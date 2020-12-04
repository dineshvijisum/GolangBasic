package main
import (
    "fmt"
    "encoding/json"
    "log"
    "strings"
)
var dataAsString = ``
type Item struct {
    Id          int    `json:"id"`
    Category    string `json:"category"`
    Name        string `json:"name"`
    Description string `json:"description"`
}
type CategoryToItemSliceMap map[string][]Item
type CategoryToIndexItemMap map[string]map[int]Item
func main() {
    decoder := json.NewDecoder(strings.NewReader(dataAsString))
    var ourData []Item
    for decoder.More() {
        var it Item
        err := decoder.Decode(&it)
        if err != nil {
            log.Fatalln(err)
        }
        ourData = append(ourData, it)
    }
    catToItemSlice := CategoryToItemSliceMap{}
    for _,v := range ourData {
        catToItemSlice[v.Category] = append(catToItemSlice[v.Category],v)
    }
    catToIndexItemMap := CategoryToIndexItemMap{}
    for k,v := range catToItemSlice {
        if catToIndexItemMap[k] == nil {
            catToIndexItemMap[k] = map[int]Item{}
        }
        for index, item := range v {
           catToIndexItemMap[k][index] = item
        }
    }
    fmt.Printf("elements: ")
    out, err := json.MarshalIndent(catToIndexItemMap, "", "    ")
    if err != nil {
        log.Fatalln(err)
    }
    fmt.Println(string(out))
}
