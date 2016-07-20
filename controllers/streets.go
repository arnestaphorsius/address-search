package controllers

import (
    "github.com/gin-gonic/gin"
    "github.com/arnestaphorsius/liveop-go/db"
    "github.com/arnestaphorsius/address-search/models"
    "fmt"
)

func GetAllStreetNames(c *gin.Context) {

    conn := db.Connection()

    streets := []models.Street{}

    if err := conn.Model(&streets).
        ColumnExpr("distinct(dm_adres_search_straatnaam)").
        Order("dm_adres_search_straatnaam").
        Select(); err != nil {

        NotFound(c)
        fmt.Println(err)

    } else {

        c.IndentedJSON(200, gin.H{"result": toStringArray(streets)})
    }
}

func GetStreetSuggestions(c *gin.Context) {

    straatnaam := c.Param("straatnaam")

    conn := db.Connection()

    streets := []models.Street{}

    if err := conn.Model(&streets).
        ColumnExpr("distinct(dm_adres_search_straatnaam)").
        Where("dm_adres_search_straatnaam ILIKE ?", straatnaam + "%").
        Order("dm_adres_search_straatnaam").
        Limit(10).
        Select(); err != nil {

        NotFound(c)
        fmt.Println(err)

    } else {

        c.IndentedJSON(200, gin.H{"result": toStringArray(streets)})
    }
}

func GetHouseNumbers(c *gin.Context) {

    straatnaam := c.Param("straatnaam")

    conn := db.Connection()

    streets := []models.Street{}

    if err := conn.Model(&streets).
        ColumnExpr("dm_adres_search_straatnaam, dm_adres_search_huisnummer").
        Where("dm_adres_search_straatnaam ILIKE ?", straatnaam + "%").
        Group("dm_adres_search_straatnaam, dm_adres_search_huisnummer").
        Order("dm_adres_search_huisnummer").
        Select(); err != nil {

        NotFound(c)
        fmt.Println(err)

    } else {

        results := make(map[string][]int64)

        for _, v := range streets {
            results[v.Dm_adres_search_straatnaam] =
                append(results[v.Dm_adres_search_straatnaam], v.Dm_adres_search_huisnummer)
        }

        c.IndentedJSON(200, gin.H{"results": results})
    }
}

func toStringArray(s []models.Street) (results []string) {

    for _, v := range s {
        results = append(results, v.Dm_adres_search_straatnaam)
    }

    return results
}