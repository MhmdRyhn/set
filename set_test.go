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

func TestContains(t *testing.T) {
	si := New[int]()
	assert.Equal(t, si.Len(), 0)
	ss := New("abc", "def")
	assert.Equal(t, si.Len(), 2)
	for _, item := range ss.Members() {
		assert.True(t, ss.Contains(item))
	}
}

func TestAdd(t *testing.T) {
	si := New(1, 2)
	assert.Equal(t, si.Len(), 2)
	for _, item := range si.Members() {
		assert.True(t, si.Contains(item))
	}
	si.Add(3)
	assert.Equal(t, si.Members(), []int{1, 2, 3})
	assert.Equal(t, si.Len(), 3)
	for _, item := range si.Members() {
		assert.True(t, si.Contains(item))
	}
}
func TestRemove(t *testing.T) {
	si := New(1, 2)
	assert.Equal(t, si.Members(), []int{1, 2})
	si.Remove(1)
	assert.Equal(t, si.Members(), []int{2})
}

func TestAddThenRemove(t *testing.T) {
	si := New(1, 2)
	assert.Equal(t, si.Members(), []int{1, 2})
	si.Add(3)
	assert.Equal(t, si.Len(), 3)
	for _, item := range si.Members() {
		assert.True(t, si.Contains(item))
	}
	si.Remove(2)
	assert.Equal(t, si.Len(), 2)
	for _, item := range si.Members() {
		assert.True(t, si.Contains(item))
	}
}
