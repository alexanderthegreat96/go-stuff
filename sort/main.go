package main

import (
	"fmt"
	"sort"
)

type person struct {
	name     string
	age      int
	location string
}

// attached func to type person
func (p person) String() string {
	return fmt.Sprintf("%s: %d from %s", p.name, p.age, p.location)
}

type ByAge []person

func (a ByAge) Len() int {
	return len(a)
}
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].age < a[j].age }

// sorting can be done easily through the sort package

func main() {
	numbers := []int{3, 2, 1, 6, 7, 8, 9}
	sort.Ints(numbers)

	strings := []string{"mathew", "james", "car", "gandolfini"}
	sort.Strings(strings)

	fmt.Println(numbers)
	fmt.Println(strings)

	// custom sorting

	p1 := person{"james", 22, "nebraska"}
	p2 := person{"mathew", 44, "new york"}
	p3 := person{"alana", 27, "california"}
	p4 := person{"joey", 35, "alaska"}

	people := []person{p1, p2, p3, p4}

	sort.Sort(ByAge(people))
	for _, person := range people {
		fmt.Println(person.String())
	}

}
