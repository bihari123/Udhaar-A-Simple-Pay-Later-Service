package controllers

import (
	"github.com/google/uuid"
	e "udhaar/utilities"
    "github.com/spf13/cast"
 
    )


func Transaction(command []string,userList *UserList,merchantList *MerchantList)(txnID uuid.UUID,err error){
    
    userName := command[2]
	merchantName := command[3]
	
	
	listUser := *(userList)
	_, found := listUser[userName]
	if !found {
		err= e.UserMissing()
		return
	}
     user:=listUser[userName]

	listMerchant := *(merchantList)
	_, found = listMerchant[merchantName]
	if !found {
		err= e.MerchantMissing()
		return
	}
	merchant:=listMerchant[merchantName]

    Amount := cast.ToFloat64(command[4])
    //after discount
    discount:=float64(merchant.DiscountOffered*Amount/100)
    
    txnAmount := float64(Amount - discount)
  
	if user.Dues + txnAmount > user.CreditLimit{
		err=e.CreditLimitExceeded()
		return
	}
    
  
	user.Dues+=txnAmount
	merchant.TotalDiscount +=discount
	txnID,err = uuid.NewUUID()

	user.TransactionID =append(user.TransactionID,txnID)
	merchant.TransactionID =append(merchant.TransactionID,txnID)

    listMerchant[merchantName] =merchant
    listUser[userName]=user
    
    return

}


func PayBack(command []string,userList *UserList)(dues float64,err error){
	userName:= command[1]
	amount:=cast.ToFloat64(command[2])
	listUser:=*(userList)
	_, found := listUser[userName]
	if !found {
		err= e.UserMissing()
		return
	}
     user:=listUser[userName]

     user.Dues -=amount
     txnID,err := uuid.NewUUID()
	 user.TransactionID =append(user.TransactionID,txnID)
     
     dues =user.Dues
    
     listUser[userName]=user

    return
}