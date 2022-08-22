package main

import (
	"fmt"
	"net/http"
	"userInfoService/routers"
)

func main() {

	r := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", 8080),
		Handler:        r,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
