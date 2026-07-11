package main

import (
	"embed"
	"log"

	"video-merger/backend"
	"video-merger/bridge"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var LAssetBundle embed.FS

//go:embed version.json
var LVersionData []byte

func main() {
	app := bridge.LProgramCreate(LVersionData)

	err := wails.Run(&options.App{
		Title:     "Video Merger",
		Width:     1400,
		Height:    900,
		MinWidth:  1200,
		MinHeight: 820,
		Frameless: true,
		AssetServer: &assetserver.Options{
			Assets:     LAssetBundle,
			Middleware: backend.LAssetMiddlewareCreate,
		},
		OnStartup:     app.LProgramStart,
		OnBeforeClose: app.LProgramStop,
		OnShutdown:    app.LProgramShutdown,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}
