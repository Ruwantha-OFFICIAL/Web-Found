// Copyright 2026 Lasith Ruwantha Amarawansha
// SPDX-License-Identifier: ISC

package requests

import (
  "github.com/go-resty/resty/v2"
  "sync"
  "fmt"
  "time"
  "strings"
  )
  
  
type OutPut struct{
  Url string
  Status int
}
  
func Send(res chan OutPut,url string,wg *sync.WaitGroup,redirect bool){
  defer wg.Done()
  client:= resty.New()
  client.SetTimeout(20 * time.Second)
  if redirect{
    client.SetRedirectPolicy(resty.NoRedirectPolicy())

  }
  req:= client.R()
  req.SetHeader("User-Agent","Mozilla/5.0 (Macintosh; Intel Mac OS X 10.9; rv:50.0) Gecko/20100101 Firefox/50.0")
  req.SetHeader("Accept","text/html,application/xhtml+xml,application/xml,application/json")
  
  respones,err:= req.Head(url)
  if err != nil && !strings.HasPrefix(err.Error(),"Head"){
    fmt.Println(err)
  }
  res<-OutPut{
    url,
    respones.StatusCode(),
  }
}