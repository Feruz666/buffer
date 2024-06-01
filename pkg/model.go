package pkg

type SaveFactRequest struct {
	PeriodStart         string `form:"period_start" binding:"required"`
	PeriodEnd           string `form:"period_end" binding:"required"`
	PeriodKey           string `form:"period_key" binding:"required"`
	IndicatorToMoId     string `form:"indicator_to_mo_id" binding:"required"`
	IndicatorToMoFactId string `form:"indicator_to_mo_fact_id" binding:"required"`
	Value               string `form:"value" binding:"required"`
	FactTime            string `form:"fact_time" binding:"required"`
	IsPlan              string `form:"is_plan" binding:"required"`
	AuthUserId          string `form:"auth_user_id" binding:"required"`
	Comment             string `form:"comment" binding:"required"`
}

type SaveFactResponse struct {
	Messages struct {
		Error   interface{} `json:"error"`
		Warning interface{} `json:"warning"`
		Info    []string    `json:"info"`
	} `json:"MESSAGES"`
	Data struct {
		IndicatorToMoFactID int `json:"indicator_to_mo_fact_id"`
	} `json:"DATA"`
	Status string `json:"STATUS"`
}
