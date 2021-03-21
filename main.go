package main

// import pkg
import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strconv"
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

func welcome() {
	// say welcome when nothing entered except gtpsr command

	fmt.Println("\n> gpshr@" + Version)

	fmt.Println(string(colorGreen), "\n Usage: gpshr -command\n", string(colorReset))
	fmt.Println("	gpshr -(un)install [directory] -hooks [hooks]   select [hooks] hooks (push, commit and checkout) then (un)install sound sctipts into [directory]")
	fmt.Println("	gpshr -import foo                               import sound from <foo>\n")

	fmt.Println(string(colorGreen), "All commands:\n", string(colorReset))
	fmt.Println("	-install, -uninstall, -import -hooks\n")

}

func dirsearch(dir string) []string {

	files, err := ioutil.ReadDir(dir)

	if err != nil {
		fmt.Println(string(colorRed), "! some error occured #002, missing sound files", string(colorReset))
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
		fmt.Println(string(colorRed), "! some error occured #001, path is broke", string(colorReset))
	}
	return p
}

func mkdir_file(p string, useafplay bool, result string, hooks string) {

	// make directory
	err := os.MkdirAll(p+"/.git/hooks", 0755)
	if err != nil {
		fmt.Println(string(colorRed), "! some error occured #003, git init first", string(colorReset))
	}
	fmt.Printf("> .git/hooks reinitialized\n")

    if(hooks != "checkout") {

	    // remove pre-push
	    os.Remove(p + "/.git/hooks/pre-" + hooks)
	    fmt.Printf("> .git/hooks/pre-" + hooks + " removed\n")
    
	    // create file
	    file, err := os.Create(p + "/.git/hooks/pre-" + hooks)
	    if err != nil {
		    fmt.Println(string(colorRed), "! some error occured #005, cannot create pre-" + hooks +" scripts", string(colorReset))
	    }
	    fmt.Printf("> .git/hooks/pre-" + hooks + " reinitialized\n")

	    defer file.Close()

	    // write
	    var output string
	    if useafplay == true {
		    output = "#!/bin/bash\nafplay " + absfile(result) + " &"
	    } else {
		    output = "#!/bin/bash\naplay " + absfile(result) + " &"
	    }

	    file.Write(([]byte)(output))
	    fmt.Printf("> .git/hooks/pre done\n")

	    _ = os.Chmod(p+"/.git/hooks/pre-" + hooks, 0700)
    } else {
    
	    // remove post-push
        os.Remove(p + "/.git/hooks/post-" + hooks)
	    fmt.Printf("> .git/hooks/post-" + hooks + " removed\n")
    
	    // create file
	    file, err := os.Create(p + "/.git/hooks/post-" + hooks)
	    if err != nil {
		    fmt.Println(string(colorRed), "! some error occured #005, cannot create post-" + hooks +" scripts", string(colorReset))
	    }
	    fmt.Printf("> .git/hooks/post-" + hooks + " reinitialized\n")

	    defer file.Close()

	    // write
	    var output string
	    if useafplay == true {
		    output = "#!/bin/bash\nafplay " + absfile(result) + " &"
	    } else {
		    output = "#!/bin/bash\naplay " + absfile(result) + " &"
	    }

	    file.Write(([]byte)(output))
	    fmt.Printf("> .git/hooks/post done\n")

	    _ = os.Chmod(p+"/.git/hooks/post-" + hooks, 0700)
    }
	fmt.Printf("> chmod 0700\n")
}

func gpshrInstall(install string, hooks string) {
	// abs file path
	p := absfile(install)
	// let check usr path right or not
	fmt.Printf("\n> gpsher will install into " + p)
	fmt.Println("\n> make sure there is .git folder in the directory\n")

	// sound selection
	fmt.Println(string(colorCyan), "? set ur mp3 file:\n", string(colorReset))
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

	if result == "" {
		fmt.Println(string(colorRed), "! some error occured #009, null is selected", string(colorReset))
	}
	fmt.Printf("\n> " + result + " is selected.\n")
	fmt.Printf("> gpsher installing scripts...\n")

	// check osx or not
	useafplay := isosx()

	// mkdir, write file
	mkdir_file(p, useafplay, result, hooks)

	fmt.Println(string(colorGreen), "\n! gpsher installation succeeded\n")
}

func gpshrUninstall(uninstall string, hooks string) {
	// abs file path
	p := absfile(uninstall)
	// let check usr path right or not
	fmt.Printf("\n> gpsher will uninstall from " + p)
	fmt.Println("\n> make sure git init alrealy in the directory\n")

	if(hooks != "checkout") {
        os.Remove(p + "/.git/hooks/pre-" + hooks)
	    fmt.Printf("> .git/hooks/pre-" + hooks + " removed\n")
    } else {
        os.Remove(p + "/.git/hooks/post-" + hooks)
        fmt.Printf("> .git/hooks/post-" + hooks + " removed\n")

    }
	fmt.Println(string(colorGreen), "\n! gpsher uninstallation succeeded\n")
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
		welcome()
	} else if *install != "" {
		
		if *hooks == "commit" || *hooks == "push" || *hooks == "checkout" {
			 gpshrInstall(*install, *hooks) 
			} else { 
				fmt.Println(string(colorRed), "! some error occured #012, bad hook parameters", string(colorReset))
			}

	} else if *uninstall != "" {
		gpshrUninstall(*uninstall, *hooks)
	} else if *imports != "" {
		gpshrAddSound(*imports)
	}
}
