package main

import "../rulebook/aws"

func main() {
	for _, service := range aws.InstMap {
		service.ExecuteRules(" ")
	}
}
