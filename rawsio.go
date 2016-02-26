package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func (r RawEnemyList) jsonexport(file string) {
	data, _ := json.MarshalIndent(&r, "", "\t")
	ioutil.WriteFile(file, data, os.ModePerm)
}

func (r *RawEnemyList) jsonimport(file string) {
	var data RawEnemyList
	jsfi, _ := ioutil.ReadFile(file)
	json.Unmarshal(jsfi, &data)
	for k, v := range data {
		(*r)[k] = v
	}
}
