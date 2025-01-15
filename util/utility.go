package util

import (
	"fmt"
	"strconv"
	"unicode"
	"strings"
	"math"
	"regexp"
	"receipt-processor/model"
)

//probable batman's utility belt
func RegexEvaluate(regx string, toEvaluate string) bool {
	regex := regexp.MustCompile(regx)
	if !regex.MatchString(toEvaluate) {
		return true
	}
	return false
}

func ToFloat(someString string) float64 {
	val, err := strconv.ParseFloat(someString, 64)
	if err != nil {
		fmt.Printf("Error converting '%s' to float64: %v\n", someString, err)
		return 0
	}
	return val
}

func ToInt(someString string) int64 {
	val, err := strconv.ParseInt(someString, 10, 64) 
	if err != nil {
		fmt.Printf("Error converting '%s' to int64: %v\n", someString, err)
		return 0
	}
	return val
}

func CountAlphaNumericCharacters(someString string) int64{
	count := int64(0)
	for _, ch := range someString {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
			count++
		}
	}
	return count
}

func ProcessItemDescription(item model.Item) int64 {
	
	if len(strings.TrimSpace(item.ShortDescription))%3 == 0 {
		price := ToFloat(item.Price)

		//well you said round up does
		roundedPrice := math.Ceil(price * 0.2)

		fmt.Printf("Description: '%s', Original Price: %.2f, Rounded Price: %.0f\n", item.ShortDescription, price, roundedPrice)
		return int64(roundedPrice)
	}

	return 0
}

func GetDayasInt(someDateString string) int64 {
	return ToInt(someDateString[len(someDateString)-2:])
}

