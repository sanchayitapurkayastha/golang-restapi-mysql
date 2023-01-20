package utils

import (
    "encoding/json"
    "io/ioutil"
    "net/http"
)

 //parse the received body which we have requested. it will be in json
 //unmarshal json
 func ParseBody(r *http.Request, x interface{}) {
    if body, err := ioutil.ReadAll(r.Body); err == nil {
        if err := json.Unmarshal([]byte(body), x); err != nil {
            return
        }
    }
 }