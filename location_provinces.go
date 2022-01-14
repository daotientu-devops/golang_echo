 package main

 import (
    "github.com/labstack/echo/v4"
    "database/sql"
    "net/http"
    "fmt"
    "log"
    _ "github.com/go-sql-driver/mysql"
 )

 type LocationProvince struct {
    ID int `json:"province_id"`
    Code string `json:"province_code"`
    Name string `json:"province_name"`
    Order int `json:"province_order"`
    Status string `json:"province_status"`
    CreatedBy string `json:"created_by"`
    CreatedOn string `json:"created_on"`
    LastModifiedBy string `json:"last_modified_by"`
    LastModifiedOn string `json:"last_modified_on"`
    isDeleted int `json:"is_deleted"`
 }

 func main() {
    // Echo instance
    server := echo.New()
    fmt.Println("Location provinces")
    // Routes
    server.GET("/provinces", getListProvinces)
    // Start server at localhost:1323
    server.Logger.Fatal(server.Start(":1323"))
 }

 func getListProvinces(c echo.Context) error {
    db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/n2cms")
    defer db.Close()
    if err != nil {
        log.Fatal(err)
    }
    result, err := db.Query("SELECT * FROM location_province")
    defer result.Close()
    if err != nil {
        log.Fatal(err)
    }
    var listProvinces []LocationProvince
    for result.Next() {
        var p LocationProvince
        _ = result.Scan(&p.ID, &p.Code, &p.Name, &p.Order, &p.Status, &p.CreatedBy, &p.CreatedOn, &p.LastModifiedBy, &p.LastModifiedOn, &p.isDeleted)
        listProvinces = append(listProvinces, p)
    }
    return c.JSON(http.StatusOK, listProvinces)
 }