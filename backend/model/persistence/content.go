package persistence

type Post struct{
	Id          string
	UserId      string
	IsAd        bool
	Type        PostType
	Description string
	Location    string
	DateCreated string
}

type Story struct{
	Post
	IsCloseFriends bool
}

type Media struct{
	Id      string
	Type    MediaType
	PostId  string
	Content string
}

type Tag struct {
	UserId string
	MediaId string
}

type Collection struct {
	Name string
	Id string
	UserId string
}

type Favorites struct {
	PostId string
	UserId string
	CollectionId string
}

type PostLikes struct {
	PostId string
	UserId string
	IsLike bool
}

type PostComments struct {
	PostId string
	UserId string
	Content string
	DateCreated string //TODO
}

type HighLights struct {
	UserId string
	Id string
	Name string
}

type RegistrationRequest struct {
	Id        string
	UserId    string
	Timestamp string //TODO
	Status    RequestStatus
}

type Ad struct {
	Id string
	Link string
	CampaignId string
	PostId string
	LinkClickNum int
}

type Campaign struct {
	Id string
	IsOneTime bool
	StartDate string //TODO
	EndDate string //TODO
	PlacementNum int
	AgentId string
	IdAdCategory string
	LastUpdated string //TODO
}

type CampaignInfluencerRequest struct {
	CampaignId   string
	InfluencerId string
	Status       RequestStatus
}

type ContentComplaint struct {
	Id       string
	Category ComplaintCategory
	PostId   string
	Status   RequestStatus
}

type AdCategory struct {
	Id string
	Name string
}

type UserAdCategories struct {
	UserId string
	IdAdCategory string
}

type CampaignChanges struct {
	CampaignId string
	AdCategoryId string
	StartDate string
	EndDate string
	PlacementNum int
}



