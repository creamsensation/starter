package main

import (
	"log"
	
	"github.com/creamsensation/cp/env"
	"github.com/creamsensation/quirk"
	"github.com/creamsensation/quirk/migrator"
)

var manager = new(migrator.Manager)

func main() {
	env.EnableDevelopmentMode()
	db, err := quirk.Connect(
		quirk.WithPostgres(),
		quirk.WithHost("localhost"),
		quirk.WithPort(5432),
		quirk.WithDbname("starter"),
		quirk.WithUser("starter"),
		quirk.WithPassword("starter"),
		quirk.WithSslDisable(),
	)
	if err != nil {
		log.Fatalln(err)
	}
	m := migrator.New(
		"migrations",
		map[string]*quirk.DB{
			"main": db,
		},
		manager.GetAll(),
	)
	m.Run()
}
