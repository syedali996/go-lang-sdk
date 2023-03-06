package models

import (
    "encoding/json"
    "log"
    "time"
)

type GetTransactionReportFileResponse struct {
    Name *string    `json:"name"`
    Date *time.Time `json:"date"`
}

func (g *GetTransactionReportFileResponse) MarshalJSON() (
    []byte,
    error) {
    var DateVal *string
    if g.Date != nil {
        str := g.Date.Format(time.RFC3339)
        DateVal = &str
    } else {
        DateVal = nil
    }
    type marshallable GetTransactionReportFileResponse
    return json.Marshal(struct {
        marshallable 
        Date         *string `json:"date"` 
    }{
        Date:         DateVal,          
        marshallable: marshallable(*g), 
    })
}

func (g *GetTransactionReportFileResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Name *string 
        Date *string 
    }{}
    err := json.Unmarshal(input, temp)
    if err != nil {
    	return err
    }
    
    g.Name = temp.Name
    if temp.Date != nil {
        DateVal, err := time.Parse(time.RFC3339, *temp.Date)
        if err != nil {
            log.Fatalf("Cannot Parse date as % s format.", time.RFC3339)
        }
        g.Date = &DateVal
    } else {
        g.Date = nil
    }
    return nil
}
