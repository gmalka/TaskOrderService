package postgresservice

import (
	"fmt"
	"userService/internal/database"
	"userService/internal/model"

	"github.com/jmoiron/sqlx"
)

type postgresService struct {
	db *sqlx.DB
}

func NewPostgresService(db *sqlx.DB) database.DatabaseService {
	return postgresService{db: db}
}

func (p postgresService) UpdateBalance(username string, change int) error {
	res, err := p.db.Exec("UPDATE users SET balance=balance+$1 WHERE username=$2", change, username)
	if err != nil {
		return fmt.Errorf("can't update users balance %s: %v", username, err)
	}
	if i, _ := res.RowsAffected(); i < 1 {
		return fmt.Errorf("can't find usere %s", username)
	}

	return nil
}

func (p postgresService) TryToBuyTask(username string, price int) error {
	var balance int

	tx, err := p.db.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = tx.QueryRow("SELECT balance FROM users WHERE username=$1 FOR UPDATE", username).Scan(&balance)
	if err != nil {
		return fmt.Errorf("can't get balance of user %s: %v", username, err)
	}

	if balance-price < -1000 {
		return fmt.Errorf("not enought money for operation for user %s: %v", username, err)
	}

	_, err = tx.Exec("UPDATE users SET balance=$1 WHERE username=$2", balance-price, username)
	if err != nil {
		return fmt.Errorf("can't update balance of user %s: %v", username, err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("can't commit changes for user %s: %v", username, err)
	}

	return nil
}

func (p postgresService) Create(user model.UserWithRole) error {
	_, err := p.db.Exec("INSERT INTO users(username,password,firstname,lastname,surname,user_group,role,balance) VALUES($1,$2,$3,$4,$5,$6,$7,0)",
		user.User.Username, user.User.Password, user.User.Info.Firstname, user.User.Info.Lastname, user.User.Info.Surname, user.User.Info.Group, user.Role)
	if err != nil {
		return fmt.Errorf("can't create user %s: %v", user.User.Username, err)
	}

	return nil
}

func (p postgresService) GetByUsername(username string) (model.UserWithRole, error) {
	var user model.UserWithRole

	err := p.db.QueryRow("SELECT password,firstname,lastname,surname,user_group,balance,role FROM users WHERE username=$1", username).Scan(&user.User.Password, &user.User.Info.Firstname,
		&user.User.Info.Lastname, &user.User.Info.Surname, &user.User.Info.Group, &user.User.Info.Balance, &user.Role)
	if err != nil {
		return user, fmt.Errorf("can't get info about user %s: %v", username, err)
	}

	user.User.Username = username
	return user, nil
}

func (p postgresService) GetAllUsers() ([]string, error) {
	var users []string

	users = make([]string, 0, 10)
	result, err := p.db.Query("SELECT username FROM users")
	if err != nil {
		return nil, fmt.Errorf("can't get all users: %v", err)
	}

	defer result.Close()
	for result.Next() {
		str := ""

		err = result.Scan(&str)
		if err != nil {
			return nil, fmt.Errorf("can't scan user info: %v", err)
		}
		users = append(users, str)
	}

	return users, nil
}

func (p postgresService) Update(user model.UserForUpdate) error {
	_, err := p.db.Exec("UPDATE users SET password=$1,firstname=$2,lastname=$3,surname=$4,user_group=$5 WHERE username=$6",
		user.Password, user.Info.Firstname, user.Info.Lastname, user.Info.Surname, user.Info.Group, user.Username)
	if err != nil {
		return fmt.Errorf("can't update user %s: %v", user.Username, err)
	}
	return nil
}

func (p postgresService) Delete(username string) error {
	_, err := p.db.Exec("DELETE FROM users WHERE username=$1", username)
	if err != nil {
		return fmt.Errorf("can't delete user %s: %v", username, err)
	}
	return err
}
