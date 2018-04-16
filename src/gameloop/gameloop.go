package gameloop

import (
  "fmt"
  "time"
  "../game"
)

type GameLoop struct {
  Fps time.Duration
  FrameLength time.Duration
  FrameCount int
  FrameTicker *time.Ticker
  Game *game.Game
}

func (gl *GameLoop) Start() {
  fmt.Println("Started GameLoop")

  var lastFrameTime = time.Now()

  for currentFrameTime := range gl.FrameTicker.C {
    gl.FrameCount = gl.FrameCount + 1
    gl.Game.Update()
    fmt.Println("Frame: ", gl.FrameCount, currentFrameTime.Sub(lastFrameTime), gl.Game.Players)
    lastFrameTime = currentFrameTime
  }
}

func (gl *GameLoop) Stop() {
  gl.FrameTicker.Stop()
}

const fpsDefault time.Duration = 30

func NewGameLoop(game *game.Game) GameLoop {
  gl := GameLoop{}
  gl.Fps = fpsDefault
  gl.FrameLength = time.Duration(time.Second / gl.Fps)
  gl.FrameTicker = time.NewTicker(gl.FrameLength)
  gl.Game = game
  return gl
}
