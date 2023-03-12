package tickets

import (
	"fmt"
	"log"
	"os"
)

type Ticket struct {
}

func readFile() {
	data, err := os.ReadFile("../../tickets.csv")
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(data)

}

// ejemplo 1
//func GetTotalTickets(destination string) (int, error) {}

// ejemplo 2
//func GetMornings(time string) (int, error) {}

// ejemplo 3
func AverageDestination(destination string, total int) (int, error) {}
