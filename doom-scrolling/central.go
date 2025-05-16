package main

//
//import (
//	"fmt"
//	"github.com/couchbase/gocb/v2"
//	"log"
//	"time"
//)
//
//func main() {
//	connStr := "db2.lan"
//	username := "jaba_admin"
//	password := "jaba_pwd"
//	gocb.SetLogger(gocb.VerboseStdioLogger())
//	opts := gocb.ClusterOptions{
//		Authenticator: gocb.PasswordAuthenticator{
//			Username: username,
//			Password: password,
//		},
//	}
//
//	cluster, err := gocb.Connect(connStr, opts)
//	if err != nil {
//		log.Fatalf("Ошибка подключения к кластеру: %v", err)
//	}
//	err = cluster.WaitUntilReady(5*time.Second, nil)
//	if err != nil {
//		log.Fatalf("Кластер не готов: %v", err)
//	}
//
//	fmt.Println("Успешно подключились к Couchbase!")
//}
