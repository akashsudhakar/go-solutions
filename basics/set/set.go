package main

import "fmt"

// Add Remove Contains  Union  Intersection

type Set map[interface{}]struct{}

func NewSet() *Set {
	s := Set{}
	return &s
}

func (s *Set) Print() {
	fmt.Println("Contents of Set : ")
	for key := range *s {
		fmt.Print(key, " ")
	}
}

func (s *Set) Add(key interface{}) {
	(*s)[key] = struct{}{}
}

func (s *Set) Remove(key interface{}) {
	delete((*s), key)
}

func (s *Set) Contains(key interface{}) bool {
	_, ok := (*s)[key]
	return ok
}

func (s *Set) Union(other *Set) *Set {
	newSet := NewSet()
	for key := range *s {
		(*newSet)[key] = struct{}{}
	}
	for key := range *other {
		(*newSet)[key] = struct{}{}
	}
	return newSet
}

func (s *Set) Intersection(other *Set) *Set {
	newSet := NewSet()
	for key := range *other {
		if _, ok := (*s)[key]; ok {
			(*newSet)[key] = struct{}{}
		}
	}
	return newSet
}
