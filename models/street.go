package models

type Street struct {

    TableName struct {}																	`sql:"dm_adres_search,alias:street" json:"-"`

    Dm_adres_search_straatnaam string                   `json:"dm_adres_search_straatnaam"`
    Dm_adres_search_huisnummer int64                    `json:"-"`
}
