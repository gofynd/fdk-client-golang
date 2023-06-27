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
            Type string  `json:"type"`
         
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
    

    
    // ThemeReq used by Theme
    type ThemeReq struct {

        
            MarketplaceThemeID string  `json:"marketplace_theme_id"`
         
    }
    
    // ThemeSchema used by Theme
    type ThemeSchema struct {

        
            ID string  `json:"_id"`
            Name string  `json:"name"`
            MarketplaceThemeID MarketplaceThemeId  `json:"marketplace_theme_id"`
            CompanyID float64  `json:"company_id"`
            Meta ThemeMeta  `json:"meta"`
            CreatedAt string  `json:"created_at"`
            UpdatedAt string  `json:"updated_at"`
         
    }
    
    // MarketplaceThemeId used by Theme
    type MarketplaceThemeId struct {

        
            ID string  `json:"_id"`
            IsDefault bool  `json:"is_default"`
         
    }
    
    // ThemeMeta used by Theme
    type ThemeMeta struct {

        
            Payment ThemePayment  `json:"payment"`
            Industry []string  `json:"industry"`
            Description string  `json:"description"`
            Images ThemeImages  `json:"images"`
            Slug string  `json:"slug"`
         
    }
    
    // ThemePayment used by Theme
    type ThemePayment struct {

        
            IsPaid bool  `json:"is_paid"`
            Amount float64  `json:"amount"`
         
    }
    
    // ThemeImages used by Theme
    type ThemeImages struct {

        
            Desktop string  `json:"desktop"`
            Mobile string  `json:"mobile"`
         
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
    
    // MarketplaceThemeResponse used by Theme
    type MarketplaceThemeResponse struct {

        
            Status float64  `json:"status"`
            Body MarketplaceThemeResponseBody  `json:"body"`
         
    }
    
    // MarketplaceThemeResponseBody used by Theme
    type MarketplaceThemeResponseBody struct {

        
            Themes []MarketplaceTheme  `json:"themes"`
            Page PaginationSchema  `json:"page"`
         
    }
    
    // ArrayOfMarketplaceTheme used by Theme
    type ArrayOfMarketplaceTheme struct {

        
            Body []MarketplaceTheme  `json:"body"`
         
    }
    
    // ThemeCreateRequest used by Theme
    type ThemeCreateRequest struct {

        
            Src string  `json:"src"`
            Release Release  `json:"release"`
         
    }
    
    // MarketplaceTheme used by Theme
    type MarketplaceTheme struct {

        
            ID string  `json:"_id"`
            Payment PaymentInfo  `json:"payment"`
            Contact ContactInfo  `json:"contact"`
            Industry []string  `json:"industry"`
            IsUpdate bool  `json:"is_update"`
            IsDefault bool  `json:"is_default"`
            Name string  `json:"name"`
            Tagline string  `json:"tagline"`
            Description string  `json:"description"`
            CatalogSize CatalogSize  `json:"catalog_size"`
            Images MarketplaceThemeImages  `json:"images"`
            Carousel []CarouselItem  `json:"carousel"`
            Src string  `json:"src"`
            Explore ExploreInfo  `json:"explore"`
            Features []Feature  `json:"features"`
            Highlights []Highlight  `json:"highlights"`
            Variations []Variation  `json:"variations"`
            Documentation Documentation  `json:"documentation"`
            Status string  `json:"status"`
            Step float64  `json:"step"`
            Comments Comments  `json:"comments"`
            Release Release  `json:"release"`
            Slug string  `json:"slug"`
            OrganizationID string  `json:"organization_id"`
            UserID string  `json:"user_id"`
            CreatedAt string  `json:"created_at"`
            UpdatedAt string  `json:"updated_at"`
            TemplateThemeID string  `json:"template_theme_id"`
         
    }
    
    // PaymentInfo used by Theme
    type PaymentInfo struct {

        
            IsPaid bool  `json:"is_paid"`
            Amount float64  `json:"amount"`
         
    }
    
    // ContactInfo used by Theme
    type ContactInfo struct {

        
            DeveloperContact []string  `json:"developer_contact"`
            SellerContact string  `json:"seller_contact"`
         
    }
    
    // CatalogSize used by Theme
    type CatalogSize struct {

        
            Min float64  `json:"min"`
            Max float64  `json:"max"`
         
    }
    
    // MarketplaceThemeImages used by Theme
    type MarketplaceThemeImages struct {

        
            Desktop string  `json:"desktop"`
            Mobile string  `json:"mobile"`
         
    }
    
    // CarouselItem used by Theme
    type CarouselItem struct {

        
            Desktop string  `json:"desktop"`
            Mobile string  `json:"mobile"`
         
    }
    
    // ExploreInfo used by Theme
    type ExploreInfo struct {

        
            Title string  `json:"title"`
            Description string  `json:"description"`
         
    }
    
    // Feature used by Theme
    type Feature struct {

        
            Category string  `json:"category"`
            List []FeatureItem  `json:"list"`
         
    }
    
    // FeatureItem used by Theme
    type FeatureItem struct {

        
            Label string  `json:"label"`
            Description string  `json:"description"`
         
    }
    
    // Highlight used by Theme
    type Highlight struct {

        
            Title string  `json:"title"`
            Description string  `json:"description"`
            Image string  `json:"image"`
         
    }
    
    // Variation used by Theme
    type Variation struct {

        
            Name string  `json:"name"`
            Color string  `json:"color"`
            DemoURL string  `json:"demo_url"`
            Images MarketplaceThemeImages  `json:"images"`
         
    }
    
    // Documentation used by Theme
    type Documentation struct {

        
            Notes string  `json:"notes"`
            URL string  `json:"url"`
         
    }
    
    // Comments used by Theme
    type Comments struct {

        
            DeveloperRemark string  `json:"developer_remark"`
            ReviewerFeedback string  `json:"reviewer_feedback"`
         
    }
    
    // Release used by Theme
    type Release struct {

        
            Version string  `json:"version"`
            Notes string  `json:"notes"`
         
    }
    
    // ThemeSlugResponse used by Theme
    type ThemeSlugResponse struct {

        
            Theme MarketplaceTheme  `json:"theme"`
            Organization Organization  `json:"organization"`
            User []ThemeCreator  `json:"user"`
         
    }
    
    // Organization used by Theme
    type Organization struct {

        
            Meta OrganizationMeta  `json:"meta"`
            ID string  `json:"_id"`
         
    }
    
    // OrganizationMeta used by Theme
    type OrganizationMeta struct {

        
            EcommPlatformUsed []string  `json:"ecomm_platform_used"`
            Goals []string  `json:"goals"`
         
    }
    
    // ThemeCreator used by Theme
    type ThemeCreator struct {

        
            ID string  `json:"_id"`
            Gender string  `json:"gender"`
            AccountType string  `json:"account_type"`
            Active bool  `json:"active"`
            FirstName string  `json:"first_name"`
            LastName string  `json:"last_name"`
            PhoneNumbers []PhoneNumber  `json:"phone_numbers"`
            Emails []Email  `json:"emails"`
         
    }
    
    // PhoneNumber used by Theme
    type PhoneNumber struct {

        
            Active bool  `json:"active"`
            Primary bool  `json:"primary"`
            Verified bool  `json:"verified"`
            Phone string  `json:"phone"`
            CountryCode float64  `json:"country_code"`
         
    }
    
    // Email used by Theme
    type Email struct {

        
            Active bool  `json:"active"`
            Primary bool  `json:"primary"`
            Verified bool  `json:"verified"`
            Email string  `json:"email"`
         
    }
    
    // ThemeAndUserDetailsResponse used by Theme
    type ThemeAndUserDetailsResponse struct {

        
            Themes []MarketplaceTheme  `json:"themes"`
            User []ThemeCreator  `json:"user"`
         
    }
    
    // ThemeRejectionReasons used by Theme
    type ThemeRejectionReasons struct {

        
            ID string  `json:"_id"`
            Message string  `json:"message"`
            ThemeID string  `json:"theme_id"`
            OrganizationID string  `json:"organization_id"`
            AdminID string  `json:"admin_id"`
            UserID string  `json:"user_id"`
            Status string  `json:"status"`
            RejectionReasons map[string]map[string]interface{}  `json:"rejection_reasons"`
            CreatedAt string  `json:"created_at"`
            UpdatedAt string  `json:"updated_at"`
         
    }
    
    // RejectionReason used by Theme
    type RejectionReason struct {

        
            Message string  `json:"message"`
         
    }
    
    // ThemeReviewRequest used by Theme
    type ThemeReviewRequest struct {

        
            DynamicProperties map[string]map[string]interface{}  `json:"dynamic_properties"`
         
    }
    
    // UpdateReviewStatusRequest used by Theme
    type UpdateReviewStatusRequest struct {

        
            Status string  `json:"status"`
         
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
    
    // ApplyThemeRequestV2 used by Theme
    type ApplyThemeRequestV2 struct {

        
            MarketplaceThemeID string  `json:"marketplace_theme_id"`
         
    }
    
    // ApplyThemeResponseV2 used by Theme
    type ApplyThemeResponseV2 struct {

        
            Font FontV2  `json:"font"`
            Config ConfigV2  `json:"config"`
            Applied bool  `json:"applied"`
            IsPrivate bool  `json:"is_private"`
            Tags []string  `json:"tags"`
            ID string  `json:"_id"`
            ApplicationID string  `json:"application_id"`
            MarketplaceThemeID string  `json:"marketplace_theme_id"`
            Meta MetaV2  `json:"meta"`
            Name string  `json:"name"`
            TemplateThemeID string  `json:"template_theme_id"`
            Version string  `json:"version"`
            Styles map[string]interface{}  `json:"styles"`
            CreatedAt string  `json:"created_at"`
            UpdatedAt string  `json:"updated_at"`
         
    }
    
    // AllThemesApplicationResponseV2 used by Theme
    type AllThemesApplicationResponseV2 struct {

        
            Font FontV2  `json:"font"`
            Config ConfigV2  `json:"config"`
            Applied bool  `json:"applied"`
            IsPrivate bool  `json:"is_private"`
            Tags []string  `json:"tags"`
            ID string  `json:"_id"`
            ApplicationID string  `json:"application_id"`
            MarketplaceThemeID string  `json:"marketplace_theme_id"`
            Meta MetaV2  `json:"meta"`
            Name string  `json:"name"`
            TemplateThemeID string  `json:"template_theme_id"`
            Version string  `json:"version"`
            Styles map[string]interface{}  `json:"styles"`
            CreatedAt string  `json:"created_at"`
            UpdatedAt string  `json:"updated_at"`
            Assets AssetsV2  `json:"assets"`
            AvailableSections []SectionItem  `json:"available_sections"`
         
    }
    
    // ThemeUpgradableResponseV2 used by Theme
    type ThemeUpgradableResponseV2 struct {

        
            Upgrade bool  `json:"upgrade"`
            Versions ThemeVersionsV2  `json:"versions"`
            Message string  `json:"message"`
         
    }
    
    // UpdateThemeNameRequestBodyV2 used by Theme
    type UpdateThemeNameRequestBodyV2 struct {

        
            Name string  `json:"name"`
         
    }
    
    // UpdateThemeRequestBodyV2 used by Theme
    type UpdateThemeRequestBodyV2 struct {

        
            Config ConfigV2  `json:"config"`
            Font FontV2  `json:"font"`
         
    }
    
    // FontV2 used by Theme
    type FontV2 struct {

        
            Variants FontVariantsV2  `json:"variants"`
            Family string  `json:"family"`
         
    }
    
    // FontVariantsV2 used by Theme
    type FontVariantsV2 struct {

        
            Light FontVariantV2  `json:"light"`
            Regular FontVariantV2  `json:"regular"`
            Medium FontVariantV2  `json:"medium"`
            SemiBold FontVariantV2  `json:"semi_bold"`
            Bold FontVariantV2  `json:"bold"`
         
    }
    
    // FontVariantV2 used by Theme
    type FontVariantV2 struct {

        
            Name string  `json:"name"`
            File string  `json:"file"`
         
    }
    
    // ConfigV2 used by Theme
    type ConfigV2 struct {

        
            Current string  `json:"current"`
            List []ConfigurationV2  `json:"list"`
            GlobalSchema GlobalSchemaV2  `json:"global_schema"`
            Preset PresetV2  `json:"preset"`
         
    }
    
    // ConfigurationV2 used by Theme
    type ConfigurationV2 struct {

        
            Name string  `json:"name"`
            GlobalConfig GlobalConfigV2  `json:"global_config"`
            Custom CustomConfigV2  `json:"custom"`
            Page []string  `json:"page"`
         
    }
    
    // GlobalConfigV2 used by Theme
    type GlobalConfigV2 struct {

        
            Statics StaticConfigV2  `json:"statics"`
            Auth AuthConfigV2  `json:"auth"`
            Palette PaletteConfigV2  `json:"palette"`
         
    }
    
    // StaticConfigV2 used by Theme
    type StaticConfigV2 struct {

        
            Props StaticPropsV2  `json:"props"`
         
    }
    
    // AuthConfigV2 used by Theme
    type AuthConfigV2 struct {

        
            ShowHeaderAuth bool  `json:"show_header_auth"`
            ShowFooterAuth bool  `json:"show_footer_auth"`
         
    }
    
    // PaletteConfigV2 used by Theme
    type PaletteConfigV2 struct {

        
            GeneralSetting GeneralSettingV2  `json:"general_setting"`
            AdvanceSetting AdvanceSettingV2  `json:"advance_setting"`
         
    }
    
    // CustomConfigV2 used by Theme
    type CustomConfigV2 struct {

        
            Props CustomPropsV2  `json:"props"`
         
    }
    
    // MetaV2 used by Theme
    type MetaV2 struct {

        
            Payment PaymentV2  `json:"payment"`
            Description string  `json:"description"`
            Industry []string  `json:"industry"`
            Release ReleaseV2  `json:"release"`
            Images ImagesV2  `json:"images"`
            Slug string  `json:"slug"`
            Name string  `json:"name"`
         
    }
    
    // PaymentV2 used by Theme
    type PaymentV2 struct {

        
            IsPaid bool  `json:"is_paid"`
            Amount float64  `json:"amount"`
         
    }
    
    // ReleaseV2 used by Theme
    type ReleaseV2 struct {

        
            Notes string  `json:"notes"`
            Version string  `json:"version"`
         
    }
    
    // ImagesV2 used by Theme
    type ImagesV2 struct {

        
            Desktop string  `json:"desktop"`
            Mobile string  `json:"mobile"`
         
    }
    
    // StaticPropsV2 used by Theme
    type StaticPropsV2 struct {

        
            Colors ColorsV2  `json:"colors"`
            Auth AuthConfigV2  `json:"auth"`
         
    }
    
    // ColorsV2 used by Theme
    type ColorsV2 struct {

        
            PrimaryColor string  `json:"primary_color"`
            SecondaryColor string  `json:"secondary_color"`
            AccentColor string  `json:"accent_color"`
            LinkColor string  `json:"link_color"`
            ButtonSecondaryColor string  `json:"button_secondary_color"`
            BgColor string  `json:"bg_color"`
         
    }
    
    // GeneralSettingV2 used by Theme
    type GeneralSettingV2 struct {

        
            Theme ThemeSettingV2  `json:"theme"`
            Text TextSettingV2  `json:"text"`
            Button ButtonSettingV2  `json:"button"`
            SaleDiscount SaleDiscountSettingV2  `json:"sale_discount"`
            Header HeaderSettingV2  `json:"header"`
            Footer FooterSettingV2  `json:"footer"`
         
    }
    
    // AdvanceSettingV2 used by Theme
    type AdvanceSettingV2 struct {

        
            OverlayPopup OverlayPopupSettingV2  `json:"overlay_popup"`
            DividerStrokeHighlight DividerStrokeHighlightSettingV2  `json:"divider_stroke_highlight"`
            UserAlerts UserAlertsSettingV2  `json:"user_alerts"`
         
    }
    
    // ThemeSettingV2 used by Theme
    type ThemeSettingV2 struct {

        
            PageBackground string  `json:"page_background"`
            ThemeAccent string  `json:"theme_accent"`
         
    }
    
    // TextSettingV2 used by Theme
    type TextSettingV2 struct {

        
            TextHeading string  `json:"text_heading"`
            TextBody string  `json:"text_body"`
            TextLabel string  `json:"text_label"`
            TextSecondary string  `json:"text_secondary"`
         
    }
    
    // ButtonSettingV2 used by Theme
    type ButtonSettingV2 struct {

        
            ButtonPrimary string  `json:"button_primary"`
            ButtonSecondary string  `json:"button_secondary"`
            ButtonLink string  `json:"button_link"`
         
    }
    
    // SaleDiscountSettingV2 used by Theme
    type SaleDiscountSettingV2 struct {

        
            SaleBadgeBackground string  `json:"sale_badge_background"`
            SaleBadgeText string  `json:"sale_badge_text"`
            SaleDiscountText string  `json:"sale_discount_text"`
            SaleTimer string  `json:"sale_timer"`
         
    }
    
    // HeaderSettingV2 used by Theme
    type HeaderSettingV2 struct {

        
            HeaderBackground string  `json:"header_background"`
            HeaderNav string  `json:"header_nav"`
            HeaderIcon string  `json:"header_icon"`
         
    }
    
    // FooterSettingV2 used by Theme
    type FooterSettingV2 struct {

        
            FooterBackground string  `json:"footer_background"`
            FooterBottomBackground string  `json:"footer_bottom_background"`
            FooterHeadingText string  `json:"footer_heading_text"`
            FooterBodyText string  `json:"footer_body_text"`
            FooterIcon string  `json:"footer_icon"`
         
    }
    
    // OverlayPopupSettingV2 used by Theme
    type OverlayPopupSettingV2 struct {

        
            DialogBackgroung string  `json:"dialog_backgroung"`
            Overlay string  `json:"overlay"`
         
    }
    
    // DividerStrokeHighlightSettingV2 used by Theme
    type DividerStrokeHighlightSettingV2 struct {

        
            DividerStrokes string  `json:"divider_strokes"`
            Highlight string  `json:"highlight"`
         
    }
    
    // UserAlertsSettingV2 used by Theme
    type UserAlertsSettingV2 struct {

        
            SuccessBackground string  `json:"success_background"`
            SuccessText string  `json:"success_text"`
            ErrorBackground string  `json:"error_background"`
            ErrorText string  `json:"error_text"`
            InfoBackground string  `json:"info_background"`
            InfoText string  `json:"info_text"`
         
    }
    
    // CustomPropsV2 used by Theme
    type CustomPropsV2 struct {

        
            HeaderBgColor string  `json:"header_bg_color"`
            HeaderTextColor string  `json:"header_text_color"`
            HeaderBorderColor string  `json:"header_border_color"`
            HeaderIconColor string  `json:"header_icon_color"`
            HeaderCartNotificationBgColor string  `json:"header_cart_notification_bg_color"`
            HeaderCartNotificationTextColor string  `json:"header_cart_notification_text_color"`
            HeaderNavHoverColor string  `json:"header_nav_hover_color"`
            ButtonPrimaryColor string  `json:"button_primary_color"`
            ButtonPrimaryLabelColor string  `json:"button_primary_label_color"`
            ButtonAddToCartColor string  `json:"button_add_to_cart_color"`
            ButtonAddToCartLabelColor string  `json:"button_add_to_cart_label_color"`
            ButtonSecondaryColor string  `json:"button_secondary_color"`
            ButtonSecondaryLabelColor string  `json:"button_secondary_label_color"`
            ButtonTertiaryColor string  `json:"button_tertiary_color"`
            ButtonTertiaryLabelColor string  `json:"button_tertiary_label_color"`
            ButtonTertiaryHoverColor string  `json:"button_tertiary_hover_color"`
            ButtonTertiaryHoverTextColor string  `json:"button_tertiary_hover_text_color"`
            TextHeadingLinkColor string  `json:"text_heading_link_color"`
            TextBodyColor string  `json:"text_body_color"`
            TextPriceColor string  `json:"text_price_color"`
            TextSalePriceColor string  `json:"text_sale_price_color"`
            TextStrikethroughPriceColor string  `json:"text_strikethrough_price_color"`
            TextDiscountColor string  `json:"text_discount_color"`
            FooterBgColor string  `json:"footer_bg_color"`
            FooterTextColor string  `json:"footer_text_color"`
            FooterBorderColor string  `json:"footer_border_color"`
            FooterNavHoverColor string  `json:"footer_nav_hover_color"`
            DisableCart bool  `json:"disable_cart"`
            IsMenuBelowLogo bool  `json:"is_menu_below_logo"`
            MenuPosition string  `json:"menu_position"`
         
    }
    
    // GlobalSchemaV2 used by Theme
    type GlobalSchemaV2 struct {

        
            Props []PropV2  `json:"props"`
         
    }
    
    // PropV2 used by Theme
    type PropV2 struct {

        
            Type string  `json:"type"`
            Category string  `json:"category"`
            ID string  `json:"id"`
            Label string  `json:"label"`
            Info string  `json:"info"`
         
    }
    
    // AssetsV2 used by Theme
    type AssetsV2 struct {

        
            UmdJs UMDJs  `json:"umd_js"`
            CommonJs CommonJS  `json:"common_js"`
            Css CSS  `json:"css"`
         
    }
    
    // UMDJs used by Theme
    type UMDJs struct {

        
            Links []string  `json:"links"`
         
    }
    
    // CommonJS used by Theme
    type CommonJS struct {

        
            Link string  `json:"link"`
         
    }
    
    // CSS used by Theme
    type CSS struct {

        
            Links []string  `json:"links"`
         
    }
    
    // SectionItem used by Theme
    type SectionItem struct {

        
            Props []interface{}  `json:"props"`
            Blocks []interface{}  `json:"blocks"`
            Name string  `json:"name"`
            Label string  `json:"label"`
         
    }
    
    // PresetV2 used by Theme
    type PresetV2 struct {

        
            Pages []PageV2  `json:"pages"`
         
    }
    
    // PageV2 used by Theme
    type PageV2 struct {

        
            Sections []SectionV2  `json:"sections"`
            Value string  `json:"value"`
         
    }
    
    // SectionV2 used by Theme
    type SectionV2 struct {

        
            Blocks []BlockV2  `json:"blocks"`
            Predicate PredicateV2  `json:"predicate"`
            Name string  `json:"name"`
            Props SectionPropsV2  `json:"props"`
            Preset SectionPresetV2  `json:"preset"`
         
    }
    
    // BlockV2 used by Theme
    type BlockV2 struct {

        
            Type string  `json:"type"`
            Name string  `json:"name"`
            Props BlockPropsV2  `json:"props"`
         
    }
    
    // PredicateV2 used by Theme
    type PredicateV2 struct {

        
            Screen ScreenV2  `json:"screen"`
            User UserV2  `json:"user"`
            Route RouteV2  `json:"route"`
         
    }
    
    // ScreenV2 used by Theme
    type ScreenV2 struct {

        
            Mobile bool  `json:"mobile"`
            Desktop bool  `json:"desktop"`
            Tablet bool  `json:"tablet"`
         
    }
    
    // UserV2 used by Theme
    type UserV2 struct {

        
            Authenticated bool  `json:"authenticated"`
            Anonymous bool  `json:"anonymous"`
         
    }
    
    // RouteV2 used by Theme
    type RouteV2 struct {

        
            Selected string  `json:"selected"`
            ExactURL string  `json:"exact_url"`
         
    }
    
    // SectionPropsV2 used by Theme
    type SectionPropsV2 struct {

        
            Title TextPropV2  `json:"title"`
            ItemMargin TextPropV2  `json:"item_margin"`
            Autoplay CheckboxPropV2  `json:"autoplay"`
            SlideInterval RangePropV2  `json:"slide_interval"`
         
    }
    
    // SectionPresetV2 used by Theme
    type SectionPresetV2 struct {

        
            Blocks []BlockV2  `json:"blocks"`
         
    }
    
    // BlockPropsV2 used by Theme
    type BlockPropsV2 struct {

        
            Image ImagePickerPropV2  `json:"image"`
            SlideLink UrlPropV2  `json:"slide_link"`
         
    }
    
    // TextPropV2 used by Theme
    type TextPropV2 struct {

        
            Value string  `json:"value"`
            Type string  `json:"type"`
         
    }
    
    // CheckboxPropV2 used by Theme
    type CheckboxPropV2 struct {

        
            Value bool  `json:"value"`
            Type string  `json:"type"`
         
    }
    
    // RangePropV2 used by Theme
    type RangePropV2 struct {

        
            Value float64  `json:"value"`
            Type string  `json:"type"`
         
    }
    
    // ImagePickerPropV2 used by Theme
    type ImagePickerPropV2 struct {

        
            Type string  `json:"type"`
            Value string  `json:"value"`
         
    }
    
    // UrlPropV2 used by Theme
    type UrlPropV2 struct {

        
            Type string  `json:"type"`
            Value string  `json:"value"`
         
    }
    
    // ThemeVersionsV2 used by Theme
    type ThemeVersionsV2 struct {

        
            ParentTheme string  `json:"parent_theme"`
            AppliedTheme string  `json:"applied_theme"`
         
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
            CaptchaCode string  `json:"captcha_code"`
         
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
            Type string  `json:"type"`
         
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
            Type string  `json:"type"`
         
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

        
            Items []SessionListResponseInfo  `json:"items"`
         
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
    
    // SessionListResponseInfo used by User
    type SessionListResponseInfo struct {

        
            SessionID string  `json:"session_id"`
            UserAgent string  `json:"user_agent"`
            Ip string  `json:"ip"`
            Domain string  `json:"domain"`
            ExpireIn string  `json:"expire_in"`
         
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
    
    // UserGroupResponseSchema used by User
    type UserGroupResponseSchema struct {

        
            Name string  `json:"name"`
            Description string  `json:"description"`
            FileURL string  `json:"file_url"`
            ID string  `json:"_id"`
            Status string  `json:"status"`
            UID float64  `json:"uid"`
            ApplicationID string  `json:"application_id"`
            CreatedAt string  `json:"created_at"`
            ModifiedAt string  `json:"modified_at"`
            V float64  `json:"__v"`
         
    }
    
    // UserGroupListResponseSchema used by User
    type UserGroupListResponseSchema struct {

        
            Items []UserGroupResponseSchema  `json:"items"`
            Page PaginationSchema  `json:"page"`
         
    }
    
    // CreateUserGroupSchema used by User
    type CreateUserGroupSchema struct {

        
            Name string  `json:"name"`
            Description string  `json:"description"`
            FileURL string  `json:"file_url"`
         
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
            VoiceOtp bool  `json:"voice_otp"`
         
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
    
    // UpdateUserGroupSchema used by User
    type UpdateUserGroupSchema struct {

        
            Name string  `json:"name"`
            Description string  `json:"description"`
            FileURL string  `json:"file_url"`
         
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
    

    
    // GenerateSEOContent used by Content
    type GenerateSEOContent struct {

        
            Text string  `json:"text"`
            ExistingText string  `json:"existing_text"`
            Keywords []string  `json:"keywords"`
            Type string  `json:"type"`
         
    }
    
    // GeneratedSEOContent used by Content
    type GeneratedSEOContent struct {

        
            Title string  `json:"title"`
            Description string  `json:"description"`
         
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
            Source PathSourceSchema  `json:"__source"`
         
    }
    
    // PathSourceSchema used by Content
    type PathSourceSchema struct {

        
            Type string  `json:"type"`
            ID string  `json:"id"`
         
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
            CannonicalEnabled bool  `json:"cannonical_enabled"`
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
            ImageURL string  `json:"image_url"`
         
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
            Media []SlideshowMedia  `json:"media"`
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
            Config map[string]interface{}  `json:"config"`
         
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
            ChannelType string  `json:"channel_type"`
         
    }
    
    // SubscriptionStatus used by Billing
    type SubscriptionStatus struct {

        
            IsEnabled bool  `json:"is_enabled"`
            Subscription Subscription  `json:"subscription"`
            LatestInvoice InvoicesData  `json:"latest_invoice"`
            NextPlan Plan  `json:"next_plan"`
            CurrentSubscriptions []Subscription  `json:"current_subscriptions"`
            MandateAmount string  `json:"mandate_amount"`
         
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
    
    // EmailProviderAdmin used by Communication
    type EmailProviderAdmin struct {

        
            Type string  `json:"type"`
            Provider string  `json:"provider"`
            FromAddress []EmailProviderReqFrom  `json:"from_address"`
            ID string  `json:"_id"`
            Name string  `json:"name"`
            Description string  `json:"description"`
            APIKey string  `json:"api_key"`
            CreatedAt string  `json:"created_at"`
            UpdatedAt string  `json:"updated_at"`
            Slug string  `json:"slug"`
            V float64  `json:"__v"`
         
    }
    
    // EmailProvidersAdmin used by Communication
    type EmailProvidersAdmin struct {

        
            Items []EmailProvider  `json:"items"`
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
    
    // DefaultEmailProviders used by Communication
    type DefaultEmailProviders struct {

        
            ID string  `json:"_id"`
            FromAddress []DefaultEmailProvidersObjFrom  `json:"from_address"`
            Name string  `json:"name"`
            IsDefault bool  `json:"is_default"`
         
    }
    
    // DefaultEmailProvidersObjFrom used by Communication
    type DefaultEmailProvidersObjFrom struct {

        
            Name string  `json:"name"`
            Email string  `json:"email"`
            IsDefault bool  `json:"is_default"`
         
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
    
    // TriggerJobResponse used by Communication
    type TriggerJobResponse struct {

        
            Status float64  `json:"status"`
         
    }
    
    // TriggerJobRequest used by Communication
    type TriggerJobRequest struct {

        
            JobID string  `json:"job_id"`
         
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
            Source string  `json:"source"`
            ChannelType string  `json:"channel_type"`
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
    
    // SendOtpEmailCommsProvider used by Communication
    type SendOtpEmailCommsProvider struct {

        
            Slug string  `json:"slug"`
            ID string  `json:"_id"`
         
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
            Provider SendOtpEmailCommsProvider  `json:"provider"`
         
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
    
    // AppProvider used by Communication
    type AppProvider struct {

        
            Email AppProviderRes  `json:"email"`
            Sms AppProviderRes  `json:"sms"`
            Voice AppProviderResVoice  `json:"voice"`
            ID string  `json:"_id"`
            Application string  `json:"application"`
            CreatedAt string  `json:"created_at"`
            UpdatedAt string  `json:"updated_at"`
            V float64  `json:"__v"`
         
    }
    
    // AppProviderRes used by Communication
    type AppProviderRes struct {

        
            Transaction AppProviderResObj  `json:"transaction"`
            Promotional AppProviderResObj  `json:"promotional"`
            Otp AppProviderResObj  `json:"otp"`
         
    }
    
    // AppProviderResVoice used by Communication
    type AppProviderResVoice struct {

        
            Transaction AppProviderResObj  `json:"transaction"`
            Otp AppProviderResObj  `json:"otp"`
         
    }
    
    // AppProviderResObj used by Communication
    type AppProviderResObj struct {

        
            Provider string  `json:"provider"`
         
    }
    
    // GlobalProviders used by Communication
    type GlobalProviders struct {

        
            Email []GlobalProvidersResObj  `json:"email"`
            Sms []GlobalProvidersResObj  `json:"sms"`
            Voice []GlobalProvidersResObj  `json:"voice"`
         
    }
    
    // GlobalProvidersResObj used by Communication
    type GlobalProvidersResObj struct {

        
            ID string  `json:"_id"`
            Name string  `json:"name"`
         
    }
    
    // AppProviderAdmin used by Communication
    type AppProviderAdmin struct {

        
            Email AppProviderAdminObj  `json:"email"`
            Sms AppProviderAdminObj  `json:"sms"`
            Voice AppProviderAdminObj  `json:"voice"`
         
    }
    
    // AppProviderAdminObj used by Communication
    type AppProviderAdminObj struct {

        
            Transaction AppProviderResObj  `json:"transaction"`
            Otp AppProviderResObj  `json:"otp"`
         
    }
    
    // UpdateGlobalProviders used by Communication
    type UpdateGlobalProviders struct {

        
            Email UpdateGlobalProvidersObj  `json:"email"`
            Sms UpdateGlobalProvidersObj  `json:"sms"`
            Voice UpdateGlobalProvidersObj  `json:"voice"`
            ID string  `json:"_id"`
            CreatedAt string  `json:"created_at"`
            UpdatedAt string  `json:"updated_at"`
            V float64  `json:"__v"`
         
    }
    
    // UpdateGlobalProvidersObj used by Communication
    type UpdateGlobalProvidersObj struct {

        
            DefaultProvider string  `json:"default_provider"`
            OtpProvider string  `json:"otp_provider"`
         
    }
    
    // AppProviderReq used by Communication
    type AppProviderReq struct {

        
            Email AppProviderRes  `json:"email"`
            Sms AppProviderRes  `json:"sms"`
            Voice AppProviderResVoice  `json:"voice"`
         
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
    
    // SmsProviderAdmin used by Communication
    type SmsProviderAdmin struct {

        
            Rpt float64  `json:"rpt"`
            Type string  `json:"type"`
            Provider string  `json:"provider"`
            ID string  `json:"_id"`
            Name string  `json:"name"`
            EntityID string  `json:"entity_id"`
            Description string  `json:"description"`
            Sender string  `json:"sender"`
            Username string  `json:"username"`
            Authkey string  `json:"authkey"`
            CreatedAt string  `json:"created_at"`
            UpdatedAt string  `json:"updated_at"`
            Slug string  `json:"slug"`
            V float64  `json:"__v"`
         
    }
    
    // SmsProvidersAdmin used by Communication
    type SmsProvidersAdmin struct {

        
            Items []SmsProviderAdmin  `json:"items"`
            Page Page  `json:"page"`
         
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
            EntityID string  `json:"entity_id"`
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
    
    // DefaultSmsProviders used by Communication
    type DefaultSmsProviders struct {

        
            ID string  `json:"_id"`
            Name string  `json:"name"`
            IsDefault bool  `json:"is_default"`
         
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
    
    // VoiceProviderReq used by Communication
    type VoiceProviderReq struct {

        
            Name string  `json:"name"`
            Description string  `json:"description"`
            Sender string  `json:"sender"`
            Username string  `json:"username"`
            Authkey string  `json:"authkey"`
            Type string  `json:"type"`
            Provider string  `json:"provider"`
         
    }
    
    // VoiceProvider used by Communication
    type VoiceProvider struct {

        
            ID string  `json:"_id"`
            Application string  `json:"application"`
            Type string  `json:"type"`
            Name string  `json:"name"`
            Description string  `json:"description"`
            Slug string  `json:"slug"`
            Provider string  `json:"provider"`
            CreatedAt string  `json:"created_at"`
            UpdatedAt string  `json:"updated_at"`
            V float64  `json:"__v"`
            Discriminator string  `json:"discriminator"`
            Username string  `json:"username"`
            Password string  `json:"password"`
            CallerID string  `json:"caller_id"`
            AppletURL string  `json:"applet_url"`
            WhitelistedIp []string  `json:"whitelisted_ip"`
         
    }
    
    // VoiceProviders used by Communication
    type VoiceProviders struct {

        
            Items []VoiceProvider  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // VoiceTemplateDeleteSuccessRes used by Communication
    type VoiceTemplateDeleteSuccessRes struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // VoiceTemplateDeleteFailureRes used by Communication
    type VoiceTemplateDeleteFailureRes struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // VoiceTemplateMessage used by Communication
    type VoiceTemplateMessage struct {

        
            TemplateType string  `json:"template_type"`
            Template string  `json:"template"`
         
    }
    
    // VoiceTemplateReq used by Communication
    type VoiceTemplateReq struct {

        
            Name string  `json:"name"`
            Description string  `json:"description"`
            Message VoiceTemplateMessage  `json:"message"`
            TemplateVariables interface{}  `json:"template_variables"`
            Attachments []interface{}  `json:"attachments"`
            Priority string  `json:"priority"`
         
    }
    
    // VoiceTemplateRes used by Communication
    type VoiceTemplateRes struct {

        
            IsSystem bool  `json:"is_system"`
            IsInternal bool  `json:"is_internal"`
            Description string  `json:"description"`
            Tags []interface{}  `json:"tags"`
            Priority string  `json:"priority"`
            Published bool  `json:"published"`
            ID string  `json:"_id"`
            Name string  `json:"name"`
            Message VoiceTemplateMessage  `json:"message"`
            TemplateVariables interface{}  `json:"template_variables"`
            CreatedAt string  `json:"created_at"`
            UpdatedAt string  `json:"updated_at"`
            Slug string  `json:"slug"`
            V float64  `json:"__v"`
         
    }
    
    // VoiceTemplate used by Communication
    type VoiceTemplate struct {

        
            IsSystem bool  `json:"is_system"`
            IsInternal bool  `json:"is_internal"`
            Description string  `json:"description"`
            Priority string  `json:"priority"`
            Tags []interface{}  `json:"tags"`
            Published bool  `json:"published"`
            ID string  `json:"_id"`
            Slug string  `json:"slug"`
            Name string  `json:"name"`
            Message VoiceTemplateMessage  `json:"message"`
            TemplateVariables interface{}  `json:"template_variables"`
            CreatedAt string  `json:"created_at"`
            UpdatedAt string  `json:"updated_at"`
            V float64  `json:"__v"`
         
    }
    
    // SystemVoiceTemplate used by Communication
    type SystemVoiceTemplate struct {

        
            IsSystem bool  `json:"is_system"`
            IsInternal bool  `json:"is_internal"`
            Description string  `json:"description"`
            Tags []interface{}  `json:"tags"`
            Priority string  `json:"priority"`
            Published bool  `json:"published"`
            ID string  `json:"_id"`
            Slug string  `json:"slug"`
            Name string  `json:"name"`
            Message VoiceTemplateMessage  `json:"message"`
            TemplateVariables interface{}  `json:"template_variables"`
            CreatedAt string  `json:"created_at"`
            UpdatedAt string  `json:"updated_at"`
            V float64  `json:"__v"`
         
    }
    
    // VoiceTemplates used by Communication
    type VoiceTemplates struct {

        
            Items []VoiceTemplate  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // SystemVoiceTemplates used by Communication
    type SystemVoiceTemplates struct {

        
            Items []SystemVoiceTemplate  `json:"items"`
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
    
    // GenericSuccess used by Communication
    type GenericSuccess struct {

        
            Success bool  `json:"success"`
         
    }
    
    // InvalidRangeErrorReqPositive used by Communication
    type InvalidRangeErrorReqPositive struct {

        
            Message string  `json:"message"`
            Code float64  `json:"code"`
            Sentry string  `json:"sentry"`
         
    }
    
    // InvalidInputRequiredByteOrHexError used by Communication
    type InvalidInputRequiredByteOrHexError struct {

        
            Message string  `json:"message"`
            Sentry string  `json:"sentry"`
         
    }
    
    // NameValidatorError used by Communication
    type NameValidatorError struct {

        
            Message NameValidatorErrorMessage  `json:"message"`
            Sentry string  `json:"sentry"`
         
    }
    
    // NameValidatorErrorMessage used by Communication
    type NameValidatorErrorMessage struct {

        
            Name ValidatorErrorBody  `json:"name"`
         
    }
    
    // ApikeyValidatorError used by Communication
    type ApikeyValidatorError struct {

        
            Message ApikeyValidatorErrorMessage  `json:"message"`
            Sentry string  `json:"sentry"`
         
    }
    
    // ApikeyValidatorErrorMessage used by Communication
    type ApikeyValidatorErrorMessage struct {

        
            APIKey ValidatorErrorBody  `json:"api_key"`
         
    }
    
    // FeedidValidatorError used by Communication
    type FeedidValidatorError struct {

        
            Message FeedidValidatorErrorMessage  `json:"message"`
            Sentry string  `json:"sentry"`
         
    }
    
    // FeedidValidatorErrorMessage used by Communication
    type FeedidValidatorErrorMessage struct {

        
            Feedid ValidatorErrorBody  `json:"feedid"`
         
    }
    
    // UsernameValidatorError used by Communication
    type UsernameValidatorError struct {

        
            Message UsernameValidatorErrorMessage  `json:"message"`
            Sentry string  `json:"sentry"`
         
    }
    
    // UsernameValidatorErrorMessage used by Communication
    type UsernameValidatorErrorMessage struct {

        
            Username ValidatorErrorBody  `json:"username"`
         
    }
    
    // PasswordValidatorError used by Communication
    type PasswordValidatorError struct {

        
            Message PasswordValidatorErrorMessage  `json:"message"`
            Sentry string  `json:"sentry"`
         
    }
    
    // PasswordValidatorErrorMessage used by Communication
    type PasswordValidatorErrorMessage struct {

        
            Password ValidatorErrorBody  `json:"password"`
         
    }
    
    // ValidatorErrorBody used by Communication
    type ValidatorErrorBody struct {

        
            Name string  `json:"name"`
            Message string  `json:"message"`
            Properties ValidatorErrorMessageProperties  `json:"properties"`
            Kind string  `json:"kind"`
            Path string  `json:"path"`
         
    }
    
    // ValidatorErrorMessageProperties used by Communication
    type ValidatorErrorMessageProperties struct {

        
            Message string  `json:"message"`
            Type string  `json:"type"`
            Path string  `json:"path"`
         
    }
    
    // CastToStringFail used by Communication
    type CastToStringFail struct {

        
            Message string  `json:"message"`
            Sentry string  `json:"sentry"`
         
    }
    
    // InvalidID used by Communication
    type InvalidID struct {

        
            Message string  `json:"message"`
            Sentry string  `json:"sentry"`
         
    }
    

    
    // PaymentGatewayConfigResponse used by Payment
    type PaymentGatewayConfigResponse struct {

        
            Aggregators []map[string]interface{}  `json:"aggregators"`
            Created bool  `json:"created"`
            ExcludedFields []string  `json:"excluded_fields"`
            AppID string  `json:"app_id"`
            Success bool  `json:"success"`
            DisplayFields []string  `json:"display_fields"`
         
    }
    
    // ErrorCodeDescription used by Payment
    type ErrorCodeDescription struct {

        
            Success bool  `json:"success"`
            Description string  `json:"description"`
            Code string  `json:"code"`
         
    }
    
    // PaymentGatewayConfig used by Payment
    type PaymentGatewayConfig struct {

        
            Key string  `json:"key"`
            Secret string  `json:"secret"`
            MerchantSalt string  `json:"merchant_salt"`
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

        
            Success bool  `json:"success"`
            Aggregator []string  `json:"aggregator"`
         
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
    
    // PaymentModeLogo used by Payment
    type PaymentModeLogo struct {

        
            Small string  `json:"small"`
            Large string  `json:"large"`
         
    }
    
    // IntentApp used by Payment
    type IntentApp struct {

        
            PackageName string  `json:"package_name"`
            DisplayName string  `json:"display_name"`
            Logos PaymentModeLogo  `json:"logos"`
            Code string  `json:"code"`
         
    }
    
    // IntentAppErrorList used by Payment
    type IntentAppErrorList struct {

        
            PackageName string  `json:"package_name"`
            Code string  `json:"code"`
         
    }
    
    // PaymentModeList used by Payment
    type PaymentModeList struct {

        
            CardName string  `json:"card_name"`
            ExpYear float64  `json:"exp_year"`
            Nickname string  `json:"nickname"`
            CodLimit float64  `json:"cod_limit"`
            CardToken string  `json:"card_token"`
            CardIssuer string  `json:"card_issuer"`
            Timeout float64  `json:"timeout"`
            Expired bool  `json:"expired"`
            AggregatorName string  `json:"aggregator_name"`
            Code string  `json:"code"`
            IntentAppErrorList []string  `json:"intent_app_error_list"`
            CompliantWithTokenisationGuidelines bool  `json:"compliant_with_tokenisation_guidelines"`
            CodLimitPerOrder float64  `json:"cod_limit_per_order"`
            LogoURL PaymentModeLogo  `json:"logo_url"`
            RemainingLimit float64  `json:"remaining_limit"`
            DisplayPriority float64  `json:"display_priority"`
            DisplayName string  `json:"display_name"`
            IntentApp []IntentApp  `json:"intent_app"`
            FyndVpa string  `json:"fynd_vpa"`
            IntentAppErrorDictList []IntentAppErrorList  `json:"intent_app_error_dict_list"`
            CardBrandImage string  `json:"card_brand_image"`
            CardBrand string  `json:"card_brand"`
            RetryCount float64  `json:"retry_count"`
            ExpMonth float64  `json:"exp_month"`
            CardNumber string  `json:"card_number"`
            Name string  `json:"name"`
            CardFingerprint string  `json:"card_fingerprint"`
            MerchantCode string  `json:"merchant_code"`
            CardReference string  `json:"card_reference"`
            CardIsin string  `json:"card_isin"`
            CardID string  `json:"card_id"`
            CardType string  `json:"card_type"`
            IntentFlow bool  `json:"intent_flow"`
         
    }
    
    // RootPaymentMode used by Payment
    type RootPaymentMode struct {

        
            Name string  `json:"name"`
            AddCardEnabled bool  `json:"add_card_enabled"`
            List []PaymentModeList  `json:"list"`
            IsPayByCardPl bool  `json:"is_pay_by_card_pl"`
            AnonymousEnable bool  `json:"anonymous_enable"`
            SaveCard bool  `json:"save_card"`
            DisplayPriority float64  `json:"display_priority"`
            DisplayName string  `json:"display_name"`
            AggregatorName string  `json:"aggregator_name"`
         
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
    
    // PayoutsResponse used by Payment
    type PayoutsResponse struct {

        
            IsDefault bool  `json:"is_default"`
            MoreAttributes map[string]interface{}  `json:"more_attributes"`
            TransferType string  `json:"transfer_type"`
            UniqueTransferNo map[string]interface{}  `json:"unique_transfer_no"`
            IsActive bool  `json:"is_active"`
            Customers map[string]interface{}  `json:"customers"`
            PayoutsAggregators []map[string]interface{}  `json:"payouts_aggregators"`
         
    }
    
    // PayoutBankDetails used by Payment
    type PayoutBankDetails struct {

        
            Country string  `json:"country"`
            AccountHolder string  `json:"account_holder"`
            City string  `json:"city"`
            IfscCode string  `json:"ifsc_code"`
            BankName string  `json:"bank_name"`
            Pincode float64  `json:"pincode"`
            AccountNo string  `json:"account_no"`
            State string  `json:"state"`
            BranchName string  `json:"branch_name"`
            AccountType string  `json:"account_type"`
         
    }
    
    // PayoutRequest used by Payment
    type PayoutRequest struct {

        
            Aggregator string  `json:"aggregator"`
            TransferType string  `json:"transfer_type"`
            UniqueExternalID string  `json:"unique_external_id"`
            BankDetails PayoutBankDetails  `json:"bank_details"`
            IsActive bool  `json:"is_active"`
            Users map[string]interface{}  `json:"users"`
         
    }
    
    // PayoutResponse used by Payment
    type PayoutResponse struct {

        
            Created bool  `json:"created"`
            Aggregator string  `json:"aggregator"`
            PaymentStatus string  `json:"payment_status"`
            UniqueTransferNo string  `json:"unique_transfer_no"`
            BankDetails map[string]interface{}  `json:"bank_details"`
            TransferType string  `json:"transfer_type"`
            IsActive bool  `json:"is_active"`
            Payouts map[string]interface{}  `json:"payouts"`
            Users map[string]interface{}  `json:"users"`
            Success bool  `json:"success"`
         
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
            IsDefault bool  `json:"is_default"`
            UniqueExternalID string  `json:"unique_external_id"`
         
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

        
            Success bool  `json:"success"`
            Config map[string]interface{}  `json:"config"`
            Aggregator string  `json:"aggregator"`
         
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

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
            Data map[string]interface{}  `json:"data"`
            IsVerifiedFlag bool  `json:"is_verified_flag"`
         
    }
    
    // NotFoundResourceError used by Payment
    type NotFoundResourceError struct {

        
            Success bool  `json:"success"`
            Description string  `json:"description"`
            Code string  `json:"code"`
         
    }
    
    // BankDetailsForOTP used by Payment
    type BankDetailsForOTP struct {

        
            AccountHolder string  `json:"account_holder"`
            IfscCode string  `json:"ifsc_code"`
            BankName string  `json:"bank_name"`
            AccountNo string  `json:"account_no"`
            BranchName string  `json:"branch_name"`
         
    }
    
    // AddBeneficiaryDetailsOTPRequest used by Payment
    type AddBeneficiaryDetailsOTPRequest struct {

        
            OrderID string  `json:"order_id"`
            Details BankDetailsForOTP  `json:"details"`
         
    }
    
    // IfscCodeResponse used by Payment
    type IfscCodeResponse struct {

        
            Success bool  `json:"success"`
            BankName string  `json:"bank_name"`
            BranchName string  `json:"branch_name"`
         
    }
    
    // OrderBeneficiaryDetails used by Payment
    type OrderBeneficiaryDetails struct {

        
            Email string  `json:"email"`
            TransferMode string  `json:"transfer_mode"`
            BankName string  `json:"bank_name"`
            CreatedOn string  `json:"created_on"`
            Title string  `json:"title"`
            IsActive bool  `json:"is_active"`
            ModifiedOn string  `json:"modified_on"`
            Subtitle string  `json:"subtitle"`
            AccountHolder string  `json:"account_holder"`
            BeneficiaryID string  `json:"beneficiary_id"`
            IfscCode string  `json:"ifsc_code"`
            Mobile string  `json:"mobile"`
            Comment string  `json:"comment"`
            DelightsUserName string  `json:"delights_user_name"`
            DisplayName string  `json:"display_name"`
            BranchName string  `json:"branch_name"`
            Address string  `json:"address"`
            AccountNo string  `json:"account_no"`
            ID float64  `json:"id"`
         
    }
    
    // OrderBeneficiaryResponse used by Payment
    type OrderBeneficiaryResponse struct {

        
            Beneficiaries []OrderBeneficiaryDetails  `json:"beneficiaries"`
            ShowBeneficiaryDetails bool  `json:"show_beneficiary_details"`
         
    }
    
    // MultiTenderPaymentMeta used by Payment
    type MultiTenderPaymentMeta struct {

        
            PaymentGateway string  `json:"payment_gateway"`
            PaymentID string  `json:"payment_id"`
            OrderID string  `json:"order_id"`
            CurrentStatus string  `json:"current_status"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
         
    }
    
    // MultiTenderPaymentMethod used by Payment
    type MultiTenderPaymentMethod struct {

        
            Amount float64  `json:"amount"`
            Name string  `json:"name"`
            Meta MultiTenderPaymentMeta  `json:"meta"`
            Mode string  `json:"mode"`
         
    }
    
    // PaymentConfirmationRequest used by Payment
    type PaymentConfirmationRequest struct {

        
            PaymentMethods []MultiTenderPaymentMethod  `json:"payment_methods"`
            OrderID string  `json:"order_id"`
         
    }
    
    // PaymentConfirmationResponse used by Payment
    type PaymentConfirmationResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
            OrderID string  `json:"order_id"`
         
    }
    
    // CODdata used by Payment
    type CODdata struct {

        
            Limit float64  `json:"limit"`
            UserID string  `json:"user_id"`
            RemainingLimit float64  `json:"remaining_limit"`
            IsActive bool  `json:"is_active"`
            Usages float64  `json:"usages"`
         
    }
    
    // GetUserCODLimitResponse used by Payment
    type GetUserCODLimitResponse struct {

        
            UserCodData CODdata  `json:"user_cod_data"`
            Success bool  `json:"success"`
         
    }
    
    // SetCODForUserRequest used by Payment
    type SetCODForUserRequest struct {

        
            IsActive bool  `json:"is_active"`
            Mobileno string  `json:"mobileno"`
            MerchantUserID string  `json:"merchant_user_id"`
         
    }
    
    // SetCODOptionResponse used by Payment
    type SetCODOptionResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // EdcModelData used by Payment
    type EdcModelData struct {

        
            Models []string  `json:"models"`
            Aggregator string  `json:"aggregator"`
            AggregatorID float64  `json:"aggregator_id"`
         
    }
    
    // EdcAggregatorAndModelListResponse used by Payment
    type EdcAggregatorAndModelListResponse struct {

        
            Success bool  `json:"success"`
            Data []EdcModelData  `json:"data"`
         
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

        
            EdcDeviceSerialNo string  `json:"edc_device_serial_no"`
            EdcModel string  `json:"edc_model"`
            DeviceTag string  `json:"device_tag"`
            AggregatorID float64  `json:"aggregator_id"`
            StoreID float64  `json:"store_id"`
            TerminalSerialNo string  `json:"terminal_serial_no"`
         
    }
    
    // EdcDevice used by Payment
    type EdcDevice struct {

        
            EdcDeviceSerialNo string  `json:"edc_device_serial_no"`
            EdcModel string  `json:"edc_model"`
            DeviceTag string  `json:"device_tag"`
            AggregatorID float64  `json:"aggregator_id"`
            StoreID float64  `json:"store_id"`
            IsActive bool  `json:"is_active"`
            ApplicationID string  `json:"application_id"`
            MerchantStorePosCode string  `json:"merchant_store_pos_code"`
            TerminalSerialNo string  `json:"terminal_serial_no"`
            TerminalUniqueIdentifier string  `json:"terminal_unique_identifier"`
            AggregatorName string  `json:"aggregator_name"`
         
    }
    
    // EdcDeviceAddResponse used by Payment
    type EdcDeviceAddResponse struct {

        
            Success bool  `json:"success"`
            Data EdcDevice  `json:"data"`
         
    }
    
    // EdcDeviceDetailsResponse used by Payment
    type EdcDeviceDetailsResponse struct {

        
            Success bool  `json:"success"`
            Data EdcDevice  `json:"data"`
         
    }
    
    // EdcUpdateRequest used by Payment
    type EdcUpdateRequest struct {

        
            EdcDeviceSerialNo string  `json:"edc_device_serial_no"`
            EdcModel string  `json:"edc_model"`
            DeviceTag string  `json:"device_tag"`
            AggregatorID float64  `json:"aggregator_id"`
            StoreID float64  `json:"store_id"`
            IsActive bool  `json:"is_active"`
            MerchantStorePosCode string  `json:"merchant_store_pos_code"`
         
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

        
            Success bool  `json:"success"`
            Page Page  `json:"page"`
            Items []EdcDevice  `json:"items"`
         
    }
    
    // PaymentInitializationRequest used by Payment
    type PaymentInitializationRequest struct {

        
            DeviceID string  `json:"device_id"`
            MerchantOrderID string  `json:"merchant_order_id"`
            Amount float64  `json:"amount"`
            Email string  `json:"email"`
            CustomerID string  `json:"customer_id"`
            Aggregator string  `json:"aggregator"`
            RazorpayPaymentID string  `json:"razorpay_payment_id"`
            OrderID string  `json:"order_id"`
            Contact string  `json:"contact"`
            Timeout float64  `json:"timeout"`
            Vpa string  `json:"vpa"`
            Method string  `json:"method"`
            Currency string  `json:"currency"`
         
    }
    
    // PaymentInitializationResponse used by Payment
    type PaymentInitializationResponse struct {

        
            Status string  `json:"status"`
            DeviceID string  `json:"device_id"`
            MerchantOrderID string  `json:"merchant_order_id"`
            Amount float64  `json:"amount"`
            CustomerID string  `json:"customer_id"`
            Vpa string  `json:"vpa"`
            Aggregator string  `json:"aggregator"`
            RazorpayPaymentID string  `json:"razorpay_payment_id"`
            PollingURL string  `json:"polling_url"`
            VirtualID string  `json:"virtual_id"`
            AggregatorOrderID string  `json:"aggregator_order_id"`
            UpiPollURL string  `json:"upi_poll_url"`
            Timeout float64  `json:"timeout"`
            BqrImage string  `json:"bqr_image"`
            Success bool  `json:"success"`
            Method string  `json:"method"`
            Currency string  `json:"currency"`
         
    }
    
    // PaymentStatusUpdateRequest used by Payment
    type PaymentStatusUpdateRequest struct {

        
            Status string  `json:"status"`
            MerchantOrderID string  `json:"merchant_order_id"`
            Amount float64  `json:"amount"`
            Email string  `json:"email"`
            CustomerID string  `json:"customer_id"`
            Aggregator string  `json:"aggregator"`
            OrderID string  `json:"order_id"`
            Contact string  `json:"contact"`
            DeviceID string  `json:"device_id"`
            Vpa string  `json:"vpa"`
            MerchantTransactionID string  `json:"merchant_transaction_id"`
            Method string  `json:"method"`
            Currency string  `json:"currency"`
         
    }
    
    // PaymentStatusUpdateResponse used by Payment
    type PaymentStatusUpdateResponse struct {

        
            Status string  `json:"status"`
            Retry bool  `json:"retry"`
            RedirectURL string  `json:"redirect_url"`
            Success bool  `json:"success"`
            AggregatorName string  `json:"aggregator_name"`
         
    }
    
    // ResendOrCancelPaymentRequest used by Payment
    type ResendOrCancelPaymentRequest struct {

        
            RequestType string  `json:"request_type"`
            OrderID string  `json:"order_id"`
            DeviceID string  `json:"device_id"`
         
    }
    
    // LinkStatus used by Payment
    type LinkStatus struct {

        
            Status bool  `json:"status"`
            Message string  `json:"message"`
         
    }
    
    // ResendOrCancelPaymentResponse used by Payment
    type ResendOrCancelPaymentResponse struct {

        
            Success bool  `json:"success"`
            Data LinkStatus  `json:"data"`
         
    }
    
    // PaymentStatusBulkHandlerRequest used by Payment
    type PaymentStatusBulkHandlerRequest struct {

        
            MerchantOrderID []string  `json:"merchant_order_id"`
         
    }
    
    // PaymentObjectListSerializer used by Payment
    type PaymentObjectListSerializer struct {

        
            AllStatus []string  `json:"all_status"`
            PaymentModeIdentifier string  `json:"payment_mode_identifier"`
            CollectedBy string  `json:"collected_by"`
            AmountInPaisa string  `json:"amount_in_paisa"`
            PaymentMode string  `json:"payment_mode"`
            PaymentID string  `json:"payment_id"`
            ID string  `json:"id"`
            CreatedOn string  `json:"created_on"`
            CurrentStatus string  `json:"current_status"`
            AggregatorPaymentObject map[string]interface{}  `json:"aggregator_payment_object"`
            ApplicationID string  `json:"application_id"`
            RefundedBy string  `json:"refunded_by"`
            CompanyID string  `json:"company_id"`
            ModifiedOn string  `json:"modified_on"`
            UserObject map[string]interface{}  `json:"user_object"`
            RefundObject map[string]interface{}  `json:"refund_object"`
            PaymentGateway string  `json:"payment_gateway"`
            Currency string  `json:"currency"`
         
    }
    
    // PaymentStatusObject used by Payment
    type PaymentStatusObject struct {

        
            MerchantOrderID string  `json:"merchant_order_id"`
            PaymentObjectList []PaymentObjectListSerializer  `json:"payment_object_list"`
         
    }
    
    // PaymentStatusBulkHandlerResponse used by Payment
    type PaymentStatusBulkHandlerResponse struct {

        
            Status float64  `json:"status"`
            Data []PaymentStatusObject  `json:"data"`
            Error string  `json:"error"`
            Count float64  `json:"count"`
            Success string  `json:"success"`
         
    }
    
    // GetOauthUrlResponse used by Payment
    type GetOauthUrlResponse struct {

        
            Success bool  `json:"success"`
            URL string  `json:"url"`
         
    }
    
    // RevokeOAuthToken used by Payment
    type RevokeOAuthToken struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // RepaymentRequestDetails used by Payment
    type RepaymentRequestDetails struct {

        
            PaymentModeIdentifier string  `json:"payment_mode_identifier"`
            MerchantOrderID string  `json:"merchant_order_id"`
            Amount float64  `json:"amount"`
            FwdShipmentID string  `json:"fwd_shipment_id"`
            PaymentMode string  `json:"payment_mode"`
            Aggregator string  `json:"aggregator"`
            CurrentStatus string  `json:"current_status"`
            OutstandingDetailsID float64  `json:"outstanding_details_id"`
            AggregatorTransactionID string  `json:"aggregator_transaction_id"`
            AggregatorOrderID string  `json:"aggregator_order_id"`
         
    }
    
    // RepaymentDetailsSerialiserPayAll used by Payment
    type RepaymentDetailsSerialiserPayAll struct {

        
            ExtensionOrderID string  `json:"extension_order_id"`
            TotalAmount float64  `json:"total_amount"`
            AggregatorTransactionID string  `json:"aggregator_transaction_id"`
            AggregatorOrderID string  `json:"aggregator_order_id"`
            ShipmentDetails []RepaymentRequestDetails  `json:"shipment_details"`
         
    }
    
    // RepaymentResponse used by Payment
    type RepaymentResponse struct {

        
            Success bool  `json:"success"`
            Data map[string]interface{}  `json:"data"`
         
    }
    
    // MerchantOnBoardingRequest used by Payment
    type MerchantOnBoardingRequest struct {

        
            Status string  `json:"status"`
            UserID string  `json:"user_id"`
            Aggregator string  `json:"aggregator"`
            CreditLineID string  `json:"credit_line_id"`
            AppID string  `json:"app_id"`
         
    }
    
    // MerchantOnBoardingResponse used by Payment
    type MerchantOnBoardingResponse struct {

        
            Success bool  `json:"success"`
            Data map[string]interface{}  `json:"data"`
         
    }
    
    // ValidateCustomerRequest used by Payment
    type ValidateCustomerRequest struct {

        
            BillingAddress map[string]interface{}  `json:"billing_address"`
            TransactionAmountInPaise float64  `json:"transaction_amount_in_paise"`
            Aggregator string  `json:"aggregator"`
            PhoneNumber string  `json:"phone_number"`
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
            Payload string  `json:"payload"`
            MerchantParams map[string]interface{}  `json:"merchant_params"`
            OrderItems []map[string]interface{}  `json:"order_items"`
         
    }
    
    // ValidateCustomerResponse used by Payment
    type ValidateCustomerResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
            Data map[string]interface{}  `json:"data"`
         
    }
    
    // GetPaymentLinkResponse used by Payment
    type GetPaymentLinkResponse struct {

        
            PollingTimeout float64  `json:"polling_timeout"`
            Message string  `json:"message"`
            Amount float64  `json:"amount"`
            StatusCode float64  `json:"status_code"`
            MerchantName string  `json:"merchant_name"`
            ExternalOrderID string  `json:"external_order_id"`
            PaymentLinkURL string  `json:"payment_link_url"`
            Success bool  `json:"success"`
            PaymentLinkCurrentStatus string  `json:"payment_link_current_status"`
         
    }
    
    // ErrorDescription used by Payment
    type ErrorDescription struct {

        
            MerchantOrderID string  `json:"merchant_order_id"`
            Amount float64  `json:"amount"`
            MerchantName string  `json:"merchant_name"`
            PaymentTransactionID string  `json:"payment_transaction_id"`
            InvalidID bool  `json:"invalid_id"`
            Cancelled bool  `json:"cancelled"`
            Expired bool  `json:"expired"`
            Msg string  `json:"msg"`
         
    }
    
    // ErrorResponse used by Payment
    type ErrorResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
            Error ErrorDescription  `json:"error"`
            StatusCode float64  `json:"status_code"`
         
    }
    
    // CreatePaymentLinkMeta used by Payment
    type CreatePaymentLinkMeta struct {

        
            AssignCardID string  `json:"assign_card_id"`
            Amount string  `json:"amount"`
            Pincode string  `json:"pincode"`
            CheckoutMode string  `json:"checkout_mode"`
            CartID string  `json:"cart_id"`
         
    }
    
    // CreatePaymentLinkRequest used by Payment
    type CreatePaymentLinkRequest struct {

        
            Amount float64  `json:"amount"`
            Email string  `json:"email"`
            Meta CreatePaymentLinkMeta  `json:"meta"`
            ExternalOrderID string  `json:"external_order_id"`
            MobileNumber string  `json:"mobile_number"`
            Description string  `json:"description"`
         
    }
    
    // CreatePaymentLinkResponse used by Payment
    type CreatePaymentLinkResponse struct {

        
            PollingTimeout float64  `json:"polling_timeout"`
            PaymentLinkID string  `json:"payment_link_id"`
            Message string  `json:"message"`
            StatusCode float64  `json:"status_code"`
            PaymentLinkURL string  `json:"payment_link_url"`
            Success bool  `json:"success"`
         
    }
    
    // PollingPaymentLinkResponse used by Payment
    type PollingPaymentLinkResponse struct {

        
            Status string  `json:"status"`
            PaymentLinkID string  `json:"payment_link_id"`
            Message string  `json:"message"`
            Amount float64  `json:"amount"`
            StatusCode float64  `json:"status_code"`
            OrderID string  `json:"order_id"`
            HttpStatus float64  `json:"http_status"`
            RedirectURL string  `json:"redirect_url"`
            Success bool  `json:"success"`
            AggregatorName string  `json:"aggregator_name"`
         
    }
    
    // CancelOrResendPaymentLinkRequest used by Payment
    type CancelOrResendPaymentLinkRequest struct {

        
            PaymentLinkID string  `json:"payment_link_id"`
         
    }
    
    // ResendPaymentLinkResponse used by Payment
    type ResendPaymentLinkResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
            PollingTimeout float64  `json:"polling_timeout"`
            StatusCode float64  `json:"status_code"`
         
    }
    
    // CancelPaymentLinkResponse used by Payment
    type CancelPaymentLinkResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
            StatusCode float64  `json:"status_code"`
         
    }
    
    // ExtensionPaymentUpdateRequestSerializer used by Payment
    type ExtensionPaymentUpdateRequestSerializer struct {

        
            Status string  `json:"status"`
            Gid string  `json:"gid"`
            OrderDetails map[string]interface{}  `json:"order_details"`
            TotalAmount float64  `json:"total_amount"`
            Currency string  `json:"currency"`
            PaymentDetails map[string]interface{}  `json:"payment_details"`
         
    }
    
    // ExtensionPaymentUpdateResponseSerializer used by Payment
    type ExtensionPaymentUpdateResponseSerializer struct {

        
            Status string  `json:"status"`
            Gid string  `json:"gid"`
            TotalAmount float64  `json:"total_amount"`
            PlatformTransactionDetails map[string]interface{}  `json:"platform_transaction_details"`
            Currency string  `json:"currency"`
         
    }
    
    // Code used by Payment
    type Code struct {

        
            MerchantCode string  `json:"merchant_code"`
            Name string  `json:"name"`
            Code string  `json:"code"`
         
    }
    
    // PaymentCode used by Payment
    type PaymentCode struct {

        
            Name string  `json:"name"`
            Codes Code  `json:"codes"`
            Networks string  `json:"networks"`
            Types string  `json:"types"`
         
    }
    
    // GetPaymentCode used by Payment
    type GetPaymentCode struct {

        
            MethodCode PaymentCode  `json:"method_code"`
         
    }
    
    // GetPaymentCodeResponse used by Payment
    type GetPaymentCodeResponse struct {

        
            Success bool  `json:"success"`
            Data GetPaymentCode  `json:"data"`
         
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
    
    // PlatformDeliveryAddress used by Order
    type PlatformDeliveryAddress struct {

        
            Address2 string  `json:"address2"`
            AddressType string  `json:"address_type"`
            UpdatedAt string  `json:"updated_at"`
            Area string  `json:"area"`
            State string  `json:"state"`
            Longitude float64  `json:"longitude"`
            Phone string  `json:"phone"`
            Version string  `json:"version"`
            CreatedAt string  `json:"created_at"`
            Latitude float64  `json:"latitude"`
            Pincode string  `json:"pincode"`
            ContactPerson string  `json:"contact_person"`
            Landmark string  `json:"landmark"`
            Country string  `json:"country"`
            Address1 string  `json:"address1"`
            Email string  `json:"email"`
            City string  `json:"city"`
            AddressCategory string  `json:"address_category"`
         
    }
    
    // UserDataInfo used by Order
    type UserDataInfo struct {

        
            AvisUserID string  `json:"avis_user_id"`
            Name string  `json:"name"`
            UID float64  `json:"uid"`
            FirstName string  `json:"first_name"`
            Mobile string  `json:"mobile"`
            Gender string  `json:"gender"`
            LastName string  `json:"last_name"`
            Email string  `json:"email"`
            IsAnonymousUser bool  `json:"is_anonymous_user"`
         
    }
    
    // ShipmentStatus used by Order
    type ShipmentStatus struct {

        
            CurrentShipmentStatus string  `json:"current_shipment_status"`
            ShipmentStatusID float64  `json:"shipment_status_id"`
            StatusCreatedAt string  `json:"status_created_at"`
            BagList []string  `json:"bag_list"`
            CreatedAt string  `json:"created_at"`
            ShipmentID string  `json:"shipment_id"`
            Meta map[string]interface{}  `json:"meta"`
            Status string  `json:"status"`
            Title string  `json:"title"`
         
    }
    
    // ShipmentListingChannel used by Order
    type ShipmentListingChannel struct {

        
            IsAffiliate bool  `json:"is_affiliate"`
            ChannelShipmentID string  `json:"channel_shipment_id"`
            Name string  `json:"name"`
            Logo string  `json:"logo"`
         
    }
    
    // Prices used by Order
    type Prices struct {

        
            AmountPaidRoundoff float64  `json:"amount_paid_roundoff"`
            FyndCredits float64  `json:"fynd_credits"`
            PriceMarked float64  `json:"price_marked"`
            PmPriceSplit float64  `json:"pm_price_split"`
            Discount float64  `json:"discount"`
            ValueOfGood float64  `json:"value_of_good"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            RefundCredit float64  `json:"refund_credit"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            Cashback float64  `json:"cashback"`
            DeliveryCharge float64  `json:"delivery_charge"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            PriceEffective float64  `json:"price_effective"`
            GiftPrice float64  `json:"gift_price"`
            CouponValue float64  `json:"coupon_value"`
            TransferPrice float64  `json:"transfer_price"`
            RefundAmount float64  `json:"refund_amount"`
            AmountPaid float64  `json:"amount_paid"`
            PromotionEffectiveDiscount float64  `json:"promotion_effective_discount"`
            CodCharges float64  `json:"cod_charges"`
            CashbackApplied float64  `json:"cashback_applied"`
         
    }
    
    // BagStateMapper used by Order
    type BagStateMapper struct {

        
            DisplayName string  `json:"display_name"`
            IsActive bool  `json:"is_active"`
            StateType string  `json:"state_type"`
            Name string  `json:"name"`
            ID float64  `json:"id"`
            AppFacing bool  `json:"app_facing"`
            AppStateName string  `json:"app_state_name"`
            NotifyCustomer bool  `json:"notify_customer"`
            AppDisplayName string  `json:"app_display_name"`
            JourneyType string  `json:"journey_type"`
         
    }
    
    // BagStatusHistory used by Order
    type BagStatusHistory struct {

        
            Reasons []map[string]interface{}  `json:"reasons"`
            DisplayName string  `json:"display_name"`
            BshID float64  `json:"bsh_id"`
            Forward bool  `json:"forward"`
            UpdatedAt string  `json:"updated_at"`
            KafkaSync bool  `json:"kafka_sync"`
            StoreID float64  `json:"store_id"`
            DeliveryPartnerID float64  `json:"delivery_partner_id"`
            BagID float64  `json:"bag_id"`
            StateType string  `json:"state_type"`
            CreatedAt string  `json:"created_at"`
            ShipmentID string  `json:"shipment_id"`
            AppDisplayName string  `json:"app_display_name"`
            BagStateMapper BagStateMapper  `json:"bag_state_mapper"`
            Status string  `json:"status"`
            DeliveryAwbNumber string  `json:"delivery_awb_number"`
            StateID float64  `json:"state_id"`
         
    }
    
    // Dates used by Order
    type Dates struct {

        
            DeliveryDate string  `json:"delivery_date"`
            OrderCreated string  `json:"order_created"`
         
    }
    
    // ShipmentListingBrand used by Order
    type ShipmentListingBrand struct {

        
            Name string  `json:"name"`
            Logo string  `json:"logo"`
            LogoBase64 string  `json:"logo_base64"`
            CreatedOn string  `json:"created_on"`
         
    }
    
    // ReplacementDetails used by Order
    type ReplacementDetails struct {

        
            OriginalAffiliateOrderID string  `json:"original_affiliate_order_id"`
            ReplacementType string  `json:"replacement_type"`
         
    }
    
    // AffiliateMeta used by Order
    type AffiliateMeta struct {

        
            ChannelShipmentID string  `json:"channel_shipment_id"`
            CouponCode string  `json:"coupon_code"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            EmployeeDiscount float64  `json:"employee_discount"`
            OrderItemID string  `json:"order_item_id"`
            MarketplaceInvoiceID string  `json:"marketplace_invoice_id"`
            BoxType string  `json:"box_type"`
            ChannelOrderID string  `json:"channel_order_id"`
            Quantity float64  `json:"quantity"`
            ReplacementDetails ReplacementDetails  `json:"replacement_details"`
            IsPriority bool  `json:"is_priority"`
            SizeLevelTotalQty float64  `json:"size_level_total_qty"`
            DueDate string  `json:"due_date"`
         
    }
    
    // AffiliateBagDetails used by Order
    type AffiliateBagDetails struct {

        
            AffiliateOrderID string  `json:"affiliate_order_id"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            EmployeeDiscount float64  `json:"employee_discount"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
            AffiliateMeta AffiliateMeta  `json:"affiliate_meta"`
         
    }
    
    // PlatformArticleAttributes used by Order
    type PlatformArticleAttributes struct {

        
            Currency string  `json:"currency"`
         
    }
    
    // PlatformItem used by Order
    type PlatformItem struct {

        
            L1Category []string  `json:"l1_category"`
            BranchURL string  `json:"branch_url"`
            Name string  `json:"name"`
            BrandID float64  `json:"brand_id"`
            L3CategoryName string  `json:"l3_category_name"`
            Attributes PlatformArticleAttributes  `json:"attributes"`
            CanCancel bool  `json:"can_cancel"`
            Brand string  `json:"brand"`
            Code string  `json:"code"`
            DepartmentID float64  `json:"department_id"`
            Images []string  `json:"images"`
            L2Category []string  `json:"l2_category"`
            ID float64  `json:"id"`
            Size string  `json:"size"`
            SlugKey string  `json:"slug_key"`
            L3Category float64  `json:"l3_category"`
            Image []string  `json:"image"`
            Color string  `json:"color"`
            CanReturn bool  `json:"can_return"`
            Meta map[string]interface{}  `json:"meta"`
            LastUpdatedAt string  `json:"last_updated_at"`
         
    }
    
    // BagReturnableCancelableStatus used by Order
    type BagReturnableCancelableStatus struct {

        
            IsActive bool  `json:"is_active"`
            CanBeCancelled bool  `json:"can_be_cancelled"`
            EnableTracking bool  `json:"enable_tracking"`
            IsCustomerReturnAllowed bool  `json:"is_customer_return_allowed"`
            IsReturnable bool  `json:"is_returnable"`
         
    }
    
    // Dimensions used by Order
    type Dimensions struct {

        
            Unit string  `json:"unit"`
            Height float64  `json:"height"`
            Width float64  `json:"width"`
            Length float64  `json:"length"`
            IsDefault bool  `json:"is_default"`
         
    }
    
    // ReturnConfig used by Order
    type ReturnConfig struct {

        
            Time float64  `json:"time"`
            Unit string  `json:"unit"`
            Returnable bool  `json:"returnable"`
         
    }
    
    // Weight used by Order
    type Weight struct {

        
            Shipping float64  `json:"shipping"`
            Unit string  `json:"unit"`
            IsDefault bool  `json:"is_default"`
         
    }
    
    // Article used by Order
    type Article struct {

        
            ChildDetails map[string]interface{}  `json:"child_details"`
            Dimensions Dimensions  `json:"dimensions"`
            EspModified bool  `json:"esp_modified"`
            SellerIdentifier string  `json:"seller_identifier"`
            Currency map[string]interface{}  `json:"currency"`
            UID string  `json:"uid"`
            ASet map[string]interface{}  `json:"a_set"`
            Identifiers map[string]interface{}  `json:"identifiers"`
            Code string  `json:"code"`
            ReturnConfig ReturnConfig  `json:"return_config"`
            ID string  `json:"_id"`
            Size string  `json:"size"`
            Weight Weight  `json:"weight"`
            RawMeta string  `json:"raw_meta"`
            IsSet bool  `json:"is_set"`
         
    }
    
    // GSTDetailsData used by Order
    type GSTDetailsData struct {

        
            GstTag string  `json:"gst_tag"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            IgstGstFee string  `json:"igst_gst_fee"`
            HsnCodeID string  `json:"hsn_code_id"`
            IgstTaxPercentage float64  `json:"igst_tax_percentage"`
            CgstTaxPercentage float64  `json:"cgst_tax_percentage"`
            HsnCode string  `json:"hsn_code"`
            GstinCode string  `json:"gstin_code"`
            GstFee float64  `json:"gst_fee"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            ValueOfGood float64  `json:"value_of_good"`
            CgstGstFee string  `json:"cgst_gst_fee"`
            SgstGstFee string  `json:"sgst_gst_fee"`
            IsDefaultHsnCode bool  `json:"is_default_hsn_code"`
            SgstTaxPercentage float64  `json:"sgst_tax_percentage"`
         
    }
    
    // Identifier used by Order
    type Identifier struct {

        
            Isbn string  `json:"isbn"`
            Alu string  `json:"alu"`
            Upc string  `json:"upc"`
            Ean string  `json:"ean"`
            SkuCode string  `json:"sku_code"`
         
    }
    
    // FinancialBreakup used by Order
    type FinancialBreakup struct {

        
            AmountPaidRoundoff float64  `json:"amount_paid_roundoff"`
            HsnCode string  `json:"hsn_code"`
            FyndCredits float64  `json:"fynd_credits"`
            TotalUnits float64  `json:"total_units"`
            PriceMarked float64  `json:"price_marked"`
            AddedToFyndCash bool  `json:"added_to_fynd_cash"`
            Identifiers Identifier  `json:"identifiers"`
            Discount float64  `json:"discount"`
            ValueOfGood float64  `json:"value_of_good"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            RefundCredit float64  `json:"refund_credit"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            Cashback float64  `json:"cashback"`
            GstTag string  `json:"gst_tag"`
            DeliveryCharge float64  `json:"delivery_charge"`
            Size string  `json:"size"`
            GstFee float64  `json:"gst_fee"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            PriceEffective float64  `json:"price_effective"`
            CouponValue float64  `json:"coupon_value"`
            TransferPrice float64  `json:"transfer_price"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            AmountPaid float64  `json:"amount_paid"`
            PromotionEffectiveDiscount float64  `json:"promotion_effective_discount"`
            ItemName string  `json:"item_name"`
            CodCharges float64  `json:"cod_charges"`
            CashbackApplied float64  `json:"cashback_applied"`
         
    }
    
    // BagUnit used by Order
    type BagUnit struct {

        
            BagStatus []BagStatusHistory  `json:"bag_status"`
            Dates Dates  `json:"dates"`
            CurrentStatus BagStatusHistory  `json:"current_status"`
            CanCancel bool  `json:"can_cancel"`
            Brand ShipmentListingBrand  `json:"brand"`
            AffiliateBagDetails AffiliateBagDetails  `json:"affiliate_bag_details"`
            Reasons []map[string]interface{}  `json:"reasons"`
            Item PlatformItem  `json:"item"`
            Size string  `json:"size"`
            CurrentOperationalStatus BagStatusHistory  `json:"current_operational_status"`
            Status BagReturnableCancelableStatus  `json:"status"`
            Article Article  `json:"article"`
            Prices Prices  `json:"prices"`
            Gst GSTDetailsData  `json:"gst"`
            DisplayName string  `json:"display_name"`
            BagID float64  `json:"bag_id"`
            CanReturn bool  `json:"can_return"`
            EntityType string  `json:"entity_type"`
            BagExpiryDate string  `json:"bag_expiry_date"`
            BagType string  `json:"bag_type"`
            ProductQuantity float64  `json:"product_quantity"`
            FinancialBreakup []FinancialBreakup  `json:"financial_breakup"`
            Meta map[string]interface{}  `json:"meta"`
            LineNumber float64  `json:"line_number"`
         
    }
    
    // ShipmentItemFulFillingStore used by Order
    type ShipmentItemFulFillingStore struct {

        
            BrandStoreTags string  `json:"brand_store_tags"`
            LocationType string  `json:"location_type"`
            State string  `json:"state"`
            StoreEmail string  `json:"store_email"`
            Name string  `json:"name"`
            Phone string  `json:"phone"`
            ID float64  `json:"id"`
            Address string  `json:"address"`
            Code string  `json:"code"`
            Meta map[string]interface{}  `json:"meta"`
            Pincode string  `json:"pincode"`
            City string  `json:"city"`
         
    }
    
    // LockData used by Order
    type LockData struct {

        
            Locked bool  `json:"locked"`
            Mto bool  `json:"mto"`
            LockMessage string  `json:"lock_message"`
         
    }
    
    // ShipmentTags used by Order
    type ShipmentTags struct {

        
            DisplayText string  `json:"display_text"`
            Slug string  `json:"slug"`
            EntityType string  `json:"entity_type"`
         
    }
    
    // Formatted used by Order
    type Formatted struct {

        
            Min string  `json:"min"`
            Max string  `json:"max"`
         
    }
    
    // ShipmentTimeStamp used by Order
    type ShipmentTimeStamp struct {

        
            Min float64  `json:"min"`
            Max float64  `json:"max"`
         
    }
    
    // ShipmentItemMeta used by Order
    type ShipmentItemMeta struct {

        
            LockData LockData  `json:"lock_data"`
            External map[string]interface{}  `json:"external"`
            PackagingName string  `json:"packaging_name"`
            ShippingZone string  `json:"shipping_zone"`
            Sla float64  `json:"sla"`
            AutoTriggerDpAssignmentAcf bool  `json:"auto_trigger_dp_assignment_acf"`
            SameStoreAvailable bool  `json:"same_store_available"`
            BagWeight map[string]interface{}  `json:"bag_weight"`
            StoreInvoiceUpdatedDate string  `json:"store_invoice_updated_date"`
            DpSortKey string  `json:"dp_sort_key"`
            ShipmentVolumetricWeight float64  `json:"shipment_volumetric_weight"`
            EwaybillInfo map[string]interface{}  `json:"ewaybill_info"`
            ShipmentChargeableWeight float64  `json:"shipment_chargeable_weight"`
            DebugInfo map[string]interface{}  `json:"debug_info"`
            ParentDpID string  `json:"parent_dp_id"`
            AssignDpFromSb bool  `json:"assign_dp_from_sb"`
            Weight float64  `json:"weight"`
            ShipmentWeight float64  `json:"shipment_weight"`
            FulfilmentPriorityText string  `json:"fulfilment_priority_text"`
            IsInternational bool  `json:"is_international"`
            DpOptions map[string]interface{}  `json:"dp_options"`
            OrderType string  `json:"order_type"`
            ExistingDpList []string  `json:"existing_dp_list"`
            Tags []map[string]interface{}  `json:"tags"`
            ShipmentTags []ShipmentTags  `json:"shipment_tags"`
            Formatted Formatted  `json:"formatted"`
            ActivityComment string  `json:"activity_comment"`
            PdfMedia []map[string]interface{}  `json:"pdf_media"`
            Timestamp ShipmentTimeStamp  `json:"timestamp"`
         
    }
    
    // ShipmentItem used by Order
    type ShipmentItem struct {

        
            OrderDate string  `json:"order_date"`
            EstimatedSlaTime string  `json:"estimated_sla_time"`
            DeliveryAddress PlatformDeliveryAddress  `json:"delivery_address"`
            StatusCreatedAt string  `json:"status_created_at"`
            LockStatus bool  `json:"lock_status"`
            ShipmentCreatedAt string  `json:"shipment_created_at"`
            User UserDataInfo  `json:"user"`
            ShipmentStatus ShipmentStatus  `json:"shipment_status"`
            OrderID string  `json:"order_id"`
            Channel ShipmentListingChannel  `json:"channel"`
            TotalBags float64  `json:"total_bags"`
            Prices Prices  `json:"prices"`
            InvoiceID string  `json:"invoice_id"`
            Bags []BagUnit  `json:"bags"`
            PreviousShipmentID string  `json:"previous_shipment_id"`
            DisplayName string  `json:"display_name"`
            PaymentMode string  `json:"payment_mode"`
            FulfillingStore ShipmentItemFulFillingStore  `json:"fulfilling_store"`
            PaymentMethods map[string]interface{}  `json:"payment_methods"`
            CustomerNote string  `json:"customer_note"`
            CanProcess bool  `json:"can_process"`
            Meta ShipmentItemMeta  `json:"meta"`
            OrderingChannnel string  `json:"ordering_channnel"`
            ShipmentID string  `json:"shipment_id"`
         
    }
    
    // ShipmentInternalPlatformViewResponse used by Order
    type ShipmentInternalPlatformViewResponse struct {

        
            Success bool  `json:"success"`
            Page Page  `json:"page"`
            TotalCount float64  `json:"total_count"`
            Message string  `json:"message"`
            Items []ShipmentItem  `json:"items"`
            Lane string  `json:"lane"`
         
    }
    
    // Error used by Order
    type Error struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // TrackingList used by Order
    type TrackingList struct {

        
            IsCurrent bool  `json:"is_current"`
            Time string  `json:"time"`
            IsPassed bool  `json:"is_passed"`
            Text string  `json:"text"`
            Status string  `json:"status"`
         
    }
    
    // OrderingStoreDetails used by Order
    type OrderingStoreDetails struct {

        
            State string  `json:"state"`
            Phone string  `json:"phone"`
            ID float64  `json:"id"`
            StoreName string  `json:"store_name"`
            Code string  `json:"code"`
            Address string  `json:"address"`
            ContactPerson string  `json:"contact_person"`
            Pincode string  `json:"pincode"`
            Meta map[string]interface{}  `json:"meta"`
            Country string  `json:"country"`
            City string  `json:"city"`
         
    }
    
    // UserDetailsData used by Order
    type UserDetailsData struct {

        
            AddressType string  `json:"address_type"`
            Area string  `json:"area"`
            State string  `json:"state"`
            Name string  `json:"name"`
            Phone string  `json:"phone"`
            Address string  `json:"address"`
            Pincode string  `json:"pincode"`
            Landmark string  `json:"landmark"`
            Country string  `json:"country"`
            Address1 string  `json:"address1"`
            Email string  `json:"email"`
            City string  `json:"city"`
         
    }
    
    // PhoneDetails used by Order
    type PhoneDetails struct {

        
            CountryCode float64  `json:"country_code"`
            Number string  `json:"number"`
         
    }
    
    // ContactDetails used by Order
    type ContactDetails struct {

        
            Emails []string  `json:"emails"`
            Phone []PhoneDetails  `json:"phone"`
         
    }
    
    // CompanyDetails used by Order
    type CompanyDetails struct {

        
            CompanyID float64  `json:"company_id"`
            CompanyCin string  `json:"company_cin"`
            CompanyContact ContactDetails  `json:"company_contact"`
            CompanyGst string  `json:"company_gst"`
            Address map[string]interface{}  `json:"address"`
            CompanyName string  `json:"company_name"`
         
    }
    
    // DPDetailsData used by Order
    type DPDetailsData struct {

        
            EwayBillID string  `json:"eway_bill_id"`
            TrackURL string  `json:"track_url"`
            GstTag string  `json:"gst_tag"`
            AwbNo string  `json:"awb_no"`
            Name string  `json:"name"`
            ID float64  `json:"id"`
            Pincode string  `json:"pincode"`
            Country string  `json:"country"`
         
    }
    
    // BuyerDetails used by Order
    type BuyerDetails struct {

        
            State string  `json:"state"`
            Name string  `json:"name"`
            Address string  `json:"address"`
            Gstin string  `json:"gstin"`
            Pincode float64  `json:"pincode"`
            AjioSiteID string  `json:"ajio_site_id"`
            City string  `json:"city"`
         
    }
    
    // DebugInfo used by Order
    type DebugInfo struct {

        
            StormbreakerUUID string  `json:"stormbreaker_uuid"`
         
    }
    
    // EinvoiceInfo used by Order
    type EinvoiceInfo struct {

        
            Invoice map[string]interface{}  `json:"invoice"`
            CreditNote map[string]interface{}  `json:"credit_note"`
         
    }
    
    // ShipmentMeta used by Order
    type ShipmentMeta struct {

        
            LockData LockData  `json:"lock_data"`
            External map[string]interface{}  `json:"external"`
            PackagingName string  `json:"packaging_name"`
            PoNumber string  `json:"po_number"`
            AutoTriggerDpAssignmentAcf bool  `json:"auto_trigger_dp_assignment_acf"`
            SameStoreAvailable bool  `json:"same_store_available"`
            ReturnAwbNumber string  `json:"return_awb_number"`
            AwbNumber string  `json:"awb_number"`
            DpID string  `json:"dp_id"`
            ReturnStoreNode float64  `json:"return_store_node"`
            MarketplaceStoreID string  `json:"marketplace_store_id"`
            BagWeight map[string]interface{}  `json:"bag_weight"`
            StoreInvoiceUpdatedDate string  `json:"store_invoice_updated_date"`
            B2bBuyerDetails BuyerDetails  `json:"b2b_buyer_details"`
            DpSortKey string  `json:"dp_sort_key"`
            BoxType string  `json:"box_type"`
            ShipmentVolumetricWeight float64  `json:"shipment_volumetric_weight"`
            EwaybillInfo map[string]interface{}  `json:"ewaybill_info"`
            DueDate string  `json:"due_date"`
            ForwardAffiliateShipmentID string  `json:"forward_affiliate_shipment_id"`
            DebugInfo DebugInfo  `json:"debug_info"`
            ParentDpID string  `json:"parent_dp_id"`
            AssignDpFromSb bool  `json:"assign_dp_from_sb"`
            ReturnDetails map[string]interface{}  `json:"return_details"`
            ReturnAffiliateOrderID string  `json:"return_affiliate_order_id"`
            Weight float64  `json:"weight"`
            ShipmentWeight float64  `json:"shipment_weight"`
            B2cBuyerDetails map[string]interface{}  `json:"b2c_buyer_details"`
            FulfilmentPriorityText string  `json:"fulfilment_priority_text"`
            DpOptions map[string]interface{}  `json:"dp_options"`
            OrderType string  `json:"order_type"`
            Dimension Dimensions  `json:"dimension"`
            ReturnAffiliateShipmentID string  `json:"return_affiliate_shipment_id"`
            ForwardAffiliateOrderID string  `json:"forward_affiliate_order_id"`
            ShipmentTags []ShipmentTags  `json:"shipment_tags"`
            Formatted Formatted  `json:"formatted"`
            DpName string  `json:"dp_name"`
            EinvoiceInfo EinvoiceInfo  `json:"einvoice_info"`
            Timestamp ShipmentTimeStamp  `json:"timestamp"`
         
    }
    
    // PDFLinks used by Order
    type PDFLinks struct {

        
            Invoice string  `json:"invoice"`
            InvoiceType string  `json:"invoice_type"`
            LabelA4 string  `json:"label_a4"`
            LabelA6 string  `json:"label_a6"`
            InvoiceA6 string  `json:"invoice_a6"`
            LabelPos string  `json:"label_pos"`
            Label string  `json:"label"`
            LabelType string  `json:"label_type"`
            CreditNoteURL string  `json:"credit_note_url"`
            PoInvoice string  `json:"po_invoice"`
            LabelExport string  `json:"label_export"`
            InvoiceExport string  `json:"invoice_export"`
            B2b string  `json:"b2b"`
            InvoicePos string  `json:"invoice_pos"`
            InvoiceA4 string  `json:"invoice_a4"`
            DeliveryChallanA4 string  `json:"delivery_challan_a4"`
         
    }
    
    // AffiliateDetails used by Order
    type AffiliateDetails struct {

        
            AffiliateOrderID string  `json:"affiliate_order_id"`
            AffiliateStoreID string  `json:"affiliate_store_id"`
            AffiliateID string  `json:"affiliate_id"`
            CompanyAffiliateTag string  `json:"company_affiliate_tag"`
            ShipmentMeta ShipmentMeta  `json:"shipment_meta"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            AdID string  `json:"ad_id"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
            AffiliateMeta AffiliateMeta  `json:"affiliate_meta"`
            PdfLinks PDFLinks  `json:"pdf_links"`
         
    }
    
    // InvoiceInfo used by Order
    type InvoiceInfo struct {

        
            ExternalInvoiceID string  `json:"external_invoice_id"`
            StoreInvoiceID string  `json:"store_invoice_id"`
            InvoiceURL string  `json:"invoice_url"`
            LabelURL string  `json:"label_url"`
            CreditNoteID string  `json:"credit_note_id"`
            UpdatedDate string  `json:"updated_date"`
         
    }
    
    // OrderDetailsData used by Order
    type OrderDetailsData struct {

        
            TaxDetails map[string]interface{}  `json:"tax_details"`
            AffiliateID string  `json:"affiliate_id"`
            OrderValue string  `json:"order_value"`
            FyndOrderID string  `json:"fynd_order_id"`
            OrderingChannelLogo map[string]interface{}  `json:"ordering_channel_logo"`
            Source string  `json:"source"`
            OrderDate string  `json:"order_date"`
            OrderingChannel string  `json:"ordering_channel"`
            CodCharges string  `json:"cod_charges"`
         
    }
    
    // ShipmentStatusData used by Order
    type ShipmentStatusData struct {

        
            DisplayName string  `json:"display_name"`
            ID float64  `json:"id"`
            BagList []string  `json:"bag_list"`
            CreatedAt string  `json:"created_at"`
            ShipmentID string  `json:"shipment_id"`
            Meta map[string]interface{}  `json:"meta"`
            Status string  `json:"status"`
         
    }
    
    // CurrentStatus used by Order
    type CurrentStatus struct {

        
            UpdatedAt float64  `json:"updated_at"`
            KafkaSync bool  `json:"kafka_sync"`
            StoreID float64  `json:"store_id"`
            StateType string  `json:"state_type"`
            DeliveryPartnerID float64  `json:"delivery_partner_id"`
            BagID float64  `json:"bag_id"`
            ID float64  `json:"id"`
            CreatedAt string  `json:"created_at"`
            ShipmentID string  `json:"shipment_id"`
            BagStateMapper BagStateMapper  `json:"bag_state_mapper"`
            Status string  `json:"status"`
            DeliveryAwbNumber string  `json:"delivery_awb_number"`
            StateID float64  `json:"state_id"`
         
    }
    
    // OrderBrandName used by Order
    type OrderBrandName struct {

        
            BrandName string  `json:"brand_name"`
            Company float64  `json:"company"`
            ID float64  `json:"id"`
            ModifiedOn string  `json:"modified_on"`
            Logo string  `json:"logo"`
            CreatedOn string  `json:"created_on"`
         
    }
    
    // BagConfigs used by Order
    type BagConfigs struct {

        
            IsActive bool  `json:"is_active"`
            CanBeCancelled bool  `json:"can_be_cancelled"`
            EnableTracking bool  `json:"enable_tracking"`
            AllowForceReturn bool  `json:"allow_force_return"`
            IsCustomerReturnAllowed bool  `json:"is_customer_return_allowed"`
            IsReturnable bool  `json:"is_returnable"`
         
    }
    
    // AffiliateBagsDetails used by Order
    type AffiliateBagsDetails struct {

        
            AffiliateBagID string  `json:"affiliate_bag_id"`
            CouponCode string  `json:"coupon_code"`
         
    }
    
    // BagGST used by Order
    type BagGST struct {

        
            GstTag string  `json:"gst_tag"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            IgstGstFee string  `json:"igst_gst_fee"`
            HsnCodeID string  `json:"hsn_code_id"`
            IgstTaxPercentage float64  `json:"igst_tax_percentage"`
            CgstTaxPercentage float64  `json:"cgst_tax_percentage"`
            HsnCode string  `json:"hsn_code"`
            GstinCode string  `json:"gstin_code"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            GstFee float64  `json:"gst_fee"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            ValueOfGood float64  `json:"value_of_good"`
            CgstGstFee string  `json:"cgst_gst_fee"`
            SgstGstFee string  `json:"sgst_gst_fee"`
            IsDefaultHsnCode bool  `json:"is_default_hsn_code"`
            SgstTaxPercentage float64  `json:"sgst_tax_percentage"`
         
    }
    
    // ReturnConfig1 used by Order
    type ReturnConfig1 struct {

        
            Time float64  `json:"time"`
            Unit string  `json:"unit"`
            Returnable bool  `json:"returnable"`
         
    }
    
    // OrderBagArticle used by Order
    type OrderBagArticle struct {

        
            ReturnConfig ReturnConfig1  `json:"return_config"`
            Size string  `json:"size"`
            UID string  `json:"uid"`
            Identifiers map[string]interface{}  `json:"identifiers"`
         
    }
    
    // DiscountRules used by Order
    type DiscountRules struct {

        
            Type string  `json:"type"`
            Value float64  `json:"value"`
         
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
    
    // AppliedPromos used by Order
    type AppliedPromos struct {

        
            DiscountRules []DiscountRules  `json:"discount_rules"`
            MrpPromotion bool  `json:"mrp_promotion"`
            ArticleQuantity float64  `json:"article_quantity"`
            PromotionName string  `json:"promotion_name"`
            Amount float64  `json:"amount"`
            PromotionType string  `json:"promotion_type"`
            BuyRules []BuyRules  `json:"buy_rules"`
            PromoID string  `json:"promo_id"`
         
    }
    
    // GiftCard used by Order
    type GiftCard struct {

        
            GiftMessage string  `json:"gift_message"`
            DisplayText string  `json:"display_text"`
            GiftPrice float64  `json:"gift_price"`
            IsGiftApplied bool  `json:"is_gift_applied"`
         
    }
    
    // B2BPODetails used by Order
    type B2BPODetails struct {

        
            TotalGstPercentage float64  `json:"total_gst_percentage"`
            ItemBasePrice float64  `json:"item_base_price"`
            DocketNumber string  `json:"docket_number"`
            PoTaxAmount float64  `json:"po_tax_amount"`
            PartialCanRet bool  `json:"partial_can_ret"`
            PoLineAmount float64  `json:"po_line_amount"`
         
    }
    
    // BagMeta used by Order
    type BagMeta struct {

        
            GiftCard GiftCard  `json:"gift_card"`
            CustomMessage string  `json:"custom_message"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            DocketNumber string  `json:"docket_number"`
            GroupID string  `json:"group_id"`
            B2bPoDetails B2BPODetails  `json:"b2b_po_details"`
            PartialCanRet bool  `json:"partial_can_ret"`
            CustomJson map[string]interface{}  `json:"custom_json"`
         
    }
    
    // OrderBags used by Order
    type OrderBags struct {

        
            SellerIdentifier string  `json:"seller_identifier"`
            CurrentStatus CurrentStatus  `json:"current_status"`
            DeliveryAddress PlatformDeliveryAddress  `json:"delivery_address"`
            CanCancel bool  `json:"can_cancel"`
            Brand OrderBrandName  `json:"brand"`
            BagConfigs BagConfigs  `json:"bag_configs"`
            ParentPromoBags map[string]interface{}  `json:"parent_promo_bags"`
            Identifier string  `json:"identifier"`
            AffiliateBagDetails AffiliateBagsDetails  `json:"affiliate_bag_details"`
            GstDetails BagGST  `json:"gst_details"`
            Item PlatformItem  `json:"item"`
            GroupID string  `json:"group_id"`
            Quantity float64  `json:"quantity"`
            Article OrderBagArticle  `json:"article"`
            Prices Prices  `json:"prices"`
            DisplayName string  `json:"display_name"`
            BagID float64  `json:"bag_id"`
            CanReturn bool  `json:"can_return"`
            EntityType string  `json:"entity_type"`
            FinancialBreakup FinancialBreakup  `json:"financial_breakup"`
            AppliedPromos []AppliedPromos  `json:"applied_promos"`
            LineNumber float64  `json:"line_number"`
            Meta BagMeta  `json:"meta"`
         
    }
    
    // FulfillingStore used by Order
    type FulfillingStore struct {

        
            State string  `json:"state"`
            Phone string  `json:"phone"`
            ID float64  `json:"id"`
            StoreName string  `json:"store_name"`
            Code string  `json:"code"`
            Address string  `json:"address"`
            ContactPerson string  `json:"contact_person"`
            Pincode string  `json:"pincode"`
            FulfillmentChannel string  `json:"fulfillment_channel"`
            Country string  `json:"country"`
            Meta map[string]interface{}  `json:"meta"`
            City string  `json:"city"`
         
    }
    
    // ShipmentDetails used by Order
    type ShipmentDetails struct {

        
            LockMessage string  `json:"lock_message"`
            ActionToStatus map[string]interface{}  `json:"action_to_status"`
            LockStatus bool  `json:"lock_status"`
         
    }
    
    // ShipmentPayments used by Order
    type ShipmentPayments struct {

        
            Mode string  `json:"mode"`
            Source string  `json:"source"`
            Logo string  `json:"logo"`
         
    }
    
    // PlatformShipment used by Order
    type PlatformShipment struct {

        
            ForwardShipmentID string  `json:"forward_shipment_id"`
            DpAssignment bool  `json:"dp_assignment"`
            TrackingList []TrackingList  `json:"tracking_list"`
            TotalItems float64  `json:"total_items"`
            ShipmentQuantity float64  `json:"shipment_quantity"`
            EstimatedSlaTime string  `json:"estimated_sla_time"`
            OrderingStore OrderingStoreDetails  `json:"ordering_store"`
            PdfLinks map[string]interface{}  `json:"pdf_links"`
            BagStatusHistory []BagStatusHistory  `json:"bag_status_history"`
            CanUpdateDimension bool  `json:"can_update_dimension"`
            EnableDpTracking bool  `json:"enable_dp_tracking"`
            BillingDetails UserDetailsData  `json:"billing_details"`
            CustomMessage string  `json:"custom_message"`
            ShipmentImages []string  `json:"shipment_images"`
            LockStatus bool  `json:"lock_status"`
            CompanyDetails CompanyDetails  `json:"company_details"`
            IsDpAssignEnabled bool  `json:"is_dp_assign_enabled"`
            DpDetails DPDetailsData  `json:"dp_details"`
            AffiliateDetails AffiliateDetails  `json:"affiliate_details"`
            FulfilmentPriority float64  `json:"fulfilment_priority"`
            GstDetails GSTDetailsData  `json:"gst_details"`
            Invoice InvoiceInfo  `json:"invoice"`
            UserAgent string  `json:"user_agent"`
            DeliveryDetails UserDetailsData  `json:"delivery_details"`
            ShipmentCreatedAt string  `json:"shipment_created_at"`
            User UserDataInfo  `json:"user"`
            Order OrderDetailsData  `json:"order"`
            ShipmentStatus string  `json:"shipment_status"`
            Vertical string  `json:"vertical"`
            TotalBags float64  `json:"total_bags"`
            Status ShipmentStatusData  `json:"status"`
            PriorityText string  `json:"priority_text"`
            JourneyType string  `json:"journey_type"`
            Prices Prices  `json:"prices"`
            Bags []OrderBags  `json:"bags"`
            OperationalStatus string  `json:"operational_status"`
            PackagingType string  `json:"packaging_type"`
            DeliverySlot map[string]interface{}  `json:"delivery_slot"`
            PickedDate string  `json:"picked_date"`
            InvoiceID string  `json:"invoice_id"`
            Coupon map[string]interface{}  `json:"coupon"`
            PaymentMode string  `json:"payment_mode"`
            FulfillingStore FulfillingStore  `json:"fulfilling_store"`
            PaymentMethods map[string]interface{}  `json:"payment_methods"`
            ShipmentDetails ShipmentDetails  `json:"shipment_details"`
            CustomMeta []map[string]interface{}  `json:"custom_meta"`
            ShipmentID string  `json:"shipment_id"`
            Meta ShipmentMeta  `json:"meta"`
            PlatformLogo string  `json:"platform_logo"`
            Payments ShipmentPayments  `json:"payments"`
         
    }
    
    // ShipmentInfoResponse used by Order
    type ShipmentInfoResponse struct {

        
            Shipments []PlatformShipment  `json:"shipments"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // TaxDetails used by Order
    type TaxDetails struct {

        
            Gstin string  `json:"gstin"`
            PanNo string  `json:"pan_no"`
         
    }
    
    // BillingStaffDetails used by Order
    type BillingStaffDetails struct {

        
            StaffID float64  `json:"staff_id"`
            User string  `json:"user"`
            FirstName string  `json:"first_name"`
            EmployeeCode string  `json:"employee_code"`
            LastName string  `json:"last_name"`
         
    }
    
    // TransactionData used by Order
    type TransactionData struct {

        
            Currency string  `json:"currency"`
            TransactionID string  `json:"transaction_id"`
            AmountPaid float64  `json:"amount_paid"`
            UniqueReferenceNumber string  `json:"unique_reference_number"`
            TerminalID string  `json:"terminal_id"`
            Status string  `json:"status"`
            Entity string  `json:"entity"`
            PaymentID string  `json:"payment_id"`
         
    }
    
    // PlatformUserDetails used by Order
    type PlatformUserDetails struct {

        
            PlatformUserID string  `json:"platform_user_id"`
            PlatformUserLastName string  `json:"platform_user_last_name"`
            PlatformUserFirstName string  `json:"platform_user_first_name"`
            PlatformUserEmployeeCode string  `json:"platform_user_employee_code"`
         
    }
    
    // OrderMeta used by Order
    type OrderMeta struct {

        
            BillingStaffDetails BillingStaffDetails  `json:"billing_staff_details"`
            MongoCartID float64  `json:"mongo_cart_id"`
            OrderingStore float64  `json:"ordering_store"`
            OrderChildEntities []string  `json:"order_child_entities"`
            OrderPlatform string  `json:"order_platform"`
            Files []map[string]interface{}  `json:"files"`
            TransactionData TransactionData  `json:"transaction_data"`
            EmployeeID float64  `json:"employee_id"`
            PaymentType string  `json:"payment_type"`
            Comment string  `json:"comment"`
            OrderTags []map[string]interface{}  `json:"order_tags"`
            PlatformUserDetails PlatformUserDetails  `json:"platform_user_details"`
            CartID float64  `json:"cart_id"`
            CompanyLogo string  `json:"company_logo"`
            Staff map[string]interface{}  `json:"staff"`
            CurrencySymbol string  `json:"currency_symbol"`
            OrderType string  `json:"order_type"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            CustomerNote string  `json:"customer_note"`
         
    }
    
    // OrderData used by Order
    type OrderData struct {

        
            TaxDetails TaxDetails  `json:"tax_details"`
            FyndOrderID string  `json:"fynd_order_id"`
            PaymentMethods map[string]interface{}  `json:"payment_methods"`
            OrderDate string  `json:"order_date"`
            Meta OrderMeta  `json:"meta"`
            Prices Prices  `json:"prices"`
         
    }
    
    // OrderDetailsResponse used by Order
    type OrderDetailsResponse struct {

        
            Order OrderData  `json:"order"`
            Success bool  `json:"success"`
            Shipments []PlatformShipment  `json:"shipments"`
         
    }
    
    // SubLane used by Order
    type SubLane struct {

        
            Index float64  `json:"index"`
            Value string  `json:"value"`
            TotalItems float64  `json:"total_items"`
            Text string  `json:"text"`
            Actions []map[string]interface{}  `json:"actions"`
         
    }
    
    // SuperLane used by Order
    type SuperLane struct {

        
            Text string  `json:"text"`
            Options []SubLane  `json:"options"`
            Value string  `json:"value"`
            TotalItems float64  `json:"total_items"`
         
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

        
            Display string  `json:"display"`
            Name string  `json:"name"`
            Value string  `json:"value"`
         
    }
    
    // PlatformOrderItems used by Order
    type PlatformOrderItems struct {

        
            UserInfo UserDataInfo  `json:"user_info"`
            PaymentMode string  `json:"payment_mode"`
            OrderValue float64  `json:"order_value"`
            OrderCreatedTime string  `json:"order_created_time"`
            Shipments []PlatformShipment  `json:"shipments"`
            Meta map[string]interface{}  `json:"meta"`
            TotalOrderValue float64  `json:"total_order_value"`
            OrderID string  `json:"order_id"`
            Channel PlatformChannel  `json:"channel"`
            BreakupValues []PlatformBreakupValues  `json:"breakup_values"`
         
    }
    
    // OrderListingResponse used by Order
    type OrderListingResponse struct {

        
            Success bool  `json:"success"`
            Page Page  `json:"page"`
            TotalCount float64  `json:"total_count"`
            Message string  `json:"message"`
            Items []PlatformOrderItems  `json:"items"`
            Lane string  `json:"lane"`
         
    }
    
    // PlatformTrack used by Order
    type PlatformTrack struct {

        
            ShipmentType string  `json:"shipment_type"`
            UpdatedAt string  `json:"updated_at"`
            RawStatus string  `json:"raw_status"`
            UpdatedTime string  `json:"updated_time"`
            Awb string  `json:"awb"`
            Meta map[string]interface{}  `json:"meta"`
            AccountName string  `json:"account_name"`
            Status string  `json:"status"`
            Reason string  `json:"reason"`
            LastLocationRecievedAt string  `json:"last_location_recieved_at"`
         
    }
    
    // PlatformShipmentTrack used by Order
    type PlatformShipmentTrack struct {

        
            Results []PlatformTrack  `json:"results"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // FilterInfoOption used by Order
    type FilterInfoOption struct {

        
            Value string  `json:"value"`
            PlaceholderText string  `json:"placeholder_text"`
            Name string  `json:"name"`
            MinSearchSize float64  `json:"min_search_size"`
            Text string  `json:"text"`
            ShowUI bool  `json:"show_ui"`
         
    }
    
    // FiltersInfo used by Order
    type FiltersInfo struct {

        
            Value string  `json:"value"`
            Required bool  `json:"required"`
            PlaceholderText string  `json:"placeholder_text"`
            Options []FilterInfoOption  `json:"options"`
            Type string  `json:"type"`
            Text string  `json:"text"`
         
    }
    
    // AdvanceFilterInfo used by Order
    type AdvanceFilterInfo struct {

        
            AppliedFilters map[string]interface{}  `json:"applied_filters"`
            Returned []FiltersInfo  `json:"returned"`
            Unfulfilled []FiltersInfo  `json:"unfulfilled"`
            Page map[string]interface{}  `json:"page"`
            ActionCentre []FiltersInfo  `json:"action_centre"`
            Processed []FiltersInfo  `json:"processed"`
            Filters []FiltersInfo  `json:"filters"`
         
    }
    
    // FiltersResponse used by Order
    type FiltersResponse struct {

        
            GlobalFilter []FiltersInfo  `json:"global_filter"`
            AdvanceFilter AdvanceFilterInfo  `json:"advance_filter"`
         
    }
    
    // URL used by Order
    type URL struct {

        
            URL string  `json:"url"`
         
    }
    
    // FileResponse used by Order
    type FileResponse struct {

        
            Cdn URL  `json:"cdn"`
            FileName string  `json:"file_name"`
         
    }
    
    // BulkActionTemplate used by Order
    type BulkActionTemplate struct {

        
            Text string  `json:"text"`
            Value string  `json:"value"`
         
    }
    
    // BulkActionTemplateResponse used by Order
    type BulkActionTemplateResponse struct {

        
            TemplateXSlug []BulkActionTemplate  `json:"template_x_slug"`
         
    }
    
    // QuestionSet used by Order
    type QuestionSet struct {

        
            DisplayName string  `json:"display_name"`
            ID float64  `json:"id"`
         
    }
    
    // Reason used by Order
    type Reason struct {

        
            QuestionSet []QuestionSet  `json:"question_set"`
            DisplayName string  `json:"display_name"`
            QcType []string  `json:"qc_type"`
            ID float64  `json:"id"`
         
    }
    
    // PlatformShipmentReasonsResponse used by Order
    type PlatformShipmentReasonsResponse struct {

        
            Reasons []Reason  `json:"reasons"`
            Success bool  `json:"success"`
         
    }
    
    // ShipmentResponseReasons used by Order
    type ShipmentResponseReasons struct {

        
            ReasonID float64  `json:"reason_id"`
            Reason string  `json:"reason"`
         
    }
    
    // ShipmentReasonsResponse used by Order
    type ShipmentReasonsResponse struct {

        
            Reasons []ShipmentResponseReasons  `json:"reasons"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // ArticleDetails used by Order
    type ArticleDetails struct {

        
            Status map[string]interface{}  `json:"status"`
         
    }
    
    // StoreAddress used by Order
    type StoreAddress struct {

        
            AddressType string  `json:"address_type"`
            Phone string  `json:"phone"`
            Latitude float64  `json:"latitude"`
            Pincode float64  `json:"pincode"`
            Longitude float64  `json:"longitude"`
            Address2 string  `json:"address2"`
            AddressCategory string  `json:"address_category"`
            Area string  `json:"area"`
            CreatedAt string  `json:"created_at"`
            CountryCode string  `json:"country_code"`
            Country string  `json:"country"`
            Address1 string  `json:"address1"`
            Email string  `json:"email"`
            UpdatedAt string  `json:"updated_at"`
            State string  `json:"state"`
            Version string  `json:"version"`
            ContactPerson string  `json:"contact_person"`
            Landmark string  `json:"landmark"`
            City string  `json:"city"`
         
    }
    
    // EInvoicePortalDetails used by Order
    type EInvoicePortalDetails struct {

        
            Password string  `json:"password"`
            User string  `json:"user"`
            Username string  `json:"username"`
         
    }
    
    // StoreEwaybill used by Order
    type StoreEwaybill struct {

        
            Enabled bool  `json:"enabled"`
         
    }
    
    // StoreEinvoice used by Order
    type StoreEinvoice struct {

        
            Password string  `json:"password"`
            Enabled bool  `json:"enabled"`
            User string  `json:"user"`
            Username string  `json:"username"`
         
    }
    
    // StoreGstCredentials used by Order
    type StoreGstCredentials struct {

        
            EWaybill StoreEwaybill  `json:"e_waybill"`
            EInvoice StoreEinvoice  `json:"e_invoice"`
         
    }
    
    // Document used by Order
    type Document struct {

        
            Value string  `json:"value"`
            DsType string  `json:"ds_type"`
            Verified bool  `json:"verified"`
            LegalName string  `json:"legal_name"`
            URL string  `json:"url"`
         
    }
    
    // StoreDocuments used by Order
    type StoreDocuments struct {

        
            Gst Document  `json:"gst"`
         
    }
    
    // StoreMeta used by Order
    type StoreMeta struct {

        
            EinvoicePortalDetails EInvoicePortalDetails  `json:"einvoice_portal_details"`
            DisplayName string  `json:"display_name"`
            EwaybillPortalDetails map[string]interface{}  `json:"ewaybill_portal_details"`
            GstNumber string  `json:"gst_number"`
            NotificationEmails []string  `json:"notification_emails"`
            AdditionalContactDetails map[string]interface{}  `json:"additional_contact_details"`
            Stage string  `json:"stage"`
            Timing []map[string]interface{}  `json:"timing"`
            GstCredentials StoreGstCredentials  `json:"gst_credentials"`
            Documents StoreDocuments  `json:"documents"`
            ProductReturnConfig map[string]interface{}  `json:"product_return_config"`
         
    }
    
    // Store used by Order
    type Store struct {

        
            StoreAddressJson StoreAddress  `json:"store_address_json"`
            StoreEmail string  `json:"store_email"`
            Name string  `json:"name"`
            Phone float64  `json:"phone"`
            VatNo string  `json:"vat_no"`
            BrandID interface{}  `json:"brand_id"`
            Latitude float64  `json:"latitude"`
            Pincode string  `json:"pincode"`
            Address2 string  `json:"address2"`
            Longitude float64  `json:"longitude"`
            BrandStoreTags []string  `json:"brand_store_tags"`
            LocationType string  `json:"location_type"`
            ParentStoreID float64  `json:"parent_store_id"`
            Code string  `json:"code"`
            CreatedAt string  `json:"created_at"`
            FulfillmentChannel string  `json:"fulfillment_channel"`
            MallArea string  `json:"mall_area"`
            AlohomoraUserID float64  `json:"alohomora_user_id"`
            IsActive bool  `json:"is_active"`
            IsArchived bool  `json:"is_archived"`
            CompanyID float64  `json:"company_id"`
            PackagingMaterialCount float64  `json:"packaging_material_count"`
            Country string  `json:"country"`
            Address1 string  `json:"address1"`
            SID string  `json:"s_id"`
            MallName string  `json:"mall_name"`
            UpdatedAt string  `json:"updated_at"`
            State string  `json:"state"`
            OrderIntegrationID string  `json:"order_integration_id"`
            StoreActiveFrom string  `json:"store_active_from"`
            ContactPerson string  `json:"contact_person"`
            Meta StoreMeta  `json:"meta"`
            LoginUsername string  `json:"login_username"`
            IsEnabledForRecon bool  `json:"is_enabled_for_recon"`
            City string  `json:"city"`
         
    }
    
    // Brand used by Order
    type Brand struct {

        
            IsVirtualInvoice bool  `json:"is_virtual_invoice"`
            CreditNoteExpiryDays float64  `json:"credit_note_expiry_days"`
            BrandName string  `json:"brand_name"`
            Company string  `json:"company"`
            ModifiedOn float64  `json:"modified_on"`
            BrandID float64  `json:"brand_id"`
            CreditNoteAllowed bool  `json:"credit_note_allowed"`
            StartDate string  `json:"start_date"`
            Logo string  `json:"logo"`
            CreatedOn float64  `json:"created_on"`
            PickupLocation string  `json:"pickup_location"`
            InvoicePrefix string  `json:"invoice_prefix"`
            ScriptLastRan string  `json:"script_last_ran"`
         
    }
    
    // BagGSTDetails used by Order
    type BagGSTDetails struct {

        
            GstTag string  `json:"gst_tag"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            IgstGstFee string  `json:"igst_gst_fee"`
            HsnCodeID string  `json:"hsn_code_id"`
            IgstTaxPercentage float64  `json:"igst_tax_percentage"`
            CgstTaxPercentage float64  `json:"cgst_tax_percentage"`
            HsnCode string  `json:"hsn_code"`
            GstinCode string  `json:"gstin_code"`
            GstFee float64  `json:"gst_fee"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            ValueOfGood float64  `json:"value_of_good"`
            CgstGstFee string  `json:"cgst_gst_fee"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            SgstGstFee string  `json:"sgst_gst_fee"`
            IsDefaultHsnCode bool  `json:"is_default_hsn_code"`
            SgstTaxPercentage float64  `json:"sgst_tax_percentage"`
         
    }
    
    // Attributes used by Order
    type Attributes struct {

        
            PrimaryMaterial string  `json:"primary_material"`
            BrandName string  `json:"brand_name"`
            Name string  `json:"name"`
            PrimaryColorHex string  `json:"primary_color_hex"`
            Essential string  `json:"essential"`
            Gender []string  `json:"gender"`
            MarketerName string  `json:"marketer_name"`
            PrimaryColor string  `json:"primary_color"`
            MarketerAddress string  `json:"marketer_address"`
         
    }
    
    // Item used by Order
    type Item struct {

        
            L1Category []string  `json:"l1_category"`
            BranchURL string  `json:"branch_url"`
            Name string  `json:"name"`
            BrandID float64  `json:"brand_id"`
            L3CategoryName string  `json:"l3_category_name"`
            CanCancel bool  `json:"can_cancel"`
            Attributes Attributes  `json:"attributes"`
            Brand string  `json:"brand"`
            Code string  `json:"code"`
            DepartmentID float64  `json:"department_id"`
            L2Category []string  `json:"l2_category"`
            L1CategoryID float64  `json:"l1_category_id"`
            ItemID float64  `json:"item_id"`
            Gender string  `json:"gender"`
            Size string  `json:"size"`
            WebstoreProductURL string  `json:"webstore_product_url"`
            SlugKey string  `json:"slug_key"`
            L3Category float64  `json:"l3_category"`
            Image []string  `json:"image"`
            L2CategoryID float64  `json:"l2_category_id"`
            Color string  `json:"color"`
            CanReturn bool  `json:"can_return"`
            Meta map[string]interface{}  `json:"meta"`
            LastUpdatedAt string  `json:"last_updated_at"`
         
    }
    
    // BagReturnableCancelableStatus1 used by Order
    type BagReturnableCancelableStatus1 struct {

        
            IsActive bool  `json:"is_active"`
            CanBeCancelled bool  `json:"can_be_cancelled"`
            EnableTracking bool  `json:"enable_tracking"`
            IsCustomerReturnAllowed bool  `json:"is_customer_return_allowed"`
            IsReturnable bool  `json:"is_returnable"`
         
    }
    
    // BagDetailsPlatformResponse used by Order
    type BagDetailsPlatformResponse struct {

        
            SellerIdentifier string  `json:"seller_identifier"`
            BagStatus []BagStatusHistory  `json:"bag_status"`
            ArticleDetails ArticleDetails  `json:"article_details"`
            Dates Dates  `json:"dates"`
            CurrentStatus BagStatusHistory  `json:"current_status"`
            OrderingStore Store  `json:"ordering_store"`
            BagStatusHistory BagStatusHistory  `json:"bag_status_history"`
            NoOfBagsOrder float64  `json:"no_of_bags_order"`
            OriginalBagList []float64  `json:"original_bag_list"`
            Brand Brand  `json:"brand"`
            Type string  `json:"type"`
            Identifier string  `json:"identifier"`
            AffiliateDetails AffiliateDetails  `json:"affiliate_details"`
            ParentPromoBags map[string]interface{}  `json:"parent_promo_bags"`
            AffiliateBagDetails AffiliateBagDetails  `json:"affiliate_bag_details"`
            GstDetails BagGSTDetails  `json:"gst_details"`
            Reasons []map[string]interface{}  `json:"reasons"`
            RestoreCoupon bool  `json:"restore_coupon"`
            Item Item  `json:"item"`
            ID float64  `json:"id"`
            Quantity float64  `json:"quantity"`
            CurrentOperationalStatus BagStatusHistory  `json:"current_operational_status"`
            Status BagReturnableCancelableStatus1  `json:"status"`
            Article Article  `json:"article"`
            JourneyType string  `json:"journey_type"`
            Prices Prices  `json:"prices"`
            RestorePromos map[string]interface{}  `json:"restore_promos"`
            OperationalStatus string  `json:"operational_status"`
            QcRequired interface{}  `json:"qc_required"`
            DisplayName string  `json:"display_name"`
            OrderIntegrationID string  `json:"order_integration_id"`
            Tags []string  `json:"tags"`
            EntityType string  `json:"entity_type"`
            BagUpdateTime float64  `json:"bag_update_time"`
            ShipmentID string  `json:"shipment_id"`
            FinancialBreakup []FinancialBreakup  `json:"financial_breakup"`
            AppliedPromos []map[string]interface{}  `json:"applied_promos"`
            Meta BagMeta  `json:"meta"`
            LineNumber float64  `json:"line_number"`
         
    }
    
    // ErrorResponse used by Order
    type ErrorResponse struct {

        
            Error string  `json:"error"`
            Message string  `json:"message"`
         
    }
    
    // BagsPage used by Order
    type BagsPage struct {

        
            PageType string  `json:"page_type"`
            ItemTotal float64  `json:"item_total"`
            Size float64  `json:"size"`
            HasNext bool  `json:"has_next"`
            Current float64  `json:"current"`
         
    }
    
    // GetBagsPlatformResponse used by Order
    type GetBagsPlatformResponse struct {

        
            Page BagsPage  `json:"page"`
            Items []BagDetailsPlatformResponse  `json:"items"`
         
    }
    
    // InvalidateShipmentCachePayload used by Order
    type InvalidateShipmentCachePayload struct {

        
            ShipmentIds []string  `json:"shipment_ids"`
            AffiliateBagIds []string  `json:"affiliate_bag_ids"`
            BagIds []string  `json:"bag_ids"`
         
    }
    
    // InvalidateShipmentCacheNestedResponse used by Order
    type InvalidateShipmentCacheNestedResponse struct {

        
            Message string  `json:"message"`
            Status float64  `json:"status"`
            Error string  `json:"error"`
            ShipmentID string  `json:"shipment_id"`
         
    }
    
    // InvalidateShipmentCacheResponse used by Order
    type InvalidateShipmentCacheResponse struct {

        
            Response []InvalidateShipmentCacheNestedResponse  `json:"response"`
         
    }
    
    // ErrorResponse1 used by Order
    type ErrorResponse1 struct {

        
            Message string  `json:"message"`
            Status float64  `json:"status"`
            ErrorTrace string  `json:"error_trace"`
            Success bool  `json:"success"`
         
    }
    
    // StoreReassign used by Order
    type StoreReassign struct {

        
            ItemID string  `json:"item_id"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            StoreID float64  `json:"store_id"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
            MongoArticleID string  `json:"mongo_article_id"`
            ReasonIds []float64  `json:"reason_ids"`
            SetID string  `json:"set_id"`
            FyndOrderID string  `json:"fynd_order_id"`
            AffiliateID string  `json:"affiliate_id"`
            BagID float64  `json:"bag_id"`
         
    }
    
    // StoreReassignResponse used by Order
    type StoreReassignResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // Entities used by Order
    type Entities struct {

        
            ID string  `json:"id"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
            ReasonText string  `json:"reason_text"`
            AffiliateID string  `json:"affiliate_id"`
         
    }
    
    // UpdateShipmentLockPayload used by Order
    type UpdateShipmentLockPayload struct {

        
            EntityType string  `json:"entity_type"`
            ActionType string  `json:"action_type"`
            Action string  `json:"action"`
            Entities []Entities  `json:"entities"`
         
    }
    
    // OriginalFilter used by Order
    type OriginalFilter struct {

        
            AffiliateID string  `json:"affiliate_id"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
         
    }
    
    // Bags used by Order
    type Bags struct {

        
            AffiliateBagID string  `json:"affiliate_bag_id"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            BagID float64  `json:"bag_id"`
            IsLocked bool  `json:"is_locked"`
         
    }
    
    // CheckResponse used by Order
    type CheckResponse struct {

        
            IsShipmentLocked bool  `json:"is_shipment_locked"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            OriginalFilter OriginalFilter  `json:"original_filter"`
            Status string  `json:"status"`
            AffiliateID string  `json:"affiliate_id"`
            LockStatus string  `json:"lock_status"`
            Bags []Bags  `json:"bags"`
            ShipmentID string  `json:"shipment_id"`
            IsBagLocked bool  `json:"is_bag_locked"`
         
    }
    
    // UpdateShipmentLockResponse used by Order
    type UpdateShipmentLockResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
            CheckResponse []CheckResponse  `json:"check_response"`
         
    }
    
    // AnnouncementResponse used by Order
    type AnnouncementResponse struct {

        
            ID float64  `json:"id"`
            CreatedAt string  `json:"created_at"`
            LogoURL string  `json:"logo_url"`
            CompanyID float64  `json:"company_id"`
            Title string  `json:"title"`
            FromDatetime string  `json:"from_datetime"`
            ToDatetime string  `json:"to_datetime"`
            PlatformID string  `json:"platform_id"`
            PlatformName string  `json:"platform_name"`
            Description string  `json:"description"`
         
    }
    
    // AnnouncementsResponse used by Order
    type AnnouncementsResponse struct {

        
            Message string  `json:"message"`
            Announcements []AnnouncementResponse  `json:"announcements"`
            Success bool  `json:"success"`
         
    }
    
    // BaseResponse used by Order
    type BaseResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // Click2CallResponse used by Order
    type Click2CallResponse struct {

        
            CallID string  `json:"call_id"`
            Success bool  `json:"success"`
         
    }
    
    // ErrorDetail used by Order
    type ErrorDetail struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // Products used by Order
    type Products struct {

        
            LineNumber float64  `json:"line_number"`
            Identifier string  `json:"identifier"`
            Quantity float64  `json:"quantity"`
         
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
    
    // ProductsReasonsFilters used by Order
    type ProductsReasonsFilters struct {

        
            LineNumber float64  `json:"line_number"`
            Identifier string  `json:"identifier"`
            Quantity float64  `json:"quantity"`
         
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
    
    // ReasonsData used by Order
    type ReasonsData struct {

        
            Products []ProductsReasons  `json:"products"`
            Entities []EntitiesReasons  `json:"entities"`
         
    }
    
    // ShipmentsRequest used by Order
    type ShipmentsRequest struct {

        
            Products []Products  `json:"products"`
            DataUpdates DataUpdates  `json:"data_updates"`
            Reasons ReasonsData  `json:"reasons"`
            Identifier string  `json:"identifier"`
         
    }
    
    // StatuesRequest used by Order
    type StatuesRequest struct {

        
            Shipments []ShipmentsRequest  `json:"shipments"`
            SplitShipment bool  `json:"split_shipment"`
            ExcludeBagsNextState string  `json:"exclude_bags_next_state"`
            Status string  `json:"status"`
         
    }
    
    // UpdateShipmentStatusRequest used by Order
    type UpdateShipmentStatusRequest struct {

        
            Task bool  `json:"task"`
            Statuses []StatuesRequest  `json:"statuses"`
            UnlockBeforeTransition bool  `json:"unlock_before_transition"`
            LockAfterTransition bool  `json:"lock_after_transition"`
            ForceTransition bool  `json:"force_transition"`
         
    }
    
    // ShipmentsResponse used by Order
    type ShipmentsResponse struct {

        
            Code string  `json:"code"`
            Meta map[string]interface{}  `json:"meta"`
            FinalState map[string]interface{}  `json:"final_state"`
            Status float64  `json:"status"`
            StackTrace string  `json:"stack_trace"`
            Exception string  `json:"exception"`
            Identifier string  `json:"identifier"`
            Message string  `json:"message"`
         
    }
    
    // StatuesResponse used by Order
    type StatuesResponse struct {

        
            Shipments []ShipmentsResponse  `json:"shipments"`
         
    }
    
    // UpdateShipmentStatusResponseBody used by Order
    type UpdateShipmentStatusResponseBody struct {

        
            Statuses []StatuesResponse  `json:"statuses"`
         
    }
    
    // OrderUser used by Order
    type OrderUser struct {

        
            FirstName string  `json:"first_name"`
            City string  `json:"city"`
            State string  `json:"state"`
            Country string  `json:"country"`
            Mobile float64  `json:"mobile"`
            LastName string  `json:"last_name"`
            Phone float64  `json:"phone"`
            Address2 string  `json:"address2"`
            Email string  `json:"email"`
            Address1 string  `json:"address1"`
            Pincode string  `json:"pincode"`
         
    }
    
    // OrderPriority used by Order
    type OrderPriority struct {

        
            AffiliatePriorityCode string  `json:"affiliate_priority_code"`
            FulfilmentPriority float64  `json:"fulfilment_priority"`
            FulfilmentPriorityText string  `json:"fulfilment_priority_text"`
         
    }
    
    // ArticleDetails1 used by Order
    type ArticleDetails1 struct {

        
            Category map[string]interface{}  `json:"category"`
            Attributes map[string]interface{}  `json:"attributes"`
            ID string  `json:"_id"`
            Weight map[string]interface{}  `json:"weight"`
            Quantity float64  `json:"quantity"`
            BrandID float64  `json:"brand_id"`
            Dimension map[string]interface{}  `json:"dimension"`
         
    }
    
    // ShipmentDetails1 used by Order
    type ShipmentDetails1 struct {

        
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            Meta map[string]interface{}  `json:"meta"`
            Shipments float64  `json:"shipments"`
            Articles []ArticleDetails1  `json:"articles"`
            DpID float64  `json:"dp_id"`
            FulfillmentID float64  `json:"fulfillment_id"`
            BoxType string  `json:"box_type"`
         
    }
    
    // LocationDetails used by Order
    type LocationDetails struct {

        
            FulfillmentID float64  `json:"fulfillment_id"`
            FulfillmentType string  `json:"fulfillment_type"`
            Articles []ArticleDetails1  `json:"articles"`
         
    }
    
    // ShipmentConfig used by Order
    type ShipmentConfig struct {

        
            Source string  `json:"source"`
            ToPincode string  `json:"to_pincode"`
            Shipment []ShipmentDetails1  `json:"shipment"`
            PaymentMode string  `json:"payment_mode"`
            Action string  `json:"action"`
            Identifier string  `json:"identifier"`
            Journey string  `json:"journey"`
            LocationDetails LocationDetails  `json:"location_details"`
         
    }
    
    // ShipmentData used by Order
    type ShipmentData struct {

        
            ShipmentData ShipmentConfig  `json:"shipment_data"`
         
    }
    
    // UserData used by Order
    type UserData struct {

        
            BillingUser OrderUser  `json:"billing_user"`
            ShippingUser OrderUser  `json:"shipping_user"`
         
    }
    
    // MarketPlacePdf used by Order
    type MarketPlacePdf struct {

        
            Invoice string  `json:"invoice"`
            Label string  `json:"label"`
         
    }
    
    // AffiliateBag used by Order
    type AffiliateBag struct {

        
            PdfLinks MarketPlacePdf  `json:"pdf_links"`
            StoreID float64  `json:"store_id"`
            CompanyID float64  `json:"company_id"`
            SellerIdentifier string  `json:"seller_identifier"`
            UnitPrice float64  `json:"unit_price"`
            Identifier map[string]interface{}  `json:"identifier"`
            DeliveryCharge float64  `json:"delivery_charge"`
            ItemSize string  `json:"item_size"`
            ID string  `json:"_id"`
            Discount float64  `json:"discount"`
            PriceEffective float64  `json:"price_effective"`
            AmountPaid float64  `json:"amount_paid"`
            ItemID float64  `json:"item_id"`
            AffiliateStoreID string  `json:"affiliate_store_id"`
            HsnCodeID string  `json:"hsn_code_id"`
            Quantity float64  `json:"quantity"`
            TransferPrice float64  `json:"transfer_price"`
            AffiliateMeta map[string]interface{}  `json:"affiliate_meta"`
            ModifiedOn string  `json:"modified_on"`
            Sku string  `json:"sku"`
            AvlQty float64  `json:"avl_qty"`
            FyndStoreID string  `json:"fynd_store_id"`
            PriceMarked float64  `json:"price_marked"`
         
    }
    
    // OrderInfo used by Order
    type OrderInfo struct {

        
            ShippingAddress OrderUser  `json:"shipping_address"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            Discount float64  `json:"discount"`
            OrderPriority OrderPriority  `json:"order_priority"`
            CodCharges float64  `json:"cod_charges"`
            DeliveryCharges float64  `json:"delivery_charges"`
            Coupon string  `json:"coupon"`
            Shipment ShipmentData  `json:"shipment"`
            User UserData  `json:"user"`
            PaymentMode string  `json:"payment_mode"`
            Payment map[string]interface{}  `json:"payment"`
            Bags []AffiliateBag  `json:"bags"`
            Items map[string]interface{}  `json:"items"`
            BillingAddress OrderUser  `json:"billing_address"`
            OrderValue float64  `json:"order_value"`
         
    }
    
    // AffiliateStoreIdMapping used by Order
    type AffiliateStoreIdMapping struct {

        
            StoreID float64  `json:"store_id"`
            MarketplaceStoreID string  `json:"marketplace_store_id"`
         
    }
    
    // AffiliateAppConfigMeta used by Order
    type AffiliateAppConfigMeta struct {

        
            Name string  `json:"name"`
            Value string  `json:"value"`
         
    }
    
    // AffiliateAppConfig used by Order
    type AffiliateAppConfig struct {

        
            ID string  `json:"id"`
            CreatedAt string  `json:"created_at"`
            Meta []AffiliateAppConfigMeta  `json:"meta"`
            Secret string  `json:"secret"`
            Owner string  `json:"owner"`
            Name string  `json:"name"`
            Description string  `json:"description"`
            UpdatedAt string  `json:"updated_at"`
            Token string  `json:"token"`
         
    }
    
    // AffiliateInventoryStoreConfig used by Order
    type AffiliateInventoryStoreConfig struct {

        
            Store map[string]interface{}  `json:"store"`
         
    }
    
    // AffiliateInventoryOrderConfig used by Order
    type AffiliateInventoryOrderConfig struct {

        
            ForceReassignment bool  `json:"force_reassignment"`
         
    }
    
    // AffiliateInventoryPaymentConfig used by Order
    type AffiliateInventoryPaymentConfig struct {

        
            ModeOfPayment string  `json:"mode_of_payment"`
            Source string  `json:"source"`
         
    }
    
    // AffiliateInventoryLogisticsConfig used by Order
    type AffiliateInventoryLogisticsConfig struct {

        
            DpAssignment bool  `json:"dp_assignment"`
         
    }
    
    // AffiliateInventoryArticleAssignmentConfig used by Order
    type AffiliateInventoryArticleAssignmentConfig struct {

        
            PostOrderReassignment bool  `json:"post_order_reassignment"`
         
    }
    
    // AffiliateInventoryConfig used by Order
    type AffiliateInventoryConfig struct {

        
            Inventory AffiliateInventoryStoreConfig  `json:"inventory"`
            Order AffiliateInventoryOrderConfig  `json:"order"`
            Payment AffiliateInventoryPaymentConfig  `json:"payment"`
            Logistics AffiliateInventoryLogisticsConfig  `json:"logistics"`
            ArticleAssignment AffiliateInventoryArticleAssignmentConfig  `json:"article_assignment"`
         
    }
    
    // AffiliateConfig used by Order
    type AffiliateConfig struct {

        
            App AffiliateAppConfig  `json:"app"`
            Inventory AffiliateInventoryConfig  `json:"inventory"`
         
    }
    
    // Affiliate used by Order
    type Affiliate struct {

        
            ID string  `json:"id"`
            Config AffiliateConfig  `json:"config"`
            Token string  `json:"token"`
         
    }
    
    // OrderConfig used by Order
    type OrderConfig struct {

        
            AffiliateStoreIDMapping []AffiliateStoreIdMapping  `json:"affiliate_store_id_mapping"`
            ArticleLookup string  `json:"article_lookup"`
            Affiliate Affiliate  `json:"affiliate"`
            CreateUser bool  `json:"create_user"`
            StoreLookup string  `json:"store_lookup"`
            BagEndState string  `json:"bag_end_state"`
         
    }
    
    // CreateOrderPayload used by Order
    type CreateOrderPayload struct {

        
            OrderInfo OrderInfo  `json:"order_info"`
            AffiliateID string  `json:"affiliate_id"`
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
            DisplayText string  `json:"display_text"`
            Slug string  `json:"slug"`
            Description string  `json:"description"`
         
    }
    
    // GetActionsResponse used by Order
    type GetActionsResponse struct {

        
            Permissions []ActionInfo  `json:"permissions"`
         
    }
    
    // HistoryReason used by Order
    type HistoryReason struct {

        
            Category string  `json:"category"`
            State string  `json:"state"`
            Text string  `json:"text"`
            Quantity float64  `json:"quantity"`
            DislayName string  `json:"dislay_name"`
            Code float64  `json:"code"`
         
    }
    
    // HistoryMeta used by Order
    type HistoryMeta struct {

        
            StoreID float64  `json:"store_id"`
            Status2 string  `json:"status2"`
            ActivityType string  `json:"activity_type"`
            Slug string  `json:"slug"`
            ChannelType string  `json:"channel_type"`
            Status string  `json:"status"`
            ShortLink string  `json:"short_link"`
            Callerid string  `json:"callerid"`
            CallID string  `json:"call_id"`
            Receiver string  `json:"receiver"`
            Endtime string  `json:"endtime"`
            Billsec string  `json:"billsec"`
            Message string  `json:"message"`
            Recipient string  `json:"recipient"`
            Starttime string  `json:"starttime"`
            Status1 string  `json:"status1"`
            StoreName string  `json:"store_name"`
            Reason HistoryReason  `json:"reason"`
            ActivityComment string  `json:"activity_comment"`
            StoreCode string  `json:"store_code"`
            Recordpath string  `json:"recordpath"`
            Duration string  `json:"duration"`
            Caller string  `json:"caller"`
         
    }
    
    // HistoryDict used by Order
    type HistoryDict struct {

        
            TicketID string  `json:"ticket_id"`
            Meta HistoryMeta  `json:"meta"`
            L2Detail string  `json:"l2_detail"`
            Createdat string  `json:"createdat"`
            User string  `json:"user"`
            L3Detail string  `json:"l3_detail"`
            AssignedAgent string  `json:"assigned_agent"`
            DisplayMessage string  `json:"display_message"`
            TicketURL string  `json:"ticket_url"`
            Type string  `json:"type"`
            Message string  `json:"message"`
            BagID float64  `json:"bag_id"`
            L1Detail string  `json:"l1_detail"`
         
    }
    
    // ShipmentHistoryResponse used by Order
    type ShipmentHistoryResponse struct {

        
            ActivityHistory []HistoryDict  `json:"activity_history"`
            Success bool  `json:"success"`
         
    }
    
    // PostHistoryFilters used by Order
    type PostHistoryFilters struct {

        
            ShipmentID string  `json:"shipment_id"`
            LineNumber string  `json:"line_number"`
            Identifier string  `json:"identifier"`
         
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

        
            OrderID string  `json:"order_id"`
            ShipmentID float64  `json:"shipment_id"`
            PhoneNumber float64  `json:"phone_number"`
            CountryCode string  `json:"country_code"`
            BrandName string  `json:"brand_name"`
            PaymentMode string  `json:"payment_mode"`
            Message string  `json:"message"`
            CustomerName string  `json:"customer_name"`
            AmountPaid float64  `json:"amount_paid"`
         
    }
    
    // SendSmsPayload used by Order
    type SendSmsPayload struct {

        
            Slug string  `json:"slug"`
            BagID float64  `json:"bag_id"`
            Data SmsDataPayload  `json:"data"`
         
    }
    
    // OrderDetails used by Order
    type OrderDetails struct {

        
            CreatedAt string  `json:"created_at"`
            FyndOrderID string  `json:"fynd_order_id"`
         
    }
    
    // Meta used by Order
    type Meta struct {

        
            StateManagerUsed string  `json:"state_manager_used"`
            KafkaEmissionStatus float64  `json:"kafka_emission_status"`
         
    }
    
    // ShipmentDetail used by Order
    type ShipmentDetail struct {

        
            ID float64  `json:"id"`
            Meta Meta  `json:"meta"`
            BagList []float64  `json:"bag_list"`
            Status string  `json:"status"`
            ShipmentID string  `json:"shipment_id"`
            Remarks string  `json:"remarks"`
         
    }
    
    // OrderStatusData used by Order
    type OrderStatusData struct {

        
            OrderDetails OrderDetails  `json:"order_details"`
            Errors []string  `json:"errors"`
            ShipmentDetails []ShipmentDetail  `json:"shipment_details"`
         
    }
    
    // OrderStatusResult used by Order
    type OrderStatusResult struct {

        
            Result []OrderStatusData  `json:"result"`
            Success string  `json:"success"`
         
    }
    
    // Dimension used by Order
    type Dimension struct {

        
            Weight string  `json:"weight"`
            Height string  `json:"height"`
            Width float64  `json:"width"`
            Length float64  `json:"length"`
            PackagingType string  `json:"packaging_type"`
         
    }
    
    // UpdatePackagingDimensionsPayload used by Order
    type UpdatePackagingDimensionsPayload struct {

        
            ShipmentID string  `json:"shipment_id"`
            CurrentStatus string  `json:"current_status"`
            Dimension []Dimension  `json:"dimension"`
         
    }
    
    // UpdatePackagingDimensionsResponse used by Order
    type UpdatePackagingDimensionsResponse struct {

        
            Message string  `json:"message"`
         
    }
    
    // Tax used by Order
    type Tax struct {

        
            Breakup []map[string]interface{}  `json:"breakup"`
            Name string  `json:"name"`
            Rate float64  `json:"rate"`
            Amount map[string]interface{}  `json:"amount"`
         
    }
    
    // Charge used by Order
    type Charge struct {

        
            Type string  `json:"type"`
            Name string  `json:"name"`
            Amount map[string]interface{}  `json:"amount"`
            Code string  `json:"code"`
            Tax Tax  `json:"tax"`
         
    }
    
    // PaymentMethod used by Order
    type PaymentMethod struct {

        
            Meta map[string]interface{}  `json:"meta"`
            CollectBy string  `json:"collect_by"`
            RefundBy string  `json:"refund_by"`
            Name string  `json:"name"`
            Amount float64  `json:"amount"`
            TransactionData map[string]interface{}  `json:"transaction_data"`
            Mode string  `json:"mode"`
         
    }
    
    // PaymentInfo used by Order
    type PaymentInfo struct {

        
            PaymentMethods []PaymentMethod  `json:"payment_methods"`
            PrimaryMode string  `json:"primary_mode"`
         
    }
    
    // LineItem used by Order
    type LineItem struct {

        
            Charges []Charge  `json:"charges"`
            Meta map[string]interface{}  `json:"meta"`
            Quantity float64  `json:"quantity"`
            SellerIdentifier string  `json:"seller_identifier"`
            CustomMessasge string  `json:"custom_messasge"`
            ExternalLineID string  `json:"external_line_id"`
         
    }
    
    // ProcessingDates used by Order
    type ProcessingDates struct {

        
            DpPickupSlot map[string]interface{}  `json:"dp_pickup_slot"`
            PackByDate string  `json:"pack_by_date"`
            CustomerPickupSlot map[string]interface{}  `json:"customer_pickup_slot"`
            ConfirmByDate string  `json:"confirm_by_date"`
            DispatchByDate string  `json:"dispatch_by_date"`
            DispatchAfterDate string  `json:"dispatch_after_date"`
         
    }
    
    // Shipment used by Order
    type Shipment struct {

        
            ExternalShipmentID string  `json:"external_shipment_id"`
            LineItems []LineItem  `json:"line_items"`
            ProcessingDates ProcessingDates  `json:"processing_dates"`
            Meta map[string]interface{}  `json:"meta"`
            LocationID float64  `json:"location_id"`
            Priority float64  `json:"priority"`
         
    }
    
    // TaxInfo used by Order
    type TaxInfo struct {

        
            B2bGstinNumber string  `json:"b2b_gstin_number"`
            Gstin string  `json:"gstin"`
         
    }
    
    // BillingInfo used by Order
    type BillingInfo struct {

        
            FirstName string  `json:"first_name"`
            State string  `json:"state"`
            LastName string  `json:"last_name"`
            HouseNo string  `json:"house_no"`
            Address2 string  `json:"address2"`
            FloorNo string  `json:"floor_no"`
            AlternateEmail string  `json:"alternate_email"`
            PrimaryEmail string  `json:"primary_email"`
            City string  `json:"city"`
            Title string  `json:"title"`
            CountryCode string  `json:"country_code"`
            MiddleName string  `json:"middle_name"`
            PrimaryMobileNumber string  `json:"primary_mobile_number"`
            Country string  `json:"country"`
            AlternateMobileNumber string  `json:"alternate_mobile_number"`
            StateCode string  `json:"state_code"`
            Address1 string  `json:"address1"`
            Gender string  `json:"gender"`
            ExternalCustomerCode string  `json:"external_customer_code"`
            CustomerCode string  `json:"customer_code"`
            Pincode string  `json:"pincode"`
         
    }
    
    // ShippingInfo used by Order
    type ShippingInfo struct {

        
            FirstName string  `json:"first_name"`
            State string  `json:"state"`
            Landmark string  `json:"landmark"`
            LastName string  `json:"last_name"`
            HouseNo string  `json:"house_no"`
            GeoLocation map[string]interface{}  `json:"geo_location"`
            Address2 string  `json:"address2"`
            FloorNo string  `json:"floor_no"`
            AlternateEmail string  `json:"alternate_email"`
            PrimaryEmail string  `json:"primary_email"`
            City string  `json:"city"`
            Title string  `json:"title"`
            CountryCode string  `json:"country_code"`
            AddressType string  `json:"address_type"`
            Slot []map[string]interface{}  `json:"slot"`
            MiddleName string  `json:"middle_name"`
            PrimaryMobileNumber string  `json:"primary_mobile_number"`
            Country string  `json:"country"`
            ShippingType string  `json:"shipping_type"`
            AlternateMobileNumber string  `json:"alternate_mobile_number"`
            StateCode string  `json:"state_code"`
            Address1 string  `json:"address1"`
            Gender string  `json:"gender"`
            ExternalCustomerCode string  `json:"external_customer_code"`
            CustomerCode string  `json:"customer_code"`
            Pincode string  `json:"pincode"`
         
    }
    
    // CreateOrderAPI used by Order
    type CreateOrderAPI struct {

        
            Charges []Charge  `json:"charges"`
            Meta map[string]interface{}  `json:"meta"`
            CurrencyInfo map[string]interface{}  `json:"currency_info"`
            PaymentInfo PaymentInfo  `json:"payment_info"`
            Shipments []Shipment  `json:"shipments"`
            ExternalCreationDate string  `json:"external_creation_date"`
            TaxInfo TaxInfo  `json:"tax_info"`
            BillingInfo BillingInfo  `json:"billing_info"`
            Config map[string]interface{}  `json:"config"`
            ShippingInfo ShippingInfo  `json:"shipping_info"`
            ExternalOrderID string  `json:"external_order_id"`
         
    }
    
    // CreateOrderErrorReponse used by Order
    type CreateOrderErrorReponse struct {

        
            Info interface{}  `json:"info"`
            Meta string  `json:"meta"`
            Status float64  `json:"status"`
            StackTrace string  `json:"stack_trace"`
            Exception string  `json:"exception"`
            Code string  `json:"code"`
            Message string  `json:"message"`
            RequestID string  `json:"request_id"`
         
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
            ModeOfPayment string  `json:"mode_of_payment"`
            Source string  `json:"source"`
         
    }
    
    // DpConfiguration used by Order
    type DpConfiguration struct {

        
            ShippingBy string  `json:"shipping_by"`
         
    }
    
    // CreateChannelConfig used by Order
    type CreateChannelConfig struct {

        
            PaymentInfo CreateChannelPaymentInfo  `json:"payment_info"`
            LockStates []string  `json:"lock_states"`
            DpConfiguration DpConfiguration  `json:"dp_configuration"`
            LogoURL map[string]interface{}  `json:"logo_url"`
            ShipmentAssignment string  `json:"shipment_assignment"`
            LocationReassignment bool  `json:"location_reassignment"`
         
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

        
            IsUpserted bool  `json:"is_upserted"`
            IsInserted bool  `json:"is_inserted"`
            Acknowledged bool  `json:"acknowledged"`
         
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

        
            StartDate string  `json:"start_date"`
            OrderDetails []FyndOrderIdList  `json:"order_details"`
            EndDate string  `json:"end_date"`
            Mobile float64  `json:"mobile"`
         
    }
    
    // BagStateTransitionMap used by Order
    type BagStateTransitionMap struct {

        
            Fynd map[string]interface{}  `json:"fynd"`
            Affiliate map[string]interface{}  `json:"affiliate"`
         
    }
    

    
    // DeleteResponse used by Catalog
    type DeleteResponse struct {

        
            Message string  `json:"message"`
         
    }
    
    // ErrorResponse used by Catalog
    type ErrorResponse struct {

        
            Status float64  `json:"status"`
            Message string  `json:"message"`
            Error string  `json:"error"`
            Meta map[string]interface{}  `json:"meta"`
            Code string  `json:"code"`
         
    }
    
    // SearchKeywordResult used by Catalog
    type SearchKeywordResult struct {

        
            Query map[string]interface{}  `json:"query"`
            SortOn string  `json:"sort_on"`
         
    }
    
    // CreateSearchKeyword used by Catalog
    type CreateSearchKeyword struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Words []string  `json:"words"`
            IsActive bool  `json:"is_active"`
            AppID string  `json:"app_id"`
            Result SearchKeywordResult  `json:"result"`
         
    }
    
    // GetSearchWordsData used by Catalog
    type GetSearchWordsData struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Words []string  `json:"words"`
            IsActive bool  `json:"is_active"`
            AppID string  `json:"app_id"`
            UID string  `json:"uid"`
            Result map[string]interface{}  `json:"result"`
         
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
    
    // GetSearchWordsDetailResponse used by Catalog
    type GetSearchWordsDetailResponse struct {

        
            Items GetSearchWordsData  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // GetSearchWordsResponse used by Catalog
    type GetSearchWordsResponse struct {

        
            Items []GetSearchWordsData  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // AutocompletePageAction used by Catalog
    type AutocompletePageAction struct {

        
            Type string  `json:"type"`
            Params map[string]interface{}  `json:"params"`
            Query map[string]interface{}  `json:"query"`
            URL string  `json:"url"`
         
    }
    
    // AutocompleteAction used by Catalog
    type AutocompleteAction struct {

        
            Page AutocompletePageAction  `json:"page"`
            Type string  `json:"type"`
         
    }
    
    // AutoCompleteMedia used by Catalog
    type AutoCompleteMedia struct {

        
            Type string  `json:"type"`
            AspectRatio string  `json:"aspect_ratio"`
            URL string  `json:"url"`
         
    }
    
    // AutocompleteResult used by Catalog
    type AutocompleteResult struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Action AutocompleteAction  `json:"action"`
            Display string  `json:"display"`
            Logo AutoCompleteMedia  `json:"logo"`
         
    }
    
    // CreateAutocompleteKeyword used by Catalog
    type CreateAutocompleteKeyword struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Words []string  `json:"words"`
            IsActive bool  `json:"is_active"`
            Results []AutocompleteResult  `json:"results"`
            AppID string  `json:"app_id"`
         
    }
    
    // GetAutocompleteWordsData used by Catalog
    type GetAutocompleteWordsData struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Words []string  `json:"words"`
            Results []map[string]interface{}  `json:"results"`
            AppID string  `json:"app_id"`
            UID string  `json:"uid"`
         
    }
    
    // GetAutocompleteWordsResponse used by Catalog
    type GetAutocompleteWordsResponse struct {

        
            Items []GetAutocompleteWordsData  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // CreateAutocompleteWordsResponse used by Catalog
    type CreateAutocompleteWordsResponse struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Words []string  `json:"words"`
            Results []map[string]interface{}  `json:"results"`
            AppID string  `json:"app_id"`
         
    }
    
    // ProductBundleItem used by Catalog
    type ProductBundleItem struct {

        
            AllowRemove bool  `json:"allow_remove"`
            ProductUID float64  `json:"product_uid"`
            MinQuantity float64  `json:"min_quantity"`
            AutoSelect bool  `json:"auto_select"`
            AutoAddToCart bool  `json:"auto_add_to_cart"`
            MaxQuantity float64  `json:"max_quantity"`
         
    }
    
    // GetProductBundleCreateResponse used by Catalog
    type GetProductBundleCreateResponse struct {

        
            CreatedOn string  `json:"created_on"`
            Logo string  `json:"logo"`
            ID string  `json:"id"`
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            Products []ProductBundleItem  `json:"products"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Choice string  `json:"choice"`
            PageVisibility []string  `json:"page_visibility"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            Slug string  `json:"slug"`
            ModifiedOn string  `json:"modified_on"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Meta map[string]interface{}  `json:"meta"`
            CompanyID float64  `json:"company_id"`
         
    }
    
    // GetProductBundleListingResponse used by Catalog
    type GetProductBundleListingResponse struct {

        
            Items []GetProductBundleCreateResponse  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // ProductBundleRequest used by Catalog
    type ProductBundleRequest struct {

        
            CreatedOn string  `json:"created_on"`
            Logo string  `json:"logo"`
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            Products []ProductBundleItem  `json:"products"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Choice string  `json:"choice"`
            PageVisibility []string  `json:"page_visibility"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            Slug string  `json:"slug"`
            ModifiedOn string  `json:"modified_on"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Meta map[string]interface{}  `json:"meta"`
            CompanyID float64  `json:"company_id"`
         
    }
    
    // Size used by Catalog
    type Size struct {

        
            Value string  `json:"value"`
            IsAvailable bool  `json:"is_available"`
            Quantity float64  `json:"quantity"`
            Display string  `json:"display"`
         
    }
    
    // LimitedProductData used by Catalog
    type LimitedProductData struct {

        
            Images []string  `json:"images"`
            Sizes []string  `json:"sizes"`
            Identifier map[string]interface{}  `json:"identifier"`
            Quantity float64  `json:"quantity"`
            ShortDescription string  `json:"short_description"`
            Name string  `json:"name"`
            ItemCode string  `json:"item_code"`
            UID float64  `json:"uid"`
            Slug string  `json:"slug"`
            CountryOfOrigin string  `json:"country_of_origin"`
            Price map[string]interface{}  `json:"price"`
            Attributes map[string]interface{}  `json:"attributes"`
         
    }
    
    // Price used by Catalog
    type Price struct {

        
            MaxMarked float64  `json:"max_marked"`
            MinMarked float64  `json:"min_marked"`
            Currency string  `json:"currency"`
            MinEffective float64  `json:"min_effective"`
            MaxEffective float64  `json:"max_effective"`
         
    }
    
    // GetProducts used by Catalog
    type GetProducts struct {

        
            Sizes []Size  `json:"sizes"`
            AllowRemove bool  `json:"allow_remove"`
            ProductUID float64  `json:"product_uid"`
            MinQuantity float64  `json:"min_quantity"`
            AutoSelect bool  `json:"auto_select"`
            AutoAddToCart bool  `json:"auto_add_to_cart"`
            ProductDetails LimitedProductData  `json:"product_details"`
            Price Price  `json:"price"`
            MaxQuantity float64  `json:"max_quantity"`
         
    }
    
    // GetProductBundleResponse used by Catalog
    type GetProductBundleResponse struct {

        
            Logo string  `json:"logo"`
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            Products []GetProducts  `json:"products"`
            Choice string  `json:"choice"`
            PageVisibility []string  `json:"page_visibility"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            Slug string  `json:"slug"`
            Meta map[string]interface{}  `json:"meta"`
            CompanyID float64  `json:"company_id"`
         
    }
    
    // ProductBundleUpdateRequest used by Catalog
    type ProductBundleUpdateRequest struct {

        
            Logo string  `json:"logo"`
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            Products []ProductBundleItem  `json:"products"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Choice string  `json:"choice"`
            PageVisibility []string  `json:"page_visibility"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            Slug string  `json:"slug"`
            ModifiedOn string  `json:"modified_on"`
            Meta map[string]interface{}  `json:"meta"`
            CompanyID float64  `json:"company_id"`
         
    }
    
    // ListSizeGuide used by Catalog
    type ListSizeGuide struct {

        
            Items []map[string]interface{}  `json:"items"`
            Page map[string]interface{}  `json:"page"`
         
    }
    
    // Meta used by Catalog
    type Meta struct {

        
            Values []map[string]interface{}  `json:"values"`
            Headers map[string]interface{}  `json:"headers"`
            Unit string  `json:"unit"`
         
    }
    
    // Guide used by Catalog
    type Guide struct {

        
            Meta Meta  `json:"meta"`
         
    }
    
    // ValidateSizeGuide used by Catalog
    type ValidateSizeGuide struct {

        
            CreatedOn string  `json:"created_on"`
            Title string  `json:"title"`
            ID string  `json:"id"`
            Tag string  `json:"tag"`
            Guide Guide  `json:"guide"`
            Name string  `json:"name"`
            BrandID float64  `json:"brand_id"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Description string  `json:"description"`
            Image string  `json:"image"`
            CompanyID float64  `json:"company_id"`
            Active bool  `json:"active"`
            ModifiedOn string  `json:"modified_on"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Subtitle string  `json:"subtitle"`
         
    }
    
    // SuccessResponse used by Catalog
    type SuccessResponse struct {

        
            Success bool  `json:"success"`
         
    }
    
    // SizeGuideResponse used by Catalog
    type SizeGuideResponse struct {

        
            CreatedOn string  `json:"created_on"`
            Title string  `json:"title"`
            ID string  `json:"id"`
            Tag string  `json:"tag"`
            Guide map[string]interface{}  `json:"guide"`
            Name string  `json:"name"`
            BrandID float64  `json:"brand_id"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Subtitle string  `json:"subtitle"`
            ModifiedOn string  `json:"modified_on"`
            Active bool  `json:"active"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            CompanyID float64  `json:"company_id"`
         
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

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            AltText map[string]interface{}  `json:"alt_text"`
            IsCod bool  `json:"is_cod"`
            IsGift bool  `json:"is_gift"`
            CustomMeta []MetaFields  `json:"_custom_meta"`
            Moq ApplicationItemMOQ  `json:"moq"`
            Seo ApplicationItemSEO  `json:"seo"`
         
    }
    
    // SuccessResponse1 used by Catalog
    type SuccessResponse1 struct {

        
            UID float64  `json:"uid"`
            Success bool  `json:"success"`
         
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

        
            AltText map[string]interface{}  `json:"alt_text"`
            IsCod bool  `json:"is_cod"`
            IsGift bool  `json:"is_gift"`
            Moq MOQData  `json:"moq"`
            Seo SEOData  `json:"seo"`
         
    }
    
    // GetConfigMetadataResponse used by Catalog
    type GetConfigMetadataResponse struct {

        
            Data []map[string]interface{}  `json:"data"`
            Values []map[string]interface{}  `json:"values"`
            Condition []map[string]interface{}  `json:"condition"`
         
    }
    
    // PageResponseType used by Catalog
    type PageResponseType struct {

        
            Current float64  `json:"current"`
            Next float64  `json:"next"`
            HasNext bool  `json:"has_next"`
            TotalCount float64  `json:"total_count"`
         
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

        
            Key string  `json:"key"`
            Priority float64  `json:"priority"`
            Unit string  `json:"unit"`
            IsActive bool  `json:"is_active"`
            DisplayType string  `json:"display_type"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Logo string  `json:"logo"`
         
    }
    
    // AppConfigurationDetail used by Catalog
    type AppConfigurationDetail struct {

        
            Priority float64  `json:"priority"`
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            TemplateSlugs []string  `json:"template_slugs"`
            AppID string  `json:"app_id"`
            IsDefault bool  `json:"is_default"`
            Slug string  `json:"slug"`
            Attributes []AttributeDetailsGroup  `json:"attributes"`
            Logo string  `json:"logo"`
         
    }
    
    // ConfigSuccessResponse used by Catalog
    type ConfigSuccessResponse struct {

        
            Message string  `json:"message"`
         
    }
    
    // AppConfigurationsSort used by Catalog
    type AppConfigurationsSort struct {

        
            Key string  `json:"key"`
            Priority float64  `json:"priority"`
            DefaultKey string  `json:"default_key"`
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            AppID string  `json:"app_id"`
            IsDefault bool  `json:"is_default"`
            Logo string  `json:"logo"`
         
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
            Similar map[string]interface{}  `json:"similar"`
            Compare map[string]interface{}  `json:"compare"`
            Detail map[string]interface{}  `json:"detail"`
         
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

        
            Key string  `json:"key"`
            FilterTypes []string  `json:"filter_types"`
            Display string  `json:"display"`
            Units []map[string]interface{}  `json:"units"`
         
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
    
    // GetCatalogConfigurationMetaData used by Catalog
    type GetCatalogConfigurationMetaData struct {

        
            Product GetCatalogConfigurationDetailsProduct  `json:"product"`
            Listing MetaDataListingResponse  `json:"listing"`
         
    }
    
    // ProductSize used by Catalog
    type ProductSize struct {

        
            Max float64  `json:"max"`
            Min float64  `json:"min"`
         
    }
    
    // ConfigurationProductVariantConfig used by Catalog
    type ConfigurationProductVariantConfig struct {

        
            Key string  `json:"key"`
            Priority float64  `json:"priority"`
            IsActive bool  `json:"is_active"`
            DisplayType string  `json:"display_type"`
            Name string  `json:"name"`
            Size ProductSize  `json:"size"`
            Logo string  `json:"logo"`
         
    }
    
    // ConfigurationProductVariant used by Catalog
    type ConfigurationProductVariant struct {

        
            Config []ConfigurationProductVariantConfig  `json:"config"`
         
    }
    
    // ConfigurationProductConfig used by Catalog
    type ConfigurationProductConfig struct {

        
            Key string  `json:"key"`
            Priority float64  `json:"priority"`
            Title string  `json:"title"`
            IsActive bool  `json:"is_active"`
            Subtitle string  `json:"subtitle"`
            Size ProductSize  `json:"size"`
            Logo string  `json:"logo"`
         
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
    
    // ConfigurationListingSortConfig used by Catalog
    type ConfigurationListingSortConfig struct {

        
            Key string  `json:"key"`
            Priority float64  `json:"priority"`
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            Logo string  `json:"logo"`
         
    }
    
    // ConfigurationListingSort used by Catalog
    type ConfigurationListingSort struct {

        
            DefaultKey string  `json:"default_key"`
            Config []ConfigurationListingSortConfig  `json:"config"`
         
    }
    
    // ConfigurationBucketPoints used by Catalog
    type ConfigurationBucketPoints struct {

        
            Start float64  `json:"start"`
            End float64  `json:"end"`
            Display string  `json:"display"`
         
    }
    
    // ConfigurationListingFilterValue used by Catalog
    type ConfigurationListingFilterValue struct {

        
            Value string  `json:"value"`
            Priority []string  `json:"priority"`
            MapValues []map[string]interface{}  `json:"map_values"`
            Sort string  `json:"sort"`
            Map map[string]interface{}  `json:"map"`
            BucketPoints []ConfigurationBucketPoints  `json:"bucket_points"`
            Condition string  `json:"condition"`
         
    }
    
    // ConfigurationListingFilterConfig used by Catalog
    type ConfigurationListingFilterConfig struct {

        
            Key string  `json:"key"`
            Priority float64  `json:"priority"`
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            DisplayName string  `json:"display_name"`
            Type string  `json:"type"`
            ValueConfig ConfigurationListingFilterValue  `json:"value_config"`
            Logo string  `json:"logo"`
         
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
    
    // AppCatalogConfiguration used by Catalog
    type AppCatalogConfiguration struct {

        
            CreatedOn string  `json:"created_on"`
            ID string  `json:"id"`
            ConfigID string  `json:"config_id"`
            AppID string  `json:"app_id"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Type string  `json:"type"`
            ModifiedOn string  `json:"modified_on"`
            ConfigType string  `json:"config_type"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Product ConfigurationProduct  `json:"product"`
            Listing ConfigurationListing  `json:"listing"`
         
    }
    
    // GetAppCatalogConfiguration used by Catalog
    type GetAppCatalogConfiguration struct {

        
            Data AppCatalogConfiguration  `json:"data"`
            IsDefault bool  `json:"is_default"`
         
    }
    
    // AppConfiguration used by Catalog
    type AppConfiguration struct {

        
            CreatedOn string  `json:"created_on"`
            ConfigID string  `json:"config_id"`
            AppID string  `json:"app_id"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Type string  `json:"type"`
            ModifiedOn string  `json:"modified_on"`
            ConfigType string  `json:"config_type"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Product ConfigurationProduct  `json:"product"`
            Listing ConfigurationListing  `json:"listing"`
         
    }
    
    // GetCatalogConfigurationDetailsSchemaListing used by Catalog
    type GetCatalogConfigurationDetailsSchemaListing struct {

        
            Sort map[string]interface{}  `json:"sort"`
            Filter map[string]interface{}  `json:"filter"`
         
    }
    
    // EntityConfiguration used by Catalog
    type EntityConfiguration struct {

        
            ID string  `json:"id"`
            ConfigID string  `json:"config_id"`
            AppID string  `json:"app_id"`
            ConfigType string  `json:"config_type"`
            Product GetCatalogConfigurationDetailsProduct  `json:"product"`
            Listing GetCatalogConfigurationDetailsSchemaListing  `json:"listing"`
         
    }
    
    // GetAppCatalogEntityConfiguration used by Catalog
    type GetAppCatalogEntityConfiguration struct {

        
            Data EntityConfiguration  `json:"data"`
            IsDefault bool  `json:"is_default"`
         
    }
    
    // ProductFiltersKey used by Catalog
    type ProductFiltersKey struct {

        
            Name string  `json:"name"`
            Kind string  `json:"kind"`
            Display string  `json:"display"`
            Operators []string  `json:"operators"`
            Logo string  `json:"logo"`
         
    }
    
    // ProductFiltersValue used by Catalog
    type ProductFiltersValue struct {

        
            Value interface{}  `json:"value"`
            QueryFormat string  `json:"query_format"`
            Min float64  `json:"min"`
            Count float64  `json:"count"`
            CurrencySymbol string  `json:"currency_symbol"`
            DisplayFormat string  `json:"display_format"`
            SelectedMax float64  `json:"selected_max"`
            SelectedMin float64  `json:"selected_min"`
            CurrencyCode string  `json:"currency_code"`
            IsSelected bool  `json:"is_selected"`
            Max float64  `json:"max"`
            Display string  `json:"display"`
         
    }
    
    // ProductFilters used by Catalog
    type ProductFilters struct {

        
            Key ProductFiltersKey  `json:"key"`
            Values []ProductFiltersValue  `json:"values"`
         
    }
    
    // ProductSortOn used by Catalog
    type ProductSortOn struct {

        
            Name string  `json:"name"`
            IsSelected bool  `json:"is_selected"`
            Value string  `json:"value"`
         
    }
    
    // GetCollectionQueryOptionResponse used by Catalog
    type GetCollectionQueryOptionResponse struct {

        
            Filters []ProductFilters  `json:"filters"`
            Operators map[string]string  `json:"operators"`
            SortOn []ProductSortOn  `json:"sort_on"`
         
    }
    
    // BannerImage used by Catalog
    type BannerImage struct {

        
            Type string  `json:"type"`
            AspectRatio string  `json:"aspect_ratio"`
            URL string  `json:"url"`
         
    }
    
    // ImageUrls used by Catalog
    type ImageUrls struct {

        
            Landscape BannerImage  `json:"landscape"`
            Portrait BannerImage  `json:"portrait"`
         
    }
    
    // CollectionBadge used by Catalog
    type CollectionBadge struct {

        
            Text string  `json:"text"`
            Color string  `json:"color"`
         
    }
    
    // CollectionQuery used by Catalog
    type CollectionQuery struct {

        
            Op string  `json:"op"`
            Attribute string  `json:"attribute"`
            Value []interface{}  `json:"value"`
         
    }
    
    // CollectionScheduleStartEnd used by Catalog
    type CollectionScheduleStartEnd struct {

        
            Start string  `json:"start"`
            End string  `json:"end"`
         
    }
    
    // CollectionSchedule used by Catalog
    type CollectionSchedule struct {

        
            Start string  `json:"start"`
            NextSchedule []CollectionScheduleStartEnd  `json:"next_schedule"`
            End string  `json:"end"`
         
    }
    
    // CollectionActionPageQuery used by Catalog
    type CollectionActionPageQuery struct {

        
            Op string  `json:"op"`
            Attribute string  `json:"attribute"`
            Value []string  `json:"value"`
         
    }
    
    // CollectionActionPage used by Catalog
    type CollectionActionPage struct {

        
            Type string  `json:"type"`
            Query CollectionActionPageQuery  `json:"query"`
         
    }
    
    // Media used by Catalog
    type Media struct {

        
            Type string  `json:"type"`
            Meta map[string]interface{}  `json:"meta"`
            URL string  `json:"url"`
         
    }
    
    // GetCollectionDetailNest used by Catalog
    type GetCollectionDetailNest struct {

        
            AppID string  `json:"app_id"`
            Banners ImageUrls  `json:"banners"`
            Tags []string  `json:"tags"`
            Type string  `json:"type"`
            Badge CollectionBadge  `json:"badge"`
            Query []CollectionQuery  `json:"query"`
            Slug string  `json:"slug"`
            Name string  `json:"name"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Description string  `json:"description"`
            AllowFacets bool  `json:"allow_facets"`
            AllowSort bool  `json:"allow_sort"`
            Schedule CollectionSchedule  `json:"_schedule"`
            Priority float64  `json:"priority"`
            IsActive bool  `json:"is_active"`
            UID string  `json:"uid"`
            Meta map[string]interface{}  `json:"meta"`
            Action CollectionActionPage  `json:"action"`
            Logo Media  `json:"logo"`
         
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

        
            Items []GetCollectionDetailNest  `json:"items"`
            Page Page  `json:"page"`
            Filters CollectionListingFilter  `json:"filters"`
         
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
    
    // CollectionBadge1 used by Catalog
    type CollectionBadge1 struct {

        
            Color string  `json:"color"`
            Text string  `json:"text"`
         
    }
    
    // SeoDetail used by Catalog
    type SeoDetail struct {

        
            Description string  `json:"description"`
            Title string  `json:"title"`
         
    }
    
    // NextSchedule used by Catalog
    type NextSchedule struct {

        
            Start string  `json:"start"`
            End string  `json:"end"`
         
    }
    
    // CollectionSchedule1 used by Catalog
    type CollectionSchedule1 struct {

        
            Start string  `json:"start"`
            End string  `json:"end"`
            NextSchedule []NextSchedule  `json:"next_schedule"`
            Cron string  `json:"cron"`
            Duration float64  `json:"duration"`
         
    }
    
    // UserInfo used by Catalog
    type UserInfo struct {

        
            Email string  `json:"email"`
            UID string  `json:"uid"`
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
         
    }
    
    // CreateCollection used by Catalog
    type CreateCollection struct {

        
            AppID string  `json:"app_id"`
            Banners CollectionBanner  `json:"banners"`
            Tags []string  `json:"tags"`
            Type string  `json:"type"`
            Badge CollectionBadge1  `json:"badge"`
            Slug string  `json:"slug"`
            Query []CollectionQuery  `json:"query"`
            Seo SeoDetail  `json:"seo"`
            Name string  `json:"name"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Description string  `json:"description"`
            AllowFacets bool  `json:"allow_facets"`
            AllowSort bool  `json:"allow_sort"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            Schedule CollectionSchedule1  `json:"_schedule"`
            Priority float64  `json:"priority"`
            IsActive bool  `json:"is_active"`
            ModifiedBy UserInfo  `json:"modified_by"`
            CreatedBy UserInfo  `json:"created_by"`
            Meta map[string]interface{}  `json:"meta"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            SortOn string  `json:"sort_on"`
            Published bool  `json:"published"`
            IsVisible bool  `json:"is_visible"`
            Logo CollectionImage  `json:"logo"`
         
    }
    
    // CollectionCreateResponse used by Catalog
    type CollectionCreateResponse struct {

        
            Tag []string  `json:"tag"`
            AppID string  `json:"app_id"`
            Banners ImageUrls  `json:"banners"`
            Type string  `json:"type"`
            Badge map[string]interface{}  `json:"badge"`
            Query []CollectionQuery  `json:"query"`
            Slug string  `json:"slug"`
            Name string  `json:"name"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Description string  `json:"description"`
            AllowFacets bool  `json:"allow_facets"`
            AllowSort bool  `json:"allow_sort"`
            Schedule map[string]interface{}  `json:"_schedule"`
            Priority float64  `json:"priority"`
            IsActive bool  `json:"is_active"`
            Cron map[string]interface{}  `json:"cron"`
            Meta map[string]interface{}  `json:"meta"`
            SortOn string  `json:"sort_on"`
            Logo BannerImage  `json:"logo"`
         
    }
    
    // CollectionDetailResponse used by Catalog
    type CollectionDetailResponse struct {

        
            Priority float64  `json:"priority"`
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            AppID string  `json:"app_id"`
            Banners ImageUrls  `json:"banners"`
            Tags []string  `json:"tags"`
            Description string  `json:"description"`
            AllowFacets bool  `json:"allow_facets"`
            Type string  `json:"type"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Badge CollectionBadge  `json:"badge"`
            Query []CollectionQuery  `json:"query"`
            Slug string  `json:"slug"`
            AllowSort bool  `json:"allow_sort"`
            UID string  `json:"uid"`
            Meta map[string]interface{}  `json:"meta"`
            Schedule CollectionSchedule  `json:"_schedule"`
            Logo Media  `json:"logo"`
         
    }
    
    // UpdateCollection used by Catalog
    type UpdateCollection struct {

        
            Banners CollectionBanner  `json:"banners"`
            Tags []string  `json:"tags"`
            Badge CollectionBadge1  `json:"badge"`
            Type string  `json:"type"`
            Slug string  `json:"slug"`
            Query []CollectionQuery  `json:"query"`
            Seo SeoDetail  `json:"seo"`
            Name string  `json:"name"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Description string  `json:"description"`
            AllowFacets bool  `json:"allow_facets"`
            AllowSort bool  `json:"allow_sort"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            Schedule CollectionSchedule1  `json:"_schedule"`
            Priority float64  `json:"priority"`
            IsActive bool  `json:"is_active"`
            ModifiedBy UserInfo  `json:"modified_by"`
            Meta map[string]interface{}  `json:"meta"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            SortOn string  `json:"sort_on"`
            Published bool  `json:"published"`
            IsVisible bool  `json:"is_visible"`
            Logo CollectionImage  `json:"logo"`
         
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
    
    // ProductBrand used by Catalog
    type ProductBrand struct {

        
            Name string  `json:"name"`
            UID float64  `json:"uid"`
            Action Action  `json:"action"`
            Logo Media  `json:"logo"`
         
    }
    
    // Price1 used by Catalog
    type Price1 struct {

        
            CurrencyCode string  `json:"currency_code"`
            Max float64  `json:"max"`
            CurrencySymbol string  `json:"currency_symbol"`
            Min float64  `json:"min"`
         
    }
    
    // ProductListingPrice used by Catalog
    type ProductListingPrice struct {

        
            Effective Price1  `json:"effective"`
            Marked Price1  `json:"marked"`
         
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

        
            Brand ProductBrand  `json:"brand"`
            Color string  `json:"color"`
            Type string  `json:"type"`
            Slug string  `json:"slug"`
            ProductOnlineDate string  `json:"product_online_date"`
            ItemCode string  `json:"item_code"`
            Name string  `json:"name"`
            Description string  `json:"description"`
            Sellable bool  `json:"sellable"`
            Medias []Media  `json:"medias"`
            ItemType string  `json:"item_type"`
            Price ProductListingPrice  `json:"price"`
            Highlights []string  `json:"highlights"`
            Tryouts []string  `json:"tryouts"`
            Rating float64  `json:"rating"`
            Similars []string  `json:"similars"`
            RatingCount float64  `json:"rating_count"`
            UID float64  `json:"uid"`
            HasVariant bool  `json:"has_variant"`
            ImageNature string  `json:"image_nature"`
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
            Attributes map[string]interface{}  `json:"attributes"`
            PromoMeta map[string]interface{}  `json:"promo_meta"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            ShortDescription string  `json:"short_description"`
            Discount string  `json:"discount"`
         
    }
    
    // GetCollectionItemsResponse used by Catalog
    type GetCollectionItemsResponse struct {

        
            Items []ProductListingDetail  `json:"items"`
            Page Page  `json:"page"`
            Filters []ProductFilters  `json:"filters"`
            SortOn []ProductSortOn  `json:"sort_on"`
         
    }
    
    // CollectionItem used by Catalog
    type CollectionItem struct {

        
            ItemID float64  `json:"item_id"`
            Priority float64  `json:"priority"`
            Action string  `json:"action"`
         
    }
    
    // CollectionItemUpdate used by Catalog
    type CollectionItemUpdate struct {

        
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Items []CollectionItem  `json:"items"`
            AllowFacets bool  `json:"allow_facets"`
            Type string  `json:"type"`
            Query []CollectionQuery  `json:"query"`
            AllowSort bool  `json:"allow_sort"`
         
    }
    
    // UpdatedResponse used by Catalog
    type UpdatedResponse struct {

        
            Message string  `json:"message"`
            ItemsNotUpdated []float64  `json:"items_not_updated"`
         
    }
    
    // CatalogInsightBrand used by Catalog
    type CatalogInsightBrand struct {

        
            AvailableSizes float64  `json:"available_sizes"`
            ArticleFreshness float64  `json:"article_freshness"`
            Name string  `json:"name"`
            TotalSizes float64  `json:"total_sizes"`
            AvailableArticles float64  `json:"available_articles"`
            TotalArticles float64  `json:"total_articles"`
         
    }
    
    // CatalogInsightItem used by Catalog
    type CatalogInsightItem struct {

        
            SellableCount float64  `json:"sellable_count"`
            OutOfStockCount float64  `json:"out_of_stock_count"`
            Count float64  `json:"count"`
         
    }
    
    // CatalogInsightResponse used by Catalog
    type CatalogInsightResponse struct {

        
            BrandDistribution CatalogInsightBrand  `json:"brand_distribution"`
            Item CatalogInsightItem  `json:"item"`
         
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

        
            BrandIds []float64  `json:"brand_ids"`
            StoreIds []float64  `json:"store_ids"`
            Enabled bool  `json:"enabled"`
            Platform string  `json:"platform"`
            OptLevel string  `json:"opt_level"`
            CompanyID float64  `json:"company_id"`
         
    }
    
    // CompanyOptIn used by Catalog
    type CompanyOptIn struct {

        
            CreatedOn float64  `json:"created_on"`
            BrandIds []float64  `json:"brand_ids"`
            StoreIds []float64  `json:"store_ids"`
            Enabled bool  `json:"enabled"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Platform string  `json:"platform"`
            OptLevel string  `json:"opt_level"`
            ModifiedOn float64  `json:"modified_on"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            CompanyID float64  `json:"company_id"`
         
    }
    
    // GetOptInPlatform used by Catalog
    type GetOptInPlatform struct {

        
            Items []CompanyOptIn  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // OptinCompanyDetail used by Catalog
    type OptinCompanyDetail struct {

        
            BusinessType string  `json:"business_type"`
            UID float64  `json:"uid"`
            CompanyType string  `json:"company_type"`
            Name string  `json:"name"`
         
    }
    
    // CompanyBrandDetail used by Catalog
    type CompanyBrandDetail struct {

        
            BrandID float64  `json:"brand_id"`
            BrandName string  `json:"brand_name"`
            TotalArticle float64  `json:"total_article"`
            CompanyID float64  `json:"company_id"`
         
    }
    
    // OptinCompanyBrandDetailsView used by Catalog
    type OptinCompanyBrandDetailsView struct {

        
            Items []CompanyBrandDetail  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // OptinCompanyMetrics used by Catalog
    type OptinCompanyMetrics struct {

        
            Company string  `json:"company"`
            Store float64  `json:"store"`
            Brand float64  `json:"brand"`
         
    }
    
    // StoreDetail used by Catalog
    type StoreDetail struct {

        
            CreatedOn string  `json:"created_on"`
            StoreType string  `json:"store_type"`
            AdditionalContacts []map[string]interface{}  `json:"additional_contacts"`
            StoreCode string  `json:"store_code"`
            Name string  `json:"name"`
            DisplayName string  `json:"display_name"`
            UID float64  `json:"uid"`
            Address map[string]interface{}  `json:"address"`
            Manager map[string]interface{}  `json:"manager"`
            CompanyID float64  `json:"company_id"`
            ModifiedOn string  `json:"modified_on"`
            Timing map[string]interface{}  `json:"timing"`
            Documents []map[string]interface{}  `json:"documents"`
         
    }
    
    // OptinStoreDetails used by Catalog
    type OptinStoreDetails struct {

        
            Items []StoreDetail  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // AttributeSchemaRange used by Catalog
    type AttributeSchemaRange struct {

        
            Max float64  `json:"max"`
            Min float64  `json:"min"`
         
    }
    
    // AttributeMaster used by Catalog
    type AttributeMaster struct {

        
            AllowedValues []string  `json:"allowed_values"`
            Mandatory bool  `json:"mandatory"`
            Multi bool  `json:"multi"`
            Type string  `json:"type"`
            Format string  `json:"format"`
            Range AttributeSchemaRange  `json:"range"`
         
    }
    
    // AttributeMasterFilter used by Catalog
    type AttributeMasterFilter struct {

        
            DependsOn []string  `json:"depends_on"`
            Priority float64  `json:"priority"`
            Indexing bool  `json:"indexing"`
         
    }
    
    // AttributeMasterDetails used by Catalog
    type AttributeMasterDetails struct {

        
            DisplayType string  `json:"display_type"`
         
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
    
    // GenderDetail used by Catalog
    type GenderDetail struct {

        
            EnabledForEndConsumer bool  `json:"enabled_for_end_consumer"`
            Schema AttributeMaster  `json:"schema"`
            Departments []string  `json:"departments"`
            ID string  `json:"id"`
            Name string  `json:"name"`
            Filters AttributeMasterFilter  `json:"filters"`
            IsNested bool  `json:"is_nested"`
            Description string  `json:"description"`
            Slug string  `json:"slug"`
            Details AttributeMasterDetails  `json:"details"`
            Meta AttributeMasterMeta  `json:"meta"`
            Logo string  `json:"logo"`
         
    }
    
    // CategoriesResponse used by Catalog
    type CategoriesResponse struct {

        
            TemplateSlug string  `json:"template_slug"`
            SlugKey string  `json:"slug_key"`
            Name string  `json:"name"`
            UID float64  `json:"uid"`
            Slug string  `json:"slug"`
         
    }
    
    // ProdcutTemplateCategoriesResponse used by Catalog
    type ProdcutTemplateCategoriesResponse struct {

        
            Items []CategoriesResponse  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // PTErrorResponse used by Catalog
    type PTErrorResponse struct {

        
            Status float64  `json:"status"`
            Message string  `json:"message"`
            Errors map[string]interface{}  `json:"errors"`
            Meta map[string]interface{}  `json:"meta"`
            Code string  `json:"code"`
         
    }
    
    // UserSerializer used by Catalog
    type UserSerializer struct {

        
            ID string  `json:"_id"`
            Contact string  `json:"contact"`
            UID string  `json:"uid"`
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
         
    }
    
    // GetDepartment used by Catalog
    type GetDepartment struct {

        
            CreatedOn string  `json:"created_on"`
            IsActive bool  `json:"is_active"`
            PriorityOrder float64  `json:"priority_order"`
            Name string  `json:"name"`
            Synonyms []string  `json:"synonyms"`
            PageSize float64  `json:"page_size"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            UID float64  `json:"uid"`
            Search string  `json:"search"`
            Slug string  `json:"slug"`
            PageNo float64  `json:"page_no"`
            ModifiedOn string  `json:"modified_on"`
            ItemType string  `json:"item_type"`
            CreatedBy UserSerializer  `json:"created_by"`
            Logo string  `json:"logo"`
         
    }
    
    // DepartmentsResponse used by Catalog
    type DepartmentsResponse struct {

        
            Items []GetDepartment  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // DepartmentErrorResponse used by Catalog
    type DepartmentErrorResponse struct {

        
            Status float64  `json:"status"`
            Message string  `json:"message"`
            Errors map[string]interface{}  `json:"errors"`
            Meta map[string]interface{}  `json:"meta"`
            Code string  `json:"code"`
         
    }
    
    // DepartmentCreateUpdate used by Catalog
    type DepartmentCreateUpdate struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            PriorityOrder float64  `json:"priority_order"`
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            Cls string  `json:"_cls"`
            Synonyms []string  `json:"synonyms"`
            Tags []string  `json:"tags"`
            UID float64  `json:"uid"`
            Slug string  `json:"slug"`
            Platforms map[string]interface{}  `json:"platforms"`
            Logo string  `json:"logo"`
         
    }
    
    // DepartmentCreateResponse used by Catalog
    type DepartmentCreateResponse struct {

        
            UID float64  `json:"uid"`
            Message string  `json:"message"`
         
    }
    
    // DepartmentCreateErrorResponse used by Catalog
    type DepartmentCreateErrorResponse struct {

        
            Error string  `json:"error"`
         
    }
    
    // UserDetail used by Catalog
    type UserDetail struct {

        
            Username string  `json:"username"`
            SuperUser bool  `json:"super_user"`
            UserID string  `json:"user_id"`
            Contact string  `json:"contact"`
         
    }
    
    // DepartmentModel used by Catalog
    type DepartmentModel struct {

        
            CreatedOn string  `json:"created_on"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ID interface{}  `json:"_id"`
            VerifiedOn string  `json:"verified_on"`
            IsActive bool  `json:"is_active"`
            PriorityOrder float64  `json:"priority_order"`
            Name string  `json:"name"`
            Cls string  `json:"_cls"`
            VerifiedBy UserDetail  `json:"verified_by"`
            Synonyms []string  `json:"synonyms"`
            ModifiedBy UserDetail  `json:"modified_by"`
            UID float64  `json:"uid"`
            ModifiedOn string  `json:"modified_on"`
            Slug string  `json:"slug"`
            CreatedBy UserDetail  `json:"created_by"`
            Logo string  `json:"logo"`
         
    }
    
    // ProductTemplate used by Catalog
    type ProductTemplate struct {

        
            CreatedOn string  `json:"created_on"`
            Departments []string  `json:"departments"`
            IsExpirable bool  `json:"is_expirable"`
            Tag string  `json:"tag"`
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            IsArchived bool  `json:"is_archived"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Description string  `json:"description"`
            IsPhysical bool  `json:"is_physical"`
            Slug string  `json:"slug"`
            ModifiedOn string  `json:"modified_on"`
            Categories []string  `json:"categories"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Attributes []string  `json:"attributes"`
            Logo string  `json:"logo"`
         
    }
    
    // TemplatesResponse used by Catalog
    type TemplatesResponse struct {

        
            Items ProductTemplate  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // Properties used by Catalog
    type Properties struct {

        
            ReturnConfig map[string]interface{}  `json:"return_config"`
            Tags map[string]interface{}  `json:"tags"`
            CategorySlug map[string]interface{}  `json:"category_slug"`
            Slug map[string]interface{}  `json:"slug"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            Command map[string]interface{}  `json:"command"`
            ItemCode map[string]interface{}  `json:"item_code"`
            Name map[string]interface{}  `json:"name"`
            IsDependent map[string]interface{}  `json:"is_dependent"`
            ProductPublish map[string]interface{}  `json:"product_publish"`
            Description map[string]interface{}  `json:"description"`
            ItemType map[string]interface{}  `json:"item_type"`
            Currency map[string]interface{}  `json:"currency"`
            ProductGroupTag map[string]interface{}  `json:"product_group_tag"`
            Highlights map[string]interface{}  `json:"highlights"`
            NoOfBoxes map[string]interface{}  `json:"no_of_boxes"`
            BrandUID map[string]interface{}  `json:"brand_uid"`
            HsnCode map[string]interface{}  `json:"hsn_code"`
            IsActive map[string]interface{}  `json:"is_active"`
            CountryOfOrigin map[string]interface{}  `json:"country_of_origin"`
            MultiSize map[string]interface{}  `json:"multi_size"`
            Sizes map[string]interface{}  `json:"sizes"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            TraderType map[string]interface{}  `json:"trader_type"`
            ShortDescription map[string]interface{}  `json:"short_description"`
            Variants map[string]interface{}  `json:"variants"`
            Trader map[string]interface{}  `json:"trader"`
            Media map[string]interface{}  `json:"media"`
            SizeGuide map[string]interface{}  `json:"size_guide"`
         
    }
    
    // GlobalValidation used by Catalog
    type GlobalValidation struct {

        
            Required []string  `json:"required"`
            Title string  `json:"title"`
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
    
    // TemplateDetails used by Catalog
    type TemplateDetails struct {

        
            Departments []string  `json:"departments"`
            ID string  `json:"id"`
            IsExpirable bool  `json:"is_expirable"`
            Tag string  `json:"tag"`
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            IsArchived bool  `json:"is_archived"`
            Description string  `json:"description"`
            IsPhysical bool  `json:"is_physical"`
            Slug string  `json:"slug"`
            Categories []string  `json:"categories"`
            Attributes []string  `json:"attributes"`
            Logo string  `json:"logo"`
         
    }
    
    // TemplatesValidationResponse used by Catalog
    type TemplatesValidationResponse struct {

        
            Data TemplateValidationData  `json:"data"`
            TemplateDetails TemplateDetails  `json:"template_details"`
         
    }
    
    // InventoryValidationResponse used by Catalog
    type InventoryValidationResponse struct {

        
            Data map[string]interface{}  `json:"data"`
            Message string  `json:"message"`
         
    }
    
    // HSNData used by Catalog
    type HSNData struct {

        
            CountryOfOrigin []string  `json:"country_of_origin"`
            HsnCode []string  `json:"hsn_code"`
         
    }
    
    // HSNCodesResponse used by Catalog
    type HSNCodesResponse struct {

        
            Data HSNData  `json:"data"`
            Message string  `json:"message"`
         
    }
    
    // UserInfo1 used by Catalog
    type UserInfo1 struct {

        
            Email string  `json:"email"`
            UID string  `json:"uid"`
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
         
    }
    
    // ProductTemplateExportResponse used by Catalog
    type ProductTemplateExportResponse struct {

        
            CompletedOn string  `json:"completed_on"`
            Status string  `json:"status"`
            SellerID float64  `json:"seller_id"`
            Filters map[string]interface{}  `json:"filters"`
            Type string  `json:"type"`
            ModifiedOn string  `json:"modified_on"`
            TaskID string  `json:"task_id"`
            URL string  `json:"url"`
            CreatedBy UserInfo1  `json:"created_by"`
            NotificationEmails []string  `json:"notification_emails"`
         
    }
    
    // ProductDownloadsResponse used by Catalog
    type ProductDownloadsResponse struct {

        
            Items []ProductTemplateExportResponse  `json:"items"`
         
    }
    
    // ProductTemplateExportFilterRequest used by Catalog
    type ProductTemplateExportFilterRequest struct {

        
            CatalogueTypes []string  `json:"catalogue_types"`
            FromDate string  `json:"from_date"`
            Brands []string  `json:"brands"`
            Templates []string  `json:"templates"`
            ToDate string  `json:"to_date"`
         
    }
    
    // ProductTemplateDownloadsExport used by Catalog
    type ProductTemplateDownloadsExport struct {

        
            Type string  `json:"type"`
            Filters ProductTemplateExportFilterRequest  `json:"filters"`
            NotificationEmails []string  `json:"notification_emails"`
         
    }
    
    // ProductConfigurationDownloads used by Catalog
    type ProductConfigurationDownloads struct {

        
            Multivalue bool  `json:"multivalue"`
            Data []map[string]interface{}  `json:"data"`
         
    }
    
    // CategoryMappingValues used by Catalog
    type CategoryMappingValues struct {

        
            Name string  `json:"name"`
            CatalogID float64  `json:"catalog_id"`
         
    }
    
    // CategoryMapping used by Catalog
    type CategoryMapping struct {

        
            Ajio CategoryMappingValues  `json:"ajio"`
            Google CategoryMappingValues  `json:"google"`
            Facebook CategoryMappingValues  `json:"facebook"`
         
    }
    
    // Media1 used by Catalog
    type Media1 struct {

        
            Landscape string  `json:"landscape"`
            Portrait string  `json:"portrait"`
            Logo string  `json:"logo"`
         
    }
    
    // Hierarchy used by Catalog
    type Hierarchy struct {

        
            Department float64  `json:"department"`
            L2 float64  `json:"l2"`
            L1 float64  `json:"l1"`
         
    }
    
    // Category used by Catalog
    type Category struct {

        
            CreatedOn string  `json:"created_on"`
            Priority float64  `json:"priority"`
            Departments []float64  `json:"departments"`
            Tryouts []string  `json:"tryouts"`
            ID string  `json:"id"`
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            Synonyms []string  `json:"synonyms"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            UID float64  `json:"uid"`
            Slug string  `json:"slug"`
            Level float64  `json:"level"`
            ModifiedOn string  `json:"modified_on"`
            Marketplaces CategoryMapping  `json:"marketplaces"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Media Media1  `json:"media"`
            Hierarchy []Hierarchy  `json:"hierarchy"`
         
    }
    
    // CategoryResponse used by Catalog
    type CategoryResponse struct {

        
            Items []Category  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // CategoryRequestBody used by Catalog
    type CategoryRequestBody struct {

        
            Priority float64  `json:"priority"`
            Departments []float64  `json:"departments"`
            Tryouts []string  `json:"tryouts"`
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            Synonyms []string  `json:"synonyms"`
            Slug string  `json:"slug"`
            Level float64  `json:"level"`
            Marketplaces CategoryMapping  `json:"marketplaces"`
            Media Media1  `json:"media"`
            Hierarchy []Hierarchy  `json:"hierarchy"`
         
    }
    
    // CategoryCreateResponse used by Catalog
    type CategoryCreateResponse struct {

        
            UID float64  `json:"uid"`
            Message string  `json:"message"`
         
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
    
    // Image used by Catalog
    type Image struct {

        
            SecureURL string  `json:"secure_url"`
            AspectRatio string  `json:"aspect_ratio"`
            AspectRatioF float64  `json:"aspect_ratio_f"`
            URL string  `json:"url"`
         
    }
    
    // Logo used by Catalog
    type Logo struct {

        
            SecureURL string  `json:"secure_url"`
            AspectRatio string  `json:"aspect_ratio"`
            AspectRatioF float64  `json:"aspect_ratio_f"`
            URL string  `json:"url"`
         
    }
    
    // Brand used by Catalog
    type Brand struct {

        
            Name string  `json:"name"`
            UID float64  `json:"uid"`
            Logo Logo  `json:"logo"`
         
    }
    
    // ReturnConfigResponse used by Catalog
    type ReturnConfigResponse struct {

        
            Returnable bool  `json:"returnable"`
            Unit string  `json:"unit"`
            Time float64  `json:"time"`
         
    }
    
    // ProductPublish used by Catalog
    type ProductPublish struct {

        
            ProductOnlineDate string  `json:"product_online_date"`
            IsSet bool  `json:"is_set"`
         
    }
    
    // NetQuantityResponse used by Catalog
    type NetQuantityResponse struct {

        
            Value float64  `json:"value"`
            Unit string  `json:"unit"`
         
    }
    
    // VerifiedBy used by Catalog
    type VerifiedBy struct {

        
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
         
    }
    
    // Trader used by Catalog
    type Trader struct {

        
            Name string  `json:"name"`
            Type string  `json:"type"`
            Address []string  `json:"address"`
         
    }
    
    // ProductSchemaV2 used by Catalog
    type ProductSchemaV2 struct {

        
            Images []Image  `json:"images"`
            VerifiedOn string  `json:"verified_on"`
            Brand Brand  `json:"brand"`
            ReturnConfig ReturnConfigResponse  `json:"return_config"`
            Color string  `json:"color"`
            Tags []string  `json:"tags"`
            PrimaryColor string  `json:"primary_color"`
            CategorySlug string  `json:"category_slug"`
            IsPhysical bool  `json:"is_physical"`
            Slug string  `json:"slug"`
            AllIdentifiers []string  `json:"all_identifiers"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            ItemCode string  `json:"item_code"`
            TemplateTag string  `json:"template_tag"`
            CreatedOn string  `json:"created_on"`
            CompanyID float64  `json:"company_id"`
            Stage string  `json:"stage"`
            IsExpirable bool  `json:"is_expirable"`
            Name string  `json:"name"`
            IsDependent bool  `json:"is_dependent"`
            ProductPublish ProductPublish  `json:"product_publish"`
            L3Mapping []string  `json:"l3_mapping"`
            Description string  `json:"description"`
            AllSizes []map[string]interface{}  `json:"all_sizes"`
            NetQuantity NetQuantityResponse  `json:"net_quantity"`
            ItemType string  `json:"item_type"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            Currency string  `json:"currency"`
            ProductGroupTag []string  `json:"product_group_tag"`
            Highlights []string  `json:"highlights"`
            VariantMedia map[string]interface{}  `json:"variant_media"`
            NoOfBoxes float64  `json:"no_of_boxes"`
            BrandUID float64  `json:"brand_uid"`
            IsSet bool  `json:"is_set"`
            HsnCode string  `json:"hsn_code"`
            IsActive bool  `json:"is_active"`
            VerifiedBy VerifiedBy  `json:"verified_by"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            UID float64  `json:"uid"`
            Pending string  `json:"pending"`
            CountryOfOrigin string  `json:"country_of_origin"`
            MultiSize bool  `json:"multi_size"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ImageNature string  `json:"image_nature"`
            Attributes map[string]interface{}  `json:"attributes"`
            AllCompanyIds []float64  `json:"all_company_ids"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Sizes []map[string]interface{}  `json:"sizes"`
            VariantGroup map[string]interface{}  `json:"variant_group"`
            Departments []float64  `json:"departments"`
            ID string  `json:"id"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            ShortDescription string  `json:"short_description"`
            CategoryUID float64  `json:"category_uid"`
            Variants map[string]interface{}  `json:"variants"`
            Trader []Trader  `json:"trader"`
            Category map[string]interface{}  `json:"category"`
            ModifiedOn string  `json:"modified_on"`
            Moq map[string]interface{}  `json:"moq"`
            Media []Media  `json:"media"`
            SizeGuide string  `json:"size_guide"`
            IsImageLessProduct bool  `json:"is_image_less_product"`
         
    }
    
    // ProductListingResponseV2 used by Catalog
    type ProductListingResponseV2 struct {

        
            Items []ProductSchemaV2  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // ReturnConfig used by Catalog
    type ReturnConfig struct {

        
            Returnable bool  `json:"returnable"`
            Unit string  `json:"unit"`
            Time float64  `json:"time"`
         
    }
    
    // CustomOrder used by Catalog
    type CustomOrder struct {

        
            IsCustomOrder bool  `json:"is_custom_order"`
            ManufacturingTimeUnit string  `json:"manufacturing_time_unit"`
            ManufacturingTime float64  `json:"manufacturing_time"`
         
    }
    
    // ProductPublish1 used by Catalog
    type ProductPublish1 struct {

        
            ProductOnlineDate string  `json:"product_online_date"`
            IsSet bool  `json:"is_set"`
         
    }
    
    // NetQuantity used by Catalog
    type NetQuantity struct {

        
            Value float64  `json:"value"`
            Unit string  `json:"unit"`
         
    }
    
    // TaxIdentifier used by Catalog
    type TaxIdentifier struct {

        
            ReportingHsn string  `json:"reporting_hsn"`
            HsnCodeID string  `json:"hsn_code_id"`
            HsnCode string  `json:"hsn_code"`
         
    }
    
    // TeaserTag used by Catalog
    type TeaserTag struct {

        
            URL string  `json:"url"`
            Tag string  `json:"tag"`
         
    }
    
    // ProductCreateUpdateSchemaV2 used by Catalog
    type ProductCreateUpdateSchemaV2 struct {

        
            ReturnConfig ReturnConfig  `json:"return_config"`
            Tags []string  `json:"tags"`
            CategorySlug string  `json:"category_slug"`
            Slug string  `json:"slug"`
            CustomOrder CustomOrder  `json:"custom_order"`
            ItemCode string  `json:"item_code"`
            TemplateTag string  `json:"template_tag"`
            CompanyID float64  `json:"company_id"`
            Name string  `json:"name"`
            IsDependent bool  `json:"is_dependent"`
            ProductPublish ProductPublish1  `json:"product_publish"`
            Requester string  `json:"requester"`
            NetQuantity NetQuantity  `json:"net_quantity"`
            Description string  `json:"description"`
            ItemType string  `json:"item_type"`
            TaxIdentifier TaxIdentifier  `json:"tax_identifier"`
            Currency string  `json:"currency"`
            ProductGroupTag []string  `json:"product_group_tag"`
            VariantMedia map[string]interface{}  `json:"variant_media"`
            Highlights []string  `json:"highlights"`
            NoOfBoxes float64  `json:"no_of_boxes"`
            BrandUID float64  `json:"brand_uid"`
            IsSet bool  `json:"is_set"`
            IsActive bool  `json:"is_active"`
            UID float64  `json:"uid"`
            CountryOfOrigin string  `json:"country_of_origin"`
            MultiSize bool  `json:"multi_size"`
            Attributes map[string]interface{}  `json:"attributes"`
            BulkJobID string  `json:"bulk_job_id"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Sizes []map[string]interface{}  `json:"sizes"`
            VariantGroup map[string]interface{}  `json:"variant_group"`
            TeaserTag TeaserTag  `json:"teaser_tag"`
            Departments []float64  `json:"departments"`
            Action string  `json:"action"`
            ShortDescription string  `json:"short_description"`
            Variants map[string]interface{}  `json:"variants"`
            Trader []Trader  `json:"trader"`
            ChangeRequestID interface{}  `json:"change_request_id"`
            Media []Media  `json:"media"`
            SizeGuide string  `json:"size_guide"`
            IsImageLessProduct bool  `json:"is_image_less_product"`
         
    }
    
    // ProductVariants used by Catalog
    type ProductVariants struct {

        
            BrandUID float64  `json:"brand_uid"`
            Name string  `json:"name"`
            CategoryUID float64  `json:"category_uid"`
            UID float64  `json:"uid"`
            Media []Media  `json:"media"`
            ItemCode string  `json:"item_code"`
         
    }
    
    // ProductVariantsResponse used by Catalog
    type ProductVariantsResponse struct {

        
            Page Page  `json:"page"`
            Variants []ProductVariants  `json:"variants"`
         
    }
    
    // AttributeMasterSerializer used by Catalog
    type AttributeMasterSerializer struct {

        
            Synonyms map[string]interface{}  `json:"synonyms"`
            Tags []string  `json:"tags"`
            Slug string  `json:"slug"`
            Example string  `json:"example"`
            CreatedOn string  `json:"created_on"`
            Name string  `json:"name"`
            Filters AttributeMasterFilter  `json:"filters"`
            IsNested bool  `json:"is_nested"`
            Description string  `json:"description"`
            EnabledForEndConsumer bool  `json:"enabled_for_end_consumer"`
            RawKey string  `json:"raw_key"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Details AttributeMasterDetails  `json:"details"`
            Suggestion string  `json:"suggestion"`
            Departments []string  `json:"departments"`
            Schema AttributeMaster  `json:"schema"`
            Unit string  `json:"unit"`
            ModifiedOn string  `json:"modified_on"`
            Variant bool  `json:"variant"`
            Logo string  `json:"logo"`
         
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

        
            ItemWidth float64  `json:"item_width"`
            ItemWeightUnitOfMeasure string  `json:"item_weight_unit_of_measure"`
            Identifiers []ValidateIdentifier  `json:"identifiers"`
            ItemDimensionsUnitOfMeasure string  `json:"item_dimensions_unit_of_measure"`
            ItemLength float64  `json:"item_length"`
            ItemHeight float64  `json:"item_height"`
            ItemWeight float64  `json:"item_weight"`
            Size string  `json:"size"`
         
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

        
            ProductOnlineDate float64  `json:"product_online_date"`
            IsSet bool  `json:"is_set"`
         
    }
    
    // Product used by Catalog
    type Product struct {

        
            Images []Image  `json:"images"`
            VerifiedOn string  `json:"verified_on"`
            Brand Brand  `json:"brand"`
            ReturnConfig ReturnConfigResponse  `json:"return_config"`
            Color string  `json:"color"`
            Tags []string  `json:"tags"`
            PrimaryColor string  `json:"primary_color"`
            CategorySlug string  `json:"category_slug"`
            IsPhysical bool  `json:"is_physical"`
            Slug string  `json:"slug"`
            AllIdentifiers []string  `json:"all_identifiers"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            ItemCode string  `json:"item_code"`
            TemplateTag string  `json:"template_tag"`
            CreatedOn string  `json:"created_on"`
            CompanyID float64  `json:"company_id"`
            Stage string  `json:"stage"`
            IsExpirable bool  `json:"is_expirable"`
            Name string  `json:"name"`
            IsDependent bool  `json:"is_dependent"`
            ProductPublish ProductPublished  `json:"product_publish"`
            L3Mapping []string  `json:"l3_mapping"`
            Description string  `json:"description"`
            AllSizes []map[string]interface{}  `json:"all_sizes"`
            NetQuantity NetQuantityResponse  `json:"net_quantity"`
            ItemType string  `json:"item_type"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            Currency string  `json:"currency"`
            ProductGroupTag []string  `json:"product_group_tag"`
            Highlights []string  `json:"highlights"`
            VariantMedia map[string]interface{}  `json:"variant_media"`
            NoOfBoxes float64  `json:"no_of_boxes"`
            BrandUID float64  `json:"brand_uid"`
            IsSet bool  `json:"is_set"`
            HsnCode string  `json:"hsn_code"`
            IsActive bool  `json:"is_active"`
            VerifiedBy VerifiedBy  `json:"verified_by"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            UID float64  `json:"uid"`
            Pending string  `json:"pending"`
            CountryOfOrigin string  `json:"country_of_origin"`
            MultiSize bool  `json:"multi_size"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ImageNature string  `json:"image_nature"`
            Attributes map[string]interface{}  `json:"attributes"`
            AllCompanyIds []float64  `json:"all_company_ids"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Sizes []map[string]interface{}  `json:"sizes"`
            VariantGroup map[string]interface{}  `json:"variant_group"`
            Departments []float64  `json:"departments"`
            ID string  `json:"id"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            ShortDescription string  `json:"short_description"`
            CategoryUID float64  `json:"category_uid"`
            Variants map[string]interface{}  `json:"variants"`
            Trader []Trader  `json:"trader"`
            Category map[string]interface{}  `json:"category"`
            ModifiedOn string  `json:"modified_on"`
            Moq map[string]interface{}  `json:"moq"`
            Media []Media  `json:"media"`
            SizeGuide string  `json:"size_guide"`
            IsImageLessProduct bool  `json:"is_image_less_product"`
         
    }
    
    // ProductListingResponse used by Catalog
    type ProductListingResponse struct {

        
            Items []Product  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // UserDetail1 used by Catalog
    type UserDetail1 struct {

        
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
            FullName string  `json:"full_name"`
         
    }
    
    // ProductBulkRequest used by Catalog
    type ProductBulkRequest struct {

        
            CreatedOn string  `json:"created_on"`
            TemplateTag string  `json:"template_tag"`
            Template ProductTemplate  `json:"template"`
            FailedRecords []string  `json:"failed_records"`
            IsActive bool  `json:"is_active"`
            Total float64  `json:"total"`
            Stage string  `json:"stage"`
            Succeed float64  `json:"succeed"`
            ModifiedBy UserDetail1  `json:"modified_by"`
            FilePath string  `json:"file_path"`
            ModifiedOn string  `json:"modified_on"`
            Cancelled float64  `json:"cancelled"`
            Failed float64  `json:"failed"`
            CreatedBy UserDetail1  `json:"created_by"`
            CancelledRecords []string  `json:"cancelled_records"`
            CompanyID float64  `json:"company_id"`
         
    }
    
    // ProductBulkRequestList used by Catalog
    type ProductBulkRequestList struct {

        
            Items ProductBulkRequest  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // BulkJob used by Catalog
    type BulkJob struct {

        
            CreatedOn string  `json:"created_on"`
            TemplateTag string  `json:"template_tag"`
            TrackingURL string  `json:"tracking_url"`
            CustomTemplateTag string  `json:"custom_template_tag"`
            Total float64  `json:"total"`
            IsActive bool  `json:"is_active"`
            FailedRecords []map[string]interface{}  `json:"failed_records"`
            Succeed float64  `json:"succeed"`
            Stage string  `json:"stage"`
            ModifiedBy UserInfo1  `json:"modified_by"`
            FilePath string  `json:"file_path"`
            ModifiedOn string  `json:"modified_on"`
            Cancelled float64  `json:"cancelled"`
            Failed float64  `json:"failed"`
            CreatedBy UserInfo1  `json:"created_by"`
            CancelledRecords []map[string]interface{}  `json:"cancelled_records"`
            CompanyID float64  `json:"company_id"`
         
    }
    
    // BulkResponse used by Catalog
    type BulkResponse struct {

        
            CreatedOn string  `json:"created_on"`
            BatchID string  `json:"batch_id"`
            IsActive bool  `json:"is_active"`
            ModifiedBy UserInfo1  `json:"modified_by"`
            ModifiedOn string  `json:"modified_on"`
            CreatedBy UserInfo1  `json:"created_by"`
         
    }
    
    // BulkProductRequest used by Catalog
    type BulkProductRequest struct {

        
            TemplateTag string  `json:"template_tag"`
            Data []map[string]interface{}  `json:"data"`
            BatchID string  `json:"batch_id"`
            CompanyID float64  `json:"company_id"`
         
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

        
            CreatedOn string  `json:"created_on"`
            TrackingURL string  `json:"tracking_url"`
            ID string  `json:"id"`
            FailedRecords []string  `json:"failed_records"`
            IsActive bool  `json:"is_active"`
            Retry float64  `json:"retry"`
            Total float64  `json:"total"`
            Succeed float64  `json:"succeed"`
            ModifiedBy UserCommon  `json:"modified_by"`
            FilePath string  `json:"file_path"`
            CompanyID float64  `json:"company_id"`
            ModifiedOn string  `json:"modified_on"`
            Cancelled float64  `json:"cancelled"`
            Failed float64  `json:"failed"`
            CreatedBy UserCommon  `json:"created_by"`
            CancelledRecords []string  `json:"cancelled_records"`
            Stage string  `json:"stage"`
         
    }
    
    // BulkAssetResponse used by Catalog
    type BulkAssetResponse struct {

        
            Items []Items  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // ProductBulkAssets used by Catalog
    type ProductBulkAssets struct {

        
            User map[string]interface{}  `json:"user"`
            CompanyID float64  `json:"company_id"`
            BatchID string  `json:"batch_id"`
            URL string  `json:"url"`
         
    }
    
    // ProductSizeDeleteDataResponse used by Catalog
    type ProductSizeDeleteDataResponse struct {

        
            ItemID float64  `json:"item_id"`
            Size string  `json:"size"`
            CompanyID float64  `json:"company_id"`
         
    }
    
    // ProductSizeDeleteResponse used by Catalog
    type ProductSizeDeleteResponse struct {

        
            Data ProductSizeDeleteDataResponse  `json:"data"`
            Success bool  `json:"success"`
         
    }
    
    // InventoryResponse used by Catalog
    type InventoryResponse struct {

        
            SellableQuantity float64  `json:"sellable_quantity"`
            ItemID float64  `json:"item_id"`
            Identifiers map[string]interface{}  `json:"identifiers"`
            Quantity float64  `json:"quantity"`
            SellerIdentifier string  `json:"seller_identifier"`
            Store map[string]interface{}  `json:"store"`
            PriceTransfer float64  `json:"price_transfer"`
            PriceEffective float64  `json:"price_effective"`
            UID string  `json:"uid"`
            Currency string  `json:"currency"`
            InventoryUpdatedOn string  `json:"inventory_updated_on"`
            Price float64  `json:"price"`
            Size string  `json:"size"`
         
    }
    
    // InventoryResponsePaginated used by Catalog
    type InventoryResponsePaginated struct {

        
            Items []InventoryResponse  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // GTIN used by Catalog
    type GTIN struct {

        
            GtinValue string  `json:"gtin_value"`
            Primary bool  `json:"primary"`
            GtinType string  `json:"gtin_type"`
         
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

        
            SizeDistribution SizeDistribution  `json:"size_distribution"`
            Quantity float64  `json:"quantity"`
            Name string  `json:"name"`
         
    }
    
    // InvSize used by Catalog
    type InvSize struct {

        
            ExpirationDate string  `json:"expiration_date"`
            ItemWidth float64  `json:"item_width"`
            IsSet bool  `json:"is_set"`
            StoreCode string  `json:"store_code"`
            ItemWeightUnitOfMeasure string  `json:"item_weight_unit_of_measure"`
            Quantity float64  `json:"quantity"`
            ItemDimensionsUnitOfMeasure string  `json:"item_dimensions_unit_of_measure"`
            Identifiers []GTIN  `json:"identifiers"`
            Currency string  `json:"currency"`
            PriceTransfer float64  `json:"price_transfer"`
            Set InventorySet  `json:"set"`
            ItemHeight float64  `json:"item_height"`
            ItemLength float64  `json:"item_length"`
            ItemWeight float64  `json:"item_weight"`
            PriceEffective float64  `json:"price_effective"`
            Price float64  `json:"price"`
            Size string  `json:"size"`
         
    }
    
    // ItemQuery used by Catalog
    type ItemQuery struct {

        
            UID float64  `json:"uid"`
            BrandUID float64  `json:"brand_uid"`
            ItemCode string  `json:"item_code"`
         
    }
    
    // InventoryRequest used by Catalog
    type InventoryRequest struct {

        
            Sizes []InvSize  `json:"sizes"`
            Item ItemQuery  `json:"item"`
            CompanyID float64  `json:"company_id"`
         
    }
    
    // QuantityBase used by Catalog
    type QuantityBase struct {

        
            Count float64  `json:"count"`
            UpdatedAt string  `json:"updated_at"`
         
    }
    
    // Quantities used by Catalog
    type Quantities struct {

        
            OrderCommitted QuantityBase  `json:"order_committed"`
            NotAvailable QuantityBase  `json:"not_available"`
            Sellable QuantityBase  `json:"sellable"`
            Damaged QuantityBase  `json:"damaged"`
         
    }
    
    // BrandMeta used by Catalog
    type BrandMeta struct {

        
            Name string  `json:"name"`
            ID float64  `json:"id"`
         
    }
    
    // ReturnConfig1 used by Catalog
    type ReturnConfig1 struct {

        
            Returnable bool  `json:"returnable"`
            Unit string  `json:"unit"`
            Time float64  `json:"time"`
         
    }
    
    // WeightResponse used by Catalog
    type WeightResponse struct {

        
            IsDefault bool  `json:"is_default"`
            Shipping float64  `json:"shipping"`
            Unit string  `json:"unit"`
         
    }
    
    // CompanyMeta used by Catalog
    type CompanyMeta struct {

        
            ID float64  `json:"id"`
         
    }
    
    // PriceMeta used by Catalog
    type PriceMeta struct {

        
            Effective float64  `json:"effective"`
            UpdatedAt string  `json:"updated_at"`
            Marked float64  `json:"marked"`
            Currency string  `json:"currency"`
            Transfer float64  `json:"transfer"`
            TpNotes map[string]interface{}  `json:"tp_notes"`
         
    }
    
    // DimensionResponse used by Catalog
    type DimensionResponse struct {

        
            Unit string  `json:"unit"`
            Length float64  `json:"length"`
            IsDefault bool  `json:"is_default"`
            Height float64  `json:"height"`
            Width float64  `json:"width"`
         
    }
    
    // Trader1 used by Catalog
    type Trader1 struct {

        
            Name string  `json:"name"`
            Type string  `json:"type"`
            Address []string  `json:"address"`
         
    }
    
    // ManufacturerResponse used by Catalog
    type ManufacturerResponse struct {

        
            Name string  `json:"name"`
            Address string  `json:"address"`
            IsDefault bool  `json:"is_default"`
         
    }
    
    // StoreMeta used by Catalog
    type StoreMeta struct {

        
            ID float64  `json:"id"`
         
    }
    
    // InventorySellerResponse used by Catalog
    type InventorySellerResponse struct {

        
            Quantities Quantities  `json:"quantities"`
            Brand BrandMeta  `json:"brand"`
            ReturnConfig ReturnConfig1  `json:"return_config"`
            FyndItemCode string  `json:"fynd_item_code"`
            ItemID float64  `json:"item_id"`
            Tags []string  `json:"tags"`
            TrackInventory bool  `json:"track_inventory"`
            Stage string  `json:"stage"`
            ExpirationDate string  `json:"expiration_date"`
            Identifier map[string]interface{}  `json:"identifier"`
            Weight WeightResponse  `json:"weight"`
            Company CompanyMeta  `json:"company"`
            TraceID string  `json:"trace_id"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            Price PriceMeta  `json:"price"`
            IsSet bool  `json:"is_set"`
            IsActive bool  `json:"is_active"`
            RawMeta map[string]interface{}  `json:"raw_meta"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            UID string  `json:"uid"`
            CountryOfOrigin string  `json:"country_of_origin"`
            CreatedBy UserSerializer  `json:"created_by"`
            Meta map[string]interface{}  `json:"meta"`
            Size string  `json:"size"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            TotalQuantity float64  `json:"total_quantity"`
            AddedOnStore string  `json:"added_on_store"`
            FyndArticleCode string  `json:"fynd_article_code"`
            SellerIdentifier string  `json:"seller_identifier"`
            FyndMeta map[string]interface{}  `json:"fynd_meta"`
            Dimension DimensionResponse  `json:"dimension"`
            Set InventorySet  `json:"set"`
            Trader []Trader1  `json:"trader"`
            Manufacturer ManufacturerResponse  `json:"manufacturer"`
            Fragile bool  `json:"fragile"`
            Store StoreMeta  `json:"store"`
         
    }
    
    // InventorySellerIdentifierResponsePaginated used by Catalog
    type InventorySellerIdentifierResponsePaginated struct {

        
            Items []InventorySellerResponse  `json:"items"`
            Page Page  `json:"page"`
         
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
    
    // DateMeta used by Catalog
    type DateMeta struct {

        
            CreatedOn string  `json:"created_on"`
            InventoryUpdatedOn string  `json:"inventory_updated_on"`
            ModifiedOn string  `json:"modified_on"`
            AddedOnStore string  `json:"added_on_store"`
         
    }
    
    // BrandMeta1 used by Catalog
    type BrandMeta1 struct {

        
            Name string  `json:"name"`
            ID float64  `json:"id"`
         
    }
    
    // ReturnConfig2 used by Catalog
    type ReturnConfig2 struct {

        
            Returnable bool  `json:"returnable"`
            Unit string  `json:"unit"`
            Time float64  `json:"time"`
         
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
    
    // PriceArticle used by Catalog
    type PriceArticle struct {

        
            Effective float64  `json:"effective"`
            Marked float64  `json:"marked"`
            Currency string  `json:"currency"`
            Transfer float64  `json:"transfer"`
            TpNotes map[string]interface{}  `json:"tp_notes"`
         
    }
    
    // DimensionResponse1 used by Catalog
    type DimensionResponse1 struct {

        
            Width float64  `json:"width"`
            Length float64  `json:"length"`
            Unit string  `json:"unit"`
            Height float64  `json:"height"`
         
    }
    
    // Trader2 used by Catalog
    type Trader2 struct {

        
            Name string  `json:"name"`
            Type string  `json:"type"`
            Address []string  `json:"address"`
         
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
            StoreType string  `json:"store_type"`
            UID float64  `json:"uid"`
            StoreCode string  `json:"store_code"`
         
    }
    
    // GetInventories used by Catalog
    type GetInventories struct {

        
            Quantities QuantitiesArticle  `json:"quantities"`
            DateMeta DateMeta  `json:"date_meta"`
            Brand BrandMeta1  `json:"brand"`
            ReturnConfig ReturnConfig2  `json:"return_config"`
            ItemID float64  `json:"item_id"`
            Tags []string  `json:"tags"`
            TrackInventory bool  `json:"track_inventory"`
            InventoryUpdatedOn string  `json:"inventory_updated_on"`
            Platforms map[string]interface{}  `json:"platforms"`
            Stage string  `json:"stage"`
            ExpirationDate string  `json:"expiration_date"`
            Identifier map[string]interface{}  `json:"identifier"`
            Weight WeightResponse1  `json:"weight"`
            Company CompanyMeta1  `json:"company"`
            TraceID string  `json:"trace_id"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            Price PriceArticle  `json:"price"`
            IsSet bool  `json:"is_set"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            UID string  `json:"uid"`
            CountryOfOrigin string  `json:"country_of_origin"`
            CreatedBy UserSerializer  `json:"created_by"`
            Size string  `json:"size"`
            ID string  `json:"id"`
            TotalQuantity float64  `json:"total_quantity"`
            Dimension DimensionResponse1  `json:"dimension"`
            SellerIdentifier string  `json:"seller_identifier"`
            Trader []Trader2  `json:"trader"`
            Manufacturer ManufacturerResponse1  `json:"manufacturer"`
            Store ArticleStoreResponse  `json:"store"`
         
    }
    
    // GetInventoriesResponse used by Catalog
    type GetInventoriesResponse struct {

        
            Items []GetInventories  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // BulkInventoryGetItems used by Catalog
    type BulkInventoryGetItems struct {

        
            CreatedOn string  `json:"created_on"`
            ID string  `json:"id"`
            FailedRecords []string  `json:"failed_records"`
            IsActive bool  `json:"is_active"`
            Total float64  `json:"total"`
            Stage string  `json:"stage"`
            Succeed float64  `json:"succeed"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            FilePath string  `json:"file_path"`
            ModifiedOn string  `json:"modified_on"`
            Cancelled float64  `json:"cancelled"`
            Failed float64  `json:"failed"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            CancelledRecords []string  `json:"cancelled_records"`
            CompanyID float64  `json:"company_id"`
         
    }
    
    // BulkInventoryGet used by Catalog
    type BulkInventoryGet struct {

        
            Items []BulkInventoryGetItems  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // InventoryJobPayload used by Catalog
    type InventoryJobPayload struct {

        
            ExpirationDate string  `json:"expiration_date"`
            PriceMarked float64  `json:"price_marked"`
            StoreCode string  `json:"store_code"`
            TotalQuantity float64  `json:"total_quantity"`
            ItemWeightUnitOfMeasure string  `json:"item_weight_unit_of_measure"`
            Currency string  `json:"currency"`
            SellerIdentifier string  `json:"seller_identifier"`
            ItemDimensionsUnitOfMeasure string  `json:"item_dimensions_unit_of_measure"`
            Quantity float64  `json:"quantity"`
            Tags []string  `json:"tags"`
            TraceID string  `json:"trace_id"`
            PriceEffective float64  `json:"price_effective"`
            Price float64  `json:"price"`
         
    }
    
    // InventoryBulkRequest used by Catalog
    type InventoryBulkRequest struct {

        
            User map[string]interface{}  `json:"user"`
            Sizes []InventoryJobPayload  `json:"sizes"`
            BatchID string  `json:"batch_id"`
            CompanyID float64  `json:"company_id"`
         
    }
    
    // InventoryExportQuantityFilter used by Catalog
    type InventoryExportQuantityFilter struct {

        
            Max float64  `json:"max"`
            Min float64  `json:"min"`
            Operators string  `json:"operators"`
         
    }
    
    // InventoryExportAdvanceOption used by Catalog
    type InventoryExportAdvanceOption struct {

        
            BrandIds []float64  `json:"brand_ids"`
            StoreIds []float64  `json:"store_ids"`
            Quantity InventoryExportQuantityFilter  `json:"quantity"`
            FromDate string  `json:"from_date"`
            ToDate string  `json:"to_date"`
         
    }
    
    // InventoryExportJob used by Catalog
    type InventoryExportJob struct {

        
            CompletedOn string  `json:"completed_on"`
            Status string  `json:"status"`
            SellerID float64  `json:"seller_id"`
            Filters InventoryExportAdvanceOption  `json:"filters"`
            Type string  `json:"type"`
            TaskID string  `json:"task_id"`
            URL string  `json:"url"`
            NotificationEmails []string  `json:"notification_emails"`
         
    }
    
    // InventoryExportRequest used by Catalog
    type InventoryExportRequest struct {

        
            Type string  `json:"type"`
            Store []float64  `json:"store"`
            Brand []float64  `json:"brand"`
         
    }
    
    // InventoryExportResponse used by Catalog
    type InventoryExportResponse struct {

        
            CreatedOn string  `json:"created_on"`
            Status string  `json:"status"`
            SellerID float64  `json:"seller_id"`
            Filters map[string]interface{}  `json:"filters"`
            Type string  `json:"type"`
            ModifiedOn string  `json:"modified_on"`
            TaskID string  `json:"task_id"`
            CreatedBy string  `json:"created_by"`
            NotificationEmails []string  `json:"notification_emails"`
         
    }
    
    // InventoryJobFilters used by Catalog
    type InventoryJobFilters struct {

        
            Quantity InventoryExportQuantityFilter  `json:"quantity"`
            Stores []string  `json:"stores"`
            FromDate string  `json:"from_date"`
            Brands []string  `json:"brands"`
            ToDate string  `json:"to_date"`
         
    }
    
    // InventoryJobDetailResponse used by Catalog
    type InventoryJobDetailResponse struct {

        
            CreatedOn string  `json:"created_on"`
            CompletedOn string  `json:"completed_on"`
            ID string  `json:"id"`
            Status string  `json:"status"`
            SellerID float64  `json:"seller_id"`
            Filters InventoryJobFilters  `json:"filters"`
            Type string  `json:"type"`
            ModifiedOn string  `json:"modified_on"`
            TaskID string  `json:"task_id"`
            URL string  `json:"url"`
            CreatedBy UserDetail  `json:"created_by"`
            CancelledBy UserDetail  `json:"cancelled_by"`
            CancelledOn string  `json:"cancelled_on"`
            NotificationEmails []string  `json:"notification_emails"`
         
    }
    
    // InventoryExportJobListResponse used by Catalog
    type InventoryExportJobListResponse struct {

        
            Items InventoryJobDetailResponse  `json:"items"`
         
    }
    
    // InventoryExportFilter used by Catalog
    type InventoryExportFilter struct {

        
            BrandIds []float64  `json:"brand_ids"`
            StoreIds []float64  `json:"store_ids"`
            Quantity InventoryExportQuantityFilter  `json:"quantity"`
            FromDate string  `json:"from_date"`
            ToDate string  `json:"to_date"`
         
    }
    
    // InventoryCreateRequest used by Catalog
    type InventoryCreateRequest struct {

        
            Data []string  `json:"data"`
            Type string  `json:"type"`
            Filters InventoryExportFilter  `json:"filters"`
            NotificationEmails []string  `json:"notification_emails"`
         
    }
    
    // FilerList used by Catalog
    type FilerList struct {

        
            Value string  `json:"value"`
            Display string  `json:"display"`
         
    }
    
    // InventoryConfig used by Catalog
    type InventoryConfig struct {

        
            Data []FilerList  `json:"data"`
            Multivalues bool  `json:"multivalues"`
         
    }
    
    // InventoryFailedReason used by Catalog
    type InventoryFailedReason struct {

        
            Message string  `json:"message"`
            Errors string  `json:"errors"`
         
    }
    
    // InventoryPayload used by Catalog
    type InventoryPayload struct {

        
            ExpirationDate string  `json:"expiration_date"`
            PriceMarked float64  `json:"price_marked"`
            StoreID float64  `json:"store_id"`
            TotalQuantity float64  `json:"total_quantity"`
            SellerIdentifier string  `json:"seller_identifier"`
            Tags []string  `json:"tags"`
            TraceID string  `json:"trace_id"`
            PriceEffective float64  `json:"price_effective"`
         
    }
    
    // InventoryResponseItem used by Catalog
    type InventoryResponseItem struct {

        
            Reason InventoryFailedReason  `json:"reason"`
            Data InventoryPayload  `json:"data"`
         
    }
    
    // InventoryUpdateResponse used by Catalog
    type InventoryUpdateResponse struct {

        
            Items []InventoryResponseItem  `json:"items"`
            Message string  `json:"message"`
         
    }
    
    // InventoryRequestSchemaV2 used by Catalog
    type InventoryRequestSchemaV2 struct {

        
            Payload []InventoryPayload  `json:"payload"`
            Meta map[string]interface{}  `json:"meta"`
            CompanyID float64  `json:"company_id"`
         
    }
    
    // HsnCodesObject used by Catalog
    type HsnCodesObject struct {

        
            Tax2 float64  `json:"tax2"`
            Tax1 float64  `json:"tax1"`
            TaxOnEsp bool  `json:"tax_on_esp"`
            ID string  `json:"id"`
            Threshold1 float64  `json:"threshold1"`
            HsnCode string  `json:"hsn_code"`
            Threshold2 float64  `json:"threshold2"`
            ModifiedOn string  `json:"modified_on"`
            TaxOnMrp bool  `json:"tax_on_mrp"`
            Hs2Code string  `json:"hs2_code"`
            CompanyID float64  `json:"company_id"`
         
    }
    
    // HsnCode used by Catalog
    type HsnCode struct {

        
            Data HsnCodesObject  `json:"data"`
         
    }
    
    // HsnUpsert used by Catalog
    type HsnUpsert struct {

        
            Tax2 float64  `json:"tax2"`
            Tax1 float64  `json:"tax1"`
            TaxOnEsp bool  `json:"tax_on_esp"`
            Threshold1 float64  `json:"threshold1"`
            HsnCode string  `json:"hsn_code"`
            IsActive bool  `json:"is_active"`
            UID float64  `json:"uid"`
            Threshold2 float64  `json:"threshold2"`
            TaxOnMrp bool  `json:"tax_on_mrp"`
            Hs2Code string  `json:"hs2_code"`
            CompanyID float64  `json:"company_id"`
         
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

        
            Threshold float64  `json:"threshold"`
            EffectiveDate string  `json:"effective_date"`
            Cess float64  `json:"cess"`
            Rate float64  `json:"rate"`
         
    }
    
    // HSNDataInsertV2 used by Catalog
    type HSNDataInsertV2 struct {

        
            CreatedOn string  `json:"created_on"`
            HsnCode string  `json:"hsn_code"`
            CountryCode string  `json:"country_code"`
            HsnCodeID string  `json:"hsn_code_id"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Description string  `json:"description"`
            ReportingHsn string  `json:"reporting_hsn"`
            Type string  `json:"type"`
            ModifiedOn string  `json:"modified_on"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Taxes []TaxSlab  `json:"taxes"`
         
    }
    
    // PageResponse used by Catalog
    type PageResponse struct {

        
            ItemTotal float64  `json:"item_total"`
            HasNext bool  `json:"has_next"`
            Current string  `json:"current"`
            Size float64  `json:"size"`
            HasPrevious bool  `json:"has_previous"`
         
    }
    
    // HsnCodesListingResponseSchemaV2 used by Catalog
    type HsnCodesListingResponseSchemaV2 struct {

        
            Items []HSNDataInsertV2  `json:"items"`
            Page PageResponse  `json:"page"`
         
    }
    
    // Media2 used by Catalog
    type Media2 struct {

        
            Type string  `json:"type"`
            AspectRatio string  `json:"aspect_ratio"`
            URL string  `json:"url"`
         
    }
    
    // BrandItem used by Catalog
    type BrandItem struct {

        
            Departments []string  `json:"departments"`
            Action Action  `json:"action"`
            Name string  `json:"name"`
            Discount string  `json:"discount"`
            Banners ImageUrls  `json:"banners"`
            UID float64  `json:"uid"`
            Slug string  `json:"slug"`
            Logo Media2  `json:"logo"`
         
    }
    
    // BrandListingResponse used by Catalog
    type BrandListingResponse struct {

        
            Items []BrandItem  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // Department used by Catalog
    type Department struct {

        
            PriorityOrder float64  `json:"priority_order"`
            Name string  `json:"name"`
            UID float64  `json:"uid"`
            Slug string  `json:"slug"`
            Logo Media2  `json:"logo"`
         
    }
    
    // DepartmentResponse used by Catalog
    type DepartmentResponse struct {

        
            Items []Department  `json:"items"`
         
    }
    
    // ThirdLevelChild used by Catalog
    type ThirdLevelChild struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Childs []map[string]interface{}  `json:"childs"`
            Name string  `json:"name"`
            Action Action  `json:"action"`
            Banners ImageUrls  `json:"banners"`
            UID float64  `json:"uid"`
            Slug string  `json:"slug"`
         
    }
    
    // SecondLevelChild used by Catalog
    type SecondLevelChild struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Childs []ThirdLevelChild  `json:"childs"`
            Name string  `json:"name"`
            Action Action  `json:"action"`
            Banners ImageUrls  `json:"banners"`
            UID float64  `json:"uid"`
            Slug string  `json:"slug"`
         
    }
    
    // Child used by Catalog
    type Child struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Childs []SecondLevelChild  `json:"childs"`
            Name string  `json:"name"`
            Action Action  `json:"action"`
            Banners ImageUrls  `json:"banners"`
            UID float64  `json:"uid"`
            Slug string  `json:"slug"`
         
    }
    
    // CategoryItems used by Catalog
    type CategoryItems struct {

        
            Childs []Child  `json:"childs"`
            Action Action  `json:"action"`
            Name string  `json:"name"`
            Banners ImageUrls  `json:"banners"`
            UID float64  `json:"uid"`
            Slug string  `json:"slug"`
         
    }
    
    // DepartmentCategoryTree used by Catalog
    type DepartmentCategoryTree struct {

        
            Department string  `json:"department"`
            Items []CategoryItems  `json:"items"`
         
    }
    
    // DepartmentIdentifier used by Catalog
    type DepartmentIdentifier struct {

        
            UID float64  `json:"uid"`
            Slug string  `json:"slug"`
         
    }
    
    // CategoryListingResponse used by Catalog
    type CategoryListingResponse struct {

        
            Data []DepartmentCategoryTree  `json:"data"`
            Departments []DepartmentIdentifier  `json:"departments"`
         
    }
    
    // ApplicationProductListingResponse used by Catalog
    type ApplicationProductListingResponse struct {

        
            SortOn []ProductSortOn  `json:"sort_on"`
            Filters []ProductFilters  `json:"filters"`
            Items []ProductListingDetail  `json:"items"`
            Page Page  `json:"page"`
            Operators map[string]interface{}  `json:"operators"`
         
    }
    
    // ProductDetail used by Catalog
    type ProductDetail struct {

        
            Brand ProductBrand  `json:"brand"`
            Color string  `json:"color"`
            Type string  `json:"type"`
            Slug string  `json:"slug"`
            ProductOnlineDate string  `json:"product_online_date"`
            ItemCode string  `json:"item_code"`
            Name string  `json:"name"`
            Description string  `json:"description"`
            Medias []Media  `json:"medias"`
            ItemType string  `json:"item_type"`
            Highlights []string  `json:"highlights"`
            Tryouts []string  `json:"tryouts"`
            Rating float64  `json:"rating"`
            Similars []string  `json:"similars"`
            RatingCount float64  `json:"rating_count"`
            UID float64  `json:"uid"`
            HasVariant bool  `json:"has_variant"`
            ImageNature string  `json:"image_nature"`
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
            Attributes map[string]interface{}  `json:"attributes"`
            PromoMeta map[string]interface{}  `json:"promo_meta"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            ShortDescription string  `json:"short_description"`
         
    }
    
    // InventoryPage used by Catalog
    type InventoryPage struct {

        
            ItemTotal float64  `json:"item_total"`
            HasNext bool  `json:"has_next"`
            Type string  `json:"type"`
            NextID string  `json:"next_id"`
            HasPrevious bool  `json:"has_previous"`
         
    }
    
    // InventoryStockResponse used by Catalog
    type InventoryStockResponse struct {

        
            Items []map[string]interface{}  `json:"items"`
            Page InventoryPage  `json:"page"`
         
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

        
            GroupID string  `json:"group_id"`
            Quantity float64  `json:"quantity"`
            Query ArticleQuery  `json:"query"`
            ArticleAssignment ArticleAssignment  `json:"article_assignment"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // AssignStore used by Catalog
    type AssignStore struct {

        
            StoreIds []float64  `json:"store_ids"`
            ChannelType string  `json:"channel_type"`
            AppID string  `json:"app_id"`
            ChannelIdentifier string  `json:"channel_identifier"`
            Articles []AssignStoreArticle  `json:"articles"`
            Pincode string  `json:"pincode"`
            CompanyID float64  `json:"company_id"`
         
    }
    
    // ArticleAssignment1 used by Catalog
    type ArticleAssignment1 struct {

        
            Strategy string  `json:"strategy"`
            Level string  `json:"level"`
         
    }
    
    // StoreAssignResponse used by Catalog
    type StoreAssignResponse struct {

        
            GroupID string  `json:"group_id"`
            ID string  `json:"_id"`
            PriceMarked float64  `json:"price_marked"`
            StoreID float64  `json:"store_id"`
            Status bool  `json:"status"`
            StrategyWiseListing []map[string]interface{}  `json:"strategy_wise_listing"`
            ItemID float64  `json:"item_id"`
            Quantity float64  `json:"quantity"`
            UID string  `json:"uid"`
            PriceEffective float64  `json:"price_effective"`
            StorePincode float64  `json:"store_pincode"`
            ArticleAssignment ArticleAssignment1  `json:"article_assignment"`
            Index float64  `json:"index"`
            SCity string  `json:"s_city"`
            Meta map[string]interface{}  `json:"meta"`
            Size string  `json:"size"`
            CompanyID float64  `json:"company_id"`
         
    }
    
    // GetAddressSerializer used by Catalog
    type GetAddressSerializer struct {

        
            City string  `json:"city"`
            Latitude float64  `json:"latitude"`
            Country string  `json:"country"`
            Landmark string  `json:"landmark"`
            CountryCode string  `json:"country_code"`
            Address2 string  `json:"address2"`
            AddressType string  `json:"address_type"`
            Longitude float64  `json:"longitude"`
            State string  `json:"state"`
            Pincode float64  `json:"pincode"`
            Address1 string  `json:"address1"`
         
    }
    
    // LocationTimingSerializer used by Catalog
    type LocationTimingSerializer struct {

        
            Minute float64  `json:"minute"`
            Hour float64  `json:"hour"`
         
    }
    
    // LocationDayWiseSerializer used by Catalog
    type LocationDayWiseSerializer struct {

        
            Open bool  `json:"open"`
            Weekday string  `json:"weekday"`
            Opening LocationTimingSerializer  `json:"opening"`
            Closing LocationTimingSerializer  `json:"closing"`
         
    }
    
    // SellerPhoneNumber used by Catalog
    type SellerPhoneNumber struct {

        
            CountryCode float64  `json:"country_code"`
            Number string  `json:"number"`
         
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
    
    // UserSerializer1 used by Catalog
    type UserSerializer1 struct {

        
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
            Contact string  `json:"contact"`
         
    }
    
    // GetCompanySerializer used by Catalog
    type GetCompanySerializer struct {

        
            CreatedOn string  `json:"created_on"`
            VerifiedOn string  `json:"verified_on"`
            BusinessType string  `json:"business_type"`
            Name string  `json:"name"`
            CompanyType string  `json:"company_type"`
            VerifiedBy UserSerializer1  `json:"verified_by"`
            ModifiedBy UserSerializer1  `json:"modified_by"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            UID float64  `json:"uid"`
            ModifiedOn string  `json:"modified_on"`
            RejectReason string  `json:"reject_reason"`
            CreatedBy UserSerializer1  `json:"created_by"`
            Stage string  `json:"stage"`
         
    }
    
    // LocationIntegrationType used by Catalog
    type LocationIntegrationType struct {

        
            Order string  `json:"order"`
            Inventory string  `json:"inventory"`
         
    }
    
    // UserSerializer2 used by Catalog
    type UserSerializer2 struct {

        
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
            Contact string  `json:"contact"`
         
    }
    
    // ProductReturnConfigSerializer used by Catalog
    type ProductReturnConfigSerializer struct {

        
            StoreUID float64  `json:"store_uid"`
            OnSameStore bool  `json:"on_same_store"`
         
    }
    
    // LocationManagerSerializer used by Catalog
    type LocationManagerSerializer struct {

        
            Email string  `json:"email"`
            Name string  `json:"name"`
            MobileNo SellerPhoneNumber  `json:"mobile_no"`
         
    }
    
    // Document used by Catalog
    type Document struct {

        
            Value string  `json:"value"`
            LegalName string  `json:"legal_name"`
            Verified bool  `json:"verified"`
            Type string  `json:"type"`
            URL string  `json:"url"`
         
    }
    
    // GetLocationSerializer used by Catalog
    type GetLocationSerializer struct {

        
            VerifiedOn string  `json:"verified_on"`
            Address GetAddressSerializer  `json:"address"`
            Timing []LocationDayWiseSerializer  `json:"timing"`
            Stage string  `json:"stage"`
            CreatedOn string  `json:"created_on"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            Code string  `json:"code"`
            PhoneNumber string  `json:"phone_number"`
            Name string  `json:"name"`
            GstCredentials InvoiceDetailsSerializer  `json:"gst_credentials"`
            Company GetCompanySerializer  `json:"company"`
            IntegrationType LocationIntegrationType  `json:"integration_type"`
            Warnings map[string]interface{}  `json:"warnings"`
            VerifiedBy UserSerializer2  `json:"verified_by"`
            ModifiedBy UserSerializer2  `json:"modified_by"`
            DisplayName string  `json:"display_name"`
            UID float64  `json:"uid"`
            CreatedBy UserSerializer2  `json:"created_by"`
            NotificationEmails []string  `json:"notification_emails"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            StoreType string  `json:"store_type"`
            ProductReturnConfig ProductReturnConfigSerializer  `json:"product_return_config"`
            Manager LocationManagerSerializer  `json:"manager"`
            ModifiedOn string  `json:"modified_on"`
            Documents []Document  `json:"documents"`
         
    }
    
    // LocationListSerializer used by Catalog
    type LocationListSerializer struct {

        
            Items []GetLocationSerializer  `json:"items"`
            Page Page  `json:"page"`
         
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

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            AppID string  `json:"app_id"`
            UID float64  `json:"uid"`
            Logo string  `json:"logo"`
         
    }
    
    // ApplicationDepartmentListingResponse used by Catalog
    type ApplicationDepartmentListingResponse struct {

        
            Items []ApplicationDepartment  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // ApplicationDepartmentJson used by Catalog
    type ApplicationDepartmentJson struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // ApplicationStoreJson used by Catalog
    type ApplicationStoreJson struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    

    
    // SellerPhoneNumber used by CompanyProfile
    type SellerPhoneNumber struct {

        
            Number string  `json:"number"`
            CountryCode float64  `json:"country_code"`
         
    }
    
    // ContactDetails used by CompanyProfile
    type ContactDetails struct {

        
            Emails []string  `json:"emails"`
            Phone []SellerPhoneNumber  `json:"phone"`
         
    }
    
    // BusinessCountryInfo used by CompanyProfile
    type BusinessCountryInfo struct {

        
            CountryCode string  `json:"country_code"`
            Country string  `json:"country"`
         
    }
    
    // UserSerializer used by CompanyProfile
    type UserSerializer struct {

        
            Contact string  `json:"contact"`
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
         
    }
    
    // Document used by CompanyProfile
    type Document struct {

        
            URL string  `json:"url"`
            Verified bool  `json:"verified"`
            Type string  `json:"type"`
            LegalName string  `json:"legal_name"`
            Value string  `json:"value"`
         
    }
    
    // GetAddressSerializer used by CompanyProfile
    type GetAddressSerializer struct {

        
            Address2 string  `json:"address2"`
            CountryCode string  `json:"country_code"`
            Country string  `json:"country"`
            Pincode float64  `json:"pincode"`
            City string  `json:"city"`
            Landmark string  `json:"landmark"`
            Latitude float64  `json:"latitude"`
            AddressType string  `json:"address_type"`
            State string  `json:"state"`
            Longitude float64  `json:"longitude"`
            Address1 string  `json:"address1"`
         
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

        
            Enable bool  `json:"enable"`
            Rate float64  `json:"rate"`
            EffectiveDate string  `json:"effective_date"`
         
    }
    
    // GetCompanyProfileSerializerResponse used by CompanyProfile
    type GetCompanyProfileSerializerResponse struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ContactDetails ContactDetails  `json:"contact_details"`
            BusinessInfo string  `json:"business_info"`
            BusinessCountryInfo BusinessCountryInfo  `json:"business_country_info"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            Mode string  `json:"mode"`
            CompanyType string  `json:"company_type"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            CreatedBy UserSerializer  `json:"created_by"`
            UID float64  `json:"uid"`
            Stage string  `json:"stage"`
            Warnings map[string]interface{}  `json:"warnings"`
            FranchiseEnabled bool  `json:"franchise_enabled"`
            Name string  `json:"name"`
            Documents []Document  `json:"documents"`
            CreatedOn string  `json:"created_on"`
            VerifiedOn string  `json:"verified_on"`
            BusinessType string  `json:"business_type"`
            NotificationEmails []string  `json:"notification_emails"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            BusinessDetails BusinessDetails  `json:"business_details"`
            Taxes []CompanyTaxesSerializer  `json:"taxes"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // ErrorResponse used by CompanyProfile
    type ErrorResponse struct {

        
            Status float64  `json:"status"`
            Code string  `json:"code"`
            Meta map[string]interface{}  `json:"meta"`
            Message string  `json:"message"`
         
    }
    
    // CreateUpdateAddressSerializer used by CompanyProfile
    type CreateUpdateAddressSerializer struct {

        
            CountryCode string  `json:"country_code"`
            Country string  `json:"country"`
            Pincode float64  `json:"pincode"`
            City string  `json:"city"`
            Landmark string  `json:"landmark"`
            Latitude float64  `json:"latitude"`
            AddressType string  `json:"address_type"`
            Address2 string  `json:"address2"`
            State string  `json:"state"`
            Longitude float64  `json:"longitude"`
            Address1 string  `json:"address1"`
         
    }
    
    // CompanyTaxesSerializer1 used by CompanyProfile
    type CompanyTaxesSerializer1 struct {

        
            Enable bool  `json:"enable"`
            Rate float64  `json:"rate"`
            EffectiveDate string  `json:"effective_date"`
         
    }
    
    // UpdateCompany used by CompanyProfile
    type UpdateCompany struct {

        
            CompanyType string  `json:"company_type"`
            Warnings map[string]interface{}  `json:"warnings"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ContactDetails ContactDetails  `json:"contact_details"`
            BusinessType string  `json:"business_type"`
            BusinessInfo string  `json:"business_info"`
            Name string  `json:"name"`
            Addresses []CreateUpdateAddressSerializer  `json:"addresses"`
            NotificationEmails []string  `json:"notification_emails"`
            BusinessDetails BusinessDetails  `json:"business_details"`
            FranchiseEnabled bool  `json:"franchise_enabled"`
            Taxes []CompanyTaxesSerializer1  `json:"taxes"`
            Documents []Document  `json:"documents"`
            Slug string  `json:"slug"`
            RejectReason string  `json:"reject_reason"`
         
    }
    
    // ProfileSuccessResponse used by CompanyProfile
    type ProfileSuccessResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
            UID float64  `json:"uid"`
         
    }
    
    // DocumentsObj used by CompanyProfile
    type DocumentsObj struct {

        
            Verified float64  `json:"verified"`
            Pending float64  `json:"pending"`
         
    }
    
    // MetricsSerializer used by CompanyProfile
    type MetricsSerializer struct {

        
            Brand DocumentsObj  `json:"brand"`
            UID float64  `json:"uid"`
            Product DocumentsObj  `json:"product"`
            CompanyDocuments DocumentsObj  `json:"company_documents"`
            StoreDocuments DocumentsObj  `json:"store_documents"`
            Store DocumentsObj  `json:"store"`
            Stage string  `json:"stage"`
         
    }
    
    // BrandBannerSerializer used by CompanyProfile
    type BrandBannerSerializer struct {

        
            Landscape string  `json:"landscape"`
            Portrait string  `json:"portrait"`
         
    }
    
    // GetBrandResponseSerializer used by CompanyProfile
    type GetBrandResponseSerializer struct {

        
            Synonyms []string  `json:"synonyms"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            Banner BrandBannerSerializer  `json:"banner"`
            Mode string  `json:"mode"`
            RejectReason string  `json:"reject_reason"`
            SlugKey string  `json:"slug_key"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            CreatedBy UserSerializer  `json:"created_by"`
            Description string  `json:"description"`
            ModifiedOn string  `json:"modified_on"`
            UID float64  `json:"uid"`
            Stage string  `json:"stage"`
            Warnings map[string]interface{}  `json:"warnings"`
            Logo string  `json:"logo"`
            Name string  `json:"name"`
            CreatedOn string  `json:"created_on"`
            VerifiedOn string  `json:"verified_on"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
         
    }
    
    // CreateUpdateBrandRequestSerializer used by CompanyProfile
    type CreateUpdateBrandRequestSerializer struct {

        
            Synonyms []string  `json:"synonyms"`
            CompanyID float64  `json:"company_id"`
            Description string  `json:"description"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Logo string  `json:"logo"`
            Name string  `json:"name"`
            Banner BrandBannerSerializer  `json:"banner"`
            BrandTier string  `json:"brand_tier"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            UID float64  `json:"uid"`
         
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

        
            CompanyType string  `json:"company_type"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            UID float64  `json:"uid"`
            CreatedBy UserSerializer  `json:"created_by"`
            VerifiedOn string  `json:"verified_on"`
            Stage string  `json:"stage"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            BusinessType string  `json:"business_type"`
            MarketChannels []string  `json:"market_channels"`
            Name string  `json:"name"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            Details CompanyDetails  `json:"details"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            NotificationEmails []string  `json:"notification_emails"`
            BusinessCountryInfo BusinessCountryInfo  `json:"business_country_info"`
            ModifiedOn string  `json:"modified_on"`
            RejectReason string  `json:"reject_reason"`
            CreatedOn string  `json:"created_on"`
         
    }
    
    // CompanyBrandSerializer used by CompanyProfile
    type CompanyBrandSerializer struct {

        
            Brand GetBrandResponseSerializer  `json:"brand"`
            Company CompanySerializer  `json:"company"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            UID float64  `json:"uid"`
            CreatedBy UserSerializer  `json:"created_by"`
            VerifiedOn string  `json:"verified_on"`
            Stage string  `json:"stage"`
            Warnings map[string]interface{}  `json:"warnings"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            ModifiedOn string  `json:"modified_on"`
            RejectReason string  `json:"reject_reason"`
            CreatedOn string  `json:"created_on"`
         
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
    
    // GetCompanySerializer used by CompanyProfile
    type GetCompanySerializer struct {

        
            CompanyType string  `json:"company_type"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            UID float64  `json:"uid"`
            CreatedBy UserSerializer  `json:"created_by"`
            VerifiedOn string  `json:"verified_on"`
            Stage string  `json:"stage"`
            BusinessType string  `json:"business_type"`
            Name string  `json:"name"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            ModifiedOn string  `json:"modified_on"`
            RejectReason string  `json:"reject_reason"`
            CreatedOn string  `json:"created_on"`
         
    }
    
    // ProductReturnConfigSerializer used by CompanyProfile
    type ProductReturnConfigSerializer struct {

        
            OnSameStore bool  `json:"on_same_store"`
            StoreUID float64  `json:"store_uid"`
         
    }
    
    // LocationManagerSerializer used by CompanyProfile
    type LocationManagerSerializer struct {

        
            Name string  `json:"name"`
            Email string  `json:"email"`
            MobileNo SellerPhoneNumber  `json:"mobile_no"`
         
    }
    
    // LocationTimingSerializer used by CompanyProfile
    type LocationTimingSerializer struct {

        
            Hour float64  `json:"hour"`
            Minute float64  `json:"minute"`
         
    }
    
    // LocationDayWiseSerializer used by CompanyProfile
    type LocationDayWiseSerializer struct {

        
            Opening LocationTimingSerializer  `json:"opening"`
            Open bool  `json:"open"`
            Closing LocationTimingSerializer  `json:"closing"`
            Weekday string  `json:"weekday"`
         
    }
    
    // InvoiceCredSerializer used by CompanyProfile
    type InvoiceCredSerializer struct {

        
            Password string  `json:"password"`
            Username string  `json:"username"`
            Enabled bool  `json:"enabled"`
         
    }
    
    // InvoiceDetailsSerializer used by CompanyProfile
    type InvoiceDetailsSerializer struct {

        
            EWaybill InvoiceCredSerializer  `json:"e_waybill"`
            EInvoice InvoiceCredSerializer  `json:"e_invoice"`
         
    }
    
    // HolidayDateSerializer used by CompanyProfile
    type HolidayDateSerializer struct {

        
            EndDate string  `json:"end_date"`
            StartDate string  `json:"start_date"`
         
    }
    
    // HolidaySchemaSerializer used by CompanyProfile
    type HolidaySchemaSerializer struct {

        
            HolidayType string  `json:"holiday_type"`
            Title string  `json:"title"`
            Date HolidayDateSerializer  `json:"date"`
         
    }
    
    // GetLocationSerializer used by CompanyProfile
    type GetLocationSerializer struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            AutoInvoice bool  `json:"auto_invoice"`
            Company GetCompanySerializer  `json:"company"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            CreatedBy UserSerializer  `json:"created_by"`
            ProductReturnConfig ProductReturnConfigSerializer  `json:"product_return_config"`
            CreditNote bool  `json:"credit_note"`
            UID float64  `json:"uid"`
            Manager LocationManagerSerializer  `json:"manager"`
            Stage string  `json:"stage"`
            Warnings map[string]interface{}  `json:"warnings"`
            Timing []LocationDayWiseSerializer  `json:"timing"`
            Code string  `json:"code"`
            DisplayName string  `json:"display_name"`
            Name string  `json:"name"`
            PhoneNumber string  `json:"phone_number"`
            GstCredentials InvoiceDetailsSerializer  `json:"gst_credentials"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            Documents []Document  `json:"documents"`
            CreatedOn string  `json:"created_on"`
            VerifiedOn string  `json:"verified_on"`
            Address GetAddressSerializer  `json:"address"`
            Holiday []HolidaySchemaSerializer  `json:"holiday"`
            StoreType string  `json:"store_type"`
            NotificationEmails []string  `json:"notification_emails"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // LocationListSerializer used by CompanyProfile
    type LocationListSerializer struct {

        
            Items []GetLocationSerializer  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // AddressSerializer used by CompanyProfile
    type AddressSerializer struct {

        
            Address2 string  `json:"address2"`
            CountryCode string  `json:"country_code"`
            Country string  `json:"country"`
            Pincode float64  `json:"pincode"`
            City string  `json:"city"`
            Landmark string  `json:"landmark"`
            Latitude float64  `json:"latitude"`
            AddressType string  `json:"address_type"`
            State string  `json:"state"`
            Longitude float64  `json:"longitude"`
            Address1 string  `json:"address1"`
         
    }
    
    // LocationSerializer used by CompanyProfile
    type LocationSerializer struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            AutoInvoice bool  `json:"auto_invoice"`
            Company float64  `json:"company"`
            ProductReturnConfig ProductReturnConfigSerializer  `json:"product_return_config"`
            CreditNote bool  `json:"credit_note"`
            Stage string  `json:"stage"`
            Manager LocationManagerSerializer  `json:"manager"`
            UID float64  `json:"uid"`
            Warnings map[string]interface{}  `json:"warnings"`
            Timing []LocationDayWiseSerializer  `json:"timing"`
            Code string  `json:"code"`
            DisplayName string  `json:"display_name"`
            Name string  `json:"name"`
            GstCredentials InvoiceDetailsSerializer  `json:"gst_credentials"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            Documents []Document  `json:"documents"`
            Address AddressSerializer  `json:"address"`
            Holiday []HolidaySchemaSerializer  `json:"holiday"`
            StoreType string  `json:"store_type"`
            NotificationEmails []string  `json:"notification_emails"`
            Slug string  `json:"slug"`
         
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
    

    
    // ClickStatsResponse used by Share
    type ClickStatsResponse struct {

        
            ClickStats []ClickStatsItem  `json:"click_stats"`
         
    }
    
    // ClickStatsItem used by Share
    type ClickStatsItem struct {

        
            Display string  `json:"display"`
            Total float64  `json:"total"`
         
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
            Communication CommunicationConfig  `json:"communication"`
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
            AnonymousCod bool  `json:"anonymous_cod"`
         
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
            Communication CommunicationConfig  `json:"communication"`
         
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
            CodThresholdAmount float64  `json:"cod_threshold_amount"`
            OnlineThresholdAmount float64  `json:"online_threshold_amount"`
         
    }
    
    // ValidationFailedResponse used by Configuration
    type ValidationFailedResponse struct {

        
            Message string  `json:"message"`
         
    }
    
    // NotFound used by Configuration
    type NotFound struct {

        
            Message string  `json:"message"`
         
    }
    
    // CommunicationConfig used by Configuration
    type CommunicationConfig struct {

        
            Email CommsConfig  `json:"email"`
            Sms CommsConfig  `json:"sms"`
            Voice CommsConfig  `json:"voice"`
         
    }
    
    // CommsConfig used by Configuration
    type CommsConfig struct {

        
            Enabled bool  `json:"enabled"`
         
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
            ID string  `json:"_id"`
            Verified bool  `json:"verified"`
            IsPrimary bool  `json:"is_primary"`
            IsShortlink bool  `json:"is_shortlink"`
            Message string  `json:"message"`
            TxtRecords []string  `json:"txt_records"`
         
    }
    
    // DomainAddRequest used by Configuration
    type DomainAddRequest struct {

        
            Domain DomainAdd  `json:"domain"`
         
    }
    
    // Domain used by Configuration
    type Domain struct {

        
            Name string  `json:"name"`
            ID string  `json:"_id"`
            Verified bool  `json:"verified"`
            IsPrimary bool  `json:"is_primary"`
            IsShortlink bool  `json:"is_shortlink"`
            IsPredefined bool  `json:"is_predefined"`
         
    }
    
    // DomainsResponse used by Configuration
    type DomainsResponse struct {

        
            Domains []Domain  `json:"domains"`
         
    }
    
    // UpdateDomain used by Configuration
    type UpdateDomain struct {

        
            Name string  `json:"name"`
            ID string  `json:"_id"`
            Verified bool  `json:"verified"`
            IsPrimary bool  `json:"is_primary"`
            IsShortlink bool  `json:"is_shortlink"`
         
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
    
    // SuccessMessageResponse used by Configuration
    type SuccessMessageResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
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
            Sort string  `json:"sort"`
         
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
    
    // UnhandledError used by Configuration
    type UnhandledError struct {

        
            Message string  `json:"message"`
         
    }
    
    // InvalidPayloadRequest used by Configuration
    type InvalidPayloadRequest struct {

        
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
    
    // State used by Cart
    type State struct {

        
            IsPublic bool  `json:"is_public"`
            IsDisplay bool  `json:"is_display"`
            IsArchived bool  `json:"is_archived"`
         
    }
    
    // Validation used by Cart
    type Validation struct {

        
            Anonymous bool  `json:"anonymous"`
            AppID []string  `json:"app_id"`
            UserRegisteredAfter string  `json:"user_registered_after"`
         
    }
    
    // CouponAction used by Cart
    type CouponAction struct {

        
            TxnMode string  `json:"txn_mode"`
            ActionDate string  `json:"action_date"`
         
    }
    
    // CouponSchedule used by Cart
    type CouponSchedule struct {

        
            NextSchedule []map[string]interface{}  `json:"next_schedule"`
            Start string  `json:"start"`
            Cron string  `json:"cron"`
            End string  `json:"end"`
            Duration float64  `json:"duration"`
         
    }
    
    // RuleDefinition used by Cart
    type RuleDefinition struct {

        
            CurrencyCode string  `json:"currency_code"`
            AutoApply bool  `json:"auto_apply"`
            Type string  `json:"type"`
            IsExact bool  `json:"is_exact"`
            Scope []string  `json:"scope"`
            CalculateOn string  `json:"calculate_on"`
            ApplicableOn string  `json:"applicable_on"`
            ValueType string  `json:"value_type"`
         
    }
    
    // Rule used by Cart
    type Rule struct {

        
            Min float64  `json:"min"`
            Value float64  `json:"value"`
            DiscountQty float64  `json:"discount_qty"`
            Key float64  `json:"key"`
            Max float64  `json:"max"`
         
    }
    
    // CouponDateMeta used by Cart
    type CouponDateMeta struct {

        
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
         
    }
    
    // Ownership used by Cart
    type Ownership struct {

        
            PayableCategory string  `json:"payable_category"`
            PayableBy string  `json:"payable_by"`
         
    }
    
    // CouponAuthor used by Cart
    type CouponAuthor struct {

        
            CreatedBy string  `json:"created_by"`
            ModifiedBy string  `json:"modified_by"`
         
    }
    
    // Validity used by Cart
    type Validity struct {

        
            Priority float64  `json:"priority"`
         
    }
    
    // Identifier used by Cart
    type Identifier struct {

        
            EmailDomain []string  `json:"email_domain"`
            ArticleID []string  `json:"article_id"`
            UserID []string  `json:"user_id"`
            CollectionID []string  `json:"collection_id"`
            ExcludeBrandID []float64  `json:"exclude_brand_id"`
            CategoryID []float64  `json:"category_id"`
            StoreID []float64  `json:"store_id"`
            ItemID []float64  `json:"item_id"`
            CompanyID []float64  `json:"company_id"`
            BrandID []float64  `json:"brand_id"`
         
    }
    
    // PaymentAllowValue used by Cart
    type PaymentAllowValue struct {

        
            Max float64  `json:"max"`
         
    }
    
    // PaymentModes used by Cart
    type PaymentModes struct {

        
            Uses PaymentAllowValue  `json:"uses"`
            Codes []string  `json:"codes"`
            Types []string  `json:"types"`
            Iins []string  `json:"iins"`
            Networks []string  `json:"networks"`
         
    }
    
    // PriceRange used by Cart
    type PriceRange struct {

        
            Min float64  `json:"min"`
            Max float64  `json:"max"`
         
    }
    
    // UsesRemaining used by Cart
    type UsesRemaining struct {

        
            App float64  `json:"app"`
            User float64  `json:"user"`
            Total float64  `json:"total"`
         
    }
    
    // UsesRestriction used by Cart
    type UsesRestriction struct {

        
            Remaining UsesRemaining  `json:"remaining"`
            Maximum UsesRemaining  `json:"maximum"`
         
    }
    
    // BulkBundleRestriction used by Cart
    type BulkBundleRestriction struct {

        
            MultiStoreAllowed bool  `json:"multi_store_allowed"`
         
    }
    
    // PostOrder used by Cart
    type PostOrder struct {

        
            ReturnAllowed bool  `json:"return_allowed"`
            CancellationAllowed bool  `json:"cancellation_allowed"`
         
    }
    
    // Restrictions used by Cart
    type Restrictions struct {

        
            Payments map[string]PaymentModes  `json:"payments"`
            PriceRange PriceRange  `json:"price_range"`
            CouponAllowed bool  `json:"coupon_allowed"`
            Platforms []string  `json:"platforms"`
            Uses UsesRestriction  `json:"uses"`
            UserType string  `json:"user_type"`
            OrderingStores []float64  `json:"ordering_stores"`
            BulkBundle BulkBundleRestriction  `json:"bulk_bundle"`
            UserGroups []float64  `json:"user_groups"`
            PostOrder PostOrder  `json:"post_order"`
         
    }
    
    // CouponAdd used by Cart
    type CouponAdd struct {

        
            DisplayMeta DisplayMeta  `json:"display_meta"`
            State State  `json:"state"`
            Validation Validation  `json:"validation"`
            TypeSlug string  `json:"type_slug"`
            Code string  `json:"code"`
            Action CouponAction  `json:"action"`
            Schedule CouponSchedule  `json:"_schedule"`
            RuleDefinition RuleDefinition  `json:"rule_definition"`
            Rule []Rule  `json:"rule"`
            DateMeta CouponDateMeta  `json:"date_meta"`
            Ownership Ownership  `json:"ownership"`
            Author CouponAuthor  `json:"author"`
            Validity Validity  `json:"validity"`
            Identifiers Identifier  `json:"identifiers"`
            Restrictions Restrictions  `json:"restrictions"`
            Tags []string  `json:"tags"`
         
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

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // OperationErrorResponse used by Cart
    type OperationErrorResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // CouponUpdate used by Cart
    type CouponUpdate struct {

        
            DisplayMeta DisplayMeta  `json:"display_meta"`
            State State  `json:"state"`
            Validation Validation  `json:"validation"`
            TypeSlug string  `json:"type_slug"`
            Code string  `json:"code"`
            Action CouponAction  `json:"action"`
            Schedule CouponSchedule  `json:"_schedule"`
            RuleDefinition RuleDefinition  `json:"rule_definition"`
            Rule []Rule  `json:"rule"`
            DateMeta CouponDateMeta  `json:"date_meta"`
            Ownership Ownership  `json:"ownership"`
            Author CouponAuthor  `json:"author"`
            Validity Validity  `json:"validity"`
            Identifiers Identifier  `json:"identifiers"`
            Restrictions Restrictions  `json:"restrictions"`
            Tags []string  `json:"tags"`
         
    }
    
    // CouponPartialUpdate used by Cart
    type CouponPartialUpdate struct {

        
            Schedule CouponSchedule  `json:"schedule"`
            Archive bool  `json:"archive"`
         
    }
    
    // PromotionSchedule used by Cart
    type PromotionSchedule struct {

        
            NextSchedule []map[string]interface{}  `json:"next_schedule"`
            Published bool  `json:"published"`
            Start string  `json:"start"`
            Cron string  `json:"cron"`
            End string  `json:"end"`
            Duration float64  `json:"duration"`
         
    }
    
    // Ownership1 used by Cart
    type Ownership1 struct {

        
            PayableCategory string  `json:"payable_category"`
            PayableBy string  `json:"payable_by"`
         
    }
    
    // DiscountOffer used by Cart
    type DiscountOffer struct {

        
            MaxUsagePerTransaction float64  `json:"max_usage_per_transaction"`
            Code string  `json:"code"`
            MaxDiscountAmount float64  `json:"max_discount_amount"`
            PartialCanRet bool  `json:"partial_can_ret"`
            MinOfferQuantity float64  `json:"min_offer_quantity"`
            ApportionDiscount bool  `json:"apportion_discount"`
            DiscountPercentage float64  `json:"discount_percentage"`
            MaxOfferQuantity float64  `json:"max_offer_quantity"`
            DiscountPrice float64  `json:"discount_price"`
            DiscountAmount float64  `json:"discount_amount"`
         
    }
    
    // CompareObject used by Cart
    type CompareObject struct {

        
            GreaterThanEquals float64  `json:"greater_than_equals"`
            LessThan float64  `json:"less_than"`
            LessThanEquals float64  `json:"less_than_equals"`
            GreaterThan float64  `json:"greater_than"`
            Equals float64  `json:"equals"`
         
    }
    
    // ItemCriteria used by Cart
    type ItemCriteria struct {

        
            CartUniqueItemAmount CompareObject  `json:"cart_unique_item_amount"`
            CartTotal CompareObject  `json:"cart_total"`
            ItemExcludeL1Category []float64  `json:"item_exclude_l1_category"`
            ItemExcludeDepartment []float64  `json:"item_exclude_department"`
            ItemDepartment []float64  `json:"item_department"`
            ProductTags []string  `json:"product_tags"`
            AllItems bool  `json:"all_items"`
            ItemTags []string  `json:"item_tags"`
            ItemCompany []float64  `json:"item_company"`
            ItemStore []float64  `json:"item_store"`
            ItemSize []string  `json:"item_size"`
            ItemID []float64  `json:"item_id"`
            ItemExcludeCategory []float64  `json:"item_exclude_category"`
            ItemExcludeID []float64  `json:"item_exclude_id"`
            ItemL1Category []float64  `json:"item_l1_category"`
            BuyRules []string  `json:"buy_rules"`
            AvailableZones []string  `json:"available_zones"`
            CartUniqueItemQuantity CompareObject  `json:"cart_unique_item_quantity"`
            ItemExcludeSku []string  `json:"item_exclude_sku"`
            CartQuantity CompareObject  `json:"cart_quantity"`
            ItemExcludeL2Category []float64  `json:"item_exclude_l2_category"`
            ItemBrand []float64  `json:"item_brand"`
            ItemCategory []float64  `json:"item_category"`
            ItemExcludeStore []float64  `json:"item_exclude_store"`
            ItemL2Category []float64  `json:"item_l2_category"`
            ItemExcludeCompany []float64  `json:"item_exclude_company"`
            ItemExcludeBrand []float64  `json:"item_exclude_brand"`
            ItemSku []string  `json:"item_sku"`
         
    }
    
    // DiscountRule used by Cart
    type DiscountRule struct {

        
            Offer DiscountOffer  `json:"offer"`
            ItemCriteria ItemCriteria  `json:"item_criteria"`
            DiscountType string  `json:"discount_type"`
            BuyCondition string  `json:"buy_condition"`
         
    }
    
    // PromotionAction used by Cart
    type PromotionAction struct {

        
            ActionType string  `json:"action_type"`
            ActionDate string  `json:"action_date"`
         
    }
    
    // Visibility used by Cart
    type Visibility struct {

        
            Pdp bool  `json:"pdp"`
            CouponList bool  `json:"coupon_list"`
         
    }
    
    // DisplayMeta1 used by Cart
    type DisplayMeta1 struct {

        
            Name string  `json:"name"`
            Description string  `json:"description"`
            OfferLabel string  `json:"offer_label"`
            OfferText string  `json:"offer_text"`
         
    }
    
    // PromotionDateMeta used by Cart
    type PromotionDateMeta struct {

        
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
         
    }
    
    // PromotionAuthor used by Cart
    type PromotionAuthor struct {

        
            CreatedBy string  `json:"created_by"`
            ModifiedBy string  `json:"modified_by"`
         
    }
    
    // PaymentAllowValue1 used by Cart
    type PaymentAllowValue1 struct {

        
            Max float64  `json:"max"`
         
    }
    
    // PromotionPaymentModes used by Cart
    type PromotionPaymentModes struct {

        
            Codes []string  `json:"codes"`
            Type string  `json:"type"`
            Uses PaymentAllowValue1  `json:"uses"`
         
    }
    
    // UserRegistered used by Cart
    type UserRegistered struct {

        
            End string  `json:"end"`
            Start string  `json:"start"`
         
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
    
    // PostOrder1 used by Cart
    type PostOrder1 struct {

        
            ReturnAllowed bool  `json:"return_allowed"`
            CancellationAllowed bool  `json:"cancellation_allowed"`
         
    }
    
    // Restrictions1 used by Cart
    type Restrictions1 struct {

        
            Payments []PromotionPaymentModes  `json:"payments"`
            UserID []string  `json:"user_id"`
            Platforms []string  `json:"platforms"`
            UserRegistered UserRegistered  `json:"user_registered"`
            Uses UsesRestriction1  `json:"uses"`
            AnonymousUsers bool  `json:"anonymous_users"`
            OrderingStores []float64  `json:"ordering_stores"`
            OrderQuantity float64  `json:"order_quantity"`
            UserGroups []float64  `json:"user_groups"`
            PostOrder PostOrder1  `json:"post_order"`
         
    }
    
    // PromotionListItem used by Cart
    type PromotionListItem struct {

        
            PromoGroup string  `json:"promo_group"`
            Schedule PromotionSchedule  `json:"_schedule"`
            Code string  `json:"code"`
            Ownership Ownership1  `json:"ownership"`
            Stackable bool  `json:"stackable"`
            DiscountRules []DiscountRule  `json:"discount_rules"`
            PromotionType string  `json:"promotion_type"`
            ApplyPriority float64  `json:"apply_priority"`
            PostOrderAction PromotionAction  `json:"post_order_action"`
            Mode string  `json:"mode"`
            ApplyExclusive string  `json:"apply_exclusive"`
            BuyRules map[string]ItemCriteria  `json:"buy_rules"`
            Visiblility Visibility  `json:"visiblility"`
            DisplayMeta DisplayMeta1  `json:"display_meta"`
            Currency string  `json:"currency"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            DateMeta PromotionDateMeta  `json:"date_meta"`
            ApplicationID string  `json:"application_id"`
            ApplyAllDiscount bool  `json:"apply_all_discount"`
            Author PromotionAuthor  `json:"author"`
            CalculateOn string  `json:"calculate_on"`
            Restrictions Restrictions1  `json:"restrictions"`
         
    }
    
    // PromotionsResponse used by Cart
    type PromotionsResponse struct {

        
            Items PromotionListItem  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // PromotionAdd used by Cart
    type PromotionAdd struct {

        
            PromoGroup string  `json:"promo_group"`
            Schedule PromotionSchedule  `json:"_schedule"`
            Code string  `json:"code"`
            Ownership Ownership1  `json:"ownership"`
            Stackable bool  `json:"stackable"`
            DiscountRules []DiscountRule  `json:"discount_rules"`
            PromotionType string  `json:"promotion_type"`
            ApplyPriority float64  `json:"apply_priority"`
            PostOrderAction PromotionAction  `json:"post_order_action"`
            Mode string  `json:"mode"`
            ApplyExclusive string  `json:"apply_exclusive"`
            BuyRules map[string]ItemCriteria  `json:"buy_rules"`
            Visiblility Visibility  `json:"visiblility"`
            DisplayMeta DisplayMeta1  `json:"display_meta"`
            Currency string  `json:"currency"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            DateMeta PromotionDateMeta  `json:"date_meta"`
            ApplicationID string  `json:"application_id"`
            ApplyAllDiscount bool  `json:"apply_all_discount"`
            Author PromotionAuthor  `json:"author"`
            CalculateOn string  `json:"calculate_on"`
            Restrictions Restrictions1  `json:"restrictions"`
         
    }
    
    // PromotionUpdate used by Cart
    type PromotionUpdate struct {

        
            PromoGroup string  `json:"promo_group"`
            Schedule PromotionSchedule  `json:"_schedule"`
            Code string  `json:"code"`
            Ownership Ownership1  `json:"ownership"`
            Stackable bool  `json:"stackable"`
            DiscountRules []DiscountRule  `json:"discount_rules"`
            PromotionType string  `json:"promotion_type"`
            ApplyPriority float64  `json:"apply_priority"`
            PostOrderAction PromotionAction  `json:"post_order_action"`
            Mode string  `json:"mode"`
            ApplyExclusive string  `json:"apply_exclusive"`
            BuyRules map[string]ItemCriteria  `json:"buy_rules"`
            Visiblility Visibility  `json:"visiblility"`
            DisplayMeta DisplayMeta1  `json:"display_meta"`
            Currency string  `json:"currency"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            DateMeta PromotionDateMeta  `json:"date_meta"`
            ApplicationID string  `json:"application_id"`
            ApplyAllDiscount bool  `json:"apply_all_discount"`
            Author PromotionAuthor  `json:"author"`
            CalculateOn string  `json:"calculate_on"`
            Restrictions Restrictions1  `json:"restrictions"`
         
    }
    
    // PromotionPartialUpdate used by Cart
    type PromotionPartialUpdate struct {

        
            Schedule PromotionSchedule  `json:"schedule"`
            Archive bool  `json:"archive"`
         
    }
    
    // ActivePromosResponse used by Cart
    type ActivePromosResponse struct {

        
            IsHidden bool  `json:"is_hidden"`
            EntityType string  `json:"entity_type"`
            EntitySlug string  `json:"entity_slug"`
            Subtitle string  `json:"subtitle"`
            Title string  `json:"title"`
            Description string  `json:"description"`
            CreatedOn string  `json:"created_on"`
            Example string  `json:"example"`
            ModifiedOn string  `json:"modified_on"`
            Type string  `json:"type"`
         
    }
    
    // Charges used by Cart
    type Charges struct {

        
            Threshold float64  `json:"threshold"`
            Charges float64  `json:"charges"`
         
    }
    
    // DeliveryCharges used by Cart
    type DeliveryCharges struct {

        
            Enabled bool  `json:"enabled"`
            Charges []Charges  `json:"charges"`
         
    }
    
    // CartMetaConfigUpdate used by Cart
    type CartMetaConfigUpdate struct {

        
            GiftDisplayText string  `json:"gift_display_text"`
            Enabled bool  `json:"enabled"`
            BulkCoupons bool  `json:"bulk_coupons"`
            MinCartValue float64  `json:"min_cart_value"`
            GiftPricing float64  `json:"gift_pricing"`
            RevenueEngineCoupon bool  `json:"revenue_engine_coupon"`
            DeliveryCharges DeliveryCharges  `json:"delivery_charges"`
            MaxCartItems float64  `json:"max_cart_items"`
         
    }
    
    // CartMetaConfigAdd used by Cart
    type CartMetaConfigAdd struct {

        
            GiftDisplayText string  `json:"gift_display_text"`
            Enabled bool  `json:"enabled"`
            BulkCoupons bool  `json:"bulk_coupons"`
            MinCartValue float64  `json:"min_cart_value"`
            GiftPricing float64  `json:"gift_pricing"`
            RevenueEngineCoupon bool  `json:"revenue_engine_coupon"`
            DeliveryCharges DeliveryCharges  `json:"delivery_charges"`
            MaxCartItems float64  `json:"max_cart_items"`
         
    }
    
    // Article used by Cart
    type Article struct {

        
            ArticleID string  `json:"article_id"`
            Value float64  `json:"value"`
            Code string  `json:"code"`
            Meta map[string]interface{}  `json:"meta"`
            Type string  `json:"type"`
         
    }
    
    // Collecttion used by Cart
    type Collecttion struct {

        
            CollectedBy string  `json:"collected_by"`
            RefundBy string  `json:"refund_by"`
         
    }
    
    // CartDynamicInjectionUpdate used by Cart
    type CartDynamicInjectionUpdate struct {

        
            AllowedRefund bool  `json:"allowed_refund"`
            CartValue float64  `json:"cart_value"`
            Value float64  `json:"value"`
            Type string  `json:"type"`
            UserID string  `json:"user_id"`
            ArticleIds []Article  `json:"article_ids"`
            Collection Collecttion  `json:"collection"`
            ArticleLevelDistribution bool  `json:"article_level_distribution"`
            IsAuthenticated bool  `json:"is_authenticated"`
            ApplyExpiry string  `json:"apply_expiry"`
            Meta map[string]interface{}  `json:"meta"`
            Message string  `json:"message"`
            CartID string  `json:"cart_id"`
         
    }
    
    // CartDynamicInjection used by Cart
    type CartDynamicInjection struct {

        
            AllowedRefund bool  `json:"allowed_refund"`
            CartValue float64  `json:"cart_value"`
            Value float64  `json:"value"`
            Type string  `json:"type"`
            UserID string  `json:"user_id"`
            ArticleIds []Article  `json:"article_ids"`
            CartID string  `json:"cart_id"`
            Collection Collecttion  `json:"collection"`
            ArticleLevelDistribution bool  `json:"article_level_distribution"`
            ApplyExpiry string  `json:"apply_expiry"`
            IsAuthenticated bool  `json:"is_authenticated"`
            Meta map[string]interface{}  `json:"meta"`
            Message string  `json:"message"`
            InjectionID string  `json:"injection_id"`
         
    }
    
    // CartDynamicInjectionResponse used by Cart
    type CartDynamicInjectionResponse struct {

        
            Data CartDynamicInjection  `json:"data"`
         
    }
    
    // CartDynamicInjectionAdd used by Cart
    type CartDynamicInjectionAdd struct {

        
            AllowedRefund bool  `json:"allowed_refund"`
            CartValue float64  `json:"cart_value"`
            Value float64  `json:"value"`
            Type string  `json:"type"`
            UserID string  `json:"user_id"`
            ArticleIds []Article  `json:"article_ids"`
            Collection Collecttion  `json:"collection"`
            ArticleLevelDistribution bool  `json:"article_level_distribution"`
            IsAuthenticated bool  `json:"is_authenticated"`
            ApplyExpiry string  `json:"apply_expiry"`
            Meta map[string]interface{}  `json:"meta"`
            Message string  `json:"message"`
            CartID string  `json:"cart_id"`
         
    }
    
    // CartItem used by Cart
    type CartItem struct {

        
            Quantity float64  `json:"quantity"`
            ProductID string  `json:"product_id"`
            Size string  `json:"size"`
         
    }
    
    // OpenapiCartDetailsRequest used by Cart
    type OpenapiCartDetailsRequest struct {

        
            CartItems []CartItem  `json:"cart_items"`
         
    }
    
    // RawBreakup used by Cart
    type RawBreakup struct {

        
            Coupon float64  `json:"coupon"`
            CodCharge float64  `json:"cod_charge"`
            ConvenienceFee float64  `json:"convenience_fee"`
            MrpTotal float64  `json:"mrp_total"`
            GstCharges float64  `json:"gst_charges"`
            Subtotal float64  `json:"subtotal"`
            DeliveryCharge float64  `json:"delivery_charge"`
            GiftCard float64  `json:"gift_card"`
            Total float64  `json:"total"`
            Discount float64  `json:"discount"`
            YouSaved float64  `json:"you_saved"`
            Vog float64  `json:"vog"`
            FyndCash float64  `json:"fynd_cash"`
         
    }
    
    // DisplayBreakup used by Cart
    type DisplayBreakup struct {

        
            CurrencyCode string  `json:"currency_code"`
            Value float64  `json:"value"`
            CurrencySymbol string  `json:"currency_symbol"`
            Display string  `json:"display"`
            Message []string  `json:"message"`
            Key string  `json:"key"`
         
    }
    
    // CouponBreakup used by Cart
    type CouponBreakup struct {

        
            CouponValue float64  `json:"coupon_value"`
            MaxDiscountValue float64  `json:"max_discount_value"`
            Value float64  `json:"value"`
            Code string  `json:"code"`
            Title string  `json:"title"`
            CouponType string  `json:"coupon_type"`
            UID string  `json:"uid"`
            MinimumCartValue float64  `json:"minimum_cart_value"`
            Description string  `json:"description"`
            SubTitle string  `json:"sub_title"`
            Message string  `json:"message"`
            Type string  `json:"type"`
            IsApplied bool  `json:"is_applied"`
         
    }
    
    // LoyaltyPoints used by Cart
    type LoyaltyPoints struct {

        
            Applicable float64  `json:"applicable"`
            Description string  `json:"description"`
            IsApplied bool  `json:"is_applied"`
            Total float64  `json:"total"`
         
    }
    
    // CartBreakup used by Cart
    type CartBreakup struct {

        
            Raw RawBreakup  `json:"raw"`
            Display []DisplayBreakup  `json:"display"`
            Coupon CouponBreakup  `json:"coupon"`
            LoyaltyPoints LoyaltyPoints  `json:"loyalty_points"`
         
    }
    
    // ProductAvailabilitySize used by Cart
    type ProductAvailabilitySize struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Value string  `json:"value"`
            Display string  `json:"display"`
            IsAvailable bool  `json:"is_available"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // ProductAvailability used by Cart
    type ProductAvailability struct {

        
            OutOfStock bool  `json:"out_of_stock"`
            Sizes []string  `json:"sizes"`
            OtherStoreQuantity float64  `json:"other_store_quantity"`
            AvailableSizes []ProductAvailabilitySize  `json:"available_sizes"`
            IsValid bool  `json:"is_valid"`
            Deliverable bool  `json:"deliverable"`
         
    }
    
    // CouponDetails used by Cart
    type CouponDetails struct {

        
            DiscountSingleQuantity float64  `json:"discount_single_quantity"`
            Code string  `json:"code"`
            DiscountTotalQuantity float64  `json:"discount_total_quantity"`
         
    }
    
    // DiscountRulesApp used by Cart
    type DiscountRulesApp struct {

        
            MatchedBuyRules []string  `json:"matched_buy_rules"`
            RawOffer map[string]interface{}  `json:"raw_offer"`
            ItemCriteria map[string]interface{}  `json:"item_criteria"`
            Offer map[string]interface{}  `json:"offer"`
         
    }
    
    // FreeGiftItem used by Cart
    type FreeGiftItem struct {

        
            ItemImagesURL []string  `json:"item_images_url"`
            ItemSlug string  `json:"item_slug"`
            ItemPriceDetails map[string]interface{}  `json:"item_price_details"`
            ItemBrandName string  `json:"item_brand_name"`
            ItemName string  `json:"item_name"`
            ItemID float64  `json:"item_id"`
         
    }
    
    // AppliedFreeArticles used by Cart
    type AppliedFreeArticles struct {

        
            FreeGiftItemDetails FreeGiftItem  `json:"free_gift_item_details"`
            Quantity float64  `json:"quantity"`
            ParentItemIdentifier string  `json:"parent_item_identifier"`
            ArticleID string  `json:"article_id"`
         
    }
    
    // Ownership2 used by Cart
    type Ownership2 struct {

        
            PayableCategory string  `json:"payable_category"`
            PayableBy string  `json:"payable_by"`
         
    }
    
    // BuyRules used by Cart
    type BuyRules struct {

        
            CartConditions map[string]interface{}  `json:"cart_conditions"`
            ItemCriteria map[string]interface{}  `json:"item_criteria"`
         
    }
    
    // AppliedPromotion used by Cart
    type AppliedPromotion struct {

        
            MrpPromotion bool  `json:"mrp_promotion"`
            Amount float64  `json:"amount"`
            DiscountRules []DiscountRulesApp  `json:"discount_rules"`
            PromotionType string  `json:"promotion_type"`
            PromotionName string  `json:"promotion_name"`
            PromoID string  `json:"promo_id"`
            PromotionOfferText string  `json:"promotion_offer_text"`
            AppliedFreeArticles []AppliedFreeArticles  `json:"applied_free_articles"`
            Ownership Ownership2  `json:"ownership"`
            PromotionGroup string  `json:"promotion_group"`
            ArticleQuantity float64  `json:"article_quantity"`
            BuyRules []BuyRules  `json:"buy_rules"`
         
    }
    
    // ProductPrice used by Cart
    type ProductPrice struct {

        
            CurrencyCode string  `json:"currency_code"`
            AddOn float64  `json:"add_on"`
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
    
    // BasePrice used by Cart
    type BasePrice struct {

        
            CurrencyCode string  `json:"currency_code"`
            CurrencySymbol string  `json:"currency_symbol"`
            Effective float64  `json:"effective"`
            Marked float64  `json:"marked"`
         
    }
    
    // ArticlePriceInfo used by Cart
    type ArticlePriceInfo struct {

        
            Base BasePrice  `json:"base"`
            Converted BasePrice  `json:"converted"`
         
    }
    
    // BaseInfo used by Cart
    type BaseInfo struct {

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // StoreInfo used by Cart
    type StoreInfo struct {

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            StoreCode string  `json:"store_code"`
         
    }
    
    // ProductArticle used by Cart
    type ProductArticle struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            SellerIdentifier string  `json:"seller_identifier"`
            MtoQuantity float64  `json:"mto_quantity"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            ProductGroupTags []string  `json:"product_group_tags"`
            Size string  `json:"size"`
            IsGiftVisible bool  `json:"is_gift_visible"`
            UID string  `json:"uid"`
            Quantity float64  `json:"quantity"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            CartItemMeta map[string]interface{}  `json:"cart_item_meta"`
            Price ArticlePriceInfo  `json:"price"`
            Seller BaseInfo  `json:"seller"`
            Type string  `json:"type"`
            Identifier map[string]interface{}  `json:"identifier"`
            Store StoreInfo  `json:"store"`
            GiftCard map[string]interface{}  `json:"gift_card"`
         
    }
    
    // PromiseTimestamp used by Cart
    type PromiseTimestamp struct {

        
            Min float64  `json:"min"`
            Max float64  `json:"max"`
         
    }
    
    // PromiseFormatted used by Cart
    type PromiseFormatted struct {

        
            Min string  `json:"min"`
            Max string  `json:"max"`
         
    }
    
    // ShipmentPromise used by Cart
    type ShipmentPromise struct {

        
            Timestamp PromiseTimestamp  `json:"timestamp"`
            Formatted PromiseFormatted  `json:"formatted"`
         
    }
    
    // NetQuantity used by Cart
    type NetQuantity struct {

        
            Unit string  `json:"unit"`
            Value string  `json:"value"`
         
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
    
    // CategoryInfo used by Cart
    type CategoryInfo struct {

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // ProductImage used by Cart
    type ProductImage struct {

        
            SecureURL string  `json:"secure_url"`
            URL string  `json:"url"`
            AspectRatio string  `json:"aspect_ratio"`
         
    }
    
    // Tags used by Cart
    type Tags struct {

        
            Tags map[string]interface{}  `json:"tags"`
         
    }
    
    // CartProduct used by Cart
    type CartProduct struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ItemCode string  `json:"item_code"`
            Tags []string  `json:"tags"`
            NetQuantity NetQuantity  `json:"net_quantity"`
            Action ProductAction  `json:"action"`
            Categories []CategoryInfo  `json:"categories"`
            Images []ProductImage  `json:"images"`
            TeaserTag Tags  `json:"teaser_tag"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Type string  `json:"type"`
            Brand BaseInfo  `json:"brand"`
         
    }
    
    // PromoMeta used by Cart
    type PromoMeta struct {

        
            Message string  `json:"message"`
         
    }
    
    // CartProductIdentifer used by Cart
    type CartProductIdentifer struct {

        
            Identifier string  `json:"identifier"`
         
    }
    
    // CartProductInfo used by Cart
    type CartProductInfo struct {

        
            CouponMessage string  `json:"coupon_message"`
            Availability ProductAvailability  `json:"availability"`
            IsSet bool  `json:"is_set"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            BulkOffer map[string]interface{}  `json:"bulk_offer"`
            Quantity float64  `json:"quantity"`
            Discount string  `json:"discount"`
            Coupon CouponDetails  `json:"coupon"`
            PromotionsApplied []AppliedPromotion  `json:"promotions_applied"`
            Moq map[string]interface{}  `json:"moq"`
            Price ProductPriceInfo  `json:"price"`
            Message string  `json:"message"`
            Article ProductArticle  `json:"article"`
            PricePerUnit ProductPriceInfo  `json:"price_per_unit"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            Product CartProduct  `json:"product"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            PromoMeta PromoMeta  `json:"promo_meta"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
            Key string  `json:"key"`
         
    }
    
    // OpenapiCartDetailsResponse used by Cart
    type OpenapiCartDetailsResponse struct {

        
            Message string  `json:"message"`
            IsValid bool  `json:"is_valid"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            Items []CartProductInfo  `json:"items"`
         
    }
    
    // OpenApiErrorResponse used by Cart
    type OpenApiErrorResponse struct {

        
            Message string  `json:"message"`
            Errors map[string]interface{}  `json:"errors"`
            Success bool  `json:"success"`
         
    }
    
    // ShippingAddress used by Cart
    type ShippingAddress struct {

        
            CountryCode string  `json:"country_code"`
            CountryPhoneCode string  `json:"country_phone_code"`
            Phone float64  `json:"phone"`
            Address string  `json:"address"`
            AddressType string  `json:"address_type"`
            Pincode float64  `json:"pincode"`
            State string  `json:"state"`
            AreaCode string  `json:"area_code"`
            Country string  `json:"country"`
            Landmark string  `json:"landmark"`
            Area string  `json:"area"`
            CountryIsoCode string  `json:"country_iso_code"`
            Name string  `json:"name"`
            City string  `json:"city"`
            AreaCodeSlug string  `json:"area_code_slug"`
            Email string  `json:"email"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // OpenApiCartServiceabilityRequest used by Cart
    type OpenApiCartServiceabilityRequest struct {

        
            ShippingAddress ShippingAddress  `json:"shipping_address"`
            CartItems []CartItem  `json:"cart_items"`
         
    }
    
    // OpenApiCartServiceabilityResponse used by Cart
    type OpenApiCartServiceabilityResponse struct {

        
            BreakupValues CartBreakup  `json:"breakup_values"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            IsValid bool  `json:"is_valid"`
            Items []CartProductInfo  `json:"items"`
            Message string  `json:"message"`
         
    }
    
    // MultiTenderPaymentMeta used by Cart
    type MultiTenderPaymentMeta struct {

        
            OrderID string  `json:"order_id"`
            PaymentGateway string  `json:"payment_gateway"`
            PaymentID string  `json:"payment_id"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            CurrentStatus string  `json:"current_status"`
         
    }
    
    // MultiTenderPaymentMethod used by Cart
    type MultiTenderPaymentMethod struct {

        
            Name string  `json:"name"`
            Meta MultiTenderPaymentMeta  `json:"meta"`
            Amount float64  `json:"amount"`
            Mode string  `json:"mode"`
         
    }
    
    // OpenApiFiles used by Cart
    type OpenApiFiles struct {

        
            Key string  `json:"key"`
            Values []string  `json:"values"`
         
    }
    
    // CartItemMeta used by Cart
    type CartItemMeta struct {

        
            PrimaryItem bool  `json:"primary_item"`
            GroupID string  `json:"group_id"`
         
    }
    
    // OpenApiOrderItem used by Cart
    type OpenApiOrderItem struct {

        
            PaymentMethods []MultiTenderPaymentMethod  `json:"payment_methods"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            AmountPaid float64  `json:"amount_paid"`
            CodCharges float64  `json:"cod_charges"`
            Quantity float64  `json:"quantity"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            Size string  `json:"size"`
            CashbackApplied float64  `json:"cashback_applied"`
            ProductID float64  `json:"product_id"`
            Discount float64  `json:"discount"`
            EmployeeDiscount float64  `json:"employee_discount"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            Files []OpenApiFiles  `json:"files"`
            PriceMarked float64  `json:"price_marked"`
            Meta CartItemMeta  `json:"meta"`
            DeliveryCharges float64  `json:"delivery_charges"`
            PriceEffective float64  `json:"price_effective"`
         
    }
    
    // OpenApiPlatformCheckoutReq used by Cart
    type OpenApiPlatformCheckoutReq struct {

        
            CurrencyCode string  `json:"currency_code"`
            OrderID string  `json:"order_id"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            CartItems []OpenApiOrderItem  `json:"cart_items"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            EmployeeDiscount map[string]interface{}  `json:"employee_discount"`
            Comment string  `json:"comment"`
            CashbackApplied float64  `json:"cashback_applied"`
            PaymentMethods []MultiTenderPaymentMethod  `json:"payment_methods"`
            Files []OpenApiFiles  `json:"files"`
            BillingAddress ShippingAddress  `json:"billing_address"`
            DeliveryCharges float64  `json:"delivery_charges"`
            CouponValue float64  `json:"coupon_value"`
            Coupon string  `json:"coupon"`
            Gstin string  `json:"gstin"`
            CartValue float64  `json:"cart_value"`
            ShippingAddress ShippingAddress  `json:"shipping_address"`
            CodCharges float64  `json:"cod_charges"`
            CouponCode string  `json:"coupon_code"`
            PaymentMode string  `json:"payment_mode"`
         
    }
    
    // OpenApiCheckoutResponse used by Cart
    type OpenApiCheckoutResponse struct {

        
            Message string  `json:"message"`
            OrderID string  `json:"order_id"`
            Success bool  `json:"success"`
            OrderRefID string  `json:"order_ref_id"`
         
    }
    
    // AbandonedCart used by Cart
    type AbandonedCart struct {

        
            Payments map[string]interface{}  `json:"payments"`
            UserID string  `json:"user_id"`
            FcIndexMap []float64  `json:"fc_index_map"`
            IsArchive bool  `json:"is_archive"`
            Comment string  `json:"comment"`
            Discount float64  `json:"discount"`
            PaymentMethods []map[string]interface{}  `json:"payment_methods"`
            Cashback map[string]interface{}  `json:"cashback"`
            Coupon map[string]interface{}  `json:"coupon"`
            Gstin string  `json:"gstin"`
            CheckoutMode string  `json:"checkout_mode"`
            CodCharges map[string]interface{}  `json:"cod_charges"`
            PickUpCustomerDetails map[string]interface{}  `json:"pick_up_customer_details"`
            IsActive bool  `json:"is_active"`
            FyndCredits map[string]interface{}  `json:"fynd_credits"`
            OrderID string  `json:"order_id"`
            LastModified string  `json:"last_modified"`
            BulkCouponDiscount float64  `json:"bulk_coupon_discount"`
            UID float64  `json:"uid"`
            AppID string  `json:"app_id"`
            Articles []map[string]interface{}  `json:"articles"`
            MergeQty bool  `json:"merge_qty"`
            Meta map[string]interface{}  `json:"meta"`
            DeliveryCharges map[string]interface{}  `json:"delivery_charges"`
            Promotion map[string]interface{}  `json:"promotion"`
            BuyNow bool  `json:"buy_now"`
            CartValue float64  `json:"cart_value"`
            IsDefault bool  `json:"is_default"`
            ID string  `json:"_id"`
            CreatedOn string  `json:"created_on"`
            ExpireAt string  `json:"expire_at"`
            Shipments []map[string]interface{}  `json:"shipments"`
            PaymentMode string  `json:"payment_mode"`
         
    }
    
    // AbandonedCartResponse used by Cart
    type AbandonedCartResponse struct {

        
            Result map[string]interface{}  `json:"result"`
            Success bool  `json:"success"`
            Page Page  `json:"page"`
            Items []AbandonedCart  `json:"items"`
            Message string  `json:"message"`
         
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
            Items []CartProductInfo  `json:"items"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            Comment string  `json:"comment"`
            CouponText string  `json:"coupon_text"`
            Gstin string  `json:"gstin"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            CheckoutMode string  `json:"checkout_mode"`
            Message string  `json:"message"`
            AppliedPromoDetails []AppliedPromotion  `json:"applied_promo_details"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            Currency CartCurrency  `json:"currency"`
            LastModified string  `json:"last_modified"`
            PanConfig map[string]interface{}  `json:"pan_config"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            IsValid bool  `json:"is_valid"`
            BuyNow bool  `json:"buy_now"`
            ID string  `json:"id"`
            PanNo string  `json:"pan_no"`
         
    }
    
    // AddProductCart used by Cart
    type AddProductCart struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Pos bool  `json:"pos"`
            ArticleID string  `json:"article_id"`
            ParentItemIdentifiers []map[string]string  `json:"parent_item_identifiers"`
            ProductGroupTags []string  `json:"product_group_tags"`
            SellerID float64  `json:"seller_id"`
            ArticleAssignment map[string]interface{}  `json:"article_assignment"`
            Display string  `json:"display"`
            Quantity float64  `json:"quantity"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            ItemSize string  `json:"item_size"`
            StoreID float64  `json:"store_id"`
            ItemID float64  `json:"item_id"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // AddCartRequest used by Cart
    type AddCartRequest struct {

        
            Items []AddProductCart  `json:"items"`
            NewCart bool  `json:"new_cart"`
         
    }
    
    // AddCartDetailResponse used by Cart
    type AddCartDetailResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
            Cart CartDetailResponse  `json:"cart"`
            Partial bool  `json:"partial"`
         
    }
    
    // UpdateProductCart used by Cart
    type UpdateProductCart struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ArticleID string  `json:"article_id"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            Quantity float64  `json:"quantity"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            ItemSize string  `json:"item_size"`
            ItemID float64  `json:"item_id"`
            ItemIndex float64  `json:"item_index"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // UpdateCartRequest used by Cart
    type UpdateCartRequest struct {

        
            Items []UpdateProductCart  `json:"items"`
            Operation string  `json:"operation"`
         
    }
    
    // UpdateCartDetailResponse used by Cart
    type UpdateCartDetailResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
            Cart CartDetailResponse  `json:"cart"`
         
    }
    
    // OverrideCartItemPromo used by Cart
    type OverrideCartItemPromo struct {

        
            PromoID string  `json:"promo_id"`
            PromoAmount string  `json:"promo_amount"`
            RwrdTndr string  `json:"rwrd_tndr"`
            PromoDesc string  `json:"promo_desc"`
            ItemList []map[string]interface{}  `json:"item_list"`
         
    }
    
    // OverrideCartItem used by Cart
    type OverrideCartItem struct {

        
            PromoList []OverrideCartItemPromo  `json:"promo_list"`
            SellerIdentifier string  `json:"seller_identifier"`
            AmountPaid float64  `json:"amount_paid"`
            Quantity float64  `json:"quantity"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            Size string  `json:"size"`
            ItemID float64  `json:"item_id"`
            Discount float64  `json:"discount"`
            PriceMarked float64  `json:"price_marked"`
            PriceEffective float64  `json:"price_effective"`
         
    }
    
    // OverrideCheckoutReq used by Cart
    type OverrideCheckoutReq struct {

        
            CurrencyCode string  `json:"currency_code"`
            OrderType string  `json:"order_type"`
            ShippingAddress map[string]interface{}  `json:"shipping_address"`
            Aggregator string  `json:"aggregator"`
            CartItems []OverrideCartItem  `json:"cart_items"`
            CartID string  `json:"cart_id"`
            PaymentIdentifier string  `json:"payment_identifier"`
            OrderingStore float64  `json:"ordering_store"`
            MerchantCode string  `json:"merchant_code"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
            PaymentMode string  `json:"payment_mode"`
         
    }
    
    // OverrideCheckoutResponse used by Cart
    type OverrideCheckoutResponse struct {

        
            OrderID string  `json:"order_id"`
            Success string  `json:"success"`
            Cart map[string]interface{}  `json:"cart"`
            Message string  `json:"message"`
            Data map[string]interface{}  `json:"data"`
         
    }
    
    // GetShareCartLinkRequest used by Cart
    type GetShareCartLinkRequest struct {

        
            ID string  `json:"id"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // GetShareCartLinkResponse used by Cart
    type GetShareCartLinkResponse struct {

        
            ShareURL string  `json:"share_url"`
            Token string  `json:"token"`
         
    }
    
    // SharedCartDetails used by Cart
    type SharedCartDetails struct {

        
            Token string  `json:"token"`
            Source map[string]interface{}  `json:"source"`
            User map[string]interface{}  `json:"user"`
            CreatedOn string  `json:"created_on"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // SharedCart used by Cart
    type SharedCart struct {

        
            BreakupValues CartBreakup  `json:"breakup_values"`
            Items []CartProductInfo  `json:"items"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            Comment string  `json:"comment"`
            CouponText string  `json:"coupon_text"`
            Gstin string  `json:"gstin"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            CheckoutMode string  `json:"checkout_mode"`
            Message string  `json:"message"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            Currency CartCurrency  `json:"currency"`
            LastModified string  `json:"last_modified"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            SharedCartDetails SharedCartDetails  `json:"shared_cart_details"`
            UID string  `json:"uid"`
            IsValid bool  `json:"is_valid"`
            CartID float64  `json:"cart_id"`
            BuyNow bool  `json:"buy_now"`
            ID string  `json:"id"`
         
    }
    
    // SharedCartResponse used by Cart
    type SharedCartResponse struct {

        
            Cart SharedCart  `json:"cart"`
            Error string  `json:"error"`
         
    }
    
    // CartList used by Cart
    type CartList struct {

        
            CartValue float64  `json:"cart_value"`
            UserID string  `json:"user_id"`
            CreatedOn string  `json:"created_on"`
            CartID string  `json:"cart_id"`
            PickUpCustomerDetails map[string]interface{}  `json:"pick_up_customer_details"`
            ItemCounts float64  `json:"item_counts"`
         
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

        
            Gender string  `json:"gender"`
            CreatedAt string  `json:"created_at"`
            ExternalID string  `json:"external_id"`
            FirstName string  `json:"first_name"`
            ID string  `json:"_id"`
            Mobile string  `json:"mobile"`
            UID string  `json:"uid"`
            ModifiedOn string  `json:"modified_on"`
            LastName string  `json:"last_name"`
         
    }
    
    // UserCartMappingResponse used by Cart
    type UserCartMappingResponse struct {

        
            BreakupValues CartBreakup  `json:"breakup_values"`
            Items []CartProductInfo  `json:"items"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            Comment string  `json:"comment"`
            CouponText string  `json:"coupon_text"`
            Gstin string  `json:"gstin"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            CheckoutMode string  `json:"checkout_mode"`
            Message string  `json:"message"`
            AppliedPromoDetails []AppliedPromotion  `json:"applied_promo_details"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            Currency CartCurrency  `json:"currency"`
            LastModified string  `json:"last_modified"`
            PanConfig map[string]interface{}  `json:"pan_config"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            User UserInfo  `json:"user"`
            IsValid bool  `json:"is_valid"`
            BuyNow bool  `json:"buy_now"`
            ID string  `json:"id"`
            PanNo string  `json:"pan_no"`
         
    }
    
    // PlatformAddCartRequest used by Cart
    type PlatformAddCartRequest struct {

        
            Items []AddProductCart  `json:"items"`
            NewCart bool  `json:"new_cart"`
            UserID string  `json:"user_id"`
         
    }
    
    // PlatformUpdateCartRequest used by Cart
    type PlatformUpdateCartRequest struct {

        
            Items []UpdateProductCart  `json:"items"`
            Operation string  `json:"operation"`
            UserID string  `json:"user_id"`
         
    }
    
    // DeleteCartRequest used by Cart
    type DeleteCartRequest struct {

        
            CartIDList []string  `json:"cart_id_list"`
         
    }
    
    // DeleteCartDetailResponse used by Cart
    type DeleteCartDetailResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // CartItemCountResponse used by Cart
    type CartItemCountResponse struct {

        
            UserCartItemsCount float64  `json:"user_cart_items_count"`
         
    }
    
    // Coupon used by Cart
    type Coupon struct {

        
            CouponValue float64  `json:"coupon_value"`
            MaxDiscountValue float64  `json:"max_discount_value"`
            IsApplicable bool  `json:"is_applicable"`
            ExpiresOn string  `json:"expires_on"`
            Title string  `json:"title"`
            CouponType string  `json:"coupon_type"`
            MinimumCartValue float64  `json:"minimum_cart_value"`
            Description string  `json:"description"`
            SubTitle string  `json:"sub_title"`
            Message string  `json:"message"`
            IsApplied bool  `json:"is_applied"`
            CouponCode string  `json:"coupon_code"`
         
    }
    
    // PageCoupon used by Cart
    type PageCoupon struct {

        
            Current float64  `json:"current"`
            HasNext bool  `json:"has_next"`
            HasPrevious bool  `json:"has_previous"`
            TotalItemCount float64  `json:"total_item_count"`
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
    
    // PlatformAddress used by Cart
    type PlatformAddress struct {

        
            IsDefaultAddress bool  `json:"is_default_address"`
            UserID string  `json:"user_id"`
            City string  `json:"city"`
            Email string  `json:"email"`
            AddressType string  `json:"address_type"`
            State string  `json:"state"`
            Country string  `json:"country"`
            CheckoutMode string  `json:"checkout_mode"`
            Landmark string  `json:"landmark"`
            IsActive bool  `json:"is_active"`
            Tags []string  `json:"tags"`
            GoogleMapPoint map[string]interface{}  `json:"google_map_point"`
            Area string  `json:"area"`
            Name string  `json:"name"`
            AreaCodeSlug string  `json:"area_code_slug"`
            Meta map[string]interface{}  `json:"meta"`
            CartID string  `json:"cart_id"`
            CountryCode string  `json:"country_code"`
            Phone string  `json:"phone"`
            ID string  `json:"id"`
            Address string  `json:"address"`
            AreaCode string  `json:"area_code"`
            CreatedByUserID string  `json:"created_by_user_id"`
            GeoLocation GeoLocation  `json:"geo_location"`
         
    }
    
    // PlatformGetAddressesResponse used by Cart
    type PlatformGetAddressesResponse struct {

        
            Address []PlatformAddress  `json:"address"`
         
    }
    
    // SaveAddressResponse used by Cart
    type SaveAddressResponse struct {

        
            Success bool  `json:"success"`
            IsDefaultAddress bool  `json:"is_default_address"`
            ID string  `json:"id"`
         
    }
    
    // UpdateAddressResponse used by Cart
    type UpdateAddressResponse struct {

        
            IsDefaultAddress bool  `json:"is_default_address"`
            Success bool  `json:"success"`
            ID string  `json:"id"`
            IsUpdated bool  `json:"is_updated"`
         
    }
    
    // DeleteAddressResponse used by Cart
    type DeleteAddressResponse struct {

        
            IsDeleted bool  `json:"is_deleted"`
            ID string  `json:"id"`
         
    }
    
    // PlatformSelectCartAddressRequest used by Cart
    type PlatformSelectCartAddressRequest struct {

        
            ID string  `json:"id"`
            UserID string  `json:"user_id"`
            CheckoutMode string  `json:"checkout_mode"`
            BillingAddressID string  `json:"billing_address_id"`
            CartID string  `json:"cart_id"`
         
    }
    
    // ShipmentArticle used by Cart
    type ShipmentArticle struct {

        
            Quantity string  `json:"quantity"`
            Meta string  `json:"meta"`
            ArticleID string  `json:"article_id"`
         
    }
    
    // PlatformShipmentResponse used by Cart
    type PlatformShipmentResponse struct {

        
            ShipmentType string  `json:"shipment_type"`
            Promise ShipmentPromise  `json:"promise"`
            FulfillmentType string  `json:"fulfillment_type"`
            OrderType string  `json:"order_type"`
            BoxType string  `json:"box_type"`
            DpOptions map[string]interface{}  `json:"dp_options"`
            Items []CartProductInfo  `json:"items"`
            Articles []ShipmentArticle  `json:"articles"`
            DpID string  `json:"dp_id"`
            Shipments float64  `json:"shipments"`
            FulfillmentID float64  `json:"fulfillment_id"`
         
    }
    
    // PlatformCartShipmentsResponse used by Cart
    type PlatformCartShipmentsResponse struct {

        
            StaffUserID string  `json:"staff_user_id"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            Error bool  `json:"error"`
            Items []CartProductInfo  `json:"items"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            Comment string  `json:"comment"`
            CouponText string  `json:"coupon_text"`
            Gstin string  `json:"gstin"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            CheckoutMode string  `json:"checkout_mode"`
            Message string  `json:"message"`
            AppliedPromoDetails []AppliedPromotion  `json:"applied_promo_details"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            Currency CartCurrency  `json:"currency"`
            LastModified string  `json:"last_modified"`
            PanConfig map[string]interface{}  `json:"pan_config"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            IsValid bool  `json:"is_valid"`
            BuyNow bool  `json:"buy_now"`
            ID string  `json:"id"`
            Shipments []PlatformShipmentResponse  `json:"shipments"`
            PanNo string  `json:"pan_no"`
         
    }
    
    // UpdateCartShipmentItem used by Cart
    type UpdateCartShipmentItem struct {

        
            ShipmentType string  `json:"shipment_type"`
            Quantity float64  `json:"quantity"`
            ArticleUID string  `json:"article_uid"`
         
    }
    
    // UpdateCartShipmentRequest used by Cart
    type UpdateCartShipmentRequest struct {

        
            Shipments []UpdateCartShipmentItem  `json:"shipments"`
         
    }
    
    // PlatformCartMetaRequest used by Cart
    type PlatformCartMetaRequest struct {

        
            Gstin string  `json:"gstin"`
            CheckoutMode string  `json:"checkout_mode"`
            StaffUserID string  `json:"staff_user_id"`
            GiftDetails map[string]interface{}  `json:"gift_details"`
            Comment string  `json:"comment"`
            PanNo string  `json:"pan_no"`
            PickUpCustomerDetails map[string]interface{}  `json:"pick_up_customer_details"`
         
    }
    
    // CartMetaResponse used by Cart
    type CartMetaResponse struct {

        
            Message string  `json:"message"`
            IsValid bool  `json:"is_valid"`
         
    }
    
    // CartMetaMissingResponse used by Cart
    type CartMetaMissingResponse struct {

        
            Errors []string  `json:"errors"`
         
    }
    
    // StaffCheckout used by Cart
    type StaffCheckout struct {

        
            FirstName string  `json:"first_name"`
            ID string  `json:"_id"`
            User string  `json:"user"`
            LastName string  `json:"last_name"`
            EmployeeCode string  `json:"employee_code"`
         
    }
    
    // Files used by Cart
    type Files struct {

        
            Key string  `json:"key"`
            Values []string  `json:"values"`
         
    }
    
    // PlatformCartCheckoutDetailRequest used by Cart
    type PlatformCartCheckoutDetailRequest struct {

        
            Pos bool  `json:"pos"`
            UserID string  `json:"user_id"`
            Staff StaffCheckout  `json:"staff"`
            BillingAddressID string  `json:"billing_address_id"`
            Files []Files  `json:"files"`
            PaymentParams map[string]interface{}  `json:"payment_params"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
            DeviceID string  `json:"device_id"`
            OrderType string  `json:"order_type"`
            CheckoutMode string  `json:"checkout_mode"`
            AddressID string  `json:"address_id"`
            CallbackURL string  `json:"callback_url"`
            Aggregator string  `json:"aggregator"`
            PaymentIdentifier string  `json:"payment_identifier"`
            OrderingStore float64  `json:"ordering_store"`
            MerchantCode string  `json:"merchant_code"`
            Meta map[string]interface{}  `json:"meta"`
            ID string  `json:"id"`
            PaymentAutoConfirm bool  `json:"payment_auto_confirm"`
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
            PickAtStoreUID float64  `json:"pick_at_store_uid"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            EmployeeCode string  `json:"employee_code"`
            PaymentMode string  `json:"payment_mode"`
         
    }
    
    // CheckCart used by Cart
    type CheckCart struct {

        
            CodAvailable bool  `json:"cod_available"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            UserType string  `json:"user_type"`
            Items []CartProductInfo  `json:"items"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            Comment string  `json:"comment"`
            DeliveryChargeOrderValue float64  `json:"delivery_charge_order_value"`
            CouponText string  `json:"coupon_text"`
            Gstin string  `json:"gstin"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            CheckoutMode string  `json:"checkout_mode"`
            StoreEmps []map[string]interface{}  `json:"store_emps"`
            CodCharges float64  `json:"cod_charges"`
            Message string  `json:"message"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            ErrorMessage string  `json:"error_message"`
            Currency CartCurrency  `json:"currency"`
            OrderID string  `json:"order_id"`
            LastModified string  `json:"last_modified"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            UID string  `json:"uid"`
            IsValid bool  `json:"is_valid"`
            StoreCode string  `json:"store_code"`
            CartID float64  `json:"cart_id"`
            CodMessage string  `json:"cod_message"`
            DeliveryCharges float64  `json:"delivery_charges"`
            BuyNow bool  `json:"buy_now"`
            ID string  `json:"id"`
            Success bool  `json:"success"`
         
    }
    
    // CartCheckoutResponse used by Cart
    type CartCheckoutResponse struct {

        
            OrderID string  `json:"order_id"`
            Success bool  `json:"success"`
            CallbackURL string  `json:"callback_url"`
            PaymentConfirmURL string  `json:"payment_confirm_url"`
            Cart CheckCart  `json:"cart"`
            AppInterceptURL string  `json:"app_intercept_url"`
            Message string  `json:"message"`
            Data map[string]interface{}  `json:"data"`
         
    }
    
    // CartDeliveryModesResponse used by Cart
    type CartDeliveryModesResponse struct {

        
            PickupStores []float64  `json:"pickup_stores"`
            AvailableModes []string  `json:"available_modes"`
         
    }
    
    // PickupStoreDetail used by Cart
    type PickupStoreDetail struct {

        
            Phone string  `json:"phone"`
            AddressType string  `json:"address_type"`
            Address string  `json:"address"`
            Pincode float64  `json:"pincode"`
            Country string  `json:"country"`
            AreaCode string  `json:"area_code"`
            State string  `json:"state"`
            ID float64  `json:"id"`
            Landmark string  `json:"landmark"`
            Area string  `json:"area"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            City string  `json:"city"`
            StoreCode string  `json:"store_code"`
            AreaCodeSlug string  `json:"area_code_slug"`
            Email string  `json:"email"`
         
    }
    
    // StoreDetailsResponse used by Cart
    type StoreDetailsResponse struct {

        
            Items []PickupStoreDetail  `json:"items"`
         
    }
    
    // UpdateCartPaymentRequest used by Cart
    type UpdateCartPaymentRequest struct {

        
            ID string  `json:"id"`
            PaymentIdentifier string  `json:"payment_identifier"`
            AggregatorName string  `json:"aggregator_name"`
            MerchantCode string  `json:"merchant_code"`
            PaymentMode string  `json:"payment_mode"`
            AddressID string  `json:"address_id"`
         
    }
    
    // CouponValidity used by Cart
    type CouponValidity struct {

        
            Valid bool  `json:"valid"`
            Code string  `json:"code"`
            DisplayMessageEn string  `json:"display_message_en"`
            NextValidationRequired bool  `json:"next_validation_required"`
            Title string  `json:"title"`
            Discount float64  `json:"discount"`
         
    }
    
    // PaymentCouponValidate used by Cart
    type PaymentCouponValidate struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
            CouponValidity CouponValidity  `json:"coupon_validity"`
         
    }
    
    // PaymentMeta used by Cart
    type PaymentMeta struct {

        
            Type string  `json:"type"`
            PaymentIdentifier string  `json:"payment_identifier"`
            PaymentGateway string  `json:"payment_gateway"`
            MerchantCode string  `json:"merchant_code"`
         
    }
    
    // PaymentMethod used by Cart
    type PaymentMethod struct {

        
            Amount float64  `json:"amount"`
            Payment string  `json:"payment"`
            Mode string  `json:"mode"`
            Name string  `json:"name"`
            PaymentMeta PaymentMeta  `json:"payment_meta"`
         
    }
    
    // PlatformCartCheckoutDetailV2Request used by Cart
    type PlatformCartCheckoutDetailV2Request struct {

        
            Pos bool  `json:"pos"`
            UserID string  `json:"user_id"`
            Staff StaffCheckout  `json:"staff"`
            BillingAddressID string  `json:"billing_address_id"`
            PaymentMethods []PaymentMethod  `json:"payment_methods"`
            CustomMeta map[string]interface{}  `json:"custom_meta"`
            Files []Files  `json:"files"`
            PaymentParams map[string]interface{}  `json:"payment_params"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
            DeviceID string  `json:"device_id"`
            OrderType string  `json:"order_type"`
            CheckoutMode string  `json:"checkout_mode"`
            AddressID string  `json:"address_id"`
            CallbackURL string  `json:"callback_url"`
            Aggregator string  `json:"aggregator"`
            PaymentIdentifier string  `json:"payment_identifier"`
            OrderingStore float64  `json:"ordering_store"`
            MerchantCode string  `json:"merchant_code"`
            Meta map[string]interface{}  `json:"meta"`
            ID string  `json:"id"`
            PaymentAutoConfirm bool  `json:"payment_auto_confirm"`
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
            PickAtStoreUID float64  `json:"pick_at_store_uid"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            EmployeeCode string  `json:"employee_code"`
            PaymentMode string  `json:"payment_mode"`
         
    }
    
    // UpdateCartPaymentRequestV2 used by Cart
    type UpdateCartPaymentRequestV2 struct {

        
            ID string  `json:"id"`
            PaymentIdentifier string  `json:"payment_identifier"`
            AggregatorName string  `json:"aggregator_name"`
            MerchantCode string  `json:"merchant_code"`
            PaymentMethods []PaymentMethod  `json:"payment_methods"`
            PaymentMode string  `json:"payment_mode"`
            AddressID string  `json:"address_id"`
         
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
    
    // DiscountMeta used by Discount
    type DiscountMeta struct {

        
            Timer bool  `json:"timer"`
            NumberOfMinutes float64  `json:"number_of_minutes"`
         
    }
    
    // DiscountItems used by Discount
    type DiscountItems struct {

        
            ItemCode string  `json:"item_code"`
            BrandUID float64  `json:"brand_uid"`
            SellerIdentifier string  `json:"seller_identifier"`
            DiscountType string  `json:"discount_type"`
            Value float64  `json:"value"`
            DiscountMeta DiscountMeta  `json:"discount_meta"`
         
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
    
    // FileJobRequest used by Discount
    type FileJobRequest struct {

        
            Name string  `json:"name"`
            IsActive bool  `json:"is_active"`
            AppIds []string  `json:"app_ids"`
            JobType string  `json:"job_type"`
            DiscountType string  `json:"discount_type"`
            DiscountLevel string  `json:"discount_level"`
            FilePath string  `json:"file_path"`
            BrandIds []float64  `json:"brand_ids"`
            StoreIds []float64  `json:"store_ids"`
            Validity ValidityObject  `json:"validity"`
            Meta map[string]interface{}  `json:"meta"`
         
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
    

    
    // ServiceabilityErrorResponse used by Serviceability
    type ServiceabilityErrorResponse struct {

        
            Value string  `json:"value"`
            Type string  `json:"type"`
            Message string  `json:"message"`
         
    }
    
    // ApplicationServiceabilityConfig used by Serviceability
    type ApplicationServiceabilityConfig struct {

        
            ChannelType string  `json:"channel_type"`
            ChannelID string  `json:"channel_id"`
            ServiceabilityType string  `json:"serviceability_type"`
         
    }
    
    // ApplicationServiceabilityConfigResponse used by Serviceability
    type ApplicationServiceabilityConfigResponse struct {

        
            Error ServiceabilityErrorResponse  `json:"error"`
            Data ApplicationServiceabilityConfig  `json:"data"`
            Success bool  `json:"success"`
         
    }
    
    // EntityRegionView_Request used by Serviceability
    type EntityRegionView_Request struct {

        
            SubType []string  `json:"sub_type"`
            ParentID []string  `json:"parent_id"`
         
    }
    
    // EntityRegionView_Items used by Serviceability
    type EntityRegionView_Items struct {

        
            Name string  `json:"name"`
            SubType string  `json:"sub_type"`
            UID string  `json:"uid"`
         
    }
    
    // EntityRegionView_Error used by Serviceability
    type EntityRegionView_Error struct {

        
            Value string  `json:"value"`
            Type string  `json:"type"`
            Message string  `json:"message"`
         
    }
    
    // EntityRegionView_page used by Serviceability
    type EntityRegionView_page struct {

        
            Size float64  `json:"size"`
            Current float64  `json:"current"`
            ItemTotal float64  `json:"item_total"`
            HasNext bool  `json:"has_next"`
            Type string  `json:"type"`
         
    }
    
    // EntityRegionView_Response used by Serviceability
    type EntityRegionView_Response struct {

        
            Data []EntityRegionView_Items  `json:"data"`
            Success bool  `json:"success"`
            Error EntityRegionView_Error  `json:"error"`
            Page EntityRegionView_page  `json:"page"`
         
    }
    
    // ListViewProduct used by Serviceability
    type ListViewProduct struct {

        
            Count float64  `json:"count"`
            Type string  `json:"type"`
         
    }
    
    // ListViewChannels used by Serviceability
    type ListViewChannels struct {

        
            ChannelType string  `json:"channel_type"`
            ChannelID string  `json:"channel_id"`
         
    }
    
    // ListViewItems used by Serviceability
    type ListViewItems struct {

        
            StoresCount float64  `json:"stores_count"`
            Slug string  `json:"slug"`
            PincodesCount float64  `json:"pincodes_count"`
            IsActive bool  `json:"is_active"`
            ZoneID string  `json:"zone_id"`
            CompanyID float64  `json:"company_id"`
            Product ListViewProduct  `json:"product"`
            Name string  `json:"name"`
            Channels ListViewChannels  `json:"channels"`
         
    }
    
    // ZoneDataItem used by Serviceability
    type ZoneDataItem struct {

        
            Size float64  `json:"size"`
            Type string  `json:"type"`
            Current float64  `json:"current"`
            ItemTotal float64  `json:"item_total"`
            HasNext bool  `json:"has_next"`
         
    }
    
    // ListViewSummary used by Serviceability
    type ListViewSummary struct {

        
            TotalPincodesServed float64  `json:"total_pincodes_served"`
            TotalZones float64  `json:"total_zones"`
            TotalActiveZones float64  `json:"total_active_zones"`
         
    }
    
    // ListViewResponse used by Serviceability
    type ListViewResponse struct {

        
            Items []ListViewItems  `json:"items"`
            Page []ZoneDataItem  `json:"page"`
            Summary []ListViewSummary  `json:"summary"`
         
    }
    
    // CompanyStoreView_PageItems used by Serviceability
    type CompanyStoreView_PageItems struct {

        
            Size float64  `json:"size"`
            Current float64  `json:"current"`
            ItemTotal float64  `json:"item_total"`
            HasNext bool  `json:"has_next"`
            Type string  `json:"type"`
         
    }
    
    // CompanyStoreView_Response used by Serviceability
    type CompanyStoreView_Response struct {

        
            Items []map[string]interface{}  `json:"items"`
            Page []CompanyStoreView_PageItems  `json:"page"`
         
    }
    
    // GetZoneDataViewChannels used by Serviceability
    type GetZoneDataViewChannels struct {

        
            ChannelType string  `json:"channel_type"`
            ChannelID string  `json:"channel_id"`
         
    }
    
    // ZoneProductTypes used by Serviceability
    type ZoneProductTypes struct {

        
            Tags []string  `json:"tags"`
            Type string  `json:"type"`
         
    }
    
    // ZoneMappingType used by Serviceability
    type ZoneMappingType struct {

        
            Pincode []string  `json:"pincode"`
            State []string  `json:"state"`
            Country string  `json:"country"`
         
    }
    
    // UpdateZoneData used by Serviceability
    type UpdateZoneData struct {

        
            ZoneID string  `json:"zone_id"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            CompanyID float64  `json:"company_id"`
            IsActive bool  `json:"is_active"`
            Channels []GetZoneDataViewChannels  `json:"channels"`
            Product ZoneProductTypes  `json:"product"`
            StoreIds []float64  `json:"store_ids"`
            RegionType string  `json:"region_type"`
            Mapping []ZoneMappingType  `json:"mapping"`
            AssignmentPreference string  `json:"assignment_preference"`
         
    }
    
    // ZoneUpdateRequest used by Serviceability
    type ZoneUpdateRequest struct {

        
            Data UpdateZoneData  `json:"data"`
            Identifier string  `json:"identifier"`
         
    }
    
    // ZoneSuccessResponse used by Serviceability
    type ZoneSuccessResponse struct {

        
            Success bool  `json:"success"`
            StatusCode float64  `json:"status_code"`
         
    }
    
    // GetZoneDataViewItems used by Serviceability
    type GetZoneDataViewItems struct {

        
            ZoneID string  `json:"zone_id"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            CompanyID float64  `json:"company_id"`
            IsActive bool  `json:"is_active"`
            Channels []GetZoneDataViewChannels  `json:"channels"`
            Product ZoneProductTypes  `json:"product"`
            StoreIds []float64  `json:"store_ids"`
            RegionType string  `json:"region_type"`
            Mapping []ZoneMappingType  `json:"mapping"`
            AssignmentPreference string  `json:"assignment_preference"`
            StoresCount float64  `json:"stores_count"`
            PincodesCount float64  `json:"pincodes_count"`
         
    }
    
    // GetSingleZoneDataViewResponse used by Serviceability
    type GetSingleZoneDataViewResponse struct {

        
            Data GetZoneDataViewItems  `json:"data"`
         
    }
    
    // CreateZoneData used by Serviceability
    type CreateZoneData struct {

        
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            CompanyID float64  `json:"company_id"`
            IsActive bool  `json:"is_active"`
            Channels []GetZoneDataViewChannels  `json:"channels"`
            Product ZoneProductTypes  `json:"product"`
            StoreIds []float64  `json:"store_ids"`
            RegionType string  `json:"region_type"`
            Mapping []ZoneMappingType  `json:"mapping"`
            AssignmentPreference string  `json:"assignment_preference"`
         
    }
    
    // ZoneRequest used by Serviceability
    type ZoneRequest struct {

        
            Data CreateZoneData  `json:"data"`
            Identifier string  `json:"identifier"`
         
    }
    
    // ZoneResponse used by Serviceability
    type ZoneResponse struct {

        
            Success bool  `json:"success"`
            StatusCode float64  `json:"status_code"`
            ZoneID string  `json:"zone_id"`
         
    }
    
    // GetZoneFromPincodeViewRequest used by Serviceability
    type GetZoneFromPincodeViewRequest struct {

        
            Country string  `json:"country"`
            Pincode string  `json:"pincode"`
         
    }
    
    // GetZoneFromPincodeViewResponse used by Serviceability
    type GetZoneFromPincodeViewResponse struct {

        
            ServiceabilityType string  `json:"serviceability_type"`
            Zones []string  `json:"zones"`
         
    }
    
    // GetZoneFromApplicationIdViewResponse used by Serviceability
    type GetZoneFromApplicationIdViewResponse struct {

        
            Items []ListViewItems  `json:"items"`
            Page []ZoneDataItem  `json:"page"`
         
    }
    
    // WarningsResponse used by Serviceability
    type WarningsResponse struct {

        
            StoreAddress string  `json:"store_address"`
         
    }
    
    // IntegrationTypeResponse used by Serviceability
    type IntegrationTypeResponse struct {

        
            Inventory string  `json:"inventory"`
            Order string  `json:"order"`
         
    }
    
    // CreatedByResponse used by Serviceability
    type CreatedByResponse struct {

        
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
         
    }
    
    // ProductReturnConfigResponse used by Serviceability
    type ProductReturnConfigResponse struct {

        
            OnSameStore bool  `json:"on_same_store"`
         
    }
    
    // Dp used by Serviceability
    type Dp struct {

        
            AreaCode float64  `json:"area_code"`
            AssignDpFromSb bool  `json:"assign_dp_from_sb"`
            Operations []string  `json:"operations"`
            TransportMode string  `json:"transport_mode"`
            FmPriority float64  `json:"fm_priority"`
            PaymentMode string  `json:"payment_mode"`
            LmPriority float64  `json:"lm_priority"`
            InternalAccountID string  `json:"internal_account_id"`
            RvpPriority float64  `json:"rvp_priority"`
            ExternalAccountID string  `json:"external_account_id"`
         
    }
    
    // LogisticsResponse used by Serviceability
    type LogisticsResponse struct {

        
            Override bool  `json:"override"`
            Dp Dp  `json:"dp"`
         
    }
    
    // ContactNumberResponse used by Serviceability
    type ContactNumberResponse struct {

        
            Number string  `json:"number"`
            CountryCode float64  `json:"country_code"`
         
    }
    
    // EinvoiceResponse used by Serviceability
    type EinvoiceResponse struct {

        
            Enabled bool  `json:"enabled"`
         
    }
    
    // EwayBillResponse used by Serviceability
    type EwayBillResponse struct {

        
            Enabled bool  `json:"enabled"`
         
    }
    
    // GstCredentialsResponse used by Serviceability
    type GstCredentialsResponse struct {

        
            EInvoice EinvoiceResponse  `json:"e_invoice"`
            EWaybill EwayBillResponse  `json:"e_waybill"`
         
    }
    
    // ModifiedByResponse used by Serviceability
    type ModifiedByResponse struct {

        
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
         
    }
    
    // MobileNo used by Serviceability
    type MobileNo struct {

        
            Number string  `json:"number"`
            CountryCode float64  `json:"country_code"`
         
    }
    
    // ManagerResponse used by Serviceability
    type ManagerResponse struct {

        
            Name string  `json:"name"`
            MobileNo MobileNo  `json:"mobile_no"`
            Email string  `json:"email"`
         
    }
    
    // DocumentsResponse used by Serviceability
    type DocumentsResponse struct {

        
            Value string  `json:"value"`
            Verified bool  `json:"verified"`
            Type string  `json:"type"`
            LegalName string  `json:"legal_name"`
         
    }
    
    // OpeningClosing used by Serviceability
    type OpeningClosing struct {

        
            Hour float64  `json:"hour"`
            Minute float64  `json:"minute"`
         
    }
    
    // TimmingResponse used by Serviceability
    type TimmingResponse struct {

        
            Open bool  `json:"open"`
            Opening OpeningClosing  `json:"opening"`
            Closing OpeningClosing  `json:"closing"`
            Weekday string  `json:"weekday"`
         
    }
    
    // AddressResponse used by Serviceability
    type AddressResponse struct {

        
            City string  `json:"city"`
            Country string  `json:"country"`
            Pincode float64  `json:"pincode"`
            Landmark string  `json:"landmark"`
            Address1 string  `json:"address1"`
            Longitude float64  `json:"longitude"`
            State string  `json:"state"`
            Address2 string  `json:"address2"`
            Latitude float64  `json:"latitude"`
         
    }
    
    // ItemResponse used by Serviceability
    type ItemResponse struct {

        
            Code string  `json:"code"`
            Warnings WarningsResponse  `json:"warnings"`
            Cls string  `json:"_cls"`
            IntegrationType IntegrationTypeResponse  `json:"integration_type"`
            VerifiedOn string  `json:"verified_on"`
            Company float64  `json:"company"`
            CreatedBy CreatedByResponse  `json:"created_by"`
            Stage string  `json:"stage"`
            NotificationEmails []string  `json:"notification_emails"`
            StoreType string  `json:"store_type"`
            ProductReturnConfig ProductReturnConfigResponse  `json:"product_return_config"`
            SubType string  `json:"sub_type"`
            Name string  `json:"name"`
            CreatedOn string  `json:"created_on"`
            Logistics LogisticsResponse  `json:"logistics"`
            ContactNumbers []ContactNumberResponse  `json:"contact_numbers"`
            GstCredentials GstCredentialsResponse  `json:"gst_credentials"`
            VerifiedBy ModifiedByResponse  `json:"verified_by"`
            CompanyID float64  `json:"company_id"`
            DisplayName string  `json:"display_name"`
            Manager ManagerResponse  `json:"manager"`
            ModifiedBy ModifiedByResponse  `json:"modified_by"`
            ModifiedOn string  `json:"modified_on"`
            Documents []DocumentsResponse  `json:"documents"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Timing []TimmingResponse  `json:"timing"`
            UID float64  `json:"uid"`
            Address AddressResponse  `json:"address"`
         
    }
    
    // ServiceabilityPageResponse used by Serviceability
    type ServiceabilityPageResponse struct {

        
            Size float64  `json:"size"`
            Current float64  `json:"current"`
            ItemTotal float64  `json:"item_total"`
            HasNext bool  `json:"has_next"`
            Type string  `json:"type"`
         
    }
    
    // GetStoresViewResponse used by Serviceability
    type GetStoresViewResponse struct {

        
            Items []ItemResponse  `json:"items"`
            Page ServiceabilityPageResponse  `json:"page"`
         
    }
    
    // ReAssignStoreRequest used by Serviceability
    type ReAssignStoreRequest struct {

        
            Configuration map[string]interface{}  `json:"configuration"`
            Articles []map[string]interface{}  `json:"articles"`
            ToPincode string  `json:"to_pincode"`
            IgnoredLocations []string  `json:"ignored_locations"`
            Identifier string  `json:"identifier"`
         
    }
    
    // ReAssignStoreResponse used by Serviceability
    type ReAssignStoreResponse struct {

        
            Error map[string]interface{}  `json:"error"`
            Articles []map[string]interface{}  `json:"articles"`
            ToPincode string  `json:"to_pincode"`
            Success bool  `json:"success"`
         
    }
    
    // ApplicationCompanyDpViewRequest used by Serviceability
    type ApplicationCompanyDpViewRequest struct {

        
            DpID string  `json:"dp_id"`
         
    }
    
    // ApplicationCompanyDpViewResponse used by Serviceability
    type ApplicationCompanyDpViewResponse struct {

        
            CourierPartnerID float64  `json:"courier_partner_id"`
            Success bool  `json:"success"`
            ApplicationID string  `json:"application_id"`
            CompanyID float64  `json:"company_id"`
         
    }
    
    // PincodeMopData used by Serviceability
    type PincodeMopData struct {

        
            Pincodes []float64  `json:"pincodes"`
            Country string  `json:"country"`
            Action string  `json:"action"`
         
    }
    
    // PincodeMopUpdateResponse used by Serviceability
    type PincodeMopUpdateResponse struct {

        
            Pincode float64  `json:"pincode"`
            ChannelID string  `json:"channel_id"`
            Country string  `json:"country"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // PincodeMOPresponse used by Serviceability
    type PincodeMOPresponse struct {

        
            Success bool  `json:"success"`
            StatusCode float64  `json:"status_code"`
            BatchID string  `json:"batch_id"`
            Country string  `json:"country"`
            Action string  `json:"action"`
            Pincodes []float64  `json:"pincodes"`
            UpdatedPincodes []PincodeMopUpdateResponse  `json:"updated_pincodes"`
         
    }
    
    // CommonError used by Serviceability
    type CommonError struct {

        
            Error interface{}  `json:"error"`
            Success string  `json:"success"`
            StatusCode string  `json:"status_code"`
         
    }
    
    // PincodeMopBulkData used by Serviceability
    type PincodeMopBulkData struct {

        
            BatchID string  `json:"batch_id"`
            S3URL string  `json:"s3_url"`
         
    }
    
    // PincodeBulkViewResponse used by Serviceability
    type PincodeBulkViewResponse struct {

        
            BatchID string  `json:"batch_id"`
            S3URL string  `json:"s3_url"`
         
    }
    
    // PincodeCodStatusListingRequest used by Serviceability
    type PincodeCodStatusListingRequest struct {

        
            Country string  `json:"country"`
            IsActive bool  `json:"is_active"`
            Pincode float64  `json:"pincode"`
            Current float64  `json:"current"`
            PageSize float64  `json:"page_size"`
         
    }
    
    // PincodeCodStatusListingResponse used by Serviceability
    type PincodeCodStatusListingResponse struct {

        
            Country string  `json:"country"`
            Data []PincodeCodStatusListingResponse  `json:"data"`
            Success bool  `json:"success"`
            Errors []Error  `json:"errors"`
            Page PincodeCodStatusListingPage  `json:"page"`
            Summary PincodeCodStatusListingSummary  `json:"summary"`
         
    }
    
    // Error used by Serviceability
    type Error struct {

        
            Type string  `json:"type"`
            Value string  `json:"value"`
            Message string  `json:"message"`
         
    }
    
    // PincodeCodStatusListingPage used by Serviceability
    type PincodeCodStatusListingPage struct {

        
            Size float64  `json:"size"`
            Current float64  `json:"current"`
            ItemTotal float64  `json:"item_total"`
            HasNext bool  `json:"has_next"`
            Type string  `json:"type"`
         
    }
    
    // PincodeCodStatusListingSummary used by Serviceability
    type PincodeCodStatusListingSummary struct {

        
            TotalActivePincodes float64  `json:"total_active_pincodes"`
            TotalInactivePincodes float64  `json:"total_inactive_pincodes"`
         
    }
    
    // PincodeMopUpdateAuditHistoryRequest used by Serviceability
    type PincodeMopUpdateAuditHistoryRequest struct {

        
            EntityType string  `json:"entity_type"`
            FileName string  `json:"file_name"`
         
    }
    
    // PincodeMopUpdateAuditHistoryPaging used by Serviceability
    type PincodeMopUpdateAuditHistoryPaging struct {

        
            Type string  `json:"type"`
            Size float64  `json:"size"`
            Current float64  `json:"current"`
            HasNext bool  `json:"has_next"`
            ItemTotal float64  `json:"item_total"`
         
    }
    
    // PincodeMopUpdateAuditHistoryResponse used by Serviceability
    type PincodeMopUpdateAuditHistoryResponse struct {

        
            BatchID string  `json:"batch_id"`
            EntityType string  `json:"entity_type"`
            ErrorFileS3URL string  `json:"error_file_s3_url"`
            S3URL string  `json:"s3_url"`
            FileName string  `json:"file_name"`
            UpdatedAt string  `json:"updated_at"`
            UpdatedBy string  `json:"updated_by"`
            Success bool  `json:"success"`
         
    }
    
    // PincodeMopUpdateAuditHistoryResponseData used by Serviceability
    type PincodeMopUpdateAuditHistoryResponseData struct {

        
            EntityType string  `json:"entity_type"`
            Page PincodeMopUpdateAuditHistoryPaging  `json:"page"`
            Data []PincodeMopUpdateAuditHistoryResponse  `json:"data"`
         
    }
    
    // BulkRegionJobSerializer used by Serviceability
    type BulkRegionJobSerializer struct {

        
            BatchID string  `json:"batch_id"`
            CountryIsoCode string  `json:"country_iso_code"`
            Action string  `json:"action"`
            JobAction string  `json:"job_action"`
            FileURL string  `json:"file_url"`
         
    }
    
    // PostBulkRegionJobResponse used by Serviceability
    type PostBulkRegionJobResponse struct {

        
            BatchID string  `json:"batch_id"`
            EventEmitted bool  `json:"event_emitted"`
            Response bool  `json:"response"`
            Message string  `json:"message"`
         
    }
    
    // BulkRecordError used by Serviceability
    type BulkRecordError struct {

        
            Error []string  `json:"error"`
            IsError bool  `json:"is_error"`
         
    }
    
    // CSVFileRecord used by Serviceability
    type CSVFileRecord struct {

        
            IsError bool  `json:"is_error"`
            RegionType string  `json:"region_type"`
            Country string  `json:"country"`
            ToRegion string  `json:"to_region"`
            Error []string  `json:"error"`
            FromRegion string  `json:"from_region"`
            TatType string  `json:"tat_type"`
            DpID float64  `json:"dp_id"`
            MaxTat float64  `json:"max_tat"`
            PlanID float64  `json:"plan_id"`
            MinTat float64  `json:"min_tat"`
            SNo float64  `json:"s_no"`
         
    }
    
    // BulkRegionData used by Serviceability
    type BulkRegionData struct {

        
            FilePath string  `json:"file_path"`
            BatchID string  `json:"batch_id"`
            Error BulkRecordError  `json:"error"`
            Action string  `json:"action"`
            FailedCount float64  `json:"failed_count"`
            SuccessCount float64  `json:"success_count"`
            FailedRec []CSVFileRecord  `json:"failed_rec"`
            Stage string  `json:"stage"`
            TotalRec float64  `json:"total_rec"`
            CreatedOn string  `json:"created_on"`
         
    }
    
    // GetBulkRegionJobResponse used by Serviceability
    type GetBulkRegionJobResponse struct {

        
            BatchID string  `json:"batch_id"`
            Data []BulkRegionData  `json:"data"`
            CurrentPageNumber float64  `json:"current_page_number"`
         
    }
    
    // Dp1 used by Serviceability
    type Dp1 struct {

        
            IsSelfShip bool  `json:"is_self_ship"`
            PlanRules map[string]interface{}  `json:"plan_rules"`
            DpID string  `json:"dp_id"`
            AccountID string  `json:"account_id"`
            PlanID string  `json:"plan_id"`
            Stage string  `json:"stage"`
            Name string  `json:"name"`
         
    }
    
    // CompanyDpAccountRequest used by Serviceability
    type CompanyDpAccountRequest struct {

        
            Data []Dp1  `json:"data"`
         
    }
    
    // CompanyDpAccountResponse used by Serviceability
    type CompanyDpAccountResponse struct {

        
            Success bool  `json:"success"`
         
    }
    
    // ErrorResponse used by Serviceability
    type ErrorResponse struct {

        
            Value string  `json:"value"`
            Type string  `json:"type"`
            Message string  `json:"message"`
         
    }
    
    // DpAccountFailureResponse used by Serviceability
    type DpAccountFailureResponse struct {

        
            Error []ErrorResponse  `json:"error"`
            Success bool  `json:"success"`
            StatusCode float64  `json:"status_code"`
         
    }
    
    // Page used by Serviceability
    type Page struct {

        
            ItemTotal float64  `json:"item_total"`
            NextID string  `json:"next_id"`
            HasPrevious bool  `json:"has_previous"`
            HasNext bool  `json:"has_next"`
            Current float64  `json:"current"`
            Type string  `json:"type"`
            Size float64  `json:"size"`
         
    }
    
    // CompanyDpAccountListResponse used by Serviceability
    type CompanyDpAccountListResponse struct {

        
            Page Page  `json:"page"`
            Items []Dp1  `json:"items"`
            Success bool  `json:"success"`
         
    }
    
    // DpRulesUpdateRequest used by Serviceability
    type DpRulesUpdateRequest struct {

        
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            DpIds map[string]map[string]interface{}  `json:"dp_ids"`
            Conditions []map[string]interface{}  `json:"conditions"`
         
    }
    
    // DpRuleResponse used by Serviceability
    type DpRuleResponse struct {

        
            CreatedBy map[string]interface{}  `json:"created_by"`
            IsActive bool  `json:"is_active"`
            CompanyID float64  `json:"company_id"`
            Conditions []string  `json:"conditions"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            UID string  `json:"uid"`
            Name string  `json:"name"`
            DpIds map[string]interface{}  `json:"dp_ids"`
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // DpRuleUpdateSuccessResponse used by Serviceability
    type DpRuleUpdateSuccessResponse struct {

        
            Data DpRuleResponse  `json:"data"`
            Success bool  `json:"success"`
            StatusCode float64  `json:"status_code"`
         
    }
    
    // FailureResponse used by Serviceability
    type FailureResponse struct {

        
            Error []ErrorResponse  `json:"error"`
            Success bool  `json:"success"`
            StatusCode float64  `json:"status_code"`
         
    }
    
    // DpSchemaInRuleListing used by Serviceability
    type DpSchemaInRuleListing struct {

        
            Priority float64  `json:"priority"`
            IsSelfShip bool  `json:"is_self_ship"`
            PlanRules map[string]interface{}  `json:"plan_rules"`
            DpID string  `json:"dp_id"`
            AccountID string  `json:"account_id"`
            PlanID string  `json:"plan_id"`
            Stage string  `json:"stage"`
            Name string  `json:"name"`
         
    }
    
    // DpRule used by Serviceability
    type DpRule struct {

        
            IsActive bool  `json:"is_active"`
            CompanyID float64  `json:"company_id"`
            Conditions []map[string]interface{}  `json:"conditions"`
            Name string  `json:"name"`
            DpIds map[string]DpSchemaInRuleListing  `json:"dp_ids"`
         
    }
    
    // DpRuleSuccessResponse used by Serviceability
    type DpRuleSuccessResponse struct {

        
            Data DpRule  `json:"data"`
            Success bool  `json:"success"`
            StatusCode float64  `json:"status_code"`
         
    }
    
    // DpIds used by Serviceability
    type DpIds struct {

        
            Enabled bool  `json:"enabled"`
            Meta map[string]interface{}  `json:"meta"`
            Priority float64  `json:"priority"`
         
    }
    
    // DpRuleRequest used by Serviceability
    type DpRuleRequest struct {

        
            IsActive bool  `json:"is_active"`
            CompanyID float64  `json:"company_id"`
            Conditions []map[string]interface{}  `json:"conditions"`
            Name string  `json:"name"`
            DpIds map[string]DpIds  `json:"dp_ids"`
         
    }
    
    // DpMultipleRuleSuccessResponse used by Serviceability
    type DpMultipleRuleSuccessResponse struct {

        
            Page Page  `json:"page"`
            Items []DpRule  `json:"items"`
            Success bool  `json:"success"`
         
    }
    
    // DPCompanyRuleRequest used by Serviceability
    type DPCompanyRuleRequest struct {

        
            RuleIds []string  `json:"rule_ids"`
         
    }
    
    // DPCompanyRuleResponse used by Serviceability
    type DPCompanyRuleResponse struct {

        
            Data []DpRuleResponse  `json:"data"`
            Success bool  `json:"success"`
            StatusCode float64  `json:"status_code"`
         
    }
    
    // DPApplicationRuleRequest used by Serviceability
    type DPApplicationRuleRequest struct {

        
            ShippingRules []string  `json:"shipping_rules"`
         
    }
    
    // DPApplicationRuleResponse used by Serviceability
    type DPApplicationRuleResponse struct {

        
            Data []DpRuleResponse  `json:"data"`
            Success bool  `json:"success"`
            StatusCode bool  `json:"status_code"`
         
    }
    
    // SelfShipResponse used by Serviceability
    type SelfShipResponse struct {

        
            Tat float64  `json:"tat"`
            Active bool  `json:"active"`
         
    }
    
    // ApplicationSelfShipConfig used by Serviceability
    type ApplicationSelfShipConfig struct {

        
            SelfShip SelfShipResponse  `json:"self_ship"`
         
    }
    
    // ApplicationSelfShipConfigResponse used by Serviceability
    type ApplicationSelfShipConfigResponse struct {

        
            ID string  `json:"id"`
            SelfShip ApplicationSelfShipConfig  `json:"self_ship"`
            Error ServiceabilityErrorResponse  `json:"error"`
            Success bool  `json:"success"`
         
    }
    

    
    // GenerateReportFilters used by Finance
    type GenerateReportFilters struct {

        
            Company []string  `json:"company"`
            Channel []string  `json:"channel"`
            Brand []string  `json:"brand"`
         
    }
    
    // GenerateReportMeta used by Finance
    type GenerateReportMeta struct {

        
            Company string  `json:"company"`
            Channel string  `json:"channel"`
            Brand string  `json:"brand"`
         
    }
    
    // GenerateReportPlatform used by Finance
    type GenerateReportPlatform struct {

        
            Filters GenerateReportFilters  `json:"filters"`
            Meta GenerateReportMeta  `json:"meta"`
            StartDate string  `json:"start_date"`
            EndDate string  `json:"end_date"`
            ReportID string  `json:"report_id"`
         
    }
    
    // GenerateReportRequest used by Finance
    type GenerateReportRequest struct {

        
            Data GenerateReportPlatform  `json:"data"`
         
    }
    
    // Page used by Finance
    type Page struct {

        
            ItemTotal float64  `json:"item_total"`
            NextID string  `json:"next_id"`
            HasPrevious bool  `json:"has_previous"`
            HasNext bool  `json:"has_next"`
            Current float64  `json:"current"`
            Type string  `json:"type"`
            Size float64  `json:"size"`
         
    }
    
    // GenerateReportJson used by Finance
    type GenerateReportJson struct {

        
            ItemCount float64  `json:"item_count"`
            Items [][]string  `json:"items"`
            Headers []string  `json:"headers"`
            StartDate string  `json:"start_date"`
            EndDate string  `json:"end_date"`
            Page Page  `json:"page"`
         
    }
    
    // Error used by Finance
    type Error struct {

        
            Success bool  `json:"success"`
            Reason string  `json:"reason"`
         
    }
    
    // DownloadReport used by Finance
    type DownloadReport struct {

        
            EndDate string  `json:"end_date"`
            Page float64  `json:"page"`
            Pagesize float64  `json:"pagesize"`
            StartDate string  `json:"start_date"`
         
    }
    
    // DownloadReportItems used by Finance
    type DownloadReportItems struct {

        
            Filters GenerateReportFilters  `json:"filters"`
            Meta GenerateReportMeta  `json:"meta"`
            TypeOfRequest string  `json:"type_of_request"`
            StartDate string  `json:"start_date"`
            EndDate string  `json:"end_date"`
            ReportID string  `json:"report_id"`
         
    }
    
    // DownloadReportList used by Finance
    type DownloadReportList struct {

        
            Page Page  `json:"page"`
            Items []DownloadReportItems  `json:"items"`
            ItemCount float64  `json:"item_count"`
         
    }
    
    // GetEngineFilters used by Finance
    type GetEngineFilters struct {

        
            ConfigField string  `json:"config_field"`
         
    }
    
    // GetEngineData used by Finance
    type GetEngineData struct {

        
            Filters GetEngineFilters  `json:"filters"`
            TableName string  `json:"table_name"`
            Project []string  `json:"project"`
         
    }
    
    // GetEngineRequest used by Finance
    type GetEngineRequest struct {

        
            Data GetEngineData  `json:"data"`
         
    }
    
    // GetEngineResponse used by Finance
    type GetEngineResponse struct {

        
            Page Page  `json:"page"`
            Success bool  `json:"success"`
            Items []map[string]interface{}  `json:"items"`
            ItemCount float64  `json:"item_count"`
         
    }
    
    // GetReason used by Finance
    type GetReason struct {

        
            ReasonType string  `json:"reason_type"`
         
    }
    
    // GetReasonRequest used by Finance
    type GetReasonRequest struct {

        
            Data GetReason  `json:"data"`
         
    }
    
    // GetDocs used by Finance
    type GetDocs struct {

        
            Docs []map[string]interface{}  `json:"docs"`
            Items []map[string]interface{}  `json:"items"`
         
    }
    
    // GetReasonResponse used by Finance
    type GetReasonResponse struct {

        
            Success bool  `json:"success"`
            Data GetDocs  `json:"data"`
         
    }
    
    // GetReportListData used by Finance
    type GetReportListData struct {

        
            RoleName string  `json:"role_name"`
            ListingEnabled bool  `json:"listing_enabled"`
         
    }
    
    // GetReportListRequest used by Finance
    type GetReportListRequest struct {

        
            Data GetReportListData  `json:"data"`
         
    }
    
    // GetAffiliate used by Finance
    type GetAffiliate struct {

        
            CompanyID float64  `json:"company_id"`
         
    }
    
    // GetAffiliateResponse used by Finance
    type GetAffiliateResponse struct {

        
            Success bool  `json:"success"`
            Docs []map[string]interface{}  `json:"docs"`
         
    }
    
    // DownloadCreditDebitNote used by Finance
    type DownloadCreditDebitNote struct {

        
            NoteID []string  `json:"note_id"`
         
    }
    
    // DownloadCreditDebitNoteRequest used by Finance
    type DownloadCreditDebitNoteRequest struct {

        
            Data DownloadCreditDebitNote  `json:"data"`
         
    }
    
    // DownloadCreditDebitNoteResponseData used by Finance
    type DownloadCreditDebitNoteResponseData struct {

        
            ID string  `json:"id"`
            PdfS3URL string  `json:"pdf_s3_url"`
         
    }
    
    // DownloadCreditDebitNoteResponse used by Finance
    type DownloadCreditDebitNoteResponse struct {

        
            Success bool  `json:"success"`
            Data []DownloadCreditDebitNoteResponseData  `json:"data"`
         
    }
    
    // PaymentProcessPayload used by Finance
    type PaymentProcessPayload struct {

        
            TotalAmount string  `json:"total_amount"`
            ModeOfPayment string  `json:"mode_of_payment"`
            Meta map[string]interface{}  `json:"meta"`
            SourceReference string  `json:"source_reference"`
            Amount string  `json:"amount"`
            InvoiceNumber string  `json:"invoice_number"`
            Platform string  `json:"platform"`
            TransactionType string  `json:"transaction_type"`
            Currency string  `json:"currency"`
            SellerID string  `json:"seller_id"`
         
    }
    
    // PaymentProcessRequest used by Finance
    type PaymentProcessRequest struct {

        
            Data PaymentProcessPayload  `json:"data"`
         
    }
    
    // PaymentProcessResponse used by Finance
    type PaymentProcessResponse struct {

        
            Message string  `json:"message"`
            Code float64  `json:"code"`
            Meta map[string]interface{}  `json:"meta"`
            TransactionID string  `json:"transaction_id"`
            RedirectURL string  `json:"redirect_url"`
         
    }
    
    // CreditlineDataPlatformPayload used by Finance
    type CreditlineDataPlatformPayload struct {

        
            EndEnd string  `json:"end_end"`
            StartEnd string  `json:"start_end"`
            Page float64  `json:"page"`
            Pagesize float64  `json:"pagesize"`
            SellerID string  `json:"seller_id"`
         
    }
    
    // CreditlineDataPlatformRequest used by Finance
    type CreditlineDataPlatformRequest struct {

        
            Data CreditlineDataPlatformPayload  `json:"data"`
         
    }
    
    // CreditlineDataPlatformResponse used by Finance
    type CreditlineDataPlatformResponse struct {

        
            Message string  `json:"message"`
            Code float64  `json:"code"`
            ItemCount float64  `json:"item_count"`
            Items []map[string]interface{}  `json:"items"`
            ShowMr bool  `json:"show_mr"`
            Headers []string  `json:"headers"`
            Page map[string]interface{}  `json:"page"`
         
    }
    
    // IsCreditlinePayload used by Finance
    type IsCreditlinePayload struct {

        
            SellerID string  `json:"seller_id"`
         
    }
    
    // IsCreditlinePlatformRequest used by Finance
    type IsCreditlinePlatformRequest struct {

        
            Data IsCreditlinePayload  `json:"data"`
         
    }
    
    // IsCreditlinePlatformResponse used by Finance
    type IsCreditlinePlatformResponse struct {

        
            IsCreditlineOpted bool  `json:"is_creditline_opted"`
            Code float64  `json:"code"`
         
    }
    
    // InvoiceTypePayloadData used by Finance
    type InvoiceTypePayloadData struct {

        
            IsActive bool  `json:"is_active"`
         
    }
    
    // InvoiceTypeRequest used by Finance
    type InvoiceTypeRequest struct {

        
            Data InvoiceTypePayloadData  `json:"data"`
         
    }
    
    // InvoiceTypeResponseItems used by Finance
    type InvoiceTypeResponseItems struct {

        
            Text string  `json:"text"`
            Value string  `json:"value"`
         
    }
    
    // InvoiceTypeResponse used by Finance
    type InvoiceTypeResponse struct {

        
            PaymentStatusList []InvoiceTypeResponseItems  `json:"payment_status_list"`
            Success bool  `json:"success"`
            InvoiceTypeList []InvoiceTypeResponseItems  `json:"invoice_type_list"`
         
    }
    
    // InoviceListingPayloadDataFilters used by Finance
    type InoviceListingPayloadDataFilters struct {

        
            CompanyID []string  `json:"company_id"`
            InvoiceType []string  `json:"invoice_type"`
            PaymentStatus []string  `json:"payment_status"`
         
    }
    
    // InvoiceListingPayloadData used by Finance
    type InvoiceListingPayloadData struct {

        
            Search string  `json:"search"`
            Filters InoviceListingPayloadDataFilters  `json:"filters"`
            PageSize float64  `json:"page_size"`
            StartDate string  `json:"start_date"`
            EndDate string  `json:"end_date"`
            Page float64  `json:"page"`
         
    }
    
    // InvoiceListingRequest used by Finance
    type InvoiceListingRequest struct {

        
            Data InvoiceListingPayloadData  `json:"data"`
         
    }
    
    // InvoiceListingResponseItems used by Finance
    type InvoiceListingResponseItems struct {

        
            Period string  `json:"period"`
            InvoiceDate string  `json:"invoice_date"`
            Amount string  `json:"amount"`
            DueDate string  `json:"due_date"`
            IsDownloadable bool  `json:"is_downloadable"`
            InvoiceNumber string  `json:"invoice_number"`
            InvoiceID string  `json:"invoice_id"`
            Company string  `json:"company"`
            InvoiceType string  `json:"invoice_type"`
            Status string  `json:"status"`
         
    }
    
    // UnpaidInvoiceDataItems used by Finance
    type UnpaidInvoiceDataItems struct {

        
            TotalUnpaidInvoiceCount float64  `json:"total_unpaid_invoice_count"`
            Currency string  `json:"currency"`
            TotalUnpaidAmount float64  `json:"total_unpaid_amount"`
         
    }
    
    // InvoiceListingResponse used by Finance
    type InvoiceListingResponse struct {

        
            Page Page  `json:"page"`
            Items []InvoiceListingResponseItems  `json:"items"`
            UnpaidInvoiceData UnpaidInvoiceDataItems  `json:"unpaid_invoice_data"`
            ItemCount float64  `json:"item_count"`
         
    }
    
    // InvoicePdfPayloadData used by Finance
    type InvoicePdfPayloadData struct {

        
            InvoiceNumber []string  `json:"invoice_number"`
         
    }
    
    // InvoicePdfRequest used by Finance
    type InvoicePdfRequest struct {

        
            Data InvoicePdfPayloadData  `json:"data"`
         
    }
    
    // InvoicePdfResponse used by Finance
    type InvoicePdfResponse struct {

        
            Error []string  `json:"error"`
            Success bool  `json:"success"`
            Data []string  `json:"data"`
         
    }
    
