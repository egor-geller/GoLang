package main

import (
	"fmt"
	"os"
	"sort"
	"time"
)

type Person struct {
	firstName string
	lastName  string
	birthDay  time.Time
}

/*func (p Person) String() string {
	return fmt.Sprintf("{%s %s %s}", p.firstName, p.lastName, p.birthDay.Format("2006-01-02"))
}*/

type People []Person

func (p People) Len() int { return len(p) }
func (p People) Less(i, j int) bool {
	if p[i].birthDay == p[j].birthDay {
		return p[i].firstName < p[j].firstName
	}
	return p[j].birthDay.Before(p[i].birthDay)
}
func (p People) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func main() {
	str := "2005-08-10"
	birthDay, err := time.Parse("2006-08-02", str)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	p := []Person{
		{"Ivan", "Ivanov", birthDay.AddDate(0, 7, 0)},
		{"Artiom", "Ivanov", birthDay.AddDate(-6, 7, 0)},
		{"Ivan", "Ivanov", birthDay.AddDate(-2, 7, -5)},
		{"Artiom", "Ivanov", birthDay.AddDate(0, 7, 0)},
	}

	sort.Sort(People(p))
	fmt.Println(p)
}
