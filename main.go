package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"math/rand"
	"strconv"
	"time"
)

type week struct {
	WeekNumber		int
	TributeId			int
}

type employee struct {
	Id							int
	Name						string
	StartingWeek		int
	LastTributeWeek	int
	Eligible				bool
}

func (week week) print(employees []employee) {
	fmt.Printf("Week %v:", week.WeekNumber)
	if week.TributeId > -1 {
		fmt.Printf("  Tribute: %v \n", employees[week.TributeId].Name)
	} else {
		fmt.Printf("  No Tribute \n")
	}
}

func main() {
	employees := make([]employee, 0)
	weeks := make([]week, 0)

	fmt.Printf("Building employee roster \n")
	employeeFile, _ := os.Open("employees.csv")
	employeeReader := csv.NewReader(employeeFile)

	for {
		row, err := employeeReader.Read()
		if err != nil {break;}

		var employee employee
		employee.Id = len(employees)
		employee.Name = row[0]
		val, _ := strconv.ParseInt(row[1],10,0)
		employee.StartingWeek = int(val)
		employee.LastTributeWeek = int(val)
		employee.Eligible = true

		employees = append(employees, employee)
	}
	fmt.Printf("  loaded %v employees \n", len(employees))

	fmt.Printf("Building tribute history \n")
	weekFile, _ := os.Open("weeks.csv")
	weekReader := csv.NewReader(weekFile)

	for {
		row, err := weekReader.Read()
		if err != nil {break;}

		var week week
		week.WeekNumber = len(weeks)
		val, _ := strconv.ParseInt(row[0],10,0)
		week.TributeId = int(val)
		if week.TributeId > -1 {
			employees[week.TributeId].LastTributeWeek = week.WeekNumber
		}

		weeks = append(weeks, week)
		week.print(employees)
	}
	fmt.Printf("  loaded %v weeks \n", len(weeks))

	fmt.Printf("Calculating ticket totals \n")
	lastWeek := len(weeks)-1
	tickets := make([]int, len(employees))
	totalTickets := 0
	for _, employee := range employees {
		weeksSinceLastTribute := lastWeek - employee.LastTributeWeek
		tickets[employee.Id] = weeksSinceLastTribute
		fmt.Printf("  %v: %v\n", employee.Name, weeksSinceLastTribute)
		totalTickets += weeksSinceLastTribute
	}

	fmt.Printf("Total tickets: %v \n", totalTickets)

	fmt.Printf("Selecting a tribute \n")
	fmt.Printf(" ... (press enter) \n")

	var throwaway string
	n, err := fmt.Scanf("%s", &throwaway)
	if err != nil || n != 1 {
		fmt.Printf(err.Error())
	}

	source := rand.NewSource(time.Now().Unix())
	random := rand.New(source)
	randomTicket := random.Intn(totalTickets)
	fmt.Printf("  Picked ticket %v \n", randomTicket)
	newTributeId := 0
	for _, employee := range employees {
		newTributeId = employee.Id
		randomTicket -= tickets[employee.Id]
		if randomTicket <= 0 {break;}
	}
	fmt.Printf("New tribute: %v \n", employees[newTributeId].Name)
}
