package main

import (
	"encoding/json"
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	c, _ := redis.Dial("tcp", "localhost:6379")
	c.Do("auth", "abcdefg")
	// c.Do("set", "po", "15")
	c.Do("set", "hihihi10", "你你你你你你你您你你in你您你i")
	nn, _ := c.Do("get", "hihihi")
	mi, _ := json.Marshal(nn)
	// mi := string(nn)
	fmt.Println(string(mi))
	// fmt.Println(redis.String(c.Do("get", "po")))
	// t, _ := redis.Bool(c.Do("exists", "po"))
	// defer c.Close()
	// fmt.Println(t)
	// imap := map[string]string{"name": "waiwaigo", "phone": "13498739038"}
	// value, _ := json.Marshal(imap)
	// fmt.Println(value)
	// fmt.Println(string(value))

	// psc := redis.PubSubConn{c}
	// psc.Subscribe("example")
	// for {
	// 	switch v := psc.Receive().(type) {
	// 	case redis.Message:
	// 		fmt.Printf("%s: message: %s\n", v.Channel, v.Data)
	// 	case redis.Subscription:
	// 		fmt.Printf("sub%s: %s %d\n", v.Channel, v.Kind, v.Count)
	// 	case error:
	// 		fmt.Println("err")
	// 		break
	// 	}
	// }
}
