package main

import (
	"testing"
)

// bottlenecks are performance blockers 
// problems that make your application slower
// 

func TestCalculate(t *testing.T){

	var tests = []struct {
		input int
		expected int
	}{
		{2,4}, 
		{-1, 1},
		{0, 2}, 
		{-5, -3}, 
		{99999, 100001}, 
	}

	for _, test := range tests{
		if output := Calculate(test.input);  output != test.expected{
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
		}
	}



}

// resources 

// writing test in go tutorial 

//https://tutorialedge.net/golang/intro-testing-in-go/