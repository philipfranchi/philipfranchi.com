package main

import (
	"log"
	"net/http"
)

func main() {
	config := CreateConfigFromEnv()
	blog := CreateBlog(config)
	handler := CreateAPIHandler(blog)
	router := CreateAPIRouter(handler)
	fullAddress := config.ApplicationAddress + ":" + config.ApplicationPort
	log.Println("Application is running on " + fullAddress)
	http.ListenAndServe(fullAddress, router)
}
