package controllers

import (
  //Standard library
  "errors"
  //"fmt"

  //Custom packages
  "katas/bowling/models"
)

var ErrTooManyPins = errors.New("You can not knock down more than 10 pins in one frame except the tenth")
var frames []models.Frame

type (
    // UserController represents the controller for operating on the User resource
    TenPinRulesController struct{}
)

func NewTenPinRulesController() *TenPinRulesController {
    frames = make([]models.Frame, 10)

    return &TenPinRulesController{}
}

func (rules *TenPinRulesController) RollFrame(frame models.Frame)(int, error){
  if frame.Ball1 + frame.Ball2 > 10 {
    return -1, ErrTooManyPins
  }

  if frame.Number <= 1 {
    frame.Score = frame.Ball1 + frame.Ball2
  } else {
    if (frames[frame.Number-2].Ball1 < 10 && (frames[frame.Number-2].Ball1 + frames[frame.Number-2].Ball2) == 10) {
      frame.Score =+ frames[frame.Number-2].Score + frame.Ball1
      frames[frame.Number-2].Score = frame.Score
    } else if (frames[frame.Number-2].Ball1 == 10){
      frame.Score =+ frames[frame.Number-2].Score + frame.Ball1 + frame.Ball2
      frames[frame.Number-2].Score = frame.Score
    }

    frame.Score =+ frames[frame.Number-2].Score + frame.Ball1 + frame.Ball2
  }

  frames[frame.Number-1] = frame

  return frame.Score, nil
}
