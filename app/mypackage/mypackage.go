package stuff

import (
	"errors"
	"strconv"
	"time"
)

var Name string = "Archish"

func IntArrToStrArr(intArr []int) []string {
	var strArr []string
	for _, v := range intArr {
		strArr = append(strArr, strconv.Itoa(v))
	}
	return strArr
}

type Date struct{
	day int
	month int
	year int
}

func (d *Date) SetDay(day int) error {
	if(day < 1) || (day > 31) {
		return errors.New("incorrect day value")
	}
	d.day = day
	return nil
}

func (d *Date) SetMonth(month int) error {
	if(month < 1) || (month > 12) {
		return errors.New("incorrect day value")
	}
	d.month = month
	return nil
}


func (d *Date) SetYear(year int) error {
	if(year < 1875) || (year > time.Now().Year()) {
		return errors.New("incorrect day value")
	}
	d.year = year
	return nil
}

func (d *Date) Day() int {
	return d.day
}

func (d *Date) Month() int {
	return d.month
}

func (d *Date) Year() int {
	return d.year
}