package main

import "fmt"

type shoppingList []string

func newShoppingList() shoppingList {
	list := shoppingList{"Milk","Bread","Eggs"}
	return list
}

func (sl shoppingList) print() {
	for i,list := range sl {
		fmt.Println(i,list)
	}
}