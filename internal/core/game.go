package core

import (
	"log"

	"github.com/servusdei2018/neon-arena/internal/server"
)

// Game represents the main game instance.
type Game struct {
	Arena  *Arena
	Server *server.Server

	commandManager *CommandManager
	greeting       *string
	players        map[PlayerID]*Player // PlayerID -> Player
	logger         *log.Logger
}

// NewGame creates a new Game instance.
func NewGame(logger *log.Logger, port string) *Game {
	s, err := server.NewServer(port)
	if err != nil {
		panic(err)
	}

	g := &Game{
		Arena:          NewArena(),
		Server:         s,
		commandManager: NewCommandManager(),
		players:        make(map[PlayerID]*Player),
		logger:         logger,
	}
	g.registerDefaultCommands()

	return g
}

// WithGreeting sets the greeting.
func (g *Game) WithGreeting(greeting string) *Game {
	g.greeting = &greeting
	return g
}

// Handle processes a command received from a player.
func (g *Game) Handle(player *Player, command string) {
	g.ExecuteCommand(player, command)

	if player.State != PLAYER_STATE_LOGIN {
		g.Server.Send(player.ID.String(), "~> ")
	}
}

// ListenAndServe starts the server and handles incoming connections.
func (g *Game) ListenAndServe() {
	go func() {
		if err := g.Server.ListenAndServe(*g.greeting); err != nil {
			g.logger.Fatal(err)
		}
	}()

	go func() {
		for {
			dc := <-g.Server.Disconnects

			player, ok := g.players[PlayerID(dc)]
			if ok {
				g.Handle(player, "quit")
			}
		}
	}()

	for {
		msg := <-g.Server.Queue

		player, ok := g.players[PlayerID(msg.ClientID)]
		if !ok {
			player = NewPlayer(msg.ClientID)
			g.players[player.ID] = player
		}

		g.Handle(player, msg.Text)
	}
}

// registerDefaultCommands registers basic movement and utility commands.
func (g *Game) registerDefaultCommands() {
	g.commandManager.RegisterCommandAlias("look", "l", g.CmdLook)
	g.commandManager.RegisterCommand("quit", g.CmdQuit)
	g.commandManager.RegisterCommand("say", g.CmdSay)

	g.commandManager.RegisterCommandAlias("north", "n", g.CmdNorth)
	g.commandManager.RegisterCommandAlias("south", "s", g.CmdSouth)
	g.commandManager.RegisterCommandAlias("west", "w", g.CmdWest)
	g.commandManager.RegisterCommandAlias("east", "e", g.CmdEast)
}
