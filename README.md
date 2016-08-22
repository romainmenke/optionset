[![godoc](https://godoc.org/github.com/romainmenke/optionset?status.svg)](https://godoc.org/github.com/romainmenke/optionset)

Coming from Swift, OptionSets are the one thing I miss in Go. Each option must be representable by a unique integer value. Each integer value must be a power of two.
OptionSets are sums of these integer values.

### Why?

An `OptionSet` is just a `uint32` so it is:
- light
- comparable / equatable

### But the overhead?

There is none. The Proof:

8 elements `Marshalled` and `Unmarshalled` as `Slice` and as `OptionSet`
More than 2x speed-up
```
BenchmarkMarshalSlice-4         	  500000	      3854 ns/op
BenchmarkMarshalOptionSet-4     	 1000000	      1684 ns/op
```

2 elements `Marshalled` and `Unmarshalled` as `Slice` and as `OptionSet`
Almost 0.5x speed-up
```
BenchmarkMarshalSmallArray-4    	 1000000	      1547 ns/op
BenchmarkMarshalSmallOptionSet-4	 1000000	      1076 ns/op
```

### What you won't get

There will be no fancy functions added to this package:
- Contained in OptionSet
- Filter
- Super/Sub OptionSet

Why? These are all iterative functions that can be optimized with the right context. Without that context it will result in slower code.
Once in memory Arrays are also really really fast and good snippets exist for all functions you might need. Doing these inline will result in faster execution.
