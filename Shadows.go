package shadowusers

import (
	"errors"
	"fmt"
	"strings"
)

type Shadows struct {
	Shadows []*Shadow
}

func NewShadowsFromString(txt string) (*Shadows, error) {
	self := new(Shadows)
	ss := strings.Split(txt, "\n")
	for k, v := range ss {
		if IsEmptyOrSpacesLine(v) {
			continue
		}
		parsed, err := NewShadowFromString(v)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("line %d parse error: %s", k, err.Error()))
		}
		self.Shadows = append(self.Shadows, parsed)
	}
	return self, nil
}

func (self *Shadows) GetByLogin(login string) (*Shadow, error) {
	for _, i := range self.Shadows {
		if i.Login == login {
			return i, nil
		}
	}
	return nil, errors.New("not found")
}

func (self *Shadows) SortByLogin() {
	lg := len(self.Shadows)
	if lg < 2 {
		return
	}
	for i := 0; i != lg-1; i++ {
		for j := i + 1; j != lg; j++ {
			if self.Shadows[i].Login > self.Shadows[j].Login {
				z := self.Shadows[i]
				self.Shadows[i] = self.Shadows[j]
				self.Shadows[j] = z
			}
		}
	}
}

func (self *Shadows) ShalowCopy() *Shadows {
	ret := &Shadows{}
	for _, i := range self.Shadows {
		ret.Shadows = append(ret.Shadows, i)
	}
	return ret
}

func (self *Shadows) Render() (string, error) {

	c := self.ShalowCopy()
	c.SortByLogin()

	ret := ""
	for k, v := range c.Shadows {
		t, err := v.Render()
		if err != nil {
			return "", errors.New(
				fmt.Sprintf(
					"error rendering Shadows line %d: %s",
					k,
					err.Error(),
				),
			)
		}

		ret += t + "\n"
	}
	return ret, nil
}
