package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
// Т. к. в описании задания ничего не сказано про прекращение работы программы,
// и возврат ошибок тоже не предусматривается, логично будет в худшем случае
// возвращать []int{}
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}
	slice := make([]int, size)
	for i := range slice {
		slice[i] = rand.Int()
	}
	return slice
}

// maximum возвращает максимальное число в слайсе.
// По аналогии с generateRandomElements(), логично в худшем случае возвращать 0
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}
	max := data[0]
	for _, v := range data[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

// maxChunks returns the maximum number of elements in a chunks.
// По аналогии с generateRandomElements(), логично в худшем случае возвращать 0
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}
	chunkSize := len(data) / CHUNKS
	chunksMaximums := make([]int, CHUNKS)
	var wg sync.WaitGroup

	for i := 0; i < CHUNKS; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == CHUNKS-1 {
			end = len(data)
		}
		wg.Add(1)
		go func(chunkNumber, chunkStart, chunkEnd int) {
			defer wg.Done()
			chunksMaximums[chunkNumber] = maximum(data[chunkStart:chunkEnd])
		}(i, start, end)
	}
	wg.Wait()
	return maximum(chunksMaximums)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	// ваш код здесь
	slice := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	// ваш код здесь
	start := time.Now()
	max := maximum(slice)
	elapsed := time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d us\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	// ваш код здесь
	start = time.Now()
	max = maxChunks(slice)
	elapsed = time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d us\n", max, elapsed)
}
