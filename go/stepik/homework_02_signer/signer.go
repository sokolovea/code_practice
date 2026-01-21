package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"sync"
)

// type endOfChan {}

var SingleHash job = func(in, out chan interface{}) {
	fmt.Printf("SingleHash!!!\n");
	// InputElemsCount := len(in)
	// fmt.Printf("InputElemsCount = %v!!!\n", InputElemsCount);
	for {
		data, ok := <-in
		if !ok {
			fmt.Println("Канал SingleHash закрыт")
			break
		}
		fmt.Printf("SingleHash get data!\n");
		dataString, ok := data.(string)
		if !ok {
			dataInt, ok := data.(int)
			if !ok {
				fmt.Printf("SingleHash not ok!\n");
				continue
			}
			dataString = strconv.Itoa(dataInt)
		}
		fmt.Printf("SingleHash: data = %v \n", dataString);
		out <- DataSignerCrc32(dataString) + "~" + DataSignerCrc32(DataSignerMd5(dataString))
	}
}

var MultiHash job = func(in, out chan interface{}) {
	fmt.Printf("MultiHash!!!\n");
	preResult := make([]string, 6)
	for {
		data, ok := <-in
		if !ok {
			fmt.Println("Канал MultiHash закрыт")
			break
		}
		fmt.Printf("MultiHash get data!\n")
		dataString, ok := data.(string)
		if !ok {
			fmt.Printf("MultiHash data not ok!\n");
			continue
		}
		for th := 0; th <= 5; th++ {
			preResult[th] = DataSignerCrc32(strconv.Itoa(th) + dataString)
		}
		preResult := strings.Join(preResult, "")
		fmt.Printf("MultiHash: data = %v \n", preResult);
		out <- preResult
	}
}

var CombineResults job = func(in, out chan interface{}) {
	fmt.Printf("CombineResults!!!\n");
	var allData []string = make([]string, 0, 200)
	for {
		data, ok := <-in
		if !ok {
			fmt.Println("Канал CombineResults закрыт")
			break
		}
		fmt.Printf("CombineResults get data!\n");
		dataString, ok := data.(string)
		if !ok {
			fmt.Printf("CombineResults data not ok!\n");
			continue
		}
		allData = append(allData, dataString)
	}
	sort.Slice(allData, func(i, j int) bool {
		return allData[i] < allData[j]
	})
	str := strings.Join(allData, "_") 
	fmt.Printf("CombineResults = %v\n", str);
	out <- str
}

func ExecutePipeline(jobs ...job) {
	fmt.Printf("ExecutePipeline!!!\n");
	channels := make([](chan interface{}), len(jobs) + 1)
	for i := 0; i < len(jobs) + 1; i++ {
		channels[i] = make(chan interface{}, 200)
	}

	fmt.Printf("%v\n", jobs)
	var wg sync.WaitGroup
	wg.Add(len(jobs))
	close(channels[0])
	for i, job := range(jobs) {
		go func(){
			job(channels[i], channels[i + 1])
			close(channels[i + 1])
			fmt.Println("job done!!!!!!!!!!!!!!!!!!!!!!!!!")
			wg.Done()
		}()
	}

	wg.Wait()
}