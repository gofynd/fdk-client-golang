package application

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
            
            
            //URGENT defines constant for the `urgent` 
            URGENT PriorityEnum = "urgent"
            
            
       )

    //IsValid return error if enum is invalid
    func (pr PriorityEnum) IsValid() error {
      switch pr {
        case  LOW, MEDIUM, HIGH, URGENT : 
        return nil
      }
      return errors.New("Invalid PriorityEnum type")
    }



    //HistoryTypeEnum used by Lead
    type HistoryTypeEnum  string
    
    const (
            
            //RATING defines constant for the `rating` 
            RATING HistoryTypeEnum = "rating"
            
            
            //LOG defines constant for the `log` 
            LOG HistoryTypeEnum = "log"
            
            
            //COMMENT defines constant for the `comment` 
            COMMENT HistoryTypeEnum = "comment"
            
            
       )

    //IsValid return error if enum is invalid
    func (hi HistoryTypeEnum) IsValid() error {
      switch hi {
        case  RATING, LOG, COMMENT : 
        return nil
      }
      return errors.New("Invalid HistoryTypeEnum type")
    }



    //TicketAssetType used by Lead
    type TicketAssetType  string
    
    const (
            
            //IMAGE defines constant for the `image` 
            IMAGE TicketAssetType = "image"
            
            
            //VIDEO defines constant for the `video` 
            VIDEO TicketAssetType = "video"
            
            
            //FILE defines constant for the `file` 
            FILE TicketAssetType = "file"
            
            
            //YOUTUBE defines constant for the `youtube` 
            YOUTUBE TicketAssetType = "youtube"
            
            
            //PRODUCT defines constant for the `product` 
            PRODUCT TicketAssetType = "product"
            
            
            //COLLECTION defines constant for the `collection` 
            COLLECTION TicketAssetType = "collection"
            
            
            //BRAND defines constant for the `brand` 
            BRAND TicketAssetType = "brand"
            
            
            //SHIPMENT defines constant for the `shipment` 
            SHIPMENT TicketAssetType = "shipment"
            
            
            //ORDER defines constant for the `order` 
            ORDER TicketAssetType = "order"
            
            
       )

    //IsValid return error if enum is invalid
    func (ti TicketAssetType) IsValid() error {
      switch ti {
        case  IMAGE, VIDEO, FILE, YOUTUBE, PRODUCT, COLLECTION, BRAND, SHIPMENT, ORDER : 
        return nil
      }
      return errors.New("Invalid TicketAssetType type")
    }



    //TicketSourceEnum used by Lead
    type TicketSourceEnum  string
    
    const (
            
            //PLATFORM_PANEL defines constant for the `platform_panel` 
            PLATFORM_PANEL TicketSourceEnum = "platform_panel"
            
            
            //SALES_CHANNEL defines constant for the `sales_channel` 
            SALES_CHANNEL TicketSourceEnum = "sales_channel"
            
            
       )

    //IsValid return error if enum is invalid
    func (ti TicketSourceEnum) IsValid() error {
      switch ti {
        case  PLATFORM_PANEL, SALES_CHANNEL : 
        return nil
      }
      return errors.New("Invalid TicketSourceEnum type")
    }















