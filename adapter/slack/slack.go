package slack

import (
	"fmt"
	"os"
	"strings"

	"github.com/danryan/env"
	"github.com/davecgh/go-spew/spew"
	"github.com/mattouille/hal"
	"github.com/nlopes/slack"
)

func init() {
	hal.RegisterAdapter("slack", New)
}

type adapter struct {
	hal.BasicAdapter
	token          string
	team           string
	mode           string
	channels       []string
	channelMode    string
	botname        string
	responseMethod string
	linkNames      int
	rtm            *slack.RTM
}

type config struct {
	Token          string `env:"key=HAL_SLACK_TOKEN required"`
	Team           string `env:"key=HAL_SLACK_TEAM"`
	Channels       string `env:"key=HAL_SLACK_CHANNELS"`
	Mode           string `env:"key=HAL_SLACK_MODE"`
	Botname        string `env:"key=HAL_SLACK_BOTNAME default=hal"`
	ResponseMethod string `env:"key=HAL_SLACK_RESPONSE_METHOD default=rtm"`
	ChannelMode    string `env:"key=HAL_SLACK_CHANNEL_MODE default=none"`
}

// New returns an initialized adapter
func New(r *hal.Robot) (hal.Adapter, error) {
	c := &config{}
	env.MustProcess(c)
	channels := strings.Split(c.Channels, ",")
	a := &adapter{
		token:          c.Token,
		team:           c.Team,
		channels:       channels,
		channelMode:    c.ChannelMode,
		mode:           c.Mode,
		botname:        c.Botname,
		responseMethod: c.ResponseMethod,
	}
	spew.Dump(c)
	hal.Logger.Debugf("%v", os.Getenv("HAL_SLACK_CHANNEL_MODE"))
	hal.Logger.Debugf("channel mode: %v", a.channelMode)

	a.SetRobot(r)
	return a, nil
}

// Send sends a regular response
func (a *adapter) Send(res *hal.Response, strings ...string) error {
	for _, str := range strings {
		out := a.rtm.NewOutgoingMessage(str, res.Message.Room)
		a.rtm.SendMessage(out)
	}

	return nil
}

// Reply sends a direct response
func (a *adapter) Reply(res *hal.Response, strings ...string) error {
	newStrings := make([]string, 0)
	for _, str := range strings {
		newStrings = append(newStrings, fmt.Sprintf("<@%s> %s", res.UserID(), str))
	}

	return a.Send(res, newStrings...)
}

// Emote is not implemented.
func (a *adapter) Emote(res *hal.Response, strings ...string) error {
	return nil
}

// Topic sets the topic
func (a *adapter) Topic(res *hal.Response, strings ...string) error {
	for range strings {
	}
	return nil
}

// Play is not implemented.
func (a *adapter) Play(res *hal.Response, strings ...string) error {
	return nil
}

// Receive forwards a message to the robot
func (a *adapter) Receive(msg *hal.Message) error {
	hal.Logger.Debug("slack - adapter received message")

	if len(a.channels) > 0 && a.channelMode != "none" {
		if a.channelMode == "blacklist" {
			if !a.inChannels(msg.Room) {
				hal.Logger.Debugf("slack - %s not in blacklist", msg.Room)
				hal.Logger.Debug("slack - adapter sent message to robot")
				return a.Robot.Receive(msg)
			}
			hal.Logger.Debug("slack - message ignored due to blacklist")
			return nil
		}

		if a.inChannels(msg.Room) {
			hal.Logger.Debugf("slack - %s in whitelist", msg.Room)
			hal.Logger.Debug("slack - adapter sent message to robot")
			return a.Robot.Receive(msg)
		}
		hal.Logger.Debug("slack - message ignored due to whitelist")
		return nil
	}

	hal.Logger.Debug("slack - adapter sent message to robot")
	return a.Robot.Receive(msg)
}

// Run starts the adapter
func (a *adapter) Run() error {
	// set up a connection to RTM API
	hal.Logger.Debug("slack - starting RTM connection")
	go a.startConnection()
	hal.Logger.Debug("slack - started RTM connection")

	hal.Logger.Debugf("slack - channelmode=%v channels=%v", a.channelMode, a.channels)
	return nil
}

// Stop shuts down the adapter
func (a *adapter) Stop() error {
	return nil
}

func (a *adapter) inChannels(room string) bool {
	for _, r := range a.channels {
		if r == room {
			return true
		}
	}

	return false
}
