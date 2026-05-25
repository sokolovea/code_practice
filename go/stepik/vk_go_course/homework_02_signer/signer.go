package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"sync"
)

var mutexMd5 sync.Mutex
var crc32ResultsMap map[string]string = make(map[string]string)

var SingleHash job = func(in, out chan interface{}) {
	fmt.Printf("SingleHash!!!\n")
	var wgCrc32 sync.WaitGroup
	var wgSingleHash sync.WaitGroup
	for {
		data, ok := <-in
		if !ok {
			fmt.Println("Канал SingleHash закрыт")
			break
		}
		wgSingleHash.Add(1)
		go func() {
			fmt.Printf("SingleHash get data!\n")
			dataString, ok := data.(string)
			if !ok {
				dataInt, ok := data.(int)
				if !ok {
					fmt.Printf("SingleHash not ok!\n")
					return
				}
				dataString = strconv.Itoa(dataInt)
			}
			fmt.Printf("SingleHash: data = %v \n", dataString)
			wgCrc32.Add(2)
			var firstSum string
			var secondSum string
			go func() {
				firstSum = DataSignerCrc32(dataString)
				wgCrc32.Done()
			}()
			go func() {
				mutexMd5.Lock()
				preResult := DataSignerMd5(dataString)
				mutexMd5.Unlock()
				secondSum = DataSignerCrc32(preResult)
				wgCrc32.Done()
			}()
			wgCrc32.Wait()
			out <- firstSum + "~" + secondSum
			wgSingleHash.Done()
		}()
	}
	wgSingleHash.Wait()
}

var MultiHash job = func(in, out chan interface{}) {
	fmt.Printf("MultiHash!!!\n")
	var wgMultiHash sync.WaitGroup
	for {
		data, ok := <-in
		if !ok {
			fmt.Println("Канал MultiHash закрыт")
			break
		}
		fmt.Printf("MultiHash get data!\n")
		dataString, ok := data.(string)
		if !ok {
			fmt.Printf("MultiHash data not ok!\n")
			continue
		}
		wgMultiHash.Add(1)
		go func(dataStringCopy string) {
			preResult := make([]string, 6)
			var wg sync.WaitGroup
			wg.Add(6)
			for th := 0; th <= 5; th++ {
				go func() {
					preResult[th] = DataSignerCrc32(strconv.Itoa(th) + dataStringCopy)
					wg.Done()
				}()
			}
			wg.Wait()
			result := strings.Join(preResult, "")
			fmt.Printf("MultiHash: data = %v \n", result)
			out <- result
			wgMultiHash.Done()
		}(dataString)
	}
	wgMultiHash.Wait()
}

var CombineResults job = func(in, out chan interface{}) {
	fmt.Printf("CombineResults!!!\n")
	var allData []string = make([]string, 0, 200)
	for {
		data, ok := <-in
		if !ok {
			fmt.Println("Канал CombineResults закрыт")
			break
		}
		fmt.Printf("CombineResults get data!\n")
		dataString, ok := data.(string)
		if !ok {
			fmt.Printf("CombineResults data not ok!\n")
			continue
		}
		allData = append(allData, dataString)
	}
	sort.Slice(allData, func(i, j int) bool {
		return allData[i] < allData[j]
	})
	str := strings.Join(allData, "_")
	fmt.Printf("CombineResults = %v\n", str)
	out <- str
}

func ExecutePipeline(jobs ...job) {
	fmt.Printf("ExecutePipeline!!!\n")
	channels := make([](chan interface{}), len(jobs)+1)
	for i := 0; i < len(jobs)+1; i++ {
		channels[i] = make(chan interface{}, 200)
	}

	fmt.Printf("%v\n", jobs)
	var wg sync.WaitGroup
	wg.Add(len(jobs))
	close(channels[0])
	for i, job := range jobs {
		go func() {
			job(channels[i], channels[i+1])
			close(channels[i+1])
			fmt.Println("job done!!!!!!!!!!!!!!!!!!!!!!!!!")
			wg.Done()
		}()
	}

	wg.Wait()
}
