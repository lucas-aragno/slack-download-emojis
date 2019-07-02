package main

import (
    "os"
    "fmt"
    "net/http"
    "encoding/json"
    "io/ioutil"
    "github.com/smallnest/goreq"
)

var SLACK_TOKEN string = os.Getenv("SLACK_TOKEN") 

func check (e error) {
    if e != nil {
        panic(e)
    }
}

func getImage (emojiName string, url string) {
    client := &http.Client{}
    resp, body, err := goreq.New().SetClient(client).Get(url).End()

    if (err == nil && resp.StatusCode == 200) {
        var name string = fmt.Sprintf("emojis/%s.png", emojiName)
        err := ioutil.WriteFile(name, []byte(body), 0644)
        check(err)
    }
}

func download () {
    var url string = fmt.Sprintf("https://slack.com/api/emoji.list?token=%s", SLACK_TOKEN) 
    client := &http.Client{}
    resp, body, err := goreq.New().SetClient(client).
        Get(url).
        End()
    
    if (err != nil) {
        fmt.Println("Error fetching emojis")
    } else if (resp.StatusCode == 200) {
        jsonMap :=  make(map[string]interface{})
        err := json.Unmarshal([] byte(body), &jsonMap)
        if err != nil {
            fmt.Println("Error unrmashalling response")
        }

        emojis := jsonMap["emoji"].(map[string]interface{})

        for key, value := range emojis {
            getImage(key, value.(string))
        }
    }
}

func main () {
    download()
}
