package main

type TaskRepository interface {
	Save(task *Task) error
	GetByID(id string) (*Task, error)
}

type InMemoryRepository struct {
	tasks map[string]*Task
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		tasks: make(map[string]*Task),
	}
}

func (r *InMemoryRepository) Save(task *Task) error {
	r.tasks[task.ID] = task
	return nil
}

func (r *InMemoryRepository) GetByID(id string) (*Task, error) {
	task, ok := r.tasks[id]
	if !ok {
		return nil, ErrTaskNotFound
	}
	return task, nil
}
