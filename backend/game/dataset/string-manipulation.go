// Package dataset a handful of algorithms to work with strings, io
package dataset

import (
	"fmt"
	"regexp"
	"slices"
	"strings"
)

// StringShift return shift any given value at provided index
func StringShift(valueToShift string, from string, atIndex int) string {
	str := strings.Split(from, "")
	str = Shift(valueToShift, str, atIndex)
	from = strings.Join(str, "")
	return from
}

// AllErase return erases all the elements in the given string
func AllErase(str []string) []string {
	str = append(str[:0], str[:0]...)
	return str
}

// AfterEraseFrom return erases all the element after given limit
// example: input: 1,2,3,4
// output: limit[2]: result: 1,2
// another one: output: limit[1]:result: 1
func AfterEraseFrom(str []string, from int) []string {
	if from == len(str)-1 {
		fmt.Println(true)
		return str
	}
	str = append(str[0:from], str[from:from]...)
	return str
}

// EraseDuplicate return if found; erases the duplicate value and sorts the array
// example:
// input: 1,2,1,3,2
// output: 1,2,3
// NOTE this works for int too but i am unable to use generics
func EraseDuplicate(str []string) []string {
	if len(str) == 0 {
		return str
	}
	// sort the array
	slices.Sort(str)
	i := 0
	for j := 1; j < len(str); j++ {
		if str[j] != str[i] {
			i++
			str[i] = str[j]
		}
	}
	return str[:i+1]
}

// EraseBefore return erases all the elements before index
// example: input: 1,2,3,4
// output: index=1: 2,3,4
func EraseBefore(str []string, before int) []string {
	str = append(str[before:], str[before:before]...)
	return str
}

// EraseOnPos return erases the element at the given index position
// input: 1,2,3
// output: at pos[0]: 2,3
// at post[1]: 1,3
func EraseOnPos[T any](s []T, pos int) []T {
	// input:1 2 3 4 5
	// output: limit[2]: 1 2 4 5
	//  check if the position is within the valid index range [0, len(s)-1]
	if pos < 0 || pos >= len(s) {
		// Position is out of bounds, return the slice unchanged
		// Or, for stricter behavior, you could panic/return an error
		return s
	}

	// this creates a new slice by appending the elements before 'pos'
	// with the elements starting from 'pos + 1'
	s = append(s[:pos], s[pos+1:]...)

	return s
}

// EraseAfter return erases all the elements after given limit while keeping the element at given limit
// example: input: 1,2,3,4
// output: limit[2]: 123
func EraseAfter(str []string, before int) []string {
	str = append(str[:before], str[before:before]...)
	return str
}

// EraseLimit return erases the elements in the given limit
// example: input: 1,2,3,4
// ouput: given limit [0,3]: result: 4
func EraseLimit(str []string, from int, to int) []string {
	if from > to {
		panic("from, to belongs to [0,...size of str] where from<to ")
	}
	str = append(str[to:], str[:from]...)
	return str
}

// Shift return shift any given value at provided index
// NOTE: you can use it for any given array type
func Shift[T any](valueToShift T, from []T, atIndex int) []T {
	temp := []T{}

	//////normal swapping process////
	temp = append(temp, from...)

	//// prepends the value at given index
	from = append(from[:atIndex], valueToShift)

	//// appends back the value at given index after the appended value
	from = append(from, temp[len(from)-1:]...)

	//////end of swapping procss////
	return from
}

// Pattern return [][]string of between the limits (opne, close]
func Pattern(str string, _open string, close_ string) [][]string {
	pattern := regexp.QuoteMeta(_open) + `(.//?)` + regexp.QuoteMeta(close_)
	re, err := regexp.Compile(pattern)
	if err != nil {
		panic(err)
	}
	s := re.FindAllStringSubmatch(str, -1)
	return s
}

// StringToByte conversion
func StringToByte(from []string, to []byte) {
	for r := range from {
		to = append(to, []byte(from[r])...)
	}
}

// GetIndex return index of found value
func GetIndex(str []string, search string) int {
	found := -1
	EqualizeString_(str)
	search = strings.ToLower(search)

	for r := range str {
		if str[r] == search {
			found = r
		}
	}
	if found > 0 {
		return found
	} else {
		return found - 2
	}
}

// Replace return replaces the value at given string
func Replace(str []string, search string, replace string) {
	// if not found it will not do anything
	found := -1
	for r := range str {
		if str[r] == search {
			found = r
			if found >= 0 {
				str[found] = replace
			}
		}
	}
}

// LastIndex return last index of the string[]
func LastIndex(str []string) int {
	lasti, err := -1, 2
	for r := range str {
		lasti = r
	}
	if lasti == -1 {
		panic(err)
	}
	return lasti
}

// SecondLastIndex return 2nd last index
func SecondLastIndex(str []string, search string) int {
	inx := []int{}
	EqualizeString_(str)
	search = strings.ToLower(search)

	for r := range str {
		if str[r] == search {
			inx = append(inx, r)
		}
	}
	f := -1
	if len(inx) > 1 {
		f = inx[len(inx)-2]
	} else if len(inx) == 1 {
		f = inx[0]
	} else {
		f = 0
	}
	return f
}

// ElementRepeated return: number of times the element repeated
// example: [1,1,2,3] res=1
func ElementRepeated(str []string, search string) int {
	inx := []int{}
	repeated := 0
	EqualizeString_(str)
	search = strings.ToLower(search)
	for r := range str {
		if str[r] == search {
			inx = append(inx, r)
			repeated = len(inx) - 1
		}
	}
	return repeated

}

// Includes return true if the value is in the string
func Includes(str []string, search string) bool {
	has := false
	EqualizeString_(str)
	search = strings.ToLower(search)
	for r := range str {
		if str[r] == search {
			has = true
		}
	}
	return has
}

// GetLastRepeationIndex return last index of repeated value in a string
func GetLastRepeationIndex(str []string, search string) int {
	found := []int{}
	foundinx := -1
	EqualizeString_(str)
	search = strings.ToLower(search)
	for r := range str {
		if str[r] == search {
			found = append(found, r)
		}
	}
	if len(found) != 0 {
		foundinx = found[len(found)-1]
	}
	return foundinx
}

// EqualizeString_ return equals the [] string not matter if it has first letter capital or last letter capital
func EqualizeString_(str []string) {
	for r := range str {
		var temp = strings.ToLower(str[r])
		str[r] = temp
	}
}
