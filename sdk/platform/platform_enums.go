package platform

import (
    "errors"
)





    //PriorityEnum used by Lead
    type PriorityEnum  string
    
    const (
            
            //LOW defines constant for the `low` 
            LOW PriorityEnum = "low"
            
            
            //MEDIUM defines constant for the `medium` 
            MEDIUM PriorityEnum = "medium"
            
            
            //HIGH defines constant for the `high` 
            HIGH PriorityEnum = "high"
            
            
       )

    //IsValid return error if enum is invalid
    func (pr PriorityEnum) IsValid() error {
      switch pr {
        case  LOW, MEDIUM, HIGH : 
        return nil
      }
      return errors.New("Invalid PriorityEnum type")
    }






















