/*
 * MumbleDJ
 * By Matthieu Grieger
 * bot/mumbledj.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package bot

import (
	"log"

	"github.com/layeh/gumble/gumble"
	"github.com/layeh/gumble/gumbleffmpeg"
)

// MumbleDJ is a struct that keeps track of all aspects of the bot's state.
type MumbleDJ struct {
	Client       *gumble.Client
	GumbleConfig *gumble.Config
	AudioStream  *gumbleffmpeg.Stream
	BotConfig    *Config
	Queue        *Queue
	Cache        *Cache
	Skips        *SkipTracker
	Log          *log.Logger
	KeepAlive    chan bool
}

// DJ is a struct that keeps track of all aspects of MumbleDJ's environment.
var DJ *MumbleDJ

// NewMumbleDJ initializes and returns a MumbleDJ type.
func NewMumbleDJ() *MumbleDJ {
	return nil
}

// OnConnect event. First moves MumbleDJ into the default channel if one exists.
// The configuration is loaded and the audio stream is initialized.
func (dj *MumbleDJ) OnConnect(e *gumble.ConnectEvent) {

}

// OnDisconnect event. Terminates MumbleDJ process or retries connection if
// automatic connection retries are enabled.
func (dj *MumbleDJ) OnDisconnect(e *gumble.ConnectEvent) {

}

// OnTextMessage event. Checks for command prefix and passes it to the Commander
// if it exists. Ignores the incoming message otherwise.
func (dj *MumbleDJ) OnTextMessage(e *gumble.TextMessageEvent) {

}

// OnUserChange event. Checks UserChange type and adjusts skip trackers to
// reflect the current status of the users on the server.
func (dj *MumbleDJ) OnUserChange(e *gumble.UserChangeEvent) {

}

// SendPrivateMessage sends a private message to the specified user. This method
// verifies that the targeted user is still present in the server before attempting
// to send the message.
func (dj *MumbleDJ) SendPrivateMessage(user *gumble.User, message string) {

}

// CheckDependencies checks whether or not the dependencies for MumbleDJ
// (most notably youtube-dl) are installed and executable. Returns nil if
// no dependencies are misconfigured/missing, returns an error otherwise
// describing the issue.
func (dj *MumbleDJ) CheckDependencies() error {
	return nil
}

func main() {
	DJ = NewMumbleDJ()
}
