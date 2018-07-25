// 146. LRU Cache
// https://leetcode.com/problems/lru-cache/description/

package main

import (
	"fmt"
)

type Value struct {
	key, value             int
	moreRecent, lessRecent *Value
}

type LRUCache struct {
	cap      int
	cache    map[int]*Value
	mru, lru *Value
}

func Constructor(cap int) LRUCache {
	return LRUCache{
		cap:   cap,
		cache: make(map[int]*Value),
		mru:   nil, lru: nil,
	}
}

func (c *LRUCache) touch(v *Value) {
	if c.mru == v {
		return
	}
	v.moreRecent.lessRecent = v.lessRecent
	if v.lessRecent == nil {
		c.lru = v.moreRecent
	} else {
		v.lessRecent.moreRecent = v.moreRecent
	}
	v.moreRecent, v.lessRecent = nil, c.mru
	c.mru, c.mru.moreRecent = v, v
}

func (c *LRUCache) Get(key int) int {
	v, ok := c.cache[key]
	if !ok {
		return -1
	}
	c.touch(v)
	return v.value
}

func (c *LRUCache) Put(key int, value int) {
	if c.cap < 1 {
		return
	}
	v, ok := c.cache[key]
	if ok {
		v.value = value
		c.touch(v)
		return
	}
	v = &Value{key, value, nil, nil}
	if len(c.cache) == c.cap {
		delete(c.cache, c.lru.key)
		c.lru = c.lru.moreRecent
		if c.lru != nil {
			c.lru.lessRecent = nil
		}
	}
	if c.cap > 1 && c.mru != nil {
		v.lessRecent, c.mru.moreRecent = c.mru, v
	}
	if c.lru == nil {
		c.lru = v
	}
	c.mru, c.cache[key] = v, v
}

func main() {
	c := Constructor(1)
	c.Put(1, 1)
	fmt.Printf("%v %v\n", c.mru, c.lru)
	fmt.Printf("%p %v\n", c.cache[1], c.cache[1])
	c.Put(2, 2)
	fmt.Printf("%v %v\n", c.mru, c.lru)
	fmt.Printf("%p %v\n", c.cache[1], c.cache[1])
	fmt.Printf("%p %v\n", c.cache[2], c.cache[2])
	c.Put(3, 3) // evicts key 2
	fmt.Printf("%v %v\n", c.mru, c.lru)
	fmt.Printf("%p %v\n", c.cache[1], c.cache[1])
	fmt.Printf("%p %v\n", c.cache[2], c.cache[2])
	fmt.Printf("%p %v\n", c.cache[3], c.cache[3])
	fmt.Println(c.Get(1)) // returns 1
	fmt.Println(c.Get(2)) // returns -1 (not found)
	fmt.Println(c.Get(3)) // returns -1 (not found)
	c.Put(4, 4)           // evicts key 1
	fmt.Println(c.Get(1)) // returns -1 (not found)
	fmt.Println(c.Get(2)) // returns -1
	fmt.Println(c.Get(3)) // returns 3
	fmt.Println(c.Get(4)) // returns 4
}
