package handler

import "github.com/mattouille/hal"

// TableFlip is an example of a Handler
var TableFlip = &hal.Handler{
	Method:  hal.HEAR,
	Usage:   "Say tableflip",
	Pattern: `tableflip`,
	Run: func(res *hal.Response) error {
		return res.Send(`(╯°□°）╯︵ ┻━┻`)
	},
}
