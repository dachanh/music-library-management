package model

import (
	"music-library-management/sdk/db"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MusicTrack struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title"`
	Artist      string             `bson:"artist"`
	Album       string             `bson:"album"`
	Genre       string             `bson:"genre"`
	ReleaseYear int                `bson:"releaseYear"`
	Duration    int                `bson:"duration"` // Duration in seconds
	MP3File     string             `bson:"mp3File"`  // URL or path to the file
}

var MusicTrackDB = &db.Instance{
	ColName:        "music_track",
	TemplateObject: &MusicTrack{},
}
