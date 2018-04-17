package theater

import (
	"github.com/Synaxis/bfheroesFesl/inter/network"
	"github.com/Synaxis/bfheroesFesl/inter/network/codec"

	"github.com/sirupsen/logrus"
)

type ansRGAM struct {
	GameID  string `fesl:"GID"`
	LobbyID string `fesl:"LID"`
}

// RGAM - A
func (tm *Theater) RGAM(event network.EventClientProcess) {
	logrus.Println("RGAM REQUEST")

	event.Client.Answer(&codec.Packet{
		Message: thtrENCL,
		Content: ansRGAM{
			event.Process.Msg["GID"],
			event.Process.Msg["LID"],
		},
	})
}