package mock

import (
    "central-control/internal/model"
    "time"
    "sync"
)

type MockAdapter struct {
    mu sync.Mutex
    status model.ControllerStatus
}

func NewMockAdapter() *MockAdapter {
    return &MockAdapter{
        status: model.ControllerStatus{
            ActivePhase: 0,
            RemainingSec: 0,
        },
    }
}

func (m *MockAdapter) ApplyPlan(plan model.Plan) error {
    go func() {
        for _, ph := range plan.Phases {
            m.mu.Lock()
            m.status.ActivePhase = ph.ID
            m.status.RemainingSec = ph.Duration
            m.mu.Unlock()
            for i := ph.Duration; i > 0; i-- {
                time.Sleep(time.Second)
                m.mu.Lock()
                m.status.RemainingSec = i
                m.mu.Unlock()
            }
        }
    }()
    return nil
}

func (m *MockAdapter) GetStatus() model.ControllerStatus {
    m.mu.Lock()
    defer m.mu.Unlock()
    return m.status
}
