package slice

import (
	"fmt"
)

type (
	Person     uint32
	PersonSlice []Person
)

func (n *PersonSlice) Add(pPerson *Person) {
	*n = append(*n, *pPerson)
}

func (n *PersonSlice) Remove(pPerson *Person) error {
	if len(*n) == 0 {
		return fmt.Errorf("nodeSlice is nil")
	}
	for i, v := range *n {
		if *pPerson == v {
			*n = append((*n)[:i], (*n)[i+1:]...)
		}
	}
	return nil
}

func (n *PersonSlice) Find(pPerson *Person) *Person {
	for i, v := range *n {
		if *pPerson == v {
			return &(*n)[i]
		}
	}
	return nil
}
