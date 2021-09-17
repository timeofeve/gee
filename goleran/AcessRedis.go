package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {

	// redis connection
	c, err := redis.Dial("tcp", "10.151.3.131:6379")
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return
	}

	// fmt.Println("redis conn success")
	defer c.Close()

	// @ string set and get 操作
	// // redis set value
	// _, err = c.Do("Set", "abc", 100)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// // redis get value
	// r, err := redis.Int(c.Do("Get", "abc"))
	// if err != nil {
	// 	fmt.Println("get abc failed,", err)
	// 	return
	// }

	// fmt.Println(r)

	// @批量操作string
	_, err = c.Do("MSet", "abc2", 100, "efg", 300)
	if err != nil {
		fmt.Println(err)
		return
	}

	r, err := redis.Ints(c.Do("MGet", "abc2", "efg"))
	if err != nil {
		fmt.Println("get abc failed,", err)
		return
	}

	for _, v := range r {
		fmt.Println(v)
	}

}
