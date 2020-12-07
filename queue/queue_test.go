package queue

import (
	"errors"
	"testing"
)

func TestQueue(t *testing.T) {
	q := Init(11)
	if !q.isEmpty() {
		t.Errorf("init wrong on isEmpty")
	}
	if q.Len() != 0 {
		t.Errorf("init wrong on length")
	}

	for i := 1; i <= 10; i++ {
		err := q.Put(i)
		if err != nil {
			t.Errorf("put %d:%s", i, err)
		}
		if q.Len() != i {
			t.Errorf("put %d:wrong length as %d", i, q.Len())
		}
	}

	err := q.Put(11)
	if !errors.Is(err, ErrorQueueOverflow) {
		t.Errorf("try overflow error wrong")
	}

	var elements []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := 1; i <= 10; i++ {
		if q.GetHead() != i {
			t.Errorf("get header wrong at %d:%d", i, q.GetHead())
		}

		data, err := q.Poll()
		if err != nil {
			t.Errorf("poll %d:%s", i, err)
		}
		if data.(int) != i {
			t.Errorf("poll %d:wrong data as %d", i, data)
		}

		elements = elements[1:]
		for k, v := range q.GetAll() {
			if elements[k] != v.(int) {
				t.Errorf("getAll error:expect %v get %v at %d",elements,q.GetAll(),i)
			}
		}
	}

	if q.Len() != 0 {
		t.Errorf("wrong on length after poll all")
	}
}
