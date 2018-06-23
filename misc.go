package shadowusers

import (
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/AnimusPEXUS/utils/set"
)

var EmptyOrSpacesLineRe = regexp.MustCompile(`^\s*$`)

func IsEmptyOrSpacesLine(txt string) bool {
	return EmptyOrSpacesLineRe.MatchString(txt)
}

func SplitLine(line string) []string {
	line = strings.Trim(line, "\r\n\x00")
	ret := strings.Split(line, ":")
	return ret
}

func StringValueValid(str string) bool {
	return !strings.ContainsAny(str, ":\n\r\x00")
}

func AtoiEmptyIsMinus1(value string, field_name string) (int, error) {
	if field_name == "" {
		return -1, nil
	} else {
		if t, err := strconv.Atoi(field_name); err != nil {
			return -100, errors.New(fmt.Sprintf("invalid %s", field_name))
		} else {
			if t < 0 {
				t = -1
			}
			return t, nil
		}
	}
}

func ItoaMinusIsEmpty(value int) string {
	if value < 0 {
		return ""
	} else {
		return strconv.Itoa(value)
	}
}

func ParseUserList(txt string) []string {
	s := set.NewSetString()
	s.AddStrings(strings.Split(txt, ",")...)
	ret := s.ListStrings()
	sort.Strings(ret)
	return ret
}

func RenderUserList(lst []string) (string, error) {
	for _, i := range lst {
		if !StringValueValid(i) {
			return "", errors.New("user names contains unacceptable symbols")
		}
	}

	s := set.NewSetString()
	s.AddStrings(lst...)
	lst = s.ListStrings()
	sort.Strings(lst)
	return strings.Join(lst, ","), nil
}
