package main

import (
    "fmt"
    "gopkg.in/yaml.v3"
    "io/ioutil"
    "os"
)

type Config struct {
    Fields []string `yaml: "fields"`
    Another yaml.Node `yaml: "another"`
}

type Dataset struct {
    Time string `yaml: "time"`
    Fruit map[string]int
}

func main() {
    filename := os.Args[1]
    datafile := os.Args[2]
    fmt.Printf("config file: %s\n\n", filename)
    fmt.Printf("data file: %s\n\n", datafile)
    var config Config
    var dataset Dataset

    source, err := ioutil.ReadFile(filename)
    if err != nil {
        panic(err)
    }

    err = yaml.Unmarshal(source, &config)
    if err != nil {
        panic(err)
    }

//    fmt.Printf("Cfg Value : %#v\n", config.Fields)
//    fmt.Printf("Cfg Value2: %#v\n", config.Another)

    data, err := ioutil.ReadFile(datafile)
    if err != nil {
        panic(err)
    }
//    fmt.Printf("Raw Data: \n%s\n\n", data)
    err = yaml.Unmarshal(data, &dataset)
    if err != nil {
        panic(err)
    }
//    fmt.Printf("Data Value: %#v\n", dataset.Time)
//    fmt.Printf("Data Value: %#v\n", dataset.Fruit)

    for k, v := range dataset.Fruit {
        exist := false
//        fmt.Printf("%s ...\n", k)
        for _, d := range config.Fields {
//            fmt.Printf(".. %s ..", d)
            
            if d == k {
                exist = true
                break
            }
        }
        if exist {
            fmt.Printf("%s -> %d\n", k, v)
        }
    }
}
