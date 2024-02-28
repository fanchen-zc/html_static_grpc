package util

import "regexp"

type re struct {
}

var Re re

func (re) Ip(ip string) bool {
	ipReg := `^(((1?\d{1,2})|(2[0-4]\d)|(25[0-5]))\.){3}((1?\d{1,2})|(2[0-4]\d)|(25[0-5]))$`
	match, _ := regexp.MatchString(ipReg, ip)
	return match
}
