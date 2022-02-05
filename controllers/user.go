package controllers

import (
	
    "fmt"
	e "udhaar/utilities"

	"github.com/spf13/cast"
)

func AddUser(command []string, userList *UserList) (newUser User,err error) {
	name := command[2]
	email := command[3]
	creditLimit := cast.ToFloat64(command[4])
	list := *(userList)
	_, found := list[name]
	if found {
		err= e.UserAlreadyPresent()
		return
	}
	newUser = User{
		Name:        name,
		Email:       email,
		CreditLimit: creditLimit,
	}
	list[name] = newUser
	return 
}
func ShowUserDues(command []string,userList *UserList)(dues float64,err error){
	userName:= command[2]
	listUser:=*(userList)
	_, found := listUser[userName]
	if !found {
		err= e.UserMissing()
		return
	}

	dues= listUser[userName].Dues
	return
}

func ShowUserAtLimit(userList *UserList)(users string,err error){
    listUser:=*(userList)
    if len(listUser)>0{
    	for name,details:=range listUser{
    		if details.Dues == details.CreditLimit{
    			users+="\n"+name
    		}
    	}
    }else{
    	err=e.UserListEmpty()
    }
    return
}

func ShowTotalDues(userList *UserList)(totalDues float64,err error){
    listUser:=*(userList)
    if len(listUser)>0{
    	for name,details:=range listUser{
    		fmt.Println(name,":",details.Dues)
    		totalDues+=details.Dues
    	}
    }else{
    	err=e.UserListEmpty()
    }
    return	
}