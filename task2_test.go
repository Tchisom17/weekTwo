package main

import (
	"reflect"
	"testing"
)

func TestTableAdd(t *testing.T) {
	var tests = []struct {
		input []float64
		expected float64
	}{
		{
			[]float64{1,2,3},
			6,
		},
	}
	for _,test := range tests{
		if output := add(test.input); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, received: {}", test.input, test.expected, output)
		}
	}
}
func TestTableSubtract(t *testing.T) {
	var tests = []struct {
		input []float64
		expected float64
	}{
		{
			[]float64{10, 2},
			8,
		},
	}
	for _,test := range tests{
		if output := subtract(test.input); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, received: {}", test.input, test.expected, output)
		}
	}
}
func TestTableMultiply(t *testing.T) {
	var tests = []struct {
		input []float64
		expected float64
	}{
		{
			[]float64{1,2,3},
			6,
		},
	}
	for _,test := range tests{
		if output := multiply(test.input); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, received: {}", test.input, test.expected, output)
		}
	}
}
func TestTableDivide(t *testing.T) {
	var tests = []struct {
		input []float64
		expected float64
	}{
		{
			[]float64{27,3,3},
			3,
		},
	}
	for _,test := range tests{
		if output := divide(test.input); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, received: {}", test.input, test.expected, output)
		}
	}
}
func TestCalculate(t *testing.T){
	result, myAddition := []float64{10,9,8,0}, calc("5*2","27/3","4+4","2-2")
	if !reflect.DeepEqual(result,myAddition) {
		t.Errorf("incorrect, supposed to be %v", result)
	}
}