package main

import (
	"errors"
	"testing"
	"time"
)

// 1. Тест успешной обработки и правильности подсчета метрик
func TestProcessSuccess(t *testing.T) {
	repo := NewInMemoryRepository()
	processor := NewTextProcessor(repo)

	task := &Task{
		ID:        "1",
		Status:    StatusPending,
		Input:     "Golang is awesome",
		CreatedAt: time.Now(),
	}

	err := processor.Process(task)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if task.Status != StatusCompleted {
		t.Errorf("expected status %s, got %s", StatusCompleted, task.Status)
	}
	if task.Result.WordCount != 3 {
		t.Errorf("expected 3 words, got %d", task.Result.WordCount)
	}
	if task.Result.CharCount != 17 {
		t.Errorf("expected 17 chars, got %d", task.Result.CharCount)
	}
	if !task.Result.IsSecure {
		t.Errorf("expected IsSecure to be true for 17 chars")
	}

	// Проверяем, что таска обновилась и в репозитории
	saved, _ := repo.GetByID("1")
	if saved.Status != StatusCompleted {
		t.Errorf("expected repo status %s, got %s", StatusCompleted, saved.Status)
	}
}

// 2. Table-driven test: проверка отклонения невалидного входа (пустая строка и пробелы)
func TestEmptyInputValidation(t *testing.T) {
	repo := NewInMemoryRepository()
	processor := NewTextProcessor(repo)

	tests := []struct {
		name  string
		input string
	}{
		{name: "empty string", input: ""},
		{name: "spaces only", input: "   "},
		{name: "tabs and newlines", input: "\t \n "},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			task := &Task{
				ID:        "test-empty",
				Status:    StatusPending,
				Input:     tt.input,
				CreatedAt: time.Now(),
			}

			err := processor.Process(task)
			if !errors.Is(err, ErrEmptyInput) {
				t.Errorf("expected ErrEmptyInput for %q, got %v", tt.input, err)
			}
		})
	}
}

// 3. Тест получения таски по ID и ошибки ErrTaskNotFound
func TestRepositoryGetByID(t *testing.T) {
	repo := NewInMemoryRepository()

	task := &Task{
		ID:     "100",
		Status: StatusPending,
		Input:  "hello",
	}
	repo.Save(task)

	// Успешный поиск
	found, err := repo.GetByID("100")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if found.ID != "100" {
		t.Errorf("expected ID 100, got %s", found.ID)
	}

	// Ошибка при поиске несуществующего ID
	_, err = repo.GetByID("non-existent")
	if !errors.Is(err, ErrTaskNotFound) {
		t.Errorf("expected ErrTaskNotFound, got %v", err)
	}
}

// 4. Тест запрета повторной обработки уже завершенной таски
func TestAlreadyDoneError(t *testing.T) {
	repo := NewInMemoryRepository()
	processor := NewTextProcessor(repo)

	task := &Task{
		ID:        "done-task",
		Status:    StatusCompleted,
		Input:     "password123",
		CreatedAt: time.Now(),
	}

	err := processor.Process(task)
	if !errors.Is(err, ErrAlreadyDone) {
		t.Errorf("expected ErrAlreadyDone, got %v", err)
	}
}
