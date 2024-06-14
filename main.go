package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	seconds := getSeconds()
	
	for {
		x,y := robotgo.Location()
		robotgo.MoveSmooth(x + 200, y)
		time.Sleep(time.Second * time.Duration(seconds))
		robotgo.MoveSmooth(x, y)
		time.Sleep(time.Second * time.Duration(seconds))
	}	
}

func getSeconds() int {
	fmt.Println("Enter seconds to wait:")
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Something went wrong")
	}

	line = strings.TrimSpace(line)

	i, err := strconv.Atoi(line)
	if err != nil {
		fmt.Print("You should use only numbers. ")
		getSeconds()
	}

	return i
}