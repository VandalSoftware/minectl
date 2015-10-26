package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
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
	flag.Parse()

	resp, err := http.Get(v)
	if err != nil {
		log.Println("fetch error", err.Error())
		os.Exit(1)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Println(resp.Status)
		os.Exit(1)
		return
	}

	d := json.NewDecoder(resp.Body)

	var versions versions
	if err := d.Decode(&versions); err != nil {
		log.Println("couldn't parse %s", v)
		os.Exit(1)
		return
	}

	if len(flag.Args()) == 0 {
		fmt.Printf("release: %s, snapshot: %s\n", versions.Latest.Release, versions.Latest.Snapshot)
		return
	}

	switch flag.Args()[0] {
	case "release":
		fmt.Println(versions.Latest.Release)
	case "snapshot":
		fmt.Println(versions.Latest.Snapshot)
	default:
		fmt.Println("minectl: invalid command --", flag.Args()[0])
		fmt.Println("Try 'minectl --help' for more information")
	}
}
