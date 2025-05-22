package core

import (
	"strings"

	"github.com/servusdei2018/neon-arena/internal/crayon"
)

// HandlerFunc defines the signature for command handler functions.
type HandlerFunc func(player *Player, cmd string, args []string)

// CommandManager manages the available commands and their handlers.
type CommandManager struct {
	handlers map[string]HandlerFunc
}

// NewCommandManager creates a new CommandManager.
func NewCommandManager() *CommandManager {
	return &CommandManager{
		handlers: make(map[string]HandlerFunc),
	}
}

// RegisterCommand registers a new command and its handler.
func (cm *CommandManager) RegisterCommand(name string, handler HandlerFunc) {
	cm.handlers[name] = handler
}

// RegisterCommandAlias registers a new command with an alias and its handler.
func (cm *CommandManager) RegisterCommandAlias(name, alias string, handler HandlerFunc) {
	cm.handlers[name] = handler
	cm.handlers[alias] = handler
}

// ExecuteCommand attempts to execute the given command.
func (g *Game) ExecuteCommand(player *Player, command string) {
	parts := strings.Fields(command)

	switch player.State {
	case PLAYER_STATE_LOGIN:
		if len(parts) == 0 {
			g.Server.Send(player.ID.String(), "\nBy what name would you like to be known? ")
			return
		}

		name := parts[0]
		for _, p := range g.players {
			if p.Name == name {
				g.Server.Send(player.ID.String(), "\nA player by that name already exists!\nBy what name would you like to be known? ")
				return
			}
		}

		player.Name = name
		g.Server.Broadcast(crayon.Announce(name + " has entered the game!"))

		player.State = PLAYER_STATE_INGAME

		spawn := g.Arena.Rooms[g.Arena.SpawnRoom]
		g.Server.BroadcastTo(name+" appears in a blinding flash of light!\n", spawn.GetPlayers())
		spawn.AddPlayer(player)

		g.ExecuteCommand(player, "look")

	case PLAYER_STATE_INGAME:
		if len(parts) == 0 {
			g.Server.Send(player.ID.String(), "Invalid command.\n")
			return
		}
		cmd := parts[0]
		args := parts[1:]
		if handler, ok := g.commandManager.handlers[cmd]; ok {
			handler(player, cmd, args)
			return
		}

		g.Server.Send(player.ID.String(), "Unknown command: "+cmd+"\n")
	}
}
