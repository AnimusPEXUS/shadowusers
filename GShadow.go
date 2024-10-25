package shadowusers

import (
	"errors"
	"fmt"
)

type GShadow struct {
	Name           string
	Password       string
	Administrators []string
	Members        []string
}

func NewGShadowFromString(line string) (*GShadow, error) {
	self := new(GShadow)

	s := SplitLine(line)
	if len(s) != 4 {
		return nil, errors.New("invalid GShadow line")
	}

	self.Name = s[0]
	self.Password = s[1]

	self.Administrators = ParseUserList(s[2])
	self.Members = ParseUserList(s[3])

	return self, nil
}

func (self *GShadow) IsValid() bool {
	for _, value := range []string{
		self.Name,
		self.Password,
	} {
		if !StringValueValid(value) {
			return false
		}
	}
	return true
}

func (self *GShadow) Render() (string, error) {
	if !self.IsValid() {
		return "", errors.New("GShadow structure have invalid field values")
	}

	adms, err := RenderUserList(self.Administrators)
	if err != nil {
		return "", err
	}

	membs, err := RenderUserList(self.Members)
	if err != nil {
		return "", err
	}

	ret := fmt.Sprintf(
		"%s:%s:%s:%s",
		self.Name,
		self.Password,
		adms,
		membs,
	)
	return ret, nil
}
