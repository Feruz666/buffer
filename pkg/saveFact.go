package pkg

import (
	"github.com/go-resty/resty/v2"
	"log"
)

const (
	bearerToken = "48ab34464a5573519725deb5865cc74c"

	keyPeriodStart         = "period_start"
	keyPeriodEnd           = "period_end"
	keyPeriodKey           = "period_key"
	keyIndicatorToMoId     = "indicator_to_mo_id"
	keyIndicatorToMoFactId = "indicator_to_mo_fact_id"
	keyValue               = "value"
	keyFactTime            = "fact_time"
	keyIsPlan              = "is_plan"
	keyAuthUserId          = "auth_user_id"
	keyComment             = "comment"

	address  = "https://development.kpi-drive.ru/_api/facts/save_fact"
	statusOK = "200 OK"
)

type FactsSaver struct {
	ss *resty.Client
}

func New(ss *resty.Client) *FactsSaver {
	return &FactsSaver{ss: ss}
}

func (f *FactsSaver) SaveFact(fact *SaveFactRequest) error {

	formData := map[string]string{
		keyPeriodStart:         fact.PeriodStart,
		keyPeriodEnd:           fact.PeriodEnd,
		keyPeriodKey:           fact.PeriodKey,
		keyIndicatorToMoId:     fact.IndicatorToMoId,
		keyIndicatorToMoFactId: fact.IndicatorToMoFactId,
		keyValue:               fact.Value,
		keyFactTime:            fact.FactTime,
		keyIsPlan:              fact.IsPlan,
		keyAuthUserId:          fact.AuthUserId,
		keyComment:             fact.Comment,
	}

	var result = &SaveFactResponse{}

	response, err := resty.New().R().
		SetHeader("Authorization", "Bearer "+bearerToken).
		SetResult(result).
		SetFormData(formData).
		Post(address)
	if err != nil {
		log.Fatalln("Error while making request, err:", err)
		return err
	}

	if response.Status() != statusOK {
		log.Println("Record has not created, status code: ", response.Status())
		return err
	} else {
		log.Println("Record has been created", response.Status())
		log.Println(result.Messages.Info[0])
	}

	return nil
}
