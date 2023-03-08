package main

import "github.com/cqqqq777/simple_oauth2/app/boot"

func main() {
	boot.ViperInit()
	boot.DatabaseInit()
	boot.RoutersInit()
}
