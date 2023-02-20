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
    
    // LocationCountry used by Common
    type LocationCountry struct {

        
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
         
    }
    
    // Locations used by Common
    type Locations struct {

        
            Items []map[string]interface{}  `json:"items"`
         
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
            Debug Debug  `json:"debug"`
            HasOldPasswordHash bool  `json:"has_old_password_hash"`
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
            SubNavigation []SubNavigationReference  `json:"sub_navigation"`
         
    }
    
    // SubNavigationReference used by Content
    type SubNavigationReference struct {

        
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
    

    
    // UnauthenticatedUser used by Billing
    type UnauthenticatedUser struct {

        
            Message string  `json:"message"`
         
    }
    
    // UnauthenticatedApplication used by Billing
    type UnauthenticatedApplication struct {

        
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
    

    
    // PaymentGatewayConfigResponse used by Payment
    type PaymentGatewayConfigResponse struct {

        
            Success bool  `json:"success"`
            Aggregators []map[string]interface{}  `json:"aggregators"`
            DisplayFields []string  `json:"display_fields"`
            Created bool  `json:"created"`
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
            Secret string  `json:"secret"`
            Key string  `json:"key"`
            IsActive bool  `json:"is_active"`
            ConfigType string  `json:"config_type"`
         
    }
    
    // PaymentGatewayConfigRequest used by Payment
    type PaymentGatewayConfigRequest struct {

        
            IsActive bool  `json:"is_active"`
            AppID string  `json:"app_id"`
            AggregatorName PaymentGatewayConfig  `json:"aggregator_name"`
         
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

        
            Error ErrorCodeAndDescription  `json:"error"`
            Success bool  `json:"success"`
         
    }
    
    // PaymentModeLogo used by Payment
    type PaymentModeLogo struct {

        
            Large string  `json:"large"`
            Small string  `json:"small"`
         
    }
    
    // IntentApp used by Payment
    type IntentApp struct {

        
            DisplayName string  `json:"display_name"`
            Logos PaymentModeLogo  `json:"logos"`
            PackageName string  `json:"package_name"`
            Code string  `json:"code"`
         
    }
    
    // IntentAppErrorList used by Payment
    type IntentAppErrorList struct {

        
            PackageName string  `json:"package_name"`
            Code string  `json:"code"`
         
    }
    
    // PaymentModeList used by Payment
    type PaymentModeList struct {

        
            Name string  `json:"name"`
            AggregatorName string  `json:"aggregator_name"`
            CodLimit float64  `json:"cod_limit"`
            DisplayPriority float64  `json:"display_priority"`
            IntentAppErrorList []string  `json:"intent_app_error_list"`
            FyndVpa string  `json:"fynd_vpa"`
            ExpYear float64  `json:"exp_year"`
            IntentApp []IntentApp  `json:"intent_app"`
            CodLimitPerOrder float64  `json:"cod_limit_per_order"`
            CardToken string  `json:"card_token"`
            DisplayName string  `json:"display_name"`
            CardName string  `json:"card_name"`
            CardType string  `json:"card_type"`
            CardFingerprint string  `json:"card_fingerprint"`
            CardReference string  `json:"card_reference"`
            CardNumber string  `json:"card_number"`
            ExpMonth float64  `json:"exp_month"`
            CompliantWithTokenisationGuidelines bool  `json:"compliant_with_tokenisation_guidelines"`
            CardIssuer string  `json:"card_issuer"`
            Nickname string  `json:"nickname"`
            RemainingLimit float64  `json:"remaining_limit"`
            IntentFlow bool  `json:"intent_flow"`
            CardBrand string  `json:"card_brand"`
            CardBrandImage string  `json:"card_brand_image"`
            CardIsin string  `json:"card_isin"`
            Timeout float64  `json:"timeout"`
            CardID string  `json:"card_id"`
            IntentAppErrorDictList []IntentAppErrorList  `json:"intent_app_error_dict_list"`
            Expired bool  `json:"expired"`
            Code string  `json:"code"`
            RetryCount float64  `json:"retry_count"`
            MerchantCode string  `json:"merchant_code"`
            LogoURL PaymentModeLogo  `json:"logo_url"`
         
    }
    
    // RootPaymentMode used by Payment
    type RootPaymentMode struct {

        
            Name string  `json:"name"`
            AggregatorName string  `json:"aggregator_name"`
            DisplayName string  `json:"display_name"`
            AddCardEnabled bool  `json:"add_card_enabled"`
            IsPayByCardPl bool  `json:"is_pay_by_card_pl"`
            List []PaymentModeList  `json:"list"`
            DisplayPriority float64  `json:"display_priority"`
            SaveCard bool  `json:"save_card"`
            AnonymousEnable bool  `json:"anonymous_enable"`
         
    }
    
    // PaymentOptions used by Payment
    type PaymentOptions struct {

        
            PaymentOption []RootPaymentMode  `json:"payment_option"`
         
    }
    
    // PaymentOptionsResponse used by Payment
    type PaymentOptionsResponse struct {

        
            Success bool  `json:"success"`
            PaymentOptions PaymentOptions  `json:"payment_options"`
         
    }
    
    // Payout used by Payment
    type Payout struct {

        
            IsDefault bool  `json:"is_default"`
            PayoutsAggregators []map[string]interface{}  `json:"payouts_aggregators"`
            Customers map[string]interface{}  `json:"customers"`
            MoreAttributes map[string]interface{}  `json:"more_attributes"`
            TransferType string  `json:"transfer_type"`
            UniqueTransferNo map[string]interface{}  `json:"unique_transfer_no"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // PayoutsResponse used by Payment
    type PayoutsResponse struct {

        
            Items []Payout  `json:"items"`
            Success bool  `json:"success"`
         
    }
    
    // PayoutBankDetails used by Payment
    type PayoutBankDetails struct {

        
            Country string  `json:"country"`
            AccountNo string  `json:"account_no"`
            AccountHolder string  `json:"account_holder"`
            AccountType string  `json:"account_type"`
            BranchName string  `json:"branch_name"`
            BankName string  `json:"bank_name"`
            State string  `json:"state"`
            Pincode float64  `json:"pincode"`
            City string  `json:"city"`
            IfscCode string  `json:"ifsc_code"`
         
    }
    
    // PayoutRequest used by Payment
    type PayoutRequest struct {

        
            Aggregator string  `json:"aggregator"`
            BankDetails PayoutBankDetails  `json:"bank_details"`
            Users map[string]interface{}  `json:"users"`
            TransferType string  `json:"transfer_type"`
            IsActive bool  `json:"is_active"`
            UniqueExternalID string  `json:"unique_external_id"`
         
    }
    
    // PayoutResponse used by Payment
    type PayoutResponse struct {

        
            Success bool  `json:"success"`
            Aggregator string  `json:"aggregator"`
            BankDetails map[string]interface{}  `json:"bank_details"`
            Users map[string]interface{}  `json:"users"`
            Created bool  `json:"created"`
            TransferType string  `json:"transfer_type"`
            Payouts map[string]interface{}  `json:"payouts"`
            UniqueTransferNo string  `json:"unique_transfer_no"`
            IsActive bool  `json:"is_active"`
            PaymentStatus string  `json:"payment_status"`
         
    }
    
    // UpdatePayoutResponse used by Payment
    type UpdatePayoutResponse struct {

        
            IsActive bool  `json:"is_active"`
            Success bool  `json:"success"`
            IsDefault bool  `json:"is_default"`
         
    }
    
    // UpdatePayoutRequest used by Payment
    type UpdatePayoutRequest struct {

        
            IsActive bool  `json:"is_active"`
            UniqueExternalID string  `json:"unique_external_id"`
            IsDefault bool  `json:"is_default"`
         
    }
    
    // DeletePayoutResponse used by Payment
    type DeletePayoutResponse struct {

        
            Success bool  `json:"success"`
         
    }
    
    // SubscriptionPaymentMethodResponse used by Payment
    type SubscriptionPaymentMethodResponse struct {

        
            Success bool  `json:"success"`
            Data []map[string]interface{}  `json:"data"`
         
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

        
            Success bool  `json:"success"`
            Data map[string]interface{}  `json:"data"`
         
    }
    
    // RefundAccountResponse used by Payment
    type RefundAccountResponse struct {

        
            Message string  `json:"message"`
            IsVerifiedFlag bool  `json:"is_verified_flag"`
            Success bool  `json:"success"`
            Data map[string]interface{}  `json:"data"`
         
    }
    
    // NotFoundResourceError used by Payment
    type NotFoundResourceError struct {

        
            Description string  `json:"description"`
            Success bool  `json:"success"`
            Code string  `json:"code"`
         
    }
    
    // BankDetailsForOTP used by Payment
    type BankDetailsForOTP struct {

        
            AccountNo string  `json:"account_no"`
            AccountHolder string  `json:"account_holder"`
            BranchName string  `json:"branch_name"`
            BankName string  `json:"bank_name"`
            IfscCode string  `json:"ifsc_code"`
         
    }
    
    // AddBeneficiaryDetailsOTPRequest used by Payment
    type AddBeneficiaryDetailsOTPRequest struct {

        
            Details BankDetailsForOTP  `json:"details"`
            OrderID string  `json:"order_id"`
         
    }
    
    // IfscCodeResponse used by Payment
    type IfscCodeResponse struct {

        
            BranchName string  `json:"branch_name"`
            Success bool  `json:"success"`
            BankName string  `json:"bank_name"`
         
    }
    
    // OrderBeneficiaryDetails used by Payment
    type OrderBeneficiaryDetails struct {

        
            AccountHolder string  `json:"account_holder"`
            TransferMode string  `json:"transfer_mode"`
            ID float64  `json:"id"`
            BeneficiaryID string  `json:"beneficiary_id"`
            IsActive bool  `json:"is_active"`
            IfscCode string  `json:"ifsc_code"`
            DisplayName string  `json:"display_name"`
            Email string  `json:"email"`
            Subtitle string  `json:"subtitle"`
            BranchName string  `json:"branch_name"`
            Title string  `json:"title"`
            DelightsUserName string  `json:"delights_user_name"`
            ModifiedOn string  `json:"modified_on"`
            Address string  `json:"address"`
            AccountNo string  `json:"account_no"`
            CreatedOn string  `json:"created_on"`
            Mobile string  `json:"mobile"`
            Comment string  `json:"comment"`
            BankName string  `json:"bank_name"`
         
    }
    
    // OrderBeneficiaryResponse used by Payment
    type OrderBeneficiaryResponse struct {

        
            ShowBeneficiaryDetails bool  `json:"show_beneficiary_details"`
            Beneficiaries []OrderBeneficiaryDetails  `json:"beneficiaries"`
         
    }
    
    // MultiTenderPaymentMeta used by Payment
    type MultiTenderPaymentMeta struct {

        
            CurrentStatus string  `json:"current_status"`
            PaymentGateway string  `json:"payment_gateway"`
            PaymentID string  `json:"payment_id"`
            OrderID string  `json:"order_id"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
         
    }
    
    // MultiTenderPaymentMethod used by Payment
    type MultiTenderPaymentMethod struct {

        
            Name string  `json:"name"`
            Mode string  `json:"mode"`
            Meta MultiTenderPaymentMeta  `json:"meta"`
            Amount float64  `json:"amount"`
         
    }
    
    // PaymentConfirmationRequest used by Payment
    type PaymentConfirmationRequest struct {

        
            PaymentMethods []MultiTenderPaymentMethod  `json:"payment_methods"`
            OrderID string  `json:"order_id"`
         
    }
    
    // PaymentConfirmationResponse used by Payment
    type PaymentConfirmationResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
            OrderID string  `json:"order_id"`
         
    }
    
    // PlatformPaymentOptions used by Payment
    type PlatformPaymentOptions struct {

        
            CodAmountLimit float64  `json:"cod_amount_limit"`
            PaymentSelectionLock map[string]interface{}  `json:"payment_selection_lock"`
            Methods map[string]interface{}  `json:"methods"`
            CodCharges float64  `json:"cod_charges"`
            AnonymousCod bool  `json:"anonymous_cod"`
            CallbackURL map[string]interface{}  `json:"callback_url"`
            ModeOfPayment string  `json:"mode_of_payment"`
            Source string  `json:"source"`
            Enabled bool  `json:"enabled"`
         
    }
    
    // PlatfromPaymentConfig used by Payment
    type PlatfromPaymentConfig struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
            Data PlatformPaymentOptions  `json:"data"`
         
    }
    
    // UpdatePlatformPaymentConfig used by Payment
    type UpdatePlatformPaymentConfig struct {

        
            CodAmountLimit float64  `json:"cod_amount_limit"`
            PaymentSelectionLock map[string]interface{}  `json:"payment_selection_lock"`
            Methods map[string]interface{}  `json:"methods"`
            CodCharges float64  `json:"cod_charges"`
            AnonymousCod bool  `json:"anonymous_cod"`
         
    }
    
    // CODdata used by Payment
    type CODdata struct {

        
            Usages float64  `json:"usages"`
            Limit float64  `json:"limit"`
            UserID string  `json:"user_id"`
            RemainingLimit float64  `json:"remaining_limit"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // GetUserCODLimitResponse used by Payment
    type GetUserCODLimitResponse struct {

        
            Success bool  `json:"success"`
            UserCodData CODdata  `json:"user_cod_data"`
         
    }
    
    // SetCODForUserRequest used by Payment
    type SetCODForUserRequest struct {

        
            Mobileno string  `json:"mobileno"`
            IsActive bool  `json:"is_active"`
            MerchantUserID string  `json:"merchant_user_id"`
         
    }
    
    // SetCODOptionResponse used by Payment
    type SetCODOptionResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // EdcDevice used by Payment
    type EdcDevice struct {

        
            ApplicationID string  `json:"application_id"`
            DeviceTag string  `json:"device_tag"`
            MerchantStorePosCode string  `json:"merchant_store_pos_code"`
            EdcDeviceSerialNo string  `json:"edc_device_serial_no"`
            AggregatorID float64  `json:"aggregator_id"`
            TerminalSerialNo string  `json:"terminal_serial_no"`
            IsActive string  `json:"is_active"`
            StoreID float64  `json:"store_id"`
            TerminalUniqueIdentifier string  `json:"terminal_unique_identifier"`
         
    }
    
    // EdcDeviceDetailsResponse used by Payment
    type EdcDeviceDetailsResponse struct {

        
            Success bool  `json:"success"`
            Data EdcDevice  `json:"data"`
         
    }
    
    // EdcAddRequest used by Payment
    type EdcAddRequest struct {

        
            DeviceTag string  `json:"device_tag"`
            AggregatorID float64  `json:"aggregator_id"`
            TerminalSerialNo string  `json:"terminal_serial_no"`
            StoreID float64  `json:"store_id"`
            EdcSerialNo string  `json:"edc_serial_no"`
         
    }
    
    // EdcDeviceAddResponse used by Payment
    type EdcDeviceAddResponse struct {

        
            Success bool  `json:"success"`
            Data EdcDevice  `json:"data"`
         
    }
    
    // EdcUpdateRequest used by Payment
    type EdcUpdateRequest struct {

        
            DeviceTag string  `json:"device_tag"`
            MerchantStorePosCode string  `json:"merchant_store_pos_code"`
            AggregatorID float64  `json:"aggregator_id"`
            TerminalSerialNo string  `json:"terminal_serial_no"`
            IsActive string  `json:"is_active"`
            StoreID float64  `json:"store_id"`
         
    }
    
    // EdcDeviceUpdateResponse used by Payment
    type EdcDeviceUpdateResponse struct {

        
            Success bool  `json:"success"`
         
    }
    
    // EdcDeviceListResponse used by Payment
    type EdcDeviceListResponse struct {

        
            DeviceCount float64  `json:"device_count"`
            TotalPage float64  `json:"total_page"`
            Success bool  `json:"success"`
            NextPage bool  `json:"next_page"`
            InactiveDeviceCount float64  `json:"inactive_device_count"`
            DeviceList []EdcDevice  `json:"device_list"`
            TotalStores float64  `json:"total_stores"`
            CurrentPage float64  `json:"current_page"`
            ActiveDeviceCount float64  `json:"active_device_count"`
         
    }
    
    // PaymentInitializationRequest used by Payment
    type PaymentInitializationRequest struct {

        
            Currency string  `json:"currency"`
            Vpa string  `json:"vpa"`
            CustomerID string  `json:"customer_id"`
            Contact string  `json:"contact"`
            Aggregator string  `json:"aggregator"`
            Email string  `json:"email"`
            Method string  `json:"method"`
            RazorpayPaymentID string  `json:"razorpay_payment_id"`
            DeviceID string  `json:"device_id"`
            Amount float64  `json:"amount"`
            MerchantOrderID string  `json:"merchant_order_id"`
            OrderID string  `json:"order_id"`
            Timeout float64  `json:"timeout"`
         
    }
    
    // PaymentInitializationResponse used by Payment
    type PaymentInitializationResponse struct {

        
            Currency string  `json:"currency"`
            Vpa string  `json:"vpa"`
            CustomerID string  `json:"customer_id"`
            Success bool  `json:"success"`
            VirtualID string  `json:"virtual_id"`
            Aggregator string  `json:"aggregator"`
            Method string  `json:"method"`
            DeviceID string  `json:"device_id"`
            RazorpayPaymentID string  `json:"razorpay_payment_id"`
            UpiPollURL string  `json:"upi_poll_url"`
            AggregatorOrderID string  `json:"aggregator_order_id"`
            MerchantOrderID string  `json:"merchant_order_id"`
            Status string  `json:"status"`
            Amount float64  `json:"amount"`
            PollingURL string  `json:"polling_url"`
            Timeout float64  `json:"timeout"`
            BqrImage string  `json:"bqr_image"`
         
    }
    
    // PaymentStatusUpdateRequest used by Payment
    type PaymentStatusUpdateRequest struct {

        
            Currency string  `json:"currency"`
            Vpa string  `json:"vpa"`
            CustomerID string  `json:"customer_id"`
            Contact string  `json:"contact"`
            Aggregator string  `json:"aggregator"`
            Email string  `json:"email"`
            Method string  `json:"method"`
            DeviceID string  `json:"device_id"`
            Amount float64  `json:"amount"`
            MerchantOrderID string  `json:"merchant_order_id"`
            Status string  `json:"status"`
            OrderID string  `json:"order_id"`
         
    }
    
    // PaymentStatusUpdateResponse used by Payment
    type PaymentStatusUpdateResponse struct {

        
            Success bool  `json:"success"`
            AggregatorName string  `json:"aggregator_name"`
            RedirectURL string  `json:"redirect_url"`
            Status string  `json:"status"`
            Retry bool  `json:"retry"`
         
    }
    

    
    // UserDataInfo used by Order
    type UserDataInfo struct {

        
            Gender string  `json:"gender"`
            Mobile string  `json:"mobile"`
            LastName string  `json:"last_name"`
            FirstName string  `json:"first_name"`
            Email string  `json:"email"`
            Name string  `json:"name"`
            IsAnonymousUser bool  `json:"is_anonymous_user"`
            UID float64  `json:"uid"`
            AvisUserID string  `json:"avis_user_id"`
         
    }
    
    // ShipmentItemFulFillingStore used by Order
    type ShipmentItemFulFillingStore struct {

        
            ID string  `json:"id"`
            Code string  `json:"code"`
         
    }
    
    // PaymentModeInfo used by Order
    type PaymentModeInfo struct {

        
            Type string  `json:"type"`
            Logo string  `json:"logo"`
         
    }
    
    // GSTDetailsData used by Order
    type GSTDetailsData struct {

        
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            ValueOfGood float64  `json:"value_of_good"`
            GstFee float64  `json:"gst_fee"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            GstinCode string  `json:"gstin_code"`
         
    }
    
    // PlatformItem used by Order
    type PlatformItem struct {

        
            L3CategoryName string  `json:"l3_category_name"`
            ID float64  `json:"id"`
            L1Category []string  `json:"l1_category"`
            Image []string  `json:"image"`
            Code string  `json:"code"`
            CanCancel bool  `json:"can_cancel"`
            L3Category float64  `json:"l3_category"`
            Images []string  `json:"images"`
            DepartmentID float64  `json:"department_id"`
            Size string  `json:"size"`
            Color string  `json:"color"`
            Name string  `json:"name"`
            CanReturn bool  `json:"can_return"`
         
    }
    
    // Prices used by Order
    type Prices struct {

        
            AmountPaid float64  `json:"amount_paid"`
            ValueOfGood float64  `json:"value_of_good"`
            RefundCredit float64  `json:"refund_credit"`
            RefundAmount float64  `json:"refund_amount"`
            Discount float64  `json:"discount"`
            CouponValue float64  `json:"coupon_value"`
            CashbackApplied float64  `json:"cashback_applied"`
            PriceMarked float64  `json:"price_marked"`
            CodCharges float64  `json:"cod_charges"`
            PriceEffective float64  `json:"price_effective"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            PromotionEffectiveDiscount float64  `json:"promotion_effective_discount"`
            Cashback float64  `json:"cashback"`
            AmountPaidRoundoff float64  `json:"amount_paid_roundoff"`
            DeliveryCharge float64  `json:"delivery_charge"`
            FyndCredits float64  `json:"fynd_credits"`
         
    }
    
    // BagUnit used by Order
    type BagUnit struct {

        
            Gst GSTDetailsData  `json:"gst"`
            CanReturn bool  `json:"can_return"`
            BagID float64  `json:"bag_id"`
            TotalShipmentBags float64  `json:"total_shipment_bags"`
            CanCancel bool  `json:"can_cancel"`
            OrderingChannel string  `json:"ordering_channel"`
            ShipmentID string  `json:"shipment_id"`
            ItemQuantity float64  `json:"item_quantity"`
            Item PlatformItem  `json:"item"`
            Status map[string]interface{}  `json:"status"`
            Prices Prices  `json:"prices"`
         
    }
    
    // ShipmentStatus used by Order
    type ShipmentStatus struct {

        
            HexCode string  `json:"hex_code"`
            Title string  `json:"title"`
            OpsStatus string  `json:"ops_status"`
            ActualStatus string  `json:"actual_status"`
            Status string  `json:"status"`
         
    }
    
    // ShipmentItem used by Order
    type ShipmentItem struct {

        
            User UserDataInfo  `json:"user"`
            FulfillingStore ShipmentItemFulFillingStore  `json:"fulfilling_store"`
            Sla map[string]interface{}  `json:"sla"`
            PaymentMethods map[string]interface{}  `json:"payment_methods"`
            ID string  `json:"id"`
            PaymentModeInfo PaymentModeInfo  `json:"payment_mode_info"`
            CreatedAt string  `json:"created_at"`
            TotalBagsCount float64  `json:"total_bags_count"`
            Channel map[string]interface{}  `json:"channel"`
            FulfillingCentre string  `json:"fulfilling_centre"`
            Bags []BagUnit  `json:"bags"`
            ShipmentStatus ShipmentStatus  `json:"shipment_status"`
            Application map[string]interface{}  `json:"application"`
            TotalShipmentsInOrder float64  `json:"total_shipments_in_order"`
            ShipmentCreatedAt float64  `json:"shipment_created_at"`
            Prices Prices  `json:"prices"`
         
    }
    
    // FilterInfoOption used by Order
    type FilterInfoOption struct {

        
            Value string  `json:"value"`
            Text string  `json:"text"`
            Name string  `json:"name"`
         
    }
    
    // FiltersInfo used by Order
    type FiltersInfo struct {

        
            Value string  `json:"value"`
            Type string  `json:"type"`
            Text string  `json:"text"`
            Options []FilterInfoOption  `json:"options"`
         
    }
    
    // ShipmentInternalPlatformViewResponse used by Order
    type ShipmentInternalPlatformViewResponse struct {

        
            Page map[string]interface{}  `json:"page"`
            AppliedFilters map[string]interface{}  `json:"applied_filters"`
            Items []ShipmentItem  `json:"items"`
            Filters []FiltersInfo  `json:"filters"`
         
    }
    
    // Error used by Order
    type Error struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // FulfillingStore used by Order
    type FulfillingStore struct {

        
            Meta map[string]interface{}  `json:"meta"`
            City string  `json:"city"`
            ID float64  `json:"id"`
            Pincode string  `json:"pincode"`
            StoreName string  `json:"store_name"`
            FulfillmentChannel string  `json:"fulfillment_channel"`
            Code string  `json:"code"`
            Country string  `json:"country"`
            Address string  `json:"address"`
            State string  `json:"state"`
            Phone string  `json:"phone"`
            ContactPerson string  `json:"contact_person"`
         
    }
    
    // OrderDetailsData used by Order
    type OrderDetailsData struct {

        
            TaxDetails map[string]interface{}  `json:"tax_details"`
            Source string  `json:"source"`
            CodCharges string  `json:"cod_charges"`
            OrderingChannel string  `json:"ordering_channel"`
            AffiliateID string  `json:"affiliate_id"`
            OrderValue string  `json:"order_value"`
            OrderDate string  `json:"order_date"`
            OrderingChannelLogo map[string]interface{}  `json:"ordering_channel_logo"`
            FyndOrderID string  `json:"fynd_order_id"`
         
    }
    
    // TrackingList used by Order
    type TrackingList struct {

        
            Text string  `json:"text"`
            IsCurrent bool  `json:"is_current"`
            IsPassed bool  `json:"is_passed"`
            Time string  `json:"time"`
            Status string  `json:"status"`
         
    }
    
    // BagStateMapper used by Order
    type BagStateMapper struct {

        
            AppFacing bool  `json:"app_facing"`
            IsActive bool  `json:"is_active"`
            JourneyType string  `json:"journey_type"`
            StateType string  `json:"state_type"`
            DisplayName string  `json:"display_name"`
            AppDisplayName string  `json:"app_display_name"`
            AppStateName string  `json:"app_state_name"`
            BsID float64  `json:"bs_id"`
            Name string  `json:"name"`
            NotifyCustomer bool  `json:"notify_customer"`
         
    }
    
    // BagStatusHistory used by Order
    type BagStatusHistory struct {

        
            Forward bool  `json:"forward"`
            Reasons []map[string]interface{}  `json:"reasons"`
            KafkaSync bool  `json:"kafka_sync"`
            CreatedAt string  `json:"created_at"`
            DeliveryPartnerID float64  `json:"delivery_partner_id"`
            BagStateMapper BagStateMapper  `json:"bag_state_mapper"`
            DisplayName string  `json:"display_name"`
            BagID float64  `json:"bag_id"`
            AppDisplayName string  `json:"app_display_name"`
            BshID float64  `json:"bsh_id"`
            StoreID float64  `json:"store_id"`
            UpdatedAt string  `json:"updated_at"`
            ShipmentID string  `json:"shipment_id"`
            StateID float64  `json:"state_id"`
            DeliveryAwbNumber string  `json:"delivery_awb_number"`
            Status string  `json:"status"`
            StateType string  `json:"state_type"`
         
    }
    
    // DPDetailsData used by Order
    type DPDetailsData struct {

        
            AwbNo string  `json:"awb_no"`
            ID float64  `json:"id"`
            Pincode string  `json:"pincode"`
            EwayBillID string  `json:"eway_bill_id"`
            GstTag string  `json:"gst_tag"`
            TrackURL string  `json:"track_url"`
            Country string  `json:"country"`
            Name string  `json:"name"`
         
    }
    
    // UserDetailsData used by Order
    type UserDetailsData struct {

        
            City string  `json:"city"`
            Pincode string  `json:"pincode"`
            Country string  `json:"country"`
            Address string  `json:"address"`
            Email string  `json:"email"`
            Name string  `json:"name"`
            State string  `json:"state"`
            Phone string  `json:"phone"`
         
    }
    
    // OrderBagArticle used by Order
    type OrderBagArticle struct {

        
            UID string  `json:"uid"`
            ReturnConfig map[string]interface{}  `json:"return_config"`
            Identifiers map[string]interface{}  `json:"identifiers"`
         
    }
    
    // Identifier used by Order
    type Identifier struct {

        
            Ean string  `json:"ean"`
         
    }
    
    // FinancialBreakup used by Order
    type FinancialBreakup struct {

        
            ItemName string  `json:"item_name"`
            Cashback float64  `json:"cashback"`
            AmountPaid float64  `json:"amount_paid"`
            GstFee float64  `json:"gst_fee"`
            PriceEffective float64  `json:"price_effective"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            CashbackApplied float64  `json:"cashback_applied"`
            Identifiers Identifier  `json:"identifiers"`
            HsnCode string  `json:"hsn_code"`
            AmountPaidRoundoff float64  `json:"amount_paid_roundoff"`
            AddedToFyndCash bool  `json:"added_to_fynd_cash"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            RefundCredit float64  `json:"refund_credit"`
            Discount float64  `json:"discount"`
            CouponValue float64  `json:"coupon_value"`
            Size string  `json:"size"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            FyndCredits float64  `json:"fynd_credits"`
            TransferPrice float64  `json:"transfer_price"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            ValueOfGood float64  `json:"value_of_good"`
            CodCharges float64  `json:"cod_charges"`
            PriceMarked float64  `json:"price_marked"`
            PromotionEffectiveDiscount float64  `json:"promotion_effective_discount"`
            GstTag string  `json:"gst_tag"`
            TotalUnits float64  `json:"total_units"`
            DeliveryCharge float64  `json:"delivery_charge"`
         
    }
    
    // CurrentStatus used by Order
    type CurrentStatus struct {

        
            KafkaSync bool  `json:"kafka_sync"`
            CreatedAt string  `json:"created_at"`
            DeliveryPartnerID float64  `json:"delivery_partner_id"`
            BagStateMapper BagStateMapper  `json:"bag_state_mapper"`
            CurrentStatusID float64  `json:"current_status_id"`
            BagID float64  `json:"bag_id"`
            StoreID float64  `json:"store_id"`
            ShipmentID string  `json:"shipment_id"`
            UpdatedAt float64  `json:"updated_at"`
            StateID float64  `json:"state_id"`
            DeliveryAwbNumber string  `json:"delivery_awb_number"`
            Status string  `json:"status"`
            StateType string  `json:"state_type"`
         
    }
    
    // BagGST used by Order
    type BagGST struct {

        
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            ValueOfGood float64  `json:"value_of_good"`
            GstFee float64  `json:"gst_fee"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            GstTag string  `json:"gst_tag"`
            IsDefaultHsnCode bool  `json:"is_default_hsn_code"`
            HsnCode string  `json:"hsn_code"`
            GstinCode string  `json:"gstin_code"`
         
    }
    
    // OrderBrandName used by Order
    type OrderBrandName struct {

        
            CreatedOn string  `json:"created_on"`
            ID float64  `json:"id"`
            Company string  `json:"company"`
            Logo string  `json:"logo"`
            ModifiedOn string  `json:"modified_on"`
            BrandName string  `json:"brand_name"`
         
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

        
            Value float64  `json:"value"`
            Type string  `json:"type"`
         
    }
    
    // AppliedPromos used by Order
    type AppliedPromos struct {

        
            PromotionName string  `json:"promotion_name"`
            Amount float64  `json:"amount"`
            BuyRules []BuyRules  `json:"buy_rules"`
            PromoID string  `json:"promo_id"`
            ArticleQuantity float64  `json:"article_quantity"`
            PromotionType string  `json:"promotion_type"`
            DiscountRules []DiscountRules  `json:"discount_rules"`
            MrpPromotion bool  `json:"mrp_promotion"`
         
    }
    
    // BagConfigs used by Order
    type BagConfigs struct {

        
            IsActive bool  `json:"is_active"`
            IsReturnable bool  `json:"is_returnable"`
            IsCustomerReturnAllowed bool  `json:"is_customer_return_allowed"`
            EnableTracking bool  `json:"enable_tracking"`
            CanBeCancelled bool  `json:"can_be_cancelled"`
            AllowForceReturn bool  `json:"allow_force_return"`
         
    }
    
    // PlatformDeliveryAddress used by Order
    type PlatformDeliveryAddress struct {

        
            Address2 string  `json:"address2"`
            Latitude float64  `json:"latitude"`
            Landmark string  `json:"landmark"`
            City string  `json:"city"`
            Pincode string  `json:"pincode"`
            Longitude float64  `json:"longitude"`
            Version string  `json:"version"`
            Area string  `json:"area"`
            CreatedAt string  `json:"created_at"`
            AddressType string  `json:"address_type"`
            AddressCategory string  `json:"address_category"`
            Country string  `json:"country"`
            UpdatedAt string  `json:"updated_at"`
            Address1 string  `json:"address1"`
            Email string  `json:"email"`
            State string  `json:"state"`
            Phone string  `json:"phone"`
            ContactPerson string  `json:"contact_person"`
         
    }
    
    // OrderBags used by Order
    type OrderBags struct {

        
            Article OrderBagArticle  `json:"article"`
            CanCancel bool  `json:"can_cancel"`
            FinancialBreakup FinancialBreakup  `json:"financial_breakup"`
            EntityType string  `json:"entity_type"`
            LineNumber float64  `json:"line_number"`
            Item PlatformItem  `json:"item"`
            CurrentStatus CurrentStatus  `json:"current_status"`
            GstDetails BagGST  `json:"gst_details"`
            DisplayName string  `json:"display_name"`
            Brand OrderBrandName  `json:"brand"`
            Prices Prices  `json:"prices"`
            Quantity float64  `json:"quantity"`
            ParentPromoBags map[string]interface{}  `json:"parent_promo_bags"`
            SellerIdentifier string  `json:"seller_identifier"`
            AppliedPromos []AppliedPromos  `json:"applied_promos"`
            BagConfigs BagConfigs  `json:"bag_configs"`
            Identifier string  `json:"identifier"`
            DeliveryAddress PlatformDeliveryAddress  `json:"delivery_address"`
            BagID float64  `json:"bag_id"`
            CanReturn bool  `json:"can_return"`
         
    }
    
    // ShipmentPayments used by Order
    type ShipmentPayments struct {

        
            Mode string  `json:"mode"`
            Logo string  `json:"logo"`
            Source string  `json:"source"`
         
    }
    
    // ShipmentStatusData used by Order
    type ShipmentStatusData struct {

        
            ID float64  `json:"id"`
            CreatedAt string  `json:"created_at"`
            Status string  `json:"status"`
            ShipmentID string  `json:"shipment_id"`
            BagList []string  `json:"bag_list"`
         
    }
    
    // ShipmentInfoResponse used by Order
    type ShipmentInfoResponse struct {

        
            Coupon map[string]interface{}  `json:"coupon"`
            FulfillingStore FulfillingStore  `json:"fulfilling_store"`
            Order OrderDetailsData  `json:"order"`
            PlatformLogo bool  `json:"platform_logo"`
            OrderCreatedTime string  `json:"order_created_time"`
            PickedDate string  `json:"picked_date"`
            LockStatus string  `json:"lock_status"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            IsInvoiced bool  `json:"is_invoiced"`
            CurrentShipmentStatus map[string]interface{}  `json:"current_shipment_status"`
            KiranaStoreID string  `json:"kirana_store_id"`
            TrackingList []TrackingList  `json:"tracking_list"`
            UserAgent string  `json:"user_agent"`
            PriorityText string  `json:"priority_text"`
            TotalItems float64  `json:"total_items"`
            Company map[string]interface{}  `json:"company"`
            CanCancel bool  `json:"can_cancel"`
            StatusProgress float64  `json:"status_progress"`
            BankData map[string]interface{}  `json:"bank_data"`
            ForwardTrackingList []map[string]interface{}  `json:"forward_tracking_list"`
            BeneficiaryDetails bool  `json:"beneficiary_details"`
            ShipmentQuantity float64  `json:"shipment_quantity"`
            GstDetails GSTDetailsData  `json:"gst_details"`
            TotalBags float64  `json:"total_bags"`
            BagStatusHistory []BagStatusHistory  `json:"bag_status_history"`
            CanReturn bool  `json:"can_return"`
            IsFyndCoupon bool  `json:"is_fynd_coupon"`
            TrackingURL string  `json:"tracking_url"`
            Vertical string  `json:"vertical"`
            OrderingStore map[string]interface{}  `json:"ordering_store"`
            Invoice map[string]interface{}  `json:"invoice"`
            DeliveryStatus []map[string]interface{}  `json:"delivery_status"`
            EnableTracking bool  `json:"enable_tracking"`
            IsFyndStore string  `json:"is_fynd_store"`
            Escalation map[string]interface{}  `json:"escalation"`
            Items []map[string]interface{}  `json:"items"`
            PackagingType string  `json:"packaging_type"`
            ForwardShipmentStatus []map[string]interface{}  `json:"forward_shipment_status"`
            DpDetails DPDetailsData  `json:"dp_details"`
            FyndstoreEmp map[string]interface{}  `json:"fyndstore_emp"`
            EnableDpTracking bool  `json:"enable_dp_tracking"`
            IsPackagingOrder bool  `json:"is_packaging_order"`
            ForwardOrderStatus []map[string]interface{}  `json:"forward_order_status"`
            UserInfo map[string]interface{}  `json:"user_info"`
            GoGreen bool  `json:"go_green"`
            OperationalStatus string  `json:"operational_status"`
            Prices Prices  `json:"prices"`
            ReplacementDetails string  `json:"replacement_details"`
            JourneyType string  `json:"journey_type"`
            ShipmentImages []string  `json:"shipment_images"`
            DeliverySlot map[string]interface{}  `json:"delivery_slot"`
            RefundText string  `json:"refund_text"`
            BillingDetails UserDetailsData  `json:"billing_details"`
            IsPdsr string  `json:"is_pdsr"`
            UserID string  `json:"user_id"`
            CustomMeta []map[string]interface{}  `json:"custom_meta"`
            ShipmentID string  `json:"shipment_id"`
            ShipmentStatus string  `json:"shipment_status"`
            CreditNoteID string  `json:"credit_note_id"`
            SecuredDeliveryFlag string  `json:"secured_delivery_flag"`
            DueDate string  `json:"due_date"`
            DeliveryDetails UserDetailsData  `json:"delivery_details"`
            Bags []OrderBags  `json:"bags"`
            CanBreak string  `json:"can_break"`
            EmailID string  `json:"email_id"`
            OrderType string  `json:"order_type"`
            IsNotFyndSource bool  `json:"is_not_fynd_source"`
            OrderStatus map[string]interface{}  `json:"order_status"`
            ChildNodes []string  `json:"child_nodes"`
            PaymentMode string  `json:"payment_mode"`
            Mid string  `json:"mid"`
            RefundDetails map[string]interface{}  `json:"refund_details"`
            Payments ShipmentPayments  `json:"payments"`
            Status ShipmentStatusData  `json:"status"`
            PayButton string  `json:"pay_button"`
         
    }
    
    // PlatformShipment used by Order
    type PlatformShipment struct {

        
            Coupon map[string]interface{}  `json:"coupon"`
            FulfillingStore FulfillingStore  `json:"fulfilling_store"`
            Order OrderDetailsData  `json:"order"`
            PlatformLogo string  `json:"platform_logo"`
            PickedDate string  `json:"picked_date"`
            TrackingList []TrackingList  `json:"tracking_list"`
            UserAgent string  `json:"user_agent"`
            PriorityText string  `json:"priority_text"`
            TotalItems float64  `json:"total_items"`
            ShipmentQuantity float64  `json:"shipment_quantity"`
            GstDetails GSTDetailsData  `json:"gst_details"`
            TotalBags float64  `json:"total_bags"`
            BagStatusHistory []BagStatusHistory  `json:"bag_status_history"`
            Vertical string  `json:"vertical"`
            PackagingType string  `json:"packaging_type"`
            DpDetails DPDetailsData  `json:"dp_details"`
            EnableDpTracking bool  `json:"enable_dp_tracking"`
            OperationalStatus string  `json:"operational_status"`
            Prices Prices  `json:"prices"`
            JourneyType string  `json:"journey_type"`
            ShipmentImages []string  `json:"shipment_images"`
            DeliverySlot map[string]interface{}  `json:"delivery_slot"`
            BillingDetails UserDetailsData  `json:"billing_details"`
            CustomMeta []map[string]interface{}  `json:"custom_meta"`
            ShipmentID string  `json:"shipment_id"`
            ShipmentStatus string  `json:"shipment_status"`
            DeliveryDetails UserDetailsData  `json:"delivery_details"`
            Bags []OrderBags  `json:"bags"`
            PaymentMode string  `json:"payment_mode"`
            Payments ShipmentPayments  `json:"payments"`
            Status ShipmentStatusData  `json:"status"`
         
    }
    
    // OrderMeta used by Order
    type OrderMeta struct {

        
            OrderType string  `json:"order_type"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            Staff map[string]interface{}  `json:"staff"`
            Files []map[string]interface{}  `json:"files"`
            EmployeeID float64  `json:"employee_id"`
            OrderTags []map[string]interface{}  `json:"order_tags"`
            PaymentType string  `json:"payment_type"`
            OrderingStore float64  `json:"ordering_store"`
            Comment string  `json:"comment"`
            OrderPlatform string  `json:"order_platform"`
            MongoCartID float64  `json:"mongo_cart_id"`
            OrderChildEntities []string  `json:"order_child_entities"`
            CustomerNote string  `json:"customer_note"`
            CurrencySymbol string  `json:"currency_symbol"`
            CartID float64  `json:"cart_id"`
         
    }
    
    // OrderDict used by Order
    type OrderDict struct {

        
            Meta OrderMeta  `json:"meta"`
            PaymentMethods map[string]interface{}  `json:"payment_methods"`
            FyndOrderID string  `json:"fynd_order_id"`
            OrderDate string  `json:"order_date"`
            Prices Prices  `json:"prices"`
         
    }
    
    // ShipmentDetailsResponse used by Order
    type ShipmentDetailsResponse struct {

        
            Success bool  `json:"success"`
            Shipments []PlatformShipment  `json:"shipments"`
            Order OrderDict  `json:"order"`
         
    }
    
    // SubLane used by Order
    type SubLane struct {

        
            Text string  `json:"text"`
            TotalItems float64  `json:"total_items"`
            Value string  `json:"value"`
            Actions []map[string]interface{}  `json:"actions"`
            Index float64  `json:"index"`
         
    }
    
    // SuperLane used by Order
    type SuperLane struct {

        
            Value string  `json:"value"`
            Text string  `json:"text"`
            TotalItems float64  `json:"total_items"`
            Options []SubLane  `json:"options"`
         
    }
    
    // LaneConfigResponse used by Order
    type LaneConfigResponse struct {

        
            SuperLanes []SuperLane  `json:"super_lanes"`
         
    }
    
    // PlatformChannel used by Order
    type PlatformChannel struct {

        
            Name string  `json:"name"`
            Logo string  `json:"logo"`
         
    }
    
    // PlatformBreakupValues used by Order
    type PlatformBreakupValues struct {

        
            Name string  `json:"name"`
            Value string  `json:"value"`
            Display string  `json:"display"`
         
    }
    
    // PlatformOrderItems used by Order
    type PlatformOrderItems struct {

        
            OrderCreatedTime string  `json:"order_created_time"`
            OrderID string  `json:"order_id"`
            TotalOrderValue float64  `json:"total_order_value"`
            PaymentMode string  `json:"payment_mode"`
            Shipments []PlatformShipment  `json:"shipments"`
            Channel PlatformChannel  `json:"channel"`
            OrderValue float64  `json:"order_value"`
            BreakupValues []PlatformBreakupValues  `json:"breakup_values"`
            UserInfo UserDataInfo  `json:"user_info"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // OrderListingResponse used by Order
    type OrderListingResponse struct {

        
            Lane string  `json:"lane"`
            Success bool  `json:"success"`
            Items []PlatformOrderItems  `json:"items"`
            Page Page  `json:"page"`
            Message string  `json:"message"`
            TotalCount float64  `json:"total_count"`
         
    }
    
    // Options used by Order
    type Options struct {

        
            Value float64  `json:"value"`
            Text string  `json:"text"`
         
    }
    
    // MetricsCount used by Order
    type MetricsCount struct {

        
            Value float64  `json:"value"`
            Text string  `json:"text"`
            Key string  `json:"key"`
            Options []Options  `json:"options"`
         
    }
    
    // MetricCountResponse used by Order
    type MetricCountResponse struct {

        
            Items []MetricsCount  `json:"items"`
         
    }
    
    // PlatformTrack used by Order
    type PlatformTrack struct {

        
            Meta map[string]interface{}  `json:"meta"`
            LastLocationRecievedAt string  `json:"last_location_recieved_at"`
            Reason string  `json:"reason"`
            UpdatedTime string  `json:"updated_time"`
            AccountName string  `json:"account_name"`
            ShipmentType string  `json:"shipment_type"`
            UpdatedAt string  `json:"updated_at"`
            Status string  `json:"status"`
            Awb string  `json:"awb"`
            RawStatus string  `json:"raw_status"`
         
    }
    
    // PlatformShipmentTrack used by Order
    type PlatformShipmentTrack struct {

        
            Results []PlatformTrack  `json:"results"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // AdvanceFilterInfo used by Order
    type AdvanceFilterInfo struct {

        
            Filters []FiltersInfo  `json:"filters"`
            ActionCentre []FiltersInfo  `json:"action_centre"`
            Unfulfilled []FiltersInfo  `json:"unfulfilled"`
            Processed []FiltersInfo  `json:"processed"`
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

        
            S3Key string  `json:"s3_key"`
            ReportRequestedAt string  `json:"report_requested_at"`
            DisplayName string  `json:"display_name"`
            ReportID string  `json:"report_id"`
            RequestDetails map[string]interface{}  `json:"request_details"`
            ReportName string  `json:"report_name"`
            ReportType string  `json:"report_type"`
            Status string  `json:"status"`
            ReportCreatedAt string  `json:"report_created_at"`
         
    }
    
    // JioCodeUpsertDataSet used by Order
    type JioCodeUpsertDataSet struct {

        
            JioCode string  `json:"jio_code"`
            CompanyID string  `json:"company_id"`
            ItemID string  `json:"item_id"`
            ArticleID string  `json:"article_id"`
         
    }
    
    // JioCodeUpsertPayload used by Order
    type JioCodeUpsertPayload struct {

        
            Data []JioCodeUpsertDataSet  `json:"data"`
         
    }
    
    // NestedErrorSchemaDataSet used by Order
    type NestedErrorSchemaDataSet struct {

        
            Value string  `json:"value"`
            Type string  `json:"type"`
            Message string  `json:"message"`
         
    }
    
    // JioCodeUpsertResponse used by Order
    type JioCodeUpsertResponse struct {

        
            Error []NestedErrorSchemaDataSet  `json:"error"`
            Data []map[string]interface{}  `json:"data"`
            TraceID string  `json:"trace_id"`
            Identifier string  `json:"identifier"`
            Success bool  `json:"success"`
         
    }
    
    // BulkInvoicingResponse used by Order
    type BulkInvoicingResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // BulkInvoiceLabelResponse used by Order
    type BulkInvoiceLabelResponse struct {

        
            CompanyID string  `json:"company_id"`
            InvoiceStatus string  `json:"invoice_status"`
            BatchID string  `json:"batch_id"`
            Label map[string]interface{}  `json:"label"`
            Data map[string]interface{}  `json:"data"`
            Invoice map[string]interface{}  `json:"invoice"`
            StoreName string  `json:"store_name"`
            StoreCode string  `json:"store_code"`
            StoreID string  `json:"store_id"`
            DoInvoiceLabelGenerated bool  `json:"do_invoice_label_generated"`
         
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

        
            FileName string  `json:"file_name"`
            Namespace string  `json:"namespace"`
            ContentType string  `json:"content_type"`
            Operation string  `json:"operation"`
            Size float64  `json:"size"`
            Upload FileUploadResponse  `json:"upload"`
            Method string  `json:"method"`
            Cdn URL  `json:"cdn"`
            Tags []string  `json:"tags"`
            FilePath string  `json:"file_path"`
         
    }
    
    // bulkListingData used by Order
    type bulkListingData struct {

        
            ProcessingShipments []string  `json:"processing_shipments"`
            FileName string  `json:"file_name"`
            ID string  `json:"id"`
            UserName string  `json:"user_name"`
            SuccessfulShipments []map[string]interface{}  `json:"successful_shipments"`
            CompanyID float64  `json:"company_id"`
            Failed float64  `json:"failed"`
            Total float64  `json:"total"`
            StoreCode string  `json:"store_code"`
            UploadedOn string  `json:"uploaded_on"`
            Processing float64  `json:"processing"`
            Successful float64  `json:"successful"`
            StoreName string  `json:"store_name"`
            UserID string  `json:"user_id"`
            BatchID string  `json:"batch_id"`
            ExcelURL string  `json:"excel_url"`
            StoreID float64  `json:"store_id"`
            FailedShipments []map[string]interface{}  `json:"failed_shipments"`
            Status string  `json:"status"`
         
    }
    
    // BulkListingPage used by Order
    type BulkListingPage struct {

        
            HasPrevious bool  `json:"has_previous"`
            Total float64  `json:"total"`
            Size float64  `json:"size"`
            Type string  `json:"type"`
            HasNext bool  `json:"has_next"`
            Current float64  `json:"current"`
         
    }
    
    // BulkListingResponse used by Order
    type BulkListingResponse struct {

        
            Data []bulkListingData  `json:"data"`
            Success bool  `json:"success"`
            Error string  `json:"error"`
            Page BulkListingPage  `json:"page"`
         
    }
    
    // QuestionSet used by Order
    type QuestionSet struct {

        
            ID float64  `json:"id"`
            DisplayName string  `json:"display_name"`
         
    }
    
    // Reason used by Order
    type Reason struct {

        
            ID float64  `json:"id"`
            DisplayName string  `json:"display_name"`
            QcType []string  `json:"qc_type"`
            QuestionSet []QuestionSet  `json:"question_set"`
         
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

        
            Status bool  `json:"status"`
            Message string  `json:"message"`
         
    }
    
    // BulkActionDetailsDataField used by Order
    type BulkActionDetailsDataField struct {

        
            TotalShipmentsCount float64  `json:"total_shipments_count"`
            CompanyID string  `json:"company_id"`
            FailedShipmentsCount float64  `json:"failed_shipments_count"`
            BatchID string  `json:"batch_id"`
            SuccessfulShipmentIds []string  `json:"successful_shipment_ids"`
            SuccessfulShipmentsCount float64  `json:"successful_shipments_count"`
            ProcessingShipmentsCount float64  `json:"processing_shipments_count"`
         
    }
    
    // BulkActionDetailsResponse used by Order
    type BulkActionDetailsResponse struct {

        
            Error []string  `json:"error"`
            UploadedBy string  `json:"uploaded_by"`
            Data []BulkActionDetailsDataField  `json:"data"`
            FailedRecords []string  `json:"failed_records"`
            UserID string  `json:"user_id"`
            Success string  `json:"success"`
            UploadedOn string  `json:"uploaded_on"`
            Message string  `json:"message"`
            Status bool  `json:"status"`
         
    }
    
    // ReturnConfig used by Order
    type ReturnConfig struct {

        
            Returnable bool  `json:"returnable"`
            Time float64  `json:"time"`
            Unit string  `json:"unit"`
         
    }
    
    // Dimensions used by Order
    type Dimensions struct {

        
            Length float64  `json:"length"`
            Width float64  `json:"width"`
            IsDefault bool  `json:"is_default"`
            Height float64  `json:"height"`
            Unit string  `json:"unit"`
         
    }
    
    // Weight used by Order
    type Weight struct {

        
            IsDefault bool  `json:"is_default"`
            Shipping float64  `json:"shipping"`
            Unit string  `json:"unit"`
         
    }
    
    // Article used by Order
    type Article struct {

        
            SellerIdentifier string  `json:"seller_identifier"`
            ChildDetails map[string]interface{}  `json:"child_details"`
            EspModified interface{}  `json:"esp_modified"`
            ASet map[string]interface{}  `json:"a_set"`
            Size string  `json:"size"`
            Identifiers Identifier  `json:"identifiers"`
            Code string  `json:"code"`
            ReturnConfig ReturnConfig  `json:"return_config"`
            RawMeta interface{}  `json:"raw_meta"`
            ID string  `json:"_id"`
            Dimensions Dimensions  `json:"dimensions"`
            Weight Weight  `json:"weight"`
            UID string  `json:"uid"`
            IsSet bool  `json:"is_set"`
         
    }
    
    // Attributes used by Order
    type Attributes struct {

        
            MarketerName string  `json:"marketer_name"`
            PrimaryColorHex string  `json:"primary_color_hex"`
            Gender []string  `json:"gender"`
            PrimaryMaterial string  `json:"primary_material"`
            PrimaryColor string  `json:"primary_color"`
            Essential string  `json:"essential"`
            MarketerAddress string  `json:"marketer_address"`
            Name string  `json:"name"`
            BrandName string  `json:"brand_name"`
         
    }
    
    // Item used by Order
    type Item struct {

        
            BranchURL string  `json:"branch_url"`
            CanCancel bool  `json:"can_cancel"`
            ItemID float64  `json:"item_id"`
            Meta map[string]interface{}  `json:"meta"`
            L2Category []string  `json:"l2_category"`
            L2CategoryID float64  `json:"l2_category_id"`
            Image []string  `json:"image"`
            Gender string  `json:"gender"`
            DepartmentID float64  `json:"department_id"`
            LastUpdatedAt string  `json:"last_updated_at"`
            Color string  `json:"color"`
            Brand string  `json:"brand"`
            BrandID float64  `json:"brand_id"`
            Size string  `json:"size"`
            SlugKey string  `json:"slug_key"`
            L3CategoryName string  `json:"l3_category_name"`
            Attributes Attributes  `json:"attributes"`
            WebstoreProductURL string  `json:"webstore_product_url"`
            L1Category []string  `json:"l1_category"`
            L1CategoryID float64  `json:"l1_category_id"`
            Code string  `json:"code"`
            L3Category float64  `json:"l3_category"`
            Name string  `json:"name"`
            CanReturn bool  `json:"can_return"`
         
    }
    
    // B2BPODetails used by Order
    type B2BPODetails struct {

        
            DockerNumber string  `json:"docker_number"`
            PoLineAmount float64  `json:"po_line_amount"`
            PartialCanRet bool  `json:"partial_can_ret"`
            TotalGstPercentage float64  `json:"total_gst_percentage"`
            PoTaxAmount float64  `json:"po_tax_amount"`
            ItemBasePrice float64  `json:"item_base_price"`
         
    }
    
    // BagMeta used by Order
    type BagMeta struct {

        
            B2bPoDetails B2BPODetails  `json:"b2b_po_details"`
         
    }
    
    // BagGSTDetails used by Order
    type BagGSTDetails struct {

        
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            ValueOfGood float64  `json:"value_of_good"`
            GstFee float64  `json:"gst_fee"`
            SgstGstFee string  `json:"sgst_gst_fee"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            CgstGstFee string  `json:"cgst_gst_fee"`
            HsnCodeID string  `json:"hsn_code_id"`
            SgstTaxPercentage float64  `json:"sgst_tax_percentage"`
            GstTag string  `json:"gst_tag"`
            CgstTaxPercentage float64  `json:"cgst_tax_percentage"`
            IsDefaultHsnCode bool  `json:"is_default_hsn_code"`
            HsnCode string  `json:"hsn_code"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            IgstTaxPercentage float64  `json:"igst_tax_percentage"`
            IgstGstFee string  `json:"igst_gst_fee"`
            GstinCode string  `json:"gstin_code"`
         
    }
    
    // StoreEwaybill used by Order
    type StoreEwaybill struct {

        
            Enabled bool  `json:"enabled"`
         
    }
    
    // StoreEinvoice used by Order
    type StoreEinvoice struct {

        
            Enabled bool  `json:"enabled"`
            Username string  `json:"username"`
            User string  `json:"user"`
            Password string  `json:"password"`
         
    }
    
    // StoreGstCredentials used by Order
    type StoreGstCredentials struct {

        
            EWaybill StoreEwaybill  `json:"e_waybill"`
            EInvoice StoreEinvoice  `json:"e_invoice"`
         
    }
    
    // Document used by Order
    type Document struct {

        
            Verified bool  `json:"verified"`
            URL string  `json:"url"`
            Value string  `json:"value"`
            DsType string  `json:"ds_type"`
            LegalName string  `json:"legal_name"`
         
    }
    
    // StoreDocuments used by Order
    type StoreDocuments struct {

        
            Gst Document  `json:"gst"`
         
    }
    
    // EInvoicePortalDetails used by Order
    type EInvoicePortalDetails struct {

        
            User string  `json:"user"`
            Username string  `json:"username"`
            Password string  `json:"password"`
         
    }
    
    // StoreMeta used by Order
    type StoreMeta struct {

        
            AdditionalContactDetails map[string]interface{}  `json:"additional_contact_details"`
            GstCredentials StoreGstCredentials  `json:"gst_credentials"`
            EwaybillPortalDetails map[string]interface{}  `json:"ewaybill_portal_details"`
            DisplayName string  `json:"display_name"`
            Documents StoreDocuments  `json:"documents"`
            GstNumber string  `json:"gst_number"`
            ProductReturnConfig map[string]interface{}  `json:"product_return_config"`
            Stage string  `json:"stage"`
            Timing []map[string]interface{}  `json:"timing"`
            EinvoicePortalDetails EInvoicePortalDetails  `json:"einvoice_portal_details"`
            NotificationEmails []string  `json:"notification_emails"`
         
    }
    
    // StoreAddress used by Order
    type StoreAddress struct {

        
            Pincode float64  `json:"pincode"`
            CreatedAt string  `json:"created_at"`
            Area string  `json:"area"`
            Address1 string  `json:"address1"`
            Email string  `json:"email"`
            ContactPerson string  `json:"contact_person"`
            Landmark string  `json:"landmark"`
            AddressType string  `json:"address_type"`
            CountryCode string  `json:"country_code"`
            AddressCategory string  `json:"address_category"`
            Country string  `json:"country"`
            State string  `json:"state"`
            Address2 string  `json:"address2"`
            Latitude float64  `json:"latitude"`
            City string  `json:"city"`
            Longitude float64  `json:"longitude"`
            Version string  `json:"version"`
            UpdatedAt string  `json:"updated_at"`
            Phone string  `json:"phone"`
         
    }
    
    // Store used by Order
    type Store struct {

        
            CreatedAt string  `json:"created_at"`
            Pincode string  `json:"pincode"`
            FulfillmentChannel string  `json:"fulfillment_channel"`
            Address1 string  `json:"address1"`
            StoreEmail string  `json:"store_email"`
            Meta StoreMeta  `json:"meta"`
            ContactPerson string  `json:"contact_person"`
            OrderIntegrationID string  `json:"order_integration_id"`
            IsArchived bool  `json:"is_archived"`
            CompanyID float64  `json:"company_id"`
            StoreActiveFrom string  `json:"store_active_from"`
            AlohomoraUserID float64  `json:"alohomora_user_id"`
            StoreAddressJson StoreAddress  `json:"store_address_json"`
            IsActive bool  `json:"is_active"`
            IsEnabledForRecon bool  `json:"is_enabled_for_recon"`
            ParentStoreID float64  `json:"parent_store_id"`
            BrandID interface{}  `json:"brand_id"`
            Country string  `json:"country"`
            State string  `json:"state"`
            SID string  `json:"s_id"`
            LocationType string  `json:"location_type"`
            MallName string  `json:"mall_name"`
            Address2 string  `json:"address2"`
            Latitude float64  `json:"latitude"`
            City string  `json:"city"`
            Longitude float64  `json:"longitude"`
            Code string  `json:"code"`
            BrandStoreTags []string  `json:"brand_store_tags"`
            LoginUsername string  `json:"login_username"`
            PackagingMaterialCount float64  `json:"packaging_material_count"`
            UpdatedAt string  `json:"updated_at"`
            Name string  `json:"name"`
            Phone float64  `json:"phone"`
            VatNo string  `json:"vat_no"`
            MallArea string  `json:"mall_area"`
         
    }
    
    // ArticleDetails used by Order
    type ArticleDetails struct {

        
            Status map[string]interface{}  `json:"status"`
         
    }
    
    // Brand used by Order
    type Brand struct {

        
            ScriptLastRan string  `json:"script_last_ran"`
            CreatedOn float64  `json:"created_on"`
            BrandID float64  `json:"brand_id"`
            ModifiedOn float64  `json:"modified_on"`
            PickupLocation string  `json:"pickup_location"`
            CreditNoteExpiryDays float64  `json:"credit_note_expiry_days"`
            Company string  `json:"company"`
            CreditNoteAllowed bool  `json:"credit_note_allowed"`
            StartDate string  `json:"start_date"`
            Logo string  `json:"logo"`
            IsVirtualInvoice bool  `json:"is_virtual_invoice"`
            BrandName string  `json:"brand_name"`
            InvoicePrefix string  `json:"invoice_prefix"`
         
    }
    
    // Dates used by Order
    type Dates struct {

        
            DeliveryDate interface{}  `json:"delivery_date"`
            OrderCreated string  `json:"order_created"`
         
    }
    
    // AffiliateMeta used by Order
    type AffiliateMeta struct {

        
            CouponCode string  `json:"coupon_code"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            ChannelShipmentID string  `json:"channel_shipment_id"`
            EmployeeDiscount float64  `json:"employee_discount"`
            Quantity float64  `json:"quantity"`
            SizeLevelTotalQty float64  `json:"size_level_total_qty"`
            ChannelOrderID string  `json:"channel_order_id"`
            OrderItemID string  `json:"order_item_id"`
            BoxType string  `json:"box_type"`
            DueDate string  `json:"due_date"`
            IsPriority bool  `json:"is_priority"`
         
    }
    
    // AffiliateBagDetails used by Order
    type AffiliateBagDetails struct {

        
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            EmployeeDiscount float64  `json:"employee_discount"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
            AffiliateMeta AffiliateMeta  `json:"affiliate_meta"`
         
    }
    
    // PDFLinks used by Order
    type PDFLinks struct {

        
            LabelA6 string  `json:"label_a6"`
            PoInvoice string  `json:"po_invoice"`
            LabelType string  `json:"label_type"`
            Label string  `json:"label"`
            InvoiceType string  `json:"invoice_type"`
            InvoicePos string  `json:"invoice_pos"`
            LabelA4 string  `json:"label_a4"`
            CreditNoteURL string  `json:"credit_note_url"`
            Invoice string  `json:"invoice"`
            LabelPos string  `json:"label_pos"`
            B2b string  `json:"b2b"`
            InvoiceA6 string  `json:"invoice_a6"`
            InvoiceA4 string  `json:"invoice_a4"`
         
    }
    
    // EInvoice used by Order
    type EInvoice struct {

        
            ErrorMessage string  `json:"error_message"`
            SignedInvoice string  `json:"signed_invoice"`
            ErrorCode string  `json:"error_code"`
            AcknowledgeNo float64  `json:"acknowledge_no"`
            SignedQrCode string  `json:"signed_qr_code"`
            Irn string  `json:"irn"`
            AcknowledgeDate string  `json:"acknowledge_date"`
         
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
    
    // ShipmentTimeStamp used by Order
    type ShipmentTimeStamp struct {

        
            TMin string  `json:"t_min"`
            TMax string  `json:"t_max"`
         
    }
    
    // BuyerDetails used by Order
    type BuyerDetails struct {

        
            City string  `json:"city"`
            Pincode float64  `json:"pincode"`
            Address string  `json:"address"`
            Gstin string  `json:"gstin"`
            Name string  `json:"name"`
            AjioSiteID string  `json:"ajio_site_id"`
            State string  `json:"state"`
         
    }
    
    // LockData used by Order
    type LockData struct {

        
            Mto bool  `json:"mto"`
            Locked bool  `json:"locked"`
            LockMessage string  `json:"lock_message"`
         
    }
    
    // Formatted used by Order
    type Formatted struct {

        
            FMin string  `json:"f_min"`
            FMax string  `json:"f_max"`
         
    }
    
    // ShipmentMeta used by Order
    type ShipmentMeta struct {

        
            External map[string]interface{}  `json:"external"`
            ReturnAwbNumber string  `json:"return_awb_number"`
            AwbNumber string  `json:"awb_number"`
            ReturnAffiliateShipmentID string  `json:"return_affiliate_shipment_id"`
            BoxType string  `json:"box_type"`
            BagWeight map[string]interface{}  `json:"bag_weight"`
            ShipmentVolumetricWeight float64  `json:"shipment_volumetric_weight"`
            EinvoiceInfo EinvoiceInfo  `json:"einvoice_info"`
            EwaybillInfo map[string]interface{}  `json:"ewaybill_info"`
            ReturnStoreNode float64  `json:"return_store_node"`
            StoreInvoiceUpdatedDate string  `json:"store_invoice_updated_date"`
            ShipmentWeight float64  `json:"shipment_weight"`
            DebugInfo DebugInfo  `json:"debug_info"`
            Timestamp ShipmentTimeStamp  `json:"timestamp"`
            SameStoreAvailable bool  `json:"same_store_available"`
            PackagingName string  `json:"packaging_name"`
            ReturnDetails map[string]interface{}  `json:"return_details"`
            MarketplaceStoreID string  `json:"marketplace_store_id"`
            AssignDpFromSb bool  `json:"assign_dp_from_sb"`
            Weight float64  `json:"weight"`
            DueDate string  `json:"due_date"`
            B2bBuyerDetails BuyerDetails  `json:"b2b_buyer_details"`
            FulfilmentPriorityText string  `json:"fulfilment_priority_text"`
            LockData LockData  `json:"lock_data"`
            OrderType string  `json:"order_type"`
            Formatted Formatted  `json:"formatted"`
            ForwardAffiliateOrderID string  `json:"forward_affiliate_order_id"`
            DpID string  `json:"dp_id"`
            DpOptions map[string]interface{}  `json:"dp_options"`
            PoNumber string  `json:"po_number"`
            ReturnAffiliateOrderID string  `json:"return_affiliate_order_id"`
            DpSortKey string  `json:"dp_sort_key"`
            B2cBuyerDetails map[string]interface{}  `json:"b2c_buyer_details"`
            ForwardAffiliateShipmentID string  `json:"forward_affiliate_shipment_id"`
            AutoTriggerDpAssignmentAcf bool  `json:"auto_trigger_dp_assignment_acf"`
            DpName string  `json:"dp_name"`
         
    }
    
    // AffiliateDetails used by Order
    type AffiliateDetails struct {

        
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            AdID string  `json:"ad_id"`
            PdfLinks PDFLinks  `json:"pdf_links"`
            CompanyAffiliateTag string  `json:"company_affiliate_tag"`
            AffiliateID string  `json:"affiliate_id"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            AffiliateStoreID string  `json:"affiliate_store_id"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
            ShipmentMeta ShipmentMeta  `json:"shipment_meta"`
            AffiliateMeta AffiliateMeta  `json:"affiliate_meta"`
         
    }
    
    // BagReturnableCancelableStatus used by Order
    type BagReturnableCancelableStatus struct {

        
            IsActive bool  `json:"is_active"`
            IsReturnable bool  `json:"is_returnable"`
            EnableTracking bool  `json:"enable_tracking"`
            IsCustomerReturnAllowed bool  `json:"is_customer_return_allowed"`
            CanBeCancelled bool  `json:"can_be_cancelled"`
         
    }
    
    // BagDetailsPlatformResponse used by Order
    type BagDetailsPlatformResponse struct {

        
            BagUpdateTime float64  `json:"bag_update_time"`
            Article Article  `json:"article"`
            OriginalBagList []float64  `json:"original_bag_list"`
            FinancialBreakup []FinancialBreakup  `json:"financial_breakup"`
            EntityType string  `json:"entity_type"`
            LineNumber float64  `json:"line_number"`
            BType string  `json:"b_type"`
            NoOfBagsOrder float64  `json:"no_of_bags_order"`
            Item Item  `json:"item"`
            Meta BagMeta  `json:"meta"`
            OrderIntegrationID string  `json:"order_integration_id"`
            CurrentStatus BagStatusHistory  `json:"current_status"`
            GstDetails BagGSTDetails  `json:"gst_details"`
            BagStatusHistory BagStatusHistory  `json:"bag_status_history"`
            BID float64  `json:"b_id"`
            DisplayName string  `json:"display_name"`
            OrderingStore Store  `json:"ordering_store"`
            QcRequired interface{}  `json:"qc_required"`
            ArticleDetails ArticleDetails  `json:"article_details"`
            RestorePromos map[string]interface{}  `json:"restore_promos"`
            Prices Prices  `json:"prices"`
            Brand Brand  `json:"brand"`
            Dates Dates  `json:"dates"`
            OperationalStatus string  `json:"operational_status"`
            JourneyType string  `json:"journey_type"`
            Quantity float64  `json:"quantity"`
            RestoreCoupon bool  `json:"restore_coupon"`
            ShipmentID string  `json:"shipment_id"`
            CurrentOperationalStatus BagStatusHistory  `json:"current_operational_status"`
            BagStatus []BagStatusHistory  `json:"bag_status"`
            Tags []string  `json:"tags"`
            ParentPromoBags map[string]interface{}  `json:"parent_promo_bags"`
            SellerIdentifier string  `json:"seller_identifier"`
            AffiliateBagDetails AffiliateBagDetails  `json:"affiliate_bag_details"`
            AppliedPromos []map[string]interface{}  `json:"applied_promos"`
            Reasons []map[string]interface{}  `json:"reasons"`
            Identifier string  `json:"identifier"`
            AffiliateDetails AffiliateDetails  `json:"affiliate_details"`
            Status BagReturnableCancelableStatus  `json:"status"`
         
    }
    
    // ErrorResponse used by Order
    type ErrorResponse struct {

        
            Error string  `json:"error"`
            Message string  `json:"message"`
         
    }
    
    // Page1 used by Order
    type Page1 struct {

        
            ItemTotal float64  `json:"item_total"`
            Size float64  `json:"size"`
            PageType string  `json:"page_type"`
            HasNext bool  `json:"has_next"`
            Current float64  `json:"current"`
         
    }
    
    // GetBagsPlatformResponse used by Order
    type GetBagsPlatformResponse struct {

        
            Items []BagDetailsPlatformResponse  `json:"items"`
            Page Page1  `json:"page"`
         
    }
    
    // InvalidateShipmentCachePayload used by Order
    type InvalidateShipmentCachePayload struct {

        
            ShipmentIds []string  `json:"shipment_ids"`
         
    }
    
    // InvalidateShipmentCacheNestedResponse used by Order
    type InvalidateShipmentCacheNestedResponse struct {

        
            Error string  `json:"error"`
            Message string  `json:"message"`
            ShipmentID string  `json:"shipment_id"`
            Status float64  `json:"status"`
         
    }
    
    // InvalidateShipmentCacheResponse used by Order
    type InvalidateShipmentCacheResponse struct {

        
            Response []InvalidateShipmentCacheNestedResponse  `json:"response"`
         
    }
    
    // ErrorResponse1 used by Order
    type ErrorResponse1 struct {

        
            Message string  `json:"message"`
            ErrorTrace string  `json:"error_trace"`
            Status float64  `json:"status"`
         
    }
    
    // StoreReassign used by Order
    type StoreReassign struct {

        
            ItemID string  `json:"item_id"`
            BagID float64  `json:"bag_id"`
            AffiliateID string  `json:"affiliate_id"`
            StoreID float64  `json:"store_id"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            MongoArticleID string  `json:"mongo_article_id"`
            ReasonIds []float64  `json:"reason_ids"`
            FyndOrderID string  `json:"fynd_order_id"`
            SetID string  `json:"set_id"`
         
    }
    
    // StoreReassignResponse used by Order
    type StoreReassignResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // Entities used by Order
    type Entities struct {

        
            AffiliateID string  `json:"affiliate_id"`
            ID string  `json:"id"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            ReasonText string  `json:"reason_text"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
         
    }
    
    // UpdateShipmentLockPayload used by Order
    type UpdateShipmentLockPayload struct {

        
            EntityType string  `json:"entity_type"`
            Action string  `json:"action"`
            ActionType string  `json:"action_type"`
            Entities []Entities  `json:"entities"`
         
    }
    
    // Bags used by Order
    type Bags struct {

        
            AffiliateBagID string  `json:"affiliate_bag_id"`
            BagID float64  `json:"bag_id"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            IsLocked bool  `json:"is_locked"`
         
    }
    
    // OriginalFilter used by Order
    type OriginalFilter struct {

        
            AffiliateID string  `json:"affiliate_id"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
         
    }
    
    // CheckResponse used by Order
    type CheckResponse struct {

        
            AffiliateID string  `json:"affiliate_id"`
            Status string  `json:"status"`
            ShipmentID string  `json:"shipment_id"`
            Bags []Bags  `json:"bags"`
            IsBagLocked bool  `json:"is_bag_locked"`
            OriginalFilter OriginalFilter  `json:"original_filter"`
            IsShipmentLocked bool  `json:"is_shipment_locked"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            LockStatus bool  `json:"lock_status"`
         
    }
    
    // UpdateShipmentLockResponse used by Order
    type UpdateShipmentLockResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
            CheckResponse []CheckResponse  `json:"check_response"`
         
    }
    
    // AnnouncementResponse used by Order
    type AnnouncementResponse struct {

        
            LogoURL string  `json:"logo_url"`
            ID float64  `json:"id"`
            PlatformName string  `json:"platform_name"`
            CompanyID float64  `json:"company_id"`
            ToDatetime string  `json:"to_datetime"`
            Title string  `json:"title"`
            Description string  `json:"description"`
            FromDatetime string  `json:"from_datetime"`
            PlatformID string  `json:"platform_id"`
            CreatedAt string  `json:"created_at"`
         
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
    
    // ProductsDataUpdatesFilters used by Order
    type ProductsDataUpdatesFilters struct {

        
            Identifier string  `json:"identifier"`
            LineNumber float64  `json:"line_number"`
         
    }
    
    // ProductsDataUpdates used by Order
    type ProductsDataUpdates struct {

        
            Filters []ProductsDataUpdatesFilters  `json:"filters"`
            Data map[string]interface{}  `json:"data"`
         
    }
    
    // EntitiesDataUpdates used by Order
    type EntitiesDataUpdates struct {

        
            Filters []map[string]interface{}  `json:"filters"`
            Data map[string]interface{}  `json:"data"`
         
    }
    
    // DataUpdates used by Order
    type DataUpdates struct {

        
            Products []ProductsDataUpdates  `json:"products"`
            Entities []EntitiesDataUpdates  `json:"entities"`
         
    }
    
    // Products used by Order
    type Products struct {

        
            Identifier string  `json:"identifier"`
            Quantity float64  `json:"quantity"`
            LineNumber float64  `json:"line_number"`
         
    }
    
    // ProductsReasonsFilters used by Order
    type ProductsReasonsFilters struct {

        
            Identifier string  `json:"identifier"`
            Quantity float64  `json:"quantity"`
            LineNumber float64  `json:"line_number"`
         
    }
    
    // ProductsReasonsData used by Order
    type ProductsReasonsData struct {

        
            ReasonText string  `json:"reason_text"`
            ReasonID float64  `json:"reason_id"`
         
    }
    
    // ProductsReasons used by Order
    type ProductsReasons struct {

        
            Filters []ProductsReasonsFilters  `json:"filters"`
            Data ProductsReasonsData  `json:"data"`
         
    }
    
    // EntityReasonData used by Order
    type EntityReasonData struct {

        
            ReasonText string  `json:"reason_text"`
            ReasonID float64  `json:"reason_id"`
         
    }
    
    // EntitiesReasons used by Order
    type EntitiesReasons struct {

        
            Filters []map[string]interface{}  `json:"filters"`
            Data EntityReasonData  `json:"data"`
         
    }
    
    // ReasonsData used by Order
    type ReasonsData struct {

        
            Products []ProductsReasons  `json:"products"`
            Entities []EntitiesReasons  `json:"entities"`
         
    }
    
    // ShipmentsRequest used by Order
    type ShipmentsRequest struct {

        
            Identifier string  `json:"identifier"`
            DataUpdates DataUpdates  `json:"data_updates"`
            Products []Products  `json:"products"`
            Reasons ReasonsData  `json:"reasons"`
         
    }
    
    // StatuesRequest used by Order
    type StatuesRequest struct {

        
            Status string  `json:"status"`
            Shipments []ShipmentsRequest  `json:"shipments"`
            ExcludeBagsNextState string  `json:"exclude_bags_next_state"`
         
    }
    
    // UpdateShipmentStatusRequest used by Order
    type UpdateShipmentStatusRequest struct {

        
            LockAfterTransition bool  `json:"lock_after_transition"`
            Statuses []StatuesRequest  `json:"statuses"`
            Task bool  `json:"task"`
            ForceTransition bool  `json:"force_transition"`
            UnlockBeforeTransition bool  `json:"unlock_before_transition"`
         
    }
    
    // ShipmentsResponse used by Order
    type ShipmentsResponse struct {

        
            Code string  `json:"code"`
            StackTrace string  `json:"stack_trace"`
            Status float64  `json:"status"`
            Message string  `json:"message"`
            Identifier string  `json:"identifier"`
            Meta map[string]interface{}  `json:"meta"`
            FinalState map[string]interface{}  `json:"final_state"`
            Exception string  `json:"exception"`
         
    }
    
    // StatuesResponse used by Order
    type StatuesResponse struct {

        
            Shipments []ShipmentsResponse  `json:"shipments"`
         
    }
    
    // UpdateShipmentStatusResponseBody used by Order
    type UpdateShipmentStatusResponseBody struct {

        
            Statuses []StatuesResponse  `json:"statuses"`
         
    }
    
    // AffiliateAppConfigMeta used by Order
    type AffiliateAppConfigMeta struct {

        
            Value string  `json:"value"`
            Name string  `json:"name"`
         
    }
    
    // AffiliateAppConfig used by Order
    type AffiliateAppConfig struct {

        
            UpdatedAt string  `json:"updated_at"`
            ID string  `json:"id"`
            Name string  `json:"name"`
            Token string  `json:"token"`
            Owner string  `json:"owner"`
            Secret string  `json:"secret"`
            Description string  `json:"description"`
            Meta []AffiliateAppConfigMeta  `json:"meta"`
            CreatedAt string  `json:"created_at"`
         
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
    
    // AffiliateInventoryStoreConfig used by Order
    type AffiliateInventoryStoreConfig struct {

        
            Store map[string]interface{}  `json:"store"`
         
    }
    
    // AffiliateInventoryLogisticsConfig used by Order
    type AffiliateInventoryLogisticsConfig struct {

        
            DpAssignment bool  `json:"dp_assignment"`
         
    }
    
    // AffiliateInventoryOrderConfig used by Order
    type AffiliateInventoryOrderConfig struct {

        
            ForceReassignment bool  `json:"force_reassignment"`
         
    }
    
    // AffiliateInventoryConfig used by Order
    type AffiliateInventoryConfig struct {

        
            Payment AffiliateInventoryPaymentConfig  `json:"payment"`
            ArticleAssignment AffiliateInventoryArticleAssignmentConfig  `json:"article_assignment"`
            Inventory AffiliateInventoryStoreConfig  `json:"inventory"`
            Logistics AffiliateInventoryLogisticsConfig  `json:"logistics"`
            Order AffiliateInventoryOrderConfig  `json:"order"`
         
    }
    
    // AffiliateConfig used by Order
    type AffiliateConfig struct {

        
            App AffiliateAppConfig  `json:"app"`
            Inventory AffiliateInventoryConfig  `json:"inventory"`
         
    }
    
    // Affiliate used by Order
    type Affiliate struct {

        
            Token string  `json:"token"`
            ID string  `json:"id"`
            Config AffiliateConfig  `json:"config"`
         
    }
    
    // AffiliateStoreIdMapping used by Order
    type AffiliateStoreIdMapping struct {

        
            MarketplaceStoreID string  `json:"marketplace_store_id"`
            StoreID float64  `json:"store_id"`
         
    }
    
    // OrderConfig used by Order
    type OrderConfig struct {

        
            BagEndState string  `json:"bag_end_state"`
            Affiliate Affiliate  `json:"affiliate"`
            StoreLookup string  `json:"store_lookup"`
            CreateUser bool  `json:"create_user"`
            ArticleLookup string  `json:"article_lookup"`
            AffiliateStoreIDMapping []AffiliateStoreIdMapping  `json:"affiliate_store_id_mapping"`
         
    }
    
    // OrderPriority used by Order
    type OrderPriority struct {

        
            AffiliatePriorityCode string  `json:"affiliate_priority_code"`
            FulfilmentPriority float64  `json:"fulfilment_priority"`
            FulfilmentPriorityText string  `json:"fulfilment_priority_text"`
         
    }
    
    // OrderUser used by Order
    type OrderUser struct {

        
            Mobile float64  `json:"mobile"`
            Address1 string  `json:"address1"`
            Pincode string  `json:"pincode"`
            Email string  `json:"email"`
            Address2 string  `json:"address2"`
            Country string  `json:"country"`
            Phone float64  `json:"phone"`
            State string  `json:"state"`
            LastName string  `json:"last_name"`
            City string  `json:"city"`
            FirstName string  `json:"first_name"`
         
    }
    
    // ArticleDetails1 used by Order
    type ArticleDetails1 struct {

        
            BrandID float64  `json:"brand_id"`
            ID string  `json:"_id"`
            Weight map[string]interface{}  `json:"weight"`
            Attributes map[string]interface{}  `json:"attributes"`
            Dimension map[string]interface{}  `json:"dimension"`
            Quantity float64  `json:"quantity"`
            Category map[string]interface{}  `json:"category"`
         
    }
    
    // ShipmentDetails used by Order
    type ShipmentDetails struct {

        
            BoxType string  `json:"box_type"`
            Articles []ArticleDetails1  `json:"articles"`
            FulfillmentID float64  `json:"fulfillment_id"`
            Meta map[string]interface{}  `json:"meta"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            Shipments float64  `json:"shipments"`
            DpID float64  `json:"dp_id"`
         
    }
    
    // LocationDetails used by Order
    type LocationDetails struct {

        
            Articles []ArticleDetails1  `json:"articles"`
            FulfillmentType string  `json:"fulfillment_type"`
            FulfillmentID float64  `json:"fulfillment_id"`
         
    }
    
    // ShipmentConfig used by Order
    type ShipmentConfig struct {

        
            PaymentMode string  `json:"payment_mode"`
            Shipment []ShipmentDetails  `json:"shipment"`
            ToPincode string  `json:"to_pincode"`
            Action string  `json:"action"`
            LocationDetails LocationDetails  `json:"location_details"`
            Identifier string  `json:"identifier"`
            Source string  `json:"source"`
            Journey string  `json:"journey"`
         
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
            ItemID float64  `json:"item_id"`
            UnitPrice float64  `json:"unit_price"`
            StoreID float64  `json:"store_id"`
            AmountPaid float64  `json:"amount_paid"`
            ModifiedOn string  `json:"modified_on"`
            HsnCodeID string  `json:"hsn_code_id"`
            AvlQty float64  `json:"avl_qty"`
            AffiliateMeta map[string]interface{}  `json:"affiliate_meta"`
            PriceMarked float64  `json:"price_marked"`
            FyndStoreID string  `json:"fynd_store_id"`
            PriceEffective float64  `json:"price_effective"`
            ID string  `json:"_id"`
            Sku string  `json:"sku"`
            PdfLinks MarketPlacePdf  `json:"pdf_links"`
            DeliveryCharge float64  `json:"delivery_charge"`
            Quantity float64  `json:"quantity"`
            SellerIdentifier string  `json:"seller_identifier"`
            ItemSize string  `json:"item_size"`
            AffiliateStoreID string  `json:"affiliate_store_id"`
            CompanyID float64  `json:"company_id"`
            TransferPrice float64  `json:"transfer_price"`
            Identifier map[string]interface{}  `json:"identifier"`
         
    }
    
    // UserData used by Order
    type UserData struct {

        
            BillingUser OrderUser  `json:"billing_user"`
            ShippingUser OrderUser  `json:"shipping_user"`
         
    }
    
    // OrderInfo used by Order
    type OrderInfo struct {

        
            Payment map[string]interface{}  `json:"payment"`
            Discount float64  `json:"discount"`
            OrderPriority OrderPriority  `json:"order_priority"`
            BillingAddress OrderUser  `json:"billing_address"`
            DeliveryCharges float64  `json:"delivery_charges"`
            PaymentMode string  `json:"payment_mode"`
            OrderValue float64  `json:"order_value"`
            Coupon string  `json:"coupon"`
            Shipment ShipmentData  `json:"shipment"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            Bags []AffiliateBag  `json:"bags"`
            Items map[string]interface{}  `json:"items"`
            CodCharges float64  `json:"cod_charges"`
            ShippingAddress OrderUser  `json:"shipping_address"`
            User UserData  `json:"user"`
         
    }
    
    // CreateOrderPayload used by Order
    type CreateOrderPayload struct {

        
            OrderConfig OrderConfig  `json:"order_config"`
            AffiliateID string  `json:"affiliate_id"`
            OrderInfo OrderInfo  `json:"order_info"`
         
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

        
            Slug string  `json:"slug"`
            Description string  `json:"description"`
            ID float64  `json:"id"`
            DisplayText string  `json:"display_text"`
         
    }
    
    // GetActionsResponse used by Order
    type GetActionsResponse struct {

        
            Permissions ActionInfo  `json:"permissions"`
         
    }
    
    // PostHistoryFilters used by Order
    type PostHistoryFilters struct {

        
            Identifier string  `json:"identifier"`
            ShipmentID string  `json:"shipment_id"`
            LineNumber string  `json:"line_number"`
         
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
    
    // HistoryDict used by Order
    type HistoryDict struct {

        
            L3Detail string  `json:"l3_detail"`
            BagID float64  `json:"bag_id"`
            Message string  `json:"message"`
            TicketID string  `json:"ticket_id"`
            User string  `json:"user"`
            L2Detail string  `json:"l2_detail"`
            TicketURL string  `json:"ticket_url"`
            Createdat string  `json:"createdat"`
            L1Detail string  `json:"l1_detail"`
            Type string  `json:"type"`
         
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
    
    // SmsDataPayload used by Order
    type SmsDataPayload struct {

        
            OrderID string  `json:"order_id"`
            PhoneNumber float64  `json:"phone_number"`
            ShipmentID float64  `json:"shipment_id"`
            PaymentMode string  `json:"payment_mode"`
            CustomerName string  `json:"customer_name"`
            Message string  `json:"message"`
            BrandName string  `json:"brand_name"`
            AmountPaid float64  `json:"amount_paid"`
            CountryCode string  `json:"country_code"`
         
    }
    
    // SendSmsPayload used by Order
    type SendSmsPayload struct {

        
            Slug string  `json:"slug"`
            BagID float64  `json:"bag_id"`
            Data SmsDataPayload  `json:"data"`
         
    }
    
    // Meta used by Order
    type Meta struct {

        
            StateManagerUsed string  `json:"state_manager_used"`
            KafkaEmissionStatus float64  `json:"kafka_emission_status"`
         
    }
    
    // ShipmentDetail used by Order
    type ShipmentDetail struct {

        
            Status string  `json:"status"`
            ID float64  `json:"id"`
            ShipmentID string  `json:"shipment_id"`
            Remarks string  `json:"remarks"`
            BagList []float64  `json:"bag_list"`
            Meta Meta  `json:"meta"`
         
    }
    
    // OrderDetails used by Order
    type OrderDetails struct {

        
            FyndOrderID string  `json:"fynd_order_id"`
            CreatedAt string  `json:"created_at"`
         
    }
    
    // OrderStatusData used by Order
    type OrderStatusData struct {

        
            ShipmentDetails []ShipmentDetail  `json:"shipment_details"`
            OrderDetails OrderDetails  `json:"order_details"`
            Errors []string  `json:"errors"`
         
    }
    
    // OrderStatusResult used by Order
    type OrderStatusResult struct {

        
            Success string  `json:"success"`
            Result []OrderStatusData  `json:"result"`
         
    }
    
    // ManualAssignDPToShipment used by Order
    type ManualAssignDPToShipment struct {

        
            ShipmentIds []string  `json:"shipment_ids"`
            QcRequired string  `json:"qc_required"`
            OrderType string  `json:"order_type"`
            DpID float64  `json:"dp_id"`
         
    }
    
    // ManualAssignDPToShipmentResponse used by Order
    type ManualAssignDPToShipmentResponse struct {

        
            Success string  `json:"success"`
            Errors []string  `json:"errors"`
         
    }
    
    // PaymentMethod used by Order
    type PaymentMethod struct {

        
            RefundBy string  `json:"refund_by"`
            CollectBy string  `json:"collect_by"`
            Name string  `json:"name"`
            Amount float64  `json:"amount"`
            TransactionData map[string]interface{}  `json:"transaction_data"`
            Meta map[string]interface{}  `json:"meta"`
            Mode string  `json:"mode"`
         
    }
    
    // PaymentInfo used by Order
    type PaymentInfo struct {

        
            PaymentMethods []PaymentMethod  `json:"payment_methods"`
            PrimaryMode string  `json:"primary_mode"`
         
    }
    
    // Tax used by Order
    type Tax struct {

        
            Name string  `json:"name"`
            Breakup []map[string]interface{}  `json:"breakup"`
            Rate float64  `json:"rate"`
            Amount map[string]interface{}  `json:"amount"`
         
    }
    
    // Charge used by Order
    type Charge struct {

        
            Code string  `json:"code"`
            Name string  `json:"name"`
            Amount map[string]interface{}  `json:"amount"`
            Tax Tax  `json:"tax"`
            Type string  `json:"type"`
         
    }
    
    // LineItem used by Order
    type LineItem struct {

        
            ExternalLineID string  `json:"external_line_id"`
            Charges []Charge  `json:"charges"`
            Meta map[string]interface{}  `json:"meta"`
            Quantity float64  `json:"quantity"`
            CustomMessasge string  `json:"custom_messasge"`
            SellerIdentifier string  `json:"seller_identifier"`
         
    }
    
    // ProcessingDates used by Order
    type ProcessingDates struct {

        
            DpPickupSlot map[string]interface{}  `json:"dp_pickup_slot"`
            CustomerPickupSlot map[string]interface{}  `json:"customer_pickup_slot"`
            DispatchAfterDate string  `json:"dispatch_after_date"`
            PackByDate string  `json:"pack_by_date"`
            DispatchByDate string  `json:"dispatch_by_date"`
            ConfirmByDate string  `json:"confirm_by_date"`
         
    }
    
    // Shipment used by Order
    type Shipment struct {

        
            LocationID float64  `json:"location_id"`
            LineItems []LineItem  `json:"line_items"`
            Priority float64  `json:"priority"`
            ProcessingDates ProcessingDates  `json:"processing_dates"`
            Meta map[string]interface{}  `json:"meta"`
            ExternalShipmentID string  `json:"external_shipment_id"`
         
    }
    
    // TaxInfo used by Order
    type TaxInfo struct {

        
            Gstin string  `json:"gstin"`
            B2bGstinNumber string  `json:"b2b_gstin_number"`
         
    }
    
    // BillingInfo used by Order
    type BillingInfo struct {

        
            Country string  `json:"country"`
            PrimaryMobileNumber string  `json:"primary_mobile_number"`
            CustomerCode string  `json:"customer_code"`
            LastName string  `json:"last_name"`
            Title string  `json:"title"`
            Address1 string  `json:"address1"`
            State string  `json:"state"`
            HouseNo string  `json:"house_no"`
            FloorNo string  `json:"floor_no"`
            ExternalCustomerCode string  `json:"external_customer_code"`
            AlternateEmail string  `json:"alternate_email"`
            Gender string  `json:"gender"`
            CountryCode string  `json:"country_code"`
            FirstName string  `json:"first_name"`
            Pincode string  `json:"pincode"`
            PrimaryEmail string  `json:"primary_email"`
            Address2 string  `json:"address2"`
            MiddleName string  `json:"middle_name"`
            StateCode string  `json:"state_code"`
            AlternateMobileNumber string  `json:"alternate_mobile_number"`
            City string  `json:"city"`
         
    }
    
    // ShippingInfo used by Order
    type ShippingInfo struct {

        
            Country string  `json:"country"`
            Landmark string  `json:"landmark"`
            ShippingType string  `json:"shipping_type"`
            PrimaryMobileNumber string  `json:"primary_mobile_number"`
            CustomerCode string  `json:"customer_code"`
            LastName string  `json:"last_name"`
            Title string  `json:"title"`
            Address1 string  `json:"address1"`
            State string  `json:"state"`
            HouseNo string  `json:"house_no"`
            AddressType string  `json:"address_type"`
            FloorNo string  `json:"floor_no"`
            ExternalCustomerCode string  `json:"external_customer_code"`
            AlternateEmail string  `json:"alternate_email"`
            Gender string  `json:"gender"`
            CountryCode string  `json:"country_code"`
            FirstName string  `json:"first_name"`
            GeoLocation map[string]interface{}  `json:"geo_location"`
            Pincode string  `json:"pincode"`
            PrimaryEmail string  `json:"primary_email"`
            Address2 string  `json:"address2"`
            MiddleName string  `json:"middle_name"`
            Slot []map[string]interface{}  `json:"slot"`
            StateCode string  `json:"state_code"`
            AlternateMobileNumber string  `json:"alternate_mobile_number"`
            City string  `json:"city"`
         
    }
    
    // CreateOrderAPI used by Order
    type CreateOrderAPI struct {

        
            PaymentInfo PaymentInfo  `json:"payment_info"`
            Shipments []Shipment  `json:"shipments"`
            TaxInfo TaxInfo  `json:"tax_info"`
            BillingInfo BillingInfo  `json:"billing_info"`
            ExternalCreationDate string  `json:"external_creation_date"`
            ExternalOrderID string  `json:"external_order_id"`
            Charges []Charge  `json:"charges"`
            Meta map[string]interface{}  `json:"meta"`
            CurrencyInfo map[string]interface{}  `json:"currency_info"`
            ShippingInfo ShippingInfo  `json:"shipping_info"`
         
    }
    
    // CreateOrderErrorReponse used by Order
    type CreateOrderErrorReponse struct {

        
            Code string  `json:"code"`
            StackTrace string  `json:"stack_trace"`
            Status float64  `json:"status"`
            Message string  `json:"message"`
            RequestID string  `json:"request_id"`
            Info interface{}  `json:"info"`
            Meta string  `json:"meta"`
            Exception string  `json:"exception"`
         
    }
    
    // PaymentMethods used by Order
    type PaymentMethods struct {

        
            RefundBy string  `json:"refund_by"`
            CollectBy string  `json:"collect_by"`
            Mode string  `json:"mode"`
         
    }
    
    // CreateChannelPaymentInfo used by Order
    type CreateChannelPaymentInfo struct {

        
            PaymentMethods []PaymentMethods  `json:"payment_methods"`
            Source string  `json:"source"`
            ModeOfPayment string  `json:"mode_of_payment"`
         
    }
    
    // DpConfiguration used by Order
    type DpConfiguration struct {

        
            ShippingBy string  `json:"shipping_by"`
         
    }
    
    // CreateChannelConfig used by Order
    type CreateChannelConfig struct {

        
            PaymentInfo CreateChannelPaymentInfo  `json:"payment_info"`
            LogoURL map[string]interface{}  `json:"logo_url"`
            LocationReassignment bool  `json:"location_reassignment"`
            ShipmentAssignment string  `json:"shipment_assignment"`
            DpConfiguration DpConfiguration  `json:"dp_configuration"`
            LockStates []string  `json:"lock_states"`
         
    }
    
    // CreateChannelConfigData used by Order
    type CreateChannelConfigData struct {

        
            ConfigData CreateChannelConfig  `json:"config_data"`
         
    }
    
    // CreateChannelConfigResponse used by Order
    type CreateChannelConfigResponse struct {

        
            IsInserted bool  `json:"is_inserted"`
            Acknowledged bool  `json:"acknowledged"`
            IsUpserted bool  `json:"is_upserted"`
         
    }
    
    // CreateChannelConifgErrorResponse used by Order
    type CreateChannelConifgErrorResponse struct {

        
            Error string  `json:"error"`
         
    }
    
    // UploadConsent used by Order
    type UploadConsent struct {

        
            ConsentURL string  `json:"consent_url"`
            ManifestID string  `json:"manifest_id"`
         
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
            EndDate string  `json:"end_date"`
            OrderDetails []FyndOrderIdList  `json:"order_details"`
         
    }
    

    
    // DeleteResponse used by Catalog
    type DeleteResponse struct {

        
            Message string  `json:"message"`
         
    }
    
    // GetSearchWordsData used by Catalog
    type GetSearchWordsData struct {

        
            Words []string  `json:"words"`
            UID string  `json:"uid"`
            AppID string  `json:"app_id"`
            Result map[string]interface{}  `json:"result"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // GetSearchWordsDetailResponse used by Catalog
    type GetSearchWordsDetailResponse struct {

        
            Page Page  `json:"page"`
            Items GetSearchWordsData  `json:"items"`
         
    }
    
    // SearchKeywordResult used by Catalog
    type SearchKeywordResult struct {

        
            Query map[string]interface{}  `json:"query"`
            SortOn string  `json:"sort_on"`
         
    }
    
    // CreateSearchKeyword used by Catalog
    type CreateSearchKeyword struct {

        
            Words []string  `json:"words"`
            AppID string  `json:"app_id"`
            Result SearchKeywordResult  `json:"result"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // GetSearchWordsResponse used by Catalog
    type GetSearchWordsResponse struct {

        
            Page Page  `json:"page"`
            Items []GetSearchWordsData  `json:"items"`
         
    }
    
    // GetAutocompleteWordsData used by Catalog
    type GetAutocompleteWordsData struct {

        
            Words []string  `json:"words"`
            Results []map[string]interface{}  `json:"results"`
            UID string  `json:"uid"`
            AppID string  `json:"app_id"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // GetAutocompleteWordsResponse used by Catalog
    type GetAutocompleteWordsResponse struct {

        
            Page Page  `json:"page"`
            Items []GetAutocompleteWordsData  `json:"items"`
         
    }
    
    // AutocompletePageAction used by Catalog
    type AutocompletePageAction struct {

        
            Type string  `json:"type"`
            URL string  `json:"url"`
            Query map[string]interface{}  `json:"query"`
            Params map[string]interface{}  `json:"params"`
         
    }
    
    // AutocompleteAction used by Catalog
    type AutocompleteAction struct {

        
            Page AutocompletePageAction  `json:"page"`
            Type string  `json:"type"`
         
    }
    
    // Media used by Catalog
    type Media struct {

        
            Type string  `json:"type"`
            URL string  `json:"url"`
            AspectRatio string  `json:"aspect_ratio"`
         
    }
    
    // AutocompleteResult used by Catalog
    type AutocompleteResult struct {

        
            Display string  `json:"display"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Action AutocompleteAction  `json:"action"`
            Logo Media  `json:"logo"`
         
    }
    
    // CreateAutocompleteKeyword used by Catalog
    type CreateAutocompleteKeyword struct {

        
            Words []string  `json:"words"`
            Results []AutocompleteResult  `json:"results"`
            AppID string  `json:"app_id"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // CreateAutocompleteWordsResponse used by Catalog
    type CreateAutocompleteWordsResponse struct {

        
            AppID string  `json:"app_id"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Results []map[string]interface{}  `json:"results"`
            Words []string  `json:"words"`
         
    }
    
    // ProductBundleItem used by Catalog
    type ProductBundleItem struct {

        
            MaxQuantity float64  `json:"max_quantity"`
            ProductUID float64  `json:"product_uid"`
            MinQuantity float64  `json:"min_quantity"`
            AutoSelect bool  `json:"auto_select"`
            AllowRemove bool  `json:"allow_remove"`
            AutoAddToCart bool  `json:"auto_add_to_cart"`
         
    }
    
    // ProductBundleRequest used by Catalog
    type ProductBundleRequest struct {

        
            PageVisibility []string  `json:"page_visibility"`
            Products []ProductBundleItem  `json:"products"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Meta map[string]interface{}  `json:"meta"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Choice string  `json:"choice"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
            CompanyID float64  `json:"company_id"`
            Logo string  `json:"logo"`
            IsActive bool  `json:"is_active"`
            CreatedBy map[string]interface{}  `json:"created_by"`
         
    }
    
    // GetProductBundleCreateResponse used by Catalog
    type GetProductBundleCreateResponse struct {

        
            PageVisibility []string  `json:"page_visibility"`
            Products []ProductBundleItem  `json:"products"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Meta map[string]interface{}  `json:"meta"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Choice string  `json:"choice"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
            CompanyID float64  `json:"company_id"`
            Logo string  `json:"logo"`
            IsActive bool  `json:"is_active"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ID string  `json:"id"`
         
    }
    
    // GetProductBundleListingResponse used by Catalog
    type GetProductBundleListingResponse struct {

        
            Page Page  `json:"page"`
            Items []GetProductBundleCreateResponse  `json:"items"`
         
    }
    
    // LimitedProductData used by Catalog
    type LimitedProductData struct {

        
            Quantity float64  `json:"quantity"`
            Images []string  `json:"images"`
            UID float64  `json:"uid"`
            Price map[string]interface{}  `json:"price"`
            Slug string  `json:"slug"`
            Name string  `json:"name"`
            ItemCode string  `json:"item_code"`
            CountryOfOrigin string  `json:"country_of_origin"`
            Sizes []string  `json:"sizes"`
            Identifier map[string]interface{}  `json:"identifier"`
            ShortDescription string  `json:"short_description"`
            Attributes map[string]interface{}  `json:"attributes"`
         
    }
    
    // Price used by Catalog
    type Price struct {

        
            MinMarked float64  `json:"min_marked"`
            MaxEffective float64  `json:"max_effective"`
            MaxMarked float64  `json:"max_marked"`
            MinEffective float64  `json:"min_effective"`
            Currency string  `json:"currency"`
         
    }
    
    // Size used by Catalog
    type Size struct {

        
            IsAvailable bool  `json:"is_available"`
            Value string  `json:"value"`
            Display string  `json:"display"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // GetProducts used by Catalog
    type GetProducts struct {

        
            MaxQuantity float64  `json:"max_quantity"`
            ProductDetails LimitedProductData  `json:"product_details"`
            ProductUID float64  `json:"product_uid"`
            Price Price  `json:"price"`
            MinQuantity float64  `json:"min_quantity"`
            Sizes []Size  `json:"sizes"`
            AllowRemove bool  `json:"allow_remove"`
            AutoSelect bool  `json:"auto_select"`
            AutoAddToCart bool  `json:"auto_add_to_cart"`
         
    }
    
    // GetProductBundleResponse used by Catalog
    type GetProductBundleResponse struct {

        
            PageVisibility []string  `json:"page_visibility"`
            Products []GetProducts  `json:"products"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Meta map[string]interface{}  `json:"meta"`
            Choice string  `json:"choice"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            CompanyID float64  `json:"company_id"`
            Logo string  `json:"logo"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // ProductBundleUpdateRequest used by Catalog
    type ProductBundleUpdateRequest struct {

        
            PageVisibility []string  `json:"page_visibility"`
            Products []ProductBundleItem  `json:"products"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Meta map[string]interface{}  `json:"meta"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Choice string  `json:"choice"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            ModifiedOn string  `json:"modified_on"`
            CompanyID float64  `json:"company_id"`
            Logo string  `json:"logo"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // Guide used by Catalog
    type Guide struct {

        
            Meta Meta  `json:"meta"`
         
    }
    
    // ValidateSizeGuide used by Catalog
    type ValidateSizeGuide struct {

        
            Tag string  `json:"tag"`
            Guide Guide  `json:"guide"`
            Active bool  `json:"active"`
            Title string  `json:"title"`
            Name string  `json:"name"`
            BrandID float64  `json:"brand_id"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
            CompanyID float64  `json:"company_id"`
            Image string  `json:"image"`
            Subtitle string  `json:"subtitle"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Description string  `json:"description"`
            ID string  `json:"id"`
         
    }
    
    // ListSizeGuide used by Catalog
    type ListSizeGuide struct {

        
            Page map[string]interface{}  `json:"page"`
            Items []map[string]interface{}  `json:"items"`
         
    }
    
    // SizeGuideResponse used by Catalog
    type SizeGuideResponse struct {

        
            Tag string  `json:"tag"`
            Guide map[string]interface{}  `json:"guide"`
            Active bool  `json:"active"`
            Title string  `json:"title"`
            Name string  `json:"name"`
            BrandID float64  `json:"brand_id"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
            CompanyID float64  `json:"company_id"`
            Subtitle string  `json:"subtitle"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ID string  `json:"id"`
         
    }
    
    // MOQData used by Catalog
    type MOQData struct {

        
            IncrementUnit float64  `json:"increment_unit"`
            Maximum float64  `json:"maximum"`
            Minimum float64  `json:"minimum"`
         
    }
    
    // SEOData used by Catalog
    type SEOData struct {

        
            Description interface{}  `json:"description"`
            Title interface{}  `json:"title"`
         
    }
    
    // OwnerAppItemResponse used by Catalog
    type OwnerAppItemResponse struct {

        
            IsGift bool  `json:"is_gift"`
            Moq MOQData  `json:"moq"`
            Seo SEOData  `json:"seo"`
            IsCod bool  `json:"is_cod"`
            AltText map[string]interface{}  `json:"alt_text"`
         
    }
    
    // ApplicationItemMOQ used by Catalog
    type ApplicationItemMOQ struct {

        
            IncrementUnit float64  `json:"increment_unit"`
            Maximum float64  `json:"maximum"`
            Minimum float64  `json:"minimum"`
         
    }
    
    // MetaFields used by Catalog
    type MetaFields struct {

        
            Key interface{}  `json:"key"`
            Value interface{}  `json:"value"`
         
    }
    
    // ApplicationItemSEO used by Catalog
    type ApplicationItemSEO struct {

        
            Description interface{}  `json:"description"`
            Title interface{}  `json:"title"`
         
    }
    
    // ApplicationItemMeta used by Catalog
    type ApplicationItemMeta struct {

        
            IsGift bool  `json:"is_gift"`
            Moq ApplicationItemMOQ  `json:"moq"`
            CustomMeta []MetaFields  `json:"_custom_meta"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Seo ApplicationItemSEO  `json:"seo"`
            IsCod bool  `json:"is_cod"`
            AltText map[string]interface{}  `json:"alt_text"`
         
    }
    
    // SuccessResponse1 used by Catalog
    type SuccessResponse1 struct {

        
            Success bool  `json:"success"`
            UID float64  `json:"uid"`
         
    }
    
    // GetConfigMetadataResponse used by Catalog
    type GetConfigMetadataResponse struct {

        
            Condition []map[string]interface{}  `json:"condition"`
            Data []map[string]interface{}  `json:"data"`
            Values []map[string]interface{}  `json:"values"`
         
    }
    
    // AttributeDetailsGroup used by Catalog
    type AttributeDetailsGroup struct {

        
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Priority float64  `json:"priority"`
            Key string  `json:"key"`
            Logo string  `json:"logo"`
            Unit string  `json:"unit"`
            DisplayType string  `json:"display_type"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // AppConfigurationDetail used by Catalog
    type AppConfigurationDetail struct {

        
            IsDefault bool  `json:"is_default"`
            AppID string  `json:"app_id"`
            Slug string  `json:"slug"`
            Name string  `json:"name"`
            Priority float64  `json:"priority"`
            TemplateSlugs []string  `json:"template_slugs"`
            Logo string  `json:"logo"`
            IsActive bool  `json:"is_active"`
            Attributes []AttributeDetailsGroup  `json:"attributes"`
         
    }
    
    // ConfigErrorResponse used by Catalog
    type ConfigErrorResponse struct {

        
            Message string  `json:"message"`
         
    }
    
    // PageResponseType used by Catalog
    type PageResponseType struct {

        
            Current float64  `json:"current"`
            HasNext bool  `json:"has_next"`
            TotalCount float64  `json:"total_count"`
            Next float64  `json:"next"`
         
    }
    
    // GetConfigResponse used by Catalog
    type GetConfigResponse struct {

        
            Page PageResponseType  `json:"page"`
            Data []map[string]interface{}  `json:"data"`
         
    }
    
    // ConfigSuccessResponse used by Catalog
    type ConfigSuccessResponse struct {

        
            Message string  `json:"message"`
         
    }
    
    // AppConfigurationsSort used by Catalog
    type AppConfigurationsSort struct {

        
            DefaultKey string  `json:"default_key"`
            IsDefault bool  `json:"is_default"`
            AppID string  `json:"app_id"`
            Name string  `json:"name"`
            Priority float64  `json:"priority"`
            Key string  `json:"key"`
            Logo string  `json:"logo"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // AllowSingleRequest used by Catalog
    type AllowSingleRequest struct {

        
            AllowSingle bool  `json:"allow_single"`
         
    }
    
    // DefaultKeyRequest used by Catalog
    type DefaultKeyRequest struct {

        
            DefaultKey string  `json:"default_key"`
         
    }
    
    // GetCatalogConfigurationDetailsProduct used by Catalog
    type GetCatalogConfigurationDetailsProduct struct {

        
            Variant map[string]interface{}  `json:"variant"`
            Compare map[string]interface{}  `json:"compare"`
            Detail map[string]interface{}  `json:"detail"`
            Similar map[string]interface{}  `json:"similar"`
         
    }
    
    // MetaDataListingFilterMetaResponse used by Catalog
    type MetaDataListingFilterMetaResponse struct {

        
            Key string  `json:"key"`
            Display string  `json:"display"`
            FilterTypes []string  `json:"filter_types"`
            Units []map[string]interface{}  `json:"units"`
         
    }
    
    // MetaDataListingFilterResponse used by Catalog
    type MetaDataListingFilterResponse struct {

        
            Data []MetaDataListingFilterMetaResponse  `json:"data"`
         
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
    
    // MetaDataListingResponse used by Catalog
    type MetaDataListingResponse struct {

        
            Filter MetaDataListingFilterResponse  `json:"filter"`
            Sort MetaDataListingSortResponse  `json:"sort"`
         
    }
    
    // GetCatalogConfigurationMetaData used by Catalog
    type GetCatalogConfigurationMetaData struct {

        
            Product GetCatalogConfigurationDetailsProduct  `json:"product"`
            Listing MetaDataListingResponse  `json:"listing"`
         
    }
    
    // ProductSize used by Catalog
    type ProductSize struct {

        
            Min float64  `json:"min"`
            Max float64  `json:"max"`
         
    }
    
    // ConfigurationProductVariantConfig used by Catalog
    type ConfigurationProductVariantConfig struct {

        
            Name string  `json:"name"`
            Priority float64  `json:"priority"`
            Key string  `json:"key"`
            Size ProductSize  `json:"size"`
            Logo string  `json:"logo"`
            DisplayType string  `json:"display_type"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // ConfigurationProductVariant used by Catalog
    type ConfigurationProductVariant struct {

        
            Config []ConfigurationProductVariantConfig  `json:"config"`
         
    }
    
    // ConfigurationProductConfig used by Catalog
    type ConfigurationProductConfig struct {

        
            Title string  `json:"title"`
            Priority float64  `json:"priority"`
            Key string  `json:"key"`
            Size ProductSize  `json:"size"`
            Logo string  `json:"logo"`
            Subtitle string  `json:"subtitle"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // ConfigurationProductSimilar used by Catalog
    type ConfigurationProductSimilar struct {

        
            Config []ConfigurationProductConfig  `json:"config"`
         
    }
    
    // ConfigurationProduct used by Catalog
    type ConfigurationProduct struct {

        
            Variant ConfigurationProductVariant  `json:"variant"`
            Similar ConfigurationProductSimilar  `json:"similar"`
         
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
            BucketPoints []ConfigurationBucketPoints  `json:"bucket_points"`
            Value string  `json:"value"`
            Map map[string]interface{}  `json:"map"`
            Condition string  `json:"condition"`
            MapValues []map[string]interface{}  `json:"map_values"`
         
    }
    
    // ConfigurationListingFilterConfig used by Catalog
    type ConfigurationListingFilterConfig struct {

        
            Name string  `json:"name"`
            Type string  `json:"type"`
            Priority float64  `json:"priority"`
            Key string  `json:"key"`
            Logo string  `json:"logo"`
            IsActive bool  `json:"is_active"`
            DisplayName string  `json:"display_name"`
            ValueConfig ConfigurationListingFilterValue  `json:"value_config"`
         
    }
    
    // ConfigurationListingFilter used by Catalog
    type ConfigurationListingFilter struct {

        
            AttributeConfig []ConfigurationListingFilterConfig  `json:"attribute_config"`
            AllowSingle bool  `json:"allow_single"`
         
    }
    
    // ConfigurationListingSortConfig used by Catalog
    type ConfigurationListingSortConfig struct {

        
            Name string  `json:"name"`
            Priority float64  `json:"priority"`
            Key string  `json:"key"`
            Logo string  `json:"logo"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // ConfigurationListingSort used by Catalog
    type ConfigurationListingSort struct {

        
            DefaultKey string  `json:"default_key"`
            Config []ConfigurationListingSortConfig  `json:"config"`
         
    }
    
    // ConfigurationListing used by Catalog
    type ConfigurationListing struct {

        
            Filter ConfigurationListingFilter  `json:"filter"`
            Sort ConfigurationListingSort  `json:"sort"`
         
    }
    
    // AppConfiguration used by Catalog
    type AppConfiguration struct {

        
            ConfigID string  `json:"config_id"`
            AppID string  `json:"app_id"`
            Type string  `json:"type"`
            Product ConfigurationProduct  `json:"product"`
            Listing ConfigurationListing  `json:"listing"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
            ConfigType string  `json:"config_type"`
            CreatedBy map[string]interface{}  `json:"created_by"`
         
    }
    
    // AppCatalogConfiguration used by Catalog
    type AppCatalogConfiguration struct {

        
            ConfigID string  `json:"config_id"`
            AppID string  `json:"app_id"`
            Type string  `json:"type"`
            Product ConfigurationProduct  `json:"product"`
            Listing ConfigurationListing  `json:"listing"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
            ConfigType string  `json:"config_type"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ID string  `json:"id"`
         
    }
    
    // GetAppCatalogConfiguration used by Catalog
    type GetAppCatalogConfiguration struct {

        
            Data AppCatalogConfiguration  `json:"data"`
            IsDefault bool  `json:"is_default"`
         
    }
    
    // GetCatalogConfigurationDetailsSchemaListing used by Catalog
    type GetCatalogConfigurationDetailsSchemaListing struct {

        
            Filter map[string]interface{}  `json:"filter"`
            Sort map[string]interface{}  `json:"sort"`
         
    }
    
    // EntityConfiguration used by Catalog
    type EntityConfiguration struct {

        
            ConfigID string  `json:"config_id"`
            AppID string  `json:"app_id"`
            Product GetCatalogConfigurationDetailsProduct  `json:"product"`
            Listing GetCatalogConfigurationDetailsSchemaListing  `json:"listing"`
            ConfigType string  `json:"config_type"`
            ID string  `json:"id"`
         
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
            Kind string  `json:"kind"`
            Logo string  `json:"logo"`
         
    }
    
    // ProductFiltersValue used by Catalog
    type ProductFiltersValue struct {

        
            SelectedMax float64  `json:"selected_max"`
            Max float64  `json:"max"`
            QueryFormat string  `json:"query_format"`
            CurrencyCode string  `json:"currency_code"`
            DisplayFormat string  `json:"display_format"`
            Value interface{}  `json:"value"`
            Display string  `json:"display"`
            Min float64  `json:"min"`
            CurrencySymbol string  `json:"currency_symbol"`
            Count float64  `json:"count"`
            SelectedMin float64  `json:"selected_min"`
            IsSelected bool  `json:"is_selected"`
         
    }
    
    // ProductFilters used by Catalog
    type ProductFilters struct {

        
            Key ProductFiltersKey  `json:"key"`
            Values []ProductFiltersValue  `json:"values"`
         
    }
    
    // GetCollectionQueryOptionResponse used by Catalog
    type GetCollectionQueryOptionResponse struct {

        
            Operators map[string]string  `json:"operators"`
            SortOn []ProductSortOn  `json:"sort_on"`
            Filters []ProductFilters  `json:"filters"`
         
    }
    
    // CollectionImage used by Catalog
    type CollectionImage struct {

        
            URL string  `json:"url"`
            AspectRatio string  `json:"aspect_ratio"`
         
    }
    
    // CollectionBanner used by Catalog
    type CollectionBanner struct {

        
            Landscape CollectionImage  `json:"landscape"`
            Portrait CollectionImage  `json:"portrait"`
         
    }
    
    // UserInfo used by Catalog
    type UserInfo struct {

        
            UID string  `json:"uid"`
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
            Email string  `json:"email"`
         
    }
    
    // SeoDetail used by Catalog
    type SeoDetail struct {

        
            Description string  `json:"description"`
            Title string  `json:"title"`
         
    }
    
    // CollectionBadge used by Catalog
    type CollectionBadge struct {

        
            Color string  `json:"color"`
            Text string  `json:"text"`
         
    }
    
    // CollectionQuery used by Catalog
    type CollectionQuery struct {

        
            Attribute string  `json:"attribute"`
            Value []interface{}  `json:"value"`
            Op string  `json:"op"`
         
    }
    
    // CollectionSchedule used by Catalog
    type CollectionSchedule struct {

        
            Cron string  `json:"cron"`
            NextSchedule []NextSchedule  `json:"next_schedule"`
            End string  `json:"end"`
            Duration float64  `json:"duration"`
            Start string  `json:"start"`
         
    }
    
    // CreateCollection used by Catalog
    type CreateCollection struct {

        
            IsVisible bool  `json:"is_visible"`
            Name string  `json:"name"`
            SortOn string  `json:"sort_on"`
            Logo CollectionImage  `json:"logo"`
            IsActive bool  `json:"is_active"`
            Banners CollectionBanner  `json:"banners"`
            CreatedBy UserInfo  `json:"created_by"`
            AllowSort bool  `json:"allow_sort"`
            Description string  `json:"description"`
            Type string  `json:"type"`
            Slug string  `json:"slug"`
            Priority float64  `json:"priority"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Seo SeoDetail  `json:"seo"`
            AllowFacets bool  `json:"allow_facets"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            Badge CollectionBadge  `json:"badge"`
            Published bool  `json:"published"`
            Meta map[string]interface{}  `json:"meta"`
            Query []CollectionQuery  `json:"query"`
            ModifiedBy UserInfo  `json:"modified_by"`
            AppID string  `json:"app_id"`
            Schedule CollectionSchedule  `json:"_schedule"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Tags []string  `json:"tags"`
         
    }
    
    // BannerImage used by Catalog
    type BannerImage struct {

        
            Type string  `json:"type"`
            URL string  `json:"url"`
            AspectRatio string  `json:"aspect_ratio"`
         
    }
    
    // ImageUrls used by Catalog
    type ImageUrls struct {

        
            Landscape BannerImage  `json:"landscape"`
            Portrait BannerImage  `json:"portrait"`
         
    }
    
    // CollectionCreateResponse used by Catalog
    type CollectionCreateResponse struct {

        
            Name string  `json:"name"`
            SortOn string  `json:"sort_on"`
            Logo BannerImage  `json:"logo"`
            IsActive bool  `json:"is_active"`
            Banners ImageUrls  `json:"banners"`
            AllowSort bool  `json:"allow_sort"`
            Description string  `json:"description"`
            Slug string  `json:"slug"`
            Type string  `json:"type"`
            Priority float64  `json:"priority"`
            AllowFacets bool  `json:"allow_facets"`
            Badge map[string]interface{}  `json:"badge"`
            Meta map[string]interface{}  `json:"meta"`
            Query []CollectionQuery  `json:"query"`
            Cron map[string]interface{}  `json:"cron"`
            Tag []string  `json:"tag"`
            AppID string  `json:"app_id"`
            Schedule map[string]interface{}  `json:"_schedule"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
         
    }
    
    // Media1 used by Catalog
    type Media1 struct {

        
            Meta map[string]interface{}  `json:"meta"`
            Type string  `json:"type"`
            URL string  `json:"url"`
         
    }
    
    // GetCollectionDetailNest used by Catalog
    type GetCollectionDetailNest struct {

        
            UID string  `json:"uid"`
            Name string  `json:"name"`
            Logo Media1  `json:"logo"`
            IsActive bool  `json:"is_active"`
            Banners ImageUrls  `json:"banners"`
            AllowSort bool  `json:"allow_sort"`
            Description string  `json:"description"`
            Slug string  `json:"slug"`
            Type string  `json:"type"`
            Priority float64  `json:"priority"`
            AllowFacets bool  `json:"allow_facets"`
            Badge map[string]interface{}  `json:"badge"`
            Meta map[string]interface{}  `json:"meta"`
            Query []CollectionQuery  `json:"query"`
            Cron map[string]interface{}  `json:"cron"`
            Tag []string  `json:"tag"`
            Action Action  `json:"action"`
            AppID string  `json:"app_id"`
            Schedule map[string]interface{}  `json:"_schedule"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
         
    }
    
    // CollectionListingFilterType used by Catalog
    type CollectionListingFilterType struct {

        
            Name string  `json:"name"`
            Display string  `json:"display"`
            IsSelected bool  `json:"is_selected"`
         
    }
    
    // CollectionListingFilterTag used by Catalog
    type CollectionListingFilterTag struct {

        
            Name string  `json:"name"`
            Display string  `json:"display"`
            IsSelected bool  `json:"is_selected"`
         
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
    
    // CollectionDetailResponse used by Catalog
    type CollectionDetailResponse struct {

        
            UID string  `json:"uid"`
            Name string  `json:"name"`
            Logo Media1  `json:"logo"`
            IsActive bool  `json:"is_active"`
            Banners ImageUrls  `json:"banners"`
            AllowSort bool  `json:"allow_sort"`
            Description string  `json:"description"`
            Slug string  `json:"slug"`
            Type string  `json:"type"`
            Priority float64  `json:"priority"`
            AllowFacets bool  `json:"allow_facets"`
            Badge map[string]interface{}  `json:"badge"`
            Meta map[string]interface{}  `json:"meta"`
            Query []CollectionQuery  `json:"query"`
            Cron map[string]interface{}  `json:"cron"`
            Tag []string  `json:"tag"`
            AppID string  `json:"app_id"`
            Schedule map[string]interface{}  `json:"_schedule"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
         
    }
    
    // UpdateCollection used by Catalog
    type UpdateCollection struct {

        
            IsVisible bool  `json:"is_visible"`
            Name string  `json:"name"`
            SortOn string  `json:"sort_on"`
            Logo CollectionImage  `json:"logo"`
            IsActive bool  `json:"is_active"`
            Banners CollectionBanner  `json:"banners"`
            AllowSort bool  `json:"allow_sort"`
            Description string  `json:"description"`
            Slug string  `json:"slug"`
            Type string  `json:"type"`
            Priority float64  `json:"priority"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Seo SeoDetail  `json:"seo"`
            AllowFacets bool  `json:"allow_facets"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            Badge CollectionBadge  `json:"badge"`
            Published bool  `json:"published"`
            Meta map[string]interface{}  `json:"meta"`
            Query []CollectionQuery  `json:"query"`
            ModifiedBy UserInfo  `json:"modified_by"`
            Schedule CollectionSchedule  `json:"_schedule"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Tags []string  `json:"tags"`
         
    }
    
    // ItemQueryForUserCollection used by Catalog
    type ItemQueryForUserCollection struct {

        
            Action string  `json:"action"`
            ItemID float64  `json:"item_id"`
         
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
    
    // ProductBrand used by Catalog
    type ProductBrand struct {

        
            Name string  `json:"name"`
            Logo Media1  `json:"logo"`
            Action Action  `json:"action"`
            UID float64  `json:"uid"`
         
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
    
    // ProductListingDetail used by Catalog
    type ProductListingDetail struct {

        
            Similars []string  `json:"similars"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            Color string  `json:"color"`
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
            Description string  `json:"description"`
            Type string  `json:"type"`
            Slug string  `json:"slug"`
            Tryouts []string  `json:"tryouts"`
            Highlights []string  `json:"highlights"`
            ProductOnlineDate string  `json:"product_online_date"`
            ItemType string  `json:"item_type"`
            Attributes map[string]interface{}  `json:"attributes"`
            RatingCount float64  `json:"rating_count"`
            ImageNature string  `json:"image_nature"`
            Rating float64  `json:"rating"`
            ItemCode string  `json:"item_code"`
            Brand ProductBrand  `json:"brand"`
            ShortDescription string  `json:"short_description"`
            Discount string  `json:"discount"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            Price ProductListingPrice  `json:"price"`
            PromoMeta map[string]interface{}  `json:"promo_meta"`
            Sellable bool  `json:"sellable"`
            Medias []Media1  `json:"medias"`
            HasVariant bool  `json:"has_variant"`
         
    }
    
    // GetCollectionItemsResponse used by Catalog
    type GetCollectionItemsResponse struct {

        
            Page Page  `json:"page"`
            SortOn []ProductSortOn  `json:"sort_on"`
            Items []ProductListingDetail  `json:"items"`
            Filters []ProductFilters  `json:"filters"`
         
    }
    
    // CatalogInsightItem used by Catalog
    type CatalogInsightItem struct {

        
            OutOfStockCount float64  `json:"out_of_stock_count"`
            Count float64  `json:"count"`
            SellableCount float64  `json:"sellable_count"`
         
    }
    
    // CatalogInsightBrand used by Catalog
    type CatalogInsightBrand struct {

        
            AvailableSizes float64  `json:"available_sizes"`
            TotalArticles float64  `json:"total_articles"`
            Name string  `json:"name"`
            TotalSizes float64  `json:"total_sizes"`
            ArticleFreshness float64  `json:"article_freshness"`
            AvailableArticles float64  `json:"available_articles"`
         
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

        
            Enabled bool  `json:"enabled"`
            BrandIds []float64  `json:"brand_ids"`
            StoreIds []float64  `json:"store_ids"`
            Platform string  `json:"platform"`
            CompanyID float64  `json:"company_id"`
            OptLevel string  `json:"opt_level"`
         
    }
    
    // CompanyOptIn used by Catalog
    type CompanyOptIn struct {

        
            CreatedBy map[string]interface{}  `json:"created_by"`
            Enabled bool  `json:"enabled"`
            BrandIds []float64  `json:"brand_ids"`
            StoreIds []float64  `json:"store_ids"`
            Platform string  `json:"platform"`
            ModifiedOn float64  `json:"modified_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            CreatedOn float64  `json:"created_on"`
            CompanyID float64  `json:"company_id"`
            OptLevel string  `json:"opt_level"`
         
    }
    
    // GetOptInPlatform used by Catalog
    type GetOptInPlatform struct {

        
            Page Page  `json:"page"`
            Items []CompanyOptIn  `json:"items"`
         
    }
    
    // OptinCompanyDetail used by Catalog
    type OptinCompanyDetail struct {

        
            Name string  `json:"name"`
            CompanyType string  `json:"company_type"`
            BusinessType string  `json:"business_type"`
            UID float64  `json:"uid"`
         
    }
    
    // CompanyBrandDetail used by Catalog
    type CompanyBrandDetail struct {

        
            BrandID float64  `json:"brand_id"`
            TotalArticle float64  `json:"total_article"`
            BrandName string  `json:"brand_name"`
            CompanyID float64  `json:"company_id"`
         
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
            Name string  `json:"name"`
            Documents []map[string]interface{}  `json:"documents"`
            Address map[string]interface{}  `json:"address"`
            StoreType string  `json:"store_type"`
            StoreCode string  `json:"store_code"`
            ModifiedOn string  `json:"modified_on"`
            Manager map[string]interface{}  `json:"manager"`
            CreatedOn string  `json:"created_on"`
            CompanyID float64  `json:"company_id"`
            AdditionalContacts []map[string]interface{}  `json:"additional_contacts"`
            Timing map[string]interface{}  `json:"timing"`
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

        
            Enriched bool  `json:"enriched"`
            MandatoryDetails AttributeMasterMandatoryDetails  `json:"mandatory_details"`
         
    }
    
    // AttributeMasterDetails used by Catalog
    type AttributeMasterDetails struct {

        
            DisplayType string  `json:"display_type"`
         
    }
    
    // AttributeSchemaRange used by Catalog
    type AttributeSchemaRange struct {

        
            Min float64  `json:"min"`
            Max float64  `json:"max"`
         
    }
    
    // AttributeMaster used by Catalog
    type AttributeMaster struct {

        
            AllowedValues []string  `json:"allowed_values"`
            Multi bool  `json:"multi"`
            Type string  `json:"type"`
            Mandatory bool  `json:"mandatory"`
            Range AttributeSchemaRange  `json:"range"`
            Format string  `json:"format"`
         
    }
    
    // AttributeMasterFilter used by Catalog
    type AttributeMasterFilter struct {

        
            DependsOn []string  `json:"depends_on"`
            Priority float64  `json:"priority"`
            Indexing bool  `json:"indexing"`
         
    }
    
    // GenderDetail used by Catalog
    type GenderDetail struct {

        
            IsNested bool  `json:"is_nested"`
            Meta AttributeMasterMeta  `json:"meta"`
            Slug string  `json:"slug"`
            Name string  `json:"name"`
            Details AttributeMasterDetails  `json:"details"`
            EnabledForEndConsumer bool  `json:"enabled_for_end_consumer"`
            Schema AttributeMaster  `json:"schema"`
            Filters AttributeMasterFilter  `json:"filters"`
            Logo string  `json:"logo"`
            Departments []string  `json:"departments"`
            Description string  `json:"description"`
            ID string  `json:"id"`
         
    }
    
    // ProdcutTemplateCategoriesResponse used by Catalog
    type ProdcutTemplateCategoriesResponse struct {

        
            Page Page  `json:"page"`
            Items []map[string]interface{}  `json:"items"`
         
    }
    
    // PTErrorResponse used by Catalog
    type PTErrorResponse struct {

        
            Status float64  `json:"status"`
            Code string  `json:"code"`
            Meta map[string]interface{}  `json:"meta"`
            Message string  `json:"message"`
            Errors map[string]interface{}  `json:"errors"`
         
    }
    
    // DepartmentCreateUpdate used by Catalog
    type DepartmentCreateUpdate struct {

        
            Cls string  `json:"_cls"`
            Tags []string  `json:"tags"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            Platforms map[string]interface{}  `json:"platforms"`
            Slug string  `json:"slug"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Synonyms []string  `json:"synonyms"`
            Logo string  `json:"logo"`
            IsActive bool  `json:"is_active"`
            PriorityOrder float64  `json:"priority_order"`
         
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
    
    // UserSerializer used by Catalog
    type UserSerializer struct {

        
            UID string  `json:"uid"`
            Contact string  `json:"contact"`
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
            ID string  `json:"_id"`
         
    }
    
    // GetDepartment used by Catalog
    type GetDepartment struct {

        
            PageNo float64  `json:"page_no"`
            CreatedBy UserSerializer  `json:"created_by"`
            Search string  `json:"search"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            PageSize float64  `json:"page_size"`
            ModifiedOn string  `json:"modified_on"`
            Synonyms []string  `json:"synonyms"`
            CreatedOn string  `json:"created_on"`
            Logo string  `json:"logo"`
            ItemType string  `json:"item_type"`
            IsActive bool  `json:"is_active"`
            PriorityOrder float64  `json:"priority_order"`
         
    }
    
    // DepartmentsResponse used by Catalog
    type DepartmentsResponse struct {

        
            Page Page  `json:"page"`
            Items []GetDepartment  `json:"items"`
         
    }
    
    // DepartmentErrorResponse used by Catalog
    type DepartmentErrorResponse struct {

        
            Status float64  `json:"status"`
            Code string  `json:"code"`
            Meta map[string]interface{}  `json:"meta"`
            Message string  `json:"message"`
            Errors map[string]interface{}  `json:"errors"`
         
    }
    
    // UserDetail used by Catalog
    type UserDetail struct {

        
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
            SuperUser bool  `json:"super_user"`
            Contact string  `json:"contact"`
         
    }
    
    // DepartmentModel used by Catalog
    type DepartmentModel struct {

        
            Cls interface{}  `json:"_cls"`
            PriorityOrder float64  `json:"priority_order"`
            UID float64  `json:"uid"`
            Name interface{}  `json:"name"`
            Slug interface{}  `json:"slug"`
            ModifiedBy UserDetail  `json:"modified_by"`
            ModifiedOn string  `json:"modified_on"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            VerifiedBy UserDetail  `json:"verified_by"`
            CreatedOn string  `json:"created_on"`
            Synonyms []interface{}  `json:"synonyms"`
            Logo interface{}  `json:"logo"`
            ID interface{}  `json:"_id"`
            IsActive bool  `json:"is_active"`
            CreatedBy UserDetail  `json:"created_by"`
            VerifiedOn string  `json:"verified_on"`
         
    }
    
    // ProductTemplate used by Catalog
    type ProductTemplate struct {

        
            CreatedBy map[string]interface{}  `json:"created_by"`
            Tag string  `json:"tag"`
            AttributesSchema []map[string]interface{}  `json:"attributes_schema"`
            IsPhysical bool  `json:"is_physical"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Categories []string  `json:"categories"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
            IsArchived bool  `json:"is_archived"`
            Logo string  `json:"logo"`
            IsExpirable bool  `json:"is_expirable"`
            IsActive bool  `json:"is_active"`
            Departments []string  `json:"departments"`
            Description string  `json:"description"`
            Attributes []string  `json:"attributes"`
         
    }
    
    // TemplatesResponse used by Catalog
    type TemplatesResponse struct {

        
            Page Page  `json:"page"`
            Items ProductTemplate  `json:"items"`
         
    }
    
    // TemplateDetails used by Catalog
    type TemplateDetails struct {

        
            Tag string  `json:"tag"`
            AttributesSchema []map[string]interface{}  `json:"attributes_schema"`
            IsPhysical bool  `json:"is_physical"`
            ID string  `json:"id"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Categories []string  `json:"categories"`
            IsArchived bool  `json:"is_archived"`
            Logo string  `json:"logo"`
            IsExpirable bool  `json:"is_expirable"`
            IsActive bool  `json:"is_active"`
            Departments []string  `json:"departments"`
            Description string  `json:"description"`
            Attributes []string  `json:"attributes"`
         
    }
    
    // Properties used by Catalog
    type Properties struct {

        
            MultiSize map[string]interface{}  `json:"multi_size"`
            Name map[string]interface{}  `json:"name"`
            IsActive map[string]interface{}  `json:"is_active"`
            Description map[string]interface{}  `json:"description"`
            Tags map[string]interface{}  `json:"tags"`
            ProductPublish map[string]interface{}  `json:"product_publish"`
            Slug map[string]interface{}  `json:"slug"`
            Sizes map[string]interface{}  `json:"sizes"`
            Highlights map[string]interface{}  `json:"highlights"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            IsDependent map[string]interface{}  `json:"is_dependent"`
            SizeGuide map[string]interface{}  `json:"size_guide"`
            ItemType map[string]interface{}  `json:"item_type"`
            Variants map[string]interface{}  `json:"variants"`
            CountryOfOrigin map[string]interface{}  `json:"country_of_origin"`
            ItemCode map[string]interface{}  `json:"item_code"`
            TraderType map[string]interface{}  `json:"trader_type"`
            Media map[string]interface{}  `json:"media"`
            NoOfBoxes map[string]interface{}  `json:"no_of_boxes"`
            ShortDescription map[string]interface{}  `json:"short_description"`
            CategorySlug map[string]interface{}  `json:"category_slug"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            BrandUID map[string]interface{}  `json:"brand_uid"`
            HsnCode map[string]interface{}  `json:"hsn_code"`
            ProductGroupTag map[string]interface{}  `json:"product_group_tag"`
            ReturnConfig map[string]interface{}  `json:"return_config"`
            Command map[string]interface{}  `json:"command"`
            Trader map[string]interface{}  `json:"trader"`
            Currency map[string]interface{}  `json:"currency"`
         
    }
    
    // GlobalValidation used by Catalog
    type GlobalValidation struct {

        
            Title string  `json:"title"`
            Type string  `json:"type"`
            Definitions map[string]interface{}  `json:"definitions"`
            Properties Properties  `json:"properties"`
            Required []string  `json:"required"`
            Description string  `json:"description"`
         
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

        
            Data map[string]interface{}  `json:"data"`
            Message string  `json:"message"`
         
    }
    
    // HSNData used by Catalog
    type HSNData struct {

        
            HsnCode []string  `json:"hsn_code"`
            CountryOfOrigin []string  `json:"country_of_origin"`
         
    }
    
    // HSNCodesResponse used by Catalog
    type HSNCodesResponse struct {

        
            Data HSNData  `json:"data"`
            Message string  `json:"message"`
         
    }
    
    // ProductDownloadItemsData used by Catalog
    type ProductDownloadItemsData struct {

        
            Type string  `json:"type"`
            Templates []string  `json:"templates"`
            Brand []string  `json:"brand"`
         
    }
    
    // VerifiedBy used by Catalog
    type VerifiedBy struct {

        
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
         
    }
    
    // ProductDownloadsItems used by Catalog
    type ProductDownloadsItems struct {

        
            TriggerOn string  `json:"trigger_on"`
            Status string  `json:"status"`
            Data ProductDownloadItemsData  `json:"data"`
            SellerID float64  `json:"seller_id"`
            TaskID string  `json:"task_id"`
            URL string  `json:"url"`
            TemplateTags map[string]interface{}  `json:"template_tags"`
            CreatedBy VerifiedBy  `json:"created_by"`
            CompletedOn string  `json:"completed_on"`
            ID string  `json:"id"`
         
    }
    
    // ProductDownloadsResponse used by Catalog
    type ProductDownloadsResponse struct {

        
            Page Page  `json:"page"`
            Items ProductDownloadsItems  `json:"items"`
         
    }
    
    // ProductConfigurationDownloads used by Catalog
    type ProductConfigurationDownloads struct {

        
            Multivalue bool  `json:"multivalue"`
            Data []map[string]interface{}  `json:"data"`
         
    }
    
    // Hierarchy used by Catalog
    type Hierarchy struct {

        
            L2 float64  `json:"l2"`
            L1 float64  `json:"l1"`
            Department float64  `json:"department"`
         
    }
    
    // Media2 used by Catalog
    type Media2 struct {

        
            Landscape string  `json:"landscape"`
            Portrait string  `json:"portrait"`
            Logo string  `json:"logo"`
         
    }
    
    // CategoryMappingValues used by Catalog
    type CategoryMappingValues struct {

        
            Name string  `json:"name"`
            CatalogID float64  `json:"catalog_id"`
         
    }
    
    // CategoryMapping used by Catalog
    type CategoryMapping struct {

        
            Google CategoryMappingValues  `json:"google"`
            Facebook CategoryMappingValues  `json:"facebook"`
            Ajio CategoryMappingValues  `json:"ajio"`
         
    }
    
    // CategoryRequestBody used by Catalog
    type CategoryRequestBody struct {

        
            Name string  `json:"name"`
            Level float64  `json:"level"`
            Slug string  `json:"slug"`
            Priority float64  `json:"priority"`
            Hierarchy []Hierarchy  `json:"hierarchy"`
            Tryouts []string  `json:"tryouts"`
            Media Media2  `json:"media"`
            Synonyms []string  `json:"synonyms"`
            IsActive bool  `json:"is_active"`
            Departments []float64  `json:"departments"`
            Marketplaces CategoryMapping  `json:"marketplaces"`
         
    }
    
    // CategoryCreateResponse used by Catalog
    type CategoryCreateResponse struct {

        
            Message string  `json:"message"`
            UID float64  `json:"uid"`
         
    }
    
    // Category used by Catalog
    type Category struct {

        
            CreatedBy map[string]interface{}  `json:"created_by"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            Level float64  `json:"level"`
            Slug string  `json:"slug"`
            Priority float64  `json:"priority"`
            Hierarchy []Hierarchy  `json:"hierarchy"`
            Tryouts []string  `json:"tryouts"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Media Media2  `json:"media"`
            ModifiedOn string  `json:"modified_on"`
            Synonyms []string  `json:"synonyms"`
            CreatedOn string  `json:"created_on"`
            IsActive bool  `json:"is_active"`
            Departments []float64  `json:"departments"`
            Marketplaces CategoryMapping  `json:"marketplaces"`
            ID string  `json:"id"`
         
    }
    
    // CategoryResponse used by Catalog
    type CategoryResponse struct {

        
            Page Page  `json:"page"`
            Items []Category  `json:"items"`
         
    }
    
    // SingleCategoryResponse used by Catalog
    type SingleCategoryResponse struct {

        
            Data Category  `json:"data"`
         
    }
    
    // CategoryUpdateResponse used by Catalog
    type CategoryUpdateResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // TaxIdentifier used by Catalog
    type TaxIdentifier struct {

        
            HsnCode string  `json:"hsn_code"`
            HsnCodeID string  `json:"hsn_code_id"`
            ReportingHsn string  `json:"reporting_hsn"`
         
    }
    
    // ProductPublish used by Catalog
    type ProductPublish struct {

        
            ProductOnlineDate string  `json:"product_online_date"`
            IsSet bool  `json:"is_set"`
         
    }
    
    // CustomOrder used by Catalog
    type CustomOrder struct {

        
            ManufacturingTimeUnit string  `json:"manufacturing_time_unit"`
            IsCustomOrder bool  `json:"is_custom_order"`
            ManufacturingTime float64  `json:"manufacturing_time"`
         
    }
    
    // TeaserTag used by Catalog
    type TeaserTag struct {

        
            Tag string  `json:"tag"`
            URL string  `json:"url"`
         
    }
    
    // NetQuantity used by Catalog
    type NetQuantity struct {

        
            Unit interface{}  `json:"unit"`
            Value float64  `json:"value"`
         
    }
    
    // Trader used by Catalog
    type Trader struct {

        
            Name interface{}  `json:"name"`
            Type string  `json:"type"`
            Address []string  `json:"address"`
         
    }
    
    // ProductCreateUpdateSchemaV2 used by Catalog
    type ProductCreateUpdateSchemaV2 struct {

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            TaxIdentifier TaxIdentifier  `json:"tax_identifier"`
            MultiSize bool  `json:"multi_size"`
            BulkJobID string  `json:"bulk_job_id"`
            IsActive bool  `json:"is_active"`
            Description string  `json:"description"`
            Tags []string  `json:"tags"`
            VariantGroup map[string]interface{}  `json:"variant_group"`
            ProductPublish ProductPublish  `json:"product_publish"`
            Slug string  `json:"slug"`
            Sizes []map[string]interface{}  `json:"sizes"`
            Highlights []string  `json:"highlights"`
            CustomOrder CustomOrder  `json:"custom_order"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            IsDependent bool  `json:"is_dependent"`
            CompanyID float64  `json:"company_id"`
            SizeGuide string  `json:"size_guide"`
            ItemType string  `json:"item_type"`
            Variants map[string]interface{}  `json:"variants"`
            Attributes map[string]interface{}  `json:"attributes"`
            CountryOfOrigin string  `json:"country_of_origin"`
            ItemCode string  `json:"item_code"`
            Media []Media1  `json:"media"`
            NoOfBoxes float64  `json:"no_of_boxes"`
            ShortDescription string  `json:"short_description"`
            CategorySlug string  `json:"category_slug"`
            Departments []float64  `json:"departments"`
            BrandUID float64  `json:"brand_uid"`
            TeaserTag TeaserTag  `json:"teaser_tag"`
            ProductGroupTag []string  `json:"product_group_tag"`
            ChangeRequestID interface{}  `json:"change_request_id"`
            NetQuantity NetQuantity  `json:"net_quantity"`
            Action string  `json:"action"`
            IsSet bool  `json:"is_set"`
            Requester string  `json:"requester"`
            ReturnConfig ReturnConfig  `json:"return_config"`
            VariantMedia map[string]interface{}  `json:"variant_media"`
            Trader []Trader  `json:"trader"`
            Currency string  `json:"currency"`
            TemplateTag string  `json:"template_tag"`
            IsImageLessProduct bool  `json:"is_image_less_product"`
         
    }
    
    // ProductPublished used by Catalog
    type ProductPublished struct {

        
            ProductOnlineDate float64  `json:"product_online_date"`
            IsSet bool  `json:"is_set"`
         
    }
    
    // Logo used by Catalog
    type Logo struct {

        
            AspectRatioF float64  `json:"aspect_ratio_f"`
            URL string  `json:"url"`
            AspectRatio string  `json:"aspect_ratio"`
            SecureURL string  `json:"secure_url"`
         
    }
    
    // Image used by Catalog
    type Image struct {

        
            AspectRatioF float64  `json:"aspect_ratio_f"`
            URL string  `json:"url"`
            AspectRatio string  `json:"aspect_ratio"`
            SecureURL string  `json:"secure_url"`
         
    }
    
    // NetQuantityResponse used by Catalog
    type NetQuantityResponse struct {

        
            Unit string  `json:"unit"`
            Value float64  `json:"value"`
         
    }
    
    // ReturnConfigResponse used by Catalog
    type ReturnConfigResponse struct {

        
            Time float64  `json:"time"`
            Unit string  `json:"unit"`
            Returnable bool  `json:"returnable"`
         
    }
    
    // Product used by Catalog
    type Product struct {

        
            CategoryUID float64  `json:"category_uid"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            MultiSize bool  `json:"multi_size"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            Color string  `json:"color"`
            AllCompanyIds []float64  `json:"all_company_ids"`
            Moq map[string]interface{}  `json:"moq"`
            Stage string  `json:"stage"`
            AllIdentifiers []string  `json:"all_identifiers"`
            VerifiedBy VerifiedBy  `json:"verified_by"`
            IsActive bool  `json:"is_active"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Description string  `json:"description"`
            Pending string  `json:"pending"`
            IsPhysical bool  `json:"is_physical"`
            PrimaryColor string  `json:"primary_color"`
            VariantGroup map[string]interface{}  `json:"variant_group"`
            ProductPublish ProductPublished  `json:"product_publish"`
            Slug string  `json:"slug"`
            Sizes []map[string]interface{}  `json:"sizes"`
            Highlights []string  `json:"highlights"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            IsDependent bool  `json:"is_dependent"`
            CreatedOn string  `json:"created_on"`
            SizeGuide string  `json:"size_guide"`
            CompanyID float64  `json:"company_id"`
            ItemType string  `json:"item_type"`
            Variants map[string]interface{}  `json:"variants"`
            AllSizes []map[string]interface{}  `json:"all_sizes"`
            Attributes map[string]interface{}  `json:"attributes"`
            VerifiedOn string  `json:"verified_on"`
            IsImageLessProduct bool  `json:"is_image_less_product"`
            ImageNature string  `json:"image_nature"`
            CountryOfOrigin string  `json:"country_of_origin"`
            ItemCode string  `json:"item_code"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Media []Media1  `json:"media"`
            NoOfBoxes float64  `json:"no_of_boxes"`
            CategorySlug string  `json:"category_slug"`
            ShortDescription string  `json:"short_description"`
            Brand Brand  `json:"brand"`
            IsExpirable bool  `json:"is_expirable"`
            Departments []float64  `json:"departments"`
            BrandUID float64  `json:"brand_uid"`
            L3Mapping []string  `json:"l3_mapping"`
            HsnCode string  `json:"hsn_code"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            ProductGroupTag []string  `json:"product_group_tag"`
            Images []Image  `json:"images"`
            NetQuantity NetQuantityResponse  `json:"net_quantity"`
            IsSet bool  `json:"is_set"`
            Trader []map[string]interface{}  `json:"trader"`
            ModifiedOn string  `json:"modified_on"`
            ReturnConfig ReturnConfigResponse  `json:"return_config"`
            VariantMedia map[string]interface{}  `json:"variant_media"`
            Tags []string  `json:"tags"`
            Currency string  `json:"currency"`
            Category map[string]interface{}  `json:"category"`
            TemplateTag string  `json:"template_tag"`
            ID string  `json:"id"`
         
    }
    
    // ProductListingResponse used by Catalog
    type ProductListingResponse struct {

        
            Page Page  `json:"page"`
            Items []Product  `json:"items"`
         
    }
    
    // ProductVariants used by Catalog
    type ProductVariants struct {

        
            CategoryUID float64  `json:"category_uid"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            ItemCode string  `json:"item_code"`
            Media []Media1  `json:"media"`
            BrandUID float64  `json:"brand_uid"`
         
    }
    
    // ProductVariantsResponse used by Catalog
    type ProductVariantsResponse struct {

        
            Page Page  `json:"page"`
            Variants []ProductVariants  `json:"variants"`
         
    }
    
    // AttributeMasterSerializer used by Catalog
    type AttributeMasterSerializer struct {

        
            Name string  `json:"name"`
            EnabledForEndConsumer bool  `json:"enabled_for_end_consumer"`
            Schema AttributeMaster  `json:"schema"`
            Filters AttributeMasterFilter  `json:"filters"`
            Logo string  `json:"logo"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Description string  `json:"description"`
            Slug string  `json:"slug"`
            Suggestion string  `json:"suggestion"`
            RawKey string  `json:"raw_key"`
            Synonyms map[string]interface{}  `json:"synonyms"`
            CreatedOn string  `json:"created_on"`
            IsNested bool  `json:"is_nested"`
            Variant bool  `json:"variant"`
            Details AttributeMasterDetails  `json:"details"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Departments []string  `json:"departments"`
            ModifiedOn string  `json:"modified_on"`
            Tags []string  `json:"tags"`
            Unit string  `json:"unit"`
         
    }
    
    // ProductAttributesResponse used by Catalog
    type ProductAttributesResponse struct {

        
            Items []AttributeMasterSerializer  `json:"items"`
         
    }
    
    // SingleProductResponse used by Catalog
    type SingleProductResponse struct {

        
            Data Product  `json:"data"`
         
    }
    
    // ProductCreateUpdate used by Catalog
    type ProductCreateUpdate struct {

        
            UID float64  `json:"uid"`
            Name interface{}  `json:"name"`
            TaxIdentifier TaxIdentifier  `json:"tax_identifier"`
            MultiSize bool  `json:"multi_size"`
            BulkJobID string  `json:"bulk_job_id"`
            IsActive bool  `json:"is_active"`
            Description string  `json:"description"`
            Tags []string  `json:"tags"`
            VariantGroup map[string]interface{}  `json:"variant_group"`
            ProductPublish ProductPublish  `json:"product_publish"`
            Slug string  `json:"slug"`
            Highlights []string  `json:"highlights"`
            CustomOrder CustomOrder  `json:"custom_order"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            IsDependent bool  `json:"is_dependent"`
            CompanyID float64  `json:"company_id"`
            SizeGuide string  `json:"size_guide"`
            ItemType string  `json:"item_type"`
            Variants map[string]interface{}  `json:"variants"`
            CountryOfOrigin string  `json:"country_of_origin"`
            ItemCode interface{}  `json:"item_code"`
            Media []Media1  `json:"media"`
            NoOfBoxes float64  `json:"no_of_boxes"`
            ShortDescription string  `json:"short_description"`
            CategorySlug string  `json:"category_slug"`
            Departments []float64  `json:"departments"`
            BrandUID float64  `json:"brand_uid"`
            TeaserTag TeaserTag  `json:"teaser_tag"`
            ProductGroupTag []string  `json:"product_group_tag"`
            ChangeRequestID interface{}  `json:"change_request_id"`
            NetQuantity NetQuantity  `json:"net_quantity"`
            Action string  `json:"action"`
            IsSet bool  `json:"is_set"`
            Requester string  `json:"requester"`
            ReturnConfig ReturnConfig  `json:"return_config"`
            VariantMedia map[string]interface{}  `json:"variant_media"`
            Trader []Trader  `json:"trader"`
            Currency string  `json:"currency"`
            TemplateTag string  `json:"template_tag"`
            IsImageLessProduct bool  `json:"is_image_less_product"`
         
    }
    
    // ValidateIdentifier used by Catalog
    type ValidateIdentifier struct {

        
            Primary bool  `json:"primary"`
            GtinType string  `json:"gtin_type"`
            GtinValue string  `json:"gtin_value"`
         
    }
    
    // AllSizes used by Catalog
    type AllSizes struct {

        
            ItemLength float64  `json:"item_length"`
            ItemWeight float64  `json:"item_weight"`
            ItemWidth float64  `json:"item_width"`
            ItemWeightUnitOfMeasure interface{}  `json:"item_weight_unit_of_measure"`
            ItemHeight float64  `json:"item_height"`
            Size interface{}  `json:"size"`
            ItemDimensionsUnitOfMeasure string  `json:"item_dimensions_unit_of_measure"`
            Identifiers []ValidateIdentifier  `json:"identifiers"`
         
    }
    
    // GetAllSizes used by Catalog
    type GetAllSizes struct {

        
            AllSizes []AllSizes  `json:"all_sizes"`
         
    }
    
    // ValidateProduct used by Catalog
    type ValidateProduct struct {

        
            Valid bool  `json:"valid"`
         
    }
    
    // UserInfo1 used by Catalog
    type UserInfo1 struct {

        
            UID string  `json:"uid"`
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
            Email string  `json:"email"`
         
    }
    
    // BulkJob used by Catalog
    type BulkJob struct {

        
            Failed float64  `json:"failed"`
            Succeed float64  `json:"succeed"`
            FilePath string  `json:"file_path"`
            TrackingURL string  `json:"tracking_url"`
            CustomTemplateTag string  `json:"custom_template_tag"`
            Total float64  `json:"total"`
            Cancelled float64  `json:"cancelled"`
            Stage string  `json:"stage"`
            ModifiedBy UserInfo1  `json:"modified_by"`
            ModifiedOn string  `json:"modified_on"`
            CancelledRecords []map[string]interface{}  `json:"cancelled_records"`
            FailedRecords []map[string]interface{}  `json:"failed_records"`
            CreatedOn string  `json:"created_on"`
            CompanyID float64  `json:"company_id"`
            IsActive bool  `json:"is_active"`
            CreatedBy UserInfo1  `json:"created_by"`
            TemplateTag string  `json:"template_tag"`
         
    }
    
    // BulkResponse used by Catalog
    type BulkResponse struct {

        
            ModifiedBy UserInfo1  `json:"modified_by"`
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
            IsActive bool  `json:"is_active"`
            CreatedBy UserInfo1  `json:"created_by"`
            BatchID string  `json:"batch_id"`
         
    }
    
    // UserDetail1 used by Catalog
    type UserDetail1 struct {

        
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
            FullName string  `json:"full_name"`
         
    }
    
    // ProductBulkRequest used by Catalog
    type ProductBulkRequest struct {

        
            Failed float64  `json:"failed"`
            FilePath string  `json:"file_path"`
            Succeed float64  `json:"succeed"`
            Total float64  `json:"total"`
            Cancelled float64  `json:"cancelled"`
            Stage string  `json:"stage"`
            CancelledRecords []string  `json:"cancelled_records"`
            ModifiedOn string  `json:"modified_on"`
            FailedRecords []string  `json:"failed_records"`
            ModifiedBy UserDetail1  `json:"modified_by"`
            CreatedOn string  `json:"created_on"`
            CompanyID float64  `json:"company_id"`
            Template ProductTemplate  `json:"template"`
            IsActive bool  `json:"is_active"`
            CreatedBy UserDetail1  `json:"created_by"`
            TemplateTag string  `json:"template_tag"`
         
    }
    
    // ProductBulkRequestList used by Catalog
    type ProductBulkRequestList struct {

        
            Page Page  `json:"page"`
            Items ProductBulkRequest  `json:"items"`
         
    }
    
    // BulkProductRequest used by Catalog
    type BulkProductRequest struct {

        
            Data []map[string]interface{}  `json:"data"`
            TemplateTag string  `json:"template_tag"`
            CompanyID float64  `json:"company_id"`
            BatchID string  `json:"batch_id"`
         
    }
    
    // NestedTags used by Catalog
    type NestedTags struct {

        
            Tags []string  `json:"tags"`
         
    }
    
    // ProductTagsViewResponse used by Catalog
    type ProductTagsViewResponse struct {

        
            Items NestedTags  `json:"items"`
         
    }
    
    // ProductBulkAssets used by Catalog
    type ProductBulkAssets struct {

        
            URL string  `json:"url"`
            CompanyID float64  `json:"company_id"`
            User map[string]interface{}  `json:"user"`
         
    }
    
    // UserCommon used by Catalog
    type UserCommon struct {

        
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
            CompanyID float64  `json:"company_id"`
         
    }
    
    // Items used by Catalog
    type Items struct {

        
            Failed float64  `json:"failed"`
            IsActive bool  `json:"is_active"`
            Retry float64  `json:"retry"`
            FilePath string  `json:"file_path"`
            TrackingURL string  `json:"tracking_url"`
            Total float64  `json:"total"`
            Cancelled float64  `json:"cancelled"`
            Stage string  `json:"stage"`
            ModifiedBy UserCommon  `json:"modified_by"`
            ModifiedOn string  `json:"modified_on"`
            CancelledRecords []string  `json:"cancelled_records"`
            FailedRecords []string  `json:"failed_records"`
            CreatedOn string  `json:"created_on"`
            CompanyID float64  `json:"company_id"`
            Succeed float64  `json:"succeed"`
            CreatedBy UserCommon  `json:"created_by"`
            ID string  `json:"id"`
         
    }
    
    // BulkAssetResponse used by Catalog
    type BulkAssetResponse struct {

        
            Page Page  `json:"page"`
            Items []Items  `json:"items"`
         
    }
    
    // ProductSizeDeleteDataResponse used by Catalog
    type ProductSizeDeleteDataResponse struct {

        
            Size string  `json:"size"`
            CompanyID float64  `json:"company_id"`
            ItemID float64  `json:"item_id"`
         
    }
    
    // ProductSizeDeleteResponse used by Catalog
    type ProductSizeDeleteResponse struct {

        
            Success bool  `json:"success"`
            Data ProductSizeDeleteDataResponse  `json:"data"`
         
    }
    
    // ItemQuery used by Catalog
    type ItemQuery struct {

        
            ItemCode string  `json:"item_code"`
            BrandUID float64  `json:"brand_uid"`
            UID float64  `json:"uid"`
         
    }
    
    // SetSize used by Catalog
    type SetSize struct {

        
            Size string  `json:"size"`
            Pieces float64  `json:"pieces"`
         
    }
    
    // SizeDistribution used by Catalog
    type SizeDistribution struct {

        
            Sizes []SetSize  `json:"sizes"`
         
    }
    
    // InventorySet used by Catalog
    type InventorySet struct {

        
            Name string  `json:"name"`
            Quantity float64  `json:"quantity"`
            SizeDistribution SizeDistribution  `json:"size_distribution"`
         
    }
    
    // GTIN used by Catalog
    type GTIN struct {

        
            Primary bool  `json:"primary"`
            GtinType string  `json:"gtin_type"`
            GtinValue interface{}  `json:"gtin_value"`
         
    }
    
    // InvSize used by Catalog
    type InvSize struct {

        
            Quantity float64  `json:"quantity"`
            ItemLength float64  `json:"item_length"`
            ItemWeight float64  `json:"item_weight"`
            ItemWidth float64  `json:"item_width"`
            Price float64  `json:"price"`
            PriceTransfer float64  `json:"price_transfer"`
            ItemWeightUnitOfMeasure string  `json:"item_weight_unit_of_measure"`
            IsSet bool  `json:"is_set"`
            StoreCode string  `json:"store_code"`
            ItemHeight float64  `json:"item_height"`
            Size interface{}  `json:"size"`
            ItemDimensionsUnitOfMeasure string  `json:"item_dimensions_unit_of_measure"`
            ExpirationDate string  `json:"expiration_date"`
            PriceEffective float64  `json:"price_effective"`
            Set InventorySet  `json:"set"`
            Currency string  `json:"currency"`
            Identifiers []GTIN  `json:"identifiers"`
         
    }
    
    // InventoryRequest used by Catalog
    type InventoryRequest struct {

        
            Item ItemQuery  `json:"item"`
            Sizes []InvSize  `json:"sizes"`
            CompanyID float64  `json:"company_id"`
         
    }
    
    // InventoryResponse used by Catalog
    type InventoryResponse struct {

        
            Quantity float64  `json:"quantity"`
            InventoryUpdatedOn string  `json:"inventory_updated_on"`
            SellerIdentifier string  `json:"seller_identifier"`
            Store map[string]interface{}  `json:"store"`
            SellableQuantity float64  `json:"sellable_quantity"`
            Price float64  `json:"price"`
            PriceTransfer float64  `json:"price_transfer"`
            UID string  `json:"uid"`
            ItemID float64  `json:"item_id"`
            Size string  `json:"size"`
            PriceEffective float64  `json:"price_effective"`
            Currency string  `json:"currency"`
            Identifiers map[string]interface{}  `json:"identifiers"`
         
    }
    
    // InventoryResponsePaginated used by Catalog
    type InventoryResponsePaginated struct {

        
            Page Page  `json:"page"`
            Items []InventoryResponse  `json:"items"`
         
    }
    
    // ManufacturerResponse used by Catalog
    type ManufacturerResponse struct {

        
            Name string  `json:"name"`
            Address string  `json:"address"`
            IsDefault bool  `json:"is_default"`
         
    }
    
    // BrandMeta used by Catalog
    type BrandMeta struct {

        
            Name string  `json:"name"`
            ID float64  `json:"id"`
         
    }
    
    // DimensionResponse used by Catalog
    type DimensionResponse struct {

        
            IsDefault bool  `json:"is_default"`
            Height float64  `json:"height"`
            Length float64  `json:"length"`
            Unit string  `json:"unit"`
            Width float64  `json:"width"`
         
    }
    
    // WeightResponse used by Catalog
    type WeightResponse struct {

        
            Unit string  `json:"unit"`
            Shipping float64  `json:"shipping"`
            IsDefault bool  `json:"is_default"`
         
    }
    
    // PriceMeta used by Catalog
    type PriceMeta struct {

        
            UpdatedAt string  `json:"updated_at"`
            Effective float64  `json:"effective"`
            TpNotes map[string]interface{}  `json:"tp_notes"`
            Marked float64  `json:"marked"`
            Currency string  `json:"currency"`
            Transfer float64  `json:"transfer"`
         
    }
    
    // QuantityBase used by Catalog
    type QuantityBase struct {

        
            UpdatedAt string  `json:"updated_at"`
            Count float64  `json:"count"`
         
    }
    
    // Quantities used by Catalog
    type Quantities struct {

        
            NotAvailable QuantityBase  `json:"not_available"`
            OrderCommitted QuantityBase  `json:"order_committed"`
            Sellable QuantityBase  `json:"sellable"`
            Damaged QuantityBase  `json:"damaged"`
         
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
    
    // Trader1 used by Catalog
    type Trader1 struct {

        
            Name string  `json:"name"`
            Type string  `json:"type"`
            Address []string  `json:"address"`
         
    }
    
    // InventorySellerResponse used by Catalog
    type InventorySellerResponse struct {

        
            SellerIdentifier string  `json:"seller_identifier"`
            Manufacturer ManufacturerResponse  `json:"manufacturer"`
            UID string  `json:"uid"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            Stage string  `json:"stage"`
            Identifier map[string]interface{}  `json:"identifier"`
            IsActive bool  `json:"is_active"`
            FyndMeta map[string]interface{}  `json:"fynd_meta"`
            CreatedBy UserSerializer  `json:"created_by"`
            TraceID string  `json:"trace_id"`
            Tags []string  `json:"tags"`
            AddedOnStore string  `json:"added_on_store"`
            ItemID float64  `json:"item_id"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Size string  `json:"size"`
            RawMeta map[string]interface{}  `json:"raw_meta"`
            Store StoreMeta  `json:"store"`
            Meta map[string]interface{}  `json:"meta"`
            CountryOfOrigin string  `json:"country_of_origin"`
            FyndArticleCode string  `json:"fynd_article_code"`
            TrackInventory bool  `json:"track_inventory"`
            Fragile bool  `json:"fragile"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            Brand BrandMeta  `json:"brand"`
            Set InventorySet  `json:"set"`
            Dimension DimensionResponse  `json:"dimension"`
            TotalQuantity float64  `json:"total_quantity"`
            Weight WeightResponse  `json:"weight"`
            Price PriceMeta  `json:"price"`
            Quantities Quantities  `json:"quantities"`
            FyndItemCode string  `json:"fynd_item_code"`
            Company CompanyMeta  `json:"company"`
            IsSet bool  `json:"is_set"`
            ReturnConfig ReturnConfig1  `json:"return_config"`
            ExpirationDate string  `json:"expiration_date"`
            Trader []Trader1  `json:"trader"`
         
    }
    
    // InventorySellerIdentifierResponsePaginated used by Catalog
    type InventorySellerIdentifierResponsePaginated struct {

        
            Page Page  `json:"page"`
            Items []InventorySellerResponse  `json:"items"`
         
    }
    
    // BulkInventoryGetItems used by Catalog
    type BulkInventoryGetItems struct {

        
            Failed float64  `json:"failed"`
            FilePath string  `json:"file_path"`
            Succeed float64  `json:"succeed"`
            Total float64  `json:"total"`
            Cancelled float64  `json:"cancelled"`
            Stage string  `json:"stage"`
            CancelledRecords []string  `json:"cancelled_records"`
            FailedRecords []string  `json:"failed_records"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
            CompanyID float64  `json:"company_id"`
            IsActive bool  `json:"is_active"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ID string  `json:"id"`
         
    }
    
    // BulkInventoryGet used by Catalog
    type BulkInventoryGet struct {

        
            Page Page  `json:"page"`
            Items []BulkInventoryGetItems  `json:"items"`
         
    }
    
    // InventoryJobPayload used by Catalog
    type InventoryJobPayload struct {

        
            Quantity float64  `json:"quantity"`
            SellerIdentifier string  `json:"seller_identifier"`
            Price float64  `json:"price"`
            PriceMarked float64  `json:"price_marked"`
            ItemWeightUnitOfMeasure string  `json:"item_weight_unit_of_measure"`
            TraceID string  `json:"trace_id"`
            StoreCode string  `json:"store_code"`
            ItemDimensionsUnitOfMeasure string  `json:"item_dimensions_unit_of_measure"`
            ExpirationDate string  `json:"expiration_date"`
            Tags []string  `json:"tags"`
            PriceEffective float64  `json:"price_effective"`
            Currency string  `json:"currency"`
            TotalQuantity float64  `json:"total_quantity"`
         
    }
    
    // InventoryBulkRequest used by Catalog
    type InventoryBulkRequest struct {

        
            User map[string]interface{}  `json:"user"`
            Sizes []InventoryJobPayload  `json:"sizes"`
            CompanyID float64  `json:"company_id"`
            BatchID string  `json:"batch_id"`
         
    }
    
    // InventoryExportRequest used by Catalog
    type InventoryExportRequest struct {

        
            Type string  `json:"type"`
            Store []float64  `json:"store"`
            Brand []float64  `json:"brand"`
         
    }
    
    // InventoryExportResponse used by Catalog
    type InventoryExportResponse struct {

        
            TriggerOn string  `json:"trigger_on"`
            Status string  `json:"status"`
            RequestParams map[string]interface{}  `json:"request_params"`
            SellerID float64  `json:"seller_id"`
            TaskID string  `json:"task_id"`
         
    }
    
    // InventoryExportJob used by Catalog
    type InventoryExportJob struct {

        
            TriggerOn string  `json:"trigger_on"`
            Status string  `json:"status"`
            RequestParams map[string]interface{}  `json:"request_params"`
            SellerID float64  `json:"seller_id"`
            TaskID string  `json:"task_id"`
            URL string  `json:"url"`
            CompletedOn string  `json:"completed_on"`
         
    }
    
    // ManufacturerResponse1 used by Catalog
    type ManufacturerResponse1 struct {

        
            Name string  `json:"name"`
            Address string  `json:"address"`
            IsDefault bool  `json:"is_default"`
         
    }
    
    // ArticleStoreResponse used by Catalog
    type ArticleStoreResponse struct {

        
            Name string  `json:"name"`
            StoreCode string  `json:"store_code"`
            StoreType string  `json:"store_type"`
            UID float64  `json:"uid"`
         
    }
    
    // BrandMeta1 used by Catalog
    type BrandMeta1 struct {

        
            Name string  `json:"name"`
            ID float64  `json:"id"`
         
    }
    
    // DimensionResponse1 used by Catalog
    type DimensionResponse1 struct {

        
            Unit string  `json:"unit"`
            Width float64  `json:"width"`
            Length float64  `json:"length"`
            Height float64  `json:"height"`
         
    }
    
    // WeightResponse1 used by Catalog
    type WeightResponse1 struct {

        
            Unit string  `json:"unit"`
            Shipping float64  `json:"shipping"`
         
    }
    
    // PriceArticle used by Catalog
    type PriceArticle struct {

        
            Effective float64  `json:"effective"`
            TpNotes map[string]interface{}  `json:"tp_notes"`
            Marked float64  `json:"marked"`
            Currency string  `json:"currency"`
            Transfer float64  `json:"transfer"`
         
    }
    
    // Quantity used by Catalog
    type Quantity struct {

        
            Count float64  `json:"count"`
         
    }
    
    // QuantitiesArticle used by Catalog
    type QuantitiesArticle struct {

        
            NotAvailable Quantity  `json:"not_available"`
            OrderCommitted Quantity  `json:"order_committed"`
            Sellable Quantity  `json:"sellable"`
            Damaged Quantity  `json:"damaged"`
         
    }
    
    // CompanyMeta1 used by Catalog
    type CompanyMeta1 struct {

        
            ID float64  `json:"id"`
         
    }
    
    // Trader2 used by Catalog
    type Trader2 struct {

        
            Name string  `json:"name"`
            Type string  `json:"type"`
            Address []string  `json:"address"`
         
    }
    
    // ReturnConfig2 used by Catalog
    type ReturnConfig2 struct {

        
            Time float64  `json:"time"`
            Unit string  `json:"unit"`
            Returnable bool  `json:"returnable"`
         
    }
    
    // GetInventories used by Catalog
    type GetInventories struct {

        
            SellerIdentifier string  `json:"seller_identifier"`
            Manufacturer ManufacturerResponse1  `json:"manufacturer"`
            UID string  `json:"uid"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            Stage string  `json:"stage"`
            Identifier map[string]interface{}  `json:"identifier"`
            TraceID string  `json:"trace_id"`
            CreatedBy UserSerializer  `json:"created_by"`
            ItemID float64  `json:"item_id"`
            Size string  `json:"size"`
            DateMeta DateMeta  `json:"date_meta"`
            InventoryUpdatedOn string  `json:"inventory_updated_on"`
            Store ArticleStoreResponse  `json:"store"`
            TrackInventory bool  `json:"track_inventory"`
            CountryOfOrigin string  `json:"country_of_origin"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            Brand BrandMeta1  `json:"brand"`
            Dimension DimensionResponse1  `json:"dimension"`
            Weight WeightResponse1  `json:"weight"`
            TotalQuantity float64  `json:"total_quantity"`
            Platforms map[string]interface{}  `json:"platforms"`
            Price PriceArticle  `json:"price"`
            Quantities QuantitiesArticle  `json:"quantities"`
            Company CompanyMeta1  `json:"company"`
            IsSet bool  `json:"is_set"`
            Trader []Trader2  `json:"trader"`
            ReturnConfig ReturnConfig2  `json:"return_config"`
            ExpirationDate string  `json:"expiration_date"`
            Tags []string  `json:"tags"`
            ID string  `json:"id"`
         
    }
    
    // GetInventoriesResponse used by Catalog
    type GetInventoriesResponse struct {

        
            Page Page  `json:"page"`
            Items []GetInventories  `json:"items"`
         
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

        
            SellerIdentifier string  `json:"seller_identifier"`
            PriceMarked float64  `json:"price_marked"`
            TraceID string  `json:"trace_id"`
            ExpirationDate string  `json:"expiration_date"`
            Tags []string  `json:"tags"`
            PriceEffective float64  `json:"price_effective"`
            TotalQuantity float64  `json:"total_quantity"`
            StoreID float64  `json:"store_id"`
         
    }
    
    // InventoryRequestSchemaV2 used by Catalog
    type InventoryRequestSchemaV2 struct {

        
            Meta map[string]interface{}  `json:"meta"`
            CompanyID float64  `json:"company_id"`
            Payload []InventoryPayload  `json:"payload"`
         
    }
    
    // InventoryFailedReason used by Catalog
    type InventoryFailedReason struct {

        
            Errors string  `json:"errors"`
            Message string  `json:"message"`
         
    }
    
    // InventoryResponseItem used by Catalog
    type InventoryResponseItem struct {

        
            Data InventoryPayload  `json:"data"`
            Reason InventoryFailedReason  `json:"reason"`
         
    }
    
    // InventoryUpdateResponse used by Catalog
    type InventoryUpdateResponse struct {

        
            Items []InventoryResponseItem  `json:"items"`
            Message string  `json:"message"`
         
    }
    
    // HsnUpsert used by Catalog
    type HsnUpsert struct {

        
            HsnCode string  `json:"hsn_code"`
            Threshold1 float64  `json:"threshold1"`
            UID float64  `json:"uid"`
            Tax1 float64  `json:"tax1"`
            Tax2 float64  `json:"tax2"`
            TaxOnEsp bool  `json:"tax_on_esp"`
            Threshold2 float64  `json:"threshold2"`
            CompanyID float64  `json:"company_id"`
            TaxOnMrp bool  `json:"tax_on_mrp"`
            IsActive bool  `json:"is_active"`
            Hs2Code string  `json:"hs2_code"`
         
    }
    
    // HsnCodesObject used by Catalog
    type HsnCodesObject struct {

        
            HsnCode string  `json:"hsn_code"`
            Threshold1 float64  `json:"threshold1"`
            Tax1 float64  `json:"tax1"`
            Tax2 float64  `json:"tax2"`
            ModifiedOn string  `json:"modified_on"`
            TaxOnEsp bool  `json:"tax_on_esp"`
            Threshold2 float64  `json:"threshold2"`
            CompanyID float64  `json:"company_id"`
            TaxOnMrp bool  `json:"tax_on_mrp"`
            Hs2Code string  `json:"hs2_code"`
            ID string  `json:"id"`
         
    }
    
    // HsnCode used by Catalog
    type HsnCode struct {

        
            Data HsnCodesObject  `json:"data"`
         
    }
    
    // PageResponse used by Catalog
    type PageResponse struct {

        
            Current string  `json:"current"`
            HasNext bool  `json:"has_next"`
            Size float64  `json:"size"`
            ItemTotal float64  `json:"item_total"`
            HasPrevious bool  `json:"has_previous"`
         
    }
    
    // HsnCodesListingResponse used by Catalog
    type HsnCodesListingResponse struct {

        
            Page PageResponse  `json:"page"`
            Items []HsnCodesObject  `json:"items"`
         
    }
    
    // BulkHsnUpsert used by Catalog
    type BulkHsnUpsert struct {

        
            Data []HsnUpsert  `json:"data"`
         
    }
    
    // BulkHsnResponse used by Catalog
    type BulkHsnResponse struct {

        
            Success bool  `json:"success"`
         
    }
    
    // TaxSlab used by Catalog
    type TaxSlab struct {

        
            Rate float64  `json:"rate"`
            Threshold float64  `json:"threshold"`
            Cess float64  `json:"cess"`
            EffectiveDate string  `json:"effective_date"`
         
    }
    
    // HSNDataInsertV2 used by Catalog
    type HSNDataInsertV2 struct {

        
            HsnCode string  `json:"hsn_code"`
            CountryCode string  `json:"country_code"`
            Taxes []TaxSlab  `json:"taxes"`
            Type string  `json:"type"`
            ReportingHsn string  `json:"reporting_hsn"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Description string  `json:"description"`
         
    }
    
    // HsnCodesListingResponseSchemaV2 used by Catalog
    type HsnCodesListingResponseSchemaV2 struct {

        
            Page PageResponse  `json:"page"`
            Items []HSNDataInsertV2  `json:"items"`
         
    }
    
    // BrandItem used by Catalog
    type BrandItem struct {

        
            Action Action  `json:"action"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Logo Media  `json:"logo"`
            Departments []string  `json:"departments"`
            Discount string  `json:"discount"`
            Banners ImageUrls  `json:"banners"`
         
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
            Slug string  `json:"slug"`
            Logo Media  `json:"logo"`
            PriorityOrder float64  `json:"priority_order"`
         
    }
    
    // DepartmentResponse used by Catalog
    type DepartmentResponse struct {

        
            Items []Department  `json:"items"`
         
    }
    
    // ThirdLevelChild used by Catalog
    type ThirdLevelChild struct {

        
            Childs []map[string]interface{}  `json:"childs"`
            Action Action  `json:"action"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Banners ImageUrls  `json:"banners"`
         
    }
    
    // SecondLevelChild used by Catalog
    type SecondLevelChild struct {

        
            Childs []ThirdLevelChild  `json:"childs"`
            Action Action  `json:"action"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Banners ImageUrls  `json:"banners"`
         
    }
    
    // Child used by Catalog
    type Child struct {

        
            Childs []SecondLevelChild  `json:"childs"`
            Action Action  `json:"action"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Banners ImageUrls  `json:"banners"`
         
    }
    
    // CategoryItems used by Catalog
    type CategoryItems struct {

        
            Childs []Child  `json:"childs"`
            Action Action  `json:"action"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Banners ImageUrls  `json:"banners"`
         
    }
    
    // DepartmentCategoryTree used by Catalog
    type DepartmentCategoryTree struct {

        
            Items []CategoryItems  `json:"items"`
            Department string  `json:"department"`
         
    }
    
    // DepartmentIdentifier used by Catalog
    type DepartmentIdentifier struct {

        
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
         
    }
    
    // CategoryListingResponse used by Catalog
    type CategoryListingResponse struct {

        
            Data []DepartmentCategoryTree  `json:"data"`
            Departments []DepartmentIdentifier  `json:"departments"`
         
    }
    
    // ApplicationProductListingResponse used by Catalog
    type ApplicationProductListingResponse struct {

        
            Page Page  `json:"page"`
            Operators map[string]interface{}  `json:"operators"`
            SortOn []ProductSortOn  `json:"sort_on"`
            Filters []ProductFilters  `json:"filters"`
            Items []ProductListingDetail  `json:"items"`
         
    }
    
    // ProductDetail used by Catalog
    type ProductDetail struct {

        
            Similars []string  `json:"similars"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            Color string  `json:"color"`
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
            Description string  `json:"description"`
            Type string  `json:"type"`
            Slug string  `json:"slug"`
            Tryouts []string  `json:"tryouts"`
            Highlights []string  `json:"highlights"`
            ProductOnlineDate string  `json:"product_online_date"`
            ItemType string  `json:"item_type"`
            Attributes map[string]interface{}  `json:"attributes"`
            RatingCount float64  `json:"rating_count"`
            ImageNature string  `json:"image_nature"`
            Rating float64  `json:"rating"`
            ItemCode string  `json:"item_code"`
            Brand ProductBrand  `json:"brand"`
            ShortDescription string  `json:"short_description"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            PromoMeta map[string]interface{}  `json:"promo_meta"`
            Medias []Media1  `json:"medias"`
            HasVariant bool  `json:"has_variant"`
         
    }
    
    // InventoryPage used by Catalog
    type InventoryPage struct {

        
            Type string  `json:"type"`
            HasNext bool  `json:"has_next"`
            NextID string  `json:"next_id"`
            ItemTotal float64  `json:"item_total"`
            HasPrevious bool  `json:"has_previous"`
         
    }
    
    // InventoryStockResponse used by Catalog
    type InventoryStockResponse struct {

        
            Page InventoryPage  `json:"page"`
            Items []map[string]interface{}  `json:"items"`
         
    }
    
    // LocationIntegrationType used by Catalog
    type LocationIntegrationType struct {

        
            Order string  `json:"order"`
            Inventory string  `json:"inventory"`
         
    }
    
    // UserSerializer1 used by Catalog
    type UserSerializer1 struct {

        
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
            Contact string  `json:"contact"`
         
    }
    
    // GetAddressSerializer used by Catalog
    type GetAddressSerializer struct {

        
            Longitude float64  `json:"longitude"`
            Address2 string  `json:"address2"`
            CountryCode string  `json:"country_code"`
            AddressType string  `json:"address_type"`
            State string  `json:"state"`
            Address1 string  `json:"address1"`
            Latitude float64  `json:"latitude"`
            Pincode float64  `json:"pincode"`
            City string  `json:"city"`
            Landmark string  `json:"landmark"`
            Country string  `json:"country"`
         
    }
    
    // InvoiceCredSerializer used by Catalog
    type InvoiceCredSerializer struct {

        
            Enabled bool  `json:"enabled"`
            Username string  `json:"username"`
            Password string  `json:"password"`
         
    }
    
    // InvoiceDetailsSerializer used by Catalog
    type InvoiceDetailsSerializer struct {

        
            EInvoice InvoiceCredSerializer  `json:"e_invoice"`
            EWaybill InvoiceCredSerializer  `json:"e_waybill"`
         
    }
    
    // ProductReturnConfigSerializer used by Catalog
    type ProductReturnConfigSerializer struct {

        
            StoreUID float64  `json:"store_uid"`
            OnSameStore bool  `json:"on_same_store"`
         
    }
    
    // SellerPhoneNumber used by Catalog
    type SellerPhoneNumber struct {

        
            Number string  `json:"number"`
            CountryCode float64  `json:"country_code"`
         
    }
    
    // UserSerializer2 used by Catalog
    type UserSerializer2 struct {

        
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
            Contact string  `json:"contact"`
         
    }
    
    // GetCompanySerializer used by Catalog
    type GetCompanySerializer struct {

        
            CompanyType string  `json:"company_type"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            BusinessType string  `json:"business_type"`
            Stage string  `json:"stage"`
            ModifiedBy UserSerializer2  `json:"modified_by"`
            ModifiedOn string  `json:"modified_on"`
            VerifiedBy UserSerializer2  `json:"verified_by"`
            CreatedOn string  `json:"created_on"`
            RejectReason string  `json:"reject_reason"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            CreatedBy UserSerializer2  `json:"created_by"`
            VerifiedOn string  `json:"verified_on"`
         
    }
    
    // LocationManagerSerializer used by Catalog
    type LocationManagerSerializer struct {

        
            Name string  `json:"name"`
            MobileNo SellerPhoneNumber  `json:"mobile_no"`
            Email string  `json:"email"`
         
    }
    
    // LocationTimingSerializer used by Catalog
    type LocationTimingSerializer struct {

        
            Minute float64  `json:"minute"`
            Hour float64  `json:"hour"`
         
    }
    
    // LocationDayWiseSerializer used by Catalog
    type LocationDayWiseSerializer struct {

        
            Opening LocationTimingSerializer  `json:"opening"`
            Open bool  `json:"open"`
            Weekday string  `json:"weekday"`
            Closing LocationTimingSerializer  `json:"closing"`
         
    }
    
    // GetLocationSerializer used by Catalog
    type GetLocationSerializer struct {

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            IntegrationType LocationIntegrationType  `json:"integration_type"`
            Stage string  `json:"stage"`
            VerifiedBy UserSerializer1  `json:"verified_by"`
            CreatedBy UserSerializer1  `json:"created_by"`
            DisplayName string  `json:"display_name"`
            Code string  `json:"code"`
            Address GetAddressSerializer  `json:"address"`
            GstCredentials InvoiceDetailsSerializer  `json:"gst_credentials"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            CreatedOn string  `json:"created_on"`
            VerifiedOn string  `json:"verified_on"`
            Documents []Document  `json:"documents"`
            Warnings map[string]interface{}  `json:"warnings"`
            StoreType string  `json:"store_type"`
            ModifiedBy UserSerializer1  `json:"modified_by"`
            ProductReturnConfig ProductReturnConfigSerializer  `json:"product_return_config"`
            NotificationEmails []string  `json:"notification_emails"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            Company GetCompanySerializer  `json:"company"`
            ModifiedOn string  `json:"modified_on"`
            Manager LocationManagerSerializer  `json:"manager"`
            Timing []LocationDayWiseSerializer  `json:"timing"`
            PhoneNumber string  `json:"phone_number"`
         
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
            AppID string  `json:"app_id"`
            Name string  `json:"name"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Logo string  `json:"logo"`
            IsActive bool  `json:"is_active"`
         
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
    

    
    // Website used by CompanyProfile
    type Website struct {

        
            URL string  `json:"url"`
         
    }
    
    // BusinessDetails used by CompanyProfile
    type BusinessDetails struct {

        
            Website Website  `json:"website"`
         
    }
    
    // CompanyTaxesSerializer used by CompanyProfile
    type CompanyTaxesSerializer struct {

        
            Rate float64  `json:"rate"`
            Enable bool  `json:"enable"`
            EffectiveDate string  `json:"effective_date"`
         
    }
    
    // BusinessCountryInfo used by CompanyProfile
    type BusinessCountryInfo struct {

        
            CountryCode string  `json:"country_code"`
            Country string  `json:"country"`
         
    }
    
    // ContactDetails used by CompanyProfile
    type ContactDetails struct {

        
            Phone []SellerPhoneNumber  `json:"phone"`
            Emails []string  `json:"emails"`
         
    }
    
    // GetCompanyProfileSerializerResponse used by CompanyProfile
    type GetCompanyProfileSerializerResponse struct {

        
            BusinessDetails BusinessDetails  `json:"business_details"`
            Taxes []CompanyTaxesSerializer  `json:"taxes"`
            ModifiedOn string  `json:"modified_on"`
            Stage string  `json:"stage"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            FranchiseEnabled bool  `json:"franchise_enabled"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            CreatedOn string  `json:"created_on"`
            UID float64  `json:"uid"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            BusinessCountryInfo BusinessCountryInfo  `json:"business_country_info"`
            NotificationEmails []string  `json:"notification_emails"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            VerifiedOn string  `json:"verified_on"`
            Mode string  `json:"mode"`
            Warnings map[string]interface{}  `json:"warnings"`
            ContactDetails ContactDetails  `json:"contact_details"`
            CompanyType string  `json:"company_type"`
            BusinessType string  `json:"business_type"`
            BusinessInfo string  `json:"business_info"`
            Documents []Document  `json:"documents"`
            Name string  `json:"name"`
            CreatedBy UserSerializer  `json:"created_by"`
         
    }
    
    // CompanyTaxesSerializer1 used by CompanyProfile
    type CompanyTaxesSerializer1 struct {

        
            Rate float64  `json:"rate"`
            Enable bool  `json:"enable"`
            EffectiveDate string  `json:"effective_date"`
         
    }
    
    // CreateUpdateAddressSerializer used by CompanyProfile
    type CreateUpdateAddressSerializer struct {

        
            State string  `json:"state"`
            Address1 string  `json:"address1"`
            AddressType string  `json:"address_type"`
            Latitude float64  `json:"latitude"`
            Pincode float64  `json:"pincode"`
            City string  `json:"city"`
            Address2 string  `json:"address2"`
            Country string  `json:"country"`
            Longitude float64  `json:"longitude"`
            CountryCode string  `json:"country_code"`
            Landmark string  `json:"landmark"`
         
    }
    
    // UpdateCompany used by CompanyProfile
    type UpdateCompany struct {

        
            RejectReason string  `json:"reject_reason"`
            NotificationEmails []string  `json:"notification_emails"`
            BusinessDetails BusinessDetails  `json:"business_details"`
            ContactDetails ContactDetails  `json:"contact_details"`
            CompanyType string  `json:"company_type"`
            BusinessType string  `json:"business_type"`
            Taxes []CompanyTaxesSerializer1  `json:"taxes"`
            BusinessInfo string  `json:"business_info"`
            Documents []Document  `json:"documents"`
            Name string  `json:"name"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Addresses []CreateUpdateAddressSerializer  `json:"addresses"`
            Warnings map[string]interface{}  `json:"warnings"`
            FranchiseEnabled bool  `json:"franchise_enabled"`
         
    }
    
    // ProfileSuccessResponse used by CompanyProfile
    type ProfileSuccessResponse struct {

        
            Success bool  `json:"success"`
            UID float64  `json:"uid"`
         
    }
    
    // DocumentsObj used by CompanyProfile
    type DocumentsObj struct {

        
            Pending float64  `json:"pending"`
            Verified float64  `json:"verified"`
         
    }
    
    // MetricsSerializer used by CompanyProfile
    type MetricsSerializer struct {

        
            Store DocumentsObj  `json:"store"`
            CompanyDocuments DocumentsObj  `json:"company_documents"`
            Product DocumentsObj  `json:"product"`
            UID float64  `json:"uid"`
            Stage string  `json:"stage"`
            Brand DocumentsObj  `json:"brand"`
            StoreDocuments DocumentsObj  `json:"store_documents"`
         
    }
    
    // BrandBannerSerializer used by CompanyProfile
    type BrandBannerSerializer struct {

        
            Landscape string  `json:"landscape"`
            Portrait string  `json:"portrait"`
         
    }
    
    // GetBrandResponseSerializer used by CompanyProfile
    type GetBrandResponseSerializer struct {

        
            Banner BrandBannerSerializer  `json:"banner"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            ModifiedOn string  `json:"modified_on"`
            Stage string  `json:"stage"`
            Synonyms []string  `json:"synonyms"`
            Description string  `json:"description"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            CreatedOn string  `json:"created_on"`
            UID float64  `json:"uid"`
            SlugKey string  `json:"slug_key"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            VerifiedOn string  `json:"verified_on"`
            Mode string  `json:"mode"`
            Warnings map[string]interface{}  `json:"warnings"`
            Logo string  `json:"logo"`
            Name string  `json:"name"`
            CreatedBy UserSerializer  `json:"created_by"`
            RejectReason string  `json:"reject_reason"`
         
    }
    
    // CreateUpdateBrandRequestSerializer used by CompanyProfile
    type CreateUpdateBrandRequestSerializer struct {

        
            Logo string  `json:"logo"`
            Banner BrandBannerSerializer  `json:"banner"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            CompanyID float64  `json:"company_id"`
            UID float64  `json:"uid"`
            BrandTier string  `json:"brand_tier"`
            Name string  `json:"name"`
            Synonyms []string  `json:"synonyms"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Description string  `json:"description"`
         
    }
    
    // CompanySocialAccounts used by CompanyProfile
    type CompanySocialAccounts struct {

        
            Name string  `json:"name"`
            URL string  `json:"url"`
         
    }
    
    // CompanyDetails used by CompanyProfile
    type CompanyDetails struct {

        
            WebsiteURL string  `json:"website_url"`
            Socials []CompanySocialAccounts  `json:"socials"`
         
    }
    
    // CompanySerializer used by CompanyProfile
    type CompanySerializer struct {

        
            RejectReason string  `json:"reject_reason"`
            NotificationEmails []string  `json:"notification_emails"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            Details CompanyDetails  `json:"details"`
            CreatedOn string  `json:"created_on"`
            CompanyType string  `json:"company_type"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            BusinessType string  `json:"business_type"`
            VerifiedOn string  `json:"verified_on"`
            ModifiedOn string  `json:"modified_on"`
            UID float64  `json:"uid"`
            MarketChannels []string  `json:"market_channels"`
            Stage string  `json:"stage"`
            Name string  `json:"name"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            CreatedBy UserSerializer  `json:"created_by"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            BusinessCountryInfo BusinessCountryInfo  `json:"business_country_info"`
         
    }
    
    // CompanyBrandSerializer used by CompanyProfile
    type CompanyBrandSerializer struct {

        
            VerifiedBy UserSerializer  `json:"verified_by"`
            CreatedOn string  `json:"created_on"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            VerifiedOn string  `json:"verified_on"`
            ModifiedOn string  `json:"modified_on"`
            UID float64  `json:"uid"`
            Stage string  `json:"stage"`
            Brand GetBrandResponseSerializer  `json:"brand"`
            CreatedBy UserSerializer  `json:"created_by"`
            RejectReason string  `json:"reject_reason"`
            Company CompanySerializer  `json:"company"`
            Warnings map[string]interface{}  `json:"warnings"`
         
    }
    
    // CompanyBrandListSerializer used by CompanyProfile
    type CompanyBrandListSerializer struct {

        
            Items []CompanyBrandSerializer  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // CompanyBrandPostRequestSerializer used by CompanyProfile
    type CompanyBrandPostRequestSerializer struct {

        
            Brands []float64  `json:"brands"`
            Company float64  `json:"company"`
            UID float64  `json:"uid"`
         
    }
    
    // HolidayDateSerializer used by CompanyProfile
    type HolidayDateSerializer struct {

        
            StartDate string  `json:"start_date"`
            EndDate string  `json:"end_date"`
         
    }
    
    // HolidaySchemaSerializer used by CompanyProfile
    type HolidaySchemaSerializer struct {

        
            Date HolidayDateSerializer  `json:"date"`
            HolidayType string  `json:"holiday_type"`
            Title string  `json:"title"`
         
    }
    
    // AddressSerializer used by CompanyProfile
    type AddressSerializer struct {

        
            Address1 string  `json:"address1"`
            State string  `json:"state"`
            AddressType string  `json:"address_type"`
            Latitude float64  `json:"latitude"`
            Pincode float64  `json:"pincode"`
            City string  `json:"city"`
            Address2 string  `json:"address2"`
            Country string  `json:"country"`
            Longitude float64  `json:"longitude"`
            CountryCode string  `json:"country_code"`
            Landmark string  `json:"landmark"`
         
    }
    
    // LocationSerializer used by CompanyProfile
    type LocationSerializer struct {

        
            NotificationEmails []string  `json:"notification_emails"`
            DisplayName string  `json:"display_name"`
            Code string  `json:"code"`
            Holiday []HolidaySchemaSerializer  `json:"holiday"`
            UID float64  `json:"uid"`
            GstCredentials InvoiceDetailsSerializer  `json:"gst_credentials"`
            StoreType string  `json:"store_type"`
            Stage string  `json:"stage"`
            Documents []Document  `json:"documents"`
            Name string  `json:"name"`
            Address AddressSerializer  `json:"address"`
            ProductReturnConfig ProductReturnConfigSerializer  `json:"product_return_config"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            Company float64  `json:"company"`
            Timing []LocationDayWiseSerializer  `json:"timing"`
            Manager LocationManagerSerializer  `json:"manager"`
            Warnings map[string]interface{}  `json:"warnings"`
         
    }
    
    // BulkLocationSerializer used by CompanyProfile
    type BulkLocationSerializer struct {

        
            Data []LocationSerializer  `json:"data"`
         
    }
    
    // _ArticleQuery used by CompanyProfile
    type _ArticleQuery struct {

        
            ItemID float64  `json:"item_id"`
            Size string  `json:"size"`
            IgnoredStores []float64  `json:"ignored_stores"`
         
    }
    
    // _ArticleAssignment used by CompanyProfile
    type _ArticleAssignment struct {

        
            Level string  `json:"level"`
            Strategy string  `json:"strategy"`
         
    }
    
    // _AssignStoreArticle used by CompanyProfile
    type _AssignStoreArticle struct {

        
            Meta map[string]interface{}  `json:"meta"`
            GroupID string  `json:"group_id"`
            Quantity float64  `json:"quantity"`
            Query _ArticleQuery  `json:"query"`
            ArticleAssignment _ArticleAssignment  `json:"article_assignment"`
         
    }
    
    // AssignStoreRequestValidator used by CompanyProfile
    type AssignStoreRequestValidator struct {

        
            AppID string  `json:"app_id"`
            ChannelIdentifier string  `json:"channel_identifier"`
            Articles []_AssignStoreArticle  `json:"articles"`
            CompanyID float64  `json:"company_id"`
            Pincode string  `json:"pincode"`
            ChannelType string  `json:"channel_type"`
            StoreIds []float64  `json:"store_ids"`
         
    }
    
    // AssignStoreResponseSerializer used by CompanyProfile
    type AssignStoreResponseSerializer struct {

        
            ItemID float64  `json:"item_id"`
            ID string  `json:"_id"`
            Meta map[string]interface{}  `json:"meta"`
            PriceEffective float64  `json:"price_effective"`
            PriceMarked float64  `json:"price_marked"`
            Quantity float64  `json:"quantity"`
            StoreID float64  `json:"store_id"`
            CompanyID float64  `json:"company_id"`
            UID string  `json:"uid"`
            SCity string  `json:"s_city"`
            StorePincode string  `json:"store_pincode"`
            ArticleAssignment _ArticleAssignment  `json:"article_assignment"`
            Size string  `json:"size"`
            Status bool  `json:"status"`
            Index float64  `json:"index"`
         
    }
    

    
    // FailedResponse used by FileStorage
    type FailedResponse struct {

        
            Message string  `json:"message"`
         
    }
    
    // CDN used by FileStorage
    type CDN struct {

        
            URL string  `json:"url"`
         
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
    

    
    // CouponSchedule used by Cart
    type CouponSchedule struct {

        
            End string  `json:"end"`
            Cron string  `json:"cron"`
            Start string  `json:"start"`
            NextSchedule []map[string]interface{}  `json:"next_schedule"`
            Duration float64  `json:"duration"`
         
    }
    
    // Ownership used by Cart
    type Ownership struct {

        
            PayableCategory string  `json:"payable_category"`
            PayableBy string  `json:"payable_by"`
         
    }
    
    // Rule used by Cart
    type Rule struct {

        
            Key float64  `json:"key"`
            Value float64  `json:"value"`
            Max float64  `json:"max"`
            Min float64  `json:"min"`
            DiscountQty float64  `json:"discount_qty"`
         
    }
    
    // DisplayMetaDict used by Cart
    type DisplayMetaDict struct {

        
            Subtitle string  `json:"subtitle"`
            Title string  `json:"title"`
         
    }
    
    // DisplayMeta used by Cart
    type DisplayMeta struct {

        
            Title string  `json:"title"`
            Auto DisplayMetaDict  `json:"auto"`
            Apply DisplayMetaDict  `json:"apply"`
            Description string  `json:"description"`
            Remove DisplayMetaDict  `json:"remove"`
            Subtitle string  `json:"subtitle"`
         
    }
    
    // Validation used by Cart
    type Validation struct {

        
            UserRegisteredAfter string  `json:"user_registered_after"`
            Anonymous bool  `json:"anonymous"`
            AppID []string  `json:"app_id"`
         
    }
    
    // State used by Cart
    type State struct {

        
            IsArchived bool  `json:"is_archived"`
            IsPublic bool  `json:"is_public"`
            IsDisplay bool  `json:"is_display"`
         
    }
    
    // RuleDefinition used by Cart
    type RuleDefinition struct {

        
            AutoApply bool  `json:"auto_apply"`
            Type string  `json:"type"`
            CalculateOn string  `json:"calculate_on"`
            Scope []string  `json:"scope"`
            ApplicableOn string  `json:"applicable_on"`
            IsExact bool  `json:"is_exact"`
            CurrencyCode string  `json:"currency_code"`
            ValueType string  `json:"value_type"`
         
    }
    
    // CouponAction used by Cart
    type CouponAction struct {

        
            ActionDate string  `json:"action_date"`
            TxnMode string  `json:"txn_mode"`
         
    }
    
    // Validity used by Cart
    type Validity struct {

        
            Priority float64  `json:"priority"`
         
    }
    
    // CouponDateMeta used by Cart
    type CouponDateMeta struct {

        
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // BulkBundleRestriction used by Cart
    type BulkBundleRestriction struct {

        
            MultiStoreAllowed bool  `json:"multi_store_allowed"`
         
    }
    
    // PriceRange used by Cart
    type PriceRange struct {

        
            Max float64  `json:"max"`
            Min float64  `json:"min"`
         
    }
    
    // PaymentAllowValue used by Cart
    type PaymentAllowValue struct {

        
            Max float64  `json:"max"`
         
    }
    
    // PaymentModes used by Cart
    type PaymentModes struct {

        
            Uses PaymentAllowValue  `json:"uses"`
            Types []string  `json:"types"`
            Networks []string  `json:"networks"`
            Codes []string  `json:"codes"`
         
    }
    
    // PostOrder used by Cart
    type PostOrder struct {

        
            ReturnAllowed bool  `json:"return_allowed"`
            CancellationAllowed bool  `json:"cancellation_allowed"`
         
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
    
    // Restrictions used by Cart
    type Restrictions struct {

        
            BulkBundle BulkBundleRestriction  `json:"bulk_bundle"`
            Platforms []string  `json:"platforms"`
            PriceRange PriceRange  `json:"price_range"`
            CouponAllowed bool  `json:"coupon_allowed"`
            Payments map[string]PaymentModes  `json:"payments"`
            UserType string  `json:"user_type"`
            OrderingStores []float64  `json:"ordering_stores"`
            UserGroups []float64  `json:"user_groups"`
            PostOrder PostOrder  `json:"post_order"`
            Uses UsesRestriction  `json:"uses"`
         
    }
    
    // CouponAuthor used by Cart
    type CouponAuthor struct {

        
            ModifiedBy string  `json:"modified_by"`
            CreatedBy string  `json:"created_by"`
         
    }
    
    // CouponAdd used by Cart
    type CouponAdd struct {

        
            Schedule CouponSchedule  `json:"_schedule"`
            Ownership Ownership  `json:"ownership"`
            Rule []Rule  `json:"rule"`
            DisplayMeta DisplayMeta  `json:"display_meta"`
            Validation Validation  `json:"validation"`
            State State  `json:"state"`
            Identifiers Identifier  `json:"identifiers"`
            RuleDefinition RuleDefinition  `json:"rule_definition"`
            Action CouponAction  `json:"action"`
            Tags []string  `json:"tags"`
            Validity Validity  `json:"validity"`
            DateMeta CouponDateMeta  `json:"date_meta"`
            TypeSlug string  `json:"type_slug"`
            Restrictions Restrictions  `json:"restrictions"`
            Code string  `json:"code"`
            Author CouponAuthor  `json:"author"`
         
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

        
            Schedule CouponSchedule  `json:"_schedule"`
            Ownership Ownership  `json:"ownership"`
            Rule []Rule  `json:"rule"`
            DisplayMeta DisplayMeta  `json:"display_meta"`
            Validation Validation  `json:"validation"`
            State State  `json:"state"`
            Identifiers Identifier  `json:"identifiers"`
            RuleDefinition RuleDefinition  `json:"rule_definition"`
            Action CouponAction  `json:"action"`
            Tags []string  `json:"tags"`
            Validity Validity  `json:"validity"`
            DateMeta CouponDateMeta  `json:"date_meta"`
            TypeSlug string  `json:"type_slug"`
            Restrictions Restrictions  `json:"restrictions"`
            Code string  `json:"code"`
            Author CouponAuthor  `json:"author"`
         
    }
    
    // CouponPartialUpdate used by Cart
    type CouponPartialUpdate struct {

        
            Schedule CouponSchedule  `json:"schedule"`
            Archive bool  `json:"archive"`
         
    }
    
    // DisplayMeta1 used by Cart
    type DisplayMeta1 struct {

        
            OfferText string  `json:"offer_text"`
            Name string  `json:"name"`
            OfferLabel string  `json:"offer_label"`
            Description string  `json:"description"`
         
    }
    
    // DiscountOffer used by Cart
    type DiscountOffer struct {

        
            DiscountAmount float64  `json:"discount_amount"`
            MinOfferQuantity float64  `json:"min_offer_quantity"`
            MaxOfferQuantity float64  `json:"max_offer_quantity"`
            DiscountPercentage float64  `json:"discount_percentage"`
            PartialCanRet bool  `json:"partial_can_ret"`
            ApportionDiscount bool  `json:"apportion_discount"`
            Code string  `json:"code"`
            MaxDiscountAmount float64  `json:"max_discount_amount"`
            DiscountPrice float64  `json:"discount_price"`
            MaxUsagePerTransaction float64  `json:"max_usage_per_transaction"`
         
    }
    
    // CompareObject used by Cart
    type CompareObject struct {

        
            LessThan float64  `json:"less_than"`
            GreaterThanEquals float64  `json:"greater_than_equals"`
            GreaterThan float64  `json:"greater_than"`
            LessThanEquals float64  `json:"less_than_equals"`
            Equals float64  `json:"equals"`
         
    }
    
    // ItemCriteria used by Cart
    type ItemCriteria struct {

        
            ItemExcludeCompany []float64  `json:"item_exclude_company"`
            CartUniqueItemQuantity CompareObject  `json:"cart_unique_item_quantity"`
            ItemExcludeID []float64  `json:"item_exclude_id"`
            CartQuantity CompareObject  `json:"cart_quantity"`
            AvailableZones []string  `json:"available_zones"`
            ItemCompany []float64  `json:"item_company"`
            ItemExcludeSku []string  `json:"item_exclude_sku"`
            ItemSize []string  `json:"item_size"`
            ItemExcludeCategory []float64  `json:"item_exclude_category"`
            AllItems bool  `json:"all_items"`
            ItemCategory []float64  `json:"item_category"`
            ItemExcludeStore []float64  `json:"item_exclude_store"`
            ItemExcludeBrand []float64  `json:"item_exclude_brand"`
            BuyRules []string  `json:"buy_rules"`
            ItemBrand []float64  `json:"item_brand"`
            CartTotal CompareObject  `json:"cart_total"`
            ItemSku []string  `json:"item_sku"`
            ItemStore []float64  `json:"item_store"`
            ProductTags []string  `json:"product_tags"`
            ItemID []float64  `json:"item_id"`
            CartUniqueItemAmount CompareObject  `json:"cart_unique_item_amount"`
         
    }
    
    // DiscountRule used by Cart
    type DiscountRule struct {

        
            Offer DiscountOffer  `json:"offer"`
            DiscountType string  `json:"discount_type"`
            BuyCondition string  `json:"buy_condition"`
            ItemCriteria ItemCriteria  `json:"item_criteria"`
         
    }
    
    // PaymentAllowValue1 used by Cart
    type PaymentAllowValue1 struct {

        
            Max float64  `json:"max"`
         
    }
    
    // PromotionPaymentModes used by Cart
    type PromotionPaymentModes struct {

        
            Type string  `json:"type"`
            Uses PaymentAllowValue1  `json:"uses"`
            Codes []string  `json:"codes"`
         
    }
    
    // PostOrder1 used by Cart
    type PostOrder1 struct {

        
            ReturnAllowed bool  `json:"return_allowed"`
            CancellationAllowed bool  `json:"cancellation_allowed"`
         
    }
    
    // UserRegistered used by Cart
    type UserRegistered struct {

        
            Start string  `json:"start"`
            End string  `json:"end"`
         
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
    
    // Restrictions1 used by Cart
    type Restrictions1 struct {

        
            UserID []string  `json:"user_id"`
            Platforms []string  `json:"platforms"`
            AnonymousUsers bool  `json:"anonymous_users"`
            Payments []PromotionPaymentModes  `json:"payments"`
            UserGroups []float64  `json:"user_groups"`
            PostOrder PostOrder1  `json:"post_order"`
            UserRegistered UserRegistered  `json:"user_registered"`
            OrderQuantity float64  `json:"order_quantity"`
            Uses UsesRestriction1  `json:"uses"`
         
    }
    
    // Ownership1 used by Cart
    type Ownership1 struct {

        
            PayableCategory string  `json:"payable_category"`
            PayableBy string  `json:"payable_by"`
         
    }
    
    // Visibility used by Cart
    type Visibility struct {

        
            Pdp bool  `json:"pdp"`
            CouponList bool  `json:"coupon_list"`
         
    }
    
    // PromotionDateMeta used by Cart
    type PromotionDateMeta struct {

        
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // PromotionAuthor used by Cart
    type PromotionAuthor struct {

        
            ModifiedBy string  `json:"modified_by"`
            CreatedBy string  `json:"created_by"`
         
    }
    
    // PromotionSchedule used by Cart
    type PromotionSchedule struct {

        
            Published bool  `json:"published"`
            End string  `json:"end"`
            Cron string  `json:"cron"`
            Start string  `json:"start"`
            NextSchedule []map[string]interface{}  `json:"next_schedule"`
            Duration float64  `json:"duration"`
         
    }
    
    // PromotionAction used by Cart
    type PromotionAction struct {

        
            ActionType string  `json:"action_type"`
            ActionDate string  `json:"action_date"`
         
    }
    
    // PromotionListItem used by Cart
    type PromotionListItem struct {

        
            DisplayMeta DisplayMeta1  `json:"display_meta"`
            DiscountRules []DiscountRule  `json:"discount_rules"`
            Restrictions Restrictions1  `json:"restrictions"`
            Ownership Ownership1  `json:"ownership"`
            Visiblility Visibility  `json:"visiblility"`
            ApplyAllDiscount bool  `json:"apply_all_discount"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            CalculateOn string  `json:"calculate_on"`
            Stackable bool  `json:"stackable"`
            ApplyExclusive string  `json:"apply_exclusive"`
            PromoGroup string  `json:"promo_group"`
            ApplyPriority float64  `json:"apply_priority"`
            Currency string  `json:"currency"`
            DateMeta PromotionDateMeta  `json:"date_meta"`
            BuyRules map[string]ItemCriteria  `json:"buy_rules"`
            Code string  `json:"code"`
            Author PromotionAuthor  `json:"author"`
            Schedule PromotionSchedule  `json:"_schedule"`
            PromotionType string  `json:"promotion_type"`
            ApplicationID string  `json:"application_id"`
            PostOrderAction PromotionAction  `json:"post_order_action"`
            Mode string  `json:"mode"`
         
    }
    
    // PromotionsResponse used by Cart
    type PromotionsResponse struct {

        
            Items PromotionListItem  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // PromotionAdd used by Cart
    type PromotionAdd struct {

        
            DisplayMeta DisplayMeta1  `json:"display_meta"`
            DiscountRules []DiscountRule  `json:"discount_rules"`
            Restrictions Restrictions1  `json:"restrictions"`
            Ownership Ownership1  `json:"ownership"`
            Visiblility Visibility  `json:"visiblility"`
            ApplyAllDiscount bool  `json:"apply_all_discount"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            CalculateOn string  `json:"calculate_on"`
            Stackable bool  `json:"stackable"`
            ApplyExclusive string  `json:"apply_exclusive"`
            PromoGroup string  `json:"promo_group"`
            ApplyPriority float64  `json:"apply_priority"`
            Currency string  `json:"currency"`
            DateMeta PromotionDateMeta  `json:"date_meta"`
            BuyRules map[string]ItemCriteria  `json:"buy_rules"`
            Code string  `json:"code"`
            Author PromotionAuthor  `json:"author"`
            Schedule PromotionSchedule  `json:"_schedule"`
            PromotionType string  `json:"promotion_type"`
            ApplicationID string  `json:"application_id"`
            PostOrderAction PromotionAction  `json:"post_order_action"`
            Mode string  `json:"mode"`
         
    }
    
    // PromotionUpdate used by Cart
    type PromotionUpdate struct {

        
            DisplayMeta DisplayMeta1  `json:"display_meta"`
            DiscountRules []DiscountRule  `json:"discount_rules"`
            Restrictions Restrictions1  `json:"restrictions"`
            Ownership Ownership1  `json:"ownership"`
            Visiblility Visibility  `json:"visiblility"`
            ApplyAllDiscount bool  `json:"apply_all_discount"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            CalculateOn string  `json:"calculate_on"`
            Stackable bool  `json:"stackable"`
            ApplyExclusive string  `json:"apply_exclusive"`
            PromoGroup string  `json:"promo_group"`
            ApplyPriority float64  `json:"apply_priority"`
            Currency string  `json:"currency"`
            DateMeta PromotionDateMeta  `json:"date_meta"`
            BuyRules map[string]ItemCriteria  `json:"buy_rules"`
            Code string  `json:"code"`
            Author PromotionAuthor  `json:"author"`
            Schedule PromotionSchedule  `json:"_schedule"`
            PromotionType string  `json:"promotion_type"`
            ApplicationID string  `json:"application_id"`
            PostOrderAction PromotionAction  `json:"post_order_action"`
            Mode string  `json:"mode"`
         
    }
    
    // PromotionPartialUpdate used by Cart
    type PromotionPartialUpdate struct {

        
            Schedule PromotionSchedule  `json:"schedule"`
            Archive bool  `json:"archive"`
         
    }
    
    // ActivePromosResponse used by Cart
    type ActivePromosResponse struct {

        
            Type string  `json:"type"`
            IsHidden bool  `json:"is_hidden"`
            Title string  `json:"title"`
            ModifiedOn string  `json:"modified_on"`
            Description string  `json:"description"`
            CreatedOn string  `json:"created_on"`
            EntitySlug string  `json:"entity_slug"`
            EntityType string  `json:"entity_type"`
            Subtitle string  `json:"subtitle"`
            Example string  `json:"example"`
         
    }
    
    // CartItem used by Cart
    type CartItem struct {

        
            ProductID string  `json:"product_id"`
            Quantity float64  `json:"quantity"`
            Size string  `json:"size"`
         
    }
    
    // OpenapiCartDetailsRequest used by Cart
    type OpenapiCartDetailsRequest struct {

        
            CartItems CartItem  `json:"cart_items"`
         
    }
    
    // DisplayBreakup used by Cart
    type DisplayBreakup struct {

        
            Key string  `json:"key"`
            Value float64  `json:"value"`
            CurrencySymbol string  `json:"currency_symbol"`
            CurrencyCode string  `json:"currency_code"`
            Display string  `json:"display"`
            Message []string  `json:"message"`
         
    }
    
    // LoyaltyPoints used by Cart
    type LoyaltyPoints struct {

        
            Applicable float64  `json:"applicable"`
            IsApplied bool  `json:"is_applied"`
            Description string  `json:"description"`
            Total float64  `json:"total"`
         
    }
    
    // CouponBreakup used by Cart
    type CouponBreakup struct {

        
            IsApplied bool  `json:"is_applied"`
            Type string  `json:"type"`
            CouponType string  `json:"coupon_type"`
            CouponValue float64  `json:"coupon_value"`
            Title string  `json:"title"`
            Value float64  `json:"value"`
            MaxDiscountValue float64  `json:"max_discount_value"`
            Description string  `json:"description"`
            UID string  `json:"uid"`
            SubTitle string  `json:"sub_title"`
            Code string  `json:"code"`
            MinimumCartValue float64  `json:"minimum_cart_value"`
            Message string  `json:"message"`
         
    }
    
    // RawBreakup used by Cart
    type RawBreakup struct {

        
            Vog float64  `json:"vog"`
            Subtotal float64  `json:"subtotal"`
            DeliveryCharge float64  `json:"delivery_charge"`
            Coupon float64  `json:"coupon"`
            FyndCash float64  `json:"fynd_cash"`
            CodCharge float64  `json:"cod_charge"`
            MrpTotal float64  `json:"mrp_total"`
            GstCharges float64  `json:"gst_charges"`
            Discount float64  `json:"discount"`
            Total float64  `json:"total"`
            ConvenienceFee float64  `json:"convenience_fee"`
            YouSaved float64  `json:"you_saved"`
         
    }
    
    // CartBreakup used by Cart
    type CartBreakup struct {

        
            Display []DisplayBreakup  `json:"display"`
            LoyaltyPoints LoyaltyPoints  `json:"loyalty_points"`
            Coupon CouponBreakup  `json:"coupon"`
            Raw RawBreakup  `json:"raw"`
         
    }
    
    // Ownership2 used by Cart
    type Ownership2 struct {

        
            PayableCategory string  `json:"payable_category"`
            PayableBy string  `json:"payable_by"`
         
    }
    
    // DiscountRulesApp used by Cart
    type DiscountRulesApp struct {

        
            MatchedBuyRules []string  `json:"matched_buy_rules"`
            Offer map[string]interface{}  `json:"offer"`
            RawOffer map[string]interface{}  `json:"raw_offer"`
            ItemCriteria map[string]interface{}  `json:"item_criteria"`
         
    }
    
    // FreeGiftItem used by Cart
    type FreeGiftItem struct {

        
            ItemName string  `json:"item_name"`
            ItemBrandName string  `json:"item_brand_name"`
            ItemPriceDetails map[string]interface{}  `json:"item_price_details"`
            ItemID float64  `json:"item_id"`
            ItemImagesURL []string  `json:"item_images_url"`
            ItemSlug string  `json:"item_slug"`
         
    }
    
    // AppliedFreeArticles used by Cart
    type AppliedFreeArticles struct {

        
            Quantity float64  `json:"quantity"`
            ParentItemIdentifier string  `json:"parent_item_identifier"`
            FreeGiftItemDetails FreeGiftItem  `json:"free_gift_item_details"`
            ArticleID string  `json:"article_id"`
         
    }
    
    // AppliedPromotion used by Cart
    type AppliedPromotion struct {

        
            Ownership Ownership2  `json:"ownership"`
            Amount float64  `json:"amount"`
            PromotionType string  `json:"promotion_type"`
            PromotionGroup string  `json:"promotion_group"`
            PromotionName string  `json:"promotion_name"`
            DiscountRules []DiscountRulesApp  `json:"discount_rules"`
            OfferText string  `json:"offer_text"`
            MrpPromotion bool  `json:"mrp_promotion"`
            PromoID string  `json:"promo_id"`
            ArticleQuantity float64  `json:"article_quantity"`
            BuyRules []BuyRules  `json:"buy_rules"`
            AppliedFreeArticles []AppliedFreeArticles  `json:"applied_free_articles"`
         
    }
    
    // ProductPrice used by Cart
    type ProductPrice struct {

        
            AddOn float64  `json:"add_on"`
            Effective float64  `json:"effective"`
            Selling float64  `json:"selling"`
            Marked float64  `json:"marked"`
            CurrencySymbol string  `json:"currency_symbol"`
            CurrencyCode string  `json:"currency_code"`
         
    }
    
    // ProductPriceInfo used by Cart
    type ProductPriceInfo struct {

        
            Converted ProductPrice  `json:"converted"`
            Base ProductPrice  `json:"base"`
         
    }
    
    // ProductImage used by Cart
    type ProductImage struct {

        
            SecureURL string  `json:"secure_url"`
            AspectRatio string  `json:"aspect_ratio"`
            URL string  `json:"url"`
         
    }
    
    // CategoryInfo used by Cart
    type CategoryInfo struct {

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // ActionQuery used by Cart
    type ActionQuery struct {

        
            ProductSlug []string  `json:"product_slug"`
         
    }
    
    // ProductAction used by Cart
    type ProductAction struct {

        
            Type string  `json:"type"`
            Query ActionQuery  `json:"query"`
            URL string  `json:"url"`
         
    }
    
    // BaseInfo used by Cart
    type BaseInfo struct {

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // CartProduct used by Cart
    type CartProduct struct {

        
            Images []ProductImage  `json:"images"`
            Type string  `json:"type"`
            NetQuantity NetQuantity  `json:"net_quantity"`
            Categories []CategoryInfo  `json:"categories"`
            Action ProductAction  `json:"action"`
            Brand BaseInfo  `json:"brand"`
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // CartProductIdentifer used by Cart
    type CartProductIdentifer struct {

        
            Identifier string  `json:"identifier"`
         
    }
    
    // ProductAvailability used by Cart
    type ProductAvailability struct {

        
            IsValid bool  `json:"is_valid"`
            Sizes []string  `json:"sizes"`
            OtherStoreQuantity float64  `json:"other_store_quantity"`
            OutOfStock bool  `json:"out_of_stock"`
            Deliverable bool  `json:"deliverable"`
         
    }
    
    // BasePrice used by Cart
    type BasePrice struct {

        
            CurrencySymbol string  `json:"currency_symbol"`
            Effective float64  `json:"effective"`
            Marked float64  `json:"marked"`
            CurrencyCode string  `json:"currency_code"`
         
    }
    
    // ArticlePriceInfo used by Cart
    type ArticlePriceInfo struct {

        
            Converted BasePrice  `json:"converted"`
            Base BasePrice  `json:"base"`
         
    }
    
    // ProductArticle used by Cart
    type ProductArticle struct {

        
            Type string  `json:"type"`
            Price ArticlePriceInfo  `json:"price"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            Quantity float64  `json:"quantity"`
            ProductGroupTags []string  `json:"product_group_tags"`
            Seller BaseInfo  `json:"seller"`
            Store BaseInfo  `json:"store"`
            Size string  `json:"size"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            UID string  `json:"uid"`
         
    }
    
    // PromoMeta used by Cart
    type PromoMeta struct {

        
            Message string  `json:"message"`
         
    }
    
    // CartProductInfo used by Cart
    type CartProductInfo struct {

        
            PromotionsApplied []AppliedPromotion  `json:"promotions_applied"`
            Key string  `json:"key"`
            Price ProductPriceInfo  `json:"price"`
            Product CartProduct  `json:"product"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            CouponMessage string  `json:"coupon_message"`
            Quantity float64  `json:"quantity"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
            IsSet bool  `json:"is_set"`
            Discount string  `json:"discount"`
            BulkOffer map[string]interface{}  `json:"bulk_offer"`
            Availability ProductAvailability  `json:"availability"`
            Article ProductArticle  `json:"article"`
            PricePerUnit ProductPriceInfo  `json:"price_per_unit"`
            PromoMeta PromoMeta  `json:"promo_meta"`
            Message string  `json:"message"`
         
    }
    
    // OpenapiCartDetailsResponse used by Cart
    type OpenapiCartDetailsResponse struct {

        
            BreakupValues CartBreakup  `json:"breakup_values"`
            Items []CartProductInfo  `json:"items"`
            IsValid bool  `json:"is_valid"`
            Message string  `json:"message"`
         
    }
    
    // OpenApiErrorResponse used by Cart
    type OpenApiErrorResponse struct {

        
            Errors map[string]interface{}  `json:"errors"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // ShippingAddress used by Cart
    type ShippingAddress struct {

        
            AddressType string  `json:"address_type"`
            State string  `json:"state"`
            Phone float64  `json:"phone"`
            Address string  `json:"address"`
            Landmark string  `json:"landmark"`
            AreaCode string  `json:"area_code"`
            City string  `json:"city"`
            Area string  `json:"area"`
            CountryCode string  `json:"country_code"`
            AreaCodeSlug string  `json:"area_code_slug"`
            Pincode float64  `json:"pincode"`
            Email string  `json:"email"`
            Country string  `json:"country"`
            Name string  `json:"name"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // OpenApiCartServiceabilityRequest used by Cart
    type OpenApiCartServiceabilityRequest struct {

        
            ShippingAddress ShippingAddress  `json:"shipping_address"`
            CartItems CartItem  `json:"cart_items"`
         
    }
    
    // PromiseTimestamp used by Cart
    type PromiseTimestamp struct {

        
            Max float64  `json:"max"`
            Min float64  `json:"min"`
         
    }
    
    // PromiseFormatted used by Cart
    type PromiseFormatted struct {

        
            Max string  `json:"max"`
            Min string  `json:"min"`
         
    }
    
    // ShipmentPromise used by Cart
    type ShipmentPromise struct {

        
            Timestamp PromiseTimestamp  `json:"timestamp"`
            Formatted PromiseFormatted  `json:"formatted"`
         
    }
    
    // OpenApiCartServiceabilityResponse used by Cart
    type OpenApiCartServiceabilityResponse struct {

        
            IsValid bool  `json:"is_valid"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            Items []CartProductInfo  `json:"items"`
            Message string  `json:"message"`
         
    }
    
    // OpenApiFiles used by Cart
    type OpenApiFiles struct {

        
            Key string  `json:"key"`
            Values []string  `json:"values"`
         
    }
    
    // CartItemMeta used by Cart
    type CartItemMeta struct {

        
            GroupID string  `json:"group_id"`
            PrimaryItem bool  `json:"primary_item"`
         
    }
    
    // OpenApiOrderItem used by Cart
    type OpenApiOrderItem struct {

        
            PaymentMethods []MultiTenderPaymentMethod  `json:"payment_methods"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            PriceMarked float64  `json:"price_marked"`
            CashbackApplied float64  `json:"cashback_applied"`
            Quantity float64  `json:"quantity"`
            CodCharges float64  `json:"cod_charges"`
            DeliveryCharges float64  `json:"delivery_charges"`
            Files []OpenApiFiles  `json:"files"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            Size string  `json:"size"`
            Discount float64  `json:"discount"`
            EmployeeDiscount float64  `json:"employee_discount"`
            PriceEffective float64  `json:"price_effective"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            ProductID float64  `json:"product_id"`
            AmountPaid float64  `json:"amount_paid"`
            Meta CartItemMeta  `json:"meta"`
         
    }
    
    // OpenApiPlatformCheckoutReq used by Cart
    type OpenApiPlatformCheckoutReq struct {

        
            CouponValue float64  `json:"coupon_value"`
            Files []OpenApiFiles  `json:"files"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            CodCharges float64  `json:"cod_charges"`
            DeliveryCharges float64  `json:"delivery_charges"`
            CashbackApplied float64  `json:"cashback_applied"`
            EmployeeDiscount map[string]interface{}  `json:"employee_discount"`
            Comment string  `json:"comment"`
            Gstin string  `json:"gstin"`
            CartValue float64  `json:"cart_value"`
            ShippingAddress ShippingAddress  `json:"shipping_address"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            Coupon string  `json:"coupon"`
            PaymentMethods []MultiTenderPaymentMethod  `json:"payment_methods"`
            CouponCode string  `json:"coupon_code"`
            OrderID string  `json:"order_id"`
            CurrencyCode string  `json:"currency_code"`
            CartItems []OpenApiOrderItem  `json:"cart_items"`
            BillingAddress ShippingAddress  `json:"billing_address"`
            PaymentMode string  `json:"payment_mode"`
         
    }
    
    // OpenApiCheckoutResponse used by Cart
    type OpenApiCheckoutResponse struct {

        
            OrderID string  `json:"order_id"`
            Success bool  `json:"success"`
            OrderRefID string  `json:"order_ref_id"`
            Message string  `json:"message"`
         
    }
    
    // AbandonedCart used by Cart
    type AbandonedCart struct {

        
            IsActive bool  `json:"is_active"`
            CheckoutMode string  `json:"checkout_mode"`
            BulkCouponDiscount float64  `json:"bulk_coupon_discount"`
            PickUpCustomerDetails map[string]interface{}  `json:"pick_up_customer_details"`
            Discount float64  `json:"discount"`
            Comment string  `json:"comment"`
            Gstin string  `json:"gstin"`
            LastModified string  `json:"last_modified"`
            CartValue float64  `json:"cart_value"`
            Cashback map[string]interface{}  `json:"cashback"`
            UserID string  `json:"user_id"`
            AppID string  `json:"app_id"`
            CreatedOn string  `json:"created_on"`
            Promotion map[string]interface{}  `json:"promotion"`
            Payments map[string]interface{}  `json:"payments"`
            BuyNow bool  `json:"buy_now"`
            Articles []map[string]interface{}  `json:"articles"`
            UID float64  `json:"uid"`
            Shipments []map[string]interface{}  `json:"shipments"`
            IsDefault bool  `json:"is_default"`
            CodCharges map[string]interface{}  `json:"cod_charges"`
            DeliveryCharges map[string]interface{}  `json:"delivery_charges"`
            Meta map[string]interface{}  `json:"meta"`
            MergeQty bool  `json:"merge_qty"`
            Coupon map[string]interface{}  `json:"coupon"`
            PaymentMethods []map[string]interface{}  `json:"payment_methods"`
            FcIndexMap []float64  `json:"fc_index_map"`
            IsArchive bool  `json:"is_archive"`
            OrderID string  `json:"order_id"`
            ExpireAt string  `json:"expire_at"`
            FyndCredits map[string]interface{}  `json:"fynd_credits"`
            ID string  `json:"_id"`
            PaymentMode string  `json:"payment_mode"`
         
    }
    
    // AbandonedCartResponse used by Cart
    type AbandonedCartResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
            Result map[string]interface{}  `json:"result"`
            Items []AbandonedCart  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // CartCurrency used by Cart
    type CartCurrency struct {

        
            Code string  `json:"code"`
            Symbol string  `json:"symbol"`
         
    }
    
    // CartDetailResponse used by Cart
    type CartDetailResponse struct {

        
            CouponText string  `json:"coupon_text"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            ID string  `json:"id"`
            IsValid bool  `json:"is_valid"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            CheckoutMode string  `json:"checkout_mode"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            Currency CartCurrency  `json:"currency"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            Comment string  `json:"comment"`
            Gstin string  `json:"gstin"`
            BuyNow bool  `json:"buy_now"`
            LastModified string  `json:"last_modified"`
            Items []CartProductInfo  `json:"items"`
            Message string  `json:"message"`
         
    }
    
    // AddProductCart used by Cart
    type AddProductCart struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            ArticleID string  `json:"article_id"`
            Quantity float64  `json:"quantity"`
            Display string  `json:"display"`
            ProductGroupTags []string  `json:"product_group_tags"`
            StoreID float64  `json:"store_id"`
            ItemSize string  `json:"item_size"`
            Pos bool  `json:"pos"`
            ArticleAssignment map[string]interface{}  `json:"article_assignment"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            ItemID float64  `json:"item_id"`
            SellerID float64  `json:"seller_id"`
         
    }
    
    // AddCartRequest used by Cart
    type AddCartRequest struct {

        
            Items []AddProductCart  `json:"items"`
         
    }
    
    // AddCartDetailResponse used by Cart
    type AddCartDetailResponse struct {

        
            Partial bool  `json:"partial"`
            Cart CartDetailResponse  `json:"cart"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // UpdateProductCart used by Cart
    type UpdateProductCart struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ArticleID string  `json:"article_id"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            Quantity float64  `json:"quantity"`
            ItemIndex float64  `json:"item_index"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
            ItemSize string  `json:"item_size"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            ItemID float64  `json:"item_id"`
         
    }
    
    // UpdateCartRequest used by Cart
    type UpdateCartRequest struct {

        
            Operation string  `json:"operation"`
            Items []UpdateProductCart  `json:"items"`
         
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
            Source map[string]interface{}  `json:"source"`
            Token string  `json:"token"`
            CreatedOn string  `json:"created_on"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // SharedCart used by Cart
    type SharedCart struct {

        
            CheckoutMode string  `json:"checkout_mode"`
            SharedCartDetails SharedCartDetails  `json:"shared_cart_details"`
            Comment string  `json:"comment"`
            Gstin string  `json:"gstin"`
            LastModified string  `json:"last_modified"`
            Message string  `json:"message"`
            ID string  `json:"id"`
            BuyNow bool  `json:"buy_now"`
            UID string  `json:"uid"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            CartID float64  `json:"cart_id"`
            Currency CartCurrency  `json:"currency"`
            CouponText string  `json:"coupon_text"`
            IsValid bool  `json:"is_valid"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            Items []CartProductInfo  `json:"items"`
         
    }
    
    // SharedCartResponse used by Cart
    type SharedCartResponse struct {

        
            Cart SharedCart  `json:"cart"`
            Error string  `json:"error"`
         
    }
    
    // CartList used by Cart
    type CartList struct {

        
            UserID string  `json:"user_id"`
            CreatedOn string  `json:"created_on"`
            CartID string  `json:"cart_id"`
            ItemCounts string  `json:"item_counts"`
            CartValue string  `json:"cart_value"`
         
    }
    
    // UserCartMappingResponse used by Cart
    type UserCartMappingResponse struct {

        
            User UserInfo  `json:"user"`
            CouponText string  `json:"coupon_text"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            ID string  `json:"id"`
            IsValid bool  `json:"is_valid"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            CheckoutMode string  `json:"checkout_mode"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            Currency CartCurrency  `json:"currency"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            Comment string  `json:"comment"`
            Gstin string  `json:"gstin"`
            BuyNow bool  `json:"buy_now"`
            LastModified string  `json:"last_modified"`
            Items []CartProductInfo  `json:"items"`
            Message string  `json:"message"`
         
    }
    
    // CartItemCountResponse used by Cart
    type CartItemCountResponse struct {

        
            UserCartItemsCount float64  `json:"user_cart_items_count"`
         
    }
    
    // Coupon used by Cart
    type Coupon struct {

        
            IsApplied bool  `json:"is_applied"`
            CouponValue float64  `json:"coupon_value"`
            Title string  `json:"title"`
            MaxDiscountValue float64  `json:"max_discount_value"`
            CouponCode string  `json:"coupon_code"`
            Description string  `json:"description"`
            SubTitle string  `json:"sub_title"`
            IsApplicable bool  `json:"is_applicable"`
            ExpiresOn string  `json:"expires_on"`
            CouponType string  `json:"coupon_type"`
            MinimumCartValue float64  `json:"minimum_cart_value"`
            Message string  `json:"message"`
         
    }
    
    // PageCoupon used by Cart
    type PageCoupon struct {

        
            TotalItemCount float64  `json:"total_item_count"`
            HasPrevious bool  `json:"has_previous"`
            HasNext bool  `json:"has_next"`
            Current float64  `json:"current"`
            Total float64  `json:"total"`
         
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
    
    // Address used by Cart
    type Address struct {

        
            CountryPhoneCode string  `json:"country_phone_code"`
            IsActive bool  `json:"is_active"`
            AddressType string  `json:"address_type"`
            State string  `json:"state"`
            CheckoutMode string  `json:"checkout_mode"`
            GeoLocation GeoLocation  `json:"geo_location"`
            Address string  `json:"address"`
            City string  `json:"city"`
            Country string  `json:"country"`
            Name string  `json:"name"`
            UserID string  `json:"user_id"`
            ID string  `json:"id"`
            Phone string  `json:"phone"`
            Area string  `json:"area"`
            AreaCodeSlug string  `json:"area_code_slug"`
            CreatedByUserID string  `json:"created_by_user_id"`
            AreaCode string  `json:"area_code"`
            Email string  `json:"email"`
            Meta map[string]interface{}  `json:"meta"`
            CountryIsoCode string  `json:"country_iso_code"`
            Landmark string  `json:"landmark"`
            IsDefaultAddress bool  `json:"is_default_address"`
            CountryCode string  `json:"country_code"`
            GoogleMapPoint map[string]interface{}  `json:"google_map_point"`
            Tags []string  `json:"tags"`
         
    }
    
    // GetAddressesResponse used by Cart
    type GetAddressesResponse struct {

        
            Address []Address  `json:"address"`
         
    }
    
    // GeoLocation1 used by Cart
    type GeoLocation1 struct {

        
            Latitude float64  `json:"latitude"`
            Longitude float64  `json:"longitude"`
         
    }
    
    // PlatformAddress used by Cart
    type PlatformAddress struct {

        
            IsActive bool  `json:"is_active"`
            AddressType string  `json:"address_type"`
            State string  `json:"state"`
            CheckoutMode string  `json:"checkout_mode"`
            GeoLocation GeoLocation1  `json:"geo_location"`
            Address string  `json:"address"`
            City string  `json:"city"`
            Country string  `json:"country"`
            Name string  `json:"name"`
            UserID string  `json:"user_id"`
            ID string  `json:"id"`
            Phone string  `json:"phone"`
            Area string  `json:"area"`
            AreaCodeSlug string  `json:"area_code_slug"`
            CreatedByUserID string  `json:"created_by_user_id"`
            CartID string  `json:"cart_id"`
            AreaCode string  `json:"area_code"`
            Email string  `json:"email"`
            Meta map[string]interface{}  `json:"meta"`
            Landmark string  `json:"landmark"`
            IsDefaultAddress bool  `json:"is_default_address"`
            CountryCode string  `json:"country_code"`
            GoogleMapPoint map[string]interface{}  `json:"google_map_point"`
            Tags []string  `json:"tags"`
         
    }
    
    // SaveAddressResponse used by Cart
    type SaveAddressResponse struct {

        
            IsDefaultAddress bool  `json:"is_default_address"`
            Success bool  `json:"success"`
            ID string  `json:"id"`
         
    }
    
    // UpdateAddressResponse used by Cart
    type UpdateAddressResponse struct {

        
            IsDefaultAddress bool  `json:"is_default_address"`
            Success bool  `json:"success"`
            IsUpdated bool  `json:"is_updated"`
            ID string  `json:"id"`
         
    }
    
    // DeleteAddressResponse used by Cart
    type DeleteAddressResponse struct {

        
            IsDeleted bool  `json:"is_deleted"`
            ID string  `json:"id"`
         
    }
    
    // PlatformSelectCartAddressRequest used by Cart
    type PlatformSelectCartAddressRequest struct {

        
            CheckoutMode string  `json:"checkout_mode"`
            CartID string  `json:"cart_id"`
            ID string  `json:"id"`
            BillingAddressID string  `json:"billing_address_id"`
         
    }
    
    // ShipmentResponse used by Cart
    type ShipmentResponse struct {

        
            ShipmentType string  `json:"shipment_type"`
            FulfillmentID float64  `json:"fulfillment_id"`
            DpID string  `json:"dp_id"`
            FulfillmentType string  `json:"fulfillment_type"`
            Promise ShipmentPromise  `json:"promise"`
            BoxType string  `json:"box_type"`
            DpOptions map[string]interface{}  `json:"dp_options"`
            Items []CartProductInfo  `json:"items"`
            OrderType string  `json:"order_type"`
            Shipments float64  `json:"shipments"`
         
    }
    
    // CartShipmentsResponse used by Cart
    type CartShipmentsResponse struct {

        
            BreakupValues CartBreakup  `json:"breakup_values"`
            CheckoutMode string  `json:"checkout_mode"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            CartID float64  `json:"cart_id"`
            Currency CartCurrency  `json:"currency"`
            Comment string  `json:"comment"`
            Gstin string  `json:"gstin"`
            LastModified string  `json:"last_modified"`
            Error bool  `json:"error"`
            Message string  `json:"message"`
            CouponText string  `json:"coupon_text"`
            ID string  `json:"id"`
            IsValid bool  `json:"is_valid"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            BuyNow bool  `json:"buy_now"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            UID string  `json:"uid"`
            Shipments []ShipmentResponse  `json:"shipments"`
         
    }
    
    // CartMetaRequest used by Cart
    type CartMetaRequest struct {

        
            PickUpCustomerDetails map[string]interface{}  `json:"pick_up_customer_details"`
            CheckoutMode string  `json:"checkout_mode"`
            Comment string  `json:"comment"`
            Gstin string  `json:"gstin"`
         
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
            ID string  `json:"_id"`
            LastName string  `json:"last_name"`
            FirstName string  `json:"first_name"`
         
    }
    
    // PlatformCartCheckoutDetailRequest used by Cart
    type PlatformCartCheckoutDetailRequest struct {

        
            UserID string  `json:"user_id"`
            PaymentIdentifier string  `json:"payment_identifier"`
            Staff StaffCheckout  `json:"staff"`
            Aggregator string  `json:"aggregator"`
            OrderingStore float64  `json:"ordering_store"`
            PaymentAutoConfirm bool  `json:"payment_auto_confirm"`
            ID string  `json:"id"`
            PaymentParams map[string]interface{}  `json:"payment_params"`
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
            Meta map[string]interface{}  `json:"meta"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            AddressID string  `json:"address_id"`
            MerchantCode string  `json:"merchant_code"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
            PaymentMode string  `json:"payment_mode"`
            CallbackURL string  `json:"callback_url"`
            BillingAddressID string  `json:"billing_address_id"`
         
    }
    
    // CheckCart used by Cart
    type CheckCart struct {

        
            StoreEmps []map[string]interface{}  `json:"store_emps"`
            CheckoutMode string  `json:"checkout_mode"`
            Comment string  `json:"comment"`
            Gstin string  `json:"gstin"`
            LastModified string  `json:"last_modified"`
            UserType string  `json:"user_type"`
            Message string  `json:"message"`
            ID string  `json:"id"`
            BuyNow bool  `json:"buy_now"`
            UID string  `json:"uid"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            CodCharges float64  `json:"cod_charges"`
            DeliveryCharges float64  `json:"delivery_charges"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            CartID float64  `json:"cart_id"`
            Currency CartCurrency  `json:"currency"`
            DeliveryChargeOrderValue float64  `json:"delivery_charge_order_value"`
            Success bool  `json:"success"`
            CodAvailable bool  `json:"cod_available"`
            CouponText string  `json:"coupon_text"`
            IsValid bool  `json:"is_valid"`
            ErrorMessage string  `json:"error_message"`
            StoreCode string  `json:"store_code"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            OrderID string  `json:"order_id"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            Items []CartProductInfo  `json:"items"`
            CodMessage string  `json:"cod_message"`
         
    }
    
    // CartCheckoutResponse used by Cart
    type CartCheckoutResponse struct {

        
            Cart CheckCart  `json:"cart"`
            AppInterceptURL string  `json:"app_intercept_url"`
            Data map[string]interface{}  `json:"data"`
            OrderID string  `json:"order_id"`
            Success bool  `json:"success"`
            PaymentConfirmURL string  `json:"payment_confirm_url"`
            CallbackURL string  `json:"callback_url"`
            Message string  `json:"message"`
         
    }
    
    // CartDeliveryModesResponse used by Cart
    type CartDeliveryModesResponse struct {

        
            PickupStores []float64  `json:"pickup_stores"`
            AvailableModes []string  `json:"available_modes"`
         
    }
    
    // PickupStoreDetail used by Cart
    type PickupStoreDetail struct {

        
            AddressType string  `json:"address_type"`
            State string  `json:"state"`
            ID float64  `json:"id"`
            Phone string  `json:"phone"`
            StoreCode string  `json:"store_code"`
            UID float64  `json:"uid"`
            Address string  `json:"address"`
            Landmark string  `json:"landmark"`
            AreaCode string  `json:"area_code"`
            City string  `json:"city"`
            Area string  `json:"area"`
            AreaCodeSlug string  `json:"area_code_slug"`
            Pincode float64  `json:"pincode"`
            Email string  `json:"email"`
            Country string  `json:"country"`
            Name string  `json:"name"`
         
    }
    
    // StoreDetailsResponse used by Cart
    type StoreDetailsResponse struct {

        
            Items []PickupStoreDetail  `json:"items"`
         
    }
    
    // UpdateCartPaymentRequest used by Cart
    type UpdateCartPaymentRequest struct {

        
            PaymentIdentifier string  `json:"payment_identifier"`
            ID string  `json:"id"`
            AggregatorName string  `json:"aggregator_name"`
            AddressID string  `json:"address_id"`
            MerchantCode string  `json:"merchant_code"`
            PaymentMode string  `json:"payment_mode"`
         
    }
    
    // CouponValidity used by Cart
    type CouponValidity struct {

        
            Title string  `json:"title"`
            Discount float64  `json:"discount"`
            DisplayMessageEn string  `json:"display_message_en"`
            Valid bool  `json:"valid"`
            Code string  `json:"code"`
         
    }
    
    // PaymentCouponValidate used by Cart
    type PaymentCouponValidate struct {

        
            CouponValidity CouponValidity  `json:"coupon_validity"`
            Success bool  `json:"success"`
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
    
    // RewardsRule used by Rewards
    type RewardsRule struct {

        
            Amount float64  `json:"amount"`
         
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
    
    // DiscountItems used by Discount
    type DiscountItems struct {

        
            ItemCode string  `json:"item_code"`
            BrandUID float64  `json:"brand_uid"`
            SellerIdentifier string  `json:"seller_identifier"`
            DiscountType string  `json:"discount_type"`
            Value float64  `json:"value"`
         
    }
    
    // BulkDiscount used by Discount
    type BulkDiscount struct {

        
            CompanyID float64  `json:"company_id"`
            Items []DiscountItems  `json:"items"`
         
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
    
    // APIError used by Partner
    type APIError struct {

        
            Code string  `json:"code"`
            Message string  `json:"message"`
            Info string  `json:"info"`
            RequestID string  `json:"request_id"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // RemoveProxyResponse used by Partner
    type RemoveProxyResponse struct {

        
            Message string  `json:"message"`
            Data map[string]interface{}  `json:"data"`
         
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
    
    // EntityTypesResponse used by AuditTrail
    type EntityTypesResponse struct {

        
            Items []EntityTypeObj  `json:"items"`
         
    }
    
    // EntityTypeObj used by AuditTrail
    type EntityTypeObj struct {

        
            EntityValue string  `json:"entity_value"`
            DisplayName string  `json:"display_name"`
         
    }
    
