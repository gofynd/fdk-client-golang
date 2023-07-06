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
    
    // EmailTemplateAdmin used by Communication
    type EmailTemplateAdmin struct {

        
            IsSystem bool  `json:"is_system"`
            IsInternal bool  `json:"is_internal"`
            Name string  `json:"name"`
            Description string  `json:"description"`
            EditorType string  `json:"editor_type"`
            EditorMeta string  `json:"editor_meta"`
            StaticTo []string  `json:"static_to"`
            StaticCc []string  `json:"static_cc"`
            StaticBcc []string  `json:"static_bcc"`
            ReplyTo string  `json:"reply_to"`
            Tags []string  `json:"tags"`
            Subject TemplateAndType  `json:"subject"`
            Html TemplateAndType  `json:"html"`
            URLShorten EnabledObj  `json:"url_shorten"`
            Priority string  `json:"priority"`
            TemplateVariables map[string]interface{}  `json:"template_variables"`
            Published bool  `json:"published"`
            Category string  `json:"category"`
            ID string  `json:"_id"`
            Headers []EmailTemplateHeaders  `json:"headers"`
            Attachments []string  `json:"attachments"`
            CreatedAt string  `json:"created_at"`
            UpdatedAt string  `json:"updated_at"`
            Slug string  `json:"slug"`
            V float64  `json:"__v"`
            FromName string  `json:"from_name"`
            Text TemplateAndType  `json:"text"`
         
    }
    
    // EmailTemplatesAdmin used by Communication
    type EmailTemplatesAdmin struct {

        
            Items []EmailTemplateAdmin  `json:"items"`
            Page Page  `json:"page"`
         
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
            FromName string  `json:"from_name"`
            StaticTo []string  `json:"static_to"`
            StaticCc []string  `json:"static_cc"`
            StaticBcc []string  `json:"static_bcc"`
            ReplyTo string  `json:"reply_to"`
            Priority string  `json:"priority"`
            Tags []string  `json:"tags"`
            TemplateVariables map[string]interface{}  `json:"template_variables"`
            Published bool  `json:"published"`
            Subject TemplateAndType  `json:"subject"`
            Html TemplateAndType  `json:"html"`
            EditorType string  `json:"editor_type"`
            EditorMeta string  `json:"editor_meta"`
            Attachments []string  `json:"attachments"`
            Headers []EmailTemplateHeaders  `json:"headers"`
            Keys EmailTemplateKeys  `json:"keys"`
            Text TemplateAndType  `json:"text"`
         
    }
    
    // TemplateAndType used by Communication
    type TemplateAndType struct {

        
            TemplateType string  `json:"template_type"`
            Template string  `json:"template"`
         
    }
    
    // EmailTemplate used by Communication
    type EmailTemplate struct {

        
            Application string  `json:"application"`
            IsSystem bool  `json:"is_system"`
            IsInternal bool  `json:"is_internal"`
            Name string  `json:"name"`
            Description string  `json:"description"`
            EditorType string  `json:"editor_type"`
            EditorMeta string  `json:"editor_meta"`
            StaticTo []string  `json:"static_to"`
            StaticCc []string  `json:"static_cc"`
            StaticBcc []string  `json:"static_bcc"`
            ReplyTo string  `json:"reply_to"`
            Tags []string  `json:"tags"`
            Subject TemplateAndType  `json:"subject"`
            Html TemplateAndType  `json:"html"`
            URLShorten EnabledObj  `json:"url_shorten"`
            Priority string  `json:"priority"`
            TemplateVariables map[string]interface{}  `json:"template_variables"`
            Published bool  `json:"published"`
            Category string  `json:"category"`
            ID string  `json:"_id"`
            Headers []EmailTemplateHeaders  `json:"headers"`
            Attachments []string  `json:"attachments"`
            CreatedAt string  `json:"created_at"`
            UpdatedAt string  `json:"updated_at"`
            Slug string  `json:"slug"`
            V float64  `json:"__v"`
            FromName string  `json:"from_name"`
            Text TemplateAndType  `json:"text"`
         
    }
    
    // SystemEmailTemplate used by Communication
    type SystemEmailTemplate struct {

        
            IsSystem bool  `json:"is_system"`
            IsInternal bool  `json:"is_internal"`
            Name string  `json:"name"`
            Description string  `json:"description"`
            StaticTo []string  `json:"static_to"`
            StaticCc []string  `json:"static_cc"`
            StaticBcc []string  `json:"static_bcc"`
            Tags []string  `json:"tags"`
            Subject TemplateAndType  `json:"subject"`
            Html TemplateAndType  `json:"html"`
            URLShorten EnabledObj  `json:"url_shorten"`
            Priority string  `json:"priority"`
            TemplateVariables map[string]interface{}  `json:"template_variables"`
            Published bool  `json:"published"`
            Category string  `json:"category"`
            ID string  `json:"_id"`
            Headers []EmailTemplateHeaders  `json:"headers"`
            Attachments []string  `json:"attachments"`
            CreatedAt string  `json:"created_at"`
            UpdatedAt string  `json:"updated_at"`
            Slug string  `json:"slug"`
            V float64  `json:"__v"`
            Text TemplateAndType  `json:"text"`
         
    }
    
    // EmailTemplates used by Communication
    type EmailTemplates struct {

        
            Items []EmailTemplate  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // SystemEmailTemplates used by Communication
    type SystemEmailTemplates struct {

        
            Items []SystemEmailTemplate  `json:"items"`
         
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
    
    // SmsTemplatesAdmin used by Communication
    type SmsTemplatesAdmin struct {

        
            Items []SmsTemplateAdmin  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // SmsTemplateAdmin used by Communication
    type SmsTemplateAdmin struct {

        
            URLShorten EnabledObj  `json:"url_shorten"`
            ID string  `json:"_id"`
            IsSystem bool  `json:"is_system"`
            IsInternal bool  `json:"is_internal"`
            Meta metaObj  `json:"meta"`
            Name string  `json:"name"`
            Description string  `json:"description"`
            Message SmsTemplateMessage  `json:"message"`
            Priority string  `json:"priority"`
            Tags []string  `json:"tags"`
            TemplateVariables map[string]interface{}  `json:"template_variables"`
            TemplateID string  `json:"template_id"`
            Published bool  `json:"published"`
            Category string  `json:"category"`
            CreatedAt string  `json:"created_at"`
            UpdatedAt string  `json:"updated_at"`
            Slug string  `json:"slug"`
            V float64  `json:"__v"`
         
    }
    
    // SmsTemplateMessage used by Communication
    type SmsTemplateMessage struct {

        
            TemplateType string  `json:"template_type"`
            Template string  `json:"template"`
         
    }
    
    // SmsTemplates used by Communication
    type SmsTemplates struct {

        
            Items []SmsTemplate  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // SmsTemplate used by Communication
    type SmsTemplate struct {

        
            URLShorten EnabledObj  `json:"url_shorten"`
            ID string  `json:"_id"`
            Application string  `json:"application"`
            IsSystem bool  `json:"is_system"`
            IsInternal bool  `json:"is_internal"`
            Meta metaObj  `json:"meta"`
            Name string  `json:"name"`
            Description string  `json:"description"`
            Message SmsTemplateMessage  `json:"message"`
            Priority string  `json:"priority"`
            Tags []string  `json:"tags"`
            TemplateVariables map[string]interface{}  `json:"template_variables"`
            TemplateID string  `json:"template_id"`
            Published bool  `json:"published"`
            Category string  `json:"category"`
            CreatedAt string  `json:"created_at"`
            UpdatedAt string  `json:"updated_at"`
            Slug string  `json:"slug"`
            V float64  `json:"__v"`
         
    }
    
    // SystemSmsTemplates used by Communication
    type SystemSmsTemplates struct {

        
            URLShorten EnabledObj  `json:"url_shorten"`
            ID string  `json:"_id"`
            IsSystem bool  `json:"is_system"`
            IsInternal bool  `json:"is_internal"`
            Name string  `json:"name"`
            Description string  `json:"description"`
            Message SmsTemplateMessage  `json:"message"`
            Priority string  `json:"priority"`
            Tags []string  `json:"tags"`
            TemplateVariables map[string]interface{}  `json:"template_variables"`
            TemplateID string  `json:"template_id"`
            Published bool  `json:"published"`
            Category string  `json:"category"`
            CreatedAt string  `json:"created_at"`
            UpdatedAt string  `json:"updated_at"`
            Slug string  `json:"slug"`
            V float64  `json:"__v"`
         
    }
    
    // metaObj used by Communication
    type metaObj struct {

        
            Type string  `json:"type"`
            IsSystem bool  `json:"is_system"`
            Template string  `json:"template"`
         
    }
    
    // SmsTemplateReq used by Communication
    type SmsTemplateReq struct {

        
            Name string  `json:"name"`
            Description string  `json:"description"`
            Priority string  `json:"priority"`
            TemplateID string  `json:"template_id"`
            Meta metaObj  `json:"meta"`
            TemplateVariables map[string]interface{}  `json:"template_variables"`
            Published bool  `json:"published"`
            Message SmsTemplateMessage  `json:"message"`
         
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
    
    // GenericPage used by Communication
    type GenericPage struct {

        
            Type string  `json:"type"`
            Size float64  `json:"size"`
            Current float64  `json:"current"`
            HasNext bool  `json:"has_next"`
            ItemTotal float64  `json:"item_total"`
         
    }
    
    // GenericSuccess used by Communication
    type GenericSuccess struct {

        
            Success bool  `json:"success"`
         
    }
    
    // GenericError used by Communication
    type GenericError struct {

        
            Message Message  `json:"message"`
            Sentry string  `json:"sentry"`
         
    }
    
    // GenericDelete used by Communication
    type GenericDelete struct {

        
            Message string  `json:"message"`
            Acknowledged bool  `json:"acknowledged"`
            Affected float64  `json:"affected"`
            Operation string  `json:"operation"`
         
    }
    
    // Message used by Communication
    type Message struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
            Info string  `json:"info"`
            Operation string  `json:"operation"`
         
    }
    
    // EnabledObj used by Communication
    type EnabledObj struct {

        
            Enabled bool  `json:"enabled"`
         
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
            AppID string  `json:"app_id"`
            Success bool  `json:"success"`
            ExcludedFields []string  `json:"excluded_fields"`
            Created bool  `json:"created"`
            DisplayFields []string  `json:"display_fields"`
         
    }
    
    // ErrorCodeDescription used by Payment
    type ErrorCodeDescription struct {

        
            Code string  `json:"code"`
            Success bool  `json:"success"`
            Description string  `json:"description"`
         
    }
    
    // PaymentGatewayConfig used by Payment
    type PaymentGatewayConfig struct {

        
            IsActive bool  `json:"is_active"`
            Key string  `json:"key"`
            ConfigType string  `json:"config_type"`
            Secret string  `json:"secret"`
            MerchantSalt string  `json:"merchant_salt"`
         
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

        
            Code string  `json:"code"`
            Description string  `json:"description"`
         
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
            Code string  `json:"code"`
            Logos PaymentModeLogo  `json:"logos"`
            PackageName string  `json:"package_name"`
         
    }
    
    // IntentAppErrorList used by Payment
    type IntentAppErrorList struct {

        
            Code string  `json:"code"`
            PackageName string  `json:"package_name"`
         
    }
    
    // PaymentModeList used by Payment
    type PaymentModeList struct {

        
            CodLimit float64  `json:"cod_limit"`
            RemainingLimit float64  `json:"remaining_limit"`
            IntentFlow bool  `json:"intent_flow"`
            CardType string  `json:"card_type"`
            AggregatorName string  `json:"aggregator_name"`
            Expired bool  `json:"expired"`
            CardIsin string  `json:"card_isin"`
            CardFingerprint string  `json:"card_fingerprint"`
            IntentApp []IntentApp  `json:"intent_app"`
            LogoURL PaymentModeLogo  `json:"logo_url"`
            Name string  `json:"name"`
            CardBrandImage string  `json:"card_brand_image"`
            CodLimitPerOrder float64  `json:"cod_limit_per_order"`
            Nickname string  `json:"nickname"`
            DisplayName string  `json:"display_name"`
            CardName string  `json:"card_name"`
            CardReference string  `json:"card_reference"`
            Timeout float64  `json:"timeout"`
            Code string  `json:"code"`
            ExpMonth float64  `json:"exp_month"`
            DisplayPriority float64  `json:"display_priority"`
            FyndVpa string  `json:"fynd_vpa"`
            CardID string  `json:"card_id"`
            IntentAppErrorDictList []IntentAppErrorList  `json:"intent_app_error_dict_list"`
            IntentAppErrorList []string  `json:"intent_app_error_list"`
            CardIssuer string  `json:"card_issuer"`
            CardToken string  `json:"card_token"`
            RetryCount float64  `json:"retry_count"`
            CardNumber string  `json:"card_number"`
            MerchantCode string  `json:"merchant_code"`
            ExpYear float64  `json:"exp_year"`
            CardBrand string  `json:"card_brand"`
            CompliantWithTokenisationGuidelines bool  `json:"compliant_with_tokenisation_guidelines"`
         
    }
    
    // RootPaymentMode used by Payment
    type RootPaymentMode struct {

        
            SaveCard bool  `json:"save_card"`
            AggregatorName string  `json:"aggregator_name"`
            DisplayPriority float64  `json:"display_priority"`
            DisplayName string  `json:"display_name"`
            IsPayByCardPl bool  `json:"is_pay_by_card_pl"`
            AnonymousEnable bool  `json:"anonymous_enable"`
            AddCardEnabled bool  `json:"add_card_enabled"`
            Name string  `json:"name"`
            List []PaymentModeList  `json:"list"`
         
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
    
    // PayoutsResponse used by Payment
    type PayoutsResponse struct {

        
            IsActive bool  `json:"is_active"`
            PayoutsAggregators []map[string]interface{}  `json:"payouts_aggregators"`
            MoreAttributes map[string]interface{}  `json:"more_attributes"`
            UniqueTransferNo map[string]interface{}  `json:"unique_transfer_no"`
            TransferType string  `json:"transfer_type"`
            IsDefault bool  `json:"is_default"`
            Customers map[string]interface{}  `json:"customers"`
         
    }
    
    // PayoutBankDetails used by Payment
    type PayoutBankDetails struct {

        
            BranchName string  `json:"branch_name"`
            City string  `json:"city"`
            Pincode float64  `json:"pincode"`
            AccountNo string  `json:"account_no"`
            AccountHolder string  `json:"account_holder"`
            State string  `json:"state"`
            AccountType string  `json:"account_type"`
            IfscCode string  `json:"ifsc_code"`
            BankName string  `json:"bank_name"`
            Country string  `json:"country"`
         
    }
    
    // PayoutRequest used by Payment
    type PayoutRequest struct {

        
            IsActive bool  `json:"is_active"`
            BankDetails PayoutBankDetails  `json:"bank_details"`
            Users map[string]interface{}  `json:"users"`
            UniqueExternalID string  `json:"unique_external_id"`
            TransferType string  `json:"transfer_type"`
            Aggregator string  `json:"aggregator"`
         
    }
    
    // PayoutResponse used by Payment
    type PayoutResponse struct {

        
            IsActive bool  `json:"is_active"`
            BankDetails map[string]interface{}  `json:"bank_details"`
            Payouts map[string]interface{}  `json:"payouts"`
            Users map[string]interface{}  `json:"users"`
            Success bool  `json:"success"`
            PaymentStatus string  `json:"payment_status"`
            UniqueTransferNo string  `json:"unique_transfer_no"`
            Created bool  `json:"created"`
            TransferType string  `json:"transfer_type"`
            Aggregator string  `json:"aggregator"`
         
    }
    
    // UpdatePayoutResponse used by Payment
    type UpdatePayoutResponse struct {

        
            IsActive bool  `json:"is_active"`
            IsDefault bool  `json:"is_default"`
            Success bool  `json:"success"`
         
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

        
            Data []map[string]interface{}  `json:"data"`
            Success bool  `json:"success"`
         
    }
    
    // DeleteSubscriptionPaymentMethodResponse used by Payment
    type DeleteSubscriptionPaymentMethodResponse struct {

        
            Success bool  `json:"success"`
         
    }
    
    // SubscriptionConfigResponse used by Payment
    type SubscriptionConfigResponse struct {

        
            Config map[string]interface{}  `json:"config"`
            Success bool  `json:"success"`
            Aggregator string  `json:"aggregator"`
         
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

        
            Data map[string]interface{}  `json:"data"`
            Message string  `json:"message"`
            Success bool  `json:"success"`
            IsVerifiedFlag bool  `json:"is_verified_flag"`
         
    }
    
    // NotFoundResourceError used by Payment
    type NotFoundResourceError struct {

        
            Code string  `json:"code"`
            Success bool  `json:"success"`
            Description string  `json:"description"`
         
    }
    
    // BankDetailsForOTP used by Payment
    type BankDetailsForOTP struct {

        
            BranchName string  `json:"branch_name"`
            AccountNo string  `json:"account_no"`
            AccountHolder string  `json:"account_holder"`
            IfscCode string  `json:"ifsc_code"`
            BankName string  `json:"bank_name"`
         
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

        
            BranchName string  `json:"branch_name"`
            ID float64  `json:"id"`
            Email string  `json:"email"`
            CreatedOn string  `json:"created_on"`
            Address string  `json:"address"`
            DelightsUserName string  `json:"delights_user_name"`
            Mobile string  `json:"mobile"`
            Title string  `json:"title"`
            AccountHolder string  `json:"account_holder"`
            DisplayName string  `json:"display_name"`
            Comment string  `json:"comment"`
            IsActive bool  `json:"is_active"`
            Subtitle string  `json:"subtitle"`
            ModifiedOn string  `json:"modified_on"`
            IfscCode string  `json:"ifsc_code"`
            BankName string  `json:"bank_name"`
            BeneficiaryID string  `json:"beneficiary_id"`
            TransferMode string  `json:"transfer_mode"`
            AccountNo string  `json:"account_no"`
         
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
            CurrentStatus string  `json:"current_status"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            OrderID string  `json:"order_id"`
         
    }
    
    // MultiTenderPaymentMethod used by Payment
    type MultiTenderPaymentMethod struct {

        
            Amount float64  `json:"amount"`
            Name string  `json:"name"`
            Mode string  `json:"mode"`
            Meta MultiTenderPaymentMeta  `json:"meta"`
         
    }
    
    // PaymentConfirmationRequest used by Payment
    type PaymentConfirmationRequest struct {

        
            OrderID string  `json:"order_id"`
            PaymentMethods []MultiTenderPaymentMethod  `json:"payment_methods"`
         
    }
    
    // PaymentConfirmationResponse used by Payment
    type PaymentConfirmationResponse struct {

        
            OrderID string  `json:"order_id"`
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // CODdata used by Payment
    type CODdata struct {

        
            IsActive bool  `json:"is_active"`
            UserID string  `json:"user_id"`
            RemainingLimit float64  `json:"remaining_limit"`
            Limit float64  `json:"limit"`
            Usages float64  `json:"usages"`
         
    }
    
    // GetUserCODLimitResponse used by Payment
    type GetUserCODLimitResponse struct {

        
            Success bool  `json:"success"`
            UserCodData CODdata  `json:"user_cod_data"`
         
    }
    
    // SetCODForUserRequest used by Payment
    type SetCODForUserRequest struct {

        
            IsActive bool  `json:"is_active"`
            MerchantUserID string  `json:"merchant_user_id"`
            Mobileno string  `json:"mobileno"`
         
    }
    
    // SetCODOptionResponse used by Payment
    type SetCODOptionResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // EdcModelData used by Payment
    type EdcModelData struct {

        
            Models []string  `json:"models"`
            Aggregator string  `json:"aggregator"`
            AggregatorID float64  `json:"aggregator_id"`
         
    }
    
    // EdcAggregatorAndModelListResponse used by Payment
    type EdcAggregatorAndModelListResponse struct {

        
            Data []EdcModelData  `json:"data"`
            Success bool  `json:"success"`
         
    }
    
    // StatisticsData used by Payment
    type StatisticsData struct {

        
            InactiveDeviceCount float64  `json:"inactive_device_count"`
            ActiveDeviceCount float64  `json:"active_device_count"`
         
    }
    
    // EdcDeviceStatsResponse used by Payment
    type EdcDeviceStatsResponse struct {

        
            Statistics StatisticsData  `json:"statistics"`
            Success bool  `json:"success"`
         
    }
    
    // EdcAddRequest used by Payment
    type EdcAddRequest struct {

        
            TerminalSerialNo string  `json:"terminal_serial_no"`
            DeviceTag string  `json:"device_tag"`
            StoreID float64  `json:"store_id"`
            EdcModel string  `json:"edc_model"`
            EdcDeviceSerialNo string  `json:"edc_device_serial_no"`
            AggregatorID float64  `json:"aggregator_id"`
         
    }
    
    // EdcDevice used by Payment
    type EdcDevice struct {

        
            IsActive bool  `json:"is_active"`
            TerminalSerialNo string  `json:"terminal_serial_no"`
            DeviceTag string  `json:"device_tag"`
            AggregatorID float64  `json:"aggregator_id"`
            AggregatorName string  `json:"aggregator_name"`
            MerchantStorePosCode string  `json:"merchant_store_pos_code"`
            StoreID float64  `json:"store_id"`
            EdcModel string  `json:"edc_model"`
            EdcDeviceSerialNo string  `json:"edc_device_serial_no"`
            TerminalUniqueIdentifier string  `json:"terminal_unique_identifier"`
            ApplicationID string  `json:"application_id"`
         
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

        
            IsActive bool  `json:"is_active"`
            DeviceTag string  `json:"device_tag"`
            StoreID float64  `json:"store_id"`
            EdcModel string  `json:"edc_model"`
            EdcDeviceSerialNo string  `json:"edc_device_serial_no"`
            MerchantStorePosCode string  `json:"merchant_store_pos_code"`
            AggregatorID float64  `json:"aggregator_id"`
         
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

        
            Page Page  `json:"page"`
            Items []EdcDevice  `json:"items"`
            Success bool  `json:"success"`
         
    }
    
    // PaymentInitializationRequest used by Payment
    type PaymentInitializationRequest struct {

        
            Timeout float64  `json:"timeout"`
            Method string  `json:"method"`
            CustomerID string  `json:"customer_id"`
            Amount float64  `json:"amount"`
            Email string  `json:"email"`
            MerchantOrderID string  `json:"merchant_order_id"`
            Vpa string  `json:"vpa"`
            Currency string  `json:"currency"`
            Contact string  `json:"contact"`
            RazorpayPaymentID string  `json:"razorpay_payment_id"`
            DeviceID string  `json:"device_id"`
            OrderID string  `json:"order_id"`
            Aggregator string  `json:"aggregator"`
         
    }
    
    // PaymentInitializationResponse used by Payment
    type PaymentInitializationResponse struct {

        
            Timeout float64  `json:"timeout"`
            Method string  `json:"method"`
            BqrImage string  `json:"bqr_image"`
            CustomerID string  `json:"customer_id"`
            Amount float64  `json:"amount"`
            MerchantOrderID string  `json:"merchant_order_id"`
            Vpa string  `json:"vpa"`
            Status string  `json:"status"`
            Success bool  `json:"success"`
            AggregatorOrderID string  `json:"aggregator_order_id"`
            PollingURL string  `json:"polling_url"`
            Currency string  `json:"currency"`
            RazorpayPaymentID string  `json:"razorpay_payment_id"`
            DeviceID string  `json:"device_id"`
            UpiPollURL string  `json:"upi_poll_url"`
            Aggregator string  `json:"aggregator"`
            VirtualID string  `json:"virtual_id"`
         
    }
    
    // PaymentStatusUpdateRequest used by Payment
    type PaymentStatusUpdateRequest struct {

        
            Method string  `json:"method"`
            CustomerID string  `json:"customer_id"`
            Amount float64  `json:"amount"`
            Email string  `json:"email"`
            MerchantOrderID string  `json:"merchant_order_id"`
            Status string  `json:"status"`
            Vpa string  `json:"vpa"`
            Currency string  `json:"currency"`
            MerchantTransactionID string  `json:"merchant_transaction_id"`
            Contact string  `json:"contact"`
            DeviceID string  `json:"device_id"`
            OrderID string  `json:"order_id"`
            Aggregator string  `json:"aggregator"`
         
    }
    
    // PaymentStatusUpdateResponse used by Payment
    type PaymentStatusUpdateResponse struct {

        
            Status string  `json:"status"`
            Success bool  `json:"success"`
            AggregatorName string  `json:"aggregator_name"`
            Retry bool  `json:"retry"`
            RedirectURL string  `json:"redirect_url"`
         
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
    
    // PaymentStatusBulkHandlerRequest used by Payment
    type PaymentStatusBulkHandlerRequest struct {

        
            MerchantOrderID []string  `json:"merchant_order_id"`
         
    }
    
    // PaymentObjectListSerializer used by Payment
    type PaymentObjectListSerializer struct {

        
            ID string  `json:"id"`
            AggregatorPaymentObject map[string]interface{}  `json:"aggregator_payment_object"`
            CollectedBy string  `json:"collected_by"`
            CompanyID string  `json:"company_id"`
            RefundedBy string  `json:"refunded_by"`
            PaymentMode string  `json:"payment_mode"`
            UserObject map[string]interface{}  `json:"user_object"`
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
            RefundObject map[string]interface{}  `json:"refund_object"`
            PaymentGateway string  `json:"payment_gateway"`
            Currency string  `json:"currency"`
            PaymentID string  `json:"payment_id"`
            CurrentStatus string  `json:"current_status"`
            PaymentModeIdentifier string  `json:"payment_mode_identifier"`
            AmountInPaisa string  `json:"amount_in_paisa"`
            AllStatus []string  `json:"all_status"`
            ApplicationID string  `json:"application_id"`
         
    }
    
    // PaymentStatusObject used by Payment
    type PaymentStatusObject struct {

        
            MerchantOrderID string  `json:"merchant_order_id"`
            PaymentObjectList []PaymentObjectListSerializer  `json:"payment_object_list"`
         
    }
    
    // PaymentStatusBulkHandlerResponse used by Payment
    type PaymentStatusBulkHandlerResponse struct {

        
            Error string  `json:"error"`
            Status float64  `json:"status"`
            Success string  `json:"success"`
            Count float64  `json:"count"`
            Data []PaymentStatusObject  `json:"data"`
         
    }
    
    // GetOauthUrlResponse used by Payment
    type GetOauthUrlResponse struct {

        
            Success bool  `json:"success"`
            URL string  `json:"url"`
         
    }
    
    // RevokeOAuthToken used by Payment
    type RevokeOAuthToken struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // RepaymentRequestDetails used by Payment
    type RepaymentRequestDetails struct {

        
            FwdShipmentID string  `json:"fwd_shipment_id"`
            Amount float64  `json:"amount"`
            MerchantOrderID string  `json:"merchant_order_id"`
            PaymentMode string  `json:"payment_mode"`
            AggregatorTransactionID string  `json:"aggregator_transaction_id"`
            AggregatorOrderID string  `json:"aggregator_order_id"`
            CurrentStatus string  `json:"current_status"`
            PaymentModeIdentifier string  `json:"payment_mode_identifier"`
            OutstandingDetailsID float64  `json:"outstanding_details_id"`
            Aggregator string  `json:"aggregator"`
         
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

        
            Data map[string]interface{}  `json:"data"`
            Success bool  `json:"success"`
         
    }
    
    // MerchantOnBoardingRequest used by Payment
    type MerchantOnBoardingRequest struct {

        
            UserID string  `json:"user_id"`
            AppID string  `json:"app_id"`
            Status string  `json:"status"`
            CreditLineID string  `json:"credit_line_id"`
            Aggregator string  `json:"aggregator"`
         
    }
    
    // MerchantOnBoardingResponse used by Payment
    type MerchantOnBoardingResponse struct {

        
            Data map[string]interface{}  `json:"data"`
            Success bool  `json:"success"`
         
    }
    
    // ValidateCustomerRequest used by Payment
    type ValidateCustomerRequest struct {

        
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
            PhoneNumber string  `json:"phone_number"`
            OrderItems []map[string]interface{}  `json:"order_items"`
            Payload string  `json:"payload"`
            TransactionAmountInPaise float64  `json:"transaction_amount_in_paise"`
            MerchantParams map[string]interface{}  `json:"merchant_params"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
            Aggregator string  `json:"aggregator"`
         
    }
    
    // ValidateCustomerResponse used by Payment
    type ValidateCustomerResponse struct {

        
            Data map[string]interface{}  `json:"data"`
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // GetPaymentLinkResponse used by Payment
    type GetPaymentLinkResponse struct {

        
            PaymentLinkURL string  `json:"payment_link_url"`
            Amount float64  `json:"amount"`
            Success bool  `json:"success"`
            PaymentLinkCurrentStatus string  `json:"payment_link_current_status"`
            StatusCode float64  `json:"status_code"`
            PollingTimeout float64  `json:"polling_timeout"`
            ExternalOrderID string  `json:"external_order_id"`
            Message string  `json:"message"`
            MerchantName string  `json:"merchant_name"`
         
    }
    
    // ErrorDescription used by Payment
    type ErrorDescription struct {

        
            Amount float64  `json:"amount"`
            MerchantOrderID string  `json:"merchant_order_id"`
            Expired bool  `json:"expired"`
            Cancelled bool  `json:"cancelled"`
            Msg string  `json:"msg"`
            PaymentTransactionID string  `json:"payment_transaction_id"`
            InvalidID bool  `json:"invalid_id"`
            MerchantName string  `json:"merchant_name"`
         
    }
    
    // ErrorResponse used by Payment
    type ErrorResponse struct {

        
            Error ErrorDescription  `json:"error"`
            StatusCode float64  `json:"status_code"`
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // CreatePaymentLinkMeta used by Payment
    type CreatePaymentLinkMeta struct {

        
            Pincode string  `json:"pincode"`
            Amount string  `json:"amount"`
            CheckoutMode string  `json:"checkout_mode"`
            CartID string  `json:"cart_id"`
            AssignCardID string  `json:"assign_card_id"`
         
    }
    
    // CreatePaymentLinkRequest used by Payment
    type CreatePaymentLinkRequest struct {

        
            Description string  `json:"description"`
            Amount float64  `json:"amount"`
            Email string  `json:"email"`
            Meta CreatePaymentLinkMeta  `json:"meta"`
            MobileNumber string  `json:"mobile_number"`
            ExternalOrderID string  `json:"external_order_id"`
         
    }
    
    // CreatePaymentLinkResponse used by Payment
    type CreatePaymentLinkResponse struct {

        
            PaymentLinkURL string  `json:"payment_link_url"`
            Success bool  `json:"success"`
            PaymentLinkID string  `json:"payment_link_id"`
            StatusCode float64  `json:"status_code"`
            PollingTimeout float64  `json:"polling_timeout"`
            Message string  `json:"message"`
         
    }
    
    // PollingPaymentLinkResponse used by Payment
    type PollingPaymentLinkResponse struct {

        
            Amount float64  `json:"amount"`
            Status string  `json:"status"`
            AggregatorName string  `json:"aggregator_name"`
            Success bool  `json:"success"`
            PaymentLinkID string  `json:"payment_link_id"`
            StatusCode float64  `json:"status_code"`
            OrderID string  `json:"order_id"`
            Message string  `json:"message"`
            HttpStatus float64  `json:"http_status"`
            RedirectURL string  `json:"redirect_url"`
         
    }
    
    // CancelOrResendPaymentLinkRequest used by Payment
    type CancelOrResendPaymentLinkRequest struct {

        
            PaymentLinkID string  `json:"payment_link_id"`
         
    }
    
    // ResendPaymentLinkResponse used by Payment
    type ResendPaymentLinkResponse struct {

        
            PollingTimeout float64  `json:"polling_timeout"`
            StatusCode float64  `json:"status_code"`
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // CancelPaymentLinkResponse used by Payment
    type CancelPaymentLinkResponse struct {

        
            StatusCode float64  `json:"status_code"`
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // ExtensionPaymentUpdateRequestSerializer used by Payment
    type ExtensionPaymentUpdateRequestSerializer struct {

        
            TotalAmount float64  `json:"total_amount"`
            Status string  `json:"status"`
            PaymentDetails map[string]interface{}  `json:"payment_details"`
            Currency string  `json:"currency"`
            Gid string  `json:"gid"`
            OrderDetails map[string]interface{}  `json:"order_details"`
         
    }
    
    // ExtensionPaymentUpdateResponseSerializer used by Payment
    type ExtensionPaymentUpdateResponseSerializer struct {

        
            TotalAmount float64  `json:"total_amount"`
            Status string  `json:"status"`
            PlatformTransactionDetails map[string]interface{}  `json:"platform_transaction_details"`
            Currency string  `json:"currency"`
            Gid string  `json:"gid"`
         
    }
    
    // Code used by Payment
    type Code struct {

        
            Code string  `json:"code"`
            Name string  `json:"name"`
            MerchantCode string  `json:"merchant_code"`
         
    }
    
    // PaymentCode used by Payment
    type PaymentCode struct {

        
            Types string  `json:"types"`
            Codes Code  `json:"codes"`
            Name string  `json:"name"`
            Networks string  `json:"networks"`
         
    }
    
    // GetPaymentCode used by Payment
    type GetPaymentCode struct {

        
            MethodCode PaymentCode  `json:"method_code"`
         
    }
    
    // GetPaymentCodeResponse used by Payment
    type GetPaymentCodeResponse struct {

        
            Data GetPaymentCode  `json:"data"`
            Success bool  `json:"success"`
         
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
    
    // ShipmentListingChannel used by Order
    type ShipmentListingChannel struct {

        
            IsAffiliate bool  `json:"is_affiliate"`
            ChannelShipmentID string  `json:"channel_shipment_id"`
            Name string  `json:"name"`
            Logo string  `json:"logo"`
         
    }
    
    // ShipmentItemFulFillingStore used by Order
    type ShipmentItemFulFillingStore struct {

        
            Meta map[string]interface{}  `json:"meta"`
            Code string  `json:"code"`
            Address string  `json:"address"`
            Pincode string  `json:"pincode"`
            StoreEmail string  `json:"store_email"`
            ID float64  `json:"id"`
            Name string  `json:"name"`
            City string  `json:"city"`
            Phone string  `json:"phone"`
            BrandStoreTags string  `json:"brand_store_tags"`
            LocationType string  `json:"location_type"`
            State string  `json:"state"`
         
    }
    
    // PlatformDeliveryAddress used by Order
    type PlatformDeliveryAddress struct {

        
            Address1 string  `json:"address1"`
            CreatedAt string  `json:"created_at"`
            ContactPerson string  `json:"contact_person"`
            Pincode string  `json:"pincode"`
            AddressCategory string  `json:"address_category"`
            Email string  `json:"email"`
            Country string  `json:"country"`
            Latitude float64  `json:"latitude"`
            Version string  `json:"version"`
            UpdatedAt string  `json:"updated_at"`
            Area string  `json:"area"`
            Phone string  `json:"phone"`
            City string  `json:"city"`
            Address2 string  `json:"address2"`
            Longitude float64  `json:"longitude"`
            AddressType string  `json:"address_type"`
            Landmark string  `json:"landmark"`
            State string  `json:"state"`
         
    }
    
    // UserDataInfo used by Order
    type UserDataInfo struct {

        
            AvisUserID string  `json:"avis_user_id"`
            Email string  `json:"email"`
            Name string  `json:"name"`
            LastName string  `json:"last_name"`
            UID float64  `json:"uid"`
            Mobile string  `json:"mobile"`
            FirstName string  `json:"first_name"`
            IsAnonymousUser bool  `json:"is_anonymous_user"`
            Gender string  `json:"gender"`
         
    }
    
    // ShipmentStatus used by Order
    type ShipmentStatus struct {

        
            CreatedAt string  `json:"created_at"`
            Meta map[string]interface{}  `json:"meta"`
            ShipmentStatusID float64  `json:"shipment_status_id"`
            Status string  `json:"status"`
            ShipmentID string  `json:"shipment_id"`
            BagList []string  `json:"bag_list"`
            Title string  `json:"title"`
            CurrentShipmentStatus string  `json:"current_shipment_status"`
            StatusCreatedAt string  `json:"status_created_at"`
         
    }
    
    // LockData used by Order
    type LockData struct {

        
            Locked bool  `json:"locked"`
            LockMessage string  `json:"lock_message"`
            Mto bool  `json:"mto"`
         
    }
    
    // Formatted used by Order
    type Formatted struct {

        
            Max string  `json:"max"`
            Min string  `json:"min"`
         
    }
    
    // ShipmentTimeStamp used by Order
    type ShipmentTimeStamp struct {

        
            Max float64  `json:"max"`
            Min float64  `json:"min"`
         
    }
    
    // ShipmentTags used by Order
    type ShipmentTags struct {

        
            Slug string  `json:"slug"`
            EntityType string  `json:"entity_type"`
            DisplayText string  `json:"display_text"`
         
    }
    
    // ShipmentItemMeta used by Order
    type ShipmentItemMeta struct {

        
            ShipmentChargeableWeight float64  `json:"shipment_chargeable_weight"`
            IsInternational bool  `json:"is_international"`
            ShippingZone string  `json:"shipping_zone"`
            DpOptions map[string]interface{}  `json:"dp_options"`
            ExistingDpList []string  `json:"existing_dp_list"`
            OrderType string  `json:"order_type"`
            BagWeight map[string]interface{}  `json:"bag_weight"`
            EwaybillInfo map[string]interface{}  `json:"ewaybill_info"`
            LockData LockData  `json:"lock_data"`
            ShipmentVolumetricWeight float64  `json:"shipment_volumetric_weight"`
            AutoTriggerDpAssignmentAcf bool  `json:"auto_trigger_dp_assignment_acf"`
            External map[string]interface{}  `json:"external"`
            Formatted Formatted  `json:"formatted"`
            PackagingName string  `json:"packaging_name"`
            ActivityComment string  `json:"activity_comment"`
            Sla float64  `json:"sla"`
            ParentDpID string  `json:"parent_dp_id"`
            DebugInfo map[string]interface{}  `json:"debug_info"`
            PdfMedia []map[string]interface{}  `json:"pdf_media"`
            Timestamp ShipmentTimeStamp  `json:"timestamp"`
            DpSortKey string  `json:"dp_sort_key"`
            Weight float64  `json:"weight"`
            AssignDpFromSb bool  `json:"assign_dp_from_sb"`
            ShipmentWeight float64  `json:"shipment_weight"`
            SameStoreAvailable bool  `json:"same_store_available"`
            StoreInvoiceUpdatedDate string  `json:"store_invoice_updated_date"`
            ShipmentTags []ShipmentTags  `json:"shipment_tags"`
            Tags []map[string]interface{}  `json:"tags"`
            FulfilmentPriorityText string  `json:"fulfilment_priority_text"`
         
    }
    
    // BagStateMapper used by Order
    type BagStateMapper struct {

        
            AppFacing bool  `json:"app_facing"`
            JourneyType string  `json:"journey_type"`
            StateType string  `json:"state_type"`
            NotifyCustomer bool  `json:"notify_customer"`
            ID float64  `json:"id"`
            Name string  `json:"name"`
            AppStateName string  `json:"app_state_name"`
            DisplayName string  `json:"display_name"`
            AppDisplayName string  `json:"app_display_name"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // BagStatusHistory used by Order
    type BagStatusHistory struct {

        
            CreatedAt string  `json:"created_at"`
            DeliveryPartnerID float64  `json:"delivery_partner_id"`
            BagStateMapper BagStateMapper  `json:"bag_state_mapper"`
            StateID float64  `json:"state_id"`
            StoreID float64  `json:"store_id"`
            Status string  `json:"status"`
            Forward bool  `json:"forward"`
            AppDisplayName string  `json:"app_display_name"`
            ShipmentID string  `json:"shipment_id"`
            UpdatedAt string  `json:"updated_at"`
            BagID float64  `json:"bag_id"`
            DeliveryAwbNumber string  `json:"delivery_awb_number"`
            Reasons []map[string]interface{}  `json:"reasons"`
            DisplayName string  `json:"display_name"`
            KafkaSync bool  `json:"kafka_sync"`
            StateType string  `json:"state_type"`
            BshID float64  `json:"bsh_id"`
         
    }
    
    // ReplacementDetails used by Order
    type ReplacementDetails struct {

        
            ReplacementType string  `json:"replacement_type"`
            OriginalAffiliateOrderID string  `json:"original_affiliate_order_id"`
         
    }
    
    // AffiliateMeta used by Order
    type AffiliateMeta struct {

        
            ChannelOrderID string  `json:"channel_order_id"`
            CouponCode string  `json:"coupon_code"`
            SizeLevelTotalQty float64  `json:"size_level_total_qty"`
            EmployeeDiscount float64  `json:"employee_discount"`
            Quantity float64  `json:"quantity"`
            OrderItemID string  `json:"order_item_id"`
            MarketplaceInvoiceID string  `json:"marketplace_invoice_id"`
            BoxType string  `json:"box_type"`
            IsPriority bool  `json:"is_priority"`
            DueDate string  `json:"due_date"`
            ChannelShipmentID string  `json:"channel_shipment_id"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            ReplacementDetails ReplacementDetails  `json:"replacement_details"`
         
    }
    
    // AffiliateBagDetails used by Order
    type AffiliateBagDetails struct {

        
            AffiliateOrderID string  `json:"affiliate_order_id"`
            EmployeeDiscount float64  `json:"employee_discount"`
            AffiliateMeta AffiliateMeta  `json:"affiliate_meta"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
         
    }
    
    // PlatformArticleAttributes used by Order
    type PlatformArticleAttributes struct {

        
            Currency string  `json:"currency"`
         
    }
    
    // PlatformItem used by Order
    type PlatformItem struct {

        
            LastUpdatedAt string  `json:"last_updated_at"`
            Name string  `json:"name"`
            ID float64  `json:"id"`
            CanReturn bool  `json:"can_return"`
            CanCancel bool  `json:"can_cancel"`
            Image []string  `json:"image"`
            Images []string  `json:"images"`
            Attributes PlatformArticleAttributes  `json:"attributes"`
            L2Category []string  `json:"l2_category"`
            Brand string  `json:"brand"`
            Meta map[string]interface{}  `json:"meta"`
            BrandID float64  `json:"brand_id"`
            Size string  `json:"size"`
            Code string  `json:"code"`
            L3Category float64  `json:"l3_category"`
            BranchURL string  `json:"branch_url"`
            L3CategoryName string  `json:"l3_category_name"`
            Color string  `json:"color"`
            SlugKey string  `json:"slug_key"`
            DepartmentID float64  `json:"department_id"`
            L1Category []string  `json:"l1_category"`
         
    }
    
    // GSTDetailsData used by Order
    type GSTDetailsData struct {

        
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            GstTag string  `json:"gst_tag"`
            HsnCode string  `json:"hsn_code"`
            SgstGstFee string  `json:"sgst_gst_fee"`
            CgstGstFee string  `json:"cgst_gst_fee"`
            IsDefaultHsnCode bool  `json:"is_default_hsn_code"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            GstFee float64  `json:"gst_fee"`
            GstinCode string  `json:"gstin_code"`
            CgstTaxPercentage float64  `json:"cgst_tax_percentage"`
            ValueOfGood float64  `json:"value_of_good"`
            SgstTaxPercentage float64  `json:"sgst_tax_percentage"`
            IgstTaxPercentage float64  `json:"igst_tax_percentage"`
            IgstGstFee string  `json:"igst_gst_fee"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            HsnCodeID string  `json:"hsn_code_id"`
         
    }
    
    // ShipmentListingBrand used by Order
    type ShipmentListingBrand struct {

        
            LogoBase64 string  `json:"logo_base64"`
            Logo string  `json:"logo"`
            CreatedOn string  `json:"created_on"`
            Name string  `json:"name"`
         
    }
    
    // Dates used by Order
    type Dates struct {

        
            DeliveryDate string  `json:"delivery_date"`
            OrderCreated string  `json:"order_created"`
         
    }
    
    // BagReturnableCancelableStatus used by Order
    type BagReturnableCancelableStatus struct {

        
            IsReturnable bool  `json:"is_returnable"`
            EnableTracking bool  `json:"enable_tracking"`
            CanBeCancelled bool  `json:"can_be_cancelled"`
            IsCustomerReturnAllowed bool  `json:"is_customer_return_allowed"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // Dimensions used by Order
    type Dimensions struct {

        
            Width float64  `json:"width"`
            Unit string  `json:"unit"`
            Height float64  `json:"height"`
            Length float64  `json:"length"`
            IsDefault bool  `json:"is_default"`
         
    }
    
    // Weight used by Order
    type Weight struct {

        
            Unit string  `json:"unit"`
            Shipping float64  `json:"shipping"`
            IsDefault bool  `json:"is_default"`
         
    }
    
    // ReturnConfig used by Order
    type ReturnConfig struct {

        
            Unit string  `json:"unit"`
            Returnable bool  `json:"returnable"`
            Time float64  `json:"time"`
         
    }
    
    // Article used by Order
    type Article struct {

        
            RawMeta string  `json:"raw_meta"`
            Size string  `json:"size"`
            Code string  `json:"code"`
            Dimensions Dimensions  `json:"dimensions"`
            ASet map[string]interface{}  `json:"a_set"`
            Weight Weight  `json:"weight"`
            EspModified bool  `json:"esp_modified"`
            Identifiers map[string]interface{}  `json:"identifiers"`
            IsSet bool  `json:"is_set"`
            UID string  `json:"uid"`
            ID string  `json:"_id"`
            Currency map[string]interface{}  `json:"currency"`
            ChildDetails map[string]interface{}  `json:"child_details"`
            ReturnConfig ReturnConfig  `json:"return_config"`
            SellerIdentifier string  `json:"seller_identifier"`
         
    }
    
    // Prices used by Order
    type Prices struct {

        
            FyndCredits float64  `json:"fynd_credits"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            AmountPaidRoundoff float64  `json:"amount_paid_roundoff"`
            PriceEffective float64  `json:"price_effective"`
            CashbackApplied float64  `json:"cashback_applied"`
            ValueOfGood float64  `json:"value_of_good"`
            RefundCredit float64  `json:"refund_credit"`
            TransferPrice float64  `json:"transfer_price"`
            PmPriceSplit float64  `json:"pm_price_split"`
            Discount float64  `json:"discount"`
            GiftPrice float64  `json:"gift_price"`
            DeliveryCharge float64  `json:"delivery_charge"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            Cashback float64  `json:"cashback"`
            AmountPaid float64  `json:"amount_paid"`
            PriceMarked float64  `json:"price_marked"`
            PromotionEffectiveDiscount float64  `json:"promotion_effective_discount"`
            CodCharges float64  `json:"cod_charges"`
            RefundAmount float64  `json:"refund_amount"`
            CouponValue float64  `json:"coupon_value"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
         
    }
    
    // Identifier used by Order
    type Identifier struct {

        
            Ean string  `json:"ean"`
            SkuCode string  `json:"sku_code"`
            Alu string  `json:"alu"`
            Isbn string  `json:"isbn"`
            Upc string  `json:"upc"`
         
    }
    
    // FinancialBreakup used by Order
    type FinancialBreakup struct {

        
            FyndCredits float64  `json:"fynd_credits"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            AmountPaidRoundoff float64  `json:"amount_paid_roundoff"`
            PriceEffective float64  `json:"price_effective"`
            Identifiers Identifier  `json:"identifiers"`
            AddedToFyndCash bool  `json:"added_to_fynd_cash"`
            ItemName string  `json:"item_name"`
            CashbackApplied float64  `json:"cashback_applied"`
            ValueOfGood float64  `json:"value_of_good"`
            TotalUnits float64  `json:"total_units"`
            RefundCredit float64  `json:"refund_credit"`
            TransferPrice float64  `json:"transfer_price"`
            Size string  `json:"size"`
            HsnCode string  `json:"hsn_code"`
            Discount float64  `json:"discount"`
            DeliveryCharge float64  `json:"delivery_charge"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            GstFee float64  `json:"gst_fee"`
            Cashback float64  `json:"cashback"`
            AmountPaid float64  `json:"amount_paid"`
            PriceMarked float64  `json:"price_marked"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            PromotionEffectiveDiscount float64  `json:"promotion_effective_discount"`
            GstTag string  `json:"gst_tag"`
            CodCharges float64  `json:"cod_charges"`
            CouponValue float64  `json:"coupon_value"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
         
    }
    
    // BagUnit used by Order
    type BagUnit struct {

        
            CurrentStatus BagStatusHistory  `json:"current_status"`
            AffiliateBagDetails AffiliateBagDetails  `json:"affiliate_bag_details"`
            EntityType string  `json:"entity_type"`
            CanReturn bool  `json:"can_return"`
            CurrentOperationalStatus BagStatusHistory  `json:"current_operational_status"`
            CanCancel bool  `json:"can_cancel"`
            Item PlatformItem  `json:"item"`
            LineNumber float64  `json:"line_number"`
            Gst GSTDetailsData  `json:"gst"`
            Brand ShipmentListingBrand  `json:"brand"`
            Meta map[string]interface{}  `json:"meta"`
            Size string  `json:"size"`
            BagStatus []BagStatusHistory  `json:"bag_status"`
            Dates Dates  `json:"dates"`
            Status BagReturnableCancelableStatus  `json:"status"`
            Article Article  `json:"article"`
            BagID float64  `json:"bag_id"`
            ProductQuantity float64  `json:"product_quantity"`
            BagExpiryDate string  `json:"bag_expiry_date"`
            Prices Prices  `json:"prices"`
            Reasons []map[string]interface{}  `json:"reasons"`
            DisplayName string  `json:"display_name"`
            FinancialBreakup []FinancialBreakup  `json:"financial_breakup"`
            BagType string  `json:"bag_type"`
         
    }
    
    // ShipmentItem used by Order
    type ShipmentItem struct {

        
            Channel ShipmentListingChannel  `json:"channel"`
            ShipmentCreatedAt string  `json:"shipment_created_at"`
            PaymentMethods map[string]interface{}  `json:"payment_methods"`
            CanProcess bool  `json:"can_process"`
            FulfillingStore ShipmentItemFulFillingStore  `json:"fulfilling_store"`
            OrderingChannnel string  `json:"ordering_channnel"`
            DeliveryAddress PlatformDeliveryAddress  `json:"delivery_address"`
            StatusCreatedAt string  `json:"status_created_at"`
            TotalBags float64  `json:"total_bags"`
            OrderID string  `json:"order_id"`
            User UserDataInfo  `json:"user"`
            InvoiceID string  `json:"invoice_id"`
            CustomerNote string  `json:"customer_note"`
            ShipmentStatus ShipmentStatus  `json:"shipment_status"`
            PaymentMode string  `json:"payment_mode"`
            Meta ShipmentItemMeta  `json:"meta"`
            PreviousShipmentID string  `json:"previous_shipment_id"`
            OrderDate string  `json:"order_date"`
            Bags []BagUnit  `json:"bags"`
            Prices Prices  `json:"prices"`
            ShipmentID string  `json:"shipment_id"`
            EstimatedSlaTime string  `json:"estimated_sla_time"`
            DisplayName string  `json:"display_name"`
            LockStatus bool  `json:"lock_status"`
         
    }
    
    // ShipmentInternalPlatformViewResponse used by Order
    type ShipmentInternalPlatformViewResponse struct {

        
            TotalCount float64  `json:"total_count"`
            Success bool  `json:"success"`
            Page Page  `json:"page"`
            Items []ShipmentItem  `json:"items"`
            Lane string  `json:"lane"`
            Message string  `json:"message"`
         
    }
    
    // Error used by Order
    type Error struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // ShipmentPayments used by Order
    type ShipmentPayments struct {

        
            Logo string  `json:"logo"`
            Source string  `json:"source"`
            Mode string  `json:"mode"`
         
    }
    
    // FulfillingStore used by Order
    type FulfillingStore struct {

        
            Meta map[string]interface{}  `json:"meta"`
            Code string  `json:"code"`
            ContactPerson string  `json:"contact_person"`
            Address string  `json:"address"`
            Pincode string  `json:"pincode"`
            ID float64  `json:"id"`
            Country string  `json:"country"`
            City string  `json:"city"`
            Phone string  `json:"phone"`
            StoreName string  `json:"store_name"`
            FulfillmentChannel string  `json:"fulfillment_channel"`
            State string  `json:"state"`
         
    }
    
    // PDFLinks used by Order
    type PDFLinks struct {

        
            InvoicePos string  `json:"invoice_pos"`
            CreditNoteURL string  `json:"credit_note_url"`
            LabelType string  `json:"label_type"`
            Invoice string  `json:"invoice"`
            InvoiceExport string  `json:"invoice_export"`
            InvoiceType string  `json:"invoice_type"`
            InvoiceA4 string  `json:"invoice_a4"`
            Label string  `json:"label"`
            LabelA6 string  `json:"label_a6"`
            DeliveryChallanA4 string  `json:"delivery_challan_a4"`
            B2b string  `json:"b2b"`
            LabelExport string  `json:"label_export"`
            LabelPos string  `json:"label_pos"`
            LabelA4 string  `json:"label_a4"`
            InvoiceA6 string  `json:"invoice_a6"`
            PoInvoice string  `json:"po_invoice"`
         
    }
    
    // BuyerDetails used by Order
    type BuyerDetails struct {

        
            AjioSiteID string  `json:"ajio_site_id"`
            Address string  `json:"address"`
            Pincode float64  `json:"pincode"`
            Gstin string  `json:"gstin"`
            Name string  `json:"name"`
            City string  `json:"city"`
            State string  `json:"state"`
         
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

        
            B2cBuyerDetails map[string]interface{}  `json:"b2c_buyer_details"`
            DpOptions map[string]interface{}  `json:"dp_options"`
            OrderType string  `json:"order_type"`
            BagWeight map[string]interface{}  `json:"bag_weight"`
            PoNumber string  `json:"po_number"`
            EwaybillInfo map[string]interface{}  `json:"ewaybill_info"`
            LockData LockData  `json:"lock_data"`
            B2bBuyerDetails BuyerDetails  `json:"b2b_buyer_details"`
            ReturnAwbNumber string  `json:"return_awb_number"`
            Dimension Dimensions  `json:"dimension"`
            ShipmentVolumetricWeight float64  `json:"shipment_volumetric_weight"`
            AutoTriggerDpAssignmentAcf bool  `json:"auto_trigger_dp_assignment_acf"`
            External map[string]interface{}  `json:"external"`
            Formatted Formatted  `json:"formatted"`
            PackagingName string  `json:"packaging_name"`
            ParentDpID string  `json:"parent_dp_id"`
            DebugInfo DebugInfo  `json:"debug_info"`
            ForwardAffiliateOrderID string  `json:"forward_affiliate_order_id"`
            EinvoiceInfo EinvoiceInfo  `json:"einvoice_info"`
            AwbNumber string  `json:"awb_number"`
            ReturnStoreNode float64  `json:"return_store_node"`
            ReturnAffiliateOrderID string  `json:"return_affiliate_order_id"`
            MarketplaceStoreID string  `json:"marketplace_store_id"`
            ReturnDetails map[string]interface{}  `json:"return_details"`
            Timestamp ShipmentTimeStamp  `json:"timestamp"`
            DpSortKey string  `json:"dp_sort_key"`
            Weight float64  `json:"weight"`
            DpName string  `json:"dp_name"`
            AssignDpFromSb bool  `json:"assign_dp_from_sb"`
            ShipmentWeight float64  `json:"shipment_weight"`
            ReturnAffiliateShipmentID string  `json:"return_affiliate_shipment_id"`
            SameStoreAvailable bool  `json:"same_store_available"`
            StoreInvoiceUpdatedDate string  `json:"store_invoice_updated_date"`
            BoxType string  `json:"box_type"`
            ForwardAffiliateShipmentID string  `json:"forward_affiliate_shipment_id"`
            DueDate string  `json:"due_date"`
            ShipmentTags []ShipmentTags  `json:"shipment_tags"`
            DpID string  `json:"dp_id"`
            FulfilmentPriorityText string  `json:"fulfilment_priority_text"`
         
    }
    
    // AffiliateDetails used by Order
    type AffiliateDetails struct {

        
            AffiliateOrderID string  `json:"affiliate_order_id"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            AffiliateMeta AffiliateMeta  `json:"affiliate_meta"`
            PdfLinks PDFLinks  `json:"pdf_links"`
            AdID string  `json:"ad_id"`
            AffiliateStoreID string  `json:"affiliate_store_id"`
            CompanyAffiliateTag string  `json:"company_affiliate_tag"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
            ShipmentMeta ShipmentMeta  `json:"shipment_meta"`
            AffiliateID string  `json:"affiliate_id"`
         
    }
    
    // InvoiceInfo used by Order
    type InvoiceInfo struct {

        
            LabelURL string  `json:"label_url"`
            CreditNoteID string  `json:"credit_note_id"`
            ExternalInvoiceID string  `json:"external_invoice_id"`
            UpdatedDate string  `json:"updated_date"`
            StoreInvoiceID string  `json:"store_invoice_id"`
            InvoiceURL string  `json:"invoice_url"`
         
    }
    
    // UserDetailsData used by Order
    type UserDetailsData struct {

        
            Address1 string  `json:"address1"`
            Address string  `json:"address"`
            Pincode string  `json:"pincode"`
            Email string  `json:"email"`
            Name string  `json:"name"`
            Country string  `json:"country"`
            City string  `json:"city"`
            Phone string  `json:"phone"`
            Area string  `json:"area"`
            AddressType string  `json:"address_type"`
            Landmark string  `json:"landmark"`
            State string  `json:"state"`
         
    }
    
    // ShipmentDetails used by Order
    type ShipmentDetails struct {

        
            LockMessage string  `json:"lock_message"`
            LockStatus bool  `json:"lock_status"`
            ActionToStatus map[string]interface{}  `json:"action_to_status"`
         
    }
    
    // ShipmentStatusData used by Order
    type ShipmentStatusData struct {

        
            CreatedAt string  `json:"created_at"`
            Meta map[string]interface{}  `json:"meta"`
            Status string  `json:"status"`
            ShipmentID string  `json:"shipment_id"`
            BagList []string  `json:"bag_list"`
            ID float64  `json:"id"`
            DisplayName string  `json:"display_name"`
         
    }
    
    // OrderingStoreDetails used by Order
    type OrderingStoreDetails struct {

        
            Meta map[string]interface{}  `json:"meta"`
            Code string  `json:"code"`
            ContactPerson string  `json:"contact_person"`
            Address string  `json:"address"`
            Pincode string  `json:"pincode"`
            ID float64  `json:"id"`
            Country string  `json:"country"`
            City string  `json:"city"`
            Phone string  `json:"phone"`
            StoreName string  `json:"store_name"`
            State string  `json:"state"`
         
    }
    
    // CurrentStatus used by Order
    type CurrentStatus struct {

        
            CreatedAt string  `json:"created_at"`
            DeliveryPartnerID float64  `json:"delivery_partner_id"`
            BagStateMapper BagStateMapper  `json:"bag_state_mapper"`
            StateID float64  `json:"state_id"`
            StoreID float64  `json:"store_id"`
            Status string  `json:"status"`
            BagID float64  `json:"bag_id"`
            ID float64  `json:"id"`
            ShipmentID string  `json:"shipment_id"`
            DeliveryAwbNumber string  `json:"delivery_awb_number"`
            UpdatedAt float64  `json:"updated_at"`
            KafkaSync bool  `json:"kafka_sync"`
            StateType string  `json:"state_type"`
         
    }
    
    // BagGST used by Order
    type BagGST struct {

        
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            GstTag string  `json:"gst_tag"`
            HsnCode string  `json:"hsn_code"`
            IgstGstFee string  `json:"igst_gst_fee"`
            SgstGstFee string  `json:"sgst_gst_fee"`
            CgstGstFee string  `json:"cgst_gst_fee"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            GstFee float64  `json:"gst_fee"`
            GstinCode string  `json:"gstin_code"`
            CgstTaxPercentage float64  `json:"cgst_tax_percentage"`
            ValueOfGood float64  `json:"value_of_good"`
            SgstTaxPercentage float64  `json:"sgst_tax_percentage"`
            IgstTaxPercentage float64  `json:"igst_tax_percentage"`
            IsDefaultHsnCode bool  `json:"is_default_hsn_code"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            HsnCodeID string  `json:"hsn_code_id"`
         
    }
    
    // DiscountRules used by Order
    type DiscountRules struct {

        
            Value float64  `json:"value"`
            Type string  `json:"type"`
         
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
            PromotionType string  `json:"promotion_type"`
            ArticleQuantity float64  `json:"article_quantity"`
            PromoID string  `json:"promo_id"`
            PromotionName string  `json:"promotion_name"`
            MrpPromotion bool  `json:"mrp_promotion"`
            BuyRules []BuyRules  `json:"buy_rules"`
            Amount float64  `json:"amount"`
         
    }
    
    // BagPaymentMethods used by Order
    type BagPaymentMethods struct {

        
            Amount float64  `json:"amount"`
            Mode string  `json:"mode"`
         
    }
    
    // AffiliateBagsDetails used by Order
    type AffiliateBagsDetails struct {

        
            CouponCode string  `json:"coupon_code"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
         
    }
    
    // OrderBrandName used by Order
    type OrderBrandName struct {

        
            ModifiedOn string  `json:"modified_on"`
            Logo string  `json:"logo"`
            CreatedOn string  `json:"created_on"`
            Company float64  `json:"company"`
            ID float64  `json:"id"`
            BrandName string  `json:"brand_name"`
         
    }
    
    // B2BPODetails used by Order
    type B2BPODetails struct {

        
            TotalGstPercentage float64  `json:"total_gst_percentage"`
            PartialCanRet bool  `json:"partial_can_ret"`
            PoTaxAmount float64  `json:"po_tax_amount"`
            ItemBasePrice float64  `json:"item_base_price"`
            DockerNumber string  `json:"docker_number"`
            PoLineAmount float64  `json:"po_line_amount"`
         
    }
    
    // GiftCard used by Order
    type GiftCard struct {

        
            GiftMessage string  `json:"gift_message"`
            DisplayText string  `json:"display_text"`
            IsGiftApplied bool  `json:"is_gift_applied"`
            GiftPrice float64  `json:"gift_price"`
         
    }
    
    // BagMeta used by Order
    type BagMeta struct {

        
            B2bPoDetails B2BPODetails  `json:"b2b_po_details"`
            CustomJson map[string]interface{}  `json:"custom_json"`
            DocketNumber string  `json:"docket_number"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            PartialCanRet bool  `json:"partial_can_ret"`
            GiftCard GiftCard  `json:"gift_card"`
            GroupID string  `json:"group_id"`
            CustomMessage string  `json:"custom_message"`
         
    }
    
    // OrderBagArticle used by Order
    type OrderBagArticle struct {

        
            UID string  `json:"uid"`
            Size string  `json:"size"`
            ReturnConfig ReturnConfig  `json:"return_config"`
            Identifiers map[string]interface{}  `json:"identifiers"`
         
    }
    
    // BagConfigs used by Order
    type BagConfigs struct {

        
            IsActive bool  `json:"is_active"`
            IsReturnable bool  `json:"is_returnable"`
            CanBeCancelled bool  `json:"can_be_cancelled"`
            AllowForceReturn bool  `json:"allow_force_return"`
            IsCustomerReturnAllowed bool  `json:"is_customer_return_allowed"`
            EnableTracking bool  `json:"enable_tracking"`
         
    }
    
    // OrderBags used by Order
    type OrderBags struct {

        
            CurrentStatus CurrentStatus  `json:"current_status"`
            GstDetails BagGST  `json:"gst_details"`
            AppliedPromos []AppliedPromos  `json:"applied_promos"`
            PaymentMethods []BagPaymentMethods  `json:"payment_methods"`
            AffiliateBagDetails AffiliateBagsDetails  `json:"affiliate_bag_details"`
            EntityType string  `json:"entity_type"`
            CanReturn bool  `json:"can_return"`
            DeliveryAddress PlatformDeliveryAddress  `json:"delivery_address"`
            CanCancel bool  `json:"can_cancel"`
            Item PlatformItem  `json:"item"`
            LineNumber float64  `json:"line_number"`
            SellerIdentifier string  `json:"seller_identifier"`
            Brand OrderBrandName  `json:"brand"`
            Meta BagMeta  `json:"meta"`
            Quantity float64  `json:"quantity"`
            Article OrderBagArticle  `json:"article"`
            BagID float64  `json:"bag_id"`
            ParentPromoBags map[string]interface{}  `json:"parent_promo_bags"`
            Prices Prices  `json:"prices"`
            BagConfigs BagConfigs  `json:"bag_configs"`
            DisplayName string  `json:"display_name"`
            GroupID string  `json:"group_id"`
            FinancialBreakup FinancialBreakup  `json:"financial_breakup"`
            Identifier string  `json:"identifier"`
         
    }
    
    // PhoneDetails used by Order
    type PhoneDetails struct {

        
            CountryCode float64  `json:"country_code"`
            Number string  `json:"number"`
         
    }
    
    // ContactDetails used by Order
    type ContactDetails struct {

        
            Phone []PhoneDetails  `json:"phone"`
            Emails []string  `json:"emails"`
         
    }
    
    // CompanyDetails used by Order
    type CompanyDetails struct {

        
            CompanyCin string  `json:"company_cin"`
            Address map[string]interface{}  `json:"address"`
            CompanyContact ContactDetails  `json:"company_contact"`
            CompanyGst string  `json:"company_gst"`
            CompanyName string  `json:"company_name"`
            CompanyID float64  `json:"company_id"`
         
    }
    
    // OrderDetailsData used by Order
    type OrderDetailsData struct {

        
            OrderDate string  `json:"order_date"`
            FyndOrderID string  `json:"fynd_order_id"`
            OrderValue string  `json:"order_value"`
            OrderingChannelLogo map[string]interface{}  `json:"ordering_channel_logo"`
            CodCharges string  `json:"cod_charges"`
            Source string  `json:"source"`
            TaxDetails map[string]interface{}  `json:"tax_details"`
            OrderingChannel string  `json:"ordering_channel"`
            AffiliateID string  `json:"affiliate_id"`
         
    }
    
    // TrackingList used by Order
    type TrackingList struct {

        
            IsPassed bool  `json:"is_passed"`
            IsCurrent bool  `json:"is_current"`
            Status string  `json:"status"`
            Time string  `json:"time"`
            Text string  `json:"text"`
         
    }
    
    // DPDetailsData used by Order
    type DPDetailsData struct {

        
            GstTag string  `json:"gst_tag"`
            EwayBillID string  `json:"eway_bill_id"`
            Pincode string  `json:"pincode"`
            AwbNo string  `json:"awb_no"`
            Name string  `json:"name"`
            ID float64  `json:"id"`
            Country string  `json:"country"`
            TrackURL string  `json:"track_url"`
         
    }
    
    // PlatformShipment used by Order
    type PlatformShipment struct {

        
            DpAssignment bool  `json:"dp_assignment"`
            GstDetails GSTDetailsData  `json:"gst_details"`
            ShipmentCreatedAt string  `json:"shipment_created_at"`
            PaymentMethods map[string]interface{}  `json:"payment_methods"`
            BagStatusHistory []BagStatusHistory  `json:"bag_status_history"`
            OperationalStatus string  `json:"operational_status"`
            Payments ShipmentPayments  `json:"payments"`
            FulfillingStore FulfillingStore  `json:"fulfilling_store"`
            ForwardShipmentID string  `json:"forward_shipment_id"`
            Coupon map[string]interface{}  `json:"coupon"`
            TotalBags float64  `json:"total_bags"`
            IsDpAssignEnabled bool  `json:"is_dp_assign_enabled"`
            CustomMeta []map[string]interface{}  `json:"custom_meta"`
            AffiliateDetails AffiliateDetails  `json:"affiliate_details"`
            PriorityText string  `json:"priority_text"`
            Invoice InvoiceInfo  `json:"invoice"`
            User UserDataInfo  `json:"user"`
            InvoiceID string  `json:"invoice_id"`
            UserAgent string  `json:"user_agent"`
            TotalItems float64  `json:"total_items"`
            DeliveryDetails UserDetailsData  `json:"delivery_details"`
            CustomMessage string  `json:"custom_message"`
            ShipmentDetails ShipmentDetails  `json:"shipment_details"`
            ShipmentStatus string  `json:"shipment_status"`
            PaymentMode string  `json:"payment_mode"`
            PlatformLogo string  `json:"platform_logo"`
            Meta ShipmentMeta  `json:"meta"`
            FulfilmentPriority float64  `json:"fulfilment_priority"`
            JourneyType string  `json:"journey_type"`
            BillingDetails UserDetailsData  `json:"billing_details"`
            Status ShipmentStatusData  `json:"status"`
            OrderingStore OrderingStoreDetails  `json:"ordering_store"`
            ShipmentImages []string  `json:"shipment_images"`
            Bags []OrderBags  `json:"bags"`
            PackagingType string  `json:"packaging_type"`
            EnableDpTracking bool  `json:"enable_dp_tracking"`
            CompanyDetails CompanyDetails  `json:"company_details"`
            Prices Prices  `json:"prices"`
            Vertical string  `json:"vertical"`
            ShipmentID string  `json:"shipment_id"`
            EstimatedSlaTime string  `json:"estimated_sla_time"`
            ShipmentQuantity float64  `json:"shipment_quantity"`
            Order OrderDetailsData  `json:"order"`
            TrackingList []TrackingList  `json:"tracking_list"`
            CanUpdateDimension bool  `json:"can_update_dimension"`
            DpDetails DPDetailsData  `json:"dp_details"`
            PdfLinks map[string]interface{}  `json:"pdf_links"`
            DeliverySlot map[string]interface{}  `json:"delivery_slot"`
            LockStatus bool  `json:"lock_status"`
            PickedDate string  `json:"picked_date"`
         
    }
    
    // ShipmentInfoResponse used by Order
    type ShipmentInfoResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
            Shipments []PlatformShipment  `json:"shipments"`
         
    }
    
    // BillingStaffDetails used by Order
    type BillingStaffDetails struct {

        
            StaffID float64  `json:"staff_id"`
            LastName string  `json:"last_name"`
            User string  `json:"user"`
            FirstName string  `json:"first_name"`
            EmployeeCode string  `json:"employee_code"`
         
    }
    
    // TransactionData used by Order
    type TransactionData struct {

        
            PaymentID string  `json:"payment_id"`
            UniqueReferenceNumber string  `json:"unique_reference_number"`
            TransactionID string  `json:"transaction_id"`
            Entity string  `json:"entity"`
            Status string  `json:"status"`
            TerminalID string  `json:"terminal_id"`
            Currency string  `json:"currency"`
            AmountPaid float64  `json:"amount_paid"`
         
    }
    
    // PlatformUserDetails used by Order
    type PlatformUserDetails struct {

        
            PlatformUserFirstName string  `json:"platform_user_first_name"`
            PlatformUserEmployeeCode string  `json:"platform_user_employee_code"`
            PlatformUserLastName string  `json:"platform_user_last_name"`
            PlatformUserID string  `json:"platform_user_id"`
         
    }
    
    // OrderMeta used by Order
    type OrderMeta struct {

        
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            OrderType string  `json:"order_type"`
            OrderPlatform string  `json:"order_platform"`
            Comment string  `json:"comment"`
            PaymentType string  `json:"payment_type"`
            MongoCartID float64  `json:"mongo_cart_id"`
            Files []map[string]interface{}  `json:"files"`
            Staff map[string]interface{}  `json:"staff"`
            CustomerNote string  `json:"customer_note"`
            BillingStaffDetails BillingStaffDetails  `json:"billing_staff_details"`
            TransactionData TransactionData  `json:"transaction_data"`
            PlatformUserDetails PlatformUserDetails  `json:"platform_user_details"`
            CurrencySymbol string  `json:"currency_symbol"`
            OrderingStore float64  `json:"ordering_store"`
            OrderTags []map[string]interface{}  `json:"order_tags"`
            CartID float64  `json:"cart_id"`
            CompanyLogo string  `json:"company_logo"`
            EmployeeID float64  `json:"employee_id"`
            OrderChildEntities []string  `json:"order_child_entities"`
         
    }
    
    // TaxDetails used by Order
    type TaxDetails struct {

        
            Gstin string  `json:"gstin"`
            PanNo string  `json:"pan_no"`
         
    }
    
    // OrderData used by Order
    type OrderData struct {

        
            Meta OrderMeta  `json:"meta"`
            OrderDate string  `json:"order_date"`
            FyndOrderID string  `json:"fynd_order_id"`
            Prices Prices  `json:"prices"`
            PaymentMethods map[string]interface{}  `json:"payment_methods"`
            TaxDetails TaxDetails  `json:"tax_details"`
         
    }
    
    // OrderDetailsResponse used by Order
    type OrderDetailsResponse struct {

        
            Success bool  `json:"success"`
            Shipments []PlatformShipment  `json:"shipments"`
            Order OrderData  `json:"order"`
         
    }
    
    // SubLane used by Order
    type SubLane struct {

        
            Index float64  `json:"index"`
            Actions []map[string]interface{}  `json:"actions"`
            TotalItems float64  `json:"total_items"`
            Text string  `json:"text"`
            Value string  `json:"value"`
         
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

        
            Logo string  `json:"logo"`
            Name string  `json:"name"`
         
    }
    
    // PlatformBreakupValues used by Order
    type PlatformBreakupValues struct {

        
            Display string  `json:"display"`
            Value string  `json:"value"`
            Name string  `json:"name"`
         
    }
    
    // PlatformOrderItems used by Order
    type PlatformOrderItems struct {

        
            Channel PlatformChannel  `json:"channel"`
            Meta map[string]interface{}  `json:"meta"`
            OrderValue float64  `json:"order_value"`
            BreakupValues []PlatformBreakupValues  `json:"breakup_values"`
            UserInfo UserDataInfo  `json:"user_info"`
            OrderID string  `json:"order_id"`
            OrderCreatedTime string  `json:"order_created_time"`
            Shipments []PlatformShipment  `json:"shipments"`
            TotalOrderValue float64  `json:"total_order_value"`
            PaymentMode string  `json:"payment_mode"`
         
    }
    
    // OrderListingResponse used by Order
    type OrderListingResponse struct {

        
            TotalCount float64  `json:"total_count"`
            Success bool  `json:"success"`
            Page Page  `json:"page"`
            Items []PlatformOrderItems  `json:"items"`
            Lane string  `json:"lane"`
            Message string  `json:"message"`
         
    }
    
    // PlatformTrack used by Order
    type PlatformTrack struct {

        
            Meta map[string]interface{}  `json:"meta"`
            Status string  `json:"status"`
            ShipmentType string  `json:"shipment_type"`
            UpdatedAt string  `json:"updated_at"`
            UpdatedTime string  `json:"updated_time"`
            Awb string  `json:"awb"`
            AccountName string  `json:"account_name"`
            RawStatus string  `json:"raw_status"`
            Reason string  `json:"reason"`
            LastLocationRecievedAt string  `json:"last_location_recieved_at"`
         
    }
    
    // PlatformShipmentTrack used by Order
    type PlatformShipmentTrack struct {

        
            Meta map[string]interface{}  `json:"meta"`
            Results []PlatformTrack  `json:"results"`
         
    }
    
    // FilterInfoOption used by Order
    type FilterInfoOption struct {

        
            PlaceholderText string  `json:"placeholder_text"`
            MinSearchSize float64  `json:"min_search_size"`
            Name string  `json:"name"`
            ShowUI bool  `json:"show_ui"`
            Text string  `json:"text"`
            Value string  `json:"value"`
         
    }
    
    // FiltersInfo used by Order
    type FiltersInfo struct {

        
            PlaceholderText string  `json:"placeholder_text"`
            Type string  `json:"type"`
            Options []FilterInfoOption  `json:"options"`
            Required bool  `json:"required"`
            Text string  `json:"text"`
            Value string  `json:"value"`
         
    }
    
    // AdvanceFilterInfo used by Order
    type AdvanceFilterInfo struct {

        
            ActionCentre []FiltersInfo  `json:"action_centre"`
            Returned []FiltersInfo  `json:"returned"`
            AppliedFilters map[string]interface{}  `json:"applied_filters"`
            Page map[string]interface{}  `json:"page"`
            Unfulfilled []FiltersInfo  `json:"unfulfilled"`
            Filters []FiltersInfo  `json:"filters"`
            Processed []FiltersInfo  `json:"processed"`
         
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

        
            DisplayName string  `json:"display_name"`
            ID float64  `json:"id"`
            QcType []string  `json:"qc_type"`
            QuestionSet []QuestionSet  `json:"question_set"`
         
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
    
    // BagGSTDetails used by Order
    type BagGSTDetails struct {

        
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            GstTag string  `json:"gst_tag"`
            HsnCode string  `json:"hsn_code"`
            IgstGstFee string  `json:"igst_gst_fee"`
            SgstGstFee string  `json:"sgst_gst_fee"`
            CgstGstFee string  `json:"cgst_gst_fee"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            GstFee float64  `json:"gst_fee"`
            GstinCode string  `json:"gstin_code"`
            CgstTaxPercentage float64  `json:"cgst_tax_percentage"`
            ValueOfGood float64  `json:"value_of_good"`
            SgstTaxPercentage float64  `json:"sgst_tax_percentage"`
            IgstTaxPercentage float64  `json:"igst_tax_percentage"`
            IsDefaultHsnCode bool  `json:"is_default_hsn_code"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            HsnCodeID string  `json:"hsn_code_id"`
         
    }
    
    // Attributes used by Order
    type Attributes struct {

        
            PrimaryColor string  `json:"primary_color"`
            PrimaryMaterial string  `json:"primary_material"`
            PrimaryColorHex string  `json:"primary_color_hex"`
            Name string  `json:"name"`
            MarketerName string  `json:"marketer_name"`
            Essential string  `json:"essential"`
            MarketerAddress string  `json:"marketer_address"`
            BrandName string  `json:"brand_name"`
            Gender []string  `json:"gender"`
         
    }
    
    // Item used by Order
    type Item struct {

        
            LastUpdatedAt string  `json:"last_updated_at"`
            Name string  `json:"name"`
            L2CategoryID float64  `json:"l2_category_id"`
            CanReturn bool  `json:"can_return"`
            CanCancel bool  `json:"can_cancel"`
            Image []string  `json:"image"`
            Attributes Attributes  `json:"attributes"`
            L2Category []string  `json:"l2_category"`
            BrandID float64  `json:"brand_id"`
            Meta map[string]interface{}  `json:"meta"`
            Brand string  `json:"brand"`
            Size string  `json:"size"`
            L1CategoryID float64  `json:"l1_category_id"`
            Code string  `json:"code"`
            L3Category float64  `json:"l3_category"`
            BranchURL string  `json:"branch_url"`
            Gender string  `json:"gender"`
            WebstoreProductURL string  `json:"webstore_product_url"`
            ItemID float64  `json:"item_id"`
            L3CategoryName string  `json:"l3_category_name"`
            Color string  `json:"color"`
            SlugKey string  `json:"slug_key"`
            DepartmentID float64  `json:"department_id"`
            L1Category []string  `json:"l1_category"`
         
    }
    
    // Brand used by Order
    type Brand struct {

        
            BrandID float64  `json:"brand_id"`
            CreditNoteExpiryDays float64  `json:"credit_note_expiry_days"`
            ModifiedOn float64  `json:"modified_on"`
            ScriptLastRan string  `json:"script_last_ran"`
            InvoicePrefix string  `json:"invoice_prefix"`
            StartDate string  `json:"start_date"`
            PickupLocation string  `json:"pickup_location"`
            Logo string  `json:"logo"`
            IsVirtualInvoice bool  `json:"is_virtual_invoice"`
            CreatedOn float64  `json:"created_on"`
            Company string  `json:"company"`
            BrandName string  `json:"brand_name"`
            CreditNoteAllowed bool  `json:"credit_note_allowed"`
         
    }
    
    // Document used by Order
    type Document struct {

        
            Value string  `json:"value"`
            LegalName string  `json:"legal_name"`
            DsType string  `json:"ds_type"`
            URL string  `json:"url"`
            Verified bool  `json:"verified"`
         
    }
    
    // StoreDocuments used by Order
    type StoreDocuments struct {

        
            Gst Document  `json:"gst"`
         
    }
    
    // StoreEinvoice used by Order
    type StoreEinvoice struct {

        
            User string  `json:"user"`
            Password string  `json:"password"`
            Enabled bool  `json:"enabled"`
            Username string  `json:"username"`
         
    }
    
    // StoreEwaybill used by Order
    type StoreEwaybill struct {

        
            Enabled bool  `json:"enabled"`
         
    }
    
    // StoreGstCredentials used by Order
    type StoreGstCredentials struct {

        
            EInvoice StoreEinvoice  `json:"e_invoice"`
            EWaybill StoreEwaybill  `json:"e_waybill"`
         
    }
    
    // EInvoicePortalDetails used by Order
    type EInvoicePortalDetails struct {

        
            Password string  `json:"password"`
            User string  `json:"user"`
            Username string  `json:"username"`
         
    }
    
    // StoreMeta used by Order
    type StoreMeta struct {

        
            GstNumber string  `json:"gst_number"`
            Documents StoreDocuments  `json:"documents"`
            EwaybillPortalDetails map[string]interface{}  `json:"ewaybill_portal_details"`
            GstCredentials StoreGstCredentials  `json:"gst_credentials"`
            NotificationEmails []string  `json:"notification_emails"`
            Stage string  `json:"stage"`
            EinvoicePortalDetails EInvoicePortalDetails  `json:"einvoice_portal_details"`
            DisplayName string  `json:"display_name"`
            Timing []map[string]interface{}  `json:"timing"`
            AdditionalContactDetails map[string]interface{}  `json:"additional_contact_details"`
            ProductReturnConfig map[string]interface{}  `json:"product_return_config"`
         
    }
    
    // StoreAddress used by Order
    type StoreAddress struct {

        
            Address1 string  `json:"address1"`
            Pincode float64  `json:"pincode"`
            Latitude float64  `json:"latitude"`
            CountryCode string  `json:"country_code"`
            Area string  `json:"area"`
            Address2 string  `json:"address2"`
            Landmark string  `json:"landmark"`
            ContactPerson string  `json:"contact_person"`
            UpdatedAt string  `json:"updated_at"`
            Country string  `json:"country"`
            City string  `json:"city"`
            Longitude float64  `json:"longitude"`
            AddressCategory string  `json:"address_category"`
            Email string  `json:"email"`
            State string  `json:"state"`
            CreatedAt string  `json:"created_at"`
            Version string  `json:"version"`
            Phone string  `json:"phone"`
            AddressType string  `json:"address_type"`
         
    }
    
    // Store used by Order
    type Store struct {

        
            Address1 string  `json:"address1"`
            MallArea string  `json:"mall_area"`
            Pincode string  `json:"pincode"`
            OrderIntegrationID string  `json:"order_integration_id"`
            Latitude float64  `json:"latitude"`
            Name string  `json:"name"`
            ParentStoreID float64  `json:"parent_store_id"`
            StoreActiveFrom string  `json:"store_active_from"`
            IsEnabledForRecon bool  `json:"is_enabled_for_recon"`
            Address2 string  `json:"address2"`
            FulfillmentChannel string  `json:"fulfillment_channel"`
            IsActive bool  `json:"is_active"`
            ContactPerson string  `json:"contact_person"`
            LoginUsername string  `json:"login_username"`
            Country string  `json:"country"`
            UpdatedAt string  `json:"updated_at"`
            City string  `json:"city"`
            Longitude float64  `json:"longitude"`
            BrandID interface{}  `json:"brand_id"`
            Meta StoreMeta  `json:"meta"`
            Code string  `json:"code"`
            PackagingMaterialCount float64  `json:"packaging_material_count"`
            LocationType string  `json:"location_type"`
            SID string  `json:"s_id"`
            CompanyID float64  `json:"company_id"`
            AlohomoraUserID float64  `json:"alohomora_user_id"`
            State string  `json:"state"`
            CreatedAt string  `json:"created_at"`
            StoreAddressJson StoreAddress  `json:"store_address_json"`
            StoreEmail string  `json:"store_email"`
            MallName string  `json:"mall_name"`
            BrandStoreTags []string  `json:"brand_store_tags"`
            Phone float64  `json:"phone"`
            VatNo string  `json:"vat_no"`
            IsArchived bool  `json:"is_archived"`
         
    }
    
    // ArticleDetails used by Order
    type ArticleDetails struct {

        
            Status map[string]interface{}  `json:"status"`
         
    }
    
    // BagDetailsPlatformResponse used by Order
    type BagDetailsPlatformResponse struct {

        
            CurrentStatus BagStatusHistory  `json:"current_status"`
            AppliedPromos []map[string]interface{}  `json:"applied_promos"`
            BagUpdateTime float64  `json:"bag_update_time"`
            GstDetails BagGSTDetails  `json:"gst_details"`
            AffiliateBagDetails AffiliateBagDetails  `json:"affiliate_bag_details"`
            EntityType string  `json:"entity_type"`
            OrderIntegrationID string  `json:"order_integration_id"`
            ID float64  `json:"id"`
            BagStatusHistory BagStatusHistory  `json:"bag_status_history"`
            OperationalStatus string  `json:"operational_status"`
            QcRequired interface{}  `json:"qc_required"`
            CurrentOperationalStatus BagStatusHistory  `json:"current_operational_status"`
            AffiliateDetails AffiliateDetails  `json:"affiliate_details"`
            Item Item  `json:"item"`
            LineNumber float64  `json:"line_number"`
            OriginalBagList []float64  `json:"original_bag_list"`
            RestoreCoupon bool  `json:"restore_coupon"`
            Identifier string  `json:"identifier"`
            RestorePromos map[string]interface{}  `json:"restore_promos"`
            Brand Brand  `json:"brand"`
            Meta BagMeta  `json:"meta"`
            NoOfBagsOrder float64  `json:"no_of_bags_order"`
            SellerIdentifier string  `json:"seller_identifier"`
            JourneyType string  `json:"journey_type"`
            Type string  `json:"type"`
            BagStatus []BagStatusHistory  `json:"bag_status"`
            Dates Dates  `json:"dates"`
            Status BagReturnableCancelableStatus  `json:"status"`
            Quantity float64  `json:"quantity"`
            Article Article  `json:"article"`
            ParentPromoBags map[string]interface{}  `json:"parent_promo_bags"`
            OrderingStore Store  `json:"ordering_store"`
            ArticleDetails ArticleDetails  `json:"article_details"`
            Prices Prices  `json:"prices"`
            ShipmentID string  `json:"shipment_id"`
            Reasons []map[string]interface{}  `json:"reasons"`
            DisplayName string  `json:"display_name"`
            Tags []string  `json:"tags"`
            FinancialBreakup []FinancialBreakup  `json:"financial_breakup"`
         
    }
    
    // ErrorResponse used by Order
    type ErrorResponse struct {

        
            Message string  `json:"message"`
            Error string  `json:"error"`
         
    }
    
    // BagsPage used by Order
    type BagsPage struct {

        
            PageType string  `json:"page_type"`
            Size float64  `json:"size"`
            Current float64  `json:"current"`
            HasNext bool  `json:"has_next"`
            ItemTotal float64  `json:"item_total"`
         
    }
    
    // GetBagsPlatformResponse used by Order
    type GetBagsPlatformResponse struct {

        
            Items []BagDetailsPlatformResponse  `json:"items"`
            Page BagsPage  `json:"page"`
         
    }
    
    // GeneratePosOrderReceiptResponse used by Order
    type GeneratePosOrderReceiptResponse struct {

        
            MerchantCnReceipt string  `json:"merchant_cn_receipt"`
            PaymentReceipt string  `json:"payment_receipt"`
            OrderID string  `json:"order_id"`
            CustomerCnReceipt string  `json:"customer_cn_receipt"`
            Success bool  `json:"success"`
            InvoiceReceipt string  `json:"invoice_receipt"`
         
    }
    
    // InvalidateShipmentCachePayload used by Order
    type InvalidateShipmentCachePayload struct {

        
            AffiliateBagIds []string  `json:"affiliate_bag_ids"`
            ShipmentIds []string  `json:"shipment_ids"`
            BagIds []string  `json:"bag_ids"`
         
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
            ErrorTrace string  `json:"error_trace"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // StoreReassign used by Order
    type StoreReassign struct {

        
            BagID float64  `json:"bag_id"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            SetID string  `json:"set_id"`
            AffiliateID string  `json:"affiliate_id"`
            ReasonIds []float64  `json:"reason_ids"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
            MongoArticleID string  `json:"mongo_article_id"`
            ItemID string  `json:"item_id"`
            FyndOrderID string  `json:"fynd_order_id"`
            StoreID float64  `json:"store_id"`
         
    }
    
    // StoreReassignResponse used by Order
    type StoreReassignResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // Entities used by Order
    type Entities struct {

        
            ReasonText string  `json:"reason_text"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            AffiliateID string  `json:"affiliate_id"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            ID string  `json:"id"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
         
    }
    
    // UpdateShipmentLockPayload used by Order
    type UpdateShipmentLockPayload struct {

        
            EntityType string  `json:"entity_type"`
            Action string  `json:"action"`
            Entities []Entities  `json:"entities"`
            ActionType string  `json:"action_type"`
         
    }
    
    // Bags used by Order
    type Bags struct {

        
            BagID float64  `json:"bag_id"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
            IsLocked bool  `json:"is_locked"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
         
    }
    
    // OriginalFilter used by Order
    type OriginalFilter struct {

        
            AffiliateID string  `json:"affiliate_id"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
         
    }
    
    // CheckResponse used by Order
    type CheckResponse struct {

        
            Status string  `json:"status"`
            AffiliateID string  `json:"affiliate_id"`
            Bags []Bags  `json:"bags"`
            IsBagLocked bool  `json:"is_bag_locked"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            IsShipmentLocked bool  `json:"is_shipment_locked"`
            LockStatus string  `json:"lock_status"`
            OriginalFilter OriginalFilter  `json:"original_filter"`
            ShipmentID string  `json:"shipment_id"`
         
    }
    
    // UpdateShipmentLockResponse used by Order
    type UpdateShipmentLockResponse struct {

        
            CheckResponse []CheckResponse  `json:"check_response"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // AnnouncementResponse used by Order
    type AnnouncementResponse struct {

        
            CompanyID float64  `json:"company_id"`
            CreatedAt string  `json:"created_at"`
            PlatformName string  `json:"platform_name"`
            ID float64  `json:"id"`
            LogoURL string  `json:"logo_url"`
            PlatformID string  `json:"platform_id"`
            ToDatetime string  `json:"to_datetime"`
            Title string  `json:"title"`
            Description string  `json:"description"`
            FromDatetime string  `json:"from_datetime"`
         
    }
    
    // AnnouncementsResponse used by Order
    type AnnouncementsResponse struct {

        
            Announcements []AnnouncementResponse  `json:"announcements"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // BaseResponse used by Order
    type BaseResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // Click2CallResponse used by Order
    type Click2CallResponse struct {

        
            Success bool  `json:"success"`
            CallID string  `json:"call_id"`
         
    }
    
    // ErrorDetail used by Order
    type ErrorDetail struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // Products used by Order
    type Products struct {

        
            Quantity float64  `json:"quantity"`
            LineNumber float64  `json:"line_number"`
            Identifier string  `json:"identifier"`
         
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
    
    // OrderItemDataUpdates used by Order
    type OrderItemDataUpdates struct {

        
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
            OrderItemStatus []OrderItemDataUpdates  `json:"order_item_status"`
            Entities []EntitiesDataUpdates  `json:"entities"`
         
    }
    
    // ProductsReasonsFilters used by Order
    type ProductsReasonsFilters struct {

        
            Quantity float64  `json:"quantity"`
            LineNumber float64  `json:"line_number"`
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
            Identifier string  `json:"identifier"`
            DataUpdates DataUpdates  `json:"data_updates"`
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

        
            ForceTransition bool  `json:"force_transition"`
            Statuses []StatuesRequest  `json:"statuses"`
            Task bool  `json:"task"`
            UnlockBeforeTransition bool  `json:"unlock_before_transition"`
            LockAfterTransition bool  `json:"lock_after_transition"`
         
    }
    
    // ShipmentsResponse used by Order
    type ShipmentsResponse struct {

        
            Status float64  `json:"status"`
            StackTrace string  `json:"stack_trace"`
            Identifier string  `json:"identifier"`
            Exception string  `json:"exception"`
            Code string  `json:"code"`
            Meta map[string]interface{}  `json:"meta"`
            FinalState map[string]interface{}  `json:"final_state"`
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
    
    // DateRange used by Order
    type DateRange struct {

        
            FromDate string  `json:"from_date"`
            ToDate string  `json:"to_date"`
         
    }
    
    // FiltersRequest used by Order
    type FiltersRequest struct {

        
            Logo string  `json:"logo"`
            StoreName string  `json:"store_name"`
            DateRange DateRange  `json:"date_range"`
            Lane string  `json:"lane"`
            DpIds float64  `json:"dp_ids"`
            DpName string  `json:"dp_name"`
            Stores float64  `json:"stores"`
         
    }
    
    // ProcessManifest used by Order
    type ProcessManifest struct {

        
            Filters FiltersRequest  `json:"filters"`
            Action string  `json:"action"`
            ManifestID string  `json:"manifest_id"`
            UniqueID string  `json:"unique_id"`
         
    }
    
    // Filters used by Order
    type Filters struct {

        
            Logo string  `json:"logo"`
            StoreName string  `json:"store_name"`
            DateRange DateRange  `json:"date_range"`
            FromDate string  `json:"from_date"`
            SelectedShipments string  `json:"selected_shipments"`
            Lane string  `json:"lane"`
            DpIds float64  `json:"dp_ids"`
            DpName string  `json:"dp_name"`
            ToDate string  `json:"to_date"`
            Stores float64  `json:"stores"`
         
    }
    
    // ProcessManifestResponse used by Order
    type ProcessManifestResponse struct {

        
            CreatedBy string  `json:"created_by"`
            Action string  `json:"action"`
            UserID string  `json:"user_id"`
            CompanyID float64  `json:"company_id"`
            Filters Filters  `json:"filters"`
            UID string  `json:"uid"`
            ManifestID string  `json:"manifest_id"`
         
    }
    
    // ProcessManifestItemResponse used by Order
    type ProcessManifestItemResponse struct {

        
            Items ProcessManifestResponse  `json:"items"`
         
    }
    
    // FilterInfoOption1 used by Order
    type FilterInfoOption1 struct {

        
            PlaceholderText string  `json:"placeholder_text"`
            ShowUI bool  `json:"show_ui"`
            Value string  `json:"value"`
            Text string  `json:"text"`
            Name string  `json:"name"`
            MinSearchSize float64  `json:"min_search_size"`
         
    }
    
    // FiltersInfo1 used by Order
    type FiltersInfo1 struct {

        
            PlaceholderText string  `json:"placeholder_text"`
            Options []FilterInfoOption1  `json:"options"`
            Required bool  `json:"required"`
            Value string  `json:"value"`
            Text string  `json:"text"`
            Type string  `json:"type"`
         
    }
    
    // ManifestFiltersResponse used by Order
    type ManifestFiltersResponse struct {

        
            GlobalFilter []FiltersInfo1  `json:"global_filter"`
            AdvanceFilter []FiltersInfo1  `json:"advance_filter"`
         
    }
    
    // DispatchManifest used by Order
    type DispatchManifest struct {

        
            ManifestID string  `json:"manifest_id"`
         
    }
    
    // SuccessResponse used by Order
    type SuccessResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
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
    
    // HistoryReason used by Order
    type HistoryReason struct {

        
            Quantity float64  `json:"quantity"`
            DislayName string  `json:"dislay_name"`
            Category string  `json:"category"`
            Text string  `json:"text"`
            State string  `json:"state"`
            Code float64  `json:"code"`
         
    }
    
    // HistoryMeta used by Order
    type HistoryMeta struct {

        
            Endtime string  `json:"endtime"`
            Slug string  `json:"slug"`
            Billsec string  `json:"billsec"`
            ActivityComment string  `json:"activity_comment"`
            Callerid string  `json:"callerid"`
            Reason HistoryReason  `json:"reason"`
            CallID string  `json:"call_id"`
            StoreName string  `json:"store_name"`
            Receiver string  `json:"receiver"`
            Recordpath string  `json:"recordpath"`
            Duration string  `json:"duration"`
            Recipient string  `json:"recipient"`
            ShortLink string  `json:"short_link"`
            Status string  `json:"status"`
            Status1 string  `json:"status1"`
            StoreCode string  `json:"store_code"`
            Starttime string  `json:"starttime"`
            Caller string  `json:"caller"`
            ChannelType string  `json:"channel_type"`
            ActivityType string  `json:"activity_type"`
            Status2 string  `json:"status2"`
            Message string  `json:"message"`
            StoreID float64  `json:"store_id"`
         
    }
    
    // HistoryDict used by Order
    type HistoryDict struct {

        
            BagID float64  `json:"bag_id"`
            L1Detail string  `json:"l1_detail"`
            L3Detail string  `json:"l3_detail"`
            TicketID string  `json:"ticket_id"`
            Createdat string  `json:"createdat"`
            TicketURL string  `json:"ticket_url"`
            AssignedAgent string  `json:"assigned_agent"`
            DisplayMessage string  `json:"display_message"`
            L2Detail string  `json:"l2_detail"`
            User string  `json:"user"`
            Meta HistoryMeta  `json:"meta"`
            Message string  `json:"message"`
            Type string  `json:"type"`
         
    }
    
    // ShipmentHistoryResponse used by Order
    type ShipmentHistoryResponse struct {

        
            ActivityHistory []HistoryDict  `json:"activity_history"`
            Success bool  `json:"success"`
         
    }
    
    // SmsDataPayload used by Order
    type SmsDataPayload struct {

        
            OrderID string  `json:"order_id"`
            PhoneNumber float64  `json:"phone_number"`
            CountryCode string  `json:"country_code"`
            BrandName string  `json:"brand_name"`
            PaymentMode string  `json:"payment_mode"`
            AmountPaid float64  `json:"amount_paid"`
            CustomerName string  `json:"customer_name"`
            Message string  `json:"message"`
            ShipmentID float64  `json:"shipment_id"`
         
    }
    
    // SendSmsPayload used by Order
    type SendSmsPayload struct {

        
            BagID float64  `json:"bag_id"`
            Slug string  `json:"slug"`
            Data SmsDataPayload  `json:"data"`
         
    }
    
    // OrderDetails used by Order
    type OrderDetails struct {

        
            FyndOrderID string  `json:"fynd_order_id"`
            CreatedAt string  `json:"created_at"`
         
    }
    
    // Meta used by Order
    type Meta struct {

        
            KafkaEmissionStatus float64  `json:"kafka_emission_status"`
            StateManagerUsed string  `json:"state_manager_used"`
         
    }
    
    // ShipmentDetail used by Order
    type ShipmentDetail struct {

        
            Status string  `json:"status"`
            Remarks string  `json:"remarks"`
            ID float64  `json:"id"`
            BagList []float64  `json:"bag_list"`
            Meta Meta  `json:"meta"`
            ShipmentID string  `json:"shipment_id"`
         
    }
    
    // OrderStatusData used by Order
    type OrderStatusData struct {

        
            Errors []string  `json:"errors"`
            OrderDetails OrderDetails  `json:"order_details"`
            ShipmentDetails []ShipmentDetail  `json:"shipment_details"`
         
    }
    
    // OrderStatusResult used by Order
    type OrderStatusResult struct {

        
            Success string  `json:"success"`
            Result []OrderStatusData  `json:"result"`
         
    }
    
    // Dimension used by Order
    type Dimension struct {

        
            Width float64  `json:"width"`
            PackagingType string  `json:"packaging_type"`
            Length float64  `json:"length"`
            Height string  `json:"height"`
            Weight string  `json:"weight"`
         
    }
    
    // UpdatePackagingDimensionsPayload used by Order
    type UpdatePackagingDimensionsPayload struct {

        
            CurrentStatus string  `json:"current_status"`
            Dimension []Dimension  `json:"dimension"`
            ShipmentID string  `json:"shipment_id"`
         
    }
    
    // UpdatePackagingDimensionsResponse used by Order
    type UpdatePackagingDimensionsResponse struct {

        
            Message string  `json:"message"`
         
    }
    
    // BillingInfo used by Order
    type BillingInfo struct {

        
            AlternateMobileNumber string  `json:"alternate_mobile_number"`
            MiddleName string  `json:"middle_name"`
            Gender string  `json:"gender"`
            StateCode string  `json:"state_code"`
            Country string  `json:"country"`
            State string  `json:"state"`
            PrimaryEmail string  `json:"primary_email"`
            AlternateEmail string  `json:"alternate_email"`
            Pincode string  `json:"pincode"`
            Address1 string  `json:"address1"`
            CountryCode string  `json:"country_code"`
            HouseNo string  `json:"house_no"`
            CustomerCode string  `json:"customer_code"`
            PrimaryMobileNumber string  `json:"primary_mobile_number"`
            City string  `json:"city"`
            FirstName string  `json:"first_name"`
            ExternalCustomerCode string  `json:"external_customer_code"`
            FloorNo string  `json:"floor_no"`
            Address2 string  `json:"address2"`
            Title string  `json:"title"`
            LastName string  `json:"last_name"`
         
    }
    
    // ProcessingDates used by Order
    type ProcessingDates struct {

        
            DispatchAfterDate string  `json:"dispatch_after_date"`
            DpPickupSlot map[string]interface{}  `json:"dp_pickup_slot"`
            ConfirmByDate string  `json:"confirm_by_date"`
            PackByDate string  `json:"pack_by_date"`
            DispatchByDate string  `json:"dispatch_by_date"`
            CustomerPickupSlot map[string]interface{}  `json:"customer_pickup_slot"`
         
    }
    
    // Tax used by Order
    type Tax struct {

        
            Amount map[string]interface{}  `json:"amount"`
            Breakup []map[string]interface{}  `json:"breakup"`
            Rate float64  `json:"rate"`
            Name string  `json:"name"`
         
    }
    
    // Charge used by Order
    type Charge struct {

        
            Tax Tax  `json:"tax"`
            Amount map[string]interface{}  `json:"amount"`
            Name string  `json:"name"`
            Code string  `json:"code"`
            Type string  `json:"type"`
         
    }
    
    // LineItem used by Order
    type LineItem struct {

        
            Quantity float64  `json:"quantity"`
            ExternalLineID string  `json:"external_line_id"`
            SellerIdentifier string  `json:"seller_identifier"`
            CustomMessasge string  `json:"custom_messasge"`
            Meta map[string]interface{}  `json:"meta"`
            Charges []Charge  `json:"charges"`
         
    }
    
    // Shipment used by Order
    type Shipment struct {

        
            ProcessingDates ProcessingDates  `json:"processing_dates"`
            LocationID float64  `json:"location_id"`
            ExternalShipmentID string  `json:"external_shipment_id"`
            Priority float64  `json:"priority"`
            Meta map[string]interface{}  `json:"meta"`
            LineItems []LineItem  `json:"line_items"`
         
    }
    
    // ShippingInfo used by Order
    type ShippingInfo struct {

        
            ShippingType string  `json:"shipping_type"`
            GeoLocation map[string]interface{}  `json:"geo_location"`
            AddressType string  `json:"address_type"`
            AlternateMobileNumber string  `json:"alternate_mobile_number"`
            MiddleName string  `json:"middle_name"`
            Gender string  `json:"gender"`
            StateCode string  `json:"state_code"`
            Country string  `json:"country"`
            State string  `json:"state"`
            PrimaryEmail string  `json:"primary_email"`
            AlternateEmail string  `json:"alternate_email"`
            Pincode string  `json:"pincode"`
            Address1 string  `json:"address1"`
            Landmark string  `json:"landmark"`
            CountryCode string  `json:"country_code"`
            HouseNo string  `json:"house_no"`
            CustomerCode string  `json:"customer_code"`
            Slot []map[string]interface{}  `json:"slot"`
            PrimaryMobileNumber string  `json:"primary_mobile_number"`
            City string  `json:"city"`
            ExternalCustomerCode string  `json:"external_customer_code"`
            FirstName string  `json:"first_name"`
            FloorNo string  `json:"floor_no"`
            Address2 string  `json:"address2"`
            Title string  `json:"title"`
            LastName string  `json:"last_name"`
         
    }
    
    // TaxInfo used by Order
    type TaxInfo struct {

        
            B2bGstinNumber string  `json:"b2b_gstin_number"`
            Gstin string  `json:"gstin"`
         
    }
    
    // PaymentMethod used by Order
    type PaymentMethod struct {

        
            Amount float64  `json:"amount"`
            Mode string  `json:"mode"`
            TransactionData map[string]interface{}  `json:"transaction_data"`
            Name string  `json:"name"`
            RefundBy string  `json:"refund_by"`
            Meta map[string]interface{}  `json:"meta"`
            CollectBy string  `json:"collect_by"`
         
    }
    
    // PaymentInfo used by Order
    type PaymentInfo struct {

        
            PaymentMethods []PaymentMethod  `json:"payment_methods"`
            PrimaryMode string  `json:"primary_mode"`
         
    }
    
    // CreateOrderAPI used by Order
    type CreateOrderAPI struct {

        
            CurrencyInfo map[string]interface{}  `json:"currency_info"`
            Config map[string]interface{}  `json:"config"`
            BillingInfo BillingInfo  `json:"billing_info"`
            ExternalCreationDate string  `json:"external_creation_date"`
            Shipments []Shipment  `json:"shipments"`
            ShippingInfo ShippingInfo  `json:"shipping_info"`
            TaxInfo TaxInfo  `json:"tax_info"`
            Meta map[string]interface{}  `json:"meta"`
            Charges []Charge  `json:"charges"`
            PaymentInfo PaymentInfo  `json:"payment_info"`
            ExternalOrderID string  `json:"external_order_id"`
         
    }
    
    // CreateOrderResponse used by Order
    type CreateOrderResponse struct {

        
            FyndOrderID string  `json:"fynd_order_id"`
         
    }
    
    // CreateOrderErrorReponse used by Order
    type CreateOrderErrorReponse struct {

        
            Status float64  `json:"status"`
            Info interface{}  `json:"info"`
            StackTrace string  `json:"stack_trace"`
            Exception string  `json:"exception"`
            RequestID string  `json:"request_id"`
            Code string  `json:"code"`
            Meta string  `json:"meta"`
            Message string  `json:"message"`
         
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

        
            Success bool  `json:"success"`
            Message []string  `json:"message"`
         
    }
    
    // FyndOrderIdList used by Order
    type FyndOrderIdList struct {

        
            FyndOrderID []string  `json:"fynd_order_id"`
         
    }
    
    // OrderStatus used by Order
    type OrderStatus struct {

        
            EndDate string  `json:"end_date"`
            StartDate string  `json:"start_date"`
            Mobile float64  `json:"mobile"`
            OrderDetails []FyndOrderIdList  `json:"order_details"`
         
    }
    
    // BagStateTransitionMap used by Order
    type BagStateTransitionMap struct {

        
            Affiliate map[string]interface{}  `json:"affiliate"`
            Fynd map[string]interface{}  `json:"fynd"`
         
    }
    
    // ManifestPageInfo used by Order
    type ManifestPageInfo struct {

        
            Current float64  `json:"current"`
            HasPrevious bool  `json:"has_previous"`
            Total float64  `json:"total"`
            Size float64  `json:"size"`
            HasNext bool  `json:"has_next"`
            Type string  `json:"type"`
         
    }
    
    // ManifestItemDetails used by Order
    type ManifestItemDetails struct {

        
            OrderID string  `json:"order_id"`
            ShipmentCreatedAt string  `json:"shipment_created_at"`
            ItemQty float64  `json:"item_qty"`
            InvoiceID string  `json:"invoice_id"`
            Awb string  `json:"awb"`
            ShipmentID string  `json:"shipment_id"`
         
    }
    
    // ManifestShipmentListing used by Order
    type ManifestShipmentListing struct {

        
            Status float64  `json:"status"`
            TotalCount float64  `json:"total_count"`
            Lane string  `json:"lane"`
            Success bool  `json:"success"`
            Page ManifestPageInfo  `json:"page"`
            Items []ManifestItemDetails  `json:"items"`
            Message string  `json:"message"`
         
    }
    
    // ManifestFile used by Order
    type ManifestFile struct {

        
            Key string  `json:"key"`
            Bucket string  `json:"bucket"`
            Region string  `json:"region"`
         
    }
    
    // ManifestMediaUpdate used by Order
    type ManifestMediaUpdate struct {

        
            Status bool  `json:"status"`
            MediaType string  `json:"media_type"`
            Link string  `json:"link"`
            Entity string  `json:"entity"`
            File ManifestFile  `json:"file"`
            Code float64  `json:"code"`
         
    }
    
    // PDFMeta used by Order
    type PDFMeta struct {

        
            Consent string  `json:"consent"`
            MediaUpdates []ManifestMediaUpdate  `json:"media_updates"`
         
    }
    
    // TotalShipmentPricesCount used by Order
    type TotalShipmentPricesCount struct {

        
            TotalPrice float64  `json:"total_price"`
            ShipmentCount float64  `json:"shipment_count"`
         
    }
    
    // ManifestMeta used by Order
    type ManifestMeta struct {

        
            TotalShipmentPricesCount TotalShipmentPricesCount  `json:"total_shipment_prices_count"`
            Filters Filters  `json:"filters"`
         
    }
    
    // Manifest used by Order
    type Manifest struct {

        
            Status string  `json:"status"`
            CreatedBy string  `json:"created_by"`
            UserID string  `json:"user_id"`
            CompanyID float64  `json:"company_id"`
            CreatedAt string  `json:"created_at"`
            ID float64  `json:"id"`
            Filters Filters  `json:"filters"`
            IsActive bool  `json:"is_active"`
            PdfMeta PDFMeta  `json:"pdf_meta"`
            Meta ManifestMeta  `json:"meta"`
            UID string  `json:"uid"`
            ManifestID string  `json:"manifest_id"`
         
    }
    
    // ManifestList used by Order
    type ManifestList struct {

        
            Page ManifestPageInfo  `json:"page"`
            Items []Manifest  `json:"items"`
         
    }
    
    // ManifestDetails used by Order
    type ManifestDetails struct {

        
            AdditionalShipmentCount float64  `json:"additional_shipment_count"`
            ManifestDetails []Manifest  `json:"manifest_details"`
            Page ManifestPageInfo  `json:"page"`
            Items []ManifestItemDetails  `json:"items"`
         
    }
    
    // FetchCreditBalanceRequestPayload used by Order
    type FetchCreditBalanceRequestPayload struct {

        
            CustomerMobileNumber string  `json:"customer_mobile_number"`
            AffiliateID string  `json:"affiliate_id"`
            SellerID string  `json:"seller_id"`
         
    }
    
    // CreditBalanceInfo used by Order
    type CreditBalanceInfo struct {

        
            CustomerMobileNumber string  `json:"customer_mobile_number"`
            Reason string  `json:"reason"`
            TotalCreditedBalance string  `json:"total_credited_balance"`
         
    }
    
    // FetchCreditBalanceResponsePayload used by Order
    type FetchCreditBalanceResponsePayload struct {

        
            Success bool  `json:"success"`
            Data CreditBalanceInfo  `json:"data"`
         
    }
    
    // RefundModeConfigRequestPayload used by Order
    type RefundModeConfigRequestPayload struct {

        
            CustomerMobileNumber string  `json:"customer_mobile_number"`
            AffiliateID string  `json:"affiliate_id"`
            SellerID float64  `json:"seller_id"`
            FyndOrderID string  `json:"fynd_order_id"`
         
    }
    
    // RefundOption used by Order
    type RefundOption struct {

        
            Type string  `json:"type"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Value string  `json:"value"`
         
    }
    
    // RefundModeInfo used by Order
    type RefundModeInfo struct {

        
            Options []RefundOption  `json:"options"`
            DisplayName string  `json:"display_name"`
            IsActive bool  `json:"is_active"`
            Slug string  `json:"slug"`
         
    }
    
    // RefundModeConfigResponsePayload used by Order
    type RefundModeConfigResponsePayload struct {

        
            Success bool  `json:"success"`
            Data []RefundModeInfo  `json:"data"`
         
    }
    
    // AttachUserInfo used by Order
    type AttachUserInfo struct {

        
            FirstName string  `json:"first_name"`
            LastName string  `json:"last_name"`
            Mobile string  `json:"mobile"`
            CountryCode string  `json:"country_code"`
         
    }
    
    // AttachUserOtpData used by Order
    type AttachUserOtpData struct {

        
            RequestID string  `json:"request_id"`
         
    }
    
    // AttachOrderUser used by Order
    type AttachOrderUser struct {

        
            UserInfo AttachUserInfo  `json:"user_info"`
            FyndOrderID string  `json:"fynd_order_id"`
            OtpData AttachUserOtpData  `json:"otp_data"`
         
    }
    
    // AttachOrderUserResponse used by Order
    type AttachOrderUserResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // SendUserMobileOTP used by Order
    type SendUserMobileOTP struct {

        
            Mobile string  `json:"mobile"`
            CountryCode string  `json:"country_code"`
         
    }
    
    // PointBlankOtpData used by Order
    type PointBlankOtpData struct {

        
            RequestID string  `json:"request_id"`
            ResendTimer float64  `json:"resend_timer"`
            Mobile float64  `json:"mobile"`
            Message string  `json:"message"`
         
    }
    
    // SendUserMobileOtpResponse used by Order
    type SendUserMobileOtpResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
            Data PointBlankOtpData  `json:"data"`
         
    }
    
    // VerifyOtpData used by Order
    type VerifyOtpData struct {

        
            RequestID string  `json:"request_id"`
            OtpCode float64  `json:"otp_code"`
            Mobile string  `json:"mobile"`
         
    }
    
    // VerifyMobileOTP used by Order
    type VerifyMobileOTP struct {

        
            FyndOrderID string  `json:"fynd_order_id"`
            OtpData VerifyOtpData  `json:"otp_data"`
         
    }
    
    // VerifyOtpResponseData used by Order
    type VerifyOtpResponseData struct {

        
            Mobile string  `json:"mobile"`
            FyndOrderID string  `json:"fynd_order_id"`
            Message string  `json:"message"`
            CountryCode string  `json:"country_code"`
         
    }
    
    // VerifyOtpResponse used by Order
    type VerifyOtpResponse struct {

        
            Status float64  `json:"status"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
            Data VerifyOtpResponseData  `json:"data"`
         
    }
    

    
    // DeleteResponse used by Catalog
    type DeleteResponse struct {

        
            Message string  `json:"message"`
         
    }
    
    // GetSearchWordsData used by Catalog
    type GetSearchWordsData struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Result map[string]interface{}  `json:"result"`
            Words []string  `json:"words"`
            UID string  `json:"uid"`
            IsActive bool  `json:"is_active"`
            AppID string  `json:"app_id"`
         
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
    
    // SearchKeywordResult used by Catalog
    type SearchKeywordResult struct {

        
            SortOn string  `json:"sort_on"`
            Query map[string]interface{}  `json:"query"`
         
    }
    
    // CreateSearchKeyword used by Catalog
    type CreateSearchKeyword struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Result SearchKeywordResult  `json:"result"`
            Words []string  `json:"words"`
            IsActive bool  `json:"is_active"`
            AppID string  `json:"app_id"`
         
    }
    
    // GetSearchWordsResponse used by Catalog
    type GetSearchWordsResponse struct {

        
            Items []GetSearchWordsData  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // GetAutocompleteWordsData used by Catalog
    type GetAutocompleteWordsData struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Results []map[string]interface{}  `json:"results"`
            Words []string  `json:"words"`
            UID string  `json:"uid"`
            AppID string  `json:"app_id"`
         
    }
    
    // GetAutocompleteWordsResponse used by Catalog
    type GetAutocompleteWordsResponse struct {

        
            Items []GetAutocompleteWordsData  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // AutocompletePageAction used by Catalog
    type AutocompletePageAction struct {

        
            Type string  `json:"type"`
            URL string  `json:"url"`
            Params map[string]interface{}  `json:"params"`
            Query map[string]interface{}  `json:"query"`
         
    }
    
    // AutocompleteAction used by Catalog
    type AutocompleteAction struct {

        
            Page AutocompletePageAction  `json:"page"`
            Type string  `json:"type"`
         
    }
    
    // AutoCompleteMedia used by Catalog
    type AutoCompleteMedia struct {

        
            URL string  `json:"url"`
            AspectRatio string  `json:"aspect_ratio"`
            Type string  `json:"type"`
         
    }
    
    // AutocompleteResult used by Catalog
    type AutocompleteResult struct {

        
            Action AutocompleteAction  `json:"action"`
            Logo AutoCompleteMedia  `json:"logo"`
            Display string  `json:"display"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // CreateAutocompleteKeyword used by Catalog
    type CreateAutocompleteKeyword struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Words []string  `json:"words"`
            Results []AutocompleteResult  `json:"results"`
            IsActive bool  `json:"is_active"`
            AppID string  `json:"app_id"`
         
    }
    
    // CreateAutocompleteWordsResponse used by Catalog
    type CreateAutocompleteWordsResponse struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Results []map[string]interface{}  `json:"results"`
            Words []string  `json:"words"`
            AppID string  `json:"app_id"`
         
    }
    
    // ProductBundleItem used by Catalog
    type ProductBundleItem struct {

        
            MinQuantity float64  `json:"min_quantity"`
            AllowRemove bool  `json:"allow_remove"`
            AutoSelect bool  `json:"auto_select"`
            ProductUID float64  `json:"product_uid"`
            MaxQuantity float64  `json:"max_quantity"`
            AutoAddToCart bool  `json:"auto_add_to_cart"`
         
    }
    
    // GetProductBundleCreateResponse used by Catalog
    type GetProductBundleCreateResponse struct {

        
            Slug string  `json:"slug"`
            Products []ProductBundleItem  `json:"products"`
            CreatedOn string  `json:"created_on"`
            CompanyID float64  `json:"company_id"`
            Logo string  `json:"logo"`
            ID string  `json:"id"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            PageVisibility []string  `json:"page_visibility"`
            ModifiedOn string  `json:"modified_on"`
            Meta map[string]interface{}  `json:"meta"`
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            Choice string  `json:"choice"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
         
    }
    
    // GetProductBundleListingResponse used by Catalog
    type GetProductBundleListingResponse struct {

        
            Items []GetProductBundleCreateResponse  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // ProductBundleRequest used by Catalog
    type ProductBundleRequest struct {

        
            Slug string  `json:"slug"`
            Products []ProductBundleItem  `json:"products"`
            CreatedOn string  `json:"created_on"`
            CompanyID float64  `json:"company_id"`
            Logo string  `json:"logo"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            PageVisibility []string  `json:"page_visibility"`
            ModifiedOn string  `json:"modified_on"`
            Meta map[string]interface{}  `json:"meta"`
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            Choice string  `json:"choice"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
         
    }
    
    // ProductBundleUpdateRequest used by Catalog
    type ProductBundleUpdateRequest struct {

        
            Slug string  `json:"slug"`
            Products []ProductBundleItem  `json:"products"`
            CompanyID float64  `json:"company_id"`
            Logo string  `json:"logo"`
            PageVisibility []string  `json:"page_visibility"`
            ModifiedOn string  `json:"modified_on"`
            Meta map[string]interface{}  `json:"meta"`
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            Choice string  `json:"choice"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
         
    }
    
    // Size used by Catalog
    type Size struct {

        
            Quantity float64  `json:"quantity"`
            IsAvailable bool  `json:"is_available"`
            Value string  `json:"value"`
            Display string  `json:"display"`
         
    }
    
    // LimitedProductData used by Catalog
    type LimitedProductData struct {

        
            Quantity float64  `json:"quantity"`
            Slug string  `json:"slug"`
            Sizes []string  `json:"sizes"`
            Price map[string]interface{}  `json:"price"`
            Images []string  `json:"images"`
            CountryOfOrigin string  `json:"country_of_origin"`
            Identifier map[string]interface{}  `json:"identifier"`
            ItemCode string  `json:"item_code"`
            ShortDescription string  `json:"short_description"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            Attributes map[string]interface{}  `json:"attributes"`
         
    }
    
    // Price used by Catalog
    type Price struct {

        
            MinEffective float64  `json:"min_effective"`
            MaxMarked float64  `json:"max_marked"`
            MinMarked float64  `json:"min_marked"`
            MaxEffective float64  `json:"max_effective"`
            Currency string  `json:"currency"`
         
    }
    
    // GetProducts used by Catalog
    type GetProducts struct {

        
            Sizes []Size  `json:"sizes"`
            MinQuantity float64  `json:"min_quantity"`
            ProductDetails LimitedProductData  `json:"product_details"`
            Price Price  `json:"price"`
            AllowRemove bool  `json:"allow_remove"`
            AutoSelect bool  `json:"auto_select"`
            ProductUID float64  `json:"product_uid"`
            MaxQuantity float64  `json:"max_quantity"`
            AutoAddToCart bool  `json:"auto_add_to_cart"`
         
    }
    
    // GetProductBundleResponse used by Catalog
    type GetProductBundleResponse struct {

        
            Slug string  `json:"slug"`
            Products []GetProducts  `json:"products"`
            CompanyID float64  `json:"company_id"`
            Logo string  `json:"logo"`
            PageVisibility []string  `json:"page_visibility"`
            Meta map[string]interface{}  `json:"meta"`
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            Choice string  `json:"choice"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
         
    }
    
    // ListSizeGuide used by Catalog
    type ListSizeGuide struct {

        
            Items []map[string]interface{}  `json:"items"`
            Page map[string]interface{}  `json:"page"`
         
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

        
            Tag string  `json:"tag"`
            Active bool  `json:"active"`
            Subtitle string  `json:"subtitle"`
            CompanyID float64  `json:"company_id"`
            ID string  `json:"id"`
            Image string  `json:"image"`
            Title string  `json:"title"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Guide Guide  `json:"guide"`
            Description string  `json:"description"`
            BrandID float64  `json:"brand_id"`
            CreatedOn string  `json:"created_on"`
            Name string  `json:"name"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // SuccessResponse used by Catalog
    type SuccessResponse struct {

        
            Success bool  `json:"success"`
         
    }
    
    // SizeGuideResponse used by Catalog
    type SizeGuideResponse struct {

        
            Tag string  `json:"tag"`
            Active bool  `json:"active"`
            CompanyID float64  `json:"company_id"`
            Subtitle string  `json:"subtitle"`
            ID string  `json:"id"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Title string  `json:"title"`
            Guide map[string]interface{}  `json:"guide"`
            BrandID float64  `json:"brand_id"`
            CreatedOn string  `json:"created_on"`
            Name string  `json:"name"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // SEOData used by Catalog
    type SEOData struct {

        
            Title string  `json:"title"`
            Description string  `json:"description"`
         
    }
    
    // MOQData used by Catalog
    type MOQData struct {

        
            Maximum float64  `json:"maximum"`
            Minimum float64  `json:"minimum"`
            IncrementUnit float64  `json:"increment_unit"`
         
    }
    
    // OwnerAppItemResponse used by Catalog
    type OwnerAppItemResponse struct {

        
            Seo SEOData  `json:"seo"`
            IsCod bool  `json:"is_cod"`
            IsGift bool  `json:"is_gift"`
            AltText map[string]interface{}  `json:"alt_text"`
            Moq MOQData  `json:"moq"`
         
    }
    
    // ApplicationItemSEO used by Catalog
    type ApplicationItemSEO struct {

        
            Title string  `json:"title"`
            Description string  `json:"description"`
         
    }
    
    // ApplicationItemMOQ used by Catalog
    type ApplicationItemMOQ struct {

        
            Maximum float64  `json:"maximum"`
            Minimum float64  `json:"minimum"`
            IncrementUnit float64  `json:"increment_unit"`
         
    }
    
    // MetaFields used by Catalog
    type MetaFields struct {

        
            Value string  `json:"value"`
            Key string  `json:"key"`
         
    }
    
    // ApplicationItemMeta used by Catalog
    type ApplicationItemMeta struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Seo ApplicationItemSEO  `json:"seo"`
            IsCod bool  `json:"is_cod"`
            IsGift bool  `json:"is_gift"`
            AltText map[string]interface{}  `json:"alt_text"`
            Moq ApplicationItemMOQ  `json:"moq"`
            CustomMeta []MetaFields  `json:"_custom_meta"`
         
    }
    
    // SuccessResponse1 used by Catalog
    type SuccessResponse1 struct {

        
            Success bool  `json:"success"`
            UID float64  `json:"uid"`
         
    }
    
    // GetConfigMetadataResponse used by Catalog
    type GetConfigMetadataResponse struct {

        
            Data []map[string]interface{}  `json:"data"`
            Values []map[string]interface{}  `json:"values"`
            Condition []map[string]interface{}  `json:"condition"`
         
    }
    
    // PageResponseType used by Catalog
    type PageResponseType struct {

        
            Next float64  `json:"next"`
            Current float64  `json:"current"`
            TotalCount float64  `json:"total_count"`
            HasNext bool  `json:"has_next"`
         
    }
    
    // GetConfigResponse used by Catalog
    type GetConfigResponse struct {

        
            Data []map[string]interface{}  `json:"data"`
            Page PageResponseType  `json:"page"`
         
    }
    
    // ConfigErrorResponse used by Catalog
    type ConfigErrorResponse struct {

        
            Message string  `json:"message"`
         
    }
    
    // AttributeDetailsGroup used by Catalog
    type AttributeDetailsGroup struct {

        
            Slug string  `json:"slug"`
            DisplayType string  `json:"display_type"`
            Logo string  `json:"logo"`
            Priority float64  `json:"priority"`
            Key string  `json:"key"`
            Unit string  `json:"unit"`
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
         
    }
    
    // AppConfigurationDetail used by Catalog
    type AppConfigurationDetail struct {

        
            Slug string  `json:"slug"`
            TemplateSlugs []string  `json:"template_slugs"`
            IsDefault bool  `json:"is_default"`
            Logo string  `json:"logo"`
            Priority float64  `json:"priority"`
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            Attributes []AttributeDetailsGroup  `json:"attributes"`
            AppID string  `json:"app_id"`
         
    }
    
    // ConfigSuccessResponse used by Catalog
    type ConfigSuccessResponse struct {

        
            Message string  `json:"message"`
         
    }
    
    // AppConfigurationsSort used by Catalog
    type AppConfigurationsSort struct {

        
            IsDefault bool  `json:"is_default"`
            Logo string  `json:"logo"`
            Priority float64  `json:"priority"`
            Key string  `json:"key"`
            DefaultKey string  `json:"default_key"`
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            AppID string  `json:"app_id"`
         
    }
    
    // AllowSingleRequest used by Catalog
    type AllowSingleRequest struct {

        
            AllowSingle bool  `json:"allow_single"`
         
    }
    
    // DefaultKeyRequest used by Catalog
    type DefaultKeyRequest struct {

        
            DefaultKey string  `json:"default_key"`
         
    }
    
    // MetaDataListingFilterMetaResponse used by Catalog
    type MetaDataListingFilterMetaResponse struct {

        
            Units []map[string]interface{}  `json:"units"`
            Display string  `json:"display"`
            FilterTypes []string  `json:"filter_types"`
            Key string  `json:"key"`
         
    }
    
    // MetaDataListingFilterResponse used by Catalog
    type MetaDataListingFilterResponse struct {

        
            Data []MetaDataListingFilterMetaResponse  `json:"data"`
         
    }
    
    // MetaDataListingSortMetaResponse used by Catalog
    type MetaDataListingSortMetaResponse struct {

        
            Display string  `json:"display"`
            Key string  `json:"key"`
         
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
    
    // GetCatalogConfigurationDetailsProduct used by Catalog
    type GetCatalogConfigurationDetailsProduct struct {

        
            Variant map[string]interface{}  `json:"variant"`
            Similar map[string]interface{}  `json:"similar"`
            Compare map[string]interface{}  `json:"compare"`
            Detail map[string]interface{}  `json:"detail"`
         
    }
    
    // GetCatalogConfigurationMetaData used by Catalog
    type GetCatalogConfigurationMetaData struct {

        
            Listing MetaDataListingResponse  `json:"listing"`
            Product GetCatalogConfigurationDetailsProduct  `json:"product"`
         
    }
    
    // ProductSize used by Catalog
    type ProductSize struct {

        
            Min float64  `json:"min"`
            Max float64  `json:"max"`
         
    }
    
    // ConfigurationProductVariantConfig used by Catalog
    type ConfigurationProductVariantConfig struct {

        
            DisplayType string  `json:"display_type"`
            Logo string  `json:"logo"`
            Priority float64  `json:"priority"`
            Key string  `json:"key"`
            Size ProductSize  `json:"size"`
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
         
    }
    
    // ConfigurationProductVariant used by Catalog
    type ConfigurationProductVariant struct {

        
            Config []ConfigurationProductVariantConfig  `json:"config"`
         
    }
    
    // ConfigurationProductConfig used by Catalog
    type ConfigurationProductConfig struct {

        
            Logo string  `json:"logo"`
            Subtitle string  `json:"subtitle"`
            Priority float64  `json:"priority"`
            Key string  `json:"key"`
            Size ProductSize  `json:"size"`
            Title string  `json:"title"`
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

        
            End float64  `json:"end"`
            Display string  `json:"display"`
            Start float64  `json:"start"`
         
    }
    
    // ConfigurationListingFilterValue used by Catalog
    type ConfigurationListingFilterValue struct {

        
            Map map[string]interface{}  `json:"map"`
            MapValues []map[string]interface{}  `json:"map_values"`
            Value string  `json:"value"`
            Priority []string  `json:"priority"`
            Condition string  `json:"condition"`
            Sort string  `json:"sort"`
            BucketPoints []ConfigurationBucketPoints  `json:"bucket_points"`
         
    }
    
    // ConfigurationListingFilterConfig used by Catalog
    type ConfigurationListingFilterConfig struct {

        
            ValueConfig ConfigurationListingFilterValue  `json:"value_config"`
            Logo string  `json:"logo"`
            Priority float64  `json:"priority"`
            Key string  `json:"key"`
            Type string  `json:"type"`
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            DisplayName string  `json:"display_name"`
         
    }
    
    // ConfigurationListingFilter used by Catalog
    type ConfigurationListingFilter struct {

        
            AllowSingle bool  `json:"allow_single"`
            AttributeConfig []ConfigurationListingFilterConfig  `json:"attribute_config"`
         
    }
    
    // ConfigurationListingSortConfig used by Catalog
    type ConfigurationListingSortConfig struct {

        
            Logo string  `json:"logo"`
            Priority float64  `json:"priority"`
            Key string  `json:"key"`
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
         
    }
    
    // ConfigurationListingSort used by Catalog
    type ConfigurationListingSort struct {

        
            Config []ConfigurationListingSortConfig  `json:"config"`
            DefaultKey string  `json:"default_key"`
         
    }
    
    // ConfigurationListing used by Catalog
    type ConfigurationListing struct {

        
            Filter ConfigurationListingFilter  `json:"filter"`
            Sort ConfigurationListingSort  `json:"sort"`
         
    }
    
    // AppCatalogConfiguration used by Catalog
    type AppCatalogConfiguration struct {

        
            ConfigID string  `json:"config_id"`
            Product ConfigurationProduct  `json:"product"`
            ID string  `json:"id"`
            ConfigType string  `json:"config_type"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Type string  `json:"type"`
            Listing ConfigurationListing  `json:"listing"`
            CreatedOn string  `json:"created_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            ModifiedOn string  `json:"modified_on"`
            AppID string  `json:"app_id"`
         
    }
    
    // GetAppCatalogConfiguration used by Catalog
    type GetAppCatalogConfiguration struct {

        
            IsDefault bool  `json:"is_default"`
            Data AppCatalogConfiguration  `json:"data"`
         
    }
    
    // AppConfiguration used by Catalog
    type AppConfiguration struct {

        
            ConfigID string  `json:"config_id"`
            Product ConfigurationProduct  `json:"product"`
            ConfigType string  `json:"config_type"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Type string  `json:"type"`
            Listing ConfigurationListing  `json:"listing"`
            CreatedOn string  `json:"created_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            ModifiedOn string  `json:"modified_on"`
            AppID string  `json:"app_id"`
         
    }
    
    // GetCatalogConfigurationDetailsSchemaListing used by Catalog
    type GetCatalogConfigurationDetailsSchemaListing struct {

        
            Filter map[string]interface{}  `json:"filter"`
            Sort map[string]interface{}  `json:"sort"`
         
    }
    
    // EntityConfiguration used by Catalog
    type EntityConfiguration struct {

        
            ConfigID string  `json:"config_id"`
            Product GetCatalogConfigurationDetailsProduct  `json:"product"`
            ID string  `json:"id"`
            ConfigType string  `json:"config_type"`
            Listing GetCatalogConfigurationDetailsSchemaListing  `json:"listing"`
            AppID string  `json:"app_id"`
         
    }
    
    // GetAppCatalogEntityConfiguration used by Catalog
    type GetAppCatalogEntityConfiguration struct {

        
            IsDefault bool  `json:"is_default"`
            Data EntityConfiguration  `json:"data"`
         
    }
    
    // ProductFiltersValue used by Catalog
    type ProductFiltersValue struct {

        
            Count float64  `json:"count"`
            Max float64  `json:"max"`
            CurrencyCode string  `json:"currency_code"`
            Value interface{}  `json:"value"`
            IsSelected bool  `json:"is_selected"`
            DisplayFormat string  `json:"display_format"`
            SelectedMin float64  `json:"selected_min"`
            SelectedMax float64  `json:"selected_max"`
            CurrencySymbol string  `json:"currency_symbol"`
            Min float64  `json:"min"`
            Display string  `json:"display"`
            QueryFormat string  `json:"query_format"`
         
    }
    
    // ProductFiltersKey used by Catalog
    type ProductFiltersKey struct {

        
            Operators []string  `json:"operators"`
            Logo string  `json:"logo"`
            Kind string  `json:"kind"`
            Name string  `json:"name"`
            Display string  `json:"display"`
         
    }
    
    // ProductFilters used by Catalog
    type ProductFilters struct {

        
            Values []ProductFiltersValue  `json:"values"`
            Key ProductFiltersKey  `json:"key"`
         
    }
    
    // ProductSortOn used by Catalog
    type ProductSortOn struct {

        
            Name string  `json:"name"`
            Value string  `json:"value"`
            IsSelected bool  `json:"is_selected"`
         
    }
    
    // GetCollectionQueryOptionResponse used by Catalog
    type GetCollectionQueryOptionResponse struct {

        
            Operators map[string]string  `json:"operators"`
            Filters []ProductFilters  `json:"filters"`
            SortOn []ProductSortOn  `json:"sort_on"`
         
    }
    
    // CollectionScheduleStartEnd used by Catalog
    type CollectionScheduleStartEnd struct {

        
            End string  `json:"end"`
            Start string  `json:"start"`
         
    }
    
    // CollectionSchedule used by Catalog
    type CollectionSchedule struct {

        
            NextSchedule []CollectionScheduleStartEnd  `json:"next_schedule"`
            End string  `json:"end"`
            Start string  `json:"start"`
         
    }
    
    // CollectionBadge used by Catalog
    type CollectionBadge struct {

        
            Color string  `json:"color"`
            Text string  `json:"text"`
         
    }
    
    // Media used by Catalog
    type Media struct {

        
            Meta map[string]interface{}  `json:"meta"`
            URL string  `json:"url"`
            Type string  `json:"type"`
         
    }
    
    // CollectionActionPageQuery used by Catalog
    type CollectionActionPageQuery struct {

        
            Op string  `json:"op"`
            Value []string  `json:"value"`
            Attribute string  `json:"attribute"`
         
    }
    
    // CollectionActionPage used by Catalog
    type CollectionActionPage struct {

        
            Type string  `json:"type"`
            Query CollectionActionPageQuery  `json:"query"`
         
    }
    
    // CollectionQuery used by Catalog
    type CollectionQuery struct {

        
            Op string  `json:"op"`
            Value []interface{}  `json:"value"`
            Attribute string  `json:"attribute"`
         
    }
    
    // BannerImage used by Catalog
    type BannerImage struct {

        
            URL string  `json:"url"`
            AspectRatio string  `json:"aspect_ratio"`
            Type string  `json:"type"`
         
    }
    
    // ImageUrls used by Catalog
    type ImageUrls struct {

        
            Portrait BannerImage  `json:"portrait"`
            Landscape BannerImage  `json:"landscape"`
         
    }
    
    // GetCollectionDetailNest used by Catalog
    type GetCollectionDetailNest struct {

        
            Slug string  `json:"slug"`
            AllowSort bool  `json:"allow_sort"`
            Description string  `json:"description"`
            Meta map[string]interface{}  `json:"meta"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            AllowFacets bool  `json:"allow_facets"`
            Priority float64  `json:"priority"`
            Name string  `json:"name"`
            AppID string  `json:"app_id"`
            Schedule CollectionSchedule  `json:"_schedule"`
            Badge CollectionBadge  `json:"badge"`
            Logo Media  `json:"logo"`
            Type string  `json:"type"`
            Action CollectionActionPage  `json:"action"`
            Query []CollectionQuery  `json:"query"`
            Tags []string  `json:"tags"`
            UID string  `json:"uid"`
            IsActive bool  `json:"is_active"`
            Banners ImageUrls  `json:"banners"`
         
    }
    
    // CollectionListingFilterTag used by Catalog
    type CollectionListingFilterTag struct {

        
            Name string  `json:"name"`
            Display string  `json:"display"`
            IsSelected bool  `json:"is_selected"`
         
    }
    
    // CollectionListingFilterType used by Catalog
    type CollectionListingFilterType struct {

        
            Name string  `json:"name"`
            Display string  `json:"display"`
            IsSelected bool  `json:"is_selected"`
         
    }
    
    // CollectionListingFilter used by Catalog
    type CollectionListingFilter struct {

        
            Tags []CollectionListingFilterTag  `json:"tags"`
            Type []CollectionListingFilterType  `json:"type"`
         
    }
    
    // GetCollectionListingResponse used by Catalog
    type GetCollectionListingResponse struct {

        
            Items []GetCollectionDetailNest  `json:"items"`
            Filters CollectionListingFilter  `json:"filters"`
            Page Page  `json:"page"`
         
    }
    
    // SeoDetail used by Catalog
    type SeoDetail struct {

        
            Title string  `json:"title"`
            Description string  `json:"description"`
         
    }
    
    // CollectionBadge1 used by Catalog
    type CollectionBadge1 struct {

        
            Color string  `json:"color"`
            Text string  `json:"text"`
         
    }
    
    // NextSchedule used by Catalog
    type NextSchedule struct {

        
            End string  `json:"end"`
            Start string  `json:"start"`
         
    }
    
    // CollectionSchedule1 used by Catalog
    type CollectionSchedule1 struct {

        
            End string  `json:"end"`
            Cron string  `json:"cron"`
            Start string  `json:"start"`
            Duration float64  `json:"duration"`
            NextSchedule []NextSchedule  `json:"next_schedule"`
         
    }
    
    // CollectionImage used by Catalog
    type CollectionImage struct {

        
            URL string  `json:"url"`
            AspectRatio string  `json:"aspect_ratio"`
         
    }
    
    // UserInfo used by Catalog
    type UserInfo struct {

        
            Username string  `json:"username"`
            Email string  `json:"email"`
            UserID string  `json:"user_id"`
            UID string  `json:"uid"`
         
    }
    
    // CollectionBanner used by Catalog
    type CollectionBanner struct {

        
            Portrait CollectionImage  `json:"portrait"`
            Landscape CollectionImage  `json:"landscape"`
         
    }
    
    // CreateCollection used by Catalog
    type CreateCollection struct {

        
            Slug string  `json:"slug"`
            Seo SeoDetail  `json:"seo"`
            AllowSort bool  `json:"allow_sort"`
            Description string  `json:"description"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Meta map[string]interface{}  `json:"meta"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            AllowFacets bool  `json:"allow_facets"`
            Priority float64  `json:"priority"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            IsVisible bool  `json:"is_visible"`
            Name string  `json:"name"`
            AppID string  `json:"app_id"`
            Badge CollectionBadge1  `json:"badge"`
            Schedule CollectionSchedule1  `json:"_schedule"`
            SortOn string  `json:"sort_on"`
            Logo CollectionImage  `json:"logo"`
            Type string  `json:"type"`
            Published bool  `json:"published"`
            ModifiedBy UserInfo  `json:"modified_by"`
            Query []CollectionQuery  `json:"query"`
            Tags []string  `json:"tags"`
            CreatedBy UserInfo  `json:"created_by"`
            IsActive bool  `json:"is_active"`
            Banners CollectionBanner  `json:"banners"`
         
    }
    
    // CollectionCreateResponse used by Catalog
    type CollectionCreateResponse struct {

        
            Slug string  `json:"slug"`
            AllowSort bool  `json:"allow_sort"`
            Description string  `json:"description"`
            Meta map[string]interface{}  `json:"meta"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Cron map[string]interface{}  `json:"cron"`
            AllowFacets bool  `json:"allow_facets"`
            Priority float64  `json:"priority"`
            Name string  `json:"name"`
            SortOn string  `json:"sort_on"`
            Schedule map[string]interface{}  `json:"_schedule"`
            Tag []string  `json:"tag"`
            Badge map[string]interface{}  `json:"badge"`
            AppID string  `json:"app_id"`
            Logo BannerImage  `json:"logo"`
            Type string  `json:"type"`
            Query []CollectionQuery  `json:"query"`
            IsActive bool  `json:"is_active"`
            Banners ImageUrls  `json:"banners"`
         
    }
    
    // CollectionDetailResponse used by Catalog
    type CollectionDetailResponse struct {

        
            Schedule CollectionSchedule  `json:"_schedule"`
            Slug string  `json:"slug"`
            Badge CollectionBadge  `json:"badge"`
            Tags []string  `json:"tags"`
            Logo Media  `json:"logo"`
            Type string  `json:"type"`
            Priority float64  `json:"priority"`
            AllowFacets bool  `json:"allow_facets"`
            Query []CollectionQuery  `json:"query"`
            UID string  `json:"uid"`
            AllowSort bool  `json:"allow_sort"`
            Description string  `json:"description"`
            Meta map[string]interface{}  `json:"meta"`
            IsActive bool  `json:"is_active"`
            Banners ImageUrls  `json:"banners"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Name string  `json:"name"`
            AppID string  `json:"app_id"`
         
    }
    
    // UpdateCollection used by Catalog
    type UpdateCollection struct {

        
            Slug string  `json:"slug"`
            Seo SeoDetail  `json:"seo"`
            AllowSort bool  `json:"allow_sort"`
            Description string  `json:"description"`
            Meta map[string]interface{}  `json:"meta"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            AllowFacets bool  `json:"allow_facets"`
            Priority float64  `json:"priority"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            IsVisible bool  `json:"is_visible"`
            Name string  `json:"name"`
            SortOn string  `json:"sort_on"`
            Schedule CollectionSchedule1  `json:"_schedule"`
            Badge CollectionBadge1  `json:"badge"`
            Logo CollectionImage  `json:"logo"`
            Type string  `json:"type"`
            Published bool  `json:"published"`
            ModifiedBy UserInfo  `json:"modified_by"`
            Query []CollectionQuery  `json:"query"`
            Tags []string  `json:"tags"`
            IsActive bool  `json:"is_active"`
            Banners CollectionBanner  `json:"banners"`
         
    }
    
    // ProductDetailAttribute used by Catalog
    type ProductDetailAttribute struct {

        
            Value string  `json:"value"`
            Key string  `json:"key"`
            Type string  `json:"type"`
         
    }
    
    // ProductDetailGroupedAttribute used by Catalog
    type ProductDetailGroupedAttribute struct {

        
            Title string  `json:"title"`
            Details []ProductDetailAttribute  `json:"details"`
         
    }
    
    // Price1 used by Catalog
    type Price1 struct {

        
            CurrencyCode string  `json:"currency_code"`
            Min float64  `json:"min"`
            CurrencySymbol string  `json:"currency_symbol"`
            Max float64  `json:"max"`
         
    }
    
    // ProductListingPrice used by Catalog
    type ProductListingPrice struct {

        
            Effective Price1  `json:"effective"`
            Marked Price1  `json:"marked"`
         
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

        
            Action Action  `json:"action"`
            Logo Media  `json:"logo"`
            Name string  `json:"name"`
            UID float64  `json:"uid"`
         
    }
    
    // ProductListingDetail used by Catalog
    type ProductListingDetail struct {

        
            Slug string  `json:"slug"`
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
            Color string  `json:"color"`
            Discount string  `json:"discount"`
            Description string  `json:"description"`
            ItemType string  `json:"item_type"`
            Sellable bool  `json:"sellable"`
            Name string  `json:"name"`
            PromoMeta map[string]interface{}  `json:"promo_meta"`
            Price ProductListingPrice  `json:"price"`
            ImageNature string  `json:"image_nature"`
            Highlights []string  `json:"highlights"`
            HasVariant bool  `json:"has_variant"`
            ItemCode string  `json:"item_code"`
            Type string  `json:"type"`
            ProductOnlineDate string  `json:"product_online_date"`
            Similars []string  `json:"similars"`
            Rating float64  `json:"rating"`
            Attributes map[string]interface{}  `json:"attributes"`
            Medias []Media  `json:"medias"`
            ShortDescription string  `json:"short_description"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            UID float64  `json:"uid"`
            Brand ProductBrand  `json:"brand"`
            RatingCount float64  `json:"rating_count"`
            Tryouts []string  `json:"tryouts"`
         
    }
    
    // GetCollectionItemsResponse used by Catalog
    type GetCollectionItemsResponse struct {

        
            Items []ProductListingDetail  `json:"items"`
            Filters []ProductFilters  `json:"filters"`
            Page Page  `json:"page"`
            SortOn []ProductSortOn  `json:"sort_on"`
         
    }
    
    // CollectionItem used by Catalog
    type CollectionItem struct {

        
            Action string  `json:"action"`
            Priority float64  `json:"priority"`
            ItemID float64  `json:"item_id"`
         
    }
    
    // CollectionItemUpdate used by Catalog
    type CollectionItemUpdate struct {

        
            Items []CollectionItem  `json:"items"`
            AllowFacets bool  `json:"allow_facets"`
            AllowSort bool  `json:"allow_sort"`
            Type string  `json:"type"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Query []CollectionQuery  `json:"query"`
         
    }
    
    // UpdatedResponse used by Catalog
    type UpdatedResponse struct {

        
            Message string  `json:"message"`
            ItemsNotUpdated []float64  `json:"items_not_updated"`
         
    }
    
    // CatalogInsightItem used by Catalog
    type CatalogInsightItem struct {

        
            Count float64  `json:"count"`
            OutOfStockCount float64  `json:"out_of_stock_count"`
            SellableCount float64  `json:"sellable_count"`
         
    }
    
    // CatalogInsightBrand used by Catalog
    type CatalogInsightBrand struct {

        
            TotalSizes float64  `json:"total_sizes"`
            TotalArticles float64  `json:"total_articles"`
            ArticleFreshness float64  `json:"article_freshness"`
            Name string  `json:"name"`
            AvailableSizes float64  `json:"available_sizes"`
            AvailableArticles float64  `json:"available_articles"`
         
    }
    
    // CatalogInsightResponse used by Catalog
    type CatalogInsightResponse struct {

        
            Item CatalogInsightItem  `json:"item"`
            BrandDistribution CatalogInsightBrand  `json:"brand_distribution"`
         
    }
    
    // CrossSellingData used by Catalog
    type CrossSellingData struct {

        
            Articles float64  `json:"articles"`
            Products float64  `json:"products"`
         
    }
    
    // CrossSellingResponse used by Catalog
    type CrossSellingResponse struct {

        
            Data CrossSellingData  `json:"data"`
            BrandDistribution CatalogInsightBrand  `json:"brand_distribution"`
         
    }
    
    // OptInPostRequest used by Catalog
    type OptInPostRequest struct {

        
            Platform string  `json:"platform"`
            StoreIds []float64  `json:"store_ids"`
            CompanyID float64  `json:"company_id"`
            OptLevel string  `json:"opt_level"`
            BrandIds []float64  `json:"brand_ids"`
            Enabled bool  `json:"enabled"`
         
    }
    
    // CompanyOptIn used by Catalog
    type CompanyOptIn struct {

        
            Platform string  `json:"platform"`
            StoreIds []float64  `json:"store_ids"`
            CompanyID float64  `json:"company_id"`
            OptLevel string  `json:"opt_level"`
            BrandIds []float64  `json:"brand_ids"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            CreatedOn float64  `json:"created_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            ModifiedOn float64  `json:"modified_on"`
            Enabled bool  `json:"enabled"`
         
    }
    
    // GetOptInPlatform used by Catalog
    type GetOptInPlatform struct {

        
            Items []CompanyOptIn  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // OptinCompanyDetail used by Catalog
    type OptinCompanyDetail struct {

        
            BusinessType string  `json:"business_type"`
            CompanyType string  `json:"company_type"`
            Name string  `json:"name"`
            UID float64  `json:"uid"`
         
    }
    
    // CompanyBrandDetail used by Catalog
    type CompanyBrandDetail struct {

        
            BrandID float64  `json:"brand_id"`
            CompanyID float64  `json:"company_id"`
            TotalArticle float64  `json:"total_article"`
            BrandName string  `json:"brand_name"`
         
    }
    
    // OptinCompanyBrandDetailsView used by Catalog
    type OptinCompanyBrandDetailsView struct {

        
            Items []CompanyBrandDetail  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // OptinCompanyMetrics used by Catalog
    type OptinCompanyMetrics struct {

        
            Company string  `json:"company"`
            Brand float64  `json:"brand"`
            Store float64  `json:"store"`
         
    }
    
    // StoreDetail used by Catalog
    type StoreDetail struct {

        
            StoreType string  `json:"store_type"`
            AdditionalContacts []map[string]interface{}  `json:"additional_contacts"`
            DisplayName string  `json:"display_name"`
            CompanyID float64  `json:"company_id"`
            Timing map[string]interface{}  `json:"timing"`
            Manager map[string]interface{}  `json:"manager"`
            UID float64  `json:"uid"`
            StoreCode string  `json:"store_code"`
            CreatedOn string  `json:"created_on"`
            Name string  `json:"name"`
            Address map[string]interface{}  `json:"address"`
            ModifiedOn string  `json:"modified_on"`
            Documents []map[string]interface{}  `json:"documents"`
         
    }
    
    // OptinStoreDetails used by Catalog
    type OptinStoreDetails struct {

        
            Items []StoreDetail  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // AttributeSchemaRange used by Catalog
    type AttributeSchemaRange struct {

        
            Min float64  `json:"min"`
            Max float64  `json:"max"`
         
    }
    
    // AttributeMaster used by Catalog
    type AttributeMaster struct {

        
            Multi bool  `json:"multi"`
            Format string  `json:"format"`
            AllowedValues []string  `json:"allowed_values"`
            Type string  `json:"type"`
            Range AttributeSchemaRange  `json:"range"`
            Mandatory bool  `json:"mandatory"`
         
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

        
            Enriched bool  `json:"enriched"`
            MandatoryDetails AttributeMasterMandatoryDetails  `json:"mandatory_details"`
         
    }
    
    // AttributeMasterFilter used by Catalog
    type AttributeMasterFilter struct {

        
            Indexing bool  `json:"indexing"`
            Priority float64  `json:"priority"`
            DependsOn []string  `json:"depends_on"`
         
    }
    
    // GenderDetail used by Catalog
    type GenderDetail struct {

        
            Slug string  `json:"slug"`
            EnabledForEndConsumer bool  `json:"enabled_for_end_consumer"`
            Departments []string  `json:"departments"`
            Logo string  `json:"logo"`
            Schema AttributeMaster  `json:"schema"`
            ID string  `json:"id"`
            Details AttributeMasterDetails  `json:"details"`
            IsNested bool  `json:"is_nested"`
            Description string  `json:"description"`
            Meta AttributeMasterMeta  `json:"meta"`
            Name string  `json:"name"`
            Filters AttributeMasterFilter  `json:"filters"`
         
    }
    
    // CategoriesResponse used by Catalog
    type CategoriesResponse struct {

        
            Slug string  `json:"slug"`
            SlugKey string  `json:"slug_key"`
            UID float64  `json:"uid"`
            TemplateSlug string  `json:"template_slug"`
            Name string  `json:"name"`
         
    }
    
    // ProdcutTemplateCategoriesResponse used by Catalog
    type ProdcutTemplateCategoriesResponse struct {

        
            Items []CategoriesResponse  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // PTErrorResponse used by Catalog
    type PTErrorResponse struct {

        
            Message string  `json:"message"`
            Code string  `json:"code"`
            Errors map[string]interface{}  `json:"errors"`
            Meta map[string]interface{}  `json:"meta"`
            Status float64  `json:"status"`
         
    }
    
    // UserSerializer used by Catalog
    type UserSerializer struct {

        
            UserID string  `json:"user_id"`
            Contact string  `json:"contact"`
            Username string  `json:"username"`
            ID string  `json:"_id"`
            UID string  `json:"uid"`
         
    }
    
    // GetDepartment used by Catalog
    type GetDepartment struct {

        
            PriorityOrder float64  `json:"priority_order"`
            Slug string  `json:"slug"`
            ItemType string  `json:"item_type"`
            PageSize float64  `json:"page_size"`
            CreatedOn string  `json:"created_on"`
            Logo string  `json:"logo"`
            PageNo float64  `json:"page_no"`
            CreatedBy UserSerializer  `json:"created_by"`
            Search string  `json:"search"`
            UID float64  `json:"uid"`
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            ModifiedOn string  `json:"modified_on"`
            Synonyms []string  `json:"synonyms"`
         
    }
    
    // DepartmentsResponse used by Catalog
    type DepartmentsResponse struct {

        
            Items []GetDepartment  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // DepartmentErrorResponse used by Catalog
    type DepartmentErrorResponse struct {

        
            Message string  `json:"message"`
            Code string  `json:"code"`
            Errors map[string]interface{}  `json:"errors"`
            Meta map[string]interface{}  `json:"meta"`
            Status float64  `json:"status"`
         
    }
    
    // DepartmentCreateUpdate used by Catalog
    type DepartmentCreateUpdate struct {

        
            PriorityOrder float64  `json:"priority_order"`
            Slug string  `json:"slug"`
            Tags []string  `json:"tags"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Platforms map[string]interface{}  `json:"platforms"`
            Logo string  `json:"logo"`
            Cls string  `json:"_cls"`
            UID float64  `json:"uid"`
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            Synonyms []string  `json:"synonyms"`
         
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

        
            Contact string  `json:"contact"`
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
            SuperUser bool  `json:"super_user"`
         
    }
    
    // DepartmentModel used by Catalog
    type DepartmentModel struct {

        
            PriorityOrder float64  `json:"priority_order"`
            Slug string  `json:"slug"`
            VerifiedOn string  `json:"verified_on"`
            VerifiedBy UserDetail  `json:"verified_by"`
            IsActive bool  `json:"is_active"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Logo string  `json:"logo"`
            Cls string  `json:"_cls"`
            CreatedBy UserDetail  `json:"created_by"`
            ID interface{}  `json:"_id"`
            UID float64  `json:"uid"`
            CreatedOn string  `json:"created_on"`
            Name string  `json:"name"`
            ModifiedBy UserDetail  `json:"modified_by"`
            ModifiedOn string  `json:"modified_on"`
            Synonyms []string  `json:"synonyms"`
         
    }
    
    // ProductTemplate used by Catalog
    type ProductTemplate struct {

        
            Slug string  `json:"slug"`
            Tag string  `json:"tag"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            CreatedOn string  `json:"created_on"`
            Departments []string  `json:"departments"`
            Logo string  `json:"logo"`
            IsExpirable bool  `json:"is_expirable"`
            Categories []string  `json:"categories"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Description string  `json:"description"`
            IsArchived bool  `json:"is_archived"`
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            Attributes []string  `json:"attributes"`
            ModifiedOn string  `json:"modified_on"`
            IsPhysical bool  `json:"is_physical"`
         
    }
    
    // TemplatesResponse used by Catalog
    type TemplatesResponse struct {

        
            Items ProductTemplate  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // TemplateDetails used by Catalog
    type TemplateDetails struct {

        
            Slug string  `json:"slug"`
            Tag string  `json:"tag"`
            Departments []string  `json:"departments"`
            Logo string  `json:"logo"`
            IsExpirable bool  `json:"is_expirable"`
            ID string  `json:"id"`
            Categories []string  `json:"categories"`
            Description string  `json:"description"`
            IsArchived bool  `json:"is_archived"`
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            Attributes []string  `json:"attributes"`
            IsPhysical bool  `json:"is_physical"`
         
    }
    
    // Properties used by Catalog
    type Properties struct {

        
            Sizes map[string]interface{}  `json:"sizes"`
            Slug map[string]interface{}  `json:"slug"`
            Variants map[string]interface{}  `json:"variants"`
            Trader map[string]interface{}  `json:"trader"`
            MultiSize map[string]interface{}  `json:"multi_size"`
            Description map[string]interface{}  `json:"description"`
            NoOfBoxes map[string]interface{}  `json:"no_of_boxes"`
            ItemType map[string]interface{}  `json:"item_type"`
            SizeGuide map[string]interface{}  `json:"size_guide"`
            Media map[string]interface{}  `json:"media"`
            ProductPublish map[string]interface{}  `json:"product_publish"`
            CountryOfOrigin map[string]interface{}  `json:"country_of_origin"`
            ProductGroupTag map[string]interface{}  `json:"product_group_tag"`
            Name map[string]interface{}  `json:"name"`
            BrandUID map[string]interface{}  `json:"brand_uid"`
            Command map[string]interface{}  `json:"command"`
            IsDependent map[string]interface{}  `json:"is_dependent"`
            HsnCode map[string]interface{}  `json:"hsn_code"`
            ReturnConfig map[string]interface{}  `json:"return_config"`
            CategorySlug map[string]interface{}  `json:"category_slug"`
            Highlights map[string]interface{}  `json:"highlights"`
            ItemCode map[string]interface{}  `json:"item_code"`
            Currency map[string]interface{}  `json:"currency"`
            Tags map[string]interface{}  `json:"tags"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            TraderType map[string]interface{}  `json:"trader_type"`
            ShortDescription map[string]interface{}  `json:"short_description"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            IsActive map[string]interface{}  `json:"is_active"`
         
    }
    
    // GlobalValidation used by Catalog
    type GlobalValidation struct {

        
            Required []string  `json:"required"`
            Properties Properties  `json:"properties"`
            Definitions map[string]interface{}  `json:"definitions"`
            Description string  `json:"description"`
            Title string  `json:"title"`
            Type string  `json:"type"`
         
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
    
    // UserInfo1 used by Catalog
    type UserInfo1 struct {

        
            Username string  `json:"username"`
            Email string  `json:"email"`
            UserID string  `json:"user_id"`
            UID string  `json:"uid"`
         
    }
    
    // ProductTemplateExportResponse used by Catalog
    type ProductTemplateExportResponse struct {

        
            TaskID string  `json:"task_id"`
            Type string  `json:"type"`
            NotificationEmails []string  `json:"notification_emails"`
            CreatedBy UserInfo1  `json:"created_by"`
            URL string  `json:"url"`
            CompletedOn string  `json:"completed_on"`
            SellerID float64  `json:"seller_id"`
            Status string  `json:"status"`
            Filters map[string]interface{}  `json:"filters"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // ProductDownloadsResponse used by Catalog
    type ProductDownloadsResponse struct {

        
            Items []ProductTemplateExportResponse  `json:"items"`
         
    }
    
    // ProductTemplateExportFilterRequest used by Catalog
    type ProductTemplateExportFilterRequest struct {

        
            FromDate string  `json:"from_date"`
            CatalogueTypes []string  `json:"catalogue_types"`
            ToDate string  `json:"to_date"`
            Brands []string  `json:"brands"`
            Templates []string  `json:"templates"`
         
    }
    
    // ProductTemplateDownloadsExport used by Catalog
    type ProductTemplateDownloadsExport struct {

        
            NotificationEmails []string  `json:"notification_emails"`
            Filters ProductTemplateExportFilterRequest  `json:"filters"`
            Type string  `json:"type"`
         
    }
    
    // ProductConfigurationDownloads used by Catalog
    type ProductConfigurationDownloads struct {

        
            Data []map[string]interface{}  `json:"data"`
            Multivalue bool  `json:"multivalue"`
         
    }
    
    // Media1 used by Catalog
    type Media1 struct {

        
            Portrait string  `json:"portrait"`
            Logo string  `json:"logo"`
            Landscape string  `json:"landscape"`
         
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
    
    // Hierarchy used by Catalog
    type Hierarchy struct {

        
            L2 float64  `json:"l2"`
            Department float64  `json:"department"`
            L1 float64  `json:"l1"`
         
    }
    
    // Category used by Catalog
    type Category struct {

        
            Slug string  `json:"slug"`
            Tryouts []string  `json:"tryouts"`
            Media Media1  `json:"media"`
            CreatedOn string  `json:"created_on"`
            Departments []float64  `json:"departments"`
            Marketplaces CategoryMapping  `json:"marketplaces"`
            Priority float64  `json:"priority"`
            ID string  `json:"id"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ModifiedOn string  `json:"modified_on"`
            UID float64  `json:"uid"`
            Level float64  `json:"level"`
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Hierarchy []Hierarchy  `json:"hierarchy"`
            Synonyms []string  `json:"synonyms"`
         
    }
    
    // CategoryResponse used by Catalog
    type CategoryResponse struct {

        
            Items []Category  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // CategoryRequestBody used by Catalog
    type CategoryRequestBody struct {

        
            Slug string  `json:"slug"`
            Tryouts []string  `json:"tryouts"`
            Media Media1  `json:"media"`
            Departments []float64  `json:"departments"`
            Marketplaces CategoryMapping  `json:"marketplaces"`
            Priority float64  `json:"priority"`
            Level float64  `json:"level"`
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            Hierarchy []Hierarchy  `json:"hierarchy"`
            Synonyms []string  `json:"synonyms"`
         
    }
    
    // CategoryCreateResponse used by Catalog
    type CategoryCreateResponse struct {

        
            Message string  `json:"message"`
            UID float64  `json:"uid"`
         
    }
    
    // CategoryUpdateResponse used by Catalog
    type CategoryUpdateResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // SingleCategoryResponse used by Catalog
    type SingleCategoryResponse struct {

        
            Data Category  `json:"data"`
         
    }
    
    // VerifiedBy used by Catalog
    type VerifiedBy struct {

        
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
         
    }
    
    // Trader used by Catalog
    type Trader struct {

        
            Name string  `json:"name"`
            Address []string  `json:"address"`
            Type string  `json:"type"`
         
    }
    
    // ProductPublish used by Catalog
    type ProductPublish struct {

        
            ProductOnlineDate string  `json:"product_online_date"`
            IsSet bool  `json:"is_set"`
         
    }
    
    // ReturnConfigResponse used by Catalog
    type ReturnConfigResponse struct {

        
            Returnable bool  `json:"returnable"`
            Time float64  `json:"time"`
            Unit string  `json:"unit"`
         
    }
    
    // Image used by Catalog
    type Image struct {

        
            SecureURL string  `json:"secure_url"`
            AspectRatioF float64  `json:"aspect_ratio_f"`
            AspectRatio string  `json:"aspect_ratio"`
            URL string  `json:"url"`
         
    }
    
    // NetQuantityResponse used by Catalog
    type NetQuantityResponse struct {

        
            Value float64  `json:"value"`
            Unit string  `json:"unit"`
         
    }
    
    // Logo used by Catalog
    type Logo struct {

        
            SecureURL string  `json:"secure_url"`
            AspectRatioF float64  `json:"aspect_ratio_f"`
            AspectRatio string  `json:"aspect_ratio"`
            URL string  `json:"url"`
         
    }
    
    // Brand used by Catalog
    type Brand struct {

        
            Logo Logo  `json:"logo"`
            Name string  `json:"name"`
            UID float64  `json:"uid"`
         
    }
    
    // ProductSchemaV2 used by Catalog
    type ProductSchemaV2 struct {

        
            Sizes []map[string]interface{}  `json:"sizes"`
            Slug string  `json:"slug"`
            VerifiedOn string  `json:"verified_on"`
            Variants map[string]interface{}  `json:"variants"`
            VerifiedBy VerifiedBy  `json:"verified_by"`
            Color string  `json:"color"`
            ID string  `json:"id"`
            IsExpirable bool  `json:"is_expirable"`
            Trader []Trader  `json:"trader"`
            MultiSize bool  `json:"multi_size"`
            Description string  `json:"description"`
            TemplateTag string  `json:"template_tag"`
            NoOfBoxes float64  `json:"no_of_boxes"`
            ItemType string  `json:"item_type"`
            SizeGuide string  `json:"size_guide"`
            Media []Media  `json:"media"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ProductPublish ProductPublish  `json:"product_publish"`
            VariantGroup map[string]interface{}  `json:"variant_group"`
            CountryOfOrigin string  `json:"country_of_origin"`
            ProductGroupTag []string  `json:"product_group_tag"`
            Name string  `json:"name"`
            BrandUID float64  `json:"brand_uid"`
            Category map[string]interface{}  `json:"category"`
            IsDependent bool  `json:"is_dependent"`
            HsnCode string  `json:"hsn_code"`
            CategoryUID float64  `json:"category_uid"`
            ImageNature string  `json:"image_nature"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            CreatedOn string  `json:"created_on"`
            IsImageLessProduct bool  `json:"is_image_less_product"`
            CompanyID float64  `json:"company_id"`
            ReturnConfig ReturnConfigResponse  `json:"return_config"`
            CategorySlug string  `json:"category_slug"`
            AllCompanyIds []float64  `json:"all_company_ids"`
            Highlights []string  `json:"highlights"`
            ItemCode string  `json:"item_code"`
            Currency string  `json:"currency"`
            VariantMedia map[string]interface{}  `json:"variant_media"`
            IsPhysical bool  `json:"is_physical"`
            Stage string  `json:"stage"`
            Moq map[string]interface{}  `json:"moq"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            IsActive bool  `json:"is_active"`
            AllIdentifiers []string  `json:"all_identifiers"`
            PrimaryColor string  `json:"primary_color"`
            Tags []string  `json:"tags"`
            AllSizes []map[string]interface{}  `json:"all_sizes"`
            Departments []float64  `json:"departments"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            Images []Image  `json:"images"`
            NetQuantity NetQuantityResponse  `json:"net_quantity"`
            ShortDescription string  `json:"short_description"`
            L3Mapping []string  `json:"l3_mapping"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            UID float64  `json:"uid"`
            Pending string  `json:"pending"`
            Brand Brand  `json:"brand"`
            IsSet bool  `json:"is_set"`
            Attributes map[string]interface{}  `json:"attributes"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // ProductListingResponseV2 used by Catalog
    type ProductListingResponseV2 struct {

        
            Items []ProductSchemaV2  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // ProductPublish1 used by Catalog
    type ProductPublish1 struct {

        
            ProductOnlineDate string  `json:"product_online_date"`
            IsSet bool  `json:"is_set"`
         
    }
    
    // TaxIdentifier used by Catalog
    type TaxIdentifier struct {

        
            ReportingHsn string  `json:"reporting_hsn"`
            HsnCode string  `json:"hsn_code"`
            HsnCodeID string  `json:"hsn_code_id"`
         
    }
    
    // ReturnConfig used by Catalog
    type ReturnConfig struct {

        
            Returnable bool  `json:"returnable"`
            Time float64  `json:"time"`
            Unit string  `json:"unit"`
         
    }
    
    // CustomOrder used by Catalog
    type CustomOrder struct {

        
            ManufacturingTimeUnit string  `json:"manufacturing_time_unit"`
            ManufacturingTime float64  `json:"manufacturing_time"`
            IsCustomOrder bool  `json:"is_custom_order"`
         
    }
    
    // NetQuantity used by Catalog
    type NetQuantity struct {

        
            Value float64  `json:"value"`
            Unit string  `json:"unit"`
         
    }
    
    // TeaserTag used by Catalog
    type TeaserTag struct {

        
            Tag string  `json:"tag"`
            URL string  `json:"url"`
         
    }
    
    // ProductCreateUpdateSchemaV2 used by Catalog
    type ProductCreateUpdateSchemaV2 struct {

        
            Sizes []map[string]interface{}  `json:"sizes"`
            Slug string  `json:"slug"`
            Variants map[string]interface{}  `json:"variants"`
            Trader []Trader  `json:"trader"`
            MultiSize bool  `json:"multi_size"`
            Description string  `json:"description"`
            TemplateTag string  `json:"template_tag"`
            NoOfBoxes float64  `json:"no_of_boxes"`
            ItemType string  `json:"item_type"`
            SizeGuide string  `json:"size_guide"`
            Media []Media  `json:"media"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ProductPublish ProductPublish1  `json:"product_publish"`
            VariantGroup map[string]interface{}  `json:"variant_group"`
            CountryOfOrigin string  `json:"country_of_origin"`
            BulkJobID string  `json:"bulk_job_id"`
            ProductGroupTag []string  `json:"product_group_tag"`
            Name string  `json:"name"`
            BrandUID float64  `json:"brand_uid"`
            IsDependent bool  `json:"is_dependent"`
            TaxIdentifier TaxIdentifier  `json:"tax_identifier"`
            IsImageLessProduct bool  `json:"is_image_less_product"`
            CompanyID float64  `json:"company_id"`
            ReturnConfig ReturnConfig  `json:"return_config"`
            CategorySlug string  `json:"category_slug"`
            Highlights []string  `json:"highlights"`
            ItemCode string  `json:"item_code"`
            Currency string  `json:"currency"`
            VariantMedia map[string]interface{}  `json:"variant_media"`
            Action string  `json:"action"`
            Tags []string  `json:"tags"`
            CustomOrder CustomOrder  `json:"custom_order"`
            Requester string  `json:"requester"`
            Departments []float64  `json:"departments"`
            NetQuantity NetQuantity  `json:"net_quantity"`
            ShortDescription string  `json:"short_description"`
            ChangeRequestID interface{}  `json:"change_request_id"`
            TeaserTag TeaserTag  `json:"teaser_tag"`
            UID float64  `json:"uid"`
            IsActive bool  `json:"is_active"`
            IsSet bool  `json:"is_set"`
            Attributes map[string]interface{}  `json:"attributes"`
         
    }
    
    // ProductVariants used by Catalog
    type ProductVariants struct {

        
            CategoryUID float64  `json:"category_uid"`
            Media []Media  `json:"media"`
            ItemCode string  `json:"item_code"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            BrandUID float64  `json:"brand_uid"`
         
    }
    
    // ProductVariantsResponse used by Catalog
    type ProductVariantsResponse struct {

        
            Variants []ProductVariants  `json:"variants"`
            Page Page  `json:"page"`
         
    }
    
    // AttributeMasterSerializer used by Catalog
    type AttributeMasterSerializer struct {

        
            Slug string  `json:"slug"`
            EnabledForEndConsumer bool  `json:"enabled_for_end_consumer"`
            Schema AttributeMaster  `json:"schema"`
            Description string  `json:"description"`
            Synonyms map[string]interface{}  `json:"synonyms"`
            Variant bool  `json:"variant"`
            Example string  `json:"example"`
            Unit string  `json:"unit"`
            Name string  `json:"name"`
            Filters AttributeMasterFilter  `json:"filters"`
            Suggestion string  `json:"suggestion"`
            Logo string  `json:"logo"`
            RawKey string  `json:"raw_key"`
            IsNested bool  `json:"is_nested"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Tags []string  `json:"tags"`
            Departments []string  `json:"departments"`
            Details AttributeMasterDetails  `json:"details"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
         
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

        
            Primary bool  `json:"primary"`
            GtinType string  `json:"gtin_type"`
            GtinValue string  `json:"gtin_value"`
         
    }
    
    // AllSizes used by Catalog
    type AllSizes struct {

        
            ItemWeightUnitOfMeasure string  `json:"item_weight_unit_of_measure"`
            ItemLength float64  `json:"item_length"`
            Identifiers []ValidateIdentifier  `json:"identifiers"`
            ItemDimensionsUnitOfMeasure string  `json:"item_dimensions_unit_of_measure"`
            ItemHeight float64  `json:"item_height"`
            ItemWidth float64  `json:"item_width"`
            Size string  `json:"size"`
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

        
            ProductOnlineDate float64  `json:"product_online_date"`
            IsSet bool  `json:"is_set"`
         
    }
    
    // Product used by Catalog
    type Product struct {

        
            Sizes []map[string]interface{}  `json:"sizes"`
            Slug string  `json:"slug"`
            VerifiedOn string  `json:"verified_on"`
            Variants map[string]interface{}  `json:"variants"`
            VerifiedBy VerifiedBy  `json:"verified_by"`
            Color string  `json:"color"`
            ID string  `json:"id"`
            IsExpirable bool  `json:"is_expirable"`
            Trader []Trader  `json:"trader"`
            MultiSize bool  `json:"multi_size"`
            Description string  `json:"description"`
            TemplateTag string  `json:"template_tag"`
            NoOfBoxes float64  `json:"no_of_boxes"`
            ItemType string  `json:"item_type"`
            SizeGuide string  `json:"size_guide"`
            Media []Media  `json:"media"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ProductPublish ProductPublished  `json:"product_publish"`
            VariantGroup map[string]interface{}  `json:"variant_group"`
            CountryOfOrigin string  `json:"country_of_origin"`
            ProductGroupTag []string  `json:"product_group_tag"`
            Name string  `json:"name"`
            BrandUID float64  `json:"brand_uid"`
            Category map[string]interface{}  `json:"category"`
            IsDependent bool  `json:"is_dependent"`
            HsnCode string  `json:"hsn_code"`
            CategoryUID float64  `json:"category_uid"`
            ImageNature string  `json:"image_nature"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            CreatedOn string  `json:"created_on"`
            IsImageLessProduct bool  `json:"is_image_less_product"`
            CompanyID float64  `json:"company_id"`
            ReturnConfig ReturnConfigResponse  `json:"return_config"`
            CategorySlug string  `json:"category_slug"`
            AllCompanyIds []float64  `json:"all_company_ids"`
            Highlights []string  `json:"highlights"`
            ItemCode string  `json:"item_code"`
            Currency string  `json:"currency"`
            VariantMedia map[string]interface{}  `json:"variant_media"`
            IsPhysical bool  `json:"is_physical"`
            Stage string  `json:"stage"`
            Moq map[string]interface{}  `json:"moq"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            IsActive bool  `json:"is_active"`
            AllIdentifiers []string  `json:"all_identifiers"`
            PrimaryColor string  `json:"primary_color"`
            Tags []string  `json:"tags"`
            AllSizes []map[string]interface{}  `json:"all_sizes"`
            Departments []float64  `json:"departments"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            Images []Image  `json:"images"`
            NetQuantity NetQuantityResponse  `json:"net_quantity"`
            ShortDescription string  `json:"short_description"`
            L3Mapping []string  `json:"l3_mapping"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            UID float64  `json:"uid"`
            Pending string  `json:"pending"`
            Brand Brand  `json:"brand"`
            IsSet bool  `json:"is_set"`
            Attributes map[string]interface{}  `json:"attributes"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // ProductListingResponse used by Catalog
    type ProductListingResponse struct {

        
            Items []Product  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // UserDetail1 used by Catalog
    type UserDetail1 struct {

        
            Username string  `json:"username"`
            FullName string  `json:"full_name"`
            UserID string  `json:"user_id"`
         
    }
    
    // ProductBulkRequest used by Catalog
    type ProductBulkRequest struct {

        
            Succeed float64  `json:"succeed"`
            CancelledRecords []string  `json:"cancelled_records"`
            CreatedOn string  `json:"created_on"`
            CompanyID float64  `json:"company_id"`
            CreatedBy UserDetail1  `json:"created_by"`
            FilePath string  `json:"file_path"`
            Total float64  `json:"total"`
            ModifiedOn string  `json:"modified_on"`
            Cancelled float64  `json:"cancelled"`
            Stage string  `json:"stage"`
            Template ProductTemplate  `json:"template"`
            IsActive bool  `json:"is_active"`
            FailedRecords []string  `json:"failed_records"`
            TemplateTag string  `json:"template_tag"`
            ModifiedBy UserDetail1  `json:"modified_by"`
            Failed float64  `json:"failed"`
         
    }
    
    // ProductBulkRequestList used by Catalog
    type ProductBulkRequestList struct {

        
            Items ProductBulkRequest  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // BulkJob used by Catalog
    type BulkJob struct {

        
            Succeed float64  `json:"succeed"`
            CancelledRecords []map[string]interface{}  `json:"cancelled_records"`
            CustomTemplateTag string  `json:"custom_template_tag"`
            CreatedOn string  `json:"created_on"`
            TrackingURL string  `json:"tracking_url"`
            CompanyID float64  `json:"company_id"`
            CreatedBy UserInfo1  `json:"created_by"`
            FilePath string  `json:"file_path"`
            Total float64  `json:"total"`
            Cancelled float64  `json:"cancelled"`
            Stage string  `json:"stage"`
            IsActive bool  `json:"is_active"`
            FailedRecords []map[string]interface{}  `json:"failed_records"`
            TemplateTag string  `json:"template_tag"`
            ModifiedBy UserInfo1  `json:"modified_by"`
            ModifiedOn string  `json:"modified_on"`
            Failed float64  `json:"failed"`
         
    }
    
    // BulkResponse used by Catalog
    type BulkResponse struct {

        
            CreatedOn string  `json:"created_on"`
            BatchID string  `json:"batch_id"`
            CreatedBy UserInfo1  `json:"created_by"`
            IsActive bool  `json:"is_active"`
            ModifiedBy UserInfo1  `json:"modified_by"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // BulkProductRequest used by Catalog
    type BulkProductRequest struct {

        
            CompanyID float64  `json:"company_id"`
            Data []map[string]interface{}  `json:"data"`
            TemplateTag string  `json:"template_tag"`
            BatchID string  `json:"batch_id"`
         
    }
    
    // ProductTagsViewResponse used by Catalog
    type ProductTagsViewResponse struct {

        
            Items []string  `json:"items"`
         
    }
    
    // UserCommon used by Catalog
    type UserCommon struct {

        
            CompanyID float64  `json:"company_id"`
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
         
    }
    
    // Items used by Catalog
    type Items struct {

        
            Succeed float64  `json:"succeed"`
            CancelledRecords []string  `json:"cancelled_records"`
            IsActive bool  `json:"is_active"`
            TrackingURL string  `json:"tracking_url"`
            CompanyID float64  `json:"company_id"`
            ID string  `json:"id"`
            CreatedBy UserCommon  `json:"created_by"`
            FilePath string  `json:"file_path"`
            Total float64  `json:"total"`
            Cancelled float64  `json:"cancelled"`
            Stage string  `json:"stage"`
            CreatedOn string  `json:"created_on"`
            FailedRecords []string  `json:"failed_records"`
            Retry float64  `json:"retry"`
            ModifiedBy UserCommon  `json:"modified_by"`
            ModifiedOn string  `json:"modified_on"`
            Failed float64  `json:"failed"`
         
    }
    
    // BulkAssetResponse used by Catalog
    type BulkAssetResponse struct {

        
            Items []Items  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // ProductBulkAssets used by Catalog
    type ProductBulkAssets struct {

        
            CompanyID float64  `json:"company_id"`
            BatchID string  `json:"batch_id"`
            URL string  `json:"url"`
            User map[string]interface{}  `json:"user"`
         
    }
    
    // ProductSizeDeleteDataResponse used by Catalog
    type ProductSizeDeleteDataResponse struct {

        
            CompanyID float64  `json:"company_id"`
            ItemID float64  `json:"item_id"`
            Size string  `json:"size"`
         
    }
    
    // ProductSizeDeleteResponse used by Catalog
    type ProductSizeDeleteResponse struct {

        
            Success bool  `json:"success"`
            Data ProductSizeDeleteDataResponse  `json:"data"`
         
    }
    
    // InventoryResponse used by Catalog
    type InventoryResponse struct {

        
            Quantity float64  `json:"quantity"`
            Price float64  `json:"price"`
            Identifiers map[string]interface{}  `json:"identifiers"`
            PriceTransfer float64  `json:"price_transfer"`
            InventoryUpdatedOn string  `json:"inventory_updated_on"`
            PriceEffective float64  `json:"price_effective"`
            SellerIdentifier string  `json:"seller_identifier"`
            Size string  `json:"size"`
            Store map[string]interface{}  `json:"store"`
            SellableQuantity float64  `json:"sellable_quantity"`
            Currency string  `json:"currency"`
            UID string  `json:"uid"`
            ItemID float64  `json:"item_id"`
         
    }
    
    // InventoryResponsePaginated used by Catalog
    type InventoryResponsePaginated struct {

        
            Items []InventoryResponse  `json:"items"`
            Page Page  `json:"page"`
         
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
    
    // GTIN used by Catalog
    type GTIN struct {

        
            Primary bool  `json:"primary"`
            GtinType string  `json:"gtin_type"`
            GtinValue string  `json:"gtin_value"`
         
    }
    
    // InvSize used by Catalog
    type InvSize struct {

        
            ItemWeightUnitOfMeasure string  `json:"item_weight_unit_of_measure"`
            ItemLength float64  `json:"item_length"`
            ExpirationDate string  `json:"expiration_date"`
            Set InventorySet  `json:"set"`
            Price float64  `json:"price"`
            ItemDimensionsUnitOfMeasure string  `json:"item_dimensions_unit_of_measure"`
            Identifiers []GTIN  `json:"identifiers"`
            Quantity float64  `json:"quantity"`
            ItemHeight float64  `json:"item_height"`
            PriceTransfer float64  `json:"price_transfer"`
            PriceEffective float64  `json:"price_effective"`
            ItemWidth float64  `json:"item_width"`
            Size string  `json:"size"`
            Currency string  `json:"currency"`
            ItemWeight float64  `json:"item_weight"`
            StoreCode string  `json:"store_code"`
            IsSet bool  `json:"is_set"`
         
    }
    
    // ItemQuery used by Catalog
    type ItemQuery struct {

        
            ItemCode string  `json:"item_code"`
            BrandUID float64  `json:"brand_uid"`
            UID float64  `json:"uid"`
         
    }
    
    // InventoryRequest used by Catalog
    type InventoryRequest struct {

        
            Sizes []InvSize  `json:"sizes"`
            CompanyID float64  `json:"company_id"`
            Item ItemQuery  `json:"item"`
         
    }
    
    // Seller used by Catalog
    type Seller struct {

        
            Count float64  `json:"count"`
            Name string  `json:"name"`
            UID float64  `json:"uid"`
         
    }
    
    // Store used by Catalog
    type Store struct {

        
            Count float64  `json:"count"`
            Name string  `json:"name"`
            UID float64  `json:"uid"`
         
    }
    
    // ArticleAssignment used by Catalog
    type ArticleAssignment struct {

        
            Level string  `json:"level"`
            Strategy string  `json:"strategy"`
         
    }
    
    // ProductSizePriceResponse used by Catalog
    type ProductSizePriceResponse struct {

        
            Quantity float64  `json:"quantity"`
            Seller Seller  `json:"seller"`
            Pincode float64  `json:"pincode"`
            Store Store  `json:"store"`
            ArticleID string  `json:"article_id"`
            ArticleAssignment ArticleAssignment  `json:"article_assignment"`
         
    }
    
    // ProductSizeSellerFilter used by Catalog
    type ProductSizeSellerFilter struct {

        
            Name string  `json:"name"`
            Value string  `json:"value"`
            IsSelected bool  `json:"is_selected"`
         
    }
    
    // ProductSizeSellersResponse used by Catalog
    type ProductSizeSellersResponse struct {

        
            Items []ProductSizePriceResponse  `json:"items"`
            Page Page  `json:"page"`
            SortOn []ProductSizeSellerFilter  `json:"sort_on"`
         
    }
    
    // ManufacturerResponse used by Catalog
    type ManufacturerResponse struct {

        
            IsDefault bool  `json:"is_default"`
            Name string  `json:"name"`
            Address string  `json:"address"`
         
    }
    
    // CompanyMeta used by Catalog
    type CompanyMeta struct {

        
            ID float64  `json:"id"`
         
    }
    
    // DimensionResponse used by Catalog
    type DimensionResponse struct {

        
            Width float64  `json:"width"`
            IsDefault bool  `json:"is_default"`
            Length float64  `json:"length"`
            Unit string  `json:"unit"`
            Height float64  `json:"height"`
         
    }
    
    // Trader1 used by Catalog
    type Trader1 struct {

        
            Name string  `json:"name"`
            Address []string  `json:"address"`
            Type string  `json:"type"`
         
    }
    
    // QuantityBase used by Catalog
    type QuantityBase struct {

        
            Count float64  `json:"count"`
            UpdatedAt string  `json:"updated_at"`
         
    }
    
    // Quantities used by Catalog
    type Quantities struct {

        
            OrderCommitted QuantityBase  `json:"order_committed"`
            Sellable QuantityBase  `json:"sellable"`
            NotAvailable QuantityBase  `json:"not_available"`
            Damaged QuantityBase  `json:"damaged"`
         
    }
    
    // BrandMeta used by Catalog
    type BrandMeta struct {

        
            Name string  `json:"name"`
            ID float64  `json:"id"`
         
    }
    
    // WeightResponse used by Catalog
    type WeightResponse struct {

        
            IsDefault bool  `json:"is_default"`
            Unit string  `json:"unit"`
            Shipping float64  `json:"shipping"`
         
    }
    
    // PriceMeta used by Catalog
    type PriceMeta struct {

        
            Marked float64  `json:"marked"`
            UpdatedAt string  `json:"updated_at"`
            Currency string  `json:"currency"`
            Effective float64  `json:"effective"`
            Transfer float64  `json:"transfer"`
            TpNotes map[string]interface{}  `json:"tp_notes"`
         
    }
    
    // ReturnConfig1 used by Catalog
    type ReturnConfig1 struct {

        
            Returnable bool  `json:"returnable"`
            Time float64  `json:"time"`
            Unit string  `json:"unit"`
         
    }
    
    // StoreMeta used by Catalog
    type StoreMeta struct {

        
            ID float64  `json:"id"`
         
    }
    
    // InventorySellerResponse used by Catalog
    type InventorySellerResponse struct {

        
            Manufacturer ManufacturerResponse  `json:"manufacturer"`
            Company CompanyMeta  `json:"company"`
            Dimension DimensionResponse  `json:"dimension"`
            Size string  `json:"size"`
            Identifier map[string]interface{}  `json:"identifier"`
            Trader []Trader1  `json:"trader"`
            TraceID string  `json:"trace_id"`
            Quantities Quantities  `json:"quantities"`
            Meta map[string]interface{}  `json:"meta"`
            FyndArticleCode string  `json:"fynd_article_code"`
            Brand BrandMeta  `json:"brand"`
            Weight WeightResponse  `json:"weight"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            CountryOfOrigin string  `json:"country_of_origin"`
            ItemID float64  `json:"item_id"`
            RawMeta map[string]interface{}  `json:"raw_meta"`
            FyndItemCode string  `json:"fynd_item_code"`
            Fragile bool  `json:"fragile"`
            Price PriceMeta  `json:"price"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            TotalQuantity float64  `json:"total_quantity"`
            ReturnConfig ReturnConfig1  `json:"return_config"`
            SellerIdentifier string  `json:"seller_identifier"`
            Store StoreMeta  `json:"store"`
            Stage string  `json:"stage"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            ExpirationDate string  `json:"expiration_date"`
            Set InventorySet  `json:"set"`
            Tags []string  `json:"tags"`
            TrackInventory bool  `json:"track_inventory"`
            FyndMeta map[string]interface{}  `json:"fynd_meta"`
            AddedOnStore string  `json:"added_on_store"`
            CreatedBy UserSerializer  `json:"created_by"`
            UID string  `json:"uid"`
            IsActive bool  `json:"is_active"`
            IsSet bool  `json:"is_set"`
         
    }
    
    // InventorySellerIdentifierResponsePaginated used by Catalog
    type InventorySellerIdentifierResponsePaginated struct {

        
            Items []InventorySellerResponse  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // ManufacturerResponse1 used by Catalog
    type ManufacturerResponse1 struct {

        
            IsDefault bool  `json:"is_default"`
            Name string  `json:"name"`
            Address string  `json:"address"`
         
    }
    
    // CompanyMeta1 used by Catalog
    type CompanyMeta1 struct {

        
            ID float64  `json:"id"`
         
    }
    
    // DimensionResponse1 used by Catalog
    type DimensionResponse1 struct {

        
            Length float64  `json:"length"`
            Width float64  `json:"width"`
            Height float64  `json:"height"`
            Unit string  `json:"unit"`
         
    }
    
    // Trader2 used by Catalog
    type Trader2 struct {

        
            Name string  `json:"name"`
            Address []string  `json:"address"`
            Type string  `json:"type"`
         
    }
    
    // Quantity used by Catalog
    type Quantity struct {

        
            Count float64  `json:"count"`
         
    }
    
    // QuantitiesArticle used by Catalog
    type QuantitiesArticle struct {

        
            OrderCommitted Quantity  `json:"order_committed"`
            Sellable Quantity  `json:"sellable"`
            NotAvailable Quantity  `json:"not_available"`
            Damaged Quantity  `json:"damaged"`
         
    }
    
    // WeightResponse1 used by Catalog
    type WeightResponse1 struct {

        
            Unit string  `json:"unit"`
            Shipping float64  `json:"shipping"`
         
    }
    
    // DateMeta used by Catalog
    type DateMeta struct {

        
            InventoryUpdatedOn string  `json:"inventory_updated_on"`
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
            AddedOnStore string  `json:"added_on_store"`
         
    }
    
    // PriceArticle used by Catalog
    type PriceArticle struct {

        
            Marked float64  `json:"marked"`
            Currency string  `json:"currency"`
            Effective float64  `json:"effective"`
            Transfer float64  `json:"transfer"`
            TpNotes map[string]interface{}  `json:"tp_notes"`
         
    }
    
    // ReturnConfig2 used by Catalog
    type ReturnConfig2 struct {

        
            Returnable bool  `json:"returnable"`
            Time float64  `json:"time"`
            Unit string  `json:"unit"`
         
    }
    
    // ArticleStoreResponse used by Catalog
    type ArticleStoreResponse struct {

        
            StoreCode string  `json:"store_code"`
            Name string  `json:"name"`
            StoreType string  `json:"store_type"`
            UID float64  `json:"uid"`
         
    }
    
    // BrandMeta1 used by Catalog
    type BrandMeta1 struct {

        
            Name string  `json:"name"`
            ID float64  `json:"id"`
         
    }
    
    // GetInventories used by Catalog
    type GetInventories struct {

        
            Manufacturer ManufacturerResponse1  `json:"manufacturer"`
            Company CompanyMeta1  `json:"company"`
            Platforms map[string]interface{}  `json:"platforms"`
            Dimension DimensionResponse1  `json:"dimension"`
            ID string  `json:"id"`
            Size string  `json:"size"`
            Identifier map[string]interface{}  `json:"identifier"`
            TraceID string  `json:"trace_id"`
            Trader []Trader2  `json:"trader"`
            Quantities QuantitiesArticle  `json:"quantities"`
            Weight WeightResponse1  `json:"weight"`
            CountryOfOrigin string  `json:"country_of_origin"`
            DateMeta DateMeta  `json:"date_meta"`
            ItemID float64  `json:"item_id"`
            Price PriceArticle  `json:"price"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            TotalQuantity float64  `json:"total_quantity"`
            ReturnConfig ReturnConfig2  `json:"return_config"`
            SellerIdentifier string  `json:"seller_identifier"`
            Store ArticleStoreResponse  `json:"store"`
            Stage string  `json:"stage"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            ExpirationDate string  `json:"expiration_date"`
            Tags []string  `json:"tags"`
            TrackInventory bool  `json:"track_inventory"`
            InventoryUpdatedOn string  `json:"inventory_updated_on"`
            CreatedBy UserSerializer  `json:"created_by"`
            UID string  `json:"uid"`
            Brand BrandMeta1  `json:"brand"`
            IsSet bool  `json:"is_set"`
         
    }
    
    // GetInventoriesResponse used by Catalog
    type GetInventoriesResponse struct {

        
            Items []GetInventories  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // BulkInventoryGetItems used by Catalog
    type BulkInventoryGetItems struct {

        
            Succeed float64  `json:"succeed"`
            CancelledRecords []string  `json:"cancelled_records"`
            IsActive bool  `json:"is_active"`
            CompanyID float64  `json:"company_id"`
            ID string  `json:"id"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            FilePath string  `json:"file_path"`
            Total float64  `json:"total"`
            ModifiedOn string  `json:"modified_on"`
            Cancelled float64  `json:"cancelled"`
            Stage string  `json:"stage"`
            CreatedOn string  `json:"created_on"`
            FailedRecords []string  `json:"failed_records"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Failed float64  `json:"failed"`
         
    }
    
    // BulkInventoryGet used by Catalog
    type BulkInventoryGet struct {

        
            Items []BulkInventoryGetItems  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // InventoryJobPayload used by Catalog
    type InventoryJobPayload struct {

        
            ExpirationDate string  `json:"expiration_date"`
            ItemWeightUnitOfMeasure string  `json:"item_weight_unit_of_measure"`
            Quantity float64  `json:"quantity"`
            Price float64  `json:"price"`
            ItemDimensionsUnitOfMeasure string  `json:"item_dimensions_unit_of_measure"`
            Tags []string  `json:"tags"`
            PriceMarked float64  `json:"price_marked"`
            PriceEffective float64  `json:"price_effective"`
            TotalQuantity float64  `json:"total_quantity"`
            SellerIdentifier string  `json:"seller_identifier"`
            TraceID string  `json:"trace_id"`
            Currency string  `json:"currency"`
            StoreCode string  `json:"store_code"`
         
    }
    
    // InventoryBulkRequest used by Catalog
    type InventoryBulkRequest struct {

        
            Sizes []InventoryJobPayload  `json:"sizes"`
            CompanyID float64  `json:"company_id"`
            User map[string]interface{}  `json:"user"`
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

        
            Quantity InventoryExportQuantityFilter  `json:"quantity"`
            FromDate string  `json:"from_date"`
            StoreIds []float64  `json:"store_ids"`
            BrandIds []float64  `json:"brand_ids"`
            ToDate string  `json:"to_date"`
         
    }
    
    // InventoryExportJob used by Catalog
    type InventoryExportJob struct {

        
            TaskID string  `json:"task_id"`
            NotificationEmails []string  `json:"notification_emails"`
            CompletedOn string  `json:"completed_on"`
            URL string  `json:"url"`
            Type string  `json:"type"`
            SellerID float64  `json:"seller_id"`
            Status string  `json:"status"`
            Filters InventoryExportAdvanceOption  `json:"filters"`
         
    }
    
    // InventoryExportRequest used by Catalog
    type InventoryExportRequest struct {

        
            Brand []float64  `json:"brand"`
            Store []float64  `json:"store"`
            Type string  `json:"type"`
         
    }
    
    // InventoryExportResponse used by Catalog
    type InventoryExportResponse struct {

        
            TaskID string  `json:"task_id"`
            NotificationEmails []string  `json:"notification_emails"`
            CreatedBy string  `json:"created_by"`
            Type string  `json:"type"`
            SellerID float64  `json:"seller_id"`
            CreatedOn string  `json:"created_on"`
            Status string  `json:"status"`
            Filters map[string]interface{}  `json:"filters"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // InventoryJobFilters used by Catalog
    type InventoryJobFilters struct {

        
            Quantity InventoryExportQuantityFilter  `json:"quantity"`
            FromDate string  `json:"from_date"`
            Stores []string  `json:"stores"`
            ToDate string  `json:"to_date"`
            Brands []string  `json:"brands"`
         
    }
    
    // InventoryJobDetailResponse used by Catalog
    type InventoryJobDetailResponse struct {

        
            TaskID string  `json:"task_id"`
            Type string  `json:"type"`
            NotificationEmails []string  `json:"notification_emails"`
            ID string  `json:"id"`
            CreatedBy UserDetail  `json:"created_by"`
            URL string  `json:"url"`
            CancelledBy UserDetail  `json:"cancelled_by"`
            CancelledOn string  `json:"cancelled_on"`
            CompletedOn string  `json:"completed_on"`
            SellerID float64  `json:"seller_id"`
            CreatedOn string  `json:"created_on"`
            Status string  `json:"status"`
            Filters InventoryJobFilters  `json:"filters"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // InventoryExportJobListResponse used by Catalog
    type InventoryExportJobListResponse struct {

        
            Items InventoryJobDetailResponse  `json:"items"`
         
    }
    
    // InventoryExportFilter used by Catalog
    type InventoryExportFilter struct {

        
            Quantity InventoryExportQuantityFilter  `json:"quantity"`
            FromDate string  `json:"from_date"`
            StoreIds []float64  `json:"store_ids"`
            BrandIds []float64  `json:"brand_ids"`
            ToDate string  `json:"to_date"`
         
    }
    
    // InventoryCreateRequest used by Catalog
    type InventoryCreateRequest struct {

        
            Data []string  `json:"data"`
            NotificationEmails []string  `json:"notification_emails"`
            Filters InventoryExportFilter  `json:"filters"`
            Type string  `json:"type"`
         
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
            StoreID float64  `json:"store_id"`
            Tags []string  `json:"tags"`
            PriceEffective float64  `json:"price_effective"`
            TotalQuantity float64  `json:"total_quantity"`
            SellerIdentifier string  `json:"seller_identifier"`
            TraceID string  `json:"trace_id"`
            PriceMarked float64  `json:"price_marked"`
         
    }
    
    // InventoryResponseItem used by Catalog
    type InventoryResponseItem struct {

        
            Reason InventoryFailedReason  `json:"reason"`
            Data InventoryPayload  `json:"data"`
         
    }
    
    // InventoryUpdateResponse used by Catalog
    type InventoryUpdateResponse struct {

        
            Message string  `json:"message"`
            Items []InventoryResponseItem  `json:"items"`
         
    }
    
    // InventoryRequestSchemaV2 used by Catalog
    type InventoryRequestSchemaV2 struct {

        
            Meta map[string]interface{}  `json:"meta"`
            CompanyID float64  `json:"company_id"`
            Payload []InventoryPayload  `json:"payload"`
         
    }
    
    // HsnUpsert used by Catalog
    type HsnUpsert struct {

        
            HsnCode string  `json:"hsn_code"`
            Hs2Code string  `json:"hs2_code"`
            CompanyID float64  `json:"company_id"`
            Tax1 float64  `json:"tax1"`
            Threshold2 float64  `json:"threshold2"`
            TaxOnMrp bool  `json:"tax_on_mrp"`
            Threshold1 float64  `json:"threshold1"`
            UID float64  `json:"uid"`
            Tax2 float64  `json:"tax2"`
            TaxOnEsp bool  `json:"tax_on_esp"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // HsnCodesObject used by Catalog
    type HsnCodesObject struct {

        
            HsnCode string  `json:"hsn_code"`
            Hs2Code string  `json:"hs2_code"`
            CompanyID float64  `json:"company_id"`
            Tax1 float64  `json:"tax1"`
            Threshold2 float64  `json:"threshold2"`
            ID string  `json:"id"`
            TaxOnMrp bool  `json:"tax_on_mrp"`
            Threshold1 float64  `json:"threshold1"`
            Tax2 float64  `json:"tax2"`
            TaxOnEsp bool  `json:"tax_on_esp"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // HsnCode used by Catalog
    type HsnCode struct {

        
            Data HsnCodesObject  `json:"data"`
         
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

        
            HsnCode string  `json:"hsn_code"`
            Description string  `json:"description"`
            HsnCodeID string  `json:"hsn_code_id"`
            ReportingHsn string  `json:"reporting_hsn"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Type string  `json:"type"`
            Taxes []TaxSlab  `json:"taxes"`
            CreatedOn string  `json:"created_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            ModifiedOn string  `json:"modified_on"`
            CountryCode string  `json:"country_code"`
         
    }
    
    // PageResponse used by Catalog
    type PageResponse struct {

        
            ItemTotal float64  `json:"item_total"`
            HasPrevious bool  `json:"has_previous"`
            Current string  `json:"current"`
            Size float64  `json:"size"`
            HasNext bool  `json:"has_next"`
         
    }
    
    // HsnCodesListingResponseSchemaV2 used by Catalog
    type HsnCodesListingResponseSchemaV2 struct {

        
            Items []HSNDataInsertV2  `json:"items"`
            Page PageResponse  `json:"page"`
         
    }
    
    // Media2 used by Catalog
    type Media2 struct {

        
            URL string  `json:"url"`
            AspectRatio string  `json:"aspect_ratio"`
            Type string  `json:"type"`
         
    }
    
    // BrandItem used by Catalog
    type BrandItem struct {

        
            Slug string  `json:"slug"`
            Departments []string  `json:"departments"`
            Logo Media2  `json:"logo"`
            Discount string  `json:"discount"`
            UID float64  `json:"uid"`
            Banners ImageUrls  `json:"banners"`
            Action Action  `json:"action"`
            Name string  `json:"name"`
         
    }
    
    // BrandListingResponse used by Catalog
    type BrandListingResponse struct {

        
            Items []BrandItem  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // Department used by Catalog
    type Department struct {

        
            PriorityOrder float64  `json:"priority_order"`
            Slug string  `json:"slug"`
            Logo Media2  `json:"logo"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // DepartmentResponse used by Catalog
    type DepartmentResponse struct {

        
            Items []Department  `json:"items"`
         
    }
    
    // DepartmentIdentifier used by Catalog
    type DepartmentIdentifier struct {

        
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
         
    }
    
    // ThirdLevelChild used by Catalog
    type ThirdLevelChild struct {

        
            Slug string  `json:"slug"`
            Childs []map[string]interface{}  `json:"childs"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            UID float64  `json:"uid"`
            Banners ImageUrls  `json:"banners"`
            Action Action  `json:"action"`
            Name string  `json:"name"`
         
    }
    
    // SecondLevelChild used by Catalog
    type SecondLevelChild struct {

        
            Slug string  `json:"slug"`
            Childs []ThirdLevelChild  `json:"childs"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            UID float64  `json:"uid"`
            Banners ImageUrls  `json:"banners"`
            Action Action  `json:"action"`
            Name string  `json:"name"`
         
    }
    
    // Child used by Catalog
    type Child struct {

        
            Slug string  `json:"slug"`
            Childs []SecondLevelChild  `json:"childs"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            UID float64  `json:"uid"`
            Banners ImageUrls  `json:"banners"`
            Action Action  `json:"action"`
            Name string  `json:"name"`
         
    }
    
    // CategoryItems used by Catalog
    type CategoryItems struct {

        
            Slug string  `json:"slug"`
            Childs []Child  `json:"childs"`
            UID float64  `json:"uid"`
            Banners ImageUrls  `json:"banners"`
            Action Action  `json:"action"`
            Name string  `json:"name"`
         
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

        
            Operators map[string]interface{}  `json:"operators"`
            Items []ProductListingDetail  `json:"items"`
            Page Page  `json:"page"`
            Filters []ProductFilters  `json:"filters"`
            SortOn []ProductSortOn  `json:"sort_on"`
         
    }
    
    // ProductDetail used by Catalog
    type ProductDetail struct {

        
            Slug string  `json:"slug"`
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
            Color string  `json:"color"`
            Description string  `json:"description"`
            ItemType string  `json:"item_type"`
            Name string  `json:"name"`
            PromoMeta map[string]interface{}  `json:"promo_meta"`
            ImageNature string  `json:"image_nature"`
            Highlights []string  `json:"highlights"`
            HasVariant bool  `json:"has_variant"`
            ItemCode string  `json:"item_code"`
            Type string  `json:"type"`
            ProductOnlineDate string  `json:"product_online_date"`
            Similars []string  `json:"similars"`
            Rating float64  `json:"rating"`
            Attributes map[string]interface{}  `json:"attributes"`
            Medias []Media  `json:"medias"`
            ShortDescription string  `json:"short_description"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            UID float64  `json:"uid"`
            Brand ProductBrand  `json:"brand"`
            RatingCount float64  `json:"rating_count"`
            Tryouts []string  `json:"tryouts"`
         
    }
    
    // InventoryPage used by Catalog
    type InventoryPage struct {

        
            NextID string  `json:"next_id"`
            ItemTotal float64  `json:"item_total"`
            HasPrevious bool  `json:"has_previous"`
            HasNext bool  `json:"has_next"`
            Type string  `json:"type"`
         
    }
    
    // InventoryStockResponse used by Catalog
    type InventoryStockResponse struct {

        
            Items []map[string]interface{}  `json:"items"`
            Page InventoryPage  `json:"page"`
         
    }
    
    // ArticleQuery used by Catalog
    type ArticleQuery struct {

        
            IgnoredStores []float64  `json:"ignored_stores"`
            ItemID float64  `json:"item_id"`
            Size string  `json:"size"`
         
    }
    
    // ArticleAssignment1 used by Catalog
    type ArticleAssignment1 struct {

        
            Level string  `json:"level"`
            Strategy string  `json:"strategy"`
         
    }
    
    // AssignStoreArticle used by Catalog
    type AssignStoreArticle struct {

        
            Quantity float64  `json:"quantity"`
            Query ArticleQuery  `json:"query"`
            Meta map[string]interface{}  `json:"meta"`
            GroupID string  `json:"group_id"`
            ArticleAssignment ArticleAssignment1  `json:"article_assignment"`
         
    }
    
    // AssignStore used by Catalog
    type AssignStore struct {

        
            ChannelIdentifier string  `json:"channel_identifier"`
            Articles []AssignStoreArticle  `json:"articles"`
            StoreIds []float64  `json:"store_ids"`
            Pincode string  `json:"pincode"`
            CompanyID float64  `json:"company_id"`
            ChannelType string  `json:"channel_type"`
            AppID string  `json:"app_id"`
         
    }
    
    // ArticleAssignment2 used by Catalog
    type ArticleAssignment2 struct {

        
            Level string  `json:"level"`
            Strategy string  `json:"strategy"`
         
    }
    
    // StoreAssignResponse used by Catalog
    type StoreAssignResponse struct {

        
            Quantity float64  `json:"quantity"`
            StoreID float64  `json:"store_id"`
            SCity string  `json:"s_city"`
            StrategyWiseListing []map[string]interface{}  `json:"strategy_wise_listing"`
            CompanyID float64  `json:"company_id"`
            ItemID float64  `json:"item_id"`
            PriceEffective float64  `json:"price_effective"`
            Size string  `json:"size"`
            Index float64  `json:"index"`
            StorePincode float64  `json:"store_pincode"`
            GroupID string  `json:"group_id"`
            ID string  `json:"_id"`
            UID string  `json:"uid"`
            Meta map[string]interface{}  `json:"meta"`
            PriceMarked float64  `json:"price_marked"`
            Status bool  `json:"status"`
            ArticleAssignment ArticleAssignment2  `json:"article_assignment"`
         
    }
    
    // UserSerializer1 used by Catalog
    type UserSerializer1 struct {

        
            Contact string  `json:"contact"`
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
         
    }
    
    // UserSerializer2 used by Catalog
    type UserSerializer2 struct {

        
            Contact string  `json:"contact"`
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
         
    }
    
    // GetAddressSerializer used by Catalog
    type GetAddressSerializer struct {

        
            Longitude float64  `json:"longitude"`
            Address1 string  `json:"address1"`
            AddressType string  `json:"address_type"`
            Pincode float64  `json:"pincode"`
            Country string  `json:"country"`
            State string  `json:"state"`
            Landmark string  `json:"landmark"`
            Address2 string  `json:"address2"`
            CountryCode string  `json:"country_code"`
            City string  `json:"city"`
            Latitude float64  `json:"latitude"`
         
    }
    
    // GetCompanySerializer used by Catalog
    type GetCompanySerializer struct {

        
            CompanyType string  `json:"company_type"`
            VerifiedOn string  `json:"verified_on"`
            VerifiedBy UserSerializer2  `json:"verified_by"`
            BusinessType string  `json:"business_type"`
            CreatedBy UserSerializer2  `json:"created_by"`
            RejectReason string  `json:"reject_reason"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            UID float64  `json:"uid"`
            Stage string  `json:"stage"`
            CreatedOn string  `json:"created_on"`
            Name string  `json:"name"`
            ModifiedBy UserSerializer2  `json:"modified_by"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // SellerPhoneNumber used by Catalog
    type SellerPhoneNumber struct {

        
            Number string  `json:"number"`
            CountryCode float64  `json:"country_code"`
         
    }
    
    // ProductReturnConfigSerializer used by Catalog
    type ProductReturnConfigSerializer struct {

        
            StoreUID float64  `json:"store_uid"`
            OnSameStore bool  `json:"on_same_store"`
         
    }
    
    // LocationIntegrationType used by Catalog
    type LocationIntegrationType struct {

        
            Order string  `json:"order"`
            Inventory string  `json:"inventory"`
         
    }
    
    // LocationTimingSerializer used by Catalog
    type LocationTimingSerializer struct {

        
            Hour float64  `json:"hour"`
            Minute float64  `json:"minute"`
         
    }
    
    // LocationDayWiseSerializer used by Catalog
    type LocationDayWiseSerializer struct {

        
            Open bool  `json:"open"`
            Opening LocationTimingSerializer  `json:"opening"`
            Weekday string  `json:"weekday"`
            Closing LocationTimingSerializer  `json:"closing"`
         
    }
    
    // LocationManagerSerializer used by Catalog
    type LocationManagerSerializer struct {

        
            Email string  `json:"email"`
            Name string  `json:"name"`
            MobileNo SellerPhoneNumber  `json:"mobile_no"`
         
    }
    
    // InvoiceCredSerializer used by Catalog
    type InvoiceCredSerializer struct {

        
            Username string  `json:"username"`
            Password string  `json:"password"`
            Enabled bool  `json:"enabled"`
         
    }
    
    // InvoiceDetailsSerializer used by Catalog
    type InvoiceDetailsSerializer struct {

        
            EWaybill InvoiceCredSerializer  `json:"e_waybill"`
            EInvoice InvoiceCredSerializer  `json:"e_invoice"`
         
    }
    
    // Document used by Catalog
    type Document struct {

        
            Value string  `json:"value"`
            URL string  `json:"url"`
            Type string  `json:"type"`
            Verified bool  `json:"verified"`
            LegalName string  `json:"legal_name"`
         
    }
    
    // GetLocationSerializer used by Catalog
    type GetLocationSerializer struct {

        
            VerifiedOn string  `json:"verified_on"`
            VerifiedBy UserSerializer1  `json:"verified_by"`
            Company GetCompanySerializer  `json:"company"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            DisplayName string  `json:"display_name"`
            ProductReturnConfig ProductReturnConfigSerializer  `json:"product_return_config"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            NotificationEmails []string  `json:"notification_emails"`
            IntegrationType LocationIntegrationType  `json:"integration_type"`
            Name string  `json:"name"`
            PhoneNumber string  `json:"phone_number"`
            StoreType string  `json:"store_type"`
            Timing []LocationDayWiseSerializer  `json:"timing"`
            Warnings map[string]interface{}  `json:"warnings"`
            Manager LocationManagerSerializer  `json:"manager"`
            Stage string  `json:"stage"`
            GstCredentials InvoiceDetailsSerializer  `json:"gst_credentials"`
            ModifiedBy UserSerializer1  `json:"modified_by"`
            Code string  `json:"code"`
            CreatedBy UserSerializer1  `json:"created_by"`
            UID float64  `json:"uid"`
            CreatedOn string  `json:"created_on"`
            Address GetAddressSerializer  `json:"address"`
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
            Logo string  `json:"logo"`
            UID float64  `json:"uid"`
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            AppID string  `json:"app_id"`
         
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
    

    
    // BusinessCountryInfo used by CompanyProfile
    type BusinessCountryInfo struct {

        
            Country string  `json:"country"`
            CountryCode string  `json:"country_code"`
         
    }
    
    // Website used by CompanyProfile
    type Website struct {

        
            URL string  `json:"url"`
         
    }
    
    // BusinessDetails used by CompanyProfile
    type BusinessDetails struct {

        
            Website Website  `json:"website"`
         
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
    
    // GetAddressSerializer used by CompanyProfile
    type GetAddressSerializer struct {

        
            CountryCode string  `json:"country_code"`
            Longitude float64  `json:"longitude"`
            AddressType string  `json:"address_type"`
            City string  `json:"city"`
            Country string  `json:"country"`
            Latitude float64  `json:"latitude"`
            Pincode float64  `json:"pincode"`
            Landmark string  `json:"landmark"`
            State string  `json:"state"`
            Address2 string  `json:"address2"`
            Address1 string  `json:"address1"`
         
    }
    
    // UserSerializer used by CompanyProfile
    type UserSerializer struct {

        
            UserID string  `json:"user_id"`
            Contact string  `json:"contact"`
            Username string  `json:"username"`
         
    }
    
    // CompanyTaxesSerializer used by CompanyProfile
    type CompanyTaxesSerializer struct {

        
            Enable bool  `json:"enable"`
            EffectiveDate string  `json:"effective_date"`
            Rate float64  `json:"rate"`
         
    }
    
    // Document used by CompanyProfile
    type Document struct {

        
            Value string  `json:"value"`
            LegalName string  `json:"legal_name"`
            Verified bool  `json:"verified"`
            URL string  `json:"url"`
            Type string  `json:"type"`
         
    }
    
    // GetCompanyProfileSerializerResponse used by CompanyProfile
    type GetCompanyProfileSerializerResponse struct {

        
            Mode string  `json:"mode"`
            Name string  `json:"name"`
            BusinessCountryInfo BusinessCountryInfo  `json:"business_country_info"`
            BusinessDetails BusinessDetails  `json:"business_details"`
            ContactDetails ContactDetails  `json:"contact_details"`
            NotificationEmails []string  `json:"notification_emails"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            FranchiseEnabled bool  `json:"franchise_enabled"`
            BusinessInfo string  `json:"business_info"`
            CompanyType string  `json:"company_type"`
            UID float64  `json:"uid"`
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            CreatedBy UserSerializer  `json:"created_by"`
            VerifiedOn string  `json:"verified_on"`
            Taxes []CompanyTaxesSerializer  `json:"taxes"`
            Documents []Document  `json:"documents"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            Warnings map[string]interface{}  `json:"warnings"`
            BusinessType string  `json:"business_type"`
            Stage string  `json:"stage"`
         
    }
    
    // ErrorResponse used by CompanyProfile
    type ErrorResponse struct {

        
            Message string  `json:"message"`
            Meta map[string]interface{}  `json:"meta"`
            Code string  `json:"code"`
            Status float64  `json:"status"`
         
    }
    
    // CompanyTaxesSerializer1 used by CompanyProfile
    type CompanyTaxesSerializer1 struct {

        
            Enable bool  `json:"enable"`
            EffectiveDate string  `json:"effective_date"`
            Rate float64  `json:"rate"`
         
    }
    
    // CreateUpdateAddressSerializer used by CompanyProfile
    type CreateUpdateAddressSerializer struct {

        
            CountryCode string  `json:"country_code"`
            Longitude float64  `json:"longitude"`
            City string  `json:"city"`
            AddressType string  `json:"address_type"`
            Country string  `json:"country"`
            Latitude float64  `json:"latitude"`
            Pincode float64  `json:"pincode"`
            Landmark string  `json:"landmark"`
            State string  `json:"state"`
            Address2 string  `json:"address2"`
            Address1 string  `json:"address1"`
         
    }
    
    // UpdateCompany used by CompanyProfile
    type UpdateCompany struct {

        
            Taxes []CompanyTaxesSerializer1  `json:"taxes"`
            NotificationEmails []string  `json:"notification_emails"`
            Name string  `json:"name"`
            Addresses []CreateUpdateAddressSerializer  `json:"addresses"`
            FranchiseEnabled bool  `json:"franchise_enabled"`
            Documents []Document  `json:"documents"`
            RejectReason string  `json:"reject_reason"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Slug string  `json:"slug"`
            BusinessInfo string  `json:"business_info"`
            Warnings map[string]interface{}  `json:"warnings"`
            BusinessType string  `json:"business_type"`
            BusinessDetails BusinessDetails  `json:"business_details"`
            CompanyType string  `json:"company_type"`
            ContactDetails ContactDetails  `json:"contact_details"`
         
    }
    
    // ProfileSuccessResponse used by CompanyProfile
    type ProfileSuccessResponse struct {

        
            Message string  `json:"message"`
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

        
            Brand DocumentsObj  `json:"brand"`
            CompanyDocuments DocumentsObj  `json:"company_documents"`
            Product DocumentsObj  `json:"product"`
            StoreDocuments DocumentsObj  `json:"store_documents"`
            Stage string  `json:"stage"`
            Store DocumentsObj  `json:"store"`
            UID float64  `json:"uid"`
         
    }
    
    // BrandBannerSerializer used by CompanyProfile
    type BrandBannerSerializer struct {

        
            Portrait string  `json:"portrait"`
            Landscape string  `json:"landscape"`
         
    }
    
    // GetBrandResponseSerializer used by CompanyProfile
    type GetBrandResponseSerializer struct {

        
            Mode string  `json:"mode"`
            Name string  `json:"name"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            Logo string  `json:"logo"`
            UID float64  `json:"uid"`
            SlugKey string  `json:"slug_key"`
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            CreatedBy UserSerializer  `json:"created_by"`
            VerifiedOn string  `json:"verified_on"`
            Description string  `json:"description"`
            Synonyms []string  `json:"synonyms"`
            Banner BrandBannerSerializer  `json:"banner"`
            RejectReason string  `json:"reject_reason"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            Warnings map[string]interface{}  `json:"warnings"`
            Stage string  `json:"stage"`
         
    }
    
    // CreateUpdateBrandRequestSerializer used by CompanyProfile
    type CreateUpdateBrandRequestSerializer struct {

        
            BrandTier string  `json:"brand_tier"`
            Name string  `json:"name"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            CompanyID float64  `json:"company_id"`
            Logo string  `json:"logo"`
            Description string  `json:"description"`
            Synonyms []string  `json:"synonyms"`
            Banner BrandBannerSerializer  `json:"banner"`
            UID float64  `json:"uid"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
         
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

        
            ModifiedOn string  `json:"modified_on"`
            NotificationEmails []string  `json:"notification_emails"`
            Name string  `json:"name"`
            CreatedOn string  `json:"created_on"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            RejectReason string  `json:"reject_reason"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            CreatedBy UserSerializer  `json:"created_by"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            BusinessCountryInfo BusinessCountryInfo  `json:"business_country_info"`
            VerifiedOn string  `json:"verified_on"`
            Details CompanyDetails  `json:"details"`
            BusinessType string  `json:"business_type"`
            Stage string  `json:"stage"`
            CompanyType string  `json:"company_type"`
            UID float64  `json:"uid"`
            MarketChannels []string  `json:"market_channels"`
         
    }
    
    // CompanyBrandSerializer used by CompanyProfile
    type CompanyBrandSerializer struct {

        
            Brand GetBrandResponseSerializer  `json:"brand"`
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            Company CompanySerializer  `json:"company"`
            RejectReason string  `json:"reject_reason"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            CreatedBy UserSerializer  `json:"created_by"`
            VerifiedOn string  `json:"verified_on"`
            Warnings map[string]interface{}  `json:"warnings"`
            Stage string  `json:"stage"`
            UID float64  `json:"uid"`
         
    }
    
    // CompanyBrandListSerializer used by CompanyProfile
    type CompanyBrandListSerializer struct {

        
            Page Page  `json:"page"`
            Items []CompanyBrandSerializer  `json:"items"`
         
    }
    
    // CompanyBrandPostRequestSerializer used by CompanyProfile
    type CompanyBrandPostRequestSerializer struct {

        
            Company float64  `json:"company"`
            Brands []float64  `json:"brands"`
            UID float64  `json:"uid"`
         
    }
    
    // GetCompanySerializer used by CompanyProfile
    type GetCompanySerializer struct {

        
            ModifiedOn string  `json:"modified_on"`
            Name string  `json:"name"`
            CreatedOn string  `json:"created_on"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            RejectReason string  `json:"reject_reason"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            CreatedBy UserSerializer  `json:"created_by"`
            VerifiedOn string  `json:"verified_on"`
            BusinessType string  `json:"business_type"`
            Stage string  `json:"stage"`
            CompanyType string  `json:"company_type"`
            UID float64  `json:"uid"`
         
    }
    
    // InvoiceCredSerializer used by CompanyProfile
    type InvoiceCredSerializer struct {

        
            Enabled bool  `json:"enabled"`
            Password string  `json:"password"`
            Username string  `json:"username"`
         
    }
    
    // InvoiceDetailsSerializer used by CompanyProfile
    type InvoiceDetailsSerializer struct {

        
            EWaybill InvoiceCredSerializer  `json:"e_waybill"`
            EInvoice InvoiceCredSerializer  `json:"e_invoice"`
         
    }
    
    // LocationTimingSerializer used by CompanyProfile
    type LocationTimingSerializer struct {

        
            Minute float64  `json:"minute"`
            Hour float64  `json:"hour"`
         
    }
    
    // LocationDayWiseSerializer used by CompanyProfile
    type LocationDayWiseSerializer struct {

        
            Closing LocationTimingSerializer  `json:"closing"`
            Opening LocationTimingSerializer  `json:"opening"`
            Open bool  `json:"open"`
            Weekday string  `json:"weekday"`
         
    }
    
    // ProductReturnConfigSerializer used by CompanyProfile
    type ProductReturnConfigSerializer struct {

        
            StoreUID float64  `json:"store_uid"`
            OnSameStore bool  `json:"on_same_store"`
         
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
    
    // AverageOrderProcessingTime used by CompanyProfile
    type AverageOrderProcessingTime struct {

        
            Duration float64  `json:"duration"`
            DurationType string  `json:"duration_type"`
         
    }
    
    // LocationManagerSerializer used by CompanyProfile
    type LocationManagerSerializer struct {

        
            MobileNo SellerPhoneNumber  `json:"mobile_no"`
            Name string  `json:"name"`
            Email string  `json:"email"`
         
    }
    
    // GetLocationSerializer used by CompanyProfile
    type GetLocationSerializer struct {

        
            AutoInvoice bool  `json:"auto_invoice"`
            Name string  `json:"name"`
            CreditNote bool  `json:"credit_note"`
            NotificationEmails []string  `json:"notification_emails"`
            Company GetCompanySerializer  `json:"company"`
            Address GetAddressSerializer  `json:"address"`
            Code string  `json:"code"`
            UID float64  `json:"uid"`
            GstCredentials InvoiceDetailsSerializer  `json:"gst_credentials"`
            OrderAcceptanceTiming []LocationDayWiseSerializer  `json:"order_acceptance_timing"`
            ModifiedOn string  `json:"modified_on"`
            DisplayName string  `json:"display_name"`
            CreatedOn string  `json:"created_on"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            PhoneNumber string  `json:"phone_number"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            CreatedBy UserSerializer  `json:"created_by"`
            VerifiedOn string  `json:"verified_on"`
            StoreType string  `json:"store_type"`
            Tags []string  `json:"tags"`
            ProductReturnConfig ProductReturnConfigSerializer  `json:"product_return_config"`
            Holiday []HolidaySchemaSerializer  `json:"holiday"`
            AvgOrderProcessingTime AverageOrderProcessingTime  `json:"avg_order_processing_time"`
            Documents []Document  `json:"documents"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            Manager LocationManagerSerializer  `json:"manager"`
            Stage string  `json:"stage"`
            Warnings map[string]interface{}  `json:"warnings"`
            Timing []LocationDayWiseSerializer  `json:"timing"`
            DefaultOrderAcceptanceTiming bool  `json:"default_order_acceptance_timing"`
         
    }
    
    // LocationListSerializer used by CompanyProfile
    type LocationListSerializer struct {

        
            Page Page  `json:"page"`
            Items []GetLocationSerializer  `json:"items"`
         
    }
    
    // AddressSerializer used by CompanyProfile
    type AddressSerializer struct {

        
            CountryCode string  `json:"country_code"`
            Longitude float64  `json:"longitude"`
            AddressType string  `json:"address_type"`
            City string  `json:"city"`
            Country string  `json:"country"`
            Latitude float64  `json:"latitude"`
            Pincode float64  `json:"pincode"`
            Landmark string  `json:"landmark"`
            State string  `json:"state"`
            Address2 string  `json:"address2"`
            Address1 string  `json:"address1"`
         
    }
    
    // LocationSerializer used by CompanyProfile
    type LocationSerializer struct {

        
            AutoInvoice bool  `json:"auto_invoice"`
            Name string  `json:"name"`
            CreditNote bool  `json:"credit_note"`
            NotificationEmails []string  `json:"notification_emails"`
            Company float64  `json:"company"`
            Address AddressSerializer  `json:"address"`
            Code string  `json:"code"`
            UID float64  `json:"uid"`
            GstCredentials InvoiceDetailsSerializer  `json:"gst_credentials"`
            OrderAcceptanceTiming []LocationDayWiseSerializer  `json:"order_acceptance_timing"`
            DisplayName string  `json:"display_name"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Slug string  `json:"slug"`
            StoreType string  `json:"store_type"`
            Tags []string  `json:"tags"`
            ProductReturnConfig ProductReturnConfigSerializer  `json:"product_return_config"`
            Holiday []HolidaySchemaSerializer  `json:"holiday"`
            AvgOrderProcessingTime AverageOrderProcessingTime  `json:"avg_order_processing_time"`
            Documents []Document  `json:"documents"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            Manager LocationManagerSerializer  `json:"manager"`
            Stage string  `json:"stage"`
            Warnings map[string]interface{}  `json:"warnings"`
            Timing []LocationDayWiseSerializer  `json:"timing"`
            DefaultOrderAcceptanceTiming bool  `json:"default_order_acceptance_timing"`
         
    }
    
    // BulkLocationSerializer used by CompanyProfile
    type BulkLocationSerializer struct {

        
            Data []LocationSerializer  `json:"data"`
         
    }
    
    // StoreTagsResponse used by CompanyProfile
    type StoreTagsResponse struct {

        
            Tags []string  `json:"tags"`
            Success bool  `json:"success"`
         
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
    
    // Validity used by Cart
    type Validity struct {

        
            Priority float64  `json:"priority"`
         
    }
    
    // Validation used by Cart
    type Validation struct {

        
            Anonymous bool  `json:"anonymous"`
            AppID []string  `json:"app_id"`
            UserRegisteredAfter string  `json:"user_registered_after"`
         
    }
    
    // PaymentAllowValue used by Cart
    type PaymentAllowValue struct {

        
            Max float64  `json:"max"`
         
    }
    
    // PaymentModes used by Cart
    type PaymentModes struct {

        
            Iins []string  `json:"iins"`
            Types []string  `json:"types"`
            Codes []string  `json:"codes"`
            Uses PaymentAllowValue  `json:"uses"`
            Networks []string  `json:"networks"`
         
    }
    
    // PriceRange used by Cart
    type PriceRange struct {

        
            Max float64  `json:"max"`
            Min float64  `json:"min"`
         
    }
    
    // PostOrder used by Cart
    type PostOrder struct {

        
            CancellationAllowed bool  `json:"cancellation_allowed"`
            ReturnAllowed bool  `json:"return_allowed"`
         
    }
    
    // UsesRemaining used by Cart
    type UsesRemaining struct {

        
            App float64  `json:"app"`
            Total float64  `json:"total"`
            User float64  `json:"user"`
         
    }
    
    // UsesRestriction used by Cart
    type UsesRestriction struct {

        
            Maximum UsesRemaining  `json:"maximum"`
            Remaining UsesRemaining  `json:"remaining"`
         
    }
    
    // BulkBundleRestriction used by Cart
    type BulkBundleRestriction struct {

        
            MultiStoreAllowed bool  `json:"multi_store_allowed"`
         
    }
    
    // Restrictions used by Cart
    type Restrictions struct {

        
            Payments map[string]PaymentModes  `json:"payments"`
            PriceRange PriceRange  `json:"price_range"`
            PostOrder PostOrder  `json:"post_order"`
            Uses UsesRestriction  `json:"uses"`
            CouponAllowed bool  `json:"coupon_allowed"`
            OrderingStores []float64  `json:"ordering_stores"`
            UserGroups []float64  `json:"user_groups"`
            UserType string  `json:"user_type"`
            Platforms []string  `json:"platforms"`
            BulkBundle BulkBundleRestriction  `json:"bulk_bundle"`
         
    }
    
    // CouponAuthor used by Cart
    type CouponAuthor struct {

        
            ModifiedBy string  `json:"modified_by"`
            CreatedBy string  `json:"created_by"`
         
    }
    
    // RuleDefinition used by Cart
    type RuleDefinition struct {

        
            CalculateOn string  `json:"calculate_on"`
            AutoApply bool  `json:"auto_apply"`
            Scope []string  `json:"scope"`
            IsExact bool  `json:"is_exact"`
            CurrencyCode string  `json:"currency_code"`
            Type string  `json:"type"`
            ApplicableOn string  `json:"applicable_on"`
            ValueType string  `json:"value_type"`
         
    }
    
    // Identifier used by Cart
    type Identifier struct {

        
            UserID []string  `json:"user_id"`
            CategoryID []float64  `json:"category_id"`
            ArticleID []string  `json:"article_id"`
            BrandID []float64  `json:"brand_id"`
            CollectionID []string  `json:"collection_id"`
            ItemID []float64  `json:"item_id"`
            ExcludeBrandID []float64  `json:"exclude_brand_id"`
            CompanyID []float64  `json:"company_id"`
            EmailDomain []string  `json:"email_domain"`
            StoreID []float64  `json:"store_id"`
         
    }
    
    // Rule used by Cart
    type Rule struct {

        
            Value float64  `json:"value"`
            Key float64  `json:"key"`
            Min float64  `json:"min"`
            DiscountQty float64  `json:"discount_qty"`
            Max float64  `json:"max"`
         
    }
    
    // DisplayMetaDict used by Cart
    type DisplayMetaDict struct {

        
            Subtitle string  `json:"subtitle"`
            Title string  `json:"title"`
         
    }
    
    // DisplayMeta used by Cart
    type DisplayMeta struct {

        
            Apply DisplayMetaDict  `json:"apply"`
            Subtitle string  `json:"subtitle"`
            Title string  `json:"title"`
            Remove DisplayMetaDict  `json:"remove"`
            Auto DisplayMetaDict  `json:"auto"`
            Description string  `json:"description"`
         
    }
    
    // CouponAction used by Cart
    type CouponAction struct {

        
            ActionDate string  `json:"action_date"`
            TxnMode string  `json:"txn_mode"`
         
    }
    
    // CouponSchedule used by Cart
    type CouponSchedule struct {

        
            NextSchedule []map[string]interface{}  `json:"next_schedule"`
            Cron string  `json:"cron"`
            End string  `json:"end"`
            Start string  `json:"start"`
            Duration float64  `json:"duration"`
         
    }
    
    // Ownership used by Cart
    type Ownership struct {

        
            PayableCategory string  `json:"payable_category"`
            PayableBy string  `json:"payable_by"`
         
    }
    
    // CouponDateMeta used by Cart
    type CouponDateMeta struct {

        
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // State used by Cart
    type State struct {

        
            IsArchived bool  `json:"is_archived"`
            IsDisplay bool  `json:"is_display"`
            IsPublic bool  `json:"is_public"`
         
    }
    
    // CouponAdd used by Cart
    type CouponAdd struct {

        
            Validity Validity  `json:"validity"`
            Validation Validation  `json:"validation"`
            Tags []string  `json:"tags"`
            Restrictions Restrictions  `json:"restrictions"`
            Author CouponAuthor  `json:"author"`
            RuleDefinition RuleDefinition  `json:"rule_definition"`
            Identifiers Identifier  `json:"identifiers"`
            Rule []Rule  `json:"rule"`
            DisplayMeta DisplayMeta  `json:"display_meta"`
            Action CouponAction  `json:"action"`
            Schedule CouponSchedule  `json:"_schedule"`
            TypeSlug string  `json:"type_slug"`
            Code string  `json:"code"`
            Ownership Ownership  `json:"ownership"`
            DateMeta CouponDateMeta  `json:"date_meta"`
            State State  `json:"state"`
         
    }
    
    // CouponsResponse used by Cart
    type CouponsResponse struct {

        
            Page Page  `json:"page"`
            Items CouponAdd  `json:"items"`
         
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

        
            Validity Validity  `json:"validity"`
            Validation Validation  `json:"validation"`
            Tags []string  `json:"tags"`
            Restrictions Restrictions  `json:"restrictions"`
            Author CouponAuthor  `json:"author"`
            RuleDefinition RuleDefinition  `json:"rule_definition"`
            Identifiers Identifier  `json:"identifiers"`
            Rule []Rule  `json:"rule"`
            DisplayMeta DisplayMeta  `json:"display_meta"`
            Action CouponAction  `json:"action"`
            Schedule CouponSchedule  `json:"_schedule"`
            TypeSlug string  `json:"type_slug"`
            Code string  `json:"code"`
            Ownership Ownership  `json:"ownership"`
            DateMeta CouponDateMeta  `json:"date_meta"`
            State State  `json:"state"`
         
    }
    
    // CouponPartialUpdate used by Cart
    type CouponPartialUpdate struct {

        
            Archive bool  `json:"archive"`
            Schedule CouponSchedule  `json:"schedule"`
         
    }
    
    // PromotionAuthor used by Cart
    type PromotionAuthor struct {

        
            ModifiedBy string  `json:"modified_by"`
            CreatedBy string  `json:"created_by"`
         
    }
    
    // CompareObject used by Cart
    type CompareObject struct {

        
            LessThanEquals float64  `json:"less_than_equals"`
            GreaterThan float64  `json:"greater_than"`
            LessThan float64  `json:"less_than"`
            GreaterThanEquals float64  `json:"greater_than_equals"`
            Equals float64  `json:"equals"`
         
    }
    
    // ItemCriteria used by Cart
    type ItemCriteria struct {

        
            CartUniqueItemAmount CompareObject  `json:"cart_unique_item_amount"`
            ItemExcludeID []float64  `json:"item_exclude_id"`
            AllItems bool  `json:"all_items"`
            BuyRules []string  `json:"buy_rules"`
            ItemL1Category []float64  `json:"item_l1_category"`
            ItemExcludeSku []string  `json:"item_exclude_sku"`
            ItemStore []float64  `json:"item_store"`
            ItemExcludeBrand []float64  `json:"item_exclude_brand"`
            ItemBrand []float64  `json:"item_brand"`
            ItemTags []string  `json:"item_tags"`
            ProductTags []string  `json:"product_tags"`
            AvailableZones []string  `json:"available_zones"`
            ItemID []float64  `json:"item_id"`
            ItemCategory []float64  `json:"item_category"`
            CartTotal CompareObject  `json:"cart_total"`
            ItemExcludeL1Category []float64  `json:"item_exclude_l1_category"`
            CartQuantity CompareObject  `json:"cart_quantity"`
            ItemExcludeStore []float64  `json:"item_exclude_store"`
            ItemCompany []float64  `json:"item_company"`
            ItemExcludeL2Category []float64  `json:"item_exclude_l2_category"`
            ItemExcludeDepartment []float64  `json:"item_exclude_department"`
            ItemExcludeCategory []float64  `json:"item_exclude_category"`
            ItemSize []string  `json:"item_size"`
            CartUniqueItemQuantity CompareObject  `json:"cart_unique_item_quantity"`
            ItemL2Category []float64  `json:"item_l2_category"`
            ItemDepartment []float64  `json:"item_department"`
            ItemExcludeCompany []float64  `json:"item_exclude_company"`
            ItemSku []string  `json:"item_sku"`
         
    }
    
    // Ownership1 used by Cart
    type Ownership1 struct {

        
            PayableCategory string  `json:"payable_category"`
            PayableBy string  `json:"payable_by"`
         
    }
    
    // PromotionDateMeta used by Cart
    type PromotionDateMeta struct {

        
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // DisplayMeta1 used by Cart
    type DisplayMeta1 struct {

        
            Description string  `json:"description"`
            OfferText string  `json:"offer_text"`
            OfferLabel string  `json:"offer_label"`
            Name string  `json:"name"`
         
    }
    
    // PromotionAction used by Cart
    type PromotionAction struct {

        
            ActionDate string  `json:"action_date"`
            ActionType string  `json:"action_type"`
         
    }
    
    // PromotionSchedule used by Cart
    type PromotionSchedule struct {

        
            Published bool  `json:"published"`
            NextSchedule []map[string]interface{}  `json:"next_schedule"`
            Cron string  `json:"cron"`
            End string  `json:"end"`
            Start string  `json:"start"`
            Duration float64  `json:"duration"`
         
    }
    
    // Visibility used by Cart
    type Visibility struct {

        
            CouponList bool  `json:"coupon_list"`
            Pdp bool  `json:"pdp"`
         
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
    
    // UserRegistered used by Cart
    type UserRegistered struct {

        
            Start string  `json:"start"`
            End string  `json:"end"`
         
    }
    
    // PostOrder1 used by Cart
    type PostOrder1 struct {

        
            CancellationAllowed bool  `json:"cancellation_allowed"`
            ReturnAllowed bool  `json:"return_allowed"`
         
    }
    
    // UsesRemaining1 used by Cart
    type UsesRemaining1 struct {

        
            Total float64  `json:"total"`
            User float64  `json:"user"`
         
    }
    
    // UsesRestriction1 used by Cart
    type UsesRestriction1 struct {

        
            Maximum UsesRemaining1  `json:"maximum"`
            Remaining UsesRemaining1  `json:"remaining"`
         
    }
    
    // Restrictions1 used by Cart
    type Restrictions1 struct {

        
            Payments []PromotionPaymentModes  `json:"payments"`
            UserID []string  `json:"user_id"`
            UserRegistered UserRegistered  `json:"user_registered"`
            OrderQuantity float64  `json:"order_quantity"`
            AnonymousUsers bool  `json:"anonymous_users"`
            PostOrder PostOrder1  `json:"post_order"`
            Uses UsesRestriction1  `json:"uses"`
            OrderingStores []float64  `json:"ordering_stores"`
            UserGroups []float64  `json:"user_groups"`
            Platforms []string  `json:"platforms"`
         
    }
    
    // DiscountOffer used by Cart
    type DiscountOffer struct {

        
            MaxOfferQuantity float64  `json:"max_offer_quantity"`
            MinOfferQuantity float64  `json:"min_offer_quantity"`
            DiscountPrice float64  `json:"discount_price"`
            ApportionDiscount bool  `json:"apportion_discount"`
            DiscountAmount float64  `json:"discount_amount"`
            DiscountPercentage float64  `json:"discount_percentage"`
            Code string  `json:"code"`
            MaxUsagePerTransaction float64  `json:"max_usage_per_transaction"`
            PartialCanRet bool  `json:"partial_can_ret"`
            MaxDiscountAmount float64  `json:"max_discount_amount"`
         
    }
    
    // DiscountRule used by Cart
    type DiscountRule struct {

        
            Offer DiscountOffer  `json:"offer"`
            ItemCriteria ItemCriteria  `json:"item_criteria"`
            DiscountType string  `json:"discount_type"`
            BuyCondition string  `json:"buy_condition"`
         
    }
    
    // PromotionListItem used by Cart
    type PromotionListItem struct {

        
            Author PromotionAuthor  `json:"author"`
            ApplyAllDiscount bool  `json:"apply_all_discount"`
            BuyRules map[string]ItemCriteria  `json:"buy_rules"`
            Code string  `json:"code"`
            Ownership Ownership1  `json:"ownership"`
            DateMeta PromotionDateMeta  `json:"date_meta"`
            ApplicationID string  `json:"application_id"`
            Stackable bool  `json:"stackable"`
            CalculateOn string  `json:"calculate_on"`
            Currency string  `json:"currency"`
            DisplayMeta DisplayMeta1  `json:"display_meta"`
            PostOrderAction PromotionAction  `json:"post_order_action"`
            ApplyExclusive string  `json:"apply_exclusive"`
            Schedule PromotionSchedule  `json:"_schedule"`
            Visiblility Visibility  `json:"visiblility"`
            PromotionType string  `json:"promotion_type"`
            PromoGroup string  `json:"promo_group"`
            Mode string  `json:"mode"`
            Restrictions Restrictions1  `json:"restrictions"`
            ApplyPriority float64  `json:"apply_priority"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            DiscountRules []DiscountRule  `json:"discount_rules"`
         
    }
    
    // PromotionsResponse used by Cart
    type PromotionsResponse struct {

        
            Page Page  `json:"page"`
            Items PromotionListItem  `json:"items"`
         
    }
    
    // PromotionAdd used by Cart
    type PromotionAdd struct {

        
            Author PromotionAuthor  `json:"author"`
            ApplyAllDiscount bool  `json:"apply_all_discount"`
            BuyRules map[string]ItemCriteria  `json:"buy_rules"`
            Code string  `json:"code"`
            Ownership Ownership1  `json:"ownership"`
            DateMeta PromotionDateMeta  `json:"date_meta"`
            ApplicationID string  `json:"application_id"`
            Stackable bool  `json:"stackable"`
            CalculateOn string  `json:"calculate_on"`
            Currency string  `json:"currency"`
            DisplayMeta DisplayMeta1  `json:"display_meta"`
            PostOrderAction PromotionAction  `json:"post_order_action"`
            ApplyExclusive string  `json:"apply_exclusive"`
            Schedule PromotionSchedule  `json:"_schedule"`
            Visiblility Visibility  `json:"visiblility"`
            PromotionType string  `json:"promotion_type"`
            PromoGroup string  `json:"promo_group"`
            Mode string  `json:"mode"`
            Restrictions Restrictions1  `json:"restrictions"`
            ApplyPriority float64  `json:"apply_priority"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            DiscountRules []DiscountRule  `json:"discount_rules"`
         
    }
    
    // PromotionUpdate used by Cart
    type PromotionUpdate struct {

        
            Author PromotionAuthor  `json:"author"`
            ApplyAllDiscount bool  `json:"apply_all_discount"`
            BuyRules map[string]ItemCriteria  `json:"buy_rules"`
            Code string  `json:"code"`
            Ownership Ownership1  `json:"ownership"`
            DateMeta PromotionDateMeta  `json:"date_meta"`
            ApplicationID string  `json:"application_id"`
            Stackable bool  `json:"stackable"`
            CalculateOn string  `json:"calculate_on"`
            Currency string  `json:"currency"`
            DisplayMeta DisplayMeta1  `json:"display_meta"`
            PostOrderAction PromotionAction  `json:"post_order_action"`
            ApplyExclusive string  `json:"apply_exclusive"`
            Schedule PromotionSchedule  `json:"_schedule"`
            Visiblility Visibility  `json:"visiblility"`
            PromotionType string  `json:"promotion_type"`
            PromoGroup string  `json:"promo_group"`
            Mode string  `json:"mode"`
            Restrictions Restrictions1  `json:"restrictions"`
            ApplyPriority float64  `json:"apply_priority"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            DiscountRules []DiscountRule  `json:"discount_rules"`
         
    }
    
    // PromotionPartialUpdate used by Cart
    type PromotionPartialUpdate struct {

        
            Archive bool  `json:"archive"`
            Schedule PromotionSchedule  `json:"schedule"`
         
    }
    
    // ActivePromosResponse used by Cart
    type ActivePromosResponse struct {

        
            IsHidden bool  `json:"is_hidden"`
            EntitySlug string  `json:"entity_slug"`
            ModifiedOn string  `json:"modified_on"`
            Type string  `json:"type"`
            Subtitle string  `json:"subtitle"`
            CreatedOn string  `json:"created_on"`
            Title string  `json:"title"`
            EntityType string  `json:"entity_type"`
            Description string  `json:"description"`
            Example string  `json:"example"`
         
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

        
            BulkCoupons bool  `json:"bulk_coupons"`
            GiftPricing float64  `json:"gift_pricing"`
            RevenueEngineCoupon bool  `json:"revenue_engine_coupon"`
            Enabled bool  `json:"enabled"`
            DeliveryCharges DeliveryCharges  `json:"delivery_charges"`
            GiftDisplayText string  `json:"gift_display_text"`
            MaxCartItems float64  `json:"max_cart_items"`
            MinCartValue float64  `json:"min_cart_value"`
         
    }
    
    // CartMetaConfigAdd used by Cart
    type CartMetaConfigAdd struct {

        
            BulkCoupons bool  `json:"bulk_coupons"`
            GiftPricing float64  `json:"gift_pricing"`
            RevenueEngineCoupon bool  `json:"revenue_engine_coupon"`
            Enabled bool  `json:"enabled"`
            DeliveryCharges DeliveryCharges  `json:"delivery_charges"`
            GiftDisplayText string  `json:"gift_display_text"`
            MaxCartItems float64  `json:"max_cart_items"`
            MinCartValue float64  `json:"min_cart_value"`
         
    }
    
    // Collecttion used by Cart
    type Collecttion struct {

        
            RefundBy string  `json:"refund_by"`
            CollectedBy string  `json:"collected_by"`
         
    }
    
    // Article used by Cart
    type Article struct {

        
            Value float64  `json:"value"`
            ArticleID string  `json:"article_id"`
            Meta map[string]interface{}  `json:"meta"`
            Type string  `json:"type"`
            Code string  `json:"code"`
         
    }
    
    // CartDynamicInjectionUpdate used by Cart
    type CartDynamicInjectionUpdate struct {

        
            ApplyExpiry string  `json:"apply_expiry"`
            Value float64  `json:"value"`
            UserID string  `json:"user_id"`
            Collection Collecttion  `json:"collection"`
            Meta map[string]interface{}  `json:"meta"`
            CartID string  `json:"cart_id"`
            CartValue float64  `json:"cart_value"`
            AllowedRefund bool  `json:"allowed_refund"`
            Type string  `json:"type"`
            IsAuthenticated bool  `json:"is_authenticated"`
            Message string  `json:"message"`
            ArticleIds []Article  `json:"article_ids"`
            ArticleLevelDistribution bool  `json:"article_level_distribution"`
         
    }
    
    // CartDynamicInjection used by Cart
    type CartDynamicInjection struct {

        
            ApplyExpiry string  `json:"apply_expiry"`
            Value float64  `json:"value"`
            UserID string  `json:"user_id"`
            Collection Collecttion  `json:"collection"`
            Meta map[string]interface{}  `json:"meta"`
            CartID string  `json:"cart_id"`
            CartValue float64  `json:"cart_value"`
            AllowedRefund bool  `json:"allowed_refund"`
            Type string  `json:"type"`
            IsAuthenticated bool  `json:"is_authenticated"`
            InjectionID string  `json:"injection_id"`
            Message string  `json:"message"`
            ArticleIds []Article  `json:"article_ids"`
            ArticleLevelDistribution bool  `json:"article_level_distribution"`
         
    }
    
    // CartDynamicInjectionResponse used by Cart
    type CartDynamicInjectionResponse struct {

        
            Data CartDynamicInjection  `json:"data"`
         
    }
    
    // CartDynamicInjectionAdd used by Cart
    type CartDynamicInjectionAdd struct {

        
            ApplyExpiry string  `json:"apply_expiry"`
            Value float64  `json:"value"`
            UserID string  `json:"user_id"`
            Collection Collecttion  `json:"collection"`
            Meta map[string]interface{}  `json:"meta"`
            CartID string  `json:"cart_id"`
            CartValue float64  `json:"cart_value"`
            AllowedRefund bool  `json:"allowed_refund"`
            Type string  `json:"type"`
            IsAuthenticated bool  `json:"is_authenticated"`
            Message string  `json:"message"`
            ArticleIds []Article  `json:"article_ids"`
            ArticleLevelDistribution bool  `json:"article_level_distribution"`
         
    }
    
    // CartItem used by Cart
    type CartItem struct {

        
            Quantity float64  `json:"quantity"`
            Size string  `json:"size"`
            ProductID string  `json:"product_id"`
         
    }
    
    // OpenapiCartDetailsRequest used by Cart
    type OpenapiCartDetailsRequest struct {

        
            CartItems []CartItem  `json:"cart_items"`
         
    }
    
    // CouponBreakup used by Cart
    type CouponBreakup struct {

        
            CouponType string  `json:"coupon_type"`
            Value float64  `json:"value"`
            Description string  `json:"description"`
            CouponValue float64  `json:"coupon_value"`
            SubTitle string  `json:"sub_title"`
            Type string  `json:"type"`
            IsApplied bool  `json:"is_applied"`
            Title string  `json:"title"`
            Code string  `json:"code"`
            MinimumCartValue float64  `json:"minimum_cart_value"`
            MaxDiscountValue float64  `json:"max_discount_value"`
            Message string  `json:"message"`
            UID string  `json:"uid"`
         
    }
    
    // DisplayBreakup used by Cart
    type DisplayBreakup struct {

        
            Value float64  `json:"value"`
            CurrencySymbol string  `json:"currency_symbol"`
            Key string  `json:"key"`
            CurrencyCode string  `json:"currency_code"`
            Message []string  `json:"message"`
            Display string  `json:"display"`
         
    }
    
    // RawBreakup used by Cart
    type RawBreakup struct {

        
            ConvenienceFee float64  `json:"convenience_fee"`
            Coupon float64  `json:"coupon"`
            GstCharges float64  `json:"gst_charges"`
            GiftCard float64  `json:"gift_card"`
            Discount float64  `json:"discount"`
            YouSaved float64  `json:"you_saved"`
            FyndCash float64  `json:"fynd_cash"`
            Total float64  `json:"total"`
            MrpTotal float64  `json:"mrp_total"`
            Vog float64  `json:"vog"`
            CodCharge float64  `json:"cod_charge"`
            DeliveryCharge float64  `json:"delivery_charge"`
            Subtotal float64  `json:"subtotal"`
         
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

        
            Coupon CouponBreakup  `json:"coupon"`
            Display []DisplayBreakup  `json:"display"`
            Raw RawBreakup  `json:"raw"`
            LoyaltyPoints LoyaltyPoints  `json:"loyalty_points"`
         
    }
    
    // CouponDetails used by Cart
    type CouponDetails struct {

        
            DiscountSingleQuantity float64  `json:"discount_single_quantity"`
            Code string  `json:"code"`
            DiscountTotalQuantity float64  `json:"discount_total_quantity"`
         
    }
    
    // ProductAvailabilitySize used by Cart
    type ProductAvailabilitySize struct {

        
            Value string  `json:"value"`
            Meta map[string]interface{}  `json:"meta"`
            IsAvailable bool  `json:"is_available"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Display string  `json:"display"`
         
    }
    
    // ProductAvailability used by Cart
    type ProductAvailability struct {

        
            AvailableSizes []ProductAvailabilitySize  `json:"available_sizes"`
            IsValid bool  `json:"is_valid"`
            Sizes []string  `json:"sizes"`
            OtherStoreQuantity float64  `json:"other_store_quantity"`
            OutOfStock bool  `json:"out_of_stock"`
            Deliverable bool  `json:"deliverable"`
         
    }
    
    // CartProductIdentifer used by Cart
    type CartProductIdentifer struct {

        
            Identifier string  `json:"identifier"`
         
    }
    
    // FreeGiftItem used by Cart
    type FreeGiftItem struct {

        
            ItemImagesURL []string  `json:"item_images_url"`
            ItemSlug string  `json:"item_slug"`
            ItemName string  `json:"item_name"`
            ItemBrandName string  `json:"item_brand_name"`
            ItemID float64  `json:"item_id"`
            ItemPriceDetails map[string]interface{}  `json:"item_price_details"`
         
    }
    
    // AppliedFreeArticles used by Cart
    type AppliedFreeArticles struct {

        
            Quantity float64  `json:"quantity"`
            ArticleID string  `json:"article_id"`
            ParentItemIdentifier string  `json:"parent_item_identifier"`
            FreeGiftItemDetails FreeGiftItem  `json:"free_gift_item_details"`
         
    }
    
    // BuyRules used by Cart
    type BuyRules struct {

        
            CartConditions map[string]interface{}  `json:"cart_conditions"`
            ItemCriteria map[string]interface{}  `json:"item_criteria"`
         
    }
    
    // Ownership2 used by Cart
    type Ownership2 struct {

        
            PayableCategory string  `json:"payable_category"`
            PayableBy string  `json:"payable_by"`
         
    }
    
    // DiscountRulesApp used by Cart
    type DiscountRulesApp struct {

        
            Offer map[string]interface{}  `json:"offer"`
            ItemCriteria map[string]interface{}  `json:"item_criteria"`
            MatchedBuyRules []string  `json:"matched_buy_rules"`
            RawOffer map[string]interface{}  `json:"raw_offer"`
         
    }
    
    // AppliedPromotion used by Cart
    type AppliedPromotion struct {

        
            PromotionGroup string  `json:"promotion_group"`
            ArticleQuantity float64  `json:"article_quantity"`
            AppliedFreeArticles []AppliedFreeArticles  `json:"applied_free_articles"`
            PromotionOfferText string  `json:"promotion_offer_text"`
            MrpPromotion bool  `json:"mrp_promotion"`
            BuyRules []BuyRules  `json:"buy_rules"`
            PromotionName string  `json:"promotion_name"`
            Ownership Ownership2  `json:"ownership"`
            DiscountRules []DiscountRulesApp  `json:"discount_rules"`
            Amount float64  `json:"amount"`
            PromotionType string  `json:"promotion_type"`
            PromoID string  `json:"promo_id"`
         
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
    
    // ProductPrice used by Cart
    type ProductPrice struct {

        
            Selling float64  `json:"selling"`
            CurrencySymbol string  `json:"currency_symbol"`
            CurrencyCode string  `json:"currency_code"`
            Effective float64  `json:"effective"`
            Marked float64  `json:"marked"`
            AddOn float64  `json:"add_on"`
         
    }
    
    // ProductPriceInfo used by Cart
    type ProductPriceInfo struct {

        
            Converted ProductPrice  `json:"converted"`
            Base ProductPrice  `json:"base"`
         
    }
    
    // ProductImage used by Cart
    type ProductImage struct {

        
            URL string  `json:"url"`
            AspectRatio string  `json:"aspect_ratio"`
            SecureURL string  `json:"secure_url"`
         
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
            URL string  `json:"url"`
            Query ActionQuery  `json:"query"`
         
    }
    
    // BaseInfo used by Cart
    type BaseInfo struct {

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // Tags used by Cart
    type Tags struct {

        
            Tags map[string]interface{}  `json:"tags"`
         
    }
    
    // NetQuantity used by Cart
    type NetQuantity struct {

        
            Value string  `json:"value"`
            Unit string  `json:"unit"`
         
    }
    
    // CartProduct used by Cart
    type CartProduct struct {

        
            Images []ProductImage  `json:"images"`
            ItemCode string  `json:"item_code"`
            Tags []string  `json:"tags"`
            Slug string  `json:"slug"`
            Type string  `json:"type"`
            Categories []CategoryInfo  `json:"categories"`
            Action ProductAction  `json:"action"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Name string  `json:"name"`
            Brand BaseInfo  `json:"brand"`
            TeaserTag Tags  `json:"teaser_tag"`
            UID float64  `json:"uid"`
            NetQuantity NetQuantity  `json:"net_quantity"`
         
    }
    
    // PromoMeta used by Cart
    type PromoMeta struct {

        
            Message string  `json:"message"`
         
    }
    
    // BasePrice used by Cart
    type BasePrice struct {

        
            Marked float64  `json:"marked"`
            Effective float64  `json:"effective"`
            CurrencySymbol string  `json:"currency_symbol"`
            CurrencyCode string  `json:"currency_code"`
         
    }
    
    // ArticlePriceInfo used by Cart
    type ArticlePriceInfo struct {

        
            Converted BasePrice  `json:"converted"`
            Base BasePrice  `json:"base"`
         
    }
    
    // StoreInfo used by Cart
    type StoreInfo struct {

        
            StoreCode string  `json:"store_code"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // ProductArticle used by Cart
    type ProductArticle struct {

        
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            CartItemMeta map[string]interface{}  `json:"cart_item_meta"`
            Size string  `json:"size"`
            MtoQuantity float64  `json:"mto_quantity"`
            Price ArticlePriceInfo  `json:"price"`
            GiftCard map[string]interface{}  `json:"gift_card"`
            IsGiftVisible bool  `json:"is_gift_visible"`
            Type string  `json:"type"`
            ProductGroupTags []string  `json:"product_group_tags"`
            Quantity float64  `json:"quantity"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Seller BaseInfo  `json:"seller"`
            Store StoreInfo  `json:"store"`
            Identifier map[string]interface{}  `json:"identifier"`
            SellerIdentifier string  `json:"seller_identifier"`
            UID string  `json:"uid"`
         
    }
    
    // CartProductInfo used by Cart
    type CartProductInfo struct {

        
            CouponMessage string  `json:"coupon_message"`
            Coupon CouponDetails  `json:"coupon"`
            BulkOffer map[string]interface{}  `json:"bulk_offer"`
            Availability ProductAvailability  `json:"availability"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
            PromotionsApplied []AppliedPromotion  `json:"promotions_applied"`
            Quantity float64  `json:"quantity"`
            Message string  `json:"message"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            Price ProductPriceInfo  `json:"price"`
            Key string  `json:"key"`
            IsSet bool  `json:"is_set"`
            Product CartProduct  `json:"product"`
            Discount string  `json:"discount"`
            PromoMeta PromoMeta  `json:"promo_meta"`
            PricePerUnit ProductPriceInfo  `json:"price_per_unit"`
            Moq map[string]interface{}  `json:"moq"`
            Article ProductArticle  `json:"article"`
         
    }
    
    // OpenapiCartDetailsResponse used by Cart
    type OpenapiCartDetailsResponse struct {

        
            IsValid bool  `json:"is_valid"`
            Message string  `json:"message"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            Items []CartProductInfo  `json:"items"`
         
    }
    
    // OpenApiErrorResponse used by Cart
    type OpenApiErrorResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
            Errors map[string]interface{}  `json:"errors"`
         
    }
    
    // ShippingAddress used by Cart
    type ShippingAddress struct {

        
            Area string  `json:"area"`
            CountryPhoneCode string  `json:"country_phone_code"`
            Pincode float64  `json:"pincode"`
            AddressType string  `json:"address_type"`
            CountryIsoCode string  `json:"country_iso_code"`
            City string  `json:"city"`
            Meta map[string]interface{}  `json:"meta"`
            AreaCodeSlug string  `json:"area_code_slug"`
            AreaCode string  `json:"area_code"`
            Landmark string  `json:"landmark"`
            Email string  `json:"email"`
            Country string  `json:"country"`
            CountryCode string  `json:"country_code"`
            Address string  `json:"address"`
            Phone float64  `json:"phone"`
            Name string  `json:"name"`
            State string  `json:"state"`
         
    }
    
    // OpenApiCartServiceabilityRequest used by Cart
    type OpenApiCartServiceabilityRequest struct {

        
            ShippingAddress ShippingAddress  `json:"shipping_address"`
            CartItems []CartItem  `json:"cart_items"`
         
    }
    
    // OpenApiCartServiceabilityResponse used by Cart
    type OpenApiCartServiceabilityResponse struct {

        
            IsValid bool  `json:"is_valid"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            Message string  `json:"message"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            Items []CartProductInfo  `json:"items"`
         
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
    
    // MultiTenderPaymentMeta used by Cart
    type MultiTenderPaymentMeta struct {

        
            PaymentGateway string  `json:"payment_gateway"`
            CurrentStatus string  `json:"current_status"`
            PaymentID string  `json:"payment_id"`
            OrderID string  `json:"order_id"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
         
    }
    
    // MultiTenderPaymentMethod used by Cart
    type MultiTenderPaymentMethod struct {

        
            Meta MultiTenderPaymentMeta  `json:"meta"`
            Amount float64  `json:"amount"`
            Name string  `json:"name"`
            Mode string  `json:"mode"`
         
    }
    
    // OpenApiOrderItem used by Cart
    type OpenApiOrderItem struct {

        
            PriceEffective float64  `json:"price_effective"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            CashbackApplied float64  `json:"cashback_applied"`
            AmountPaid float64  `json:"amount_paid"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            Meta CartItemMeta  `json:"meta"`
            Size string  `json:"size"`
            EmployeeDiscount float64  `json:"employee_discount"`
            Discount float64  `json:"discount"`
            Quantity float64  `json:"quantity"`
            DeliveryCharges float64  `json:"delivery_charges"`
            Files []OpenApiFiles  `json:"files"`
            CodCharges float64  `json:"cod_charges"`
            ProductID float64  `json:"product_id"`
            PaymentMethods []MultiTenderPaymentMethod  `json:"payment_methods"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            PriceMarked float64  `json:"price_marked"`
         
    }
    
    // OpenApiPlatformCheckoutReq used by Cart
    type OpenApiPlatformCheckoutReq struct {

        
            CashbackApplied float64  `json:"cashback_applied"`
            Coupon string  `json:"coupon"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            EmployeeDiscount map[string]interface{}  `json:"employee_discount"`
            CurrencyCode string  `json:"currency_code"`
            Files []OpenApiFiles  `json:"files"`
            OrderID string  `json:"order_id"`
            Comment string  `json:"comment"`
            CouponCode string  `json:"coupon_code"`
            ShippingAddress ShippingAddress  `json:"shipping_address"`
            CartItems []OpenApiOrderItem  `json:"cart_items"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            CouponValue float64  `json:"coupon_value"`
            Gstin string  `json:"gstin"`
            CartValue float64  `json:"cart_value"`
            PaymentMode string  `json:"payment_mode"`
            BillingAddress ShippingAddress  `json:"billing_address"`
            DeliveryCharges float64  `json:"delivery_charges"`
            CodCharges float64  `json:"cod_charges"`
            PaymentMethods []MultiTenderPaymentMethod  `json:"payment_methods"`
         
    }
    
    // OpenApiCheckoutResponse used by Cart
    type OpenApiCheckoutResponse struct {

        
            OrderRefID string  `json:"order_ref_id"`
            Message string  `json:"message"`
            Success bool  `json:"success"`
            OrderID string  `json:"order_id"`
         
    }
    
    // AbandonedCart used by Cart
    type AbandonedCart struct {

        
            BulkCouponDiscount float64  `json:"bulk_coupon_discount"`
            Coupon map[string]interface{}  `json:"coupon"`
            Cashback map[string]interface{}  `json:"cashback"`
            Shipments []map[string]interface{}  `json:"shipments"`
            CreatedOn string  `json:"created_on"`
            OrderID string  `json:"order_id"`
            CheckoutMode string  `json:"checkout_mode"`
            BuyNow bool  `json:"buy_now"`
            Comment string  `json:"comment"`
            FyndCredits map[string]interface{}  `json:"fynd_credits"`
            Meta map[string]interface{}  `json:"meta"`
            Gstin string  `json:"gstin"`
            MergeQty bool  `json:"merge_qty"`
            DeliveryCharges map[string]interface{}  `json:"delivery_charges"`
            PaymentMode string  `json:"payment_mode"`
            Articles []map[string]interface{}  `json:"articles"`
            PaymentMethods []map[string]interface{}  `json:"payment_methods"`
            AppID string  `json:"app_id"`
            Payments map[string]interface{}  `json:"payments"`
            ExpireAt string  `json:"expire_at"`
            PickUpCustomerDetails map[string]interface{}  `json:"pick_up_customer_details"`
            FcIndexMap []float64  `json:"fc_index_map"`
            ID string  `json:"_id"`
            LastModified string  `json:"last_modified"`
            UserID string  `json:"user_id"`
            CartValue float64  `json:"cart_value"`
            IsDefault bool  `json:"is_default"`
            Discount float64  `json:"discount"`
            IsArchive bool  `json:"is_archive"`
            CodCharges map[string]interface{}  `json:"cod_charges"`
            IsActive bool  `json:"is_active"`
            UID float64  `json:"uid"`
            Promotion map[string]interface{}  `json:"promotion"`
         
    }
    
    // AbandonedCartResponse used by Cart
    type AbandonedCartResponse struct {

        
            Success bool  `json:"success"`
            Page Page  `json:"page"`
            Result map[string]interface{}  `json:"result"`
            Message string  `json:"message"`
            Items []AbandonedCart  `json:"items"`
         
    }
    
    // CartCurrency used by Cart
    type CartCurrency struct {

        
            Symbol string  `json:"symbol"`
            Code string  `json:"code"`
         
    }
    
    // PaymentSelectionLock used by Cart
    type PaymentSelectionLock struct {

        
            Enabled bool  `json:"enabled"`
            PaymentIdentifier string  `json:"payment_identifier"`
            DefaultOptions string  `json:"default_options"`
         
    }
    
    // CartDetailResponse used by Cart
    type CartDetailResponse struct {

        
            PanNo string  `json:"pan_no"`
            CheckoutMode string  `json:"checkout_mode"`
            ID string  `json:"id"`
            BuyNow bool  `json:"buy_now"`
            IsValid bool  `json:"is_valid"`
            Comment string  `json:"comment"`
            Gstin string  `json:"gstin"`
            Currency CartCurrency  `json:"currency"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            Message string  `json:"message"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            AppliedPromoDetails []AppliedPromotion  `json:"applied_promo_details"`
            LastModified string  `json:"last_modified"`
            CouponText string  `json:"coupon_text"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            PanConfig map[string]interface{}  `json:"pan_config"`
            Items []CartProductInfo  `json:"items"`
         
    }
    
    // AddProductCart used by Cart
    type AddProductCart struct {

        
            ParentItemIdentifiers []map[string]string  `json:"parent_item_identifiers"`
            Pos bool  `json:"pos"`
            ArticleID string  `json:"article_id"`
            Meta map[string]interface{}  `json:"meta"`
            SellerID float64  `json:"seller_id"`
            ItemSize string  `json:"item_size"`
            ProductGroupTags []string  `json:"product_group_tags"`
            Quantity float64  `json:"quantity"`
            ArticleAssignment map[string]interface{}  `json:"article_assignment"`
            Display string  `json:"display"`
            ItemID float64  `json:"item_id"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            StoreID float64  `json:"store_id"`
         
    }
    
    // AddCartRequest used by Cart
    type AddCartRequest struct {

        
            Items []AddProductCart  `json:"items"`
            NewCart bool  `json:"new_cart"`
         
    }
    
    // AddCartDetailResponse used by Cart
    type AddCartDetailResponse struct {

        
            Cart CartDetailResponse  `json:"cart"`
            Message string  `json:"message"`
            Success bool  `json:"success"`
            Partial bool  `json:"partial"`
         
    }
    
    // UpdateProductCart used by Cart
    type UpdateProductCart struct {

        
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            ItemIndex float64  `json:"item_index"`
            ArticleID string  `json:"article_id"`
            Meta map[string]interface{}  `json:"meta"`
            ItemSize string  `json:"item_size"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
            Quantity float64  `json:"quantity"`
            ItemID float64  `json:"item_id"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
         
    }
    
    // UpdateCartRequest used by Cart
    type UpdateCartRequest struct {

        
            Items []UpdateProductCart  `json:"items"`
            Operation string  `json:"operation"`
         
    }
    
    // UpdateCartDetailResponse used by Cart
    type UpdateCartDetailResponse struct {

        
            Cart CartDetailResponse  `json:"cart"`
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // OverrideCartItemPromo used by Cart
    type OverrideCartItemPromo struct {

        
            PromoDesc string  `json:"promo_desc"`
            ItemList []map[string]interface{}  `json:"item_list"`
            RwrdTndr string  `json:"rwrd_tndr"`
            PromoAmount string  `json:"promo_amount"`
            PromoID string  `json:"promo_id"`
         
    }
    
    // OverrideCartItem used by Cart
    type OverrideCartItem struct {

        
            PriceEffective float64  `json:"price_effective"`
            PromoList []OverrideCartItemPromo  `json:"promo_list"`
            AmountPaid float64  `json:"amount_paid"`
            Size string  `json:"size"`
            Discount float64  `json:"discount"`
            Quantity float64  `json:"quantity"`
            ItemID float64  `json:"item_id"`
            SellerIdentifier string  `json:"seller_identifier"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            PriceMarked float64  `json:"price_marked"`
         
    }
    
    // OverrideCheckoutReq used by Cart
    type OverrideCheckoutReq struct {

        
            PaymentIdentifier string  `json:"payment_identifier"`
            CartItems []OverrideCartItem  `json:"cart_items"`
            ShippingAddress map[string]interface{}  `json:"shipping_address"`
            OrderingStore float64  `json:"ordering_store"`
            Aggregator string  `json:"aggregator"`
            CartID string  `json:"cart_id"`
            CurrencyCode string  `json:"currency_code"`
            PaymentMode string  `json:"payment_mode"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
            MerchantCode string  `json:"merchant_code"`
            OrderType string  `json:"order_type"`
         
    }
    
    // OverrideCheckoutResponse used by Cart
    type OverrideCheckoutResponse struct {

        
            Data map[string]interface{}  `json:"data"`
            Cart map[string]interface{}  `json:"cart"`
            Success string  `json:"success"`
            OrderID string  `json:"order_id"`
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

        
            Meta map[string]interface{}  `json:"meta"`
            CreatedOn string  `json:"created_on"`
            Token string  `json:"token"`
            Source map[string]interface{}  `json:"source"`
            User map[string]interface{}  `json:"user"`
         
    }
    
    // SharedCart used by Cart
    type SharedCart struct {

        
            CheckoutMode string  `json:"checkout_mode"`
            ID string  `json:"id"`
            BuyNow bool  `json:"buy_now"`
            IsValid bool  `json:"is_valid"`
            Comment string  `json:"comment"`
            Gstin string  `json:"gstin"`
            Currency CartCurrency  `json:"currency"`
            Message string  `json:"message"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            CartID float64  `json:"cart_id"`
            LastModified string  `json:"last_modified"`
            CouponText string  `json:"coupon_text"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            SharedCartDetails SharedCartDetails  `json:"shared_cart_details"`
            UID string  `json:"uid"`
            BreakupValues CartBreakup  `json:"breakup_values"`
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
            CartID string  `json:"cart_id"`
            CartValue float64  `json:"cart_value"`
            PickUpCustomerDetails map[string]interface{}  `json:"pick_up_customer_details"`
            CreatedOn string  `json:"created_on"`
            ItemCounts float64  `json:"item_counts"`
         
    }
    
    // MultiCartResponse used by Cart
    type MultiCartResponse struct {

        
            Data []CartList  `json:"data"`
            Success bool  `json:"success"`
         
    }
    
    // UpdateUserCartMapping used by Cart
    type UpdateUserCartMapping struct {

        
            UserID string  `json:"user_id"`
         
    }
    
    // UserInfo used by Cart
    type UserInfo struct {

        
            ID string  `json:"_id"`
            UID string  `json:"uid"`
            Gender string  `json:"gender"`
            ModifiedOn string  `json:"modified_on"`
            LastName string  `json:"last_name"`
            ExternalID string  `json:"external_id"`
            FirstName string  `json:"first_name"`
            Mobile string  `json:"mobile"`
            CreatedAt string  `json:"created_at"`
         
    }
    
    // UserCartMappingResponse used by Cart
    type UserCartMappingResponse struct {

        
            PanNo string  `json:"pan_no"`
            CheckoutMode string  `json:"checkout_mode"`
            ID string  `json:"id"`
            BuyNow bool  `json:"buy_now"`
            IsValid bool  `json:"is_valid"`
            Comment string  `json:"comment"`
            Gstin string  `json:"gstin"`
            Currency CartCurrency  `json:"currency"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            Message string  `json:"message"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            AppliedPromoDetails []AppliedPromotion  `json:"applied_promo_details"`
            LastModified string  `json:"last_modified"`
            CouponText string  `json:"coupon_text"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            User UserInfo  `json:"user"`
            PanConfig map[string]interface{}  `json:"pan_config"`
            Items []CartProductInfo  `json:"items"`
         
    }
    
    // PlatformAddCartRequest used by Cart
    type PlatformAddCartRequest struct {

        
            Items []AddProductCart  `json:"items"`
            UserID string  `json:"user_id"`
            NewCart bool  `json:"new_cart"`
         
    }
    
    // PlatformUpdateCartRequest used by Cart
    type PlatformUpdateCartRequest struct {

        
            Items []UpdateProductCart  `json:"items"`
            UserID string  `json:"user_id"`
            Operation string  `json:"operation"`
         
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

        
            CouponCode string  `json:"coupon_code"`
            CouponType string  `json:"coupon_type"`
            CouponValue float64  `json:"coupon_value"`
            Message string  `json:"message"`
            IsApplicable bool  `json:"is_applicable"`
            SubTitle string  `json:"sub_title"`
            IsApplied bool  `json:"is_applied"`
            Title string  `json:"title"`
            MinimumCartValue float64  `json:"minimum_cart_value"`
            MaxDiscountValue float64  `json:"max_discount_value"`
            Description string  `json:"description"`
            ExpiresOn string  `json:"expires_on"`
         
    }
    
    // PageCoupon used by Cart
    type PageCoupon struct {

        
            TotalItemCount float64  `json:"total_item_count"`
            HasNext bool  `json:"has_next"`
            Current float64  `json:"current"`
            Total float64  `json:"total"`
            HasPrevious bool  `json:"has_previous"`
         
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

        
            Longitude float64  `json:"longitude"`
            Latitude float64  `json:"latitude"`
         
    }
    
    // PlatformAddress used by Cart
    type PlatformAddress struct {

        
            Area string  `json:"area"`
            CreatedByUserID string  `json:"created_by_user_id"`
            City string  `json:"city"`
            Landmark string  `json:"landmark"`
            Country string  `json:"country"`
            CheckoutMode string  `json:"checkout_mode"`
            ID string  `json:"id"`
            AreaCodeSlug string  `json:"area_code_slug"`
            Meta map[string]interface{}  `json:"meta"`
            AddressType string  `json:"address_type"`
            Tags []string  `json:"tags"`
            CartID string  `json:"cart_id"`
            Name string  `json:"name"`
            IsDefaultAddress bool  `json:"is_default_address"`
            State string  `json:"state"`
            UserID string  `json:"user_id"`
            AreaCode string  `json:"area_code"`
            CountryCode string  `json:"country_code"`
            GoogleMapPoint map[string]interface{}  `json:"google_map_point"`
            Address string  `json:"address"`
            Phone string  `json:"phone"`
            Email string  `json:"email"`
            IsActive bool  `json:"is_active"`
            GeoLocation GeoLocation  `json:"geo_location"`
         
    }
    
    // PlatformGetAddressesResponse used by Cart
    type PlatformGetAddressesResponse struct {

        
            Address []PlatformAddress  `json:"address"`
         
    }
    
    // SaveAddressResponse used by Cart
    type SaveAddressResponse struct {

        
            IsDefaultAddress bool  `json:"is_default_address"`
            ID string  `json:"id"`
            Success bool  `json:"success"`
         
    }
    
    // UpdateAddressResponse used by Cart
    type UpdateAddressResponse struct {

        
            IsDefaultAddress bool  `json:"is_default_address"`
            IsUpdated bool  `json:"is_updated"`
            ID string  `json:"id"`
            Success bool  `json:"success"`
         
    }
    
    // DeleteAddressResponse used by Cart
    type DeleteAddressResponse struct {

        
            IsDeleted bool  `json:"is_deleted"`
            ID string  `json:"id"`
         
    }
    
    // PlatformSelectCartAddressRequest used by Cart
    type PlatformSelectCartAddressRequest struct {

        
            UserID string  `json:"user_id"`
            CartID string  `json:"cart_id"`
            CheckoutMode string  `json:"checkout_mode"`
            ID string  `json:"id"`
            BillingAddressID string  `json:"billing_address_id"`
         
    }
    
    // ShipmentArticle used by Cart
    type ShipmentArticle struct {

        
            Quantity string  `json:"quantity"`
            ArticleID string  `json:"article_id"`
            Meta string  `json:"meta"`
         
    }
    
    // PlatformShipmentResponse used by Cart
    type PlatformShipmentResponse struct {

        
            FulfillmentID float64  `json:"fulfillment_id"`
            DpOptions map[string]interface{}  `json:"dp_options"`
            Shipments float64  `json:"shipments"`
            Promise ShipmentPromise  `json:"promise"`
            OrderType string  `json:"order_type"`
            Items []CartProductInfo  `json:"items"`
            FulfillmentType string  `json:"fulfillment_type"`
            DpID string  `json:"dp_id"`
            BoxType string  `json:"box_type"`
            ShipmentType string  `json:"shipment_type"`
            Articles []ShipmentArticle  `json:"articles"`
         
    }
    
    // PlatformCartShipmentsResponse used by Cart
    type PlatformCartShipmentsResponse struct {

        
            StaffUserID string  `json:"staff_user_id"`
            PanNo string  `json:"pan_no"`
            Shipments []PlatformShipmentResponse  `json:"shipments"`
            CheckoutMode string  `json:"checkout_mode"`
            ID string  `json:"id"`
            BuyNow bool  `json:"buy_now"`
            IsValid bool  `json:"is_valid"`
            Comment string  `json:"comment"`
            Gstin string  `json:"gstin"`
            Currency CartCurrency  `json:"currency"`
            Error bool  `json:"error"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            Message string  `json:"message"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            AppliedPromoDetails []AppliedPromotion  `json:"applied_promo_details"`
            LastModified string  `json:"last_modified"`
            CouponText string  `json:"coupon_text"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            PanConfig map[string]interface{}  `json:"pan_config"`
            Items []CartProductInfo  `json:"items"`
         
    }
    
    // UpdateCartShipmentItem used by Cart
    type UpdateCartShipmentItem struct {

        
            Quantity float64  `json:"quantity"`
            ArticleUID string  `json:"article_uid"`
            ShipmentType string  `json:"shipment_type"`
         
    }
    
    // UpdateCartShipmentRequest used by Cart
    type UpdateCartShipmentRequest struct {

        
            Shipments []UpdateCartShipmentItem  `json:"shipments"`
         
    }
    
    // PlatformCartMetaRequest used by Cart
    type PlatformCartMetaRequest struct {

        
            Comment string  `json:"comment"`
            StaffUserID string  `json:"staff_user_id"`
            GiftDetails map[string]interface{}  `json:"gift_details"`
            PanNo string  `json:"pan_no"`
            Gstin string  `json:"gstin"`
            PickUpCustomerDetails map[string]interface{}  `json:"pick_up_customer_details"`
            CheckoutMode string  `json:"checkout_mode"`
         
    }
    
    // CartMetaResponse used by Cart
    type CartMetaResponse struct {

        
            IsValid bool  `json:"is_valid"`
            Message string  `json:"message"`
         
    }
    
    // CartMetaMissingResponse used by Cart
    type CartMetaMissingResponse struct {

        
            Errors []string  `json:"errors"`
         
    }
    
    // StaffCheckout used by Cart
    type StaffCheckout struct {

        
            ID string  `json:"_id"`
            User string  `json:"user"`
            LastName string  `json:"last_name"`
            EmployeeCode string  `json:"employee_code"`
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
            CheckoutMode string  `json:"checkout_mode"`
            ID string  `json:"id"`
            PaymentIdentifier string  `json:"payment_identifier"`
            AddressID string  `json:"address_id"`
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
            Meta map[string]interface{}  `json:"meta"`
            PaymentMode string  `json:"payment_mode"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
            MerchantCode string  `json:"merchant_code"`
            PaymentParams map[string]interface{}  `json:"payment_params"`
            Staff StaffCheckout  `json:"staff"`
            Files []Files  `json:"files"`
            BillingAddressID string  `json:"billing_address_id"`
            EmployeeCode string  `json:"employee_code"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            CallbackURL string  `json:"callback_url"`
            PaymentAutoConfirm bool  `json:"payment_auto_confirm"`
            Pos bool  `json:"pos"`
            DeviceID string  `json:"device_id"`
            OrderingStore float64  `json:"ordering_store"`
            UserID string  `json:"user_id"`
            OrderType string  `json:"order_type"`
         
    }
    
    // CheckCart used by Cart
    type CheckCart struct {

        
            StoreCode string  `json:"store_code"`
            CheckoutMode string  `json:"checkout_mode"`
            ID string  `json:"id"`
            BuyNow bool  `json:"buy_now"`
            OrderID string  `json:"order_id"`
            CodAvailable bool  `json:"cod_available"`
            IsValid bool  `json:"is_valid"`
            Comment string  `json:"comment"`
            Gstin string  `json:"gstin"`
            Currency CartCurrency  `json:"currency"`
            DeliveryCharges float64  `json:"delivery_charges"`
            ErrorMessage string  `json:"error_message"`
            Message string  `json:"message"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            CartID float64  `json:"cart_id"`
            CodMessage string  `json:"cod_message"`
            StoreEmps []map[string]interface{}  `json:"store_emps"`
            UserType string  `json:"user_type"`
            DeliveryChargeOrderValue float64  `json:"delivery_charge_order_value"`
            LastModified string  `json:"last_modified"`
            CouponText string  `json:"coupon_text"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            Success bool  `json:"success"`
            CodCharges float64  `json:"cod_charges"`
            UID string  `json:"uid"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            Items []CartProductInfo  `json:"items"`
         
    }
    
    // CartCheckoutResponse used by Cart
    type CartCheckoutResponse struct {

        
            Data map[string]interface{}  `json:"data"`
            Cart CheckCart  `json:"cart"`
            AppInterceptURL string  `json:"app_intercept_url"`
            Success bool  `json:"success"`
            OrderID string  `json:"order_id"`
            Message string  `json:"message"`
            PaymentConfirmURL string  `json:"payment_confirm_url"`
            CallbackURL string  `json:"callback_url"`
         
    }
    
    // CartDeliveryModesResponse used by Cart
    type CartDeliveryModesResponse struct {

        
            AvailableModes []string  `json:"available_modes"`
            PickupStores []float64  `json:"pickup_stores"`
         
    }
    
    // PickupStoreDetail used by Cart
    type PickupStoreDetail struct {

        
            Area string  `json:"area"`
            UID float64  `json:"uid"`
            Pincode float64  `json:"pincode"`
            AddressType string  `json:"address_type"`
            Name string  `json:"name"`
            City string  `json:"city"`
            AreaCodeSlug string  `json:"area_code_slug"`
            StoreCode string  `json:"store_code"`
            AreaCode string  `json:"area_code"`
            Landmark string  `json:"landmark"`
            Country string  `json:"country"`
            Address string  `json:"address"`
            Phone string  `json:"phone"`
            Email string  `json:"email"`
            ID float64  `json:"id"`
            State string  `json:"state"`
         
    }
    
    // StoreDetailsResponse used by Cart
    type StoreDetailsResponse struct {

        
            Items []PickupStoreDetail  `json:"items"`
         
    }
    
    // UpdateCartPaymentRequest used by Cart
    type UpdateCartPaymentRequest struct {

        
            PaymentIdentifier string  `json:"payment_identifier"`
            AddressID string  `json:"address_id"`
            PaymentMode string  `json:"payment_mode"`
            MerchantCode string  `json:"merchant_code"`
            ID string  `json:"id"`
            AggregatorName string  `json:"aggregator_name"`
         
    }
    
    // CouponValidity used by Cart
    type CouponValidity struct {

        
            DisplayMessageEn string  `json:"display_message_en"`
            Discount float64  `json:"discount"`
            Title string  `json:"title"`
            Code string  `json:"code"`
            NextValidationRequired bool  `json:"next_validation_required"`
            Valid bool  `json:"valid"`
         
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
            PaymentGateway string  `json:"payment_gateway"`
            MerchantCode string  `json:"merchant_code"`
            PaymentIdentifier string  `json:"payment_identifier"`
         
    }
    
    // PaymentMethod used by Cart
    type PaymentMethod struct {

        
            PaymentMeta PaymentMeta  `json:"payment_meta"`
            Payment string  `json:"payment"`
            Name string  `json:"name"`
            Amount float64  `json:"amount"`
            Mode string  `json:"mode"`
         
    }
    
    // PlatformCartCheckoutDetailV2Request used by Cart
    type PlatformCartCheckoutDetailV2Request struct {

        
            PickAtStoreUID float64  `json:"pick_at_store_uid"`
            Aggregator string  `json:"aggregator"`
            CheckoutMode string  `json:"checkout_mode"`
            ID string  `json:"id"`
            PaymentIdentifier string  `json:"payment_identifier"`
            AddressID string  `json:"address_id"`
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
            Meta map[string]interface{}  `json:"meta"`
            PaymentMode string  `json:"payment_mode"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
            MerchantCode string  `json:"merchant_code"`
            PaymentParams map[string]interface{}  `json:"payment_params"`
            PaymentMethods []PaymentMethod  `json:"payment_methods"`
            CustomMeta map[string]interface{}  `json:"custom_meta"`
            Staff StaffCheckout  `json:"staff"`
            Files []Files  `json:"files"`
            BillingAddressID string  `json:"billing_address_id"`
            EmployeeCode string  `json:"employee_code"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            CallbackURL string  `json:"callback_url"`
            PaymentAutoConfirm bool  `json:"payment_auto_confirm"`
            Pos bool  `json:"pos"`
            DeviceID string  `json:"device_id"`
            OrderingStore float64  `json:"ordering_store"`
            UserID string  `json:"user_id"`
            OrderType string  `json:"order_type"`
         
    }
    
    // UpdateCartPaymentRequestV2 used by Cart
    type UpdateCartPaymentRequestV2 struct {

        
            PaymentIdentifier string  `json:"payment_identifier"`
            AddressID string  `json:"address_id"`
            PaymentMode string  `json:"payment_mode"`
            MerchantCode string  `json:"merchant_code"`
            ID string  `json:"id"`
            PaymentMethods []PaymentMethod  `json:"payment_methods"`
            AggregatorName string  `json:"aggregator_name"`
         
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
            ID string  `json:"_id"`
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
    
    // EventConfigDetails used by Webhook
    type EventConfigDetails struct {

        
            EventName string  `json:"event_name"`
            EventType string  `json:"event_type"`
            EventCategory string  `json:"event_category"`
            Version string  `json:"version"`
            DisplayName string  `json:"display_name"`
            Description string  `json:"description"`
            CreatedOn string  `json:"created_on"`
         
    }
    
    // SubscriberConfigDetails used by Webhook
    type SubscriberConfigDetails struct {

        
            Name string  `json:"name"`
            WebhookURL string  `json:"webhook_url"`
            Association Association  `json:"association"`
            CustomHeaders map[string]interface{}  `json:"custom_headers"`
            Status SubscriberStatus  `json:"status"`
            EmailID string  `json:"email_id"`
            AuthMeta AuthMeta  `json:"auth_meta"`
            EventID []float64  `json:"event_id"`
         
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
    

    
    // ApplicationServiceabilityConfig used by Serviceability
    type ApplicationServiceabilityConfig struct {

        
            ServiceabilityType string  `json:"serviceability_type"`
            ChannelType string  `json:"channel_type"`
            ChannelID string  `json:"channel_id"`
         
    }
    
    // ServiceabilityErrorResponse used by Serviceability
    type ServiceabilityErrorResponse struct {

        
            Type string  `json:"type"`
            Message string  `json:"message"`
            Value string  `json:"value"`
         
    }
    
    // ApplicationServiceabilityConfigResponse used by Serviceability
    type ApplicationServiceabilityConfigResponse struct {

        
            Data ApplicationServiceabilityConfig  `json:"data"`
            Success bool  `json:"success"`
            Error ServiceabilityErrorResponse  `json:"error"`
         
    }
    
    // EntityRegionView_Request used by Serviceability
    type EntityRegionView_Request struct {

        
            SubType []string  `json:"sub_type"`
            ParentID []string  `json:"parent_id"`
         
    }
    
    // EntityRegionView_page used by Serviceability
    type EntityRegionView_page struct {

        
            Current float64  `json:"current"`
            Type string  `json:"type"`
            Size float64  `json:"size"`
            ItemTotal float64  `json:"item_total"`
            HasNext bool  `json:"has_next"`
         
    }
    
    // EntityRegionView_Items used by Serviceability
    type EntityRegionView_Items struct {

        
            Name string  `json:"name"`
            SubType string  `json:"sub_type"`
            UID string  `json:"uid"`
         
    }
    
    // EntityRegionView_Error used by Serviceability
    type EntityRegionView_Error struct {

        
            Type string  `json:"type"`
            Message string  `json:"message"`
            Value string  `json:"value"`
         
    }
    
    // EntityRegionView_Response used by Serviceability
    type EntityRegionView_Response struct {

        
            Success bool  `json:"success"`
            Page EntityRegionView_page  `json:"page"`
            Data []EntityRegionView_Items  `json:"data"`
            Error EntityRegionView_Error  `json:"error"`
         
    }
    
    // ZoneDataItem used by Serviceability
    type ZoneDataItem struct {

        
            Current float64  `json:"current"`
            Type string  `json:"type"`
            ItemTotal float64  `json:"item_total"`
            Size float64  `json:"size"`
            HasNext bool  `json:"has_next"`
         
    }
    
    // ListViewSummary used by Serviceability
    type ListViewSummary struct {

        
            TotalZones float64  `json:"total_zones"`
            TotalPincodesServed float64  `json:"total_pincodes_served"`
            TotalActiveZones float64  `json:"total_active_zones"`
         
    }
    
    // ListViewChannels used by Serviceability
    type ListViewChannels struct {

        
            ChannelType string  `json:"channel_type"`
            ChannelID string  `json:"channel_id"`
         
    }
    
    // ListViewProduct used by Serviceability
    type ListViewProduct struct {

        
            Count float64  `json:"count"`
            Type string  `json:"type"`
         
    }
    
    // ListViewItems used by Serviceability
    type ListViewItems struct {

        
            Name string  `json:"name"`
            StoresCount float64  `json:"stores_count"`
            CompanyID float64  `json:"company_id"`
            PincodesCount float64  `json:"pincodes_count"`
            ZoneID string  `json:"zone_id"`
            Channels ListViewChannels  `json:"channels"`
            IsActive bool  `json:"is_active"`
            Product ListViewProduct  `json:"product"`
            Slug string  `json:"slug"`
         
    }
    
    // ListViewResponse used by Serviceability
    type ListViewResponse struct {

        
            Page []ZoneDataItem  `json:"page"`
            Summary []ListViewSummary  `json:"summary"`
            Items []ListViewItems  `json:"items"`
         
    }
    
    // CompanyStoreView_PageItems used by Serviceability
    type CompanyStoreView_PageItems struct {

        
            Current float64  `json:"current"`
            Type string  `json:"type"`
            Size float64  `json:"size"`
            ItemTotal float64  `json:"item_total"`
            HasNext bool  `json:"has_next"`
         
    }
    
    // CompanyStoreView_Response used by Serviceability
    type CompanyStoreView_Response struct {

        
            Page []CompanyStoreView_PageItems  `json:"page"`
            Items []map[string]interface{}  `json:"items"`
         
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

        
            State []string  `json:"state"`
            Country string  `json:"country"`
            Pincode []string  `json:"pincode"`
         
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

        
            Identifier string  `json:"identifier"`
            Data UpdateZoneData  `json:"data"`
         
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

        
            Identifier string  `json:"identifier"`
            Data CreateZoneData  `json:"data"`
         
    }
    
    // ZoneResponse used by Serviceability
    type ZoneResponse struct {

        
            ZoneID string  `json:"zone_id"`
            Success bool  `json:"success"`
            StatusCode float64  `json:"status_code"`
         
    }
    
    // GetZoneFromPincodeViewRequest used by Serviceability
    type GetZoneFromPincodeViewRequest struct {

        
            Country string  `json:"country"`
            Pincode string  `json:"pincode"`
         
    }
    
    // GetZoneFromPincodeViewResponse used by Serviceability
    type GetZoneFromPincodeViewResponse struct {

        
            Zones []string  `json:"zones"`
            ServiceabilityType string  `json:"serviceability_type"`
         
    }
    
    // GetZoneFromApplicationIdViewResponse used by Serviceability
    type GetZoneFromApplicationIdViewResponse struct {

        
            Page []ZoneDataItem  `json:"page"`
            Items []ListViewItems  `json:"items"`
         
    }
    
    // ServiceabilityPageResponse used by Serviceability
    type ServiceabilityPageResponse struct {

        
            Current float64  `json:"current"`
            Type string  `json:"type"`
            Size float64  `json:"size"`
            ItemTotal float64  `json:"item_total"`
            HasNext bool  `json:"has_next"`
         
    }
    
    // CreatedByResponse used by Serviceability
    type CreatedByResponse struct {

        
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
         
    }
    
    // IntegrationTypeResponse used by Serviceability
    type IntegrationTypeResponse struct {

        
            Inventory string  `json:"inventory"`
            Order string  `json:"order"`
         
    }
    
    // ModifiedByResponse used by Serviceability
    type ModifiedByResponse struct {

        
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
         
    }
    
    // ProductReturnConfigResponse used by Serviceability
    type ProductReturnConfigResponse struct {

        
            OnSameStore bool  `json:"on_same_store"`
         
    }
    
    // OpeningClosing used by Serviceability
    type OpeningClosing struct {

        
            Hour float64  `json:"hour"`
            Minute float64  `json:"minute"`
         
    }
    
    // TimmingResponse used by Serviceability
    type TimmingResponse struct {

        
            Opening OpeningClosing  `json:"opening"`
            Closing OpeningClosing  `json:"closing"`
            Open bool  `json:"open"`
            Weekday string  `json:"weekday"`
         
    }
    
    // AddressResponse used by Serviceability
    type AddressResponse struct {

        
            Landmark string  `json:"landmark"`
            City string  `json:"city"`
            Pincode float64  `json:"pincode"`
            State string  `json:"state"`
            Longitude float64  `json:"longitude"`
            Address2 string  `json:"address2"`
            Latitude float64  `json:"latitude"`
            Country string  `json:"country"`
            Address1 string  `json:"address1"`
         
    }
    
    // EwayBillResponse used by Serviceability
    type EwayBillResponse struct {

        
            Enabled bool  `json:"enabled"`
         
    }
    
    // EinvoiceResponse used by Serviceability
    type EinvoiceResponse struct {

        
            Enabled bool  `json:"enabled"`
         
    }
    
    // GstCredentialsResponse used by Serviceability
    type GstCredentialsResponse struct {

        
            EWaybill EwayBillResponse  `json:"e_waybill"`
            EInvoice EinvoiceResponse  `json:"e_invoice"`
         
    }
    
    // DocumentsResponse used by Serviceability
    type DocumentsResponse struct {

        
            Type string  `json:"type"`
            Verified bool  `json:"verified"`
            LegalName string  `json:"legal_name"`
            Value string  `json:"value"`
         
    }
    
    // MobileNo used by Serviceability
    type MobileNo struct {

        
            Number string  `json:"number"`
            CountryCode float64  `json:"country_code"`
         
    }
    
    // ManagerResponse used by Serviceability
    type ManagerResponse struct {

        
            Name string  `json:"name"`
            Email string  `json:"email"`
            MobileNo MobileNo  `json:"mobile_no"`
         
    }
    
    // Dp used by Serviceability
    type Dp struct {

        
            AssignDpFromSb bool  `json:"assign_dp_from_sb"`
            Operations []string  `json:"operations"`
            RvpPriority float64  `json:"rvp_priority"`
            LmPriority float64  `json:"lm_priority"`
            FmPriority float64  `json:"fm_priority"`
            InternalAccountID string  `json:"internal_account_id"`
            ExternalAccountID string  `json:"external_account_id"`
            TransportMode string  `json:"transport_mode"`
            AreaCode float64  `json:"area_code"`
            PaymentMode string  `json:"payment_mode"`
         
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
    
    // WarningsResponse used by Serviceability
    type WarningsResponse struct {

        
            StoreAddress string  `json:"store_address"`
         
    }
    
    // ItemResponse used by Serviceability
    type ItemResponse struct {

        
            CreatedOn string  `json:"created_on"`
            CreatedBy CreatedByResponse  `json:"created_by"`
            StoreType string  `json:"store_type"`
            IntegrationType IntegrationTypeResponse  `json:"integration_type"`
            Name string  `json:"name"`
            Stage string  `json:"stage"`
            ModifiedBy ModifiedByResponse  `json:"modified_by"`
            ProductReturnConfig ProductReturnConfigResponse  `json:"product_return_config"`
            ModifiedOn string  `json:"modified_on"`
            Timing []TimmingResponse  `json:"timing"`
            VerifiedOn string  `json:"verified_on"`
            Address AddressResponse  `json:"address"`
            DisplayName string  `json:"display_name"`
            Code string  `json:"code"`
            GstCredentials GstCredentialsResponse  `json:"gst_credentials"`
            CompanyID float64  `json:"company_id"`
            Company float64  `json:"company"`
            Documents []DocumentsResponse  `json:"documents"`
            NotificationEmails []string  `json:"notification_emails"`
            VerifiedBy ModifiedByResponse  `json:"verified_by"`
            Manager ManagerResponse  `json:"manager"`
            UID float64  `json:"uid"`
            Logistics LogisticsResponse  `json:"logistics"`
            SubType string  `json:"sub_type"`
            Cls string  `json:"_cls"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ContactNumbers []ContactNumberResponse  `json:"contact_numbers"`
            Warnings WarningsResponse  `json:"warnings"`
         
    }
    
    // GetStoresViewResponse used by Serviceability
    type GetStoresViewResponse struct {

        
            Page ServiceabilityPageResponse  `json:"page"`
            Items []ItemResponse  `json:"items"`
         
    }
    
    // ReAssignStoreRequest used by Serviceability
    type ReAssignStoreRequest struct {

        
            Articles []map[string]interface{}  `json:"articles"`
            Identifier string  `json:"identifier"`
            ToPincode string  `json:"to_pincode"`
            Configuration map[string]interface{}  `json:"configuration"`
            IgnoredLocations []string  `json:"ignored_locations"`
         
    }
    
    // ReAssignStoreResponse used by Serviceability
    type ReAssignStoreResponse struct {

        
            Articles []map[string]interface{}  `json:"articles"`
            ToPincode string  `json:"to_pincode"`
            Success bool  `json:"success"`
            Error map[string]interface{}  `json:"error"`
         
    }
    
    // ApplicationCompanyDpViewRequest used by Serviceability
    type ApplicationCompanyDpViewRequest struct {

        
            DpID string  `json:"dp_id"`
         
    }
    
    // ApplicationCompanyDpViewResponse used by Serviceability
    type ApplicationCompanyDpViewResponse struct {

        
            ApplicationID string  `json:"application_id"`
            CourierPartnerID float64  `json:"courier_partner_id"`
            Success bool  `json:"success"`
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

        
            Success string  `json:"success"`
            StatusCode string  `json:"status_code"`
            Error interface{}  `json:"error"`
         
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

        
            Current float64  `json:"current"`
            Type string  `json:"type"`
            Size float64  `json:"size"`
            ItemTotal float64  `json:"item_total"`
            HasNext bool  `json:"has_next"`
         
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
    
    // Dp1 used by Serviceability
    type Dp1 struct {

        
            Name string  `json:"name"`
            PlanRules map[string]interface{}  `json:"plan_rules"`
            Stage string  `json:"stage"`
            AccountID string  `json:"account_id"`
            DpID string  `json:"dp_id"`
            IsSelfShip bool  `json:"is_self_ship"`
            PlanID string  `json:"plan_id"`
         
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

        
            Type string  `json:"type"`
            Message string  `json:"message"`
            Value string  `json:"value"`
         
    }
    
    // DpAccountFailureResponse used by Serviceability
    type DpAccountFailureResponse struct {

        
            Success bool  `json:"success"`
            StatusCode float64  `json:"status_code"`
            Error []ErrorResponse  `json:"error"`
         
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

        
            Success bool  `json:"success"`
            Page Page  `json:"page"`
            Items []Dp1  `json:"items"`
         
    }
    
    // DpRulesUpdateRequest used by Serviceability
    type DpRulesUpdateRequest struct {

        
            Conditions []map[string]interface{}  `json:"conditions"`
            Name string  `json:"name"`
            IsActive bool  `json:"is_active"`
            DpIds map[string]map[string]interface{}  `json:"dp_ids"`
         
    }
    
    // DpRuleResponse used by Serviceability
    type DpRuleResponse struct {

        
            CreatedOn string  `json:"created_on"`
            Name string  `json:"name"`
            CompanyID float64  `json:"company_id"`
            Conditions []string  `json:"conditions"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            ModifiedOn string  `json:"modified_on"`
            IsActive bool  `json:"is_active"`
            DpIds map[string]interface{}  `json:"dp_ids"`
            UID string  `json:"uid"`
         
    }
    
    // DpRuleUpdateSuccessResponse used by Serviceability
    type DpRuleUpdateSuccessResponse struct {

        
            StatusCode float64  `json:"status_code"`
            Success bool  `json:"success"`
            Data DpRuleResponse  `json:"data"`
         
    }
    
    // FailureResponse used by Serviceability
    type FailureResponse struct {

        
            Success bool  `json:"success"`
            StatusCode float64  `json:"status_code"`
            Error []ErrorResponse  `json:"error"`
         
    }
    
    // DpSchemaInRuleListing used by Serviceability
    type DpSchemaInRuleListing struct {

        
            Name string  `json:"name"`
            Priority float64  `json:"priority"`
            PlanRules map[string]interface{}  `json:"plan_rules"`
            Stage string  `json:"stage"`
            AccountID string  `json:"account_id"`
            DpID string  `json:"dp_id"`
            IsSelfShip bool  `json:"is_self_ship"`
            PlanID string  `json:"plan_id"`
         
    }
    
    // DpRule used by Serviceability
    type DpRule struct {

        
            Name string  `json:"name"`
            CompanyID float64  `json:"company_id"`
            Conditions []map[string]interface{}  `json:"conditions"`
            IsActive bool  `json:"is_active"`
            DpIds map[string]DpSchemaInRuleListing  `json:"dp_ids"`
         
    }
    
    // DpRuleSuccessResponse used by Serviceability
    type DpRuleSuccessResponse struct {

        
            StatusCode float64  `json:"status_code"`
            Success bool  `json:"success"`
            Data DpRule  `json:"data"`
         
    }
    
    // DpIds used by Serviceability
    type DpIds struct {

        
            Priority float64  `json:"priority"`
            Meta map[string]interface{}  `json:"meta"`
            Enabled bool  `json:"enabled"`
         
    }
    
    // DpRuleRequest used by Serviceability
    type DpRuleRequest struct {

        
            Name string  `json:"name"`
            CompanyID float64  `json:"company_id"`
            Conditions []map[string]interface{}  `json:"conditions"`
            IsActive bool  `json:"is_active"`
            DpIds map[string]DpIds  `json:"dp_ids"`
         
    }
    
    // DpMultipleRuleSuccessResponse used by Serviceability
    type DpMultipleRuleSuccessResponse struct {

        
            Success bool  `json:"success"`
            Page Page  `json:"page"`
            Items []DpRule  `json:"items"`
         
    }
    
    // DPCompanyRuleRequest used by Serviceability
    type DPCompanyRuleRequest struct {

        
            RuleIds []string  `json:"rule_ids"`
         
    }
    
    // DPCompanyRuleResponse used by Serviceability
    type DPCompanyRuleResponse struct {

        
            StatusCode float64  `json:"status_code"`
            Success bool  `json:"success"`
            Data []DpRuleResponse  `json:"data"`
         
    }
    
    // DPApplicationRuleRequest used by Serviceability
    type DPApplicationRuleRequest struct {

        
            ShippingRules []string  `json:"shipping_rules"`
         
    }
    
    // DPApplicationRuleResponse used by Serviceability
    type DPApplicationRuleResponse struct {

        
            StatusCode bool  `json:"status_code"`
            Success bool  `json:"success"`
            Data []DpRuleResponse  `json:"data"`
         
    }
    
    // SelfShipResponse used by Serviceability
    type SelfShipResponse struct {

        
            Tat float64  `json:"tat"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // ApplicationSelfShipConfig used by Serviceability
    type ApplicationSelfShipConfig struct {

        
            SelfShip SelfShipResponse  `json:"self_ship"`
         
    }
    
    // ApplicationSelfShipConfigResponse used by Serviceability
    type ApplicationSelfShipConfigResponse struct {

        
            Success bool  `json:"success"`
            Data ApplicationSelfShipConfig  `json:"data"`
            Error ServiceabilityErrorResponse  `json:"error"`
         
    }
    

    
    // GenerateReportFilters used by Finance
    type GenerateReportFilters struct {

        
            Brand []string  `json:"brand"`
            Channel []string  `json:"channel"`
            Company []string  `json:"company"`
         
    }
    
    // GenerateReportMeta used by Finance
    type GenerateReportMeta struct {

        
            Brand string  `json:"brand"`
            Channel string  `json:"channel"`
            Company string  `json:"company"`
         
    }
    
    // GenerateReportPlatform used by Finance
    type GenerateReportPlatform struct {

        
            EndDate string  `json:"end_date"`
            Filters GenerateReportFilters  `json:"filters"`
            Meta GenerateReportMeta  `json:"meta"`
            StartDate string  `json:"start_date"`
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

        
            Items [][]string  `json:"items"`
            EndDate string  `json:"end_date"`
            Headers []string  `json:"headers"`
            ItemCount float64  `json:"item_count"`
            StartDate string  `json:"start_date"`
            Page Page  `json:"page"`
         
    }
    
    // Error used by Finance
    type Error struct {

        
            Success bool  `json:"success"`
            Reason string  `json:"reason"`
         
    }
    
    // DownloadReport used by Finance
    type DownloadReport struct {

        
            StartDate string  `json:"start_date"`
            Page float64  `json:"page"`
            Pagesize float64  `json:"pagesize"`
            EndDate string  `json:"end_date"`
         
    }
    
    // DownloadReportItems used by Finance
    type DownloadReportItems struct {

        
            EndDate string  `json:"end_date"`
            Filters GenerateReportFilters  `json:"filters"`
            Meta GenerateReportMeta  `json:"meta"`
            TypeOfRequest string  `json:"type_of_request"`
            StartDate string  `json:"start_date"`
            ReportID string  `json:"report_id"`
         
    }
    
    // DownloadReportList used by Finance
    type DownloadReportList struct {

        
            Items []DownloadReportItems  `json:"items"`
            ItemCount float64  `json:"item_count"`
            Page Page  `json:"page"`
         
    }
    
    // GetEngineFilters used by Finance
    type GetEngineFilters struct {

        
            ConfigField string  `json:"config_field"`
         
    }
    
    // GetEngineData used by Finance
    type GetEngineData struct {

        
            TableName string  `json:"table_name"`
            Filters GetEngineFilters  `json:"filters"`
            Project []string  `json:"project"`
         
    }
    
    // GetEngineRequest used by Finance
    type GetEngineRequest struct {

        
            Data GetEngineData  `json:"data"`
         
    }
    
    // GetEngineResponse used by Finance
    type GetEngineResponse struct {

        
            Success bool  `json:"success"`
            Page Page  `json:"page"`
            ItemCount float64  `json:"item_count"`
            Items []map[string]interface{}  `json:"items"`
         
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

        
            Items []map[string]interface{}  `json:"items"`
            Docs []map[string]interface{}  `json:"docs"`
         
    }
    
    // GetReasonResponse used by Finance
    type GetReasonResponse struct {

        
            Success bool  `json:"success"`
            Data GetDocs  `json:"data"`
         
    }
    
    // GetReportListData used by Finance
    type GetReportListData struct {

        
            ListingEnabled bool  `json:"listing_enabled"`
            RoleName string  `json:"role_name"`
         
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

        
            PdfS3URL string  `json:"pdf_s3_url"`
            ID string  `json:"id"`
         
    }
    
    // DownloadCreditDebitNoteResponse used by Finance
    type DownloadCreditDebitNoteResponse struct {

        
            Success bool  `json:"success"`
            Data []DownloadCreditDebitNoteResponseData  `json:"data"`
         
    }
    
    // PaymentProcessPayload used by Finance
    type PaymentProcessPayload struct {

        
            Amount string  `json:"amount"`
            TotalAmount string  `json:"total_amount"`
            ModeOfPayment string  `json:"mode_of_payment"`
            InvoiceNumber string  `json:"invoice_number"`
            Meta map[string]interface{}  `json:"meta"`
            TransactionType string  `json:"transaction_type"`
            SellerID string  `json:"seller_id"`
            Platform string  `json:"platform"`
            SourceReference string  `json:"source_reference"`
            Currency string  `json:"currency"`
         
    }
    
    // PaymentProcessRequest used by Finance
    type PaymentProcessRequest struct {

        
            Data PaymentProcessPayload  `json:"data"`
         
    }
    
    // PaymentProcessResponse used by Finance
    type PaymentProcessResponse struct {

        
            Message string  `json:"message"`
            TransactionID string  `json:"transaction_id"`
            Code float64  `json:"code"`
            Meta map[string]interface{}  `json:"meta"`
            RedirectURL string  `json:"redirect_url"`
         
    }
    
    // CreditlineDataPlatformPayload used by Finance
    type CreditlineDataPlatformPayload struct {

        
            EndEnd string  `json:"end_end"`
            StartEnd string  `json:"start_end"`
            Pagesize float64  `json:"pagesize"`
            SellerID string  `json:"seller_id"`
            Page float64  `json:"page"`
         
    }
    
    // CreditlineDataPlatformRequest used by Finance
    type CreditlineDataPlatformRequest struct {

        
            Data CreditlineDataPlatformPayload  `json:"data"`
         
    }
    
    // CreditlineDataPlatformResponse used by Finance
    type CreditlineDataPlatformResponse struct {

        
            Items []map[string]interface{}  `json:"items"`
            ShowMr bool  `json:"show_mr"`
            Message string  `json:"message"`
            Code float64  `json:"code"`
            Headers []string  `json:"headers"`
            ItemCount float64  `json:"item_count"`
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

        
            Code float64  `json:"code"`
            IsCreditlineOpted bool  `json:"is_creditline_opted"`
         
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

        
            Value string  `json:"value"`
            Text string  `json:"text"`
         
    }
    
    // InvoiceTypeResponse used by Finance
    type InvoiceTypeResponse struct {

        
            Success bool  `json:"success"`
            InvoiceTypeList []InvoiceTypeResponseItems  `json:"invoice_type_list"`
            PaymentStatusList []InvoiceTypeResponseItems  `json:"payment_status_list"`
         
    }
    
    // InoviceListingPayloadDataFilters used by Finance
    type InoviceListingPayloadDataFilters struct {

        
            PaymentStatus []string  `json:"payment_status"`
            InvoiceType []string  `json:"invoice_type"`
            CompanyID []string  `json:"company_id"`
         
    }
    
    // InvoiceListingPayloadData used by Finance
    type InvoiceListingPayloadData struct {

        
            Search string  `json:"search"`
            EndDate string  `json:"end_date"`
            PageSize float64  `json:"page_size"`
            Filters InoviceListingPayloadDataFilters  `json:"filters"`
            StartDate string  `json:"start_date"`
            Page float64  `json:"page"`
         
    }
    
    // InvoiceListingRequest used by Finance
    type InvoiceListingRequest struct {

        
            Data InvoiceListingPayloadData  `json:"data"`
         
    }
    
    // InvoiceListingResponseItems used by Finance
    type InvoiceListingResponseItems struct {

        
            Amount string  `json:"amount"`
            IsDownloadable bool  `json:"is_downloadable"`
            Status string  `json:"status"`
            InvoiceNumber string  `json:"invoice_number"`
            InvoiceType string  `json:"invoice_type"`
            DueDate string  `json:"due_date"`
            InvoiceID string  `json:"invoice_id"`
            InvoiceDate string  `json:"invoice_date"`
            Period string  `json:"period"`
            Company string  `json:"company"`
         
    }
    
    // UnpaidInvoiceDataItems used by Finance
    type UnpaidInvoiceDataItems struct {

        
            TotalUnpaidAmount float64  `json:"total_unpaid_amount"`
            Currency string  `json:"currency"`
            TotalUnpaidInvoiceCount float64  `json:"total_unpaid_invoice_count"`
         
    }
    
    // InvoiceListingResponse used by Finance
    type InvoiceListingResponse struct {

        
            Items []InvoiceListingResponseItems  `json:"items"`
            UnpaidInvoiceData UnpaidInvoiceDataItems  `json:"unpaid_invoice_data"`
            Page Page  `json:"page"`
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

        
            Success bool  `json:"success"`
            Data []string  `json:"data"`
            Error []string  `json:"error"`
         
    }
    
    // AsCnRefundData used by Finance
    type AsCnRefundData struct {

        
            AffiliateID string  `json:"affiliate_id"`
            SellerID float64  `json:"seller_id"`
            ToggleEditRequired bool  `json:"toggle_edit_required"`
         
    }
    
    // AsCnRefundRequest used by Finance
    type AsCnRefundRequest struct {

        
            Data AsCnRefundData  `json:"data"`
         
    }
    
    // AsCnRefundResponseData used by Finance
    type AsCnRefundResponseData struct {

        
            IsFirstTimeUser bool  `json:"is_first_time_user"`
         
    }
    
    // AsCnRefundResponse used by Finance
    type AsCnRefundResponse struct {

        
            Success bool  `json:"success"`
            Data AsCnRefundResponseData  `json:"data"`
         
    }
    
    // CreditNoteConfigNotificationEvents used by Finance
    type CreditNoteConfigNotificationEvents struct {

        
            ExpirationReminderToCustomer float64  `json:"expiration_reminder_to_customer"`
         
    }
    
    // CreateSellerCreditNoteConfig used by Finance
    type CreateSellerCreditNoteConfig struct {

        
            SourceChannel []string  `json:"source_channel"`
            IsCnAsRefundMethod bool  `json:"is_cn_as_refund_method"`
            SalesChannelName string  `json:"sales_channel_name"`
            OrderingChannel []string  `json:"ordering_channel"`
            AffiliateID string  `json:"affiliate_id"`
            CurrencyType string  `json:"currency_type"`
            SellerID float64  `json:"seller_id"`
            SlugValues []string  `json:"slug_values"`
            NotificationEvents CreditNoteConfigNotificationEvents  `json:"notification_events"`
            Validity float64  `json:"validity"`
         
    }
    
    // CreateSellerCreditNoteConfigRequest used by Finance
    type CreateSellerCreditNoteConfigRequest struct {

        
            Data CreateSellerCreditNoteConfig  `json:"data"`
         
    }
    
    // CreateSellerCreditNoteConfigResponse used by Finance
    type CreateSellerCreditNoteConfigResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // DeleteConfig used by Finance
    type DeleteConfig struct {

        
            AffiliateID string  `json:"affiliate_id"`
            SellerID float64  `json:"seller_id"`
            SlugValues []string  `json:"slug_values"`
         
    }
    
    // DeleteConfigRequest used by Finance
    type DeleteConfigRequest struct {

        
            Data DeleteConfig  `json:"data"`
         
    }
    
    // DeleteConfigResponse used by Finance
    type DeleteConfigResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // ChannelDisplayName used by Finance
    type ChannelDisplayName struct {

        
            PlatformPos string  `json:"platform_pos"`
         
    }
    
    // ChannelDisplayNameResponse used by Finance
    type ChannelDisplayNameResponse struct {

        
            Success bool  `json:"success"`
            Data ChannelDisplayName  `json:"data"`
         
    }
    
    // CnReferenceNumber used by Finance
    type CnReferenceNumber struct {

        
            CnReferenceNumber string  `json:"cn_reference_number"`
         
    }
    
    // GetPdfUrlViewRequest used by Finance
    type GetPdfUrlViewRequest struct {

        
            Data CnReferenceNumber  `json:"data"`
         
    }
    
    // GetPdfUrlViewResponseData used by Finance
    type GetPdfUrlViewResponseData struct {

        
            S3PdfLink string  `json:"s3_pdf_link"`
            CnReferenceNumber string  `json:"cn_reference_number"`
         
    }
    
    // GetPdfUrlViewResponse used by Finance
    type GetPdfUrlViewResponse struct {

        
            Success bool  `json:"success"`
            Data GetPdfUrlViewResponseData  `json:"data"`
         
    }
    
    // CreditNoteDetailsRequest used by Finance
    type CreditNoteDetailsRequest struct {

        
            Data CnReferenceNumber  `json:"data"`
         
    }
    
    // RedemptionDetails used by Finance
    type RedemptionDetails struct {

        
            StoreID string  `json:"store_id"`
            InvoiceNumber string  `json:"invoice_number"`
            AmountDebited float64  `json:"amount_debited"`
            ShipmentID string  `json:"shipment_id"`
            OrderingChannel string  `json:"ordering_channel"`
            StaffID string  `json:"staff_id"`
            OrderID string  `json:"order_id"`
            CreatedAt string  `json:"created_at"`
         
    }
    
    // CnDetails used by Finance
    type CnDetails struct {

        
            ChannelOfIssuance string  `json:"channel_of_issuance"`
            StoreID string  `json:"store_id"`
            InvoiceNumber string  `json:"invoice_number"`
            ShipmentID string  `json:"shipment_id"`
            DateIssued string  `json:"date_issued"`
            OrderingChannel string  `json:"ordering_channel"`
            StaffID string  `json:"staff_id"`
            ExpiryDate string  `json:"expiry_date"`
            OrderID string  `json:"order_id"`
         
    }
    
    // CreditNoteDetails used by Finance
    type CreditNoteDetails struct {

        
            AvailableCnBalance float64  `json:"available_cn_balance"`
            RedemptionDetails []RedemptionDetails  `json:"redemption_details"`
            CnDetails []CnDetails  `json:"cn_details"`
            CnStatus string  `json:"cn_status"`
            CustomerMobileNumber string  `json:"customer_mobile_number"`
            CnAmount float64  `json:"cn_amount"`
            RemainingCnAmount float64  `json:"remaining_cn_amount"`
            CnReferenceNumber string  `json:"cn_reference_number"`
         
    }
    
    // CreditNoteDetailsResponse used by Finance
    type CreditNoteDetailsResponse struct {

        
            Success bool  `json:"success"`
            Data CreditNoteDetails  `json:"data"`
         
    }
    
    // GetCustomerCreditBalance used by Finance
    type GetCustomerCreditBalance struct {

        
            AffiliateID string  `json:"affiliate_id"`
            SellerID float64  `json:"seller_id"`
            CustomerMobileNumber string  `json:"customer_mobile_number"`
         
    }
    
    // GetCustomerCreditBalanceRequest used by Finance
    type GetCustomerCreditBalanceRequest struct {

        
            Data GetCustomerCreditBalance  `json:"data"`
         
    }
    
    // GetCustomerCreditBalanceResponseData used by Finance
    type GetCustomerCreditBalanceResponseData struct {

        
            TotalCreditedBalance float64  `json:"total_credited_balance"`
            CustomerMobileNumber string  `json:"customer_mobile_number"`
         
    }
    
    // GetCustomerCreditBalanceResponse used by Finance
    type GetCustomerCreditBalanceResponse struct {

        
            Success bool  `json:"success"`
            Data GetCustomerCreditBalanceResponseData  `json:"data"`
         
    }
    
    // GetCnConfigRequest used by Finance
    type GetCnConfigRequest struct {

        
            Data DeleteConfig  `json:"data"`
         
    }
    
    // GetCnConfigResponseMeta used by Finance
    type GetCnConfigResponseMeta struct {

        
            SourceChannel []string  `json:"source_channel"`
            Reason string  `json:"reason"`
         
    }
    
    // GetCnConfigResponseData used by Finance
    type GetCnConfigResponseData struct {

        
            RedemptionOrderingChannel []string  `json:"redemption_ordering_channel"`
            IsCnAsRefundMethod bool  `json:"is_cn_as_refund_method"`
            Meta GetCnConfigResponseMeta  `json:"meta"`
            AffiliateID string  `json:"affiliate_id"`
            CurrencyType string  `json:"currency_type"`
            SellerID float64  `json:"seller_id"`
            NotificationEvents CreditNoteConfigNotificationEvents  `json:"notification_events"`
            Validity float64  `json:"validity"`
         
    }
    
    // GetCnConfigResponse used by Finance
    type GetCnConfigResponse struct {

        
            Success bool  `json:"success"`
            Data GetCnConfigResponseData  `json:"data"`
         
    }
    
    // CnGenerateReportFilters used by Finance
    type CnGenerateReportFilters struct {

        
            ChannelOfIssuance []string  `json:"channel_of_issuance"`
            Utilisation []string  `json:"utilisation"`
            StoreID []float64  `json:"store_id"`
            OrderingChannel []string  `json:"ordering_channel"`
            TypesOfTransaction []string  `json:"types_of_transaction"`
            StaffID []string  `json:"staff_id"`
         
    }
    
    // CnGenerateReport used by Finance
    type CnGenerateReport struct {

        
            Search string  `json:"search"`
            EndDate string  `json:"end_date"`
            SearchType string  `json:"search_type"`
            Pagesize float64  `json:"pagesize"`
            Filters CnGenerateReportFilters  `json:"filters"`
            Meta GenerateReportFilters  `json:"meta"`
            AffiliateID string  `json:"affiliate_id"`
            StartDate string  `json:"start_date"`
            Page float64  `json:"page"`
            ReportID string  `json:"report_id"`
         
    }
    
    // GenerateReportCustomerCnRequest used by Finance
    type GenerateReportCustomerCnRequest struct {

        
            Data CnGenerateReport  `json:"data"`
         
    }
    
    // CnGenerateReportItems used by Finance
    type CnGenerateReportItems struct {

        
            TotalAmount float64  `json:"total_amount"`
            Status string  `json:"status"`
            InvoiceNumber string  `json:"invoice_number"`
            ShipmentID string  `json:"shipment_id"`
            DateIssued string  `json:"date_issued"`
            ExpiryDate string  `json:"expiry_date"`
            OrderID string  `json:"order_id"`
            CreditNoteNumber string  `json:"credit_note_number"`
         
    }
    
    // GenerateReportCustomerCnResponse used by Finance
    type GenerateReportCustomerCnResponse struct {

        
            Items []CnGenerateReportItems  `json:"items"`
            RowHeaderDisplayOrder map[string]interface{}  `json:"row_header_display_order"`
            EndDate string  `json:"end_date"`
            Headers []string  `json:"headers"`
            AllowedFilters []string  `json:"allowed_filters"`
            ItemCount float64  `json:"item_count"`
            StartDate string  `json:"start_date"`
            PrimaryHeaders []string  `json:"primary_headers"`
            Page Page  `json:"page"`
         
    }
    
    // CnDownloadReport used by Finance
    type CnDownloadReport struct {

        
            Search string  `json:"search"`
            EndDate string  `json:"end_date"`
            SearchType string  `json:"search_type"`
            Pagesize float64  `json:"pagesize"`
            Status []string  `json:"status"`
            AffiliateID string  `json:"affiliate_id"`
            StartDate string  `json:"start_date"`
            Page float64  `json:"page"`
         
    }
    
    // DownloadReportCustomerCnRequest used by Finance
    type DownloadReportCustomerCnRequest struct {

        
            Data CnDownloadReport  `json:"data"`
         
    }
    
    // DownloadReportResponseData used by Finance
    type DownloadReportResponseData struct {

        
            ReportName string  `json:"report_name"`
            Msg string  `json:"msg"`
            FullName string  `json:"full_name"`
            EndDate string  `json:"end_date"`
            Status string  `json:"status"`
            Filters map[string]interface{}  `json:"filters"`
            RequestDict map[string]interface{}  `json:"request_dict"`
            Meta map[string]interface{}  `json:"meta"`
            DownloadLink string  `json:"download_link"`
            DisplayName string  `json:"display_name"`
            RequestedBy string  `json:"requested_by"`
            StartDate string  `json:"start_date"`
            CreatedAt string  `json:"created_at"`
            ReportConfigID string  `json:"report_config_id"`
         
    }
    
    // DownloadReportCustomerCnResponse used by Finance
    type DownloadReportCustomerCnResponse struct {

        
            Data []DownloadReportResponseData  `json:"data"`
         
    }
    
    // GetReportingFilters used by Finance
    type GetReportingFilters struct {

        
            Value string  `json:"value"`
            Type string  `json:"type"`
            Options []map[string]interface{}  `json:"options"`
            Text string  `json:"text"`
         
    }
    
    // GetReportingNestedFilters used by Finance
    type GetReportingNestedFilters struct {

        
            Value string  `json:"value"`
            Options []map[string]interface{}  `json:"options"`
            Type string  `json:"type"`
            Required bool  `json:"required"`
            Text string  `json:"text"`
            PlaceholderText string  `json:"placeholder_text"`
         
    }
    
    // GetReportingFiltersResponse used by Finance
    type GetReportingFiltersResponse struct {

        
            Status GetReportingFilters  `json:"status"`
            Search GetReportingFilters  `json:"search"`
            Filters []GetReportingNestedFilters  `json:"filters"`
         
    }
    
