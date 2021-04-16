package functions

// import pkg
import (
	"fmt"
)

// variables
var (
	Version  string = "v0.9"
	Revision string = "12/03/2021"
	spath    string = "/.gpshr/sounds"

	colorReset string = "\033[0m"
	colorRed   string = "\033[31m"
	colorGreen string = "\033[32m"
	colorCyan  string = "\033[36m"
)

func Welcome() {
	// say welcome when nothing entered except gtpsr command

	fmt.Println("\n> gpshr@" + Version)

	fmt.Println(string(colorGreen), "\n Usage: gpshr -command\n", string(colorReset))
	fmt.Println("	gpshr -(un)install [directory] -hooks [hooks]   select [hooks] hooks (push, commit and checkout) then (un)install sound sctipts into [directory]")
	fmt.Println("	gpshr -import foo                               import sound from <foo>\n")

	fmt.Println(string(colorGreen), "All commands:\n", string(colorReset))
	fmt.Println("	-install, -uninstall, -import -hooks\n")

}
