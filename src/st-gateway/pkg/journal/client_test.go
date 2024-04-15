package journal_test

import (
	"github.com/menezmethod/st-server/src/st-gateway/configs"
	"github.com/menezmethod/st-server/src/st-gateway/pkg/journal"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"log"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

var _ = Describe("JournalClient: Journal Service", func() {
	BeforeEach(func() {
		lis = bufconn.Listen(bufSize)
		s := grpc.NewServer()

		go func() {
			if err := s.Serve(lis); err != nil {
				log.Fatalf("Server exited with error: %v", err)
			}
		}()
	})

	Describe("Function InitRecordServiceClient()", func() {
		Context("Initializing the journal service client", func() {
			It("Should successfully create a service client", func() {
				c := &configs.Config{
					JournalSvcUrl: "bufnet",
				}

				client := journal.InitJournalServiceClient(c)
				Expect(client).NotTo(BeNil())
			})
		})
	})

	AfterEach(func() {
		err := lis.Close()
		if err != nil {
			return
		}
	})
})
