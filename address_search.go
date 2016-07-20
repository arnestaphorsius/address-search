package main

import (
    "github.com/gin-gonic/gin"
    . "github.com/arnestaphorsius/address-search/controllers"
)

func main() {

    router := gin.Default()

    api := router.Group("api")
    {

        api.GET("/straat", GetAllStreetNames)
        api.GET("/straat/:straatnaam", GetStreetSuggestions)
        api.GET("/straat/:straatnaam/nummer/:huisnummer", GetAddressSpecifics)

        api.GET("/straat/:straatnaam/nummers", GetHouseNumbers)

    }

    router.Run(":8020")
}