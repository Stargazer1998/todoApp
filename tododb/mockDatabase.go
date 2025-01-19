package tododb

import "errors"

type MockDatabase struct {
	data   map[int]TodoItem
	nextId int
}

func NewMockDatabase() *MockDatabase {
	return &MockDatabase{
		data:   make(map[int]TodoItem),
		nextId: 1,
	}
}

func (m *MockDatabase) CreateTodoItem(item TodoItem) (TodoItem, error) {
	item.ID = m.nextId
	m.data[item.ID] = item
	m.nextId++
	return item, nil
}

func (m *MockDatabase) GetTodoItem(id int) (TodoItem, error) {
	item, exists := m.data[id]

	if !exists {
		return TodoItem{}, errors.New("Item not found")
	}

	return item, nil
}

func (m *MockDatabase) UpdateTodoItem(id int, item TodoItem) (TodoItem, error) {

	existingItem, exists := m.data[id]
	if !exists {
		return TodoItem{}, errors.New("Item not Found")
	}

	existingItem.ID = item.ID
	existingItem.Completed = item.Completed
	existingItem.Title = item.Title

	return existingItem, nil

}

func (m *MockDatabase) DeleteTodoItem(id int) error {
	if _, exists := m.data[id]; !exists {
		return errors.New("Item not found")
	}
	delete(m.data, id)
	return nil
}
