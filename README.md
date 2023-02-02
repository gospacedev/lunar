# Lunar - Fast Terminal Audio Player

[![Go](https://github.com/gospacedev/lunar/actions/workflows/go.yml/badge.svg)](https://github.com/gospacedev/lunar/actions/workflows/go.yml)
[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://github.com/gospacedev/lunar/blob/master/LICENSE)

Lunar is an audio player specifically designed for terminal enthusiasts. With its fast performance and support for multiple audio formats, it's the perfect tool for playing music directly from the terminal.

## Features

- Lunar supports `WAV`, `MP3`, `OGG`, and `FLAC` audio formats.

- Customize the look of Lunar with different built-in color themes.

- Control audio playback with options to loop, pause/resume, change position, change volume, and adjust playback speed.

## Installation

To build Lunar from the source, make sure you have [Go](https://go.dev/) installed. Then run the following command:
```go
go install github.com\gospacedev\lunar@latest
```

You can download the binary from the [release](https://github.com/gospacedev/lunar/releases) page.

## Getting Started

To change the filepath to your audio files, run the following command:

```
lunar newpath "C:\Users\grant\Music"
```

Now you're ready to run Lunar:

```
lunar
```

You can then select the audio you want to play. For more information on [color themes](https://github.com/gospacedev/lunar/blob/master/doc/themes.md) and how to change them, please refer to the Lunar documentation:

![Screenshot (230)](https://user-images.githubusercontent.com/83633399/201524750-94267cca-fc3f-4c16-91b5-ed7f9b535016.png)

![Screenshot (235)](https://user-images.githubusercontent.com/83633399/201524728-1ae12760-1ae1-4120-939c-0f2579a31560.png)
