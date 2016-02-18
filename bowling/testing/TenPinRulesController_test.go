package controllers_test

import (
  "katas/bowling/controllers"
  "testing"
  //"fmt"
)

func TestSingleFrameScore(t *testing.T) {
  sut := controllers.NewTenPinRulesController()

  gameScore, err := sut.AddScore(0, 9, 0)

  if err != nil {
    t.Error("\tThere was an error adding score.", err)
  }

  if gameScore != 9 {
    t.Error("\tScore was incorrect")
//    fmt.Printf("\nIt was %d instead of 9", gameScore)
  }
}

func TestNonSpareStrikeScoreAddedToPreviousFrame(t *testing.T) {
  sut := controllers.NewTenPinRulesController()

  gameScore, err := sut.AddScore(5, 9, 0)

  if err != nil {
    t.Error("\tThere was an error adding score.", err)
  }

  if gameScore != 14 {
    t.Error("\tScore was incorrect")
//    fmt.Printf("\nIt was %d instead of 9", gameScore)
  }
}
 
//TODO Rethink the AddScore method.  Shouldnt pass in previous frame score
func TestSpareFirstFrameOtherNonMarkSecondFrame(t *testing.T) {
  sut := controllers.NewTenPinRulesController()

  gameScore, err := sut.AddScore(0, 9, 1)
  gameScore, err = sut.AddScore(10, 6, 2)

  if err != nil {
    t.Error("\tThere was an error adding score.", err)
  }

  if gameScore != 24 {
    t.Error("\tScore was incorrect")
//    fmt.Printf("\nIt was %d instead of 9", gameScore)
  }
}
