package booking

import "errors"


var (
	ErrSeatAlreadyBooked = errors.New("Seat already taken")
)

type Booking struct {
	Id string
	MovieId string
	SeatId string
	UserId string
	Status string
}

type BookingStore interface {
	Book (b Booking) error 
	ListBookings(movieId string) []Booking
}