package kindergarten

import (
	"errors"
	"sort"
	"strings"
	"unicode"
)

var defaultChildren = []string{
	"Alice", "Bob", "Charlie", "David",
	"Eve", "Fred", "Ginny", "Harriet",
	"Ileana", "Joseph", "Kincaid", "Larry",
}

var plants = map[rune]string{
	'G': "grass",
	'C': "clover",
	'R': "radishes",
	'V': "violets",
}

type Garden struct {
	row1     []rune
	row2     []rune
	students []string
}

// The diagram argument starts each row with a '\n'.  This allows Go's
// raw string literals to present diagrams in source code nicely as two
// rows flush left, for example,
//
//     diagram := `
//     VVCCGG
//     VVCCGG`

func NewGarden(diagram string, children []string) (*Garden, error) {
	if children == nil || len(children) == 0 {
		children = defaultChildren
	}
	if hasDuplicates(children) {
		return nil, errors.New("duplicate name")
	}
	sorted := make([]string, len(children))
	copy(sorted, children)
	sort.Strings(sorted)
	if strings.Count(diagram, "\n") != 2 {
		return nil, errors.New("wrong diagram format")
	}
	arr := strings.Split(strings.TrimSpace(diagram), "\n")
	if len(arr) < 2 {
		return nil, errors.New("invalid diagram")
	}

	rune1 := []rune(arr[0])
	rune2 := []rune(arr[1])
	for i := 0; i < 2; i++ {
		if unicode.IsLower(rune1[i]) || unicode.IsLower(rune2[i]) {
			return nil, errors.New("invalid cup codes")
		}
	}
	if len(rune1) != len(rune2) || len(rune1)%2 != 0 || len(rune2)%2 != 0 {
		return nil, errors.New("mismatched rows")
	}
	newGarden := Garden{
		row1:     []rune(arr[0]),
		row2:     []rune(arr[1]),
		students: sorted,
	}
	return &newGarden, nil
}

func (g *Garden) Plants(child string) ([]string, bool) {
	var result []string
	isFound := false
	for i, childName := range g.students {
		if childName == child {
			plant1 := []string{plants[g.row1[i*2]], plants[g.row1[i*2+1]]}
			plant2 := []string{plants[g.row2[i*2]], plants[g.row2[i*2+1]]}
			result = append(result, plant1...)
			result = append(result, plant2...)
			isFound = true
			break

		}
	}
	if isFound {
		return result, isFound
	}
	return nil, isFound
}

func hasDuplicates(items []string) bool {
	seen := make(map[string]bool)
	for _, item := range items {
		if seen[item] {
			return true // Duplicate found
		}
		seen[item] = true
	}
	return false // No duplicates
}
