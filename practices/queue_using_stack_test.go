package dsalgo

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
)

func TestQueueEnqueue(t *testing.T) {
	q := new(QueueWithStack)
	q.New()

	Convey("Test adding a value to queue", t, func() {
		Convey("should add the first value to the queue", func() {
			q.Enqueue(1)
			assert.Equal(t, 1, q.GetFirst().GetValue())
			assert.Equal(t, 1, q.GetLast().GetValue())
			assert.Equal(t, 1, q.Len())
		})

		Convey("should add the 2nd value to the queue", func() {
			q.Enqueue(2)
			assert.Equal(t, 1, q.GetFirst().GetValue())
			assert.Equal(t, 2, q.GetLast().GetValue())
			assert.Equal(t, 2, q.Len())
		})

		Convey("should add the 3rd value to the queue", func() {
			q.Enqueue(3)
			assert.Equal(t, 1, q.GetFirst().GetValue())
			assert.Equal(t, 3, q.GetLast().GetValue())
			assert.Equal(t, 3, q.Len())
		})
	})
}

func TestQueueDequeue(t *testing.T) {
	q := new(QueueWithStack)
	q.New()

	Convey("Test removing a value from queue", t, func() {
		Convey("should return error if queue is empty", func() {
			_, err := q.Dequeue()
			assert.NotEmpty(t, err)
		})

		Convey("should remove when only 1 item in the queue", func() {
			q.Enqueue(1)
			_, err := q.Dequeue()
			assert.Empty(t, q.GetFirst())
			assert.Equal(t, nil, err)
			assert.Empty(t, q.GetLast())
			assert.Equal(t, 0, q.Len())
		})

		Convey("should remove the first one when there're 2 items in the queue", func() {
			q.Enqueue(1)
			q.Enqueue(2)
			q.Dequeue()
			assert.Equal(t, 2, q.GetFirst().GetValue())
			assert.Equal(t, 2, q.GetLast().GetValue())
			assert.Equal(t, 1, q.Len())

			q.Dequeue()
			assert.Empty(t, q.GetFirst())
			assert.Empty(t, q.GetLast())
			assert.Equal(t, 0, q.Len())
		})

		Convey("should remove the first one when there're 3 items in the queue", func() {
			q.Enqueue(1)
			q.Enqueue(2)
			q.Enqueue(3)
			q.Dequeue()
			assert.Equal(t, 2, q.GetFirst().GetValue())
			assert.Equal(t, 3, q.GetLast().GetValue())
			assert.Equal(t, 2, q.Len())
		})
	})
}

func TestGetLast(t *testing.T) {
	q := new(QueueWithStack)
	q.New()

	Convey("Test Get Last function", t, func() {
		Convey("should return the last item in the queue", func() {
			q.Enqueue(1)
			q.Enqueue(2)
			q.Enqueue(3)
			q.Enqueue(4)
			assert.Equal(t, 4, q.GetLast().GetValue())
		})
	})
}
