package store

type BufferData struct {
	PeriodStart         string `json:"period_start"`
	PeriodEnd           string `json:"period_end"`
	PeriodKey           string `json:"period_key"`
	IndicatorToMoId     string `json:"indicator_to_mo_id"`
	IndicatorToMoFactId string `json:"indicator_to_mo_fact_id"`
	Value               string `json:"value"`
	FactTime            string `json:"fact_time"`
	IsPlan              string `json:"is_plan"`
	AuthUserId          string `json:"auth_user_id"`
	Comment             string `json:"comment"`
}
