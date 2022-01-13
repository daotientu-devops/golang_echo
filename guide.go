package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"io"
	"net/http"
	"os"
)

func main() {
	// Echo instance
	e := echo.New()
	// Root level middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// Group level middleware
	g := e.Group("/admin")
	g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "admin" && password == "admin" {
			return true, nil
		}
		return false, nil
	}))
	// Route level middleware
	track := func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			println("request to /users")
			return next(c)
		}
	}
	e.GET("/users", func(c echo.Context) error {
		return c.String(http.StatusOK, "/users")
	}, track)
	// Routes
	e.GET("/", hello)
	// Path Parameters
	// e.GET("/users/:id", getUser)
	e.GET("/users/:id", getUser)
	// Query Parameters
	// /show?team=x-men&member=wolverine
	e.GET("/show", showUser)
	// Form application/x-www-form-urlencoded
	e.POST("/save", saveUser)
	// Form multipart/form-data
	e.POST("/save/avatar", saveAvatar)
	// Handing request
	// + Bind json, xml, form or query payload into Go struct based on Content-Type request header
	// + Render response as json or xml with status code
	type User struct {
		Name  string `json:"name" xml:"name" form:"name" query:"name"`
		Email string `json:"email" xml:"email" form:"email" query:"name"`
	}
	e.POST("/users", func(c echo.Context) error {
		u := new(User)
		if err := c.Bind(u); err != nil {
			return err
		}
		return c.XML(http.StatusCreated, u)
	})
	// Static Content
	// Serve any file from static directory for path /static/*
	e.Static("/static", "static")
	// Start server at localhost:1323
	e.Logger.Fatal(e.Start(":1323"))
}

func saveAvatar(c echo.Context) error {
	// Get name
	name := c.FormValue("name")
	// Get avatar
	avatar, err := c.FormFile("avatar")
	if err != nil {
		return err
	}
	// Source
	src, err := avatar.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	// Destination
	dst, err := os.Create(avatar.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()
	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}
	return c.HTML(http.StatusOK, "<b>Thank you! "+name+"</b>")
}

func saveUser(c echo.Context) error {
	// Get nam and email
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name: "+name+", email: "+email)
}

func showUser(c echo.Context) error {
	// Get team and member from the query string
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team: "+team+", member: "+member)
}

func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
