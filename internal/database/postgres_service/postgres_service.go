package postgresservice

import (
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
		log.Println("Error while creating new user: ", user.User.Username)
		return err
	}

	return nil
}

func (p postgresService) GetByUsername(username string) (model.UserInfoWithRoleruct, error) {
	var userInfo model.UserInfoWithRoleruct

	err := p.db.QueryRow("SELECT firstname,lastname,surname,group,balance,role FROM users WHERE username=$1", username).Scan(&userInfo.Info.Firstname,
		&userInfo.Info.Lastname, &userInfo.Info.Surname, &userInfo.Info.Group, &userInfo.Info.Balance, &userInfo.Role)
	if err != nil {
		log.Println("Error while geting info about user: ", username)
		return userInfo, err
	}

	return userInfo, nil
}

func (p postgresService) GetAllUsers() ([]string, error) {
	var users []string

	users = make([]string, 0, 10)
	result, err := p.db.Query("SELECT username FROM users")
	if err != nil {
		log.Println("Error while geting all users user")
		return nil, err
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
		log.Println("Error while Updating user: ", user.Username)
		return err
	}
	return nil
}

func (p postgresService) Delete(username string) error {
	_, err := p.db.Exec("DELETE FROM users WHERE username=$1", username)
	if err != nil {
		log.Println("Error while Deleting user: ", username)
		return err
	}
	return err
}
