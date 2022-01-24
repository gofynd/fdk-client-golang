package platform



import (
	"encoding/json"
	"fmt"
	"github.com/gofynd/fdk-client-golang/sdk/common"
)
// PlatformClient holds all the PlatformClient object properties  
type PlatformClient struct {
	config            *PlatformConfig
	
		Common  *PlatformCommon
	
		Lead  *PlatformLead
	
		Billing  *PlatformBilling
	
		Communication  *PlatformCommunication
	
		Payment  *PlatformPayment
	
		Order  *PlatformOrder
	
		Catalog  *PlatformCatalog
	
		CompanyProfile  *PlatformCompanyProfile
	
		FileStorage  *PlatformFileStorage
	
		Inventory  *PlatformInventory
	
		Configuration  *PlatformConfiguration
	
		Analytics  *PlatformAnalytics
	
		Discount  *PlatformDiscount
	
		Webhook  *PlatformWebhook
	
	ApplicationClient *ApplicationClient //ApplicationClient embedded
}

// NewPlatformClient returns platform company level service instances
func NewPlatformClient(config *PlatformConfig) *PlatformClient {
		return &PlatformClient{
			config: config,
			
				Common:  NewPlatformCommon(config),
			
				Lead:  NewPlatformLead(config),
			
				Billing:  NewPlatformBilling(config),
			
				Communication:  NewPlatformCommunication(config),
			
				Payment:  NewPlatformPayment(config),
			
				Order:  NewPlatformOrder(config),
			
				Catalog:  NewPlatformCatalog(config),
			
				CompanyProfile:  NewPlatformCompanyProfile(config),
			
				FileStorage:  NewPlatformFileStorage(config),
			
				Inventory:  NewPlatformInventory(config),
			
				Configuration:  NewPlatformConfiguration(config),
			
				Analytics:  NewPlatformAnalytics(config),
			
				Discount:  NewPlatformDiscount(config),
			
				Webhook:  NewPlatformWebhook(config),
			
			ApplicationClient: &ApplicationClient{},
		}
}

