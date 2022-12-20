package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type State struct {
	lists        ListMap
	selectedList string
}

func (state State) GetSelected() string {
	return state.selectedList
}

func (s *State) SetSelected(listName string) error {
	if _, ok := s.lists[listName]; ok {
		s.selectedList = listName
		return nil
	}

	return fmt.Errorf("Not in the list")
}

func (s State) GetListKeys() []string {
	var keys []string

	for k := range s.lists {
		keys = append(keys, k)
	}

	return keys
}

func (s *State) AddList(listName string) error {
	if _, ok := s.lists[listName]; ok {
		return fmt.Errorf("List \"%s\" already exists", listName)
	}

	s.lists[listName] = []Item{}

	return nil
}

func (s *State) AddItemToList(listName, itemName string) error {
	if _, ok := s.lists[listName]; !ok {
		return fmt.Errorf("List \"%s\" does not exist", listName)
	}

	item := NewItem(itemName)
	s.lists[listName] = append(s.lists[listName], item)

	return nil
}

func (state *State) SaveToFile(path string) error {
	bytes, err := json.Marshal(state.lists)
	if err != nil {
		return err
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(bytes)
	if err != nil {
		return err
	}

	return nil
}

func NewStateFromFile(path string) (*State, error) {
	var state = &State{}

	loaded, err := NewListMapFromFile(listPath)
	if err != nil {
		return state, err
	}

	state.lists = loaded

	err = state.SetSelected(state.lists.First())
	if err != nil {
		return state, err
	}
	fmt.Printf("Setting selected to %s\n", state.selectedList)

	return state, nil
}
