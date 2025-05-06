package main

import (
	"context"
	"embed"
	"fmt"
	"os"

	// "fmt"
	"path/filepath"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	app := NewApp()
	of := NewOperateFile()

	appMenu := menu.NewMenuFromItems(

		&menu.MenuItem{
			Label: "File",
			Type:  menu.SubmenuType,
			SubMenu: menu.NewMenuFromItems(
				&menu.MenuItem{
					Label:       "Reload",
					Accelerator: keys.Control("r"),
					Type:        menu.TextType,
					Click: func(cd *menu.CallbackData) {
						fmt.Println("wails menu ctrl")
						runtime.WindowExecJS(of.ctx, "window.reloadMd();")
					},
					Hidden:   false,
					Disabled: false,
				},
				&menu.MenuItem{
					Label:       "Open",
					Accelerator: keys.Control("o"),
					Type:        menu.TextType,
					Click: func(cd *menu.CallbackData) {
						fmt.Println("wails menu open")
						err := of.SetFilePathWithDialog()
						if err != nil {
							fmt.Fprintln(os.Stderr, "Error: Invalid file path")

							return
						}
						runtime.WindowExecJS(of.ctx, "window.reloadMd()")
					},
				},
			),
		},
		&menu.MenuItem{
			Label: "Hot Reload",
			Type:  menu.SubmenuType,
			SubMenu: CreateNewMenu(CreateRadioMenus(
				[]RadioMenuLabelsValues[int]{
					{
						label: "Stop Hot Reload",
						value: 0,
					},
					{
						label: "0.5sec",
						value: 500,
					},
					{
						label: "1sec",
						value: 1000,
					}, {
						label: "2sec",
						value: 2000,
					},
				},
				func(label string, value int) func(*menu.CallbackData) {
					return func(cd *menu.CallbackData) {
						/// TODO
						of.SetHotReloadTime(value)
					}
				})...),
		},
	)

	//
	// Arguments
	//

	if IsProduction() {
		fmt.Println("Prod mode")
		args := os.Args
		if len(args) != 2 {
			// panic("Error: Requiring a file path")
			fmt.Fprintln(os.Stderr, "Error: Requiring a file path")
			os.Exit(0)
		}
		path := args[1]

		path = sanitizePath(path)

		var err0 error
		// Create an instance of the app structure
		of.path, err0 = filepath.Abs(path)

		if err0 != nil {
			panic("Error: Cannot convert to abs path")
		}

		if !of.ExistsFile() {
			fmt.Println("Error: File does not exist")
		}
	}

	//
	// end arguments
	//

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "wails-markdown",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		// BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup: func(ctx context.Context) {
			app.SetContext(ctx)
			of.SetupFileOperation(ctx)
			// AppMenuGlobal.SetContext(ctx)
			// AppMenuGlobal.MenuSet(ctx)
		},
		Bind: []interface{}{
			app,
			of,
		},
		// Menu: &AppMenuGlobal.menu,
		Menu: appMenu,
	})

	if err != nil {
		println("Error:", err.Error())
	}

}
