package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	_ "os"
)

type latest struct {
	Snapshot string `json:"snapshot"`
	Release  string `json:"release"`
}

type version struct {
	Id          string `json:"id"`
	Time        string `json:"time"`
	ReleaseTime string `json:"releaseTime"`
	Type        string `json:"snapshot"`
}

type versions struct {
	Latest   latest    `json:"latest"`
	Versions []version `json:"versions"`
}

const v = "https://s3.amazonaws.com/Minecraft.Download/versions/versions.json"

func main() {
	// update
	resp, err := http.Get(v)
	if err != nil {
		log.Println("fetch error", err.Error())
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Println(resp.Status)
		return
	}

	//check
	fname := "versions.json"
	/*
		f, err := os.Open(fname)
		if err != nil {
			log.Println("couldn't open %s", fname)
			return
		}
		defer f.Close()
	*/

	d := json.NewDecoder(resp.Body)

	var versions versions
	if err := d.Decode(&versions); err != nil {
		log.Println("couldn't parse %s", fname)
		return
	}

	fmt.Printf("release: %s, snapshot: %s\n", versions.Latest.Release, versions.Latest.Snapshot)
}
