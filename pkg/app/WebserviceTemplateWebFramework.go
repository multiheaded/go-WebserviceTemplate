package app

import (
	"embed"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"io/fs"
	"net/http"
	"os"
	"strings"
	"time"
)

// Filesystem embedding:
// - e.g. if your backend (this software right here) ships the frontend as well as provides API access

//go:embed ui/*
var uiFiles embed.FS

func uiFS() http.FileSystem {
	sub, err := fs.Sub(uiFiles, "ui")

	if err != nil {
		panic(err)
	}

	return http.FS(sub)
}

func initWebFramework() (*gin.Engine, error) {
	// initialize go-gin
	r := gin.Default()

	// CORS might be a problem during development and debugging. Therefore, provide
	// an switch to make this webservice unsafe by granting generous permissions
	// CORS for https://foo.com and https://github.com origins, allowing:
	// - PUT and PATCH methods
	// - Origin header
	// - Credentials share
	// - Preflight requests cached for 12 hours
	unsafeCors := strings.TrimSpace(os.Getenv("WEBSERVICETEMPLATE_CORS_UNSAFE"))
	if len(unsafeCors) != 0 {
		fmt.Println("Using unsafe CORS settings for debug purposes")
		r.Use(cors.New(cors.Config{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "OPTIONS", "DELETE"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}))
	}

	// serve the user interface as static filesystem
	r.StaticFS("/ui", uiFS())

	return r, nil
}
