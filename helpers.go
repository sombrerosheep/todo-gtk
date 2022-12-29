package main

import (
	"fmt"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func GetLabel(template *gtk.Builder, id string) (*gtk.Label, error) {
	o, err := template.GetObject(id)
	if err != nil {
		return nil, err
	}

	return IsLabel(o)
}

func FindSetLabel(template *gtk.Builder, objName, labelValue string) error {
	label, err := GetLabel(template, objName)
	if err != nil {
		return err
	}

	label.SetLabel(labelValue)

	return nil
}

func GetBox(template *gtk.Builder, id string) (*gtk.Box, error) {
	o, err := template.GetObject(id)
	if err != nil {
		return nil, err
	}

	return IsBox(o)
}

func GetButton(template *gtk.Builder, id string) (*gtk.Button, error) {
	o, err := template.GetObject(id)
	if err != nil {
		return nil, err
	}

	return IsButton(o)
}

func GetEntry(template *gtk.Builder, id string) (*gtk.Entry, error) {
	o, err := template.GetObject(id)
	if err != nil {
		return nil, err
	}

	return IsEntry(o)
}

func GetWindow(template *gtk.Builder, id string) (*gtk.Window, error) {
	o, err := template.GetObject(id)
	if err != nil {
		return nil, err
	}

	return IsWindow(o)
}

func GetListBox(template *gtk.Builder, id string) (*gtk.ListBox, error) {
	o, err := template.GetObject(id)
	if err != nil {
		return nil, err
	}

	return IsListBox(o)
}

func GetCheckButton(template *gtk.Builder, id string) (*gtk.CheckButton, error) {
	o, err := template.GetObject(id)
	if err != nil {
		return nil, err
	}

	return IsCheckButton(o)
}

func IsLabel(o interface{}) (*gtk.Label, error) {
	if l, ok := o.(*gtk.Label); ok {
		return l, nil
	}

	return nil, fmt.Errorf("object is not a label")
}

func IsBox(o interface{}) (*gtk.Box, error) {
	if box, ok := o.(*gtk.Box); ok {
		return box, nil
	}

	return nil, fmt.Errorf("object is not a box")
}

func IsWindow(o glib.IObject) (*gtk.Window, error) {
	if win, ok := o.(*gtk.Window); ok {
		return win, nil
	}

	return nil, fmt.Errorf("object is not a window")
}

func IsListBox(o glib.IObject) (*gtk.ListBox, error) {
	if listBox, ok := o.(*gtk.ListBox); ok {
		return listBox, nil
	}

	return nil, fmt.Errorf("object is not a list box")
}

func IsButton(o glib.IObject) (*gtk.Button, error) {
	if btn, ok := o.(*gtk.Button); ok {
		return btn, nil
	}

	return nil, fmt.Errorf("object is not a button")
}

func IsCheckButton(o glib.IObject) (*gtk.CheckButton, error) {
	if check, ok := o.(*gtk.CheckButton); ok {
		return check, nil
	}

	return nil, fmt.Errorf("object is not a check button")
}

func IsEntry(o glib.IObject) (*gtk.Entry, error) {
	if entry, ok := o.(*gtk.Entry); ok {
		return entry, nil
	}

	return nil, fmt.Errorf("object is not an entry")
}
