package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	rt.router.PUT("/profiles/:userid/subscribe/:userid", rt.wrap(rt.followUser))

	rt.router.GET("/fountains/", rt.wrap(rt.listFountains))
	rt.router.POST("/fountains/", rt.wrap(rt.createFountain))
	rt.router.PUT("/fountains/:id", rt.wrap(rt.updateFountain))
	rt.router.DELETE("/fountains/:id", rt.wrap(rt.deleteFountain))

	rt.router.GET("/profiles", rt.wrap(rt.userSearch))
	rt.router.POST("/session", rt.wrap(rt.session))
	rt.router.GET("/posts", rt.wrap(rt.getStream))

	// profile component is below
	// userprofile is below:::: check for that data inside

	// rt.router.GET("/profiles/:userid", rt.wrap(rt.getMyProfile))
	rt.router.GET("/profiles/:userid/posts", rt.wrap(rt.getProfilePosts))
	rt.router.POST("/profiles/:userid/posts", rt.wrap(rt.createProfilePost))
	rt.router.PUT("/profiles/:userid/posts/:id", rt.wrap(rt.updateProfilePost))
	rt.router.DELETE("/profiles/:userid/posts/:id", rt.wrap(rt.deleteProfilePost))

	// below u have data for implementing likes and dislikes inside of your component
	rt.router.PUT("/posts/:postid/like/:userid", rt.wrap(rt.addLikePost))
	rt.router.DELETE("/posts/:postid/like/:userid", rt.wrap(rt.deleteLikePost))
	rt.router.PUT("/posts/:postid/dislike/:userid", rt.wrap(rt.addDislikePost))
	rt.router.DELETE("/posts/:postid/dislike/:userid", rt.wrap(rt.deleteDislikePost))

	// below u should put the first line of code u inserted above:::::

	rt.router.DELETE("/profiles/:userid/subscribe/:userid", rt.wrap(rt.unfollowUser))
	// below u should implement comments data for better performance and check how it works

	rt.router.GET("/posts/:postid/comments", rt.wrap(rt.getCommentsRelatedToPost))
	rt.router.POST("/posts/:postid/comments", rt.wrap(rt.createCommentRelatedToPost))
	rt.router.PUT("/posts/:postid/comments/:commentid", rt.wrap(rt.updateCommentRelatedToPost))
	rt.router.DELETE("/posts/:postid/comments/:commentid", rt.wrap(rt.deleteCommentRelatedToPost))
	rt.router.PUT("/posts/:postid/comments/:commentid/like/:userid", rt.wrap(rt.addLikeToCommentRelatedToPost))
	rt.router.DELETE("/posts/:postid/comments/:commentid/like/:userid", rt.wrap(rt.removeLikeFromCommentRelatedToPost))
	rt.router.PUT("/posts/:postid/comments/:commentid/dislike/:userid", rt.wrap(rt.addDislikeToCommentRelatedToPost))
	rt.router.DELETE("/posts/:postid/comments/:commentid/dislike/:userid", rt.wrap(rt.removeDislikeFromCommentRelatedToPost))

	rt.router.GET("/profiles/:userid", rt.wrap(rt.getProfile))
	rt.router.PUT("/profiles/:userid/banuser/:userid", rt.wrap(rt.banUser))
	rt.router.DELETE("/profiles/:userid/banuser/:userid", rt.wrap(rt.unbanUser))
	rt.router.GET("/profiles/:userid/banuser", rt.wrap(rt.getBannedUsers))
	rt.router.PATCH("/profiles/:userid/changePassword", rt.wrap(rt.changePassword))
	rt.router.PATCH("/profiles/:userid/changeAvatar", rt.wrap(rt.changeAvatar))
	rt.router.PATCH("/profiles/:userid/changeUsername", rt.wrap(rt.changeUsername))
	rt.router.DELETE("/profiles/:userid/deleteAccount", rt.wrap(rt.deleteAccount))

	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
