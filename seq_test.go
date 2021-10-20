package cobble

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCreationEmpty(t *testing.T) {
	c, e := NewSeq([]int{})
	if e != nil {
		t.Fatal("Error in construction: ", e)
	}
	if c.Size() != 0 {
		t.Fatal("Size MUST be 0, found: ", c.Size())
	}
}

func TestCreationMix(t *testing.T) {
	c, e := NewSeq([]interface{}{1, "false", 3})
	if e != nil {
		t.Fatal("Error in construction: ", e)
	}
	if c.Size() != 3 {
		t.Fatal("Size MUST be 3, found: ", c.Size())
	}
}

func TestSize(t *testing.T) {
	c, e := NewSeq([]int{1, 2, 3})
	if e != nil {
		t.Fatal("Error in construction: ", e)
	}
	if c.Size() != 3 {
		t.Fatal("Size MUST be 3, found: ", c.Size())
	}
}

func TestHead(t *testing.T) {
	if c, e := NewSeq([]int{1, 2, 3}); e == nil {
		if v, n := c.Head(); n == nil {
			if v != 1 {
				t.Fatal("Head MUST be 1; found: ", v.(int))
			}
		} else {
			t.Fatal(n)
		}
	} else {
		t.Fatal(e)
	}

}

func TestAppend1(t *testing.T) {
	if c, e := NewSeq([]int{1, 2, 3}); e == nil {
		c2 := c.Append(4)
		if !c2.IsError() && c2.Size() != 4 {
			t.Fatal("Expect collection size 4, got ", c2.Size())
		}
	} else {
		t.Fatal(e)
	}
}

func TestAppend2(t *testing.T) {
	if c, e := NewSeq([]int{1, 2, 3}); e == nil {
		c2 := c.Append("4")
		if !c2.IsError() {
			t.Fatal("Expect failure, got ", c2)
		}
	} else {
		t.Fatal(e)
	}
}

func TestExtendPositive(t *testing.T) {
	c1, e1 := NewSeq([]int{1, 2, 3})
	c2, e2 := NewSeq([]int{4, 5, 6})
	if e1 == nil && e2 == nil {
		c3 := c1.Extend(c2)
		if c3.IsError() || c3.Size() != 6 {
			t.Fatal("Expectation failed (no error, size == 6): ", c3.Error(), c3.Size())
		}
	} else {
		t.Fatal("Can't create sequences(s)", e1, e2)
	}
}

func TestExtendNegative(t *testing.T) {
	c1, e1 := NewSeq([]int{1, 2, 3})
	c2, e2 := NewSeq([]string{"4", "5", "6"})
	if e1 == nil && e2 == nil {
		c3 := c1.Extend(c2)
		if !c3.IsError() || c3.Size() != 3 {
			t.Fatal("Expectation failed (type-mismatch-error, size == 6): ", c3.Error(), c3.Size())
		}
	} else {
		t.Fatal("Can't create sequences(s)", e1, e2)
	}
}

func TestExtendInterfacePositive(t *testing.T) {
	c1, e1 := NewSeq([]int{1, 2, 3})
	c2, e2 := NewSeq([]interface{}{4, 5, 6})
	if e1 == nil && e2 == nil {
		c3 := c1.Extend(c2)
		if c3.IsError() || c3.Size() != 6 {
			t.Fatal("Expectation failed (type-mismatch-error, size == 6): ", c3.Error(), c3.Size())
		}
	} else {
		t.Fatal("Can't create sequences(s)", e1, e2)
	}
}

func TestExtendInterfaceNegative(t *testing.T) {
	c1, e1 := NewSeq([]int{1, 2, 3})
	c2, e2 := NewSeq([]interface{}{4, "5", 6})
	if e1 == nil && e2 == nil {
		c3 := c1.Extend(c2)
		if !c3.IsError() {
			t.Fatal("Expectation failed (type-mismatch-error): ", c3.Error())
		}
	} else {
		t.Fatal("Can't create sequences(s)", e1, e2)
	}
}

func TestExtendNativePositive(t *testing.T) {
	c1, e1 := NewSeq([]int{1, 2, 3})
	if e1 == nil {
		c3 := c1.ExtendN([]interface{}{4, 5, 6})
		if !c3.IsError() && c3.Size() != 6 {
			t.Fatal("Expectation failed error: ", c3.Error())
		}
	} else {
		t.Fatal("Can't create sequences(s)", e1)
	}
}

func TestExtendNativeNegative(t *testing.T) {
	c1, e1 := NewSeq([]int{1, 2, 3})
	if e1 == nil {
		c3 := c1.ExtendN(4)
		if !c3.IsError() {
			t.Fatal("Expectation failed (error): ", c3.Error())
		}
	} else {
		t.Fatal("Can't create sequences(s)", e1)
	}
}

func TestGet(t *testing.T) {
	c1, e1 := NewSeq([]int{1, 2, 3})
	if e1 == nil {
		el, _ := c1.Get(1)
		if reflect.TypeOf(el).Kind() != reflect.Int || el != 2 {
			t.Fatal("Expectation failed (error): ", reflect.TypeOf(el), el)
		}
		_, e2 := c1.Get(-1)
		if e2 == nil {
			t.Fatal("Must be error on negative index")
		}
		_, e3 := c1.Get(5)
		if e3 == nil {
			t.Fatal("Must be error too big index")
		}
	} else {
		t.Fatal("Can't create sequences(s)", e1)
	}
}

func TestToArray(t *testing.T) {
	c1, e1 := NewSeq([]int{1, 2, 3})
	if e1 == nil {
		ar := c1.ToArray()
		t.Logf("Got %T, %s", ar, ar)
		ax := ar.([]int)
		for i, v := range ax {
			t.Logf("%d: %d", i, v)
		}
		if len(ax) != 3 || ax[2] != 3 {
			t.Fatal("Expectation failed")
		}
	} else {
		t.Fatal("Can't create sequences(s)", e1)
	}
}

func ExampleSeq_ToArray_choose() {
	c1, _ := NewSeq([]int{1, 2, 3})
	array := c1.ToArray()
	switch array.(type) {
	case []int:
		res := array.([]int)
		fmt.Println("Result is ", res)
	default:
		fmt.Println("Unknown type: ", reflect.TypeOf(array))
	}
	// Output:
	// Result is  [1 2 3]
}

func ExampleSeq_ToArray_sure() {
	c1, _ := NewSeq([]int{1, 2, 3})
	array := c1.ToArray().([]int)
	fmt.Println("Result is ", array)
	// Output:
	// Result is  [1 2 3]
}

func ExampleNewSeq() {
	_, e1 := NewSeq([]int{}) // empty sequence of ints
	if e1 == nil {
	} // everything is ok
	c2, e2 := NewSeq([]int{1, 2, 3}) // sequence of 3 ints
	if e2 == nil && c2.Size() == 3 {
	} // everything is ok
	c3, e3 := NewSeq([]interface{}{1, "2", 3.1415}) // sequence of different types. yes, it's possible, but on your own risk
	if e3 == nil && c3.Size() == 3 {
	} // everything is ok
}
