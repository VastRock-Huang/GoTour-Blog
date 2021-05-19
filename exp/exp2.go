//2-11-2 fsnotify demo
package main

import (
	"github.com/fsnotify/fsnotify"
	"log"
)

func main() {
	watcher,_:=fsnotify.NewWatcher()
	defer watcher.Close()

	done:=make(chan bool)
	go func() {
		for{
			select {
			case event,ok := <-watcher.Events:
				if !ok{
					return
				}
				log.Println("event:",event)
				if event.Op&fsnotify.Write==fsnotify.Write{
					log.Println("modified file:",event.Name)
				}
			case err,ok:=<-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:",err)
			}
		}
	}()
	path:="E:\\Computer\\Go\\gotour-blog\\configs\\config.yaml"
	_ = watcher.Add(path)
	//以下两种方法都可以使主程序阻塞
	<-done	//从管道中读取一个数据,由于管道为空因此会阻塞程序
	//done<-true	//向管道中写入一个数据, 在数据被读取之前一直阻塞程序
}
