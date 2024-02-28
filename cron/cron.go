package cron

import (
	"github.com/robfig/cron/v3"
	"log"
)

func newWithSeconds() *cron.Cron {
	secondParser := cron.NewParser(cron.Second | cron.Minute |
		cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor)
	return cron.New(cron.WithParser(secondParser), cron.WithChain())
}
func SetUp() {
	go func() {
		crontab := newWithSeconds()

		//if _, err := crontab.AddFunc("0 0 */1 * * ?", CheckGptSk); err != nil { //每小时一次
		//	log.Print("定时任务: checkGptSk failed" + err.Error())
		//}

		crontab.Start()
		log.Println("cron start success")
	}()
}
