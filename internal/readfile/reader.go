// Copyright 2026 Lasith Ruwantha Amarawansha
// SPDX-License-Identifier: ISC

package readfile

import (
  "os"
  "bufio"
  )
  
func Read(word chan string,path string){
  
  file,_:= os.Open(path)
  
  defer file.Close()
  
  scaner:= bufio.NewScanner(file)
  for scaner.Scan(){
    
    word<-scaner.Text()
  }
  close(word)
}