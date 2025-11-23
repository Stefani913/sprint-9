package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Пишите тесты в этом файле
func TestGenerateRandomElementsBySize(t *testing.T) {
	size := 5
	slice := generateRandomElements(size)

	assert.Equal(t, size, len(slice))
	require.NotEmpty(t, slice)
}

func TestNotGenerateRandomElementsBySize(t *testing.T) {
	size := 0
	slice := generateRandomElements(size)

	assert.Equal(t, size, len(slice))
	require.Empty(t, slice)
}

func TestReturnMax(t *testing.T) {
	var data = []int{1, 2, 3, 4, 5}
	result := maximum(data)

	assert.Equal(t, 5, result)
}

func TestReturnMaxSingleElement(t *testing.T) {
	var data = []int{1}
	result := maximum(data)

	assert.Equal(t, 1, result)
}

func TestMaxForEmptySlice(t *testing.T) {
	data := make([]int, 0)
	result := maximum(data)

	assert.Equal(t, 0, result)
}
