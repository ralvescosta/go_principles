package repositories

import (
	"context"
	"goapi/src/entities"
	"log"

	"cloud.google.com/go/firestore"
)

type PostRepository interface {
	Save(post *entities.Post) (*entities.Post, error)
	FindAll() ([]entities.Post, error)
}

type repository struct{}

const (
	projectId      string = "id"
	collectionName string = "posts"
)

func PostRepositoryConstructor() PostRepository {
	return &repository{}
}

func (*repository) Save(post *entities.Post) (*entities.Post, error) {
	ctx := context.Background()

	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatal("Failed to create a Firestore Client")
		return nil, err
	}
	defer client.Close()

	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"id":    post.Id,
		"title": post.Title,
		"text":  post.Text,
	})
	if err != nil {
		log.Fatal("Failed to Add on Firestore")
		return nil, err
	}

	return post, nil
}

func (*repository) FindAll() ([]entities.Post, error) {
	ctx := context.Background()

	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatal("Failed to create a Firestore Client")
		return nil, err
	}
	defer client.Close()

	var posts []entities.Post
	interator := client.Collection(collectionName).Documents(ctx)

	for {
		doc, err := interator.Next()
		if err != nil {
			log.Fatal("Failed to interate Post List")
			return nil, err
		}
		post := entities.Post{
			Id:    doc.Data()["id"].(int64),
			Title: doc.Data()["title"].(string),
			Text:  doc.Data()["text"].(string),
		}

		posts = append(posts, post)
	}

	return posts, nil
}
