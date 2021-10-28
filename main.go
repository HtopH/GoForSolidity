package main

import "SecondTest/app"

func main() {

	AuctionContractFunc()
}

func AuctionContractFunc()  {
	err := app.AuctionContractService.InitAuctionContract()
	if err != nil {
		return
	}
	defer app.AuctionContractService.Client.Close()

	_ = app.AuctionContractService.ReadBlockLog()
}
