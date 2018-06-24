package util

import (
    "testing"
    "fmt"
)

func TestCheckIsSameRawUrl(t *testing.T) {
    referer := "http://localhost:8002/site/login"

    isLoginPage, urlLoginErr := IsLoginPage(referer)
    isRegisterPage, urlRegisterErr := IsRegisterPage(referer)

    fmt.Printf("isLoginPage:%v\n", isLoginPage)
    fmt.Printf("urlLoginErr:%+v\n", urlLoginErr)

    fmt.Printf("isRegisterPage:%v\n", isRegisterPage)
    fmt.Printf("urlRegisterErr:%+v\n", urlRegisterErr)
}
