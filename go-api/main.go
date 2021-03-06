package main

import (
	"fmt"
	"github.com/ccdle12/Blocksage/go-api/injector"
	"log"
	"net/http"
)

func main() {
	// TODO: Throwing crawler here, will be separated later
	// TODO: Make crawler run on a separate go routine
	inj := injector.DependencyInjector{}

	go func() {
		bitcoinCrawler := inj.InjectBitcoinCrawler()
		bitcoinCrawler.Start()
	}()

	API := inj.InjectMainnetAPI()
	router := inj.InjectRouter()

	router.HandleFunc("/api/network-info", API.NetworkInfo).Methods("GET")
	router.HandleFunc("/api/blocks/{blockhash}", API.Blocks).Methods("GET")
	router.HandleFunc("/api/txs/{txhash}", API.Transactions).Methods("GET")

	fmt.Println("Serving on Port 8548...")
	log.Fatal(http.ListenAndServe(":8548", router))
}
