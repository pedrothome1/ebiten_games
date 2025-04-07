package main

import (
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

var (
	audioContext = audio.NewContext(44100)
	players      = make(map[*wav.Stream]*audio.Player)
)

func MustPlayWavSound(stream *wav.Stream) {
	player, ok := players[stream]
	if !ok {
		p, err := audioContext.NewPlayer(stream)
		if err != nil {
			panic(err)
		}

		player = p
		players[stream] = player
	}

	player.Rewind()
	player.Play()
}
