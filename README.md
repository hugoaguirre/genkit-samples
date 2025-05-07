# Genkit Samples (Go)

A collection of Genkit samples using the Go runtime

## Dependencies

Please follow this [quick](https://firebase.google.com/docs/genkit-go/devtools) guide

## How to run it

```shell
## Provide your Google AI API key
export GOOGLE_GENAI_API_KEY=foobar

## If using Vertex AI or ModelGarden, set the following env vars
export GOOGLE_CLOUD_LOCATION=us-west-1
export GOOGLE_CLOUD_PROJECT=my-cool-project

## Set the environment
export GENKIT_ENV=dev

## Use Genkit CLI to run the sample
genkit start -- go run tools/main.go

```
