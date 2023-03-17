package platform



    
    // ApplicationResponse used by Common
    type ApplicationResponse struct {

        
            Application ApplicationData  `json:"application"`
         
    }
    
    // Currency used by Common
    type Currency struct {

        
            ID string  `json:"_id"`
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            Code string  `json:"code"`
            CreatedAt string  `json:"created_at"`
            UpdatedAt string  `json:"updated_at"`
            DecimalDigits float64  `json:"decimal_digits"`
            Symbol string  `json:"symbol"`
         
    }
    
    // Domain used by Common
    type Domain struct {

        
            Verified bool  `json:"verified"`
            IsPrimary bool  `json:"is_primary"`
            IsShortlink bool  `json:"is_shortlink"`
            ID string  `json:"_id"`
            Name string  `json:"name"`
            IsPredefined bool  `json:"is_predefined"`
         
    }
    
    // ApplicationWebsite used by Common
    type ApplicationWebsite struct {

        
            Enabled bool  `json:"enabled"`
            Basepath string  `json:"basepath"`
         
    }
    
    // ApplicationCors used by Common
    type ApplicationCors struct {

        
            Domains []string  `json:"domains"`
         
    }
    
    // ApplicationAuth used by Common
    type ApplicationAuth struct {

        
            Enabled bool  `json:"enabled"`
         
    }
    
    // ApplicationRedirections used by Common
    type ApplicationRedirections struct {

        
            RedirectFrom string  `json:"redirect_from"`
            RedirectTo string  `json:"redirect_to"`
            Type string  `json:"type"`
         
    }
    
    // ApplicationMeta used by Common
    type ApplicationMeta struct {

        
            Name string  `json:"name"`
            Value string  `json:"value"`
         
    }
    
    // SecureUrl used by Common
    type SecureUrl struct {

        
            SecureURL string  `json:"secure_url"`
         
    }
    
    // ApplicationData used by Common
    type ApplicationData struct {

        
            Website ApplicationWebsite  `json:"website"`
            Cors ApplicationCors  `json:"cors"`
            Auth ApplicationAuth  `json:"auth"`
            Description string  `json:"description"`
            ChannelType string  `json:"channel_type"`
            CacheTtl float64  `json:"cache_ttl"`
            IsInternal bool  `json:"is_internal"`
            IsActive bool  `json:"is_active"`
            ID string  `json:"_id"`
            Name string  `json:"name"`
            Owner string  `json:"owner"`
            CompanyID float64  `json:"company_id"`
            Token string  `json:"token"`
            Redirections []ApplicationRedirections  `json:"redirections"`
            Meta []ApplicationMeta  `json:"meta"`
            CreatedAt string  `json:"created_at"`
            UpdatedAt string  `json:"updated_at"`
            V float64  `json:"__v"`
            Banner SecureUrl  `json:"banner"`
            Logo SecureUrl  `json:"logo"`
            Favicon SecureUrl  `json:"favicon"`
            Domains []Domain  `json:"domains"`
            AppType string  `json:"app_type"`
            MobileLogo SecureUrl  `json:"mobile_logo"`
            Domain Domain  `json:"domain"`
            Slug string  `json:"slug"`
         
    }
    
    // NotFound used by Common
    type NotFound struct {

        
            Message string  `json:"message"`
         
    }
    
    // BadRequest used by Common
    type BadRequest struct {

        
            Message string  `json:"message"`
         
    }
    
    // Page used by Common
    type Page struct {

        
            ItemTotal float64  `json:"item_total"`
            NextID string  `json:"next_id"`
            HasPrevious bool  `json:"has_previous"`
            HasNext bool  `json:"has_next"`
            Current float64  `json:"current"`
            Type string  `json:"type"`
            Size float64  `json:"size"`
         
    }
    
    // LocationDefaultLanguage used by Common
    type LocationDefaultLanguage struct {

        
            Name string  `json:"name"`
            Code string  `json:"code"`
         
    }
    
    // LocationDefaultCurrency used by Common
    type LocationDefaultCurrency struct {

        
            Name string  `json:"name"`
            Symbol string  `json:"symbol"`
            Code string  `json:"code"`
         
    }
    
    // LocationDetails used by Common
    type LocationDetails struct {

        
            Capital string  `json:"capital"`
            Currency string  `json:"currency"`
            Iso2 string  `json:"iso2"`
            Iso3 string  `json:"iso3"`
            Name string  `json:"name"`
            Parent string  `json:"parent"`
            PhoneCode string  `json:"phone_code"`
            Type string  `json:"type"`
            UID float64  `json:"uid"`
            V float64  `json:"__v"`
            ID string  `json:"_id"`
            DefaultCurrency LocationDefaultCurrency  `json:"default_currency"`
            DefaultLanguage LocationDefaultLanguage  `json:"default_language"`
            StateCode string  `json:"state_code"`
            CountryCode string  `json:"country_code"`
            Latitude string  `json:"latitude"`
            Longitude string  `json:"longitude"`
         
    }
    
    // Locations used by Common
    type Locations struct {

        
            Items []LocationDetails  `json:"items"`
         
    }
    

    
    // TicketList used by Lead
    type TicketList struct {

        
            Items []Ticket  `json:"items"`
            Filters Filter  `json:"filters"`
            Page Page  `json:"page"`
         
    }
    
    // Page used by Lead
    type Page struct {

        
            ItemTotal float64  `json:"item_total"`
            NextID string  `json:"next_id"`
            HasPrevious bool  `json:"has_previous"`
            HasNext bool  `json:"has_next"`
            Current float64  `json:"current"`
            Type string  `json:"type"`
            Size float64  `json:"size"`
         
    }
    
    // TicketHistoryList used by Lead
    type TicketHistoryList struct {

        
            Items []TicketHistory  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // CustomFormList used by Lead
    type CustomFormList struct {

        
            Items []CustomForm  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // CreateCustomFormPayload used by Lead
    type CreateCustomFormPayload struct {

        
            Slug string  `json:"slug"`
            Title string  `json:"title"`
            Inputs []map[string]interface{}  `json:"inputs"`
            Description string  `json:"description"`
            HeaderImage string  `json:"header_image"`
            Priority PriorityEnum  `json:"priority"`
            ShouldNotify bool  `json:"should_notify"`
            SuccessMessage string  `json:"success_message"`
            PollForAssignment PollForAssignment  `json:"poll_for_assignment"`
         
    }
    
    // EditCustomFormPayload used by Lead
    type EditCustomFormPayload struct {

        
            Title string  `json:"title"`
            Inputs []map[string]interface{}  `json:"inputs"`
            Description string  `json:"description"`
            Priority PriorityEnum  `json:"priority"`
            HeaderImage string  `json:"header_image"`
            ShouldNotify bool  `json:"should_notify"`
            LoginRequired bool  `json:"login_required"`
            SuccessMessage string  `json:"success_message"`
            PollForAssignment PollForAssignment  `json:"poll_for_assignment"`
         
    }
    
    // EditTicketPayload used by Lead
    type EditTicketPayload struct {

        
            Content TicketContent  `json:"content"`
            Category string  `json:"category"`
            SubCategory string  `json:"sub_category"`
            Source string  `json:"source"`
            Status string  `json:"status"`
            Priority PriorityEnum  `json:"priority"`
            AssignedTo AgentChangePayload  `json:"assigned_to"`
            Tags []string  `json:"tags"`
         
    }
    
    // AgentChangePayload used by Lead
    type AgentChangePayload struct {

        
            AgentID string  `json:"agent_id"`
         
    }
    
    // CreateVideoRoomResponse used by Lead
    type CreateVideoRoomResponse struct {

        
            UniqueName string  `json:"unique_name"`
         
    }
    
    // CloseVideoRoomResponse used by Lead
    type CloseVideoRoomResponse struct {

        
            Success bool  `json:"success"`
         
    }
    
    // CreateVideoRoomPayload used by Lead
    type CreateVideoRoomPayload struct {

        
            UniqueName string  `json:"unique_name"`
            Notify []NotifyUser  `json:"notify"`
         
    }
    
    // NotifyUser used by Lead
    type NotifyUser struct {

        
            CountryCode string  `json:"country_code"`
            PhoneNumber string  `json:"phone_number"`
         
    }
    
    // Filter used by Lead
    type Filter struct {

        
            Priorities []Priority  `json:"priorities"`
            Categories []TicketCategory  `json:"categories"`
            Statuses []Status  `json:"statuses"`
            Assignees []map[string]interface{}  `json:"assignees"`
         
    }
    
    // TicketHistoryPayload used by Lead
    type TicketHistoryPayload struct {

        
            Value map[string]interface{}  `json:"value"`
            Type HistoryTypeEnum  `json:"type"`
         
    }
    
    // CustomFormSubmissionPayload used by Lead
    type CustomFormSubmissionPayload struct {

        
            Response []map[string]interface{}  `json:"response"`
            Attachments []TicketAsset  `json:"attachments"`
         
    }
    
    // GetTokenForVideoRoomResponse used by Lead
    type GetTokenForVideoRoomResponse struct {

        
            AccessToken string  `json:"access_token"`
         
    }
    
    // GetParticipantsInsideVideoRoomResponse used by Lead
    type GetParticipantsInsideVideoRoomResponse struct {

        
            Participants []Participant  `json:"participants"`
         
    }
    
    // Participant used by Lead
    type Participant struct {

        
            User UserSchema  `json:"user"`
            Identity string  `json:"identity"`
            Status string  `json:"status"`
         
    }
    
    // UserSchema used by Lead
    type UserSchema struct {

        
            ApplicationID string  `json:"application_id"`
            UserID string  `json:"user_id"`
            FirstName string  `json:"first_name"`
            Meta map[string]interface{}  `json:"meta"`
            LastName string  `json:"last_name"`
            PhoneNumbers []PhoneNumber  `json:"phone_numbers"`
            Emails []Email  `json:"emails"`
            Gender string  `json:"gender"`
            Dob string  `json:"dob"`
            Active bool  `json:"active"`
            ProfilePicURL string  `json:"profile_pic_url"`
            Username string  `json:"username"`
            AccountType string  `json:"account_type"`
            ID string  `json:"_id"`
            CreatedAt string  `json:"created_at"`
            UpdatedAt string  `json:"updated_at"`
         
    }
    
    // PhoneNumber used by Lead
    type PhoneNumber struct {

        
            Active bool  `json:"active"`
            Primary bool  `json:"primary"`
            Verified bool  `json:"verified"`
            Phone string  `json:"phone"`
            CountryCode float64  `json:"country_code"`
         
    }
    
    // Email used by Lead
    type Email struct {

        
            Primary bool  `json:"primary"`
            Verified bool  `json:"verified"`
            Email string  `json:"email"`
            Active bool  `json:"active"`
         
    }
    
    // Debug used by Lead
    type Debug struct {

        
            Source string  `json:"source"`
            Platform string  `json:"platform"`
         
    }
    
    // SubmitCustomFormResponse used by Lead
    type SubmitCustomFormResponse struct {

        
            Message string  `json:"message"`
            Ticket Ticket  `json:"ticket"`
         
    }
    
    // TicketContext used by Lead
    type TicketContext struct {

        
            ApplicationID string  `json:"application_id"`
            CompanyID string  `json:"company_id"`
         
    }
    
    // CreatedOn used by Lead
    type CreatedOn struct {

        
            UserAgent string  `json:"user_agent"`
         
    }
    
    // TicketAsset used by Lead
    type TicketAsset struct {

        
            Display string  `json:"display"`
            Value string  `json:"value"`
            Type TicketAssetTypeEnum  `json:"type"`
         
    }
    
    // TicketContent used by Lead
    type TicketContent struct {

        
            Title string  `json:"title"`
            Description string  `json:"description"`
            Attachments []TicketAsset  `json:"attachments"`
         
    }
    
    // AddTicketPayload used by Lead
    type AddTicketPayload struct {

        
            CreatedBy map[string]interface{}  `json:"created_by"`
            Status string  `json:"status"`
            Priority PriorityEnum  `json:"priority"`
            Category string  `json:"category"`
            Content TicketContent  `json:"content"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // Priority used by Lead
    type Priority struct {

        
            Key PriorityEnum  `json:"key"`
            Display string  `json:"display"`
            Color string  `json:"color"`
         
    }
    
    // Status used by Lead
    type Status struct {

        
            Key string  `json:"key"`
            Display string  `json:"display"`
            Color string  `json:"color"`
         
    }
    
    // TicketFeedbackForm used by Lead
    type TicketFeedbackForm struct {

        
            Title string  `json:"title"`
            Display []map[string]interface{}  `json:"display"`
         
    }
    
    // TicketFeedbackList used by Lead
    type TicketFeedbackList struct {

        
            Items []TicketFeedback  `json:"items"`
         
    }
    
    // TicketFeedbackPayload used by Lead
    type TicketFeedbackPayload struct {

        
            FormResponse map[string]interface{}  `json:"form_response"`
         
    }
    
    // SubmitButton used by Lead
    type SubmitButton struct {

        
            Title string  `json:"title"`
            TitleColor string  `json:"title_color"`
            BackgroundColor string  `json:"background_color"`
         
    }
    
    // PollForAssignment used by Lead
    type PollForAssignment struct {

        
            Duration float64  `json:"duration"`
            Message string  `json:"message"`
            SuccessMessage string  `json:"success_message"`
            FailureMessage string  `json:"failure_message"`
         
    }
    
    // CustomForm used by Lead
    type CustomForm struct {

        
            ApplicationID string  `json:"application_id"`
            Slug string  `json:"slug"`
            HeaderImage string  `json:"header_image"`
            Title string  `json:"title"`
            Description string  `json:"description"`
            Priority Priority  `json:"priority"`
            LoginRequired bool  `json:"login_required"`
            ShouldNotify bool  `json:"should_notify"`
            SuccessMessage string  `json:"success_message"`
            SubmitButton SubmitButton  `json:"submit_button"`
            Inputs []map[string]interface{}  `json:"inputs"`
            CreatedOn CreatedOn  `json:"created_on"`
            PollForAssignment PollForAssignment  `json:"poll_for_assignment"`
            ID string  `json:"_id"`
         
    }
    
    // CommunicationDetails used by Lead
    type CommunicationDetails struct {

        
            Type string  `json:"type"`
            Title string  `json:"title"`
            Value string  `json:"value"`
            Description string  `json:"description"`
            Enabled bool  `json:"enabled"`
         
    }
    
    // SupportGeneralConfig used by Lead
    type SupportGeneralConfig struct {

        
            ID string  `json:"_id"`
            SupportEmail CommunicationDetails  `json:"support_email"`
            SupportPhone CommunicationDetails  `json:"support_phone"`
            SupportFaq CommunicationDetails  `json:"support_faq"`
            ShowCommunicationInfo bool  `json:"show_communication_info"`
            SupportCommunication CommunicationDetails  `json:"support_communication"`
            ShowSupportDris bool  `json:"show_support_dris"`
            Integration map[string]interface{}  `json:"integration"`
         
    }
    
    // FeedbackForm used by Lead
    type FeedbackForm struct {

        
            Inputs map[string]interface{}  `json:"inputs"`
            Title string  `json:"title"`
            Timestamps map[string]interface{}  `json:"timestamps"`
         
    }
    
    // TicketSubCategory used by Lead
    type TicketSubCategory struct {

        
            Key string  `json:"key"`
            Display string  `json:"display"`
            SubCategories *TicketSubCategory  `json:"sub_categories"`
         
    }
    
    // TicketCategory used by Lead
    type TicketCategory struct {

        
            Display string  `json:"display"`
            Key string  `json:"key"`
            SubCategories *TicketCategory  `json:"sub_categories"`
            GroupID float64  `json:"group_id"`
            FeedbackForm FeedbackForm  `json:"feedback_form"`
         
    }
    
    // CategoryData used by Lead
    type CategoryData struct {

        
            List TicketCategory  `json:"list"`
         
    }
    
    // IntegrationConfig used by Lead
    type IntegrationConfig struct {

        
            ID string  `json:"_id"`
            IntegrationType string  `json:"integration_type"`
            BaseURL string  `json:"base_url"`
            CreateTicketApikey string  `json:"create_ticket_apikey"`
            UpdateTicketApikey string  `json:"update_ticket_apikey"`
            CategorySyncApikey string  `json:"category_sync_apikey"`
            CategoryData CategoryData  `json:"category_data"`
            WebhookApikey string  `json:"webhook_apikey"`
            ConfigCompleted bool  `json:"config_completed"`
            AllowTicketCreation bool  `json:"allow_ticket_creation"`
            ShowListing bool  `json:"show_listing"`
         
    }
    
    // FeedbackResponseItem used by Lead
    type FeedbackResponseItem struct {

        
            Display string  `json:"display"`
            Key string  `json:"key"`
            Value string  `json:"value"`
         
    }
    
    // TicketFeedback used by Lead
    type TicketFeedback struct {

        
            ID string  `json:"_id"`
            TicketID string  `json:"ticket_id"`
            CompanyID string  `json:"company_id"`
            Response []FeedbackResponseItem  `json:"response"`
            Category string  `json:"category"`
            User map[string]interface{}  `json:"user"`
            UpdatedAt string  `json:"updated_at"`
            CreatedAt string  `json:"created_at"`
         
    }
    
    // TicketHistory used by Lead
    type TicketHistory struct {

        
            Type string  `json:"type"`
            Value map[string]interface{}  `json:"value"`
            TicketID string  `json:"ticket_id"`
            CreatedOn CreatedOn  `json:"created_on"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ID string  `json:"_id"`
            UpdatedAt string  `json:"updated_at"`
            CreatedAt string  `json:"created_at"`
         
    }
    
    // Ticket used by Lead
    type Ticket struct {

        
            Context TicketContext  `json:"context"`
            CreatedOn CreatedOn  `json:"created_on"`
            ResponseID string  `json:"response_id"`
            Content TicketContent  `json:"content"`
            Category TicketCategory  `json:"category"`
            SubCategory string  `json:"sub_category"`
            Source TicketSourceEnum  `json:"source"`
            Status Status  `json:"status"`
            Priority Priority  `json:"priority"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            AssignedTo map[string]interface{}  `json:"assigned_to"`
            Tags []string  `json:"tags"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            IsFeedbackPending bool  `json:"is_feedback_pending"`
            Integration map[string]interface{}  `json:"integration"`
            ID string  `json:"_id"`
            UpdatedAt string  `json:"updated_at"`
            CreatedAt string  `json:"created_at"`
         
    }
    

    
    // Activity used by Feedback
    type Activity struct {

        
            CurrentState map[string]interface{}  `json:"current_state"`
            DocumentID string  `json:"document_id"`
            PreviousState map[string]interface{}  `json:"previous_state"`
         
    }
    
    // ActivityDump used by Feedback
    type ActivityDump struct {

        
            Activity Activity  `json:"activity"`
            CreatedBy CreatedBy  `json:"created_by"`
            DateMeta DateMeta  `json:"date_meta"`
            ID string  `json:"id"`
            Type string  `json:"type"`
         
    }
    
    // AddMediaListRequest used by Feedback
    type AddMediaListRequest struct {

        
            EntityID string  `json:"entity_id"`
            EntityType string  `json:"entity_type"`
            MediaList []AddMediaRequest  `json:"media_list"`
            RefID string  `json:"ref_id"`
            RefType string  `json:"ref_type"`
         
    }
    
    // AddMediaRequest used by Feedback
    type AddMediaRequest struct {

        
            CloudID string  `json:"cloud_id"`
            CloudName string  `json:"cloud_name"`
            CloudProvider string  `json:"cloud_provider"`
            EntityID string  `json:"entity_id"`
            EntityType string  `json:"entity_type"`
            MediaURL string  `json:"media_url"`
            RefID string  `json:"ref_id"`
            RefType string  `json:"ref_type"`
            Tags []string  `json:"tags"`
            ThumbnailURL string  `json:"thumbnail_url"`
            Type string  `json:"type"`
         
    }
    
    // ApproveRequest used by Feedback
    type ApproveRequest struct {

        
            Approve bool  `json:"approve"`
            EntityType string  `json:"entity_type"`
            ID string  `json:"id"`
            Reason string  `json:"reason"`
         
    }
    
    // Attribute used by Feedback
    type Attribute struct {

        
            DateMeta DateMeta  `json:"date_meta"`
            Description string  `json:"description"`
            ID string  `json:"id"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Tags []TagMeta  `json:"tags"`
         
    }
    
    // AttributeObject used by Feedback
    type AttributeObject struct {

        
            Description string  `json:"description"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Title string  `json:"title"`
            Type string  `json:"type"`
            Value float64  `json:"value"`
         
    }
    
    // CreatedBy used by Feedback
    type CreatedBy struct {

        
            ID string  `json:"id"`
            Name string  `json:"name"`
            Tags []TagMeta  `json:"tags"`
         
    }
    
    // CursorGetResponse used by Feedback
    type CursorGetResponse struct {

        
            Items []map[string]interface{}  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // DateMeta used by Feedback
    type DateMeta struct {

        
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // DeviceMeta used by Feedback
    type DeviceMeta struct {

        
            AppVersion string  `json:"app_version"`
            Platform string  `json:"platform"`
         
    }
    
    // Entity used by Feedback
    type Entity struct {

        
            ID string  `json:"id"`
            Type string  `json:"type"`
         
    }
    
    // EntityRequest used by Feedback
    type EntityRequest struct {

        
            EntityID string  `json:"entity_id"`
            EntityType string  `json:"entity_type"`
         
    }
    
    // FeedbackAttributes used by Feedback
    type FeedbackAttributes struct {

        
            Items []Attribute  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // FeedbackError used by Feedback
    type FeedbackError struct {

        
            Code map[string]interface{}  `json:"code"`
            Exception string  `json:"exception"`
            Info string  `json:"info"`
            Message string  `json:"message"`
            Meta map[string]interface{}  `json:"meta"`
            RequestID string  `json:"request_id"`
            StackTrace string  `json:"stack_trace"`
            Status float64  `json:"status"`
         
    }
    
    // FeedbackState used by Feedback
    type FeedbackState struct {

        
            Active bool  `json:"active"`
            Archive bool  `json:"archive"`
            Media string  `json:"media"`
            Qna bool  `json:"qna"`
            Rating bool  `json:"rating"`
            Review bool  `json:"review"`
         
    }
    
    // GetResponse used by Feedback
    type GetResponse struct {

        
            Data map[string]interface{}  `json:"data"`
            Page Page  `json:"page"`
         
    }
    
    // GetReviewResponse used by Feedback
    type GetReviewResponse struct {

        
            Facets []ReviewFacet  `json:"facets"`
            Items []map[string]interface{}  `json:"items"`
            Page Page  `json:"page"`
            Sort []SortMethod  `json:"sort"`
         
    }
    
    // InsertResponse used by Feedback
    type InsertResponse struct {

        
            Count float64  `json:"count"`
         
    }
    
    // MediaMeta used by Feedback
    type MediaMeta struct {

        
            MaxCount float64  `json:"max_count"`
            Size float64  `json:"size"`
            Type string  `json:"type"`
         
    }
    
    // MediaMetaRequest used by Feedback
    type MediaMetaRequest struct {

        
            MaxCount float64  `json:"max_count"`
            Size float64  `json:"size"`
         
    }
    
    // NumberGetResponse used by Feedback
    type NumberGetResponse struct {

        
            Items []map[string]interface{}  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // Page used by Feedback
    type Page struct {

        
            ItemTotal float64  `json:"item_total"`
            NextID string  `json:"next_id"`
            HasPrevious bool  `json:"has_previous"`
            HasNext bool  `json:"has_next"`
            Current float64  `json:"current"`
            Type string  `json:"type"`
            Size float64  `json:"size"`
         
    }
    
    // PageCursor used by Feedback
    type PageCursor struct {

        
            Current float64  `json:"current"`
            HasNext bool  `json:"has_next"`
            HasPrevious bool  `json:"has_previous"`
            ItemTotal float64  `json:"item_total"`
            NextID string  `json:"next_id"`
            Size float64  `json:"size"`
            Type string  `json:"type"`
         
    }
    
    // PageNumber used by Feedback
    type PageNumber struct {

        
            Current float64  `json:"current"`
            HasNext bool  `json:"has_next"`
            ItemTotal float64  `json:"item_total"`
            Size float64  `json:"size"`
            Type string  `json:"type"`
         
    }
    
    // Rating used by Feedback
    type Rating struct {

        
            Attributes []Attribute  `json:"attributes"`
            AttributesSlugs []string  `json:"attributes_slugs"`
            UI UI  `json:"ui"`
         
    }
    
    // RatingRequest used by Feedback
    type RatingRequest struct {

        
            Attributes []string  `json:"attributes"`
            UI UI  `json:"ui"`
         
    }
    
    // ReportAbuseRequest used by Feedback
    type ReportAbuseRequest struct {

        
            Description string  `json:"description"`
            EntityID string  `json:"entity_id"`
            EntityType string  `json:"entity_type"`
         
    }
    
    // Review used by Feedback
    type Review struct {

        
            Description string  `json:"description"`
            Header string  `json:"header"`
            ImageMeta MediaMeta  `json:"image_meta"`
            Title string  `json:"title"`
            VideoMeta MediaMeta  `json:"video_meta"`
            VoteAllowed bool  `json:"vote_allowed"`
         
    }
    
    // ReviewFacet used by Feedback
    type ReviewFacet struct {

        
            Display string  `json:"display"`
            Name string  `json:"name"`
            Selected bool  `json:"selected"`
            Slug string  `json:"slug"`
            Type string  `json:"type"`
         
    }
    
    // ReviewRequest used by Feedback
    type ReviewRequest struct {

        
            Description string  `json:"description"`
            Header string  `json:"header"`
            ImageMeta MediaMetaRequest  `json:"image_meta"`
            IsVoteAllowed bool  `json:"is_vote_allowed"`
            Title string  `json:"title"`
            VideoMeta MediaMetaRequest  `json:"video_meta"`
         
    }
    
    // SaveAttributeRequest used by Feedback
    type SaveAttributeRequest struct {

        
            Description string  `json:"description"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
         
    }
    
    // SortMethod used by Feedback
    type SortMethod struct {

        
            Name string  `json:"name"`
            Selected bool  `json:"selected"`
            Type string  `json:"type"`
         
    }
    
    // TagMeta used by Feedback
    type TagMeta struct {

        
            Media []MediaMeta  `json:"media"`
            Name string  `json:"name"`
            Type string  `json:"type"`
         
    }
    
    // Template used by Feedback
    type Template struct {

        
            DateMeta DateMeta  `json:"date_meta"`
            Entity Entity  `json:"entity"`
            ID string  `json:"id"`
            Name string  `json:"name"`
            Rating Rating  `json:"rating"`
            Review Review  `json:"review"`
            State FeedbackState  `json:"state"`
            Tags []TagMeta  `json:"tags"`
         
    }
    
    // TemplateGetResponse used by Feedback
    type TemplateGetResponse struct {

        
            Items []Template  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // TemplateRequest used by Feedback
    type TemplateRequest struct {

        
            Active bool  `json:"active"`
            EnableMediaType string  `json:"enable_media_type"`
            EnableQna bool  `json:"enable_qna"`
            EnableRating bool  `json:"enable_rating"`
            EnableReview bool  `json:"enable_review"`
            Entity EntityRequest  `json:"entity"`
            Rating RatingRequest  `json:"rating"`
            Review ReviewRequest  `json:"review"`
         
    }
    
    // TemplateRequestList used by Feedback
    type TemplateRequestList struct {

        
            TemplateList []TemplateRequest  `json:"template_list"`
         
    }
    
    // UI used by Feedback
    type UI struct {

        
            FeedbackQuestion []string  `json:"feedback_question"`
            Icon UIIcon  `json:"icon"`
            Text []string  `json:"text"`
            Type string  `json:"type"`
         
    }
    
    // UIIcon used by Feedback
    type UIIcon struct {

        
            Active string  `json:"active"`
            Inactive string  `json:"inactive"`
            Selected []string  `json:"selected"`
         
    }
    
    // UpdateAttributeRequest used by Feedback
    type UpdateAttributeRequest struct {

        
            Description string  `json:"description"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
         
    }
    
    // UpdateResponse used by Feedback
    type UpdateResponse struct {

        
            Count float64  `json:"count"`
         
    }
    
    // UpdateReviewRequest used by Feedback
    type UpdateReviewRequest struct {

        
            Active bool  `json:"active"`
            Application string  `json:"application"`
            Approve bool  `json:"approve"`
            Archive bool  `json:"archive"`
            AttributesRating []AttributeObject  `json:"attributes_rating"`
            Description string  `json:"description"`
            DeviceMeta DeviceMeta  `json:"device_meta"`
            EntityID string  `json:"entity_id"`
            EntityType string  `json:"entity_type"`
            MediaResource []MediaMeta  `json:"media_resource"`
            Rating float64  `json:"rating"`
            ReviewID string  `json:"review_id"`
            TemplateID string  `json:"template_id"`
            Title string  `json:"title"`
         
    }
    
    // UpdateTemplateRequest used by Feedback
    type UpdateTemplateRequest struct {

        
            Active bool  `json:"active"`
            EnableMediaType string  `json:"enable_media_type"`
            EnableQna bool  `json:"enable_qna"`
            EnableRating bool  `json:"enable_rating"`
            EnableReview bool  `json:"enable_review"`
            Entity EntityRequest  `json:"entity"`
            Rating RatingRequest  `json:"rating"`
            Review ReviewRequest  `json:"review"`
         
    }
    
    // UpdateTemplateStatusRequest used by Feedback
    type UpdateTemplateStatusRequest struct {

        
            Active bool  `json:"active"`
            Archive bool  `json:"archive"`
         
    }
    

    
    // AvailablePageSchema used by Theme
    type AvailablePageSchema struct {

        
            Value string  `json:"value"`
            Text string  `json:"text"`
            Path string  `json:"path"`
            Type string  `json:"type"`
            Sections []AvailablePageSchemaSections  `json:"sections"`
            SectionsMeta []AvailablePageSectionMetaAttributes  `json:"sections_meta"`
            Theme string  `json:"theme"`
            Seo AvailablePageSeo  `json:"seo"`
            Props []map[string]interface{}  `json:"props"`
            ID string  `json:"_id"`
         
    }
    
    // AvailablePageSectionMetaAttributes used by Theme
    type AvailablePageSectionMetaAttributes struct {

        
            Attributes map[string]interface{}  `json:"attributes"`
         
    }
    
    // AvailablePageSeo used by Theme
    type AvailablePageSeo struct {

        
            Title string  `json:"title"`
            Description string  `json:"description"`
            ID string  `json:"_id"`
         
    }
    
    // AvailablePageSchemaSections used by Theme
    type AvailablePageSchemaSections struct {

        
            Name string  `json:"name"`
            Label string  `json:"label"`
            Props map[string]interface{}  `json:"props"`
            Blocks []map[string]interface{}  `json:"blocks"`
            Preset map[string]interface{}  `json:"preset"`
            Predicate AvailablePagePredicate  `json:"predicate"`
         
    }
    
    // AvailablePageScreenPredicate used by Theme
    type AvailablePageScreenPredicate struct {

        
            Mobile bool  `json:"mobile"`
            Desktop bool  `json:"desktop"`
            Tablet bool  `json:"tablet"`
         
    }
    
    // AvailablePageUserPredicate used by Theme
    type AvailablePageUserPredicate struct {

        
            Authenticated bool  `json:"authenticated"`
            Anonymous bool  `json:"anonymous"`
         
    }
    
    // AvailablePageRoutePredicate used by Theme
    type AvailablePageRoutePredicate struct {

        
            Selected string  `json:"selected"`
            ExactURL string  `json:"exact_url"`
            Query map[string]interface{}  `json:"query"`
         
    }
    
    // AvailablePagePredicate used by Theme
    type AvailablePagePredicate struct {

        
            Screen AvailablePageScreenPredicate  `json:"screen"`
            User AvailablePageUserPredicate  `json:"user"`
            Route AvailablePageRoutePredicate  `json:"route"`
         
    }
    
    // AllAvailablePageSchema used by Theme
    type AllAvailablePageSchema struct {

        
            Pages []AvailablePageSchema  `json:"pages"`
         
    }
    
    // PaginationSchema used by Theme
    type PaginationSchema struct {

        
            Size float64  `json:"size"`
            ItemTotal float64  `json:"item_total"`
            HasNext bool  `json:"has_next"`
            Type string  `json:"type"`
            Current float64  `json:"current"`
         
    }
    
    // ThemesListingResponseSchema used by Theme
    type ThemesListingResponseSchema struct {

        
            Items []ThemesSchema  `json:"items"`
            Page PaginationSchema  `json:"page"`
         
    }
    
    // AddThemeRequestSchema used by Theme
    type AddThemeRequestSchema struct {

        
            ThemeID string  `json:"theme_id"`
         
    }
    
    // UpgradableThemeSchema used by Theme
    type UpgradableThemeSchema struct {

        
            ParentTheme string  `json:"parent_theme"`
            AppliedTheme string  `json:"applied_theme"`
            Upgrade bool  `json:"upgrade"`
         
    }
    
    // FontsSchema used by Theme
    type FontsSchema struct {

        
            Items FontsSchemaItems  `json:"items"`
            Kind string  `json:"kind"`
         
    }
    
    // BlitzkriegApiErrorSchema used by Theme
    type BlitzkriegApiErrorSchema struct {

        
            Message string  `json:"message"`
         
    }
    
    // BlitzkriegNotFoundSchema used by Theme
    type BlitzkriegNotFoundSchema struct {

        
            Message string  `json:"message"`
         
    }
    
    // BlitzkriegInternalServerErrorSchema used by Theme
    type BlitzkriegInternalServerErrorSchema struct {

        
            Message string  `json:"message"`
         
    }
    
    // FontsSchemaItems used by Theme
    type FontsSchemaItems struct {

        
            Family string  `json:"family"`
            Variants []string  `json:"variants"`
            Subsets []string  `json:"subsets"`
            Version string  `json:"version"`
            LastModified string  `json:"last_modified"`
            Files FontsSchemaItemsFiles  `json:"files"`
            Category string  `json:"category"`
            Kind string  `json:"kind"`
         
    }
    
    // FontsSchemaItemsFiles used by Theme
    type FontsSchemaItemsFiles struct {

        
            Regular string  `json:"regular"`
            Italic string  `json:"italic"`
            Bold string  `json:"bold"`
         
    }
    
    // ThemesSchema used by Theme
    type ThemesSchema struct {

        
            Application string  `json:"application"`
            Applied bool  `json:"applied"`
            Customized bool  `json:"customized"`
            Published bool  `json:"published"`
            Archived bool  `json:"archived"`
            CreatedAt string  `json:"created_at"`
            UpdatedAt string  `json:"updated_at"`
            Version string  `json:"version"`
            ParentThemeVersion string  `json:"parent_theme_version"`
            ParentTheme string  `json:"parent_theme"`
            Information Information  `json:"information"`
            Tags []string  `json:"tags"`
            Src Src  `json:"src"`
            Assets AssetsSchema  `json:"assets"`
            AvailableSections []availableSectionSchema  `json:"available_sections"`
            Styles map[string]interface{}  `json:"styles"`
            Config Config  `json:"config"`
            Font Font  `json:"font"`
            ID string  `json:"_id"`
            V float64  `json:"__v"`
            Colors Colors  `json:"colors"`
         
    }
    
    // availableSectionSchema used by Theme
    type availableSectionSchema struct {

        
            Blocks []Blocks  `json:"blocks"`
            Name string  `json:"name"`
            Label string  `json:"label"`
            Props []BlocksProps  `json:"props"`
         
    }
    
    // Information used by Theme
    type Information struct {

        
            Images Images  `json:"images"`
            Features []string  `json:"features"`
            Name string  `json:"name"`
            Description string  `json:"description"`
         
    }
    
    // Images used by Theme
    type Images struct {

        
            Desktop []string  `json:"desktop"`
            Android []string  `json:"android"`
            Ios []string  `json:"ios"`
            Thumbnail []string  `json:"thumbnail"`
         
    }
    
    // Src used by Theme
    type Src struct {

        
            Link string  `json:"link"`
         
    }
    
    // AssetsSchema used by Theme
    type AssetsSchema struct {

        
            UmdJs UmdJs  `json:"umd_js"`
            CommonJs CommonJs  `json:"common_js"`
            Css Css  `json:"css"`
         
    }
    
    // UmdJs used by Theme
    type UmdJs struct {

        
            Link string  `json:"link"`
            Links []string  `json:"links"`
         
    }
    
    // CommonJs used by Theme
    type CommonJs struct {

        
            Link string  `json:"link"`
         
    }
    
    // Css used by Theme
    type Css struct {

        
            Link string  `json:"link"`
            Links []string  `json:"links"`
         
    }
    
    // Sections used by Theme
    type Sections struct {

        
            Attributes string  `json:"attributes"`
         
    }
    
    // Config used by Theme
    type Config struct {

        
            Preset Preset  `json:"preset"`
            GlobalSchema GlobalSchema  `json:"global_schema"`
            Current string  `json:"current"`
            List []ListSchemaItem  `json:"list"`
         
    }
    
    // Preset used by Theme
    type Preset struct {

        
            Pages []AvailablePageSchema  `json:"pages"`
         
    }
    
    // GlobalSchema used by Theme
    type GlobalSchema struct {

        
            Props []GlobalSchemaProps  `json:"props"`
         
    }
    
    // ListSchemaItem used by Theme
    type ListSchemaItem struct {

        
            GlobalConfig map[string]interface{}  `json:"global_config"`
            Page []ConfigPage  `json:"page"`
            Name string  `json:"name"`
         
    }
    
    // Colors used by Theme
    type Colors struct {

        
            BgColor string  `json:"bg_color"`
            PrimaryColor string  `json:"primary_color"`
            SecondaryColor string  `json:"secondary_color"`
            AccentColor string  `json:"accent_color"`
            LinkColor string  `json:"link_color"`
            ButtonSecondaryColor string  `json:"button_secondary_color"`
         
    }
    
    // Custom used by Theme
    type Custom struct {

        
            Props map[string]interface{}  `json:"props"`
         
    }
    
    // ConfigPage used by Theme
    type ConfigPage struct {

        
            Settings map[string]interface{}  `json:"settings"`
            Page string  `json:"page"`
         
    }
    
    // Font used by Theme
    type Font struct {

        
            Family string  `json:"family"`
            Variants Variants  `json:"variants"`
         
    }
    
    // Variants used by Theme
    type Variants struct {

        
            Medium Medium  `json:"medium"`
            SemiBold SemiBold  `json:"semi_bold"`
            Bold Bold  `json:"bold"`
            Light Light  `json:"light"`
            Regular Regular  `json:"regular"`
         
    }
    
    // Medium used by Theme
    type Medium struct {

        
            Name string  `json:"name"`
            File string  `json:"file"`
         
    }
    
    // SemiBold used by Theme
    type SemiBold struct {

        
            Name string  `json:"name"`
            File string  `json:"file"`
         
    }
    
    // Bold used by Theme
    type Bold struct {

        
            Name string  `json:"name"`
            File string  `json:"file"`
         
    }
    
    // Light used by Theme
    type Light struct {

        
            Name string  `json:"name"`
            File string  `json:"file"`
         
    }
    
    // Regular used by Theme
    type Regular struct {

        
            Name string  `json:"name"`
            File string  `json:"file"`
         
    }
    
    // Blocks used by Theme
    type Blocks struct {

        
            Type string  `json:"type"`
            Name string  `json:"name"`
            Props []BlocksProps  `json:"props"`
         
    }
    
    // GlobalSchemaProps used by Theme
    type GlobalSchemaProps struct {

        
            ID string  `json:"id"`
            Label string  `json:"label"`
            Type string  `json:"type"`
            Category string  `json:"category"`
         
    }
    
    // BlocksProps used by Theme
    type BlocksProps struct {

        
            ID string  `json:"id"`
            Label string  `json:"label"`
            Type string  `json:"type"`
         
    }
    

    
    // BlockUserRequestSchema used by User
    type BlockUserRequestSchema struct {

        
            Status bool  `json:"status"`
            UserID []string  `json:"user_id"`
            Reason string  `json:"reason"`
         
    }
    
    // ArchiveUserRequestSchema used by User
    type ArchiveUserRequestSchema struct {

        
            UserID string  `json:"user_id"`
         
    }
    
    // DeleteApplicationUserRequestSchema used by User
    type DeleteApplicationUserRequestSchema struct {

        
            UserID string  `json:"user_id"`
            Reason string  `json:"reason"`
            ReasonID string  `json:"reason_id"`
            RequestID string  `json:"request_id"`
            Otp string  `json:"otp"`
         
    }
    
    // UnDeleteUserRequestSchema used by User
    type UnDeleteUserRequestSchema struct {

        
            UserID string  `json:"user_id"`
            Reason string  `json:"reason"`
            ReasonID string  `json:"reason_id"`
         
    }
    
    // EditEmailRequestSchema used by User
    type EditEmailRequestSchema struct {

        
            Email string  `json:"email"`
         
    }
    
    // SendVerificationLinkMobileRequestSchema used by User
    type SendVerificationLinkMobileRequestSchema struct {

        
            Verified bool  `json:"verified"`
            Active bool  `json:"active"`
            CountryCode string  `json:"country_code"`
            Phone string  `json:"phone"`
            Primary bool  `json:"primary"`
         
    }
    
    // EditMobileRequestSchema used by User
    type EditMobileRequestSchema struct {

        
            CountryCode string  `json:"country_code"`
            Phone string  `json:"phone"`
         
    }
    
    // EditProfileRequestSchema used by User
    type EditProfileRequestSchema struct {

        
            FirstName string  `json:"first_name"`
            LastName string  `json:"last_name"`
            Mobile EditProfileMobileSchema  `json:"mobile"`
            CountryCode string  `json:"country_code"`
            Email string  `json:"email"`
            Gender string  `json:"gender"`
            Dob string  `json:"dob"`
            ProfilePicURL string  `json:"profile_pic_url"`
            AndroidHash string  `json:"android_hash"`
            Sender string  `json:"sender"`
            RegisterToken string  `json:"register_token"`
         
    }
    
    // EditProfileMobileSchema used by User
    type EditProfileMobileSchema struct {

        
            Phone string  `json:"phone"`
            CountryCode string  `json:"country_code"`
         
    }
    
    // SendEmailOtpRequestSchema used by User
    type SendEmailOtpRequestSchema struct {

        
            Email string  `json:"email"`
            Action string  `json:"action"`
            Token string  `json:"token"`
            RegisterToken string  `json:"register_token"`
         
    }
    
    // VerifyEmailOtpRequestSchema used by User
    type VerifyEmailOtpRequestSchema struct {

        
            Email string  `json:"email"`
            Action string  `json:"action"`
            RegisterToken string  `json:"register_token"`
            Otp string  `json:"otp"`
         
    }
    
    // VerifyOtpRequestSchema used by User
    type VerifyOtpRequestSchema struct {

        
            RequestID string  `json:"request_id"`
            RegisterToken string  `json:"register_token"`
            Otp string  `json:"otp"`
         
    }
    
    // SendMobileOtpRequestSchema used by User
    type SendMobileOtpRequestSchema struct {

        
            Mobile string  `json:"mobile"`
            CountryCode string  `json:"country_code"`
            Action string  `json:"action"`
            Token string  `json:"token"`
            AndroidHash string  `json:"android_hash"`
            Force string  `json:"force"`
            CaptchaCode string  `json:"captcha_code"`
         
    }
    
    // UpdatePasswordRequestSchema used by User
    type UpdatePasswordRequestSchema struct {

        
            OldPassword string  `json:"old_password"`
            NewPassword string  `json:"new_password"`
         
    }
    
    // FormRegisterRequestSchema used by User
    type FormRegisterRequestSchema struct {

        
            FirstName string  `json:"first_name"`
            LastName string  `json:"last_name"`
            Gender string  `json:"gender"`
            Email string  `json:"email"`
            Password string  `json:"password"`
            Phone FormRegisterRequestSchemaPhone  `json:"phone"`
            RegisterToken string  `json:"register_token"`
         
    }
    
    // TokenRequestBodySchema used by User
    type TokenRequestBodySchema struct {

        
            Token string  `json:"token"`
         
    }
    
    // ForgotPasswordRequestSchema used by User
    type ForgotPasswordRequestSchema struct {

        
            Code string  `json:"code"`
            Password string  `json:"password"`
         
    }
    
    // CodeRequestBodySchema used by User
    type CodeRequestBodySchema struct {

        
            Code string  `json:"code"`
         
    }
    
    // SendResetPasswordEmailRequestSchema used by User
    type SendResetPasswordEmailRequestSchema struct {

        
            Email string  `json:"email"`
            CaptchaCode string  `json:"captcha_code"`
         
    }
    
    // SendResetPasswordMobileRequestSchema used by User
    type SendResetPasswordMobileRequestSchema struct {

        
            CountryCode string  `json:"country_code"`
            Mobile string  `json:"mobile"`
            CaptchaCode string  `json:"captcha_code"`
         
    }
    
    // PasswordLoginRequestSchema used by User
    type PasswordLoginRequestSchema struct {

        
            CaptchaCode string  `json:"captcha_code"`
            Password string  `json:"password"`
            Username string  `json:"username"`
         
    }
    
    // SendOtpRequestSchema used by User
    type SendOtpRequestSchema struct {

        
            CountryCode string  `json:"country_code"`
            CaptchaCode string  `json:"captcha_code"`
            Mobile string  `json:"mobile"`
            AndroidHash string  `json:"android_hash"`
         
    }
    
    // OAuthRequestSchema used by User
    type OAuthRequestSchema struct {

        
            IsSignedIn bool  `json:"is_signed_in"`
            Oauth2 OAuthRequestSchemaOauth2  `json:"oauth2"`
            Profile OAuthRequestSchemaProfile  `json:"profile"`
         
    }
    
    // OAuthRequestAppleSchema used by User
    type OAuthRequestAppleSchema struct {

        
            UserIdentifier string  `json:"user_identifier"`
            Oauth OAuthRequestAppleSchemaOauth  `json:"oauth"`
            Profile OAuthRequestAppleSchemaProfile  `json:"profile"`
         
    }
    
    // UserObjectSchema used by User
    type UserObjectSchema struct {

        
            User UserSchema  `json:"user"`
         
    }
    
    // AuthSuccess used by User
    type AuthSuccess struct {

        
            RegisterToken string  `json:"register_token"`
            UserExists bool  `json:"user_exists"`
            User UserSchema  `json:"user"`
         
    }
    
    // SendOtpResponse used by User
    type SendOtpResponse struct {

        
            ResendTimer float64  `json:"resend_timer"`
            ResendToken string  `json:"resend_token"`
            Success bool  `json:"success"`
            RequestID string  `json:"request_id"`
            Message string  `json:"message"`
            Mobile string  `json:"mobile"`
            CountryCode string  `json:"country_code"`
            Email string  `json:"email"`
            ResendEmailToken string  `json:"resend_email_token"`
            RegisterToken string  `json:"register_token"`
            VerifyEmailOtp bool  `json:"verify_email_otp"`
            VerifyMobileOtp bool  `json:"verify_mobile_otp"`
            UserExists bool  `json:"user_exists"`
         
    }
    
    // ProfileEditSuccess used by User
    type ProfileEditSuccess struct {

        
            User UserSchema  `json:"user"`
            RegisterToken string  `json:"register_token"`
            ResendEmailToken string  `json:"resend_email_token"`
            UserExists bool  `json:"user_exists"`
            VerifyEmailLink bool  `json:"verify_email_link"`
            VerifyEmailOtp bool  `json:"verify_email_otp"`
            VerifyMobileOtp bool  `json:"verify_mobile_otp"`
            Email string  `json:"email"`
            RequestID string  `json:"request_id"`
            CountryCode string  `json:"country_code"`
            Mobile string  `json:"mobile"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
            ResendTimer float64  `json:"resend_timer"`
            ResendToken string  `json:"resend_token"`
         
    }
    
    // LoginSuccess used by User
    type LoginSuccess struct {

        
            User UserSchema  `json:"user"`
            RequestID string  `json:"request_id"`
            RegisterToken string  `json:"register_token"`
         
    }
    
    // VerifyOtpSuccess used by User
    type VerifyOtpSuccess struct {

        
            User UserSchema  `json:"user"`
            UserExists bool  `json:"user_exists"`
            RegisterToken string  `json:"register_token"`
         
    }
    
    // ResetPasswordSuccess used by User
    type ResetPasswordSuccess struct {

        
            Status string  `json:"status"`
         
    }
    
    // RegisterFormSuccess used by User
    type RegisterFormSuccess struct {

        
            Email string  `json:"email"`
            ResendTimer float64  `json:"resend_timer"`
            ResendToken string  `json:"resend_token"`
            ResendEmailToken string  `json:"resend_email_token"`
            RegisterToken string  `json:"register_token"`
            Success bool  `json:"success"`
            RequestID string  `json:"request_id"`
            Message string  `json:"message"`
            Mobile string  `json:"mobile"`
            CountryCode string  `json:"country_code"`
            VerifyEmailOtp bool  `json:"verify_email_otp"`
            VerifyMobileOtp bool  `json:"verify_mobile_otp"`
            UserExists bool  `json:"user_exists"`
         
    }
    
    // VerifyEmailSuccess used by User
    type VerifyEmailSuccess struct {

        
            Message string  `json:"message"`
         
    }
    
    // HasPasswordSuccess used by User
    type HasPasswordSuccess struct {

        
            Result bool  `json:"result"`
         
    }
    
    // LogoutSuccess used by User
    type LogoutSuccess struct {

        
            Logout bool  `json:"logout"`
         
    }
    
    // BlockUserSuccess used by User
    type BlockUserSuccess struct {

        
            Success bool  `json:"success"`
         
    }
    
    // ArchiveUserSuccess used by User
    type ArchiveUserSuccess struct {

        
            Success bool  `json:"success"`
         
    }
    
    // DeleteUserSuccess used by User
    type DeleteUserSuccess struct {

        
            Success bool  `json:"success"`
         
    }
    
    // UnDeleteUserSuccess used by User
    type UnDeleteUserSuccess struct {

        
            Success bool  `json:"success"`
         
    }
    
    // OtpSuccess used by User
    type OtpSuccess struct {

        
            ResendTimer float64  `json:"resend_timer"`
            ResendToken string  `json:"resend_token"`
            RegisterToken string  `json:"register_token"`
            Success bool  `json:"success"`
            RequestID string  `json:"request_id"`
            Message string  `json:"message"`
            Mobile string  `json:"mobile"`
            CountryCode string  `json:"country_code"`
         
    }
    
    // EmailOtpSuccess used by User
    type EmailOtpSuccess struct {

        
            Success bool  `json:"success"`
         
    }
    
    // SessionListSuccess used by User
    type SessionListSuccess struct {

        
            Sessions []string  `json:"sessions"`
         
    }
    
    // VerifyMobileOTPSuccess used by User
    type VerifyMobileOTPSuccess struct {

        
            User UserSchema  `json:"user"`
            VerifyMobileLink bool  `json:"verify_mobile_link"`
         
    }
    
    // VerifyEmailOTPSuccess used by User
    type VerifyEmailOTPSuccess struct {

        
            User UserSchema  `json:"user"`
            VerifyEmailLink bool  `json:"verify_email_link"`
         
    }
    
    // SendMobileVerifyLinkSuccess used by User
    type SendMobileVerifyLinkSuccess struct {

        
            VerifyMobileLink bool  `json:"verify_mobile_link"`
         
    }
    
    // SendEmailVerifyLinkSuccess used by User
    type SendEmailVerifyLinkSuccess struct {

        
            VerifyEmailLink bool  `json:"verify_email_link"`
         
    }
    
    // UserSearchResponseSchema used by User
    type UserSearchResponseSchema struct {

        
            Users []UserSchema  `json:"users"`
         
    }
    
    // CustomerListResponseSchema used by User
    type CustomerListResponseSchema struct {

        
            Items []UserSchema  `json:"items"`
            Page PaginationSchema  `json:"page"`
         
    }
    
    // PaginationSchema used by User
    type PaginationSchema struct {

        
            Size float64  `json:"size"`
            ItemTotal float64  `json:"item_total"`
            HasNext bool  `json:"has_next"`
            Type string  `json:"type"`
            Current float64  `json:"current"`
         
    }
    
    // SessionListResponseSchema used by User
    type SessionListResponseSchema struct {

        
            Items []string  `json:"items"`
         
    }
    
    // SessionDeleteResponseSchema used by User
    type SessionDeleteResponseSchema struct {

        
            Items []string  `json:"items"`
         
    }
    
    // UnauthorizedSchema used by User
    type UnauthorizedSchema struct {

        
            Message string  `json:"message"`
         
    }
    
    // UnauthenticatedSchema used by User
    type UnauthenticatedSchema struct {

        
            Authenticated bool  `json:"authenticated"`
         
    }
    
    // NotFoundSchema used by User
    type NotFoundSchema struct {

        
            Message string  `json:"message"`
         
    }
    
    // AuthenticationInternalServerErrorSchema used by User
    type AuthenticationInternalServerErrorSchema struct {

        
            Message string  `json:"message"`
         
    }
    
    // AuthenticationApiErrorSchema used by User
    type AuthenticationApiErrorSchema struct {

        
            Message string  `json:"message"`
         
    }
    
    // ProfileEditSuccessSchema used by User
    type ProfileEditSuccessSchema struct {

        
            Email string  `json:"email"`
            VerifyEmailOtp bool  `json:"verify_email_otp"`
            VerifyEmailLink bool  `json:"verify_email_link"`
            VerifyMobileOtp bool  `json:"verify_mobile_otp"`
            User string  `json:"user"`
            RegisterToken string  `json:"register_token"`
            UserExists bool  `json:"user_exists"`
         
    }
    
    // FormRegisterRequestSchemaPhone used by User
    type FormRegisterRequestSchemaPhone struct {

        
            CountryCode string  `json:"country_code"`
            Mobile string  `json:"mobile"`
         
    }
    
    // OAuthRequestSchemaOauth2 used by User
    type OAuthRequestSchemaOauth2 struct {

        
            AccessToken string  `json:"access_token"`
            Expiry float64  `json:"expiry"`
            RefreshToken string  `json:"refresh_token"`
         
    }
    
    // OAuthRequestSchemaProfile used by User
    type OAuthRequestSchemaProfile struct {

        
            LastName string  `json:"last_name"`
            Image string  `json:"image"`
            ID string  `json:"id"`
            Email string  `json:"email"`
            FullName string  `json:"full_name"`
            FirstName string  `json:"first_name"`
         
    }
    
    // OAuthRequestAppleSchemaOauth used by User
    type OAuthRequestAppleSchemaOauth struct {

        
            IdentityToken string  `json:"identity_token"`
         
    }
    
    // OAuthRequestAppleSchemaProfile used by User
    type OAuthRequestAppleSchemaProfile struct {

        
            LastName string  `json:"last_name"`
            FullName string  `json:"full_name"`
            FirstName string  `json:"first_name"`
         
    }
    
    // AuthSuccessUser used by User
    type AuthSuccessUser struct {

        
            FirstName string  `json:"first_name"`
            LastName string  `json:"last_name"`
            Debug AuthSuccessUserDebug  `json:"debug"`
            Active bool  `json:"active"`
            Emails AuthSuccessUserEmails  `json:"emails"`
         
    }
    
    // AuthSuccessUserDebug used by User
    type AuthSuccessUserDebug struct {

        
            Platform string  `json:"platform"`
         
    }
    
    // AuthSuccessUserEmails used by User
    type AuthSuccessUserEmails struct {

        
            Email string  `json:"email"`
            Verified bool  `json:"verified"`
            Primary bool  `json:"primary"`
            Active bool  `json:"active"`
         
    }
    
    // CreateUserRequestSchema used by User
    type CreateUserRequestSchema struct {

        
            PhoneNumber string  `json:"phone_number"`
            Email string  `json:"email"`
            FirstName string  `json:"first_name"`
            LastName string  `json:"last_name"`
            Gender string  `json:"gender"`
            Username string  `json:"username"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // CreateUserResponseSchema used by User
    type CreateUserResponseSchema struct {

        
            User UserSchema  `json:"user"`
         
    }
    
    // CreateUserSessionRequestSchema used by User
    type CreateUserSessionRequestSchema struct {

        
            Domain string  `json:"domain"`
            MaxAge float64  `json:"max_age"`
            UserID string  `json:"user_id"`
         
    }
    
    // CreateUserSessionResponseSchema used by User
    type CreateUserSessionResponseSchema struct {

        
            Domain string  `json:"domain"`
            MaxAge float64  `json:"max_age"`
            Secure bool  `json:"secure"`
            HttpOnly bool  `json:"http_only"`
            Cookie map[string]interface{}  `json:"cookie"`
         
    }
    
    // PlatformSchema used by User
    type PlatformSchema struct {

        
            Display string  `json:"display"`
            LookAndFeel LookAndFeel  `json:"look_and_feel"`
            UpdatedAt string  `json:"updated_at"`
            Active bool  `json:"active"`
            ForgotPassword bool  `json:"forgot_password"`
            Login Login  `json:"login"`
            SkipCaptcha bool  `json:"skip_captcha"`
            Name string  `json:"name"`
            Meta MetaSchema  `json:"meta"`
            ID string  `json:"_id"`
            Social Social  `json:"social"`
            RequiredFields RequiredFields  `json:"required_fields"`
            RegisterRequiredFields RegisterRequiredFields  `json:"register_required_fields"`
            SkipLogin bool  `json:"skip_login"`
            FlashCard FlashCard  `json:"flash_card"`
            Subtext string  `json:"subtext"`
            SocialTokens SocialTokens  `json:"social_tokens"`
            CreatedAt string  `json:"created_at"`
            Register bool  `json:"register"`
            MobileImage string  `json:"mobile_image"`
            DesktopImage string  `json:"desktop_image"`
            DeleteAccountDay float64  `json:"delete_account_day"`
            DeleteAccountReasons []DeleteAccountReasons  `json:"delete_account_reasons"`
            DeleteAccountConsent map[string]interface{}  `json:"delete_account_consent"`
            SessionConfig map[string]interface{}  `json:"session_config"`
         
    }
    
    // LookAndFeel used by User
    type LookAndFeel struct {

        
            CardPosition string  `json:"card_position"`
            BackgroundColor string  `json:"background_color"`
         
    }
    
    // Login used by User
    type Login struct {

        
            Password bool  `json:"password"`
            Otp bool  `json:"otp"`
         
    }
    
    // MetaSchema used by User
    type MetaSchema struct {

        
            FyndDefault bool  `json:"fynd_default"`
         
    }
    
    // Social used by User
    type Social struct {

        
            AccountKit bool  `json:"account_kit"`
            Facebook bool  `json:"facebook"`
            Google bool  `json:"google"`
            Apple bool  `json:"apple"`
         
    }
    
    // RequiredFields used by User
    type RequiredFields struct {

        
            Email PlatformEmail  `json:"email"`
            Mobile PlatformMobile  `json:"mobile"`
         
    }
    
    // PlatformEmail used by User
    type PlatformEmail struct {

        
            IsRequired bool  `json:"is_required"`
            Level string  `json:"level"`
         
    }
    
    // PlatformMobile used by User
    type PlatformMobile struct {

        
            IsRequired bool  `json:"is_required"`
            Level string  `json:"level"`
         
    }
    
    // RegisterRequiredFields used by User
    type RegisterRequiredFields struct {

        
            Email RegisterRequiredFieldsEmail  `json:"email"`
            Mobile RegisterRequiredFieldsMobile  `json:"mobile"`
         
    }
    
    // RegisterRequiredFieldsEmail used by User
    type RegisterRequiredFieldsEmail struct {

        
            IsRequired bool  `json:"is_required"`
            Level string  `json:"level"`
         
    }
    
    // RegisterRequiredFieldsMobile used by User
    type RegisterRequiredFieldsMobile struct {

        
            IsRequired bool  `json:"is_required"`
            Level string  `json:"level"`
         
    }
    
    // FlashCard used by User
    type FlashCard struct {

        
            Text string  `json:"text"`
            TextColor string  `json:"text_color"`
            BackgroundColor string  `json:"background_color"`
         
    }
    
    // SocialTokens used by User
    type SocialTokens struct {

        
            Facebook Facebook  `json:"facebook"`
            AccountKit Accountkit  `json:"account_kit"`
            Google Google  `json:"google"`
         
    }
    
    // DeleteAccountReasons used by User
    type DeleteAccountReasons struct {

        
            ReasonText string  `json:"reason_text"`
            ReasonID string  `json:"reason_id"`
            ShowTextArea bool  `json:"show_text_area"`
         
    }
    
    // DeleteAccountConsent used by User
    type DeleteAccountConsent struct {

        
            ConsentText string  `json:"consent_text"`
         
    }
    
    // Facebook used by User
    type Facebook struct {

        
            AppID string  `json:"app_id"`
         
    }
    
    // Accountkit used by User
    type Accountkit struct {

        
            AppID string  `json:"app_id"`
         
    }
    
    // Google used by User
    type Google struct {

        
            AppID string  `json:"app_id"`
         
    }
    
    // SessionExpiry used by User
    type SessionExpiry struct {

        
            Duration float64  `json:"duration"`
            Type string  `json:"type"`
            IsRolling bool  `json:"is_rolling"`
         
    }
    
    // UpdateUserRequestSchema used by User
    type UpdateUserRequestSchema struct {

        
            FirstName string  `json:"first_name"`
            LastName string  `json:"last_name"`
            Gender string  `json:"gender"`
            ExternalID string  `json:"external_id"`
            Meta map[string]interface{}  `json:"meta"`
            PhoneNumbers []UserPhoneNumbers  `json:"phone_numbers"`
            Emails []UserEmails  `json:"emails"`
         
    }
    
    // UserEmails used by User
    type UserEmails struct {

        
            Active bool  `json:"active"`
            Primary bool  `json:"primary"`
            Verified bool  `json:"verified"`
            Email string  `json:"email"`
         
    }
    
    // UserPhoneNumbers used by User
    type UserPhoneNumbers struct {

        
            Active bool  `json:"active"`
            Primary bool  `json:"primary"`
            Verified bool  `json:"verified"`
            Phone string  `json:"phone"`
            CountryCode string  `json:"country_code"`
         
    }
    
    // UserSchema used by User
    type UserSchema struct {

        
            ApplicationID string  `json:"application_id"`
            UserID string  `json:"user_id"`
            FirstName string  `json:"first_name"`
            Meta map[string]interface{}  `json:"meta"`
            LastName string  `json:"last_name"`
            PhoneNumbers []PhoneNumber  `json:"phone_numbers"`
            Emails []Email  `json:"emails"`
            Gender string  `json:"gender"`
            Dob string  `json:"dob"`
            Active bool  `json:"active"`
            ProfilePicURL string  `json:"profile_pic_url"`
            Username string  `json:"username"`
            AccountType string  `json:"account_type"`
            ID string  `json:"_id"`
            CreatedAt string  `json:"created_at"`
            UpdatedAt string  `json:"updated_at"`
         
    }
    
    // PhoneNumber used by User
    type PhoneNumber struct {

        
            Active bool  `json:"active"`
            Primary bool  `json:"primary"`
            Verified bool  `json:"verified"`
            Phone string  `json:"phone"`
            CountryCode float64  `json:"country_code"`
         
    }
    
    // Email used by User
    type Email struct {

        
            Primary bool  `json:"primary"`
            Verified bool  `json:"verified"`
            Email string  `json:"email"`
            Active bool  `json:"active"`
         
    }
    

    
    // ApplicationLegal used by Content
    type ApplicationLegal struct {

        
            Application string  `json:"application"`
            Tnc string  `json:"tnc"`
            Policy string  `json:"policy"`
            Shipping string  `json:"shipping"`
            Returns string  `json:"returns"`
            Faq []ApplicationLegalFAQ  `json:"faq"`
            ID string  `json:"_id"`
            UpdatedAt string  `json:"updated_at"`
            CreatedAt string  `json:"created_at"`
         
    }
    
    // ApplicationLegalFAQ used by Content
    type ApplicationLegalFAQ struct {

        
            Question string  `json:"question"`
            Answer string  `json:"answer"`
         
    }
    
    // PathMappingSchema used by Content
    type PathMappingSchema struct {

        
            Application string  `json:"application"`
            ID string  `json:"_id"`
            RedirectFrom string  `json:"redirect_from"`
            RedirectTo string  `json:"redirect_to"`
            UpdatedAt string  `json:"updated_at"`
            CreatedAt string  `json:"created_at"`
            Source TagSourceSchema  `json:"__source"`
         
    }
    
    // SeoComponent used by Content
    type SeoComponent struct {

        
            Seo SeoSchema  `json:"seo"`
         
    }
    
    // SeoSchema used by Content
    type SeoSchema struct {

        
            App string  `json:"app"`
            ID string  `json:"_id"`
            RobotsTxt string  `json:"robots_txt"`
            SitemapEnabled bool  `json:"sitemap_enabled"`
            CustomMetaTags []map[string]interface{}  `json:"custom_meta_tags"`
            Details Detail  `json:"details"`
            CreatedAt string  `json:"created_at"`
            UpdatedAt string  `json:"updated_at"`
         
    }
    
    // CustomMetaTag used by Content
    type CustomMetaTag struct {

        
            Name string  `json:"name"`
            Content string  `json:"content"`
            ID string  `json:"_id"`
         
    }
    
    // Detail used by Content
    type Detail struct {

        
            Title string  `json:"title"`
            Description string  `json:"description"`
         
    }
    
    // AnnouncementPageSchema used by Content
    type AnnouncementPageSchema struct {

        
            PageSlug string  `json:"page_slug"`
            Type string  `json:"type"`
         
    }
    
    // EditorMeta used by Content
    type EditorMeta struct {

        
            ForegroundColor string  `json:"foreground_color"`
            BackgroundColor string  `json:"background_color"`
            ContentType string  `json:"content_type"`
            Content string  `json:"content"`
         
    }
    
    // AnnouncementAuthorSchema used by Content
    type AnnouncementAuthorSchema struct {

        
            CreatedBy string  `json:"created_by"`
            ModifiedBy string  `json:"modified_by"`
         
    }
    
    // AdminAnnouncementSchema used by Content
    type AdminAnnouncementSchema struct {

        
            ID string  `json:"_id"`
            Platforms []string  `json:"platforms"`
            Title string  `json:"title"`
            Announcement string  `json:"announcement"`
            Pages []AnnouncementPageSchema  `json:"pages"`
            EditorMeta EditorMeta  `json:"editor_meta"`
            Author AnnouncementAuthorSchema  `json:"author"`
            CreatedAt string  `json:"created_at"`
            App string  `json:"app"`
            ModifiedAt string  `json:"modified_at"`
            Schedule ScheduleSchema  `json:"_schedule"`
         
    }
    
    // ScheduleSchema used by Content
    type ScheduleSchema struct {

        
            Cron string  `json:"cron"`
            Start string  `json:"start"`
            End string  `json:"end"`
            Duration float64  `json:"duration"`
            NextSchedule []map[string]interface{}  `json:"next_schedule"`
         
    }
    
    // NextSchedule used by Content
    type NextSchedule struct {

        
            Start string  `json:"start"`
            End string  `json:"end"`
         
    }
    
    // AnnouncementSchema used by Content
    type AnnouncementSchema struct {

        
            Announcement string  `json:"announcement"`
            Schedule ScheduleStartSchema  `json:"schedule"`
         
    }
    
    // ScheduleStartSchema used by Content
    type ScheduleStartSchema struct {

        
            Start string  `json:"start"`
            End string  `json:"end"`
         
    }
    
    // BlogGetResponse used by Content
    type BlogGetResponse struct {

        
            Items []BlogSchema  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // ResourceContent used by Content
    type ResourceContent struct {

        
            Type string  `json:"type"`
            Value string  `json:"value"`
         
    }
    
    // Asset used by Content
    type Asset struct {

        
            AspectRatio string  `json:"aspect_ratio"`
            ID string  `json:"id"`
            SecureURL string  `json:"secure_url"`
         
    }
    
    // Author used by Content
    type Author struct {

        
            Designation string  `json:"designation"`
            ID string  `json:"id"`
            Name string  `json:"name"`
         
    }
    
    // BlogSchema used by Content
    type BlogSchema struct {

        
            ID string  `json:"_id"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Application string  `json:"application"`
            Archived bool  `json:"archived"`
            Author Author  `json:"author"`
            Content []ResourceContent  `json:"content"`
            FeatureImage Asset  `json:"feature_image"`
            Published bool  `json:"published"`
            ReadingTime string  `json:"reading_time"`
            Slug string  `json:"slug"`
            Tags []string  `json:"tags"`
            Seo SEO  `json:"seo"`
            Schedule CronSchedule  `json:"_schedule"`
            Title string  `json:"title"`
            DateMeta DateMeta  `json:"date_meta"`
         
    }
    
    // SEO used by Content
    type SEO struct {

        
            Description string  `json:"description"`
            Image SEOImage  `json:"image"`
            Title string  `json:"title"`
         
    }
    
    // SEOImage used by Content
    type SEOImage struct {

        
            URL string  `json:"url"`
         
    }
    
    // DateMeta used by Content
    type DateMeta struct {

        
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // BlogRequest used by Content
    type BlogRequest struct {

        
            Application string  `json:"application"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Author Author  `json:"author"`
            Content []ResourceContent  `json:"content"`
            FeatureImage Asset  `json:"feature_image"`
            Published bool  `json:"published"`
            ReadingTime string  `json:"reading_time"`
            Slug string  `json:"slug"`
            Tags []string  `json:"tags"`
            Title string  `json:"title"`
            Seo SEO  `json:"seo"`
            Schedule CronSchedule  `json:"_schedule"`
         
    }
    
    // GetAnnouncementListSchema used by Content
    type GetAnnouncementListSchema struct {

        
            Items []AdminAnnouncementSchema  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // CreateAnnouncementSchema used by Content
    type CreateAnnouncementSchema struct {

        
            Message string  `json:"message"`
            Data AdminAnnouncementSchema  `json:"data"`
         
    }
    
    // DataLoaderResponseSchema used by Content
    type DataLoaderResponseSchema struct {

        
            Application string  `json:"application"`
            Company string  `json:"company"`
            ID string  `json:"_id"`
            Name string  `json:"name"`
            Service string  `json:"service"`
            OperationID string  `json:"operation_id"`
            Type string  `json:"type"`
            URL string  `json:"url"`
            Content string  `json:"content"`
            Source DataLoaderSourceSchema  `json:"__source"`
         
    }
    
    // DataLoaderResetResponseSchema used by Content
    type DataLoaderResetResponseSchema struct {

        
            Reset string  `json:"reset"`
         
    }
    
    // Navigation used by Content
    type Navigation struct {

        
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Orientation string  `json:"orientation"`
            CreatedBy CreatedBySchema  `json:"created_by"`
            DateMeta DateMeta  `json:"date_meta"`
            ID string  `json:"_id"`
            Position string  `json:"position"`
            Application string  `json:"application"`
            Platform string  `json:"platform"`
            Navigation NavigationReference  `json:"navigation"`
         
    }
    
    // LocaleLanguage used by Content
    type LocaleLanguage struct {

        
            Hi Language  `json:"hi"`
            Ar Language  `json:"ar"`
            EnUs Language  `json:"en_us"`
         
    }
    
    // Language used by Content
    type Language struct {

        
            Display string  `json:"display"`
         
    }
    
    // Action used by Content
    type Action struct {

        
            Page ActionPage  `json:"page"`
            Popup ActionPage  `json:"popup"`
            Type string  `json:"type"`
         
    }
    
    // ActionPage used by Content
    type ActionPage struct {

        
            Params map[string][]string  `json:"params"`
            Query map[string][]string  `json:"query"`
            URL string  `json:"url"`
            Type PageType  `json:"type"`
         
    }
    
    // NavigationReference used by Content
    type NavigationReference struct {

        
            Acl []string  `json:"acl"`
            Tags []string  `json:"tags"`
            LocaleLanguage LocaleLanguage  `json:"_locale_language"`
            Image string  `json:"image"`
            Type string  `json:"type"`
            Action Action  `json:"action"`
            Active bool  `json:"active"`
            Display string  `json:"display"`
            SortOrder float64  `json:"sort_order"`
            SubNavigation []NavigationReference  `json:"sub_navigation"`
         
    }
    
    // LandingPage used by Content
    type LandingPage struct {

        
            Data LandingPageSchema  `json:"data"`
            Success bool  `json:"success"`
         
    }
    
    // ConfigurationSchema used by Content
    type ConfigurationSchema struct {

        
            SleepTime float64  `json:"sleep_time"`
            StartOnLaunch bool  `json:"start_on_launch"`
            Duration float64  `json:"duration"`
            SlideDirection string  `json:"slide_direction"`
         
    }
    
    // SlideshowMedia used by Content
    type SlideshowMedia struct {

        
            Type string  `json:"type"`
            URL string  `json:"url"`
            BgColor string  `json:"bg_color"`
            Duration float64  `json:"duration"`
            AutoDecideDuration bool  `json:"auto_decide_duration"`
            Action Action  `json:"action"`
         
    }
    
    // Slideshow used by Content
    type Slideshow struct {

        
            Data SlideshowSchema  `json:"data"`
            Success bool  `json:"success"`
         
    }
    
    // AnnouncementsResponseSchema used by Content
    type AnnouncementsResponseSchema struct {

        
            Announcements map[string][]AnnouncementSchema  `json:"announcements"`
            RefreshRate float64  `json:"refresh_rate"`
            RefreshPages []string  `json:"refresh_pages"`
         
    }
    
    // FaqResponseSchema used by Content
    type FaqResponseSchema struct {

        
            Faqs []FaqSchema  `json:"faqs"`
         
    }
    
    // UpdateHandpickedSchema used by Content
    type UpdateHandpickedSchema struct {

        
            Tag HandpickedTagSchema  `json:"tag"`
         
    }
    
    // HandpickedTagSchema used by Content
    type HandpickedTagSchema struct {

        
            Position string  `json:"position"`
            Attributes map[string]interface{}  `json:"attributes"`
            Name string  `json:"name"`
            URL string  `json:"url"`
            Type string  `json:"type"`
            SubType string  `json:"sub_type"`
            Content string  `json:"content"`
         
    }
    
    // RemoveHandpickedSchema used by Content
    type RemoveHandpickedSchema struct {

        
            Tags []string  `json:"tags"`
         
    }
    
    // CreateTagSchema used by Content
    type CreateTagSchema struct {

        
            Name string  `json:"name"`
            SubType string  `json:"sub_type"`
            ID string  `json:"_id"`
            Type string  `json:"type"`
            URL string  `json:"url"`
            Position string  `json:"position"`
            Attributes map[string]interface{}  `json:"attributes"`
            Pages []map[string]interface{}  `json:"pages"`
            Content string  `json:"content"`
         
    }
    
    // CreateTagRequestSchema used by Content
    type CreateTagRequestSchema struct {

        
            Tags []CreateTagSchema  `json:"tags"`
         
    }
    
    // DataLoaderSchema used by Content
    type DataLoaderSchema struct {

        
            Name string  `json:"name"`
            Service string  `json:"service"`
            OperationID string  `json:"operation_id"`
            Type string  `json:"type"`
            URL string  `json:"url"`
            Content string  `json:"content"`
            Source DataLoaderSourceSchema  `json:"__source"`
            ID string  `json:"_id"`
         
    }
    
    // DataLoaderSourceSchema used by Content
    type DataLoaderSourceSchema struct {

        
            Type string  `json:"type"`
            ID string  `json:"id"`
         
    }
    
    // DataLoadersSchema used by Content
    type DataLoadersSchema struct {

        
            Items []DataLoaderSchema  `json:"items"`
         
    }
    
    // TagDeleteSuccessResponse used by Content
    type TagDeleteSuccessResponse struct {

        
            Success bool  `json:"success"`
         
    }
    
    // ContentAPIError used by Content
    type ContentAPIError struct {

        
            Message string  `json:"message"`
            Status float64  `json:"status"`
            Code string  `json:"code"`
            Exception string  `json:"exception"`
            Info string  `json:"info"`
            RequestID string  `json:"request_id"`
            StackTrace string  `json:"stack_trace"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // CommonError used by Content
    type CommonError struct {

        
            Message string  `json:"message"`
         
    }
    
    // CategorySchema used by Content
    type CategorySchema struct {

        
            Index float64  `json:"index"`
            Title string  `json:"title"`
            Description string  `json:"description"`
            Children []string  `json:"children"`
            ID string  `json:"_id"`
            Slug string  `json:"slug"`
            Application string  `json:"application"`
            IconURL string  `json:"icon_url"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // ChildrenSchema used by Content
    type ChildrenSchema struct {

        
            Question string  `json:"question"`
            Answer string  `json:"answer"`
            Slug string  `json:"slug"`
            Application string  `json:"application"`
            ID string  `json:"_id"`
         
    }
    
    // CategoryRequestSchema used by Content
    type CategoryRequestSchema struct {

        
            Slug string  `json:"slug"`
            Title string  `json:"title"`
         
    }
    
    // FAQCategorySchema used by Content
    type FAQCategorySchema struct {

        
            Index float64  `json:"index"`
            Title string  `json:"title"`
            Description string  `json:"description"`
            Children []ChildrenSchema  `json:"children"`
            ID string  `json:"_id"`
            Slug string  `json:"slug"`
            Application string  `json:"application"`
            IconURL string  `json:"icon_url"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // FaqSchema used by Content
    type FaqSchema struct {

        
            Slug string  `json:"slug"`
            Application string  `json:"application"`
            ID string  `json:"_id"`
            Question string  `json:"question"`
            Answer string  `json:"answer"`
            Tags []string  `json:"tags"`
         
    }
    
    // FAQ used by Content
    type FAQ struct {

        
            Slug string  `json:"slug"`
            Question string  `json:"question"`
            Answer string  `json:"answer"`
         
    }
    
    // CreateFaqResponseSchema used by Content
    type CreateFaqResponseSchema struct {

        
            Faq FaqSchema  `json:"faq"`
         
    }
    
    // CreateFaqSchema used by Content
    type CreateFaqSchema struct {

        
            Faq FAQ  `json:"faq"`
         
    }
    
    // GetFaqSchema used by Content
    type GetFaqSchema struct {

        
            Faqs []FaqSchema  `json:"faqs"`
         
    }
    
    // UpdateFaqCategoryRequestSchema used by Content
    type UpdateFaqCategoryRequestSchema struct {

        
            Category CategorySchema  `json:"category"`
         
    }
    
    // CreateFaqCategoryRequestSchema used by Content
    type CreateFaqCategoryRequestSchema struct {

        
            Category CategoryRequestSchema  `json:"category"`
         
    }
    
    // CreateFaqCategorySchema used by Content
    type CreateFaqCategorySchema struct {

        
            Category CategorySchema  `json:"category"`
         
    }
    
    // GetFaqCategoriesSchema used by Content
    type GetFaqCategoriesSchema struct {

        
            Categories []CategorySchema  `json:"categories"`
         
    }
    
    // GetFaqCategoryBySlugSchema used by Content
    type GetFaqCategoryBySlugSchema struct {

        
            Category FAQCategorySchema  `json:"category"`
         
    }
    
    // Page used by Content
    type Page struct {

        
            ItemTotal float64  `json:"item_total"`
            NextID string  `json:"next_id"`
            HasPrevious bool  `json:"has_previous"`
            HasNext bool  `json:"has_next"`
            Current float64  `json:"current"`
            Type string  `json:"type"`
            Size float64  `json:"size"`
         
    }
    
    // LandingPageGetResponse used by Content
    type LandingPageGetResponse struct {

        
            Items []LandingPageSchema  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // LandingPageSchema used by Content
    type LandingPageSchema struct {

        
            Slug string  `json:"slug"`
            Action Action  `json:"action"`
            Platform []string  `json:"platform"`
            CreatedBy CreatedBySchema  `json:"created_by"`
            DateMeta DateMeta  `json:"date_meta"`
            ID string  `json:"_id"`
            Application string  `json:"application"`
            Archived bool  `json:"archived"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // DefaultNavigationResponse used by Content
    type DefaultNavigationResponse struct {

        
            Items []NavigationSchema  `json:"items"`
         
    }
    
    // NavigationGetResponse used by Content
    type NavigationGetResponse struct {

        
            Items []NavigationSchema  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // Orientation used by Content
    type Orientation struct {

        
            Portrait []string  `json:"portrait"`
            Landscape []string  `json:"landscape"`
         
    }
    
    // NavigationSchema used by Content
    type NavigationSchema struct {

        
            ID string  `json:"_id"`
            Application string  `json:"application"`
            Archived bool  `json:"archived"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Platform []string  `json:"platform"`
            CreatedBy CreatedBySchema  `json:"created_by"`
            DateMeta DateMeta  `json:"date_meta"`
            Orientation Orientation  `json:"orientation"`
            Version float64  `json:"version"`
            Navigation []NavigationReference  `json:"navigation"`
         
    }
    
    // NavigationRequest used by Content
    type NavigationRequest struct {

        
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Platform []string  `json:"platform"`
            Orientation Orientation  `json:"orientation"`
            Navigation []NavigationReference  `json:"navigation"`
         
    }
    
    // CustomPageSchema used by Content
    type CustomPageSchema struct {

        
            ID string  `json:"_id"`
            Platform string  `json:"platform"`
            Title string  `json:"title"`
            Slug string  `json:"slug"`
            Type string  `json:"type"`
            Orientation string  `json:"orientation"`
            Application string  `json:"application"`
            Description string  `json:"description"`
            Published bool  `json:"published"`
            Tags []string  `json:"tags"`
            Content []map[string]interface{}  `json:"content"`
            CreatedBy CreatedBySchema  `json:"created_by"`
            DateMeta DateMeta  `json:"date_meta"`
            Schedule ScheduleSchema  `json:"_schedule"`
         
    }
    
    // ContentSchema used by Content
    type ContentSchema struct {

        
            Type string  `json:"type"`
            Value map[string]interface{}  `json:"value"`
         
    }
    
    // CustomPage used by Content
    type CustomPage struct {

        
            Data CustomPageSchema  `json:"data"`
         
    }
    
    // FeatureImage used by Content
    type FeatureImage struct {

        
            SecureURL string  `json:"secure_url"`
         
    }
    
    // PageGetResponse used by Content
    type PageGetResponse struct {

        
            Items []PageSchema  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // PageSpec used by Content
    type PageSpec struct {

        
            Specifications []map[string]interface{}  `json:"specifications"`
         
    }
    
    // PageSpecParam used by Content
    type PageSpecParam struct {

        
            Key string  `json:"key"`
            Required bool  `json:"required"`
         
    }
    
    // PageSpecItem used by Content
    type PageSpecItem struct {

        
            PageType string  `json:"page_type"`
            DisplayName string  `json:"display_name"`
            Params []PageSpecParam  `json:"params"`
            Query []PageSpecParam  `json:"query"`
         
    }
    
    // PageSchema used by Content
    type PageSchema struct {

        
            ID string  `json:"_id"`
            Application string  `json:"application"`
            ComponentIds []string  `json:"component_ids"`
            Content []map[string]interface{}  `json:"content"`
            ContentPath string  `json:"content_path"`
            CreatedBy CreatedBySchema  `json:"created_by"`
            DateMeta DateMeta  `json:"date_meta"`
            Description string  `json:"description"`
            FeatureImage Asset  `json:"feature_image"`
            PageMeta []map[string]interface{}  `json:"page_meta"`
            Schedule ScheduleSchema  `json:"_schedule"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Orientation string  `json:"orientation"`
            Platform string  `json:"platform"`
            Published bool  `json:"published"`
            Slug string  `json:"slug"`
            Tags []string  `json:"tags"`
            Title string  `json:"title"`
            Type string  `json:"type"`
            Seo SEO  `json:"seo"`
            Visibility map[string]interface{}  `json:"visibility"`
            Archived bool  `json:"archived"`
         
    }
    
    // CreatedBySchema used by Content
    type CreatedBySchema struct {

        
            ID string  `json:"id"`
         
    }
    
    // PageContent used by Content
    type PageContent struct {

        
            Type string  `json:"type"`
            Value map[string]interface{}  `json:"value"`
         
    }
    
    // PageMeta used by Content
    type PageMeta struct {

        
            Key string  `json:"key"`
            Value map[string]interface{}  `json:"value"`
         
    }
    
    // PageRequest used by Content
    type PageRequest struct {

        
            Schedule CronSchedule  `json:"_schedule"`
            Application string  `json:"application"`
            Author Author  `json:"author"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Orientation string  `json:"orientation"`
            Content []map[string]interface{}  `json:"content"`
            FeatureImage Asset  `json:"feature_image"`
            Published bool  `json:"published"`
            ReadingTime string  `json:"reading_time"`
            Slug string  `json:"slug"`
            Tags []string  `json:"tags"`
            Seo SEO  `json:"seo"`
            Title string  `json:"title"`
         
    }
    
    // CronSchedule used by Content
    type CronSchedule struct {

        
            Cron string  `json:"cron"`
            Start string  `json:"start"`
            End string  `json:"end"`
            Duration float64  `json:"duration"`
         
    }
    
    // PagePublishRequest used by Content
    type PagePublishRequest struct {

        
            Publish bool  `json:"publish"`
         
    }
    
    // PageMetaSchema used by Content
    type PageMetaSchema struct {

        
            SystemPages []NavigationSchema  `json:"system_pages"`
            CustomPages []PageSchema  `json:"custom_pages"`
            ApplicationID string  `json:"application_id"`
         
    }
    
    // SlideshowGetResponse used by Content
    type SlideshowGetResponse struct {

        
            Items []SlideshowSchema  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // SlideshowSchema used by Content
    type SlideshowSchema struct {

        
            ID string  `json:"_id"`
            Slug string  `json:"slug"`
            DateMeta DateMeta  `json:"date_meta"`
            Application string  `json:"application"`
            Platform string  `json:"platform"`
            Configuration ConfigurationSchema  `json:"configuration"`
            Media []SlideshowMedia  `json:"media"`
            Active bool  `json:"active"`
            Archived bool  `json:"archived"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // SlideshowRequest used by Content
    type SlideshowRequest struct {

        
            Slug string  `json:"slug"`
            Platform string  `json:"platform"`
            Configuration ConfigurationSchema  `json:"configuration"`
            Media SlideshowMedia  `json:"media"`
            Active bool  `json:"active"`
         
    }
    
    // Support used by Content
    type Support struct {

        
            Created bool  `json:"created"`
            ID string  `json:"_id"`
            ConfigType string  `json:"config_type"`
            Application string  `json:"application"`
            CreatedAt string  `json:"created_at"`
            UpdatedAt string  `json:"updated_at"`
            Contact ContactSchema  `json:"contact"`
         
    }
    
    // PhoneProperties used by Content
    type PhoneProperties struct {

        
            Key string  `json:"key"`
            Code string  `json:"code"`
            Number string  `json:"number"`
         
    }
    
    // PhoneSchema used by Content
    type PhoneSchema struct {

        
            Active bool  `json:"active"`
            Phone []PhoneProperties  `json:"phone"`
         
    }
    
    // EmailProperties used by Content
    type EmailProperties struct {

        
            Key string  `json:"key"`
            Value string  `json:"value"`
         
    }
    
    // EmailSchema used by Content
    type EmailSchema struct {

        
            Active bool  `json:"active"`
            Email []EmailProperties  `json:"email"`
         
    }
    
    // ContactSchema used by Content
    type ContactSchema struct {

        
            Phone PhoneSchema  `json:"phone"`
            Email EmailSchema  `json:"email"`
         
    }
    
    // TagsSchema used by Content
    type TagsSchema struct {

        
            Application string  `json:"application"`
            ID string  `json:"_id"`
            Tags []TagSchema  `json:"tags"`
         
    }
    
    // TagSchema used by Content
    type TagSchema struct {

        
            Name string  `json:"name"`
            URL string  `json:"url"`
            Type string  `json:"type"`
            SubType string  `json:"sub_type"`
            ID string  `json:"_id"`
            Position string  `json:"position"`
            Attributes map[string]interface{}  `json:"attributes"`
            Content string  `json:"content"`
            Pages []map[string]interface{}  `json:"pages"`
            Source TagSourceSchema  `json:"__source"`
         
    }
    
    // TagSourceSchema used by Content
    type TagSourceSchema struct {

        
            Type string  `json:"type"`
            ID string  `json:"id"`
         
    }
    

    
    // Page used by Billing
    type Page struct {

        
            ItemTotal float64  `json:"item_total"`
            NextID string  `json:"next_id"`
            HasPrevious bool  `json:"has_previous"`
            HasNext bool  `json:"has_next"`
            Current float64  `json:"current"`
            Type string  `json:"type"`
            Size float64  `json:"size"`
         
    }
    
    // UnauthenticatedUser used by Billing
    type UnauthenticatedUser struct {

        
            Message string  `json:"message"`
         
    }
    
    // UnauthenticatedApplication used by Billing
    type UnauthenticatedApplication struct {

        
            Message string  `json:"message"`
         
    }
    
    // BadRequest used by Billing
    type BadRequest struct {

        
            Message string  `json:"message"`
         
    }
    
    // ResourceNotFound used by Billing
    type ResourceNotFound struct {

        
            Message string  `json:"message"`
         
    }
    
    // InternalServerError used by Billing
    type InternalServerError struct {

        
            Message string  `json:"message"`
            Code string  `json:"code"`
         
    }
    
    // CheckValidityResponse used by Billing
    type CheckValidityResponse struct {

        
            IsValid bool  `json:"is_valid"`
            DiscountAmount float64  `json:"discount_amount"`
         
    }
    
    // PlanRecurring used by Billing
    type PlanRecurring struct {

        
            Interval string  `json:"interval"`
            IntervalCount float64  `json:"interval_count"`
         
    }
    
    // Plan used by Billing
    type Plan struct {

        
            Recurring PlanRecurring  `json:"recurring"`
            IsTrialPlan bool  `json:"is_trial_plan"`
            PlanGroup string  `json:"plan_group"`
            TagLines []string  `json:"tag_lines"`
            Currency string  `json:"currency"`
            IsActive bool  `json:"is_active"`
            IsVisible bool  `json:"is_visible"`
            TrialPeriod float64  `json:"trial_period"`
            Addons []string  `json:"addons"`
            Tags []string  `json:"tags"`
            Type string  `json:"type"`
            Country string  `json:"country"`
            ID string  `json:"_id"`
            Name string  `json:"name"`
            Description string  `json:"description"`
            Amount float64  `json:"amount"`
            ProductSuiteID string  `json:"product_suite_id"`
            CreatedAt string  `json:"created_at"`
            ModifiedAt string  `json:"modified_at"`
         
    }
    
    // DetailedPlanComponents used by Billing
    type DetailedPlanComponents struct {

        
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Description string  `json:"description"`
            Group string  `json:"group"`
            Icon string  `json:"icon"`
            Links map[string]interface{}  `json:"links"`
            Enabled bool  `json:"enabled"`
            DisplayText string  `json:"display_text"`
         
    }
    
    // DetailedPlan used by Billing
    type DetailedPlan struct {

        
            Recurring PlanRecurring  `json:"recurring"`
            IsTrialPlan bool  `json:"is_trial_plan"`
            PlanGroup string  `json:"plan_group"`
            TagLines []string  `json:"tag_lines"`
            Currency string  `json:"currency"`
            IsActive bool  `json:"is_active"`
            IsVisible bool  `json:"is_visible"`
            TrialPeriod float64  `json:"trial_period"`
            Addons []string  `json:"addons"`
            Tags []string  `json:"tags"`
            Type string  `json:"type"`
            Country string  `json:"country"`
            ID string  `json:"_id"`
            Name string  `json:"name"`
            Description string  `json:"description"`
            Amount float64  `json:"amount"`
            ProductSuiteID string  `json:"product_suite_id"`
            CreatedAt string  `json:"created_at"`
            ModifiedAt string  `json:"modified_at"`
            Components []DetailedPlanComponents  `json:"components"`
         
    }
    
    // SubscriptionTrialPeriod used by Billing
    type SubscriptionTrialPeriod struct {

        
            StartDate string  `json:"start_date"`
            EndDate string  `json:"end_date"`
         
    }
    
    // EntityChargePrice used by Billing
    type EntityChargePrice struct {

        
            Amount float64  `json:"amount"`
            CurrencyCode string  `json:"currency_code"`
         
    }
    
    // EntityChargeRecurring used by Billing
    type EntityChargeRecurring struct {

        
            Interval string  `json:"interval"`
         
    }
    
    // ChargeLineItem used by Billing
    type ChargeLineItem struct {

        
            Name string  `json:"name"`
            Term string  `json:"term"`
            PricingType string  `json:"pricing_type"`
            Price EntityChargePrice  `json:"price"`
            Recurring EntityChargeRecurring  `json:"recurring"`
            CappedAmount float64  `json:"capped_amount"`
            TrialDays float64  `json:"trial_days"`
            IsTest bool  `json:"is_test"`
            Metadata map[string]interface{}  `json:"metadata"`
         
    }
    
    // CreateSubscriptionCharge used by Billing
    type CreateSubscriptionCharge struct {

        
            Name string  `json:"name"`
            TrialDays float64  `json:"trial_days"`
            LineItems []ChargeLineItem  `json:"line_items"`
            IsTest bool  `json:"is_test"`
            ReturnURL string  `json:"return_url"`
         
    }
    
    // OneTimeChargeItem used by Billing
    type OneTimeChargeItem struct {

        
            Name string  `json:"name"`
            Term string  `json:"term"`
            PricingType string  `json:"pricing_type"`
            Price EntityChargePrice  `json:"price"`
            CappedAmount float64  `json:"capped_amount"`
            IsTest bool  `json:"is_test"`
            Metadata map[string]interface{}  `json:"metadata"`
         
    }
    
    // CreateOneTimeCharge used by Billing
    type CreateOneTimeCharge struct {

        
            Name string  `json:"name"`
            Charge OneTimeChargeItem  `json:"charge"`
            IsTest bool  `json:"is_test"`
            ReturnURL string  `json:"return_url"`
         
    }
    
    // CurrentPeriod used by Billing
    type CurrentPeriod struct {

        
            StartDate string  `json:"start_date"`
            EndDate string  `json:"end_date"`
         
    }
    
    // SubscriptionCharge used by Billing
    type SubscriptionCharge struct {

        
            ID string  `json:"_id"`
            Name string  `json:"name"`
            Term string  `json:"term"`
            PricingType string  `json:"pricing_type"`
            Price EntityChargePrice  `json:"price"`
            Recurring EntityChargeRecurring  `json:"recurring"`
            CappedAmount float64  `json:"capped_amount"`
            ActivatedOn string  `json:"activated_on"`
            CancelledOn string  `json:"cancelled_on"`
            BillingDate string  `json:"billing_date"`
            CurrentPeriod CurrentPeriod  `json:"current_period"`
            Status string  `json:"status"`
            IsTest bool  `json:"is_test"`
            Metadata map[string]interface{}  `json:"metadata"`
         
    }
    
    // EntitySubscription used by Billing
    type EntitySubscription struct {

        
            ID string  `json:"_id"`
            Name string  `json:"name"`
            Status string  `json:"status"`
            CompanyID float64  `json:"company_id"`
            ActivatedOn string  `json:"activated_on"`
            CancelledOn string  `json:"cancelled_on"`
            TrialDays float64  `json:"trial_days"`
            TrialPeriod SubscriptionTrialPeriod  `json:"trial_period"`
            Metadata map[string]interface{}  `json:"metadata"`
            LineItems []SubscriptionCharge  `json:"line_items"`
         
    }
    
    // OneTimeChargeEntity used by Billing
    type OneTimeChargeEntity struct {

        
            ID string  `json:"_id"`
            Name string  `json:"name"`
            Status string  `json:"status"`
            ActivatedOn string  `json:"activated_on"`
            CancelledOn string  `json:"cancelled_on"`
            Metadata map[string]interface{}  `json:"metadata"`
            ReturnURL string  `json:"return_url"`
            IsTest bool  `json:"is_test"`
            PricingType string  `json:"pricing_type"`
            SubscriberID string  `json:"subscriber_id"`
            EntityType string  `json:"entity_type"`
            EntityID string  `json:"entity_id"`
            Meta map[string]interface{}  `json:"meta"`
            Price EntityChargePrice  `json:"price"`
         
    }
    
    // CreateOneTimeChargeResponse used by Billing
    type CreateOneTimeChargeResponse struct {

        
            Charge OneTimeChargeEntity  `json:"charge"`
            ConfirmURL string  `json:"confirm_url"`
         
    }
    
    // CreateSubscriptionResponse used by Billing
    type CreateSubscriptionResponse struct {

        
            Subscription EntitySubscription  `json:"subscription"`
            ConfirmURL string  `json:"confirm_url"`
         
    }
    
    // InvoiceDetailsPeriod used by Billing
    type InvoiceDetailsPeriod struct {

        
            Start string  `json:"start"`
            End string  `json:"end"`
         
    }
    
    // InvoiceDetailsClient used by Billing
    type InvoiceDetailsClient struct {

        
            AddressLines []string  `json:"address_lines"`
            Name string  `json:"name"`
            Email string  `json:"email"`
            Phone string  `json:"phone"`
         
    }
    
    // InvoiceDetailsStatusTrail used by Billing
    type InvoiceDetailsStatusTrail struct {

        
            ID string  `json:"_id"`
            Value string  `json:"value"`
            Timestamp string  `json:"timestamp"`
         
    }
    
    // InvoiceDetailsPaymentMethodsDataChecks used by Billing
    type InvoiceDetailsPaymentMethodsDataChecks struct {

        
            CvcCheck string  `json:"cvc_check"`
            AddressLine1Check string  `json:"address_line1_check"`
            AddressPostalCodeCheck string  `json:"address_postal_code_check"`
         
    }
    
    // InvoiceDetailsPaymentMethodsDataNetworks used by Billing
    type InvoiceDetailsPaymentMethodsDataNetworks struct {

        
            Available []string  `json:"available"`
            Preferred string  `json:"preferred"`
         
    }
    
    // InvoiceDetailsPaymentMethodsDataThreeDSecureUsage used by Billing
    type InvoiceDetailsPaymentMethodsDataThreeDSecureUsage struct {

        
            Supported bool  `json:"supported"`
         
    }
    
    // InvoiceDetailsPaymentMethodsData used by Billing
    type InvoiceDetailsPaymentMethodsData struct {

        
            Brand string  `json:"brand"`
            Last4 string  `json:"last4"`
            Checks InvoiceDetailsPaymentMethodsDataChecks  `json:"checks"`
            Wallet string  `json:"wallet"`
            Country string  `json:"country"`
            Funding string  `json:"funding"`
            ExpYear float64  `json:"exp_year"`
            Networks InvoiceDetailsPaymentMethodsDataNetworks  `json:"networks"`
            ExpMonth float64  `json:"exp_month"`
            Fingerprint string  `json:"fingerprint"`
            GeneratedFrom string  `json:"generated_from"`
            ThreeDSecureUsage InvoiceDetailsPaymentMethodsDataThreeDSecureUsage  `json:"three_d_secure_usage"`
         
    }
    
    // InvoiceDetailsPaymentMethods used by Billing
    type InvoiceDetailsPaymentMethods struct {

        
            ID float64  `json:"id"`
            Type string  `json:"type"`
            PgPaymentMethodID string  `json:"pg_payment_method_id"`
            Data InvoiceDetailsPaymentMethodsData  `json:"data"`
            IsDefault bool  `json:"is_default"`
         
    }
    
    // InvoicePaymentMethod used by Billing
    type InvoicePaymentMethod struct {

        
            PgPaymentMethodID string  `json:"pg_payment_method_id"`
         
    }
    
    // InvoiceDetails used by Billing
    type InvoiceDetails struct {

        
            Period InvoiceDetailsPeriod  `json:"period"`
            Client InvoiceDetailsClient  `json:"client"`
            AutoAdvance bool  `json:"auto_advance"`
            Currency string  `json:"currency"`
            Paid bool  `json:"paid"`
            Attemp float64  `json:"attemp"`
            ID string  `json:"_id"`
            CollectionMethod string  `json:"collection_method"`
            SubscriberID string  `json:"subscriber_id"`
            InvoiceURL string  `json:"invoice_url"`
            Number string  `json:"number"`
            PgData map[string]interface{}  `json:"pg_data"`
            ReceiptNumber string  `json:"receipt_number"`
            StatementDescriptor string  `json:"statement_descriptor"`
            CurrentStatus string  `json:"current_status"`
            StatusTrail []InvoiceDetailsStatusTrail  `json:"status_trail"`
            Subtotal float64  `json:"subtotal"`
            Total float64  `json:"total"`
            Subscription string  `json:"subscription"`
            NextActionTime string  `json:"next_action_time"`
            CreatedAt string  `json:"created_at"`
            ModifiedAt string  `json:"modified_at"`
            HashIdentifier string  `json:"hash_identifier"`
            PaymentMethod InvoicePaymentMethod  `json:"payment_method"`
         
    }
    
    // InvoiceItemsPlanRecurring used by Billing
    type InvoiceItemsPlanRecurring struct {

        
            Interval string  `json:"interval"`
            IntervalCount float64  `json:"interval_count"`
         
    }
    
    // InvoiceItemsPlan used by Billing
    type InvoiceItemsPlan struct {

        
            Recurring InvoiceItemsPlanRecurring  `json:"recurring"`
            IsTrialPlan bool  `json:"is_trial_plan"`
            PlanGroup string  `json:"plan_group"`
            TagLines []string  `json:"tag_lines"`
            Currency string  `json:"currency"`
            IsActive bool  `json:"is_active"`
            IsVisible bool  `json:"is_visible"`
            TrialPeriod float64  `json:"trial_period"`
            Addons []string  `json:"addons"`
            Tags []string  `json:"tags"`
            Type string  `json:"type"`
            Country string  `json:"country"`
            ID string  `json:"_id"`
            Name string  `json:"name"`
            Description string  `json:"description"`
            Amount float64  `json:"amount"`
            ProductSuiteID string  `json:"product_suite_id"`
            CreatedAt string  `json:"created_at"`
            ModifiedAt string  `json:"modified_at"`
         
    }
    
    // InvoiceItemsPeriod used by Billing
    type InvoiceItemsPeriod struct {

        
            Start string  `json:"start"`
            End string  `json:"end"`
         
    }
    
    // InvoiceItems used by Billing
    type InvoiceItems struct {

        
            ID string  `json:"_id"`
            Currency string  `json:"currency"`
            Plan InvoiceItemsPlan  `json:"plan"`
            Name string  `json:"name"`
            Quantity float64  `json:"quantity"`
            Description string  `json:"description"`
            Period InvoiceItemsPeriod  `json:"period"`
            UnitAmount float64  `json:"unit_amount"`
            Amount float64  `json:"amount"`
            Type string  `json:"type"`
            InvoiceID string  `json:"invoice_id"`
            CreatedAt string  `json:"created_at"`
            ModifiedAt string  `json:"modified_at"`
         
    }
    
    // Invoice used by Billing
    type Invoice struct {

        
            Invoice InvoiceDetails  `json:"invoice"`
            InvoiceItems []InvoiceItems  `json:"invoice_items"`
         
    }
    
    // InvoicesDataClient used by Billing
    type InvoicesDataClient struct {

        
            Name string  `json:"name"`
            Email string  `json:"email"`
            Phone string  `json:"phone"`
            AddressLines []string  `json:"address_lines"`
         
    }
    
    // InvoicesDataPeriod used by Billing
    type InvoicesDataPeriod struct {

        
            Start string  `json:"start"`
            End string  `json:"end"`
         
    }
    
    // InvoicesDataPaymentMethod used by Billing
    type InvoicesDataPaymentMethod struct {

        
            PgPaymentMethodID string  `json:"pg_payment_method_id"`
         
    }
    
    // InvoicesData used by Billing
    type InvoicesData struct {

        
            ID string  `json:"_id"`
            Client InvoicesDataClient  `json:"client"`
            AutoAdvance bool  `json:"auto_advance"`
            Currency string  `json:"currency"`
            Paid bool  `json:"paid"`
            Attemp float64  `json:"attemp"`
            CollectionMethod string  `json:"collection_method"`
            SubscriberID string  `json:"subscriber_id"`
            InvoiceURL string  `json:"invoice_url"`
            Number string  `json:"number"`
            PgData map[string]interface{}  `json:"pg_data"`
            Period InvoicesDataPeriod  `json:"period"`
            ReceiptNumber string  `json:"receipt_number"`
            StatementDescriptor string  `json:"statement_descriptor"`
            CurrentStatus string  `json:"current_status"`
            StatusTrail []InvoiceDetailsStatusTrail  `json:"status_trail"`
            Subtotal float64  `json:"subtotal"`
            Total float64  `json:"total"`
            Subscription string  `json:"subscription"`
            NextActionTime string  `json:"next_action_time"`
            CreatedAt string  `json:"created_at"`
            ModifiedAt string  `json:"modified_at"`
            HashIdentifier string  `json:"hash_identifier"`
            PaymentMethod InvoicesDataPaymentMethod  `json:"payment_method"`
            InvoiceItems []InvoiceItems  `json:"invoice_items"`
         
    }
    
    // Invoices used by Billing
    type Invoices struct {

        
            Data []InvoicesData  `json:"data"`
            Start float64  `json:"start"`
            End float64  `json:"end"`
            Limit float64  `json:"limit"`
            Page float64  `json:"page"`
            Total float64  `json:"total"`
         
    }
    
    // Phone used by Billing
    type Phone struct {

        
            PhoneNumber string  `json:"phone_number"`
            PhoneCountryCode string  `json:"phone_country_code"`
         
    }
    
    // SubscriptionBillingAddress used by Billing
    type SubscriptionBillingAddress struct {

        
            Country string  `json:"country"`
            State string  `json:"state"`
            City string  `json:"city"`
            Line1 string  `json:"line1"`
            Line2 string  `json:"line2"`
            PostalCode string  `json:"postal_code"`
         
    }
    
    // SubscriptionCustomer used by Billing
    type SubscriptionCustomer struct {

        
            Phone Phone  `json:"phone"`
            BillingAddress SubscriptionBillingAddress  `json:"billing_address"`
            ID string  `json:"_id"`
            UniqueID string  `json:"unique_id"`
            Type string  `json:"type"`
            Name string  `json:"name"`
            Email string  `json:"email"`
            CreatedAt string  `json:"created_at"`
            ModifiedAt string  `json:"modified_at"`
            CreditBalance float64  `json:"credit_balance"`
            Data map[string]interface{}  `json:"data"`
         
    }
    
    // SubscriptionCustomerCreate used by Billing
    type SubscriptionCustomerCreate struct {

        
            Phone Phone  `json:"phone"`
            BillingAddress SubscriptionBillingAddress  `json:"billing_address"`
            UniqueID string  `json:"unique_id"`
            Type string  `json:"type"`
            Name string  `json:"name"`
            Email string  `json:"email"`
         
    }
    
    // SubscriptionCurrentPeriod used by Billing
    type SubscriptionCurrentPeriod struct {

        
            Start string  `json:"start"`
            End string  `json:"end"`
         
    }
    
    // SubscriptionPauseCollection used by Billing
    type SubscriptionPauseCollection struct {

        
            Behavior string  `json:"behavior"`
            ResumeAt string  `json:"resume_at"`
         
    }
    
    // SubscriptionTrial used by Billing
    type SubscriptionTrial struct {

        
            Start string  `json:"start"`
            End string  `json:"end"`
         
    }
    
    // SubscriptionInvoiceSettings used by Billing
    type SubscriptionInvoiceSettings struct {

        
            Generation bool  `json:"generation"`
            Charging bool  `json:"charging"`
         
    }
    
    // Subscription used by Billing
    type Subscription struct {

        
            CurrentPeriod SubscriptionCurrentPeriod  `json:"current_period"`
            PauseCollection SubscriptionPauseCollection  `json:"pause_collection"`
            Trial SubscriptionTrial  `json:"trial"`
            InvoiceSettings SubscriptionInvoiceSettings  `json:"invoice_settings"`
            IsActive bool  `json:"is_active"`
            CancelAtPeriodEnd bool  `json:"cancel_at_period_end"`
            ID string  `json:"_id"`
            SubscriberID string  `json:"subscriber_id"`
            PlanID string  `json:"plan_id"`
            ProductSuiteID string  `json:"product_suite_id"`
            PlanData Plan  `json:"plan_data"`
            CurrentStatus string  `json:"current_status"`
            CollectionMethod string  `json:"collection_method"`
            CreatedAt string  `json:"created_at"`
            ModifiedAt string  `json:"modified_at"`
            LatestInvoice string  `json:"latest_invoice"`
         
    }
    
    // SubscriptionStatus used by Billing
    type SubscriptionStatus struct {

        
            IsEnabled bool  `json:"is_enabled"`
            Subscription Subscription  `json:"subscription"`
         
    }
    
    // SubscriptionLimitApplication used by Billing
    type SubscriptionLimitApplication struct {

        
            Enabled bool  `json:"enabled"`
            HardLimit float64  `json:"hard_limit"`
            SoftLimit float64  `json:"soft_limit"`
         
    }
    
    // SubscriptionLimitMarketplace used by Billing
    type SubscriptionLimitMarketplace struct {

        
            Enabled bool  `json:"enabled"`
         
    }
    
    // SubscriptionLimitOtherPlatform used by Billing
    type SubscriptionLimitOtherPlatform struct {

        
            Enabled bool  `json:"enabled"`
         
    }
    
    // SubscriptionLimitTeam used by Billing
    type SubscriptionLimitTeam struct {

        
            Limit float64  `json:"limit"`
         
    }
    
    // SubscriptionLimitProducts used by Billing
    type SubscriptionLimitProducts struct {

        
            Bulk bool  `json:"bulk"`
            Limit float64  `json:"limit"`
         
    }
    
    // SubscriptionLimitExtensions used by Billing
    type SubscriptionLimitExtensions struct {

        
            Enabled bool  `json:"enabled"`
            Limit float64  `json:"limit"`
         
    }
    
    // SubscriptionLimitIntegrations used by Billing
    type SubscriptionLimitIntegrations struct {

        
            Enabled bool  `json:"enabled"`
            Limit float64  `json:"limit"`
         
    }
    
    // SubscriptionLimit used by Billing
    type SubscriptionLimit struct {

        
            Application SubscriptionLimitApplication  `json:"application"`
            Marketplace SubscriptionLimitMarketplace  `json:"marketplace"`
            OtherPlatform SubscriptionLimitOtherPlatform  `json:"other_platform"`
            Team SubscriptionLimitTeam  `json:"team"`
            Products SubscriptionLimitProducts  `json:"products"`
            Extensions SubscriptionLimitExtensions  `json:"extensions"`
            Integrations SubscriptionLimitIntegrations  `json:"integrations"`
            IsTrialPlan bool  `json:"is_trial_plan"`
         
    }
    
    // SubscriptionActivateReq used by Billing
    type SubscriptionActivateReq struct {

        
            UniqueID string  `json:"unique_id"`
            Type string  `json:"type"`
            ProductSuite string  `json:"product_suite"`
            PlanID string  `json:"plan_id"`
            PaymentMethod string  `json:"payment_method"`
         
    }
    
    // SubscriptionActivateRes used by Billing
    type SubscriptionActivateRes struct {

        
            Success bool  `json:"success"`
            Data Subscription  `json:"data"`
         
    }
    
    // CancelSubscriptionReq used by Billing
    type CancelSubscriptionReq struct {

        
            UniqueID string  `json:"unique_id"`
            Type string  `json:"type"`
            ProductSuite string  `json:"product_suite"`
            SubscriptionID string  `json:"subscription_id"`
         
    }
    
    // CancelSubscriptionRes used by Billing
    type CancelSubscriptionRes struct {

        
            Success bool  `json:"success"`
            Data Subscription  `json:"data"`
         
    }
    

    
    // StatsImported used by Communication
    type StatsImported struct {

        
            Count float64  `json:"count"`
         
    }
    
    // StatsProcessedEmail used by Communication
    type StatsProcessedEmail struct {

        
            Success float64  `json:"success"`
            Failed float64  `json:"failed"`
            Suppressed float64  `json:"suppressed"`
         
    }
    
    // StatsProcessedSms used by Communication
    type StatsProcessedSms struct {

        
            Success float64  `json:"success"`
            Failed float64  `json:"failed"`
            Suppressed float64  `json:"suppressed"`
         
    }
    
    // StatsProcessed used by Communication
    type StatsProcessed struct {

        
            Email StatsProcessedEmail  `json:"email"`
            Sms StatsProcessedSms  `json:"sms"`
         
    }
    
    // Stats used by Communication
    type Stats struct {

        
            ID string  `json:"_id"`
            Imported interface{}  `json:"imported"`
            Processed interface{}  `json:"processed"`
         
    }
    
    // GetStats used by Communication
    type GetStats struct {

        
            Items []Stats  `json:"items"`
         
    }
    
    // CampaignReq used by Communication
    type CampaignReq struct {

        
            Description string  `json:"description"`
            Tags []string  `json:"tags"`
            Headers []string  `json:"headers"`
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            FileURL string  `json:"file_url"`
            Type string  `json:"type"`
            RecordsCount float64  `json:"records_count"`
            Application string  `json:"application"`
         
    }
    
    // RecipientHeaders used by Communication
    type RecipientHeaders struct {

        
            Email string  `json:"email"`
         
    }
    
    // CampaignEmailTemplate used by Communication
    type CampaignEmailTemplate struct {

        
            Key string  `json:"key"`
            Value string  `json:"value"`
         
    }
    
    // CampignEmailProvider used by Communication
    type CampignEmailProvider struct {

        
            ID string  `json:"_id"`
            FromName string  `json:"from_name"`
            FromEmail string  `json:"from_email"`
         
    }
    
    // CampaignEmail used by Communication
    type CampaignEmail struct {

        
            Template CampaignEmailTemplate  `json:"template"`
            Provider CampignEmailProvider  `json:"provider"`
         
    }
    
    // Campaign used by Communication
    type Campaign struct {

        
            RecipientHeaders RecipientHeaders  `json:"recipient_headers"`
            Email CampaignEmail  `json:"email"`
            Description string  `json:"description"`
            Tags []interface{}  `json:"tags"`
            IsActive bool  `json:"is_active"`
            ID string  `json:"_id"`
            Datasource string  `json:"datasource"`
            Type string  `json:"type"`
            Name string  `json:"name"`
            Application string  `json:"application"`
            CreatedAt string  `json:"created_at"`
            UpdatedAt string  `json:"updated_at"`
            Slug string  `json:"slug"`
            V float64  `json:"__v"`
         
    }
    
    // Campaigns used by Communication
    type Campaigns struct {

        
            Items []Campaign  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // BadRequestSchema used by Communication
    type BadRequestSchema struct {

        
            Status string  `json:"status"`
            Message string  `json:"message"`
         
    }
    
    // NotFound used by Communication
    type NotFound struct {

        
            Message string  `json:"message"`
         
    }
    
    // BigqueryHeadersReq used by Communication
    type BigqueryHeadersReq struct {

        
            Query string  `json:"query"`
            Type string  `json:"type"`
         
    }
    
    // BigqueryHeadersResHeaders used by Communication
    type BigqueryHeadersResHeaders struct {

        
            Name string  `json:"name"`
            Type string  `json:"type"`
         
    }
    
    // BigqueryHeadersRes used by Communication
    type BigqueryHeadersRes struct {

        
            Headers []BigqueryHeadersResHeaders  `json:"headers"`
         
    }
    
    // GetNRecordsCsvReq used by Communication
    type GetNRecordsCsvReq struct {

        
            URL string  `json:"url"`
            Header bool  `json:"header"`
            Count float64  `json:"count"`
         
    }
    
    // GetNRecordsCsvResItems used by Communication
    type GetNRecordsCsvResItems struct {

        
            PhoneNumber string  `json:"phone_number"`
            Email string  `json:"email"`
            Firstname string  `json:"firstname"`
            Lastname string  `json:"lastname"`
            Orderid string  `json:"orderid"`
         
    }
    
    // GetNRecordsCsvRes used by Communication
    type GetNRecordsCsvRes struct {

        
            Items []GetNRecordsCsvResItems  `json:"items"`
         
    }
    
    // AudienceReq used by Communication
    type AudienceReq struct {

        
            Description string  `json:"description"`
            Tags []string  `json:"tags"`
            Headers []string  `json:"headers"`
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            FileURL string  `json:"file_url"`
            Type string  `json:"type"`
            RecordsCount float64  `json:"records_count"`
            Application string  `json:"application"`
         
    }
    
    // Audience used by Communication
    type Audience struct {

        
            Description string  `json:"description"`
            Tags []string  `json:"tags"`
            Headers []string  `json:"headers"`
            IsActive bool  `json:"is_active"`
            ID string  `json:"_id"`
            Name string  `json:"name"`
            FileURL string  `json:"file_url"`
            Type string  `json:"type"`
            RecordsCount float64  `json:"records_count"`
            Application string  `json:"application"`
            CreatedAt string  `json:"created_at"`
            UpdatedAt string  `json:"updated_at"`
            Slug string  `json:"slug"`
            V float64  `json:"__v"`
         
    }
    
    // Audiences used by Communication
    type Audiences struct {

        
            Items []Audience  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // EmailProviderReqFrom used by Communication
    type EmailProviderReqFrom struct {

        
            Name string  `json:"name"`
            Email string  `json:"email"`
            IsDefault bool  `json:"is_default"`
         
    }
    
    // EmailProviderReq used by Communication
    type EmailProviderReq struct {

        
            Name string  `json:"name"`
            Description string  `json:"description"`
            APIKey string  `json:"api_key"`
            Type string  `json:"type"`
            Provider string  `json:"provider"`
            FromAddress []EmailProviderReqFrom  `json:"from_address"`
         
    }
    
    // EmailProvider used by Communication
    type EmailProvider struct {

        
            Type string  `json:"type"`
            Provider string  `json:"provider"`
            FromAddress []EmailProviderReqFrom  `json:"from_address"`
            ID string  `json:"_id"`
            Name string  `json:"name"`
            Description string  `json:"description"`
            APIKey string  `json:"api_key"`
            Application string  `json:"application"`
            CreatedAt string  `json:"created_at"`
            UpdatedAt string  `json:"updated_at"`
            Slug string  `json:"slug"`
            V float64  `json:"__v"`
         
    }
    
    // EmailProviders used by Communication
    type EmailProviders struct {

        
            Items []EmailProvider  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // EmailTemplateDeleteSuccessRes used by Communication
    type EmailTemplateDeleteSuccessRes struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // EmailTemplateDeleteFailureRes used by Communication
    type EmailTemplateDeleteFailureRes struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // EmailTemplateKeys used by Communication
    type EmailTemplateKeys struct {

        
            To string  `json:"to"`
            Cc string  `json:"cc"`
            Bcc string  `json:"bcc"`
         
    }
    
    // EmailTemplateHeaders used by Communication
    type EmailTemplateHeaders struct {

        
            Key string  `json:"key"`
            Value string  `json:"value"`
         
    }
    
    // EmailTemplateReq used by Communication
    type EmailTemplateReq struct {

        
            Name string  `json:"name"`
            Description string  `json:"description"`
            Keys EmailTemplateKeys  `json:"keys"`
            StaticTo []string  `json:"static_to"`
            StaticCc []string  `json:"static_cc"`
            StaticBcc []string  `json:"static_bcc"`
            ReplyTo string  `json:"reply_to"`
            Headers []EmailTemplateHeaders  `json:"headers"`
            Subject TemplateAndType  `json:"subject"`
            Html TemplateAndType  `json:"html"`
            Text TemplateAndType  `json:"text"`
            Attachments []interface{}  `json:"attachments"`
            Priority string  `json:"priority"`
         
    }
    
    // TemplateAndType used by Communication
    type TemplateAndType struct {

        
            TemplateType string  `json:"template_type"`
            Template string  `json:"template"`
         
    }
    
    // EmailTemplateRes used by Communication
    type EmailTemplateRes struct {

        
            IsSystem bool  `json:"is_system"`
            IsInternal bool  `json:"is_internal"`
            Description string  `json:"description"`
            StaticTo []string  `json:"static_to"`
            StaticCc []string  `json:"static_cc"`
            StaticBcc []string  `json:"static_bcc"`
            Tags []interface{}  `json:"tags"`
            Priority string  `json:"priority"`
            Published bool  `json:"published"`
            ID string  `json:"_id"`
            Name string  `json:"name"`
            Keys EmailTemplateKeys  `json:"keys"`
            ReplyTo string  `json:"reply_to"`
            Headers []EmailTemplateHeaders  `json:"headers"`
            Subject TemplateAndType  `json:"subject"`
            Html TemplateAndType  `json:"html"`
            Text TemplateAndType  `json:"text"`
            Attachments []interface{}  `json:"attachments"`
            CreatedAt string  `json:"created_at"`
            UpdatedAt string  `json:"updated_at"`
            Slug string  `json:"slug"`
            V float64  `json:"__v"`
         
    }
    
    // EmailTemplate used by Communication
    type EmailTemplate struct {

        
            IsSystem bool  `json:"is_system"`
            IsInternal bool  `json:"is_internal"`
            Description string  `json:"description"`
            StaticTo []interface{}  `json:"static_to"`
            StaticCc []interface{}  `json:"static_cc"`
            StaticBcc []interface{}  `json:"static_bcc"`
            Tags []interface{}  `json:"tags"`
            Priority string  `json:"priority"`
            Published bool  `json:"published"`
            ID string  `json:"_id"`
            Slug string  `json:"slug"`
            Name string  `json:"name"`
            FromName string  `json:"from_name"`
            Subject TemplateAndType  `json:"subject"`
            Html TemplateAndType  `json:"html"`
            Text TemplateAndType  `json:"text"`
            Headers []interface{}  `json:"headers"`
            Attachments []interface{}  `json:"attachments"`
            CreatedAt string  `json:"created_at"`
            UpdatedAt string  `json:"updated_at"`
            V float64  `json:"__v"`
         
    }
    
    // SystemEmailTemplate used by Communication
    type SystemEmailTemplate struct {

        
            IsSystem bool  `json:"is_system"`
            IsInternal bool  `json:"is_internal"`
            Description string  `json:"description"`
            StaticTo []interface{}  `json:"static_to"`
            StaticCc []interface{}  `json:"static_cc"`
            StaticBcc []interface{}  `json:"static_bcc"`
            Tags []interface{}  `json:"tags"`
            Priority string  `json:"priority"`
            Published bool  `json:"published"`
            ID string  `json:"_id"`
            Slug string  `json:"slug"`
            Name string  `json:"name"`
            FromName string  `json:"from_name"`
            Subject TemplateAndType  `json:"subject"`
            Html TemplateAndType  `json:"html"`
            Text TemplateAndType  `json:"text"`
            Headers []interface{}  `json:"headers"`
            Attachments []interface{}  `json:"attachments"`
            CreatedAt string  `json:"created_at"`
            UpdatedAt string  `json:"updated_at"`
            V float64  `json:"__v"`
         
    }
    
    // EmailTemplates used by Communication
    type EmailTemplates struct {

        
            Items []EmailTemplate  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // SystemEmailTemplates used by Communication
    type SystemEmailTemplates struct {

        
            Items []SystemEmailTemplate  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // PayloadEmailTemplateStructure used by Communication
    type PayloadEmailTemplateStructure struct {

        
            Key string  `json:"key"`
            Value interface{}  `json:"value"`
         
    }
    
    // PayloadEmailProviderStructure used by Communication
    type PayloadEmailProviderStructure struct {

        
            ID string  `json:"_id"`
         
    }
    
    // PayloadEmailStructure used by Communication
    type PayloadEmailStructure struct {

        
            Template PayloadEmailTemplateStructure  `json:"template"`
            Provider PayloadEmailProviderStructure  `json:"provider"`
         
    }
    
    // PayloadSmsTemplateStructure used by Communication
    type PayloadSmsTemplateStructure struct {

        
            Key string  `json:"key"`
            Value interface{}  `json:"value"`
         
    }
    
    // PayloadSmsProviderStructure used by Communication
    type PayloadSmsProviderStructure struct {

        
            ID string  `json:"_id"`
         
    }
    
    // PayloadSmsStructure used by Communication
    type PayloadSmsStructure struct {

        
            Template PayloadSmsTemplateStructure  `json:"template"`
            Provider PayloadSmsProviderStructure  `json:"provider"`
         
    }
    
    // PayloadStructure used by Communication
    type PayloadStructure struct {

        
            Data []map[string]interface{}  `json:"data"`
            Email PayloadEmailStructure  `json:"email"`
            Sms PayloadSmsStructure  `json:"sms"`
            Application string  `json:"application"`
         
    }
    
    // MetaStructure used by Communication
    type MetaStructure struct {

        
            JobType string  `json:"job_type"`
            Action string  `json:"action"`
            Trace string  `json:"trace"`
            Timestamp string  `json:"timestamp"`
         
    }
    
    // EngineRequest used by Communication
    type EngineRequest struct {

        
            Payload PayloadStructure  `json:"payload"`
            Meta MetaStructure  `json:"meta"`
         
    }
    
    // EngineResponse used by Communication
    type EngineResponse struct {

        
            Success bool  `json:"success"`
         
    }
    
    // EventSubscriptionTemplateSms used by Communication
    type EventSubscriptionTemplateSms struct {

        
            Subscribed bool  `json:"subscribed"`
            Template string  `json:"template"`
         
    }
    
    // EventSubscriptionTemplateEmail used by Communication
    type EventSubscriptionTemplateEmail struct {

        
            Subscribed bool  `json:"subscribed"`
            Template string  `json:"template"`
         
    }
    
    // EventSubscriptionTemplate used by Communication
    type EventSubscriptionTemplate struct {

        
            Sms EventSubscriptionTemplateSms  `json:"sms"`
            Email EventSubscriptionTemplateEmail  `json:"email"`
         
    }
    
    // EventSubscription used by Communication
    type EventSubscription struct {

        
            Template EventSubscriptionTemplate  `json:"template"`
            IsDefault bool  `json:"is_default"`
            ID string  `json:"_id"`
            Application string  `json:"application"`
            Event string  `json:"event"`
            Slug string  `json:"slug"`
            CreatedAt string  `json:"created_at"`
            UpdatedAt string  `json:"updated_at"`
            V float64  `json:"__v"`
         
    }
    
    // EventSubscriptions used by Communication
    type EventSubscriptions struct {

        
            Items []EventSubscription  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // TriggerJobResponse used by Communication
    type TriggerJobResponse struct {

        
            Status float64  `json:"status"`
         
    }
    
    // TriggerJobRequest used by Communication
    type TriggerJobRequest struct {

        
            JobID string  `json:"job_id"`
         
    }
    
    // Job used by Communication
    type Job struct {

        
            Completed bool  `json:"completed"`
            IsActive bool  `json:"is_active"`
            ID string  `json:"_id"`
            Campaign string  `json:"campaign"`
            Application string  `json:"application"`
            CreatedAt string  `json:"created_at"`
            UpdatedAt string  `json:"updated_at"`
            V float64  `json:"__v"`
         
    }
    
    // Jobs used by Communication
    type Jobs struct {

        
            Items []Job  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // JobLog used by Communication
    type JobLog struct {

        
            Imported interface{}  `json:"imported"`
            Processed interface{}  `json:"processed"`
            ID string  `json:"_id"`
            Job string  `json:"job"`
            Campaign string  `json:"campaign"`
            CreatedAt string  `json:"created_at"`
            UpdatedAt string  `json:"updated_at"`
            V float64  `json:"__v"`
         
    }
    
    // JobLogs used by Communication
    type JobLogs struct {

        
            Items []JobLog  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // LogEmail used by Communication
    type LogEmail struct {

        
            Template string  `json:"template"`
         
    }
    
    // LogPushnotification used by Communication
    type LogPushnotification struct {

        
            Pushtokens []string  `json:"pushtokens"`
         
    }
    
    // LogMeta used by Communication
    type LogMeta struct {

        
            Type string  `json:"type"`
            Identifier string  `json:"identifier"`
            Key string  `json:"key"`
            Offset string  `json:"offset"`
            Partition string  `json:"partition"`
            Topic string  `json:"topic"`
         
    }
    
    // Log used by Communication
    type Log struct {

        
            Email LogEmail  `json:"email"`
            Pushnotification LogPushnotification  `json:"pushnotification"`
            Meta LogMeta  `json:"meta"`
            ID string  `json:"_id"`
            Application string  `json:"application"`
            Service string  `json:"service"`
            Step string  `json:"step"`
            Status string  `json:"status"`
            Data interface{}  `json:"data"`
            ExpireAt string  `json:"expire_at"`
            CreatedAt string  `json:"created_at"`
         
    }
    
    // Logs used by Communication
    type Logs struct {

        
            Items []Log  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // SendOtpSmsCommsTemplate used by Communication
    type SendOtpSmsCommsTemplate struct {

        
            Key string  `json:"key"`
            Value interface{}  `json:"value"`
         
    }
    
    // SendOtpSmsCommsProvider used by Communication
    type SendOtpSmsCommsProvider struct {

        
            Slug string  `json:"slug"`
            ID string  `json:"_id"`
         
    }
    
    // SendOtpEmailCommsTemplate used by Communication
    type SendOtpEmailCommsTemplate struct {

        
            Key string  `json:"key"`
            Value interface{}  `json:"value"`
         
    }
    
    // SendOtpCommsReqData used by Communication
    type SendOtpCommsReqData struct {

        
            SendSameOtpToChannel bool  `json:"send_same_otp_to_channel"`
            Mobile string  `json:"mobile"`
            CountryCode string  `json:"country_code"`
            To string  `json:"to"`
         
    }
    
    // SendOtpCommsReqSms used by Communication
    type SendOtpCommsReqSms struct {

        
            OtpLength float64  `json:"otp_length"`
            Expiry float64  `json:"expiry"`
            Template SendOtpSmsCommsTemplate  `json:"template"`
            Provider SendOtpSmsCommsProvider  `json:"provider"`
         
    }
    
    // SendOtpCommsReqEmail used by Communication
    type SendOtpCommsReqEmail struct {

        
            OtpLength float64  `json:"otp_length"`
            Expiry float64  `json:"expiry"`
            Template SendOtpEmailCommsTemplate  `json:"template"`
         
    }
    
    // SendOtpCommsResSms used by Communication
    type SendOtpCommsResSms struct {

        
            Success bool  `json:"success"`
            RequestID string  `json:"request_id"`
            Message string  `json:"message"`
            Mobile string  `json:"mobile"`
            CountryCode string  `json:"country_code"`
            ResendTimer float64  `json:"resend_timer"`
         
    }
    
    // SendOtpCommsResEmail used by Communication
    type SendOtpCommsResEmail struct {

        
            Success bool  `json:"success"`
            RequestID string  `json:"request_id"`
            Message string  `json:"message"`
            To string  `json:"to"`
            ResendTimer float64  `json:"resend_timer"`
         
    }
    
    // SendOtpCommsReq used by Communication
    type SendOtpCommsReq struct {

        
            Data SendOtpCommsReqData  `json:"data"`
            Sms SendOtpCommsReqSms  `json:"sms"`
            Email SendOtpCommsReqEmail  `json:"email"`
         
    }
    
    // SendOtpCommsRes used by Communication
    type SendOtpCommsRes struct {

        
            Sms SendOtpCommsResSms  `json:"sms"`
            Email SendOtpCommsResEmail  `json:"email"`
         
    }
    
    // VerifyOtpCommsReq used by Communication
    type VerifyOtpCommsReq struct {

        
            RequestID string  `json:"request_id"`
            Otp string  `json:"otp"`
         
    }
    
    // VerifyOtpCommsSuccessRes used by Communication
    type VerifyOtpCommsSuccessRes struct {

        
            Success bool  `json:"success"`
            Mobile string  `json:"mobile"`
            CountryCode string  `json:"country_code"`
            Message string  `json:"message"`
         
    }
    
    // VerifyOtpCommsErrorRes used by Communication
    type VerifyOtpCommsErrorRes struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // PushtokenReq used by Communication
    type PushtokenReq struct {

        
            Action string  `json:"action"`
            BundleIdentifier string  `json:"bundle_identifier"`
            PushToken string  `json:"push_token"`
            UniqueDeviceID string  `json:"unique_device_id"`
         
    }
    
    // PushtokenRes used by Communication
    type PushtokenRes struct {

        
            ID string  `json:"_id"`
            BundleIdentifier string  `json:"bundle_identifier"`
            PushToken string  `json:"push_token"`
            UniqueDeviceID string  `json:"unique_device_id"`
            Type string  `json:"type"`
            Platform string  `json:"platform"`
            ApplicationID string  `json:"application_id"`
            UserID string  `json:"user_id"`
            CreatedAt string  `json:"created_at"`
            UpdatedAt string  `json:"updated_at"`
            ExpiredAt string  `json:"expired_at"`
         
    }
    
    // SmsProviderReq used by Communication
    type SmsProviderReq struct {

        
            Name string  `json:"name"`
            Description string  `json:"description"`
            Sender string  `json:"sender"`
            Username string  `json:"username"`
            Authkey string  `json:"authkey"`
            Type string  `json:"type"`
            Provider string  `json:"provider"`
         
    }
    
    // SmsProvider used by Communication
    type SmsProvider struct {

        
            Rpt float64  `json:"rpt"`
            Type string  `json:"type"`
            Provider string  `json:"provider"`
            ID string  `json:"_id"`
            Name string  `json:"name"`
            Description string  `json:"description"`
            Sender string  `json:"sender"`
            Username string  `json:"username"`
            Authkey string  `json:"authkey"`
            Application string  `json:"application"`
            CreatedAt string  `json:"created_at"`
            UpdatedAt string  `json:"updated_at"`
            Slug string  `json:"slug"`
            V float64  `json:"__v"`
         
    }
    
    // SmsProviders used by Communication
    type SmsProviders struct {

        
            Items []SmsProvider  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // SmsTemplateDeleteSuccessRes used by Communication
    type SmsTemplateDeleteSuccessRes struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // SmsTemplateDeleteFailureRes used by Communication
    type SmsTemplateDeleteFailureRes struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // SmsTemplateMessage used by Communication
    type SmsTemplateMessage struct {

        
            TemplateType string  `json:"template_type"`
            Template string  `json:"template"`
         
    }
    
    // SmsTemplateReq used by Communication
    type SmsTemplateReq struct {

        
            Name string  `json:"name"`
            Description string  `json:"description"`
            Message SmsTemplateMessage  `json:"message"`
            TemplateVariables interface{}  `json:"template_variables"`
            Attachments []interface{}  `json:"attachments"`
            Priority string  `json:"priority"`
         
    }
    
    // SmsTemplateRes used by Communication
    type SmsTemplateRes struct {

        
            IsSystem bool  `json:"is_system"`
            IsInternal bool  `json:"is_internal"`
            Description string  `json:"description"`
            Tags []interface{}  `json:"tags"`
            Priority string  `json:"priority"`
            Published bool  `json:"published"`
            ID string  `json:"_id"`
            Name string  `json:"name"`
            Message SmsTemplateMessage  `json:"message"`
            TemplateVariables interface{}  `json:"template_variables"`
            CreatedAt string  `json:"created_at"`
            UpdatedAt string  `json:"updated_at"`
            Slug string  `json:"slug"`
            V float64  `json:"__v"`
         
    }
    
    // SmsTemplate used by Communication
    type SmsTemplate struct {

        
            IsSystem bool  `json:"is_system"`
            IsInternal bool  `json:"is_internal"`
            Description string  `json:"description"`
            Priority string  `json:"priority"`
            Tags []interface{}  `json:"tags"`
            Published bool  `json:"published"`
            ID string  `json:"_id"`
            Slug string  `json:"slug"`
            Name string  `json:"name"`
            Message SmsTemplateMessage  `json:"message"`
            TemplateVariables interface{}  `json:"template_variables"`
            CreatedAt string  `json:"created_at"`
            UpdatedAt string  `json:"updated_at"`
            V float64  `json:"__v"`
         
    }
    
    // SystemSmsTemplate used by Communication
    type SystemSmsTemplate struct {

        
            IsSystem bool  `json:"is_system"`
            IsInternal bool  `json:"is_internal"`
            Description string  `json:"description"`
            Tags []interface{}  `json:"tags"`
            Priority string  `json:"priority"`
            Published bool  `json:"published"`
            ID string  `json:"_id"`
            Slug string  `json:"slug"`
            Name string  `json:"name"`
            Message SmsTemplateMessage  `json:"message"`
            TemplateVariables interface{}  `json:"template_variables"`
            CreatedAt string  `json:"created_at"`
            UpdatedAt string  `json:"updated_at"`
            V float64  `json:"__v"`
         
    }
    
    // SmsTemplates used by Communication
    type SmsTemplates struct {

        
            Items []SmsTemplate  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // SystemSmsTemplates used by Communication
    type SystemSmsTemplates struct {

        
            Items []SystemSmsTemplate  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // Notification used by Communication
    type Notification struct {

        
            Title string  `json:"title"`
            Body string  `json:"body"`
            Subtitle string  `json:"subtitle"`
            Icon string  `json:"icon"`
            Deeplink string  `json:"deeplink"`
            ClickAction string  `json:"click_action"`
         
    }
    
    // SystemNotificationUser used by Communication
    type SystemNotificationUser struct {

        
            Type string  `json:"type"`
            Value string  `json:"value"`
         
    }
    
    // SystemNotificationSettings used by Communication
    type SystemNotificationSettings struct {

        
            Sound bool  `json:"sound"`
            Priority string  `json:"priority"`
            TimeToLive string  `json:"time_to_live"`
         
    }
    
    // SystemNotification used by Communication
    type SystemNotification struct {

        
            Notification Notification  `json:"notification"`
            User SystemNotificationUser  `json:"user"`
            Settings SystemNotificationUser  `json:"settings"`
            ID string  `json:"_id"`
            Group string  `json:"group"`
            CreatedAt string  `json:"created_at"`
         
    }
    
    // SystemNotificationsPage used by Communication
    type SystemNotificationsPage struct {

        
            Type string  `json:"type"`
            Current float64  `json:"current"`
            Size float64  `json:"size"`
            ItemTotal float64  `json:"item_total"`
            HasNext bool  `json:"has_next"`
         
    }
    
    // SystemNotifications used by Communication
    type SystemNotifications struct {

        
            Items []SystemNotification  `json:"items"`
            LastReadAnchor float64  `json:"last_read_anchor"`
            Page Page  `json:"page"`
         
    }
    
    // Page used by Communication
    type Page struct {

        
            ItemTotal float64  `json:"item_total"`
            NextID string  `json:"next_id"`
            HasPrevious bool  `json:"has_previous"`
            HasNext bool  `json:"has_next"`
            Current float64  `json:"current"`
            Type string  `json:"type"`
            Size float64  `json:"size"`
         
    }
    

    
    // PaymentGatewayConfigResponse used by Payment
    type PaymentGatewayConfigResponse struct {

        
            DisplayFields []string  `json:"display_fields"`
            Success bool  `json:"success"`
            Created bool  `json:"created"`
            Aggregators []map[string]interface{}  `json:"aggregators"`
            ExcludedFields []string  `json:"excluded_fields"`
            AppID string  `json:"app_id"`
         
    }
    
    // ErrorCodeDescription used by Payment
    type ErrorCodeDescription struct {

        
            Description string  `json:"description"`
            Success bool  `json:"success"`
            Code string  `json:"code"`
         
    }
    
    // PaymentGatewayConfig used by Payment
    type PaymentGatewayConfig struct {

        
            MerchantSalt string  `json:"merchant_salt"`
            IsActive bool  `json:"is_active"`
            Key string  `json:"key"`
            Secret string  `json:"secret"`
            ConfigType string  `json:"config_type"`
         
    }
    
    // PaymentGatewayConfigRequest used by Payment
    type PaymentGatewayConfigRequest struct {

        
            AggregatorName PaymentGatewayConfig  `json:"aggregator_name"`
            AppID string  `json:"app_id"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // PaymentGatewayToBeReviewed used by Payment
    type PaymentGatewayToBeReviewed struct {

        
            Aggregator []string  `json:"aggregator"`
            Success bool  `json:"success"`
         
    }
    
    // ErrorCodeAndDescription used by Payment
    type ErrorCodeAndDescription struct {

        
            Description string  `json:"description"`
            Code string  `json:"code"`
         
    }
    
    // HttpErrorCodeAndResponse used by Payment
    type HttpErrorCodeAndResponse struct {

        
            Success bool  `json:"success"`
            Error ErrorCodeAndDescription  `json:"error"`
         
    }
    
    // IntentAppErrorList used by Payment
    type IntentAppErrorList struct {

        
            PackageName string  `json:"package_name"`
            Code string  `json:"code"`
         
    }
    
    // PaymentModeLogo used by Payment
    type PaymentModeLogo struct {

        
            Small string  `json:"small"`
            Large string  `json:"large"`
         
    }
    
    // IntentApp used by Payment
    type IntentApp struct {

        
            DisplayName string  `json:"display_name"`
            PackageName string  `json:"package_name"`
            Logos PaymentModeLogo  `json:"logos"`
            Code string  `json:"code"`
         
    }
    
    // PaymentModeList used by Payment
    type PaymentModeList struct {

        
            DisplayPriority float64  `json:"display_priority"`
            Code string  `json:"code"`
            ExpMonth float64  `json:"exp_month"`
            RemainingLimit float64  `json:"remaining_limit"`
            CardBrandImage string  `json:"card_brand_image"`
            AggregatorName string  `json:"aggregator_name"`
            CardNumber string  `json:"card_number"`
            IntentAppErrorDictList []IntentAppErrorList  `json:"intent_app_error_dict_list"`
            CardIsin string  `json:"card_isin"`
            CardIssuer string  `json:"card_issuer"`
            DisplayName string  `json:"display_name"`
            IntentApp []IntentApp  `json:"intent_app"`
            IntentAppErrorList []string  `json:"intent_app_error_list"`
            CardReference string  `json:"card_reference"`
            MerchantCode string  `json:"merchant_code"`
            CodLimit float64  `json:"cod_limit"`
            LogoURL PaymentModeLogo  `json:"logo_url"`
            ExpYear float64  `json:"exp_year"`
            Expired bool  `json:"expired"`
            Nickname string  `json:"nickname"`
            CardID string  `json:"card_id"`
            Name string  `json:"name"`
            CodLimitPerOrder float64  `json:"cod_limit_per_order"`
            Timeout float64  `json:"timeout"`
            IntentFlow bool  `json:"intent_flow"`
            CardToken string  `json:"card_token"`
            CardName string  `json:"card_name"`
            FyndVpa string  `json:"fynd_vpa"`
            CompliantWithTokenisationGuidelines bool  `json:"compliant_with_tokenisation_guidelines"`
            CardType string  `json:"card_type"`
            RetryCount float64  `json:"retry_count"`
            CardBrand string  `json:"card_brand"`
            CardFingerprint string  `json:"card_fingerprint"`
         
    }
    
    // RootPaymentMode used by Payment
    type RootPaymentMode struct {

        
            AggregatorName string  `json:"aggregator_name"`
            DisplayPriority float64  `json:"display_priority"`
            DisplayName string  `json:"display_name"`
            IsPayByCardPl bool  `json:"is_pay_by_card_pl"`
            SaveCard bool  `json:"save_card"`
            Name string  `json:"name"`
            List []PaymentModeList  `json:"list"`
            AddCardEnabled bool  `json:"add_card_enabled"`
            AnonymousEnable bool  `json:"anonymous_enable"`
         
    }
    
    // PaymentOptions used by Payment
    type PaymentOptions struct {

        
            PaymentOption []RootPaymentMode  `json:"payment_option"`
         
    }
    
    // PaymentOptionsResponse used by Payment
    type PaymentOptionsResponse struct {

        
            PaymentOptions PaymentOptions  `json:"payment_options"`
            Success bool  `json:"success"`
         
    }
    
    // PayoutMoreAttributes used by Payment
    type PayoutMoreAttributes struct {

        
            Country string  `json:"country"`
            BankName string  `json:"bank_name"`
            IfscCode string  `json:"ifsc_code"`
            AccountNo string  `json:"account_no"`
            AccountType string  `json:"account_type"`
            State string  `json:"state"`
            City string  `json:"city"`
            AccountHolder string  `json:"account_holder"`
            BranchName string  `json:"branch_name"`
         
    }
    
    // PayoutAggregator used by Payment
    type PayoutAggregator struct {

        
            AggregatorFundID float64  `json:"aggregator_fund_id"`
            AggregatorID float64  `json:"aggregator_id"`
            PayoutDetailsID float64  `json:"payout_details_id"`
         
    }
    
    // PayoutCustomer used by Payment
    type PayoutCustomer struct {

        
            Name string  `json:"name"`
            Mobile string  `json:"mobile"`
            ID float64  `json:"id"`
            UniqueExternalID string  `json:"unique_external_id"`
         
    }
    
    // Payout used by Payment
    type Payout struct {

        
            TransferType string  `json:"transfer_type"`
            UniqueTransferNo string  `json:"unique_transfer_no"`
            IsActive bool  `json:"is_active"`
            IsDefault bool  `json:"is_default"`
            MoreAttributes PayoutMoreAttributes  `json:"more_attributes"`
            PayoutsAggregators PayoutAggregator  `json:"payouts_aggregators"`
            Customers PayoutCustomer  `json:"customers"`
         
    }
    
    // PayoutsResponse used by Payment
    type PayoutsResponse struct {

        
            Items []Payout  `json:"items"`
            Success bool  `json:"success"`
         
    }
    
    // PayoutBankDetails used by Payment
    type PayoutBankDetails struct {

        
            BankName string  `json:"bank_name"`
            IfscCode string  `json:"ifsc_code"`
            AccountNo string  `json:"account_no"`
            AccountType string  `json:"account_type"`
            State string  `json:"state"`
            Country string  `json:"country"`
            City string  `json:"city"`
            AccountHolder string  `json:"account_holder"`
            Pincode float64  `json:"pincode"`
            BranchName string  `json:"branch_name"`
         
    }
    
    // PayoutRequest used by Payment
    type PayoutRequest struct {

        
            TransferType string  `json:"transfer_type"`
            IsActive bool  `json:"is_active"`
            Aggregator string  `json:"aggregator"`
            BankDetails PayoutBankDetails  `json:"bank_details"`
            UniqueExternalID string  `json:"unique_external_id"`
            Users map[string]interface{}  `json:"users"`
         
    }
    
    // PayoutResponse used by Payment
    type PayoutResponse struct {

        
            TransferType string  `json:"transfer_type"`
            UniqueTransferNo string  `json:"unique_transfer_no"`
            Success bool  `json:"success"`
            Created bool  `json:"created"`
            IsActive bool  `json:"is_active"`
            Aggregator string  `json:"aggregator"`
            Payouts map[string]interface{}  `json:"payouts"`
            BankDetails map[string]interface{}  `json:"bank_details"`
            PaymentStatus string  `json:"payment_status"`
            Users map[string]interface{}  `json:"users"`
         
    }
    
    // UpdatePayoutResponse used by Payment
    type UpdatePayoutResponse struct {

        
            IsDefault bool  `json:"is_default"`
            Success bool  `json:"success"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // UpdatePayoutRequest used by Payment
    type UpdatePayoutRequest struct {

        
            IsDefault bool  `json:"is_default"`
            UniqueExternalID string  `json:"unique_external_id"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // DeletePayoutResponse used by Payment
    type DeletePayoutResponse struct {

        
            Success bool  `json:"success"`
         
    }
    
    // SubscriptionPaymentMethodResponse used by Payment
    type SubscriptionPaymentMethodResponse struct {

        
            Data []map[string]interface{}  `json:"data"`
            Success bool  `json:"success"`
         
    }
    
    // DeleteSubscriptionPaymentMethodResponse used by Payment
    type DeleteSubscriptionPaymentMethodResponse struct {

        
            Success bool  `json:"success"`
         
    }
    
    // SubscriptionConfigResponse used by Payment
    type SubscriptionConfigResponse struct {

        
            Aggregator string  `json:"aggregator"`
            Success bool  `json:"success"`
            Config map[string]interface{}  `json:"config"`
         
    }
    
    // SaveSubscriptionSetupIntentRequest used by Payment
    type SaveSubscriptionSetupIntentRequest struct {

        
            UniqueExternalID string  `json:"unique_external_id"`
         
    }
    
    // SaveSubscriptionSetupIntentResponse used by Payment
    type SaveSubscriptionSetupIntentResponse struct {

        
            Data map[string]interface{}  `json:"data"`
            Success bool  `json:"success"`
         
    }
    
    // RefundAccountResponse used by Payment
    type RefundAccountResponse struct {

        
            Message string  `json:"message"`
            Data map[string]interface{}  `json:"data"`
            Success bool  `json:"success"`
            IsVerifiedFlag bool  `json:"is_verified_flag"`
         
    }
    
    // NotFoundResourceError used by Payment
    type NotFoundResourceError struct {

        
            Description string  `json:"description"`
            Success bool  `json:"success"`
            Code string  `json:"code"`
         
    }
    
    // BankDetailsForOTP used by Payment
    type BankDetailsForOTP struct {

        
            BankName string  `json:"bank_name"`
            IfscCode string  `json:"ifsc_code"`
            AccountNo string  `json:"account_no"`
            AccountHolder string  `json:"account_holder"`
            BranchName string  `json:"branch_name"`
         
    }
    
    // AddBeneficiaryDetailsOTPRequest used by Payment
    type AddBeneficiaryDetailsOTPRequest struct {

        
            Details BankDetailsForOTP  `json:"details"`
            OrderID string  `json:"order_id"`
         
    }
    
    // IfscCodeResponse used by Payment
    type IfscCodeResponse struct {

        
            BranchName string  `json:"branch_name"`
            BankName string  `json:"bank_name"`
            Success bool  `json:"success"`
         
    }
    
    // OrderBeneficiaryDetails used by Payment
    type OrderBeneficiaryDetails struct {

        
            Mobile string  `json:"mobile"`
            CreatedOn string  `json:"created_on"`
            AccountHolder string  `json:"account_holder"`
            TransferMode string  `json:"transfer_mode"`
            BankName string  `json:"bank_name"`
            IsActive bool  `json:"is_active"`
            DisplayName string  `json:"display_name"`
            ID float64  `json:"id"`
            Address string  `json:"address"`
            ModifiedOn string  `json:"modified_on"`
            Subtitle string  `json:"subtitle"`
            IfscCode string  `json:"ifsc_code"`
            Email string  `json:"email"`
            Comment string  `json:"comment"`
            BeneficiaryID string  `json:"beneficiary_id"`
            DelightsUserName string  `json:"delights_user_name"`
            AccountNo string  `json:"account_no"`
            Title string  `json:"title"`
            BranchName string  `json:"branch_name"`
         
    }
    
    // OrderBeneficiaryResponse used by Payment
    type OrderBeneficiaryResponse struct {

        
            ShowBeneficiaryDetails bool  `json:"show_beneficiary_details"`
            Beneficiaries []OrderBeneficiaryDetails  `json:"beneficiaries"`
         
    }
    
    // MultiTenderPaymentMeta used by Payment
    type MultiTenderPaymentMeta struct {

        
            PaymentGateway string  `json:"payment_gateway"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            CurrentStatus string  `json:"current_status"`
            PaymentID string  `json:"payment_id"`
            OrderID string  `json:"order_id"`
         
    }
    
    // MultiTenderPaymentMethod used by Payment
    type MultiTenderPaymentMethod struct {

        
            Name string  `json:"name"`
            Amount float64  `json:"amount"`
            Mode string  `json:"mode"`
            Meta MultiTenderPaymentMeta  `json:"meta"`
         
    }
    
    // PaymentConfirmationRequest used by Payment
    type PaymentConfirmationRequest struct {

        
            PaymentMethods []MultiTenderPaymentMethod  `json:"payment_methods"`
            OrderID string  `json:"order_id"`
         
    }
    
    // PaymentConfirmationResponse used by Payment
    type PaymentConfirmationResponse struct {

        
            Message string  `json:"message"`
            OrderID string  `json:"order_id"`
            Success bool  `json:"success"`
         
    }
    
    // PlatformPaymentOptions used by Payment
    type PlatformPaymentOptions struct {

        
            ModeOfPayment string  `json:"mode_of_payment"`
            AnonymousCod bool  `json:"anonymous_cod"`
            Enabled bool  `json:"enabled"`
            Source string  `json:"source"`
            CodAmountLimit float64  `json:"cod_amount_limit"`
            PaymentSelectionLock map[string]interface{}  `json:"payment_selection_lock"`
            Methods map[string]interface{}  `json:"methods"`
            CodCharges float64  `json:"cod_charges"`
            CallbackURL map[string]interface{}  `json:"callback_url"`
         
    }
    
    // PlatfromPaymentConfig used by Payment
    type PlatfromPaymentConfig struct {

        
            Message string  `json:"message"`
            Data PlatformPaymentOptions  `json:"data"`
            Success bool  `json:"success"`
         
    }
    
    // UpdatePlatformPaymentConfig used by Payment
    type UpdatePlatformPaymentConfig struct {

        
            AnonymousCod bool  `json:"anonymous_cod"`
            CodAmountLimit float64  `json:"cod_amount_limit"`
            PaymentSelectionLock map[string]interface{}  `json:"payment_selection_lock"`
            Methods map[string]interface{}  `json:"methods"`
            CodCharges float64  `json:"cod_charges"`
         
    }
    
    // CODdata used by Payment
    type CODdata struct {

        
            Usages float64  `json:"usages"`
            IsActive bool  `json:"is_active"`
            Limit float64  `json:"limit"`
            UserID string  `json:"user_id"`
            RemainingLimit float64  `json:"remaining_limit"`
         
    }
    
    // GetUserCODLimitResponse used by Payment
    type GetUserCODLimitResponse struct {

        
            UserCodData CODdata  `json:"user_cod_data"`
            Success bool  `json:"success"`
         
    }
    
    // SetCODForUserRequest used by Payment
    type SetCODForUserRequest struct {

        
            MerchantUserID string  `json:"merchant_user_id"`
            Mobileno string  `json:"mobileno"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // SetCODOptionResponse used by Payment
    type SetCODOptionResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // EdcModelData used by Payment
    type EdcModelData struct {

        
            Aggregator string  `json:"aggregator"`
            AggregatorID float64  `json:"aggregator_id"`
            Models []string  `json:"models"`
         
    }
    
    // EdcAggregatorAndModelListResponse used by Payment
    type EdcAggregatorAndModelListResponse struct {

        
            Data []EdcModelData  `json:"data"`
            Success bool  `json:"success"`
         
    }
    
    // StatisticsData used by Payment
    type StatisticsData struct {

        
            ActiveDeviceCount float64  `json:"active_device_count"`
            InactiveDeviceCount float64  `json:"inactive_device_count"`
         
    }
    
    // EdcDeviceStatsResponse used by Payment
    type EdcDeviceStatsResponse struct {

        
            Success bool  `json:"success"`
            Statistics StatisticsData  `json:"statistics"`
         
    }
    
    // EdcAddRequest used by Payment
    type EdcAddRequest struct {

        
            EdcModel string  `json:"edc_model"`
            AggregatorID float64  `json:"aggregator_id"`
            StoreID float64  `json:"store_id"`
            EdcDeviceSerialNo string  `json:"edc_device_serial_no"`
            TerminalSerialNo string  `json:"terminal_serial_no"`
            DeviceTag string  `json:"device_tag"`
         
    }
    
    // EdcDevice used by Payment
    type EdcDevice struct {

        
            AggregatorName string  `json:"aggregator_name"`
            DeviceTag string  `json:"device_tag"`
            EdcModel string  `json:"edc_model"`
            TerminalUniqueIdentifier string  `json:"terminal_unique_identifier"`
            IsActive bool  `json:"is_active"`
            MerchantStorePosCode string  `json:"merchant_store_pos_code"`
            StoreID float64  `json:"store_id"`
            EdcDeviceSerialNo string  `json:"edc_device_serial_no"`
            ApplicationID string  `json:"application_id"`
            TerminalSerialNo string  `json:"terminal_serial_no"`
            AggregatorID float64  `json:"aggregator_id"`
         
    }
    
    // EdcDeviceAddResponse used by Payment
    type EdcDeviceAddResponse struct {

        
            Data EdcDevice  `json:"data"`
            Success bool  `json:"success"`
         
    }
    
    // EdcDeviceDetailsResponse used by Payment
    type EdcDeviceDetailsResponse struct {

        
            Data EdcDevice  `json:"data"`
            Success bool  `json:"success"`
         
    }
    
    // EdcUpdateRequest used by Payment
    type EdcUpdateRequest struct {

        
            EdcModel string  `json:"edc_model"`
            IsActive string  `json:"is_active"`
            MerchantStorePosCode string  `json:"merchant_store_pos_code"`
            AggregatorID float64  `json:"aggregator_id"`
            StoreID float64  `json:"store_id"`
            EdcDeviceSerialNo string  `json:"edc_device_serial_no"`
            DeviceTag string  `json:"device_tag"`
         
    }
    
    // EdcDeviceUpdateResponse used by Payment
    type EdcDeviceUpdateResponse struct {

        
            Success bool  `json:"success"`
         
    }
    
    // Page used by Payment
    type Page struct {

        
            ItemTotal float64  `json:"item_total"`
            NextID string  `json:"next_id"`
            HasPrevious bool  `json:"has_previous"`
            HasNext bool  `json:"has_next"`
            Current float64  `json:"current"`
            Type string  `json:"type"`
            Size float64  `json:"size"`
         
    }
    
    // EdcDeviceListResponse used by Payment
    type EdcDeviceListResponse struct {

        
            Items []EdcDevice  `json:"items"`
            Page Page  `json:"page"`
            Success bool  `json:"success"`
         
    }
    
    // PaymentInitializationRequest used by Payment
    type PaymentInitializationRequest struct {

        
            Timeout float64  `json:"timeout"`
            Aggregator string  `json:"aggregator"`
            Contact string  `json:"contact"`
            Email string  `json:"email"`
            Currency string  `json:"currency"`
            MerchantOrderID string  `json:"merchant_order_id"`
            DeviceID string  `json:"device_id"`
            CustomerID string  `json:"customer_id"`
            RazorpayPaymentID string  `json:"razorpay_payment_id"`
            Amount float64  `json:"amount"`
            Method string  `json:"method"`
            OrderID string  `json:"order_id"`
            Vpa string  `json:"vpa"`
         
    }
    
    // PaymentInitializationResponse used by Payment
    type PaymentInitializationResponse struct {

        
            Timeout float64  `json:"timeout"`
            Success bool  `json:"success"`
            PollingURL string  `json:"polling_url"`
            Aggregator string  `json:"aggregator"`
            BqrImage string  `json:"bqr_image"`
            Currency string  `json:"currency"`
            UpiPollURL string  `json:"upi_poll_url"`
            MerchantOrderID string  `json:"merchant_order_id"`
            DeviceID string  `json:"device_id"`
            CustomerID string  `json:"customer_id"`
            RazorpayPaymentID string  `json:"razorpay_payment_id"`
            AggregatorOrderID string  `json:"aggregator_order_id"`
            Method string  `json:"method"`
            VirtualID string  `json:"virtual_id"`
            Amount float64  `json:"amount"`
            Vpa string  `json:"vpa"`
            Status string  `json:"status"`
         
    }
    
    // PaymentStatusUpdateRequest used by Payment
    type PaymentStatusUpdateRequest struct {

        
            Aggregator string  `json:"aggregator"`
            Contact string  `json:"contact"`
            Email string  `json:"email"`
            Currency string  `json:"currency"`
            MerchantOrderID string  `json:"merchant_order_id"`
            CustomerID string  `json:"customer_id"`
            DeviceID string  `json:"device_id"`
            Amount float64  `json:"amount"`
            Method string  `json:"method"`
            OrderID string  `json:"order_id"`
            Vpa string  `json:"vpa"`
            Status string  `json:"status"`
         
    }
    
    // PaymentStatusUpdateResponse used by Payment
    type PaymentStatusUpdateResponse struct {

        
            AggregatorName string  `json:"aggregator_name"`
            Success bool  `json:"success"`
            RedirectURL string  `json:"redirect_url"`
            Retry bool  `json:"retry"`
            Status string  `json:"status"`
         
    }
    
    // ResendOrCancelPaymentRequest used by Payment
    type ResendOrCancelPaymentRequest struct {

        
            OrderID string  `json:"order_id"`
            DeviceID string  `json:"device_id"`
            RequestType string  `json:"request_type"`
         
    }
    
    // LinkStatus used by Payment
    type LinkStatus struct {

        
            Message string  `json:"message"`
            Status bool  `json:"status"`
         
    }
    
    // ResendOrCancelPaymentResponse used by Payment
    type ResendOrCancelPaymentResponse struct {

        
            Data LinkStatus  `json:"data"`
            Success bool  `json:"success"`
         
    }
    

    
    // Prices used by Order
    type Prices struct {

        
            ValueOfGood float64  `json:"value_of_good"`
            RefundAmount float64  `json:"refund_amount"`
            CashbackApplied float64  `json:"cashback_applied"`
            Discount float64  `json:"discount"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            PromotionEffectiveDiscount float64  `json:"promotion_effective_discount"`
            CouponValue float64  `json:"coupon_value"`
            PriceMarked float64  `json:"price_marked"`
            Cashback float64  `json:"cashback"`
            FyndCredits float64  `json:"fynd_credits"`
            DeliveryCharge float64  `json:"delivery_charge"`
            CodCharges float64  `json:"cod_charges"`
            RefundCredit float64  `json:"refund_credit"`
            AmountPaidRoundoff float64  `json:"amount_paid_roundoff"`
            AmountPaid float64  `json:"amount_paid"`
            PriceEffective float64  `json:"price_effective"`
         
    }
    
    // ShipmentStatus used by Order
    type ShipmentStatus struct {

        
            Title string  `json:"title"`
            HexCode string  `json:"hex_code"`
            OpsStatus string  `json:"ops_status"`
            ActualStatus string  `json:"actual_status"`
            Status string  `json:"status"`
         
    }
    
    // GSTDetailsData used by Order
    type GSTDetailsData struct {

        
            ValueOfGood float64  `json:"value_of_good"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            GstinCode string  `json:"gstin_code"`
            GstFee float64  `json:"gst_fee"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
         
    }
    
    // PlatformItem used by Order
    type PlatformItem struct {

        
            CanCancel bool  `json:"can_cancel"`
            L3Category float64  `json:"l3_category"`
            CanReturn bool  `json:"can_return"`
            Color string  `json:"color"`
            L1Category []string  `json:"l1_category"`
            ID float64  `json:"id"`
            Code string  `json:"code"`
            L3CategoryName string  `json:"l3_category_name"`
            Image []string  `json:"image"`
            Name string  `json:"name"`
            DepartmentID float64  `json:"department_id"`
            Size string  `json:"size"`
            Images []string  `json:"images"`
         
    }
    
    // BagUnit used by Order
    type BagUnit struct {

        
            BagID float64  `json:"bag_id"`
            Prices Prices  `json:"prices"`
            CanCancel bool  `json:"can_cancel"`
            CanReturn bool  `json:"can_return"`
            Gst GSTDetailsData  `json:"gst"`
            OrderingChannel string  `json:"ordering_channel"`
            TotalShipmentBags float64  `json:"total_shipment_bags"`
            ShipmentID string  `json:"shipment_id"`
            Item PlatformItem  `json:"item"`
            ItemQuantity float64  `json:"item_quantity"`
            Status map[string]interface{}  `json:"status"`
         
    }
    
    // ShipmentItemFulFillingStore used by Order
    type ShipmentItemFulFillingStore struct {

        
            Code string  `json:"code"`
            ID string  `json:"id"`
         
    }
    
    // PaymentModeInfo used by Order
    type PaymentModeInfo struct {

        
            Type string  `json:"type"`
            Logo string  `json:"logo"`
         
    }
    
    // UserDataInfo used by Order
    type UserDataInfo struct {

        
            Gender string  `json:"gender"`
            UID float64  `json:"uid"`
            LastName string  `json:"last_name"`
            Mobile string  `json:"mobile"`
            AvisUserID string  `json:"avis_user_id"`
            Email string  `json:"email"`
            IsAnonymousUser bool  `json:"is_anonymous_user"`
            Name string  `json:"name"`
            FirstName string  `json:"first_name"`
         
    }
    
    // ShipmentItem used by Order
    type ShipmentItem struct {

        
            Prices Prices  `json:"prices"`
            ShipmentStatus ShipmentStatus  `json:"shipment_status"`
            FulfillingCentre string  `json:"fulfilling_centre"`
            CreatedAt string  `json:"created_at"`
            PaymentMethods map[string]interface{}  `json:"payment_methods"`
            Bags []BagUnit  `json:"bags"`
            Company map[string]interface{}  `json:"company"`
            ID string  `json:"id"`
            FulfillingStore ShipmentItemFulFillingStore  `json:"fulfilling_store"`
            TotalShipmentsInOrder float64  `json:"total_shipments_in_order"`
            Channel map[string]interface{}  `json:"channel"`
            TotalBagsCount float64  `json:"total_bags_count"`
            ShipmentCreatedAt float64  `json:"shipment_created_at"`
            PaymentModeInfo PaymentModeInfo  `json:"payment_mode_info"`
            Sla map[string]interface{}  `json:"sla"`
            User UserDataInfo  `json:"user"`
            Application map[string]interface{}  `json:"application"`
         
    }
    
    // FilterInfoOption used by Order
    type FilterInfoOption struct {

        
            Text string  `json:"text"`
            Name string  `json:"name"`
            Value string  `json:"value"`
         
    }
    
    // FiltersInfo used by Order
    type FiltersInfo struct {

        
            Type string  `json:"type"`
            Text string  `json:"text"`
            Options []FilterInfoOption  `json:"options"`
            Value string  `json:"value"`
         
    }
    
    // ShipmentInternalPlatformViewResponse used by Order
    type ShipmentInternalPlatformViewResponse struct {

        
            Items []ShipmentItem  `json:"items"`
            Page map[string]interface{}  `json:"page"`
            AppliedFilters map[string]interface{}  `json:"applied_filters"`
            Filters []FiltersInfo  `json:"filters"`
         
    }
    
    // Error used by Order
    type Error struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // ShipmentPayments used by Order
    type ShipmentPayments struct {

        
            Mode string  `json:"mode"`
            Logo string  `json:"logo"`
            Source string  `json:"source"`
         
    }
    
    // Dimensions used by Order
    type Dimensions struct {

        
            Unit string  `json:"unit"`
            Length float64  `json:"length"`
            Height float64  `json:"height"`
            IsDefault bool  `json:"is_default"`
            Width float64  `json:"width"`
         
    }
    
    // Meta used by Order
    type Meta struct {

        
            Dimension Dimensions  `json:"dimension"`
         
    }
    
    // BagStateMapper used by Order
    type BagStateMapper struct {

        
            AppFacing bool  `json:"app_facing"`
            StateType string  `json:"state_type"`
            DisplayName string  `json:"display_name"`
            IsActive bool  `json:"is_active"`
            AppStateName string  `json:"app_state_name"`
            AppDisplayName string  `json:"app_display_name"`
            BsID float64  `json:"bs_id"`
            Name string  `json:"name"`
            JourneyType string  `json:"journey_type"`
            NotifyCustomer bool  `json:"notify_customer"`
         
    }
    
    // BagStatusHistory used by Order
    type BagStatusHistory struct {

        
            BagID float64  `json:"bag_id"`
            KafkaSync bool  `json:"kafka_sync"`
            StateType string  `json:"state_type"`
            BshID float64  `json:"bsh_id"`
            Forward bool  `json:"forward"`
            DeliveryAwbNumber string  `json:"delivery_awb_number"`
            CreatedAt string  `json:"created_at"`
            BagStateMapper BagStateMapper  `json:"bag_state_mapper"`
            Reasons []map[string]interface{}  `json:"reasons"`
            DisplayName string  `json:"display_name"`
            DeliveryPartnerID float64  `json:"delivery_partner_id"`
            ShipmentID string  `json:"shipment_id"`
            UpdatedAt string  `json:"updated_at"`
            AppDisplayName string  `json:"app_display_name"`
            StateID float64  `json:"state_id"`
            StoreID float64  `json:"store_id"`
            Status string  `json:"status"`
         
    }
    
    // OrderingStoreDetails used by Order
    type OrderingStoreDetails struct {

        
            StoreName string  `json:"store_name"`
            Country string  `json:"country"`
            Phone string  `json:"phone"`
            Code string  `json:"code"`
            OrderingStoreID float64  `json:"ordering_store_id"`
            City string  `json:"city"`
            Address string  `json:"address"`
            Pincode string  `json:"pincode"`
            ContactPerson string  `json:"contact_person"`
            State string  `json:"state"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // UserDetailsData used by Order
    type UserDetailsData struct {

        
            Country string  `json:"country"`
            Phone string  `json:"phone"`
            Email string  `json:"email"`
            City string  `json:"city"`
            Name string  `json:"name"`
            Address string  `json:"address"`
            Pincode string  `json:"pincode"`
            State string  `json:"state"`
         
    }
    
    // OrderDetailsData used by Order
    type OrderDetailsData struct {

        
            AffiliateID string  `json:"affiliate_id"`
            OrderDate string  `json:"order_date"`
            OrderingChannel string  `json:"ordering_channel"`
            Source string  `json:"source"`
            TaxDetails map[string]interface{}  `json:"tax_details"`
            CodCharges string  `json:"cod_charges"`
            OrderingChannelLogo map[string]interface{}  `json:"ordering_channel_logo"`
            FyndOrderID string  `json:"fynd_order_id"`
            OrderValue string  `json:"order_value"`
         
    }
    
    // FulfillingStore used by Order
    type FulfillingStore struct {

        
            StoreName string  `json:"store_name"`
            FulfillmentChannel string  `json:"fulfillment_channel"`
            Country string  `json:"country"`
            Phone string  `json:"phone"`
            ID float64  `json:"id"`
            Code string  `json:"code"`
            City string  `json:"city"`
            Address string  `json:"address"`
            Pincode string  `json:"pincode"`
            ContactPerson string  `json:"contact_person"`
            State string  `json:"state"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // OrderBrandName used by Order
    type OrderBrandName struct {

        
            CreatedOn string  `json:"created_on"`
            BrandName string  `json:"brand_name"`
            Company string  `json:"company"`
            ModifiedOn string  `json:"modified_on"`
            ID float64  `json:"id"`
            Logo string  `json:"logo"`
         
    }
    
    // BagGST used by Order
    type BagGST struct {

        
            HsnCode string  `json:"hsn_code"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            ValueOfGood float64  `json:"value_of_good"`
            GstTag string  `json:"gst_tag"`
            GstinCode string  `json:"gstin_code"`
            IsDefaultHsnCode bool  `json:"is_default_hsn_code"`
            GstFee float64  `json:"gst_fee"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
         
    }
    
    // CurrentStatus used by Order
    type CurrentStatus struct {

        
            BagID float64  `json:"bag_id"`
            StateType string  `json:"state_type"`
            KafkaSync bool  `json:"kafka_sync"`
            CurrentStatusID float64  `json:"current_status_id"`
            DeliveryAwbNumber string  `json:"delivery_awb_number"`
            CreatedAt string  `json:"created_at"`
            BagStateMapper BagStateMapper  `json:"bag_state_mapper"`
            DeliveryPartnerID float64  `json:"delivery_partner_id"`
            ShipmentID string  `json:"shipment_id"`
            UpdatedAt float64  `json:"updated_at"`
            StateID float64  `json:"state_id"`
            StoreID float64  `json:"store_id"`
            Status string  `json:"status"`
         
    }
    
    // BagConfigs used by Order
    type BagConfigs struct {

        
            CanBeCancelled bool  `json:"can_be_cancelled"`
            IsCustomerReturnAllowed bool  `json:"is_customer_return_allowed"`
            IsActive bool  `json:"is_active"`
            EnableTracking bool  `json:"enable_tracking"`
            AllowForceReturn bool  `json:"allow_force_return"`
            IsReturnable bool  `json:"is_returnable"`
         
    }
    
    // ItemCriterias used by Order
    type ItemCriterias struct {

        
            ItemBrand []float64  `json:"item_brand"`
         
    }
    
    // BuyRules used by Order
    type BuyRules struct {

        
            CartConditions map[string]interface{}  `json:"cart_conditions"`
            ItemCriteria ItemCriterias  `json:"item_criteria"`
         
    }
    
    // DiscountRules used by Order
    type DiscountRules struct {

        
            Type string  `json:"type"`
            Value float64  `json:"value"`
         
    }
    
    // AppliedPromos used by Order
    type AppliedPromos struct {

        
            ArticleQuantity float64  `json:"article_quantity"`
            BuyRules []BuyRules  `json:"buy_rules"`
            Amount float64  `json:"amount"`
            DiscountRules []DiscountRules  `json:"discount_rules"`
            MrpPromotion bool  `json:"mrp_promotion"`
            PromotionName string  `json:"promotion_name"`
            PromotionType string  `json:"promotion_type"`
            PromoID string  `json:"promo_id"`
         
    }
    
    // Identifier used by Order
    type Identifier struct {

        
            Ean string  `json:"ean"`
         
    }
    
    // FinancialBreakup used by Order
    type FinancialBreakup struct {

        
            ValueOfGood float64  `json:"value_of_good"`
            HsnCode string  `json:"hsn_code"`
            PromotionEffectiveDiscount float64  `json:"promotion_effective_discount"`
            PriceMarked float64  `json:"price_marked"`
            DeliveryCharge float64  `json:"delivery_charge"`
            CodCharges float64  `json:"cod_charges"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            Size string  `json:"size"`
            TransferPrice float64  `json:"transfer_price"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            CashbackApplied float64  `json:"cashback_applied"`
            CouponValue float64  `json:"coupon_value"`
            GstFee float64  `json:"gst_fee"`
            FyndCredits float64  `json:"fynd_credits"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            AddedToFyndCash bool  `json:"added_to_fynd_cash"`
            AmountPaid float64  `json:"amount_paid"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            RefundCredit float64  `json:"refund_credit"`
            Identifiers Identifier  `json:"identifiers"`
            ItemName string  `json:"item_name"`
            Discount float64  `json:"discount"`
            GstTag string  `json:"gst_tag"`
            Cashback float64  `json:"cashback"`
            TotalUnits float64  `json:"total_units"`
            AmountPaidRoundoff float64  `json:"amount_paid_roundoff"`
            PriceEffective float64  `json:"price_effective"`
         
    }
    
    // PlatformDeliveryAddress used by Order
    type PlatformDeliveryAddress struct {

        
            Country string  `json:"country"`
            Area string  `json:"area"`
            Phone string  `json:"phone"`
            CreatedAt string  `json:"created_at"`
            AddressCategory string  `json:"address_category"`
            Email string  `json:"email"`
            Latitude float64  `json:"latitude"`
            Version string  `json:"version"`
            Landmark string  `json:"landmark"`
            UpdatedAt string  `json:"updated_at"`
            Address2 string  `json:"address2"`
            City string  `json:"city"`
            AddressType string  `json:"address_type"`
            Pincode string  `json:"pincode"`
            ContactPerson string  `json:"contact_person"`
            State string  `json:"state"`
            Longitude float64  `json:"longitude"`
            Address1 string  `json:"address1"`
         
    }
    
    // OrderBagArticle used by Order
    type OrderBagArticle struct {

        
            ReturnConfig map[string]interface{}  `json:"return_config"`
            UID string  `json:"uid"`
            Identifiers map[string]interface{}  `json:"identifiers"`
         
    }
    
    // OrderBags used by Order
    type OrderBags struct {

        
            BagID float64  `json:"bag_id"`
            Prices Prices  `json:"prices"`
            Brand OrderBrandName  `json:"brand"`
            GstDetails BagGST  `json:"gst_details"`
            Identifier string  `json:"identifier"`
            CurrentStatus CurrentStatus  `json:"current_status"`
            Quantity float64  `json:"quantity"`
            EntityType string  `json:"entity_type"`
            SellerIdentifier string  `json:"seller_identifier"`
            CanCancel bool  `json:"can_cancel"`
            CanReturn bool  `json:"can_return"`
            BagConfigs BagConfigs  `json:"bag_configs"`
            AppliedPromos []AppliedPromos  `json:"applied_promos"`
            FinancialBreakup FinancialBreakup  `json:"financial_breakup"`
            ParentPromoBags map[string]interface{}  `json:"parent_promo_bags"`
            DisplayName string  `json:"display_name"`
            DeliveryAddress PlatformDeliveryAddress  `json:"delivery_address"`
            Article OrderBagArticle  `json:"article"`
            Item PlatformItem  `json:"item"`
            LineNumber float64  `json:"line_number"`
         
    }
    
    // ShipmentStatusData used by Order
    type ShipmentStatusData struct {

        
            CreatedAt string  `json:"created_at"`
            BagList []string  `json:"bag_list"`
            ID float64  `json:"id"`
            ShipmentID string  `json:"shipment_id"`
            Status string  `json:"status"`
         
    }
    
    // CompanyDetails used by Order
    type CompanyDetails struct {

        
            CompanyID float64  `json:"company_id"`
            CompanyName string  `json:"company_name"`
            Address map[string]interface{}  `json:"address"`
            CompanyCin string  `json:"company_cin"`
         
    }
    
    // TrackingList used by Order
    type TrackingList struct {

        
            Text string  `json:"text"`
            IsPassed bool  `json:"is_passed"`
            Time string  `json:"time"`
            IsCurrent bool  `json:"is_current"`
            Status string  `json:"status"`
         
    }
    
    // PDFLinks used by Order
    type PDFLinks struct {

        
            InvoiceA6 string  `json:"invoice_a6"`
            LabelA4 string  `json:"label_a4"`
            LabelType string  `json:"label_type"`
            Label string  `json:"label"`
            InvoiceA4 string  `json:"invoice_a4"`
            InvoiceType string  `json:"invoice_type"`
            PoInvoice string  `json:"po_invoice"`
            Invoice string  `json:"invoice"`
            B2b string  `json:"b2b"`
            InvoicePos string  `json:"invoice_pos"`
            LabelPos string  `json:"label_pos"`
            CreditNoteURL string  `json:"credit_note_url"`
            LabelA6 string  `json:"label_a6"`
         
    }
    
    // LockData used by Order
    type LockData struct {

        
            Locked bool  `json:"locked"`
            LockMessage string  `json:"lock_message"`
            Mto bool  `json:"mto"`
         
    }
    
    // ShipmentTimeStamp used by Order
    type ShipmentTimeStamp struct {

        
            TMin string  `json:"t_min"`
            TMax string  `json:"t_max"`
         
    }
    
    // BuyerDetails used by Order
    type BuyerDetails struct {

        
            Gstin string  `json:"gstin"`
            Name string  `json:"name"`
            City string  `json:"city"`
            Address string  `json:"address"`
            Pincode float64  `json:"pincode"`
            State string  `json:"state"`
            AjioSiteID string  `json:"ajio_site_id"`
         
    }
    
    // EInvoice used by Order
    type EInvoice struct {

        
            Irn string  `json:"irn"`
            SignedInvoice string  `json:"signed_invoice"`
            SignedQrCode string  `json:"signed_qr_code"`
            AcknowledgeNo float64  `json:"acknowledge_no"`
            ErrorCode string  `json:"error_code"`
            AcknowledgeDate string  `json:"acknowledge_date"`
            ErrorMessage string  `json:"error_message"`
         
    }
    
    // EinvoiceInfo used by Order
    type EinvoiceInfo struct {

        
            Invoice EInvoice  `json:"invoice"`
            CreditNote EInvoice  `json:"credit_note"`
         
    }
    
    // DebugInfo used by Order
    type DebugInfo struct {

        
            StormbreakerUUID string  `json:"stormbreaker_uuid"`
         
    }
    
    // Formatted used by Order
    type Formatted struct {

        
            FMin string  `json:"f_min"`
            FMax string  `json:"f_max"`
         
    }
    
    // ShipmentMeta used by Order
    type ShipmentMeta struct {

        
            DpSortKey string  `json:"dp_sort_key"`
            ReturnAffiliateShipmentID string  `json:"return_affiliate_shipment_id"`
            ReturnAffiliateOrderID string  `json:"return_affiliate_order_id"`
            B2cBuyerDetails map[string]interface{}  `json:"b2c_buyer_details"`
            DpID string  `json:"dp_id"`
            LockData LockData  `json:"lock_data"`
            AssignDpFromSb bool  `json:"assign_dp_from_sb"`
            Timestamp ShipmentTimeStamp  `json:"timestamp"`
            PoNumber string  `json:"po_number"`
            MarketplaceStoreID string  `json:"marketplace_store_id"`
            B2bBuyerDetails BuyerDetails  `json:"b2b_buyer_details"`
            ReturnStoreNode float64  `json:"return_store_node"`
            BoxType string  `json:"box_type"`
            EinvoiceInfo EinvoiceInfo  `json:"einvoice_info"`
            FulfilmentPriorityText string  `json:"fulfilment_priority_text"`
            EwaybillInfo map[string]interface{}  `json:"ewaybill_info"`
            DueDate string  `json:"due_date"`
            ForwardAffiliateOrderID string  `json:"forward_affiliate_order_id"`
            OrderType string  `json:"order_type"`
            External map[string]interface{}  `json:"external"`
            ReturnAwbNumber string  `json:"return_awb_number"`
            PackagingName string  `json:"packaging_name"`
            AwbNumber string  `json:"awb_number"`
            BagWeight map[string]interface{}  `json:"bag_weight"`
            DebugInfo DebugInfo  `json:"debug_info"`
            ShipmentWeight float64  `json:"shipment_weight"`
            ReturnDetails map[string]interface{}  `json:"return_details"`
            AutoTriggerDpAssignmentAcf bool  `json:"auto_trigger_dp_assignment_acf"`
            DpOptions map[string]interface{}  `json:"dp_options"`
            ForwardAffiliateShipmentID string  `json:"forward_affiliate_shipment_id"`
            Formatted Formatted  `json:"formatted"`
            Weight float64  `json:"weight"`
            SameStoreAvailable bool  `json:"same_store_available"`
            DpName string  `json:"dp_name"`
            StoreInvoiceUpdatedDate string  `json:"store_invoice_updated_date"`
            ShipmentVolumetricWeight float64  `json:"shipment_volumetric_weight"`
         
    }
    
    // AffiliateMeta used by Order
    type AffiliateMeta struct {

        
            IsPriority bool  `json:"is_priority"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            ChannelOrderID string  `json:"channel_order_id"`
            CouponCode string  `json:"coupon_code"`
            OrderItemID string  `json:"order_item_id"`
            BoxType string  `json:"box_type"`
            EmployeeDiscount float64  `json:"employee_discount"`
            Quantity float64  `json:"quantity"`
            SizeLevelTotalQty float64  `json:"size_level_total_qty"`
            ChannelShipmentID string  `json:"channel_shipment_id"`
            DueDate string  `json:"due_date"`
         
    }
    
    // AffiliateDetails used by Order
    type AffiliateDetails struct {

        
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            AffiliateID string  `json:"affiliate_id"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
            PdfLinks PDFLinks  `json:"pdf_links"`
            AdID string  `json:"ad_id"`
            ShipmentMeta ShipmentMeta  `json:"shipment_meta"`
            AffiliateStoreID string  `json:"affiliate_store_id"`
            CompanyAffiliateTag string  `json:"company_affiliate_tag"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            AffiliateMeta AffiliateMeta  `json:"affiliate_meta"`
         
    }
    
    // DPDetailsData used by Order
    type DPDetailsData struct {

        
            Country string  `json:"country"`
            GstTag string  `json:"gst_tag"`
            EwayBillID string  `json:"eway_bill_id"`
            TrackURL string  `json:"track_url"`
            AwbNo string  `json:"awb_no"`
            ID float64  `json:"id"`
            Name string  `json:"name"`
            Pincode string  `json:"pincode"`
         
    }
    
    // PlatformShipment used by Order
    type PlatformShipment struct {

        
            Prices Prices  `json:"prices"`
            UserAgent string  `json:"user_agent"`
            ShipmentImages []string  `json:"shipment_images"`
            PriorityText string  `json:"priority_text"`
            GstDetails GSTDetailsData  `json:"gst_details"`
            ShipmentID string  `json:"shipment_id"`
            JourneyType string  `json:"journey_type"`
            CustomMeta []map[string]interface{}  `json:"custom_meta"`
            PickedDate string  `json:"picked_date"`
            Payments ShipmentPayments  `json:"payments"`
            Meta Meta  `json:"meta"`
            Coupon map[string]interface{}  `json:"coupon"`
            PlatformLogo string  `json:"platform_logo"`
            OperationalStatus string  `json:"operational_status"`
            BagStatusHistory []BagStatusHistory  `json:"bag_status_history"`
            ShipmentQuantity float64  `json:"shipment_quantity"`
            TotalItems float64  `json:"total_items"`
            OrderingStore OrderingStoreDetails  `json:"ordering_store"`
            FulfilmentPriority float64  `json:"fulfilment_priority"`
            Vertical string  `json:"vertical"`
            User UserDataInfo  `json:"user"`
            PackagingType string  `json:"packaging_type"`
            DeliveryDetails UserDetailsData  `json:"delivery_details"`
            DeliverySlot map[string]interface{}  `json:"delivery_slot"`
            Order OrderDetailsData  `json:"order"`
            PaymentMode string  `json:"payment_mode"`
            FulfillingStore FulfillingStore  `json:"fulfilling_store"`
            Bags []OrderBags  `json:"bags"`
            Status ShipmentStatusData  `json:"status"`
            ShipmentStatus string  `json:"shipment_status"`
            InvoiceID string  `json:"invoice_id"`
            TotalBags float64  `json:"total_bags"`
            CompanyDetails CompanyDetails  `json:"company_details"`
            TrackingList []TrackingList  `json:"tracking_list"`
            AffiliateDetails AffiliateDetails  `json:"affiliate_details"`
            BillingDetails UserDetailsData  `json:"billing_details"`
            Invoice EinvoiceInfo  `json:"invoice"`
            DpDetails DPDetailsData  `json:"dp_details"`
            EnableDpTracking bool  `json:"enable_dp_tracking"`
         
    }
    
    // ShipmentInfoResponse used by Order
    type ShipmentInfoResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
            Shipments []PlatformShipment  `json:"shipments"`
         
    }
    
    // TaxDetails used by Order
    type TaxDetails struct {

        
            Gstin string  `json:"gstin"`
            PanNo string  `json:"pan_no"`
         
    }
    
    // PlatformUserDetails used by Order
    type PlatformUserDetails struct {

        
            PlatformUserID string  `json:"platform_user_id"`
            PlatformUserFirstName string  `json:"platform_user_first_name"`
            PlatformUserLastName string  `json:"platform_user_last_name"`
         
    }
    
    // TransactionData used by Order
    type TransactionData struct {

        
            Currency string  `json:"currency"`
            TransactionID string  `json:"transaction_id"`
            Entity string  `json:"entity"`
            TerminalID string  `json:"terminal_id"`
            UniqueReferenceNumber string  `json:"unique_reference_number"`
            PaymentID string  `json:"payment_id"`
            AmountPaid string  `json:"amount_paid"`
            Status string  `json:"status"`
         
    }
    
    // BillingStaffDetails used by Order
    type BillingStaffDetails struct {

        
            LastName string  `json:"last_name"`
            User string  `json:"user"`
            StaffID float64  `json:"staff_id"`
            FirstName string  `json:"first_name"`
         
    }
    
    // OrderMeta used by Order
    type OrderMeta struct {

        
            CustomerNote string  `json:"customer_note"`
            CurrencySymbol string  `json:"currency_symbol"`
            EmployeeID float64  `json:"employee_id"`
            OrderTags []map[string]interface{}  `json:"order_tags"`
            OrderPlatform string  `json:"order_platform"`
            OrderingStore float64  `json:"ordering_store"`
            PlatformUserDetails PlatformUserDetails  `json:"platform_user_details"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            OrderType string  `json:"order_type"`
            MongoCartID float64  `json:"mongo_cart_id"`
            TransactionData TransactionData  `json:"transaction_data"`
            Comment string  `json:"comment"`
            CompanyLogo string  `json:"company_logo"`
            Staff map[string]interface{}  `json:"staff"`
            PaymentType string  `json:"payment_type"`
            CartID float64  `json:"cart_id"`
            OrderChildEntities []string  `json:"order_child_entities"`
            BillingStaffDetails BillingStaffDetails  `json:"billing_staff_details"`
            Files []map[string]interface{}  `json:"files"`
         
    }
    
    // OrderDict used by Order
    type OrderDict struct {

        
            Prices Prices  `json:"prices"`
            OrderDate string  `json:"order_date"`
            PaymentMethods map[string]interface{}  `json:"payment_methods"`
            TaxDetails TaxDetails  `json:"tax_details"`
            Meta OrderMeta  `json:"meta"`
            FyndOrderID string  `json:"fynd_order_id"`
         
    }
    
    // ShipmentDetailsResponse used by Order
    type ShipmentDetailsResponse struct {

        
            Order OrderDict  `json:"order"`
            Success bool  `json:"success"`
            Shipments []PlatformShipment  `json:"shipments"`
         
    }
    
    // SubLane used by Order
    type SubLane struct {

        
            Index float64  `json:"index"`
            Text string  `json:"text"`
            TotalItems float64  `json:"total_items"`
            Actions []map[string]interface{}  `json:"actions"`
            Value string  `json:"value"`
         
    }
    
    // SuperLane used by Order
    type SuperLane struct {

        
            TotalItems float64  `json:"total_items"`
            Text string  `json:"text"`
            Options []SubLane  `json:"options"`
            Value string  `json:"value"`
         
    }
    
    // LaneConfigResponse used by Order
    type LaneConfigResponse struct {

        
            SuperLanes []SuperLane  `json:"super_lanes"`
         
    }
    
    // PlatformBreakupValues used by Order
    type PlatformBreakupValues struct {

        
            Display string  `json:"display"`
            Name string  `json:"name"`
            Value string  `json:"value"`
         
    }
    
    // PlatformChannel used by Order
    type PlatformChannel struct {

        
            Logo string  `json:"logo"`
            Name string  `json:"name"`
         
    }
    
    // PlatformOrderItems used by Order
    type PlatformOrderItems struct {

        
            TotalOrderValue float64  `json:"total_order_value"`
            OrderID string  `json:"order_id"`
            BreakupValues []PlatformBreakupValues  `json:"breakup_values"`
            Shipments []PlatformShipment  `json:"shipments"`
            PaymentMode string  `json:"payment_mode"`
            Channel PlatformChannel  `json:"channel"`
            UserInfo UserDataInfo  `json:"user_info"`
            OrderCreatedTime string  `json:"order_created_time"`
            Meta map[string]interface{}  `json:"meta"`
            OrderValue float64  `json:"order_value"`
         
    }
    
    // Page used by Order
    type Page struct {

        
            ItemTotal float64  `json:"item_total"`
            NextID string  `json:"next_id"`
            HasPrevious bool  `json:"has_previous"`
            HasNext bool  `json:"has_next"`
            Current float64  `json:"current"`
            Type string  `json:"type"`
            Size float64  `json:"size"`
         
    }
    
    // OrderListingResponse used by Order
    type OrderListingResponse struct {

        
            Items []PlatformOrderItems  `json:"items"`
            Lane string  `json:"lane"`
            Page Page  `json:"page"`
            TotalCount float64  `json:"total_count"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // Options used by Order
    type Options struct {

        
            Text string  `json:"text"`
            Value float64  `json:"value"`
         
    }
    
    // MetricsCount used by Order
    type MetricsCount struct {

        
            Text string  `json:"text"`
            Options []Options  `json:"options"`
            Key string  `json:"key"`
            Value float64  `json:"value"`
         
    }
    
    // MetricCountResponse used by Order
    type MetricCountResponse struct {

        
            Items []MetricsCount  `json:"items"`
         
    }
    
    // PlatformTrack used by Order
    type PlatformTrack struct {

        
            RawStatus string  `json:"raw_status"`
            AccountName string  `json:"account_name"`
            Awb string  `json:"awb"`
            UpdatedAt string  `json:"updated_at"`
            ShipmentType string  `json:"shipment_type"`
            Reason string  `json:"reason"`
            UpdatedTime string  `json:"updated_time"`
            LastLocationRecievedAt string  `json:"last_location_recieved_at"`
            Meta map[string]interface{}  `json:"meta"`
            Status string  `json:"status"`
         
    }
    
    // PlatformShipmentTrack used by Order
    type PlatformShipmentTrack struct {

        
            Results []PlatformTrack  `json:"results"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // AdvanceFilterInfo used by Order
    type AdvanceFilterInfo struct {

        
            Processed []FiltersInfo  `json:"processed"`
            Unfulfilled []FiltersInfo  `json:"unfulfilled"`
            ActionCentre []FiltersInfo  `json:"action_centre"`
            Filters []FiltersInfo  `json:"filters"`
            Returned []FiltersInfo  `json:"returned"`
         
    }
    
    // FiltersResponse used by Order
    type FiltersResponse struct {

        
            AdvanceFilter AdvanceFilterInfo  `json:"advance_filter"`
            GlobalFilter []FiltersInfo  `json:"global_filter"`
         
    }
    
    // Success used by Order
    type Success struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // OmsReports used by Order
    type OmsReports struct {

        
            ReportName string  `json:"report_name"`
            S3Key string  `json:"s3_key"`
            DisplayName string  `json:"display_name"`
            ReportCreatedAt string  `json:"report_created_at"`
            ReportType string  `json:"report_type"`
            ReportRequestedAt string  `json:"report_requested_at"`
            RequestDetails map[string]interface{}  `json:"request_details"`
            ReportID string  `json:"report_id"`
            Status string  `json:"status"`
         
    }
    
    // JioCodeUpsertDataSet used by Order
    type JioCodeUpsertDataSet struct {

        
            CompanyID string  `json:"company_id"`
            ItemID string  `json:"item_id"`
            ArticleID string  `json:"article_id"`
            JioCode string  `json:"jio_code"`
         
    }
    
    // JioCodeUpsertPayload used by Order
    type JioCodeUpsertPayload struct {

        
            Data []JioCodeUpsertDataSet  `json:"data"`
         
    }
    
    // NestedErrorSchemaDataSet used by Order
    type NestedErrorSchemaDataSet struct {

        
            Message string  `json:"message"`
            Type string  `json:"type"`
            Value string  `json:"value"`
         
    }
    
    // JioCodeUpsertResponse used by Order
    type JioCodeUpsertResponse struct {

        
            Error []NestedErrorSchemaDataSet  `json:"error"`
            TraceID string  `json:"trace_id"`
            Identifier string  `json:"identifier"`
            Success bool  `json:"success"`
            Data []map[string]interface{}  `json:"data"`
         
    }
    
    // BulkInvoicingResponse used by Order
    type BulkInvoicingResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // BulkInvoiceLabelResponse used by Order
    type BulkInvoiceLabelResponse struct {

        
            CompanyID string  `json:"company_id"`
            StoreName string  `json:"store_name"`
            Label map[string]interface{}  `json:"label"`
            InvoiceStatus string  `json:"invoice_status"`
            DoInvoiceLabelGenerated bool  `json:"do_invoice_label_generated"`
            Invoice map[string]interface{}  `json:"invoice"`
            Data map[string]interface{}  `json:"data"`
            BatchID string  `json:"batch_id"`
            StoreID string  `json:"store_id"`
            StoreCode string  `json:"store_code"`
         
    }
    
    // FileUploadResponse used by Order
    type FileUploadResponse struct {

        
            URL string  `json:"url"`
            Expiry float64  `json:"expiry"`
         
    }
    
    // URL used by Order
    type URL struct {

        
            URL string  `json:"url"`
         
    }
    
    // FileResponse used by Order
    type FileResponse struct {

        
            Namespace string  `json:"namespace"`
            ContentType string  `json:"content_type"`
            FilePath string  `json:"file_path"`
            Operation string  `json:"operation"`
            Upload FileUploadResponse  `json:"upload"`
            Tags []string  `json:"tags"`
            FileName string  `json:"file_name"`
            Method string  `json:"method"`
            Size float64  `json:"size"`
            Cdn URL  `json:"cdn"`
         
    }
    
    // BulkListingPage used by Order
    type BulkListingPage struct {

        
            HasNext bool  `json:"has_next"`
            Current float64  `json:"current"`
            Type string  `json:"type"`
            HasPrevious bool  `json:"has_previous"`
            Size float64  `json:"size"`
            Total float64  `json:"total"`
         
    }
    
    // bulkListingData used by Order
    type bulkListingData struct {

        
            CompanyID float64  `json:"company_id"`
            FailedShipments []map[string]interface{}  `json:"failed_shipments"`
            UserID string  `json:"user_id"`
            Total float64  `json:"total"`
            StoreCode string  `json:"store_code"`
            StoreName string  `json:"store_name"`
            UploadedOn string  `json:"uploaded_on"`
            SuccessfulShipments []map[string]interface{}  `json:"successful_shipments"`
            ProcessingShipments []string  `json:"processing_shipments"`
            UserName string  `json:"user_name"`
            ID string  `json:"id"`
            FileName string  `json:"file_name"`
            BatchID string  `json:"batch_id"`
            StoreID float64  `json:"store_id"`
            ExcelURL string  `json:"excel_url"`
            Successful float64  `json:"successful"`
            Status string  `json:"status"`
            Processing float64  `json:"processing"`
            Failed float64  `json:"failed"`
         
    }
    
    // BulkListingResponse used by Order
    type BulkListingResponse struct {

        
            Success bool  `json:"success"`
            Error string  `json:"error"`
            Page BulkListingPage  `json:"page"`
            Data []bulkListingData  `json:"data"`
         
    }
    
    // QuestionSet used by Order
    type QuestionSet struct {

        
            DisplayName string  `json:"display_name"`
            ID float64  `json:"id"`
         
    }
    
    // Reason used by Order
    type Reason struct {

        
            DisplayName string  `json:"display_name"`
            QcType []string  `json:"qc_type"`
            QuestionSet []QuestionSet  `json:"question_set"`
            ID float64  `json:"id"`
         
    }
    
    // PlatformShipmentReasonsResponse used by Order
    type PlatformShipmentReasonsResponse struct {

        
            Success bool  `json:"success"`
            Reasons []Reason  `json:"reasons"`
         
    }
    
    // BulkActionPayload used by Order
    type BulkActionPayload struct {

        
            URL string  `json:"url"`
         
    }
    
    // BulkActionResponse used by Order
    type BulkActionResponse struct {

        
            Message string  `json:"message"`
            Status bool  `json:"status"`
         
    }
    
    // BulkActionDetailsDataField used by Order
    type BulkActionDetailsDataField struct {

        
            CompanyID string  `json:"company_id"`
            TotalShipmentsCount float64  `json:"total_shipments_count"`
            ProcessingShipmentsCount float64  `json:"processing_shipments_count"`
            SuccessfulShipmentsCount float64  `json:"successful_shipments_count"`
            SuccessfulShipmentIds []string  `json:"successful_shipment_ids"`
            BatchID string  `json:"batch_id"`
            FailedShipmentsCount float64  `json:"failed_shipments_count"`
         
    }
    
    // BulkActionDetailsResponse used by Order
    type BulkActionDetailsResponse struct {

        
            Error []string  `json:"error"`
            UploadedOn string  `json:"uploaded_on"`
            FailedRecords []string  `json:"failed_records"`
            Success string  `json:"success"`
            Message string  `json:"message"`
            UserID string  `json:"user_id"`
            UploadedBy string  `json:"uploaded_by"`
            Status bool  `json:"status"`
            Data []BulkActionDetailsDataField  `json:"data"`
         
    }
    
    // Brand used by Order
    type Brand struct {

        
            IsVirtualInvoice bool  `json:"is_virtual_invoice"`
            BrandID float64  `json:"brand_id"`
            CreatedOn float64  `json:"created_on"`
            CreditNoteExpiryDays float64  `json:"credit_note_expiry_days"`
            BrandName string  `json:"brand_name"`
            CreditNoteAllowed bool  `json:"credit_note_allowed"`
            Company string  `json:"company"`
            ModifiedOn float64  `json:"modified_on"`
            InvoicePrefix string  `json:"invoice_prefix"`
            PickupLocation string  `json:"pickup_location"`
            Logo string  `json:"logo"`
            StartDate string  `json:"start_date"`
            ScriptLastRan string  `json:"script_last_ran"`
         
    }
    
    // BagGSTDetails used by Order
    type BagGSTDetails struct {

        
            HsnCode string  `json:"hsn_code"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            ValueOfGood float64  `json:"value_of_good"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            IgstTaxPercentage float64  `json:"igst_tax_percentage"`
            SgstTaxPercentage float64  `json:"sgst_tax_percentage"`
            CgstGstFee string  `json:"cgst_gst_fee"`
            SgstGstFee string  `json:"sgst_gst_fee"`
            GstTag string  `json:"gst_tag"`
            GstinCode string  `json:"gstin_code"`
            CgstTaxPercentage float64  `json:"cgst_tax_percentage"`
            IgstGstFee string  `json:"igst_gst_fee"`
            IsDefaultHsnCode bool  `json:"is_default_hsn_code"`
            GstFee float64  `json:"gst_fee"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            HsnCodeID string  `json:"hsn_code_id"`
         
    }
    
    // B2BPODetails used by Order
    type B2BPODetails struct {

        
            PoLineAmount float64  `json:"po_line_amount"`
            DockerNumber string  `json:"docker_number"`
            ItemBasePrice float64  `json:"item_base_price"`
            PartialCanRet bool  `json:"partial_can_ret"`
            PoTaxAmount float64  `json:"po_tax_amount"`
            TotalGstPercentage float64  `json:"total_gst_percentage"`
         
    }
    
    // BagMeta used by Order
    type BagMeta struct {

        
            B2bPoDetails B2BPODetails  `json:"b2b_po_details"`
         
    }
    
    // Dates used by Order
    type Dates struct {

        
            DeliveryDate interface{}  `json:"delivery_date"`
            OrderCreated string  `json:"order_created"`
         
    }
    
    // StoreEwaybill used by Order
    type StoreEwaybill struct {

        
            Enabled bool  `json:"enabled"`
         
    }
    
    // StoreEinvoice used by Order
    type StoreEinvoice struct {

        
            Username string  `json:"username"`
            Password string  `json:"password"`
            User string  `json:"user"`
            Enabled bool  `json:"enabled"`
         
    }
    
    // StoreGstCredentials used by Order
    type StoreGstCredentials struct {

        
            EWaybill StoreEwaybill  `json:"e_waybill"`
            EInvoice StoreEinvoice  `json:"e_invoice"`
         
    }
    
    // EInvoicePortalDetails used by Order
    type EInvoicePortalDetails struct {

        
            Username string  `json:"username"`
            Password string  `json:"password"`
            User string  `json:"user"`
         
    }
    
    // Document used by Order
    type Document struct {

        
            URL string  `json:"url"`
            LegalName string  `json:"legal_name"`
            DsType string  `json:"ds_type"`
            Verified bool  `json:"verified"`
            Value string  `json:"value"`
         
    }
    
    // StoreDocuments used by Order
    type StoreDocuments struct {

        
            Gst Document  `json:"gst"`
         
    }
    
    // StoreMeta used by Order
    type StoreMeta struct {

        
            Stage string  `json:"stage"`
            EwaybillPortalDetails map[string]interface{}  `json:"ewaybill_portal_details"`
            DisplayName string  `json:"display_name"`
            GstCredentials StoreGstCredentials  `json:"gst_credentials"`
            GstNumber string  `json:"gst_number"`
            NotificationEmails []string  `json:"notification_emails"`
            EinvoicePortalDetails EInvoicePortalDetails  `json:"einvoice_portal_details"`
            AdditionalContactDetails map[string]interface{}  `json:"additional_contact_details"`
            Timing []map[string]interface{}  `json:"timing"`
            ProductReturnConfig map[string]interface{}  `json:"product_return_config"`
            Documents StoreDocuments  `json:"documents"`
         
    }
    
    // StoreAddress used by Order
    type StoreAddress struct {

        
            Landmark string  `json:"landmark"`
            Address2 string  `json:"address2"`
            Area string  `json:"area"`
            Latitude float64  `json:"latitude"`
            ContactPerson string  `json:"contact_person"`
            Country string  `json:"country"`
            AddressCategory string  `json:"address_category"`
            Email string  `json:"email"`
            CountryCode string  `json:"country_code"`
            Version string  `json:"version"`
            UpdatedAt string  `json:"updated_at"`
            AddressType string  `json:"address_type"`
            Address1 string  `json:"address1"`
            Phone string  `json:"phone"`
            CreatedAt string  `json:"created_at"`
            City string  `json:"city"`
            Pincode float64  `json:"pincode"`
            State string  `json:"state"`
            Longitude float64  `json:"longitude"`
         
    }
    
    // Store used by Order
    type Store struct {

        
            CompanyID float64  `json:"company_id"`
            FulfillmentChannel string  `json:"fulfillment_channel"`
            StoreActiveFrom string  `json:"store_active_from"`
            BrandStoreTags []string  `json:"brand_store_tags"`
            Code string  `json:"code"`
            Address2 string  `json:"address2"`
            Meta StoreMeta  `json:"meta"`
            MallArea string  `json:"mall_area"`
            StoreAddressJson StoreAddress  `json:"store_address_json"`
            Latitude float64  `json:"latitude"`
            LocationType string  `json:"location_type"`
            ContactPerson string  `json:"contact_person"`
            AlohomoraUserID float64  `json:"alohomora_user_id"`
            BrandID interface{}  `json:"brand_id"`
            Country string  `json:"country"`
            IsArchived bool  `json:"is_archived"`
            IsActive bool  `json:"is_active"`
            UpdatedAt string  `json:"updated_at"`
            IsEnabledForRecon bool  `json:"is_enabled_for_recon"`
            Address1 string  `json:"address1"`
            LoginUsername string  `json:"login_username"`
            StoreEmail string  `json:"store_email"`
            VatNo string  `json:"vat_no"`
            Phone float64  `json:"phone"`
            CreatedAt string  `json:"created_at"`
            OrderIntegrationID string  `json:"order_integration_id"`
            ParentStoreID float64  `json:"parent_store_id"`
            SID string  `json:"s_id"`
            MallName string  `json:"mall_name"`
            PackagingMaterialCount float64  `json:"packaging_material_count"`
            City string  `json:"city"`
            Name string  `json:"name"`
            Pincode string  `json:"pincode"`
            State string  `json:"state"`
            Longitude float64  `json:"longitude"`
         
    }
    
    // ArticleDetails used by Order
    type ArticleDetails struct {

        
            Status map[string]interface{}  `json:"status"`
         
    }
    
    // BagReturnableCancelableStatus used by Order
    type BagReturnableCancelableStatus struct {

        
            CanBeCancelled bool  `json:"can_be_cancelled"`
            IsCustomerReturnAllowed bool  `json:"is_customer_return_allowed"`
            IsActive bool  `json:"is_active"`
            EnableTracking bool  `json:"enable_tracking"`
            IsReturnable bool  `json:"is_returnable"`
         
    }
    
    // Weight used by Order
    type Weight struct {

        
            IsDefault bool  `json:"is_default"`
            Shipping float64  `json:"shipping"`
            Unit string  `json:"unit"`
         
    }
    
    // ReturnConfig used by Order
    type ReturnConfig struct {

        
            Returnable bool  `json:"returnable"`
            Time float64  `json:"time"`
            Unit string  `json:"unit"`
         
    }
    
    // Article used by Order
    type Article struct {

        
            SellerIdentifier string  `json:"seller_identifier"`
            UID string  `json:"uid"`
            RawMeta interface{}  `json:"raw_meta"`
            Weight Weight  `json:"weight"`
            Code string  `json:"code"`
            IsSet bool  `json:"is_set"`
            ReturnConfig ReturnConfig  `json:"return_config"`
            ID string  `json:"_id"`
            ChildDetails map[string]interface{}  `json:"child_details"`
            EspModified interface{}  `json:"esp_modified"`
            ASet map[string]interface{}  `json:"a_set"`
            Size string  `json:"size"`
            Identifiers Identifier  `json:"identifiers"`
            Dimensions Dimensions  `json:"dimensions"`
         
    }
    
    // Attributes used by Order
    type Attributes struct {

        
            PrimaryMaterial string  `json:"primary_material"`
            MarketerAddress string  `json:"marketer_address"`
            Gender []string  `json:"gender"`
            Essential string  `json:"essential"`
            BrandName string  `json:"brand_name"`
            MarketerName string  `json:"marketer_name"`
            PrimaryColorHex string  `json:"primary_color_hex"`
            Name string  `json:"name"`
            PrimaryColor string  `json:"primary_color"`
         
    }
    
    // Item used by Order
    type Item struct {

        
            Gender string  `json:"gender"`
            L2Category []string  `json:"l2_category"`
            BranchURL string  `json:"branch_url"`
            WebstoreProductURL string  `json:"webstore_product_url"`
            Brand string  `json:"brand"`
            Attributes Attributes  `json:"attributes"`
            Code string  `json:"code"`
            ItemID float64  `json:"item_id"`
            Size string  `json:"size"`
            Meta map[string]interface{}  `json:"meta"`
            CanCancel bool  `json:"can_cancel"`
            CanReturn bool  `json:"can_return"`
            SlugKey string  `json:"slug_key"`
            BrandID float64  `json:"brand_id"`
            L2CategoryID float64  `json:"l2_category_id"`
            Color string  `json:"color"`
            L1Category []string  `json:"l1_category"`
            L1CategoryID float64  `json:"l1_category_id"`
            L3CategoryName string  `json:"l3_category_name"`
            DepartmentID float64  `json:"department_id"`
            L3Category float64  `json:"l3_category"`
            Image []string  `json:"image"`
            Name string  `json:"name"`
            LastUpdatedAt string  `json:"last_updated_at"`
         
    }
    
    // AffiliateBagDetails used by Order
    type AffiliateBagDetails struct {

        
            AffiliateBagID string  `json:"affiliate_bag_id"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            EmployeeDiscount float64  `json:"employee_discount"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            AffiliateMeta AffiliateMeta  `json:"affiliate_meta"`
         
    }
    
    // BagDetailsPlatformResponse used by Order
    type BagDetailsPlatformResponse struct {

        
            Prices Prices  `json:"prices"`
            BID float64  `json:"b_id"`
            Brand Brand  `json:"brand"`
            GstDetails BagGSTDetails  `json:"gst_details"`
            Identifier string  `json:"identifier"`
            ShipmentID string  `json:"shipment_id"`
            CurrentStatus BagStatusHistory  `json:"current_status"`
            Quantity float64  `json:"quantity"`
            BType string  `json:"b_type"`
            Tags []string  `json:"tags"`
            JourneyType string  `json:"journey_type"`
            Meta BagMeta  `json:"meta"`
            Dates Dates  `json:"dates"`
            EntityType string  `json:"entity_type"`
            OriginalBagList []float64  `json:"original_bag_list"`
            SellerIdentifier string  `json:"seller_identifier"`
            BagStatusHistory BagStatusHistory  `json:"bag_status_history"`
            OperationalStatus string  `json:"operational_status"`
            OrderingStore Store  `json:"ordering_store"`
            CurrentOperationalStatus BagStatusHistory  `json:"current_operational_status"`
            BagUpdateTime float64  `json:"bag_update_time"`
            ArticleDetails ArticleDetails  `json:"article_details"`
            RestoreCoupon bool  `json:"restore_coupon"`
            AppliedPromos []map[string]interface{}  `json:"applied_promos"`
            NoOfBagsOrder float64  `json:"no_of_bags_order"`
            FinancialBreakup []FinancialBreakup  `json:"financial_breakup"`
            ParentPromoBags map[string]interface{}  `json:"parent_promo_bags"`
            RestorePromos map[string]interface{}  `json:"restore_promos"`
            Status BagReturnableCancelableStatus  `json:"status"`
            DisplayName string  `json:"display_name"`
            OrderIntegrationID string  `json:"order_integration_id"`
            Reasons []map[string]interface{}  `json:"reasons"`
            Article Article  `json:"article"`
            BagStatus []BagStatusHistory  `json:"bag_status"`
            AffiliateDetails AffiliateDetails  `json:"affiliate_details"`
            Item Item  `json:"item"`
            AffiliateBagDetails AffiliateBagDetails  `json:"affiliate_bag_details"`
            QcRequired interface{}  `json:"qc_required"`
            LineNumber float64  `json:"line_number"`
         
    }
    
    // ErrorResponse used by Order
    type ErrorResponse struct {

        
            Error string  `json:"error"`
            Message string  `json:"message"`
         
    }
    
    // Page1 used by Order
    type Page1 struct {

        
            HasNext bool  `json:"has_next"`
            PageType string  `json:"page_type"`
            Current float64  `json:"current"`
            Size float64  `json:"size"`
            ItemTotal float64  `json:"item_total"`
         
    }
    
    // GetBagsPlatformResponse used by Order
    type GetBagsPlatformResponse struct {

        
            Items []BagDetailsPlatformResponse  `json:"items"`
            Page Page1  `json:"page"`
         
    }
    
    // GeneratePosOrderReceiptResponse used by Order
    type GeneratePosOrderReceiptResponse struct {

        
            OrderID string  `json:"order_id"`
            Success bool  `json:"success"`
            PaymentReceipt string  `json:"payment_receipt"`
            InvoiceReceipt string  `json:"invoice_receipt"`
         
    }
    
    // InvalidateShipmentCachePayload used by Order
    type InvalidateShipmentCachePayload struct {

        
            ShipmentIds []string  `json:"shipment_ids"`
         
    }
    
    // InvalidateShipmentCacheNestedResponse used by Order
    type InvalidateShipmentCacheNestedResponse struct {

        
            Status float64  `json:"status"`
            Message string  `json:"message"`
            ShipmentID string  `json:"shipment_id"`
            Error string  `json:"error"`
         
    }
    
    // InvalidateShipmentCacheResponse used by Order
    type InvalidateShipmentCacheResponse struct {

        
            Response []InvalidateShipmentCacheNestedResponse  `json:"response"`
         
    }
    
    // ErrorResponse1 used by Order
    type ErrorResponse1 struct {

        
            Status float64  `json:"status"`
            Message string  `json:"message"`
            ErrorTrace string  `json:"error_trace"`
         
    }
    
    // StoreReassign used by Order
    type StoreReassign struct {

        
            SetID string  `json:"set_id"`
            AffiliateID string  `json:"affiliate_id"`
            FyndOrderID string  `json:"fynd_order_id"`
            ReasonIds []float64  `json:"reason_ids"`
            ItemID string  `json:"item_id"`
            MongoArticleID string  `json:"mongo_article_id"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            StoreID float64  `json:"store_id"`
            BagID float64  `json:"bag_id"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
         
    }
    
    // StoreReassignResponse used by Order
    type StoreReassignResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // Entities used by Order
    type Entities struct {

        
            AffiliateBagID string  `json:"affiliate_bag_id"`
            AffiliateID string  `json:"affiliate_id"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            ID string  `json:"id"`
            ReasonText string  `json:"reason_text"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
         
    }
    
    // UpdateShipmentLockPayload used by Order
    type UpdateShipmentLockPayload struct {

        
            Entities []Entities  `json:"entities"`
            EntityType string  `json:"entity_type"`
            ActionType string  `json:"action_type"`
            Action string  `json:"action"`
         
    }
    
    // OriginalFilter used by Order
    type OriginalFilter struct {

        
            AffiliateID string  `json:"affiliate_id"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
         
    }
    
    // Bags used by Order
    type Bags struct {

        
            AffiliateBagID string  `json:"affiliate_bag_id"`
            IsLocked bool  `json:"is_locked"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            BagID float64  `json:"bag_id"`
         
    }
    
    // CheckResponse used by Order
    type CheckResponse struct {

        
            AffiliateID string  `json:"affiliate_id"`
            LockStatus bool  `json:"lock_status"`
            IsShipmentLocked bool  `json:"is_shipment_locked"`
            OriginalFilter OriginalFilter  `json:"original_filter"`
            ShipmentID string  `json:"shipment_id"`
            IsBagLocked bool  `json:"is_bag_locked"`
            Status string  `json:"status"`
            Bags []Bags  `json:"bags"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
         
    }
    
    // UpdateShipmentLockResponse used by Order
    type UpdateShipmentLockResponse struct {

        
            CheckResponse []CheckResponse  `json:"check_response"`
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // AnnouncementResponse used by Order
    type AnnouncementResponse struct {

        
            FromDatetime string  `json:"from_datetime"`
            ToDatetime string  `json:"to_datetime"`
            CreatedAt string  `json:"created_at"`
            Description string  `json:"description"`
            PlatformName string  `json:"platform_name"`
            Title string  `json:"title"`
            LogoURL string  `json:"logo_url"`
            ID float64  `json:"id"`
            PlatformID string  `json:"platform_id"`
            CompanyID float64  `json:"company_id"`
         
    }
    
    // AnnouncementsResponse used by Order
    type AnnouncementsResponse struct {

        
            Announcements []AnnouncementResponse  `json:"announcements"`
         
    }
    
    // BaseResponse used by Order
    type BaseResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // Click2CallResponse used by Order
    type Click2CallResponse struct {

        
            Status bool  `json:"status"`
            CallID string  `json:"call_id"`
         
    }
    
    // EntityReasonData used by Order
    type EntityReasonData struct {

        
            ReasonID float64  `json:"reason_id"`
            ReasonText string  `json:"reason_text"`
         
    }
    
    // EntitiesReasons used by Order
    type EntitiesReasons struct {

        
            Filters []map[string]interface{}  `json:"filters"`
            Data EntityReasonData  `json:"data"`
         
    }
    
    // ProductsReasonsFilters used by Order
    type ProductsReasonsFilters struct {

        
            LineNumber float64  `json:"line_number"`
            Quantity float64  `json:"quantity"`
            Identifier string  `json:"identifier"`
         
    }
    
    // ProductsReasonsData used by Order
    type ProductsReasonsData struct {

        
            ReasonID float64  `json:"reason_id"`
            ReasonText string  `json:"reason_text"`
         
    }
    
    // ProductsReasons used by Order
    type ProductsReasons struct {

        
            Filters []ProductsReasonsFilters  `json:"filters"`
            Data ProductsReasonsData  `json:"data"`
         
    }
    
    // ReasonsData used by Order
    type ReasonsData struct {

        
            Entities []EntitiesReasons  `json:"entities"`
            Products []ProductsReasons  `json:"products"`
         
    }
    
    // EntitiesDataUpdates used by Order
    type EntitiesDataUpdates struct {

        
            Filters []map[string]interface{}  `json:"filters"`
            Data map[string]interface{}  `json:"data"`
         
    }
    
    // ProductsDataUpdatesFilters used by Order
    type ProductsDataUpdatesFilters struct {

        
            LineNumber float64  `json:"line_number"`
            Identifier string  `json:"identifier"`
         
    }
    
    // ProductsDataUpdates used by Order
    type ProductsDataUpdates struct {

        
            Filters []ProductsDataUpdatesFilters  `json:"filters"`
            Data map[string]interface{}  `json:"data"`
         
    }
    
    // DataUpdates used by Order
    type DataUpdates struct {

        
            Entities []EntitiesDataUpdates  `json:"entities"`
            Products []ProductsDataUpdates  `json:"products"`
         
    }
    
    // Products used by Order
    type Products struct {

        
            LineNumber float64  `json:"line_number"`
            Quantity float64  `json:"quantity"`
            Identifier string  `json:"identifier"`
         
    }
    
    // ShipmentsRequest used by Order
    type ShipmentsRequest struct {

        
            Reasons ReasonsData  `json:"reasons"`
            DataUpdates DataUpdates  `json:"data_updates"`
            Identifier string  `json:"identifier"`
            Products []Products  `json:"products"`
         
    }
    
    // StatuesRequest used by Order
    type StatuesRequest struct {

        
            Status string  `json:"status"`
            ExcludeBagsNextState string  `json:"exclude_bags_next_state"`
            Shipments []ShipmentsRequest  `json:"shipments"`
         
    }
    
    // UpdateShipmentStatusRequest used by Order
    type UpdateShipmentStatusRequest struct {

        
            LockAfterTransition bool  `json:"lock_after_transition"`
            UnlockBeforeTransition bool  `json:"unlock_before_transition"`
            Statuses []StatuesRequest  `json:"statuses"`
            Task bool  `json:"task"`
            ForceTransition bool  `json:"force_transition"`
         
    }
    
    // ShipmentsResponse used by Order
    type ShipmentsResponse struct {

        
            StackTrace string  `json:"stack_trace"`
            Exception string  `json:"exception"`
            Message string  `json:"message"`
            Code string  `json:"code"`
            Meta map[string]interface{}  `json:"meta"`
            Identifier string  `json:"identifier"`
            Status float64  `json:"status"`
            FinalState map[string]interface{}  `json:"final_state"`
         
    }
    
    // StatuesResponse used by Order
    type StatuesResponse struct {

        
            Shipments []ShipmentsResponse  `json:"shipments"`
         
    }
    
    // UpdateShipmentStatusResponseBody used by Order
    type UpdateShipmentStatusResponseBody struct {

        
            Statuses []StatuesResponse  `json:"statuses"`
         
    }
    
    // OrderPriority used by Order
    type OrderPriority struct {

        
            AffiliatePriorityCode string  `json:"affiliate_priority_code"`
            FulfilmentPriority float64  `json:"fulfilment_priority"`
            FulfilmentPriorityText string  `json:"fulfilment_priority_text"`
         
    }
    
    // OrderUser used by Order
    type OrderUser struct {

        
            State string  `json:"state"`
            Phone float64  `json:"phone"`
            Pincode string  `json:"pincode"`
            City string  `json:"city"`
            Address2 string  `json:"address2"`
            LastName string  `json:"last_name"`
            Address1 string  `json:"address1"`
            Email string  `json:"email"`
            Country string  `json:"country"`
            Mobile float64  `json:"mobile"`
            FirstName string  `json:"first_name"`
         
    }
    
    // ArticleDetails1 used by Order
    type ArticleDetails1 struct {

        
            ID string  `json:"_id"`
            BrandID float64  `json:"brand_id"`
            Category map[string]interface{}  `json:"category"`
            Dimension map[string]interface{}  `json:"dimension"`
            Weight map[string]interface{}  `json:"weight"`
            Quantity float64  `json:"quantity"`
            Attributes map[string]interface{}  `json:"attributes"`
         
    }
    
    // LocationDetails used by Order
    type LocationDetails struct {

        
            Articles []ArticleDetails1  `json:"articles"`
            FulfillmentType string  `json:"fulfillment_type"`
            FulfillmentID float64  `json:"fulfillment_id"`
         
    }
    
    // ShipmentDetails used by Order
    type ShipmentDetails struct {

        
            DpID float64  `json:"dp_id"`
            BoxType string  `json:"box_type"`
            FulfillmentID float64  `json:"fulfillment_id"`
            Meta map[string]interface{}  `json:"meta"`
            Articles []ArticleDetails1  `json:"articles"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            Shipments float64  `json:"shipments"`
         
    }
    
    // ShipmentConfig used by Order
    type ShipmentConfig struct {

        
            Source string  `json:"source"`
            LocationDetails LocationDetails  `json:"location_details"`
            ToPincode string  `json:"to_pincode"`
            Journey string  `json:"journey"`
            PaymentMode string  `json:"payment_mode"`
            Identifier string  `json:"identifier"`
            Shipment []ShipmentDetails  `json:"shipment"`
            Action string  `json:"action"`
         
    }
    
    // ShipmentData used by Order
    type ShipmentData struct {

        
            ShipmentData ShipmentConfig  `json:"shipment_data"`
         
    }
    
    // MarketPlacePdf used by Order
    type MarketPlacePdf struct {

        
            Invoice string  `json:"invoice"`
            Label string  `json:"label"`
         
    }
    
    // AffiliateBag used by Order
    type AffiliateBag struct {

        
            Discount float64  `json:"discount"`
            ModifiedOn string  `json:"modified_on"`
            ItemID float64  `json:"item_id"`
            PriceMarked float64  `json:"price_marked"`
            Identifier map[string]interface{}  `json:"identifier"`
            SellerIdentifier string  `json:"seller_identifier"`
            CompanyID float64  `json:"company_id"`
            ItemSize string  `json:"item_size"`
            ID string  `json:"_id"`
            AffiliateStoreID string  `json:"affiliate_store_id"`
            Sku string  `json:"sku"`
            FyndStoreID string  `json:"fynd_store_id"`
            HsnCodeID string  `json:"hsn_code_id"`
            PriceEffective float64  `json:"price_effective"`
            AvlQty float64  `json:"avl_qty"`
            AmountPaid float64  `json:"amount_paid"`
            StoreID float64  `json:"store_id"`
            DeliveryCharge float64  `json:"delivery_charge"`
            TransferPrice float64  `json:"transfer_price"`
            Quantity float64  `json:"quantity"`
            UnitPrice float64  `json:"unit_price"`
            AffiliateMeta map[string]interface{}  `json:"affiliate_meta"`
            PdfLinks MarketPlacePdf  `json:"pdf_links"`
         
    }
    
    // UserData used by Order
    type UserData struct {

        
            ShippingUser OrderUser  `json:"shipping_user"`
            BillingUser OrderUser  `json:"billing_user"`
         
    }
    
    // OrderInfo used by Order
    type OrderInfo struct {

        
            OrderPriority OrderPriority  `json:"order_priority"`
            OrderValue float64  `json:"order_value"`
            DeliveryCharges float64  `json:"delivery_charges"`
            Items map[string]interface{}  `json:"items"`
            Discount float64  `json:"discount"`
            Payment map[string]interface{}  `json:"payment"`
            CodCharges float64  `json:"cod_charges"`
            PaymentMode string  `json:"payment_mode"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            ShippingAddress OrderUser  `json:"shipping_address"`
            BillingAddress OrderUser  `json:"billing_address"`
            Shipment ShipmentData  `json:"shipment"`
            Bags []AffiliateBag  `json:"bags"`
            User UserData  `json:"user"`
            Coupon string  `json:"coupon"`
         
    }
    
    // AffiliateInventoryPaymentConfig used by Order
    type AffiliateInventoryPaymentConfig struct {

        
            Source string  `json:"source"`
            ModeOfPayment string  `json:"mode_of_payment"`
         
    }
    
    // AffiliateInventoryArticleAssignmentConfig used by Order
    type AffiliateInventoryArticleAssignmentConfig struct {

        
            PostOrderReassignment bool  `json:"post_order_reassignment"`
         
    }
    
    // AffiliateInventoryLogisticsConfig used by Order
    type AffiliateInventoryLogisticsConfig struct {

        
            DpAssignment bool  `json:"dp_assignment"`
         
    }
    
    // AffiliateInventoryStoreConfig used by Order
    type AffiliateInventoryStoreConfig struct {

        
            Store map[string]interface{}  `json:"store"`
         
    }
    
    // AffiliateInventoryOrderConfig used by Order
    type AffiliateInventoryOrderConfig struct {

        
            ForceReassignment bool  `json:"force_reassignment"`
         
    }
    
    // AffiliateInventoryConfig used by Order
    type AffiliateInventoryConfig struct {

        
            Payment AffiliateInventoryPaymentConfig  `json:"payment"`
            ArticleAssignment AffiliateInventoryArticleAssignmentConfig  `json:"article_assignment"`
            Logistics AffiliateInventoryLogisticsConfig  `json:"logistics"`
            Inventory AffiliateInventoryStoreConfig  `json:"inventory"`
            Order AffiliateInventoryOrderConfig  `json:"order"`
         
    }
    
    // AffiliateAppConfigMeta used by Order
    type AffiliateAppConfigMeta struct {

        
            Value string  `json:"value"`
            Name string  `json:"name"`
         
    }
    
    // AffiliateAppConfig used by Order
    type AffiliateAppConfig struct {

        
            Secret string  `json:"secret"`
            CreatedAt string  `json:"created_at"`
            Meta []AffiliateAppConfigMeta  `json:"meta"`
            Owner string  `json:"owner"`
            Description string  `json:"description"`
            Token string  `json:"token"`
            ID string  `json:"id"`
            Name string  `json:"name"`
            UpdatedAt string  `json:"updated_at"`
         
    }
    
    // AffiliateConfig used by Order
    type AffiliateConfig struct {

        
            Inventory AffiliateInventoryConfig  `json:"inventory"`
            App AffiliateAppConfig  `json:"app"`
         
    }
    
    // Affiliate used by Order
    type Affiliate struct {

        
            ID string  `json:"id"`
            Config AffiliateConfig  `json:"config"`
            Token string  `json:"token"`
         
    }
    
    // AffiliateStoreIdMapping used by Order
    type AffiliateStoreIdMapping struct {

        
            StoreID float64  `json:"store_id"`
            MarketplaceStoreID string  `json:"marketplace_store_id"`
         
    }
    
    // OrderConfig used by Order
    type OrderConfig struct {

        
            Affiliate Affiliate  `json:"affiliate"`
            ArticleLookup string  `json:"article_lookup"`
            StoreLookup string  `json:"store_lookup"`
            BagEndState string  `json:"bag_end_state"`
            AffiliateStoreIDMapping []AffiliateStoreIdMapping  `json:"affiliate_store_id_mapping"`
            CreateUser bool  `json:"create_user"`
         
    }
    
    // CreateOrderPayload used by Order
    type CreateOrderPayload struct {

        
            AffiliateID string  `json:"affiliate_id"`
            OrderInfo OrderInfo  `json:"order_info"`
            OrderConfig OrderConfig  `json:"order_config"`
         
    }
    
    // CreateOrderResponse used by Order
    type CreateOrderResponse struct {

        
            FyndOrderID string  `json:"fynd_order_id"`
         
    }
    
    // DispatchManifest used by Order
    type DispatchManifest struct {

        
            ManifestID string  `json:"manifest_id"`
         
    }
    
    // SuccessResponse used by Order
    type SuccessResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // ActionInfo used by Order
    type ActionInfo struct {

        
            ID float64  `json:"id"`
            Description string  `json:"description"`
            DisplayText string  `json:"display_text"`
            Slug string  `json:"slug"`
         
    }
    
    // GetActionsResponse used by Order
    type GetActionsResponse struct {

        
            Permissions ActionInfo  `json:"permissions"`
         
    }
    
    // HistoryDict used by Order
    type HistoryDict struct {

        
            Message string  `json:"message"`
            L2Detail string  `json:"l2_detail"`
            Createdat string  `json:"createdat"`
            BagID float64  `json:"bag_id"`
            TicketID string  `json:"ticket_id"`
            L1Detail string  `json:"l1_detail"`
            L3Detail string  `json:"l3_detail"`
            TicketURL string  `json:"ticket_url"`
            Type string  `json:"type"`
            User string  `json:"user"`
         
    }
    
    // ShipmentHistoryResponse used by Order
    type ShipmentHistoryResponse struct {

        
            ActivityHistory []HistoryDict  `json:"activity_history"`
         
    }
    
    // ErrorDetail used by Order
    type ErrorDetail struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // PostHistoryFilters used by Order
    type PostHistoryFilters struct {

        
            LineNumber string  `json:"line_number"`
            Identifier string  `json:"identifier"`
            ShipmentID string  `json:"shipment_id"`
         
    }
    
    // PostHistoryData used by Order
    type PostHistoryData struct {

        
            Message string  `json:"message"`
            UserName string  `json:"user_name"`
         
    }
    
    // PostActivityHistory used by Order
    type PostActivityHistory struct {

        
            Filters []PostHistoryFilters  `json:"filters"`
            Data PostHistoryData  `json:"data"`
         
    }
    
    // PostHistoryDict used by Order
    type PostHistoryDict struct {

        
            ActivityHistory PostActivityHistory  `json:"activity_history"`
         
    }
    
    // PostShipmentHistory used by Order
    type PostShipmentHistory struct {

        
            ActivityHistory []PostHistoryDict  `json:"activity_history"`
         
    }
    
    // SmsDataPayload used by Order
    type SmsDataPayload struct {

        
            CustomerName string  `json:"customer_name"`
            Message string  `json:"message"`
            PhoneNumber float64  `json:"phone_number"`
            BrandName string  `json:"brand_name"`
            PaymentMode string  `json:"payment_mode"`
            OrderID string  `json:"order_id"`
            AmountPaid float64  `json:"amount_paid"`
            ShipmentID float64  `json:"shipment_id"`
            CountryCode string  `json:"country_code"`
         
    }
    
    // SendSmsPayload used by Order
    type SendSmsPayload struct {

        
            Data SmsDataPayload  `json:"data"`
            Slug string  `json:"slug"`
            BagID float64  `json:"bag_id"`
         
    }
    
    // Meta1 used by Order
    type Meta1 struct {

        
            StateManagerUsed string  `json:"state_manager_used"`
            KafkaEmissionStatus float64  `json:"kafka_emission_status"`
         
    }
    
    // ShipmentDetail used by Order
    type ShipmentDetail struct {

        
            ID float64  `json:"id"`
            Remarks string  `json:"remarks"`
            Meta Meta1  `json:"meta"`
            BagList []float64  `json:"bag_list"`
            ShipmentID string  `json:"shipment_id"`
            Status string  `json:"status"`
         
    }
    
    // OrderDetails used by Order
    type OrderDetails struct {

        
            CreatedAt string  `json:"created_at"`
            FyndOrderID string  `json:"fynd_order_id"`
         
    }
    
    // OrderStatusData used by Order
    type OrderStatusData struct {

        
            ShipmentDetails []ShipmentDetail  `json:"shipment_details"`
            Errors []string  `json:"errors"`
            OrderDetails OrderDetails  `json:"order_details"`
         
    }
    
    // OrderStatusResult used by Order
    type OrderStatusResult struct {

        
            Success string  `json:"success"`
            Result []OrderStatusData  `json:"result"`
         
    }
    
    // ManualAssignDPToShipment used by Order
    type ManualAssignDPToShipment struct {

        
            QcRequired string  `json:"qc_required"`
            OrderType string  `json:"order_type"`
            DpID float64  `json:"dp_id"`
            ShipmentIds []string  `json:"shipment_ids"`
         
    }
    
    // ManualAssignDPToShipmentResponse used by Order
    type ManualAssignDPToShipmentResponse struct {

        
            Errors []string  `json:"errors"`
            Success string  `json:"success"`
         
    }
    
    // ShippingInfo used by Order
    type ShippingInfo struct {

        
            PrimaryMobileNumber string  `json:"primary_mobile_number"`
            AlternateEmail string  `json:"alternate_email"`
            StateCode string  `json:"state_code"`
            Pincode string  `json:"pincode"`
            AlternateMobileNumber string  `json:"alternate_mobile_number"`
            Address2 string  `json:"address2"`
            CustomerCode string  `json:"customer_code"`
            Title string  `json:"title"`
            Country string  `json:"country"`
            CountryCode string  `json:"country_code"`
            Gender string  `json:"gender"`
            FloorNo string  `json:"floor_no"`
            State string  `json:"state"`
            Slot []map[string]interface{}  `json:"slot"`
            ExternalCustomerCode string  `json:"external_customer_code"`
            Landmark string  `json:"landmark"`
            AddressType string  `json:"address_type"`
            City string  `json:"city"`
            PrimaryEmail string  `json:"primary_email"`
            LastName string  `json:"last_name"`
            GeoLocation map[string]interface{}  `json:"geo_location"`
            Address1 string  `json:"address1"`
            MiddleName string  `json:"middle_name"`
            HouseNo string  `json:"house_no"`
            FirstName string  `json:"first_name"`
            ShippingType string  `json:"shipping_type"`
         
    }
    
    // Tax used by Order
    type Tax struct {

        
            Name string  `json:"name"`
            Rate float64  `json:"rate"`
            Amount map[string]interface{}  `json:"amount"`
            Breakup []map[string]interface{}  `json:"breakup"`
         
    }
    
    // Charge used by Order
    type Charge struct {

        
            Tax Tax  `json:"tax"`
            Code string  `json:"code"`
            Amount map[string]interface{}  `json:"amount"`
            Type string  `json:"type"`
            Name string  `json:"name"`
         
    }
    
    // TaxInfo used by Order
    type TaxInfo struct {

        
            B2bGstinNumber string  `json:"b2b_gstin_number"`
            Gstin string  `json:"gstin"`
         
    }
    
    // BillingInfo used by Order
    type BillingInfo struct {

        
            PrimaryMobileNumber string  `json:"primary_mobile_number"`
            AlternateEmail string  `json:"alternate_email"`
            StateCode string  `json:"state_code"`
            Pincode string  `json:"pincode"`
            AlternateMobileNumber string  `json:"alternate_mobile_number"`
            Address2 string  `json:"address2"`
            CustomerCode string  `json:"customer_code"`
            Title string  `json:"title"`
            Country string  `json:"country"`
            CountryCode string  `json:"country_code"`
            Gender string  `json:"gender"`
            FloorNo string  `json:"floor_no"`
            State string  `json:"state"`
            ExternalCustomerCode string  `json:"external_customer_code"`
            City string  `json:"city"`
            PrimaryEmail string  `json:"primary_email"`
            LastName string  `json:"last_name"`
            Address1 string  `json:"address1"`
            MiddleName string  `json:"middle_name"`
            HouseNo string  `json:"house_no"`
            FirstName string  `json:"first_name"`
         
    }
    
    // PaymentMethod used by Order
    type PaymentMethod struct {

        
            Mode string  `json:"mode"`
            TransactionData map[string]interface{}  `json:"transaction_data"`
            RefundBy string  `json:"refund_by"`
            Meta map[string]interface{}  `json:"meta"`
            Amount float64  `json:"amount"`
            CollectBy string  `json:"collect_by"`
            Name string  `json:"name"`
         
    }
    
    // PaymentInfo used by Order
    type PaymentInfo struct {

        
            PaymentMethods []PaymentMethod  `json:"payment_methods"`
            PrimaryMode string  `json:"primary_mode"`
         
    }
    
    // ProcessingDates used by Order
    type ProcessingDates struct {

        
            ConfirmByDate string  `json:"confirm_by_date"`
            PackByDate string  `json:"pack_by_date"`
            DispatchByDate string  `json:"dispatch_by_date"`
            CustomerPickupSlot map[string]interface{}  `json:"customer_pickup_slot"`
            DpPickupSlot map[string]interface{}  `json:"dp_pickup_slot"`
            DispatchAfterDate string  `json:"dispatch_after_date"`
         
    }
    
    // LineItem used by Order
    type LineItem struct {

        
            Charges []Charge  `json:"charges"`
            Meta map[string]interface{}  `json:"meta"`
            CustomMessasge string  `json:"custom_messasge"`
            SellerIdentifier string  `json:"seller_identifier"`
            Quantity float64  `json:"quantity"`
            ExternalLineID string  `json:"external_line_id"`
         
    }
    
    // Shipment used by Order
    type Shipment struct {

        
            ProcessingDates ProcessingDates  `json:"processing_dates"`
            ExternalShipmentID string  `json:"external_shipment_id"`
            Meta map[string]interface{}  `json:"meta"`
            Priority float64  `json:"priority"`
            LocationID float64  `json:"location_id"`
            LineItems []LineItem  `json:"line_items"`
         
    }
    
    // CreateOrderAPI used by Order
    type CreateOrderAPI struct {

        
            CurrencyInfo map[string]interface{}  `json:"currency_info"`
            ShippingInfo ShippingInfo  `json:"shipping_info"`
            Charges []Charge  `json:"charges"`
            TaxInfo TaxInfo  `json:"tax_info"`
            Meta map[string]interface{}  `json:"meta"`
            BillingInfo BillingInfo  `json:"billing_info"`
            ExternalOrderID string  `json:"external_order_id"`
            PaymentInfo PaymentInfo  `json:"payment_info"`
            ExternalCreationDate string  `json:"external_creation_date"`
            Shipments []Shipment  `json:"shipments"`
         
    }
    
    // CreateOrderErrorReponse used by Order
    type CreateOrderErrorReponse struct {

        
            StackTrace string  `json:"stack_trace"`
            Message string  `json:"message"`
            Exception string  `json:"exception"`
            Info interface{}  `json:"info"`
            Code string  `json:"code"`
            Meta string  `json:"meta"`
            RequestID string  `json:"request_id"`
            Status float64  `json:"status"`
         
    }
    
    // DpConfiguration used by Order
    type DpConfiguration struct {

        
            ShippingBy string  `json:"shipping_by"`
         
    }
    
    // PaymentMethods used by Order
    type PaymentMethods struct {

        
            RefundBy string  `json:"refund_by"`
            Mode string  `json:"mode"`
            CollectBy string  `json:"collect_by"`
         
    }
    
    // CreateChannelPaymentInfo used by Order
    type CreateChannelPaymentInfo struct {

        
            Source string  `json:"source"`
            PaymentMethods []PaymentMethods  `json:"payment_methods"`
            ModeOfPayment string  `json:"mode_of_payment"`
         
    }
    
    // CreateChannelConfig used by Order
    type CreateChannelConfig struct {

        
            DpConfiguration DpConfiguration  `json:"dp_configuration"`
            ShipmentAssignment string  `json:"shipment_assignment"`
            LocationReassignment bool  `json:"location_reassignment"`
            LogoURL map[string]interface{}  `json:"logo_url"`
            PaymentInfo CreateChannelPaymentInfo  `json:"payment_info"`
            LockStates []string  `json:"lock_states"`
         
    }
    
    // CreateChannelConfigData used by Order
    type CreateChannelConfigData struct {

        
            ConfigData CreateChannelConfig  `json:"config_data"`
         
    }
    
    // CreateChannelConifgErrorResponse used by Order
    type CreateChannelConifgErrorResponse struct {

        
            Error string  `json:"error"`
         
    }
    
    // CreateChannelConfigResponse used by Order
    type CreateChannelConfigResponse struct {

        
            Acknowledged bool  `json:"acknowledged"`
            IsUpserted bool  `json:"is_upserted"`
            IsInserted bool  `json:"is_inserted"`
         
    }
    
    // UploadConsent used by Order
    type UploadConsent struct {

        
            ManifestID string  `json:"manifest_id"`
            ConsentURL string  `json:"consent_url"`
         
    }
    
    // PlatformOrderUpdate used by Order
    type PlatformOrderUpdate struct {

        
            OrderID string  `json:"order_id"`
         
    }
    
    // ResponseDetail used by Order
    type ResponseDetail struct {

        
            Message []string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // FyndOrderIdList used by Order
    type FyndOrderIdList struct {

        
            FyndOrderID []string  `json:"fynd_order_id"`
         
    }
    
    // OrderStatus used by Order
    type OrderStatus struct {

        
            Mobile float64  `json:"mobile"`
            StartDate string  `json:"start_date"`
            OrderDetails []FyndOrderIdList  `json:"order_details"`
            EndDate string  `json:"end_date"`
         
    }
    

    
    // DeleteResponse used by Catalog
    type DeleteResponse struct {

        
            Message string  `json:"message"`
         
    }
    
    // ErrorResponse used by Catalog
    type ErrorResponse struct {

        
            Message string  `json:"message"`
            Meta map[string]interface{}  `json:"meta"`
            Status float64  `json:"status"`
            Code string  `json:"code"`
            Error string  `json:"error"`
         
    }
    
    // Page used by Catalog
    type Page struct {

        
            ItemTotal float64  `json:"item_total"`
            NextID string  `json:"next_id"`
            HasPrevious bool  `json:"has_previous"`
            HasNext bool  `json:"has_next"`
            Current float64  `json:"current"`
            Type string  `json:"type"`
            Size float64  `json:"size"`
         
    }
    
    // GetSearchWordsData used by Catalog
    type GetSearchWordsData struct {

        
            UID string  `json:"uid"`
            Result map[string]interface{}  `json:"result"`
            Words []string  `json:"words"`
            IsActive bool  `json:"is_active"`
            AppID string  `json:"app_id"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // GetSearchWordsDetailResponse used by Catalog
    type GetSearchWordsDetailResponse struct {

        
            Page Page  `json:"page"`
            Items GetSearchWordsData  `json:"items"`
         
    }
    
    // SearchKeywordResult used by Catalog
    type SearchKeywordResult struct {

        
            SortOn string  `json:"sort_on"`
            Query map[string]interface{}  `json:"query"`
         
    }
    
    // CreateSearchKeyword used by Catalog
    type CreateSearchKeyword struct {

        
            Result SearchKeywordResult  `json:"result"`
            Words []string  `json:"words"`
            IsActive bool  `json:"is_active"`
            AppID string  `json:"app_id"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // GetSearchWordsResponse used by Catalog
    type GetSearchWordsResponse struct {

        
            Page Page  `json:"page"`
            Items []GetSearchWordsData  `json:"items"`
         
    }
    
    // GetAutocompleteWordsData used by Catalog
    type GetAutocompleteWordsData struct {

        
            UID string  `json:"uid"`
            Words []string  `json:"words"`
            AppID string  `json:"app_id"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Results []map[string]interface{}  `json:"results"`
         
    }
    
    // GetAutocompleteWordsResponse used by Catalog
    type GetAutocompleteWordsResponse struct {

        
            Page Page  `json:"page"`
            Items []GetAutocompleteWordsData  `json:"items"`
         
    }
    
    // Media used by Catalog
    type Media struct {

        
            AspectRatio string  `json:"aspect_ratio"`
            Type string  `json:"type"`
            URL string  `json:"url"`
         
    }
    
    // AutocompletePageAction used by Catalog
    type AutocompletePageAction struct {

        
            Params map[string]interface{}  `json:"params"`
            Type string  `json:"type"`
            Query map[string]interface{}  `json:"query"`
            URL string  `json:"url"`
         
    }
    
    // AutocompleteAction used by Catalog
    type AutocompleteAction struct {

        
            Page AutocompletePageAction  `json:"page"`
            Type string  `json:"type"`
         
    }
    
    // AutocompleteResult used by Catalog
    type AutocompleteResult struct {

        
            Logo Media  `json:"logo"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Action AutocompleteAction  `json:"action"`
            Display string  `json:"display"`
         
    }
    
    // CreateAutocompleteKeyword used by Catalog
    type CreateAutocompleteKeyword struct {

        
            Words []string  `json:"words"`
            IsActive bool  `json:"is_active"`
            AppID string  `json:"app_id"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Results []AutocompleteResult  `json:"results"`
         
    }
    
    // CreateAutocompleteWordsResponse used by Catalog
    type CreateAutocompleteWordsResponse struct {

        
            Words []string  `json:"words"`
            Results []map[string]interface{}  `json:"results"`
            AppID string  `json:"app_id"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // ProductBundleItem used by Catalog
    type ProductBundleItem struct {

        
            ProductUID float64  `json:"product_uid"`
            AutoAddToCart bool  `json:"auto_add_to_cart"`
            AutoSelect bool  `json:"auto_select"`
            MaxQuantity float64  `json:"max_quantity"`
            AllowRemove bool  `json:"allow_remove"`
            MinQuantity float64  `json:"min_quantity"`
         
    }
    
    // GetProductBundleCreateResponse used by Catalog
    type GetProductBundleCreateResponse struct {

        
            Name string  `json:"name"`
            Products []ProductBundleItem  `json:"products"`
            ModifiedOn string  `json:"modified_on"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            Meta map[string]interface{}  `json:"meta"`
            PageVisibility []string  `json:"page_visibility"`
            Choice string  `json:"choice"`
            Slug string  `json:"slug"`
            CreatedOn string  `json:"created_on"`
            ID string  `json:"id"`
            IsActive bool  `json:"is_active"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Logo string  `json:"logo"`
            CompanyID float64  `json:"company_id"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
         
    }
    
    // GetProductBundleListingResponse used by Catalog
    type GetProductBundleListingResponse struct {

        
            Page Page  `json:"page"`
            Items []GetProductBundleCreateResponse  `json:"items"`
         
    }
    
    // ProductBundleRequest used by Catalog
    type ProductBundleRequest struct {

        
            Name string  `json:"name"`
            Products []ProductBundleItem  `json:"products"`
            ModifiedOn string  `json:"modified_on"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            Meta map[string]interface{}  `json:"meta"`
            PageVisibility []string  `json:"page_visibility"`
            Choice string  `json:"choice"`
            Slug string  `json:"slug"`
            CreatedOn string  `json:"created_on"`
            IsActive bool  `json:"is_active"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Logo string  `json:"logo"`
            CompanyID float64  `json:"company_id"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
         
    }
    
    // Size used by Catalog
    type Size struct {

        
            IsAvailable bool  `json:"is_available"`
            Quantity float64  `json:"quantity"`
            Value string  `json:"value"`
            Display string  `json:"display"`
         
    }
    
    // LimitedProductData used by Catalog
    type LimitedProductData struct {

        
            Identifier map[string]interface{}  `json:"identifier"`
            UID float64  `json:"uid"`
            ShortDescription string  `json:"short_description"`
            Sizes []string  `json:"sizes"`
            Name string  `json:"name"`
            Price map[string]interface{}  `json:"price"`
            Slug string  `json:"slug"`
            Images []string  `json:"images"`
            CountryOfOrigin string  `json:"country_of_origin"`
            Attributes map[string]interface{}  `json:"attributes"`
            Quantity float64  `json:"quantity"`
            ItemCode string  `json:"item_code"`
         
    }
    
    // Price used by Catalog
    type Price struct {

        
            MinEffective float64  `json:"min_effective"`
            MaxEffective float64  `json:"max_effective"`
            MaxMarked float64  `json:"max_marked"`
            Currency string  `json:"currency"`
            MinMarked float64  `json:"min_marked"`
         
    }
    
    // GetProducts used by Catalog
    type GetProducts struct {

        
            ProductUID float64  `json:"product_uid"`
            Sizes []Size  `json:"sizes"`
            AutoAddToCart bool  `json:"auto_add_to_cart"`
            ProductDetails LimitedProductData  `json:"product_details"`
            AutoSelect bool  `json:"auto_select"`
            AllowRemove bool  `json:"allow_remove"`
            MaxQuantity float64  `json:"max_quantity"`
            MinQuantity float64  `json:"min_quantity"`
            Price Price  `json:"price"`
         
    }
    
    // GetProductBundleResponse used by Catalog
    type GetProductBundleResponse struct {

        
            Products []GetProducts  `json:"products"`
            Name string  `json:"name"`
            PageVisibility []string  `json:"page_visibility"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            Meta map[string]interface{}  `json:"meta"`
            Slug string  `json:"slug"`
            Logo string  `json:"logo"`
            IsActive bool  `json:"is_active"`
            Choice string  `json:"choice"`
            CompanyID float64  `json:"company_id"`
         
    }
    
    // ProductBundleUpdateRequest used by Catalog
    type ProductBundleUpdateRequest struct {

        
            Name string  `json:"name"`
            Products []ProductBundleItem  `json:"products"`
            ModifiedOn string  `json:"modified_on"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            Meta map[string]interface{}  `json:"meta"`
            PageVisibility []string  `json:"page_visibility"`
            Choice string  `json:"choice"`
            Slug string  `json:"slug"`
            IsActive bool  `json:"is_active"`
            Logo string  `json:"logo"`
            CompanyID float64  `json:"company_id"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
         
    }
    
    // ListSizeGuide used by Catalog
    type ListSizeGuide struct {

        
            Page map[string]interface{}  `json:"page"`
            Items []map[string]interface{}  `json:"items"`
         
    }
    
    // Meta used by Catalog
    type Meta struct {

        
            Headers map[string]interface{}  `json:"headers"`
            Values []map[string]interface{}  `json:"values"`
            Unit string  `json:"unit"`
         
    }
    
    // Guide used by Catalog
    type Guide struct {

        
            Meta Meta  `json:"meta"`
         
    }
    
    // ValidateSizeGuide used by Catalog
    type ValidateSizeGuide struct {

        
            Guide Guide  `json:"guide"`
            Name string  `json:"name"`
            ID string  `json:"id"`
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
            Subtitle string  `json:"subtitle"`
            Title string  `json:"title"`
            Tag string  `json:"tag"`
            Image string  `json:"image"`
            BrandID float64  `json:"brand_id"`
            Active bool  `json:"active"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Description string  `json:"description"`
            CompanyID float64  `json:"company_id"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
         
    }
    
    // SuccessResponse used by Catalog
    type SuccessResponse struct {

        
            Success bool  `json:"success"`
         
    }
    
    // SizeGuideResponse used by Catalog
    type SizeGuideResponse struct {

        
            Guide map[string]interface{}  `json:"guide"`
            Name string  `json:"name"`
            ID string  `json:"id"`
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
            Subtitle string  `json:"subtitle"`
            Title string  `json:"title"`
            Tag string  `json:"tag"`
            BrandID float64  `json:"brand_id"`
            Active bool  `json:"active"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            CompanyID float64  `json:"company_id"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
         
    }
    
    // MOQData used by Catalog
    type MOQData struct {

        
            IncrementUnit float64  `json:"increment_unit"`
            Maximum float64  `json:"maximum"`
            Minimum float64  `json:"minimum"`
         
    }
    
    // SEOData used by Catalog
    type SEOData struct {

        
            Description string  `json:"description"`
            Title string  `json:"title"`
         
    }
    
    // OwnerAppItemResponse used by Catalog
    type OwnerAppItemResponse struct {

        
            IsCod bool  `json:"is_cod"`
            AltText map[string]interface{}  `json:"alt_text"`
            Moq MOQData  `json:"moq"`
            Seo SEOData  `json:"seo"`
            IsGift bool  `json:"is_gift"`
         
    }
    
    // MetaFields used by Catalog
    type MetaFields struct {

        
            Key string  `json:"key"`
            Value string  `json:"value"`
         
    }
    
    // ApplicationItemMOQ used by Catalog
    type ApplicationItemMOQ struct {

        
            IncrementUnit float64  `json:"increment_unit"`
            Maximum float64  `json:"maximum"`
            Minimum float64  `json:"minimum"`
         
    }
    
    // ApplicationItemSEO used by Catalog
    type ApplicationItemSEO struct {

        
            Description string  `json:"description"`
            Title string  `json:"title"`
         
    }
    
    // ApplicationItemMeta used by Catalog
    type ApplicationItemMeta struct {

        
            CustomMeta []MetaFields  `json:"_custom_meta"`
            IsCod bool  `json:"is_cod"`
            AltText map[string]interface{}  `json:"alt_text"`
            Moq ApplicationItemMOQ  `json:"moq"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Seo ApplicationItemSEO  `json:"seo"`
            IsGift bool  `json:"is_gift"`
         
    }
    
    // SuccessResponse1 used by Catalog
    type SuccessResponse1 struct {

        
            UID float64  `json:"uid"`
            Success bool  `json:"success"`
         
    }
    
    // GetConfigMetadataResponse used by Catalog
    type GetConfigMetadataResponse struct {

        
            Values []map[string]interface{}  `json:"values"`
            Data []map[string]interface{}  `json:"data"`
            Condition []map[string]interface{}  `json:"condition"`
         
    }
    
    // PageResponseType used by Catalog
    type PageResponseType struct {

        
            TotalCount float64  `json:"total_count"`
            HasNext bool  `json:"has_next"`
            Current float64  `json:"current"`
            Next float64  `json:"next"`
         
    }
    
    // GetConfigResponse used by Catalog
    type GetConfigResponse struct {

        
            Page PageResponseType  `json:"page"`
            Data []map[string]interface{}  `json:"data"`
         
    }
    
    // ConfigErrorResponse used by Catalog
    type ConfigErrorResponse struct {

        
            Message string  `json:"message"`
         
    }
    
    // AttributeDetailsGroup used by Catalog
    type AttributeDetailsGroup struct {

        
            Name string  `json:"name"`
            Unit string  `json:"unit"`
            DisplayType string  `json:"display_type"`
            Slug string  `json:"slug"`
            Key string  `json:"key"`
            IsActive bool  `json:"is_active"`
            Logo string  `json:"logo"`
            Priority float64  `json:"priority"`
         
    }
    
    // AppConfigurationDetail used by Catalog
    type AppConfigurationDetail struct {

        
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Attributes []AttributeDetailsGroup  `json:"attributes"`
            TemplateSlugs []string  `json:"template_slugs"`
            IsActive bool  `json:"is_active"`
            AppID string  `json:"app_id"`
            IsDefault bool  `json:"is_default"`
            Logo string  `json:"logo"`
            Priority float64  `json:"priority"`
         
    }
    
    // ConfigSuccessResponse used by Catalog
    type ConfigSuccessResponse struct {

        
            Message string  `json:"message"`
         
    }
    
    // AppConfigurationsSort used by Catalog
    type AppConfigurationsSort struct {

        
            Name string  `json:"name"`
            DefaultKey string  `json:"default_key"`
            Key string  `json:"key"`
            IsActive bool  `json:"is_active"`
            AppID string  `json:"app_id"`
            IsDefault bool  `json:"is_default"`
            Logo string  `json:"logo"`
            Priority float64  `json:"priority"`
         
    }
    
    // AllowSingleRequest used by Catalog
    type AllowSingleRequest struct {

        
            AllowSingle bool  `json:"allow_single"`
         
    }
    
    // DefaultKeyRequest used by Catalog
    type DefaultKeyRequest struct {

        
            DefaultKey string  `json:"default_key"`
         
    }
    
    // MetaDataListingSortMetaResponse used by Catalog
    type MetaDataListingSortMetaResponse struct {

        
            Key string  `json:"key"`
            Display string  `json:"display"`
         
    }
    
    // MetaDataListingSortResponse used by Catalog
    type MetaDataListingSortResponse struct {

        
            Data []MetaDataListingSortMetaResponse  `json:"data"`
         
    }
    
    // MetaDataListingFilterMetaResponse used by Catalog
    type MetaDataListingFilterMetaResponse struct {

        
            Units []map[string]interface{}  `json:"units"`
            Key string  `json:"key"`
            FilterTypes []string  `json:"filter_types"`
            Display string  `json:"display"`
         
    }
    
    // MetaDataListingFilterResponse used by Catalog
    type MetaDataListingFilterResponse struct {

        
            Data []MetaDataListingFilterMetaResponse  `json:"data"`
         
    }
    
    // MetaDataListingResponse used by Catalog
    type MetaDataListingResponse struct {

        
            Sort MetaDataListingSortResponse  `json:"sort"`
            Filter MetaDataListingFilterResponse  `json:"filter"`
         
    }
    
    // GetCatalogConfigurationDetailsProduct used by Catalog
    type GetCatalogConfigurationDetailsProduct struct {

        
            Compare map[string]interface{}  `json:"compare"`
            Similar map[string]interface{}  `json:"similar"`
            Variant map[string]interface{}  `json:"variant"`
            Detail map[string]interface{}  `json:"detail"`
         
    }
    
    // GetCatalogConfigurationMetaData used by Catalog
    type GetCatalogConfigurationMetaData struct {

        
            Listing MetaDataListingResponse  `json:"listing"`
            Product GetCatalogConfigurationDetailsProduct  `json:"product"`
         
    }
    
    // ConfigurationListingSortConfig used by Catalog
    type ConfigurationListingSortConfig struct {

        
            Name string  `json:"name"`
            Key string  `json:"key"`
            IsActive bool  `json:"is_active"`
            Logo string  `json:"logo"`
            Priority float64  `json:"priority"`
         
    }
    
    // ConfigurationListingSort used by Catalog
    type ConfigurationListingSort struct {

        
            Config []ConfigurationListingSortConfig  `json:"config"`
            DefaultKey string  `json:"default_key"`
         
    }
    
    // ConfigurationBucketPoints used by Catalog
    type ConfigurationBucketPoints struct {

        
            Display string  `json:"display"`
            Start float64  `json:"start"`
            End float64  `json:"end"`
         
    }
    
    // ConfigurationListingFilterValue used by Catalog
    type ConfigurationListingFilterValue struct {

        
            Sort string  `json:"sort"`
            MapValues []map[string]interface{}  `json:"map_values"`
            BucketPoints []ConfigurationBucketPoints  `json:"bucket_points"`
            Condition string  `json:"condition"`
            Map map[string]interface{}  `json:"map"`
            Value string  `json:"value"`
         
    }
    
    // ConfigurationListingFilterConfig used by Catalog
    type ConfigurationListingFilterConfig struct {

        
            Name string  `json:"name"`
            ValueConfig ConfigurationListingFilterValue  `json:"value_config"`
            Key string  `json:"key"`
            IsActive bool  `json:"is_active"`
            Type string  `json:"type"`
            Logo string  `json:"logo"`
            Priority float64  `json:"priority"`
            DisplayName string  `json:"display_name"`
         
    }
    
    // ConfigurationListingFilter used by Catalog
    type ConfigurationListingFilter struct {

        
            AllowSingle bool  `json:"allow_single"`
            AttributeConfig []ConfigurationListingFilterConfig  `json:"attribute_config"`
         
    }
    
    // ConfigurationListing used by Catalog
    type ConfigurationListing struct {

        
            Sort ConfigurationListingSort  `json:"sort"`
            Filter ConfigurationListingFilter  `json:"filter"`
         
    }
    
    // ProductSize used by Catalog
    type ProductSize struct {

        
            Min float64  `json:"min"`
            Max float64  `json:"max"`
         
    }
    
    // ConfigurationProductConfig used by Catalog
    type ConfigurationProductConfig struct {

        
            Subtitle string  `json:"subtitle"`
            Title string  `json:"title"`
            Size ProductSize  `json:"size"`
            Key string  `json:"key"`
            IsActive bool  `json:"is_active"`
            Logo string  `json:"logo"`
            Priority float64  `json:"priority"`
         
    }
    
    // ConfigurationProductSimilar used by Catalog
    type ConfigurationProductSimilar struct {

        
            Config []ConfigurationProductConfig  `json:"config"`
         
    }
    
    // ConfigurationProductVariantConfig used by Catalog
    type ConfigurationProductVariantConfig struct {

        
            Name string  `json:"name"`
            DisplayType string  `json:"display_type"`
            Size ProductSize  `json:"size"`
            Key string  `json:"key"`
            IsActive bool  `json:"is_active"`
            Logo string  `json:"logo"`
            Priority float64  `json:"priority"`
         
    }
    
    // ConfigurationProductVariant used by Catalog
    type ConfigurationProductVariant struct {

        
            Config []ConfigurationProductVariantConfig  `json:"config"`
         
    }
    
    // ConfigurationProduct used by Catalog
    type ConfigurationProduct struct {

        
            Similar ConfigurationProductSimilar  `json:"similar"`
            Variant ConfigurationProductVariant  `json:"variant"`
         
    }
    
    // AppCatalogConfiguration used by Catalog
    type AppCatalogConfiguration struct {

        
            Listing ConfigurationListing  `json:"listing"`
            ID string  `json:"id"`
            ModifiedOn string  `json:"modified_on"`
            ConfigType string  `json:"config_type"`
            CreatedOn string  `json:"created_on"`
            ConfigID string  `json:"config_id"`
            Type string  `json:"type"`
            AppID string  `json:"app_id"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Product ConfigurationProduct  `json:"product"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
         
    }
    
    // GetAppCatalogConfiguration used by Catalog
    type GetAppCatalogConfiguration struct {

        
            Data AppCatalogConfiguration  `json:"data"`
            IsDefault bool  `json:"is_default"`
         
    }
    
    // AppConfiguration used by Catalog
    type AppConfiguration struct {

        
            Listing ConfigurationListing  `json:"listing"`
            ModifiedOn string  `json:"modified_on"`
            ConfigType string  `json:"config_type"`
            CreatedOn string  `json:"created_on"`
            ConfigID string  `json:"config_id"`
            Type string  `json:"type"`
            AppID string  `json:"app_id"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Product ConfigurationProduct  `json:"product"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
         
    }
    
    // GetCatalogConfigurationDetailsSchemaListing used by Catalog
    type GetCatalogConfigurationDetailsSchemaListing struct {

        
            Sort map[string]interface{}  `json:"sort"`
            Filter map[string]interface{}  `json:"filter"`
         
    }
    
    // EntityConfiguration used by Catalog
    type EntityConfiguration struct {

        
            Listing GetCatalogConfigurationDetailsSchemaListing  `json:"listing"`
            ID string  `json:"id"`
            ConfigType string  `json:"config_type"`
            ConfigID string  `json:"config_id"`
            AppID string  `json:"app_id"`
            Product GetCatalogConfigurationDetailsProduct  `json:"product"`
         
    }
    
    // GetAppCatalogEntityConfiguration used by Catalog
    type GetAppCatalogEntityConfiguration struct {

        
            Data EntityConfiguration  `json:"data"`
            IsDefault bool  `json:"is_default"`
         
    }
    
    // ProductSortOn used by Catalog
    type ProductSortOn struct {

        
            Name string  `json:"name"`
            Value string  `json:"value"`
            IsSelected bool  `json:"is_selected"`
         
    }
    
    // ProductFiltersKey used by Catalog
    type ProductFiltersKey struct {

        
            Name string  `json:"name"`
            Operators []string  `json:"operators"`
            Display string  `json:"display"`
            Logo string  `json:"logo"`
            Kind string  `json:"kind"`
         
    }
    
    // ProductFiltersValue used by Catalog
    type ProductFiltersValue struct {

        
            Min float64  `json:"min"`
            IsSelected bool  `json:"is_selected"`
            Max float64  `json:"max"`
            SelectedMax float64  `json:"selected_max"`
            CurrencySymbol string  `json:"currency_symbol"`
            SelectedMin float64  `json:"selected_min"`
            Count float64  `json:"count"`
            DisplayFormat string  `json:"display_format"`
            Display string  `json:"display"`
            Value interface{}  `json:"value"`
            CurrencyCode string  `json:"currency_code"`
            QueryFormat string  `json:"query_format"`
         
    }
    
    // ProductFilters used by Catalog
    type ProductFilters struct {

        
            Key ProductFiltersKey  `json:"key"`
            Values []ProductFiltersValue  `json:"values"`
         
    }
    
    // GetCollectionQueryOptionResponse used by Catalog
    type GetCollectionQueryOptionResponse struct {

        
            SortOn []ProductSortOn  `json:"sort_on"`
            Operators map[string]string  `json:"operators"`
            Filters []ProductFilters  `json:"filters"`
         
    }
    
    // CollectionQuery used by Catalog
    type CollectionQuery struct {

        
            Op string  `json:"op"`
            Attribute string  `json:"attribute"`
            Value []interface{}  `json:"value"`
         
    }
    
    // BannerImage used by Catalog
    type BannerImage struct {

        
            AspectRatio string  `json:"aspect_ratio"`
            Type string  `json:"type"`
            URL string  `json:"url"`
         
    }
    
    // ImageUrls used by Catalog
    type ImageUrls struct {

        
            Landscape BannerImage  `json:"landscape"`
            Portrait BannerImage  `json:"portrait"`
         
    }
    
    // Media1 used by Catalog
    type Media1 struct {

        
            Meta map[string]interface{}  `json:"meta"`
            Type string  `json:"type"`
            URL string  `json:"url"`
         
    }
    
    // ActionPage used by Catalog
    type ActionPage struct {

        
            Params map[string][]string  `json:"params"`
            Query map[string][]string  `json:"query"`
            URL string  `json:"url"`
            Type PageType  `json:"type"`
         
    }
    
    // Action used by Catalog
    type Action struct {

        
            Page ActionPage  `json:"page"`
            Popup ActionPage  `json:"popup"`
            Type string  `json:"type"`
         
    }
    
    // GetCollectionDetailNest used by Catalog
    type GetCollectionDetailNest struct {

        
            UID string  `json:"uid"`
            Query []CollectionQuery  `json:"query"`
            AllowSort bool  `json:"allow_sort"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Tag []string  `json:"tag"`
            Schedule map[string]interface{}  `json:"_schedule"`
            AllowFacets bool  `json:"allow_facets"`
            Cron map[string]interface{}  `json:"cron"`
            Banners ImageUrls  `json:"banners"`
            Type string  `json:"type"`
            Logo Media1  `json:"logo"`
            Priority float64  `json:"priority"`
            Description string  `json:"description"`
            Name string  `json:"name"`
            Badge map[string]interface{}  `json:"badge"`
            Meta map[string]interface{}  `json:"meta"`
            Slug string  `json:"slug"`
            Action Action  `json:"action"`
            IsActive bool  `json:"is_active"`
            AppID string  `json:"app_id"`
         
    }
    
    // CollectionListingFilterType used by Catalog
    type CollectionListingFilterType struct {

        
            Name string  `json:"name"`
            IsSelected bool  `json:"is_selected"`
            Display string  `json:"display"`
         
    }
    
    // CollectionListingFilterTag used by Catalog
    type CollectionListingFilterTag struct {

        
            Name string  `json:"name"`
            IsSelected bool  `json:"is_selected"`
            Display string  `json:"display"`
         
    }
    
    // CollectionListingFilter used by Catalog
    type CollectionListingFilter struct {

        
            Type []CollectionListingFilterType  `json:"type"`
            Tags []CollectionListingFilterTag  `json:"tags"`
         
    }
    
    // GetCollectionListingResponse used by Catalog
    type GetCollectionListingResponse struct {

        
            Page Page  `json:"page"`
            Items []GetCollectionDetailNest  `json:"items"`
            Filters CollectionListingFilter  `json:"filters"`
         
    }
    
    // SeoDetail used by Catalog
    type SeoDetail struct {

        
            Description string  `json:"description"`
            Title string  `json:"title"`
         
    }
    
    // UserInfo used by Catalog
    type UserInfo struct {

        
            Username string  `json:"username"`
            Email string  `json:"email"`
            UserID string  `json:"user_id"`
            UID string  `json:"uid"`
         
    }
    
    // NextSchedule used by Catalog
    type NextSchedule struct {

        
            Start string  `json:"start"`
            End string  `json:"end"`
         
    }
    
    // CollectionSchedule used by Catalog
    type CollectionSchedule struct {

        
            Duration float64  `json:"duration"`
            NextSchedule []NextSchedule  `json:"next_schedule"`
            Start string  `json:"start"`
            End string  `json:"end"`
            Cron string  `json:"cron"`
         
    }
    
    // CollectionImage used by Catalog
    type CollectionImage struct {

        
            AspectRatio string  `json:"aspect_ratio"`
            URL string  `json:"url"`
         
    }
    
    // CollectionBanner used by Catalog
    type CollectionBanner struct {

        
            Landscape CollectionImage  `json:"landscape"`
            Portrait CollectionImage  `json:"portrait"`
         
    }
    
    // CollectionBadge used by Catalog
    type CollectionBadge struct {

        
            Text string  `json:"text"`
            Color string  `json:"color"`
         
    }
    
    // CreateCollection used by Catalog
    type CreateCollection struct {

        
            Query []CollectionQuery  `json:"query"`
            AllowSort bool  `json:"allow_sort"`
            Tags []string  `json:"tags"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Seo SeoDetail  `json:"seo"`
            ModifiedBy UserInfo  `json:"modified_by"`
            IsVisible bool  `json:"is_visible"`
            Schedule CollectionSchedule  `json:"_schedule"`
            AllowFacets bool  `json:"allow_facets"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            Banners CollectionBanner  `json:"banners"`
            Type string  `json:"type"`
            CreatedBy UserInfo  `json:"created_by"`
            Logo CollectionImage  `json:"logo"`
            Description string  `json:"description"`
            Priority float64  `json:"priority"`
            Name string  `json:"name"`
            Badge CollectionBadge  `json:"badge"`
            Meta map[string]interface{}  `json:"meta"`
            Slug string  `json:"slug"`
            SortOn string  `json:"sort_on"`
            IsActive bool  `json:"is_active"`
            AppID string  `json:"app_id"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Published bool  `json:"published"`
         
    }
    
    // CollectionCreateResponse used by Catalog
    type CollectionCreateResponse struct {

        
            Query []CollectionQuery  `json:"query"`
            AllowSort bool  `json:"allow_sort"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Tag []string  `json:"tag"`
            Schedule map[string]interface{}  `json:"_schedule"`
            AllowFacets bool  `json:"allow_facets"`
            Cron map[string]interface{}  `json:"cron"`
            Banners ImageUrls  `json:"banners"`
            Type string  `json:"type"`
            Logo BannerImage  `json:"logo"`
            Priority float64  `json:"priority"`
            Description string  `json:"description"`
            Name string  `json:"name"`
            Badge map[string]interface{}  `json:"badge"`
            Meta map[string]interface{}  `json:"meta"`
            Slug string  `json:"slug"`
            SortOn string  `json:"sort_on"`
            IsActive bool  `json:"is_active"`
            AppID string  `json:"app_id"`
         
    }
    
    // CollectionDetailResponse used by Catalog
    type CollectionDetailResponse struct {

        
            UID string  `json:"uid"`
            Query []CollectionQuery  `json:"query"`
            AllowSort bool  `json:"allow_sort"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Tag []string  `json:"tag"`
            Schedule map[string]interface{}  `json:"_schedule"`
            AllowFacets bool  `json:"allow_facets"`
            Cron map[string]interface{}  `json:"cron"`
            Banners ImageUrls  `json:"banners"`
            Type string  `json:"type"`
            Logo Media1  `json:"logo"`
            Priority float64  `json:"priority"`
            Description string  `json:"description"`
            Name string  `json:"name"`
            Badge map[string]interface{}  `json:"badge"`
            Meta map[string]interface{}  `json:"meta"`
            Slug string  `json:"slug"`
            IsActive bool  `json:"is_active"`
            AppID string  `json:"app_id"`
         
    }
    
    // UpdateCollection used by Catalog
    type UpdateCollection struct {

        
            Query []CollectionQuery  `json:"query"`
            AllowSort bool  `json:"allow_sort"`
            Tags []string  `json:"tags"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Seo SeoDetail  `json:"seo"`
            ModifiedBy UserInfo  `json:"modified_by"`
            IsVisible bool  `json:"is_visible"`
            Schedule CollectionSchedule  `json:"_schedule"`
            AllowFacets bool  `json:"allow_facets"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            Banners CollectionBanner  `json:"banners"`
            Type string  `json:"type"`
            Logo CollectionImage  `json:"logo"`
            Description string  `json:"description"`
            Priority float64  `json:"priority"`
            Name string  `json:"name"`
            Badge CollectionBadge  `json:"badge"`
            Meta map[string]interface{}  `json:"meta"`
            Slug string  `json:"slug"`
            SortOn string  `json:"sort_on"`
            IsActive bool  `json:"is_active"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Published bool  `json:"published"`
         
    }
    
    // ProductBrand used by Catalog
    type ProductBrand struct {

        
            UID float64  `json:"uid"`
            Action Action  `json:"action"`
            Name string  `json:"name"`
            Logo Media1  `json:"logo"`
         
    }
    
    // Price1 used by Catalog
    type Price1 struct {

        
            CurrencySymbol string  `json:"currency_symbol"`
            CurrencyCode string  `json:"currency_code"`
            Min float64  `json:"min"`
            Max float64  `json:"max"`
         
    }
    
    // ProductListingPrice used by Catalog
    type ProductListingPrice struct {

        
            Marked Price1  `json:"marked"`
            Effective Price1  `json:"effective"`
         
    }
    
    // ProductDetailAttribute used by Catalog
    type ProductDetailAttribute struct {

        
            Key string  `json:"key"`
            Type string  `json:"type"`
            Value string  `json:"value"`
         
    }
    
    // ProductDetailGroupedAttribute used by Catalog
    type ProductDetailGroupedAttribute struct {

        
            Details []ProductDetailAttribute  `json:"details"`
            Title string  `json:"title"`
         
    }
    
    // ProductListingDetail used by Catalog
    type ProductListingDetail struct {

        
            UID float64  `json:"uid"`
            Similars []string  `json:"similars"`
            RatingCount float64  `json:"rating_count"`
            Medias []Media1  `json:"medias"`
            Discount string  `json:"discount"`
            ImageNature string  `json:"image_nature"`
            Highlights []string  `json:"highlights"`
            Rating float64  `json:"rating"`
            Color string  `json:"color"`
            Tryouts []string  `json:"tryouts"`
            Brand ProductBrand  `json:"brand"`
            HasVariant bool  `json:"has_variant"`
            ItemCode string  `json:"item_code"`
            Price ProductListingPrice  `json:"price"`
            PromoMeta map[string]interface{}  `json:"promo_meta"`
            Type string  `json:"type"`
            Description string  `json:"description"`
            ShortDescription string  `json:"short_description"`
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
            Name string  `json:"name"`
            ProductOnlineDate string  `json:"product_online_date"`
            Slug string  `json:"slug"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            Sellable bool  `json:"sellable"`
            Attributes map[string]interface{}  `json:"attributes"`
            ItemType string  `json:"item_type"`
         
    }
    
    // GetCollectionItemsResponse used by Catalog
    type GetCollectionItemsResponse struct {

        
            SortOn []ProductSortOn  `json:"sort_on"`
            Page Page  `json:"page"`
            Items []ProductListingDetail  `json:"items"`
            Filters []ProductFilters  `json:"filters"`
         
    }
    
    // ItemQueryForUserCollection used by Catalog
    type ItemQueryForUserCollection struct {

        
            ItemID float64  `json:"item_id"`
            Action string  `json:"action"`
         
    }
    
    // CollectionItemRequest used by Catalog
    type CollectionItemRequest struct {

        
            Item []ItemQueryForUserCollection  `json:"item"`
            Type string  `json:"type"`
            Query []CollectionQuery  `json:"query"`
         
    }
    
    // UpdatedResponse used by Catalog
    type UpdatedResponse struct {

        
            Message string  `json:"message"`
            ItemsNotUpdated []float64  `json:"items_not_updated"`
         
    }
    
    // CatalogInsightItem used by Catalog
    type CatalogInsightItem struct {

        
            OutOfStockCount float64  `json:"out_of_stock_count"`
            SellableCount float64  `json:"sellable_count"`
            Count float64  `json:"count"`
         
    }
    
    // CatalogInsightBrand used by Catalog
    type CatalogInsightBrand struct {

        
            Name string  `json:"name"`
            TotalArticles float64  `json:"total_articles"`
            AvailableArticles float64  `json:"available_articles"`
            ArticleFreshness float64  `json:"article_freshness"`
            AvailableSizes float64  `json:"available_sizes"`
            TotalSizes float64  `json:"total_sizes"`
         
    }
    
    // CatalogInsightResponse used by Catalog
    type CatalogInsightResponse struct {

        
            Item CatalogInsightItem  `json:"item"`
            BrandDistribution CatalogInsightBrand  `json:"brand_distribution"`
         
    }
    
    // CrossSellingData used by Catalog
    type CrossSellingData struct {

        
            Products float64  `json:"products"`
            Articles float64  `json:"articles"`
         
    }
    
    // CrossSellingResponse used by Catalog
    type CrossSellingResponse struct {

        
            BrandDistribution CatalogInsightBrand  `json:"brand_distribution"`
            Data CrossSellingData  `json:"data"`
         
    }
    
    // OptInPostRequest used by Catalog
    type OptInPostRequest struct {

        
            OptLevel string  `json:"opt_level"`
            Enabled bool  `json:"enabled"`
            StoreIds []float64  `json:"store_ids"`
            BrandIds []float64  `json:"brand_ids"`
            CompanyID float64  `json:"company_id"`
            Platform string  `json:"platform"`
         
    }
    
    // CompanyOptIn used by Catalog
    type CompanyOptIn struct {

        
            OptLevel string  `json:"opt_level"`
            ModifiedOn float64  `json:"modified_on"`
            Enabled bool  `json:"enabled"`
            CreatedOn float64  `json:"created_on"`
            StoreIds []float64  `json:"store_ids"`
            BrandIds []float64  `json:"brand_ids"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            CompanyID float64  `json:"company_id"`
            Platform string  `json:"platform"`
         
    }
    
    // GetOptInPlatform used by Catalog
    type GetOptInPlatform struct {

        
            Page Page  `json:"page"`
            Items []CompanyOptIn  `json:"items"`
         
    }
    
    // OptinCompanyDetail used by Catalog
    type OptinCompanyDetail struct {

        
            UID float64  `json:"uid"`
            BusinessType string  `json:"business_type"`
            Name string  `json:"name"`
            CompanyType string  `json:"company_type"`
         
    }
    
    // CompanyBrandDetail used by Catalog
    type CompanyBrandDetail struct {

        
            BrandID float64  `json:"brand_id"`
            TotalArticle float64  `json:"total_article"`
            CompanyID float64  `json:"company_id"`
            BrandName string  `json:"brand_name"`
         
    }
    
    // OptinCompanyBrandDetailsView used by Catalog
    type OptinCompanyBrandDetailsView struct {

        
            Page Page  `json:"page"`
            Items []CompanyBrandDetail  `json:"items"`
         
    }
    
    // OptinCompanyMetrics used by Catalog
    type OptinCompanyMetrics struct {

        
            Store float64  `json:"store"`
            Company string  `json:"company"`
            Brand float64  `json:"brand"`
         
    }
    
    // StoreDetail used by Catalog
    type StoreDetail struct {

        
            UID float64  `json:"uid"`
            Address map[string]interface{}  `json:"address"`
            Manager map[string]interface{}  `json:"manager"`
            Name string  `json:"name"`
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
            CompanyID float64  `json:"company_id"`
            StoreCode string  `json:"store_code"`
            Documents []map[string]interface{}  `json:"documents"`
            StoreType string  `json:"store_type"`
            Timing map[string]interface{}  `json:"timing"`
            AdditionalContacts []map[string]interface{}  `json:"additional_contacts"`
            DisplayName string  `json:"display_name"`
         
    }
    
    // OptinStoreDetails used by Catalog
    type OptinStoreDetails struct {

        
            Page Page  `json:"page"`
            Items []StoreDetail  `json:"items"`
         
    }
    
    // AttributeMasterMandatoryDetails used by Catalog
    type AttributeMasterMandatoryDetails struct {

        
            L3Keys []string  `json:"l3_keys"`
         
    }
    
    // AttributeMasterMeta used by Catalog
    type AttributeMasterMeta struct {

        
            MandatoryDetails AttributeMasterMandatoryDetails  `json:"mandatory_details"`
            Enriched bool  `json:"enriched"`
         
    }
    
    // AttributeMasterDetails used by Catalog
    type AttributeMasterDetails struct {

        
            DisplayType string  `json:"display_type"`
         
    }
    
    // AttributeMasterFilter used by Catalog
    type AttributeMasterFilter struct {

        
            Priority float64  `json:"priority"`
            DependsOn []string  `json:"depends_on"`
            Indexing bool  `json:"indexing"`
         
    }
    
    // AttributeSchemaRange used by Catalog
    type AttributeSchemaRange struct {

        
            Min float64  `json:"min"`
            Max float64  `json:"max"`
         
    }
    
    // AttributeMaster used by Catalog
    type AttributeMaster struct {

        
            Multi bool  `json:"multi"`
            AllowedValues []string  `json:"allowed_values"`
            Type string  `json:"type"`
            Format string  `json:"format"`
            Range AttributeSchemaRange  `json:"range"`
            Mandatory bool  `json:"mandatory"`
         
    }
    
    // GenderDetail used by Catalog
    type GenderDetail struct {

        
            EnabledForEndConsumer bool  `json:"enabled_for_end_consumer"`
            Name string  `json:"name"`
            ID string  `json:"id"`
            Meta AttributeMasterMeta  `json:"meta"`
            Details AttributeMasterDetails  `json:"details"`
            Slug string  `json:"slug"`
            IsNested bool  `json:"is_nested"`
            Filters AttributeMasterFilter  `json:"filters"`
            Departments []string  `json:"departments"`
            Logo string  `json:"logo"`
            Description string  `json:"description"`
            Schema AttributeMaster  `json:"schema"`
         
    }
    
    // ProductTemplateCategory used by Catalog
    type ProductTemplateCategory struct {

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            SlugKey string  `json:"slug_key"`
            Slug string  `json:"slug"`
            TemplateSlug string  `json:"template_slug"`
         
    }
    
    // ProdcutTemplateCategoriesResponse used by Catalog
    type ProdcutTemplateCategoriesResponse struct {

        
            Items []ProductTemplateCategory  `json:"items"`
         
    }
    
    // PTErrorResponse used by Catalog
    type PTErrorResponse struct {

        
            Message string  `json:"message"`
            Meta map[string]interface{}  `json:"meta"`
            Status float64  `json:"status"`
            Code string  `json:"code"`
            Errors map[string]interface{}  `json:"errors"`
         
    }
    
    // UserSerializer used by Catalog
    type UserSerializer struct {

        
            ID string  `json:"_id"`
            UID string  `json:"uid"`
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
            Contact string  `json:"contact"`
         
    }
    
    // GetDepartment used by Catalog
    type GetDepartment struct {

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            PageNo float64  `json:"page_no"`
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
            Slug string  `json:"slug"`
            PriorityOrder float64  `json:"priority_order"`
            Synonyms []string  `json:"synonyms"`
            Search string  `json:"search"`
            IsActive bool  `json:"is_active"`
            PageSize float64  `json:"page_size"`
            CreatedBy UserSerializer  `json:"created_by"`
            Logo string  `json:"logo"`
            ItemType string  `json:"item_type"`
            ModifiedBy UserSerializer  `json:"modified_by"`
         
    }
    
    // DepartmentsResponse used by Catalog
    type DepartmentsResponse struct {

        
            Page Page  `json:"page"`
            Items []GetDepartment  `json:"items"`
         
    }
    
    // DepartmentErrorResponse used by Catalog
    type DepartmentErrorResponse struct {

        
            Message string  `json:"message"`
            Meta map[string]interface{}  `json:"meta"`
            Status float64  `json:"status"`
            Code string  `json:"code"`
            Errors map[string]interface{}  `json:"errors"`
         
    }
    
    // DepartmentCreateUpdate used by Catalog
    type DepartmentCreateUpdate struct {

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            PriorityOrder float64  `json:"priority_order"`
            Slug string  `json:"slug"`
            Cls string  `json:"_cls"`
            Synonyms []string  `json:"synonyms"`
            IsActive bool  `json:"is_active"`
            Tags []string  `json:"tags"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Logo string  `json:"logo"`
            Platforms map[string]interface{}  `json:"platforms"`
         
    }
    
    // DepartmentCreateResponse used by Catalog
    type DepartmentCreateResponse struct {

        
            Message string  `json:"message"`
            UID float64  `json:"uid"`
         
    }
    
    // DepartmentCreateErrorResponse used by Catalog
    type DepartmentCreateErrorResponse struct {

        
            Error string  `json:"error"`
         
    }
    
    // UserDetail used by Catalog
    type UserDetail struct {

        
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
            Contact string  `json:"contact"`
            SuperUser bool  `json:"super_user"`
         
    }
    
    // DepartmentModel used by Catalog
    type DepartmentModel struct {

        
            ID interface{}  `json:"_id"`
            UID float64  `json:"uid"`
            VerifiedBy UserDetail  `json:"verified_by"`
            Name string  `json:"name"`
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
            Slug string  `json:"slug"`
            VerifiedOn string  `json:"verified_on"`
            PriorityOrder float64  `json:"priority_order"`
            Cls string  `json:"_cls"`
            Synonyms []string  `json:"synonyms"`
            IsActive bool  `json:"is_active"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            CreatedBy UserDetail  `json:"created_by"`
            Logo string  `json:"logo"`
            ModifiedBy UserDetail  `json:"modified_by"`
         
    }
    
    // ProductTemplate used by Catalog
    type ProductTemplate struct {

        
            IsPhysical bool  `json:"is_physical"`
            Name string  `json:"name"`
            IsArchived bool  `json:"is_archived"`
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
            Slug string  `json:"slug"`
            Tag string  `json:"tag"`
            IsExpirable bool  `json:"is_expirable"`
            Departments []string  `json:"departments"`
            IsActive bool  `json:"is_active"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Attributes []string  `json:"attributes"`
            Categories []string  `json:"categories"`
            Logo string  `json:"logo"`
            Description string  `json:"description"`
            CreatedBy map[string]interface{}  `json:"created_by"`
         
    }
    
    // TemplatesResponse used by Catalog
    type TemplatesResponse struct {

        
            Page Page  `json:"page"`
            Items ProductTemplate  `json:"items"`
         
    }
    
    // TemplateDetails used by Catalog
    type TemplateDetails struct {

        
            IsPhysical bool  `json:"is_physical"`
            Name string  `json:"name"`
            IsArchived bool  `json:"is_archived"`
            ID string  `json:"id"`
            Slug string  `json:"slug"`
            Tag string  `json:"tag"`
            IsExpirable bool  `json:"is_expirable"`
            Departments []string  `json:"departments"`
            IsActive bool  `json:"is_active"`
            Attributes []string  `json:"attributes"`
            Categories []string  `json:"categories"`
            Logo string  `json:"logo"`
            Description string  `json:"description"`
         
    }
    
    // Properties used by Catalog
    type Properties struct {

        
            Trader map[string]interface{}  `json:"trader"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            NoOfBoxes map[string]interface{}  `json:"no_of_boxes"`
            CountryOfOrigin map[string]interface{}  `json:"country_of_origin"`
            Tags map[string]interface{}  `json:"tags"`
            Highlights map[string]interface{}  `json:"highlights"`
            Media map[string]interface{}  `json:"media"`
            Command map[string]interface{}  `json:"command"`
            MultiSize map[string]interface{}  `json:"multi_size"`
            ReturnConfig map[string]interface{}  `json:"return_config"`
            CategorySlug map[string]interface{}  `json:"category_slug"`
            ItemCode map[string]interface{}  `json:"item_code"`
            Variants map[string]interface{}  `json:"variants"`
            HsnCode map[string]interface{}  `json:"hsn_code"`
            IsDependent map[string]interface{}  `json:"is_dependent"`
            ProductGroupTag map[string]interface{}  `json:"product_group_tag"`
            Currency map[string]interface{}  `json:"currency"`
            SizeGuide map[string]interface{}  `json:"size_guide"`
            Description map[string]interface{}  `json:"description"`
            ShortDescription map[string]interface{}  `json:"short_description"`
            Sizes map[string]interface{}  `json:"sizes"`
            Name map[string]interface{}  `json:"name"`
            TraderType map[string]interface{}  `json:"trader_type"`
            BrandUID map[string]interface{}  `json:"brand_uid"`
            Slug map[string]interface{}  `json:"slug"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            IsActive map[string]interface{}  `json:"is_active"`
            ItemType map[string]interface{}  `json:"item_type"`
            ProductPublish map[string]interface{}  `json:"product_publish"`
         
    }
    
    // GlobalValidation used by Catalog
    type GlobalValidation struct {

        
            Title string  `json:"title"`
            Required []string  `json:"required"`
            Definitions map[string]interface{}  `json:"definitions"`
            Description string  `json:"description"`
            Type string  `json:"type"`
            Properties Properties  `json:"properties"`
         
    }
    
    // TemplateValidationData used by Catalog
    type TemplateValidationData struct {

        
            GlobalValidation GlobalValidation  `json:"global_validation"`
            TemplateValidation map[string]interface{}  `json:"template_validation"`
         
    }
    
    // TemplatesValidationResponse used by Catalog
    type TemplatesValidationResponse struct {

        
            TemplateDetails TemplateDetails  `json:"template_details"`
            Data TemplateValidationData  `json:"data"`
         
    }
    
    // InventoryValidationResponse used by Catalog
    type InventoryValidationResponse struct {

        
            Message string  `json:"message"`
            Data map[string]interface{}  `json:"data"`
         
    }
    
    // HSNData used by Catalog
    type HSNData struct {

        
            HsnCode []string  `json:"hsn_code"`
            CountryOfOrigin []string  `json:"country_of_origin"`
         
    }
    
    // HSNCodesResponse used by Catalog
    type HSNCodesResponse struct {

        
            Message string  `json:"message"`
            Data HSNData  `json:"data"`
         
    }
    
    // UserInfo1 used by Catalog
    type UserInfo1 struct {

        
            Username string  `json:"username"`
            Email string  `json:"email"`
            UserID string  `json:"user_id"`
            UID string  `json:"uid"`
         
    }
    
    // ProductTemplateExportResponse used by Catalog
    type ProductTemplateExportResponse struct {

        
            ModifiedOn string  `json:"modified_on"`
            NotificationEmails []string  `json:"notification_emails"`
            Status string  `json:"status"`
            SellerID float64  `json:"seller_id"`
            Filters map[string]interface{}  `json:"filters"`
            CompletedOn string  `json:"completed_on"`
            Type string  `json:"type"`
            CreatedBy UserInfo1  `json:"created_by"`
            TaskID string  `json:"task_id"`
            URL string  `json:"url"`
         
    }
    
    // ProductDownloadsResponse used by Catalog
    type ProductDownloadsResponse struct {

        
            Items []ProductTemplateExportResponse  `json:"items"`
         
    }
    
    // ProductTemplateExportFilterRequest used by Catalog
    type ProductTemplateExportFilterRequest struct {

        
            Templates []string  `json:"templates"`
            Brands []string  `json:"brands"`
            FromDate string  `json:"from_date"`
            ToDate string  `json:"to_date"`
            CatalogueTypes []string  `json:"catalogue_types"`
         
    }
    
    // ProductTemplateDownloadsExport used by Catalog
    type ProductTemplateDownloadsExport struct {

        
            Type string  `json:"type"`
            NotificationEmails []string  `json:"notification_emails"`
            Filters ProductTemplateExportFilterRequest  `json:"filters"`
         
    }
    
    // ProductPartialExportRequest used by Catalog
    type ProductPartialExportRequest struct {

        
            Status string  `json:"status"`
            NotificationEmails []string  `json:"notification_emails"`
         
    }
    
    // ProductConfigurationDownloads used by Catalog
    type ProductConfigurationDownloads struct {

        
            Multivalue bool  `json:"multivalue"`
            Data []map[string]interface{}  `json:"data"`
         
    }
    
    // Media2 used by Catalog
    type Media2 struct {

        
            Logo string  `json:"logo"`
            Landscape string  `json:"landscape"`
            Portrait string  `json:"portrait"`
         
    }
    
    // Hierarchy used by Catalog
    type Hierarchy struct {

        
            L2 float64  `json:"l2"`
            Department float64  `json:"department"`
            L1 float64  `json:"l1"`
         
    }
    
    // CategoryMappingValues used by Catalog
    type CategoryMappingValues struct {

        
            CatalogID float64  `json:"catalog_id"`
            Name string  `json:"name"`
         
    }
    
    // CategoryMapping used by Catalog
    type CategoryMapping struct {

        
            Ajio CategoryMappingValues  `json:"ajio"`
            Facebook CategoryMappingValues  `json:"facebook"`
            Google CategoryMappingValues  `json:"google"`
         
    }
    
    // Category used by Catalog
    type Category struct {

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            Media Media2  `json:"media"`
            ModifiedOn string  `json:"modified_on"`
            ID string  `json:"id"`
            Hierarchy []Hierarchy  `json:"hierarchy"`
            CreatedOn string  `json:"created_on"`
            Slug string  `json:"slug"`
            Level float64  `json:"level"`
            Departments []float64  `json:"departments"`
            Synonyms []string  `json:"synonyms"`
            IsActive bool  `json:"is_active"`
            Tryouts []string  `json:"tryouts"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Marketplaces CategoryMapping  `json:"marketplaces"`
            Priority float64  `json:"priority"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
         
    }
    
    // CategoryResponse used by Catalog
    type CategoryResponse struct {

        
            Page Page  `json:"page"`
            Items []Category  `json:"items"`
         
    }
    
    // CategoryRequestBody used by Catalog
    type CategoryRequestBody struct {

        
            Name string  `json:"name"`
            Media Media2  `json:"media"`
            Hierarchy []Hierarchy  `json:"hierarchy"`
            Slug string  `json:"slug"`
            Level float64  `json:"level"`
            Departments []float64  `json:"departments"`
            Synonyms []string  `json:"synonyms"`
            IsActive bool  `json:"is_active"`
            Tryouts []string  `json:"tryouts"`
            Marketplaces CategoryMapping  `json:"marketplaces"`
            Priority float64  `json:"priority"`
         
    }
    
    // CategoryCreateResponse used by Catalog
    type CategoryCreateResponse struct {

        
            Message string  `json:"message"`
            UID float64  `json:"uid"`
         
    }
    
    // SingleCategoryResponse used by Catalog
    type SingleCategoryResponse struct {

        
            Data Category  `json:"data"`
         
    }
    
    // CategoryUpdateResponse used by Catalog
    type CategoryUpdateResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // Trader used by Catalog
    type Trader struct {

        
            Address []string  `json:"address"`
            Name string  `json:"name"`
            Type string  `json:"type"`
         
    }
    
    // VerifiedBy used by Catalog
    type VerifiedBy struct {

        
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
         
    }
    
    // Image used by Catalog
    type Image struct {

        
            AspectRatioF float64  `json:"aspect_ratio_f"`
            SecureURL string  `json:"secure_url"`
            AspectRatio string  `json:"aspect_ratio"`
            URL string  `json:"url"`
         
    }
    
    // ReturnConfigResponse used by Catalog
    type ReturnConfigResponse struct {

        
            Time float64  `json:"time"`
            Unit string  `json:"unit"`
            Returnable bool  `json:"returnable"`
         
    }
    
    // Logo used by Catalog
    type Logo struct {

        
            AspectRatioF float64  `json:"aspect_ratio_f"`
            SecureURL string  `json:"secure_url"`
            AspectRatio string  `json:"aspect_ratio"`
            URL string  `json:"url"`
         
    }
    
    // Brand used by Catalog
    type Brand struct {

        
            Logo Logo  `json:"logo"`
            Name string  `json:"name"`
            UID float64  `json:"uid"`
         
    }
    
    // NetQuantityResponse used by Catalog
    type NetQuantityResponse struct {

        
            Unit string  `json:"unit"`
            Value float64  `json:"value"`
         
    }
    
    // ProductPublish used by Catalog
    type ProductPublish struct {

        
            IsSet bool  `json:"is_set"`
            ProductOnlineDate string  `json:"product_online_date"`
         
    }
    
    // ProductSchemaV2 used by Catalog
    type ProductSchemaV2 struct {

        
            UID float64  `json:"uid"`
            Trader []Trader  `json:"trader"`
            IsImageLessProduct bool  `json:"is_image_less_product"`
            VerifiedBy VerifiedBy  `json:"verified_by"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            Images []Image  `json:"images"`
            PrimaryColor string  `json:"primary_color"`
            NoOfBoxes float64  `json:"no_of_boxes"`
            CountryOfOrigin string  `json:"country_of_origin"`
            Departments []float64  `json:"departments"`
            AllCompanyIds []float64  `json:"all_company_ids"`
            ImageNature string  `json:"image_nature"`
            AllSizes []map[string]interface{}  `json:"all_sizes"`
            Tags []string  `json:"tags"`
            Category map[string]interface{}  `json:"category"`
            Pending string  `json:"pending"`
            L3Mapping []string  `json:"l3_mapping"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            IsSet bool  `json:"is_set"`
            Highlights []string  `json:"highlights"`
            Media []Media1  `json:"media"`
            CreatedOn string  `json:"created_on"`
            VariantMedia map[string]interface{}  `json:"variant_media"`
            MultiSize bool  `json:"multi_size"`
            Attributes map[string]interface{}  `json:"attributes"`
            Moq map[string]interface{}  `json:"moq"`
            Color string  `json:"color"`
            ReturnConfig ReturnConfigResponse  `json:"return_config"`
            Brand Brand  `json:"brand"`
            CategorySlug string  `json:"category_slug"`
            ItemCode string  `json:"item_code"`
            Variants map[string]interface{}  `json:"variants"`
            ID string  `json:"id"`
            ModifiedOn string  `json:"modified_on"`
            TemplateTag string  `json:"template_tag"`
            NetQuantity NetQuantityResponse  `json:"net_quantity"`
            HsnCode string  `json:"hsn_code"`
            IsDependent bool  `json:"is_dependent"`
            VariantGroup map[string]interface{}  `json:"variant_group"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            CategoryUID float64  `json:"category_uid"`
            ProductGroupTag []string  `json:"product_group_tag"`
            Currency string  `json:"currency"`
            IsExpirable bool  `json:"is_expirable"`
            SizeGuide string  `json:"size_guide"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Description string  `json:"description"`
            CompanyID float64  `json:"company_id"`
            Stage string  `json:"stage"`
            ShortDescription string  `json:"short_description"`
            Sizes []map[string]interface{}  `json:"sizes"`
            Name string  `json:"name"`
            BrandUID float64  `json:"brand_uid"`
            Slug string  `json:"slug"`
            VerifiedOn string  `json:"verified_on"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            IsActive bool  `json:"is_active"`
            AllIdentifiers []string  `json:"all_identifiers"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            IsPhysical bool  `json:"is_physical"`
            ItemType string  `json:"item_type"`
            ProductPublish ProductPublish  `json:"product_publish"`
         
    }
    
    // ProductListingResponseV2 used by Catalog
    type ProductListingResponseV2 struct {

        
            Page Page  `json:"page"`
            Items []ProductSchemaV2  `json:"items"`
         
    }
    
    // CustomOrder used by Catalog
    type CustomOrder struct {

        
            IsCustomOrder bool  `json:"is_custom_order"`
            ManufacturingTime float64  `json:"manufacturing_time"`
            ManufacturingTimeUnit string  `json:"manufacturing_time_unit"`
         
    }
    
    // ReturnConfig used by Catalog
    type ReturnConfig struct {

        
            Time float64  `json:"time"`
            Unit string  `json:"unit"`
            Returnable bool  `json:"returnable"`
         
    }
    
    // NetQuantity used by Catalog
    type NetQuantity struct {

        
            Unit string  `json:"unit"`
            Value float64  `json:"value"`
         
    }
    
    // TaxIdentifier used by Catalog
    type TaxIdentifier struct {

        
            HsnCode string  `json:"hsn_code"`
            HsnCodeID string  `json:"hsn_code_id"`
            ReportingHsn string  `json:"reporting_hsn"`
         
    }
    
    // TeaserTag used by Catalog
    type TeaserTag struct {

        
            Tag string  `json:"tag"`
            URL string  `json:"url"`
         
    }
    
    // ProductPublish1 used by Catalog
    type ProductPublish1 struct {

        
            IsSet bool  `json:"is_set"`
            ProductOnlineDate string  `json:"product_online_date"`
         
    }
    
    // ProductCreateUpdateSchemaV2 used by Catalog
    type ProductCreateUpdateSchemaV2 struct {

        
            Trader []Trader  `json:"trader"`
            UID float64  `json:"uid"`
            IsImageLessProduct bool  `json:"is_image_less_product"`
            CustomOrder CustomOrder  `json:"custom_order"`
            NoOfBoxes float64  `json:"no_of_boxes"`
            CountryOfOrigin string  `json:"country_of_origin"`
            Departments []float64  `json:"departments"`
            Tags []string  `json:"tags"`
            BulkJobID string  `json:"bulk_job_id"`
            IsSet bool  `json:"is_set"`
            Highlights []string  `json:"highlights"`
            Media []Media1  `json:"media"`
            VariantMedia map[string]interface{}  `json:"variant_media"`
            MultiSize bool  `json:"multi_size"`
            Attributes map[string]interface{}  `json:"attributes"`
            ReturnConfig ReturnConfig  `json:"return_config"`
            CategorySlug string  `json:"category_slug"`
            ItemCode string  `json:"item_code"`
            Variants map[string]interface{}  `json:"variants"`
            Requester string  `json:"requester"`
            TemplateTag string  `json:"template_tag"`
            NetQuantity NetQuantity  `json:"net_quantity"`
            IsDependent bool  `json:"is_dependent"`
            VariantGroup map[string]interface{}  `json:"variant_group"`
            TaxIdentifier TaxIdentifier  `json:"tax_identifier"`
            ProductGroupTag []string  `json:"product_group_tag"`
            Currency string  `json:"currency"`
            SizeGuide string  `json:"size_guide"`
            Description string  `json:"description"`
            CompanyID float64  `json:"company_id"`
            ShortDescription string  `json:"short_description"`
            Sizes []map[string]interface{}  `json:"sizes"`
            Name string  `json:"name"`
            ChangeRequestID interface{}  `json:"change_request_id"`
            BrandUID float64  `json:"brand_uid"`
            Slug string  `json:"slug"`
            Action string  `json:"action"`
            TeaserTag TeaserTag  `json:"teaser_tag"`
            IsActive bool  `json:"is_active"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ItemType string  `json:"item_type"`
            ProductPublish ProductPublish1  `json:"product_publish"`
         
    }
    
    // ProductVariants used by Catalog
    type ProductVariants struct {

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            Media []Media1  `json:"media"`
            BrandUID float64  `json:"brand_uid"`
            CategoryUID float64  `json:"category_uid"`
            ItemCode string  `json:"item_code"`
         
    }
    
    // ProductVariantsResponse used by Catalog
    type ProductVariantsResponse struct {

        
            Variants []ProductVariants  `json:"variants"`
            Page Page  `json:"page"`
         
    }
    
    // AttributeMasterSerializer used by Catalog
    type AttributeMasterSerializer struct {

        
            Filters AttributeMasterFilter  `json:"filters"`
            Departments []string  `json:"departments"`
            Synonyms map[string]interface{}  `json:"synonyms"`
            RawKey string  `json:"raw_key"`
            Tags []string  `json:"tags"`
            Schema AttributeMaster  `json:"schema"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            EnabledForEndConsumer bool  `json:"enabled_for_end_consumer"`
            CreatedOn string  `json:"created_on"`
            IsNested bool  `json:"is_nested"`
            Unit string  `json:"unit"`
            ModifiedOn string  `json:"modified_on"`
            Details AttributeMasterDetails  `json:"details"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Logo string  `json:"logo"`
            Suggestion string  `json:"suggestion"`
            Description string  `json:"description"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Variant bool  `json:"variant"`
         
    }
    
    // ProductAttributesResponse used by Catalog
    type ProductAttributesResponse struct {

        
            Items []AttributeMasterSerializer  `json:"items"`
         
    }
    
    // SingleProductResponse used by Catalog
    type SingleProductResponse struct {

        
            Data ProductSchemaV2  `json:"data"`
         
    }
    
    // ValidateIdentifier used by Catalog
    type ValidateIdentifier struct {

        
            GtinValue string  `json:"gtin_value"`
            Primary bool  `json:"primary"`
            GtinType string  `json:"gtin_type"`
         
    }
    
    // AllSizes used by Catalog
    type AllSizes struct {

        
            Identifiers []ValidateIdentifier  `json:"identifiers"`
            ItemHeight float64  `json:"item_height"`
            ItemDimensionsUnitOfMeasure string  `json:"item_dimensions_unit_of_measure"`
            ItemWeightUnitOfMeasure string  `json:"item_weight_unit_of_measure"`
            Size string  `json:"size"`
            ItemLength float64  `json:"item_length"`
            ItemWidth float64  `json:"item_width"`
            ItemWeight float64  `json:"item_weight"`
         
    }
    
    // GetAllSizes used by Catalog
    type GetAllSizes struct {

        
            AllSizes []AllSizes  `json:"all_sizes"`
         
    }
    
    // ValidateProduct used by Catalog
    type ValidateProduct struct {

        
            Valid bool  `json:"valid"`
         
    }
    
    // ProductPublished used by Catalog
    type ProductPublished struct {

        
            IsSet bool  `json:"is_set"`
            ProductOnlineDate float64  `json:"product_online_date"`
         
    }
    
    // Product used by Catalog
    type Product struct {

        
            UID float64  `json:"uid"`
            Trader []Trader  `json:"trader"`
            IsImageLessProduct bool  `json:"is_image_less_product"`
            VerifiedBy VerifiedBy  `json:"verified_by"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            Images []Image  `json:"images"`
            PrimaryColor string  `json:"primary_color"`
            NoOfBoxes float64  `json:"no_of_boxes"`
            CountryOfOrigin string  `json:"country_of_origin"`
            Departments []float64  `json:"departments"`
            AllCompanyIds []float64  `json:"all_company_ids"`
            ImageNature string  `json:"image_nature"`
            AllSizes []map[string]interface{}  `json:"all_sizes"`
            Tags []string  `json:"tags"`
            Category map[string]interface{}  `json:"category"`
            Pending string  `json:"pending"`
            L3Mapping []string  `json:"l3_mapping"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            IsSet bool  `json:"is_set"`
            Highlights []string  `json:"highlights"`
            Media []Media1  `json:"media"`
            CreatedOn string  `json:"created_on"`
            VariantMedia map[string]interface{}  `json:"variant_media"`
            MultiSize bool  `json:"multi_size"`
            Attributes map[string]interface{}  `json:"attributes"`
            Moq map[string]interface{}  `json:"moq"`
            Color string  `json:"color"`
            ReturnConfig ReturnConfigResponse  `json:"return_config"`
            Brand Brand  `json:"brand"`
            CategorySlug string  `json:"category_slug"`
            ItemCode string  `json:"item_code"`
            Variants map[string]interface{}  `json:"variants"`
            ID string  `json:"id"`
            ModifiedOn string  `json:"modified_on"`
            TemplateTag string  `json:"template_tag"`
            NetQuantity NetQuantityResponse  `json:"net_quantity"`
            HsnCode string  `json:"hsn_code"`
            IsDependent bool  `json:"is_dependent"`
            VariantGroup map[string]interface{}  `json:"variant_group"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            CategoryUID float64  `json:"category_uid"`
            ProductGroupTag []string  `json:"product_group_tag"`
            Currency string  `json:"currency"`
            IsExpirable bool  `json:"is_expirable"`
            SizeGuide string  `json:"size_guide"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Description string  `json:"description"`
            CompanyID float64  `json:"company_id"`
            Stage string  `json:"stage"`
            ShortDescription string  `json:"short_description"`
            Sizes []map[string]interface{}  `json:"sizes"`
            Name string  `json:"name"`
            BrandUID float64  `json:"brand_uid"`
            Slug string  `json:"slug"`
            VerifiedOn string  `json:"verified_on"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            IsActive bool  `json:"is_active"`
            AllIdentifiers []string  `json:"all_identifiers"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            IsPhysical bool  `json:"is_physical"`
            ItemType string  `json:"item_type"`
            ProductPublish ProductPublished  `json:"product_publish"`
         
    }
    
    // ProductListingResponse used by Catalog
    type ProductListingResponse struct {

        
            Page Page  `json:"page"`
            Items []Product  `json:"items"`
         
    }
    
    // UserDetail1 used by Catalog
    type UserDetail1 struct {

        
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
            FullName string  `json:"full_name"`
         
    }
    
    // ProductBulkRequest used by Catalog
    type ProductBulkRequest struct {

        
            Stage string  `json:"stage"`
            CancelledRecords []string  `json:"cancelled_records"`
            ModifiedOn string  `json:"modified_on"`
            TemplateTag string  `json:"template_tag"`
            CreatedOn string  `json:"created_on"`
            Total float64  `json:"total"`
            FilePath string  `json:"file_path"`
            Failed float64  `json:"failed"`
            IsActive bool  `json:"is_active"`
            Cancelled float64  `json:"cancelled"`
            Template ProductTemplate  `json:"template"`
            CreatedBy UserDetail1  `json:"created_by"`
            Succeed float64  `json:"succeed"`
            CompanyID float64  `json:"company_id"`
            FailedRecords []string  `json:"failed_records"`
            ModifiedBy UserDetail1  `json:"modified_by"`
         
    }
    
    // ProductBulkRequestList used by Catalog
    type ProductBulkRequestList struct {

        
            Page Page  `json:"page"`
            Items ProductBulkRequest  `json:"items"`
         
    }
    
    // BulkJob used by Catalog
    type BulkJob struct {

        
            Stage string  `json:"stage"`
            CancelledRecords []map[string]interface{}  `json:"cancelled_records"`
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
            TemplateTag string  `json:"template_tag"`
            Total float64  `json:"total"`
            CustomTemplateTag string  `json:"custom_template_tag"`
            FilePath string  `json:"file_path"`
            Cancelled float64  `json:"cancelled"`
            IsActive bool  `json:"is_active"`
            Failed float64  `json:"failed"`
            TrackingURL string  `json:"tracking_url"`
            CreatedBy UserInfo1  `json:"created_by"`
            Succeed float64  `json:"succeed"`
            CompanyID float64  `json:"company_id"`
            FailedRecords []map[string]interface{}  `json:"failed_records"`
            ModifiedBy UserInfo1  `json:"modified_by"`
         
    }
    
    // BulkResponse used by Catalog
    type BulkResponse struct {

        
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
            IsActive bool  `json:"is_active"`
            BatchID string  `json:"batch_id"`
            CreatedBy UserInfo1  `json:"created_by"`
            ModifiedBy UserInfo1  `json:"modified_by"`
         
    }
    
    // BulkProductRequest used by Catalog
    type BulkProductRequest struct {

        
            TemplateTag string  `json:"template_tag"`
            BatchID string  `json:"batch_id"`
            CompanyID float64  `json:"company_id"`
            Data []map[string]interface{}  `json:"data"`
         
    }
    
    // ProductTagsViewResponse used by Catalog
    type ProductTagsViewResponse struct {

        
            Items []string  `json:"items"`
         
    }
    
    // UserCommon used by Catalog
    type UserCommon struct {

        
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
            CompanyID float64  `json:"company_id"`
         
    }
    
    // Items used by Catalog
    type Items struct {

        
            Stage string  `json:"stage"`
            CancelledRecords []string  `json:"cancelled_records"`
            ID string  `json:"id"`
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
            Total float64  `json:"total"`
            FilePath string  `json:"file_path"`
            Failed float64  `json:"failed"`
            IsActive bool  `json:"is_active"`
            Retry float64  `json:"retry"`
            Cancelled float64  `json:"cancelled"`
            TrackingURL string  `json:"tracking_url"`
            CreatedBy UserCommon  `json:"created_by"`
            Succeed float64  `json:"succeed"`
            CompanyID float64  `json:"company_id"`
            FailedRecords []string  `json:"failed_records"`
            ModifiedBy UserCommon  `json:"modified_by"`
         
    }
    
    // BulkAssetResponse used by Catalog
    type BulkAssetResponse struct {

        
            Page Page  `json:"page"`
            Items []Items  `json:"items"`
         
    }
    
    // ProductBulkAssets used by Catalog
    type ProductBulkAssets struct {

        
            User map[string]interface{}  `json:"user"`
            CompanyID float64  `json:"company_id"`
            URL string  `json:"url"`
         
    }
    
    // ProductSizeDeleteDataResponse used by Catalog
    type ProductSizeDeleteDataResponse struct {

        
            ItemID float64  `json:"item_id"`
            CompanyID float64  `json:"company_id"`
            Size string  `json:"size"`
         
    }
    
    // ProductSizeDeleteResponse used by Catalog
    type ProductSizeDeleteResponse struct {

        
            Success bool  `json:"success"`
            Data ProductSizeDeleteDataResponse  `json:"data"`
         
    }
    
    // InventoryResponse used by Catalog
    type InventoryResponse struct {

        
            Identifiers map[string]interface{}  `json:"identifiers"`
            UID string  `json:"uid"`
            Store map[string]interface{}  `json:"store"`
            SellableQuantity float64  `json:"sellable_quantity"`
            PriceEffective float64  `json:"price_effective"`
            PriceTransfer float64  `json:"price_transfer"`
            Size string  `json:"size"`
            InventoryUpdatedOn string  `json:"inventory_updated_on"`
            ItemID float64  `json:"item_id"`
            SellerIdentifier string  `json:"seller_identifier"`
            Currency string  `json:"currency"`
            Quantity float64  `json:"quantity"`
            Price float64  `json:"price"`
         
    }
    
    // InventoryResponsePaginated used by Catalog
    type InventoryResponsePaginated struct {

        
            Page Page  `json:"page"`
            Items []InventoryResponse  `json:"items"`
         
    }
    
    // ItemQuery used by Catalog
    type ItemQuery struct {

        
            UID float64  `json:"uid"`
            ItemCode string  `json:"item_code"`
            BrandUID float64  `json:"brand_uid"`
         
    }
    
    // GTIN used by Catalog
    type GTIN struct {

        
            GtinValue string  `json:"gtin_value"`
            Primary bool  `json:"primary"`
            GtinType string  `json:"gtin_type"`
         
    }
    
    // SetSize used by Catalog
    type SetSize struct {

        
            Pieces float64  `json:"pieces"`
            Size string  `json:"size"`
         
    }
    
    // SizeDistribution used by Catalog
    type SizeDistribution struct {

        
            Sizes []SetSize  `json:"sizes"`
         
    }
    
    // InventorySet used by Catalog
    type InventorySet struct {

        
            Quantity float64  `json:"quantity"`
            Name string  `json:"name"`
            SizeDistribution SizeDistribution  `json:"size_distribution"`
         
    }
    
    // InvSize used by Catalog
    type InvSize struct {

        
            ItemHeight float64  `json:"item_height"`
            IsSet bool  `json:"is_set"`
            Identifiers []GTIN  `json:"identifiers"`
            ItemDimensionsUnitOfMeasure string  `json:"item_dimensions_unit_of_measure"`
            ExpirationDate string  `json:"expiration_date"`
            PriceEffective float64  `json:"price_effective"`
            ItemWeightUnitOfMeasure string  `json:"item_weight_unit_of_measure"`
            Set InventorySet  `json:"set"`
            PriceTransfer float64  `json:"price_transfer"`
            Size string  `json:"size"`
            StoreCode string  `json:"store_code"`
            ItemLength float64  `json:"item_length"`
            ItemWidth float64  `json:"item_width"`
            ItemWeight float64  `json:"item_weight"`
            Currency string  `json:"currency"`
            Quantity float64  `json:"quantity"`
            Price float64  `json:"price"`
         
    }
    
    // InventoryRequest used by Catalog
    type InventoryRequest struct {

        
            Item ItemQuery  `json:"item"`
            Sizes []InvSize  `json:"sizes"`
            CompanyID float64  `json:"company_id"`
         
    }
    
    // Trader1 used by Catalog
    type Trader1 struct {

        
            Address []string  `json:"address"`
            Name string  `json:"name"`
            Type string  `json:"type"`
         
    }
    
    // WeightResponse used by Catalog
    type WeightResponse struct {

        
            Shipping float64  `json:"shipping"`
            Unit string  `json:"unit"`
            IsDefault bool  `json:"is_default"`
         
    }
    
    // CompanyMeta used by Catalog
    type CompanyMeta struct {

        
            ID float64  `json:"id"`
         
    }
    
    // ReturnConfig1 used by Catalog
    type ReturnConfig1 struct {

        
            Time float64  `json:"time"`
            Unit string  `json:"unit"`
            Returnable bool  `json:"returnable"`
         
    }
    
    // BrandMeta used by Catalog
    type BrandMeta struct {

        
            ID float64  `json:"id"`
            Name string  `json:"name"`
         
    }
    
    // PriceMeta used by Catalog
    type PriceMeta struct {

        
            Marked float64  `json:"marked"`
            UpdatedAt string  `json:"updated_at"`
            Transfer float64  `json:"transfer"`
            TpNotes map[string]interface{}  `json:"tp_notes"`
            Currency string  `json:"currency"`
            Effective float64  `json:"effective"`
         
    }
    
    // StoreMeta used by Catalog
    type StoreMeta struct {

        
            ID float64  `json:"id"`
         
    }
    
    // ManufacturerResponse used by Catalog
    type ManufacturerResponse struct {

        
            Address string  `json:"address"`
            Name string  `json:"name"`
            IsDefault bool  `json:"is_default"`
         
    }
    
    // DimensionResponse used by Catalog
    type DimensionResponse struct {

        
            Height float64  `json:"height"`
            Unit string  `json:"unit"`
            Width float64  `json:"width"`
            IsDefault bool  `json:"is_default"`
            Length float64  `json:"length"`
         
    }
    
    // QuantityBase used by Catalog
    type QuantityBase struct {

        
            UpdatedAt string  `json:"updated_at"`
            Count float64  `json:"count"`
         
    }
    
    // Quantities used by Catalog
    type Quantities struct {

        
            OrderCommitted QuantityBase  `json:"order_committed"`
            NotAvailable QuantityBase  `json:"not_available"`
            Sellable QuantityBase  `json:"sellable"`
            Damaged QuantityBase  `json:"damaged"`
         
    }
    
    // InventorySellerResponse used by Catalog
    type InventorySellerResponse struct {

        
            UID string  `json:"uid"`
            Trader []Trader1  `json:"trader"`
            CountryOfOrigin string  `json:"country_of_origin"`
            Tags []string  `json:"tags"`
            Identifier map[string]interface{}  `json:"identifier"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            FyndItemCode string  `json:"fynd_item_code"`
            IsSet bool  `json:"is_set"`
            Weight WeightResponse  `json:"weight"`
            Set InventorySet  `json:"set"`
            Company CompanyMeta  `json:"company"`
            Size string  `json:"size"`
            TrackInventory bool  `json:"track_inventory"`
            RawMeta map[string]interface{}  `json:"raw_meta"`
            ReturnConfig ReturnConfig1  `json:"return_config"`
            TraceID string  `json:"trace_id"`
            Brand BrandMeta  `json:"brand"`
            Price PriceMeta  `json:"price"`
            FyndArticleCode string  `json:"fynd_article_code"`
            Store StoreMeta  `json:"store"`
            FyndMeta map[string]interface{}  `json:"fynd_meta"`
            Fragile bool  `json:"fragile"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            AddedOnStore string  `json:"added_on_store"`
            ItemID float64  `json:"item_id"`
            Manufacturer ManufacturerResponse  `json:"manufacturer"`
            CreatedBy UserSerializer  `json:"created_by"`
            Stage string  `json:"stage"`
            TotalQuantity float64  `json:"total_quantity"`
            Meta map[string]interface{}  `json:"meta"`
            ExpirationDate string  `json:"expiration_date"`
            SellerIdentifier string  `json:"seller_identifier"`
            IsActive bool  `json:"is_active"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Dimension DimensionResponse  `json:"dimension"`
            Quantities Quantities  `json:"quantities"`
         
    }
    
    // InventorySellerIdentifierResponsePaginated used by Catalog
    type InventorySellerIdentifierResponsePaginated struct {

        
            Page Page  `json:"page"`
            Items []InventorySellerResponse  `json:"items"`
         
    }
    
    // Trader2 used by Catalog
    type Trader2 struct {

        
            Address []string  `json:"address"`
            Name string  `json:"name"`
            Type string  `json:"type"`
         
    }
    
    // WeightResponse1 used by Catalog
    type WeightResponse1 struct {

        
            Shipping float64  `json:"shipping"`
            Unit string  `json:"unit"`
         
    }
    
    // CompanyMeta1 used by Catalog
    type CompanyMeta1 struct {

        
            ID float64  `json:"id"`
         
    }
    
    // ReturnConfig2 used by Catalog
    type ReturnConfig2 struct {

        
            Time float64  `json:"time"`
            Unit string  `json:"unit"`
            Returnable bool  `json:"returnable"`
         
    }
    
    // BrandMeta1 used by Catalog
    type BrandMeta1 struct {

        
            ID float64  `json:"id"`
            Name string  `json:"name"`
         
    }
    
    // PriceArticle used by Catalog
    type PriceArticle struct {

        
            Marked float64  `json:"marked"`
            Transfer float64  `json:"transfer"`
            TpNotes map[string]interface{}  `json:"tp_notes"`
            Currency string  `json:"currency"`
            Effective float64  `json:"effective"`
         
    }
    
    // ArticleStoreResponse used by Catalog
    type ArticleStoreResponse struct {

        
            UID float64  `json:"uid"`
            StoreType string  `json:"store_type"`
            Name string  `json:"name"`
            StoreCode string  `json:"store_code"`
         
    }
    
    // DateMeta used by Catalog
    type DateMeta struct {

        
            CreatedOn string  `json:"created_on"`
            InventoryUpdatedOn string  `json:"inventory_updated_on"`
            AddedOnStore string  `json:"added_on_store"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // ManufacturerResponse1 used by Catalog
    type ManufacturerResponse1 struct {

        
            Address string  `json:"address"`
            Name string  `json:"name"`
            IsDefault bool  `json:"is_default"`
         
    }
    
    // DimensionResponse1 used by Catalog
    type DimensionResponse1 struct {

        
            Length float64  `json:"length"`
            Height float64  `json:"height"`
            Unit string  `json:"unit"`
            Width float64  `json:"width"`
         
    }
    
    // Quantity used by Catalog
    type Quantity struct {

        
            Count float64  `json:"count"`
         
    }
    
    // QuantitiesArticle used by Catalog
    type QuantitiesArticle struct {

        
            OrderCommitted Quantity  `json:"order_committed"`
            NotAvailable Quantity  `json:"not_available"`
            Sellable Quantity  `json:"sellable"`
            Damaged Quantity  `json:"damaged"`
         
    }
    
    // GetInventories used by Catalog
    type GetInventories struct {

        
            Trader []Trader2  `json:"trader"`
            UID string  `json:"uid"`
            CountryOfOrigin string  `json:"country_of_origin"`
            Tags []string  `json:"tags"`
            Identifier map[string]interface{}  `json:"identifier"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            IsSet bool  `json:"is_set"`
            Weight WeightResponse1  `json:"weight"`
            Company CompanyMeta1  `json:"company"`
            Size string  `json:"size"`
            TrackInventory bool  `json:"track_inventory"`
            ReturnConfig ReturnConfig2  `json:"return_config"`
            TraceID string  `json:"trace_id"`
            Brand BrandMeta1  `json:"brand"`
            Price PriceArticle  `json:"price"`
            Store ArticleStoreResponse  `json:"store"`
            ID string  `json:"id"`
            DateMeta DateMeta  `json:"date_meta"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            InventoryUpdatedOn string  `json:"inventory_updated_on"`
            ItemID float64  `json:"item_id"`
            Manufacturer ManufacturerResponse1  `json:"manufacturer"`
            CreatedBy UserSerializer  `json:"created_by"`
            Stage string  `json:"stage"`
            TotalQuantity float64  `json:"total_quantity"`
            ExpirationDate string  `json:"expiration_date"`
            SellerIdentifier string  `json:"seller_identifier"`
            Dimension DimensionResponse1  `json:"dimension"`
            Platforms map[string]interface{}  `json:"platforms"`
            Quantities QuantitiesArticle  `json:"quantities"`
         
    }
    
    // GetInventoriesResponse used by Catalog
    type GetInventoriesResponse struct {

        
            Page Page  `json:"page"`
            Items []GetInventories  `json:"items"`
         
    }
    
    // BulkInventoryGetItems used by Catalog
    type BulkInventoryGetItems struct {

        
            Stage string  `json:"stage"`
            CancelledRecords []string  `json:"cancelled_records"`
            ID string  `json:"id"`
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
            Total float64  `json:"total"`
            FilePath string  `json:"file_path"`
            Cancelled float64  `json:"cancelled"`
            Failed float64  `json:"failed"`
            IsActive bool  `json:"is_active"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Succeed float64  `json:"succeed"`
            CompanyID float64  `json:"company_id"`
            FailedRecords []string  `json:"failed_records"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
         
    }
    
    // BulkInventoryGet used by Catalog
    type BulkInventoryGet struct {

        
            Page Page  `json:"page"`
            Items []BulkInventoryGetItems  `json:"items"`
         
    }
    
    // InventoryJobPayload used by Catalog
    type InventoryJobPayload struct {

        
            ItemDimensionsUnitOfMeasure string  `json:"item_dimensions_unit_of_measure"`
            TotalQuantity float64  `json:"total_quantity"`
            ExpirationDate string  `json:"expiration_date"`
            PriceEffective float64  `json:"price_effective"`
            ItemWeightUnitOfMeasure string  `json:"item_weight_unit_of_measure"`
            StoreCode string  `json:"store_code"`
            SellerIdentifier string  `json:"seller_identifier"`
            PriceMarked float64  `json:"price_marked"`
            Tags []string  `json:"tags"`
            Currency string  `json:"currency"`
            TraceID string  `json:"trace_id"`
            Quantity float64  `json:"quantity"`
            Price float64  `json:"price"`
         
    }
    
    // InventoryBulkRequest used by Catalog
    type InventoryBulkRequest struct {

        
            User map[string]interface{}  `json:"user"`
            Sizes []InventoryJobPayload  `json:"sizes"`
            CompanyID float64  `json:"company_id"`
            BatchID string  `json:"batch_id"`
         
    }
    
    // InventoryExportQuantityFilter used by Catalog
    type InventoryExportQuantityFilter struct {

        
            Operators string  `json:"operators"`
            Min float64  `json:"min"`
            Max float64  `json:"max"`
         
    }
    
    // InventoryExportAdvanceOption used by Catalog
    type InventoryExportAdvanceOption struct {

        
            ToDate string  `json:"to_date"`
            StoreIds []float64  `json:"store_ids"`
            FromDate string  `json:"from_date"`
            BrandIds []float64  `json:"brand_ids"`
            Quantity InventoryExportQuantityFilter  `json:"quantity"`
         
    }
    
    // InventoryExportJob used by Catalog
    type InventoryExportJob struct {

        
            NotificationEmails []string  `json:"notification_emails"`
            Status string  `json:"status"`
            SellerID float64  `json:"seller_id"`
            Filters InventoryExportAdvanceOption  `json:"filters"`
            CompletedOn string  `json:"completed_on"`
            Type string  `json:"type"`
            TaskID string  `json:"task_id"`
            URL string  `json:"url"`
         
    }
    
    // InventoryExportRequest used by Catalog
    type InventoryExportRequest struct {

        
            Store []float64  `json:"store"`
            Type string  `json:"type"`
            Brand []float64  `json:"brand"`
         
    }
    
    // InventoryExportResponse used by Catalog
    type InventoryExportResponse struct {

        
            ModifiedOn string  `json:"modified_on"`
            NotificationEmails []string  `json:"notification_emails"`
            CreatedOn string  `json:"created_on"`
            Status string  `json:"status"`
            SellerID float64  `json:"seller_id"`
            Filters map[string]interface{}  `json:"filters"`
            Type string  `json:"type"`
            CreatedBy string  `json:"created_by"`
            TaskID string  `json:"task_id"`
         
    }
    
    // InventoryJobFilters used by Catalog
    type InventoryJobFilters struct {

        
            Brands []string  `json:"brands"`
            ToDate string  `json:"to_date"`
            FromDate string  `json:"from_date"`
            Stores []string  `json:"stores"`
            Quantity InventoryExportQuantityFilter  `json:"quantity"`
         
    }
    
    // InventoryJobDetailResponse used by Catalog
    type InventoryJobDetailResponse struct {

        
            ID string  `json:"id"`
            ModifiedOn string  `json:"modified_on"`
            NotificationEmails []string  `json:"notification_emails"`
            CreatedOn string  `json:"created_on"`
            Status string  `json:"status"`
            SellerID float64  `json:"seller_id"`
            Filters InventoryJobFilters  `json:"filters"`
            CompletedOn string  `json:"completed_on"`
            Type string  `json:"type"`
            CancelledOn string  `json:"cancelled_on"`
            CreatedBy UserDetail  `json:"created_by"`
            CancelledBy UserDetail  `json:"cancelled_by"`
            TaskID string  `json:"task_id"`
            URL string  `json:"url"`
         
    }
    
    // InventoryExportJobListResponse used by Catalog
    type InventoryExportJobListResponse struct {

        
            Items InventoryJobDetailResponse  `json:"items"`
         
    }
    
    // InventoryExportFilter used by Catalog
    type InventoryExportFilter struct {

        
            ToDate string  `json:"to_date"`
            StoreIds []float64  `json:"store_ids"`
            FromDate string  `json:"from_date"`
            BrandIds []float64  `json:"brand_ids"`
            Quantity InventoryExportQuantityFilter  `json:"quantity"`
         
    }
    
    // InventoryCreateRequest used by Catalog
    type InventoryCreateRequest struct {

        
            Data []string  `json:"data"`
            Type string  `json:"type"`
            NotificationEmails []string  `json:"notification_emails"`
            Filters InventoryExportFilter  `json:"filters"`
         
    }
    
    // InventoryPartialExportRequest used by Catalog
    type InventoryPartialExportRequest struct {

        
            Status string  `json:"status"`
            NotificationEmails []string  `json:"notification_emails"`
         
    }
    
    // InventoryExportCreateFilters used by Catalog
    type InventoryExportCreateFilters struct {

        
            ToDate string  `json:"to_date"`
            StoreIds []float64  `json:"store_ids"`
            FromDate string  `json:"from_date"`
            BrandIds []float64  `json:"brand_ids"`
            Quantity InventoryExportQuantityFilter  `json:"quantity"`
         
    }
    
    // InventoryExportCreateResponse used by Catalog
    type InventoryExportCreateResponse struct {

        
            ID string  `json:"id"`
            ModifiedOn string  `json:"modified_on"`
            NotificationEmails []string  `json:"notification_emails"`
            CreatedOn string  `json:"created_on"`
            Status string  `json:"status"`
            SellerID float64  `json:"seller_id"`
            Filters InventoryExportCreateFilters  `json:"filters"`
            CompletedOn string  `json:"completed_on"`
            Type string  `json:"type"`
            CancelledOn string  `json:"cancelled_on"`
            CreatedBy UserDetail  `json:"created_by"`
            CancelledBy UserDetail  `json:"cancelled_by"`
            TaskID string  `json:"task_id"`
            URL string  `json:"url"`
         
    }
    
    // FilerList used by Catalog
    type FilerList struct {

        
            Value string  `json:"value"`
            Display string  `json:"display"`
         
    }
    
    // InventoryConfig used by Catalog
    type InventoryConfig struct {

        
            Multivalues bool  `json:"multivalues"`
            Data []FilerList  `json:"data"`
         
    }
    
    // InventoryPayload used by Catalog
    type InventoryPayload struct {

        
            TotalQuantity float64  `json:"total_quantity"`
            ExpirationDate string  `json:"expiration_date"`
            PriceEffective float64  `json:"price_effective"`
            StoreID float64  `json:"store_id"`
            SellerIdentifier string  `json:"seller_identifier"`
            PriceMarked float64  `json:"price_marked"`
            Tags []string  `json:"tags"`
            TraceID string  `json:"trace_id"`
         
    }
    
    // InventoryRequestSchemaV2 used by Catalog
    type InventoryRequestSchemaV2 struct {

        
            Meta map[string]interface{}  `json:"meta"`
            Payload []InventoryPayload  `json:"payload"`
            CompanyID float64  `json:"company_id"`
         
    }
    
    // InventoryFailedReason used by Catalog
    type InventoryFailedReason struct {

        
            Message string  `json:"message"`
            Errors string  `json:"errors"`
         
    }
    
    // InventoryResponseItem used by Catalog
    type InventoryResponseItem struct {

        
            Data InventoryPayload  `json:"data"`
            Reason InventoryFailedReason  `json:"reason"`
         
    }
    
    // InventoryUpdateResponse used by Catalog
    type InventoryUpdateResponse struct {

        
            Message string  `json:"message"`
            Items []InventoryResponseItem  `json:"items"`
         
    }
    
    // HsnCodesObject used by Catalog
    type HsnCodesObject struct {

        
            TaxOnMrp bool  `json:"tax_on_mrp"`
            ID string  `json:"id"`
            Threshold1 float64  `json:"threshold1"`
            ModifiedOn string  `json:"modified_on"`
            Threshold2 float64  `json:"threshold2"`
            HsnCode string  `json:"hsn_code"`
            Tax1 float64  `json:"tax1"`
            Tax2 float64  `json:"tax2"`
            Hs2Code string  `json:"hs2_code"`
            CompanyID float64  `json:"company_id"`
            TaxOnEsp bool  `json:"tax_on_esp"`
         
    }
    
    // HsnCode used by Catalog
    type HsnCode struct {

        
            Data HsnCodesObject  `json:"data"`
         
    }
    
    // HsnUpsert used by Catalog
    type HsnUpsert struct {

        
            TaxOnMrp bool  `json:"tax_on_mrp"`
            UID float64  `json:"uid"`
            Threshold1 float64  `json:"threshold1"`
            Threshold2 float64  `json:"threshold2"`
            HsnCode string  `json:"hsn_code"`
            Tax1 float64  `json:"tax1"`
            Tax2 float64  `json:"tax2"`
            IsActive bool  `json:"is_active"`
            Hs2Code string  `json:"hs2_code"`
            CompanyID float64  `json:"company_id"`
            TaxOnEsp bool  `json:"tax_on_esp"`
         
    }
    
    // BulkHsnUpsert used by Catalog
    type BulkHsnUpsert struct {

        
            Data []HsnUpsert  `json:"data"`
         
    }
    
    // BulkHsnResponse used by Catalog
    type BulkHsnResponse struct {

        
            Success bool  `json:"success"`
         
    }
    
    // PageResponse used by Catalog
    type PageResponse struct {

        
            HasNext bool  `json:"has_next"`
            Size float64  `json:"size"`
            Current string  `json:"current"`
            ItemTotal float64  `json:"item_total"`
            HasPrevious bool  `json:"has_previous"`
         
    }
    
    // TaxSlab used by Catalog
    type TaxSlab struct {

        
            Threshold float64  `json:"threshold"`
            Rate float64  `json:"rate"`
            Cess float64  `json:"cess"`
            EffectiveDate string  `json:"effective_date"`
         
    }
    
    // HSNDataInsertV2 used by Catalog
    type HSNDataInsertV2 struct {

        
            CountryCode string  `json:"country_code"`
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
            HsnCode string  `json:"hsn_code"`
            Taxes []TaxSlab  `json:"taxes"`
            Type string  `json:"type"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Description string  `json:"description"`
            HsnCodeID string  `json:"hsn_code_id"`
            ReportingHsn string  `json:"reporting_hsn"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
         
    }
    
    // HsnCodesListingResponseSchemaV2 used by Catalog
    type HsnCodesListingResponseSchemaV2 struct {

        
            Page PageResponse  `json:"page"`
            Items []HSNDataInsertV2  `json:"items"`
         
    }
    
    // BrandItem used by Catalog
    type BrandItem struct {

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Discount string  `json:"discount"`
            Action Action  `json:"action"`
            Departments []string  `json:"departments"`
            Banners ImageUrls  `json:"banners"`
            Logo Media  `json:"logo"`
         
    }
    
    // BrandListingResponse used by Catalog
    type BrandListingResponse struct {

        
            Page Page  `json:"page"`
            Items []BrandItem  `json:"items"`
         
    }
    
    // Department used by Catalog
    type Department struct {

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            PriorityOrder float64  `json:"priority_order"`
            Slug string  `json:"slug"`
            Logo Media  `json:"logo"`
         
    }
    
    // DepartmentResponse used by Catalog
    type DepartmentResponse struct {

        
            Items []Department  `json:"items"`
         
    }
    
    // DepartmentIdentifier used by Catalog
    type DepartmentIdentifier struct {

        
            UID float64  `json:"uid"`
            Slug string  `json:"slug"`
         
    }
    
    // ThirdLevelChild used by Catalog
    type ThirdLevelChild struct {

        
            UID float64  `json:"uid"`
            Childs []map[string]interface{}  `json:"childs"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Action Action  `json:"action"`
            Banners ImageUrls  `json:"banners"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // SecondLevelChild used by Catalog
    type SecondLevelChild struct {

        
            UID float64  `json:"uid"`
            Childs []ThirdLevelChild  `json:"childs"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Action Action  `json:"action"`
            Banners ImageUrls  `json:"banners"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // Child used by Catalog
    type Child struct {

        
            UID float64  `json:"uid"`
            Childs []SecondLevelChild  `json:"childs"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Action Action  `json:"action"`
            Banners ImageUrls  `json:"banners"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // CategoryItems used by Catalog
    type CategoryItems struct {

        
            UID float64  `json:"uid"`
            Childs []Child  `json:"childs"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Action Action  `json:"action"`
            Banners ImageUrls  `json:"banners"`
         
    }
    
    // DepartmentCategoryTree used by Catalog
    type DepartmentCategoryTree struct {

        
            Items []CategoryItems  `json:"items"`
            Department string  `json:"department"`
         
    }
    
    // CategoryListingResponse used by Catalog
    type CategoryListingResponse struct {

        
            Departments []DepartmentIdentifier  `json:"departments"`
            Data []DepartmentCategoryTree  `json:"data"`
         
    }
    
    // ApplicationProductListingResponse used by Catalog
    type ApplicationProductListingResponse struct {

        
            Page Page  `json:"page"`
            Filters []ProductFilters  `json:"filters"`
            SortOn []ProductSortOn  `json:"sort_on"`
            Operators map[string]interface{}  `json:"operators"`
            Items []ProductListingDetail  `json:"items"`
         
    }
    
    // ProductDetail used by Catalog
    type ProductDetail struct {

        
            UID float64  `json:"uid"`
            Similars []string  `json:"similars"`
            RatingCount float64  `json:"rating_count"`
            Medias []Media1  `json:"medias"`
            ImageNature string  `json:"image_nature"`
            Highlights []string  `json:"highlights"`
            Rating float64  `json:"rating"`
            Color string  `json:"color"`
            Tryouts []string  `json:"tryouts"`
            Brand ProductBrand  `json:"brand"`
            HasVariant bool  `json:"has_variant"`
            ItemCode string  `json:"item_code"`
            PromoMeta map[string]interface{}  `json:"promo_meta"`
            Type string  `json:"type"`
            Description string  `json:"description"`
            ShortDescription string  `json:"short_description"`
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
            Name string  `json:"name"`
            ProductOnlineDate string  `json:"product_online_date"`
            Slug string  `json:"slug"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            Attributes map[string]interface{}  `json:"attributes"`
            ItemType string  `json:"item_type"`
         
    }
    
    // InventoryPage used by Catalog
    type InventoryPage struct {

        
            HasNext bool  `json:"has_next"`
            ItemTotal float64  `json:"item_total"`
            Type string  `json:"type"`
            HasPrevious bool  `json:"has_previous"`
            NextID string  `json:"next_id"`
         
    }
    
    // InventoryStockResponse used by Catalog
    type InventoryStockResponse struct {

        
            Page InventoryPage  `json:"page"`
            Items []map[string]interface{}  `json:"items"`
         
    }
    
    // ArticleQuery used by Catalog
    type ArticleQuery struct {

        
            ItemID float64  `json:"item_id"`
            IgnoredStores []float64  `json:"ignored_stores"`
            Size string  `json:"size"`
         
    }
    
    // ArticleAssignment used by Catalog
    type ArticleAssignment struct {

        
            Strategy string  `json:"strategy"`
            Level string  `json:"level"`
         
    }
    
    // AssignStoreArticle used by Catalog
    type AssignStoreArticle struct {

        
            Meta map[string]interface{}  `json:"meta"`
            Query ArticleQuery  `json:"query"`
            ArticleAssignment ArticleAssignment  `json:"article_assignment"`
            Quantity float64  `json:"quantity"`
            GroupID string  `json:"group_id"`
         
    }
    
    // AssignStore used by Catalog
    type AssignStore struct {

        
            Pincode string  `json:"pincode"`
            ChannelType string  `json:"channel_type"`
            ChannelIdentifier string  `json:"channel_identifier"`
            StoreIds []float64  `json:"store_ids"`
            AppID string  `json:"app_id"`
            Articles []AssignStoreArticle  `json:"articles"`
            CompanyID float64  `json:"company_id"`
         
    }
    
    // ArticleAssignment1 used by Catalog
    type ArticleAssignment1 struct {

        
            Strategy string  `json:"strategy"`
            Level string  `json:"level"`
         
    }
    
    // StoreAssignResponse used by Catalog
    type StoreAssignResponse struct {

        
            ID string  `json:"_id"`
            UID string  `json:"uid"`
            Meta map[string]interface{}  `json:"meta"`
            Status bool  `json:"status"`
            PriceEffective float64  `json:"price_effective"`
            GroupID string  `json:"group_id"`
            StoreID float64  `json:"store_id"`
            Size string  `json:"size"`
            ArticleAssignment ArticleAssignment1  `json:"article_assignment"`
            ItemID float64  `json:"item_id"`
            PriceMarked float64  `json:"price_marked"`
            StrategyWiseListing []map[string]interface{}  `json:"strategy_wise_listing"`
            Index float64  `json:"index"`
            Quantity float64  `json:"quantity"`
            CompanyID float64  `json:"company_id"`
            SCity string  `json:"s_city"`
            StorePincode float64  `json:"store_pincode"`
         
    }
    
    // UserSerializer1 used by Catalog
    type UserSerializer1 struct {

        
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
            Contact string  `json:"contact"`
         
    }
    
    // GetAddressSerializer used by Catalog
    type GetAddressSerializer struct {

        
            Address2 string  `json:"address2"`
            Pincode float64  `json:"pincode"`
            CountryCode string  `json:"country_code"`
            Country string  `json:"country"`
            State string  `json:"state"`
            City string  `json:"city"`
            Latitude float64  `json:"latitude"`
            Longitude float64  `json:"longitude"`
            Landmark string  `json:"landmark"`
            Address1 string  `json:"address1"`
            AddressType string  `json:"address_type"`
         
    }
    
    // UserSerializer2 used by Catalog
    type UserSerializer2 struct {

        
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
            Contact string  `json:"contact"`
         
    }
    
    // GetCompanySerializer used by Catalog
    type GetCompanySerializer struct {

        
            Stage string  `json:"stage"`
            UID float64  `json:"uid"`
            BusinessType string  `json:"business_type"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            VerifiedBy UserSerializer2  `json:"verified_by"`
            Name string  `json:"name"`
            ModifiedOn string  `json:"modified_on"`
            CompanyType string  `json:"company_type"`
            CreatedOn string  `json:"created_on"`
            VerifiedOn string  `json:"verified_on"`
            RejectReason string  `json:"reject_reason"`
            CreatedBy UserSerializer2  `json:"created_by"`
            ModifiedBy UserSerializer2  `json:"modified_by"`
         
    }
    
    // ProductReturnConfigSerializer used by Catalog
    type ProductReturnConfigSerializer struct {

        
            StoreUID float64  `json:"store_uid"`
            OnSameStore bool  `json:"on_same_store"`
         
    }
    
    // LocationIntegrationType used by Catalog
    type LocationIntegrationType struct {

        
            Inventory string  `json:"inventory"`
            Order string  `json:"order"`
         
    }
    
    // SellerPhoneNumber used by Catalog
    type SellerPhoneNumber struct {

        
            Number string  `json:"number"`
            CountryCode float64  `json:"country_code"`
         
    }
    
    // LocationManagerSerializer used by Catalog
    type LocationManagerSerializer struct {

        
            Email string  `json:"email"`
            MobileNo SellerPhoneNumber  `json:"mobile_no"`
            Name string  `json:"name"`
         
    }
    
    // InvoiceCredSerializer used by Catalog
    type InvoiceCredSerializer struct {

        
            Username string  `json:"username"`
            Password string  `json:"password"`
            Enabled bool  `json:"enabled"`
         
    }
    
    // InvoiceDetailsSerializer used by Catalog
    type InvoiceDetailsSerializer struct {

        
            EInvoice InvoiceCredSerializer  `json:"e_invoice"`
            EWaybill InvoiceCredSerializer  `json:"e_waybill"`
         
    }
    
    // Document used by Catalog
    type Document struct {

        
            LegalName string  `json:"legal_name"`
            Type string  `json:"type"`
            Value string  `json:"value"`
            Verified bool  `json:"verified"`
            URL string  `json:"url"`
         
    }
    
    // LocationTimingSerializer used by Catalog
    type LocationTimingSerializer struct {

        
            Minute float64  `json:"minute"`
            Hour float64  `json:"hour"`
         
    }
    
    // LocationDayWiseSerializer used by Catalog
    type LocationDayWiseSerializer struct {

        
            Weekday string  `json:"weekday"`
            Closing LocationTimingSerializer  `json:"closing"`
            Open bool  `json:"open"`
            Opening LocationTimingSerializer  `json:"opening"`
         
    }
    
    // GetLocationSerializer used by Catalog
    type GetLocationSerializer struct {

        
            UID float64  `json:"uid"`
            VerifiedBy UserSerializer1  `json:"verified_by"`
            PhoneNumber string  `json:"phone_number"`
            ModifiedBy UserSerializer1  `json:"modified_by"`
            NotificationEmails []string  `json:"notification_emails"`
            CreatedOn string  `json:"created_on"`
            Company GetCompanySerializer  `json:"company"`
            ProductReturnConfig ProductReturnConfigSerializer  `json:"product_return_config"`
            StoreType string  `json:"store_type"`
            Address GetAddressSerializer  `json:"address"`
            ModifiedOn string  `json:"modified_on"`
            Code string  `json:"code"`
            Warnings map[string]interface{}  `json:"warnings"`
            CreatedBy UserSerializer1  `json:"created_by"`
            IntegrationType LocationIntegrationType  `json:"integration_type"`
            DisplayName string  `json:"display_name"`
            Stage string  `json:"stage"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            Manager LocationManagerSerializer  `json:"manager"`
            GstCredentials InvoiceDetailsSerializer  `json:"gst_credentials"`
            Name string  `json:"name"`
            VerifiedOn string  `json:"verified_on"`
            Documents []Document  `json:"documents"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Timing []LocationDayWiseSerializer  `json:"timing"`
         
    }
    
    // LocationListSerializer used by Catalog
    type LocationListSerializer struct {

        
            Page Page  `json:"page"`
            Items []GetLocationSerializer  `json:"items"`
         
    }
    
    // ApplicationBrandJson used by Catalog
    type ApplicationBrandJson struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // ApplicationCategoryJson used by Catalog
    type ApplicationCategoryJson struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // ApplicationDepartment used by Catalog
    type ApplicationDepartment struct {

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            IsActive bool  `json:"is_active"`
            AppID string  `json:"app_id"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Logo string  `json:"logo"`
         
    }
    
    // ApplicationDepartmentListingResponse used by Catalog
    type ApplicationDepartmentListingResponse struct {

        
            Page Page  `json:"page"`
            Items []ApplicationDepartment  `json:"items"`
         
    }
    
    // ApplicationDepartmentJson used by Catalog
    type ApplicationDepartmentJson struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // ApplicationStoreJson used by Catalog
    type ApplicationStoreJson struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    

    
    // UserSerializer used by CompanyProfile
    type UserSerializer struct {

        
            UserID string  `json:"user_id"`
            Contact string  `json:"contact"`
            Username string  `json:"username"`
         
    }
    
    // Document used by CompanyProfile
    type Document struct {

        
            Value string  `json:"value"`
            URL string  `json:"url"`
            Verified bool  `json:"verified"`
            LegalName string  `json:"legal_name"`
            Type string  `json:"type"`
         
    }
    
    // GetAddressSerializer used by CompanyProfile
    type GetAddressSerializer struct {

        
            Longitude float64  `json:"longitude"`
            City string  `json:"city"`
            Address1 string  `json:"address1"`
            State string  `json:"state"`
            CountryCode string  `json:"country_code"`
            Address2 string  `json:"address2"`
            Pincode float64  `json:"pincode"`
            AddressType string  `json:"address_type"`
            Country string  `json:"country"`
            Landmark string  `json:"landmark"`
            Latitude float64  `json:"latitude"`
         
    }
    
    // SellerPhoneNumber used by CompanyProfile
    type SellerPhoneNumber struct {

        
            Number string  `json:"number"`
            CountryCode float64  `json:"country_code"`
         
    }
    
    // ContactDetails used by CompanyProfile
    type ContactDetails struct {

        
            Phone []SellerPhoneNumber  `json:"phone"`
            Emails []string  `json:"emails"`
         
    }
    
    // BusinessCountryInfo used by CompanyProfile
    type BusinessCountryInfo struct {

        
            Country string  `json:"country"`
            CountryCode string  `json:"country_code"`
         
    }
    
    // CompanyTaxesSerializer used by CompanyProfile
    type CompanyTaxesSerializer struct {

        
            EffectiveDate string  `json:"effective_date"`
            Enable bool  `json:"enable"`
            Rate float64  `json:"rate"`
         
    }
    
    // Website used by CompanyProfile
    type Website struct {

        
            URL string  `json:"url"`
         
    }
    
    // BusinessDetails used by CompanyProfile
    type BusinessDetails struct {

        
            Website Website  `json:"website"`
         
    }
    
    // GetCompanyProfileSerializerResponse used by CompanyProfile
    type GetCompanyProfileSerializerResponse struct {

        
            FranchiseEnabled bool  `json:"franchise_enabled"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            UID float64  `json:"uid"`
            Warnings map[string]interface{}  `json:"warnings"`
            ModifiedOn string  `json:"modified_on"`
            Documents []Document  `json:"documents"`
            Mode string  `json:"mode"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            VerifiedOn string  `json:"verified_on"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            ContactDetails ContactDetails  `json:"contact_details"`
            BusinessCountryInfo BusinessCountryInfo  `json:"business_country_info"`
            Name string  `json:"name"`
            BusinessInfo string  `json:"business_info"`
            Stage string  `json:"stage"`
            CompanyType string  `json:"company_type"`
            NotificationEmails []string  `json:"notification_emails"`
            Taxes []CompanyTaxesSerializer  `json:"taxes"`
            CreatedOn string  `json:"created_on"`
            BusinessDetails BusinessDetails  `json:"business_details"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            BusinessType string  `json:"business_type"`
            CreatedBy UserSerializer  `json:"created_by"`
         
    }
    
    // ErrorResponse used by CompanyProfile
    type ErrorResponse struct {

        
            Status float64  `json:"status"`
            Code string  `json:"code"`
            Message string  `json:"message"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // CompanyTaxesSerializer1 used by CompanyProfile
    type CompanyTaxesSerializer1 struct {

        
            EffectiveDate string  `json:"effective_date"`
            Enable bool  `json:"enable"`
            Rate float64  `json:"rate"`
         
    }
    
    // CreateUpdateAddressSerializer used by CompanyProfile
    type CreateUpdateAddressSerializer struct {

        
            Longitude float64  `json:"longitude"`
            State string  `json:"state"`
            Address1 string  `json:"address1"`
            City string  `json:"city"`
            CountryCode string  `json:"country_code"`
            Address2 string  `json:"address2"`
            Pincode float64  `json:"pincode"`
            AddressType string  `json:"address_type"`
            Country string  `json:"country"`
            Landmark string  `json:"landmark"`
            Latitude float64  `json:"latitude"`
         
    }
    
    // UpdateCompany used by CompanyProfile
    type UpdateCompany struct {

        
            Taxes []CompanyTaxesSerializer1  `json:"taxes"`
            Name string  `json:"name"`
            BusinessDetails BusinessDetails  `json:"business_details"`
            FranchiseEnabled bool  `json:"franchise_enabled"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Warnings map[string]interface{}  `json:"warnings"`
            Documents []Document  `json:"documents"`
            ContactDetails ContactDetails  `json:"contact_details"`
            BusinessType string  `json:"business_type"`
            BusinessInfo string  `json:"business_info"`
            Addresses []CreateUpdateAddressSerializer  `json:"addresses"`
            RejectReason string  `json:"reject_reason"`
            CompanyType string  `json:"company_type"`
            NotificationEmails []string  `json:"notification_emails"`
         
    }
    
    // ProfileSuccessResponse used by CompanyProfile
    type ProfileSuccessResponse struct {

        
            UID float64  `json:"uid"`
            Success bool  `json:"success"`
         
    }
    
    // DocumentsObj used by CompanyProfile
    type DocumentsObj struct {

        
            Verified float64  `json:"verified"`
            Pending float64  `json:"pending"`
         
    }
    
    // MetricsSerializer used by CompanyProfile
    type MetricsSerializer struct {

        
            StoreDocuments DocumentsObj  `json:"store_documents"`
            Store DocumentsObj  `json:"store"`
            UID float64  `json:"uid"`
            Product DocumentsObj  `json:"product"`
            Brand DocumentsObj  `json:"brand"`
            Stage string  `json:"stage"`
            CompanyDocuments DocumentsObj  `json:"company_documents"`
         
    }
    
    // BrandBannerSerializer used by CompanyProfile
    type BrandBannerSerializer struct {

        
            Landscape string  `json:"landscape"`
            Portrait string  `json:"portrait"`
         
    }
    
    // GetBrandResponseSerializer used by CompanyProfile
    type GetBrandResponseSerializer struct {

        
            VerifiedBy UserSerializer  `json:"verified_by"`
            UID float64  `json:"uid"`
            SlugKey string  `json:"slug_key"`
            Warnings map[string]interface{}  `json:"warnings"`
            ModifiedOn string  `json:"modified_on"`
            Logo string  `json:"logo"`
            Mode string  `json:"mode"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            VerifiedOn string  `json:"verified_on"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            Name string  `json:"name"`
            Synonyms []string  `json:"synonyms"`
            Banner BrandBannerSerializer  `json:"banner"`
            Stage string  `json:"stage"`
            Description string  `json:"description"`
            CreatedOn string  `json:"created_on"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            CreatedBy UserSerializer  `json:"created_by"`
            RejectReason string  `json:"reject_reason"`
         
    }
    
    // CreateUpdateBrandRequestSerializer used by CompanyProfile
    type CreateUpdateBrandRequestSerializer struct {

        
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            Name string  `json:"name"`
            UID float64  `json:"uid"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Synonyms []string  `json:"synonyms"`
            BrandTier string  `json:"brand_tier"`
            Banner BrandBannerSerializer  `json:"banner"`
            Logo string  `json:"logo"`
            Description string  `json:"description"`
            CompanyID float64  `json:"company_id"`
         
    }
    
    // Page used by CompanyProfile
    type Page struct {

        
            ItemTotal float64  `json:"item_total"`
            NextID string  `json:"next_id"`
            HasPrevious bool  `json:"has_previous"`
            HasNext bool  `json:"has_next"`
            Current float64  `json:"current"`
            Type string  `json:"type"`
            Size float64  `json:"size"`
         
    }
    
    // CompanySocialAccounts used by CompanyProfile
    type CompanySocialAccounts struct {

        
            URL string  `json:"url"`
            Name string  `json:"name"`
         
    }
    
    // CompanyDetails used by CompanyProfile
    type CompanyDetails struct {

        
            WebsiteURL string  `json:"website_url"`
            Socials []CompanySocialAccounts  `json:"socials"`
         
    }
    
    // CompanySerializer used by CompanyProfile
    type CompanySerializer struct {

        
            BusinessCountryInfo BusinessCountryInfo  `json:"business_country_info"`
            Name string  `json:"name"`
            CreatedOn string  `json:"created_on"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            UID float64  `json:"uid"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            VerifiedOn string  `json:"verified_on"`
            ModifiedOn string  `json:"modified_on"`
            BusinessType string  `json:"business_type"`
            Details CompanyDetails  `json:"details"`
            CreatedBy UserSerializer  `json:"created_by"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            MarketChannels []string  `json:"market_channels"`
            RejectReason string  `json:"reject_reason"`
            Stage string  `json:"stage"`
            CompanyType string  `json:"company_type"`
            NotificationEmails []string  `json:"notification_emails"`
         
    }
    
    // CompanyBrandSerializer used by CompanyProfile
    type CompanyBrandSerializer struct {

        
            CreatedOn string  `json:"created_on"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            UID float64  `json:"uid"`
            RejectReason string  `json:"reject_reason"`
            Warnings map[string]interface{}  `json:"warnings"`
            VerifiedOn string  `json:"verified_on"`
            Company CompanySerializer  `json:"company"`
            ModifiedOn string  `json:"modified_on"`
            CreatedBy UserSerializer  `json:"created_by"`
            Brand GetBrandResponseSerializer  `json:"brand"`
            Stage string  `json:"stage"`
         
    }
    
    // CompanyBrandListSerializer used by CompanyProfile
    type CompanyBrandListSerializer struct {

        
            Page Page  `json:"page"`
            Items []CompanyBrandSerializer  `json:"items"`
         
    }
    
    // CompanyBrandPostRequestSerializer used by CompanyProfile
    type CompanyBrandPostRequestSerializer struct {

        
            Company float64  `json:"company"`
            UID float64  `json:"uid"`
            Brands []float64  `json:"brands"`
         
    }
    
    // LocationManagerSerializer used by CompanyProfile
    type LocationManagerSerializer struct {

        
            MobileNo SellerPhoneNumber  `json:"mobile_no"`
            Name string  `json:"name"`
            Email string  `json:"email"`
         
    }
    
    // LocationTimingSerializer used by CompanyProfile
    type LocationTimingSerializer struct {

        
            Hour float64  `json:"hour"`
            Minute float64  `json:"minute"`
         
    }
    
    // LocationDayWiseSerializer used by CompanyProfile
    type LocationDayWiseSerializer struct {

        
            Weekday string  `json:"weekday"`
            Open bool  `json:"open"`
            Closing LocationTimingSerializer  `json:"closing"`
            Opening LocationTimingSerializer  `json:"opening"`
         
    }
    
    // ProductReturnConfigSerializer used by CompanyProfile
    type ProductReturnConfigSerializer struct {

        
            OnSameStore bool  `json:"on_same_store"`
            StoreUID float64  `json:"store_uid"`
         
    }
    
    // InvoiceCredSerializer used by CompanyProfile
    type InvoiceCredSerializer struct {

        
            Enabled bool  `json:"enabled"`
            Username string  `json:"username"`
            Password string  `json:"password"`
         
    }
    
    // InvoiceDetailsSerializer used by CompanyProfile
    type InvoiceDetailsSerializer struct {

        
            EWaybill InvoiceCredSerializer  `json:"e_waybill"`
            EInvoice InvoiceCredSerializer  `json:"e_invoice"`
         
    }
    
    // GetCompanySerializer used by CompanyProfile
    type GetCompanySerializer struct {

        
            Name string  `json:"name"`
            CreatedOn string  `json:"created_on"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            UID float64  `json:"uid"`
            VerifiedOn string  `json:"verified_on"`
            ModifiedOn string  `json:"modified_on"`
            BusinessType string  `json:"business_type"`
            CreatedBy UserSerializer  `json:"created_by"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            RejectReason string  `json:"reject_reason"`
            Stage string  `json:"stage"`
            CompanyType string  `json:"company_type"`
         
    }
    
    // HolidayDateSerializer used by CompanyProfile
    type HolidayDateSerializer struct {

        
            StartDate string  `json:"start_date"`
            EndDate string  `json:"end_date"`
         
    }
    
    // HolidaySchemaSerializer used by CompanyProfile
    type HolidaySchemaSerializer struct {

        
            HolidayType string  `json:"holiday_type"`
            Title string  `json:"title"`
            Date HolidayDateSerializer  `json:"date"`
         
    }
    
    // GetLocationSerializer used by CompanyProfile
    type GetLocationSerializer struct {

        
            Code string  `json:"code"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            UID float64  `json:"uid"`
            Warnings map[string]interface{}  `json:"warnings"`
            ModifiedOn string  `json:"modified_on"`
            DisplayName string  `json:"display_name"`
            Documents []Document  `json:"documents"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            VerifiedOn string  `json:"verified_on"`
            Manager LocationManagerSerializer  `json:"manager"`
            Timing []LocationDayWiseSerializer  `json:"timing"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            ProductReturnConfig ProductReturnConfigSerializer  `json:"product_return_config"`
            Name string  `json:"name"`
            StoreType string  `json:"store_type"`
            Address GetAddressSerializer  `json:"address"`
            Stage string  `json:"stage"`
            NotificationEmails []string  `json:"notification_emails"`
            GstCredentials InvoiceDetailsSerializer  `json:"gst_credentials"`
            CreatedOn string  `json:"created_on"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Company GetCompanySerializer  `json:"company"`
            PhoneNumber string  `json:"phone_number"`
            CreatedBy UserSerializer  `json:"created_by"`
            Holiday []HolidaySchemaSerializer  `json:"holiday"`
         
    }
    
    // LocationListSerializer used by CompanyProfile
    type LocationListSerializer struct {

        
            Page Page  `json:"page"`
            Items []GetLocationSerializer  `json:"items"`
         
    }
    
    // AddressSerializer used by CompanyProfile
    type AddressSerializer struct {

        
            Longitude float64  `json:"longitude"`
            City string  `json:"city"`
            Address1 string  `json:"address1"`
            State string  `json:"state"`
            CountryCode string  `json:"country_code"`
            Address2 string  `json:"address2"`
            Pincode float64  `json:"pincode"`
            AddressType string  `json:"address_type"`
            Country string  `json:"country"`
            Landmark string  `json:"landmark"`
            Latitude float64  `json:"latitude"`
         
    }
    
    // LocationSerializer used by CompanyProfile
    type LocationSerializer struct {

        
            Code string  `json:"code"`
            NotificationEmails []string  `json:"notification_emails"`
            Name string  `json:"name"`
            GstCredentials InvoiceDetailsSerializer  `json:"gst_credentials"`
            StoreType string  `json:"store_type"`
            UID float64  `json:"uid"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Warnings map[string]interface{}  `json:"warnings"`
            Company float64  `json:"company"`
            DisplayName string  `json:"display_name"`
            Manager LocationManagerSerializer  `json:"manager"`
            Address AddressSerializer  `json:"address"`
            Timing []LocationDayWiseSerializer  `json:"timing"`
            Stage string  `json:"stage"`
            Documents []Document  `json:"documents"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            Holiday []HolidaySchemaSerializer  `json:"holiday"`
            ProductReturnConfig ProductReturnConfigSerializer  `json:"product_return_config"`
         
    }
    
    // BulkLocationSerializer used by CompanyProfile
    type BulkLocationSerializer struct {

        
            Data []LocationSerializer  `json:"data"`
         
    }
    

    
    // FailedResponse used by FileStorage
    type FailedResponse struct {

        
            Message string  `json:"message"`
         
    }
    
    // CDN used by FileStorage
    type CDN struct {

        
            URL string  `json:"url"`
            AbsoluteURL string  `json:"absolute_url"`
            RelativeURL string  `json:"relative_url"`
         
    }
    
    // Upload used by FileStorage
    type Upload struct {

        
            Expiry float64  `json:"expiry"`
            URL string  `json:"url"`
         
    }
    
    // StartResponse used by FileStorage
    type StartResponse struct {

        
            FileName string  `json:"file_name"`
            FilePath string  `json:"file_path"`
            ContentType string  `json:"content_type"`
            Method string  `json:"method"`
            Namespace string  `json:"namespace"`
            Operation string  `json:"operation"`
            Size float64  `json:"size"`
            Upload Upload  `json:"upload"`
            Cdn CDN  `json:"cdn"`
            Tags []string  `json:"tags"`
         
    }
    
    // StartRequest used by FileStorage
    type StartRequest struct {

        
            FileName string  `json:"file_name"`
            ContentType string  `json:"content_type"`
            Size float64  `json:"size"`
            Tags []string  `json:"tags"`
            Params map[string]interface{}  `json:"params"`
         
    }
    
    // CompleteResponse used by FileStorage
    type CompleteResponse struct {

        
            ID string  `json:"_id"`
            FileName string  `json:"file_name"`
            FilePath string  `json:"file_path"`
            ContentType string  `json:"content_type"`
            Namespace string  `json:"namespace"`
            Operation string  `json:"operation"`
            Size float64  `json:"size"`
            Upload Upload  `json:"upload"`
            Cdn CDN  `json:"cdn"`
            Success bool  `json:"success"`
            Tags []string  `json:"tags"`
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // Opts used by FileStorage
    type Opts struct {

        
            Attempts float64  `json:"attempts"`
            Timestamp float64  `json:"timestamp"`
            Delay float64  `json:"delay"`
         
    }
    
    // CopyFileTask used by FileStorage
    type CopyFileTask struct {

        
            ID string  `json:"id"`
            Name string  `json:"name"`
            Data BulkRequest  `json:"data"`
            Opts Opts  `json:"opts"`
            Progress float64  `json:"progress"`
            Delay float64  `json:"delay"`
            Timestamp float64  `json:"timestamp"`
            AttemptsMade float64  `json:"attempts_made"`
            Stacktrace []string  `json:"stacktrace"`
            FinishedOn float64  `json:"finished_on"`
            ProcessedOn float64  `json:"processed_on"`
         
    }
    
    // BulkUploadResponse used by FileStorage
    type BulkUploadResponse struct {

        
            TrackingURL string  `json:"tracking_url"`
            Task CopyFileTask  `json:"task"`
         
    }
    
    // ReqConfiguration used by FileStorage
    type ReqConfiguration struct {

        
            Concurrency float64  `json:"concurrency"`
         
    }
    
    // Destination used by FileStorage
    type Destination struct {

        
            Namespace string  `json:"namespace"`
            Rewrite string  `json:"rewrite"`
            Basepath string  `json:"basepath"`
         
    }
    
    // BulkRequest used by FileStorage
    type BulkRequest struct {

        
            Urls []string  `json:"urls"`
            Destination Destination  `json:"destination"`
            Configuration ReqConfiguration  `json:"configuration"`
         
    }
    
    // Urls used by FileStorage
    type Urls struct {

        
            URL string  `json:"url"`
            SignedURL string  `json:"signed_url"`
            Expiry float64  `json:"expiry"`
         
    }
    
    // SignUrlResponse used by FileStorage
    type SignUrlResponse struct {

        
            Urls []Urls  `json:"urls"`
         
    }
    
    // SignUrlRequest used by FileStorage
    type SignUrlRequest struct {

        
            Expiry float64  `json:"expiry"`
            Urls []string  `json:"urls"`
         
    }
    
    // Page used by FileStorage
    type Page struct {

        
            ItemTotal float64  `json:"item_total"`
            NextID string  `json:"next_id"`
            HasPrevious bool  `json:"has_previous"`
            HasNext bool  `json:"has_next"`
            Current float64  `json:"current"`
            Type string  `json:"type"`
            Size float64  `json:"size"`
         
    }
    
    // DbRecord used by FileStorage
    type DbRecord struct {

        
            Success bool  `json:"success"`
            Tags []string  `json:"tags"`
            ID string  `json:"_id"`
            FileName string  `json:"file_name"`
            Operation string  `json:"operation"`
            Namespace string  `json:"namespace"`
            ContentType string  `json:"content_type"`
            FilePath string  `json:"file_path"`
            Upload Upload  `json:"upload"`
            Cdn CDN  `json:"cdn"`
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // BrowseResponse used by FileStorage
    type BrowseResponse struct {

        
            Items []DbRecord  `json:"items"`
            Page Page  `json:"page"`
         
    }
    

    
    // RedirectDevice used by Share
    type RedirectDevice struct {

        
            Link string  `json:"link"`
            Type string  `json:"type"`
         
    }
    
    // WebRedirect used by Share
    type WebRedirect struct {

        
            Link string  `json:"link"`
            Type string  `json:"type"`
         
    }
    
    // Redirects used by Share
    type Redirects struct {

        
            Ios RedirectDevice  `json:"ios"`
            Android RedirectDevice  `json:"android"`
            Web WebRedirect  `json:"web"`
            ForceWeb bool  `json:"force_web"`
         
    }
    
    // CampaignShortLink used by Share
    type CampaignShortLink struct {

        
            Source string  `json:"source"`
            Medium string  `json:"medium"`
         
    }
    
    // Attribution used by Share
    type Attribution struct {

        
            CampaignCookieExpiry string  `json:"campaign_cookie_expiry"`
         
    }
    
    // SocialMediaTags used by Share
    type SocialMediaTags struct {

        
            Title string  `json:"title"`
            Description string  `json:"description"`
            Image string  `json:"image"`
         
    }
    
    // ShortLinkReq used by Share
    type ShortLinkReq struct {

        
            Title string  `json:"title"`
            URL string  `json:"url"`
            Hash string  `json:"hash"`
            Active bool  `json:"active"`
            ExpireAt string  `json:"expire_at"`
            EnableTracking bool  `json:"enable_tracking"`
            Personalized bool  `json:"personalized"`
            Campaign CampaignShortLink  `json:"campaign"`
            Redirects Redirects  `json:"redirects"`
            Attribution Attribution  `json:"attribution"`
            SocialMediaTags SocialMediaTags  `json:"social_media_tags"`
            Count float64  `json:"count"`
         
    }
    
    // UrlInfo used by Share
    type UrlInfo struct {

        
            Original string  `json:"original"`
            Short string  `json:"short"`
            Hash string  `json:"hash"`
         
    }
    
    // ShortLinkRes used by Share
    type ShortLinkRes struct {

        
            Title string  `json:"title"`
            URL UrlInfo  `json:"url"`
            CreatedBy string  `json:"created_by"`
            AppRedirect bool  `json:"app_redirect"`
            Fallback string  `json:"fallback"`
            Active bool  `json:"active"`
            ID string  `json:"_id"`
            EnableTracking bool  `json:"enable_tracking"`
            ExpireAt string  `json:"expire_at"`
            Application string  `json:"application"`
            UserID string  `json:"user_id"`
            CreatedAt string  `json:"created_at"`
            Meta map[string]interface{}  `json:"meta"`
            UpdatedAt string  `json:"updated_at"`
            Personalized bool  `json:"personalized"`
            Campaign CampaignShortLink  `json:"campaign"`
            Redirects Redirects  `json:"redirects"`
            Attribution Attribution  `json:"attribution"`
            SocialMediaTags SocialMediaTags  `json:"social_media_tags"`
            Count float64  `json:"count"`
         
    }
    
    // Page used by Share
    type Page struct {

        
            ItemTotal float64  `json:"item_total"`
            NextID string  `json:"next_id"`
            HasPrevious bool  `json:"has_previous"`
            HasNext bool  `json:"has_next"`
            Current float64  `json:"current"`
            Type string  `json:"type"`
            Size float64  `json:"size"`
         
    }
    
    // ShortLinkList used by Share
    type ShortLinkList struct {

        
            Items []ShortLinkRes  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // ErrorRes used by Share
    type ErrorRes struct {

        
            Message string  `json:"message"`
         
    }
    

    
    // DataTresholdDTO used by Inventory
    type DataTresholdDTO struct {

        
            MinPrice float64  `json:"min_price"`
            SafeStock float64  `json:"safe_stock"`
            PeriodThreshold float64  `json:"period_threshold"`
            PeriodThresholdType string  `json:"period_threshold_type"`
            PeriodTypeList []GenericDTO  `json:"period_type_list"`
         
    }
    
    // GenericDTO used by Inventory
    type GenericDTO struct {

        
            Text string  `json:"text"`
            Value map[string]interface{}  `json:"value"`
         
    }
    
    // JobConfigDTO used by Inventory
    type JobConfigDTO struct {

        
            IntegrationData map[string]map[string]interface{}  `json:"integration_data"`
            CompanyName string  `json:"company_name"`
            Integration string  `json:"integration"`
            CompanyID float64  `json:"company_id"`
            TaskDetails TaskDTO  `json:"task_details"`
            ThresholdDetails DataTresholdDTO  `json:"threshold_details"`
            JobCode string  `json:"job_code"`
            Alias string  `json:"alias"`
         
    }
    
    // TaskDTO used by Inventory
    type TaskDTO struct {

        
            Type float64  `json:"type"`
            GroupList []GenericDTO  `json:"group_list"`
         
    }
    
    // Page used by Inventory
    type Page struct {

        
            ItemTotal float64  `json:"item_total"`
            NextID string  `json:"next_id"`
            HasPrevious bool  `json:"has_previous"`
            HasNext bool  `json:"has_next"`
            Current float64  `json:"current"`
            Type string  `json:"type"`
            Size float64  `json:"size"`
         
    }
    
    // ResponseEnvelopeString used by Inventory
    type ResponseEnvelopeString struct {

        
            Timestamp string  `json:"timestamp"`
            Status float64  `json:"status"`
            Error string  `json:"error"`
            Exception string  `json:"exception"`
            Message string  `json:"message"`
            TotalTimeTakenInMillis float64  `json:"total_time_taken_in_millis"`
            HttpStatus string  `json:"http_status"`
            Items string  `json:"items"`
            Payload string  `json:"payload"`
            TraceID string  `json:"trace_id"`
            Page Page  `json:"page"`
         
    }
    
    // KafkaMetaModel used by Inventory
    type KafkaMetaModel struct {

        
            JobType string  `json:"job_type"`
            BatchID float64  `json:"batch_id"`
            Action string  `json:"action"`
            Trace []string  `json:"trace"`
            CreatedOn string  `json:"created_on"`
            CreatedTimestamp float64  `json:"created_timestamp"`
         
    }
    
    // SuppressStoreModel used by Inventory
    type SuppressStoreModel struct {

        
            Stores []float64  `json:"stores"`
         
    }
    
    // SuppressStorePayload used by Inventory
    type SuppressStorePayload struct {

        
            Payload []SuppressStoreModel  `json:"payload"`
            Meta KafkaMetaModel  `json:"meta"`
         
    }
    
    // KafkaResponse used by Inventory
    type KafkaResponse struct {

        
            Offset float64  `json:"offset"`
            Partition float64  `json:"partition"`
         
    }
    
    // ResponseEnvelopeKafkaResponse used by Inventory
    type ResponseEnvelopeKafkaResponse struct {

        
            Timestamp string  `json:"timestamp"`
            Status float64  `json:"status"`
            Error string  `json:"error"`
            Exception string  `json:"exception"`
            Message string  `json:"message"`
            TotalTimeTakenInMillis float64  `json:"total_time_taken_in_millis"`
            HttpStatus string  `json:"http_status"`
            Items KafkaResponse  `json:"items"`
            Payload KafkaResponse  `json:"payload"`
            TraceID string  `json:"trace_id"`
            Page Page  `json:"page"`
         
    }
    
    // GCompany used by Inventory
    type GCompany struct {

        
            ID string  `json:"_id"`
            Integration string  `json:"integration"`
            Level string  `json:"level"`
            UID float64  `json:"uid"`
            Opted bool  `json:"opted"`
            Permissions []string  `json:"permissions"`
            Token string  `json:"token"`
            Name string  `json:"name"`
            Stores []GStore  `json:"stores"`
         
    }
    
    // GStore used by Inventory
    type GStore struct {

        
            ID string  `json:"_id"`
            Integration string  `json:"integration"`
            Level string  `json:"level"`
            UID float64  `json:"uid"`
            Opted bool  `json:"opted"`
            Permissions []string  `json:"permissions"`
            Token string  `json:"token"`
            Code string  `json:"code"`
            Name string  `json:"name"`
            Data StoreData  `json:"data"`
         
    }
    
    // Metum used by Inventory
    type Metum struct {

        
            ID string  `json:"_id"`
            Name string  `json:"name"`
            Value string  `json:"value"`
         
    }
    
    // ResponseEnvelopeListSlingshotConfigurationDetail used by Inventory
    type ResponseEnvelopeListSlingshotConfigurationDetail struct {

        
            Timestamp string  `json:"timestamp"`
            Status float64  `json:"status"`
            Error string  `json:"error"`
            Exception string  `json:"exception"`
            Message string  `json:"message"`
            TotalTimeTakenInMillis float64  `json:"total_time_taken_in_millis"`
            HttpStatus string  `json:"http_status"`
            Items []SlingshotConfigurationDetail  `json:"items"`
            Payload []SlingshotConfigurationDetail  `json:"payload"`
            TraceID string  `json:"trace_id"`
            Page Page  `json:"page"`
         
    }
    
    // SlingshotConfigurationDetail used by Inventory
    type SlingshotConfigurationDetail struct {

        
            Integration SlingshotIntegration  `json:"integration"`
            Companies []GCompany  `json:"companies"`
         
    }
    
    // SlingshotIntegration used by Inventory
    type SlingshotIntegration struct {

        
            ID string  `json:"_id"`
            Description string  `json:"description"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Meta []Metum  `json:"meta"`
         
    }
    
    // StoreData used by Inventory
    type StoreData struct {

        
            LocationID string  `json:"location_id"`
         
    }
    
    // AWSS3config used by Inventory
    type AWSS3config struct {

        
            Bucket string  `json:"bucket"`
            Region string  `json:"region"`
            Dir string  `json:"dir"`
            AccessKey string  `json:"access_key"`
            SecretKey string  `json:"secret_key"`
            LocalFilePath string  `json:"local_file_path"`
            ArchivePath string  `json:"archive_path"`
            Archive bool  `json:"archive"`
            Delete bool  `json:"delete"`
            Unzip bool  `json:"unzip"`
            ZipFormat string  `json:"zip_format"`
            FileRegex string  `json:"file_regex"`
            ArchiveConfig ArchiveConfig  `json:"archive_config"`
         
    }
    
    // ArchiveConfig used by Inventory
    type ArchiveConfig struct {

        
            Delete bool  `json:"delete"`
            Archive bool  `json:"archive"`
            ArchiveDir string  `json:"archive_dir"`
         
    }
    
    // Audit used by Inventory
    type Audit struct {

        
            CreatedBy string  `json:"created_by"`
            ModifiedBy string  `json:"modified_by"`
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // CatalogMasterConfig used by Inventory
    type CatalogMasterConfig struct {

        
            SourceSlug string  `json:"source_slug"`
         
    }
    
    // CompanyConfig used by Inventory
    type CompanyConfig struct {

        
            CompanyID float64  `json:"company_id"`
            ExcludeSteps []float64  `json:"exclude_steps"`
            HiddenClosetKeys []string  `json:"hidden_closet_keys"`
            OpenTags map[string]map[string]interface{}  `json:"open_tags"`
            TaxIdentifiers []string  `json:"tax_identifiers"`
            DeleteQuantityThreshold float64  `json:"delete_quantity_threshold"`
         
    }
    
    // DBConfig used by Inventory
    type DBConfig struct {

        
            Vendor string  `json:"vendor"`
            Host string  `json:"host"`
            Port float64  `json:"port"`
            Username string  `json:"username"`
            Password string  `json:"password"`
            Dbname string  `json:"dbname"`
            Query string  `json:"query"`
            Procedure bool  `json:"procedure"`
            DriverClass string  `json:"driver_class"`
            JdbcURL string  `json:"jdbc_url"`
            OptionalProperties map[string]string  `json:"optional_properties"`
         
    }
    
    // DBConnectionProfile used by Inventory
    type DBConnectionProfile struct {

        
            Inventory string  `json:"inventory"`
         
    }
    
    // DBParamConfig used by Inventory
    type DBParamConfig struct {

        
            Params map[string]map[string]interface{}  `json:"params"`
         
    }
    
    // DefaultHeadersDTO used by Inventory
    type DefaultHeadersDTO struct {

        
            Store PropBeanDTO  `json:"store"`
            IntfArticleID PropBeanDTO  `json:"intf_article_id"`
            PriceEffective PropBeanDTO  `json:"price_effective"`
            Quantity PropBeanDTO  `json:"quantity"`
         
    }
    
    // DocMappingConfig used by Inventory
    type DocMappingConfig struct {

        
            Properties map[string]map[string]interface{}  `json:"properties"`
            JunkDataThresholdCount float64  `json:"junk_data_threshold_count"`
            PropBeanConfigs []PropBeanConfig  `json:"prop_bean_configs"`
            UnwindField string  `json:"unwind_field"`
            DefaultHeaders DefaultHeadersDTO  `json:"default_headers"`
         
    }
    
    // EmailConfig used by Inventory
    type EmailConfig struct {

        
            Recepient string  `json:"recepient"`
            Host string  `json:"host"`
            Username string  `json:"username"`
            Password string  `json:"password"`
            Unzip bool  `json:"unzip"`
            ReadFromContent bool  `json:"read_from_content"`
            FilterBasedOnRecepients bool  `json:"filter_based_on_recepients"`
            Pcol string  `json:"pcol"`
            SubjectLineRegex string  `json:"subject_line_regex"`
            SenderAddress string  `json:"sender_address"`
            LocalDir string  `json:"local_dir"`
            FolderNameHierarchies []string  `json:"folder_name_hierarchies"`
            AttachmentRegex string  `json:"attachment_regex"`
            BodyContentRegex string  `json:"body_content_regex"`
            PasswordProtected bool  `json:"password_protected"`
            ZipFormat string  `json:"zip_format"`
            AttachmentMandate bool  `json:"attachment_mandate"`
            FilterFilesAfterExtraction bool  `json:"filter_files_after_extraction"`
            ArchiveConfig ArchiveConfig  `json:"archive_config"`
            ReadAllUnreadMails bool  `json:"read_all_unread_mails"`
            ContentType string  `json:"content_type"`
            DownloadableLink bool  `json:"downloadable_link"`
            Properties map[string]string  `json:"properties"`
         
    }
    
    // FTPConfig used by Inventory
    type FTPConfig struct {

        
            Host string  `json:"host"`
            Port float64  `json:"port"`
            Username string  `json:"username"`
            Password string  `json:"password"`
            Unzip bool  `json:"unzip"`
            Retries float64  `json:"retries"`
            Interval float64  `json:"interval"`
            LocalDir string  `json:"local_dir"`
            RemoteDir string  `json:"remote_dir"`
            ZipFileRegex string  `json:"zip_file_regex"`
            ArchiveConfig ArchiveConfig  `json:"archive_config"`
            FileRegex string  `json:"file_regex"`
            ZipFormat string  `json:"zip_format"`
            ReadAllFiles bool  `json:"read_all_files"`
         
    }
    
    // FileConfig used by Inventory
    type FileConfig struct {

        
            Delimiter string  `json:"delimiter"`
            Charset string  `json:"charset"`
            Properties map[string]map[string]interface{}  `json:"properties"`
            FileHasHeader bool  `json:"file_has_header"`
            HeaderIndex float64  `json:"header_index"`
            HeaderArray []string  `json:"header_array"`
            DataStartIndex float64  `json:"data_start_index"`
            PropBeanConfigs []PropBeanConfig  `json:"prop_bean_configs"`
            JunkDataThresholdCount float64  `json:"junk_data_threshold_count"`
            FileType string  `json:"file_type"`
            LineValidityCheck bool  `json:"line_validity_check"`
            SheetNames []string  `json:"sheet_names"`
            ReadAllSheets bool  `json:"read_all_sheets"`
            QuoteChar string  `json:"quote_char"`
            EscapeChar string  `json:"escape_char"`
            DefaultHeaders DefaultHeadersDTO  `json:"default_headers"`
         
    }
    
    // GoogleSpreadSheetConfig used by Inventory
    type GoogleSpreadSheetConfig struct {

        
            Range string  `json:"range"`
            SheetID string  `json:"sheet_id"`
            ClientSecretLocation string  `json:"client_secret_location"`
            CredStorageDirectory string  `json:"cred_storage_directory"`
            LocalDir string  `json:"local_dir"`
            ArchiveConfig ArchiveConfig  `json:"archive_config"`
         
    }
    
    // HttpConfig used by Inventory
    type HttpConfig struct {

        
            Hosturl string  `json:"hosturl"`
            Username string  `json:"username"`
            Password string  `json:"password"`
            RequestParams map[string]string  `json:"request_params"`
            HttpMethod string  `json:"http_method"`
            RequestPayload string  `json:"request_payload"`
            LocalPath string  `json:"local_path"`
            ArchiveConfig ArchiveConfig  `json:"archive_config"`
         
    }
    
    // JobConfig used by Inventory
    type JobConfig struct {

        
            ID float64  `json:"_id"`
            JobCode string  `json:"job_code"`
            TaskType string  `json:"task_type"`
            SyncDelay float64  `json:"sync_delay"`
            CronExpression string  `json:"cron_expression"`
            StoreFilter StoreFilter  `json:"store_filter"`
            ProcessConfig ProcessConfig  `json:"process_config"`
            StoreConfig []StoreConfig  `json:"store_config"`
            Properties map[string]string  `json:"properties"`
            ImmediateFirstRun bool  `json:"immediate_first_run"`
            Disable bool  `json:"disable"`
            DependentJobCodes []string  `json:"dependent_job_codes"`
            CompanyConfig []CompanyConfig  `json:"company_config"`
            CompanyIds []float64  `json:"company_ids"`
            TaxIdentifiers []string  `json:"tax_identifiers"`
            Priority string  `json:"priority"`
            PeriodThreshold float64  `json:"period_threshold"`
            PeriodThresholdType string  `json:"period_threshold_type"`
            DbConnectionProfile DBConnectionProfile  `json:"db_connection_profile"`
            Params map[string]map[string]interface{}  `json:"params"`
            OpenTags map[string]map[string]interface{}  `json:"open_tags"`
            DeleteQuantityThreshold float64  `json:"delete_quantity_threshold"`
            CatalogMasterConfig CatalogMasterConfig  `json:"catalog_master_config"`
            AggregatorTypes []string  `json:"aggregator_types"`
            IntegrationType string  `json:"integration_type"`
            MinPrice float64  `json:"min_price"`
            Audit Audit  `json:"audit"`
            Version float64  `json:"version"`
            Alias string  `json:"alias"`
         
    }
    
    // JobConfigRawDTO used by Inventory
    type JobConfigRawDTO struct {

        
            CompanyName string  `json:"company_name"`
            Integration string  `json:"integration"`
            CompanyID float64  `json:"company_id"`
            Data JobConfig  `json:"data"`
         
    }
    
    // JsonDocConfig used by Inventory
    type JsonDocConfig struct {

        
            PropBeanConfigs []PropBeanConfig  `json:"prop_bean_configs"`
         
    }
    
    // LocalFileConfig used by Inventory
    type LocalFileConfig struct {

        
            Retries float64  `json:"retries"`
            Interval float64  `json:"interval"`
            LocalDir string  `json:"local_dir"`
            WorkingDir string  `json:"working_dir"`
            Unzip bool  `json:"unzip"`
            ZipFileRegex string  `json:"zip_file_regex"`
            FileRegex string  `json:"file_regex"`
            ZipFormat string  `json:"zip_format"`
            ArchiveConfig ArchiveConfig  `json:"archive_config"`
            ReadAllFiles bool  `json:"read_all_files"`
         
    }
    
    // MongoDocConfig used by Inventory
    type MongoDocConfig struct {

        
            CollectionName string  `json:"collection_name"`
            FindQuery map[string]map[string]interface{}  `json:"find_query"`
            ProjectionQuery map[string]map[string]interface{}  `json:"projection_query"`
            PropBeanConfigs []PropBeanConfig  `json:"prop_bean_configs"`
            AggregatePipeline []map[string]map[string]interface{}  `json:"aggregate_pipeline"`
            SkipSave bool  `json:"skip_save"`
         
    }
    
    // OAuthConfig used by Inventory
    type OAuthConfig struct {

        
            Limit float64  `json:"limit"`
            Pages float64  `json:"pages"`
            Interval float64  `json:"interval"`
            ConsumerKey string  `json:"consumer_key"`
            ConsumerSecret string  `json:"consumer_secret"`
            Token string  `json:"token"`
            TokenSecret string  `json:"token_secret"`
            RestURL string  `json:"rest_url"`
            RestBaseURL string  `json:"rest_base_url"`
            FunctionName string  `json:"function_name"`
         
    }
    
    // ProcessConfig used by Inventory
    type ProcessConfig struct {

        
            DbConfig DBConfig  `json:"db_config"`
            DbParamConfig DBParamConfig  `json:"db_param_config"`
            SftpConfig SFTPConfig  `json:"sftp_config"`
            AwsS3Config AWSS3config  `json:"aws_s3_config"`
            MongoDocConfig MongoDocConfig  `json:"mongo_doc_config"`
            FtpConfig FTPConfig  `json:"ftp_config"`
            EmailConfig EmailConfig  `json:"email_config"`
            FileConfig FileConfig  `json:"file_config"`
            JsonDocConfig JsonDocConfig  `json:"json_doc_config"`
            DocMappingConfig DocMappingConfig  `json:"doc_mapping_config"`
            TaskStepConfig TaskStepConfig  `json:"task_step_config"`
            HttpConfig HttpConfig  `json:"http_config"`
            LocalFileConfig LocalFileConfig  `json:"local_file_config"`
            OauthConfig OAuthConfig  `json:"oauth_config"`
            GoogleSpreadsheetConfig GoogleSpreadSheetConfig  `json:"google_spreadsheet_config"`
         
    }
    
    // PropBeanConfig used by Inventory
    type PropBeanConfig struct {

        
            Required bool  `json:"required"`
            Optional bool  `json:"optional"`
            Send Send  `json:"send"`
            Validations []map[string]map[string]interface{}  `json:"validations"`
            Values []string  `json:"values"`
            Include bool  `json:"include"`
            SourceField string  `json:"source_field"`
            SourceFields []string  `json:"source_fields"`
            DestinationField string  `json:"destination_field"`
            DataType string  `json:"data_type"`
            DefaultValue map[string]interface{}  `json:"default_value"`
            ConstValue map[string]interface{}  `json:"const_value"`
            ConcatStr string  `json:"concat_str"`
            FunctionName string  `json:"function_name"`
            TransformerName string  `json:"transformer_name"`
            AllParamFunctionName string  `json:"all_param_function_name"`
            SubSeparator string  `json:"sub_separator"`
            IndexField string  `json:"index_field"`
            IgnoreIfNotExists bool  `json:"ignore_if_not_exists"`
            IdentifierType string  `json:"identifier_type"`
            ProjectionQuery map[string]map[string]interface{}  `json:"projection_query"`
            EnrichFromMaster bool  `json:"enrich_from_master"`
         
    }
    
    // PropBeanDTO used by Inventory
    type PropBeanDTO struct {

        
            Required bool  `json:"required"`
            Optional bool  `json:"optional"`
            Include bool  `json:"include"`
            SourceField string  `json:"source_field"`
            SourceFields []string  `json:"source_fields"`
            DestinationField string  `json:"destination_field"`
            DataType string  `json:"data_type"`
            DefaultValue map[string]interface{}  `json:"default_value"`
            ConstValue map[string]interface{}  `json:"const_value"`
            ConcatStr string  `json:"concat_str"`
            FunctionName string  `json:"function_name"`
            TransformerName string  `json:"transformer_name"`
            AllParamFunctionName string  `json:"all_param_function_name"`
            SubSeparator string  `json:"sub_separator"`
            IndexField string  `json:"index_field"`
            IgnoreIfNotExists bool  `json:"ignore_if_not_exists"`
            IdentifierType string  `json:"identifier_type"`
            ProjectionQuery map[string]map[string]interface{}  `json:"projection_query"`
            EnrichFromMaster bool  `json:"enrich_from_master"`
         
    }
    
    // ResponseEnvelopeListJobConfigRawDTO used by Inventory
    type ResponseEnvelopeListJobConfigRawDTO struct {

        
            Timestamp string  `json:"timestamp"`
            Status float64  `json:"status"`
            Error string  `json:"error"`
            Exception string  `json:"exception"`
            Message string  `json:"message"`
            TotalTimeTakenInMillis float64  `json:"total_time_taken_in_millis"`
            HttpStatus string  `json:"http_status"`
            Items []JobConfigRawDTO  `json:"items"`
            Payload []JobConfigRawDTO  `json:"payload"`
            TraceID string  `json:"trace_id"`
            Page Page  `json:"page"`
         
    }
    
    // SFTPConfig used by Inventory
    type SFTPConfig struct {

        
            Host string  `json:"host"`
            Port float64  `json:"port"`
            Username string  `json:"username"`
            Password string  `json:"password"`
            Unzip bool  `json:"unzip"`
            Retries float64  `json:"retries"`
            Interval float64  `json:"interval"`
            PrivateKeyPath string  `json:"private_key_path"`
            StrictHostKeyChecking bool  `json:"strict_host_key_checking"`
            LocalDir string  `json:"local_dir"`
            RemoteDir string  `json:"remote_dir"`
            PasswordProtected bool  `json:"password_protected"`
            ZipFileRegex string  `json:"zip_file_regex"`
            FileRegex string  `json:"file_regex"`
            ZipFormat string  `json:"zip_format"`
            ArchiveConfig ArchiveConfig  `json:"archive_config"`
            ReadAllFiles bool  `json:"read_all_files"`
         
    }
    
    // Send used by Inventory
    type Send struct {

        
            Raw bool  `json:"raw"`
            Processed bool  `json:"processed"`
         
    }
    
    // StoreConfig used by Inventory
    type StoreConfig struct {

        
            JobCode string  `json:"job_code"`
            Storeid string  `json:"storeid"`
            StoreAlias string  `json:"store_alias"`
            StoreFileRegex string  `json:"store_file_regex"`
            StoreFileName string  `json:"store_file_name"`
            ProcessConfig ProcessConfig  `json:"process_config"`
            Properties map[string]string  `json:"properties"`
         
    }
    
    // StoreFilter used by Inventory
    type StoreFilter struct {

        
            IncludeTags []string  `json:"include_tags"`
            ExcludeTags []string  `json:"exclude_tags"`
            Query map[string]map[string]interface{}  `json:"query"`
         
    }
    
    // TaskConfig used by Inventory
    type TaskConfig struct {

        
            Name string  `json:"name"`
            TaskConfigID float64  `json:"task_config_id"`
            TaskParams []TaskParam  `json:"task_params"`
         
    }
    
    // TaskParam used by Inventory
    type TaskParam struct {

        
            Name string  `json:"name"`
            Value map[string]interface{}  `json:"value"`
         
    }
    
    // TaskStepConfig used by Inventory
    type TaskStepConfig struct {

        
            TaskConfigs []TaskConfig  `json:"task_configs"`
            TaskConfigIds []float64  `json:"task_config_ids"`
            TaskConfigGroupIds []float64  `json:"task_config_group_ids"`
         
    }
    
    // JobStepsDTO used by Inventory
    type JobStepsDTO struct {

        
            StepName string  `json:"step_name"`
            Type string  `json:"type"`
            StepExecutionTime float64  `json:"step_execution_time"`
            StartCount float64  `json:"start_count"`
            EndCount float64  `json:"end_count"`
            DeletedCount float64  `json:"deleted_count"`
            ProcessedStartTime string  `json:"processed_start_time"`
            ProcessedAt string  `json:"processed_at"`
         
    }
    
    // ResponseEnvelopeListJobStepsDTO used by Inventory
    type ResponseEnvelopeListJobStepsDTO struct {

        
            Timestamp string  `json:"timestamp"`
            Status float64  `json:"status"`
            Error string  `json:"error"`
            Exception string  `json:"exception"`
            Message string  `json:"message"`
            TotalTimeTakenInMillis float64  `json:"total_time_taken_in_millis"`
            HttpStatus string  `json:"http_status"`
            Items []JobStepsDTO  `json:"items"`
            Payload []JobStepsDTO  `json:"payload"`
            TraceID string  `json:"trace_id"`
            Page Page  `json:"page"`
         
    }
    
    // ResponseEnvelopeListJobConfigDTO used by Inventory
    type ResponseEnvelopeListJobConfigDTO struct {

        
            Timestamp string  `json:"timestamp"`
            Status float64  `json:"status"`
            Error string  `json:"error"`
            Exception string  `json:"exception"`
            Message string  `json:"message"`
            TotalTimeTakenInMillis float64  `json:"total_time_taken_in_millis"`
            HttpStatus string  `json:"http_status"`
            Items []JobConfigDTO  `json:"items"`
            Payload []JobConfigDTO  `json:"payload"`
            TraceID string  `json:"trace_id"`
            Page Page  `json:"page"`
         
    }
    
    // ResponseEnvelopeJobConfigDTO used by Inventory
    type ResponseEnvelopeJobConfigDTO struct {

        
            Timestamp string  `json:"timestamp"`
            Status float64  `json:"status"`
            Error string  `json:"error"`
            Exception string  `json:"exception"`
            Message string  `json:"message"`
            TotalTimeTakenInMillis float64  `json:"total_time_taken_in_millis"`
            HttpStatus string  `json:"http_status"`
            Items JobConfigDTO  `json:"items"`
            Payload JobConfigDTO  `json:"payload"`
            TraceID string  `json:"trace_id"`
            Page Page  `json:"page"`
         
    }
    
    // JobHistoryDto used by Inventory
    type JobHistoryDto struct {

        
            TotalAddedCount float64  `json:"total_added_count"`
            TotalUpdatedCount float64  `json:"total_updated_count"`
            TotalSuppressedCount float64  `json:"total_suppressed_count"`
            TotalInitialCount float64  `json:"total_initial_count"`
            JobID float64  `json:"job_id"`
            Status string  `json:"status"`
            JobCode string  `json:"job_code"`
            ProcessedOn string  `json:"processed_on"`
            Filename []string  `json:"filename"`
            ErrorType string  `json:"error_type"`
            Message string  `json:"message"`
         
    }
    
    // JobMetricsDto used by Inventory
    type JobMetricsDto struct {

        
            JobCode string  `json:"job_code"`
            IsRunMoreThanUsual string  `json:"is_run_more_than_usual"`
            JobHistory []JobHistoryDto  `json:"job_history"`
            TotalSuccessCount float64  `json:"total_success_count"`
            TotalFailureCount float64  `json:"total_failure_count"`
            TotalWarningCount float64  `json:"total_warning_count"`
            TotalSuppressedCount float64  `json:"total_suppressed_count"`
            TotalJobRuns float64  `json:"total_job_runs"`
         
    }
    
    // ResponseEnvelopeJobMetricsDto used by Inventory
    type ResponseEnvelopeJobMetricsDto struct {

        
            Timestamp string  `json:"timestamp"`
            Status float64  `json:"status"`
            Error string  `json:"error"`
            Exception string  `json:"exception"`
            Message string  `json:"message"`
            TotalTimeTakenInMillis float64  `json:"total_time_taken_in_millis"`
            HttpStatus string  `json:"http_status"`
            Items JobMetricsDto  `json:"items"`
            Payload JobMetricsDto  `json:"payload"`
            TraceID string  `json:"trace_id"`
            Page Page  `json:"page"`
         
    }
    
    // JobConfigListDTO used by Inventory
    type JobConfigListDTO struct {

        
            Code string  `json:"code"`
            Alias string  `json:"alias"`
            ModifiedBy string  `json:"modified_by"`
            CreatedBy string  `json:"created_by"`
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
            Active bool  `json:"active"`
            Type string  `json:"type"`
         
    }
    
    // ResponseEnvelopeListJobConfigListDTO used by Inventory
    type ResponseEnvelopeListJobConfigListDTO struct {

        
            Timestamp string  `json:"timestamp"`
            Status float64  `json:"status"`
            Error string  `json:"error"`
            Exception string  `json:"exception"`
            Message string  `json:"message"`
            TotalTimeTakenInMillis float64  `json:"total_time_taken_in_millis"`
            HttpStatus string  `json:"http_status"`
            Items []JobConfigListDTO  `json:"items"`
            Payload []JobConfigListDTO  `json:"payload"`
            TraceID string  `json:"trace_id"`
            Page Page  `json:"page"`
         
    }
    

    
    // ApplicationInventory used by Configuration
    type ApplicationInventory struct {

        
            Inventory AppInventoryConfig  `json:"inventory"`
            Authentication AuthenticationConfig  `json:"authentication"`
            ArticleAssignment ArticleAssignmentConfig  `json:"article_assignment"`
            RewardPoints RewardPointsConfig  `json:"reward_points"`
            Cart AppCartConfig  `json:"cart"`
            Payment AppPaymentConfig  `json:"payment"`
            Order AppOrderConfig  `json:"order"`
            Logistics AppLogisticsConfig  `json:"logistics"`
            Business string  `json:"business"`
            CommsEnabled bool  `json:"comms_enabled"`
            Platforms []string  `json:"platforms"`
            ID string  `json:"_id"`
            LoyaltyPoints LoyaltyPointsConfig  `json:"loyalty_points"`
            App string  `json:"app"`
            CreatedAt string  `json:"created_at"`
            UpdatedAt string  `json:"updated_at"`
            ModifiedBy string  `json:"modified_by"`
         
    }
    
    // AppInventoryConfig used by Configuration
    type AppInventoryConfig struct {

        
            Brand InventoryBrand  `json:"brand"`
            Store InventoryStore  `json:"store"`
            Category InventoryCategory  `json:"category"`
            Price InventoryPrice  `json:"price"`
            Discount InventoryDiscount  `json:"discount"`
            OutOfStock bool  `json:"out_of_stock"`
            OnlyVerifiedProducts bool  `json:"only_verified_products"`
            FranchiseEnabled bool  `json:"franchise_enabled"`
            ExcludeCategory []interface{}  `json:"exclude_category"`
            Image []string  `json:"image"`
            CompanyStore []interface{}  `json:"company_store"`
         
    }
    
    // InventoryBrand used by Configuration
    type InventoryBrand struct {

        
            Criteria string  `json:"criteria"`
            Brands []interface{}  `json:"brands"`
         
    }
    
    // InventoryStore used by Configuration
    type InventoryStore struct {

        
            Criteria string  `json:"criteria"`
            Stores []interface{}  `json:"stores"`
            Rules AppStoreRules  `json:"rules"`
         
    }
    
    // AppStoreRules used by Configuration
    type AppStoreRules struct {

        
            Companies []float64  `json:"companies"`
            Brands []interface{}  `json:"brands"`
         
    }
    
    // InventoryCategory used by Configuration
    type InventoryCategory struct {

        
            Criteria string  `json:"criteria"`
            Categories []interface{}  `json:"categories"`
         
    }
    
    // InventoryPrice used by Configuration
    type InventoryPrice struct {

        
            Min float64  `json:"min"`
            Max float64  `json:"max"`
         
    }
    
    // InventoryDiscount used by Configuration
    type InventoryDiscount struct {

        
            Min float64  `json:"min"`
            Max float64  `json:"max"`
         
    }
    
    // AuthenticationConfig used by Configuration
    type AuthenticationConfig struct {

        
            Required bool  `json:"required"`
            Provider string  `json:"provider"`
         
    }
    
    // ArticleAssignmentConfig used by Configuration
    type ArticleAssignmentConfig struct {

        
            Rules ArticleAssignmentRules  `json:"rules"`
            PostOrderReassignment bool  `json:"post_order_reassignment"`
         
    }
    
    // ArticleAssignmentRules used by Configuration
    type ArticleAssignmentRules struct {

        
            StorePriority StorePriority  `json:"store_priority"`
         
    }
    
    // StorePriority used by Configuration
    type StorePriority struct {

        
            Enabled bool  `json:"enabled"`
            StoretypeOrder []interface{}  `json:"storetype_order"`
         
    }
    
    // AppCartConfig used by Configuration
    type AppCartConfig struct {

        
            DeliveryCharges DeliveryCharges  `json:"delivery_charges"`
            Enabled bool  `json:"enabled"`
            MaxCartItems float64  `json:"max_cart_items"`
            MinCartValue float64  `json:"min_cart_value"`
            BulkCoupons bool  `json:"bulk_coupons"`
            RevenueEngineCoupon bool  `json:"revenue_engine_coupon"`
            EmptyCart bool  `json:"empty_cart"`
            PanCard PanCardConfig  `json:"pan_card"`
         
    }
    
    // DeliveryCharges used by Configuration
    type DeliveryCharges struct {

        
            Enabled bool  `json:"enabled"`
            Charges Charges  `json:"charges"`
         
    }
    
    // Charges used by Configuration
    type Charges struct {

        
            Threshold float64  `json:"threshold"`
            Charges float64  `json:"charges"`
         
    }
    
    // AppPaymentConfig used by Configuration
    type AppPaymentConfig struct {

        
            CallbackURL CallbackUrl  `json:"callback_url"`
            Methods Methods  `json:"methods"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            ModeOfPayment string  `json:"mode_of_payment"`
            Source string  `json:"source"`
            Enabled bool  `json:"enabled"`
            CodAmountLimit float64  `json:"cod_amount_limit"`
            CodCharges float64  `json:"cod_charges"`
         
    }
    
    // CallbackUrl used by Configuration
    type CallbackUrl struct {

        
            App string  `json:"app"`
            Web string  `json:"web"`
         
    }
    
    // Methods used by Configuration
    type Methods struct {

        
            Pl PaymentModeConfig  `json:"pl"`
            Card PaymentModeConfig  `json:"card"`
            Nb PaymentModeConfig  `json:"nb"`
            Wl PaymentModeConfig  `json:"wl"`
            Ps PaymentModeConfig  `json:"ps"`
            Upi PaymentModeConfig  `json:"upi"`
            Qr PaymentModeConfig  `json:"qr"`
            Cod PaymentModeConfig  `json:"cod"`
            Pp PaymentModeConfig  `json:"pp"`
            Jp PaymentModeConfig  `json:"jp"`
            Pac PaymentModeConfig  `json:"pac"`
            Fc PaymentModeConfig  `json:"fc"`
            Jiopp PaymentModeConfig  `json:"jiopp"`
            Stripepg PaymentModeConfig  `json:"stripepg"`
            Juspaypg PaymentModeConfig  `json:"juspaypg"`
            Payubizpg PaymentModeConfig  `json:"payubizpg"`
            Payumoneypg PaymentModeConfig  `json:"payumoneypg"`
            Rupifipg PaymentModeConfig  `json:"rupifipg"`
            Simpl PaymentModeConfig  `json:"simpl"`
         
    }
    
    // PaymentModeConfig used by Configuration
    type PaymentModeConfig struct {

        
            Enabled bool  `json:"enabled"`
         
    }
    
    // PaymentSelectionLock used by Configuration
    type PaymentSelectionLock struct {

        
            Enabled bool  `json:"enabled"`
            DefaultOptions string  `json:"default_options"`
            PaymentIdentifier string  `json:"payment_identifier"`
         
    }
    
    // AppOrderConfig used by Configuration
    type AppOrderConfig struct {

        
            Enabled bool  `json:"enabled"`
            ForceReassignment bool  `json:"force_reassignment"`
            Message string  `json:"message"`
         
    }
    
    // AppLogisticsConfig used by Configuration
    type AppLogisticsConfig struct {

        
            LogisticsBySeller bool  `json:"logistics_by_seller"`
            ServiceabilityCheck bool  `json:"serviceability_check"`
            SameDayDelivery bool  `json:"same_day_delivery"`
            DpAssignment bool  `json:"dp_assignment"`
         
    }
    
    // LoyaltyPointsConfig used by Configuration
    type LoyaltyPointsConfig struct {

        
            Enabled bool  `json:"enabled"`
            AutoApply bool  `json:"auto_apply"`
         
    }
    
    // AppInventoryPartialUpdate used by Configuration
    type AppInventoryPartialUpdate struct {

        
            RewardPoints RewardPointsConfig  `json:"reward_points"`
            Cart AppCartConfig  `json:"cart"`
            Payment AppPaymentConfig  `json:"payment"`
            LoyaltyPoints LoyaltyPointsConfig  `json:"loyalty_points"`
            CommsEnabled bool  `json:"comms_enabled"`
         
    }
    
    // BrandCompanyInfo used by Configuration
    type BrandCompanyInfo struct {

        
            CompanyName string  `json:"company_name"`
            CompanyID float64  `json:"company_id"`
         
    }
    
    // CompanyByBrandsRequest used by Configuration
    type CompanyByBrandsRequest struct {

        
            Brands float64  `json:"brands"`
            SearchText string  `json:"search_text"`
         
    }
    
    // CompanyByBrandsResponse used by Configuration
    type CompanyByBrandsResponse struct {

        
            Items []BrandCompanyInfo  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // StoreByBrandsRequest used by Configuration
    type StoreByBrandsRequest struct {

        
            CompanyID float64  `json:"company_id"`
            Brands float64  `json:"brands"`
            SearchText string  `json:"search_text"`
         
    }
    
    // StoreByBrandsResponse used by Configuration
    type StoreByBrandsResponse struct {

        
            Items []BrandStoreInfo  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // BrandStoreInfo used by Configuration
    type BrandStoreInfo struct {

        
            StoreName string  `json:"store_name"`
            StoreID float64  `json:"store_id"`
            StoreType string  `json:"store_type"`
            StoreCode string  `json:"store_code"`
            StoreAddress OptedStoreAddress  `json:"store_address"`
            Company OptedCompany  `json:"company"`
         
    }
    
    // CompanyBrandInfo used by Configuration
    type CompanyBrandInfo struct {

        
            Name string  `json:"name"`
            Value float64  `json:"value"`
            BrandLogoURL string  `json:"brand_logo_url"`
            BrandBannerURL string  `json:"brand_banner_url"`
            BrandBannerPortraitURL string  `json:"brand_banner_portrait_url"`
         
    }
    
    // BrandsByCompanyResponse used by Configuration
    type BrandsByCompanyResponse struct {

        
            Brands CompanyBrandInfo  `json:"brands"`
         
    }
    
    // PanCardConfig used by Configuration
    type PanCardConfig struct {

        
            Enabled bool  `json:"enabled"`
            ThresholdAmount float64  `json:"threshold_amount"`
         
    }
    
    // CreateApplicationRequest used by Configuration
    type CreateApplicationRequest struct {

        
            App App  `json:"app"`
            Configuration AppInventory  `json:"configuration"`
            Domain AppDomain  `json:"domain"`
         
    }
    
    // CreateAppResponse used by Configuration
    type CreateAppResponse struct {

        
            App Application  `json:"app"`
            Configuration ApplicationInventory  `json:"configuration"`
         
    }
    
    // ApplicationsResponse used by Configuration
    type ApplicationsResponse struct {

        
            Items []Application  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // MobileAppConfiguration used by Configuration
    type MobileAppConfiguration struct {

        
            IsActive bool  `json:"is_active"`
            ID string  `json:"_id"`
            AppName string  `json:"app_name"`
            LandingImage LandingImage  `json:"landing_image"`
            SplashImage SplashImage  `json:"splash_image"`
            Application string  `json:"application"`
            PlatformType string  `json:"platform_type"`
            CreatedAt string  `json:"created_at"`
            UpdatedAt string  `json:"updated_at"`
            V float64  `json:"__v"`
            PackageName string  `json:"package_name"`
         
    }
    
    // LandingImage used by Configuration
    type LandingImage struct {

        
            AspectRatio string  `json:"aspect_ratio"`
            SecureURL string  `json:"secure_url"`
         
    }
    
    // SplashImage used by Configuration
    type SplashImage struct {

        
            AspectRatio string  `json:"aspect_ratio"`
            SecureURL string  `json:"secure_url"`
         
    }
    
    // MobileAppConfigRequest used by Configuration
    type MobileAppConfigRequest struct {

        
            AppName string  `json:"app_name"`
            LandingImage LandingImage  `json:"landing_image"`
            SplashImage SplashImage  `json:"splash_image"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // BuildVersionHistory used by Configuration
    type BuildVersionHistory struct {

        
            Versions BuildVersion  `json:"versions"`
            LatestAvailableVersionName string  `json:"latest_available_version_name"`
         
    }
    
    // BuildVersion used by Configuration
    type BuildVersion struct {

        
            ID string  `json:"_id"`
            Application string  `json:"application"`
            PlatformType string  `json:"platform_type"`
            BuildStatus string  `json:"build_status"`
            VersionName string  `json:"version_name"`
            VersionCode float64  `json:"version_code"`
            CreatedAt string  `json:"created_at"`
            UpdatedAt string  `json:"updated_at"`
            V float64  `json:"__v"`
         
    }
    
    // AppSupportedCurrency used by Configuration
    type AppSupportedCurrency struct {

        
            ID string  `json:"_id"`
            SupportedCurrency []string  `json:"supported_currency"`
            Application string  `json:"application"`
            DefaultCurrency DefaultCurrency  `json:"default_currency"`
            CreatedAt string  `json:"created_at"`
            UpdatedAt string  `json:"updated_at"`
         
    }
    
    // DefaultCurrency used by Configuration
    type DefaultCurrency struct {

        
            Ref string  `json:"ref"`
            Code string  `json:"code"`
         
    }
    
    // CurrencyConfig used by Configuration
    type CurrencyConfig struct {

        
            ID string  `json:"_id"`
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            Code string  `json:"code"`
            CreatedAt string  `json:"created_at"`
            UpdatedAt string  `json:"updated_at"`
            DecimalDigits float64  `json:"decimal_digits"`
            Symbol string  `json:"symbol"`
         
    }
    
    // DomainAdd used by Configuration
    type DomainAdd struct {

        
            Name string  `json:"name"`
         
    }
    
    // DomainAddRequest used by Configuration
    type DomainAddRequest struct {

        
            Domain DomainAdd  `json:"domain"`
         
    }
    
    // DomainsResponse used by Configuration
    type DomainsResponse struct {

        
            Domains []Domain  `json:"domains"`
         
    }
    
    // UpdateDomain used by Configuration
    type UpdateDomain struct {

        
            ID string  `json:"_id"`
         
    }
    
    // UpdateDomainTypeRequest used by Configuration
    type UpdateDomainTypeRequest struct {

        
            Domain UpdateDomain  `json:"domain"`
            Action string  `json:"action"`
         
    }
    
    // DomainStatusRequest used by Configuration
    type DomainStatusRequest struct {

        
            DomainURL string  `json:"domain_url"`
         
    }
    
    // DomainStatus used by Configuration
    type DomainStatus struct {

        
            Display string  `json:"display"`
            Status bool  `json:"status"`
         
    }
    
    // DomainStatusResponse used by Configuration
    type DomainStatusResponse struct {

        
            Connected bool  `json:"connected"`
            Status []DomainStatus  `json:"status"`
         
    }
    
    // DomainSuggestionsRequest used by Configuration
    type DomainSuggestionsRequest struct {

        
            DomainURL string  `json:"domain_url"`
            Custom bool  `json:"custom"`
         
    }
    
    // DomainSuggestion used by Configuration
    type DomainSuggestion struct {

        
            Name string  `json:"name"`
            Unsupported bool  `json:"unsupported"`
            IsAvailable bool  `json:"is_available"`
            Price float64  `json:"price"`
            Currency string  `json:"currency"`
         
    }
    
    // DomainSuggestionsResponse used by Configuration
    type DomainSuggestionsResponse struct {

        
            Domains []DomainSuggestion  `json:"domains"`
         
    }
    
    // GetIntegrationsOptInsResponse used by Configuration
    type GetIntegrationsOptInsResponse struct {

        
            Items []IntegrationOptIn  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // IntegrationOptIn used by Configuration
    type IntegrationOptIn struct {

        
            Validators Validators  `json:"validators"`
            Description string  `json:"description"`
            DescriptionHtml string  `json:"description_html"`
            Constants string  `json:"constants"`
            Companies []map[string]interface{}  `json:"companies"`
            Support []string  `json:"support"`
            ID string  `json:"_id"`
            Name string  `json:"name"`
            Meta []IntegrationMeta  `json:"meta"`
            Icon string  `json:"icon"`
            Owner string  `json:"owner"`
            CreatedAt string  `json:"created_at"`
            UpdatedAt string  `json:"updated_at"`
            Token string  `json:"token"`
            Secret string  `json:"secret"`
            V float64  `json:"__v"`
         
    }
    
    // Validators used by Configuration
    type Validators struct {

        
            Company CompanyValidator  `json:"company"`
            Store StoreValidator  `json:"store"`
            Inventory InventoryValidator  `json:"inventory"`
            Order OrderValidator  `json:"order"`
         
    }
    
    // CompanyValidator used by Configuration
    type CompanyValidator struct {

        
            JsonSchema []JsonSchema  `json:"json_schema"`
            BrowserScript string  `json:"browser_script"`
         
    }
    
    // JsonSchema used by Configuration
    type JsonSchema struct {

        
            Display string  `json:"display"`
            Key string  `json:"key"`
            Type string  `json:"type"`
            Tooltip string  `json:"tooltip"`
         
    }
    
    // StoreValidator used by Configuration
    type StoreValidator struct {

        
            JsonSchema []JsonSchema  `json:"json_schema"`
            BrowserScript string  `json:"browser_script"`
         
    }
    
    // InventoryValidator used by Configuration
    type InventoryValidator struct {

        
            JsonSchema []JsonSchema  `json:"json_schema"`
            BrowserScript string  `json:"browser_script"`
         
    }
    
    // OrderValidator used by Configuration
    type OrderValidator struct {

        
            JsonSchema []JsonSchema  `json:"json_schema"`
            BrowserScript string  `json:"browser_script"`
         
    }
    
    // IntegrationMeta used by Configuration
    type IntegrationMeta struct {

        
            IsPublic bool  `json:"is_public"`
            ID string  `json:"_id"`
            Name string  `json:"name"`
            Value string  `json:"value"`
         
    }
    
    // Integration used by Configuration
    type Integration struct {

        
            Validators Validators  `json:"validators"`
            Description string  `json:"description"`
            DescriptionHtml string  `json:"description_html"`
            Constants map[string]interface{}  `json:"constants"`
            Companies []map[string]interface{}  `json:"companies"`
            Support []string  `json:"support"`
            ID string  `json:"_id"`
            Name string  `json:"name"`
            Meta []IntegrationMeta  `json:"meta"`
            Icon string  `json:"icon"`
            Owner string  `json:"owner"`
            CreatedAt string  `json:"created_at"`
            UpdatedAt string  `json:"updated_at"`
            Token string  `json:"token"`
            Secret string  `json:"secret"`
            V float64  `json:"__v"`
         
    }
    
    // IntegrationConfigResponse used by Configuration
    type IntegrationConfigResponse struct {

        
            Items []IntegrationLevel  `json:"items"`
         
    }
    
    // IntegrationLevel used by Configuration
    type IntegrationLevel struct {

        
            Opted bool  `json:"opted"`
            Permissions []map[string]interface{}  `json:"permissions"`
            LastPatch []LastPatch  `json:"last_patch"`
            ID string  `json:"_id"`
            Integration string  `json:"integration"`
            Level string  `json:"level"`
            UID float64  `json:"uid"`
            Meta []IntegrationMeta  `json:"meta"`
            Token string  `json:"token"`
            CreatedAt string  `json:"created_at"`
            UpdatedAt string  `json:"updated_at"`
            V float64  `json:"__v"`
            Data map[string]interface{}  `json:"data"`
         
    }
    
    // UpdateIntegrationLevelRequest used by Configuration
    type UpdateIntegrationLevelRequest struct {

        
            Items []IntegrationLevel  `json:"items"`
         
    }
    
    // OptedStoreIntegration used by Configuration
    type OptedStoreIntegration struct {

        
            OtherOpted bool  `json:"other_opted"`
            OtherIntegration IntegrationOptIn  `json:"other_integration"`
            OtherEntity OtherEntity  `json:"other_entity"`
         
    }
    
    // OtherEntity used by Configuration
    type OtherEntity struct {

        
            Opted bool  `json:"opted"`
            Permissions []string  `json:"permissions"`
            LastPatch []LastPatch  `json:"last_patch"`
            ID string  `json:"_id"`
            Integration string  `json:"integration"`
            Level string  `json:"level"`
            UID float64  `json:"uid"`
            Data OtherEntityData  `json:"data"`
            Meta []map[string]interface{}  `json:"meta"`
            Token string  `json:"token"`
            CreatedAt string  `json:"created_at"`
            UpdatedAt string  `json:"updated_at"`
            V float64  `json:"__v"`
         
    }
    
    // LastPatch used by Configuration
    type LastPatch struct {

        
            Op string  `json:"op"`
            Path string  `json:"path"`
            Value string  `json:"value"`
         
    }
    
    // OtherEntityData used by Configuration
    type OtherEntityData struct {

        
            ArticleIdentifier string  `json:"article_identifier"`
         
    }
    
    // App used by Configuration
    type App struct {

        
            CompanyID string  `json:"company_id"`
            ChannelType string  `json:"channel_type"`
            Auth ApplicationAuth  `json:"auth"`
            Name string  `json:"name"`
            Desc string  `json:"desc"`
         
    }
    
    // AppInventory used by Configuration
    type AppInventory struct {

        
            Brand InventoryBrandRule  `json:"brand"`
            Store InventoryStoreRule  `json:"store"`
            Image []string  `json:"image"`
            FranchiseEnabled bool  `json:"franchise_enabled"`
            OutOfStock bool  `json:"out_of_stock"`
            OnlyVerifiedProducts bool  `json:"only_verified_products"`
            Payment InventoryPaymentConfig  `json:"payment"`
            ArticleAssignment InventoryArticleAssignment  `json:"article_assignment"`
         
    }
    
    // AppDomain used by Configuration
    type AppDomain struct {

        
            Name string  `json:"name"`
         
    }
    
    // CompaniesResponse used by Configuration
    type CompaniesResponse struct {

        
            Items AppInventoryCompanies  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // AppInventoryCompanies used by Configuration
    type AppInventoryCompanies struct {

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            CompanyType string  `json:"company_type"`
         
    }
    
    // StoresResponse used by Configuration
    type StoresResponse struct {

        
            Items AppInventoryStores  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // AppInventoryStores used by Configuration
    type AppInventoryStores struct {

        
            ID string  `json:"_id"`
            ModifiedOn string  `json:"modified_on"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            DisplayName string  `json:"display_name"`
            StoreType string  `json:"store_type"`
            StoreCode string  `json:"store_code"`
            CompanyID float64  `json:"company_id"`
         
    }
    
    // FilterOrderingStoreRequest used by Configuration
    type FilterOrderingStoreRequest struct {

        
            AllStores bool  `json:"all_stores"`
            DeployedStores []float64  `json:"deployed_stores"`
            Q string  `json:"q"`
            OnlyDeployed bool  `json:"only_deployed"`
         
    }
    
    // DeploymentMeta used by Configuration
    type DeploymentMeta struct {

        
            DeployedStores []float64  `json:"deployed_stores"`
            AllStores bool  `json:"all_stores"`
            Enabled bool  `json:"enabled"`
            Type string  `json:"type"`
            ID string  `json:"_id"`
            App string  `json:"app"`
         
    }
    
    // OrderingStoreConfig used by Configuration
    type OrderingStoreConfig struct {

        
            DeploymentMeta DeploymentMeta  `json:"deployment_meta"`
         
    }
    
    // OtherSellerCompany used by Configuration
    type OtherSellerCompany struct {

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // OtherSellerApplication used by Configuration
    type OtherSellerApplication struct {

        
            Name string  `json:"name"`
            Description string  `json:"description"`
            ID string  `json:"_id"`
            Domain string  `json:"domain"`
            Company OtherSellerCompany  `json:"company"`
            OptType string  `json:"opt_type"`
         
    }
    
    // OtherSellerApplications used by Configuration
    type OtherSellerApplications struct {

        
            Items []OtherSellerApplication  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // OptedApplicationResponse used by Configuration
    type OptedApplicationResponse struct {

        
            Name string  `json:"name"`
            Description string  `json:"description"`
            ID string  `json:"_id"`
            Domain string  `json:"domain"`
            Company OptedCompany  `json:"company"`
            OptedInventory OptedInventory  `json:"opted_inventory"`
            OptOutInventory OptOutInventory  `json:"opt_out_inventory"`
         
    }
    
    // OptedCompany used by Configuration
    type OptedCompany struct {

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // OptedInventory used by Configuration
    type OptedInventory struct {

        
            OptType OptType  `json:"opt_type"`
            Items interface{}  `json:"items"`
         
    }
    
    // OptType used by Configuration
    type OptType struct {

        
            Key string  `json:"key"`
            Display string  `json:"display"`
         
    }
    
    // OptedStore used by Configuration
    type OptedStore struct {

        
            Name string  `json:"name"`
            StoreCode string  `json:"store_code"`
            ID string  `json:"_id"`
            ModifiedOn string  `json:"modified_on"`
            UID float64  `json:"uid"`
            Address OptedStoreAddress  `json:"address"`
            DisplayName string  `json:"display_name"`
            StoreType string  `json:"store_type"`
            CompanyID float64  `json:"company_id"`
         
    }
    
    // OptOutInventory used by Configuration
    type OptOutInventory struct {

        
            Store []float64  `json:"store"`
            Company []float64  `json:"company"`
         
    }
    
    // TokenResponse used by Configuration
    type TokenResponse struct {

        
            Tokens Tokens  `json:"tokens"`
            ID string  `json:"_id"`
            Application string  `json:"application"`
            CreatedAt string  `json:"created_at"`
            UpdatedAt string  `json:"updated_at"`
            V float64  `json:"__v"`
         
    }
    
    // Tokens used by Configuration
    type Tokens struct {

        
            Firebase Firebase  `json:"firebase"`
            Moengage Moengage  `json:"moengage"`
            Segment Segment  `json:"segment"`
            Gtm Gtm  `json:"gtm"`
            Freshchat Freshchat  `json:"freshchat"`
            Safetynet Safetynet  `json:"safetynet"`
            FyndRewards FyndRewards  `json:"fynd_rewards"`
            GoogleMap GoogleMap  `json:"google_map"`
         
    }
    
    // Firebase used by Configuration
    type Firebase struct {

        
            Credentials Credentials  `json:"credentials"`
            Enabled bool  `json:"enabled"`
         
    }
    
    // Credentials used by Configuration
    type Credentials struct {

        
            Ios Ios  `json:"ios"`
            Android Android  `json:"android"`
            ProjectID string  `json:"project_id"`
            GcmSenderID string  `json:"gcm_sender_id"`
            ApplicationID string  `json:"application_id"`
            APIKey string  `json:"api_key"`
         
    }
    
    // Ios used by Configuration
    type Ios struct {

        
            ApplicationID string  `json:"application_id"`
            APIKey string  `json:"api_key"`
         
    }
    
    // Android used by Configuration
    type Android struct {

        
            ApplicationID string  `json:"application_id"`
            APIKey string  `json:"api_key"`
         
    }
    
    // Moengage used by Configuration
    type Moengage struct {

        
            Credentials MoengageCredentials  `json:"credentials"`
            Enabled bool  `json:"enabled"`
         
    }
    
    // MoengageCredentials used by Configuration
    type MoengageCredentials struct {

        
            AppID string  `json:"app_id"`
         
    }
    
    // Segment used by Configuration
    type Segment struct {

        
            Credentials SegmentCredentials  `json:"credentials"`
            Enabled bool  `json:"enabled"`
         
    }
    
    // SegmentCredentials used by Configuration
    type SegmentCredentials struct {

        
            WriteKey string  `json:"write_key"`
         
    }
    
    // Gtm used by Configuration
    type Gtm struct {

        
            Credentials GtmCredentials  `json:"credentials"`
            Enabled bool  `json:"enabled"`
         
    }
    
    // GtmCredentials used by Configuration
    type GtmCredentials struct {

        
            APIKey string  `json:"api_key"`
         
    }
    
    // Freshchat used by Configuration
    type Freshchat struct {

        
            Credentials FreshchatCredentials  `json:"credentials"`
            Enabled bool  `json:"enabled"`
         
    }
    
    // FreshchatCredentials used by Configuration
    type FreshchatCredentials struct {

        
            AppID string  `json:"app_id"`
            AppKey string  `json:"app_key"`
            WebToken string  `json:"web_token"`
         
    }
    
    // Safetynet used by Configuration
    type Safetynet struct {

        
            Credentials SafetynetCredentials  `json:"credentials"`
            Enabled bool  `json:"enabled"`
         
    }
    
    // SafetynetCredentials used by Configuration
    type SafetynetCredentials struct {

        
            APIKey string  `json:"api_key"`
         
    }
    
    // FyndRewards used by Configuration
    type FyndRewards struct {

        
            Credentials FyndRewardsCredentials  `json:"credentials"`
         
    }
    
    // FyndRewardsCredentials used by Configuration
    type FyndRewardsCredentials struct {

        
            PublicKey string  `json:"public_key"`
         
    }
    
    // GoogleMap used by Configuration
    type GoogleMap struct {

        
            Credentials GoogleMapCredentials  `json:"credentials"`
         
    }
    
    // GoogleMapCredentials used by Configuration
    type GoogleMapCredentials struct {

        
            APIKey string  `json:"api_key"`
         
    }
    
    // RewardPointsConfig used by Configuration
    type RewardPointsConfig struct {

        
            Credit Credit  `json:"credit"`
            Debit Debit  `json:"debit"`
         
    }
    
    // Credit used by Configuration
    type Credit struct {

        
            Enabled bool  `json:"enabled"`
         
    }
    
    // Debit used by Configuration
    type Debit struct {

        
            Enabled bool  `json:"enabled"`
            AutoApply bool  `json:"auto_apply"`
            StrategyChannel string  `json:"strategy_channel"`
         
    }
    
    // ProductDetailFeature used by Configuration
    type ProductDetailFeature struct {

        
            Similar []string  `json:"similar"`
            SellerSelection bool  `json:"seller_selection"`
            UpdateProductMeta bool  `json:"update_product_meta"`
            RequestProduct bool  `json:"request_product"`
         
    }
    
    // LaunchPage used by Configuration
    type LaunchPage struct {

        
            PageType string  `json:"page_type"`
            Params map[string]interface{}  `json:"params"`
            Query map[string]interface{}  `json:"query"`
         
    }
    
    // LandingPageFeature used by Configuration
    type LandingPageFeature struct {

        
            LaunchPage LaunchPage  `json:"launch_page"`
            ContinueAsGuest bool  `json:"continue_as_guest"`
            LoginBtnText string  `json:"login_btn_text"`
            ShowDomainTextbox bool  `json:"show_domain_textbox"`
            ShowRegisterBtn bool  `json:"show_register_btn"`
         
    }
    
    // RegistrationPageFeature used by Configuration
    type RegistrationPageFeature struct {

        
            AskStoreAddress bool  `json:"ask_store_address"`
         
    }
    
    // AppFeature used by Configuration
    type AppFeature struct {

        
            ProductDetail ProductDetailFeature  `json:"product_detail"`
            LandingPage LandingPageFeature  `json:"landing_page"`
            RegistrationPage RegistrationPageFeature  `json:"registration_page"`
            HomePage HomePageFeature  `json:"home_page"`
            Common CommonFeature  `json:"common"`
            Cart CartFeature  `json:"cart"`
            Qr QrFeature  `json:"qr"`
            Pcr PcrFeature  `json:"pcr"`
            Order OrderFeature  `json:"order"`
            ID string  `json:"_id"`
            App string  `json:"app"`
            CreatedAt string  `json:"created_at"`
            UpdatedAt string  `json:"updated_at"`
            V float64  `json:"__v"`
         
    }
    
    // HomePageFeature used by Configuration
    type HomePageFeature struct {

        
            OrderProcessing bool  `json:"order_processing"`
         
    }
    
    // CommonFeature used by Configuration
    type CommonFeature struct {

        
            CommunicationOptinDialog CommunicationOptinDialogFeature  `json:"communication_optin_dialog"`
            DeploymentStoreSelection DeploymentStoreSelectionFeature  `json:"deployment_store_selection"`
            ListingPrice ListingPriceFeature  `json:"listing_price"`
            Currency CurrencyFeature  `json:"currency"`
            RevenueEngine RevenueEngineFeature  `json:"revenue_engine"`
            Feedback FeedbackFeature  `json:"feedback"`
            CompareProducts CompareProductsFeature  `json:"compare_products"`
            RewardPoints RewardPointsConfig  `json:"reward_points"`
         
    }
    
    // CommunicationOptinDialogFeature used by Configuration
    type CommunicationOptinDialogFeature struct {

        
            Visibility bool  `json:"visibility"`
         
    }
    
    // DeploymentStoreSelectionFeature used by Configuration
    type DeploymentStoreSelectionFeature struct {

        
            Enabled bool  `json:"enabled"`
            Type string  `json:"type"`
         
    }
    
    // ListingPriceFeature used by Configuration
    type ListingPriceFeature struct {

        
            Value string  `json:"value"`
         
    }
    
    // CurrencyFeature used by Configuration
    type CurrencyFeature struct {

        
            Value []string  `json:"value"`
            Type string  `json:"type"`
            DefaultCurrency string  `json:"default_currency"`
         
    }
    
    // RevenueEngineFeature used by Configuration
    type RevenueEngineFeature struct {

        
            Enabled bool  `json:"enabled"`
         
    }
    
    // FeedbackFeature used by Configuration
    type FeedbackFeature struct {

        
            Enabled bool  `json:"enabled"`
         
    }
    
    // CompareProductsFeature used by Configuration
    type CompareProductsFeature struct {

        
            Enabled bool  `json:"enabled"`
         
    }
    
    // CartFeature used by Configuration
    type CartFeature struct {

        
            GstInput bool  `json:"gst_input"`
            StaffSelection bool  `json:"staff_selection"`
            PlacingForCustomer bool  `json:"placing_for_customer"`
            GoogleMap bool  `json:"google_map"`
            RevenueEngineCoupon bool  `json:"revenue_engine_coupon"`
         
    }
    
    // QrFeature used by Configuration
    type QrFeature struct {

        
            Application bool  `json:"application"`
            Products bool  `json:"products"`
            Collections bool  `json:"collections"`
         
    }
    
    // PcrFeature used by Configuration
    type PcrFeature struct {

        
            StaffSelection bool  `json:"staff_selection"`
         
    }
    
    // OrderFeature used by Configuration
    type OrderFeature struct {

        
            BuyAgain bool  `json:"buy_again"`
         
    }
    
    // AppFeatureRequest used by Configuration
    type AppFeatureRequest struct {

        
            Feature AppFeature  `json:"feature"`
         
    }
    
    // AppFeatureResponse used by Configuration
    type AppFeatureResponse struct {

        
            Feature AppFeature  `json:"feature"`
         
    }
    
    // Currency used by Configuration
    type Currency struct {

        
            ID string  `json:"_id"`
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            Code string  `json:"code"`
            CreatedAt string  `json:"created_at"`
            UpdatedAt string  `json:"updated_at"`
            DecimalDigits float64  `json:"decimal_digits"`
            Symbol string  `json:"symbol"`
         
    }
    
    // Domain used by Configuration
    type Domain struct {

        
            Verified bool  `json:"verified"`
            IsPrimary bool  `json:"is_primary"`
            IsShortlink bool  `json:"is_shortlink"`
            ID string  `json:"_id"`
            Name string  `json:"name"`
            IsPredefined bool  `json:"is_predefined"`
         
    }
    
    // ApplicationWebsite used by Configuration
    type ApplicationWebsite struct {

        
            Enabled bool  `json:"enabled"`
            Basepath string  `json:"basepath"`
         
    }
    
    // ApplicationCors used by Configuration
    type ApplicationCors struct {

        
            Domains []string  `json:"domains"`
         
    }
    
    // ApplicationAuth used by Configuration
    type ApplicationAuth struct {

        
            Enabled bool  `json:"enabled"`
         
    }
    
    // ApplicationRedirections used by Configuration
    type ApplicationRedirections struct {

        
            RedirectFrom string  `json:"redirect_from"`
            RedirectTo string  `json:"redirect_to"`
            Type string  `json:"type"`
         
    }
    
    // ApplicationMeta used by Configuration
    type ApplicationMeta struct {

        
            Name string  `json:"name"`
            Value string  `json:"value"`
         
    }
    
    // SecureUrl used by Configuration
    type SecureUrl struct {

        
            SecureURL string  `json:"secure_url"`
         
    }
    
    // Application used by Configuration
    type Application struct {

        
            Website ApplicationWebsite  `json:"website"`
            Cors ApplicationCors  `json:"cors"`
            Auth ApplicationAuth  `json:"auth"`
            Description string  `json:"description"`
            ChannelType string  `json:"channel_type"`
            CacheTtl float64  `json:"cache_ttl"`
            IsInternal bool  `json:"is_internal"`
            IsActive bool  `json:"is_active"`
            ID string  `json:"_id"`
            Name string  `json:"name"`
            Owner string  `json:"owner"`
            CompanyID float64  `json:"company_id"`
            Token string  `json:"token"`
            Redirections []ApplicationRedirections  `json:"redirections"`
            Meta []ApplicationMeta  `json:"meta"`
            CreatedAt string  `json:"created_at"`
            UpdatedAt string  `json:"updated_at"`
            V float64  `json:"__v"`
            Banner SecureUrl  `json:"banner"`
            Logo SecureUrl  `json:"logo"`
            Favicon SecureUrl  `json:"favicon"`
            Domains []Domain  `json:"domains"`
            AppType string  `json:"app_type"`
            MobileLogo SecureUrl  `json:"mobile_logo"`
            Domain Domain  `json:"domain"`
            Slug string  `json:"slug"`
         
    }
    
    // NotFound used by Configuration
    type NotFound struct {

        
            Message string  `json:"message"`
         
    }
    
    // UnhandledError used by Configuration
    type UnhandledError struct {

        
            Message string  `json:"message"`
         
    }
    
    // InvalidPayloadRequest used by Configuration
    type InvalidPayloadRequest struct {

        
            Message string  `json:"message"`
         
    }
    
    // SuccessMessageResponse used by Configuration
    type SuccessMessageResponse struct {

        
            Message string  `json:"message"`
         
    }
    
    // InventoryBrandRule used by Configuration
    type InventoryBrandRule struct {

        
            Criteria string  `json:"criteria"`
            Brands []float64  `json:"brands"`
         
    }
    
    // StoreCriteriaRule used by Configuration
    type StoreCriteriaRule struct {

        
            Companies []float64  `json:"companies"`
            Brands []float64  `json:"brands"`
         
    }
    
    // InventoryStoreRule used by Configuration
    type InventoryStoreRule struct {

        
            Criteria string  `json:"criteria"`
            Rules []StoreCriteriaRule  `json:"rules"`
            Stores []float64  `json:"stores"`
         
    }
    
    // InventoryPaymentConfig used by Configuration
    type InventoryPaymentConfig struct {

        
            ModeOfPayment string  `json:"mode_of_payment"`
            Source string  `json:"source"`
         
    }
    
    // StorePriorityRule used by Configuration
    type StorePriorityRule struct {

        
            Enabled bool  `json:"enabled"`
            StoretypeOrder []string  `json:"storetype_order"`
         
    }
    
    // ArticleAssignmentRule used by Configuration
    type ArticleAssignmentRule struct {

        
            StorePriority StorePriorityRule  `json:"store_priority"`
         
    }
    
    // InventoryArticleAssignment used by Configuration
    type InventoryArticleAssignment struct {

        
            PostOrderReassignment bool  `json:"post_order_reassignment"`
            Rules ArticleAssignmentRule  `json:"rules"`
         
    }
    
    // CompanyAboutAddress used by Configuration
    type CompanyAboutAddress struct {

        
            Pincode float64  `json:"pincode"`
            Address1 string  `json:"address1"`
            Address2 string  `json:"address2"`
            City string  `json:"city"`
            State string  `json:"state"`
            Country string  `json:"country"`
            AddressType string  `json:"address_type"`
         
    }
    
    // UserEmail used by Configuration
    type UserEmail struct {

        
            Active bool  `json:"active"`
            Primary bool  `json:"primary"`
            Verified bool  `json:"verified"`
            Email string  `json:"email"`
         
    }
    
    // UserPhoneNumber used by Configuration
    type UserPhoneNumber struct {

        
            Active bool  `json:"active"`
            Primary bool  `json:"primary"`
            Verified bool  `json:"verified"`
            CountryCode float64  `json:"country_code"`
            Phone string  `json:"phone"`
         
    }
    
    // Page used by Configuration
    type Page struct {

        
            ItemTotal float64  `json:"item_total"`
            NextID string  `json:"next_id"`
            HasPrevious bool  `json:"has_previous"`
            HasNext bool  `json:"has_next"`
            Current float64  `json:"current"`
            Type string  `json:"type"`
            Size float64  `json:"size"`
         
    }
    
    // ApplicationInformation used by Configuration
    type ApplicationInformation struct {

        
            Address InformationAddress  `json:"address"`
            Support InformationSupport  `json:"support"`
            SocialLinks SocialLinks  `json:"social_links"`
            Links Links  `json:"links"`
            CopyrightText string  `json:"copyright_text"`
            ID string  `json:"_id"`
            BusinessHighlights BusinessHighlights  `json:"business_highlights"`
            Application string  `json:"application"`
            CreatedAt string  `json:"created_at"`
            UpdatedAt string  `json:"updated_at"`
            V float64  `json:"__v"`
         
    }
    
    // InformationAddress used by Configuration
    type InformationAddress struct {

        
            Loc string  `json:"loc"`
            AddressLine []string  `json:"address_line"`
            Phone []InformationPhone  `json:"phone"`
            City string  `json:"city"`
            Country string  `json:"country"`
            Pincode float64  `json:"pincode"`
         
    }
    
    // InformationPhone used by Configuration
    type InformationPhone struct {

        
            Code string  `json:"code"`
            Number string  `json:"number"`
         
    }
    
    // InformationSupport used by Configuration
    type InformationSupport struct {

        
            Phone []string  `json:"phone"`
            Email []string  `json:"email"`
            Timing string  `json:"timing"`
         
    }
    
    // SocialLinks used by Configuration
    type SocialLinks struct {

        
            Facebook FacebookLink  `json:"facebook"`
            Instagram InstagramLink  `json:"instagram"`
            Twitter TwitterLink  `json:"twitter"`
            Pinterest PinterestLink  `json:"pinterest"`
            GooglePlus GooglePlusLink  `json:"google_plus"`
            Youtube YoutubeLink  `json:"youtube"`
            LinkedIn LinkedInLink  `json:"linked_in"`
            Vimeo VimeoLink  `json:"vimeo"`
            BlogLink BlogLink  `json:"blog_link"`
         
    }
    
    // FacebookLink used by Configuration
    type FacebookLink struct {

        
            Title string  `json:"title"`
            Icon string  `json:"icon"`
            Link string  `json:"link"`
         
    }
    
    // InstagramLink used by Configuration
    type InstagramLink struct {

        
            Title string  `json:"title"`
            Icon string  `json:"icon"`
            Link string  `json:"link"`
         
    }
    
    // TwitterLink used by Configuration
    type TwitterLink struct {

        
            Title string  `json:"title"`
            Icon string  `json:"icon"`
            Link string  `json:"link"`
         
    }
    
    // PinterestLink used by Configuration
    type PinterestLink struct {

        
            Title string  `json:"title"`
            Icon string  `json:"icon"`
            Link string  `json:"link"`
         
    }
    
    // GooglePlusLink used by Configuration
    type GooglePlusLink struct {

        
            Title string  `json:"title"`
            Icon string  `json:"icon"`
            Link string  `json:"link"`
         
    }
    
    // YoutubeLink used by Configuration
    type YoutubeLink struct {

        
            Title string  `json:"title"`
            Icon string  `json:"icon"`
            Link string  `json:"link"`
         
    }
    
    // LinkedInLink used by Configuration
    type LinkedInLink struct {

        
            Title string  `json:"title"`
            Icon string  `json:"icon"`
            Link string  `json:"link"`
         
    }
    
    // VimeoLink used by Configuration
    type VimeoLink struct {

        
            Title string  `json:"title"`
            Icon string  `json:"icon"`
            Link string  `json:"link"`
         
    }
    
    // BlogLink used by Configuration
    type BlogLink struct {

        
            Title string  `json:"title"`
            Icon string  `json:"icon"`
            Link string  `json:"link"`
         
    }
    
    // Links used by Configuration
    type Links struct {

        
            Title string  `json:"title"`
            Link string  `json:"link"`
         
    }
    
    // BusinessHighlights used by Configuration
    type BusinessHighlights struct {

        
            ID string  `json:"_id"`
            Title string  `json:"title"`
            Icon string  `json:"icon"`
            SubTitle string  `json:"sub_title"`
         
    }
    
    // ApplicationDetail used by Configuration
    type ApplicationDetail struct {

        
            Name string  `json:"name"`
            Description string  `json:"description"`
            Logo SecureUrl  `json:"logo"`
            MobileLogo SecureUrl  `json:"mobile_logo"`
            Favicon SecureUrl  `json:"favicon"`
            Banner SecureUrl  `json:"banner"`
            Domain Domain  `json:"domain"`
            Domains []Domain  `json:"domains"`
            ID string  `json:"_id"`
            Slug string  `json:"slug"`
         
    }
    
    // CurrenciesResponse used by Configuration
    type CurrenciesResponse struct {

        
            Items []Currency  `json:"items"`
         
    }
    
    // AppCurrencyResponse used by Configuration
    type AppCurrencyResponse struct {

        
            Application string  `json:"application"`
            DefaultCurrency DefaultCurrency  `json:"default_currency"`
            SupportedCurrency []Currency  `json:"supported_currency"`
         
    }
    
    // StoreLatLong used by Configuration
    type StoreLatLong struct {

        
            Type string  `json:"type"`
            Coordinates []float64  `json:"coordinates"`
         
    }
    
    // OptedStoreAddress used by Configuration
    type OptedStoreAddress struct {

        
            State string  `json:"state"`
            Address1 string  `json:"address1"`
            LatLong StoreLatLong  `json:"lat_long"`
            Address2 string  `json:"address2"`
            Pincode float64  `json:"pincode"`
            Country string  `json:"country"`
            City string  `json:"city"`
         
    }
    
    // OrderingStore used by Configuration
    type OrderingStore struct {

        
            Address OptedStoreAddress  `json:"address"`
            ID string  `json:"_id"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            DisplayName string  `json:"display_name"`
            StoreType string  `json:"store_type"`
            StoreCode string  `json:"store_code"`
            Pincode float64  `json:"pincode"`
            Code string  `json:"code"`
         
    }
    
    // OrderingStores used by Configuration
    type OrderingStores struct {

        
            Page Page  `json:"page"`
            Items []OrderingStore  `json:"items"`
            DeployedStores []float64  `json:"deployed_stores"`
            AllStores bool  `json:"all_stores"`
            Enabled bool  `json:"enabled"`
            Type string  `json:"type"`
            ID string  `json:"_id"`
            App string  `json:"app"`
            V float64  `json:"__v"`
         
    }
    
    // OrderingStoresResponse used by Configuration
    type OrderingStoresResponse struct {

        
            Page Page  `json:"page"`
            Items []OrderingStore  `json:"items"`
         
    }
    

    
    // Validation used by Cart
    type Validation struct {

        
            Anonymous bool  `json:"anonymous"`
            AppID []string  `json:"app_id"`
            UserRegisteredAfter string  `json:"user_registered_after"`
         
    }
    
    // CouponSchedule used by Cart
    type CouponSchedule struct {

        
            Cron string  `json:"cron"`
            NextSchedule []map[string]interface{}  `json:"next_schedule"`
            Duration float64  `json:"duration"`
            Start string  `json:"start"`
            End string  `json:"end"`
         
    }
    
    // CouponDateMeta used by Cart
    type CouponDateMeta struct {

        
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
         
    }
    
    // Rule used by Cart
    type Rule struct {

        
            Max float64  `json:"max"`
            Value float64  `json:"value"`
            DiscountQty float64  `json:"discount_qty"`
            Min float64  `json:"min"`
            Key float64  `json:"key"`
         
    }
    
    // Identifier used by Cart
    type Identifier struct {

        
            StoreID []float64  `json:"store_id"`
            CategoryID []float64  `json:"category_id"`
            ExcludeBrandID []float64  `json:"exclude_brand_id"`
            CollectionID []string  `json:"collection_id"`
            BrandID []float64  `json:"brand_id"`
            CompanyID []float64  `json:"company_id"`
            ArticleID []string  `json:"article_id"`
            ItemID []float64  `json:"item_id"`
            UserID []string  `json:"user_id"`
         
    }
    
    // State used by Cart
    type State struct {

        
            IsPublic bool  `json:"is_public"`
            IsArchived bool  `json:"is_archived"`
            IsDisplay bool  `json:"is_display"`
         
    }
    
    // Ownership used by Cart
    type Ownership struct {

        
            PayableCategory string  `json:"payable_category"`
            PayableBy string  `json:"payable_by"`
         
    }
    
    // DisplayMetaDict used by Cart
    type DisplayMetaDict struct {

        
            Subtitle string  `json:"subtitle"`
            Title string  `json:"title"`
         
    }
    
    // DisplayMeta used by Cart
    type DisplayMeta struct {

        
            Subtitle string  `json:"subtitle"`
            Title string  `json:"title"`
            Auto DisplayMetaDict  `json:"auto"`
            Description string  `json:"description"`
            Remove DisplayMetaDict  `json:"remove"`
            Apply DisplayMetaDict  `json:"apply"`
         
    }
    
    // PostOrder used by Cart
    type PostOrder struct {

        
            ReturnAllowed bool  `json:"return_allowed"`
            CancellationAllowed bool  `json:"cancellation_allowed"`
         
    }
    
    // BulkBundleRestriction used by Cart
    type BulkBundleRestriction struct {

        
            MultiStoreAllowed bool  `json:"multi_store_allowed"`
         
    }
    
    // UsesRemaining used by Cart
    type UsesRemaining struct {

        
            User float64  `json:"user"`
            Total float64  `json:"total"`
            App float64  `json:"app"`
         
    }
    
    // UsesRestriction used by Cart
    type UsesRestriction struct {

        
            Remaining UsesRemaining  `json:"remaining"`
            Maximum UsesRemaining  `json:"maximum"`
         
    }
    
    // PaymentAllowValue used by Cart
    type PaymentAllowValue struct {

        
            Max float64  `json:"max"`
         
    }
    
    // PaymentModes used by Cart
    type PaymentModes struct {

        
            Codes []string  `json:"codes"`
            Networks []string  `json:"networks"`
            Uses PaymentAllowValue  `json:"uses"`
            Types []string  `json:"types"`
         
    }
    
    // PriceRange used by Cart
    type PriceRange struct {

        
            Max float64  `json:"max"`
            Min float64  `json:"min"`
         
    }
    
    // Restrictions used by Cart
    type Restrictions struct {

        
            PostOrder PostOrder  `json:"post_order"`
            CouponAllowed bool  `json:"coupon_allowed"`
            UserGroups []float64  `json:"user_groups"`
            UserType string  `json:"user_type"`
            BulkBundle BulkBundleRestriction  `json:"bulk_bundle"`
            Uses UsesRestriction  `json:"uses"`
            OrderingStores []float64  `json:"ordering_stores"`
            Payments map[string]PaymentModes  `json:"payments"`
            Platforms []string  `json:"platforms"`
            PriceRange PriceRange  `json:"price_range"`
         
    }
    
    // CouponAuthor used by Cart
    type CouponAuthor struct {

        
            ModifiedBy string  `json:"modified_by"`
            CreatedBy string  `json:"created_by"`
         
    }
    
    // CouponAction used by Cart
    type CouponAction struct {

        
            TxnMode string  `json:"txn_mode"`
            ActionDate string  `json:"action_date"`
         
    }
    
    // RuleDefinition used by Cart
    type RuleDefinition struct {

        
            CalculateOn string  `json:"calculate_on"`
            IsExact bool  `json:"is_exact"`
            Scope []string  `json:"scope"`
            CurrencyCode string  `json:"currency_code"`
            ValueType string  `json:"value_type"`
            AutoApply bool  `json:"auto_apply"`
            Type string  `json:"type"`
            ApplicableOn string  `json:"applicable_on"`
         
    }
    
    // Validity used by Cart
    type Validity struct {

        
            Priority float64  `json:"priority"`
         
    }
    
    // CouponAdd used by Cart
    type CouponAdd struct {

        
            Tags []string  `json:"tags"`
            Validation Validation  `json:"validation"`
            Schedule CouponSchedule  `json:"_schedule"`
            DateMeta CouponDateMeta  `json:"date_meta"`
            Rule []Rule  `json:"rule"`
            Identifiers Identifier  `json:"identifiers"`
            State State  `json:"state"`
            Ownership Ownership  `json:"ownership"`
            DisplayMeta DisplayMeta  `json:"display_meta"`
            Code string  `json:"code"`
            Restrictions Restrictions  `json:"restrictions"`
            Author CouponAuthor  `json:"author"`
            TypeSlug string  `json:"type_slug"`
            Action CouponAction  `json:"action"`
            RuleDefinition RuleDefinition  `json:"rule_definition"`
            Validity Validity  `json:"validity"`
         
    }
    
    // Page used by Cart
    type Page struct {

        
            ItemTotal float64  `json:"item_total"`
            NextID string  `json:"next_id"`
            HasPrevious bool  `json:"has_previous"`
            HasNext bool  `json:"has_next"`
            Current float64  `json:"current"`
            Type string  `json:"type"`
            Size float64  `json:"size"`
         
    }
    
    // CouponsResponse used by Cart
    type CouponsResponse struct {

        
            Items CouponAdd  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // SuccessMessage used by Cart
    type SuccessMessage struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // OperationErrorResponse used by Cart
    type OperationErrorResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // CouponUpdate used by Cart
    type CouponUpdate struct {

        
            Tags []string  `json:"tags"`
            Validation Validation  `json:"validation"`
            Schedule CouponSchedule  `json:"_schedule"`
            DateMeta CouponDateMeta  `json:"date_meta"`
            Rule []Rule  `json:"rule"`
            Identifiers Identifier  `json:"identifiers"`
            State State  `json:"state"`
            Ownership Ownership  `json:"ownership"`
            DisplayMeta DisplayMeta  `json:"display_meta"`
            Code string  `json:"code"`
            Restrictions Restrictions  `json:"restrictions"`
            Author CouponAuthor  `json:"author"`
            TypeSlug string  `json:"type_slug"`
            Action CouponAction  `json:"action"`
            RuleDefinition RuleDefinition  `json:"rule_definition"`
            Validity Validity  `json:"validity"`
         
    }
    
    // CouponPartialUpdate used by Cart
    type CouponPartialUpdate struct {

        
            Archive bool  `json:"archive"`
            Schedule CouponSchedule  `json:"schedule"`
         
    }
    
    // CompareObject used by Cart
    type CompareObject struct {

        
            LessThan float64  `json:"less_than"`
            LessThanEquals float64  `json:"less_than_equals"`
            Equals float64  `json:"equals"`
            GreaterThan float64  `json:"greater_than"`
            GreaterThanEquals float64  `json:"greater_than_equals"`
         
    }
    
    // ItemCriteria used by Cart
    type ItemCriteria struct {

        
            ItemExcludeSku []string  `json:"item_exclude_sku"`
            ItemSize []string  `json:"item_size"`
            CartTotal CompareObject  `json:"cart_total"`
            CartQuantity CompareObject  `json:"cart_quantity"`
            BuyRules []string  `json:"buy_rules"`
            ItemSku []string  `json:"item_sku"`
            ItemBrand []float64  `json:"item_brand"`
            AvailableZones []string  `json:"available_zones"`
            ItemExcludeCategory []float64  `json:"item_exclude_category"`
            ItemExcludeCompany []float64  `json:"item_exclude_company"`
            ItemExcludeStore []float64  `json:"item_exclude_store"`
            ProductTags []string  `json:"product_tags"`
            CartUniqueItemQuantity CompareObject  `json:"cart_unique_item_quantity"`
            ItemStore []float64  `json:"item_store"`
            ItemCategory []float64  `json:"item_category"`
            AllItems bool  `json:"all_items"`
            ItemID []float64  `json:"item_id"`
            ItemCompany []float64  `json:"item_company"`
            CartUniqueItemAmount CompareObject  `json:"cart_unique_item_amount"`
            ItemExcludeID []float64  `json:"item_exclude_id"`
            ItemExcludeBrand []float64  `json:"item_exclude_brand"`
         
    }
    
    // DiscountOffer used by Cart
    type DiscountOffer struct {

        
            MaxUsagePerTransaction float64  `json:"max_usage_per_transaction"`
            DiscountPrice float64  `json:"discount_price"`
            MinOfferQuantity float64  `json:"min_offer_quantity"`
            DiscountAmount float64  `json:"discount_amount"`
            MaxOfferQuantity float64  `json:"max_offer_quantity"`
            DiscountPercentage float64  `json:"discount_percentage"`
            Code string  `json:"code"`
            PartialCanRet bool  `json:"partial_can_ret"`
            MaxDiscountAmount float64  `json:"max_discount_amount"`
            ApportionDiscount bool  `json:"apportion_discount"`
         
    }
    
    // DiscountRule used by Cart
    type DiscountRule struct {

        
            DiscountType string  `json:"discount_type"`
            BuyCondition string  `json:"buy_condition"`
            ItemCriteria ItemCriteria  `json:"item_criteria"`
            Offer DiscountOffer  `json:"offer"`
         
    }
    
    // Ownership1 used by Cart
    type Ownership1 struct {

        
            PayableCategory string  `json:"payable_category"`
            PayableBy string  `json:"payable_by"`
         
    }
    
    // PostOrder1 used by Cart
    type PostOrder1 struct {

        
            ReturnAllowed bool  `json:"return_allowed"`
            CancellationAllowed bool  `json:"cancellation_allowed"`
         
    }
    
    // UsesRemaining1 used by Cart
    type UsesRemaining1 struct {

        
            User float64  `json:"user"`
            Total float64  `json:"total"`
         
    }
    
    // UsesRestriction1 used by Cart
    type UsesRestriction1 struct {

        
            Remaining UsesRemaining1  `json:"remaining"`
            Maximum UsesRemaining1  `json:"maximum"`
         
    }
    
    // UserRegistered used by Cart
    type UserRegistered struct {

        
            Start string  `json:"start"`
            End string  `json:"end"`
         
    }
    
    // PaymentAllowValue1 used by Cart
    type PaymentAllowValue1 struct {

        
            Max float64  `json:"max"`
         
    }
    
    // PromotionPaymentModes used by Cart
    type PromotionPaymentModes struct {

        
            Codes []string  `json:"codes"`
            Uses PaymentAllowValue1  `json:"uses"`
            Type string  `json:"type"`
         
    }
    
    // Restrictions1 used by Cart
    type Restrictions1 struct {

        
            PostOrder PostOrder1  `json:"post_order"`
            AnonymousUsers bool  `json:"anonymous_users"`
            OrderQuantity float64  `json:"order_quantity"`
            UserGroups []float64  `json:"user_groups"`
            Uses UsesRestriction1  `json:"uses"`
            UserRegistered UserRegistered  `json:"user_registered"`
            Payments []PromotionPaymentModes  `json:"payments"`
            Platforms []string  `json:"platforms"`
            UserID []string  `json:"user_id"`
         
    }
    
    // PromotionSchedule used by Cart
    type PromotionSchedule struct {

        
            Cron string  `json:"cron"`
            NextSchedule []map[string]interface{}  `json:"next_schedule"`
            Duration float64  `json:"duration"`
            Start string  `json:"start"`
            End string  `json:"end"`
            Published bool  `json:"published"`
         
    }
    
    // DisplayMeta1 used by Cart
    type DisplayMeta1 struct {

        
            Name string  `json:"name"`
            OfferLabel string  `json:"offer_label"`
            Description string  `json:"description"`
            OfferText string  `json:"offer_text"`
         
    }
    
    // PromotionDateMeta used by Cart
    type PromotionDateMeta struct {

        
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
         
    }
    
    // PromotionAuthor used by Cart
    type PromotionAuthor struct {

        
            ModifiedBy string  `json:"modified_by"`
            CreatedBy string  `json:"created_by"`
         
    }
    
    // Visibility used by Cart
    type Visibility struct {

        
            CouponList bool  `json:"coupon_list"`
            Pdp bool  `json:"pdp"`
         
    }
    
    // PromotionAction used by Cart
    type PromotionAction struct {

        
            ActionType string  `json:"action_type"`
            ActionDate string  `json:"action_date"`
         
    }
    
    // PromotionListItem used by Cart
    type PromotionListItem struct {

        
            Mode string  `json:"mode"`
            DiscountRules []DiscountRule  `json:"discount_rules"`
            Ownership Ownership1  `json:"ownership"`
            Restrictions Restrictions1  `json:"restrictions"`
            PromotionType string  `json:"promotion_type"`
            ApplyAllDiscount bool  `json:"apply_all_discount"`
            PromoGroup string  `json:"promo_group"`
            BuyRules map[string]ItemCriteria  `json:"buy_rules"`
            ApplicationID string  `json:"application_id"`
            CalculateOn string  `json:"calculate_on"`
            Schedule PromotionSchedule  `json:"_schedule"`
            DisplayMeta DisplayMeta1  `json:"display_meta"`
            Code string  `json:"code"`
            Stackable bool  `json:"stackable"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Currency string  `json:"currency"`
            DateMeta PromotionDateMeta  `json:"date_meta"`
            Author PromotionAuthor  `json:"author"`
            ApplyPriority float64  `json:"apply_priority"`
            Visiblility Visibility  `json:"visiblility"`
            ApplyExclusive string  `json:"apply_exclusive"`
            PostOrderAction PromotionAction  `json:"post_order_action"`
         
    }
    
    // PromotionsResponse used by Cart
    type PromotionsResponse struct {

        
            Items PromotionListItem  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // PromotionAdd used by Cart
    type PromotionAdd struct {

        
            Mode string  `json:"mode"`
            DiscountRules []DiscountRule  `json:"discount_rules"`
            Ownership Ownership1  `json:"ownership"`
            Restrictions Restrictions1  `json:"restrictions"`
            PromotionType string  `json:"promotion_type"`
            ApplyAllDiscount bool  `json:"apply_all_discount"`
            PromoGroup string  `json:"promo_group"`
            BuyRules map[string]ItemCriteria  `json:"buy_rules"`
            ApplicationID string  `json:"application_id"`
            CalculateOn string  `json:"calculate_on"`
            Schedule PromotionSchedule  `json:"_schedule"`
            DisplayMeta DisplayMeta1  `json:"display_meta"`
            Code string  `json:"code"`
            Stackable bool  `json:"stackable"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Currency string  `json:"currency"`
            DateMeta PromotionDateMeta  `json:"date_meta"`
            Author PromotionAuthor  `json:"author"`
            ApplyPriority float64  `json:"apply_priority"`
            Visiblility Visibility  `json:"visiblility"`
            ApplyExclusive string  `json:"apply_exclusive"`
            PostOrderAction PromotionAction  `json:"post_order_action"`
         
    }
    
    // PromotionUpdate used by Cart
    type PromotionUpdate struct {

        
            Mode string  `json:"mode"`
            DiscountRules []DiscountRule  `json:"discount_rules"`
            Ownership Ownership1  `json:"ownership"`
            Restrictions Restrictions1  `json:"restrictions"`
            PromotionType string  `json:"promotion_type"`
            ApplyAllDiscount bool  `json:"apply_all_discount"`
            PromoGroup string  `json:"promo_group"`
            BuyRules map[string]ItemCriteria  `json:"buy_rules"`
            ApplicationID string  `json:"application_id"`
            CalculateOn string  `json:"calculate_on"`
            Schedule PromotionSchedule  `json:"_schedule"`
            DisplayMeta DisplayMeta1  `json:"display_meta"`
            Code string  `json:"code"`
            Stackable bool  `json:"stackable"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Currency string  `json:"currency"`
            DateMeta PromotionDateMeta  `json:"date_meta"`
            Author PromotionAuthor  `json:"author"`
            ApplyPriority float64  `json:"apply_priority"`
            Visiblility Visibility  `json:"visiblility"`
            ApplyExclusive string  `json:"apply_exclusive"`
            PostOrderAction PromotionAction  `json:"post_order_action"`
         
    }
    
    // PromotionPartialUpdate used by Cart
    type PromotionPartialUpdate struct {

        
            Archive bool  `json:"archive"`
            Schedule PromotionSchedule  `json:"schedule"`
         
    }
    
    // ActivePromosResponse used by Cart
    type ActivePromosResponse struct {

        
            Subtitle string  `json:"subtitle"`
            Title string  `json:"title"`
            EntityType string  `json:"entity_type"`
            EntitySlug string  `json:"entity_slug"`
            CreatedOn string  `json:"created_on"`
            Example string  `json:"example"`
            Type string  `json:"type"`
            ModifiedOn string  `json:"modified_on"`
            IsHidden bool  `json:"is_hidden"`
            Description string  `json:"description"`
         
    }
    
    // CartItem used by Cart
    type CartItem struct {

        
            Quantity float64  `json:"quantity"`
            ProductID string  `json:"product_id"`
            Size string  `json:"size"`
         
    }
    
    // OpenapiCartDetailsRequest used by Cart
    type OpenapiCartDetailsRequest struct {

        
            CartItems CartItem  `json:"cart_items"`
         
    }
    
    // RawBreakup used by Cart
    type RawBreakup struct {

        
            Subtotal float64  `json:"subtotal"`
            ConvenienceFee float64  `json:"convenience_fee"`
            Vog float64  `json:"vog"`
            FyndCash float64  `json:"fynd_cash"`
            DeliveryCharge float64  `json:"delivery_charge"`
            YouSaved float64  `json:"you_saved"`
            CodCharge float64  `json:"cod_charge"`
            Coupon float64  `json:"coupon"`
            Total float64  `json:"total"`
            Discount float64  `json:"discount"`
            GstCharges float64  `json:"gst_charges"`
            MrpTotal float64  `json:"mrp_total"`
         
    }
    
    // CouponBreakup used by Cart
    type CouponBreakup struct {

        
            MaxDiscountValue float64  `json:"max_discount_value"`
            UID string  `json:"uid"`
            SubTitle string  `json:"sub_title"`
            Title string  `json:"title"`
            Value float64  `json:"value"`
            Message string  `json:"message"`
            Code string  `json:"code"`
            Type string  `json:"type"`
            CouponType string  `json:"coupon_type"`
            CouponValue float64  `json:"coupon_value"`
            MinimumCartValue float64  `json:"minimum_cart_value"`
            Description string  `json:"description"`
            IsApplied bool  `json:"is_applied"`
         
    }
    
    // DisplayBreakup used by Cart
    type DisplayBreakup struct {

        
            Display string  `json:"display"`
            CurrencyCode string  `json:"currency_code"`
            Value float64  `json:"value"`
            Message []string  `json:"message"`
            CurrencySymbol string  `json:"currency_symbol"`
            Key string  `json:"key"`
         
    }
    
    // LoyaltyPoints used by Cart
    type LoyaltyPoints struct {

        
            Description string  `json:"description"`
            IsApplied bool  `json:"is_applied"`
            Total float64  `json:"total"`
            Applicable float64  `json:"applicable"`
         
    }
    
    // CartBreakup used by Cart
    type CartBreakup struct {

        
            Raw RawBreakup  `json:"raw"`
            Coupon CouponBreakup  `json:"coupon"`
            Display []DisplayBreakup  `json:"display"`
            LoyaltyPoints LoyaltyPoints  `json:"loyalty_points"`
         
    }
    
    // ProductPrice used by Cart
    type ProductPrice struct {

        
            AddOn float64  `json:"add_on"`
            CurrencyCode string  `json:"currency_code"`
            Effective float64  `json:"effective"`
            Marked float64  `json:"marked"`
            CurrencySymbol string  `json:"currency_symbol"`
            Selling float64  `json:"selling"`
         
    }
    
    // ProductPriceInfo used by Cart
    type ProductPriceInfo struct {

        
            Base ProductPrice  `json:"base"`
            Converted ProductPrice  `json:"converted"`
         
    }
    
    // BaseInfo used by Cart
    type BaseInfo struct {

        
            Name string  `json:"name"`
            UID float64  `json:"uid"`
         
    }
    
    // BasePrice used by Cart
    type BasePrice struct {

        
            Marked float64  `json:"marked"`
            CurrencySymbol string  `json:"currency_symbol"`
            Effective float64  `json:"effective"`
            CurrencyCode string  `json:"currency_code"`
         
    }
    
    // ArticlePriceInfo used by Cart
    type ArticlePriceInfo struct {

        
            Base BasePrice  `json:"base"`
            Converted BasePrice  `json:"converted"`
         
    }
    
    // ProductArticle used by Cart
    type ProductArticle struct {

        
            UID string  `json:"uid"`
            Seller BaseInfo  `json:"seller"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            ProductGroupTags []string  `json:"product_group_tags"`
            Store BaseInfo  `json:"store"`
            Type string  `json:"type"`
            Size string  `json:"size"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Price ArticlePriceInfo  `json:"price"`
            Quantity float64  `json:"quantity"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
         
    }
    
    // PromoMeta used by Cart
    type PromoMeta struct {

        
            Message string  `json:"message"`
         
    }
    
    // ProductAvailability used by Cart
    type ProductAvailability struct {

        
            Deliverable bool  `json:"deliverable"`
            Sizes []string  `json:"sizes"`
            IsValid bool  `json:"is_valid"`
            OutOfStock bool  `json:"out_of_stock"`
            OtherStoreQuantity float64  `json:"other_store_quantity"`
         
    }
    
    // BuyRules used by Cart
    type BuyRules struct {

        
            ItemCriteria map[string]interface{}  `json:"item_criteria"`
            CartConditions map[string]interface{}  `json:"cart_conditions"`
         
    }
    
    // DiscountRulesApp used by Cart
    type DiscountRulesApp struct {

        
            RawOffer map[string]interface{}  `json:"raw_offer"`
            Offer map[string]interface{}  `json:"offer"`
            ItemCriteria map[string]interface{}  `json:"item_criteria"`
            MatchedBuyRules []string  `json:"matched_buy_rules"`
         
    }
    
    // Ownership2 used by Cart
    type Ownership2 struct {

        
            PayableCategory string  `json:"payable_category"`
            PayableBy string  `json:"payable_by"`
         
    }
    
    // FreeGiftItem used by Cart
    type FreeGiftItem struct {

        
            ItemName string  `json:"item_name"`
            ItemImagesURL []string  `json:"item_images_url"`
            ItemBrandName string  `json:"item_brand_name"`
            ItemPriceDetails map[string]interface{}  `json:"item_price_details"`
            ItemSlug string  `json:"item_slug"`
            ItemID float64  `json:"item_id"`
         
    }
    
    // AppliedFreeArticles used by Cart
    type AppliedFreeArticles struct {

        
            FreeGiftItemDetails FreeGiftItem  `json:"free_gift_item_details"`
            ParentItemIdentifier string  `json:"parent_item_identifier"`
            Quantity float64  `json:"quantity"`
            ArticleID string  `json:"article_id"`
         
    }
    
    // AppliedPromotion used by Cart
    type AppliedPromotion struct {

        
            OfferText string  `json:"offer_text"`
            PromotionGroup string  `json:"promotion_group"`
            BuyRules []BuyRules  `json:"buy_rules"`
            PromoID string  `json:"promo_id"`
            MrpPromotion bool  `json:"mrp_promotion"`
            PromotionName string  `json:"promotion_name"`
            ArticleQuantity float64  `json:"article_quantity"`
            DiscountRules []DiscountRulesApp  `json:"discount_rules"`
            Amount float64  `json:"amount"`
            Ownership Ownership2  `json:"ownership"`
            AppliedFreeArticles []AppliedFreeArticles  `json:"applied_free_articles"`
            PromotionType string  `json:"promotion_type"`
         
    }
    
    // CartProductIdentifer used by Cart
    type CartProductIdentifer struct {

        
            Identifier string  `json:"identifier"`
         
    }
    
    // CategoryInfo used by Cart
    type CategoryInfo struct {

        
            Name string  `json:"name"`
            UID float64  `json:"uid"`
         
    }
    
    // ProductImage used by Cart
    type ProductImage struct {

        
            SecureURL string  `json:"secure_url"`
            URL string  `json:"url"`
            AspectRatio string  `json:"aspect_ratio"`
         
    }
    
    // NetQuantity used by Cart
    type NetQuantity struct {

        
            Value string  `json:"value"`
            Unit string  `json:"unit"`
         
    }
    
    // ActionQuery used by Cart
    type ActionQuery struct {

        
            ProductSlug []string  `json:"product_slug"`
         
    }
    
    // ProductAction used by Cart
    type ProductAction struct {

        
            URL string  `json:"url"`
            Query ActionQuery  `json:"query"`
            Type string  `json:"type"`
         
    }
    
    // CartProduct used by Cart
    type CartProduct struct {

        
            Categories []CategoryInfo  `json:"categories"`
            Images []ProductImage  `json:"images"`
            ItemCode string  `json:"item_code"`
            UID float64  `json:"uid"`
            NetQuantity NetQuantity  `json:"net_quantity"`
            Type string  `json:"type"`
            Name string  `json:"name"`
            Action ProductAction  `json:"action"`
            Slug string  `json:"slug"`
            Brand BaseInfo  `json:"brand"`
         
    }
    
    // CartProductInfo used by Cart
    type CartProductInfo struct {

        
            PricePerUnit ProductPriceInfo  `json:"price_per_unit"`
            Discount string  `json:"discount"`
            Article ProductArticle  `json:"article"`
            PromoMeta PromoMeta  `json:"promo_meta"`
            Availability ProductAvailability  `json:"availability"`
            PromotionsApplied []AppliedPromotion  `json:"promotions_applied"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
            Message string  `json:"message"`
            CouponMessage string  `json:"coupon_message"`
            IsSet bool  `json:"is_set"`
            BulkOffer map[string]interface{}  `json:"bulk_offer"`
            Price ProductPriceInfo  `json:"price"`
            Quantity float64  `json:"quantity"`
            Key string  `json:"key"`
            Product CartProduct  `json:"product"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
         
    }
    
    // OpenapiCartDetailsResponse used by Cart
    type OpenapiCartDetailsResponse struct {

        
            BreakupValues CartBreakup  `json:"breakup_values"`
            Items []CartProductInfo  `json:"items"`
            Message string  `json:"message"`
            IsValid bool  `json:"is_valid"`
         
    }
    
    // OpenApiErrorResponse used by Cart
    type OpenApiErrorResponse struct {

        
            Errors map[string]interface{}  `json:"errors"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // ShippingAddress used by Cart
    type ShippingAddress struct {

        
            Country string  `json:"country"`
            Area string  `json:"area"`
            AreaCode string  `json:"area_code"`
            AreaCodeSlug string  `json:"area_code_slug"`
            City string  `json:"city"`
            CountryCode string  `json:"country_code"`
            Meta map[string]interface{}  `json:"meta"`
            AddressType string  `json:"address_type"`
            Pincode float64  `json:"pincode"`
            Email string  `json:"email"`
            State string  `json:"state"`
            Phone float64  `json:"phone"`
            Name string  `json:"name"`
            Landmark string  `json:"landmark"`
            Address string  `json:"address"`
            CountryIsoCode string  `json:"country_iso_code"`
            CountryPhoneCode string  `json:"country_phone_code"`
         
    }
    
    // OpenApiCartServiceabilityRequest used by Cart
    type OpenApiCartServiceabilityRequest struct {

        
            CartItems CartItem  `json:"cart_items"`
            ShippingAddress ShippingAddress  `json:"shipping_address"`
         
    }
    
    // PromiseFormatted used by Cart
    type PromiseFormatted struct {

        
            Max string  `json:"max"`
            Min string  `json:"min"`
         
    }
    
    // PromiseTimestamp used by Cart
    type PromiseTimestamp struct {

        
            Max float64  `json:"max"`
            Min float64  `json:"min"`
         
    }
    
    // ShipmentPromise used by Cart
    type ShipmentPromise struct {

        
            Formatted PromiseFormatted  `json:"formatted"`
            Timestamp PromiseTimestamp  `json:"timestamp"`
         
    }
    
    // OpenApiCartServiceabilityResponse used by Cart
    type OpenApiCartServiceabilityResponse struct {

        
            BreakupValues CartBreakup  `json:"breakup_values"`
            Items []CartProductInfo  `json:"items"`
            Message string  `json:"message"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            IsValid bool  `json:"is_valid"`
         
    }
    
    // CartItemMeta used by Cart
    type CartItemMeta struct {

        
            GroupID string  `json:"group_id"`
            PrimaryItem bool  `json:"primary_item"`
         
    }
    
    // OpenApiFiles used by Cart
    type OpenApiFiles struct {

        
            Key string  `json:"key"`
            Values []string  `json:"values"`
         
    }
    
    // MultiTenderPaymentMeta used by Cart
    type MultiTenderPaymentMeta struct {

        
            OrderID string  `json:"order_id"`
            PaymentGateway string  `json:"payment_gateway"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            PaymentID string  `json:"payment_id"`
            CurrentStatus string  `json:"current_status"`
         
    }
    
    // MultiTenderPaymentMethod used by Cart
    type MultiTenderPaymentMethod struct {

        
            Mode string  `json:"mode"`
            Amount float64  `json:"amount"`
            Meta MultiTenderPaymentMeta  `json:"meta"`
            Name string  `json:"name"`
         
    }
    
    // OpenApiOrderItem used by Cart
    type OpenApiOrderItem struct {

        
            CodCharges float64  `json:"cod_charges"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            AmountPaid float64  `json:"amount_paid"`
            DeliveryCharges float64  `json:"delivery_charges"`
            Meta CartItemMeta  `json:"meta"`
            PriceEffective float64  `json:"price_effective"`
            Size string  `json:"size"`
            PriceMarked float64  `json:"price_marked"`
            Files []OpenApiFiles  `json:"files"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            PaymentMethods []MultiTenderPaymentMethod  `json:"payment_methods"`
            CashbackApplied float64  `json:"cashback_applied"`
            EmployeeDiscount float64  `json:"employee_discount"`
            Discount float64  `json:"discount"`
            Quantity float64  `json:"quantity"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            ProductID float64  `json:"product_id"`
         
    }
    
    // OpenApiPlatformCheckoutReq used by Cart
    type OpenApiPlatformCheckoutReq struct {

        
            BillingAddress ShippingAddress  `json:"billing_address"`
            CodCharges float64  `json:"cod_charges"`
            CartItems []OpenApiOrderItem  `json:"cart_items"`
            CartValue float64  `json:"cart_value"`
            Coupon string  `json:"coupon"`
            CashbackApplied float64  `json:"cashback_applied"`
            EmployeeDiscount map[string]interface{}  `json:"employee_discount"`
            CouponValue float64  `json:"coupon_value"`
            PaymentMode string  `json:"payment_mode"`
            CouponCode string  `json:"coupon_code"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            OrderID string  `json:"order_id"`
            Gstin string  `json:"gstin"`
            CurrencyCode string  `json:"currency_code"`
            Files []OpenApiFiles  `json:"files"`
            Comment string  `json:"comment"`
            ShippingAddress ShippingAddress  `json:"shipping_address"`
            PaymentMethods []MultiTenderPaymentMethod  `json:"payment_methods"`
            DeliveryCharges float64  `json:"delivery_charges"`
         
    }
    
    // OpenApiCheckoutResponse used by Cart
    type OpenApiCheckoutResponse struct {

        
            OrderRefID string  `json:"order_ref_id"`
            Success bool  `json:"success"`
            OrderID string  `json:"order_id"`
            Message string  `json:"message"`
         
    }
    
    // AbandonedCart used by Cart
    type AbandonedCart struct {

        
            PickUpCustomerDetails map[string]interface{}  `json:"pick_up_customer_details"`
            Meta map[string]interface{}  `json:"meta"`
            CartValue float64  `json:"cart_value"`
            FyndCredits map[string]interface{}  `json:"fynd_credits"`
            CreatedOn string  `json:"created_on"`
            IsDefault bool  `json:"is_default"`
            IsArchive bool  `json:"is_archive"`
            LastModified string  `json:"last_modified"`
            PaymentMode string  `json:"payment_mode"`
            UID float64  `json:"uid"`
            OrderID string  `json:"order_id"`
            Gstin string  `json:"gstin"`
            Articles []map[string]interface{}  `json:"articles"`
            Comment string  `json:"comment"`
            ExpireAt string  `json:"expire_at"`
            MergeQty bool  `json:"merge_qty"`
            Discount float64  `json:"discount"`
            Promotion map[string]interface{}  `json:"promotion"`
            CodCharges map[string]interface{}  `json:"cod_charges"`
            Cashback map[string]interface{}  `json:"cashback"`
            FcIndexMap []float64  `json:"fc_index_map"`
            Shipments []map[string]interface{}  `json:"shipments"`
            ID string  `json:"_id"`
            AppID string  `json:"app_id"`
            Coupon map[string]interface{}  `json:"coupon"`
            IsActive bool  `json:"is_active"`
            CheckoutMode string  `json:"checkout_mode"`
            BulkCouponDiscount float64  `json:"bulk_coupon_discount"`
            BuyNow bool  `json:"buy_now"`
            Payments map[string]interface{}  `json:"payments"`
            PaymentMethods []map[string]interface{}  `json:"payment_methods"`
            DeliveryCharges map[string]interface{}  `json:"delivery_charges"`
            UserID string  `json:"user_id"`
         
    }
    
    // AbandonedCartResponse used by Cart
    type AbandonedCartResponse struct {

        
            Items []AbandonedCart  `json:"items"`
            Message string  `json:"message"`
            Page Page  `json:"page"`
            Success bool  `json:"success"`
            Result map[string]interface{}  `json:"result"`
         
    }
    
    // PaymentSelectionLock used by Cart
    type PaymentSelectionLock struct {

        
            PaymentIdentifier string  `json:"payment_identifier"`
            Enabled bool  `json:"enabled"`
            DefaultOptions string  `json:"default_options"`
         
    }
    
    // CartCurrency used by Cart
    type CartCurrency struct {

        
            Code string  `json:"code"`
            Symbol string  `json:"symbol"`
         
    }
    
    // CartDetailResponse used by Cart
    type CartDetailResponse struct {

        
            BreakupValues CartBreakup  `json:"breakup_values"`
            PanNo string  `json:"pan_no"`
            Items []CartProductInfo  `json:"items"`
            Gstin string  `json:"gstin"`
            CouponText string  `json:"coupon_text"`
            ID string  `json:"id"`
            PanConfig map[string]interface{}  `json:"pan_config"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            Message string  `json:"message"`
            Comment string  `json:"comment"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            LastModified string  `json:"last_modified"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            BuyNow bool  `json:"buy_now"`
            IsValid bool  `json:"is_valid"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            CheckoutMode string  `json:"checkout_mode"`
            Currency CartCurrency  `json:"currency"`
         
    }
    
    // AddProductCart used by Cart
    type AddProductCart struct {

        
            StoreID float64  `json:"store_id"`
            Display string  `json:"display"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ItemSize string  `json:"item_size"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            ProductGroupTags []string  `json:"product_group_tags"`
            ArticleAssignment map[string]interface{}  `json:"article_assignment"`
            SellerID float64  `json:"seller_id"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            Pos bool  `json:"pos"`
            Quantity float64  `json:"quantity"`
            ArticleID string  `json:"article_id"`
            ItemID float64  `json:"item_id"`
         
    }
    
    // AddCartRequest used by Cart
    type AddCartRequest struct {

        
            Items []AddProductCart  `json:"items"`
            NewCart bool  `json:"new_cart"`
         
    }
    
    // AddCartDetailResponse used by Cart
    type AddCartDetailResponse struct {

        
            Cart CartDetailResponse  `json:"cart"`
            Success bool  `json:"success"`
            Partial bool  `json:"partial"`
            Message string  `json:"message"`
         
    }
    
    // UpdateProductCart used by Cart
    type UpdateProductCart struct {

        
            ItemSize string  `json:"item_size"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
            ItemID float64  `json:"item_id"`
            ItemIndex float64  `json:"item_index"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Quantity float64  `json:"quantity"`
            ArticleID string  `json:"article_id"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
         
    }
    
    // UpdateCartRequest used by Cart
    type UpdateCartRequest struct {

        
            Items []UpdateProductCart  `json:"items"`
            Operation string  `json:"operation"`
         
    }
    
    // UpdateCartDetailResponse used by Cart
    type UpdateCartDetailResponse struct {

        
            Cart CartDetailResponse  `json:"cart"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // GetShareCartLinkRequest used by Cart
    type GetShareCartLinkRequest struct {

        
            Meta map[string]interface{}  `json:"meta"`
            ID string  `json:"id"`
         
    }
    
    // GetShareCartLinkResponse used by Cart
    type GetShareCartLinkResponse struct {

        
            ShareURL string  `json:"share_url"`
            Token string  `json:"token"`
         
    }
    
    // SharedCartDetails used by Cart
    type SharedCartDetails struct {

        
            User map[string]interface{}  `json:"user"`
            Meta map[string]interface{}  `json:"meta"`
            Token string  `json:"token"`
            CreatedOn string  `json:"created_on"`
            Source map[string]interface{}  `json:"source"`
         
    }
    
    // SharedCart used by Cart
    type SharedCart struct {

        
            Message string  `json:"message"`
            IsValid bool  `json:"is_valid"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            LastModified string  `json:"last_modified"`
            SharedCartDetails SharedCartDetails  `json:"shared_cart_details"`
            UID string  `json:"uid"`
            Gstin string  `json:"gstin"`
            CouponText string  `json:"coupon_text"`
            ID string  `json:"id"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            Comment string  `json:"comment"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            Currency CartCurrency  `json:"currency"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            Items []CartProductInfo  `json:"items"`
            CartID float64  `json:"cart_id"`
            BuyNow bool  `json:"buy_now"`
            CheckoutMode string  `json:"checkout_mode"`
         
    }
    
    // SharedCartResponse used by Cart
    type SharedCartResponse struct {

        
            Cart SharedCart  `json:"cart"`
            Error string  `json:"error"`
         
    }
    
    // CartList used by Cart
    type CartList struct {

        
            PickUpCustomerDetails string  `json:"pick_up_customer_details"`
            CartValue string  `json:"cart_value"`
            CreatedOn string  `json:"created_on"`
            CartID string  `json:"cart_id"`
            ItemCounts string  `json:"item_counts"`
            UserID string  `json:"user_id"`
         
    }
    
    // MultiCartResponse used by Cart
    type MultiCartResponse struct {

        
            Success bool  `json:"success"`
            Data []CartList  `json:"data"`
         
    }
    
    // UpdateUserCartMapping used by Cart
    type UpdateUserCartMapping struct {

        
            UserID string  `json:"user_id"`
         
    }
    
    // UserInfo used by Cart
    type UserInfo struct {

        
            UID string  `json:"uid"`
            Mobile string  `json:"mobile"`
            ID string  `json:"_id"`
            LastName string  `json:"last_name"`
            CreatedAt string  `json:"created_at"`
            ExternalID string  `json:"external_id"`
            Gender string  `json:"gender"`
            FirstName string  `json:"first_name"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // UserCartMappingResponse used by Cart
    type UserCartMappingResponse struct {

        
            Message string  `json:"message"`
            IsValid bool  `json:"is_valid"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            LastModified string  `json:"last_modified"`
            User UserInfo  `json:"user"`
            Gstin string  `json:"gstin"`
            CouponText string  `json:"coupon_text"`
            ID string  `json:"id"`
            PanConfig map[string]interface{}  `json:"pan_config"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            Comment string  `json:"comment"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            Currency CartCurrency  `json:"currency"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            PanNo string  `json:"pan_no"`
            Items []CartProductInfo  `json:"items"`
            BuyNow bool  `json:"buy_now"`
            CheckoutMode string  `json:"checkout_mode"`
         
    }
    
    // DeleteCartDetailResponse used by Cart
    type DeleteCartDetailResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // CartItemCountResponse used by Cart
    type CartItemCountResponse struct {

        
            UserCartItemsCount float64  `json:"user_cart_items_count"`
         
    }
    
    // Coupon used by Cart
    type Coupon struct {

        
            MaxDiscountValue float64  `json:"max_discount_value"`
            SubTitle string  `json:"sub_title"`
            Title string  `json:"title"`
            IsApplied bool  `json:"is_applied"`
            ExpiresOn string  `json:"expires_on"`
            Message string  `json:"message"`
            CouponType string  `json:"coupon_type"`
            CouponValue float64  `json:"coupon_value"`
            IsApplicable bool  `json:"is_applicable"`
            MinimumCartValue float64  `json:"minimum_cart_value"`
            Description string  `json:"description"`
            CouponCode string  `json:"coupon_code"`
         
    }
    
    // PageCoupon used by Cart
    type PageCoupon struct {

        
            HasPrevious bool  `json:"has_previous"`
            TotalItemCount float64  `json:"total_item_count"`
            Current float64  `json:"current"`
            Total float64  `json:"total"`
            HasNext bool  `json:"has_next"`
         
    }
    
    // GetCouponResponse used by Cart
    type GetCouponResponse struct {

        
            AvailableCouponList []Coupon  `json:"available_coupon_list"`
            Page PageCoupon  `json:"page"`
         
    }
    
    // ApplyCouponRequest used by Cart
    type ApplyCouponRequest struct {

        
            CouponCode string  `json:"coupon_code"`
         
    }
    
    // GeoLocation used by Cart
    type GeoLocation struct {

        
            Latitude float64  `json:"latitude"`
            Longitude float64  `json:"longitude"`
         
    }
    
    // PlatformAddress used by Cart
    type PlatformAddress struct {

        
            AreaCode string  `json:"area_code"`
            CountryCode string  `json:"country_code"`
            Meta map[string]interface{}  `json:"meta"`
            Email string  `json:"email"`
            AddressType string  `json:"address_type"`
            Landmark string  `json:"landmark"`
            Address string  `json:"address"`
            Country string  `json:"country"`
            City string  `json:"city"`
            ID string  `json:"id"`
            State string  `json:"state"`
            Phone string  `json:"phone"`
            GoogleMapPoint map[string]interface{}  `json:"google_map_point"`
            GeoLocation GeoLocation  `json:"geo_location"`
            Name string  `json:"name"`
            Tags []string  `json:"tags"`
            CreatedByUserID string  `json:"created_by_user_id"`
            IsDefaultAddress bool  `json:"is_default_address"`
            Area string  `json:"area"`
            AreaCodeSlug string  `json:"area_code_slug"`
            IsActive bool  `json:"is_active"`
            CartID string  `json:"cart_id"`
            CheckoutMode string  `json:"checkout_mode"`
            UserID string  `json:"user_id"`
         
    }
    
    // PlatformGetAddressesResponse used by Cart
    type PlatformGetAddressesResponse struct {

        
            Address []PlatformAddress  `json:"address"`
         
    }
    
    // SaveAddressResponse used by Cart
    type SaveAddressResponse struct {

        
            Success bool  `json:"success"`
            ID string  `json:"id"`
            IsDefaultAddress bool  `json:"is_default_address"`
         
    }
    
    // UpdateAddressResponse used by Cart
    type UpdateAddressResponse struct {

        
            Success bool  `json:"success"`
            ID string  `json:"id"`
            IsUpdated bool  `json:"is_updated"`
            IsDefaultAddress bool  `json:"is_default_address"`
         
    }
    
    // DeleteAddressResponse used by Cart
    type DeleteAddressResponse struct {

        
            ID string  `json:"id"`
            IsDeleted bool  `json:"is_deleted"`
         
    }
    
    // PlatformSelectCartAddressRequest used by Cart
    type PlatformSelectCartAddressRequest struct {

        
            CheckoutMode string  `json:"checkout_mode"`
            ID string  `json:"id"`
            BillingAddressID string  `json:"billing_address_id"`
            CartID string  `json:"cart_id"`
         
    }
    
    // ShipmentResponse used by Cart
    type ShipmentResponse struct {

        
            DpID string  `json:"dp_id"`
            BoxType string  `json:"box_type"`
            ShipmentType string  `json:"shipment_type"`
            Items []CartProductInfo  `json:"items"`
            FulfillmentID float64  `json:"fulfillment_id"`
            Shipments float64  `json:"shipments"`
            Promise ShipmentPromise  `json:"promise"`
            DpOptions map[string]interface{}  `json:"dp_options"`
            OrderType string  `json:"order_type"`
            FulfillmentType string  `json:"fulfillment_type"`
         
    }
    
    // CartShipmentsResponse used by Cart
    type CartShipmentsResponse struct {

        
            Shipments []ShipmentResponse  `json:"shipments"`
            Error bool  `json:"error"`
            Message string  `json:"message"`
            IsValid bool  `json:"is_valid"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            LastModified string  `json:"last_modified"`
            Currency CartCurrency  `json:"currency"`
            UID string  `json:"uid"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            Gstin string  `json:"gstin"`
            CouponText string  `json:"coupon_text"`
            ID string  `json:"id"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            Comment string  `json:"comment"`
            CartID float64  `json:"cart_id"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            BuyNow bool  `json:"buy_now"`
            CheckoutMode string  `json:"checkout_mode"`
         
    }
    
    // UpdateCartShipmentItem used by Cart
    type UpdateCartShipmentItem struct {

        
            ShipmentType string  `json:"shipment_type"`
            ArticleUID string  `json:"article_uid"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // UpdateCartShipmentRequest used by Cart
    type UpdateCartShipmentRequest struct {

        
            Shipments []UpdateCartShipmentItem  `json:"shipments"`
         
    }
    
    // PlatformCartMetaRequest used by Cart
    type PlatformCartMetaRequest struct {

        
            PanNo string  `json:"pan_no"`
            PickUpCustomerDetails map[string]interface{}  `json:"pick_up_customer_details"`
            Gstin string  `json:"gstin"`
            Comment string  `json:"comment"`
            CheckoutMode string  `json:"checkout_mode"`
         
    }
    
    // CartMetaResponse used by Cart
    type CartMetaResponse struct {

        
            Message string  `json:"message"`
         
    }
    
    // CartMetaMissingResponse used by Cart
    type CartMetaMissingResponse struct {

        
            Errors []string  `json:"errors"`
         
    }
    
    // StaffCheckout used by Cart
    type StaffCheckout struct {

        
            User string  `json:"user"`
            EmployeeCode string  `json:"employee_code"`
            ID string  `json:"_id"`
            LastName string  `json:"last_name"`
            FirstName string  `json:"first_name"`
         
    }
    
    // Files used by Cart
    type Files struct {

        
            Key string  `json:"key"`
            Values []string  `json:"values"`
         
    }
    
    // PlatformCartCheckoutDetailRequest used by Cart
    type PlatformCartCheckoutDetailRequest struct {

        
            PickAtStoreUID float64  `json:"pick_at_store_uid"`
            Aggregator string  `json:"aggregator"`
            Meta map[string]interface{}  `json:"meta"`
            Staff StaffCheckout  `json:"staff"`
            PaymentMode string  `json:"payment_mode"`
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
            ID string  `json:"id"`
            Files []Files  `json:"files"`
            DeviceID string  `json:"device_id"`
            BillingAddressID string  `json:"billing_address_id"`
            Pos bool  `json:"pos"`
            PaymentAutoConfirm bool  `json:"payment_auto_confirm"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            OrderingStore float64  `json:"ordering_store"`
            MerchantCode string  `json:"merchant_code"`
            OrderType string  `json:"order_type"`
            PaymentIdentifier string  `json:"payment_identifier"`
            CallbackURL string  `json:"callback_url"`
            AddressID string  `json:"address_id"`
            CheckoutMode string  `json:"checkout_mode"`
            PaymentParams map[string]interface{}  `json:"payment_params"`
            UserID string  `json:"user_id"`
         
    }
    
    // CheckCart used by Cart
    type CheckCart struct {

        
            Message string  `json:"message"`
            StoreCode string  `json:"store_code"`
            IsValid bool  `json:"is_valid"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            LastModified string  `json:"last_modified"`
            Success bool  `json:"success"`
            UID string  `json:"uid"`
            CodAvailable bool  `json:"cod_available"`
            OrderID string  `json:"order_id"`
            Gstin string  `json:"gstin"`
            CouponText string  `json:"coupon_text"`
            ID string  `json:"id"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            Comment string  `json:"comment"`
            ErrorMessage string  `json:"error_message"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            CodCharges float64  `json:"cod_charges"`
            StoreEmps []map[string]interface{}  `json:"store_emps"`
            CodMessage string  `json:"cod_message"`
            Currency CartCurrency  `json:"currency"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            Items []CartProductInfo  `json:"items"`
            CartID float64  `json:"cart_id"`
            UserType string  `json:"user_type"`
            DeliveryChargeOrderValue float64  `json:"delivery_charge_order_value"`
            BuyNow bool  `json:"buy_now"`
            CheckoutMode string  `json:"checkout_mode"`
            DeliveryCharges float64  `json:"delivery_charges"`
         
    }
    
    // CartCheckoutResponse used by Cart
    type CartCheckoutResponse struct {

        
            CallbackURL string  `json:"callback_url"`
            Cart CheckCart  `json:"cart"`
            OrderID string  `json:"order_id"`
            Data map[string]interface{}  `json:"data"`
            AppInterceptURL string  `json:"app_intercept_url"`
            Message string  `json:"message"`
            PaymentConfirmURL string  `json:"payment_confirm_url"`
            Success bool  `json:"success"`
         
    }
    
    // CartDeliveryModesResponse used by Cart
    type CartDeliveryModesResponse struct {

        
            PickupStores []float64  `json:"pickup_stores"`
            AvailableModes []string  `json:"available_modes"`
         
    }
    
    // PickupStoreDetail used by Cart
    type PickupStoreDetail struct {

        
            Country string  `json:"country"`
            UID float64  `json:"uid"`
            Area string  `json:"area"`
            AreaCode string  `json:"area_code"`
            AreaCodeSlug string  `json:"area_code_slug"`
            City string  `json:"city"`
            ID float64  `json:"id"`
            AddressType string  `json:"address_type"`
            Email string  `json:"email"`
            Pincode float64  `json:"pincode"`
            State string  `json:"state"`
            StoreCode string  `json:"store_code"`
            Phone string  `json:"phone"`
            Name string  `json:"name"`
            Landmark string  `json:"landmark"`
            Address string  `json:"address"`
         
    }
    
    // StoreDetailsResponse used by Cart
    type StoreDetailsResponse struct {

        
            Items []PickupStoreDetail  `json:"items"`
         
    }
    
    // UpdateCartPaymentRequest used by Cart
    type UpdateCartPaymentRequest struct {

        
            PaymentIdentifier string  `json:"payment_identifier"`
            AddressID string  `json:"address_id"`
            ID string  `json:"id"`
            MerchantCode string  `json:"merchant_code"`
            AggregatorName string  `json:"aggregator_name"`
            PaymentMode string  `json:"payment_mode"`
         
    }
    
    // CouponValidity used by Cart
    type CouponValidity struct {

        
            Title string  `json:"title"`
            DisplayMessageEn string  `json:"display_message_en"`
            Code string  `json:"code"`
            Valid bool  `json:"valid"`
            Discount float64  `json:"discount"`
         
    }
    
    // PaymentCouponValidate used by Cart
    type PaymentCouponValidate struct {

        
            Success bool  `json:"success"`
            CouponValidity CouponValidity  `json:"coupon_validity"`
            Message string  `json:"message"`
         
    }
    

    
    // E used by Rewards
    type E struct {

        
            Code float64  `json:"code"`
            Exception string  `json:"exception"`
            Info string  `json:"info"`
            Message string  `json:"message"`
         
    }
    
    // GiveawayResponse used by Rewards
    type GiveawayResponse struct {

        
            Items []Giveaway  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // Giveaway used by Rewards
    type Giveaway struct {

        
            ID string  `json:"_id"`
            Schedule Schedule  `json:"_schedule"`
            Active bool  `json:"active"`
            ApplicationID string  `json:"application_id"`
            Audience RewardsAudience  `json:"audience"`
            BannerImage Asset  `json:"banner_image"`
            CreatedAt string  `json:"created_at"`
            Description string  `json:"description"`
            Name string  `json:"name"`
            Rule RewardsRule  `json:"rule"`
            Title string  `json:"title"`
            UpdatedAt string  `json:"updated_at"`
         
    }
    
    // Schedule used by Rewards
    type Schedule struct {

        
            Duration float64  `json:"duration"`
            End string  `json:"end"`
            Start string  `json:"start"`
            Cron string  `json:"cron"`
         
    }
    
    // RewardsAudience used by Rewards
    type RewardsAudience struct {

        
            HeaderUserID string  `json:"header_user_id"`
            ID string  `json:"id"`
         
    }
    
    // Asset used by Rewards
    type Asset struct {

        
            AspectRatio string  `json:"aspect_ratio"`
            ID string  `json:"id"`
            SecureURL string  `json:"secure_url"`
         
    }
    
    // RewardsRule used by Rewards
    type RewardsRule struct {

        
            Amount float64  `json:"amount"`
         
    }
    
    // Page used by Rewards
    type Page struct {

        
            ItemTotal float64  `json:"item_total"`
            NextID string  `json:"next_id"`
            HasPrevious bool  `json:"has_previous"`
            HasNext bool  `json:"has_next"`
            Current float64  `json:"current"`
            Type string  `json:"type"`
            Size float64  `json:"size"`
         
    }
    
    // Offer used by Rewards
    type Offer struct {

        
            Schedule Schedule  `json:"_schedule"`
            Active bool  `json:"active"`
            ApplicationID string  `json:"application_id"`
            BannerImage Asset  `json:"banner_image"`
            CreatedAt string  `json:"created_at"`
            Name string  `json:"name"`
            Rule map[string]interface{}  `json:"rule"`
            Share ShareMessages  `json:"share"`
            SubText string  `json:"sub_text"`
            Text string  `json:"text"`
            Type string  `json:"type"`
            UpdatedAt string  `json:"updated_at"`
            UpdatedBy string  `json:"updated_by"`
            URL string  `json:"url"`
         
    }
    
    // ShareMessages used by Rewards
    type ShareMessages struct {

        
            Email float64  `json:"email"`
            Facebook string  `json:"facebook"`
            Fallback string  `json:"fallback"`
            Message string  `json:"message"`
            Messenger string  `json:"messenger"`
            Sms string  `json:"sms"`
            Text string  `json:"text"`
            Twitter string  `json:"twitter"`
            Whatsapp string  `json:"whatsapp"`
         
    }
    
    // UserRes used by Rewards
    type UserRes struct {

        
            Points Points  `json:"points"`
            User RewardUser  `json:"user"`
         
    }
    
    // Points used by Rewards
    type Points struct {

        
            Available float64  `json:"available"`
         
    }
    
    // RewardUser used by Rewards
    type RewardUser struct {

        
            ID string  `json:"_id"`
            Active bool  `json:"active"`
            CreatedAt string  `json:"created_at"`
            Referral Referral  `json:"referral"`
            UID float64  `json:"uid"`
            UpdatedAt string  `json:"updated_at"`
            UserBlockReason string  `json:"user_block_reason"`
            UserID string  `json:"user_id"`
         
    }
    
    // Referral used by Rewards
    type Referral struct {

        
            Code string  `json:"code"`
         
    }
    
    // AppUser used by Rewards
    type AppUser struct {

        
            ID string  `json:"_id"`
            Active bool  `json:"active"`
            ApplicationID string  `json:"application_id"`
            BlockReason string  `json:"block_reason"`
            UpdatedAt string  `json:"updated_at"`
            UpdatedBy string  `json:"updated_by"`
            UserID string  `json:"user_id"`
         
    }
    
    // GiveawayAudience used by Rewards
    type GiveawayAudience struct {

        
            AudienceID string  `json:"audience_id"`
            CurrentCount float64  `json:"current_count"`
         
    }
    
    // HistoryRes used by Rewards
    type HistoryRes struct {

        
            Items []PointsHistory  `json:"items"`
            Page Page  `json:"page"`
            Points float64  `json:"points"`
         
    }
    
    // PointsHistory used by Rewards
    type PointsHistory struct {

        
            ID string  `json:"_id"`
            ApplicationID string  `json:"application_id"`
            Claimed bool  `json:"claimed"`
            CreatedAt string  `json:"created_at"`
            ExpiresOn string  `json:"expires_on"`
            Meta map[string]interface{}  `json:"meta"`
            Points float64  `json:"points"`
            RemainingPoints float64  `json:"remaining_points"`
            Text1 string  `json:"text_1"`
            Text2 string  `json:"text_2"`
            Text3 string  `json:"text_3"`
            TxnName string  `json:"txn_name"`
            UpdatedAt string  `json:"updated_at"`
            UserID string  `json:"user_id"`
         
    }
    
    // ConfigurationRes used by Rewards
    type ConfigurationRes struct {

        
            ValidAndroidPackages []string  `json:"valid_android_packages"`
            TermsConditionsLink string  `json:"terms_conditions_link"`
            ApplicationID string  `json:"application_id"`
            Success bool  `json:"success"`
         
    }
    
    // SetConfigurationRes used by Rewards
    type SetConfigurationRes struct {

        
            Success bool  `json:"success"`
         
    }
    
    // ConfigurationRequest used by Rewards
    type ConfigurationRequest struct {

        
            ValidAndroidPackages []string  `json:"valid_android_packages"`
            TermsConditionsLink string  `json:"terms_conditions_link"`
         
    }
    

    
    // StatGroup used by Analytics
    type StatGroup struct {

        
            Key string  `json:"key"`
            URL string  `json:"url"`
            Title string  `json:"title"`
         
    }
    
    // ErrorRes used by Analytics
    type ErrorRes struct {

        
            Message string  `json:"message"`
         
    }
    
    // StatsGroups used by Analytics
    type StatsGroups struct {

        
            Groups []StatGroup  `json:"groups"`
         
    }
    
    // StatsGroupComponent used by Analytics
    type StatsGroupComponent struct {

        
            Key string  `json:"key"`
            URL string  `json:"url"`
            Title string  `json:"title"`
            Type string  `json:"type"`
            Filters map[string]interface{}  `json:"filters"`
         
    }
    
    // StatsGroupComponents used by Analytics
    type StatsGroupComponents struct {

        
            Title string  `json:"title"`
            Components []StatsGroupComponent  `json:"components"`
         
    }
    
    // StatsRes used by Analytics
    type StatsRes struct {

        
            Key string  `json:"key"`
            Title string  `json:"title"`
            Type string  `json:"type"`
            Data map[string]interface{}  `json:"data"`
         
    }
    
    // ReceivedAt used by Analytics
    type ReceivedAt struct {

        
            Value string  `json:"value"`
         
    }
    
    // AbandonCartsDetail used by Analytics
    type AbandonCartsDetail struct {

        
            PropertiesCartID string  `json:"properties_cart_id"`
            ContextTraitsFirstName string  `json:"context_traits_first_name"`
            ContextTraitsLastName string  `json:"context_traits_last_name"`
            ContextTraitsPhoneNumber string  `json:"context_traits_phone_number"`
            ContextTraitsEmail string  `json:"context_traits_email"`
            ContextAppApplicationID string  `json:"context_app_application_id"`
            PropertiesBreakupValuesRawTotal string  `json:"properties_breakup_values_raw_total"`
            ReceivedAt ReceivedAt  `json:"received_at"`
         
    }
    
    // Page used by Analytics
    type Page struct {

        
            ItemTotal float64  `json:"item_total"`
            NextID string  `json:"next_id"`
            HasPrevious bool  `json:"has_previous"`
            HasNext bool  `json:"has_next"`
            Current float64  `json:"current"`
            Type string  `json:"type"`
            Size float64  `json:"size"`
         
    }
    
    // AbandonCartsList used by Analytics
    type AbandonCartsList struct {

        
            Items []AbandonCartsDetail  `json:"items"`
            CartTotal string  `json:"cart_total"`
            Page Page  `json:"page"`
         
    }
    
    // AbandonCartDetail used by Analytics
    type AbandonCartDetail struct {

        
            ID string  `json:"_id"`
            UserID string  `json:"user_id"`
            CartValue string  `json:"cart_value"`
            Articles []map[string]interface{}  `json:"articles"`
            Breakup map[string]interface{}  `json:"breakup"`
            Address map[string]interface{}  `json:"address"`
         
    }
    
    // ExportJobReq used by Analytics
    type ExportJobReq struct {

        
            MarketplaceName string  `json:"marketplace_name"`
            StartTime string  `json:"start_time"`
            EndTime string  `json:"end_time"`
            EventType string  `json:"event_type"`
            TraceID string  `json:"trace_id"`
         
    }
    
    // ExportJobRes used by Analytics
    type ExportJobRes struct {

        
            Status string  `json:"status"`
            JobID string  `json:"job_id"`
         
    }
    
    // ExportJobStatusRes used by Analytics
    type ExportJobStatusRes struct {

        
            Status string  `json:"status"`
            JobID string  `json:"job_id"`
            DownloadURL string  `json:"download_url"`
         
    }
    
    // GetLogsListReq used by Analytics
    type GetLogsListReq struct {

        
            MarketplaceName string  `json:"marketplace_name"`
            StartDate string  `json:"start_date"`
            CompanyID string  `json:"company_id"`
            EndDate string  `json:"end_date"`
         
    }
    
    // MkpLogsResp used by Analytics
    type MkpLogsResp struct {

        
            StartTimeIso string  `json:"start_time_iso"`
            EndTimeIso string  `json:"end_time_iso"`
            EventType string  `json:"event_type"`
            TraceID string  `json:"trace_id"`
            Count string  `json:"count"`
            Status string  `json:"status"`
         
    }
    
    // GetLogsListRes used by Analytics
    type GetLogsListRes struct {

        
            Items []MkpLogsResp  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // SearchLogReq used by Analytics
    type SearchLogReq struct {

        
            MarketplaceName string  `json:"marketplace_name"`
            StartDate string  `json:"start_date"`
            CompanyID string  `json:"company_id"`
            EndDate string  `json:"end_date"`
            Identifier string  `json:"identifier"`
            IdentifierValue string  `json:"identifier_value"`
         
    }
    
    // LogInfo used by Analytics
    type LogInfo struct {

        
            ID string  `json:"_id"`
            Status string  `json:"status"`
            EventType string  `json:"event_type"`
            MarketplaceName string  `json:"marketplace_name"`
            Event string  `json:"event"`
            TraceID string  `json:"trace_id"`
            CompanyID float64  `json:"company_id"`
            BrandID float64  `json:"brand_id"`
            StoreCode string  `json:"store_code"`
            StoreID float64  `json:"store_id"`
            ItemID float64  `json:"item_id"`
            ArticleID string  `json:"article_id"`
            SellerIdentifier string  `json:"seller_identifier"`
         
    }
    
    // SearchLogRes used by Analytics
    type SearchLogRes struct {

        
            Items []LogInfo  `json:"items"`
            Page Page  `json:"page"`
         
    }
    

    
    // ValidityObject used by Discount
    type ValidityObject struct {

        
            Start string  `json:"start"`
            End string  `json:"end"`
         
    }
    
    // CreateUpdateDiscount used by Discount
    type CreateUpdateDiscount struct {

        
            Name string  `json:"name"`
            CompanyID float64  `json:"company_id"`
            IsActive bool  `json:"is_active"`
            AppIds []string  `json:"app_ids"`
            ExtensionIds []string  `json:"extension_ids"`
            JobType string  `json:"job_type"`
            DiscountType string  `json:"discount_type"`
            DiscountLevel string  `json:"discount_level"`
            Value float64  `json:"value"`
            FilePath string  `json:"file_path"`
            BrandIds []float64  `json:"brand_ids"`
            StoreIds []float64  `json:"store_ids"`
            Validity ValidityObject  `json:"validity"`
         
    }
    
    // DiscountJob used by Discount
    type DiscountJob struct {

        
            ID string  `json:"_id"`
            Name string  `json:"name"`
            CompanyID float64  `json:"company_id"`
            IsActive bool  `json:"is_active"`
            AppIds []string  `json:"app_ids"`
            JobType string  `json:"job_type"`
            DiscountType string  `json:"discount_type"`
            DiscountLevel string  `json:"discount_level"`
            Value float64  `json:"value"`
            FilePath string  `json:"file_path"`
            BrandIds []float64  `json:"brand_ids"`
            StoreIds []float64  `json:"store_ids"`
            Validity ValidityObject  `json:"validity"`
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
            CreatedBy UserDetails  `json:"created_by"`
            ModifiedBy UserDetails  `json:"modified_by"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // ListOrCalender used by Discount
    type ListOrCalender struct {

        
            Items []DiscountJob  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // FileJobResponse used by Discount
    type FileJobResponse struct {

        
            Stage string  `json:"stage"`
            Total float64  `json:"total"`
            Failed float64  `json:"failed"`
            CompanyID float64  `json:"company_id"`
            Body map[string]interface{}  `json:"body"`
            Type string  `json:"type"`
            FileType string  `json:"file_type"`
         
    }
    
    // DownloadFileJob used by Discount
    type DownloadFileJob struct {

        
            BrandIds []float64  `json:"brand_ids"`
            StoreIds []float64  `json:"store_ids"`
         
    }
    
    // CancelJobResponse used by Discount
    type CancelJobResponse struct {

        
            Success bool  `json:"success"`
         
    }
    
    // Page used by Discount
    type Page struct {

        
            ItemTotal float64  `json:"item_total"`
            NextID string  `json:"next_id"`
            HasPrevious bool  `json:"has_previous"`
            HasNext bool  `json:"has_next"`
            Current float64  `json:"current"`
            Type string  `json:"type"`
            Size float64  `json:"size"`
         
    }
    
    // UserDetails used by Discount
    type UserDetails struct {

        
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
         
    }
    
    // BadRequestObject used by Discount
    type BadRequestObject struct {

        
            Message string  `json:"message"`
         
    }
    

    
    // AddProxyReq used by Partner
    type AddProxyReq struct {

        
            AttachedPath string  `json:"attached_path"`
            ProxyURL string  `json:"proxy_url"`
         
    }
    
    // AddProxyResponse used by Partner
    type AddProxyResponse struct {

        
            ID string  `json:"_id"`
            AttachedPath string  `json:"attached_path"`
            ProxyURL string  `json:"proxy_url"`
            CompanyID string  `json:"company_id"`
            ApplicationID string  `json:"application_id"`
            ExtensionID string  `json:"extension_id"`
            CreatedAt string  `json:"created_at"`
            ModifiedAt string  `json:"modified_at"`
         
    }
    
    // RemoveProxyResponse used by Partner
    type RemoveProxyResponse struct {

        
            Message string  `json:"message"`
            Data map[string]interface{}  `json:"data"`
         
    }
    
    // APIError used by Partner
    type APIError struct {

        
            Code string  `json:"code"`
            Message string  `json:"message"`
            Info string  `json:"info"`
            RequestID string  `json:"request_id"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    

    
    // EventConfig used by Webhook
    type EventConfig struct {

        
            ID float64  `json:"id"`
            EventName string  `json:"event_name"`
            EventType string  `json:"event_type"`
            EventCategory string  `json:"event_category"`
            Version string  `json:"version"`
            DisplayName string  `json:"display_name"`
            Description string  `json:"description"`
            CreatedOn string  `json:"created_on"`
         
    }
    
    // EventConfigList used by Webhook
    type EventConfigList struct {

        
            Items []EventConfig  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // EventConfigResponse used by Webhook
    type EventConfigResponse struct {

        
            EventConfigs []EventConfig  `json:"event_configs"`
         
    }
    
    // SubscriberConfigList used by Webhook
    type SubscriberConfigList struct {

        
            Items []SubscriberResponse  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // Page used by Webhook
    type Page struct {

        
            ItemTotal float64  `json:"item_total"`
            NextID string  `json:"next_id"`
            HasPrevious bool  `json:"has_previous"`
            HasNext bool  `json:"has_next"`
            Current float64  `json:"current"`
            Type string  `json:"type"`
            Size float64  `json:"size"`
         
    }
    
    // EventProcessedStatus used by Webhook
    type EventProcessedStatus struct {

        
            ID float64  `json:"id"`
            SubscriberID string  `json:"subscriber_id"`
            Attempt float64  `json:"attempt"`
            ResponseCode string  `json:"response_code"`
            ResponseMessage string  `json:"response_message"`
            CreatedOn string  `json:"created_on"`
            ProcessedOn string  `json:"processed_on"`
            Status bool  `json:"status"`
         
    }
    
    // EventPayload used by Webhook
    type EventPayload struct {

        
            ID float64  `json:"id"`
            EventTraceID string  `json:"event_trace_id"`
            MessageID string  `json:"message_id"`
            EventName string  `json:"event_name"`
            EventType string  `json:"event_type"`
            Version string  `json:"version"`
            Status bool  `json:"status"`
         
    }
    
    // SubscriberConfig used by Webhook
    type SubscriberConfig struct {

        
            ID float64  `json:"id"`
            Name string  `json:"name"`
            WebhookURL string  `json:"webhook_url"`
            Association Association  `json:"association"`
            CustomHeaders map[string]interface{}  `json:"custom_headers"`
            Status SubscriberStatus  `json:"status"`
            EmailID string  `json:"email_id"`
            AuthMeta AuthMeta  `json:"auth_meta"`
            EventID []float64  `json:"event_id"`
         
    }
    
    // SubscriberResponse used by Webhook
    type SubscriberResponse struct {

        
            ID float64  `json:"id"`
            Name string  `json:"name"`
            WebhookURL string  `json:"webhook_url"`
            Association Association  `json:"association"`
            CustomHeaders map[string]interface{}  `json:"custom_headers"`
            EmailID string  `json:"email_id"`
            Status SubscriberStatus  `json:"status"`
            AuthMeta AuthMeta  `json:"auth_meta"`
            CreatedOn string  `json:"created_on"`
            UpdatedOn string  `json:"updated_on"`
            EventConfigs []EventConfig  `json:"event_configs"`
         
    }
    
    // SubscriberEvent used by Webhook
    type SubscriberEvent struct {

        
            ID float64  `json:"id"`
            SubscriberID float64  `json:"subscriber_id"`
            EventID float64  `json:"event_id"`
            CreatedDate string  `json:"created_date"`
         
    }
    
    // AuthMeta used by Webhook
    type AuthMeta struct {

        
            Type string  `json:"type"`
            Secret string  `json:"secret"`
         
    }
    
    // Association used by Webhook
    type Association struct {

        
            CompanyID float64  `json:"company_id"`
            ApplicationID []string  `json:"application_id"`
            ExtensionID string  `json:"extension_id"`
            Criteria string  `json:"criteria"`
         
    }
    
    // EventConfigBase used by Webhook
    type EventConfigBase struct {

        
            EventName string  `json:"event_name"`
            EventType string  `json:"event_type"`
            EventCategory string  `json:"event_category"`
            Version string  `json:"version"`
         
    }
    

    
    // RequestBodyAuditLog used by AuditTrail
    type RequestBodyAuditLog struct {

        
            LogMeta LogMetaObj  `json:"log_meta"`
            LogPayload map[string]interface{}  `json:"log_payload"`
         
    }
    
    // CreateLogResponse used by AuditTrail
    type CreateLogResponse struct {

        
            Message string  `json:"message"`
            InternalMessage string  `json:"internal_message"`
         
    }
    
    // LogMetaObj used by AuditTrail
    type LogMetaObj struct {

        
            Modifier map[string]interface{}  `json:"modifier"`
            Application string  `json:"application"`
            Entity EntityObject  `json:"entity"`
            DeviceInfo map[string]interface{}  `json:"device_info"`
            Location map[string]interface{}  `json:"location"`
         
    }
    
    // EntityObject used by AuditTrail
    type EntityObject struct {

        
            ID string  `json:"id"`
            Type string  `json:"type"`
            Action string  `json:"action"`
         
    }
    
    // LogSchemaResponse used by AuditTrail
    type LogSchemaResponse struct {

        
            Docs []LogDocs  `json:"docs"`
         
    }
    
    // LogDocs used by AuditTrail
    type LogDocs struct {

        
            Entity EntityObj  `json:"entity"`
            Modifier Modifier  `json:"modifier"`
            DeviceInfo DeviceInfo  `json:"device_info"`
            Location Location  `json:"location"`
            ID string  `json:"_id"`
            Company string  `json:"company"`
            Application string  `json:"application"`
            Sessions string  `json:"sessions"`
            Date string  `json:"date"`
            Logs map[string]interface{}  `json:"logs"`
         
    }
    
    // EntityObj used by AuditTrail
    type EntityObj struct {

        
            ID string  `json:"id"`
            Type string  `json:"type"`
            Action string  `json:"action"`
            EntityDetails map[string]interface{}  `json:"entity_details"`
         
    }
    
    // Modifier used by AuditTrail
    type Modifier struct {

        
            UserID string  `json:"user_id"`
            AsAdministrator bool  `json:"as_administrator"`
            UserDetails map[string]interface{}  `json:"user_details"`
         
    }
    
    // DeviceInfo used by AuditTrail
    type DeviceInfo struct {

        
            UserAgent string  `json:"user_agent"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
         
    }
    
    // Location used by AuditTrail
    type Location struct {

        
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
         
    }
    
    // BadRequest used by AuditTrail
    type BadRequest struct {

        
            Message string  `json:"message"`
         
    }
    
    // ResourceNotFound used by AuditTrail
    type ResourceNotFound struct {

        
            Message string  `json:"message"`
         
    }
    
    // InternalServerError used by AuditTrail
    type InternalServerError struct {

        
            Message string  `json:"message"`
            Code string  `json:"code"`
         
    }
    
    // EntityTypesResponse used by AuditTrail
    type EntityTypesResponse struct {

        
            Items []EntityTypeObj  `json:"items"`
         
    }
    
    // EntityTypeObj used by AuditTrail
    type EntityTypeObj struct {

        
            EntityValue string  `json:"entity_value"`
            DisplayName string  `json:"display_name"`
         
    }
    
