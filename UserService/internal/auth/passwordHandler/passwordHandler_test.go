package passwordHandler_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"golang.org/x/crypto/bcrypt"

	"userService/internal/auth/passwordHandler"
)

var _ = Describe("PasswordHandler", func() {
	var passworder passwordHandler.PasswordHandler
	BeforeEach(func() {
		passworder = passwordHandler.NewPasswordManager()
	})
	Context("passwordHandler", func() {
		Context("HashPassword", func() {
			It("regular", func() {
				password := "1234"
				b, _ := passworder.HashPassword(password)
				Expect(bcrypt.CompareHashAndPassword([]byte(b), []byte(password))).Should(Succeed())
			})
			It("wrong password", func() {
				password := "123456"
				b, _ := passworder.HashPassword(password)
				Expect(bcrypt.CompareHashAndPassword([]byte(b), []byte("12345"))).ShouldNot(Succeed())
			})
		})

		Context("CheckPassword", func() {
			It("regular", func() {
				password := "123456"
				b, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
				Expect(passworder.CheckPassword(password, string(b))).Should(Succeed())
			})
			It("wrong password", func() {
				password := "1234"
				b, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
				Expect(passworder.CheckPassword("12345", string(b))).ShouldNot(Succeed())
			})
		})
	})
})
