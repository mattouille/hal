package handler

import "github.com/mattouille/hal"

// Commands outputs a list of commands the bot accepts
var Commands = &hal.Handler{
	Method:  hal.RESPOND,
	Usage:   "help",
	Pattern: `help`,
	Run: func(res *hal.Response) error {
		return res.Reply("Multiline\nhelp\nmessage")
	},
}
