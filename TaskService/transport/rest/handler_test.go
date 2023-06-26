package rest_test

import (
	"bytes"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"taskServer/model"
	"taskServer/transport/rest"

	. "github.com/gcapizzi/moka"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var (
	logger rest.Log
)

var _ = BeforeSuite(func() {
	loggerErr := log.New(os.Stderr, "ERROR:\t ", log.Lshortfile|log.Ltime)
	loggerInfo := log.New(os.Stdout, "INFO:\t ", log.Lshortfile|log.Ltime)
	logger = rest.Log{loggerErr, loggerInfo}
})

var _ = Describe("Handler", func() {
	var (
		controller ControllerDouble
		h          rest.Handler
		//ts         *httptest.Server
	)

	BeforeEach(func() {
		controller = NewDbDouble()
		h = rest.NewHandler(controller, logger)
	})

	Context("Handler testing ", func() {
		Context("testing getTasks ", func() {
			It("regular", func() {
				AllowDouble(controller).To(ReceiveCallTo("GetAllTasks").With().AndReturn([]model.Task{
					{Id: 1, Count: 2, Heights: []int64{1, 2}, Price: 500, Answer: 2},
				}, nil))

				req, err := http.NewRequest("GET", "/", nil)
				res := httptest.NewRecorder()
				h.InitRouter().ServeHTTP(res, req)

				js := `[{"id":1,"quantity":2,"heights":[1,2],"price":500,"answer":2}]`

				Expect(err).Should(Succeed())
				Expect(res.Result().StatusCode).To(Equal(200))
				Expect(res.Header().Get("Content-Type")).To(Equal("application/json"))
				Expect(ioutil.ReadAll(res.Body)).To(Equal([]byte(js)))
			})

			It("error", func() {
				AllowDouble(controller).To(ReceiveCallTo("GetAllTasks").With().AndReturn([]model.Task{
					{Id: 1, Count: 2, Heights: []int64{1, 2}, Price: 500, Answer: 2},
				}, errors.New("some error")))

				req, err := http.NewRequest("GET", "/", nil)
				res := httptest.NewRecorder()
				h.InitRouter().ServeHTTP(res, req)

				Expect(err).Should(Succeed())
				Expect(res.Result().StatusCode).To(Equal(400))
				want := []byte(`message: some error`)
				want = append(want, '\n')
				Expect(ioutil.ReadAll(res.Body)).To(Equal(want))
			})
		})

		Context("testing getTask", func() {
			It("regular", func() {
				AllowDouble(controller).To(ReceiveCallTo("CheckAndGetTask").With("", 2).AndReturn(model.Task{Id: 1, Count: 2, Heights: []int64{1, 2}, Price: 500, Answer: 2}, nil))

				res := httptest.NewRecorder()
				req, err := http.NewRequest("GET", "/2", nil)

				h.InitRouter().ServeHTTP(res, req)

				js := `{"id":1,"quantity":2,"heights":[1,2],"price":500,"answer":2}`

				Expect(err).Should(Succeed())
				Expect(res.Result().StatusCode).To(Equal(200))
				Expect(res.Header().Get("Content-Type")).To(Equal("application/json"))
				Expect(ioutil.ReadAll(res.Body)).To(Equal([]byte(js)))
			})


			It("error", func() {
				AllowDouble(controller).To(ReceiveCallTo("CheckAndGetTask").With("", 2).AndReturn(model.Task{}, errors.New("some error")))

				req, err := http.NewRequest("GET", "/2", nil)
				res := httptest.NewRecorder()
				h.InitRouter().ServeHTTP(res, req)

				Expect(err).Should(Succeed())
				Expect(res.Result().StatusCode).To(Equal(400))
				want := []byte(`message: some error`)
				want = append(want, '\n')
				Expect(ioutil.ReadAll(res.Body)).To(Equal(want))
			})
		})

		Context("testing createTask", func() {
			It("regular", func() {
				AllowDouble(controller).To(ReceiveCallTo("CreateTask").With(model.Task{Id: 0, Count: 2, Heights: []int64{1, 2}, Price: 500, Answer: 2}).AndReturn(nil))

				res := httptest.NewRecorder()
				js := `{"id":0,"quantity":2,"heights":[1,2],"price":500,"answer":2}`
				req, err := http.NewRequest("POST", "/", bytes.NewReader([]byte(js)))

				h.InitRouter().ServeHTTP(res, req)

				Expect(err).Should(Succeed())
				Expect(res.Result().StatusCode).To(Equal(200))
				Expect(res.Header().Get("Content-Type")).To(Equal("application/json"))
				Expect(ioutil.ReadAll(res.Body)).To(Equal([]byte(`{"message":"success create"}`)))
			})

			It("error", func() {
				AllowDouble(controller).To(ReceiveCallTo("CreateTask").With(model.Task{Id: 0, Count: 2, Heights: []int64{1, 2}, Price: 500, Answer: 2}).AndReturn(errors.New("some error")))

				res := httptest.NewRecorder()
				js := `{"id":0,"quantity":2,"heights":[1,2],"price":500,"answer":2}`
				req, err := http.NewRequest("POST", "/", bytes.NewReader([]byte(js)))

				h.InitRouter().ServeHTTP(res, req)

				Expect(err).Should(Succeed())
				Expect(res.Result().StatusCode).To(Equal(400))
				want := []byte(`message: some error`)
				want = append(want, '\n')
				Expect(ioutil.ReadAll(res.Body)).To(Equal(want))
			})
		})

		Context("testing changeTaskPrice", func() {
			It("regular", func() {
				AllowDouble(controller).To(ReceiveCallTo("ChangeTaskPrice").With(1, 2).AndReturn(nil))

				res := httptest.NewRecorder()
				js := `{"orderId":1,"price":2}`
				req, err := http.NewRequest("PATCH", "/", bytes.NewReader([]byte(js)))

				h.InitRouter().ServeHTTP(res, req)

				Expect(err).Should(Succeed())
				Expect(res.Result().StatusCode).To(Equal(200))
				Expect(res.Header().Get("Content-Type")).To(Equal("application/json"))
				Expect(ioutil.ReadAll(res.Body)).To(Equal([]byte(`{"message":"success update"}`)))
			})

			It("error", func() {
				AllowDouble(controller).To(ReceiveCallTo("ChangeTaskPrice").With(1, 2).AndReturn(errors.New("some error")))

				res := httptest.NewRecorder()
				js := `{"orderId":1,"price":2}`
				req, err := http.NewRequest("PATCH", "/", bytes.NewReader([]byte(js)))

				h.InitRouter().ServeHTTP(res, req)

				Expect(err).Should(Succeed())
				Expect(res.Result().StatusCode).To(Equal(400))
				want := []byte(`message: some error`)
				want = append(want, '\n')
				Expect(ioutil.ReadAll(res.Body)).To(Equal(want))
			})
		})

		Context("testing deleteTaskForUser", func() {
			It("regular", func() {
				AllowDouble(controller).To(ReceiveCallTo("DeleteTask").With(1).AndReturn(nil))

				res := httptest.NewRecorder()
				req, err := http.NewRequest("DELETE", "/1", nil)

				h.InitRouter().ServeHTTP(res, req)

				Expect(err).Should(Succeed())
				Expect(res.Result().StatusCode).To(Equal(200))
				Expect(res.Header().Get("Content-Type")).To(Equal("application/json"))
				Expect(ioutil.ReadAll(res.Body)).To(Equal([]byte(`{"message":"success delete"}`)))
			})

			It("error", func() {
				AllowDouble(controller).To(ReceiveCallTo("DeleteTask").With(1).AndReturn(errors.New("some error")))

				res := httptest.NewRecorder()
				req, err := http.NewRequest("DELETE", "/1", nil)

				h.InitRouter().ServeHTTP(res, req)

				Expect(err).Should(Succeed())
				Expect(res.Result().StatusCode).To(Equal(400))
				want := []byte(`message: some error`)
				want = append(want, '\n')
				Expect(ioutil.ReadAll(res.Body)).To(Equal(want))
			})
		})

		Context("testing deleteTask", func() {
			It("regular", func() {
				AllowDouble(controller).To(ReceiveCallTo("DeleteAllTasksOfUser").With("root").AndReturn(nil))

				res := httptest.NewRecorder()
				req, err := http.NewRequest("DELETE", "/users/root", nil)

				h.InitRouter().ServeHTTP(res, req)

				Expect(err).Should(Succeed())
				Expect(res.Result().StatusCode).To(Equal(200))
				Expect(res.Header().Get("Content-Type")).To(Equal("application/json"))
				Expect(ioutil.ReadAll(res.Body)).To(Equal([]byte(`{"message":"success delete"}`)))
			})

			It("regular", func() {
				AllowDouble(controller).To(ReceiveCallTo("DeleteAllTasksOfUser").With("root").AndReturn(errors.New("some error")))

				res := httptest.NewRecorder()
				req, err := http.NewRequest("DELETE", "/users/root", nil)

				h.InitRouter().ServeHTTP(res, req)

				Expect(err).Should(Succeed())
				Expect(res.Result().StatusCode).To(Equal(400))
				want := []byte(`message: some error`)
				want = append(want, '\n')
				Expect(ioutil.ReadAll(res.Body)).To(Equal(want))
			})
		})
	})
})

// Id      int     `db:"id" json:"id,omitempty"`
// Count   int     `db:"quantity" json:"quantity"`
// Heights []int64 `db:"heights" json:"heights"`
// Price   int     `db:"price" json:"price"`
// Answer  int     `db:"answer" json:"answer"`
