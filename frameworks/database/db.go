package database

import "go-clean-architecture-boilerplate/entities"

type DB struct {
    // Initialize your database connection here
}

func (db *DB) CreateTask(name string) (*entities.Task, error) {
    // Implement task creation logic
}

func (db *DB) ListTasks() ([]*entities.Task, error) {
    // Implement task listing logic
}