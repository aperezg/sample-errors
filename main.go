package main

import (
	"log"

	"github.com/aperezg/sample-errors/behaviour"
	"github.com/aperezg/sample-errors/behaviour113"
	"github.com/aperezg/sample-errors/sentinels"
)

func main() {
	// sentinels
	fetchSomethingWithSentinels()
	fetchSomethingWithSentinelsWithWrap()

	// behaviour
	fetchSomethingBehaviour()
	fetchSomethingWithWrapBehaviour()

	// 113
	fetchSomething113()
	fetchSomethingWithWrap113()
}

func fetchSomethingWithSentinels() {
	err := sentinels.SearchSomething()

	if err == sentinels.ErrEmptyData {
		log.Println(err)
	} else if err != nil {
		log.Fatal("fatal sentinel")
	}
}

func fetchSomethingWithSentinelsWithWrap() {
	err := sentinels.SearchWithWraping()
	if err != nil {
		log.Println(err)
	}
}

func fetchSomethingBehaviour() {
	err := behaviour.SearchSomething()
	if behaviour.IsNotFound(err) {
		log.Println(err)
	} else if err != nil {
		log.Fatal("fatal behaviour")
	}
}

func fetchSomethingWithWrapBehaviour() {
	err := behaviour.SearchWithWraping()
	if behaviour.IsNotFound(err) {
		log.Println(err)
	} else if err != nil {
		log.Fatal("fatal wrap behaviour")
	}
}

func fetchSomething113() {
	err := behaviour113.SearchSomething()
	if behaviour113.IsNotFound(err) {
		log.Println(err)
	} else if err != nil {
		log.Fatal("fatal 113")
	}
}

func fetchSomethingWithWrap113() {
	err := behaviour113.SearchWithWraping()
	if behaviour113.IsNotFound(err) {
		log.Println(err)
	} else if err != nil {
		log.Fatal("fatal wrap 113")
	}
}