package dsalgo

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
)

func TestLinkedListAppend(t *testing.T) {
	linkedList := LinkedList{}

	Convey("Test linked list append", t, func() {
		Convey("should add the first value", func() {
			linkedList.Append(1)
			lastValue := linkedList.GetLastValue()
			assert.Equal(t, 1, lastValue)
		})

		Convey("should add the 2nd value", func() {
			linkedList.Append(0)
			linkedList.Append(1)
			linkedList.Append(2)
			lastValue := linkedList.GetLastValue()
			assert.Equal(t, 2, lastValue)
		})
	})
}

func TestLinkedListGetNthValue(t *testing.T) {
	linkedList := LinkedList{}

	Convey("Test get nth value", t, func() {
		Convey("should return first value", func() {
			linkedList.Append(1)
			nthValue := linkedList.GetValueByIndex(0)
			assert.Equal(t, 1, nthValue)
		})

		Convey("should get the 2nd value", func() {
			linkedList.Append(2)
			linkedList.Append(3)
			linkedList.Append(4)
			nthValue := linkedList.GetValueByIndex(1)
			assert.Equal(t, 2, nthValue)
		})

		Convey("should add the 3nd value", func() {
			linkedList.Append(5)
			nthValue := linkedList.GetValueByIndex(2)
			assert.Equal(t, 3, nthValue)
		})
	})
}

func TestLinkedListPrepend(t *testing.T) {
	linkedList := LinkedList{}

	Convey("Test linked list prepend", t, func() {
		Convey("should add the first value", func() {
			linkedList.Prepend(1)
			lastValue := linkedList.GetLastValue()
			assert.Equal(t, 1, lastValue)
		})

		Convey("should add the 2nd value", func() {
			linkedList.Append(0)
			linkedList.Append(1)
			linkedList.Append(2)
			lastValue := linkedList.GetLastValue()
			assert.Equal(t, 2, lastValue)
		})
	})
}
