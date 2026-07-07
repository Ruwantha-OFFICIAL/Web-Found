// Copyright 2026 Lasith Ruwantha Amarawansha
// SPDX-License-Identifier: ISC
package main 

import (
  "fmt"
  "flag"
  "github.com/mbndr/figlet4go"
  "web-found/internal/readfile"
  "web-found/internal/requests"
  "sync"
  "embed"
  )

//go:embed assets/3-d.flf
var FontFS embed.FS

var catch = []int{200,403,204,401,301,302}

func main(){
  logo()
  //get flags
  url:= flag.String("url","","Atack Domain Name.")
  passwd:= flag.String("word","","Word List Path.")
  redirect:= flag.Bool("r",false,"redirect mod")
  flag.Parse()
  if *url == "" || *passwd == ""{
    return
  }
  
  req:= Web{
    url:*url,
    path: *passwd,
  }
  
  log("\033[0;32m─────────────────────")
  checkup:= req.CheckFile()
  if !checkup{
    return
  }
  
  checkurl:= req.CheckUrl()
  if !checkurl{
    return
  }
  log("\033[0;32m─────────────────────")
  log("\033[0m")
  Start(*url,*passwd,*redirect)
  return
  
}

func log(message string){
  fmt.Println(message)
}
  
func logo(){
  as:= figlet4go.NewAsciiRender()
  option:= figlet4go.NewRenderOptions()
  option.FontName = "3-d"
  fontData,err := FontFS.ReadFile("assets/3-d.flf")
  if err != nil{
    fmt.Println(err)
  }
  as.LoadBindataFont(fontData,"3-d")
  c1,_:= figlet4go.NewTrueColorFromHexString("11FEBE")
  option.FontColor = []figlet4go.Color{
    c1,
  }
  data,err:= as.RenderOpts("WEB FOUND",option)
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(data)
}

func Start(url string,path string,redirect bool){
  word:= make(chan string)
  out:= make(chan requests.OutPut)
  go readfile.Read(word,path)
  wg:= sync.WaitGroup{}
  go func(){
    for urlPath:= range word{
      if urlPath != ""{
        wg.Add(1)
        go requests.Send(out,url+"/"+urlPath,&wg,redirect)
      }
    }
    
    wg.Wait()
    close(out)
  }()
  var found int64 = 0
  for res:= range out{
    for _,code:= range catch{
      if code == res.Status{
        found++
        fmt.Printf("\033[1;32m[+] Code:\033[0m %d \033[1;32mAddres:\033[0m %s\n",res.Status,res.Url)
        break
      }
    }
  }
  log("\033[0;32m─────────────────────")
  fmt.Printf("\033[0;35m[🤖] Pages:\033[0m %d\n\033[0;35m[🤖] Target:\033[0m %s\n",found,url)
  fmt.Printf("\033[0;35m[🤖] Search:\033[0m Done\n")
  log("\033[0;32m─────────────────────")
}