package carbon_helper

import (
	"github.com/uniplaces/carbon"
)

var units = map[string]string{"week": "week", "weeks": "week", "day": "day", "days": "day", "d": "day", "month": "month", "months": "month", "m": "month", "quarter": "quarter", "q": "quarter", "quarters": "quarter", "year": "year", "years": "years", "y": "year"}

type CarbonHelper struct {
	Carbon *carbon.Carbon
}

func New(c *carbon.Carbon) *CarbonHelper {
	return &CarbonHelper{Carbon: c}
}

func (c *CarbonHelper) Add(unit string, amount int) *CarbonHelper {
	validUnit := units[unit]

	switch validUnit {
	case "day":
		c.Carbon = c.Carbon.AddDays(amount)
		break
	case "week":
		c.Carbon = c.Carbon.AddWeeks(amount)
		break
	case "month":
		c.Carbon = c.Carbon.AddMonths(amount)
		break
	case "year":
		c.Carbon = c.Carbon.AddYears(amount)
		break
	case "quartal":
		c.Carbon = c.Carbon.AddQuarters(amount)
		break
	}
	return c

}

func (c *CarbonHelper) Sub(unit string, amount int) *CarbonHelper {
	validUnit := units[unit]

	switch validUnit {
	case "day":
		c.Carbon = c.Carbon.SubDays(amount)
		break
	case "week":
		c.Carbon = c.Carbon.SubWeeks(amount)
		break
	case "month":
		c.Carbon = c.Carbon.SubMonths(amount)
		break
	case "year":
		c.Carbon = c.Carbon.SubYears(amount)
		break
	case "quartal":
		c.Carbon = c.Carbon.SubQuarters(amount)
		break
	}
	return c

}

func (c *CarbonHelper) StartOf(unit string) *CarbonHelper {
	validUnit := units[unit]

	switch validUnit {
	case "day":
		c.Carbon = c.Carbon.StartOfDay()
		break
	case "week":
		c.Carbon = c.Carbon.StartOfDay()
		break
	case "month":
		c.Carbon = c.Carbon.StartOfMonth()
		break
	case "year":
		c.Carbon = c.Carbon.StartOfYear()
		break
	case "quarter":
		c.Carbon = c.Carbon.StartOfQuarter()
		break
	}
	return c
}

func (c *CarbonHelper) EndOf(unit string) *CarbonHelper {
	validUnit := units[unit]
	switch validUnit {
	case "day":
		c.Carbon = c.Carbon.EndOfDay()
		break
	case "week":
		c.Carbon = c.Carbon.EndOfDay()
		break
	case "month":
		c.Carbon = c.Carbon.EndOfMonth()
		break
	case "year":
		c.Carbon = c.Carbon.EndOfYear()
		break
	case "quarter":
		c.Carbon = c.Carbon.EndOfQuarter()
		break
	}
	return c
}
