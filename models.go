package main

import "time"

type Comic struct {
	ID          string    `bson:"_id,omitempty" json:"id"`
	Title       string    `bson:"title" json:"title"`
	Image       string    `bson:"image" json:"image"`
	Desc        string    `bson:"desc" json:"desc"`
	Type        string    `bson:"type" json:"type"`
	Endpoint    string    `bson:"endpoint" json:"endpoint"`
	Thumbnail   string    `bson:"thumbnail" json:"thumbnail"`
	Author      string    `bson:"author" json:"author"`
	Status      string    `bson:"status" json:"status"`
	Rating      string    `bson:"rating" json:"rating"`
	Genre       []string  `bson:"genre" json:"genre"`
	ChapterList []Chapter `bson:"chapter_list" json:"chapter_list"`
	CreatedAt   time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at" json:"updated_at"`
}

type Chapter struct {
	Name     string `bson:"name" json:"name"`
	Endpoint string `bson:"endpoint" json:"endpoint"`
}

type ChapterDetail struct {
	Title  string   `bson:"title" json:"title"`
	Images []string `bson:"images" json:"images"`
}
