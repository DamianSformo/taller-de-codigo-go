package main

import (
	"github.com/DamianSformo/taller-de-codigo-go/internal/tickets"
	"fmt"
	"errors"
)

func main() {

	doc := "tickets.csv"
	destination := "Colombia"
	time := "noche"

	ts := tickets.Tickets{
		Tickets: tickets.ReadFile(doc),
	}

	getTotalDestinatio, err1 := ts.GetTotalTickets(destination)
	if errors.Is(err1, tickets.ErrTicketsNull) {
		fmt.Println(err1)
	} else {
		fmt.Printf("En el día de hoy se han registrado un total de %d vuelos a %s \n", getTotalDestinatio, destination)
	}

	getTotalMornings, err2 := ts.GetMornings(time)
	if errors.Is(err2, tickets.ErrTicketsNull) {
		fmt.Println(err2)
	} else {
		fmt.Printf("Durante la %s viajaron en total %d personas \n", time, getTotalMornings)
	}

	getAverage, err3 := ts.AverageDestination(destination)
	if errors.Is(err3, tickets.ErrTicketsNull) {
		fmt.Println(err3)
	} else {
		fmt.Printf("En el día de hoy un %.2f%s de personas viajaron a %s \n", getAverage, "%", destination)
	}
	
}