package main

import (
	"context"
	"fmt"
	"path/filepath"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"crypto/sha1"

	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
	"github.com/minio/minio-go"
)

const maxFileSize = 5048576 // 1MB

func main() {
	mux := gin.Default()

	mux.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          60 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))
	//config.AllowAllOrigins = true

	mux.POST("/", imgAdd)

	srv := &http.Server{
		Addr:    ":8081",
		Handler: mux,
	}

	gracefulStop := make(chan os.Signal)
	signal.Notify(gracefulStop, os.Interrupt)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("Server started: %s\n", err)
		}
	}()

	// This part will be executed only after receiving SIGINT
	<-gracefulStop
	log.Printf("Shutting down the server...\n")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server graceful shutdown error:", err)
	}
	log.Printf("Server exited")
}

func imgAdd(ctx *gin.Context) {
	//	form, _ := ctx.MultipartForm()
	file, header, err := ctx.Request.FormFile("file")

	if err != nil {
		message := map[string]interface{}{}
		message["status"] = err
		ctx.JSON(403, message)
		return
	}
	name := header.Filename
	extension := filepath.Ext(name)
	//fmt.Println("file path %s", x)
	// fmt.Println()
	if header.Size > maxFileSize {
		message := map[string]interface{}{}
		message["status"] = "Invalid Size"
		ctx.JSON(403, message)
		return
	}

	var (
		contentType = header.Header.Get("Content-Type")
		bucket      = "climbmentors"
	)

	if contentType == "application/octet-stream" {
		contentType = "audio/mpeg"
	}
	// extension := ".jpg"
	// if strings.Contains(name, "png") {
	// 	extension = ".png"
	// 	contentType = "image/png"
	// }

	sum := sha1.Sum([]byte(header.Filename + time.Now().String()))
	shaStr := fmt.Sprintf("%x", sum)

	go func() {

		ssl := true

		s3Client, err := minio.New("s3.amazonaws.com", "AKIAJUAMUZHR3G3U3NXQ", "+Egvp4LBhBCyZmISNqSm3Wd/HJi+nw+ubIkUnjg8", ssl)
		if err != nil {
			fmt.Println(err)
			return
		}

		_, err = s3Client.PutObject(bucket, shaStr+extension, file, -1, minio.PutObjectOptions{ContentType: contentType})
		if err != nil {
			fmt.Println(err)
			return
		}
	}()

	message := map[string]interface{}{}
	message["status"] = shaStr + extension
	ctx.JSON(200, message)
}
