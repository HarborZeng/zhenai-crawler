package persist

import (
	"context"
	"github.com/olivere/elastic"
	"log"
	"os/exec"
	"zhenai-crawler/crawler/common/reporter"
	"zhenai-crawler/crawler/model"
)

func ItemSaver() (chan model.Profile, error) {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		reporter.ReportError("创建elastic客户端出错", err)
		cmdResult := exec.Command("powershell", "scripts/fcheck_elastic_docker_run.ps1")
		if result, err := cmdResult.Output(); err != nil {
			return nil, err
		} else {
			log.Printf("%s\n", result)
		}
	}

	out := make(chan model.Profile)
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

func save(client *elastic.Client, item model.Profile) (id string, err error) {
	response, err := client.Index().
		Index("dating_profile").Type("zhenai").Id(item.Id).
		BodyJson(item).Do(context.Background())
	if err != nil {
		reporter.ReportError("es插入数据出错", err)
		return "", err
	}
	return response.Id, nil
}
