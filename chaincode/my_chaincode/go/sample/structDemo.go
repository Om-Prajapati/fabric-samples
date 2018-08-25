package main

import (
	"fmt"
)

type Employee struct {
	FirstName   string
	LastName    string
	TotalLeaves int
	LeavesTaken int
}

func (emp Employee) LeavesRemaining() {
	fmt.Println("Emp detaits  is -->  ", emp.FirstName, emp.LastName, (emp.TotalLeaves - emp.LeavesTaken))
}

func main() {
	emp := Employee{
		FirstName:   "Om",
		LastName:    "Prajapati",
		TotalLeaves: 35,
		LeavesTaken: 20,
	}
	emp.LeavesRemaining()
}
