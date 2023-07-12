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
    
    // GlobalVariablesGetResponse used by Communication
    type GlobalVariablesGetResponse struct {

        
            ReadOnly map[string]interface{}  `json:"read_only"`
            Editable map[string]interface{}  `json:"editable"`
         
    }
    
    // GlobalVariablesPostResponse used by Communication
    type GlobalVariablesPostResponse struct {

        
            ID string  `json:"_id"`
            Category string  `json:"category"`
            Application map[string]interface{}  `json:"application"`
            GlobalVariables map[string]interface{}  `json:"global_variables"`
            CreatedAt map[string]interface{}  `json:"created_at"`
         
    }
    
    // GlobalVariablesPostResponseAdmin used by Communication
    type GlobalVariablesPostResponseAdmin struct {

        
            ID string  `json:"_id"`
            Category string  `json:"category"`
            GlobalVariables map[string]interface{}  `json:"global_variables"`
            CreatedAt map[string]interface{}  `json:"created_at"`
         
    }
    
    // GlobalVariablesReq used by Communication
    type GlobalVariablesReq struct {

        
            GlobalVariables map[string]interface{}  `json:"global_variables"`
         
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

        
            Created bool  `json:"created"`
            Success bool  `json:"success"`
            ExcludedFields []string  `json:"excluded_fields"`
            AppID string  `json:"app_id"`
            Aggregators []map[string]interface{}  `json:"aggregators"`
            DisplayFields []string  `json:"display_fields"`
         
    }
    
    // ErrorCodeDescription used by Payment
    type ErrorCodeDescription struct {

        
            Description string  `json:"description"`
            Success bool  `json:"success"`
            Code string  `json:"code"`
         
    }
    
    // PaymentGatewayConfig used by Payment
    type PaymentGatewayConfig struct {

        
            Key string  `json:"key"`
            MerchantSalt string  `json:"merchant_salt"`
            Secret string  `json:"secret"`
            IsActive bool  `json:"is_active"`
            ConfigType string  `json:"config_type"`
         
    }
    
    // PaymentGatewayConfigRequest used by Payment
    type PaymentGatewayConfigRequest struct {

        
            AggregatorName PaymentGatewayConfig  `json:"aggregator_name"`
            IsActive bool  `json:"is_active"`
            AppID string  `json:"app_id"`
         
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

        
            Large string  `json:"large"`
            Small string  `json:"small"`
         
    }
    
    // IntentApp used by Payment
    type IntentApp struct {

        
            Logos PaymentModeLogo  `json:"logos"`
            Code string  `json:"code"`
            DisplayName string  `json:"display_name"`
            PackageName string  `json:"package_name"`
         
    }
    
    // IntentAppErrorList used by Payment
    type IntentAppErrorList struct {

        
            PackageName string  `json:"package_name"`
            Code string  `json:"code"`
         
    }
    
    // PaymentModeList used by Payment
    type PaymentModeList struct {

        
            CardNumber string  `json:"card_number"`
            CardIsin string  `json:"card_isin"`
            Timeout float64  `json:"timeout"`
            CardReference string  `json:"card_reference"`
            IntentApp []IntentApp  `json:"intent_app"`
            Code string  `json:"code"`
            LogoURL PaymentModeLogo  `json:"logo_url"`
            RemainingLimit float64  `json:"remaining_limit"`
            CardToken string  `json:"card_token"`
            CardIssuer string  `json:"card_issuer"`
            CardFingerprint string  `json:"card_fingerprint"`
            Nickname string  `json:"nickname"`
            ExpYear float64  `json:"exp_year"`
            MerchantCode string  `json:"merchant_code"`
            RetryCount float64  `json:"retry_count"`
            CardBrand string  `json:"card_brand"`
            CodLimitPerOrder float64  `json:"cod_limit_per_order"`
            IntentAppErrorDictList []IntentAppErrorList  `json:"intent_app_error_dict_list"`
            Expired bool  `json:"expired"`
            DisplayName string  `json:"display_name"`
            IntentAppErrorList []string  `json:"intent_app_error_list"`
            AggregatorName string  `json:"aggregator_name"`
            CardBrandImage string  `json:"card_brand_image"`
            CardID string  `json:"card_id"`
            CodLimit float64  `json:"cod_limit"`
            FyndVpa string  `json:"fynd_vpa"`
            CompliantWithTokenisationGuidelines bool  `json:"compliant_with_tokenisation_guidelines"`
            CardName string  `json:"card_name"`
            DisplayPriority float64  `json:"display_priority"`
            Name string  `json:"name"`
            IntentFlow bool  `json:"intent_flow"`
            ExpMonth float64  `json:"exp_month"`
            CardType string  `json:"card_type"`
         
    }
    
    // RootPaymentMode used by Payment
    type RootPaymentMode struct {

        
            DisplayPriority float64  `json:"display_priority"`
            List []PaymentModeList  `json:"list"`
            Name string  `json:"name"`
            AddCardEnabled bool  `json:"add_card_enabled"`
            AggregatorName string  `json:"aggregator_name"`
            SaveCard bool  `json:"save_card"`
            DisplayName string  `json:"display_name"`
            IsPayByCardPl bool  `json:"is_pay_by_card_pl"`
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
    
    // PayoutsResponse used by Payment
    type PayoutsResponse struct {

        
            MoreAttributes map[string]interface{}  `json:"more_attributes"`
            PayoutsAggregators []map[string]interface{}  `json:"payouts_aggregators"`
            IsDefault bool  `json:"is_default"`
            Customers map[string]interface{}  `json:"customers"`
            UniqueTransferNo map[string]interface{}  `json:"unique_transfer_no"`
            IsActive bool  `json:"is_active"`
            TransferType string  `json:"transfer_type"`
         
    }
    
    // PayoutBankDetails used by Payment
    type PayoutBankDetails struct {

        
            City string  `json:"city"`
            Country string  `json:"country"`
            AccountType string  `json:"account_type"`
            State string  `json:"state"`
            BankName string  `json:"bank_name"`
            Pincode float64  `json:"pincode"`
            IfscCode string  `json:"ifsc_code"`
            AccountNo string  `json:"account_no"`
            BranchName string  `json:"branch_name"`
            AccountHolder string  `json:"account_holder"`
         
    }
    
    // PayoutRequest used by Payment
    type PayoutRequest struct {

        
            BankDetails PayoutBankDetails  `json:"bank_details"`
            Aggregator string  `json:"aggregator"`
            Users map[string]interface{}  `json:"users"`
            UniqueExternalID string  `json:"unique_external_id"`
            IsActive bool  `json:"is_active"`
            TransferType string  `json:"transfer_type"`
         
    }
    
    // PayoutResponse used by Payment
    type PayoutResponse struct {

        
            BankDetails map[string]interface{}  `json:"bank_details"`
            Aggregator string  `json:"aggregator"`
            Created bool  `json:"created"`
            PaymentStatus string  `json:"payment_status"`
            Users map[string]interface{}  `json:"users"`
            Payouts map[string]interface{}  `json:"payouts"`
            Success bool  `json:"success"`
            UniqueTransferNo string  `json:"unique_transfer_no"`
            IsActive bool  `json:"is_active"`
            TransferType string  `json:"transfer_type"`
         
    }
    
    // UpdatePayoutResponse used by Payment
    type UpdatePayoutResponse struct {

        
            IsDefault bool  `json:"is_default"`
            Success bool  `json:"success"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // UpdatePayoutRequest used by Payment
    type UpdatePayoutRequest struct {

        
            UniqueExternalID string  `json:"unique_external_id"`
            IsActive bool  `json:"is_active"`
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

        
            Success bool  `json:"success"`
            Aggregator string  `json:"aggregator"`
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

        
            Success bool  `json:"success"`
            IsVerifiedFlag bool  `json:"is_verified_flag"`
            Message string  `json:"message"`
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

        
            BankName string  `json:"bank_name"`
            IfscCode string  `json:"ifsc_code"`
            AccountNo string  `json:"account_no"`
            BranchName string  `json:"branch_name"`
            AccountHolder string  `json:"account_holder"`
         
    }
    
    // AddBeneficiaryDetailsOTPRequest used by Payment
    type AddBeneficiaryDetailsOTPRequest struct {

        
            Details BankDetailsForOTP  `json:"details"`
            OrderID string  `json:"order_id"`
         
    }
    
    // IfscCodeResponse used by Payment
    type IfscCodeResponse struct {

        
            BankName string  `json:"bank_name"`
            Success bool  `json:"success"`
            BranchName string  `json:"branch_name"`
         
    }
    
    // OrderBeneficiaryDetails used by Payment
    type OrderBeneficiaryDetails struct {

        
            CreatedOn string  `json:"created_on"`
            Mobile string  `json:"mobile"`
            BankName string  `json:"bank_name"`
            ID float64  `json:"id"`
            Title string  `json:"title"`
            BeneficiaryID string  `json:"beneficiary_id"`
            Subtitle string  `json:"subtitle"`
            IsActive bool  `json:"is_active"`
            TransferMode string  `json:"transfer_mode"`
            DisplayName string  `json:"display_name"`
            BranchName string  `json:"branch_name"`
            Comment string  `json:"comment"`
            DelightsUserName string  `json:"delights_user_name"`
            Email string  `json:"email"`
            ModifiedOn string  `json:"modified_on"`
            AccountNo string  `json:"account_no"`
            IfscCode string  `json:"ifsc_code"`
            Address string  `json:"address"`
            AccountHolder string  `json:"account_holder"`
         
    }
    
    // OrderBeneficiaryResponse used by Payment
    type OrderBeneficiaryResponse struct {

        
            ShowBeneficiaryDetails bool  `json:"show_beneficiary_details"`
            Beneficiaries []OrderBeneficiaryDetails  `json:"beneficiaries"`
         
    }
    
    // MultiTenderPaymentMeta used by Payment
    type MultiTenderPaymentMeta struct {

        
            PaymentGateway string  `json:"payment_gateway"`
            OrderID string  `json:"order_id"`
            CurrentStatus string  `json:"current_status"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            PaymentID string  `json:"payment_id"`
         
    }
    
    // MultiTenderPaymentMethod used by Payment
    type MultiTenderPaymentMethod struct {

        
            Amount float64  `json:"amount"`
            Meta MultiTenderPaymentMeta  `json:"meta"`
            Name string  `json:"name"`
            Mode string  `json:"mode"`
         
    }
    
    // PaymentConfirmationRequest used by Payment
    type PaymentConfirmationRequest struct {

        
            OrderID string  `json:"order_id"`
            PaymentMethods []MultiTenderPaymentMethod  `json:"payment_methods"`
         
    }
    
    // PaymentConfirmationResponse used by Payment
    type PaymentConfirmationResponse struct {

        
            Success bool  `json:"success"`
            OrderID string  `json:"order_id"`
            Message string  `json:"message"`
         
    }
    
    // CODdata used by Payment
    type CODdata struct {

        
            UserID string  `json:"user_id"`
            Limit float64  `json:"limit"`
            RemainingLimit float64  `json:"remaining_limit"`
            IsActive bool  `json:"is_active"`
            Usages float64  `json:"usages"`
         
    }
    
    // GetUserCODLimitResponse used by Payment
    type GetUserCODLimitResponse struct {

        
            Success bool  `json:"success"`
            UserCodData CODdata  `json:"user_cod_data"`
         
    }
    
    // SetCODForUserRequest used by Payment
    type SetCODForUserRequest struct {

        
            MerchantUserID string  `json:"merchant_user_id"`
            IsActive bool  `json:"is_active"`
            Mobileno string  `json:"mobileno"`
         
    }
    
    // SetCODOptionResponse used by Payment
    type SetCODOptionResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // EdcModelData used by Payment
    type EdcModelData struct {

        
            Models []string  `json:"models"`
            AggregatorID float64  `json:"aggregator_id"`
            Aggregator string  `json:"aggregator"`
         
    }
    
    // EdcAggregatorAndModelListResponse used by Payment
    type EdcAggregatorAndModelListResponse struct {

        
            Success bool  `json:"success"`
            Data []EdcModelData  `json:"data"`
         
    }
    
    // StatisticsData used by Payment
    type StatisticsData struct {

        
            InactiveDeviceCount float64  `json:"inactive_device_count"`
            ActiveDeviceCount float64  `json:"active_device_count"`
         
    }
    
    // EdcDeviceStatsResponse used by Payment
    type EdcDeviceStatsResponse struct {

        
            Success bool  `json:"success"`
            Statistics StatisticsData  `json:"statistics"`
         
    }
    
    // EdcAddRequest used by Payment
    type EdcAddRequest struct {

        
            EdcDeviceSerialNo string  `json:"edc_device_serial_no"`
            StoreID float64  `json:"store_id"`
            EdcModel string  `json:"edc_model"`
            AggregatorID float64  `json:"aggregator_id"`
            TerminalSerialNo string  `json:"terminal_serial_no"`
            DeviceTag string  `json:"device_tag"`
         
    }
    
    // EdcDevice used by Payment
    type EdcDevice struct {

        
            ApplicationID string  `json:"application_id"`
            MerchantStorePosCode string  `json:"merchant_store_pos_code"`
            AggregatorName string  `json:"aggregator_name"`
            AggregatorID float64  `json:"aggregator_id"`
            IsActive bool  `json:"is_active"`
            StoreID float64  `json:"store_id"`
            EdcModel string  `json:"edc_model"`
            TerminalUniqueIdentifier string  `json:"terminal_unique_identifier"`
            EdcDeviceSerialNo string  `json:"edc_device_serial_no"`
            TerminalSerialNo string  `json:"terminal_serial_no"`
            DeviceTag string  `json:"device_tag"`
         
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

        
            MerchantStorePosCode string  `json:"merchant_store_pos_code"`
            AggregatorID float64  `json:"aggregator_id"`
            StoreID float64  `json:"store_id"`
            EdcModel string  `json:"edc_model"`
            IsActive bool  `json:"is_active"`
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

        
            Success bool  `json:"success"`
            Items []EdcDevice  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // PaymentInitializationRequest used by Payment
    type PaymentInitializationRequest struct {

        
            DeviceID string  `json:"device_id"`
            RazorpayPaymentID string  `json:"razorpay_payment_id"`
            Timeout float64  `json:"timeout"`
            Aggregator string  `json:"aggregator"`
            OrderID string  `json:"order_id"`
            Currency string  `json:"currency"`
            Contact string  `json:"contact"`
            CustomerID string  `json:"customer_id"`
            Amount float64  `json:"amount"`
            MerchantOrderID string  `json:"merchant_order_id"`
            Vpa string  `json:"vpa"`
            Email string  `json:"email"`
            Method string  `json:"method"`
         
    }
    
    // PaymentInitializationResponse used by Payment
    type PaymentInitializationResponse struct {

        
            DeviceID string  `json:"device_id"`
            RazorpayPaymentID string  `json:"razorpay_payment_id"`
            Timeout float64  `json:"timeout"`
            Aggregator string  `json:"aggregator"`
            VirtualID string  `json:"virtual_id"`
            Currency string  `json:"currency"`
            UpiPollURL string  `json:"upi_poll_url"`
            MerchantOrderID string  `json:"merchant_order_id"`
            Amount float64  `json:"amount"`
            Success bool  `json:"success"`
            CustomerID string  `json:"customer_id"`
            Vpa string  `json:"vpa"`
            PollingURL string  `json:"polling_url"`
            BqrImage string  `json:"bqr_image"`
            Method string  `json:"method"`
            AggregatorOrderID string  `json:"aggregator_order_id"`
            Status string  `json:"status"`
         
    }
    
    // PaymentStatusUpdateRequest used by Payment
    type PaymentStatusUpdateRequest struct {

        
            DeviceID string  `json:"device_id"`
            Aggregator string  `json:"aggregator"`
            OrderID string  `json:"order_id"`
            Currency string  `json:"currency"`
            Contact string  `json:"contact"`
            CustomerID string  `json:"customer_id"`
            Amount float64  `json:"amount"`
            MerchantOrderID string  `json:"merchant_order_id"`
            Vpa string  `json:"vpa"`
            Email string  `json:"email"`
            Method string  `json:"method"`
            MerchantTransactionID string  `json:"merchant_transaction_id"`
            Status string  `json:"status"`
         
    }
    
    // PaymentStatusUpdateResponse used by Payment
    type PaymentStatusUpdateResponse struct {

        
            RedirectURL string  `json:"redirect_url"`
            AggregatorName string  `json:"aggregator_name"`
            Success bool  `json:"success"`
            Retry bool  `json:"retry"`
            Status string  `json:"status"`
         
    }
    
    // ResendOrCancelPaymentRequest used by Payment
    type ResendOrCancelPaymentRequest struct {

        
            DeviceID string  `json:"device_id"`
            OrderID string  `json:"order_id"`
            RequestType string  `json:"request_type"`
         
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

        
            CompanyID string  `json:"company_id"`
            ApplicationID string  `json:"application_id"`
            PaymentGateway string  `json:"payment_gateway"`
            CreatedOn string  `json:"created_on"`
            AllStatus []string  `json:"all_status"`
            PaymentMode string  `json:"payment_mode"`
            ID string  `json:"id"`
            Currency string  `json:"currency"`
            CurrentStatus string  `json:"current_status"`
            UserObject map[string]interface{}  `json:"user_object"`
            RefundObject map[string]interface{}  `json:"refund_object"`
            RefundedBy string  `json:"refunded_by"`
            AmountInPaisa string  `json:"amount_in_paisa"`
            ModifiedOn string  `json:"modified_on"`
            PaymentModeIdentifier string  `json:"payment_mode_identifier"`
            AggregatorPaymentObject map[string]interface{}  `json:"aggregator_payment_object"`
            CollectedBy string  `json:"collected_by"`
            PaymentID string  `json:"payment_id"`
         
    }
    
    // PaymentStatusObject used by Payment
    type PaymentStatusObject struct {

        
            PaymentObjectList []PaymentObjectListSerializer  `json:"payment_object_list"`
            MerchantOrderID string  `json:"merchant_order_id"`
         
    }
    
    // PaymentStatusBulkHandlerResponse used by Payment
    type PaymentStatusBulkHandlerResponse struct {

        
            Error string  `json:"error"`
            Count float64  `json:"count"`
            Success string  `json:"success"`
            Data []PaymentStatusObject  `json:"data"`
            Status float64  `json:"status"`
         
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

        
            Aggregator string  `json:"aggregator"`
            PaymentMode string  `json:"payment_mode"`
            PaymentModeIdentifier string  `json:"payment_mode_identifier"`
            CurrentStatus string  `json:"current_status"`
            AggregatorTransactionID string  `json:"aggregator_transaction_id"`
            MerchantOrderID string  `json:"merchant_order_id"`
            Amount float64  `json:"amount"`
            OutstandingDetailsID float64  `json:"outstanding_details_id"`
            FwdShipmentID string  `json:"fwd_shipment_id"`
            AggregatorOrderID string  `json:"aggregator_order_id"`
         
    }
    
    // RepaymentDetailsSerialiserPayAll used by Payment
    type RepaymentDetailsSerialiserPayAll struct {

        
            TotalAmount float64  `json:"total_amount"`
            ExtensionOrderID string  `json:"extension_order_id"`
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

        
            Aggregator string  `json:"aggregator"`
            UserID string  `json:"user_id"`
            AppID string  `json:"app_id"`
            CreditLineID string  `json:"credit_line_id"`
            Status string  `json:"status"`
         
    }
    
    // MerchantOnBoardingResponse used by Payment
    type MerchantOnBoardingResponse struct {

        
            Success bool  `json:"success"`
            Data map[string]interface{}  `json:"data"`
         
    }
    
    // ValidateCustomerRequest used by Payment
    type ValidateCustomerRequest struct {

        
            Aggregator string  `json:"aggregator"`
            TransactionAmountInPaise float64  `json:"transaction_amount_in_paise"`
            OrderItems []map[string]interface{}  `json:"order_items"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
            PhoneNumber string  `json:"phone_number"`
            MerchantParams map[string]interface{}  `json:"merchant_params"`
            Payload string  `json:"payload"`
         
    }
    
    // ValidateCustomerResponse used by Payment
    type ValidateCustomerResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
            Data map[string]interface{}  `json:"data"`
         
    }
    
    // GetPaymentLinkResponse used by Payment
    type GetPaymentLinkResponse struct {

        
            MerchantName string  `json:"merchant_name"`
            StatusCode float64  `json:"status_code"`
            PaymentLinkURL string  `json:"payment_link_url"`
            Amount float64  `json:"amount"`
            Success bool  `json:"success"`
            PaymentLinkCurrentStatus string  `json:"payment_link_current_status"`
            ExternalOrderID string  `json:"external_order_id"`
            PollingTimeout float64  `json:"polling_timeout"`
            Message string  `json:"message"`
         
    }
    
    // ErrorDescription used by Payment
    type ErrorDescription struct {

        
            PaymentTransactionID string  `json:"payment_transaction_id"`
            MerchantName string  `json:"merchant_name"`
            MerchantOrderID string  `json:"merchant_order_id"`
            Amount float64  `json:"amount"`
            Expired bool  `json:"expired"`
            Msg string  `json:"msg"`
            Cancelled bool  `json:"cancelled"`
            InvalidID bool  `json:"invalid_id"`
         
    }
    
    // ErrorResponse used by Payment
    type ErrorResponse struct {

        
            Success bool  `json:"success"`
            StatusCode float64  `json:"status_code"`
            Message string  `json:"message"`
            Error ErrorDescription  `json:"error"`
         
    }
    
    // CreatePaymentLinkMeta used by Payment
    type CreatePaymentLinkMeta struct {

        
            AssignCardID string  `json:"assign_card_id"`
            CheckoutMode string  `json:"checkout_mode"`
            CartID string  `json:"cart_id"`
            Amount string  `json:"amount"`
            Pincode string  `json:"pincode"`
         
    }
    
    // CreatePaymentLinkRequest used by Payment
    type CreatePaymentLinkRequest struct {

        
            Description string  `json:"description"`
            Meta CreatePaymentLinkMeta  `json:"meta"`
            MobileNumber string  `json:"mobile_number"`
            Amount float64  `json:"amount"`
            Email string  `json:"email"`
            ExternalOrderID string  `json:"external_order_id"`
         
    }
    
    // CreatePaymentLinkResponse used by Payment
    type CreatePaymentLinkResponse struct {

        
            PaymentLinkID string  `json:"payment_link_id"`
            StatusCode float64  `json:"status_code"`
            PaymentLinkURL string  `json:"payment_link_url"`
            Success bool  `json:"success"`
            PollingTimeout float64  `json:"polling_timeout"`
            Message string  `json:"message"`
         
    }
    
    // PollingPaymentLinkResponse used by Payment
    type PollingPaymentLinkResponse struct {

        
            PaymentLinkID string  `json:"payment_link_id"`
            OrderID string  `json:"order_id"`
            RedirectURL string  `json:"redirect_url"`
            AggregatorName string  `json:"aggregator_name"`
            StatusCode float64  `json:"status_code"`
            Status string  `json:"status"`
            Amount float64  `json:"amount"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
            HttpStatus float64  `json:"http_status"`
         
    }
    
    // CancelOrResendPaymentLinkRequest used by Payment
    type CancelOrResendPaymentLinkRequest struct {

        
            PaymentLinkID string  `json:"payment_link_id"`
         
    }
    
    // ResendPaymentLinkResponse used by Payment
    type ResendPaymentLinkResponse struct {

        
            Success bool  `json:"success"`
            PollingTimeout float64  `json:"polling_timeout"`
            StatusCode float64  `json:"status_code"`
            Message string  `json:"message"`
         
    }
    
    // CancelPaymentLinkResponse used by Payment
    type CancelPaymentLinkResponse struct {

        
            Success bool  `json:"success"`
            StatusCode float64  `json:"status_code"`
            Message string  `json:"message"`
         
    }
    
    // Code used by Payment
    type Code struct {

        
            Code string  `json:"code"`
            MerchantCode string  `json:"merchant_code"`
            Name string  `json:"name"`
         
    }
    
    // PaymentCode used by Payment
    type PaymentCode struct {

        
            Networks string  `json:"networks"`
            Name string  `json:"name"`
            Types string  `json:"types"`
            Codes Code  `json:"codes"`
         
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
    

    
    // Formatted used by Order
    type Formatted struct {

        
            Min string  `json:"min"`
            Max string  `json:"max"`
         
    }
    
    // ShipmentTags used by Order
    type ShipmentTags struct {

        
            EntityType string  `json:"entity_type"`
            Slug string  `json:"slug"`
            DisplayText string  `json:"display_text"`
         
    }
    
    // ShipmentTimeStamp used by Order
    type ShipmentTimeStamp struct {

        
            Min float64  `json:"min"`
            Max float64  `json:"max"`
         
    }
    
    // LockData used by Order
    type LockData struct {

        
            LockMessage string  `json:"lock_message"`
            Locked bool  `json:"locked"`
            Mto bool  `json:"mto"`
         
    }
    
    // ShipmentItemMeta used by Order
    type ShipmentItemMeta struct {

        
            DpOptions map[string]interface{}  `json:"dp_options"`
            EwaybillInfo map[string]interface{}  `json:"ewaybill_info"`
            ShipmentChargeableWeight float64  `json:"shipment_chargeable_weight"`
            DpSortKey string  `json:"dp_sort_key"`
            AssignDpFromSb bool  `json:"assign_dp_from_sb"`
            ActivityComment string  `json:"activity_comment"`
            Formatted Formatted  `json:"formatted"`
            Tags []map[string]interface{}  `json:"tags"`
            ShippingZone string  `json:"shipping_zone"`
            DebugInfo map[string]interface{}  `json:"debug_info"`
            FulfilmentPriorityText string  `json:"fulfilment_priority_text"`
            ShipmentTags []ShipmentTags  `json:"shipment_tags"`
            ShipmentVolumetricWeight float64  `json:"shipment_volumetric_weight"`
            IsInternational bool  `json:"is_international"`
            OrderType string  `json:"order_type"`
            External map[string]interface{}  `json:"external"`
            StoreInvoiceUpdatedDate string  `json:"store_invoice_updated_date"`
            Timestamp ShipmentTimeStamp  `json:"timestamp"`
            Weight float64  `json:"weight"`
            SameStoreAvailable bool  `json:"same_store_available"`
            PackagingName string  `json:"packaging_name"`
            LockData LockData  `json:"lock_data"`
            AutoTriggerDpAssignmentAcf bool  `json:"auto_trigger_dp_assignment_acf"`
            ShipmentWeight float64  `json:"shipment_weight"`
            Sla float64  `json:"sla"`
            ParentDpID string  `json:"parent_dp_id"`
            ExistingDpList []string  `json:"existing_dp_list"`
            BagWeight map[string]interface{}  `json:"bag_weight"`
            PdfMedia []map[string]interface{}  `json:"pdf_media"`
         
    }
    
    // ShipmentStatus used by Order
    type ShipmentStatus struct {

        
            Meta map[string]interface{}  `json:"meta"`
            CurrentShipmentStatus string  `json:"current_shipment_status"`
            CreatedAt string  `json:"created_at"`
            ShipmentID string  `json:"shipment_id"`
            Title string  `json:"title"`
            ShipmentStatusID float64  `json:"shipment_status_id"`
            BagList []string  `json:"bag_list"`
            Status string  `json:"status"`
            StatusCreatedAt string  `json:"status_created_at"`
         
    }
    
    // ShipmentItemFulFillingStore used by Order
    type ShipmentItemFulFillingStore struct {

        
            Pincode string  `json:"pincode"`
            State string  `json:"state"`
            Code string  `json:"code"`
            Meta map[string]interface{}  `json:"meta"`
            Phone string  `json:"phone"`
            City string  `json:"city"`
            StoreEmail string  `json:"store_email"`
            Address string  `json:"address"`
            ID float64  `json:"id"`
            BrandStoreTags string  `json:"brand_store_tags"`
            LocationType string  `json:"location_type"`
            Name string  `json:"name"`
         
    }
    
    // UserDataInfo used by Order
    type UserDataInfo struct {

        
            UID float64  `json:"uid"`
            Gender string  `json:"gender"`
            LastName string  `json:"last_name"`
            Mobile string  `json:"mobile"`
            Email string  `json:"email"`
            AvisUserID string  `json:"avis_user_id"`
            FirstName string  `json:"first_name"`
            IsAnonymousUser bool  `json:"is_anonymous_user"`
            Name string  `json:"name"`
         
    }
    
    // ShipmentListingBrand used by Order
    type ShipmentListingBrand struct {

        
            LogoBase64 string  `json:"logo_base64"`
            CreatedOn string  `json:"created_on"`
            Logo string  `json:"logo"`
            Name string  `json:"name"`
         
    }
    
    // BagStateMapper used by Order
    type BagStateMapper struct {

        
            AppFacing bool  `json:"app_facing"`
            ID float64  `json:"id"`
            StateType string  `json:"state_type"`
            NotifyCustomer bool  `json:"notify_customer"`
            AppStateName string  `json:"app_state_name"`
            AppDisplayName string  `json:"app_display_name"`
            JourneyType string  `json:"journey_type"`
            DisplayName string  `json:"display_name"`
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
         
    }
    
    // BagStatusHistory used by Order
    type BagStatusHistory struct {

        
            StateID float64  `json:"state_id"`
            KafkaSync bool  `json:"kafka_sync"`
            Reasons []map[string]interface{}  `json:"reasons"`
            BshID float64  `json:"bsh_id"`
            StoreID float64  `json:"store_id"`
            ShipmentID string  `json:"shipment_id"`
            CreatedAt string  `json:"created_at"`
            DeliveryPartnerID float64  `json:"delivery_partner_id"`
            Forward bool  `json:"forward"`
            DeliveryAwbNumber string  `json:"delivery_awb_number"`
            StateType string  `json:"state_type"`
            BagStateMapper BagStateMapper  `json:"bag_state_mapper"`
            BagID float64  `json:"bag_id"`
            AppDisplayName string  `json:"app_display_name"`
            Status string  `json:"status"`
            DisplayName string  `json:"display_name"`
            UpdatedAt string  `json:"updated_at"`
         
    }
    
    // ReplacementDetails used by Order
    type ReplacementDetails struct {

        
            ReplacementType string  `json:"replacement_type"`
            OriginalAffiliateOrderID string  `json:"original_affiliate_order_id"`
         
    }
    
    // AffiliateMeta used by Order
    type AffiliateMeta struct {

        
            DueDate string  `json:"due_date"`
            IsPriority bool  `json:"is_priority"`
            ChannelOrderID string  `json:"channel_order_id"`
            SizeLevelTotalQty float64  `json:"size_level_total_qty"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            Quantity float64  `json:"quantity"`
            ChannelShipmentID string  `json:"channel_shipment_id"`
            ReplacementDetails ReplacementDetails  `json:"replacement_details"`
            MarketplaceInvoiceID string  `json:"marketplace_invoice_id"`
            CouponCode string  `json:"coupon_code"`
            BoxType string  `json:"box_type"`
            EmployeeDiscount float64  `json:"employee_discount"`
            OrderItemID string  `json:"order_item_id"`
         
    }
    
    // AffiliateBagDetails used by Order
    type AffiliateBagDetails struct {

        
            AffiliateMeta AffiliateMeta  `json:"affiliate_meta"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
            EmployeeDiscount float64  `json:"employee_discount"`
         
    }
    
    // Weight used by Order
    type Weight struct {

        
            IsDefault bool  `json:"is_default"`
            Shipping float64  `json:"shipping"`
            Unit string  `json:"unit"`
         
    }
    
    // ReturnConfig used by Order
    type ReturnConfig struct {

        
            Time float64  `json:"time"`
            Returnable bool  `json:"returnable"`
            Unit string  `json:"unit"`
         
    }
    
    // Dimensions used by Order
    type Dimensions struct {

        
            Length float64  `json:"length"`
            Height float64  `json:"height"`
            Width float64  `json:"width"`
            Unit string  `json:"unit"`
            IsDefault bool  `json:"is_default"`
         
    }
    
    // Article used by Order
    type Article struct {

        
            EspModified bool  `json:"esp_modified"`
            SellerIdentifier string  `json:"seller_identifier"`
            Code string  `json:"code"`
            UID string  `json:"uid"`
            ChildDetails map[string]interface{}  `json:"child_details"`
            IsSet bool  `json:"is_set"`
            Weight Weight  `json:"weight"`
            RawMeta string  `json:"raw_meta"`
            ID string  `json:"_id"`
            ReturnConfig ReturnConfig  `json:"return_config"`
            Dimensions Dimensions  `json:"dimensions"`
            Currency map[string]interface{}  `json:"currency"`
            ASet map[string]interface{}  `json:"a_set"`
            Identifiers map[string]interface{}  `json:"identifiers"`
            Size string  `json:"size"`
         
    }
    
    // Dates used by Order
    type Dates struct {

        
            DeliveryDate string  `json:"delivery_date"`
            OrderCreated string  `json:"order_created"`
         
    }
    
    // GSTDetailsData used by Order
    type GSTDetailsData struct {

        
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            HsnCodeID string  `json:"hsn_code_id"`
            SgstTaxPercentage float64  `json:"sgst_tax_percentage"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            GstTag string  `json:"gst_tag"`
            IsDefaultHsnCode bool  `json:"is_default_hsn_code"`
            IgstGstFee string  `json:"igst_gst_fee"`
            CgstTaxPercentage float64  `json:"cgst_tax_percentage"`
            GstFee float64  `json:"gst_fee"`
            SgstGstFee string  `json:"sgst_gst_fee"`
            ValueOfGood float64  `json:"value_of_good"`
            CgstGstFee string  `json:"cgst_gst_fee"`
            IgstTaxPercentage float64  `json:"igst_tax_percentage"`
            HsnCode string  `json:"hsn_code"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            GstinCode string  `json:"gstin_code"`
         
    }
    
    // BagReturnableCancelableStatus used by Order
    type BagReturnableCancelableStatus struct {

        
            IsCustomerReturnAllowed bool  `json:"is_customer_return_allowed"`
            IsReturnable bool  `json:"is_returnable"`
            CanBeCancelled bool  `json:"can_be_cancelled"`
            EnableTracking bool  `json:"enable_tracking"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // PlatformArticleAttributes used by Order
    type PlatformArticleAttributes struct {

        
            Currency string  `json:"currency"`
         
    }
    
    // PlatformItem used by Order
    type PlatformItem struct {

        
            Meta map[string]interface{}  `json:"meta"`
            Attributes PlatformArticleAttributes  `json:"attributes"`
            DepartmentID float64  `json:"department_id"`
            Brand string  `json:"brand"`
            CanReturn bool  `json:"can_return"`
            BrandID float64  `json:"brand_id"`
            ID float64  `json:"id"`
            Color string  `json:"color"`
            L1Category []string  `json:"l1_category"`
            CanCancel bool  `json:"can_cancel"`
            Image []string  `json:"image"`
            BranchURL string  `json:"branch_url"`
            Code string  `json:"code"`
            L3Category float64  `json:"l3_category"`
            L2Category []string  `json:"l2_category"`
            SlugKey string  `json:"slug_key"`
            Images []string  `json:"images"`
            LastUpdatedAt string  `json:"last_updated_at"`
            L3CategoryName string  `json:"l3_category_name"`
            Size string  `json:"size"`
            Name string  `json:"name"`
         
    }
    
    // Prices used by Order
    type Prices struct {

        
            AmountPaid float64  `json:"amount_paid"`
            GiftPrice float64  `json:"gift_price"`
            FyndCredits float64  `json:"fynd_credits"`
            RefundCredit float64  `json:"refund_credit"`
            CouponValue float64  `json:"coupon_value"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            TransferPrice float64  `json:"transfer_price"`
            PmPriceSplit float64  `json:"pm_price_split"`
            DeliveryCharge float64  `json:"delivery_charge"`
            PriceEffective float64  `json:"price_effective"`
            RefundAmount float64  `json:"refund_amount"`
            AmountPaidRoundoff float64  `json:"amount_paid_roundoff"`
            Cashback float64  `json:"cashback"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            Discount float64  `json:"discount"`
            PriceMarked float64  `json:"price_marked"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            CashbackApplied float64  `json:"cashback_applied"`
            ValueOfGood float64  `json:"value_of_good"`
            CodCharges float64  `json:"cod_charges"`
            PromotionEffectiveDiscount float64  `json:"promotion_effective_discount"`
         
    }
    
    // Identifier used by Order
    type Identifier struct {

        
            Isbn string  `json:"isbn"`
            Alu string  `json:"alu"`
            SkuCode string  `json:"sku_code"`
            Upc string  `json:"upc"`
            Ean string  `json:"ean"`
         
    }
    
    // FinancialBreakup used by Order
    type FinancialBreakup struct {

        
            AmountPaid float64  `json:"amount_paid"`
            GstTag string  `json:"gst_tag"`
            FyndCredits float64  `json:"fynd_credits"`
            RefundCredit float64  `json:"refund_credit"`
            CouponValue float64  `json:"coupon_value"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            TotalUnits float64  `json:"total_units"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            TransferPrice float64  `json:"transfer_price"`
            DeliveryCharge float64  `json:"delivery_charge"`
            CodCharges float64  `json:"cod_charges"`
            PriceEffective float64  `json:"price_effective"`
            AddedToFyndCash bool  `json:"added_to_fynd_cash"`
            AmountPaidRoundoff float64  `json:"amount_paid_roundoff"`
            Cashback float64  `json:"cashback"`
            ItemName string  `json:"item_name"`
            Size string  `json:"size"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            HsnCode string  `json:"hsn_code"`
            PriceMarked float64  `json:"price_marked"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            GstFee float64  `json:"gst_fee"`
            CashbackApplied float64  `json:"cashback_applied"`
            ValueOfGood float64  `json:"value_of_good"`
            Identifiers Identifier  `json:"identifiers"`
            Discount float64  `json:"discount"`
            PromotionEffectiveDiscount float64  `json:"promotion_effective_discount"`
         
    }
    
    // BagUnit used by Order
    type BagUnit struct {

        
            Meta map[string]interface{}  `json:"meta"`
            LineNumber float64  `json:"line_number"`
            Brand ShipmentListingBrand  `json:"brand"`
            CanReturn bool  `json:"can_return"`
            BagStatus []BagStatusHistory  `json:"bag_status"`
            CanCancel bool  `json:"can_cancel"`
            CurrentStatus BagStatusHistory  `json:"current_status"`
            DisplayName string  `json:"display_name"`
            AffiliateBagDetails AffiliateBagDetails  `json:"affiliate_bag_details"`
            BagExpiryDate string  `json:"bag_expiry_date"`
            Article Article  `json:"article"`
            Dates Dates  `json:"dates"`
            Gst GSTDetailsData  `json:"gst"`
            CurrentOperationalStatus BagStatusHistory  `json:"current_operational_status"`
            EntityType string  `json:"entity_type"`
            ProductQuantity float64  `json:"product_quantity"`
            Status BagReturnableCancelableStatus  `json:"status"`
            Item PlatformItem  `json:"item"`
            Reasons []map[string]interface{}  `json:"reasons"`
            BagType string  `json:"bag_type"`
            Prices Prices  `json:"prices"`
            BagID float64  `json:"bag_id"`
            FinancialBreakup []FinancialBreakup  `json:"financial_breakup"`
            Size string  `json:"size"`
         
    }
    
    // ShipmentListingChannel used by Order
    type ShipmentListingChannel struct {

        
            Name string  `json:"name"`
            IsAffiliate bool  `json:"is_affiliate"`
            Logo string  `json:"logo"`
            ChannelShipmentID string  `json:"channel_shipment_id"`
         
    }
    
    // PlatformDeliveryAddress used by Order
    type PlatformDeliveryAddress struct {

        
            State string  `json:"state"`
            Pincode string  `json:"pincode"`
            Longitude float64  `json:"longitude"`
            UpdatedAt string  `json:"updated_at"`
            Version string  `json:"version"`
            Address2 string  `json:"address2"`
            City string  `json:"city"`
            Email string  `json:"email"`
            Latitude float64  `json:"latitude"`
            CreatedAt string  `json:"created_at"`
            Country string  `json:"country"`
            Address1 string  `json:"address1"`
            ContactPerson string  `json:"contact_person"`
            AddressCategory string  `json:"address_category"`
            Phone string  `json:"phone"`
            Landmark string  `json:"landmark"`
            Area string  `json:"area"`
            AddressType string  `json:"address_type"`
         
    }
    
    // ShipmentItem used by Order
    type ShipmentItem struct {

        
            Meta ShipmentItemMeta  `json:"meta"`
            ShipmentStatus ShipmentStatus  `json:"shipment_status"`
            PaymentMode string  `json:"payment_mode"`
            FulfillingStore ShipmentItemFulFillingStore  `json:"fulfilling_store"`
            ShipmentCreatedAt string  `json:"shipment_created_at"`
            EstimatedSlaTime string  `json:"estimated_sla_time"`
            ShipmentID string  `json:"shipment_id"`
            TotalBags float64  `json:"total_bags"`
            OrderingChannnel string  `json:"ordering_channnel"`
            DisplayName string  `json:"display_name"`
            StatusCreatedAt string  `json:"status_created_at"`
            User UserDataInfo  `json:"user"`
            Bags []BagUnit  `json:"bags"`
            CanProcess bool  `json:"can_process"`
            Channel ShipmentListingChannel  `json:"channel"`
            OrderID string  `json:"order_id"`
            OrderDate string  `json:"order_date"`
            DeliveryAddress PlatformDeliveryAddress  `json:"delivery_address"`
            LockStatus bool  `json:"lock_status"`
            InvoiceID string  `json:"invoice_id"`
            Prices Prices  `json:"prices"`
            PaymentMethods map[string]interface{}  `json:"payment_methods"`
            PreviousShipmentID string  `json:"previous_shipment_id"`
            CustomerNote string  `json:"customer_note"`
         
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
    
    // ShipmentInternalPlatformViewResponse used by Order
    type ShipmentInternalPlatformViewResponse struct {

        
            Message string  `json:"message"`
            Lane string  `json:"lane"`
            Items []ShipmentItem  `json:"items"`
            Success bool  `json:"success"`
            TotalCount float64  `json:"total_count"`
            Page Page  `json:"page"`
         
    }
    
    // Error used by Order
    type Error struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // ShipmentDetails used by Order
    type ShipmentDetails struct {

        
            ActionToStatus map[string]interface{}  `json:"action_to_status"`
            LockMessage string  `json:"lock_message"`
            LockStatus bool  `json:"lock_status"`
         
    }
    
    // EinvoiceInfo used by Order
    type EinvoiceInfo struct {

        
            Invoice map[string]interface{}  `json:"invoice"`
            CreditNote map[string]interface{}  `json:"credit_note"`
         
    }
    
    // DebugInfo used by Order
    type DebugInfo struct {

        
            StormbreakerUUID string  `json:"stormbreaker_uuid"`
         
    }
    
    // BuyerDetails used by Order
    type BuyerDetails struct {

        
            State string  `json:"state"`
            Pincode float64  `json:"pincode"`
            AjioSiteID string  `json:"ajio_site_id"`
            City string  `json:"city"`
            Address string  `json:"address"`
            Gstin string  `json:"gstin"`
            Name string  `json:"name"`
         
    }
    
    // ShipmentMeta used by Order
    type ShipmentMeta struct {

        
            DpOptions map[string]interface{}  `json:"dp_options"`
            EwaybillInfo map[string]interface{}  `json:"ewaybill_info"`
            DpSortKey string  `json:"dp_sort_key"`
            AssignDpFromSb bool  `json:"assign_dp_from_sb"`
            ReturnAwbNumber string  `json:"return_awb_number"`
            ForwardAffiliateShipmentID string  `json:"forward_affiliate_shipment_id"`
            Formatted Formatted  `json:"formatted"`
            AwbNumber string  `json:"awb_number"`
            DpID string  `json:"dp_id"`
            EinvoiceInfo EinvoiceInfo  `json:"einvoice_info"`
            ReturnDetails map[string]interface{}  `json:"return_details"`
            Dimension Dimensions  `json:"dimension"`
            DpName string  `json:"dp_name"`
            DebugInfo DebugInfo  `json:"debug_info"`
            FulfilmentPriorityText string  `json:"fulfilment_priority_text"`
            ShipmentTags []ShipmentTags  `json:"shipment_tags"`
            ShipmentVolumetricWeight float64  `json:"shipment_volumetric_weight"`
            ReturnAffiliateOrderID string  `json:"return_affiliate_order_id"`
            PoNumber string  `json:"po_number"`
            ForwardAffiliateOrderID string  `json:"forward_affiliate_order_id"`
            OrderType string  `json:"order_type"`
            DueDate string  `json:"due_date"`
            ReturnStoreNode float64  `json:"return_store_node"`
            ReturnAffiliateShipmentID string  `json:"return_affiliate_shipment_id"`
            External map[string]interface{}  `json:"external"`
            Weight float64  `json:"weight"`
            Timestamp ShipmentTimeStamp  `json:"timestamp"`
            B2cBuyerDetails map[string]interface{}  `json:"b2c_buyer_details"`
            StoreInvoiceUpdatedDate string  `json:"store_invoice_updated_date"`
            SameStoreAvailable bool  `json:"same_store_available"`
            PackagingName string  `json:"packaging_name"`
            LockData LockData  `json:"lock_data"`
            AutoTriggerDpAssignmentAcf bool  `json:"auto_trigger_dp_assignment_acf"`
            ShipmentWeight float64  `json:"shipment_weight"`
            ParentDpID string  `json:"parent_dp_id"`
            MarketplaceStoreID string  `json:"marketplace_store_id"`
            BagWeight map[string]interface{}  `json:"bag_weight"`
            B2bBuyerDetails BuyerDetails  `json:"b2b_buyer_details"`
            BoxType string  `json:"box_type"`
         
    }
    
    // ShipmentPayments used by Order
    type ShipmentPayments struct {

        
            Mode string  `json:"mode"`
            Source string  `json:"source"`
            Logo string  `json:"logo"`
         
    }
    
    // FulfillingStore used by Order
    type FulfillingStore struct {

        
            State string  `json:"state"`
            Pincode string  `json:"pincode"`
            Code string  `json:"code"`
            Meta map[string]interface{}  `json:"meta"`
            FulfillmentChannel string  `json:"fulfillment_channel"`
            City string  `json:"city"`
            Address string  `json:"address"`
            ID float64  `json:"id"`
            Country string  `json:"country"`
            ContactPerson string  `json:"contact_person"`
            Phone string  `json:"phone"`
            StoreName string  `json:"store_name"`
         
    }
    
    // UserDetailsData used by Order
    type UserDetailsData struct {

        
            State string  `json:"state"`
            Pincode string  `json:"pincode"`
            Email string  `json:"email"`
            City string  `json:"city"`
            Address string  `json:"address"`
            Country string  `json:"country"`
            Address1 string  `json:"address1"`
            AddressType string  `json:"address_type"`
            Phone string  `json:"phone"`
            Landmark string  `json:"landmark"`
            Area string  `json:"area"`
            Name string  `json:"name"`
         
    }
    
    // DPDetailsData used by Order
    type DPDetailsData struct {

        
            Pincode string  `json:"pincode"`
            EwayBillID string  `json:"eway_bill_id"`
            GstTag string  `json:"gst_tag"`
            ID float64  `json:"id"`
            Country string  `json:"country"`
            TrackURL string  `json:"track_url"`
            AwbNo string  `json:"awb_no"`
            Name string  `json:"name"`
         
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

        
            Address map[string]interface{}  `json:"address"`
            CompanyGst string  `json:"company_gst"`
            CompanyName string  `json:"company_name"`
            CompanyContact ContactDetails  `json:"company_contact"`
            CompanyCin string  `json:"company_cin"`
            CompanyID float64  `json:"company_id"`
         
    }
    
    // TrackingList used by Order
    type TrackingList struct {

        
            Time string  `json:"time"`
            IsPassed bool  `json:"is_passed"`
            Text string  `json:"text"`
            Status string  `json:"status"`
            IsCurrent bool  `json:"is_current"`
         
    }
    
    // GiftCard used by Order
    type GiftCard struct {

        
            GiftPrice float64  `json:"gift_price"`
            IsGiftApplied bool  `json:"is_gift_applied"`
            GiftMessage string  `json:"gift_message"`
            DisplayText string  `json:"display_text"`
         
    }
    
    // B2BPODetails used by Order
    type B2BPODetails struct {

        
            PartialCanRet bool  `json:"partial_can_ret"`
            DockerNumber string  `json:"docker_number"`
            PoTaxAmount float64  `json:"po_tax_amount"`
            ItemBasePrice float64  `json:"item_base_price"`
            TotalGstPercentage float64  `json:"total_gst_percentage"`
            PoLineAmount float64  `json:"po_line_amount"`
         
    }
    
    // BagMeta used by Order
    type BagMeta struct {

        
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            GiftCard GiftCard  `json:"gift_card"`
            PartialCanRet bool  `json:"partial_can_ret"`
            CustomMessage string  `json:"custom_message"`
            CustomJson map[string]interface{}  `json:"custom_json"`
            B2bPoDetails B2BPODetails  `json:"b2b_po_details"`
            DocketNumber string  `json:"docket_number"`
            GroupID string  `json:"group_id"`
         
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

        
            ArticleQuantity float64  `json:"article_quantity"`
            Amount float64  `json:"amount"`
            BuyRules []BuyRules  `json:"buy_rules"`
            PromotionName string  `json:"promotion_name"`
            PromotionType string  `json:"promotion_type"`
            MrpPromotion bool  `json:"mrp_promotion"`
            DiscountRules []DiscountRules  `json:"discount_rules"`
            PromoID string  `json:"promo_id"`
         
    }
    
    // OrderBrandName used by Order
    type OrderBrandName struct {

        
            CreatedOn string  `json:"created_on"`
            Logo string  `json:"logo"`
            ID float64  `json:"id"`
            BrandName string  `json:"brand_name"`
            ModifiedOn string  `json:"modified_on"`
            Company float64  `json:"company"`
         
    }
    
    // CurrentStatus used by Order
    type CurrentStatus struct {

        
            StateID float64  `json:"state_id"`
            KafkaSync bool  `json:"kafka_sync"`
            StoreID float64  `json:"store_id"`
            ShipmentID string  `json:"shipment_id"`
            CreatedAt string  `json:"created_at"`
            ID float64  `json:"id"`
            DeliveryPartnerID float64  `json:"delivery_partner_id"`
            DeliveryAwbNumber string  `json:"delivery_awb_number"`
            StateType string  `json:"state_type"`
            BagStateMapper BagStateMapper  `json:"bag_state_mapper"`
            BagID float64  `json:"bag_id"`
            Status string  `json:"status"`
            UpdatedAt float64  `json:"updated_at"`
         
    }
    
    // AffiliateBagsDetails used by Order
    type AffiliateBagsDetails struct {

        
            AffiliateBagID string  `json:"affiliate_bag_id"`
            CouponCode string  `json:"coupon_code"`
         
    }
    
    // OrderBagArticle used by Order
    type OrderBagArticle struct {

        
            UID string  `json:"uid"`
            Identifiers map[string]interface{}  `json:"identifiers"`
            ReturnConfig ReturnConfig  `json:"return_config"`
            Size string  `json:"size"`
         
    }
    
    // BagGST used by Order
    type BagGST struct {

        
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            HsnCodeID string  `json:"hsn_code_id"`
            SgstTaxPercentage float64  `json:"sgst_tax_percentage"`
            IsDefaultHsnCode bool  `json:"is_default_hsn_code"`
            GstTag string  `json:"gst_tag"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            IgstGstFee string  `json:"igst_gst_fee"`
            CgstTaxPercentage float64  `json:"cgst_tax_percentage"`
            GstFee float64  `json:"gst_fee"`
            SgstGstFee string  `json:"sgst_gst_fee"`
            ValueOfGood float64  `json:"value_of_good"`
            CgstGstFee string  `json:"cgst_gst_fee"`
            IgstTaxPercentage float64  `json:"igst_tax_percentage"`
            HsnCode string  `json:"hsn_code"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            GstinCode string  `json:"gstin_code"`
         
    }
    
    // BagConfigs used by Order
    type BagConfigs struct {

        
            IsCustomerReturnAllowed bool  `json:"is_customer_return_allowed"`
            AllowForceReturn bool  `json:"allow_force_return"`
            IsReturnable bool  `json:"is_returnable"`
            CanBeCancelled bool  `json:"can_be_cancelled"`
            EnableTracking bool  `json:"enable_tracking"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // BagPaymentMethods used by Order
    type BagPaymentMethods struct {

        
            Mode string  `json:"mode"`
            Amount float64  `json:"amount"`
         
    }
    
    // OrderBags used by Order
    type OrderBags struct {

        
            SellerIdentifier string  `json:"seller_identifier"`
            Meta BagMeta  `json:"meta"`
            AppliedPromos []AppliedPromos  `json:"applied_promos"`
            LineNumber float64  `json:"line_number"`
            Brand OrderBrandName  `json:"brand"`
            CanReturn bool  `json:"can_return"`
            ParentPromoBags map[string]interface{}  `json:"parent_promo_bags"`
            CanCancel bool  `json:"can_cancel"`
            CurrentStatus CurrentStatus  `json:"current_status"`
            DisplayName string  `json:"display_name"`
            AffiliateBagDetails AffiliateBagsDetails  `json:"affiliate_bag_details"`
            Article OrderBagArticle  `json:"article"`
            GstDetails BagGST  `json:"gst_details"`
            Identifier string  `json:"identifier"`
            GroupID string  `json:"group_id"`
            EntityType string  `json:"entity_type"`
            Item PlatformItem  `json:"item"`
            BagConfigs BagConfigs  `json:"bag_configs"`
            DeliveryAddress PlatformDeliveryAddress  `json:"delivery_address"`
            Quantity float64  `json:"quantity"`
            Prices Prices  `json:"prices"`
            BagID float64  `json:"bag_id"`
            FinancialBreakup FinancialBreakup  `json:"financial_breakup"`
            PaymentMethods []BagPaymentMethods  `json:"payment_methods"`
         
    }
    
    // OrderingStoreDetails used by Order
    type OrderingStoreDetails struct {

        
            State string  `json:"state"`
            Pincode string  `json:"pincode"`
            Code string  `json:"code"`
            Meta map[string]interface{}  `json:"meta"`
            City string  `json:"city"`
            Address string  `json:"address"`
            ID float64  `json:"id"`
            Country string  `json:"country"`
            ContactPerson string  `json:"contact_person"`
            Phone string  `json:"phone"`
            StoreName string  `json:"store_name"`
         
    }
    
    // PDFLinks used by Order
    type PDFLinks struct {

        
            InvoiceExport string  `json:"invoice_export"`
            Invoice string  `json:"invoice"`
            PoInvoice string  `json:"po_invoice"`
            InvoiceType string  `json:"invoice_type"`
            LabelType string  `json:"label_type"`
            LabelExport string  `json:"label_export"`
            LabelA4 string  `json:"label_a4"`
            B2b string  `json:"b2b"`
            Label string  `json:"label"`
            InvoicePos string  `json:"invoice_pos"`
            LabelPos string  `json:"label_pos"`
            LabelA6 string  `json:"label_a6"`
            InvoiceA6 string  `json:"invoice_a6"`
            InvoiceA4 string  `json:"invoice_a4"`
            CreditNoteURL string  `json:"credit_note_url"`
            DeliveryChallanA4 string  `json:"delivery_challan_a4"`
         
    }
    
    // AffiliateDetails used by Order
    type AffiliateDetails struct {

        
            AffiliateID string  `json:"affiliate_id"`
            AffiliateMeta AffiliateMeta  `json:"affiliate_meta"`
            PdfLinks PDFLinks  `json:"pdf_links"`
            AdID string  `json:"ad_id"`
            AffiliateStoreID string  `json:"affiliate_store_id"`
            CompanyAffiliateTag string  `json:"company_affiliate_tag"`
            ShipmentMeta ShipmentMeta  `json:"shipment_meta"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
         
    }
    
    // ShipmentStatusData used by Order
    type ShipmentStatusData struct {

        
            Meta map[string]interface{}  `json:"meta"`
            ShipmentID string  `json:"shipment_id"`
            CreatedAt string  `json:"created_at"`
            ID float64  `json:"id"`
            BagList []string  `json:"bag_list"`
            Status string  `json:"status"`
            DisplayName string  `json:"display_name"`
         
    }
    
    // OrderDetailsData used by Order
    type OrderDetailsData struct {

        
            AffiliateID string  `json:"affiliate_id"`
            OrderingChannel string  `json:"ordering_channel"`
            OrderingChannelLogo map[string]interface{}  `json:"ordering_channel_logo"`
            Source string  `json:"source"`
            TaxDetails map[string]interface{}  `json:"tax_details"`
            OrderDate string  `json:"order_date"`
            OrderValue string  `json:"order_value"`
            CodCharges string  `json:"cod_charges"`
            FyndOrderID string  `json:"fynd_order_id"`
         
    }
    
    // InvoiceInfo used by Order
    type InvoiceInfo struct {

        
            CreditNoteID string  `json:"credit_note_id"`
            UpdatedDate string  `json:"updated_date"`
            LabelURL string  `json:"label_url"`
            StoreInvoiceID string  `json:"store_invoice_id"`
            InvoiceURL string  `json:"invoice_url"`
            ExternalInvoiceID string  `json:"external_invoice_id"`
         
    }
    
    // PlatformShipment used by Order
    type PlatformShipment struct {

        
            Vertical string  `json:"vertical"`
            ShipmentDetails ShipmentDetails  `json:"shipment_details"`
            Meta ShipmentMeta  `json:"meta"`
            ShipmentStatus string  `json:"shipment_status"`
            ShipmentQuantity float64  `json:"shipment_quantity"`
            Payments ShipmentPayments  `json:"payments"`
            ShipmentImages []string  `json:"shipment_images"`
            OperationalStatus string  `json:"operational_status"`
            PaymentMode string  `json:"payment_mode"`
            FulfillingStore FulfillingStore  `json:"fulfilling_store"`
            DeliveryDetails UserDetailsData  `json:"delivery_details"`
            ShipmentCreatedAt string  `json:"shipment_created_at"`
            EstimatedSlaTime string  `json:"estimated_sla_time"`
            BagStatusHistory []BagStatusHistory  `json:"bag_status_history"`
            PackagingType string  `json:"packaging_type"`
            CustomMessage string  `json:"custom_message"`
            IsDpAssignEnabled bool  `json:"is_dp_assign_enabled"`
            ShipmentID string  `json:"shipment_id"`
            PriorityText string  `json:"priority_text"`
            UserAgent string  `json:"user_agent"`
            TotalBags float64  `json:"total_bags"`
            DpDetails DPDetailsData  `json:"dp_details"`
            DpAssignment bool  `json:"dp_assignment"`
            CompanyDetails CompanyDetails  `json:"company_details"`
            User UserDataInfo  `json:"user"`
            PickedDate string  `json:"picked_date"`
            GstDetails GSTDetailsData  `json:"gst_details"`
            TrackingList []TrackingList  `json:"tracking_list"`
            Bags []OrderBags  `json:"bags"`
            OrderingStore OrderingStoreDetails  `json:"ordering_store"`
            PlatformLogo string  `json:"platform_logo"`
            AffiliateDetails AffiliateDetails  `json:"affiliate_details"`
            BillingDetails UserDetailsData  `json:"billing_details"`
            FulfilmentPriority float64  `json:"fulfilment_priority"`
            DeliverySlot map[string]interface{}  `json:"delivery_slot"`
            Status ShipmentStatusData  `json:"status"`
            CustomMeta []map[string]interface{}  `json:"custom_meta"`
            Order OrderDetailsData  `json:"order"`
            Invoice InvoiceInfo  `json:"invoice"`
            ForwardShipmentID string  `json:"forward_shipment_id"`
            PdfLinks map[string]interface{}  `json:"pdf_links"`
            IsSelfShip bool  `json:"is_self_ship"`
            LockStatus bool  `json:"lock_status"`
            Coupon map[string]interface{}  `json:"coupon"`
            Prices Prices  `json:"prices"`
            CanUpdateDimension bool  `json:"can_update_dimension"`
            PaymentMethods map[string]interface{}  `json:"payment_methods"`
            JourneyType string  `json:"journey_type"`
            InvoiceID string  `json:"invoice_id"`
            TotalItems float64  `json:"total_items"`
            EnableDpTracking bool  `json:"enable_dp_tracking"`
         
    }
    
    // ShipmentInfoResponse used by Order
    type ShipmentInfoResponse struct {

        
            Shipments []PlatformShipment  `json:"shipments"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // AssetByShipment used by Order
    type AssetByShipment struct {

        
            PresignedUrls map[string]string  `json:"presigned_urls"`
            PresignedType string  `json:"presigned_type"`
            ShipmentID string  `json:"shipment_id"`
            ExpiresIn string  `json:"expires_in"`
            Success bool  `json:"success"`
         
    }
    
    // ResponseGetAssetShipment used by Order
    type ResponseGetAssetShipment struct {

        
            Data []AssetByShipment  `json:"data"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // PlatformUserDetails used by Order
    type PlatformUserDetails struct {

        
            PlatformUserEmployeeCode string  `json:"platform_user_employee_code"`
            PlatformUserLastName string  `json:"platform_user_last_name"`
            PlatformUserFirstName string  `json:"platform_user_first_name"`
            PlatformUserID string  `json:"platform_user_id"`
         
    }
    
    // TransactionData used by Order
    type TransactionData struct {

        
            Entity string  `json:"entity"`
            PaymentID string  `json:"payment_id"`
            UniqueReferenceNumber string  `json:"unique_reference_number"`
            AmountPaid float64  `json:"amount_paid"`
            TerminalID string  `json:"terminal_id"`
            Currency string  `json:"currency"`
            Status string  `json:"status"`
            TransactionID string  `json:"transaction_id"`
         
    }
    
    // BillingStaffDetails used by Order
    type BillingStaffDetails struct {

        
            User string  `json:"user"`
            LastName string  `json:"last_name"`
            StaffID float64  `json:"staff_id"`
            FirstName string  `json:"first_name"`
            EmployeeCode string  `json:"employee_code"`
         
    }
    
    // OrderMeta used by Order
    type OrderMeta struct {

        
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            CurrencySymbol string  `json:"currency_symbol"`
            PaymentType string  `json:"payment_type"`
            OrderChildEntities []string  `json:"order_child_entities"`
            PlatformUserDetails PlatformUserDetails  `json:"platform_user_details"`
            EmployeeID float64  `json:"employee_id"`
            Staff map[string]interface{}  `json:"staff"`
            CartID float64  `json:"cart_id"`
            OrderType string  `json:"order_type"`
            OrderPlatform string  `json:"order_platform"`
            OrderingStore float64  `json:"ordering_store"`
            TransactionData TransactionData  `json:"transaction_data"`
            Comment string  `json:"comment"`
            MongoCartID float64  `json:"mongo_cart_id"`
            Files []map[string]interface{}  `json:"files"`
            OrderTags []map[string]interface{}  `json:"order_tags"`
            BillingStaffDetails BillingStaffDetails  `json:"billing_staff_details"`
            CompanyLogo string  `json:"company_logo"`
            CustomerNote string  `json:"customer_note"`
         
    }
    
    // TaxDetails used by Order
    type TaxDetails struct {

        
            Gstin string  `json:"gstin"`
            PanNo string  `json:"pan_no"`
         
    }
    
    // OrderData used by Order
    type OrderData struct {

        
            Meta OrderMeta  `json:"meta"`
            TaxDetails TaxDetails  `json:"tax_details"`
            Prices Prices  `json:"prices"`
            OrderDate string  `json:"order_date"`
            PaymentMethods map[string]interface{}  `json:"payment_methods"`
            FyndOrderID string  `json:"fynd_order_id"`
         
    }
    
    // OrderDetailsResponse used by Order
    type OrderDetailsResponse struct {

        
            Shipments []PlatformShipment  `json:"shipments"`
            Success bool  `json:"success"`
            Order OrderData  `json:"order"`
         
    }
    
    // SubLane used by Order
    type SubLane struct {

        
            Value string  `json:"value"`
            Actions []map[string]interface{}  `json:"actions"`
            Index float64  `json:"index"`
            Text string  `json:"text"`
            TotalItems float64  `json:"total_items"`
         
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

        
            Logo string  `json:"logo"`
            Name string  `json:"name"`
         
    }
    
    // PlatformBreakupValues used by Order
    type PlatformBreakupValues struct {

        
            Value string  `json:"value"`
            Display string  `json:"display"`
            Name string  `json:"name"`
         
    }
    
    // PlatformOrderItems used by Order
    type PlatformOrderItems struct {

        
            Shipments []PlatformShipment  `json:"shipments"`
            OrderCreatedTime string  `json:"order_created_time"`
            Meta map[string]interface{}  `json:"meta"`
            Channel PlatformChannel  `json:"channel"`
            OrderID string  `json:"order_id"`
            PaymentMode string  `json:"payment_mode"`
            UserInfo UserDataInfo  `json:"user_info"`
            BreakupValues []PlatformBreakupValues  `json:"breakup_values"`
            TotalOrderValue float64  `json:"total_order_value"`
            OrderValue float64  `json:"order_value"`
         
    }
    
    // OrderListingResponse used by Order
    type OrderListingResponse struct {

        
            Message string  `json:"message"`
            Lane string  `json:"lane"`
            Items []PlatformOrderItems  `json:"items"`
            Success bool  `json:"success"`
            TotalCount float64  `json:"total_count"`
            Page Page  `json:"page"`
         
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

        
            Reason string  `json:"reason"`
            Awb string  `json:"awb"`
            UpdatedTime string  `json:"updated_time"`
            Meta map[string]interface{}  `json:"meta"`
            RawStatus string  `json:"raw_status"`
            AccountName string  `json:"account_name"`
            ShipmentType string  `json:"shipment_type"`
            LastLocationRecievedAt string  `json:"last_location_recieved_at"`
            Status string  `json:"status"`
            UpdatedAt string  `json:"updated_at"`
         
    }
    
    // PlatformShipmentTrack used by Order
    type PlatformShipmentTrack struct {

        
            Results []PlatformTrack  `json:"results"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // FilterInfoOption used by Order
    type FilterInfoOption struct {

        
            ShowUI bool  `json:"show_ui"`
            Value string  `json:"value"`
            PlaceholderText string  `json:"placeholder_text"`
            Text string  `json:"text"`
            MinSearchSize float64  `json:"min_search_size"`
            Name string  `json:"name"`
         
    }
    
    // FiltersInfo used by Order
    type FiltersInfo struct {

        
            Required bool  `json:"required"`
            Value string  `json:"value"`
            Type string  `json:"type"`
            PlaceholderText string  `json:"placeholder_text"`
            Options []FilterInfoOption  `json:"options"`
            Text string  `json:"text"`
         
    }
    
    // AdvanceFilterInfo used by Order
    type AdvanceFilterInfo struct {

        
            Returned []FiltersInfo  `json:"returned"`
            AppliedFilters map[string]interface{}  `json:"applied_filters"`
            Filters []FiltersInfo  `json:"filters"`
            ActionCentre []FiltersInfo  `json:"action_centre"`
            Unfulfilled []FiltersInfo  `json:"unfulfilled"`
            Processed []FiltersInfo  `json:"processed"`
            Page map[string]interface{}  `json:"page"`
         
    }
    
    // FiltersResponse used by Order
    type FiltersResponse struct {

        
            GlobalFilter []FiltersInfo  `json:"global_filter"`
            AdvanceFilter AdvanceFilterInfo  `json:"advance_filter"`
         
    }
    
    // Success used by Order
    type Success struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // OmsReports used by Order
    type OmsReports struct {

        
            ReportCreatedAt string  `json:"report_created_at"`
            ReportID string  `json:"report_id"`
            RequestDetails map[string]interface{}  `json:"request_details"`
            ReportType string  `json:"report_type"`
            ReportName string  `json:"report_name"`
            Status string  `json:"status"`
            DisplayName string  `json:"display_name"`
            S3Key string  `json:"s3_key"`
            ReportRequestedAt string  `json:"report_requested_at"`
         
    }
    
    // JioCodeUpsertDataSet used by Order
    type JioCodeUpsertDataSet struct {

        
            CompanyID string  `json:"company_id"`
            JioCode string  `json:"jio_code"`
            ArticleID string  `json:"article_id"`
            ItemID string  `json:"item_id"`
         
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

        
            Identifier string  `json:"identifier"`
            Data []map[string]interface{}  `json:"data"`
            Error []NestedErrorSchemaDataSet  `json:"error"`
            Success bool  `json:"success"`
            TraceID string  `json:"trace_id"`
         
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

        
            Value string  `json:"value"`
            Text string  `json:"text"`
         
    }
    
    // BulkActionTemplateResponse used by Order
    type BulkActionTemplateResponse struct {

        
            TemplateXSlug []BulkActionTemplate  `json:"template_x_slug"`
         
    }
    
    // QuestionSet used by Order
    type QuestionSet struct {

        
            ID float64  `json:"id"`
            DisplayName string  `json:"display_name"`
         
    }
    
    // Reason used by Order
    type Reason struct {

        
            QcType []string  `json:"qc_type"`
            ID float64  `json:"id"`
            DisplayName string  `json:"display_name"`
            QuestionSet []QuestionSet  `json:"question_set"`
         
    }
    
    // PlatformShipmentReasonsResponse used by Order
    type PlatformShipmentReasonsResponse struct {

        
            Success bool  `json:"success"`
            Reasons []Reason  `json:"reasons"`
         
    }
    
    // ShipmentResponseReasons used by Order
    type ShipmentResponseReasons struct {

        
            Reason string  `json:"reason"`
            ReasonID float64  `json:"reason_id"`
         
    }
    
    // ShipmentReasonsResponse used by Order
    type ShipmentReasonsResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
            Reasons []ShipmentResponseReasons  `json:"reasons"`
         
    }
    
    // Brand used by Order
    type Brand struct {

        
            StartDate string  `json:"start_date"`
            CreatedOn float64  `json:"created_on"`
            Logo string  `json:"logo"`
            ScriptLastRan string  `json:"script_last_ran"`
            InvoicePrefix string  `json:"invoice_prefix"`
            CreditNoteAllowed bool  `json:"credit_note_allowed"`
            PickupLocation string  `json:"pickup_location"`
            BrandName string  `json:"brand_name"`
            IsVirtualInvoice bool  `json:"is_virtual_invoice"`
            CreditNoteExpiryDays float64  `json:"credit_note_expiry_days"`
            ModifiedOn float64  `json:"modified_on"`
            Company string  `json:"company"`
            BrandID float64  `json:"brand_id"`
         
    }
    
    // ArticleDetails used by Order
    type ArticleDetails struct {

        
            Status map[string]interface{}  `json:"status"`
         
    }
    
    // BagGSTDetails used by Order
    type BagGSTDetails struct {

        
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            HsnCodeID string  `json:"hsn_code_id"`
            SgstTaxPercentage float64  `json:"sgst_tax_percentage"`
            IsDefaultHsnCode bool  `json:"is_default_hsn_code"`
            GstTag string  `json:"gst_tag"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            IgstGstFee string  `json:"igst_gst_fee"`
            CgstTaxPercentage float64  `json:"cgst_tax_percentage"`
            GstFee float64  `json:"gst_fee"`
            SgstGstFee string  `json:"sgst_gst_fee"`
            ValueOfGood float64  `json:"value_of_good"`
            CgstGstFee string  `json:"cgst_gst_fee"`
            IgstTaxPercentage float64  `json:"igst_tax_percentage"`
            HsnCode string  `json:"hsn_code"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            GstinCode string  `json:"gstin_code"`
         
    }
    
    // StoreEwaybill used by Order
    type StoreEwaybill struct {

        
            Enabled bool  `json:"enabled"`
         
    }
    
    // StoreEinvoice used by Order
    type StoreEinvoice struct {

        
            Enabled bool  `json:"enabled"`
            User string  `json:"user"`
            Password string  `json:"password"`
            Username string  `json:"username"`
         
    }
    
    // StoreGstCredentials used by Order
    type StoreGstCredentials struct {

        
            EWaybill StoreEwaybill  `json:"e_waybill"`
            EInvoice StoreEinvoice  `json:"e_invoice"`
         
    }
    
    // Document used by Order
    type Document struct {

        
            LegalName string  `json:"legal_name"`
            Verified bool  `json:"verified"`
            Value string  `json:"value"`
            URL string  `json:"url"`
            DsType string  `json:"ds_type"`
         
    }
    
    // StoreDocuments used by Order
    type StoreDocuments struct {

        
            Gst Document  `json:"gst"`
         
    }
    
    // EInvoicePortalDetails used by Order
    type EInvoicePortalDetails struct {

        
            Password string  `json:"password"`
            User string  `json:"user"`
            Username string  `json:"username"`
         
    }
    
    // StoreMeta used by Order
    type StoreMeta struct {

        
            Stage string  `json:"stage"`
            GstCredentials StoreGstCredentials  `json:"gst_credentials"`
            Timing []map[string]interface{}  `json:"timing"`
            EwaybillPortalDetails map[string]interface{}  `json:"ewaybill_portal_details"`
            ProductReturnConfig map[string]interface{}  `json:"product_return_config"`
            Documents StoreDocuments  `json:"documents"`
            NotificationEmails []string  `json:"notification_emails"`
            GstNumber string  `json:"gst_number"`
            AdditionalContactDetails map[string]interface{}  `json:"additional_contact_details"`
            DisplayName string  `json:"display_name"`
            EinvoicePortalDetails EInvoicePortalDetails  `json:"einvoice_portal_details"`
         
    }
    
    // StoreAddress used by Order
    type StoreAddress struct {

        
            Longitude float64  `json:"longitude"`
            City string  `json:"city"`
            Email string  `json:"email"`
            Country string  `json:"country"`
            Area string  `json:"area"`
            Pincode float64  `json:"pincode"`
            Version string  `json:"version"`
            Address1 string  `json:"address1"`
            CountryCode string  `json:"country_code"`
            AddressType string  `json:"address_type"`
            Latitude float64  `json:"latitude"`
            CreatedAt string  `json:"created_at"`
            ContactPerson string  `json:"contact_person"`
            AddressCategory string  `json:"address_category"`
            Landmark string  `json:"landmark"`
            UpdatedAt string  `json:"updated_at"`
            State string  `json:"state"`
            Address2 string  `json:"address2"`
            Phone string  `json:"phone"`
         
    }
    
    // Store used by Order
    type Store struct {

        
            Longitude float64  `json:"longitude"`
            VatNo string  `json:"vat_no"`
            Meta StoreMeta  `json:"meta"`
            City string  `json:"city"`
            Country string  `json:"country"`
            SID string  `json:"s_id"`
            PackagingMaterialCount float64  `json:"packaging_material_count"`
            LoginUsername string  `json:"login_username"`
            BrandID interface{}  `json:"brand_id"`
            IsActive bool  `json:"is_active"`
            Pincode string  `json:"pincode"`
            OrderIntegrationID string  `json:"order_integration_id"`
            Address1 string  `json:"address1"`
            MallName string  `json:"mall_name"`
            StoreAddressJson StoreAddress  `json:"store_address_json"`
            Code string  `json:"code"`
            FulfillmentChannel string  `json:"fulfillment_channel"`
            StoreEmail string  `json:"store_email"`
            Latitude float64  `json:"latitude"`
            CreatedAt string  `json:"created_at"`
            MallArea string  `json:"mall_area"`
            StoreActiveFrom string  `json:"store_active_from"`
            ContactPerson string  `json:"contact_person"`
            ParentStoreID float64  `json:"parent_store_id"`
            CompanyID float64  `json:"company_id"`
            UpdatedAt string  `json:"updated_at"`
            LocationType string  `json:"location_type"`
            State string  `json:"state"`
            IsEnabledForRecon bool  `json:"is_enabled_for_recon"`
            Phone float64  `json:"phone"`
            Address2 string  `json:"address2"`
            IsArchived bool  `json:"is_archived"`
            BrandStoreTags []string  `json:"brand_store_tags"`
            AlohomoraUserID float64  `json:"alohomora_user_id"`
            Name string  `json:"name"`
         
    }
    
    // Attributes used by Order
    type Attributes struct {

        
            PrimaryMaterial string  `json:"primary_material"`
            Gender []string  `json:"gender"`
            MarketerName string  `json:"marketer_name"`
            Name string  `json:"name"`
            BrandName string  `json:"brand_name"`
            PrimaryColorHex string  `json:"primary_color_hex"`
            MarketerAddress string  `json:"marketer_address"`
            PrimaryColor string  `json:"primary_color"`
            Essential string  `json:"essential"`
         
    }
    
    // Item used by Order
    type Item struct {

        
            Meta map[string]interface{}  `json:"meta"`
            Attributes Attributes  `json:"attributes"`
            DepartmentID float64  `json:"department_id"`
            Brand string  `json:"brand"`
            CanReturn bool  `json:"can_return"`
            L2CategoryID float64  `json:"l2_category_id"`
            BrandID float64  `json:"brand_id"`
            Gender string  `json:"gender"`
            Color string  `json:"color"`
            L1Category []string  `json:"l1_category"`
            CanCancel bool  `json:"can_cancel"`
            Image []string  `json:"image"`
            BranchURL string  `json:"branch_url"`
            Code string  `json:"code"`
            L3Category float64  `json:"l3_category"`
            ItemID float64  `json:"item_id"`
            L2Category []string  `json:"l2_category"`
            SlugKey string  `json:"slug_key"`
            L1CategoryID float64  `json:"l1_category_id"`
            LastUpdatedAt string  `json:"last_updated_at"`
            WebstoreProductURL string  `json:"webstore_product_url"`
            L3CategoryName string  `json:"l3_category_name"`
            Size string  `json:"size"`
            Name string  `json:"name"`
         
    }
    
    // BagDetailsPlatformResponse used by Order
    type BagDetailsPlatformResponse struct {

        
            SellerIdentifier string  `json:"seller_identifier"`
            Meta BagMeta  `json:"meta"`
            AppliedPromos []map[string]interface{}  `json:"applied_promos"`
            LineNumber float64  `json:"line_number"`
            Brand Brand  `json:"brand"`
            Type string  `json:"type"`
            ParentPromoBags map[string]interface{}  `json:"parent_promo_bags"`
            OriginalBagList []float64  `json:"original_bag_list"`
            OperationalStatus string  `json:"operational_status"`
            QcRequired interface{}  `json:"qc_required"`
            ArticleDetails ArticleDetails  `json:"article_details"`
            BagStatusHistory BagStatusHistory  `json:"bag_status_history"`
            RestorePromos map[string]interface{}  `json:"restore_promos"`
            Tags []string  `json:"tags"`
            RestoreCoupon bool  `json:"restore_coupon"`
            OrderIntegrationID string  `json:"order_integration_id"`
            ShipmentID string  `json:"shipment_id"`
            ID float64  `json:"id"`
            NoOfBagsOrder float64  `json:"no_of_bags_order"`
            BagStatus []BagStatusHistory  `json:"bag_status"`
            CurrentStatus BagStatusHistory  `json:"current_status"`
            DisplayName string  `json:"display_name"`
            AffiliateBagDetails AffiliateBagDetails  `json:"affiliate_bag_details"`
            Article Article  `json:"article"`
            Dates Dates  `json:"dates"`
            GstDetails BagGSTDetails  `json:"gst_details"`
            Identifier string  `json:"identifier"`
            OrderingStore Store  `json:"ordering_store"`
            AffiliateDetails AffiliateDetails  `json:"affiliate_details"`
            CurrentOperationalStatus BagStatusHistory  `json:"current_operational_status"`
            EntityType string  `json:"entity_type"`
            Status BagReturnableCancelableStatus  `json:"status"`
            Item Item  `json:"item"`
            Reasons []map[string]interface{}  `json:"reasons"`
            Quantity float64  `json:"quantity"`
            BagUpdateTime float64  `json:"bag_update_time"`
            Prices Prices  `json:"prices"`
            FinancialBreakup []FinancialBreakup  `json:"financial_breakup"`
            JourneyType string  `json:"journey_type"`
         
    }
    
    // ErrorResponse used by Order
    type ErrorResponse struct {

        
            Message string  `json:"message"`
            Error string  `json:"error"`
         
    }
    
    // BagsPage used by Order
    type BagsPage struct {

        
            Current float64  `json:"current"`
            ItemTotal float64  `json:"item_total"`
            PageType string  `json:"page_type"`
            Size float64  `json:"size"`
            HasNext bool  `json:"has_next"`
         
    }
    
    // GetBagsPlatformResponse used by Order
    type GetBagsPlatformResponse struct {

        
            Items []BagDetailsPlatformResponse  `json:"items"`
            Page BagsPage  `json:"page"`
         
    }
    
    // GeneratePosOrderReceiptResponse used by Order
    type GeneratePosOrderReceiptResponse struct {

        
            CustomerCnReceipt string  `json:"customer_cn_receipt"`
            MerchantCnReceipt string  `json:"merchant_cn_receipt"`
            PaymentReceipt string  `json:"payment_receipt"`
            OrderID string  `json:"order_id"`
            Success bool  `json:"success"`
            InvoiceReceipt string  `json:"invoice_receipt"`
         
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
            ShipmentID string  `json:"shipment_id"`
            Error string  `json:"error"`
         
    }
    
    // InvalidateShipmentCacheResponse used by Order
    type InvalidateShipmentCacheResponse struct {

        
            Response []InvalidateShipmentCacheNestedResponse  `json:"response"`
         
    }
    
    // ErrorResponse1 used by Order
    type ErrorResponse1 struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
            Status float64  `json:"status"`
            ErrorTrace string  `json:"error_trace"`
         
    }
    
    // StoreReassign used by Order
    type StoreReassign struct {

        
            ItemID string  `json:"item_id"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
            SetID string  `json:"set_id"`
            MongoArticleID string  `json:"mongo_article_id"`
            BagID float64  `json:"bag_id"`
            StoreID float64  `json:"store_id"`
            ReasonIds []float64  `json:"reason_ids"`
            AffiliateID string  `json:"affiliate_id"`
            FyndOrderID string  `json:"fynd_order_id"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
         
    }
    
    // StoreReassignResponse used by Order
    type StoreReassignResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // Entities used by Order
    type Entities struct {

        
            AffiliateBagID string  `json:"affiliate_bag_id"`
            ReasonText string  `json:"reason_text"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            AffiliateID string  `json:"affiliate_id"`
            ID string  `json:"id"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
         
    }
    
    // UpdateShipmentLockPayload used by Order
    type UpdateShipmentLockPayload struct {

        
            Action string  `json:"action"`
            Entities []Entities  `json:"entities"`
            EntityType string  `json:"entity_type"`
            ActionType string  `json:"action_type"`
         
    }
    
    // OriginalFilter used by Order
    type OriginalFilter struct {

        
            AffiliateID string  `json:"affiliate_id"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
         
    }
    
    // Bags used by Order
    type Bags struct {

        
            BagID float64  `json:"bag_id"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
            IsLocked bool  `json:"is_locked"`
         
    }
    
    // CheckResponse used by Order
    type CheckResponse struct {

        
            Status string  `json:"status"`
            IsShipmentLocked bool  `json:"is_shipment_locked"`
            OriginalFilter OriginalFilter  `json:"original_filter"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            LockStatus string  `json:"lock_status"`
            ShipmentID string  `json:"shipment_id"`
            AffiliateID string  `json:"affiliate_id"`
            IsBagLocked bool  `json:"is_bag_locked"`
            Bags []Bags  `json:"bags"`
         
    }
    
    // UpdateShipmentLockResponse used by Order
    type UpdateShipmentLockResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
            CheckResponse []CheckResponse  `json:"check_response"`
         
    }
    
    // AnnouncementResponse used by Order
    type AnnouncementResponse struct {

        
            ToDatetime string  `json:"to_datetime"`
            LogoURL string  `json:"logo_url"`
            PlatformName string  `json:"platform_name"`
            FromDatetime string  `json:"from_datetime"`
            PlatformID string  `json:"platform_id"`
            Description string  `json:"description"`
            CreatedAt string  `json:"created_at"`
            ID float64  `json:"id"`
            CompanyID float64  `json:"company_id"`
            Title string  `json:"title"`
         
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

        
            Success bool  `json:"success"`
            CallID string  `json:"call_id"`
         
    }
    
    // ErrorDetail used by Order
    type ErrorDetail struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // EntitiesDataUpdates used by Order
    type EntitiesDataUpdates struct {

        
            Data map[string]interface{}  `json:"data"`
            Filters []map[string]interface{}  `json:"filters"`
         
    }
    
    // OrderItemDataUpdates used by Order
    type OrderItemDataUpdates struct {

        
            Data map[string]interface{}  `json:"data"`
         
    }
    
    // ProductsDataUpdatesFilters used by Order
    type ProductsDataUpdatesFilters struct {

        
            LineNumber float64  `json:"line_number"`
            Identifier string  `json:"identifier"`
         
    }
    
    // ProductsDataUpdates used by Order
    type ProductsDataUpdates struct {

        
            Data map[string]interface{}  `json:"data"`
            Filters []ProductsDataUpdatesFilters  `json:"filters"`
         
    }
    
    // DataUpdates used by Order
    type DataUpdates struct {

        
            Entities []EntitiesDataUpdates  `json:"entities"`
            OrderItemStatus []OrderItemDataUpdates  `json:"order_item_status"`
            Products []ProductsDataUpdates  `json:"products"`
         
    }
    
    // Products used by Order
    type Products struct {

        
            LineNumber float64  `json:"line_number"`
            Quantity float64  `json:"quantity"`
            Identifier string  `json:"identifier"`
         
    }
    
    // EntityReasonData used by Order
    type EntityReasonData struct {

        
            ReasonText string  `json:"reason_text"`
            ReasonID float64  `json:"reason_id"`
         
    }
    
    // EntitiesReasons used by Order
    type EntitiesReasons struct {

        
            Data EntityReasonData  `json:"data"`
            Filters []map[string]interface{}  `json:"filters"`
         
    }
    
    // ProductsReasonsData used by Order
    type ProductsReasonsData struct {

        
            ReasonText string  `json:"reason_text"`
            ReasonID float64  `json:"reason_id"`
         
    }
    
    // ProductsReasonsFilters used by Order
    type ProductsReasonsFilters struct {

        
            LineNumber float64  `json:"line_number"`
            Quantity float64  `json:"quantity"`
            Identifier string  `json:"identifier"`
         
    }
    
    // ProductsReasons used by Order
    type ProductsReasons struct {

        
            Data ProductsReasonsData  `json:"data"`
            Filters []ProductsReasonsFilters  `json:"filters"`
         
    }
    
    // ReasonsData used by Order
    type ReasonsData struct {

        
            Entities []EntitiesReasons  `json:"entities"`
            Products []ProductsReasons  `json:"products"`
         
    }
    
    // ShipmentsRequest used by Order
    type ShipmentsRequest struct {

        
            DataUpdates DataUpdates  `json:"data_updates"`
            Identifier string  `json:"identifier"`
            Products []Products  `json:"products"`
            Reasons ReasonsData  `json:"reasons"`
         
    }
    
    // StatuesRequest used by Order
    type StatuesRequest struct {

        
            Shipments []ShipmentsRequest  `json:"shipments"`
            Status string  `json:"status"`
            ExcludeBagsNextState string  `json:"exclude_bags_next_state"`
         
    }
    
    // UpdateShipmentStatusRequest used by Order
    type UpdateShipmentStatusRequest struct {

        
            Statuses []StatuesRequest  `json:"statuses"`
            ForceTransition bool  `json:"force_transition"`
            LockAfterTransition bool  `json:"lock_after_transition"`
            UnlockBeforeTransition bool  `json:"unlock_before_transition"`
            Task bool  `json:"task"`
         
    }
    
    // ShipmentsResponse used by Order
    type ShipmentsResponse struct {

        
            Status float64  `json:"status"`
            Meta map[string]interface{}  `json:"meta"`
            FinalState map[string]interface{}  `json:"final_state"`
            Identifier string  `json:"identifier"`
            StackTrace string  `json:"stack_trace"`
            Message string  `json:"message"`
            Code string  `json:"code"`
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
    
    // DateRange used by Order
    type DateRange struct {

        
            ToDate string  `json:"to_date"`
            FromDate string  `json:"from_date"`
         
    }
    
    // FiltersRequest used by Order
    type FiltersRequest struct {

        
            Lane string  `json:"lane"`
            DateRange DateRange  `json:"date_range"`
            DpIds float64  `json:"dp_ids"`
            Logo string  `json:"logo"`
            Stores float64  `json:"stores"`
            DpName string  `json:"dp_name"`
            StoreName string  `json:"store_name"`
         
    }
    
    // ProcessManifest used by Order
    type ProcessManifest struct {

        
            Action string  `json:"action"`
            UniqueID string  `json:"unique_id"`
            Filters FiltersRequest  `json:"filters"`
            ManifestID string  `json:"manifest_id"`
         
    }
    
    // Filters used by Order
    type Filters struct {

        
            Lane string  `json:"lane"`
            DateRange DateRange  `json:"date_range"`
            DpIds float64  `json:"dp_ids"`
            Logo string  `json:"logo"`
            Stores float64  `json:"stores"`
            DpName string  `json:"dp_name"`
            StoreName string  `json:"store_name"`
            SelectedShipments string  `json:"selected_shipments"`
            ToDate string  `json:"to_date"`
            FromDate string  `json:"from_date"`
         
    }
    
    // ProcessManifestResponse used by Order
    type ProcessManifestResponse struct {

        
            ManifestID string  `json:"manifest_id"`
            CreatedBy string  `json:"created_by"`
            Action string  `json:"action"`
            UID string  `json:"uid"`
            CompanyID float64  `json:"company_id"`
            UserID string  `json:"user_id"`
            Filters Filters  `json:"filters"`
         
    }
    
    // ProcessManifestItemResponse used by Order
    type ProcessManifestItemResponse struct {

        
            Items ProcessManifestResponse  `json:"items"`
         
    }
    
    // FilterInfoOption1 used by Order
    type FilterInfoOption1 struct {

        
            Value string  `json:"value"`
            MinSearchSize float64  `json:"min_search_size"`
            Name string  `json:"name"`
            ShowUI bool  `json:"show_ui"`
            Text string  `json:"text"`
            PlaceholderText string  `json:"placeholder_text"`
         
    }
    
    // FiltersInfo1 used by Order
    type FiltersInfo1 struct {

        
            Type string  `json:"type"`
            Value string  `json:"value"`
            Options []FilterInfoOption1  `json:"options"`
            Required bool  `json:"required"`
            Text string  `json:"text"`
            PlaceholderText string  `json:"placeholder_text"`
         
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

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // ActionInfo used by Order
    type ActionInfo struct {

        
            Description string  `json:"description"`
            DisplayText string  `json:"display_text"`
            Slug string  `json:"slug"`
            ID float64  `json:"id"`
         
    }
    
    // GetActionsResponse used by Order
    type GetActionsResponse struct {

        
            Permissions []ActionInfo  `json:"permissions"`
         
    }
    
    // HistoryReason used by Order
    type HistoryReason struct {

        
            DislayName string  `json:"dislay_name"`
            Category string  `json:"category"`
            State string  `json:"state"`
            Code float64  `json:"code"`
            Quantity float64  `json:"quantity"`
            Text string  `json:"text"`
         
    }
    
    // HistoryMeta used by Order
    type HistoryMeta struct {

        
            ActivityComment string  `json:"activity_comment"`
            Status string  `json:"status"`
            Caller string  `json:"caller"`
            Recipient string  `json:"recipient"`
            Slug string  `json:"slug"`
            StoreName string  `json:"store_name"`
            StoreID float64  `json:"store_id"`
            StoreCode string  `json:"store_code"`
            CallID string  `json:"call_id"`
            Status1 string  `json:"status1"`
            Recordpath string  `json:"recordpath"`
            ActivityType string  `json:"activity_type"`
            ChannelType string  `json:"channel_type"`
            Message string  `json:"message"`
            Callerid string  `json:"callerid"`
            Endtime string  `json:"endtime"`
            ShortLink string  `json:"short_link"`
            Starttime string  `json:"starttime"`
            Status2 string  `json:"status2"`
            Billsec string  `json:"billsec"`
            Duration string  `json:"duration"`
            Reason HistoryReason  `json:"reason"`
            Receiver string  `json:"receiver"`
         
    }
    
    // HistoryDict used by Order
    type HistoryDict struct {

        
            User string  `json:"user"`
            Meta HistoryMeta  `json:"meta"`
            AssignedAgent string  `json:"assigned_agent"`
            Type string  `json:"type"`
            TicketID string  `json:"ticket_id"`
            L3Detail string  `json:"l3_detail"`
            TicketURL string  `json:"ticket_url"`
            Message string  `json:"message"`
            Createdat string  `json:"createdat"`
            BagID float64  `json:"bag_id"`
            DisplayMessage string  `json:"display_message"`
            L1Detail string  `json:"l1_detail"`
            L2Detail string  `json:"l2_detail"`
         
    }
    
    // ShipmentHistoryResponse used by Order
    type ShipmentHistoryResponse struct {

        
            Success bool  `json:"success"`
            ActivityHistory []HistoryDict  `json:"activity_history"`
         
    }
    
    // PostHistoryData used by Order
    type PostHistoryData struct {

        
            Message string  `json:"message"`
            UserName string  `json:"user_name"`
         
    }
    
    // PostHistoryFilters used by Order
    type PostHistoryFilters struct {

        
            LineNumber string  `json:"line_number"`
            Identifier string  `json:"identifier"`
            ShipmentID string  `json:"shipment_id"`
         
    }
    
    // PostActivityHistory used by Order
    type PostActivityHistory struct {

        
            Data PostHistoryData  `json:"data"`
            Filters []PostHistoryFilters  `json:"filters"`
         
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
            BrandName string  `json:"brand_name"`
            Message string  `json:"message"`
            CountryCode string  `json:"country_code"`
            PhoneNumber float64  `json:"phone_number"`
            ShipmentID float64  `json:"shipment_id"`
            PaymentMode string  `json:"payment_mode"`
            AmountPaid float64  `json:"amount_paid"`
            OrderID string  `json:"order_id"`
         
    }
    
    // SendSmsPayload used by Order
    type SendSmsPayload struct {

        
            BagID float64  `json:"bag_id"`
            Data SmsDataPayload  `json:"data"`
            Slug string  `json:"slug"`
         
    }
    
    // Meta used by Order
    type Meta struct {

        
            StateManagerUsed string  `json:"state_manager_used"`
            KafkaEmissionStatus float64  `json:"kafka_emission_status"`
         
    }
    
    // ShipmentDetail used by Order
    type ShipmentDetail struct {

        
            BagList []float64  `json:"bag_list"`
            Meta Meta  `json:"meta"`
            Status string  `json:"status"`
            ShipmentID string  `json:"shipment_id"`
            ID float64  `json:"id"`
            Remarks string  `json:"remarks"`
         
    }
    
    // OrderDetails used by Order
    type OrderDetails struct {

        
            FyndOrderID string  `json:"fynd_order_id"`
            CreatedAt string  `json:"created_at"`
         
    }
    
    // OrderStatusData used by Order
    type OrderStatusData struct {

        
            Errors []string  `json:"errors"`
            ShipmentDetails []ShipmentDetail  `json:"shipment_details"`
            OrderDetails OrderDetails  `json:"order_details"`
         
    }
    
    // OrderStatusResult used by Order
    type OrderStatusResult struct {

        
            Success string  `json:"success"`
            Result []OrderStatusData  `json:"result"`
         
    }
    
    // Dimension used by Order
    type Dimension struct {

        
            PackagingType string  `json:"packaging_type"`
            Width float64  `json:"width"`
            Weight string  `json:"weight"`
            Height string  `json:"height"`
            Length float64  `json:"length"`
         
    }
    
    // UpdatePackagingDimensionsPayload used by Order
    type UpdatePackagingDimensionsPayload struct {

        
            Dimension []Dimension  `json:"dimension"`
            ShipmentID string  `json:"shipment_id"`
            CurrentStatus string  `json:"current_status"`
         
    }
    
    // UpdatePackagingDimensionsResponse used by Order
    type UpdatePackagingDimensionsResponse struct {

        
            Message string  `json:"message"`
         
    }
    
    // TaxInfo used by Order
    type TaxInfo struct {

        
            B2bGstinNumber string  `json:"b2b_gstin_number"`
            Gstin string  `json:"gstin"`
         
    }
    
    // BillingInfo used by Order
    type BillingInfo struct {

        
            LastName string  `json:"last_name"`
            City string  `json:"city"`
            AlternateMobileNumber string  `json:"alternate_mobile_number"`
            MiddleName string  `json:"middle_name"`
            Gender string  `json:"gender"`
            State string  `json:"state"`
            PrimaryMobileNumber string  `json:"primary_mobile_number"`
            PrimaryEmail string  `json:"primary_email"`
            StateCode string  `json:"state_code"`
            Pincode string  `json:"pincode"`
            ExternalCustomerCode string  `json:"external_customer_code"`
            Address2 string  `json:"address2"`
            CountryCode string  `json:"country_code"`
            Title string  `json:"title"`
            Country string  `json:"country"`
            CustomerCode string  `json:"customer_code"`
            FloorNo string  `json:"floor_no"`
            HouseNo string  `json:"house_no"`
            AlternateEmail string  `json:"alternate_email"`
            FirstName string  `json:"first_name"`
            Address1 string  `json:"address1"`
         
    }
    
    // ShippingInfo used by Order
    type ShippingInfo struct {

        
            LastName string  `json:"last_name"`
            City string  `json:"city"`
            AlternateMobileNumber string  `json:"alternate_mobile_number"`
            GeoLocation map[string]interface{}  `json:"geo_location"`
            MiddleName string  `json:"middle_name"`
            Gender string  `json:"gender"`
            State string  `json:"state"`
            PrimaryMobileNumber string  `json:"primary_mobile_number"`
            PrimaryEmail string  `json:"primary_email"`
            AddressType string  `json:"address_type"`
            StateCode string  `json:"state_code"`
            AlternateEmail string  `json:"alternate_email"`
            Pincode string  `json:"pincode"`
            ExternalCustomerCode string  `json:"external_customer_code"`
            Slot []map[string]interface{}  `json:"slot"`
            Address2 string  `json:"address2"`
            Landmark string  `json:"landmark"`
            CountryCode string  `json:"country_code"`
            Title string  `json:"title"`
            Country string  `json:"country"`
            CustomerCode string  `json:"customer_code"`
            FloorNo string  `json:"floor_no"`
            HouseNo string  `json:"house_no"`
            ShippingType string  `json:"shipping_type"`
            FirstName string  `json:"first_name"`
            Address1 string  `json:"address1"`
         
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

        
            Type string  `json:"type"`
            Tax Tax  `json:"tax"`
            Code string  `json:"code"`
            Amount map[string]interface{}  `json:"amount"`
            Name string  `json:"name"`
         
    }
    
    // LineItem used by Order
    type LineItem struct {

        
            CustomMessasge string  `json:"custom_messasge"`
            Meta map[string]interface{}  `json:"meta"`
            ExternalLineID string  `json:"external_line_id"`
            Charges []Charge  `json:"charges"`
            SellerIdentifier string  `json:"seller_identifier"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // ProcessingDates used by Order
    type ProcessingDates struct {

        
            CustomerPickupSlot map[string]interface{}  `json:"customer_pickup_slot"`
            DpPickupSlot map[string]interface{}  `json:"dp_pickup_slot"`
            ConfirmByDate string  `json:"confirm_by_date"`
            DispatchAfterDate string  `json:"dispatch_after_date"`
            DispatchByDate string  `json:"dispatch_by_date"`
            PackByDate string  `json:"pack_by_date"`
         
    }
    
    // Shipment used by Order
    type Shipment struct {

        
            LineItems []LineItem  `json:"line_items"`
            ProcessingDates ProcessingDates  `json:"processing_dates"`
            Meta map[string]interface{}  `json:"meta"`
            ExternalShipmentID string  `json:"external_shipment_id"`
            Priority float64  `json:"priority"`
            LocationID float64  `json:"location_id"`
         
    }
    
    // PaymentMethod used by Order
    type PaymentMethod struct {

        
            Mode string  `json:"mode"`
            Meta map[string]interface{}  `json:"meta"`
            CollectBy string  `json:"collect_by"`
            Amount float64  `json:"amount"`
            RefundBy string  `json:"refund_by"`
            TransactionData map[string]interface{}  `json:"transaction_data"`
            Name string  `json:"name"`
         
    }
    
    // PaymentInfo used by Order
    type PaymentInfo struct {

        
            PaymentMethods []PaymentMethod  `json:"payment_methods"`
            PrimaryMode string  `json:"primary_mode"`
         
    }
    
    // CreateOrderAPI used by Order
    type CreateOrderAPI struct {

        
            TaxInfo TaxInfo  `json:"tax_info"`
            BillingInfo BillingInfo  `json:"billing_info"`
            Meta map[string]interface{}  `json:"meta"`
            ShippingInfo ShippingInfo  `json:"shipping_info"`
            Config map[string]interface{}  `json:"config"`
            ExternalCreationDate string  `json:"external_creation_date"`
            Charges []Charge  `json:"charges"`
            ExternalOrderID string  `json:"external_order_id"`
            Shipments []Shipment  `json:"shipments"`
            CurrencyInfo map[string]interface{}  `json:"currency_info"`
            PaymentInfo PaymentInfo  `json:"payment_info"`
         
    }
    
    // CreateOrderResponse used by Order
    type CreateOrderResponse struct {

        
            FyndOrderID string  `json:"fynd_order_id"`
         
    }
    
    // CreateOrderErrorReponse used by Order
    type CreateOrderErrorReponse struct {

        
            RequestID string  `json:"request_id"`
            Status float64  `json:"status"`
            Meta string  `json:"meta"`
            StackTrace string  `json:"stack_trace"`
            Message string  `json:"message"`
            Code string  `json:"code"`
            Exception string  `json:"exception"`
            Info interface{}  `json:"info"`
         
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
    
    // BagStateTransitionMap used by Order
    type BagStateTransitionMap struct {

        
            Fynd map[string]interface{}  `json:"fynd"`
            Affiliate map[string]interface{}  `json:"affiliate"`
         
    }
    
    // ManifestPageInfo used by Order
    type ManifestPageInfo struct {

        
            Total float64  `json:"total"`
            Type string  `json:"type"`
            Size float64  `json:"size"`
            HasNext bool  `json:"has_next"`
            HasPrevious bool  `json:"has_previous"`
            Current float64  `json:"current"`
         
    }
    
    // ManifestItemDetails used by Order
    type ManifestItemDetails struct {

        
            ShipmentCreatedAt string  `json:"shipment_created_at"`
            Awb string  `json:"awb"`
            ShipmentID string  `json:"shipment_id"`
            ItemQty float64  `json:"item_qty"`
            InvoiceID string  `json:"invoice_id"`
            OrderID string  `json:"order_id"`
         
    }
    
    // ManifestShipmentListing used by Order
    type ManifestShipmentListing struct {

        
            Lane string  `json:"lane"`
            Status float64  `json:"status"`
            TotalCount float64  `json:"total_count"`
            Page ManifestPageInfo  `json:"page"`
            Message string  `json:"message"`
            Items []ManifestItemDetails  `json:"items"`
            Success bool  `json:"success"`
         
    }
    
    // TotalShipmentPricesCount used by Order
    type TotalShipmentPricesCount struct {

        
            ShipmentCount float64  `json:"shipment_count"`
            TotalPrice float64  `json:"total_price"`
         
    }
    
    // ManifestMeta used by Order
    type ManifestMeta struct {

        
            Filters Filters  `json:"filters"`
            TotalShipmentPricesCount TotalShipmentPricesCount  `json:"total_shipment_prices_count"`
         
    }
    
    // ManifestFile used by Order
    type ManifestFile struct {

        
            Region string  `json:"region"`
            Bucket string  `json:"bucket"`
            Key string  `json:"key"`
         
    }
    
    // ManifestMediaUpdate used by Order
    type ManifestMediaUpdate struct {

        
            Link string  `json:"link"`
            Status bool  `json:"status"`
            Entity string  `json:"entity"`
            Code float64  `json:"code"`
            File ManifestFile  `json:"file"`
            MediaType string  `json:"media_type"`
         
    }
    
    // PDFMeta used by Order
    type PDFMeta struct {

        
            MediaUpdates []ManifestMediaUpdate  `json:"media_updates"`
            Consent string  `json:"consent"`
         
    }
    
    // Manifest used by Order
    type Manifest struct {

        
            Meta ManifestMeta  `json:"meta"`
            Status string  `json:"status"`
            ManifestID string  `json:"manifest_id"`
            PdfMeta PDFMeta  `json:"pdf_meta"`
            Filters Filters  `json:"filters"`
            CreatedAt string  `json:"created_at"`
            CreatedBy string  `json:"created_by"`
            ID float64  `json:"id"`
            UID string  `json:"uid"`
            CompanyID float64  `json:"company_id"`
            UserID string  `json:"user_id"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // ManifestList used by Order
    type ManifestList struct {

        
            Page ManifestPageInfo  `json:"page"`
            Items []Manifest  `json:"items"`
         
    }
    
    // ManifestDetails used by Order
    type ManifestDetails struct {

        
            AdditionalShipmentCount float64  `json:"additional_shipment_count"`
            Page ManifestPageInfo  `json:"page"`
            Items []ManifestItemDetails  `json:"items"`
            ManifestDetails []Manifest  `json:"manifest_details"`
         
    }
    
    // FetchCreditBalanceRequestPayload used by Order
    type FetchCreditBalanceRequestPayload struct {

        
            SellerID string  `json:"seller_id"`
            AffiliateID string  `json:"affiliate_id"`
            CustomerMobileNumber string  `json:"customer_mobile_number"`
         
    }
    
    // CreditBalanceInfo used by Order
    type CreditBalanceInfo struct {

        
            TotalCreditedBalance string  `json:"total_credited_balance"`
            Reason string  `json:"reason"`
            CustomerMobileNumber string  `json:"customer_mobile_number"`
         
    }
    
    // FetchCreditBalanceResponsePayload used by Order
    type FetchCreditBalanceResponsePayload struct {

        
            Success bool  `json:"success"`
            Data CreditBalanceInfo  `json:"data"`
         
    }
    
    // RefundModeConfigRequestPayload used by Order
    type RefundModeConfigRequestPayload struct {

        
            FyndOrderID string  `json:"fynd_order_id"`
            CustomerMobileNumber string  `json:"customer_mobile_number"`
            AffiliateID string  `json:"affiliate_id"`
            SellerID float64  `json:"seller_id"`
         
    }
    
    // RefundOption used by Order
    type RefundOption struct {

        
            Name string  `json:"name"`
            Value string  `json:"value"`
            Type string  `json:"type"`
            Slug string  `json:"slug"`
         
    }
    
    // RefundModeInfo used by Order
    type RefundModeInfo struct {

        
            DisplayName string  `json:"display_name"`
            Options []RefundOption  `json:"options"`
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

        
            Mobile string  `json:"mobile"`
            CountryCode string  `json:"country_code"`
            FirstName string  `json:"first_name"`
            LastName string  `json:"last_name"`
         
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

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // SendUserMobileOTP used by Order
    type SendUserMobileOTP struct {

        
            Mobile string  `json:"mobile"`
            CountryCode string  `json:"country_code"`
         
    }
    
    // PointBlankOtpData used by Order
    type PointBlankOtpData struct {

        
            RequestID string  `json:"request_id"`
            Message string  `json:"message"`
            Mobile float64  `json:"mobile"`
            ResendTimer float64  `json:"resend_timer"`
         
    }
    
    // SendUserMobileOtpResponse used by Order
    type SendUserMobileOtpResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
            Data PointBlankOtpData  `json:"data"`
         
    }
    
    // VerifyOtpData used by Order
    type VerifyOtpData struct {

        
            RequestID string  `json:"request_id"`
            Mobile string  `json:"mobile"`
            OtpCode float64  `json:"otp_code"`
         
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
            CountryCode string  `json:"country_code"`
            Message string  `json:"message"`
         
    }
    
    // VerifyOtpResponse used by Order
    type VerifyOtpResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
            Status float64  `json:"status"`
            Data VerifyOtpResponseData  `json:"data"`
         
    }
    

    
    // SearchKeywordResult used by Catalog
    type SearchKeywordResult struct {

        
            SortOn string  `json:"sort_on"`
            Query map[string]interface{}  `json:"query"`
         
    }
    
    // CreateSearchKeyword used by Catalog
    type CreateSearchKeyword struct {

        
            Result SearchKeywordResult  `json:"result"`
            AppID string  `json:"app_id"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Words []string  `json:"words"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // GetSearchWordsData used by Catalog
    type GetSearchWordsData struct {

        
            Result map[string]interface{}  `json:"result"`
            AppID string  `json:"app_id"`
            UID string  `json:"uid"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Words []string  `json:"words"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // ErrorResponse used by Catalog
    type ErrorResponse struct {

        
            Error string  `json:"error"`
            Meta map[string]interface{}  `json:"meta"`
            Message string  `json:"message"`
            Code string  `json:"code"`
            Status float64  `json:"status"`
         
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

        
            Page Page  `json:"page"`
            Items GetSearchWordsData  `json:"items"`
         
    }
    
    // DeleteResponse used by Catalog
    type DeleteResponse struct {

        
            Message string  `json:"message"`
         
    }
    
    // GetSearchWordsResponse used by Catalog
    type GetSearchWordsResponse struct {

        
            Page Page  `json:"page"`
            Items []GetSearchWordsData  `json:"items"`
         
    }
    
    // AutocompletePageAction used by Catalog
    type AutocompletePageAction struct {

        
            URL string  `json:"url"`
            Type string  `json:"type"`
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

        
            AspectRatio string  `json:"aspect_ratio"`
            URL string  `json:"url"`
            Type string  `json:"type"`
         
    }
    
    // AutocompleteResult used by Catalog
    type AutocompleteResult struct {

        
            Action AutocompleteAction  `json:"action"`
            Display string  `json:"display"`
            Logo AutoCompleteMedia  `json:"logo"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // CreateAutocompleteKeyword used by Catalog
    type CreateAutocompleteKeyword struct {

        
            AppID string  `json:"app_id"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Words []string  `json:"words"`
            Results []AutocompleteResult  `json:"results"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // GetAutocompleteWordsData used by Catalog
    type GetAutocompleteWordsData struct {

        
            AppID string  `json:"app_id"`
            UID string  `json:"uid"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Words []string  `json:"words"`
            Results []map[string]interface{}  `json:"results"`
         
    }
    
    // GetAutocompleteWordsResponse used by Catalog
    type GetAutocompleteWordsResponse struct {

        
            Page Page  `json:"page"`
            Items []GetAutocompleteWordsData  `json:"items"`
         
    }
    
    // CreateAutocompleteWordsResponse used by Catalog
    type CreateAutocompleteWordsResponse struct {

        
            Results []map[string]interface{}  `json:"results"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Words []string  `json:"words"`
            AppID string  `json:"app_id"`
         
    }
    
    // ProductBundleItem used by Catalog
    type ProductBundleItem struct {

        
            AllowRemove bool  `json:"allow_remove"`
            MinQuantity float64  `json:"min_quantity"`
            MaxQuantity float64  `json:"max_quantity"`
            ProductUID float64  `json:"product_uid"`
            AutoSelect bool  `json:"auto_select"`
            AutoAddToCart bool  `json:"auto_add_to_cart"`
         
    }
    
    // ProductBundleRequest used by Catalog
    type ProductBundleRequest struct {

        
            Logo string  `json:"logo"`
            PageVisibility []string  `json:"page_visibility"`
            CompanyID float64  `json:"company_id"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Meta map[string]interface{}  `json:"meta"`
            Name string  `json:"name"`
            Choice string  `json:"choice"`
            Products []ProductBundleItem  `json:"products"`
            ModifiedOn string  `json:"modified_on"`
            Slug string  `json:"slug"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            CreatedOn string  `json:"created_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // GetProductBundleCreateResponse used by Catalog
    type GetProductBundleCreateResponse struct {

        
            ID string  `json:"id"`
            Logo string  `json:"logo"`
            PageVisibility []string  `json:"page_visibility"`
            CompanyID float64  `json:"company_id"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Meta map[string]interface{}  `json:"meta"`
            Name string  `json:"name"`
            Choice string  `json:"choice"`
            Products []ProductBundleItem  `json:"products"`
            ModifiedOn string  `json:"modified_on"`
            Slug string  `json:"slug"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            CreatedOn string  `json:"created_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // GetProductBundleListingResponse used by Catalog
    type GetProductBundleListingResponse struct {

        
            Page Page  `json:"page"`
            Items []GetProductBundleCreateResponse  `json:"items"`
         
    }
    
    // ProductBundleUpdateRequest used by Catalog
    type ProductBundleUpdateRequest struct {

        
            Logo string  `json:"logo"`
            PageVisibility []string  `json:"page_visibility"`
            CompanyID float64  `json:"company_id"`
            Meta map[string]interface{}  `json:"meta"`
            Name string  `json:"name"`
            Choice string  `json:"choice"`
            Products []ProductBundleItem  `json:"products"`
            ModifiedOn string  `json:"modified_on"`
            Slug string  `json:"slug"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // LimitedProductData used by Catalog
    type LimitedProductData struct {

        
            Images []string  `json:"images"`
            Sizes []string  `json:"sizes"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            Identifier map[string]interface{}  `json:"identifier"`
            Quantity float64  `json:"quantity"`
            CountryOfOrigin string  `json:"country_of_origin"`
            ItemCode string  `json:"item_code"`
            ShortDescription string  `json:"short_description"`
            Slug string  `json:"slug"`
            Attributes map[string]interface{}  `json:"attributes"`
            Price map[string]interface{}  `json:"price"`
         
    }
    
    // Size used by Catalog
    type Size struct {

        
            IsAvailable bool  `json:"is_available"`
            Display string  `json:"display"`
            Value string  `json:"value"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // Price used by Catalog
    type Price struct {

        
            MaxEffective float64  `json:"max_effective"`
            Currency string  `json:"currency"`
            MinEffective float64  `json:"min_effective"`
            MaxMarked float64  `json:"max_marked"`
            MinMarked float64  `json:"min_marked"`
         
    }
    
    // GetProducts used by Catalog
    type GetProducts struct {

        
            AllowRemove bool  `json:"allow_remove"`
            MaxQuantity float64  `json:"max_quantity"`
            MinQuantity float64  `json:"min_quantity"`
            ProductUID float64  `json:"product_uid"`
            AutoSelect bool  `json:"auto_select"`
            ProductDetails LimitedProductData  `json:"product_details"`
            AutoAddToCart bool  `json:"auto_add_to_cart"`
            Sizes []Size  `json:"sizes"`
            Price Price  `json:"price"`
         
    }
    
    // GetProductBundleResponse used by Catalog
    type GetProductBundleResponse struct {

        
            Logo string  `json:"logo"`
            PageVisibility []string  `json:"page_visibility"`
            CompanyID float64  `json:"company_id"`
            Meta map[string]interface{}  `json:"meta"`
            Name string  `json:"name"`
            Choice string  `json:"choice"`
            Products []GetProducts  `json:"products"`
            Slug string  `json:"slug"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            IsActive bool  `json:"is_active"`
         
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

        
            ID string  `json:"id"`
            Image string  `json:"image"`
            BrandID float64  `json:"brand_id"`
            CompanyID float64  `json:"company_id"`
            Active bool  `json:"active"`
            ModifiedOn string  `json:"modified_on"`
            Subtitle string  `json:"subtitle"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Title string  `json:"title"`
            Description string  `json:"description"`
            Name string  `json:"name"`
            Guide Guide  `json:"guide"`
            Tag string  `json:"tag"`
            CreatedOn string  `json:"created_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
         
    }
    
    // SuccessResponse used by Catalog
    type SuccessResponse struct {

        
            Success bool  `json:"success"`
         
    }
    
    // ListSizeGuide used by Catalog
    type ListSizeGuide struct {

        
            Page map[string]interface{}  `json:"page"`
            Items []map[string]interface{}  `json:"items"`
         
    }
    
    // SizeGuideResponse used by Catalog
    type SizeGuideResponse struct {

        
            ID string  `json:"id"`
            BrandID float64  `json:"brand_id"`
            CompanyID float64  `json:"company_id"`
            Active bool  `json:"active"`
            Subtitle string  `json:"subtitle"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Title string  `json:"title"`
            Name string  `json:"name"`
            Guide map[string]interface{}  `json:"guide"`
            ModifiedOn string  `json:"modified_on"`
            Tag string  `json:"tag"`
            CreatedOn string  `json:"created_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
         
    }
    
    // ApplicationItemMOQ used by Catalog
    type ApplicationItemMOQ struct {

        
            Minimum float64  `json:"minimum"`
            IncrementUnit float64  `json:"increment_unit"`
            Maximum float64  `json:"maximum"`
         
    }
    
    // MetaFields used by Catalog
    type MetaFields struct {

        
            Key string  `json:"key"`
            Value string  `json:"value"`
         
    }
    
    // ApplicationItemSEO used by Catalog
    type ApplicationItemSEO struct {

        
            Title string  `json:"title"`
            Description string  `json:"description"`
         
    }
    
    // ApplicationItemMeta used by Catalog
    type ApplicationItemMeta struct {

        
            AltText map[string]interface{}  `json:"alt_text"`
            Moq ApplicationItemMOQ  `json:"moq"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            CustomMeta []MetaFields  `json:"_custom_meta"`
            Seo ApplicationItemSEO  `json:"seo"`
            IsGift bool  `json:"is_gift"`
            IsCod bool  `json:"is_cod"`
         
    }
    
    // SuccessResponse1 used by Catalog
    type SuccessResponse1 struct {

        
            UID float64  `json:"uid"`
            Success bool  `json:"success"`
         
    }
    
    // MOQData used by Catalog
    type MOQData struct {

        
            Minimum float64  `json:"minimum"`
            IncrementUnit float64  `json:"increment_unit"`
            Maximum float64  `json:"maximum"`
         
    }
    
    // SEOData used by Catalog
    type SEOData struct {

        
            Title string  `json:"title"`
            Description string  `json:"description"`
         
    }
    
    // OwnerAppItemResponse used by Catalog
    type OwnerAppItemResponse struct {

        
            AltText map[string]interface{}  `json:"alt_text"`
            Moq MOQData  `json:"moq"`
            IsGift bool  `json:"is_gift"`
            Seo SEOData  `json:"seo"`
            IsCod bool  `json:"is_cod"`
         
    }
    
    // GetConfigMetadataResponse used by Catalog
    type GetConfigMetadataResponse struct {

        
            Data []map[string]interface{}  `json:"data"`
            Condition []map[string]interface{}  `json:"condition"`
            Values []map[string]interface{}  `json:"values"`
         
    }
    
    // AttributeDetailsGroup used by Catalog
    type AttributeDetailsGroup struct {

        
            Logo string  `json:"logo"`
            DisplayType string  `json:"display_type"`
            Name string  `json:"name"`
            Key string  `json:"key"`
            Priority float64  `json:"priority"`
            Unit string  `json:"unit"`
            Slug string  `json:"slug"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // AppConfigurationDetail used by Catalog
    type AppConfigurationDetail struct {

        
            Logo string  `json:"logo"`
            AppID string  `json:"app_id"`
            IsDefault bool  `json:"is_default"`
            Name string  `json:"name"`
            TemplateSlugs []string  `json:"template_slugs"`
            Priority float64  `json:"priority"`
            Slug string  `json:"slug"`
            Attributes []AttributeDetailsGroup  `json:"attributes"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // ConfigErrorResponse used by Catalog
    type ConfigErrorResponse struct {

        
            Message string  `json:"message"`
         
    }
    
    // PageResponseType used by Catalog
    type PageResponseType struct {

        
            HasNext bool  `json:"has_next"`
            Next float64  `json:"next"`
            TotalCount float64  `json:"total_count"`
            Current float64  `json:"current"`
         
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

        
            Logo string  `json:"logo"`
            AppID string  `json:"app_id"`
            IsDefault bool  `json:"is_default"`
            Name string  `json:"name"`
            Key string  `json:"key"`
            Priority float64  `json:"priority"`
            DefaultKey string  `json:"default_key"`
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

        
            FilterTypes []string  `json:"filter_types"`
            Key string  `json:"key"`
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
    
    // GetCatalogConfigurationDetailsProduct used by Catalog
    type GetCatalogConfigurationDetailsProduct struct {

        
            Similar map[string]interface{}  `json:"similar"`
            Compare map[string]interface{}  `json:"compare"`
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

        
            Logo string  `json:"logo"`
            Name string  `json:"name"`
            Key string  `json:"key"`
            Priority float64  `json:"priority"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // ConfigurationListingSort used by Catalog
    type ConfigurationListingSort struct {

        
            Config []ConfigurationListingSortConfig  `json:"config"`
            DefaultKey string  `json:"default_key"`
         
    }
    
    // ConfigurationBucketPoints used by Catalog
    type ConfigurationBucketPoints struct {

        
            Start float64  `json:"start"`
            Display string  `json:"display"`
            End float64  `json:"end"`
         
    }
    
    // ConfigurationListingFilterValue used by Catalog
    type ConfigurationListingFilterValue struct {

        
            Sort string  `json:"sort"`
            BucketPoints []ConfigurationBucketPoints  `json:"bucket_points"`
            MapValues []map[string]interface{}  `json:"map_values"`
            Value string  `json:"value"`
            Priority []string  `json:"priority"`
            Condition string  `json:"condition"`
            Map map[string]interface{}  `json:"map"`
         
    }
    
    // ConfigurationListingFilterConfig used by Catalog
    type ConfigurationListingFilterConfig struct {

        
            Logo string  `json:"logo"`
            Type string  `json:"type"`
            Name string  `json:"name"`
            Key string  `json:"key"`
            ValueConfig ConfigurationListingFilterValue  `json:"value_config"`
            Priority float64  `json:"priority"`
            DisplayName string  `json:"display_name"`
            IsActive bool  `json:"is_active"`
         
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

        
            Logo string  `json:"logo"`
            Subtitle string  `json:"subtitle"`
            Size ProductSize  `json:"size"`
            Title string  `json:"title"`
            Key string  `json:"key"`
            Priority float64  `json:"priority"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // ConfigurationProductSimilar used by Catalog
    type ConfigurationProductSimilar struct {

        
            Config []ConfigurationProductConfig  `json:"config"`
         
    }
    
    // ConfigurationProductVariantConfig used by Catalog
    type ConfigurationProductVariantConfig struct {

        
            Logo string  `json:"logo"`
            Size ProductSize  `json:"size"`
            Name string  `json:"name"`
            Key string  `json:"key"`
            Priority float64  `json:"priority"`
            DisplayType string  `json:"display_type"`
            IsActive bool  `json:"is_active"`
         
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
    
    // AppConfiguration used by Catalog
    type AppConfiguration struct {

        
            Listing ConfigurationListing  `json:"listing"`
            AppID string  `json:"app_id"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Product ConfigurationProduct  `json:"product"`
            ModifiedOn string  `json:"modified_on"`
            ConfigType string  `json:"config_type"`
            ConfigID string  `json:"config_id"`
            CreatedOn string  `json:"created_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Type string  `json:"type"`
         
    }
    
    // AppCatalogConfiguration used by Catalog
    type AppCatalogConfiguration struct {

        
            Listing ConfigurationListing  `json:"listing"`
            ID string  `json:"id"`
            AppID string  `json:"app_id"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Product ConfigurationProduct  `json:"product"`
            ModifiedOn string  `json:"modified_on"`
            ConfigType string  `json:"config_type"`
            ConfigID string  `json:"config_id"`
            CreatedOn string  `json:"created_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Type string  `json:"type"`
         
    }
    
    // GetAppCatalogConfiguration used by Catalog
    type GetAppCatalogConfiguration struct {

        
            Data AppCatalogConfiguration  `json:"data"`
            IsDefault bool  `json:"is_default"`
         
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
            AppID string  `json:"app_id"`
            Product GetCatalogConfigurationDetailsProduct  `json:"product"`
            ConfigType string  `json:"config_type"`
            ConfigID string  `json:"config_id"`
         
    }
    
    // GetAppCatalogEntityConfiguration used by Catalog
    type GetAppCatalogEntityConfiguration struct {

        
            Data EntityConfiguration  `json:"data"`
            IsDefault bool  `json:"is_default"`
         
    }
    
    // ProductSortOn used by Catalog
    type ProductSortOn struct {

        
            Name string  `json:"name"`
            IsSelected bool  `json:"is_selected"`
            Value string  `json:"value"`
         
    }
    
    // ProductFiltersKey used by Catalog
    type ProductFiltersKey struct {

        
            Display string  `json:"display"`
            Logo string  `json:"logo"`
            Kind string  `json:"kind"`
            Name string  `json:"name"`
            Operators []string  `json:"operators"`
         
    }
    
    // ProductFiltersValue used by Catalog
    type ProductFiltersValue struct {

        
            Display string  `json:"display"`
            Count float64  `json:"count"`
            Min float64  `json:"min"`
            QueryFormat string  `json:"query_format"`
            SelectedMax float64  `json:"selected_max"`
            CurrencyCode string  `json:"currency_code"`
            Max float64  `json:"max"`
            IsSelected bool  `json:"is_selected"`
            Value interface{}  `json:"value"`
            DisplayFormat string  `json:"display_format"`
            SelectedMin float64  `json:"selected_min"`
            CurrencySymbol string  `json:"currency_symbol"`
         
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
    
    // UserInfo used by Catalog
    type UserInfo struct {

        
            Email string  `json:"email"`
            UID string  `json:"uid"`
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
         
    }
    
    // NextSchedule used by Catalog
    type NextSchedule struct {

        
            Start string  `json:"start"`
            End string  `json:"end"`
         
    }
    
    // CollectionSchedule used by Catalog
    type CollectionSchedule struct {

        
            Start string  `json:"start"`
            Cron string  `json:"cron"`
            NextSchedule []NextSchedule  `json:"next_schedule"`
            End string  `json:"end"`
            Duration float64  `json:"duration"`
         
    }
    
    // CollectionImage used by Catalog
    type CollectionImage struct {

        
            AspectRatio string  `json:"aspect_ratio"`
            URL string  `json:"url"`
         
    }
    
    // SeoDetail used by Catalog
    type SeoDetail struct {

        
            Title string  `json:"title"`
            Description string  `json:"description"`
         
    }
    
    // CollectionBadge used by Catalog
    type CollectionBadge struct {

        
            Text string  `json:"text"`
            Color string  `json:"color"`
         
    }
    
    // CollectionBanner used by Catalog
    type CollectionBanner struct {

        
            Portrait CollectionImage  `json:"portrait"`
            Landscape CollectionImage  `json:"landscape"`
         
    }
    
    // CollectionQuery used by Catalog
    type CollectionQuery struct {

        
            Attribute string  `json:"attribute"`
            Value []interface{}  `json:"value"`
            Op string  `json:"op"`
         
    }
    
    // CreateCollection used by Catalog
    type CreateCollection struct {

        
            Tags []string  `json:"tags"`
            IsVisible bool  `json:"is_visible"`
            Description string  `json:"description"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            AllowFacets bool  `json:"allow_facets"`
            CreatedBy UserInfo  `json:"created_by"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            IsActive bool  `json:"is_active"`
            Schedule CollectionSchedule  `json:"_schedule"`
            Logo CollectionImage  `json:"logo"`
            Type string  `json:"type"`
            AppID string  `json:"app_id"`
            SortOn string  `json:"sort_on"`
            Meta map[string]interface{}  `json:"meta"`
            Name string  `json:"name"`
            Seo SeoDetail  `json:"seo"`
            Priority float64  `json:"priority"`
            Slug string  `json:"slug"`
            Badge CollectionBadge  `json:"badge"`
            ModifiedBy UserInfo  `json:"modified_by"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Banners CollectionBanner  `json:"banners"`
            Published bool  `json:"published"`
            AllowSort bool  `json:"allow_sort"`
            Query []CollectionQuery  `json:"query"`
         
    }
    
    // BannerImage used by Catalog
    type BannerImage struct {

        
            AspectRatio string  `json:"aspect_ratio"`
            URL string  `json:"url"`
            Type string  `json:"type"`
         
    }
    
    // ImageUrls used by Catalog
    type ImageUrls struct {

        
            Landscape BannerImage  `json:"landscape"`
            Portrait BannerImage  `json:"portrait"`
         
    }
    
    // CollectionCreateResponse used by Catalog
    type CollectionCreateResponse struct {

        
            Query []CollectionQuery  `json:"query"`
            Description string  `json:"description"`
            Cron map[string]interface{}  `json:"cron"`
            Tag []string  `json:"tag"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            AllowFacets bool  `json:"allow_facets"`
            IsActive bool  `json:"is_active"`
            Schedule map[string]interface{}  `json:"_schedule"`
            Logo BannerImage  `json:"logo"`
            AppID string  `json:"app_id"`
            SortOn string  `json:"sort_on"`
            Meta map[string]interface{}  `json:"meta"`
            Name string  `json:"name"`
            Priority float64  `json:"priority"`
            Slug string  `json:"slug"`
            Badge map[string]interface{}  `json:"badge"`
            Banners ImageUrls  `json:"banners"`
            AllowSort bool  `json:"allow_sort"`
            Type string  `json:"type"`
         
    }
    
    // CollectionActionPageQuery used by Catalog
    type CollectionActionPageQuery struct {

        
            Attribute string  `json:"attribute"`
            Value []string  `json:"value"`
            Op string  `json:"op"`
         
    }
    
    // CollectionActionPage used by Catalog
    type CollectionActionPage struct {

        
            Type string  `json:"type"`
            Query CollectionActionPageQuery  `json:"query"`
         
    }
    
    // CollectionScheduleStartEnd used by Catalog
    type CollectionScheduleStartEnd struct {

        
            Start string  `json:"start"`
            End string  `json:"end"`
         
    }
    
    // CollectionSchedule1 used by Catalog
    type CollectionSchedule1 struct {

        
            Start string  `json:"start"`
            NextSchedule []CollectionScheduleStartEnd  `json:"next_schedule"`
            End string  `json:"end"`
         
    }
    
    // Media used by Catalog
    type Media struct {

        
            URL string  `json:"url"`
            Meta map[string]interface{}  `json:"meta"`
            Type string  `json:"type"`
         
    }
    
    // CollectionBadge1 used by Catalog
    type CollectionBadge1 struct {

        
            Text string  `json:"text"`
            Color string  `json:"color"`
         
    }
    
    // GetCollectionDetailNest used by Catalog
    type GetCollectionDetailNest struct {

        
            Tags []string  `json:"tags"`
            Query []CollectionQuery  `json:"query"`
            Description string  `json:"description"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Action CollectionActionPage  `json:"action"`
            AllowFacets bool  `json:"allow_facets"`
            IsActive bool  `json:"is_active"`
            Schedule CollectionSchedule1  `json:"_schedule"`
            Logo Media  `json:"logo"`
            AppID string  `json:"app_id"`
            Meta map[string]interface{}  `json:"meta"`
            Name string  `json:"name"`
            Priority float64  `json:"priority"`
            Slug string  `json:"slug"`
            Badge CollectionBadge1  `json:"badge"`
            UID string  `json:"uid"`
            Banners ImageUrls  `json:"banners"`
            AllowSort bool  `json:"allow_sort"`
            Type string  `json:"type"`
         
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

        
            Page Page  `json:"page"`
            Items []GetCollectionDetailNest  `json:"items"`
            Filters CollectionListingFilter  `json:"filters"`
         
    }
    
    // CollectionDetailResponse used by Catalog
    type CollectionDetailResponse struct {

        
            Schedule CollectionSchedule1  `json:"_schedule"`
            Logo Media  `json:"logo"`
            Tags []string  `json:"tags"`
            Type string  `json:"type"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            AppID string  `json:"app_id"`
            AllowFacets bool  `json:"allow_facets"`
            UID string  `json:"uid"`
            Query []CollectionQuery  `json:"query"`
            Meta map[string]interface{}  `json:"meta"`
            Description string  `json:"description"`
            Banners ImageUrls  `json:"banners"`
            Name string  `json:"name"`
            Priority float64  `json:"priority"`
            Slug string  `json:"slug"`
            Badge CollectionBadge1  `json:"badge"`
            AllowSort bool  `json:"allow_sort"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // UpdateCollection used by Catalog
    type UpdateCollection struct {

        
            Tags []string  `json:"tags"`
            IsVisible bool  `json:"is_visible"`
            Description string  `json:"description"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            AllowFacets bool  `json:"allow_facets"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            IsActive bool  `json:"is_active"`
            Schedule CollectionSchedule  `json:"_schedule"`
            Logo CollectionImage  `json:"logo"`
            Type string  `json:"type"`
            SortOn string  `json:"sort_on"`
            Meta map[string]interface{}  `json:"meta"`
            Name string  `json:"name"`
            Seo SeoDetail  `json:"seo"`
            Priority float64  `json:"priority"`
            Slug string  `json:"slug"`
            Badge CollectionBadge  `json:"badge"`
            ModifiedBy UserInfo  `json:"modified_by"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Banners CollectionBanner  `json:"banners"`
            Published bool  `json:"published"`
            AllowSort bool  `json:"allow_sort"`
            Query []CollectionQuery  `json:"query"`
         
    }
    
    // CollectionItem used by Catalog
    type CollectionItem struct {

        
            ItemID float64  `json:"item_id"`
            Action string  `json:"action"`
            Priority float64  `json:"priority"`
         
    }
    
    // CollectionItemUpdate used by Catalog
    type CollectionItemUpdate struct {

        
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            AllowFacets bool  `json:"allow_facets"`
            Query []CollectionQuery  `json:"query"`
            Items []CollectionItem  `json:"items"`
            AllowSort bool  `json:"allow_sort"`
            Type string  `json:"type"`
         
    }
    
    // UpdatedResponse used by Catalog
    type UpdatedResponse struct {

        
            ItemsNotUpdated []float64  `json:"items_not_updated"`
            Message string  `json:"message"`
         
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
            Action Action  `json:"action"`
            Logo Media  `json:"logo"`
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
    
    // ProductDetailAttribute used by Catalog
    type ProductDetailAttribute struct {

        
            Key string  `json:"key"`
            Value string  `json:"value"`
            Type string  `json:"type"`
         
    }
    
    // ProductDetailGroupedAttribute used by Catalog
    type ProductDetailGroupedAttribute struct {

        
            Details []ProductDetailAttribute  `json:"details"`
            Title string  `json:"title"`
         
    }
    
    // ProductListingDetail used by Catalog
    type ProductListingDetail struct {

        
            ImageNature string  `json:"image_nature"`
            Description string  `json:"description"`
            PromoMeta map[string]interface{}  `json:"promo_meta"`
            ItemType string  `json:"item_type"`
            ItemCode string  `json:"item_code"`
            Brand ProductBrand  `json:"brand"`
            Price ProductListingPrice  `json:"price"`
            Medias []Media  `json:"medias"`
            ShortDescription string  `json:"short_description"`
            ProductOnlineDate string  `json:"product_online_date"`
            RatingCount float64  `json:"rating_count"`
            Sellable bool  `json:"sellable"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            Name string  `json:"name"`
            Color string  `json:"color"`
            Slug string  `json:"slug"`
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
            UID float64  `json:"uid"`
            HasVariant bool  `json:"has_variant"`
            Rating float64  `json:"rating"`
            Tryouts []string  `json:"tryouts"`
            Highlights []string  `json:"highlights"`
            Similars []string  `json:"similars"`
            Discount string  `json:"discount"`
            Attributes map[string]interface{}  `json:"attributes"`
            Type string  `json:"type"`
         
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

        
            Count float64  `json:"count"`
            SellableCount float64  `json:"sellable_count"`
            OutOfStockCount float64  `json:"out_of_stock_count"`
         
    }
    
    // CatalogInsightBrand used by Catalog
    type CatalogInsightBrand struct {

        
            AvailableArticles float64  `json:"available_articles"`
            ArticleFreshness float64  `json:"article_freshness"`
            TotalArticles float64  `json:"total_articles"`
            Name string  `json:"name"`
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

        
            StoreIds []float64  `json:"store_ids"`
            CompanyID float64  `json:"company_id"`
            Platform string  `json:"platform"`
            BrandIds []float64  `json:"brand_ids"`
            Enabled bool  `json:"enabled"`
            OptLevel string  `json:"opt_level"`
         
    }
    
    // CompanyOptIn used by Catalog
    type CompanyOptIn struct {

        
            StoreIds []float64  `json:"store_ids"`
            CompanyID float64  `json:"company_id"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Platform string  `json:"platform"`
            BrandIds []float64  `json:"brand_ids"`
            Enabled bool  `json:"enabled"`
            ModifiedOn float64  `json:"modified_on"`
            CreatedOn float64  `json:"created_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            OptLevel string  `json:"opt_level"`
         
    }
    
    // GetOptInPlatform used by Catalog
    type GetOptInPlatform struct {

        
            Page Page  `json:"page"`
            Items []CompanyOptIn  `json:"items"`
         
    }
    
    // OptinCompanyDetail used by Catalog
    type OptinCompanyDetail struct {

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            CompanyType string  `json:"company_type"`
            BusinessType string  `json:"business_type"`
         
    }
    
    // CompanyBrandDetail used by Catalog
    type CompanyBrandDetail struct {

        
            CompanyID float64  `json:"company_id"`
            BrandName string  `json:"brand_name"`
            TotalArticle float64  `json:"total_article"`
            BrandID float64  `json:"brand_id"`
         
    }
    
    // OptinCompanyBrandDetailsView used by Catalog
    type OptinCompanyBrandDetailsView struct {

        
            Page Page  `json:"page"`
            Items []CompanyBrandDetail  `json:"items"`
         
    }
    
    // OptinCompanyMetrics used by Catalog
    type OptinCompanyMetrics struct {

        
            Company string  `json:"company"`
            Store float64  `json:"store"`
            Brand float64  `json:"brand"`
         
    }
    
    // StoreDetail used by Catalog
    type StoreDetail struct {

        
            StoreCode string  `json:"store_code"`
            Address map[string]interface{}  `json:"address"`
            Timing map[string]interface{}  `json:"timing"`
            Documents []map[string]interface{}  `json:"documents"`
            UID float64  `json:"uid"`
            CompanyID float64  `json:"company_id"`
            Name string  `json:"name"`
            StoreType string  `json:"store_type"`
            DisplayName string  `json:"display_name"`
            ModifiedOn string  `json:"modified_on"`
            AdditionalContacts []map[string]interface{}  `json:"additional_contacts"`
            CreatedOn string  `json:"created_on"`
            Manager map[string]interface{}  `json:"manager"`
         
    }
    
    // OptinStoreDetails used by Catalog
    type OptinStoreDetails struct {

        
            Page Page  `json:"page"`
            Items []StoreDetail  `json:"items"`
         
    }
    
    // AttributeMasterFilter used by Catalog
    type AttributeMasterFilter struct {

        
            DependsOn []string  `json:"depends_on"`
            Indexing bool  `json:"indexing"`
            Priority float64  `json:"priority"`
         
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
    
    // AttributeSchemaRange used by Catalog
    type AttributeSchemaRange struct {

        
            Min float64  `json:"min"`
            Max float64  `json:"max"`
         
    }
    
    // AttributeMaster used by Catalog
    type AttributeMaster struct {

        
            AllowedValues []string  `json:"allowed_values"`
            Multi bool  `json:"multi"`
            Range AttributeSchemaRange  `json:"range"`
            Mandatory bool  `json:"mandatory"`
            Format string  `json:"format"`
            Type string  `json:"type"`
         
    }
    
    // AttributeMasterDetails used by Catalog
    type AttributeMasterDetails struct {

        
            DisplayType string  `json:"display_type"`
         
    }
    
    // GenderDetail used by Catalog
    type GenderDetail struct {

        
            ID string  `json:"id"`
            Logo string  `json:"logo"`
            Filters AttributeMasterFilter  `json:"filters"`
            EnabledForEndConsumer bool  `json:"enabled_for_end_consumer"`
            Meta AttributeMasterMeta  `json:"meta"`
            Description string  `json:"description"`
            Name string  `json:"name"`
            Schema AttributeMaster  `json:"schema"`
            Details AttributeMasterDetails  `json:"details"`
            Slug string  `json:"slug"`
            IsNested bool  `json:"is_nested"`
            Departments []string  `json:"departments"`
         
    }
    
    // CategoriesResponse used by Catalog
    type CategoriesResponse struct {

        
            TemplateSlug string  `json:"template_slug"`
            SlugKey string  `json:"slug_key"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
         
    }
    
    // ProdcutTemplateCategoriesResponse used by Catalog
    type ProdcutTemplateCategoriesResponse struct {

        
            Page Page  `json:"page"`
            Items []CategoriesResponse  `json:"items"`
         
    }
    
    // PTErrorResponse used by Catalog
    type PTErrorResponse struct {

        
            Meta map[string]interface{}  `json:"meta"`
            Message string  `json:"message"`
            Errors map[string]interface{}  `json:"errors"`
            Code string  `json:"code"`
            Status float64  `json:"status"`
         
    }
    
    // DepartmentCreateUpdate used by Catalog
    type DepartmentCreateUpdate struct {

        
            Slug string  `json:"slug"`
            Logo string  `json:"logo"`
            Tags []string  `json:"tags"`
            UID float64  `json:"uid"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Synonyms []string  `json:"synonyms"`
            Name string  `json:"name"`
            PriorityOrder float64  `json:"priority_order"`
            Platforms map[string]interface{}  `json:"platforms"`
            Cls string  `json:"_cls"`
            IsActive bool  `json:"is_active"`
         
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
    
    // UserSerializer used by Catalog
    type UserSerializer struct {

        
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
            UID string  `json:"uid"`
            Contact string  `json:"contact"`
            ID string  `json:"_id"`
         
    }
    
    // GetDepartment used by Catalog
    type GetDepartment struct {

        
            Logo string  `json:"logo"`
            UID float64  `json:"uid"`
            ModifiedOn string  `json:"modified_on"`
            CreatedBy UserSerializer  `json:"created_by"`
            Synonyms []string  `json:"synonyms"`
            Name string  `json:"name"`
            Search string  `json:"search"`
            PriorityOrder float64  `json:"priority_order"`
            ItemType string  `json:"item_type"`
            PageSize float64  `json:"page_size"`
            PageNo float64  `json:"page_no"`
            Slug string  `json:"slug"`
            CreatedOn string  `json:"created_on"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // DepartmentsResponse used by Catalog
    type DepartmentsResponse struct {

        
            Page Page  `json:"page"`
            Items []GetDepartment  `json:"items"`
         
    }
    
    // DepartmentErrorResponse used by Catalog
    type DepartmentErrorResponse struct {

        
            Meta map[string]interface{}  `json:"meta"`
            Message string  `json:"message"`
            Errors map[string]interface{}  `json:"errors"`
            Code string  `json:"code"`
            Status float64  `json:"status"`
         
    }
    
    // UserDetail used by Catalog
    type UserDetail struct {

        
            SuperUser bool  `json:"super_user"`
            Contact string  `json:"contact"`
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
         
    }
    
    // DepartmentModel used by Catalog
    type DepartmentModel struct {

        
            Logo string  `json:"logo"`
            UID float64  `json:"uid"`
            Cls string  `json:"_cls"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            CreatedBy UserDetail  `json:"created_by"`
            Synonyms []string  `json:"synonyms"`
            VerifiedOn string  `json:"verified_on"`
            Name string  `json:"name"`
            PriorityOrder float64  `json:"priority_order"`
            VerifiedBy UserDetail  `json:"verified_by"`
            ID interface{}  `json:"_id"`
            ModifiedOn string  `json:"modified_on"`
            Slug string  `json:"slug"`
            CreatedOn string  `json:"created_on"`
            ModifiedBy UserDetail  `json:"modified_by"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // ProductTemplate used by Catalog
    type ProductTemplate struct {

        
            Logo string  `json:"logo"`
            IsPhysical bool  `json:"is_physical"`
            Categories []string  `json:"categories"`
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Description string  `json:"description"`
            Name string  `json:"name"`
            IsActive bool  `json:"is_active"`
            IsArchived bool  `json:"is_archived"`
            Tag string  `json:"tag"`
            Slug string  `json:"slug"`
            IsExpirable bool  `json:"is_expirable"`
            Attributes []string  `json:"attributes"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Departments []string  `json:"departments"`
         
    }
    
    // TemplatesResponse used by Catalog
    type TemplatesResponse struct {

        
            Page Page  `json:"page"`
            Items ProductTemplate  `json:"items"`
         
    }
    
    // Properties used by Catalog
    type Properties struct {

        
            Trader map[string]interface{}  `json:"trader"`
            Tags map[string]interface{}  `json:"tags"`
            TraderType map[string]interface{}  `json:"trader_type"`
            IsDependent map[string]interface{}  `json:"is_dependent"`
            BrandUID map[string]interface{}  `json:"brand_uid"`
            CategorySlug map[string]interface{}  `json:"category_slug"`
            Description map[string]interface{}  `json:"description"`
            ItemType map[string]interface{}  `json:"item_type"`
            ItemCode map[string]interface{}  `json:"item_code"`
            MultiSize map[string]interface{}  `json:"multi_size"`
            Currency map[string]interface{}  `json:"currency"`
            Command map[string]interface{}  `json:"command"`
            Media map[string]interface{}  `json:"media"`
            HsnCode map[string]interface{}  `json:"hsn_code"`
            ShortDescription map[string]interface{}  `json:"short_description"`
            IsActive map[string]interface{}  `json:"is_active"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            Name map[string]interface{}  `json:"name"`
            ProductPublish map[string]interface{}  `json:"product_publish"`
            Slug map[string]interface{}  `json:"slug"`
            NoOfBoxes map[string]interface{}  `json:"no_of_boxes"`
            Sizes map[string]interface{}  `json:"sizes"`
            SizeGuide map[string]interface{}  `json:"size_guide"`
            CountryOfOrigin map[string]interface{}  `json:"country_of_origin"`
            Highlights map[string]interface{}  `json:"highlights"`
            ReturnConfig map[string]interface{}  `json:"return_config"`
            ProductGroupTag map[string]interface{}  `json:"product_group_tag"`
            Variants map[string]interface{}  `json:"variants"`
         
    }
    
    // GlobalValidation used by Catalog
    type GlobalValidation struct {

        
            Definitions map[string]interface{}  `json:"definitions"`
            Required []string  `json:"required"`
            Properties Properties  `json:"properties"`
            Title string  `json:"title"`
            Description string  `json:"description"`
            Type string  `json:"type"`
         
    }
    
    // TemplateValidationData used by Catalog
    type TemplateValidationData struct {

        
            GlobalValidation GlobalValidation  `json:"global_validation"`
            TemplateValidation map[string]interface{}  `json:"template_validation"`
         
    }
    
    // TemplateDetails used by Catalog
    type TemplateDetails struct {

        
            ID string  `json:"id"`
            Logo string  `json:"logo"`
            IsPhysical bool  `json:"is_physical"`
            Categories []string  `json:"categories"`
            Description string  `json:"description"`
            Name string  `json:"name"`
            IsActive bool  `json:"is_active"`
            IsArchived bool  `json:"is_archived"`
            Tag string  `json:"tag"`
            Slug string  `json:"slug"`
            IsExpirable bool  `json:"is_expirable"`
            Attributes []string  `json:"attributes"`
            Departments []string  `json:"departments"`
         
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

        
            Filters map[string]interface{}  `json:"filters"`
            CompletedOn string  `json:"completed_on"`
            NotificationEmails []string  `json:"notification_emails"`
            CreatedBy UserInfo1  `json:"created_by"`
            TaskID string  `json:"task_id"`
            URL string  `json:"url"`
            SellerID float64  `json:"seller_id"`
            ModifiedOn string  `json:"modified_on"`
            Status string  `json:"status"`
            Type string  `json:"type"`
         
    }
    
    // ProductDownloadsResponse used by Catalog
    type ProductDownloadsResponse struct {

        
            Items []ProductTemplateExportResponse  `json:"items"`
         
    }
    
    // ProductTemplateExportFilterRequest used by Catalog
    type ProductTemplateExportFilterRequest struct {

        
            Templates []string  `json:"templates"`
            Brands []string  `json:"brands"`
            ToDate string  `json:"to_date"`
            FromDate string  `json:"from_date"`
            CatalogueTypes []string  `json:"catalogue_types"`
         
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
    
    // Hierarchy used by Catalog
    type Hierarchy struct {

        
            L1 float64  `json:"l1"`
            Department float64  `json:"department"`
            L2 float64  `json:"l2"`
         
    }
    
    // Media1 used by Catalog
    type Media1 struct {

        
            Logo string  `json:"logo"`
            Landscape string  `json:"landscape"`
            Portrait string  `json:"portrait"`
         
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
    
    // CategoryRequestBody used by Catalog
    type CategoryRequestBody struct {

        
            Hierarchy []Hierarchy  `json:"hierarchy"`
            Tryouts []string  `json:"tryouts"`
            Synonyms []string  `json:"synonyms"`
            Name string  `json:"name"`
            IsActive bool  `json:"is_active"`
            Media Media1  `json:"media"`
            Priority float64  `json:"priority"`
            Level float64  `json:"level"`
            Marketplaces CategoryMapping  `json:"marketplaces"`
            Slug string  `json:"slug"`
            Departments []float64  `json:"departments"`
         
    }
    
    // CategoryCreateResponse used by Catalog
    type CategoryCreateResponse struct {

        
            UID float64  `json:"uid"`
            Message string  `json:"message"`
         
    }
    
    // Category used by Catalog
    type Category struct {

        
            ID string  `json:"id"`
            UID float64  `json:"uid"`
            Hierarchy []Hierarchy  `json:"hierarchy"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Tryouts []string  `json:"tryouts"`
            Synonyms []string  `json:"synonyms"`
            Name string  `json:"name"`
            IsActive bool  `json:"is_active"`
            Media Media1  `json:"media"`
            Priority float64  `json:"priority"`
            Level float64  `json:"level"`
            ModifiedOn string  `json:"modified_on"`
            Marketplaces CategoryMapping  `json:"marketplaces"`
            Slug string  `json:"slug"`
            CreatedOn string  `json:"created_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Departments []float64  `json:"departments"`
         
    }
    
    // CategoryResponse used by Catalog
    type CategoryResponse struct {

        
            Page Page  `json:"page"`
            Items []Category  `json:"items"`
         
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
    
    // Trader used by Catalog
    type Trader struct {

        
            Name string  `json:"name"`
            Address []string  `json:"address"`
            Type string  `json:"type"`
         
    }
    
    // TaxIdentifier used by Catalog
    type TaxIdentifier struct {

        
            ReportingHsn string  `json:"reporting_hsn"`
            HsnCode string  `json:"hsn_code"`
            HsnCodeID string  `json:"hsn_code_id"`
         
    }
    
    // NetQuantity used by Catalog
    type NetQuantity struct {

        
            Value float64  `json:"value"`
            Unit string  `json:"unit"`
         
    }
    
    // CustomOrder used by Catalog
    type CustomOrder struct {

        
            IsCustomOrder bool  `json:"is_custom_order"`
            ManufacturingTime float64  `json:"manufacturing_time"`
            ManufacturingTimeUnit string  `json:"manufacturing_time_unit"`
         
    }
    
    // TeaserTag used by Catalog
    type TeaserTag struct {

        
            Tag string  `json:"tag"`
            URL string  `json:"url"`
         
    }
    
    // ProductPublish used by Catalog
    type ProductPublish struct {

        
            ProductOnlineDate string  `json:"product_online_date"`
            IsSet bool  `json:"is_set"`
         
    }
    
    // ReturnConfig used by Catalog
    type ReturnConfig struct {

        
            Time float64  `json:"time"`
            Returnable bool  `json:"returnable"`
            Unit string  `json:"unit"`
         
    }
    
    // ProductCreateUpdateSchemaV2 used by Catalog
    type ProductCreateUpdateSchemaV2 struct {

        
            Trader []Trader  `json:"trader"`
            Tags []string  `json:"tags"`
            ChangeRequestID interface{}  `json:"change_request_id"`
            IsDependent bool  `json:"is_dependent"`
            BrandUID float64  `json:"brand_uid"`
            CategorySlug string  `json:"category_slug"`
            Description string  `json:"description"`
            TaxIdentifier TaxIdentifier  `json:"tax_identifier"`
            ItemType string  `json:"item_type"`
            ItemCode string  `json:"item_code"`
            MultiSize bool  `json:"multi_size"`
            Action string  `json:"action"`
            Currency string  `json:"currency"`
            BulkJobID string  `json:"bulk_job_id"`
            NetQuantity NetQuantity  `json:"net_quantity"`
            Media []Media  `json:"media"`
            ShortDescription string  `json:"short_description"`
            VariantMedia map[string]interface{}  `json:"variant_media"`
            IsActive bool  `json:"is_active"`
            CustomOrder CustomOrder  `json:"custom_order"`
            CompanyID float64  `json:"company_id"`
            IsSet bool  `json:"is_set"`
            IsImageLessProduct bool  `json:"is_image_less_product"`
            Name string  `json:"name"`
            TeaserTag TeaserTag  `json:"teaser_tag"`
            ProductPublish ProductPublish  `json:"product_publish"`
            Requester string  `json:"requester"`
            Slug string  `json:"slug"`
            TemplateTag string  `json:"template_tag"`
            NoOfBoxes float64  `json:"no_of_boxes"`
            Sizes []map[string]interface{}  `json:"sizes"`
            Departments []float64  `json:"departments"`
            SizeGuide string  `json:"size_guide"`
            UID float64  `json:"uid"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            CountryOfOrigin string  `json:"country_of_origin"`
            Highlights []string  `json:"highlights"`
            ReturnConfig ReturnConfig  `json:"return_config"`
            VariantGroup map[string]interface{}  `json:"variant_group"`
            ProductGroupTag []string  `json:"product_group_tag"`
            Variants map[string]interface{}  `json:"variants"`
            Attributes map[string]interface{}  `json:"attributes"`
         
    }
    
    // Image used by Catalog
    type Image struct {

        
            AspectRatio string  `json:"aspect_ratio"`
            URL string  `json:"url"`
            AspectRatioF float64  `json:"aspect_ratio_f"`
            SecureURL string  `json:"secure_url"`
         
    }
    
    // Logo used by Catalog
    type Logo struct {

        
            AspectRatio string  `json:"aspect_ratio"`
            URL string  `json:"url"`
            AspectRatioF float64  `json:"aspect_ratio_f"`
            SecureURL string  `json:"secure_url"`
         
    }
    
    // Brand used by Catalog
    type Brand struct {

        
            Name string  `json:"name"`
            Logo Logo  `json:"logo"`
            UID float64  `json:"uid"`
         
    }
    
    // NetQuantityResponse used by Catalog
    type NetQuantityResponse struct {

        
            Value float64  `json:"value"`
            Unit string  `json:"unit"`
         
    }
    
    // ProductPublish1 used by Catalog
    type ProductPublish1 struct {

        
            ProductOnlineDate string  `json:"product_online_date"`
            IsSet bool  `json:"is_set"`
         
    }
    
    // ReturnConfigResponse used by Catalog
    type ReturnConfigResponse struct {

        
            Time float64  `json:"time"`
            Returnable bool  `json:"returnable"`
            Unit string  `json:"unit"`
         
    }
    
    // VerifiedBy used by Catalog
    type VerifiedBy struct {

        
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
         
    }
    
    // ProductSchemaV2 used by Catalog
    type ProductSchemaV2 struct {

        
            ImageNature string  `json:"image_nature"`
            ID string  `json:"id"`
            Images []Image  `json:"images"`
            IsPhysical bool  `json:"is_physical"`
            Tags []string  `json:"tags"`
            Trader []Trader  `json:"trader"`
            IsDependent bool  `json:"is_dependent"`
            BrandUID float64  `json:"brand_uid"`
            CategorySlug string  `json:"category_slug"`
            Description string  `json:"description"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            ItemType string  `json:"item_type"`
            ItemCode string  `json:"item_code"`
            Brand Brand  `json:"brand"`
            MultiSize bool  `json:"multi_size"`
            Category map[string]interface{}  `json:"category"`
            CreatedOn string  `json:"created_on"`
            AllCompanyIds []float64  `json:"all_company_ids"`
            Currency string  `json:"currency"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            NetQuantity NetQuantityResponse  `json:"net_quantity"`
            Stage string  `json:"stage"`
            Media []Media  `json:"media"`
            HsnCode string  `json:"hsn_code"`
            ShortDescription string  `json:"short_description"`
            IsExpirable bool  `json:"is_expirable"`
            VariantMedia map[string]interface{}  `json:"variant_media"`
            IsActive bool  `json:"is_active"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            Moq map[string]interface{}  `json:"moq"`
            CompanyID float64  `json:"company_id"`
            Pending string  `json:"pending"`
            IsSet bool  `json:"is_set"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            Name string  `json:"name"`
            IsImageLessProduct bool  `json:"is_image_less_product"`
            ProductPublish ProductPublish1  `json:"product_publish"`
            AllSizes []map[string]interface{}  `json:"all_sizes"`
            Color string  `json:"color"`
            L3Mapping []string  `json:"l3_mapping"`
            ModifiedOn string  `json:"modified_on"`
            Slug string  `json:"slug"`
            TemplateTag string  `json:"template_tag"`
            Sizes []map[string]interface{}  `json:"sizes"`
            NoOfBoxes float64  `json:"no_of_boxes"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Departments []float64  `json:"departments"`
            SizeGuide string  `json:"size_guide"`
            AllIdentifiers []string  `json:"all_identifiers"`
            UID float64  `json:"uid"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            PrimaryColor string  `json:"primary_color"`
            VerifiedOn string  `json:"verified_on"`
            CountryOfOrigin string  `json:"country_of_origin"`
            Highlights []string  `json:"highlights"`
            ReturnConfig ReturnConfigResponse  `json:"return_config"`
            VariantGroup map[string]interface{}  `json:"variant_group"`
            CategoryUID float64  `json:"category_uid"`
            ProductGroupTag []string  `json:"product_group_tag"`
            VerifiedBy VerifiedBy  `json:"verified_by"`
            Variants map[string]interface{}  `json:"variants"`
            Attributes map[string]interface{}  `json:"attributes"`
         
    }
    
    // ProductListingResponseV2 used by Catalog
    type ProductListingResponseV2 struct {

        
            Page Page  `json:"page"`
            Items []ProductSchemaV2  `json:"items"`
         
    }
    
    // ProductVariants used by Catalog
    type ProductVariants struct {

        
            UID float64  `json:"uid"`
            BrandUID float64  `json:"brand_uid"`
            Name string  `json:"name"`
            Media []Media  `json:"media"`
            ItemCode string  `json:"item_code"`
            CategoryUID float64  `json:"category_uid"`
         
    }
    
    // ProductVariantsResponse used by Catalog
    type ProductVariantsResponse struct {

        
            Page Page  `json:"page"`
            Variants []ProductVariants  `json:"variants"`
         
    }
    
    // AttributeMasterSerializer used by Catalog
    type AttributeMasterSerializer struct {

        
            Tags []string  `json:"tags"`
            EnabledForEndConsumer bool  `json:"enabled_for_end_consumer"`
            Description string  `json:"description"`
            Schema AttributeMaster  `json:"schema"`
            CreatedOn string  `json:"created_on"`
            Variant bool  `json:"variant"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Synonyms map[string]interface{}  `json:"synonyms"`
            Unit string  `json:"unit"`
            Example string  `json:"example"`
            Logo string  `json:"logo"`
            Filters AttributeMasterFilter  `json:"filters"`
            RawKey string  `json:"raw_key"`
            Name string  `json:"name"`
            Details AttributeMasterDetails  `json:"details"`
            Slug string  `json:"slug"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Departments []string  `json:"departments"`
            ModifiedOn string  `json:"modified_on"`
            IsNested bool  `json:"is_nested"`
            Suggestion string  `json:"suggestion"`
         
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
            ItemWeight float64  `json:"item_weight"`
            Size string  `json:"size"`
            ItemLength float64  `json:"item_length"`
            Identifiers []ValidateIdentifier  `json:"identifiers"`
            ItemDimensionsUnitOfMeasure string  `json:"item_dimensions_unit_of_measure"`
            ItemHeight float64  `json:"item_height"`
            ItemWidth float64  `json:"item_width"`
         
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

        
            ImageNature string  `json:"image_nature"`
            ID string  `json:"id"`
            Images []Image  `json:"images"`
            IsPhysical bool  `json:"is_physical"`
            Tags []string  `json:"tags"`
            Trader []Trader  `json:"trader"`
            IsDependent bool  `json:"is_dependent"`
            BrandUID float64  `json:"brand_uid"`
            CategorySlug string  `json:"category_slug"`
            Description string  `json:"description"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            ItemType string  `json:"item_type"`
            ItemCode string  `json:"item_code"`
            Brand Brand  `json:"brand"`
            MultiSize bool  `json:"multi_size"`
            Category map[string]interface{}  `json:"category"`
            CreatedOn string  `json:"created_on"`
            AllCompanyIds []float64  `json:"all_company_ids"`
            Currency string  `json:"currency"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            NetQuantity NetQuantityResponse  `json:"net_quantity"`
            Stage string  `json:"stage"`
            Media []Media  `json:"media"`
            HsnCode string  `json:"hsn_code"`
            ShortDescription string  `json:"short_description"`
            IsExpirable bool  `json:"is_expirable"`
            VariantMedia map[string]interface{}  `json:"variant_media"`
            IsActive bool  `json:"is_active"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            Moq map[string]interface{}  `json:"moq"`
            CompanyID float64  `json:"company_id"`
            Pending string  `json:"pending"`
            IsSet bool  `json:"is_set"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            Name string  `json:"name"`
            IsImageLessProduct bool  `json:"is_image_less_product"`
            ProductPublish ProductPublished  `json:"product_publish"`
            AllSizes []map[string]interface{}  `json:"all_sizes"`
            Color string  `json:"color"`
            L3Mapping []string  `json:"l3_mapping"`
            ModifiedOn string  `json:"modified_on"`
            Slug string  `json:"slug"`
            TemplateTag string  `json:"template_tag"`
            Sizes []map[string]interface{}  `json:"sizes"`
            NoOfBoxes float64  `json:"no_of_boxes"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Departments []float64  `json:"departments"`
            SizeGuide string  `json:"size_guide"`
            AllIdentifiers []string  `json:"all_identifiers"`
            UID float64  `json:"uid"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            PrimaryColor string  `json:"primary_color"`
            VerifiedOn string  `json:"verified_on"`
            CountryOfOrigin string  `json:"country_of_origin"`
            Highlights []string  `json:"highlights"`
            ReturnConfig ReturnConfigResponse  `json:"return_config"`
            VariantGroup map[string]interface{}  `json:"variant_group"`
            CategoryUID float64  `json:"category_uid"`
            ProductGroupTag []string  `json:"product_group_tag"`
            VerifiedBy VerifiedBy  `json:"verified_by"`
            Variants map[string]interface{}  `json:"variants"`
            Attributes map[string]interface{}  `json:"attributes"`
         
    }
    
    // ProductListingResponse used by Catalog
    type ProductListingResponse struct {

        
            Page Page  `json:"page"`
            Items []Product  `json:"items"`
         
    }
    
    // BulkJob used by Catalog
    type BulkJob struct {

        
            Succeed float64  `json:"succeed"`
            TrackingURL string  `json:"tracking_url"`
            FailedRecords []map[string]interface{}  `json:"failed_records"`
            FilePath string  `json:"file_path"`
            CustomTemplateTag string  `json:"custom_template_tag"`
            CompanyID float64  `json:"company_id"`
            Failed float64  `json:"failed"`
            CreatedBy UserInfo1  `json:"created_by"`
            Cancelled float64  `json:"cancelled"`
            Total float64  `json:"total"`
            CancelledRecords []map[string]interface{}  `json:"cancelled_records"`
            Stage string  `json:"stage"`
            ModifiedOn string  `json:"modified_on"`
            TemplateTag string  `json:"template_tag"`
            CreatedOn string  `json:"created_on"`
            ModifiedBy UserInfo1  `json:"modified_by"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // BulkResponse used by Catalog
    type BulkResponse struct {

        
            CreatedBy UserInfo1  `json:"created_by"`
            ModifiedOn string  `json:"modified_on"`
            BatchID string  `json:"batch_id"`
            CreatedOn string  `json:"created_on"`
            ModifiedBy UserInfo1  `json:"modified_by"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // UserDetail1 used by Catalog
    type UserDetail1 struct {

        
            FullName string  `json:"full_name"`
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
         
    }
    
    // ProductBulkRequest used by Catalog
    type ProductBulkRequest struct {

        
            Succeed float64  `json:"succeed"`
            FailedRecords []string  `json:"failed_records"`
            FilePath string  `json:"file_path"`
            CompanyID float64  `json:"company_id"`
            Failed float64  `json:"failed"`
            Template ProductTemplate  `json:"template"`
            CreatedBy UserDetail1  `json:"created_by"`
            Cancelled float64  `json:"cancelled"`
            Total float64  `json:"total"`
            CancelledRecords []string  `json:"cancelled_records"`
            Stage string  `json:"stage"`
            ModifiedOn string  `json:"modified_on"`
            TemplateTag string  `json:"template_tag"`
            CreatedOn string  `json:"created_on"`
            ModifiedBy UserDetail1  `json:"modified_by"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // ProductBulkRequestList used by Catalog
    type ProductBulkRequestList struct {

        
            Page Page  `json:"page"`
            Items ProductBulkRequest  `json:"items"`
         
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
    
    // ProductBulkAssets used by Catalog
    type ProductBulkAssets struct {

        
            CompanyID float64  `json:"company_id"`
            URL string  `json:"url"`
            BatchID string  `json:"batch_id"`
            User map[string]interface{}  `json:"user"`
         
    }
    
    // UserCommon used by Catalog
    type UserCommon struct {

        
            CompanyID float64  `json:"company_id"`
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
         
    }
    
    // Items used by Catalog
    type Items struct {

        
            ID string  `json:"id"`
            TrackingURL string  `json:"tracking_url"`
            FailedRecords []string  `json:"failed_records"`
            Failed float64  `json:"failed"`
            CompanyID float64  `json:"company_id"`
            CreatedOn string  `json:"created_on"`
            CreatedBy UserCommon  `json:"created_by"`
            Cancelled float64  `json:"cancelled"`
            Total float64  `json:"total"`
            CancelledRecords []string  `json:"cancelled_records"`
            IsActive bool  `json:"is_active"`
            Stage string  `json:"stage"`
            ModifiedOn string  `json:"modified_on"`
            Retry float64  `json:"retry"`
            FilePath string  `json:"file_path"`
            ModifiedBy UserCommon  `json:"modified_by"`
            Succeed float64  `json:"succeed"`
         
    }
    
    // BulkAssetResponse used by Catalog
    type BulkAssetResponse struct {

        
            Page Page  `json:"page"`
            Items []Items  `json:"items"`
         
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
    
    // ItemQuery used by Catalog
    type ItemQuery struct {

        
            UID float64  `json:"uid"`
            BrandUID float64  `json:"brand_uid"`
            ItemCode string  `json:"item_code"`
         
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

        
            Name string  `json:"name"`
            Quantity float64  `json:"quantity"`
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

        
            StoreCode string  `json:"store_code"`
            ItemWeightUnitOfMeasure string  `json:"item_weight_unit_of_measure"`
            Currency string  `json:"currency"`
            PriceEffective float64  `json:"price_effective"`
            Set InventorySet  `json:"set"`
            PriceTransfer float64  `json:"price_transfer"`
            ItemWeight float64  `json:"item_weight"`
            Size string  `json:"size"`
            ExpirationDate string  `json:"expiration_date"`
            ItemLength float64  `json:"item_length"`
            IsSet bool  `json:"is_set"`
            Quantity float64  `json:"quantity"`
            Identifiers []GTIN  `json:"identifiers"`
            ItemDimensionsUnitOfMeasure string  `json:"item_dimensions_unit_of_measure"`
            ItemHeight float64  `json:"item_height"`
            ItemWidth float64  `json:"item_width"`
            Price float64  `json:"price"`
         
    }
    
    // InventoryRequest used by Catalog
    type InventoryRequest struct {

        
            CompanyID float64  `json:"company_id"`
            Item ItemQuery  `json:"item"`
            Sizes []InvSize  `json:"sizes"`
         
    }
    
    // InventoryResponse used by Catalog
    type InventoryResponse struct {

        
            Currency string  `json:"currency"`
            InventoryUpdatedOn string  `json:"inventory_updated_on"`
            PriceEffective float64  `json:"price_effective"`
            SellerIdentifier string  `json:"seller_identifier"`
            PriceTransfer float64  `json:"price_transfer"`
            SellableQuantity float64  `json:"sellable_quantity"`
            Size string  `json:"size"`
            UID string  `json:"uid"`
            ItemID float64  `json:"item_id"`
            Quantity float64  `json:"quantity"`
            Identifiers map[string]interface{}  `json:"identifiers"`
            Store map[string]interface{}  `json:"store"`
            Price float64  `json:"price"`
         
    }
    
    // InventoryResponsePaginated used by Catalog
    type InventoryResponsePaginated struct {

        
            Page Page  `json:"page"`
            Items []InventoryResponse  `json:"items"`
         
    }
    
    // ProductSizeSellerFilter used by Catalog
    type ProductSizeSellerFilter struct {

        
            Name string  `json:"name"`
            IsSelected bool  `json:"is_selected"`
            Value string  `json:"value"`
         
    }
    
    // ArticleAssignment used by Catalog
    type ArticleAssignment struct {

        
            Strategy string  `json:"strategy"`
            Level string  `json:"level"`
         
    }
    
    // Seller used by Catalog
    type Seller struct {

        
            UID float64  `json:"uid"`
            Count float64  `json:"count"`
            Name string  `json:"name"`
         
    }
    
    // Store used by Catalog
    type Store struct {

        
            UID float64  `json:"uid"`
            Count float64  `json:"count"`
            Name string  `json:"name"`
         
    }
    
    // ProductSizePriceResponse used by Catalog
    type ProductSizePriceResponse struct {

        
            ArticleID string  `json:"article_id"`
            Pincode float64  `json:"pincode"`
            ArticleAssignment ArticleAssignment  `json:"article_assignment"`
            Seller Seller  `json:"seller"`
            Quantity float64  `json:"quantity"`
            Store Store  `json:"store"`
         
    }
    
    // ProductSizeSellersResponse used by Catalog
    type ProductSizeSellersResponse struct {

        
            Page Page  `json:"page"`
            SortOn []ProductSizeSellerFilter  `json:"sort_on"`
            Items []ProductSizePriceResponse  `json:"items"`
         
    }
    
    // CompanyMeta used by Catalog
    type CompanyMeta struct {

        
            ID float64  `json:"id"`
         
    }
    
    // Trader1 used by Catalog
    type Trader1 struct {

        
            Name string  `json:"name"`
            Address []string  `json:"address"`
            Type string  `json:"type"`
         
    }
    
    // BrandMeta used by Catalog
    type BrandMeta struct {

        
            Name string  `json:"name"`
            ID float64  `json:"id"`
         
    }
    
    // PriceMeta used by Catalog
    type PriceMeta struct {

        
            Transfer float64  `json:"transfer"`
            UpdatedAt string  `json:"updated_at"`
            Currency string  `json:"currency"`
            Marked float64  `json:"marked"`
            TpNotes map[string]interface{}  `json:"tp_notes"`
            Effective float64  `json:"effective"`
         
    }
    
    // DimensionResponse used by Catalog
    type DimensionResponse struct {

        
            Width float64  `json:"width"`
            Length float64  `json:"length"`
            IsDefault bool  `json:"is_default"`
            Unit string  `json:"unit"`
            Height float64  `json:"height"`
         
    }
    
    // WeightResponse used by Catalog
    type WeightResponse struct {

        
            Shipping float64  `json:"shipping"`
            IsDefault bool  `json:"is_default"`
            Unit string  `json:"unit"`
         
    }
    
    // QuantityBase used by Catalog
    type QuantityBase struct {

        
            Count float64  `json:"count"`
            UpdatedAt string  `json:"updated_at"`
         
    }
    
    // Quantities used by Catalog
    type Quantities struct {

        
            NotAvailable QuantityBase  `json:"not_available"`
            Sellable QuantityBase  `json:"sellable"`
            OrderCommitted QuantityBase  `json:"order_committed"`
            Damaged QuantityBase  `json:"damaged"`
         
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
    
    // ReturnConfig1 used by Catalog
    type ReturnConfig1 struct {

        
            Time float64  `json:"time"`
            Returnable bool  `json:"returnable"`
            Unit string  `json:"unit"`
         
    }
    
    // InventorySellerResponse used by Catalog
    type InventorySellerResponse struct {

        
            Company CompanyMeta  `json:"company"`
            Trader []Trader1  `json:"trader"`
            Tags []string  `json:"tags"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            Brand BrandMeta  `json:"brand"`
            FyndMeta map[string]interface{}  `json:"fynd_meta"`
            FyndItemCode string  `json:"fynd_item_code"`
            FyndArticleCode string  `json:"fynd_article_code"`
            Price PriceMeta  `json:"price"`
            Dimension DimensionResponse  `json:"dimension"`
            Size string  `json:"size"`
            CreatedBy UserSerializer  `json:"created_by"`
            ItemID float64  `json:"item_id"`
            Weight WeightResponse  `json:"weight"`
            Stage string  `json:"stage"`
            AddedOnStore string  `json:"added_on_store"`
            IsActive bool  `json:"is_active"`
            TotalQuantity float64  `json:"total_quantity"`
            TraceID string  `json:"trace_id"`
            Quantities Quantities  `json:"quantities"`
            Meta map[string]interface{}  `json:"meta"`
            IsSet bool  `json:"is_set"`
            TrackInventory bool  `json:"track_inventory"`
            Manufacturer ManufacturerResponse  `json:"manufacturer"`
            Store StoreMeta  `json:"store"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            SellerIdentifier string  `json:"seller_identifier"`
            Set InventorySet  `json:"set"`
            UID string  `json:"uid"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ExpirationDate string  `json:"expiration_date"`
            RawMeta map[string]interface{}  `json:"raw_meta"`
            Identifier map[string]interface{}  `json:"identifier"`
            ReturnConfig ReturnConfig1  `json:"return_config"`
            CountryOfOrigin string  `json:"country_of_origin"`
            Fragile bool  `json:"fragile"`
         
    }
    
    // InventorySellerIdentifierResponsePaginated used by Catalog
    type InventorySellerIdentifierResponsePaginated struct {

        
            Page Page  `json:"page"`
            Items []InventorySellerResponse  `json:"items"`
         
    }
    
    // CompanyMeta1 used by Catalog
    type CompanyMeta1 struct {

        
            ID float64  `json:"id"`
         
    }
    
    // DateMeta used by Catalog
    type DateMeta struct {

        
            ModifiedOn string  `json:"modified_on"`
            AddedOnStore string  `json:"added_on_store"`
            CreatedOn string  `json:"created_on"`
            InventoryUpdatedOn string  `json:"inventory_updated_on"`
         
    }
    
    // Trader2 used by Catalog
    type Trader2 struct {

        
            Name string  `json:"name"`
            Address []string  `json:"address"`
            Type string  `json:"type"`
         
    }
    
    // BrandMeta1 used by Catalog
    type BrandMeta1 struct {

        
            Name string  `json:"name"`
            ID float64  `json:"id"`
         
    }
    
    // PriceArticle used by Catalog
    type PriceArticle struct {

        
            Transfer float64  `json:"transfer"`
            Currency string  `json:"currency"`
            Marked float64  `json:"marked"`
            TpNotes map[string]interface{}  `json:"tp_notes"`
            Effective float64  `json:"effective"`
         
    }
    
    // DimensionResponse1 used by Catalog
    type DimensionResponse1 struct {

        
            Length float64  `json:"length"`
            Width float64  `json:"width"`
            Height float64  `json:"height"`
            Unit string  `json:"unit"`
         
    }
    
    // WeightResponse1 used by Catalog
    type WeightResponse1 struct {

        
            Shipping float64  `json:"shipping"`
            Unit string  `json:"unit"`
         
    }
    
    // Quantity used by Catalog
    type Quantity struct {

        
            Count float64  `json:"count"`
         
    }
    
    // QuantitiesArticle used by Catalog
    type QuantitiesArticle struct {

        
            NotAvailable Quantity  `json:"not_available"`
            Sellable Quantity  `json:"sellable"`
            OrderCommitted Quantity  `json:"order_committed"`
            Damaged Quantity  `json:"damaged"`
         
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
    
    // ReturnConfig2 used by Catalog
    type ReturnConfig2 struct {

        
            Time float64  `json:"time"`
            Returnable bool  `json:"returnable"`
            Unit string  `json:"unit"`
         
    }
    
    // GetInventories used by Catalog
    type GetInventories struct {

        
            ID string  `json:"id"`
            Company CompanyMeta1  `json:"company"`
            DateMeta DateMeta  `json:"date_meta"`
            Tags []string  `json:"tags"`
            Trader []Trader2  `json:"trader"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            Brand BrandMeta1  `json:"brand"`
            Platforms map[string]interface{}  `json:"platforms"`
            Price PriceArticle  `json:"price"`
            Dimension DimensionResponse1  `json:"dimension"`
            Size string  `json:"size"`
            CreatedBy UserSerializer  `json:"created_by"`
            ItemID float64  `json:"item_id"`
            Weight WeightResponse1  `json:"weight"`
            Stage string  `json:"stage"`
            InventoryUpdatedOn string  `json:"inventory_updated_on"`
            TotalQuantity float64  `json:"total_quantity"`
            Quantities QuantitiesArticle  `json:"quantities"`
            TraceID string  `json:"trace_id"`
            IsSet bool  `json:"is_set"`
            TrackInventory bool  `json:"track_inventory"`
            Manufacturer ManufacturerResponse1  `json:"manufacturer"`
            Store ArticleStoreResponse  `json:"store"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            SellerIdentifier string  `json:"seller_identifier"`
            UID string  `json:"uid"`
            ExpirationDate string  `json:"expiration_date"`
            Identifier map[string]interface{}  `json:"identifier"`
            ReturnConfig ReturnConfig2  `json:"return_config"`
            CountryOfOrigin string  `json:"country_of_origin"`
         
    }
    
    // GetInventoriesResponse used by Catalog
    type GetInventoriesResponse struct {

        
            Page Page  `json:"page"`
            Items []GetInventories  `json:"items"`
         
    }
    
    // BulkInventoryGetItems used by Catalog
    type BulkInventoryGetItems struct {

        
            ID string  `json:"id"`
            Succeed float64  `json:"succeed"`
            FailedRecords []string  `json:"failed_records"`
            FilePath string  `json:"file_path"`
            CompanyID float64  `json:"company_id"`
            Failed float64  `json:"failed"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Cancelled float64  `json:"cancelled"`
            Total float64  `json:"total"`
            CancelledRecords []string  `json:"cancelled_records"`
            Stage string  `json:"stage"`
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // BulkInventoryGet used by Catalog
    type BulkInventoryGet struct {

        
            Page Page  `json:"page"`
            Items []BulkInventoryGetItems  `json:"items"`
         
    }
    
    // InventoryJobPayload used by Catalog
    type InventoryJobPayload struct {

        
            StoreCode string  `json:"store_code"`
            ItemWeightUnitOfMeasure string  `json:"item_weight_unit_of_measure"`
            Tags []string  `json:"tags"`
            SellerIdentifier string  `json:"seller_identifier"`
            Currency string  `json:"currency"`
            PriceEffective float64  `json:"price_effective"`
            ExpirationDate string  `json:"expiration_date"`
            TotalQuantity float64  `json:"total_quantity"`
            TraceID string  `json:"trace_id"`
            PriceMarked float64  `json:"price_marked"`
            Quantity float64  `json:"quantity"`
            ItemDimensionsUnitOfMeasure string  `json:"item_dimensions_unit_of_measure"`
            Price float64  `json:"price"`
         
    }
    
    // InventoryBulkRequest used by Catalog
    type InventoryBulkRequest struct {

        
            CompanyID float64  `json:"company_id"`
            Sizes []InventoryJobPayload  `json:"sizes"`
            BatchID string  `json:"batch_id"`
            User map[string]interface{}  `json:"user"`
         
    }
    
    // InventoryExportRequest used by Catalog
    type InventoryExportRequest struct {

        
            Brand []float64  `json:"brand"`
            Store []float64  `json:"store"`
            Type string  `json:"type"`
         
    }
    
    // InventoryExportResponse used by Catalog
    type InventoryExportResponse struct {

        
            Status string  `json:"status"`
            Filters map[string]interface{}  `json:"filters"`
            NotificationEmails []string  `json:"notification_emails"`
            CreatedBy string  `json:"created_by"`
            SellerID float64  `json:"seller_id"`
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
            TaskID string  `json:"task_id"`
            Type string  `json:"type"`
         
    }
    
    // InventoryExportQuantityFilter used by Catalog
    type InventoryExportQuantityFilter struct {

        
            Operators string  `json:"operators"`
            Min float64  `json:"min"`
            Max float64  `json:"max"`
         
    }
    
    // InventoryExportAdvanceOption used by Catalog
    type InventoryExportAdvanceOption struct {

        
            StoreIds []float64  `json:"store_ids"`
            ToDate string  `json:"to_date"`
            Quantity InventoryExportQuantityFilter  `json:"quantity"`
            BrandIds []float64  `json:"brand_ids"`
            FromDate string  `json:"from_date"`
         
    }
    
    // InventoryExportJob used by Catalog
    type InventoryExportJob struct {

        
            Status string  `json:"status"`
            Filters InventoryExportAdvanceOption  `json:"filters"`
            CompletedOn string  `json:"completed_on"`
            NotificationEmails []string  `json:"notification_emails"`
            URL string  `json:"url"`
            SellerID float64  `json:"seller_id"`
            TaskID string  `json:"task_id"`
            Type string  `json:"type"`
         
    }
    
    // InventoryExportFilter used by Catalog
    type InventoryExportFilter struct {

        
            StoreIds []float64  `json:"store_ids"`
            ToDate string  `json:"to_date"`
            Quantity InventoryExportQuantityFilter  `json:"quantity"`
            BrandIds []float64  `json:"brand_ids"`
            FromDate string  `json:"from_date"`
         
    }
    
    // InventoryCreateRequest used by Catalog
    type InventoryCreateRequest struct {

        
            Data []string  `json:"data"`
            NotificationEmails []string  `json:"notification_emails"`
            Filters InventoryExportFilter  `json:"filters"`
            Type string  `json:"type"`
         
    }
    
    // InventoryJobFilters used by Catalog
    type InventoryJobFilters struct {

        
            Brands []string  `json:"brands"`
            ToDate string  `json:"to_date"`
            Quantity InventoryExportQuantityFilter  `json:"quantity"`
            FromDate string  `json:"from_date"`
            Stores []string  `json:"stores"`
         
    }
    
    // InventoryJobDetailResponse used by Catalog
    type InventoryJobDetailResponse struct {

        
            ID string  `json:"id"`
            Type string  `json:"type"`
            Status string  `json:"status"`
            Filters InventoryJobFilters  `json:"filters"`
            CompletedOn string  `json:"completed_on"`
            NotificationEmails []string  `json:"notification_emails"`
            CreatedBy UserDetail  `json:"created_by"`
            URL string  `json:"url"`
            SellerID float64  `json:"seller_id"`
            ModifiedOn string  `json:"modified_on"`
            CancelledOn string  `json:"cancelled_on"`
            CreatedOn string  `json:"created_on"`
            TaskID string  `json:"task_id"`
            CancelledBy UserDetail  `json:"cancelled_by"`
         
    }
    
    // InventoryExportJobListResponse used by Catalog
    type InventoryExportJobListResponse struct {

        
            Items InventoryJobDetailResponse  `json:"items"`
         
    }
    
    // FilerList used by Catalog
    type FilerList struct {

        
            Display string  `json:"display"`
            Value string  `json:"value"`
         
    }
    
    // InventoryConfig used by Catalog
    type InventoryConfig struct {

        
            Data []FilerList  `json:"data"`
            Multivalues bool  `json:"multivalues"`
         
    }
    
    // InventoryPayload used by Catalog
    type InventoryPayload struct {

        
            StoreID float64  `json:"store_id"`
            Tags []string  `json:"tags"`
            SellerIdentifier string  `json:"seller_identifier"`
            PriceEffective float64  `json:"price_effective"`
            ExpirationDate string  `json:"expiration_date"`
            TotalQuantity float64  `json:"total_quantity"`
            TraceID string  `json:"trace_id"`
            PriceMarked float64  `json:"price_marked"`
         
    }
    
    // InventoryRequestSchemaV2 used by Catalog
    type InventoryRequestSchemaV2 struct {

        
            CompanyID float64  `json:"company_id"`
            Payload []InventoryPayload  `json:"payload"`
            Meta map[string]interface{}  `json:"meta"`
         
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

        
            Tax1 float64  `json:"tax1"`
            CompanyID float64  `json:"company_id"`
            UID float64  `json:"uid"`
            Tax2 float64  `json:"tax2"`
            TaxOnEsp bool  `json:"tax_on_esp"`
            HsnCode string  `json:"hsn_code"`
            Hs2Code string  `json:"hs2_code"`
            Threshold1 float64  `json:"threshold1"`
            TaxOnMrp bool  `json:"tax_on_mrp"`
            Threshold2 float64  `json:"threshold2"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // HsnCodesObject used by Catalog
    type HsnCodesObject struct {

        
            ID string  `json:"id"`
            Tax1 float64  `json:"tax1"`
            CompanyID float64  `json:"company_id"`
            Tax2 float64  `json:"tax2"`
            TaxOnEsp bool  `json:"tax_on_esp"`
            HsnCode string  `json:"hsn_code"`
            ModifiedOn string  `json:"modified_on"`
            Hs2Code string  `json:"hs2_code"`
            Threshold1 float64  `json:"threshold1"`
            TaxOnMrp bool  `json:"tax_on_mrp"`
            Threshold2 float64  `json:"threshold2"`
         
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
    
    // PageResponse used by Catalog
    type PageResponse struct {

        
            HasNext bool  `json:"has_next"`
            ItemTotal float64  `json:"item_total"`
            Size float64  `json:"size"`
            HasPrevious bool  `json:"has_previous"`
            Current string  `json:"current"`
         
    }
    
    // TaxSlab used by Catalog
    type TaxSlab struct {

        
            EffectiveDate string  `json:"effective_date"`
            Rate float64  `json:"rate"`
            Threshold float64  `json:"threshold"`
            Cess float64  `json:"cess"`
         
    }
    
    // HSNDataInsertV2 used by Catalog
    type HSNDataInsertV2 struct {

        
            CountryCode string  `json:"country_code"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Taxes []TaxSlab  `json:"taxes"`
            Description string  `json:"description"`
            HsnCodeID string  `json:"hsn_code_id"`
            HsnCode string  `json:"hsn_code"`
            ModifiedOn string  `json:"modified_on"`
            ReportingHsn string  `json:"reporting_hsn"`
            CreatedOn string  `json:"created_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Type string  `json:"type"`
         
    }
    
    // HsnCodesListingResponseSchemaV2 used by Catalog
    type HsnCodesListingResponseSchemaV2 struct {

        
            Page PageResponse  `json:"page"`
            Items []HSNDataInsertV2  `json:"items"`
         
    }
    
    // Media2 used by Catalog
    type Media2 struct {

        
            AspectRatio string  `json:"aspect_ratio"`
            URL string  `json:"url"`
            Type string  `json:"type"`
         
    }
    
    // BrandItem used by Catalog
    type BrandItem struct {

        
            Action Action  `json:"action"`
            Logo Media2  `json:"logo"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            Banners ImageUrls  `json:"banners"`
            Discount string  `json:"discount"`
            Slug string  `json:"slug"`
            Departments []string  `json:"departments"`
         
    }
    
    // BrandListingResponse used by Catalog
    type BrandListingResponse struct {

        
            Page Page  `json:"page"`
            Items []BrandItem  `json:"items"`
         
    }
    
    // Department used by Catalog
    type Department struct {

        
            Logo Media2  `json:"logo"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            PriorityOrder float64  `json:"priority_order"`
            Slug string  `json:"slug"`
         
    }
    
    // DepartmentResponse used by Catalog
    type DepartmentResponse struct {

        
            Items []Department  `json:"items"`
         
    }
    
    // ThirdLevelChild used by Catalog
    type ThirdLevelChild struct {

        
            Action Action  `json:"action"`
            UID float64  `json:"uid"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Name string  `json:"name"`
            Banners ImageUrls  `json:"banners"`
            Childs []map[string]interface{}  `json:"childs"`
            Slug string  `json:"slug"`
         
    }
    
    // SecondLevelChild used by Catalog
    type SecondLevelChild struct {

        
            Action Action  `json:"action"`
            UID float64  `json:"uid"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Name string  `json:"name"`
            Banners ImageUrls  `json:"banners"`
            Childs []ThirdLevelChild  `json:"childs"`
            Slug string  `json:"slug"`
         
    }
    
    // Child used by Catalog
    type Child struct {

        
            Action Action  `json:"action"`
            UID float64  `json:"uid"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Name string  `json:"name"`
            Banners ImageUrls  `json:"banners"`
            Childs []SecondLevelChild  `json:"childs"`
            Slug string  `json:"slug"`
         
    }
    
    // CategoryItems used by Catalog
    type CategoryItems struct {

        
            Action Action  `json:"action"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            Banners ImageUrls  `json:"banners"`
            Childs []Child  `json:"childs"`
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

        
            Filters []ProductFilters  `json:"filters"`
            Page Page  `json:"page"`
            SortOn []ProductSortOn  `json:"sort_on"`
            Items []ProductListingDetail  `json:"items"`
            Operators map[string]interface{}  `json:"operators"`
         
    }
    
    // ProductDetail used by Catalog
    type ProductDetail struct {

        
            ImageNature string  `json:"image_nature"`
            Description string  `json:"description"`
            PromoMeta map[string]interface{}  `json:"promo_meta"`
            ItemType string  `json:"item_type"`
            ItemCode string  `json:"item_code"`
            Brand ProductBrand  `json:"brand"`
            Medias []Media  `json:"medias"`
            ShortDescription string  `json:"short_description"`
            ProductOnlineDate string  `json:"product_online_date"`
            RatingCount float64  `json:"rating_count"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            Name string  `json:"name"`
            Color string  `json:"color"`
            Slug string  `json:"slug"`
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
            UID float64  `json:"uid"`
            HasVariant bool  `json:"has_variant"`
            Rating float64  `json:"rating"`
            Tryouts []string  `json:"tryouts"`
            Highlights []string  `json:"highlights"`
            Similars []string  `json:"similars"`
            Attributes map[string]interface{}  `json:"attributes"`
            Type string  `json:"type"`
         
    }
    
    // InventoryPage used by Catalog
    type InventoryPage struct {

        
            HasNext bool  `json:"has_next"`
            ItemTotal float64  `json:"item_total"`
            HasPrevious bool  `json:"has_previous"`
            NextID string  `json:"next_id"`
            Type string  `json:"type"`
         
    }
    
    // InventoryStockResponse used by Catalog
    type InventoryStockResponse struct {

        
            Page InventoryPage  `json:"page"`
            Items []map[string]interface{}  `json:"items"`
         
    }
    
    // ArticleAssignment1 used by Catalog
    type ArticleAssignment1 struct {

        
            Strategy string  `json:"strategy"`
            Level string  `json:"level"`
         
    }
    
    // ArticleQuery used by Catalog
    type ArticleQuery struct {

        
            ItemID float64  `json:"item_id"`
            IgnoredStores []float64  `json:"ignored_stores"`
            Size string  `json:"size"`
         
    }
    
    // AssignStoreArticle used by Catalog
    type AssignStoreArticle struct {

        
            ArticleAssignment ArticleAssignment1  `json:"article_assignment"`
            Meta map[string]interface{}  `json:"meta"`
            GroupID string  `json:"group_id"`
            Quantity float64  `json:"quantity"`
            Query ArticleQuery  `json:"query"`
         
    }
    
    // AssignStore used by Catalog
    type AssignStore struct {

        
            Pincode string  `json:"pincode"`
            StoreIds []float64  `json:"store_ids"`
            CompanyID float64  `json:"company_id"`
            Articles []AssignStoreArticle  `json:"articles"`
            AppID string  `json:"app_id"`
            ChannelIdentifier string  `json:"channel_identifier"`
            ChannelType string  `json:"channel_type"`
         
    }
    
    // ArticleAssignment2 used by Catalog
    type ArticleAssignment2 struct {

        
            Strategy string  `json:"strategy"`
            Level string  `json:"level"`
         
    }
    
    // StoreAssignResponse used by Catalog
    type StoreAssignResponse struct {

        
            StoreID float64  `json:"store_id"`
            SCity string  `json:"s_city"`
            PriceEffective float64  `json:"price_effective"`
            UID string  `json:"uid"`
            CompanyID float64  `json:"company_id"`
            Size string  `json:"size"`
            ArticleAssignment ArticleAssignment2  `json:"article_assignment"`
            Meta map[string]interface{}  `json:"meta"`
            PriceMarked float64  `json:"price_marked"`
            Index float64  `json:"index"`
            ItemID float64  `json:"item_id"`
            StorePincode float64  `json:"store_pincode"`
            Quantity float64  `json:"quantity"`
            GroupID string  `json:"group_id"`
            ID string  `json:"_id"`
            StrategyWiseListing []map[string]interface{}  `json:"strategy_wise_listing"`
            Status bool  `json:"status"`
         
    }
    
    // GetAddressSerializer used by Catalog
    type GetAddressSerializer struct {

        
            Country string  `json:"country"`
            CountryCode string  `json:"country_code"`
            Pincode float64  `json:"pincode"`
            Address2 string  `json:"address2"`
            State string  `json:"state"`
            AddressType string  `json:"address_type"`
            City string  `json:"city"`
            Address1 string  `json:"address1"`
            Longitude float64  `json:"longitude"`
            Latitude float64  `json:"latitude"`
            Landmark string  `json:"landmark"`
         
    }
    
    // UserSerializer1 used by Catalog
    type UserSerializer1 struct {

        
            Contact string  `json:"contact"`
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
         
    }
    
    // GetCompanySerializer used by Catalog
    type GetCompanySerializer struct {

        
            Addresses []GetAddressSerializer  `json:"addresses"`
            UID float64  `json:"uid"`
            CreatedBy UserSerializer1  `json:"created_by"`
            VerifiedOn string  `json:"verified_on"`
            Name string  `json:"name"`
            Stage string  `json:"stage"`
            RejectReason string  `json:"reject_reason"`
            CompanyType string  `json:"company_type"`
            VerifiedBy UserSerializer1  `json:"verified_by"`
            BusinessType string  `json:"business_type"`
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
            ModifiedBy UserSerializer1  `json:"modified_by"`
         
    }
    
    // ProductReturnConfigSerializer used by Catalog
    type ProductReturnConfigSerializer struct {

        
            StoreUID float64  `json:"store_uid"`
            OnSameStore bool  `json:"on_same_store"`
         
    }
    
    // LocationTimingSerializer used by Catalog
    type LocationTimingSerializer struct {

        
            Minute float64  `json:"minute"`
            Hour float64  `json:"hour"`
         
    }
    
    // LocationDayWiseSerializer used by Catalog
    type LocationDayWiseSerializer struct {

        
            Closing LocationTimingSerializer  `json:"closing"`
            Weekday string  `json:"weekday"`
            Opening LocationTimingSerializer  `json:"opening"`
            Open bool  `json:"open"`
         
    }
    
    // UserSerializer2 used by Catalog
    type UserSerializer2 struct {

        
            Contact string  `json:"contact"`
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
         
    }
    
    // SellerPhoneNumber used by Catalog
    type SellerPhoneNumber struct {

        
            Number string  `json:"number"`
            CountryCode float64  `json:"country_code"`
         
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

        
            URL string  `json:"url"`
            LegalName string  `json:"legal_name"`
            Value string  `json:"value"`
            Verified bool  `json:"verified"`
            Type string  `json:"type"`
         
    }
    
    // LocationIntegrationType used by Catalog
    type LocationIntegrationType struct {

        
            Order string  `json:"order"`
            Inventory string  `json:"inventory"`
         
    }
    
    // LocationManagerSerializer used by Catalog
    type LocationManagerSerializer struct {

        
            Name string  `json:"name"`
            Email string  `json:"email"`
            MobileNo SellerPhoneNumber  `json:"mobile_no"`
         
    }
    
    // GetLocationSerializer used by Catalog
    type GetLocationSerializer struct {

        
            Company GetCompanySerializer  `json:"company"`
            PhoneNumber string  `json:"phone_number"`
            DisplayName string  `json:"display_name"`
            Code string  `json:"code"`
            CreatedOn string  `json:"created_on"`
            ProductReturnConfig ProductReturnConfigSerializer  `json:"product_return_config"`
            Address GetAddressSerializer  `json:"address"`
            Timing []LocationDayWiseSerializer  `json:"timing"`
            CreatedBy UserSerializer2  `json:"created_by"`
            Stage string  `json:"stage"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            Warnings map[string]interface{}  `json:"warnings"`
            NotificationEmails []string  `json:"notification_emails"`
            Name string  `json:"name"`
            ModifiedBy UserSerializer2  `json:"modified_by"`
            GstCredentials InvoiceDetailsSerializer  `json:"gst_credentials"`
            UID float64  `json:"uid"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Documents []Document  `json:"documents"`
            IntegrationType LocationIntegrationType  `json:"integration_type"`
            VerifiedOn string  `json:"verified_on"`
            StoreType string  `json:"store_type"`
            VerifiedBy UserSerializer2  `json:"verified_by"`
            ModifiedOn string  `json:"modified_on"`
            Manager LocationManagerSerializer  `json:"manager"`
         
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

        
            Logo string  `json:"logo"`
            AppID string  `json:"app_id"`
            UID float64  `json:"uid"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Name string  `json:"name"`
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
    

    
    // BusinessCountryInfo used by CompanyProfile
    type BusinessCountryInfo struct {

        
            CountryCode string  `json:"country_code"`
            Country string  `json:"country"`
         
    }
    
    // UserSerializer used by CompanyProfile
    type UserSerializer struct {

        
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
            Contact string  `json:"contact"`
         
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

        
            CountryCode float64  `json:"country_code"`
            Number string  `json:"number"`
         
    }
    
    // ContactDetails used by CompanyProfile
    type ContactDetails struct {

        
            Emails []string  `json:"emails"`
            Phone []SellerPhoneNumber  `json:"phone"`
         
    }
    
    // GetAddressSerializer used by CompanyProfile
    type GetAddressSerializer struct {

        
            State string  `json:"state"`
            Landmark string  `json:"landmark"`
            City string  `json:"city"`
            Longitude float64  `json:"longitude"`
            Latitude float64  `json:"latitude"`
            CountryCode string  `json:"country_code"`
            Pincode float64  `json:"pincode"`
            Country string  `json:"country"`
            AddressType string  `json:"address_type"`
            Address2 string  `json:"address2"`
            Address1 string  `json:"address1"`
         
    }
    
    // Document used by CompanyProfile
    type Document struct {

        
            Value string  `json:"value"`
            LegalName string  `json:"legal_name"`
            URL string  `json:"url"`
            Type string  `json:"type"`
            Verified bool  `json:"verified"`
         
    }
    
    // CompanyTaxesSerializer used by CompanyProfile
    type CompanyTaxesSerializer struct {

        
            Rate float64  `json:"rate"`
            Enable bool  `json:"enable"`
            EffectiveDate string  `json:"effective_date"`
         
    }
    
    // GetCompanyProfileSerializerResponse used by CompanyProfile
    type GetCompanyProfileSerializerResponse struct {

        
            BusinessCountryInfo BusinessCountryInfo  `json:"business_country_info"`
            ModifiedOn string  `json:"modified_on"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            CompanyType string  `json:"company_type"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            Warnings map[string]interface{}  `json:"warnings"`
            BusinessDetails BusinessDetails  `json:"business_details"`
            VerifiedOn string  `json:"verified_on"`
            BusinessType string  `json:"business_type"`
            Name string  `json:"name"`
            CreatedBy UserSerializer  `json:"created_by"`
            Mode string  `json:"mode"`
            ContactDetails ContactDetails  `json:"contact_details"`
            BusinessInfo string  `json:"business_info"`
            FranchiseEnabled bool  `json:"franchise_enabled"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            Documents []Document  `json:"documents"`
            Taxes []CompanyTaxesSerializer  `json:"taxes"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            CreatedOn string  `json:"created_on"`
            UID float64  `json:"uid"`
            Stage string  `json:"stage"`
            NotificationEmails []string  `json:"notification_emails"`
         
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

        
            State string  `json:"state"`
            Landmark string  `json:"landmark"`
            City string  `json:"city"`
            Longitude float64  `json:"longitude"`
            Latitude float64  `json:"latitude"`
            CountryCode string  `json:"country_code"`
            Pincode float64  `json:"pincode"`
            Country string  `json:"country"`
            AddressType string  `json:"address_type"`
            Address2 string  `json:"address2"`
            Address1 string  `json:"address1"`
         
    }
    
    // CompanyTaxesSerializer1 used by CompanyProfile
    type CompanyTaxesSerializer1 struct {

        
            Rate float64  `json:"rate"`
            Enable bool  `json:"enable"`
            EffectiveDate string  `json:"effective_date"`
         
    }
    
    // UpdateCompany used by CompanyProfile
    type UpdateCompany struct {

        
            BusinessType string  `json:"business_type"`
            CompanyType string  `json:"company_type"`
            Addresses []CreateUpdateAddressSerializer  `json:"addresses"`
            FranchiseEnabled bool  `json:"franchise_enabled"`
            Documents []Document  `json:"documents"`
            Name string  `json:"name"`
            Warnings map[string]interface{}  `json:"warnings"`
            Taxes []CompanyTaxesSerializer1  `json:"taxes"`
            BusinessDetails BusinessDetails  `json:"business_details"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            NotificationEmails []string  `json:"notification_emails"`
            RejectReason string  `json:"reject_reason"`
            Slug string  `json:"slug"`
            ContactDetails ContactDetails  `json:"contact_details"`
            BusinessInfo string  `json:"business_info"`
         
    }
    
    // ProfileSuccessResponse used by CompanyProfile
    type ProfileSuccessResponse struct {

        
            Success bool  `json:"success"`
            UID float64  `json:"uid"`
            Message string  `json:"message"`
         
    }
    
    // DocumentsObj used by CompanyProfile
    type DocumentsObj struct {

        
            Pending float64  `json:"pending"`
            Verified float64  `json:"verified"`
         
    }
    
    // MetricsSerializer used by CompanyProfile
    type MetricsSerializer struct {

        
            StoreDocuments DocumentsObj  `json:"store_documents"`
            CompanyDocuments DocumentsObj  `json:"company_documents"`
            Stage string  `json:"stage"`
            UID float64  `json:"uid"`
            Brand DocumentsObj  `json:"brand"`
            Store DocumentsObj  `json:"store"`
            Product DocumentsObj  `json:"product"`
         
    }
    
    // BrandBannerSerializer used by CompanyProfile
    type BrandBannerSerializer struct {

        
            Landscape string  `json:"landscape"`
            Portrait string  `json:"portrait"`
         
    }
    
    // GetBrandResponseSerializer used by CompanyProfile
    type GetBrandResponseSerializer struct {

        
            Synonyms []string  `json:"synonyms"`
            SlugKey string  `json:"slug_key"`
            ModifiedOn string  `json:"modified_on"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Banner BrandBannerSerializer  `json:"banner"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            Warnings map[string]interface{}  `json:"warnings"`
            VerifiedOn string  `json:"verified_on"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            Name string  `json:"name"`
            CreatedBy UserSerializer  `json:"created_by"`
            Mode string  `json:"mode"`
            RejectReason string  `json:"reject_reason"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            CreatedOn string  `json:"created_on"`
            UID float64  `json:"uid"`
            Stage string  `json:"stage"`
            Description string  `json:"description"`
            Logo string  `json:"logo"`
         
    }
    
    // CreateUpdateBrandRequestSerializer used by CompanyProfile
    type CreateUpdateBrandRequestSerializer struct {

        
            CompanyID float64  `json:"company_id"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            Synonyms []string  `json:"synonyms"`
            Name string  `json:"name"`
            Description string  `json:"description"`
            UID float64  `json:"uid"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            BrandTier string  `json:"brand_tier"`
            Logo string  `json:"logo"`
            Banner BrandBannerSerializer  `json:"banner"`
         
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

        
            BusinessType string  `json:"business_type"`
            CompanyType string  `json:"company_type"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            BusinessCountryInfo BusinessCountryInfo  `json:"business_country_info"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            Name string  `json:"name"`
            CreatedBy UserSerializer  `json:"created_by"`
            Details CompanyDetails  `json:"details"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            CreatedOn string  `json:"created_on"`
            UID float64  `json:"uid"`
            ModifiedOn string  `json:"modified_on"`
            MarketChannels []string  `json:"market_channels"`
            NotificationEmails []string  `json:"notification_emails"`
            RejectReason string  `json:"reject_reason"`
            Stage string  `json:"stage"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            VerifiedOn string  `json:"verified_on"`
         
    }
    
    // CompanyBrandSerializer used by CompanyProfile
    type CompanyBrandSerializer struct {

        
            ModifiedBy UserSerializer  `json:"modified_by"`
            Company CompanySerializer  `json:"company"`
            Warnings map[string]interface{}  `json:"warnings"`
            CreatedBy UserSerializer  `json:"created_by"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            CreatedOn string  `json:"created_on"`
            UID float64  `json:"uid"`
            ModifiedOn string  `json:"modified_on"`
            Brand GetBrandResponseSerializer  `json:"brand"`
            Stage string  `json:"stage"`
            RejectReason string  `json:"reject_reason"`
            VerifiedOn string  `json:"verified_on"`
         
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

        
            Company float64  `json:"company"`
            UID float64  `json:"uid"`
            Brands []float64  `json:"brands"`
         
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
    
    // AverageOrderProcessingTime used by CompanyProfile
    type AverageOrderProcessingTime struct {

        
            DurationType string  `json:"duration_type"`
            Duration float64  `json:"duration"`
         
    }
    
    // HolidayDateSerializer used by CompanyProfile
    type HolidayDateSerializer struct {

        
            EndDate string  `json:"end_date"`
            StartDate string  `json:"start_date"`
         
    }
    
    // HolidaySchemaSerializer used by CompanyProfile
    type HolidaySchemaSerializer struct {

        
            Title string  `json:"title"`
            HolidayType string  `json:"holiday_type"`
            Date HolidayDateSerializer  `json:"date"`
         
    }
    
    // ProductReturnConfigSerializer used by CompanyProfile
    type ProductReturnConfigSerializer struct {

        
            StoreUID float64  `json:"store_uid"`
            OnSameStore bool  `json:"on_same_store"`
         
    }
    
    // GetCompanySerializer used by CompanyProfile
    type GetCompanySerializer struct {

        
            BusinessType string  `json:"business_type"`
            CompanyType string  `json:"company_type"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            Name string  `json:"name"`
            CreatedBy UserSerializer  `json:"created_by"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            CreatedOn string  `json:"created_on"`
            UID float64  `json:"uid"`
            ModifiedOn string  `json:"modified_on"`
            Stage string  `json:"stage"`
            RejectReason string  `json:"reject_reason"`
            VerifiedOn string  `json:"verified_on"`
         
    }
    
    // LocationManagerSerializer used by CompanyProfile
    type LocationManagerSerializer struct {

        
            MobileNo SellerPhoneNumber  `json:"mobile_no"`
            Name string  `json:"name"`
            Email string  `json:"email"`
         
    }
    
    // InvoiceCredSerializer used by CompanyProfile
    type InvoiceCredSerializer struct {

        
            Enabled bool  `json:"enabled"`
            Username string  `json:"username"`
            Password string  `json:"password"`
         
    }
    
    // InvoiceDetailsSerializer used by CompanyProfile
    type InvoiceDetailsSerializer struct {

        
            EInvoice InvoiceCredSerializer  `json:"e_invoice"`
            EWaybill InvoiceCredSerializer  `json:"e_waybill"`
         
    }
    
    // GetLocationSerializer used by CompanyProfile
    type GetLocationSerializer struct {

        
            CreditNote bool  `json:"credit_note"`
            Code string  `json:"code"`
            Address GetAddressSerializer  `json:"address"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            Tags []string  `json:"tags"`
            ModifiedOn string  `json:"modified_on"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            AutoInvoice bool  `json:"auto_invoice"`
            BulkShipment bool  `json:"bulk_shipment"`
            Timing []LocationDayWiseSerializer  `json:"timing"`
            AvgOrderProcessingTime AverageOrderProcessingTime  `json:"avg_order_processing_time"`
            Holiday []HolidaySchemaSerializer  `json:"holiday"`
            DefaultOrderAcceptanceTiming bool  `json:"default_order_acceptance_timing"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            ProductReturnConfig ProductReturnConfigSerializer  `json:"product_return_config"`
            Warnings map[string]interface{}  `json:"warnings"`
            OrderAcceptanceTiming []LocationDayWiseSerializer  `json:"order_acceptance_timing"`
            VerifiedOn string  `json:"verified_on"`
            Company GetCompanySerializer  `json:"company"`
            Name string  `json:"name"`
            CreatedBy UserSerializer  `json:"created_by"`
            StoreType string  `json:"store_type"`
            PhoneNumber string  `json:"phone_number"`
            Documents []Document  `json:"documents"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            CreatedOn string  `json:"created_on"`
            UID float64  `json:"uid"`
            Stage string  `json:"stage"`
            Manager LocationManagerSerializer  `json:"manager"`
            NotificationEmails []string  `json:"notification_emails"`
            GstCredentials InvoiceDetailsSerializer  `json:"gst_credentials"`
            DisplayName string  `json:"display_name"`
         
    }
    
    // LocationListSerializer used by CompanyProfile
    type LocationListSerializer struct {

        
            Items []GetLocationSerializer  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // AddressSerializer used by CompanyProfile
    type AddressSerializer struct {

        
            State string  `json:"state"`
            Landmark string  `json:"landmark"`
            City string  `json:"city"`
            Longitude float64  `json:"longitude"`
            Latitude float64  `json:"latitude"`
            CountryCode string  `json:"country_code"`
            Pincode float64  `json:"pincode"`
            Country string  `json:"country"`
            AddressType string  `json:"address_type"`
            Address2 string  `json:"address2"`
            Address1 string  `json:"address1"`
         
    }
    
    // LocationSerializer used by CompanyProfile
    type LocationSerializer struct {

        
            CreditNote bool  `json:"credit_note"`
            Code string  `json:"code"`
            Address AddressSerializer  `json:"address"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            Tags []string  `json:"tags"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            AutoInvoice bool  `json:"auto_invoice"`
            BulkShipment bool  `json:"bulk_shipment"`
            Timing []LocationDayWiseSerializer  `json:"timing"`
            AvgOrderProcessingTime AverageOrderProcessingTime  `json:"avg_order_processing_time"`
            Holiday []HolidaySchemaSerializer  `json:"holiday"`
            DefaultOrderAcceptanceTiming bool  `json:"default_order_acceptance_timing"`
            ProductReturnConfig ProductReturnConfigSerializer  `json:"product_return_config"`
            Warnings map[string]interface{}  `json:"warnings"`
            OrderAcceptanceTiming []LocationDayWiseSerializer  `json:"order_acceptance_timing"`
            Company float64  `json:"company"`
            Name string  `json:"name"`
            StoreType string  `json:"store_type"`
            Documents []Document  `json:"documents"`
            Stage string  `json:"stage"`
            UID float64  `json:"uid"`
            GstCredentials InvoiceDetailsSerializer  `json:"gst_credentials"`
            Manager LocationManagerSerializer  `json:"manager"`
            NotificationEmails []string  `json:"notification_emails"`
            Slug string  `json:"slug"`
            DisplayName string  `json:"display_name"`
         
    }
    
    // BulkLocationSerializer used by CompanyProfile
    type BulkLocationSerializer struct {

        
            Data []LocationSerializer  `json:"data"`
         
    }
    
    // StoreTagsResponse used by CompanyProfile
    type StoreTagsResponse struct {

        
            Success bool  `json:"success"`
            Tags []string  `json:"tags"`
         
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
    
    // RuleDefinition used by Cart
    type RuleDefinition struct {

        
            CalculateOn string  `json:"calculate_on"`
            ValueType string  `json:"value_type"`
            Scope []string  `json:"scope"`
            Type string  `json:"type"`
            CurrencyCode string  `json:"currency_code"`
            AutoApply bool  `json:"auto_apply"`
            IsExact bool  `json:"is_exact"`
            ApplicableOn string  `json:"applicable_on"`
         
    }
    
    // CouponSchedule used by Cart
    type CouponSchedule struct {

        
            Start string  `json:"start"`
            End string  `json:"end"`
            NextSchedule []map[string]interface{}  `json:"next_schedule"`
            Cron string  `json:"cron"`
            Duration float64  `json:"duration"`
         
    }
    
    // CouponAction used by Cart
    type CouponAction struct {

        
            TxnMode string  `json:"txn_mode"`
            ActionDate string  `json:"action_date"`
         
    }
    
    // Validation used by Cart
    type Validation struct {

        
            Anonymous bool  `json:"anonymous"`
            UserRegisteredAfter string  `json:"user_registered_after"`
            AppID []string  `json:"app_id"`
         
    }
    
    // Identifier used by Cart
    type Identifier struct {

        
            EmailDomain []string  `json:"email_domain"`
            ExcludeBrandID []float64  `json:"exclude_brand_id"`
            UserID []string  `json:"user_id"`
            StoreID []float64  `json:"store_id"`
            ItemID []float64  `json:"item_id"`
            CollectionID []string  `json:"collection_id"`
            ArticleID []string  `json:"article_id"`
            BrandID []float64  `json:"brand_id"`
            CompanyID []float64  `json:"company_id"`
            CategoryID []float64  `json:"category_id"`
         
    }
    
    // CouponAuthor used by Cart
    type CouponAuthor struct {

        
            CreatedBy string  `json:"created_by"`
            ModifiedBy string  `json:"modified_by"`
         
    }
    
    // PriceRange used by Cart
    type PriceRange struct {

        
            Max float64  `json:"max"`
            Min float64  `json:"min"`
         
    }
    
    // UsesRemaining used by Cart
    type UsesRemaining struct {

        
            Total float64  `json:"total"`
            User float64  `json:"user"`
            App float64  `json:"app"`
         
    }
    
    // UsesRestriction used by Cart
    type UsesRestriction struct {

        
            Maximum UsesRemaining  `json:"maximum"`
            Remaining UsesRemaining  `json:"remaining"`
         
    }
    
    // PostOrder used by Cart
    type PostOrder struct {

        
            CancellationAllowed bool  `json:"cancellation_allowed"`
            ReturnAllowed bool  `json:"return_allowed"`
         
    }
    
    // BulkBundleRestriction used by Cart
    type BulkBundleRestriction struct {

        
            MultiStoreAllowed bool  `json:"multi_store_allowed"`
         
    }
    
    // PaymentAllowValue used by Cart
    type PaymentAllowValue struct {

        
            Max float64  `json:"max"`
         
    }
    
    // PaymentModes used by Cart
    type PaymentModes struct {

        
            Uses PaymentAllowValue  `json:"uses"`
            Networks []string  `json:"networks"`
            Codes []string  `json:"codes"`
            Iins []string  `json:"iins"`
            Types []string  `json:"types"`
         
    }
    
    // Restrictions used by Cart
    type Restrictions struct {

        
            UserGroups []float64  `json:"user_groups"`
            PriceRange PriceRange  `json:"price_range"`
            CouponAllowed bool  `json:"coupon_allowed"`
            Uses UsesRestriction  `json:"uses"`
            PostOrder PostOrder  `json:"post_order"`
            UserType string  `json:"user_type"`
            BulkBundle BulkBundleRestriction  `json:"bulk_bundle"`
            Payments map[string]PaymentModes  `json:"payments"`
            Platforms []string  `json:"platforms"`
            OrderingStores []float64  `json:"ordering_stores"`
         
    }
    
    // State used by Cart
    type State struct {

        
            IsPublic bool  `json:"is_public"`
            IsDisplay bool  `json:"is_display"`
            IsArchived bool  `json:"is_archived"`
         
    }
    
    // CouponDateMeta used by Cart
    type CouponDateMeta struct {

        
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // DisplayMetaDict used by Cart
    type DisplayMetaDict struct {

        
            Title string  `json:"title"`
            Subtitle string  `json:"subtitle"`
         
    }
    
    // DisplayMeta used by Cart
    type DisplayMeta struct {

        
            Title string  `json:"title"`
            Subtitle string  `json:"subtitle"`
            Remove DisplayMetaDict  `json:"remove"`
            Description string  `json:"description"`
            Auto DisplayMetaDict  `json:"auto"`
            Apply DisplayMetaDict  `json:"apply"`
         
    }
    
    // Ownership used by Cart
    type Ownership struct {

        
            PayableBy string  `json:"payable_by"`
            PayableCategory string  `json:"payable_category"`
         
    }
    
    // Rule used by Cart
    type Rule struct {

        
            Key float64  `json:"key"`
            Value float64  `json:"value"`
            Min float64  `json:"min"`
            Max float64  `json:"max"`
            DiscountQty float64  `json:"discount_qty"`
         
    }
    
    // Validity used by Cart
    type Validity struct {

        
            Priority float64  `json:"priority"`
         
    }
    
    // CouponAdd used by Cart
    type CouponAdd struct {

        
            Code string  `json:"code"`
            RuleDefinition RuleDefinition  `json:"rule_definition"`
            Schedule CouponSchedule  `json:"_schedule"`
            Action CouponAction  `json:"action"`
            Validation Validation  `json:"validation"`
            Identifiers Identifier  `json:"identifiers"`
            Author CouponAuthor  `json:"author"`
            Restrictions Restrictions  `json:"restrictions"`
            Tags []string  `json:"tags"`
            State State  `json:"state"`
            DateMeta CouponDateMeta  `json:"date_meta"`
            DisplayMeta DisplayMeta  `json:"display_meta"`
            Ownership Ownership  `json:"ownership"`
            Rule []Rule  `json:"rule"`
            Validity Validity  `json:"validity"`
            TypeSlug string  `json:"type_slug"`
         
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

        
            Code string  `json:"code"`
            RuleDefinition RuleDefinition  `json:"rule_definition"`
            Schedule CouponSchedule  `json:"_schedule"`
            Action CouponAction  `json:"action"`
            Validation Validation  `json:"validation"`
            Identifiers Identifier  `json:"identifiers"`
            Author CouponAuthor  `json:"author"`
            Restrictions Restrictions  `json:"restrictions"`
            Tags []string  `json:"tags"`
            State State  `json:"state"`
            DateMeta CouponDateMeta  `json:"date_meta"`
            DisplayMeta DisplayMeta  `json:"display_meta"`
            Ownership Ownership  `json:"ownership"`
            Rule []Rule  `json:"rule"`
            Validity Validity  `json:"validity"`
            TypeSlug string  `json:"type_slug"`
         
    }
    
    // CouponPartialUpdate used by Cart
    type CouponPartialUpdate struct {

        
            Archive bool  `json:"archive"`
            Schedule CouponSchedule  `json:"schedule"`
         
    }
    
    // Visibility used by Cart
    type Visibility struct {

        
            CouponList bool  `json:"coupon_list"`
            Pdp bool  `json:"pdp"`
         
    }
    
    // PromotionSchedule used by Cart
    type PromotionSchedule struct {

        
            Start string  `json:"start"`
            End string  `json:"end"`
            NextSchedule []map[string]interface{}  `json:"next_schedule"`
            Published bool  `json:"published"`
            Cron string  `json:"cron"`
            Duration float64  `json:"duration"`
         
    }
    
    // CompareObject used by Cart
    type CompareObject struct {

        
            Equals float64  `json:"equals"`
            GreaterThan float64  `json:"greater_than"`
            LessThanEquals float64  `json:"less_than_equals"`
            GreaterThanEquals float64  `json:"greater_than_equals"`
            LessThan float64  `json:"less_than"`
         
    }
    
    // ItemCriteria used by Cart
    type ItemCriteria struct {

        
            ItemExcludeStore []float64  `json:"item_exclude_store"`
            ItemExcludeL1Category []float64  `json:"item_exclude_l1_category"`
            AllItems bool  `json:"all_items"`
            BuyRules []string  `json:"buy_rules"`
            ItemCategory []float64  `json:"item_category"`
            ItemSize []string  `json:"item_size"`
            AvailableZones []string  `json:"available_zones"`
            ItemCompany []float64  `json:"item_company"`
            ItemDepartment []float64  `json:"item_department"`
            ItemExcludeCompany []float64  `json:"item_exclude_company"`
            ItemL2Category []float64  `json:"item_l2_category"`
            CartUniqueItemQuantity CompareObject  `json:"cart_unique_item_quantity"`
            ItemStore []float64  `json:"item_store"`
            ItemExcludeCategory []float64  `json:"item_exclude_category"`
            ItemSku []string  `json:"item_sku"`
            ItemExcludeSku []string  `json:"item_exclude_sku"`
            ItemExcludeID []float64  `json:"item_exclude_id"`
            CartQuantity CompareObject  `json:"cart_quantity"`
            CartTotal CompareObject  `json:"cart_total"`
            ItemExcludeBrand []float64  `json:"item_exclude_brand"`
            ItemID []float64  `json:"item_id"`
            ProductTags []string  `json:"product_tags"`
            ItemExcludeL2Category []float64  `json:"item_exclude_l2_category"`
            CartUniqueItemAmount CompareObject  `json:"cart_unique_item_amount"`
            ItemL1Category []float64  `json:"item_l1_category"`
            ItemBrand []float64  `json:"item_brand"`
            ItemTags []string  `json:"item_tags"`
            ItemExcludeDepartment []float64  `json:"item_exclude_department"`
         
    }
    
    // UsesRemaining1 used by Cart
    type UsesRemaining1 struct {

        
            User float64  `json:"user"`
            Total float64  `json:"total"`
         
    }
    
    // UsesRestriction1 used by Cart
    type UsesRestriction1 struct {

        
            Maximum UsesRemaining1  `json:"maximum"`
            Remaining UsesRemaining1  `json:"remaining"`
         
    }
    
    // PostOrder1 used by Cart
    type PostOrder1 struct {

        
            CancellationAllowed bool  `json:"cancellation_allowed"`
            ReturnAllowed bool  `json:"return_allowed"`
         
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

        
            Type string  `json:"type"`
            Uses PaymentAllowValue1  `json:"uses"`
            Codes []string  `json:"codes"`
         
    }
    
    // Restrictions1 used by Cart
    type Restrictions1 struct {

        
            UserGroups []float64  `json:"user_groups"`
            UserID []string  `json:"user_id"`
            Uses UsesRestriction1  `json:"uses"`
            PostOrder PostOrder1  `json:"post_order"`
            UserRegistered UserRegistered  `json:"user_registered"`
            Payments []PromotionPaymentModes  `json:"payments"`
            Platforms []string  `json:"platforms"`
            OrderingStores []float64  `json:"ordering_stores"`
            OrderQuantity float64  `json:"order_quantity"`
            AnonymousUsers bool  `json:"anonymous_users"`
         
    }
    
    // Ownership1 used by Cart
    type Ownership1 struct {

        
            PayableBy string  `json:"payable_by"`
            PayableCategory string  `json:"payable_category"`
         
    }
    
    // PromotionAction used by Cart
    type PromotionAction struct {

        
            ActionType string  `json:"action_type"`
            ActionDate string  `json:"action_date"`
         
    }
    
    // DisplayMeta1 used by Cart
    type DisplayMeta1 struct {

        
            Name string  `json:"name"`
            OfferLabel string  `json:"offer_label"`
            OfferText string  `json:"offer_text"`
            Description string  `json:"description"`
         
    }
    
    // PromotionAuthor used by Cart
    type PromotionAuthor struct {

        
            CreatedBy string  `json:"created_by"`
            ModifiedBy string  `json:"modified_by"`
         
    }
    
    // PromotionDateMeta used by Cart
    type PromotionDateMeta struct {

        
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // DiscountOffer used by Cart
    type DiscountOffer struct {

        
            MaxUsagePerTransaction float64  `json:"max_usage_per_transaction"`
            Code string  `json:"code"`
            DiscountAmount float64  `json:"discount_amount"`
            MaxOfferQuantity float64  `json:"max_offer_quantity"`
            ApportionDiscount bool  `json:"apportion_discount"`
            MaxDiscountAmount float64  `json:"max_discount_amount"`
            DiscountPrice float64  `json:"discount_price"`
            PartialCanRet bool  `json:"partial_can_ret"`
            DiscountPercentage float64  `json:"discount_percentage"`
            MinOfferQuantity float64  `json:"min_offer_quantity"`
         
    }
    
    // DiscountRule used by Cart
    type DiscountRule struct {

        
            BuyCondition string  `json:"buy_condition"`
            DiscountType string  `json:"discount_type"`
            ItemCriteria ItemCriteria  `json:"item_criteria"`
            Offer DiscountOffer  `json:"offer"`
         
    }
    
    // PromotionListItem used by Cart
    type PromotionListItem struct {

        
            PromoGroup string  `json:"promo_group"`
            Visiblility Visibility  `json:"visiblility"`
            Schedule PromotionSchedule  `json:"_schedule"`
            BuyRules map[string]ItemCriteria  `json:"buy_rules"`
            Restrictions Restrictions1  `json:"restrictions"`
            ApplyExclusive string  `json:"apply_exclusive"`
            Ownership Ownership1  `json:"ownership"`
            PostOrderAction PromotionAction  `json:"post_order_action"`
            Currency string  `json:"currency"`
            DisplayMeta DisplayMeta1  `json:"display_meta"`
            ApplyPriority float64  `json:"apply_priority"`
            Author PromotionAuthor  `json:"author"`
            ApplicationID string  `json:"application_id"`
            Mode string  `json:"mode"`
            DateMeta PromotionDateMeta  `json:"date_meta"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Code string  `json:"code"`
            Stackable bool  `json:"stackable"`
            CalculateOn string  `json:"calculate_on"`
            DiscountRules []DiscountRule  `json:"discount_rules"`
            ApplyAllDiscount bool  `json:"apply_all_discount"`
            PromotionType string  `json:"promotion_type"`
         
    }
    
    // PromotionsResponse used by Cart
    type PromotionsResponse struct {

        
            Page Page  `json:"page"`
            Items PromotionListItem  `json:"items"`
         
    }
    
    // PromotionAdd used by Cart
    type PromotionAdd struct {

        
            PromoGroup string  `json:"promo_group"`
            Visiblility Visibility  `json:"visiblility"`
            Schedule PromotionSchedule  `json:"_schedule"`
            BuyRules map[string]ItemCriteria  `json:"buy_rules"`
            Restrictions Restrictions1  `json:"restrictions"`
            ApplyExclusive string  `json:"apply_exclusive"`
            Ownership Ownership1  `json:"ownership"`
            PostOrderAction PromotionAction  `json:"post_order_action"`
            Currency string  `json:"currency"`
            DisplayMeta DisplayMeta1  `json:"display_meta"`
            ApplyPriority float64  `json:"apply_priority"`
            Author PromotionAuthor  `json:"author"`
            ApplicationID string  `json:"application_id"`
            Mode string  `json:"mode"`
            DateMeta PromotionDateMeta  `json:"date_meta"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Code string  `json:"code"`
            Stackable bool  `json:"stackable"`
            CalculateOn string  `json:"calculate_on"`
            DiscountRules []DiscountRule  `json:"discount_rules"`
            ApplyAllDiscount bool  `json:"apply_all_discount"`
            PromotionType string  `json:"promotion_type"`
         
    }
    
    // PromotionUpdate used by Cart
    type PromotionUpdate struct {

        
            PromoGroup string  `json:"promo_group"`
            Visiblility Visibility  `json:"visiblility"`
            Schedule PromotionSchedule  `json:"_schedule"`
            BuyRules map[string]ItemCriteria  `json:"buy_rules"`
            Restrictions Restrictions1  `json:"restrictions"`
            ApplyExclusive string  `json:"apply_exclusive"`
            Ownership Ownership1  `json:"ownership"`
            PostOrderAction PromotionAction  `json:"post_order_action"`
            Currency string  `json:"currency"`
            DisplayMeta DisplayMeta1  `json:"display_meta"`
            ApplyPriority float64  `json:"apply_priority"`
            Author PromotionAuthor  `json:"author"`
            ApplicationID string  `json:"application_id"`
            Mode string  `json:"mode"`
            DateMeta PromotionDateMeta  `json:"date_meta"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Code string  `json:"code"`
            Stackable bool  `json:"stackable"`
            CalculateOn string  `json:"calculate_on"`
            DiscountRules []DiscountRule  `json:"discount_rules"`
            ApplyAllDiscount bool  `json:"apply_all_discount"`
            PromotionType string  `json:"promotion_type"`
         
    }
    
    // PromotionPartialUpdate used by Cart
    type PromotionPartialUpdate struct {

        
            Archive bool  `json:"archive"`
            Schedule PromotionSchedule  `json:"schedule"`
         
    }
    
    // ActivePromosResponse used by Cart
    type ActivePromosResponse struct {

        
            IsHidden bool  `json:"is_hidden"`
            Example string  `json:"example"`
            EntitySlug string  `json:"entity_slug"`
            Type string  `json:"type"`
            Title string  `json:"title"`
            Subtitle string  `json:"subtitle"`
            Description string  `json:"description"`
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
            EntityType string  `json:"entity_type"`
         
    }
    
    // Charges used by Cart
    type Charges struct {

        
            Threshold float64  `json:"threshold"`
            Charges float64  `json:"charges"`
         
    }
    
    // DeliveryCharges used by Cart
    type DeliveryCharges struct {

        
            Charges []Charges  `json:"charges"`
            Enabled bool  `json:"enabled"`
         
    }
    
    // CartMetaConfigUpdate used by Cart
    type CartMetaConfigUpdate struct {

        
            GiftDisplayText string  `json:"gift_display_text"`
            BulkCoupons bool  `json:"bulk_coupons"`
            Enabled bool  `json:"enabled"`
            MaxCartItems float64  `json:"max_cart_items"`
            GiftPricing float64  `json:"gift_pricing"`
            RevenueEngineCoupon bool  `json:"revenue_engine_coupon"`
            DeliveryCharges DeliveryCharges  `json:"delivery_charges"`
            MinCartValue float64  `json:"min_cart_value"`
         
    }
    
    // CartMetaConfigAdd used by Cart
    type CartMetaConfigAdd struct {

        
            GiftDisplayText string  `json:"gift_display_text"`
            BulkCoupons bool  `json:"bulk_coupons"`
            Enabled bool  `json:"enabled"`
            MaxCartItems float64  `json:"max_cart_items"`
            GiftPricing float64  `json:"gift_pricing"`
            RevenueEngineCoupon bool  `json:"revenue_engine_coupon"`
            DeliveryCharges DeliveryCharges  `json:"delivery_charges"`
            MinCartValue float64  `json:"min_cart_value"`
         
    }
    
    // Article used by Cart
    type Article struct {

        
            Code string  `json:"code"`
            Meta map[string]interface{}  `json:"meta"`
            Value float64  `json:"value"`
            Type string  `json:"type"`
            ArticleID string  `json:"article_id"`
         
    }
    
    // Collecttion used by Cart
    type Collecttion struct {

        
            RefundBy string  `json:"refund_by"`
            CollectedBy string  `json:"collected_by"`
         
    }
    
    // CartDynamicInjectionUpdate used by Cart
    type CartDynamicInjectionUpdate struct {

        
            CartValue float64  `json:"cart_value"`
            ArticleIds []Article  `json:"article_ids"`
            Meta map[string]interface{}  `json:"meta"`
            UserID string  `json:"user_id"`
            Value float64  `json:"value"`
            Type string  `json:"type"`
            ApplyExpiry string  `json:"apply_expiry"`
            IsAuthenticated bool  `json:"is_authenticated"`
            ArticleLevelDistribution bool  `json:"article_level_distribution"`
            CartID string  `json:"cart_id"`
            Message string  `json:"message"`
            Collection Collecttion  `json:"collection"`
            AllowedRefund bool  `json:"allowed_refund"`
         
    }
    
    // CartDynamicInjection used by Cart
    type CartDynamicInjection struct {

        
            CartValue float64  `json:"cart_value"`
            ArticleIds []Article  `json:"article_ids"`
            Meta map[string]interface{}  `json:"meta"`
            UserID string  `json:"user_id"`
            Value float64  `json:"value"`
            Type string  `json:"type"`
            ApplyExpiry string  `json:"apply_expiry"`
            IsAuthenticated bool  `json:"is_authenticated"`
            ArticleLevelDistribution bool  `json:"article_level_distribution"`
            CartID string  `json:"cart_id"`
            Message string  `json:"message"`
            Collection Collecttion  `json:"collection"`
            AllowedRefund bool  `json:"allowed_refund"`
            InjectionID string  `json:"injection_id"`
         
    }
    
    // CartDynamicInjectionResponse used by Cart
    type CartDynamicInjectionResponse struct {

        
            Data CartDynamicInjection  `json:"data"`
         
    }
    
    // CartDynamicInjectionAdd used by Cart
    type CartDynamicInjectionAdd struct {

        
            CartValue float64  `json:"cart_value"`
            ArticleIds []Article  `json:"article_ids"`
            Meta map[string]interface{}  `json:"meta"`
            UserID string  `json:"user_id"`
            Value float64  `json:"value"`
            Type string  `json:"type"`
            ApplyExpiry string  `json:"apply_expiry"`
            IsAuthenticated bool  `json:"is_authenticated"`
            ArticleLevelDistribution bool  `json:"article_level_distribution"`
            CartID string  `json:"cart_id"`
            Message string  `json:"message"`
            Collection Collecttion  `json:"collection"`
            AllowedRefund bool  `json:"allowed_refund"`
         
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
    
    // DisplayBreakup used by Cart
    type DisplayBreakup struct {

        
            Key string  `json:"key"`
            CurrencySymbol string  `json:"currency_symbol"`
            Value float64  `json:"value"`
            Message []string  `json:"message"`
            Display string  `json:"display"`
            CurrencyCode string  `json:"currency_code"`
         
    }
    
    // RawBreakup used by Cart
    type RawBreakup struct {

        
            GstCharges float64  `json:"gst_charges"`
            DeliveryCharge float64  `json:"delivery_charge"`
            MrpTotal float64  `json:"mrp_total"`
            FyndCash float64  `json:"fynd_cash"`
            Coupon float64  `json:"coupon"`
            Vog float64  `json:"vog"`
            Total float64  `json:"total"`
            GiftCard float64  `json:"gift_card"`
            ConvenienceFee float64  `json:"convenience_fee"`
            Discount float64  `json:"discount"`
            Subtotal float64  `json:"subtotal"`
            YouSaved float64  `json:"you_saved"`
            CodCharge float64  `json:"cod_charge"`
         
    }
    
    // LoyaltyPoints used by Cart
    type LoyaltyPoints struct {

        
            Applicable float64  `json:"applicable"`
            IsApplied bool  `json:"is_applied"`
            Total float64  `json:"total"`
            Description string  `json:"description"`
         
    }
    
    // CouponBreakup used by Cart
    type CouponBreakup struct {

        
            Code string  `json:"code"`
            IsApplied bool  `json:"is_applied"`
            SubTitle string  `json:"sub_title"`
            MaxDiscountValue float64  `json:"max_discount_value"`
            Value float64  `json:"value"`
            Type string  `json:"type"`
            Title string  `json:"title"`
            Description string  `json:"description"`
            Message string  `json:"message"`
            CouponValue float64  `json:"coupon_value"`
            MinimumCartValue float64  `json:"minimum_cart_value"`
            CouponType string  `json:"coupon_type"`
            UID string  `json:"uid"`
         
    }
    
    // CartBreakup used by Cart
    type CartBreakup struct {

        
            Display []DisplayBreakup  `json:"display"`
            Raw RawBreakup  `json:"raw"`
            LoyaltyPoints LoyaltyPoints  `json:"loyalty_points"`
            Coupon CouponBreakup  `json:"coupon"`
         
    }
    
    // FreeGiftItem used by Cart
    type FreeGiftItem struct {

        
            ItemID float64  `json:"item_id"`
            ItemName string  `json:"item_name"`
            ItemBrandName string  `json:"item_brand_name"`
            ItemPriceDetails map[string]interface{}  `json:"item_price_details"`
            ItemImagesURL []string  `json:"item_images_url"`
            ItemSlug string  `json:"item_slug"`
         
    }
    
    // AppliedFreeArticles used by Cart
    type AppliedFreeArticles struct {

        
            FreeGiftItemDetails FreeGiftItem  `json:"free_gift_item_details"`
            ParentItemIdentifier string  `json:"parent_item_identifier"`
            ArticleID string  `json:"article_id"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // BuyRules used by Cart
    type BuyRules struct {

        
            CartConditions map[string]interface{}  `json:"cart_conditions"`
            ItemCriteria map[string]interface{}  `json:"item_criteria"`
         
    }
    
    // DiscountRulesApp used by Cart
    type DiscountRulesApp struct {

        
            RawOffer map[string]interface{}  `json:"raw_offer"`
            Offer map[string]interface{}  `json:"offer"`
            MatchedBuyRules []string  `json:"matched_buy_rules"`
            ItemCriteria map[string]interface{}  `json:"item_criteria"`
         
    }
    
    // Ownership2 used by Cart
    type Ownership2 struct {

        
            PayableBy string  `json:"payable_by"`
            PayableCategory string  `json:"payable_category"`
         
    }
    
    // AppliedPromotion used by Cart
    type AppliedPromotion struct {

        
            MrpPromotion bool  `json:"mrp_promotion"`
            ArticleQuantity float64  `json:"article_quantity"`
            AppliedFreeArticles []AppliedFreeArticles  `json:"applied_free_articles"`
            PromoID string  `json:"promo_id"`
            PromotionGroup string  `json:"promotion_group"`
            PromotionOfferText string  `json:"promotion_offer_text"`
            BuyRules []BuyRules  `json:"buy_rules"`
            Amount float64  `json:"amount"`
            DiscountRules []DiscountRulesApp  `json:"discount_rules"`
            PromotionName string  `json:"promotion_name"`
            Ownership Ownership2  `json:"ownership"`
            PromotionType string  `json:"promotion_type"`
         
    }
    
    // CouponDetails used by Cart
    type CouponDetails struct {

        
            Code string  `json:"code"`
            DiscountTotalQuantity float64  `json:"discount_total_quantity"`
            DiscountSingleQuantity float64  `json:"discount_single_quantity"`
         
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
    
    // ProductAvailabilitySize used by Cart
    type ProductAvailabilitySize struct {

        
            Meta map[string]interface{}  `json:"meta"`
            Value string  `json:"value"`
            IsAvailable bool  `json:"is_available"`
            Display string  `json:"display"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // ProductAvailability used by Cart
    type ProductAvailability struct {

        
            IsValid bool  `json:"is_valid"`
            OtherStoreQuantity float64  `json:"other_store_quantity"`
            AvailableSizes []ProductAvailabilitySize  `json:"available_sizes"`
            Deliverable bool  `json:"deliverable"`
            Sizes []string  `json:"sizes"`
            OutOfStock bool  `json:"out_of_stock"`
         
    }
    
    // PromoMeta used by Cart
    type PromoMeta struct {

        
            Message string  `json:"message"`
         
    }
    
    // ProductPrice used by Cart
    type ProductPrice struct {

        
            Effective float64  `json:"effective"`
            CurrencySymbol string  `json:"currency_symbol"`
            Selling float64  `json:"selling"`
            Marked float64  `json:"marked"`
            CurrencyCode string  `json:"currency_code"`
            AddOn float64  `json:"add_on"`
         
    }
    
    // ProductPriceInfo used by Cart
    type ProductPriceInfo struct {

        
            Base ProductPrice  `json:"base"`
            Converted ProductPrice  `json:"converted"`
         
    }
    
    // ProductImage used by Cart
    type ProductImage struct {

        
            SecureURL string  `json:"secure_url"`
            URL string  `json:"url"`
            AspectRatio string  `json:"aspect_ratio"`
         
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

        
            Name string  `json:"name"`
            UID float64  `json:"uid"`
         
    }
    
    // Tags used by Cart
    type Tags struct {

        
            Tags map[string]interface{}  `json:"tags"`
         
    }
    
    // CategoryInfo used by Cart
    type CategoryInfo struct {

        
            Name string  `json:"name"`
            UID float64  `json:"uid"`
         
    }
    
    // NetQuantity used by Cart
    type NetQuantity struct {

        
            Value string  `json:"value"`
            Unit string  `json:"unit"`
         
    }
    
    // CartProduct used by Cart
    type CartProduct struct {

        
            Images []ProductImage  `json:"images"`
            Action ProductAction  `json:"action"`
            Type string  `json:"type"`
            Slug string  `json:"slug"`
            Brand BaseInfo  `json:"brand"`
            ItemCode string  `json:"item_code"`
            Name string  `json:"name"`
            Tags []string  `json:"tags"`
            TeaserTag Tags  `json:"teaser_tag"`
            Categories []CategoryInfo  `json:"categories"`
            NetQuantity NetQuantity  `json:"net_quantity"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            UID float64  `json:"uid"`
         
    }
    
    // BasePrice used by Cart
    type BasePrice struct {

        
            CurrencyCode string  `json:"currency_code"`
            Effective float64  `json:"effective"`
            CurrencySymbol string  `json:"currency_symbol"`
            Marked float64  `json:"marked"`
         
    }
    
    // ArticlePriceInfo used by Cart
    type ArticlePriceInfo struct {

        
            Base BasePrice  `json:"base"`
            Converted BasePrice  `json:"converted"`
         
    }
    
    // StoreInfo used by Cart
    type StoreInfo struct {

        
            Name string  `json:"name"`
            UID float64  `json:"uid"`
            StoreCode string  `json:"store_code"`
         
    }
    
    // ProductArticle used by Cart
    type ProductArticle struct {

        
            Size string  `json:"size"`
            ProductGroupTags []string  `json:"product_group_tags"`
            Price ArticlePriceInfo  `json:"price"`
            SellerIdentifier string  `json:"seller_identifier"`
            Type string  `json:"type"`
            Seller BaseInfo  `json:"seller"`
            IsGiftVisible bool  `json:"is_gift_visible"`
            Quantity float64  `json:"quantity"`
            GiftCard map[string]interface{}  `json:"gift_card"`
            CartItemMeta map[string]interface{}  `json:"cart_item_meta"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            Store StoreInfo  `json:"store"`
            MtoQuantity float64  `json:"mto_quantity"`
            Identifier map[string]interface{}  `json:"identifier"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            UID string  `json:"uid"`
         
    }
    
    // CartProductIdentifer used by Cart
    type CartProductIdentifer struct {

        
            Identifier string  `json:"identifier"`
         
    }
    
    // CartProductInfo used by Cart
    type CartProductInfo struct {

        
            PromotionsApplied []AppliedPromotion  `json:"promotions_applied"`
            Key string  `json:"key"`
            Coupon CouponDetails  `json:"coupon"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            Availability ProductAvailability  `json:"availability"`
            Moq map[string]interface{}  `json:"moq"`
            PromoMeta PromoMeta  `json:"promo_meta"`
            PricePerUnit ProductPriceInfo  `json:"price_per_unit"`
            Product CartProduct  `json:"product"`
            CouponMessage string  `json:"coupon_message"`
            Quantity float64  `json:"quantity"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            Article ProductArticle  `json:"article"`
            IsSet bool  `json:"is_set"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
            Message string  `json:"message"`
            Price ProductPriceInfo  `json:"price"`
            BulkOffer map[string]interface{}  `json:"bulk_offer"`
            Discount string  `json:"discount"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
         
    }
    
    // OpenapiCartDetailsResponse used by Cart
    type OpenapiCartDetailsResponse struct {

        
            BreakupValues CartBreakup  `json:"breakup_values"`
            Message string  `json:"message"`
            IsValid bool  `json:"is_valid"`
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

        
            Meta map[string]interface{}  `json:"meta"`
            Pincode float64  `json:"pincode"`
            CountryIsoCode string  `json:"country_iso_code"`
            Address string  `json:"address"`
            Landmark string  `json:"landmark"`
            AreaCode string  `json:"area_code"`
            Area string  `json:"area"`
            Name string  `json:"name"`
            Phone float64  `json:"phone"`
            CountryPhoneCode string  `json:"country_phone_code"`
            AddressType string  `json:"address_type"`
            Email string  `json:"email"`
            State string  `json:"state"`
            Country string  `json:"country"`
            City string  `json:"city"`
            AreaCodeSlug string  `json:"area_code_slug"`
            CountryCode string  `json:"country_code"`
         
    }
    
    // OpenApiCartServiceabilityRequest used by Cart
    type OpenApiCartServiceabilityRequest struct {

        
            CartItems []CartItem  `json:"cart_items"`
            ShippingAddress ShippingAddress  `json:"shipping_address"`
         
    }
    
    // OpenApiCartServiceabilityResponse used by Cart
    type OpenApiCartServiceabilityResponse struct {

        
            BreakupValues CartBreakup  `json:"breakup_values"`
            IsValid bool  `json:"is_valid"`
            Items []CartProductInfo  `json:"items"`
            Message string  `json:"message"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
         
    }
    
    // OpenApiFiles used by Cart
    type OpenApiFiles struct {

        
            Values []string  `json:"values"`
            Key string  `json:"key"`
         
    }
    
    // MultiTenderPaymentMeta used by Cart
    type MultiTenderPaymentMeta struct {

        
            OrderID string  `json:"order_id"`
            CurrentStatus string  `json:"current_status"`
            PaymentGateway string  `json:"payment_gateway"`
            PaymentID string  `json:"payment_id"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
         
    }
    
    // MultiTenderPaymentMethod used by Cart
    type MultiTenderPaymentMethod struct {

        
            Name string  `json:"name"`
            Amount float64  `json:"amount"`
            Meta MultiTenderPaymentMeta  `json:"meta"`
            Mode string  `json:"mode"`
         
    }
    
    // CartItemMeta used by Cart
    type CartItemMeta struct {

        
            GroupID string  `json:"group_id"`
            PrimaryItem bool  `json:"primary_item"`
         
    }
    
    // OpenApiOrderItem used by Cart
    type OpenApiOrderItem struct {

        
            Files []OpenApiFiles  `json:"files"`
            Meta CartItemMeta  `json:"meta"`
            PaymentMethods []MultiTenderPaymentMethod  `json:"payment_methods"`
            ProductID float64  `json:"product_id"`
            AmountPaid float64  `json:"amount_paid"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            CashbackApplied float64  `json:"cashback_applied"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            Discount float64  `json:"discount"`
            PriceMarked float64  `json:"price_marked"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            PriceEffective float64  `json:"price_effective"`
            Quantity float64  `json:"quantity"`
            EmployeeDiscount float64  `json:"employee_discount"`
            CodCharges float64  `json:"cod_charges"`
            Size string  `json:"size"`
            DeliveryCharges float64  `json:"delivery_charges"`
         
    }
    
    // OpenApiPlatformCheckoutReq used by Cart
    type OpenApiPlatformCheckoutReq struct {

        
            CartValue float64  `json:"cart_value"`
            Files []OpenApiFiles  `json:"files"`
            PaymentMethods []MultiTenderPaymentMethod  `json:"payment_methods"`
            Coupon string  `json:"coupon"`
            PaymentMode string  `json:"payment_mode"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            CashbackApplied float64  `json:"cashback_applied"`
            Gstin string  `json:"gstin"`
            CouponValue float64  `json:"coupon_value"`
            CurrencyCode string  `json:"currency_code"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            CouponCode string  `json:"coupon_code"`
            CartItems []OpenApiOrderItem  `json:"cart_items"`
            OrderID string  `json:"order_id"`
            BillingAddress ShippingAddress  `json:"billing_address"`
            Comment string  `json:"comment"`
            ShippingAddress ShippingAddress  `json:"shipping_address"`
            EmployeeDiscount map[string]interface{}  `json:"employee_discount"`
            CodCharges float64  `json:"cod_charges"`
            DeliveryCharges float64  `json:"delivery_charges"`
         
    }
    
    // OpenApiCheckoutResponse used by Cart
    type OpenApiCheckoutResponse struct {

        
            OrderRefID string  `json:"order_ref_id"`
            Message string  `json:"message"`
            OrderID string  `json:"order_id"`
            Success bool  `json:"success"`
         
    }
    
    // AbandonedCart used by Cart
    type AbandonedCart struct {

        
            CartValue float64  `json:"cart_value"`
            FcIndexMap []float64  `json:"fc_index_map"`
            Articles []map[string]interface{}  `json:"articles"`
            FyndCredits map[string]interface{}  `json:"fynd_credits"`
            Coupon map[string]interface{}  `json:"coupon"`
            PaymentMode string  `json:"payment_mode"`
            LastModified string  `json:"last_modified"`
            Meta map[string]interface{}  `json:"meta"`
            Payments map[string]interface{}  `json:"payments"`
            ID string  `json:"_id"`
            AppID string  `json:"app_id"`
            IsArchive bool  `json:"is_archive"`
            IsActive bool  `json:"is_active"`
            CheckoutMode string  `json:"checkout_mode"`
            ExpireAt string  `json:"expire_at"`
            IsDefault bool  `json:"is_default"`
            Cashback map[string]interface{}  `json:"cashback"`
            PaymentMethods []map[string]interface{}  `json:"payment_methods"`
            UserID string  `json:"user_id"`
            Gstin string  `json:"gstin"`
            CreatedOn string  `json:"created_on"`
            Shipments []map[string]interface{}  `json:"shipments"`
            UID float64  `json:"uid"`
            OrderID string  `json:"order_id"`
            Comment string  `json:"comment"`
            MergeQty bool  `json:"merge_qty"`
            PickUpCustomerDetails map[string]interface{}  `json:"pick_up_customer_details"`
            BulkCouponDiscount float64  `json:"bulk_coupon_discount"`
            Discount float64  `json:"discount"`
            Promotion map[string]interface{}  `json:"promotion"`
            BuyNow bool  `json:"buy_now"`
            CodCharges map[string]interface{}  `json:"cod_charges"`
            DeliveryCharges map[string]interface{}  `json:"delivery_charges"`
         
    }
    
    // AbandonedCartResponse used by Cart
    type AbandonedCartResponse struct {

        
            Page Page  `json:"page"`
            Success bool  `json:"success"`
            Items []AbandonedCart  `json:"items"`
            Message string  `json:"message"`
            Result map[string]interface{}  `json:"result"`
         
    }
    
    // PaymentSelectionLock used by Cart
    type PaymentSelectionLock struct {

        
            DefaultOptions string  `json:"default_options"`
            PaymentIdentifier string  `json:"payment_identifier"`
            Enabled bool  `json:"enabled"`
         
    }
    
    // CartCurrency used by Cart
    type CartCurrency struct {

        
            Code string  `json:"code"`
            Symbol string  `json:"symbol"`
         
    }
    
    // CartDetailResponse used by Cart
    type CartDetailResponse struct {

        
            LastModified string  `json:"last_modified"`
            AppliedPromoDetails []AppliedPromotion  `json:"applied_promo_details"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            IsValid bool  `json:"is_valid"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            Currency CartCurrency  `json:"currency"`
            Items []CartProductInfo  `json:"items"`
            CheckoutMode string  `json:"checkout_mode"`
            PanConfig map[string]interface{}  `json:"pan_config"`
            CouponText string  `json:"coupon_text"`
            Gstin string  `json:"gstin"`
            Message string  `json:"message"`
            ID string  `json:"id"`
            PanNo string  `json:"pan_no"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            Comment string  `json:"comment"`
            BuyNow bool  `json:"buy_now"`
         
    }
    
    // AddProductCart used by Cart
    type AddProductCart struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ProductGroupTags []string  `json:"product_group_tags"`
            Meta map[string]interface{}  `json:"meta"`
            StoreID float64  `json:"store_id"`
            ItemID float64  `json:"item_id"`
            ArticleAssignment map[string]interface{}  `json:"article_assignment"`
            Quantity float64  `json:"quantity"`
            Pos bool  `json:"pos"`
            SellerID float64  `json:"seller_id"`
            Display string  `json:"display"`
            ArticleID string  `json:"article_id"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            ParentItemIdentifiers []map[string]string  `json:"parent_item_identifiers"`
            ItemSize string  `json:"item_size"`
         
    }
    
    // AddCartRequest used by Cart
    type AddCartRequest struct {

        
            NewCart bool  `json:"new_cart"`
            Items []AddProductCart  `json:"items"`
         
    }
    
    // AddCartDetailResponse used by Cart
    type AddCartDetailResponse struct {

        
            Message string  `json:"message"`
            Partial bool  `json:"partial"`
            Success bool  `json:"success"`
            Cart CartDetailResponse  `json:"cart"`
         
    }
    
    // UpdateProductCart used by Cart
    type UpdateProductCart struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Meta map[string]interface{}  `json:"meta"`
            ItemID float64  `json:"item_id"`
            ItemIndex float64  `json:"item_index"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
            Quantity float64  `json:"quantity"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            ArticleID string  `json:"article_id"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            ItemSize string  `json:"item_size"`
         
    }
    
    // UpdateCartRequest used by Cart
    type UpdateCartRequest struct {

        
            Operation string  `json:"operation"`
            Items []UpdateProductCart  `json:"items"`
         
    }
    
    // UpdateCartDetailResponse used by Cart
    type UpdateCartDetailResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
            Cart CartDetailResponse  `json:"cart"`
         
    }
    
    // OverrideCartItemPromo used by Cart
    type OverrideCartItemPromo struct {

        
            ItemList []map[string]interface{}  `json:"item_list"`
            PromoID string  `json:"promo_id"`
            PromoAmount string  `json:"promo_amount"`
            RwrdTndr string  `json:"rwrd_tndr"`
            PromoDesc string  `json:"promo_desc"`
         
    }
    
    // OverrideCartItem used by Cart
    type OverrideCartItem struct {

        
            SellerIdentifier string  `json:"seller_identifier"`
            ItemID float64  `json:"item_id"`
            AmountPaid float64  `json:"amount_paid"`
            Discount float64  `json:"discount"`
            PriceMarked float64  `json:"price_marked"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            PriceEffective float64  `json:"price_effective"`
            Quantity float64  `json:"quantity"`
            PromoList []OverrideCartItemPromo  `json:"promo_list"`
            Size string  `json:"size"`
         
    }
    
    // OverrideCheckoutReq used by Cart
    type OverrideCheckoutReq struct {

        
            CartItems []OverrideCartItem  `json:"cart_items"`
            MerchantCode string  `json:"merchant_code"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
            Aggregator string  `json:"aggregator"`
            PaymentMode string  `json:"payment_mode"`
            ShippingAddress map[string]interface{}  `json:"shipping_address"`
            OrderType string  `json:"order_type"`
            CartID string  `json:"cart_id"`
            PaymentIdentifier string  `json:"payment_identifier"`
            OrderingStore float64  `json:"ordering_store"`
            CurrencyCode string  `json:"currency_code"`
         
    }
    
    // OverrideCheckoutResponse used by Cart
    type OverrideCheckoutResponse struct {

        
            OrderID string  `json:"order_id"`
            Success string  `json:"success"`
            Message string  `json:"message"`
            Cart map[string]interface{}  `json:"cart"`
            Data map[string]interface{}  `json:"data"`
         
    }
    
    // GetShareCartLinkRequest used by Cart
    type GetShareCartLinkRequest struct {

        
            ID string  `json:"id"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // GetShareCartLinkResponse used by Cart
    type GetShareCartLinkResponse struct {

        
            Token string  `json:"token"`
            ShareURL string  `json:"share_url"`
         
    }
    
    // SharedCartDetails used by Cart
    type SharedCartDetails struct {

        
            Meta map[string]interface{}  `json:"meta"`
            Source map[string]interface{}  `json:"source"`
            User map[string]interface{}  `json:"user"`
            CreatedOn string  `json:"created_on"`
            Token string  `json:"token"`
         
    }
    
    // SharedCart used by Cart
    type SharedCart struct {

        
            LastModified string  `json:"last_modified"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            IsValid bool  `json:"is_valid"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            Currency CartCurrency  `json:"currency"`
            Items []CartProductInfo  `json:"items"`
            CartID float64  `json:"cart_id"`
            CheckoutMode string  `json:"checkout_mode"`
            CouponText string  `json:"coupon_text"`
            Gstin string  `json:"gstin"`
            Message string  `json:"message"`
            ID string  `json:"id"`
            UID string  `json:"uid"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            Comment string  `json:"comment"`
            SharedCartDetails SharedCartDetails  `json:"shared_cart_details"`
            BuyNow bool  `json:"buy_now"`
         
    }
    
    // SharedCartResponse used by Cart
    type SharedCartResponse struct {

        
            Error string  `json:"error"`
            Cart SharedCart  `json:"cart"`
         
    }
    
    // CartList used by Cart
    type CartList struct {

        
            CartValue float64  `json:"cart_value"`
            UserID string  `json:"user_id"`
            PickUpCustomerDetails map[string]interface{}  `json:"pick_up_customer_details"`
            CartID string  `json:"cart_id"`
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

        
            Mobile string  `json:"mobile"`
            UID string  `json:"uid"`
            Gender string  `json:"gender"`
            CreatedAt string  `json:"created_at"`
            FirstName string  `json:"first_name"`
            ExternalID string  `json:"external_id"`
            ModifiedOn string  `json:"modified_on"`
            ID string  `json:"_id"`
            LastName string  `json:"last_name"`
         
    }
    
    // UserCartMappingResponse used by Cart
    type UserCartMappingResponse struct {

        
            LastModified string  `json:"last_modified"`
            AppliedPromoDetails []AppliedPromotion  `json:"applied_promo_details"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            IsValid bool  `json:"is_valid"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            User UserInfo  `json:"user"`
            Currency CartCurrency  `json:"currency"`
            Items []CartProductInfo  `json:"items"`
            CheckoutMode string  `json:"checkout_mode"`
            PanConfig map[string]interface{}  `json:"pan_config"`
            CouponText string  `json:"coupon_text"`
            Gstin string  `json:"gstin"`
            Message string  `json:"message"`
            ID string  `json:"id"`
            PanNo string  `json:"pan_no"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            Comment string  `json:"comment"`
            BuyNow bool  `json:"buy_now"`
         
    }
    
    // PlatformAddCartRequest used by Cart
    type PlatformAddCartRequest struct {

        
            UserID string  `json:"user_id"`
            NewCart bool  `json:"new_cart"`
            Items []AddProductCart  `json:"items"`
         
    }
    
    // PlatformUpdateCartRequest used by Cart
    type PlatformUpdateCartRequest struct {

        
            UserID string  `json:"user_id"`
            Operation string  `json:"operation"`
            Items []UpdateProductCart  `json:"items"`
         
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

        
            SubTitle string  `json:"sub_title"`
            MaxDiscountValue float64  `json:"max_discount_value"`
            IsApplied bool  `json:"is_applied"`
            IsBankOffer bool  `json:"is_bank_offer"`
            Savings float64  `json:"savings"`
            ExpiresOn string  `json:"expires_on"`
            Title string  `json:"title"`
            OfferText string  `json:"offer_text"`
            Description string  `json:"description"`
            Message string  `json:"message"`
            CouponValue float64  `json:"coupon_value"`
            MinimumCartValue float64  `json:"minimum_cart_value"`
            IsApplicable bool  `json:"is_applicable"`
            CouponType string  `json:"coupon_type"`
            CouponCode string  `json:"coupon_code"`
         
    }
    
    // PageCoupon used by Cart
    type PageCoupon struct {

        
            HasNext bool  `json:"has_next"`
            Total float64  `json:"total"`
            Current float64  `json:"current"`
            TotalItemCount float64  `json:"total_item_count"`
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

        
            Latitude float64  `json:"latitude"`
            Longitude float64  `json:"longitude"`
         
    }
    
    // PlatformAddress used by Cart
    type PlatformAddress struct {

        
            Meta map[string]interface{}  `json:"meta"`
            Area string  `json:"area"`
            CreatedByUserID string  `json:"created_by_user_id"`
            IsActive bool  `json:"is_active"`
            AreaCode string  `json:"area_code"`
            GoogleMapPoint map[string]interface{}  `json:"google_map_point"`
            CartID string  `json:"cart_id"`
            CheckoutMode string  `json:"checkout_mode"`
            Email string  `json:"email"`
            City string  `json:"city"`
            State string  `json:"state"`
            UserID string  `json:"user_id"`
            Name string  `json:"name"`
            ID string  `json:"id"`
            AddressType string  `json:"address_type"`
            IsDefaultAddress bool  `json:"is_default_address"`
            AreaCodeSlug string  `json:"area_code_slug"`
            CountryCode string  `json:"country_code"`
            GeoLocation GeoLocation  `json:"geo_location"`
            Address string  `json:"address"`
            Landmark string  `json:"landmark"`
            Tags []string  `json:"tags"`
            Phone string  `json:"phone"`
            Country string  `json:"country"`
         
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

        
            IsDefaultAddress bool  `json:"is_default_address"`
            Success bool  `json:"success"`
            ID string  `json:"id"`
            IsUpdated bool  `json:"is_updated"`
         
    }
    
    // DeleteAddressResponse used by Cart
    type DeleteAddressResponse struct {

        
            ID string  `json:"id"`
            IsDeleted bool  `json:"is_deleted"`
         
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
            Meta string  `json:"meta"`
            ArticleID string  `json:"article_id"`
         
    }
    
    // PlatformShipmentResponse used by Cart
    type PlatformShipmentResponse struct {

        
            BoxType string  `json:"box_type"`
            Articles []ShipmentArticle  `json:"articles"`
            ShipmentType string  `json:"shipment_type"`
            Items []CartProductInfo  `json:"items"`
            Promise ShipmentPromise  `json:"promise"`
            OrderType string  `json:"order_type"`
            FulfillmentType string  `json:"fulfillment_type"`
            DpOptions map[string]interface{}  `json:"dp_options"`
            Shipments float64  `json:"shipments"`
            FulfillmentID float64  `json:"fulfillment_id"`
            DpID string  `json:"dp_id"`
         
    }
    
    // PlatformCartShipmentsResponse used by Cart
    type PlatformCartShipmentsResponse struct {

        
            LastModified string  `json:"last_modified"`
            AppliedPromoDetails []AppliedPromotion  `json:"applied_promo_details"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            IsValid bool  `json:"is_valid"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            Currency CartCurrency  `json:"currency"`
            Items []CartProductInfo  `json:"items"`
            CheckoutMode string  `json:"checkout_mode"`
            PanConfig map[string]interface{}  `json:"pan_config"`
            CouponText string  `json:"coupon_text"`
            Error bool  `json:"error"`
            Gstin string  `json:"gstin"`
            Message string  `json:"message"`
            ID string  `json:"id"`
            Shipments []PlatformShipmentResponse  `json:"shipments"`
            PanNo string  `json:"pan_no"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            StaffUserID string  `json:"staff_user_id"`
            Comment string  `json:"comment"`
            BuyNow bool  `json:"buy_now"`
         
    }
    
    // UpdateCartShipmentItem used by Cart
    type UpdateCartShipmentItem struct {

        
            Quantity float64  `json:"quantity"`
            ShipmentType string  `json:"shipment_type"`
            ArticleUID string  `json:"article_uid"`
         
    }
    
    // UpdateCartShipmentRequest used by Cart
    type UpdateCartShipmentRequest struct {

        
            Shipments []UpdateCartShipmentItem  `json:"shipments"`
         
    }
    
    // PlatformCartMetaRequest used by Cart
    type PlatformCartMetaRequest struct {

        
            StaffUserID string  `json:"staff_user_id"`
            Comment string  `json:"comment"`
            GiftDetails map[string]interface{}  `json:"gift_details"`
            PickUpCustomerDetails map[string]interface{}  `json:"pick_up_customer_details"`
            Gstin string  `json:"gstin"`
            CheckoutMode string  `json:"checkout_mode"`
            PanNo string  `json:"pan_no"`
         
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

        
            User string  `json:"user"`
            EmployeeCode string  `json:"employee_code"`
            FirstName string  `json:"first_name"`
            ID string  `json:"_id"`
            LastName string  `json:"last_name"`
         
    }
    
    // Files used by Cart
    type Files struct {

        
            Values []string  `json:"values"`
            Key string  `json:"key"`
         
    }
    
    // PlatformCartCheckoutDetailRequest used by Cart
    type PlatformCartCheckoutDetailRequest struct {

        
            Staff StaffCheckout  `json:"staff"`
            MerchantCode string  `json:"merchant_code"`
            CallbackURL string  `json:"callback_url"`
            PaymentMode string  `json:"payment_mode"`
            Meta map[string]interface{}  `json:"meta"`
            AddressID string  `json:"address_id"`
            EmployeeCode string  `json:"employee_code"`
            OrderType string  `json:"order_type"`
            Pos bool  `json:"pos"`
            PickAtStoreUID float64  `json:"pick_at_store_uid"`
            DeviceID string  `json:"device_id"`
            PaymentParams map[string]interface{}  `json:"payment_params"`
            PaymentAutoConfirm bool  `json:"payment_auto_confirm"`
            CheckoutMode string  `json:"checkout_mode"`
            BillingAddressID string  `json:"billing_address_id"`
            Files []Files  `json:"files"`
            UserID string  `json:"user_id"`
            ID string  `json:"id"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            OrderingStore float64  `json:"ordering_store"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
            Aggregator string  `json:"aggregator"`
            PaymentIdentifier string  `json:"payment_identifier"`
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
         
    }
    
    // CheckCart used by Cart
    type CheckCart struct {

        
            LastModified string  `json:"last_modified"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            StoreCode string  `json:"store_code"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            IsValid bool  `json:"is_valid"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            Currency CartCurrency  `json:"currency"`
            UserType string  `json:"user_type"`
            Items []CartProductInfo  `json:"items"`
            CartID float64  `json:"cart_id"`
            CheckoutMode string  `json:"checkout_mode"`
            CodAvailable bool  `json:"cod_available"`
            CouponText string  `json:"coupon_text"`
            Success bool  `json:"success"`
            ErrorMessage string  `json:"error_message"`
            Gstin string  `json:"gstin"`
            Message string  `json:"message"`
            ID string  `json:"id"`
            DeliveryChargeOrderValue float64  `json:"delivery_charge_order_value"`
            UID string  `json:"uid"`
            CodMessage string  `json:"cod_message"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            OrderID string  `json:"order_id"`
            Comment string  `json:"comment"`
            BuyNow bool  `json:"buy_now"`
            StoreEmps []map[string]interface{}  `json:"store_emps"`
            CodCharges float64  `json:"cod_charges"`
            DeliveryCharges float64  `json:"delivery_charges"`
         
    }
    
    // CartCheckoutResponse used by Cart
    type CartCheckoutResponse struct {

        
            OrderID string  `json:"order_id"`
            CallbackURL string  `json:"callback_url"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
            PaymentConfirmURL string  `json:"payment_confirm_url"`
            AppInterceptURL string  `json:"app_intercept_url"`
            Cart CheckCart  `json:"cart"`
            Data map[string]interface{}  `json:"data"`
         
    }
    
    // CartDeliveryModesResponse used by Cart
    type CartDeliveryModesResponse struct {

        
            PickupStores []float64  `json:"pickup_stores"`
            AvailableModes []string  `json:"available_modes"`
         
    }
    
    // PickupStoreDetail used by Cart
    type PickupStoreDetail struct {

        
            UID float64  `json:"uid"`
            Pincode float64  `json:"pincode"`
            Address string  `json:"address"`
            Landmark string  `json:"landmark"`
            AreaCode string  `json:"area_code"`
            Area string  `json:"area"`
            Name string  `json:"name"`
            Phone string  `json:"phone"`
            State string  `json:"state"`
            AddressType string  `json:"address_type"`
            Email string  `json:"email"`
            StoreCode string  `json:"store_code"`
            ID float64  `json:"id"`
            Country string  `json:"country"`
            City string  `json:"city"`
            AreaCodeSlug string  `json:"area_code_slug"`
         
    }
    
    // StoreDetailsResponse used by Cart
    type StoreDetailsResponse struct {

        
            Items []PickupStoreDetail  `json:"items"`
         
    }
    
    // UpdateCartPaymentRequest used by Cart
    type UpdateCartPaymentRequest struct {

        
            MerchantCode string  `json:"merchant_code"`
            PaymentMode string  `json:"payment_mode"`
            AddressID string  `json:"address_id"`
            AggregatorName string  `json:"aggregator_name"`
            PaymentIdentifier string  `json:"payment_identifier"`
            ID string  `json:"id"`
         
    }
    
    // CouponValidity used by Cart
    type CouponValidity struct {

        
            Code string  `json:"code"`
            DisplayMessageEn string  `json:"display_message_en"`
            Title string  `json:"title"`
            NextValidationRequired bool  `json:"next_validation_required"`
            Discount float64  `json:"discount"`
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

        
            PaymentGateway string  `json:"payment_gateway"`
            PaymentIdentifier string  `json:"payment_identifier"`
            MerchantCode string  `json:"merchant_code"`
            Type string  `json:"type"`
         
    }
    
    // PaymentMethod used by Cart
    type PaymentMethod struct {

        
            Name string  `json:"name"`
            Amount float64  `json:"amount"`
            Mode string  `json:"mode"`
            Payment string  `json:"payment"`
            PaymentMeta PaymentMeta  `json:"payment_meta"`
         
    }
    
    // PlatformCartCheckoutDetailV2Request used by Cart
    type PlatformCartCheckoutDetailV2Request struct {

        
            Staff StaffCheckout  `json:"staff"`
            MerchantCode string  `json:"merchant_code"`
            CallbackURL string  `json:"callback_url"`
            PaymentMode string  `json:"payment_mode"`
            Meta map[string]interface{}  `json:"meta"`
            AddressID string  `json:"address_id"`
            EmployeeCode string  `json:"employee_code"`
            OrderType string  `json:"order_type"`
            Pos bool  `json:"pos"`
            PickAtStoreUID float64  `json:"pick_at_store_uid"`
            DeviceID string  `json:"device_id"`
            PaymentParams map[string]interface{}  `json:"payment_params"`
            PaymentAutoConfirm bool  `json:"payment_auto_confirm"`
            CheckoutMode string  `json:"checkout_mode"`
            BillingAddressID string  `json:"billing_address_id"`
            Files []Files  `json:"files"`
            PaymentMethods []PaymentMethod  `json:"payment_methods"`
            UserID string  `json:"user_id"`
            ID string  `json:"id"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            OrderingStore float64  `json:"ordering_store"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
            Aggregator string  `json:"aggregator"`
            PaymentIdentifier string  `json:"payment_identifier"`
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
            CustomMeta map[string]interface{}  `json:"custom_meta"`
         
    }
    
    // UpdateCartPaymentRequestV2 used by Cart
    type UpdateCartPaymentRequestV2 struct {

        
            MerchantCode string  `json:"merchant_code"`
            PaymentMethods []PaymentMethod  `json:"payment_methods"`
            PaymentMode string  `json:"payment_mode"`
            AddressID string  `json:"address_id"`
            AggregatorName string  `json:"aggregator_name"`
            PaymentIdentifier string  `json:"payment_identifier"`
            ID string  `json:"id"`
         
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

        
            Message string  `json:"message"`
            Type string  `json:"type"`
            Value string  `json:"value"`
         
    }
    
    // ApplicationServiceabilityConfig used by Serviceability
    type ApplicationServiceabilityConfig struct {

        
            ChannelID string  `json:"channel_id"`
            ServiceabilityType string  `json:"serviceability_type"`
            ChannelType string  `json:"channel_type"`
         
    }
    
    // ApplicationServiceabilityConfigResponse used by Serviceability
    type ApplicationServiceabilityConfigResponse struct {

        
            Success bool  `json:"success"`
            Error ServiceabilityErrorResponse  `json:"error"`
            Data ApplicationServiceabilityConfig  `json:"data"`
         
    }
    
    // EntityRegionView_Request used by Serviceability
    type EntityRegionView_Request struct {

        
            ParentID []string  `json:"parent_id"`
            SubType []string  `json:"sub_type"`
         
    }
    
    // EntityRegionView_Error used by Serviceability
    type EntityRegionView_Error struct {

        
            Message string  `json:"message"`
            Type string  `json:"type"`
            Value string  `json:"value"`
         
    }
    
    // EntityRegionView_Items used by Serviceability
    type EntityRegionView_Items struct {

        
            UID string  `json:"uid"`
            Name string  `json:"name"`
            SubType string  `json:"sub_type"`
         
    }
    
    // EntityRegionView_page used by Serviceability
    type EntityRegionView_page struct {

        
            ItemTotal float64  `json:"item_total"`
            Current float64  `json:"current"`
            Type string  `json:"type"`
            Size float64  `json:"size"`
            HasNext bool  `json:"has_next"`
         
    }
    
    // EntityRegionView_Response used by Serviceability
    type EntityRegionView_Response struct {

        
            Success bool  `json:"success"`
            Error EntityRegionView_Error  `json:"error"`
            Data []EntityRegionView_Items  `json:"data"`
            Page EntityRegionView_page  `json:"page"`
         
    }
    
    // ListViewChannels used by Serviceability
    type ListViewChannels struct {

        
            ChannelID string  `json:"channel_id"`
            ChannelType string  `json:"channel_type"`
         
    }
    
    // ListViewProduct used by Serviceability
    type ListViewProduct struct {

        
            Type string  `json:"type"`
            Count float64  `json:"count"`
         
    }
    
    // ListViewItems used by Serviceability
    type ListViewItems struct {

        
            Channels ListViewChannels  `json:"channels"`
            IsActive bool  `json:"is_active"`
            Slug string  `json:"slug"`
            StoresCount float64  `json:"stores_count"`
            ZoneID string  `json:"zone_id"`
            PincodesCount float64  `json:"pincodes_count"`
            CompanyID float64  `json:"company_id"`
            Product ListViewProduct  `json:"product"`
            Name string  `json:"name"`
         
    }
    
    // ListViewSummary used by Serviceability
    type ListViewSummary struct {

        
            TotalPincodesServed float64  `json:"total_pincodes_served"`
            TotalZones float64  `json:"total_zones"`
            TotalActiveZones float64  `json:"total_active_zones"`
         
    }
    
    // ZoneDataItem used by Serviceability
    type ZoneDataItem struct {

        
            Type string  `json:"type"`
            ItemTotal float64  `json:"item_total"`
            Current float64  `json:"current"`
            Size float64  `json:"size"`
            HasNext bool  `json:"has_next"`
         
    }
    
    // ListViewResponse used by Serviceability
    type ListViewResponse struct {

        
            Items []ListViewItems  `json:"items"`
            Summary []ListViewSummary  `json:"summary"`
            Page []ZoneDataItem  `json:"page"`
         
    }
    
    // CompanyStoreView_PageItems used by Serviceability
    type CompanyStoreView_PageItems struct {

        
            ItemTotal float64  `json:"item_total"`
            Current float64  `json:"current"`
            Type string  `json:"type"`
            Size float64  `json:"size"`
            HasNext bool  `json:"has_next"`
         
    }
    
    // CompanyStoreView_Response used by Serviceability
    type CompanyStoreView_Response struct {

        
            Items []map[string]interface{}  `json:"items"`
            Page []CompanyStoreView_PageItems  `json:"page"`
         
    }
    
    // GetZoneDataViewChannels used by Serviceability
    type GetZoneDataViewChannels struct {

        
            ChannelID string  `json:"channel_id"`
            ChannelType string  `json:"channel_type"`
         
    }
    
    // ZoneProductTypes used by Serviceability
    type ZoneProductTypes struct {

        
            Type string  `json:"type"`
            Tags []string  `json:"tags"`
         
    }
    
    // ZoneMappingType used by Serviceability
    type ZoneMappingType struct {

        
            Country string  `json:"country"`
            State []string  `json:"state"`
            Pincode []string  `json:"pincode"`
         
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

        
            Success bool  `json:"success"`
            ZoneID string  `json:"zone_id"`
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

        
            Items []ListViewItems  `json:"items"`
            Page []ZoneDataItem  `json:"page"`
         
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
    
    // Dp used by Serviceability
    type Dp struct {

        
            InternalAccountID string  `json:"internal_account_id"`
            PaymentMode string  `json:"payment_mode"`
            TransportMode string  `json:"transport_mode"`
            LmPriority float64  `json:"lm_priority"`
            RvpPriority float64  `json:"rvp_priority"`
            AreaCode float64  `json:"area_code"`
            FmPriority float64  `json:"fm_priority"`
            AssignDpFromSb bool  `json:"assign_dp_from_sb"`
            ExternalAccountID string  `json:"external_account_id"`
            Operations []string  `json:"operations"`
         
    }
    
    // LogisticsResponse used by Serviceability
    type LogisticsResponse struct {

        
            Override bool  `json:"override"`
            Dp Dp  `json:"dp"`
         
    }
    
    // IntegrationTypeResponse used by Serviceability
    type IntegrationTypeResponse struct {

        
            Inventory string  `json:"inventory"`
            Order string  `json:"order"`
         
    }
    
    // ModifiedByResponse used by Serviceability
    type ModifiedByResponse struct {

        
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
         
    }
    
    // AddressResponse used by Serviceability
    type AddressResponse struct {

        
            Address2 string  `json:"address2"`
            Country string  `json:"country"`
            Longitude float64  `json:"longitude"`
            Pincode float64  `json:"pincode"`
            Landmark string  `json:"landmark"`
            Latitude float64  `json:"latitude"`
            City string  `json:"city"`
            Address1 string  `json:"address1"`
            State string  `json:"state"`
         
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
    
    // OpeningClosing used by Serviceability
    type OpeningClosing struct {

        
            Hour float64  `json:"hour"`
            Minute float64  `json:"minute"`
         
    }
    
    // TimmingResponse used by Serviceability
    type TimmingResponse struct {

        
            Closing OpeningClosing  `json:"closing"`
            Opening OpeningClosing  `json:"opening"`
            Open bool  `json:"open"`
            Weekday string  `json:"weekday"`
         
    }
    
    // ContactNumberResponse used by Serviceability
    type ContactNumberResponse struct {

        
            Number string  `json:"number"`
            CountryCode float64  `json:"country_code"`
         
    }
    
    // DocumentsResponse used by Serviceability
    type DocumentsResponse struct {

        
            LegalName string  `json:"legal_name"`
            Type string  `json:"type"`
            Value string  `json:"value"`
            Verified bool  `json:"verified"`
         
    }
    
    // CreatedByResponse used by Serviceability
    type CreatedByResponse struct {

        
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
         
    }
    
    // WarningsResponse used by Serviceability
    type WarningsResponse struct {

        
            StoreAddress string  `json:"store_address"`
         
    }
    
    // ProductReturnConfigResponse used by Serviceability
    type ProductReturnConfigResponse struct {

        
            OnSameStore bool  `json:"on_same_store"`
         
    }
    
    // ItemResponse used by Serviceability
    type ItemResponse struct {

        
            Stage string  `json:"stage"`
            Company float64  `json:"company"`
            NotificationEmails []string  `json:"notification_emails"`
            GstCredentials GstCredentialsResponse  `json:"gst_credentials"`
            Logistics LogisticsResponse  `json:"logistics"`
            IntegrationType IntegrationTypeResponse  `json:"integration_type"`
            ModifiedBy ModifiedByResponse  `json:"modified_by"`
            DisplayName string  `json:"display_name"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            CompanyID float64  `json:"company_id"`
            Address AddressResponse  `json:"address"`
            UID float64  `json:"uid"`
            Cls string  `json:"_cls"`
            Manager ManagerResponse  `json:"manager"`
            SubType string  `json:"sub_type"`
            Name string  `json:"name"`
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
            Timing []TimmingResponse  `json:"timing"`
            Code string  `json:"code"`
            StoreType string  `json:"store_type"`
            VerifiedBy ModifiedByResponse  `json:"verified_by"`
            VerifiedOn string  `json:"verified_on"`
            ContactNumbers []ContactNumberResponse  `json:"contact_numbers"`
            Documents []DocumentsResponse  `json:"documents"`
            CreatedBy CreatedByResponse  `json:"created_by"`
            Warnings WarningsResponse  `json:"warnings"`
            ProductReturnConfig ProductReturnConfigResponse  `json:"product_return_config"`
         
    }
    
    // ServiceabilityPageResponse used by Serviceability
    type ServiceabilityPageResponse struct {

        
            ItemTotal float64  `json:"item_total"`
            Current float64  `json:"current"`
            Type string  `json:"type"`
            Size float64  `json:"size"`
            HasNext bool  `json:"has_next"`
         
    }
    
    // GetStoresViewResponse used by Serviceability
    type GetStoresViewResponse struct {

        
            Items []ItemResponse  `json:"items"`
            Page ServiceabilityPageResponse  `json:"page"`
         
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

        
            Success bool  `json:"success"`
            Articles []map[string]interface{}  `json:"articles"`
            Error map[string]interface{}  `json:"error"`
            ToPincode string  `json:"to_pincode"`
         
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
            Error interface{}  `json:"error"`
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

        
            ItemTotal float64  `json:"item_total"`
            Current float64  `json:"current"`
            Type string  `json:"type"`
            Size float64  `json:"size"`
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
    
    // BulkRegionJobSerializer used by Serviceability
    type BulkRegionJobSerializer struct {

        
            Action string  `json:"action"`
            CountryIsoCode string  `json:"country_iso_code"`
            JobAction string  `json:"job_action"`
            FileURL string  `json:"file_url"`
            BatchID string  `json:"batch_id"`
         
    }
    
    // PostBulkRegionJobResponse used by Serviceability
    type PostBulkRegionJobResponse struct {

        
            Message string  `json:"message"`
            Response bool  `json:"response"`
            EventEmitted bool  `json:"event_emitted"`
            BatchID string  `json:"batch_id"`
         
    }
    
    // BulkRecordError used by Serviceability
    type BulkRecordError struct {

        
            IsError bool  `json:"is_error"`
            Error []string  `json:"error"`
         
    }
    
    // CSVFileRecord used by Serviceability
    type CSVFileRecord struct {

        
            Country string  `json:"country"`
            TatType string  `json:"tat_type"`
            Error []string  `json:"error"`
            DpID float64  `json:"dp_id"`
            FromRegion string  `json:"from_region"`
            PlanID float64  `json:"plan_id"`
            IsError bool  `json:"is_error"`
            MinTat float64  `json:"min_tat"`
            MaxTat float64  `json:"max_tat"`
            SNo float64  `json:"s_no"`
            ToRegion string  `json:"to_region"`
            RegionType string  `json:"region_type"`
         
    }
    
    // BulkRegionData used by Serviceability
    type BulkRegionData struct {

        
            SuccessCount float64  `json:"success_count"`
            Action string  `json:"action"`
            Error BulkRecordError  `json:"error"`
            CreatedOn string  `json:"created_on"`
            Stage string  `json:"stage"`
            FailedRec []CSVFileRecord  `json:"failed_rec"`
            BatchID string  `json:"batch_id"`
            TotalRec float64  `json:"total_rec"`
            FilePath string  `json:"file_path"`
            FailedCount float64  `json:"failed_count"`
         
    }
    
    // GetBulkRegionJobResponse used by Serviceability
    type GetBulkRegionJobResponse struct {

        
            BatchID string  `json:"batch_id"`
            CurrentPageNumber float64  `json:"current_page_number"`
            Data []BulkRegionData  `json:"data"`
         
    }
    
    // Dp1 used by Serviceability
    type Dp1 struct {

        
            AccountID string  `json:"account_id"`
            DpID string  `json:"dp_id"`
            Stage string  `json:"stage"`
            PlanID string  `json:"plan_id"`
            IsSelfShip bool  `json:"is_self_ship"`
            Name string  `json:"name"`
            PlanRules map[string]interface{}  `json:"plan_rules"`
         
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

        
            Message string  `json:"message"`
            Type string  `json:"type"`
            Value string  `json:"value"`
         
    }
    
    // DpAccountFailureResponse used by Serviceability
    type DpAccountFailureResponse struct {

        
            Success bool  `json:"success"`
            Error []ErrorResponse  `json:"error"`
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

        
            Items []Dp1  `json:"items"`
            Success bool  `json:"success"`
            Page Page  `json:"page"`
         
    }
    
    // DpSchemaInRuleListing used by Serviceability
    type DpSchemaInRuleListing struct {

        
            AccountID string  `json:"account_id"`
            DpID string  `json:"dp_id"`
            Stage string  `json:"stage"`
            PlanID string  `json:"plan_id"`
            IsSelfShip bool  `json:"is_self_ship"`
            Priority float64  `json:"priority"`
            Name string  `json:"name"`
            PlanRules map[string]interface{}  `json:"plan_rules"`
         
    }
    
    // DpRule used by Serviceability
    type DpRule struct {

        
            IsActive bool  `json:"is_active"`
            CompanyID float64  `json:"company_id"`
            DpIds map[string]DpSchemaInRuleListing  `json:"dp_ids"`
            Name string  `json:"name"`
            Conditions []map[string]interface{}  `json:"conditions"`
         
    }
    
    // DpRuleSuccessResponse used by Serviceability
    type DpRuleSuccessResponse struct {

        
            Success bool  `json:"success"`
            Data DpRule  `json:"data"`
            StatusCode float64  `json:"status_code"`
         
    }
    
    // FailureResponse used by Serviceability
    type FailureResponse struct {

        
            Success bool  `json:"success"`
            Error []ErrorResponse  `json:"error"`
            StatusCode float64  `json:"status_code"`
         
    }
    
    // DpRulesUpdateRequest used by Serviceability
    type DpRulesUpdateRequest struct {

        
            DpIds map[string]map[string]interface{}  `json:"dp_ids"`
            Name string  `json:"name"`
            Conditions []map[string]interface{}  `json:"conditions"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // DpRuleResponse used by Serviceability
    type DpRuleResponse struct {

        
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            UID string  `json:"uid"`
            CreatedOn string  `json:"created_on"`
            IsActive bool  `json:"is_active"`
            CompanyID float64  `json:"company_id"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            DpIds map[string]interface{}  `json:"dp_ids"`
            Name string  `json:"name"`
            Conditions []string  `json:"conditions"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // DpRuleUpdateSuccessResponse used by Serviceability
    type DpRuleUpdateSuccessResponse struct {

        
            Success bool  `json:"success"`
            Data DpRuleResponse  `json:"data"`
            StatusCode float64  `json:"status_code"`
         
    }
    
    // DpIds used by Serviceability
    type DpIds struct {

        
            Priority float64  `json:"priority"`
            Meta map[string]interface{}  `json:"meta"`
            Enabled bool  `json:"enabled"`
         
    }
    
    // DpRuleRequest used by Serviceability
    type DpRuleRequest struct {

        
            IsActive bool  `json:"is_active"`
            CompanyID float64  `json:"company_id"`
            DpIds map[string]DpIds  `json:"dp_ids"`
            Name string  `json:"name"`
            Conditions []map[string]interface{}  `json:"conditions"`
         
    }
    
    // DpMultipleRuleSuccessResponse used by Serviceability
    type DpMultipleRuleSuccessResponse struct {

        
            Items []DpRule  `json:"items"`
            Success bool  `json:"success"`
            Page Page  `json:"page"`
         
    }
    
    // DPCompanyRuleResponse used by Serviceability
    type DPCompanyRuleResponse struct {

        
            Success bool  `json:"success"`
            Data []DpRuleResponse  `json:"data"`
            StatusCode float64  `json:"status_code"`
         
    }
    
    // DPCompanyRuleRequest used by Serviceability
    type DPCompanyRuleRequest struct {

        
            RuleIds []string  `json:"rule_ids"`
         
    }
    
    // DPApplicationRuleResponse used by Serviceability
    type DPApplicationRuleResponse struct {

        
            Success bool  `json:"success"`
            Data []DpRuleResponse  `json:"data"`
            StatusCode bool  `json:"status_code"`
         
    }
    
    // DPApplicationRuleRequest used by Serviceability
    type DPApplicationRuleRequest struct {

        
            ShippingRules []string  `json:"shipping_rules"`
         
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
            Error ServiceabilityErrorResponse  `json:"error"`
            Data ApplicationSelfShipConfig  `json:"data"`
         
    }
    

    
    // GenerateReportMeta used by Finance
    type GenerateReportMeta struct {

        
            Company string  `json:"company"`
            Brand string  `json:"brand"`
            Channel string  `json:"channel"`
         
    }
    
    // GenerateReportFilters used by Finance
    type GenerateReportFilters struct {

        
            Company []string  `json:"company"`
            Brand []string  `json:"brand"`
            Channel []string  `json:"channel"`
         
    }
    
    // GenerateReportPlatform used by Finance
    type GenerateReportPlatform struct {

        
            Meta GenerateReportMeta  `json:"meta"`
            StartDate string  `json:"start_date"`
            EndDate string  `json:"end_date"`
            Filters GenerateReportFilters  `json:"filters"`
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

        
            StartDate string  `json:"start_date"`
            ItemCount float64  `json:"item_count"`
            EndDate string  `json:"end_date"`
            Headers []string  `json:"headers"`
            Items [][]string  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // Error used by Finance
    type Error struct {

        
            Success bool  `json:"success"`
            Reason string  `json:"reason"`
         
    }
    
    // DownloadReport used by Finance
    type DownloadReport struct {

        
            Page float64  `json:"page"`
            Pagesize float64  `json:"pagesize"`
            EndDate string  `json:"end_date"`
            StartDate string  `json:"start_date"`
         
    }
    
    // DownloadReportItems used by Finance
    type DownloadReportItems struct {

        
            TypeOfRequest string  `json:"type_of_request"`
            Meta GenerateReportMeta  `json:"meta"`
            StartDate string  `json:"start_date"`
            EndDate string  `json:"end_date"`
            Filters GenerateReportFilters  `json:"filters"`
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

        
            Project []string  `json:"project"`
            TableName string  `json:"table_name"`
            Filters GetEngineFilters  `json:"filters"`
         
    }
    
    // GetEngineRequest used by Finance
    type GetEngineRequest struct {

        
            Data GetEngineData  `json:"data"`
         
    }
    
    // GetEngineResponse used by Finance
    type GetEngineResponse struct {

        
            Success bool  `json:"success"`
            Items []map[string]interface{}  `json:"items"`
            ItemCount float64  `json:"item_count"`
            Page Page  `json:"page"`
         
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

        
            Meta map[string]interface{}  `json:"meta"`
            InvoiceNumber string  `json:"invoice_number"`
            SourceReference string  `json:"source_reference"`
            Amount string  `json:"amount"`
            Platform string  `json:"platform"`
            TotalAmount string  `json:"total_amount"`
            TransactionType string  `json:"transaction_type"`
            ModeOfPayment string  `json:"mode_of_payment"`
            SellerID string  `json:"seller_id"`
            Currency string  `json:"currency"`
         
    }
    
    // PaymentProcessRequest used by Finance
    type PaymentProcessRequest struct {

        
            Data PaymentProcessPayload  `json:"data"`
         
    }
    
    // PaymentProcessResponse used by Finance
    type PaymentProcessResponse struct {

        
            Meta map[string]interface{}  `json:"meta"`
            Message string  `json:"message"`
            TransactionID string  `json:"transaction_id"`
            Code float64  `json:"code"`
            RedirectURL string  `json:"redirect_url"`
         
    }
    
    // CreditlineDataPlatformPayload used by Finance
    type CreditlineDataPlatformPayload struct {

        
            StartEnd string  `json:"start_end"`
            EndEnd string  `json:"end_end"`
            SellerID string  `json:"seller_id"`
            Pagesize float64  `json:"pagesize"`
            Page float64  `json:"page"`
         
    }
    
    // CreditlineDataPlatformRequest used by Finance
    type CreditlineDataPlatformRequest struct {

        
            Data CreditlineDataPlatformPayload  `json:"data"`
         
    }
    
    // CreditlineDataPlatformResponse used by Finance
    type CreditlineDataPlatformResponse struct {

        
            ShowMr bool  `json:"show_mr"`
            Message string  `json:"message"`
            ItemCount float64  `json:"item_count"`
            Code float64  `json:"code"`
            Headers []string  `json:"headers"`
            Items []map[string]interface{}  `json:"items"`
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

        
            Success bool  `json:"success"`
            InvoiceTypeList []InvoiceTypeResponseItems  `json:"invoice_type_list"`
            PaymentStatusList []InvoiceTypeResponseItems  `json:"payment_status_list"`
         
    }
    
    // InoviceListingPayloadDataFilters used by Finance
    type InoviceListingPayloadDataFilters struct {

        
            InvoiceType []string  `json:"invoice_type"`
            PaymentStatus []string  `json:"payment_status"`
            CompanyID []string  `json:"company_id"`
         
    }
    
    // InvoiceListingPayloadData used by Finance
    type InvoiceListingPayloadData struct {

        
            Search string  `json:"search"`
            StartDate string  `json:"start_date"`
            PageSize float64  `json:"page_size"`
            EndDate string  `json:"end_date"`
            Filters InoviceListingPayloadDataFilters  `json:"filters"`
            Page float64  `json:"page"`
         
    }
    
    // InvoiceListingRequest used by Finance
    type InvoiceListingRequest struct {

        
            Data InvoiceListingPayloadData  `json:"data"`
         
    }
    
    // UnpaidInvoiceDataItems used by Finance
    type UnpaidInvoiceDataItems struct {

        
            Currency string  `json:"currency"`
            TotalUnpaidAmount float64  `json:"total_unpaid_amount"`
            TotalUnpaidInvoiceCount float64  `json:"total_unpaid_invoice_count"`
         
    }
    
    // InvoiceListingResponseItems used by Finance
    type InvoiceListingResponseItems struct {

        
            InvoiceNumber string  `json:"invoice_number"`
            Status string  `json:"status"`
            InvoiceDate string  `json:"invoice_date"`
            Amount string  `json:"amount"`
            InvoiceID string  `json:"invoice_id"`
            Company string  `json:"company"`
            DueDate string  `json:"due_date"`
            Period string  `json:"period"`
            InvoiceType string  `json:"invoice_type"`
            IsDownloadable bool  `json:"is_downloadable"`
         
    }
    
    // InvoiceListingResponse used by Finance
    type InvoiceListingResponse struct {

        
            UnpaidInvoiceData UnpaidInvoiceDataItems  `json:"unpaid_invoice_data"`
            Items []InvoiceListingResponseItems  `json:"items"`
            ItemCount float64  `json:"item_count"`
            Page Page  `json:"page"`
         
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
            Error []string  `json:"error"`
            Data []string  `json:"data"`
         
    }
    
    // AsCnRefundData used by Finance
    type AsCnRefundData struct {

        
            SellerID float64  `json:"seller_id"`
            AffiliateID string  `json:"affiliate_id"`
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

        
            AffiliateID string  `json:"affiliate_id"`
            NotificationEvents CreditNoteConfigNotificationEvents  `json:"notification_events"`
            SourceChannel []string  `json:"source_channel"`
            IsCnAsRefundMethod bool  `json:"is_cn_as_refund_method"`
            Validity float64  `json:"validity"`
            SellerID float64  `json:"seller_id"`
            OrderingChannel []string  `json:"ordering_channel"`
            SalesChannelName string  `json:"sales_channel_name"`
            CurrencyType string  `json:"currency_type"`
            SlugValues []string  `json:"slug_values"`
         
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

        
            SellerID float64  `json:"seller_id"`
            AffiliateID string  `json:"affiliate_id"`
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

        
            CreatedAt string  `json:"created_at"`
            ShipmentID string  `json:"shipment_id"`
            InvoiceNumber string  `json:"invoice_number"`
            AmountDebited float64  `json:"amount_debited"`
            OrderingChannel string  `json:"ordering_channel"`
            StoreID string  `json:"store_id"`
            StaffID string  `json:"staff_id"`
            OrderID string  `json:"order_id"`
         
    }
    
    // CnDetails used by Finance
    type CnDetails struct {

        
            ShipmentID string  `json:"shipment_id"`
            InvoiceNumber string  `json:"invoice_number"`
            OrderingChannel string  `json:"ordering_channel"`
            StoreID string  `json:"store_id"`
            DateIssued string  `json:"date_issued"`
            StaffID string  `json:"staff_id"`
            ChannelOfIssuance string  `json:"channel_of_issuance"`
            OrderID string  `json:"order_id"`
            ExpiryDate string  `json:"expiry_date"`
         
    }
    
    // CreditNoteDetails used by Finance
    type CreditNoteDetails struct {

        
            RemainingCnAmount float64  `json:"remaining_cn_amount"`
            CustomerMobileNumber string  `json:"customer_mobile_number"`
            RedemptionDetails []RedemptionDetails  `json:"redemption_details"`
            CnReferenceNumber string  `json:"cn_reference_number"`
            CnAmount float64  `json:"cn_amount"`
            AvailableCnBalance float64  `json:"available_cn_balance"`
            CnDetails []CnDetails  `json:"cn_details"`
            CnStatus string  `json:"cn_status"`
         
    }
    
    // CreditNoteDetailsResponse used by Finance
    type CreditNoteDetailsResponse struct {

        
            Success bool  `json:"success"`
            Data CreditNoteDetails  `json:"data"`
         
    }
    
    // GetCustomerCreditBalance used by Finance
    type GetCustomerCreditBalance struct {

        
            SellerID float64  `json:"seller_id"`
            AffiliateID string  `json:"affiliate_id"`
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

        
            Reason string  `json:"reason"`
            SourceChannel []string  `json:"source_channel"`
         
    }
    
    // GetCnConfigResponseData used by Finance
    type GetCnConfigResponseData struct {

        
            AffiliateID string  `json:"affiliate_id"`
            Meta GetCnConfigResponseMeta  `json:"meta"`
            NotificationEvents CreditNoteConfigNotificationEvents  `json:"notification_events"`
            IsCnAsRefundMethod bool  `json:"is_cn_as_refund_method"`
            Validity float64  `json:"validity"`
            RedemptionOrderingChannel []string  `json:"redemption_ordering_channel"`
            SellerID float64  `json:"seller_id"`
            CurrencyType string  `json:"currency_type"`
         
    }
    
    // GetCnConfigResponse used by Finance
    type GetCnConfigResponse struct {

        
            Success bool  `json:"success"`
            Data GetCnConfigResponseData  `json:"data"`
         
    }
    
    // CnGenerateReportFilters used by Finance
    type CnGenerateReportFilters struct {

        
            Utilisation []string  `json:"utilisation"`
            TypesOfTransaction []string  `json:"types_of_transaction"`
            OrderingChannel []string  `json:"ordering_channel"`
            StoreID []float64  `json:"store_id"`
            StaffID []string  `json:"staff_id"`
            ChannelOfIssuance []string  `json:"channel_of_issuance"`
         
    }
    
    // CnGenerateReport used by Finance
    type CnGenerateReport struct {

        
            AffiliateID string  `json:"affiliate_id"`
            Meta GenerateReportFilters  `json:"meta"`
            Search string  `json:"search"`
            StartDate string  `json:"start_date"`
            EndDate string  `json:"end_date"`
            Filters CnGenerateReportFilters  `json:"filters"`
            SearchType string  `json:"search_type"`
            Pagesize float64  `json:"pagesize"`
            ReportID string  `json:"report_id"`
            Page float64  `json:"page"`
         
    }
    
    // GenerateReportCustomerCnRequest used by Finance
    type GenerateReportCustomerCnRequest struct {

        
            Data CnGenerateReport  `json:"data"`
         
    }
    
    // CnGenerateReportItems used by Finance
    type CnGenerateReportItems struct {

        
            ShipmentID string  `json:"shipment_id"`
            CreditNoteNumber string  `json:"credit_note_number"`
            InvoiceNumber string  `json:"invoice_number"`
            Status string  `json:"status"`
            DateIssued string  `json:"date_issued"`
            TotalAmount float64  `json:"total_amount"`
            OrderID string  `json:"order_id"`
            ExpiryDate string  `json:"expiry_date"`
         
    }
    
    // GenerateReportCustomerCnResponse used by Finance
    type GenerateReportCustomerCnResponse struct {

        
            RowHeaderDisplayOrder map[string]interface{}  `json:"row_header_display_order"`
            StartDate string  `json:"start_date"`
            EndDate string  `json:"end_date"`
            ItemCount float64  `json:"item_count"`
            PrimaryHeaders []string  `json:"primary_headers"`
            Headers []string  `json:"headers"`
            AllowedFilters []string  `json:"allowed_filters"`
            Items []CnGenerateReportItems  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // CnDownloadReport used by Finance
    type CnDownloadReport struct {

        
            AffiliateID string  `json:"affiliate_id"`
            Search string  `json:"search"`
            Status []string  `json:"status"`
            StartDate string  `json:"start_date"`
            EndDate string  `json:"end_date"`
            SearchType string  `json:"search_type"`
            Pagesize float64  `json:"pagesize"`
            Page float64  `json:"page"`
         
    }
    
    // DownloadReportCustomerCnRequest used by Finance
    type DownloadReportCustomerCnRequest struct {

        
            Data CnDownloadReport  `json:"data"`
         
    }
    
    // DownloadReportResponseData used by Finance
    type DownloadReportResponseData struct {

        
            RequestedBy string  `json:"requested_by"`
            Meta map[string]interface{}  `json:"meta"`
            CreatedAt string  `json:"created_at"`
            DisplayName string  `json:"display_name"`
            Msg string  `json:"msg"`
            FullName string  `json:"full_name"`
            Status string  `json:"status"`
            StartDate string  `json:"start_date"`
            EndDate string  `json:"end_date"`
            ReportConfigID string  `json:"report_config_id"`
            Filters map[string]interface{}  `json:"filters"`
            DownloadLink string  `json:"download_link"`
            ReportName string  `json:"report_name"`
            RequestDict map[string]interface{}  `json:"request_dict"`
         
    }
    
    // DownloadReportCustomerCnResponse used by Finance
    type DownloadReportCustomerCnResponse struct {

        
            Data []DownloadReportResponseData  `json:"data"`
         
    }
    
    // GetReportingFilters used by Finance
    type GetReportingFilters struct {

        
            Type string  `json:"type"`
            Text string  `json:"text"`
            Value string  `json:"value"`
            Options []map[string]interface{}  `json:"options"`
         
    }
    
    // GetReportingNestedFilters used by Finance
    type GetReportingNestedFilters struct {

        
            Required bool  `json:"required"`
            Options []map[string]interface{}  `json:"options"`
            Type string  `json:"type"`
            Text string  `json:"text"`
            PlaceholderText string  `json:"placeholder_text"`
            Value string  `json:"value"`
         
    }
    
    // GetReportingFiltersResponse used by Finance
    type GetReportingFiltersResponse struct {

        
            Search GetReportingFilters  `json:"search"`
            Status GetReportingFilters  `json:"status"`
            Filters []GetReportingNestedFilters  `json:"filters"`
         
    }
    
