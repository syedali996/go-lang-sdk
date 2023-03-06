package models

import (
    "encoding/json"
    "log"
    "time"
)

type CreatePeriodRequest struct {
    EndAt time.Time `json:"end_at,omitempty"`
}

func (c *CreatePeriodRequest) MarshalJSON() (
    []byte,
    error) {
    type marshallable CreatePeriodRequest
    return json.Marshal(struct {
        marshallable 
        EndAt        string `json:"endAt"` 
    }{
        EndAt:        c.EndAt.Format(time.RFC3339), 
        marshallable: marshallable(*c),             
    })
}

func (c *CreatePeriodRequest) UnmarshalJSON(input []byte) error {
    temp := &struct {
        EndAt string 
    }{}
    err := json.Unmarshal(input, temp)
    if err != nil {
    	return err
    }
    
    c.EndAt, err = time.Parse(time.RFC3339, temp.EndAt)
    if err != nil {
        log.Fatalf("Cannot Parse end_at as % s format.", time.RFC3339)
    }
    return nil
}
