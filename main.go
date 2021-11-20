package main

import (
	"fmt"

	"github.com/pranavshukla1011/golang_webservice_sample/models"
)

func main(){
	u := models.User{
		ID: 1,
		FirstName: "Pranav",
		LastName: "Shukla",
	}

	fmt.Println(u)
}