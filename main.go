package main

import (
	"fmt"
	"math/rand"
	"slices"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		fmt.Println("Переменная size не должна быть меньше или равна 0.")
		return make([]int, 0)
	}

	randomElements := make([]int, size)
	for i := range size {
		randomElements[i] = rand.Intn(size)
	}
	return randomElements
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) < 1 {
		fmt.Println("Длина слайса data не должна быть меньше 1.")
		return 0
	}
	return slices.Max(data)
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	var wg sync.WaitGroup
	size := len(data)
	chunkSize := size / CHUNKS
	maxValues := make([]int, CHUNKS)
	index := 0

	wg.Add(CHUNKS)

	for chunk := range slices.Chunk(data, chunkSize) {
		go func() {
			defer wg.Done()
			maxValues[index] = maximum(chunk)
			index++
		}()
	}

	wg.Wait()
	return maximum(maxValues)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	randNumSlice := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max := maximum(randNumSlice)
	end := time.Now()
	elapsed := end.Sub(start)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)
	// ваш код здесь
	start = time.Now()
	max = maxChunks(randNumSlice)
	end = time.Now()
	elapsed = end.Sub(start)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
