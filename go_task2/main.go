package main

import "fmt"

func usingSwitch(amount int){
var hundred,fifty,twenty,ten,five,two,one int
switch amount>0 {
	case amount>=100:
		hundred = amount/100
	    amount = amount%100
		fallthrough
	case amount>=50:
		fifty = amount/50
	    amount = amount%50
		fallthrough
	case amount>=20:
		twenty = amount/20
	    amount = amount%20
		fallthrough
	case amount>=10:
		ten = amount/10
	    amount = amount%10
		fallthrough
	case amount>=5:
		five = amount/5
	    amount = amount%5
		fallthrough
	case two>=2:
		two = amount/2
	    amount = amount%2
		fallthrough
	case one>=1:
		one = amount/1
		amount = amount%1
	default:
		fmt.Println("invalid")
	}
	fmt.Printf("%d Note(s) of 100.00\n",hundred)
	fmt.Printf("%d Note(s) of 50.00\n",fifty)
	fmt.Printf("%d Note(s) of 20.00\n",twenty)
	fmt.Printf("%d Note(s) of 10.00\n",ten)
	fmt.Printf("%d Note(s) of 5.00\n",five)
	fmt.Printf("%d Note(s) of 2.00\n",two)
	fmt.Printf("%d Note(s) of 1.00\n",one)
}

func usingifElse(amount int){
	var hundred,fifty,twenty,ten,five,two,one int
	for amount>0{
		if amount>=100{
			hundred = amount/100
			amount = amount%100
		}else if amount>=50{
			fifty = amount/50
			amount = amount%50
		}else if amount>=20{
			twenty = amount/20
			amount = amount%20
		}else if amount>=10{
			ten = amount/10
			amount = amount%10
		}else if amount>=5{
			five = amount/5
			amount = amount%5
		}else if amount>=2{
			two = amount/2
			amount = amount%2
		}else if amount>=1{
			one = amount/1
			amount = amount%1
		}else{
			fmt.Println("invalid")
		}
	}
	fmt.Printf("%d Note(s) of 100.00\n",hundred)
	fmt.Printf("%d Note(s) of 50.00\n",fifty)
	fmt.Printf("%d Note(s) of 20.00\n",twenty)
	fmt.Printf("%d Note(s) of 10.00\n",ten)
	fmt.Printf("%d Note(s) of 5.00\n",five)
	fmt.Printf("%d Note(s) of 2.00\n",two)
	fmt.Printf("%d Note(s) of 1.00\n",one)
}

func main() {
	var amount int
	fmt.Print("Enter the Amount:")
	if _, err := fmt.Scanln(&amount); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("There are:")
	usingifElse(amount)
}