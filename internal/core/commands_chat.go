package core

import (
	"fmt"
	"strings"

	"github.com/servusdei2018/neon-arena/internal/crayon"
)

// CmdSay handles the say command.
func (g *Game) CmdSay(player *Player, cmd string, args []string) {
	room := g.Arena.GetPlayerLocation(player)

	var include []string
	for _, p := range room.GetPlayers() {
		if p != player.ID.String() {
			include = append(include, p)
		}
	}

	msg := strings.Join(args, " ")
	g.Server.BroadcastTo(crayon.Say(fmt.Sprintf("%s says, '%s'\n", player.Name, msg)), include)
	g.Server.Send(player.ID.String(), crayon.Say(fmt.Sprintf("You say, '%s'\n", msg)))
}
