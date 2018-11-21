package persist

import (
	"context"
	"gopkg.in/olivere/elastic.v5"
	"log"
	"zhenai-crawler/crawler/common/reporter"
	"zhenai-crawler/crawler/model"
)

func ItemSaver() (chan interface{}, error) {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		reporter.ReportError("创建elastic客户端出错", err)
		return nil, err
	}

	out := make(chan interface{})
	go func() {
		itemCounter := 0
		for {
			item := <-out
			itemCounter++
			//log.Printf("ItemSaver #%d: %+v\n", itemCounter, item)
			id, err := save(client, item)
			if err != nil {
				panic(err)
			}
			log.Printf("Item %s saved", id)
		}
	}()
	return out, nil
}

func save(client *elastic.Client, item interface{}) (id string, err error) {
	switch item.(type) {
	case model.Profile:
		response, err := client.Index().
			Index("dating_profile").Type("zhenai").Id(item.(model.Profile).Id).
			BodyJson(item).Do(context.Background())
		if err != nil {
			reporter.ReportError("es插入数据出错", err)
			return "", err
		}
		return response.Id, nil
	default:
		return "", err
	}
}
