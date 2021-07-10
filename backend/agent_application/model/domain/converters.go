package domain

import (
	protopb "github.com/david-drvar/xws2021-nistagram/common/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (l LoginRequest) ConvertFromGrpc(request *protopb.LoginRequestAgentApp) LoginRequest {
	return LoginRequest{
		Email:    request.Email,
		Password: request.Password,
	}
}

func (s CampaignStats) ConvertToGrpc() *protopb.CampaignStats{
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

func (s InfluencerStats) ConvertToGrpc() *protopb.InfluencerStats{
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

func (a AdStats) ConvertToGrpc() *protopb.AdStats{
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

func ConvertMultipleAdStatsToGrpc(adStats []AdStats) []*protopb.AdStats{
	result := []*protopb.AdStats{}
	for _, ad := range adStats {
		result =append(result, ad.ConvertToGrpc())
	}

	return result
}

func ConvertMultipleInfluencerStatsToGrpc(influencerStats []InfluencerStats) []*protopb.InfluencerStats{
	result := []*protopb.InfluencerStats{}
	for _, stat := range influencerStats {
		result =append(result, stat.ConvertToGrpc())
	}

	return result
}

// FromGrpc
func (s CampaignStats) ConvertFromGrpc(stats *protopb.CampaignStats) CampaignStats{
	return CampaignStats{
		Id:          stats.Id,
		Name:        stats.Name,
		IsOneTime:   stats.IsOneTime,
		StartDate:   stats.StartDate.AsTime(),
		EndDate:     stats.EndDate.AsTime(),
		StartTime:   int(stats.StartTime),
		EndTime:     int(stats.EndTime),
		Placements:  int(stats.Placements),
		Category:    stats.Category,
		Type:        stats.Type,
		Influencers: ConvertMultipleInfluencerStatsFromGrpc(stats.Influencers),
		Likes:       int(stats.Likes),
		Dislikes:    int(stats.Dislikes),
		Comments:    int(stats.Comments),
		Clicks:      int(stats.Clicks),
	}
}

func (s InfluencerStats) ConvertFromGrpc(stats *protopb.InfluencerStats) InfluencerStats{
	return InfluencerStats{
		Id:            stats.Id,
		Username:      stats.Username,
		Ads:           ConvertMultipleAdStatsFromGrpc(stats.Ads),
		TotalLikes:    int(stats.TotalLikes),
		TotalDislikes: int(stats.TotalDislikes),
		TotalComments: int(stats.TotalComments),
		TotalClicks:   int(stats.TotalClicks),
	}
}

func (a AdStats) ConvertFromGrpc(stats *protopb.AdStats) AdStats{
	return AdStats{
		Id:       stats.Id,
		Media:    stats.Media,
		Type:     stats.Type,
		Hashtags: stats.Hashtags,
		Location: stats.Location,
		Likes:    int(stats.Likes),
		Dislikes: int(stats.Dislikes),
		Comments: int(stats.Comments),
		Clicks:   int(stats.Clicks),
	}
}

func ConvertMultipleAdStatsFromGrpc(adStats []*protopb.AdStats) []AdStats{
	result := []AdStats{}
	for _, ad := range adStats {
		var domainStats AdStats
		result = append(result, domainStats.ConvertFromGrpc(ad))
	}

	return result
}

func ConvertMultipleInfluencerStatsFromGrpc(influencerStats []*protopb.InfluencerStats) []InfluencerStats{
	result := []InfluencerStats{}
	for _, stat := range influencerStats {
		var domainStats InfluencerStats
		result =append(result, domainStats.ConvertFromGrpc(stat))
	}

	return result
}