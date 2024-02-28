package helper

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base32"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/leeqvip/gophp"
	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
	"html_static_grpc/pkg/util"
	"unicode"

	//"github.com/lionsoul2014/ip2region/tree/master/binding/golang/ip2region"
	"io/ioutil"
	"math/rand"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func MkMd5(str string) string {
	has := md5.Sum([]byte(str))
	return fmt.Sprintf("%x", has)
}

func IsEmpty(params interface{}) bool {
	//初始化变量
	var (
		flag          bool = true
		default_value reflect.Value
	)

	r := reflect.ValueOf(params)

	//获取对应类型默认值
	default_value = reflect.Zero(r.Type())
	//由于params 接口类型 所以default_value也要获取对应接口类型的值 如果获取不为接口类型 一直为返回false
	if !reflect.DeepEqual(r.Interface(), default_value.Interface()) {
		flag = false
	}
	return flag
}

// jwt加密
func JwtEncode(jwtinfo jwt.MapClaims, secret_key []byte) (jwt_token string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtinfo)
	tokenString, err := token.SignedString(secret_key)
	return tokenString, err
}

// jwt解密
func JwtDncode(token_string string, secret_key interface{}) (token_info map[string]interface{}, err error) {
	token, err := jwt.Parse(token_string, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return secret_key, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}

func Interface2String(inter interface{}) (str string) {
	str = ""
	switch inter.(type) {
	case string:
		str = inter.(string)
	default:

	}
	return str
}

/**
 * 获取谷歌验证码
 */
func MkGaCode(secret string) (code uint32, err error) {

	// decode the key from the first argument
	inputNoSpaces := strings.Replace(secret, " ", "", -1)
	inputNoSpacesUpper := strings.ToUpper(inputNoSpaces)
	fmt.Println(inputNoSpacesUpper)
	key, err := base32.StdEncoding.DecodeString(inputNoSpacesUpper)
	if err != nil {
		return 0, err
	}

	// generate a one-time password using the time at 30-second intervals
	epochSeconds := time.Now().Unix()
	code = oneTimePassword(key, toBytes(epochSeconds/30))

	return code, nil
}

func oneTimePassword(key []byte, value []byte) uint32 {
	// sign the value using HMAC-SHA1
	hmacSha1 := hmac.New(sha1.New, key)
	hmacSha1.Write(value)
	hash := hmacSha1.Sum(nil)

	// We're going to use a subset of the generated hash.
	// Using the last nibble (half-byte) to choose the index to start from.
	// This number is always appropriate as it's maximum decimal 15, the hash will
	// have the maximum index 19 (20 bytes of SHA1) and we need 4 bytes.
	offset := hash[len(hash)-1] & 0x0F

	// get a 32-bit (4-byte) chunk from the hash starting at offset
	hashParts := hash[offset : offset+4]

	// ignore the most significant bit as per RFC 4226
	hashParts[0] = hashParts[0] & 0x7F

	number := toUint32(hashParts)

	// size to 6 digits
	// one million is the first number with 7 digits so the remainder
	// of the division will always return < 7 digits
	pwd := number % 1000000

	return pwd
}

func toBytes(value int64) []byte {
	var result []byte
	mask := int64(0xFF)
	shifts := [8]uint16{56, 48, 40, 32, 24, 16, 8, 0}
	for _, shift := range shifts {
		result = append(result, byte((value>>shift)&mask))
	}
	return result
}

func toUint32(bytes []byte) uint32 {
	return (uint32(bytes[0]) << 24) + (uint32(bytes[1]) << 16) +
		(uint32(bytes[2]) << 8) + uint32(bytes[3])
}

func IsTimeStr(str string) bool {
	loc, _ := time.LoadLocation("Local")                                  //重要：获取时区
	theTime, err := time.ParseInLocation("2006-01-02 15:04:05", str, loc) //使用模板在对应时区转化为time.time类型
	if err != nil {
		return false
	}
	if theTime.Unix() > 0 {
		return true
	}
	return false
}

// 时间格式转换
func DateToDateTime(date string) string {
	timeTemplate := "2006-01-02T15:04:05+08:00" //常规类型
	toTemplate := "2006-01-02 15:04:05"
	stamp, _ := time.ParseInLocation(timeTemplate, date, time.Local)
	return time.Unix(stamp.Unix(), 0).Format(toTemplate)
}

// 时间格式转换
func TimestampToDateTime(date int64) string {
	toTemplate := "2006-01-02 15:04:05"
	return time.Unix(date, 0).Format(toTemplate)
}

// 日期转化为时间戳 类型是int64
func DateTime2Int64(str string) int64 {
	tmp, err := time.ParseInLocation("2006-01-02 15:04:05", str, time.Local)
	if err != nil {
		return 0
	}
	return tmp.Unix()
}

func GetAppDir() string {
	appDir, err := os.Getwd()
	if err != nil {
		file, _ := exec.LookPath(os.Args[0])
		applicationPath, _ := filepath.Abs(file)
		appDir, _ = filepath.Split(applicationPath)
	}
	return appDir
}

// 识别手机号码
func Check_phone(mobile string) bool {
	isorno, _ := regexp.MatchString(`^(1[3-9][0-9]\d{4,8})$`, mobile)
	if isorno {
		return true
	}
	return false
}

// 判断字符串是否含有中文
func IsChineseChar(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) {
			return true
		}
	}
	return false
}
func RandInt64(min, max int64) int64 {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	rand.Seed(time.Now().UnixNano())
	return rand.Int63n(max-min) + min
}

