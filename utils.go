package main

import "fmt"

func PPrintln(a ...any) {

	if IsDebug() || !IsProduction() {
		fmt.Println(a...)
	}
}
