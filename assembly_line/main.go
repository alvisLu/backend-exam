package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Item interface {
	Process()
	Name() string
}

type Item1 struct{ id int }

func (i Item1) Process()     { time.Sleep(1000 * time.Millisecond) }
func (i Item1) Name() string { return fmt.Sprintf("Item1-%02d", i.id) }

type Item2 struct{ id int }

func (i Item2) Process()     { time.Sleep(2000 * time.Millisecond) }
func (i Item2) Name() string { return fmt.Sprintf("Item2-%02d", i.id) }

type Item3 struct{ id int }

func (i Item3) Process()     { time.Sleep(3000 * time.Millisecond) }
func (i Item3) Name() string { return fmt.Sprintf("Item3-%02d", i.id) }

type Employee struct {
	id    int
	count int
}

func (e *Employee) Work(items <-chan Item, wg *sync.WaitGroup) {
	defer wg.Done()
	for item := range items {
		fmt.Printf("[員工 %d] 開始處理 %s\n", e.id, item.Name())
		item.Process()
		fmt.Printf("[員工 %d] 完成處理 %s\n", e.id, item.Name())
		e.count++
	}
}

func main() {
	const itemsPerType = 10
	const numEmployees = 5

	items := make([]Item, 0, itemsPerType*3)
	for i := 1; i <= itemsPerType; i++ {
		items = append(items, Item1{id: i}, Item2{id: i}, Item3{id: i})
	}

	rand.Shuffle(len(items), func(i, j int) { items[i], items[j] = items[j], items[i] })

	ch := make(chan Item, len(items))
	for _, it := range items {
		ch <- it
	}
	close(ch)

	employees := make([]*Employee, numEmployees)
	var wg sync.WaitGroup
	start := time.Now()
	for i := range employees {
		employees[i] = &Employee{id: i + 1}
		wg.Add(1)
		go employees[i].Work(ch, &wg)
	}
	wg.Wait()
	elapsed := time.Since(start)

	fmt.Println("\n=== 統計結果 ===")
	fmt.Printf("總處理時間: %v\n", elapsed)
	total := 0
	for _, e := range employees {
		fmt.Printf("員工 %d 共處理 %d 件物品\n", e.id, e.count)
		total += e.count
	}
	fmt.Printf("總處理物品數: %d 件\n", total)
}
