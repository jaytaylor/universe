package db

import (
	"fmt"
	"io"
	"sync"
	"testing"
)

func TestPostgresQueue(t *testing.T) {
	config := NewPostgresConfig("dbname=andromeda_test host=/var/run/postgresql")

	q := NewPostgresQueue(config)
	if err := q.Open(); err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := q.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	if err := q.Destroy(TableToCrawl); err != nil {
		t.Fatal(err)
	}

	{
		l0, err := q.Len(TableToCrawl, 0)
		if err != nil {
			t.Fatal(err)
		}
		if expected, actual := 0, l0; actual != expected {
			t.Fatalf("Expected queue=%v len=%v but actual=%v", TableToCrawl, expected, actual)
		}
	}
	{
		if _, err := q.Dequeue(TableToCrawl); err != io.EOF {
			t.Fatalf("Expected err=io.EOF but actual=%[1]T/%[1]s", err)
		}
	}

	if err := q.Enqueue(TableToCrawl, 5, []byte("hello world")); err != nil {
		t.Fatal(err)
	}
	if err := q.Enqueue(TableToCrawl, 5, []byte("<witty test-value here>")); err != nil {
		t.Fatal(err)
	}

	{
		l1, err := q.Len(TableToCrawl, 0)
		if err != nil {
			t.Fatal(err)
		}
		if expected, actual := 2, l1; actual != expected {
			t.Fatalf("Expected queue=%v len=%v but actual=%v", TableToCrawl, expected, actual)
		}
	}
	{
		l2, err := q.Len(TableToCrawl, 5)
		if err != nil {
			t.Fatal(err)
		}
		if expected, actual := 2, l2; actual != expected {
			t.Fatalf("Expected queue=%v len=%v but actual=%v", TableToCrawl, expected, actual)
		}
	}
	{
		l3, err := q.Len(TableToCrawl, 4)
		if err != nil {
			t.Fatal(err)
		}
		if expected, actual := 0, l3; actual != expected {
			t.Fatalf("Expected queue=%v len=%v but actual=%v", TableToCrawl, expected, actual)
		}
	}
	{
		l4, err := q.Len(TableToCrawl, 6)
		if err != nil {
			t.Fatal(err)
		}
		if expected, actual := 0, l4; actual != expected {
			t.Fatalf("Expected queue=%v len=%v but actual=%v", TableToCrawl, expected, actual)
		}
	}

	if err := q.Enqueue(TableToCrawl, 5, []byte("threeve")); err != nil {
		t.Fatal(err)
	}

	{
		v, err := q.Dequeue(TableToCrawl)
		if err != nil {
			t.Fatal(err)
		}
		if expected, actual := "hello world", string(v); actual != expected {
			t.Fatalf("Expected dequeue to produce value=%v but actual=%v", expected, actual)
		}
	}
	{
		v, err := q.Dequeue(TableToCrawl)
		if err != nil {
			t.Fatal(err)
		}
		if expected, actual := "<witty test-value here>", string(v); actual != expected {
			t.Fatalf("Expected dequeue to produce value=%v but actual=%v", expected, actual)
		}
	}
	{
		v, err := q.Dequeue(TableToCrawl)
		if err != nil {
			t.Fatal(err)
		}
		if expected, actual := "threeve", string(v); actual != expected {
			t.Fatalf("Expected dequeue to produce value=%v but actual=%v", expected, actual)
		}
	}
	{
		if _, err := q.Dequeue(TableToCrawl); err != io.EOF {
			t.Fatalf("Expected err=io.EOF but actual=%[1]T/%[1]s", err)
		}
	}
}

func TestPostgresQueueConcurrencyCorrectness(t *testing.T) {
	config := NewPostgresConfig("dbname=andromeda_test host=/var/run/postgresql")

	q := NewPostgresQueue(config)
	if err := q.Open(); err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := q.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	if err := q.Destroy(TableToCrawl); err != nil {
		t.Fatal(err)
	}

	{
		l0, err := q.Len(TableToCrawl, 0)
		if err != nil {
			t.Fatal(err)
		}
		if expected, actual := 0, l0; actual != expected {
			t.Fatalf("Expected queue=%v len=%v but actual=%v", TableToCrawl, expected, actual)
		}
	}

	for i := 0; i < 500; i++ {
		if err := q.Enqueue(TableToCrawl, 5, []byte(fmt.Sprint(i))); err != nil {
			t.Fatal(err)
		}
	}

	var (
		dqMap = map[string]struct{}{}
		mu    sync.Mutex
		wg    sync.WaitGroup
	)

	for i := 0; i < 25; i++ {
		wg.Add(1)
		go func(i int) {
			for {
				dq, err := q.Dequeue(TableToCrawl)
				if err == io.EOF {
					wg.Done()
					return
				}
				mu.Lock()
				if _, ok := dqMap[string(dq)]; ok {
					t.Errorf("dqMap already contained %q", string(dq))
				}
				dqMap[string(dq)] = struct{}{}
				mu.Unlock()
			}
		}(i)
	}

	wg.Wait()
}
