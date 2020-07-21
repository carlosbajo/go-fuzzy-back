package main

import (
	"github.com/go-martini/martini"
)

func main() {
	m := martini.Classic()
	m.Post("/signin", Signin)
	m.Post("/welcome", Welcome)
	m.Post("/refresh", Refresh)
	m.Get("/", func() (s string) {
		s = "Main page"
		return
	})
	m.Run()
}
