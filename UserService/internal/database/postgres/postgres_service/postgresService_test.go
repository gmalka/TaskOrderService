package postgresservice_test

import (
	"errors"
	"log"
	"regexp"

	"github.com/jmoiron/sqlx"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	sqlmock "github.com/zhashkevych/go-sqlxmock"

	"userService/internal/database"
	postgresservice "userService/internal/database/postgres/postgres_service"
	"userService/internal/model"
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
		Context("testing UpdateBalance", func() {
			It("regular", func() {
				mockBehavior := func(change int, username string) {
					mock.ExpectExec(regexp.QuoteMeta("UPDATE users SET balance=balance+$1 WHERE username=$2")).
						WithArgs(change, username).
						WillReturnResult(sqlmock.NewResult(1, 1))
				}

				mockBehavior(1, "root")
				Expect(r.UpdateBalance("root", 1)).Should(Succeed())
			})

			It("error", func() {
				mockBehavior := func(change int, username string) {
					mock.ExpectExec(regexp.QuoteMeta("UPDATE users SET balance=balance+$1 WHERE username=$2")).
						WithArgs(change, username).
						WillReturnResult(sqlmock.NewResult(1, 1)).
						WillReturnError(errors.New("some error"))
				}

				mockBehavior(1, "root")
				Expect(r.UpdateBalance("root", 1)).ShouldNot(Succeed())
			})
		})

		Context("testing TryToBuyTask", func() {
			It("regular", func() {
				mockBehavior := func(change int, username string) {
					mock.ExpectBegin()

					row := sqlmock.NewRows([]string{"balance"}).AddRow("500")

					mock.ExpectQuery(regexp.QuoteMeta("SELECT balance FROM users WHERE username=$1 FOR UPDATE")).
						WithArgs(username).
						WillReturnRows(row)

					mock.ExpectExec(regexp.QuoteMeta("UPDATE users SET balance=$1 WHERE username=$2")).
						WithArgs(500-change, username).
						WillReturnResult(sqlmock.NewResult(1, 1))

					mock.ExpectCommit()
					mock.ExpectRollback()
				}

				mockBehavior(100, "root")
				Expect(r.TryToBuyTask("root", 100)).Should(Succeed())
			})
			It("error 1", func() {
				mockBehavior := func(change int, username string) {
					mock.ExpectBegin()

					row := sqlmock.NewRows([]string{"balance"}).AddRow("500")

					mock.ExpectQuery(regexp.QuoteMeta("SELECT balance FROM users WHERE username=$1 FOR UPDATE")).
						WithArgs(username).
						WillReturnRows(row).
						WillReturnError(errors.New("some error"))

					mock.ExpectExec(regexp.QuoteMeta("UPDATE users SET balance=$1 WHERE username=$2")).
						WithArgs(500-change, username).
						WillReturnResult(sqlmock.NewResult(1, 1))

					mock.ExpectCommit()
					mock.ExpectRollback()
				}

				mockBehavior(100, "root")
				Expect(r.TryToBuyTask("root", 100)).ShouldNot(Succeed())
			})
			It("error 2", func() {
				mockBehavior := func(change int, username string) {
					mock.ExpectBegin()

					row := sqlmock.NewRows([]string{"balance"}).AddRow("500")

					mock.ExpectQuery(regexp.QuoteMeta("SELECT balance FROM users WHERE username=$1 FOR UPDATE")).
						WithArgs(username).
						WillReturnRows(row)

					mock.ExpectExec(regexp.QuoteMeta("UPDATE users SET balance=$1 WHERE username=$2")).
						WithArgs(500-change, username).
						WillReturnResult(sqlmock.NewResult(1, 1)).
						WillReturnError(errors.New("some error"))

					mock.ExpectCommit()
					mock.ExpectRollback()
				}

				mockBehavior(100, "root")
				Expect(r.TryToBuyTask("root", 100)).ShouldNot(Succeed())
			})
			It("error 2", func() {
				mockBehavior := func(change int, username string) {
					mock.ExpectBegin()

					row := sqlmock.NewRows([]string{"balance"}).AddRow("500")

					mock.ExpectQuery(regexp.QuoteMeta("SELECT balance FROM users WHERE username=$1 FOR UPDATE")).
						WithArgs(username).
						WillReturnRows(row)

					mock.ExpectExec(regexp.QuoteMeta("UPDATE users SET balance=$1 WHERE username=$2")).
						WithArgs(500-change, username).
						WillReturnResult(sqlmock.NewResult(1, 1))

					mock.ExpectCommit()
					mock.ExpectRollback()
				}

				mockBehavior(100, "root")
				Expect(r.TryToBuyTask("root", 1501)).ShouldNot(Succeed())
			})
		})

		Context("testing Create", func() {
			It("regular", func() {
				mockBehavior := func(username, password, firstname, lastname, surname, user_group, role string) {
					mock.ExpectExec(regexp.QuoteMeta("INSERT INTO users(username,password,firstname,lastname,surname,user_group,role,balance) VALUES($1,$2,$3,$4,$5,$6,$7,0)")).
						WithArgs(username, password, firstname, lastname, surname, user_group, role).
						WillReturnResult(sqlmock.NewResult(1, 1))
				}

				mockBehavior("root", "1234", "m", "m", "m", "g-12", "admin")
				Expect(r.Create(model.UserWithRole{
					User: model.User{
						Username: "root",
						Password: "1234",
						Info: model.UserInfo{
							Firstname: "m",
							Lastname:  "m",
							Surname:   "m",
							Group:     "g-12",
							Balance:   0,
						},
					},
					Role: "admin",
				})).Should(Succeed())
			})
			It("error", func() {
				mockBehavior := func(username, password, firstname, lastname, surname, user_group, role string) {
					mock.ExpectExec(regexp.QuoteMeta("INSERT INTO users(username,password,firstname,lastname,surname,user_group,role,balance) VALUES($1,$2,$3,$4,$5,$6,$7,0)")).
						WithArgs(username, password, firstname, lastname, surname, user_group, role).
						WillReturnResult(sqlmock.NewResult(1, 1)).
						WillReturnError(errors.New("some error"))
				}

				mockBehavior("root", "1234", "m", "m", "m", "g-12", "admin")
				Expect(r.Create(model.UserWithRole{
					User: model.User{
						Username: "root",
						Password: "1234",
						Info: model.UserInfo{
							Firstname: "m",
							Lastname:  "m",
							Surname:   "m",
							Group:     "g-12",
							Balance:   0,
						},
					},
					Role: "admin",
				})).ShouldNot(Succeed())
			})
		})
		Context("testing GetByUsername", func() {
			It("regular", func() {
				mockBehavior := func(username string) {
					row := sqlmock.NewRows([]string{"password","firstname","lastname","surname","user_group","balance","role"}).
						AddRow("1234", "m", "m", "m", "g-12", "500", "admin")

					mock.ExpectQuery(regexp.QuoteMeta("SELECT password,firstname,lastname,surname,user_group,balance,role FROM users WHERE username=$1")).
						WithArgs(username).
						WillReturnRows(row)
				}

				mockBehavior("root")
				Expect(r.GetByUsername("root")).To(Equal(
					model.UserWithRole{
						User: model.User{
							Username: "root",
							Password: "1234",
							Info: model.UserInfo{
								Firstname: "m",
								Lastname:  "m",
								Surname:   "m",
								Group:     "g-12",
								Balance:   500,
							},
						},
						Role: "admin",
					},
				))
			})
			It("error", func() {
				mockBehavior := func(username string) {
					row := sqlmock.NewRows([]string{"password","firstname","lastname","surname","user_group","balance","role"}).
						AddRow("1234", "m", "m", "m", "g-12", "500", "admin")

					mock.ExpectQuery(regexp.QuoteMeta("SELECT password,firstname,lastname,surname,user_group,balance,role FROM users WHERE username=$1")).
						WithArgs(username).
						WillReturnRows(row).
						WillReturnError(errors.New("some error"))
				}

				mockBehavior("root")
				_, err := r.GetByUsername("root")
				Expect(err).ShouldNot(Succeed())
			})
		})

		Context("testing GetAllUsers", func() {
			It("regular", func() {
				row := sqlmock.NewRows([]string{"username"}).AddRow("root").AddRow("gmalka").AddRow("Misha")
				mockBehavior := func() {
					mock.ExpectQuery(regexp.QuoteMeta("SELECT username FROM users")).
						WillReturnRows(row)
				}

				mockBehavior()
				Expect(r.GetAllUsers()).To(Equal([]string{"root", "gmalka", "Misha"}))
			})

			It("error", func() {
				row := sqlmock.NewRows([]string{"username"}).AddRow("root").AddRow("gmalka").AddRow("Misha")
				mockBehavior := func() {
					mock.ExpectQuery(regexp.QuoteMeta("SELECT username FROM users")).
						WillReturnRows(row).
						WillReturnError(errors.New("some error"))
				}

				mockBehavior()
				_, err := r.GetAllUsers()
				Expect(err).ShouldNot(Succeed())
			})
		})

		Context("testing Update", func() {
			It("regular", func() {
				user := model.UserForUpdate{Username: "root", Password: "1234", Info: model.UserInfoForUpdate{Firstname: "m", Lastname: "m", Surname: "m", Group: "p-14"}}
				mockBehavior := func(user model.UserForUpdate) {
					mock.ExpectExec(regexp.QuoteMeta("UPDATE users SET password=$1,firstname=$2,lastname=$3,surname=$4,user_group=$5 WHERE username=$6")).
						WithArgs(user.Password, user.Info.Firstname, user.Info.Lastname, user.Info.Surname, user.Info.Group, user.Username).
						WillReturnResult(sqlmock.NewResult(1, 2))
				}

				mockBehavior(user)
				Expect(r.Update(user)).Should(Succeed())
			})
			It("regular", func() {
				user := model.UserForUpdate{Username: "root", Password: "1234", Info: model.UserInfoForUpdate{Firstname: "m", Lastname: "m", Surname: "m", Group: "p-14"}}
				mockBehavior := func(user model.UserForUpdate) {
					mock.ExpectExec(regexp.QuoteMeta("UPDATE users SET password=$1,firstname=$2,lastname=$3,surname=$4,user_group=$5 WHERE username=$6")).
						WithArgs(user.Password, user.Info.Firstname, user.Info.Lastname, user.Info.Surname, user.Info.Group, user.Username).
						WillReturnResult(sqlmock.NewResult(1, 2)).
						WillReturnError(errors.New("some error"))
				}

				mockBehavior(user)
				Expect(r.Update(user)).ShouldNot(Succeed())
			})
		})

		Context("testing Delete", func() {
			It("regular", func() {
				mockBehavior := func(username string) {
					mock.ExpectExec(regexp.QuoteMeta("DELETE FROM users WHERE username=$1")).
						WithArgs(username).
						WillReturnResult(sqlmock.NewResult(1, 1))
				}

				mockBehavior("root")
				Expect(r.Delete("root")).Should(Succeed())
			})
			It("regular", func() {
				mockBehavior := func(username string) {
					mock.ExpectExec(regexp.QuoteMeta("DELETE FROM users WHERE username=$1")).
						WithArgs(username).
						WillReturnResult(sqlmock.NewResult(1, 1)).
						WillReturnError(errors.New("some error"))
				}

				mockBehavior("root")
				Expect(r.Delete("root")).ShouldNot(Succeed())
			})
		})
	})
})
