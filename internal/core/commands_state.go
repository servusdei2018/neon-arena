package core

import (
	"fmt"

	"github.com/servusdei2018/neon-arena/internal/crayon"
)

// CmdQuit handles the quit command.
func (g *Game) CmdQuit(player *Player, cmd string, args []string) {
	g.Server.Send(player.ID.String(), "Goodbye.")

	delete(g.players, player.ID)
	g.Server.Eject(player.ID.String())

	g.Server.BroadcastTo(fmt.Sprintf("%s disappears in a blinding flash of light!\n", player.Name), g.Arena.GetPlayerLocation(player).GetPlayers())
	g.Server.Broadcast(crayon.Announce(fmt.Sprintf("%s has left the game.", player.Name)))
}
