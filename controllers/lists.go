package controllers

import "udhaar/models"

type User = models.User
type Merchant = models.Merchant

type UserList = map[string]User
type MerchantList = map[string]Merchant
