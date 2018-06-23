package shadowusers

import (
	"errors"
	"fmt"
	"strconv"
)

type Passwd struct {
	Login    string
	Password string
	UserId   int
	GroupId  int
	Comment  string
	Home     string
	Shell    string
}

func NewPasswdFromString(line string) (*Passwd, error) {
	self := new(Passwd)

	s := SplitLine(line)
	if len(s) != 7 {
		return nil, errors.New("invalid passwd line")
	}

	self.Login = s[0]
	self.Password = s[1]

	if t, err := strconv.Atoi(s[2]); err != nil {
		return nil, errors.New("invalid UserId")
	} else {
		self.UserId = t
	}

	if t, err := strconv.Atoi(s[3]); err != nil {
		return nil, errors.New("invalid GroupId")
	} else {
		self.GroupId = t
	}

	self.Comment = s[4]
	self.Home = s[5]
	self.Shell = s[6]

	return self, nil
}

func (self *Passwd) IsValid() bool {
	for _, value := range []string{
		self.Login,
		self.Password,
		self.Comment,
		self.Home,
		self.Shell,
	} {
		if !StringValueValid(value) {
			return false
		}
	}
	return true
}

func (self *Passwd) Render() (string, error) {

	if !self.IsValid() {
		return "", errors.New("Passwd structure have invalid field values")
	}

	ret := fmt.Sprintf(
		"%s:%s:%d:%d:%s:%s:%s",
		self.Login,
		self.Password,
		self.UserId,
		self.GroupId,
		self.Comment,
		self.Home,
		self.Shell,
	)

	return ret, nil
}
