package main

import (
	"fmt"
	"net/http"
	"log"
	"sync"
    "io/ioutil"
)

func main() {
	out := make(chan string)
	var wg sync.WaitGroup

	wg.Add(1)   
	go getOneData(out,&wg)
	wg.Add(1)   
	go getSecondData(out,&wg)
	
	go func() {
		wg.Wait()	
		close(out)   
	}() 

	for elem:= range out {
        fmt.Println(elem)
    }

}

func getOneData(out chan string,wg *sync.WaitGroup)  {
	defer (*wg).Done()
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
	   log.Fatalln(err,resp)
	}	
	body, err := ioutil.ReadAll(resp.Body)

    if err != nil {
        log.Fatal(err)
    }
	out <- string(body)
}

func getSecondData(out chan string,wg *sync.WaitGroup)  {
	defer (*wg).Done()
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/2")
	if err != nil {
	   log.Fatalln(err,resp)
	}	
	body, err := ioutil.ReadAll(resp.Body)

    if err != nil {
        log.Fatal(err)
    }
	out <- string(body)
}