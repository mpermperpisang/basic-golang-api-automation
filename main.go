package main

import (
	"github.com/basic-golang-api-automation/features/helper"
	"github.com/joho/godotenv"
)

func init() {
	env := godotenv.Load()
	helper.LogPanicln(env)
}

func main() {
	// can be blank for now
}
