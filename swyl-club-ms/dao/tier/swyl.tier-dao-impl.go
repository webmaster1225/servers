/*
   @dev: Logan (Nam) Nguyen
   @course: SUNY Oswego - CSC 495 - Capstone
   @instructor: Professor Bastian Tenbergen
   @version: 1.0
*/

// @package
package dao

// @import
import (
	"Swyl/servers/swyl-club-ms/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// @notice Root struct for other methods in dao-impl
type TierDaoImpl struct {
   ctx               context.Context
   mongoCollection   *mongo.Collection
} 


// @dev Constructor
func TierDaoConstructor(ctx context.Context, mongoCollection *mongo.Collection) TierDao {
   return &TierDaoImpl {
      ctx: ctx,
      mongoCollection: mongoCollection,
   }
}

// @notice Method of TierDaoImpl struct
// 
// @dev Lets a club owner create a tier to internal database
// 
// @param tier *models.Tier
// 
// @return error
func (ti *TierDaoImpl) CreateTier(tier *models.Tier) error {
   // updated tier.Tier_ID
   tier.Tier_ID = primitive.NewObjectID()

   // updated tier.Created_at
   tier.Created_at = uint64(time.Now().Unix())

   // insert the tier to the internal database
   _, err := ti.mongoCollection.InsertOne(ti.ctx, tier)
   if err != nil {return err}

   // return OK
   return nil
}


// @notice Method of TierDaoImpl struct
// 
// @dev Gets a Tier at tierId and clubOwner
// 
// @param clubId *uint64
// 
// @return *models.Tier
// 
// @return error
func (ti *TierDaoImpl) GetTierAt(tierId *string) (*models.Tier, error) {
   // prepare tier struct
   tier := &models.Tier{}

   // set up objectId
   objectId, err := primitive.ObjectIDFromHex(*tierId)
   if err != nil {return nil, err}

   // find the document with _id = tierId in
   if err := ti.mongoCollection.FindOne(ti.ctx, bson.M{"_id": objectId}).Decode(tier); err != nil {return nil, err}

   // return OK
   return tier, nil
}


// @notice Method of TierDaoImpl struct
// 
// @dev Gets all tiers owned by clubOwner
// 
// @param clubOwner *string
// 
// @return *[]models.Tier
// 
// @return error
func (ti *TierDaoImpl) GetAllTiersOwnedBy(clubOwner *string) (*[]models.Tier, error) {
   // prepare tier struct holder slice
   tiers := &[]models.Tier{}

   // find documents in database
   cursor, err := ti.mongoCollection.Find(ti.ctx, bson.M{"club_owner": clubOwner})
   if err != nil {return nil, err}

   // decode cursor into declared slice
   if err := cursor.All(ti.ctx, tiers); err != nil {return nil, err}

   // return OK
   return tiers, nil
}


// @notice Method of TierDaoImpl struct
// 
// @dev Lets a clubOwner update a tier
// 
// @param tier *models.Tier
// 
// @return error
func (ti *TierDaoImpl) UpdateTier(tier *models.Tier) error {
   return nil
}


// @notice Lets a clubOwner delete a tier
// 
// @param tierId *uint64
// 
// @param clubOwner *string
// 
// @return error
func (ti *TierDaoImpl) DeleteTier(tierId *uint64, clubOwner *string) error {
   return nil
}