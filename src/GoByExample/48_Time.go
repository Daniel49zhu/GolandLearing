package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	now := time.Now()
	p(now)

	then := time.Date(2018, 12, 12, 11, 23, 57, 651387234, time.UTC)
	p(then)

	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	p(then.Weekday())

	//判断两个日期的先后顺序
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	//获取两个时间之间的差
	diff := now.Sub(then)
	p(diff)
	//按照不同的时间度量来衡量时间差
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	p(then.Add(diff))
	p(then.Add(-diff))
}
