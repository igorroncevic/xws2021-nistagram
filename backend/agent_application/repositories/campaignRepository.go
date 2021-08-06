package repositories

import (
	"context"
	"fmt"
	"github.com/beevik/etree"
	"github.com/igorroncevic/xws2021-nistagram/agent_application/model/domain"
	"github.com/igorroncevic/xws2021-nistagram/common/tracer"
	"gorm.io/gorm"
)

type CampaignRepository interface {
	CreateCampaignReport(context.Context, domain.CampaignStats) error
}

type campaignRepository struct {
	DB *gorm.DB
}

func NewCampaignRepo(db *gorm.DB) (*campaignRepository, error) {
	if db == nil {
		panic("CampaignRepository not created, gorm.DB is nil")
	}

	return &campaignRepository{DB: db}, nil
}

func (repository *campaignRepository) CreateCampaignReport(ctx context.Context, stats domain.CampaignStats) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateCampaignReport")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	doc := etree.NewDocument()
	doc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)
	//doc.CreateProcInst("xml-stylesheet", `type="text/xsl" href="style.xsl"`)
	campaign := doc.CreateElement("campaign")
	campaign.CreateComment("Root Campaign element")

	name := campaign.CreateElement("name")
	name.CreateText(stats.Name)

	durationType := campaign.CreateElement("duration")
	durationValue := ""
	if stats.IsOneTime {
		durationValue = "Short term"
	} else {
		durationValue = "Long term"
	}

	campaignType := stats.Type + " campaign"
	durationType.CreateText(durationValue + " " + campaignType)

	dates := campaign.CreateElement("dates")
	dates.CreateText("Campaign was placed between" + stats.StartDate.Format("02 January 2006") + " - " + stats.EndDate.Format("02 January 2006"))

	times := campaign.CreateElement("times")
	times.CreateText(fmt.Sprintf("Campaign is placed daily between %dh and %dh", stats.StartTime, stats.EndTime))

	placements := campaign.CreateElement("placements")
	placements.CreateText(fmt.Sprintf("Total number of placements: %d", stats.Placements))

	category := campaign.CreateElement("category")
	category.CreateText("Category: " + stats.Category)

	likes := campaign.CreateElement("likes")
	likes.CreateText(fmt.Sprintf("Total number of likes in campaign: %d", stats.Likes))

	dislikes := campaign.CreateElement("dislikes")
	dislikes.CreateText(fmt.Sprintf("Total number of dislikes in campaign: %d", stats.Dislikes))

	comments := campaign.CreateElement("comments")
	comments.CreateText(fmt.Sprintf("Total number of comments in campaign: %d", stats.Comments))

	influencers := campaign.CreateElement("influencers")
	influencers.CreateText("Influencer stats")
	for _, influencer := range stats.Influencers {
		singleInfluencer := influencers.CreateElement("influencer")
		singleInfluencer.CreateAttr("id", influencer.Id)

		username := singleInfluencer.CreateElement("username")
		username.CreateText("Username: " + influencer.Username)

		likes := singleInfluencer.CreateElement("likes")
		likes.CreateText(fmt.Sprintf("Total number of likes from this influencer: %d", influencer.TotalLikes))

		dislikes := singleInfluencer.CreateElement("dislikes")
		dislikes.CreateText(fmt.Sprintf("Total number of dislikes from this influencer: %d", influencer.TotalDislikes))

		comments := singleInfluencer.CreateElement("comments")
		comments.CreateText(fmt.Sprintf("Total number of comments from this influencer: %d", influencer.TotalComments))

		ads := singleInfluencer.CreateElement("ads")
		ads.CreateText(influencer.Username + "'s ads")
		for _, ad := range influencer.Ads {
			singleAd := ads.CreateElement("ad")
			singleAd.CreateAttr("id", ad.Id)

			adType := singleAd.CreateElement("ad_type")
			adType.CreateText("Type: " + ad.Type)

			location := singleAd.CreateElement("ad_location")
			location.CreateText("Type: " + ad.Location)

			hashtags := singleAd.CreateElement("hashtags")
			if len(ad.Hashtags) == 0 {
				hashtags.CreateText(" None")
			} else {
				hashtagsText := "Hashtags:"
				for _, hashtag := range ad.Hashtags {
					hashtagsText += " #" + hashtag
				}
				hashtags.CreateText(hashtagsText)
			}

			if ad.Type == "Post" {
				likes := singleAd.CreateElement("likes")
				likes.CreateText(fmt.Sprintf("Number of likes on this ad: %d", ad.Likes))

				dislikes := singleAd.CreateElement("dislikes")
				dislikes.CreateText(fmt.Sprintf("Number of dislikes on this ad: %d", ad.Dislikes))

				comments := singleAd.CreateElement("comments")
				comments.CreateText(fmt.Sprintf("Number of comments on this ad: %d", ad.Comments))
			}
		}
	}

	return nil
}
