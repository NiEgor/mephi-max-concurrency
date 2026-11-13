package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name     string
		size     int
		wantSize int
	}{
		{name: "нулевой размер", size: 0, wantSize: 0},
		{name: "отрицательный размер", size: -5, wantSize: 0},
		{name: "один элемент", size: 1, wantSize: 1},
		{name: "несколько элементов", size: 10, wantSize: 10},
		{name: "много элементов", size: 100_000, wantSize: 100_000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateRandomElements(tt.size)

			if got == nil {
				t.Fatalf("generateRandomElements(%d) вернула nil, ожидался слайс", tt.size)
			}
			if len(got) != tt.wantSize {
				t.Fatalf("generateRandomElements(%d): длина %d, ожидалась %d", tt.size, len(got), tt.wantSize)
			}
			for i, value := range got {
				if value <= 0 {
					t.Fatalf("generateRandomElements(%d): элемент [%d] = %d, ожидалось положительное число", tt.size, i, value)
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "nil слайс", data: nil, want: 0},
		{name: "пустой слайс", data: []int{}, want: 0},
		{name: "один элемент", data: []int{42}, want: 42},
		{name: "максимум в начале", data: []int{9, 3, 5, 1}, want: 9},
		{name: "максимум в конце", data: []int{1, 3, 5, 9}, want: 9},
		{name: "максимум в середине", data: []int{1, 9, 5, 3}, want: 9},
		{name: "одинаковые элементы", data: []int{7, 7, 7}, want: 7},
		{name: "несколько максимумов", data: []int{2, 8, 3, 8}, want: 8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.data); got != tt.want {
				t.Errorf("maximum(%v) = %d, ожидалось %d", tt.data, got, tt.want)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "пустой слайс", data: []int{}, want: 0},
		{name: "меньше элементов, чем частей", data: []int{3, 1, 4}, want: 4},
		{name: "ровно CHUNKS элементов", data: []int{1, 2, 3, 4, 5, 6, 7, 8}, want: 8},
		{name: "максимум в остатке", data: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 99}, want: 99},
		{name: "максимум в первой части", data: []int{50, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, want: 50},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxChunks(tt.data); got != tt.want {
				t.Errorf("maxChunks(%v) = %d, ожидалось %d", tt.data, got, tt.want)
			}
		})
	}

	t.Run("совпадает с maximum на случайных данных", func(t *testing.T) {
		data := generateRandomElements(1_000_003)
		if got, want := maxChunks(data), maximum(data); got != want {
			t.Errorf("maxChunks = %d, maximum = %d", got, want)
		}
	})
}
