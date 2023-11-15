package validate

import (
	"fmt"
	"regexp"
	"unicode/utf8"
)

type ErrprTest struct {
	msg string
}

func (pe *ErrprTest) Error() string {
	return pe.msg
}

//************************
type PasswordError struct {
	msg string
}

func (pe *PasswordError) Error() string {
	return pe.msg
}

func (pe *PasswordError) WrapBy(err error) {
	pe.msg = pe.msg + err.Error() + "\n"
}

// check 7 - 16 Char
func VAlidate_PAssword(pw string) error {
	pwError := PasswordError{}

	// เช็คจำนวน string
	err := check_Lenght(pw)
	if err != nil {
		// เพิ่ม error ลงไปใน pwError
		pwError.msg = err.Error() + "\n"
	}

	// check small letter
	err = contain_SmallLetter(pw)
	if err != nil {
		// เพิ่ม error ลงไปใน pwError
		pwError.msg = pwError.msg + err.Error() + "\n"
	}

	// check Big letter
	err = contain_BigLetter(pw)
	if err != nil {
		// เพิ่ม error ลงไปใน pwError
		pwError.msg = pwError.msg + err.Error() + "\n"
	}

	//***************************************
	// pwError มีค่า น่าจะเอาไว้อันล่างสุด
	if pwError.msg != "" {
		return &pwError
	}

	// กรณี ไม่มี error
	return nil
}
func regex_Helper(pw, pattern, msg string) error {
	re := regexp.MustCompile(pattern)
	result := re.FindString(pw)
	if result == "" {
		return fmt.Errorf("Password must contain Small letter")
	}
	return nil
}

// เช็ค Len
func check_Lenght(pw string) error {
	// lib เอาไว้นับ string
	pwLen := utf8.RuneCountInString(pw)
	if pwLen < 7 || pwLen > 16 {
		return fmt.Errorf("Password Lenght should be 7 - 16")
	}
	return nil
}

// เช็ค ตัวอักษรเล็ก
func contain_SmallLetter(pw string) error {
	// ใช้ regex  มาใช้ [a-z]
	// re := regex_Helper(`[a-z]`)
	// result := re.FindString(pw,`[a-z]`,"Password must contain Small letter")
	// if result == "" {
	// 	return fmt.Errorf("Password must contain Small letter")
	// }

	// func helper มัน return error ออกมาอยู่แล้วเลยเอา Func return ออกไปได้เลย
	return regex_Helper(pw, "[a-z]", "Password must contain Small letter")

}

// เช็ค ตัวอักษรใหญ๋
func contain_BigLetter(pw string) error {
	// func helper มัน return error ออกมาอยู่แล้วเลยเอา Func return ออกไปได้เลย
	return regex_Helper(pw, "[A-Z]", "Password must contain Big letter")

}

// เช็คว่ามีตัวเลขหรือไม่
func contain_Digit(pw string) error {

	return regex_Helper(pw, "[0-9]", "PAssword must contain digit")
}
