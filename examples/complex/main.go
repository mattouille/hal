package main

import (
	"os"

	"github.com/mattouille/hal"
	_ "github.com/mattouille/hal/adapter/slack"
	"github.com/mattouille/hal/handler"
	_ "github.com/mattouille/hal/store/memory"
)

func run() int {
	robot, err := hal.NewRobot()
	if err != nil {
		hal.Logger.Fatal(err)
	}

	robot.Handle(
		handler.Ping,
		handler.Echo,
		handler.TableFlip,
		handler.Commands,
	)

	if err := robot.Run(); err != nil {
		hal.Logger.Fatal(err)
	}
	return 0
}

func main() {
	os.Exit(run())
}
