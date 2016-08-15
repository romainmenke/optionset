package optionset_test

import (
	"fmt"

	"github.com/romainmenke/optionset"
)

type AnimalTrait uint32

const (
	Fluffy AnimalTrait = iota + 1
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

func Example() {
	cat := optionset.New(Fluffy, EatsMeat, Wild)
	traits := cat.Options(Fluffy)
	for _, trait := range traits {
		fmt.Println(trait)
	}
}
