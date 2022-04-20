
package dto

type ImpactStat struct {
	Impact   string `json:"impact"`
	Count int `json:"count"`
}


type FoundStat struct {
	TypeFound   string `json:"typeFound"`
	Count int `json:"count"`
}
