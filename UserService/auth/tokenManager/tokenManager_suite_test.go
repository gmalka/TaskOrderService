package tokenManager_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestTokenManager(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TokenManager Suite")
}
