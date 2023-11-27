package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	gettext "text-to-speech/gettext"

	echotrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/labstack/echo.v4"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"

	texttospeech "cloud.google.com/go/texttospeech/apiv1"
	"cloud.google.com/go/texttospeech/apiv1/texttospeechpb"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"google.golang.org/api/option"
)

// Handler
func echotts(c echo.Context) error {
	return c.String(http.StatusOK, "text-to-peech Go!")
}

func main() {
	// will be passed as text-to-speech req input
	books := gettext.Gettex()

	// Instantiates a client.
	ctx := context.Background()

	client, err := texttospeech.NewClient(ctx, option.WithCredentialsFile(os.Getenv("SA_PATH")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// Perform the text-to-speech request on the text input with the selected
	// voice parameters and audio file type.
	req := texttospeechpb.SynthesizeSpeechRequest{
		// Set the text input to be synthesized.
		Input: &texttospeechpb.SynthesisInput{
			InputSource: &texttospeechpb.SynthesisInput_Text{Text: books},
		},
		// Build the voice request, select the language code ("en-US") and the SSML
		// voice gender ("neutral").
		Voice: &texttospeechpb.VoiceSelectionParams{
			LanguageCode: "en-US",
			SsmlGender:   texttospeechpb.SsmlVoiceGender_NEUTRAL,
		},
		// Select the type of audio file you want returned.
		AudioConfig: &texttospeechpb.AudioConfig{
			AudioEncoding: texttospeechpb.AudioEncoding_MP3,
		},
	}

	resp, err := client.SynthesizeSpeech(ctx, &req)
	if err != nil {
		log.Fatal(err)
	}

	// The resp's AudioContent is binary.
	filename := "output.mp3"
	err = os.WriteFile(filename, resp.AudioContent, 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Audio content written to file: %v\n", filename)

	tracer.Start(
		tracer.WithEnv("prod"),
		tracer.WithService("text-to-speech"),
		tracer.WithServiceVersion("v1"),
	)
	// When the tracer is stopped, it will flush everything it has to the Datadog Agent before quitting.
	defer tracer.Stop()

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(echotrace.Middleware(echotrace.WithServiceName("text-to-speech"), echotrace.WithCustomTag("env", "prod")))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", echotts)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
