package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Print("***************************************\n")
	fmt.Print("       Welcome To BHOLENATH CAFE\n")
	fmt.Print("***************************************\n")
	fmt.Print("C: Coffee, Rs: 40/-\nD: Dosa, Rs: 80/-\nT: Tomato Soup, Rs: 20/-\nI : Idli , Rs: 20/-\nV : Vada, Rs: 25/-\nB: Banana Shake, Rs: 30/-\nP: Paneer Pakoda, Rs: 30/-\nM: Manchurian, Rs: 90/-\nH: Hakka Noodle, Rs: 70/-\nF: French Fries, Rs: 30/-\nJ: Jalebi, Rs: 30/-\nL: Lemonade, Rs: 15/-\nS: Spring Roll, Rs: 20/-\n")
	var n int
	var code, c string
	var total []int
	var bill int = 0
	price := map[string]int{
		"C": 40, "D": 80, "T": 20, "I": 20, "V": 25, "B": 30, "P": 30, "M": 90, "H": 70, "F": 30, "J": 30, "L": 15, "S": 20,
	}
x:
	for true {
		fmt.Print("C: Coffee, Rs: 40/-\nD: Dosa, Rs: 80/-\nT: Tomato Soup, Rs: 20/-\nI: Idli , Rs: 20/-\nV: Vada, Rs: 25/-\nB: Banana Shake, Rs: 30/-\nP: Paneer Pakoda, Rs: 30/-\nM: Manchurian, Rs: 90/-\nH: Hakka Noodle, Rs: 70/-\nF: French Fries, Rs: 30/-\nJ: Jalebi, Rs: 30/-\nL: Lemonade, Rs: 15/-\nS: Spring Roll, Rs: 20/-\n")
		fmt.Print("\nPlace the order: ")
		fmt.Scan(&code)
		if code == "END" || code == "end" {
			var total_income int

			for i := 0; i < len(total); i++ {
				total_income = total[i] + total_income
			}
			fmt.Println("**************************")
			fmt.Println("      BHOLENATH CAFE")
			fmt.Println("**************************")
			fmt.Println(" Total Income of the day:", total_income)
			fmt.Println("******END OF THE DAY******")
			os.Exit(0)
		}
		fmt.Scan(&n)
		bill += n * price[code]
	z:
		fmt.Println("Want to add anything else(y/n): ")
		fmt.Scan(&c)
		if c == "y" || c == "Y" {
			goto x
		} else if c == "n" || c == "N" {
			goto y
		} else {
			fmt.Println("Wrong Input, Please Try Again!")
			goto z
		}
	y:
		fmt.Println("Your total bill is:", bill)

		total = append(total, bill)
		bill = 0
	}
}
