package main

// import pkg
import (
	"flag"
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

func main() {

	// init flags
	hooks := flag.String("hooks", "", "select hooks (push, commit and checkout)")
	install := flag.String("install", "", "install sound sctipts")
	uninstall := flag.String("uninstall", "", "uninstall sound sctipts")
	imports := flag.String("import", "", "import sound")

	// parse flags
	flag.Parse()

	// cmd selector
	if *install == "" && *uninstall == "" && *imports == "" && *hooks == "" {
		Welcome()
	} else if *install != "" {

		if *hooks == "commit" || *hooks == "push" || *hooks == "checkout" {
			GpshrInstall(*install, *hooks)
		} else {
			fmt.Println(string(colorRed), "! some error occured #012, bad hook parameters", string(colorReset))
		}

	} else if *uninstall != "" {
		GpshrUninstall(*uninstall, *hooks)
	} else if *imports != "" {
		GpshrAddSound(*imports)
	}
}
