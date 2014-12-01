package main

import (
	"fmt"
	mgo "gopkg.in/mgo.v2"
	"os"
)

func main() {
	uri := os.Getenv("MONGOHQ_URL")
	if uri == "" {
		fmt.Println("no connection string provided")
		os.Exit(1)
	}
	sess, err := mgo.Dial(uri)
	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		os.Exit(1)
	}
	defer sess.Close()
}
