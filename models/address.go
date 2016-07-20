package models

type Address struct {

    TableName struct {}																	`sql:"dm_adres_search,alias:address" json:"-"`

    Dm_adres_search_id string                           `json:"-"`
    Dm_adres_search_straatnaam string                   `json:"dm_adres_search_straatnaam"`
    Dm_adres_search_huisnummer int64                    `json:"dm_adres_search_huisnummer"`
    Dm_adres_search_huisletter string                   `json:"-"`
    Dm_adres_search_huisnummertoevoeging string         `json:"dm_adres_search_huisnummertoevoeging"`
    Dm_adres_search_postcode string                     `json:"dm_adres_search_postcode"`
    Dm_adres_search_gemeentenaam string                 `json:"dm_adres_search_gemeentenaam"`
    Dm_adres_search_woonplaats string                   `json:"-"`
    Dm_adres_search_bag int64                           `json:"dm_adres_search_bag"`
}