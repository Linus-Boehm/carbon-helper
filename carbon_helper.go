package carbon_helper

import (
	"errors"
	"github.com/uniplaces/carbon"
)

var units = map[string]string{"week": "week", "weeks": "week", "day": "day", "days": "day", "d": "day", "month": "month", "months": "month", "m": "month", "quarter": "quarter", "q": "quarter", "quarters": "quarter", "year": "year", "years": "years", "y": "year"}

type CarbonHelper struct {
	Carbon *carbon.Carbon
}
func New(c *carbon.Carbon) *CarbonHelper{
	return &CarbonHelper{Carbon:c}
}

func (c *CarbonHelper) Add(unit string, amount int) (*carbon.Carbon, error) {
	validUnit, ok := units[unit]
	if !ok {
		return c.Carbon, errors.New("Unsopported unit: " + unit)
	}
	switch validUnit {
	case "day":
		return c.Carbon.AddDays(amount), nil
	case "week":
		return c.Carbon.AddWeeks(amount), nil
	case "month":
		return c.Carbon.AddMonths(amount), nil
	case "year":
		return c.Carbon.AddYears(amount), nil
	case "quartal":
		return c.Carbon.AddQuarters(amount), nil
	}
	return c.Carbon, errors.New("Add not implemented for unit: " + unit)
}

func (c *CarbonHelper) Sub(unit string, amount int) (*carbon.Carbon, error) {
	validUnit, ok := units[unit]
	if !ok {
		return c.Carbon, errors.New("Unsopported unit: " + unit)
	}
	switch validUnit {
	case "day":
		return c.Carbon.SubDays(amount), nil
	case "week":
		return c.Carbon.SubWeeks(amount), nil
	case "month":
		return c.Carbon.SubMonths(amount), nil
	case "year":
		return c.Carbon.SubYears(amount), nil
	case "quartal":
		return c.Carbon.SubQuarters(amount), nil
	}
	return c.Carbon, errors.New("Sub not implemented for unit: " + unit)
}

func (c *CarbonHelper) StartOf(unit string) (*carbon.Carbon, error) {
	validUnit, ok := units[unit]
	if !ok {
		return c.Carbon, errors.New("Unsopported unit: " + unit)
	}
	switch validUnit {
	case "day":
		return c.Carbon.StartOfDay(), nil
	case "week":
		return c.Carbon.StartOfDay(), nil
	case "month":
		return c.Carbon.StartOfMonth(), nil
	case "year":
		return c.Carbon.StartOfYear(), nil
	case "quarter":
		return c.Carbon.StartOfQuarter(), nil
	}
	return c.Carbon, errors.New("StartOf not implemented for unit: " + unit)
}

func (c *CarbonHelper) EndOf(unit string) (*carbon.Carbon, error) {
	validUnit, ok := units[unit]
	if !ok {
		return c.Carbon, errors.New("Unsopported unit: " + unit)
	}
	switch validUnit {
	case "day":
		return c.Carbon.EndOfDay(), nil
	case "week":
		return c.Carbon.EndOfDay(), nil
	case "month":
		return c.Carbon.EndOfMonth(), nil
	case "year":
		return c.Carbon.EndOfYear(), nil
	case "quarter":
		return c.Carbon.EndOfQuarter(), nil
	}
	return c.Carbon, errors.New("EndOf not implemented for unit: " + unit)
}
