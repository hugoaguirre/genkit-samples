// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"fmt"
	"log"
	"math"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
	"github.com/firebase/genkit/go/plugins/vertexai/modelgarden"
)

// Input struct for the song suggestion prompt and tool.
type SongSuggestionQuery struct {
	Genre    string  `json:"genre" jsonschema_description:"the preferred song genre by the user"`
	Mood     string  `json:"mood" jsonschema_description:"the user's current mood"`
	Length   float32 `json:"length,omitempty" jsonschema_description:"the lenght of the song"`
	Language string  `json:"language,omitempty" jsonschema_description:"the song language"`
}

// Output struct for the song suggestion prompt.
type SongSuggestion struct {
	SongName string `json:"songName"`
}

type GablorkenInput struct {
	Value float64 `json:"value"`
	Over  float64 `json:"over"`
}

func main() {
	ctx := context.Background()

	g, err := genkit.Init(ctx,
		genkit.WithDefaultModel("anthropic/claude-3-5-sonnet-v2"),
		genkit.WithPlugins(&modelgarden.Anthropic{}),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Define a tool that simulates suggesting a song.
	songTool := genkit.DefineTool(g, "songSuggestionTool",
		"Use this tool to retrieve a song suggestion based on the details given by the user, the user will specify the mood and the genre of the song in the inputs, as an output the tool should return a string with the song name suggestion",
		func(ctx *ai.ToolContext, input SongSuggestionQuery) (SongSuggestion, error) {
			return SongSuggestion{
				SongName: "A good song would be I drink wine by Adele",
			}, nil
		},
	)

	gablorkenTool := genkit.DefineTool(g, "gablorken",
		"use when need to calculate a gablorken",
		func(ctx *ai.ToolContext, input GablorkenInput) (float64, error) {
			return math.Pow(input.Value, input.Over), nil
		},
	)

	genkit.DefineFlow(g, "songFlow", func(ctx context.Context, input string) (string, error) {
		resp, err := genkit.Generate(ctx, g,
			ai.WithConfig(&ai.GenerationCommonConfig{Temperature: 0.0}),
			ai.WithPrompt(input),
			ai.WithTools(songTool, gablorkenTool),
			ai.WithToolChoice(ai.ToolChoiceAuto),
			ai.WithMaxTurns(25),
		)
		if err != nil {
			return "", fmt.Errorf("error executing song suggestion prompt: %w", err)
		}

		return resp.Text(), nil
	})

	<-ctx.Done()
}
