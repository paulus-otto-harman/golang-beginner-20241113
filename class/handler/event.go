package handler

import (
	"20241113/class/lib"
	"20241113/class/model"
	"20241113/class/service"
	"encoding/json"
	"net/http"
	"strconv"
)

type EventHandler struct {
	EventService service.EventService
}

func InitEventHandler(eventService service.EventService) EventHandler {
	return EventHandler{EventService: eventService}
}

func (handler EventHandler) All(w http.ResponseWriter, r *http.Request) {
	date := r.URL.Query().Get("date")

	page := 1
	var err error
	if q := r.URL.Query().Get("page"); q != "" {
		page, err = strconv.Atoi(q)
	}

	if err != nil {
		lib.JsonResponse(w).Fail(http.StatusBadRequest, "Invalid Page")
		return
	}

	sort := r.URL.Query().Get("sort")

	if err != nil {
		lib.JsonResponse(w).Fail(http.StatusBadRequest, "Invalid Limit")

	}

	totalItems, totalPages, items, err := handler.EventService.All(date, page, sort)
	if err != nil {
		lib.JsonResponse(w).Fail(0, err.Error())
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(model.DataPage{
		Success:    true,
		Page:       page,
		Limit:      6,
		TotalItems: totalItems,
		TotalPages: totalPages,
		Data:       items,
	})
}
