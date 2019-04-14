package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"parking_lot/cmd/parkinglot/services"
)

func main() {
	commandService := services.NewCommandService()

	if len(os.Args) > 1 {
		fileName := os.Args[1]
		
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Println(err)
		}
    	defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			commandService.Execute(scanner.Text())
		}
	} else {
		reader := bufio.NewReader(os.Stdin)
		line, _ := reader.ReadString('\n')
		for strings.Compare(strings.TrimRight(line, "\n"), "exit") != 0 {
			line = strings.TrimRight(line, "\n")
			commandService.Execute(line)
			line, _ = reader.ReadString('\n')
		}
	}
}
