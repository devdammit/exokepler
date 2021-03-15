package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	addr := flag.String("addr", "192.168.1.14", "Сетевой адрес веб сервера megad")
	pwd := flag.String("password", "sec", "Пароль доступка к megad")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	infoLog.Println("Creating backup megad")
	infoLog.Println(*addr)
	infoLog.Println(*pwd)
}
