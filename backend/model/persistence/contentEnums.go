package persistence

type PostType string
const(
	TypePost  PostType = "Post"
	TypeStory          = "Story"
)

type MediaType string
const(
	TypeImage MediaType = "Image"
	TypeVideo           = "Video"
)

type RequestStatus string
const(
	Pending  RequestStatus = "Pending"
	Accepted               = "Accepted"
	Rejected               = "Rejected"
)

type ComplaintCategory string
const(
	Gore       ComplaintCategory = "Gore"
	Nudity                       = "Nudity"
	Violence                     = "Violence"
	Suicide                      = "Suicide"
	FakeNews                     = "Fake News"
	Spam                         = "Spam"
	HateSpeech                   = "Hate Speech"
	Terrorism                    = "Terrorism"
	Harassment                   = "Harassment"
	Other                        = "Other"
)