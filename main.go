package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	// Import Genkit and the Google AI plugin
	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
	"github.com/firebase/genkit/go/plugins/googleai"
)

type (
	InputType  string
	OutputType string
	StreamType string
)

func makeMenuItemSuggestion(restaurantTheme InputType, stream chan<- StreamType) error {
	// TODO: check if this is the right place to init the AI model
	m := googleai.Model("gemini-1.5-flash")
	if m == nil {
		return errors.New("menuSuggestionFlow: failed to find model")
	}

	resp, err := ai.Generate(
		context.Background(),
		m,
		ai.WithConfig(&ai.GenerationCommonConfig{Temperature: 1}),
		ai.WithTextPrompt(fmt.Sprintf(`Suggest an item for the menu of a %s themed restaurant`, restaurantTheme)),
		ai.WithStreaming(
			func(ctx context.Context, grc *ai.GenerateResponseChunk) error {
				fmt.Printf("Chunk: %s\n", grc.Text())
				return nil
			}))
	if err != nil {
		return err
	}

	stream <- StreamType(resp.Text())
	return nil
}

func main() {
	ctx := context.Background()

	// Initialize the Google AI plugin. When you pass nil for the
	// Config parameter, the Google AI plugin will get the API key from the
	// GOOGLE_GENAI_API_KEY environment variable, which is the recommended
	// practice.
	if err := googleai.Init(ctx, nil); err != nil {
		log.Fatal(err)
	}

	genkit.DefineStreamingFlow("menuSuggestionFlow",
		func(ctx context.Context, restaurantTheme InputType, callback func(context.Context, StreamType) error) (OutputType, error) {
			var menu strings.Builder
			menuChunks := make(chan StreamType)
			go makeMenuItemSuggestion(restaurantTheme, menuChunks)
			for {
				chunk, ok := <-menuChunks
				if !ok {
					break
				}
				if callback != nil {
					callback(context.Background(), chunk)
				}
				menu.WriteString(string(chunk))
			}
			return OutputType(menu.String()), nil
		})

	// Initialize Genkit and start a flow server. This call must come last,
	// after all of your plug-in configuration and flow definitions. When you
	// pass a nil configuration to Init, Genkit starts a local flow server,
	// which you can interact with using the developer UI.
	if err := genkit.Init(ctx, nil); err != nil {
		log.Fatal(err)
	}
}
