package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
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

	if len(flag.Args()) == 0 {
		versions, err := fetchVersions()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
			return
		}
		fmt.Printf("release: %s, snapshot: %s\n", versions.Latest.Release, versions.Latest.Snapshot)
		return
	}

	switch flag.Args()[0] {
	case "release":
		versions, err := fetchVersions()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
			return
		}
		fmt.Println(versions.Latest.Release)
	case "snapshot":
		versions, err := fetchVersions()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
			return
		}
		fmt.Println(versions.Latest.Snapshot)
	default:
		fmt.Println("minectl: invalid command --", flag.Args()[0])
		fmt.Println("Try 'minectl --help' for more information")
	}
}

func fetchVersions() (*versions, error) {
	resp, err := http.Get(v)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	d := json.NewDecoder(resp.Body)

	var versions versions
	if err := d.Decode(&versions); err != nil {
		return nil, errors.New("parse error")
	}
	return &versions, nil
}
