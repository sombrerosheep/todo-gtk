package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

const (
	listPath = "lists.json"
)

var state *State

func newThingCallbackNoOp(_ NewThingData) {}

type WidgetCallback func()

func onAppShutdown() {
	log.Println("onAppShutdown")

	log.Printf("saving lists to file: %s\n", listPath)
	err := state.SaveToFile(listPath)
	if err != nil {
		log.Printf("Error saving state to file: %s\n", err)
	}
}

func main() {
	var err error
	fmt.Println("Hello, World")

	state, err = NewStateFromFile(listPath)
	if err != nil {
		log.Printf("Error loading data: %s\n", err)
		return
	}

	app, err := gtk.ApplicationNew("com.github.sombrerosheep.todo-gtk", glib.APPLICATION_FLAGS_NONE)
	if err != nil {
		log.Printf("err creating app: %s\n", err)
		return
	}

	app.Connect("activate", func() {
		log.Println("app activate")

		win, err := NewTodoWindow()
		if err != nil {
			log.Printf("error creating todo window: %s\n", err)
			return
		}

		win.ShowAll()
		app.AddWindow(win)
	})

	app.Connect("shutdown", onAppShutdown)

	os.Exit(app.Run(os.Args))
}
