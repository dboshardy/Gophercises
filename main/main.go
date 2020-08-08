package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
)

func main() {
	storyFile := flag.String("story", "gopher.json", "the json file containing the story you wish to read.")

	story, err := parseStory(storyFile)
	if err != nil {
		_ = fmt.Errorf("Error reading story file %s", err)
	}

	tellStory(story)
}

func tellStory(story *Story) {
	storyTeller := NewStoryTeller(story)
	storyTeller.TellStory()
}

func parseStory(file *string) (*Story, error) {
	var storyMap map[string]StoryArc

	fileBytes, err := ioutil.ReadFile(*file)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(fileBytes, &storyMap)
	if err != nil {
		return nil, err
	}
	return &Story{StoryArcs: storyMap}, nil
}
