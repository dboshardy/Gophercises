package cyoa

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
	for {
		paths := printOptions(arc)
		choice := awaitChoice()
		if val, ok := paths[choice]; ok {
			return val
		}
		if choice == "q" || choice == "quit" {
			return ""
		}
		fmt.Printf(`"%s" is not a valid option, please choose a valid option. 
To quit, type "q" or "quit".
`, choice)
	}
}

func awaitChoice() string {
	var choice string
	_, err := fmt.Scanf("%s\n", &choice)
	if err != nil {
		fmt.Errorf("Error reading option: &s", choice)
		return ""
	}
	return choice
}

func printStory(story []string) {
	for _, line := range story {
		fmt.Printf("%s\n(press Enter to continue)", line)
		var wait string
		_, _ = fmt.Scanf("\n", &wait)
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
