package adapter

import "central-control/internal/model"

type ControllerAdapter interface {
    ApplyPlan(plan model.Plan) error
    GetStatus() model.ControllerStatus
}
