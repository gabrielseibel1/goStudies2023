package cp3pg12

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
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

func BinarySearch() {
	fmt.Print("How many sorted integers to search on? ")
	var n int
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		panic(err)
	}

	fmt.Println("Give me the sorted integers to search on (sep by line break):")
	s := make([]int, n)
	for i := 0; i < n; i++ {
		_, err := fmt.Scanf("%d", &s[i])
		if err != nil {
			panic(err)
		}
	}

	fmt.Print("What integer do you want me to search? ")
	var x int
	_, err = fmt.Scanf("%d", &x)
	if err != nil {
		panic(err)
	}

	i, j := 0, n-1

	for i <= j {
		mean := (i + j) / 2

		if s[mean] == x {
			fmt.Printf("Found %d at position %d\n", x, mean)
			return
		}
		if x < s[mean] {
			j = mean - 1
		} else {
			i = mean + 1
		}
	}

	fmt.Printf("Didn't find %d\n", x)
}

// 7. Generate all possible permutations of {‘A’, ‘B’, ‘C’, . . . , ‘J’}, the first N = 10 letters
// in the alphabet (see Section 3.2.1).

func PermutationsOfAToJ() {
	letters := []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J'}

	permutations := make([][]byte, 0, factorial(len(letters)))

	wg := new(sync.WaitGroup)
	wg.Add(1)
	addPermutations(letters, &permutations, []byte{}, wg)
	wg.Wait()

	fmt.Println(permutations)
}

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func addPermutations(letters []byte, permutations *[][]byte, permutation []byte, fwg *sync.WaitGroup) {
	defer fwg.Done()

	if len(letters) == 0 {
		pc := make([]byte, len(permutation))
		copy(pc, permutation)
		*permutations = append(*permutations, pc)
		return
	}

	wg := new(sync.WaitGroup)
	wg.Add(len(letters))
	for _, l := range letters {
		// copy letters without current letter
		lessLetters := make([]byte, 0, len(letters)-1)
		for _, v := range letters {
			if v != l {
				lessLetters = append(lessLetters, v)
			}
		}

		pc := make([]byte, len(permutation), len(permutation)+1)
		copy(pc, permutation)
		pc = append(pc, l)
		go addPermutations(lessLetters, permutations, pc, wg)
	}
	wg.Wait()
}

// 8. Generate all possible subsets of {0, 1, 2, . . . , N-1}, for N = 20 (see Section 3.2.1).

func PowerSet0to20() {
	set := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	pSetInts := powerSetInts(len(set))

	powerSet := intsToCombinations(set, pSetInts)

	fmt.Println(powerSet)
}

func powerSetInts(n int) []int {
	states := make([]int, 0, int(math.Pow(2, float64(n))))
	for i := 0; i <= n; i++ {
		intsWithBitsOn(n, &states, 0, i, 0)
	}
	return states
}

func intsWithBitsOn(n int, states *[]int, state, mustOn, at int) {
	if mustOn == 0 {
		*states = append(*states, state)
		return
	}

	for i := at; i < n; i++ {
		intsWithBitsOn(n, states, withOn(state, i), mustOn-1, i+1)
	}
}

func isOn(state, i int) bool {
	return (state & (1 << i)) == (1 << i)
}

func withOn(state, i int) int {
	return (state | (1 << i))
}

func intsToCombinations(set, ints []int) [][]int {
	powerSet := make([][]int, 0, len(ints))
	for _, state := range ints {
		comb := make([]int, 0)
		for i := 0; i <= len(set); i++ {
			if isOn(state, i) {
				comb = append(comb, set[i])
			}
		}
		combCopy := make([]int, len(comb))
		copy(combCopy, comb) // needed?
		powerSet = append(powerSet, combCopy)
	}
	return powerSet
}

// 9. Given a string that represents a base X number, convert it to an equivalent string in
// base Y, 2 ≤ X, Y ≤ 36. For example: “FF” in base X = 16 (hexadecimal) is “255” in
// base Y1 = 10 (decimal) and “11111111” in base Y2 = 2 (binary). See Section 5.3.2.

func BaseChange() {
	fmt.Print("What is the base of your number? ")
	var x int
	_, err := fmt.Scanf("%d", &x)
	if err != nil {
		panic(err)
	}

	fmt.Printf("What is your number in base %d? ", x)
	var sx string
	_, err = fmt.Scanf("%s", &sx)
	if err != nil {
		panic(err)
	}

	i, err := strconv.ParseInt(sx, x, 64)
	if err != nil {
		panic(fmt.Errorf("unable to parse %s as base %d number", sx, x))
	}

	fmt.Printf("What is the base you desire to convert %s to? ", sx)
	var y int
	_, err = fmt.Scanf("%d", &y)
	if err != nil {
		panic(err)
	}

	sy := strconv.FormatInt(i, y)
	fmt.Printf("In base %d, your number is %s\n", y, sy)
}

// 10. Let’s define a ‘special word’ as a lowercase alphabet followed by two consecutive digits.
// Given a string, replace all ‘special words’ of length 3 with 3 stars “***”, e.g.
// S = “line: a70 and z72 will be replaced, aa24 and a872 will not”
// should be transformed into:
// S = “line: *** and *** will be replaced, aa24 and a872 will not”.

func SpecialWords() {
	fmt.Println("Give me a string, I will change [a-Z][0-9][0-9] words to ***!")
	scanner := bufio.NewScanner(os.Stdin)

	re := regexp.MustCompile(`^[A-Za-z][0-9][0-9]$`)
	var redacted string

	for scanner.Scan() {
		words := strings.Split(scanner.Text(), " ")
		for i := 0; i < len(words); i++ {
			if re.MatchString(words[i]) {
				redacted += "***"
			} else {
				redacted += words[i]
			}
			if i+1 < len(words) {
				redacted += " "
			}
		}
		redacted += "\n"
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println("Your redacted content:")
	fmt.Print(redacted)
}
