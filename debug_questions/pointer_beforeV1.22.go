package main

import (
	"fmt"
	"math/rand"
)

type UserDAO struct {
	ID  int
	DOB *string
}

// go 1.21
func main() {
	outputUsers := make([]*UserDAO, 0)
	inputDOB := []string{
		"1990-01-01",
		"1990-01-02",
		"1990-01-03",
	}

	for _, dob := range inputDOB {
		fmt.Println(&dob)
		outputUsers = append(outputUsers, &UserDAO{
			ID:  rand.Int(),
			DOB: &dob,
		})
	}
	for _, user := range outputUsers {
		fmt.Printf("ID: %d, DOB: %s\n", user.ID, *user.DOB)
	}
}
