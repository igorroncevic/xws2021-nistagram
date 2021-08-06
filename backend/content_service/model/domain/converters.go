package domain

import (
	protopb "github.com/igorroncevic/xws2021-nistagram/common/proto"
	"github.com/igorroncevic/xws2021-nistagram/content_service/model"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Post Converters
func (p *Post) ConvertFromGrpc(post *protopb.Post) *Post {
	if post == nil {
		post = &protopb.Post{}
	}
	return &Post{
		Objava: Objava{
			Id:          post.Id,
			UserId:      post.UserId,
			IsAd:        post.IsAd,
			Type:        model.GetPostType(post.Type),
			Description: post.Description,
			Location:    post.Location,
			CreatedAt:   post.CreatedAt.AsTime(),
			Media:       ConvertMultipleMediaFromGrpc(post.Media),
			Hashtags:    ConvertMultipleHashtagFromGrpc(post.Hashtags),
		},
		Comments: ConvertMultipleCommentsFromGrpc(post.Comments),
		Likes:    ConvertMultipleLikesFromGrpc(post.Likes),
		Dislikes: ConvertMultipleLikesFromGrpc(post.Dislikes),
	}
}

func (p Post) ConvertToGrpc() *protopb.Post {
	return &protopb.Post{
		Id:          p.Id,
		UserId:      p.UserId,
		IsAd:        p.IsAd,
		Type:        p.Type.String(),
		Description: p.Description,
		Location:    p.Location,
		CreatedAt:   timestamppb.New(p.CreatedAt),
		Media:       ConvertMultipleMediaToGrpc(p.Media),
		Comments:    ConvertMultipleCommentsToGrpc(p.Comments),
		Likes:       ConvertMultipleLikesToGrpc(p.Likes),
		Dislikes:    ConvertMultipleLikesToGrpc(p.Dislikes),
		Hashtags:    ConvertMultipleHashtagToGrpc(p.Hashtags),
	}
}

func ConvertMultiplePostsToGrpc(posts []Post) []*protopb.Post {
	grpcPosts := []*protopb.Post{}
	for _, post := range posts {
		grpcPosts = append(grpcPosts, post.ConvertToGrpc())
	}

	return grpcPosts
}

// Story Converters
func (s Story) ConvertToGrpc() *protopb.Story {
	return &protopb.Story{
		Id:             s.Id,
		UserId:         s.UserId,
		IsAd:           s.IsAd,
		Type:           s.Type.String(),
		Description:    s.Description,
		Location:       s.Location,
		CreatedAt:      timestamppb.New(s.CreatedAt),
		Media:          ConvertMultipleMediaToGrpc(s.Media),
		Hashtags:       ConvertMultipleHashtagToGrpc(s.Hashtags),
		IsCloseFriends: s.IsCloseFriends,
	}
}

func (s *Story) ConvertFromGrpc(story *protopb.Story) *Story {
	if story == nil {
		story = &protopb.Story{}
	}
	return &Story{
		Objava: Objava{
			Id:          story.Id,
			UserId:      story.UserId,
			IsAd:        story.IsAd,
			Type:        model.GetPostType(story.Type),
			Description: story.Description,
			Location:    story.Location,
			CreatedAt:   story.CreatedAt.AsTime(),
			Media:       ConvertMultipleMediaFromGrpc(story.Media),
			Hashtags:    ConvertMultipleHashtagFromGrpc(story.Hashtags),
		},
		IsCloseFriends: story.IsCloseFriends,
	}
}

func (s *StoriesHome) ConvertToGrpc(ads []*protopb.StoryAdHome) *protopb.StoriesHome {
	storiesHome := &protopb.StoriesHome{
		Ads: ads,
	}
	for _, storyHome := range s.Stories {
		storiesHome.Stories = append(storiesHome.Stories, &protopb.StoryHome{
			UserId:    storyHome.UserId,
			Username:  storyHome.Username,
			UserPhoto: storyHome.UserPhoto,
			Stories:   ConvertMultipleStoriesToGrpc(storyHome.Stories),
		})
	}

	return storiesHome
}

// HighlightRequest Converters
func (hr HighlightRequest) ConvertToGrpc() *protopb.HighlightRequest {
	return &protopb.HighlightRequest{
		UserId:      hr.UserId,
		HighlightId: hr.HighlightId,
		StoryId:     hr.StoryId,
	}
}

func (hr *HighlightRequest) ConvertFromGrpc(request *protopb.HighlightRequest) *HighlightRequest {
	if request == nil {
		request = &protopb.HighlightRequest{}
	}
	return &HighlightRequest{
		UserId:      request.UserId,
		HighlightId: request.HighlightId,
		StoryId:     request.StoryId,
	}
}

// HighlightRequest Converters
func (h Highlight) ConvertToGrpc() *protopb.Highlight {
	return &protopb.Highlight{
		Id:      h.Id,
		Name:    h.Name,
		UserId:  h.UserId,
		Stories: ConvertMultipleStoriesToGrpc(h.Stories),
	}
}

func (h *Highlight) ConvertFromGrpc(request *protopb.Highlight) *Highlight {
	if request == nil {
		request = &protopb.Highlight{}
	}
	return &Highlight{
		Id:      request.Id,
		Name:    request.Name,
		UserId:  request.UserId,
		Stories: ConvertMultipleStoriesFromGrpc(request.Stories),
	}
}

func ConvertMultipleStoriesFromGrpc(stories []*protopb.Story) []Story {
	convertedStories := []Story{}
	for _, story := range stories {
		var converted *Story
		converted = converted.ConvertFromGrpc(story)

		convertedStories = append(convertedStories, *converted)
	}

	return convertedStories
}

func ConvertMultipleStoriesToGrpc(stories []Story) []*protopb.Story {
	grpcStories := []*protopb.Story{}
	for _, story := range stories {
		grpcStories = append(grpcStories, story.ConvertToGrpc())
	}

	return grpcStories
}

// ReducedPost Converters
func (p ReducedPost) ConvertToGrpc() *protopb.ReducedPost {
	return &protopb.ReducedPost{
		Id:          p.Id,
		UserId:      p.UserId,
		IsAd:        p.IsAd,
		Type:        p.Type.String(),
		Description: p.Description,
		Location:    p.Location,
		CreatedAt:   timestamppb.New(p.CreatedAt),
		Media:       ConvertMultipleMediaToGrpc(p.Media),
		CommentsNum: p.CommentsNum,
		LikesNum:    p.LikesNum,
		DislikesNum: p.DislikesNum,
	}
}

func (p ReducedPost) ConvertFromGrpc(post *protopb.ReducedPost) ReducedPost {
	return ReducedPost{
		Objava: Objava{
			Id:          post.Id,
			UserId:      post.UserId,
			IsAd:        post.IsAd,
			Type:        model.GetPostType(post.Type),
			Description: post.Description,
			Location:    post.Location,
			CreatedAt:   post.CreatedAt.AsTime(),
			Media:       ConvertMultipleMediaFromGrpc(post.Media),
		},
		CommentsNum: post.CommentsNum,
		LikesNum:    post.LikesNum,
		DislikesNum: post.DislikesNum,
	}
}

func ConvertMultipleReducedPostsToGrpc(posts []ReducedPost) []*protopb.ReducedPost {
	grpcPosts := []*protopb.ReducedPost{}
	for _, post := range posts {
		grpcPosts = append(grpcPosts, post.ConvertToGrpc())
	}

	return grpcPosts
}

func ConvertMultipleReducedPostsFromGrpc(posts []*protopb.ReducedPost) []ReducedPost {
	convertedPosts := []ReducedPost{}
	for _, post := range posts {
		var converted ReducedPost
		converted = converted.ConvertFromGrpc(post)

		convertedPosts = append(convertedPosts, converted)
	}

	return convertedPosts
}

func ConvertMultiplePostsFromGrpc(posts []*protopb.Post) []Post {
	convertedPosts := []Post{}
	for _, post := range posts {
		var converted *Post
		converted = converted.ConvertFromGrpc(post)

		convertedPosts = append(convertedPosts, *converted)
	}

	return convertedPosts
}

// Media converters
func (m *Media) ConvertFromGrpc(media *protopb.Media) *Media {
	if m == nil {
		m = &Media{}
	}
	return &Media{
		Id:       media.Id,
		Type:     model.GetMediaType(media.Type),
		PostId:   media.PostId,
		Content:  media.Content,
		OrderNum: media.OrderNum,
		Tags:     ConvertMultipleTagFromGrpc(media.Tags),
	}
}

func (m *Media) ConvertToGrpc() *protopb.Media {
	if m == nil {
		m = &Media{}
	}
	return &protopb.Media{
		Id:       m.Id,
		Type:     m.Type.String(),
		PostId:   m.PostId,
		Content:  m.Content,
		OrderNum: m.OrderNum,
		Tags:     ConvertMultipleTagsToGrpc(m.Tags),
	}
}

func ConvertMultipleMediaFromGrpc(m []*protopb.Media) []Media {
	media := []Media{}
	if m != nil {
		for _, protoMedia := range m {
			var domainMedia *Media
			domainMedia = domainMedia.ConvertFromGrpc(protoMedia)
			media = append(media, *domainMedia)
		}
	}
	return media
}

func ConvertMultipleMediaToGrpc(m []Media) []*protopb.Media {
	media := []*protopb.Media{}
	if m != nil {
		for _, domainMedia := range m {
			protoMedia := domainMedia.ConvertToGrpc()
			media = append(media, protoMedia)
		}
	}
	return media
}

// Comment converters
func (c *Comment) ConvertFromGrpc(comment *protopb.Comment) *Comment {
	if c == nil {
		c = &Comment{}
	}
	return &Comment{
		Id:        comment.Id,
		PostId:    comment.PostId,
		UserId:    comment.UserId,
		Username:  comment.Username,
		Content:   comment.Content,
		CreatedAt: comment.CreatedAt.AsTime(),
	}
}

func (c *Comment) ConvertToGrpc() *protopb.Comment {
	if c == nil {
		c = &Comment{}
	}
	return &protopb.Comment{
		Id:        c.Id,
		PostId:    c.PostId,
		UserId:    c.UserId,
		Username:  c.Username,
		Content:   c.Content,
		CreatedAt: timestamppb.New(c.CreatedAt),
	}
}

func ConvertMultipleCommentsFromGrpc(c []*protopb.Comment) []Comment {
	comments := []Comment{}
	if c != nil {
		for _, protoComment := range c {
			var domainComment *Comment
			domainComment = domainComment.ConvertFromGrpc(protoComment)
			comments = append(comments, *domainComment)
		}
	}
	return comments
}

func ConvertMultipleCommentsToGrpc(c []Comment) []*protopb.Comment {
	comments := []*protopb.Comment{}
	if c != nil {
		for _, domainComment := range c {
			protoComment := domainComment.ConvertToGrpc()
			comments = append(comments, protoComment)
		}
	}
	return comments
}

// Tag converters
func (t *Tag) ConvertFromGrpc(tag *protopb.Tag) *Tag {
	if t == nil {
		t = &Tag{}
	}
	return &Tag{
		MediaId:  tag.MediaId,
		UserId:   tag.UserId,
		Username: tag.Username,
	}
}

func (t *Tag) ConvertToGrpc() *protopb.Tag {
	if t == nil {
		t = &Tag{}
	}
	return &protopb.Tag{
		MediaId:  t.MediaId,
		UserId:   t.UserId,
		Username: t.Username,
	}
}

func ConvertMultipleTagFromGrpc(t []*protopb.Tag) []Tag {
	tags := []Tag{}
	if t != nil {
		for _, protoTag := range t {
			var domainTag *Tag
			domainTag = domainTag.ConvertFromGrpc(protoTag)
			tags = append(tags, *domainTag)
		}
	}

	return tags
}

func ConvertMultipleTagsToGrpc(t []Tag) []*protopb.Tag {
	tags := []*protopb.Tag{}
	if t != nil {
		for _, domainTag := range t {
			tags = append(tags, domainTag.ConvertToGrpc())
		}
	}
	return tags
}

// Like converters
func (l *Like) ConvertFromGrpc(like *protopb.Like) *Like {
	if l == nil {
		l = &Like{}
	}
	return &Like{
		PostId: like.PostId,
		UserId: like.UserId,
		IsLike: like.IsLike,
	}
}

func (l *Like) ConvertToGrpc() *protopb.Like {
	if l == nil {
		l = &Like{}
	}
	return &protopb.Like{
		PostId:   l.PostId,
		UserId:   l.UserId,
		IsLike:   l.IsLike,
		Username: l.Username,
	}
}

func ConvertMultipleLikesFromGrpc(l []*protopb.Like) []Like {
	likes := []Like{}
	if l != nil {
		for _, protoLike := range l {
			var domainLike *Like
			domainLike.ConvertFromGrpc(protoLike)
			likes = append(likes, *domainLike)
		}
	}

	return likes
}

func ConvertMultipleLikesToGrpc(l []Like) []*protopb.Like {
	likes := []*protopb.Like{}
	if l != nil {
		for _, domainLike := range l {
			likes = append(likes, domainLike.ConvertToGrpc())
		}
	}
	return likes
}

func (c Collection) ConvertFromGrpc(collection *protopb.Collection) Collection {
	return Collection{
		Id:     collection.Id,
		Name:   collection.Name,
		UserId: collection.UserId,
		Posts:  ConvertMultiplePostsFromGrpc(collection.Posts),
	}
}

func (c Collection) ConvertToGrpc() *protopb.Collection {
	return &protopb.Collection{
		Id:     c.Id,
		Name:   c.Name,
		UserId: c.UserId,
		Posts:  ConvertMultiplePostsToGrpc(c.Posts),
	}
}

func ConvertMultipleCollectionsToGrpc(collections []Collection) []*protopb.Collection {
	converted := []*protopb.Collection{}
	for _, collection := range collections {
		converted = append(converted, collection.ConvertToGrpc())
	}

	return converted
}

func (f Favorites) ConvertToGrpc() *protopb.Favorites {
	return &protopb.Favorites{
		UserId:       f.UserId,
		Collections:  ConvertMultipleCollectionsToGrpc(f.Collections),
		Unclassified: ConvertMultiplePostsToGrpc(f.Unclassified),
	}
}

func (fr *FavoritesRequest) ConvertFromGrpc(request *protopb.FavoritesRequest) FavoritesRequest {
	return FavoritesRequest{
		PostId:       request.PostId,
		CollectionId: request.CollectionId,
		UserId:       request.UserId,
	}
}

// Hashtag converters
func (h *Hashtag) ConvertFromGrpc(hashtag *protopb.Hashtag) *Hashtag {
	if h == nil {
		h = &Hashtag{}
	}
	return &Hashtag{
		Id:   hashtag.Id,
		Text: hashtag.Text,
	}
}

func ConvertMultipleHashtagFromGrpc(t []*protopb.Hashtag) []Hashtag {
	hashtags := []Hashtag{}
	if t != nil {
		for _, protoTag := range t {
			var domainHashtag *Hashtag
			domainHashtag = domainHashtag.ConvertFromGrpc(protoTag)
			hashtags = append(hashtags, *domainHashtag)
		}
	}

	return hashtags
}

func ConvertMultipleHashtagToGrpc(t []Hashtag) []*protopb.Hashtag {
	hashtags := []*protopb.Hashtag{}
	if t != nil {
		for _, hashtag := range t {
			hashtags = append(hashtags, hashtag.ConvertToGrpc())
		}
	}

	return hashtags
}

func (h *Hashtag) ConvertToGrpc() *protopb.Hashtag {
	if h == nil {
		h = &Hashtag{}
	}
	return &protopb.Hashtag{
		Id:   h.Id,
		Text: h.Text,
	}
}

func (c *ContentComplaint) ConvertFromGrpc(contentComplaint *protopb.ContentComplaint) *ContentComplaint {
	if c == nil {
		c = &ContentComplaint{}
	}
	return &ContentComplaint{
		Id:       contentComplaint.Id,
		Category: model.ComplaintCategory(contentComplaint.Category),
		PostId:   contentComplaint.PostId,
		Status:   model.RequestStatus(contentComplaint.Status),
		IsPost:   contentComplaint.IsPost,
		UserId:   contentComplaint.UserId,
	}
}

func (c *ContentComplaint) ConvertToGrpc() *protopb.ContentComplaint {
	return &protopb.ContentComplaint{
		Id:       c.Id,
		Category: string(c.Category),
		PostId:   c.PostId,
		Status:   string(c.Status),
		IsPost:   c.IsPost,
		UserId:   c.UserId,
	}
}

func (a *Ad) ConvertToGrpc() *protopb.Ad {
	if a == nil {
		a = &Ad{}
	}
	return &protopb.Ad{
		Id:         a.Id,
		Link:       a.Link,
		CampaignId: a.CampaignId,
		LinkClicks: int32(a.LinkClicks),
		Post: &protopb.Post{
			Id:          a.Post.Id,
			UserId:      a.Post.UserId,
			IsAd:        a.Post.IsAd,
			Type:        a.Post.Type.String(),
			Description: a.Post.Description,
			Location:    a.Post.Location,
			CreatedAt:   timestamppb.New(a.Post.CreatedAt),
			Media:       ConvertMultipleMediaToGrpc(a.Post.Media),
			Comments:    ConvertMultipleCommentsToGrpc(a.Post.Comments),
			Likes:       ConvertMultipleLikesToGrpc(a.Post.Likes),
			Dislikes:    ConvertMultipleLikesToGrpc(a.Post.Dislikes),
			Hashtags:    ConvertMultipleHashtagToGrpc(a.Post.Hashtags),
		},
	}
}

func (a *Ad) ConvertFromGrpc(ad *protopb.Ad) *Ad {
	if a == nil {
		a = &Ad{}
	}
	if ad == nil {
		ad = &protopb.Ad{}
	}
	return &Ad{
		Id:         ad.Id,
		Link:       ad.Link,
		CampaignId: ad.CampaignId,
		Post: Post{
			Objava: Objava{
				Id:          ad.Post.Id,
				UserId:      ad.Post.UserId,
				IsAd:        ad.Post.IsAd,
				Type:        model.GetPostType(ad.Post.Type),
				Description: ad.Post.Description,
				Location:    ad.Post.Location,
				CreatedAt:   ad.Post.CreatedAt.AsTime(),
				Media:       ConvertMultipleMediaFromGrpc(ad.Post.Media),
				Hashtags:    ConvertMultipleHashtagFromGrpc(ad.Post.Hashtags),
			},
		},
		LinkClicks: int(ad.LinkClicks),
	}
}

func (a Ad) CreateStoryAd() *Story {
	return &Story{
		Objava: Objava{
			Id:          a.Post.Id,
			UserId:      a.Post.UserId,
			IsAd:        true,
			Type:        model.TypeStory,
			Description: a.Post.Description,
			Location:    a.Post.Location,
			CreatedAt:   a.Post.CreatedAt,
			Media:       a.Post.Media,
			Hashtags:    a.Post.Hashtags,
		},
		IsCloseFriends: false,
	}
}

func ConvertMultipleAdsToGrpc(ads []Ad) []*protopb.Ad {
	grpcAds := []*protopb.Ad{}
	for _, ad := range ads {
		grpcAds = append(grpcAds, ad.ConvertToGrpc())
	}
	return grpcAds
}

func ConvertMultipleAdsFromGrpc(ads []*protopb.Ad) []Ad {
	domainAds := []Ad{}
	for _, ad := range ads {
		var domainAd Ad
		domainAds = append(domainAds, *domainAd.ConvertFromGrpc(ad))
	}
	return domainAds
}

func (c Campaign) ConvertToGrpc() *protopb.Campaign {
	return &protopb.Campaign{
		Id:          c.Id,
		Name:        c.Name,
		IsOneTime:   c.IsOneTime,
		StartDate:   timestamppb.New(c.StartDate),
		EndDate:     timestamppb.New(c.EndDate),
		StartTime:   int32(c.StartTime),
		EndTime:     int32(c.EndTime),
		Placements:  int32(c.Placements),
		AgentId:     c.AgentId,
		Category:    c.Category.ConvertToGrpc(),
		LastUpdated: timestamppb.New(c.LastUpdated),
		Ads:         ConvertMultipleAdsToGrpc(c.Ads),
		Type:        c.Type.String(),
	}
}

func (c Campaign) ConvertFromGrpc(campaign *protopb.Campaign) Campaign {
	if campaign == nil {
		campaign = &protopb.Campaign{}
	}
	category := AdCategory{}
	return Campaign{
		Id:          campaign.Id,
		Name:        campaign.Name,
		IsOneTime:   campaign.IsOneTime,
		StartDate:   campaign.StartDate.AsTime(),
		EndDate:     campaign.EndDate.AsTime(),
		StartTime:   int(campaign.StartTime),
		EndTime:     int(campaign.EndTime),
		Placements:  int(campaign.Placements),
		AgentId:     campaign.AgentId,
		Category:    category.ConvertFromGrpc(campaign.Category),
		LastUpdated: campaign.LastUpdated.AsTime(),
		Ads:         ConvertMultipleAdsFromGrpc(campaign.Ads),
		Type:        model.PostType(campaign.Type),
	}
}

func (ac AdCategory) ConvertToGrpc() *protopb.AdCategory {
	return &protopb.AdCategory{
		Id:   ac.Id,
		Name: ac.Name,
	}
}

func (ac AdCategory) ConvertFromGrpc(adCategory *protopb.AdCategory) AdCategory {
	return AdCategory{
		Id:   adCategory.Id,
		Name: adCategory.Name,
	}
}

func (s CampaignStats) ConvertToGrpc() *protopb.CampaignStats {
	return &protopb.CampaignStats{
		Id:          s.Id,
		Name:        s.Name,
		IsOneTime:   s.IsOneTime,
		StartDate:   timestamppb.New(s.StartDate),
		EndDate:     timestamppb.New(s.EndDate),
		StartTime:   int32(s.StartTime),
		EndTime:     int32(s.EndTime),
		Placements:  int32(s.Placements),
		Category:    s.Category,
		Type:        s.Type,
		Influencers: ConvertMultipleInfluencerStatsToGrpc(s.Influencers),
		Likes:       int32(s.Likes),
		Dislikes:    int32(s.Dislikes),
		Comments:    int32(s.Comments),
		Clicks:      int32(s.Clicks),
	}
}

func (s InfluencerStats) ConvertToGrpc() *protopb.InfluencerStats {
	return &protopb.InfluencerStats{
		Id:            s.Id,
		Username:      s.Username,
		Ads:           ConvertMultipleAdStatsToGrpc(s.Ads),
		TotalLikes:    int32(s.TotalLikes),
		TotalDislikes: int32(s.TotalDislikes),
		TotalComments: int32(s.TotalComments),
		TotalClicks:   int32(s.TotalClicks),
	}
}

func (a AdStats) ConvertToGrpc() *protopb.AdStats {
	return &protopb.AdStats{
		Id:       a.Id,
		Media:    a.Media,
		Type:     a.Type,
		Hashtags: a.Hashtags,
		Location: a.Location,
		Likes:    int32(a.Likes),
		Dislikes: int32(a.Dislikes),
		Comments: int32(a.Comments),
		Clicks:   int32(a.Clicks),
	}
}

func ConvertMultipleAdStatsToGrpc(adStats []AdStats) []*protopb.AdStats {
	result := []*protopb.AdStats{}
	for _, ad := range adStats {
		result = append(result, ad.ConvertToGrpc())
	}

	return result
}

func ConvertMultipleInfluencerStatsToGrpc(influencerStats []InfluencerStats) []*protopb.InfluencerStats {
	result := []*protopb.InfluencerStats{}
	for _, stat := range influencerStats {
		result = append(result, stat.ConvertToGrpc())
	}

	return result
}

func (u CampaignInfluencerRequest) ConvertToGrpc() *protopb.CampaignInfluencerRequest {
	return &protopb.CampaignInfluencerRequest{
		Id:           u.Id,
		AgentId:      u.AgentId,
		InfluencerId: u.InfluencerId,
		CampaignId:   u.CampaignId,
		Status:       string(u.Status),
		PostAt:       timestamppb.New(u.PostAt),
	}
}

func (campaignRequest *CampaignInfluencerRequest) ConvertFromGrpc(n *protopb.CampaignInfluencerRequest) *CampaignInfluencerRequest {
	return &CampaignInfluencerRequest{
		Id:           n.Id,
		AgentId:      n.AgentId,
		InfluencerId: n.InfluencerId,
		CampaignId:   n.CampaignId,
		Status:       model.RequestStatus(n.Status),
		PostAt:       n.PostAt.AsTime(),
	}
}
