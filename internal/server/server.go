package server

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"nhooyr.io/websocket"
)

// Default is the structure for the application.
type Default struct {
	SubscriberMessageBuffer int
	Mux                     http.ServeMux
	SubscribersMu           sync.Mutex
	Subscribers             map[*subscriber]struct{}
}

// subscriber is the structure for the subscriber.
type subscriber struct {
	msgs chan []byte
}

// NewServer is a function that creates a new Server.
func NewServer() *Default {
	server := &Default{
		SubscriberMessageBuffer: 10,
		Subscribers:             make(map[*subscriber]struct{}),
	}

	server.Mux.Handle("/", http.FileServer(http.Dir("./htmx")))

	return server
}

// SubscribeHandler is the handler for the subscribe endpoint.
func (s *Default) SubscribeHandler(w http.ResponseWriter, r *http.Request) {
	err := s.Subscribe(r.Context(), w, r)
	if err != nil {
		fmt.Println(err)
		return
	}
}

// AddSubscriber adds a subscriber to the server.
func (s *Default) AddSubscriber(subscriber *subscriber) {
	s.SubscribersMu.Lock()
	s.Subscribers[subscriber] = struct{}{}
	s.SubscribersMu.Unlock()
	fmt.Println("Added subscriber", subscriber)
}

// Subscribe is the method that subscribes to the server.
func (s *Default) Subscribe(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var c *websocket.Conn
	subscriber := &subscriber{
		msgs: make(chan []byte, s.SubscriberMessageBuffer),
	}
	s.AddSubscriber(subscriber)

	c, err := websocket.Accept(w, r, nil)
	if err != nil {
		return err
	}
	defer c.CloseNow()

	ctx = c.CloseRead(ctx)
	for {
		select {
		case msg := <-subscriber.msgs:
			ctx, cancel := context.WithTimeout(ctx, time.Second*5)
			defer cancel()
			err := c.Write(ctx, websocket.MessageText, msg)
			if err != nil {
				return err
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (cs *Default) PublishMsg(msg []byte) {
	cs.SubscribersMu.Lock()
	defer cs.SubscribersMu.Unlock()

	for s := range cs.Subscribers {
		s.msgs <- msg
	}
}
