package main

import(
	"fmt"
)

func main() {

	var ip int
	var cID int
	var Aall string
	var cDob string
	var data int 

	var custN = map[int]string{
		0001: "Customer 1" ,
		0002: "Customer 2" ,
		0003: "Customer 3" ,
		0004: "Customer 4" ,
		0005: "Customer 5" ,
		9999: "Admin" ,
	}
	var custDob = map[int]string{
		//dob in mmddyyyy is given as string
		0001: "01021997",
		0002: "09091995",
		0003: "09111993",
		0004: "12231994",
		0005: "05181994",
	}
	var custBal = map[int]int {
		0001: 3000,
		0002: 2234,
		0003: 3343,
		0004: 242,
		0005: 933,
	}
	var custAdd = map[int]string{
		0001: "Florida",
		0002: "Newyork",
		0003: "Arizona",
		0004: "Memphis",
		0005: "Kansas",
		}

for {
fmt.Println("Welcome to banking application")
fmt.Println("Please follow the below list to proceed. Press the respective numbers \n 1. Login \n 2. New Customer \n 3. Exit")
fmt.Scanf("%d", &ip)
if ip == 1{

	fmt.Println("Please give you 4 digit customer ID")
	fmt.Scanf("%d", &cID)
		if cID == 9999 { 
			//Admin Access
			fmt.Println("Admin Access!!!\n \tAdmin Access!!!\n \t \tAdmin Access!!!")
			fmt.Println(" To Print Customer Name Press n\n Customer Account Balance Press b \n Customer Address Press l\n Customer Date of Birth Press d\n To print All data Press a")
			fmt.Scanf("%s", &Aall)
				if Aall == "a" {
					adminN(custN)
					adminBal(custBal)
					adminAdd(custAdd)
					admindob(custDob)
					}else if Aall == "n" {
					adminN(custN)
					}else if Aall == "b" {
					adminBal(custBal)
					}else if Aall == "l" {
					adminAdd(custAdd)
					}else if Aall == "d" {
					admindob(custDob)
					}else {
					return
					} 
			}else{
				fmt.Println("Please enter you date of birth in mmddyyyy format to login")
				fmt.Scanf("%s", &cDob)
					if _, exists := custDob[cID]; exists && custDob[cID] == cDob {
						
						fmt.Printf("Thank you for being our customer %v with Dob %v \n", cID, cDob)
						fmt.Printf("Please press the details required:\n 1. Customer Information \n 2. Customer Bank Balance\n 3. Withdraw Money\n 4. Deposit Money\n 5. Exit\n")
						for {
						fmt.Scanf("%d", &data)
							switch data {
								case 1:
									fmt.Println("The information is below")
									fmt.Println(custN[cID])
									fmt.Println(custAdd[cID])
									fmt.Println(custDob[cID])
									fmt.Print("Press \n1. Customer Information \n 2. Customer Bank Balance\n 3. Withdraw Money\n 4. Deposit Money\n 5. Exit\n")
								
								case 2:
									fmt.Println("The Amount in the bank for customer is below:")
									fmt.Println(custN[cID])
									fmt.Println(custBal[cID])
									fmt.Print("Press \n1. Customer Information \n 2. Customer Bank Balance\n 3. Withdraw Money\n 4. Deposit Money\n 5. Exit\n")
									
								case 3:
									fmt.Println("To withdraw Money")
									wdbal(custBal)
									fmt.Print("Press \n1. Customer Information \n 2. Customer Bank Balance\n 3. Withdraw Money\n 4. Deposit Money\n 5. Exit\n")
									
								case 4:
									fmt.Println("To deposit Money")
									dpbal(custBal)
									fmt.Print("Press \n1. Customer Information \n 2. Customer Bank Balance\n 3. Withdraw Money\n 4. Deposit Money\n 5. Exit\n")
									
								default:
									fmt.Println("Thank you for banking!!!")
									return
								}


							}	
						 
						}else{
							fmt.Printf("Key %v and %v does not match in the cutsomer data", cID, cDob)
						}

	
	}
} else if ip == 2 {

	// adding new customer 

	leng := len(custDob)
	var NewCN string // New Customer Name
	var NewCADD string // New customer Address
	var NewCDP int //New customer Deposit
	var NewCDB string // New Customer Dob

	
	fmt.Println("enter the customer name")
	fmt.Scan(&NewCN)
	leng = leng + 1
	custN[leng] = NewCN

	fmt.Println("enter the Address name")
	fmt.Scan(&NewCADD)
	custAdd[leng] = NewCADD
	
	fmt.Println("enter the Deposit amount")
	fmt.Scan(&NewCDP)
	custBal[leng] = NewCDP

	fmt.Println("enter the Date of birth")
	fmt.Scan(&NewCDB)
	custDob[leng] = NewCDB

	fmt.Println("Thank you for being part of this bank \n Please find the below details")
	fmt.Print("Customer Name: ")
	fmt.Println(custN[leng])
	fmt.Print("Customer Date of Birth: ")
	fmt.Println(custDob[leng])
	fmt.Print("Customer Bank Balance: ")
	fmt.Println(custBal[leng])
	fmt.Print("Customer Address: ")
	fmt.Println(custAdd[leng])
	fmt.Print("Customer Account Number: ")
	fmt.Println(leng)
	fmt.Print("Please proceed with Login \n \n")


}else if ip == 3 {
fmt.Println("Please visit again. Thank You!!!")
	return
}
}
	}

	func wdbal(custBal map[int]int){
		
		var Amt int
		var cn int

		fmt.Println("Please reconfirm your Customer Account no")
		fmt.Scan(&cn)

		fmt.Println("Enter the Amount")
		fmt.Scan(&Amt)

		if balance, ok := custBal[cn]; ok {
			if balance >= Amt {
				// Withdraw the amount
				custBal[cn] -= Amt
				fmt.Printf("Withdrawal of amount %d successful for customer %d\n", Amt, cn)
			} else {
				fmt.Printf("Insufficient balance for customer %d\n", cn)
			}
		} else {
			fmt.Printf("Customer %d not found\n", cn)
		}

	}
func adminN(custN map[int]string) {
	fmt.Println(custN)
}
func admindob(custDob map[int]string) {
	fmt.Println(custDob)
}
func adminBal(custBal map[int]int){
	fmt.Println(custBal)
}

func adminAdd(custAdd map[int]string){
	fmt.Println(custAdd)
}

func dpbal(custBal map[int]int){
		
	var Amt int
	var cn int

	fmt.Println("Please reconfirm your Customer Account no")
	fmt.Scan(&cn)

	fmt.Println("Enter the Amount")
	fmt.Scan(&Amt)

	if _, ok := custBal[cn]; ok {
			custBal[cn] += Amt
			fmt.Printf("Deposit of amount %d successful for customer %d\n", Amt, cn)
		} else {
		fmt.Printf("Customer %d not found\n", cn)
	}


}

// Test Case 1
// Login with correct ID
// check 1,2,5
// check Withdrawal in limit
// check withdrawal limit cross
// check deposit
// New customer
// check the new customer entry
