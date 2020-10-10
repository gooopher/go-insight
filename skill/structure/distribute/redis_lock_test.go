package distribute

import (
	"fmt"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/encoder/yaml"
	"github.com/micro/go-micro/config/source"
	"github.com/micro/go-micro/config/source/file"
	"go-insight/driver"
	"log"
	"sync"
	"testing"
)

func TestLockRedis(t *testing.T) {
	err := config.Load(
		file.NewSource(
			file.WithPath("../../../config/redis.yaml"),
			source.WithEncoder(yaml.NewEncoder()),
		),
	)
	if err != nil {
		log.Fatalf("load config file fail: %s", err.Error())
	}

	redisClient, err := driver.GetRedis("base")
	if err != nil {
		log.Fatalf("redis config file fail: %s", err.Error())
	}

	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			lockIns := NewLockRedis("book", 1, redisClient)
			err := lockIns.Lock()
			if err != nil {
				log.Fatalf("lock fail: %s", err.Error())
			}
			err = lockIns.Proccess(func() error {
				fmt.Printf("task : %d success\n", i)
				return nil
			})
			if err != nil {
				log.Fatalf("deal fail: %s", err.Error())
			}
			wg.Done()
		}(i)
	}

	wg.Wait()
}
