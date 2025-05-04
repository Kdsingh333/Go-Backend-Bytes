package queue

import (
    "testing"
)

func TestQueueEquals(t *testing.T) {
    q1 := ConstructQueue[int]()
    q2 := ConstructQueue[int]()
    q3 := ConstructQueue[int]()

    q1.Enqueue(1)
    q1.Enqueue(2)
    q1.Enqueue(3)

    q2.Enqueue(1)
    q2.Enqueue(2)
    q2.Enqueue(3)

  
    q3.Enqueue(1)
    q3.Enqueue(2)
    q3.Enqueue(4)

    if !q1.Equals(q2.(*Queue[int])) {
        t.Errorf("Expected queues to be equal, but they are not")
    }

    if q1.Equals(q3.(*Queue[int])) {
        t.Errorf("Expected queues to be different, but they are equal")
    }

    q2.Dequeue() 

    if q1.Equals(q2.(*Queue[int])) {
        t.Errorf("Expected queues with different sizes to be unequal")
    }
}
