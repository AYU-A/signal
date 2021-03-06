package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pkg/errors"

	"github.com/AYU-A/signal/model"
	"github.com/AYU-A/signal/scheduler"
)

// ScheduleCall accepts a list of rooms to call at certain time, saves them in Postgres and creates scheduler job if not exists yet
func (s *Server) ScheduleCall(w http.ResponseWriter, r *http.Request) {
	var rooms []*model.Room
	if err := s.decode(w, r, &rooms); err != nil {
		s.error(w, r, err, http.StatusBadRequest)
		return
	}
	if len(rooms) == 0 {
		s.error(w, r, errors.New("No rooms provided"), http.StatusBadRequest)
		return
	}
	// Validate room details provided
	for _, room := range rooms {
		if room.RoomNumber == 0 {
			s.error(w, r, fmt.Errorf("No room number specified for guest %s %s", room.Firstname, room.Lastname), http.StatusBadRequest)
			return
		}
		if room.Firstname == "" {
			s.error(w, r, fmt.Errorf("No first name provided for room number %d", room.RoomNumber), http.StatusBadRequest)
			return
		}
		if room.Lastname == "" {
			s.error(w, r, fmt.Errorf("No last name specified for room number %d", room.RoomNumber), http.StatusBadRequest)
			return
		}
		if room.CallTime == "" {
			s.error(w, r, fmt.Errorf("No call time specified for room number %d", room.RoomNumber), http.StatusBadRequest)
			return
		}
		if err := validateCallTime(room.CallTime); err != nil {
			s.error(w, r, fmt.Errorf("Call time (%s) has incorrect format (must be hh:mm) for room number %d", room.CallTime, room.RoomNumber), http.StatusBadRequest)
			return
		}
	}
	// Save metadata to Postgres
	if err := s.db.AddRooms(rooms); err != nil {
		s.error(w, r, err, http.StatusInternalServerError)
		return
	}
	// Schedule cron job
	if err := s.scheduleJob(rooms); err != nil {
		s.error(w, r, err, http.StatusInternalServerError)
		return
	}
	s.respond(w, r, rooms, http.StatusOK)
}

// scheduleJob adds cron job in cloud scheduler
func (s *Server) scheduleJob(rooms []*model.Room) error {
	if len(rooms) == 0 {
		return errors.New("No rooms provided")
	}
	for _, room := range rooms {
		// Create scheduler job
		callRoomURL := s.config.CallRoomEndpoint + "?call_time=" + room.CallTime
		if err := s.scheduler.CreateJob(scheduler.CreateJobArgs{
			RoomNumber:        room.RoomNumber,
			CallTime:          room.CallTime,
			CallURL:           callRoomURL,
			SchedulerLocation: s.config.SchedulerLocation,
			SchedulerTimezone: s.config.SchedulerTimeZone,
		}); err != nil {
			return err
		}
		log.Printf("[Room %d] Created scheduler job for wake up time %s", room.RoomNumber, room.CallTime)
	}
	return nil
}
