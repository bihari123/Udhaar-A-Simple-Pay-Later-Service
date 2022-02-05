package cmd

import (
	prefix "udhaar/constants"
	utilities "udhaar/utilities"
	"fmt"
	"strings"
    controllers "udhaar/controllers"
	"github.com/spf13/cast"
)

var userList controllers.UserList=make(controllers.UserList, 0)
var merchantList controllers.MerchantList=make(controllers.MerchantList, 0)

func ExecuteComm(input string) error {

	command := strings.Split(input, " ")
	var rootCommand string

	if (len(command) >= 4 && len(command) < 6) || command[0]=="report" {
		rootCommand = string(command[0]) + " " + string(command[1])
	} else if len(command) >= 2 {
		rootCommand = string(command[0])
	} else {
		return utilities.InvalidCommand()
	}

	
	switch rootCommand {
	case prefix.NewUser:
		user,err:=controllers.AddUser(command,&userList)
		PrintMsg(err,"userAdd",user.Name)

	case prefix.NewMerchant:
		merchant,err:=controllers.AddMerchant(command,&merchantList)
		PrintMsg(err,"merchantAdd",merchant.Name)
	
	case prefix.NewTransaction:
		txnID,err:=controllers.Transaction(command,&userList,&merchantList)
		PrintMsg(err,"transaction",cast.ToString(txnID))

	case prefix.UpdateMerchant:
		newDiscount,err:=controllers.UpdateDiscount(command,&merchantList)
		PrintMsg(err,"updateDiscount",cast.ToString(newDiscount))
	case prefix.PayBack:
		due,err:= controllers.PayBack(command,&userList)
		PrintMsg(err,"payback",cast.ToString(due))
	case prefix.ReportDiscount:
		discount,err:=controllers.ShowMerchantDiscount(command,&merchantList)
		PrintMsg(err,"merchantDiscount",cast.ToString(discount))

	case prefix.ReportDues:
		dues,err:=controllers.ShowUserDues(command,&userList)
		PrintMsg(err,"userDues",cast.ToString(dues))

	case prefix.ReportUserAtCreditLimit:
		users,err:=controllers.ShowUserAtLimit(&userList)
		PrintMsg(err,"usersAtLimit",users)
	case prefix.ReportTotalDues:
		totalDues,err:=controllers.ShowTotalDues(&userList)
		PrintMsg(err,"totalDues",cast.ToString(totalDues))
	default:
		/* code */
		return utilities.InvalidCommand()
	}

	return nil
}

func PrintMsg(err error,msg string,val string){
	fmt.Println("")
	if err!=nil{
			fmt.Println(err)
		}else{
			switch msg{
			    case "userAdd":
			    	fmt.Println("user added :",val)

			    case "merchantAdd":
			        fmt.Println("merchant added ",val)

			    case "updateDiscount":
			        fmt.Println("merchant discount updated.",val)   
			    case "transaction":
			        fmt.Println("transaction complete. Txn Id:", val)
			    case "payback":
			        fmt.Println("Paid back to the merchant. Amount remaining: ",val)   
			    case "merchantDiscount":
			        fmt.Println("the total discount offered by this merchant is ",val)
			    case "userDues":
			        fmt.Println("the amount due on this user is: ",val)  
			    case "usersAtLimit":
			    	if len(val)>0{
			    	fmt.Println("the users at credit limit are the following: \n",val)                	
				    }else{
				    	fmt.Println("none of the users is at the credit limit")
				    }
				case "totalDues":
				    fmt.Println("total due on users is :",val)    
			        
			}

		}

	fmt.Println("\n\n")	
}
