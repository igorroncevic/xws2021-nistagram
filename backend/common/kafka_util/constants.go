package kafka_util

import "time"

const (
	ExampleGroupId         = "groupId"
	PerformanceTopic       = "performance"
	UserEventsTopic        = "user-events"
	RetryTopic             = "retry"
	RegularConsumerMaxWait = time.Duration(10) * time.Second
	RetryConsumerMaxWait   = time.Duration(5) * time.Second

	/* Services */
	UserService    			= "UserService"
	ContentService 			= "ContentService"
	RecommendationService 	= "RecommendationService"

	/* Functions */
	LoginFunction              			= "Login"
	CreateNotificationFunction 			= "CreateNotification"
	GenerateApiTokenFunction   			= "GenerateApiToken"

	CreateAdFunction 		   			= "CreateAd"
	CreateAdCategoryFunction   			= "CreateAdCategory"
	GetUsersAdCategoriesFunction 		= "GetUsersAdCategories"
	CreateCampaignFunction	   			= "CreateCampaign"
	UpdateCampaignFunction     			= "UpdateCampaignFunction"
	CreateCommentFunction	   			= "CreateComment"
	CreateLikeFunction	   	   			= "CreateLike"
	CreateContentComplaintFunction 	   	= "CreateContentComplaint"
	CreateFavoriteFunction	   		   	= "CreateFavorite"
	RemoveFavoriteFunction	   		   	= "RemoveFavorite"
	CreateCollectionFunction	   	   	= "CreateCollection"
	RemoveCollectionFunction	   	   	= "RemoveCollection"
	CreateHighlightStoryFunction	   	= "CreateHighlightStory"
	RemoveHighlightStoryFunction	   	= "RemoveHighlightStory"
	CreateHighlightFunction	   			= "CreateHighlight"
	RemoveHighlightFunction	   			= "RemoveHighlight"
	CreatePostFunction					= "CreatePost"
	RemovePostFunction					= "RemovePost"
	CreateStoryFunction					= "CreateStory"
	RemoveStoryFunction					= "RemoveStory"

	CreateUserConnectionFunction		= "CreateUserConnection"
	DeleteBiDirectedConnectionFunction  = "DeleteBiDirectedConnection"
	DeleteDirectedConnectionFunction	= "DeleteDirectedConnection"
	CreateUserFunction					= "CreateUser"
	UpdateUserConnectionFunction		= "UpdateUserConnection"
	AcceptFollowRequestFunction			= "AcceptFollowRequest"

)
