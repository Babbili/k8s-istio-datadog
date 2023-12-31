from flask import Flask
from get_text.get_text import gettext
import os
from google.cloud import texttospeech

app = Flask(__name__)


def main():
    # get books from goapi to pass them as a texttospeech input 
    books = gettext()
    # Instantiates a client
    client = texttospeech.TextToSpeechClient.from_service_account_json(os.getenv('SA_PATH'))
    
    # Set the text input to be synthesized
    synthesis_input = texttospeech.SynthesisInput(text=books)
    
    # Build the voice request, select the language code ("en-US") and the ssml
    # voice gender ("neutral")
    voice = texttospeech.VoiceSelectionParams(
        language_code="en-US", ssml_gender=texttospeech.SsmlVoiceGender.NEUTRAL
    )
    
    # Select the type of audio file you want returned
    audio_config = texttospeech.AudioConfig(
        audio_encoding=texttospeech.AudioEncoding.MP3
    )
    
    # Perform the text-to-speech request on the text input with the selected
    # voice parameters and audio file type
    response = client.synthesize_speech(
        input=synthesis_input, voice=voice, audio_config=audio_config
    )
    
    # The response's audio_content is binary.
    with open("output.mp3", "wb") as out:
        # Write the response to the output file.
        out.write(response.audio_content)
        print('Audio content written to file "output.mp3"')

@app.route('/')
def root():
    return "TTS canary version (Python)!"

if __name__ == "__main__":
    main()
    app.run(debug=True, host="0.0.0.0", port=1323)