package cp3pg12

import (
	"fmt"
	"math"
	"sort"
	"time"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/number"
)

// 1. Using Java, read in a double (e.g. 1.4732, 15.324547327, etc.), echo it,
// but with a minimum field width of 7 and 3 digits after the decimal point
// (e.g. ss1.473 (where ‘s’ denotes a space), s15.325, etc.)

func ReadDoubleEchoWidth7Precision3() {
	fmt.Print("Give me a double: ")

	var f float64
	i, err := fmt.Scanf("%f", &f)
	if err != nil || i != 1 {
		panic("Failed to read exactly one double precision number (float64)")
	}

	fmt.Printf("%7.3f\n", f)
}

// 2. Given an integer n (n ≤ 15), print π to n digits after the decimal point (rounded).
// (e.g. for n = 2, print 3.14; for n = 4, print 3.1416; for n = 5, prints 3.14159.)

func NPiDigits() {
	fmt.Print("How many π digits after decimal point? ")

	var n int
	i, err := fmt.Scanf("%d", &n)
	if err != nil || i != 1 {
		panic("Failed to read exactly one integer (int32)")
	}
	fmt.Printf("%d digits: ", n)

	pi := number.Decimal(math.Pi, number.MaxFractionDigits(n))
	f := message.NewPrinter(language.English)
	f.Println(pi)
}

// 3. Given a date, determine the day of the week (Monday, . . . , Sunday) on that day.
// (e.g. 9 August 2010—the launch date of the first edition of this book—is a Monday.)

func DayOfTheWeek() {
	fmt.Print("Give me a date (as DD/MM/YYYY): ")

	var input string
	n, err := fmt.Scanf("%s", &input)
	if err != nil || n != 1 || input == "" {
		panic(err)
	}

	loc, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		panic(err)
	}

	t, err := time.ParseInLocation("02/01/2006", input, loc)
	if err != nil {
		panic(err)
	}

	fmt.Printf("It's a %s\n", t.Weekday())
}

// 4. Given n random integers, print the distinct (unique) integers in sorted order.

func SortNRandomInts() {
	fmt.Print("How many integers to sort? ")

	var n int
	i, err := fmt.Scanf("%d", &n)
	if err != nil || i != 1 {
		panic("Failed to read exactly one integer (int32)")
	}

	fmt.Printf("The %d integers to sort (sep by line break):\n", n)

	m := make(map[int]struct{})
	for i := 0; i < n; i++ {
		var a int
		i, err := fmt.Scanf("%d\n", &a)
		if err != nil || i != 1 {
			panic(fmt.Errorf("failed to read exactly one integer (int32) at position %d", i))
		}
		m[a] = struct{}{}
	}

	s := make(sort.IntSlice, 0, len(m))
	for k := range m {
		s = append(s, k)
	}

	s.Sort()

	fmt.Println(s)
}

// 5. Given the distinct and valid birthdates of n people as triples (DD, MM, YYYY), order
// them first by ascending birth months (MM), then by ascending birth dates (DD), and
// finally by ascending age.

func SortBirthdatesMonthDateAge() {
	fmt.Print("How many birthdates? ")

	var n int
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		panic(err)
	}

	loc, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		panic(err)
	}

	ts := make(TimesSorter, 0, n)
	for i := 0; i < n; i++ {
		var d, m, y int
		_, err := fmt.Scanf("(%02d, %02d, %04d)\n", &d, &m, &y)
		if err != nil {
			panic(err)
		}

		s := fmt.Sprintf("%02d/%02d/%04d", d, m, y)
		t, err := time.ParseInLocation("02/01/2006", s, loc)
		if err != nil {
			panic(err)
		}

		ts = append(ts, t)
	}

	sort.Sort(&ts)

	fmt.Println(ts)
}

type TimesSorter []time.Time

func (ts *TimesSorter) Len() int {
	return len(*ts)
}

func (ts *TimesSorter) Swap(i, j int) {
	t := (*ts)[i]
	(*ts)[i] = (*ts)[j]
	(*ts)[j] = t
}

func (ts *TimesSorter) Less(i, j int) bool {
	ith := (*ts)[i]
	jth := (*ts)[j]

	if ith.Month() < jth.Month() {
		return true
	} else if ith.Month() > jth.Month() {
		return false
	}

	if ith.Day() < jth.Day() {
		return true
	} else if ith.Day() > jth.Day() {
		return false
	}

	if time.Since(ith) < time.Since(jth) {
		return true
	} else if time.Since(ith) > time.Since(jth) {
		return false
	}

	return false
}

// 6. Given a list of sorted integers L of size up to 1M items, determine whether a value v
// exists in L with no more than 20 comparisons (more details in Section 2.2).

// 7. Generate all possible permutations of {‘A’, ‘B’, ‘C’, . . . , ‘J’}, the first N = 10 letters
// in the alphabet (see Section 3.2.1).

// 8. Generate all possible subsets of {0, 1, 2, . . . , N-1}, for N = 20 (see Section 3.2.1).

// 9. Given a string that represents a base X number, convert it to an equivalent string in
// base Y, 2 ≤ X, Y ≤ 36. For example: “FF” in base X = 16 (hexadecimal) is “255” in
// base Y1 = 10 (decimal) and “11111111” in base Y2 = 2 (binary). See Section 5.3.2.

// 10. Let’s define a ‘special word’ as a lowercase alphabet followed by two consecutive digits.
// Given a string, replace all ‘special words’ of length 3 with 3 stars “***”, e.g.
// S = “line: a70 and z72 will be replaced, aa24 and a872 will not”
// should be transformed into:
// S = “line: *** and *** will be replaced, aa24 and a872 will not”.

// 11. Given a valid mathematical expression involving ‘+’, ‘-’, ‘*’, ‘/’, ‘(’, and ‘)’ in a single
// line, evaluate that expression. (e.g. a rather complicated but valid expression 3 + (8 -
// 7.5) * 10 / 5 - (2 + 5 * 7) should produce -33.0 when evaluated with standard
// operator precedence.)
