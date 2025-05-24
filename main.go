package main

func main() {
	// lists := newShoppingList()
	// lists.print()
	// lists.saveToFile("my_list")
	lists := newShoppingListFromFile("my_list")
	lists.print()
}