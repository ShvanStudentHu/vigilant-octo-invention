package api

import (
    "github.com/gin-gonic/gin"
    "github.com/hashicorp/vault/api"


)

func CreateKeyRoute(r *gin.Engine, client *api.Client, keyName string) {
    r.POST("/create-key", CreateTransitKeyHandler(client, keyName))
}
