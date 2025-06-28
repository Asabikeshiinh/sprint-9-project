package main

import (
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	slice := generateRandomElements(0)
	if len(slice) != 0 {
		t.Errorf("Ожидался пустой слайс, получен %v", slice)
	}
	slice = generateRandomElements(5)
	if len(slice) != 5 {
		t.Errorf("Ожидался слайс длины 5, получен слайс длиной %d", len(slice))
	}
}

func TestMaximum(t *testing.T) {
	if maximum([]int{}) != 0 {
		t.Error("Для пустого слайса функция должна возвращать 0")
	}
	if maximum([]int{42}) != 42 {
		t.Error("Для слайса, содержащего 1 число, функция должна возвращать это число")
	}
	if max := maximum([]int{1, 2, 3, 2, 1}); max != 3 {
		t.Errorf("Слайс {1, 2, 3, 2, 1}, ожидалось 3, получено %d", max)
	}
}
