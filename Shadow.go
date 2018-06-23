package shadowusers

import (
	"errors"
	"fmt"
)

type Shadow struct {
	Login                 string
	Password              string
	LastChangeDays        int
	MinAgeDays            int
	MaxAgeDays            int
	WarningPeriodDays     int
	InactivityPeriodDays  int
	AccountExpirationDays int
	// one more - reserved
}

func NewShadowFromString(line string) (*Shadow, error) {
	self := new(Shadow)

	s := SplitLine(line)
	if len(s) != 9 {
		return nil, errors.New("invalid Shadow line")
	}

	self.Login = s[0]
	self.Password = s[1]

	t, err := AtoiEmptyIsMinus1(s[2], "LastChangeDays")
	if err != nil {
		return nil, err
	}
	self.LastChangeDays = t

	t, err = AtoiEmptyIsMinus1(s[3], "MinAgeDays")
	if err != nil {
		return nil, err
	}
	self.MinAgeDays = t

	t, err = AtoiEmptyIsMinus1(s[4], "MaxAgeDays")
	if err != nil {
		return nil, err
	}
	self.MaxAgeDays = t

	t, err = AtoiEmptyIsMinus1(s[5], "WarningPeriodDays")
	if err != nil {
		return nil, err
	}
	self.WarningPeriodDays = t

	t, err = AtoiEmptyIsMinus1(s[6], "InactivityPeriodDays")
	if err != nil {
		return nil, err
	}
	self.InactivityPeriodDays = t

	t, err = AtoiEmptyIsMinus1(s[7], "AccountExpirationDays")
	if err != nil {
		return nil, err
	}
	self.AccountExpirationDays = t

	return self, nil
}

func (self *Shadow) IsValid() bool {
	for _, value := range []string{
		self.Login,
		self.Password,
	} {
		if !StringValueValid(value) {
			return false
		}
	}
	return true
}

func (self *Shadow) Render() (string, error) {
	if !self.IsValid() {
		return "", errors.New("Shadow structure have invalid field values")
	}
	ret := fmt.Sprintf(
		"%s:%s:%s:%s:%s:%s:%s:%s:", // NOTE: reserved field at the end
		self.Login,
		self.Password,
		ItoaMinusIsEmpty(self.LastChangeDays),
		ItoaMinusIsEmpty(self.MinAgeDays),
		ItoaMinusIsEmpty(self.MaxAgeDays),
		ItoaMinusIsEmpty(self.WarningPeriodDays),
		ItoaMinusIsEmpty(self.InactivityPeriodDays),
		ItoaMinusIsEmpty(self.AccountExpirationDays),
	)
	return ret, nil
}
