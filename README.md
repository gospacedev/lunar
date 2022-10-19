# Lunar

[![Apache Licence](https://img.shields.io/badge/licence-Apache%20License%202.0-blue)](https://www.apache.org/licenses/LICENSE-2.0)

Lunar is a music player for terminal enthusiasts.

## Build

Install Lunar by running this command:

```powershell
go install github.com/gospacedev/lunar
```

## Usage
Run Lunar to explore music:

```cmd
lunar

Use the arrow keys to navigate: ↓ ↑ → ←
? Select music:
  > Charlie Puth.mp3
    The Weekend.mp3
    Twenty One Pliots.mp3
```

Control the audio with keys:
```
Playing Charlie Puth...
Pause and play: [ENTER]
Back: [BACKSPACE]
Quit: [ESC]
```

Change filepath by running new:

```cmd
lunar new

Enter new filepath...
C:/Users/grant/Music
Filepath successfully added
```

## Todo List

- enter filepath as argument to lunar new command
- add new options: change speed, shuffle, playback