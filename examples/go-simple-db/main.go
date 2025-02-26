package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

type Record struct {
	ID         int               `json:"id"`
	Attributes map[string]string `json:"attributes"`
}

type DBStore struct {
	data   []Record
	mu     sync.RWMutex
	file   string
	nextID int
}

// NewStore initializes the store
func DBNewStore(file string) *DBStore {
	store := &DBStore{
		data:   []Record{},
		file:   file,
		nextID: 1,
	}
	store.load()
	return store
}

// load initializes data from a file
func (s *DBStore) load() {
	file, err := os.Open(s.file)
	if err == nil {
		defer file.Close()
		json.NewDecoder(file).Decode(&s.data)
		if len(s.data) > 0 {
			s.nextID = s.data[len(s.data)-1].ID + 1
		}
	}
}

// Insert adds a new record
func (s *DBStore) Insert(attributes map[string]string) int {
	s.mu.Lock()
	defer s.mu.Unlock()

	record := Record{ID: s.nextID, Attributes: attributes}
	s.data = append(s.data, record) // ✅ Append works correctly

	s.nextID++
	s.save()

	return record.ID
}

// save persists data to a file
func (s *DBStore) save() {
	file, _ := os.Create(s.file)
	defer file.Close()
	json.NewEncoder(file).Encode(s.data)
}

// Get retrieves a record by ID
func (s *DBStore) Get(id int) (*Record, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	fmt.Println(s.data[id])
	for _, record := range s.data {
		if record.ID == id {
			return &record, nil
		}
	}
	return nil, errors.New("record not found")
}

// Delete removes a record
func (s *DBStore) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, record := range s.data {
		if record.ID == id {
			s.data = append(s.data[:i], s.data[i+1:]...)
			s.save()
			return nil
		}
	}
	return errors.New("record not found")
}

// Update modifies an existing record
func (s *DBStore) Update(id int, attributes map[string]string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i := range s.data {
		if s.data[i].ID == id {
			s.data[i].Attributes = attributes
			s.save()
			return nil
		}
	}
	return errors.New("record not found")
}

// GetAllRecords retrieves all records
func (s *DBStore) GetAllRecords() []Record {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.data
}

func main() {
	store := DBNewStore("mdb.json")
	r := gin.Default()
	// id3 := store.Insert(map[string]string{"name": "Java", "framework": "Spring"})
	// fmt.Println(id3)

	r.POST("/insert", func(c *gin.Context) {
		var attributes map[string]string
		if err := c.ShouldBindJSON(&attributes); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		id := store.Insert(attributes)
		c.JSON(http.StatusOK, gin.H{"message": "Inserted", "id": id})
	})

	r.PUT("/update/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}
		var attributes map[string]string
		if err := c.ShouldBindJSON(&attributes); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := store.Update(id, attributes); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Updated", "id": id})
	})

	r.GET("/get/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}
		record, err := store.Get(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, record)
	})

	r.DELETE("/delete/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}
		if err := store.Delete(id); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Deleted", "id": id})
	})

	r.GET("/all", func(c *gin.Context) {
		records := store.GetAllRecords()
		c.JSON(http.StatusOK, records)
	})

	fmt.Println("Server running on http://localhost:8080")
	r.Run(":8081")
}
