package menu

import (
	"os/user"

	"github.com/libretro/ludo/settings"

	"github.com/libretro/ludo/core"
	"github.com/libretro/ludo/notifications"
	"github.com/libretro/ludo/state"
)

type screenMain struct {
	entry
}

func buildMainMenu() Scene {
	var list screenMain
	list.label = "Main Menu"

	usr, _ := user.Current()

	if state.Global.CoreRunning {
		list.children = append(list.children, entry{
			label: "Quick Menu",
			icon:  "subsetting",
			callbackOK: func() {
				list.segueNext()
				menu.stack = append(menu.stack, buildQuickMenu())
			},
		})
	}

	list.children = append(list.children, entry{
		label: "Load Core",
		icon:  "subsetting",
		callbackOK: func() {
			list.segueNext()
			menu.stack = append(menu.stack, buildExplorer(
				settings.Current.CoresDirectory,
				[]string{".dll", ".dylib", ".so"},
				func(path string) error {
					err := core.Load(path)
					if err == nil {
						notifications.DisplayAndLog("Core", "Core loaded.")
					}
					return err
				},
				nil,
			))
		},
	})

	list.children = append(list.children, entry{
		label: "Load Game",
		icon:  "subsetting",
		callbackOK: func() {
			list.segueNext()
			menu.stack = append(menu.stack, buildExplorer(usr.HomeDir, nil, core.LoadGame, nil))
		},
	})

	list.children = append(list.children, entry{
		label: "Settings",
		icon:  "subsetting",
		callbackOK: func() {
			list.segueNext()
			menu.stack = append(menu.stack, buildSettings())
		},
	})

	list.children = append(list.children, entry{
		label: "Help",
		icon:  "subsetting",
		callbackOK: func() {
			notifications.DisplayAndLog("Menu", "Not implemented yet.")
		},
	})

	list.children = append(list.children, entry{
		label: "Quit",
		icon:  "subsetting",
		callbackOK: func() {
			vid.Window.SetShouldClose(true)
		},
	})

	list.segueMount()

	return &list
}

func (main *screenMain) Entry() *entry {
	return &main.entry
}

func (main *screenMain) segueMount() {
	genericSegueMount(&main.entry)
}

func (main *screenMain) segueBack() {
	genericAnimate(&main.entry)
}

func (main *screenMain) segueNext() {
	genericSegueNext(&main.entry)
}

func (main *screenMain) update() {
	genericInput(&main.entry)
}

func (main *screenMain) render() {
	genericRender(&main.entry)
}

func (main *screenMain) drawHintBar() {
	genericDrawHintBar()
}
