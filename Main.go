package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/julienschmidt/httprouter"
)

func main() {
	srv := &http.Server{
		Addr:    ":10101",
		Handler: newRouter(),
	}
	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		signal.Notify(sigint, syscall.SIGTERM)
		<-sigint
		log.Println("service interrupt received")
		log.Println("http server shutting down")
		time.Sleep(5 * time.Second)

		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			log.Printf("http server shutdown error: %v", err)
		}

		log.Println("shutdown complete")

		close(idleConnsClosed)

	}()

	log.Printf("Starting server on port 10101")
	if err := srv.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("fatal http server failed to start: %v", err)
		}
	}

	<-idleConnsClosed
	log.Println("Service Stop")

}

func newRouter() *httprouter.Router {
	mux := httprouter.New()
	// ytApiKey := os.Getenv("YOUTUB_API_KEY")

	// if ytApiKey == "" {
	// 	log.Fatal("youtube API key not provided")
	// }

	mux.GET("/channel/stats", GetChannelStats())
	mux.GET("/ping", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("pong"))
		if err != nil {
			panic(err)
		}
	})
	return mux
}
