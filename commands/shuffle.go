/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/shuffle.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package commands

import (
	"errors"

	"github.com/layeh/gumble/gumble"
)

// ShuffleCommand is a command that shuffles the audio queue.
type ShuffleCommand struct{}

// Aliases returns the current aliases for the command.
func (c *ShuffleCommand) Aliases() []string {
	return DJ.BotConfig.Aliases.Shuffle
}

// Description returns the description for the command.
func (c *ShuffleCommand) Description() string {
	return DJ.BotConfig.Descriptions.Shuffle
}

// IsAdminCommand returns true if the command is only for admin use, and
// returns false otherwise.
func (c *ShuffleCommand) IsAdminCommand() bool {
	return DJ.BotConfig.Permissions.Shuffle
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
func (c *ShuffleCommand) Execute(user *gumble.User, args ...string) (string, bool, error) {
	if len(DJ.Queue.Queue) == 0 {
		return "", true, errors.New("There are no tracks currently in the queue")
	}
	if len(DJ.Queue.Queue) <= 2 {
		return "", true, errors.New("There are not enough tracks in the queue to execute a shuffle")
	}

	DJ.Queue.ShuffleTracks()

	return "The audio queue has been shuffled.", false, nil
}
