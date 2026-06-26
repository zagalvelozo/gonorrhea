package api

import (
	"crypto/rand"
	"encoding/hex"
	"sync"
	"time"
)

// User represents an authenticated user.
type User struct {
	UID   string `json:"uid"`
	Email string `json:"email,omitempty"`
}

// Message is a single chat message.
type Message struct {
	ID        string `json:"id"`
	Text      string `json:"text"`
	Sender    string `json:"sender"`
	CreatedAt int64  `json:"createdAt"`
	AudioURL  string `json:"audioURL,omitempty"`
}

// Store is a thread-safe in-memory data store.
type Store struct {
	mu       sync.RWMutex
	users    map[string]*User    // uid → user
	messages map[string][]Message // chatID → messages
}

// NewStore creates a new empty Store.
func NewStore() *Store {
	return &Store{
		users:    make(map[string]*User),
		messages: make(map[string][]Message),
	}
}

func randomID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}

// --- Users ---

// CreateUser stores a new user and returns it.
func (s *Store) CreateUser(email string) *User {
	s.mu.Lock()
	defer s.mu.Unlock()
	u := &User{UID: randomID(), Email: email}
	s.users[u.UID] = u
	return u
}

// GetUser retrieves a user by UID.
func (s *Store) GetUser(uid string) *User {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.users[uid]
}

// FindUserByEmail returns the first user with the given email, or nil.
func (s *Store) FindUserByEmail(email string) *User {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, u := range s.users {
		if u.Email == email {
			return u
		}
	}
	return nil
}

// --- Messages ---

// AddMessage appends a message to a chat and returns it with an ID.
func (s *Store) AddMessage(chatID, text, sender string, audioURL string) Message {
	s.mu.Lock()
	defer s.mu.Unlock()
	m := Message{
		ID:        randomID(),
		Text:      text,
		Sender:    sender,
		CreatedAt: time.Now().UnixMilli(),
		AudioURL:  audioURL,
	}
	s.messages[chatID] = append(s.messages[chatID], m)
	return m
}

// GetMessages returns the last `limit` messages for a chat.
func (s *Store) GetMessages(chatID string, limit int) []Message {
	s.mu.RLock()
	defer s.mu.RUnlock()
	msgs := s.messages[chatID]
	if limit <= 0 || limit > len(msgs) {
		limit = len(msgs)
	}
	// Return the last `limit` messages.
	out := make([]Message, limit)
	copy(out, msgs[len(msgs)-limit:])
	return out
}
