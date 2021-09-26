package main

import (
	"log"

	"github.com/delihiros/shockv/pkg/jsonutil"

	"github.com/delihiros/shockv/pkg/client"
)

func main() {
	c := client.New("http://localhost", 8080)
	database := "hello"
	_, err := c.NewDB(database, true)
	if err != nil {
		log.Println(err)
	}
	_, err = c.Set(database, "1", "abc", 0)
	if err != nil {
		panic(err)
	}
	_, err = c.Set(database, "2", "def", 0)
	if err != nil {
		panic(err)
	}
	_, err = c.Set(database, "3", "xyz", 0)
	if err != nil {
		panic(err)
	}
	list, err := c.List(database)
	jsonutil.PrintJSON(list, false)

	r, err := c.Delete(database, "3")
	if err != nil {
		panic(err)
	}
	jsonutil.PrintJSON(r, false)
	list, err = c.List(database)
	if err != nil {
		panic(err)
	}
	jsonutil.PrintJSON(list, false)
}
