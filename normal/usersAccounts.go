package main

import (
	"fmt"
	"sort"
)

/*
You are given two integer arrays, users and accounts,
where users[i] represents the ID of a user and accounts[i] represents the account ID associated with the user at index i.
It is possible for users to have multiple accounts, and similarly, multiple users can share the same account ID.
Your task is to write a function that identifies all unique users who share all same accounts ID with another user.
The function should return a list of these users' IDs without duplicates.

Input:
users = [1, 2, 3, 4, 2, 5, 1]
accounts = [1, 1, 2, 2, 3, 4, 3]

Output:
[1, 2, 3, 4]

Explanation:
Users 1 and 2 share account 1, 3.
Users 3 and 4 share account 2.

Input:
users = [1, 2, 3, 4, 2, 1, 1]
accounts = [1, 1, 2, 2, 3, 4, 3]

Output:
[3, 4]

Explanation:
Users 1 and 2 share account 1, 3. but 1 has 4 so it not the same.
Users 3 and 4 share account 2.
*/

func main() {
	users := []int{1, 2, 3, 4, 2, 5, 1}
	accounts := []int{1, 1, 2, 2, 3, 4, 3}
	fmt.Println(findUsersWithIdenticalAccounts(users, accounts)) // [1 2 3 4]
	users = []int{1, 2, 3, 4, 2, 1, 1}
	accounts = []int{1, 1, 2, 2, 3, 4, 3}
	fmt.Println(findUsersWithIdenticalAccounts(users, accounts)) // [3 4]
	users = []int{1, 2, 3, 4, 5, 6}
	accounts = []int{1, 2, 3, 4, 5, 6}
	fmt.Println(findUsersWithIdenticalAccounts(users, accounts)) // []
}

func findUsersWithSameAccounts(users, accounts []int) []int {
	userToAccounts := make(map[int][]int)
	for i, user := range users {
		account := accounts[i]
		userToAccounts[user] = append(userToAccounts[user], account)
	}

	// Sort accounts for each user
	for _, accs := range userToAccounts {
		sort.Ints(accs)
	}

	var result []int
	visited := make(map[int]bool)

	for user1, accs1 := range userToAccounts {
		if visited[user1] {
			continue
		}
		for user2, accs2 := range userToAccounts {
			if user1 != user2 && isEqualSlice(accs1, accs2) {
				if !visited[user1] {
					result = append(result, user1)
					visited[user1] = true
				}
				if !visited[user2] {
					result = append(result, user2)
					visited[user2] = true
				}
			}
		}
	}

	return result
}

func isEqualSlice(slice1, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}

func findUsersWithIdenticalAccounts(users, accounts []int) []int {
	userToAccounts := make(map[int]map[int]bool)
	for i, user := range users {
		account := accounts[i]
		if userToAccounts[user] == nil {
			userToAccounts[user] = make(map[int]bool)
		}
		userToAccounts[user][account] = true
	}

	var result []int
	visited := make(map[int]bool)

	for user1, accs1 := range userToAccounts {
		if visited[user1] {
			continue
		}
		for user2, accs2 := range userToAccounts {
			if user1 != user2 && !visited[user2] && isEqualSets(accs1, accs2) {
				result = append(result, user1, user2)
				visited[user1] = true
				visited[user2] = true
				break // Found a match, no need to compare user1 with others
			}
		}
	}
	return result
}

// isEqualSets checks if two sets (represented as map[int]bool) are equal
func isEqualSets(set1, set2 map[int]bool) bool {
	if len(set1) != len(set2) {
		return false
	}
	for key := range set1 {
		if !set2[key] {
			return false
		}
	}
	return true
}
