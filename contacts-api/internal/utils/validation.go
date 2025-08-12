package utils

import (
	"fmt"
	"regexp"
	"strings"
)

func onlyNumbers(s string) string {
	re := regexp.MustCompile(`\D`) 
	return re.ReplaceAllString(s, "")
}

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

func ValidateCpfCnpj(doc string) (bool, string) {
	digits := regexp.MustCompile(`\D`).ReplaceAllString(doc, "")

	if len(digits) == 11 {
		return validateCPF(digits), digits
	}
	if len(digits) == 14 {
		return validateCNPJ(digits), digits
	}
	return false, ""
}

func validateCPF(cpf string) bool {
	cpf = onlyNumbers(cpf)
	if len(cpf) != 11 {
		return false
	}

	invalids := []string{
		"00000000000", "11111111111", "22222222222",
		"33333333333", "44444444444", "55555555555",
		"66666666666", "77777777777", "88888888888", "99999999999",
	}
	for _, inv := range invalids {
		if cpf == inv {
				return false
		}
	}

	sum := 0
	for i := 0; i < 9; i++ {
		sum += int(cpf[i]-'0') * (10 - i)
	}
	d1 := (sum * 10 % 11) % 10
	if d1 != int(cpf[9]-'0') {
		return false
	}
	sum = 0
	for i := 0; i < 10; i++ {
		sum += int(cpf[i]-'0') * (11 - i)
	}
	d2 := (sum * 10 % 11) % 10
	return d2 == int(cpf[10]-'0')
}

func validateCNPJ(cnpj string) bool {
	cnpj = onlyNumbers(cnpj)
	if len(cnpj) != 14 {
		return false
	}
	invalids := []string{
		"00000000000000", "11111111111111", "22222222222222",
		"33333333333333", "44444444444444", "55555555555555",
		"66666666666666", "77777777777777", "88888888888888",
		"99999999999999",
	}
	for _, inv := range invalids {
		if cnpj == inv {
				return false
		}
	}

	var calcDigit = func(weights []int, numbers string) int {
		sum := 0
		for i, w := range weights {
			sum += w * int(numbers[i]-'0')
		}
		res := sum % 11
		if res < 2 {
			return 0
		}
		return 11 - res
	}

	weights1 := []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	weights2 := []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}

	firstDigit := calcDigit(weights1, cnpj[:12])
	if firstDigit != int(cnpj[12]-'0') {
		return false
	}

	secondDigit := calcDigit(weights2, cnpj[:13])
	if secondDigit != int(cnpj[13]-'0') {
		return false
	}

	return true
}
