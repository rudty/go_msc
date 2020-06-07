package main

const selectAuction = "select AuctionID, ItemID, BidPrice, ExpireTime, BidUserID from AuctionItem "
const selectAuctionOrderByRandomLimitX = selectAuction + "order by random() limit ?;"
const selectAuctionWhereItemIDLimitX = selectAuction + "where ItemID = ? limit ?;"
const selectAuctionWhereExpire = selectAuction + " where ExpireTime < ?;"

const selectBidUserIDByAuctionID = `select BidUserID from AuctionItem where AuctionID = ?;`

const updateBidPriceAndBidUserIDWhereAuctionIDBidPrice = `update AuctionItem set BidPrice = ?, BidUserID = ? where AuctionID = ? and BidPrice < ?;`
