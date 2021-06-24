package cfg

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	DaysHolidaysThreshold = "days"
	GlUsers               = "glusers"
	MailReceiver          = "mail_receiver"
	ResetDays             = "resetdays"
	TestUser              = "testuser"
	BaseUrl               = "base_url"
)

type Items struct {
	Items []Item
}
type Item struct {
	Key   string
	Value interface{}
}

var conf Items

func Read(stage string) {
	if stage == "local" {
		log.Println("read local properties")
		readFileConfig(&conf)
	} else {
		log.Println("read db properties")
	}
}
func readFileConfig(conf *Items) {
	var items []Item
	cwd, _ := os.Getwd()
	filename := filepath.Join(cwd, "config/config")
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if equal := strings.Index(line, "="); equal >= 0 {
			if key := strings.TrimSpace(line[:equal]); len(key) > 0 {
				value := ""
				if len(line) > equal {
					value = strings.TrimSpace(line[equal+1:])
				}
				item := Item{Key: key, Value: value}
				items = append(items, item)
			}
		}
	}
	conf.Items = items
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func Get(key string) interface{} {
	log.Println("get config with key: " + key)
	for _, item := range conf.Items {
		if item.Key == key {
			log.Println("found config for key: " + key + " with value: " + fmt.Sprintf("%v", item.Value))
			return item.Value
		}
	}
	return ""
}

func GetStringArray(key string) []string {
	var aString []string
	val := Get(key)
	switch v := val.(type) {
	case string:
		if comma := strings.Index(val.(string), ","); comma >= 0 {
			aString = strings.Split(val.(string), ",")
		} else {
			aString = append(aString, val.(string))
		}
	case []interface{}:
		aInterface := val.([]interface{})
		aString = make([]string, 0)
		for _, va := range aInterface {
			aString = append(aString, va.(string))
		}
	default:
		log.Fatalf("undefined type: %s", v)
	}
	return aString
}
