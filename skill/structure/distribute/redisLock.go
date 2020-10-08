package distribute

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"math/rand"
	"sync"
	"time"
)

func NewLockRedis(lockkey string, timeout int, client *redis.Client) Lock {
	s := rand.NewSource(time.Now().UnixNano())
	token := rand.New(s).Intn(1000000)
	return &lockRedis{
		LockKey:     lockkey,
		TimeOut:     timeout,
		Token:       token,
		RedisClient: client,
	}
}

type lockRedis struct {
	LockKey     string
	TimeOut     int
	Token       int
	RedisClient *redis.Client
}

/**
加锁
若发生竞态，延时队列获取解锁消息，再次加锁
*/
func (l *lockRedis) Lock() error {
	r, err := l.RedisClient.Do("SET", l.LockKey, l.Token, "EX", l.TimeOut, "NX").Result()
	if err != nil {
		return err
	}
	if r == nil {
		for {
			q, err := l.RedisClient.BLPop(time.Second, l.LockKey+"_queue").Result()
			if err != nil {
				return err
			}
			if len(q) > 0 {
				l.RedisClient.Del(l.LockKey + "_queue")
				return l.Lock()
			}
		}
	}
	return nil
}

/**
解锁，并入队通知其余客户端
*/
func (l *lockRedis) Unlock() error {
	value, err := l.RedisClient.Get(l.LockKey).Result()
	if err != nil {
		return err
	}
	if value != fmt.Sprintf("%d", l.Token) {
		return errors.New("unlock client not lock client")
	}
	err = l.RedisClient.Del(l.LockKey).Err()
	if err != nil {
		_, err = l.RedisClient.RPush(l.LockKey+"_queue", 1).Result()
	}

	return err
}

/**
业务代码处理
业务代码超时会自动续期
执行成功或失败，会释放锁
*/
func (l *lockRedis) Proccess(dealFunc func() error) error {
	defer l.Unlock()

	done := make(chan bool, 1)
	var err error
	go func() {
		err = dealFunc()
		done <- true
	}()

	timeInterval := (l.TimeOut / 10) * 8
	ticker := time.NewTicker(time.Second * time.Duration(timeInterval))
	for {
		var isOver bool
		select {
		case r := <-done:
			if r {
				isOver = true
			}
		case <-ticker.C:
			_, err = l.RedisClient.Expire(l.LockKey, time.Duration(l.TimeOut)).Result()
		}
		if isOver {
			break
		}
	}

	return err
}
