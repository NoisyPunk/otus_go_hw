package internalhttp

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/storage"
	"io"
	"log"
	"net/http"
)

func (s *HTTPEventServer) CreateEvent(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	resp := &CreateEventResponse{}
	if r.Method != http.MethodPost {
		resp.Error.Message = fmt.Sprintf("method %s not not supported on uri %s", r.Method, r.URL.Path)
		w.WriteHeader(http.StatusMethodNotAllowed)
		writeResponse(w, resp)
		return
	}
	buf := make([]byte, r.ContentLength)
	_, err := r.Body.Read(buf)
	if err != nil && err != io.EOF {
		resp.Error.Message = err.Error()
		w.WriteHeader(http.StatusBadRequest)
		writeResponse(w, resp)
		return
	}

	req := &CreateEventRequest{}

	err = json.Unmarshal(buf, req)
	if err != nil {
		resp.Error.Message = err.Error()
		w.WriteHeader(http.StatusBadRequest)
		writeResponse(w, resp)
		return
	}

	eventData := storage.Event{
		Title:        req.Title,
		DateAndTime:  req.DateAndTime,
		Duration:     req.Duration,
		Description:  req.Description,
		UserID:       req.UserID,
		TimeToNotify: req.TimeToNotify,
	}

	event, err := s.application.CreateEvent(ctx, eventData, eventData.UserID)
	if err != nil {
		if err != nil {
			resp.Error.Message = err.Error()
			w.WriteHeader(http.StatusBadRequest)
			writeResponse(w, resp)
			return
		}
	}
	resp.ID = event.ID
	resp.UserID = event.UserID
	resp.DateAndTime = event.DateAndTime
	resp.TimeToNotify = event.TimeToNotify
	resp.Duration = event.Duration
	resp.Description = event.Description
	resp.Title = event.Title
	resp.Message = fmt.Sprintf("event created successfully")

	w.WriteHeader(http.StatusOK)
	writeResponse(w, resp)
	return
}

func (s *HTTPEventServer) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	resp := &UpdateEventResponse{}
	if r.Method != http.MethodPost {
		resp.Error.Message = fmt.Sprintf("method %s not not supported on uri %s", r.Method, r.URL.Path)
		w.WriteHeader(http.StatusMethodNotAllowed)
		writeResponse(w, resp)
		return
	}
	buf := make([]byte, r.ContentLength)
	_, err := r.Body.Read(buf)
	if err != nil && err != io.EOF {
		resp.Error.Message = err.Error()
		w.WriteHeader(http.StatusBadRequest)
		writeResponse(w, resp)
		return
	}

	req := &UpdateEventRequest{}

	err = json.Unmarshal(buf, req)
	if err != nil {
		resp.Error.Message = err.Error()
		w.WriteHeader(http.StatusBadRequest)
		writeResponse(w, resp)
		return
	}
	eventData := storage.Event{
		ID:           req.EventID,
		Title:        req.Event.Title,
		DateAndTime:  req.Event.DateAndTime,
		Duration:     req.Event.Duration,
		Description:  req.Event.Description,
		UserID:       req.Event.UserID,
		TimeToNotify: req.Event.TimeToNotify,
	}

	err = s.application.UpdateEvent(ctx, req.EventID, eventData)
	if err != nil {
		if err != nil {
			resp.Error.Message = err.Error()
			w.WriteHeader(http.StatusBadRequest)
			writeResponse(w, resp)
			return
		}
	}
	resp.EventID = req.EventID
	resp.Message = fmt.Sprintf("event updated successfully")

	w.WriteHeader(http.StatusOK)
	writeResponse(w, resp)
	return
}

func (s *HTTPEventServer) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (s *HTTPEventServer) EventsDailyList(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (s *HTTPEventServer) EventsWeeklyList(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (s *HTTPEventServer) EventsMonthlyList(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func writeResponse(w http.ResponseWriter, resp interface{}) {
	resBuf, err := json.Marshal(resp)
	if err != nil {
		log.Printf("responce marshal error: %s", err)
	}
	_, err = w.Write(resBuf)
	if err != nil {
		log.Printf("responce marshal error: %s", err)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return
}
