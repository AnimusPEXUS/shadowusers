package shadowusers

import (
	"errors"
	"fmt"
	"strconv"
)

type Group struct {
	Name     string
	Password string
	GID      int
	UserList []string
}

func NewGroupFromString(line string) (*Group, error) {
	self := new(Group)

	s := SplitLine(line)
	if len(s) != 4 {
		return nil, errors.New("invalid Group line")
	}

	self.Name = s[0]
	self.Password = s[1]

	if t, err := strconv.Atoi(s[2]); err != nil {
		return nil, errors.New("invalid GID")
	} else {
		self.GID = t
	}

	self.UserList = ParseUserList(s[3])

	return self, nil
}

func (self *Group) IsValid() bool {
	for _, value := range []string{
		self.Name,
		self.Password,
	} {
		if !StringValueValid(value) {
			return false
		}
	}

	for _, value := range self.UserList {
		if !StringValueValid(value) {
			return false
		}
	}
	return true
}

func (self *Group) Render() (string, error) {

	if !self.IsValid() {
		return "", errors.New("Group structure have invalid field values")
	}

	user_list, err := RenderUserList(self.UserList)
	if err != nil {
		return "", err
	}

	ret := fmt.Sprintf(
		"%s:%s:%d:%s",
		self.Name,
		self.Password,
		self.GID,
		user_list,
	)

	return ret, nil
}
