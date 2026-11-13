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

// maxRandomValue — верхняя граница генерируемых чисел.
const maxRandomValue = 1_000_000_000

// generateRandomElements generates random elements.
// Возвращает слайс длиной size со случайными целыми положительными числами.
// Для size <= 0 возвращает пустой слайс.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}

	data := make([]int, size)
	for i := range data {
		// +1, чтобы числа были строго положительными
		data[i] = rand.Intn(maxRandomValue) + 1
	}

	return data
}

// maximum returns the maximum number of elements.
// Для пустого слайса возвращает 0.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	maxValue := data[0]
	for _, value := range data[1:] {
		if value > maxValue {
			maxValue = value
		}
	}

	return maxValue
}

// maxChunks returns the maximum number of elements in a chunks.
// Делит слайс на CHUNKS частей, ищет максимум каждой части в отдельной горутине,
// а затем находит максимум среди найденных значений.
func maxChunks(data []int) int {
	// если элементов меньше, чем частей, делить нечего
	if len(data) < CHUNKS {
		return maximum(data)
	}

	chunkSize := len(data) / CHUNKS
	maxValues := make([]int, CHUNKS)

	var wg sync.WaitGroup

	for i := 0; i < CHUNKS; i++ {
		start := i * chunkSize
		end := start + chunkSize
		// последняя часть забирает остаток, если длина не делится на CHUNKS
		if i == CHUNKS-1 {
			end = len(data)
		}

		wg.Add(1)
		go func(chunk []int, index int) {
			defer wg.Done()
			// каждая горутина пишет только в свою ячейку — гонки нет
			maxValues[index] = maximum(chunk)
		}(data[start:end], i)
	}

	wg.Wait()

	return maximum(maxValues)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	maxValue := maximum(data)
	elapsed := time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", maxValue, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	maxValue = maxChunks(data)
	elapsed = time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", maxValue, elapsed)
}
