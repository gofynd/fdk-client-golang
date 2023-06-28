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
            DisplayFields []string  `json:"display_fields"`
            ExcludedFields []string  `json:"excluded_fields"`
            Success bool  `json:"success"`
            Created bool  `json:"created"`
            AppID string  `json:"app_id"`
         
    }
    
    // ErrorCodeDescription used by Payment
    type ErrorCodeDescription struct {

        
            Code string  `json:"code"`
            Success bool  `json:"success"`
            Description string  `json:"description"`
         
    }
    
    // PaymentGatewayConfig used by Payment
    type PaymentGatewayConfig struct {

        
            ConfigType string  `json:"config_type"`
            MerchantSalt string  `json:"merchant_salt"`
            IsActive bool  `json:"is_active"`
            Key string  `json:"key"`
            Secret string  `json:"secret"`
         
    }
    
    // PaymentGatewayConfigRequest used by Payment
    type PaymentGatewayConfigRequest struct {

        
            AppID string  `json:"app_id"`
            AggregatorName PaymentGatewayConfig  `json:"aggregator_name"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // PaymentGatewayToBeReviewed used by Payment
    type PaymentGatewayToBeReviewed struct {

        
            Aggregator []string  `json:"aggregator"`
            Success bool  `json:"success"`
         
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

        
            PackageName string  `json:"package_name"`
            DisplayName string  `json:"display_name"`
            Code string  `json:"code"`
            Logos PaymentModeLogo  `json:"logos"`
         
    }
    
    // PaymentModeList used by Payment
    type PaymentModeList struct {

        
            CardBrand string  `json:"card_brand"`
            CardName string  `json:"card_name"`
            Nickname string  `json:"nickname"`
            CompliantWithTokenisationGuidelines bool  `json:"compliant_with_tokenisation_guidelines"`
            IntentAppErrorList []string  `json:"intent_app_error_list"`
            DisplayPriority float64  `json:"display_priority"`
            CardIssuer string  `json:"card_issuer"`
            CardReference string  `json:"card_reference"`
            AggregatorName string  `json:"aggregator_name"`
            CodLimit float64  `json:"cod_limit"`
            IntentAppErrorDictList []IntentAppErrorList  `json:"intent_app_error_dict_list"`
            CardBrandImage string  `json:"card_brand_image"`
            CardType string  `json:"card_type"`
            MerchantCode string  `json:"merchant_code"`
            Timeout float64  `json:"timeout"`
            RetryCount float64  `json:"retry_count"`
            IntentApp []IntentApp  `json:"intent_app"`
            DisplayName string  `json:"display_name"`
            IntentFlow bool  `json:"intent_flow"`
            CardNumber string  `json:"card_number"`
            LogoURL PaymentModeLogo  `json:"logo_url"`
            RemainingLimit float64  `json:"remaining_limit"`
            CardFingerprint string  `json:"card_fingerprint"`
            FyndVpa string  `json:"fynd_vpa"`
            Name string  `json:"name"`
            CardIsin string  `json:"card_isin"`
            Expired bool  `json:"expired"`
            ExpYear float64  `json:"exp_year"`
            Code string  `json:"code"`
            ExpMonth float64  `json:"exp_month"`
            CardID string  `json:"card_id"`
            CodLimitPerOrder float64  `json:"cod_limit_per_order"`
            CardToken string  `json:"card_token"`
         
    }
    
    // RootPaymentMode used by Payment
    type RootPaymentMode struct {

        
            List []PaymentModeList  `json:"list"`
            DisplayPriority float64  `json:"display_priority"`
            AddCardEnabled bool  `json:"add_card_enabled"`
            DisplayName string  `json:"display_name"`
            AnonymousEnable bool  `json:"anonymous_enable"`
            SaveCard bool  `json:"save_card"`
            AggregatorName string  `json:"aggregator_name"`
            IsPayByCardPl bool  `json:"is_pay_by_card_pl"`
            Name string  `json:"name"`
         
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

        
            TransferType string  `json:"transfer_type"`
            PayoutsAggregators []map[string]interface{}  `json:"payouts_aggregators"`
            Customers map[string]interface{}  `json:"customers"`
            UniqueTransferNo map[string]interface{}  `json:"unique_transfer_no"`
            IsDefault bool  `json:"is_default"`
            IsActive bool  `json:"is_active"`
            MoreAttributes map[string]interface{}  `json:"more_attributes"`
         
    }
    
    // PayoutBankDetails used by Payment
    type PayoutBankDetails struct {

        
            AccountType string  `json:"account_type"`
            State string  `json:"state"`
            AccountNo string  `json:"account_no"`
            AccountHolder string  `json:"account_holder"`
            BankName string  `json:"bank_name"`
            City string  `json:"city"`
            Pincode float64  `json:"pincode"`
            IfscCode string  `json:"ifsc_code"`
            BranchName string  `json:"branch_name"`
            Country string  `json:"country"`
         
    }
    
    // PayoutRequest used by Payment
    type PayoutRequest struct {

        
            TransferType string  `json:"transfer_type"`
            Users map[string]interface{}  `json:"users"`
            IsActive bool  `json:"is_active"`
            Aggregator string  `json:"aggregator"`
            UniqueExternalID string  `json:"unique_external_id"`
            BankDetails PayoutBankDetails  `json:"bank_details"`
         
    }
    
    // PayoutResponse used by Payment
    type PayoutResponse struct {

        
            TransferType string  `json:"transfer_type"`
            PaymentStatus string  `json:"payment_status"`
            UniqueTransferNo string  `json:"unique_transfer_no"`
            Users map[string]interface{}  `json:"users"`
            IsActive bool  `json:"is_active"`
            Success bool  `json:"success"`
            Aggregator string  `json:"aggregator"`
            Created bool  `json:"created"`
            BankDetails map[string]interface{}  `json:"bank_details"`
            Payouts map[string]interface{}  `json:"payouts"`
         
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
            IsDefault bool  `json:"is_default"`
            UniqueExternalID string  `json:"unique_external_id"`
         
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
            Config map[string]interface{}  `json:"config"`
            Success bool  `json:"success"`
         
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
            IsVerifiedFlag bool  `json:"is_verified_flag"`
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // NotFoundResourceError used by Payment
    type NotFoundResourceError struct {

        
            Code string  `json:"code"`
            Success bool  `json:"success"`
            Description string  `json:"description"`
         
    }
    
    // BankDetailsForOTP used by Payment
    type BankDetailsForOTP struct {

        
            AccountNo string  `json:"account_no"`
            BankName string  `json:"bank_name"`
            AccountHolder string  `json:"account_holder"`
            IfscCode string  `json:"ifsc_code"`
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
            Success bool  `json:"success"`
            BankName string  `json:"bank_name"`
         
    }
    
    // OrderBeneficiaryDetails used by Payment
    type OrderBeneficiaryDetails struct {

        
            Comment string  `json:"comment"`
            AccountNo string  `json:"account_no"`
            ModifiedOn string  `json:"modified_on"`
            Mobile string  `json:"mobile"`
            ID float64  `json:"id"`
            BranchName string  `json:"branch_name"`
            Subtitle string  `json:"subtitle"`
            IsActive bool  `json:"is_active"`
            Email string  `json:"email"`
            CreatedOn string  `json:"created_on"`
            IfscCode string  `json:"ifsc_code"`
            Title string  `json:"title"`
            Address string  `json:"address"`
            DisplayName string  `json:"display_name"`
            BeneficiaryID string  `json:"beneficiary_id"`
            AccountHolder string  `json:"account_holder"`
            TransferMode string  `json:"transfer_mode"`
            DelightsUserName string  `json:"delights_user_name"`
            BankName string  `json:"bank_name"`
         
    }
    
    // OrderBeneficiaryResponse used by Payment
    type OrderBeneficiaryResponse struct {

        
            Beneficiaries []OrderBeneficiaryDetails  `json:"beneficiaries"`
            ShowBeneficiaryDetails bool  `json:"show_beneficiary_details"`
         
    }
    
    // MultiTenderPaymentMeta used by Payment
    type MultiTenderPaymentMeta struct {

        
            OrderID string  `json:"order_id"`
            PaymentGateway string  `json:"payment_gateway"`
            CurrentStatus string  `json:"current_status"`
            PaymentID string  `json:"payment_id"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
         
    }
    
    // MultiTenderPaymentMethod used by Payment
    type MultiTenderPaymentMethod struct {

        
            Name string  `json:"name"`
            Meta MultiTenderPaymentMeta  `json:"meta"`
            Mode string  `json:"mode"`
            Amount float64  `json:"amount"`
         
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

        
            UserID string  `json:"user_id"`
            RemainingLimit float64  `json:"remaining_limit"`
            Usages float64  `json:"usages"`
            IsActive bool  `json:"is_active"`
            Limit float64  `json:"limit"`
         
    }
    
    // GetUserCODLimitResponse used by Payment
    type GetUserCODLimitResponse struct {

        
            UserCodData CODdata  `json:"user_cod_data"`
            Success bool  `json:"success"`
         
    }
    
    // SetCODForUserRequest used by Payment
    type SetCODForUserRequest struct {

        
            MerchantUserID string  `json:"merchant_user_id"`
            IsActive bool  `json:"is_active"`
            Mobileno string  `json:"mobileno"`
         
    }
    
    // SetCODOptionResponse used by Payment
    type SetCODOptionResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // EdcModelData used by Payment
    type EdcModelData struct {

        
            AggregatorID float64  `json:"aggregator_id"`
            Models []string  `json:"models"`
            Aggregator string  `json:"aggregator"`
         
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

        
            Statistics StatisticsData  `json:"statistics"`
            Success bool  `json:"success"`
         
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

        
            ApplicationID string  `json:"application_id"`
            EdcDeviceSerialNo string  `json:"edc_device_serial_no"`
            MerchantStorePosCode string  `json:"merchant_store_pos_code"`
            EdcModel string  `json:"edc_model"`
            AggregatorID float64  `json:"aggregator_id"`
            DeviceTag string  `json:"device_tag"`
            AggregatorName string  `json:"aggregator_name"`
            TerminalUniqueIdentifier string  `json:"terminal_unique_identifier"`
            IsActive bool  `json:"is_active"`
            StoreID float64  `json:"store_id"`
            TerminalSerialNo string  `json:"terminal_serial_no"`
         
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

        
            EdcDeviceSerialNo string  `json:"edc_device_serial_no"`
            MerchantStorePosCode string  `json:"merchant_store_pos_code"`
            EdcModel string  `json:"edc_model"`
            DeviceTag string  `json:"device_tag"`
            AggregatorID float64  `json:"aggregator_id"`
            IsActive bool  `json:"is_active"`
            StoreID float64  `json:"store_id"`
         
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
            Success bool  `json:"success"`
            Items []EdcDevice  `json:"items"`
         
    }
    
    // PaymentInitializationRequest used by Payment
    type PaymentInitializationRequest struct {

        
            Vpa string  `json:"vpa"`
            OrderID string  `json:"order_id"`
            Timeout float64  `json:"timeout"`
            Method string  `json:"method"`
            RazorpayPaymentID string  `json:"razorpay_payment_id"`
            Contact string  `json:"contact"`
            DeviceID string  `json:"device_id"`
            Currency string  `json:"currency"`
            MerchantOrderID string  `json:"merchant_order_id"`
            Aggregator string  `json:"aggregator"`
            Email string  `json:"email"`
            CustomerID string  `json:"customer_id"`
            Amount float64  `json:"amount"`
         
    }
    
    // PaymentInitializationResponse used by Payment
    type PaymentInitializationResponse struct {

        
            Vpa string  `json:"vpa"`
            Timeout float64  `json:"timeout"`
            Method string  `json:"method"`
            AggregatorOrderID string  `json:"aggregator_order_id"`
            RazorpayPaymentID string  `json:"razorpay_payment_id"`
            PollingURL string  `json:"polling_url"`
            VirtualID string  `json:"virtual_id"`
            BqrImage string  `json:"bqr_image"`
            Currency string  `json:"currency"`
            MerchantOrderID string  `json:"merchant_order_id"`
            DeviceID string  `json:"device_id"`
            Status string  `json:"status"`
            Aggregator string  `json:"aggregator"`
            Success bool  `json:"success"`
            UpiPollURL string  `json:"upi_poll_url"`
            CustomerID string  `json:"customer_id"`
            Amount float64  `json:"amount"`
         
    }
    
    // PaymentStatusUpdateRequest used by Payment
    type PaymentStatusUpdateRequest struct {

        
            Vpa string  `json:"vpa"`
            OrderID string  `json:"order_id"`
            MerchantTransactionID string  `json:"merchant_transaction_id"`
            Method string  `json:"method"`
            Contact string  `json:"contact"`
            DeviceID string  `json:"device_id"`
            Currency string  `json:"currency"`
            MerchantOrderID string  `json:"merchant_order_id"`
            Status string  `json:"status"`
            Aggregator string  `json:"aggregator"`
            Email string  `json:"email"`
            CustomerID string  `json:"customer_id"`
            Amount float64  `json:"amount"`
         
    }
    
    // PaymentStatusUpdateResponse used by Payment
    type PaymentStatusUpdateResponse struct {

        
            Retry bool  `json:"retry"`
            AggregatorName string  `json:"aggregator_name"`
            Status string  `json:"status"`
            Success bool  `json:"success"`
            RedirectURL string  `json:"redirect_url"`
         
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

        
            Data LinkStatus  `json:"data"`
            Success bool  `json:"success"`
         
    }
    
    // PaymentStatusBulkHandlerRequest used by Payment
    type PaymentStatusBulkHandlerRequest struct {

        
            MerchantOrderID []string  `json:"merchant_order_id"`
         
    }
    
    // PaymentObjectListSerializer used by Payment
    type PaymentObjectListSerializer struct {

        
            ApplicationID string  `json:"application_id"`
            PaymentMode string  `json:"payment_mode"`
            PaymentModeIdentifier string  `json:"payment_mode_identifier"`
            CollectedBy string  `json:"collected_by"`
            ModifiedOn string  `json:"modified_on"`
            UserObject map[string]interface{}  `json:"user_object"`
            PaymentGateway string  `json:"payment_gateway"`
            RefundedBy string  `json:"refunded_by"`
            CurrentStatus string  `json:"current_status"`
            AmountInPaisa string  `json:"amount_in_paisa"`
            Currency string  `json:"currency"`
            PaymentID string  `json:"payment_id"`
            CompanyID string  `json:"company_id"`
            AggregatorPaymentObject map[string]interface{}  `json:"aggregator_payment_object"`
            CreatedOn string  `json:"created_on"`
            ID string  `json:"id"`
            RefundObject map[string]interface{}  `json:"refund_object"`
            AllStatus []string  `json:"all_status"`
         
    }
    
    // PaymentStatusObject used by Payment
    type PaymentStatusObject struct {

        
            MerchantOrderID string  `json:"merchant_order_id"`
            PaymentObjectList []PaymentObjectListSerializer  `json:"payment_object_list"`
         
    }
    
    // PaymentStatusBulkHandlerResponse used by Payment
    type PaymentStatusBulkHandlerResponse struct {

        
            Error string  `json:"error"`
            Count float64  `json:"count"`
            Data []PaymentStatusObject  `json:"data"`
            Status float64  `json:"status"`
            Success string  `json:"success"`
         
    }
    
    // GetOauthUrlResponse used by Payment
    type GetOauthUrlResponse struct {

        
            URL string  `json:"url"`
            Success bool  `json:"success"`
         
    }
    
    // RevokeOAuthToken used by Payment
    type RevokeOAuthToken struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // RepaymentRequestDetails used by Payment
    type RepaymentRequestDetails struct {

        
            PaymentMode string  `json:"payment_mode"`
            PaymentModeIdentifier string  `json:"payment_mode_identifier"`
            FwdShipmentID string  `json:"fwd_shipment_id"`
            AggregatorOrderID string  `json:"aggregator_order_id"`
            CurrentStatus string  `json:"current_status"`
            AggregatorTransactionID string  `json:"aggregator_transaction_id"`
            OutstandingDetailsID float64  `json:"outstanding_details_id"`
            MerchantOrderID string  `json:"merchant_order_id"`
            Aggregator string  `json:"aggregator"`
            Amount float64  `json:"amount"`
         
    }
    
    // RepaymentDetailsSerialiserPayAll used by Payment
    type RepaymentDetailsSerialiserPayAll struct {

        
            ShipmentDetails []RepaymentRequestDetails  `json:"shipment_details"`
            AggregatorOrderID string  `json:"aggregator_order_id"`
            AggregatorTransactionID string  `json:"aggregator_transaction_id"`
            ExtensionOrderID string  `json:"extension_order_id"`
            TotalAmount float64  `json:"total_amount"`
         
    }
    
    // RepaymentResponse used by Payment
    type RepaymentResponse struct {

        
            Data map[string]interface{}  `json:"data"`
            Success bool  `json:"success"`
         
    }
    
    // MerchantOnBoardingRequest used by Payment
    type MerchantOnBoardingRequest struct {

        
            UserID string  `json:"user_id"`
            Status string  `json:"status"`
            Aggregator string  `json:"aggregator"`
            CreditLineID string  `json:"credit_line_id"`
            AppID string  `json:"app_id"`
         
    }
    
    // MerchantOnBoardingResponse used by Payment
    type MerchantOnBoardingResponse struct {

        
            Data map[string]interface{}  `json:"data"`
            Success bool  `json:"success"`
         
    }
    
    // ValidateCustomerRequest used by Payment
    type ValidateCustomerRequest struct {

        
            TransactionAmountInPaise float64  `json:"transaction_amount_in_paise"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
            OrderItems []map[string]interface{}  `json:"order_items"`
            Aggregator string  `json:"aggregator"`
            PhoneNumber string  `json:"phone_number"`
            MerchantParams map[string]interface{}  `json:"merchant_params"`
            Payload string  `json:"payload"`
         
    }
    
    // ValidateCustomerResponse used by Payment
    type ValidateCustomerResponse struct {

        
            Data map[string]interface{}  `json:"data"`
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // GetPaymentLinkResponse used by Payment
    type GetPaymentLinkResponse struct {

        
            ExternalOrderID string  `json:"external_order_id"`
            StatusCode float64  `json:"status_code"`
            PaymentLinkCurrentStatus string  `json:"payment_link_current_status"`
            MerchantName string  `json:"merchant_name"`
            PollingTimeout float64  `json:"polling_timeout"`
            Message string  `json:"message"`
            PaymentLinkURL string  `json:"payment_link_url"`
            Success bool  `json:"success"`
            Amount float64  `json:"amount"`
         
    }
    
    // ErrorDescription used by Payment
    type ErrorDescription struct {

        
            Msg string  `json:"msg"`
            InvalidID bool  `json:"invalid_id"`
            MerchantName string  `json:"merchant_name"`
            MerchantOrderID string  `json:"merchant_order_id"`
            Cancelled bool  `json:"cancelled"`
            PaymentTransactionID string  `json:"payment_transaction_id"`
            Expired bool  `json:"expired"`
            Amount float64  `json:"amount"`
         
    }
    
    // ErrorResponse used by Payment
    type ErrorResponse struct {

        
            Error ErrorDescription  `json:"error"`
            Success bool  `json:"success"`
            StatusCode float64  `json:"status_code"`
            Message string  `json:"message"`
         
    }
    
    // CreatePaymentLinkMeta used by Payment
    type CreatePaymentLinkMeta struct {

        
            CartID string  `json:"cart_id"`
            Pincode string  `json:"pincode"`
            CheckoutMode string  `json:"checkout_mode"`
            AssignCardID string  `json:"assign_card_id"`
            Amount string  `json:"amount"`
         
    }
    
    // CreatePaymentLinkRequest used by Payment
    type CreatePaymentLinkRequest struct {

        
            ExternalOrderID string  `json:"external_order_id"`
            MobileNumber string  `json:"mobile_number"`
            Meta CreatePaymentLinkMeta  `json:"meta"`
            Email string  `json:"email"`
            Amount float64  `json:"amount"`
            Description string  `json:"description"`
         
    }
    
    // CreatePaymentLinkResponse used by Payment
    type CreatePaymentLinkResponse struct {

        
            StatusCode float64  `json:"status_code"`
            PollingTimeout float64  `json:"polling_timeout"`
            Message string  `json:"message"`
            PaymentLinkURL string  `json:"payment_link_url"`
            PaymentLinkID string  `json:"payment_link_id"`
            Success bool  `json:"success"`
         
    }
    
    // PollingPaymentLinkResponse used by Payment
    type PollingPaymentLinkResponse struct {

        
            StatusCode float64  `json:"status_code"`
            OrderID string  `json:"order_id"`
            Amount float64  `json:"amount"`
            Message string  `json:"message"`
            HttpStatus float64  `json:"http_status"`
            AggregatorName string  `json:"aggregator_name"`
            PaymentLinkID string  `json:"payment_link_id"`
            Status string  `json:"status"`
            Success bool  `json:"success"`
            RedirectURL string  `json:"redirect_url"`
         
    }
    
    // CancelOrResendPaymentLinkRequest used by Payment
    type CancelOrResendPaymentLinkRequest struct {

        
            PaymentLinkID string  `json:"payment_link_id"`
         
    }
    
    // ResendPaymentLinkResponse used by Payment
    type ResendPaymentLinkResponse struct {

        
            StatusCode float64  `json:"status_code"`
            Message string  `json:"message"`
            PollingTimeout float64  `json:"polling_timeout"`
            Success bool  `json:"success"`
         
    }
    
    // CancelPaymentLinkResponse used by Payment
    type CancelPaymentLinkResponse struct {

        
            StatusCode float64  `json:"status_code"`
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // Code used by Payment
    type Code struct {

        
            Name string  `json:"name"`
            MerchantCode string  `json:"merchant_code"`
            Code string  `json:"code"`
         
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
    
    // Prices used by Order
    type Prices struct {

        
            TransferPrice float64  `json:"transfer_price"`
            PromotionEffectiveDiscount float64  `json:"promotion_effective_discount"`
            DeliveryCharge float64  `json:"delivery_charge"`
            GiftPrice float64  `json:"gift_price"`
            CouponValue float64  `json:"coupon_value"`
            RefundAmount float64  `json:"refund_amount"`
            ValueOfGood float64  `json:"value_of_good"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            RefundCredit float64  `json:"refund_credit"`
            AmountPaidRoundoff float64  `json:"amount_paid_roundoff"`
            CodCharges float64  `json:"cod_charges"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            FyndCredits float64  `json:"fynd_credits"`
            PriceMarked float64  `json:"price_marked"`
            AmountPaid float64  `json:"amount_paid"`
            Cashback float64  `json:"cashback"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            PmPriceSplit float64  `json:"pm_price_split"`
            CashbackApplied float64  `json:"cashback_applied"`
            Discount float64  `json:"discount"`
            PriceEffective float64  `json:"price_effective"`
         
    }
    
    // UserDataInfo used by Order
    type UserDataInfo struct {

        
            Mobile string  `json:"mobile"`
            IsAnonymousUser bool  `json:"is_anonymous_user"`
            AvisUserID string  `json:"avis_user_id"`
            LastName string  `json:"last_name"`
            Email string  `json:"email"`
            Name string  `json:"name"`
            FirstName string  `json:"first_name"`
            UID float64  `json:"uid"`
            Gender string  `json:"gender"`
         
    }
    
    // ShipmentItemFulFillingStore used by Order
    type ShipmentItemFulFillingStore struct {

        
            Address string  `json:"address"`
            Meta map[string]interface{}  `json:"meta"`
            Code string  `json:"code"`
            City string  `json:"city"`
            Phone string  `json:"phone"`
            ID float64  `json:"id"`
            LocationType string  `json:"location_type"`
            State string  `json:"state"`
            Name string  `json:"name"`
            Pincode string  `json:"pincode"`
            BrandStoreTags string  `json:"brand_store_tags"`
            StoreEmail string  `json:"store_email"`
         
    }
    
    // LockData used by Order
    type LockData struct {

        
            Mto bool  `json:"mto"`
            Locked bool  `json:"locked"`
            LockMessage string  `json:"lock_message"`
         
    }
    
    // Formatted used by Order
    type Formatted struct {

        
            Max string  `json:"max"`
            Min string  `json:"min"`
         
    }
    
    // ShipmentTags used by Order
    type ShipmentTags struct {

        
            Slug string  `json:"slug"`
            EntityType string  `json:"entity_type"`
            DisplayText string  `json:"display_text"`
         
    }
    
    // ShipmentTimeStamp used by Order
    type ShipmentTimeStamp struct {

        
            Max float64  `json:"max"`
            Min float64  `json:"min"`
         
    }
    
    // ShipmentItemMeta used by Order
    type ShipmentItemMeta struct {

        
            ShipmentWeight float64  `json:"shipment_weight"`
            ShipmentVolumetricWeight float64  `json:"shipment_volumetric_weight"`
            ShipmentChargeableWeight float64  `json:"shipment_chargeable_weight"`
            FulfilmentPriorityText string  `json:"fulfilment_priority_text"`
            Sla float64  `json:"sla"`
            ParentDpID string  `json:"parent_dp_id"`
            DpSortKey string  `json:"dp_sort_key"`
            LockData LockData  `json:"lock_data"`
            External map[string]interface{}  `json:"external"`
            PackagingName string  `json:"packaging_name"`
            ActivityComment string  `json:"activity_comment"`
            SameStoreAvailable bool  `json:"same_store_available"`
            PdfMedia []map[string]interface{}  `json:"pdf_media"`
            ExistingDpList []string  `json:"existing_dp_list"`
            Formatted Formatted  `json:"formatted"`
            BagWeight map[string]interface{}  `json:"bag_weight"`
            DpOptions map[string]interface{}  `json:"dp_options"`
            Weight float64  `json:"weight"`
            OrderType string  `json:"order_type"`
            AssignDpFromSb bool  `json:"assign_dp_from_sb"`
            DebugInfo map[string]interface{}  `json:"debug_info"`
            EwaybillInfo map[string]interface{}  `json:"ewaybill_info"`
            ShipmentTags []ShipmentTags  `json:"shipment_tags"`
            IsInternational bool  `json:"is_international"`
            AutoTriggerDpAssignmentAcf bool  `json:"auto_trigger_dp_assignment_acf"`
            Tags []map[string]interface{}  `json:"tags"`
            Timestamp ShipmentTimeStamp  `json:"timestamp"`
            ShippingZone string  `json:"shipping_zone"`
            StoreInvoiceUpdatedDate string  `json:"store_invoice_updated_date"`
         
    }
    
    // BagStateMapper used by Order
    type BagStateMapper struct {

        
            IsActive bool  `json:"is_active"`
            DisplayName string  `json:"display_name"`
            AppStateName string  `json:"app_state_name"`
            AppDisplayName string  `json:"app_display_name"`
            NotifyCustomer bool  `json:"notify_customer"`
            ID float64  `json:"id"`
            StateType string  `json:"state_type"`
            Name string  `json:"name"`
            JourneyType string  `json:"journey_type"`
            AppFacing bool  `json:"app_facing"`
         
    }
    
    // BagStatusHistory used by Order
    type BagStatusHistory struct {

        
            DisplayName string  `json:"display_name"`
            BagStateMapper BagStateMapper  `json:"bag_state_mapper"`
            AppDisplayName string  `json:"app_display_name"`
            ShipmentID string  `json:"shipment_id"`
            CreatedAt string  `json:"created_at"`
            KafkaSync bool  `json:"kafka_sync"`
            StateType string  `json:"state_type"`
            StoreID float64  `json:"store_id"`
            BshID float64  `json:"bsh_id"`
            Reasons []map[string]interface{}  `json:"reasons"`
            BagID float64  `json:"bag_id"`
            Forward bool  `json:"forward"`
            UpdatedAt string  `json:"updated_at"`
            DeliveryAwbNumber string  `json:"delivery_awb_number"`
            Status string  `json:"status"`
            DeliveryPartnerID float64  `json:"delivery_partner_id"`
            StateID float64  `json:"state_id"`
         
    }
    
    // Identifier used by Order
    type Identifier struct {

        
            Isbn string  `json:"isbn"`
            SkuCode string  `json:"sku_code"`
            Alu string  `json:"alu"`
            Upc string  `json:"upc"`
            Ean string  `json:"ean"`
         
    }
    
    // FinancialBreakup used by Order
    type FinancialBreakup struct {

        
            TotalUnits float64  `json:"total_units"`
            TransferPrice float64  `json:"transfer_price"`
            PromotionEffectiveDiscount float64  `json:"promotion_effective_discount"`
            Identifiers Identifier  `json:"identifiers"`
            DeliveryCharge float64  `json:"delivery_charge"`
            Size string  `json:"size"`
            CouponValue float64  `json:"coupon_value"`
            ValueOfGood float64  `json:"value_of_good"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            AddedToFyndCash bool  `json:"added_to_fynd_cash"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            RefundCredit float64  `json:"refund_credit"`
            HsnCode string  `json:"hsn_code"`
            AmountPaidRoundoff float64  `json:"amount_paid_roundoff"`
            GstTag string  `json:"gst_tag"`
            CodCharges float64  `json:"cod_charges"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            FyndCredits float64  `json:"fynd_credits"`
            PriceMarked float64  `json:"price_marked"`
            AmountPaid float64  `json:"amount_paid"`
            ItemName string  `json:"item_name"`
            Cashback float64  `json:"cashback"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            GstFee float64  `json:"gst_fee"`
            CashbackApplied float64  `json:"cashback_applied"`
            Discount float64  `json:"discount"`
            PriceEffective float64  `json:"price_effective"`
         
    }
    
    // ShipmentListingBrand used by Order
    type ShipmentListingBrand struct {

        
            Logo string  `json:"logo"`
            CreatedOn string  `json:"created_on"`
            Name string  `json:"name"`
            LogoBase64 string  `json:"logo_base64"`
         
    }
    
    // GSTDetailsData used by Order
    type GSTDetailsData struct {

        
            CgstGstFee string  `json:"cgst_gst_fee"`
            IgstGstFee string  `json:"igst_gst_fee"`
            IgstTaxPercentage float64  `json:"igst_tax_percentage"`
            IsDefaultHsnCode bool  `json:"is_default_hsn_code"`
            GstTag string  `json:"gst_tag"`
            ValueOfGood float64  `json:"value_of_good"`
            GstinCode string  `json:"gstin_code"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            SgstTaxPercentage float64  `json:"sgst_tax_percentage"`
            HsnCodeID string  `json:"hsn_code_id"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            CgstTaxPercentage float64  `json:"cgst_tax_percentage"`
            GstFee float64  `json:"gst_fee"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            SgstGstFee string  `json:"sgst_gst_fee"`
            HsnCode string  `json:"hsn_code"`
         
    }
    
    // PlatformArticleAttributes used by Order
    type PlatformArticleAttributes struct {

        
            Currency string  `json:"currency"`
         
    }
    
    // PlatformItem used by Order
    type PlatformItem struct {

        
            BrandID float64  `json:"brand_id"`
            L1Category []string  `json:"l1_category"`
            Color string  `json:"color"`
            CanReturn bool  `json:"can_return"`
            L3Category float64  `json:"l3_category"`
            DepartmentID float64  `json:"department_id"`
            Meta map[string]interface{}  `json:"meta"`
            Size string  `json:"size"`
            Brand string  `json:"brand"`
            LastUpdatedAt string  `json:"last_updated_at"`
            Images []string  `json:"images"`
            BranchURL string  `json:"branch_url"`
            Code string  `json:"code"`
            SlugKey string  `json:"slug_key"`
            ID float64  `json:"id"`
            L3CategoryName string  `json:"l3_category_name"`
            L2Category []string  `json:"l2_category"`
            Name string  `json:"name"`
            Image []string  `json:"image"`
            CanCancel bool  `json:"can_cancel"`
            Attributes PlatformArticleAttributes  `json:"attributes"`
         
    }
    
    // BagReturnableCancelableStatus used by Order
    type BagReturnableCancelableStatus struct {

        
            IsActive bool  `json:"is_active"`
            IsCustomerReturnAllowed bool  `json:"is_customer_return_allowed"`
            IsReturnable bool  `json:"is_returnable"`
            CanBeCancelled bool  `json:"can_be_cancelled"`
            EnableTracking bool  `json:"enable_tracking"`
         
    }
    
    // Dates used by Order
    type Dates struct {

        
            DeliveryDate string  `json:"delivery_date"`
            OrderCreated string  `json:"order_created"`
         
    }
    
    // ReturnConfig used by Order
    type ReturnConfig struct {

        
            Unit string  `json:"unit"`
            Time float64  `json:"time"`
            Returnable bool  `json:"returnable"`
         
    }
    
    // Weight used by Order
    type Weight struct {

        
            Unit string  `json:"unit"`
            IsDefault bool  `json:"is_default"`
            Shipping float64  `json:"shipping"`
         
    }
    
    // Dimensions used by Order
    type Dimensions struct {

        
            Length float64  `json:"length"`
            IsDefault bool  `json:"is_default"`
            Width float64  `json:"width"`
            Height float64  `json:"height"`
            Unit string  `json:"unit"`
         
    }
    
    // Article used by Order
    type Article struct {

        
            Size string  `json:"size"`
            ASet map[string]interface{}  `json:"a_set"`
            RawMeta string  `json:"raw_meta"`
            Code string  `json:"code"`
            ReturnConfig ReturnConfig  `json:"return_config"`
            ChildDetails map[string]interface{}  `json:"child_details"`
            EspModified bool  `json:"esp_modified"`
            Weight Weight  `json:"weight"`
            Currency map[string]interface{}  `json:"currency"`
            SellerIdentifier string  `json:"seller_identifier"`
            Identifiers map[string]interface{}  `json:"identifiers"`
            ID string  `json:"_id"`
            Dimensions Dimensions  `json:"dimensions"`
            IsSet bool  `json:"is_set"`
            UID string  `json:"uid"`
         
    }
    
    // ReplacementDetails used by Order
    type ReplacementDetails struct {

        
            OriginalAffiliateOrderID string  `json:"original_affiliate_order_id"`
            ReplacementType string  `json:"replacement_type"`
         
    }
    
    // AffiliateMeta used by Order
    type AffiliateMeta struct {

        
            OrderItemID string  `json:"order_item_id"`
            DueDate string  `json:"due_date"`
            CouponCode string  `json:"coupon_code"`
            ChannelOrderID string  `json:"channel_order_id"`
            IsPriority bool  `json:"is_priority"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            BoxType string  `json:"box_type"`
            EmployeeDiscount float64  `json:"employee_discount"`
            ReplacementDetails ReplacementDetails  `json:"replacement_details"`
            MarketplaceInvoiceID string  `json:"marketplace_invoice_id"`
            SizeLevelTotalQty float64  `json:"size_level_total_qty"`
            ChannelShipmentID string  `json:"channel_shipment_id"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // AffiliateBagDetails used by Order
    type AffiliateBagDetails struct {

        
            AffiliateMeta AffiliateMeta  `json:"affiliate_meta"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            EmployeeDiscount float64  `json:"employee_discount"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
         
    }
    
    // BagUnit used by Order
    type BagUnit struct {

        
            BagStatus []BagStatusHistory  `json:"bag_status"`
            Prices Prices  `json:"prices"`
            FinancialBreakup []FinancialBreakup  `json:"financial_breakup"`
            LineNumber float64  `json:"line_number"`
            CanReturn bool  `json:"can_return"`
            Meta map[string]interface{}  `json:"meta"`
            Size string  `json:"size"`
            CurrentOperationalStatus BagStatusHistory  `json:"current_operational_status"`
            Brand ShipmentListingBrand  `json:"brand"`
            EntityType string  `json:"entity_type"`
            BagType string  `json:"bag_type"`
            DisplayName string  `json:"display_name"`
            CurrentStatus BagStatusHistory  `json:"current_status"`
            Gst GSTDetailsData  `json:"gst"`
            Reasons []map[string]interface{}  `json:"reasons"`
            BagID float64  `json:"bag_id"`
            Item PlatformItem  `json:"item"`
            BagExpiryDate string  `json:"bag_expiry_date"`
            Status BagReturnableCancelableStatus  `json:"status"`
            CanCancel bool  `json:"can_cancel"`
            Dates Dates  `json:"dates"`
            Article Article  `json:"article"`
            ProductQuantity float64  `json:"product_quantity"`
            AffiliateBagDetails AffiliateBagDetails  `json:"affiliate_bag_details"`
         
    }
    
    // PlatformDeliveryAddress used by Order
    type PlatformDeliveryAddress struct {

        
            AddressCategory string  `json:"address_category"`
            City string  `json:"city"`
            AddressType string  `json:"address_type"`
            Area string  `json:"area"`
            Phone string  `json:"phone"`
            Longitude float64  `json:"longitude"`
            CreatedAt string  `json:"created_at"`
            Version string  `json:"version"`
            Latitude float64  `json:"latitude"`
            State string  `json:"state"`
            ContactPerson string  `json:"contact_person"`
            Email string  `json:"email"`
            Address2 string  `json:"address2"`
            Landmark string  `json:"landmark"`
            UpdatedAt string  `json:"updated_at"`
            Pincode string  `json:"pincode"`
            Address1 string  `json:"address1"`
            Country string  `json:"country"`
         
    }
    
    // ShipmentListingChannel used by Order
    type ShipmentListingChannel struct {

        
            Logo string  `json:"logo"`
            Name string  `json:"name"`
            ChannelShipmentID string  `json:"channel_shipment_id"`
            IsAffiliate bool  `json:"is_affiliate"`
         
    }
    
    // ShipmentStatus used by Order
    type ShipmentStatus struct {

        
            Meta map[string]interface{}  `json:"meta"`
            CurrentShipmentStatus string  `json:"current_shipment_status"`
            CreatedAt string  `json:"created_at"`
            BagList []string  `json:"bag_list"`
            ShipmentID string  `json:"shipment_id"`
            Status string  `json:"status"`
            StatusCreatedAt string  `json:"status_created_at"`
            Title string  `json:"title"`
            ShipmentStatusID float64  `json:"shipment_status_id"`
         
    }
    
    // ShipmentItem used by Order
    type ShipmentItem struct {

        
            Prices Prices  `json:"prices"`
            User UserDataInfo  `json:"user"`
            LockStatus bool  `json:"lock_status"`
            ShipmentID string  `json:"shipment_id"`
            PaymentMode string  `json:"payment_mode"`
            TotalBags float64  `json:"total_bags"`
            FulfillingStore ShipmentItemFulFillingStore  `json:"fulfilling_store"`
            Meta ShipmentItemMeta  `json:"meta"`
            Bags []BagUnit  `json:"bags"`
            InvoiceID string  `json:"invoice_id"`
            StatusCreatedAt string  `json:"status_created_at"`
            OrderID string  `json:"order_id"`
            OrderingChannnel string  `json:"ordering_channnel"`
            PreviousShipmentID string  `json:"previous_shipment_id"`
            DisplayName string  `json:"display_name"`
            EstimatedSlaTime string  `json:"estimated_sla_time"`
            PaymentMethods map[string]interface{}  `json:"payment_methods"`
            OrderDate string  `json:"order_date"`
            DeliveryAddress PlatformDeliveryAddress  `json:"delivery_address"`
            Channel ShipmentListingChannel  `json:"channel"`
            ShipmentStatus ShipmentStatus  `json:"shipment_status"`
            ShipmentCreatedAt string  `json:"shipment_created_at"`
            CanProcess bool  `json:"can_process"`
            CustomerNote string  `json:"customer_note"`
         
    }
    
    // ShipmentInternalPlatformViewResponse used by Order
    type ShipmentInternalPlatformViewResponse struct {

        
            Page Page  `json:"page"`
            Success bool  `json:"success"`
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
    
    // DPDetailsData used by Order
    type DPDetailsData struct {

        
            TrackURL string  `json:"track_url"`
            AwbNo string  `json:"awb_no"`
            GstTag string  `json:"gst_tag"`
            ID float64  `json:"id"`
            EwayBillID string  `json:"eway_bill_id"`
            Name string  `json:"name"`
            Pincode string  `json:"pincode"`
            Country string  `json:"country"`
         
    }
    
    // TrackingList used by Order
    type TrackingList struct {

        
            IsCurrent bool  `json:"is_current"`
            Text string  `json:"text"`
            Time string  `json:"time"`
            Status string  `json:"status"`
            IsPassed bool  `json:"is_passed"`
         
    }
    
    // FulfillingStore used by Order
    type FulfillingStore struct {

        
            Address string  `json:"address"`
            Meta map[string]interface{}  `json:"meta"`
            Code string  `json:"code"`
            City string  `json:"city"`
            Phone string  `json:"phone"`
            ID float64  `json:"id"`
            ContactPerson string  `json:"contact_person"`
            FulfillmentChannel string  `json:"fulfillment_channel"`
            State string  `json:"state"`
            StoreName string  `json:"store_name"`
            Pincode string  `json:"pincode"`
            Country string  `json:"country"`
         
    }
    
    // BuyerDetails used by Order
    type BuyerDetails struct {

        
            Address string  `json:"address"`
            City string  `json:"city"`
            State string  `json:"state"`
            AjioSiteID string  `json:"ajio_site_id"`
            Gstin string  `json:"gstin"`
            Name string  `json:"name"`
            Pincode float64  `json:"pincode"`
         
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

        
            ShipmentWeight float64  `json:"shipment_weight"`
            ShipmentVolumetricWeight float64  `json:"shipment_volumetric_weight"`
            B2bBuyerDetails BuyerDetails  `json:"b2b_buyer_details"`
            FulfilmentPriorityText string  `json:"fulfilment_priority_text"`
            MarketplaceStoreID string  `json:"marketplace_store_id"`
            DpName string  `json:"dp_name"`
            ParentDpID string  `json:"parent_dp_id"`
            DpSortKey string  `json:"dp_sort_key"`
            LockData LockData  `json:"lock_data"`
            External map[string]interface{}  `json:"external"`
            PackagingName string  `json:"packaging_name"`
            Dimension Dimensions  `json:"dimension"`
            ReturnAwbNumber string  `json:"return_awb_number"`
            ReturnAffiliateShipmentID string  `json:"return_affiliate_shipment_id"`
            ReturnDetails map[string]interface{}  `json:"return_details"`
            SameStoreAvailable bool  `json:"same_store_available"`
            DueDate string  `json:"due_date"`
            AwbNumber string  `json:"awb_number"`
            Formatted Formatted  `json:"formatted"`
            BagWeight map[string]interface{}  `json:"bag_weight"`
            DpOptions map[string]interface{}  `json:"dp_options"`
            Weight float64  `json:"weight"`
            OrderType string  `json:"order_type"`
            AssignDpFromSb bool  `json:"assign_dp_from_sb"`
            DebugInfo DebugInfo  `json:"debug_info"`
            EwaybillInfo map[string]interface{}  `json:"ewaybill_info"`
            ForwardAffiliateOrderID string  `json:"forward_affiliate_order_id"`
            EinvoiceInfo EinvoiceInfo  `json:"einvoice_info"`
            ShipmentTags []ShipmentTags  `json:"shipment_tags"`
            AutoTriggerDpAssignmentAcf bool  `json:"auto_trigger_dp_assignment_acf"`
            B2cBuyerDetails map[string]interface{}  `json:"b2c_buyer_details"`
            ReturnStoreNode float64  `json:"return_store_node"`
            BoxType string  `json:"box_type"`
            ReturnAffiliateOrderID string  `json:"return_affiliate_order_id"`
            DpID string  `json:"dp_id"`
            Timestamp ShipmentTimeStamp  `json:"timestamp"`
            PoNumber string  `json:"po_number"`
            ForwardAffiliateShipmentID string  `json:"forward_affiliate_shipment_id"`
            StoreInvoiceUpdatedDate string  `json:"store_invoice_updated_date"`
         
    }
    
    // B2BPODetails used by Order
    type B2BPODetails struct {

        
            PoTaxAmount float64  `json:"po_tax_amount"`
            ItemBasePrice float64  `json:"item_base_price"`
            PoLineAmount float64  `json:"po_line_amount"`
            DocketNumber string  `json:"docket_number"`
            TotalGstPercentage float64  `json:"total_gst_percentage"`
            PartialCanRet bool  `json:"partial_can_ret"`
         
    }
    
    // GiftCard used by Order
    type GiftCard struct {

        
            GiftMessage string  `json:"gift_message"`
            GiftPrice float64  `json:"gift_price"`
            IsGiftApplied bool  `json:"is_gift_applied"`
            DisplayText string  `json:"display_text"`
         
    }
    
    // BagMeta used by Order
    type BagMeta struct {

        
            B2bPoDetails B2BPODetails  `json:"b2b_po_details"`
            CustomJson map[string]interface{}  `json:"custom_json"`
            GiftCard GiftCard  `json:"gift_card"`
            DocketNumber string  `json:"docket_number"`
            PartialCanRet bool  `json:"partial_can_ret"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            CustomMessage string  `json:"custom_message"`
            GroupID string  `json:"group_id"`
         
    }
    
    // OrderBrandName used by Order
    type OrderBrandName struct {

        
            ModifiedOn string  `json:"modified_on"`
            ID float64  `json:"id"`
            BrandName string  `json:"brand_name"`
            CreatedOn string  `json:"created_on"`
            Company float64  `json:"company"`
            Logo string  `json:"logo"`
         
    }
    
    // BagGST used by Order
    type BagGST struct {

        
            CgstGstFee string  `json:"cgst_gst_fee"`
            IgstGstFee string  `json:"igst_gst_fee"`
            IsDefaultHsnCode bool  `json:"is_default_hsn_code"`
            IgstTaxPercentage float64  `json:"igst_tax_percentage"`
            GstTag string  `json:"gst_tag"`
            ValueOfGood float64  `json:"value_of_good"`
            GstinCode string  `json:"gstin_code"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            SgstTaxPercentage float64  `json:"sgst_tax_percentage"`
            HsnCodeID string  `json:"hsn_code_id"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            CgstTaxPercentage float64  `json:"cgst_tax_percentage"`
            GstFee float64  `json:"gst_fee"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            SgstGstFee string  `json:"sgst_gst_fee"`
            HsnCode string  `json:"hsn_code"`
         
    }
    
    // CurrentStatus used by Order
    type CurrentStatus struct {

        
            BagStateMapper BagStateMapper  `json:"bag_state_mapper"`
            ID float64  `json:"id"`
            CreatedAt string  `json:"created_at"`
            Status string  `json:"status"`
            ShipmentID string  `json:"shipment_id"`
            StateType string  `json:"state_type"`
            KafkaSync bool  `json:"kafka_sync"`
            StoreID float64  `json:"store_id"`
            BagID float64  `json:"bag_id"`
            UpdatedAt float64  `json:"updated_at"`
            DeliveryAwbNumber string  `json:"delivery_awb_number"`
            DeliveryPartnerID float64  `json:"delivery_partner_id"`
            StateID float64  `json:"state_id"`
         
    }
    
    // BagPaymentMethods used by Order
    type BagPaymentMethods struct {

        
            Amount float64  `json:"amount"`
            Mode string  `json:"mode"`
         
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

        
            BuyRules []BuyRules  `json:"buy_rules"`
            PromotionType string  `json:"promotion_type"`
            PromoID string  `json:"promo_id"`
            DiscountRules []DiscountRules  `json:"discount_rules"`
            PromotionName string  `json:"promotion_name"`
            ArticleQuantity float64  `json:"article_quantity"`
            Amount float64  `json:"amount"`
            MrpPromotion bool  `json:"mrp_promotion"`
         
    }
    
    // ReturnConfig1 used by Order
    type ReturnConfig1 struct {

        
            Unit string  `json:"unit"`
            Time float64  `json:"time"`
            Returnable bool  `json:"returnable"`
         
    }
    
    // OrderBagArticle used by Order
    type OrderBagArticle struct {

        
            Size string  `json:"size"`
            ReturnConfig ReturnConfig1  `json:"return_config"`
            UID string  `json:"uid"`
            Identifiers map[string]interface{}  `json:"identifiers"`
         
    }
    
    // BagConfigs used by Order
    type BagConfigs struct {

        
            IsCustomerReturnAllowed bool  `json:"is_customer_return_allowed"`
            IsActive bool  `json:"is_active"`
            IsReturnable bool  `json:"is_returnable"`
            CanBeCancelled bool  `json:"can_be_cancelled"`
            AllowForceReturn bool  `json:"allow_force_return"`
            EnableTracking bool  `json:"enable_tracking"`
         
    }
    
    // AffiliateBagsDetails used by Order
    type AffiliateBagsDetails struct {

        
            CouponCode string  `json:"coupon_code"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
         
    }
    
    // OrderBags used by Order
    type OrderBags struct {

        
            Prices Prices  `json:"prices"`
            FinancialBreakup FinancialBreakup  `json:"financial_breakup"`
            LineNumber float64  `json:"line_number"`
            CanReturn bool  `json:"can_return"`
            Meta BagMeta  `json:"meta"`
            Brand OrderBrandName  `json:"brand"`
            EntityType string  `json:"entity_type"`
            SellerIdentifier string  `json:"seller_identifier"`
            Identifier string  `json:"identifier"`
            GstDetails BagGST  `json:"gst_details"`
            DisplayName string  `json:"display_name"`
            CurrentStatus CurrentStatus  `json:"current_status"`
            PaymentMethods []BagPaymentMethods  `json:"payment_methods"`
            BagID float64  `json:"bag_id"`
            AppliedPromos []AppliedPromos  `json:"applied_promos"`
            Item PlatformItem  `json:"item"`
            DeliveryAddress PlatformDeliveryAddress  `json:"delivery_address"`
            GroupID string  `json:"group_id"`
            ParentPromoBags map[string]interface{}  `json:"parent_promo_bags"`
            Quantity float64  `json:"quantity"`
            CanCancel bool  `json:"can_cancel"`
            Article OrderBagArticle  `json:"article"`
            BagConfigs BagConfigs  `json:"bag_configs"`
            AffiliateBagDetails AffiliateBagsDetails  `json:"affiliate_bag_details"`
         
    }
    
    // OrderingStoreDetails used by Order
    type OrderingStoreDetails struct {

        
            Address string  `json:"address"`
            Meta map[string]interface{}  `json:"meta"`
            Code string  `json:"code"`
            City string  `json:"city"`
            Phone string  `json:"phone"`
            ID float64  `json:"id"`
            ContactPerson string  `json:"contact_person"`
            State string  `json:"state"`
            StoreName string  `json:"store_name"`
            Pincode string  `json:"pincode"`
            Country string  `json:"country"`
         
    }
    
    // UserDetailsData used by Order
    type UserDetailsData struct {

        
            Address string  `json:"address"`
            City string  `json:"city"`
            AddressType string  `json:"address_type"`
            Area string  `json:"area"`
            Phone string  `json:"phone"`
            State string  `json:"state"`
            Email string  `json:"email"`
            Name string  `json:"name"`
            Landmark string  `json:"landmark"`
            Pincode string  `json:"pincode"`
            Address1 string  `json:"address1"`
            Country string  `json:"country"`
         
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

        
            CompanyContact ContactDetails  `json:"company_contact"`
            Address map[string]interface{}  `json:"address"`
            CompanyGst string  `json:"company_gst"`
            CompanyID float64  `json:"company_id"`
            CompanyName string  `json:"company_name"`
            CompanyCin string  `json:"company_cin"`
         
    }
    
    // ShipmentDetails used by Order
    type ShipmentDetails struct {

        
            ActionToStatus map[string]interface{}  `json:"action_to_status"`
            LockStatus bool  `json:"lock_status"`
            LockMessage string  `json:"lock_message"`
         
    }
    
    // PDFLinks used by Order
    type PDFLinks struct {

        
            LabelPos string  `json:"label_pos"`
            CreditNoteURL string  `json:"credit_note_url"`
            InvoiceType string  `json:"invoice_type"`
            B2b string  `json:"b2b"`
            InvoicePos string  `json:"invoice_pos"`
            Label string  `json:"label"`
            LabelExport string  `json:"label_export"`
            PoInvoice string  `json:"po_invoice"`
            Invoice string  `json:"invoice"`
            LabelA4 string  `json:"label_a4"`
            LabelType string  `json:"label_type"`
            InvoiceExport string  `json:"invoice_export"`
            LabelA6 string  `json:"label_a6"`
            InvoiceA6 string  `json:"invoice_a6"`
            DeliveryChallanA4 string  `json:"delivery_challan_a4"`
            InvoiceA4 string  `json:"invoice_a4"`
         
    }
    
    // AffiliateDetails used by Order
    type AffiliateDetails struct {

        
            AffiliateID string  `json:"affiliate_id"`
            AffiliateMeta AffiliateMeta  `json:"affiliate_meta"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            AdID string  `json:"ad_id"`
            CompanyAffiliateTag string  `json:"company_affiliate_tag"`
            ShipmentMeta ShipmentMeta  `json:"shipment_meta"`
            PdfLinks PDFLinks  `json:"pdf_links"`
            AffiliateStoreID string  `json:"affiliate_store_id"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
         
    }
    
    // ShipmentStatusData used by Order
    type ShipmentStatusData struct {

        
            Meta map[string]interface{}  `json:"meta"`
            DisplayName string  `json:"display_name"`
            ShipmentID string  `json:"shipment_id"`
            BagList []string  `json:"bag_list"`
            CreatedAt string  `json:"created_at"`
            ID float64  `json:"id"`
            Status string  `json:"status"`
         
    }
    
    // InvoiceInfo used by Order
    type InvoiceInfo struct {

        
            LabelURL string  `json:"label_url"`
            UpdatedDate string  `json:"updated_date"`
            CreditNoteID string  `json:"credit_note_id"`
            ExternalInvoiceID string  `json:"external_invoice_id"`
            InvoiceURL string  `json:"invoice_url"`
            StoreInvoiceID string  `json:"store_invoice_id"`
         
    }
    
    // ShipmentPayments used by Order
    type ShipmentPayments struct {

        
            Logo string  `json:"logo"`
            Source string  `json:"source"`
            Mode string  `json:"mode"`
         
    }
    
    // OrderDetailsData used by Order
    type OrderDetailsData struct {

        
            AffiliateID string  `json:"affiliate_id"`
            TaxDetails map[string]interface{}  `json:"tax_details"`
            OrderingChannel string  `json:"ordering_channel"`
            Source string  `json:"source"`
            OrderValue string  `json:"order_value"`
            CodCharges string  `json:"cod_charges"`
            OrderDate string  `json:"order_date"`
            FyndOrderID string  `json:"fynd_order_id"`
            OrderingChannelLogo map[string]interface{}  `json:"ordering_channel_logo"`
         
    }
    
    // PlatformShipment used by Order
    type PlatformShipment struct {

        
            Prices Prices  `json:"prices"`
            Vertical string  `json:"vertical"`
            ShipmentImages []string  `json:"shipment_images"`
            User UserDataInfo  `json:"user"`
            TotalBags float64  `json:"total_bags"`
            ShipmentID string  `json:"shipment_id"`
            PaymentMode string  `json:"payment_mode"`
            LockStatus bool  `json:"lock_status"`
            ForwardShipmentID string  `json:"forward_shipment_id"`
            DpDetails DPDetailsData  `json:"dp_details"`
            TrackingList []TrackingList  `json:"tracking_list"`
            PlatformLogo string  `json:"platform_logo"`
            FulfillingStore FulfillingStore  `json:"fulfilling_store"`
            TotalItems float64  `json:"total_items"`
            PickedDate string  `json:"picked_date"`
            Meta ShipmentMeta  `json:"meta"`
            Bags []OrderBags  `json:"bags"`
            OrderingStore OrderingStoreDetails  `json:"ordering_store"`
            InvoiceID string  `json:"invoice_id"`
            DeliveryDetails UserDetailsData  `json:"delivery_details"`
            FulfilmentPriority float64  `json:"fulfilment_priority"`
            DeliverySlot map[string]interface{}  `json:"delivery_slot"`
            CustomMeta []map[string]interface{}  `json:"custom_meta"`
            CustomMessage string  `json:"custom_message"`
            JourneyType string  `json:"journey_type"`
            GstDetails GSTDetailsData  `json:"gst_details"`
            CompanyDetails CompanyDetails  `json:"company_details"`
            ShipmentDetails ShipmentDetails  `json:"shipment_details"`
            EnableDpTracking bool  `json:"enable_dp_tracking"`
            PriorityText string  `json:"priority_text"`
            Coupon map[string]interface{}  `json:"coupon"`
            CanUpdateDimension bool  `json:"can_update_dimension"`
            PackagingType string  `json:"packaging_type"`
            AffiliateDetails AffiliateDetails  `json:"affiliate_details"`
            EstimatedSlaTime string  `json:"estimated_sla_time"`
            PaymentMethods map[string]interface{}  `json:"payment_methods"`
            IsDpAssignEnabled bool  `json:"is_dp_assign_enabled"`
            Status ShipmentStatusData  `json:"status"`
            BillingDetails UserDetailsData  `json:"billing_details"`
            UserAgent string  `json:"user_agent"`
            ShipmentStatus string  `json:"shipment_status"`
            Invoice InvoiceInfo  `json:"invoice"`
            IsSelfShip bool  `json:"is_self_ship"`
            ShipmentCreatedAt string  `json:"shipment_created_at"`
            OperationalStatus string  `json:"operational_status"`
            BagStatusHistory []BagStatusHistory  `json:"bag_status_history"`
            ShipmentQuantity float64  `json:"shipment_quantity"`
            Payments ShipmentPayments  `json:"payments"`
            DpAssignment bool  `json:"dp_assignment"`
            Order OrderDetailsData  `json:"order"`
         
    }
    
    // ShipmentInfoResponse used by Order
    type ShipmentInfoResponse struct {

        
            Success bool  `json:"success"`
            Shipments []PlatformShipment  `json:"shipments"`
            Message string  `json:"message"`
         
    }
    
    // AssetByShipment used by Order
    type AssetByShipment struct {

        
            PresignedUrls map[string]string  `json:"presigned_urls"`
            ExpiresIn string  `json:"expires_in"`
            ShipmentID string  `json:"shipment_id"`
            PresignedType string  `json:"presigned_type"`
            Success bool  `json:"success"`
         
    }
    
    // ResponseGetAssetShipment used by Order
    type ResponseGetAssetShipment struct {

        
            Success bool  `json:"success"`
            Data []AssetByShipment  `json:"data"`
            Message string  `json:"message"`
         
    }
    
    // PlatformUserDetails used by Order
    type PlatformUserDetails struct {

        
            PlatformUserLastName string  `json:"platform_user_last_name"`
            PlatformUserFirstName string  `json:"platform_user_first_name"`
            PlatformUserID string  `json:"platform_user_id"`
            PlatformUserEmployeeCode string  `json:"platform_user_employee_code"`
         
    }
    
    // TransactionData used by Order
    type TransactionData struct {

        
            PaymentID string  `json:"payment_id"`
            Status string  `json:"status"`
            TransactionID string  `json:"transaction_id"`
            TerminalID string  `json:"terminal_id"`
            Entity string  `json:"entity"`
            Currency string  `json:"currency"`
            UniqueReferenceNumber string  `json:"unique_reference_number"`
            AmountPaid float64  `json:"amount_paid"`
         
    }
    
    // BillingStaffDetails used by Order
    type BillingStaffDetails struct {

        
            StaffID float64  `json:"staff_id"`
            User string  `json:"user"`
            LastName string  `json:"last_name"`
            EmployeeCode string  `json:"employee_code"`
            FirstName string  `json:"first_name"`
         
    }
    
    // OrderMeta used by Order
    type OrderMeta struct {

        
            OrderChildEntities []string  `json:"order_child_entities"`
            MongoCartID float64  `json:"mongo_cart_id"`
            PlatformUserDetails PlatformUserDetails  `json:"platform_user_details"`
            EmployeeID float64  `json:"employee_id"`
            TransactionData TransactionData  `json:"transaction_data"`
            OrderingStore float64  `json:"ordering_store"`
            CartID float64  `json:"cart_id"`
            CompanyLogo string  `json:"company_logo"`
            Files []map[string]interface{}  `json:"files"`
            Comment string  `json:"comment"`
            Staff map[string]interface{}  `json:"staff"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            OrderType string  `json:"order_type"`
            OrderPlatform string  `json:"order_platform"`
            BillingStaffDetails BillingStaffDetails  `json:"billing_staff_details"`
            OrderTags []map[string]interface{}  `json:"order_tags"`
            CurrencySymbol string  `json:"currency_symbol"`
            PaymentType string  `json:"payment_type"`
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
            Prices Prices  `json:"prices"`
            TaxDetails TaxDetails  `json:"tax_details"`
            PaymentMethods map[string]interface{}  `json:"payment_methods"`
            OrderDate string  `json:"order_date"`
            FyndOrderID string  `json:"fynd_order_id"`
         
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
            Value string  `json:"value"`
            Text string  `json:"text"`
            TotalItems float64  `json:"total_items"`
         
    }
    
    // SuperLane used by Order
    type SuperLane struct {

        
            Text string  `json:"text"`
            TotalItems float64  `json:"total_items"`
            Value string  `json:"value"`
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

        
            Channel PlatformChannel  `json:"channel"`
            Meta map[string]interface{}  `json:"meta"`
            UserInfo UserDataInfo  `json:"user_info"`
            TotalOrderValue float64  `json:"total_order_value"`
            OrderCreatedTime string  `json:"order_created_time"`
            OrderValue float64  `json:"order_value"`
            PaymentMode string  `json:"payment_mode"`
            Shipments []PlatformShipment  `json:"shipments"`
            BreakupValues []PlatformBreakupValues  `json:"breakup_values"`
            OrderID string  `json:"order_id"`
         
    }
    
    // OrderListingResponse used by Order
    type OrderListingResponse struct {

        
            Page Page  `json:"page"`
            TotalCount float64  `json:"total_count"`
            Success bool  `json:"success"`
            Items []PlatformOrderItems  `json:"items"`
            Message string  `json:"message"`
            Lane string  `json:"lane"`
         
    }
    
    // Options used by Order
    type Options struct {

        
            Text string  `json:"text"`
            Value float64  `json:"value"`
         
    }
    
    // MetricsCount used by Order
    type MetricsCount struct {

        
            Text string  `json:"text"`
            Key string  `json:"key"`
            Value float64  `json:"value"`
            Options []Options  `json:"options"`
         
    }
    
    // MetricCountResponse used by Order
    type MetricCountResponse struct {

        
            Items []MetricsCount  `json:"items"`
         
    }
    
    // PlatformTrack used by Order
    type PlatformTrack struct {

        
            Meta map[string]interface{}  `json:"meta"`
            UpdatedTime string  `json:"updated_time"`
            Reason string  `json:"reason"`
            Awb string  `json:"awb"`
            ShipmentType string  `json:"shipment_type"`
            LastLocationRecievedAt string  `json:"last_location_recieved_at"`
            AccountName string  `json:"account_name"`
            UpdatedAt string  `json:"updated_at"`
            Status string  `json:"status"`
            RawStatus string  `json:"raw_status"`
         
    }
    
    // PlatformShipmentTrack used by Order
    type PlatformShipmentTrack struct {

        
            Meta map[string]interface{}  `json:"meta"`
            Results []PlatformTrack  `json:"results"`
         
    }
    
    // FilterInfoOption used by Order
    type FilterInfoOption struct {

        
            PlaceholderText string  `json:"placeholder_text"`
            ShowUI bool  `json:"show_ui"`
            Name string  `json:"name"`
            Value string  `json:"value"`
            Text string  `json:"text"`
            MinSearchSize float64  `json:"min_search_size"`
         
    }
    
    // FiltersInfo used by Order
    type FiltersInfo struct {

        
            PlaceholderText string  `json:"placeholder_text"`
            Value string  `json:"value"`
            Options []FilterInfoOption  `json:"options"`
            Text string  `json:"text"`
            Required bool  `json:"required"`
            Type string  `json:"type"`
         
    }
    
    // AdvanceFilterInfo used by Order
    type AdvanceFilterInfo struct {

        
            Processed []FiltersInfo  `json:"processed"`
            Filters []FiltersInfo  `json:"filters"`
            Returned []FiltersInfo  `json:"returned"`
            Unfulfilled []FiltersInfo  `json:"unfulfilled"`
            Page map[string]interface{}  `json:"page"`
            AppliedFilters map[string]interface{}  `json:"applied_filters"`
            ActionCentre []FiltersInfo  `json:"action_centre"`
         
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

        
            ReportType string  `json:"report_type"`
            DisplayName string  `json:"display_name"`
            S3Key string  `json:"s3_key"`
            ReportID string  `json:"report_id"`
            ReportRequestedAt string  `json:"report_requested_at"`
            RequestDetails map[string]interface{}  `json:"request_details"`
            ReportCreatedAt string  `json:"report_created_at"`
            ReportName string  `json:"report_name"`
            Status string  `json:"status"`
         
    }
    
    // JioCodeUpsertDataSet used by Order
    type JioCodeUpsertDataSet struct {

        
            ArticleID string  `json:"article_id"`
            ItemID string  `json:"item_id"`
            CompanyID string  `json:"company_id"`
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

        
            Data []map[string]interface{}  `json:"data"`
            TraceID string  `json:"trace_id"`
            Error []NestedErrorSchemaDataSet  `json:"error"`
            Success bool  `json:"success"`
            Identifier string  `json:"identifier"`
         
    }
    
    // URL used by Order
    type URL struct {

        
            URL string  `json:"url"`
         
    }
    
    // FileResponse used by Order
    type FileResponse struct {

        
            FileName string  `json:"file_name"`
            Cdn URL  `json:"cdn"`
         
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
            ID float64  `json:"id"`
            QcType []string  `json:"qc_type"`
         
    }
    
    // PlatformShipmentReasonsResponse used by Order
    type PlatformShipmentReasonsResponse struct {

        
            Success bool  `json:"success"`
            Reasons []Reason  `json:"reasons"`
         
    }
    
    // ShipmentResponseReasons used by Order
    type ShipmentResponseReasons struct {

        
            ReasonID float64  `json:"reason_id"`
            Reason string  `json:"reason"`
         
    }
    
    // ShipmentReasonsResponse used by Order
    type ShipmentReasonsResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
            Reasons []ShipmentResponseReasons  `json:"reasons"`
         
    }
    
    // ArticleDetails used by Order
    type ArticleDetails struct {

        
            Status map[string]interface{}  `json:"status"`
         
    }
    
    // StoreEinvoice used by Order
    type StoreEinvoice struct {

        
            Username string  `json:"username"`
            Password string  `json:"password"`
            Enabled bool  `json:"enabled"`
            User string  `json:"user"`
         
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

        
            User string  `json:"user"`
            Password string  `json:"password"`
            Username string  `json:"username"`
         
    }
    
    // Document used by Order
    type Document struct {

        
            Verified bool  `json:"verified"`
            LegalName string  `json:"legal_name"`
            Value string  `json:"value"`
            DsType string  `json:"ds_type"`
            URL string  `json:"url"`
         
    }
    
    // StoreDocuments used by Order
    type StoreDocuments struct {

        
            Gst Document  `json:"gst"`
         
    }
    
    // StoreMeta used by Order
    type StoreMeta struct {

        
            AdditionalContactDetails map[string]interface{}  `json:"additional_contact_details"`
            DisplayName string  `json:"display_name"`
            GstNumber string  `json:"gst_number"`
            NotificationEmails []string  `json:"notification_emails"`
            ProductReturnConfig map[string]interface{}  `json:"product_return_config"`
            Timing []map[string]interface{}  `json:"timing"`
            EwaybillPortalDetails map[string]interface{}  `json:"ewaybill_portal_details"`
            Stage string  `json:"stage"`
            GstCredentials StoreGstCredentials  `json:"gst_credentials"`
            EinvoicePortalDetails EInvoicePortalDetails  `json:"einvoice_portal_details"`
            Documents StoreDocuments  `json:"documents"`
         
    }
    
    // StoreAddress used by Order
    type StoreAddress struct {

        
            AddressType string  `json:"address_type"`
            Area string  `json:"area"`
            Landmark string  `json:"landmark"`
            Latitude float64  `json:"latitude"`
            Country string  `json:"country"`
            CountryCode string  `json:"country_code"`
            ContactPerson string  `json:"contact_person"`
            State string  `json:"state"`
            Address2 string  `json:"address2"`
            UpdatedAt string  `json:"updated_at"`
            AddressCategory string  `json:"address_category"`
            Longitude float64  `json:"longitude"`
            Email string  `json:"email"`
            Pincode float64  `json:"pincode"`
            City string  `json:"city"`
            Phone string  `json:"phone"`
            CreatedAt string  `json:"created_at"`
            Version string  `json:"version"`
            Address1 string  `json:"address1"`
         
    }
    
    // Store used by Order
    type Store struct {

        
            BrandID interface{}  `json:"brand_id"`
            Latitude float64  `json:"latitude"`
            Country string  `json:"country"`
            Meta StoreMeta  `json:"meta"`
            SID string  `json:"s_id"`
            ContactPerson string  `json:"contact_person"`
            AlohomoraUserID float64  `json:"alohomora_user_id"`
            State string  `json:"state"`
            MallArea string  `json:"mall_area"`
            Address2 string  `json:"address2"`
            UpdatedAt string  `json:"updated_at"`
            LoginUsername string  `json:"login_username"`
            BrandStoreTags []string  `json:"brand_store_tags"`
            StoreActiveFrom string  `json:"store_active_from"`
            OrderIntegrationID string  `json:"order_integration_id"`
            IsActive bool  `json:"is_active"`
            PackagingMaterialCount float64  `json:"packaging_material_count"`
            Code string  `json:"code"`
            MallName string  `json:"mall_name"`
            IsArchived bool  `json:"is_archived"`
            Longitude float64  `json:"longitude"`
            FulfillmentChannel string  `json:"fulfillment_channel"`
            Name string  `json:"name"`
            Pincode string  `json:"pincode"`
            IsEnabledForRecon bool  `json:"is_enabled_for_recon"`
            City string  `json:"city"`
            Phone float64  `json:"phone"`
            CompanyID float64  `json:"company_id"`
            CreatedAt string  `json:"created_at"`
            LocationType string  `json:"location_type"`
            ParentStoreID float64  `json:"parent_store_id"`
            StoreAddressJson StoreAddress  `json:"store_address_json"`
            VatNo string  `json:"vat_no"`
            Address1 string  `json:"address1"`
            StoreEmail string  `json:"store_email"`
         
    }
    
    // Brand used by Order
    type Brand struct {

        
            CreditNoteExpiryDays float64  `json:"credit_note_expiry_days"`
            BrandID float64  `json:"brand_id"`
            InvoicePrefix string  `json:"invoice_prefix"`
            ModifiedOn float64  `json:"modified_on"`
            PickupLocation string  `json:"pickup_location"`
            StartDate string  `json:"start_date"`
            BrandName string  `json:"brand_name"`
            ScriptLastRan string  `json:"script_last_ran"`
            IsVirtualInvoice bool  `json:"is_virtual_invoice"`
            CreatedOn float64  `json:"created_on"`
            CreditNoteAllowed bool  `json:"credit_note_allowed"`
            Company string  `json:"company"`
            Logo string  `json:"logo"`
         
    }
    
    // BagGSTDetails used by Order
    type BagGSTDetails struct {

        
            IgstTaxPercentage float64  `json:"igst_tax_percentage"`
            CgstGstFee string  `json:"cgst_gst_fee"`
            IsDefaultHsnCode bool  `json:"is_default_hsn_code"`
            IgstGstFee string  `json:"igst_gst_fee"`
            GstTag string  `json:"gst_tag"`
            ValueOfGood float64  `json:"value_of_good"`
            GstinCode string  `json:"gstin_code"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            SgstTaxPercentage float64  `json:"sgst_tax_percentage"`
            HsnCodeID string  `json:"hsn_code_id"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            CgstTaxPercentage float64  `json:"cgst_tax_percentage"`
            GstFee float64  `json:"gst_fee"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            SgstGstFee string  `json:"sgst_gst_fee"`
            HsnCode string  `json:"hsn_code"`
         
    }
    
    // Attributes used by Order
    type Attributes struct {

        
            Essential string  `json:"essential"`
            MarketerAddress string  `json:"marketer_address"`
            MarketerName string  `json:"marketer_name"`
            PrimaryMaterial string  `json:"primary_material"`
            PrimaryColor string  `json:"primary_color"`
            BrandName string  `json:"brand_name"`
            Name string  `json:"name"`
            PrimaryColorHex string  `json:"primary_color_hex"`
            Gender []string  `json:"gender"`
         
    }
    
    // Item used by Order
    type Item struct {

        
            BrandID float64  `json:"brand_id"`
            L1Category []string  `json:"l1_category"`
            ItemID float64  `json:"item_id"`
            Color string  `json:"color"`
            CanReturn bool  `json:"can_return"`
            L3Category float64  `json:"l3_category"`
            DepartmentID float64  `json:"department_id"`
            Meta map[string]interface{}  `json:"meta"`
            Size string  `json:"size"`
            Brand string  `json:"brand"`
            LastUpdatedAt string  `json:"last_updated_at"`
            Gender string  `json:"gender"`
            BranchURL string  `json:"branch_url"`
            Code string  `json:"code"`
            SlugKey string  `json:"slug_key"`
            L3CategoryName string  `json:"l3_category_name"`
            L1CategoryID float64  `json:"l1_category_id"`
            L2Category []string  `json:"l2_category"`
            WebstoreProductURL string  `json:"webstore_product_url"`
            Name string  `json:"name"`
            CanCancel bool  `json:"can_cancel"`
            Image []string  `json:"image"`
            L2CategoryID float64  `json:"l2_category_id"`
            Attributes Attributes  `json:"attributes"`
         
    }
    
    // BagReturnableCancelableStatus1 used by Order
    type BagReturnableCancelableStatus1 struct {

        
            IsActive bool  `json:"is_active"`
            IsCustomerReturnAllowed bool  `json:"is_customer_return_allowed"`
            IsReturnable bool  `json:"is_returnable"`
            CanBeCancelled bool  `json:"can_be_cancelled"`
            EnableTracking bool  `json:"enable_tracking"`
         
    }
    
    // BagDetailsPlatformResponse used by Order
    type BagDetailsPlatformResponse struct {

        
            BagStatus []BagStatusHistory  `json:"bag_status"`
            ArticleDetails ArticleDetails  `json:"article_details"`
            Prices Prices  `json:"prices"`
            FinancialBreakup []FinancialBreakup  `json:"financial_breakup"`
            LineNumber float64  `json:"line_number"`
            ShipmentID string  `json:"shipment_id"`
            RestorePromos map[string]interface{}  `json:"restore_promos"`
            Type string  `json:"type"`
            OriginalBagList []float64  `json:"original_bag_list"`
            Meta BagMeta  `json:"meta"`
            OrderingStore Store  `json:"ordering_store"`
            CurrentOperationalStatus BagStatusHistory  `json:"current_operational_status"`
            Brand Brand  `json:"brand"`
            BagUpdateTime float64  `json:"bag_update_time"`
            EntityType string  `json:"entity_type"`
            SellerIdentifier string  `json:"seller_identifier"`
            JourneyType string  `json:"journey_type"`
            Identifier string  `json:"identifier"`
            GstDetails BagGSTDetails  `json:"gst_details"`
            OrderIntegrationID string  `json:"order_integration_id"`
            DisplayName string  `json:"display_name"`
            CurrentStatus BagStatusHistory  `json:"current_status"`
            QcRequired interface{}  `json:"qc_required"`
            AffiliateDetails AffiliateDetails  `json:"affiliate_details"`
            ID float64  `json:"id"`
            NoOfBagsOrder float64  `json:"no_of_bags_order"`
            Reasons []map[string]interface{}  `json:"reasons"`
            AppliedPromos []map[string]interface{}  `json:"applied_promos"`
            Item Item  `json:"item"`
            Status BagReturnableCancelableStatus1  `json:"status"`
            ParentPromoBags map[string]interface{}  `json:"parent_promo_bags"`
            Quantity float64  `json:"quantity"`
            Dates Dates  `json:"dates"`
            Article Article  `json:"article"`
            BagStatusHistory BagStatusHistory  `json:"bag_status_history"`
            OperationalStatus string  `json:"operational_status"`
            RestoreCoupon bool  `json:"restore_coupon"`
            Tags []string  `json:"tags"`
            AffiliateBagDetails AffiliateBagDetails  `json:"affiliate_bag_details"`
         
    }
    
    // ErrorResponse used by Order
    type ErrorResponse struct {

        
            Message string  `json:"message"`
            Error string  `json:"error"`
         
    }
    
    // BagsPage used by Order
    type BagsPage struct {

        
            Size float64  `json:"size"`
            PageType string  `json:"page_type"`
            ItemTotal float64  `json:"item_total"`
            Current float64  `json:"current"`
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
            PaymentReceipt string  `json:"payment_receipt"`
            MerchantCnReceipt string  `json:"merchant_cn_receipt"`
            Success bool  `json:"success"`
            OrderID string  `json:"order_id"`
            InvoiceReceipt string  `json:"invoice_receipt"`
         
    }
    
    // InvalidateShipmentCachePayload used by Order
    type InvalidateShipmentCachePayload struct {

        
            AffiliateBagIds []string  `json:"affiliate_bag_ids"`
            BagIds []string  `json:"bag_ids"`
            ShipmentIds []string  `json:"shipment_ids"`
         
    }
    
    // InvalidateShipmentCacheNestedResponse used by Order
    type InvalidateShipmentCacheNestedResponse struct {

        
            ShipmentID string  `json:"shipment_id"`
            Error string  `json:"error"`
            Status float64  `json:"status"`
            Message string  `json:"message"`
         
    }
    
    // InvalidateShipmentCacheResponse used by Order
    type InvalidateShipmentCacheResponse struct {

        
            Response []InvalidateShipmentCacheNestedResponse  `json:"response"`
         
    }
    
    // ErrorResponse1 used by Order
    type ErrorResponse1 struct {

        
            ErrorTrace string  `json:"error_trace"`
            Status float64  `json:"status"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // StoreReassign used by Order
    type StoreReassign struct {

        
            AffiliateOrderID string  `json:"affiliate_order_id"`
            ItemID string  `json:"item_id"`
            FyndOrderID string  `json:"fynd_order_id"`
            AffiliateID string  `json:"affiliate_id"`
            BagID float64  `json:"bag_id"`
            ReasonIds []float64  `json:"reason_ids"`
            StoreID float64  `json:"store_id"`
            MongoArticleID string  `json:"mongo_article_id"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
            SetID string  `json:"set_id"`
         
    }
    
    // StoreReassignResponse used by Order
    type StoreReassignResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // Entities used by Order
    type Entities struct {

        
            AffiliateOrderID string  `json:"affiliate_order_id"`
            ReasonText string  `json:"reason_text"`
            ID string  `json:"id"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
            AffiliateID string  `json:"affiliate_id"`
         
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

        
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            AffiliateID string  `json:"affiliate_id"`
         
    }
    
    // Bags used by Order
    type Bags struct {

        
            AffiliateOrderID string  `json:"affiliate_order_id"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
            IsLocked bool  `json:"is_locked"`
            BagID float64  `json:"bag_id"`
         
    }
    
    // CheckResponse used by Order
    type CheckResponse struct {

        
            IsBagLocked bool  `json:"is_bag_locked"`
            IsShipmentLocked bool  `json:"is_shipment_locked"`
            AffiliateID string  `json:"affiliate_id"`
            Status string  `json:"status"`
            ShipmentID string  `json:"shipment_id"`
            OriginalFilter OriginalFilter  `json:"original_filter"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            Bags []Bags  `json:"bags"`
            LockStatus string  `json:"lock_status"`
         
    }
    
    // UpdateShipmentLockResponse used by Order
    type UpdateShipmentLockResponse struct {

        
            CheckResponse []CheckResponse  `json:"check_response"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // AnnouncementResponse used by Order
    type AnnouncementResponse struct {

        
            ID float64  `json:"id"`
            Description string  `json:"description"`
            Title string  `json:"title"`
            FromDatetime string  `json:"from_datetime"`
            CompanyID float64  `json:"company_id"`
            PlatformID string  `json:"platform_id"`
            CreatedAt string  `json:"created_at"`
            LogoURL string  `json:"logo_url"`
            ToDatetime string  `json:"to_datetime"`
            PlatformName string  `json:"platform_name"`
         
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

        
            CallID string  `json:"call_id"`
            Success bool  `json:"success"`
         
    }
    
    // ErrorDetail used by Order
    type ErrorDetail struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // ProductsReasonsData used by Order
    type ProductsReasonsData struct {

        
            ReasonText string  `json:"reason_text"`
            ReasonID float64  `json:"reason_id"`
         
    }
    
    // ProductsReasonsFilters used by Order
    type ProductsReasonsFilters struct {

        
            LineNumber float64  `json:"line_number"`
            Identifier string  `json:"identifier"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // ProductsReasons used by Order
    type ProductsReasons struct {

        
            Data ProductsReasonsData  `json:"data"`
            Filters []ProductsReasonsFilters  `json:"filters"`
         
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
    
    // ReasonsData used by Order
    type ReasonsData struct {

        
            Products []ProductsReasons  `json:"products"`
            Entities []EntitiesReasons  `json:"entities"`
         
    }
    
    // Products used by Order
    type Products struct {

        
            LineNumber float64  `json:"line_number"`
            Identifier string  `json:"identifier"`
            Quantity float64  `json:"quantity"`
         
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
    
    // EntitiesDataUpdates used by Order
    type EntitiesDataUpdates struct {

        
            Data map[string]interface{}  `json:"data"`
            Filters []map[string]interface{}  `json:"filters"`
         
    }
    
    // DataUpdates used by Order
    type DataUpdates struct {

        
            OrderItemStatus []OrderItemDataUpdates  `json:"order_item_status"`
            Products []ProductsDataUpdates  `json:"products"`
            Entities []EntitiesDataUpdates  `json:"entities"`
         
    }
    
    // ShipmentsRequest used by Order
    type ShipmentsRequest struct {

        
            Reasons ReasonsData  `json:"reasons"`
            Products []Products  `json:"products"`
            DataUpdates DataUpdates  `json:"data_updates"`
            Identifier string  `json:"identifier"`
         
    }
    
    // StatuesRequest used by Order
    type StatuesRequest struct {

        
            Shipments []ShipmentsRequest  `json:"shipments"`
            ExcludeBagsNextState string  `json:"exclude_bags_next_state"`
            Status string  `json:"status"`
         
    }
    
    // UpdateShipmentStatusRequest used by Order
    type UpdateShipmentStatusRequest struct {

        
            UnlockBeforeTransition bool  `json:"unlock_before_transition"`
            Statuses []StatuesRequest  `json:"statuses"`
            Task bool  `json:"task"`
            ForceTransition bool  `json:"force_transition"`
            LockAfterTransition bool  `json:"lock_after_transition"`
         
    }
    
    // ShipmentsResponse used by Order
    type ShipmentsResponse struct {

        
            Exception string  `json:"exception"`
            Identifier string  `json:"identifier"`
            Status float64  `json:"status"`
            Code string  `json:"code"`
            Meta map[string]interface{}  `json:"meta"`
            StackTrace string  `json:"stack_trace"`
            Message string  `json:"message"`
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
    
    // DateRange used by Order
    type DateRange struct {

        
            FromDate string  `json:"from_date"`
            ToDate string  `json:"to_date"`
         
    }
    
    // FiltersRequest used by Order
    type FiltersRequest struct {

        
            Stores float64  `json:"stores"`
            Lane string  `json:"lane"`
            DpIds float64  `json:"dp_ids"`
            DpName string  `json:"dp_name"`
            Logo string  `json:"logo"`
            StoreName string  `json:"store_name"`
            DateRange DateRange  `json:"date_range"`
         
    }
    
    // ProcessManifest used by Order
    type ProcessManifest struct {

        
            Action string  `json:"action"`
            ManifestID string  `json:"manifest_id"`
            UniqueID string  `json:"unique_id"`
            Filters FiltersRequest  `json:"filters"`
         
    }
    
    // Filters used by Order
    type Filters struct {

        
            Stores float64  `json:"stores"`
            Lane string  `json:"lane"`
            ToDate string  `json:"to_date"`
            DpIds float64  `json:"dp_ids"`
            SelectedShipments string  `json:"selected_shipments"`
            DpName string  `json:"dp_name"`
            Logo string  `json:"logo"`
            StoreName string  `json:"store_name"`
            FromDate string  `json:"from_date"`
            DateRange DateRange  `json:"date_range"`
         
    }
    
    // ProcessManifestResponse used by Order
    type ProcessManifestResponse struct {

        
            ManifestID string  `json:"manifest_id"`
            UID string  `json:"uid"`
            CreatedBy string  `json:"created_by"`
            CompanyID float64  `json:"company_id"`
            Action string  `json:"action"`
            Filters Filters  `json:"filters"`
            UserID string  `json:"user_id"`
         
    }
    
    // ProcessManifestItemResponse used by Order
    type ProcessManifestItemResponse struct {

        
            Items ProcessManifestResponse  `json:"items"`
         
    }
    
    // FilterInfoOption1 used by Order
    type FilterInfoOption1 struct {

        
            Value string  `json:"value"`
            Text string  `json:"text"`
            Name string  `json:"name"`
            PlaceholderText string  `json:"placeholder_text"`
            ShowUI bool  `json:"show_ui"`
            MinSearchSize float64  `json:"min_search_size"`
         
    }
    
    // FiltersInfo1 used by Order
    type FiltersInfo1 struct {

        
            Value string  `json:"value"`
            Text string  `json:"text"`
            Options []FilterInfoOption1  `json:"options"`
            PlaceholderText string  `json:"placeholder_text"`
            Type string  `json:"type"`
            Required bool  `json:"required"`
         
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

        
            DisplayText string  `json:"display_text"`
            ID float64  `json:"id"`
            Description string  `json:"description"`
            Slug string  `json:"slug"`
         
    }
    
    // GetActionsResponse used by Order
    type GetActionsResponse struct {

        
            Permissions ActionInfo  `json:"permissions"`
         
    }
    
    // PostHistoryData used by Order
    type PostHistoryData struct {

        
            UserName string  `json:"user_name"`
            Message string  `json:"message"`
         
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
    
    // HistoryDict used by Order
    type HistoryDict struct {

        
            DisplayMessage string  `json:"display_message"`
            L3Detail string  `json:"l3_detail"`
            BagID float64  `json:"bag_id"`
            TicketURL string  `json:"ticket_url"`
            TicketID string  `json:"ticket_id"`
            Meta map[string]interface{}  `json:"meta"`
            User string  `json:"user"`
            Createdat string  `json:"createdat"`
            Type string  `json:"type"`
            Message string  `json:"message"`
            L2Detail string  `json:"l2_detail"`
            L1Detail string  `json:"l1_detail"`
            AssignedAgent string  `json:"assigned_agent"`
         
    }
    
    // ShipmentHistoryResponse used by Order
    type ShipmentHistoryResponse struct {

        
            ActivityHistory []HistoryDict  `json:"activity_history"`
            Success bool  `json:"success"`
         
    }
    
    // SmsDataPayload used by Order
    type SmsDataPayload struct {

        
            CustomerName string  `json:"customer_name"`
            ShipmentID float64  `json:"shipment_id"`
            PaymentMode string  `json:"payment_mode"`
            PhoneNumber float64  `json:"phone_number"`
            CountryCode string  `json:"country_code"`
            Message string  `json:"message"`
            AmountPaid float64  `json:"amount_paid"`
            BrandName string  `json:"brand_name"`
            OrderID string  `json:"order_id"`
         
    }
    
    // SendSmsPayload used by Order
    type SendSmsPayload struct {

        
            Data SmsDataPayload  `json:"data"`
            Slug string  `json:"slug"`
            BagID float64  `json:"bag_id"`
         
    }
    
    // OrderDetails used by Order
    type OrderDetails struct {

        
            CreatedAt string  `json:"created_at"`
            FyndOrderID string  `json:"fynd_order_id"`
         
    }
    
    // Meta used by Order
    type Meta struct {

        
            KafkaEmissionStatus float64  `json:"kafka_emission_status"`
            StateManagerUsed string  `json:"state_manager_used"`
         
    }
    
    // ShipmentDetail used by Order
    type ShipmentDetail struct {

        
            Remarks string  `json:"remarks"`
            BagList []float64  `json:"bag_list"`
            ID float64  `json:"id"`
            Status string  `json:"status"`
            ShipmentID string  `json:"shipment_id"`
            Meta Meta  `json:"meta"`
         
    }
    
    // OrderStatusData used by Order
    type OrderStatusData struct {

        
            Errors []string  `json:"errors"`
            OrderDetails OrderDetails  `json:"order_details"`
            ShipmentDetails []ShipmentDetail  `json:"shipment_details"`
         
    }
    
    // OrderStatusResult used by Order
    type OrderStatusResult struct {

        
            Result []OrderStatusData  `json:"result"`
            Success string  `json:"success"`
         
    }
    
    // Dimension used by Order
    type Dimension struct {

        
            Width float64  `json:"width"`
            Length float64  `json:"length"`
            Height string  `json:"height"`
            Weight string  `json:"weight"`
            PackagingType string  `json:"packaging_type"`
         
    }
    
    // UpdatePackagingDimensionsPayload used by Order
    type UpdatePackagingDimensionsPayload struct {

        
            Dimension []Dimension  `json:"dimension"`
            CurrentStatus string  `json:"current_status"`
            ShipmentID string  `json:"shipment_id"`
         
    }
    
    // UpdatePackagingDimensionsResponse used by Order
    type UpdatePackagingDimensionsResponse struct {

        
            Message string  `json:"message"`
         
    }
    
    // BillingInfo used by Order
    type BillingInfo struct {

        
            State string  `json:"state"`
            Address1 string  `json:"address1"`
            Pincode string  `json:"pincode"`
            AlternateEmail string  `json:"alternate_email"`
            FloorNo string  `json:"floor_no"`
            PrimaryMobileNumber string  `json:"primary_mobile_number"`
            PrimaryEmail string  `json:"primary_email"`
            MiddleName string  `json:"middle_name"`
            City string  `json:"city"`
            Gender string  `json:"gender"`
            CustomerCode string  `json:"customer_code"`
            AlternateMobileNumber string  `json:"alternate_mobile_number"`
            HouseNo string  `json:"house_no"`
            Country string  `json:"country"`
            ExternalCustomerCode string  `json:"external_customer_code"`
            Title string  `json:"title"`
            LastName string  `json:"last_name"`
            FirstName string  `json:"first_name"`
            Address2 string  `json:"address2"`
            CountryCode string  `json:"country_code"`
            StateCode string  `json:"state_code"`
         
    }
    
    // PaymentMethod used by Order
    type PaymentMethod struct {

        
            Mode string  `json:"mode"`
            RefundBy string  `json:"refund_by"`
            Name string  `json:"name"`
            CollectBy string  `json:"collect_by"`
            TransactionData map[string]interface{}  `json:"transaction_data"`
            Meta map[string]interface{}  `json:"meta"`
            Amount float64  `json:"amount"`
         
    }
    
    // PaymentInfo used by Order
    type PaymentInfo struct {

        
            PrimaryMode string  `json:"primary_mode"`
            PaymentMethods []PaymentMethod  `json:"payment_methods"`
         
    }
    
    // TaxInfo used by Order
    type TaxInfo struct {

        
            B2bGstinNumber string  `json:"b2b_gstin_number"`
            Gstin string  `json:"gstin"`
         
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

        
            Name string  `json:"name"`
            Code string  `json:"code"`
            Type string  `json:"type"`
            Amount map[string]interface{}  `json:"amount"`
            Tax Tax  `json:"tax"`
         
    }
    
    // LineItem used by Order
    type LineItem struct {

        
            Quantity float64  `json:"quantity"`
            ExternalLineID string  `json:"external_line_id"`
            Meta map[string]interface{}  `json:"meta"`
            SellerIdentifier string  `json:"seller_identifier"`
            CustomMessasge string  `json:"custom_messasge"`
            Charges []Charge  `json:"charges"`
         
    }
    
    // ProcessingDates used by Order
    type ProcessingDates struct {

        
            DispatchByDate string  `json:"dispatch_by_date"`
            CustomerPickupSlot map[string]interface{}  `json:"customer_pickup_slot"`
            PackByDate string  `json:"pack_by_date"`
            DpPickupSlot map[string]interface{}  `json:"dp_pickup_slot"`
            ConfirmByDate string  `json:"confirm_by_date"`
            DispatchAfterDate string  `json:"dispatch_after_date"`
         
    }
    
    // Shipment used by Order
    type Shipment struct {

        
            LineItems []LineItem  `json:"line_items"`
            Meta map[string]interface{}  `json:"meta"`
            Priority float64  `json:"priority"`
            ExternalShipmentID string  `json:"external_shipment_id"`
            LocationID float64  `json:"location_id"`
            ProcessingDates ProcessingDates  `json:"processing_dates"`
         
    }
    
    // ShippingInfo used by Order
    type ShippingInfo struct {

        
            State string  `json:"state"`
            Address1 string  `json:"address1"`
            Pincode string  `json:"pincode"`
            AlternateEmail string  `json:"alternate_email"`
            GeoLocation map[string]interface{}  `json:"geo_location"`
            FloorNo string  `json:"floor_no"`
            PrimaryMobileNumber string  `json:"primary_mobile_number"`
            PrimaryEmail string  `json:"primary_email"`
            Landmark string  `json:"landmark"`
            ShippingType string  `json:"shipping_type"`
            MiddleName string  `json:"middle_name"`
            City string  `json:"city"`
            Gender string  `json:"gender"`
            Slot []map[string]interface{}  `json:"slot"`
            CustomerCode string  `json:"customer_code"`
            AlternateMobileNumber string  `json:"alternate_mobile_number"`
            HouseNo string  `json:"house_no"`
            Country string  `json:"country"`
            ExternalCustomerCode string  `json:"external_customer_code"`
            AddressType string  `json:"address_type"`
            Title string  `json:"title"`
            LastName string  `json:"last_name"`
            FirstName string  `json:"first_name"`
            Address2 string  `json:"address2"`
            CountryCode string  `json:"country_code"`
            StateCode string  `json:"state_code"`
         
    }
    
    // CreateOrderAPI used by Order
    type CreateOrderAPI struct {

        
            CurrencyInfo map[string]interface{}  `json:"currency_info"`
            Config map[string]interface{}  `json:"config"`
            BillingInfo BillingInfo  `json:"billing_info"`
            PaymentInfo PaymentInfo  `json:"payment_info"`
            TaxInfo TaxInfo  `json:"tax_info"`
            Meta map[string]interface{}  `json:"meta"`
            Shipments []Shipment  `json:"shipments"`
            ExternalCreationDate string  `json:"external_creation_date"`
            ExternalOrderID string  `json:"external_order_id"`
            ShippingInfo ShippingInfo  `json:"shipping_info"`
            Charges []Charge  `json:"charges"`
         
    }
    
    // CreateOrderResponse used by Order
    type CreateOrderResponse struct {

        
            FyndOrderID string  `json:"fynd_order_id"`
         
    }
    
    // CreateOrderErrorReponse used by Order
    type CreateOrderErrorReponse struct {

        
            Exception string  `json:"exception"`
            Status float64  `json:"status"`
            RequestID string  `json:"request_id"`
            Code string  `json:"code"`
            Meta string  `json:"meta"`
            StackTrace string  `json:"stack_trace"`
            Info interface{}  `json:"info"`
            Message string  `json:"message"`
         
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
            Mobile float64  `json:"mobile"`
            OrderDetails []FyndOrderIdList  `json:"order_details"`
            StartDate string  `json:"start_date"`
         
    }
    
    // BagStateTransitionMap used by Order
    type BagStateTransitionMap struct {

        
            Fynd map[string]interface{}  `json:"fynd"`
            Affiliate map[string]interface{}  `json:"affiliate"`
         
    }
    
    // ManifestPageInfo used by Order
    type ManifestPageInfo struct {

        
            Current float64  `json:"current"`
            HasNext bool  `json:"has_next"`
            Type string  `json:"type"`
            HasPrevious bool  `json:"has_previous"`
            Total float64  `json:"total"`
            Size float64  `json:"size"`
         
    }
    
    // ManifestItemDetails used by Order
    type ManifestItemDetails struct {

        
            ItemQty float64  `json:"item_qty"`
            ShipmentCreatedAt string  `json:"shipment_created_at"`
            ShipmentID string  `json:"shipment_id"`
            InvoiceID string  `json:"invoice_id"`
            Awb string  `json:"awb"`
            OrderID string  `json:"order_id"`
         
    }
    
    // ManifestShipmentListing used by Order
    type ManifestShipmentListing struct {

        
            Lane string  `json:"lane"`
            Status float64  `json:"status"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
            TotalCount float64  `json:"total_count"`
            Page ManifestPageInfo  `json:"page"`
            Items []ManifestItemDetails  `json:"items"`
         
    }
    
    // ManifestFile used by Order
    type ManifestFile struct {

        
            Bucket string  `json:"bucket"`
            Region string  `json:"region"`
            Key string  `json:"key"`
         
    }
    
    // ManifestMediaUpdate used by Order
    type ManifestMediaUpdate struct {

        
            Entity string  `json:"entity"`
            Status bool  `json:"status"`
            Code float64  `json:"code"`
            MediaType string  `json:"media_type"`
            File ManifestFile  `json:"file"`
            Link string  `json:"link"`
         
    }
    
    // PDFMeta used by Order
    type PDFMeta struct {

        
            MediaUpdates []ManifestMediaUpdate  `json:"media_updates"`
            Consent string  `json:"consent"`
         
    }
    
    // TotalShipmentPricesCount used by Order
    type TotalShipmentPricesCount struct {

        
            TotalPrice float64  `json:"total_price"`
            ShipmentCount float64  `json:"shipment_count"`
         
    }
    
    // ManifestMeta used by Order
    type ManifestMeta struct {

        
            Filters Filters  `json:"filters"`
            TotalShipmentPricesCount TotalShipmentPricesCount  `json:"total_shipment_prices_count"`
         
    }
    
    // Manifest used by Order
    type Manifest struct {

        
            ManifestID string  `json:"manifest_id"`
            UID string  `json:"uid"`
            ID float64  `json:"id"`
            CreatedBy string  `json:"created_by"`
            PdfMeta PDFMeta  `json:"pdf_meta"`
            Status string  `json:"status"`
            UserID string  `json:"user_id"`
            CompanyID float64  `json:"company_id"`
            Meta ManifestMeta  `json:"meta"`
            Filters Filters  `json:"filters"`
            IsActive bool  `json:"is_active"`
            CreatedAt string  `json:"created_at"`
         
    }
    
    // ManifestList used by Order
    type ManifestList struct {

        
            Page ManifestPageInfo  `json:"page"`
            Items []Manifest  `json:"items"`
         
    }
    
    // ManifestDetails used by Order
    type ManifestDetails struct {

        
            Page ManifestPageInfo  `json:"page"`
            Items []ManifestItemDetails  `json:"items"`
            AdditionalShipmentCount float64  `json:"additional_shipment_count"`
            ManifestDetails []Manifest  `json:"manifest_details"`
         
    }
    
    // FetchCreditBalanceRequestPayload used by Order
    type FetchCreditBalanceRequestPayload struct {

        
            SellerID string  `json:"seller_id"`
            CustomerMobileNumber string  `json:"customer_mobile_number"`
            AffiliateID string  `json:"affiliate_id"`
         
    }
    
    // CreditBalanceInfo used by Order
    type CreditBalanceInfo struct {

        
            CustomerMobileNumber string  `json:"customer_mobile_number"`
            TotalCreditedBalance string  `json:"total_credited_balance"`
            Reason string  `json:"reason"`
         
    }
    
    // FetchCreditBalanceResponsePayload used by Order
    type FetchCreditBalanceResponsePayload struct {

        
            Data CreditBalanceInfo  `json:"data"`
            Success bool  `json:"success"`
         
    }
    
    // RefundModeConfigRequestPayload used by Order
    type RefundModeConfigRequestPayload struct {

        
            CustomerMobileNumber string  `json:"customer_mobile_number"`
            SellerID string  `json:"seller_id"`
            FyndOrderID string  `json:"fynd_order_id"`
            AffiliateID string  `json:"affiliate_id"`
         
    }
    
    // RefundOption used by Order
    type RefundOption struct {

        
            Value string  `json:"value"`
            Name string  `json:"name"`
            Type string  `json:"type"`
            Slug string  `json:"slug"`
         
    }
    
    // RefundModeInfo used by Order
    type RefundModeInfo struct {

        
            Options []RefundOption  `json:"options"`
            Slug string  `json:"slug"`
            DisplayName string  `json:"display_name"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // RefundModeConfigResponsePayload used by Order
    type RefundModeConfigResponsePayload struct {

        
            Data []RefundModeInfo  `json:"data"`
            Success bool  `json:"success"`
         
    }
    
    // AttachUserInfo used by Order
    type AttachUserInfo struct {

        
            Mobile string  `json:"mobile"`
            LastName string  `json:"last_name"`
            CountryCode string  `json:"country_code"`
            FirstName string  `json:"first_name"`
         
    }
    
    // AttachUserOtpData used by Order
    type AttachUserOtpData struct {

        
            RequestID string  `json:"request_id"`
         
    }
    
    // AttachOrderUser used by Order
    type AttachOrderUser struct {

        
            UserInfo AttachUserInfo  `json:"user_info"`
            OtpData AttachUserOtpData  `json:"otp_data"`
            FyndOrderID string  `json:"fynd_order_id"`
         
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

        
            Data PointBlankOtpData  `json:"data"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // VerifyOtpData used by Order
    type VerifyOtpData struct {

        
            RequestID string  `json:"request_id"`
            Mobile string  `json:"mobile"`
            OtpCode float64  `json:"otp_code"`
         
    }
    
    // VerifyMobileOTP used by Order
    type VerifyMobileOTP struct {

        
            OtpData VerifyOtpData  `json:"otp_data"`
            FyndOrderID string  `json:"fynd_order_id"`
         
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

        
            Data VerifyOtpResponseData  `json:"data"`
            Success bool  `json:"success"`
            Status float64  `json:"status"`
            Message string  `json:"message"`
         
    }
    

    
    // DeleteResponse used by Catalog
    type DeleteResponse struct {

        
            Message string  `json:"message"`
         
    }
    
    // ErrorResponse used by Catalog
    type ErrorResponse struct {

        
            Status float64  `json:"status"`
            Meta map[string]interface{}  `json:"meta"`
            Message string  `json:"message"`
            Error string  `json:"error"`
            Code string  `json:"code"`
         
    }
    
    // GetSearchWordsData used by Catalog
    type GetSearchWordsData struct {

        
            AppID string  `json:"app_id"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            UID string  `json:"uid"`
            Words []string  `json:"words"`
            IsActive bool  `json:"is_active"`
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
    
    // SearchKeywordResult used by Catalog
    type SearchKeywordResult struct {

        
            SortOn string  `json:"sort_on"`
            Query map[string]interface{}  `json:"query"`
         
    }
    
    // CreateSearchKeyword used by Catalog
    type CreateSearchKeyword struct {

        
            AppID string  `json:"app_id"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Words []string  `json:"words"`
            IsActive bool  `json:"is_active"`
            Result SearchKeywordResult  `json:"result"`
         
    }
    
    // GetSearchWordsResponse used by Catalog
    type GetSearchWordsResponse struct {

        
            Items []GetSearchWordsData  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // GetAutocompleteWordsData used by Catalog
    type GetAutocompleteWordsData struct {

        
            AppID string  `json:"app_id"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            UID string  `json:"uid"`
            Words []string  `json:"words"`
            Results []map[string]interface{}  `json:"results"`
         
    }
    
    // GetAutocompleteWordsResponse used by Catalog
    type GetAutocompleteWordsResponse struct {

        
            Items []GetAutocompleteWordsData  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // AutoCompleteMedia used by Catalog
    type AutoCompleteMedia struct {

        
            URL string  `json:"url"`
            AspectRatio string  `json:"aspect_ratio"`
            Type string  `json:"type"`
         
    }
    
    // AutocompletePageAction used by Catalog
    type AutocompletePageAction struct {

        
            URL string  `json:"url"`
            Query map[string]interface{}  `json:"query"`
            Params map[string]interface{}  `json:"params"`
            Type string  `json:"type"`
         
    }
    
    // AutocompleteAction used by Catalog
    type AutocompleteAction struct {

        
            Page AutocompletePageAction  `json:"page"`
            Type string  `json:"type"`
         
    }
    
    // AutocompleteResult used by Catalog
    type AutocompleteResult struct {

        
            Logo AutoCompleteMedia  `json:"logo"`
            Display string  `json:"display"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Action AutocompleteAction  `json:"action"`
         
    }
    
    // CreateAutocompleteKeyword used by Catalog
    type CreateAutocompleteKeyword struct {

        
            AppID string  `json:"app_id"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Words []string  `json:"words"`
            Results []AutocompleteResult  `json:"results"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // CreateAutocompleteWordsResponse used by Catalog
    type CreateAutocompleteWordsResponse struct {

        
            Words []string  `json:"words"`
            Results []map[string]interface{}  `json:"results"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            AppID string  `json:"app_id"`
         
    }
    
    // ProductBundleItem used by Catalog
    type ProductBundleItem struct {

        
            AutoAddToCart bool  `json:"auto_add_to_cart"`
            ProductUID float64  `json:"product_uid"`
            MaxQuantity float64  `json:"max_quantity"`
            AllowRemove bool  `json:"allow_remove"`
            MinQuantity float64  `json:"min_quantity"`
            AutoSelect bool  `json:"auto_select"`
         
    }
    
    // ProductBundleRequest used by Catalog
    type ProductBundleRequest struct {

        
            CompanyID float64  `json:"company_id"`
            Products []ProductBundleItem  `json:"products"`
            Slug string  `json:"slug"`
            Logo string  `json:"logo"`
            CreatedOn string  `json:"created_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            PageVisibility []string  `json:"page_visibility"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Name string  `json:"name"`
            Choice string  `json:"choice"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            Meta map[string]interface{}  `json:"meta"`
            ModifiedOn string  `json:"modified_on"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // GetProductBundleCreateResponse used by Catalog
    type GetProductBundleCreateResponse struct {

        
            CompanyID float64  `json:"company_id"`
            Products []ProductBundleItem  `json:"products"`
            Slug string  `json:"slug"`
            Logo string  `json:"logo"`
            CreatedOn string  `json:"created_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            PageVisibility []string  `json:"page_visibility"`
            ID string  `json:"id"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Name string  `json:"name"`
            Choice string  `json:"choice"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            Meta map[string]interface{}  `json:"meta"`
            ModifiedOn string  `json:"modified_on"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // GetProductBundleListingResponse used by Catalog
    type GetProductBundleListingResponse struct {

        
            Items []GetProductBundleCreateResponse  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // ProductBundleUpdateRequest used by Catalog
    type ProductBundleUpdateRequest struct {

        
            CompanyID float64  `json:"company_id"`
            Products []ProductBundleItem  `json:"products"`
            Slug string  `json:"slug"`
            Logo string  `json:"logo"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            PageVisibility []string  `json:"page_visibility"`
            Name string  `json:"name"`
            Choice string  `json:"choice"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            Meta map[string]interface{}  `json:"meta"`
            ModifiedOn string  `json:"modified_on"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // Size used by Catalog
    type Size struct {

        
            Quantity float64  `json:"quantity"`
            Display string  `json:"display"`
            Value string  `json:"value"`
            IsAvailable bool  `json:"is_available"`
         
    }
    
    // Price used by Catalog
    type Price struct {

        
            MaxEffective float64  `json:"max_effective"`
            MinEffective float64  `json:"min_effective"`
            MaxMarked float64  `json:"max_marked"`
            MinMarked float64  `json:"min_marked"`
            Currency string  `json:"currency"`
         
    }
    
    // LimitedProductData used by Catalog
    type LimitedProductData struct {

        
            Attributes map[string]interface{}  `json:"attributes"`
            Identifier map[string]interface{}  `json:"identifier"`
            Slug string  `json:"slug"`
            Sizes []string  `json:"sizes"`
            UID float64  `json:"uid"`
            Price map[string]interface{}  `json:"price"`
            CountryOfOrigin string  `json:"country_of_origin"`
            ShortDescription string  `json:"short_description"`
            Name string  `json:"name"`
            Images []string  `json:"images"`
            Quantity float64  `json:"quantity"`
            ItemCode string  `json:"item_code"`
         
    }
    
    // GetProducts used by Catalog
    type GetProducts struct {

        
            AutoAddToCart bool  `json:"auto_add_to_cart"`
            ProductUID float64  `json:"product_uid"`
            Sizes []Size  `json:"sizes"`
            Price Price  `json:"price"`
            ProductDetails LimitedProductData  `json:"product_details"`
            AllowRemove bool  `json:"allow_remove"`
            MaxQuantity float64  `json:"max_quantity"`
            MinQuantity float64  `json:"min_quantity"`
            AutoSelect bool  `json:"auto_select"`
         
    }
    
    // GetProductBundleResponse used by Catalog
    type GetProductBundleResponse struct {

        
            CompanyID float64  `json:"company_id"`
            Products []GetProducts  `json:"products"`
            Slug string  `json:"slug"`
            Logo string  `json:"logo"`
            PageVisibility []string  `json:"page_visibility"`
            Name string  `json:"name"`
            Choice string  `json:"choice"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            Meta map[string]interface{}  `json:"meta"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // Meta used by Catalog
    type Meta struct {

        
            Unit string  `json:"unit"`
            Values []map[string]interface{}  `json:"values"`
            Headers map[string]interface{}  `json:"headers"`
         
    }
    
    // Guide used by Catalog
    type Guide struct {

        
            Meta Meta  `json:"meta"`
         
    }
    
    // ValidateSizeGuide used by Catalog
    type ValidateSizeGuide struct {

        
            CompanyID float64  `json:"company_id"`
            Guide Guide  `json:"guide"`
            CreatedOn string  `json:"created_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            BrandID float64  `json:"brand_id"`
            ID string  `json:"id"`
            Image string  `json:"image"`
            Name string  `json:"name"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Active bool  `json:"active"`
            Tag string  `json:"tag"`
            Title string  `json:"title"`
            Description string  `json:"description"`
            Subtitle string  `json:"subtitle"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // SuccessResponse used by Catalog
    type SuccessResponse struct {

        
            Success bool  `json:"success"`
         
    }
    
    // ListSizeGuide used by Catalog
    type ListSizeGuide struct {

        
            Items []map[string]interface{}  `json:"items"`
            Page map[string]interface{}  `json:"page"`
         
    }
    
    // SizeGuideResponse used by Catalog
    type SizeGuideResponse struct {

        
            CompanyID float64  `json:"company_id"`
            Guide map[string]interface{}  `json:"guide"`
            CreatedOn string  `json:"created_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            BrandID float64  `json:"brand_id"`
            ID string  `json:"id"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Name string  `json:"name"`
            Active bool  `json:"active"`
            Tag string  `json:"tag"`
            Title string  `json:"title"`
            Subtitle string  `json:"subtitle"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // SEOData used by Catalog
    type SEOData struct {

        
            Title string  `json:"title"`
            Description string  `json:"description"`
         
    }
    
    // MOQData used by Catalog
    type MOQData struct {

        
            IncrementUnit float64  `json:"increment_unit"`
            Maximum float64  `json:"maximum"`
            Minimum float64  `json:"minimum"`
         
    }
    
    // OwnerAppItemResponse used by Catalog
    type OwnerAppItemResponse struct {

        
            Seo SEOData  `json:"seo"`
            Moq MOQData  `json:"moq"`
            AltText map[string]interface{}  `json:"alt_text"`
            IsCod bool  `json:"is_cod"`
            IsGift bool  `json:"is_gift"`
         
    }
    
    // ApplicationItemSEO used by Catalog
    type ApplicationItemSEO struct {

        
            Title string  `json:"title"`
            Description string  `json:"description"`
         
    }
    
    // ApplicationItemMOQ used by Catalog
    type ApplicationItemMOQ struct {

        
            IncrementUnit float64  `json:"increment_unit"`
            Maximum float64  `json:"maximum"`
            Minimum float64  `json:"minimum"`
         
    }
    
    // MetaFields used by Catalog
    type MetaFields struct {

        
            Value string  `json:"value"`
            Key string  `json:"key"`
         
    }
    
    // ApplicationItemMeta used by Catalog
    type ApplicationItemMeta struct {

        
            Seo ApplicationItemSEO  `json:"seo"`
            Moq ApplicationItemMOQ  `json:"moq"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            AltText map[string]interface{}  `json:"alt_text"`
            IsCod bool  `json:"is_cod"`
            CustomMeta []MetaFields  `json:"_custom_meta"`
            IsGift bool  `json:"is_gift"`
         
    }
    
    // SuccessResponse1 used by Catalog
    type SuccessResponse1 struct {

        
            Success bool  `json:"success"`
            UID float64  `json:"uid"`
         
    }
    
    // GetConfigMetadataResponse used by Catalog
    type GetConfigMetadataResponse struct {

        
            Values []map[string]interface{}  `json:"values"`
            Condition []map[string]interface{}  `json:"condition"`
            Data []map[string]interface{}  `json:"data"`
         
    }
    
    // AttributeDetailsGroup used by Catalog
    type AttributeDetailsGroup struct {

        
            Slug string  `json:"slug"`
            Priority float64  `json:"priority"`
            Key string  `json:"key"`
            Logo string  `json:"logo"`
            Name string  `json:"name"`
            DisplayType string  `json:"display_type"`
            Unit string  `json:"unit"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // AppConfigurationDetail used by Catalog
    type AppConfigurationDetail struct {

        
            Attributes []AttributeDetailsGroup  `json:"attributes"`
            Slug string  `json:"slug"`
            IsDefault bool  `json:"is_default"`
            AppID string  `json:"app_id"`
            Logo string  `json:"logo"`
            Priority float64  `json:"priority"`
            Name string  `json:"name"`
            TemplateSlugs []string  `json:"template_slugs"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // ConfigErrorResponse used by Catalog
    type ConfigErrorResponse struct {

        
            Message string  `json:"message"`
         
    }
    
    // PageResponseType used by Catalog
    type PageResponseType struct {

        
            TotalCount float64  `json:"total_count"`
            HasNext bool  `json:"has_next"`
            Next float64  `json:"next"`
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

        
            IsDefault bool  `json:"is_default"`
            Key string  `json:"key"`
            Logo string  `json:"logo"`
            AppID string  `json:"app_id"`
            Priority float64  `json:"priority"`
            Name string  `json:"name"`
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

        
            Display string  `json:"display"`
            Key string  `json:"key"`
         
    }
    
    // MetaDataListingSortResponse used by Catalog
    type MetaDataListingSortResponse struct {

        
            Data []MetaDataListingSortMetaResponse  `json:"data"`
         
    }
    
    // MetaDataListingFilterMetaResponse used by Catalog
    type MetaDataListingFilterMetaResponse struct {

        
            Display string  `json:"display"`
            FilterTypes []string  `json:"filter_types"`
            Units []map[string]interface{}  `json:"units"`
            Key string  `json:"key"`
         
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

        
            Detail map[string]interface{}  `json:"detail"`
            Similar map[string]interface{}  `json:"similar"`
            Variant map[string]interface{}  `json:"variant"`
            Compare map[string]interface{}  `json:"compare"`
         
    }
    
    // GetCatalogConfigurationMetaData used by Catalog
    type GetCatalogConfigurationMetaData struct {

        
            Listing MetaDataListingResponse  `json:"listing"`
            Product GetCatalogConfigurationDetailsProduct  `json:"product"`
         
    }
    
    // ProductSize used by Catalog
    type ProductSize struct {

        
            Max float64  `json:"max"`
            Min float64  `json:"min"`
         
    }
    
    // ConfigurationProductConfig used by Catalog
    type ConfigurationProductConfig struct {

        
            Priority float64  `json:"priority"`
            Key string  `json:"key"`
            Logo string  `json:"logo"`
            Title string  `json:"title"`
            Size ProductSize  `json:"size"`
            Subtitle string  `json:"subtitle"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // ConfigurationProductSimilar used by Catalog
    type ConfigurationProductSimilar struct {

        
            Config []ConfigurationProductConfig  `json:"config"`
         
    }
    
    // ConfigurationProductVariantConfig used by Catalog
    type ConfigurationProductVariantConfig struct {

        
            Priority float64  `json:"priority"`
            Key string  `json:"key"`
            Logo string  `json:"logo"`
            Name string  `json:"name"`
            DisplayType string  `json:"display_type"`
            Size ProductSize  `json:"size"`
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
    
    // ConfigurationListingSortConfig used by Catalog
    type ConfigurationListingSortConfig struct {

        
            Priority float64  `json:"priority"`
            Key string  `json:"key"`
            Logo string  `json:"logo"`
            Name string  `json:"name"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // ConfigurationListingSort used by Catalog
    type ConfigurationListingSort struct {

        
            DefaultKey string  `json:"default_key"`
            Config []ConfigurationListingSortConfig  `json:"config"`
         
    }
    
    // ConfigurationBucketPoints used by Catalog
    type ConfigurationBucketPoints struct {

        
            Display string  `json:"display"`
            Start float64  `json:"start"`
            End float64  `json:"end"`
         
    }
    
    // ConfigurationListingFilterValue used by Catalog
    type ConfigurationListingFilterValue struct {

        
            BucketPoints []ConfigurationBucketPoints  `json:"bucket_points"`
            Priority []string  `json:"priority"`
            Value string  `json:"value"`
            Condition string  `json:"condition"`
            Sort string  `json:"sort"`
            Map map[string]interface{}  `json:"map"`
            MapValues []map[string]interface{}  `json:"map_values"`
         
    }
    
    // ConfigurationListingFilterConfig used by Catalog
    type ConfigurationListingFilterConfig struct {

        
            DisplayName string  `json:"display_name"`
            Type string  `json:"type"`
            Priority float64  `json:"priority"`
            Key string  `json:"key"`
            Logo string  `json:"logo"`
            Name string  `json:"name"`
            ValueConfig ConfigurationListingFilterValue  `json:"value_config"`
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
    
    // AppConfiguration used by Catalog
    type AppConfiguration struct {

        
            ConfigID string  `json:"config_id"`
            ConfigType string  `json:"config_type"`
            Type string  `json:"type"`
            AppID string  `json:"app_id"`
            CreatedOn string  `json:"created_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Product ConfigurationProduct  `json:"product"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Listing ConfigurationListing  `json:"listing"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // AppCatalogConfiguration used by Catalog
    type AppCatalogConfiguration struct {

        
            ConfigID string  `json:"config_id"`
            ConfigType string  `json:"config_type"`
            Type string  `json:"type"`
            AppID string  `json:"app_id"`
            CreatedOn string  `json:"created_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Product ConfigurationProduct  `json:"product"`
            ID string  `json:"id"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Listing ConfigurationListing  `json:"listing"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // GetAppCatalogConfiguration used by Catalog
    type GetAppCatalogConfiguration struct {

        
            IsDefault bool  `json:"is_default"`
            Data AppCatalogConfiguration  `json:"data"`
         
    }
    
    // GetCatalogConfigurationDetailsSchemaListing used by Catalog
    type GetCatalogConfigurationDetailsSchemaListing struct {

        
            Sort map[string]interface{}  `json:"sort"`
            Filter map[string]interface{}  `json:"filter"`
         
    }
    
    // EntityConfiguration used by Catalog
    type EntityConfiguration struct {

        
            ConfigID string  `json:"config_id"`
            ConfigType string  `json:"config_type"`
            AppID string  `json:"app_id"`
            Product GetCatalogConfigurationDetailsProduct  `json:"product"`
            ID string  `json:"id"`
            Listing GetCatalogConfigurationDetailsSchemaListing  `json:"listing"`
         
    }
    
    // GetAppCatalogEntityConfiguration used by Catalog
    type GetAppCatalogEntityConfiguration struct {

        
            IsDefault bool  `json:"is_default"`
            Data EntityConfiguration  `json:"data"`
         
    }
    
    // ProductSortOn used by Catalog
    type ProductSortOn struct {

        
            Name string  `json:"name"`
            Value string  `json:"value"`
            IsSelected bool  `json:"is_selected"`
         
    }
    
    // ProductFiltersValue used by Catalog
    type ProductFiltersValue struct {

        
            DisplayFormat string  `json:"display_format"`
            Display string  `json:"display"`
            Count float64  `json:"count"`
            CurrencySymbol string  `json:"currency_symbol"`
            Value interface{}  `json:"value"`
            Min float64  `json:"min"`
            IsSelected bool  `json:"is_selected"`
            SelectedMin float64  `json:"selected_min"`
            Max float64  `json:"max"`
            SelectedMax float64  `json:"selected_max"`
            CurrencyCode string  `json:"currency_code"`
            QueryFormat string  `json:"query_format"`
         
    }
    
    // ProductFiltersKey used by Catalog
    type ProductFiltersKey struct {

        
            Display string  `json:"display"`
            Logo string  `json:"logo"`
            Operators []string  `json:"operators"`
            Name string  `json:"name"`
            Kind string  `json:"kind"`
         
    }
    
    // ProductFilters used by Catalog
    type ProductFilters struct {

        
            Values []ProductFiltersValue  `json:"values"`
            Key ProductFiltersKey  `json:"key"`
         
    }
    
    // GetCollectionQueryOptionResponse used by Catalog
    type GetCollectionQueryOptionResponse struct {

        
            SortOn []ProductSortOn  `json:"sort_on"`
            Filters []ProductFilters  `json:"filters"`
            Operators map[string]string  `json:"operators"`
         
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

        
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
            Email string  `json:"email"`
            UID string  `json:"uid"`
         
    }
    
    // CollectionBadge used by Catalog
    type CollectionBadge struct {

        
            Text string  `json:"text"`
            Color string  `json:"color"`
         
    }
    
    // NextSchedule used by Catalog
    type NextSchedule struct {

        
            Start string  `json:"start"`
            End string  `json:"end"`
         
    }
    
    // CollectionSchedule used by Catalog
    type CollectionSchedule struct {

        
            Duration float64  `json:"duration"`
            Start string  `json:"start"`
            NextSchedule []NextSchedule  `json:"next_schedule"`
            End string  `json:"end"`
            Cron string  `json:"cron"`
         
    }
    
    // SeoDetail used by Catalog
    type SeoDetail struct {

        
            Title string  `json:"title"`
            Description string  `json:"description"`
         
    }
    
    // CollectionQuery used by Catalog
    type CollectionQuery struct {

        
            Value []interface{}  `json:"value"`
            Op string  `json:"op"`
            Attribute string  `json:"attribute"`
         
    }
    
    // CreateCollection used by Catalog
    type CreateCollection struct {

        
            Slug string  `json:"slug"`
            Logo CollectionImage  `json:"logo"`
            Published bool  `json:"published"`
            Banners CollectionBanner  `json:"banners"`
            CreatedBy UserInfo  `json:"created_by"`
            Badge CollectionBadge  `json:"badge"`
            Schedule CollectionSchedule  `json:"_schedule"`
            AllowFacets bool  `json:"allow_facets"`
            Type string  `json:"type"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Meta map[string]interface{}  `json:"meta"`
            SortOn string  `json:"sort_on"`
            Seo SeoDetail  `json:"seo"`
            Tags []string  `json:"tags"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            Name string  `json:"name"`
            AllowSort bool  `json:"allow_sort"`
            IsVisible bool  `json:"is_visible"`
            AppID string  `json:"app_id"`
            Priority float64  `json:"priority"`
            Query []CollectionQuery  `json:"query"`
            ModifiedBy UserInfo  `json:"modified_by"`
            Description string  `json:"description"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // BannerImage used by Catalog
    type BannerImage struct {

        
            URL string  `json:"url"`
            AspectRatio string  `json:"aspect_ratio"`
            Type string  `json:"type"`
         
    }
    
    // ImageUrls used by Catalog
    type ImageUrls struct {

        
            Landscape BannerImage  `json:"landscape"`
            Portrait BannerImage  `json:"portrait"`
         
    }
    
    // CollectionCreateResponse used by Catalog
    type CollectionCreateResponse struct {

        
            Slug string  `json:"slug"`
            Logo BannerImage  `json:"logo"`
            Banners ImageUrls  `json:"banners"`
            Badge map[string]interface{}  `json:"badge"`
            Schedule map[string]interface{}  `json:"_schedule"`
            AllowFacets bool  `json:"allow_facets"`
            Type string  `json:"type"`
            Meta map[string]interface{}  `json:"meta"`
            SortOn string  `json:"sort_on"`
            Cron map[string]interface{}  `json:"cron"`
            Name string  `json:"name"`
            Tag []string  `json:"tag"`
            AllowSort bool  `json:"allow_sort"`
            AppID string  `json:"app_id"`
            Priority float64  `json:"priority"`
            Query []CollectionQuery  `json:"query"`
            Description string  `json:"description"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // CollectionActionPageQuery used by Catalog
    type CollectionActionPageQuery struct {

        
            Value []string  `json:"value"`
            Op string  `json:"op"`
            Attribute string  `json:"attribute"`
         
    }
    
    // CollectionActionPage used by Catalog
    type CollectionActionPage struct {

        
            Query CollectionActionPageQuery  `json:"query"`
            Type string  `json:"type"`
         
    }
    
    // Media used by Catalog
    type Media struct {

        
            URL string  `json:"url"`
            Type string  `json:"type"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // CollectionBadge1 used by Catalog
    type CollectionBadge1 struct {

        
            Text string  `json:"text"`
            Color string  `json:"color"`
         
    }
    
    // CollectionScheduleStartEnd used by Catalog
    type CollectionScheduleStartEnd struct {

        
            Start string  `json:"start"`
            End string  `json:"end"`
         
    }
    
    // CollectionSchedule1 used by Catalog
    type CollectionSchedule1 struct {

        
            NextSchedule []CollectionScheduleStartEnd  `json:"next_schedule"`
            Start string  `json:"start"`
            End string  `json:"end"`
         
    }
    
    // GetCollectionDetailNest used by Catalog
    type GetCollectionDetailNest struct {

        
            Action CollectionActionPage  `json:"action"`
            Slug string  `json:"slug"`
            Logo Media  `json:"logo"`
            UID string  `json:"uid"`
            Banners ImageUrls  `json:"banners"`
            Badge CollectionBadge1  `json:"badge"`
            Schedule CollectionSchedule1  `json:"_schedule"`
            AllowFacets bool  `json:"allow_facets"`
            Type string  `json:"type"`
            Meta map[string]interface{}  `json:"meta"`
            Tags []string  `json:"tags"`
            Name string  `json:"name"`
            AllowSort bool  `json:"allow_sort"`
            AppID string  `json:"app_id"`
            Priority float64  `json:"priority"`
            Query []CollectionQuery  `json:"query"`
            Description string  `json:"description"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            IsActive bool  `json:"is_active"`
         
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
    
    // CollectionDetailResponse used by Catalog
    type CollectionDetailResponse struct {

        
            AllowFacets bool  `json:"allow_facets"`
            Slug string  `json:"slug"`
            Type string  `json:"type"`
            AppID string  `json:"app_id"`
            Priority float64  `json:"priority"`
            Logo Media  `json:"logo"`
            Query []CollectionQuery  `json:"query"`
            Tags []string  `json:"tags"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            UID string  `json:"uid"`
            Banners ImageUrls  `json:"banners"`
            Name string  `json:"name"`
            Badge CollectionBadge1  `json:"badge"`
            Description string  `json:"description"`
            Meta map[string]interface{}  `json:"meta"`
            Schedule CollectionSchedule1  `json:"_schedule"`
            AllowSort bool  `json:"allow_sort"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // UpdateCollection used by Catalog
    type UpdateCollection struct {

        
            Slug string  `json:"slug"`
            Logo CollectionImage  `json:"logo"`
            Published bool  `json:"published"`
            Banners CollectionBanner  `json:"banners"`
            Badge CollectionBadge  `json:"badge"`
            Schedule CollectionSchedule  `json:"_schedule"`
            AllowFacets bool  `json:"allow_facets"`
            Type string  `json:"type"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Meta map[string]interface{}  `json:"meta"`
            SortOn string  `json:"sort_on"`
            Seo SeoDetail  `json:"seo"`
            Tags []string  `json:"tags"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            Name string  `json:"name"`
            AllowSort bool  `json:"allow_sort"`
            IsVisible bool  `json:"is_visible"`
            Priority float64  `json:"priority"`
            Query []CollectionQuery  `json:"query"`
            ModifiedBy UserInfo  `json:"modified_by"`
            Description string  `json:"description"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // CollectionItem used by Catalog
    type CollectionItem struct {

        
            Action string  `json:"action"`
            ItemID float64  `json:"item_id"`
            Priority float64  `json:"priority"`
         
    }
    
    // CollectionItemUpdate used by Catalog
    type CollectionItemUpdate struct {

        
            AllowFacets bool  `json:"allow_facets"`
            Type string  `json:"type"`
            Items []CollectionItem  `json:"items"`
            Query []CollectionQuery  `json:"query"`
            AllowSort bool  `json:"allow_sort"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
         
    }
    
    // UpdatedResponse used by Catalog
    type UpdatedResponse struct {

        
            Message string  `json:"message"`
            ItemsNotUpdated []float64  `json:"items_not_updated"`
         
    }
    
    // ProductDetailAttribute used by Catalog
    type ProductDetailAttribute struct {

        
            Value string  `json:"value"`
            Type string  `json:"type"`
            Key string  `json:"key"`
         
    }
    
    // ProductDetailGroupedAttribute used by Catalog
    type ProductDetailGroupedAttribute struct {

        
            Title string  `json:"title"`
            Details []ProductDetailAttribute  `json:"details"`
         
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
            Logo Media  `json:"logo"`
            UID float64  `json:"uid"`
            Action Action  `json:"action"`
         
    }
    
    // Price1 used by Catalog
    type Price1 struct {

        
            Max float64  `json:"max"`
            CurrencySymbol string  `json:"currency_symbol"`
            CurrencyCode string  `json:"currency_code"`
            Min float64  `json:"min"`
         
    }
    
    // ProductListingPrice used by Catalog
    type ProductListingPrice struct {

        
            Marked Price1  `json:"marked"`
            Effective Price1  `json:"effective"`
         
    }
    
    // ProductListingDetail used by Catalog
    type ProductListingDetail struct {

        
            Slug string  `json:"slug"`
            Highlights []string  `json:"highlights"`
            UID float64  `json:"uid"`
            RatingCount float64  `json:"rating_count"`
            ShortDescription string  `json:"short_description"`
            ProductOnlineDate string  `json:"product_online_date"`
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
            Type string  `json:"type"`
            ImageNature string  `json:"image_nature"`
            HasVariant bool  `json:"has_variant"`
            Sellable bool  `json:"sellable"`
            Color string  `json:"color"`
            Similars []string  `json:"similars"`
            Discount string  `json:"discount"`
            Brand ProductBrand  `json:"brand"`
            Tryouts []string  `json:"tryouts"`
            Price ProductListingPrice  `json:"price"`
            Name string  `json:"name"`
            PromoMeta map[string]interface{}  `json:"promo_meta"`
            ItemCode string  `json:"item_code"`
            Attributes map[string]interface{}  `json:"attributes"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            ItemType string  `json:"item_type"`
            Description string  `json:"description"`
            Medias []Media  `json:"medias"`
            Rating float64  `json:"rating"`
         
    }
    
    // GetCollectionItemsResponse used by Catalog
    type GetCollectionItemsResponse struct {

        
            SortOn []ProductSortOn  `json:"sort_on"`
            Items []ProductListingDetail  `json:"items"`
            Filters []ProductFilters  `json:"filters"`
            Page Page  `json:"page"`
         
    }
    
    // CatalogInsightItem used by Catalog
    type CatalogInsightItem struct {

        
            Count float64  `json:"count"`
            OutOfStockCount float64  `json:"out_of_stock_count"`
            SellableCount float64  `json:"sellable_count"`
         
    }
    
    // CatalogInsightBrand used by Catalog
    type CatalogInsightBrand struct {

        
            TotalArticles float64  `json:"total_articles"`
            AvailableSizes float64  `json:"available_sizes"`
            Name string  `json:"name"`
            AvailableArticles float64  `json:"available_articles"`
            ArticleFreshness float64  `json:"article_freshness"`
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

        
            CompanyID float64  `json:"company_id"`
            BrandIds []float64  `json:"brand_ids"`
            Enabled bool  `json:"enabled"`
            StoreIds []float64  `json:"store_ids"`
            Platform string  `json:"platform"`
            OptLevel string  `json:"opt_level"`
         
    }
    
    // CompanyOptIn used by Catalog
    type CompanyOptIn struct {

        
            CompanyID float64  `json:"company_id"`
            BrandIds []float64  `json:"brand_ids"`
            CreatedOn float64  `json:"created_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Enabled bool  `json:"enabled"`
            StoreIds []float64  `json:"store_ids"`
            Platform string  `json:"platform"`
            OptLevel string  `json:"opt_level"`
            ModifiedOn float64  `json:"modified_on"`
         
    }
    
    // GetOptInPlatform used by Catalog
    type GetOptInPlatform struct {

        
            Items []CompanyOptIn  `json:"items"`
            Page Page  `json:"page"`
         
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

        
            CompanyID float64  `json:"company_id"`
            BrandName string  `json:"brand_name"`
            BrandID float64  `json:"brand_id"`
            TotalArticle float64  `json:"total_article"`
         
    }
    
    // OptinCompanyBrandDetailsView used by Catalog
    type OptinCompanyBrandDetailsView struct {

        
            Items []CompanyBrandDetail  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // OptinCompanyMetrics used by Catalog
    type OptinCompanyMetrics struct {

        
            Brand float64  `json:"brand"`
            Company string  `json:"company"`
            Store float64  `json:"store"`
         
    }
    
    // StoreDetail used by Catalog
    type StoreDetail struct {

        
            CompanyID float64  `json:"company_id"`
            Documents []map[string]interface{}  `json:"documents"`
            DisplayName string  `json:"display_name"`
            CreatedOn string  `json:"created_on"`
            AdditionalContacts []map[string]interface{}  `json:"additional_contacts"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            Timing map[string]interface{}  `json:"timing"`
            Address map[string]interface{}  `json:"address"`
            Manager map[string]interface{}  `json:"manager"`
            StoreCode string  `json:"store_code"`
            ModifiedOn string  `json:"modified_on"`
            StoreType string  `json:"store_type"`
         
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

        
            Multi bool  `json:"multi"`
            Type string  `json:"type"`
            Range AttributeSchemaRange  `json:"range"`
            Mandatory bool  `json:"mandatory"`
            AllowedValues []string  `json:"allowed_values"`
            Format string  `json:"format"`
         
    }
    
    // AttributeMasterDetails used by Catalog
    type AttributeMasterDetails struct {

        
            DisplayType string  `json:"display_type"`
         
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

        
            MandatoryDetails AttributeMasterMandatoryDetails  `json:"mandatory_details"`
            Enriched bool  `json:"enriched"`
         
    }
    
    // GenderDetail used by Catalog
    type GenderDetail struct {

        
            Schema AttributeMaster  `json:"schema"`
            Departments []string  `json:"departments"`
            Slug string  `json:"slug"`
            EnabledForEndConsumer bool  `json:"enabled_for_end_consumer"`
            Logo string  `json:"logo"`
            Details AttributeMasterDetails  `json:"details"`
            ID string  `json:"id"`
            Name string  `json:"name"`
            Filters AttributeMasterFilter  `json:"filters"`
            Description string  `json:"description"`
            Meta AttributeMasterMeta  `json:"meta"`
            IsNested bool  `json:"is_nested"`
         
    }
    
    // CategoriesResponse used by Catalog
    type CategoriesResponse struct {

        
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            TemplateSlug string  `json:"template_slug"`
            SlugKey string  `json:"slug_key"`
         
    }
    
    // ProdcutTemplateCategoriesResponse used by Catalog
    type ProdcutTemplateCategoriesResponse struct {

        
            Items []CategoriesResponse  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // PTErrorResponse used by Catalog
    type PTErrorResponse struct {

        
            Status float64  `json:"status"`
            Errors map[string]interface{}  `json:"errors"`
            Message string  `json:"message"`
            Meta map[string]interface{}  `json:"meta"`
            Code string  `json:"code"`
         
    }
    
    // DepartmentCreateUpdate used by Catalog
    type DepartmentCreateUpdate struct {

        
            PriorityOrder float64  `json:"priority_order"`
            Slug string  `json:"slug"`
            Logo string  `json:"logo"`
            Tags []string  `json:"tags"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            Synonyms []string  `json:"synonyms"`
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

        
            Contact string  `json:"contact"`
            UID string  `json:"uid"`
            ID string  `json:"_id"`
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
         
    }
    
    // GetDepartment used by Catalog
    type GetDepartment struct {

        
            PriorityOrder float64  `json:"priority_order"`
            Slug string  `json:"slug"`
            PageSize float64  `json:"page_size"`
            Logo string  `json:"logo"`
            CreatedOn string  `json:"created_on"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            UID float64  `json:"uid"`
            Search string  `json:"search"`
            PageNo float64  `json:"page_no"`
            CreatedBy UserSerializer  `json:"created_by"`
            Name string  `json:"name"`
            Synonyms []string  `json:"synonyms"`
            ItemType string  `json:"item_type"`
            ModifiedOn string  `json:"modified_on"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // DepartmentsResponse used by Catalog
    type DepartmentsResponse struct {

        
            Items []GetDepartment  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // DepartmentErrorResponse used by Catalog
    type DepartmentErrorResponse struct {

        
            Status float64  `json:"status"`
            Errors map[string]interface{}  `json:"errors"`
            Message string  `json:"message"`
            Meta map[string]interface{}  `json:"meta"`
            Code string  `json:"code"`
         
    }
    
    // UserDetail used by Catalog
    type UserDetail struct {

        
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
            SuperUser bool  `json:"super_user"`
            Contact string  `json:"contact"`
         
    }
    
    // DepartmentModel used by Catalog
    type DepartmentModel struct {

        
            PriorityOrder float64  `json:"priority_order"`
            Slug string  `json:"slug"`
            CreatedOn string  `json:"created_on"`
            Logo string  `json:"logo"`
            ModifiedBy UserDetail  `json:"modified_by"`
            UID float64  `json:"uid"`
            ID interface{}  `json:"_id"`
            IsActive bool  `json:"is_active"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            CreatedBy UserDetail  `json:"created_by"`
            Name string  `json:"name"`
            Synonyms []string  `json:"synonyms"`
            VerifiedBy UserDetail  `json:"verified_by"`
            Cls string  `json:"_cls"`
            ModifiedOn string  `json:"modified_on"`
            VerifiedOn string  `json:"verified_on"`
         
    }
    
    // ProductTemplate used by Catalog
    type ProductTemplate struct {

        
            Attributes []string  `json:"attributes"`
            Departments []string  `json:"departments"`
            Slug string  `json:"slug"`
            Categories []string  `json:"categories"`
            Logo string  `json:"logo"`
            CreatedOn string  `json:"created_on"`
            IsArchived bool  `json:"is_archived"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Name string  `json:"name"`
            Tag string  `json:"tag"`
            Description string  `json:"description"`
            IsExpirable bool  `json:"is_expirable"`
            IsPhysical bool  `json:"is_physical"`
            ModifiedOn string  `json:"modified_on"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // TemplatesResponse used by Catalog
    type TemplatesResponse struct {

        
            Items ProductTemplate  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // TemplateDetails used by Catalog
    type TemplateDetails struct {

        
            Attributes []string  `json:"attributes"`
            Departments []string  `json:"departments"`
            Slug string  `json:"slug"`
            Categories []string  `json:"categories"`
            Logo string  `json:"logo"`
            IsArchived bool  `json:"is_archived"`
            ID string  `json:"id"`
            Name string  `json:"name"`
            Tag string  `json:"tag"`
            Description string  `json:"description"`
            IsExpirable bool  `json:"is_expirable"`
            IsPhysical bool  `json:"is_physical"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // Properties used by Catalog
    type Properties struct {

        
            TraderType map[string]interface{}  `json:"trader_type"`
            Slug map[string]interface{}  `json:"slug"`
            ReturnConfig map[string]interface{}  `json:"return_config"`
            Highlights map[string]interface{}  `json:"highlights"`
            SizeGuide map[string]interface{}  `json:"size_guide"`
            Trader map[string]interface{}  `json:"trader"`
            ShortDescription map[string]interface{}  `json:"short_description"`
            IsDependent map[string]interface{}  `json:"is_dependent"`
            HsnCode map[string]interface{}  `json:"hsn_code"`
            Media map[string]interface{}  `json:"media"`
            ProductPublish map[string]interface{}  `json:"product_publish"`
            Command map[string]interface{}  `json:"command"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            Currency map[string]interface{}  `json:"currency"`
            MultiSize map[string]interface{}  `json:"multi_size"`
            Tags map[string]interface{}  `json:"tags"`
            Variants map[string]interface{}  `json:"variants"`
            CountryOfOrigin map[string]interface{}  `json:"country_of_origin"`
            Name map[string]interface{}  `json:"name"`
            CategorySlug map[string]interface{}  `json:"category_slug"`
            BrandUID map[string]interface{}  `json:"brand_uid"`
            ItemCode map[string]interface{}  `json:"item_code"`
            Sizes map[string]interface{}  `json:"sizes"`
            NoOfBoxes map[string]interface{}  `json:"no_of_boxes"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            ProductGroupTag map[string]interface{}  `json:"product_group_tag"`
            ItemType map[string]interface{}  `json:"item_type"`
            Description map[string]interface{}  `json:"description"`
            IsActive map[string]interface{}  `json:"is_active"`
         
    }
    
    // GlobalValidation used by Catalog
    type GlobalValidation struct {

        
            Definitions map[string]interface{}  `json:"definitions"`
            Properties Properties  `json:"properties"`
            Type string  `json:"type"`
            Required []string  `json:"required"`
            Title string  `json:"title"`
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

        
            Message string  `json:"message"`
            Data map[string]interface{}  `json:"data"`
         
    }
    
    // HSNData used by Catalog
    type HSNData struct {

        
            CountryOfOrigin []string  `json:"country_of_origin"`
            HsnCode []string  `json:"hsn_code"`
         
    }
    
    // HSNCodesResponse used by Catalog
    type HSNCodesResponse struct {

        
            Message string  `json:"message"`
            Data HSNData  `json:"data"`
         
    }
    
    // UserInfo1 used by Catalog
    type UserInfo1 struct {

        
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
            Email string  `json:"email"`
            UID string  `json:"uid"`
         
    }
    
    // ProductTemplateExportResponse used by Catalog
    type ProductTemplateExportResponse struct {

        
            CompletedOn string  `json:"completed_on"`
            NotificationEmails []string  `json:"notification_emails"`
            Type string  `json:"type"`
            Status string  `json:"status"`
            SellerID float64  `json:"seller_id"`
            CreatedBy UserInfo1  `json:"created_by"`
            TaskID string  `json:"task_id"`
            Filters map[string]interface{}  `json:"filters"`
            URL string  `json:"url"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // ProductDownloadsResponse used by Catalog
    type ProductDownloadsResponse struct {

        
            Items []ProductTemplateExportResponse  `json:"items"`
         
    }
    
    // ProductTemplateExportFilterRequest used by Catalog
    type ProductTemplateExportFilterRequest struct {

        
            ToDate string  `json:"to_date"`
            Templates []string  `json:"templates"`
            CatalogueTypes []string  `json:"catalogue_types"`
            FromDate string  `json:"from_date"`
            Brands []string  `json:"brands"`
         
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

        
            L2 float64  `json:"l2"`
            L1 float64  `json:"l1"`
            Department float64  `json:"department"`
         
    }
    
    // CategoryMappingValues used by Catalog
    type CategoryMappingValues struct {

        
            Name string  `json:"name"`
            CatalogID float64  `json:"catalog_id"`
         
    }
    
    // CategoryMapping used by Catalog
    type CategoryMapping struct {

        
            Google CategoryMappingValues  `json:"google"`
            Ajio CategoryMappingValues  `json:"ajio"`
            Facebook CategoryMappingValues  `json:"facebook"`
         
    }
    
    // Media1 used by Catalog
    type Media1 struct {

        
            Logo string  `json:"logo"`
            Landscape string  `json:"landscape"`
            Portrait string  `json:"portrait"`
         
    }
    
    // CategoryRequestBody used by Catalog
    type CategoryRequestBody struct {

        
            Departments []float64  `json:"departments"`
            Level float64  `json:"level"`
            Slug string  `json:"slug"`
            Priority float64  `json:"priority"`
            Tryouts []string  `json:"tryouts"`
            Hierarchy []Hierarchy  `json:"hierarchy"`
            Name string  `json:"name"`
            Synonyms []string  `json:"synonyms"`
            Marketplaces CategoryMapping  `json:"marketplaces"`
            Media Media1  `json:"media"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // CategoryCreateResponse used by Catalog
    type CategoryCreateResponse struct {

        
            UID float64  `json:"uid"`
            Message string  `json:"message"`
         
    }
    
    // Category used by Catalog
    type Category struct {

        
            Departments []float64  `json:"departments"`
            Level float64  `json:"level"`
            Slug string  `json:"slug"`
            Priority float64  `json:"priority"`
            Tryouts []string  `json:"tryouts"`
            CreatedOn string  `json:"created_on"`
            Hierarchy []Hierarchy  `json:"hierarchy"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            UID float64  `json:"uid"`
            ID string  `json:"id"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Name string  `json:"name"`
            Synonyms []string  `json:"synonyms"`
            Marketplaces CategoryMapping  `json:"marketplaces"`
            Media Media1  `json:"media"`
            ModifiedOn string  `json:"modified_on"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // CategoryResponse used by Catalog
    type CategoryResponse struct {

        
            Items []Category  `json:"items"`
            Page Page  `json:"page"`
         
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
    
    // ReturnConfig used by Catalog
    type ReturnConfig struct {

        
            Returnable bool  `json:"returnable"`
            Unit string  `json:"unit"`
            Time float64  `json:"time"`
         
    }
    
    // Trader used by Catalog
    type Trader struct {

        
            Name string  `json:"name"`
            Address []string  `json:"address"`
            Type string  `json:"type"`
         
    }
    
    // NetQuantity used by Catalog
    type NetQuantity struct {

        
            Unit string  `json:"unit"`
            Value float64  `json:"value"`
         
    }
    
    // TaxIdentifier used by Catalog
    type TaxIdentifier struct {

        
            HsnCodeID string  `json:"hsn_code_id"`
            ReportingHsn string  `json:"reporting_hsn"`
            HsnCode string  `json:"hsn_code"`
         
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

        
            URL string  `json:"url"`
            Tag string  `json:"tag"`
         
    }
    
    // ProductCreateUpdateSchemaV2 used by Catalog
    type ProductCreateUpdateSchemaV2 struct {

        
            Action string  `json:"action"`
            Slug string  `json:"slug"`
            ReturnConfig ReturnConfig  `json:"return_config"`
            Highlights []string  `json:"highlights"`
            SizeGuide string  `json:"size_guide"`
            Trader []Trader  `json:"trader"`
            UID float64  `json:"uid"`
            ShortDescription string  `json:"short_description"`
            TemplateTag string  `json:"template_tag"`
            NetQuantity NetQuantity  `json:"net_quantity"`
            Requester string  `json:"requester"`
            IsDependent bool  `json:"is_dependent"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            TaxIdentifier TaxIdentifier  `json:"tax_identifier"`
            Media []Media  `json:"media"`
            BulkJobID string  `json:"bulk_job_id"`
            ProductPublish ProductPublish  `json:"product_publish"`
            VariantGroup map[string]interface{}  `json:"variant_group"`
            CustomOrder CustomOrder  `json:"custom_order"`
            Currency string  `json:"currency"`
            CompanyID float64  `json:"company_id"`
            Departments []float64  `json:"departments"`
            MultiSize bool  `json:"multi_size"`
            Tags []string  `json:"tags"`
            Variants map[string]interface{}  `json:"variants"`
            CountryOfOrigin string  `json:"country_of_origin"`
            Name string  `json:"name"`
            ChangeRequestID interface{}  `json:"change_request_id"`
            CategorySlug string  `json:"category_slug"`
            BrandUID float64  `json:"brand_uid"`
            ItemCode string  `json:"item_code"`
            Attributes map[string]interface{}  `json:"attributes"`
            IsSet bool  `json:"is_set"`
            Sizes []map[string]interface{}  `json:"sizes"`
            NoOfBoxes float64  `json:"no_of_boxes"`
            VariantMedia map[string]interface{}  `json:"variant_media"`
            TeaserTag TeaserTag  `json:"teaser_tag"`
            ProductGroupTag []string  `json:"product_group_tag"`
            ItemType string  `json:"item_type"`
            Description string  `json:"description"`
            IsActive bool  `json:"is_active"`
            IsImageLessProduct bool  `json:"is_image_less_product"`
         
    }
    
    // ReturnConfigResponse used by Catalog
    type ReturnConfigResponse struct {

        
            Returnable bool  `json:"returnable"`
            Unit string  `json:"unit"`
            Time float64  `json:"time"`
         
    }
    
    // VerifiedBy used by Catalog
    type VerifiedBy struct {

        
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
         
    }
    
    // NetQuantityResponse used by Catalog
    type NetQuantityResponse struct {

        
            Unit string  `json:"unit"`
            Value float64  `json:"value"`
         
    }
    
    // ProductPublish1 used by Catalog
    type ProductPublish1 struct {

        
            ProductOnlineDate string  `json:"product_online_date"`
            IsSet bool  `json:"is_set"`
         
    }
    
    // Logo used by Catalog
    type Logo struct {

        
            URL string  `json:"url"`
            AspectRatio string  `json:"aspect_ratio"`
            SecureURL string  `json:"secure_url"`
            AspectRatioF float64  `json:"aspect_ratio_f"`
         
    }
    
    // Brand used by Catalog
    type Brand struct {

        
            Name string  `json:"name"`
            Logo Logo  `json:"logo"`
            UID float64  `json:"uid"`
         
    }
    
    // Image used by Catalog
    type Image struct {

        
            SecureURL string  `json:"secure_url"`
            AspectRatio string  `json:"aspect_ratio"`
            URL string  `json:"url"`
            AspectRatioF float64  `json:"aspect_ratio_f"`
         
    }
    
    // ProductSchemaV2 used by Catalog
    type ProductSchemaV2 struct {

        
            Slug string  `json:"slug"`
            ReturnConfig ReturnConfigResponse  `json:"return_config"`
            Highlights []string  `json:"highlights"`
            SizeGuide string  `json:"size_guide"`
            Trader []Trader  `json:"trader"`
            UID float64  `json:"uid"`
            ShortDescription string  `json:"short_description"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            L3Mapping []string  `json:"l3_mapping"`
            VerifiedBy VerifiedBy  `json:"verified_by"`
            TemplateTag string  `json:"template_tag"`
            Category map[string]interface{}  `json:"category"`
            NetQuantity NetQuantityResponse  `json:"net_quantity"`
            IsDependent bool  `json:"is_dependent"`
            Moq map[string]interface{}  `json:"moq"`
            Stage string  `json:"stage"`
            CreatedOn string  `json:"created_on"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            HsnCode string  `json:"hsn_code"`
            ImageNature string  `json:"image_nature"`
            Media []Media  `json:"media"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            Color string  `json:"color"`
            ProductPublish ProductPublish1  `json:"product_publish"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            VariantGroup map[string]interface{}  `json:"variant_group"`
            VerifiedOn string  `json:"verified_on"`
            Currency string  `json:"currency"`
            CompanyID float64  `json:"company_id"`
            Brand Brand  `json:"brand"`
            Departments []float64  `json:"departments"`
            MultiSize bool  `json:"multi_size"`
            Tags []string  `json:"tags"`
            Variants map[string]interface{}  `json:"variants"`
            CountryOfOrigin string  `json:"country_of_origin"`
            Name string  `json:"name"`
            AllSizes []map[string]interface{}  `json:"all_sizes"`
            Images []Image  `json:"images"`
            CategoryUID float64  `json:"category_uid"`
            AllCompanyIds []float64  `json:"all_company_ids"`
            CategorySlug string  `json:"category_slug"`
            BrandUID float64  `json:"brand_uid"`
            ModifiedOn string  `json:"modified_on"`
            AllIdentifiers []string  `json:"all_identifiers"`
            ItemCode string  `json:"item_code"`
            Attributes map[string]interface{}  `json:"attributes"`
            PrimaryColor string  `json:"primary_color"`
            Pending string  `json:"pending"`
            IsSet bool  `json:"is_set"`
            Sizes []map[string]interface{}  `json:"sizes"`
            NoOfBoxes float64  `json:"no_of_boxes"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            VariantMedia map[string]interface{}  `json:"variant_media"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            ID string  `json:"id"`
            ProductGroupTag []string  `json:"product_group_tag"`
            ItemType string  `json:"item_type"`
            Description string  `json:"description"`
            IsExpirable bool  `json:"is_expirable"`
            IsPhysical bool  `json:"is_physical"`
            IsActive bool  `json:"is_active"`
            IsImageLessProduct bool  `json:"is_image_less_product"`
         
    }
    
    // ProductListingResponseV2 used by Catalog
    type ProductListingResponseV2 struct {

        
            Items []ProductSchemaV2  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // ProductVariants used by Catalog
    type ProductVariants struct {

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            CategoryUID float64  `json:"category_uid"`
            Media []Media  `json:"media"`
            BrandUID float64  `json:"brand_uid"`
            ItemCode string  `json:"item_code"`
         
    }
    
    // ProductVariantsResponse used by Catalog
    type ProductVariantsResponse struct {

        
            Variants []ProductVariants  `json:"variants"`
            Page Page  `json:"page"`
         
    }
    
    // AttributeMasterSerializer used by Catalog
    type AttributeMasterSerializer struct {

        
            Schema AttributeMaster  `json:"schema"`
            Slug string  `json:"slug"`
            Logo string  `json:"logo"`
            Details AttributeMasterDetails  `json:"details"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Filters AttributeMasterFilter  `json:"filters"`
            CreatedOn string  `json:"created_on"`
            Suggestion string  `json:"suggestion"`
            Variant bool  `json:"variant"`
            Departments []string  `json:"departments"`
            Tags []string  `json:"tags"`
            Name string  `json:"name"`
            Unit string  `json:"unit"`
            ModifiedOn string  `json:"modified_on"`
            IsNested bool  `json:"is_nested"`
            EnabledForEndConsumer bool  `json:"enabled_for_end_consumer"`
            Example string  `json:"example"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            RawKey string  `json:"raw_key"`
            Synonyms map[string]interface{}  `json:"synonyms"`
            Description string  `json:"description"`
         
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
            GtinValue string  `json:"gtin_value"`
            GtinType string  `json:"gtin_type"`
         
    }
    
    // AllSizes used by Catalog
    type AllSizes struct {

        
            ItemWidth float64  `json:"item_width"`
            ItemDimensionsUnitOfMeasure string  `json:"item_dimensions_unit_of_measure"`
            Identifiers []ValidateIdentifier  `json:"identifiers"`
            ItemHeight float64  `json:"item_height"`
            ItemLength float64  `json:"item_length"`
            ItemWeight float64  `json:"item_weight"`
            Size string  `json:"size"`
            ItemWeightUnitOfMeasure string  `json:"item_weight_unit_of_measure"`
         
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

        
            Slug string  `json:"slug"`
            ReturnConfig ReturnConfigResponse  `json:"return_config"`
            Highlights []string  `json:"highlights"`
            SizeGuide string  `json:"size_guide"`
            Trader []Trader  `json:"trader"`
            UID float64  `json:"uid"`
            ShortDescription string  `json:"short_description"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            L3Mapping []string  `json:"l3_mapping"`
            VerifiedBy VerifiedBy  `json:"verified_by"`
            TemplateTag string  `json:"template_tag"`
            Category map[string]interface{}  `json:"category"`
            NetQuantity NetQuantityResponse  `json:"net_quantity"`
            IsDependent bool  `json:"is_dependent"`
            Moq map[string]interface{}  `json:"moq"`
            Stage string  `json:"stage"`
            CreatedOn string  `json:"created_on"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            HsnCode string  `json:"hsn_code"`
            ImageNature string  `json:"image_nature"`
            Media []Media  `json:"media"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            Color string  `json:"color"`
            ProductPublish ProductPublished  `json:"product_publish"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            VariantGroup map[string]interface{}  `json:"variant_group"`
            VerifiedOn string  `json:"verified_on"`
            Currency string  `json:"currency"`
            CompanyID float64  `json:"company_id"`
            Brand Brand  `json:"brand"`
            Departments []float64  `json:"departments"`
            MultiSize bool  `json:"multi_size"`
            Tags []string  `json:"tags"`
            Variants map[string]interface{}  `json:"variants"`
            CountryOfOrigin string  `json:"country_of_origin"`
            Name string  `json:"name"`
            AllSizes []map[string]interface{}  `json:"all_sizes"`
            Images []Image  `json:"images"`
            CategoryUID float64  `json:"category_uid"`
            AllCompanyIds []float64  `json:"all_company_ids"`
            CategorySlug string  `json:"category_slug"`
            BrandUID float64  `json:"brand_uid"`
            ModifiedOn string  `json:"modified_on"`
            AllIdentifiers []string  `json:"all_identifiers"`
            ItemCode string  `json:"item_code"`
            Attributes map[string]interface{}  `json:"attributes"`
            PrimaryColor string  `json:"primary_color"`
            Pending string  `json:"pending"`
            IsSet bool  `json:"is_set"`
            Sizes []map[string]interface{}  `json:"sizes"`
            NoOfBoxes float64  `json:"no_of_boxes"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            VariantMedia map[string]interface{}  `json:"variant_media"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            ID string  `json:"id"`
            ProductGroupTag []string  `json:"product_group_tag"`
            ItemType string  `json:"item_type"`
            Description string  `json:"description"`
            IsExpirable bool  `json:"is_expirable"`
            IsPhysical bool  `json:"is_physical"`
            IsActive bool  `json:"is_active"`
            IsImageLessProduct bool  `json:"is_image_less_product"`
         
    }
    
    // ProductListingResponse used by Catalog
    type ProductListingResponse struct {

        
            Items []Product  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // BulkJob used by Catalog
    type BulkJob struct {

        
            CompanyID float64  `json:"company_id"`
            TrackingURL string  `json:"tracking_url"`
            FailedRecords []map[string]interface{}  `json:"failed_records"`
            Stage string  `json:"stage"`
            CreatedOn string  `json:"created_on"`
            Failed float64  `json:"failed"`
            ModifiedBy UserInfo1  `json:"modified_by"`
            Cancelled float64  `json:"cancelled"`
            CustomTemplateTag string  `json:"custom_template_tag"`
            CreatedBy UserInfo1  `json:"created_by"`
            CancelledRecords []map[string]interface{}  `json:"cancelled_records"`
            FilePath string  `json:"file_path"`
            TemplateTag string  `json:"template_tag"`
            Total float64  `json:"total"`
            ModifiedOn string  `json:"modified_on"`
            IsActive bool  `json:"is_active"`
            Succeed float64  `json:"succeed"`
         
    }
    
    // BulkResponse used by Catalog
    type BulkResponse struct {

        
            CreatedOn string  `json:"created_on"`
            ModifiedBy UserInfo1  `json:"modified_by"`
            CreatedBy UserInfo1  `json:"created_by"`
            ModifiedOn string  `json:"modified_on"`
            IsActive bool  `json:"is_active"`
            BatchID string  `json:"batch_id"`
         
    }
    
    // UserDetail1 used by Catalog
    type UserDetail1 struct {

        
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
            FullName string  `json:"full_name"`
         
    }
    
    // ProductBulkRequest used by Catalog
    type ProductBulkRequest struct {

        
            CompanyID float64  `json:"company_id"`
            Stage string  `json:"stage"`
            FailedRecords []string  `json:"failed_records"`
            Failed float64  `json:"failed"`
            CreatedOn string  `json:"created_on"`
            Cancelled float64  `json:"cancelled"`
            ModifiedBy UserDetail1  `json:"modified_by"`
            CancelledRecords []string  `json:"cancelled_records"`
            CreatedBy UserDetail1  `json:"created_by"`
            FilePath string  `json:"file_path"`
            TemplateTag string  `json:"template_tag"`
            Template ProductTemplate  `json:"template"`
            Total float64  `json:"total"`
            ModifiedOn string  `json:"modified_on"`
            IsActive bool  `json:"is_active"`
            Succeed float64  `json:"succeed"`
         
    }
    
    // ProductBulkRequestList used by Catalog
    type ProductBulkRequestList struct {

        
            Items ProductBulkRequest  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // BulkProductRequest used by Catalog
    type BulkProductRequest struct {

        
            CompanyID float64  `json:"company_id"`
            TemplateTag string  `json:"template_tag"`
            Data []map[string]interface{}  `json:"data"`
            BatchID string  `json:"batch_id"`
         
    }
    
    // ProductTagsViewResponse used by Catalog
    type ProductTagsViewResponse struct {

        
            Items []string  `json:"items"`
         
    }
    
    // ProductBulkAssets used by Catalog
    type ProductBulkAssets struct {

        
            URL string  `json:"url"`
            User map[string]interface{}  `json:"user"`
            BatchID string  `json:"batch_id"`
            CompanyID float64  `json:"company_id"`
         
    }
    
    // UserCommon used by Catalog
    type UserCommon struct {

        
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
            CompanyID float64  `json:"company_id"`
         
    }
    
    // Items used by Catalog
    type Items struct {

        
            CompanyID float64  `json:"company_id"`
            Retry float64  `json:"retry"`
            TrackingURL string  `json:"tracking_url"`
            FailedRecords []string  `json:"failed_records"`
            Stage string  `json:"stage"`
            CreatedOn string  `json:"created_on"`
            Failed float64  `json:"failed"`
            ModifiedBy UserCommon  `json:"modified_by"`
            Cancelled float64  `json:"cancelled"`
            ID string  `json:"id"`
            CreatedBy UserCommon  `json:"created_by"`
            FilePath string  `json:"file_path"`
            CancelledRecords []string  `json:"cancelled_records"`
            Total float64  `json:"total"`
            ModifiedOn string  `json:"modified_on"`
            IsActive bool  `json:"is_active"`
            Succeed float64  `json:"succeed"`
         
    }
    
    // BulkAssetResponse used by Catalog
    type BulkAssetResponse struct {

        
            Items []Items  `json:"items"`
            Page Page  `json:"page"`
         
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
    
    // ItemQuery used by Catalog
    type ItemQuery struct {

        
            BrandUID float64  `json:"brand_uid"`
            UID float64  `json:"uid"`
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
            SizeDistribution SizeDistribution  `json:"size_distribution"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // GTIN used by Catalog
    type GTIN struct {

        
            Primary bool  `json:"primary"`
            GtinValue string  `json:"gtin_value"`
            GtinType string  `json:"gtin_type"`
         
    }
    
    // InvSize used by Catalog
    type InvSize struct {

        
            ExpirationDate string  `json:"expiration_date"`
            IsSet bool  `json:"is_set"`
            PriceEffective float64  `json:"price_effective"`
            StoreCode string  `json:"store_code"`
            ItemWidth float64  `json:"item_width"`
            Set InventorySet  `json:"set"`
            Price float64  `json:"price"`
            ItemDimensionsUnitOfMeasure string  `json:"item_dimensions_unit_of_measure"`
            PriceTransfer float64  `json:"price_transfer"`
            Identifiers []GTIN  `json:"identifiers"`
            ItemHeight float64  `json:"item_height"`
            ItemLength float64  `json:"item_length"`
            ItemWeight float64  `json:"item_weight"`
            Size string  `json:"size"`
            ItemWeightUnitOfMeasure string  `json:"item_weight_unit_of_measure"`
            Quantity float64  `json:"quantity"`
            Currency string  `json:"currency"`
         
    }
    
    // InventoryRequest used by Catalog
    type InventoryRequest struct {

        
            CompanyID float64  `json:"company_id"`
            Item ItemQuery  `json:"item"`
            Sizes []InvSize  `json:"sizes"`
         
    }
    
    // InventoryResponse used by Catalog
    type InventoryResponse struct {

        
            Store map[string]interface{}  `json:"store"`
            PriceEffective float64  `json:"price_effective"`
            Quantity float64  `json:"quantity"`
            Price float64  `json:"price"`
            UID string  `json:"uid"`
            PriceTransfer float64  `json:"price_transfer"`
            Identifiers map[string]interface{}  `json:"identifiers"`
            SellableQuantity float64  `json:"sellable_quantity"`
            SellerIdentifier string  `json:"seller_identifier"`
            ItemID float64  `json:"item_id"`
            Size string  `json:"size"`
            InventoryUpdatedOn string  `json:"inventory_updated_on"`
            Currency string  `json:"currency"`
         
    }
    
    // InventoryResponsePaginated used by Catalog
    type InventoryResponsePaginated struct {

        
            Items []InventoryResponse  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // ReturnConfig1 used by Catalog
    type ReturnConfig1 struct {

        
            Returnable bool  `json:"returnable"`
            Unit string  `json:"unit"`
            Time float64  `json:"time"`
         
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
            NotAvailable QuantityBase  `json:"not_available"`
            Sellable QuantityBase  `json:"sellable"`
            Damaged QuantityBase  `json:"damaged"`
         
    }
    
    // CompanyMeta used by Catalog
    type CompanyMeta struct {

        
            ID float64  `json:"id"`
         
    }
    
    // DimensionResponse used by Catalog
    type DimensionResponse struct {

        
            IsDefault bool  `json:"is_default"`
            Height float64  `json:"height"`
            Width float64  `json:"width"`
            Length float64  `json:"length"`
            Unit string  `json:"unit"`
         
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
            Marked float64  `json:"marked"`
            Effective float64  `json:"effective"`
            Currency string  `json:"currency"`
            TpNotes map[string]interface{}  `json:"tp_notes"`
         
    }
    
    // ManufacturerResponse used by Catalog
    type ManufacturerResponse struct {

        
            Name string  `json:"name"`
            Address string  `json:"address"`
            IsDefault bool  `json:"is_default"`
         
    }
    
    // WeightResponse used by Catalog
    type WeightResponse struct {

        
            Unit string  `json:"unit"`
            IsDefault bool  `json:"is_default"`
            Shipping float64  `json:"shipping"`
         
    }
    
    // StoreMeta used by Catalog
    type StoreMeta struct {

        
            ID float64  `json:"id"`
         
    }
    
    // InventorySellerResponse used by Catalog
    type InventorySellerResponse struct {

        
            Identifier map[string]interface{}  `json:"identifier"`
            TotalQuantity float64  `json:"total_quantity"`
            ReturnConfig ReturnConfig1  `json:"return_config"`
            Trader []Trader1  `json:"trader"`
            UID string  `json:"uid"`
            CreatedBy UserSerializer  `json:"created_by"`
            SellerIdentifier string  `json:"seller_identifier"`
            ItemID float64  `json:"item_id"`
            Fragile bool  `json:"fragile"`
            Quantities Quantities  `json:"quantities"`
            TrackInventory bool  `json:"track_inventory"`
            Stage string  `json:"stage"`
            Company CompanyMeta  `json:"company"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Dimension DimensionResponse  `json:"dimension"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            RawMeta map[string]interface{}  `json:"raw_meta"`
            Meta map[string]interface{}  `json:"meta"`
            Brand BrandMeta  `json:"brand"`
            FyndItemCode string  `json:"fynd_item_code"`
            Tags []string  `json:"tags"`
            FyndMeta map[string]interface{}  `json:"fynd_meta"`
            Price PriceMeta  `json:"price"`
            CountryOfOrigin string  `json:"country_of_origin"`
            FyndArticleCode string  `json:"fynd_article_code"`
            Size string  `json:"size"`
            ExpirationDate string  `json:"expiration_date"`
            TraceID string  `json:"trace_id"`
            IsSet bool  `json:"is_set"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            Set InventorySet  `json:"set"`
            Manufacturer ManufacturerResponse  `json:"manufacturer"`
            Weight WeightResponse  `json:"weight"`
            IsActive bool  `json:"is_active"`
            AddedOnStore string  `json:"added_on_store"`
            Store StoreMeta  `json:"store"`
         
    }
    
    // InventorySellerIdentifierResponsePaginated used by Catalog
    type InventorySellerIdentifierResponsePaginated struct {

        
            Items []InventorySellerResponse  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // ReturnConfig2 used by Catalog
    type ReturnConfig2 struct {

        
            Returnable bool  `json:"returnable"`
            Unit string  `json:"unit"`
            Time float64  `json:"time"`
         
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
    
    // CompanyMeta1 used by Catalog
    type CompanyMeta1 struct {

        
            ID float64  `json:"id"`
         
    }
    
    // DimensionResponse1 used by Catalog
    type DimensionResponse1 struct {

        
            Width float64  `json:"width"`
            Unit string  `json:"unit"`
            Height float64  `json:"height"`
            Length float64  `json:"length"`
         
    }
    
    // BrandMeta1 used by Catalog
    type BrandMeta1 struct {

        
            Name string  `json:"name"`
            ID float64  `json:"id"`
         
    }
    
    // PriceArticle used by Catalog
    type PriceArticle struct {

        
            Transfer float64  `json:"transfer"`
            Marked float64  `json:"marked"`
            Effective float64  `json:"effective"`
            Currency string  `json:"currency"`
            TpNotes map[string]interface{}  `json:"tp_notes"`
         
    }
    
    // ManufacturerResponse1 used by Catalog
    type ManufacturerResponse1 struct {

        
            Name string  `json:"name"`
            Address string  `json:"address"`
            IsDefault bool  `json:"is_default"`
         
    }
    
    // WeightResponse1 used by Catalog
    type WeightResponse1 struct {

        
            Unit string  `json:"unit"`
            Shipping float64  `json:"shipping"`
         
    }
    
    // ArticleStoreResponse used by Catalog
    type ArticleStoreResponse struct {

        
            Name string  `json:"name"`
            UID float64  `json:"uid"`
            StoreType string  `json:"store_type"`
            StoreCode string  `json:"store_code"`
         
    }
    
    // GetInventories used by Catalog
    type GetInventories struct {

        
            Identifier map[string]interface{}  `json:"identifier"`
            TotalQuantity float64  `json:"total_quantity"`
            ReturnConfig ReturnConfig2  `json:"return_config"`
            Trader []Trader2  `json:"trader"`
            UID string  `json:"uid"`
            CreatedBy UserSerializer  `json:"created_by"`
            SellerIdentifier string  `json:"seller_identifier"`
            ItemID float64  `json:"item_id"`
            Quantities QuantitiesArticle  `json:"quantities"`
            DateMeta DateMeta  `json:"date_meta"`
            TrackInventory bool  `json:"track_inventory"`
            Stage string  `json:"stage"`
            Company CompanyMeta1  `json:"company"`
            Dimension DimensionResponse1  `json:"dimension"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            Platforms map[string]interface{}  `json:"platforms"`
            Brand BrandMeta1  `json:"brand"`
            Tags []string  `json:"tags"`
            Price PriceArticle  `json:"price"`
            CountryOfOrigin string  `json:"country_of_origin"`
            Size string  `json:"size"`
            InventoryUpdatedOn string  `json:"inventory_updated_on"`
            ExpirationDate string  `json:"expiration_date"`
            TraceID string  `json:"trace_id"`
            IsSet bool  `json:"is_set"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            Manufacturer ManufacturerResponse1  `json:"manufacturer"`
            ID string  `json:"id"`
            Weight WeightResponse1  `json:"weight"`
            Store ArticleStoreResponse  `json:"store"`
         
    }
    
    // GetInventoriesResponse used by Catalog
    type GetInventoriesResponse struct {

        
            Items []GetInventories  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // BulkInventoryGetItems used by Catalog
    type BulkInventoryGetItems struct {

        
            CompanyID float64  `json:"company_id"`
            FailedRecords []string  `json:"failed_records"`
            Stage string  `json:"stage"`
            CreatedOn string  `json:"created_on"`
            Failed float64  `json:"failed"`
            Cancelled float64  `json:"cancelled"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            ID string  `json:"id"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            CancelledRecords []string  `json:"cancelled_records"`
            FilePath string  `json:"file_path"`
            Total float64  `json:"total"`
            ModifiedOn string  `json:"modified_on"`
            IsActive bool  `json:"is_active"`
            Succeed float64  `json:"succeed"`
         
    }
    
    // BulkInventoryGet used by Catalog
    type BulkInventoryGet struct {

        
            Items []BulkInventoryGetItems  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // InventoryJobPayload used by Catalog
    type InventoryJobPayload struct {

        
            ExpirationDate string  `json:"expiration_date"`
            TraceID string  `json:"trace_id"`
            TotalQuantity float64  `json:"total_quantity"`
            PriceEffective float64  `json:"price_effective"`
            Tags []string  `json:"tags"`
            Price float64  `json:"price"`
            ItemDimensionsUnitOfMeasure string  `json:"item_dimensions_unit_of_measure"`
            SellerIdentifier string  `json:"seller_identifier"`
            ItemWeightUnitOfMeasure string  `json:"item_weight_unit_of_measure"`
            StoreCode string  `json:"store_code"`
            PriceMarked float64  `json:"price_marked"`
            Quantity float64  `json:"quantity"`
            Currency string  `json:"currency"`
         
    }
    
    // InventoryBulkRequest used by Catalog
    type InventoryBulkRequest struct {

        
            CompanyID float64  `json:"company_id"`
            User map[string]interface{}  `json:"user"`
            BatchID string  `json:"batch_id"`
            Sizes []InventoryJobPayload  `json:"sizes"`
         
    }
    
    // InventoryExportRequest used by Catalog
    type InventoryExportRequest struct {

        
            Brand []float64  `json:"brand"`
            Type string  `json:"type"`
            Store []float64  `json:"store"`
         
    }
    
    // InventoryExportResponse used by Catalog
    type InventoryExportResponse struct {

        
            NotificationEmails []string  `json:"notification_emails"`
            Type string  `json:"type"`
            CreatedOn string  `json:"created_on"`
            Status string  `json:"status"`
            SellerID float64  `json:"seller_id"`
            CreatedBy string  `json:"created_by"`
            TaskID string  `json:"task_id"`
            Filters map[string]interface{}  `json:"filters"`
            ModifiedOn string  `json:"modified_on"`
         
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
            ToDate string  `json:"to_date"`
            StoreIds []float64  `json:"store_ids"`
            FromDate string  `json:"from_date"`
            Quantity InventoryExportQuantityFilter  `json:"quantity"`
         
    }
    
    // InventoryExportJob used by Catalog
    type InventoryExportJob struct {

        
            CompletedOn string  `json:"completed_on"`
            NotificationEmails []string  `json:"notification_emails"`
            Type string  `json:"type"`
            Status string  `json:"status"`
            SellerID float64  `json:"seller_id"`
            TaskID string  `json:"task_id"`
            Filters InventoryExportAdvanceOption  `json:"filters"`
            URL string  `json:"url"`
         
    }
    
    // InventoryExportFilter used by Catalog
    type InventoryExportFilter struct {

        
            BrandIds []float64  `json:"brand_ids"`
            ToDate string  `json:"to_date"`
            StoreIds []float64  `json:"store_ids"`
            FromDate string  `json:"from_date"`
            Quantity InventoryExportQuantityFilter  `json:"quantity"`
         
    }
    
    // InventoryCreateRequest used by Catalog
    type InventoryCreateRequest struct {

        
            NotificationEmails []string  `json:"notification_emails"`
            Filters InventoryExportFilter  `json:"filters"`
            Type string  `json:"type"`
            Data []string  `json:"data"`
         
    }
    
    // InventoryJobFilters used by Catalog
    type InventoryJobFilters struct {

        
            Stores []string  `json:"stores"`
            ToDate string  `json:"to_date"`
            FromDate string  `json:"from_date"`
            Quantity InventoryExportQuantityFilter  `json:"quantity"`
            Brands []string  `json:"brands"`
         
    }
    
    // InventoryJobDetailResponse used by Catalog
    type InventoryJobDetailResponse struct {

        
            CompletedOn string  `json:"completed_on"`
            NotificationEmails []string  `json:"notification_emails"`
            CancelledBy UserDetail  `json:"cancelled_by"`
            Type string  `json:"type"`
            CreatedOn string  `json:"created_on"`
            Status string  `json:"status"`
            SellerID float64  `json:"seller_id"`
            ID string  `json:"id"`
            CreatedBy UserDetail  `json:"created_by"`
            TaskID string  `json:"task_id"`
            Filters InventoryJobFilters  `json:"filters"`
            URL string  `json:"url"`
            ModifiedOn string  `json:"modified_on"`
            CancelledOn string  `json:"cancelled_on"`
         
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

        
            Multivalues bool  `json:"multivalues"`
            Data []FilerList  `json:"data"`
         
    }
    
    // InventoryPayload used by Catalog
    type InventoryPayload struct {

        
            ExpirationDate string  `json:"expiration_date"`
            TraceID string  `json:"trace_id"`
            TotalQuantity float64  `json:"total_quantity"`
            PriceEffective float64  `json:"price_effective"`
            Tags []string  `json:"tags"`
            SellerIdentifier string  `json:"seller_identifier"`
            PriceMarked float64  `json:"price_marked"`
            StoreID float64  `json:"store_id"`
         
    }
    
    // InventoryRequestSchemaV2 used by Catalog
    type InventoryRequestSchemaV2 struct {

        
            CompanyID float64  `json:"company_id"`
            Payload []InventoryPayload  `json:"payload"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // InventoryFailedReason used by Catalog
    type InventoryFailedReason struct {

        
            Message string  `json:"message"`
            Errors string  `json:"errors"`
         
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
    
    // HsnUpsert used by Catalog
    type HsnUpsert struct {

        
            CompanyID float64  `json:"company_id"`
            Hs2Code string  `json:"hs2_code"`
            Threshold1 float64  `json:"threshold1"`
            Tax2 float64  `json:"tax2"`
            TaxOnMrp bool  `json:"tax_on_mrp"`
            UID float64  `json:"uid"`
            HsnCode string  `json:"hsn_code"`
            Tax1 float64  `json:"tax1"`
            Threshold2 float64  `json:"threshold2"`
            TaxOnEsp bool  `json:"tax_on_esp"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // HsnCodesObject used by Catalog
    type HsnCodesObject struct {

        
            CompanyID float64  `json:"company_id"`
            Hs2Code string  `json:"hs2_code"`
            Threshold1 float64  `json:"threshold1"`
            Tax2 float64  `json:"tax2"`
            TaxOnMrp bool  `json:"tax_on_mrp"`
            HsnCode string  `json:"hsn_code"`
            ID string  `json:"id"`
            Tax1 float64  `json:"tax1"`
            Threshold2 float64  `json:"threshold2"`
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

        
            Cess float64  `json:"cess"`
            Rate float64  `json:"rate"`
            Threshold float64  `json:"threshold"`
            EffectiveDate string  `json:"effective_date"`
         
    }
    
    // HSNDataInsertV2 used by Catalog
    type HSNDataInsertV2 struct {

        
            ReportingHsn string  `json:"reporting_hsn"`
            Type string  `json:"type"`
            CreatedOn string  `json:"created_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            HsnCode string  `json:"hsn_code"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Taxes []TaxSlab  `json:"taxes"`
            Description string  `json:"description"`
            HsnCodeID string  `json:"hsn_code_id"`
            ModifiedOn string  `json:"modified_on"`
            CountryCode string  `json:"country_code"`
         
    }
    
    // PageResponse used by Catalog
    type PageResponse struct {

        
            ItemTotal float64  `json:"item_total"`
            Current string  `json:"current"`
            HasNext bool  `json:"has_next"`
            HasPrevious bool  `json:"has_previous"`
            Size float64  `json:"size"`
         
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

        
            Action Action  `json:"action"`
            Departments []string  `json:"departments"`
            Slug string  `json:"slug"`
            Logo Media2  `json:"logo"`
            UID float64  `json:"uid"`
            Banners ImageUrls  `json:"banners"`
            Name string  `json:"name"`
            Discount string  `json:"discount"`
         
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

        
            UID float64  `json:"uid"`
            Slug string  `json:"slug"`
         
    }
    
    // ThirdLevelChild used by Catalog
    type ThirdLevelChild struct {

        
            Action Action  `json:"action"`
            Slug string  `json:"slug"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            UID float64  `json:"uid"`
            Banners ImageUrls  `json:"banners"`
            Name string  `json:"name"`
            Childs []map[string]interface{}  `json:"childs"`
         
    }
    
    // SecondLevelChild used by Catalog
    type SecondLevelChild struct {

        
            Action Action  `json:"action"`
            Slug string  `json:"slug"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            UID float64  `json:"uid"`
            Banners ImageUrls  `json:"banners"`
            Name string  `json:"name"`
            Childs []ThirdLevelChild  `json:"childs"`
         
    }
    
    // Child used by Catalog
    type Child struct {

        
            Action Action  `json:"action"`
            Slug string  `json:"slug"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            UID float64  `json:"uid"`
            Banners ImageUrls  `json:"banners"`
            Name string  `json:"name"`
            Childs []SecondLevelChild  `json:"childs"`
         
    }
    
    // CategoryItems used by Catalog
    type CategoryItems struct {

        
            Action Action  `json:"action"`
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
            Banners ImageUrls  `json:"banners"`
            Name string  `json:"name"`
            Childs []Child  `json:"childs"`
         
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

        
            SortOn []ProductSortOn  `json:"sort_on"`
            Page Page  `json:"page"`
            Items []ProductListingDetail  `json:"items"`
            Operators map[string]interface{}  `json:"operators"`
            Filters []ProductFilters  `json:"filters"`
         
    }
    
    // ProductDetail used by Catalog
    type ProductDetail struct {

        
            Slug string  `json:"slug"`
            Highlights []string  `json:"highlights"`
            UID float64  `json:"uid"`
            RatingCount float64  `json:"rating_count"`
            ShortDescription string  `json:"short_description"`
            ProductOnlineDate string  `json:"product_online_date"`
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
            Type string  `json:"type"`
            ImageNature string  `json:"image_nature"`
            HasVariant bool  `json:"has_variant"`
            Color string  `json:"color"`
            Similars []string  `json:"similars"`
            Brand ProductBrand  `json:"brand"`
            Tryouts []string  `json:"tryouts"`
            Name string  `json:"name"`
            PromoMeta map[string]interface{}  `json:"promo_meta"`
            ItemCode string  `json:"item_code"`
            Attributes map[string]interface{}  `json:"attributes"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            ItemType string  `json:"item_type"`
            Description string  `json:"description"`
            Medias []Media  `json:"medias"`
            Rating float64  `json:"rating"`
         
    }
    
    // InventoryPage used by Catalog
    type InventoryPage struct {

        
            ItemTotal float64  `json:"item_total"`
            Type string  `json:"type"`
            HasNext bool  `json:"has_next"`
            HasPrevious bool  `json:"has_previous"`
            NextID string  `json:"next_id"`
         
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
    
    // ArticleAssignment used by Catalog
    type ArticleAssignment struct {

        
            Strategy string  `json:"strategy"`
            Level string  `json:"level"`
         
    }
    
    // AssignStoreArticle used by Catalog
    type AssignStoreArticle struct {

        
            Query ArticleQuery  `json:"query"`
            Meta map[string]interface{}  `json:"meta"`
            Quantity float64  `json:"quantity"`
            ArticleAssignment ArticleAssignment  `json:"article_assignment"`
            GroupID string  `json:"group_id"`
         
    }
    
    // AssignStore used by Catalog
    type AssignStore struct {

        
            CompanyID float64  `json:"company_id"`
            ChannelIdentifier string  `json:"channel_identifier"`
            AppID string  `json:"app_id"`
            StoreIds []float64  `json:"store_ids"`
            Pincode string  `json:"pincode"`
            Articles []AssignStoreArticle  `json:"articles"`
            ChannelType string  `json:"channel_type"`
         
    }
    
    // ArticleAssignment1 used by Catalog
    type ArticleAssignment1 struct {

        
            Strategy string  `json:"strategy"`
            Level string  `json:"level"`
         
    }
    
    // StoreAssignResponse used by Catalog
    type StoreAssignResponse struct {

        
            CompanyID float64  `json:"company_id"`
            StorePincode float64  `json:"store_pincode"`
            PriceEffective float64  `json:"price_effective"`
            Status bool  `json:"status"`
            Quantity float64  `json:"quantity"`
            UID string  `json:"uid"`
            ID string  `json:"_id"`
            GroupID string  `json:"group_id"`
            SCity string  `json:"s_city"`
            ItemID float64  `json:"item_id"`
            Meta map[string]interface{}  `json:"meta"`
            PriceMarked float64  `json:"price_marked"`
            StrategyWiseListing []map[string]interface{}  `json:"strategy_wise_listing"`
            Index float64  `json:"index"`
            ArticleAssignment ArticleAssignment1  `json:"article_assignment"`
            StoreID float64  `json:"store_id"`
            Size string  `json:"size"`
         
    }
    
    // LocationIntegrationType used by Catalog
    type LocationIntegrationType struct {

        
            Order string  `json:"order"`
            Inventory string  `json:"inventory"`
         
    }
    
    // UserSerializer1 used by Catalog
    type UserSerializer1 struct {

        
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
            Contact string  `json:"contact"`
         
    }
    
    // GetAddressSerializer used by Catalog
    type GetAddressSerializer struct {

        
            Longitude float64  `json:"longitude"`
            State string  `json:"state"`
            Address1 string  `json:"address1"`
            AddressType string  `json:"address_type"`
            Latitude float64  `json:"latitude"`
            Pincode float64  `json:"pincode"`
            City string  `json:"city"`
            Landmark string  `json:"landmark"`
            Address2 string  `json:"address2"`
            CountryCode string  `json:"country_code"`
            Country string  `json:"country"`
         
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

        
            Type string  `json:"type"`
            Value string  `json:"value"`
            Verified bool  `json:"verified"`
            URL string  `json:"url"`
            LegalName string  `json:"legal_name"`
         
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
            CreatedOn string  `json:"created_on"`
            CompanyType string  `json:"company_type"`
            ModifiedBy UserSerializer2  `json:"modified_by"`
            UID float64  `json:"uid"`
            CreatedBy UserSerializer2  `json:"created_by"`
            Name string  `json:"name"`
            VerifiedBy UserSerializer2  `json:"verified_by"`
            BusinessType string  `json:"business_type"`
            RejectReason string  `json:"reject_reason"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            ModifiedOn string  `json:"modified_on"`
            VerifiedOn string  `json:"verified_on"`
         
    }
    
    // SellerPhoneNumber used by Catalog
    type SellerPhoneNumber struct {

        
            CountryCode float64  `json:"country_code"`
            Number string  `json:"number"`
         
    }
    
    // ProductReturnConfigSerializer used by Catalog
    type ProductReturnConfigSerializer struct {

        
            OnSameStore bool  `json:"on_same_store"`
            StoreUID float64  `json:"store_uid"`
         
    }
    
    // LocationTimingSerializer used by Catalog
    type LocationTimingSerializer struct {

        
            Hour float64  `json:"hour"`
            Minute float64  `json:"minute"`
         
    }
    
    // LocationDayWiseSerializer used by Catalog
    type LocationDayWiseSerializer struct {

        
            Weekday string  `json:"weekday"`
            Opening LocationTimingSerializer  `json:"opening"`
            Closing LocationTimingSerializer  `json:"closing"`
            Open bool  `json:"open"`
         
    }
    
    // LocationManagerSerializer used by Catalog
    type LocationManagerSerializer struct {

        
            Name string  `json:"name"`
            Email string  `json:"email"`
            MobileNo SellerPhoneNumber  `json:"mobile_no"`
         
    }
    
    // GetLocationSerializer used by Catalog
    type GetLocationSerializer struct {

        
            DisplayName string  `json:"display_name"`
            IntegrationType LocationIntegrationType  `json:"integration_type"`
            Warnings map[string]interface{}  `json:"warnings"`
            UID float64  `json:"uid"`
            CreatedBy UserSerializer1  `json:"created_by"`
            VerifiedBy UserSerializer1  `json:"verified_by"`
            Address GetAddressSerializer  `json:"address"`
            GstCredentials InvoiceDetailsSerializer  `json:"gst_credentials"`
            Documents []Document  `json:"documents"`
            NotificationEmails []string  `json:"notification_emails"`
            Stage string  `json:"stage"`
            CreatedOn string  `json:"created_on"`
            Company GetCompanySerializer  `json:"company"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            VerifiedOn string  `json:"verified_on"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            Name string  `json:"name"`
            ModifiedOn string  `json:"modified_on"`
            StoreType string  `json:"store_type"`
            ModifiedBy UserSerializer1  `json:"modified_by"`
            ProductReturnConfig ProductReturnConfigSerializer  `json:"product_return_config"`
            Timing []LocationDayWiseSerializer  `json:"timing"`
            PhoneNumber string  `json:"phone_number"`
            Code string  `json:"code"`
            Manager LocationManagerSerializer  `json:"manager"`
         
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

        
            AppID string  `json:"app_id"`
            Logo string  `json:"logo"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            IsActive bool  `json:"is_active"`
         
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
    

    
    // UserSerializer used by CompanyProfile
    type UserSerializer struct {

        
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
            Contact string  `json:"contact"`
         
    }
    
    // SellerPhoneNumber used by CompanyProfile
    type SellerPhoneNumber struct {

        
            CountryCode float64  `json:"country_code"`
            Number string  `json:"number"`
         
    }
    
    // ContactDetails used by CompanyProfile
    type ContactDetails struct {

        
            Phone []SellerPhoneNumber  `json:"phone"`
            Emails []string  `json:"emails"`
         
    }
    
    // BusinessCountryInfo used by CompanyProfile
    type BusinessCountryInfo struct {

        
            CountryCode string  `json:"country_code"`
            Country string  `json:"country"`
         
    }
    
    // Website used by CompanyProfile
    type Website struct {

        
            URL string  `json:"url"`
         
    }
    
    // BusinessDetails used by CompanyProfile
    type BusinessDetails struct {

        
            Website Website  `json:"website"`
         
    }
    
    // Document used by CompanyProfile
    type Document struct {

        
            Value string  `json:"value"`
            LegalName string  `json:"legal_name"`
            Verified bool  `json:"verified"`
            URL string  `json:"url"`
            Type string  `json:"type"`
         
    }
    
    // GetAddressSerializer used by CompanyProfile
    type GetAddressSerializer struct {

        
            Country string  `json:"country"`
            Pincode float64  `json:"pincode"`
            CountryCode string  `json:"country_code"`
            Address2 string  `json:"address2"`
            City string  `json:"city"`
            State string  `json:"state"`
            Longitude float64  `json:"longitude"`
            AddressType string  `json:"address_type"`
            Landmark string  `json:"landmark"`
            Address1 string  `json:"address1"`
            Latitude float64  `json:"latitude"`
         
    }
    
    // CompanyTaxesSerializer used by CompanyProfile
    type CompanyTaxesSerializer struct {

        
            EffectiveDate string  `json:"effective_date"`
            Enable bool  `json:"enable"`
            Rate float64  `json:"rate"`
         
    }
    
    // GetCompanyProfileSerializerResponse used by CompanyProfile
    type GetCompanyProfileSerializerResponse struct {

        
            ModifiedBy UserSerializer  `json:"modified_by"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            VerifiedOn string  `json:"verified_on"`
            ModifiedOn string  `json:"modified_on"`
            BusinessType string  `json:"business_type"`
            ContactDetails ContactDetails  `json:"contact_details"`
            Warnings map[string]interface{}  `json:"warnings"`
            FranchiseEnabled bool  `json:"franchise_enabled"`
            CompanyType string  `json:"company_type"`
            BusinessCountryInfo BusinessCountryInfo  `json:"business_country_info"`
            CreatedOn string  `json:"created_on"`
            Name string  `json:"name"`
            Stage string  `json:"stage"`
            Mode string  `json:"mode"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            BusinessDetails BusinessDetails  `json:"business_details"`
            BusinessInfo string  `json:"business_info"`
            CreatedBy UserSerializer  `json:"created_by"`
            Documents []Document  `json:"documents"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            NotificationEmails []string  `json:"notification_emails"`
            UID float64  `json:"uid"`
            Taxes []CompanyTaxesSerializer  `json:"taxes"`
         
    }
    
    // ErrorResponse used by CompanyProfile
    type ErrorResponse struct {

        
            Message string  `json:"message"`
            Status float64  `json:"status"`
            Meta map[string]interface{}  `json:"meta"`
            Code string  `json:"code"`
         
    }
    
    // CreateUpdateAddressSerializer used by CompanyProfile
    type CreateUpdateAddressSerializer struct {

        
            Country string  `json:"country"`
            Address1 string  `json:"address1"`
            Pincode float64  `json:"pincode"`
            CountryCode string  `json:"country_code"`
            Address2 string  `json:"address2"`
            State string  `json:"state"`
            Longitude float64  `json:"longitude"`
            AddressType string  `json:"address_type"`
            Landmark string  `json:"landmark"`
            City string  `json:"city"`
            Latitude float64  `json:"latitude"`
         
    }
    
    // CompanyTaxesSerializer1 used by CompanyProfile
    type CompanyTaxesSerializer1 struct {

        
            EffectiveDate string  `json:"effective_date"`
            Enable bool  `json:"enable"`
            Rate float64  `json:"rate"`
         
    }
    
    // UpdateCompany used by CompanyProfile
    type UpdateCompany struct {

        
            Warnings map[string]interface{}  `json:"warnings"`
            Documents []Document  `json:"documents"`
            FranchiseEnabled bool  `json:"franchise_enabled"`
            Slug string  `json:"slug"`
            Addresses []CreateUpdateAddressSerializer  `json:"addresses"`
            NotificationEmails []string  `json:"notification_emails"`
            CompanyType string  `json:"company_type"`
            Taxes []CompanyTaxesSerializer1  `json:"taxes"`
            BusinessDetails BusinessDetails  `json:"business_details"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            BusinessInfo string  `json:"business_info"`
            Name string  `json:"name"`
            BusinessType string  `json:"business_type"`
            ContactDetails ContactDetails  `json:"contact_details"`
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

        
            Store DocumentsObj  `json:"store"`
            Product DocumentsObj  `json:"product"`
            UID float64  `json:"uid"`
            StoreDocuments DocumentsObj  `json:"store_documents"`
            CompanyDocuments DocumentsObj  `json:"company_documents"`
            Stage string  `json:"stage"`
            Brand DocumentsObj  `json:"brand"`
         
    }
    
    // BrandBannerSerializer used by CompanyProfile
    type BrandBannerSerializer struct {

        
            Landscape string  `json:"landscape"`
            Portrait string  `json:"portrait"`
         
    }
    
    // GetBrandResponseSerializer used by CompanyProfile
    type GetBrandResponseSerializer struct {

        
            ModifiedBy UserSerializer  `json:"modified_by"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            VerifiedOn string  `json:"verified_on"`
            ModifiedOn string  `json:"modified_on"`
            Synonyms []string  `json:"synonyms"`
            Warnings map[string]interface{}  `json:"warnings"`
            Logo string  `json:"logo"`
            Description string  `json:"description"`
            SlugKey string  `json:"slug_key"`
            CreatedOn string  `json:"created_on"`
            Name string  `json:"name"`
            Stage string  `json:"stage"`
            Mode string  `json:"mode"`
            Banner BrandBannerSerializer  `json:"banner"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            CreatedBy UserSerializer  `json:"created_by"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            UID float64  `json:"uid"`
            RejectReason string  `json:"reject_reason"`
         
    }
    
    // CreateUpdateBrandRequestSerializer used by CompanyProfile
    type CreateUpdateBrandRequestSerializer struct {

        
            BrandTier string  `json:"brand_tier"`
            Logo string  `json:"logo"`
            CompanyID float64  `json:"company_id"`
            Description string  `json:"description"`
            UID float64  `json:"uid"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Name string  `json:"name"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            Synonyms []string  `json:"synonyms"`
            Banner BrandBannerSerializer  `json:"banner"`
         
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

        
            Socials []CompanySocialAccounts  `json:"socials"`
            WebsiteURL string  `json:"website_url"`
         
    }
    
    // CompanySerializer used by CompanyProfile
    type CompanySerializer struct {

        
            VerifiedBy UserSerializer  `json:"verified_by"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            MarketChannels []string  `json:"market_channels"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            NotificationEmails []string  `json:"notification_emails"`
            UID float64  `json:"uid"`
            CompanyType string  `json:"company_type"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            BusinessCountryInfo BusinessCountryInfo  `json:"business_country_info"`
            CreatedOn string  `json:"created_on"`
            VerifiedOn string  `json:"verified_on"`
            ModifiedOn string  `json:"modified_on"`
            CreatedBy UserSerializer  `json:"created_by"`
            Name string  `json:"name"`
            BusinessType string  `json:"business_type"`
            Details CompanyDetails  `json:"details"`
            Stage string  `json:"stage"`
            RejectReason string  `json:"reject_reason"`
         
    }
    
    // CompanyBrandSerializer used by CompanyProfile
    type CompanyBrandSerializer struct {

        
            Warnings map[string]interface{}  `json:"warnings"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            Company CompanySerializer  `json:"company"`
            UID float64  `json:"uid"`
            CreatedOn string  `json:"created_on"`
            VerifiedOn string  `json:"verified_on"`
            ModifiedOn string  `json:"modified_on"`
            CreatedBy UserSerializer  `json:"created_by"`
            Stage string  `json:"stage"`
            Brand GetBrandResponseSerializer  `json:"brand"`
            RejectReason string  `json:"reject_reason"`
         
    }
    
    // CompanyBrandListSerializer used by CompanyProfile
    type CompanyBrandListSerializer struct {

        
            Page Page  `json:"page"`
            Items []CompanyBrandSerializer  `json:"items"`
         
    }
    
    // CompanyBrandPostRequestSerializer used by CompanyProfile
    type CompanyBrandPostRequestSerializer struct {

        
            Brands []float64  `json:"brands"`
            Company float64  `json:"company"`
            UID float64  `json:"uid"`
         
    }
    
    // ProductReturnConfigSerializer used by CompanyProfile
    type ProductReturnConfigSerializer struct {

        
            StoreUID float64  `json:"store_uid"`
            OnSameStore bool  `json:"on_same_store"`
         
    }
    
    // LocationTimingSerializer used by CompanyProfile
    type LocationTimingSerializer struct {

        
            Minute float64  `json:"minute"`
            Hour float64  `json:"hour"`
         
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

        
            Enabled bool  `json:"enabled"`
            Username string  `json:"username"`
            Password string  `json:"password"`
         
    }
    
    // InvoiceDetailsSerializer used by CompanyProfile
    type InvoiceDetailsSerializer struct {

        
            EWaybill InvoiceCredSerializer  `json:"e_waybill"`
            EInvoice InvoiceCredSerializer  `json:"e_invoice"`
         
    }
    
    // LocationManagerSerializer used by CompanyProfile
    type LocationManagerSerializer struct {

        
            Email string  `json:"email"`
            Name string  `json:"name"`
            MobileNo SellerPhoneNumber  `json:"mobile_no"`
         
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
    
    // GetCompanySerializer used by CompanyProfile
    type GetCompanySerializer struct {

        
            VerifiedBy UserSerializer  `json:"verified_by"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            UID float64  `json:"uid"`
            CompanyType string  `json:"company_type"`
            CreatedOn string  `json:"created_on"`
            VerifiedOn string  `json:"verified_on"`
            ModifiedOn string  `json:"modified_on"`
            CreatedBy UserSerializer  `json:"created_by"`
            Name string  `json:"name"`
            BusinessType string  `json:"business_type"`
            Stage string  `json:"stage"`
            RejectReason string  `json:"reject_reason"`
         
    }
    
    // GetLocationSerializer used by CompanyProfile
    type GetLocationSerializer struct {

        
            AutoInvoice bool  `json:"auto_invoice"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            ProductReturnConfig ProductReturnConfigSerializer  `json:"product_return_config"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            VerifiedOn string  `json:"verified_on"`
            ModifiedOn string  `json:"modified_on"`
            Timing []LocationDayWiseSerializer  `json:"timing"`
            Warnings map[string]interface{}  `json:"warnings"`
            CreditNote bool  `json:"credit_note"`
            GstCredentials InvoiceDetailsSerializer  `json:"gst_credentials"`
            CreatedOn string  `json:"created_on"`
            Name string  `json:"name"`
            Stage string  `json:"stage"`
            Manager LocationManagerSerializer  `json:"manager"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            Address GetAddressSerializer  `json:"address"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            Holiday []HolidaySchemaSerializer  `json:"holiday"`
            CreatedBy UserSerializer  `json:"created_by"`
            DisplayName string  `json:"display_name"`
            Documents []Document  `json:"documents"`
            StoreType string  `json:"store_type"`
            PhoneNumber string  `json:"phone_number"`
            Company GetCompanySerializer  `json:"company"`
            NotificationEmails []string  `json:"notification_emails"`
            UID float64  `json:"uid"`
            Code string  `json:"code"`
         
    }
    
    // LocationListSerializer used by CompanyProfile
    type LocationListSerializer struct {

        
            Page Page  `json:"page"`
            Items []GetLocationSerializer  `json:"items"`
         
    }
    
    // AddressSerializer used by CompanyProfile
    type AddressSerializer struct {

        
            Country string  `json:"country"`
            Pincode float64  `json:"pincode"`
            CountryCode string  `json:"country_code"`
            Address2 string  `json:"address2"`
            City string  `json:"city"`
            State string  `json:"state"`
            Longitude float64  `json:"longitude"`
            AddressType string  `json:"address_type"`
            Landmark string  `json:"landmark"`
            Address1 string  `json:"address1"`
            Latitude float64  `json:"latitude"`
         
    }
    
    // LocationSerializer used by CompanyProfile
    type LocationSerializer struct {

        
            AutoInvoice bool  `json:"auto_invoice"`
            Slug string  `json:"slug"`
            ProductReturnConfig ProductReturnConfigSerializer  `json:"product_return_config"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Timing []LocationDayWiseSerializer  `json:"timing"`
            Warnings map[string]interface{}  `json:"warnings"`
            CreditNote bool  `json:"credit_note"`
            GstCredentials InvoiceDetailsSerializer  `json:"gst_credentials"`
            Name string  `json:"name"`
            Stage string  `json:"stage"`
            Manager LocationManagerSerializer  `json:"manager"`
            Address AddressSerializer  `json:"address"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            Holiday []HolidaySchemaSerializer  `json:"holiday"`
            DisplayName string  `json:"display_name"`
            Documents []Document  `json:"documents"`
            StoreType string  `json:"store_type"`
            Company float64  `json:"company"`
            NotificationEmails []string  `json:"notification_emails"`
            UID float64  `json:"uid"`
            Code string  `json:"code"`
         
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
    

    
    // Identifier used by Cart
    type Identifier struct {

        
            UserID []string  `json:"user_id"`
            StoreID []float64  `json:"store_id"`
            ItemID []float64  `json:"item_id"`
            CollectionID []string  `json:"collection_id"`
            ExcludeBrandID []float64  `json:"exclude_brand_id"`
            ArticleID []string  `json:"article_id"`
            CategoryID []float64  `json:"category_id"`
            EmailDomain []string  `json:"email_domain"`
            CompanyID []float64  `json:"company_id"`
            BrandID []float64  `json:"brand_id"`
         
    }
    
    // Validity used by Cart
    type Validity struct {

        
            Priority float64  `json:"priority"`
         
    }
    
    // State used by Cart
    type State struct {

        
            IsDisplay bool  `json:"is_display"`
            IsArchived bool  `json:"is_archived"`
            IsPublic bool  `json:"is_public"`
         
    }
    
    // DisplayMetaDict used by Cart
    type DisplayMetaDict struct {

        
            Subtitle string  `json:"subtitle"`
            Title string  `json:"title"`
         
    }
    
    // DisplayMeta used by Cart
    type DisplayMeta struct {

        
            Apply DisplayMetaDict  `json:"apply"`
            Title string  `json:"title"`
            Auto DisplayMetaDict  `json:"auto"`
            Description string  `json:"description"`
            Subtitle string  `json:"subtitle"`
            Remove DisplayMetaDict  `json:"remove"`
         
    }
    
    // CouponAction used by Cart
    type CouponAction struct {

        
            TxnMode string  `json:"txn_mode"`
            ActionDate string  `json:"action_date"`
         
    }
    
    // CouponAuthor used by Cart
    type CouponAuthor struct {

        
            CreatedBy string  `json:"created_by"`
            ModifiedBy string  `json:"modified_by"`
         
    }
    
    // CouponDateMeta used by Cart
    type CouponDateMeta struct {

        
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
         
    }
    
    // Validation used by Cart
    type Validation struct {

        
            UserRegisteredAfter string  `json:"user_registered_after"`
            Anonymous bool  `json:"anonymous"`
            AppID []string  `json:"app_id"`
         
    }
    
    // Rule used by Cart
    type Rule struct {

        
            Key float64  `json:"key"`
            Value float64  `json:"value"`
            Max float64  `json:"max"`
            Min float64  `json:"min"`
            DiscountQty float64  `json:"discount_qty"`
         
    }
    
    // CouponSchedule used by Cart
    type CouponSchedule struct {

        
            Cron string  `json:"cron"`
            End string  `json:"end"`
            Start string  `json:"start"`
            NextSchedule []map[string]interface{}  `json:"next_schedule"`
            Duration float64  `json:"duration"`
         
    }
    
    // RuleDefinition used by Cart
    type RuleDefinition struct {

        
            Scope []string  `json:"scope"`
            AutoApply bool  `json:"auto_apply"`
            ValueType string  `json:"value_type"`
            CalculateOn string  `json:"calculate_on"`
            Type string  `json:"type"`
            ApplicableOn string  `json:"applicable_on"`
            IsExact bool  `json:"is_exact"`
            CurrencyCode string  `json:"currency_code"`
         
    }
    
    // PriceRange used by Cart
    type PriceRange struct {

        
            Max float64  `json:"max"`
            Min float64  `json:"min"`
         
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

        
            Iins []string  `json:"iins"`
            Networks []string  `json:"networks"`
            Codes []string  `json:"codes"`
            Types []string  `json:"types"`
            Uses PaymentAllowValue  `json:"uses"`
         
    }
    
    // PostOrder used by Cart
    type PostOrder struct {

        
            CancellationAllowed bool  `json:"cancellation_allowed"`
            ReturnAllowed bool  `json:"return_allowed"`
         
    }
    
    // UsesRemaining used by Cart
    type UsesRemaining struct {

        
            Total float64  `json:"total"`
            App float64  `json:"app"`
            User float64  `json:"user"`
         
    }
    
    // UsesRestriction used by Cart
    type UsesRestriction struct {

        
            Remaining UsesRemaining  `json:"remaining"`
            Maximum UsesRemaining  `json:"maximum"`
         
    }
    
    // Restrictions used by Cart
    type Restrictions struct {

        
            UserType string  `json:"user_type"`
            PriceRange PriceRange  `json:"price_range"`
            BulkBundle BulkBundleRestriction  `json:"bulk_bundle"`
            Payments map[string]PaymentModes  `json:"payments"`
            CouponAllowed bool  `json:"coupon_allowed"`
            PostOrder PostOrder  `json:"post_order"`
            UserGroups []float64  `json:"user_groups"`
            Platforms []string  `json:"platforms"`
            Uses UsesRestriction  `json:"uses"`
            OrderingStores []float64  `json:"ordering_stores"`
         
    }
    
    // Ownership used by Cart
    type Ownership struct {

        
            PayableCategory string  `json:"payable_category"`
            PayableBy string  `json:"payable_by"`
         
    }
    
    // CouponAdd used by Cart
    type CouponAdd struct {

        
            Identifiers Identifier  `json:"identifiers"`
            Validity Validity  `json:"validity"`
            State State  `json:"state"`
            DisplayMeta DisplayMeta  `json:"display_meta"`
            Action CouponAction  `json:"action"`
            TypeSlug string  `json:"type_slug"`
            Author CouponAuthor  `json:"author"`
            Tags []string  `json:"tags"`
            DateMeta CouponDateMeta  `json:"date_meta"`
            Code string  `json:"code"`
            Validation Validation  `json:"validation"`
            Rule []Rule  `json:"rule"`
            Schedule CouponSchedule  `json:"_schedule"`
            RuleDefinition RuleDefinition  `json:"rule_definition"`
            Restrictions Restrictions  `json:"restrictions"`
            Ownership Ownership  `json:"ownership"`
         
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

        
            Identifiers Identifier  `json:"identifiers"`
            Validity Validity  `json:"validity"`
            State State  `json:"state"`
            DisplayMeta DisplayMeta  `json:"display_meta"`
            Action CouponAction  `json:"action"`
            TypeSlug string  `json:"type_slug"`
            Author CouponAuthor  `json:"author"`
            Tags []string  `json:"tags"`
            DateMeta CouponDateMeta  `json:"date_meta"`
            Code string  `json:"code"`
            Validation Validation  `json:"validation"`
            Rule []Rule  `json:"rule"`
            Schedule CouponSchedule  `json:"_schedule"`
            RuleDefinition RuleDefinition  `json:"rule_definition"`
            Restrictions Restrictions  `json:"restrictions"`
            Ownership Ownership  `json:"ownership"`
         
    }
    
    // CouponPartialUpdate used by Cart
    type CouponPartialUpdate struct {

        
            Schedule CouponSchedule  `json:"schedule"`
            Archive bool  `json:"archive"`
         
    }
    
    // CompareObject used by Cart
    type CompareObject struct {

        
            GreaterThan float64  `json:"greater_than"`
            LessThan float64  `json:"less_than"`
            GreaterThanEquals float64  `json:"greater_than_equals"`
            Equals float64  `json:"equals"`
            LessThanEquals float64  `json:"less_than_equals"`
         
    }
    
    // ItemCriteria used by Cart
    type ItemCriteria struct {

        
            ItemL2Category []float64  `json:"item_l2_category"`
            ItemExcludeBrand []float64  `json:"item_exclude_brand"`
            CartQuantity CompareObject  `json:"cart_quantity"`
            BuyRules []string  `json:"buy_rules"`
            CartUniqueItemQuantity CompareObject  `json:"cart_unique_item_quantity"`
            ItemExcludeCompany []float64  `json:"item_exclude_company"`
            ItemSize []string  `json:"item_size"`
            ItemBrand []float64  `json:"item_brand"`
            ItemStore []float64  `json:"item_store"`
            CartTotal CompareObject  `json:"cart_total"`
            ItemDepartment []float64  `json:"item_department"`
            ProductTags []string  `json:"product_tags"`
            AvailableZones []string  `json:"available_zones"`
            ItemTags []string  `json:"item_tags"`
            ItemExcludeL2Category []float64  `json:"item_exclude_l2_category"`
            ItemExcludeL1Category []float64  `json:"item_exclude_l1_category"`
            ItemL1Category []float64  `json:"item_l1_category"`
            ItemCategory []float64  `json:"item_category"`
            ItemExcludeID []float64  `json:"item_exclude_id"`
            AllItems bool  `json:"all_items"`
            ItemExcludeDepartment []float64  `json:"item_exclude_department"`
            ItemExcludeSku []string  `json:"item_exclude_sku"`
            ItemID []float64  `json:"item_id"`
            ItemSku []string  `json:"item_sku"`
            ItemExcludeCategory []float64  `json:"item_exclude_category"`
            ItemCompany []float64  `json:"item_company"`
            ItemExcludeStore []float64  `json:"item_exclude_store"`
            CartUniqueItemAmount CompareObject  `json:"cart_unique_item_amount"`
         
    }
    
    // DiscountOffer used by Cart
    type DiscountOffer struct {

        
            MaxDiscountAmount float64  `json:"max_discount_amount"`
            MinOfferQuantity float64  `json:"min_offer_quantity"`
            MaxOfferQuantity float64  `json:"max_offer_quantity"`
            ApportionDiscount bool  `json:"apportion_discount"`
            DiscountPercentage float64  `json:"discount_percentage"`
            PartialCanRet bool  `json:"partial_can_ret"`
            DiscountAmount float64  `json:"discount_amount"`
            Code string  `json:"code"`
            MaxUsagePerTransaction float64  `json:"max_usage_per_transaction"`
            DiscountPrice float64  `json:"discount_price"`
         
    }
    
    // DiscountRule used by Cart
    type DiscountRule struct {

        
            DiscountType string  `json:"discount_type"`
            ItemCriteria ItemCriteria  `json:"item_criteria"`
            Offer DiscountOffer  `json:"offer"`
            BuyCondition string  `json:"buy_condition"`
         
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
    
    // PromotionSchedule used by Cart
    type PromotionSchedule struct {

        
            Cron string  `json:"cron"`
            End string  `json:"end"`
            Start string  `json:"start"`
            NextSchedule []map[string]interface{}  `json:"next_schedule"`
            Published bool  `json:"published"`
            Duration float64  `json:"duration"`
         
    }
    
    // Ownership1 used by Cart
    type Ownership1 struct {

        
            PayableCategory string  `json:"payable_category"`
            PayableBy string  `json:"payable_by"`
         
    }
    
    // PromotionAuthor used by Cart
    type PromotionAuthor struct {

        
            CreatedBy string  `json:"created_by"`
            ModifiedBy string  `json:"modified_by"`
         
    }
    
    // PromotionDateMeta used by Cart
    type PromotionDateMeta struct {

        
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
         
    }
    
    // DisplayMeta1 used by Cart
    type DisplayMeta1 struct {

        
            OfferLabel string  `json:"offer_label"`
            OfferText string  `json:"offer_text"`
            Name string  `json:"name"`
            Description string  `json:"description"`
         
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

        
            Remaining UsesRemaining1  `json:"remaining"`
            Maximum UsesRemaining1  `json:"maximum"`
         
    }
    
    // Restrictions1 used by Cart
    type Restrictions1 struct {

        
            UserRegistered UserRegistered  `json:"user_registered"`
            UserID []string  `json:"user_id"`
            OrderingStores []float64  `json:"ordering_stores"`
            OrderQuantity float64  `json:"order_quantity"`
            Payments []PromotionPaymentModes  `json:"payments"`
            PostOrder PostOrder1  `json:"post_order"`
            Platforms []string  `json:"platforms"`
            Uses UsesRestriction1  `json:"uses"`
            AnonymousUsers bool  `json:"anonymous_users"`
            UserGroups []float64  `json:"user_groups"`
         
    }
    
    // PromotionListItem used by Cart
    type PromotionListItem struct {

        
            DiscountRules []DiscountRule  `json:"discount_rules"`
            CalculateOn string  `json:"calculate_on"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            BuyRules map[string]ItemCriteria  `json:"buy_rules"`
            Stackable bool  `json:"stackable"`
            Visiblility Visibility  `json:"visiblility"`
            PromotionType string  `json:"promotion_type"`
            PostOrderAction PromotionAction  `json:"post_order_action"`
            Code string  `json:"code"`
            Schedule PromotionSchedule  `json:"_schedule"`
            Ownership Ownership1  `json:"ownership"`
            Mode string  `json:"mode"`
            Currency string  `json:"currency"`
            ApplyExclusive string  `json:"apply_exclusive"`
            ApplyAllDiscount bool  `json:"apply_all_discount"`
            Author PromotionAuthor  `json:"author"`
            DateMeta PromotionDateMeta  `json:"date_meta"`
            ApplyPriority float64  `json:"apply_priority"`
            PromoGroup string  `json:"promo_group"`
            ApplicationID string  `json:"application_id"`
            DisplayMeta DisplayMeta1  `json:"display_meta"`
            Restrictions Restrictions1  `json:"restrictions"`
         
    }
    
    // PromotionsResponse used by Cart
    type PromotionsResponse struct {

        
            Items PromotionListItem  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // PromotionAdd used by Cart
    type PromotionAdd struct {

        
            DiscountRules []DiscountRule  `json:"discount_rules"`
            CalculateOn string  `json:"calculate_on"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            BuyRules map[string]ItemCriteria  `json:"buy_rules"`
            Stackable bool  `json:"stackable"`
            Visiblility Visibility  `json:"visiblility"`
            PromotionType string  `json:"promotion_type"`
            PostOrderAction PromotionAction  `json:"post_order_action"`
            Code string  `json:"code"`
            Schedule PromotionSchedule  `json:"_schedule"`
            Ownership Ownership1  `json:"ownership"`
            Mode string  `json:"mode"`
            Currency string  `json:"currency"`
            ApplyExclusive string  `json:"apply_exclusive"`
            ApplyAllDiscount bool  `json:"apply_all_discount"`
            Author PromotionAuthor  `json:"author"`
            DateMeta PromotionDateMeta  `json:"date_meta"`
            ApplyPriority float64  `json:"apply_priority"`
            PromoGroup string  `json:"promo_group"`
            ApplicationID string  `json:"application_id"`
            DisplayMeta DisplayMeta1  `json:"display_meta"`
            Restrictions Restrictions1  `json:"restrictions"`
         
    }
    
    // PromotionUpdate used by Cart
    type PromotionUpdate struct {

        
            DiscountRules []DiscountRule  `json:"discount_rules"`
            CalculateOn string  `json:"calculate_on"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            BuyRules map[string]ItemCriteria  `json:"buy_rules"`
            Stackable bool  `json:"stackable"`
            Visiblility Visibility  `json:"visiblility"`
            PromotionType string  `json:"promotion_type"`
            PostOrderAction PromotionAction  `json:"post_order_action"`
            Code string  `json:"code"`
            Schedule PromotionSchedule  `json:"_schedule"`
            Ownership Ownership1  `json:"ownership"`
            Mode string  `json:"mode"`
            Currency string  `json:"currency"`
            ApplyExclusive string  `json:"apply_exclusive"`
            ApplyAllDiscount bool  `json:"apply_all_discount"`
            Author PromotionAuthor  `json:"author"`
            DateMeta PromotionDateMeta  `json:"date_meta"`
            ApplyPriority float64  `json:"apply_priority"`
            PromoGroup string  `json:"promo_group"`
            ApplicationID string  `json:"application_id"`
            DisplayMeta DisplayMeta1  `json:"display_meta"`
            Restrictions Restrictions1  `json:"restrictions"`
         
    }
    
    // PromotionPartialUpdate used by Cart
    type PromotionPartialUpdate struct {

        
            Schedule PromotionSchedule  `json:"schedule"`
            Archive bool  `json:"archive"`
         
    }
    
    // ActivePromosResponse used by Cart
    type ActivePromosResponse struct {

        
            ModifiedOn string  `json:"modified_on"`
            Title string  `json:"title"`
            EntitySlug string  `json:"entity_slug"`
            Example string  `json:"example"`
            Type string  `json:"type"`
            CreatedOn string  `json:"created_on"`
            Description string  `json:"description"`
            IsHidden bool  `json:"is_hidden"`
            Subtitle string  `json:"subtitle"`
            EntityType string  `json:"entity_type"`
         
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
            MaxCartItems float64  `json:"max_cart_items"`
            GiftDisplayText string  `json:"gift_display_text"`
            MinCartValue float64  `json:"min_cart_value"`
            RevenueEngineCoupon bool  `json:"revenue_engine_coupon"`
            Enabled bool  `json:"enabled"`
            DeliveryCharges DeliveryCharges  `json:"delivery_charges"`
            GiftPricing float64  `json:"gift_pricing"`
         
    }
    
    // CartMetaConfigAdd used by Cart
    type CartMetaConfigAdd struct {

        
            BulkCoupons bool  `json:"bulk_coupons"`
            MaxCartItems float64  `json:"max_cart_items"`
            GiftDisplayText string  `json:"gift_display_text"`
            MinCartValue float64  `json:"min_cart_value"`
            RevenueEngineCoupon bool  `json:"revenue_engine_coupon"`
            Enabled bool  `json:"enabled"`
            DeliveryCharges DeliveryCharges  `json:"delivery_charges"`
            GiftPricing float64  `json:"gift_pricing"`
         
    }
    
    // Collecttion used by Cart
    type Collecttion struct {

        
            CollectedBy string  `json:"collected_by"`
            RefundBy string  `json:"refund_by"`
         
    }
    
    // Article used by Cart
    type Article struct {

        
            Meta map[string]interface{}  `json:"meta"`
            Value float64  `json:"value"`
            Type string  `json:"type"`
            ArticleID string  `json:"article_id"`
            Code string  `json:"code"`
         
    }
    
    // CartDynamicInjectionUpdate used by Cart
    type CartDynamicInjectionUpdate struct {

        
            IsAuthenticated bool  `json:"is_authenticated"`
            ApplyExpiry string  `json:"apply_expiry"`
            UserID string  `json:"user_id"`
            Meta map[string]interface{}  `json:"meta"`
            Value float64  `json:"value"`
            Collection Collecttion  `json:"collection"`
            ArticleLevelDistribution bool  `json:"article_level_distribution"`
            ArticleIds []Article  `json:"article_ids"`
            CartID string  `json:"cart_id"`
            Type string  `json:"type"`
            AllowedRefund bool  `json:"allowed_refund"`
            Message string  `json:"message"`
            CartValue float64  `json:"cart_value"`
         
    }
    
    // CartDynamicInjection used by Cart
    type CartDynamicInjection struct {

        
            IsAuthenticated bool  `json:"is_authenticated"`
            ApplyExpiry string  `json:"apply_expiry"`
            UserID string  `json:"user_id"`
            Meta map[string]interface{}  `json:"meta"`
            Value float64  `json:"value"`
            Collection Collecttion  `json:"collection"`
            ArticleLevelDistribution bool  `json:"article_level_distribution"`
            ArticleIds []Article  `json:"article_ids"`
            CartID string  `json:"cart_id"`
            Type string  `json:"type"`
            AllowedRefund bool  `json:"allowed_refund"`
            Message string  `json:"message"`
            InjectionID string  `json:"injection_id"`
            CartValue float64  `json:"cart_value"`
         
    }
    
    // CartDynamicInjectionResponse used by Cart
    type CartDynamicInjectionResponse struct {

        
            Data CartDynamicInjection  `json:"data"`
         
    }
    
    // CartDynamicInjectionAdd used by Cart
    type CartDynamicInjectionAdd struct {

        
            IsAuthenticated bool  `json:"is_authenticated"`
            ApplyExpiry string  `json:"apply_expiry"`
            UserID string  `json:"user_id"`
            Meta map[string]interface{}  `json:"meta"`
            Value float64  `json:"value"`
            Collection Collecttion  `json:"collection"`
            ArticleLevelDistribution bool  `json:"article_level_distribution"`
            ArticleIds []Article  `json:"article_ids"`
            CartID string  `json:"cart_id"`
            Type string  `json:"type"`
            AllowedRefund bool  `json:"allowed_refund"`
            Message string  `json:"message"`
            CartValue float64  `json:"cart_value"`
         
    }
    
    // CartItem used by Cart
    type CartItem struct {

        
            ProductID string  `json:"product_id"`
            Quantity float64  `json:"quantity"`
            Size string  `json:"size"`
         
    }
    
    // OpenapiCartDetailsRequest used by Cart
    type OpenapiCartDetailsRequest struct {

        
            CartItems []CartItem  `json:"cart_items"`
         
    }
    
    // RawBreakup used by Cart
    type RawBreakup struct {

        
            GiftCard float64  `json:"gift_card"`
            Subtotal float64  `json:"subtotal"`
            Total float64  `json:"total"`
            YouSaved float64  `json:"you_saved"`
            ConvenienceFee float64  `json:"convenience_fee"`
            FyndCash float64  `json:"fynd_cash"`
            Coupon float64  `json:"coupon"`
            CodCharge float64  `json:"cod_charge"`
            Vog float64  `json:"vog"`
            GstCharges float64  `json:"gst_charges"`
            MrpTotal float64  `json:"mrp_total"`
            Discount float64  `json:"discount"`
            DeliveryCharge float64  `json:"delivery_charge"`
         
    }
    
    // CouponBreakup used by Cart
    type CouponBreakup struct {

        
            MinimumCartValue float64  `json:"minimum_cart_value"`
            Description string  `json:"description"`
            Value float64  `json:"value"`
            Title string  `json:"title"`
            Type string  `json:"type"`
            IsApplied bool  `json:"is_applied"`
            Message string  `json:"message"`
            CouponType string  `json:"coupon_type"`
            SubTitle string  `json:"sub_title"`
            MaxDiscountValue float64  `json:"max_discount_value"`
            Code string  `json:"code"`
            UID string  `json:"uid"`
            CouponValue float64  `json:"coupon_value"`
         
    }
    
    // LoyaltyPoints used by Cart
    type LoyaltyPoints struct {

        
            IsApplied bool  `json:"is_applied"`
            Total float64  `json:"total"`
            Applicable float64  `json:"applicable"`
            Description string  `json:"description"`
         
    }
    
    // DisplayBreakup used by Cart
    type DisplayBreakup struct {

        
            Display string  `json:"display"`
            Key string  `json:"key"`
            CurrencySymbol string  `json:"currency_symbol"`
            Value float64  `json:"value"`
            Message []string  `json:"message"`
            CurrencyCode string  `json:"currency_code"`
         
    }
    
    // CartBreakup used by Cart
    type CartBreakup struct {

        
            Raw RawBreakup  `json:"raw"`
            Coupon CouponBreakup  `json:"coupon"`
            LoyaltyPoints LoyaltyPoints  `json:"loyalty_points"`
            Display []DisplayBreakup  `json:"display"`
         
    }
    
    // ProductPrice used by Cart
    type ProductPrice struct {

        
            Marked float64  `json:"marked"`
            AddOn float64  `json:"add_on"`
            CurrencySymbol string  `json:"currency_symbol"`
            Selling float64  `json:"selling"`
            Effective float64  `json:"effective"`
            CurrencyCode string  `json:"currency_code"`
         
    }
    
    // ProductPriceInfo used by Cart
    type ProductPriceInfo struct {

        
            Converted ProductPrice  `json:"converted"`
            Base ProductPrice  `json:"base"`
         
    }
    
    // CouponDetails used by Cart
    type CouponDetails struct {

        
            DiscountTotalQuantity float64  `json:"discount_total_quantity"`
            DiscountSingleQuantity float64  `json:"discount_single_quantity"`
            Code string  `json:"code"`
         
    }
    
    // ProductAvailabilitySize used by Cart
    type ProductAvailabilitySize struct {

        
            Display string  `json:"display"`
            IsAvailable bool  `json:"is_available"`
            Meta map[string]interface{}  `json:"meta"`
            Value string  `json:"value"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // ProductAvailability used by Cart
    type ProductAvailability struct {

        
            IsValid bool  `json:"is_valid"`
            Sizes []string  `json:"sizes"`
            OtherStoreQuantity float64  `json:"other_store_quantity"`
            Deliverable bool  `json:"deliverable"`
            OutOfStock bool  `json:"out_of_stock"`
            AvailableSizes []ProductAvailabilitySize  `json:"available_sizes"`
         
    }
    
    // BaseInfo used by Cart
    type BaseInfo struct {

        
            Name string  `json:"name"`
            UID float64  `json:"uid"`
         
    }
    
    // StoreInfo used by Cart
    type StoreInfo struct {

        
            StoreCode string  `json:"store_code"`
            Name string  `json:"name"`
            UID float64  `json:"uid"`
         
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
    
    // ProductArticle used by Cart
    type ProductArticle struct {

        
            GiftCard map[string]interface{}  `json:"gift_card"`
            Quantity float64  `json:"quantity"`
            Size string  `json:"size"`
            UID string  `json:"uid"`
            Identifier map[string]interface{}  `json:"identifier"`
            Seller BaseInfo  `json:"seller"`
            Type string  `json:"type"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            MtoQuantity float64  `json:"mto_quantity"`
            IsGiftVisible bool  `json:"is_gift_visible"`
            Store StoreInfo  `json:"store"`
            ProductGroupTags []string  `json:"product_group_tags"`
            Price ArticlePriceInfo  `json:"price"`
            SellerIdentifier string  `json:"seller_identifier"`
            CartItemMeta map[string]interface{}  `json:"cart_item_meta"`
         
    }
    
    // CategoryInfo used by Cart
    type CategoryInfo struct {

        
            Name string  `json:"name"`
            UID float64  `json:"uid"`
         
    }
    
    // ActionQuery used by Cart
    type ActionQuery struct {

        
            ProductSlug []string  `json:"product_slug"`
         
    }
    
    // ProductAction used by Cart
    type ProductAction struct {

        
            Query ActionQuery  `json:"query"`
            Type string  `json:"type"`
            URL string  `json:"url"`
         
    }
    
    // ProductImage used by Cart
    type ProductImage struct {

        
            AspectRatio string  `json:"aspect_ratio"`
            SecureURL string  `json:"secure_url"`
            URL string  `json:"url"`
         
    }
    
    // NetQuantity used by Cart
    type NetQuantity struct {

        
            Unit string  `json:"unit"`
            Value string  `json:"value"`
         
    }
    
    // Tags used by Cart
    type Tags struct {

        
            Tags map[string]interface{}  `json:"tags"`
         
    }
    
    // CartProduct used by Cart
    type CartProduct struct {

        
            Brand BaseInfo  `json:"brand"`
            ItemCode string  `json:"item_code"`
            Type string  `json:"type"`
            Slug string  `json:"slug"`
            Categories []CategoryInfo  `json:"categories"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Action ProductAction  `json:"action"`
            Images []ProductImage  `json:"images"`
            Tags []string  `json:"tags"`
            NetQuantity NetQuantity  `json:"net_quantity"`
            UID float64  `json:"uid"`
            TeaserTag Tags  `json:"teaser_tag"`
            Name string  `json:"name"`
         
    }
    
    // Ownership2 used by Cart
    type Ownership2 struct {

        
            PayableCategory string  `json:"payable_category"`
            PayableBy string  `json:"payable_by"`
         
    }
    
    // BuyRules used by Cart
    type BuyRules struct {

        
            ItemCriteria map[string]interface{}  `json:"item_criteria"`
            CartConditions map[string]interface{}  `json:"cart_conditions"`
         
    }
    
    // FreeGiftItem used by Cart
    type FreeGiftItem struct {

        
            ItemBrandName string  `json:"item_brand_name"`
            ItemName string  `json:"item_name"`
            ItemPriceDetails map[string]interface{}  `json:"item_price_details"`
            ItemSlug string  `json:"item_slug"`
            ItemID float64  `json:"item_id"`
            ItemImagesURL []string  `json:"item_images_url"`
         
    }
    
    // AppliedFreeArticles used by Cart
    type AppliedFreeArticles struct {

        
            ParentItemIdentifier string  `json:"parent_item_identifier"`
            ArticleID string  `json:"article_id"`
            Quantity float64  `json:"quantity"`
            FreeGiftItemDetails FreeGiftItem  `json:"free_gift_item_details"`
         
    }
    
    // DiscountRulesApp used by Cart
    type DiscountRulesApp struct {

        
            ItemCriteria map[string]interface{}  `json:"item_criteria"`
            RawOffer map[string]interface{}  `json:"raw_offer"`
            MatchedBuyRules []string  `json:"matched_buy_rules"`
            Offer map[string]interface{}  `json:"offer"`
         
    }
    
    // AppliedPromotion used by Cart
    type AppliedPromotion struct {

        
            ArticleQuantity float64  `json:"article_quantity"`
            PromotionType string  `json:"promotion_type"`
            Ownership Ownership2  `json:"ownership"`
            PromotionName string  `json:"promotion_name"`
            BuyRules []BuyRules  `json:"buy_rules"`
            MrpPromotion bool  `json:"mrp_promotion"`
            PromotionOfferText string  `json:"promotion_offer_text"`
            PromotionGroup string  `json:"promotion_group"`
            PromoID string  `json:"promo_id"`
            AppliedFreeArticles []AppliedFreeArticles  `json:"applied_free_articles"`
            DiscountRules []DiscountRulesApp  `json:"discount_rules"`
            Amount float64  `json:"amount"`
         
    }
    
    // PromoMeta used by Cart
    type PromoMeta struct {

        
            Message string  `json:"message"`
         
    }
    
    // CartProductIdentifer used by Cart
    type CartProductIdentifer struct {

        
            Identifier string  `json:"identifier"`
         
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
    
    // CartProductInfo used by Cart
    type CartProductInfo struct {

        
            BulkOffer map[string]interface{}  `json:"bulk_offer"`
            Key string  `json:"key"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            PricePerUnit ProductPriceInfo  `json:"price_per_unit"`
            Coupon CouponDetails  `json:"coupon"`
            Availability ProductAvailability  `json:"availability"`
            Moq map[string]interface{}  `json:"moq"`
            CouponMessage string  `json:"coupon_message"`
            Price ProductPriceInfo  `json:"price"`
            Article ProductArticle  `json:"article"`
            IsSet bool  `json:"is_set"`
            Product CartProduct  `json:"product"`
            Message string  `json:"message"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            PromotionsApplied []AppliedPromotion  `json:"promotions_applied"`
            Quantity float64  `json:"quantity"`
            PromoMeta PromoMeta  `json:"promo_meta"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            Discount string  `json:"discount"`
         
    }
    
    // OpenapiCartDetailsResponse used by Cart
    type OpenapiCartDetailsResponse struct {

        
            Message string  `json:"message"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            Items []CartProductInfo  `json:"items"`
            IsValid bool  `json:"is_valid"`
         
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
            Email string  `json:"email"`
            CountryIsoCode string  `json:"country_iso_code"`
            State string  `json:"state"`
            CountryPhoneCode string  `json:"country_phone_code"`
            Pincode float64  `json:"pincode"`
            CountryCode string  `json:"country_code"`
            Landmark string  `json:"landmark"`
            Address string  `json:"address"`
            Country string  `json:"country"`
            Phone float64  `json:"phone"`
            City string  `json:"city"`
            AreaCodeSlug string  `json:"area_code_slug"`
            AreaCode string  `json:"area_code"`
            AddressType string  `json:"address_type"`
            Area string  `json:"area"`
            Name string  `json:"name"`
         
    }
    
    // OpenApiCartServiceabilityRequest used by Cart
    type OpenApiCartServiceabilityRequest struct {

        
            ShippingAddress ShippingAddress  `json:"shipping_address"`
            CartItems []CartItem  `json:"cart_items"`
         
    }
    
    // OpenApiCartServiceabilityResponse used by Cart
    type OpenApiCartServiceabilityResponse struct {

        
            BreakupValues CartBreakup  `json:"breakup_values"`
            Items []CartProductInfo  `json:"items"`
            IsValid bool  `json:"is_valid"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            Message string  `json:"message"`
         
    }
    
    // CartItemMeta used by Cart
    type CartItemMeta struct {

        
            PrimaryItem bool  `json:"primary_item"`
            GroupID string  `json:"group_id"`
         
    }
    
    // MultiTenderPaymentMeta used by Cart
    type MultiTenderPaymentMeta struct {

        
            PaymentGateway string  `json:"payment_gateway"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            CurrentStatus string  `json:"current_status"`
            OrderID string  `json:"order_id"`
            PaymentID string  `json:"payment_id"`
         
    }
    
    // MultiTenderPaymentMethod used by Cart
    type MultiTenderPaymentMethod struct {

        
            Meta MultiTenderPaymentMeta  `json:"meta"`
            Name string  `json:"name"`
            Mode string  `json:"mode"`
            Amount float64  `json:"amount"`
         
    }
    
    // OpenApiFiles used by Cart
    type OpenApiFiles struct {

        
            Values []string  `json:"values"`
            Key string  `json:"key"`
         
    }
    
    // OpenApiOrderItem used by Cart
    type OpenApiOrderItem struct {

        
            Quantity float64  `json:"quantity"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            Meta CartItemMeta  `json:"meta"`
            EmployeeDiscount float64  `json:"employee_discount"`
            PaymentMethods []MultiTenderPaymentMethod  `json:"payment_methods"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            Size string  `json:"size"`
            PriceEffective float64  `json:"price_effective"`
            Files []OpenApiFiles  `json:"files"`
            CashbackApplied float64  `json:"cashback_applied"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            AmountPaid float64  `json:"amount_paid"`
            CodCharges float64  `json:"cod_charges"`
            PriceMarked float64  `json:"price_marked"`
            DeliveryCharges float64  `json:"delivery_charges"`
            ProductID float64  `json:"product_id"`
            Discount float64  `json:"discount"`
         
    }
    
    // OpenApiPlatformCheckoutReq used by Cart
    type OpenApiPlatformCheckoutReq struct {

        
            ShippingAddress ShippingAddress  `json:"shipping_address"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            CartItems []OpenApiOrderItem  `json:"cart_items"`
            CashbackApplied float64  `json:"cashback_applied"`
            CartValue float64  `json:"cart_value"`
            DeliveryCharges float64  `json:"delivery_charges"`
            CurrencyCode string  `json:"currency_code"`
            BillingAddress ShippingAddress  `json:"billing_address"`
            Gstin string  `json:"gstin"`
            PaymentMethods []MultiTenderPaymentMethod  `json:"payment_methods"`
            PaymentMode string  `json:"payment_mode"`
            EmployeeDiscount map[string]interface{}  `json:"employee_discount"`
            Comment string  `json:"comment"`
            Coupon string  `json:"coupon"`
            Files []OpenApiFiles  `json:"files"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            CouponCode string  `json:"coupon_code"`
            CodCharges float64  `json:"cod_charges"`
            OrderID string  `json:"order_id"`
            CouponValue float64  `json:"coupon_value"`
         
    }
    
    // OpenApiCheckoutResponse used by Cart
    type OpenApiCheckoutResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
            OrderRefID string  `json:"order_ref_id"`
            OrderID string  `json:"order_id"`
         
    }
    
    // AbandonedCart used by Cart
    type AbandonedCart struct {

        
            UserID string  `json:"user_id"`
            CheckoutMode string  `json:"checkout_mode"`
            FyndCredits map[string]interface{}  `json:"fynd_credits"`
            CartValue float64  `json:"cart_value"`
            DeliveryCharges map[string]interface{}  `json:"delivery_charges"`
            PaymentMode string  `json:"payment_mode"`
            MergeQty bool  `json:"merge_qty"`
            IsArchive bool  `json:"is_archive"`
            Coupon map[string]interface{}  `json:"coupon"`
            ExpireAt string  `json:"expire_at"`
            FcIndexMap []float64  `json:"fc_index_map"`
            BuyNow bool  `json:"buy_now"`
            AppID string  `json:"app_id"`
            IsActive bool  `json:"is_active"`
            Cashback map[string]interface{}  `json:"cashback"`
            Shipments []map[string]interface{}  `json:"shipments"`
            ID string  `json:"_id"`
            Payments map[string]interface{}  `json:"payments"`
            BulkCouponDiscount float64  `json:"bulk_coupon_discount"`
            UID float64  `json:"uid"`
            PickUpCustomerDetails map[string]interface{}  `json:"pick_up_customer_details"`
            LastModified string  `json:"last_modified"`
            CreatedOn string  `json:"created_on"`
            Gstin string  `json:"gstin"`
            PaymentMethods []map[string]interface{}  `json:"payment_methods"`
            IsDefault bool  `json:"is_default"`
            Meta map[string]interface{}  `json:"meta"`
            Comment string  `json:"comment"`
            Articles []map[string]interface{}  `json:"articles"`
            CodCharges map[string]interface{}  `json:"cod_charges"`
            OrderID string  `json:"order_id"`
            Discount float64  `json:"discount"`
            Promotion map[string]interface{}  `json:"promotion"`
         
    }
    
    // AbandonedCartResponse used by Cart
    type AbandonedCartResponse struct {

        
            Items []AbandonedCart  `json:"items"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
            Result map[string]interface{}  `json:"result"`
            Page Page  `json:"page"`
         
    }
    
    // PaymentSelectionLock used by Cart
    type PaymentSelectionLock struct {

        
            Enabled bool  `json:"enabled"`
            PaymentIdentifier string  `json:"payment_identifier"`
            DefaultOptions string  `json:"default_options"`
         
    }
    
    // CartCurrency used by Cart
    type CartCurrency struct {

        
            Symbol string  `json:"symbol"`
            Code string  `json:"code"`
         
    }
    
    // CartDetailResponse used by Cart
    type CartDetailResponse struct {

        
            BreakupValues CartBreakup  `json:"breakup_values"`
            Items []CartProductInfo  `json:"items"`
            CheckoutMode string  `json:"checkout_mode"`
            ID string  `json:"id"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            BuyNow bool  `json:"buy_now"`
            IsValid bool  `json:"is_valid"`
            Currency CartCurrency  `json:"currency"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            Message string  `json:"message"`
            CouponText string  `json:"coupon_text"`
            LastModified string  `json:"last_modified"`
            Gstin string  `json:"gstin"`
            Comment string  `json:"comment"`
            AppliedPromoDetails []AppliedPromotion  `json:"applied_promo_details"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            PanNo string  `json:"pan_no"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            PanConfig map[string]interface{}  `json:"pan_config"`
         
    }
    
    // AddProductCart used by Cart
    type AddProductCart struct {

        
            Quantity float64  `json:"quantity"`
            StoreID float64  `json:"store_id"`
            Display string  `json:"display"`
            Meta map[string]interface{}  `json:"meta"`
            Pos bool  `json:"pos"`
            ItemSize string  `json:"item_size"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ParentItemIdentifiers []map[string]string  `json:"parent_item_identifiers"`
            SellerID float64  `json:"seller_id"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            ArticleID string  `json:"article_id"`
            ProductGroupTags []string  `json:"product_group_tags"`
            ItemID float64  `json:"item_id"`
            ArticleAssignment map[string]interface{}  `json:"article_assignment"`
         
    }
    
    // AddCartRequest used by Cart
    type AddCartRequest struct {

        
            Items []AddProductCart  `json:"items"`
            NewCart bool  `json:"new_cart"`
         
    }
    
    // AddCartDetailResponse used by Cart
    type AddCartDetailResponse struct {

        
            Message string  `json:"message"`
            Cart CartDetailResponse  `json:"cart"`
            Success bool  `json:"success"`
            Partial bool  `json:"partial"`
         
    }
    
    // UpdateProductCart used by Cart
    type UpdateProductCart struct {

        
            Quantity float64  `json:"quantity"`
            Meta map[string]interface{}  `json:"meta"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
            ItemSize string  `json:"item_size"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            ItemIndex float64  `json:"item_index"`
            ArticleID string  `json:"article_id"`
            ItemID float64  `json:"item_id"`
         
    }
    
    // UpdateCartRequest used by Cart
    type UpdateCartRequest struct {

        
            Items []UpdateProductCart  `json:"items"`
            Operation string  `json:"operation"`
         
    }
    
    // UpdateCartDetailResponse used by Cart
    type UpdateCartDetailResponse struct {

        
            Message string  `json:"message"`
            Cart CartDetailResponse  `json:"cart"`
            Success bool  `json:"success"`
         
    }
    
    // OverrideCartItemPromo used by Cart
    type OverrideCartItemPromo struct {

        
            PromoDesc string  `json:"promo_desc"`
            RwrdTndr string  `json:"rwrd_tndr"`
            ItemList []map[string]interface{}  `json:"item_list"`
            PromoAmount string  `json:"promo_amount"`
            PromoID string  `json:"promo_id"`
         
    }
    
    // OverrideCartItem used by Cart
    type OverrideCartItem struct {

        
            Quantity float64  `json:"quantity"`
            Discount float64  `json:"discount"`
            Size string  `json:"size"`
            PriceEffective float64  `json:"price_effective"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            AmountPaid float64  `json:"amount_paid"`
            PriceMarked float64  `json:"price_marked"`
            ItemID float64  `json:"item_id"`
            PromoList []OverrideCartItemPromo  `json:"promo_list"`
            SellerIdentifier string  `json:"seller_identifier"`
         
    }
    
    // OverrideCheckoutReq used by Cart
    type OverrideCheckoutReq struct {

        
            OrderingStore float64  `json:"ordering_store"`
            PaymentMode string  `json:"payment_mode"`
            PaymentIdentifier string  `json:"payment_identifier"`
            OrderType string  `json:"order_type"`
            ShippingAddress map[string]interface{}  `json:"shipping_address"`
            CartID string  `json:"cart_id"`
            CartItems []OverrideCartItem  `json:"cart_items"`
            MerchantCode string  `json:"merchant_code"`
            Aggregator string  `json:"aggregator"`
            CurrencyCode string  `json:"currency_code"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
         
    }
    
    // OverrideCheckoutResponse used by Cart
    type OverrideCheckoutResponse struct {

        
            Data map[string]interface{}  `json:"data"`
            Cart map[string]interface{}  `json:"cart"`
            Success string  `json:"success"`
            Message string  `json:"message"`
            OrderID string  `json:"order_id"`
         
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

        
            Meta map[string]interface{}  `json:"meta"`
            Source map[string]interface{}  `json:"source"`
            Token string  `json:"token"`
            User map[string]interface{}  `json:"user"`
            CreatedOn string  `json:"created_on"`
         
    }
    
    // SharedCart used by Cart
    type SharedCart struct {

        
            BreakupValues CartBreakup  `json:"breakup_values"`
            Items []CartProductInfo  `json:"items"`
            CheckoutMode string  `json:"checkout_mode"`
            SharedCartDetails SharedCartDetails  `json:"shared_cart_details"`
            ID string  `json:"id"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            BuyNow bool  `json:"buy_now"`
            IsValid bool  `json:"is_valid"`
            Currency CartCurrency  `json:"currency"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            Message string  `json:"message"`
            CouponText string  `json:"coupon_text"`
            UID string  `json:"uid"`
            LastModified string  `json:"last_modified"`
            Gstin string  `json:"gstin"`
            Comment string  `json:"comment"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            CartID float64  `json:"cart_id"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
         
    }
    
    // SharedCartResponse used by Cart
    type SharedCartResponse struct {

        
            Cart SharedCart  `json:"cart"`
            Error string  `json:"error"`
         
    }
    
    // CartList used by Cart
    type CartList struct {

        
            UserID string  `json:"user_id"`
            ItemCounts float64  `json:"item_counts"`
            CartID string  `json:"cart_id"`
            CartValue float64  `json:"cart_value"`
            PickUpCustomerDetails map[string]interface{}  `json:"pick_up_customer_details"`
            CreatedOn string  `json:"created_on"`
         
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

        
            ModifiedOn string  `json:"modified_on"`
            CreatedAt string  `json:"created_at"`
            Gender string  `json:"gender"`
            LastName string  `json:"last_name"`
            ExternalID string  `json:"external_id"`
            ID string  `json:"_id"`
            FirstName string  `json:"first_name"`
            UID string  `json:"uid"`
            Mobile string  `json:"mobile"`
         
    }
    
    // UserCartMappingResponse used by Cart
    type UserCartMappingResponse struct {

        
            BreakupValues CartBreakup  `json:"breakup_values"`
            Items []CartProductInfo  `json:"items"`
            CheckoutMode string  `json:"checkout_mode"`
            ID string  `json:"id"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            BuyNow bool  `json:"buy_now"`
            IsValid bool  `json:"is_valid"`
            Currency CartCurrency  `json:"currency"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            Message string  `json:"message"`
            CouponText string  `json:"coupon_text"`
            LastModified string  `json:"last_modified"`
            Gstin string  `json:"gstin"`
            Comment string  `json:"comment"`
            AppliedPromoDetails []AppliedPromotion  `json:"applied_promo_details"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            PanNo string  `json:"pan_no"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            User UserInfo  `json:"user"`
            PanConfig map[string]interface{}  `json:"pan_config"`
         
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

        
            MinimumCartValue float64  `json:"minimum_cart_value"`
            Title string  `json:"title"`
            Message string  `json:"message"`
            CouponCode string  `json:"coupon_code"`
            CouponType string  `json:"coupon_type"`
            SubTitle string  `json:"sub_title"`
            IsApplied bool  `json:"is_applied"`
            MaxDiscountValue float64  `json:"max_discount_value"`
            Description string  `json:"description"`
            IsApplicable bool  `json:"is_applicable"`
            ExpiresOn string  `json:"expires_on"`
            CouponValue float64  `json:"coupon_value"`
         
    }
    
    // PageCoupon used by Cart
    type PageCoupon struct {

        
            TotalItemCount float64  `json:"total_item_count"`
            Total float64  `json:"total"`
            HasPrevious bool  `json:"has_previous"`
            HasNext bool  `json:"has_next"`
            Current float64  `json:"current"`
         
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

        
            UserID string  `json:"user_id"`
            CheckoutMode string  `json:"checkout_mode"`
            GeoLocation GeoLocation  `json:"geo_location"`
            ID string  `json:"id"`
            AreaCodeSlug string  `json:"area_code_slug"`
            CountryCode string  `json:"country_code"`
            Tags []string  `json:"tags"`
            IsActive bool  `json:"is_active"`
            AddressType string  `json:"address_type"`
            Name string  `json:"name"`
            State string  `json:"state"`
            Landmark string  `json:"landmark"`
            Address string  `json:"address"`
            Email string  `json:"email"`
            IsDefaultAddress bool  `json:"is_default_address"`
            GoogleMapPoint map[string]interface{}  `json:"google_map_point"`
            Meta map[string]interface{}  `json:"meta"`
            CartID string  `json:"cart_id"`
            Phone string  `json:"phone"`
            City string  `json:"city"`
            AreaCode string  `json:"area_code"`
            Country string  `json:"country"`
            Area string  `json:"area"`
            CreatedByUserID string  `json:"created_by_user_id"`
         
    }
    
    // PlatformGetAddressesResponse used by Cart
    type PlatformGetAddressesResponse struct {

        
            Address []PlatformAddress  `json:"address"`
         
    }
    
    // SaveAddressResponse used by Cart
    type SaveAddressResponse struct {

        
            ID string  `json:"id"`
            Success bool  `json:"success"`
            IsDefaultAddress bool  `json:"is_default_address"`
         
    }
    
    // UpdateAddressResponse used by Cart
    type UpdateAddressResponse struct {

        
            ID string  `json:"id"`
            Success bool  `json:"success"`
            IsDefaultAddress bool  `json:"is_default_address"`
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
            BillingAddressID string  `json:"billing_address_id"`
            CheckoutMode string  `json:"checkout_mode"`
            CartID string  `json:"cart_id"`
            ID string  `json:"id"`
         
    }
    
    // ShipmentArticle used by Cart
    type ShipmentArticle struct {

        
            ArticleID string  `json:"article_id"`
            Meta string  `json:"meta"`
            Quantity string  `json:"quantity"`
         
    }
    
    // PlatformShipmentResponse used by Cart
    type PlatformShipmentResponse struct {

        
            Items []CartProductInfo  `json:"items"`
            FulfillmentID float64  `json:"fulfillment_id"`
            Shipments float64  `json:"shipments"`
            DpOptions map[string]interface{}  `json:"dp_options"`
            OrderType string  `json:"order_type"`
            FulfillmentType string  `json:"fulfillment_type"`
            Articles []ShipmentArticle  `json:"articles"`
            DpID string  `json:"dp_id"`
            Promise ShipmentPromise  `json:"promise"`
            BoxType string  `json:"box_type"`
            ShipmentType string  `json:"shipment_type"`
         
    }
    
    // PlatformCartShipmentsResponse used by Cart
    type PlatformCartShipmentsResponse struct {

        
            BreakupValues CartBreakup  `json:"breakup_values"`
            Items []CartProductInfo  `json:"items"`
            CheckoutMode string  `json:"checkout_mode"`
            ID string  `json:"id"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            BuyNow bool  `json:"buy_now"`
            Shipments []PlatformShipmentResponse  `json:"shipments"`
            IsValid bool  `json:"is_valid"`
            Currency CartCurrency  `json:"currency"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            Message string  `json:"message"`
            CouponText string  `json:"coupon_text"`
            StaffUserID string  `json:"staff_user_id"`
            Error bool  `json:"error"`
            LastModified string  `json:"last_modified"`
            Gstin string  `json:"gstin"`
            Comment string  `json:"comment"`
            AppliedPromoDetails []AppliedPromotion  `json:"applied_promo_details"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            PanNo string  `json:"pan_no"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            PanConfig map[string]interface{}  `json:"pan_config"`
         
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

        
            Gstin string  `json:"gstin"`
            Comment string  `json:"comment"`
            CheckoutMode string  `json:"checkout_mode"`
            GiftDetails map[string]interface{}  `json:"gift_details"`
            PanNo string  `json:"pan_no"`
            StaffUserID string  `json:"staff_user_id"`
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

        
            LastName string  `json:"last_name"`
            ID string  `json:"_id"`
            FirstName string  `json:"first_name"`
            EmployeeCode string  `json:"employee_code"`
            User string  `json:"user"`
         
    }
    
    // Files used by Cart
    type Files struct {

        
            Values []string  `json:"values"`
            Key string  `json:"key"`
         
    }
    
    // PlatformCartCheckoutDetailRequest used by Cart
    type PlatformCartCheckoutDetailRequest struct {

        
            UserID string  `json:"user_id"`
            PickAtStoreUID float64  `json:"pick_at_store_uid"`
            OrderType string  `json:"order_type"`
            CheckoutMode string  `json:"checkout_mode"`
            Pos bool  `json:"pos"`
            EmployeeCode string  `json:"employee_code"`
            ID string  `json:"id"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
            PaymentMode string  `json:"payment_mode"`
            PaymentIdentifier string  `json:"payment_identifier"`
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            Staff StaffCheckout  `json:"staff"`
            MerchantCode string  `json:"merchant_code"`
            OrderingStore float64  `json:"ordering_store"`
            PaymentAutoConfirm bool  `json:"payment_auto_confirm"`
            DeviceID string  `json:"device_id"`
            AddressID string  `json:"address_id"`
            Meta map[string]interface{}  `json:"meta"`
            BillingAddressID string  `json:"billing_address_id"`
            Files []Files  `json:"files"`
            PaymentParams map[string]interface{}  `json:"payment_params"`
            CallbackURL string  `json:"callback_url"`
            Aggregator string  `json:"aggregator"`
         
    }
    
    // CheckCart used by Cart
    type CheckCart struct {

        
            BreakupValues CartBreakup  `json:"breakup_values"`
            Items []CartProductInfo  `json:"items"`
            CheckoutMode string  `json:"checkout_mode"`
            Success bool  `json:"success"`
            ID string  `json:"id"`
            DeliveryCharges float64  `json:"delivery_charges"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            UserType string  `json:"user_type"`
            StoreEmps []map[string]interface{}  `json:"store_emps"`
            StoreCode string  `json:"store_code"`
            BuyNow bool  `json:"buy_now"`
            CodAvailable bool  `json:"cod_available"`
            CodMessage string  `json:"cod_message"`
            IsValid bool  `json:"is_valid"`
            Currency CartCurrency  `json:"currency"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            Message string  `json:"message"`
            CouponText string  `json:"coupon_text"`
            ErrorMessage string  `json:"error_message"`
            UID string  `json:"uid"`
            LastModified string  `json:"last_modified"`
            Gstin string  `json:"gstin"`
            Comment string  `json:"comment"`
            DeliveryChargeOrderValue float64  `json:"delivery_charge_order_value"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            CartID float64  `json:"cart_id"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            CodCharges float64  `json:"cod_charges"`
            OrderID string  `json:"order_id"`
         
    }
    
    // CartCheckoutResponse used by Cart
    type CartCheckoutResponse struct {

        
            Data map[string]interface{}  `json:"data"`
            Cart CheckCart  `json:"cart"`
            Success bool  `json:"success"`
            AppInterceptURL string  `json:"app_intercept_url"`
            Message string  `json:"message"`
            OrderID string  `json:"order_id"`
            CallbackURL string  `json:"callback_url"`
            PaymentConfirmURL string  `json:"payment_confirm_url"`
         
    }
    
    // CartDeliveryModesResponse used by Cart
    type CartDeliveryModesResponse struct {

        
            PickupStores []float64  `json:"pickup_stores"`
            AvailableModes []string  `json:"available_modes"`
         
    }
    
    // PickupStoreDetail used by Cart
    type PickupStoreDetail struct {

        
            UID float64  `json:"uid"`
            ID float64  `json:"id"`
            State string  `json:"state"`
            Pincode float64  `json:"pincode"`
            StoreCode string  `json:"store_code"`
            Landmark string  `json:"landmark"`
            Address string  `json:"address"`
            Area string  `json:"area"`
            Phone string  `json:"phone"`
            Country string  `json:"country"`
            City string  `json:"city"`
            AreaCodeSlug string  `json:"area_code_slug"`
            AreaCode string  `json:"area_code"`
            AddressType string  `json:"address_type"`
            Email string  `json:"email"`
            Name string  `json:"name"`
         
    }
    
    // StoreDetailsResponse used by Cart
    type StoreDetailsResponse struct {

        
            Items []PickupStoreDetail  `json:"items"`
         
    }
    
    // UpdateCartPaymentRequest used by Cart
    type UpdateCartPaymentRequest struct {

        
            AggregatorName string  `json:"aggregator_name"`
            PaymentMode string  `json:"payment_mode"`
            PaymentIdentifier string  `json:"payment_identifier"`
            MerchantCode string  `json:"merchant_code"`
            ID string  `json:"id"`
            AddressID string  `json:"address_id"`
         
    }
    
    // CouponValidity used by Cart
    type CouponValidity struct {

        
            NextValidationRequired bool  `json:"next_validation_required"`
            Title string  `json:"title"`
            DisplayMessageEn string  `json:"display_message_en"`
            Code string  `json:"code"`
            Valid bool  `json:"valid"`
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

        
            PaymentGateway string  `json:"payment_gateway"`
            PaymentIdentifier string  `json:"payment_identifier"`
            MerchantCode string  `json:"merchant_code"`
            Type string  `json:"type"`
         
    }
    
    // PaymentMethod used by Cart
    type PaymentMethod struct {

        
            Payment string  `json:"payment"`
            Mode string  `json:"mode"`
            PaymentMeta PaymentMeta  `json:"payment_meta"`
            Name string  `json:"name"`
            Amount float64  `json:"amount"`
         
    }
    
    // PlatformCartCheckoutDetailV2Request used by Cart
    type PlatformCartCheckoutDetailV2Request struct {

        
            UserID string  `json:"user_id"`
            PickAtStoreUID float64  `json:"pick_at_store_uid"`
            OrderType string  `json:"order_type"`
            CheckoutMode string  `json:"checkout_mode"`
            Pos bool  `json:"pos"`
            EmployeeCode string  `json:"employee_code"`
            ID string  `json:"id"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
            PaymentMode string  `json:"payment_mode"`
            PaymentIdentifier string  `json:"payment_identifier"`
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            Staff StaffCheckout  `json:"staff"`
            MerchantCode string  `json:"merchant_code"`
            OrderingStore float64  `json:"ordering_store"`
            PaymentAutoConfirm bool  `json:"payment_auto_confirm"`
            DeviceID string  `json:"device_id"`
            AddressID string  `json:"address_id"`
            PaymentMethods []PaymentMethod  `json:"payment_methods"`
            Meta map[string]interface{}  `json:"meta"`
            BillingAddressID string  `json:"billing_address_id"`
            Files []Files  `json:"files"`
            PaymentParams map[string]interface{}  `json:"payment_params"`
            CustomMeta map[string]interface{}  `json:"custom_meta"`
            CallbackURL string  `json:"callback_url"`
            Aggregator string  `json:"aggregator"`
         
    }
    
    // UpdateCartPaymentRequestV2 used by Cart
    type UpdateCartPaymentRequestV2 struct {

        
            PaymentMethods []PaymentMethod  `json:"payment_methods"`
            AggregatorName string  `json:"aggregator_name"`
            PaymentMode string  `json:"payment_mode"`
            PaymentIdentifier string  `json:"payment_identifier"`
            MerchantCode string  `json:"merchant_code"`
            ID string  `json:"id"`
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
            Success bool  `json:"success"`
            Data ApplicationServiceabilityConfig  `json:"data"`
         
    }
    
    // EntityRegionView_Request used by Serviceability
    type EntityRegionView_Request struct {

        
            SubType []string  `json:"sub_type"`
            ParentID []string  `json:"parent_id"`
         
    }
    
    // EntityRegionView_Error used by Serviceability
    type EntityRegionView_Error struct {

        
            Value string  `json:"value"`
            Type string  `json:"type"`
            Message string  `json:"message"`
         
    }
    
    // EntityRegionView_page used by Serviceability
    type EntityRegionView_page struct {

        
            HasNext bool  `json:"has_next"`
            ItemTotal float64  `json:"item_total"`
            Current float64  `json:"current"`
            Type string  `json:"type"`
            Size float64  `json:"size"`
         
    }
    
    // EntityRegionView_Items used by Serviceability
    type EntityRegionView_Items struct {

        
            UID string  `json:"uid"`
            Name string  `json:"name"`
            SubType string  `json:"sub_type"`
         
    }
    
    // EntityRegionView_Response used by Serviceability
    type EntityRegionView_Response struct {

        
            Error EntityRegionView_Error  `json:"error"`
            Success bool  `json:"success"`
            Page EntityRegionView_page  `json:"page"`
            Data []EntityRegionView_Items  `json:"data"`
         
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

        
            ZoneID string  `json:"zone_id"`
            Name string  `json:"name"`
            Product ListViewProduct  `json:"product"`
            PincodesCount float64  `json:"pincodes_count"`
            IsActive bool  `json:"is_active"`
            CompanyID float64  `json:"company_id"`
            Slug string  `json:"slug"`
            StoresCount float64  `json:"stores_count"`
            Channels ListViewChannels  `json:"channels"`
         
    }
    
    // ListViewSummary used by Serviceability
    type ListViewSummary struct {

        
            TotalPincodesServed float64  `json:"total_pincodes_served"`
            TotalZones float64  `json:"total_zones"`
            TotalActiveZones float64  `json:"total_active_zones"`
         
    }
    
    // ZoneDataItem used by Serviceability
    type ZoneDataItem struct {

        
            HasNext bool  `json:"has_next"`
            ItemTotal float64  `json:"item_total"`
            Current float64  `json:"current"`
            Type string  `json:"type"`
            Size float64  `json:"size"`
         
    }
    
    // ListViewResponse used by Serviceability
    type ListViewResponse struct {

        
            Items []ListViewItems  `json:"items"`
            Summary []ListViewSummary  `json:"summary"`
            Page []ZoneDataItem  `json:"page"`
         
    }
    
    // CompanyStoreView_PageItems used by Serviceability
    type CompanyStoreView_PageItems struct {

        
            HasNext bool  `json:"has_next"`
            ItemTotal float64  `json:"item_total"`
            Current float64  `json:"current"`
            Type string  `json:"type"`
            Size float64  `json:"size"`
         
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

        
            State []string  `json:"state"`
            Country string  `json:"country"`
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

        
            StatusCode float64  `json:"status_code"`
            Success bool  `json:"success"`
         
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
            StatusCode float64  `json:"status_code"`
            Success bool  `json:"success"`
         
    }
    
    // GetZoneFromApplicationIdViewResponse used by Serviceability
    type GetZoneFromApplicationIdViewResponse struct {

        
            Items []ListViewItems  `json:"items"`
            Page []ZoneDataItem  `json:"page"`
         
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
    
    // MobileNo used by Serviceability
    type MobileNo struct {

        
            CountryCode float64  `json:"country_code"`
            Number string  `json:"number"`
         
    }
    
    // ManagerResponse used by Serviceability
    type ManagerResponse struct {

        
            MobileNo MobileNo  `json:"mobile_no"`
            Name string  `json:"name"`
            Email string  `json:"email"`
         
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
    
    // ContactNumberResponse used by Serviceability
    type ContactNumberResponse struct {

        
            CountryCode float64  `json:"country_code"`
            Number string  `json:"number"`
         
    }
    
    // DocumentsResponse used by Serviceability
    type DocumentsResponse struct {

        
            LegalName string  `json:"legal_name"`
            Value string  `json:"value"`
            Verified bool  `json:"verified"`
            Type string  `json:"type"`
         
    }
    
    // AddressResponse used by Serviceability
    type AddressResponse struct {

        
            State string  `json:"state"`
            Country string  `json:"country"`
            Landmark string  `json:"landmark"`
            Address2 string  `json:"address2"`
            City string  `json:"city"`
            Address1 string  `json:"address1"`
            Pincode float64  `json:"pincode"`
            Latitude float64  `json:"latitude"`
            Longitude float64  `json:"longitude"`
         
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
    
    // IntegrationTypeResponse used by Serviceability
    type IntegrationTypeResponse struct {

        
            Inventory string  `json:"inventory"`
            Order string  `json:"order"`
         
    }
    
    // Dp used by Serviceability
    type Dp struct {

        
            Operations []string  `json:"operations"`
            TransportMode string  `json:"transport_mode"`
            ExternalAccountID string  `json:"external_account_id"`
            AssignDpFromSb bool  `json:"assign_dp_from_sb"`
            PaymentMode string  `json:"payment_mode"`
            InternalAccountID string  `json:"internal_account_id"`
            AreaCode float64  `json:"area_code"`
            LmPriority float64  `json:"lm_priority"`
            RvpPriority float64  `json:"rvp_priority"`
            FmPriority float64  `json:"fm_priority"`
         
    }
    
    // LogisticsResponse used by Serviceability
    type LogisticsResponse struct {

        
            Override bool  `json:"override"`
            Dp Dp  `json:"dp"`
         
    }
    
    // ModifiedByResponse used by Serviceability
    type ModifiedByResponse struct {

        
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
         
    }
    
    // OpeningClosing used by Serviceability
    type OpeningClosing struct {

        
            Minute float64  `json:"minute"`
            Hour float64  `json:"hour"`
         
    }
    
    // TimmingResponse used by Serviceability
    type TimmingResponse struct {

        
            Opening OpeningClosing  `json:"opening"`
            Closing OpeningClosing  `json:"closing"`
            Open bool  `json:"open"`
            Weekday string  `json:"weekday"`
         
    }
    
    // ProductReturnConfigResponse used by Serviceability
    type ProductReturnConfigResponse struct {

        
            OnSameStore bool  `json:"on_same_store"`
         
    }
    
    // ItemResponse used by Serviceability
    type ItemResponse struct {

        
            Manager ManagerResponse  `json:"manager"`
            CreatedBy CreatedByResponse  `json:"created_by"`
            Company float64  `json:"company"`
            DisplayName string  `json:"display_name"`
            NotificationEmails []string  `json:"notification_emails"`
            ModifiedOn string  `json:"modified_on"`
            Warnings WarningsResponse  `json:"warnings"`
            Cls string  `json:"_cls"`
            CreatedOn string  `json:"created_on"`
            ContactNumbers []ContactNumberResponse  `json:"contact_numbers"`
            UID float64  `json:"uid"`
            Documents []DocumentsResponse  `json:"documents"`
            Address AddressResponse  `json:"address"`
            GstCredentials GstCredentialsResponse  `json:"gst_credentials"`
            IntegrationType IntegrationTypeResponse  `json:"integration_type"`
            Logistics LogisticsResponse  `json:"logistics"`
            SubType string  `json:"sub_type"`
            VerifiedBy ModifiedByResponse  `json:"verified_by"`
            StoreType string  `json:"store_type"`
            Stage string  `json:"stage"`
            Name string  `json:"name"`
            Timing []TimmingResponse  `json:"timing"`
            VerifiedOn string  `json:"verified_on"`
            CompanyID float64  `json:"company_id"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ModifiedBy ModifiedByResponse  `json:"modified_by"`
            Code string  `json:"code"`
            ProductReturnConfig ProductReturnConfigResponse  `json:"product_return_config"`
         
    }
    
    // PageResponse used by Serviceability
    type PageResponse struct {

        
            HasNext bool  `json:"has_next"`
            ItemTotal float64  `json:"item_total"`
            Current float64  `json:"current"`
            Type string  `json:"type"`
            Size float64  `json:"size"`
         
    }
    
    // GetStoresViewResponse used by Serviceability
    type GetStoresViewResponse struct {

        
            Items []ItemResponse  `json:"items"`
            Page PageResponse  `json:"page"`
         
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

        
            StatusCode string  `json:"status_code"`
            Success string  `json:"success"`
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

        
            HasNext bool  `json:"has_next"`
            ItemTotal float64  `json:"item_total"`
            Current float64  `json:"current"`
            Type string  `json:"type"`
            Size float64  `json:"size"`
         
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
    

    
    // GenerateReportFilters used by Finance
    type GenerateReportFilters struct {

        
            Channel []string  `json:"channel"`
            Brand []string  `json:"brand"`
            Company []string  `json:"company"`
         
    }
    
    // GenerateReportMeta used by Finance
    type GenerateReportMeta struct {

        
            Channel string  `json:"channel"`
            Brand string  `json:"brand"`
            Company string  `json:"company"`
         
    }
    
    // GenerateReportPlatform used by Finance
    type GenerateReportPlatform struct {

        
            EndDate string  `json:"end_date"`
            StartDate string  `json:"start_date"`
            ReportID string  `json:"report_id"`
            Filters GenerateReportFilters  `json:"filters"`
            Meta GenerateReportMeta  `json:"meta"`
         
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

        
            EndDate string  `json:"end_date"`
            StartDate string  `json:"start_date"`
            Items [][]string  `json:"items"`
            ItemCount float64  `json:"item_count"`
            Headers []string  `json:"headers"`
            Page Page  `json:"page"`
         
    }
    
    // Error used by Finance
    type Error struct {

        
            Reason string  `json:"reason"`
            Success bool  `json:"success"`
         
    }
    
    // DownloadReport used by Finance
    type DownloadReport struct {

        
            EndDate string  `json:"end_date"`
            Pagesize float64  `json:"pagesize"`
            StartDate string  `json:"start_date"`
            Page float64  `json:"page"`
         
    }
    
    // DownloadReportItems used by Finance
    type DownloadReportItems struct {

        
            EndDate string  `json:"end_date"`
            StartDate string  `json:"start_date"`
            TypeOfRequest string  `json:"type_of_request"`
            Meta GenerateReportMeta  `json:"meta"`
            Filters GenerateReportFilters  `json:"filters"`
            ReportID string  `json:"report_id"`
         
    }
    
    // DownloadReportList used by Finance
    type DownloadReportList struct {

        
            ItemCount float64  `json:"item_count"`
            Page Page  `json:"page"`
            Items []DownloadReportItems  `json:"items"`
         
    }
    
    // GetEngineFilters used by Finance
    type GetEngineFilters struct {

        
            ConfigField string  `json:"config_field"`
         
    }
    
    // GetEngineData used by Finance
    type GetEngineData struct {

        
            Project []string  `json:"project"`
            Filters GetEngineFilters  `json:"filters"`
            TableName string  `json:"table_name"`
         
    }
    
    // GetEngineRequest used by Finance
    type GetEngineRequest struct {

        
            Data GetEngineData  `json:"data"`
         
    }
    
    // GetEngineResponse used by Finance
    type GetEngineResponse struct {

        
            Page Page  `json:"page"`
            ItemCount float64  `json:"item_count"`
            Success bool  `json:"success"`
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

        
            ModeOfPayment string  `json:"mode_of_payment"`
            TransactionType string  `json:"transaction_type"`
            SellerID string  `json:"seller_id"`
            Currency string  `json:"currency"`
            TotalAmount string  `json:"total_amount"`
            InvoiceNumber string  `json:"invoice_number"`
            Platform string  `json:"platform"`
            SourceReference string  `json:"source_reference"`
            Amount string  `json:"amount"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // PaymentProcessRequest used by Finance
    type PaymentProcessRequest struct {

        
            Data PaymentProcessPayload  `json:"data"`
         
    }
    
    // PaymentProcessResponse used by Finance
    type PaymentProcessResponse struct {

        
            TransactionID string  `json:"transaction_id"`
            Code float64  `json:"code"`
            RedirectURL string  `json:"redirect_url"`
            Message string  `json:"message"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // CreditlineDataPlatformPayload used by Finance
    type CreditlineDataPlatformPayload struct {

        
            StartEnd string  `json:"start_end"`
            SellerID string  `json:"seller_id"`
            EndEnd string  `json:"end_end"`
            Pagesize float64  `json:"pagesize"`
            Page float64  `json:"page"`
         
    }
    
    // CreditlineDataPlatformRequest used by Finance
    type CreditlineDataPlatformRequest struct {

        
            Data CreditlineDataPlatformPayload  `json:"data"`
         
    }
    
    // CreditlineDataPlatformResponse used by Finance
    type CreditlineDataPlatformResponse struct {

        
            Code float64  `json:"code"`
            Items []map[string]interface{}  `json:"items"`
            Message string  `json:"message"`
            ShowMr bool  `json:"show_mr"`
            ItemCount float64  `json:"item_count"`
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

        
            InvoiceTypeList []InvoiceTypeResponseItems  `json:"invoice_type_list"`
            Success bool  `json:"success"`
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

        
            EndDate string  `json:"end_date"`
            StartDate string  `json:"start_date"`
            PageSize float64  `json:"page_size"`
            Search string  `json:"search"`
            Filters InoviceListingPayloadDataFilters  `json:"filters"`
            Page float64  `json:"page"`
         
    }
    
    // InvoiceListingRequest used by Finance
    type InvoiceListingRequest struct {

        
            Data InvoiceListingPayloadData  `json:"data"`
         
    }
    
    // UnpaidInvoiceDataItems used by Finance
    type UnpaidInvoiceDataItems struct {

        
            TotalUnpaidAmount float64  `json:"total_unpaid_amount"`
            TotalUnpaidInvoiceCount float64  `json:"total_unpaid_invoice_count"`
            Currency string  `json:"currency"`
         
    }
    
    // InvoiceListingResponseItems used by Finance
    type InvoiceListingResponseItems struct {

        
            InvoiceDate string  `json:"invoice_date"`
            DueDate string  `json:"due_date"`
            InvoiceNumber string  `json:"invoice_number"`
            IsDownloadable bool  `json:"is_downloadable"`
            Period string  `json:"period"`
            InvoiceID string  `json:"invoice_id"`
            Status string  `json:"status"`
            InvoiceType string  `json:"invoice_type"`
            Amount string  `json:"amount"`
            Company string  `json:"company"`
         
    }
    
    // InvoiceListingResponse used by Finance
    type InvoiceListingResponse struct {

        
            UnpaidInvoiceData UnpaidInvoiceDataItems  `json:"unpaid_invoice_data"`
            ItemCount float64  `json:"item_count"`
            Page Page  `json:"page"`
            Items []InvoiceListingResponseItems  `json:"items"`
         
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
    
