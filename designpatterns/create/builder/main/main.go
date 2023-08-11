package main

import builder2 "go-learn/designpatterns/create/builder"

func main() {
	builder := &builder2.BuilderA{}
	director := builder2.NewDirector(builder)
	director.Construct()

}
