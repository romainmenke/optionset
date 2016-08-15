package optionset

import "testing"

type enum int

func (e enum) RawValue() uint32 {
	return uint32(e)
}

func (e enum) FrowRaw(raw uint32) Option {
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

func BenchmarkPower(b *testing.B) {
	var n uint32 = 6543
	for i := 0; i < b.N; i++ {
		power := nextPowerOfTwo(n)
		_ = power
	}
}

func BenchmarkOption(b *testing.B) {
	for i := 0; i < b.N; i++ {
		myset := New(foo, fooz, g, j)
		options := myset.Options(Option(none))
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
