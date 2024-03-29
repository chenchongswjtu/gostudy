package main

import (
	"fmt"
	"log"
	"time"
)

// 获取当前时间
func getNow() {
	// 获取当前时间，返回time.Time对象
	fmt.Println(time.Now())
	// output: 2016-07-27 08:57:46.53277327 +0800 CST
	// 其中CST可视为美国，澳大利亚，古巴或中国的标准时间
	// +0800表示比UTC时间快8个小时

	// 获取当前时间戳
	fmt.Println(time.Now().Unix())
	// 精确到纳秒，通过纳秒就可以计算出毫秒和微妙
	fmt.Println(time.Now().UnixNano())
	// output:
	//    1469581066
	//    1469581438172080471
}

// 格式化时间显示
func formatUnixTime() {
	// 获取当前时间，进行格式化
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	// output: 2016-07-27 08:57:46

	// 指定的时间进行格式化
	fmt.Println(time.Unix(1469579899, 0).Format("2006-01-02 15:04:05"))
	// output: 2016-07-27 08:38:19
}

// 获取指定时间戳的年份
func getYear() {
	// 获取指定时间戳的年月日，小时分钟秒
	t := time.Unix(1469579899, 0)
	fmt.Printf("%d-%d-%d %d:%d:%d\n", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	// output: 2016-7-27 8:38:19
}

// 时间字符串转换时间戳
// 将2016-07-27 08:46:15这样的时间字符串转换时间戳
func strToUnix() {
	// 先用time.Parse对时间字符串进行分析，如果正确会得到一个time.Time对象
	// 后面就可以用time.Time对象的函数Unix进行获取
	t2, err := time.Parse("2006-01-02 15:04:05", "2016-07-27 08:46:15")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(t2)
	fmt.Println(t2.Unix())
	// output:
	//     2016-07-27 08:46:15 +0000 UTC
	//     1469609175
}

// 根据时间戳获取当日开始的时间戳
// 这个在统计功能中会常常用到
// 方法就是通过时间戳取到2016-01-01 00:00:00这样的时间格式
// 然后再转成时间戳就OK了
// 获取月开始时间和年开始时间类似
func getDayStartUnix() {
	t := time.Unix(1469581066, 0).Format("2006-01-02")
	sts, err := time.Parse("2006-01-02", t)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(sts.Unix())
	// output: 1469577600
}

// 休眠
func sleep() {
	// 休眠1秒
	// time.Millisecond    表示1毫秒
	// time.Microsecond    表示1微妙
	// time.Nanosecond    表示1纳秒
	time.Sleep(1 * time.Second)
	// 休眠100毫秒
	time.Sleep(100 * time.Millisecond)
}

func main() {
	getNow()
	formatUnixTime()
	getYear()
	strToUnix()
	getDayStartUnix()
	sleep()

	//在windows系统上，没有安装go语言环境的情况下，time.LoadLocation会加载失败。
	var sh, _ = time.LoadLocation("Asia/Shanghai")
	log.Println(time.Now().In(sh).Format("2006-01-02 15:04:05"))

	//time.FixedZone各个系统都能很好的设置时区
	log.Println(time.Now().In(time.FixedZone("CST", 8*3600)).Format("2006-01-02 15:04:05"))

	//h -- > 时
	//m -- > 分
	//s -- > 秒
	//ms -- > 毫秒
	//us -- > 纳秒
	//µs -- > 纳秒
	//ns -- > 微秒

	//10分钟前的时间
	m, _ := time.ParseDuration("-10m")
	_ = time.Now().Add(m)

	//10分钟后的时间
	m, _ = time.ParseDuration("10m")
	_ = time.Now().Add(m)

	time.Now().Add(time.Second)
	t1 := time.Now()
	t2 := time.Now()
	t2.Sub(t1)

	t := time.Now()
	time.Since(t) // 与 time.Now().Sub(t) 相同

	time1 := "2019-05-20 18:30:50"
	time2 := "2019-05-20 17:30:50"
	t1, _ = time.Parse("2006-01-02 15:04:05", time1)
	t2, _ = time.Parse("2006-01-02 15:04:05", time2)

	//t1的时间是否早于t2
	t1.Before(t2)

	//t1的时间是否晚于t2
	t1.After(t2)

	//t1的时间是否与t2相等
	t1.Equal(t2)

	now := time.Now()
	// backdate 30 days
	notBefore := now.Add(-30 * 24 * time.Hour)
	// set expiry to around 100 years
	notAfter := time.Date(now.Year()+100, now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, time.Local)

	fmt.Println(notBefore.UTC())
	fmt.Println(notAfter.UTC())
}
