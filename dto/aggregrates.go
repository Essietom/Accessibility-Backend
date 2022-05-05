package dto

type ImpactStat struct {
	Critical   int `json:"critical"`
	Serious int `json:"serious"`
	Moderate int `json:"moderate"`
	Minor int `json:"minor"`
	ImpactTotal int `json:"impactTotal"`
}

type FoundStat struct {
	Automatic   int `json:"automatic"`
	Guided int `json:"guided"`
	NeedsReview int `json:"needsReview"`
	FoundTotal int `json:"foundTotal"`
}
