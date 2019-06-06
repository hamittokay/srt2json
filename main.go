package main

import (
	"encoding/json"
	"errors"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Scene struct {
	ID        string
	Start     string
	End       string
	Sentences string
}

func getFileContent(filepath string) string {
	dat, err := ioutil.ReadFile(filepath)
	check(err)
	return string(dat)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func generateSceneItem(scene string) (*Scene, error) {
	sp := strings.Split
	sceneItems := sp(scene, "\n")

	if len(sceneItems) >= 2 {
		ID, Time, Sentences :=
			sceneItems[0], sp(sceneItems[1], " --> "), sceneItems[2]

		Start := Time[0]
		End := Time[1]

		SceneItem := Scene{
			ID:        ID,
			Start:     Start,
			End:       End,
			Sentences: Sentences,
		}

		return &SceneItem, nil
	}

	return nil, errors.New("Range Error")
}

func GetScenes(path string) []Scene {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	filepath := dir + "/" + path
	subtitles := getFileContent(filepath)

	var scenes []Scene

	sceneChunk := strings.Split(subtitles, "\n\n")

	for _, scene := range sceneChunk {
		SceneItem, err := generateSceneItem(scene)

		if err != nil {
			// fmt.Println(err)
		} else {
			scenes = append(scenes, *SceneItem)
		}
	}

	return scenes
}

func Srt2Json() {
	flag.Parse()

	filename := flag.Arg(0)
	exportFileName := flag.Arg(1)

	subt := GetScenes(filename)
	file, _ := json.MarshalIndent(subt, "  ", "  ")

	if exportFileName == "" {
		exportFileName = strings.Replace(filename, ".srt", "", -1)
	}

	_ = ioutil.WriteFile(exportFileName+".json", file, 0644)
}

func main() {
	Srt2Json()
}
