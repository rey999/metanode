package main

import (
	"fmt"
	"sync"
	"time"
)

type Schedule struct {
	name    *string
	start   time.Time
	end     time.Time
	use     int
	runFunc func(*Schedule, *sync.WaitGroup)
}

type ScheduleList struct {
	schedules []Schedule
}

func (s *Schedule) run(wg *sync.WaitGroup) {
	go s.runFunc(s, wg)
}

func (s *ScheduleList) add(runFunc func(*Schedule, *sync.WaitGroup)) {

	s.schedules = append(s.schedules, Schedule{runFunc: runFunc})
}
func main() {
	var wg sync.WaitGroup
	// var s string = "hello"
	var scheduleList ScheduleList
	scheduleList.add(func(s *Schedule, wg *sync.WaitGroup) {
		defer wg.Done()
		s.start = time.Now()
		time.Sleep(time.Second * 1)
		fmt.Println("hello")
		s.end = time.Now()
		s.use = int(s.end.Sub(s.start).Seconds())
	})

	scheduleList.add(func(s *Schedule, wg *sync.WaitGroup) {
		defer wg.Done()
		s.start = time.Now()
		time.Sleep(time.Second * 2)
		fmt.Println("word")
		s.end = time.Now()
		s.use = int(s.end.Sub(s.start).Seconds())
	})

	scheduleList.add(func(s *Schedule, wg *sync.WaitGroup) {
		defer wg.Done()
		s.start = time.Now()
		time.Sleep(time.Second * 3)
		fmt.Println("baby")
		s.end = time.Now()
		s.use = int(s.end.Sub(s.start).Seconds())
	})
	scheduleList.add(func(s *Schedule, wg *sync.WaitGroup) {
		defer wg.Done()
		s.start = time.Now()
		time.Sleep(time.Second * 4)
		fmt.Println("giaogiao")
		s.end = time.Now()
		s.use = int(s.end.Sub(s.start).Seconds())
	})

	for i := 0; i < len(scheduleList.schedules); i++ {
		wg.Add(1)
		scheduleList.schedules[i].run(&wg)
	}

	wg.Wait()

	for i := 0; i < len(scheduleList.schedules); i++ {
		fmt.Println(scheduleList.schedules[i].name, "开始时间：", scheduleList.schedules[i].start, "结束时间：", scheduleList.schedules[i].end, "运行：", scheduleList.schedules[i].use, "秒")
	}

	// time.Sleep(time.Second * 10)
}
