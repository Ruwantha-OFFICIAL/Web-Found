// Copyright 2026 Lasith Ruwantha Amarawansha
// SPDX-License-Identifier: ISC

package main

import (
  "os"
  "fmt"
  "github.com/go-resty/resty/v2"
  "strings"
  )

type Web struct{
  url string
  path string
}

  
func (w Web) CheckFile() bool{
  file,err:= os.Open(w.path)
  
  if err != nil{
    fmt.Println("\033[0;31m[x] Can't Open File\033[0m")
    return false
  }
  
  defer file.Close()
  
  _,err = file.Stat()
  
  if err != nil{
    fmt.Print("\033[0;31m[x] Can't Read File Detalls\033[0m")
    return false
  }
  
  fmt.Printf("\033[0;32m[✔] File Check Success\033[0m\n")
  return true
  
}

func (w Web) CheckUrl()bool{
  if !strings.HasPrefix(w.url, "https://"){
    fmt.Println("\033[0;31m[x] Url Must Not Start With https.")
    return false
  }
  //req build
  client:= resty.New()
  req:= client.R()
  req.SetHeader("User-Agent","Mozilla/5.0 (Macintosh; Intel Mac OS X 10.9; rv:50.0) Gecko/20100101 Firefox/50.0")
  req.SetHeader("Accept","text/html,application/xhtml+xml,application/xml,application/json")
  
  _,err:= req.Head(w.url)
  if err != nil{
    fmt.Println("\033[0;31m[x] Url Chek Faild!")
    fmt.Printf("Erro: %s\n",err)
    return false
  }
  fmt.Println("\033[0;32m[✔] Web Site Check Success")
  return true
}

