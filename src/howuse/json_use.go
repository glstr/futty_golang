package howuse

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/xeipuuv/gojsonpointer"
)

func Decodejson() {
	example := "{\"hello\": 12343434343}"
	c := map[string]interface{}{}
	dec := json.NewDecoder(strings.NewReader(example))
	err := dec.Decode(&c)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(c)
}

func EncodeArray() {
	ex := []int{1, 2, 3}
	str, err := json.Marshal(ex)
	if err != nil {
		log.Printf("err_msg:%s", err.Error())
		return
	}

	log.Printf("res:%s", str)
}

func DecodeArray() {
	array := "[1, 2, 3]"
	var res []int64
	err := json.Unmarshal([]byte(array), &res)
	if err != nil {
		log.Printf("err_msg:%s", err.Error())
		return
	}
	log.Printf("res:%v", res)
}

func JsonModify() {
	example := "{\"id\":10112889999101,\"app_id\":40011,\"type\":\"html\",\"time\":1555050164,\"level\":0,\"expire\":0,\"msg\":{\"o2o\":0,\"opentype\":1,\"ext\":{\"type\":20,\"content\":{\"data_list\":[{\"data\":{\"api\":\"splash_notice\",\"options\":{\"callback_add_params\":{\"lcs_call\":\"1\"}},\"data\":{\"search_id\":\"\",\"query\":\"\",\"ext\":\"\"}},\"meta\":[[\"cuid\",\"==\",\"9cd3b19ac5dbb5bd6072477c22fe56032921cd41\"],[\"appid\",\"==\",\"405384\"]]}],\"version\":1}},\"fg\":256},\"uid\":1070724739}"
	ret, err := modifyText(example)
	if err != nil {
		log.Printf("error:%s", err.Error())
		return
	}

	fmt.Printf("res:%v", ret)
}

func modifyText(input string) (string, error) {
	var res string
	var jsonDoc map[string]interface{}
	json.Unmarshal([]byte(input), &jsonDoc)
	pointerString := "/msg/ext/content/data_list"
	pointer, err := gojsonpointer.NewJsonPointer(pointerString)
	if err != nil {
		return res, err
	}
	data_list, _, err := pointer.Get(jsonDoc)
	if err != nil {
		return res, err
	}

	new_data_list, ok := data_list.([]interface{})
	var obj_list []interface{}
	if ok {
		for _, value := range new_data_list {
			if value_map, ok := value.(map[string]interface{}); ok {
				for key, value := range value_map {
					if key == "data" {
						obj_list = append(obj_list, value)
					}
				}
			} else {
				return res, errors.New("param error")
			}
		}
	} else {
		return res, errors.New("param error")
	}

	pointer.Set(jsonDoc, obj_list)
	b, err := json.Marshal(jsonDoc)
	if err != nil {
		return res, err
	}

	return string(b), nil
}
