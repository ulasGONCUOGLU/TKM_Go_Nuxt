package model

type Phase struct {
    ID int `json:"id"`
    Duration int `json:"duration"`
}

type Plan struct {
    ID string `json:"id"`
    Phases []Phase `json:"phases"`
}

type ControllerStatus struct {
    ActivePhase int `json:"active_phase"`
    RemainingSec int `json:"remaining_sec"`
}
