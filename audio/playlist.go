/*
 * MumbleDJ
 * By Matthieu Grieger
 * audio/playlist.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package audio

// Playlist is an interface that represents all valid playlists of multiple
// audio tracks.
type Playlist interface {
	ID() string
	Title() string
	Service() string
}
