package tickets

import (
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

// ejemplo 1
//func GetTotalTickets(destination string) (int, error) {
func GetTotalTickets(destination string) (int) {
	tickets := readFile()
	count := 0;

	for i:=0 ; i < len(tickets) ; i++{
		if tickets[i].destination == destination{
			count++
		}
	}
	return count
}

// ejemplo 2
//func GetMornings(time string) (int, error) {}
func GetMornings(time string) (int) {
	tickets := readFile()

	count:= 0

	for i:=0 ; i < len(tickets) ; i++{
		t := strings.Split(string(tickets[i].time), ":")

		h, err := strconv.Atoi(t[0]) 
		if err != nil {
			log.Fatal(err)
		}

		switch time{
		case "madrugada" :
			if h >= 0 && h < 7 {
				count++
			}
		case "mañana" :
			if h >= 7 && h < 13{
				count++
			}
		case "tarde" :
			if h >= 13 && h < 20{
				count++
			}
		case "noche" :
			if h >= 20 {
				count ++
			}
		}
	}

	return count
}

// ejemplo 3
//func AverageDestination(destination string, total int) (int, error) {}
func AverageDestination(destination string) (float64) {
	tickets := readFile()

	quantity := float64(GetTotalTickets(destination))
	total := float64(len(tickets))

	average := quantity * 100 / total

	return average
}

func readFile() []Ticket{

	tickets := []Ticket{}

	dataByte, err := os.ReadFile("tickets.csv")
	if err != nil{
		log.Fatal(err)
	}

	data := strings.Split(string(dataByte), "\n")

	for i:=0 ; i < len(data) ; i++{
		line := strings.Split(string(data[i]), ",")
		i := Ticket{line[0], line[1], line[2], line[3], line[4], line[5]}
		tickets = append(tickets, i)
	}

	return tickets
}
