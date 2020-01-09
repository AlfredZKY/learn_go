package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"go.uber.org/fx"
)

// NewLogger Logger构造函数
func NewLogger() *log.Logger {
	logger := log.New(os.Stdout, "httpLog", 0)
	logger.Print("Executing NewLogger.")
	return logger
}

// NewHandler http.Hamdler构造函数
func NewHandler(logger *log.Logger) (http.Handler, error) {
	logger.Print("Executing NewHandler")
	return http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
		logger.Print("Got a rquest.")
	}), nil
}

// NewMux http.ServerMux 构造函数
func NewMux(lc fx.Lifecycle, logger *log.Logger) *http.ServeMux {
	logger.Print("Executing NewMux.")
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    "8080",
		Handler: mux,
	}

	// 自定义生命周期过程对应的启动和关闭的行为
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Print("Starting HTTP server.")
			go server.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Print("Stopping HTTP server.")
			return server.Shutdown(ctx)
		},
	})
	return mux
}

// Register 注册 http.Handler
func Register(mux *http.ServeMux, h http.Handler) {
	mux.Handle("/", h)
}

func main() {
	app := fx.New(
		fx.Provide(
			NewLogger,
			NewHandler,
			NewMux,
		),
		// 通过 invoke 来完成Logger、Handler、ServerMux的创建
		fx.Invoke(Register),
	)
	startCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// 手动调用start
	if err := app.Start(startCtx); err != nil {
		log.Fatal(err)
	}

	// 具体操作
	http.Get("http://localhost:8080")

	stopCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := app.Stop(stopCtx); err != nil {
		log.Fatal(err)
	}
}
