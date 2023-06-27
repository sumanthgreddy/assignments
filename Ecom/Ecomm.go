package main

import (
	"fmt"
	"regexp"
)

type Inv_Item struct {
	ItemName  string
	ItemCode  string
	ItemPrice float64
	Quantity  int
}

type CartItem struct {
	Inv_Item     //InventoryItem
	Q_Cart   int //QuantityInCart
}

type Inventory struct {
	Items []Inv_Item
}

type Cart struct {
	Items []CartItem
}

func main() {
	inventory := InitializeInventory() // Initializing Inventory
	cart := Cart{}
	itemInventory := make(map[string]int) //Create map for ItemInventory match from Item code to Quantity

	for _, item := range inventory.Items {
		itemInventory[item.ItemCode] = item.Quantity
	}

	for {
		fmt.Print("\n------------- Welcome to the E-commerce Application ---------------\n")
		fmt.Println("Please select an option:")
		fmt.Println("1. View Inventory")
		fmt.Println("2. Add Item to Cart")
		fmt.Println("3. Remove Item from Cart")
		fmt.Println("4. View Cart")
		fmt.Println("5. Checkout and Calculate Total Bill")
		fmt.Println("6. Search")
		fmt.Println("7. Exit")
		fmt.Print("----------------------**********************-----------------------\n")

		var choice int
		_, err := fmt.Scanf("%d", &choice)
		if err != nil {
			fmt.Println("Invalid input. Please try again.")
			fmt.Print("\n----------------------**********************-----------------------\n")
			continue
		}

		switch choice {
		case 1:
			ViewInventory(inventory)
		case 2:
			AddItemToCart(&cart, inventory, itemInventory)
		case 3:
			RemoveItemFromCart(&cart)
		case 4:
			ViewCart(cart)
		case 5:
			Checkout(cart)
			return
		case 6:
			Search(inventory, itemInventory, &cart)
		case 7:
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
			fmt.Print("\n----------------------**********************-----------------------\n")
		}
	}
}

func InitializeInventory() Inventory {
	items := []Inv_Item{
		{ItemName: "Pencil", ItemCode: "p1", ItemPrice: 3.0, Quantity: 50},
		{ItemName: "Pen", ItemCode: "p2", ItemPrice: 7.0, Quantity: 20},
		{ItemName: "Book", ItemCode: "b1", ItemPrice: 1.0, Quantity: 14},
		{ItemName: "Ruler", ItemCode: "r1", ItemPrice: 2.0, Quantity: 58},
		{ItemName: "Bottle", ItemCode: "b2", ItemPrice: 10.0, Quantity: 67},
		{ItemName: "Table", ItemCode: "t1", ItemPrice: 40.0, Quantity: 89},
		{ItemName: "Chair", ItemCode: "c1", ItemPrice: 35.0, Quantity: 20},
		{ItemName: "Wall Stand", ItemCode: "s1", ItemPrice: 37.0, Quantity: 19},
		{ItemName: "Charging Cable", ItemCode: "c2", ItemPrice: 16.0, Quantity: 20},
		{ItemName: "Earphones", ItemCode: "e1", ItemPrice: 12.0, Quantity: 39},
	}

	return Inventory{Items: items}
}

func ViewInventory(inventory Inventory) {
	fmt.Print("\n----------------------**********************-----------------------\n")
	fmt.Println("Inventory:")
	for _, item := range inventory.Items {
		fmt.Printf("Item: %s (Code: %s) - Price: $%.2f - Quantity: %d\n", item.ItemName, item.ItemCode, item.ItemPrice, item.Quantity)
	}
	fmt.Print("\n----------------------**********************-----------------------\n")
}

func AddItemToCart(cart *Cart, inventory Inventory, itemInventory map[string]int) {
	ViewInventory(inventory) // call inventory func

	var itemCode string
	fmt.Print("Enter the item code to add to cart: ")
	_, err := fmt.Scanf("%s", &itemCode)
	if err != nil {
		fmt.Println("Invalid item code. Please try again.")
		fmt.Print("\n----------------------**********************-----------------------\n")
		return
	}

	item := findItemByCode(inventory.Items, itemCode)
	if item == nil {
		fmt.Println("Item not found in the inventory.")
		fmt.Print("\n----------------------**********************-----------------------\n")
		return
	}

	fmt.Printf("Adding item to cart: %s\n", item.ItemName)

	var quantity int
	fmt.Print("Enter the quantity to add to cart: ")
	_, err = fmt.Scanf("%d", &quantity)
	if err != nil {
		fmt.Println("Invalid quantity. Please try again.")
		return
	}
// check the Item Quantity
	if quantity > item.Quantity {
		fmt.Println("Insufficient quantity in the inventory.")
		return
	}

	if quantity > itemInventory[itemCode] {
		fmt.Println("Insufficient quantity in the inventory.")
		return
	}

	cartItem := findCartItemByCode(cart.Items, itemCode)
	if cartItem != nil {
		if quantity+cartItem.Q_Cart > item.Quantity {
			fmt.Println("Insufficient quantity in the inventory.")
			return
		}
		cartItem.Q_Cart += quantity
	} else {//Change Item quantity as per requirement
		cart.Items = append(cart.Items, CartItem{
			Inv_Item: *item,
			Q_Cart:   quantity,
		})
	}

	item.Quantity -= quantity
	itemInventory[itemCode] -= quantity
	fmt.Println("Item added to cart.")
	fmt.Print("\n----------------------**********************-----------------------\n")
}

