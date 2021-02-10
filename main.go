package main

import (
	"flag"
	"fmt"
)

func welcome() {
	// say welcome when nothing entered except gtpsr command
	fmt.Println("\nUsage: gpshr <command>\n")
	fmt.Println("	gpshr -install <foo> -sound <term>  install sound sctipts into <foo> directory and assign <term>")
	fmt.Println("	gpshr -uninstall <foo>              uninstall sound script from <foo>\n")

	fmt.Println("All commands:\n")
	fmt.Println("	-install, -uninstall, -sound\n")

	fmt.Println("git_pusher@0.1\n")
}

func main() {

	// init flags
	install := flag.String("install", "", "install sound sctipts")
	uninstall := flag.String("uninstall", "", "uninstall sound sctipts")
	sound := flag.Int("sound", -1, "a sound that play when git push")
	
	// parse flags
	flag.Parse()

	// command select
	if *install == "" && *uninstall == "" && *sound == -1 {
		welcome()
	} else if *install != ""{

	} else if *uninstall != ""{
		
	}
}