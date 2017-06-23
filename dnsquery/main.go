package main

import (
	"flag"
	"log"
	"net"
	"sync"
	"time"
)

var (
	host     *string = flag.String("host", "", "domain name")
	batch    *int    = flag.Int("batch", 30, "bach numbers per epoch")
	epoch    *int    = flag.Int("epoch", 10, "epoch numbers per second")
	duration *int    = flag.Int("duration", 1, "duration seconds")
)

func main() {
	var wg sync.WaitGroup
	var err error
	var addrs []string

	flag.Parse()
	if len(*host) == 0 {
		flag.PrintDefaults()
		return
	}

	max := 0
	sleep := 1e6 / *epoch
	log.Println("test begin...")
	for i := 0; i < *duration; i++ {
		for j := 0; j < *epoch; j++ {
			for k := 0; k < *batch; k++ {
				wg.Add(1)
				go func() {
					st := time.Now()
					addrs, err = net.LookupHost(*host)
					cost := int(time.Now().Sub(st) / 1e6)
					if cost > max {
						max = cost
					}
					wg.Done()

					if err != nil {
						log.Fatal(err)
					}
				}()
			}
			time.Sleep(time.Microsecond * time.Duration(sleep))
		}
	}

	wg.Wait()
	log.Printf("The max cost time = %d", max)
	log.Println(addrs)
	log.Println("test end...")
}
