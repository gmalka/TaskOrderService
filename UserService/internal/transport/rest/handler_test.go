package rest_test

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"time"

	. "github.com/gcapizzi/moka"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"userService/internal/auth/tokenManager"
	"userService/internal/model"
	"userService/internal/transport/rest"
)

var _ = Describe("Handler", func() {
	var (
		db     UserControllerDouble
		mygrpc GrpcServiceDouble
		pas    PasswordMangerDouble
		tok    TokenManagerDouble
		h      rest.Handler
		//ts         *httptest.Server
	)

	BeforeEach(func() {
		loggerErr := log.New(ioutil.Discard, "ERROR:\t ", log.Lshortfile|log.Ltime)
		loggerInfo := log.New(ioutil.Discard, "INFO:\t ", log.Lshortfile|log.Ltime)
		logger := rest.Log{loggerErr, loggerInfo}

		db = NewUserControllerDouble()
		mygrpc = NewGrpcServiceDouble()
		pas = NewPasswordMangerDouble()
		tok = NewTokenManagerDouble()
		h = rest.NewHandler(db, tok, mygrpc, pas, logger)
	})

	Context("Handler testing ", func() {
		Context("test Auth", func() {
			Context("test Register", func() {
				It("regular", func() {
					AllowDouble(pas).To(ReceiveCallTo("HashPassword").With("1234").AndReturn("123456", nil))
					AllowDouble(db).To(ReceiveCallTo("CreateUser").With(model.User{
						Username: "root",
						Password: "123456",
						Info: model.UserInfo{
							Firstname: "m",
							Lastname:  "m",
							Surname:   "mm",
							Group:     "t-15",
							Balance:   0,
						},
					}).AndReturn(nil))

					res := httptest.NewRecorder()
					in := `{"username":"root","password":"1234","info":{"firstname":"m","lastname":"m","surname":"mm","group":"t-15","balance":0}}`

					req, err := http.NewRequest("POST", "/auth/register", bytes.NewReader([]byte(in)))
					h.InitRouter(false).ServeHTTP(res, req)

					Expect(err).Should(Succeed())
					Expect(res.Result().StatusCode).To(Equal(200))
					Expect(res.Header().Get("Content-Type")).To(Equal("application/json"))
					Expect(ioutil.ReadAll(res.Body)).To(Equal([]byte(`{"message":"success register"}`)))
				})
				It("error hash password", func() {
					AllowDouble(pas).To(ReceiveCallTo("HashPassword").With("1234").AndReturn("123456", errors.New("some error")))

					res := httptest.NewRecorder()
					in := `{"username":"root","password":"1234","info":{"firstname":"m","lastname":"m","surname":"mm","group":"t-15","balance":0}}`

					req, err := http.NewRequest("POST", "/auth/register", bytes.NewReader([]byte(in)))
					h.InitRouter(false).ServeHTTP(res, req)

					Expect(err).Should(Succeed())
					Expect(res.Result().StatusCode).To(Equal(500))
				})
				It("error create user", func() {
					AllowDouble(pas).To(ReceiveCallTo("HashPassword").With("1234").AndReturn("123456", nil))
					AllowDouble(db).To(ReceiveCallTo("CreateUser").With(model.User{
						Username: "root",
						Password: "123456",
						Info: model.UserInfo{
							Firstname: "m",
							Lastname:  "m",
							Surname:   "mm",
							Group:     "t-15",
							Balance:   0,
						},
					}).AndReturn(errors.New("some error")))

					res := httptest.NewRecorder()
					in := `{"username":"root","password":"1234","info":{"firstname":"m","lastname":"m","surname":"mm","group":"t-15","balance":0}}`

					req, err := http.NewRequest("POST", "/auth/register", bytes.NewReader([]byte(in)))
					h.InitRouter(false).ServeHTTP(res, req)

					Expect(err).Should(Succeed())
					Expect(res.Result().StatusCode).To(Equal(400))
				})
			})

			Context("test LoginIn", func() {
				It("regular", func() {
					AllowDouble(db).To(ReceiveCallTo("GetUser").With("root").AndReturn(model.UserWithRole{
						User: model.User{Password: "1234"},
						Role: "admin",
					}, nil))
					AllowDouble(pas).To(ReceiveCallTo("CheckPassword").With("1234", "1234").AndReturn(nil))
					AllowDouble(tok).To(ReceiveCallTo("CreateToken").With(tokenManager.UserInfo{
						Username:  "",
						Role:      "admin",
						Firstname: "",
						Lastname:  "",
					}, time.Duration(15), 2).AndReturn("4321", nil))
					AllowDouble(tok).To(ReceiveCallTo("CreateToken").With(tokenManager.UserInfo{
						Username:  "",
						Role:      "admin",
						Firstname: "",
						Lastname:  "",
					}, time.Duration(60), 3).AndReturn("654321", nil))

					res := httptest.NewRecorder()
					in := `{"username":"root","password":"1234"}`

					req, _ := http.NewRequest("POST", "/auth/login", bytes.NewReader([]byte(in)))
					h.InitRouter(false).ServeHTTP(res, req)

					Expect(res.Result().StatusCode).To(Equal(200))
					Expect(ioutil.ReadAll(res.Body)).To(Equal([]byte(`{"Access":"4321","Refresh":"654321"}`)))
				})

				It("error get user", func() {
					AllowDouble(db).To(ReceiveCallTo("GetUser").With("root").AndReturn(model.UserWithRole{
						User: model.User{Password: "1234"},
						Role: "admin",
					}, errors.New("some error")))

					res := httptest.NewRecorder()
					in := `{"username":"root","password":"1234"}`

					req, _ := http.NewRequest("POST", "/auth/login", bytes.NewReader([]byte(in)))
					h.InitRouter(false).ServeHTTP(res, req)

					Expect(res.Result().StatusCode).To(Equal(400))
				})

				It("error check password", func() {
					AllowDouble(db).To(ReceiveCallTo("GetUser").With("root").AndReturn(model.UserWithRole{
						User: model.User{Password: "1234"},
						Role: "admin",
					}, nil))
					AllowDouble(pas).To(ReceiveCallTo("CheckPassword").With("1234", "1234").AndReturn(errors.New("some error")))

					res := httptest.NewRecorder()
					in := `{"username":"root","password":"1234"}`

					req, _ := http.NewRequest("POST", "/auth/login", bytes.NewReader([]byte(in)))
					h.InitRouter(false).ServeHTTP(res, req)

					Expect(res.Result().StatusCode).To(Equal(401))
				})
			})

			Context("test Refresh", func() {
				It("regular", func() {
					AllowDouble(tok).To(ReceiveCallTo("ParseToken").With("54321", 3).AndReturn(tokenManager.UserClaims{
						Username:  "root",
						Role:      "regular",
						Firstname: "m",
						Lastname:  "m",
					}, nil))
					AllowDouble(tok).To(ReceiveCallTo("CreateToken").With(tokenManager.UserInfo{
						Username:  "root",
						Role:      "regular",
						Firstname: "m",
						Lastname:  "m",
					}, time.Duration(15), 2).AndReturn("4321", nil))
					AllowDouble(tok).To(ReceiveCallTo("CreateToken").With(tokenManager.UserInfo{
						Username:  "root",
						Role:      "regular",
						Firstname: "m",
						Lastname:  "m",
					}, time.Duration(60), 3).AndReturn("654321", nil))

					res := httptest.NewRecorder()

					req, _ := http.NewRequest("POST", "/auth/refresh", nil)
					req.Header.Add("Authorization", "Bearer 54321")
					h.InitRouter(false).ServeHTTP(res, req)

					Expect(res.Result().StatusCode).To(Equal(200))
					Expect(ioutil.ReadAll(res.Body)).To(Equal([]byte(`{"Access":"4321","Refresh":"654321"}`)))
				})
				It("error CreateToken", func() {
					AllowDouble(tok).To(ReceiveCallTo("ParseToken").With("54321", 3).AndReturn(tokenManager.UserClaims{
						Username:  "root",
						Role:      "regular",
						Firstname: "m",
						Lastname:  "m",
					}, nil))
					AllowDouble(tok).To(ReceiveCallTo("CreateToken").With(tokenManager.UserInfo{
						Username:  "root",
						Role:      "regular",
						Firstname: "m",
						Lastname:  "m",
					}, time.Duration(15), 2).AndReturn("4321", errors.New("some error")))

					res := httptest.NewRecorder()

					req, _ := http.NewRequest("POST", "/auth/refresh", nil)
					req.Header.Add("Authorization", "Bearer 54321")
					h.InitRouter(false).ServeHTTP(res, req)

					Expect(res.Result().StatusCode).To(Equal(500))
				})
				It("error ParseToken", func() {
					AllowDouble(tok).To(ReceiveCallTo("ParseToken").With("54321", 3).AndReturn(tokenManager.UserClaims{
						Username:  "root",
						Role:      "regular",
						Firstname: "m",
						Lastname:  "m",
					}, errors.New("some error")))

					res := httptest.NewRecorder()
					req, _ := http.NewRequest("POST", "/auth/refresh", nil)
					req.Header.Add("Authorization", "Bearer 54321")
					h.InitRouter(false).ServeHTTP(res, req)

					Expect(res.Result().StatusCode).To(Equal(400))
				})
			})
		})

		Context("test Get tasks", func() {
			Context("test getUsersTasksWithoutAnswer", func() {
				It("regular", func() {
					AllowDouble(mygrpc).To(ReceiveCallTo("GetAllTasksWithoutAnswers").With(1).AndReturn([]model.TaskWithoutAnswer{
						{Id: 1, Count: 2, Heights: []int64{1, 2}, Price: 4},
					}, nil))

					res := httptest.NewRecorder()
					req, _ := http.NewRequest("GET", "/tasks/1", nil)
					h.InitRouter(false).ServeHTTP(res, req)

					Expect(res.Result().StatusCode).To(Equal(200))
					Expect(ioutil.ReadAll(res.Body)).
						To(Equal([]byte(`[{"id":1,"count":2,"heights":[1,2],"price":4}]`)))
				})
				It("error", func() {
					AllowDouble(mygrpc).To(ReceiveCallTo("GetAllTasksWithoutAnswers").With(1).AndReturn([]model.TaskWithoutAnswer{
						{Id: 1, Count: 2, Heights: []int64{1, 2}, Price: 4},
					}, errors.New("some error")))

					res := httptest.NewRecorder()
					req, _ := http.NewRequest("GET", "/tasks/1", nil)
					h.InitRouter(false).ServeHTTP(res, req)

					Expect(res.Result().StatusCode).To(Equal(400))
				})
			})
		})

		Context("test users handler", func() {
			Context("test getUsersNicknames", func() {
				It("regular", func() {
					AllowDouble(db).To(ReceiveCallTo("GetAllUsernames").With().
						AndReturn([]string{"1", "2", "3"}, nil))

					res := httptest.NewRecorder()
					req, _ := http.NewRequest("GET", "/users", nil)
					h.InitRouter(false).ServeHTTP(res, req)

					Expect(res.Result().StatusCode).To(Equal(200))
					Expect(ioutil.ReadAll(res.Body)).
						To(Equal([]byte(`["1","2","3"]`)))
				})
				It("error", func() {
					AllowDouble(db).To(ReceiveCallTo("GetAllUsernames").With().
						AndReturn(nil, errors.New("some error")))

					res := httptest.NewRecorder()
					req, _ := http.NewRequest("GET", "/users", nil)
					h.InitRouter(false).ServeHTTP(res, req)

					Expect(res.Result().StatusCode).To(Equal(400))
				})
			})
		})

		Context("test users profile handler", func() {
			Context("test getInfo", func() {
				It("regular", func() {
					AllowDouble(tok).To(ReceiveCallTo("ParseToken").With("54321", 2).AndReturn(tokenManager.UserClaims{
						Username:  "root",
						Role:      "regular",
						Firstname: "m",
						Lastname:  "m",
					}, nil))
					AllowDouble(db).To(ReceiveCallTo("GetUser").With("root").AndReturn(model.UserWithRole{
						User: model.User{
							Username: "root",
							Password: "1234",
							Info: model.UserInfo{
								Firstname: "gmalka",
							},
						},
						Role: "regular",
					}, nil))

					res := httptest.NewRecorder()
					req, _ := http.NewRequest("GET", "/users/root", nil)
					req.Header.Add("Authorization", "Bearer 54321")
					h.InitRouter(false).ServeHTTP(res, req)

					Expect(res.Result().StatusCode).To(Equal(200))
					Expect(ioutil.ReadAll(res.Body)).
						To(Equal([]byte(`{"firstname":"gmalka","lastname":"","surname":"","group":"","balance":0}`)))
				})
				It("error GetUser", func() {
					AllowDouble(tok).To(ReceiveCallTo("ParseToken").With("54321", 2).AndReturn(tokenManager.UserClaims{
						Username:  "root",
						Role:      "regular",
						Firstname: "m",
						Lastname:  "m",
					}, nil))
					AllowDouble(db).To(ReceiveCallTo("GetUser").With("root").AndReturn(model.UserWithRole{
						User: model.User{
							Username: "root",
							Password: "1234",
							Info: model.UserInfo{
								Firstname: "gmalka",
							},
						},
						Role: "regular",
					}, errors.New("some error")))

					res := httptest.NewRecorder()
					req, _ := http.NewRequest("GET", "/users/root", nil)
					req.Header.Add("Authorization", "Bearer 54321")
					h.InitRouter(false).ServeHTTP(res, req)

					Expect(res.Result().StatusCode).To(Equal(400))
				})
				It("error ParseToken", func() {
					AllowDouble(tok).To(ReceiveCallTo("ParseToken").With("54321", 2).AndReturn(tokenManager.UserClaims{
						Username:  "root",
						Role:      "regular",
						Firstname: "m",
						Lastname:  "m",
					}, errors.New("some error")))

					res := httptest.NewRecorder()
					req, _ := http.NewRequest("GET", "/users/root", nil)
					req.Header.Add("Authorization", "Bearer 54321")
					h.InitRouter(false).ServeHTTP(res, req)

					Expect(res.Result().StatusCode).To(Equal(400))
				})
			})

			Context("test getInfo", func() {
				It("regular not admin", func() {
					AllowDouble(tok).To(ReceiveCallTo("ParseToken").With("54321", 2).AndReturn(tokenManager.UserClaims{
						Username:  "root",
						Role:      "regular",
						Firstname: "m",
						Lastname:  "m",
					}, nil))
					AllowDouble(db).To(ReceiveCallTo("UpdateUser").With(model.UserForUpdate{
						Username: "root",
						Password: "1234",
						Info: model.UserInfoForUpdate{
							Firstname: "",
							Lastname:  "",
							Surname:   "",
							Group:     "",
						},
					}).AndReturn(nil))

					res := httptest.NewRecorder()
					in := `{"username":"root","password":"1234"}`
					req, _ := http.NewRequest("PUT", "/users/root", bytes.NewReader([]byte(in)))
					req.Header.Add("Authorization", "Bearer 54321")
					h.InitRouter(false).ServeHTTP(res, req)

					Expect(res.Result().StatusCode).To(Equal(200))
				})
				It("error admin", func() {
					AllowDouble(tok).To(ReceiveCallTo("ParseToken").With("54321", 2).AndReturn(tokenManager.UserClaims{
						Username:  "root",
						Role:      "admin",
						Firstname: "m",
						Lastname:  "m",
					}, nil))

					res := httptest.NewRecorder()
					in := `{"username":"root","password":"1234"}`
					req, _ := http.NewRequest("PUT", "/users/root", bytes.NewReader([]byte(in)))
					req.Header.Add("Authorization", "Bearer 54321")
					h.InitRouter(false).ServeHTTP(res, req)

					Expect(res.Result().StatusCode).To(Equal(403))
				})
				It("error UpdateUser", func() {
					AllowDouble(tok).To(ReceiveCallTo("ParseToken").With("54321", 2).AndReturn(tokenManager.UserClaims{
						Username:  "root",
						Role:      "regular",
						Firstname: "m",
						Lastname:  "m",
					}, nil))
					AllowDouble(db).To(ReceiveCallTo("UpdateUser").With(model.UserForUpdate{
						Username: "root",
						Password: "1234",
						Info: model.UserInfoForUpdate{
							Firstname: "",
							Lastname:  "",
							Surname:   "",
							Group:     "",
						},
					}).AndReturn(errors.New("some error")))

					res := httptest.NewRecorder()
					in := `{"username":"root","password":"1234"}`
					req, _ := http.NewRequest("PUT", "/users/root", bytes.NewReader([]byte(in)))
					req.Header.Add("Authorization", "Bearer 54321")
					h.InitRouter(false).ServeHTTP(res, req)

					Expect(res.Result().StatusCode).To(Equal(400))
				})
			})

			Context("test deleteUser", func() {
				It("regular not admin", func() {
					AllowDouble(tok).To(ReceiveCallTo("ParseToken").With("54321", 2).AndReturn(tokenManager.UserClaims{
						Username:  "root",
						Role:      "regular",
						Firstname: "m",
						Lastname:  "m",
					}, nil))
					AllowDouble(mygrpc).To(ReceiveCallTo("DeleteOrdersForUser").With("root").AndReturn(nil))
					AllowDouble(db).To(ReceiveCallTo("DeleteUser").With("root").AndReturn(nil))

					res := httptest.NewRecorder()
					req, _ := http.NewRequest("DELETE", "/users/root", nil)
					req.Header.Add("Authorization", "Bearer 54321")
					h.InitRouter(false).ServeHTTP(res, req)

					Expect(res.Result().StatusCode).To(Equal(200))
				})
				It("error admin", func() {
					AllowDouble(tok).To(ReceiveCallTo("ParseToken").With("54321", 2).AndReturn(tokenManager.UserClaims{
						Username:  "root",
						Role:      "admin",
						Firstname: "m",
						Lastname:  "m",
					}, nil))

					res := httptest.NewRecorder()
					req, _ := http.NewRequest("DELETE", "/users/root", nil)
					req.Header.Add("Authorization", "Bearer 54321")
					h.InitRouter(false).ServeHTTP(res, req)

					Expect(res.Result().StatusCode).To(Equal(403))
				})
				It("error DeleteUser", func() {
					AllowDouble(tok).To(ReceiveCallTo("ParseToken").With("54321", 2).AndReturn(tokenManager.UserClaims{
						Username:  "root",
						Role:      "regular",
						Firstname: "m",
						Lastname:  "m",
					}, nil))
					AllowDouble(mygrpc).To(ReceiveCallTo("DeleteOrdersForUser").With("root").AndReturn(nil))
					AllowDouble(db).To(ReceiveCallTo("DeleteUser").With("root").AndReturn(errors.New("some error")))

					res := httptest.NewRecorder()
					req, _ := http.NewRequest("DELETE", "/users/root", nil)
					req.Header.Add("Authorization", "Bearer 54321")
					h.InitRouter(false).ServeHTTP(res, req)

					Expect(res.Result().StatusCode).To(Equal(400))
				})
				It("error DeleteOrdersForUser", func() {
					AllowDouble(tok).To(ReceiveCallTo("ParseToken").With("54321", 2).AndReturn(tokenManager.UserClaims{
						Username:  "root",
						Role:      "regular",
						Firstname: "m",
						Lastname:  "m",
					}, nil))
					AllowDouble(mygrpc).To(ReceiveCallTo("DeleteOrdersForUser").With("root").AndReturn(errors.New("some error")))

					res := httptest.NewRecorder()
					req, _ := http.NewRequest("DELETE", "/users/root", nil)
					req.Header.Add("Authorization", "Bearer 54321")
					h.InitRouter(false).ServeHTTP(res, req)

					Expect(res.Result().StatusCode).To(Equal(400))
				})
			})

			Context("test tryToOrderTask", func() {
				It("regular", func() {
					AllowDouble(tok).To(ReceiveCallTo("ParseToken").With("54321", 2).AndReturn(tokenManager.UserClaims{
						Username:  "root",
						Role:      "regular",
						Firstname: "m",
						Lastname:  "m",
					}, nil))
					AllowDouble(mygrpc).To(ReceiveCallTo("CheckAndGetTask").With("root", 1).AndReturn(model.TaskOrderInfo{
						Id:     1,
						Answer: 2,
						Price:  500,
					}, nil))
					AllowDouble(db).To(ReceiveCallTo("TryToBuyTask").With("root", 500).AndReturn(nil))
					AllowDouble(mygrpc).To(ReceiveCallTo("BuyTaskAnswer").With("root", 1).AndReturn(nil))

					res := httptest.NewRecorder()
					req, _ := http.NewRequest("POST", "/users/root", nil)
					req.Header.Add("Authorization", "Bearer 54321")
					req.Header.Add("taskId", "1")
					h.InitRouter(false).ServeHTTP(res, req)

					Expect(res.Result().StatusCode).To(Equal(200))
					Expect(io.ReadAll(res.Body)).To(Equal([]byte(`{"answer":2}`)))
				})
			})

			Context("test updateUserBalance", func() {
				It("regular", func() {
					AllowDouble(tok).To(ReceiveCallTo("ParseToken").With("54321", 2).AndReturn(tokenManager.UserClaims{
						Username:  "root",
						Role:      "admin",
						Firstname: "m",
						Lastname:  "m",
					}, nil))
					AllowDouble(db).To(ReceiveCallTo("UpdateBalance").With("root", 2000).AndReturn(nil))

					res := httptest.NewRecorder()
					req, _ := http.NewRequest("PATCH", "/users/root", bytes.NewReader([]byte(`{"username":"root","money":2000}`)))
					req.Header.Add("Authorization", "Bearer 54321")
					h.InitRouter(false).ServeHTTP(res, req)

					Expect(res.Result().StatusCode).To(Equal(200))
				})

				It("regular not admin", func() {
					AllowDouble(tok).To(ReceiveCallTo("ParseToken").With("54321", 2).AndReturn(tokenManager.UserClaims{
						Username:  "root",
						Role:      "regular",
						Firstname: "m",
						Lastname:  "m",
					}, nil))

					res := httptest.NewRecorder()
					req, _ := http.NewRequest("PATCH", "/users/root", bytes.NewReader([]byte(`{"username":"root","money":2000}`)))
					req.Header.Add("Authorization", "Bearer 54321")
					h.InitRouter(false).ServeHTTP(res, req)

					Expect(res.Result().StatusCode).To(Equal(403))
				})
			})
		})
		Context("test Orders", func() {
			Context("test getUsersTasks", func() {
				It("regular", func() {
					AllowDouble(tok).To(ReceiveCallTo("ParseToken").With("54321", 2).AndReturn(tokenManager.UserClaims{
						Username:  "root",
						Role:      "admin",
						Firstname: "m",
						Lastname:  "m",
					}, nil))
					AllowDouble(mygrpc).To(ReceiveCallTo("GetOrdersForUser").With("root", 1).AndReturn([]model.Task{
						{Id: 1},
						{Id: 2},
					}, nil))

					res := httptest.NewRecorder()
					req, _ := http.NewRequest("GET", "/users/root/orders/purchased/1", nil)
					req.Header.Add("Authorization", "Bearer 54321")
					h.InitRouter(false).ServeHTTP(res, req)

					Expect(res.Result().StatusCode).To(Equal(200))
					Expect(io.ReadAll(res.Body)).To(Equal([]byte("page №1:\n[{\"id\":1,\"count\":0,\"heights\":null,\"price\":0,\"answer\":0},{\"id\":2,\"count\":0,\"heights\":null,\"price\":0,\"answer\":0}]")))
				})
			})

			Context("test getAllTasks", func() {
				It("regular", func() {
					AllowDouble(tok).To(ReceiveCallTo("ParseToken").With("54321", 2).AndReturn(tokenManager.UserClaims{
						Username:  "root",
						Role:      "admin",
						Firstname: "m",
						Lastname:  "m",
					}, nil))
					AllowDouble(mygrpc).To(ReceiveCallTo("GetAllTasks").With().AndReturn([]model.Task{
						{Id: 1}, {Id: 2},
					}, nil))

					res := httptest.NewRecorder()
					req, _ := http.NewRequest("GET", "/users/root/orders", nil)
					req.Header.Add("Authorization", "Bearer 54321")
					h.InitRouter(false).ServeHTTP(res, req)

					Expect(res.Result().StatusCode).To(Equal(200))
					Expect(io.ReadAll(res.Body)).To(Equal([]byte(`[{"id":1,"count":0,"heights":null,"price":0,"answer":0},{"id":2,"count":0,"heights":null,"price":0,"answer":0}]`)))
				})
			})

			Context("test updateTask", func() {
				It("regular", func() {
					AllowDouble(tok).To(ReceiveCallTo("ParseToken").With("54321", 2).AndReturn(tokenManager.UserClaims{
						Username:  "root",
						Role:      "admin",
						Firstname: "m",
						Lastname:  "m",
					}, nil))
					AllowDouble(mygrpc).To(ReceiveCallTo("UpdatePriceOfTask").With(1, 1500).AndReturn(nil))

					res := httptest.NewRecorder()
					req, _ := http.NewRequest("PUT", "/users/root/orders", bytes.NewReader([]byte(`{"taskId":1,"balance":1500}`)))
					req.Header.Add("Authorization", "Bearer 54321")
					h.InitRouter(false).ServeHTTP(res, req)

					Expect(res.Result().StatusCode).To(Equal(200))
				})
			})

			Context("test createTask", func() {
				It("regular", func() {
					AllowDouble(tok).To(ReceiveCallTo("ParseToken").With("54321", 2).AndReturn(tokenManager.UserClaims{
						Username:  "root",
						Role:      "admin",
						Firstname: "m",
						Lastname:  "m",
					}, nil))
					AllowDouble(mygrpc).To(ReceiveCallTo("CreateNewTask").With(model.Task{Id: 1, Count: 2, Heights: []int64{1, 2}, Price: 400, Answer: 2}).AndReturn(nil))

					res := httptest.NewRecorder()
					req, _ := http.NewRequest("POST", "/users/root/orders", bytes.NewReader([]byte(`{"id":1,"count":2,"heights":[1,2],"price":400,"answer":2}`)))
					req.Header.Add("Authorization", "Bearer 54321")
					h.InitRouter(false).ServeHTTP(res, req)

					Expect(res.Result().StatusCode).To(Equal(200))
				})
			})

			Context("test createTask", func() {
				It("regular", func() {
					AllowDouble(tok).To(ReceiveCallTo("ParseToken").With("54321", 2).AndReturn(tokenManager.UserClaims{
						Username:  "root",
						Role:      "admin",
						Firstname: "m",
						Lastname:  "m",
					}, nil))
					AllowDouble(mygrpc).To(ReceiveCallTo("DeleteTask").With(1).AndReturn(nil))

					res := httptest.NewRecorder()
					req, _ := http.NewRequest("DELETE", "/users/root/orders/1", nil)
					req.Header.Add("Authorization", "Bearer 54321")
					h.InitRouter(false).ServeHTTP(res, req)
				})
			})
		})
	})
})
