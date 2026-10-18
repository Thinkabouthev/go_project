package main

import (
	"fmt"
	"time"
)

func main() {
	repo := NewInMemoryRepository()
	processor := NewTextProcessor(repo)

	task1 := &Task{
		ID:        "task-1",
		Status:    StatusPending,
		Input:     "MyStrongPassword2026",
		CreatedAt: time.Now(),
	}
	repo.Save(task1)

	err := processor.Process(task1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Task %s [%s]: Words=%d, Chars=%d, IsSecure=%t\n",
		task1.ID, task1.Status, task1.Result.WordCount, task1.Result.CharCount, task1.Result.IsSecure)

	task2 := &Task{
		ID:        "task-2",
		Status:    StatusPending,
		Input:     "Hi Go",
		CreatedAt: time.Now(),
	}
	repo.Save(task2)
	_ = processor.Process(task2)
	fmt.Printf("Task %s [%s]: Words=%d, Chars=%d, IsSecure=%t\n",
		task2.ID, task2.Status, task2.Result.WordCount, task2.Result.CharCount, task2.Result.IsSecure)

	emptyTask := &Task{
		ID:        "task-3",
		Status:    StatusPending,
		Input:     "   ",
		CreatedAt: time.Now(),
	}
	err = processor.Process(emptyTask)
	fmt.Println("Expected error for empty input:", err)
}
