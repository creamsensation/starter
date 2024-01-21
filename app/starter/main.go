package main

import (
	"github.com/creamsensation/cp"
	"github.com/creamsensation/cp/env"
	
	"starter/controller"
	"starter/ui"
)

func main() {
	env.EnableDevelopmentMode()
	app := cp.New("./app/starter/config")
	app.Ui().Layout("main", ui.Layout)
	app.Controllers(
		&controller.WelcomeController{},
	)
	app.Serve()
}
