package main

import (
	"github.com/creamsensation/cp"
	"github.com/creamsensation/quirk/migrator"
)

func init() {
	manager.Add().
		Up(
			func(c migrator.Control) {
				cp.CreateUsersTable(c.DB())
				cp.CreateUserManager(c.DB().DB, 0, "").Create(
					cp.User{
						Active:   true,
						Roles:    []string{"owner"},
						Email:    "admin@cream.puff",
						Password: "asd123",
					},
				)
			},
		).
		Down(
			func(c migrator.Control) {
				cp.DropUsersTable(c.DB())
			},
		)
}
