package models

import (
    "encoding/json"
    "log"
    "time"
)

// Request for updating a charge due date
type UpdateChargeDueDateRequest struct {
    DueAt time.Time `json:"due_at,omitempty"`
}

func (u *UpdateChargeDueDateRequest) MarshalJSON() (
    []byte,
    error) {
    type marshallable UpdateChargeDueDateRequest
    return json.Marshal(struct {
        marshallable 
        DueAt        string `json:"dueAt"` 
    }{
        DueAt:        u.DueAt.Format(time.RFC3339), 
        marshallable: marshallable(*u),             
    })
}

func (u *UpdateChargeDueDateRequest) UnmarshalJSON(input []byte) error {
    temp := &struct {
        DueAt string 
    }{}
    err := json.Unmarshal(input, temp)
    if err != nil {
    	return err
    }
    
    u.DueAt, err = time.Parse(time.RFC3339, temp.DueAt)
    if err != nil {
        log.Fatalf("Cannot Parse due_at as % s format.", time.RFC3339)
    }
    return nil
}
