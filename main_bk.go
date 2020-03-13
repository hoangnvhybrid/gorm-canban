package main

import (
"fmt"
"time"
)

func main() {
	// test cho offshore-3
	//expDateFirstFrom := time.Now().AddDate(0, 0, -1*int(60))
	//fmt.Println(expDateFirstFrom)
	//
	////init the loc
	//loc, _ := time.LoadLocation("Asia/Tokyo")
	////set timezone,
	//now := time.Now().In(loc)
	//fmt.Println("now: ", now.AddDate(0, 0, -60))
	//fmt.Println("now: ", now.AddDate(0, 0, -45))

	//loc, _ := time.LoadLocation("Asia/Tokyo")
	//now := time.Now().In(loc)
	//
	//loc, _ := time.LoadLocation("Asia/Tokyo")
	//now := time.Now();
	//abc := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	//fmt.Println("abc", abc)
	//expDateFirstFrom := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc).AddDate(0, 0, -1*int(60))
	//fmt.Println(expDateFirstFrom)
	//expDateFirstTo := expDateFirstFrom.AddDate(0, 0, 1)
	//expDateSecondFrom := now.Truncate(24*time.Hour).AddDate(0, 0, -1*int(expDaySecond))
	//expDateSecondTo := expDateSecondFrom.AddDate(0, 0, 1)

	//fmt.Println("hello");
	expDateFirstFrom := time.Now().AddDate(0, 0, -1*int(180))
	fmt.Println(expDateFirstFrom)
	fmt.Println(expDateFirstFrom.AddDate(0,0, 150))
	fmt.Println(expDateFirstFrom.AddDate(0,0, 177))
	fmt.Println(expDateFirstFrom.AddDate(0,0, 179))
	fmt.Println(expDateFirstFrom.AddDate(0,0, 180))
	//ngay mua + 180= ngay hom nay+30 => ngay mua = ngay hom nay + 30 -180
	//ngay hom nay -30 + 180 = ngay het han
	//expDateFirstTo := expDateFirstFrom.AddDate(0, 0, 1)
	//expDateSecondFrom := time.Now().AddDate(0, 0, -1*int(expDaySecond))
	//expDateSecondTo := expDateSecondFrom.AddDate(0, 0, 1)
	//2019-09-06 10:56:22.288972739 +0700 +07 + 180 = ngay hom nay khong

	//layout := "2019-09-06"
	//str := "2019-09-06"
	//t, err := time.Parse(layout, str)
	//
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(t)
	//t, _ := time.Parse("2018-09-05T10:56:22288972739Z", "2019-09-05 10:56:22.288972739")
	//loc, _ := time.LoadLocation("Asia/Tokyo")
	//t1 := time.Now().In(loc).AddDate(0, 0, -1*int(150))
	////fmt.Println(t)
	//fmt.Println(t1)
}

