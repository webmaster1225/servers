/*
   @dev: Logan (Nam) Nguyen
   @course: SUNY Oswego - CSC 495 - Capstone
   @instructor: Professor Bastian Tenbergen
   @version: 1.0
*/

// @package
package controllers

import (
	"Syns/servers/syns-users-ms/models"
	"Syns/servers/syns-users-ms/utils"
	"os"

	"github.com/gin-gonic/gin"
)

// Constants
var (
	ALCHEMY_BASE_URL = "https://polygon-mumbai.g.alchemy.com/nft/v2/"
	MORALIS_BASE_URL = "https://deep-index.moralis.io/api/v2/nft/"
)

// @route `GET/get-all-syns-tokens/:asset-contract`
//
// @dev Get all Syns tokens from blockchain
//
// @honor Moralis API
//
// @param gc *gin.Context
func GetAllSynsTokens(gc *gin.Context) {

	assetContract := gc.Param("asset-contract")
  
	url := MORALIS_BASE_URL+assetContract+"?chain=mumbai&format=decimal&normalizeMetadata=true&media_items=false"
  
	// prepare response object
	resObject := &models.MoralisNFTResponse{}

	// process http request
	result := utils.DoHttp(url, "X-API-Key", os.Getenv("MORALIS_API_KEY"), resObject)

	// convert result back to the correct model
	NFTRes, _ := result.(*models.MoralisNFTResponse)
  
	  // return this to client
	gc.JSON(200, gin.H{"nfts": NFTRes.Result})
  }

// @route `GET/get-nfts-owned-by/:owner-addr/:asset-contract`
// 
// @dev Get all Syns tokens owned by an owner address
//
// @honor Alchemy API
//
// @param gc *gin.Context
func GetSynsTokensOwnedBy(gc *gin.Context) {
	// prepare params
	ownerAddr := gc.Param("owner-addr")
	assetContract := gc.Param("asset-contract")

	// prepare url
	url := ALCHEMY_BASE_URL+os.Getenv("ALCHEMY_API_KEY")+"/getNFTs?owner="+ownerAddr+"&contractAddresses[]="+assetContract+"&withMetadata=true&pageSize=100"

	// prepare response object
	resObject := &models.AlchemyNFTsByOwnerResponse{}

	// process http request
	result := utils.DoHttp(url, "", "", resObject)

	// convert result back to the correct model
	NFTRes, _ := result.(*models.AlchemyNFTsByOwnerResponse)

	// return this to client
	gc.JSON(200, gin.H{"nfts": NFTRes.OwnedNfts})
}


// @route `GET/get-token-metadata/:asset-contract/:token-id/:token-type`
// 
// @dev Get token metadata based on asset contract and tokenId
//
// @honor Alchemy API
//
// @param gc *gin.Context
func GetTokenMetadata(gc *gin.Context) {
	// // prepare params
	assetContract := gc.Param("asset-contract")
	tokenId := gc.Param("token-id")
	tokenType := gc.Param("token-type")

	// prepare url
	url := ALCHEMY_BASE_URL+os.Getenv("ALCHEMY_API_KEY")+"/getNFTMetadata?contractAddress="+assetContract+"&tokenId="+tokenId+"&tokenType="+tokenType+"&refreshCache=false"

	// prepare response object
	resObject := &models.AlchemyNftMetadataResponse{}

	// process http request
	result := utils.DoHttp(url, "", "", resObject)

	// convert result back to the correct model
	NFTRes, _ := result.(*models.AlchemyNftMetadataResponse)

	// return this to client
	gc.JSON(200,NFTRes)
}