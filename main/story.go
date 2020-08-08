package main

import (
	"fmt"
	"strconv"
)

func NewStoryTeller(story *Story) *StoryTeller {
	return &StoryTeller{story: story}
}

type StoryArc struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}
type Story struct {
	StoryArcs map[string]StoryArc
}

type StoryTeller struct {
	story *Story
}

func (s *StoryTeller) Story() Story {
	return *s.story
}

func (s *StoryTeller) TellStory() {
	key := "intro"
Book:
	for {
		key = s.read(key)
		if key == "" {
			break Book
		}
	}
}

func (s *StoryTeller) read(key string) string {
	arc := s.Story().StoryArcs[key]
	printStory(arc.Story)

	if len(arc.Options) == 0 { // check for end of book
		fmt.Println("The end!")
		return ""
	}

	paths := printOptions(arc)

	var chosenOption string
	_, err := fmt.Scanf("%s\n", &chosenOption)
	if err != nil {
		fmt.Errorf("Error reading option")
		return ""
	}
	return paths[chosenOption]
}

func printStory(story []string) {
	for _, line := range story {
		fmt.Println(line)
		fmt.Println("(to continue, press Enter)")
		var wait string
		_, _ = fmt.Scanln(&wait)
	}
}
func printOptions(arc StoryArc) map[string]string {
	fmt.Println("What would you like to do?")
	optionMap := make(map[string]string)
	for i, option := range arc.Options {
		key := strconv.Itoa(i + 1)
		optionMap[key] = option.Arc
		fmt.Printf("%s) %s\n", key, option.Text)
	}
	return optionMap
}
