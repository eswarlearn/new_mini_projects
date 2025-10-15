package main

import "fmt"

func LoopWord(wd string, nu int) {

	for i := 0; i < nu; i++ {
		fmt.Println(wd, " ... ", i)
	}
}
