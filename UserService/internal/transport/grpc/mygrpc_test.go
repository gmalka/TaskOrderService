package mygrpc_test

import (
	"context"
	"net"
	mygrpc "userService/internal/transport/grpc"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

var _ = Describe("Mygrpc", func() {
	var g mygrpc.RemoteOrderClient

	BeforeEach(func() {
		lis := bufconn.Listen(1024 * 1024)
		dialer := func(context.Context, string) (net.Conn, error) {
			return lis.Dial()
		}
		conn, _ := grpc.DialContext(context.Background(), "", grpc.WithContextDialer(dialer), grpc.WithInsecure())
		mygrpc.NewGrpcClient(conn)
	})

	Context("", func() {
		It("", func() {
		
			
		})
	})
})
