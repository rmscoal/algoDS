package mock_assestment

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"testing"
)

/**
A transaction is possibly invalid if:

the amount exceeds $1000, or;
if it occurs within (and including) 60 minutes of another transaction with the same name in a different city.
You are given an array of strings transaction where transactions[i] consists of comma-separated values representing the name, time (in minutes), amount, and city of the transaction.

Return a list of transactions that are possibly invalid. You may return the answer in any order.

Example 1:

Input: transactions = ["alice,20,800,mtv","alice,50,100,beijing"]
Output: ["alice,20,800,mtv","alice,50,100,beijing"]
Explanation: The first transaction is invalid because the second transaction occurs within a difference of 60 minutes, have the same name and is in a different city. Similarly the second one is invalid too.
Example 2:

Input: transactions = ["alice,20,800,mtv","alice,50,1200,mtv"]
Output: ["alice,50,1200,mtv"]
Example 3:

Input: transactions = ["alice,20,800,mtv","bob,50,1200,mtv"]
Output: ["bob,50,1200,mtv"]
*/

type Transaction struct {
	Name   string
	Time   int
	Amount int
	City   string
}

func (t Transaction) IsSuspicious(u Transaction) bool {
	return math.Abs(float64(t.Time)-float64(u.Time)) <= 60 &&
		t.City != u.City
}

func (t Transaction) Illegal() bool {
	return t.Amount > 1000
}

func (t Transaction) String() string {
	return fmt.Sprintf("%s,%d,%d,%s", t.Name, t.Time, t.Amount, t.City)
}

func NewTransaction(tran string) Transaction {
	var t Transaction
	split := strings.Split(tran, ",")
	t.Name = split[0]
	t.Time, _ = strconv.Atoi(split[1])
	t.Amount, _ = strconv.Atoi(split[2])
	t.City = split[3]
	return t
}

// TODO: Dynamic approach
func InvalidTransactions(transactions []string) []string {
	invalids := make([]string, 0)
	transByName := make(map[string][]Transaction)

	for _, tt := range transactions {
		transaction := NewTransaction(tt)
		transByName[transaction.Name] = append(transByName[transaction.Name],
			transaction)
	}

	for _, tt := range transactions {
		curr := NewTransaction(tt)
		if curr.Illegal() {
			invalids = append(invalids, curr.String())
			continue
		}
		for _, other := range transByName[curr.Name] {
			if curr.IsSuspicious(other) {
				invalids = append(invalids, curr.String())
				break
			}
		}
	}

	return invalids
}

func TestInvalidTransactions(t *testing.T) {
	// fmt.Println(invalidTransactions([]string{
	// 	"alice,20,800,mtv", "alice,50,1200,mtv",
	// }))
	// fmt.Println(InvalidTransactions([]string{
	// 	"alice,20,800,mtv", "alice,50,100,mtv", "alice,51,100,frankfurt",
	// }))
	// fmt.Println(InvalidTransactions([]string{
	// 	"alice,20,800,mtv", "bob,50,1200,mtv", "alice,20,800,mtv",
	// 	"alice,50,1200,mtv", "alice,20,800,mtv", "alice,50,100,beijing",
	// }))
	fmt.Println(InvalidTransactions([]string{
		"alex,676,260,bangkok", "bob,656,1366,bangkok", "alex,393,616,bangkok",
		"bob,820,990,amsterdam", "alex,596,1390,amsterdam",
	}))
}
