package mongodb

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
	"yes-blog/internal/model/post"
	"yes-blog/internal/model/user"
)

type PostMongoDriver struct {
	collection *mongo.Collection
}

const PostInsertTimeOut = 50000
const PostDeleteTimeOut = 50000
const PostUpdateTimeOut = 50000
const PostGetTimeOut = 50000
const PostGetAllTimeOut = 50000

const UserNotFound = "there is no User "
const PostNotFound = "there is no post @"
const UserNotAllowed = "permission denied"

func (p PostMongoDriver) Insert(pst *post.Post) error {
	ctx, cancel := context.WithTimeout(context.Background(), PostInsertTimeOut*time.Millisecond)
	defer cancel()

	target := bson.M{
		"name": pst.Author,
	}
	var tempUser *user.User
	findErr := p.collection.FindOne(ctx, target).Decode(&tempUser)
	if findErr != nil {
		return errors.New(UserNotFound + pst.Author)
	}

	change := bson.M{
		"$push": bson.M{
			"posts": pst,
		},
	}
	_, err := p.collection.UpdateOne(ctx, target, change)
	if err != nil {
		return errors.New(PostNotFound + pst.ID)
	}
	return nil
}

func (p PostMongoDriver) Delete(postID string, authorName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), PostDeleteTimeOut*time.Millisecond)
	defer cancel()

	target := bson.M{
		"name": authorName,
	}
	var tempUser *user.User
	findErr := p.collection.FindOne(ctx, target).Decode(&tempUser)
	if findErr != nil {
		return errors.New(UserNotFound + authorName)
	}

	if _, _, fr := post.Find(tempUser.Posts, postID); !fr {
		return errors.New(UserNotAllowed)
	}

	change := bson.M{
		"$pull": bson.M{
			"posts": bson.M{
				"id": postID,
			},
		},
	}
	_, err := p.collection.UpdateOne(ctx, target, change)
	return err
}

func (p PostMongoDriver) Update(pst *post.Post) error {
	ctx, cancel := context.WithTimeout(context.Background(), PostUpdateTimeOut*time.Millisecond)
	defer cancel()

	target := bson.M{
		"name":     pst.Author,
		"posts.id": pst.ID,
	}
	var tempUser *user.User
	findErr := p.collection.FindOne(ctx, target).Decode(&tempUser)
	if findErr != nil {
		return errors.New(PostNotFound + pst.ID)
	}

	change := bson.M{
		"$set": bson.M{
			"posts.$.title":     pst.Title,
			"posts.$.body":      pst.Body,
			"posts.$.timeStamp": pst.TimeStamp,
		},
	}
	_, err := p.collection.UpdateOne(ctx, target, change)
	return err
}

func (p PostMongoDriver) Get(postID string) (*post.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), PostGetTimeOut*time.Millisecond)
	defer cancel()

	//todo Better implementation with mongoDB built in filters
	curr, err := p.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer curr.Close(ctx)
	for curr.Next(context.Background()) {
		var usrtmp user.User
		_ = curr.Decode(&usrtmp)
		p, _, rf := post.Find(usrtmp.Posts, postID)
		if rf {
			return p, nil
		}
	}
	return nil, errors.New(PostNotFound + postID)
}

func (p PostMongoDriver) GetAll(startID string, amount int) ([]*post.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), PostGetTimeOut*time.Millisecond)
	defer cancel()

	//todo Better implementation with mongoDB built in filters
	var allPosts []*post.Post
	curr, err := p.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer curr.Close(ctx)
	for curr.Next(context.Background()) {
		var usrtmp user.User
		_ = curr.Decode(&usrtmp)
		allPosts = append(allPosts, usrtmp.Posts...)
	}
	post.Sort(allPosts)
	_, i, fr := post.Find(allPosts, startID)
	if !fr {
		return nil, errors.New(PostNotFound + startID)
	}
	if i-amount+1 < 0 {
		return allPosts[:i+1], nil
	}
	return allPosts[i-amount+1 : i+1], nil
}

func (p PostMongoDriver) GetByUser(userName string) ([]*post.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), PostGetAllTimeOut*time.Millisecond)
	defer cancel()

	target := bson.M{
		"name": userName,
	}
	var res user.User
	err := p.collection.FindOne(ctx, target).Decode(&res)
	if err != nil {
		return nil, errors.New(UserNotFound + userName)
	}
	post.Sort(res.Posts)
	return res.Posts, nil
}

func NewPostMongoDriver(db, collection string) *PostMongoDriver {
	return &PostMongoDriver{
		collection: client.Database(db).Collection(collection),
	}
}
