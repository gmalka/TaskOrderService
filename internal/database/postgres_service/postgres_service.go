package postgresservice

import (
	"fmt"
	"log"
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

func (p postgresService) Create(user model.UserWithRole) error {
	_, err := p.db.Exec("INSERT INTO users(username,password,firstname,lastname,surname,group,role) VALUES($1,$2,$3,$4,$5,$6,$7)",
		user.User.Username, user.User.Password, user.User.Info.Firstname, user.User.Info.Lastname, user.User.Info.Surname, user.User.Info.Group, user.Role)
	if err != nil {
		return fmt.Errorf("cant create user: %s", user.User.Username)
	}

	return nil
}

func (p postgresService) GetByUsername(username string) (model.UserWithRole, error) {
	var user model.UserWithRole

	err := p.db.QueryRow("SELECT password,firstname,lastname,surname,group,balance,role FROM users WHERE username=$1", username).Scan(&user.User.Password,&user.User.Info.Firstname,
		&user.User.Info.Lastname, &user.User.Info.Surname, &user.User.Info.Group, &user.User.Info.Balance, &user.Role)
	if err != nil {
		return user, fmt.Errorf("cant get info about user: %s", username)
	}

	user.User.Username = username
	return user, nil
}

func (p postgresService) GetPassword(username string) (string, error) {
	var result string

	err := p.db.QueryRow("SELECT password FROM users WHERE username=$1", username).Scan(&result)

	return result, err
}

func (p postgresService) GetAllUsers() ([]string, error) {
	var users []string

	users = make([]string, 0, 10)
	result, err := p.db.Query("SELECT username FROM users")
	if err != nil {
		return nil, fmt.Errorf("cant get all users")
	}

	defer result.Close()
	for result.Next() {
		str := ""

		err = result.Scan(&str)
		if err != nil {
			log.Println("Error while scanning")
			return nil, err
		}
		users = append(users, str)
	}

	return users, nil
}

func (p postgresService) Update(user model.User) error {
	_, err := p.db.Exec("UPDATE users SET password=$1,firstname=$2,lastname=$3,surname=$4,group=$5 WHERE username=$6",
		user.Password, user.Info.Firstname, user.Info.Lastname, user.Info.Surname, user.Info.Group, user.Username)
	if err != nil {
		return fmt.Errorf("cant update user: %s", user.Username)
	}
	return nil
}

func (p postgresService) Delete(username string) error {
	_, err := p.db.Exec("DELETE FROM users WHERE username=$1", username)
	if err != nil {
		log.Println("Error while Deleting user: ", username)
		return fmt.Errorf("cant delete user: %s", username)
	}
	return err
}
