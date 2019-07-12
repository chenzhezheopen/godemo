package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

type Error string

func (err Error) Error() string {
	return string(err)
}
func main() {

	n, _ := redis.Dial("tcp", "localhost:6379")
	n.Do("auth", "abcdefg")

	reply, ok := n.Do("get", "oppo")
	switch reply.(type) {
	case int64:
		fmt.Println("int64")
		break
	case []byte:
		fmt.Println("[]byte")
		break
	case nil:
		fmt.Println("nil")
		break
	case Error:
		fmt.Println("error")
		break
	}
	fmt.Println(ok)
	// oppo, _ := n.Do("get", "oppo")
	// oppo, _ = json.Marshal(oppo)
	// fmt.Println(oppo)
	// fmt.Println(bol)
	boll, _ := redis.Bool(n.Do("hgetall", "36673161306"))
	fmt.Println(boll)
	fmt.Println(n.Do("hgetall", "36673161306"))
}
