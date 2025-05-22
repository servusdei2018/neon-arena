package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/servusdei2018/neon-arena/internal/core"

	"github.com/logrusorgru/aurora/v4"
)

var (
	arena *string
	port  *string

	GREETING = fmt.Sprintf("\n%s%s\nBy what name would you like to be known? ",
		aurora.BrightRed(`███▄    █ ▓█████  ▒█████   ███▄    █
██ ▀█   █ ▓█   ▀ ▒██▒  ██▒ ██ ▀█   █
▓██  ▀█ ██▒▒███   ▒██░  ██▒▓██  ▀█ ██▒
▓██▒  ▐▌██▒▒▓█  ▄ ▒██   ██░▓██▒  ▐▌██▒
▒██░   ▓██░░▒████▒░ ████▓▒░▒██░   ▓██░
░ ▒░   ▒ ▒ ░░ ▒░ ░░ ▒░▒░▒░ ░ ▒░   ▒ ▒
░ ░░   ░ ▒░ ░ ░  ░  ░ ▒ ▒░ ░ ░░   ░ ▒░
  ░   ░ ░    ░   ░ ░ ░ ▒     ░   ░ ░
        ░    ░  ░    ░ ░           ░
`), aurora.BrightGreen(`
▄▄▄       ██▀███  ▓█████  ███▄    █  ▄▄▄
▒████▄    ▓██ ▒ ██▒▓█   ▀  ██ ▀█   █ ▒████▄
▒██  ▀█▄  ▓██ ░▄█ ▒▒███   ▓██  ▀█ ██▒▒██  ▀█▄
░██▄▄▄▄██ ▒██▀▀█▄  ▒▓█  ▄ ▓██▒  ▐▌██▒░██▄▄▄▄██
▓█   ▓██▒░██▓ ▒██▒░▒████▒▒██░   ▓██░ ▓█   ▓██▒
▒▒   ▓▒█░░ ▒▓ ░▒▓░░░ ▒░ ░░ ▒░   ▒ ▒  ▒▒   ▓▒█░
 ▒   ▒▒ ░  ░▒ ░ ▒░ ░ ░  ░░ ░░   ░ ▒░  ▒   ▒▒ ░
 ░   ▒     ░░   ░    ░      ░   ░ ░   ░   ▒
     ░  ░   ░        ░  ░         ░       ░  ░
`))
)

func init() {
	arena = flag.String("arena", "./config/aethelburg.json", "path to arena file")
	port = flag.String("port", "4321", "port on which to listen")
	flag.Parse()
}

func main() {
	log := log.Default()

	game := core.NewGame(log, *port).WithGreeting(GREETING)
	game.Arena.FromJSON(*arena)

	game.ListenAndServe()
}
