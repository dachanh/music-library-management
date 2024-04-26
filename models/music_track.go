package models

import (
	"music-library-management/sdk/db"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MusicTrack struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title       string             `bson:"title" json:"title"`
	Artist      string             `bson:"artist" json:"artist"`
	Album       string             `bson:"album" json:"album"`
	Genre       string             `bson:"genre" json:"genre"`
	ReleaseYear int                `bson:"release_year" json:"release_year"`
	Duration    int                `bson:"duration" json:"duration"` // Duration in seconds
	MP3File     string             `bson:"mp3File" json:"mp3_file"`  // URL or path to the file
}

var MusicTrackDB = &db.Instance{
	ColName:        "music_track",
	TemplateObject: &MusicTrack{},
}

func InitMusicTrackModel(s *mongo.Database) {
	MusicTrackDB.ApplyDatabase(s)
}
