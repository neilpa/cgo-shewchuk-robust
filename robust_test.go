package robust_test

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"testing"
)

type Vec2 struct{ X, Y float64 }
type Vec3 struct{ X, Y, Z float64 }

var result float64 // benchmark results

type testcase struct {
	args  []float64 // input args to predicate function
	sign  int       // sign of the return value from predicate
	label string    // test case label (e.g. file and line number)
}

func loadCases(t testing.TB, path string, coords int) []testcase {
	f, err := os.Open("test_data/" + path)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	var tests []testcase
	scanner := bufio.NewScanner(f)
	for line := 1; scanner.Scan(); line++ {
		tt := testcase{label: path + ":" + strconv.Itoa(line)}
		parts := strings.Split(scanner.Text(), " ")

		if strings.TrimSpace(parts[0]) == "" || parts[0][0] == '#' {
			continue // skip empty and comment lines
		}

		if len(parts) != coords+1 {
			t.Fatalf("Coord count mismatch, got: %d want: %d", len(parts)-1, coords)
		}

		for _, field := range parts[:len(parts)-1] {
			n, err := strconv.ParseFloat(field, 64)
			if err != nil {
				t.Fatal(err)
			}
			tt.args = append(tt.args, n)
		}
		tt.sign, err = strconv.Atoi(parts[len(parts)-1])
		if err != nil {
			t.Fatal(err)
		}
		tests = append(tests, tt)
	}
	if err := scanner.Err(); err != nil {
		t.Fatal(err)
	}
	return tests
}

func assert(t *testing.T, want int, got float64) {
	t.Helper()
	if sign(got) != want {
		t.Errorf("want: %d; got: sign(%g)", want, got)
	}
}

func sign(n float64) int {
	if n < 0 {
		return -1
	}
	if n > 0 {
		return 1
	}
	return 0
}
