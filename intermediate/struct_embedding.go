package main

import "fmt"

type person struct {
	name string
	age  int
}

type Employee struct {
	employeeInfo person //named field
	//person              // anonymous field embedded struct
	empId  string
	salary float64
}

func (p person) introduce() {
	fmt.Printf("Hi, I'm %s and I'm %d years old.\n", p.name, p.age)
}

// Overwrites the previous method
func (e Employee) introduce() {
	// fmt.Printf("Hi, I'm %s, employee ID: %s, and I earn %.2f.\n", e.name, e.empId, e.salary)
	fmt.Printf("Hi, I'm %s, employee ID: %s, and I earn %.2f.\n", e.employeeInfo.name, e.empId, e.salary)
}

func main() {

	emp := Employee{
		// person: person{name: "John", age: 30},
		employeeInfo: person{name: "John", age: 30},
		empId:        "E001", salary: 50000,
	}

	// since we changed the field to no longer be anonymous we also need to change the fields
	// fmt.Println("Name:", emp.name) // Accessing the embedded struct field emp.person.name
	fmt.Println("Name:", emp.employeeInfo.name)
	// fmt.Println("Age:", emp.age)
	fmt.Println("Age:", emp.employeeInfo.age)
	fmt.Println("Emp ID:", emp.empId)
	fmt.Println("Salary:", emp.salary)

	emp.introduce()
}
