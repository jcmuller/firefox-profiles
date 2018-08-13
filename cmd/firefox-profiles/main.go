package main

import (
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"path"
	"sort"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/spf13/pflag"

	"gopkg.in/ini.v1"
)

var (
	// Version of this tool
	Version          = "0.1.0"
	profilesFileName = "profiles.ini"
	profiles         = readConfig(profilesFileName)
)

func main() {
	listThemes := pflag.BoolP("profiles", "l", false, "List profile names")
	profile := pflag.StringP("open-profile", "o", "", "Open profile")
	version := pflag.BoolP("version", "V", false, "Show version")
	pflag.BoolP("help", "h", false, "Show help")

	pflag.Parse()
	urlToOpen := strings.Join(pflag.Args(), "")

	switch {
	case *listThemes:
		fmt.Println(strings.Join(profileNames(), "\n"))
		os.Exit(0)

	case len(*profile) > 0:
		validateAndOpenUrlWithProfile(profile, urlToOpen)
		os.Exit(0)

	case *version:
		fmt.Printf("%s %s\n", os.Args[0], Version)
		os.Exit(0)

	default:
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS|THEME NAME]\n", os.Args[0])
		pflag.PrintDefaults()

		os.Exit(0)
	}
}

func readConfig(name string) (file *ini.File) {
	filePath := filePathFor(name)
	file, err := ini.LoadSources(ini.LoadOptions{IgnoreInlineComment: true}, filePath)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Fail to read file: %v", err)
		os.Exit(1)
	}

	return
}

func profileNames() (list []string) {
	for _, t := range profiles.Sections() {
		if t.HasKey("Name") {
			list = append(list, t.Key("Name").String())
		}
	}

	sort.Strings(list)

	return
}

func validateAndOpenUrlWithProfile(profile *string, urlToOpen string) {
	var err error

	if !validProfile(*profile) {
		fmt.Fprintln(os.Stderr, "Invalid profile")
		os.Exit(1)
	}

	if len(urlToOpen) == 0 {
		urlToOpen, err = clipboard.ReadAll()

		if err != nil {
			fmt.Fprintln(os.Stderr, "Error getting the clipboard contents")
			os.Exit(1)
		}
	}

	_, err = url.ParseRequestURI(urlToOpen)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid URL '%s'\n", urlToOpen)
		os.Exit(1)
	}

	exec.Command("firefox", "-P", *profile, urlToOpen).Start()
}

func filePathFor(name string) (filePath string) {
	filePath = path.Join(os.Getenv("HOME"), ".mozilla", "firefox", name)

	return
}

func validProfile(profile string) bool {
	for _, p := range profileNames() {
		if p == profile {
			return true
		}
	}

	return false
}
