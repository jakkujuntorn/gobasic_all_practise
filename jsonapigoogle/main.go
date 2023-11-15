package main

import (
	"bytes"
	_ "encoding/json"
	"fmt"
	"os"

	"time"

	"github.com/gofiber/fiber/v2"
	// fiber "github.com/gofiber/fiber/v2"
	"github.com/google/jsonapi"
)

func main() {
	fmt.Println("")

	// testBlog()

	app := fiber.New()

	// app.Get("/user", GetBlog)
	// app.Post("/user", PostBlog)

	// appCustom := NewFiberCustom()
	// _ = appCustom
	//*****************

	// ************* Fiber V2 ***************
	customHandler := NewHandlerFiberV2()
	// ต้อง แปลง เป็น fiber.Handler ก่อน
	app.Get("/russy", ConvertFiberV2(customHandler.Handler_Change_Byte_to_string))
	app.Get("/getdata", GetBlog)
	app.Post("/user", ConvertFiberV2(customHandler.PostDataFiberV2))

	// **************** Fiber V3 ************
	f3 := NewFiberV3()
	_ = f3
	f33 := fiberV3{}

	app.Get("/v3", ConvertFiberV3(f33.GetData_FiberV3)) // work

	// *************** fiber V4 ****************
	fV2 := NewFiberCustom2()
	f4 := NewFiberV4(fV2)
	app.Get("/v4", f4.JsobFiberV4) // error
	app.Listen(":8000")

}

//****************************** Fiber V4 **************************
type fiberV4 struct {
	I_FiberV2
}

func NewFiberV4(in I_FiberV2) *fiberV4 {
	return &fiberV4{in}
}

func (i *fiberV4) JsobFiberV4(c *fiber.Ctx) error {
	dd := "Fiber V4"
	i.I_FiberV2.JsonFiberV2(dd)
	// c.JSON(dd)
	return nil
}

// *************************** Fiber V3 *********************************
// test fiber ทำเป็น struct
//******** fiber ส่งเป็น struct ไม่ได้ ******
type fiberV3 struct {
	*fiber.Ctx
}

func NewFiberV3() *fiberV3 {
	return &fiberV3{}
}

func (f *fiberV3) GetData_FiberV3(fi I_FiberV2, c *fiber.Ctx) {
	// f.Ctx.Status(200).JSON("Test V3")
	c.Status(200).JSON("Test V3333")
}

// *************************** Fiber V2 *****************************
// struct เอาไปใช้ใน Func ที่ใช้ convert
type FiberV2 struct {
	Ctx *fiber.Ctx
}

// ใช้เป็น paramiter จะทำงาน
// แต่ทำงานผ่าน struct โดยใสใน struct จะไม่ทำงาน เช่น Fiber V4 ***
type I_FiberV2 interface {
	Change_Byte_to_String([]byte) (string, error)
	JsonFiberV2(i interface{})
	// BodyParserFiberV2(interface{}) error // ถ้าเป็น interface ต้องใช้ * ทั้งตัวรับและตัวส่าง
	BodyParserFiberV2(*Comment) error // เป็น struct แค่ส่ง pointer
}

func (f *FiberV2) Change_Byte_to_String(by []byte) (string, error) {
	s := fmt.Sprintf(string(by))
	return s, nil
}

func (f *FiberV2) JsonFiberV2(i interface{}) {
	f.Ctx.JSON(i)
}

// func (f *FiberV2) BodyParserFiberV2(i interface{}) error {
func (f *FiberV2) BodyParserFiberV2(i *Comment) error {
	if err := f.Ctx.BodyParser(i); err != nil {
		return err
	}
	return nil
}

// อันนี้ไม่ใช้ก็ได้เพราะมันไมได้เอาไปต่อกับใคร
// ใช้แค่เป็นตัว return ใน func แปลง
func NewFiberCustom(c *fiber.Ctx) I_FiberV2 {
	return &FiberV2{Ctx: c}
}

func NewFiberCustom2() I_FiberV2 {
	return &FiberV2{}
}

func ConvertFiberV2(handlerFunc func(I_FiberV2)) fiber.Handler {
	return func(c *fiber.Ctx) error {

		// func นี้ ต้องทำ recive Func ให้ confrom ตาม interface I_FiberV2
		// handlerFunc(NewFiberCustom(c))

		// struct นี้ ต้องทำ recive Func ให้ confrom ตาม interface I_FiberV2
		handlerFunc(&FiberV2{Ctx: c})

		return nil
	}
}

