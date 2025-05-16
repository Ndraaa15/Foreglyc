package scheduler

import (
	"fmt"
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

func New() *cron.Cron {
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		log.Fatal(fmt.Sprintf("failed to load timezone: %v", err))
	}

	return cron.New(cron.WithLocation(loc), cron.WithChain(cron.Recover(cron.DefaultLogger)))

}
