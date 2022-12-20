package main

import (
	"log"

	"github.com/gotk3/gotk3/gtk"
)

func getOnDeleteListClick(state *State, listName string, listBox *gtk.ListBox) WidgetCallback {
	return func() {
		// remove from state
		err := state.RemoveList(listName)
		if err != nil {
			log.Println(err)
			return
		}

		RefreshLists(state.GetListKeys(), listBox)
	}
}

func NewListSelector(name string, parent *gtk.ListBox) (*gtk.ListBoxRow, error) {
	row, err := gtk.ListBoxRowNew()
	if err != nil {
		return nil, err
	}

	template, err := gtk.BuilderNewFromFile("list_selector.glade")
	if err != nil {
		return nil, err
	}

	err = FindSetLabel(template, "name_label", name)
	if err != nil {
		return nil, err
	}

	btn, err := GetButton(template, "delete_list_btn")
	if err != nil {
		return nil, err
	}

	btn.Connect("clicked", getOnDeleteListClick(state, name, parent))

	root, err := GetBox(template, "root_box")
	if err != nil {
		return nil, err
	}

	root.SetName(name)

	row.Add(root)

	return row, nil
}
