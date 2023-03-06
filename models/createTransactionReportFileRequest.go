package models

import (
    "encoding/json"
    "log"
    "time"
)

type CreateTransactionReportFileRequest struct {
    Name    string    `json:"name"`
    StartAt time.Time `json:"start_at,omitempty"`
    EndAt   string    `json:"end_at,omitempty"`
}

func (c *CreateTransactionReportFileRequest) MarshalJSON() (
    []byte,
    error) {
    type marshallable CreateTransactionReportFileRequest
    return json.Marshal(struct {
        marshallable 
        StartAt      string `json:"startAt"` 
    }{
        StartAt:      c.StartAt.Format(time.RFC3339), 
        marshallable: marshallable(*c),               
    })
}

func (c *CreateTransactionReportFileRequest) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Name    string 
        StartAt string 
        EndAt   string 
    }{}
    err := json.Unmarshal(input, temp)
    if err != nil {
    	return err
    }
    
    c.Name = temp.Name
    c.StartAt, err = time.Parse(time.RFC3339, temp.StartAt)
    if err != nil {
        log.Fatalf("Cannot Parse start_at as % s format.", time.RFC3339)
    }
    c.EndAt = temp.EndAt
    return nil
}
