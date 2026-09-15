package queue

import "testing"

func TestQueueCapHints(t *testing.T) {
	for capHint := range 8 {
		q := New[int](capHint)
		n := 0
		for i := 1; i <= 40; i++ {
			q.Push(i)
			if i%3 == 0 {
				n++
				if got := q.Pop(); got != n {
					t.Fatalf("New(%d): pop got %d, want %d", capHint, got, n)
				}
			}
			if q.Len() != i-n {
				t.Fatalf("New(%d): Len got %d, want %d", capHint, q.Len(), i-n)
			}
		}
		for q.Len() > 0 {
			n++
			if got := q.Pop(); got != n {
				t.Fatalf("New(%d): pop got %d, want %d", capHint, got, n)
			}
		}
	}
}