//SetPlatformApplicationClient sets platform application level service instances
func (p *PlatformClient) SetPlatformApplicationClient(appID string) {
	p.ApplicationClient = NewApplicationClient(appID, p.config)
}



	
   // PlatformCommon holds PlatformCommon object properties
    type PlatformCommon struct {
        Config *PlatformConfig
        CompanyID string
    }
    // NewPlatformCommon returns new PlatformCommon instance
    func NewPlatformCommon(config *PlatformConfig) *PlatformCommon {
        return &PlatformCommon{Config: config, CompanyID: config.CompanyID}
    }
    
    
   
  
    
    
    //PlatformGetLocationsXQuery holds query params
    type PlatformGetLocationsXQuery struct { 
        LocationType string  `url:"location_type,omitempty"` 
        ID string  `url:"id,omitempty"`  
    }
    


    // GetLocations Get countries, states, cities
     func (co *PlatformCommon)  GetLocations(xQuery PlatformGetLocationsXQuery) (Locations, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getLocationsResponse Locations
	    )

        

        
            
                
            
                
            
        

        

         
        
        
        //API call
        rawRequest = NewRequest(
            co.Config,
            "get",
            "/service/common/configuration/v1.0/location",
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return Locations{}, err
	    }
        
        err = json.Unmarshal(response, &getLocationsResponse)
        if err != nil {
             return Locations{}, common.NewFDKError(err.Error())
        }
        return getLocationsResponse, nil
        
    }
         
        
       
    


	
   // PlatformLead holds PlatformLead object properties
    type PlatformLead struct {
        Config *PlatformConfig
        CompanyID string
    }
    // NewPlatformLead returns new PlatformLead instance
    func NewPlatformLead(config *PlatformConfig) *PlatformLead {
        return &PlatformLead{Config: config, CompanyID: config.CompanyID}
    }
    
    
   
  
    
    
    //PlatformGetTicketsXQuery holds query params
    type PlatformGetTicketsXQuery struct { 
        Items bool  `url:"items,omitempty"` 
        Filters bool  `url:"filters,omitempty"` 
        Q string  `url:"q,omitempty"` 
        Status string  `url:"status,omitempty"` 
        Priority PriorityEnum  `url:"priority,omitempty"` 
        Category string  `url:"category,omitempty"` 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"`  
    }
    


    // GetTickets Gets the list of company level tickets and/or ticket filters depending on query params
     func (le *PlatformLead)  GetTickets(xQuery PlatformGetTicketsXQuery) (TicketList, error){
        
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
            le.Config,
            "get",
            fmt.Sprintf("/service/platform/lead/v1.0/company/%s/ticket",le.CompanyID),
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
         
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
            
            // GetTicketsPaginator Gets the list of company level tickets and/or ticket filters depending on query params  
            func (le *PlatformLead)  GetTicketsPaginator(
              xQuery PlatformGetTicketsXQuery ) *common.Paginator {
                paginator := common.NewPaginator("number")
                
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 xQuery.PageNo  = paginator.PageNo
                 
                 
                 
                 
                 
                 
                 
                 
                paginator.Next = func() (interface{}, error) {
                    response, err := le.GetTickets(xQuery)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, int(response.Page.Current+1), response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }
        
       
    
    
   
  
    
    


    // CreateTicket Creates a company level ticket
     func (le *PlatformLead)  CreateTicket(body  AddTicketPayload) (Ticket, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            createTicketResponse Ticket
	    )

        
        
        
        
        
        
        //enum validation inside request body
        err = body.Priority.IsValid()
        if err != nil {
             
             return Ticket{}, common.NewFDKError(err.Error())
        }
        
        
        
        
        
        

        

        

         
        
        
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
            le.Config,
            "post",
            fmt.Sprintf("/service/platform/lead/v1.0/company/%s/ticket",le.CompanyID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
             return Ticket{}, err
	    }
        
        err = json.Unmarshal(response, &createTicketResponse)
        if err != nil {
             return Ticket{}, common.NewFDKError(err.Error())
        }
        return createTicketResponse, nil
        
    }
         
        
       
    
    
    
    
   
  
    
    


    // GetTicket Retreives ticket details of a company level ticket with ticket ID
     func (le *PlatformLead)  GetTicket(ID string) (Ticket, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getTicketResponse Ticket
	    )

        

        

        
        
        

         
        
        
        //API call
        rawRequest = NewRequest(
            le.Config,
            "get",
            fmt.Sprintf("/service/platform/lead/v1.0/company/%s/ticket/%s",le.CompanyID, ID),
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
         
        
       
    
    
   
  
    
    


    // EditTicket Edits ticket details of a company level ticket
     func (le *PlatformLead)  EditTicket(ID string, body  EditTicketPayload) (Ticket, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            editTicketResponse Ticket
	    )

        
        
        
        
        
        
        
        
        
        
        
        
        //enum validation inside request body
        err = body.Priority.IsValid()
        if err != nil {
             
             return Ticket{}, common.NewFDKError(err.Error())
        }
        
        
        
        
        
        

        

        
        
        

         
        
        
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
            le.Config,
            "put",
            fmt.Sprintf("/service/platform/lead/v1.0/company/%s/ticket/%s",le.CompanyID, ID),
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
         
        
       
    
    
    
    
    
    
   
  
    
    


    // CreateHistory Create history for specific company level ticket
     func (le *PlatformLead)  CreateHistory(ID string, body  TicketHistoryPayload) (TicketHistory, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            createHistoryResponse TicketHistory
	    )

        
        
        
        
        //enum validation inside request body
        err = body.Type.IsValid()
        if err != nil {
             
             return TicketHistory{}, common.NewFDKError(err.Error())
        }
        
        

        

        
        
        

         
        
        
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
            le.Config,
            "post",
            fmt.Sprintf("/service/platform/lead/v1.0/company/%s/ticket/%s/history",le.CompanyID, ID),
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
         
        
       
    
    
   
  
    
    


    // GetTicketHistory Gets history list for specific company level ticket
     func (le *PlatformLead)  GetTicketHistory(ID string) (TicketHistoryList, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getTicketHistoryResponse TicketHistoryList
	    )

        

        

        
        
        

         
        
        
        //API call
        rawRequest = NewRequest(
            le.Config,
            "get",
            fmt.Sprintf("/service/platform/lead/v1.0/company/%s/ticket/%s/history",le.CompanyID, ID),
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
         
        
       
    
    
   
  
    
    


    // GetFeedbacks Gets a list of feedback submitted against that ticket
     func (le *PlatformLead)  GetFeedbacks(ID string) (TicketFeedbackList, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getFeedbacksResponse TicketFeedbackList
	    )

        

        

        
        
        

         
        
        
        //API call
        rawRequest = NewRequest(
            le.Config,
            "get",
            fmt.Sprintf("/service/platform/lead/v1.0/company/%s/ticket/%s/feedback",le.CompanyID, ID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return TicketFeedbackList{}, err
	    }
        
        err = json.Unmarshal(response, &getFeedbacksResponse)
        if err != nil {
             return TicketFeedbackList{}, common.NewFDKError(err.Error())
        }
        return getFeedbacksResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // SubmitFeedback Submit a response for feeback form against that ticket
     func (le *PlatformLead)  SubmitFeedback(ID string, body  TicketFeedbackPayload) (TicketFeedback, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            submitFeedbackResponse TicketFeedback
	    )

        
        
        

        

        
        
        

         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
             
             return TicketFeedback{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
               
             return TicketFeedback{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            le.Config,
            "post",
            fmt.Sprintf("/service/platform/lead/v1.0/company/%s/ticket/%s/feedback",le.CompanyID, ID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
             return TicketFeedback{}, err
	    }
        
        err = json.Unmarshal(response, &submitFeedbackResponse)
        if err != nil {
             return TicketFeedback{}, common.NewFDKError(err.Error())
        }
        return submitFeedbackResponse, nil
        
    }
         
        
       
    
    
    
    
    
    
    
    
    
    
    
    
    
    
   
  
    
    


    // GetTokenForVideoRoom Get Token to join a specific Video Room using it's unqiue name
     func (le *PlatformLead)  GetTokenForVideoRoom(UniqueName string) (GetTokenForVideoRoomResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getTokenForVideoRoomResponse GetTokenForVideoRoomResponse
	    )

        

        

        
        
        

         
        
        
        //API call
        rawRequest = NewRequest(
            le.Config,
            "get",
            fmt.Sprintf("/service/platform/lead/v1.0/company/%s/video/room/%s/token",le.CompanyID, UniqueName),
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
     func (le *PlatformLead)  GetVideoParticipants(UniqueName string) (GetParticipantsInsideVideoRoomResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getVideoParticipantsResponse GetParticipantsInsideVideoRoomResponse
	    )

        

        

        
        
        

         
        
        
        //API call
        rawRequest = NewRequest(
            le.Config,
            "get",
            fmt.Sprintf("/service/platform/lead/v1.0/company/%s/video/room/%s/participants",le.CompanyID, UniqueName),
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
         
        
       
    
    
    
    
    
    
    


	
   // PlatformBilling holds PlatformBilling object properties
    type PlatformBilling struct {
        Config *PlatformConfig
        CompanyID string
    }
    // NewPlatformBilling returns new PlatformBilling instance
    func NewPlatformBilling(config *PlatformConfig) *PlatformBilling {
        return &PlatformBilling{Config: config, CompanyID: config.CompanyID}
    }
    
    
   
  
    
    


    // CreateSubscriptionCharge Create subscription charge
     func (bi *PlatformBilling)  CreateSubscriptionCharge(ExtensionID string, body  CreateSubscriptionCharge) (CreateSubscriptionResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            createSubscriptionChargeResponse CreateSubscriptionResponse
	    )

        
        
        
        
        
        
        
        
        
        
        

        

        
        
        

         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
             
             return CreateSubscriptionResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
               
             return CreateSubscriptionResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            bi.Config,
            "post",
            fmt.Sprintf("/service/platform/billing/v1.0/company/%s/extension/%s/subscription",bi.CompanyID, ExtensionID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
             return CreateSubscriptionResponse{}, err
	    }
        
        err = json.Unmarshal(response, &createSubscriptionChargeResponse)
        if err != nil {
             return CreateSubscriptionResponse{}, common.NewFDKError(err.Error())
        }
        return createSubscriptionChargeResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // GetSubscriptionCharge Get subscription charge details
     func (bi *PlatformBilling)  GetSubscriptionCharge(ExtensionID string, SubscriptionID string) (EntitySubscription, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getSubscriptionChargeResponse EntitySubscription
	    )

        

        

        
        
        
        
        

         
        
        
        //API call
        rawRequest = NewRequest(
            bi.Config,
            "get",
            fmt.Sprintf("/service/platform/billing/v1.0/company/%s/extension/%s/subscription/%s",bi.CompanyID, ExtensionID, SubscriptionID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return EntitySubscription{}, err
	    }
        
        err = json.Unmarshal(response, &getSubscriptionChargeResponse)
        if err != nil {
             return EntitySubscription{}, common.NewFDKError(err.Error())
        }
        return getSubscriptionChargeResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // CancelSubscriptionCharge Cancel subscription charge
     func (bi *PlatformBilling)  CancelSubscriptionCharge(ExtensionID string, SubscriptionID string) (EntitySubscription, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            cancelSubscriptionChargeResponse EntitySubscription
	    )

        

        

        
        
        
        
        

         
        
        
        //API call
        rawRequest = NewRequest(
            bi.Config,
            "post",
            fmt.Sprintf("/service/platform/billing/v1.0/company/%s/extension/%s/subscription/%s/cancel",bi.CompanyID, ExtensionID, SubscriptionID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return EntitySubscription{}, err
	    }
        
        err = json.Unmarshal(response, &cancelSubscriptionChargeResponse)
        if err != nil {
             return EntitySubscription{}, common.NewFDKError(err.Error())
        }
        return cancelSubscriptionChargeResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // GetInvoices Get invoices
     func (bi *PlatformBilling)  GetInvoices() (Invoices, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getInvoicesResponse Invoices
	    )

        

        

        

         
        
        
        //API call
        rawRequest = NewRequest(
            bi.Config,
            "get",
            fmt.Sprintf("/service/platform/billing/v1.0/company/%s/invoice/list",bi.CompanyID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return Invoices{}, err
	    }
        
        err = json.Unmarshal(response, &getInvoicesResponse)
        if err != nil {
             return Invoices{}, common.NewFDKError(err.Error())
        }
        return getInvoicesResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // GetInvoiceById Get invoice by id
     func (bi *PlatformBilling)  GetInvoiceById(InvoiceID string) (Invoice, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getInvoiceByIdResponse Invoice
	    )

        

        

        
        
        

         
        
        
        //API call
        rawRequest = NewRequest(
            bi.Config,
            "get",
            fmt.Sprintf("/service/platform/billing/v1.0/company/%s/invoice/%s",bi.CompanyID, InvoiceID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return Invoice{}, err
	    }
        
        err = json.Unmarshal(response, &getInvoiceByIdResponse)
        if err != nil {
             return Invoice{}, common.NewFDKError(err.Error())
        }
        return getInvoiceByIdResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // GetCustomerDetail Get subscription customer detail
     func (bi *PlatformBilling)  GetCustomerDetail() (SubscriptionCustomer, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getCustomerDetailResponse SubscriptionCustomer
	    )

        

        

        

         
        
        
        //API call
        rawRequest = NewRequest(
            bi.Config,
            "get",
            fmt.Sprintf("/service/platform/billing/v1.0/company/%s/subscription/customer-detail",bi.CompanyID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return SubscriptionCustomer{}, err
	    }
        
        err = json.Unmarshal(response, &getCustomerDetailResponse)
        if err != nil {
             return SubscriptionCustomer{}, common.NewFDKError(err.Error())
        }
        return getCustomerDetailResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // UpsertCustomerDetail Upsert subscription customer detail
     func (bi *PlatformBilling)  UpsertCustomerDetail(body  SubscriptionCustomerCreate) (SubscriptionCustomer, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            upsertCustomerDetailResponse SubscriptionCustomer
	    )

        
        
        
        
        
        
        
        
        
        
        
        
        

        

        

         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
             
             return SubscriptionCustomer{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
               
             return SubscriptionCustomer{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            bi.Config,
            "post",
            fmt.Sprintf("/service/platform/billing/v1.0/company/%s/subscription/customer-detail",bi.CompanyID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
             return SubscriptionCustomer{}, err
	    }
        
        err = json.Unmarshal(response, &upsertCustomerDetailResponse)
        if err != nil {
             return SubscriptionCustomer{}, common.NewFDKError(err.Error())
        }
        return upsertCustomerDetailResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // GetSubscription Get current subscription detail
     func (bi *PlatformBilling)  GetSubscription() (SubscriptionStatus, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getSubscriptionResponse SubscriptionStatus
	    )

        

        

        

         
        
        
        //API call
        rawRequest = NewRequest(
            bi.Config,
            "get",
            fmt.Sprintf("/service/platform/billing/v1.0/company/%s/subscription/current",bi.CompanyID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return SubscriptionStatus{}, err
	    }
        
        err = json.Unmarshal(response, &getSubscriptionResponse)
        if err != nil {
             return SubscriptionStatus{}, common.NewFDKError(err.Error())
        }
        return getSubscriptionResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // GetFeatureLimitConfig Get subscription subscription limits
     func (bi *PlatformBilling)  GetFeatureLimitConfig() (SubscriptionLimit, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getFeatureLimitConfigResponse SubscriptionLimit
	    )

        

        

        

         
        
        
        //API call
        rawRequest = NewRequest(
            bi.Config,
            "get",
            fmt.Sprintf("/service/platform/billing/v1.0/company/%s/subscription/current-limit",bi.CompanyID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return SubscriptionLimit{}, err
	    }
        
        err = json.Unmarshal(response, &getFeatureLimitConfigResponse)
        if err != nil {
             return SubscriptionLimit{}, common.NewFDKError(err.Error())
        }
        return getFeatureLimitConfigResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // ActivateSubscriptionPlan Activate subscription
     func (bi *PlatformBilling)  ActivateSubscriptionPlan(body  SubscriptionActivateReq) (SubscriptionActivateRes, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            activateSubscriptionPlanResponse SubscriptionActivateRes
	    )

        
        
        
        
        
        
        
        
        
        
        

        

        

         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
             
             return SubscriptionActivateRes{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
               
             return SubscriptionActivateRes{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            bi.Config,
            "post",
            fmt.Sprintf("/service/platform/billing/v1.0/company/%s/subscription/activate",bi.CompanyID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
             return SubscriptionActivateRes{}, err
	    }
        
        err = json.Unmarshal(response, &activateSubscriptionPlanResponse)
        if err != nil {
             return SubscriptionActivateRes{}, common.NewFDKError(err.Error())
        }
        return activateSubscriptionPlanResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // CancelSubscriptionPlan Cancel subscription
     func (bi *PlatformBilling)  CancelSubscriptionPlan(body  CancelSubscriptionReq) (CancelSubscriptionRes, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            cancelSubscriptionPlanResponse CancelSubscriptionRes
	    )

        
        
        
        
        
        
        
        
        

        

        

         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
             
             return CancelSubscriptionRes{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
               
             return CancelSubscriptionRes{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            bi.Config,
            "post",
            fmt.Sprintf("/service/platform/billing/v1.0/company/%s/subscription/cancel",bi.CompanyID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
             return CancelSubscriptionRes{}, err
	    }
        
        err = json.Unmarshal(response, &cancelSubscriptionPlanResponse)
        if err != nil {
             return CancelSubscriptionRes{}, common.NewFDKError(err.Error())
        }
        return cancelSubscriptionPlanResponse, nil
        
    }
         
        
       
    


	
   // PlatformCommunication holds PlatformCommunication object properties
    type PlatformCommunication struct {
        Config *PlatformConfig
        CompanyID string
    }
    // NewPlatformCommunication returns new PlatformCommunication instance
    func NewPlatformCommunication(config *PlatformConfig) *PlatformCommunication {
        return &PlatformCommunication{Config: config, CompanyID: config.CompanyID}
    }
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
   
  
    
    
    //PlatformGetSystemNotificationsXQuery holds query params
    type PlatformGetSystemNotificationsXQuery struct { 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"`  
    }
    


    // GetSystemNotifications Get system notifications
     func (co *PlatformCommunication)  GetSystemNotifications(xQuery PlatformGetSystemNotificationsXQuery) (SystemNotifications, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getSystemNotificationsResponse SystemNotifications
	    )

        

        
            
                
            
                
            
        

        

         
        
        
        //API call
        rawRequest = NewRequest(
            co.Config,
            "get",
            fmt.Sprintf("/service/platform/communication/v1.0/company/%s/notification/system-notifications/",co.CompanyID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return SystemNotifications{}, err
	    }
        
        err = json.Unmarshal(response, &getSystemNotificationsResponse)
        if err != nil {
             return SystemNotifications{}, common.NewFDKError(err.Error())
        }
        return getSystemNotificationsResponse, nil
        
    }
         
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
            
            // GetSystemNotificationsPaginator Get system notifications  
            func (co *PlatformCommunication)  GetSystemNotificationsPaginator(
              xQuery PlatformGetSystemNotificationsXQuery ) *common.Paginator {
                paginator := common.NewPaginator("number")
                
                 
                 xQuery.PageNo  = paginator.PageNo
                 
                 
                 
                 
                 
                 
                 
                 
                paginator.Next = func() (interface{}, error) {
                    response, err := co.GetSystemNotifications(xQuery)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, int(response.Page.Current+1), response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }
        
       
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    


	
   // PlatformPayment holds PlatformPayment object properties
    type PlatformPayment struct {
        Config *PlatformConfig
        CompanyID string
    }
    // NewPlatformPayment returns new PlatformPayment instance
    func NewPlatformPayment(config *PlatformConfig) *PlatformPayment {
        return &PlatformPayment{Config: config, CompanyID: config.CompanyID}
    }
    
    
    
    
    
    
    
    
    
    
   
  
    
    
    //PlatformGetAllPayoutsXQuery holds query params
    type PlatformGetAllPayoutsXQuery struct { 
        UniqueExternalID string  `url:"unique_external_id,omitempty"`  
    }
    


    // GetAllPayouts Get All Payouts
     func (pa *PlatformPayment)  GetAllPayouts(xQuery PlatformGetAllPayoutsXQuery) (PayoutsResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getAllPayoutsResponse PayoutsResponse
	    )

        

        
            
                
            
        

        

         
        
        
        //API call
        rawRequest = NewRequest(
            pa.Config,
            "get",
            fmt.Sprintf("/service/platform/payment/v1.0/company/%s/payouts",pa.CompanyID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return PayoutsResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getAllPayoutsResponse)
        if err != nil {
             return PayoutsResponse{}, common.NewFDKError(err.Error())
        }
        return getAllPayoutsResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // SavePayout Save Payout
     func (pa *PlatformPayment)  SavePayout(body  PayoutRequest) (PayoutResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            savePayoutResponse PayoutResponse
	    )

        
        
        
        
        
        
        
        
        
        
        
        
        

        

        

         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
             
             return PayoutResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
               
             return PayoutResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            pa.Config,
            "post",
            fmt.Sprintf("/service/platform/payment/v1.0/company/%s/payouts",pa.CompanyID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
             return PayoutResponse{}, err
	    }
        
        err = json.Unmarshal(response, &savePayoutResponse)
        if err != nil {
             return PayoutResponse{}, common.NewFDKError(err.Error())
        }
        return savePayoutResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // UpdatePayout Update Payout
     func (pa *PlatformPayment)  UpdatePayout(UniqueTransferNo string, body  PayoutRequest) (UpdatePayoutResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            updatePayoutResponse UpdatePayoutResponse
	    )

        
        
        
        
        
        
        
        
        
        
        
        
        

        

        
        
        

         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
             
             return UpdatePayoutResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
               
             return UpdatePayoutResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            pa.Config,
            "put",
            fmt.Sprintf("/service/platform/payment/v1.0/company/%s/payouts/%s",pa.CompanyID, UniqueTransferNo),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
             return UpdatePayoutResponse{}, err
	    }
        
        err = json.Unmarshal(response, &updatePayoutResponse)
        if err != nil {
             return UpdatePayoutResponse{}, common.NewFDKError(err.Error())
        }
        return updatePayoutResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // ActivateAndDectivatePayout Partial Update Payout
     func (pa *PlatformPayment)  ActivateAndDectivatePayout(UniqueTransferNo string, body  UpdatePayoutRequest) (UpdatePayoutResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            activateAndDectivatePayoutResponse UpdatePayoutResponse
	    )

        
        
        
        
        
        
        

        

        
        
        

         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
             
             return UpdatePayoutResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
               
             return UpdatePayoutResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            pa.Config,
            "patch",
            fmt.Sprintf("/service/platform/payment/v1.0/company/%s/payouts/%s",pa.CompanyID, UniqueTransferNo),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
             return UpdatePayoutResponse{}, err
	    }
        
        err = json.Unmarshal(response, &activateAndDectivatePayoutResponse)
        if err != nil {
             return UpdatePayoutResponse{}, common.NewFDKError(err.Error())
        }
        return activateAndDectivatePayoutResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // DeletePayout Delete Payout
     func (pa *PlatformPayment)  DeletePayout(UniqueTransferNo string) (DeletePayoutResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            deletePayoutResponse DeletePayoutResponse
	    )

        

        

        
        
        

         
        
        
        //API call
        rawRequest = NewRequest(
            pa.Config,
            "delete",
            fmt.Sprintf("/service/platform/payment/v1.0/company/%s/payouts/%s",pa.CompanyID, UniqueTransferNo),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return DeletePayoutResponse{}, err
	    }
        
        err = json.Unmarshal(response, &deletePayoutResponse)
        if err != nil {
             return DeletePayoutResponse{}, common.NewFDKError(err.Error())
        }
        return deletePayoutResponse, nil
        
    }
         
        
       
    
    
   
  
    
    
    //PlatformGetSubscriptionPaymentMethodXQuery holds query params
    type PlatformGetSubscriptionPaymentMethodXQuery struct { 
        UniqueExternalID string  `url:"unique_external_id,omitempty"`  
    }
    


    // GetSubscriptionPaymentMethod List Subscription Payment Method
     func (pa *PlatformPayment)  GetSubscriptionPaymentMethod(xQuery PlatformGetSubscriptionPaymentMethodXQuery) (SubscriptionPaymentMethodResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getSubscriptionPaymentMethodResponse SubscriptionPaymentMethodResponse
	    )

        

        
            
                
            
        

        

         
        
        
        //API call
        rawRequest = NewRequest(
            pa.Config,
            "get",
            fmt.Sprintf("/service/platform/payment/v1.0/company/%s/subscription/methods",pa.CompanyID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return SubscriptionPaymentMethodResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getSubscriptionPaymentMethodResponse)
        if err != nil {
             return SubscriptionPaymentMethodResponse{}, common.NewFDKError(err.Error())
        }
        return getSubscriptionPaymentMethodResponse, nil
        
    }
         
        
       
    
    
   
  
    
    
    //PlatformDeleteSubscriptionPaymentMethodXQuery holds query params
    type PlatformDeleteSubscriptionPaymentMethodXQuery struct { 
        UniqueExternalID string  `url:"unique_external_id,omitempty"` 
        PaymentMethodID string  `url:"payment_method_id,omitempty"`  
    }
    


    // DeleteSubscriptionPaymentMethod Delete Subscription Payment Method
     func (pa *PlatformPayment)  DeleteSubscriptionPaymentMethod(xQuery PlatformDeleteSubscriptionPaymentMethodXQuery) (DeleteSubscriptionPaymentMethodResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            deleteSubscriptionPaymentMethodResponse DeleteSubscriptionPaymentMethodResponse
	    )

        

        
            
                
            
                
            
        

        

         
        
        
        //API call
        rawRequest = NewRequest(
            pa.Config,
            "delete",
            fmt.Sprintf("/service/platform/payment/v1.0/company/%s/subscription/methods",pa.CompanyID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return DeleteSubscriptionPaymentMethodResponse{}, err
	    }
        
        err = json.Unmarshal(response, &deleteSubscriptionPaymentMethodResponse)
        if err != nil {
             return DeleteSubscriptionPaymentMethodResponse{}, common.NewFDKError(err.Error())
        }
        return deleteSubscriptionPaymentMethodResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // GetSubscriptionConfig List Subscription Config
     func (pa *PlatformPayment)  GetSubscriptionConfig() (SubscriptionConfigResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getSubscriptionConfigResponse SubscriptionConfigResponse
	    )

        

        

        

         
        
        
        //API call
        rawRequest = NewRequest(
            pa.Config,
            "get",
            fmt.Sprintf("/service/platform/payment/v1.0/company/%s/subscription/configs",pa.CompanyID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return SubscriptionConfigResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getSubscriptionConfigResponse)
        if err != nil {
             return SubscriptionConfigResponse{}, common.NewFDKError(err.Error())
        }
        return getSubscriptionConfigResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // SaveSubscriptionSetupIntent Save Subscription Setup Intent
     func (pa *PlatformPayment)  SaveSubscriptionSetupIntent(body  SaveSubscriptionSetupIntentRequest) (SaveSubscriptionSetupIntentResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            saveSubscriptionSetupIntentResponse SaveSubscriptionSetupIntentResponse
	    )

        
        
        

        

        

         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
             
             return SaveSubscriptionSetupIntentResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
               
             return SaveSubscriptionSetupIntentResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            pa.Config,
            "post",
            fmt.Sprintf("/service/platform/payment/v1.0/company/%s/subscription/setup/intent",pa.CompanyID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
             return SaveSubscriptionSetupIntentResponse{}, err
	    }
        
        err = json.Unmarshal(response, &saveSubscriptionSetupIntentResponse)
        if err != nil {
             return SaveSubscriptionSetupIntentResponse{}, common.NewFDKError(err.Error())
        }
        return saveSubscriptionSetupIntentResponse, nil
        
    }
         
        
       
    
    
    
    
   
  
    
    
    //PlatformVerifyIfscCodeXQuery holds query params
    type PlatformVerifyIfscCodeXQuery struct { 
        IfscCode string  `url:"ifsc_code,omitempty"`  
    }
    


    // VerifyIfscCode Ifsc Code Verification
     func (pa *PlatformPayment)  VerifyIfscCode(xQuery PlatformVerifyIfscCodeXQuery) (IfscCodeResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            verifyIfscCodeResponse IfscCodeResponse
	    )

        

        
            
                
            
        

        

         
        
        
        //API call
        rawRequest = NewRequest(
            pa.Config,
            "get",
            fmt.Sprintf("/service/platform/payment/v1.0/company/%s/ifsc-code/verify",pa.CompanyID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return IfscCodeResponse{}, err
	    }
        
        err = json.Unmarshal(response, &verifyIfscCodeResponse)
        if err != nil {
             return IfscCodeResponse{}, common.NewFDKError(err.Error())
        }
        return verifyIfscCodeResponse, nil
        
    }
         
        
       
    
    
    
    
    
    
    


	
   // PlatformOrder holds PlatformOrder object properties
    type PlatformOrder struct {
        Config *PlatformConfig
        CompanyID string
    }
    // NewPlatformOrder returns new PlatformOrder instance
    func NewPlatformOrder(config *PlatformConfig) *PlatformOrder {
        return &PlatformOrder{Config: config, CompanyID: config.CompanyID}
    }
    
    
   
  
    
    


    // ShipmentStatusUpdate Update status of Shipment
     func (or *PlatformOrder)  ShipmentStatusUpdate(body  UpdateShipmentStatusBody) (UpdateShipmentStatusResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            shipmentStatusUpdateResponse UpdateShipmentStatusResponse
	    )

        
        
        
        
        
        
        

        

        

         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
             
             return UpdateShipmentStatusResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
               
             return UpdateShipmentStatusResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            or.Config,
            "post",
            fmt.Sprintf("/service/platform/order/v1.0/company/%s/actions/status",or.CompanyID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
             return UpdateShipmentStatusResponse{}, err
	    }
        
        err = json.Unmarshal(response, &shipmentStatusUpdateResponse)
        if err != nil {
             return UpdateShipmentStatusResponse{}, common.NewFDKError(err.Error())
        }
        return shipmentStatusUpdateResponse, nil
        
    }
         
        
       
    
    
   
  
    
    
    //PlatformActivityStatusXQuery holds query params
    type PlatformActivityStatusXQuery struct { 
        BagID string  `url:"bag_id,omitempty"`  
    }
    


    // ActivityStatus Get Activity Status
     func (or *PlatformOrder)  ActivityStatus(xQuery PlatformActivityStatusXQuery) (GetActivityStatus, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            activityStatusResponse GetActivityStatus
	    )

        

        
            
                
            
        

        

         
        
        
        //API call
        rawRequest = NewRequest(
            or.Config,
            "get",
            fmt.Sprintf("/service/platform/order/v1.0/company/%s/actions/activity/status",or.CompanyID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return GetActivityStatus{}, err
	    }
        
        err = json.Unmarshal(response, &activityStatusResponse)
        if err != nil {
             return GetActivityStatus{}, common.NewFDKError(err.Error())
        }
        return activityStatusResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // StoreProcessShipmentUpdate Update Store Process-Shipment
     func (or *PlatformOrder)  StoreProcessShipmentUpdate(body  UpdateProcessShipmenstRequestBody) (UpdateProcessShipmenstRequestResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            storeProcessShipmentUpdateResponse UpdateProcessShipmenstRequestResponse
	    )

        
        
        
        
        

        

        

         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
             
             return UpdateProcessShipmenstRequestResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
               
             return UpdateProcessShipmenstRequestResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            or.Config,
            "post",
            fmt.Sprintf("/service/platform/order/v1.0/company/%s/actions/store/process-shipments",or.CompanyID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
             return UpdateProcessShipmenstRequestResponse{}, err
	    }
        
        err = json.Unmarshal(response, &storeProcessShipmentUpdateResponse)
        if err != nil {
             return UpdateProcessShipmenstRequestResponse{}, common.NewFDKError(err.Error())
        }
        return storeProcessShipmentUpdateResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // CheckRefund Check Refund is available or not
     func (or *PlatformOrder)  CheckRefund(ShipmentID string) (map[string]interface{}, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            checkRefundResponse map[string]interface{}
	    )

        

        

        
        
        

         
        
        
        //API call
        rawRequest = NewRequest(
            or.Config,
            "get",
            fmt.Sprintf("/service/platform/order/v1.0/company/%s/actions/check-refund/%s",or.CompanyID, ShipmentID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return map[string]interface{}{}, err
	    }
        
        err = json.Unmarshal(response, &checkRefundResponse)
        if err != nil {
             return map[string]interface{}{}, common.NewFDKError(err.Error())
        }
        return checkRefundResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // ShipmentBagsCanBreak Decides if Shipment bags can break
     func (or *PlatformOrder)  ShipmentBagsCanBreak(body  CanBreakRequestBody) (CanBreakResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            ShipmentBagsCanBreakResponse CanBreakResponse
	    )

        
        
        

        

        

         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
             
             return CanBreakResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
               
             return CanBreakResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            or.Config,
            "post",
            fmt.Sprintf("/service/platform/order/v1.0/company/%s/actions/can-break",or.CompanyID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
             return CanBreakResponse{}, err
	    }
        
        err = json.Unmarshal(response, &ShipmentBagsCanBreakResponse)
        if err != nil {
             return CanBreakResponse{}, common.NewFDKError(err.Error())
        }
        return ShipmentBagsCanBreakResponse, nil
        
    }
         
        
       
    
    
   
  
    
    
    //PlatformGetOrdersByCompanyIdXQuery holds query params
    type PlatformGetOrdersByCompanyIdXQuery struct { 
        PageNo string  `url:"page_no,omitempty"` 
        PageSize string  `url:"page_size,omitempty"` 
        FromDate string  `url:"from_date,omitempty"` 
        ToDate string  `url:"to_date,omitempty"` 
        IsPrioritySort bool  `url:"is_priority_sort,omitempty"` 
        LockStatus bool  `url:"lock_status,omitempty"` 
        Q string  `url:"q,omitempty"` 
        Stage string  `url:"stage,omitempty"` 
        SalesChannels string  `url:"sales_channels,omitempty"` 
        OrderID string  `url:"order_id,omitempty"` 
        Stores string  `url:"stores,omitempty"` 
        DeploymentStores string  `url:"deployment_stores,omitempty"` 
        Status string  `url:"status,omitempty"` 
        Dp string  `url:"dp,omitempty"` 
        ShortenUrls bool  `url:"shorten_urls,omitempty"` 
        FilterType string  `url:"filter_type,omitempty"`  
    }
    


    // GetOrdersByCompanyId Get Orders for company based on Company Id
     func (or *PlatformOrder)  GetOrdersByCompanyId(xQuery PlatformGetOrdersByCompanyIdXQuery) (OrderListing, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getOrdersByCompanyIdResponse OrderListing
	    )

        

        
            
                
            
                
            
                
            
                
            
                
            
                
            
                
            
                
            
                
            
                
            
                
            
                
            
                
            
                
            
                
            
                
            
        

        

         
        
        
        //API call
        rawRequest = NewRequest(
            or.Config,
            "get",
            fmt.Sprintf("/service/platform/order/v1.0/company/%s/orders",or.CompanyID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return OrderListing{}, err
	    }
        
        err = json.Unmarshal(response, &getOrdersByCompanyIdResponse)
        if err != nil {
             return OrderListing{}, common.NewFDKError(err.Error())
        }
        return getOrdersByCompanyIdResponse, nil
        
    }
         
        
       
    
    
   
  
    
    
    //PlatformGetOrderLanesCountByCompanyIdXQuery holds query params
    type PlatformGetOrderLanesCountByCompanyIdXQuery struct { 
        PageNo string  `url:"page_no,omitempty"` 
        PageSize string  `url:"page_size,omitempty"` 
        FromDate string  `url:"from_date,omitempty"` 
        ToDate string  `url:"to_date,omitempty"` 
        Q string  `url:"q,omitempty"` 
        Stage string  `url:"stage,omitempty"` 
        SalesChannels string  `url:"sales_channels,omitempty"` 
        OrderID string  `url:"order_id,omitempty"` 
        Stores string  `url:"stores,omitempty"` 
        Status string  `url:"status,omitempty"` 
        ShortenUrls bool  `url:"shorten_urls,omitempty"` 
        FilterType string  `url:"filter_type,omitempty"`  
    }
    


    // GetOrderLanesCountByCompanyId Get Order Lanes Count for company based on Company Id
     func (or *PlatformOrder)  GetOrderLanesCountByCompanyId(xQuery PlatformGetOrderLanesCountByCompanyIdXQuery) (OrderLanesCount, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getOrderLanesCountByCompanyIdResponse OrderLanesCount
	    )

        

        
            
                
            
                
            
                
            
                
            
                
            
                
            
                
            
                
            
                
            
                
            
                
            
                
            
        

        

         
        
        
        //API call
        rawRequest = NewRequest(
            or.Config,
            "get",
            fmt.Sprintf("/service/platform/order/v1.0/company/%s/orders/lane-count",or.CompanyID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return OrderLanesCount{}, err
	    }
        
        err = json.Unmarshal(response, &getOrderLanesCountByCompanyIdResponse)
        if err != nil {
             return OrderLanesCount{}, common.NewFDKError(err.Error())
        }
        return getOrderLanesCountByCompanyIdResponse, nil
        
    }
         
        
       
    
    
   
  
    
    
    //PlatformGetOrderDetailsXQuery holds query params
    type PlatformGetOrderDetailsXQuery struct { 
        OrderID string  `url:"order_id,omitempty"` 
        Next string  `url:"next,omitempty"` 
        Previous string  `url:"previous,omitempty"`  
    }
    


    // GetOrderDetails Get Order Details for company based on Company Id and Order Id
     func (or *PlatformOrder)  GetOrderDetails(xQuery PlatformGetOrderDetailsXQuery) (OrderDetails, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getOrderDetailsResponse OrderDetails
	    )

        

        
            
                
            
                
            
                
            
        

        

         
        
        
        //API call
        rawRequest = NewRequest(
            or.Config,
            "get",
            fmt.Sprintf("/service/platform/order/v1.0/company/%s/orders/details",or.CompanyID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return OrderDetails{}, err
	    }
        
        err = json.Unmarshal(response, &getOrderDetailsResponse)
        if err != nil {
             return OrderDetails{}, common.NewFDKError(err.Error())
        }
        return getOrderDetailsResponse, nil
        
    }
         
        
       
    
    
    
    
   
  
    
    
    //PlatformGetPicklistOrdersByCompanyIdXQuery holds query params
    type PlatformGetPicklistOrdersByCompanyIdXQuery struct { 
        PageNo string  `url:"page_no,omitempty"` 
        PageSize string  `url:"page_size,omitempty"` 
        FromDate string  `url:"from_date,omitempty"` 
        ToDate string  `url:"to_date,omitempty"` 
        Q string  `url:"q,omitempty"` 
        Stage string  `url:"stage,omitempty"` 
        SalesChannels string  `url:"sales_channels,omitempty"` 
        OrderID string  `url:"order_id,omitempty"` 
        Stores string  `url:"stores,omitempty"` 
        Status string  `url:"status,omitempty"` 
        ShortenUrls bool  `url:"shorten_urls,omitempty"` 
        FilterType string  `url:"filter_type,omitempty"`  
    }
    


    // GetPicklistOrdersByCompanyId Get Orders for company based on Company Id
     func (or *PlatformOrder)  GetPicklistOrdersByCompanyId(xQuery PlatformGetPicklistOrdersByCompanyIdXQuery) (OrderPicklistListing, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getPicklistOrdersByCompanyIdResponse OrderPicklistListing
	    )

        

        
            
                
            
                
            
                
            
                
            
                
            
                
            
                
            
                
            
                
            
                
            
                
            
                
            
        

        

         
        
        
        //API call
        rawRequest = NewRequest(
            or.Config,
            "get",
            fmt.Sprintf("/service/platform/order/v1.0/company/%s/orders/picklist",or.CompanyID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return OrderPicklistListing{}, err
	    }
        
        err = json.Unmarshal(response, &getPicklistOrdersByCompanyIdResponse)
        if err != nil {
             return OrderPicklistListing{}, common.NewFDKError(err.Error())
        }
        return getPicklistOrdersByCompanyIdResponse, nil
        
    }
         
        
       
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
   
  
    
    


    // GetShipmentAddress Use this API to get address of a shipment using its shipment ID and Address Category.
     func (or *PlatformOrder)  GetShipmentAddress(ShipmentID string, AddressCategory string) (GetShipmentAddressResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getShipmentAddressResponse GetShipmentAddressResponse
	    )

        

        

        
        
        
        
        

         
        
        
        //API call
        rawRequest = NewRequest(
            or.Config,
            "get",
            fmt.Sprintf("/service/platform/order/v1.0/company/%s/orders/shipments/%s/address/%s",or.CompanyID, ShipmentID, AddressCategory),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return GetShipmentAddressResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getShipmentAddressResponse)
        if err != nil {
             return GetShipmentAddressResponse{}, common.NewFDKError(err.Error())
        }
        return getShipmentAddressResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // UpdateShipmentAddress Use this API to update address of a shipment using its shipment ID and Address Category.
     func (or *PlatformOrder)  UpdateShipmentAddress(ShipmentID string, AddressCategory string, body  UpdateShipmentAddressRequest) (UpdateShipmentAddressResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            updateShipmentAddressResponse UpdateShipmentAddressResponse
	    )

        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        

        

        
        
        
        
        

         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
             
             return UpdateShipmentAddressResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
               
             return UpdateShipmentAddressResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            or.Config,
            "post",
            fmt.Sprintf("/service/platform/order/v1.0/company/%s/orders/shipments/%s/address/%s",or.CompanyID, ShipmentID, AddressCategory),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
             return UpdateShipmentAddressResponse{}, err
	    }
        
        err = json.Unmarshal(response, &updateShipmentAddressResponse)
        if err != nil {
             return UpdateShipmentAddressResponse{}, common.NewFDKError(err.Error())
        }
        return updateShipmentAddressResponse, nil
        
    }
         
        
       
    
    
    
    
   
  
    
    


    // GetPing Get Ping
     func (or *PlatformOrder)  GetPing() (GetPingResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getPingResponse GetPingResponse
	    )

        

        

        

         
        
        
        //API call
        rawRequest = NewRequest(
            or.Config,
            "get",
            fmt.Sprintf("/service/platform/order/v1.0/company/%s/ping",or.CompanyID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return GetPingResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getPingResponse)
        if err != nil {
             return GetPingResponse{}, common.NewFDKError(err.Error())
        }
        return getPingResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // VoiceCallback Get Voice Callback
     func (or *PlatformOrder)  VoiceCallback() (GetVoiceCallbackResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            voiceCallbackResponse GetVoiceCallbackResponse
	    )

        

        

        

         
        
        
        //API call
        rawRequest = NewRequest(
            or.Config,
            "get",
            fmt.Sprintf("/service/platform/order/v1.0/company/%s/voice/callback",or.CompanyID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return GetVoiceCallbackResponse{}, err
	    }
        
        err = json.Unmarshal(response, &voiceCallbackResponse)
        if err != nil {
             return GetVoiceCallbackResponse{}, common.NewFDKError(err.Error())
        }
        return voiceCallbackResponse, nil
        
    }
         
        
       
    
    
   
  
    
    
    //PlatformVoiceClickToCallXQuery holds query params
    type PlatformVoiceClickToCallXQuery struct { 
        Caller string  `url:"caller,omitempty"` 
        Receiver string  `url:"receiver,omitempty"`  
    }
    


    // VoiceClickToCall Get Voice Click to Call
     func (or *PlatformOrder)  VoiceClickToCall(xQuery PlatformVoiceClickToCallXQuery) (GetClickToCallResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            voiceClickToCallResponse GetClickToCallResponse
	    )

        

        
            
                
            
                
            
        

        

         
        
        
        //API call
        rawRequest = NewRequest(
            or.Config,
            "get",
            fmt.Sprintf("/service/platform/order/v1.0/company/%s/voice/click-to-call",or.CompanyID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return GetClickToCallResponse{}, err
	    }
        
        err = json.Unmarshal(response, &voiceClickToCallResponse)
        if err != nil {
             return GetClickToCallResponse{}, common.NewFDKError(err.Error())
        }
        return voiceClickToCallResponse, nil
        
    }
         
        
       
    


	
   // PlatformCatalog holds PlatformCatalog object properties
    type PlatformCatalog struct {
        Config *PlatformConfig
        CompanyID string
    }
    // NewPlatformCatalog returns new PlatformCatalog instance
    func NewPlatformCatalog(config *PlatformConfig) *PlatformCatalog {
        return &PlatformCatalog{Config: config, CompanyID: config.CompanyID}
    }
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
   
  
    
    
    //PlatformGetProductBundleXQuery holds query params
    type PlatformGetProductBundleXQuery struct { 
        Q string  `url:"q,omitempty"`  
    }
    


    // GetProductBundle List all Product Bundles
     func (ca *PlatformCatalog)  GetProductBundle(xQuery PlatformGetProductBundleXQuery) (GetProductBundleListingResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getProductBundleResponse GetProductBundleListingResponse
	    )

        

        
            
                
            
        

        

         
        
        
        //API call
        rawRequest = NewRequest(
            ca.Config,
            "get",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/product-bundle/",ca.CompanyID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return GetProductBundleListingResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getProductBundleResponse)
        if err != nil {
             return GetProductBundleListingResponse{}, common.NewFDKError(err.Error())
        }
        return getProductBundleResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // CreateProductBundle Create Product Bundle
     func (ca *PlatformCatalog)  CreateProductBundle(body  ProductBundleRequest) (GetProductBundleCreateResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            createProductBundleResponse GetProductBundleCreateResponse
	    )

        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        

        

        

         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
             
             return GetProductBundleCreateResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
               
             return GetProductBundleCreateResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            ca.Config,
            "post",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/product-bundle/",ca.CompanyID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
             return GetProductBundleCreateResponse{}, err
	    }
        
        err = json.Unmarshal(response, &createProductBundleResponse)
        if err != nil {
             return GetProductBundleCreateResponse{}, common.NewFDKError(err.Error())
        }
        return createProductBundleResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // GetProductBundleDetail Get a particular Product Bundle details
     func (ca *PlatformCatalog)  GetProductBundleDetail(ID string) (GetProductBundleResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getProductBundleDetailResponse GetProductBundleResponse
	    )

        

        

        
        
        

         
        
        
        //API call
        rawRequest = NewRequest(
            ca.Config,
            "get",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/product-bundle/%s/",ca.CompanyID, ID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return GetProductBundleResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getProductBundleDetailResponse)
        if err != nil {
             return GetProductBundleResponse{}, common.NewFDKError(err.Error())
        }
        return getProductBundleDetailResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // UpdateProductBundle Update a Product Bundle
     func (ca *PlatformCatalog)  UpdateProductBundle(ID string, body  ProductBundleUpdateRequest) (GetProductBundleCreateResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            updateProductBundleResponse GetProductBundleCreateResponse
	    )

        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        

        

        
        
        

         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
             
             return GetProductBundleCreateResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
               
             return GetProductBundleCreateResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            ca.Config,
            "put",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/product-bundle/%s/",ca.CompanyID, ID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
             return GetProductBundleCreateResponse{}, err
	    }
        
        err = json.Unmarshal(response, &updateProductBundleResponse)
        if err != nil {
             return GetProductBundleCreateResponse{}, common.NewFDKError(err.Error())
        }
        return updateProductBundleResponse, nil
        
    }
         
        
       
    
    
   
  
    
    
    //PlatformGetSizeGuidesXQuery holds query params
    type PlatformGetSizeGuidesXQuery struct { 
        Active bool  `url:"active,omitempty"` 
        Q string  `url:"q,omitempty"` 
        Tag string  `url:"tag,omitempty"` 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"`  
    }
    


    // GetSizeGuides Get list of size guides
     func (ca *PlatformCatalog)  GetSizeGuides(xQuery PlatformGetSizeGuidesXQuery) (ListSizeGuide, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getSizeGuidesResponse ListSizeGuide
	    )

        

        
            
                
            
                
            
                
            
                
            
                
            
        

        

         
        
        
        //API call
        rawRequest = NewRequest(
            ca.Config,
            "get",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/sizeguide",ca.CompanyID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return ListSizeGuide{}, err
	    }
        
        err = json.Unmarshal(response, &getSizeGuidesResponse)
        if err != nil {
             return ListSizeGuide{}, common.NewFDKError(err.Error())
        }
        return getSizeGuidesResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // CreateSizeGuide Create a size guide.
     func (ca *PlatformCatalog)  CreateSizeGuide(body  ValidateSizeGuide) (SuccessResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            createSizeGuideResponse SuccessResponse
	    )

        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        

        

        

         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
             
             return SuccessResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
               
             return SuccessResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            ca.Config,
            "post",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/sizeguide",ca.CompanyID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
             return SuccessResponse{}, err
	    }
        
        err = json.Unmarshal(response, &createSizeGuideResponse)
        if err != nil {
             return SuccessResponse{}, common.NewFDKError(err.Error())
        }
        return createSizeGuideResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // GetSizeGuide Get a single size guide.
     func (ca *PlatformCatalog)  GetSizeGuide(ID string) (SizeGuideResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getSizeGuideResponse SizeGuideResponse
	    )

        

        

        
        
        

         
        
        
        //API call
        rawRequest = NewRequest(
            ca.Config,
            "get",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/sizeguide/%s/",ca.CompanyID, ID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return SizeGuideResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getSizeGuideResponse)
        if err != nil {
             return SizeGuideResponse{}, common.NewFDKError(err.Error())
        }
        return getSizeGuideResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // UpdateSizeGuide Edit a size guide.
     func (ca *PlatformCatalog)  UpdateSizeGuide(ID string, body  ValidateSizeGuide) (SuccessResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            updateSizeGuideResponse SuccessResponse
	    )

        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        

        

        
        
        

         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
             
             return SuccessResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
               
             return SuccessResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            ca.Config,
            "put",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/sizeguide/%s/",ca.CompanyID, ID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
             return SuccessResponse{}, err
	    }
        
        err = json.Unmarshal(response, &updateSizeGuideResponse)
        if err != nil {
             return SuccessResponse{}, common.NewFDKError(err.Error())
        }
        return updateSizeGuideResponse, nil
        
    }
         
        
       
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
   
  
    
    


    // GetSellerInsights Analytics data of catalog and inventory that are being cross-selled.
     func (ca *PlatformCatalog)  GetSellerInsights(SellerAppID string) (CrossSellingResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getSellerInsightsResponse CrossSellingResponse
	    )

        

        

        
        
        

         
        
        
        //API call
        rawRequest = NewRequest(
            ca.Config,
            "get",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/cross-selling/%s/analytics/insights/",ca.CompanyID, SellerAppID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return CrossSellingResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getSellerInsightsResponse)
        if err != nil {
             return CrossSellingResponse{}, common.NewFDKError(err.Error())
        }
        return getSellerInsightsResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // CreateMarketplaceOptin Create/Update opt-in infomation.
     func (ca *PlatformCatalog)  CreateMarketplaceOptin(Marketplace string, body  OptInPostRequest) (UpdatedResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            createMarketplaceOptinResponse UpdatedResponse
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
            ca.Config,
            "post",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/marketplaces/%s/optin/",ca.CompanyID, Marketplace),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
             return UpdatedResponse{}, err
	    }
        
        err = json.Unmarshal(response, &createMarketplaceOptinResponse)
        if err != nil {
             return UpdatedResponse{}, common.NewFDKError(err.Error())
        }
        return createMarketplaceOptinResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // GetMarketplaceOptinDetail Get opt-in infomation.
     func (ca *PlatformCatalog)  GetMarketplaceOptinDetail() (GetOptInPlatform, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getMarketplaceOptinDetailResponse GetOptInPlatform
	    )

        

        

        

         
        
        
        //API call
        rawRequest = NewRequest(
            ca.Config,
            "get",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/marketplaces/",ca.CompanyID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return GetOptInPlatform{}, err
	    }
        
        err = json.Unmarshal(response, &getMarketplaceOptinDetailResponse)
        if err != nil {
             return GetOptInPlatform{}, common.NewFDKError(err.Error())
        }
        return getMarketplaceOptinDetailResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // GetCompanyDetail Get the Company details.
     func (ca *PlatformCatalog)  GetCompanyDetail() (OptinCompanyDetail, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getCompanyDetailResponse OptinCompanyDetail
	    )

        

        

        

         
        
        
        //API call
        rawRequest = NewRequest(
            ca.Config,
            "get",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/marketplaces/company-details/",ca.CompanyID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return OptinCompanyDetail{}, err
	    }
        
        err = json.Unmarshal(response, &getCompanyDetailResponse)
        if err != nil {
             return OptinCompanyDetail{}, common.NewFDKError(err.Error())
        }
        return getCompanyDetailResponse, nil
        
    }
         
        
       
    
    
   
  
    
    
    //PlatformGetCompanyBrandDetailXQuery holds query params
    type PlatformGetCompanyBrandDetailXQuery struct { 
        IsActive bool  `url:"is_active,omitempty"` 
        Q bool  `url:"q,omitempty"` 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"` 
        Marketplace string  `url:"marketplace,omitempty"`  
    }
    


    // GetCompanyBrandDetail Get the Company Brand details of Optin.
     func (ca *PlatformCatalog)  GetCompanyBrandDetail(xQuery PlatformGetCompanyBrandDetailXQuery) (OptinCompanyBrandDetailsView, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getCompanyBrandDetailResponse OptinCompanyBrandDetailsView
	    )

        

        
            
                
            
                
            
                
            
                
            
                
            
        

        

         
        
        
        //API call
        rawRequest = NewRequest(
            ca.Config,
            "get",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/marketplaces/company-brand-details/",ca.CompanyID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return OptinCompanyBrandDetailsView{}, err
	    }
        
        err = json.Unmarshal(response, &getCompanyBrandDetailResponse)
        if err != nil {
             return OptinCompanyBrandDetailsView{}, common.NewFDKError(err.Error())
        }
        return getCompanyBrandDetailResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // GetCompanyMetrics Get the Company metrics
     func (ca *PlatformCatalog)  GetCompanyMetrics() (OptinCompanyMetrics, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getCompanyMetricsResponse OptinCompanyMetrics
	    )

        

        

        

         
        
        
        //API call
        rawRequest = NewRequest(
            ca.Config,
            "get",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/marketplaces/company-metrics/",ca.CompanyID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return OptinCompanyMetrics{}, err
	    }
        
        err = json.Unmarshal(response, &getCompanyMetricsResponse)
        if err != nil {
             return OptinCompanyMetrics{}, common.NewFDKError(err.Error())
        }
        return getCompanyMetricsResponse, nil
        
    }
         
        
       
    
    
   
  
    
    
    //PlatformGetStoreDetailXQuery holds query params
    type PlatformGetStoreDetailXQuery struct { 
        Q string  `url:"q,omitempty"` 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"`  
    }
    


    // GetStoreDetail Get the Store details.
     func (ca *PlatformCatalog)  GetStoreDetail(xQuery PlatformGetStoreDetailXQuery) (OptinStoreDetails, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getStoreDetailResponse OptinStoreDetails
	    )

        

        
            
                
            
                
            
                
            
        

        

         
        
        
        //API call
        rawRequest = NewRequest(
            ca.Config,
            "get",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/marketplaces/location-details/",ca.CompanyID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return OptinStoreDetails{}, err
	    }
        
        err = json.Unmarshal(response, &getStoreDetailResponse)
        if err != nil {
             return OptinStoreDetails{}, common.NewFDKError(err.Error())
        }
        return getStoreDetailResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // GetGenderAttribute Get gender attribute details
     func (ca *PlatformCatalog)  GetGenderAttribute(AttributeSlug string) (GenderDetail, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getGenderAttributeResponse GenderDetail
	    )

        

        

        
        
        

         
        
        
        //API call
        rawRequest = NewRequest(
            ca.Config,
            "get",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/product-attributes/%s",ca.CompanyID, AttributeSlug),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return GenderDetail{}, err
	    }
        
        err = json.Unmarshal(response, &getGenderAttributeResponse)
        if err != nil {
             return GenderDetail{}, common.NewFDKError(err.Error())
        }
        return getGenderAttributeResponse, nil
        
    }
         
        
       
    
    
   
  
    
    
    //PlatformListProductTemplateCategoriesXQuery holds query params
    type PlatformListProductTemplateCategoriesXQuery struct { 
        Departments string  `url:"departments,omitempty"` 
        ItemType string  `url:"item_type,omitempty"`  
    }
    


    // ListProductTemplateCategories List Department specifiec product categories
     func (ca *PlatformCatalog)  ListProductTemplateCategories(xQuery PlatformListProductTemplateCategoriesXQuery) (ProdcutTemplateCategoriesResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            listProductTemplateCategoriesResponse ProdcutTemplateCategoriesResponse
	    )

        

        
            
                
            
                
            
        

        

         
        
        
        //API call
        rawRequest = NewRequest(
            ca.Config,
            "get",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/products/templates/categories/",ca.CompanyID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return ProdcutTemplateCategoriesResponse{}, err
	    }
        
        err = json.Unmarshal(response, &listProductTemplateCategoriesResponse)
        if err != nil {
             return ProdcutTemplateCategoriesResponse{}, common.NewFDKError(err.Error())
        }
        return listProductTemplateCategoriesResponse, nil
        
    }
         
        
       
    
    
   
  
    
    
    //PlatformListDepartmentsDataXQuery holds query params
    type PlatformListDepartmentsDataXQuery struct { 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"` 
        Name string  `url:"name,omitempty"` 
        Search string  `url:"search,omitempty"` 
        IsActive bool  `url:"is_active,omitempty"`  
    }
    


    // ListDepartmentsData List all Departments
     func (ca *PlatformCatalog)  ListDepartmentsData(xQuery PlatformListDepartmentsDataXQuery) (DepartmentsResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            listDepartmentsDataResponse DepartmentsResponse
	    )

        

        
            
                
            
                
            
                
            
                
            
                
            
        

        

         
        
        
        //API call
        rawRequest = NewRequest(
            ca.Config,
            "get",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/departments/",ca.CompanyID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return DepartmentsResponse{}, err
	    }
        
        err = json.Unmarshal(response, &listDepartmentsDataResponse)
        if err != nil {
             return DepartmentsResponse{}, common.NewFDKError(err.Error())
        }
        return listDepartmentsDataResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // GetDepartmentData Get specific departments details by passing in unique id of the department
     func (ca *PlatformCatalog)  GetDepartmentData(UID string) (DepartmentsResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getDepartmentDataResponse DepartmentsResponse
	    )

        

        

        
        
        

         
        
        
        //API call
        rawRequest = NewRequest(
            ca.Config,
            "get",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/departments/%s/",ca.CompanyID, UID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return DepartmentsResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getDepartmentDataResponse)
        if err != nil {
             return DepartmentsResponse{}, common.NewFDKError(err.Error())
        }
        return getDepartmentDataResponse, nil
        
    }
         
        
       
    
    
   
  
    
    
    //PlatformListProductTemplateXQuery holds query params
    type PlatformListProductTemplateXQuery struct { 
        Department string  `url:"department,omitempty"`  
    }
    


    // ListProductTemplate List all Templates
     func (ca *PlatformCatalog)  ListProductTemplate(xQuery PlatformListProductTemplateXQuery) (TemplatesResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            listProductTemplateResponse TemplatesResponse
	    )

        

        
            
                
            
        

        

         
        
        
        //API call
        rawRequest = NewRequest(
            ca.Config,
            "get",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/products/templates/",ca.CompanyID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return TemplatesResponse{}, err
	    }
        
        err = json.Unmarshal(response, &listProductTemplateResponse)
        if err != nil {
             return TemplatesResponse{}, common.NewFDKError(err.Error())
        }
        return listProductTemplateResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // ValidateProductTemplate Validate Product Template Schema
     func (ca *PlatformCatalog)  ValidateProductTemplate(Slug string) (TemplatesValidationResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            validateProductTemplateResponse TemplatesValidationResponse
	    )

        

        

        
        
        

         
        
        
        //API call
        rawRequest = NewRequest(
            ca.Config,
            "get",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/products/templates/%s/validation/schema/",ca.CompanyID, Slug),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return TemplatesValidationResponse{}, err
	    }
        
        err = json.Unmarshal(response, &validateProductTemplateResponse)
        if err != nil {
             return TemplatesValidationResponse{}, common.NewFDKError(err.Error())
        }
        return validateProductTemplateResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // DownloadProductTemplateViews Download Product Template View
     func (ca *PlatformCatalog)  DownloadProductTemplateViews(Slug string) ([]byte, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
	    )

        

        

        
        
        

         
        
        
        //API call
        rawRequest = NewRequest(
            ca.Config,
            "get",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/products/templates/%s/download/",ca.CompanyID, Slug),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return []byte{}, err
	    }
        
        return response, nil
        
    }
         
        
       
    
    
   
  
    
    
    //PlatformDownloadProductTemplateViewXQuery holds query params
    type PlatformDownloadProductTemplateViewXQuery struct { 
        ItemType string  `url:"item_type,omitempty"`  
    }
    


    // DownloadProductTemplateView Download Product Template View
     func (ca *PlatformCatalog)  DownloadProductTemplateView(xQuery PlatformDownloadProductTemplateViewXQuery) ([]byte, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
	    )

        

        
            
                
            
        

        

         
        
        
        //API call
        rawRequest = NewRequest(
            ca.Config,
            "get",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/inventory/templates/download/",ca.CompanyID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return []byte{}, err
	    }
        
        return response, nil
        
    }
         
        
       
    
    
   
  
    
    
    //PlatformValidateProductTemplateSchemaXQuery holds query params
    type PlatformValidateProductTemplateSchemaXQuery struct { 
        ItemType string  `url:"item_type,omitempty"`  
    }
    


    // ValidateProductTemplateSchema Validate Product Template Schema
     func (ca *PlatformCatalog)  ValidateProductTemplateSchema(xQuery PlatformValidateProductTemplateSchemaXQuery) (InventoryValidationResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            validateProductTemplateSchemaResponse InventoryValidationResponse
	    )

        

        
            
                
            
        

        

         
        
        
        //API call
        rawRequest = NewRequest(
            ca.Config,
            "get",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/inventory/templates/validation/schema/",ca.CompanyID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return InventoryValidationResponse{}, err
	    }
        
        err = json.Unmarshal(response, &validateProductTemplateSchemaResponse)
        if err != nil {
             return InventoryValidationResponse{}, common.NewFDKError(err.Error())
        }
        return validateProductTemplateSchemaResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // ListHSNCodes List HSN Codes
     func (ca *PlatformCatalog)  ListHSNCodes() (HSNCodesResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            listHSNCodesResponse HSNCodesResponse
	    )

        

        

        

         
        
        
        //API call
        rawRequest = NewRequest(
            ca.Config,
            "get",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/products/hsn/",ca.CompanyID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return HSNCodesResponse{}, err
	    }
        
        err = json.Unmarshal(response, &listHSNCodesResponse)
        if err != nil {
             return HSNCodesResponse{}, common.NewFDKError(err.Error())
        }
        return listHSNCodesResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // ListProductTemplateExportDetails Allows you to list all product templates export list details
     func (ca *PlatformCatalog)  ListProductTemplateExportDetails() (ProductDownloadsResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            listProductTemplateExportDetailsResponse ProductDownloadsResponse
	    )

        

        

        

         
        
        
        //API call
        rawRequest = NewRequest(
            ca.Config,
            "get",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/products/downloads/",ca.CompanyID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return ProductDownloadsResponse{}, err
	    }
        
        err = json.Unmarshal(response, &listProductTemplateExportDetailsResponse)
        if err != nil {
             return ProductDownloadsResponse{}, common.NewFDKError(err.Error())
        }
        return listProductTemplateExportDetailsResponse, nil
        
    }
         
        
       
    
    
   
  
    
    
    //PlatformListTemplateBrandTypeValuesXQuery holds query params
    type PlatformListTemplateBrandTypeValuesXQuery struct { 
        Filter string  `url:"filter,omitempty"`  
    }
    


    // ListTemplateBrandTypeValues Allows you to list all values for Templates, Brands or Type
     func (ca *PlatformCatalog)  ListTemplateBrandTypeValues(xQuery PlatformListTemplateBrandTypeValuesXQuery) (ProductConfigurationDownloads, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            listTemplateBrandTypeValuesResponse ProductConfigurationDownloads
	    )

        

        
            
                
            
        

        

         
        
        
        //API call
        rawRequest = NewRequest(
            ca.Config,
            "get",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/downloads/configuration/",ca.CompanyID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return ProductConfigurationDownloads{}, err
	    }
        
        err = json.Unmarshal(response, &listTemplateBrandTypeValuesResponse)
        if err != nil {
             return ProductConfigurationDownloads{}, common.NewFDKError(err.Error())
        }
        return listTemplateBrandTypeValuesResponse, nil
        
    }
         
        
       
    
    
   
  
    
    
    //PlatformListCategoriesXQuery holds query params
    type PlatformListCategoriesXQuery struct { 
        Level string  `url:"level,omitempty"` 
        Departments string  `url:"departments,omitempty"` 
        Q string  `url:"q,omitempty"` 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"`  
    }
    


    // ListCategories Get product categories list
     func (ca *PlatformCatalog)  ListCategories(xQuery PlatformListCategoriesXQuery) (CategoryResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            listCategoriesResponse CategoryResponse
	    )

        

        
            
                
            
                
            
                
            
                
            
                
            
        

        

         
        
        
        //API call
        rawRequest = NewRequest(
            ca.Config,
            "get",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/category/",ca.CompanyID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return CategoryResponse{}, err
	    }
        
        err = json.Unmarshal(response, &listCategoriesResponse)
        if err != nil {
             return CategoryResponse{}, common.NewFDKError(err.Error())
        }
        return listCategoriesResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // CreateCategories Create product categories
     func (ca *PlatformCatalog)  CreateCategories(body  CategoryRequestBody) (CategoryCreateResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            createCategoriesResponse CategoryCreateResponse
	    )

        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        

        

        

         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
             
             return CategoryCreateResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
               
             return CategoryCreateResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            ca.Config,
            "post",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/category/",ca.CompanyID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
             return CategoryCreateResponse{}, err
	    }
        
        err = json.Unmarshal(response, &createCategoriesResponse)
        if err != nil {
             return CategoryCreateResponse{}, common.NewFDKError(err.Error())
        }
        return createCategoriesResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // GetCategoryData Get product category by uid
     func (ca *PlatformCatalog)  GetCategoryData(UID string) (SingleCategoryResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getCategoryDataResponse SingleCategoryResponse
	    )

        

        

        
        
        

         
        
        
        //API call
        rawRequest = NewRequest(
            ca.Config,
            "get",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/category/%s/",ca.CompanyID, UID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return SingleCategoryResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getCategoryDataResponse)
        if err != nil {
             return SingleCategoryResponse{}, common.NewFDKError(err.Error())
        }
        return getCategoryDataResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // UpdateCategory Update product categories
     func (ca *PlatformCatalog)  UpdateCategory(UID string, body  CategoryRequestBody) (CategoryUpdateResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            updateCategoryResponse CategoryUpdateResponse
	    )

        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        

        

        
        
        

         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
             
             return CategoryUpdateResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
               
             return CategoryUpdateResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            ca.Config,
            "put",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/category/%s/",ca.CompanyID, UID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
             return CategoryUpdateResponse{}, err
	    }
        
        err = json.Unmarshal(response, &updateCategoryResponse)
        if err != nil {
             return CategoryUpdateResponse{}, common.NewFDKError(err.Error())
        }
        return updateCategoryResponse, nil
        
    }
         
        
       
    
    
   
  
    
    
    //PlatformGetProductsXQuery holds query params
    type PlatformGetProductsXQuery struct { 
        BrandIds []float64  `url:"brand_ids,omitempty"` 
        CategoryIds []float64  `url:"category_ids,omitempty"` 
        ItemIds []float64  `url:"item_ids,omitempty"` 
        DepartmentIds []float64  `url:"department_ids,omitempty"` 
        ItemCode []float64  `url:"item_code,omitempty"` 
        Q string  `url:"q,omitempty"` 
        Tags []string  `url:"tags,omitempty"` 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"`  
    }
    


    // GetProducts Get product list
     func (ca *PlatformCatalog)  GetProducts(xQuery PlatformGetProductsXQuery) (ProductListingResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getProductsResponse ProductListingResponse
	    )

        

        
            
                
            
                
            
                
            
                
            
                
            
                
            
                
            
                
            
                
            
        

        

         
        
        
        //API call
        rawRequest = NewRequest(
            ca.Config,
            "get",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/products/",ca.CompanyID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return ProductListingResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getProductsResponse)
        if err != nil {
             return ProductListingResponse{}, common.NewFDKError(err.Error())
        }
        return getProductsResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // CreateProduct Create a product.
     func (ca *PlatformCatalog)  CreateProduct(body  ProductCreateUpdate) (SuccessResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            createProductResponse SuccessResponse
	    )

        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        

        

        

         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
             
             return SuccessResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
               
             return SuccessResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            ca.Config,
            "post",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/products/",ca.CompanyID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
             return SuccessResponse{}, err
	    }
        
        err = json.Unmarshal(response, &createProductResponse)
        if err != nil {
             return SuccessResponse{}, common.NewFDKError(err.Error())
        }
        return createProductResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // DeleteProduct Delete a product.
     func (ca *PlatformCatalog)  DeleteProduct(ItemID float64) (SuccessResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            deleteProductResponse SuccessResponse
	    )

        

        

        
        
        

         
        
        
        //API call
        rawRequest = NewRequest(
            ca.Config,
            "delete",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/products/undefined/",ca.CompanyID, ItemID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return SuccessResponse{}, err
	    }
        
        err = json.Unmarshal(response, &deleteProductResponse)
        if err != nil {
             return SuccessResponse{}, common.NewFDKError(err.Error())
        }
        return deleteProductResponse, nil
        
    }
         
        
       
    
    
   
  
    
    
    //PlatformGetProductXQuery holds query params
    type PlatformGetProductXQuery struct { 
        ItemCode string  `url:"item_code,omitempty"` 
        BrandUID float64  `url:"brand_uid,omitempty"`  
    }
    


    // GetProduct Get a single product.
     func (ca *PlatformCatalog)  GetProduct(ItemID float64, xQuery PlatformGetProductXQuery) (Product, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getProductResponse Product
	    )

        

        
            
                
            
                
            
        

        
        
        

         
        
        
        //API call
        rawRequest = NewRequest(
            ca.Config,
            "get",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/products/undefined/",ca.CompanyID, ItemID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return Product{}, err
	    }
        
        err = json.Unmarshal(response, &getProductResponse)
        if err != nil {
             return Product{}, common.NewFDKError(err.Error())
        }
        return getProductResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // EditProduct Edit a product.
     func (ca *PlatformCatalog)  EditProduct(ItemID float64, body  ProductCreateUpdate) (SuccessResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            editProductResponse SuccessResponse
	    )

        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        

        

        
        
        

         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
             
             return SuccessResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
               
             return SuccessResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            ca.Config,
            "put",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/products/undefined/",ca.CompanyID, ItemID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
             return SuccessResponse{}, err
	    }
        
        err = json.Unmarshal(response, &editProductResponse)
        if err != nil {
             return SuccessResponse{}, common.NewFDKError(err.Error())
        }
        return editProductResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // GetProductValidation Validate product/size data
     func (ca *PlatformCatalog)  GetProductValidation() (ValidateProduct, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getProductValidationResponse ValidateProduct
	    )

        

        

        

         
        
        
        //API call
        rawRequest = NewRequest(
            ca.Config,
            "get",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/products/validation/",ca.CompanyID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return ValidateProduct{}, err
	    }
        
        err = json.Unmarshal(response, &getProductValidationResponse)
        if err != nil {
             return ValidateProduct{}, common.NewFDKError(err.Error())
        }
        return getProductValidationResponse, nil
        
    }
         
        
       
    
    
   
  
    
    
    //PlatformGetProductSizeXQuery holds query params
    type PlatformGetProductSizeXQuery struct { 
        ItemCode string  `url:"item_code,omitempty"` 
        BrandUID float64  `url:"brand_uid,omitempty"` 
        UID float64  `url:"uid,omitempty"`  
    }
    


    // GetProductSize Get a single product size.
     func (ca *PlatformCatalog)  GetProductSize(ItemID float64, xQuery PlatformGetProductSizeXQuery) (ProductListingResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getProductSizeResponse ProductListingResponse
	    )

        

        
            
                
            
                
            
                
            
        

        
        
        

         
        
        
        //API call
        rawRequest = NewRequest(
            ca.Config,
            "get",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/products/undefined/sizes/",ca.CompanyID, ItemID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return ProductListingResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getProductSizeResponse)
        if err != nil {
             return ProductListingResponse{}, common.NewFDKError(err.Error())
        }
        return getProductSizeResponse, nil
        
    }
         
        
       
    
    
   
  
    
    
    //PlatformGetProductBulkUploadHistoryXQuery holds query params
    type PlatformGetProductBulkUploadHistoryXQuery struct { 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"`  
    }
    


    // GetProductBulkUploadHistory Get a list of all bulk product upload jobs.
     func (ca *PlatformCatalog)  GetProductBulkUploadHistory(xQuery PlatformGetProductBulkUploadHistoryXQuery) (ProductBulkRequestList, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getProductBulkUploadHistoryResponse ProductBulkRequestList
	    )

        

        
            
                
            
                
            
        

        

         
        
        
        //API call
        rawRequest = NewRequest(
            ca.Config,
            "get",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/products/bulk",ca.CompanyID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return ProductBulkRequestList{}, err
	    }
        
        err = json.Unmarshal(response, &getProductBulkUploadHistoryResponse)
        if err != nil {
             return ProductBulkRequestList{}, common.NewFDKError(err.Error())
        }
        return getProductBulkUploadHistoryResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // UpdateProductAssetsInBulk Create a Bulk asset upload Job.
     func (ca *PlatformCatalog)  UpdateProductAssetsInBulk(body  BulkJob) (SuccessResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            updateProductAssetsInBulkResponse SuccessResponse
	    )

        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        

        

        

         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
             
             return SuccessResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
               
             return SuccessResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            ca.Config,
            "post",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/products/bulk",ca.CompanyID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
             return SuccessResponse{}, err
	    }
        
        err = json.Unmarshal(response, &updateProductAssetsInBulkResponse)
        if err != nil {
             return SuccessResponse{}, common.NewFDKError(err.Error())
        }
        return updateProductAssetsInBulkResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // DeleteProductBulkJob Delete Bulk product job.
     func (ca *PlatformCatalog)  DeleteProductBulkJob(BatchID float64) (SuccessResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            deleteProductBulkJobResponse SuccessResponse
	    )

        

        

        
        
        

         
        
        
        //API call
        rawRequest = NewRequest(
            ca.Config,
            "delete",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/products/bulk/undefined",ca.CompanyID, BatchID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return SuccessResponse{}, err
	    }
        
        err = json.Unmarshal(response, &deleteProductBulkJobResponse)
        if err != nil {
             return SuccessResponse{}, common.NewFDKError(err.Error())
        }
        return deleteProductBulkJobResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // CreateProductsInBulk Create products in bulk associated with given batch Id.
     func (ca *PlatformCatalog)  CreateProductsInBulk(BatchID string, body  BulkProductRequest) (SuccessResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            createProductsInBulkResponse SuccessResponse
	    )

        
        
        
        
        
        
        
        
        

        

        
        
        

         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
             
             return SuccessResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
               
             return SuccessResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            ca.Config,
            "post",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/products/bulk/%s",ca.CompanyID, BatchID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
             return SuccessResponse{}, err
	    }
        
        err = json.Unmarshal(response, &createProductsInBulkResponse)
        if err != nil {
             return SuccessResponse{}, common.NewFDKError(err.Error())
        }
        return createProductsInBulkResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // GetProductTags Get a list of all tags associated with company.
     func (ca *PlatformCatalog)  GetProductTags() (ProductTagsViewResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getProductTagsResponse ProductTagsViewResponse
	    )

        

        

        

         
        
        
        //API call
        rawRequest = NewRequest(
            ca.Config,
            "get",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/products/tags",ca.CompanyID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return ProductTagsViewResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getProductTagsResponse)
        if err != nil {
             return ProductTagsViewResponse{}, common.NewFDKError(err.Error())
        }
        return getProductTagsResponse, nil
        
    }
         
        
       
    
    
   
  
    
    
    //PlatformGetProductAssetsInBulkXQuery holds query params
    type PlatformGetProductAssetsInBulkXQuery struct { 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"`  
    }
    


    // GetProductAssetsInBulk Get a list of all bulk asset jobs.
     func (ca *PlatformCatalog)  GetProductAssetsInBulk(xQuery PlatformGetProductAssetsInBulkXQuery) (BulkAssetResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getProductAssetsInBulkResponse BulkAssetResponse
	    )

        

        
            
                
            
                
            
        

        

         
        
        
        //API call
        rawRequest = NewRequest(
            ca.Config,
            "get",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/products/assets/bulk/",ca.CompanyID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return BulkAssetResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getProductAssetsInBulkResponse)
        if err != nil {
             return BulkAssetResponse{}, common.NewFDKError(err.Error())
        }
        return getProductAssetsInBulkResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // CreateProductAssetsInBulk Create a Bulk asset upload Job.
     func (ca *PlatformCatalog)  CreateProductAssetsInBulk(body  ProductBulkAssets) (SuccessResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            createProductAssetsInBulkResponse SuccessResponse
	    )

        
        
        
        
        
        
        

        

        

         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
             
             return SuccessResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
               
             return SuccessResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            ca.Config,
            "post",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/products/assets/bulk/",ca.CompanyID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
             return SuccessResponse{}, err
	    }
        
        err = json.Unmarshal(response, &createProductAssetsInBulkResponse)
        if err != nil {
             return SuccessResponse{}, common.NewFDKError(err.Error())
        }
        return createProductAssetsInBulkResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // DeleteSize Delete a Size associated with product.
     func (ca *PlatformCatalog)  DeleteSize(ItemID float64, Size float64) (ProductSizeDeleteResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            deleteSizeResponse ProductSizeDeleteResponse
	    )

        

        

        
        
        
        
        

         
        
        
        //API call
        rawRequest = NewRequest(
            ca.Config,
            "delete",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/products/undefined/sizes/undefined",ca.CompanyID, ItemID, Size),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return ProductSizeDeleteResponse{}, err
	    }
        
        err = json.Unmarshal(response, &deleteSizeResponse)
        if err != nil {
             return ProductSizeDeleteResponse{}, common.NewFDKError(err.Error())
        }
        return deleteSizeResponse, nil
        
    }
         
        
       
    
    
   
  
    
    
    //PlatformGetInventoryBySizeXQuery holds query params
    type PlatformGetInventoryBySizeXQuery struct { 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"` 
        Q string  `url:"q,omitempty"` 
        Sellable bool  `url:"sellable,omitempty"`  
    }
    


    // GetInventoryBySize Get Inventory for company
     func (ca *PlatformCatalog)  GetInventoryBySize(ItemID string, Size string, xQuery PlatformGetInventoryBySizeXQuery) (InventoryResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getInventoryBySizeResponse InventoryResponse
	    )

        

        
            
                
            
                
            
                
            
                
            
        

        
        
        
        
        

         
        
        
        //API call
        rawRequest = NewRequest(
            ca.Config,
            "get",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/products/%s/sizes/%s",ca.CompanyID, ItemID, Size),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return InventoryResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getInventoryBySizeResponse)
        if err != nil {
             return InventoryResponse{}, common.NewFDKError(err.Error())
        }
        return getInventoryBySizeResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // AddInventory Add Inventory for particular size and store.
     func (ca *PlatformCatalog)  AddInventory(ItemID float64, Size string, body  InventoryRequest) (SuccessResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            addInventoryResponse SuccessResponse
	    )

        
        
        
        
        
        
        

        

        
        
        
        
        

         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
             
             return SuccessResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
               
             return SuccessResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            ca.Config,
            "post",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/products/undefined/sizes/%s",ca.CompanyID, ItemID, Size),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
             return SuccessResponse{}, err
	    }
        
        err = json.Unmarshal(response, &addInventoryResponse)
        if err != nil {
             return SuccessResponse{}, common.NewFDKError(err.Error())
        }
        return addInventoryResponse, nil
        
    }
         
        
       
    
    
   
  
    
    
    //PlatformGetInventoryBySizeIdentifierXQuery holds query params
    type PlatformGetInventoryBySizeIdentifierXQuery struct { 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"` 
        Q string  `url:"q,omitempty"` 
        LocationIds []float64  `url:"location_ids,omitempty"`  
    }
    


    // GetInventoryBySizeIdentifier Get Inventory for company
     func (ca *PlatformCatalog)  GetInventoryBySizeIdentifier(ItemID string, SizeIdentifier string, xQuery PlatformGetInventoryBySizeIdentifierXQuery) (InventoryResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getInventoryBySizeIdentifierResponse InventoryResponse
	    )

        

        
            
                
            
                
            
                
            
                
            
        

        
        
        
        
        

         
        
        
        //API call
        rawRequest = NewRequest(
            ca.Config,
            "get",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/products/%s/inventory/%s",ca.CompanyID, ItemID, SizeIdentifier),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return InventoryResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getInventoryBySizeIdentifierResponse)
        if err != nil {
             return InventoryResponse{}, common.NewFDKError(err.Error())
        }
        return getInventoryBySizeIdentifierResponse, nil
        
    }
         
        
       
    
    
    
    
   
  
    
    


    // DeleteInventory Delete a Inventory.
     func (ca *PlatformCatalog)  DeleteInventory(Size string, ItemID float64, LocationID float64) (InventoryDelete, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            deleteInventoryResponse InventoryDelete
	    )

        

        

        
        
        
        
        
        
        

         
        
        
        //API call
        rawRequest = NewRequest(
            ca.Config,
            "delete",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/products/undefined/sizes/%s/location/undefined/",ca.CompanyID, Size, ItemID, LocationID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return InventoryDelete{}, err
	    }
        
        err = json.Unmarshal(response, &deleteInventoryResponse)
        if err != nil {
             return InventoryDelete{}, common.NewFDKError(err.Error())
        }
        return deleteInventoryResponse, nil
        
    }
         
        
       
    
    
   
  
    
    
    //PlatformGetInventoryBulkUploadHistoryXQuery holds query params
    type PlatformGetInventoryBulkUploadHistoryXQuery struct { 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"`  
    }
    


    // GetInventoryBulkUploadHistory Get a list of all bulk Inventory upload jobs.
     func (ca *PlatformCatalog)  GetInventoryBulkUploadHistory(xQuery PlatformGetInventoryBulkUploadHistoryXQuery) (BulkInventoryGet, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getInventoryBulkUploadHistoryResponse BulkInventoryGet
	    )

        

        
            
                
            
                
            
        

        

         
        
        
        //API call
        rawRequest = NewRequest(
            ca.Config,
            "get",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/inventory/bulk/",ca.CompanyID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return BulkInventoryGet{}, err
	    }
        
        err = json.Unmarshal(response, &getInventoryBulkUploadHistoryResponse)
        if err != nil {
             return BulkInventoryGet{}, common.NewFDKError(err.Error())
        }
        return getInventoryBulkUploadHistoryResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // CreateBulkInventoryJob Create a Bulk Inventory upload Job.
     func (ca *PlatformCatalog)  CreateBulkInventoryJob(body  BulkJob) (CommonResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            createBulkInventoryJobResponse CommonResponse
	    )

        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        

        

        

         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
             
             return CommonResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
               
             return CommonResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            ca.Config,
            "post",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/inventory/bulk/",ca.CompanyID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
             return CommonResponse{}, err
	    }
        
        err = json.Unmarshal(response, &createBulkInventoryJobResponse)
        if err != nil {
             return CommonResponse{}, common.NewFDKError(err.Error())
        }
        return createBulkInventoryJobResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // DeleteBulkInventoryJob Delete Bulk Inventory job.
     func (ca *PlatformCatalog)  DeleteBulkInventoryJob(BatchID string) (SuccessResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            deleteBulkInventoryJobResponse SuccessResponse
	    )

        

        

        
        
        

         
        
        
        //API call
        rawRequest = NewRequest(
            ca.Config,
            "delete",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/inventory/bulk/%s/",ca.CompanyID, BatchID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return SuccessResponse{}, err
	    }
        
        err = json.Unmarshal(response, &deleteBulkInventoryJobResponse)
        if err != nil {
             return SuccessResponse{}, common.NewFDKError(err.Error())
        }
        return deleteBulkInventoryJobResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // CreateBulkInventory Create products in bulk associated with given batch Id.
     func (ca *PlatformCatalog)  CreateBulkInventory(BatchID string, body  InventoryBulkRequest) (SuccessResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            createBulkInventoryResponse SuccessResponse
	    )

        
        
        
        
        
        
        
        
        

        

        
        
        

         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
             
             return SuccessResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
               
             return SuccessResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            ca.Config,
            "post",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/inventory/bulk/%s/",ca.CompanyID, BatchID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
             return SuccessResponse{}, err
	    }
        
        err = json.Unmarshal(response, &createBulkInventoryResponse)
        if err != nil {
             return SuccessResponse{}, common.NewFDKError(err.Error())
        }
        return createBulkInventoryResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // GetInventoryExport Get Inventory export history.
     func (ca *PlatformCatalog)  GetInventoryExport() (InventoryExportJob, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getInventoryExportResponse InventoryExportJob
	    )

        

        

        

         
        
        
        //API call
        rawRequest = NewRequest(
            ca.Config,
            "get",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/inventory/download/",ca.CompanyID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return InventoryExportJob{}, err
	    }
        
        err = json.Unmarshal(response, &getInventoryExportResponse)
        if err != nil {
             return InventoryExportJob{}, common.NewFDKError(err.Error())
        }
        return getInventoryExportResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // CreateInventoryExportJob Create a Inventory export Job.
     func (ca *PlatformCatalog)  CreateInventoryExportJob(body  InventoryExportRequest) (SuccessResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            createInventoryExportJobResponse SuccessResponse
	    )

        
        
        
        
        
        
        

        

        

         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
             
             return SuccessResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
               
             return SuccessResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            ca.Config,
            "post",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/inventory/download/",ca.CompanyID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
             return SuccessResponse{}, err
	    }
        
        err = json.Unmarshal(response, &createInventoryExportJobResponse)
        if err != nil {
             return SuccessResponse{}, common.NewFDKError(err.Error())
        }
        return createInventoryExportJobResponse, nil
        
    }
         
        
       
    
    
   
  
    
    
    //PlatformExportInventoryConfigXQuery holds query params
    type PlatformExportInventoryConfigXQuery struct { 
        FilterType string  `url:"filter_type,omitempty"`  
    }
    


    // ExportInventoryConfig Get List of different filters for inventory export
     func (ca *PlatformCatalog)  ExportInventoryConfig(xQuery PlatformExportInventoryConfigXQuery) (InventoryConfig, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            exportInventoryConfigResponse InventoryConfig
	    )

        

        
            
                
            
        

        

         
        
        
        //API call
        rawRequest = NewRequest(
            ca.Config,
            "get",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/inventory/download/configuration/",ca.CompanyID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return InventoryConfig{}, err
	    }
        
        err = json.Unmarshal(response, &exportInventoryConfigResponse)
        if err != nil {
             return InventoryConfig{}, common.NewFDKError(err.Error())
        }
        return exportInventoryConfigResponse, nil
        
    }
         
        
       
    
    
   
  
    
    
    //PlatformGetAllHsnCodesXQuery holds query params
    type PlatformGetAllHsnCodesXQuery struct { 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"` 
        Q string  `url:"q,omitempty"`  
    }
    


    // GetAllHsnCodes Hsn Code List.
     func (ca *PlatformCatalog)  GetAllHsnCodes(xQuery PlatformGetAllHsnCodesXQuery) (HsnCodesListingResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getAllHsnCodesResponse HsnCodesListingResponse
	    )

        

        
            
                
            
                
            
                
            
        

        

         
        
        
        //API call
        rawRequest = NewRequest(
            ca.Config,
            "get",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/hsn/",ca.CompanyID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return HsnCodesListingResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getAllHsnCodesResponse)
        if err != nil {
             return HsnCodesListingResponse{}, common.NewFDKError(err.Error())
        }
        return getAllHsnCodesResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // CreateHsnCode Create Hsn Code.
     func (ca *PlatformCatalog)  CreateHsnCode(body  HsnUpsert) (HsnCode, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            createHsnCodeResponse HsnCode
	    )

        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        

        

        

         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
             
             return HsnCode{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
               
             return HsnCode{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            ca.Config,
            "post",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/hsn/",ca.CompanyID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
             return HsnCode{}, err
	    }
        
        err = json.Unmarshal(response, &createHsnCodeResponse)
        if err != nil {
             return HsnCode{}, common.NewFDKError(err.Error())
        }
        return createHsnCodeResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // GetHsnCode Fetch Hsn Code.
     func (ca *PlatformCatalog)  GetHsnCode(ID string) (HsnCode, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getHsnCodeResponse HsnCode
	    )

        

        

        
        
        

         
        
        
        //API call
        rawRequest = NewRequest(
            ca.Config,
            "get",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/hsn/%s/",ca.CompanyID, ID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return HsnCode{}, err
	    }
        
        err = json.Unmarshal(response, &getHsnCodeResponse)
        if err != nil {
             return HsnCode{}, common.NewFDKError(err.Error())
        }
        return getHsnCodeResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // UpdateHsnCode Update Hsn Code.
     func (ca *PlatformCatalog)  UpdateHsnCode(ID string, body  HsnUpsert) (HsnCode, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            updateHsnCodeResponse HsnCode
	    )

        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        

        

        
        
        

         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
             
             return HsnCode{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
               
             return HsnCode{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            ca.Config,
            "put",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/hsn/%s/",ca.CompanyID, ID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
             return HsnCode{}, err
	    }
        
        err = json.Unmarshal(response, &updateHsnCodeResponse)
        if err != nil {
             return HsnCode{}, common.NewFDKError(err.Error())
        }
        return updateHsnCodeResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // BulkHsnCode Bulk Create or Update Hsn Code.
     func (ca *PlatformCatalog)  BulkHsnCode(body  BulkHsnUpsert) (BulkHsnResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            bulkHsnCodeResponse BulkHsnResponse
	    )

        
        
        

        

        

         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
             
             return BulkHsnResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
               
             return BulkHsnResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            ca.Config,
            "post",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/hsn/bulk/",ca.CompanyID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
             return BulkHsnResponse{}, err
	    }
        
        err = json.Unmarshal(response, &bulkHsnCodeResponse)
        if err != nil {
             return BulkHsnResponse{}, common.NewFDKError(err.Error())
        }
        return bulkHsnCodeResponse, nil
        
    }
         
        
       
    
    
    
    
    
    
    
    
    
    
    
    
    
    
   
  
    
    


    // GetOptimalLocations Location Reassignment
     func (ca *PlatformCatalog)  GetOptimalLocations(body  AssignStore) (StoreAssignResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getOptimalLocationsResponse StoreAssignResponse
	    )

        
        
        
        
        
        
        
        
        
        
        
        
        
        
        

        

        

         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
             
             return StoreAssignResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
               
             return StoreAssignResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            ca.Config,
            "post",
            fmt.Sprintf("/service/platform/catalog/v1.0/company/%s/location/reassign/",ca.CompanyID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
             return StoreAssignResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getOptimalLocationsResponse)
        if err != nil {
             return StoreAssignResponse{}, common.NewFDKError(err.Error())
        }
        return getOptimalLocationsResponse, nil
        
    }
         
        
       
    


	
   // PlatformCompanyProfile holds PlatformCompanyProfile object properties
    type PlatformCompanyProfile struct {
        Config *PlatformConfig
        CompanyID string
    }
    // NewPlatformCompanyProfile returns new PlatformCompanyProfile instance
    func NewPlatformCompanyProfile(config *PlatformConfig) *PlatformCompanyProfile {
        return &PlatformCompanyProfile{Config: config, CompanyID: config.CompanyID}
    }
    
    
   
  
    
    


    // CbsOnboardGet Get company profile
     func (co *PlatformCompanyProfile)  CbsOnboardGet() (GetCompanyProfileSerializerResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            cbsOnboardGetResponse GetCompanyProfileSerializerResponse
	    )

        

        

        

         
        
        
        //API call
        rawRequest = NewRequest(
            co.Config,
            "get",
            fmt.Sprintf("/service/platform/company-profile/v1.0/company/%s",co.CompanyID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return GetCompanyProfileSerializerResponse{}, err
	    }
        
        err = json.Unmarshal(response, &cbsOnboardGetResponse)
        if err != nil {
             return GetCompanyProfileSerializerResponse{}, common.NewFDKError(err.Error())
        }
        return cbsOnboardGetResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // UpdateCompany Edit company profile
     func (co *PlatformCompanyProfile)  UpdateCompany(body  UpdateCompany) (SuccessResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            updateCompanyResponse SuccessResponse
	    )

        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        

        

        

         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
             
             return SuccessResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
               
             return SuccessResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            co.Config,
            "patch",
            fmt.Sprintf("/service/platform/company-profile/v1.0/company/%s",co.CompanyID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
             return SuccessResponse{}, err
	    }
        
        err = json.Unmarshal(response, &updateCompanyResponse)
        if err != nil {
             return SuccessResponse{}, common.NewFDKError(err.Error())
        }
        return updateCompanyResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // GetCompanyMetrics Get company metrics
     func (co *PlatformCompanyProfile)  GetCompanyMetrics() (MetricsSerializer, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getCompanyMetricsResponse MetricsSerializer
	    )

        

        

        

         
        
        
        //API call
        rawRequest = NewRequest(
            co.Config,
            "get",
            fmt.Sprintf("/service/platform/company-profile/v1.0/company/%s/metrics",co.CompanyID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return MetricsSerializer{}, err
	    }
        
        err = json.Unmarshal(response, &getCompanyMetricsResponse)
        if err != nil {
             return MetricsSerializer{}, common.NewFDKError(err.Error())
        }
        return getCompanyMetricsResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // EditBrand Edit a brand.
     func (co *PlatformCompanyProfile)  EditBrand(BrandID string, body  CreateUpdateBrandRequestSerializer) (SuccessResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            editBrandResponse SuccessResponse
	    )

        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        

        

        
        
        

         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
             
             return SuccessResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
               
             return SuccessResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            co.Config,
            "put",
            fmt.Sprintf("/service/platform/company-profile/v1.0/company/%s/brand/%s",co.CompanyID, BrandID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
             return SuccessResponse{}, err
	    }
        
        err = json.Unmarshal(response, &editBrandResponse)
        if err != nil {
             return SuccessResponse{}, common.NewFDKError(err.Error())
        }
        return editBrandResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // GetBrand Get a single brand.
     func (co *PlatformCompanyProfile)  GetBrand(BrandID string) (GetBrandResponseSerializer, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getBrandResponse GetBrandResponseSerializer
	    )

        

        

        
        
        

         
        
        
        //API call
        rawRequest = NewRequest(
            co.Config,
            "get",
            fmt.Sprintf("/service/platform/company-profile/v1.0/company/%s/brand/%s",co.CompanyID, BrandID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return GetBrandResponseSerializer{}, err
	    }
        
        err = json.Unmarshal(response, &getBrandResponse)
        if err != nil {
             return GetBrandResponseSerializer{}, common.NewFDKError(err.Error())
        }
        return getBrandResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // CreateBrand Create a Brand.
     func (co *PlatformCompanyProfile)  CreateBrand(body  CreateUpdateBrandRequestSerializer) (SuccessResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            createBrandResponse SuccessResponse
	    )

        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        

        

        

         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
             
             return SuccessResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
               
             return SuccessResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            co.Config,
            "post",
            fmt.Sprintf("/service/platform/company-profile/v1.0/company/%s/brand",co.CompanyID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
             return SuccessResponse{}, err
	    }
        
        err = json.Unmarshal(response, &createBrandResponse)
        if err != nil {
             return SuccessResponse{}, common.NewFDKError(err.Error())
        }
        return createBrandResponse, nil
        
    }
         
        
       
    
    
   
  
    
    
    //PlatformGetBrandsXQuery holds query params
    type PlatformGetBrandsXQuery struct { 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"` 
        Q string  `url:"q,omitempty"`  
    }
    


    // GetBrands Get brands associated to a company
     func (co *PlatformCompanyProfile)  GetBrands(xQuery PlatformGetBrandsXQuery) (CompanyBrandListSerializer, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getBrandsResponse CompanyBrandListSerializer
	    )

        

        
            
                
            
                
            
                
            
        

        

         
        
        
        //API call
        rawRequest = NewRequest(
            co.Config,
            "get",
            fmt.Sprintf("/service/platform/company-profile/v1.0/company/%s/company-brand",co.CompanyID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return CompanyBrandListSerializer{}, err
	    }
        
        err = json.Unmarshal(response, &getBrandsResponse)
        if err != nil {
             return CompanyBrandListSerializer{}, common.NewFDKError(err.Error())
        }
        return getBrandsResponse, nil
        
    }
         
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
            
            // GetBrandsPaginator Get brands associated to a company  
            func (co *PlatformCompanyProfile)  GetBrandsPaginator(
              xQuery PlatformGetBrandsXQuery ) *common.Paginator {
                paginator := common.NewPaginator("number")
                
                 
                 xQuery.PageNo  = paginator.PageNo
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                paginator.Next = func() (interface{}, error) {
                    response, err := co.GetBrands(xQuery)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, int(response.Page.Current+1), response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }
        
       
    
    
   
  
    
    


    // CreateCompanyBrandMapping Create a company brand mapping.
     func (co *PlatformCompanyProfile)  CreateCompanyBrandMapping(body  CompanyBrandPostRequestSerializer) (SuccessResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            createCompanyBrandMappingResponse SuccessResponse
	    )

        
        
        
        
        
        
        

        

        

         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
             
             return SuccessResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
               
             return SuccessResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            co.Config,
            "post",
            fmt.Sprintf("/service/platform/company-profile/v1.0/company/%s/company-brand",co.CompanyID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
             return SuccessResponse{}, err
	    }
        
        err = json.Unmarshal(response, &createCompanyBrandMappingResponse)
        if err != nil {
             return SuccessResponse{}, common.NewFDKError(err.Error())
        }
        return createCompanyBrandMappingResponse, nil
        
    }
         
        
       
    
    
   
  
    
    
    //PlatformGetLocationsXQuery holds query params
    type PlatformGetLocationsXQuery struct { 
        StoreType string  `url:"store_type,omitempty"` 
        Q string  `url:"q,omitempty"` 
        Stage string  `url:"stage,omitempty"` 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"`  
    }
    


    // GetLocations Get list of locations
     func (co *PlatformCompanyProfile)  GetLocations(xQuery PlatformGetLocationsXQuery) (LocationListSerializer, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getLocationsResponse LocationListSerializer
	    )

        

        
            
                
            
                
            
                
            
                
            
                
            
        

        

         
        
        
        //API call
        rawRequest = NewRequest(
            co.Config,
            "get",
            fmt.Sprintf("/service/platform/company-profile/v1.0/company/%s/location",co.CompanyID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return LocationListSerializer{}, err
	    }
        
        err = json.Unmarshal(response, &getLocationsResponse)
        if err != nil {
             return LocationListSerializer{}, common.NewFDKError(err.Error())
        }
        return getLocationsResponse, nil
        
    }
         
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
            
            // GetLocationsPaginator Get list of locations  
            func (co *PlatformCompanyProfile)  GetLocationsPaginator(
              xQuery PlatformGetLocationsXQuery ) *common.Paginator {
                paginator := common.NewPaginator("number")
                
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 xQuery.PageNo  = paginator.PageNo
                 
                 
                 
                 
                 
                 
                 
                 
                paginator.Next = func() (interface{}, error) {
                    response, err := co.GetLocations(xQuery)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, int(response.Page.Current+1), response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }
        
       
    
    
   
  
    
    


    // CreateLocation Create a location asscoiated to a company.
     func (co *PlatformCompanyProfile)  CreateLocation(body  LocationSerializer) (SuccessResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            createLocationResponse SuccessResponse
	    )

        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        

        

        

         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
             
             return SuccessResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
               
             return SuccessResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            co.Config,
            "post",
            fmt.Sprintf("/service/platform/company-profile/v1.0/company/%s/location",co.CompanyID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
             return SuccessResponse{}, err
	    }
        
        err = json.Unmarshal(response, &createLocationResponse)
        if err != nil {
             return SuccessResponse{}, common.NewFDKError(err.Error())
        }
        return createLocationResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // UpdateLocation Edit a location asscoiated to a company.
     func (co *PlatformCompanyProfile)  UpdateLocation(LocationID string, body  LocationSerializer) (SuccessResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            updateLocationResponse SuccessResponse
	    )

        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        

        

        
        
        

         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
             
             return SuccessResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
               
             return SuccessResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            co.Config,
            "put",
            fmt.Sprintf("/service/platform/company-profile/v1.0/company/%s/location/%s",co.CompanyID, LocationID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
             return SuccessResponse{}, err
	    }
        
        err = json.Unmarshal(response, &updateLocationResponse)
        if err != nil {
             return SuccessResponse{}, common.NewFDKError(err.Error())
        }
        return updateLocationResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // GetLocationDetail Get details of a specific location.
     func (co *PlatformCompanyProfile)  GetLocationDetail(LocationID string) (GetLocationSerializer, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getLocationDetailResponse GetLocationSerializer
	    )

        

        

        
        
        

         
        
        
        //API call
        rawRequest = NewRequest(
            co.Config,
            "get",
            fmt.Sprintf("/service/platform/company-profile/v1.0/company/%s/location/%s",co.CompanyID, LocationID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return GetLocationSerializer{}, err
	    }
        
        err = json.Unmarshal(response, &getLocationDetailResponse)
        if err != nil {
             return GetLocationSerializer{}, common.NewFDKError(err.Error())
        }
        return getLocationDetailResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // CreateLocationBulk Create a location asscoiated to a company in bulk.
     func (co *PlatformCompanyProfile)  CreateLocationBulk(body  BulkLocationSerializer) (SuccessResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            createLocationBulkResponse SuccessResponse
	    )

        
        
        

        

        

         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
             
             return SuccessResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
               
             return SuccessResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            co.Config,
            "post",
            fmt.Sprintf("/service/platform/company-profile/v1.0/company/%s/location/bulk",co.CompanyID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
             return SuccessResponse{}, err
	    }
        
        err = json.Unmarshal(response, &createLocationBulkResponse)
        if err != nil {
             return SuccessResponse{}, common.NewFDKError(err.Error())
        }
        return createLocationBulkResponse, nil
        
    }
         
        
       
    


	
   // PlatformFileStorage holds PlatformFileStorage object properties
    type PlatformFileStorage struct {
        Config *PlatformConfig
        CompanyID string
    }
    // NewPlatformFileStorage returns new PlatformFileStorage instance
    func NewPlatformFileStorage(config *PlatformConfig) *PlatformFileStorage {
        return &PlatformFileStorage{Config: config, CompanyID: config.CompanyID}
    }
    
    
   
  
    
    


    // StartUpload This operation initiates upload and returns storage link which is valid for 30 Minutes. You can use that storage link to make subsequent upload request with file buffer or blob.
     func (fi *PlatformFileStorage)  StartUpload(Namespace string, body  StartRequest) (StartResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            startUploadResponse StartResponse
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
            fi.Config,
            "post",
            fmt.Sprintf("/service/platform/assets/v1.0/company/%s/namespaces/%s/upload/start/",Namespace, fi.CompanyID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
             return StartResponse{}, err
	    }
        
        err = json.Unmarshal(response, &startUploadResponse)
        if err != nil {
             return StartResponse{}, common.NewFDKError(err.Error())
        }
        return startUploadResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // CompleteUpload This will complete the upload process. After successfully uploading file, you can call this operation to complete the upload process.
     func (fi *PlatformFileStorage)  CompleteUpload(Namespace string, body  StartResponse) (CompleteResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            completeUploadResponse CompleteResponse
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
            fi.Config,
            "post",
            fmt.Sprintf("/service/platform/assets/v1.0/company/%s/namespaces/%s/upload/complete/",Namespace, fi.CompanyID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
             return CompleteResponse{}, err
	    }
        
        err = json.Unmarshal(response, &completeUploadResponse)
        if err != nil {
             return CompleteResponse{}, common.NewFDKError(err.Error())
        }
        return completeUploadResponse, nil
        
    }
         
        
       
    
    
    
    
    
    
   
  
    
    


    // GetSignUrls Explain here
     func (fi *PlatformFileStorage)  GetSignUrls(body  SignUrlRequest) (SignUrlResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getSignUrlsResponse SignUrlResponse
	    )

        
        
        
        
        

        

        

         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
             
             return SignUrlResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
               
             return SignUrlResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            fi.Config,
            "post",
            fmt.Sprintf("/service/platform/assets/v1.0/company/%s/sign-urls/",fi.CompanyID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
             return SignUrlResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getSignUrlsResponse)
        if err != nil {
             return SignUrlResponse{}, common.NewFDKError(err.Error())
        }
        return getSignUrlsResponse, nil
        
    }
         
        
       
    
    
   
  
    
    
    //PlatformCopyFilesXQuery holds query params
    type PlatformCopyFilesXQuery struct { 
        Sync bool  `url:"sync,omitempty"`  
    }
    


    // CopyFiles Copy Files
     func (fi *PlatformFileStorage)  CopyFiles(xQuery PlatformCopyFilesXQuery, body  BulkRequest) (BulkResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            copyFilesResponse BulkResponse
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
            fi.Config,
            "post",
            fmt.Sprintf("/service/platform/assets/v1.0/company/%s/uploads/copy/",fi.CompanyID),
            nil,
            xQuery,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
             return BulkResponse{}, err
	    }
        
        err = json.Unmarshal(response, &copyFilesResponse)
        if err != nil {
             return BulkResponse{}, common.NewFDKError(err.Error())
        }
        return copyFilesResponse, nil
        
    }
         
        
       
    
    
    
    
   
  
    
    
    //PlatformBrowseXQuery holds query params
    type PlatformBrowseXQuery struct { 
        PageNo float64  `url:"page_no,omitempty"`  
    }
    


    // Browse Browse Files
     func (fi *PlatformFileStorage)  Browse(Namespace string, xQuery PlatformBrowseXQuery) (BrowseResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            browseResponse BrowseResponse
	    )

        

        
            
                
            
        

        
        
        

         
        
        
        //API call
        rawRequest = NewRequest(
            fi.Config,
            "get",
            fmt.Sprintf("/service/platform/assets/v1.0/company/%s/namespaces/%s/browse/",Namespace, fi.CompanyID),
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
            func (fi *PlatformFileStorage)  BrowsePaginator(Namespace string  , 
              xQuery PlatformBrowseXQuery ) *common.Paginator {
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
        
       
    
    
    
    
   
  
    
    
    //PlatformProxyXQuery holds query params
    type PlatformProxyXQuery struct { 
        URL string  `url:"url,omitempty"`  
    }
    


    // Proxy Proxy
     func (fi *PlatformFileStorage)  Proxy(xQuery PlatformProxyXQuery) ([]byte, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
	    )

        

        
            
                
            
        

        

         
        
        
        //API call
        rawRequest = NewRequest(
            fi.Config,
            "post",
            fmt.Sprintf("/service/platform/assets/v1.0/company/%s/proxy/",fi.CompanyID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return []byte{}, err
	    }
        
        return response, nil
        
    }
         
        
       
    


	
   // PlatformInventory holds PlatformInventory object properties
    type PlatformInventory struct {
        Config *PlatformConfig
        CompanyID string
    }
    // NewPlatformInventory returns new PlatformInventory instance
    func NewPlatformInventory(config *PlatformConfig) *PlatformInventory {
        return &PlatformInventory{Config: config, CompanyID: config.CompanyID}
    }
    
    
   
  
    
    
    //PlatformGetJobsByCompanyXQuery holds query params
    type PlatformGetJobsByCompanyXQuery struct { 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"`  
    }
    


    // GetJobsByCompany Get Job Configs For A Company
     func (in *PlatformInventory)  GetJobsByCompany(xQuery PlatformGetJobsByCompanyXQuery) (ResponseEnvelopeListJobConfigRawDTO, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getJobsByCompanyResponse ResponseEnvelopeListJobConfigRawDTO
	    )

        

        
            
                
            
                
            
        

        

         
        
        
        //API call
        rawRequest = NewRequest(
            in.Config,
            "get",
            fmt.Sprintf("/service/platform/inventory/v1.0/company/%s/jobs",in.CompanyID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return ResponseEnvelopeListJobConfigRawDTO{}, err
	    }
        
        err = json.Unmarshal(response, &getJobsByCompanyResponse)
        if err != nil {
             return ResponseEnvelopeListJobConfigRawDTO{}, common.NewFDKError(err.Error())
        }
        return getJobsByCompanyResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // UpdateJob Updates An Existing Job Config
     func (in *PlatformInventory)  UpdateJob(body  JobConfigDTO) (ResponseEnvelopeString, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            updateJobResponse ResponseEnvelopeString
	    )

        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        

        

        

         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
             
             return ResponseEnvelopeString{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
               
             return ResponseEnvelopeString{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            in.Config,
            "put",
            fmt.Sprintf("/service/platform/inventory/v1.0/company/%s/jobs",in.CompanyID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
             return ResponseEnvelopeString{}, err
	    }
        
        err = json.Unmarshal(response, &updateJobResponse)
        if err != nil {
             return ResponseEnvelopeString{}, common.NewFDKError(err.Error())
        }
        return updateJobResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // CreateJob Creates A New Job Config
     func (in *PlatformInventory)  CreateJob(body  JobConfigDTO) (ResponseEnvelopeString, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            createJobResponse ResponseEnvelopeString
	    )

        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        

        

        

         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
             
             return ResponseEnvelopeString{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
               
             return ResponseEnvelopeString{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            in.Config,
            "post",
            fmt.Sprintf("/service/platform/inventory/v1.0/company/%s/jobs",in.CompanyID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
             return ResponseEnvelopeString{}, err
	    }
        
        err = json.Unmarshal(response, &createJobResponse)
        if err != nil {
             return ResponseEnvelopeString{}, common.NewFDKError(err.Error())
        }
        return createJobResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // GetJobSteps Get Job Code Steps
     func (in *PlatformInventory)  GetJobSteps(JobID float64) (ResponseEnvelopeListJobStepsDTO, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getJobStepsResponse ResponseEnvelopeListJobStepsDTO
	    )

        

        

        
        
        

         
        
        
        //API call
        rawRequest = NewRequest(
            in.Config,
            "get",
            fmt.Sprintf("/service/platform/inventory/v1.0/company/%s/jobs/steps/undefined",in.CompanyID, JobID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return ResponseEnvelopeListJobStepsDTO{}, err
	    }
        
        err = json.Unmarshal(response, &getJobStepsResponse)
        if err != nil {
             return ResponseEnvelopeListJobStepsDTO{}, common.NewFDKError(err.Error())
        }
        return getJobStepsResponse, nil
        
    }
         
        
       
    
    
   
  
    
    
    //PlatformGetJobByCompanyAndIntegrationXQuery holds query params
    type PlatformGetJobByCompanyAndIntegrationXQuery struct { 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"`  
    }
    


    // GetJobByCompanyAndIntegration Get Job Configs By Company And Integration
     func (in *PlatformInventory)  GetJobByCompanyAndIntegration(IntegrationID string, xQuery PlatformGetJobByCompanyAndIntegrationXQuery) (ResponseEnvelopeListJobConfigDTO, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getJobByCompanyAndIntegrationResponse ResponseEnvelopeListJobConfigDTO
	    )

        

        
            
                
            
                
            
        

        
        
        

         
        
        
        //API call
        rawRequest = NewRequest(
            in.Config,
            "get",
            fmt.Sprintf("/service/platform/inventory/v1.0/company/%s/jobs/integration/%s",in.CompanyID, IntegrationID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return ResponseEnvelopeListJobConfigDTO{}, err
	    }
        
        err = json.Unmarshal(response, &getJobByCompanyAndIntegrationResponse)
        if err != nil {
             return ResponseEnvelopeListJobConfigDTO{}, common.NewFDKError(err.Error())
        }
        return getJobByCompanyAndIntegrationResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // Disable Disable Job Config
     func (in *PlatformInventory)  Disable(IntegrationID string) (ResponseEnvelopeString, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            disableResponse ResponseEnvelopeString
	    )

        

        

        
        
        

         
        
        
        //API call
        rawRequest = NewRequest(
            in.Config,
            "get",
            fmt.Sprintf("/service/platform/inventory/v1.0/company/%s/jobs/disable/integration/%s",in.CompanyID, IntegrationID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return ResponseEnvelopeString{}, err
	    }
        
        err = json.Unmarshal(response, &disableResponse)
        if err != nil {
             return ResponseEnvelopeString{}, common.NewFDKError(err.Error())
        }
        return disableResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // GetJobConfigDefaults Get Job Configs Defaults
     func (in *PlatformInventory)  GetJobConfigDefaults() (ResponseEnvelopeJobConfigDTO, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getJobConfigDefaultsResponse ResponseEnvelopeJobConfigDTO
	    )

        

        

        

         
        
        
        //API call
        rawRequest = NewRequest(
            in.Config,
            "get",
            fmt.Sprintf("/service/platform/inventory/v1.0/company/%s/jobs/defaults",in.CompanyID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return ResponseEnvelopeJobConfigDTO{}, err
	    }
        
        err = json.Unmarshal(response, &getJobConfigDefaultsResponse)
        if err != nil {
             return ResponseEnvelopeJobConfigDTO{}, common.NewFDKError(err.Error())
        }
        return getJobConfigDefaultsResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // GetJobByCode Get Job Config By Code
     func (in *PlatformInventory)  GetJobByCode(Code string) (ResponseEnvelopeJobConfigDTO, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getJobByCodeResponse ResponseEnvelopeJobConfigDTO
	    )

        

        

        
        
        

         
        
        
        //API call
        rawRequest = NewRequest(
            in.Config,
            "get",
            fmt.Sprintf("/service/platform/inventory/v1.0/company/%s/jobs/code/%s",in.CompanyID, Code),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return ResponseEnvelopeJobConfigDTO{}, err
	    }
        
        err = json.Unmarshal(response, &getJobByCodeResponse)
        if err != nil {
             return ResponseEnvelopeJobConfigDTO{}, common.NewFDKError(err.Error())
        }
        return getJobByCodeResponse, nil
        
    }
         
        
       
    
    
   
  
    
    
    //PlatformGetJobCodeMetricsXQuery holds query params
    type PlatformGetJobCodeMetricsXQuery struct { 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"` 
        Status string  `url:"status,omitempty"` 
        Date string  `url:"date,omitempty"`  
    }
    


    // GetJobCodeMetrics Get Job Metrics
     func (in *PlatformInventory)  GetJobCodeMetrics(Code string, xQuery PlatformGetJobCodeMetricsXQuery) (ResponseEnvelopeJobMetricsDto, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getJobCodeMetricsResponse ResponseEnvelopeJobMetricsDto
	    )

        

        
            
                
            
                
            
                
            
                
            
        

        
        
        

         
        
        
        //API call
        rawRequest = NewRequest(
            in.Config,
            "get",
            fmt.Sprintf("/service/platform/inventory/v1.0/company/%s/jobs/code/%s/metrics",in.CompanyID, Code),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return ResponseEnvelopeJobMetricsDto{}, err
	    }
        
        err = json.Unmarshal(response, &getJobCodeMetricsResponse)
        if err != nil {
             return ResponseEnvelopeJobMetricsDto{}, common.NewFDKError(err.Error())
        }
        return getJobCodeMetricsResponse, nil
        
    }
         
        
       
    
    
   
  
    
    
    //PlatformGetJobCodesByCompanyAndIntegrationXQuery holds query params
    type PlatformGetJobCodesByCompanyAndIntegrationXQuery struct { 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"`  
    }
    


    // GetJobCodesByCompanyAndIntegration Get Job Codes By Company And Integration
     func (in *PlatformInventory)  GetJobCodesByCompanyAndIntegration(IntegrationID string, xQuery PlatformGetJobCodesByCompanyAndIntegrationXQuery) (ResponseEnvelopeListJobConfigListDTO, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getJobCodesByCompanyAndIntegrationResponse ResponseEnvelopeListJobConfigListDTO
	    )

        

        
            
                
            
                
            
        

        
        
        

         
        
        
        //API call
        rawRequest = NewRequest(
            in.Config,
            "get",
            fmt.Sprintf("/service/platform/inventory/v1.0/company/%s/jobs/code/integration/%s",in.CompanyID, IntegrationID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return ResponseEnvelopeListJobConfigListDTO{}, err
	    }
        
        err = json.Unmarshal(response, &getJobCodesByCompanyAndIntegrationResponse)
        if err != nil {
             return ResponseEnvelopeListJobConfigListDTO{}, common.NewFDKError(err.Error())
        }
        return getJobCodesByCompanyAndIntegrationResponse, nil
        
    }
         
        
       
    


	
   // PlatformConfiguration holds PlatformConfiguration object properties
    type PlatformConfiguration struct {
        Config *PlatformConfig
        CompanyID string
    }
    // NewPlatformConfiguration returns new PlatformConfiguration instance
    func NewPlatformConfiguration(config *PlatformConfig) *PlatformConfiguration {
        return &PlatformConfiguration{Config: config, CompanyID: config.CompanyID}
    }
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
   
  
    
    


    // CreateApplication Create application
     func (co *PlatformConfiguration)  CreateApplication(body  CreateApplicationRequest) (CreateAppResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            createApplicationResponse CreateAppResponse
	    )

        
        
        
        
        
        
        

        

        

         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
             
             return CreateAppResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
               
             return CreateAppResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            co.Config,
            "post",
            fmt.Sprintf("/service/platform/configuration/v1.0/company/%s/application",co.CompanyID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
             return CreateAppResponse{}, err
	    }
        
        err = json.Unmarshal(response, &createApplicationResponse)
        if err != nil {
             return CreateAppResponse{}, common.NewFDKError(err.Error())
        }
        return createApplicationResponse, nil
        
    }
         
        
       
    
    
   
  
    
    
    //PlatformGetApplicationsXQuery holds query params
    type PlatformGetApplicationsXQuery struct { 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"` 
        Q string  `url:"q,omitempty"`  
    }
    


    // GetApplications Get list of application under company
     func (co *PlatformConfiguration)  GetApplications(xQuery PlatformGetApplicationsXQuery) (ApplicationsResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getApplicationsResponse ApplicationsResponse
	    )

        

        
            
                
            
                
            
                
            
        

        

         
        
        
        //API call
        rawRequest = NewRequest(
            co.Config,
            "get",
            fmt.Sprintf("/service/platform/configuration/v1.0/company/%s/application",co.CompanyID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return ApplicationsResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getApplicationsResponse)
        if err != nil {
             return ApplicationsResponse{}, common.NewFDKError(err.Error())
        }
        return getApplicationsResponse, nil
        
    }
         
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
            
            // GetApplicationsPaginator Get list of application under company  
            func (co *PlatformConfiguration)  GetApplicationsPaginator(
              xQuery PlatformGetApplicationsXQuery ) *common.Paginator {
                paginator := common.NewPaginator("number")
                
                 
                 xQuery.PageNo  = paginator.PageNo
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                paginator.Next = func() (interface{}, error) {
                    response, err := co.GetApplications(xQuery)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, int(response.Page.Current+1), response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }
        
       
    
    
    
    
   
  
    
    


    // GetCurrencies Get all currencies
     func (co *PlatformConfiguration)  GetCurrencies() (CurrenciesResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getCurrenciesResponse CurrenciesResponse
	    )

        

        

        

         
        
        
        //API call
        rawRequest = NewRequest(
            co.Config,
            "get",
            fmt.Sprintf("/service/platform/configuration/v1.0/company/%s/currencies",co.CompanyID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return CurrenciesResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getCurrenciesResponse)
        if err != nil {
             return CurrenciesResponse{}, common.NewFDKError(err.Error())
        }
        return getCurrenciesResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // GetDomainAvailibility Check domain availibility before linking to application
     func (co *PlatformConfiguration)  GetDomainAvailibility(body  DomainSuggestionsRequest) (DomainSuggestionsResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getDomainAvailibilityResponse DomainSuggestionsResponse
	    )

        
        
        
        
        

        

        

         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
             
             return DomainSuggestionsResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
               
             return DomainSuggestionsResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            co.Config,
            "post",
            fmt.Sprintf("/service/platform/configuration/v1.0/company/%s/domain/suggestions",co.CompanyID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
             return DomainSuggestionsResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getDomainAvailibilityResponse)
        if err != nil {
             return DomainSuggestionsResponse{}, common.NewFDKError(err.Error())
        }
        return getDomainAvailibilityResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // GetIntegrationById Get integration data
     func (co *PlatformConfiguration)  GetIntegrationById(ID float64) (Integration, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getIntegrationByIdResponse Integration
	    )

        

        

        
        
        

         
        
        
        //API call
        rawRequest = NewRequest(
            co.Config,
            "get",
            fmt.Sprintf("/service/platform/configuration/v1.0/company/%s/integration/undefined",co.CompanyID, ID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return Integration{}, err
	    }
        
        err = json.Unmarshal(response, &getIntegrationByIdResponse)
        if err != nil {
             return Integration{}, common.NewFDKError(err.Error())
        }
        return getIntegrationByIdResponse, nil
        
    }
         
        
       
    
    
   
  
    
    
    //PlatformGetAvailableOptInsXQuery holds query params
    type PlatformGetAvailableOptInsXQuery struct { 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"`  
    }
    


    // GetAvailableOptIns Get all available integration opt-ins
     func (co *PlatformConfiguration)  GetAvailableOptIns(xQuery PlatformGetAvailableOptInsXQuery) (GetIntegrationsOptInsResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getAvailableOptInsResponse GetIntegrationsOptInsResponse
	    )

        

        
            
                
            
                
            
        

        

         
        
        
        //API call
        rawRequest = NewRequest(
            co.Config,
            "get",
            fmt.Sprintf("/service/platform/configuration/v1.0/company/%s/integration-opt-in/available",co.CompanyID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return GetIntegrationsOptInsResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getAvailableOptInsResponse)
        if err != nil {
             return GetIntegrationsOptInsResponse{}, common.NewFDKError(err.Error())
        }
        return getAvailableOptInsResponse, nil
        
    }
         
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
            
            // GetAvailableOptInsPaginator Get all available integration opt-ins  
            func (co *PlatformConfiguration)  GetAvailableOptInsPaginator(
              xQuery PlatformGetAvailableOptInsXQuery ) *common.Paginator {
                paginator := common.NewPaginator("number")
                
                 
                 xQuery.PageNo  = paginator.PageNo
                 
                 
                 
                 
                 
                 
                 
                 
                paginator.Next = func() (interface{}, error) {
                    response, err := co.GetAvailableOptIns(xQuery)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, int(response.Page.Current+1), response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }
        
       
    
    
   
  
    
    
    //PlatformGetSelectedOptInsXQuery holds query params
    type PlatformGetSelectedOptInsXQuery struct { 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"`  
    }
    


    // GetSelectedOptIns Get company/store level integration opt-ins
     func (co *PlatformConfiguration)  GetSelectedOptIns(Level string, UID float64, xQuery PlatformGetSelectedOptInsXQuery) (GetIntegrationsOptInsResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getSelectedOptInsResponse GetIntegrationsOptInsResponse
	    )

        

        
            
                
            
                
            
        

        
        
        
        
        

         
        
        
        //API call
        rawRequest = NewRequest(
            co.Config,
            "get",
            fmt.Sprintf("/service/platform/configuration/v1.0/company/%s/integration-opt-in/selected/%s/undefined",co.CompanyID, Level, UID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return GetIntegrationsOptInsResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getSelectedOptInsResponse)
        if err != nil {
             return GetIntegrationsOptInsResponse{}, common.NewFDKError(err.Error())
        }
        return getSelectedOptInsResponse, nil
        
    }
         
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
            
            // GetSelectedOptInsPaginator Get company/store level integration opt-ins  
            func (co *PlatformConfiguration)  GetSelectedOptInsPaginator(Level string  , UID float64  , 
              xQuery PlatformGetSelectedOptInsXQuery ) *common.Paginator {
                paginator := common.NewPaginator("number")
                
                 
                 xQuery.PageNo  = paginator.PageNo
                 
                 
                 
                 
                 
                 
                 
                 
                paginator.Next = func() (interface{}, error) {
                    response, err := co.GetSelectedOptIns(Level, UID, xQuery)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, int(response.Page.Current+1), response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }
        
       
    
    
   
  
    
    
    //PlatformGetIntegrationLevelConfigXQuery holds query params
    type PlatformGetIntegrationLevelConfigXQuery struct { 
        Opted bool  `url:"opted,omitempty"` 
        CheckPermission bool  `url:"check_permission,omitempty"`  
    }
    


    // GetIntegrationLevelConfig Get integration level config
     func (co *PlatformConfiguration)  GetIntegrationLevelConfig(ID string, Level string, xQuery PlatformGetIntegrationLevelConfigXQuery) (IntegrationConfigResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getIntegrationLevelConfigResponse IntegrationConfigResponse
	    )

        

        
            
                
            
                
            
        

        
        
        
        
        

         
        
        
        //API call
        rawRequest = NewRequest(
            co.Config,
            "get",
            fmt.Sprintf("/service/platform/configuration/v1.0/company/%s/integration-opt-in/configuration/%s/%s",co.CompanyID, ID, Level),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return IntegrationConfigResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getIntegrationLevelConfigResponse)
        if err != nil {
             return IntegrationConfigResponse{}, common.NewFDKError(err.Error())
        }
        return getIntegrationLevelConfigResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // UpdateLevelIntegration Update a store level opt-in for integration
     func (co *PlatformConfiguration)  UpdateLevelIntegration(ID string, Level string, body  UpdateIntegrationLevelRequest) (IntegrationLevel, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            updateLevelIntegrationResponse IntegrationLevel
	    )

        
        
        

        

        
        
        
        
        

         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
             
             return IntegrationLevel{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
               
             return IntegrationLevel{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            co.Config,
            "put",
            fmt.Sprintf("/service/platform/configuration/v1.0/company/%s/integration-opt-in/configuration/%s/%s",co.CompanyID, ID, Level),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
             return IntegrationLevel{}, err
	    }
        
        err = json.Unmarshal(response, &updateLevelIntegrationResponse)
        if err != nil {
             return IntegrationLevel{}, common.NewFDKError(err.Error())
        }
        return updateLevelIntegrationResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // GetIntegrationByLevelId Get level data for integration
     func (co *PlatformConfiguration)  GetIntegrationByLevelId(ID string, Level string, UID float64) (IntegrationLevel, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getIntegrationByLevelIdResponse IntegrationLevel
	    )

        

        

        
        
        
        
        
        
        

         
        
        
        //API call
        rawRequest = NewRequest(
            co.Config,
            "get",
            fmt.Sprintf("/service/platform/configuration/v1.0/company/%s/integration-opt-in/configuration/%s/%s/undefined",co.CompanyID, ID, Level, UID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return IntegrationLevel{}, err
	    }
        
        err = json.Unmarshal(response, &getIntegrationByLevelIdResponse)
        if err != nil {
             return IntegrationLevel{}, common.NewFDKError(err.Error())
        }
        return getIntegrationByLevelIdResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // UpdateLevelUidIntegration Update a store level opt-in for integration
     func (co *PlatformConfiguration)  UpdateLevelUidIntegration(ID string, Level string, UID float64, body  IntegrationLevel) (IntegrationLevel, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            updateLevelUidIntegrationResponse IntegrationLevel
	    )

        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        

        

        
        
        
        
        
        
        

         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
             
             return IntegrationLevel{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
               
             return IntegrationLevel{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            co.Config,
            "put",
            fmt.Sprintf("/service/platform/configuration/v1.0/company/%s/integration-opt-in/configuration/%s/%s/undefined",co.CompanyID, ID, Level, UID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
             return IntegrationLevel{}, err
	    }
        
        err = json.Unmarshal(response, &updateLevelUidIntegrationResponse)
        if err != nil {
             return IntegrationLevel{}, common.NewFDKError(err.Error())
        }
        return updateLevelUidIntegrationResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // GetLevelActiveIntegrations Check store has active integration
     func (co *PlatformConfiguration)  GetLevelActiveIntegrations(ID string, Level string, UID float64) (OptedStoreIntegration, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getLevelActiveIntegrationsResponse OptedStoreIntegration
	    )

        

        

        
        
        
        
        
        
        

         
        
        
        //API call
        rawRequest = NewRequest(
            co.Config,
            "get",
            fmt.Sprintf("/service/platform/configuration/v1.0/company/%s/integration-opt-in/check/configuration/%s/%s/undefined",co.CompanyID, ID, Level, UID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return OptedStoreIntegration{}, err
	    }
        
        err = json.Unmarshal(response, &getLevelActiveIntegrationsResponse)
        if err != nil {
             return OptedStoreIntegration{}, common.NewFDKError(err.Error())
        }
        return getLevelActiveIntegrationsResponse, nil
        
    }
         
        
       
    
    
   
  
    
    
    //PlatformGetBrandsByCompanyXQuery holds query params
    type PlatformGetBrandsByCompanyXQuery struct { 
        Q string  `url:"q,omitempty"`  
    }
    


    // GetBrandsByCompany Get brands by company
     func (co *PlatformConfiguration)  GetBrandsByCompany(xQuery PlatformGetBrandsByCompanyXQuery) (BrandsByCompanyResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getBrandsByCompanyResponse BrandsByCompanyResponse
	    )

        

        
            
                
            
        

        

         
        
        
        //API call
        rawRequest = NewRequest(
            co.Config,
            "get",
            fmt.Sprintf("/service/platform/configuration/v1.0/company/%s/inventory/brands-by-companies",co.CompanyID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return BrandsByCompanyResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getBrandsByCompanyResponse)
        if err != nil {
             return BrandsByCompanyResponse{}, common.NewFDKError(err.Error())
        }
        return getBrandsByCompanyResponse, nil
        
    }
         
        
       
    
    
   
  
    
    
    //PlatformGetCompanyByBrandsXQuery holds query params
    type PlatformGetCompanyByBrandsXQuery struct { 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"`  
    }
    


    // GetCompanyByBrands Get company by brand uids
     func (co *PlatformConfiguration)  GetCompanyByBrands(xQuery PlatformGetCompanyByBrandsXQuery, body  CompanyByBrandsRequest) (CompanyByBrandsResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getCompanyByBrandsResponse CompanyByBrandsResponse
	    )

        
        
        
        
        

        
            
                
            
                
            
        

        

         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
             
             return CompanyByBrandsResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
               
             return CompanyByBrandsResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            co.Config,
            "post",
            fmt.Sprintf("/service/platform/configuration/v1.0/company/%s/inventory/companies-by-brands",co.CompanyID),
            nil,
            xQuery,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
             return CompanyByBrandsResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getCompanyByBrandsResponse)
        if err != nil {
             return CompanyByBrandsResponse{}, common.NewFDKError(err.Error())
        }
        return getCompanyByBrandsResponse, nil
        
    }
         
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
            
            // GetCompanyByBrandsPaginator Get company by brand uids  
            func (co *PlatformConfiguration)  GetCompanyByBrandsPaginator(
              xQuery PlatformGetCompanyByBrandsXQuery  , body  CompanyByBrandsRequest) *common.Paginator {
                paginator := common.NewPaginator("number")
                
                 
                 xQuery.PageNo  = paginator.PageNo
                 
                 
                 
                 
                 
                 
                 
                 
                paginator.Next = func() (interface{}, error) {
                    response, err := co.GetCompanyByBrands(xQuery, body)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, int(response.Page.Current+1), response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }
        
       
    
    
   
  
    
    
    //PlatformGetStoreByBrandsXQuery holds query params
    type PlatformGetStoreByBrandsXQuery struct { 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"`  
    }
    


    // GetStoreByBrands Get stores by brand uids
     func (co *PlatformConfiguration)  GetStoreByBrands(xQuery PlatformGetStoreByBrandsXQuery, body  StoreByBrandsRequest) (StoreByBrandsResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getStoreByBrandsResponse StoreByBrandsResponse
	    )

        
        
        
        
        
        
        

        
            
                
            
                
            
        

        

         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
             
             return StoreByBrandsResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
               
             return StoreByBrandsResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            co.Config,
            "post",
            fmt.Sprintf("/service/platform/configuration/v1.0/company/%s/inventory/stores-by-brands",co.CompanyID),
            nil,
            xQuery,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
             return StoreByBrandsResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getStoreByBrandsResponse)
        if err != nil {
             return StoreByBrandsResponse{}, common.NewFDKError(err.Error())
        }
        return getStoreByBrandsResponse, nil
        
    }
         
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
            
            // GetStoreByBrandsPaginator Get stores by brand uids  
            func (co *PlatformConfiguration)  GetStoreByBrandsPaginator(
              xQuery PlatformGetStoreByBrandsXQuery  , body  StoreByBrandsRequest) *common.Paginator {
                paginator := common.NewPaginator("number")
                
                 
                 xQuery.PageNo  = paginator.PageNo
                 
                 
                 
                 
                 
                 
                 
                 
                paginator.Next = func() (interface{}, error) {
                    response, err := co.GetStoreByBrands(xQuery, body)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, int(response.Page.Current+1), response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }
        
       
    
    
   
  
    
    
    //PlatformGetOtherSellerApplicationsXQuery holds query params
    type PlatformGetOtherSellerApplicationsXQuery struct { 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"`  
    }
    


    // GetOtherSellerApplications Get other seller applications
     func (co *PlatformConfiguration)  GetOtherSellerApplications(xQuery PlatformGetOtherSellerApplicationsXQuery) (OtherSellerApplications, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getOtherSellerApplicationsResponse OtherSellerApplications
	    )

        

        
            
                
            
                
            
        

        

         
        
        
        //API call
        rawRequest = NewRequest(
            co.Config,
            "get",
            fmt.Sprintf("/service/platform/configuration/v1.0/company/%s/other-seller-applications/",co.CompanyID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return OtherSellerApplications{}, err
	    }
        
        err = json.Unmarshal(response, &getOtherSellerApplicationsResponse)
        if err != nil {
             return OtherSellerApplications{}, common.NewFDKError(err.Error())
        }
        return getOtherSellerApplicationsResponse, nil
        
    }
         
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
            
            // GetOtherSellerApplicationsPaginator Get other seller applications  
            func (co *PlatformConfiguration)  GetOtherSellerApplicationsPaginator(
              xQuery PlatformGetOtherSellerApplicationsXQuery ) *common.Paginator {
                paginator := common.NewPaginator("number")
                
                 
                 xQuery.PageNo  = paginator.PageNo
                 
                 
                 
                 
                 
                 
                 
                 
                paginator.Next = func() (interface{}, error) {
                    response, err := co.GetOtherSellerApplications(xQuery)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, int(response.Page.Current+1), response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }
        
       
    
    
   
  
    
    


    // GetOtherSellerApplicationById Get other seller applications
     func (co *PlatformConfiguration)  GetOtherSellerApplicationById(ID string) (OptedApplicationResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getOtherSellerApplicationByIdResponse OptedApplicationResponse
	    )

        

        

        
        
        

         
        
        
        //API call
        rawRequest = NewRequest(
            co.Config,
            "get",
            fmt.Sprintf("/service/platform/configuration/v1.0/company/%s/other-seller-applications/%s",co.CompanyID, ID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return OptedApplicationResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getOtherSellerApplicationByIdResponse)
        if err != nil {
             return OptedApplicationResponse{}, common.NewFDKError(err.Error())
        }
        return getOtherSellerApplicationByIdResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // OptOutFromApplication Opt out company or store from other seller application
     func (co *PlatformConfiguration)  OptOutFromApplication(ID string, body  OptOutInventory) (SuccessMessageResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            optOutFromApplicationResponse SuccessMessageResponse
	    )

        
        
        
        
        

        

        
        
        

         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
             
             return SuccessMessageResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
               
             return SuccessMessageResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            co.Config,
            "put",
            fmt.Sprintf("/service/platform/configuration/v1.0/company/%s/other-seller-applications/%s/opt_out",co.CompanyID, ID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
             return SuccessMessageResponse{}, err
	    }
        
        err = json.Unmarshal(response, &optOutFromApplicationResponse)
        if err != nil {
             return SuccessMessageResponse{}, common.NewFDKError(err.Error())
        }
        return optOutFromApplicationResponse, nil
        
    }
         
        
       
    


	
   // PlatformAnalytics holds PlatformAnalytics object properties
    type PlatformAnalytics struct {
        Config *PlatformConfig
        CompanyID string
    }
    // NewPlatformAnalytics returns new PlatformAnalytics instance
    func NewPlatformAnalytics(config *PlatformConfig) *PlatformAnalytics {
        return &PlatformAnalytics{Config: config, CompanyID: config.CompanyID}
    }
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
   
  
    
    


    // CreateExportJob Create data export job in required format
     func (an *PlatformAnalytics)  CreateExportJob(ExportType string, body  ExportJobReq) (ExportJobRes, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            createExportJobResponse ExportJobRes
	    )

        
        
        
        
        
        
        
        
        
        
        

        

        
        
        

         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
             
             return ExportJobRes{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
               
             return ExportJobRes{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            an.Config,
            "post",
            fmt.Sprintf("/service/platform/analytics/v1.0/company/%s/export/%s",an.CompanyID, ExportType),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
             return ExportJobRes{}, err
	    }
        
        err = json.Unmarshal(response, &createExportJobResponse)
        if err != nil {
             return ExportJobRes{}, common.NewFDKError(err.Error())
        }
        return createExportJobResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // GetExportJobStatus Get data export job status
     func (an *PlatformAnalytics)  GetExportJobStatus(ExportType string, JobID string) (ExportJobStatusRes, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getExportJobStatusResponse ExportJobStatusRes
	    )

        

        

        
        
        
        
        

         
        
        
        //API call
        rawRequest = NewRequest(
            an.Config,
            "get",
            fmt.Sprintf("/service/platform/analytics/v1.0/company/%s/export/%s/job/%s",an.CompanyID, ExportType, JobID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return ExportJobStatusRes{}, err
	    }
        
        err = json.Unmarshal(response, &getExportJobStatusResponse)
        if err != nil {
             return ExportJobStatusRes{}, common.NewFDKError(err.Error())
        }
        return getExportJobStatusResponse, nil
        
    }
         
        
       
    
    
   
  
    
    
    //PlatformGetLogsListXQuery holds query params
    type PlatformGetLogsListXQuery struct { 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"`  
    }
    


    // GetLogsList Get logs list
     func (an *PlatformAnalytics)  GetLogsList(LogType string, xQuery PlatformGetLogsListXQuery, body  GetLogsListReq) (GetLogsListRes, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getLogsListResponse GetLogsListRes
	    )

        
        
        
        
        
        
        
        
        

        
            
                
            
                
            
        

        
        
        

         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
             
             return GetLogsListRes{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
               
             return GetLogsListRes{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            an.Config,
            "post",
            fmt.Sprintf("/service/platform/analytics/v1.0/company/%s/logs/%s",an.CompanyID, LogType),
            nil,
            xQuery,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
             return GetLogsListRes{}, err
	    }
        
        err = json.Unmarshal(response, &getLogsListResponse)
        if err != nil {
             return GetLogsListRes{}, common.NewFDKError(err.Error())
        }
        return getLogsListResponse, nil
        
    }
         
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
            
            // GetLogsListPaginator Get logs list  
            func (an *PlatformAnalytics)  GetLogsListPaginator(LogType string  , 
              xQuery PlatformGetLogsListXQuery  , body  GetLogsListReq) *common.Paginator {
                paginator := common.NewPaginator("number")
                
                 
                 xQuery.PageNo  = paginator.PageNo
                 
                 
                 
                 
                 
                 
                 
                 
                paginator.Next = func() (interface{}, error) {
                    response, err := an.GetLogsList(LogType, xQuery, body)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, int(response.Page.Current+1), response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }
        
       
    
    
   
  
    
    
    //PlatformSearchLogsXQuery holds query params
    type PlatformSearchLogsXQuery struct { 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"`  
    }
    


    // SearchLogs Search logs
     func (an *PlatformAnalytics)  SearchLogs(LogType string, xQuery PlatformSearchLogsXQuery, body  SearchLogReq) (SearchLogRes, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            searchLogsResponse SearchLogRes
	    )

        
        
        
        
        
        
        
        
        
        
        
        
        

        
            
                
            
                
            
        

        
        
        

         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
             
             return SearchLogRes{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
               
             return SearchLogRes{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            an.Config,
            "post",
            fmt.Sprintf("/service/platform/analytics/v1.0/company/%s/logs/%s/search",an.CompanyID, LogType),
            nil,
            xQuery,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
             return SearchLogRes{}, err
	    }
        
        err = json.Unmarshal(response, &searchLogsResponse)
        if err != nil {
             return SearchLogRes{}, common.NewFDKError(err.Error())
        }
        return searchLogsResponse, nil
        
    }
         
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
            
            // SearchLogsPaginator Search logs  
            func (an *PlatformAnalytics)  SearchLogsPaginator(LogType string  , 
              xQuery PlatformSearchLogsXQuery  , body  SearchLogReq) *common.Paginator {
                paginator := common.NewPaginator("number")
                
                 
                 xQuery.PageNo  = paginator.PageNo
                 
                 
                 
                 
                 
                 
                 
                 
                paginator.Next = func() (interface{}, error) {
                    response, err := an.SearchLogs(LogType, xQuery, body)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, int(response.Page.Current+1), response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }
        
       
    


	
   // PlatformDiscount holds PlatformDiscount object properties
    type PlatformDiscount struct {
        Config *PlatformConfig
        CompanyID string
    }
    // NewPlatformDiscount returns new PlatformDiscount instance
    func NewPlatformDiscount(config *PlatformConfig) *PlatformDiscount {
        return &PlatformDiscount{Config: config, CompanyID: config.CompanyID}
    }
    
    
   
  
    
    
    //PlatformGetDiscountsXQuery holds query params
    type PlatformGetDiscountsXQuery struct { 
        View string  `url:"view,omitempty"` 
        Q string  `url:"q,omitempty"` 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"` 
        Archived bool  `url:"archived,omitempty"` 
        Month float64  `url:"month,omitempty"` 
        Year float64  `url:"year,omitempty"` 
        Type string  `url:"type,omitempty"` 
        AppIds []string  `url:"app_ids,omitempty"`  
    }
    


    // GetDiscounts Fetch discount list.
     func (di *PlatformDiscount)  GetDiscounts(xQuery PlatformGetDiscountsXQuery) (ListOrCalender, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getDiscountsResponse ListOrCalender
	    )

        

        
            
                
            
                
            
                
            
                
            
                
            
                
            
                
            
                
            
                
            
        

        

         
        
        
        //API call
        rawRequest = NewRequest(
            di.Config,
            "get",
            fmt.Sprintf("/service/platform/discount/v1.0/company/%s/job/",di.CompanyID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return ListOrCalender{}, err
	    }
        
        err = json.Unmarshal(response, &getDiscountsResponse)
        if err != nil {
             return ListOrCalender{}, common.NewFDKError(err.Error())
        }
        return getDiscountsResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // CreateDiscount Create Discount.
     func (di *PlatformDiscount)  CreateDiscount(body  CreateUpdateDiscount) (DiscountJob, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            createDiscountResponse DiscountJob
	    )

        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        

        

        

         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
             
             return DiscountJob{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
               
             return DiscountJob{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            di.Config,
            "post",
            fmt.Sprintf("/service/platform/discount/v1.0/company/%s/job/",di.CompanyID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
             return DiscountJob{}, err
	    }
        
        err = json.Unmarshal(response, &createDiscountResponse)
        if err != nil {
             return DiscountJob{}, common.NewFDKError(err.Error())
        }
        return createDiscountResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // GetDiscount Fetch discount.
     func (di *PlatformDiscount)  GetDiscount(ID string) (DiscountJob, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getDiscountResponse DiscountJob
	    )

        

        

        
        
        

         
        
        
        //API call
        rawRequest = NewRequest(
            di.Config,
            "get",
            fmt.Sprintf("/service/platform/discount/v1.0/company/%s/job/%s/",di.CompanyID, ID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return DiscountJob{}, err
	    }
        
        err = json.Unmarshal(response, &getDiscountResponse)
        if err != nil {
             return DiscountJob{}, common.NewFDKError(err.Error())
        }
        return getDiscountResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // UpdateDiscount Create Discount.
     func (di *PlatformDiscount)  UpdateDiscount(ID string, body  CreateUpdateDiscount) (DiscountJob, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            updateDiscountResponse DiscountJob
	    )

        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        

        

        
        
        

         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
             
             return DiscountJob{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
               
             return DiscountJob{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            di.Config,
            "put",
            fmt.Sprintf("/service/platform/discount/v1.0/company/%s/job/%s/",di.CompanyID, ID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
             return DiscountJob{}, err
	    }
        
        err = json.Unmarshal(response, &updateDiscountResponse)
        if err != nil {
             return DiscountJob{}, common.NewFDKError(err.Error())
        }
        return updateDiscountResponse, nil
        
    }
         
        
       
    
    
   
  
    
    
    //PlatformValidateDiscountFileXQuery holds query params
    type PlatformValidateDiscountFileXQuery struct { 
        Discount string  `url:"discount,omitempty"`  
    }
    


    // ValidateDiscountFile Validate File.
     func (di *PlatformDiscount)  ValidateDiscountFile(xQuery PlatformValidateDiscountFileXQuery, body  DiscountJob) (FileJobResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            validateDiscountFileResponse FileJobResponse
	    )

        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        

        
            
                
            
        

        

         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
             
             return FileJobResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
               
             return FileJobResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            di.Config,
            "post",
            fmt.Sprintf("/service/platform/discount/v1.0/company/%s/file/validation/",di.CompanyID),
            nil,
            xQuery,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
             return FileJobResponse{}, err
	    }
        
        err = json.Unmarshal(response, &validateDiscountFileResponse)
        if err != nil {
             return FileJobResponse{}, common.NewFDKError(err.Error())
        }
        return validateDiscountFileResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // DownloadDiscountFile Validate File.
     func (di *PlatformDiscount)  DownloadDiscountFile(Type string, body  DownloadFileJob) (FileJobResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            downloadDiscountFileResponse FileJobResponse
	    )

        
        
        
        
        

        

        
        
        

         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
             
             return FileJobResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
               
             return FileJobResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            di.Config,
            "post",
            fmt.Sprintf("/service/platform/discount/v1.0/company/%s/file/%s/download/",di.CompanyID, Type),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
             return FileJobResponse{}, err
	    }
        
        err = json.Unmarshal(response, &downloadDiscountFileResponse)
        if err != nil {
             return FileJobResponse{}, common.NewFDKError(err.Error())
        }
        return downloadDiscountFileResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // GetValidationJob Validate File Job.
     func (di *PlatformDiscount)  GetValidationJob(ID string) (FileJobResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getValidationJobResponse FileJobResponse
	    )

        

        

        
        
        

         
        
        
        //API call
        rawRequest = NewRequest(
            di.Config,
            "get",
            fmt.Sprintf("/service/platform/discount/v1.0/company/%s/file/validation/%s/",di.CompanyID, ID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return FileJobResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getValidationJobResponse)
        if err != nil {
             return FileJobResponse{}, common.NewFDKError(err.Error())
        }
        return getValidationJobResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // CancelValidationJob Cancel Validation Job.
     func (di *PlatformDiscount)  CancelValidationJob(ID string) (CancelJobResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            cancelValidationJobResponse CancelJobResponse
	    )

        

        

        
        
        

         
        
        
        //API call
        rawRequest = NewRequest(
            di.Config,
            "delete",
            fmt.Sprintf("/service/platform/discount/v1.0/company/%s/file/validation/%s/",di.CompanyID, ID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return CancelJobResponse{}, err
	    }
        
        err = json.Unmarshal(response, &cancelValidationJobResponse)
        if err != nil {
             return CancelJobResponse{}, common.NewFDKError(err.Error())
        }
        return cancelValidationJobResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // GetDownloadJob Download File Job.
     func (di *PlatformDiscount)  GetDownloadJob(ID string) (FileJobResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getDownloadJobResponse FileJobResponse
	    )

        

        

        
        
        

         
        
        
        //API call
        rawRequest = NewRequest(
            di.Config,
            "get",
            fmt.Sprintf("/service/platform/discount/v1.0/company/%s/file/download/%s/",di.CompanyID, ID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return FileJobResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getDownloadJobResponse)
        if err != nil {
             return FileJobResponse{}, common.NewFDKError(err.Error())
        }
        return getDownloadJobResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // CancelDownloadJob Cancel Download Job.
     func (di *PlatformDiscount)  CancelDownloadJob(ID string) (CancelJobResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            cancelDownloadJobResponse CancelJobResponse
	    )

        

        

        
        
        

         
        
        
        //API call
        rawRequest = NewRequest(
            di.Config,
            "delete",
            fmt.Sprintf("/service/platform/discount/v1.0/company/%s/file/download/%s/",di.CompanyID, ID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return CancelJobResponse{}, err
	    }
        
        err = json.Unmarshal(response, &cancelDownloadJobResponse)
        if err != nil {
             return CancelJobResponse{}, common.NewFDKError(err.Error())
        }
        return cancelDownloadJobResponse, nil
        
    }
         
        
       
    


	
   // PlatformWebhook holds PlatformWebhook object properties
    type PlatformWebhook struct {
        Config *PlatformConfig
        CompanyID string
    }
    // NewPlatformWebhook returns new PlatformWebhook instance
    func NewPlatformWebhook(config *PlatformConfig) *PlatformWebhook {
        return &PlatformWebhook{Config: config, CompanyID: config.CompanyID}
    }
    
    
   
  
    
    
    //PlatformGetSubscribersByCompanyXQuery holds query params
    type PlatformGetSubscribersByCompanyXQuery struct { 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"` 
        ExtensionID string  `url:"extension_id,omitempty"`  
    }
    


    // GetSubscribersByCompany Get Subscribers By Company ID
     func (we *PlatformWebhook)  GetSubscribersByCompany(xQuery PlatformGetSubscribersByCompanyXQuery) (SubscriberResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getSubscribersByCompanyResponse SubscriberResponse
	    )

        

        
            
                
            
                
            
                
            
        

        

         
        
        
        //API call
        rawRequest = NewRequest(
            we.Config,
            "get",
            fmt.Sprintf("/service/platform/webhook/v1.0/company/%s/subscriber",we.CompanyID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return SubscriberResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getSubscribersByCompanyResponse)
        if err != nil {
             return SubscriberResponse{}, common.NewFDKError(err.Error())
        }
        return getSubscribersByCompanyResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // RegisterSubscriberToEvent Register Subscriber
     func (we *PlatformWebhook)  RegisterSubscriberToEvent(body  SubscriberConfig) (SubscriberConfig, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            registerSubscriberToEventResponse SubscriberConfig
	    )

        
        
        
        
        
        
        
        
        
        
        //enum validation inside request body
        err = body.Status.IsValid()
        if err != nil {
             
             return SubscriberConfig{}, common.NewFDKError(err.Error())
        }
        
        
        
        
        
        
        
        

        

        

         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
             
             return SubscriberConfig{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
               
             return SubscriberConfig{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            we.Config,
            "post",
            fmt.Sprintf("/service/platform/webhook/v1.0/company/%s/subscriber",we.CompanyID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
             return SubscriberConfig{}, err
	    }
        
        err = json.Unmarshal(response, &registerSubscriberToEventResponse)
        if err != nil {
             return SubscriberConfig{}, common.NewFDKError(err.Error())
        }
        return registerSubscriberToEventResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // UpdateSubscriberConfig Update Subscriber
     func (we *PlatformWebhook)  UpdateSubscriberConfig(body  SubscriberConfig) (SubscriberConfig, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            updateSubscriberConfigResponse SubscriberConfig
	    )

        
        
        
        
        
        
        
        
        
        
        //enum validation inside request body
        err = body.Status.IsValid()
        if err != nil {
             
             return SubscriberConfig{}, common.NewFDKError(err.Error())
        }
        
        
        
        
        
        
        
        

        

        

         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
             
             return SubscriberConfig{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
               
             return SubscriberConfig{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            we.Config,
            "put",
            fmt.Sprintf("/service/platform/webhook/v1.0/company/%s/subscriber",we.CompanyID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
             return SubscriberConfig{}, err
	    }
        
        err = json.Unmarshal(response, &updateSubscriberConfigResponse)
        if err != nil {
             return SubscriberConfig{}, common.NewFDKError(err.Error())
        }
        return updateSubscriberConfigResponse, nil
        
    }
         
        
       
    
    
   
  
    
    
    //PlatformGetSubscribersByExtensionIdXQuery holds query params
    type PlatformGetSubscribersByExtensionIdXQuery struct { 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"`  
    }
    


    // GetSubscribersByExtensionId Get Subscribers By Extension ID
     func (we *PlatformWebhook)  GetSubscribersByExtensionId(ExtensionID string, xQuery PlatformGetSubscribersByExtensionIdXQuery) (SubscriberResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getSubscribersByExtensionIdResponse SubscriberResponse
	    )

        

        
            
                
            
                
            
        

        
        
        

         
        
        
        //API call
        rawRequest = NewRequest(
            we.Config,
            "get",
            fmt.Sprintf("/service/platform/webhook/v1.0/company/%s/extension/%s/subscriber",we.CompanyID, ExtensionID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return SubscriberResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getSubscribersByExtensionIdResponse)
        if err != nil {
             return SubscriberResponse{}, common.NewFDKError(err.Error())
        }
        return getSubscribersByExtensionIdResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // GetSubscriberById Get Subscriber By Subscriber ID
     func (we *PlatformWebhook)  GetSubscriberById(SubscriberID float64) (SubscriberResponse, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            getSubscriberByIdResponse SubscriberResponse
	    )

        

        

        
        
        

         
        
        
        //API call
        rawRequest = NewRequest(
            we.Config,
            "get",
            fmt.Sprintf("/service/platform/webhook/v1.0/company/%s/subscriber/undefined",we.CompanyID, SubscriberID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return SubscriberResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getSubscriberByIdResponse)
        if err != nil {
             return SubscriberResponse{}, common.NewFDKError(err.Error())
        }
        return getSubscriberByIdResponse, nil
        
    }
         
        
       
    
    
   
  
    
    


    // FetchAllEventConfigurations Get All Webhook Events
     func (we *PlatformWebhook)  FetchAllEventConfigurations() (EventConfigList, error){
        
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            fetchAllEventConfigurationsResponse EventConfigList
	    )

        

        

        

         
        
        
        //API call
        rawRequest = NewRequest(
            we.Config,
            "get",
            fmt.Sprintf("/service/platform/webhook/v1.0/company/%s/events",we.CompanyID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
             return EventConfigList{}, err
	    }
        
        err = json.Unmarshal(response, &fetchAllEventConfigurationsResponse)
        if err != nil {
             return EventConfigList{}, common.NewFDKError(err.Error())
        }
        return fetchAllEventConfigurationsResponse, nil
        
    }
         
        
       
    


