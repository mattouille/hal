package handler

import "github.com/mattouille/hal"

// Echo duplicates whatever the user tells it.
var Echo = &hal.Handler{
	Method:  hal.RESPOND,
	Usage:   "@<bot> echo <something>",
	Pattern: `echo (.+)`,
	Run: func(res *hal.Response) error {
		return res.Send(res.Match[1])
	},
}
