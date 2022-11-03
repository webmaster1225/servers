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
	dao "Swyl/servers/swyl-community-ms/dao/post"
	"Swyl/servers/swyl-community-ms/models"
	"Swyl/servers/swyl-community-ms/utils"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

//@notice global vars
var validate = validator.New()

// @notice Root struct for other methods in post-controller
type PostController struct {
	PostDao dao.PostDao
}

// @dev Constructor
func PostControllerConstructor(postDao dao.PostDao) *PostController {
	return &PostController {
		PostDao: postDao,
	}
}


// #############################################
// 			 		Post Handlers 
// #############################################

// @notice Method of PostController struct
// 
// @route `POST/create-post`
// 
// @dev Lets a commOwner create a post
// 
// @param gc *gin.Context
func (pi *PostController) CreatePost(gc *gin.Context) {
	// decalre param holder
	param := &models.Post{}

	// bin json post data to param holder
	if err := gc.ShouldBindJSON(param); err != nil {gc.AbortWithStatusJSON(400, gin.H{"error": err.Error()}); return;}

	// validate struct param
	if err := validate.Struct(param); err != nil {gc.AbortWithStatusJSON(400, gin.H{"error": err.Error()}); return;}

	// extra validation on param.Community_owner to match ETH Crypto wallet address convention
	matched, err := utils.TestEthAddress(param.Community_owner)
	if err != nil{gc.AbortWithStatusJSON(400, gin.H{"error": "!REGEX - cannot test Community_owner using regex"}); return;}
	if !matched {gc.AbortWithStatusJSON(400, gin.H{"error": "!ETH_ADDRESS - Community_owner is not an ETH crypto wallet address"}); return;}

	// invokde PostDao.CreatePost()
	if err := pi.PostDao.CreatePost(param); err != nil {gc.AbortWithStatusJSON(500, gin.H{"error": err.Error()}); return}

	// http response
	gc.JSON(200, "Post successfully created")
}


// @notice Method of PostController struct
// 
// @route `GET/get-post-at/:post_id`
// 
// @dev Gets a post at postId
// 
// @param gc *gin.Context
func (pi *PostController) GetPostAt(gc *gin.Context) {
	// Handle postId param
	postId := gc.Param("post_id")

	// sanitizing postId
	matched, err := regexp.MatchString(`^[a-fA-f0-9]{24}$`, postId)
	if err != nil {gc.AbortWithStatusJSON(400, gin.H{"error": "!REGEX - cannot test postId using regex"}); return;}
	if !matched {gc.AbortWithStatusJSON(400, gin.H{"error": "!TIERID - postId is not valid"}); return;}
	
	// invokde PostDao.GetPostAt()
	post, err := pi.PostDao.GetPostAt(&postId)
	if err != nil {gc.AbortWithStatusJSON(500, gin.H{"error": err.Error()}); return}

	// http response
	gc.JSON(200, post)
}


// @notice Method of PostController struct
// 
// @route `GET/get-all-posts-by/:community_owner`
// 
// @dev Gets all posts created by commOwner
// 
// @param gc *gin.Context
func (pi *PostController) GetPostsBy(gc *gin.Context) {
	// Handle commOwner param
	commOwner := gc.Param("community_owner")

	// validate commOwner to match ETH Crypto Wallet address convention
	matched, err := utils.TestEthAddress(&commOwner)
	if err != nil {gc.AbortWithStatusJSON(400, gin.H{"error": "!REGEX - cannot test commOwner using regex"}); return;}
	if !matched {gc.AbortWithStatusJSON(400, gin.H{"error": "!ETH_ADDRESS - commOwner is not an ETH crypto wallet address"}); return;}
	
	// invokde PostDao.GetPostsBy()
	posts, err := pi.PostDao.GetPostsBy(&commOwner); 
	if err != nil {gc.AbortWithStatusJSON(500, gin.H{"error": err.Error()})}

	// http response
	gc.JSON(200, posts)
}


// @notice Method of PostController struct
// 
// @route `PATCH/update-post-content`
// 
// @dev Lets a commOwner update a post - only post.Title and post.Content are allowed
// 
// @param gc *gin.Context
func (pi *PostController) UpdatePostContent(gc *gin.Context) {gc.JSON(200, "Swyl-v1")}


// @notice Method of PostController struct
// 
// @route `PATCH/react-post`
// 
// @dev Lets a user react to a post
// 
// @logic if the post never has `reaction.React_type` by reacter before, add `reaction.React_type` to post.Reaction
// 
// @logic if the post has `reaction.React_type` by reacter before, remove `reaction.React_type` from post.Reaction
// 
// @logic if the post has different `reaction.React_type` by reacter, update `reaction.React_type` in post.Reaction
// 
// @param gc *gin.Context
func (pi *PostController) ReactPost(gc *gin.Context) {gc.JSON(200, "Swyl-v1")}


