package main

import (
	"taller-de-codigo-go/internal/tickets"
	"fmt"
)
func main() {
	
	//total, err := tickets.GetTotalTickets("Brazil")
	fmt.Println(tickets.GetTotalTickets("Jamaica"))
	//fmt.Println(tickets.GetMornings("mañana"))
	fmt.Println(tickets.AverageDestination("Jamaica"))
}

