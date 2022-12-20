package main

import (
	"github.com/gotk3/gotk3/gtk"
)

func NewListSelector(name string) (*gtk.ListBoxRow, error) {
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

	root, err := GetBox(template, "root_box")
	if err != nil {
		return nil, err
	}

	root.SetName(name)

	row.Add(root)

	return row, nil
}
