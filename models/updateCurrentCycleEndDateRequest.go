package models

import (
    "encoding/json"
    "log"
    "time"
)

// Request to update the end date of the current subscription cycle
type UpdateCurrentCycleEndDateRequest struct {
    EndAt time.Time `json:"end_at,omitempty"`
}

func (u *UpdateCurrentCycleEndDateRequest) MarshalJSON() (
    []byte,
    error) {
    type marshallable UpdateCurrentCycleEndDateRequest
    return json.Marshal(struct {
        marshallable 
        EndAt        string `json:"endAt"` 
    }{
        EndAt:        u.EndAt.Format(time.RFC3339), 
        marshallable: marshallable(*u),             
    })
}

func (u *UpdateCurrentCycleEndDateRequest) UnmarshalJSON(input []byte) error {
    temp := &struct {
        EndAt string 
    }{}
    err := json.Unmarshal(input, temp)
    if err != nil {
    	return err
    }
    
    u.EndAt, err = time.Parse(time.RFC3339, temp.EndAt)
    if err != nil {
        log.Fatalf("Cannot Parse end_at as % s format.", time.RFC3339)
    }
    return nil
}
