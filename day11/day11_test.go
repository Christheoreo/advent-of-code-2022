package day11

import (
	"testing"
)

var testFile string = "problem-11-example.txt"
var realFile string = "problem-11.txt"

// func TestSolveFirst(t *testing.T) {

// 	expected := 10605
// 	actual := SolveFirst(testFile)

// 	if expected != actual {
// 		t.Errorf("Expected %d but got %d", expected, actual)
// 		t.Fail()
// 	}
// 	answer := SolveFirst(realFile)
// 	fmt.Printf("\n-----\n --ANSWER Day 11 part 1 = %d \n-----\n", answer)
// }

func TestSolveSecond(t *testing.T) {

	expected := 2713310158
	actual := SolveSecond(testFile)

	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
		t.Fail()
	}

	// answer := SolveSecond(realFile)
	// fmt.Printf("\n-----\n --ANSWER Day 11 part 2 = %d \n-----\n", answer)
}

func TestStringAddition(t *testing.T) {
	expected := "1010478"
	strOne := "454689"
	strTwo := "555789"

	actual := addStrings(strOne, strTwo)

	if expected != actual {
		t.Errorf("Expected %s but got %s", expected, actual)
		t.Fail()
	}
}

func TestStringAdditionA(t *testing.T) {
	expected := "601257"
	strOne := "45468"
	strTwo := "555789"

	actual := addStrings(strOne, strTwo)

	if expected != actual {
		t.Errorf("Expected %s but got %s", expected, actual)
		t.Fail()
	}
}

func TestStringMultiply(t *testing.T) {
	expected := "1480"
	strOne := "40"
	strTwo := "37"

	actual := multiplyStrings(strOne, strTwo)

	if expected != actual {
		t.Errorf("Expected %s but got %s", expected, actual)
		t.Fail()
	}
}

func TestStringMultiplyA(t *testing.T) {
	expected := "8604460"
	strOne := "436"
	strTwo := "19735"

	actual := multiplyStrings(strOne, strTwo)

	if expected != actual {
		t.Errorf("Expected %s but got %s", expected, actual)
		t.Fail()
	}
}

func TestSubtractStrings(t *testing.T) {
	expected := "6"
	strOne := "9"
	stringTwo := "3"

	actual := subTractStrings(strOne, stringTwo)

	if expected != actual {
		t.Errorf("Expected %s but got %s", expected, actual)
		t.Fail()
	}
}

func TestSubtractStringsA(t *testing.T) {
	expected := "120"
	strOne := "200"
	stringTwo := "80"

	actual := subTractStrings(strOne, stringTwo)

	if expected != actual {
		t.Errorf("Expected %s but got %s", expected, actual)
		t.Fail()
	}
}

func TestSubtractStringsB(t *testing.T) {
	expected := "1"
	strOne := "11"
	stringTwo := "10"

	actual := subTractStrings(strOne, stringTwo)

	if expected != actual {
		t.Errorf("Expected %s but got %s", expected, actual)
		t.Fail()
	}
}

func TestSubtractStringsC(t *testing.T) {
	expected := "11238746823648264386332"
	strOne := "11238746823648264386432"
	stringTwo := "100"

	actual := subTractStrings(strOne, stringTwo)

	if expected != actual {
		t.Errorf("Expected %s but got %s", expected, actual)
		t.Fail()
	}
}

func TestStringIsDivisable(t *testing.T) {
	expected := true
	strOne := "9"
	by := "3"

	actual := isDivisible(strOne, by)

	if expected != actual {
		t.Errorf("Expected %v but got %v", expected, actual)
		t.Fail()
	}
}

func TestStringIsDivisableA(t *testing.T) {
	expected := false
	strOne := "22"
	by := "3"

	actual := isDivisible(strOne, by)

	if expected != actual {
		t.Errorf("Expected %v but got %v", expected, actual)
		t.Fail()
	}
}

func TestStringIsDivisableB(t *testing.T) {
	expected := true
	strOne := "69"
	by := "3"

	actual := isDivisible(strOne, by)

	if expected != actual {
		t.Errorf("Expected %v but got %v", expected, actual)
		t.Fail()
	}
}

func TestStringIsDivisableC(t *testing.T) {
	expected := true
	strOne := "69"
	by := "3"

	actual := isDivisible(strOne, by)

	if expected != actual {
		t.Errorf("Expected %v but got %v", expected, actual)
		t.Fail()
	}
}

func TestStringIsDivisableD(t *testing.T) {
	expected := false
	strOne := "6900000000000"
	by := "107"

	actual := isDivisible(strOne, by)

	if expected != actual {
		t.Errorf("Expected %v but got %v", expected, actual)
		t.Fail()
	}
}

func TestStringIsDivisableE(t *testing.T) {
	expected := false
	strOne := "300000000000002"
	by := "3"

	actual := isDivisible(strOne, by)

	if expected != actual {
		t.Errorf("Expected %v but got %v", expected, actual)
		t.Fail()
	}
}

func TestStringIsDivisableF(t *testing.T) {
	expected := true
	strOne := "300000000000003"
	by := "3"

	actual := isDivisible(strOne, by)

	if expected != actual {
		t.Errorf("Expected %v but got %v", expected, actual)
		t.Fail()
	}
}

// func BenchmarkSolveFirst(b *testing.B) {
// 	for n := 0; n < b.N; n++ {
// 		SolveFirst(realFile)
// 	}
// }

// func BenchmarkSolveSecond(b *testing.B) {
// 	for n := 0; n < b.N; n++ {
// 		SolveSecond(realFile)
// 	}
// }

func BenchmarkAddition(b *testing.B) {
	const strOne = "24726487236487236487236487236487"
	const strTwo = "2472648723648723648723648723648712321312"
	for n := 0; n < b.N; n++ {
		addStrings(strOne, strTwo)
	}
}

func BenchmarkMultiplication(b *testing.B) {
	const strOne = "24726487236487236487236487236487"
	const strTwo = "2472648723648723648723648723648712321312"
	for n := 0; n < b.N; n++ {
		multiplyStrings(strOne, strTwo)
	}
}

func BenchmarkSubTraction(b *testing.B) {
	const strOne = "24726487236487236487236487236487"
	const strTwo = "24726487236487236487236487236"
	for n := 0; n < b.N; n++ {
		subTractStrings(strOne, strTwo)
	}
}

func BenchmarkDivisable(b *testing.B) {
	const strOne = "24726487236487236487236487236487"
	const strTwo = "12312"
	for n := 0; n < b.N; n++ {
		isDivisible(strOne, strTwo)
	}
}
