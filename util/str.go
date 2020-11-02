package util

import (
"crypto/md5"
"encoding/hex"
"strconv"
"strings"
"unicode/utf8"
)

type StrUtil struct {
}

func (u *StrUtil) Substr(str string, start, end int) string {
	if str == "" {
		return ""
	}

	rs := []rune(str)
	length := len(rs)
	if start < 0 || start > length {
		panic("start is valid")
	}
	if end < 0 || end > length {
		panic("end is valid")
	}
	return string(rs[start:end])
}

func (u *StrUtil) Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func (u *StrUtil) SubstrReplace(str string, start, end int, newStr string, n int) string {
	substr := u.Substr(str, start, end)
	return strings.Replace(str, substr, newStr, n)
}

func (u *StrUtil) ParseFloat(s string, bitSize int) float64 {
	value, err := strconv.ParseFloat(s, bitSize)
	if err != nil {
		return 0
	}
	return value
}

func (u *StrUtil) DesensitizationPhone(phone string) string {
	if phone == "" || len(phone) < 11 {
		return phone
	}
	formatPhone := phone[0:3] + "******" + phone[9:len(phone)]
	return formatPhone
}

func (u *StrUtil) DesensitizationEmail(email string) string {
	if email == "" || strings.Index(email, "@") == -1 {
		return email
	}

	data := strings.Split(email, "@")
	partOneLength := len(data)
	if partOneLength == 0 || partOneLength != 2 {
		return email
	}
	partOne := ""
	if partOneLength <= 3 {
		partOne = data[0][0:1] + "**"
	} else {
		partOne = data[0][0:3] + "***"
	}
	return partOne + "@" + data[1]
}

func (u *StrUtil) DesensitizationUserName(userName string) string {
	if userName == "" {
		return userName
	}
	length := utf8.RuneCountInString(userName)
	runeStr := []rune(userName)
	if length <= 1 {
		return userName
	}
	if length == 2 {
		return string(runeStr[0]) + "*"
	}
	return string(runeStr[0]) + "**" + string(runeStr[length-1])
}
