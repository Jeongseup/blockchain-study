package clist

import (
	"fmt"
	"testing"
)

func TestSmall(t *testing.T) {
	testList := New()
	fmt.Println(testList.Len())

	el1 := testList.PushBack(1)
	el2 := testList.PushBack(2)
	el3 := testList.PushBack(3)

	if testList.Len() != 3 {
		t.Error("Expected len 3, got ", testList.Len())
	}

	fmt.Printf("%p | %v\n", el1, el1)
	fmt.Printf("%p | %v\n", el2, el2)
	fmt.Printf("%p | %v\n", el3, el3)

	r1 := testList.Remove(el1)
	fmt.Println(r1)
	fmt.Printf("%p | %v\n", el1, el1)
	fmt.Printf("%p | %v\n", el2, el2)
	fmt.Printf("%p | %v\n", el3, el3)
	fmt.Println(testList.Len())

	r2 := testList.Remove(el2)

	fmt.Printf("%p %v\n", el1, el1)
	fmt.Printf("%p %v\n", el2, el2)
	fmt.Printf("%p %v\n", el3, el3)
	fmt.Println(testList.Len())

	r3 := testList.Remove(el3)
	fmt.Printf("%p %v\n", el1, el1)
	fmt.Printf("%p %v\n", el2, el2)
	fmt.Printf("%p %v\n", el3, el3)
	fmt.Println(testList.Len())

	if r1 != 1 {
		t.Error("Expected 1, got ", r1)
	}
	if r2 != 2 {
		t.Error("Expected 2, got ", r2)
	}
	if r3 != 3 {
		t.Error("Expected 3, got ", r3)
	}
	if testList.Len() != 0 {
		t.Error("Expected len 0, got ", testList.Len())
	}

}
