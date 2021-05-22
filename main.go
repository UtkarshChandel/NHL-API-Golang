package main

import (
	"io"
	"log"
	"os"
	"time"
	"github.com/utkarshchandel/NHL-API-Golang/nhlAPI"
)

func main() {
	now := time.Now()
	
	rosterFie,err := os.OpenFile("roster.txt",os.O_RDWR | os.O_CREATE | os.O_APPEND,0666)

	if err != nil {
		log.Fatalf("error opening the file roster.txt : %v",err);


	}
	//close the file after the function ends and after handling the error
	defer rosterFie.Close()

	wrt := io.MultiWriter(os.Stdout, rosterFie)
	log.SetOutput(wrt)

	teams, err := nhlAPI.GetAllTeams()

	if err != nil {
		log.Fatalf("error while geting all teams : %v",err)
	}

	for _,team := range teams{
		log.Println("---------------")
		log.Printf("Name : %s ",team.Name)
		log.Println("----------------")
	}

	log.Printf("took %v ", time.Now().Sub(now).String())

}