func ConvertFiberV3(handlerFunc func(I_FiberV2, *fiber.Ctx)) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// func นี้ ต้องทำ recive Func ให้ confrom ตาม interface I_FiberV2
		// handlerFunc(NewFiberCustom(c))

		// struct นี้ ต้องทำ recive Func ให้ confrom ตาม interface I_FiberV2
		handlerFunc(&FiberV2{Ctx: c}, c)

		return nil
	}
}

// ***************************************************************
//****************** Handler ***********************
type handlerCustom struct {
	// เอาของจาก layer มาใช้
	Ctx *fiber.Ctx
}

// type IHandlerCustomer interface {
// 	Json_Customer(interface{})
// }

func NewHandlerFiberV2() *handlerCustom {
	return &handlerCustom{}
}

// *******Func fiber.Handler *******
func (h *handlerCustom) Handler_Change_Byte_to_string(f I_FiberV2) {

	data := []byte("Russy")
	sData, _ := f.Change_Byte_to_String(data)

	// fmt.Println(sData)

	// ใช้ของ struct ตัวเองไม่ได้
	// h.Ctx.JSON(fiber.Map{
	// 	"message": sData,
	// })

	// ********** แต่ทำไม struct ด้านบนใช้ได้ ***********
	// h.Ctx.JSON(sData)

	//ต้องใช้ของ I_FiberV2
	f.JsonFiberV2(sData)

}

func (h *handlerCustom) PostDataFiberV2(f I_FiberV2) {
	dataComment := Comment{}
	err := f.BodyParserFiberV2(&dataComment)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(dataComment)
	f.JsonFiberV2(dataComment)
}

//*****************************************************

func GetBlog(c *fiber.Ctx) error {

	dataBlog1 := Blog{
		ID:     1150,
		Title:  "Harry Gin",
		Titles: "Harry Gin",
		Posts: []Post{
			{
				ID:     1,
				BlogID: 11,
				Title:  "Head",
				Body:   "Good work",
				Comments: []*Comment{
					{
						ID:     1,
						PostID: 45,
						Body:   "it ok",
					},
					{
						ID:     2,
						PostID: 45,
						Body:   "work",
					},
				},
			},
			{
				ID:     1,
				BlogID: 11,
				Title:  "Head",
				Body:   "Good work",
				Comments: []*Comment{
					{
						ID:     1,
						PostID: 45,
						Body:   "it ok",
					},
					{
						ID:     2,
						PostID: 45,
						Body:   "work",
					},
				},
			},
		}, // Post

		CurrentPost: &Post{
			ID:     11,
			BlogID: 50,
			Title:  "Gin loconic",
			Body:   "It is  New Book for all",
			Comments: []*Comment{
				{
					ID:     1,
					PostID: 45,
					Body:   "it ok",
				},
				{
					ID:     2,
					PostID: 45,
					Body:   "work",
				},
			},
		},
		CurrentPostID: 450,
		CreatedAt:     time.Now().Local().UTC(),
		ViewCount:     555,
	}

	// post:=Post2{
	// 	ID:     11,
	// 	BlogID: 50,
	// 	Title:  "Gin loconic",
	// 	Body:   "It is  New Book for all",
	// 	Comments: []*Comment{
	// 		{
	// 			ID:     1,
	// 			PostID: 45,
	// 			Body:   "it ok",
	// 		},
	// 		{
	// 			ID:     2,
	// 			PostID: 45,
	// 			Body:   "work",
	// 		},
	// 	},
	// }

	payload, err := jsonapi.Marshal(dataBlog1)
	if err != nil {
		return c.Status(500).JSON(err)
	}

	return c.Status(200).JSON(payload)

}

func PostBlog(c *fiber.Ctx) error {

	// dataComment := new(Comment)
	dataComment := Comment{}

	// ********* fiber
	// if err := c.BodyParser(&dataComment); err != nil {
	// 	return c.Status(500).JSON(err)
	// }
	// fmt.Println(dataComment)

	//*************jsonpai google

	jsonapiRuntime := jsonapi.NewRuntime().Instrument("create_comment")

	errPayload := jsonapiRuntime.UnmarshalPayload(bytes.NewReader(c.Body()), &dataComment)
	// errPayload := jsonapi.UnmarshalPayload(strings.NewReader(c.Body()), dataComment)
	if errPayload != nil {
		return c.Status(500).JSON(errPayload.Error())
	}

	// buf := &bytes.Buffer{}
	err := jsonapiRuntime.MarshalPayload(os.Stdout, &dataComment) // ตรงนี้ return ค่าออกมาด้วย
	// err := jsonapi.MarshalPayload(c,&dataComment)// ตรงนี้ return  ถ้าไม่ใส &=> models should be a struct pointer or slice of struct pointers
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println(dataComment)
	// fmt.Println("=======================")

	// 	b,_:=json.Marshal(dataComment)

	// fmt.Println(string(b))

	// err=json.Unmarshal(b,dataComment)
	// fmt.Println(dataComment)

	// payload := &jsonapi.Payload{
	// 	Data: dataComment,
	// }

	// data := &Comment{
	// 	ID:     1150,
	// 	PostID: 11,
	// 	Body:   "Is work",
	// }

	// ww:=bytes.NewReader([]byte(dada))

	// newFiberV2:=NewFiberV2(c)

	// err:= ConvertFiberHandler2(newFiberV2)
	// if err != nil {
	// fmt.Println(err)
	// }

	// err:=jsonapi.MarshalPayload(bytes.NewReader(c.Body()),dataComment)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// buf := new(bytes.Buffer)
	// json.NewEncoder(buf).Encode(dataComment)
	// // io.WriteString(buf,&dataComment)

	c.Set(fiber.HeaderContentType, jsonapi.MediaType)

	// fmt.Println(dataComment)

	return c.Status(200).JSON(dataComment)
}

