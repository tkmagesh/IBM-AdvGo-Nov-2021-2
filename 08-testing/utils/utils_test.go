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
		TestCase{Name: "Is 0 Prime", Input: 0, ExpectedResult: false},
		TestCase{Name: "Is 2 Prime", Input: 2, ExpectedResult: true},
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

func TestAllPrimeNos_2(t *testing.T) {
	testCases := []TestCase{
		TestCase{Name: "Is 0 Prime(2)", Input: 0, ExpectedResult: false},
		TestCase{Name: "Is 2 Prime(2)", Input: 2, ExpectedResult: true},
		TestCase{Name: "Is 7 Prime(2)", Input: 7, ExpectedResult: true},
		TestCase{Name: "Is 9 Prime(2)", Input: 9, ExpectedResult: false},
		TestCase{Name: "Is 11 Prime(2)", Input: 11, ExpectedResult: true},
		TestCase{Name: "Is 12 Prime(2)", Input: 12, ExpectedResult: false},
		TestCase{Name: "Is 13 Prime(2)", Input: 13, ExpectedResult: true},
		TestCase{Name: "Is 17 Prime(2)", Input: 17, ExpectedResult: true},
		TestCase{Name: "Is 19 Prime(2)", Input: 19, ExpectedResult: true},
		TestCase{Name: "Is 23 Prime(2)", Input: 23, ExpectedResult: true},
	}
	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			testCase.ActualResult = IsPrime2(testCase.Input)
			if testCase.ActualResult != testCase.ExpectedResult {
				t.Errorf("Expected %v but got %v", testCase.ExpectedResult, testCase.ActualResult)
			}
		})
	}
}

func TestAllPrimeNos_3(t *testing.T) {
	testCases := []TestCase{
		TestCase{Name: "Is 0 Prime(3)", Input: 0, ExpectedResult: false},
		TestCase{Name: "Is 2 Prime(3)", Input: 2, ExpectedResult: true},
		TestCase{Name: "Is 7 Prime(3)", Input: 7, ExpectedResult: true},
		TestCase{Name: "Is 9 Prime(3)", Input: 9, ExpectedResult: false},
		TestCase{Name: "Is 11 Prime(3)", Input: 11, ExpectedResult: true},
		TestCase{Name: "Is 12 Prime(3)", Input: 12, ExpectedResult: false},
		TestCase{Name: "Is 13 Prime(3)", Input: 13, ExpectedResult: true},
		TestCase{Name: "Is 17 Prime(3)", Input: 17, ExpectedResult: true},
		TestCase{Name: "Is 19 Prime(3)", Input: 19, ExpectedResult: true},
		TestCase{Name: "Is 23 Prime(3)", Input: 23, ExpectedResult: true},
	}
	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			testCase.ActualResult = IsPrime3(testCase.Input)
			if testCase.ActualResult != testCase.ExpectedResult {
				t.Errorf("Expected %v but got %v", testCase.ExpectedResult, testCase.ActualResult)
			}
		})
	}
}

func BenchmarkIsPrime_1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPrime(79)
	}
}

func BenchmarkIsPrime_2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPrime2(79)
	}
}

func BenchmarkIsPrime_3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPrime3(79)
	}
}
