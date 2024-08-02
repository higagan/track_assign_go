package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

type VisitorInfo struct {
	URL     string `json:"url"`
	Visitor string `json:"visitor"`
}

type AnalyticService struct {
	mu       sync.RWMutex
	visitors map[string]map[string]struct{}
}

func NewAnalyticService() *AnalyticService {
	return &AnalyticService{
		visitors: make(map[string]map[string]struct{}),
	}
}

func (s *AnalyticService) CaptureVisitor(w http.ResponseWriter, r *http.Request) {
	var info VisitorInfo
	if err := json.NewDecoder(r.Body).Decode(&info); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	s.mu.Lock()
	if _, exists := s.visitors[info.URL]; !exists {
		s.visitors[info.URL] = make(map[string]struct{})
	}
	s.visitors[info.URL][info.Visitor] = struct{}{}
	s.mu.Unlock()
	w.WriteHeader(http.StatusOK)
}

func (s *AnalyticService) QueryVisitors(w http.ResponseWriter, r *http.Request) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	type URLVisitors struct {
		URL      string `json:"url"`
		Visitors int    `json:"visitors"`
	}
	var data []URLVisitors
	for url, visitors := range s.visitors {
		data = append(data, URLVisitors{URL: url, Visitors: len(visitors)})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func main() {
	service := NewAnalyticService()
	http.HandleFunc("/capture", service.CaptureVisitor)
	http.HandleFunc("/query", service.QueryVisitors)
	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
