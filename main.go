package main

import (
	tasks "app/cmd/tasks"
	app "app/src/ui/console"
	win "app/src/ui/desktop"

	core "github.com/pastusiak/golang/tpl/core"
)

func main() {
	application := core.NewApp(false, core.ModeDesktop, true, true)

	application.RegisterCmd(&app.Example{})
	application.RegisterCmd(&tasks.Example{})

	application.RegisterWindow(&win.Example{})

	application.Run()
}
