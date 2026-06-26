package booking

import "errors"

type Service struct {
	store BookingStore
}

func NewBookingService(store BookingStore) *Service {
	return &Service{
		store: store,
	}
}

func (s *Service) CreateBooking(b Booking) error {

	if b.SeatId == "" {
		return errors.New("seat_id is required")
	}
	if b.MovieId == "" {
		return errors.New("movie_id is required")
	}
	if b.UserId == "" {
		return errors.New("user_id is required")
	}

	if b.Status == "" {
		b.Status = "confirmed"
	}

	return s.store.Book(b)
}
