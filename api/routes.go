package api

import (
    "github.com/gin-gonic/gin"
    "github.com/hashicorp/vault/api"
    "fmt"

)

func CreateKeyRoute(r *gin.Engine, client *api.Client, keyName string) {
    if keyName == "" {
        fmt.Errorf("Missing KeyName")
    }
    r.POST("/create-key", CreateTransitKeyHandler(client, keyName))
}

func encryptDataRoute(r *gin.Engine, a vaultServer) {

}