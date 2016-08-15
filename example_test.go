package optionset_test

import (
	"fmt"

	"github.com/romainmenke/optionset"
)

type AnimalTrait uint32

const (
	Fluffy AnimalTrait = 1 << iota
	EatsMeat
	EatsPlants
	Flying
	Swimming
	Friendly
	Wild
)

func (t AnimalTrait) RawValue() uint32 {
	return uint32(t)
}

func (t AnimalTrait) FromRaw(raw uint32) optionset.Option {
	return AnimalTrait(raw)
}

func (t AnimalTrait) String() string {
	switch t.RawValue() {
	case 1:
		return "Fluffy"
	case 2:
		return "Eats Meat"
	case 4:
		return "Eats Plants"
	case 8:
		return "Flying"
	case 16:
		return "Swimming"
	case 32:
		return "Friendly"
	case 64:
		return "Wild"
	default:
		return "Unknown"
	}
}

func Example() {
	cat := optionset.New(Fluffy, EatsMeat, Wild)
	traits := cat.Options(Fluffy)
	for _, trait := range traits {
		fmt.Println(trait)
	}
}
