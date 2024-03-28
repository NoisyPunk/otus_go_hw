package internalhttp

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/storage"
	"go.uber.org/zap"
	"io"
	"net/http"
)

const (
	day   = "day"
	week  = "week"
	month = "month"
)

func (s *HTTPEventServer) CreateEvent(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	buf := s.readRequest(w, r)
	if buf == nil {
		return
	}
	req := &CreateEventRequest{}
	resp := &CreateEventResponse{}
	errResp := &ErrorResponse{}

	err := json.Unmarshal(buf, req)
	if err != nil {
		errResp.Error.Message = err.Error()
		w.WriteHeader(http.StatusBadRequest)
		s.writeResponse(w, errResp)
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
			errResp.Error.Message = err.Error()
			w.WriteHeader(http.StatusBadRequest)
			s.writeResponse(w, resp)
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
	s.writeResponse(w, resp)
	return
}

func (s *HTTPEventServer) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	buf := s.readRequest(w, r)

	req := &UpdateEventRequest{}
	resp := &UpdateEventResponse{}
	errResp := &ErrorResponse{}

	err := json.Unmarshal(buf, req)
	if err != nil {
		errResp.Error.Message = err.Error()
		w.WriteHeader(http.StatusBadRequest)
		s.writeResponse(w, errResp)
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
			errResp.Error.Message = err.Error()
			w.WriteHeader(http.StatusBadRequest)
			s.writeResponse(w, errResp)
			return
		}
	}
	resp.EventID = req.EventID
	resp.Message = fmt.Sprintf("event updated successfully")

	w.WriteHeader(http.StatusOK)
	s.writeResponse(w, resp)
	return
}

func (s *HTTPEventServer) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	buf := s.readRequest(w, r)

	req := &DeleteEventRequest{}
	resp := &DeleteEventResponse{}
	errResp := &ErrorResponse{}

	err := json.Unmarshal(buf, req)
	if err != nil {
		errResp.Error.Message = err.Error()
		w.WriteHeader(http.StatusBadRequest)
		s.writeResponse(w, errResp)
		return
	}

	err = s.application.DeleteEvent(ctx, req.EventID)
	if err != nil {
		if err != nil {
			errResp.Error.Message = err.Error()
			w.WriteHeader(http.StatusBadRequest)
			s.writeResponse(w, errResp)
			return
		}
	}
	resp.EventID = req.EventID
	resp.Message = fmt.Sprintf("event deleted successfully")

	w.WriteHeader(http.StatusOK)
	s.writeResponse(w, resp)
	return
}

func (s *HTTPEventServer) EventsDailyList(w http.ResponseWriter, r *http.Request) {
	s.collectEventList(w, r, day)
}

func (s *HTTPEventServer) EventsWeeklyList(w http.ResponseWriter, r *http.Request) {
	s.collectEventList(w, r, week)
}

func (s *HTTPEventServer) EventsMonthlyList(w http.ResponseWriter, r *http.Request) {
	s.collectEventList(w, r, month)
}

func (s *HTTPEventServer) collectEventList(w http.ResponseWriter, r *http.Request, period string) {
	ctx := context.Background()
	buf := s.readRequest(w, r)

	req := &EventListRequest{}
	resp := &EventListResponse{}
	errResp := &ErrorResponse{}

	err := json.Unmarshal(buf, req)
	if err != nil {
		errResp.Error.Message = err.Error()
		w.WriteHeader(http.StatusBadRequest)
		s.writeResponse(w, errResp)
		return
	}
	var eventList []storage.Event

	switch period {
	case day:
		eventList, err = s.application.EventsDailyList(ctx, req.DateAndTime, req.UserID)
	case week:
		eventList, err = s.application.EventsWeeklyList(ctx, req.DateAndTime, req.UserID)
	case month:
		eventList, err = s.application.EventsMonthlyList(ctx, req.DateAndTime, req.UserID)
	}
	if err != nil {
		if err != nil {
			errResp.Error.Message = err.Error()
			w.WriteHeader(http.StatusBadRequest)
			s.writeResponse(w, errResp)
			return
		}
	}

	for _, event := range eventList {
		responseItem := &CreateEventResponse{
			ID:           event.ID,
			Title:        event.Title,
			DateAndTime:  event.DateAndTime,
			Duration:     event.Duration,
			Description:  event.Description,
			UserID:       event.UserID,
			TimeToNotify: event.TimeToNotify,
		}
		resp.EventList = append(resp.EventList, responseItem)

	}

	w.WriteHeader(http.StatusOK)
	s.writeResponse(w, resp)
	return

}

func (s *HTTPEventServer) writeResponse(w http.ResponseWriter, resp interface{}) {
	responseBuf, err := json.Marshal(resp)
	if err != nil {
		s.logger.Error("response marshal error:", zap.String("message:", err.Error()))
	}
	_, err = w.Write(responseBuf)
	if err != nil {
		s.logger.Error("response writer error:", zap.String("message:", err.Error()))
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return
}

func (s *HTTPEventServer) readRequest(w http.ResponseWriter, r *http.Request) []byte {
	errResp := &ErrorResponse{}
	if r.Method != http.MethodPost {
		errResp.Error.Message = fmt.Sprintf("method %s not not supported on uri %s", r.Method, r.URL.Path)
		w.WriteHeader(http.StatusMethodNotAllowed)
		s.writeResponse(w, errResp)
		return nil
	}
	buf := make([]byte, r.ContentLength)
	_, err := r.Body.Read(buf)
	if err != nil && err != io.EOF {
		errResp.Error.Message = err.Error()
		w.WriteHeader(http.StatusBadRequest)
		s.writeResponse(w, errResp)
		return nil
	}
	return buf
}
