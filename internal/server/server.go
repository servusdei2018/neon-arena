package server

import (
	"fmt"
	"log"
	"net"
	"os"
	"sync"
)

// Server represents a TCP server.
type Server struct {
	Disconnects chan string
	Queue       chan Message

	listener *net.TCPListener
	clients  map[string]*connection
	mu       sync.RWMutex
	logger   *log.Logger
}

// NewServer creates a new Server instance.
func NewServer(port string) (*Server, error) {
	addr, err := net.ResolveTCPAddr("tcp", ":"+port)
	if err != nil {
		return nil, fmt.Errorf("Error resolving address: %w", err)
	}
	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return nil, fmt.Errorf("Error creating listener: %w", err)
	}
	return &Server{
		Disconnects: make(chan string),
		Queue:       make(chan Message),
		listener:    listener,
		clients:     make(map[string]*connection),
		logger:      log.New(os.Stdout, "[SERVER] ", log.LstdFlags),
	}, nil
}

// Broadcast sends a message to all connected clients.
func (s *Server) Broadcast(message string) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, client := range s.clients {
		client.sendMessage(message)
	}
}

// BroadcastExclude sends a message to all connected clients except the specified ones.
func (s *Server) BroadcastExclude(message string, excludeIDs []string) {
	exclude := make(map[string]bool)
	for _, id := range excludeIDs {
		exclude[id] = true
	}

	s.mu.RLock()
	defer s.mu.RUnlock()
	for id, client := range s.clients {
		if !exclude[id] {
			client.sendMessage(message)
		}
	}
}

// BroadcastTo sends a message to all specified clients.
func (s *Server) BroadcastTo(message string, toIDs []string) {
	include := make(map[string]bool)
	for _, id := range toIDs {
		include[id] = true
	}

	s.mu.RLock()
	defer s.mu.RUnlock()
	for id, client := range s.clients {
		if include[id] {
			client.sendMessage(message)
		}
	}
}

// Eject ejects a client.
func (s *Server) Eject(id string) {
	if client, ok := s.clients[id]; ok {
		client.close()
	}
}

// Send sends a message to a specific client.
func (s *Server) Send(to string, message string) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for id, client := range s.clients {
		if to == id {
			client.sendMessage(message)
			return
		}
	}
}

// ListenAndServe starts the server and listens for incoming connections.
func (s *Server) ListenAndServe(greeting string) error {
	s.logger.Printf("Server listening on %s", s.listener.Addr())
	for {
		conn, err := s.listener.AcceptTCP()
		if err != nil {
			s.logger.Printf("Error accepting connection: %v", err)
			continue
		}
		client, id := s.newConnection(conn, s.Queue, s.logger)

		s.mu.Lock()
		s.clients[id] = client
		s.mu.Unlock()

		s.logger.Printf("Client connected: %s (%s)", conn.RemoteAddr(), id)
		go client.serve()
		client.sendMessage(greeting)
	}
}

// Shutdown closes the listener and any active client connections.
func (s *Server) Shutdown() error {
	s.logger.Println("Shutting down server...")
	err := s.listener.Close()
	if err != nil {
		s.logger.Printf("Error closing listener: %v", err)
	}
	s.mu.Lock()
	for _, clientConn := range s.clients {
		clientConn.close()
	}
	s.mu.Unlock()
	s.logger.Println("Server shutdown complete.")
	return nil
}
