package quotation

// Hub maintains the set of active sessions and broadcasts messages to the sessions.
type Hub struct {
	// Registered sessions.
	sessions map[*Session]bool

	// Inbound messages from the sessions.
	broadcast chan []byte

	// Register requests from the sessions.
	register chan *Session

	// Unregister requests from sessions.
	unregister chan *Session
}

// NewHub builds new hub instance
func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Session),
		unregister: make(chan *Session),
		sessions:   make(map[*Session]bool),
	}
}

// Run hub
func (h *Hub) Run() {
	for {
		select {
		case session := <-h.register:
			h.sessions[session] = true
		case session := <-h.unregister:
			if _, ok := h.sessions[session]; ok {
				delete(h.sessions, session)
				close(session.send)
			}
		case message := <-h.broadcast:
			for session := range h.sessions {
				select {
				case session.send <- message:
				default:
					close(session.send)
					delete(h.sessions, session)
				}
			}
		}
	}
}
