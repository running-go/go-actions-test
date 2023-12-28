package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
)

func main() {
	srv := &http.Server{
		Addr:    ":10101",
		Handler: newRouter(),
	}
	//idleConnsClosed := make(chan struct{})
	//go func() {
	//	sigint := make(chan os.Signal, 1)
	//	signal.Notify(sigint, os.Interrupt)
	//	signal.Notify(sigint, syscall.SIGTERM)
	//	<-sigint
	//	log.Println("service interrupt received")
	//	log.Println("http server shutting down")
	//	time.Sleep(5 * time.Second)
	//
	//}()

	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}

}

func newRouter() *httprouter.Router {
	mux := httprouter.New()
	ytApiKey := os.Getenv("YOUTUB_API_KEY")

	if ytApiKey == "" {
		log.Fatal("youtube API key not provided")
	}

	mux.GET("/channel/stats", GetChannelStats())
	return mux
}
