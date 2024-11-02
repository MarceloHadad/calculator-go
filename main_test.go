package main

import (
	"math"
	"testing"
)

func almostEqual(n1, n2, tolerance float64) bool {
	return math.Abs(n1-n2) <= tolerance
}

func TestShouldSumCorrectlyWithIntPositiveNumbers(t *testing.T) {
	expectedResult := 13.0

	result := sum(8, 5)

	if result != expectedResult {
		t.Errorf("Expected %f, got %f", expectedResult, result)
	}
}

func TestShouldSumCorrectlyWithFloatPositiveNumbers(t *testing.T) {
	expectedResult := 13.5

	result := sum(8.2, 5.3)

	if result != expectedResult {
		t.Errorf("Expected %f, got %f", expectedResult, result)
	}
}

func TestShouldSumCorrectlyWithIntNegativeNumbers(t *testing.T) {
	expectedResult := -3.0

	result := sum(-8, 5)

	if result != expectedResult {
		t.Errorf("Expected %f, got %f", expectedResult, result)
	}
}

func TestShouldSumCorrectlyWithFloatNegativeNumbers(t *testing.T) {
	expectedResult := -2.9
	tolerance := 0.0000001

	result := sum(-8.2, 5.3)

	if !almostEqual(result, expectedResult, tolerance) {
		t.Errorf("Expected %f, got %f", expectedResult, result)
	}
}

func TestShouldSumCorrectlyWithZero(t *testing.T) {
	expectedResult := 5.0

	result := sum(0, 5)

	if result != expectedResult {
		t.Errorf("Expected %f, got %f", expectedResult, result)
	}
}

func TestShouldSubCorrectlyWithIntPositiveNumbers(t *testing.T) {
	expectedResult := 2.0

	result := sub(5, 3)

	if result != expectedResult {
		t.Errorf("Expected %f, got %f", expectedResult, result)
	}
}

func TestShouldSubCorrectlyWithFloatPositiveNumbers(t *testing.T) {
	expectedResult := 2.0

	result := sub(5.3, 3.3)

	if result != expectedResult {
		t.Errorf("Expected %f, got %f", expectedResult, result)
	}
}

func TestShouldSubCorrectlyWithIntNegativeNumbers(t *testing.T) {
	expectedResult := -8.0

	result := sub(-5, 3)

	if result != expectedResult {
		t.Errorf("Expected %f, got %f", expectedResult, result)
	}
}

func TestShouldSubCorrectlyWithFloatNegativeNumbers(t *testing.T) {
	expectedResult := -8.6

	result := sub(-5.3, 3.3)

	if result != expectedResult {
		t.Errorf("Expected %f, got %f", expectedResult, result)
	}
}

func TestShouldSubCorrectlyWithZero(t *testing.T) {
	expectedResult := -5.0

	result := sub(0, 5)

	if result != expectedResult {
		t.Errorf("Expected %f, got %f", expectedResult, result)
	}
}

func TestShouldMulCorrectlyWithIntPositiveNumbers(t *testing.T) {
	expectedResult := 15.0

	result := mul(5, 3)

	if result != expectedResult {
		t.Errorf("Expected %f, got %f", expectedResult, result)
	}
}

func TestShouldMulCorrectlyWithFloatPositiveNumbers(t *testing.T) {
	expectedResult := 15.0

	result := mul(5.0, 3.0)

	if result != expectedResult {
		t.Errorf("Expected %f, got %f", expectedResult, result)
	}
}

func TestShouldMulCorrectlyWithIIntNegativeNumbers(t *testing.T) {
	expectedResult := -15.0

	result := mul(-5, 3)

	if result != expectedResult {
		t.Errorf("Expected %f, got %f", expectedResult, result)
	}
}

func TestShouldMulCorrectlyWithFloatNegativeNumbers(t *testing.T) {
	expectedResult := -15.0

	result := mul(-5.0, 3.0)

	if result != expectedResult {
		t.Errorf("Expected %f, got %f", expectedResult, result)
	}
}

func TestShouldMulCorrectlyWithZero(t *testing.T) {
	expectedResult := 0.0

	result := mul(0, 5)

	if result != expectedResult {
		t.Errorf("Expected %f, got %f", expectedResult, result)
	}
}

func TestShouldDivCorrectlyWithIntPositiveNumbers(t *testing.T) {
	expectedResult := 4.0

	result, err := div(8, 2)

	if err != nil {
		t.Errorf("Expected error nil, got %v", err)
	}
	if result != expectedResult {
		t.Errorf("Expected %f, got %f", expectedResult, result)
	}
}

func TestShouldDivCorrectlyWithFloatPositiveNumbers(t *testing.T) {
	expectedResult := 4.0

	result, err := div(8.0, 2.0)

	if err != nil {
		t.Errorf("Expected error nil, got %v", err)
	}
	if result != expectedResult {
		t.Errorf("Expected %f, got %f", expectedResult, result)
	}
}

func TestShouldDivCorrectlyWithIntNegativeNumbers(t *testing.T) {
	expectedResult := -4.0

	result, err := div(-8, 2)

	if err != nil {
		t.Errorf("Expected error nil, got %v", err)
	}
	if result != expectedResult {
		t.Errorf("Expected %f, got %f", expectedResult, result)
	}
}

func TestShouldDivCorrectlyWithFloatNegativeNumbers(t *testing.T) {
	expectedResult := -4.0

	result, err := div(-8.0, 2.0)

	if err != nil {
		t.Errorf("Expected error nil, got %v", err)
	}
	if result != expectedResult {
		t.Errorf("Expected %f, got %f", expectedResult, result)
	}
}

func TestShouldReturnErrorWhenDivByZero(t *testing.T) {
	_, err := div(10, 0)

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