// Search the Item by the Item code
func Search(inventory Inventory, itemInventory map[string]int, cart *Cart) {
	var Sip string
	fmt.Print("\n----------------------**********************-----------------------\n")
	fmt.Println("To search for an item, enter a Item Code:")
	fmt.Scanf("%s", &Sip)
	fmt.Printf("\nSearching for term: %s\n", Sip)

	for _, item := range inventory.Items {
		match, _ := regexp.MatchString(fmt.Sprintf("\\b%s\\b", Sip), item.ItemCode)
		if match {
			fmt.Println("Item found:", item)
			fmt.Println("Press 1 to add to cart, press 2 to search for another item, or press 3 to exit")
			var input int
			fmt.Scanf("%d", &input)
			switch input {
			case 1: // Item Found, adding to cart futionality 
				fmt.Printf("Adding item to cart: %s\n", item.ItemName)
				var quantity int
				fmt.Print("Enter the quantity to add to cart: ")
				_, err := fmt.Scanf("%d", &quantity)
				if err != nil {
					fmt.Println("Invalid quantity. Please try again.")
					return
				}

				if quantity > item.Quantity {
					fmt.Println("Insufficient quantity in the inventory.")
					return
				}

				if quantity > itemInventory[item.ItemCode] {
					fmt.Println("Insufficient quantity in the inventory.")
					return
				}

				cartItem := findCartItemByCode(cart.Items, item.ItemCode)
				if cartItem != nil {
					if quantity+cartItem.Q_Cart > item.Quantity {
						fmt.Println("Insufficient quantity in the inventory.")
						return
					}
					cartItem.Q_Cart += quantity
				} else {
					cart.Items = append(cart.Items, CartItem{Inv_Item: item, Q_Cart: quantity})
				}
				item.Quantity -= quantity
				itemInventory[item.ItemCode] -= quantity
				fmt.Println("Item added to cart.")
				fmt.Println()

			case 2:
				Search(inventory, itemInventory, cart)
			default:
				return
			}
			return
		}
	}
	fmt.Println("Item not found.")
}

func RemoveItemFromCart(cart *Cart) {
	ViewCart(*cart)

	var itemCode string
	fmt.Print("\n----------------------**********************-----------------------\n")
	fmt.Print("Enter the item code to remove from cart: ")
	_, err := fmt.Scanf("%s", &itemCode)
	if err != nil {
		fmt.Println("Invalid item code. Please try again.")
		fmt.Print("\n----------------------**********************-----------------------\n")
		return
	}

	cartItem := findCartItemByCode(cart.Items, itemCode)
	if cartItem == nil {
		fmt.Println("Item not found in the cart.")
		fmt.Print("\n----------------------**********************-----------------------\n")
		return
	}

	fmt.Printf("Removing item from cart: %s\n", cartItem.ItemName)

	var quantity int
	fmt.Print("Enter the quantity to remove from cart: ")
	_, err = fmt.Scanf("%d", &quantity)
	if err != nil {
		fmt.Println("Invalid quantity. Please try again.")
		return
	}

	if quantity > cartItem.Q_Cart {
		fmt.Println("Insufficient quantity in the cart.")
		return
	}

	cartItem.Q_Cart -= quantity

	if cartItem.Q_Cart == 0 {
		removeCartItem(cart, itemCode)
	}

	fmt.Println("Item removed from cart.")
	fmt.Println()
}

func removeCartItem(cart *Cart, itemCode string) {
	for i, cartItem := range cart.Items {
		if cartItem.ItemCode == itemCode {
			// Remove the item from the slice
			cart.Items = append(cart.Items[:i], cart.Items[i+1:]...)
			break
		}
	}
}

func findItemByCode(items []Inv_Item, itemCode string) *Inv_Item {
	for _, item := range items {
		if item.ItemCode == itemCode {
			return &item
		}
	}
	return nil
}

func findCartItemByCode(items []CartItem, itemCode string) *CartItem {
	for _, item := range items {
		if item.ItemCode == itemCode {
			return &item
		}
	}
	return nil
}

func ViewCart(cart Cart) {
	if len(cart.Items) == 0 {
		fmt.Println("Cart is empty.")
		fmt.Println("Please add the items to cart to view")
		main()

	} else {
		fmt.Println("Cart:")
		for _, item := range cart.Items {
			fmt.Printf("Item: %s (Code: %s) - Price: $%.2f - Quantity: %d\n", item.ItemName, item.ItemCode, item.ItemPrice, item.Q_Cart)
		}
	}
	fmt.Println()
}

func Checkout(cart Cart) {
	ViewCart(cart)

	totalBill := 0.0
	for _, item := range cart.Items {
		itemTotal := item.ItemPrice * float64(item.Q_Cart)
		totalBill += itemTotal
		fmt.Printf("Item: %s - Quantity: %d - Price: $%.2f - Total: $%.2f\n", item.ItemName, item.Q_Cart, item.ItemPrice, itemTotal)
	}

	fmt.Printf("Total Bill: $%.2f\n", totalBill)
	fmt.Println("Thank you for shopping!")
	fmt.Print("\n\n------------------------- CHECKED OUT ---------------------------\n\n")
}
