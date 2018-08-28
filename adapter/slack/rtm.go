package slack

import (
	"github.com/mattouille/hal"
	"github.com/nlopes/slack"
)

func (a *adapter) startConnection() {
	api := slack.New(a.token)

	users, err := api.GetUsers()
	if err != nil {
		hal.Logger.Error(err)
	}

	for _, user := range users {
		// Initialize a newUser object in case we need it.
		newUser := hal.User{
			ID:   user.ID,
			Name: user.Name,
		}
		// Prepopulate our users map because we can easily do so.
		// If a user doesn't exist, set it.
		u, err := a.Robot.Users.Get(user.ID)
		if err != nil {
			a.Robot.Users.Set(user.ID, newUser)
		}

		// If the user doesn't match completely (say, if someone changes their name),
		// then adjust what we have stored.
		if u.Name != user.Name {
			a.Robot.Users.Set(user.ID, newUser)
		}
	}

	// This can be nice to have at times.
	hal.Logger.Debugf("Stored users: %s", a.Robot.Users.All())
	hal.Logger.Debugf("Admin users: %s", a.Robot.Auth.Admins())
	hal.Logger.Debugf("Handlers: %s", a.Robot.Handlers())

	a.rtm = api.NewRTM()
	go a.rtm.ManageConnection()

	for {
		select {
		case msg := <-a.rtm.IncomingEvents:
			// Handle events
			switch msg.Data.(type) {

			case *slack.HelloEvent:
				hal.Logger.Debugf("Received HelloEvent")

			case *slack.ConnectingEvent:
				hal.HealthStatus.AdapterStatus = "connecting"
				hal.Logger.Debugf("Connecting...")

			case *slack.ConnectedEvent:
				hal.HealthStatus.AdapterStatus = "connected"
				hal.Logger.Debugf("Connected to Slack.")

			case *slack.DisconnectedEvent:
				hal.HealthStatus.AdapterStatus = "disconnected"
				hal.Logger.Debugf("Disconnected from Slack.")

			case *slack.MessageEvent:
				m := msg.Data.(*slack.MessageEvent)
				hal.Logger.Debugf("MessageEvent: %v", m)
				msg := a.newMessage(m)
				a.Receive(msg)

			case *slack.AckMessage:
				m := msg.Data.(*slack.AckMessage)
				// Returns epoch time
				hal.Logger.Debugf("Message Ack'd: %v", m.Timestamp)

			case *slack.UserTypingEvent:
				m := msg.Data.(*slack.UserTypingEvent)
				// Even DM's have a "channel"
				hal.Logger.Debugf("UserTypingEvent: %v in %v", m.User, m.Channel)

			case *slack.PresenceChangeEvent:
				m := msg.Data.(*slack.PresenceChangeEvent)
				hal.Logger.Debugf("PresenceChangeEvent: %v", m)

			// Come in at regular intervals
			case *slack.LatencyReport:
				m := msg.Data.(*slack.LatencyReport)
				hal.Logger.Debugf("LatencyReport: %v", m.Value)

			case *slack.TeamJoinEvent:
				m := msg.Data.(*slack.TeamJoinEvent)
				hal.Logger.Debugf("TeamJoinEvent: %v", m.User)
				// Add the new member to the user list
				if _, err := a.Robot.Users.Get(m.User.ID); err != nil {
					a.Robot.Users.Set(m.User.ID, hal.User{ID: m.User.ID, Name: m.User.Name})
				}

			default:
				hal.Logger.Debugf("Unhandled: %v, Type: %v", msg.Data, msg.Type)
			}
		}
	}
}

func (a *adapter) newMessage(msg *slack.MessageEvent) *hal.Message {
	user, _ := a.Robot.Users.Get(msg.Msg.User)
	return &hal.Message{
		User: user,
		Room: msg.Msg.Channel,
		Text: msg.Text,
	}
}
