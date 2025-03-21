package document

import (
	"github.com/couchbase/gocb/v2"
	"log/slog"
	"rshd/lab1/v2/config"
	"time"
)

// InitializeCluster возвращает пропингованный кластер
func InitializeCluster(cfg config.Config, log *slog.Logger) (*gocb.Cluster, error) {
	log.Info("Initializing Couchbase cluster")
	//gocb.SetLogger(gocb.VerboseStdioLogger()) // закоментить, если нужно заткнуть логи бд

	options := gocb.ClusterOptions{
		Authenticator: gocb.PasswordAuthenticator{
			Username: cfg.CouchBaseCfg.Username,
			Password: cfg.CouchBaseCfg.Password,
		},
	}

	if err := options.ApplyProfile(gocb.ClusterConfigProfileWanDevelopment); err != nil {
		log.Error("Failed to apply Couchbase configuration profile", "error", err)
		return nil, err
	}

	cluster, err := gocb.Connect("couchbase://"+cfg.CouchBaseCfg.URL, options)
	if err != nil {
		log.Error("Failed to connect to Couchbase cluster", "error", err)
		return nil, err
	}

	bucket := cluster.Bucket(cfg.CouchBaseCfg.Bucket)
	//
	err = bucket.WaitUntilReady(5*time.Second, nil)
	if err != nil {
		log.Error("Failed to wait Couchbase bucket ready", "error", err)
		return nil, err
	}

	//_ = bucket.Scope("test-scope").Collection("users")

	//// Create and store a Document
	//type User struct {
	//	Name      string   `json:"name"`
	//	Email     string   `json:"email"`
	//	Interests []string `json:"interests"`
	//}
	//
	//_, err = col.Upsert("jade",
	//	User{
	//		Name:      "Jade",
	//		Email:     "jaba@test-email.com",
	//		Interests: []string{"Swimming", "Rowing"},
	//	}, nil)
	//if err != nil {
	//	log.Error("Failed to update Couchbase user", "error", err)
	//	panic(err)
	//}
	//
	//// Get the document back
	//getResult, err := col.Get("u:jade", nil)
	//if err != nil {
	//	panic(err)
	//}
	////
	//var inUser User
	//err = getResult.Content(&inUser)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("User: %v\n", inUser)
	//
	//inventoryScope := bucket.Scope("inventory")
	//queryResult, err := inventoryScope.Query(
	//	fmt.Sprintf("SELECT * FROM airline WHERE id=10"),
	//	&gocb.QueryOptions{},
	//)
	//if err != nil {
	//	panic(err)
	//}
	//
	//// Print each found Row
	//for queryResult.Next() {
	//	var result interface{}
	//	err := queryResult.Row(&result)
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Println(result)
	//}
	//
	//if err := queryResult.Err(); err != nil {
	//	panic(err)
	//}
	return cluster, nil
}
