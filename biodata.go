package main

import (
	"fmt"
	"os"
	"strconv"
)

type Friend struct {
	Name         string
	Address      string
	Job          string
	ReasonForGo  string
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run biodata.go <roll_number>")
		return
	}

	rollNumber, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Roll number must be a number.")
		return
	}

	friends := initFriendData()
	if err := displayFriendData(friends, rollNumber); err != nil {
		fmt.Println(err)
	}
}

func initFriendData() []Friend {
	return []Friend{
		{"Yohanes", "Jl. Jenderal Soedirman", "Frontend", "Wants to learn a new programming language"},
		{"Harke", "Jl. Samratulangi", "Student", "Wants to develop web applications with Go"},
		{"Wauran", "Jl. Bethesda", "Programmer", "Interested in the active Go community"},
	}
}

func displayFriendData(friends []Friend, rollNumber int) error {
	if rollNumber <= 0 || rollNumber > len(friends) {
		return fmt.Errorf("Roll number is not valid")
	}

	fmt.Printf("Friend Data for Roll Number %d\n", rollNumber)
	fmt.Println("-------------------------")
	fmt.Println("Name:", friends[rollNumber-1].Name)
	fmt.Println("Address:", friends[rollNumber-1].Address)
	fmt.Println("Job:", friends[rollNumber-1].Job)
	fmt.Println("Reason for Choosing Go Class:", friends[rollNumber-1].ReasonForGo)

	return nil
}
