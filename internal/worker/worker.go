package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Feruz666/buffer/internal/store"
	"github.com/Feruz666/buffer/pkg"
	"github.com/redis/go-redis/v9"
	"log"
)

type Worker struct {
	rdb  *redis.Client
	pkgs *pkg.FactsSaver
}

func New(rdb *redis.Client, pkgs *pkg.FactsSaver) *Worker {
	return &Worker{
		rdb:  rdb,
		pkgs: pkgs,
	}
}

func (w *Worker) SaveFacts(ctx context.Context) {
	log.Println("Worker is working")

	keys, err := w.rdb.Keys(ctx, "*").Result()
	if err != nil {
		log.Println("Error while getting keys from redis, err", err)
	}

	if len(keys) > 0 {
		var valuesFromRedis = make([]store.BufferData, len(keys))

		for _, key := range keys {
			bData := store.BufferData{}

			val, err := w.rdb.Get(ctx, key).Result()
			if err != nil {
				log.Println("Error while getting data from cache: ", err)
			}

			err = json.Unmarshal([]byte(val), &bData)
			if err != nil {
				log.Println("Error while unmarshalling: ", err)
			}

			valuesFromRedis = append(valuesFromRedis, bData)

			_, err = w.rdb.Del(ctx, key).Result()
			if err != nil {
				log.Println("Error while deleting data from cache: ", err)
			}

		}

		for _, valFromRedis := range valuesFromRedis {
			// Filtering
			err := validateBufferData(valFromRedis)
			if err == nil {
				err = w.pkgs.SaveFact(&pkg.SaveFactRequest{
					PeriodStart:         valFromRedis.PeriodStart,
					PeriodEnd:           valFromRedis.PeriodEnd,
					PeriodKey:           valFromRedis.PeriodKey,
					IndicatorToMoId:     valFromRedis.IndicatorToMoId,
					IndicatorToMoFactId: valFromRedis.IndicatorToMoFactId,
					Value:               valFromRedis.Value,
					FactTime:            valFromRedis.FactTime,
					IsPlan:              valFromRedis.IsPlan,
					AuthUserId:          valFromRedis.AuthUserId,
					Comment:             valFromRedis.Comment,
				})
				if err != nil {
					log.Println("pkg.SaveFact error:", err)
					return
				}
			}
		}

	} else {
		return
	}

}

func validateField(field string, fieldName string) error {
	if field == "" {
		return fmt.Errorf("field %s is empty", fieldName)
	}
	return nil
}

func validateBufferData(data store.BufferData) error {
	if err := validateField(data.PeriodStart, "PeriodStart"); err != nil {
		return err
	}
	if err := validateField(data.PeriodEnd, "PeriodEnd"); err != nil {
		return err
	}
	if err := validateField(data.PeriodKey, "PeriodKey"); err != nil {
		return err
	}
	if err := validateField(data.IndicatorToMoId, "IndicatorToMoId"); err != nil {
		return err
	}
	if err := validateField(data.IndicatorToMoFactId, "IndicatorToMoFactId"); err != nil {
		return err
	}
	if err := validateField(data.Value, "Value"); err != nil {
		return err
	}
	if err := validateField(data.FactTime, "FactTime"); err != nil {
		return err
	}
	if err := validateField(data.IsPlan, "IsPlan"); err != nil {
		return err
	}
	if err := validateField(data.AuthUserId, "AuthUserId"); err != nil {
		return err
	}
	if err := validateField(data.Comment, "Comment"); err != nil {
		return err
	}
	return nil
}
