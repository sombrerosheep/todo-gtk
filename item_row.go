package main

import (
	"github.com/gotk3/gotk3/gtk"
)

const (
	dt_format = "January 01 2006"
)

func NewItemRow(i Item) (*gtk.ListBoxRow, error) {
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

	row, err := gtk.ListBoxRowNew()
	if err != nil {
		return nil, err
	}

	root, err := template.GetObject("root")
	if err != nil {
		return nil, err
	}
	b, err := IsBox(root)
	if err != nil {
		return nil, err
	}

	row.Add(b)

	return row, nil
}
