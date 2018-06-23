package shadowusers

import (
	"errors"
	"fmt"
	"strings"

	"github.com/AnimusPEXUS/utils/set"
)

type Groups struct {
	Groups []*Group
}

func NewGroupsFromString(txt string) (*Groups, error) {
	self := new(Groups)
	ss := strings.Split(txt, "\n")
	for k, v := range ss {
		if IsEmptyOrSpacesLine(v) {
			continue
		}
		parsed, err := NewGroupFromString(v)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("line %d parse error: %s", k, err.Error()))
		}
		self.Groups = append(self.Groups, parsed)
	}
	return self, nil
}

func (self *Groups) GetByName(name string) (*Group, error) {
	for _, i := range self.Groups {
		if i.Name == name {
			return i, nil
		}
	}
	return nil, errors.New("not found")
}

func (self *Groups) GetByGID(id int) (*Group, error) {
	for _, i := range self.Groups {
		if i.GID == id {
			return i, nil
		}
	}
	return nil, errors.New("not found")
}

func (self *Groups) GetByUser(user string) ([]*Group, error) {
	ret := make([]*Group, 0)
	for _, i := range self.Groups {
		s := set.NewSetString()
		s.AddStrings(i.UserList...)

		if s.Have(user) {
			ret = append(ret, i)
		}
	}
	return ret, nil
}

func (self *Groups) Render() (string, error) {
	ret := ""
	for k, v := range self.Groups {
		t, err := v.Render()
		if err != nil {
			return "", errors.New(
				fmt.Sprintf(
					"error rendering Groups line %d: %s",
					k,
					err.Error(),
				),
			)
		}

		ret += t + "\n"
	}
	return ret, nil
}
