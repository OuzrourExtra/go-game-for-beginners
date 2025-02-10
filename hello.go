package main

import (
	"fmt"
	"math/rand"
	"strings"
	"bufio"
	"os"
	"log"
	"strconv"
)

func main(){
	nb := rand.Intn(100)
	fmt.Println("Choose a number : ")
	succes := false
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
	if !succes {
		fmt.Println("You lose !")
		fmt.Println("The number was : ",nb)
	}
}