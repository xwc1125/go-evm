// description: joinchain 
// 
// @author: xwc1125
// @date: 2019/11/13
package test

import (
	"fmt"
	lru "github.com/hashicorp/golang-lru"
	"testing"
)

func TestLru(t *testing.T) {
	kvch, err := lru.New(100)
	fmt.Println(kvch, err)

	kvch.Add("1", "9")
	kvch.Add("2", "8")
	kvch.Add("3", "7")
	kvch.Add("4", "6")
	kvch.Add("5", "5")

	fmt.Println(kvch.Get("1"))
	fmt.Println(kvch.Get("2"))
	fmt.Println(kvch.Get("3"))
	kvch.Add("5", "6")
	fmt.Println(kvch.Get("3"))
}
