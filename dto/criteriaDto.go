package dto

import "Accessibility-Backend/entity"

type CriteriaRequestBody struct {
	ID   string `json:"criteriaId" bson:"_id,omitempty"`
	Name string `json:"name" validate:"required"`
	Note string `json:"note"`
}


func (data CriteriaRequestBody) ToCriteriaEntities() *entity.Criteria {
	return &entity.Criteria{
		ID:        data.ID,
		Name:       data.Name,
		Note: data.Note,
	}
}

func  GetIssueCriteria(criterias []CriteriaRequestBody) *[]entity.Criteria{
	var criteriaEntity []entity.Criteria
	for _, crit := range criterias {	
		criteriaEntity = append(criteriaEntity, *crit.ToCriteriaEntities())
	}
	return &criteriaEntity
}