package dto

import (
	"Accessibility-Backend/entity"

	"go.mongodb.org/mongo-driver/bson/primitive"
)



type WebpageRequestBody struct {
	Name     string             `json:"name" validate:"required"`
	Url      string             `json:"url" validate:"required"`
	ScanTime string             `json:"scanTime" validate:"required"`
	Note     string             `json:"note"`
	Issue    []IssueRequestBody `json:"issues" validate:"required"`
	Website  WebsiteRequestBody            `json:"website" validate:"required"`
}

type WebpageFullResponseBody struct {
	ID     string             `json:"id"`
	Name     string             `json:"name"`
	Url      string             `json:"url"`
	ScanTime string             `json:"scanTime"`
	Note     string             `json:"note"`
	Issue    []entity.Issue `json:"issues"`
	Website  entity.Website            `json:"website"`
	ImpactStats ImpactStat             `json:"impactStatistics"`
	FoundStats FoundStat             `json:"foundStatistics"`
}

type WebpageResponseBody struct {
	ID     string             `json:"id"`
	Name     string             `json:"name"`
	Url      string             `json:"url"`
	ScanTime string             `json:"scanTime"`
	Website  WebsiteRequestBody            `json:"website"`
}

func (data WebpageRequestBody) ToWebpageEntities() *entity.Webpage {
	return &entity.Webpage{
		ID: primitive.NewObjectID(),
		Name:        data.Name,
		Url:       data.Url,
		ScanTime: data.ScanTime,
		Note: data.Note,
		Issue:    *GetIssueEntities(data.Issue),
		Website:    *data.Website.ToWebsiteEntities(),
	}
}

type Wp entity.Webpage
func (data Wp) ToWebpageResponse() *WebpageFullResponseBody {
	return &WebpageFullResponseBody{
		ID: data.ID.String(),
		Name:        data.Name,
		Url:       data.Url,
		ScanTime: data.ScanTime,
		Note: data.Note,
		Issue:    data.Issue,
		Website:    data.Website,
	}
}