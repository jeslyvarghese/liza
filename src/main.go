package main

import (
	"github.com/gorilla/mux"
	"github.com/jeslyvarghese/liza/src/engine.go"
	"log"
	"net/http"
	"net/url"
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

	r := mux.NewRouter()
	r.HandleFunc("/resize", resize)
	http.ListenAndServe(8080, nil)
}

func resize(w http.ResponseWriter, r *http.Request) {
	imageURL =  url.QueryUnescape(r.URL.Query().Get("imageURL"))
	var response ResponseData
	if optimisedImageURL, hasOptimisedImageURL = engine.CheckHasImage(imageURL); hasOptimisedImageURL {
		response = newResponseData(optimisedImageURL, imageURL)
	} else {
		engine.DownloadImage(imageURL, func(err error, isSuccess bool, destImagePath string) {
			if isSuccess {
				imagePath, isResized := engine.ResizeImage(destImagePath, imageURL)
				if isResized {
					engine.UploadImage(imagePath, func(err error, isSuccess, uploadImageURL string){
						if isSuccess {
							engine.AddImage(imageURL, uploadImageURL)	
						}
						})
				}
			}
		})
	}
	writeResponse(w, response)
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
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(responseString)
}
