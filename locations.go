package main

// import libraries
import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
)

// Structure Model Province
type LocationProvince struct {
	ID             int    `json:"province_id"`
	Code           string `json:"province_code"`
	Name           string `json:"province_name"`
	Order          int    `json:"province_order"`
	Status         string `json:"province_status"`
	CreatedBy      string `json:"created_by"`
	CreatedOn      string `json:"created_on"`
	LastModifiedBy string `json:"last_modified_by"`
	LastModifiedOn string `json:"last_modified_on"`
	isDeleted      int    `json:"is_deleted"`
}

// Structure Model District
type LocationDistrict struct {
	ID             int    `json:"district_id"`
	ProvinceId     int    `json:"province_id"`
	Code           string `json:"district_code"`
	Name           string `json:"district_name"`
	Order          int    `json:"district_order"`
	Status         string `json:"district_status"`
	CreatedBy      string `json:"created_by"`
	CreatedOn      string `json:"created_on"`
	LastModifiedBy string `json:"last_modified_by"`
	LastModifiedOn string `json:"last_modified_on"`
	isDeleted      int    `json:"is_deleted"`
}

// Structure Model SubDistrict
type LocationSubDistrict struct {
	ID             int    `json:"sub_district_id"`
	DistrictId     int    `json:"district_id"`
	Code           string `json:"sub_district_code"`
	Name           string `json:"sub_district_name"`
	Order          int    `json:"sub_district_order"`
	Status         string `json:"sub_district_status"`
	CreatedBy      string `json:"created_by"`
	CreatedOn      string `json:"created_on"`
	LastModifiedBy string `json:"last_modified_by"`
	LastModifiedOn string `json:"last_modified_on"`
	isDeleted      int    `json:"is_deleted"`
}

// Main function
func main() {
	// Echo instance
	server := echo.New()
	fmt.Println("Location provinces")
	// Routes
	server.GET("/provinces", getListProvinces)
	server.GET("/districts/:provinceId", getListDistrictsByProvinceId)
	server.GET("/subdistricts/:districtId", getListSubDistrictsByDistrictId)
	// Start server at localhost:1323
	server.Logger.Fatal(server.Start(":1323"))
}

// Function get list provinces
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

// Function get list districts by province ID
func getListDistrictsByProvinceId(c echo.Context) error {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/n2cms")
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	provinceId, _ := strconv.Atoi(c.Param("provinceId"))
	result, err := db.Query("SELECT * FROM location_district WHERE province_id=?", provinceId)
	defer result.Close()
	if err != nil {
		log.Fatal(err)
	}
	var listDistricts []LocationDistrict
	for result.Next() {
		var d LocationDistrict
		_ = result.Scan(&d.ID, &d.ProvinceId, &d.Code, &d.Name, &d.Order, &d.Status, &d.CreatedBy, &d.CreatedOn, &d.LastModifiedBy, &d.LastModifiedOn, &d.isDeleted)
		listDistricts = append(listDistricts, d)
	}
	return c.JSON(http.StatusOK, listDistricts)
}

// Function get list subdistricts by district ID
func getListSubDistrictsByDistrictId(c echo.Context) error {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/n2cms")
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	districtId, _ := strconv.Atoi(c.Param("districtId"))
	result, err := db.Query("SELECT * FROM location_sub_district WHERE district_id=?", districtId)
	defer result.Close()
	if err != nil {
		log.Fatal(err)
	}
	var listSubDistricts []LocationSubDistrict
	for result.Next() {
		var s LocationSubDistrict
		_ = result.Scan(&s.ID, &s.DistrictId, &s.Code, &s.Name, &s.Order, &s.Status, &s.CreatedBy, &s.CreatedOn, &s.LastModifiedBy, &s.LastModifiedOn, &s.isDeleted)
		listSubDistricts = append(listSubDistricts, s)
	}
	return c.JSON(http.StatusOK, listSubDistricts)
}
