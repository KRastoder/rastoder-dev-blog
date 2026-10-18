package main
import "fmt"

type person struct{
	name string 
	age int 
}

func changePerson(prsn *person , newPerson person) {
	prsn.age = newPerson.age
	prsn.name = newPerson.name
}
func newPerson (name string)*person{
	p := person{name:name}
	p.age = 42
	return &p
}

func main() {
	fatma := person{"fatma" , 23}
	fmt.Println(fatma.name, fatma.age )

	changePerson(&fatma, person{"sarac" ,67})
	fmt.Println(fatma.name, fatma.age )

	keni := newPerson("keni")

	fmt.Println(keni.name, keni.age)
	
}
