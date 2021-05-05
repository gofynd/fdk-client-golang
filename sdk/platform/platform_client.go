package platform



import (
	"encoding/json"
	"fmt"
	"github.com/gofynd/fdk-client-golang/sdk/common"
)
// PlatformClient holds all the PlatformClient object properties  
type PlatformClient struct {
	config            *PlatformConfig
	
		Lead  *PlatformLead
	
		Billing  *PlatformBilling
	
		Communication  *PlatformCommunication
	
		Payment  *PlatformPayment
	
		Order  *PlatformOrder
	
		CompanyProfile  *PlatformCompanyProfile
	
		FileStorage  *PlatformFileStorage
	
		Inventory  *PlatformInventory
	
		Configuration  *PlatformConfiguration
	
		Marketplaces  *PlatformMarketplaces
	
		Analytics  *PlatformAnalytics
	
		Discount  *PlatformDiscount
	
		Webhook  *PlatformWebhook
	
	ApplicationClient *ApplicationClient //ApplicationClient embedded
}

// NewPlatformClient returns platform company level service instances
func NewPlatformClient(config *PlatformConfig) *PlatformClient {
		return &PlatformClient{
			config: config,
			
				Lead:  NewPlatformLead(config),
			
				Billing:  NewPlatformBilling(config),
			
				Communication:  NewPlatformCommunication(config),
			
				Payment:  NewPlatformPayment(config),
			
				Order:  NewPlatformOrder(config),
			
				CompanyProfile:  NewPlatformCompanyProfile(config),
			
				FileStorage:  NewPlatformFileStorage(config),
			
				Inventory:  NewPlatformInventory(config),
			
				Configuration:  NewPlatformConfiguration(config),
			
				Marketplaces:  NewPlatformMarketplaces(config),
			
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
        Priority string  `url:"priority,omitempty"` 
        Category string  `url:"category,omitempty"` 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"`  
    }
    
    // GetTickets Gets the list of company level tickets and/or ticket filters depending on query params
     func (le *PlatformLead)  GetTickets(
         xQuery PlatformGetTicketsXQuery) (TicketList, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getTicketsResponse TicketList
            
	    )
     
         
        
        
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
     func (le *PlatformLead)  CreateTicket(
         body  AddTicketPayload) (Ticket, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            createTicketResponse Ticket
            
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
     func (le *PlatformLead)  GetTicket(
         TicketID string) (Ticket, error){
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
            fmt.Sprintf("/service/platform/lead/v1.0/company/%s/ticket/%s",le.CompanyID, TicketID),
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
     func (le *PlatformLead)  EditTicket(
         TicketID string, body  EditTicketPayload) (Ticket, error){
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
            le.Config,
            "put",
            fmt.Sprintf("/service/platform/lead/v1.0/company/%s/ticket/%s",le.CompanyID, TicketID),
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
     func (le *PlatformLead)  CreateHistory(
         TicketID string, body  TicketHistoryPayload) (TicketHistory, error){
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
            le.Config,
            "post",
            fmt.Sprintf("/service/platform/lead/v1.0/company/%s/ticket/%s/history",le.CompanyID, TicketID),
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
     func (le *PlatformLead)  GetTicketHistory(
         TicketID string) (TicketHistoryList, error){
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
            fmt.Sprintf("/service/platform/lead/v1.0/company/%s/ticket/%s/history",le.CompanyID, TicketID),
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
         
        
       
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    


	
   // PlatformBilling holds PlatformBilling object properties
    type PlatformBilling struct {
        Config *PlatformConfig
        CompanyID string
    }
    // NewPlatformBilling returns new PlatformBilling instance
    func NewPlatformBilling(config *PlatformConfig) *PlatformBilling {
        return &PlatformBilling{Config: config, CompanyID: config.CompanyID}
    }
    
    
  
    
    
    // GetInvoices Get invoices
     func (bi *PlatformBilling)  GetInvoices(
         ) (Invoices, error){
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
     func (bi *PlatformBilling)  GetInvoiceById(
         InvoiceID string) (Invoice, error){
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
     func (bi *PlatformBilling)  GetCustomerDetail(
         ) (SubscriptionCustomer, error){
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
     func (bi *PlatformBilling)  UpsertCustomerDetail(
         body  SubscriptionCustomerCreate) (SubscriptionCustomer, error){
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
     func (bi *PlatformBilling)  GetSubscription(
         ) (SubscriptionStatus, error){
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
     func (bi *PlatformBilling)  GetFeatureLimitConfig(
         ) (SubscriptionLimit, error){
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
     func (bi *PlatformBilling)  ActivateSubscriptionPlan(
         body  SubscriptionActivateReq) (SubscriptionActivateRes, error){
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
     func (bi *PlatformBilling)  CancelSubscriptionPlan(
         body  CancelSubscriptionReq) (CancelSubscriptionRes, error){
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
     func (co *PlatformCommunication)  GetSystemNotifications(
         xQuery PlatformGetSystemNotificationsXQuery) (SystemNotifications, error){
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
     func (pa *PlatformPayment)  GetAllPayouts(
         xQuery PlatformGetAllPayoutsXQuery) (PayoutsResponse, error){
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
     func (pa *PlatformPayment)  SavePayout(
         body  PayoutRequest) (PayoutResponse, error){
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
     func (pa *PlatformPayment)  UpdatePayout(
         UniqueTransferNo string, body  PayoutRequest) (UpdatePayoutResponse, error){
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
     func (pa *PlatformPayment)  ActivateAndDectivatePayout(
         UniqueTransferNo string, body  UpdatePayoutRequest) (UpdatePayoutResponse, error){
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
     func (pa *PlatformPayment)  DeletePayout(
         UniqueTransferNo string) (DeletePayoutResponse, error){
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
         
        
       
    
    
  
    
    
    // GetSubscriptionPaymentMethod List Subscription Payment Method
     func (pa *PlatformPayment)  GetSubscriptionPaymentMethod(
         ) (SubscriptionPaymentMethodResponse, error){
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
            nil,
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
     func (pa *PlatformPayment)  DeleteSubscriptionPaymentMethod(
         xQuery PlatformDeleteSubscriptionPaymentMethodXQuery) (DeleteSubscriptionPaymentMethodResponse, error){
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
     func (pa *PlatformPayment)  GetSubscriptionConfig(
         ) (SubscriptionConfigResponse, error){
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
     func (pa *PlatformPayment)  SaveSubscriptionSetupIntent(
         body  SaveSubscriptionSetupIntentRequest) (SaveSubscriptionSetupIntentResponse, error){
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
     func (or *PlatformOrder)  ShipmentStatusUpdate(
         body  UpdateShipmentStatusBody) (UpdateShipmentStatusResponse, error){
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
     func (or *PlatformOrder)  ActivityStatus(
         xQuery PlatformActivityStatusXQuery) (GetActivityStatus, error){
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
     func (or *PlatformOrder)  StoreProcessShipmentUpdate(
         body  UpdateProcessShipmenstRequestBody) (UpdateProcessShipmenstRequestResponse, error){
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
         
        
       
    
    
  
    
    
    //PlatformGetOrdersByCompanyIdXQuery holds query params
    type PlatformGetOrdersByCompanyIdXQuery struct { 
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
    
    // GetOrdersByCompanyId Get Orders for company based on Company Id
     func (or *PlatformOrder)  GetOrdersByCompanyId(
         xQuery PlatformGetOrdersByCompanyIdXQuery) (OrderListing, error){
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
         
        
       
    
    
    
    
    
    
    
    
    
    
  
    
    
    // GetPing Get Ping
     func (or *PlatformOrder)  GetPing(
         ) (GetPingResponse, error){
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
     func (or *PlatformOrder)  VoiceCallback(
         ) (GetVoiceCallbackResponse, error){
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
     func (or *PlatformOrder)  VoiceClickToCall(
         xQuery PlatformVoiceClickToCallXQuery) (GetClickToCallResponse, error){
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
         
        
       
    


	
   // PlatformCompanyProfile holds PlatformCompanyProfile object properties
    type PlatformCompanyProfile struct {
        Config *PlatformConfig
        CompanyID string
    }
    // NewPlatformCompanyProfile returns new PlatformCompanyProfile instance
    func NewPlatformCompanyProfile(config *PlatformConfig) *PlatformCompanyProfile {
        return &PlatformCompanyProfile{Config: config, CompanyID: config.CompanyID}
    }
    
    
  
    
    
    // UpdateCompany Edit company profile
     func (co *PlatformCompanyProfile)  UpdateCompany(
         body  CompanyStoreSerializerRequest) (SuccessResponse, error){
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
         
        
       
    
    
  
    
    
    // CbsOnboardGet Get company profile
     func (co *PlatformCompanyProfile)  CbsOnboardGet(
         ) (GetCompanyProfileSerializerResponse, error){
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
         
        
       
    
    
  
    
    
    // GetCompanyMetrics Get company metrics
     func (co *PlatformCompanyProfile)  GetCompanyMetrics(
         ) (MetricsSerializer, error){
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
         
        
       
    
    
  
    
    
    // GetBrand Get a single brand.
     func (co *PlatformCompanyProfile)  GetBrand(
         BrandID string) (GetBrandResponseSerializer, error){
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
         
        
       
    
    
  
    
    
    // EditBrand Edit a brand.
     func (co *PlatformCompanyProfile)  EditBrand(
         BrandID string, body  CreateUpdateBrandRequestSerializer) (SuccessResponse, error){
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
         
        
       
    
    
  
    
    
    // CreateBrand Create a Brand.
     func (co *PlatformCompanyProfile)  CreateBrand(
         body  CreateUpdateBrandRequestSerializer) (SuccessResponse, error){
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
         
        
       
    
    
  
    
    
    // CreateCompanyBrandMapping Create a company brand mapping.
     func (co *PlatformCompanyProfile)  CreateCompanyBrandMapping(
         body  CompanyBrandPostRequestSerializer) (SuccessResponse, error){
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
         
        
       
    
    
  
    
    
    //PlatformGetBrandsXQuery holds query params
    type PlatformGetBrandsXQuery struct { 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"`  
    }
    
    // GetBrands Get brands associated to a company
     func (co *PlatformCompanyProfile)  GetBrands(
         xQuery PlatformGetBrandsXQuery) (CompanyBrandListSerializer, error){
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
        
       
    
    
  
    
    
    // CreateLocation Create a location asscoiated to a company.
     func (co *PlatformCompanyProfile)  CreateLocation(
         body  LocationSerializer) (SuccessResponse, error){
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
         
        
       
    
    
  
    
    
    //PlatformGetLocationsXQuery holds query params
    type PlatformGetLocationsXQuery struct { 
        StoreType string  `url:"store_type,omitempty"` 
        Q string  `url:"q,omitempty"` 
        Stage string  `url:"stage,omitempty"` 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"`  
    }
    
    // GetLocations Get list of locations
     func (co *PlatformCompanyProfile)  GetLocations(
         xQuery PlatformGetLocationsXQuery) (LocationListSerializer, error){
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
        
       
    
    
  
    
    
    // GetLocationDetail Get details of a specific location.
     func (co *PlatformCompanyProfile)  GetLocationDetail(
         LocationID string) (GetLocationSerializer, error){
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
         
        
       
    
    
  
    
    
    // UpdateLocation Edit a location asscoiated to a company.
     func (co *PlatformCompanyProfile)  UpdateLocation(
         LocationID string, body  LocationSerializer) (SuccessResponse, error){
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
     func (fi *PlatformFileStorage)  StartUpload(
         Namespace string, body  StartRequest) (StartResponse, error){
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
     func (fi *PlatformFileStorage)  CompleteUpload(
         Namespace string, body  StartResponse) (CompleteResponse, error){
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
     func (fi *PlatformFileStorage)  GetSignUrls(
         body  SignUrlRequest) (SignUrlResponse, error){
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
     func (fi *PlatformFileStorage)  CopyFiles(
         xQuery PlatformCopyFilesXQuery, body  BulkRequest) (BulkResponse, error){
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
     func (fi *PlatformFileStorage)  Browse(
         Namespace string, xQuery PlatformBrowseXQuery) (BrowseResponse, error){
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
     func (fi *PlatformFileStorage)  Proxy(
         xQuery PlatformProxyXQuery) (string, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            proxyResponse string
            
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
	 	   
             return "", err
            
	    }
        err = json.Unmarshal(response, &proxyResponse)
        if err != nil {
           
             return "", common.NewFDKError(err.Error())
            
        }
        return proxyResponse, nil
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
     func (in *PlatformInventory)  GetJobsByCompany(
         xQuery PlatformGetJobsByCompanyXQuery) (ResponseEnvelopeListJobConfigRawDTO, error){
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
     func (in *PlatformInventory)  UpdateJob(
         XUserData string, body  JobConfigDTO) (ResponseEnvelopeString, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            updateJobResponse ResponseEnvelopeString
            
	    )
     
        
        //Adding extra headers
        var xHeaders = make(map[string]string) 
        
        xHeaders["x-user-data"]=  XUserData;
        
        
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
            xHeaders,
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
     func (in *PlatformInventory)  CreateJob(
         XUserData string, body  JobConfigDTO) (ResponseEnvelopeString, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            createJobResponse ResponseEnvelopeString
            
	    )
     
        
        //Adding extra headers
        var xHeaders = make(map[string]string) 
        
        xHeaders["x-user-data"]=  XUserData;
        
        
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
            xHeaders,
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
         
        
       
    
    
  
    
    
    //PlatformGetJobByCompanyAndIntegrationXQuery holds query params
    type PlatformGetJobByCompanyAndIntegrationXQuery struct { 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"`  
    }
    
    // GetJobByCompanyAndIntegration Get Job Configs By Company And Integration
     func (in *PlatformInventory)  GetJobByCompanyAndIntegration(
         IntegrationID string, xQuery PlatformGetJobByCompanyAndIntegrationXQuery) (ResponseEnvelopeListJobConfigDTO, error){
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
         
        
       
    
    
  
    
    
    // GetJobConfigDefaults Get Job Configs Defaults
     func (in *PlatformInventory)  GetJobConfigDefaults(
         ) (ResponseEnvelopeJobConfigDTO, error){
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
     func (in *PlatformInventory)  GetJobByCode(
         Code string) (ResponseEnvelopeJobConfigDTO, error){
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
         
        
       
    
    
  
    
    
    //PlatformGetJobCodesByCompanyAndIntegrationXQuery holds query params
    type PlatformGetJobCodesByCompanyAndIntegrationXQuery struct { 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"`  
    }
    
    // GetJobCodesByCompanyAndIntegration Get Job Codes By Company And Integration
     func (in *PlatformInventory)  GetJobCodesByCompanyAndIntegration(
         IntegrationID string, xQuery PlatformGetJobCodesByCompanyAndIntegrationXQuery) (ResponseEnvelopeListJobConfigListDTO, error){
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
     func (co *PlatformConfiguration)  CreateApplication(
         body  CreateApplicationRequest) (CreateAppResponse, error){
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
     func (co *PlatformConfiguration)  GetApplications(
         xQuery PlatformGetApplicationsXQuery) (ApplicationsResponse, error){
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
     func (co *PlatformConfiguration)  GetCurrencies(
         ) (CurrenciesResponse, error){
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
     func (co *PlatformConfiguration)  GetDomainAvailibility(
         body  DomainSuggestionsRequest) (DomainSuggestionsResponse, error){
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
     func (co *PlatformConfiguration)  GetIntegrationById(
         ID float64) (Integration, error){
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
     func (co *PlatformConfiguration)  GetAvailableOptIns(
         xQuery PlatformGetAvailableOptInsXQuery) (GetIntegrationsOptInsResponse, error){
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
     func (co *PlatformConfiguration)  GetSelectedOptIns(
         Level string, UID float64, xQuery PlatformGetSelectedOptInsXQuery) (GetIntegrationsOptInsResponse, error){
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
     func (co *PlatformConfiguration)  GetIntegrationLevelConfig(
         ID string, Level string, xQuery PlatformGetIntegrationLevelConfigXQuery) (IntegrationConfigResponse, error){
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
         
        
       
    
    
  
    
    
    // GetIntegrationByLevelId Get level data for integration
     func (co *PlatformConfiguration)  GetIntegrationByLevelId(
         ID string, Level string, UID float64) (IntegrationLevel, error){
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
         
        
       
    
    
  
    
    
    // GetLevelActiveIntegrations Check store has active integration
     func (co *PlatformConfiguration)  GetLevelActiveIntegrations(
         ID string, Level string, UID float64) (OptedStoreIntegration, error){
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
     func (co *PlatformConfiguration)  GetBrandsByCompany(
         xQuery PlatformGetBrandsByCompanyXQuery) (BrandsByCompanyResponse, error){
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
     func (co *PlatformConfiguration)  GetCompanyByBrands(
         xQuery PlatformGetCompanyByBrandsXQuery, body  CompanyByBrandsRequest) (CompanyByBrandsResponse, error){
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
     func (co *PlatformConfiguration)  GetStoreByBrands(
         xQuery PlatformGetStoreByBrandsXQuery, body  StoreByBrandsRequest) (StoreByBrandsResponse, error){
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
     func (co *PlatformConfiguration)  GetOtherSellerApplications(
         xQuery PlatformGetOtherSellerApplicationsXQuery) (OtherSellerApplications, error){
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
     func (co *PlatformConfiguration)  GetOtherSellerApplicationById(
         ID string) (OptedApplicationResponse, error){
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
     func (co *PlatformConfiguration)  OptOutFromApplication(
         ID string, body  OptOutInventory) (SuccessMessageResponse, error){
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
         
        
       
    


	
   // PlatformMarketplaces holds PlatformMarketplaces object properties
    type PlatformMarketplaces struct {
        Config *PlatformConfig
        CompanyID string
    }
    // NewPlatformMarketplaces returns new PlatformMarketplaces instance
    func NewPlatformMarketplaces(config *PlatformConfig) *PlatformMarketplaces {
        return &PlatformMarketplaces{Config: config, CompanyID: config.CompanyID}
    }
    
    
  
    
    
    // GetAvailableChannels Get available marketplace channels
     func (ma *PlatformMarketplaces)  GetAvailableChannels(
         ) (*map[string]interface{}, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getAvailableChannelsResponse *map[string]interface{}
            
	    )
     
         
        
        
        //API call
        rawRequest = NewRequest(
            ma.Config,
            "get",
            fmt.Sprintf("/service/platform/marketplaces/v1.0/company/%s/all-channels",ma.CompanyID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &map[string]interface{}{}, err
            
	    }
        err = json.Unmarshal(response, &getAvailableChannelsResponse)
        if err != nil {
           
             return &map[string]interface{}{}, common.NewFDKError(err.Error())
            
        }
        return getAvailableChannelsResponse, nil
    }
         
        
       
    
    
  
    
    
    // GetChannels Get all registered marketplace channels for a seller
     func (ma *PlatformMarketplaces)  GetChannels(
         ) (*map[string]interface{}, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getChannelsResponse *map[string]interface{}
            
	    )
     
         
        
        
        //API call
        rawRequest = NewRequest(
            ma.Config,
            "get",
            fmt.Sprintf("/service/platform/marketplaces/v1.0/company/%s/channels/",ma.CompanyID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &map[string]interface{}{}, err
            
	    }
        err = json.Unmarshal(response, &getChannelsResponse)
        if err != nil {
           
             return &map[string]interface{}{}, common.NewFDKError(err.Error())
            
        }
        return getChannelsResponse, nil
    }
         
        
       
    
    
  
    
    
    // GetChannel Get registered marketplace channel credential configuration for a seller
     func (ma *PlatformMarketplaces)  GetChannel(
         Channel string) (*map[string]interface{}, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getChannelResponse *map[string]interface{}
            
	    )
     
         
        
        
        //API call
        rawRequest = NewRequest(
            ma.Config,
            "get",
            fmt.Sprintf("/service/platform/marketplaces/v1.0/company/%s/channels/%s",ma.CompanyID, Channel),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &map[string]interface{}{}, err
            
	    }
        err = json.Unmarshal(response, &getChannelResponse)
        if err != nil {
           
             return &map[string]interface{}{}, common.NewFDKError(err.Error())
            
        }
        return getChannelResponse, nil
    }
         
        
       
    
    
  
    
    
    // RegisterMyntraChannel Create Myntra marketplace channel for a seller
     func (ma *PlatformMarketplaces)  RegisterMyntraChannel(
         body  MyntraPayload) (MkpResp, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            registerMyntraChannelResponse MkpResp
            
	    )
     
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return MkpResp{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return MkpResp{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            ma.Config,
            "post",
            fmt.Sprintf("/service/platform/marketplaces/v1.0/company/%s/channels/myntra_in",ma.CompanyID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return MkpResp{}, err
            
	    }
        err = json.Unmarshal(response, &registerMyntraChannelResponse)
        if err != nil {
           
             return MkpResp{}, common.NewFDKError(err.Error())
            
        }
        return registerMyntraChannelResponse, nil
    }
         
        
       
    
    
  
    
    
    // UpdateMyntraChannelCredentials Update Myntra marketplace channel credentials for a seller
     func (ma *PlatformMarketplaces)  UpdateMyntraChannelCredentials(
         body  MyntraPayload) (MkpResp, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            updateMyntraChannelCredentialsResponse MkpResp
            
	    )
     
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return MkpResp{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return MkpResp{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            ma.Config,
            "put",
            fmt.Sprintf("/service/platform/marketplaces/v1.0/company/%s/channels/myntra_in",ma.CompanyID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return MkpResp{}, err
            
	    }
        err = json.Unmarshal(response, &updateMyntraChannelCredentialsResponse)
        if err != nil {
           
             return MkpResp{}, common.NewFDKError(err.Error())
            
        }
        return updateMyntraChannelCredentialsResponse, nil
    }
         
        
       
    
    
  
    
    
    // RegisterAmazonChannel Create Amazon marketplace channel for a seller
     func (ma *PlatformMarketplaces)  RegisterAmazonChannel(
         body  AmazonPayload) (MkpResp, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            registerAmazonChannelResponse MkpResp
            
	    )
     
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return MkpResp{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return MkpResp{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            ma.Config,
            "post",
            fmt.Sprintf("/service/platform/marketplaces/v1.0/company/%s/channels/amazon_in",ma.CompanyID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return MkpResp{}, err
            
	    }
        err = json.Unmarshal(response, &registerAmazonChannelResponse)
        if err != nil {
           
             return MkpResp{}, common.NewFDKError(err.Error())
            
        }
        return registerAmazonChannelResponse, nil
    }
         
        
       
    
    
  
    
    
    // UpdateAmazonChannelCredentials Update Amazon marketplace channel credentials for a seller
     func (ma *PlatformMarketplaces)  UpdateAmazonChannelCredentials(
         body  AmazonPayload) (MkpResp, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            updateAmazonChannelCredentialsResponse MkpResp
            
	    )
     
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return MkpResp{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return MkpResp{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            ma.Config,
            "put",
            fmt.Sprintf("/service/platform/marketplaces/v1.0/company/%s/channels/amazon_in",ma.CompanyID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return MkpResp{}, err
            
	    }
        err = json.Unmarshal(response, &updateAmazonChannelCredentialsResponse)
        if err != nil {
           
             return MkpResp{}, common.NewFDKError(err.Error())
            
        }
        return updateAmazonChannelCredentialsResponse, nil
    }
         
        
       
    
    
  
    
    
    // RegisterFlipkartChannel Create Flipkart / Flipkart Assured marketplace channel for a seller
     func (ma *PlatformMarketplaces)  RegisterFlipkartChannel(
         FlipkartChannel string, body  FlipkartPayload) (MkpResp, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            registerFlipkartChannelResponse MkpResp
            
	    )
     
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return MkpResp{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return MkpResp{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            ma.Config,
            "post",
            fmt.Sprintf("/service/platform/marketplaces/v1.0/company/%s/channels/%s",ma.CompanyID, FlipkartChannel),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return MkpResp{}, err
            
	    }
        err = json.Unmarshal(response, &registerFlipkartChannelResponse)
        if err != nil {
           
             return MkpResp{}, common.NewFDKError(err.Error())
            
        }
        return registerFlipkartChannelResponse, nil
    }
         
        
       
    
    
  
    
    
    // UpdateFlipkartChannelCredentials Update Flipkart / Flipkart Assured marketplace channel credentials for a seller
     func (ma *PlatformMarketplaces)  UpdateFlipkartChannelCredentials(
         FlipkartChannel string, body  FlipkartPayload) (MkpResp, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            updateFlipkartChannelCredentialsResponse MkpResp
            
	    )
     
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return MkpResp{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return MkpResp{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            ma.Config,
            "put",
            fmt.Sprintf("/service/platform/marketplaces/v1.0/company/%s/channels/%s",ma.CompanyID, FlipkartChannel),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return MkpResp{}, err
            
	    }
        err = json.Unmarshal(response, &updateFlipkartChannelCredentialsResponse)
        if err != nil {
           
             return MkpResp{}, common.NewFDKError(err.Error())
            
        }
        return updateFlipkartChannelCredentialsResponse, nil
    }
         
        
       
    
    
  
    
    
    // RegisterTatacliqChannel Create Tatacliq / Tatacliq Luxury marketplace channel for a seller
     func (ma *PlatformMarketplaces)  RegisterTatacliqChannel(
         TatacliqChannel string, body  TatacliqPayload) (MkpResp, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            registerTatacliqChannelResponse MkpResp
            
	    )
     
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return MkpResp{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return MkpResp{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            ma.Config,
            "post",
            fmt.Sprintf("/service/platform/marketplaces/v1.0/company/%s/channels/%s",ma.CompanyID, TatacliqChannel),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return MkpResp{}, err
            
	    }
        err = json.Unmarshal(response, &registerTatacliqChannelResponse)
        if err != nil {
           
             return MkpResp{}, common.NewFDKError(err.Error())
            
        }
        return registerTatacliqChannelResponse, nil
    }
         
        
       
    
    
  
    
    
    // UpdateTatacliqChannelCredentials Update Tatacliq / Tatacliq Luxury Assured marketplace channel credentials for a seller
     func (ma *PlatformMarketplaces)  UpdateTatacliqChannelCredentials(
         TatacliqChannel string, body  TatacliqPayload) (MkpResp, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            updateTatacliqChannelCredentialsResponse MkpResp
            
	    )
     
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return MkpResp{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return MkpResp{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            ma.Config,
            "put",
            fmt.Sprintf("/service/platform/marketplaces/v1.0/company/%s/channels/%s",ma.CompanyID, TatacliqChannel),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return MkpResp{}, err
            
	    }
        err = json.Unmarshal(response, &updateTatacliqChannelCredentialsResponse)
        if err != nil {
           
             return MkpResp{}, common.NewFDKError(err.Error())
            
        }
        return updateTatacliqChannelCredentialsResponse, nil
    }
         
        
       
    
    
  
    
    
    // RegisterAjioChannel Create Ajio marketplace channel for a seller
     func (ma *PlatformMarketplaces)  RegisterAjioChannel(
         body  AjioPayload) (MkpResp, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            registerAjioChannelResponse MkpResp
            
	    )
     
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return MkpResp{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return MkpResp{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            ma.Config,
            "post",
            fmt.Sprintf("/service/platform/marketplaces/v1.0/company/%s/channels/ajio_in",ma.CompanyID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return MkpResp{}, err
            
	    }
        err = json.Unmarshal(response, &registerAjioChannelResponse)
        if err != nil {
           
             return MkpResp{}, common.NewFDKError(err.Error())
            
        }
        return registerAjioChannelResponse, nil
    }
         
        
       
    
    
  
    
    
    // UpdateAjioChannelCredentials Update Ajio marketplace channel credentials for a seller
     func (ma *PlatformMarketplaces)  UpdateAjioChannelCredentials(
         body  AjioPayload) (MkpResp, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            updateAjioChannelCredentialsResponse MkpResp
            
	    )
     
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return MkpResp{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return MkpResp{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            ma.Config,
            "put",
            fmt.Sprintf("/service/platform/marketplaces/v1.0/company/%s/channels/ajio_in",ma.CompanyID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return MkpResp{}, err
            
	    }
        err = json.Unmarshal(response, &updateAjioChannelCredentialsResponse)
        if err != nil {
           
             return MkpResp{}, common.NewFDKError(err.Error())
            
        }
        return updateAjioChannelCredentialsResponse, nil
    }
         
        
       
    
    
  
    
    
    //PlatformUpdateChannelInventorySyncConfigXQuery holds query params
    type PlatformUpdateChannelInventorySyncConfigXQuery struct { 
        ValidateCred string  `url:"validate_cred,omitempty"`  
    }
    
    // UpdateChannelInventorySyncConfig Update inventory sync configuration of marketplace channel for a seller
     func (ma *PlatformMarketplaces)  UpdateChannelInventorySyncConfig(
         Channel string, xQuery PlatformUpdateChannelInventorySyncConfigXQuery, body  InventorySyncConfig) (MkpResp, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            updateChannelInventorySyncConfigResponse MkpResp
            
	    )
     
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return MkpResp{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return MkpResp{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            ma.Config,
            "put",
            fmt.Sprintf("/service/platform/marketplaces/company/%s/v1.0/channels/%s/inventory/config",ma.CompanyID, Channel),
            nil,
            xQuery,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return MkpResp{}, err
            
	    }
        err = json.Unmarshal(response, &updateChannelInventorySyncConfigResponse)
        if err != nil {
           
             return MkpResp{}, common.NewFDKError(err.Error())
            
        }
        return updateChannelInventorySyncConfigResponse, nil
    }
         
        
       
    
    
  
    
    
    // GetChannelLocationConfig Get marketplace channel location config for a seller
     func (ma *PlatformMarketplaces)  GetChannelLocationConfig(
         Channel string) (StoreMapping, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getChannelLocationConfigResponse StoreMapping
            
	    )
     
         
        
        
        //API call
        rawRequest = NewRequest(
            ma.Config,
            "get",
            fmt.Sprintf("/service/platform/marketplaces/v1.0/company/%s/channels/%s/location/config",ma.CompanyID, Channel),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return StoreMapping{}, err
            
	    }
        err = json.Unmarshal(response, &getChannelLocationConfigResponse)
        if err != nil {
           
             return StoreMapping{}, common.NewFDKError(err.Error())
            
        }
        return getChannelLocationConfigResponse, nil
    }
         
        
       
    
    
  
    
    
    // UpdateChannelLocationConfig update marketplace channel location config for a seller
     func (ma *PlatformMarketplaces)  UpdateChannelLocationConfig(
         Channel string, body  StoreMappingPayload) (StoreMapping, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            updateChannelLocationConfigResponse StoreMapping
            
	    )
     
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return StoreMapping{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return StoreMapping{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            ma.Config,
            "put",
            fmt.Sprintf("/service/platform/marketplaces/v1.0/company/%s/channels/%s/location/config",ma.CompanyID, Channel),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return StoreMapping{}, err
            
	    }
        err = json.Unmarshal(response, &updateChannelLocationConfigResponse)
        if err != nil {
           
             return StoreMapping{}, common.NewFDKError(err.Error())
            
        }
        return updateChannelLocationConfigResponse, nil
    }
         
        
       
    
    
  
    
    
    // GetChannelStatus Get marketplace channel active status for a seller
     func (ma *PlatformMarketplaces)  GetChannelStatus(
         Channel string) (StatusPayload, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getChannelStatusResponse StatusPayload
            
	    )
     
         
        
        
        //API call
        rawRequest = NewRequest(
            ma.Config,
            "get",
            fmt.Sprintf("/service/platform/marketplaces/v1.0/company/%s/channels/%s/status",ma.CompanyID, Channel),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return StatusPayload{}, err
            
	    }
        err = json.Unmarshal(response, &getChannelStatusResponse)
        if err != nil {
           
             return StatusPayload{}, common.NewFDKError(err.Error())
            
        }
        return getChannelStatusResponse, nil
    }
         
        
       
    
    
  
    
    
    // UpdateChannelStatus Update marketplace channel active status for a seller
     func (ma *PlatformMarketplaces)  UpdateChannelStatus(
         Channel string, body  StatusPayload) (StatusResp, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            updateChannelStatusResponse StatusResp
            
	    )
     
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return StatusResp{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return StatusResp{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            ma.Config,
            "put",
            fmt.Sprintf("/service/platform/marketplaces/v1.0/company/%s/channels/%s/status",ma.CompanyID, Channel),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return StatusResp{}, err
            
	    }
        err = json.Unmarshal(response, &updateChannelStatusResponse)
        if err != nil {
           
             return StatusResp{}, common.NewFDKError(err.Error())
            
        }
        return updateChannelStatusResponse, nil
    }
         
        
       
    
    
  
    
    
    // TriggerChannelInventoryUpdates Trigger marketplace channel inventory updates for a seller
     func (ma *PlatformMarketplaces)  TriggerChannelInventoryUpdates(
         Channel string, UpdateType string, body  SyncPayload) (SyncResp, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            triggerChannelInventoryUpdatesResponse SyncResp
            
	    )
     
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return SyncResp{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return SyncResp{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            ma.Config,
            "post",
            fmt.Sprintf("/service/platform/marketplaces/v1.0/company/%s/channels/%s/inventory/%s/sync",ma.CompanyID, Channel, UpdateType),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return SyncResp{}, err
            
	    }
        err = json.Unmarshal(response, &triggerChannelInventoryUpdatesResponse)
        if err != nil {
           
             return SyncResp{}, common.NewFDKError(err.Error())
            
        }
        return triggerChannelInventoryUpdatesResponse, nil
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
     func (an *PlatformAnalytics)  CreateExportJob(
         ExportType string, body  ExportJobReq) (ExportJobRes, error){
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
     func (an *PlatformAnalytics)  GetExportJobStatus(
         ExportType string, JobID string) (ExportJobStatusRes, error){
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
     func (an *PlatformAnalytics)  GetLogsList(
         LogType string, xQuery PlatformGetLogsListXQuery, body  GetLogsListReq) (GetLogsListRes, error){
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
     func (an *PlatformAnalytics)  SearchLogs(
         LogType string, xQuery PlatformSearchLogsXQuery, body  SearchLogReq) (SearchLogRes, error){
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
     func (di *PlatformDiscount)  GetDiscounts(
         xQuery PlatformGetDiscountsXQuery) (ListOrCalender, error){
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
     func (di *PlatformDiscount)  CreateDiscount(
         body  CreateUpdateDiscount) (DiscountJob, error){
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
     func (di *PlatformDiscount)  GetDiscount(
         ID string) (DiscountJob, error){
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
     func (di *PlatformDiscount)  UpdateDiscount(
         ID string, body  CreateUpdateDiscount) (DiscountJob, error){
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
     func (di *PlatformDiscount)  ValidateDiscountFile(
         xQuery PlatformValidateDiscountFileXQuery, body  DiscountJob) (FileJobResponse, error){
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
     func (di *PlatformDiscount)  DownloadDiscountFile(
         Type string, body  DownloadFileJob) (FileJobResponse, error){
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
     func (di *PlatformDiscount)  GetValidationJob(
         ID string) (FileJobResponse, error){
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
     func (di *PlatformDiscount)  CancelValidationJob(
         ID string) (CancelJobResponse, error){
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
     func (di *PlatformDiscount)  GetDownloadJob(
         ID string) (FileJobResponse, error){
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
     func (di *PlatformDiscount)  CancelDownloadJob(
         ID string) (CancelJobResponse, error){
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
    
    
  
    
    
    //PlatformGetSubscribersByCompanyAndEventIdXQuery holds query params
    type PlatformGetSubscribersByCompanyAndEventIdXQuery struct { 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"` 
        EventID float64  `url:"event_id,omitempty"`  
    }
    
    // GetSubscribersByCompanyAndEventId Get Subscribers By Company And Event
     func (we *PlatformWebhook)  GetSubscribersByCompanyAndEventId(
         xQuery PlatformGetSubscribersByCompanyAndEventIdXQuery) (SubscriberConfigList, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getSubscribersByCompanyAndEventIdResponse SubscriberConfigList
            
	    )
     
         
        
        
        //API call
        rawRequest = NewRequest(
            we.Config,
            "get",
            fmt.Sprintf("/service/platform/webhook/v1.0/company/%s/subscriber/{event_id}",we.CompanyID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return SubscriberConfigList{}, err
            
	    }
        err = json.Unmarshal(response, &getSubscribersByCompanyAndEventIdResponse)
        if err != nil {
           
             return SubscriberConfigList{}, common.NewFDKError(err.Error())
            
        }
        return getSubscribersByCompanyAndEventIdResponse, nil
    }
         
        
       
    
    
  
    
    
    // RegisterSubscriberToEvent Register Subscriber
     func (we *PlatformWebhook)  RegisterSubscriberToEvent(
         body  SubscriberConfig) (SubscriberConfig, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            registerSubscriberToEventResponse SubscriberConfig
            
	    )
     
         
        
        
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
     func (we *PlatformWebhook)  UpdateSubscriberConfig(
         body  SubscriberConfig) (SubscriberConfig, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            updateSubscriberConfigResponse SubscriberConfig
            
	    )
     
         
        
        
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
         
        
       
    


