package main

import "fmt"

type shoppingList []string

func newShoppingList() shoppingList {
	list := shoppingList{"Milk","Bread","Eggs"}
	return list
}