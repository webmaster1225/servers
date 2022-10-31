/*
   @dev: Logan (Nam) Nguyen
   @course: SUNY Oswego - CSC 495 - Capstone
   @instructor: Professor Bastian Tenbergen
   @version: 1.0
*/

// @package
package models

// @notice The information related to a Swyl club
type Club struct {
	Club_ID						*uint64 	`json:"club_id" bson:"club_id" validate:"required"`
	Club_owner					*string		`json:"club_owner" bson:"club_owner" validate:"required,len=42,alphanum"`
	Created_at					*uint64		`json:"created_at" bson:"created_at"`
	Total_member					*uint64		`json:"total_member" bson:"total_member"`
}


// @notice The information related to a Swyl Tier
type Tier struct {
	Club_ID						*uint64 	`json:"club_id" bson:"club_id" validate:"required"`
	Tier_ID						*uint64		`json:"tier_id" bson:"tier_id" validate:"required"`
	Tier_name					*string		`json:"tier_name" bson:"tier_name" validate:"required,min=2,max=20"`
	Tier_img					*string		`json:"tier_img" bson:"tier_img"`
	Tier_bio					*string		`json:"tier_bio" bson:"tier_bio" validate:"omitempty,min=2,max=200"`
	Tier_fee					*uint64		`json:"tier_fee" bson:"tier_fee"`
	Tier_welcome_msg				*string		`json:"Tier_welcome_msg" bson:"Tier_welcome_msg" validate:"omitempty,min=2,max=100"`
	Created_at					*uint64		`json:"created_at" bson:"created_at"`
}


// @notice The information related to a Swyl Subscription
type Subscription struct {
	Club_ID						*uint64 	`json:"club_id" bson:"club_id" validate:"required"`
	Tier_ID						*uint64		`json:"tier_id" bson:"tier_id" validate:"required"`
	Subscription_ID					*int64 		`json:"subscription_id" bason:"subscription_id"`
	Subscriber					*string 	`json:"club_owner" bson:"club_owner" validate:"required,len=42,alphanum"`
	Joined_at					*uint64		`json:"joined_at" bson:"joined_at"`
}
