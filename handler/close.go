package handler

import "log"

func (h *Handler) Close() {
	log.Println("Close Server...")
}
