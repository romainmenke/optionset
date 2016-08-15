package optionset

import "math"

// Option is an object that can be represented by a raw integer value.
// It must implement a RawValue method which returns a unique power of 2.
// It must also implement FromRaw which returns a new instance based on a power of 2.
type Option interface {
	RawValue() uint32
	FromRaw(raw uint32) Option
}

// OptionSet is typealias for uint64 and represents a collection of RawValues.
// These are derived from the sum of an array of Options.
type OptionSet uint32

// New returns an OptionSet from an array of Options.
func New(options ...Option) OptionSet {
	var sum uint32
	for _, value := range options {
		sum += value.RawValue()
	}
	return OptionSet(sum)
}

// Options returns an array of Options from an OptionSet
func (s OptionSet) Options(option Option) []Option {
	if uint32(s) <= 0 {
		return []Option{}
	}

	set := uint32(s)
	next := nextPowerOfTwo(set)
	maxSize := math.Log2(float64(next))
	options := make([]Option, uint(maxSize), uint(maxSize))
	newOption := option.FromRaw(1)
	var index int

	for set > 0 {
		next = nextPowerOfTwo(set)
		if set != next {
			newOption = option.FromRaw(next / 2)
			set -= next / 2
		} else {
			newOption = option.FromRaw(next)
			set -= next
		}
		options[index] = newOption
		index++
	}

	sized := options[:index]
	for i, j := 0, len(sized)-1; i < j; i, j = i+1, j-1 {
		sized[i], sized[j] = sized[j], sized[i]
	}
	return sized
}

func nextPowerOfTwo(v uint32) uint32 {
	v--
	v |= v >> 1
	v |= v >> 2
	v |= v >> 4
	v |= v >> 8
	v |= v >> 16
	v++
	return v
}
