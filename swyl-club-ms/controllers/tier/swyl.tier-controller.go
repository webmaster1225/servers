/*
   @dev: Logan (Nam) Nguyen
   @course: SUNY Oswego - CSC 495 - Capstone
   @instructor: Professor Bastian Tenbergen
   @version: 1.0
*/

// @package
package controllers

// @import
import (
	dao "Swyl/servers/swyl-club-ms/dao/tier"
	"Swyl/servers/swyl-club-ms/models"
	"Swyl/servers/swyl-club-ms/utils"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// global vars
var validate = validator.New()

// @notice Root struct for other methods in controller
type TierController struct {
   TierDao dao.TierDao
}

// @dev Constructor
func TierControllerConstructor(tierDao dao.TierDao) *TierController {
   return &TierController{
      TierDao: tierDao,
   }
}


// @notice Method of TierController struct
// 
// @route `POST/create-tier`
// 
// @dev Lets a club owner create a tier to internal database
// 
// @param gc *gin.Context
func (tc *TierController) CreateTier(gc *gin.Context) {
   // declare params
   param := &models.Tier{}

   // bind json post data to param
   if err := gc.ShouldBindJSON(param); err != nil {
      gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return;
   }

   // validate struct models.Tier
   if err := validate.Struct(param); err != nil {gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"erorr": err.Error()}); return;}

   // extra validation on param.club_owner to match ETH Crypto Wallet address convention
   matched, err := utils.TestEthAddress(param.Club_owner)
   if err != nil {gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "!REGEX - cannot test club_owner using regex"}); return;}
   if !matched {gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "!ETH_ADDRESS - club_owner is not an ETH crypto wallet address"}); return;}

   // invoke TierDao.CreateTier() api
   if err := tc.TierDao.CreateTier(param); err !=nil {
      gc.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"erorr": err.Error()})
   }

   // http reponse
   gc.JSON(200, gin.H{"msg": "Tier is sucessfully created"})
}


// @notice Method of TierController struct
// 
// @route `@GET/get-tier-at/:tier_id`
// 
// @dev Gets a Tier at tierId and clubOwner
// 
// @param gc *gin.Context
func (tc *TierController) GetTierAt(gc *gin.Context) {
   // get tierId from param
   tierId := gc.Param("tier_id")

   // sanitize tierId
   matched, err := regexp.MatchString(`^[a-zA-Z0-9]*$`, tierId)
   if err != nil {gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "!REGEX - cannot test club_owner using regex"}); return;}
   if !matched {gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "!CLUBID - clubId is not valid"}); return;}

   
   // invoke TierDao.GetTierAt
   tier, err := tc.TierDao.GetTierAt(&tierId)
   if err != nil {gc.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})}
   
   // http reponse
   gc.JSON(200, tier)
}


// @notice Method of TierController struct
// 
// @route `GET//get-all-tiers-owned-by/:clubOwner`
// 
// @dev Gets all tiers owned by clubOwner
// 
// @param gc *gin.Context
func (tc *TierController) GetAllTiersOwnedBy(gc *gin.Context) {
   gc.JSON(200, gin.H{"msg": "swyl-v1"})
}


// @notice Method of TierController struct
// 
// @route `PATCH/update-tier`
// 
// @dev Lets a clubOwner update a tier
// 
// @param gc *gin.Context
func (tc *TierController) UpdateTier(gc *gin.Context) {
   gc.JSON(200, gin.H{"msg": "swyl-v1"})
}


// @notice Method of TierController struct
// 
// @route `POST/connect`
// 
// @param gc *gin.Context
func (tc *TierController) DeleteTier(gc *gin.Context){
   gc.JSON(200, gin.H{"msg": "swyl-v1"})
}