// @notice Method of PostController struct
// 
// @route `DELETE/delete-post-at/:post_id`
// 
// @dev Lets a commOwner delete own post
// 
// @param gc *gin.Context
func (pi *PostController) DeletePostAt(gc *gin.Context) {gc.JSON(200, "Swyl-v1")}


// #############################################
// 			 		Comment Handlers 
// #############################################

// @notice Method of PostController struct
// 
// @route `POST/comment-post`
// 
// @dev Lets a user comment on a post
// 
// @param gc *gin.Context
func (pi *PostController) Comment(gc *gin.Context) {gc.JSON(200, "Swyl-v1")}


// @notice Method of PostController struct
// 
// @route `GET/get-comment-at/:comment_id`
// 
// @dev Gets a comment at commentId
// 
// @param gc *gin.Context
func (pi *PostController) GetCommentAt(gc *gin.Context) {gc.JSON(200, "Swyl-v1")}


// @notice Method of PostController struct
// 
// @route `GET/get-all-comments-at/:post_id`
// 
// @dev Gets all comments at postId
// 
// @param gc *gin.Context
func (pi *PostController) GetAllCommentsAt(gc *gin.Context) {gc.JSON(200, "Swyl-v1")}


// @notice Method of PostController struct
// 
// @route `PATCH/update-comment-content`
// 
// @dev Lets a user to update own comment - only comment.Content is allowed
// 
// @param gc *gin.Context
func (pi *PostController) UpdateCommentContent(gc *gin.Context) {gc.JSON(200, "Swyl-v1")}


// @notice Method of PostController struct
// 
// @route `PATCH/react-comment`
// 
// @dev Lets a user react to a comment
// 
// @logic if the post never has `reaction.React_type` by reacter before, add `reaction.React_type` to comment.Reaction
// 
// @logic if the post has `reaction.React_type` by reacter before, remove `reaction.React_type` from comment.Reaction
// 
// @logic if the post has different `reaction.React_type` by reacter, update `reaction.React_type` in comment.Reaction
// 
// @param gc *gin.Context
func (pi *PostController) ReactComment(gc *gin.Context) {gc.JSON(200, "Swyl-v1")}


// @notice Method of PostController struct
// 
// @route `DELETE/delete-comment-at/:comment_id`
// 
// @dev Lets a user delete own comment at commentId
// 
// @param gc *gin.Context
func (pi *PostController) DeleteCommentAt(gc *gin.Context) {gc.JSON(200, "Swyl-v1")}


// #############################################
// 			 		Reply Handlers 
// #############################################

// @notice Method of PostController struct
// 
// @route `POST/reply`
// 
// @dev Lets a user reply to a comment
// 
// @param gc *gin.Context
func (pi *PostController) Reply(gc *gin.Context) {gc.JSON(200, "Swyl-v1")}


// @notice Method of PostController struct
// 
// @route `GET/get-reply-at/:reply_id`
// 
// @dev Gets a reply at replyId
// 
// @param gc *gin.Context
func (pi *PostController) GetReplyAt(gc *gin.Context) {gc.JSON(200, "Swyl-v1")}


// @notice Method of PostController struct
// 
// @route `GET/get-all-reply-at/:comment_id`
// 
// @dev Gets all replies at commentId
// 
// @param gc *gin.Context
func (pi *PostController) GetAllRepliesAt(gc *gin.Context) {gc.JSON(200, "Swyl-v1")}


// @notice Method of PostController struct
// 
// @route `PATCH/update-reply`
// 
// @dev Lets a user update own reply - only reply.Content is allowed
// 
// @param gc *gin.Context
func (pi *PostController) UpdateReplyContent(gc *gin.Context) {gc.JSON(200, "Swyl-v1")}


// @notice Method of PostController struct
// 
// @route `PATCH/react-reply`
// 
// @dev Lets a user react a reply
// 
// @logic if the post never has `reaction.React_type` by reacter before, add `reaction.React_type` to comment.Reaction
// 
// @logic if the post has `reaction.React_type` by reacter before, remove `reaction.React_type` from comment.Reaction
// 
// @logic if the post has different `reaction.React_type` by reacter, update `reaction.React_type` in comment.Reaction
// 
// @param gc *gin.Context
func (pi *PostController) ReactReply(gc *gin.Context) {gc.JSON(200, "Swyl-v1")}


// @notice Method of PostController struct
// 
// @route `DELETE/delete-reply-at`
// 
// @dev Lets a user delete own reply at replyId
// 
// @param gc *gin.Context
func (pi *PostController) DeleteReplyAt(gc *gin.Context) {gc.JSON(200, "Swyl-v1")}