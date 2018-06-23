package shadowusers

import (
	"errors"
	"fmt"
	"strings"

	"github.com/AnimusPEXUS/utils/set"
)

type GShadows struct {
	GShadows []*GShadow
}

func NewGShadowsFromString(txt string) (*GShadows, error) {
	self := new(GShadows)
	ss := strings.Split(txt, "\n")
	for k, v := range ss {
		if IsEmptyOrSpacesLine(v) {
			continue
		}
		parsed, err := NewGShadowFromString(v)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("line %d parse error: %s", k, err.Error()))
		}
		self.GShadows = append(self.GShadows, parsed)
	}
	return self, nil
}

func (self *GShadows) GetByName(name string) (*GShadow, error) {
	for _, i := range self.GShadows {
		if i.Name == name {
			return i, nil
		}
	}
	return nil, errors.New("not found")
}

func (self *GShadows) GetByAdministrator(user string) ([]*GShadow, error) {
	ret := make([]*GShadow, 0)
	for _, i := range self.GShadows {
		s := set.NewSetString()
		s.AddStrings(i.Administrators...)

		if s.Have(user) {
			ret = append(ret, i)
		}
	}
	return ret, nil
}

func (self *GShadows) GetByMember(user string) ([]*GShadow, error) {
	ret := make([]*GShadow, 0)
	for _, i := range self.GShadows {
		s := set.NewSetString()
		s.AddStrings(i.Members...)

		if s.Have(user) {
			ret = append(ret, i)
		}
	}
	return ret, nil
}

func (self *GShadows) Render() (string, error) {
	ret := ""
	for k, v := range self.GShadows {
		t, err := v.Render()
		if err != nil {
			return "", errors.New(
				fmt.Sprintf(
					"error rendering GShadows line %d: %s",
					k,
					err.Error(),
				),
			)
		}

		ret += t + "\n"
	}
	return ret, nil
}
