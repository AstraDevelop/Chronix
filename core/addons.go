package core

import (
	"log"
	"os"

	lua "github.com/yuin/gopher-lua"
)

func LoadAllPlugins(pluginsDir string) {
	files, err := os.ReadDir(pluginsDir)
	if err != nil {
		log.Fatalf("Failed to read plugins directory: %s", err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		LoadLuaPlugin(pluginsDir + "/" + file.Name())
	}
}

func LoadLuaPlugin(pluginPath string) {
	L := lua.NewState()
	defer L.Close()

	err := L.DoFile(pluginPath)
	if err != nil {
		log.Printf("Error loading Lua plugin %s: %s", pluginPath, err)
	}
}
