package redis

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"flag"
	redisDriver "github.com/garyburd/redigo/redis"
	"log"
)

var POOL redisDriver.Pool
var (
	redisServer   = flag.String("redisServer", ":6379", "")
	redisPassword = flag.String("redisPassword", "", "")
)

const (
	LizaKey    = "liza:"
	LizaJobKey = "liza:jobs"
)

func init() {
	POOL = redisDriver.Pool{
		MaxIdle:   50,
		MaxActive: 500,
		Dial: func() (redisDriver.Conn, error) {
			c, err := redisDriver.Dial("tcp", ":6379")
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}

func GetURL(URL string) (string, bool) {
	hashString := computeHmac256(URL, "url_secret")
	c := POOL.Get()
	defer c.Close()
	log.Println(hashString)
	if urlInterface, err := c.Do("GET", LizaKey+hashString); err != nil {
		log.Fatal("Redis Error: ", err)
		return "", false
	} else {
		urlString, _ := redisDriver.String(urlInterface, err)
		if len(urlString) > 0 {
			log.Println("Redis entry found: ", urlString)
			return urlString, true
		} else {
			log.Println("No redis entry found: ", urlString)
			return "", false
		}

	}
	return "", false
}

func AddURL(URL, imageURL string) bool {
	hashString := computeHmac256(URL, "url_secret")
	c := POOL.Get()
	defer c.Close()
	log.Println("Key hash ", hashString)
	if _, err := c.Do("SET", LizaKey+hashString, imageURL); err != nil {
		log.Fatal("Redis Error: ", err)
		return false
	}
	return true
}

func QueueJob(URL string) bool {
	c := POOL.Get()
	defer c.Close()
	if _, err := c.Do("LPUSH", LizaJobKey, URL); err != nil {
		log.Fatal("Redis Error: ", err)
		return false
	}
	return true
}

func GetJob() (string, bool) {
	c := POOL.Get()
	defer c.Close()
	if urlInterface, err := c.Do("RPOP", LizaJobKey); err != nil {
		log.Fatal("Redis Error: ", err)
		return "", false
	} else {
		urlString, _ := redisDriver.String(urlInterface, err)
		log.Println("Redis entry found: ", urlString)
		return urlString, true
	}
	return "", false
}

func computeHmac256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
