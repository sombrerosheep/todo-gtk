package main

import (
	"testing"
	"time"

	"github.com/go-playground/assert/v2"
)

func assert_time_match_to_sec(t *testing.T, expected, actual time.Time) {
	assert.Equal(t, expected.Year(), actual.Year())
	assert.Equal(t, expected.Month(), actual.Month())
	assert.Equal(t, expected.Day(), actual.Day())
	assert.Equal(t, expected.Hour(), actual.Hour())
	assert.Equal(t, expected.Minute(), actual.Minute())
	assert.Equal(t, expected.Second(), actual.Second())
}

func Test_NewItem(t *testing.T) {
	t.Run("NewItem sets name", func(t *testing.T) {
		itemName := "do stuff"

		item := NewItem(itemName)

		assert.Equal(t, itemName, item.Name)
	})

	t.Run("NewItem sets created", func(t *testing.T) {
		itemName := "do stuff"

		nowTime := time.Now()
		item := NewItem(itemName)

		assert_time_match_to_sec(t, nowTime, item.Created)
	})

	t.Run("NewItem sets completed to Zero", func(t *testing.T) {
		itemName := "do stuff"

		item := NewItem(itemName)

		assert.Equal(t, true, item.Completed.IsZero())
	})
}

func Test_Item_Complete(t *testing.T) {
	t.Run("should set completed time stamp", func(t *testing.T) {
		item := NewItem("do stuff")

		nowTime := time.Now()
		item.Complete()

		assert_time_match_to_sec(t, nowTime, item.Completed)
	})
}
