package main

const selectAuction = "select AuctionID, ItemID, BidPrice, ExpireTime, BidUserID from AuctionItem "
const selectAuctionOrderByRandomLimitX = selectAuction + "order by random() limit ?;"
const selectAuctionWhereItemIDLimitX = selectAuction + "where ItemID = ? limit ?;"
const selectAuctionWhereExpire = selectAuction + " where ExpireTime < ?;"
