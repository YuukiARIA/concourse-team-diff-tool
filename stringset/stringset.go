package stringset

type stringSet map[string]struct{}
type unit struct{}

// Contruct a new string set with given strings.
func New(xs ...string) stringSet {
	s := stringSet{}
	for _, x := range xs {
		s.Add(x)
	}
	return s
}

// Returns 3 disjoint sets: as - bs, bs - as, Intersect(as, bs)
func Partition(as, bs stringSet) (stringSet, stringSet, stringSet) {
	aOnly, bOnly, intersect := New(), New(), New()
	for a := range as {
		if bs.Contains(a) {
			intersect.Add(a)
		} else {
			aOnly.Add(a)
		}
	}
	for b := range bs {
		if !as.Contains(b) {
			bOnly.Add(b)
		}
	}
	return aOnly, bOnly, intersect
}

// Adds x to this set.
func (s stringSet) Add(x string) {
	s[x] = unit{}
}

// Removes x from this set.
func (s stringSet) Remove(x string) {
	delete(s, x)
}

// Returns true if this set contains x, otherwise false.
func (s stringSet) Contains(x string) bool {
	_, exists := s[x]
	return exists
}

// Convert this set to string array.
func (s stringSet) Array() []string {
	array := make([]string, 0)
	for x := range s {
		array = append(array, x)
	}
	return array
}
