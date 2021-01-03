package client

import (
	"fmt"
	"net"
	"os"
	"strings"
	"sync"

	"github.com/JakeCooper/malcolm/pkg/proxy"
	"github.com/go-redis/redis"
)

type Malcolm struct {
	rdb       *redis.Client
	listeners map[string]chan bool
	lsmtx     sync.Mutex
}

const (
	GLOBAL_NAMESPACE = "malcolm"
	RULE_NAMESPACE   = "malcolm.rule."
)

func (m *Malcolm) addProxy(protocol string, url string, out string) error {
	done := make(chan bool)
	ln, err := net.Listen(protocol, url)
	if err != nil {
		return err
	}
	k := fmt.Sprintf("%s/%s", protocol, url)
	m.lsmtx.Lock()
	m.listeners[k] = done
	m.lsmtx.Unlock()
	go func() {
		for {
			select {
			case <-done:
				ln.Close()
				delete(m.listeners, k)
			default:
				connFrom, err := ln.Accept()
				if err != nil {
					fmt.Println(err)
				}
				connTo, err := net.Dial(protocol, out)
				if err != nil {
					fmt.Println(err)
				}
				c := proxy.New(&connFrom, &connTo)
				c.Proxy()
			}
		}
	}()
	return nil
}

// Creates a new malcolmn instance
func New() *Malcolm {
	// Construct Redis conn
	host := os.Getenv("REDISHOST")
	if host == "" {
		panic("Redis Host Required!")
	}

	port := os.Getenv("REDISPORT")
	if port == "" {
		panic("REDISPORT Required!")
	}

	pw := os.Getenv("REDISPASSWORD")

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: pw,
		DB:       0,
	})

	// Test conn
	if _, err := rdb.Ping().Result(); err != nil {
		panic(err)
	}

	m := &Malcolm{
		rdb:       rdb,
		lsmtx:     sync.Mutex{},
		listeners: make(map[string]chan bool),
	}

	rules, err := m.Rules()
	if err != nil {
		panic(err)
	}

	for in, out := range rules {
		us := len(in) - 4
		url, protocol := in[:us], in[us+1:]
		m.addProxy(protocol, url, out)
		fmt.Printf("%s Routing: %s -> %s\n", strings.ToUpper(protocol), url, out)
	}

	return m
}

func (m *Malcolm) Rules() (map[string]string, error) {
	mp := make(map[string]string)
	keys, err := m.rdb.Keys(fmt.Sprintf("%s*", RULE_NAMESPACE)).Result()
	if err != nil {
		return nil, err
	}
	if len(keys) == 0 {
		return mp, nil
	}
	values, err := m.rdb.MGet(keys...).Result()
	if err != nil {
		return nil, err
	}
	for i, key := range keys {
		mp[key[len(RULE_NAMESPACE):]] = values[i].(string)
	}
	return mp, nil
}

func (m *Malcolm) Health() (string, error) {
	return m.rdb.Ping().Result()
}
