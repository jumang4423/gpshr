package main

// import pkg
import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
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

func Isosx() bool {
	if runtime.GOOS == "darwin" {
		fmt.Println("> afplay detected. ")
		return true
	} else {
		fmt.Println("> aplay detected. ")
		return false
	}
}

func Dirsearch(dir string) []string {

	files, err := ioutil.ReadDir(dir)

	if err != nil {
		fmt.Println(string(colorRed), "! some error occured #002, missing sound files", string(colorReset))
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, Dirsearch(filepath.Join(dir, file.Name()))...)
			continue
		}
		paths = append(paths, filepath.Join(dir, file.Name()))
	}

	return paths
}

func Absfile(path string) string {
	// abs file path
	p, err := filepath.Abs(path)
	if err != nil {
		fmt.Println(string(colorRed), "! some error occured #001, path is broke", string(colorReset))
	}
	return p
}

func Mkdir_file(p string, useafplay bool, result string, hooks string) {

	// make directory
	err := os.MkdirAll(p+"/.git/hooks", 0755)
	if err != nil {
		fmt.Println(string(colorRed), "! some error occured #003, git init first", string(colorReset))
	}
	fmt.Printf("> .git/hooks reinitialized\n")

	if hooks != "checkout" {

		// remove pre-push
		os.Remove(p + "/.git/hooks/pre-" + hooks)
		fmt.Printf("> .git/hooks/pre-" + hooks + " removed\n")

		// create file
		file, err := os.Create(p + "/.git/hooks/pre-" + hooks)
		if err != nil {
			fmt.Println(string(colorRed), "! some error occured #005, cannot create pre-"+hooks+" scripts", string(colorReset))
		}
		fmt.Printf("> .git/hooks/pre-" + hooks + " reinitialized\n")

		defer file.Close()

		// write
		var output string
		if useafplay == true {
			output = "#!/bin/bash\nafplay " + Absfile(result) + " &"
		} else {
			output = "#!/bin/bash\naplay " + Absfile(result) + " &"
		}

		file.Write(([]byte)(output))
		fmt.Printf("> .git/hooks/pre done\n")

		_ = os.Chmod(p+"/.git/hooks/pre-"+hooks, 0700)
	} else {

		// remove post-push
		os.Remove(p + "/.git/hooks/post-" + hooks)
		fmt.Printf("> .git/hooks/post-" + hooks + " removed\n")

		// create file
		file, err := os.Create(p + "/.git/hooks/post-" + hooks)
		if err != nil {
			fmt.Println(string(colorRed), "! some error occured #005, cannot create post-"+hooks+" scripts", string(colorReset))
		}
		fmt.Printf("> .git/hooks/post-" + hooks + " reinitialized\n")

		defer file.Close()

		// write
		var output string
		if useafplay == true {
			output = "#!/bin/bash\nafplay " + Absfile(result) + " &"
		} else {
			output = "#!/bin/bash\naplay " + Absfile(result) + " &"
		}

		file.Write(([]byte)(output))
		fmt.Printf("> .git/hooks/post done\n")

		_ = os.Chmod(p+"/.git/hooks/post-"+hooks, 0700)
	}
	fmt.Printf("> chmod 0700\n")
}
