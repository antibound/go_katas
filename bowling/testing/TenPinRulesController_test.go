package controllers_test

import (
  "katas/bowling/controllers"
  "katas/bowling/models"
  "testing"
  //"fmt"
)

func TestSingleFrameScore(t *testing.T) {
  sut := controllers.NewTenPinRulesController()
  frame := models.Frame{Ball1: 9, Ball2: 0, Number: 1}

  gameScore, err := sut.RollFrame(frame)

  if err != nil {
    t.Error("\tThere was an error adding score.", err)
  }

  if gameScore != 9 {
    t.Errorf("\tScore should be 9 but was %d", gameScore)
//    fmt.Printf("\nIt was %d instead of 9", gameScore)
  }
}

func TestNonSpareStrikeScoreAddedToPreviousFrame(t *testing.T) {
  sut := controllers.NewTenPinRulesController()

  frame := models.Frame{Ball1: 9, Ball2: 0, Number:1}
  gameScore, err := sut.RollFrame(frame)

  frame = models.Frame{Ball1: 5, Ball2: 0, Number:2}
  gameScore, err = sut.RollFrame(frame)

  if err != nil {
    t.Error("\tThere was an error adding score.", err)
  }

  if gameScore != 14 {
    t.Errorf("\tScore should be 14 but was %d", gameScore)
//    fmt.Printf("\nIt was %d instead of 9", gameScore)
  }
}

func TestTooManyPinsError(t *testing.T){
  sut := controllers.NewTenPinRulesController()
  frame := models.Frame{Ball1: 9, Ball2: 2}
  gameScore, err := sut.RollFrame(frame)

  if gameScore > -1 {
    t.Error("\tAn error was expected and this means no score should be returned")
  }

  if err == nil || err.Error() != "You can not knock down more than 10 pins in one frame except the tenth" {
    t.Error("\tExpected error of too many pins falling but received other")
  }
}

func TestSpareFirstFrameOtherNonMarkSecondFrame(t *testing.T) {
  sut := controllers.NewTenPinRulesController()

    frame := models.Frame{Ball1: 9, Ball2: 1, Number:1}
    gameScore, err := sut.RollFrame(frame)

    frame = models.Frame{Ball1: 5, Ball2: 4, Number:2}
    gameScore, err = sut.RollFrame(frame)

  if err != nil {
    t.Error("\tThere was an error adding score.", err)
  }

  if gameScore != 24 {
    t.Errorf("\tScore should be 24 but was %d", gameScore)
  }
}

func TestFirstFrameStrikeThenOpen(t *testing.T){
  sut := controllers.NewTenPinRulesController()

  frame := models.Frame{Ball1: 10, Ball2: 0, Number:1}
  gameScore, err := sut.RollFrame(frame)

  frame = models.Frame{Ball1: 5, Ball2: 4, Number:2}
  gameScore, err = sut.RollFrame(frame)

  if err != nil {
    t.Error("\tThere was an error adding score.", err)
  }

  if gameScore != 28 {
    t.Errorf("\tScore should be 28 but was %d", gameScore)
  }
}
//FirstFrameStrikeThenSpare
//First2FramesStrikeThenOpen
//PefectGame
//FirstFrameSpareThenStrike
