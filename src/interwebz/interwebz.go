package interwebz

import (
	"github.com/go-martini/martini"
	"time"
)

func Clockwise() {
	m := martini.Classic()
	m.Get("/", func() string {
		return time.Now().String()
	})
	m.Run()
}
