package controller

import (
	"github.com/creamsensation/cp"
	. "github.com/creamsensation/gox"
)

type WelcomeController struct {
}

func (r *WelcomeController) Name() string {
	return "welcome"
}

func (r *WelcomeController) Routes() cp.Routes {
	return cp.Routes{
		cp.Route(cp.Get(""), cp.Name("main"), cp.Handler(r.Main)),
	}
}

func (r *WelcomeController) Main(c cp.Control) cp.Result {
	c.Page().Set().Title("Welcome!")
	return c.Response().Render(
		Div(
			Class("w-screen h-screen grid place-items-center text-2xl font-bold"),
			Text("Welcome to CP family! \t&#128513;"),
		),
	)
}
