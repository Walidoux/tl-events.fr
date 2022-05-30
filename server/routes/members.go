package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tigrouland/eventsstats/server/mongo"
	"github.com/tigrouland/eventsstats/server/mongo/entities"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"time"
)

func Members(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := mongo.Get().Collection("players").Find(ctx, bson.D{})
	if err != nil {
		log.Fatal("an error occurred while retrieving the players' data: ", err)
	}
	var players []entities.Player
	for cursor.Next(ctx) {
		var player entities.Player
		err = cursor.Decode(&player)
		if err != nil {
			log.Fatal(err)
		}
		var playerUUID uuid.UUID
		if !player.UUID.IsZero() {
			playerUUID, err = uuid.FromBytes(player.UUID.Data)
			if err != nil {
				log.Fatal(err)
			}
			player.DecodedUUID = playerUUID
		}
		players = append(players, player)
	}
	c.JSON(200, players)
}
