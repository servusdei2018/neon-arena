package core

import (
	"fmt"

	"github.com/logrusorgru/aurora/v4"
	"github.com/servusdei2018/neon-arena/internal/crayon"
)

// CmdLook handles the look command.
func (g *Game) CmdLook(player *Player, cmd string, args []string) {
	room := g.Arena.GetPlayerLocation(player)

	g.Server.Send(player.ID.String(), fmt.Sprintf("%s\n\t%s\n", aurora.Yellow(room.Name).String(), room.Description))

	for _, pid := range room.GetPlayers() {
		if pid == player.ID.String() {
			continue
		}
		if here, ok := g.players[PlayerID(pid)]; ok {
			g.Server.Send(player.ID.String(), crayon.Presence(here.Name+" is standing here.\n"))
		}
	}

	g.Server.Send(player.ID.String(), fmt.Sprintf("Exits: %s\n", aurora.Yellow(room.ExitString()).String()))
}

// CmdNorth handles the north command.
func (g *Game) CmdNorth(player *Player, cmd string, args []string) {
	from := g.Arena.GetPlayerLocation(player)

	if from.Exits[NORTH] == "" {
		g.Server.Send(player.ID.String(), "You cannot go north.\n")
		return
	}

	to := g.Arena.Rooms[from.Exits[NORTH]]

	from.RemovePlayer(player)
	g.Server.BroadcastTo(fmt.Sprintf("%s walks north.\n", player.Name), from.GetPlayers())
	g.Server.BroadcastTo(fmt.Sprintf("%s walks in from the south.\n", player.Name), to.GetPlayers())
	to.AddPlayer(player)
	g.ExecuteCommand(player, "look")
}

// CmdSouth handles the south command.
func (g *Game) CmdSouth(player *Player, cmd string, args []string) {
	from := g.Arena.GetPlayerLocation(player)

	if from.Exits[SOUTH] == "" {
		g.Server.Send(player.ID.String(), "You cannot go south.\n")
		return
	}

	to := g.Arena.Rooms[from.Exits[SOUTH]]

	from.RemovePlayer(player)
	g.Server.BroadcastTo(fmt.Sprintf("%s walks south.\n", player.Name), from.GetPlayers())
	g.Server.BroadcastTo(fmt.Sprintf("%s walks in from the north.\n", player.Name), to.GetPlayers())
	to.AddPlayer(player)
	g.ExecuteCommand(player, "look")
}

// CmdWest handles the west command.
func (g *Game) CmdWest(player *Player, cmd string, args []string) {
	from := g.Arena.GetPlayerLocation(player)

	if from.Exits[WEST] == "" {
		g.Server.Send(player.ID.String(), "You cannot go west.\n")
		return
	}

	to := g.Arena.Rooms[from.Exits[WEST]]

	from.RemovePlayer(player)
	g.Server.BroadcastTo(fmt.Sprintf("%s walks west.\n", player.Name), from.GetPlayers())
	g.Server.BroadcastTo(fmt.Sprintf("%s walks in from the east.\n", player.Name), to.GetPlayers())
	to.AddPlayer(player)
	g.ExecuteCommand(player, "look")
}

// CmdEast handles the east command.
func (g *Game) CmdEast(player *Player, cmd string, args []string) {
	from := g.Arena.GetPlayerLocation(player)

	if from.Exits[EAST] == "" {
		g.Server.Send(player.ID.String(), "You cannot go east.\n")
		return
	}

	to := g.Arena.Rooms[from.Exits[EAST]]

	from.RemovePlayer(player)
	g.Server.BroadcastTo(fmt.Sprintf("%s walks east.\n", player.Name), from.GetPlayers())
	g.Server.BroadcastTo(fmt.Sprintf("%s walks in from the west.\n", player.Name), to.GetPlayers())
	to.AddPlayer(player)
	g.ExecuteCommand(player, "look")
}
