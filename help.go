package main

import (
	"fmt"
	"io/ioutil"
	"os"
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

func newShoppingListFromFile(filename string) shoppingList {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:",err)
		os.Exit(1)
	}
	s := strings.Split(string(bs),",")
	return shoppingList(s)
}