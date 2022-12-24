package main

import (
	"log"

	"github.com/gotk3/gotk3/gtk"
)

const (
	dt_format string = "January 01 2006"
)

func getOnDeleteItemClick(state *State, listName, itemName string, listBox *gtk.ListBox) WidgetCallback {
	return func() {
		err := state.RemoveItemFromList(listName, itemName)
		if err != nil {
			log.Println(err)
			return
		}

		RefreshItems(listName, state.lists[listName], listBox)
	}
}

func NewItemRow(i Item, listName string, parent *gtk.ListBox) (*gtk.ListBoxRow, error) {
	template, err := gtk.BuilderNewFromFile("./item_row.glade")
	if err != nil {
		return nil, err
	}

	err = FindSetLabel(template, "name_label", i.Name)
	if err != nil {
		return nil, err
	}

	err = FindSetLabel(template, "created", i.Created.Format(dt_format))
	if err != nil {
		return nil, err
	}

	o, err := template.GetObject("complete")
	if err != nil {
		return nil, err
	}

	check, err := IsCheckButton(o)
	if err != nil {
		return nil, err
	}

	if i.Completed.IsZero() {
		// Incomplete item
		check.SetLabel("Completed?")
	} else {
		// Complete item
		check.SetLabel(i.Completed.Format(dt_format))
	}

	btn, err := GetButton(template, "delete_item_btn")
	if err != nil {
		return nil, err
	}

	btn.Connect("clicked", getOnDeleteItemClick(state, listName, i.Name, parent))

	row, err := gtk.ListBoxRowNew()
	if err != nil {
		return nil, err
	}

	b, err := GetBox(template, "root")
	if err != nil {
		return nil, err
	}

	row.Add(b)

	return row, nil
}
