package main

import (
	"log"

	"github.com/gotk3/gotk3/gtk"
)

func RefreshLists(lists []string, lb *gtk.ListBox) error {
	lb.GetChildren().Foreach(func(o interface{}) {
		if w, ok := o.(gtk.IWidget); ok {
			lb.Remove(w)
		} else {
			log.Println("could not cast child to widget")
		}
	})

	for _, name := range lists {
		row, err := NewListSelector(name, lb)
		if err != nil {
			return err
		}

		log.Printf("inserting (%s)\n", name)
		lb.Insert(row, -1)
	}

	lb.ShowAll()
	return nil
}

func RefreshItems(listName string, items []Item, lb *gtk.ListBox) error {
	lb.GetChildren().Foreach(func(o interface{}) {
		if w, ok := o.(gtk.IWidget); ok {
			lb.Remove(w)
		} else {
			log.Println("could not cast child to widget")
		}
	})

	for _, item := range items {
		// create the row from item.glade
		row, err := NewItemRow(item, listName, lb)
		if err != nil {
			return err
		}
		// insert each row
		lb.Insert(row, -1)
	}

	lb.ShowAll()
	return nil
}

func getOnNewListClick(list *gtk.ListBox) WidgetCallback {
	return func() {
		fn := func(n NewThingData) {
			state.AddList(n.Name)

			row, err := NewListSelector(n.Name, list)
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
			row, err := NewItemRow(item, listName, list)
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

func NewTodoWindow() (*gtk.Window, error) {
	builder, err := gtk.BuilderNewFromFile("todo.glade")
	if err != nil {
		return nil, err
	}

	signals := map[string]interface{}{}
	builder.ConnectSignals(signals)

	obj, err := builder.GetObject("main_win")
	if err != nil {
		return nil, err
	}

	win, err := IsWindow(obj)
	if err != nil {
		return nil, err
	}

	// Begin List Selection
	listsListBox, err := GetListBox(builder, "list_box")
	if err != nil {
		return nil, err
	}

	err = RefreshLists(state.GetListKeys(), listsListBox)
	if err != nil {
		log.Println(err)
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

		state.SetSelected(selected)

		if err != nil {
			log.Println(err)
			return
		}

		err = RefreshItems(selected, state.lists[selected], lb)
		if err != nil {
			log.Printf("error applying list: %s\n", err)
			return
		}
	})

	newListBtn, err := GetButton(builder, "add_list_btn")
	if err != nil {
		return nil, err
	}
	newListBtn.Connect("clicked", getOnNewListClick(listsListBox))

	newItemBtn, err := GetButton(builder, "add_item_btn")
	if err != nil {
		return nil, err
	}
	newItemBtn.Connect("clicked", getOnNewItemClick(lb))

	err = RefreshItems(state.GetSelected(), state.lists[state.GetSelected()], lb)
	if err != nil {
		return nil, err
	}

	return win, nil
}
