package game

import (
  "fmt"
  "../geom"
  "../player"
)
type GameInput struct {

}

type Game struct  {
  GameArea geom.Rectangle
  Players []player.Player
  PlayerAddedChan chan int
  PlayerLeftChan chan int
}

func (g *Game) Update(){
  g.handleEvents()
  fmt.Println("Updating Game State")
}

func (g *Game) handleEvents() {
  g.popEvents(g.PlayerAddedChan, "Player Added")
  g.popEvents(g.PlayerLeftChan, "Player Left")
}

func (g *Game) popEvents(chanel chan int, name string) {
  fmt.Println(name, " events:", len(chanel))
  for len(chanel) > 0  {
    playerId := <- chanel
    fmt.Println("Processing:",name, playerId)
  }
}

const gameareawidth int = 100
const gameareaheight int = 100

func NewGame() Game {
  g := Game{}
  g.GameArea = geom.Rectangle{geom.Point{0 ,0}, geom.Point{gameareawidth, gameareaheight}}
  g.PlayerAddedChan = make(chan int, 10)
  g.PlayerLeftChan = make(chan int, 10)
  return g
}
