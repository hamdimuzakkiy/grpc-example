package main

import (
	"errors"
	"log"
)

func channel(i int) error {
	errC := make(chan error)
	count := 0

	count ++
	go func(num int) {
		var err error
		defer func() { 
			errC <- err
		}()

		if num == 0	{
			err = errors.New("num = 0")
			return
		} else if num > 0 {
			err = errors.New("num > 10")
			return
		}
		err = errors.New("ngga error sebenernya")
	}(i)

	var finalErr error 
	for i := 0; i < count; i++ {
		err := <- errC
		if err != nil {
			log.Println(err)
			finalErr = err
		}
	}
	return finalErr
}

func main() {
	channel(10)
	channel(0)
	channel(-10)
}