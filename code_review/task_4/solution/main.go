package main

import (
	"container/heap"
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Task struct {
	ID       int
	Priority int // 0 = high, 1 = medium, 2 = low
	Attempts int
	Index    int // для heap.Interface
}

func (t *Task) Execute(ctx context.Context) error {
	delay := time.Duration(rand.Intn(3)+1) * time.Second

	select {
	case <-time.After(delay):
		if rand.Intn(4) == 0 { // иммитация 25% шанса ошибки
			fmt.Printf("❌ Task %d (P%d) failed\n", t.ID, t.Priority)
			return fmt.Errorf("failed")
		}
		fmt.Printf("✅ Task %d (P%d) done in %v\n", t.ID, t.Priority, delay)
		return nil
	case <-ctx.Done():
		fmt.Printf("⛔ Task %d canceled\n", t.ID)
		return ctx.Err()
	}
}

// PriorityQueue реализация очереди с приоритетом для задач
type PriorityQueue []*Task

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// меньший приоритет — более высокий (0 > 1 > 2)
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	task := x.(*Task)
	task.Index = n
	*pq = append(*pq, task)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	task := old[n-1]
	*pq = old[0 : n-1]
	return task
}

// --------- НАЧНИ СЮДА РЕАЛИЗАЦИЮ -----------

func startWorkerPool(ctx context.Context, pq *PriorityQueue, workerCount int, wg *sync.WaitGroup, mu *sync.Mutex, results chan<- string) {
	// Реализуй воркеров:
	// - они должны брать задачи с самым высоким приоритетом
	// - если задача упала — добавить её обратно с Attempts++
	// - максимум 3 попытки на задачу
	// - использовать мьютекс при работе с PriorityQueue
}

func main() {
	rand.Seed(time.Now().UnixNano())

	const (
		totalTasks  = 20
		workerCount = 4
		timeout     = 15 * time.Second
	)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	// Генерация задач
	for i := 1; i <= totalTasks; i++ {
		task := &Task{
			ID:       i,
			Priority: rand.Intn(3), // 0, 1, 2
			Attempts: 0,
		}
		heap.Push(&pq, task)
	}

	var mu sync.Mutex
	var wg sync.WaitGroup
	results := make(chan string, totalTasks)

	// Стартуй worker pool
	startWorkerPool(ctx, &pq, workerCount, &wg, &mu, results)

	// Ожидание завершения
	go func() {
		wg.Wait()
		close(results)
	}()

	success, failed := 0, 0
	for res := range results {
		switch res {
		case "ok":
			success++
		case "fail":
			failed++
		}
	}

	fmt.Println("---------------")
	fmt.Printf("✅ Success: %d\n", success)
	fmt.Printf("❌ Failed: %d\n", failed)
}
