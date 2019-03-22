package pvsv

type Employee struct {
	Name string
	Age  int
}


func (e *Employee) ChangeName(newName string) {
	e.Name = newName
}


func (e *Employee) ChangeAge(newAge int) {
	e.Age = newAge
}
