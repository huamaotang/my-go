package date

import (
	"fmt"
	"math"
	"testing"
	"time"
)

func TestSince(t *testing.T) {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	timeX, _ := time.ParseInLocation("20060102 15:04:05", "20210219 10:44:00", loc)
	t.Log(int(time.Since(timeX).Seconds()))
}

func TestName(t *testing.T) {
	fmt.Println(time.Now().Format("2006"))
	fmt.Println(time.Now().AddDate(-1, 1, -2).Format("20060102"))
}

func TestDays(t *testing.T) {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	timeX, _ := time.ParseInLocation("20060102 15:04:05", "20190325 18:00:00", loc)
	x := timeX.Unix()
	y := time.Now().Unix()
	fmt.Println(math.Ceil(float64(y-x) / 86400))
}

func TestWeekday(t *testing.T) {
	t.Log(WeekByDate(time.Now()))
}

//判断时间是当年的第几周
func WeekByDate(t time.Time) string {
	yearDay := t.YearDay()
	yearFirstDay := t.AddDate(0, 0, -yearDay+1)
	firstDayInWeek := int(yearFirstDay.Weekday())

	//今年第一周有几天
	firstWeekDays := 1
	if firstDayInWeek != 0 {
		firstWeekDays = 7 - firstDayInWeek + 1
	}
	var week int
	if yearDay <= firstWeekDays {
		week = 1
	} else {
		week = (yearDay-firstWeekDays)/7 + 2
	}
	return fmt.Sprintf("%d第%d周", t.Year(), week)
}
