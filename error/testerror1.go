package main

import (
	"encoding/json"
	"errors"
	"fmt"
)



type CustomError struct {
	Message string
	Code    int
	error
}

type Err_Russy struct {
	Message string
	Code int
	// error
}

func (er Err_Russy)Error()string  {
	return  er.Message
	// return  er.Message
}

func (ce *CustomError) MessageError1(m string, c int) *CustomError {
	return &CustomError{
		Message: m,
		Code:    c,
	}
}

func (ce *CustomError) MessageError2() CustomError {
	return CustomError{
		Message: ce.Message,
		Code:    ce.Code,
	}
}

// func (ce *CustomError) Error() string {
// 	return ce.Message
// }

// is error
func MessageError3() error {
	return errors.New("Erro 333")
}

// is error
func MessageError4() error {
	return fmt.Errorf("Error FMT")
}


func MessageError5(m string,n int) CustomError {
	return CustomError{
		Message: m,
		Code: n,
	}
}

//
type ErrChan struct {
	ErrChan chan error
}

type MyError struct {
	data json.RawMessage // คืออะไร
}

func (MyError) Error() string {
	return "My Error"
}

// ส่ง chan ที่เป็น error เข้ามาเช็ค type ว่าเป็น type อะไร
func ErrorChan(e chan error) {
	for {
		switch err := <-e; err.(type) {
		case *MyError:
			fmt.Println("MyError")
		case *CustomError:
			fmt.Println("CusTomError")
		default:
			fmt.Println("Undefined Error")
		}

	}
}

var Input1EmptyErr = errors.New("Input1 is empty")
var Input2EmptyErr = errors.New("Input2 is empty")

func DoSth(input1 string, input2 string) (err error) {
	// if input1 == "" {
	// 	err = errors.Join(err, Input1EmptyErr)
	// }
	// if input2 == "" {
	// 	err = errors.Join(err, Input2EmptyErr)
	
	// }
	return nil
}


