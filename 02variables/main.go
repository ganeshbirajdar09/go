package main

import (
	"fmt"
)

// jwtToken := "gahsasfasags"  --cant do this outside any method

const LoginToken string = "ghsdhgadghas" //PUBLIC (recommened=first letter of var name should be capital)

func main() {
	var username string = "ganesh"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float64 = 255.9416598316848941321984
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	//default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	var anotherString string
	fmt.Println(anotherString)
	fmt.Printf("Variable is of type: %T \n", anotherString)

	var anotherBool bool
	fmt.Println(anotherBool)
	fmt.Printf("Variable is of type: %T \n", anotherBool)

	//implicit type
	var website = "learncodeonline.in"
	fmt.Println(website)

	//no var style
	noOfUsers := 300000
	fmt.Println(noOfUsers)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)

}
