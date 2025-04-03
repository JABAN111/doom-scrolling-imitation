package neof4j // Исправлено название пакета для соответствия common style

import (
	"context"
	"log/slog"
	"rshd/lab1/v2/config"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Neo4jDb struct {
	log    *slog.Logger
	driver neo4j.DriverWithContext
}

func New(log *slog.Logger, cfg config.Config) (*Neo4jDb, error) {
	neo4jUsername := "neo4j"
	neo4jPassword := "password123"
	uri := "neo4j://localhost:7687"
	driver, err := neo4j.NewDriverWithContext(uri, neo4j.BasicAuth(neo4jUsername, neo4jPassword, ""))
	if err != nil {
		return nil, err
	}
	return &Neo4jDb{log: log, driver: driver}, nil
}

func (db *Neo4jDb) Close() {
	db.driver.Close(context.Background())
}

func (db *Neo4jDb) CreateUser(ctx context.Context, username string) error {
	session := db.driver.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		// Используем MERGE и уникальный constraint (должен быть создан в БД)
		return tx.Run(ctx,
			"MERGE (u:User {username: $username})",
			map[string]interface{}{"username": username},
		)
	})
	return err
}

func (db *Neo4jDb) FollowUser(ctx context.Context, followerName, followingName string) error {
	session := db.driver.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		return tx.Run(ctx, `
			MATCH (a:User {username: $follower}), (b:User {username: $following})
			MERGE (a)-[:FOLLOWS]->(b)`, // Проверено название отношения
			map[string]interface{}{"follower": followerName, "following": followingName},
		)
	})
	return err
}

func (db *Neo4jDb) LikePost(ctx context.Context, username, postID string) error { // Переименовано userID -> username
	session := db.driver.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		return tx.Run(ctx, `
			MATCH (u:User {username: $username}), (p:Post {id: $postID})
			MERGE (u)-[:LIKES]->(p)`,
			map[string]interface{}{"username": username, "postID": postID}, // Исправлен параметр
		)
	})
	return err
}

// func (db *Neo4jDb) GetFeed(ctx context.Context, username string) ([]string, error) {
// 	session := db.driver.NewSession(ctx, neo4j.SessionConfig{})
// 	defer session.Close(ctx)

// 	result, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
// 		// Исправлено: используем параметр $username вместо хардкода "alice"
// 		return tx.Run(ctx, `
//             MATCH (u:User {username: $username})-[:FOLLOWS]->(f)-[:POSTED]->(p:Post)
//             RETURN p.id AS postID
//             ORDER BY p.timestamp DESC
//             LIMIT 10`,
// 			map[string]interface{}{"username": username}, // Параметр теперь используется в запросе
// 		)
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Явное приведение типа результата
// 	neo4jResult := result.(neo4j.ResultWithContext)

// 	feed := make([]string, 0)
// 	for neo4jResult.Next(ctx) {
// 		postID, _ := neo4jResult.Record().Get("postID")
// 		feed = append(feed, postID.(string))

// 		// Для отладки
// 		db.log.Debug("Found post", "postID", postID)
// 	}

// 	if err = neo4jResult.Err(); err != nil {
// 		return nil, err
// 	}

// 	// Добавляем проверку, если результат пустой
// 	if len(feed) == 0 {
// 		db.log.Warn("No posts found in feed", "username", username)
// 	}

// 	return feed, nil
// }

func (db *Neo4jDb) GetFeedDeprecated(ctx context.Context, username string) ([]string, error) {
	session := db.driver.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)

	result, err := session.Run(ctx,
		`MATCH (u:User {username: $username})-[:FOLLOWS]->(f)-[:POSTED]->(p:Post) RETURN p ORDER BY p.timestamp DESC`,
		map[string]any{"username": username},
	)
	if err != nil {
		return nil, err
	}

	posts := make([]map[string]interface{}, 0)
	for result.Next(ctx) {
		record := result.Record()
		// fmt.Print("curr record", record)
		nodeVal, found := record.Get("p")
		if !found {
			continue
		}

		// Приводим к neo4j.Node, чтобы получить доступ к свойствам узла.
		node, ok := nodeVal.(neo4j.Node)
		if !ok {
			continue
		}

		// node.Props содержит свойства записи, например: {"id": "post1", "timestamp": 1743637070774}
		posts = append(posts, node.Props)
	}

	if err = result.Err(); err != nil {
		return nil, err
	}
	// var res []string

	// type Post struct {
	// 	id string
	// }
	// for _, val := range posts {
	// 	post := val.(Post)
	// }

	return nil, nil
}
func (db *Neo4jDb) GetFeed(ctx context.Context, username string) ([]string, error) {
	session := db.driver.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)

	result, err := session.Run(ctx,
		`MATCH (u:User {username: $username})-[:FOLLOWS]->(f)-[:POSTED]->(p:Post)
		 RETURN p.id AS postID
		 ORDER BY p.timestamp DESC`,
		map[string]any{"username": username},
	)
	if err != nil {
		return nil, err
	}

	posts := make([]string, 0)
	for result.Next(ctx) {
		record := result.Record()
		val, found := record.Get("postID")
		if !found {
			continue
		}
		id, ok := val.(string)
		if !ok {
			continue
		}
		posts = append(posts, id)
	}

	if err = result.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func (db *Neo4jDb) CreatePost(ctx context.Context, postID, username string) error { // Переименовано userId -> username
	session := db.driver.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		return tx.Run(ctx, `
			MATCH (u:User {username: $username})
			CREATE (p:Post {id: $postID, timestamp: timestamp()})
			CREATE (u)-[:POSTED]->(p)`, // Использован CREATE вместо MERGE
			map[string]interface{}{"postID": postID, "username": username},
		)
	})
	return err
}
