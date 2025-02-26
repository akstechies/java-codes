package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sync"
)

type Store struct {
	data map[string]string
	mu   sync.RWMutex
	file string
}

// RWMutex ~ Allows multiple goroutines to read (RLock) simultaneously.

// NewStore initializes a new key-value store
func NewStore(file string) *Store {
	store := &Store{
		data: make(map[string]string),
		file: file,
	}
	store.load()
	return store
}

// load initializes data from a file
func (s *Store) load() {
	file, err := os.Open(s.file)
	if err == nil {
		defer file.Close()
		json.NewDecoder(file).Decode(&s.data)
	}
}

// Set stores a key-value pair
func (s *Store) Set(key, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = value
	s.save()
}

// save persists data to a file
func (s *Store) save() {
	file, _ := os.Create(s.file)
	defer file.Close()
	json.NewEncoder(file).Encode(s.data)
}

// Get retrieves a value by key
func (s *Store) Get(key string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	val, exists := s.data[key]
	return val, exists
}

// Delete removes a key-value pair
func (s *Store) Delete(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.data, key)
	s.save()
}

// Insert adds a new key-value pair if the key does not exist
func (s *Store) Insert(key, value string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.data[key]; exists {
		return errors.New("key already exists")
	}

	s.data[key] = value
	s.save()
	return nil
}

func store() {
	store := NewStore("db.json")

	store.Set("name", "Go Database")

	getByName, isExist := store.Get("name")
	fmt.Println("Saved: ", getByName, " Exists: ", isExist)

	store.Set("language", "GoLang")

	store.Set("name", "Go ")

	// Test Insert
	if err := store.Insert("names", "Go Database"); err != nil {
		fmt.Println("Insert Error:", err)
	} else {
		fmt.Println("Inserted:", getByName)
	}
}
