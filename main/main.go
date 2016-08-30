package main

import (
	"fmt"

	"github.com/cagnosolutions/mg"
)

type User struct {
	Name   string
	Age    int
	Active bool
}

func main() {

	user := User{
		Name:   "Greg Pechiro",
		Age:    30,
		Active: true,
	}
	body, err := mg.Body("email.tmpl", map[string]interface{}{"user": user})

	if err != nil {
		panic(err)
	}
	//r, err := mg.Send("gregpechiro@gmail.com", "Totally not a virus or spam <info@test.com>", "TEST EMAIL", body, "gregpechiro@yahoo.com", "scottiecagno@gmaio.com", "cagnosolutions@gmail.com")
	email := mg.Email{
		To:      []string{"gregpechiro@gmail.com"},
		From:    "Not a virus or spam <info@test.com>",
		Subject: "Totally not a virus or spam... I promise",
		HTML:    body,
		CC:      []string{"gregpechiro@yahoo.com", "cagnosolutions@gmail.com"},
		BCC:     []string{"scottiecagno@gmail.com"},
	}
	r, err := mg.SendEmail(email)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", r)
}
