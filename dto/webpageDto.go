package dto

import (
	"Accessibility-Backend/entity"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WebpagesResult struct {
	Data []WebpageResponseBody `json:"data"`
	TotalCount int `json:"totalCount"`
	Page int `json:"page"`
	LastPage int `json:"lastPage"`
}

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

type WebpageFullResponseBodyNew struct {
	ID     string             `json:"id"`
	Name     string             `json:"name"`
	Url      string             `json:"url"`
	ScanTime string             `json:"scanTime"`
	Note     string             `json:"note"`
	Issue    []entity.Issue `json:"issues"`
	Website  entity.Website            `json:"website"`
	ImpactStats []ImpactStatNew             `json:"impactStatistics"`
	FoundStats []FoundStatNew            `json:"foundStatistics"`
}

type WebpageResponseBody struct {
	ID     string             `json:"id"`
	Name     string             `json:"name,omitempty"`
	Url      string             `json:"url,omitempty"`
	ScanTime string             `json:"scanTime,omitempty"`
	Website  string            `json:"website,omitempty"`
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
func (data Wp) ToWebpageFullResponse() *WebpageFullResponseBody {
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

func (data Wp) ToWebpageResponse() *WebpageResponseBody {
	return &WebpageResponseBody{
		ID: data.ID.Hex(),
		Name:        data.Name,
		Url:       data.Url,
		ScanTime: data.ScanTime,
		Website:    data.Website.Name,
	}
}