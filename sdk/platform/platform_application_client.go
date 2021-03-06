package platform



import (
	"encoding/json"
	"fmt"
	"github.com/gofynd/fdk-client-golang/sdk/common"
)
// ApplicationClient holds all the platform-application level services instance 
type ApplicationClient struct {
	config        *PlatformConfig
	CompanyID string
	ApplicationID string
	
		Lead  *PlatformAppLead
	 
		Feedback  *PlatformAppFeedback
	 
		Theme  *PlatformAppTheme
	 
		User  *PlatformAppUser
	 
		Content  *PlatformAppContent
	 
		Assignment  *PlatformAppAssignment
	 
		Communication  *PlatformAppCommunication
	 
		Payment  *PlatformAppPayment
	 
		Order  *PlatformAppOrder
	 
		Catalog  *PlatformAppCatalog
	 
		FileStorage  *PlatformAppFileStorage
	 
		Share  *PlatformAppShare
	 
		Configuration  *PlatformAppConfiguration
	 
		Cart  *PlatformAppCart
	 
		Rewards  *PlatformAppRewards
	 
		Analytics  *PlatformAppAnalytics
	 
		Partner  *PlatformAppPartner
	 
}

// NewApplicationClient provides platform-application client
func NewApplicationClient(appID string, config *PlatformConfig) *ApplicationClient {
		return &ApplicationClient{
			config: config,
			CompanyID: config.CompanyID,
			ApplicationID: appID,
			
				Lead:  NewPlatformAppLead(config, appID),
			 
				Feedback:  NewPlatformAppFeedback(config, appID),
			 
				Theme:  NewPlatformAppTheme(config, appID),
			 
				User:  NewPlatformAppUser(config, appID),
			 
				Content:  NewPlatformAppContent(config, appID),
			 
				Assignment:  NewPlatformAppAssignment(config, appID),
			 
				Communication:  NewPlatformAppCommunication(config, appID),
			 
				Payment:  NewPlatformAppPayment(config, appID),
			 
				Order:  NewPlatformAppOrder(config, appID),
			 
				Catalog:  NewPlatformAppCatalog(config, appID),
			 
				FileStorage:  NewPlatformAppFileStorage(config, appID),
			 
				Share:  NewPlatformAppShare(config, appID),
			 
				Configuration:  NewPlatformAppConfiguration(config, appID),
			 
				Cart:  NewPlatformAppCart(config, appID),
			 
				Rewards:  NewPlatformAppRewards(config, appID),
			 
				Analytics:  NewPlatformAppAnalytics(config, appID),
			 
				Partner:  NewPlatformAppPartner(config, appID),
			 
		}
}


	 
   // PlatformAppLead holds PlatformAppLead object properties
    type PlatformAppLead struct {
        config *PlatformConfig
        CompanyID string
        ApplicationID string
    }
    // NewPlatformAppLead returns new PlatformAppLead instance
    func NewPlatformAppLead(config *PlatformConfig, appID string) *PlatformAppLead {
        return &PlatformAppLead{config, config.CompanyID, appID}
    }
    
    
    
    
    
    
    
  

    
    //PlatformAppGetTicketsXQuery holds query params
    type PlatformAppGetTicketsXQuery struct { 
        Items bool  `url:"items,omitempty"` 
        Filters bool  `url:"filters,omitempty"` 
        Q string  `url:"q,omitempty"` 
        Status string  `url:"status,omitempty"` 
        Priority PriorityEnum  `url:"priority,omitempty"` 
        Category string  `url:"category,omitempty"`  
    }
    
    // GetTickets Gets the list of Application level Tickets and/or ticket filters depending on query params
     func (le *PlatformAppLead)  GetTickets(xQuery PlatformAppGetTicketsXQuery) (TicketList, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getTicketsResponse TicketList
	    )

        

         
            
                
            
                
            
                
            
                
            
                
                //enum validation inside query params
                err = xQuery.Priority.IsValid()
                if err != nil {
                    
                    return TicketList{}, common.NewFDKError(err.Error())
                }
                
            
                
            
        

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            le.config,
            "get",
            fmt.Sprintf("/service/platform/lead/v1.0/company/%s/application/%s/ticket",le.CompanyID, le.ApplicationID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return TicketList{}, err
	    }
        
        err = json.Unmarshal(response, &getTicketsResponse)
        if err != nil {
            return TicketList{}, common.NewFDKError(err.Error())
        }
        return getTicketsResponse, nil
        
    }
           
       
    
    
    
    
    
    
    
  

    
    // GetTicket Retreives ticket details of a application level ticket
     func (le *PlatformAppLead)  GetTicket(ID string) (Ticket, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getTicketResponse Ticket
	    )

        

         

        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            le.config,
            "get",
            fmt.Sprintf("/service/platform/lead/v1.0/company/%s/application/%s/ticket/%s",le.CompanyID, le.ApplicationID, ID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return Ticket{}, err
	    }
        
        err = json.Unmarshal(response, &getTicketResponse)
        if err != nil {
            return Ticket{}, common.NewFDKError(err.Error())
        }
        return getTicketResponse, nil
        
    }
           
       
    
    
    
  

    
    // EditTicket Edits ticket details of a application level ticket
     func (le *PlatformAppLead)  EditTicket(ID string, body  EditTicketPayload) (Ticket, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            editTicketResponse Ticket
	    )

        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        

         

        
        
        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return Ticket{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return Ticket{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            le.config,
            "put",
            fmt.Sprintf("/service/platform/lead/v1.0/company/%s/application/%s/ticket/%s",le.CompanyID, le.ApplicationID, ID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return Ticket{}, err
	    }
        
        err = json.Unmarshal(response, &editTicketResponse)
        if err != nil {
            return Ticket{}, common.NewFDKError(err.Error())
        }
        return editTicketResponse, nil
        
    }
           
       
    
    
    
    
    
    
    
    
    
    
    
  

    
    // CreateHistory Create history for specific application level ticket
     func (le *PlatformAppLead)  CreateHistory(ID string, body  TicketHistoryPayload) (TicketHistory, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            createHistoryResponse TicketHistory
	    )

        
            
        
            
        

         

        
        
        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return TicketHistory{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return TicketHistory{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            le.config,
            "post",
            fmt.Sprintf("/service/platform/lead/v1.0/company/%s/application/%s/ticket/%s/history",le.CompanyID, le.ApplicationID, ID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return TicketHistory{}, err
	    }
        
        err = json.Unmarshal(response, &createHistoryResponse)
        if err != nil {
            return TicketHistory{}, common.NewFDKError(err.Error())
        }
        return createHistoryResponse, nil
        
    }
           
       
    
    
    
  

    
    // GetTicketHistory Gets history list for specific application level ticket
     func (le *PlatformAppLead)  GetTicketHistory(ID string) (TicketHistoryList, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getTicketHistoryResponse TicketHistoryList
	    )

        

         

        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            le.config,
            "get",
            fmt.Sprintf("/service/platform/lead/v1.0/company/%s/application/%s/ticket/%s/history",le.CompanyID, le.ApplicationID, ID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return TicketHistoryList{}, err
	    }
        
        err = json.Unmarshal(response, &getTicketHistoryResponse)
        if err != nil {
            return TicketHistoryList{}, common.NewFDKError(err.Error())
        }
        return getTicketHistoryResponse, nil
        
    }
           
       
    
    
    
  

    
    // GetCustomForm Get specific custom form using it's slug
     func (le *PlatformAppLead)  GetCustomForm(Slug string) (CustomForm, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getCustomFormResponse CustomForm
	    )

        

         

        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            le.config,
            "get",
            fmt.Sprintf("/service/platform/lead/v1.0/company/%s/application/%s/form/%s",le.CompanyID, le.ApplicationID, Slug),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return CustomForm{}, err
	    }
        
        err = json.Unmarshal(response, &getCustomFormResponse)
        if err != nil {
            return CustomForm{}, common.NewFDKError(err.Error())
        }
        return getCustomFormResponse, nil
        
    }
           
       
    
    
    
  

    
    // EditCustomForm Edit the given custom form
     func (le *PlatformAppLead)  EditCustomForm(Slug string, body  EditCustomFormPayload) (CustomForm, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            editCustomFormResponse CustomForm
	    )

        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        

         

        
        
        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return CustomForm{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return CustomForm{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            le.config,
            "put",
            fmt.Sprintf("/service/platform/lead/v1.0/company/%s/application/%s/form/%s",le.CompanyID, le.ApplicationID, Slug),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return CustomForm{}, err
	    }
        
        err = json.Unmarshal(response, &editCustomFormResponse)
        if err != nil {
            return CustomForm{}, common.NewFDKError(err.Error())
        }
        return editCustomFormResponse, nil
        
    }
           
       
    
    
    
  

    
    // GetCustomForms Get list of custom form
     func (le *PlatformAppLead)  GetCustomForms() (CustomFormList, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getCustomFormsResponse CustomFormList
	    )

        

         

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            le.config,
            "get",
            fmt.Sprintf("/service/platform/lead/v1.0/company/%s/application/%s/form",le.CompanyID, le.ApplicationID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return CustomFormList{}, err
	    }
        
        err = json.Unmarshal(response, &getCustomFormsResponse)
        if err != nil {
            return CustomFormList{}, common.NewFDKError(err.Error())
        }
        return getCustomFormsResponse, nil
        
    }
           
       
    
    
    
  

    
    // CreateCustomForm Creates a new custom form
     func (le *PlatformAppLead)  CreateCustomForm(body  CreateCustomFormPayload) (CustomForm, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            createCustomFormResponse CustomForm
	    )

        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        

         

        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return CustomForm{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return CustomForm{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            le.config,
            "post",
            fmt.Sprintf("/service/platform/lead/v1.0/company/%s/application/%s/form",le.CompanyID, le.ApplicationID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return CustomForm{}, err
	    }
        
        err = json.Unmarshal(response, &createCustomFormResponse)
        if err != nil {
            return CustomForm{}, common.NewFDKError(err.Error())
        }
        return createCustomFormResponse, nil
        
    }
           
       
    
    
    
    
    
  

    
    // GetTokenForVideoRoom Get Token to join a specific Video Room using it's unqiue name
     func (le *PlatformAppLead)  GetTokenForVideoRoom(UniqueName string) (GetTokenForVideoRoomResponse, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getTokenForVideoRoomResponse GetTokenForVideoRoomResponse
	    )

        

         

        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            le.config,
            "get",
            fmt.Sprintf("/service/platform/lead/v1.0/company/%s/application/%s/video/room/%s/token",le.CompanyID, le.ApplicationID, UniqueName),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return GetTokenForVideoRoomResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getTokenForVideoRoomResponse)
        if err != nil {
            return GetTokenForVideoRoomResponse{}, common.NewFDKError(err.Error())
        }
        return getTokenForVideoRoomResponse, nil
        
    }
           
       
    
    
    
    
    
  

    
    // GetVideoParticipants Get participants of a specific Video Room using it's unique name
     func (le *PlatformAppLead)  GetVideoParticipants(UniqueName string) (GetParticipantsInsideVideoRoomResponse, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getVideoParticipantsResponse GetParticipantsInsideVideoRoomResponse
	    )

        

         

        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            le.config,
            "get",
            fmt.Sprintf("/service/platform/lead/v1.0/company/%s/application/%s/video/room/%s/participants",le.CompanyID, le.ApplicationID, UniqueName),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return GetParticipantsInsideVideoRoomResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getVideoParticipantsResponse)
        if err != nil {
            return GetParticipantsInsideVideoRoomResponse{}, common.NewFDKError(err.Error())
        }
        return getVideoParticipantsResponse, nil
        
    }
           
       
    
    
    
  

    
    // OpenVideoRoom Open a video room.
     func (le *PlatformAppLead)  OpenVideoRoom(body  CreateVideoRoomPayload) (CreateVideoRoomResponse, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            openVideoRoomResponse CreateVideoRoomResponse
	    )

        
            
        
            
        

         

        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return CreateVideoRoomResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return CreateVideoRoomResponse{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            le.config,
            "post",
            fmt.Sprintf("/service/platform/lead/v1.0/company/%s/application/%s/video/room",le.CompanyID, le.ApplicationID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return CreateVideoRoomResponse{}, err
	    }
        
        err = json.Unmarshal(response, &openVideoRoomResponse)
        if err != nil {
            return CreateVideoRoomResponse{}, common.NewFDKError(err.Error())
        }
        return openVideoRoomResponse, nil
        
    }
           
       
    
    
    
  

    
    // CloseVideoRoom Close the video room and force all participants to leave.
     func (le *PlatformAppLead)  CloseVideoRoom(UniqueName string) (CloseVideoRoomResponse, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            closeVideoRoomResponse CloseVideoRoomResponse
	    )

        

         

        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            le.config,
            "delete",
            fmt.Sprintf("/service/platform/lead/v1.0/company/%s/application/%s/video/room/%s",le.CompanyID, le.ApplicationID, UniqueName),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return CloseVideoRoomResponse{}, err
	    }
        
        err = json.Unmarshal(response, &closeVideoRoomResponse)
        if err != nil {
            return CloseVideoRoomResponse{}, common.NewFDKError(err.Error())
        }
        return closeVideoRoomResponse, nil
        
    }
           
       
    

 
	 
   // PlatformAppFeedback holds PlatformAppFeedback object properties
    type PlatformAppFeedback struct {
        config *PlatformConfig
        CompanyID string
        ApplicationID string
    }
    // NewPlatformAppFeedback returns new PlatformAppFeedback instance
    func NewPlatformAppFeedback(config *PlatformConfig, appID string) *PlatformAppFeedback {
        return &PlatformAppFeedback{config, config.CompanyID, appID}
    }
    
    
    
  

    
    //PlatformAppGetAttributesXQuery holds query params
    type PlatformAppGetAttributesXQuery struct { 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"`  
    }
    
    // GetAttributes Get list of attribute data
     func (fe *PlatformAppFeedback)  GetAttributes(xQuery PlatformAppGetAttributesXQuery) (FeedbackAttributes, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getAttributesResponse FeedbackAttributes
	    )

        

         
            
                
            
                
            
        

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            fe.config,
            "get",
            fmt.Sprintf("/service/platform/feedback/v1.0/company/%s/application/%s/attributes/",fe.CompanyID, fe.ApplicationID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return FeedbackAttributes{}, err
	    }
        
        err = json.Unmarshal(response, &getAttributesResponse)
        if err != nil {
            return FeedbackAttributes{}, common.NewFDKError(err.Error())
        }
        return getAttributesResponse, nil
        
    }
           
            
            
            
            
            
            
            
             
            
            
             
            
            
             
            
            
             
            
            
            
            
            
            
            
            
            
            // GetAttributesPaginator Get list of attribute data  
            func (fe *PlatformAppFeedback)  GetAttributesPaginator( xQuery PlatformAppGetAttributesXQuery ) *common.Paginator {
                paginator := common.NewPaginator("number")
                
                
                 xQuery.PageNo  = paginator.PageNo
                 
                 
                 
                 
                 
                 
                 
                 
                paginator.Next = func() (interface{}, error) {
                    response, err := fe.GetAttributes(xQuery)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, int(response.Page.Current+1), response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }
        
       
    
    
    
  

    
    //PlatformAppGetCustomerReviewsXQuery holds query params
    type PlatformAppGetCustomerReviewsXQuery struct { 
        ID string  `url:"id,omitempty"` 
        EntityID string  `url:"entity_id,omitempty"` 
        EntityType string  `url:"entity_type,omitempty"` 
        UserID string  `url:"user_id,omitempty"` 
        Media string  `url:"media,omitempty"` 
        Rating []float64  `url:"rating,omitempty"` 
        AttributeRating []string  `url:"attribute_rating,omitempty"` 
        Facets bool  `url:"facets,omitempty"` 
        Sort string  `url:"sort,omitempty"` 
        Next string  `url:"next,omitempty"` 
        Start string  `url:"start,omitempty"` 
        Limit string  `url:"limit,omitempty"` 
        Count string  `url:"count,omitempty"` 
        PageID string  `url:"page_id,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"`  
    }
    
    // GetCustomerReviews Get list of customer reviews [admin]
     func (fe *PlatformAppFeedback)  GetCustomerReviews(xQuery PlatformAppGetCustomerReviewsXQuery) (GetReviewResponse, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getCustomerReviewsResponse GetReviewResponse
	    )

        

         
            
                
            
                
            
                
            
                
            
                
            
                
            
                
            
                
            
                
            
                
            
                
            
                
            
                
            
                
            
                
            
        

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            fe.config,
            "get",
            fmt.Sprintf("/service/platform/feedback/v1.0/company/%s/application/%s/reviews/",fe.CompanyID, fe.ApplicationID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return GetReviewResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getCustomerReviewsResponse)
        if err != nil {
            return GetReviewResponse{}, common.NewFDKError(err.Error())
        }
        return getCustomerReviewsResponse, nil
        
    }
           
            
            
            
            
            
            
            
             
            
            
             
            
            
             
            
            
             
            
            
             
            
            
             
            
            
             
            
            
             
            
            
             
            
            
             
            
            
             
            
            
             
            
            
             
            
            
             
            
            
             
            
            
             
            
            
             
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            // GetCustomerReviewsPaginator Get list of customer reviews [admin]  
            func (fe *PlatformAppFeedback)  GetCustomerReviewsPaginator( xQuery PlatformAppGetCustomerReviewsXQuery ) *common.Paginator {
                paginator := common.NewPaginator("cursor")
                
                
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 xQuery.PageID = paginator.NextID
                 
                 
                 
                 
                 
                 
                 
                paginator.Next = func() (interface{}, error) {
                    response, err := fe.GetCustomerReviews(xQuery)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, int(response.Page.Current+1), response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }
        
       
    
    
    
  

    
    // UpdateApprove update approve details
     func (fe *PlatformAppFeedback)  UpdateApprove(ReviewID string, body  ApproveRequest) (UpdateResponse, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            updateApproveResponse UpdateResponse
	    )

        
            
        
            
        
            
        
            
        

         

        
        
        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return UpdateResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return UpdateResponse{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            fe.config,
            "put",
            fmt.Sprintf("/service/platform/feedback/v1.0/company/%s/application/%s/reviews/%s/approve/",fe.CompanyID, fe.ApplicationID, ReviewID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return UpdateResponse{}, err
	    }
        
        err = json.Unmarshal(response, &updateApproveResponse)
        if err != nil {
            return UpdateResponse{}, common.NewFDKError(err.Error())
        }
        return updateApproveResponse, nil
        
    }
           
       
    
    
    
  

    
    // GetHistory get history details
     func (fe *PlatformAppFeedback)  GetHistory(ReviewID string) ([]ActivityDump, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getHistoryResponse []ActivityDump
	    )

        

         

        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            fe.config,
            "get",
            fmt.Sprintf("/service/platform/feedback/v1.0/company/%s/application/%s/reviews/%s/history/",fe.CompanyID, fe.ApplicationID, ReviewID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return []ActivityDump{}, err
	    }
        
        err = json.Unmarshal(response, &getHistoryResponse)
        if err != nil {
            return []ActivityDump{}, common.NewFDKError(err.Error())
        }
        return getHistoryResponse, nil
        
    }
           
       
    
    
    
  

    
    //PlatformAppGetApplicationTemplatesXQuery holds query params
    type PlatformAppGetApplicationTemplatesXQuery struct { 
        PageID string  `url:"page_id,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"`  
    }
    
    // GetApplicationTemplates Get list of templates
     func (fe *PlatformAppFeedback)  GetApplicationTemplates(xQuery PlatformAppGetApplicationTemplatesXQuery) (TemplateGetResponse, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getApplicationTemplatesResponse TemplateGetResponse
	    )

        

         
            
                
            
                
            
        

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            fe.config,
            "get",
            fmt.Sprintf("/service/platform/feedback/v1.0/company/%s/application/%s/templates/",fe.CompanyID, fe.ApplicationID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return TemplateGetResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getApplicationTemplatesResponse)
        if err != nil {
            return TemplateGetResponse{}, common.NewFDKError(err.Error())
        }
        return getApplicationTemplatesResponse, nil
        
    }
           
            
            
            
            
            
            
            
             
            
            
             
            
            
             
            
            
             
            
            
            
            
            
            
            
            
            
            // GetApplicationTemplatesPaginator Get list of templates  
            func (fe *PlatformAppFeedback)  GetApplicationTemplatesPaginator( xQuery PlatformAppGetApplicationTemplatesXQuery ) *common.Paginator {
                paginator := common.NewPaginator("cursor")
                
                
                 
                 xQuery.PageID = paginator.NextID
                 
                 
                 
                 
                 
                 
                 
                paginator.Next = func() (interface{}, error) {
                    response, err := fe.GetApplicationTemplates(xQuery)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, int(response.Page.Current+1), response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }
        
       
    
    
    
  

    
    // CreateTemplate Create a new template
     func (fe *PlatformAppFeedback)  CreateTemplate(body  TemplateRequestList) (InsertResponse, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            createTemplateResponse InsertResponse
	    )

        
            
        

         

        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return InsertResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return InsertResponse{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            fe.config,
            "post",
            fmt.Sprintf("/service/platform/feedback/v1.0/company/%s/application/%s/templates/",fe.CompanyID, fe.ApplicationID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return InsertResponse{}, err
	    }
        
        err = json.Unmarshal(response, &createTemplateResponse)
        if err != nil {
            return InsertResponse{}, common.NewFDKError(err.Error())
        }
        return createTemplateResponse, nil
        
    }
           
       
    
    
    
  

    
    // GetTemplateById Get a template by ID
     func (fe *PlatformAppFeedback)  GetTemplateById(ID string) (Template, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getTemplateByIdResponse Template
	    )

        

         

        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            fe.config,
            "get",
            fmt.Sprintf("/service/platform/feedback/v1.0/company/%s/application/%s/templates/%s/",fe.CompanyID, fe.ApplicationID, ID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return Template{}, err
	    }
        
        err = json.Unmarshal(response, &getTemplateByIdResponse)
        if err != nil {
            return Template{}, common.NewFDKError(err.Error())
        }
        return getTemplateByIdResponse, nil
        
    }
           
       
    
    
    
  

    
    // UpdateTemplate Update a template's status
     func (fe *PlatformAppFeedback)  UpdateTemplate(ID string, body  UpdateTemplateRequest) (UpdateResponse, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            updateTemplateResponse UpdateResponse
	    )

        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        

         

        
        
        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return UpdateResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return UpdateResponse{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            fe.config,
            "put",
            fmt.Sprintf("/service/platform/feedback/v1.0/company/%s/application/%s/templates/%s/",fe.CompanyID, fe.ApplicationID, ID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return UpdateResponse{}, err
	    }
        
        err = json.Unmarshal(response, &updateTemplateResponse)
        if err != nil {
            return UpdateResponse{}, common.NewFDKError(err.Error())
        }
        return updateTemplateResponse, nil
        
    }
           
       
    
    
    
  

    
    // UpdateTemplateStatus Update a template's status
     func (fe *PlatformAppFeedback)  UpdateTemplateStatus(ID string, body  UpdateTemplateStatusRequest) (UpdateResponse, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            updateTemplateStatusResponse UpdateResponse
	    )

        
            
        
            
        

         

        
        
        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return UpdateResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return UpdateResponse{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            fe.config,
            "patch",
            fmt.Sprintf("/service/platform/feedback/v1.0/company/%s/application/%s/templates/%s/status/",fe.CompanyID, fe.ApplicationID, ID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return UpdateResponse{}, err
	    }
        
        err = json.Unmarshal(response, &updateTemplateStatusResponse)
        if err != nil {
            return UpdateResponse{}, common.NewFDKError(err.Error())
        }
        return updateTemplateStatusResponse, nil
        
    }
           
       
    

 
	 
   // PlatformAppTheme holds PlatformAppTheme object properties
    type PlatformAppTheme struct {
        config *PlatformConfig
        CompanyID string
        ApplicationID string
    }
    // NewPlatformAppTheme returns new PlatformAppTheme instance
    func NewPlatformAppTheme(config *PlatformConfig, appID string) *PlatformAppTheme {
        return &PlatformAppTheme{config, config.CompanyID, appID}
    }
    
    
    
  

    
    // GetAllPages Get all pages of a theme
     func (th *PlatformAppTheme)  GetAllPages(ThemeID string) (AllAvailablePageSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getAllPagesResponse AllAvailablePageSchema
	    )

        

         

        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            th.config,
            "get",
            fmt.Sprintf("/service/platform/theme/v1.0/company/%s/application/%s/%s/page",th.CompanyID, th.ApplicationID, ThemeID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return AllAvailablePageSchema{}, err
	    }
        
        err = json.Unmarshal(response, &getAllPagesResponse)
        if err != nil {
            return AllAvailablePageSchema{}, common.NewFDKError(err.Error())
        }
        return getAllPagesResponse, nil
        
    }
           
       
    
    
    
  

    
    // CreatePage Create a page 
     func (th *PlatformAppTheme)  CreatePage(ThemeID string, body  AvailablePageSchema) (AvailablePageSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            createPageResponse AvailablePageSchema
	    )

        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        

         

        
        
        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return AvailablePageSchema{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return AvailablePageSchema{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            th.config,
            "post",
            fmt.Sprintf("/service/platform/theme/v1.0/company/%s/application/%s/%s/page",th.CompanyID, th.ApplicationID, ThemeID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return AvailablePageSchema{}, err
	    }
        
        err = json.Unmarshal(response, &createPageResponse)
        if err != nil {
            return AvailablePageSchema{}, common.NewFDKError(err.Error())
        }
        return createPageResponse, nil
        
    }
           
       
    
    
    
  

    
    // UpdateMultiplePages Update multiple pages of a theme
     func (th *PlatformAppTheme)  UpdateMultiplePages(ThemeID string, body  AllAvailablePageSchema) (AllAvailablePageSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            updateMultiplePagesResponse AllAvailablePageSchema
	    )

        
            
        

         

        
        
        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return AllAvailablePageSchema{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return AllAvailablePageSchema{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            th.config,
            "put",
            fmt.Sprintf("/service/platform/theme/v1.0/company/%s/application/%s/%s/page",th.CompanyID, th.ApplicationID, ThemeID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return AllAvailablePageSchema{}, err
	    }
        
        err = json.Unmarshal(response, &updateMultiplePagesResponse)
        if err != nil {
            return AllAvailablePageSchema{}, common.NewFDKError(err.Error())
        }
        return updateMultiplePagesResponse, nil
        
    }
           
       
    
    
    
  

    
    // GetPage Get page of a theme
     func (th *PlatformAppTheme)  GetPage(ThemeID string, PageValue string) (AvailablePageSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getPageResponse AvailablePageSchema
	    )

        

         

        
        
        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            th.config,
            "get",
            fmt.Sprintf("/service/platform/theme/v1.0/company/%s/application/%s/%s/%s",th.CompanyID, th.ApplicationID, ThemeID, PageValue),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return AvailablePageSchema{}, err
	    }
        
        err = json.Unmarshal(response, &getPageResponse)
        if err != nil {
            return AvailablePageSchema{}, common.NewFDKError(err.Error())
        }
        return getPageResponse, nil
        
    }
           
       
    
    
    
  

    
    // UpdatePage Updates a page 
     func (th *PlatformAppTheme)  UpdatePage(ThemeID string, PageValue string, body  AvailablePageSchema) (AvailablePageSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            updatePageResponse AvailablePageSchema
	    )

        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        

         

        
        
        
        
        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return AvailablePageSchema{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return AvailablePageSchema{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            th.config,
            "put",
            fmt.Sprintf("/service/platform/theme/v1.0/company/%s/application/%s/%s/%s",th.CompanyID, th.ApplicationID, ThemeID, PageValue),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return AvailablePageSchema{}, err
	    }
        
        err = json.Unmarshal(response, &updatePageResponse)
        if err != nil {
            return AvailablePageSchema{}, common.NewFDKError(err.Error())
        }
        return updatePageResponse, nil
        
    }
           
       
    
    
    
  

    
    // DeletePage Deletes a page 
     func (th *PlatformAppTheme)  DeletePage(ThemeID string, PageValue string) (AvailablePageSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            deletePageResponse AvailablePageSchema
	    )

        

         

        
        
        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            th.config,
            "delete",
            fmt.Sprintf("/service/platform/theme/v1.0/company/%s/application/%s/%s/%s",th.CompanyID, th.ApplicationID, ThemeID, PageValue),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return AvailablePageSchema{}, err
	    }
        
        err = json.Unmarshal(response, &deletePageResponse)
        if err != nil {
            return AvailablePageSchema{}, common.NewFDKError(err.Error())
        }
        return deletePageResponse, nil
        
    }
           
       
    
    
    
  

    
    //PlatformAppGetThemeLibraryXQuery holds query params
    type PlatformAppGetThemeLibraryXQuery struct { 
        PageSize float64  `url:"page_size,omitempty"` 
        PageNo float64  `url:"page_no,omitempty"`  
    }
    
    // GetThemeLibrary Get a list of themes from the theme library
     func (th *PlatformAppTheme)  GetThemeLibrary(xQuery PlatformAppGetThemeLibraryXQuery) (ThemesListingResponseSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getThemeLibraryResponse ThemesListingResponseSchema
	    )

        

         
            
                
            
                
            
        

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            th.config,
            "get",
            fmt.Sprintf("/service/platform/theme/v1.0/company/%s/application/%s/library",th.CompanyID, th.ApplicationID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return ThemesListingResponseSchema{}, err
	    }
        
        err = json.Unmarshal(response, &getThemeLibraryResponse)
        if err != nil {
            return ThemesListingResponseSchema{}, common.NewFDKError(err.Error())
        }
        return getThemeLibraryResponse, nil
        
    }
           
       
    
    
    
  

    
    // AddToThemeLibrary Add a theme to the theme library
     func (th *PlatformAppTheme)  AddToThemeLibrary(body  AddThemeRequestSchema) (ThemesSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            addToThemeLibraryResponse ThemesSchema
	    )

        
            
        

         

        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return ThemesSchema{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return ThemesSchema{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            th.config,
            "post",
            fmt.Sprintf("/service/platform/theme/v1.0/company/%s/application/%s/library",th.CompanyID, th.ApplicationID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return ThemesSchema{}, err
	    }
        
        err = json.Unmarshal(response, &addToThemeLibraryResponse)
        if err != nil {
            return ThemesSchema{}, common.NewFDKError(err.Error())
        }
        return addToThemeLibraryResponse, nil
        
    }
           
       
    
    
    
  

    
    // ApplyTheme Apply a theme
     func (th *PlatformAppTheme)  ApplyTheme(body  AddThemeRequestSchema) (ThemesSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            applyThemeResponse ThemesSchema
	    )

        
            
        

         

        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return ThemesSchema{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return ThemesSchema{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            th.config,
            "post",
            fmt.Sprintf("/service/platform/theme/v1.0/company/%s/application/%s/apply",th.CompanyID, th.ApplicationID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return ThemesSchema{}, err
	    }
        
        err = json.Unmarshal(response, &applyThemeResponse)
        if err != nil {
            return ThemesSchema{}, common.NewFDKError(err.Error())
        }
        return applyThemeResponse, nil
        
    }
           
       
    
    
    
  

    
    // IsUpgradable Checks if theme is upgradable
     func (th *PlatformAppTheme)  IsUpgradable(ThemeID string) (UpgradableThemeSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            isUpgradableResponse UpgradableThemeSchema
	    )

        

         

        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            th.config,
            "get",
            fmt.Sprintf("/service/platform/theme/v1.0/company/%s/application/%s/%s/upgradable",th.CompanyID, th.ApplicationID, ThemeID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return UpgradableThemeSchema{}, err
	    }
        
        err = json.Unmarshal(response, &isUpgradableResponse)
        if err != nil {
            return UpgradableThemeSchema{}, common.NewFDKError(err.Error())
        }
        return isUpgradableResponse, nil
        
    }
           
       
    
    
    
  

    
    // UpgradeTheme Upgrade a theme
     func (th *PlatformAppTheme)  UpgradeTheme(ThemeID string) (ThemesSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            upgradeThemeResponse ThemesSchema
	    )

        

         

        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            th.config,
            "put",
            fmt.Sprintf("/service/platform/theme/v1.0/company/%s/application/%s/%s/upgrade",th.CompanyID, th.ApplicationID, ThemeID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return ThemesSchema{}, err
	    }
        
        err = json.Unmarshal(response, &upgradeThemeResponse)
        if err != nil {
            return ThemesSchema{}, common.NewFDKError(err.Error())
        }
        return upgradeThemeResponse, nil
        
    }
           
       
    
    
    
  

    
    //PlatformAppGetPublicThemesXQuery holds query params
    type PlatformAppGetPublicThemesXQuery struct { 
        PageSize float64  `url:"page_size,omitempty"` 
        PageNo float64  `url:"page_no,omitempty"`  
    }
    
    // GetPublicThemes Get all public themes
     func (th *PlatformAppTheme)  GetPublicThemes(xQuery PlatformAppGetPublicThemesXQuery) (ThemesListingResponseSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getPublicThemesResponse ThemesListingResponseSchema
	    )

        

         
            
                
            
                
            
        

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            th.config,
            "get",
            fmt.Sprintf("/service/platform/theme/v1.0/company/%s/application/%s/list/public",th.CompanyID, th.ApplicationID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return ThemesListingResponseSchema{}, err
	    }
        
        err = json.Unmarshal(response, &getPublicThemesResponse)
        if err != nil {
            return ThemesListingResponseSchema{}, common.NewFDKError(err.Error())
        }
        return getPublicThemesResponse, nil
        
    }
           
       
    
    
    
  

    
    // CreateTheme Create a new theme
     func (th *PlatformAppTheme)  CreateTheme(body  ThemesSchema) (ThemesSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            createThemeResponse ThemesSchema
	    )

        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        

         

        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return ThemesSchema{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return ThemesSchema{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            th.config,
            "post",
            fmt.Sprintf("/service/platform/theme/v1.0/company/%s/application/%s/",th.CompanyID, th.ApplicationID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return ThemesSchema{}, err
	    }
        
        err = json.Unmarshal(response, &createThemeResponse)
        if err != nil {
            return ThemesSchema{}, common.NewFDKError(err.Error())
        }
        return createThemeResponse, nil
        
    }
           
       
    
    
    
  

    
    // GetAppliedTheme Get the applied theme
     func (th *PlatformAppTheme)  GetAppliedTheme() (ThemesSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getAppliedThemeResponse ThemesSchema
	    )

        

         

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            th.config,
            "put",
            fmt.Sprintf("/service/platform/theme/v1.0/company/%s/application/%s/",th.CompanyID, th.ApplicationID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return ThemesSchema{}, err
	    }
        
        err = json.Unmarshal(response, &getAppliedThemeResponse)
        if err != nil {
            return ThemesSchema{}, common.NewFDKError(err.Error())
        }
        return getAppliedThemeResponse, nil
        
    }
           
       
    
    
    
  

    
    // GetFonts Get all the supported fonts in a theme
     func (th *PlatformAppTheme)  GetFonts() (FontsSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getFontsResponse FontsSchema
	    )

        

         

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            th.config,
            "get",
            fmt.Sprintf("/service/platform/theme/v1.0/company/%s/application/%s/fonts",th.CompanyID, th.ApplicationID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return FontsSchema{}, err
	    }
        
        err = json.Unmarshal(response, &getFontsResponse)
        if err != nil {
            return FontsSchema{}, common.NewFDKError(err.Error())
        }
        return getFontsResponse, nil
        
    }
           
       
    
    
    
  

    
    // GetThemeById Gets theme by id
     func (th *PlatformAppTheme)  GetThemeById(ThemeID string) (ThemesSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getThemeByIdResponse ThemesSchema
	    )

        

         

        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            th.config,
            "get",
            fmt.Sprintf("/service/platform/theme/v1.0/company/%s/application/%s/%s",th.CompanyID, th.ApplicationID, ThemeID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return ThemesSchema{}, err
	    }
        
        err = json.Unmarshal(response, &getThemeByIdResponse)
        if err != nil {
            return ThemesSchema{}, common.NewFDKError(err.Error())
        }
        return getThemeByIdResponse, nil
        
    }
           
       
    
    
    
  

    
    // UpdateTheme Update a theme
     func (th *PlatformAppTheme)  UpdateTheme(ThemeID string, body  ThemesSchema) (ThemesSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            updateThemeResponse ThemesSchema
	    )

        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        

         

        
        
        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return ThemesSchema{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return ThemesSchema{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            th.config,
            "put",
            fmt.Sprintf("/service/platform/theme/v1.0/company/%s/application/%s/%s",th.CompanyID, th.ApplicationID, ThemeID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return ThemesSchema{}, err
	    }
        
        err = json.Unmarshal(response, &updateThemeResponse)
        if err != nil {
            return ThemesSchema{}, common.NewFDKError(err.Error())
        }
        return updateThemeResponse, nil
        
    }
           
       
    
    
    
  

    
    // DeleteTheme Delete a theme
     func (th *PlatformAppTheme)  DeleteTheme(ThemeID string) (ThemesSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            deleteThemeResponse ThemesSchema
	    )

        

         

        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            th.config,
            "delete",
            fmt.Sprintf("/service/platform/theme/v1.0/company/%s/application/%s/%s",th.CompanyID, th.ApplicationID, ThemeID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return ThemesSchema{}, err
	    }
        
        err = json.Unmarshal(response, &deleteThemeResponse)
        if err != nil {
            return ThemesSchema{}, common.NewFDKError(err.Error())
        }
        return deleteThemeResponse, nil
        
    }
           
       
    
    
    
  

    
    // GetThemeForPreview Get a theme preview
     func (th *PlatformAppTheme)  GetThemeForPreview(ThemeID string) (ThemesSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getThemeForPreviewResponse ThemesSchema
	    )

        

         

        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            th.config,
            "get",
            fmt.Sprintf("/service/platform/theme/v1.0/company/%s/application/%s/%s/preview",th.CompanyID, th.ApplicationID, ThemeID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return ThemesSchema{}, err
	    }
        
        err = json.Unmarshal(response, &getThemeForPreviewResponse)
        if err != nil {
            return ThemesSchema{}, common.NewFDKError(err.Error())
        }
        return getThemeForPreviewResponse, nil
        
    }
           
       
    
    
    
  

    
    // PublishTheme Publish a theme
     func (th *PlatformAppTheme)  PublishTheme(ThemeID string) (ThemesSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            publishThemeResponse ThemesSchema
	    )

        

         

        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            th.config,
            "put",
            fmt.Sprintf("/service/platform/theme/v1.0/company/%s/application/%s/%s/publish",th.CompanyID, th.ApplicationID, ThemeID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return ThemesSchema{}, err
	    }
        
        err = json.Unmarshal(response, &publishThemeResponse)
        if err != nil {
            return ThemesSchema{}, common.NewFDKError(err.Error())
        }
        return publishThemeResponse, nil
        
    }
           
       
    
    
    
  

    
    // UnpublishTheme Unpublish a theme
     func (th *PlatformAppTheme)  UnpublishTheme(ThemeID string) (ThemesSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            unpublishThemeResponse ThemesSchema
	    )

        

         

        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            th.config,
            "put",
            fmt.Sprintf("/service/platform/theme/v1.0/company/%s/application/%s/%s/unpublish",th.CompanyID, th.ApplicationID, ThemeID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return ThemesSchema{}, err
	    }
        
        err = json.Unmarshal(response, &unpublishThemeResponse)
        if err != nil {
            return ThemesSchema{}, common.NewFDKError(err.Error())
        }
        return unpublishThemeResponse, nil
        
    }
           
       
    
    
    
  

    
    // ArchiveTheme Archive a theme
     func (th *PlatformAppTheme)  ArchiveTheme(ThemeID string) (ThemesSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            archiveThemeResponse ThemesSchema
	    )

        

         

        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            th.config,
            "put",
            fmt.Sprintf("/service/platform/theme/v1.0/company/%s/application/%s/%s/archive",th.CompanyID, th.ApplicationID, ThemeID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return ThemesSchema{}, err
	    }
        
        err = json.Unmarshal(response, &archiveThemeResponse)
        if err != nil {
            return ThemesSchema{}, common.NewFDKError(err.Error())
        }
        return archiveThemeResponse, nil
        
    }
           
       
    
    
    
  

    
    // UnarchiveTheme Unarchive a theme
     func (th *PlatformAppTheme)  UnarchiveTheme(ThemeID string) (ThemesSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            unarchiveThemeResponse ThemesSchema
	    )

        

         

        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            th.config,
            "put",
            fmt.Sprintf("/service/platform/theme/v1.0/company/%s/application/%s/%s/unarchive",th.CompanyID, th.ApplicationID, ThemeID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return ThemesSchema{}, err
	    }
        
        err = json.Unmarshal(response, &unarchiveThemeResponse)
        if err != nil {
            return ThemesSchema{}, common.NewFDKError(err.Error())
        }
        return unarchiveThemeResponse, nil
        
    }
           
       
    

 
	 
   // PlatformAppUser holds PlatformAppUser object properties
    type PlatformAppUser struct {
        config *PlatformConfig
        CompanyID string
        ApplicationID string
    }
    // NewPlatformAppUser returns new PlatformAppUser instance
    func NewPlatformAppUser(config *PlatformConfig, appID string) *PlatformAppUser {
        return &PlatformAppUser{config, config.CompanyID, appID}
    }
    
    
    
  

    
    //PlatformAppGetCustomersXQuery holds query params
    type PlatformAppGetCustomersXQuery struct { 
        Q string  `url:"q,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"` 
        PageNo float64  `url:"page_no,omitempty"`  
    }
    
    // GetCustomers Get a list of customers
     func (us *PlatformAppUser)  GetCustomers(xQuery PlatformAppGetCustomersXQuery) (CustomerListResponseSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getCustomersResponse CustomerListResponseSchema
	    )

        

         
            
                
            
                
            
                
            
        

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            us.config,
            "get",
            fmt.Sprintf("/service/platform/user/v1.0/company/%s/application/%s/customers/list",us.CompanyID, us.ApplicationID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return CustomerListResponseSchema{}, err
	    }
        
        err = json.Unmarshal(response, &getCustomersResponse)
        if err != nil {
            return CustomerListResponseSchema{}, common.NewFDKError(err.Error())
        }
        return getCustomersResponse, nil
        
    }
           
       
    
    
    
  

    
    //PlatformAppSearchUsersXQuery holds query params
    type PlatformAppSearchUsersXQuery struct { 
        Q string  `url:"q,omitempty"`  
    }
    
    // SearchUsers Search an existing user.
     func (us *PlatformAppUser)  SearchUsers(xQuery PlatformAppSearchUsersXQuery) (UserSearchResponseSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            searchUsersResponse UserSearchResponseSchema
	    )

        

         
            
                
            
        

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            us.config,
            "get",
            fmt.Sprintf("/service/platform/user/v1.0/company/%s/application/%s/customers/search",us.CompanyID, us.ApplicationID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return UserSearchResponseSchema{}, err
	    }
        
        err = json.Unmarshal(response, &searchUsersResponse)
        if err != nil {
            return UserSearchResponseSchema{}, common.NewFDKError(err.Error())
        }
        return searchUsersResponse, nil
        
    }
           
       
    
    
    
  

    
    // CreateUser Create user
     func (us *PlatformAppUser)  CreateUser(body  CreateUserRequestSchema) (CreateUserResponseSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            createUserResponse CreateUserResponseSchema
	    )

        
            
        
            
        
            
        
            
        
            
        
            
        
            
        

         

        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return CreateUserResponseSchema{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return CreateUserResponseSchema{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            us.config,
            "post",
            fmt.Sprintf("/service/platform/user/v1.0/company/%s/application/%s/customers",us.CompanyID, us.ApplicationID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return CreateUserResponseSchema{}, err
	    }
        
        err = json.Unmarshal(response, &createUserResponse)
        if err != nil {
            return CreateUserResponseSchema{}, common.NewFDKError(err.Error())
        }
        return createUserResponse, nil
        
    }
           
       
    
    
    
  

    
    // UpdateUser Update user
     func (us *PlatformAppUser)  UpdateUser(UserID string, body  UpdateUserRequestSchema) (CreateUserResponseSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            updateUserResponse CreateUserResponseSchema
	    )

        
            
        
            
        
            
        
            
        

         

        
        
        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return CreateUserResponseSchema{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return CreateUserResponseSchema{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            us.config,
            "put",
            fmt.Sprintf("/service/platform/user/v1.0/company/%s/application/%s/customers/%s",us.CompanyID, us.ApplicationID, UserID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return CreateUserResponseSchema{}, err
	    }
        
        err = json.Unmarshal(response, &updateUserResponse)
        if err != nil {
            return CreateUserResponseSchema{}, common.NewFDKError(err.Error())
        }
        return updateUserResponse, nil
        
    }
           
       
    
    
    
  

    
    // CreateUserSession Create user session
     func (us *PlatformAppUser)  CreateUserSession(body  CreateUserSessionRequestSchema) (CreateUserSessionResponseSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            createUserSessionResponse CreateUserSessionResponseSchema
	    )

        
            
        
            
        
            
        

         

        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return CreateUserSessionResponseSchema{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return CreateUserSessionResponseSchema{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            us.config,
            "post",
            fmt.Sprintf("/service/platform/user/v1.0/company/%s/application/%s/customers/session",us.CompanyID, us.ApplicationID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return CreateUserSessionResponseSchema{}, err
	    }
        
        err = json.Unmarshal(response, &createUserSessionResponse)
        if err != nil {
            return CreateUserSessionResponseSchema{}, common.NewFDKError(err.Error())
        }
        return createUserSessionResponse, nil
        
    }
           
       
    
    
    
  

    
    // GetPlatformConfig Get platform configurations
     func (us *PlatformAppUser)  GetPlatformConfig() (PlatformSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getPlatformConfigResponse PlatformSchema
	    )

        

         

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            us.config,
            "get",
            fmt.Sprintf("/service/platform/user/v1.0/company/%s/application/%s/platform/config",us.CompanyID, us.ApplicationID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return PlatformSchema{}, err
	    }
        
        err = json.Unmarshal(response, &getPlatformConfigResponse)
        if err != nil {
            return PlatformSchema{}, common.NewFDKError(err.Error())
        }
        return getPlatformConfigResponse, nil
        
    }
           
       
    
    
    
  

    
    // UpdatePlatformConfig Update platform configurations
     func (us *PlatformAppUser)  UpdatePlatformConfig(body  PlatformSchema) (PlatformSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            updatePlatformConfigResponse PlatformSchema
	    )

        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        

         

        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return PlatformSchema{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return PlatformSchema{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            us.config,
            "post",
            fmt.Sprintf("/service/platform/user/v1.0/company/%s/application/%s/platform/config",us.CompanyID, us.ApplicationID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return PlatformSchema{}, err
	    }
        
        err = json.Unmarshal(response, &updatePlatformConfigResponse)
        if err != nil {
            return PlatformSchema{}, common.NewFDKError(err.Error())
        }
        return updatePlatformConfigResponse, nil
        
    }
           
       
    

 
	 
   // PlatformAppContent holds PlatformAppContent object properties
    type PlatformAppContent struct {
        config *PlatformConfig
        CompanyID string
        ApplicationID string
    }
    // NewPlatformAppContent returns new PlatformAppContent instance
    func NewPlatformAppContent(config *PlatformConfig, appID string) *PlatformAppContent {
        return &PlatformAppContent{config, config.CompanyID, appID}
    }
    
    
    
  

    
    //PlatformAppGetAnnouncementsListXQuery holds query params
    type PlatformAppGetAnnouncementsListXQuery struct { 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"`  
    }
    
    // GetAnnouncementsList Get a list of announcements
     func (co *PlatformAppContent)  GetAnnouncementsList(xQuery PlatformAppGetAnnouncementsListXQuery) (GetAnnouncementListSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getAnnouncementsListResponse GetAnnouncementListSchema
	    )

        

         
            
                
            
                
            
        

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            fmt.Sprintf("/service/platform/content/v1.0/company/%s/application/%s/announcements",co.CompanyID, co.ApplicationID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return GetAnnouncementListSchema{}, err
	    }
        
        err = json.Unmarshal(response, &getAnnouncementsListResponse)
        if err != nil {
            return GetAnnouncementListSchema{}, common.NewFDKError(err.Error())
        }
        return getAnnouncementsListResponse, nil
        
    }
           
            
            
            
            
            
            
            
             
            
            
             
            
            
             
            
            
             
            
            
            
            
            
            
            
            
            
            // GetAnnouncementsListPaginator Get a list of announcements  
            func (co *PlatformAppContent)  GetAnnouncementsListPaginator( xQuery PlatformAppGetAnnouncementsListXQuery ) *common.Paginator {
                paginator := common.NewPaginator("number")
                
                
                 xQuery.PageNo  = paginator.PageNo
                 
                 
                 
                 
                 
                 
                 
                 
                paginator.Next = func() (interface{}, error) {
                    response, err := co.GetAnnouncementsList(xQuery)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, int(response.Page.Current+1), response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }
        
       
    
    
    
  

    
    // CreateAnnouncement Create an announcement
     func (co *PlatformAppContent)  CreateAnnouncement(body  AdminAnnouncementSchema) (CreateAnnouncementSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            createAnnouncementResponse CreateAnnouncementSchema
	    )

        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        

         

        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return CreateAnnouncementSchema{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return CreateAnnouncementSchema{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "post",
            fmt.Sprintf("/service/platform/content/v1.0/company/%s/application/%s/announcements",co.CompanyID, co.ApplicationID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return CreateAnnouncementSchema{}, err
	    }
        
        err = json.Unmarshal(response, &createAnnouncementResponse)
        if err != nil {
            return CreateAnnouncementSchema{}, common.NewFDKError(err.Error())
        }
        return createAnnouncementResponse, nil
        
    }
           
       
    
    
    
  

    
    // GetAnnouncementById Get announcement by ID
     func (co *PlatformAppContent)  GetAnnouncementById(AnnouncementID string) (AdminAnnouncementSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getAnnouncementByIdResponse AdminAnnouncementSchema
	    )

        

         

        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            fmt.Sprintf("/service/platform/content/v1.0/company/%s/application/%s/announcements/%s",co.CompanyID, co.ApplicationID, AnnouncementID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return AdminAnnouncementSchema{}, err
	    }
        
        err = json.Unmarshal(response, &getAnnouncementByIdResponse)
        if err != nil {
            return AdminAnnouncementSchema{}, common.NewFDKError(err.Error())
        }
        return getAnnouncementByIdResponse, nil
        
    }
           
       
    
    
    
  

    
    // UpdateAnnouncement Update an announcement
     func (co *PlatformAppContent)  UpdateAnnouncement(AnnouncementID string, body  AdminAnnouncementSchema) (CreateAnnouncementSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            updateAnnouncementResponse CreateAnnouncementSchema
	    )

        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        

         

        
        
        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return CreateAnnouncementSchema{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return CreateAnnouncementSchema{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "put",
            fmt.Sprintf("/service/platform/content/v1.0/company/%s/application/%s/announcements/%s",co.CompanyID, co.ApplicationID, AnnouncementID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return CreateAnnouncementSchema{}, err
	    }
        
        err = json.Unmarshal(response, &updateAnnouncementResponse)
        if err != nil {
            return CreateAnnouncementSchema{}, common.NewFDKError(err.Error())
        }
        return updateAnnouncementResponse, nil
        
    }
           
       
    
    
    
  

    
    // UpdateAnnouncementSchedule Update the schedule and the publish status of an announcement
     func (co *PlatformAppContent)  UpdateAnnouncementSchedule(AnnouncementID string, body  ScheduleSchema) (CreateAnnouncementSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            updateAnnouncementScheduleResponse CreateAnnouncementSchema
	    )

        
            
        
            
        
            
        
            
        
            
        

         

        
        
        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return CreateAnnouncementSchema{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return CreateAnnouncementSchema{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "patch",
            fmt.Sprintf("/service/platform/content/v1.0/company/%s/application/%s/announcements/%s",co.CompanyID, co.ApplicationID, AnnouncementID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return CreateAnnouncementSchema{}, err
	    }
        
        err = json.Unmarshal(response, &updateAnnouncementScheduleResponse)
        if err != nil {
            return CreateAnnouncementSchema{}, common.NewFDKError(err.Error())
        }
        return updateAnnouncementScheduleResponse, nil
        
    }
           
       
    
    
    
  

    
    // DeleteAnnouncement Delete announcement by id
     func (co *PlatformAppContent)  DeleteAnnouncement(AnnouncementID string) (CreateAnnouncementSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            deleteAnnouncementResponse CreateAnnouncementSchema
	    )

        

         

        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "delete",
            fmt.Sprintf("/service/platform/content/v1.0/company/%s/application/%s/announcements/%s",co.CompanyID, co.ApplicationID, AnnouncementID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return CreateAnnouncementSchema{}, err
	    }
        
        err = json.Unmarshal(response, &deleteAnnouncementResponse)
        if err != nil {
            return CreateAnnouncementSchema{}, common.NewFDKError(err.Error())
        }
        return deleteAnnouncementResponse, nil
        
    }
           
       
    
    
    
  

    
    // CreateBlog Create a blog
     func (co *PlatformAppContent)  CreateBlog(body  BlogRequest) (BlogSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            createBlogResponse BlogSchema
	    )

        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        

         

        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return BlogSchema{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return BlogSchema{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "post",
            fmt.Sprintf("/service/platform/content/v1.0/company/%s/application/%s/blogs/",co.CompanyID, co.ApplicationID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return BlogSchema{}, err
	    }
        
        err = json.Unmarshal(response, &createBlogResponse)
        if err != nil {
            return BlogSchema{}, common.NewFDKError(err.Error())
        }
        return createBlogResponse, nil
        
    }
           
       
    
    
    
  

    
    //PlatformAppGetBlogsXQuery holds query params
    type PlatformAppGetBlogsXQuery struct { 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"`  
    }
    
    // GetBlogs Get blogs
     func (co *PlatformAppContent)  GetBlogs(xQuery PlatformAppGetBlogsXQuery) (BlogGetResponse, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getBlogsResponse BlogGetResponse
	    )

        

         
            
                
            
                
            
        

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            fmt.Sprintf("/service/platform/content/v1.0/company/%s/application/%s/blogs/",co.CompanyID, co.ApplicationID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return BlogGetResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getBlogsResponse)
        if err != nil {
            return BlogGetResponse{}, common.NewFDKError(err.Error())
        }
        return getBlogsResponse, nil
        
    }
           
            
            
            
            
            
            
            
             
            
            
             
            
            
             
            
            
             
            
            
            
            
            
            
            
            
            
            // GetBlogsPaginator Get blogs  
            func (co *PlatformAppContent)  GetBlogsPaginator( xQuery PlatformAppGetBlogsXQuery ) *common.Paginator {
                paginator := common.NewPaginator("number")
                
                
                 xQuery.PageNo  = paginator.PageNo
                 
                 
                 
                 
                 
                 
                 
                 
                paginator.Next = func() (interface{}, error) {
                    response, err := co.GetBlogs(xQuery)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, int(response.Page.Current+1), response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }
        
       
    
    
    
  

    
    // UpdateBlog Update a blog
     func (co *PlatformAppContent)  UpdateBlog(ID string, body  BlogRequest) (BlogSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            updateBlogResponse BlogSchema
	    )

        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        

         

        
        
        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return BlogSchema{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return BlogSchema{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "put",
            fmt.Sprintf("/service/platform/content/v1.0/company/%s/application/%s/blogs/%s",co.CompanyID, co.ApplicationID, ID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return BlogSchema{}, err
	    }
        
        err = json.Unmarshal(response, &updateBlogResponse)
        if err != nil {
            return BlogSchema{}, common.NewFDKError(err.Error())
        }
        return updateBlogResponse, nil
        
    }
           
       
    
    
    
  

    
    // DeleteBlog Delete blogs
     func (co *PlatformAppContent)  DeleteBlog(ID string) (BlogSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            deleteBlogResponse BlogSchema
	    )

        

         

        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "delete",
            fmt.Sprintf("/service/platform/content/v1.0/company/%s/application/%s/blogs/%s",co.CompanyID, co.ApplicationID, ID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return BlogSchema{}, err
	    }
        
        err = json.Unmarshal(response, &deleteBlogResponse)
        if err != nil {
            return BlogSchema{}, common.NewFDKError(err.Error())
        }
        return deleteBlogResponse, nil
        
    }
           
       
    
    
    
  

    
    // GetComponentById Get components of a blog
     func (co *PlatformAppContent)  GetComponentById(Slug string) (BlogSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getComponentByIdResponse BlogSchema
	    )

        

         

        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            fmt.Sprintf("/service/platform/content/v1.0/company/%s/application/%s/blogs/%s",co.CompanyID, co.ApplicationID, Slug),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return BlogSchema{}, err
	    }
        
        err = json.Unmarshal(response, &getComponentByIdResponse)
        if err != nil {
            return BlogSchema{}, common.NewFDKError(err.Error())
        }
        return getComponentByIdResponse, nil
        
    }
           
       
    
    
    
  

    
    // GetFaqCategories Get a list of FAQ categories
     func (co *PlatformAppContent)  GetFaqCategories() (GetFaqCategoriesSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getFaqCategoriesResponse GetFaqCategoriesSchema
	    )

        

         

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            fmt.Sprintf("/service/platform/content/v1.0/company/%s/application/%s/faq/categories",co.CompanyID, co.ApplicationID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return GetFaqCategoriesSchema{}, err
	    }
        
        err = json.Unmarshal(response, &getFaqCategoriesResponse)
        if err != nil {
            return GetFaqCategoriesSchema{}, common.NewFDKError(err.Error())
        }
        return getFaqCategoriesResponse, nil
        
    }
           
       
    
    
    
  

    
    // GetFaqCategoryBySlugOrId Get an FAQ category by slug or id
     func (co *PlatformAppContent)  GetFaqCategoryBySlugOrId(IDOrSlug string) (GetFaqCategoryBySlugSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getFaqCategoryBySlugOrIdResponse GetFaqCategoryBySlugSchema
	    )

        

         

        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            fmt.Sprintf("/service/platform/content/v1.0/company/%s/application/%s/faq/category/%s",co.CompanyID, co.ApplicationID, IDOrSlug),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return GetFaqCategoryBySlugSchema{}, err
	    }
        
        err = json.Unmarshal(response, &getFaqCategoryBySlugOrIdResponse)
        if err != nil {
            return GetFaqCategoryBySlugSchema{}, common.NewFDKError(err.Error())
        }
        return getFaqCategoryBySlugOrIdResponse, nil
        
    }
           
       
    
    
    
  

    
    // CreateFaqCategory Create an FAQ category
     func (co *PlatformAppContent)  CreateFaqCategory(body  CreateFaqCategoryRequestSchema) (CreateFaqCategorySchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            createFaqCategoryResponse CreateFaqCategorySchema
	    )

        
            
        

         

        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return CreateFaqCategorySchema{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return CreateFaqCategorySchema{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "post",
            fmt.Sprintf("/service/platform/content/v1.0/company/%s/application/%s/faq/category",co.CompanyID, co.ApplicationID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return CreateFaqCategorySchema{}, err
	    }
        
        err = json.Unmarshal(response, &createFaqCategoryResponse)
        if err != nil {
            return CreateFaqCategorySchema{}, common.NewFDKError(err.Error())
        }
        return createFaqCategoryResponse, nil
        
    }
           
       
    
    
    
  

    
    // UpdateFaqCategory Update an FAQ category
     func (co *PlatformAppContent)  UpdateFaqCategory(ID string, body  UpdateFaqCategoryRequestSchema) (CreateFaqCategorySchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            updateFaqCategoryResponse CreateFaqCategorySchema
	    )

        
            
        

         

        
        
        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return CreateFaqCategorySchema{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return CreateFaqCategorySchema{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "put",
            fmt.Sprintf("/service/platform/content/v1.0/company/%s/application/%s/faq/category/%s",co.CompanyID, co.ApplicationID, ID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return CreateFaqCategorySchema{}, err
	    }
        
        err = json.Unmarshal(response, &updateFaqCategoryResponse)
        if err != nil {
            return CreateFaqCategorySchema{}, common.NewFDKError(err.Error())
        }
        return updateFaqCategoryResponse, nil
        
    }
           
       
    
    
    
  

    
    // DeleteFaqCategory Delete an FAQ category
     func (co *PlatformAppContent)  DeleteFaqCategory(ID string) (FaqSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            deleteFaqCategoryResponse FaqSchema
	    )

        

         

        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "delete",
            fmt.Sprintf("/service/platform/content/v1.0/company/%s/application/%s/faq/category/%s",co.CompanyID, co.ApplicationID, ID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return FaqSchema{}, err
	    }
        
        err = json.Unmarshal(response, &deleteFaqCategoryResponse)
        if err != nil {
            return FaqSchema{}, common.NewFDKError(err.Error())
        }
        return deleteFaqCategoryResponse, nil
        
    }
           
       
    
    
    
  

    
    // GetFaqsByCategoryIdOrSlug Get question and answers within an FAQ category
     func (co *PlatformAppContent)  GetFaqsByCategoryIdOrSlug(IDOrSlug string) (GetFaqSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getFaqsByCategoryIdOrSlugResponse GetFaqSchema
	    )

        

         

        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            fmt.Sprintf("/service/platform/content/v1.0/company/%s/application/%s/faq/category/%s/faqs",co.CompanyID, co.ApplicationID, IDOrSlug),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return GetFaqSchema{}, err
	    }
        
        err = json.Unmarshal(response, &getFaqsByCategoryIdOrSlugResponse)
        if err != nil {
            return GetFaqSchema{}, common.NewFDKError(err.Error())
        }
        return getFaqsByCategoryIdOrSlugResponse, nil
        
    }
           
       
    
    
    
  

    
    // AddFaq Create an FAQ
     func (co *PlatformAppContent)  AddFaq(CategoryID string, body  CreateFaqSchema) (CreateFaqResponseSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            addFaqResponse CreateFaqResponseSchema
	    )

        
            
        

         

        
        
        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return CreateFaqResponseSchema{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return CreateFaqResponseSchema{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "post",
            fmt.Sprintf("/service/platform/content/v1.0/company/%s/application/%s/faq/category/%s/faqs",co.CompanyID, co.ApplicationID, CategoryID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return CreateFaqResponseSchema{}, err
	    }
        
        err = json.Unmarshal(response, &addFaqResponse)
        if err != nil {
            return CreateFaqResponseSchema{}, common.NewFDKError(err.Error())
        }
        return addFaqResponse, nil
        
    }
           
       
    
    
    
  

    
    // UpdateFaq Update an FAQ
     func (co *PlatformAppContent)  UpdateFaq(CategoryID string, FaqID string, body  CreateFaqSchema) (CreateFaqResponseSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            updateFaqResponse CreateFaqResponseSchema
	    )

        
            
        

         

        
        
        
        
        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return CreateFaqResponseSchema{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return CreateFaqResponseSchema{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "put",
            fmt.Sprintf("/service/platform/content/v1.0/company/%s/application/%s/faq/category/%s/faq/%s",co.CompanyID, co.ApplicationID, CategoryID, FaqID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return CreateFaqResponseSchema{}, err
	    }
        
        err = json.Unmarshal(response, &updateFaqResponse)
        if err != nil {
            return CreateFaqResponseSchema{}, common.NewFDKError(err.Error())
        }
        return updateFaqResponse, nil
        
    }
           
       
    
    
    
  

    
    // DeleteFaq Delete an FAQ
     func (co *PlatformAppContent)  DeleteFaq(CategoryID string, FaqID string) (CreateFaqResponseSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            deleteFaqResponse CreateFaqResponseSchema
	    )

        

         

        
        
        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "delete",
            fmt.Sprintf("/service/platform/content/v1.0/company/%s/application/%s/faq/category/%s/faq/%s",co.CompanyID, co.ApplicationID, CategoryID, FaqID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return CreateFaqResponseSchema{}, err
	    }
        
        err = json.Unmarshal(response, &deleteFaqResponse)
        if err != nil {
            return CreateFaqResponseSchema{}, common.NewFDKError(err.Error())
        }
        return deleteFaqResponse, nil
        
    }
           
       
    
    
    
  

    
    // GetFaqByIdOrSlug Get an FAQ
     func (co *PlatformAppContent)  GetFaqByIdOrSlug(IDOrSlug string) (CreateFaqResponseSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getFaqByIdOrSlugResponse CreateFaqResponseSchema
	    )

        

         

        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            fmt.Sprintf("/service/platform/content/v1.0/company/%s/application/%s/faq/%s",co.CompanyID, co.ApplicationID, IDOrSlug),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return CreateFaqResponseSchema{}, err
	    }
        
        err = json.Unmarshal(response, &getFaqByIdOrSlugResponse)
        if err != nil {
            return CreateFaqResponseSchema{}, common.NewFDKError(err.Error())
        }
        return getFaqByIdOrSlugResponse, nil
        
    }
           
       
    
    
    
  

    
    //PlatformAppGetLandingPagesXQuery holds query params
    type PlatformAppGetLandingPagesXQuery struct { 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"`  
    }
    
    // GetLandingPages Get landing pages
     func (co *PlatformAppContent)  GetLandingPages(xQuery PlatformAppGetLandingPagesXQuery) (LandingPageGetResponse, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getLandingPagesResponse LandingPageGetResponse
	    )

        

         
            
                
            
                
            
        

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            fmt.Sprintf("/service/platform/content/v1.0/company/%s/application/%s/landing-page/",co.CompanyID, co.ApplicationID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return LandingPageGetResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getLandingPagesResponse)
        if err != nil {
            return LandingPageGetResponse{}, common.NewFDKError(err.Error())
        }
        return getLandingPagesResponse, nil
        
    }
           
            
            
            
            
            
            
            
             
            
            
             
            
            
             
            
            
             
            
            
            
            
            
            
            
            
            
            // GetLandingPagesPaginator Get landing pages  
            func (co *PlatformAppContent)  GetLandingPagesPaginator( xQuery PlatformAppGetLandingPagesXQuery ) *common.Paginator {
                paginator := common.NewPaginator("number")
                
                
                 xQuery.PageNo  = paginator.PageNo
                 
                 
                 
                 
                 
                 
                 
                 
                paginator.Next = func() (interface{}, error) {
                    response, err := co.GetLandingPages(xQuery)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, int(response.Page.Current+1), response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }
        
       
    
    
    
  

    
    // CreateLandingPage Create a landing page
     func (co *PlatformAppContent)  CreateLandingPage(body  LandingPageSchema) (LandingPageSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            createLandingPageResponse LandingPageSchema
	    )

        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        

         

        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return LandingPageSchema{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return LandingPageSchema{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "post",
            fmt.Sprintf("/service/platform/content/v1.0/company/%s/application/%s/landing-page/",co.CompanyID, co.ApplicationID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return LandingPageSchema{}, err
	    }
        
        err = json.Unmarshal(response, &createLandingPageResponse)
        if err != nil {
            return LandingPageSchema{}, common.NewFDKError(err.Error())
        }
        return createLandingPageResponse, nil
        
    }
           
       
    
    
    
  

    
    // UpdateLandingPage Update a landing page
     func (co *PlatformAppContent)  UpdateLandingPage(ID string, body  LandingPageSchema) (LandingPageSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            updateLandingPageResponse LandingPageSchema
	    )

        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        

         

        
        
        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return LandingPageSchema{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return LandingPageSchema{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "put",
            fmt.Sprintf("/service/platform/content/v1.0/company/%s/application/%s/landing-page/%s",co.CompanyID, co.ApplicationID, ID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return LandingPageSchema{}, err
	    }
        
        err = json.Unmarshal(response, &updateLandingPageResponse)
        if err != nil {
            return LandingPageSchema{}, common.NewFDKError(err.Error())
        }
        return updateLandingPageResponse, nil
        
    }
           
       
    
    
    
  

    
    // DeleteLandingPage Delete a landing page
     func (co *PlatformAppContent)  DeleteLandingPage(ID string) (LandingPageSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            deleteLandingPageResponse LandingPageSchema
	    )

        

         

        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "delete",
            fmt.Sprintf("/service/platform/content/v1.0/company/%s/application/%s/landing-page/%s",co.CompanyID, co.ApplicationID, ID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return LandingPageSchema{}, err
	    }
        
        err = json.Unmarshal(response, &deleteLandingPageResponse)
        if err != nil {
            return LandingPageSchema{}, common.NewFDKError(err.Error())
        }
        return deleteLandingPageResponse, nil
        
    }
           
       
    
    
    
  

    
    // GetLegalInformation Get legal information
     func (co *PlatformAppContent)  GetLegalInformation() (ApplicationLegal, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getLegalInformationResponse ApplicationLegal
	    )

        

         

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            fmt.Sprintf("/service/platform/content/v1.0/company/%s/application/%s/legal",co.CompanyID, co.ApplicationID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return ApplicationLegal{}, err
	    }
        
        err = json.Unmarshal(response, &getLegalInformationResponse)
        if err != nil {
            return ApplicationLegal{}, common.NewFDKError(err.Error())
        }
        return getLegalInformationResponse, nil
        
    }
           
       
    
    
    
  

    
    // UpdateLegalInformation Save legal information
     func (co *PlatformAppContent)  UpdateLegalInformation(body  ApplicationLegal) (ApplicationLegal, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            updateLegalInformationResponse ApplicationLegal
	    )

        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        

         

        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return ApplicationLegal{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return ApplicationLegal{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "post",
            fmt.Sprintf("/service/platform/content/v1.0/company/%s/application/%s/legal",co.CompanyID, co.ApplicationID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return ApplicationLegal{}, err
	    }
        
        err = json.Unmarshal(response, &updateLegalInformationResponse)
        if err != nil {
            return ApplicationLegal{}, common.NewFDKError(err.Error())
        }
        return updateLegalInformationResponse, nil
        
    }
           
       
    
    
    
  

    
    //PlatformAppGetNavigationsXQuery holds query params
    type PlatformAppGetNavigationsXQuery struct { 
        DevicePlatform string  `url:"device_platform,omitempty"` 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"`  
    }
    
    // GetNavigations Get navigations
     func (co *PlatformAppContent)  GetNavigations(xQuery PlatformAppGetNavigationsXQuery) (NavigationGetResponse, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getNavigationsResponse NavigationGetResponse
	    )

        

         
            
                
            
                
            
                
            
        

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            fmt.Sprintf("/service/platform/content/v1.0/company/%s/application/%s/navigations/",co.CompanyID, co.ApplicationID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return NavigationGetResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getNavigationsResponse)
        if err != nil {
            return NavigationGetResponse{}, common.NewFDKError(err.Error())
        }
        return getNavigationsResponse, nil
        
    }
           
            
            
            
            
            
            
            
             
            
            
             
            
            
             
            
            
             
            
            
             
            
            
            
            
            
            
            
            
            
            
            
            
            // GetNavigationsPaginator Get navigations  
            func (co *PlatformAppContent)  GetNavigationsPaginator( xQuery PlatformAppGetNavigationsXQuery ) *common.Paginator {
                paginator := common.NewPaginator("number")
                
                
                 
                 
                 
                 
                 xQuery.PageNo  = paginator.PageNo
                 
                 
                 
                 
                 
                 
                 
                 
                paginator.Next = func() (interface{}, error) {
                    response, err := co.GetNavigations(xQuery)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, int(response.Page.Current+1), response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }
        
       
    
    
    
  

    
    // CreateNavigation Create a navigation
     func (co *PlatformAppContent)  CreateNavigation(body  NavigationRequest) (NavigationSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            createNavigationResponse NavigationSchema
	    )

        
            
        
            
        
            
        
            
        
            
        

         

        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return NavigationSchema{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return NavigationSchema{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "post",
            fmt.Sprintf("/service/platform/content/v1.0/company/%s/application/%s/navigations/",co.CompanyID, co.ApplicationID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return NavigationSchema{}, err
	    }
        
        err = json.Unmarshal(response, &createNavigationResponse)
        if err != nil {
            return NavigationSchema{}, common.NewFDKError(err.Error())
        }
        return createNavigationResponse, nil
        
    }
           
       
    
    
    
  

    
    // GetDefaultNavigations Get default navigations
     func (co *PlatformAppContent)  GetDefaultNavigations() (DefaultNavigationResponse, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getDefaultNavigationsResponse DefaultNavigationResponse
	    )

        

         

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            fmt.Sprintf("/service/platform/content/v1.0/company/%s/application/%s/navigations/default",co.CompanyID, co.ApplicationID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return DefaultNavigationResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getDefaultNavigationsResponse)
        if err != nil {
            return DefaultNavigationResponse{}, common.NewFDKError(err.Error())
        }
        return getDefaultNavigationsResponse, nil
        
    }
           
       
    
    
    
  

    
    //PlatformAppGetNavigationBySlugXQuery holds query params
    type PlatformAppGetNavigationBySlugXQuery struct { 
        DevicePlatform string  `url:"device_platform,omitempty"`  
    }
    
    // GetNavigationBySlug Get a navigation by slug
     func (co *PlatformAppContent)  GetNavigationBySlug(Slug string, xQuery PlatformAppGetNavigationBySlugXQuery) (NavigationSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getNavigationBySlugResponse NavigationSchema
	    )

        

         
            
                
            
        

        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            fmt.Sprintf("/service/platform/content/v1.0/company/%s/application/%s/navigations/%s",co.CompanyID, co.ApplicationID, Slug),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return NavigationSchema{}, err
	    }
        
        err = json.Unmarshal(response, &getNavigationBySlugResponse)
        if err != nil {
            return NavigationSchema{}, common.NewFDKError(err.Error())
        }
        return getNavigationBySlugResponse, nil
        
    }
           
       
    
    
    
  

    
    // UpdateNavigation Update a navigation
     func (co *PlatformAppContent)  UpdateNavigation(ID string, body  NavigationRequest) (NavigationSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            updateNavigationResponse NavigationSchema
	    )

        
            
        
            
        
            
        
            
        
            
        

         

        
        
        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return NavigationSchema{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return NavigationSchema{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "put",
            fmt.Sprintf("/service/platform/content/v1.0/company/%s/application/%s/navigations/%s",co.CompanyID, co.ApplicationID, ID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return NavigationSchema{}, err
	    }
        
        err = json.Unmarshal(response, &updateNavigationResponse)
        if err != nil {
            return NavigationSchema{}, common.NewFDKError(err.Error())
        }
        return updateNavigationResponse, nil
        
    }
           
       
    
    
    
  

    
    // DeleteNavigation Delete a navigation
     func (co *PlatformAppContent)  DeleteNavigation(ID string) (NavigationSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            deleteNavigationResponse NavigationSchema
	    )

        

         

        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "delete",
            fmt.Sprintf("/service/platform/content/v1.0/company/%s/application/%s/navigations/%s",co.CompanyID, co.ApplicationID, ID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return NavigationSchema{}, err
	    }
        
        err = json.Unmarshal(response, &deleteNavigationResponse)
        if err != nil {
            return NavigationSchema{}, common.NewFDKError(err.Error())
        }
        return deleteNavigationResponse, nil
        
    }
           
       
    
    
    
  

    
    // GetPageMeta Get page meta
     func (co *PlatformAppContent)  GetPageMeta() (PageMetaSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getPageMetaResponse PageMetaSchema
	    )

        

         

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            fmt.Sprintf("/service/platform/content/v1.0/company/%s/application/%s/pages/meta",co.CompanyID, co.ApplicationID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return PageMetaSchema{}, err
	    }
        
        err = json.Unmarshal(response, &getPageMetaResponse)
        if err != nil {
            return PageMetaSchema{}, common.NewFDKError(err.Error())
        }
        return getPageMetaResponse, nil
        
    }
           
       
    
    
    
  

    
    // GetPageSpec Get page spec
     func (co *PlatformAppContent)  GetPageSpec() (PageSpec, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getPageSpecResponse PageSpec
	    )

        

         

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            fmt.Sprintf("/service/platform/content/v1.0/company/%s/application/%s/pages/spec",co.CompanyID, co.ApplicationID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return PageSpec{}, err
	    }
        
        err = json.Unmarshal(response, &getPageSpecResponse)
        if err != nil {
            return PageSpec{}, common.NewFDKError(err.Error())
        }
        return getPageSpecResponse, nil
        
    }
           
       
    
    
    
  

    
    // CreatePage Create a page
     func (co *PlatformAppContent)  CreatePage(body  PageRequest) (PageSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            createPageResponse PageSchema
	    )

        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        

         

        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return PageSchema{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return PageSchema{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "post",
            fmt.Sprintf("/service/platform/content/v1.0/company/%s/application/%s/pages/",co.CompanyID, co.ApplicationID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return PageSchema{}, err
	    }
        
        err = json.Unmarshal(response, &createPageResponse)
        if err != nil {
            return PageSchema{}, common.NewFDKError(err.Error())
        }
        return createPageResponse, nil
        
    }
           
       
    
    
    
  

    
    //PlatformAppGetPagesXQuery holds query params
    type PlatformAppGetPagesXQuery struct { 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"`  
    }
    
    // GetPages Get a list of pages
     func (co *PlatformAppContent)  GetPages(xQuery PlatformAppGetPagesXQuery) (PageGetResponse, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getPagesResponse PageGetResponse
	    )

        

         
            
                
            
                
            
        

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            fmt.Sprintf("/service/platform/content/v1.0/company/%s/application/%s/pages/",co.CompanyID, co.ApplicationID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return PageGetResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getPagesResponse)
        if err != nil {
            return PageGetResponse{}, common.NewFDKError(err.Error())
        }
        return getPagesResponse, nil
        
    }
           
            
            
            
            
            
            
            
             
            
            
             
            
            
             
            
            
             
            
            
            
            
            
            
            
            
            
            // GetPagesPaginator Get a list of pages  
            func (co *PlatformAppContent)  GetPagesPaginator( xQuery PlatformAppGetPagesXQuery ) *common.Paginator {
                paginator := common.NewPaginator("number")
                
                
                 xQuery.PageNo  = paginator.PageNo
                 
                 
                 
                 
                 
                 
                 
                 
                paginator.Next = func() (interface{}, error) {
                    response, err := co.GetPages(xQuery)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, int(response.Page.Current+1), response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }
        
       
    
    
    
  

    
    // CreatePagePreview Create a page preview
     func (co *PlatformAppContent)  CreatePagePreview(body  PageRequest) (PageSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            createPagePreviewResponse PageSchema
	    )

        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        

         

        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return PageSchema{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return PageSchema{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "post",
            fmt.Sprintf("/service/platform/content/v1.0/company/%s/application/%s/pages/preview/",co.CompanyID, co.ApplicationID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return PageSchema{}, err
	    }
        
        err = json.Unmarshal(response, &createPagePreviewResponse)
        if err != nil {
            return PageSchema{}, common.NewFDKError(err.Error())
        }
        return createPagePreviewResponse, nil
        
    }
           
       
    
    
    
  

    
    // UpdatePagePreview Change the publish status of a page
     func (co *PlatformAppContent)  UpdatePagePreview(Slug string, body  PagePublishRequest) (PageSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            updatePagePreviewResponse PageSchema
	    )

        
            
        

         

        
        
        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return PageSchema{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return PageSchema{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "put",
            fmt.Sprintf("/service/platform/content/v1.0/company/%s/application/%s/pages/publish/%s",co.CompanyID, co.ApplicationID, Slug),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return PageSchema{}, err
	    }
        
        err = json.Unmarshal(response, &updatePagePreviewResponse)
        if err != nil {
            return PageSchema{}, common.NewFDKError(err.Error())
        }
        return updatePagePreviewResponse, nil
        
    }
           
       
    
    
    
  

    
    // UpdatePage Update a page
     func (co *PlatformAppContent)  UpdatePage(ID string, body  PageSchema) (PageSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            updatePageResponse PageSchema
	    )

        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        

         

        
        
        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return PageSchema{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return PageSchema{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "put",
            fmt.Sprintf("/service/platform/content/v1.0/company/%s/application/%s/pages/%s",co.CompanyID, co.ApplicationID, ID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return PageSchema{}, err
	    }
        
        err = json.Unmarshal(response, &updatePageResponse)
        if err != nil {
            return PageSchema{}, common.NewFDKError(err.Error())
        }
        return updatePageResponse, nil
        
    }
           
       
    
    
    
  

    
    // DeletePage Delete a page
     func (co *PlatformAppContent)  DeletePage(ID string) (PageSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            deletePageResponse PageSchema
	    )

        

         

        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "delete",
            fmt.Sprintf("/service/platform/content/v1.0/company/%s/application/%s/pages/%s",co.CompanyID, co.ApplicationID, ID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return PageSchema{}, err
	    }
        
        err = json.Unmarshal(response, &deletePageResponse)
        if err != nil {
            return PageSchema{}, common.NewFDKError(err.Error())
        }
        return deletePageResponse, nil
        
    }
           
       
    
    
    
  

    
    // GetPageBySlug Get pages by component Id
     func (co *PlatformAppContent)  GetPageBySlug(Slug string) (PageSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getPageBySlugResponse PageSchema
	    )

        

         

        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            fmt.Sprintf("/service/platform/content/v1.0/company/%s/application/%s/pages/%s",co.CompanyID, co.ApplicationID, Slug),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return PageSchema{}, err
	    }
        
        err = json.Unmarshal(response, &getPageBySlugResponse)
        if err != nil {
            return PageSchema{}, common.NewFDKError(err.Error())
        }
        return getPageBySlugResponse, nil
        
    }
           
       
    
    
    
  

    
    // GetSEOConfiguration Get SEO configuration of an application
     func (co *PlatformAppContent)  GetSEOConfiguration() (SeoComponent, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getSEOConfigurationResponse SeoComponent
	    )

        

         

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            fmt.Sprintf("/service/platform/content/v1.0/company/%s/application/%s/seo",co.CompanyID, co.ApplicationID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return SeoComponent{}, err
	    }
        
        err = json.Unmarshal(response, &getSEOConfigurationResponse)
        if err != nil {
            return SeoComponent{}, common.NewFDKError(err.Error())
        }
        return getSEOConfigurationResponse, nil
        
    }
           
       
    
    
    
  

    
    // UpdateSEOConfiguration Update SEO of application
     func (co *PlatformAppContent)  UpdateSEOConfiguration(body  SeoComponent) (SeoSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            updateSEOConfigurationResponse SeoSchema
	    )

        
            
        

         

        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return SeoSchema{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return SeoSchema{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "post",
            fmt.Sprintf("/service/platform/content/v1.0/company/%s/application/%s/seo",co.CompanyID, co.ApplicationID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return SeoSchema{}, err
	    }
        
        err = json.Unmarshal(response, &updateSEOConfigurationResponse)
        if err != nil {
            return SeoSchema{}, common.NewFDKError(err.Error())
        }
        return updateSEOConfigurationResponse, nil
        
    }
           
       
    
    
    
  

    
    //PlatformAppGetSlideshowsXQuery holds query params
    type PlatformAppGetSlideshowsXQuery struct { 
        DevicePlatform string  `url:"device_platform,omitempty"` 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"`  
    }
    
    // GetSlideshows Get slideshows
     func (co *PlatformAppContent)  GetSlideshows(xQuery PlatformAppGetSlideshowsXQuery) (SlideshowGetResponse, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getSlideshowsResponse SlideshowGetResponse
	    )

        

         
            
                
            
                
            
                
            
        

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            fmt.Sprintf("/service/platform/content/v1.0/company/%s/application/%s/slideshows/",co.CompanyID, co.ApplicationID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return SlideshowGetResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getSlideshowsResponse)
        if err != nil {
            return SlideshowGetResponse{}, common.NewFDKError(err.Error())
        }
        return getSlideshowsResponse, nil
        
    }
           
            
            
            
            
            
            
            
             
            
            
             
            
            
             
            
            
             
            
            
             
            
            
            
            
            
            
            
            
            
            
            
            
            // GetSlideshowsPaginator Get slideshows  
            func (co *PlatformAppContent)  GetSlideshowsPaginator( xQuery PlatformAppGetSlideshowsXQuery ) *common.Paginator {
                paginator := common.NewPaginator("number")
                
                
                 
                 
                 
                 
                 xQuery.PageNo  = paginator.PageNo
                 
                 
                 
                 
                 
                 
                 
                 
                paginator.Next = func() (interface{}, error) {
                    response, err := co.GetSlideshows(xQuery)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, int(response.Page.Current+1), response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }
        
       
    
    
    
  

    
    // CreateSlideshow Create a slideshow
     func (co *PlatformAppContent)  CreateSlideshow(body  SlideshowRequest) (SlideshowSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            createSlideshowResponse SlideshowSchema
	    )

        
            
        
            
        
            
        
            
        
            
        

         

        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return SlideshowSchema{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return SlideshowSchema{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "post",
            fmt.Sprintf("/service/platform/content/v1.0/company/%s/application/%s/slideshows/",co.CompanyID, co.ApplicationID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return SlideshowSchema{}, err
	    }
        
        err = json.Unmarshal(response, &createSlideshowResponse)
        if err != nil {
            return SlideshowSchema{}, common.NewFDKError(err.Error())
        }
        return createSlideshowResponse, nil
        
    }
           
       
    
    
    
  

    
    //PlatformAppGetSlideshowBySlugXQuery holds query params
    type PlatformAppGetSlideshowBySlugXQuery struct { 
        DevicePlatform string  `url:"device_platform,omitempty"`  
    }
    
    // GetSlideshowBySlug Get slideshow by slug
     func (co *PlatformAppContent)  GetSlideshowBySlug(Slug string, xQuery PlatformAppGetSlideshowBySlugXQuery) (SlideshowSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getSlideshowBySlugResponse SlideshowSchema
	    )

        

         
            
                
            
        

        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            fmt.Sprintf("/service/platform/content/v1.0/company/%s/application/%s/slideshows/%s",co.CompanyID, co.ApplicationID, Slug),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return SlideshowSchema{}, err
	    }
        
        err = json.Unmarshal(response, &getSlideshowBySlugResponse)
        if err != nil {
            return SlideshowSchema{}, common.NewFDKError(err.Error())
        }
        return getSlideshowBySlugResponse, nil
        
    }
           
       
    
    
    
  

    
    // UpdateSlideshow Update a slideshow
     func (co *PlatformAppContent)  UpdateSlideshow(ID string, body  SlideshowRequest) (SlideshowSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            updateSlideshowResponse SlideshowSchema
	    )

        
            
        
            
        
            
        
            
        
            
        

         

        
        
        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return SlideshowSchema{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return SlideshowSchema{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "put",
            fmt.Sprintf("/service/platform/content/v1.0/company/%s/application/%s/slideshows/%s",co.CompanyID, co.ApplicationID, ID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return SlideshowSchema{}, err
	    }
        
        err = json.Unmarshal(response, &updateSlideshowResponse)
        if err != nil {
            return SlideshowSchema{}, common.NewFDKError(err.Error())
        }
        return updateSlideshowResponse, nil
        
    }
           
       
    
    
    
  

    
    // DeleteSlideshow Delete a slideshow
     func (co *PlatformAppContent)  DeleteSlideshow(ID string) (SlideshowSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            deleteSlideshowResponse SlideshowSchema
	    )

        

         

        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "delete",
            fmt.Sprintf("/service/platform/content/v1.0/company/%s/application/%s/slideshows/%s",co.CompanyID, co.ApplicationID, ID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return SlideshowSchema{}, err
	    }
        
        err = json.Unmarshal(response, &deleteSlideshowResponse)
        if err != nil {
            return SlideshowSchema{}, common.NewFDKError(err.Error())
        }
        return deleteSlideshowResponse, nil
        
    }
           
       
    
    
    
  

    
    // GetSupportInformation Get support information
     func (co *PlatformAppContent)  GetSupportInformation() (Support, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getSupportInformationResponse Support
	    )

        

         

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            fmt.Sprintf("/service/platform/content/v1.0/company/%s/application/%s/support",co.CompanyID, co.ApplicationID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return Support{}, err
	    }
        
        err = json.Unmarshal(response, &getSupportInformationResponse)
        if err != nil {
            return Support{}, common.NewFDKError(err.Error())
        }
        return getSupportInformationResponse, nil
        
    }
           
       
    
    
    
  

    
    // UpdateSupportInformation Update the support data of an application
     func (co *PlatformAppContent)  UpdateSupportInformation(body  Support) (Support, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            updateSupportInformationResponse Support
	    )

        
            
        
            
        
            
        
            
        
            
        
            
        
            
        

         

        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return Support{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return Support{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "post",
            fmt.Sprintf("/service/platform/content/v1.0/company/%s/application/%s/support",co.CompanyID, co.ApplicationID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return Support{}, err
	    }
        
        err = json.Unmarshal(response, &updateSupportInformationResponse)
        if err != nil {
            return Support{}, common.NewFDKError(err.Error())
        }
        return updateSupportInformationResponse, nil
        
    }
           
       
    
    
    
  

    
    // UpdateInjectableTag Update a tag
     func (co *PlatformAppContent)  UpdateInjectableTag(body  CreateTagRequestSchema) (TagsSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            updateInjectableTagResponse TagsSchema
	    )

        
            
        

         

        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return TagsSchema{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return TagsSchema{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "put",
            fmt.Sprintf("/service/platform/content/v1.0/company/%s/application/%s/tags",co.CompanyID, co.ApplicationID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return TagsSchema{}, err
	    }
        
        err = json.Unmarshal(response, &updateInjectableTagResponse)
        if err != nil {
            return TagsSchema{}, common.NewFDKError(err.Error())
        }
        return updateInjectableTagResponse, nil
        
    }
           
       
    
    
    
  

    
    // DeleteAllInjectableTags Delete tags in application
     func (co *PlatformAppContent)  DeleteAllInjectableTags() (TagsSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            deleteAllInjectableTagsResponse TagsSchema
	    )

        

         

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "delete",
            fmt.Sprintf("/service/platform/content/v1.0/company/%s/application/%s/tags",co.CompanyID, co.ApplicationID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return TagsSchema{}, err
	    }
        
        err = json.Unmarshal(response, &deleteAllInjectableTagsResponse)
        if err != nil {
            return TagsSchema{}, common.NewFDKError(err.Error())
        }
        return deleteAllInjectableTagsResponse, nil
        
    }
           
       
    
    
    
  

    
    // GetInjectableTags Get all the tags in an application
     func (co *PlatformAppContent)  GetInjectableTags() (TagsSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getInjectableTagsResponse TagsSchema
	    )

        

         

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            fmt.Sprintf("/service/platform/content/v1.0/company/%s/application/%s/tags",co.CompanyID, co.ApplicationID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return TagsSchema{}, err
	    }
        
        err = json.Unmarshal(response, &getInjectableTagsResponse)
        if err != nil {
            return TagsSchema{}, common.NewFDKError(err.Error())
        }
        return getInjectableTagsResponse, nil
        
    }
           
       
    
    
    
  

    
    // AddInjectableTag Add a tag
     func (co *PlatformAppContent)  AddInjectableTag(body  CreateTagRequestSchema) (TagsSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            addInjectableTagResponse TagsSchema
	    )

        
            
        

         

        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return TagsSchema{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return TagsSchema{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "put",
            fmt.Sprintf("/service/platform/content/v1.0/company/%s/application/%s/tags/add",co.CompanyID, co.ApplicationID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return TagsSchema{}, err
	    }
        
        err = json.Unmarshal(response, &addInjectableTagResponse)
        if err != nil {
            return TagsSchema{}, common.NewFDKError(err.Error())
        }
        return addInjectableTagResponse, nil
        
    }
           
       
    
    
    
  

    
    // RemoveInjectableTag Remove a tag
     func (co *PlatformAppContent)  RemoveInjectableTag(body  RemoveHandpickedSchema) (TagsSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            removeInjectableTagResponse TagsSchema
	    )

        
            
        

         

        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return TagsSchema{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return TagsSchema{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "put",
            fmt.Sprintf("/service/platform/content/v1.0/company/%s/application/%s/tags/remove/handpicked",co.CompanyID, co.ApplicationID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return TagsSchema{}, err
	    }
        
        err = json.Unmarshal(response, &removeInjectableTagResponse)
        if err != nil {
            return TagsSchema{}, common.NewFDKError(err.Error())
        }
        return removeInjectableTagResponse, nil
        
    }
           
       
    
    
    
  

    
    // EditInjectableTag Edit a tag by id
     func (co *PlatformAppContent)  EditInjectableTag(TagID string, body  UpdateHandpickedSchema) (TagsSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            editInjectableTagResponse TagsSchema
	    )

        
            
        

         

        
        
        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return TagsSchema{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return TagsSchema{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "put",
            fmt.Sprintf("/service/platform/content/v1.0/company/%s/application/%s/tags/edit/handpicked/%s",co.CompanyID, co.ApplicationID, TagID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return TagsSchema{}, err
	    }
        
        err = json.Unmarshal(response, &editInjectableTagResponse)
        if err != nil {
            return TagsSchema{}, common.NewFDKError(err.Error())
        }
        return editInjectableTagResponse, nil
        
    }
           
       
    

 
	 
   // PlatformAppAssignment holds PlatformAppAssignment object properties
    type PlatformAppAssignment struct {
        config *PlatformConfig
        CompanyID string
        ApplicationID string
    }
    // NewPlatformAppAssignment returns new PlatformAppAssignment instance
    func NewPlatformAppAssignment(config *PlatformConfig, appID string) *PlatformAppAssignment {
        return &PlatformAppAssignment{config, config.CompanyID, appID}
    }
    
    
    
  

    
    // CreatePickupLocation 
     func (as *PlatformAppAssignment)  CreatePickupLocation(body  PickupPointSchema) (Success, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            createPickupLocationResponse Success
	    )

        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        

         

        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return Success{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return Success{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            as.config,
            "post",
            fmt.Sprintf("/service/platform/assignment/v1.0/company/%s/application/%s/pickup/locations",as.CompanyID, as.ApplicationID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return Success{}, err
	    }
        
        err = json.Unmarshal(response, &createPickupLocationResponse)
        if err != nil {
            return Success{}, common.NewFDKError(err.Error())
        }
        return createPickupLocationResponse, nil
        
    }
           
       
    
    
    
  

    
    //PlatformAppGetPickupLocationXQuery holds query params
    type PlatformAppGetPickupLocationXQuery struct { 
        Q string  `url:"q,omitempty"` 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"`  
    }
    
    // GetPickupLocation 
     func (as *PlatformAppAssignment)  GetPickupLocation(xQuery PlatformAppGetPickupLocationXQuery) (PickupPointResponse, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getPickupLocationResponse PickupPointResponse
	    )

        

         
            
                
            
                
            
                
            
        

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            as.config,
            "get",
            fmt.Sprintf("/service/platform/assignment/v1.0/company/%s/application/%s/pickup/locations",as.CompanyID, as.ApplicationID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return PickupPointResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getPickupLocationResponse)
        if err != nil {
            return PickupPointResponse{}, common.NewFDKError(err.Error())
        }
        return getPickupLocationResponse, nil
        
    }
           
            
            
            
            
            
            
            
             
            
            
             
            
            
             
            
            
             
            
            
             
            
            
            
            
            
            
            
            
            
            
            
            
            // GetPickupLocationPaginator   
            func (as *PlatformAppAssignment)  GetPickupLocationPaginator( xQuery PlatformAppGetPickupLocationXQuery ) *common.Paginator {
                paginator := common.NewPaginator("number")
                
                
                 
                 
                 
                 
                 xQuery.PageNo  = paginator.PageNo
                 
                 
                 
                 
                 
                 
                 
                 
                paginator.Next = func() (interface{}, error) {
                    response, err := as.GetPickupLocation(xQuery)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, int(response.Page.Current+1), response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }
        
       
    
    
    
  

    
    // UpdatePickupLocation 
     func (as *PlatformAppAssignment)  UpdatePickupLocation(ID float64, body  PickupPointSchema) (Success, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            updatePickupLocationResponse Success
	    )

        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        

         

        
        
        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return Success{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return Success{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            as.config,
            "put",
            fmt.Sprintf("/service/platform/assignment/v1.0/company/%s/application/%s/pickup/locations/undefined",ID, as.CompanyID, as.ApplicationID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return Success{}, err
	    }
        
        err = json.Unmarshal(response, &updatePickupLocationResponse)
        if err != nil {
            return Success{}, common.NewFDKError(err.Error())
        }
        return updatePickupLocationResponse, nil
        
    }
           
       
    
    
    
  

    
    // GetPickupLocationById 
     func (as *PlatformAppAssignment)  GetPickupLocationById(ID float64, ) (PickupPointSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getPickupLocationByIdResponse PickupPointSchema
	    )

        

         

        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            as.config,
            "get",
            fmt.Sprintf("/service/platform/assignment/v1.0/company/%s/application/%s/pickup/locations/undefined",ID, as.CompanyID, as.ApplicationID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return PickupPointSchema{}, err
	    }
        
        err = json.Unmarshal(response, &getPickupLocationByIdResponse)
        if err != nil {
            return PickupPointSchema{}, common.NewFDKError(err.Error())
        }
        return getPickupLocationByIdResponse, nil
        
    }
           
       
    
    
    
  

    
    // CreatePickupConfiguration 
     func (as *PlatformAppAssignment)  CreatePickupConfiguration(body  PickupConfiguration) (Success, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            createPickupConfigurationResponse Success
	    )

        
            
        
            
        
            
        

         

        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return Success{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return Success{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            as.config,
            "post",
            fmt.Sprintf("/service/platform/assignment/v1.0/company/%s/application/%s/pickup/configurations",as.CompanyID, as.ApplicationID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return Success{}, err
	    }
        
        err = json.Unmarshal(response, &createPickupConfigurationResponse)
        if err != nil {
            return Success{}, common.NewFDKError(err.Error())
        }
        return createPickupConfigurationResponse, nil
        
    }
           
       
    
    
    
  

    
    // GetPickupConfiguration 
     func (as *PlatformAppAssignment)  GetPickupConfiguration() (PickupResponse, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getPickupConfigurationResponse PickupResponse
	    )

        

         

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            as.config,
            "get",
            fmt.Sprintf("/service/platform/assignment/v1.0/company/%s/application/%s/pickup/configurations",as.CompanyID, as.ApplicationID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return PickupResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getPickupConfigurationResponse)
        if err != nil {
            return PickupResponse{}, common.NewFDKError(err.Error())
        }
        return getPickupConfigurationResponse, nil
        
    }
           
       
    
    
    
  

    
    // GetAllocationConfiguration 
     func (as *PlatformAppAssignment)  GetAllocationConfiguration() (ShippingResponse, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getAllocationConfigurationResponse ShippingResponse
	    )

        

         

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            as.config,
            "get",
            fmt.Sprintf("/service/platform/assignment/v1.0/company/%s/application/%s/allocate/configurations",as.CompanyID, as.ApplicationID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return ShippingResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getAllocationConfigurationResponse)
        if err != nil {
            return ShippingResponse{}, common.NewFDKError(err.Error())
        }
        return getAllocationConfigurationResponse, nil
        
    }
           
       
    
    
    
  

    
    // CreateAllocationConfiguration 
     func (as *PlatformAppAssignment)  CreateAllocationConfiguration(body  ShippingSchema) (Success, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            createAllocationConfigurationResponse Success
	    )

        
            
        
            
        
            
        

         

        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return Success{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return Success{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            as.config,
            "post",
            fmt.Sprintf("/service/platform/assignment/v1.0/company/%s/application/%s/allocate/configurations",as.CompanyID, as.ApplicationID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return Success{}, err
	    }
        
        err = json.Unmarshal(response, &createAllocationConfigurationResponse)
        if err != nil {
            return Success{}, common.NewFDKError(err.Error())
        }
        return createAllocationConfigurationResponse, nil
        
    }
           
       
    
    
    
  

    
    // UpdateAllocationConfiguration 
     func (as *PlatformAppAssignment)  UpdateAllocationConfiguration(body  ShippingSchema) (Success, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            updateAllocationConfigurationResponse Success
	    )

        
            
        
            
        
            
        

         

        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return Success{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return Success{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            as.config,
            "put",
            fmt.Sprintf("/service/platform/assignment/v1.0/company/%s/application/%s/allocate/configurations",as.CompanyID, as.ApplicationID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return Success{}, err
	    }
        
        err = json.Unmarshal(response, &updateAllocationConfigurationResponse)
        if err != nil {
            return Success{}, common.NewFDKError(err.Error())
        }
        return updateAllocationConfigurationResponse, nil
        
    }
           
       
    
    
    
  

    
    //PlatformAppGetAllocationLocationsXQuery holds query params
    type PlatformAppGetAllocationLocationsXQuery struct { 
        Q string  `url:"q,omitempty"` 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"` 
        Type string  `url:"type,omitempty"`  
    }
    
    // GetAllocationLocations 
     func (as *PlatformAppAssignment)  GetAllocationLocations(xQuery PlatformAppGetAllocationLocationsXQuery) (StoreListResponse, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getAllocationLocationsResponse StoreListResponse
	    )

        

         
            
                
            
                
            
                
            
                
            
        

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            as.config,
            "get",
            fmt.Sprintf("/service/platform/assignment/v1.0/company/%s/application/%s/allocate/locations/",as.CompanyID, as.ApplicationID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return StoreListResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getAllocationLocationsResponse)
        if err != nil {
            return StoreListResponse{}, common.NewFDKError(err.Error())
        }
        return getAllocationLocationsResponse, nil
        
    }
           
            
            
            
            
            
            
            
             
            
            
             
            
            
             
            
            
             
            
            
             
            
            
             
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            // GetAllocationLocationsPaginator   
            func (as *PlatformAppAssignment)  GetAllocationLocationsPaginator( xQuery PlatformAppGetAllocationLocationsXQuery ) *common.Paginator {
                paginator := common.NewPaginator("number")
                
                
                 
                 
                 
                 
                 xQuery.PageNo  = paginator.PageNo
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                paginator.Next = func() (interface{}, error) {
                    response, err := as.GetAllocationLocations(xQuery)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, int(response.Page.Current+1), response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }
        
       
    
    
    
  

    
    // GetAllocationLocationById 
     func (as *PlatformAppAssignment)  GetAllocationLocationById(ID float64, ) (StoreResponse, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getAllocationLocationByIdResponse StoreResponse
	    )

        

         

        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            as.config,
            "get",
            fmt.Sprintf("/service/platform/assignment/v1.0/company/%s/application/%s/allocate/locations/undefined",ID, as.CompanyID, as.ApplicationID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return StoreResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getAllocationLocationByIdResponse)
        if err != nil {
            return StoreResponse{}, common.NewFDKError(err.Error())
        }
        return getAllocationLocationByIdResponse, nil
        
    }
           
       
    
    
    
  

    
    // UpdateAllocationLocation 
     func (as *PlatformAppAssignment)  UpdateAllocationLocation(ID float64, body  StoreSchema) (Success, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            updateAllocationLocationResponse Success
	    )

        
            
        
            
        

         

        
        
        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return Success{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return Success{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            as.config,
            "put",
            fmt.Sprintf("/service/platform/assignment/v1.0/company/%s/application/%s/allocate/locations/undefined",ID, as.CompanyID, as.ApplicationID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return Success{}, err
	    }
        
        err = json.Unmarshal(response, &updateAllocationLocationResponse)
        if err != nil {
            return Success{}, common.NewFDKError(err.Error())
        }
        return updateAllocationLocationResponse, nil
        
    }
           
       
    
    
    
  

    
    // CreateAllocationLocation 
     func (as *PlatformAppAssignment)  CreateAllocationLocation(body  StoreSchema) (Success, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            createAllocationLocationResponse Success
	    )

        
            
        
            
        

         

        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return Success{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return Success{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            as.config,
            "post",
            fmt.Sprintf("/service/platform/assignment/v1.0/company/%s/application/%s/allocate/locations",as.CompanyID, as.ApplicationID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return Success{}, err
	    }
        
        err = json.Unmarshal(response, &createAllocationLocationResponse)
        if err != nil {
            return Success{}, common.NewFDKError(err.Error())
        }
        return createAllocationLocationResponse, nil
        
    }
           
       
    
    
    
  

    
    //PlatformAppGetDestinationZonesXQuery holds query params
    type PlatformAppGetDestinationZonesXQuery struct { 
        Q string  `url:"q,omitempty"` 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"`  
    }
    
    // GetDestinationZones 
     func (as *PlatformAppAssignment)  GetDestinationZones(xQuery PlatformAppGetDestinationZonesXQuery) (ZoneListResponse, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getDestinationZonesResponse ZoneListResponse
	    )

        

         
            
                
            
                
            
                
            
        

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            as.config,
            "get",
            fmt.Sprintf("/service/platform/assignment/v1.0/company/%s/application/%s/destination-zones",as.CompanyID, as.ApplicationID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return ZoneListResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getDestinationZonesResponse)
        if err != nil {
            return ZoneListResponse{}, common.NewFDKError(err.Error())
        }
        return getDestinationZonesResponse, nil
        
    }
           
            
            
            
            
            
            
            
             
            
            
             
            
            
             
            
            
             
            
            
             
            
            
            
            
            
            
            
            
            
            
            
            
            // GetDestinationZonesPaginator   
            func (as *PlatformAppAssignment)  GetDestinationZonesPaginator( xQuery PlatformAppGetDestinationZonesXQuery ) *common.Paginator {
                paginator := common.NewPaginator("number")
                
                
                 
                 
                 
                 
                 xQuery.PageNo  = paginator.PageNo
                 
                 
                 
                 
                 
                 
                 
                 
                paginator.Next = func() (interface{}, error) {
                    response, err := as.GetDestinationZones(xQuery)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, int(response.Page.Current+1), response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }
        
       
    
    
    
  

    
    // PostDestinationZone 
     func (as *PlatformAppAssignment)  PostDestinationZone(body  ZoneSchema) (Success, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            postDestinationZoneResponse Success
	    )

        
            
        
            
        

         

        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return Success{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return Success{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            as.config,
            "post",
            fmt.Sprintf("/service/platform/assignment/v1.0/company/%s/application/%s/destination-zones",as.CompanyID, as.ApplicationID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return Success{}, err
	    }
        
        err = json.Unmarshal(response, &postDestinationZoneResponse)
        if err != nil {
            return Success{}, common.NewFDKError(err.Error())
        }
        return postDestinationZoneResponse, nil
        
    }
           
       
    
    
    
  

    
    // GetDestinationZoneById 
     func (as *PlatformAppAssignment)  GetDestinationZoneById(ID float64, ) (ZoneSchema, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getDestinationZoneByIdResponse ZoneSchema
	    )

        

         

        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            as.config,
            "get",
            fmt.Sprintf("/service/platform/assignment/v1.0/company/%s/application/%s/destination-zones/undefined",ID, as.CompanyID, as.ApplicationID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return ZoneSchema{}, err
	    }
        
        err = json.Unmarshal(response, &getDestinationZoneByIdResponse)
        if err != nil {
            return ZoneSchema{}, common.NewFDKError(err.Error())
        }
        return getDestinationZoneByIdResponse, nil
        
    }
           
       
    
    
    
  

    
    // UpdateDestinationZone 
     func (as *PlatformAppAssignment)  UpdateDestinationZone(ID float64, body  ZoneSchema) (Success, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            updateDestinationZoneResponse Success
	    )

        
            
        
            
        

         

        
        
        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return Success{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return Success{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            as.config,
            "put",
            fmt.Sprintf("/service/platform/assignment/v1.0/company/%s/application/%s/destination-zones/undefined",ID, as.CompanyID, as.ApplicationID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return Success{}, err
	    }
        
        err = json.Unmarshal(response, &updateDestinationZoneResponse)
        if err != nil {
            return Success{}, common.NewFDKError(err.Error())
        }
        return updateDestinationZoneResponse, nil
        
    }
           
       
    

 
	 
   // PlatformAppCommunication holds PlatformAppCommunication object properties
    type PlatformAppCommunication struct {
        config *PlatformConfig
        CompanyID string
        ApplicationID string
    }
    // NewPlatformAppCommunication returns new PlatformAppCommunication instance
    func NewPlatformAppCommunication(config *PlatformConfig, appID string) *PlatformAppCommunication {
        return &PlatformAppCommunication{config, config.CompanyID, appID}
    }
    
    
    
  

    
    //PlatformAppGetCampaignsXQuery holds query params
    type PlatformAppGetCampaignsXQuery struct { 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"` 
        Sort map[string]interface{}  `url:"sort,omitempty"`  
    }
    
    // GetCampaigns Get campaigns
     func (co *PlatformAppCommunication)  GetCampaigns(xQuery PlatformAppGetCampaignsXQuery) (Campaigns, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getCampaignsResponse Campaigns
	    )

        

         
            
                
            
                
            
                
            
        

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            fmt.Sprintf("/service/platform/communication/v1.0/company/%s/application/%s/campaigns/campaigns",co.CompanyID, co.ApplicationID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return Campaigns{}, err
	    }
        
        err = json.Unmarshal(response, &getCampaignsResponse)
        if err != nil {
            return Campaigns{}, common.NewFDKError(err.Error())
        }
        return getCampaignsResponse, nil
        
    }
           
            
            
            
            
            
            
            
             
            
            
             
            
            
             
            
            
             
            
            
             
            
            
            
            
            
            
            
            
            
            
            
            
            // GetCampaignsPaginator Get campaigns  
            func (co *PlatformAppCommunication)  GetCampaignsPaginator( xQuery PlatformAppGetCampaignsXQuery ) *common.Paginator {
                paginator := common.NewPaginator("number")
                
                
                 xQuery.PageNo  = paginator.PageNo
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                paginator.Next = func() (interface{}, error) {
                    response, err := co.GetCampaigns(xQuery)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, int(response.Page.Current+1), response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }
        
       
    
    
    
  

    
    // CreateCampaign Create campaign
     func (co *PlatformAppCommunication)  CreateCampaign(body  CampaignReq) (Campaign, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            createCampaignResponse Campaign
	    )

        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        

         

        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return Campaign{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return Campaign{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "post",
            fmt.Sprintf("/service/platform/communication/v1.0/company/%s/application/%s/campaigns/campaigns",co.CompanyID, co.ApplicationID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return Campaign{}, err
	    }
        
        err = json.Unmarshal(response, &createCampaignResponse)
        if err != nil {
            return Campaign{}, common.NewFDKError(err.Error())
        }
        return createCampaignResponse, nil
        
    }
           
       
    
    
    
  

    
    // GetCampaignById Get campaign by id
     func (co *PlatformAppCommunication)  GetCampaignById(ID string) (Campaign, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getCampaignByIdResponse Campaign
	    )

        

         

        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            fmt.Sprintf("/service/platform/communication/v1.0/company/%s/application/%s/campaigns/campaigns/%s",co.CompanyID, co.ApplicationID, ID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return Campaign{}, err
	    }
        
        err = json.Unmarshal(response, &getCampaignByIdResponse)
        if err != nil {
            return Campaign{}, common.NewFDKError(err.Error())
        }
        return getCampaignByIdResponse, nil
        
    }
           
       
    
    
    
  

    
    // UpdateCampaignById Update campaign by id
     func (co *PlatformAppCommunication)  UpdateCampaignById(ID string, body  CampaignReq) (Campaign, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            updateCampaignByIdResponse Campaign
	    )

        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        

         

        
        
        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return Campaign{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return Campaign{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "put",
            fmt.Sprintf("/service/platform/communication/v1.0/company/%s/application/%s/campaigns/campaigns/%s",co.CompanyID, co.ApplicationID, ID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return Campaign{}, err
	    }
        
        err = json.Unmarshal(response, &updateCampaignByIdResponse)
        if err != nil {
            return Campaign{}, common.NewFDKError(err.Error())
        }
        return updateCampaignByIdResponse, nil
        
    }
           
       
    
    
    
  

    
    // GetStatsOfCampaignById Get stats of campaign by id
     func (co *PlatformAppCommunication)  GetStatsOfCampaignById(ID string) (GetStats, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getStatsOfCampaignByIdResponse GetStats
	    )

        

         

        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            fmt.Sprintf("/service/platform/communication/v1.0/company/%s/application/%s/campaigns/get-stats/%s",co.CompanyID, co.ApplicationID, ID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return GetStats{}, err
	    }
        
        err = json.Unmarshal(response, &getStatsOfCampaignByIdResponse)
        if err != nil {
            return GetStats{}, common.NewFDKError(err.Error())
        }
        return getStatsOfCampaignByIdResponse, nil
        
    }
           
       
    
    
    
  

    
    //PlatformAppGetAudiencesXQuery holds query params
    type PlatformAppGetAudiencesXQuery struct { 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"` 
        Sort map[string]interface{}  `url:"sort,omitempty"`  
    }
    
    // GetAudiences Get audiences
     func (co *PlatformAppCommunication)  GetAudiences(xQuery PlatformAppGetAudiencesXQuery) (Audiences, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getAudiencesResponse Audiences
	    )

        

         
            
                
            
                
            
                
            
        

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            fmt.Sprintf("/service/platform/communication/v1.0/company/%s/application/%s/sources/datasources",co.CompanyID, co.ApplicationID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return Audiences{}, err
	    }
        
        err = json.Unmarshal(response, &getAudiencesResponse)
        if err != nil {
            return Audiences{}, common.NewFDKError(err.Error())
        }
        return getAudiencesResponse, nil
        
    }
           
            
            
            
            
            
            
            
             
            
            
             
            
            
             
            
            
             
            
            
             
            
            
            
            
            
            
            
            
            
            
            
            
            // GetAudiencesPaginator Get audiences  
            func (co *PlatformAppCommunication)  GetAudiencesPaginator( xQuery PlatformAppGetAudiencesXQuery ) *common.Paginator {
                paginator := common.NewPaginator("number")
                
                
                 xQuery.PageNo  = paginator.PageNo
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                paginator.Next = func() (interface{}, error) {
                    response, err := co.GetAudiences(xQuery)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, int(response.Page.Current+1), response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }
        
       
    
    
    
  

    
    // CreateAudience Create audience
     func (co *PlatformAppCommunication)  CreateAudience(body  AudienceReq) (Audience, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            createAudienceResponse Audience
	    )

        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        

         

        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return Audience{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return Audience{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "post",
            fmt.Sprintf("/service/platform/communication/v1.0/company/%s/application/%s/sources/datasources",co.CompanyID, co.ApplicationID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return Audience{}, err
	    }
        
        err = json.Unmarshal(response, &createAudienceResponse)
        if err != nil {
            return Audience{}, common.NewFDKError(err.Error())
        }
        return createAudienceResponse, nil
        
    }
           
       
    
    
    
  

    
    // GetBigqueryHeaders Get bigquery headers
     func (co *PlatformAppCommunication)  GetBigqueryHeaders(body  BigqueryHeadersReq) (BigqueryHeadersRes, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getBigqueryHeadersResponse BigqueryHeadersRes
	    )

        
            
        
            
        

         

        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return BigqueryHeadersRes{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return BigqueryHeadersRes{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "post",
            fmt.Sprintf("/service/platform/communication/v1.0/company/%s/application/%s/sources/bigquery-headers",co.CompanyID, co.ApplicationID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return BigqueryHeadersRes{}, err
	    }
        
        err = json.Unmarshal(response, &getBigqueryHeadersResponse)
        if err != nil {
            return BigqueryHeadersRes{}, common.NewFDKError(err.Error())
        }
        return getBigqueryHeadersResponse, nil
        
    }
           
       
    
    
    
  

    
    // GetAudienceById Get audience by id
     func (co *PlatformAppCommunication)  GetAudienceById(ID string) (Audience, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getAudienceByIdResponse Audience
	    )

        

         

        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            fmt.Sprintf("/service/platform/communication/v1.0/company/%s/application/%s/sources/datasources/%s",co.CompanyID, co.ApplicationID, ID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return Audience{}, err
	    }
        
        err = json.Unmarshal(response, &getAudienceByIdResponse)
        if err != nil {
            return Audience{}, common.NewFDKError(err.Error())
        }
        return getAudienceByIdResponse, nil
        
    }
           
       
    
    
    
  

    
    // UpdateAudienceById Update audience by id
     func (co *PlatformAppCommunication)  UpdateAudienceById(ID string, body  AudienceReq) (Audience, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            updateAudienceByIdResponse Audience
	    )

        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        

         

        
        
        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return Audience{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return Audience{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "put",
            fmt.Sprintf("/service/platform/communication/v1.0/company/%s/application/%s/sources/datasources/%s",co.CompanyID, co.ApplicationID, ID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return Audience{}, err
	    }
        
        err = json.Unmarshal(response, &updateAudienceByIdResponse)
        if err != nil {
            return Audience{}, common.NewFDKError(err.Error())
        }
        return updateAudienceByIdResponse, nil
        
    }
           
       
    
    
    
  

    
    // GetNSampleRecordsFromCsv Get n sample records from csv
     func (co *PlatformAppCommunication)  GetNSampleRecordsFromCsv(body  GetNRecordsCsvReq) (GetNRecordsCsvRes, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getNSampleRecordsFromCsvResponse GetNRecordsCsvRes
	    )

        
            
        
            
        
            
        

         

        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return GetNRecordsCsvRes{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return GetNRecordsCsvRes{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "post",
            fmt.Sprintf("/service/platform/communication/v1.0/company/%s/application/%s/sources/get-n-records",co.CompanyID, co.ApplicationID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return GetNRecordsCsvRes{}, err
	    }
        
        err = json.Unmarshal(response, &getNSampleRecordsFromCsvResponse)
        if err != nil {
            return GetNRecordsCsvRes{}, common.NewFDKError(err.Error())
        }
        return getNSampleRecordsFromCsvResponse, nil
        
    }
           
       
    
    
    
  

    
    //PlatformAppGetEmailProvidersXQuery holds query params
    type PlatformAppGetEmailProvidersXQuery struct { 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"` 
        Sort map[string]interface{}  `url:"sort,omitempty"`  
    }
    
    // GetEmailProviders Get email providers
     func (co *PlatformAppCommunication)  GetEmailProviders(xQuery PlatformAppGetEmailProvidersXQuery) (EmailProviders, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getEmailProvidersResponse EmailProviders
	    )

        

         
            
                
            
                
            
                
            
        

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            fmt.Sprintf("/service/platform/communication/v1.0/company/%s/application/%s/email/providers",co.CompanyID, co.ApplicationID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return EmailProviders{}, err
	    }
        
        err = json.Unmarshal(response, &getEmailProvidersResponse)
        if err != nil {
            return EmailProviders{}, common.NewFDKError(err.Error())
        }
        return getEmailProvidersResponse, nil
        
    }
           
            
            
            
            
            
            
            
             
            
            
             
            
            
             
            
            
             
            
            
             
            
            
            
            
            
            
            
            
            
            
            
            
            // GetEmailProvidersPaginator Get email providers  
            func (co *PlatformAppCommunication)  GetEmailProvidersPaginator( xQuery PlatformAppGetEmailProvidersXQuery ) *common.Paginator {
                paginator := common.NewPaginator("number")
                
                
                 xQuery.PageNo  = paginator.PageNo
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                paginator.Next = func() (interface{}, error) {
                    response, err := co.GetEmailProviders(xQuery)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, int(response.Page.Current+1), response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }
        
       
    
    
    
  

    
    // CreateEmailProvider Create email provider
     func (co *PlatformAppCommunication)  CreateEmailProvider(body  EmailProviderReq) (EmailProvider, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            createEmailProviderResponse EmailProvider
	    )

        
            
        
            
        
            
        
            
        
            
        
            
        

         

        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return EmailProvider{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return EmailProvider{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "post",
            fmt.Sprintf("/service/platform/communication/v1.0/company/%s/application/%s/email/providers",co.CompanyID, co.ApplicationID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return EmailProvider{}, err
	    }
        
        err = json.Unmarshal(response, &createEmailProviderResponse)
        if err != nil {
            return EmailProvider{}, common.NewFDKError(err.Error())
        }
        return createEmailProviderResponse, nil
        
    }
           
       
    
    
    
  

    
    // GetEmailProviderById Get email provider by id
     func (co *PlatformAppCommunication)  GetEmailProviderById(ID string) (EmailProvider, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getEmailProviderByIdResponse EmailProvider
	    )

        

         

        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            fmt.Sprintf("/service/platform/communication/v1.0/company/%s/application/%s/email/providers/%s",co.CompanyID, co.ApplicationID, ID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return EmailProvider{}, err
	    }
        
        err = json.Unmarshal(response, &getEmailProviderByIdResponse)
        if err != nil {
            return EmailProvider{}, common.NewFDKError(err.Error())
        }
        return getEmailProviderByIdResponse, nil
        
    }
           
       
    
    
    
  

    
    // UpdateEmailProviderById Update email provider by id
     func (co *PlatformAppCommunication)  UpdateEmailProviderById(ID string, body  EmailProviderReq) (EmailProvider, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            updateEmailProviderByIdResponse EmailProvider
	    )

        
            
        
            
        
            
        
            
        
            
        
            
        

         

        
        
        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return EmailProvider{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return EmailProvider{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "put",
            fmt.Sprintf("/service/platform/communication/v1.0/company/%s/application/%s/email/providers/%s",co.CompanyID, co.ApplicationID, ID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return EmailProvider{}, err
	    }
        
        err = json.Unmarshal(response, &updateEmailProviderByIdResponse)
        if err != nil {
            return EmailProvider{}, common.NewFDKError(err.Error())
        }
        return updateEmailProviderByIdResponse, nil
        
    }
           
       
    
    
    
  

    
    //PlatformAppGetEmailTemplatesXQuery holds query params
    type PlatformAppGetEmailTemplatesXQuery struct { 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"` 
        Sort map[string]interface{}  `url:"sort,omitempty"`  
    }
    
    // GetEmailTemplates Get email templates
     func (co *PlatformAppCommunication)  GetEmailTemplates(xQuery PlatformAppGetEmailTemplatesXQuery) (EmailTemplates, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getEmailTemplatesResponse EmailTemplates
	    )

        

         
            
                
            
                
            
                
            
        

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            fmt.Sprintf("/service/platform/communication/v1.0/company/%s/application/%s/email/templates",co.CompanyID, co.ApplicationID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return EmailTemplates{}, err
	    }
        
        err = json.Unmarshal(response, &getEmailTemplatesResponse)
        if err != nil {
            return EmailTemplates{}, common.NewFDKError(err.Error())
        }
        return getEmailTemplatesResponse, nil
        
    }
           
            
            
            
            
            
            
            
             
            
            
             
            
            
             
            
            
             
            
            
             
            
            
            
            
            
            
            
            
            
            
            
            
            // GetEmailTemplatesPaginator Get email templates  
            func (co *PlatformAppCommunication)  GetEmailTemplatesPaginator( xQuery PlatformAppGetEmailTemplatesXQuery ) *common.Paginator {
                paginator := common.NewPaginator("number")
                
                
                 xQuery.PageNo  = paginator.PageNo
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                paginator.Next = func() (interface{}, error) {
                    response, err := co.GetEmailTemplates(xQuery)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, int(response.Page.Current+1), response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }
        
       
    
    
    
  

    
    // CreateEmailTemplate Create email template
     func (co *PlatformAppCommunication)  CreateEmailTemplate(body  EmailTemplateReq) (EmailTemplateRes, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            createEmailTemplateResponse EmailTemplateRes
	    )

        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        

         

        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return EmailTemplateRes{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return EmailTemplateRes{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "post",
            fmt.Sprintf("/service/platform/communication/v1.0/company/%s/application/%s/email/templates",co.CompanyID, co.ApplicationID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return EmailTemplateRes{}, err
	    }
        
        err = json.Unmarshal(response, &createEmailTemplateResponse)
        if err != nil {
            return EmailTemplateRes{}, common.NewFDKError(err.Error())
        }
        return createEmailTemplateResponse, nil
        
    }
           
       
    
    
    
  

    
    //PlatformAppGetSystemEmailTemplatesXQuery holds query params
    type PlatformAppGetSystemEmailTemplatesXQuery struct { 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"` 
        Sort map[string]interface{}  `url:"sort,omitempty"`  
    }
    
    // GetSystemEmailTemplates Get system email templates
     func (co *PlatformAppCommunication)  GetSystemEmailTemplates(xQuery PlatformAppGetSystemEmailTemplatesXQuery) (SystemEmailTemplates, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getSystemEmailTemplatesResponse SystemEmailTemplates
	    )

        

         
            
                
            
                
            
                
            
        

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            fmt.Sprintf("/service/platform/communication/v1.0/company/%s/application/%s/email/system-templates",co.CompanyID, co.ApplicationID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return SystemEmailTemplates{}, err
	    }
        
        err = json.Unmarshal(response, &getSystemEmailTemplatesResponse)
        if err != nil {
            return SystemEmailTemplates{}, common.NewFDKError(err.Error())
        }
        return getSystemEmailTemplatesResponse, nil
        
    }
           
            
            
            
            
            
            
            
             
            
            
             
            
            
             
            
            
             
            
            
             
            
            
            
            
            
            
            
            
            
            
            
            
            // GetSystemEmailTemplatesPaginator Get system email templates  
            func (co *PlatformAppCommunication)  GetSystemEmailTemplatesPaginator( xQuery PlatformAppGetSystemEmailTemplatesXQuery ) *common.Paginator {
                paginator := common.NewPaginator("number")
                
                
                 xQuery.PageNo  = paginator.PageNo
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                paginator.Next = func() (interface{}, error) {
                    response, err := co.GetSystemEmailTemplates(xQuery)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, int(response.Page.Current+1), response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }
        
       
    
    
    
  

    
    // GetEmailTemplateById Get email template by id
     func (co *PlatformAppCommunication)  GetEmailTemplateById(ID string) (EmailTemplate, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getEmailTemplateByIdResponse EmailTemplate
	    )

        

         

        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            fmt.Sprintf("/service/platform/communication/v1.0/company/%s/application/%s/email/templates/%s",co.CompanyID, co.ApplicationID, ID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return EmailTemplate{}, err
	    }
        
        err = json.Unmarshal(response, &getEmailTemplateByIdResponse)
        if err != nil {
            return EmailTemplate{}, common.NewFDKError(err.Error())
        }
        return getEmailTemplateByIdResponse, nil
        
    }
           
       
    
    
    
  

    
    // UpdateEmailTemplateById Update email template by id
     func (co *PlatformAppCommunication)  UpdateEmailTemplateById(ID string, body  EmailTemplateReq) (EmailTemplateRes, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            updateEmailTemplateByIdResponse EmailTemplateRes
	    )

        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        

         

        
        
        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return EmailTemplateRes{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return EmailTemplateRes{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "put",
            fmt.Sprintf("/service/platform/communication/v1.0/company/%s/application/%s/email/templates/%s",co.CompanyID, co.ApplicationID, ID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return EmailTemplateRes{}, err
	    }
        
        err = json.Unmarshal(response, &updateEmailTemplateByIdResponse)
        if err != nil {
            return EmailTemplateRes{}, common.NewFDKError(err.Error())
        }
        return updateEmailTemplateByIdResponse, nil
        
    }
           
       
    
    
    
  

    
    // DeleteEmailTemplateById Delete email template by id
     func (co *PlatformAppCommunication)  DeleteEmailTemplateById(ID string) (EmailTemplateDeleteSuccessRes, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            deleteEmailTemplateByIdResponse EmailTemplateDeleteSuccessRes
	    )

        

         

        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "delete",
            fmt.Sprintf("/service/platform/communication/v1.0/company/%s/application/%s/email/templates/%s",co.CompanyID, co.ApplicationID, ID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return EmailTemplateDeleteSuccessRes{}, err
	    }
        
        err = json.Unmarshal(response, &deleteEmailTemplateByIdResponse)
        if err != nil {
            return EmailTemplateDeleteSuccessRes{}, common.NewFDKError(err.Error())
        }
        return deleteEmailTemplateByIdResponse, nil
        
    }
           
       
    
    
    
  

    
    //PlatformAppGetEventSubscriptionsXQuery holds query params
    type PlatformAppGetEventSubscriptionsXQuery struct { 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"` 
        Populate string  `url:"populate,omitempty"`  
    }
    
    // GetEventSubscriptions Get event subscriptions
     func (co *PlatformAppCommunication)  GetEventSubscriptions(xQuery PlatformAppGetEventSubscriptionsXQuery) (EventSubscriptions, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getEventSubscriptionsResponse EventSubscriptions
	    )

        

         
            
                
            
                
            
                
            
        

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            fmt.Sprintf("/service/platform/communication/v1.0/company/%s/application/%s/event/event-subscriptions",co.CompanyID, co.ApplicationID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return EventSubscriptions{}, err
	    }
        
        err = json.Unmarshal(response, &getEventSubscriptionsResponse)
        if err != nil {
            return EventSubscriptions{}, common.NewFDKError(err.Error())
        }
        return getEventSubscriptionsResponse, nil
        
    }
           
            
            
            
            
            
            
            
             
            
            
             
            
            
             
            
            
             
            
            
             
            
            
            
            
            
            
            
            
            
            
            
            
            // GetEventSubscriptionsPaginator Get event subscriptions  
            func (co *PlatformAppCommunication)  GetEventSubscriptionsPaginator( xQuery PlatformAppGetEventSubscriptionsXQuery ) *common.Paginator {
                paginator := common.NewPaginator("number")
                
                
                 xQuery.PageNo  = paginator.PageNo
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                paginator.Next = func() (interface{}, error) {
                    response, err := co.GetEventSubscriptions(xQuery)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, int(response.Page.Current+1), response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }
        
       
    
    
    
  

    
    //PlatformAppGetJobsXQuery holds query params
    type PlatformAppGetJobsXQuery struct { 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"` 
        Sort map[string]interface{}  `url:"sort,omitempty"`  
    }
    
    // GetJobs Get jobs
     func (co *PlatformAppCommunication)  GetJobs(xQuery PlatformAppGetJobsXQuery) (Jobs, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getJobsResponse Jobs
	    )

        

         
            
                
            
                
            
                
            
        

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            fmt.Sprintf("/service/platform/communication/v1.0/company/%s/application/%s/jobs/jobs",co.CompanyID, co.ApplicationID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return Jobs{}, err
	    }
        
        err = json.Unmarshal(response, &getJobsResponse)
        if err != nil {
            return Jobs{}, common.NewFDKError(err.Error())
        }
        return getJobsResponse, nil
        
    }
           
            
            
            
            
            
            
            
             
            
            
             
            
            
             
            
            
             
            
            
             
            
            
            
            
            
            
            
            
            
            
            
            
            // GetJobsPaginator Get jobs  
            func (co *PlatformAppCommunication)  GetJobsPaginator( xQuery PlatformAppGetJobsXQuery ) *common.Paginator {
                paginator := common.NewPaginator("number")
                
                
                 xQuery.PageNo  = paginator.PageNo
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                paginator.Next = func() (interface{}, error) {
                    response, err := co.GetJobs(xQuery)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, int(response.Page.Current+1), response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }
        
       
    
    
    
  

    
    // TriggerCampaignJob Trigger campaign job
     func (co *PlatformAppCommunication)  TriggerCampaignJob(body  TriggerJobRequest) (TriggerJobResponse, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            triggerCampaignJobResponse TriggerJobResponse
	    )

        
            
        

         

        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return TriggerJobResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return TriggerJobResponse{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "post",
            fmt.Sprintf("/service/platform/communication/v1.0/company/%s/application/%s/jobs/trigger-job",co.CompanyID, co.ApplicationID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return TriggerJobResponse{}, err
	    }
        
        err = json.Unmarshal(response, &triggerCampaignJobResponse)
        if err != nil {
            return TriggerJobResponse{}, common.NewFDKError(err.Error())
        }
        return triggerCampaignJobResponse, nil
        
    }
           
       
    
    
    
  

    
    //PlatformAppGetJobLogsXQuery holds query params
    type PlatformAppGetJobLogsXQuery struct { 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"` 
        Sort map[string]interface{}  `url:"sort,omitempty"`  
    }
    
    // GetJobLogs Get job logs
     func (co *PlatformAppCommunication)  GetJobLogs(xQuery PlatformAppGetJobLogsXQuery) (JobLogs, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getJobLogsResponse JobLogs
	    )

        

         
            
                
            
                
            
                
            
        

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            fmt.Sprintf("/service/platform/communication/v1.0/company/%s/application/%s/jobs/logs",co.CompanyID, co.ApplicationID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return JobLogs{}, err
	    }
        
        err = json.Unmarshal(response, &getJobLogsResponse)
        if err != nil {
            return JobLogs{}, common.NewFDKError(err.Error())
        }
        return getJobLogsResponse, nil
        
    }
           
            
            
            
            
            
            
            
             
            
            
             
            
            
             
            
            
             
            
            
             
            
            
            
            
            
            
            
            
            
            
            
            
            // GetJobLogsPaginator Get job logs  
            func (co *PlatformAppCommunication)  GetJobLogsPaginator( xQuery PlatformAppGetJobLogsXQuery ) *common.Paginator {
                paginator := common.NewPaginator("number")
                
                
                 xQuery.PageNo  = paginator.PageNo
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                paginator.Next = func() (interface{}, error) {
                    response, err := co.GetJobLogs(xQuery)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, int(response.Page.Current+1), response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }
        
       
    
    
    
  

    
    //PlatformAppGetCommunicationLogsXQuery holds query params
    type PlatformAppGetCommunicationLogsXQuery struct { 
        PageID string  `url:"page_id,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"` 
        Sort map[string]interface{}  `url:"sort,omitempty"` 
        Query map[string]interface{}  `url:"query,omitempty"`  
    }
    
    // GetCommunicationLogs Get communication logs
     func (co *PlatformAppCommunication)  GetCommunicationLogs(xQuery PlatformAppGetCommunicationLogsXQuery) (Logs, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getCommunicationLogsResponse Logs
	    )

        

         
            
                
            
                
            
                
            
                
            
        

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            fmt.Sprintf("/service/platform/communication/v1.0/company/%s/application/%s/log",co.CompanyID, co.ApplicationID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return Logs{}, err
	    }
        
        err = json.Unmarshal(response, &getCommunicationLogsResponse)
        if err != nil {
            return Logs{}, common.NewFDKError(err.Error())
        }
        return getCommunicationLogsResponse, nil
        
    }
           
            
            
            
            
            
            
            
             
            
            
             
            
            
             
            
            
             
            
            
             
            
            
             
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            // GetCommunicationLogsPaginator Get communication logs  
            func (co *PlatformAppCommunication)  GetCommunicationLogsPaginator( xQuery PlatformAppGetCommunicationLogsXQuery ) *common.Paginator {
                paginator := common.NewPaginator("cursor")
                
                
                 
                 xQuery.PageID = paginator.NextID
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                paginator.Next = func() (interface{}, error) {
                    response, err := co.GetCommunicationLogs(xQuery)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, int(response.Page.Current+1), response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }
        
       
    
    
    
    
    
  

    
    //PlatformAppGetSmsProvidersXQuery holds query params
    type PlatformAppGetSmsProvidersXQuery struct { 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"` 
        Sort map[string]interface{}  `url:"sort,omitempty"`  
    }
    
    // GetSmsProviders Get sms providers
     func (co *PlatformAppCommunication)  GetSmsProviders(xQuery PlatformAppGetSmsProvidersXQuery) (SmsProviders, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getSmsProvidersResponse SmsProviders
	    )

        

         
            
                
            
                
            
                
            
        

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            fmt.Sprintf("/service/platform/communication/v1.0/company/%s/application/%s/sms/providers",co.CompanyID, co.ApplicationID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return SmsProviders{}, err
	    }
        
        err = json.Unmarshal(response, &getSmsProvidersResponse)
        if err != nil {
            return SmsProviders{}, common.NewFDKError(err.Error())
        }
        return getSmsProvidersResponse, nil
        
    }
           
            
            
            
            
            
            
            
             
            
            
             
            
            
             
            
            
             
            
            
             
            
            
            
            
            
            
            
            
            
            
            
            
            // GetSmsProvidersPaginator Get sms providers  
            func (co *PlatformAppCommunication)  GetSmsProvidersPaginator( xQuery PlatformAppGetSmsProvidersXQuery ) *common.Paginator {
                paginator := common.NewPaginator("number")
                
                
                 xQuery.PageNo  = paginator.PageNo
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                paginator.Next = func() (interface{}, error) {
                    response, err := co.GetSmsProviders(xQuery)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, int(response.Page.Current+1), response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }
        
       
    
    
    
  

    
    // CreateSmsProvider Create sms provider
     func (co *PlatformAppCommunication)  CreateSmsProvider(body  SmsProviderReq) (SmsProvider, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            createSmsProviderResponse SmsProvider
	    )

        
            
        
            
        
            
        
            
        
            
        
            
        
            
        

         

        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return SmsProvider{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return SmsProvider{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "post",
            fmt.Sprintf("/service/platform/communication/v1.0/company/%s/application/%s/sms/providers",co.CompanyID, co.ApplicationID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return SmsProvider{}, err
	    }
        
        err = json.Unmarshal(response, &createSmsProviderResponse)
        if err != nil {
            return SmsProvider{}, common.NewFDKError(err.Error())
        }
        return createSmsProviderResponse, nil
        
    }
           
       
    
    
    
  

    
    // GetSmsProviderById Get sms provider by id
     func (co *PlatformAppCommunication)  GetSmsProviderById(ID string) (SmsProvider, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getSmsProviderByIdResponse SmsProvider
	    )

        

         

        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            fmt.Sprintf("/service/platform/communication/v1.0/company/%s/application/%s/sms/providers/%s",co.CompanyID, co.ApplicationID, ID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return SmsProvider{}, err
	    }
        
        err = json.Unmarshal(response, &getSmsProviderByIdResponse)
        if err != nil {
            return SmsProvider{}, common.NewFDKError(err.Error())
        }
        return getSmsProviderByIdResponse, nil
        
    }
           
       
    
    
    
  

    
    // UpdateSmsProviderById Update sms provider by id
     func (co *PlatformAppCommunication)  UpdateSmsProviderById(ID string, body  SmsProviderReq) (SmsProvider, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            updateSmsProviderByIdResponse SmsProvider
	    )

        
            
        
            
        
            
        
            
        
            
        
            
        
            
        

         

        
        
        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return SmsProvider{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return SmsProvider{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "put",
            fmt.Sprintf("/service/platform/communication/v1.0/company/%s/application/%s/sms/providers/%s",co.CompanyID, co.ApplicationID, ID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return SmsProvider{}, err
	    }
        
        err = json.Unmarshal(response, &updateSmsProviderByIdResponse)
        if err != nil {
            return SmsProvider{}, common.NewFDKError(err.Error())
        }
        return updateSmsProviderByIdResponse, nil
        
    }
           
       
    
    
    
  

    
    //PlatformAppGetSmsTemplatesXQuery holds query params
    type PlatformAppGetSmsTemplatesXQuery struct { 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"` 
        Sort map[string]interface{}  `url:"sort,omitempty"`  
    }
    
    // GetSmsTemplates Get sms templates
     func (co *PlatformAppCommunication)  GetSmsTemplates(xQuery PlatformAppGetSmsTemplatesXQuery) (SmsTemplates, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getSmsTemplatesResponse SmsTemplates
	    )

        

         
            
                
            
                
            
                
            
        

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            fmt.Sprintf("/service/platform/communication/v1.0/company/%s/application/%s/sms/templates",co.CompanyID, co.ApplicationID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return SmsTemplates{}, err
	    }
        
        err = json.Unmarshal(response, &getSmsTemplatesResponse)
        if err != nil {
            return SmsTemplates{}, common.NewFDKError(err.Error())
        }
        return getSmsTemplatesResponse, nil
        
    }
           
            
            
            
            
            
            
            
             
            
            
             
            
            
             
            
            
             
            
            
             
            
            
            
            
            
            
            
            
            
            
            
            
            // GetSmsTemplatesPaginator Get sms templates  
            func (co *PlatformAppCommunication)  GetSmsTemplatesPaginator( xQuery PlatformAppGetSmsTemplatesXQuery ) *common.Paginator {
                paginator := common.NewPaginator("number")
                
                
                 xQuery.PageNo  = paginator.PageNo
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                paginator.Next = func() (interface{}, error) {
                    response, err := co.GetSmsTemplates(xQuery)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, int(response.Page.Current+1), response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }
        
       
    
    
    
  

    
    // CreateSmsTemplate Create sms template
     func (co *PlatformAppCommunication)  CreateSmsTemplate(body  SmsTemplateReq) (SmsTemplateRes, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            createSmsTemplateResponse SmsTemplateRes
	    )

        
            
        
            
        
            
        
            
        
            
        
            
        

         

        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return SmsTemplateRes{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return SmsTemplateRes{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "post",
            fmt.Sprintf("/service/platform/communication/v1.0/company/%s/application/%s/sms/templates",co.CompanyID, co.ApplicationID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return SmsTemplateRes{}, err
	    }
        
        err = json.Unmarshal(response, &createSmsTemplateResponse)
        if err != nil {
            return SmsTemplateRes{}, common.NewFDKError(err.Error())
        }
        return createSmsTemplateResponse, nil
        
    }
           
       
    
    
    
  

    
    // GetSmsTemplateById Get sms template by id
     func (co *PlatformAppCommunication)  GetSmsTemplateById(ID string) (SmsTemplate, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getSmsTemplateByIdResponse SmsTemplate
	    )

        

         

        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            fmt.Sprintf("/service/platform/communication/v1.0/company/%s/application/%s/sms/templates/%s",co.CompanyID, co.ApplicationID, ID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return SmsTemplate{}, err
	    }
        
        err = json.Unmarshal(response, &getSmsTemplateByIdResponse)
        if err != nil {
            return SmsTemplate{}, common.NewFDKError(err.Error())
        }
        return getSmsTemplateByIdResponse, nil
        
    }
           
       
    
    
    
  

    
    // UpdateSmsTemplateById Update sms template by id
     func (co *PlatformAppCommunication)  UpdateSmsTemplateById(ID string, body  SmsTemplateReq) (SmsTemplateRes, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            updateSmsTemplateByIdResponse SmsTemplateRes
	    )

        
            
        
            
        
            
        
            
        
            
        
            
        

         

        
        
        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return SmsTemplateRes{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return SmsTemplateRes{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "put",
            fmt.Sprintf("/service/platform/communication/v1.0/company/%s/application/%s/sms/templates/%s",co.CompanyID, co.ApplicationID, ID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return SmsTemplateRes{}, err
	    }
        
        err = json.Unmarshal(response, &updateSmsTemplateByIdResponse)
        if err != nil {
            return SmsTemplateRes{}, common.NewFDKError(err.Error())
        }
        return updateSmsTemplateByIdResponse, nil
        
    }
           
       
    
    
    
  

    
    // DeleteSmsTemplateById Delete sms template by id
     func (co *PlatformAppCommunication)  DeleteSmsTemplateById(ID string) (SmsTemplateDeleteSuccessRes, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            deleteSmsTemplateByIdResponse SmsTemplateDeleteSuccessRes
	    )

        

         

        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "delete",
            fmt.Sprintf("/service/platform/communication/v1.0/company/%s/application/%s/sms/templates/%s",co.CompanyID, co.ApplicationID, ID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return SmsTemplateDeleteSuccessRes{}, err
	    }
        
        err = json.Unmarshal(response, &deleteSmsTemplateByIdResponse)
        if err != nil {
            return SmsTemplateDeleteSuccessRes{}, common.NewFDKError(err.Error())
        }
        return deleteSmsTemplateByIdResponse, nil
        
    }
           
       
    
    
    
  

    
    //PlatformAppGetSystemSystemTemplatesXQuery holds query params
    type PlatformAppGetSystemSystemTemplatesXQuery struct { 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"` 
        Sort map[string]interface{}  `url:"sort,omitempty"`  
    }
    
    // GetSystemSystemTemplates Get system sms templates
     func (co *PlatformAppCommunication)  GetSystemSystemTemplates(xQuery PlatformAppGetSystemSystemTemplatesXQuery) (SystemSmsTemplates, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getSystemSystemTemplatesResponse SystemSmsTemplates
	    )

        

         
            
                
            
                
            
                
            
        

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            fmt.Sprintf("/service/platform/communication/v1.0/company/%s/application/%s/sms/system-templates",co.CompanyID, co.ApplicationID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return SystemSmsTemplates{}, err
	    }
        
        err = json.Unmarshal(response, &getSystemSystemTemplatesResponse)
        if err != nil {
            return SystemSmsTemplates{}, common.NewFDKError(err.Error())
        }
        return getSystemSystemTemplatesResponse, nil
        
    }
           
            
            
            
            
            
            
            
             
            
            
             
            
            
             
            
            
             
            
            
             
            
            
            
            
            
            
            
            
            
            
            
            
            // GetSystemSystemTemplatesPaginator Get system sms templates  
            func (co *PlatformAppCommunication)  GetSystemSystemTemplatesPaginator( xQuery PlatformAppGetSystemSystemTemplatesXQuery ) *common.Paginator {
                paginator := common.NewPaginator("number")
                
                
                 xQuery.PageNo  = paginator.PageNo
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                paginator.Next = func() (interface{}, error) {
                    response, err := co.GetSystemSystemTemplates(xQuery)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, int(response.Page.Current+1), response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }
        
       
    

 
	 
   // PlatformAppPayment holds PlatformAppPayment object properties
    type PlatformAppPayment struct {
        config *PlatformConfig
        CompanyID string
        ApplicationID string
    }
    // NewPlatformAppPayment returns new PlatformAppPayment instance
    func NewPlatformAppPayment(config *PlatformConfig, appID string) *PlatformAppPayment {
        return &PlatformAppPayment{config, config.CompanyID, appID}
    }
    
    
    
  

    
    // GetBrandPaymentGatewayConfig Get All Brand Payment Gateway Config Secret
     func (pa *PlatformAppPayment)  GetBrandPaymentGatewayConfig() (PaymentGatewayConfigResponse, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getBrandPaymentGatewayConfigResponse PaymentGatewayConfigResponse
	    )

        

         

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            pa.config,
            "get",
            fmt.Sprintf("/service/platform/payment/v1.0/company/%s/application/%s/aggregator/request",pa.CompanyID, pa.ApplicationID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return PaymentGatewayConfigResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getBrandPaymentGatewayConfigResponse)
        if err != nil {
            return PaymentGatewayConfigResponse{}, common.NewFDKError(err.Error())
        }
        return getBrandPaymentGatewayConfigResponse, nil
        
    }
           
       
    
    
    
  

    
    // SaveBrandPaymentGatewayConfig Save Config Secret For Brand Payment Gateway
     func (pa *PlatformAppPayment)  SaveBrandPaymentGatewayConfig(body  PaymentGatewayConfigRequest) (PaymentGatewayToBeReviewed, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            saveBrandPaymentGatewayConfigResponse PaymentGatewayToBeReviewed
	    )

        
            
        
            
        
            
        

         

        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return PaymentGatewayToBeReviewed{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return PaymentGatewayToBeReviewed{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            pa.config,
            "post",
            fmt.Sprintf("/service/platform/payment/v1.0/company/%s/application/%s/aggregator/request",pa.CompanyID, pa.ApplicationID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return PaymentGatewayToBeReviewed{}, err
	    }
        
        err = json.Unmarshal(response, &saveBrandPaymentGatewayConfigResponse)
        if err != nil {
            return PaymentGatewayToBeReviewed{}, common.NewFDKError(err.Error())
        }
        return saveBrandPaymentGatewayConfigResponse, nil
        
    }
           
       
    
    
    
  

    
    // UpdateBrandPaymentGatewayConfig Save Config Secret For Brand Payment Gateway
     func (pa *PlatformAppPayment)  UpdateBrandPaymentGatewayConfig(body  PaymentGatewayConfigRequest) (PaymentGatewayToBeReviewed, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            updateBrandPaymentGatewayConfigResponse PaymentGatewayToBeReviewed
	    )

        
            
        
            
        
            
        

         

        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return PaymentGatewayToBeReviewed{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return PaymentGatewayToBeReviewed{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            pa.config,
            "put",
            fmt.Sprintf("/service/platform/payment/v1.0/company/%s/application/%s/aggregator/request",pa.CompanyID, pa.ApplicationID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return PaymentGatewayToBeReviewed{}, err
	    }
        
        err = json.Unmarshal(response, &updateBrandPaymentGatewayConfigResponse)
        if err != nil {
            return PaymentGatewayToBeReviewed{}, common.NewFDKError(err.Error())
        }
        return updateBrandPaymentGatewayConfigResponse, nil
        
    }
           
       
    
    
    
  

    
    //PlatformAppGetPaymentModeRoutesXQuery holds query params
    type PlatformAppGetPaymentModeRoutesXQuery struct { 
        Refresh bool  `url:"refresh,omitempty"` 
        RequestType string  `url:"request_type,omitempty"`  
    }
    
    // GetPaymentModeRoutes Get All Valid Payment Options
     func (pa *PlatformAppPayment)  GetPaymentModeRoutes(xQuery PlatformAppGetPaymentModeRoutesXQuery) (PaymentOptionsResponse, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getPaymentModeRoutesResponse PaymentOptionsResponse
	    )

        

         
            
                
            
                
            
        

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            pa.config,
            "get",
            fmt.Sprintf("/service/platform/payment/v1.0/company/%s/application/%s/payment/options",pa.CompanyID, pa.ApplicationID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return PaymentOptionsResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getPaymentModeRoutesResponse)
        if err != nil {
            return PaymentOptionsResponse{}, common.NewFDKError(err.Error())
        }
        return getPaymentModeRoutesResponse, nil
        
    }
           
       
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
  

    
    // AddBeneficiaryDetails Save bank details for cancelled/returned order
     func (pa *PlatformAppPayment)  AddBeneficiaryDetails(body  AddBeneficiaryDetailsRequest) (RefundAccountResponse, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            addBeneficiaryDetailsResponse RefundAccountResponse
	    )

        
            
        
            
        
            
        
            
        
            
        
            
        
            
        

         

        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return RefundAccountResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return RefundAccountResponse{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            pa.config,
            "post",
            fmt.Sprintf("/service/platform/payment/v1.0/company/%s/application/%s/refund/account",pa.CompanyID, pa.ApplicationID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return RefundAccountResponse{}, err
	    }
        
        err = json.Unmarshal(response, &addBeneficiaryDetailsResponse)
        if err != nil {
            return RefundAccountResponse{}, common.NewFDKError(err.Error())
        }
        return addBeneficiaryDetailsResponse, nil
        
    }
           
       
    
    
    
    
    
  

    
    //PlatformAppGetUserOrderBeneficiariesXQuery holds query params
    type PlatformAppGetUserOrderBeneficiariesXQuery struct { 
        OrderID string  `url:"order_id,omitempty"`  
    }
    
    // GetUserOrderBeneficiaries List Order Beneficiary
     func (pa *PlatformAppPayment)  GetUserOrderBeneficiaries(xQuery PlatformAppGetUserOrderBeneficiariesXQuery) (OrderBeneficiaryResponse, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getUserOrderBeneficiariesResponse OrderBeneficiaryResponse
	    )

        

         
            
                
            
        

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            pa.config,
            "get",
            fmt.Sprintf("/service/platform/payment/v1.0/company/%s/application/%s/refund/accounts/order",pa.CompanyID, pa.ApplicationID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return OrderBeneficiaryResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getUserOrderBeneficiariesResponse)
        if err != nil {
            return OrderBeneficiaryResponse{}, common.NewFDKError(err.Error())
        }
        return getUserOrderBeneficiariesResponse, nil
        
    }
           
       
    
    
    
  

    
    //PlatformAppGetUserBeneficiariesXQuery holds query params
    type PlatformAppGetUserBeneficiariesXQuery struct { 
        OrderID string  `url:"order_id,omitempty"`  
    }
    
    // GetUserBeneficiaries List User Beneficiary
     func (pa *PlatformAppPayment)  GetUserBeneficiaries(xQuery PlatformAppGetUserBeneficiariesXQuery) (OrderBeneficiaryResponse, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getUserBeneficiariesResponse OrderBeneficiaryResponse
	    )

        

         
            
                
            
        

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            pa.config,
            "get",
            fmt.Sprintf("/service/platform/payment/v1.0/company/%s/application/%s/refund/accounts/user",pa.CompanyID, pa.ApplicationID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return OrderBeneficiaryResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getUserBeneficiariesResponse)
        if err != nil {
            return OrderBeneficiaryResponse{}, common.NewFDKError(err.Error())
        }
        return getUserBeneficiariesResponse, nil
        
    }
           
       
    

 
	 
   // PlatformAppOrder holds PlatformAppOrder object properties
    type PlatformAppOrder struct {
        config *PlatformConfig
        CompanyID string
        ApplicationID string
    }
    // NewPlatformAppOrder returns new PlatformAppOrder instance
    func NewPlatformAppOrder(config *PlatformConfig, appID string) *PlatformAppOrder {
        return &PlatformAppOrder{config, config.CompanyID, appID}
    }
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
  

    
    // TrackShipmentPlatform Track Shipment by shipment id, for application based on application Id
     func (or *PlatformAppOrder)  TrackShipmentPlatform(ShipmentID string) (PlatformShipmentTrack, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            trackShipmentPlatformResponse PlatformShipmentTrack
	    )

        

         

        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            or.config,
            "get",
            fmt.Sprintf("/service/platform/order/v1.0/company/%s/application/%s/orders/shipments/%s/track",or.CompanyID, or.ApplicationID, ShipmentID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return PlatformShipmentTrack{}, err
	    }
        
        err = json.Unmarshal(response, &trackShipmentPlatformResponse)
        if err != nil {
            return PlatformShipmentTrack{}, common.NewFDKError(err.Error())
        }
        return trackShipmentPlatformResponse, nil
        
    }
           
       
    
    
    
  

    
    // TrackOrder Track Order by order id, for application based on application Id
     func (or *PlatformAppOrder)  TrackOrder(OrderID string) (PlatformOrderTrack, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            trackOrderResponse PlatformOrderTrack
	    )

        

         

        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            or.config,
            "post",
            fmt.Sprintf("/service/platform/order/v1.0/company/%s/application/%s/orders/%s/track",or.CompanyID, or.ApplicationID, OrderID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return PlatformOrderTrack{}, err
	    }
        
        err = json.Unmarshal(response, &trackOrderResponse)
        if err != nil {
            return PlatformOrderTrack{}, common.NewFDKError(err.Error())
        }
        return trackOrderResponse, nil
        
    }
           
       
    
    
    
  

    
    // FailedOrders Get all failed orders application wise
     func (or *PlatformAppOrder)  FailedOrders() (FailedOrders, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            failedOrdersResponse FailedOrders
	    )

        

         

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            or.config,
            "get",
            fmt.Sprintf("/service/platform/order/v1.0/company/%s/application/%s/orders/failed",or.CompanyID, or.ApplicationID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return FailedOrders{}, err
	    }
        
        err = json.Unmarshal(response, &failedOrdersResponse)
        if err != nil {
            return FailedOrders{}, common.NewFDKError(err.Error())
        }
        return failedOrdersResponse, nil
        
    }
           
       
    
    
    
  

    
    // ReprocessOrder Reprocess order by order id
     func (or *PlatformAppOrder)  ReprocessOrder(OrderID string) (UpdateOrderReprocessResponse, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            reprocessOrderResponse UpdateOrderReprocessResponse
	    )

        

         

        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            or.config,
            "post",
            fmt.Sprintf("/service/platform/order/v1.0/company/%s/application/%s/orders/%s/reprocess",or.CompanyID, or.ApplicationID, OrderID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return UpdateOrderReprocessResponse{}, err
	    }
        
        err = json.Unmarshal(response, &reprocessOrderResponse)
        if err != nil {
            return UpdateOrderReprocessResponse{}, common.NewFDKError(err.Error())
        }
        return reprocessOrderResponse, nil
        
    }
           
       
    
    
    
  

    
    // UpdateShipment Use this API to update the shipment using its shipment ID.
     func (or *PlatformAppOrder)  UpdateShipment(ShipmentID string, body  ShipmentUpdateRequest) (ShipmentUpdateResponse, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            updateShipmentResponse ShipmentUpdateResponse
	    )

        
            
        
            
        
            
        
            
        

         

        
        
        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return ShipmentUpdateResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return ShipmentUpdateResponse{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            or.config,
            "post",
            fmt.Sprintf("/service/platform/order/v1.0/company/%s/application/%s/orders/shipments/%s/update",or.CompanyID, or.ApplicationID, ShipmentID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return ShipmentUpdateResponse{}, err
	    }
        
        err = json.Unmarshal(response, &updateShipmentResponse)
        if err != nil {
            return ShipmentUpdateResponse{}, common.NewFDKError(err.Error())
        }
        return updateShipmentResponse, nil
        
    }
           
       
    
    
    
  

    
    // GetPlatformShipmentReasons Use this API to retrieve the issues that led to the cancellation of bags within a shipment.
     func (or *PlatformAppOrder)  GetPlatformShipmentReasons(Action string) (ShipmentReasonsResponse, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getPlatformShipmentReasonsResponse ShipmentReasonsResponse
	    )

        

         

        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            or.config,
            "get",
            fmt.Sprintf("/service/platform/order/v1.0/company/%s/application/%s/orders/shipments/reasons/%s",or.CompanyID, or.ApplicationID, Action),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return ShipmentReasonsResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getPlatformShipmentReasonsResponse)
        if err != nil {
            return ShipmentReasonsResponse{}, common.NewFDKError(err.Error())
        }
        return getPlatformShipmentReasonsResponse, nil
        
    }
           
       
    
    
    
  

    
    // GetShipmentTrackDetails Use this API to track a shipment using its shipment ID.
     func (or *PlatformAppOrder)  GetShipmentTrackDetails(OrderID string, ShipmentID string) (ShipmentTrackResponse, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getShipmentTrackDetailsResponse ShipmentTrackResponse
	    )

        

         

        
        
        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            or.config,
            "get",
            fmt.Sprintf("/service/platform/order/v1.0/company/%s/application/%s/orders/%s/shipments/%s/track",or.CompanyID, or.ApplicationID, OrderID, ShipmentID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return ShipmentTrackResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getShipmentTrackDetailsResponse)
        if err != nil {
            return ShipmentTrackResponse{}, common.NewFDKError(err.Error())
        }
        return getShipmentTrackDetailsResponse, nil
        
    }
           
       
    
    
    
    
    
    
    
    
    
    
    

 
	 
   // PlatformAppCatalog holds PlatformAppCatalog object properties
    type PlatformAppCatalog struct {
        config *PlatformConfig
        CompanyID string
        ApplicationID string
    }
    // NewPlatformAppCatalog returns new PlatformAppCatalog instance
    func NewPlatformAppCatalog(config *PlatformConfig, appID string) *PlatformAppCatalog {
        return &PlatformAppCatalog{config, config.CompanyID, appID}
    }
    
    
    
  

    
    // DeleteSearchKeywords Delete a Search Keywords
     func (ca *PlatformAppCatalog)  DeleteSearchKeywords(ID string) (DeleteResponse, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            deleteSearchKeywordsResponse DeleteResponse
	    )

        

         

        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "delete",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/application/%s/search/keyword/%s/",ca.CompanyID, ca.ApplicationID, ID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return DeleteResponse{}, err
	    }
        
        err = json.Unmarshal(response, &deleteSearchKeywordsResponse)
        if err != nil {
            return DeleteResponse{}, common.NewFDKError(err.Error())
        }
        return deleteSearchKeywordsResponse, nil
        
    }
           
       
    
    
    
  

    
    // UpdateSearchKeywords Update Search Keyword
     func (ca *PlatformAppCatalog)  UpdateSearchKeywords(ID string, body  CreateSearchKeyword) (GetSearchWordsData, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            updateSearchKeywordsResponse GetSearchWordsData
	    )

        
            
        
            
        
            
        
            
        
            
        

         

        
        
        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return GetSearchWordsData{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return GetSearchWordsData{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "put",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/application/%s/search/keyword/%s/",ca.CompanyID, ca.ApplicationID, ID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return GetSearchWordsData{}, err
	    }
        
        err = json.Unmarshal(response, &updateSearchKeywordsResponse)
        if err != nil {
            return GetSearchWordsData{}, common.NewFDKError(err.Error())
        }
        return updateSearchKeywordsResponse, nil
        
    }
           
       
    
    
    
  

    
    // GetSearchKeywords Get a Search Keywords Details
     func (ca *PlatformAppCatalog)  GetSearchKeywords(ID string) (GetSearchWordsDetailResponse, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getSearchKeywordsResponse GetSearchWordsDetailResponse
	    )

        

         

        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "get",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/application/%s/search/keyword/%s/",ca.CompanyID, ca.ApplicationID, ID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return GetSearchWordsDetailResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getSearchKeywordsResponse)
        if err != nil {
            return GetSearchWordsDetailResponse{}, common.NewFDKError(err.Error())
        }
        return getSearchKeywordsResponse, nil
        
    }
           
       
    
    
    
  

    
    // GetAllSearchKeyword List all Search Custom Keyword Listing
     func (ca *PlatformAppCatalog)  GetAllSearchKeyword() (GetSearchWordsResponse, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getAllSearchKeywordResponse GetSearchWordsResponse
	    )

        

         

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "get",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/application/%s/search/keyword/",ca.CompanyID, ca.ApplicationID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return GetSearchWordsResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getAllSearchKeywordResponse)
        if err != nil {
            return GetSearchWordsResponse{}, common.NewFDKError(err.Error())
        }
        return getAllSearchKeywordResponse, nil
        
    }
           
       
    
    
    
  

    
    // CreateCustomKeyword Add a Custom Search Keywords
     func (ca *PlatformAppCatalog)  CreateCustomKeyword(body  CreateSearchKeyword) (GetSearchWordsData, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            createCustomKeywordResponse GetSearchWordsData
	    )

        
            
        
            
        
            
        
            
        
            
        

         

        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return GetSearchWordsData{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return GetSearchWordsData{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "post",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/application/%s/search/keyword/",ca.CompanyID, ca.ApplicationID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return GetSearchWordsData{}, err
	    }
        
        err = json.Unmarshal(response, &createCustomKeywordResponse)
        if err != nil {
            return GetSearchWordsData{}, common.NewFDKError(err.Error())
        }
        return createCustomKeywordResponse, nil
        
    }
           
       
    
    
    
  

    
    // DeleteAutocompleteKeyword Delete a Autocomplete Keywords
     func (ca *PlatformAppCatalog)  DeleteAutocompleteKeyword(ID string) (DeleteResponse, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            deleteAutocompleteKeywordResponse DeleteResponse
	    )

        

         

        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "delete",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/application/%s/search/autocomplete/%s/",ca.CompanyID, ca.ApplicationID, ID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return DeleteResponse{}, err
	    }
        
        err = json.Unmarshal(response, &deleteAutocompleteKeywordResponse)
        if err != nil {
            return DeleteResponse{}, common.NewFDKError(err.Error())
        }
        return deleteAutocompleteKeywordResponse, nil
        
    }
           
       
    
    
    
  

    
    // UpdateAutocompleteKeyword Create & Update Autocomplete Keyword
     func (ca *PlatformAppCatalog)  UpdateAutocompleteKeyword(ID string, body  CreateAutocompleteKeyword) (GetAutocompleteWordsResponse, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            updateAutocompleteKeywordResponse GetAutocompleteWordsResponse
	    )

        
            
        
            
        
            
        
            
        
            
        

         

        
        
        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return GetAutocompleteWordsResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return GetAutocompleteWordsResponse{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "put",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/application/%s/search/autocomplete/%s/",ca.CompanyID, ca.ApplicationID, ID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return GetAutocompleteWordsResponse{}, err
	    }
        
        err = json.Unmarshal(response, &updateAutocompleteKeywordResponse)
        if err != nil {
            return GetAutocompleteWordsResponse{}, common.NewFDKError(err.Error())
        }
        return updateAutocompleteKeywordResponse, nil
        
    }
           
       
    
    
    
  

    
    // GetAutocompleteKeywordDetail Get a Autocomplete Keywords Details
     func (ca *PlatformAppCatalog)  GetAutocompleteKeywordDetail(ID string) (GetAutocompleteWordsResponse, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getAutocompleteKeywordDetailResponse GetAutocompleteWordsResponse
	    )

        

         

        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "get",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/application/%s/search/autocomplete/%s/",ca.CompanyID, ca.ApplicationID, ID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return GetAutocompleteWordsResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getAutocompleteKeywordDetailResponse)
        if err != nil {
            return GetAutocompleteWordsResponse{}, common.NewFDKError(err.Error())
        }
        return getAutocompleteKeywordDetailResponse, nil
        
    }
           
       
    
    
    
  

    
    // GetAutocompleteConfig List all Autocomplete Keyword Listing
     func (ca *PlatformAppCatalog)  GetAutocompleteConfig() (GetAutocompleteWordsResponse, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getAutocompleteConfigResponse GetAutocompleteWordsResponse
	    )

        

         

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "get",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/application/%s/search/autocomplete/",ca.CompanyID, ca.ApplicationID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return GetAutocompleteWordsResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getAutocompleteConfigResponse)
        if err != nil {
            return GetAutocompleteWordsResponse{}, common.NewFDKError(err.Error())
        }
        return getAutocompleteConfigResponse, nil
        
    }
           
       
    
    
    
  

    
    // CreateCustomAutocompleteRule Add a Custom Autocomplete Keywords
     func (ca *PlatformAppCatalog)  CreateCustomAutocompleteRule(body  CreateAutocompleteKeyword) (CreateAutocompleteWordsResponse, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            createCustomAutocompleteRuleResponse CreateAutocompleteWordsResponse
	    )

        
            
        
            
        
            
        
            
        
            
        

         

        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return CreateAutocompleteWordsResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return CreateAutocompleteWordsResponse{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "post",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/application/%s/search/autocomplete/",ca.CompanyID, ca.ApplicationID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return CreateAutocompleteWordsResponse{}, err
	    }
        
        err = json.Unmarshal(response, &createCustomAutocompleteRuleResponse)
        if err != nil {
            return CreateAutocompleteWordsResponse{}, common.NewFDKError(err.Error())
        }
        return createCustomAutocompleteRuleResponse, nil
        
    }
           
       
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
  

    
    // GetCatalogConfiguration Get configuration meta  details for catalog for admin panel
     func (ca *PlatformAppCatalog)  GetCatalogConfiguration() (GetCatalogConfigurationMetaData, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getCatalogConfigurationResponse GetCatalogConfigurationMetaData
	    )

        

         

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "get",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/application/%s/product-configuration/metadata/",ca.CompanyID, ca.ApplicationID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return GetCatalogConfigurationMetaData{}, err
	    }
        
        err = json.Unmarshal(response, &getCatalogConfigurationResponse)
        if err != nil {
            return GetCatalogConfigurationMetaData{}, common.NewFDKError(err.Error())
        }
        return getCatalogConfigurationResponse, nil
        
    }
           
       
    
    
    
  

    
    // GetConfigurations Get configured details for catalog
     func (ca *PlatformAppCatalog)  GetConfigurations() (GetAppCatalogConfiguration, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getConfigurationsResponse GetAppCatalogConfiguration
	    )

        

         

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "get",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/application/%s/product-configuration/",ca.CompanyID, ca.ApplicationID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return GetAppCatalogConfiguration{}, err
	    }
        
        err = json.Unmarshal(response, &getConfigurationsResponse)
        if err != nil {
            return GetAppCatalogConfiguration{}, common.NewFDKError(err.Error())
        }
        return getConfigurationsResponse, nil
        
    }
           
       
    
    
    
  

    
    // CreateConfigurationProductListing Add configuration for products & listings
     func (ca *PlatformAppCatalog)  CreateConfigurationProductListing(body  AppConfiguration) (GetAppCatalogConfiguration, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            createConfigurationProductListingResponse GetAppCatalogConfiguration
	    )

        
            
        
            
        
            
        
            
        
            
        

         

        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return GetAppCatalogConfiguration{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return GetAppCatalogConfiguration{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "post",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/application/%s/product-configuration/",ca.CompanyID, ca.ApplicationID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return GetAppCatalogConfiguration{}, err
	    }
        
        err = json.Unmarshal(response, &createConfigurationProductListingResponse)
        if err != nil {
            return GetAppCatalogConfiguration{}, common.NewFDKError(err.Error())
        }
        return createConfigurationProductListingResponse, nil
        
    }
           
       
    
    
    
  

    
    // GetConfigurationByType Get configured details for catalog
     func (ca *PlatformAppCatalog)  GetConfigurationByType(Type string) (GetAppCatalogEntityConfiguration, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getConfigurationByTypeResponse GetAppCatalogEntityConfiguration
	    )

        

         

        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "get",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/application/%s/product-configuration/%s/",ca.CompanyID, ca.ApplicationID, Type),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return GetAppCatalogEntityConfiguration{}, err
	    }
        
        err = json.Unmarshal(response, &getConfigurationByTypeResponse)
        if err != nil {
            return GetAppCatalogEntityConfiguration{}, common.NewFDKError(err.Error())
        }
        return getConfigurationByTypeResponse, nil
        
    }
           
       
    
    
    
  

    
    // CreateConfigurationByType Add configuration for categories and brands
     func (ca *PlatformAppCatalog)  CreateConfigurationByType(Type string, body  AppConfiguration) (GetAppCatalogConfiguration, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            createConfigurationByTypeResponse GetAppCatalogConfiguration
	    )

        
            
        
            
        
            
        
            
        
            
        

         

        
        
        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return GetAppCatalogConfiguration{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return GetAppCatalogConfiguration{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "post",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/application/%s/product-configuration/%s/",ca.CompanyID, ca.ApplicationID, Type),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return GetAppCatalogConfiguration{}, err
	    }
        
        err = json.Unmarshal(response, &createConfigurationByTypeResponse)
        if err != nil {
            return GetAppCatalogConfiguration{}, common.NewFDKError(err.Error())
        }
        return createConfigurationByTypeResponse, nil
        
    }
           
       
    
    
    
  

    
    // GetQueryFilters Get query filters to configure a collection
     func (ca *PlatformAppCatalog)  GetQueryFilters() (GetCollectionQueryOptionResponse, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getQueryFiltersResponse GetCollectionQueryOptionResponse
	    )

        

         

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "get",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/application/%s/collections/query-options/",ca.CompanyID, ca.ApplicationID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return GetCollectionQueryOptionResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getQueryFiltersResponse)
        if err != nil {
            return GetCollectionQueryOptionResponse{}, common.NewFDKError(err.Error())
        }
        return getQueryFiltersResponse, nil
        
    }
           
       
    
    
    
  

    
    // GetAllCollections List all the collections
     func (ca *PlatformAppCatalog)  GetAllCollections() (GetCollectionListingResponse, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getAllCollectionsResponse GetCollectionListingResponse
	    )

        

         

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "get",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/application/%s/collections/",ca.CompanyID, ca.ApplicationID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return GetCollectionListingResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getAllCollectionsResponse)
        if err != nil {
            return GetCollectionListingResponse{}, common.NewFDKError(err.Error())
        }
        return getAllCollectionsResponse, nil
        
    }
           
       
    
    
    
  

    
    // CreateCollection Add a Collection
     func (ca *PlatformAppCatalog)  CreateCollection(body  CreateCollection) (CollectionCreateResponse, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            createCollectionResponse CollectionCreateResponse
	    )

        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        

         

        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return CollectionCreateResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return CollectionCreateResponse{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "post",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/application/%s/collections/",ca.CompanyID, ca.ApplicationID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return CollectionCreateResponse{}, err
	    }
        
        err = json.Unmarshal(response, &createCollectionResponse)
        if err != nil {
            return CollectionCreateResponse{}, common.NewFDKError(err.Error())
        }
        return createCollectionResponse, nil
        
    }
           
       
    
    
    
  

    
    // GetCollectionDetail Get a particular collection
     func (ca *PlatformAppCatalog)  GetCollectionDetail(Slug string) (CollectionDetailResponse, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getCollectionDetailResponse CollectionDetailResponse
	    )

        

         

        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "get",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/application/%s/collections/%s/",ca.CompanyID, ca.ApplicationID, Slug),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return CollectionDetailResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getCollectionDetailResponse)
        if err != nil {
            return CollectionDetailResponse{}, common.NewFDKError(err.Error())
        }
        return getCollectionDetailResponse, nil
        
    }
           
       
    
    
    
  

    
    // UpdateCollection Update a collection
     func (ca *PlatformAppCatalog)  UpdateCollection(ID string, body  UpdateCollection) (UpdateCollection, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            updateCollectionResponse UpdateCollection
	    )

        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        

         

        
        
        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return UpdateCollection{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return UpdateCollection{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "put",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/application/%s/collections/%s/",ca.CompanyID, ca.ApplicationID, ID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return UpdateCollection{}, err
	    }
        
        err = json.Unmarshal(response, &updateCollectionResponse)
        if err != nil {
            return UpdateCollection{}, common.NewFDKError(err.Error())
        }
        return updateCollectionResponse, nil
        
    }
           
       
    
    
    
  

    
    // DeleteCollection Delete a Collection
     func (ca *PlatformAppCatalog)  DeleteCollection(ID string) (DeleteResponse, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            deleteCollectionResponse DeleteResponse
	    )

        

         

        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "delete",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/application/%s/collections/%s/",ca.CompanyID, ca.ApplicationID, ID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return DeleteResponse{}, err
	    }
        
        err = json.Unmarshal(response, &deleteCollectionResponse)
        if err != nil {
            return DeleteResponse{}, common.NewFDKError(err.Error())
        }
        return deleteCollectionResponse, nil
        
    }
           
       
    
    
    
  

    
    //PlatformAppGetCollectionItemsXQuery holds query params
    type PlatformAppGetCollectionItemsXQuery struct { 
        SortOn string  `url:"sort_on,omitempty"` 
        PageID string  `url:"page_id,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"`  
    }
    
    // GetCollectionItems Get the items for a collection
     func (ca *PlatformAppCatalog)  GetCollectionItems(ID string, xQuery PlatformAppGetCollectionItemsXQuery) (GetCollectionItemsResponse, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getCollectionItemsResponse GetCollectionItemsResponse
	    )

        

         
            
                
            
                
            
                
            
        

        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "get",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/application/%s/collections/%s/items/",ca.CompanyID, ca.ApplicationID, ID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return GetCollectionItemsResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getCollectionItemsResponse)
        if err != nil {
            return GetCollectionItemsResponse{}, common.NewFDKError(err.Error())
        }
        return getCollectionItemsResponse, nil
        
    }
           
       
    
    
    
  

    
    // AddCollectionItems Add items to a collection
     func (ca *PlatformAppCatalog)  AddCollectionItems(ID string, body  CollectionItemRequest) (UpdatedResponse, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            addCollectionItemsResponse UpdatedResponse
	    )

        
            
        
            
        

         

        
        
        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return UpdatedResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return UpdatedResponse{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "post",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/application/%s/collections/%s/items/",ca.CompanyID, ca.ApplicationID, ID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return UpdatedResponse{}, err
	    }
        
        err = json.Unmarshal(response, &addCollectionItemsResponse)
        if err != nil {
            return UpdatedResponse{}, common.NewFDKError(err.Error())
        }
        return addCollectionItemsResponse, nil
        
    }
           
       
    
    
    
  

    
    //PlatformAppGetCatalogInsightsXQuery holds query params
    type PlatformAppGetCatalogInsightsXQuery struct { 
        Brand string  `url:"brand,omitempty"`  
    }
    
    // GetCatalogInsights Analytics data of catalog and inventory.
     func (ca *PlatformAppCatalog)  GetCatalogInsights(xQuery PlatformAppGetCatalogInsightsXQuery) (CatalogInsightResponse, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getCatalogInsightsResponse CatalogInsightResponse
	    )

        

         
            
                
            
        

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "get",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/application/%s/analytics/insights/",ca.CompanyID, ca.ApplicationID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return CatalogInsightResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getCatalogInsightsResponse)
        if err != nil {
            return CatalogInsightResponse{}, common.NewFDKError(err.Error())
        }
        return getCatalogInsightsResponse, nil
        
    }
           
       
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
  

    
    //PlatformAppGetApplicationBrandsXQuery holds query params
    type PlatformAppGetApplicationBrandsXQuery struct { 
        Department string  `url:"department,omitempty"` 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"`  
    }
    
    // GetApplicationBrands List all the brands
     func (ca *PlatformAppCatalog)  GetApplicationBrands(xQuery PlatformAppGetApplicationBrandsXQuery) (BrandListingResponse, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getApplicationBrandsResponse BrandListingResponse
	    )

        

         
            
                
            
                
            
                
            
        

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "get",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/application/%s/brands",ca.CompanyID, ca.ApplicationID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return BrandListingResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getApplicationBrandsResponse)
        if err != nil {
            return BrandListingResponse{}, common.NewFDKError(err.Error())
        }
        return getApplicationBrandsResponse, nil
        
    }
           
            
            
            
            
            
            
            
             
            
            
             
            
            
             
            
            
             
            
            
             
            
            
            
            
            
            
            
            
            
            
            
            
            // GetApplicationBrandsPaginator List all the brands  
            func (ca *PlatformAppCatalog)  GetApplicationBrandsPaginator( xQuery PlatformAppGetApplicationBrandsXQuery ) *common.Paginator {
                paginator := common.NewPaginator("number")
                
                
                 
                 
                 
                 
                 xQuery.PageNo  = paginator.PageNo
                 
                 
                 
                 
                 
                 
                 
                 
                paginator.Next = func() (interface{}, error) {
                    response, err := ca.GetApplicationBrands(xQuery)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, int(response.Page.Current+1), response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }
        
       
    
    
    
  

    
    // GetDepartments List all the departments
     func (ca *PlatformAppCatalog)  GetDepartments() (DepartmentResponse, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getDepartmentsResponse DepartmentResponse
	    )

        

         

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "get",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/application/%s/departments",ca.CompanyID, ca.ApplicationID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return DepartmentResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getDepartmentsResponse)
        if err != nil {
            return DepartmentResponse{}, common.NewFDKError(err.Error())
        }
        return getDepartmentsResponse, nil
        
    }
           
       
    
    
    
  

    
    //PlatformAppGetCategoriesXQuery holds query params
    type PlatformAppGetCategoriesXQuery struct { 
        Department string  `url:"department,omitempty"`  
    }
    
    // GetCategories List all the categories
     func (ca *PlatformAppCatalog)  GetCategories(xQuery PlatformAppGetCategoriesXQuery) (CategoryListingResponse, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getCategoriesResponse CategoryListingResponse
	    )

        

         
            
                
            
        

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "get",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/application/%s/categories",ca.CompanyID, ca.ApplicationID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return CategoryListingResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getCategoriesResponse)
        if err != nil {
            return CategoryListingResponse{}, common.NewFDKError(err.Error())
        }
        return getCategoriesResponse, nil
        
    }
           
       
    
    
    
  

    
    //PlatformAppGetAppicationProductsXQuery holds query params
    type PlatformAppGetAppicationProductsXQuery struct { 
        Q string  `url:"q,omitempty"` 
        F string  `url:"f,omitempty"` 
        Filters bool  `url:"filters,omitempty"` 
        SortOn string  `url:"sort_on,omitempty"` 
        PageID string  `url:"page_id,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"` 
        PageNo float64  `url:"page_no,omitempty"` 
        PageType string  `url:"page_type,omitempty"`  
    }
    
    // GetAppicationProducts List the products
     func (ca *PlatformAppCatalog)  GetAppicationProducts(xQuery PlatformAppGetAppicationProductsXQuery) (ApplicationProductListingResponse, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getAppicationProductsResponse ApplicationProductListingResponse
	    )

        

         
            
                
            
                
            
                
            
                
            
                
            
                
            
                
            
                
            
        

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "get",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/application/%s/products",ca.CompanyID, ca.ApplicationID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return ApplicationProductListingResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getAppicationProductsResponse)
        if err != nil {
            return ApplicationProductListingResponse{}, common.NewFDKError(err.Error())
        }
        return getAppicationProductsResponse, nil
        
    }
           
            
            
            
            
            
            
            
             
            
            
             
            
            
             
            
            
             
            
            
             
            
            
             
            
            
             
            
            
             
            
            
             
            
            
             
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            // GetAppicationProductsPaginator List the products  
            func (ca *PlatformAppCatalog)  GetAppicationProductsPaginator( xQuery PlatformAppGetAppicationProductsXQuery ) *common.Paginator {
                paginator := common.NewPaginator("cursor")
                
                
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 xQuery.PageID = paginator.NextID
                 
                 
                 
                 
                 
                 
                 xQuery.PageNo  = paginator.PageNo
                 
                 
                 
                 
                 
                 xQuery.PageType = "cursor"
                 
                 
                paginator.Next = func() (interface{}, error) {
                    response, err := ca.GetAppicationProducts(xQuery)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, int(response.Page.Current+1), response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }
        
       
    
    
    
  

    
    // GetProductDetailBySlug Get a product
     func (ca *PlatformAppCatalog)  GetProductDetailBySlug(Slug string) (ProductDetail, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getProductDetailBySlugResponse ProductDetail
	    )

        

         

        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "get",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/application/%s/products/%s",ca.CompanyID, ca.ApplicationID, Slug),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return ProductDetail{}, err
	    }
        
        err = json.Unmarshal(response, &getProductDetailBySlugResponse)
        if err != nil {
            return ProductDetail{}, common.NewFDKError(err.Error())
        }
        return getProductDetailBySlugResponse, nil
        
    }
           
       
    

 
	 
   // PlatformAppFileStorage holds PlatformAppFileStorage object properties
    type PlatformAppFileStorage struct {
        config *PlatformConfig
        CompanyID string
        ApplicationID string
    }
    // NewPlatformAppFileStorage returns new PlatformAppFileStorage instance
    func NewPlatformAppFileStorage(config *PlatformConfig, appID string) *PlatformAppFileStorage {
        return &PlatformAppFileStorage{config, config.CompanyID, appID}
    }
    
    
    
    
    
    
    
  

    
    // AppStartUpload This operation initiates upload and returns storage link which is valid for 30 Minutes. You can use that storage link to make subsequent upload request with file buffer or blob.
     func (fi *PlatformAppFileStorage)  AppStartUpload(Namespace string, body  StartRequest) (StartResponse, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            appStartUploadResponse StartResponse
	    )

        
            
        
            
        
            
        
            
        
            
        

         

        
        
        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return StartResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return StartResponse{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            fi.config,
            "post",
            fmt.Sprintf("/service/platform/assets/v1.0/company/%s/application/%s/namespaces/%s/upload/start/",Namespace, fi.CompanyID, fi.ApplicationID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return StartResponse{}, err
	    }
        
        err = json.Unmarshal(response, &appStartUploadResponse)
        if err != nil {
            return StartResponse{}, common.NewFDKError(err.Error())
        }
        return appStartUploadResponse, nil
        
    }
           
       
    
    
    
  

    
    // AppCompleteUpload This will complete the upload process. After successfully uploading file, you can call this operation to complete the upload process.
     func (fi *PlatformAppFileStorage)  AppCompleteUpload(Namespace string, body  StartResponse) (CompleteResponse, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            appCompleteUploadResponse CompleteResponse
	    )

        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        

         

        
        
        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return CompleteResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return CompleteResponse{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            fi.config,
            "post",
            fmt.Sprintf("/service/platform/assets/v1.0/company/%s/application/%s/namespaces/%s/upload/complete/",Namespace, fi.CompanyID, fi.ApplicationID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return CompleteResponse{}, err
	    }
        
        err = json.Unmarshal(response, &appCompleteUploadResponse)
        if err != nil {
            return CompleteResponse{}, common.NewFDKError(err.Error())
        }
        return appCompleteUploadResponse, nil
        
    }
           
       
    
    
    
    
    
    
    
  

    
    //PlatformAppAppCopyFilesXQuery holds query params
    type PlatformAppAppCopyFilesXQuery struct { 
        Sync bool  `url:"sync,omitempty"`  
    }
    
    // AppCopyFiles Copy Files
     func (fi *PlatformAppFileStorage)  AppCopyFiles(xQuery PlatformAppAppCopyFilesXQuery, body  BulkRequest) (BulkResponse, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            appCopyFilesResponse BulkResponse
	    )

        
            
        
            
        
            
        

         
            
                
            
        

        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return BulkResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return BulkResponse{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            fi.config,
            "post",
            fmt.Sprintf("/service/platform/assets/v1.0/company/%s/application/%s/uploads/copy/",fi.CompanyID, fi.ApplicationID),
            nil,
            xQuery,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return BulkResponse{}, err
	    }
        
        err = json.Unmarshal(response, &appCopyFilesResponse)
        if err != nil {
            return BulkResponse{}, common.NewFDKError(err.Error())
        }
        return appCopyFilesResponse, nil
        
    }
           
       
    
    
    
    
    
  

    
    //PlatformAppBrowseXQuery holds query params
    type PlatformAppBrowseXQuery struct { 
        PageNo float64  `url:"page_no,omitempty"`  
    }
    
    // Browse Browse Files
     func (fi *PlatformAppFileStorage)  Browse(Namespace string, xQuery PlatformAppBrowseXQuery) (BrowseResponse, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            browseResponse BrowseResponse
	    )

        

         
            
                
            
        

        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            fi.config,
            "get",
            fmt.Sprintf("/service/platform/assets/v1.0/company/%s/application/%s/namespaces/%s/browse/",Namespace, fi.CompanyID, fi.ApplicationID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return BrowseResponse{}, err
	    }
        
        err = json.Unmarshal(response, &browseResponse)
        if err != nil {
            return BrowseResponse{}, common.NewFDKError(err.Error())
        }
        return browseResponse, nil
        
    }
           
            
            
            
            
            
            
            
             
            
            
            
             
            
            
             
            
            
             
            
            
            
            
            
            
            // BrowsePaginator Browse Files  
            func (fi *PlatformAppFileStorage)  BrowsePaginator(Namespace string ,  xQuery PlatformAppBrowseXQuery ) *common.Paginator {
                paginator := common.NewPaginator("number")
                
                
                 xQuery.PageNo  = paginator.PageNo
                 
                 
                 
                 
                paginator.Next = func() (interface{}, error) {
                    response, err := fi.Browse(Namespace, xQuery)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, int(response.Page.Current+1), response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }
        
       
    
    
    

 
	 
   // PlatformAppShare holds PlatformAppShare object properties
    type PlatformAppShare struct {
        config *PlatformConfig
        CompanyID string
        ApplicationID string
    }
    // NewPlatformAppShare returns new PlatformAppShare instance
    func NewPlatformAppShare(config *PlatformConfig, appID string) *PlatformAppShare {
        return &PlatformAppShare{config, config.CompanyID, appID}
    }
    
    
    
  

    
    // CreateShortLink Create short link
     func (sh *PlatformAppShare)  CreateShortLink(body  ShortLinkReq) (ShortLinkRes, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            createShortLinkResponse ShortLinkRes
	    )

        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        

         

        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return ShortLinkRes{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return ShortLinkRes{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            sh.config,
            "post",
            fmt.Sprintf("/service/platform/share/v1.0/company/%s/application/%s/links/short-link/",sh.CompanyID, sh.ApplicationID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return ShortLinkRes{}, err
	    }
        
        err = json.Unmarshal(response, &createShortLinkResponse)
        if err != nil {
            return ShortLinkRes{}, common.NewFDKError(err.Error())
        }
        return createShortLinkResponse, nil
        
    }
           
       
    
    
    
  

    
    //PlatformAppGetShortLinksXQuery holds query params
    type PlatformAppGetShortLinksXQuery struct { 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"` 
        CreatedBy string  `url:"created_by,omitempty"` 
        Active string  `url:"active,omitempty"` 
        Q string  `url:"q,omitempty"`  
    }
    
    // GetShortLinks Get short links
     func (sh *PlatformAppShare)  GetShortLinks(xQuery PlatformAppGetShortLinksXQuery) (ShortLinkList, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getShortLinksResponse ShortLinkList
	    )

        

         
            
                
            
                
            
                
            
                
            
                
            
        

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            sh.config,
            "get",
            fmt.Sprintf("/service/platform/share/v1.0/company/%s/application/%s/links/short-link/",sh.CompanyID, sh.ApplicationID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return ShortLinkList{}, err
	    }
        
        err = json.Unmarshal(response, &getShortLinksResponse)
        if err != nil {
            return ShortLinkList{}, common.NewFDKError(err.Error())
        }
        return getShortLinksResponse, nil
        
    }
           
            
            
            
            
            
            
            
             
            
            
             
            
            
             
            
            
             
            
            
             
            
            
             
            
            
             
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            // GetShortLinksPaginator Get short links  
            func (sh *PlatformAppShare)  GetShortLinksPaginator( xQuery PlatformAppGetShortLinksXQuery ) *common.Paginator {
                paginator := common.NewPaginator("number")
                
                
                 xQuery.PageNo  = paginator.PageNo
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                paginator.Next = func() (interface{}, error) {
                    response, err := sh.GetShortLinks(xQuery)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, int(response.Page.Current+1), response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }
        
       
    
    
    
  

    
    // GetShortLinkByHash Get short link by hash
     func (sh *PlatformAppShare)  GetShortLinkByHash(Hash string) (ShortLinkRes, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getShortLinkByHashResponse ShortLinkRes
	    )

        

         

        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            sh.config,
            "get",
            fmt.Sprintf("/service/platform/share/v1.0/company/%s/application/%s/links/short-link/%s/",sh.CompanyID, sh.ApplicationID, Hash),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return ShortLinkRes{}, err
	    }
        
        err = json.Unmarshal(response, &getShortLinkByHashResponse)
        if err != nil {
            return ShortLinkRes{}, common.NewFDKError(err.Error())
        }
        return getShortLinkByHashResponse, nil
        
    }
           
       
    
    
    
  

    
    // UpdateShortLinkById Update short link by id
     func (sh *PlatformAppShare)  UpdateShortLinkById(ID string, body  ShortLinkReq) (ShortLinkRes, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            updateShortLinkByIdResponse ShortLinkRes
	    )

        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        

         

        
        
        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return ShortLinkRes{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return ShortLinkRes{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            sh.config,
            "patch",
            fmt.Sprintf("/service/platform/share/v1.0/company/%s/application/%s/links/short-link/%s/",sh.CompanyID, sh.ApplicationID, ID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return ShortLinkRes{}, err
	    }
        
        err = json.Unmarshal(response, &updateShortLinkByIdResponse)
        if err != nil {
            return ShortLinkRes{}, common.NewFDKError(err.Error())
        }
        return updateShortLinkByIdResponse, nil
        
    }
           
       
    

 
	 
   // PlatformAppConfiguration holds PlatformAppConfiguration object properties
    type PlatformAppConfiguration struct {
        config *PlatformConfig
        CompanyID string
        ApplicationID string
    }
    // NewPlatformAppConfiguration returns new PlatformAppConfiguration instance
    func NewPlatformAppConfiguration(config *PlatformConfig, appID string) *PlatformAppConfiguration {
        return &PlatformAppConfiguration{config, config.CompanyID, appID}
    }
    
    
    
  

    
    // GetBuildConfig Get latest build config
     func (co *PlatformAppConfiguration)  GetBuildConfig(PlatformType string) (MobileAppConfiguration, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getBuildConfigResponse MobileAppConfiguration
	    )

        

         

        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            fmt.Sprintf("/service/platform/configuration/v1.0/company/%s/application/%s/build/%s/configuration",co.CompanyID, co.ApplicationID, PlatformType),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return MobileAppConfiguration{}, err
	    }
        
        err = json.Unmarshal(response, &getBuildConfigResponse)
        if err != nil {
            return MobileAppConfiguration{}, common.NewFDKError(err.Error())
        }
        return getBuildConfigResponse, nil
        
    }
           
       
    
    
    
  

    
    // UpdateBuildConfig Update build config for next build
     func (co *PlatformAppConfiguration)  UpdateBuildConfig(PlatformType string, body  MobileAppConfigRequest) (MobileAppConfiguration, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            updateBuildConfigResponse MobileAppConfiguration
	    )

        
            
        
            
        
            
        
            
        

         

        
        
        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return MobileAppConfiguration{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return MobileAppConfiguration{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "put",
            fmt.Sprintf("/service/platform/configuration/v1.0/company/%s/application/%s/build/%s/configuration",co.CompanyID, co.ApplicationID, PlatformType),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return MobileAppConfiguration{}, err
	    }
        
        err = json.Unmarshal(response, &updateBuildConfigResponse)
        if err != nil {
            return MobileAppConfiguration{}, common.NewFDKError(err.Error())
        }
        return updateBuildConfigResponse, nil
        
    }
           
       
    
    
    
  

    
    // GetPreviousVersions Get previous build versions
     func (co *PlatformAppConfiguration)  GetPreviousVersions(PlatformType string) (BuildVersionHistory, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getPreviousVersionsResponse BuildVersionHistory
	    )

        

         

        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            fmt.Sprintf("/service/platform/configuration/v1.0/company/%s/application/%s/build/%s/versions",co.CompanyID, co.ApplicationID, PlatformType),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return BuildVersionHistory{}, err
	    }
        
        err = json.Unmarshal(response, &getPreviousVersionsResponse)
        if err != nil {
            return BuildVersionHistory{}, common.NewFDKError(err.Error())
        }
        return getPreviousVersionsResponse, nil
        
    }
           
       
    
    
    
  

    
    // GetAppFeatures Get features of application
     func (co *PlatformAppConfiguration)  GetAppFeatures() (AppFeatureResponse, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getAppFeaturesResponse AppFeatureResponse
	    )

        

         

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            fmt.Sprintf("/service/platform/configuration/v1.0/company/%s/application/%s/feature",co.CompanyID, co.ApplicationID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return AppFeatureResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getAppFeaturesResponse)
        if err != nil {
            return AppFeatureResponse{}, common.NewFDKError(err.Error())
        }
        return getAppFeaturesResponse, nil
        
    }
           
       
    
    
    
  

    
    // UpdateAppFeatures Update features of application
     func (co *PlatformAppConfiguration)  UpdateAppFeatures(body  AppFeatureRequest) (AppFeature, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            updateAppFeaturesResponse AppFeature
	    )

        
            
        

         

        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return AppFeature{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return AppFeature{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "post",
            fmt.Sprintf("/service/platform/configuration/v1.0/company/%s/application/%s/feature",co.CompanyID, co.ApplicationID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return AppFeature{}, err
	    }
        
        err = json.Unmarshal(response, &updateAppFeaturesResponse)
        if err != nil {
            return AppFeature{}, common.NewFDKError(err.Error())
        }
        return updateAppFeaturesResponse, nil
        
    }
           
       
    
    
    
  

    
    // GetAppBasicDetails Get basic application details
     func (co *PlatformAppConfiguration)  GetAppBasicDetails() (ApplicationDetail, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getAppBasicDetailsResponse ApplicationDetail
	    )

        

         

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            fmt.Sprintf("/service/platform/configuration/v1.0/company/%s/application/%s/detail",co.CompanyID, co.ApplicationID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return ApplicationDetail{}, err
	    }
        
        err = json.Unmarshal(response, &getAppBasicDetailsResponse)
        if err != nil {
            return ApplicationDetail{}, common.NewFDKError(err.Error())
        }
        return getAppBasicDetailsResponse, nil
        
    }
           
       
    
    
    
  

    
    // UpdateAppBasicDetails Add or update application's basic details
     func (co *PlatformAppConfiguration)  UpdateAppBasicDetails(body  ApplicationDetail) (ApplicationDetail, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            updateAppBasicDetailsResponse ApplicationDetail
	    )

        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        

         

        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return ApplicationDetail{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return ApplicationDetail{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "put",
            fmt.Sprintf("/service/platform/configuration/v1.0/company/%s/application/%s/detail",co.CompanyID, co.ApplicationID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return ApplicationDetail{}, err
	    }
        
        err = json.Unmarshal(response, &updateAppBasicDetailsResponse)
        if err != nil {
            return ApplicationDetail{}, common.NewFDKError(err.Error())
        }
        return updateAppBasicDetailsResponse, nil
        
    }
           
       
    
    
    
  

    
    // GetAppContactInfo Get application information
     func (co *PlatformAppConfiguration)  GetAppContactInfo() (ApplicationInformation, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getAppContactInfoResponse ApplicationInformation
	    )

        

         

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            fmt.Sprintf("/service/platform/configuration/v1.0/company/%s/application/%s/information",co.CompanyID, co.ApplicationID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return ApplicationInformation{}, err
	    }
        
        err = json.Unmarshal(response, &getAppContactInfoResponse)
        if err != nil {
            return ApplicationInformation{}, common.NewFDKError(err.Error())
        }
        return getAppContactInfoResponse, nil
        
    }
           
       
    
    
    
  

    
    // UpdateAppContactInfo Get application information
     func (co *PlatformAppConfiguration)  UpdateAppContactInfo(body  ApplicationInformation) (ApplicationInformation, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            updateAppContactInfoResponse ApplicationInformation
	    )

        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        

         

        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return ApplicationInformation{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return ApplicationInformation{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "put",
            fmt.Sprintf("/service/platform/configuration/v1.0/company/%s/application/%s/information",co.CompanyID, co.ApplicationID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return ApplicationInformation{}, err
	    }
        
        err = json.Unmarshal(response, &updateAppContactInfoResponse)
        if err != nil {
            return ApplicationInformation{}, common.NewFDKError(err.Error())
        }
        return updateAppContactInfoResponse, nil
        
    }
           
       
    
    
    
  

    
    // GetAppApiTokens Get social tokens
     func (co *PlatformAppConfiguration)  GetAppApiTokens() (TokenResponse, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getAppApiTokensResponse TokenResponse
	    )

        

         

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            fmt.Sprintf("/service/platform/configuration/v1.0/company/%s/application/%s/token",co.CompanyID, co.ApplicationID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return TokenResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getAppApiTokensResponse)
        if err != nil {
            return TokenResponse{}, common.NewFDKError(err.Error())
        }
        return getAppApiTokensResponse, nil
        
    }
           
       
    
    
    
  

    
    // UpdateAppApiTokens Add social tokens
     func (co *PlatformAppConfiguration)  UpdateAppApiTokens(body  TokenResponse) (TokenResponse, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            updateAppApiTokensResponse TokenResponse
	    )

        
            
        
            
        
            
        
            
        
            
        
            
        

         

        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return TokenResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return TokenResponse{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "post",
            fmt.Sprintf("/service/platform/configuration/v1.0/company/%s/application/%s/token",co.CompanyID, co.ApplicationID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return TokenResponse{}, err
	    }
        
        err = json.Unmarshal(response, &updateAppApiTokensResponse)
        if err != nil {
            return TokenResponse{}, common.NewFDKError(err.Error())
        }
        return updateAppApiTokensResponse, nil
        
    }
           
       
    
    
    
  

    
    //PlatformAppGetAppCompaniesXQuery holds query params
    type PlatformAppGetAppCompaniesXQuery struct { 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"`  
    }
    
    // GetAppCompanies Application inventory enabled companies
     func (co *PlatformAppConfiguration)  GetAppCompanies(xQuery PlatformAppGetAppCompaniesXQuery) (CompaniesResponse, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getAppCompaniesResponse CompaniesResponse
	    )

        

         
            
                
            
                
            
        

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            fmt.Sprintf("/service/platform/configuration/v1.0/company/%s/application/%s/companies",co.CompanyID, co.ApplicationID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return CompaniesResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getAppCompaniesResponse)
        if err != nil {
            return CompaniesResponse{}, common.NewFDKError(err.Error())
        }
        return getAppCompaniesResponse, nil
        
    }
           
            
            
            
            
            
            
            
             
            
            
             
            
            
             
            
            
             
            
            
            
            
            
            
            
            
            
            // GetAppCompaniesPaginator Application inventory enabled companies  
            func (co *PlatformAppConfiguration)  GetAppCompaniesPaginator( xQuery PlatformAppGetAppCompaniesXQuery ) *common.Paginator {
                paginator := common.NewPaginator("number")
                
                
                 xQuery.PageNo  = paginator.PageNo
                 
                 
                 
                 
                 
                 
                 
                 
                paginator.Next = func() (interface{}, error) {
                    response, err := co.GetAppCompanies(xQuery)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, int(response.Page.Current+1), response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }
        
       
    
    
    
  

    
    //PlatformAppGetAppStoresXQuery holds query params
    type PlatformAppGetAppStoresXQuery struct { 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"`  
    }
    
    // GetAppStores Application inventory enabled stores
     func (co *PlatformAppConfiguration)  GetAppStores(xQuery PlatformAppGetAppStoresXQuery) (StoresResponse, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getAppStoresResponse StoresResponse
	    )

        

         
            
                
            
                
            
        

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            fmt.Sprintf("/service/platform/configuration/v1.0/company/%s/application/%s/stores",co.CompanyID, co.ApplicationID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return StoresResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getAppStoresResponse)
        if err != nil {
            return StoresResponse{}, common.NewFDKError(err.Error())
        }
        return getAppStoresResponse, nil
        
    }
           
            
            
            
            
            
            
            
             
            
            
             
            
            
             
            
            
             
            
            
            
            
            
            
            
            
            
            // GetAppStoresPaginator Application inventory enabled stores  
            func (co *PlatformAppConfiguration)  GetAppStoresPaginator( xQuery PlatformAppGetAppStoresXQuery ) *common.Paginator {
                paginator := common.NewPaginator("number")
                
                
                 xQuery.PageNo  = paginator.PageNo
                 
                 
                 
                 
                 
                 
                 
                 
                paginator.Next = func() (interface{}, error) {
                    response, err := co.GetAppStores(xQuery)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, int(response.Page.Current+1), response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }
        
       
    
    
    
  

    
    // GetInventoryConfig Get application configuration
     func (co *PlatformAppConfiguration)  GetInventoryConfig() (ApplicationInventory, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getInventoryConfigResponse ApplicationInventory
	    )

        

         

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            fmt.Sprintf("/service/platform/configuration/v1.0/company/%s/application/%s/configuration",co.CompanyID, co.ApplicationID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return ApplicationInventory{}, err
	    }
        
        err = json.Unmarshal(response, &getInventoryConfigResponse)
        if err != nil {
            return ApplicationInventory{}, common.NewFDKError(err.Error())
        }
        return getInventoryConfigResponse, nil
        
    }
           
       
    
    
    
  

    
    // UpdateInventoryConfig Update application configuration
     func (co *PlatformAppConfiguration)  UpdateInventoryConfig(body  ApplicationInventory) (ApplicationInventory, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            updateInventoryConfigResponse ApplicationInventory
	    )

        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        

         

        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return ApplicationInventory{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return ApplicationInventory{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "put",
            fmt.Sprintf("/service/platform/configuration/v1.0/company/%s/application/%s/configuration",co.CompanyID, co.ApplicationID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return ApplicationInventory{}, err
	    }
        
        err = json.Unmarshal(response, &updateInventoryConfigResponse)
        if err != nil {
            return ApplicationInventory{}, common.NewFDKError(err.Error())
        }
        return updateInventoryConfigResponse, nil
        
    }
           
       
    
    
    
  

    
    // PartiallyUpdateInventoryConfig Partially update application configuration
     func (co *PlatformAppConfiguration)  PartiallyUpdateInventoryConfig(body  AppInventoryPartialUpdate) (ApplicationInventory, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            partiallyUpdateInventoryConfigResponse ApplicationInventory
	    )

        
            
        
            
        
            
        
            
        
            
        

         

        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return ApplicationInventory{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return ApplicationInventory{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "patch",
            fmt.Sprintf("/service/platform/configuration/v1.0/company/%s/application/%s/configuration",co.CompanyID, co.ApplicationID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return ApplicationInventory{}, err
	    }
        
        err = json.Unmarshal(response, &partiallyUpdateInventoryConfigResponse)
        if err != nil {
            return ApplicationInventory{}, common.NewFDKError(err.Error())
        }
        return partiallyUpdateInventoryConfigResponse, nil
        
    }
           
       
    
    
    
  

    
    // GetAppCurrencyConfig Get application enabled currency list
     func (co *PlatformAppConfiguration)  GetAppCurrencyConfig() (AppSupportedCurrency, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getAppCurrencyConfigResponse AppSupportedCurrency
	    )

        

         

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            fmt.Sprintf("/service/platform/configuration/v1.0/company/%s/application/%s/currency",co.CompanyID, co.ApplicationID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return AppSupportedCurrency{}, err
	    }
        
        err = json.Unmarshal(response, &getAppCurrencyConfigResponse)
        if err != nil {
            return AppSupportedCurrency{}, common.NewFDKError(err.Error())
        }
        return getAppCurrencyConfigResponse, nil
        
    }
           
       
    
    
    
  

    
    // UpdateAppCurrencyConfig Add initial application supported currency
     func (co *PlatformAppConfiguration)  UpdateAppCurrencyConfig(body  AppSupportedCurrency) (AppSupportedCurrency, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            updateAppCurrencyConfigResponse AppSupportedCurrency
	    )

        
            
        
            
        
            
        
            
        
            
        
            
        

         

        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return AppSupportedCurrency{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return AppSupportedCurrency{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "post",
            fmt.Sprintf("/service/platform/configuration/v1.0/company/%s/application/%s/currency",co.CompanyID, co.ApplicationID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return AppSupportedCurrency{}, err
	    }
        
        err = json.Unmarshal(response, &updateAppCurrencyConfigResponse)
        if err != nil {
            return AppSupportedCurrency{}, common.NewFDKError(err.Error())
        }
        return updateAppCurrencyConfigResponse, nil
        
    }
           
       
    
    
    
  

    
    //PlatformAppGetOrderingStoresByFilterXQuery holds query params
    type PlatformAppGetOrderingStoresByFilterXQuery struct { 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"`  
    }
    
    // GetOrderingStoresByFilter Get ordering store by filter
     func (co *PlatformAppConfiguration)  GetOrderingStoresByFilter(xQuery PlatformAppGetOrderingStoresByFilterXQuery, body  FilterOrderingStoreRequest) (OrderingStores, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getOrderingStoresByFilterResponse OrderingStores
	    )

        
            
        
            
        
            
        

         
            
                
            
                
            
        

        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return OrderingStores{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return OrderingStores{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "post",
            fmt.Sprintf("/service/platform/configuration/v1.0/company/%s/application/%s/ordering-store/stores/filter",co.CompanyID, co.ApplicationID),
            nil,
            xQuery,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return OrderingStores{}, err
	    }
        
        err = json.Unmarshal(response, &getOrderingStoresByFilterResponse)
        if err != nil {
            return OrderingStores{}, common.NewFDKError(err.Error())
        }
        return getOrderingStoresByFilterResponse, nil
        
    }
           
            
            
            
            
            
            
            
             
            
            
             
            
            
             
            
            
             
            
            
            
            
            
            
            
            
            
            // GetOrderingStoresByFilterPaginator Get ordering store by filter  
            func (co *PlatformAppConfiguration)  GetOrderingStoresByFilterPaginator( xQuery PlatformAppGetOrderingStoresByFilterXQuery  , body  FilterOrderingStoreRequest) *common.Paginator {
                paginator := common.NewPaginator("number")
                
                
                 xQuery.PageNo  = paginator.PageNo
                 
                 
                 
                 
                 
                 
                 
                 
                paginator.Next = func() (interface{}, error) {
                    response, err := co.GetOrderingStoresByFilter(xQuery, body)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, int(response.Page.Current+1), response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }
        
       
    
    
    
  

    
    // UpdateOrderingStoreConfig Add/Update ordering store config
     func (co *PlatformAppConfiguration)  UpdateOrderingStoreConfig(body  OrderingStoreConfig) (DeploymentMeta, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            updateOrderingStoreConfigResponse DeploymentMeta
	    )

        
            
        

         

        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return DeploymentMeta{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return DeploymentMeta{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "post",
            fmt.Sprintf("/service/platform/configuration/v1.0/company/%s/application/%s/ordering-store",co.CompanyID, co.ApplicationID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return DeploymentMeta{}, err
	    }
        
        err = json.Unmarshal(response, &updateOrderingStoreConfigResponse)
        if err != nil {
            return DeploymentMeta{}, common.NewFDKError(err.Error())
        }
        return updateOrderingStoreConfigResponse, nil
        
    }
           
       
    
    
    
  

    
    // GetDomains Get attached domain list
     func (co *PlatformAppConfiguration)  GetDomains() (DomainsResponse, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getDomainsResponse DomainsResponse
	    )

        

         

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            fmt.Sprintf("/service/platform/configuration/v1.0/company/%s/application/%s/domain",co.CompanyID, co.ApplicationID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return DomainsResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getDomainsResponse)
        if err != nil {
            return DomainsResponse{}, common.NewFDKError(err.Error())
        }
        return getDomainsResponse, nil
        
    }
           
       
    
    
    
  

    
    // AddDomain Add new domain to application
     func (co *PlatformAppConfiguration)  AddDomain(body  DomainAddRequest) (Domain, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            addDomainResponse Domain
	    )

        
            
        

         

        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return Domain{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return Domain{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "post",
            fmt.Sprintf("/service/platform/configuration/v1.0/company/%s/application/%s/domain",co.CompanyID, co.ApplicationID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return Domain{}, err
	    }
        
        err = json.Unmarshal(response, &addDomainResponse)
        if err != nil {
            return Domain{}, common.NewFDKError(err.Error())
        }
        return addDomainResponse, nil
        
    }
           
       
    
    
    
  

    
    // RemoveDomainById Remove attached domain
     func (co *PlatformAppConfiguration)  RemoveDomainById(ID string) (SuccessMessageResponse, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            removeDomainByIdResponse SuccessMessageResponse
	    )

        

         

        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "delete",
            fmt.Sprintf("/service/platform/configuration/v1.0/company/%s/application/%s/domain/%s",co.CompanyID, co.ApplicationID, ID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return SuccessMessageResponse{}, err
	    }
        
        err = json.Unmarshal(response, &removeDomainByIdResponse)
        if err != nil {
            return SuccessMessageResponse{}, common.NewFDKError(err.Error())
        }
        return removeDomainByIdResponse, nil
        
    }
           
       
    
    
    
  

    
    // ChangeDomainType Change domain type
     func (co *PlatformAppConfiguration)  ChangeDomainType(body  UpdateDomainTypeRequest) (DomainsResponse, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            changeDomainTypeResponse DomainsResponse
	    )

        
            
        
            
        

         

        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return DomainsResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return DomainsResponse{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "post",
            fmt.Sprintf("/service/platform/configuration/v1.0/company/%s/application/%s/domain/set-domain",co.CompanyID, co.ApplicationID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return DomainsResponse{}, err
	    }
        
        err = json.Unmarshal(response, &changeDomainTypeResponse)
        if err != nil {
            return DomainsResponse{}, common.NewFDKError(err.Error())
        }
        return changeDomainTypeResponse, nil
        
    }
           
       
    
    
    
  

    
    // GetDomainStatus Get domain connected status.
     func (co *PlatformAppConfiguration)  GetDomainStatus(body  DomainStatusRequest) (DomainStatusResponse, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getDomainStatusResponse DomainStatusResponse
	    )

        
            
        

         

        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return DomainStatusResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return DomainStatusResponse{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "post",
            fmt.Sprintf("/service/platform/configuration/v1.0/company/%s/application/%s/domain/domain-status",co.CompanyID, co.ApplicationID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return DomainStatusResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getDomainStatusResponse)
        if err != nil {
            return DomainStatusResponse{}, common.NewFDKError(err.Error())
        }
        return getDomainStatusResponse, nil
        
    }
           
       
    
    
    
    
    
    
    
  

    
    // GetApplicationById Get application data from id
     func (co *PlatformAppConfiguration)  GetApplicationById() (Application, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getApplicationByIdResponse Application
	    )

        

         

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            fmt.Sprintf("/service/platform/configuration/v1.0/company/%s/application/%s",co.CompanyID, co.ApplicationID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return Application{}, err
	    }
        
        err = json.Unmarshal(response, &getApplicationByIdResponse)
        if err != nil {
            return Application{}, common.NewFDKError(err.Error())
        }
        return getApplicationByIdResponse, nil
        
    }
           
       
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

 
	 
   // PlatformAppCart holds PlatformAppCart object properties
    type PlatformAppCart struct {
        config *PlatformConfig
        CompanyID string
        ApplicationID string
    }
    // NewPlatformAppCart returns new PlatformAppCart instance
    func NewPlatformAppCart(config *PlatformConfig, appID string) *PlatformAppCart {
        return &PlatformAppCart{config, config.CompanyID, appID}
    }
    
    
    
  

    
    //PlatformAppGetCouponsXQuery holds query params
    type PlatformAppGetCouponsXQuery struct { 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"` 
        IsArchived bool  `url:"is_archived,omitempty"` 
        Title string  `url:"title,omitempty"` 
        IsPublic bool  `url:"is_public,omitempty"` 
        IsDisplay bool  `url:"is_display,omitempty"` 
        TypeSlug string  `url:"type_slug,omitempty"` 
        Code string  `url:"code,omitempty"`  
    }
    
    // GetCoupons Get with single coupon details or coupon list
     func (ca *PlatformAppCart)  GetCoupons(xQuery PlatformAppGetCouponsXQuery) (CouponsResponse, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getCouponsResponse CouponsResponse
	    )

        

         
            
                
            
                
            
                
            
                
            
                
            
                
            
                
            
                
            
        

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "get",
            fmt.Sprintf("/service/platform/cart/v1.0/company/%s/application/%s/coupon",ca.CompanyID, ca.ApplicationID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return CouponsResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getCouponsResponse)
        if err != nil {
            return CouponsResponse{}, common.NewFDKError(err.Error())
        }
        return getCouponsResponse, nil
        
    }
           
            
            
            
            
            
            
            
             
            
            
             
            
            
             
            
            
             
            
            
             
            
            
             
            
            
             
            
            
             
            
            
             
            
            
             
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            // GetCouponsPaginator Get with single coupon details or coupon list  
            func (ca *PlatformAppCart)  GetCouponsPaginator( xQuery PlatformAppGetCouponsXQuery ) *common.Paginator {
                paginator := common.NewPaginator("number")
                
                
                 xQuery.PageNo  = paginator.PageNo
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                paginator.Next = func() (interface{}, error) {
                    response, err := ca.GetCoupons(xQuery)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, int(response.Page.Current+1), response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }
        
       
    
    
    
  

    
    // CreateCoupon Create new coupon
     func (ca *PlatformAppCart)  CreateCoupon(body  CouponAdd) (interface{}, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
	    )

        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        

         

        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return []byte{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return []byte{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "post",
            fmt.Sprintf("/service/platform/cart/v1.0/company/%s/application/%s/coupon",ca.CompanyID, ca.ApplicationID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return []byte{}, err
	    }
        
        return response, nil
        
    }
           
       
    
    
    
  

    
    // GetCouponById Get with single coupon details or coupon list
     func (ca *PlatformAppCart)  GetCouponById(ID string) (CouponUpdate, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getCouponByIdResponse CouponUpdate
	    )

        

         

        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "get",
            fmt.Sprintf("/service/platform/cart/v1.0/company/%s/application/%s/coupon/%s",ca.CompanyID, ca.ApplicationID, ID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return CouponUpdate{}, err
	    }
        
        err = json.Unmarshal(response, &getCouponByIdResponse)
        if err != nil {
            return CouponUpdate{}, common.NewFDKError(err.Error())
        }
        return getCouponByIdResponse, nil
        
    }
           
       
    
    
    
  

    
    // UpdateCoupon Update existing coupon configuration
     func (ca *PlatformAppCart)  UpdateCoupon(ID string, body  CouponUpdate) (SuccessMessage, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            updateCouponResponse SuccessMessage
	    )

        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        

         

        
        
        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return SuccessMessage{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return SuccessMessage{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "put",
            fmt.Sprintf("/service/platform/cart/v1.0/company/%s/application/%s/coupon/%s",ca.CompanyID, ca.ApplicationID, ID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return SuccessMessage{}, err
	    }
        
        err = json.Unmarshal(response, &updateCouponResponse)
        if err != nil {
            return SuccessMessage{}, common.NewFDKError(err.Error())
        }
        return updateCouponResponse, nil
        
    }
           
       
    
    
    
  

    
    // UpdateCouponPartially Update coupon archive state and schedule
     func (ca *PlatformAppCart)  UpdateCouponPartially(ID string, body  CouponPartialUpdate) (SuccessMessage, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            updateCouponPartiallyResponse SuccessMessage
	    )

        
            
        
            
        

         

        
        
        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return SuccessMessage{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return SuccessMessage{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "patch",
            fmt.Sprintf("/service/platform/cart/v1.0/company/%s/application/%s/coupon/%s",ca.CompanyID, ca.ApplicationID, ID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return SuccessMessage{}, err
	    }
        
        err = json.Unmarshal(response, &updateCouponPartiallyResponse)
        if err != nil {
            return SuccessMessage{}, common.NewFDKError(err.Error())
        }
        return updateCouponPartiallyResponse, nil
        
    }
           
       
    

 
	 
   // PlatformAppRewards holds PlatformAppRewards object properties
    type PlatformAppRewards struct {
        config *PlatformConfig
        CompanyID string
        ApplicationID string
    }
    // NewPlatformAppRewards returns new PlatformAppRewards instance
    func NewPlatformAppRewards(config *PlatformConfig, appID string) *PlatformAppRewards {
        return &PlatformAppRewards{config, config.CompanyID, appID}
    }
    
    
    
  

    
    //PlatformAppGetGiveawaysXQuery holds query params
    type PlatformAppGetGiveawaysXQuery struct { 
        PageID string  `url:"page_id,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"`  
    }
    
    // GetGiveaways List of giveaways of the current application.
     func (re *PlatformAppRewards)  GetGiveaways(xQuery PlatformAppGetGiveawaysXQuery) (GiveawayResponse, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getGiveawaysResponse GiveawayResponse
	    )

        

         
            
                
            
                
            
        

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            re.config,
            "get",
            fmt.Sprintf("/service/platform/rewards/v1.0/company/%s/application/%s/giveaways/",re.CompanyID, re.ApplicationID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return GiveawayResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getGiveawaysResponse)
        if err != nil {
            return GiveawayResponse{}, common.NewFDKError(err.Error())
        }
        return getGiveawaysResponse, nil
        
    }
           
            
            
            
            
            
            
            
             
            
            
             
            
            
             
            
            
             
            
            
            
            
            
            
            
            
            
            // GetGiveawaysPaginator List of giveaways of the current application.  
            func (re *PlatformAppRewards)  GetGiveawaysPaginator( xQuery PlatformAppGetGiveawaysXQuery ) *common.Paginator {
                paginator := common.NewPaginator("cursor")
                
                
                 
                 xQuery.PageID = paginator.NextID
                 
                 
                 
                 
                 
                 
                 
                paginator.Next = func() (interface{}, error) {
                    response, err := re.GetGiveaways(xQuery)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, int(response.Page.Current+1), response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }
        
       
    
    
    
  

    
    // CreateGiveaway Adds a new giveaway.
     func (re *PlatformAppRewards)  CreateGiveaway(body  Giveaway) (Giveaway, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            createGiveawayResponse Giveaway
	    )

        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        

         

        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return Giveaway{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return Giveaway{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            re.config,
            "post",
            fmt.Sprintf("/service/platform/rewards/v1.0/company/%s/application/%s/giveaways/",re.CompanyID, re.ApplicationID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return Giveaway{}, err
	    }
        
        err = json.Unmarshal(response, &createGiveawayResponse)
        if err != nil {
            return Giveaway{}, common.NewFDKError(err.Error())
        }
        return createGiveawayResponse, nil
        
    }
           
       
    
    
    
  

    
    // GetGiveawayByID Get giveaway by ID.
     func (re *PlatformAppRewards)  GetGiveawayByID(ID string) (Giveaway, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getGiveawayByIDResponse Giveaway
	    )

        

         

        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            re.config,
            "get",
            fmt.Sprintf("/service/platform/rewards/v1.0/company/%s/application/%s/giveaways/%s/",re.CompanyID, re.ApplicationID, ID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return Giveaway{}, err
	    }
        
        err = json.Unmarshal(response, &getGiveawayByIDResponse)
        if err != nil {
            return Giveaway{}, common.NewFDKError(err.Error())
        }
        return getGiveawayByIDResponse, nil
        
    }
           
       
    
    
    
  

    
    // UpdateGiveaway Updates the giveaway by it's ID.
     func (re *PlatformAppRewards)  UpdateGiveaway(ID string, body  Giveaway) (Giveaway, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            updateGiveawayResponse Giveaway
	    )

        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        

         

        
        
        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return Giveaway{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return Giveaway{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            re.config,
            "put",
            fmt.Sprintf("/service/platform/rewards/v1.0/company/%s/application/%s/giveaways/%s/",re.CompanyID, re.ApplicationID, ID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return Giveaway{}, err
	    }
        
        err = json.Unmarshal(response, &updateGiveawayResponse)
        if err != nil {
            return Giveaway{}, common.NewFDKError(err.Error())
        }
        return updateGiveawayResponse, nil
        
    }
           
       
    
    
    
  

    
    // GetOffers List of offer of the current application.
     func (re *PlatformAppRewards)  GetOffers() ([]Offer, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getOffersResponse []Offer
	    )

        

         

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            re.config,
            "get",
            fmt.Sprintf("/service/platform/rewards/v1.0/company/%s/application/%s/offers/",re.CompanyID, re.ApplicationID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return []Offer{}, err
	    }
        
        err = json.Unmarshal(response, &getOffersResponse)
        if err != nil {
            return []Offer{}, common.NewFDKError(err.Error())
        }
        return getOffersResponse, nil
        
    }
           
       
    
    
    
  

    
    // GetOfferByName Get offer by name.
     func (re *PlatformAppRewards)  GetOfferByName(Cookie string, Name string) (Offer, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getOfferByNameResponse Offer
	    )

        

         

        
        
        
        
        
        //Adding extra headers
        var xHeaders = make(map[string]string) 
        
         
         xHeaders["cookie"] =  Cookie
         
        
        
        //API call
        rawRequest = NewRequest(
            re.config,
            "get",
            fmt.Sprintf("/service/platform/rewards/v1.0/company/%s/application/%s/offers/%s/",re.CompanyID, re.ApplicationID, Name),
            xHeaders,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return Offer{}, err
	    }
        
        err = json.Unmarshal(response, &getOfferByNameResponse)
        if err != nil {
            return Offer{}, common.NewFDKError(err.Error())
        }
        return getOfferByNameResponse, nil
        
    }
           
       
    
    
    
  

    
    // UpdateOfferByName Updates the offer by name.
     func (re *PlatformAppRewards)  UpdateOfferByName(Name string, body  Offer) (Offer, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            updateOfferByNameResponse Offer
	    )

        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        

         

        
        
        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return Offer{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return Offer{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            re.config,
            "put",
            fmt.Sprintf("/service/platform/rewards/v1.0/company/%s/application/%s/offers/%s/",re.CompanyID, re.ApplicationID, Name),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return Offer{}, err
	    }
        
        err = json.Unmarshal(response, &updateOfferByNameResponse)
        if err != nil {
            return Offer{}, common.NewFDKError(err.Error())
        }
        return updateOfferByNameResponse, nil
        
    }
           
       
    
    
    
  

    
    // GetUserAvailablePoints User's reward details.
     func (re *PlatformAppRewards)  GetUserAvailablePoints(UserID string) (UserRes, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getUserAvailablePointsResponse UserRes
	    )

        

         

        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            re.config,
            "get",
            fmt.Sprintf("/service/platform/rewards/v1.0/company/%s/application/%s/users/%s/",re.CompanyID, re.ApplicationID, UserID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return UserRes{}, err
	    }
        
        err = json.Unmarshal(response, &getUserAvailablePointsResponse)
        if err != nil {
            return UserRes{}, common.NewFDKError(err.Error())
        }
        return getUserAvailablePointsResponse, nil
        
    }
           
       
    
    
    
  

    
    // UpdateUserStatus Update User status
     func (re *PlatformAppRewards)  UpdateUserStatus(UserID string, body  AppUser) (AppUser, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            updateUserStatusResponse AppUser
	    )

        
            
        
            
        
            
        
            
        
            
        
            
        
            
        

         

        
        
        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return AppUser{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return AppUser{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            re.config,
            "patch",
            fmt.Sprintf("/service/platform/rewards/v1.0/company/%s/application/%s/users/%s/",re.CompanyID, re.ApplicationID, UserID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return AppUser{}, err
	    }
        
        err = json.Unmarshal(response, &updateUserStatusResponse)
        if err != nil {
            return AppUser{}, common.NewFDKError(err.Error())
        }
        return updateUserStatusResponse, nil
        
    }
           
       
    
    
    
  

    
    //PlatformAppGetUserPointsHistoryXQuery holds query params
    type PlatformAppGetUserPointsHistoryXQuery struct { 
        PageID string  `url:"page_id,omitempty"` 
        PageLimit float64  `url:"page_limit,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"`  
    }
    
    // GetUserPointsHistory Get list of points transactions.
     func (re *PlatformAppRewards)  GetUserPointsHistory(UserID string, xQuery PlatformAppGetUserPointsHistoryXQuery) (HistoryRes, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getUserPointsHistoryResponse HistoryRes
	    )

        

         
            
                
            
                
            
                
            
        

        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            re.config,
            "get",
            fmt.Sprintf("/service/platform/rewards/v1.0/company/%s/application/%s/users/%s/points/history/",re.CompanyID, re.ApplicationID, UserID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return HistoryRes{}, err
	    }
        
        err = json.Unmarshal(response, &getUserPointsHistoryResponse)
        if err != nil {
            return HistoryRes{}, common.NewFDKError(err.Error())
        }
        return getUserPointsHistoryResponse, nil
        
    }
           
            
            
            
            
            
            
            
             
            
            
             
            
            
             
            
            
            
             
            
            
             
            
            
             
            
            
            
            
            
            
            
            
            
            
            
            
            // GetUserPointsHistoryPaginator Get list of points transactions.  
            func (re *PlatformAppRewards)  GetUserPointsHistoryPaginator(UserID string ,  xQuery PlatformAppGetUserPointsHistoryXQuery ) *common.Paginator {
                paginator := common.NewPaginator("cursor")
                
                
                 
                 xQuery.PageID = paginator.NextID
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                paginator.Next = func() (interface{}, error) {
                    response, err := re.GetUserPointsHistory(UserID, xQuery)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, int(response.Page.Current+1), response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }
        
       
    

 
	 
   // PlatformAppAnalytics holds PlatformAppAnalytics object properties
    type PlatformAppAnalytics struct {
        config *PlatformConfig
        CompanyID string
        ApplicationID string
    }
    // NewPlatformAppAnalytics returns new PlatformAppAnalytics instance
    func NewPlatformAppAnalytics(config *PlatformConfig, appID string) *PlatformAppAnalytics {
        return &PlatformAppAnalytics{config, config.CompanyID, appID}
    }
    
    
    
  

    
    // GetStatiscticsGroups Get statistics groups
     func (an *PlatformAppAnalytics)  GetStatiscticsGroups() (StatsGroups, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getStatiscticsGroupsResponse StatsGroups
	    )

        

         

        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            an.config,
            "get",
            fmt.Sprintf("/service/platform/analytics/v1.0/company/%s/application/%s/stats/group",an.CompanyID, an.ApplicationID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return StatsGroups{}, err
	    }
        
        err = json.Unmarshal(response, &getStatiscticsGroupsResponse)
        if err != nil {
            return StatsGroups{}, common.NewFDKError(err.Error())
        }
        return getStatiscticsGroupsResponse, nil
        
    }
           
       
    
    
    
  

    
    // GetStatiscticsGroupComponents Get statistics group components
     func (an *PlatformAppAnalytics)  GetStatiscticsGroupComponents(GroupName string) (StatsGroupComponents, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getStatiscticsGroupComponentsResponse StatsGroupComponents
	    )

        

         

        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            an.config,
            "get",
            fmt.Sprintf("/service/platform/analytics/v1.0/company/%s/application/%s/stats/group/%s",an.CompanyID, an.ApplicationID, GroupName),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return StatsGroupComponents{}, err
	    }
        
        err = json.Unmarshal(response, &getStatiscticsGroupComponentsResponse)
        if err != nil {
            return StatsGroupComponents{}, common.NewFDKError(err.Error())
        }
        return getStatiscticsGroupComponentsResponse, nil
        
    }
           
       
    
    
    
  

    
    // GetComponentStatsCSV Get component statistics csv
     func (an *PlatformAppAnalytics)  GetComponentStatsCSV(ComponentName string) ([]byte, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
	    )

        

         

        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            an.config,
            "get",
            fmt.Sprintf("/service/platform/analytics/v1.0/company/%s/application/%s/stats/component/%s.csv",an.CompanyID, an.ApplicationID, ComponentName),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return []byte{}, err
	    }
        
        return response, nil
        
    }
           
       
    
    
    
  

    
    // GetComponentStatsPDF Get component statistics pdf
     func (an *PlatformAppAnalytics)  GetComponentStatsPDF(ComponentName string) ([]byte, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
	    )

        

         

        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            an.config,
            "get",
            fmt.Sprintf("/service/platform/analytics/v1.0/company/%s/application/%s/stats/component/%s.pdf",an.CompanyID, an.ApplicationID, ComponentName),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return []byte{}, err
	    }
        
        return response, nil
        
    }
           
       
    
    
    
  

    
    // GetComponentStats Get component statistics
     func (an *PlatformAppAnalytics)  GetComponentStats(ComponentName string) (StatsRes, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getComponentStatsResponse StatsRes
	    )

        

         

        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            an.config,
            "get",
            fmt.Sprintf("/service/platform/analytics/v1.0/company/%s/application/%s/stats/component/%s",an.CompanyID, an.ApplicationID, ComponentName),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return StatsRes{}, err
	    }
        
        err = json.Unmarshal(response, &getComponentStatsResponse)
        if err != nil {
            return StatsRes{}, common.NewFDKError(err.Error())
        }
        return getComponentStatsResponse, nil
        
    }
           
       
    
    
    
  

    
    //PlatformAppGetAbandonCartListXQuery holds query params
    type PlatformAppGetAbandonCartListXQuery struct { 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"`  
    }
    
    // GetAbandonCartList Get abandon carts list
     func (an *PlatformAppAnalytics)  GetAbandonCartList(FromDate string, ToDate string, xQuery PlatformAppGetAbandonCartListXQuery) (AbandonCartsList, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getAbandonCartListResponse AbandonCartsList
	    )

        

         
            
                
            
                
            
        

        
        
        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            an.config,
            "get",
            fmt.Sprintf("/service/platform/analytics/v1.0/company/%s/application/%s/cart/from/%s/to/%s/abandon-cart/",an.CompanyID, an.ApplicationID, FromDate, ToDate),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return AbandonCartsList{}, err
	    }
        
        err = json.Unmarshal(response, &getAbandonCartListResponse)
        if err != nil {
            return AbandonCartsList{}, common.NewFDKError(err.Error())
        }
        return getAbandonCartListResponse, nil
        
    }
           
            
            
            
            
            
            
            
             
            
            
             
            
            
             
            
            
            
             
            
            
            
             
            
            
             
            
            
            
            
            
            
            
            
            
            // GetAbandonCartListPaginator Get abandon carts list  
            func (an *PlatformAppAnalytics)  GetAbandonCartListPaginator(FromDate string , ToDate string ,  xQuery PlatformAppGetAbandonCartListXQuery ) *common.Paginator {
                paginator := common.NewPaginator("number")
                
                
                 xQuery.PageNo  = paginator.PageNo
                 
                 
                 
                 
                 
                 
                 
                 
                paginator.Next = func() (interface{}, error) {
                    response, err := an.GetAbandonCartList(FromDate, ToDate, xQuery)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, int(response.Page.Current+1), response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }
        
       
    
    
    
  

    
    // GetAbandonCartsCSV Get abandon carts csv
     func (an *PlatformAppAnalytics)  GetAbandonCartsCSV(FromDate string, ToDate string) ([]byte, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
	    )

        

         

        
        
        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            an.config,
            "get",
            fmt.Sprintf("/service/platform/analytics/v1.0/company/%s/application/%s/cart/%s/to/%s/abandon-cart.csv",an.CompanyID, an.ApplicationID, FromDate, ToDate),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return []byte{}, err
	    }
        
        return response, nil
        
    }
           
       
    
    
    
  

    
    // GetAbandonCartDetail Get abandon carts details
     func (an *PlatformAppAnalytics)  GetAbandonCartDetail(CartID string) (AbandonCartDetail, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getAbandonCartDetailResponse AbandonCartDetail
	    )

        

         

        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            an.config,
            "get",
            fmt.Sprintf("/service/platform/analytics/v1.0/company/%s/application/%s/cart/abandon-cart/%s",an.CompanyID, an.ApplicationID, CartID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return AbandonCartDetail{}, err
	    }
        
        err = json.Unmarshal(response, &getAbandonCartDetailResponse)
        if err != nil {
            return AbandonCartDetail{}, common.NewFDKError(err.Error())
        }
        return getAbandonCartDetailResponse, nil
        
    }
           
       
    
    
    
    
    
    
    
    
    

 
	 
   // PlatformAppPartner holds PlatformAppPartner object properties
    type PlatformAppPartner struct {
        config *PlatformConfig
        CompanyID string
        ApplicationID string
    }
    // NewPlatformAppPartner returns new PlatformAppPartner instance
    func NewPlatformAppPartner(config *PlatformConfig, appID string) *PlatformAppPartner {
        return &PlatformAppPartner{config, config.CompanyID, appID}
    }
    
    
    
  

    
    // AddProxyPath Add proxy path for external url
     func (pa *PlatformAppPartner)  AddProxyPath(ExtensionID string, body  AddProxyReq) (AddProxyResponse, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            addProxyPathResponse AddProxyResponse
	    )

        
            
        
            
        

         

        
        
        
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
            
             return AddProxyResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
            
             return AddProxyResponse{}, common.NewFDKError(err.Error())       
        }
        
        //API call
        rawRequest = NewRequest(
            pa.config,
            "post",
            fmt.Sprintf("/service/platform/partners/v1.0/company/%s/application/%s/proxy/%s",pa.CompanyID, pa.ApplicationID, ExtensionID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return AddProxyResponse{}, err
	    }
        
        err = json.Unmarshal(response, &addProxyPathResponse)
        if err != nil {
            return AddProxyResponse{}, common.NewFDKError(err.Error())
        }
        return addProxyPathResponse, nil
        
    }
           
       
    
    
    
  

    
    // RemoveProxyPath Remove proxy path for external url
     func (pa *PlatformAppPartner)  RemoveProxyPath(ExtensionID string, AttachedPath string) (interface{}, error) {
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
	    )

        

         

        
        
        
        
        
        
         
        
        
        //API call
        rawRequest = NewRequest(
            pa.config,
            "delete",
            fmt.Sprintf("/service/platform/partners/v1.0/company/%s/application/%s/proxy/%s/%s",pa.CompanyID, pa.ApplicationID, ExtensionID, AttachedPath),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return []byte{}, err
	    }
        
        return response, nil
        
    }
           
       
    

 
