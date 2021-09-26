package main

import (
	"log"

	"github.com/delihiros/shockv/pkg/client"
)

func main() {
	c := client.New("http://localhost:8080")
	database := "hello"
	err := c.NewDB(database, true)
	if err != nil {
		log.Println(err)
	}
	err = c.Set(database, "1", "abc")
	if err != nil {
		log.Println(err)
	}
	err = c.Set(database, "2", "def")
	if err != nil {
		log.Println(err)
	}
	err = c.Set(database, "3", "xyz")
	if err != nil {
		log.Println(err)
	}
	list, err := c.List(database)
	if err != nil {
		log.Println(err)
	}
	log.Println(list)
	for k, v := range list {
		log.Println(k, v)
	}

	err = c.Delete(database, "3")
	if err != nil {
		log.Println(err)
	}
	list, err = c.List(database)
	if err != nil {
		log.Println(err)
	}
	for k, v := range list {
		log.Println(k, v)
	}
}
