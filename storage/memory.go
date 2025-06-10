package storage

import (
	"math/rand"
	"sync"
	"task-api/models"
	"time"

	"github.com/google/uuid"
)

type MemoryStorage struct {
	tasks map[string]*models.Task
	mu    sync.RWMutex
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		tasks: make(map[string]*models.Task),
	}
}

func (s *MemoryStorage) CreateTask() *models.Task {
	s.mu.Lock()
	defer s.mu.Unlock()

	task := &models.Task{
		ID:        generateID(),
		Status:    models.StatusPending,
		CreatedAt: time.Now(),
	}

	s.tasks[task.ID] = task

	go s.processTask(task)

	return task
}

func (s *MemoryStorage) GetTask(id string) (*models.Task, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	task, exists := s.tasks[id]
	return task, exists
}

func (s *MemoryStorage) GetAllTasks() []*models.Task {
	s.mu.RLock()
	defer s.mu.RUnlock()

	tasks := make([]*models.Task, 0, len(s.tasks))
	for _, task := range s.tasks {
		tasks = append(tasks, task)
	}

	return tasks
}

func (s *MemoryStorage) DeleteTask(id string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.tasks[id]; exists {
		delete(s.tasks, id)
		return true
	}

	return false
}

func (s *MemoryStorage) processTask(task *models.Task) {
	s.mu.Lock()
	task.Status = models.StatusProcessing
	now := time.Now()
	task.StartedAt = &now
	s.mu.Unlock()

	time.Sleep(3*time.Minute + time.Duration(rand.Intn(120))*time.Second)

	s.mu.Lock()
	defer s.mu.Unlock()

	now = time.Now()
	task.CompletedAt = &now
	task.Status = models.StatusCompleted
	task.Result = "Task completed successfully"
	task.CalculateProcessingTime()
}

func generateID() string {
	return uuid.New().String()
}
