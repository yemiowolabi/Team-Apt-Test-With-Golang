package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var mtnNum, airtelNum, gloNum, nineMobileNum, mtelNum int

var mtnPrefix = []string{"0703", "0706", "0803", "0806", "0810", "0813", "0814", "0816", "0903", "0906", "0913", "0704", "0916", "07025", "07026"}
var airtelPrefix = []string{"0701", "0708", "0802", "0808", "0812", "0901", "0902", "0904", "0907", "0912"}
var globacomPrefix = []string{"0705", "0805", "0807", "0811", "0815", "0905", "0915"}
var nineMobilePrefix = []string{"0809", "0817", "0818", "0909", "0908"}
var mtelPrefix = []string{"0804"}

func main() {
	url := "https://grnhse-ghr-prod-s101.s3.eu-central-1.amazonaws.com/generic_attachments/attachments/402/214/110/original/PhoneNumbers.txt?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAVQGOLGY3RZPEZZOZ%2F20220429%2Feu-central-1%2Fs3%2Faws4_request&X-Amz-Date=20220429T161547Z&X-Amz-Expires=604800&X-Amz-SignedHeaders=host&X-Amz-Signature=f4dc6bbadb1ac576d38baa613f01d9c75e8794abdda11a83b796d9415a5a3517"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		fmt.Printf("%s is DOWN!!!\n", url)
	} else {
		defer resp.Body.Close()
		fmt.Printf("%s =======> status code %d\n", url, resp.StatusCode)
		if resp.StatusCode == 200 {
			bodybytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}
			file := "PhoneNumbers.txt"
			fmt.Printf("Writing response body to %s\n", file) //To show what it is doing
			err = ioutil.WriteFile(file, bodybytes, 0664)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	file, err := os.Open("PhoneNumbers.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if checkIfItContains(mtnPrefix[:13], scanner.Text()[:4]) || checkIfItContains(mtnPrefix[13:], scanner.Text()[:5]) {
			mtnNum++
		} else if checkIfItContains(airtelPrefix, scanner.Text()[:4]) {
			airtelNum++
		} else if checkIfItContains(globacomPrefix, scanner.Text()[:4]) {
			gloNum++
		} else if checkIfItContains(nineMobilePrefix, scanner.Text()[:4]) {
			nineMobileNum++
		} else if checkIfItContains(mtelPrefix, scanner.Text()[:4]) {
			mtelNum++
		}
	}
	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("+" + strings.Repeat("=", 25) + "+" + strings.Repeat("=", 25) + "+")
	fmt.Println("|   NETWORK PROVIDER \t  |   PHONE NUMBERS COUNT   |")
	fmt.Println("+" + strings.Repeat("-", 25) + "+" + strings.Repeat("-", 25) + "+")
	fmt.Println("|\t MTN    \t     \t", mtnNum, "       \t    |")
	fmt.Println("+" + strings.Repeat("-", 25) + "+" + strings.Repeat("-", 25) + "+")
	fmt.Println("|\t AIRTEL \t  \t", airtelNum, "       \t    |")
	fmt.Println("+" + strings.Repeat("-", 25) + "+" + strings.Repeat("-", 25) + "+")
	fmt.Println("|\t GLOBACOM \t  \t", gloNum, "       \t    |")
	fmt.Println("+" + strings.Repeat("-", 25) + "+" + strings.Repeat("-", 25) + "+")
	fmt.Println("|\t 9MOBILE \t  \t", nineMobileNum, "       \t    |")
	fmt.Println("+" + strings.Repeat("-", 25) + "+" + strings.Repeat("-", 25) + "+")
	fmt.Println("|\t MTEL \t  \t\t", mtelNum, "\t \t    |")
	fmt.Println("+" + strings.Repeat("=", 25) + "+" + strings.Repeat("=", 25) + "+")

}
func checkIfItContains(a []string, b string) (c bool) {
	for _, value := range a {
		if b != value {
			c = false
		} else {
			c = true
			break
		}
	}
	return c
}
