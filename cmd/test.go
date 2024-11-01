package cmd

import (
	"app/cmd/build"
	"fmt"
	"os"
	"strings"
)

func Test() {
	if len(os.Args) < 2 {
		fmt.Println("No args")

		return
	}

	fmt.Println("== SIMPLE CLI APP")

	// app version
	// app --version
	arg := strings.TrimPrefix(os.Args[1], "--")

	switch arg {
	case "version":
		fmt.Printf("%s %s %s (%s)\n", build.AppName, build.Version, build.Commit[:7], build.Date)
	case "help":
		fmt.Println("Help !")
	default:
		fmt.Println("Unknown command")
	}

	// ...
}
