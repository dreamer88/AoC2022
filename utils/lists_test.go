package utils

import (
	"math/rand"
	"testing"
	"time"
)

type CopyTest struct {
	x, y int
}

func newTestObj() CopyTest {
	return CopyTest{rand.Int(), rand.Int()}
}

func compareTestObj(a *CopyTest, b *CopyTest) bool {
	return a.x == b.x && a.y == b.y
}

func compare1D(t *testing.T, orig *[]CopyTest, new *[]CopyTest) {
	if len(*orig) != len(*new) {
		t.Errorf("lengths were not equal, original: %d new: %d", len(*orig), len(*new))
	}

	for i := 0; i < len(*orig); i++ {
		if !compareTestObj(&(*orig)[i], &(*new)[i]) {
			t.Errorf("values different at index %d", i)
		}
	}
}

func compare2D(t *testing.T, orig *[][]CopyTest, new *[][]CopyTest) {
	if len(*orig) != len(*new) {
		t.Errorf("lengths were not equal, original: %d new: %d", len(*orig), len(*new))
	}

	for i := 0; i < len(*orig); i++ {
		compare1D(t, &(*orig)[i], &(*new)[i])
	}
}

func compare3D(t *testing.T, orig *[][][]CopyTest, new *[][][]CopyTest) {
	if len(*orig) != len(*new) {
		t.Errorf("lengths were not equal, original: %d new: %d", len(*orig), len(*new))
	}

	for i := 0; i < len(*orig); i++ {
		compare2D(t, &(*orig)[i], &(*new)[i])
	}
}

func randArray1D() []CopyTest {
	count := rand.Int()%5 + 20
	vals := make([]CopyTest, count)
	for i := 0; i < count; i++ {
		vals[i] = newTestObj()
	}

	return vals
}

func randArray2D() [][]CopyTest {
	count := rand.Int()%5 + 20
	vals := make([][]CopyTest, count)
	for i := 0; i < count; i++ {
		vals[i] = randArray1D()
	}

	return vals
}

func randArray3D() [][][]CopyTest {
	count := rand.Int()%5 + 20
	vals := make([][][]CopyTest, count)
	for i := 0; i < count; i++ {
		vals[i] = randArray2D()
	}

	return vals
}

func TestCopy1D(t *testing.T) {
	seed := time.Now().UnixNano()
	t.Logf("Running with seed %d", seed)
	rand.Seed(time.Now().UnixNano())

	vals := randArray1D()
	other := Copy1D(vals)

	compare1D(t, &vals, &other)
}

func TestCopy2D(t *testing.T) {
	seed := time.Now().UnixNano()
	t.Logf("Running with seed %d", seed)
	rand.Seed(time.Now().UnixNano())

	vals := randArray2D()
	other := Copy2D(vals)

	compare2D(t, &vals, &other)
}

func TestCopy3D(t *testing.T) {
	seed := time.Now().UnixNano()
	t.Logf("Running with seed %d", seed)
	rand.Seed(time.Now().UnixNano())

	vals := randArray3D()
	other := Copy3D(vals)

	compare3D(t, &vals, &other)
}

func TestInsert(t *testing.T) {
	seed := time.Now().UnixNano()
	t.Logf("Running with seed %d", seed)
	rand.Seed(time.Now().UnixNano())

	vals := randArray1D()
	copy := Copy1D(vals)
	inserted_val := newTestObj()

	for i := 0; i < len(vals); i++ {
		new_vals, err := Insert(vals, inserted_val, i)
		if err != nil {
			t.Errorf("insert element failed, reason: %s", err.Error())
		} else {
			compare1D(t, &vals, &copy)

			if len(vals)+1 != len(new_vals) {
				t.Errorf("insert did not add an element, original: %d new: %d", len(vals), len(new_vals))
			}

			if !compareTestObj(&inserted_val, &new_vals[i]) {
				t.Errorf("insert did not set the correct inserted element")
			}
		}
	}
}
