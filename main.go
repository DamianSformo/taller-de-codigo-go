package main

import (
	//"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
)

type Ticket struct {
	id string
	name string
	mail string
	destination string
	time string
	price string
}

func main() {
	//total, err := tickets.GetTotalTickets("Brazil")
	GetMornings()
}

// ejemplo 1
//func GetTotalTickets(destination string) (int, error) {
//}

/*

func readFile() {
	dataByte, err := os.ReadFile("tickets.csv")
	if err != nil{
		log.Fatal(err)
	}

	tickets := []Ticket{}
	count := 0

	data := strings.Split(string(dataByte), "\n")
	fmt.Println(data[0])

	
	for i:=0 ; i < len(data)-1 ; i++{
		line := strings.Split(string(data[i]), ",")
		i := Ticket{line[0], line[1], line[2], line[3], line[4], line[5]}
		tickets = append(tickets, i)
	}

	for i:=0 ; i < len(tickets) ; i++{
		if tickets[i].destination == "Jamaica"{
			count++
		}
	}

	fmt.Println(count)
}
*/


// ejemplo 2
func GetMornings() {
	dataByte, err := os.ReadFile("tickets.csv")
	if err != nil{
		log.Fatal(err)
	}

	madrugada := 0
	mañana := 0
	tarde := 0
	noche := 0

	data := strings.Split(string(dataByte), "\n")

	
	for i:=0 ; i < len(data) ; i++{
		line := strings.Split(string(data[i]), ",")
		time := strings.Split(string(line[4]), ":")
		//fmt.Println(time)
		//for j:=0 ; j < len(time) ; j++{
			t, err := strconv.Atoi(time[0]) 
			if err != nil {
				log.Fatal(err)
			}
			switch{
			case (t >= 0 && t < 7) :
				madrugada ++
				break
			case (t >= 7 && t < 13) :
				mañana ++
				break
			case (t >= 13 && t < 20) :
				tarde ++
				break
			case (t >= 20) :
				noche ++
			}
		}
	
	fmt.Println(madrugada)
	fmt.Println(mañana)
	fmt.Println(tarde)
	fmt.Println(noche)
}

