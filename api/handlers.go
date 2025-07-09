package api 

import(
    "github.com/gin-gonic/gin"
	"github.com/ShvanStudentHu/vigilant-octo-invention/vault"
	"github.com/hashicorp/vault/api"

)

func CreateTransitKeyHandler(client *api.Client, keyName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := vault.CreateTransitKey(client, keyName)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "Encryption Key generated"})
	}
}