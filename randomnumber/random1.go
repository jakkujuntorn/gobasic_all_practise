package main

import (
	_ "encoding/binary"
	_ "encoding/json"
	"fmt"
	"github.com/Pallinder/go-randomdata"
	_"github.com/fractalbach/fractalnet/namegen"
	"github.com/google/uuid"
	"math/rand"
	_ "strconv"
	"time"
)

type Product struct {
	IdStore     int
	Title       string 
	Price       int// randoms   100 - 5000
	Quantity    int // randoms 10-500
	Description string //randomdata.Paragraph()
	Category    string // ทำ []string ให้ซู่มเลือกเอาไปใส ด้วย การ random เลข index 
}

func main() {
	s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)
	_= r1

	for i := 0; i < 500; i++ {
		// fmt.Println(r1.Intn(1000))
		fmt.Println(rand.Intn(10))
	}

	fmt.Println(uuid.NewDCEPerson())
	// fmt.Println(uuid.NewString())
	// fmt.Println(uuid.NewRandom())
	// fmt.Println(uuid.Future.String())

	// fmt.Println(namegen.GenerateUsername())

	// fmt.Println("*********************")

	// for i := 0; i < 10; i++ {
	// 	v:=namegen.GenerateUsername()
	// 	fmt.Println(v)
	// }
	n := []int{2, 4, 5, 8}
_= n
	for i := 0; i < 100; i++ {
		// fmt.Println(randomdata.FirstName(100))
		// fmt.Println(randomdata.LastName())
		// fmt.Println("Adjective:", randomdata.Adjective())
		// fmt.Println("Phome number:", randomdata.PhoneNumber())
		// fmt.Println("SillyName:", randomdata.SillyName())
		// fmt.Println("Title: ", randomdata.Title(5))
		// fmt.Println("Alphanumeric: ", randomdata.Alphanumeric(len(n)))
		// fmt.Println("Currency: ", randomdata.Currency())
		// fmt.Println("Digits: ", randomdata.Digits(7))
		// fmt.Println("GenerateProfile: ", randomdata.GenerateProfile(7))
		// fmt.Println("Letters: ", randomdata.Letters(8))
		fmt.Println("Noun: ", randomdata.Noun())
		fmt.Println("Paragraph: ", randomdata.Paragraph())
		// fmt.Println("ProvinceForCountry: ", randomdata.ProvinceForCountry("US"))
		// fmt.Println("RandStringRunes: ", randomdata.RandStringRunes(7))
		// fmt.Println("StringNumberExt: ", randomdata.StringNumberExt(4, "john", 7))

		// fmt.Println("UserAgentString:", randomdata.UserAgentString())
		// fmt.Println("StringSample:", randomdata.StringSample())
		// fmt.Println("Address:", randomdata.Address())
		// fmt.Println("City:", randomdata.City())
		// fmt.Println("Local:", randomdata.Locale())
		// fmt.Println("PostalCode:", randomdata.PostalCode("TJ"))
		// fmt.Println("Email:", randomdata.Email())
		// fmt.Println(randomdata.Number(1, 2, 4, 5, 8))
		fmt.Println("*************************************")
	}

	// fmt.Println([]byte("2023-04-18 15:10:23"))
	// fmt.Println(time.Now())
	// fmt.Println(time.Now().Unix())
	// fmt.Println(time.Now().UnixNano())

	// fmt.Println([]byte("0")) // 48       //power day
	// fmt.Println([]byte("1")) // 49   // day // 49 -54
	// fmt.Println([]byte("2")) // 50
	// fmt.Println([]byte("3")) // 51
	// fmt.Println([]byte("4")) // 52
	// fmt.Println([]byte("5")) // 53
	// fmt.Println([]byte("6")) // 54
	// fmt.Println([]byte("7")) // 55
	// fmt.Println([]byte("8")) // 56
	// fmt.Println([]byte("9")) // 57
	// fmt.Println([]byte("21"))

	// fmt.Println("====================")
	// data:= []int{0,1,2,3,4,5,6,7,8,9}
	// for _,value:=range data {
	// 	b,_:=json.Marshal(value)
	// 	_=b
	// 	fmt.Println(b)
	// }

}
func GuessNumber() {
	n := []int{14,
		58,
		47,
		92,
		80,
		55, // 46
		// 73, // 21
	}
	number := guessNextNumber(n)
	fmt.Println(number)
}

func guessNextNumber(n []int) int {
	if len(n) < 2 {
		return 0
	}

	// diff := n[1] - n[0]
	// for i := 1; i < len(n)-1; i++ {
	// 	if n[i+1]-n[i] != diff {
	// 		return 0
	// 	}
	// }
	// return n[len(n)-1]+diff

	rand.Seed(time.Now().UnixNano())
	next := rand.Intn(n[len(n)-1] - n[len(n)-2] + n[len(n)-1])
	return next
}
