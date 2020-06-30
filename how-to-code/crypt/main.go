package main

import "fmt"
import "golang.org/x/crypto/bcrypt"

func main() {

	p := `pass123`

	bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(bs)

	loggedIn := `pass123`

	err = bcrypt.CompareHashAndPassword(bs, []byte(loggedIn))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("LoggedIn Succesfull")

}