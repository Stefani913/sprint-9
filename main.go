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
	if len(data) == 0 {
		fmt.Println("Длина слайса data не должна равняться 0.")
		return 0
	}
	return slices.Max(data)
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) (int, error) {

	if len(data) == 0 {
		fmt.Println("Длина слайса data не должна равняться 0.")
		return 0, nil
	}

	size := len(data)
	chunkSize := size / CHUNKS

	// добавляем проверку, тк в противном случае slices.Chunk() выбросит panic
	if chunkSize < 1 {
		err := fmt.Errorf("рассчитанное значение чанка не должно быть меньше 1")
		return 0, err
	}

	maxValues := make([]int, 0)
	var wg sync.WaitGroup

	for chunk := range slices.Chunk(data, chunkSize) {
		wg.Add(1)
		go func() {
			defer wg.Done()
			maxValues = append(maxValues, maximum(chunk))
		}()
	}
	wg.Wait()
	return maximum(maxValues), nil
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	randNumSlice := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max := maximum(randNumSlice)
	elapsed := time.Since(start)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)
	// ваш код здесь
	start = time.Now()
	max, err := maxChunks(randNumSlice)
	if err != nil {
		fmt.Println("Возникла ошибка при поиске максимума: ", err)
	}
	elapsed = time.Since(start)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
