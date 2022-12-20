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

func getOnNewListClick(list *gtk.ListBox) WidgetCallback {
	return func() {
		fn := func(n NewThingData) {
			state.AddList(n.Name)

			row, err := NewListSelector(n.Name)
			if err != nil {
				log.Printf("error making selector for list \"%s\": %s\n", n.Name, err)
				return
			}

			list.Insert(row, -1)

			list.ShowAll()
		}

		win, err := new_list_window(fn)
		if err != nil {
			log.Printf("Error creating new list window: %s", err)
			return
		}

		win.ShowAll()
	}
}

func getOnNewItemClick(list *gtk.ListBox) WidgetCallback {
	return func() {
		listName := state.GetSelected()

		fn := func(n NewThingData) {
			state.AddItemToList(listName, n.Name)

			item := NewItem(n.Name)
			row, err := NewItemRow(item)
			if err != nil {
				log.Printf("error making selector for list \"%s\": %s\n", n.Name, err)
				return
			}

			list.Insert(row, -1)

			list.ShowAll()
		}

		win, err := new_item_window(listName, fn)
		if err != nil {
			log.Printf("Error creating new list window: %s", err)
			return
		}

		win.ShowAll()
	}
}

func apply_new_list(state *State, lb *gtk.ListBox) error {
	lb.GetChildren().Foreach(func(o interface{}) {
		if w, ok := o.(gtk.IWidget); ok {
			lb.Remove(w)
		} else {
			log.Println("could not cast child to widget")
		}
	})

	if items, ok := state.lists[state.selectedList]; ok {
		for _, item := range items {
			// create the row from item.glade
			row, err := NewItemRow(item)
			if err != nil {
				return err
			}
			// insert each row
			lb.Insert(row, -1)
		}
	} else {
		return fmt.Errorf("selected item not found in state.lists")
	}

	lb.ShowAll()
	return nil
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

		builder, err := gtk.BuilderNewFromFile("todo.glade")
		if err != nil {
			log.Printf("err created builder: %s\n", err)
			return
		}

		signals := map[string]interface{}{}
		builder.ConnectSignals(signals)

		obj, err := builder.GetObject("main_win")
		if err != nil {
			log.Printf("err getting main_window obj: %s", err)
			return
		}

		win, err := IsWindow(obj)
		if err != nil {
			log.Println(err)
			return
		}

		// Begin List Selection
		listsListBox, err := GetListBox(builder, "list_box")
		if err != nil {
			log.Println(err)
			return
		}

		for _, name := range state.GetListKeys() {
			row, err := NewListSelector(name)
			if err != nil {
				log.Printf("error making selector for list \"%s\": %s\n", name, err)
				return
			}

			listsListBox.Insert(row, -1)
		}

		// Begin populate selected list items
		lb, err := GetListBox(builder, "items")
		if err != nil {
			log.Printf("couldn't get items listbox: %s\n", err)
		}

		listsListBox.Connect("selected-rows-changed", func(listBox *gtk.ListBox) {
			row := listBox.GetSelectedRow()

			wid, err := row.GetChild()
			if err != nil {
				log.Printf("cant get child: %s\n", err)
			}

			box, err := IsBox(wid)
			if err != nil {
				log.Println("not a box")
			}

			selected, err := box.GetName()
			if err != nil {
				log.Println("could not get selected item name")
				return
			}

			if len(selected) < 1 {
				log.Println("row-selected event fired but label was not found")
				return
			}

			if selected == state.selectedList {
				return
			}

			state.selectedList = selected

			err = apply_new_list(state, lb)
			if err != nil {
				fmt.Printf("error applying list: %s\n", err)
				return
			}
		})

		newListBtn, err := GetButton(builder, "add_list_btn")
		if err != nil {
			log.Println(err)
			return
		}
		newListBtn.Connect("clicked", getOnNewListClick(listsListBox))

		newItemBtn, err := GetButton(builder, "add_item_btn")
		if err != nil {
			log.Println(err)
			return
		}
		newItemBtn.Connect("clicked", getOnNewItemClick(lb))

		win.ShowAll()
		app.AddWindow(win)
	})

	app.Connect("shutdown", onAppShutdown)

	os.Exit(app.Run(os.Args))
}
