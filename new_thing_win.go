package main

import (
	"fmt"
	"log"

	"github.com/gotk3/gotk3/gtk"
)

const (
	newListTitle         = "New List"
	newListLabel         = "List name"
	newItemTitleTemplate = "Adding a new item to %s"
	newItemLabel         = "Item name"
)

type NewThingData struct {
	Name string
}

type NewThingCallback func(NewThingData)

func new_list_window(callback NewThingCallback) (*gtk.Window, error) {
	return new_thing_window(newListTitle, newItemLabel, callback)
}

func new_item_window(list_name string, callback NewThingCallback) (*gtk.Window, error) {
	return new_thing_window(fmt.Sprintf(newItemTitleTemplate, list_name), newItemLabel, callback)
}

func new_thing_window(title, thingLabel string, callback func(NewThingData)) (*gtk.Window, error) {
	template, err := gtk.BuilderNewFromFile("./new_thing_win.glade")
	if err != nil {
		return nil, err
	}

	err = FindSetLabel(template, "new_thing_win_title", title)
	if err != nil {
		return nil, err
	}

	err = FindSetLabel(template, "new_thing_name_label", thingLabel)
	if err != nil {
		return nil, err
	}

	entry, err := GetEntry(template, "new_thing_name_provided")
	if err != nil {
		return nil, err
	}

	btn, err := GetButton(template, "new_thing_save_btn")
	if err != nil {
		return nil, err
	}

	win, err := GetWindow(template, "new_thing_win")
	if err != nil {
		return nil, err
	}

	btn.Connect("clicked", func(b *gtk.Button) {
		b.SetSensitive(false)
		defer b.SetSensitive(true)

		entry.SetSensitive(false)
		defer entry.SetSensitive(true)

		new_thing, err := entry.GetText()
		if err != nil {
			log.Printf("error getting text: %s\n", err)
			return
		}

		data := NewThingData{
			Name: new_thing,
		}

		callback(data)

		win.Close()
	})

	return win, nil
}
