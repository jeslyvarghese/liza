package main

import (
	"github.com/jeslyvarghese/liza/src/engine"
	"github.com/jeslyvarghese/liza/src/urlops"
	"log"
	"net/http"
	"net/url"
	"os"
	"encoding/json"
)

type ResponseData struct {
	OptimisedImageURL string `json:"optimisedImageURL,omitempty"`
	OriginalImageURL string `json:"OriginalImageURL,omitempty"`
} 

func main() {
	f, err := os.OpenFile("applog.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
	}
	defer f.Close()
	log.SetOutput(f)
	http.HandleFunc("/resize", resize)
	http.ListenAndServe(":8080", nil)
}

func resize(w http.ResponseWriter, r *http.Request) {
	imageURL, _ :=  url.QueryUnescape(r.URL.Query().Get("imageURL"))
	var response ResponseData
	if optimisedImageURL, hasOptimisedImageURL := engine.CheckHasImage(imageURL); hasOptimisedImageURL {
		log.Println("Sending optimised image")
		response = newResponseData(optimisedImageURL, imageURL)
	} else {
		var downloadImageCallBack urlops.DownloadCallBack
		log.Println("Creating optimised image")
		response = newResponseData(optimisedImageURL, imageURL)
		downloadImageCallBack = func(err error, isSuccess bool, destImagePath string) {
			if isSuccess {
				imagePath, isResized := engine.ResizeImage(destImagePath, imageURL)
				log.Println("Resized imagePath:", imagePath)
				if isResized {
					log.Println("Uploading to rackspace:", imagePath)
					engine.UploadImage(imagePath, func(err error, isSuccess bool, uploadImageURL string){
						if isSuccess {
							engine.AddImage(imageURL, uploadImageURL)	
						}
						})
				}
			}
		}
		engine.DownloadImage(imageURL, downloadImageCallBack)
	}
	writeResponse(&w, response)
}

func newResponseData(optimisedImageURL, originalImageURL string) ResponseData {
	var resp ResponseData
	resp.OriginalImageURL =  originalImageURL
	resp.OptimisedImageURL = optimisedImageURL
	return resp
}
func writeResponse(w *http.ResponseWriter, responseData ResponseData) {
	responseString, err := json.Marshal(responseData)
	if err != nil {
		log.Fatal(err)
	}
	
	(*w).Header().Add("Content-Type", "application/json; charset=utf-8")
	(*w).Write(responseString)
}
