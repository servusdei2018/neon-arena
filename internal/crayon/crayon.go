package crayon

import (
	"fmt"

	"github.com/logrusorgru/aurora/v4"
)

var AnnounceBanner = fmt.Sprint(aurora.Red("["), aurora.BrightRed("ARENA"), aurora.Red("]"))

// Announce formats a server-wide announcement message.
func Announce(msg string) string {
	return fmt.Sprintf("%s %s\n", AnnounceBanner, msg)
}

// Presence formats a presence.
func Presence(msg string) string {
	return aurora.Green(msg).String()
}

// Say formats a say message.
func Say(msg string) string {
	return aurora.Blue(msg).String()
}
