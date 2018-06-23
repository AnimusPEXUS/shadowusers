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

func (self *Shadows) Render() (string, error) {
	ret := ""
	for k, v := range self.Shadows {
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
