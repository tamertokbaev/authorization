package app

import (
	"diploma/authorization/internal/config"
	httpDelivery "diploma/authorization/internal/delivery/handlers/http"
	"diploma/authorization/internal/server"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg *config.Config) {

	healthCheckFn := func() error {
		return nil
	}

	handlerDelivery := httpDelivery.NewHandlerDelivery(
		healthCheckFn,
		"authorization",
	)

	// HTTP Server
	srv, err := server.NewServer(cfg, handlerDelivery)
	if err != nil {
		fmt.Printf(err.Error())
	}

	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("🔥 Server stopped due error")
		} else {
			fmt.Printf("✅ Server shutdown successfully")
		}
	}()

	fmt.Sprintf("🚀 Starting server at http://0.0.0.0:%s", cfg.Service.Port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit
}
