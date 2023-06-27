package passwordManager_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestPasswordHandler(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PasswordHandler Suite")
}
