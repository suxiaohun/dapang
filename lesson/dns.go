package lesson

import (
	"context"
	"flag"
	"fmt"
	"net"
	"sync/atomic"
	"time"
)

var host string
var connections int
var duration int64
var limit int64
var timeoutCount int64

func DNS() {
	// os.Args = append(os.Args, "-host", "www.baidu.com", "-c", "200", "-d", "30", "-l", "5000")

	flag.StringVar(&host, "host", "www.baidu.com", "Resolve host")
	flag.IntVar(&connections, "c", 100, "Connections")
	flag.Int64Var(&duration, "d", 0, "Duration(s)")
	flag.Int64Var(&limit, "l", 0, "Limit(ms)")
	flag.Parse()

	var count int64 = 0
	var errCount int64 = 0
	pool := make(chan interface{}, connections)
	exit := make(chan bool)
	var (
		min int64 = 0
		max int64 = 0
		sum int64 = 0
	)

	go func() {
		time.Sleep(time.Second * time.Duration(duration))
		exit <- true
	}()
endD:
	for {
		select {
		case pool <- nil:
			go func() {
				defer func() {
					<-pool
				}()
				resolver := &net.Resolver{}
				now := time.Now()
				_, err := resolver.LookupIPAddr(context.Background(), host)
				use := time.Since(now).Nanoseconds() / int64(time.Millisecond)
				if min == 0 || use < min {
					min = use
				}
				if use > max {
					max = use
				}
				sum += use
				if limit > 0 && use >= limit {
					timeoutCount++
				}
				atomic.AddInt64(&count, 1)
				if err != nil {
					fmt.Println(err.Error())
					atomic.AddInt64(&errCount, 1)
				}
			}()
		case <-exit:
			break endD
		}
	}

	time.Sleep(3*time.Second)
	fmt.Printf("request count：%d\nerror count：%d\n", count, errCount)
	fmt.Printf("request time：min(%dms) max(%dms) avg(%dms) timeout(%dn)\n", min, max, sum/count, timeoutCount)
}