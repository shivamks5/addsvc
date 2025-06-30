package main

import (
	"log"
	"net/http"

	addsvc "github.com/shivamks5/51-addsvc"
)

func main() {
	svc := addsvc.NewService()
	ep := addsvc.MakeEndpoints(svc)
	handler := addsvc.MakeHTTPHandler(ep)
	log.Println("Server : http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
