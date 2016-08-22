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

	set := uint32(s)

	if set <= 0 {
		return nil
	}

	next := set

	next--
	next |= next >> 1
	next |= next >> 2
	next |= next >> 4
	next |= next >> 8
	next |= next >> 16
	next++

	maxSize := uint(math.Log2(float64(next)))
	options := make([]Option, maxSize, maxSize)
	var index int

	for set > 0 {
		next = set
		next--
		next |= next >> 1
		next |= next >> 2
		next |= next >> 4
		next |= next >> 8
		next |= next >> 16
		next++

		if set != next {
			next = next / 2
		}
		set -= next
		options[index] = option.FromRaw(next)
		index++
	}

	sized := options[:index]
	for i, j := 0, len(sized)-1; i < j; i, j = i+1, j-1 {
		sized[i], sized[j] = sized[j], sized[i]
	}
	return sized
}
