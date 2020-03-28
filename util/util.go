package util

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "net/url"
    "os"
    "strings"
    "sync"
)

var lock sync.Mutex
var output_folder = "results"

func getEnv(key string, defaultVal string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    fmt.Fprintln(os.Stderr, key, "not found in env")
    return defaultVal
}

func getAccountId() string {
    return getEnv("AWS_ACCOUNT_ID", "")
}

func Save(service string, region string, object interface{}) error {
    lock.Lock()
    defer lock.Unlock()
    if !object.(Results).isEmpty() {
        directory := fmt.Sprintf("%s/%s", output_folder, service)
        os.MkdirAll(directory, os.ModePerm)

        accountid := getAccountId()
        path := ""
        if len(accountid) > 0 {
            path = fmt.Sprintf("%s/%s/%s.%s.json", output_folder, service, accountid, region)
        } else {
            path = fmt.Sprintf("%s/%s/%s.json", output_folder, service, region)
        }
        f, err := os.Create(path)
        if err != nil {
            fmt.Println(err)
            return err
        }
        defer f.Close()
        r, err := Marshal(object)
        if err != nil {
            fmt.Println(err)
            return err
        }
        _, err = io.Copy(f, r)
        if err != nil {
            fmt.Println(err)
            return nil
        }
        return nil
    }
    return nil
}

var Marshal = func(object interface{}) (io.Reader, error) {
    jsonOutput, err := json.MarshalIndent(object, "", "\t")
    if err != nil {
        return nil, err
    }
    return bytes.NewReader(jsonOutput), nil
}

func UniqueStrings(input []string) []string {
    uniqueStringSlice := make([]string, 0, len(input))
    uniqueMapMap := make(map[string]bool)

    for _, val := range input {
        if _, ok := uniqueMapMap[val]; !ok {
            uniqueMapMap[val] = true
            uniqueStringSlice = append(uniqueStringSlice, val)
        }
    }

    return uniqueStringSlice
}

func GetHostnameFromUrl(urlString string) string {
    parsedUrl, err := url.Parse(strings.ToLower(urlString))
    if err != nil {
        return ""
    }
    return parsedUrl.Host
}
