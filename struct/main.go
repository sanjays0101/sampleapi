package main

import (
	"fmt"
	"reflect"
)

type State struct {
	Statename string
}

type User struct {
	Name string `required max:"100"`
}

type Village struct {
	State
	name         string
	totalFarmers int
	product      []string
}

func main() {
	aPerson := Village{
		name:         "xyz",
		totalFarmers: 100,
		product: []string{
			"A",
			"B",
			"C",
		},
	}
	aPerson.Statename = "KA"
	// Instantiate struct by position, Never use this!
	/*  bPerson := Village{ */
	//         "ABC",
	//         1200,
	//         []string{
	//                 "X",
	//                 "Y",
	//         },
	// }
	// fmt.Println(aPerson)
	// fmt.Println(bPerson)
	/*  */
	// local struct or anonymus struct
	fmt.Println(aPerson)
	animals := struct{ name string }{"Dog"}
	fmt.Println(animals)

	// struct value are independent
	dog := animals
	dog.name = "Lab"
	fmt.Println(dog)


	cPerson :=
		Village{
			State:        State{Statename: "UP"}, // Embeding
			name:         "UP_Guy",
			totalFarmers: 10000,
			product: []string{
				"Z",
				"Y",
				"X",
			},
		}

	fmt.Println(cPerson)
	// Tags

	user1 := User{Name: "Sanjay"}
	fmt.Println(user1)

	// Tags
	t := reflect.TypeOf(User{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)

	// Pointer to tage struct
	BCACollege := struct{ name string }{name: "KLE"}
	MPharmCollege := &BCACollege
	fmt.Println(BCACollege)
	fmt.Println(*MPharmCollege)

}
