/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/move.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package commands

import (
	"errors"
	"fmt"
	"strings"

	"github.com/layeh/gumble/gumble"
)

// MoveCommand is a command that moves the bot from one channel to another.
type MoveCommand struct{}

// Aliases returns the current aliases for the command.
func (c *MoveCommand) Aliases() []string {
	return DJ.BotConfig.Aliases.Move
}

// Description returns the description for the command.
func (c *MoveCommand) Description() string {
	return DJ.BotConfig.Descriptions.Move
}

// IsAdminCommand returns true if the command is only for admin use, and
// returns false otherwise.
func (c *MoveCommand) IsAdminCommand() bool {
	return DJ.BotConfig.Permissions.Move
}

// Execute executes the command with the given user and arguments.
// Return value descriptions:
//    string: A message to be returned to the user upon successful execution.
//    bool:   Whether the message should be private or not. true = private,
//            false = public (sent to whole channel).
//    error:  An error message to be returned upon unsuccessful execution.
//            If no error has occurred, pass nil instead.
// Example return statement:
//    return "This is a private message!", true, nil
func (c *MoveCommand) Execute(user *gumble.User, args ...string) (string, bool, error) {
	if len(args) == 0 {
		return "", true, errors.New("A destination channel must be supplied to move the bot")
	}
	if channels := strings.Split(args[0], "/"); DJ.Client.Channels.Find(channels...) != nil {
		DJ.Client.Self.Move(DJ.Client.Channels.Find(channels...))
	} else {
		return "", true, errors.New("The provided channel does not exist")
	}

	return fmt.Sprintf("You have successfully moved the bot to <b>%s</b>.", args[0]), true, nil
}
