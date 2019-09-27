package main

import (
    "interceptor/account"
    "interceptor/proxy"
)

func main() {
    id := "100111"
    a := account.New(id, "ZhangSan", 100)
    a.Query(id)
    a.Update(id, 500)
}

func init() {
    account.New = func(id, name string, value int) account.Account {
        a := &account.AccountImpl{id, name, value}
        p := &proxy.Proxy{a}
        return p
    }
}