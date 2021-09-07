package kafka_util

func GetUserEventMessage(eventType UserEventType, success bool) string {
	switch eventType {
	case Login:
		if success {
			return "Successful login attempt."
		}
		return "Failed login attempt."
	case PasswordChange:
		if success {
			return "Successful password change."
		}
		return "Failed password change."
	case ProfileUpdate:
		if success {
			return "Successful profile update."
		}
		return "Failed profile update."
	case AdCategoryUpdate:
		if success {
			return "Successful ad categories update."
		}
		return "Failed ad categories update."
	case CampaignUpdate:
		if success {
			return "Successfully updated a campaign"
		}
		return "Failed to update a campaign"
	case DeleteCampaign:
		if success {
			return "Successfully delete a campaign"
		}
		return "Failed to delete a campaign"
	case CreateContentComplaint:
		if success {
			return "Successfully created a content complaint"
		}
		return "Failed to create content a complaint"
	case CreatePost:
		if success {
			return "Successfully created a post"
		}
		return "Failed to create a post"
	case RemovePost:
		if success {
			return "Successfully removed a post"
		}
		return "Failed to remove a post"
	case CreateStory:
		if success {
			return "Successfully created a story"
		}
		return "Failed to create a story"
	case RemoveStory:
		if success {
			return "Successfully removed a story"
		}
		return "Failed to remove a story"

	default:
		return ""
	}
}

func GetPerformanceMessage(eventType string, success bool) string {
	switch eventType {
	case CreateAdFunction:
		if success {
			return "Successful created new ad."
		}
		return "Failed to create new ad."
	case CreateAdCategoryFunction:
		if success {
			return "Successfully created ad category."
		}
		return "Failed to create new ad category."
	case GetUsersAdCategoriesFunction:
		if success {
			return "Successfully retrieved ad categories"
		}
		return "Failed to retrieve ad categories"
	case CreateCampaignFunction:
		if success {
			return "Successfully created campaign"
		}
		return "Failed to create campaign"
	case UpdateCampaignFunction:
		if success {
			return "Successfully updated campaign"
		}
		return "Failed to update campaign"
	case CreateCommentFunction:
		if success {
			return "Successfully created a comment"
		}
		return "Failed to create a comment"
	case CreateLikeFunction:
		if success {
			return "Successfully created a like"
		}
		return "Failed to create a like"
	case CreateContentComplaintFunction:
		if success {
			return "Successfully created a content complaint"
		}
		return "Failed to create a content complaint"
	case CreateFavoriteFunction:
		if success {
			return "Successfully created a favorite"
		}
		return "Failed to create a favorite"
	case RemoveFavoriteFunction:
		if success {
			return "Successfully removed a favorite"
		}
		return "Failed to remove a favorite"
	case CreateCollectionFunction:
		if success {
			return "Successfully created a collection"
		}
		return "Failed to create a collection"
	case RemoveCollectionFunction:
		if success {
			return "Successfully removed a collection"
		}
		return "Failed to remove a collection"
	case CreateHighlightStoryFunction:
		if success {
			return "Successfully created a highlight story"
		}
		return "Failed to create a highlight story"
	case CreateHighlightFunction:
		if success {
			return "Successfully created a highlight"
		}
		return "Failed to create a highlight"
	case RemoveHighlightFunction:
		if success {
			return "Successfully removed a highlight"
		}
		return "Failed to remove a post"
	case CreatePostFunction:
		if success {
			return "Successfully created a post"
		}
		return "Failed to create a post"
	case RemovePostFunction:
		if success {
			return "Successfully removed a post"
		}
		return "Failed to remove a post"
	case CreateStoryFunction:
		if success {
			return "Successfully created a story"
		}
		return "Failed to create a story"
	case RemoveStoryFunction:
		if success {
			return "Successfully removed a story"
		}
		return "Failed to remove a story"

	default:
		return ""
	}
}
