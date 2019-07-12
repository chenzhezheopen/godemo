package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/garyburd/redigo/redis"
)

type mCookie struct {
	mCookie     string
	mSession    map[string]*mSession
	maxlifetime int64
	mlock       sync.RWMutex
}
type mSession struct {
	Name    string
	nowTime time.Time
	mValues map[interface{}]interface{}
}
type mredis struct {
	now    string
	Name   string
	Value  string
	Path   string
	MaxAge string
}

var Logger *log.Logger
var mag *mCookie
var c redis.Conn
var now = time.Now()

func init() {

	var cName string
	cName = "chenzhe"
	mag = maxage(cName, 3600)
}
func main() {
	c, _ = redis.Dial("tcp", "localhost:6379")
	serve := &http.Server{
		Addr: ":4545",
	}
	http.HandleFunc("/demo", rediscookie)
	http.HandleFunc("/accept", acceptcookie)
	serve.ListenAndServe()
}

func acceptcookie(w http.ResponseWriter, r *http.Request) {
	rand.Seed(now.Unix())
	ran := rand.Intn(1000000000000)
	fmt.Println(ran)
	display, err := c.Do("hgetall", ran)
	if err != nil {
		fmt.Println("没有找到对应的数据")
		return
	}

	oi, _ := json.Marshal(display)
	biu := string(oi[:])
	fmt.Println(biu)
}
func rediscookie(w http.ResponseWriter, r *http.Request) {
	cookie := presever(r.Host, time.Now())
	http.SetCookie(w, &cookie)
	// fmt.Println(c)
}

// "Name", "cccc", "Value", h, "Path", "/", "MaxAge", int(mag.maxlifetime),
//
func presever(h string, t time.Time) http.Cookie {
	c.Do("auth", "abcdefg")
	now = time.Now()
	rand.Seed(now.Unix())
	var ran = rand.Intn(1000000000000)
	pan, _ := redis.Bool(c.Do("HMSET", ran, "now", time.Now(), "Name", ran, "Value", h, "Path", "/", "MaxAge", int(mag.maxlifetime)))
	fmt.Println(pan)
	logfile, _ := os.OpenFile("test.txt", os.O_RDWR|os.O_APPEND, 0)
	Logger = log.New(logfile, "\r\n", log.Ldate|log.Lshortfile)
	sessionID := &mSession{
		Name:    "ccccc",
		nowTime: time.Now(),
		mValues: make(map[interface{}]interface{}),
	}
	str := strconv.Itoa(ran)
	mag.mSession[str] = sessionID
	Logger.Println(str, pan)
	cookie := http.Cookie{Name: str, Value: h, Path: "/", MaxAge: int(mag.maxlifetime), HttpOnly: true}
	return cookie
}
func maxage(n string, t int64) *mCookie {

	mar := &mCookie{mCookie: n, maxlifetime: t, mSession: make(map[string]*mSession)}

	return mar
}
