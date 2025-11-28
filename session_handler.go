package phpsessgo

// SessionHandler is adoption of PHP SessionHandlerInterface
// For more reference: https://www.php.net/manual/en/class.sessionhandlerinterface.php

type SessionHandler struct {
	database map[string]string
}
type SessionHandlerInterface interface {
	Read(sessionID string) (string, error)
	Write(sessionID string, sessionData string) error
}

func NewSessionHandler() SessionHandlerInterface {
	return &SessionHandler{make(map[string]string)}
}

func (h *SessionHandler) Read(sessionID string) (data string, err error) {
	if value, ok := h.database[sessionID]; ok {
		return value, nil
	}
	return "", nil
}

func (h *SessionHandler) Write(sessionID string, sessionData string) error {
	h.database[sessionID] = sessionData
	return nil
}
