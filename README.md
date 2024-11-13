# Genkit - Menu Items Suggestion

How to create streamable flows in Genkit

This demo will talk to `Google AI - Gemini` and will suggest a menu plate based on a given restaurant theme.
However, all the communication between the flow and the AI will be streamed in chunks

This demo was created by following [Genkit Golang Flows documentation](https://firebase.google.com/docs/genkit-go/flows)

## Dependencies

- Go 1.22 or later
- Node.js 20 or later

## How to run it

```shell
## Install Genkit
$ npm i -g genkit

## On the project folder, run the following commands
$ genkit init --model googleai

## Provide your Google GEN AI API key
export GOOGLE_GENAI_API_KEY=foobar

## Run flow in stream mode (chunks of data will be sent/received)
$ genkit run:flow menuItemSuggestion '"Mexican"' -s

## optional: if you just want to open Genkit Dev UI
$ genkit start
```
