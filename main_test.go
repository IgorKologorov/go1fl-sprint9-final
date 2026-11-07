package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	data := generateRandomElements(100)

	if len(data) != 100 {
		t.Fatalf("ожидалась длина 100, получено %d", len(data))
	}

	for _, value := range data {
		if value <= 0 {
			t.Fatalf("ожидалось положительное число, получено %d", value)
		}
	}

	empty := generateRandomElements(0)

	if len(empty) != 0 {
		t.Fatalf(
			"для размера 0 ожидался пустой слайс, получена длина %d",
			len(empty),
		)
	}

	negative := generateRandomElements(-1)

	if len(negative) != 0 {
		t.Fatalf(
			"для отрицательного размера ожидался пустой слайс, получена длина %d",
			len(negative),
		)
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{
			name: "обычный слайс",
			data: []int{3, 7, 2, 9, 4},
			want: 9,
		},
		{
			name: "один элемент",
			data: []int{42},
			want: 42,
		},
		{
			name: "пустой слайс",
			data: []int{},
			want: 0,
		},
		{
			name: "отрицательные значения",
			data: []int{-8, -2, -15},
			want: -2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximum(tt.data)

			if got != tt.want {
				t.Errorf(
					"maximum() = %d, ожидалось %d",
					got,
					tt.want,
				)
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
		{
			name: "делится на восемь",
			data: []int{1, 2, 3, 4, 5, 6, 7, 8},
			want: 8,
		},
		{
			name: "есть остаток",
			data: []int{1, 10, 3, 4, 5, 6, 7, 8, 9, 2},
			want: 10,
		},
		{
			name: "меньше восьми элементов",
			data: []int{4, 1, 7},
			want: 7,
		},
		{
			name: "пустой слайс",
			data: []int{},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxChunks(tt.data)

			if got != tt.want {
				t.Errorf(
					"maxChunks() = %d, ожидалось %d",
					got,
					tt.want,
				)
			}
		})
	}
}
