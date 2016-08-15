package optionset

import "math"

// Option is an object that can be represented by a raw integer value.
// It must implement a RawValue method which returns a unique power of 2.
// It must also implement FromRaw which returns a new instance based on a power of 2.
type Option interface {
	RawValue() uint64
	FrowRaw(raw uint64) Option
}

// OptionSet is typealias for uint64 and represents a collection of RawValues.
// These are derived from the sum of an array of Options.
type OptionSet uint64

// New returns an OptionSet from an array of Options.
func New(options ...Option) OptionSet {
	var sum uint64
	for _, value := range options {
		sum += value.RawValue()
	}
	return OptionSet(sum)
}

// Options returns an array of Options from an OptionSet
func (s OptionSet) Options(option Option) []Option {

	if uint64(s) <= 0 {
		return nil
	}

	set := uint64(s)
	next := nextPowerOfTwo(set)
	maxSize := math.Log2(float64(next))
	options := make([]Option, uint(maxSize), uint(maxSize))
	newOption := option.FrowRaw(1)
	var index int

	for set > 0 {
		next = nextPowerOfTwo(set)
		if set != next {
			newOption = option.FrowRaw(next / 2)
			options[index] = newOption
			index++
			set -= next / 2
		} else {
			newOption = option.FrowRaw(next)
			options[index] = newOption
			index++
			set -= next
		}
	}

	sized := options[:index]

	for i, j := 0, len(sized)-1; i < j; i, j = i+1, j-1 {
		sized[i], sized[j] = sized[j], sized[i]
	}

	return sized
}

func nextPowerOfTwo(v uint64) uint64 {
	v--
	v |= v >> 1
	v |= v >> 2
	v |= v >> 4
	v |= v >> 8
	v |= v >> 16
	v++
	return v
}

func previousPowerOfTwo(v uint64) uint64 {
	return nextPowerOfTwo(v) / 2
}
