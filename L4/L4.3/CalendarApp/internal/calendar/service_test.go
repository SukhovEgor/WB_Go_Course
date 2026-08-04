package calendar

import (
	"testing"
	"time"
)

func newTestService() Service {
	return NewService(nil)
}

func TestCreateEvent(t *testing.T) {

	service := newTestService()

	req := CreateEventRequest{
		UserID: 1,
		Date:   "2030-01-01",
		Title:  "New Year",
	}

	event, err := service.CreateEvent(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if event.ID != 1 {
		t.Fatalf("expected id=1 got=%d", event.ID)
	}

	if event.Title != "New Year" {
		t.Fatalf("wrong title")
	}

	if event.UserID != 1 {
		t.Fatalf("wrong user")
	}
}

func TestCreateEventInvalidInput(t *testing.T) {

	service := newTestService()

	_, err := service.CreateEvent(CreateEventRequest{})

	if err != ErrInvalidInput {
		t.Fatalf("expected ErrInvalidInput")
	}
}

func TestUpdateEvent(t *testing.T) {

	service := newTestService()

	event, _ := service.CreateEvent(CreateEventRequest{
		UserID: 1,
		Date:   "2030-01-01",
		Title:  "Meeting",
	})

	updated, err := service.UpdateEvent(UpdateEventRequest{
		ID:    event.ID,
		Title: "Updated meeting",
	})

	if err != nil {
		t.Fatal(err)
	}

	if updated.Title != "Updated meeting" {
		t.Fatalf("update failed")
	}
}

func TestDeleteEvent(t *testing.T) {

	service := newTestService()

	event, _ := service.CreateEvent(CreateEventRequest{
		UserID: 1,
		Date:   "2030-01-01",
		Title:  "Delete me",
	})

	err := service.DeleteEvent(event.ID)
	if err != nil {
		t.Fatal(err)
	}

	err = service.DeleteEvent(event.ID)

	if err != ErrEventNotFound {
		t.Fatalf("expected ErrEventNotFound")
	}
}

func TestEventsForDay(t *testing.T) {

	service := newTestService()

	service.CreateEvent(CreateEventRequest{
		UserID: 1,
		Date:   "2030-01-01",
		Title:  "Event 1",
	})

	service.CreateEvent(CreateEventRequest{
		UserID: 1,
		Date:   "2030-01-01",
		Title:  "Event 2",
	})

	service.CreateEvent(CreateEventRequest{
		UserID: 1,
		Date:   "2030-01-02",
		Title:  "Event 3",
	})

	date, _ := time.Parse("2006-01-02", "2030-01-01")

	events, err := service.GetEventsForDay(1, date)
	if err != nil {
		t.Fatal(err)
	}

	if len(events) != 2 {
		t.Fatalf("expected 2 events got %d", len(events))
	}
}

func TestEventsForWeek(t *testing.T) {

	service := newTestService()

	service.CreateEvent(CreateEventRequest{
		UserID: 1,
		Date:   "2030-05-13",
		Title:  "Monday",
	})

	service.CreateEvent(CreateEventRequest{
		UserID: 1,
		Date:   "2030-05-15",
		Title:  "Wednesday",
	})

	date, _ := time.Parse("2006-01-02", "2030-05-13")

	events, err := service.GetEventsForWeek(1, date)
	if err != nil {
		t.Fatal(err)
	}

	if len(events) != 2 {
		t.Fatalf("expected 2 events")
	}
}

func TestEventsForMonth(t *testing.T) {

	service := newTestService()

	service.CreateEvent(CreateEventRequest{
		UserID: 1,
		Date:   "2030-06-10",
		Title:  "June",
	})

	service.CreateEvent(CreateEventRequest{
		UserID: 1,
		Date:   "2030-06-20",
		Title:  "June 2",
	})

	service.CreateEvent(CreateEventRequest{
		UserID: 1,
		Date:   "2030-07-10",
		Title:  "July",
	})

	date, _ := time.Parse("2006-01-02", "2030-06-01")

	events, err := service.GetEventsForMonth(1, date)
	if err != nil {
		t.Fatal(err)
	}

	if len(events) != 2 {
		t.Fatalf("expected 2 events")
	}
}