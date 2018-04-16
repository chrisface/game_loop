package main

import (
  "./game"
  "./gameloop"
  "time"
)

func main() {
  var g = game.NewGame()
  var gl = gameloop.NewGameLoop(&g)

  go gl.Start()

  time.Sleep(time.Second * 1)
  g.PlayerAddedChan <- 53
  g.PlayerAddedChan <- 54
  g.PlayerAddedChan <- 55
  g.PlayerAddedChan <- 56
  g.PlayerAddedChan <- 57
  g.PlayerAddedChan <- 58
  time.Sleep(time.Second * 1)
  g.PlayerLeftChan <- 4
  g.PlayerLeftChan <- 5
  g.PlayerLeftChan <- 6
  g.PlayerLeftChan <- 7
  g.PlayerLeftChan <- 8
  time.Sleep(time.Second * 10)
  gl.Stop()
}
