package utils

import "book-app/app/reqres"

type BookQueue struct {
	items []reqres.BookResponse
}

func (q *BookQueue) Enqueue(book reqres.BookResponse) {
	q.items = append(q.items, book)
}

func (q *BookQueue) Dequeue() (reqres.BookResponse, bool) {
	if len(q.items) == 0 {
		return reqres.BookResponse{}, false
	}
	book := q.items[0]
	q.items = q.items[1:]
	return book, true
}

func (q *BookQueue) IsEmpty() bool {
	return len(q.items) == 0
}
