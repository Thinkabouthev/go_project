package main

import (
	"strings"
	"time"
)

type TaskProcessor interface {
	Process(task *Task) error
}

type TextProcessor struct {
	repo TaskRepository
}

func NewTextProcessor(repo TaskRepository) *TextProcessor {
	return &TextProcessor{repo: repo}
}

func (p *TextProcessor) Process(task *Task) error {
	if strings.TrimSpace(task.Input) == "" {
		return ErrEmptyInput
	}

	if task.Status == StatusCompleted {
		return ErrAlreadyDone
	}

	words := strings.Fields(task.Input)
	charCount := len(task.Input)

	task.Result = TextResult{
		WordCount: len(words),
		CharCount: charCount,
		IsSecure:  charCount >= 8,
	}

	task.Status = StatusCompleted
	task.ProcessedAt = time.Now()

	return p.repo.Save(task)
}
