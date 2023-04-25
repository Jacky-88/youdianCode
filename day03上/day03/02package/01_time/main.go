package main

import (
	"fmt"
	"time"
)

func main() {
	//1、时间对象
	now := time.Now()
	fmt.Println(now)
	//2、格式化时间,格式必须是2006-01-02 15:04:05
	strTime := now.Format("2006-01-02 15:04:05")
	fmt.Println(strTime)
	//3、时间戳: 从1970年1月1日开始到现在的时间，秒，毫秒
	fmt.Println(now.Unix()) //毫秒
	//4、格式化时间转时间对象
	loc, _ := time.LoadLocation("Asia/Shanghai")
	timeObj, _ := time.ParseInLocation("2006-01-02 15:04:05", strTime, loc)
	fmt.Println(timeObj)

	//获取年、月、日
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Println(year, month, day, hour, minute, second)
	fmt.Printf("%d-%02d-%02d %02d-%02d-%02d \n", year, month, day, hour, minute, second)

	//时间戳
	timestamp1 := now.Unix() //当前时间的秒时间戳
	timestamp2 := now.UnixNano() //当前时间的纳秒时间戳
	fmt.Println(timestamp1, timestamp2)
	//秒时间戳转时间类型
	timeObj1 := time.Unix(timestamp1, 0)
	fmt.Println(timeObj1)

	//时间间隔
	//fmt.Println("sleep开始", time.Now())
	//time.Sleep(1 * time.Second)
	//fmt.Println("sleep结束", time.Now())

	//拿昨天的数据 2022-10-21
	//startTime 2022-10-20 endTime 2022-10-22
	//时间计算
	//add
	now1 := time.Now()
	fmt.Println(now1)
	duration, _ := time.ParseDuration("1m")  // 60 * time.Second
	m1 := now1.Add(duration)
	fmt.Println(m1)

	//sub
	fmt.Println(now1.Sub(m1))
}

