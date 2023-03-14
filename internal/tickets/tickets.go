package tickets

import (
	"log"
	"os"
	"strings"
	"strconv"
	"errors"
)

var ErrTicketsNull = errors.New("No se encontraron Tickets")
var ErrFileNotFind = errors.New("No se encontró el Archivo buscado")

type Ticket struct {
	id string
	name string
	mail string
	destination string
	time string
	price string
}

type Tickets struct {
	Tickets []Ticket
}

// ejercicio 1
func (t *Tickets) GetTotalTickets(destination string) (int, error) {
	count := 0;

	if len(t.Tickets) < 1 {
		return 0, ErrTicketsNull
	}

	for i:=0 ; i < len(t.Tickets) ; i++{
		if t.Tickets[i].destination == destination{
			count++
		}
	}
	return count, nil
}

// ejercicio 2
func (t *Tickets) GetMornings(time string) (int, error) {
	count:= 0

	if len(t.Tickets) < 1 {
		return 0, ErrTicketsNull
	}

	for i:=0 ; i < len(t.Tickets) ; i++{
		t := strings.Split(string(t.Tickets[i].time), ":")

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

	return count, nil
}

// ejercicio 3
func (t *Tickets) AverageDestination(destination string) (float64, error) {

	quantity, error := t.GetTotalTickets(destination)

	if error != nil {
		return 0, ErrTicketsNull
	}

	total := float64(len(t.Tickets))

	average := float64(quantity) * 100 / total

	return average, nil
}

func ReadFile(doc string) []Ticket{

	tickets := []Ticket{}

	dataByte, err := os.ReadFile(doc)
	if err != nil{
		log.Fatal(ErrFileNotFind)
	}

	data := strings.Split(string(dataByte), "\n")

	for i:=0 ; i < len(data) ; i++{
		line := strings.Split(string(data[i]), ",")
		i := Ticket{line[0], line[1], line[2], line[3], line[4], line[5]}
		tickets = append(tickets, i)
	}

	return tickets
}