// md6加密
func Md6(str string) string {
	json, _ := json.Marshal(str)
	serialize, _ := gophp.Serialize(string(json))
	return MkMd5(string(serialize))
}

func TodayTime() int64 {
	t := time.Now()
	newTime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return newTime.Unix()
}

// 创建订单号
func CreateOrderNum() string {
	time := time.Now().Format("060102150405")
	rand := RandInt64(1000, 9999)
	return time + strconv.FormatInt(rand, 10)
}

// 随机字符串
func RandSeq(t string, n int) string {
	str := ""
	switch t {
	case "numeric":
		str = "0123456789"
	case "le":
		str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	case "w":
		str = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	case "l":
		str = "0123456789abcdefghijklmnopqrstuvwxyz"
	default:
		str = "abcdefghijklmnopqrstuvwxyz"
	}
	letters := []rune(str)
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// float32 转 String
func Float32ToString(num float32) string {
	return strconv.FormatFloat(float64(num), 'f', -1, 32)
}

// float64 转 String
func Float64ToString(num float64) string {
	return strconv.FormatFloat(num, 'f', -1, 64)
}

// int64 转 String
func Int64ToString(num int64) string {
	return strconv.FormatInt(num, 10)
}

// int32 转 String
func Int32ToString(num int32) string {
	return strconv.FormatInt(int64(num), 10)
}

// int 转 String
func IntToString(num int) string {
	return strconv.FormatInt(int64(num), 10)
}

// struct转map,key会大写
func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

// struct转map,key json
func StructToMapViaJson(obj interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	j, _ := json.Marshal(obj)
	json.Unmarshal(j, &m)
	return m
}

func BuildUrlCode(mp map[string]string) string {
	data := &url.Values{}
	for k, v := range mp {
		data.Set(k, v)
	}
	return data.Encode()
}

// 手机号中间四位星号
func GetSecretPhone(phone string) string {
	//reg := regexp.MustCompile(`(\d{3})\d{4}(\d{4})`)
	//return reg.ReplaceAllString(phone, `${1}****${2}`)
	if len(phone) <= 10 {
		return phone
	}
	return phone[:3] + "****" + phone[len(phone)-4:]
}

// 去重空和重复
func UniqueFilter(data []string) []string {
	mapData := make(map[string]interface{})
	for _, v := range data {
		mapData[v] = 1
	}
	var newWhites []string
	for k, _ := range mapData {
		if k == "" {
			continue
		}
		newWhites = append(newWhites, k)
	}
	return newWhites
}

func GetAreaByIp(ip string) string {
	var dbPath = "config/ip2region.xdb"
	searcher, err := xdb.NewWithFileOnly(dbPath)
	if err != nil {
		fmt.Printf("failed to create searcher: %s\n", err.Error())
		return ""
	}

	defer searcher.Close()

	// do the search
	var tStart = time.Now()
	region, err := searcher.SearchByStr(ip)
	if err != nil {
		fmt.Printf("failed to SearchIP(%s): %s\n", ip, err)
		return ""
	}
	fmt.Printf("{region: %s, took: %s}\n", region, time.Since(tStart))
	return region
	//本地ip库
	//appDir := GetAppDir()
	//if runtime.GOOS != "windows" {
	//	appDir = GetCurrentPath()
	//}
	//region, err := ip2region.New(appDir + "/config/ip2region.db")
	//defer region.Close()
	//if err != nil {
	//	log.Println(err)
	//	return ""
	//}
	////ipRes, err = region.BinarySearch(ip)
	////ipRes, err = region.BtreeSearch(ip)
	//ipRes, err := region.MemorySearch(ip)
	//if err != nil {
	//	return ""
	//}
	//
	//return ipRes.String() //1038|中国|0|江苏省|徐州市|电信
}

func HttpGet(url string, head bool) (body []byte, err error) {
	client := &http.Client{Timeout: time.Second * 5}
	var rqt *http.Request
	rqt, err = http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	if head {
		//灵匠标志
		rqt.Header.Add("User-Agent", "Lingjiang")
	}
	var response *http.Response
	response, err = client.Do(rqt)
	if err != nil {
		return
	}
	defer response.Body.Close()
	body, err = ioutil.ReadAll(response.Body)
	return
}

func StructToMap(in interface{}, tagName string) (map[string]interface{}, error) {
	out := make(map[string]interface{})

	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct { // 非结构体返回错误提示
		return nil, fmt.Errorf("ToMap only accepts struct or struct pointer; got %T", v)
	}

	t := v.Type()
	// 遍历结构体字段
	// 指定tagName值为map中key;字段值为map中value
	for i := 0; i < v.NumField(); i++ {
		fi := t.Field(i)
		if tagValue := fi.Tag.Get(tagName); tagValue != "" {
			out[tagValue] = v.Field(i).Interface()
		}
	}
	return out, nil
}

func GetZeroTimestamp() int64 {
	currentTime := time.Now()
	startTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location()).Unix()
	return startTime
}

