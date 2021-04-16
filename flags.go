package main

import (
	"fmt"
	"io"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
)

func GpshrInstall(install string, hooks string) {
	// abs file path
	p := Absfile(install)
	// let check usr path right or not
	fmt.Printf("\n> gpsher will install into " + p)
	fmt.Println("\n> make sure there is .git folder in the directory\n")

	// sound selection
	fmt.Println(string(colorCyan), "? set ur mp3 file:\n", string(colorReset))
	// exe, _ := os.Executable()
	usr, _ := user.Current()
	res := Dirsearch(usr.HomeDir + spath)

	// show list
	for i, v := range res {
		fmt.Println("[" + strconv.Itoa(i+1) + "] " + string(v))
	}

	fmt.Print("\nenter the number :")

	//input number
	var input int
	var result string

	fmt.Scan(&input)

	for i, v := range res {
		if i+1 == input {
			result = v
			break
		}
	}

	if result == "" {
		fmt.Println(string(colorRed), "! some error occured #009, null is selected", string(colorReset))
	}
	fmt.Printf("\n> " + result + " is selected.\n")
	fmt.Printf("> gpsher installing scripts...\n")

	// check osx or not
	useafplay := Isosx()

	// mkdir, write file
	Mkdir_file(p, useafplay, result, hooks)

	fmt.Println(string(colorGreen), "\n! gpsher installation succeeded\n")
}

func GpshrUninstall(uninstall string, hooks string) {
	// abs file path
	p := Absfile(uninstall)
	// let check usr path right or not
	fmt.Printf("\n> gpsher will uninstall from " + p)
	fmt.Println("\n> make sure git init alrealy in the directory\n")

	if hooks != "checkout" {
		os.Remove(p + "/.git/hooks/pre-" + hooks)
		fmt.Printf("> .git/hooks/pre-" + hooks + " removed\n")
	} else {
		os.Remove(p + "/.git/hooks/post-" + hooks)
		fmt.Printf("> .git/hooks/post-" + hooks + " removed\n")

	}
	fmt.Println(string(colorGreen), "\n! gpsher uninstallation succeeded\n")
}

func GpshrAddSound(imports string) {
	// abs
	original := Absfile(imports)
	_, file := filepath.Split(original)

	// exe, _ := os.Executable()

	usr, _ := user.Current()
	cmd := usr.HomeDir + spath + "/" + file

	from, err := os.Open(original)
	if err != nil {
		fmt.Println(string(colorRed), "! some error occured #006, the file is missing", string(colorReset))
	}
	defer from.Close()
	fmt.Printf("> opened original file\n")

	to, err := os.OpenFile(cmd, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(string(colorRed), "! some error occured #007, permission denied", string(colorReset))
	}
	defer to.Close()

	fmt.Printf("> copying...\n")
	_, err = io.Copy(to, from)
	if err != nil {
		fmt.Println(string(colorRed), "! some error occured #008, permission denied", string(colorReset))
	}
	fmt.Printf("> installing...\n")

	fmt.Println(string(colorGreen), "\n! importing succeeded\n")
}
