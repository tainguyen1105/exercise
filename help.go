package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

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

func (sl shoppingList) toString() string {
	return strings.Join([]string(sl),",")
}

func (sl shoppingList) saveToFile(filename string) error {
	return ioutil.WriteFile(filename,[]byte(sl.toString()),0666)
}