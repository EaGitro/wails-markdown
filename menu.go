package main

import (
	"github.com/wailsapp/wails/v2/pkg/menu"
)

// type AppMenuType struct {
// 	ctx  context.Context
// 	menu menu.Menu
// }

// func NewAppMenu() *AppMenuType {
// 	return &AppMenuType{}
// }

// func (m *AppMenuType) SetAppMenu(itemfuncs ...func(ctx context.Context) *menu.MenuItem) *AppMenuType {
// 	var menuSlice []*menu.MenuItem
// 	for _, itemfunc := range itemfuncs {
// 		item := itemfunc(m.ctx)
// 		menuSlice = append(menuSlice, item)

// 	}
// 	copy(m.menu.Items, menuSlice)
// 	fmt.Println("SetAppMenu end")
// 	return m
// }

// func (m *AppMenuType) SetContext(ctx context.Context) {
// 	m.ctx = ctx
// }

// func (m *AppMenuType) MenuSet(ctx context.Context) {
// 	fmt.Println("MenuSet")
// 	runtime.MenuSetApplicationMenu(ctx, &m.menu)
// }

// var AppMenuGlobal = NewAppMenu().SetAppMenu(func(ctx context.Context) *menu.MenuItem {
// 	return &menu.MenuItem{
// 		Label:       "Reload(Ctrl+r)",
// 		Accelerator: keys.Control("r"),
// 		Type:        menu.TextType,
// 		Click: func(cd *menu.CallbackData) {
// 			runtime.WindowExecJS(ctx, "window.reloadMd();")
// 		},
// 		Hidden:   false,
// 		Disabled: false,
// 	}
// })

func CreateNewMenu(menuItems ...*menu.MenuItem) *menu.Menu {
	var menu menu.Menu
	// copy(menu.Items, menuItems)
	for _, menuItem := range menuItems {
		menu.Append(menuItem)
	}
	return &menu
}

type RadioMenuLabelsValues[ValueT int | float64 | string] struct {
	label string
	value ValueT
}

func CreateRadioMenus[ValueT int | string](radioMenuValues []RadioMenuLabelsValues[ValueT], click func(label string, value ValueT) func(*menu.CallbackData)) []*menu.MenuItem {
	var menuItems []*menu.MenuItem
	for _, radioMenuValue := range radioMenuValues {
		menuItems = append(menuItems, &menu.MenuItem{
			Label: radioMenuValue.label,
			Type:  menu.RadioType,
			Click: click(radioMenuValue.label, radioMenuValue.value),
		})
	}
	return menuItems
}
