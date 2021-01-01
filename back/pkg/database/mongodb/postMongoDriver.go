package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"strconv"
	"time"
	"yes-blog/internal/model/post"
	"yes-blog/internal/model/user"
	"yes-blog/pkg/database/DBException"
)

type PostMongoDriver struct {
	collection *mongo.Collection
}

const PostInsertTimeOut = 50000
const PostDeleteTimeOut = 50000
const PostUpdateTimeOut = 50000
const PostGetTimeOut = 50000
const PostGetAllTimeOut = 50000

func (p PostMongoDriver) Insert(pst *post.Post) (*user.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), PostInsertTimeOut*time.Millisecond)
	defer cancel()

	target := bson.M{
		"name": pst.AuthorID,
	}
	var usr user.User
	findErr := p.collection.FindOne(ctx, target).Decode(&usr)
	if findErr != nil {
		return nil, DBException.ThrowUserNotFoundException(pst.AuthorID)
	}

	change := bson.M{
		"$push": bson.M{
			"posts": pst,
		},
	}
	_, err := p.collection.UpdateOne(ctx, target, change)
	if err != nil {
		return nil, DBException.ThrowInternalDBException(err.Error())
	}
	return &usr, nil
}

func (p PostMongoDriver) Delete(postID string, authorName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), PostDeleteTimeOut*time.Millisecond)
	defer cancel()

	target := bson.M{
		"name": authorName,
	}
	var usr user.User
	findErr := p.collection.FindOne(ctx, target).Decode(&usr)
	if findErr != nil {
		return DBException.ThrowUserNotFoundException(authorName)
	}

	if _, _, fr := post.Find(usr.Posts, postID); !fr {
		return DBException.ThrowUserNotAllowedException(authorName)
	}

	change := bson.M{
		"$pull": bson.M{
			"posts": bson.M{
				"id": postID,
			},
		},
	}
	_, err := p.collection.UpdateOne(ctx, target, change)
	if err != nil {
		return DBException.ThrowInternalDBException(err.Error())
	}
	return nil
}

func (p PostMongoDriver) Update(pst *post.Post) error {
	ctx, cancel := context.WithTimeout(context.Background(), PostUpdateTimeOut*time.Millisecond)
	defer cancel()

	target := bson.M{
		"name":     pst.AuthorID,
		"posts.id": pst.ID,
	}
	var usr user.User
	findErr := p.collection.FindOne(ctx, target).Decode(&usr)
	if findErr != nil {
		return DBException.ThrowPostNotFoundException(pst.ID)
	}

	change := bson.M{
		"$set": bson.M{
			"posts.$.title":     pst.Title,
			"posts.$.body":      pst.Body,
			"posts.$.timeStamp": pst.TimeStamp,
		},
	}
	_, err := p.collection.UpdateOne(ctx, target, change)
	if err != nil {
		return DBException.ThrowInternalDBException(err.Error())
	}
	return nil
}

func (p PostMongoDriver) Get(postID string) (*post.Post, *user.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), PostGetTimeOut*time.Millisecond)
	defer cancel()

	//todo Better implementation with mongoDB built in filters
	curr, err := p.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, nil, DBException.ThrowInternalDBException(err.Error())
	}
	defer curr.Close(ctx)
	for curr.Next(context.Background()) {
		var usr user.User
		_ = curr.Decode(&usr)
		p, _, rf := post.Find(usr.Posts, postID)
		if rf {
			return p, &usr, nil
		}
	}
	return nil, nil, DBException.ThrowPostNotFoundException(postID)
}

func (p PostMongoDriver) GetAll(startIndex, amount int) ([]*post.Post, []*user.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), PostGetTimeOut*time.Millisecond)
	defer cancel()

	//todo Better implementation with mongoDB built in filters
	var allPosts []*post.Post
	curr, err := p.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, nil, DBException.ThrowInternalDBException(err.Error())
	}
	defer curr.Close(ctx)
	usrs := make(map[*post.Post]*user.User)
	for curr.Next(context.Background()) {
		var usr user.User
		_ = curr.Decode(&usr)
		allPosts = append(allPosts, usr.Posts...)
		for _, p := range allPosts {
			usrs[p] = &usr
		}
	}

	if startIndex >= len(allPosts) || startIndex < 0 {
		return nil, nil, DBException.ThrowPostNotFoundException("(index of " + strconv.Itoa(startIndex) + ")")
	}

	post.Sort(allPosts)
	resUs := make([]*user.User, amount)
	sIn := startIndex
	eIn := startIndex - amount + 1
	if startIndex-amount+1 < 0 {
		eIn = 0
	}
	for j := sIn; j >= eIn; j-- {
		resUs = append(resUs, usrs[allPosts[j]])
	}
	return allPosts[eIn : sIn+1], resUs, nil
}

func (p PostMongoDriver) GetByUser(userName string) ([]*post.Post, *user.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), PostGetAllTimeOut*time.Millisecond)
	defer cancel()

	target := bson.M{
		"name": userName,
	}
	var usr user.User
	err := p.collection.FindOne(ctx, target).Decode(&usr)
	if err != nil {
		return nil, nil, DBException.ThrowUserNotFoundException(userName)
	}
	post.Sort(usr.Posts)
	return usr.Posts, &usr, nil
}

func NewPostMongoDriver(db, collection string) *PostMongoDriver {
	return &PostMongoDriver{
		collection: client.Database(db).Collection(collection),
	}
}
