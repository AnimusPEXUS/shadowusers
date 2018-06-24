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

func (self *Passwds) SortByUid() {
	lg := len(self.Passwds)
	for i := 0; i != lg-1; i++ {
		for j := i + 1; j != lg; j++ {
			if self.Passwds[i].UserId > self.Passwds[j].UserId {
				z := self.Passwds[i]
				self.Passwds[i] = self.Passwds[j]
				self.Passwds[j] = z
			}
		}
	}
}

func (self *Passwds) ShalowCopy() *Passwds {
	ret := &Passwds{}
	for _, i := range self.Passwds {
		ret.Passwds = append(ret.Passwds, i)
	}
	return ret
}

func (self *Passwds) Render() (string, error) {

	c := self.ShalowCopy()
	c.SortByUid()

	ret := ""
	for k, v := range c.Passwds {
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
