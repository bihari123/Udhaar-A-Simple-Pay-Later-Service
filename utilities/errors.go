package utilities

import "errors"

func InvalidCommand() error {
	return errors.New("Invalid command, please try again")
}

func UserAlreadyPresent() error {
	return errors.New("User already present in the list, try adding a new user")
}

func UserMissing() error {
	return errors.New("User is missing. Can/'t complete the transaction")
}

func UserListEmpty()error{
	return errors.New("User list is empty. Please register some users")
}

func MerchantAlreadyPresent() error {
	return errors.New("Merchant already present in the list, try adding a new merchant")
}

func MerchantMissing() error {
	return errors.New("Merchant is missing. Can/'t update the discount")
}

func CreditLimitExceeded()error{
	return errors.New("Credit Limit Exceeded. Aborting transaction")
}
