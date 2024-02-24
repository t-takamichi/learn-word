package collections

import "time"

type Word struct {
	ID         string    `bson:"_id"`
	Vocabulary string    `bson:"vocabulary"`
	Meaning    string    `bson:"meaning"`
	CreatedAt  time.Time `bson:"createdAt"`
	UpdatedAt  time.Time `bson:"updatedAt"`
}
