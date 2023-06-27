package tokenManager_test

import (
	"userService/auth/tokenManager"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("TokenManager", func() {
	var tm tokenManager.TokenManager
	BeforeEach(func() {
		tm = tokenManager.NewAuthService("4123123", "4123412r12e12rf12f")
	})

	Context("tokenManager", func() {
		It("AccessToken", func() {
			wanted := tokenManager.UserInfo{
				Username:  "root",
				Role:      "regular",
				Firstname: "g",
				Lastname:  "m",
			}
			s, _ := tm.CreateToken(wanted, tokenManager.REFRESH_TOKEN_TTL, tokenManager.AccessToken)

			claims, err := tm.ParseToken(s, tokenManager.AccessToken)
			Expect(err).Should(Succeed())

			getted := tokenManager.UserInfo{
				Username:  claims.Username,
				Role:      claims.Role,
				Firstname: claims.Firstname,
				Lastname:  claims.Lastname,
			}
			Expect(wanted).To(Equal(getted))
		})
		It("error", func() {
			wanted := tokenManager.UserInfo{
				Username:  "root",
				Role:      "regular",
				Firstname: "g",
				Lastname:  "m",
			}
			s, _ := tm.CreateToken(wanted, tokenManager.REFRESH_TOKEN_TTL, tokenManager.AccessToken)

			_, err := tm.ParseToken(s, tokenManager.RefreshToken)
			Expect(err).ShouldNot(Succeed())
		})
	})
})
