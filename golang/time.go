package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

func main() {
	// TimeFormat := "2006-01-02 15:04:05"
	// DateFormat := "2006-01-02"
	// now := time.Now().Format(DateFormat)
	//now, _ := time.Parse(format, time.Now().Format(format))
	// a, _ := time.Parse(TimeFormat, "2014-03-13 00:00:00")
	// b, _ := time.Parse(TimeFormat, "2015-03-13 00:00:00")

	// fmt.Println(now.Format(format), a.Format(format), b.Format(format))
	// fmt.Println(now.After(a))
	// fmt.Println(now.Before(a))
	// fmt.Println(now.After(b))
	// fmt.Println(now.Before(b))
	// fmt.Println(a, b)
	// fmt.Println(a.Before(b))
	// fmt.Println(now.Format(format), a.Format(format), b.Format(format))
	// fmt.Println(now.Unix(), a.Unix(), b.Unix())

	//获取时间戳
	// timestamp := time.Now().Unix()
	// fmt.Println(timestamp)

	//格式化为字符串,tm为Time类型
	// tm := time.Unix(timestamp, 0)
	// fmt.Println(tm.Format("2006-01-02 03:04:05 PM"))
	// fmt.Println(tm.Format("02/01/2006 15:04:05 PM"))

	//从字符串转为时间戳，第一个参数是格式，第二个是要转换的时间字符串
	// tm2, _ := time.Parse("01/02/2006", "02/08/2015")
	// fmt.Println(tm2.Unix())

	const longForm = "Jan 2, 2006 at 3:04pm (MST)"
	t, _ := time.Parse(longForm, "Jun 21, 2017 at 0:00am (PST)")
	fmt.Println("1=====", t)

	const shortForm = "2006-Jan-02"
	t, _ = time.Parse(shortForm, "2017-Jun-21")
	fmt.Println("2=====", t)

	t, _ = time.Parse("01/02/2006 15:04:05", "06/21/2017 13:21:00")
	fmt.Println("3=====", t)
	fmt.Println("3=====", t.Unix())

	i, err := strconv.ParseInt("1498003200", 10, 64)
	if err != nil {
		panic(err)
	}
	tm := time.Unix(i, 0)
	fmt.Println("4=====", tm)

	var timestamp int64 = 1525334350
	tm2 := time.Unix(timestamp, 0)
	fmt.Println("5=====", tm2.Format("2006-01-02 15:04:05 PM"))
	fmt.Println("5=====", tm2.Format("02/01/2006 15:04:05 PM"))

	loc, _ := time.LoadLocation("Local")
	now1, _ := time.Parse("2006-01-02 15:04", time.Now().Format("2006-01-02 15:04"))
	now2, _ := time.ParseInLocation("2006-01-02 15:04", time.Now().Format("2006-01-02 15:04"), loc)
	fmt.Println(now1, now2)
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(strconv.FormatInt(now2.Unix(), 10)))
	cipherStr := md5Ctx.Sum(nil)
	fmt.Println(hex.EncodeToString(cipherStr))

}
