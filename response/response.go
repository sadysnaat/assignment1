package response

import "time"

type Responder interface {
	WaitTillComplete() ([]Result, error)
	Subscribe(func(res Result), func(err error), func(completed bool))
	Completed() bool
	SetComplete()
	SetError(err error)
	Error() error
	AddResult(result Result)
	GetResults() []Result
}

type Result struct {
	Value int
}

type Response struct {
	results   []Result
	completed bool
	err       error
}

func (r *Response) Completed() bool {
	return r.completed
}

func (r *Response) SetComplete() {
	r.completed = true
}

func (r *Response) SetError(err error) {
	r.err = err
}

func (r *Response) Error() error {
	return r.err
}

func (r *Response) AddResult(result Result) {
	r.results = append(r.results, result)
}

func (r *Response) GetResults() []Result {
	return r.results
}

func (r *Response) WaitTillComplete() ([]Result, error) {
	var results []Result
	for i := 0; i < 10; i++ {
		result := Result{
			Value: i,
		}
		results = append(results, result)
		time.Sleep(time.Second * 1)
	}
	return results, nil
}

func (r *Response) Subscribe(recb func(result Result), ercb func(err error), comcb func(completed bool)) {
	go func() {
		for i := 0; i < 10; i++ {
			result := Result{
				Value: i,
			}
			go recb(result)
			time.Sleep(time.Millisecond * 100)
		}
		go comcb(true)
	}()
}
