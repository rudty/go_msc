package message

import "fmt"

type M struct {
}
type Handler struct {
	messageChannel chan *M
	HandleMessage  func(*M)
}

func NewHandler() *Handler {
	h := &Handler{
		messageChannel: make(chan *M, 10),
	}
	return h
}

func defaultRecover() {
	r := recover()
	if r != nil {
		fmt.Println(r)
	}
}

func callHandleMessage(h *Handler, m *M) {
	defer defaultRecover()
	if h.HandleMessage != nil {
		h.HandleMessage(m)
	}
}

func (h *Handler) Loop() {
	for {
		m := <-h.messageChannel
		callHandleMessage(h, m)
	}
}

func (h *Handler) Post(m *M) {
	h.messageChannel <- m
}
