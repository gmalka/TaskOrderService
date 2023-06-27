package postgresservice_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestPostgresService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PostgresService Suite")
}