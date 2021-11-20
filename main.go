package main

import (
	"net/http"

	"github.com/pranavshukla1011/golang_webservice_sample/controllers"
)

func main(){
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}