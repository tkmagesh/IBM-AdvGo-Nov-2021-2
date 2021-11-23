package utils

import "testing"

func TestIsPrime(t *testing.T) {
	//arrange
	no := 7
	expectedResult := true

	//act
	actualResult := IsPrime(no)

	//assert
	if actualResult != expectedResult {
		t.Errorf("Expected %v but got %v", expectedResult, actualResult)
	}
}

type TestCase struct {
	Name           string
	Input          int
	ExpectedResult bool
	ActualResult   bool
}

func TestAllPrimeNos(t *testing.T) {
	testCases := []TestCase{
		TestCase{Name: "Is 7 Prime", Input: 7, ExpectedResult: true},
		TestCase{Name: "Is 9 Prime", Input: 9, ExpectedResult: false},
		TestCase{Name: "Is 11 Prime", Input: 11, ExpectedResult: true},
		TestCase{Name: "Is 12 Prime", Input: 12, ExpectedResult: false},
		TestCase{Name: "Is 13 Prime", Input: 13, ExpectedResult: true},
		TestCase{Name: "Is 17 Prime", Input: 17, ExpectedResult: true},
		TestCase{Name: "Is 19 Prime", Input: 19, ExpectedResult: true},
		TestCase{Name: "Is 23 Prime", Input: 23, ExpectedResult: true},
	}
	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			testCase.ActualResult = IsPrime(testCase.Input)
			if testCase.ActualResult != testCase.ExpectedResult {
				t.Errorf("Expected %v but got %v", testCase.ExpectedResult, testCase.ActualResult)
			}
		})
	}
}
