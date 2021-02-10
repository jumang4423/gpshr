package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"io"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strconv"
)

var (
    Version  string = "v0.5"
    Revision string = "10/02/2021"
	spath string = "/.gpshr/sounds"
)

func welcome() {
	// say welcome when nothing entered except gtpsr command
	fmt.Println("\nUsage: gpshr -command\n")
	fmt.Println("gpshr -install foo          install sound sctipts into <foo> directory")
	fmt.Println("gpshr -uninstall foo        uninstall sound script from <foo>")
	fmt.Println("gpshr -import foo           import sound from <foo>\n")

	fmt.Println("All commands:\n")
	fmt.Println("	-install, -uninstall, -import\n")

	fmt.Println("gpshr@" + Version + "\n")
}

func dirsearch(dir string) []string {
	
	files, err := ioutil.ReadDir(dir)
	
	if err != nil {
		fmt.Println("some error occured #002")
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, dirsearch(filepath.Join(dir, file.Name()))...)
			continue
		}
		paths = append(paths, filepath.Join(dir, file.Name()))
	}

	return paths
}

func isosx() bool {
	if runtime.GOOS == "darwin" {
		fmt.Println("> afplay detected. ")
		return true
	} else {
		fmt.Println("> aplay detected. ")
		return false
	}
}

func absfile(path string) string {
	// abs file path
	p, err := filepath.Abs(path)
	if err != nil {
		fmt.Println("! some error occured #001, path is broke")
	}
	return p
}

func mkdir_file(p string, useafplay bool, result string) {

	// make directory
	err := os.MkdirAll(p+"/.git/hooks", 0755)
	if err != nil {
		fmt.Println("! some error occured #003, git init first.")
	}
	fmt.Printf("> .git/hooks reinitialized\n")

	// remove pre-push
	os.Remove(p + "/.git/hooks/pre-push")
	fmt.Printf("> .git/hooks/pre-push removed\n")

	// create file
	file, err := os.Create(p + "/.git/hooks/pre-push")
	if err != nil {
		fmt.Println("! some error occured #005")
	}
	fmt.Printf("> .git/hooks/pre-push reinitialized\n")

	defer file.Close()

	// write
	var output string
	if useafplay == true {
		output = "#!/bin/bash\nafplay " + absfile(result) + " &"
	} else {
		output = "#!/bin/bash\naplay " + absfile(result) + " &"
	}
	
	file.Write(([]byte)(output))
	fmt.Printf("> .git/hooks/pre-push was written properly\n")

	_ = os.Chmod(p+"/.git/hooks/pre-push", 0700)
	fmt.Printf("> chmod 0700\n")
}

func gpshrInstall(install string) {
	// abs file path
	p := absfile(install)
	// let check usr path right or not
	fmt.Printf("\n> gpsher will install into " + p)
	fmt.Println("\n> make sure git init alrealy in the directory\n")

	// sound selection
	fmt.Println("? set ur mp3 file:\n")
	// exe, _ := os.Executable()
	usr, _ := user.Current()
	res := dirsearch(usr.HomeDir + spath)

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

	fmt.Printf("\n> " + result + " is selected.\n")
	fmt.Printf("> gpsher installing scripts...\n")

	// check osx or not
	useafplay := isosx()

	// mkdir, write file
	mkdir_file(p, useafplay, result)

	fmt.Printf("\n! gpsher installation succeeded properly!\n")
	fmt.Printf("! go to " + install + " then git push!\n")
	fmt.Printf("! if u want to uninstall the script, run: gpshr -uninstall foo\n\n")
}

func gpshrUninstall(uninstall string) {
	// abs file path
	p := absfile(uninstall)
	// let check usr path right or not
	fmt.Printf("\n> gpsher will uninstall from " + p)
	fmt.Println("\n> make sure git init alrealy in the directory\n")

	os.Remove(p + "/.git/hooks/pre-push")
	fmt.Printf("> .git/hooks/pre-push removed\n")

	fmt.Printf("\n! gpsher uninstallation succeeded properly!\n\n")
}

func gpshrAddSound(imports string) {
	// abs
	original := absfile(imports)
	_, file := filepath.Split(original)

	// exe, _ := os.Executable()
	
	usr, _ := user.Current()
	cmd := usr.HomeDir + spath + "/" + file

	from, err := os.Open(original)
	if err != nil {
		fmt.Println("! some error occured #006, the file is missing")
	}
	defer from.Close()
	fmt.Printf("> opened original file\n")

	to, err := os.OpenFile(cmd, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("! some error occured #007, permission denied")
	}
	defer to.Close()

	fmt.Printf("> copying...\n")
	_, err = io.Copy(to, from)
	if err != nil {
		fmt.Println("! some error occured #008, permission denied")
	}
	fmt.Printf("> installing...\n")


	fmt.Printf("\n! import succeeded properly!\n")
	fmt.Printf("! if u want to install the script, run: gpshr -install foo\n\n")
}

func main() {

	// init flags
	install := flag.String("install", "", "install sound sctipts")
	uninstall := flag.String("uninstall", "", "uninstall sound sctipts")
	imports := flag.String("import", "", "import sound")

	// parse flags
	flag.Parse()

	// cmd selector
	if *install == "" && *uninstall == "" && *imports == "" {
		welcome()
	} else if *install != "" {
		gpshrInstall(*install)

	} else if *uninstall != "" {
		gpshrUninstall(*uninstall)
	} else if *imports != "" {
		gpshrAddSound(*imports)
	}
}
