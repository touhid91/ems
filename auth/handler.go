package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/touhid91/ems/common"
	"net/http"
	"context"
	"github.com/mongodb/mongo-go-driver/bson"
	"log"
)

func TokenHandler(c *gin.Context) {
	db, edb := common.Db(c.GetHeader("Referer"))

	if nil != edb {
		c.JSON(http.StatusBadRequest, gin.H{"err": edb.Error()})
	}

	cur, _ := db.Collection("user").Find(context.Background(), nil)
	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		elem := bson.NewDocument()
		err := cur.Decode(elem)
		if err != nil { log.Fatal(err) }
		// do something with elem....
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
}