func testBlog() {
	person := Person2{
		FirstName: "John",
		LastName:  "Doe",
	}

	// data, err := jsonapi.MarshalPayload(os.Stdout,&person)
	// if err != nil {
	// fmt.Println("err")
	// return
	// }

	err := jsonapi.MarshalPayload(os.Stdout, &person)
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(string(data))

	// fmt.Println(person)

	sData := strings.NewReader(string("data"))
	_ = sData

	//     var person22 Person2
	//     err = jsonapi.UnmarshalPayload(sData, &person22)
	//     if err != nil {
	//         panic(err)
	//     }

	// fmt.Printf("First name: %s, Last name: %s\n", person22.FirstName, person2.LastName)

}

type Blog struct {
	ID            int       `jsonapi:"primary,blogs"`
	Title         string    `jsonapi:"attr,titless"`
	Titles        string    `json:"titleRussy"`
	Posts         []Post    `jsonapi:"relation,posts"`
	CurrentPost   *Post     `jsonapi:"relation,current_post"`
	CurrentPostID int       `jsonapi:"attr,current_post_id"`
	CreatedAt     time.Time `jsonapi:"attr,created_at"`
	ViewCount     int       `jsonapi:"attr,view_count"`
}

type Post struct {
	ID       int        `jsonapi:"primary,posts"`
	BlogID   int        `jsonapi:"attr,blog_id"`
	Title    string     `jsonapi:"attr,title"`
	Body     string     `jsonapi:"attr,body"`
	Comments []*Comment `jsonapi:"relation,comments"`
	// Comments []Comment `jsonapi:"relation,comments"`
}

type Post2 struct {
	ID       int        `jsonapi:"primary,posts"`
	BlogID   int        `jsonapi:"attr,blog_id"`
	Title    string     `jsonapi:"attr,title"`
	Body     string     `jsonapi:"attr,body"`
	Comments []*Comment `jsonapi:"relation,comments"`
}

type Comment struct {
	ID     int `jsonapi:"primary,comment,omitempty" json:"id"`
	PostID int `jsonapi:"attr,post_id"  json:"postid"`
	// PostID int    `json:"post_id"`
	Body string `jsonapi:"attr,bodysss"  json:"body"`
}

type Person struct {
	ID        int    `jsonapi:"primary,person"`
	ID_Person int    `jsonapi:"attr,id"` // input ต้องใช้ชื่อของ json เท่านั้น
	FirstName string `jsonapi:"attr,first_name"`
	LastName  string `jsonapi:"attr,last_name"`
}

type Data struct {
	FirstName string `jsonapi:"attr,fname"`
	LastName  string `jsonapi:"attr,lname"`
}

// type Comment struct {
// 	ID     int
// 	PostID int
// 	// PostID int    `json:"post_id"`
// 	Body   string
// }

// custom fiber
type fiberV2 struct {
	*fiber.Ctx
}

type IFiberV2 interface {
	Writer(p []byte) (n int, err error)
}

func (*fiberV2) Writer(p []byte) (int, error) {
	return len(p), nil
}

func NewFiberV2(c *fiber.Ctx) IFiberV2 {
	return &fiberV2{Ctx: c}
}

func ConvertFiberHandler2(handlerFunc func(IFiberV2)) fiber.Handler {
	return func(c *fiber.Ctx) error {
		//  handlerFunc(NewFiberCustom(c))
		handlerFunc(&fiberV2{Ctx: c})
		return nil
	}
}

type Person2 struct {
	FirstName string `jsonapi:"name,first_name"`
	LastName  string `jsonapi:"name,last_name"`
}

func (p Person) GetID() string {
	return fmt.Sprintf("%s %s", p.FirstName, p.LastName)
}
