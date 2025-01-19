package tests

import (
	"testing"
	tododatabase "todoApp/tododb"

	"github.com/stretchr/testify/assert"
)

func TestMockDatabase_CreateTodoItem(t *testing.T) {
	db := tododatabase.NewMockDatabase()

	itemTitle := "Test todo Item"
	itemCompleted := false

	item := tododatabase.TodoItem{
		Title:     itemTitle,
		Completed: itemCompleted,
	}

	createdTodo, err := db.CreateTodoItem(item)

	assert.NoError(t, err)

	assert.Equal(t, itemTitle, createdTodo.Title)
	assert.Equal(t, itemCompleted, createdTodo.Completed)
	assert.Greater(t, createdTodo.ID, 0)

}

func TestMockDatabase_UpdateTodoItem(t *testing.T) {
	db := tododatabase.NewMockDatabase()

	itemTitle := "Test todo Item"
	itemCompleted := false

	item, _ := db.CreateTodoItem(tododatabase.TodoItem{
		Title:     itemTitle,
		Completed: itemCompleted,
	})
	println("cretead Item: ", item.ID)
	updatedItemTitle := "Updated todo Item"
	updatedItemCompleted := true

	updatedItem, err := db.UpdateTodoItem(item.ID,
		tododatabase.TodoItem{
			Title:     updatedItemTitle,
			Completed: updatedItemCompleted,
		})

	assert.NoError(t, err)
	assert.Equal(t, updatedItemTitle, updatedItem.Title)
	assert.Equal(t, updatedItemCompleted, updatedItem.Completed)

}

func TestDataBase_DeleteTodoItem(t *testing.T) {
	db := tododatabase.NewMockDatabase()

	itemTitle := "Test todo Item"
	itemCompleted := false

	item := tododatabase.TodoItem{
		Title:     itemTitle,
		Completed: itemCompleted,
	}

	createdItem, _ := db.CreateTodoItem(item)

	err := db.DeleteTodoItem(createdItem.ID)

	assert.NoError(t, err)
}
