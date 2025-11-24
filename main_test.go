package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type MaxValueTestSuite struct {
	suite.Suite
}

func TestMaxValueSuite(t *testing.T) {
	suite.Run(t, new(MaxValueTestSuite))
}

func (suite *MaxValueTestSuite) TestGenerateRandomElements() {
	tests := []struct {
		name        string
		input       int
		factIsEmpty bool
	}{
		{
			name:        "Создание массива для нескольки значений",
			input:       5,
			factIsEmpty: false,
		},
		{
			name:        "Создание массива для одного значения",
			input:       1,
			factIsEmpty: false,
		},
		{
			name:        "Создание пустого массива, если передан 0",
			input:       0,
			factIsEmpty: true,
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			slice := generateRandomElements(tt.input)
			sliceSize := len(slice)

			if tt.factIsEmpty {
				require.Empty(suite.T(), slice)

			} else {
				require.NotEmpty(suite.T(), slice)
			}

			assert.Equal(suite.T(), tt.input, sliceSize, "generateRandomElements() возвращает значение %d. Ожидаемое %d", tt.input, sliceSize)
		})
	}
}

func (suite *MaxValueTestSuite) TestFindingMaxElement() {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "Поиск максимума в слайсе из нескольких одинаковых значений",
			input:    []int{1, 1, 1, 1, 1},
			expected: 1,
		},
		{
			name:     "Поиск максимума в слайсе из нескольких разных значений",
			input:    []int{1, 2, 3, 4, 5, 13},
			expected: 13,
		},
		{
			name:     "Поиск максимума в слайсе из одного значения",
			input:    []int{9},
			expected: 9,
		},
		{
			name:     "Поиск максимума в путом слайсе",
			input:    []int{},
			expected: 0,
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			maxNumber := maximum(tt.input)

			assert.Equal(suite.T(), tt.expected, maxNumber, "maximum() возвращает значение %d. Ожидаемое %d", tt.expected, maxNumber)
		})
	}
}
func (suite *MaxValueTestSuite) TestMaxElementsInChunks() {
	tests := []struct {
		name      string
		input     []int
		expectErr bool
		expected  int
	}{
		{
			name:      "Поиск максимума в слайсе из нескольких значений, количество которых равно 8 (1 чанк)",
			input:     []int{4, 3, 0, 1, 1, 2, 1, 5},
			expectErr: false,
			expected:  5,
		},
		{
			name:      "Поиск максимума в слайсе из нескольких значений, количество которых больше 8 (2 чанка)",
			input:     []int{1, 1, 1, 1, 1, 2, 4, 1, 51, 22, 114141, 22, 1, 0, 44, 876},
			expectErr: false,
			expected:  114141,
		},
		{
			name:      "Поиск максимума в слайсе из нескольких значений, количество которых меньше 8",
			input:     []int{9, 13},
			expectErr: true,
		},
		{
			name:      "Поиск максимума в слайсе из одного значения",
			input:     []int{9},
			expectErr: true,
		},
		{
			name:      "Поиск максимума в путом слайсе",
			input:     []int{},
			expectErr: false,
			expected:  0,
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			maxNumber, err := maxChunks(tt.input)

			if tt.expectErr {
				assert.Error(suite.T(), err, "Ожидалась ошибка %q и мы ее получили", err)
			} else {
				assert.NoError(suite.T(), err, " Неожиданных ошибок не возникло: %q", err)
				assert.Equal(suite.T(), tt.expected, maxNumber, "maximum() возвращает значение %d. Ожидаемое %d", maxNumber, tt.expected)
			}
		})
	}
}
