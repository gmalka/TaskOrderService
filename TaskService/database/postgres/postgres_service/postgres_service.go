package postgresservice

import (
	"database/sql"
	"fmt"
	"taskServer/database"
	"taskServer/model"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

const TASKS_PER_PAGE = 10

type postgresService struct {
	db *sqlx.DB
}

func NewPostgresService(db *sqlx.DB) database.DatabaseService {
	return postgresService{db: db}
}

func (p postgresService) GetAllTasksWithoutAnswers(page int) ([]model.TaskWithoutAnswer, error) {
	var tasks []model.TaskWithoutAnswer

	if page <= 0 {
		page = 1
	}

	result, err := p.db.Query("SELECT id,quantity,heights,price FROM tasks LIMIT $1 OFFSET $2", TASKS_PER_PAGE, (page-1)*TASKS_PER_PAGE)
	if err != nil {
		return nil, fmt.Errorf("can't get tasks: %s", err)
	}

	tasks = make([]model.TaskWithoutAnswer, 0, 10)

	for result.Next() {
		var task model.TaskWithoutAnswer

		err = result.Scan(&task.Id, &task.Count, (*pq.Int64Array)(&task.Heights), &task.Price)
		if err != nil {
			return nil, fmt.Errorf("can't get tasks: %s", err)
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (p postgresService) GetAllTasks() ([]model.Task, error) {
	var tasks []model.Task
	result, err := p.db.Query("SELECT * FROM tasks")
	if err != nil {
		return nil, fmt.Errorf("can't get tasks: %s", err)
	}

	tasks = make([]model.Task, 0, 10)

	for result.Next() {
		var task model.Task

		err = result.Scan(&task.Id, &task.Count, (*pq.Int64Array)(&task.Heights), &task.Price, &task.Answer)
		if err != nil {
			return nil, fmt.Errorf("can't get tasks: %s", err)
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (p postgresService) CreateTask(task model.Task) error {
	_, err := p.db.Exec("INSERT INTO tasks(quantity,heights,price,answer) VALUES($1,$2,$3,$4)", task.Count, pq.Array(task.Heights), task.Price, task.Answer)
	if err != nil {
		return fmt.Errorf("can't create task: %s", err)
	}

	return nil
}

func (p postgresService) CheckAndGetTask(username string, id int) (model.Task, error) {
	var task model.Task

	err := p.db.QueryRow("SELECT id,quantity,heights,price,answer FROM tasks LEFT JOIN userOrders ON tasks.id=userOrders.orderId WHERE id=$1 AND NOT EXISTS (SELECT 1 FROM userOrders WHERE orderId = $1 AND username = $2)",
		id, username).Scan(&task.Id, &task.Count, (*pq.Int64Array)(&task.Heights), &task.Price, &task.Answer)
	if err == sql.ErrNoRows {
		return task, fmt.Errorf("the job has already been purchased by this user\n")
	}
	if err != nil {
		return task, fmt.Errorf("can't get for task %d: %s", id, err)
	}

	return task, nil
}

func (p postgresService) ChangeTaskPrice(id int, price int) error {
	res, err := p.db.Exec("UPDATE tasks SET price=$1 WHERE id=$2", price, id)
	if err != nil {
		return fmt.Errorf("can't set for task %d price to %d: %s", id, price, err)
	}

	if i, _ := res.RowsAffected(); i != 1 {
		return fmt.Errorf("unknow task id: %d", id)
	}

	return nil
}

func (p postgresService) DeleteTask(id int) error {
	tx, err := p.db.Begin()
	if err != nil {
		return fmt.Errorf("can't delete task %d: %s", id, err)
	}
	defer tx.Rollback()

	_, err = tx.Exec("DELETE FROM tasks WHERE id=$1", id)
	if err != nil {
		return fmt.Errorf("can't delete task %d: %s", id, err)
	}

	_, err = tx.Exec("DELETE FROM userOrders WHERE orderId=$1", id)
	if err != nil {
		return fmt.Errorf("can't delete task %d: %s", id, err)
	}

	tx.Commit()
	return nil
}

func (p postgresService) GetAllTasksOfUser(username string, page int) ([]model.Task, error) {
	var tasks []model.Task

	if page <= 0 {
		page = 1
	}

	result, err := p.db.Query("SELECT id,quantity,heights,price,answer FROM userOrders LEFT JOIN tasks ON tasks.id=userOrders.orderId WHERE userOrders.username=$1 LIMIT $2 OFFSET $3",
		username, TASKS_PER_PAGE, (page-1)*TASKS_PER_PAGE)
	if err != nil {
		return nil, fmt.Errorf("can't get orders of user %s: %s", username, err)
	}

	tasks = make([]model.Task, 0, 10)

	for result.Next() {
		var task model.Task

		err = result.Scan(&task.Id, &task.Count, (*pq.Int64Array)(&task.Heights), &task.Price, &task.Answer)
		if err != nil {
			return nil, fmt.Errorf("can't get orders of user %s: %s", username, err)
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (p postgresService) DeleteAllTasksOfUser(username string) error {
	_, err := p.db.Exec("DELETE FROM userOrders WHERE username=$1", username)
	if err != nil {
		return fmt.Errorf("can't delete orders for user %s: %s", username, err)
	}

	return nil
}

func (p postgresService) BuyTaskAnswer(task model.UsersPurchase) error {
	_, err := p.db.Exec("INSERT INTO userOrders(username,orderId) VALUES($1,$2)", task.Username, task.OrderId)
	if err != nil {
		return fmt.Errorf("can't write order %d of user %s : %s", task.OrderId, task.Username, err)
	}

	return nil
}
