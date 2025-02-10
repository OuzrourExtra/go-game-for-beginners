package main
// import statement
import (
	"fmt"
	"math/rand"
	"strings"
	"bufio"
	"os"
	"log"
	"strconv"
	"time"
)
/*
 The Main function of the game
*/
func main(){
	// For Seed the random number
	seconds := time.Now().Unix()
	rand.NewSource(seconds)
	// Generate a random number
	nb := rand.Intn(100)
	// let the user choose a number
	fmt.Println("Choose a number : ")
	// initialize the variable succes to false ( he/she didn't find the number)
	succes := false
	// loop for 10 times to let the user choose a number
	for i:=0 ; i<10 ; i++{
		input , err1 := bufio.NewReader(os.Stdin).ReadString('\n')
		input = strings.TrimSpace(input)
		test , err2 := strconv.ParseFloat(input,64)
		if err1 != nil  {
			log.Fatal(err1)
		} else if err2 != nil {
			log.Fatal(err2)
		} else {
			if test == float64(nb) {
				fmt.Println("You win !")
				succes = true
			} else if test < float64(nb) {
				fmt.Println("Too Low")
			} else {
				fmt.Println("Too High")
			}
		}
	}
	// if the user didn't find the number after the 10 attempts
	if !succes {
		fmt.Println("You lose !")
		fmt.Println("The number was : ",nb)
	}
}