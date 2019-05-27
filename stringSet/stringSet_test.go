package stringSet

import (
	"reflect"
	"sort"
	"testing"
)

func Test_Partition(t *testing.T) {
	set1 := New("a", "b", "c", "d")
	set2 := New("c", "d", "e", "f")

	set1Only, set2Only, intersect := Partition(set1, set2)

	set1OnlyArray := set1Only.Array()
	set2OnlyArray := set2Only.Array()
	intersectArray := intersect.Array()

	sort.Strings(set1OnlyArray)
	sort.Strings(set2OnlyArray)
	sort.Strings(intersectArray)

	if !reflect.DeepEqual(set1OnlyArray, []string{"a", "b"}) &&
		!reflect.DeepEqual(set2OnlyArray, []string{"e", "f"}) &&
		!reflect.DeepEqual(intersectArray, []string{"c", "d"}) {
		t.Fail()
	}
}

func Test_NewWithItems_Unique(t *testing.T) {
	set := New("a", "b", "c", "b")

	got := set.Array()
	sort.Strings(got)

	expected := []string{"a", "b", "c"}

	if !reflect.DeepEqual(got, expected) {
		t.Fail()
	}
}

func Test_New_Empty(t *testing.T) {
	set := New()
	got := set.Array()
	expected := []string{}

	if !reflect.DeepEqual(got, expected) {
		t.Fail()
	}
}

func Test_Add_Contains(t *testing.T) {
	set := New("a", "b", "c")
	set.Add("x")
	if !set.Contains("x") {
		t.Fail()
	}
}

func Test_Remove_NotContains(t *testing.T) {
	set := New("a", "b", "c")
	set.Remove("b")
	if set.Contains("b") {
		t.Fail()
	}
}
