package optionset

import (
	"encoding/json"
	"fmt"
	"testing"
)

type enum int

func (e enum) RawValue() uint32 {
	return uint32(e)
}

func (e enum) FromRaw(raw uint32) Option {
	switch raw {
	case 0:
		return none
	case 1:
		return foo
	case 2:
		return foos
	case 4:
		return fooz
	case 8:
		return a
	case 16:
		return b
	case 32:
		return c
	case 64:
		return d
	case 128:
		return e
	case 265:
		return ee
	case 512:
		return f
	case 1024:
		return g
	case 2048:
		return h
	case 4096:
		return i
	case 8192:
		return j
	default:
		return none
	}
}

const (
	none enum = 0
	foo  enum = 1
	foos enum = 2
	fooz enum = 4
	a    enum = 8
	b    enum = 16
	c    enum = 32
	d    enum = 64
	e    enum = 128
	ee   enum = 265
	f    enum = 512
	g    enum = 1024
	h    enum = 2048
	i    enum = 4096
	j    enum = 8192
)

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		myset := New(foo, fooz, g, j)
		_ = myset
	}
}

// func BenchmarkPower(b *testing.B) {
// 	var n uint32 = 6543
// 	for i := 0; i < b.N; i++ {
// 		power := nextPowerOfTwo(n)
// 		_ = power
// 	}
// }

func BenchmarkOption(b *testing.B) {
	for i := 0; i < b.N; i++ {
		myset := New(foo, fooz, g, j)
		options := myset.Options(Option(none))
		_ = options
	}
}

func BenchmarkOptionB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		myset := New(foo, fooz)
		options := myset.Options(Option(none))
		_ = options
	}
}

func BenchmarkOptionC(b *testing.B) {
	for i := 0; i < b.N; i++ {
		myset := New(foo, fooz, f, d, ee, c, g, j)
		options := myset.Options(Option(none))
		_ = options
	}
}

func BenchmarkMarshalSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		array := []uint32{1, 4, 512, 64, 265, 32, 1024, 8192}
		b, err := json.Marshal(array)
		if err != nil {
			fmt.Println(err)
			return
		}

		var newArray []uint32
		err = json.Unmarshal(b, &newArray)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func BenchmarkMarshalOptionSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		set := New(foo, fooz, f, d, ee, c, g, j)
		b, err := json.Marshal(set)
		if err != nil {
			fmt.Println(err)
			return
		}

		var newRawSet uint32
		err = json.Unmarshal(b, &newRawSet)
		if err != nil {
			fmt.Println(err)
			return
		}

		newSet := OptionSet(newRawSet)

		options := newSet.Options(Option(none))
		_ = options
	}
}

func BenchmarkMarshalSmallArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		array := []uint32{1, 4}
		b, err := json.Marshal(array)
		if err != nil {
			fmt.Println(err)
			return
		}

		var newArray []uint32
		err = json.Unmarshal(b, &newArray)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func BenchmarkMarshalSmallOptionSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		set := New(foo, fooz)
		b, err := json.Marshal(set)
		if err != nil {
			fmt.Println(err)
			return
		}

		var newRawSet uint32
		err = json.Unmarshal(b, &newRawSet)
		if err != nil {
			fmt.Println(err)
			return
		}

		newSet := OptionSet(newRawSet)

		options := newSet.Options(Option(none))
		_ = options
	}
}

func TestMain(t *testing.T) {
	myset := New(foo, fooz, g, j)
	options := myset.Options(Option(none))
	//fmt.Println(options)
	_ = options
}

func TestEmpty(t *testing.T) {
	myset := New()
	options := myset.Options(Option(none))
	//fmt.Println(options)
	_ = options
}

func TestMarshal(t *testing.T) {
	set := New(foo, fooz, f, d, ee, c, g, j)
	b, err := json.Marshal(set)
	if err != nil {
		fmt.Println(err)
		return
	}

	var newRawSet uint32
	err = json.Unmarshal(b, &newRawSet)
	if err != nil {
		fmt.Println(err)
		return
	}

	newSet := OptionSet(newRawSet)

	_ = newSet.Options(Option(none))
	if newSet != set {
		t.Fail()
	}
}
