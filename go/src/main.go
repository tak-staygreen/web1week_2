package main

import (
    "bytes"
    "io/ioutil"
    "net/http"
    "regexp"
    "math/rand"
    "time"
    "encoding/json"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    // get document
    searchUrl := "XXXXX" // search query url
    document, _ := http.Get(searchUrl)
    defer document.Body.Close()

    // get html from document
    body, _ := ioutil.ReadAll(document.Body)
    buf := bytes.NewBuffer(body)
    html := buf.String()

    // make url list from html (.jpg only)
    exp := regexp.MustCompile(`https?://[\w/:%#\$&\?\(\)~\.=\+\-]+jpg`)
    urlList := exp.FindAllStringSubmatch(html, -1)

    // select 30 urls from url list at random
    var selectedUrlList []string
    rand.Seed(time.Now().UnixNano())
    for i := 0; i < 30; i++ {
        num := rand.Intn(len(urlList))
	    selectedUrlList = append(selectedUrlList, urlList[num]...)
    }

    // make response
    response, _ := json.Marshal(selectedUrlList)

    // return response
    r.GET("/getcat", func(c *gin.Context) {
        c.JSON(200, gin.H {
            "response" : string(response),
        })
    })

    r.Run(":9000")
}