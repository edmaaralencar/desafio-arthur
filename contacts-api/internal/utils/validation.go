package utils

import (
	"fmt"
	"regexp"
	"strings"
)

func FormatPhoneWithDDD(phone string) string {
	re := regexp.MustCompile(`\D`)
	digits := re.ReplaceAllString(phone, "")

	if len(digits) == 11 {
		return fmt.Sprintf("(%s) %s-%s", digits[0:2], digits[2:7], digits[7:])
	} else if len(digits) == 10 {
		return fmt.Sprintf("(%s) %s-%s", digits[0:2], digits[2:6], digits[6:])
	}

	return phone
}

func FormatCpfCnpj(doc string) string {
	digits := strings.Map(func(r rune) rune {
		if r >= '0' && r <= '9' {
			return r
		}
		return -1
	}, doc)

	if len(digits) == 11 {
		return digits[0:3] + "." + digits[3:6] + "." + digits[6:9] + "-" + digits[9:11]
	} else if len(digits) == 14 {
		return digits[0:2] + "." + digits[2:5] + "." + digits[5:8] + "/" + digits[8:12] + "-" + digits[12:14]
	}

	return doc
}
