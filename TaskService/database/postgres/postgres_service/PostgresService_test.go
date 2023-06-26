package postgresservice_test

import (
	"database/sql"
	"errors"
	"log"
	"regexp"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	sqlmock "github.com/zhashkevych/go-sqlxmock"

	"taskServer/database"
	postgresservice "taskServer/database/postgres/postgres_service"
	"taskServer/model"
)

var _ = Describe("PostgresService", func() {
	var (
		mock sqlmock.Sqlmock
		db   *sqlx.DB
		err  error
		r    database.DatabaseService
	)

	BeforeEach(OncePerOrdered, func() {
		db, mock, err = sqlmock.Newx()
		if err != nil {
			log.Fatalln(err)
		}
		r = postgresservice.NewPostgresService(db)
	})

	Context("postgres service test", func() {
		Context("testing GetAllTasksWithoutAnswers", func() {
			It("regular", func() {
				rows := sqlmock.NewRows([]string{"id", "quantity", "heights", "price"}).
					AddRow(1, 2, pq.Int64Array([]int64{1, 2}), 500).
					AddRow(14, 2, pq.Int64Array([]int64{1, 2}), 500)
				mockBehavior := func(page int) {
					mock.ExpectQuery("SELECT id,quantity,heights,price FROM tasks LIMIT \\$1 OFFSET \\$2").
						WithArgs(postgresservice.TASKS_PER_PAGE, (page-1)*postgresservice.TASKS_PER_PAGE).
						WillReturnRows(rows)
				}

				mockBehavior(16)
				Expect(r.GetAllTasksWithoutAnswers(16)).To(Equal([]model.TaskWithoutAnswer{
					{Id: 1, Count: 2, Heights: []int64{1, 2}, Price: 500},
					{Id: 14, Count: 2, Heights: []int64{1, 2}, Price: 500},
				}))
			})
			It("error", func() {
				rows := sqlmock.NewRows([]string{"id", "quantity", "heights", "price"}).
					AddRow(1, 2, pq.Int64Array([]int64{1, 2}), 500).
					AddRow(14, 2, pq.Int64Array([]int64{1, 2}), 500)
				mockBehavior := func(page int) {
					mock.ExpectQuery("SELECT id,quantity,heights,price FROM tasks LIMIT \\$1 OFFSET \\$2").
						WithArgs(postgresservice.TASKS_PER_PAGE, (page-1)*postgresservice.TASKS_PER_PAGE).
						WillReturnRows(rows).
						WillReturnError(errors.New("some error"))
				}

				mockBehavior(1)
				_, err := r.GetAllTasksWithoutAnswers(1)
				Expect(err).ShouldNot(Succeed())
			})
		})

		Context("testing GetAllTasks", func() {
			It("regular", func() {
				rows := sqlmock.NewRows([]string{"id", "quantity", "heights", "price", "answer"}).
					AddRow(1, 2, pq.Int64Array([]int64{1, 2}), 500, 4)
				mockBehavior := func() {
					mock.ExpectQuery("SELECT (.+) FROM tasks").
						WillReturnRows(rows)
				}

				mockBehavior()

				Expect(r.GetAllTasks()).To(Equal([]model.Task{
					{Id: 1, Count: 2, Heights: []int64{1, 2}, Price: 500, Answer: 4},
				}))
			})
			It("error", func() {
				rows := sqlmock.NewRows([]string{"id", "quantity", "heights", "price", "answer"}).
					AddRow(1, 2, pq.Int64Array([]int64{1, 2}), 500, 4)
				mockBehavior := func() {
					mock.ExpectQuery("SELECT (.+) FROM tasks").
						WillReturnRows(rows).
						WillReturnError(errors.New("some error"))
				}

				mockBehavior()

				_, err := r.GetAllTasks()

				Expect(err).ShouldNot(Succeed())
			})
		})

		Context("testing CreateTask", func() {
			It("regular", func() {
				task := model.Task{
					Count:   3,
					Heights: []int64{1, 2, 3},
					Price:   500,
					Answer:  4,
				}

				mockBehavior := func(task model.Task) {
					mock.ExpectExec("INSERT INTO tasks").
						WithArgs(task.Count, pq.Int64Array(task.Heights), task.Price, task.Answer).
						WillReturnResult(sqlmock.NewResult(1, 2))
				}

				mockBehavior(task)
				Expect(r.CreateTask(task)).Should(Succeed())
			})
			It("error", func() {
				task := model.Task{
					Count:   3,
					Heights: []int64{1, 2, 3},
					Price:   500,
					Answer:  4,
				}

				mockBehavior := func(task model.Task) {
					mock.ExpectExec("INSERT INTO tasks").
						WithArgs(task.Count, pq.Int64Array(task.Heights), task.Price, task.Answer).
						WillReturnResult(sqlmock.NewResult(1, 2)).
						WillReturnError(errors.New("some error"))
				}

				mockBehavior(task)
				Expect(r.CreateTask(task)).ShouldNot(Succeed())
			})
		})

		Context("testing CheckAndGetTask", func() {
			It("regular", func() {
				task := model.Task{
					Id:      1,
					Count:   4,
					Heights: pq.Int64Array([]int64{1, 2, 3, 4}),
					Price:   500,
					Answer:  2,
				}
				result := sqlmock.NewRows([]string{"id", "quantity", "heights", "price", "answer"}).
					AddRow(1, 4, pq.Int64Array([]int64{1, 2, 3, 4}), 500, 2)

				mockBehavior := func(username string, id int) {
					mock.ExpectQuery(regexp.QuoteMeta("SELECT id,quantity,heights,price,answer FROM tasks LEFT JOIN userOrders ON tasks.id=userOrders.orderId WHERE id=$1 AND NOT EXISTS (SELECT 1 FROM userOrders WHERE orderId = $1 AND username = $2)")).
						WithArgs(id, username).
						WillReturnRows(result)
				}

				mockBehavior("root", 1)
				Expect(r.CheckAndGetTask("root", 1)).To(Equal(task))
			})

			It("error", func() {
				result := sqlmock.NewRows([]string{"id", "quantity", "heights", "price", "answer"}).
					AddRow(1, 4, pq.Int64Array([]int64{1, 2, 3, 4}), 500, 2).
					RowError(0, sql.ErrNoRows)

				mockBehavior := func(username string, id int) {
					mock.ExpectQuery(regexp.QuoteMeta("SELECT id,quantity,heights,price,answer FROM tasks LEFT JOIN userOrders ON tasks.id=userOrders.orderId WHERE id=$1 AND NOT EXISTS (SELECT 1 FROM userOrders WHERE orderId = $1 AND username = $2)")).
						WithArgs(id, username).
						WillReturnRows(result)
				}

				mockBehavior("root", 1)
				_, err := r.CheckAndGetTask("root", 1)
				Expect(err).ShouldNot(Succeed())
			})
		})

		Context("testing ChangeTaskPrice", func() {
			It("regular", func() {
				mockBehavior := func(id int, price int) {
					mock.ExpectExec(regexp.QuoteMeta("UPDATE tasks SET price=$1 WHERE id=$2")).
						WithArgs(price, id).
						WillReturnResult(sqlmock.NewResult(1, 1))
				}

				mockBehavior(1, 400)
				Expect(r.ChangeTaskPrice(1, 400)).Should(Succeed())
			})
			It("error", func() {
				mockBehavior := func(id int, price int) {
					mock.ExpectExec(regexp.QuoteMeta("UPDATE tasks SET price=$1 WHERE id=$2")).
						WithArgs(price, id).
						WillReturnResult(sqlmock.NewResult(1, 1)).
						WillReturnError(errors.New("some error"))
				}

				mockBehavior(1, 400)
				Expect(r.ChangeTaskPrice(1, 400)).ShouldNot(Succeed())
			})
			It("error unknow task", func() {
				mockBehavior := func(id int, price int) {
					mock.ExpectExec(regexp.QuoteMeta("UPDATE tasks SET price=$1 WHERE id=$2")).
						WithArgs(price, id).
						WillReturnResult(sqlmock.NewResult(1, 2))
				}

				mockBehavior(1, 400)
				Expect(r.ChangeTaskPrice(1, 400)).ShouldNot(Succeed())
			})
		})

		Context("testing DeleteTask", func() {
			It("regular", func() {
				mockBehavior := func(id int) {
					mock.ExpectBegin()

					mock.ExpectExec(regexp.QuoteMeta("DELETE FROM tasks WHERE id=$1")).
						WithArgs(id).
						WillReturnResult(sqlmock.NewResult(1, 1))

					mock.ExpectExec(regexp.QuoteMeta("DELETE FROM userOrders WHERE orderId=$1")).
						WithArgs(id).
						WillReturnResult(sqlmock.NewResult(1, 1))

					mock.ExpectCommit()
					mock.ExpectRollback()
				}

				mockBehavior(1)
				Expect(r.DeleteTask(1)).Should(Succeed())
			})
			It("error tasks delete", func() {
				mockBehavior := func(id int) {
					mock.ExpectBegin()

					mock.ExpectExec(regexp.QuoteMeta("DELETE FROM tasks WHERE id=$1")).
						WithArgs(id).
						WillReturnResult(sqlmock.NewResult(1, 1)).
						WillReturnError(errors.New("some error"))

					mock.ExpectExec(regexp.QuoteMeta("DELETE FROM userOrders WHERE orderId=$1")).
						WithArgs(id).
						WillReturnResult(sqlmock.NewResult(1, 1))

					mock.ExpectCommit()
					mock.ExpectRollback()
				}

				mockBehavior(1)
				Expect(r.DeleteTask(1)).ShouldNot(Succeed())
			})
			It("error userOrders delete", func() {
				mockBehavior := func(id int) {
					mock.ExpectBegin()

					mock.ExpectExec(regexp.QuoteMeta("DELETE FROM tasks WHERE id=$1")).
						WithArgs(id).
						WillReturnResult(sqlmock.NewResult(1, 1))

					mock.ExpectExec(regexp.QuoteMeta("DELETE FROM userOrders WHERE orderId=$1")).
						WithArgs(id).
						WillReturnResult(sqlmock.NewResult(1, 1)).
						WillReturnError(errors.New("some error"))

					mock.ExpectCommit()
					mock.ExpectRollback()
				}

				mockBehavior(1)
				Expect(r.DeleteTask(1)).ShouldNot(Succeed())
			})
		})

		Context("testing GetAllTasksOfUser", func() {
			It("regular", func() {
				mockBehavior := func(username string, page int) {
					rows := sqlmock.NewRows([]string{"id", "quantity", "heights", "price", "answer"}).
						AddRow(1, 2, pq.Int64Array{1, 2}, 500, 2).
						AddRow(2, 2, pq.Int64Array{1, 2}, 500, 2)

					mock.ExpectQuery(regexp.QuoteMeta("SELECT id,quantity,heights,price,answer FROM userOrders LEFT JOIN tasks ON tasks.id=userOrders.orderId WHERE userOrders.username=$1 LIMIT $2 OFFSET $3")).
						WithArgs(username, postgresservice.TASKS_PER_PAGE, (page-1)*postgresservice.TASKS_PER_PAGE).
						WillReturnRows(rows)
				}

				mockBehavior("root", 1)
				Expect(r.GetAllTasksOfUser("root", 1)).To(Equal([]model.Task{
					{Id: 1, Count: 2, Heights: []int64{1, 2}, Price: 500, Answer: 2},
					{Id: 2, Count: 2, Heights: []int64{1, 2}, Price: 500, Answer: 2},
				}))
			})
			It("error", func() {
				mockBehavior := func(username string, page int) {
					rows := sqlmock.NewRows([]string{"id", "quantity", "heights", "price", "answer"}).
						AddRow(1, 2, pq.Int64Array{1, 2}, 500, 2).
						AddRow(2, 2, pq.Int64Array{1, 2}, 500, 2)

					mock.ExpectQuery(regexp.QuoteMeta("SELECT id,quantity,heights,price,answer FROM userOrders LEFT JOIN tasks ON tasks.id=userOrders.orderId WHERE userOrders.username=$1 LIMIT $2 OFFSET $3")).
						WithArgs(username, postgresservice.TASKS_PER_PAGE, (page-1)*postgresservice.TASKS_PER_PAGE).
						WillReturnRows(rows).
						WillReturnError(errors.New("some error"))
				}

				mockBehavior("root", 1)
				_, err := r.GetAllTasksOfUser("root", 1)
				Expect(err).ShouldNot(Succeed())
			})
		})

		Context("testing DeleteAllTasksOfUser", func() {
			It("regular", func() {
				mockBehavior := func(username string) {
					mock.ExpectExec(regexp.QuoteMeta("DELETE FROM userOrders WHERE username=$1")).
						WithArgs(username).
						WillReturnResult(sqlmock.NewResult(1, 1))
				}

				mockBehavior("root")
				Expect(r.DeleteAllTasksOfUser("root")).Should(Succeed())
			})
			It("error", func() {
				mockBehavior := func(username string) {
					mock.ExpectExec(regexp.QuoteMeta("DELETE FROM userOrders WHERE username=$1")).
						WithArgs(username).
						WillReturnResult(sqlmock.NewResult(1, 1)).
						WillReturnError(errors.New("some error"))
				}

				mockBehavior("root")
				Expect(r.DeleteAllTasksOfUser("root")).ShouldNot(Succeed())
			})
		})

		Context("testing BuyTaskAnswer", func() {
			It("regular", func() {
				mockBehavior := func(task model.UsersPurchase) {
					mock.ExpectExec(regexp.QuoteMeta("INSERT INTO userOrders(username,orderId) VALUES($1,$2)")).
						WithArgs(task.Username, task.OrderId).
						WillReturnResult(sqlmock.NewResult(1, 1))
				}

				task := model.UsersPurchase{
					Username: "root",
					OrderId: 1,
				}
				mockBehavior(task)
				Expect(r.BuyTaskAnswer(task)).Should(Succeed())
			})

			It("error", func() {
				mockBehavior := func(task model.UsersPurchase) {
					mock.ExpectExec(regexp.QuoteMeta("INSERT INTO userOrders(username,orderId) VALUES($1,$2)")).
						WithArgs(task.Username, task.OrderId).
						WillReturnResult(sqlmock.NewResult(1, 1)).
						WillReturnError(errors.New("some error"))
				}

				task := model.UsersPurchase{
					Username: "root",
					OrderId: 1,
				}
				mockBehavior(task)
				Expect(r.BuyTaskAnswer(task)).ShouldNot(Succeed())
			})
		})
		AfterEach(OncePerOrdered, func() {
			db.Close()
		})
	})
})
