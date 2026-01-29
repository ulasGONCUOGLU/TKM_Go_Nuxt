package api

import (
    "encoding/json"
    "net/http"
    "central-control/internal/model"
    "central-control/internal/controller"
)

var manager = controller.NewControllerManager()

func RegisterRoutes(mux *http.ServeMux) {
    mux.HandleFunc("/plans/apply", applyPlan)
    mux.HandleFunc("/controllers/mock/status", mockStatus)
}

func applyPlan(w http.ResponseWriter, r *http.Request) {
    var p model.Plan
    json.NewDecoder(r.Body).Decode(&p)
    manager.ApplyPlan(p)
    w.Write([]byte(`{"ok":true}`))
}

func mockStatus(w http.ResponseWriter, r *http.Request) {
    status := manager.GetStatus()
    json.NewEncoder(w).Encode(status)
}
