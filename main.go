package main

import (
	"fmt"
	"github.com/sadysnaat/assignment1/response"
)

func main() {
	r := &response.Response{}

	res, err := r.WaitTillComplete()
	if err != nil {
		fmt.Println(err)
	}

	for _, result := range res {
		r.AddResult(result)
	}
	fmt.Println("results", r.GetResults())

	r2 := &response.Response{}

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
