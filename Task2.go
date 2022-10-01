package main

import ("fmt")

type employee struct {
	name     string
	salary   int
	position string
}

type company struct {
	companyName string
	employees   []employee
}

func AddEmployee(emp employee) {

}

func main() {

	var emp1 employee
	var emp2 employee
	var emp3 employee

	emp1.name = "Tehami"
	emp1.position = "Programmer"
	emp1.salary = 100000

	emp2.name = "Ahmad"
	emp2.position = "Hacker"
	emp2.salary = 200000

	emp3.name = "Yaqoob"
	emp3.position = "Developer"
	emp3.salary = 300000

	emplys := []employee{emp1, emp2, emp3}

	var comp company

	comp.companyName = "Jazz"
	comp.employees = emplys

	fmt.Println("\n")
	fmt.Println("Company Name : ", comp.companyName)
	fmt.Println("Employees : ", comp.employees)

}
