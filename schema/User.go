package schema

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User Role as Enum
type Role string

const (
	READER Role = "READER"
	ARTIST Role = "ARTIST"
)

/*

Create the schema that is used to interact with the database
Notice the BSON here.

*/

// User of the platform, Reader who reads or Artist who publishes
type User struct {
	ID        primitive.ObjectID   `bson:"_id"`
	FirstName string               `bson:"firstName"`
	LastName  string               `bson:"lastName"`
	Username  string               `bson:"username"`
	Email     string               `bson:"email"`
	Role      Role                 `bson:"role"`
	Bio       string               `bson:"bio"`
	ImageUri  string               `bson:"imageUri"`
	CreatedAt primitive.DateTime   `bson:"createdAt"`
	Artworks  []primitive.ObjectID `bson:"artworks"`
}
