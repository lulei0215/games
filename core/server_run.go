package core

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
	Shutdown(context.Context) error
}

// initServer start
func initServer(address string, router *gin.Engine, readTimeout, writeTimeout time.Duration) {
	srv := &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	// goroutine start
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("listen: %s\n", err)
			zap.L().Error("server start fail", zap.Error(err))
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	// kill  syscall.SIGTERM
	// kill -2  syscall.SIGINT
	// kill -9  syscall.SIGKILL
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	zap.L().Info("close WEB...")

	// 5 mi
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("WEB close fail", zap.Error(err))
	}

	zap.L().Info("WEB closed")
}
