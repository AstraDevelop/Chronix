package main

import (
	"chronix/core"
	"fmt"

	"github.com/fatih/color"
)

func main() {
	green := color.New(color.FgGreen).PrintfFunc()
	blue := color.New(color.FgBlue).PrintfFunc()

	green("\nINFO: Starting Chronix server...\n\n")

	core.LoadAllPlugins("addons")
	blue("INFO: Plugins loaded successfully.\n\n")

	serverAddr := "localhost:8080"
	core.StartServer(serverAddr)
	green(fmt.Sprintf("INFO: Server started at http://%s\n", serverAddr))
}
