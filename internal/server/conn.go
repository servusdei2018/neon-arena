package server

import (
	"bufio"
	"log"
	"net"
	"time"

	"github.com/google/uuid"
)

// Message represents a message received from a client.
type Message struct {
	ClientID string
	Text     string
}

// connection represents a single TCP connection.
type connection struct {
	conn     net.Conn
	id       string
	logger   *log.Logger
	outbound chan string
	queue    chan<- Message
	reader   *bufio.Reader
	server   *Server
}

// newConnection creates a new connection instance.
func (s *Server) newConnection(conn net.Conn, queue chan<- Message, logger *log.Logger) (*connection, string) {
	id := uuid.New().String()

	return &connection{
		conn:     conn,
		id:       id,
		logger:   logger,
		outbound: make(chan string),
		queue:    queue,
		reader:   bufio.NewReader(conn),
		server:   s,
	}, id
}

// serve processes a connection's inbound and outbound messages.
func (c *connection) serve() {
	go c.writeLoop()

	for {
		c.conn.SetReadDeadline(time.Now().Add(5 * time.Minute))
		msg, err := c.reader.ReadString('\n')
		if err != nil {
			c.logger.Printf("Error reading from %s: %v", c.conn.RemoteAddr(), err)
			c.close()
			c.server.mu.Lock()
			delete(c.server.clients, c.id)
			c.server.Disconnects <- c.id
			c.server.mu.Unlock()
			return
		}
		msg = msg[:len(msg)-1]
		c.queue <- Message{ClientID: c.id, Text: msg}
	}
}

// sendMessage sends a message to the client.
func (c *connection) sendMessage(message string) {
	c.outbound <- message
}

// close closes the client connection.
func (c *connection) close() {
	err := c.conn.Close()
	if err != nil {
		c.logger.Printf("Error closing connection for client %s (%s): %v", c.conn.RemoteAddr(), c.id, err)
	} else {
		c.logger.Printf("Client disconnected: %s (%s)", c.conn.RemoteAddr(), c.id)
	}
}

func (c *connection) writeLoop() {
	for msg := range c.outbound {
		c.conn.SetWriteDeadline(time.Now().Add(time.Second * 5))
		_, err := c.conn.Write([]byte(msg))
		if err != nil {
			c.logger.Printf("Error sending to %s: %v", c.conn.RemoteAddr(), err)
			c.close()
			c.server.mu.Lock()
			delete(c.server.clients, c.id)
			c.server.Disconnects <- c.id
			c.server.mu.Unlock()
			return
		}
	}
}
