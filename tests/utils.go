package tests

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/dbtest"
	"io/ioutil"
	"math/rand"
	"time"
)

func GetTestDB() (*mgo.Session, *mgo.Database) {
	var server dbtest.DBServer
	tempDir, _ := ioutil.TempDir("", "testing")
	server.SetPath(tempDir)
	session := server.Session()

	db := session.DB("test_db")
	return session, db
}

func FreeTestDB(session *mgo.Session, db *mgo.Database) {
	db.DropDatabase()
	session.Close()
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}