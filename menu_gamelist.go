package main

import (
	"fmt"
)

type screenGameList struct {
	entry
}

func buildGameList() scene {
	var list screenGameList
	list.label = "Game List"

	for i := 1; i <= 20; i++ {
		list.children = append(list.children, entry{
			label: fmt.Sprintf("Game #%d Name Here", i),
			icon:  "subsetting",
			callbackOK: func() {
				list.segueNext()
				menu.stack = append(menu.stack, buildExplorer("/"))
			},
		})
	}

	list.segueMount()

	return &list
}

func (s *screenGameList) segueMount() {
	genericSegueMount(&s.entry)
}

func (s *screenGameList) segueNext() {
	genericSegueNext(&s.entry)
}

func (s *screenGameList) segueBack() {
	genericAnimate(&s.entry)
}

func (s *screenGameList) update() {
	genericInput(&s.entry)
}

func (s *screenGameList) render() {
	genericRender(&s.entry)
}