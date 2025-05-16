package neof4j

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"log"
	"log/slog"
	"rshd/lab1/v2/config"
)

type Neo4jDb struct {
	log    *slog.Logger
	driver neo4j.DriverWithContext
}

func New(log *slog.Logger, cfg config.Config) (*Neo4jDb, error) {
	neo4jUsername := "neo4j"
	neo4jPassword := "your_password"
	uri := "neo4j://localhost:7687"
	driver, err := neo4j.NewDriverWithContext(uri, neo4j.BasicAuth(neo4jUsername, neo4jPassword, ""))
	if err != nil {
		return nil, err
	}
	return &Neo4jDb{log: log, driver: driver}, nil
}

func (db *Neo4jDb) Close() {
	_ = db.driver.Close(context.Background())
}

func (db *Neo4jDb) CreateUser(ctx context.Context, username string) error {
	session := db.driver.NewSession(ctx, neo4j.SessionConfig{})
	defer func(session neo4j.SessionWithContext, ctx context.Context) {
		_ = session.Close(ctx)
	}(session, ctx)

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		return tx.Run(ctx,
			"MERGE (u:User {username: $username})",
			map[string]interface{}{"username": username},
		)
	})
	return err
}

func (db *Neo4jDb) FollowUser(ctx context.Context, followerName, followingName string) error {
	session := db.driver.NewSession(ctx, neo4j.SessionConfig{})
	defer func(session neo4j.SessionWithContext, ctx context.Context) {
		_ = session.Close(ctx)
	}(session, ctx)

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		return tx.Run(ctx, `
			MATCH (a:User {username: $follower}), (b:User {username: $following})
			MERGE (a)-[:FOLLOWS]->(b)`,
			map[string]interface{}{"follower": followerName, "following": followingName},
		)
	})
	return err
}

func (db *Neo4jDb) LikePost(ctx context.Context, username, postID string) error {
	session := db.driver.NewSession(ctx, neo4j.SessionConfig{})
	defer func(session neo4j.SessionWithContext, ctx context.Context) {
		_ = session.Close(ctx)
	}(session, ctx)

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		return tx.Run(ctx, `
			MATCH (u:User {username: $username}), (p:Post {id: $postID})
			MERGE (u)-[:LIKES]->(p)`,
			map[string]interface{}{"username": username, "postID": postID},
		)
	})
	return err
}

func (db *Neo4jDb) GetFeed(ctx context.Context, username string) ([]string, error) {
	session := db.driver.NewSession(ctx, neo4j.SessionConfig{})
	defer func(session neo4j.SessionWithContext, ctx context.Context) {
		err := session.Close(ctx)
		if err != nil {
			log.Print("[INFO] fail to close session")
		}
	}(session, ctx)

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
	defer func(session neo4j.SessionWithContext, ctx context.Context) {
		err := session.Close(ctx)
		if err != nil {
			log.Print("[INFO] fail to close session")
		}
	}(session, ctx)

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		return tx.Run(ctx, `
			MATCH (u:User {username: $username})
			CREATE (p:Post {id: $postID, timestamp: timestamp()})
			CREATE (u)-[:POSTED]->(p)`,
			map[string]interface{}{"postID": postID, "username": username},
		)
	})
	return err
}
