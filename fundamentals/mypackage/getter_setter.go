package mypackage

import "errors"

type Date struct {
	day, month, year int
}

func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("incorrect day Value")
	}
	d.day = day
	return nil
}

func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 31 {
		return errors.New("incorrect month Value")
	}
	d.month = month
	return nil
}

func (d *Date) SetYear(year int) error {
	if year < 1 || year > 31 {
		return errors.New("incorrect year Value")
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