// 以formdata格式提交数据
func PostWithFormData(url string, postData *map[string]string) []byte {
	body := new(bytes.Buffer)
	w := multipart.NewWriter(body)
	for k, v := range *postData {
		_ = w.WriteField(k, v)
	}
	err := w.Close()

	if err != nil {
		fmt.Println(err)
		return nil
	}
	req, err := http.NewRequest(http.MethodPost, url, body)
	req.Close = true
	if err != nil {
		return nil
	}
	req.Header.Set("Content-Type", w.FormDataContentType())
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil

	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}
	err = resp.Body.Close()
	if err != nil {
		return nil

	}
	return data
}

func CheckMobileUa(ua string) bool {
	if ua == "" {
		return false
	} else if strings.Contains(ua, "Mobile") ||
		strings.Contains(ua, "Android") ||
		strings.Contains(ua, "Silk/") ||
		strings.Contains(ua, "Kindle") ||
		strings.Contains(ua, "BlackBerry") ||
		strings.Contains(ua, "Opera Mini") ||
		strings.Contains(ua, "Opera Mobi") {
		return true
	} else {
		return false
	}
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

// 发送POST请求(带请求头)
// url:请求地址，data:POST请求提交的数据,
// contentType:
// (1)application/x-www-form-urlencoded  最常见的POST提交数据的方式，浏览器的原生form表单。后面可以跟charset=utf-8
// (2)multipart/form-data
// (3)application/json
// (4)text/xml    XML-RPC远程调用
// content:请求放回的内容
func HttpPostHeader(url string, data interface{}, contentType string, header map[string]interface{}) (content string, err error) {
	jsonStr, _ := json.Marshal(data)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("content-type", contentType)
	if len(header) > 0 {
		for k, v := range header {
			req.Header.Set(k, v.(string))
		}
	}
	if err != nil {
		return "", err
	}
	defer req.Body.Close()

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	result, _ := ioutil.ReadAll(resp.Body)
	content = string(result)
	return content, nil
}

func SendDingMessage(msg string, sn string) {
	var body = map[string]string{
		"sn":  sn,
		"msg": msg,
	}
	var body_byte, _ = json.Marshal(body)

	http.Post("http://api.qingjuhe.com/robot", "application/json", bytes.NewReader(body_byte))
}

// 向钉钉异常监控群发送钉钉群通知
func SendDingMsg(messageContent string, msgType int) (err error) {
	/*defer func() {
		if er := recover(); er != nil {
			err = errors.New("SendDingTalkMessage panic")
			return
		}
	}()*/
	messagePrefix := "通知"
	if msgType == -1 {
		messagePrefix = "异常"
	}

	access_token := "d857ce2783cd20ea9d437b33a8a042e1d8254b917029108e8b32aaa210bfbbff123a"
	apiUrl := "https://oapi.dingtalk.com/robot/send?access_token=" + access_token

	text := map[string]string{
		"content": messagePrefix + " : " + messageContent,
	}

	postData := map[string]interface{}{
		"msgtype": "text",
		"text":    text,
	}
	HttpPostHeader(apiUrl, postData, "application/json", map[string]interface{}{})

	return nil
}

func HttpGetHeader(url string, header map[string]interface{}) (body []byte, err error) {
	client := &http.Client{Timeout: time.Second * 5}
	var rqt *http.Request
	rqt, err = http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	if len(header) > 0 {
		for k, v := range header {
			rqt.Header.Set(k, v.(string))
		}
	}

	var response *http.Response
	response, err = client.Do(rqt)
	if err != nil {
		return
	}
	defer response.Body.Close()
	body, err = ioutil.ReadAll(response.Body)
	return
}
func GetCurrentPath() string {
	dir, _ := os.Executable()
	exPath := filepath.Dir(dir)
	return exPath
}

// 将页面内容保存到相对路径中
func SaveHtmlToPath(html string, staticFilePath string) error {
	util.CheckFile(staticFilePath) //检查文件夹是否存在,否则创建

	file, err := os.Create(staticFilePath)
	if err != nil {
		fmt.Println("Failed to create file:", err)
		return err
	}
	defer file.Close()

	// 将页面内容写入文件
	_, err = file.WriteString(html)
	if err != nil {
		fmt.Println("Failed to write to file:", err)
		return err
	}
	fmt.Println(staticFilePath + " successfully.")
	return nil
}
