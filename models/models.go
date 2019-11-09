package models

import (
	"github.com/golang/protobuf/ptypes/timestamp"
)

//Price of the item or the bid made on the item
type Price float32

//ID represents the id of an auciton, user or a bid
type ID string

//User represents the users making the bid
type User struct {
	UserID   ID
	UserName string
}

//Bid represents a single bid in an Auction
type Bid struct {
	BidID     ID
	AuctionID ID
	UserID    ID
	Price     Price
	Time      timestamp.Timestamp
}

//Auction represents a single auction
type Auction struct {
	AuctionID   ID
	Image       string // image encode in base64
	Description string
	EndTime     timestamp.Timestamp
}

//AuctionList is a list of auctions
type AuctionList []Auction

type Result struct {
	AuctionID ID
	WinnerID  ID
	Price     Price
}

//List of bids
type BidList []Bid
