package api

import (
    "github.com/pquerna/otp/totp"
    "github.com/gin-gonic/gin"
)

func GenerateMFA(c *gin.Context) {
    key, _ := totp.Generate(totp.GenerateOpts{
        Issuer:      "TradingBotApp",
        AccountName: "user@example.com",
    })
    c.JSON(200, gin.H{
        "otpURL": key.URL(), // URL à scanner dans Google Authenticator
    })
}
