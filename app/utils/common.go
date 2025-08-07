package utils

import (
	"book-app/config"
	"fmt"
	"math/rand"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	strip "github.com/grokify/html-strip-tags-go"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func GenerateRandomString(length int) string {
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		location = time.Local
		err = nil
		// fmt.Println("Failed to get Asia/Jakarta time when generating random string")
	}
	rand.Seed(time.Now().In(location).UnixNano())
	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZqwertyuiopasdfghjklzxcvbnm0123456789!@#$%^&*?"
	bytes := make([]byte, length)
	for i := 0; i < length; i++ {
		bytes[i] = charset[rand.Intn(len(charset))]
	}
	return string(bytes)
}

func GetTimeLocation() (response *time.Location) {
	response, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		response = time.Local
	}

	return
}

func Makerank(rank_step, min, max, target_value float64) int {
	range_size := (max - min) / rank_step
	rank := int((target_value - min) / range_size)
	if rank < 0 {
		rank = 0
	} else if rank >= int(rank_step) {
		rank = int(rank_step) - 1
	}
	return rank + 1
}

func GetRank(v int, rank_step, min, max float64) (data string) {
	for i := 1; i <= int(rank_step); i++ {
		if v == 1 {
			data = strconv.Itoa(int(min))
			return
		}
		if v == int(rank_step) {
			data = strconv.Itoa(int(max))
			return
		}
		stepSize := (max - min) / (rank_step - 1)
		data = strconv.Itoa(int(min + ((float64(v) - 1) * stepSize)))
	}
	return
}

func ConvertToCamelCase(s string) string {
	parts := strings.Split(s, "_")
	for i := range parts {
		parts[i] = TitleCase(parts[i])
	}
	return strings.Join(parts, "")
}

func TitleCase(str string) string {
	tc := cases.Title(language.Indonesian)
	return tc.String(str)
}

func RemoveDuplicates(str string) (data string) {
	s := strings.Split(str, ",")
	uniqueStrings := make([]string, 0, len(s))
	seen := make(map[string]bool)
	for _, str := range s {
		if !seen[str] {
			if str != "" {
				seen[str] = true
				uniqueStrings = append(uniqueStrings, str)
			}
		}
	}
	for _, v := range uniqueStrings {
		data += v + ","
	}
	data = strings.TrimRight(data, ",")
	return
}

func GetNumberFromStr(str string) (num int) {
	numStr := ""
	for _, char := range str {
		if char >= '0' && char <= '9' {
			numStr += string(char)
		}
	}
	num, _ = strconv.Atoi(numStr)

	return
}

func Average(numbers []float64) float64 {
	var sum float64
	for _, number := range numbers {
		sum += number
	}
	return sum / float64(len(numbers))
}

func StripTags(str string) string {
	str = strip.StripTags(str)
	return str
}

func LastId(table string) (id int) {
	type OnlyId struct {
		ID int
	}
	var last OnlyId
	config.DB.Table(table).Order("id desc").Limit(1).Scan(&last)

	id = last.ID + 1
	return
}

func GenerateRandomNumber(length int) string {
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		location = time.Local
		err = nil
		// fmt.Println("Failed to get Asia/Jakarta time when generating random pin")
	}
	rand.Seed(time.Now().In(location).UnixNano())
	charset := "0123456789"
	randomBytes := make([]byte, length)
	for i := range randomBytes {
		randomBytes[i] = charset[rand.Intn(len(charset))]
	}
	randomString := string(randomBytes)
	return randomString
}

func StripTagsFromStruct(input interface{}) {
	structValue := reflect.ValueOf(input).Elem()

	for i := 0; i < structValue.NumField(); i++ {
		fieldValue := structValue.Field(i)

		if fieldValue.Kind() == reflect.String {
			originalValue := fieldValue.String()
			strippedValue := strip.StripTags(originalValue)
			fieldValue.SetString(strippedValue)
		} else if fieldValue.Kind() == reflect.Struct {
			StripTagsFromStruct(fieldValue.Addr().Interface())
		}
	}
}

func GenerateInvoiceID(initial string, model interface{}) (invoiceID string) {
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		location = time.Local
		err = nil
		// fmt.Println("Failed to get Asia/Jakarta time when generating invoice id")
	}
	var count int64
	config.DB.Model(&model).Count(&count)
	if count == 0 {
		count = 1
	}
	var countString string
	if len(strconv.Itoa(int(count))) < 3 {
		countString = fmt.Sprintf("%0*d", 3, count)
	}

	timeString := time.Now().In(location).Format("0601021504")

	invoiceID = initial + timeString + countString

	return invoiceID
}

func IsStringInArray(str string, array []string) bool {
	for _, s := range array {
		if s == str {
			return true
		}
	}
	return false
}

func RoundToNextMidnight(t time.Time) time.Time {
	roundedTime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	roundedTime = roundedTime.Add(24 * time.Hour)
	return roundedTime
}

func ExecuteSQL(filename string, clean bool) (err error) {
	sqlFile, err := os.ReadFile(filename)
	if err != nil {
		return
	}
	if err = config.DB.Exec(string(sqlFile)).Error; err != nil {
		return
	}

	if clean {
		if err = os.Remove(filename); err != nil {
			return
		}
	}
	return
}
