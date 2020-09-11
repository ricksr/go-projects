package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	// "time"
)

func main() {
	var name string
	var userRating string

	//front end
	//take name as input
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("please enter name")
	name, _ = reader.ReadString('\n')

	//take rating from user and
	//show mssg according to rating

	readerr := bufio.NewReader(os.Stdin)
	fmt.Println("please enter your rating")
	userRating, _ = readerr.ReadString('\n')
	mynumRating, _ := strconv.ParseFloat(strings.TrimSpace(userRating), 64)

	//back end
	fmt.Printf("Hello %v\nthanks for rating our dosa center of %v \nat this time : %v\n", name, mynumRating, time.Now().Format(time.Stamp))

	if mynumRating == 5 {
		fmt.Println("Bonus to the team for 5 star service")
	} else if mynumRating == 4 || mynumRating == 3 {
		fmt.Println("We will surely improve")
	} else {
		fmt.Println("Need srs work")
	}
}
