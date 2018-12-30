package repo

import (
	"fmt"
	"os"

	"github.com/globalsign/mgo"
)

var mongoDBURL string

var mongoSession *mgo.Session

// New will create a new Repo instance tied with MongoDB implementation
func New() {
	fmt.Println("Initializing repo with the MongoDB instance ..")
	mongoDBURL = os.Getenv("MONGODB_URL")
	if mongoDBURL == "" {
		panic("MONGODB_URL needs to be defined!! e.g mongodb://<dbuser>:<dbpassword>@ds121183.mlab.com:21183/cooljoe")
	}
	session, err := mgo.Dial(mongoDBURL)
	if err != nil {
		panic(err)
	}
	mongoSession = session
	// id, err := ulid.New(ulid.Now(), ulid.Monotonic(rand.New(rand.NewSource(time.Unix(1000000, 0).UnixNano())), 0))
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("HOWTO use ulid?")
	// fmt.Println("ID: ", id, " <<****>>")
}

// ReadSecrets get data out from the Vautl server ..
func ReadSecrets() {
	fmt.Println("Reading .. secretes .. cooljoe")
	c := mongoSession.DB("cooljoe").C("quotes")
	q, err := c.Count()
	if err != nil {
		panic(err)
	}
	fmt.Println("COUNT", q)
}
