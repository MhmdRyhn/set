package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLen(t *testing.T) {
	si := New[int]()
	assert.Equal(t, si.Len(), 0)

	si = New(1, 2)
	assert.Equal(t, si.Len(), 2)
}

func TestIsEmpty(t *testing.T) {
	si := New[int]()
	assert.True(t, si.IsEmpty())

	si = New(1, 2)
	assert.False(t, si.IsEmpty())
}

func TestContains(t *testing.T) {
	si := New[int]()
	assert.Equal(t, si.Len(), 0)

	ss := New("abc", "def")
	assert.Equal(t, ss.Len(), 2)
	assert.True(t, ss.Contains("abc"))
	assert.True(t, ss.Contains("def"))
}

func TestAdd(t *testing.T) {
	si := New(1, 2)
	assert.Equal(t, si.Len(), 2)
	assert.True(t, si.Contains(1))
	assert.True(t, si.Contains(2))

	si.Add(1, 3)
	assert.Equal(t, si.Len(), 3)
	assert.True(t, si.Contains(1))
	assert.True(t, si.Contains(2))
	assert.True(t, si.Contains(3))
}

func TestRemove(t *testing.T) {
	si := New(1, 2, 3)
	si.Remove(2)
	assert.Equal(t, si.Len(), 2)
	assert.True(t, si.Contains(1))
	assert.True(t, si.Contains(3))
}

func TestAddThenRemove(t *testing.T) {
	si := New(1, 2)
	assert.Equal(t, si.Len(), 2)
	assert.True(t, si.Contains(1))
	assert.True(t, si.Contains(2))

	si.Add(3)
	assert.Equal(t, si.Len(), 3)
	assert.True(t, si.Contains(1))
	assert.True(t, si.Contains(2))
	assert.True(t, si.Contains(3))

	si.Remove(2)
	assert.Equal(t, si.Len(), 2)
	assert.True(t, si.Contains(1))
	assert.True(t, si.Contains(3))
}

func TestUnion(t *testing.T) {
	ss1 := New("a", "b")
	ss2 := New("b", "c")
	result := ss1.Union(*ss2)
	assert.Equal(t, result.Len(), 3)
	assert.True(t, result.Contains("a"))
	assert.True(t, result.Contains("b"))
	assert.True(t, result.Contains("c"))
}

func TestIntersection(t *testing.T) {
	ss1 := New("a", "b", "c")
	ss2 := New("b", "c", "d")
	result := ss1.Intersection(*ss2)
	assert.Equal(t, result.Len(), 2)
	assert.True(t, result.Contains("b"))
	assert.True(t, result.Contains("c"))
}

func TestDifference(t *testing.T) {
	ss1 := New("a", "b", "c")
	ss2 := New("c", "d")
	result := ss1.Difference(*ss2)
	assert.Equal(t, result.Len(), 2)
	assert.True(t, result.Contains("a"))
	assert.True(t, result.Contains("b"))
}

func TestIsEqual(t *testing.T) {
	si1 := New(1, 2)
	si2 := New(1)
	assert.False(t, si1.IsEqual(*si2))

	si1 = New(1, 2)
	si2 = New(1, 2)
	assert.True(t, si1.IsEqual(*si2))
}

func TestIsDisjoint(t *testing.T) {
	si1 := New[int]()
	si2 := New[int]()
	assert.False(t, si1.IsDisjoint(*si2))

	si1 = New(1, 2)
	si2 = New(2, 3)
	assert.False(t, si1.IsDisjoint(*si2))

	si1 = New[int]()
	si2 = New(1, 2, 3)
	assert.True(t, si1.IsDisjoint(*si2))

	si1 = New(1, 2)
	si2 = New(3, 4)
	assert.True(t, si1.IsDisjoint(*si2))
}

func TestIsSubsetOf(t *testing.T) {
	ss1 := New("a", "b")
	ss2 := New("a", "b", "c")
	assert.True(t, ss1.IsSubsetOf(*ss2, false))

	ss1 = New("a", "b", "c")
	ss2 = New("a", "b", "c")
	assert.True(t, ss1.IsSubsetOf(*ss2, false))

	ss1 = New("a", "d")
	ss2 = New("a", "b", "c")
	assert.False(t, ss1.IsSubsetOf(*ss2, false))

	ss1 = New[string]()
	ss2 = New[string]()
	assert.True(t, ss1.IsSubsetOf(*ss2, false))
}

func TestIsProperSubsetOf(t *testing.T) {
	ss1 := New("a", "b")
	ss2 := New("a", "b", "c")
	assert.True(t, ss1.IsSubsetOf(*ss2, true))

	ss1 = New("a", "b", "c")
	ss2 = New("a", "b", "c")
	assert.False(t, ss1.IsSubsetOf(*ss2, true))

	ss1 = New("a", "d")
	ss2 = New("a", "b", "c")
	assert.False(t, ss1.IsSubsetOf(*ss2, true))

	ss1 = New[string]()
	ss2 = New[string]()
	assert.False(t, ss1.IsSubsetOf(*ss2, true))
}
