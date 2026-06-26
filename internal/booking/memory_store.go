package booking

import "sync"

type MemoryStore struct{
	bookings map[string]Booking
	sync.RWMutex
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		bookings: map[string]Booking{},
	}
}

func (ms *MemoryStore) Book(b Booking) error {
	ms.Lock()
	defer ms.Unlock()

	if _, exists := ms.bookings[b.SeatId]; exists {
		return ErrSeatAlreadyBooked
	}

	ms.bookings[b.SeatId] = b;

	return nil
}

func (ms *MemoryStore) ListBookings(movieId string) []Booking {
	ms.RLock()
	defer ms.RUnlock()

	lisOfBookings := []Booking{}
	for _, book := range ms.bookings {
		if (book.MovieId == movieId) {
			lisOfBookings = append(lisOfBookings, book)
		}
	}

	return lisOfBookings
}