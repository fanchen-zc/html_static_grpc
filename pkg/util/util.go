package util

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/big"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

func GetCurrPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func CheckEmail(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

func CheckPwd(password string) bool {
	pattern := `^[a-zA-Z0-9]{6,20}$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(password)
}

// 生成指定区间随机数（包括纯数字／纯字母／随机）
func Krand(size int, kind int) []byte {
	ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	is_all := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if is_all { // random ikind
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return result
}

// 生成区间随机数
func RandInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Intn(max-min) + min
}

func Md5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func Md6(s string) string {
	ss := ""
	for _, item := range s {
		if item == 34 {
			ss += "\\"
		}
		ss += string(item)
	}
	s = fmt.Sprintf("\"%s\"", ss)
	return Md5(fmt.Sprintf("s:%d:\"%s\";", len(s), s))
}

// 获取目前时间戳
func NowTime() int32 {
	nowTime := int32(time.Now().Unix())
	return nowTime
}

func TodayWeeHourTimestamp() int {
	formatLayout := "2006-01-02"
	today := time.Now().Format(formatLayout)
	t, _ := time.Parse(formatLayout, today)
	return int(t.Unix() - 8*3600)
}

func GetMidNight(target time.Time) int64 {
	formatLayout := "2006-01-02"
	today := target.Format(formatLayout)
	t, _ := time.Parse(formatLayout, today)
	return t.Unix() - 8*3600
}

// 发送POST请求
// url:请求地址，data:POST请求提交的数据,
// contentType:
// (1)application/x-www-form-urlencoded  最常见的POST提交数据的方式，浏览器的原生form表单。后面可以跟charset=utf-8
// (2)multipart/form-data
// (3)application/json
// (4)text/xml    XML-RPC远程调用
// content:请求放回的内容
func HttpPost(url string, data interface{}, contentType string) (content string) {
	jsonStr, _ := json.Marshal(data)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Add("content-type", contentType)
	if err != nil {
		panic(any(err))
	}
	defer req.Body.Close()

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		panic(any(err))
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	content = string(result)
	return
}

func HttpPostForm(postUrl string, param map[string]string) (string, int) {
	data := make(url.Values)
	for k, v := range param {
		data[k] = []string{v}
	}
	resp, err := http.PostForm(postUrl, data)
	if err != nil {
		panic(any(err))
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(any(err))
	}

	return string(body), resp.StatusCode
}

// Post 请求 httpPostForm(带请求头)
func HttpPostFormHeader(postUrl string, param map[string]string, header map[string]interface{}) (err error, result string) {
	data := make(url.Values)
	for k, v := range param {
		data[k] = []string{v}
	}
	req, err := http.NewRequest("POST", postUrl, strings.NewReader(data.Encode()))
	if err != nil {
		return err, ""
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if len(header) > 0 {
		for k, v := range header {
			req.Header.Set(k, v.(string))
		}
	}
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return err, ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err, ""
	}
	return nil, string(body)
}

func HttpGET(url string) string {
	client := &http.Client{Timeout: 5 * time.Second}

	resp, err := client.Get(url)
	if err != nil {
		panic(any(err))
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK || err != nil {
		panic(any(err))
	}

	return string(body)
}

// 可以设置请求头
func HttpGetHeader(postUrl string, header map[string]interface{}) string {
	client := http.Client{}
	req, err := http.NewRequest(http.MethodGet, postUrl, nil)
	if err != nil {
		panic(any(err))
	}
	// 添加请求头
	//req.Header.Add("Content-type", "application/json;charset=utf-8")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if len(header) > 0 {
		for k, v := range header {
			req.Header.Set(k, v.(string))
		}
	}
	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		panic(any(err))
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(any(err))
	}
	return string(body)
}

// 伪装浏览器GET
func BrowserHttpGet(url string) (error, string) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err, ""
	}
	request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	request.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3")
	request.Header.Add("Connection", "keep-alive")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")

	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(request)

	if err != nil {
		return err, ""
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK || err != nil {
		return err, ""
	}

	return nil, string(body)
}

func ApiSignVerify() {
}

func Str2Int(num string, defaultNum int) (numInt int) {
	numInt, err := strconv.Atoi(num)
	if err != nil {
		return defaultNum
	}
	return
}

func Str2Float64(num string) (value float64) {
	value, _ = strconv.ParseFloat(num, 64)
	return value
}

/*
字符串ip转整形IP
*/
func InetAtoN(ip string) uint32 {
	ret := big.NewInt(0)
	ret.SetBytes(net.ParseIP(ip).To4())
	return uint32(ret.Uint64())
}

func IntToStr(intValue int64) string {
	return fmt.Sprintf("%d", intValue)
}

func ToJsonStr(entity interface{}) string {
	str, err := json.Marshal(entity)
	if err != nil {
		return ""
	}
	return string(str)
}

func InArrayStr(s string, list []string) bool {
	for _, val := range list {
		if val == s {
			return true
		}
	}
	return false
}

func DeduplicationStr(list []string) []string {
	sort.Strings(list)
	var previous = ""
	var retList []string
	for i := 0; i < len(list); i++ {
		if list[i] != previous {
			retList = append(retList, list[i])
		}
		previous = list[i]

	}
	return retList
}

func PhoneRemoveSensitiveInformation(phone string) string {
	if phone == "" {
		return ""
	}
	return phone[:3] + "****" + phone[7:]
}

func MapToUrlParam(data map[string]string) string {
	var param = ""
	for key, value := range data {
		param += fmt.Sprintf("%s=%s&", key, value)
	}
	return param
}
func HideStar(str string) (result string) {
	if str == "" {
		return "***"
	}
	if strings.Contains(str, "@") {
		res := strings.Split(str, "@")
		if len(res[0]) < 3 {
			resString := "***"
			result = resString + "@" + res[1]
		} else {
			res2 := Substr2(str, 0, 3)
			resString := res2 + "***"
			result = resString + "@" + res[1]
		}
		return result
	} else {
		reg := `^1[0-9]\d{9}$`
		rgx := regexp.MustCompile(reg)
		mobileMatch := rgx.MatchString(str)
		if mobileMatch {
			result = Substr2(str, 0, 3) + "****" + Substr2(str, 7, 11)
		} else {
			nameRune := []rune(str)
			lens := len(nameRune)

			if lens <= 1 {
				result = "***"
			} else if lens == 2 {
				result = string(nameRune[:1]) + "*"
			} else if lens == 3 {
				result = string(nameRune[:1]) + "*" + string(nameRune[2:3])
			} else if lens == 4 {
				result = string(nameRune[:1]) + "**" + string(nameRune[lens-1:lens])
			} else if lens > 4 {
				result = string(nameRune[:2]) + "***" + string(nameRune[lens-2:lens])
			}
		}
		return
	}
}

func Substr2(str string, start int, end int) string {
	rs := []rune(str)
	return string(rs[start:end])
}

/**
 *	自定义正则验证
 */
func CustomMatch(str, match string) bool {
	reg := match
	rgx := regexp.MustCompile(reg)
	return rgx.MatchString(str)
}

func QueryParamStrToMap(s string) map[string]string {

	var data = make(map[string]string)

	for _, item := range strings.Split(s, "&") {
		var itemList = strings.Split(item, "=")

		var value = ""
		if len(itemList) >= 2 {
			value = itemList[1]
		}

		data[itemList[0]] = value
	}
	return data
}

func MapToQueryParam(data map[string]string) string {
	var keyList []string
	for key, _ := range data {
		keyList = append(keyList, key)
	}

	sort.Sort(sort.StringSlice(keyList))

	var str = ""
	for _, val := range keyList {
		if val == "Signature" {
			continue
		}

		if str != "" {
			str += "&"
		}
		str += fmt.Sprintf("%s=%s", val, data[val])
	}
	return str
}

/*
**
是否在邮箱黑名单
*/
func EmailInBlacklist(email string, black []string) bool {
	for _, s := range black {
		if strings.Contains(email, s) {
			return true
		}
	}
	return false
}

// float保留两位小数
func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}
