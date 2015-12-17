// Example program with Interface, Composition and Method Overriding
package main

import (
	"fmt"
	"time"
)

type User interface {
	PrintName()
	PrintDetails()
}

type Person struct {
	FirstName, LastName string
	Dob                 time.Time
	Email, Location     string
}

//A person method
func (p Person) PrintName() {
	fmt.Printf("\n%s %s\n", p.FirstName, p.LastName)
}

//A person method
func (p Person) PrintDetails() {
	fmt.Printf("[Date of Birth: %s, Email: %s, Location: %s ]\n", p.Dob.String(), p.Email, p.Location)
}

type Admin struct {
	Person //type embedding for composition
	Roles  []string
}

//overrides PrintDetails
func (a Admin) PrintDetails() {
	//Call person PrintDetails
	a.Person.PrintDetails()
	fmt.Println("Admin Roles:")
	for _, v := range a.Roles {
		fmt.Println(v)
	}
}

type Member struct {
	Person //type embedding for composition
	Skills []string
}

//overrides PrintDetails
func (m Member) PrintDetails() {
	//Call person PrintDetails
	m.Person.PrintDetails()
	fmt.Println("Skills:")
	for _, v := range m.Skills {
		fmt.Println(v)
	}
}

type Team struct {
	Name, Description string
	Users             []User
}

func (t Team) GetTeamDetails() {
	fmt.Printf("Team: %s  - %s\n", t.Name, t.Description)
	fmt.Println("Details   of the team members:")
	for _, v := range t.Users {
		v.PrintName()
		v.PrintDetails()
	}
}

func main() {
	alex := Admin{
		Person{
			"Alex",
			"John",
			time.Date(1970, time.January, 10, 0, 0, 0, 0, time.UTC),
			"alex@email.com",
			"New York"},
		[]string{"Manage Team", "Manage Tasks"},
	}
	shiju := Member{
		Person{
			"Shiju",
			"Varghese",
			time.Date(1979, time.February, 17, 0, 0, 0, 0, time.UTC),
			"shiju@email.com",
			"Kochi"},
		[]string{"Go", "Docker", "Kubernetes"},
	}
	chris := Member{
		Person{
			"Chris",
			"Martin",
			time.Date(1978, time.March, 15, 0, 0, 0, 0, time.UTC),
			"chris@email.com",
			"Santa Clara"},
		[]string{"Go", "Docker"},
	}
	team := Team{
		"Go",
		"Golang CoE",
		[]User{alex, shiju, chris},
	}
	//get details of Team
	team.GetTeamDetails()
}
