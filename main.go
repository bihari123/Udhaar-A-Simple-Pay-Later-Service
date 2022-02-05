package main

import (
	"bufio"
	"fmt"
	//"github.com/google/uuid"
	"os"
	"strings"
	"udhaar/cmd"
)


func IsExitComm(input string) bool{
	exitComm:=[]string{
		"exit",
		"Exit",
		"EXIT",
		"quit",
		"Quit",
		"QUIT",
		"stop",
		"STOP",
		"Stop",
		string(27) , //escape key
		//string( 3 ), // ctrl + c
	}

	for _,val:=range exitComm{
		if val == input{
			return true
		}
	}
	return false
}

func process(reader bufio.Reader) {
	//userList := make(map[string]model.User)

	//merchantList := make(map[string]model.Merchant)

	//allTransactionDetails :=make(map[uuid.UUID]model.Transaction)

	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Print(err)
		}

		input = strings.Replace(input, "\n", "", -1)

		if IsExitComm(input){
			fmt.Println("program terminated!")
			fmt.Println("Thanks For Exploring Pay-Later Terminal Client")
			os.Exit(0)
		} 
        
        cmd.ExecuteComm(input)
	//command.NewCommandHandler(consoleInput, userList, merchantList, allTransactionDetails)
	}
}

func main() {

	
	reader := bufio.NewReader(os.Stdin)
    
	process(*reader)

}