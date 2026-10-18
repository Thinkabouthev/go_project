package main

import (
	"errors"
	"time"
)

const (
	StatusPending   = "PENDING"
	StatusCompleted = "COMPLETED"
)

var (
	ErrEmptyInput   = errors.New("input cannot be empty")
	ErrTaskNotFound = errors.New("task not found")
	ErrAlreadyDone  = errors.New("task is already completed")
)

type TextResult struct {
	WordCount int
	CharCount int
	IsSecure  bool
}

type Task struct {
	ID          string
	Status      string
	Input       string
	Result      TextResult
	CreatedAt   time.Time
	ProcessedAt time.Time
}
