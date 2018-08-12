package db

import (
	"sync"

	bolt "github.com/coreos/bbolt"
	boltqueue "jaytaylor.com/bboltqueue"
)

type BoltQueue struct {
	db *bolt.DB
	q  *boltqueue.PQueue
	mu sync.Mutex
}

func NewBoltQueue(db *bolt.DB) *BoltQueue {
	bq := &BoltQueue{
		db: db,
	}
	return bq
}

func (bq *BoltQueue) Open() error {
	bq.mu.Lock()
	defer bq.mu.Unlock()

	if bq.q != nil {
		return nil
	}

	bq.q = boltqueue.NewPQueue(bq.db)

	return nil
}

func (bq *BoltQueue) Close() error {
	bq.mu.Lock()
	defer bq.mu.Unlock()

	if bq.q == nil {
		return nil
	}

	if err := bq.q.Close(); err != nil {
		return err
	}
	bq.q = nil

	return nil
}

func (bq *BoltQueue) Enqueue(table string, priority int, values ...[]byte) error {
	msgs := make([]*boltqueue.Message, 0, len(values))
	for _, value := range values {
		m := boltqueue.NewMessage(string(value))
		msgs = append(msgs, m)
	}
	if err := bq.q.Enqueue(table, priority, msgs...); err != nil {
		return err
	}
	return nil
}

func (bq *BoltQueue) Dequeue(table string) ([]byte, error) {
	v, err := bq.q.Dequeue(table)
	if err != nil {
		return nil, err
	}
	return v.Value, nil
}

func (bq *BoltQueue) Scan(name string, opts *QueueOptions, fn func(value []byte)) error {
	return bq.q.Scan(name, func(m *boltqueue.Message) {
		if opts != nil && m.Priority() != opts.Priority {
			return
		}
		fn(m.Value)
	})
}

func (bq *BoltQueue) Len(name string, priority int) (int, error) {
	return bq.q.Len(name, priority)
}