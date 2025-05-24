package main

func main() {
	lists := newShoppingList()
	// lists.print()
	lists.saveToFile("my_list")
}