package main

import (
	"fmt"

	"example.com/calculate"
	"example.com/greetings"
)

func calculateSum(a, b int) {
	fmt.Println(a * b)
}

//classses(struct and interfaces)

type Person struct {
	firstname, lastname string
}

type Employee struct {
	name   Person
	salary float64
}

type Developer struct {
	name     Person
	salary   float64
	language string
}

type Salary interface {
	increaseSalary()
}

//method pointer receiver

func (emp Employee) increaseSalary() {
	emp.salary = (float64(100) + float64(20)) / float64(100) * emp.salary

}

func (dev Developer) increaseSalary() {
	salary := (float64(100) + float64(35)) / float64(100) * dev.salary
	fmt.Println(salary)

}

func calculateSalary(s Salary) {
	s.increaseSalary()
}
func main() {
	// Get a greeting message and print it.

	p1 := Person{firstname: "timothy", lastname: "lubanga"}
	p2 := Person{firstname: "samson", lastname: "agendo"}
	dev1 := Developer{p2, 198000, "java"}
	calculateSalary(dev1)
	emp1 := Employee{p1, 24000}
	emp1.increaseSalary()
	fmt.Println(emp1.salary)
	message := greetings.Hello("Gladys")
	fmt.Println(message)
	calc := calculate.Reverse("timothy")
	fmt.Println(calc)
	calculateSum(45, 78)
}
