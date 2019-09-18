package main

import (
	"fmt"
	"github.com/sadysnaat/assignment1/response"
)

func main() {
	r := &response.Response{}

	fmt.Println("Doing some heavy work, please wait a while")
	res, err := r.WaitTillComplete()
	if err != nil {
		fmt.Println(err)
	}

	for _, result := range res {
		r.AddResult(result)
	}
	fmt.Println("results", r.GetResults())

	r2 := &response.Response{}

	fmt.Println("Doing the same heavy work but async, expect result as soon as they are ready")
	r2.Subscribe(func(result response.Result) {
		fmt.Println("got", result.Value)
		r2.AddResult(result)
	}, func(err error) {
		fmt.Println("error occurred")
		r2.SetError(err)
	}, func(completed bool) {
		fmt.Println("completed")
		r2.SetComplete()
	})

	for {
		if r2.Completed() || r2.Error() != nil {
			break
		}
	}

	fmt.Println(r2.GetResults())
}
