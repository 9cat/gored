package test

// Copyright (c) 2015-2019 Bitontop Technologies Inc.
// Distributed under the MIT software license, see the accompanying
// file COPYING or http://www.opensource.org/licenses/mit-license.php.

import (
	"log"
	"testing"

	"../coin"
	"../exchange"
	"../exchange/bittrex"
	"../pair"
	"./conf"
)

/********************Public API********************/
func Test_Bittrex(t *testing.T) {
	e := InitBittrex()

	pair := e.GetPair("BTC|ETH")

	Test_Coins(e)
	Test_Pairs(e)
	Test_Pair(e, pair)
	Test_Orderbook(e, pair)
	// Test_Trading(e, pair, 0.00000001, 100)
	// Test_Withdraw(e, pair.Base, 1, "ADDRESS")
	Test_ConstraintFetch(e, pair)
	Test_Constraint(e, pair)
}

func InitBittrex() exchange.Exchange {
	coin.Init()
	pair.Init()
	config := &exchange.Config{}
	config.Source = exchange.EXCHANGE_API
	conf.Exchange(exchange.BITTREX, config)

	ex := bittrex.CreateBittrex(config)
	log.Printf("Initial [ %v ] ", ex.GetName())

	config = nil
	return ex
}
