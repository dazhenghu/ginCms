package util

import (
    "net/url"
    "strings"
    adminConsts "github.com/dazhenghu/ginCms/admin/consts"
)

/**
判断是否是登录页
 */
func IsLoginPage(rawUrl string) (res bool, err error)  {
    return CheckIsSameRawUrl(rawUrl, adminConsts.URL_LOGIN)
}

/**
判断是否是注册
 */
func IsRegisterPage(rawUrl string) (res bool, err error)  {
    return CheckIsSameRawUrl(rawUrl, adminConsts.URL_REGISTER)
}

/**
判断两个url是否是相同的，不区分大小写
 */
func CheckIsSameRawUrl(sourceUrl, targetUrl string) (res bool, err error) {
    urlInfo, err := url.Parse(sourceUrl)
    if err != nil {
        return
    }

    path := urlInfo.Path

    if strings.ToLower(strings.Trim(path, "/")) == strings.ToLower(strings.Trim(targetUrl, "/")) {
        res = true
    } else {
        res = false
    }

    return
}