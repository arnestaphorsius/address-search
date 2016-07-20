package controllers

import (
    "github.com/gin-gonic/gin"
    "github.com/arnestaphorsius/address-search/db"
    "github.com/arnestaphorsius/address-search/models"
    "fmt"
)

func GetAddressSpecifics(c *gin.Context) {

    straatnaam := c.Param("straatnaam")
    huisnummer := c.Param("huisnummer")

    conn := db.Connection()

    addresses := []models.Address{}

    if err := conn.Model(&addresses).
        ColumnExpr("distinct(dm_adres_search_bag), *").
        Where("dm_adres_search_straatnaam=?", straatnaam).
        Where("dm_adres_search_huisnummer=?", huisnummer).
        Select(); err != nil {

        NotFound(c)
        fmt.Println(err)

    } else {
        c.IndentedJSON(200, addresses)
    }
}