package main

import (
	"fmt"
)

type Human struct {
	name string
	surname string
}

type Action struct {
	Human
	login string
	password string
}
func (h *Human) WhoIs() {
	fmt.Printf("My name is %s %s\n", h.name, h.surname)
}

func(a *Action) Auth() {
	fmt.Printf("User %s %s has login: %s, and password: %s",
	 a.name, a.surname, a.login, a.password)
}

func main() {
	a := Action{
		Human{"Ivan","Ivanov",},
		"sukh","qwerty",
	}

	a.WhoIs()
	a.Auth()
}