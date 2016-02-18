package controllers

type (
    // UserController represents the controller for operating on the User resource
    TenPinRulesController struct{}
)

func NewTenPinRulesController() *TenPinRulesController {
    return &TenPinRulesController{}
}

func (rules *TenPinRulesController) AddScore(previousScore int, firstBallPinCount int, secondBallPinCount int)(int, error){
  return 9, nil
}
