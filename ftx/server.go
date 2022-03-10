package ftx

import (
	"fmt"
	"github.com/cloudingcity/go-ftx/ftx"
	"log"
)

const (
	key = ""
	secret = ""
)

func rest(){
	client := ftx.New(
		ftx.WithAuth(key, secret),
	)
	account, err := client.Accounts.GetInformation()
	if err != nil {
		log.Fatal()
	}
	fmt.Printf("%+v", account)
}

func socket()  {
	
}