package repository

import (
	// "context"
	// "net/http/httptest"
	"fmt"
	"log"
	"net/http/httptest"
	"os"
	"os/signal"
	"syscall"
	"testing"

	"62teknologi-senior-backend-test-muhammad_risky_nugraha/server"
	"github.com/stretchr/testify/assert"
)

// func (s *server.Server) InitTest() error {
// 	if err := s.MapHandlers(s.app); err != nil {
// 		// log.Panic(err)
// 	}

// 	return nil
// }

// var (
// 	app := fiber.App()
// )

func TestFunction(t *testing.T) {
	s := server.NewServer()
		// s.Start()
		if err := s.MapHandlers(s.App); err != nil {
			log.Panic(err)
		}
		server_port := "9101"
		fmt.Printf("Server running on port: %s\n", server_port)
			if err := s.App.Listen(fmt.Sprintf(":%s", server_port)); err != nil {
				log.Panic(err)
			}

		// time.Sleep(3.0)

	// se
	// t.Log(app)
	
	// businessRepo.GetAllData(ctx)

	req := httptest.NewRequest("GET", "/business/search", nil)
	
	// NewServer
	// Perform the request plain with the app,
	// the second argument is a request latency
	// (set to -1 for no latency)
	resp, _ := s.App.Test(req, 10)

	t.Log(resp.Body, "resp")
	fmt.Println(resp.Body)
	assert.Equal(t, 2, 1, "wrong int")
	assert.Equal(t, 200, resp.StatusCode, "Get API")

		// gracefull shutdown
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	
		<-quit
		fmt.Println("shutting down...")
		_ = s.App.Shutdown()
	// Verify, if the status code is as expected
	// assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
}

func Test2Function(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// mockUsecase := new(businessUsecase)
	})
}