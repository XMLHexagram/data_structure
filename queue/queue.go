package queue

import (
	"errors"
)

var (
	ErrorQueueOverflow = errors.New("queue overflow")
	ErrorQueueIsEmpty  = errors.New("queue is empty")
)

type ADT interface {
	Put(data interface{}) error
	Poll() (interface{}, error)
	Len() int
	isEmpty() bool
	GetHead() interface{}
	GetAll() []interface{}
}

type Queue struct {
	Data  []interface{}
	Front int // 队前
	Rear  int // 队尾
	MAX   int
}

// true max is the max+1
func Init(max int) *Queue {
	queue := &Queue{
		Data:  make([]interface{}, max),
		Front: 0,
		Rear:  0,
		MAX:   max,
	}
	return queue
}

func (q *Queue) Put(data interface{}) error {
	tempRear := q.Rear + 1
	tempRear %= q.MAX
	if tempRear == q.Front {
		return ErrorQueueOverflow
	}

	q.Data[q.Rear] = data
	q.Rear = tempRear
	return nil
}

func (q *Queue) Poll() (interface{}, error) {
	if q.isEmpty() {
		return nil, ErrorQueueIsEmpty
	}

	data := q.Data[q.Front]
	q.Front++
	q.Front %= q.MAX

	return data, nil
}

func (q *Queue) Len() int {
	return ((q.Rear - q.Front) + q.MAX) % q.MAX
}

func (q *Queue) isEmpty() bool {
	return q.Front == q.Rear
}

func (q *Queue) GetHead() interface{} {
	if q.Front == q.Rear {
		return nil
	}
	return q.Data[q.Front]
}

func (q Queue) GetAll() []interface{} {
	data := make([]interface{}, 0, 100)
	for q.Front != q.Rear {
		if q.Front == q.MAX {
			q.Front %= q.MAX
		}
		data = append(data, q.Data[q.Front])
		q.Front++
	}
	return data
}
