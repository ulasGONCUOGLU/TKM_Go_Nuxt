package controller

import (
    "central-control/internal/model"
    "central-control/internal/adapter/mock"
)

type ControllerManager struct {
    adapter *mock.MockAdapter
}

func NewControllerManager() *ControllerManager {
    return &ControllerManager{
        adapter: mock.NewMockAdapter(),
    }
}

func (c *ControllerManager) ApplyPlan(p model.Plan) {
    c.adapter.ApplyPlan(p)
}

func (c *ControllerManager) GetStatus() model.ControllerStatus {
    return c.adapter.GetStatus()
}
