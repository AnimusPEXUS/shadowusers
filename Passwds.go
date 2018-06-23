package shadowusers

import (
	"errors"
	"fmt"
	"strings"
)

type Passwds struct {
	Passwds []*Passwd
}

func NewPasswdsFromString(txt string) (*Passwds, error) {
	self := new(Passwds)
	ss := strings.Split(txt, "\n")
	for k, v := range ss {
		if IsEmptyOrSpacesLine(v) {
			continue
		}
		parsed, err := NewPasswdFromString(v)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("line %d parse error: %s", k, err.Error()))
		}
		self.Passwds = append(self.Passwds, parsed)
	}
	return self, nil
}

func (self *Passwds) GetByLogin(login string) (*Passwd, error) {
	for _, i := range self.Passwds {
		if i.Login == login {
			return i, nil
		}
	}
	return nil, errors.New("not found")
}

func (self *Passwds) GetByUid(uid int) (*Passwd, error) {
	for _, i := range self.Passwds {
		if i.UserId == uid {
			return i, nil
		}
	}
	return nil, errors.New("not found")
}

func (self *Passwds) GetByGid(gid int) ([]*Passwd, error) {
	ret := make([]*Passwd, 0)
	for _, i := range self.Passwds {
		if i.GroupId == gid {
			ret = append(ret, i)
		}
	}
	return ret, nil
}

func (self *Passwds) Render() (string, error) {
	ret := ""
	for k, v := range self.Passwds {
		t, err := v.Render()
		if err != nil {
			return "", errors.New(
				fmt.Sprintf(
					"error rendering Passwds line %d: %s",
					k,
					err.Error(),
				),
			)
		}

		ret += t + "\n"
	}
	return ret, nil
}
