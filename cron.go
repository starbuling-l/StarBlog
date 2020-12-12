package main

import (
	"github.com/robfig/cron"
	"github.com/starbuling-l/StarBlog/models"
	"log"
	"time"
)

//实现定时器对软删除的文章进行清理
func main() {
	log.Println("Start cron >>>>>>")
	c := cron.New()

	c.AddFunc("* * * * * *", func() {
		log.Println("Clean tags >>>>>")
		models.CleanAllTags()
	})
	c.AddFunc("* * * * * *", func() {
		log.Println("Clean Articles >>>>>")
		models.CleanAllArticles()
	})

	c.Start()

	//会创建一个新的定时器，持续你设定的时间 d 后发送一个 channel 消息
	t := time.NewTimer(10 * time.Second)
	//for + select 阻塞 select 等待 channel
	for {
		select {
		case <-t.C:
			t.Reset(10 * time.Second)
		}
	}

}
