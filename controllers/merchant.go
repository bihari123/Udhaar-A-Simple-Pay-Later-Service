package controllers

import (
    "strings"
	e "udhaar/utilities"
    "github.com/spf13/cast"
  
)

func AddMerchant(command []string, merchantList *MerchantList) (newMerchant Merchant,err error) {
	name := command[2]
	var email string
	var discount float64
	if len(command)>4{
		email= command[3]
	    discount = cast.ToFloat64((strings.Split(command[4],"%"))[0])		
	}else{
	    discount = cast.ToFloat64((strings.Split(command[3],"%"))[0])	
	}
	
	list := *(merchantList)
	_, found := list[name]
	if found {
		err= e.MerchantAlreadyPresent()
		return
	}
	newMerchant =Merchant{
			Name:            name,
			Email: email,
			DiscountOffered: discount,
		}
		
	list[name] = newMerchant
	return 
}

func UpdateDiscount(command []string,merchantList *MerchantList)(newDiscount float64, err error){
    var merchantCopy Merchant
    
    merchantName:=command[2]
    list := *(merchantList)
	_, found := list[merchantName]
	if !found {
		err= e.MerchantMissing()
		return
	}

	newDiscount= cast.ToFloat64((strings.Split(command[3],"%"))[0])
 
    merchantCopy = list[merchantName]
	merchantCopy.DiscountOffered=newDiscount

	list[merchantName]=merchantCopy
    return
}

func ShowMerchantDiscount(command []string,merchantList *MerchantList)(discount float64,err error){
	merchantName:=command[2]
	list := *(merchantList)
	_, found := list[merchantName]
	if !found {
		err= e.MerchantMissing()
		return
	}
	discount = list[merchantName].TotalDiscount
	return

}
