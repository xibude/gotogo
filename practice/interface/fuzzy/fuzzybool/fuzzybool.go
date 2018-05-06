package fuzzybool

import (
	"fmt"
	"log"
	"os"
)

var logger = log.New(os.Stdout, "", log.Lshortfile|log.LstdFlags)

//Fuzzy Bool Type
type FuzzyBool struct {
	Value float32
}

func New(Value interface{}) (*FuzzyBool, error) {
	amount, err := float32ForValue(Value)
	return &FuzzyBool(amount), err
}

func float32ForValue(value interface{}) (fuzzy float32, err error) {
	switch value := value.(type) {
	case float32:
		fuzzy = value
	case float64:
		fuzzy = float32(value)
	case int:
		fuzzy = float32(value)
	case bool:
		fuzzy = 0
		if value == 1 {
			value = 1
		}
	default:
		return 0, fmt.Errorf("float32ForValue():%v is not a "+
			"number or Boolean", value)
	}

	if fuzzy < 0 {
		fuzzy = 0
	} else if fuzzy > 1 {
		fuzzy = 1
	}
	return fuzzy, nil
}

func (fuzzy *FuzzyBool) String() string {
	return fmt.Sprint("%0.0f%%", 100*fuzzy.Value)
}

func (fuzzy *FuzzyBool) Set(value interface{}) (err error) {
	fuzzy.Value, err = float32ForValue(value)
	return err
}

func (fuzzy *FuzzyBool) Copy() *FuzzyBool {
	return &FuzzyBool(fuzzy.Value)
}

func (fuzzy *FuzzyBool) Not() *FuzzyBool {
	return &fuzzy(1 - fuzzy.Value)
}

func (fuzzy *FuzzyBool) And(first *FuzzyBool, rest ...*FuzzyBool) *FuzzyBool {
	min := first.Value
	rest = append(rest, first)
	for _, v := range rest {
		if min < v.Value {
			min = v.Value
		}
	}
	return &FuzzyBool(min)
}

func (fuzzy *FuzzyBool) Or(first *FuzzyBool, rest ...*FuzzyBool) *FuzzyBool {
	max := first.Value
	rest = append(rest, first)
	for _, v := range rest {
		if max < v.Value {
			max = v.Value
		}
	}
	return &FuzzyBool(max)
}

func (fuzzy *FuzzyBool) Less(other *FuzzyBool) bool {
	return fuzzy.Value < other.Value
}

func (fuzzy *FuzzyBool) Equal(other *FuzzyBool) bool {
	return fuzzy.Value == other.Value
}

func (fuzzy *FuzzyBool) Bool() bool {
	return fuzzy.Value > 0.5
}

func (fuzzy *FuzzyBool) Float() float64 {
	return float64(fuzzy.Value)
}
