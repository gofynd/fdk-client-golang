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
    

    
    // GenerateSEOContent used by Content
    type GenerateSEOContent struct {

        
            Text string  `json:"text"`
         
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

        
            MandateAmount float64  `json:"mandate_amount"`
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

        
            Created bool  `json:"created"`
            AppID string  `json:"app_id"`
            Aggregators []map[string]interface{}  `json:"aggregators"`
            ExcludedFields []string  `json:"excluded_fields"`
            DisplayFields []string  `json:"display_fields"`
            Success bool  `json:"success"`
         
    }
    
    // ErrorCodeDescription used by Payment
    type ErrorCodeDescription struct {

        
            Success bool  `json:"success"`
            Description string  `json:"description"`
            Code string  `json:"code"`
         
    }
    
    // PaymentGatewayConfig used by Payment
    type PaymentGatewayConfig struct {

        
            ConfigType string  `json:"config_type"`
            IsActive bool  `json:"is_active"`
            Secret string  `json:"secret"`
            Key string  `json:"key"`
            MerchantSalt string  `json:"merchant_salt"`
         
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

        
            CardBrandImage string  `json:"card_brand_image"`
            CardFingerprint string  `json:"card_fingerprint"`
            MerchantCode string  `json:"merchant_code"`
            IntentApp []IntentApp  `json:"intent_app"`
            IntentAppErrorDictList []IntentAppErrorList  `json:"intent_app_error_dict_list"`
            CardType string  `json:"card_type"`
            Timeout float64  `json:"timeout"`
            AggregatorName string  `json:"aggregator_name"`
            CardReference string  `json:"card_reference"`
            CodLimitPerOrder float64  `json:"cod_limit_per_order"`
            Expired bool  `json:"expired"`
            ExpMonth float64  `json:"exp_month"`
            IntentFlow bool  `json:"intent_flow"`
            CardNumber string  `json:"card_number"`
            DisplayPriority float64  `json:"display_priority"`
            IntentAppErrorList []string  `json:"intent_app_error_list"`
            RetryCount float64  `json:"retry_count"`
            CardID string  `json:"card_id"`
            RemainingLimit float64  `json:"remaining_limit"`
            CardIssuer string  `json:"card_issuer"`
            CodLimit float64  `json:"cod_limit"`
            ExpYear float64  `json:"exp_year"`
            CompliantWithTokenisationGuidelines bool  `json:"compliant_with_tokenisation_guidelines"`
            CardName string  `json:"card_name"`
            Nickname string  `json:"nickname"`
            CardBrand string  `json:"card_brand"`
            DisplayName string  `json:"display_name"`
            CardToken string  `json:"card_token"`
            LogoURL PaymentModeLogo  `json:"logo_url"`
            Code string  `json:"code"`
            CardIsin string  `json:"card_isin"`
            FyndVpa string  `json:"fynd_vpa"`
            Name string  `json:"name"`
         
    }
    
    // RootPaymentMode used by Payment
    type RootPaymentMode struct {

        
            AggregatorName string  `json:"aggregator_name"`
            DisplayName string  `json:"display_name"`
            IsPayByCardPl bool  `json:"is_pay_by_card_pl"`
            DisplayPriority float64  `json:"display_priority"`
            List []PaymentModeList  `json:"list"`
            SaveCard bool  `json:"save_card"`
            AddCardEnabled bool  `json:"add_card_enabled"`
            AnonymousEnable bool  `json:"anonymous_enable"`
            Name string  `json:"name"`
         
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
    
    // PayoutAggregator used by Payment
    type PayoutAggregator struct {

        
            PayoutDetailsID float64  `json:"payout_details_id"`
            AggregatorFundID float64  `json:"aggregator_fund_id"`
            AggregatorID float64  `json:"aggregator_id"`
         
    }
    
    // PayoutMoreAttributes used by Payment
    type PayoutMoreAttributes struct {

        
            State string  `json:"state"`
            AccountHolder string  `json:"account_holder"`
            IfscCode string  `json:"ifsc_code"`
            BranchName string  `json:"branch_name"`
            Country string  `json:"country"`
            City string  `json:"city"`
            AccountNo string  `json:"account_no"`
            BankName string  `json:"bank_name"`
            AccountType string  `json:"account_type"`
         
    }
    
    // PayoutCustomer used by Payment
    type PayoutCustomer struct {

        
            ID float64  `json:"id"`
            Name string  `json:"name"`
            UniqueExternalID string  `json:"unique_external_id"`
            Mobile string  `json:"mobile"`
         
    }
    
    // Payout used by Payment
    type Payout struct {

        
            UniqueTransferNo string  `json:"unique_transfer_no"`
            PayoutsAggregators []PayoutAggregator  `json:"payouts_aggregators"`
            IsActive bool  `json:"is_active"`
            MoreAttributes PayoutMoreAttributes  `json:"more_attributes"`
            TransferType string  `json:"transfer_type"`
            Customers PayoutCustomer  `json:"customers"`
            IsDefault bool  `json:"is_default"`
         
    }
    
    // PayoutsResponse used by Payment
    type PayoutsResponse struct {

        
            Success bool  `json:"success"`
            Items []Payout  `json:"items"`
         
    }
    
    // PayoutBankDetails used by Payment
    type PayoutBankDetails struct {

        
            AccountHolder string  `json:"account_holder"`
            State string  `json:"state"`
            IfscCode string  `json:"ifsc_code"`
            BranchName string  `json:"branch_name"`
            Country string  `json:"country"`
            AccountNo string  `json:"account_no"`
            City string  `json:"city"`
            Pincode float64  `json:"pincode"`
            BankName string  `json:"bank_name"`
            AccountType string  `json:"account_type"`
         
    }
    
    // PayoutRequest used by Payment
    type PayoutRequest struct {

        
            Aggregator string  `json:"aggregator"`
            UniqueExternalID string  `json:"unique_external_id"`
            Users map[string]interface{}  `json:"users"`
            IsActive bool  `json:"is_active"`
            TransferType string  `json:"transfer_type"`
            BankDetails PayoutBankDetails  `json:"bank_details"`
         
    }
    
    // PayoutResponse used by Payment
    type PayoutResponse struct {

        
            Created bool  `json:"created"`
            Aggregator string  `json:"aggregator"`
            UniqueTransferNo string  `json:"unique_transfer_no"`
            Users map[string]interface{}  `json:"users"`
            IsActive bool  `json:"is_active"`
            Success bool  `json:"success"`
            Payouts map[string]interface{}  `json:"payouts"`
            TransferType string  `json:"transfer_type"`
            PaymentStatus string  `json:"payment_status"`
            BankDetails map[string]interface{}  `json:"bank_details"`
         
    }
    
    // UpdatePayoutResponse used by Payment
    type UpdatePayoutResponse struct {

        
            Success bool  `json:"success"`
            IsActive bool  `json:"is_active"`
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
            IsVerifiedFlag bool  `json:"is_verified_flag"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
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
            BranchName string  `json:"branch_name"`
            AccountNo string  `json:"account_no"`
            BankName string  `json:"bank_name"`
         
    }
    
    // AddBeneficiaryDetailsOTPRequest used by Payment
    type AddBeneficiaryDetailsOTPRequest struct {

        
            OrderID string  `json:"order_id"`
            Details BankDetailsForOTP  `json:"details"`
         
    }
    
    // IfscCodeResponse used by Payment
    type IfscCodeResponse struct {

        
            BranchName string  `json:"branch_name"`
            BankName string  `json:"bank_name"`
            Success bool  `json:"success"`
         
    }
    
    // OrderBeneficiaryDetails used by Payment
    type OrderBeneficiaryDetails struct {

        
            Comment string  `json:"comment"`
            AccountNo string  `json:"account_no"`
            ID float64  `json:"id"`
            Title string  `json:"title"`
            IsActive bool  `json:"is_active"`
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
            Subtitle string  `json:"subtitle"`
            AccountHolder string  `json:"account_holder"`
            IfscCode string  `json:"ifsc_code"`
            Mobile string  `json:"mobile"`
            BeneficiaryID string  `json:"beneficiary_id"`
            TransferMode string  `json:"transfer_mode"`
            Address string  `json:"address"`
            DisplayName string  `json:"display_name"`
            BranchName string  `json:"branch_name"`
            DelightsUserName string  `json:"delights_user_name"`
            Email string  `json:"email"`
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
            CurrentStatus string  `json:"current_status"`
            PaymentGateway string  `json:"payment_gateway"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            PaymentID string  `json:"payment_id"`
         
    }
    
    // MultiTenderPaymentMethod used by Payment
    type MultiTenderPaymentMethod struct {

        
            Name string  `json:"name"`
            Mode string  `json:"mode"`
            Amount float64  `json:"amount"`
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
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // CODdata used by Payment
    type CODdata struct {

        
            RemainingLimit float64  `json:"remaining_limit"`
            Limit float64  `json:"limit"`
            IsActive bool  `json:"is_active"`
            Usages float64  `json:"usages"`
            UserID string  `json:"user_id"`
         
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
    

    
    // FilterInfoOption used by Order
    type FilterInfoOption struct {

        
            Text string  `json:"text"`
            Value string  `json:"value"`
            Name string  `json:"name"`
         
    }
    
    // FiltersInfo used by Order
    type FiltersInfo struct {

        
            Type string  `json:"type"`
            Text string  `json:"text"`
            Value string  `json:"value"`
            Options []FilterInfoOption  `json:"options"`
         
    }
    
    // Dimensions used by Order
    type Dimensions struct {

        
            Unit string  `json:"unit"`
            Width float64  `json:"width"`
            IsDefault bool  `json:"is_default"`
            Height float64  `json:"height"`
            Length float64  `json:"length"`
         
    }
    
    // Meta used by Order
    type Meta struct {

        
            Dimension Dimensions  `json:"dimension"`
         
    }
    
    // DPDetailsData used by Order
    type DPDetailsData struct {

        
            Country string  `json:"country"`
            TrackURL string  `json:"track_url"`
            EwayBillNumber float64  `json:"eway_bill_number"`
            AwbNo string  `json:"awb_no"`
            EwayBillID string  `json:"eway_bill_id"`
            DpCharges float64  `json:"dp_charges"`
            GstTag string  `json:"gst_tag"`
            Pincode string  `json:"pincode"`
            Name string  `json:"name"`
            ID float64  `json:"id"`
            DpReturnCharges float64  `json:"dp_return_charges"`
            AmountHandlingCharges float64  `json:"amount_handling_charges"`
         
    }
    
    // UserDetailsData used by Order
    type UserDetailsData struct {

        
            Country string  `json:"country"`
            Email string  `json:"email"`
            Address string  `json:"address"`
            Pincode string  `json:"pincode"`
            Name string  `json:"name"`
            State string  `json:"state"`
            City string  `json:"city"`
            Phone string  `json:"phone"`
         
    }
    
    // ShipmentItemFulFillingStore used by Order
    type ShipmentItemFulFillingStore struct {

        
            Code string  `json:"code"`
            ID string  `json:"id"`
         
    }
    
    // Prices used by Order
    type Prices struct {

        
            AmountPaidRoundoff float64  `json:"amount_paid_roundoff"`
            RefundAmount float64  `json:"refund_amount"`
            AmountPaid float64  `json:"amount_paid"`
            CodCharges float64  `json:"cod_charges"`
            TransferPrice float64  `json:"transfer_price"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            CouponValue float64  `json:"coupon_value"`
            Cashback float64  `json:"cashback"`
            CashbackApplied float64  `json:"cashback_applied"`
            ValueOfGood float64  `json:"value_of_good"`
            PriceEffective float64  `json:"price_effective"`
            PromotionEffectiveDiscount float64  `json:"promotion_effective_discount"`
            PriceMarked float64  `json:"price_marked"`
            DeliveryCharge float64  `json:"delivery_charge"`
            RefundCredit float64  `json:"refund_credit"`
            Discount float64  `json:"discount"`
            FyndCredits float64  `json:"fynd_credits"`
         
    }
    
    // UserDataInfo used by Order
    type UserDataInfo struct {

        
            Mobile string  `json:"mobile"`
            Email string  `json:"email"`
            LastName string  `json:"last_name"`
            FirstName string  `json:"first_name"`
            Name string  `json:"name"`
            AvisUserID string  `json:"avis_user_id"`
            UID float64  `json:"uid"`
            IsAnonymousUser bool  `json:"is_anonymous_user"`
            Gender string  `json:"gender"`
         
    }
    
    // PlatformItem used by Order
    type PlatformItem struct {

        
            Image []string  `json:"image"`
            Images []string  `json:"images"`
            L3Category float64  `json:"l3_category"`
            Size string  `json:"size"`
            CanReturn bool  `json:"can_return"`
            DepartmentID float64  `json:"department_id"`
            Name string  `json:"name"`
            L3CategoryName string  `json:"l3_category_name"`
            Code string  `json:"code"`
            ID float64  `json:"id"`
            CanCancel bool  `json:"can_cancel"`
            L1Category []string  `json:"l1_category"`
            Color string  `json:"color"`
         
    }
    
    // GSTDetailsData used by Order
    type GSTDetailsData struct {

        
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            GstFee float64  `json:"gst_fee"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            ValueOfGood float64  `json:"value_of_good"`
            GstinCode string  `json:"gstin_code"`
         
    }
    
    // BagStateMapper used by Order
    type BagStateMapper struct {

        
            AppDisplayName string  `json:"app_display_name"`
            JourneyType string  `json:"journey_type"`
            AppStateName string  `json:"app_state_name"`
            BsID float64  `json:"bs_id"`
            Name string  `json:"name"`
            NotifyCustomer bool  `json:"notify_customer"`
            IsActive bool  `json:"is_active"`
            AppFacing bool  `json:"app_facing"`
            DisplayName string  `json:"display_name"`
            StateType string  `json:"state_type"`
         
    }
    
    // BagCurrentStatus used by Order
    type BagCurrentStatus struct {

        
            BagStateMapper BagStateMapper  `json:"bag_state_mapper"`
            CreatedAt string  `json:"created_at"`
            StateID float64  `json:"state_id"`
            StoreID float64  `json:"store_id"`
            ShipmentID string  `json:"shipment_id"`
            KafkaSync bool  `json:"kafka_sync"`
            Status string  `json:"status"`
            DeliveryAwbNumber string  `json:"delivery_awb_number"`
            ID float64  `json:"id"`
            UpdatedAt string  `json:"updated_at"`
            DeliveryPartnerID float64  `json:"delivery_partner_id"`
            BagID float64  `json:"bag_id"`
            StateType string  `json:"state_type"`
         
    }
    
    // BagUnit used by Order
    type BagUnit struct {

        
            Item PlatformItem  `json:"item"`
            OrderingChannel string  `json:"ordering_channel"`
            Prices Prices  `json:"prices"`
            Status map[string]interface{}  `json:"status"`
            ShipmentID string  `json:"shipment_id"`
            CanReturn bool  `json:"can_return"`
            Gst GSTDetailsData  `json:"gst"`
            TotalShipmentBags float64  `json:"total_shipment_bags"`
            CurrentStatus BagCurrentStatus  `json:"current_status"`
            CanCancel bool  `json:"can_cancel"`
            BagID float64  `json:"bag_id"`
            ItemQuantity float64  `json:"item_quantity"`
         
    }
    
    // EInvoice used by Order
    type EInvoice struct {

        
            ErrorCode string  `json:"error_code"`
            AcknowledgeNo float64  `json:"acknowledge_no"`
            ErrorMessage string  `json:"error_message"`
            SignedInvoice string  `json:"signed_invoice"`
            Irn string  `json:"irn"`
            AcknowledgeDate string  `json:"acknowledge_date"`
            SignedQrCode string  `json:"signed_qr_code"`
         
    }
    
    // EinvoiceInfo used by Order
    type EinvoiceInfo struct {

        
            CreditNote EInvoice  `json:"credit_note"`
            Invoice EInvoice  `json:"invoice"`
         
    }
    
    // ShipmentStatus used by Order
    type ShipmentStatus struct {

        
            CreatedAt string  `json:"created_at"`
            HexCode string  `json:"hex_code"`
            Title string  `json:"title"`
            Status string  `json:"status"`
            ActualStatus string  `json:"actual_status"`
            OpsStatus string  `json:"ops_status"`
         
    }
    
    // PDFLinks used by Order
    type PDFLinks struct {

        
            LabelType string  `json:"label_type"`
            CreditNoteURL string  `json:"credit_note_url"`
            LabelA6 string  `json:"label_a6"`
            DeliveryChallanA4 string  `json:"delivery_challan_a4"`
            InvoiceType string  `json:"invoice_type"`
            B2b string  `json:"b2b"`
            LabelA4 string  `json:"label_a4"`
            LabelPos string  `json:"label_pos"`
            InvoiceA6 string  `json:"invoice_a6"`
            InvoiceA4 string  `json:"invoice_a4"`
            PoInvoice string  `json:"po_invoice"`
            Label string  `json:"label"`
            InvoicePos string  `json:"invoice_pos"`
            Invoice string  `json:"invoice"`
         
    }
    
    // PaymentModeInfo used by Order
    type PaymentModeInfo struct {

        
            Type string  `json:"type"`
            Logo string  `json:"logo"`
         
    }
    
    // Formatted used by Order
    type Formatted struct {

        
            FMax string  `json:"f_max"`
            FMin string  `json:"f_min"`
         
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
    
    // LockData used by Order
    type LockData struct {

        
            Locked bool  `json:"locked"`
            Mto bool  `json:"mto"`
            LockMessage string  `json:"lock_message"`
         
    }
    
    // BuyerDetails used by Order
    type BuyerDetails struct {

        
            AjioSiteID string  `json:"ajio_site_id"`
            Address string  `json:"address"`
            Pincode float64  `json:"pincode"`
            Name string  `json:"name"`
            Gstin string  `json:"gstin"`
            State string  `json:"state"`
            City string  `json:"city"`
         
    }
    
    // ShipmentMeta used by Order
    type ShipmentMeta struct {

        
            Formatted Formatted  `json:"formatted"`
            SameStoreAvailable bool  `json:"same_store_available"`
            AwbNumber string  `json:"awb_number"`
            StoreInvoiceUpdatedDate string  `json:"store_invoice_updated_date"`
            EinvoiceInfo EinvoiceInfo  `json:"einvoice_info"`
            FulfilmentPriorityText string  `json:"fulfilment_priority_text"`
            ForwardAffiliateShipmentID string  `json:"forward_affiliate_shipment_id"`
            Weight float64  `json:"weight"`
            DebugInfo DebugInfo  `json:"debug_info"`
            OrderType string  `json:"order_type"`
            BagWeight map[string]interface{}  `json:"bag_weight"`
            PoNumber string  `json:"po_number"`
            Timestamp ShipmentTimeStamp  `json:"timestamp"`
            DueDate string  `json:"due_date"`
            ReturnAffiliateOrderID string  `json:"return_affiliate_order_id"`
            AutoTriggerDpAssignmentAcf bool  `json:"auto_trigger_dp_assignment_acf"`
            ShipmentWeight float64  `json:"shipment_weight"`
            MarketplaceStoreID string  `json:"marketplace_store_id"`
            ForwardAffiliateOrderID string  `json:"forward_affiliate_order_id"`
            ReturnAffiliateShipmentID string  `json:"return_affiliate_shipment_id"`
            LockData LockData  `json:"lock_data"`
            DpID string  `json:"dp_id"`
            EwaybillInfo map[string]interface{}  `json:"ewaybill_info"`
            ReturnStoreNode float64  `json:"return_store_node"`
            DpName string  `json:"dp_name"`
            ReturnAwbNumber string  `json:"return_awb_number"`
            BoxType string  `json:"box_type"`
            PackagingName string  `json:"packaging_name"`
            ReturnDetails map[string]interface{}  `json:"return_details"`
            DpSortKey string  `json:"dp_sort_key"`
            External map[string]interface{}  `json:"external"`
            ShipmentVolumetricWeight float64  `json:"shipment_volumetric_weight"`
            DpOptions map[string]interface{}  `json:"dp_options"`
            B2cBuyerDetails map[string]interface{}  `json:"b2c_buyer_details"`
            B2bBuyerDetails BuyerDetails  `json:"b2b_buyer_details"`
            AssignDpFromSb bool  `json:"assign_dp_from_sb"`
         
    }
    
    // AffiliateMeta used by Order
    type AffiliateMeta struct {

        
            BoxType string  `json:"box_type"`
            OrderItemID string  `json:"order_item_id"`
            ChannelShipmentID string  `json:"channel_shipment_id"`
            DueDate string  `json:"due_date"`
            SizeLevelTotalQty float64  `json:"size_level_total_qty"`
            IsPriority bool  `json:"is_priority"`
            Quantity float64  `json:"quantity"`
            ChannelOrderID string  `json:"channel_order_id"`
            CouponCode string  `json:"coupon_code"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            EmployeeDiscount float64  `json:"employee_discount"`
         
    }
    
    // AffiliateDetails used by Order
    type AffiliateDetails struct {

        
            AffiliateBagID string  `json:"affiliate_bag_id"`
            PdfLinks PDFLinks  `json:"pdf_links"`
            ShipmentMeta ShipmentMeta  `json:"shipment_meta"`
            CompanyAffiliateTag string  `json:"company_affiliate_tag"`
            AdID string  `json:"ad_id"`
            AffiliateStoreID string  `json:"affiliate_store_id"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            AffiliateID string  `json:"affiliate_id"`
            AffiliateMeta AffiliateMeta  `json:"affiliate_meta"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
         
    }
    
    // ShipmentItem used by Order
    type ShipmentItem struct {

        
            Meta Meta  `json:"meta"`
            Sla map[string]interface{}  `json:"sla"`
            ShipmentID string  `json:"shipment_id"`
            DpDetails DPDetailsData  `json:"dp_details"`
            DeliveryDetails UserDetailsData  `json:"delivery_details"`
            ID string  `json:"id"`
            BillingDetails UserDetailsData  `json:"billing_details"`
            ShipmentCreatedAt string  `json:"shipment_created_at"`
            FulfillingStore ShipmentItemFulFillingStore  `json:"fulfilling_store"`
            InvoiceID string  `json:"invoice_id"`
            FulfilmentPriority float64  `json:"fulfilment_priority"`
            TotalShipmentsInOrder float64  `json:"total_shipments_in_order"`
            PaymentMethods map[string]interface{}  `json:"payment_methods"`
            OrderingChannel string  `json:"ordering_channel"`
            LockStatus bool  `json:"lock_status"`
            Channel map[string]interface{}  `json:"channel"`
            Prices Prices  `json:"prices"`
            JourneyType string  `json:"journey_type"`
            User UserDataInfo  `json:"user"`
            Bags []BagUnit  `json:"bags"`
            Invoice EinvoiceInfo  `json:"invoice"`
            FulfillingCentre string  `json:"fulfilling_centre"`
            CreatedAt string  `json:"created_at"`
            ShipmentStatus ShipmentStatus  `json:"shipment_status"`
            PdfLinks PDFLinks  `json:"pdf_links"`
            Application map[string]interface{}  `json:"application"`
            PaymentModeInfo PaymentModeInfo  `json:"payment_mode_info"`
            Company map[string]interface{}  `json:"company"`
            AffiliateDetails AffiliateDetails  `json:"affiliate_details"`
            TotalBagsCount float64  `json:"total_bags_count"`
         
    }
    
    // ShipmentInternalPlatformViewResponse used by Order
    type ShipmentInternalPlatformViewResponse struct {

        
            Filters []FiltersInfo  `json:"filters"`
            Items []ShipmentItem  `json:"items"`
            AppliedFilters map[string]interface{}  `json:"applied_filters"`
            Page map[string]interface{}  `json:"page"`
         
    }
    
    // Error used by Order
    type Error struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // ShipmentPayments used by Order
    type ShipmentPayments struct {

        
            Source string  `json:"source"`
            Mode string  `json:"mode"`
            Logo string  `json:"logo"`
         
    }
    
    // TrackingList used by Order
    type TrackingList struct {

        
            IsPassed bool  `json:"is_passed"`
            Status string  `json:"status"`
            IsCurrent bool  `json:"is_current"`
            Text string  `json:"text"`
            Time string  `json:"time"`
         
    }
    
    // FulfillingStore used by Order
    type FulfillingStore struct {

        
            Country string  `json:"country"`
            Meta map[string]interface{}  `json:"meta"`
            Address string  `json:"address"`
            FulfillmentType string  `json:"fulfillment_type"`
            FulfillmentChannel string  `json:"fulfillment_channel"`
            Pincode string  `json:"pincode"`
            ContactPerson string  `json:"contact_person"`
            Code string  `json:"code"`
            StoreName string  `json:"store_name"`
            ID float64  `json:"id"`
            State string  `json:"state"`
            City string  `json:"city"`
            Phone string  `json:"phone"`
         
    }
    
    // ShipmentStatusData used by Order
    type ShipmentStatusData struct {

        
            CreatedAt string  `json:"created_at"`
            ShipmentID string  `json:"shipment_id"`
            Status string  `json:"status"`
            ID float64  `json:"id"`
            BagList []string  `json:"bag_list"`
         
    }
    
    // OrderingStoreDetails used by Order
    type OrderingStoreDetails struct {

        
            Country string  `json:"country"`
            OrderingStoreID float64  `json:"ordering_store_id"`
            Meta map[string]interface{}  `json:"meta"`
            Address string  `json:"address"`
            Pincode string  `json:"pincode"`
            ContactPerson string  `json:"contact_person"`
            Code string  `json:"code"`
            StoreName string  `json:"store_name"`
            State string  `json:"state"`
            City string  `json:"city"`
            Phone string  `json:"phone"`
         
    }
    
    // BagGST used by Order
    type BagGST struct {

        
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            GstFee float64  `json:"gst_fee"`
            IsDefaultHsnCode bool  `json:"is_default_hsn_code"`
            GstTag string  `json:"gst_tag"`
            ValueOfGood float64  `json:"value_of_good"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            GstinCode string  `json:"gstin_code"`
            HsnCode string  `json:"hsn_code"`
         
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

        
            ItemCriteria ItemCriterias  `json:"item_criteria"`
            CartConditions map[string]interface{}  `json:"cart_conditions"`
         
    }
    
    // AppliedPromos used by Order
    type AppliedPromos struct {

        
            Amount float64  `json:"amount"`
            PromotionName string  `json:"promotion_name"`
            ArticleQuantity float64  `json:"article_quantity"`
            PromoID string  `json:"promo_id"`
            MrpPromotion bool  `json:"mrp_promotion"`
            PromotionType string  `json:"promotion_type"`
            DiscountRules []DiscountRules  `json:"discount_rules"`
            BuyRules []BuyRules  `json:"buy_rules"`
         
    }
    
    // OrderBrandName used by Order
    type OrderBrandName struct {

        
            BrandName string  `json:"brand_name"`
            ModifiedOn string  `json:"modified_on"`
            Logo string  `json:"logo"`
            Company string  `json:"company"`
            ID float64  `json:"id"`
            CreatedOn string  `json:"created_on"`
         
    }
    
    // CurrentStatus used by Order
    type CurrentStatus struct {

        
            BagStateMapper BagStateMapper  `json:"bag_state_mapper"`
            CreatedAt string  `json:"created_at"`
            StateID float64  `json:"state_id"`
            StoreID float64  `json:"store_id"`
            Status string  `json:"status"`
            ShipmentID string  `json:"shipment_id"`
            DeliveryAwbNumber string  `json:"delivery_awb_number"`
            CurrentStatusID float64  `json:"current_status_id"`
            DeliveryPartnerID float64  `json:"delivery_partner_id"`
            UpdatedAt string  `json:"updated_at"`
            KafkaSync bool  `json:"kafka_sync"`
            BagID float64  `json:"bag_id"`
            StateType string  `json:"state_type"`
         
    }
    
    // OrderBagArticle used by Order
    type OrderBagArticle struct {

        
            Identifiers map[string]interface{}  `json:"identifiers"`
            UID string  `json:"uid"`
            ReturnConfig map[string]interface{}  `json:"return_config"`
         
    }
    
    // Identifier used by Order
    type Identifier struct {

        
            Ean string  `json:"ean"`
         
    }
    
    // FinancialBreakup used by Order
    type FinancialBreakup struct {

        
            TransferPrice float64  `json:"transfer_price"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            PromotionEffectiveDiscount float64  `json:"promotion_effective_discount"`
            DeliveryCharge float64  `json:"delivery_charge"`
            CouponValue float64  `json:"coupon_value"`
            Identifiers Identifier  `json:"identifiers"`
            Cashback float64  `json:"cashback"`
            ItemName string  `json:"item_name"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            HsnCode string  `json:"hsn_code"`
            TotalUnits float64  `json:"total_units"`
            FyndCredits float64  `json:"fynd_credits"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            CodCharges float64  `json:"cod_charges"`
            PriceEffective float64  `json:"price_effective"`
            PriceMarked float64  `json:"price_marked"`
            RefundCredit float64  `json:"refund_credit"`
            Discount float64  `json:"discount"`
            AmountPaidRoundoff float64  `json:"amount_paid_roundoff"`
            GstFee float64  `json:"gst_fee"`
            AmountPaid float64  `json:"amount_paid"`
            Size string  `json:"size"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            AddedToFyndCash bool  `json:"added_to_fynd_cash"`
            CashbackApplied float64  `json:"cashback_applied"`
            GstTag string  `json:"gst_tag"`
            ValueOfGood float64  `json:"value_of_good"`
         
    }
    
    // PlatformDeliveryAddress used by Order
    type PlatformDeliveryAddress struct {

        
            Country string  `json:"country"`
            Latitude float64  `json:"latitude"`
            CreatedAt string  `json:"created_at"`
            Email string  `json:"email"`
            AddressType string  `json:"address_type"`
            Landmark string  `json:"landmark"`
            Address2 string  `json:"address2"`
            AddressCategory string  `json:"address_category"`
            Pincode string  `json:"pincode"`
            ContactPerson string  `json:"contact_person"`
            Longitude float64  `json:"longitude"`
            Area string  `json:"area"`
            UpdatedAt string  `json:"updated_at"`
            Address1 string  `json:"address1"`
            State string  `json:"state"`
            City string  `json:"city"`
            Phone string  `json:"phone"`
            Version string  `json:"version"`
         
    }
    
    // BagConfigs used by Order
    type BagConfigs struct {

        
            EnableTracking bool  `json:"enable_tracking"`
            CanBeCancelled bool  `json:"can_be_cancelled"`
            IsActive bool  `json:"is_active"`
            IsCustomerReturnAllowed bool  `json:"is_customer_return_allowed"`
            IsReturnable bool  `json:"is_returnable"`
            AllowForceReturn bool  `json:"allow_force_return"`
         
    }
    
    // OrderBags used by Order
    type OrderBags struct {

        
            Item PlatformItem  `json:"item"`
            CanReturn bool  `json:"can_return"`
            GstDetails BagGST  `json:"gst_details"`
            DisplayName string  `json:"display_name"`
            AppliedPromos []AppliedPromos  `json:"applied_promos"`
            Identifier string  `json:"identifier"`
            SellerIdentifier string  `json:"seller_identifier"`
            Brand OrderBrandName  `json:"brand"`
            CurrentStatus CurrentStatus  `json:"current_status"`
            CanCancel bool  `json:"can_cancel"`
            BagID float64  `json:"bag_id"`
            EntityType string  `json:"entity_type"`
            Prices Prices  `json:"prices"`
            Article OrderBagArticle  `json:"article"`
            FinancialBreakup FinancialBreakup  `json:"financial_breakup"`
            ParentPromoBags map[string]interface{}  `json:"parent_promo_bags"`
            DeliveryAddress PlatformDeliveryAddress  `json:"delivery_address"`
            Quantity float64  `json:"quantity"`
            BagConfigs BagConfigs  `json:"bag_configs"`
            LineNumber float64  `json:"line_number"`
         
    }
    
    // InvoiceInfo used by Order
    type InvoiceInfo struct {

        
            InvoiceURL string  `json:"invoice_url"`
            UpdatedDate string  `json:"updated_date"`
            CreditNoteID string  `json:"credit_note_id"`
            LabelURL string  `json:"label_url"`
            StoreInvoiceID string  `json:"store_invoice_id"`
         
    }
    
    // OrderDetailsData used by Order
    type OrderDetailsData struct {

        
            OrderingChannelLogo map[string]interface{}  `json:"ordering_channel_logo"`
            OrderValue string  `json:"order_value"`
            OrderingChannel string  `json:"ordering_channel"`
            OrderDate string  `json:"order_date"`
            CodCharges string  `json:"cod_charges"`
            FyndOrderID string  `json:"fynd_order_id"`
            AffiliateID string  `json:"affiliate_id"`
            Source string  `json:"source"`
            TaxDetails map[string]interface{}  `json:"tax_details"`
         
    }
    
    // BagStatusHistory used by Order
    type BagStatusHistory struct {

        
            BagStateMapper BagStateMapper  `json:"bag_state_mapper"`
            CreatedAt string  `json:"created_at"`
            StateID float64  `json:"state_id"`
            AppDisplayName string  `json:"app_display_name"`
            StoreID float64  `json:"store_id"`
            ShipmentID string  `json:"shipment_id"`
            Status string  `json:"status"`
            Forward bool  `json:"forward"`
            DeliveryAwbNumber string  `json:"delivery_awb_number"`
            DeliveryPartnerID float64  `json:"delivery_partner_id"`
            UpdatedAt string  `json:"updated_at"`
            KafkaSync bool  `json:"kafka_sync"`
            BagID float64  `json:"bag_id"`
            DisplayName string  `json:"display_name"`
            StateType string  `json:"state_type"`
            Reasons []map[string]interface{}  `json:"reasons"`
            BshID float64  `json:"bsh_id"`
         
    }
    
    // CompanyDetails used by Order
    type CompanyDetails struct {

        
            CompanyName string  `json:"company_name"`
            CompanyID float64  `json:"company_id"`
            CompanyCin string  `json:"company_cin"`
            ManufacturerAddress map[string]interface{}  `json:"manufacturer_address"`
         
    }
    
    // PlatformShipment used by Order
    type PlatformShipment struct {

        
            TotalBags float64  `json:"total_bags"`
            ShipmentImages []string  `json:"shipment_images"`
            Meta Meta  `json:"meta"`
            ShipmentID string  `json:"shipment_id"`
            DpDetails DPDetailsData  `json:"dp_details"`
            PackagingType string  `json:"packaging_type"`
            Vertical string  `json:"vertical"`
            DeliveryDetails UserDetailsData  `json:"delivery_details"`
            GstDetails GSTDetailsData  `json:"gst_details"`
            Payments ShipmentPayments  `json:"payments"`
            BillingDetails UserDetailsData  `json:"billing_details"`
            TrackingList []TrackingList  `json:"tracking_list"`
            ShipmentQuantity float64  `json:"shipment_quantity"`
            FulfillingStore FulfillingStore  `json:"fulfilling_store"`
            PriorityText string  `json:"priority_text"`
            InvoiceID string  `json:"invoice_id"`
            FulfilmentPriority float64  `json:"fulfilment_priority"`
            CustomMeta []map[string]interface{}  `json:"custom_meta"`
            PaymentMethods map[string]interface{}  `json:"payment_methods"`
            DeliverySlot map[string]interface{}  `json:"delivery_slot"`
            OperationalStatus string  `json:"operational_status"`
            Coupon map[string]interface{}  `json:"coupon"`
            LockStatus bool  `json:"lock_status"`
            TotalItems float64  `json:"total_items"`
            Prices Prices  `json:"prices"`
            Status ShipmentStatusData  `json:"status"`
            JourneyType string  `json:"journey_type"`
            PickedDate string  `json:"picked_date"`
            User UserDataInfo  `json:"user"`
            EnableDpTracking bool  `json:"enable_dp_tracking"`
            OrderingStore OrderingStoreDetails  `json:"ordering_store"`
            PaymentMode string  `json:"payment_mode"`
            Bags []OrderBags  `json:"bags"`
            Invoice InvoiceInfo  `json:"invoice"`
            ShipmentStatus string  `json:"shipment_status"`
            ForwardShipmentID string  `json:"forward_shipment_id"`
            PlatformLogo string  `json:"platform_logo"`
            Order OrderDetailsData  `json:"order"`
            AffiliateDetails AffiliateDetails  `json:"affiliate_details"`
            UserAgent string  `json:"user_agent"`
            BagStatusHistory []BagStatusHistory  `json:"bag_status_history"`
            CompanyDetails CompanyDetails  `json:"company_details"`
         
    }
    
    // ShipmentInfoResponse used by Order
    type ShipmentInfoResponse struct {

        
            Shipments []PlatformShipment  `json:"shipments"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // BillingStaffDetails used by Order
    type BillingStaffDetails struct {

        
            LastName string  `json:"last_name"`
            FirstName string  `json:"first_name"`
            User string  `json:"user"`
            StaffID float64  `json:"staff_id"`
         
    }
    
    // PlatformUserDetails used by Order
    type PlatformUserDetails struct {

        
            PlatformUserID string  `json:"platform_user_id"`
            PlatformUserLastName string  `json:"platform_user_last_name"`
            PlatformUserFirstName string  `json:"platform_user_first_name"`
         
    }
    
    // TransactionData used by Order
    type TransactionData struct {

        
            TransactionID string  `json:"transaction_id"`
            Entity string  `json:"entity"`
            AmountPaid string  `json:"amount_paid"`
            UniqueReferenceNumber string  `json:"unique_reference_number"`
            Currency string  `json:"currency"`
            Status string  `json:"status"`
            TerminalID string  `json:"terminal_id"`
            PaymentID string  `json:"payment_id"`
         
    }
    
    // OrderMeta used by Order
    type OrderMeta struct {

        
            PaymentType string  `json:"payment_type"`
            OrderChildEntities []string  `json:"order_child_entities"`
            OrderPlatform string  `json:"order_platform"`
            OrderType string  `json:"order_type"`
            OrderTags []map[string]interface{}  `json:"order_tags"`
            EmployeeID float64  `json:"employee_id"`
            BillingStaffDetails BillingStaffDetails  `json:"billing_staff_details"`
            CompanyLogo string  `json:"company_logo"`
            Comment string  `json:"comment"`
            PlatformUserDetails PlatformUserDetails  `json:"platform_user_details"`
            CustomerNote string  `json:"customer_note"`
            TransactionData TransactionData  `json:"transaction_data"`
            OrderingStore float64  `json:"ordering_store"`
            Staff map[string]interface{}  `json:"staff"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            Files []map[string]interface{}  `json:"files"`
            MongoCartID float64  `json:"mongo_cart_id"`
            CartID float64  `json:"cart_id"`
            CurrencySymbol string  `json:"currency_symbol"`
         
    }
    
    // TaxDetails used by Order
    type TaxDetails struct {

        
            PanNo string  `json:"pan_no"`
            Gstin string  `json:"gstin"`
         
    }
    
    // OrderDict used by Order
    type OrderDict struct {

        
            Meta OrderMeta  `json:"meta"`
            OrderDate string  `json:"order_date"`
            Prices Prices  `json:"prices"`
            PaymentMethods map[string]interface{}  `json:"payment_methods"`
            FyndOrderID string  `json:"fynd_order_id"`
            TaxDetails TaxDetails  `json:"tax_details"`
         
    }
    
    // ShipmentDetailsResponse used by Order
    type ShipmentDetailsResponse struct {

        
            Shipments []PlatformShipment  `json:"shipments"`
            Order OrderDict  `json:"order"`
            Success bool  `json:"success"`
         
    }
    
    // SubLane used by Order
    type SubLane struct {

        
            Value string  `json:"value"`
            TotalItems float64  `json:"total_items"`
            Index float64  `json:"index"`
            Actions []map[string]interface{}  `json:"actions"`
            Text string  `json:"text"`
         
    }
    
    // SuperLane used by Order
    type SuperLane struct {

        
            Text string  `json:"text"`
            Value string  `json:"value"`
            Options []SubLane  `json:"options"`
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

        
            Value string  `json:"value"`
            Display string  `json:"display"`
            Name string  `json:"name"`
         
    }
    
    // PlatformOrderItems used by Order
    type PlatformOrderItems struct {

        
            Shipments []PlatformShipment  `json:"shipments"`
            TotalOrderValue float64  `json:"total_order_value"`
            Meta map[string]interface{}  `json:"meta"`
            OrderValue float64  `json:"order_value"`
            UserInfo UserDataInfo  `json:"user_info"`
            Channel PlatformChannel  `json:"channel"`
            OrderCreatedTime string  `json:"order_created_time"`
            OrderID string  `json:"order_id"`
            PaymentMode string  `json:"payment_mode"`
            BreakupValues []PlatformBreakupValues  `json:"breakup_values"`
         
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
            TotalCount float64  `json:"total_count"`
            Message string  `json:"message"`
            Page Page  `json:"page"`
            Success bool  `json:"success"`
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
            Value float64  `json:"value"`
            Key string  `json:"key"`
            Options []Options  `json:"options"`
         
    }
    
    // MetricCountResponse used by Order
    type MetricCountResponse struct {

        
            Items []MetricsCount  `json:"items"`
         
    }
    
    // PlatformTrack used by Order
    type PlatformTrack struct {

        
            ShipmentType string  `json:"shipment_type"`
            UpdatedTime string  `json:"updated_time"`
            LastLocationRecievedAt string  `json:"last_location_recieved_at"`
            Reason string  `json:"reason"`
            Meta map[string]interface{}  `json:"meta"`
            AccountName string  `json:"account_name"`
            Status string  `json:"status"`
            Awb string  `json:"awb"`
            RawStatus string  `json:"raw_status"`
            UpdatedAt string  `json:"updated_at"`
         
    }
    
    // PlatformShipmentTrack used by Order
    type PlatformShipmentTrack struct {

        
            Results []PlatformTrack  `json:"results"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // AdvanceFilterInfo used by Order
    type AdvanceFilterInfo struct {

        
            Returned []FiltersInfo  `json:"returned"`
            ActionCentre []FiltersInfo  `json:"action_centre"`
            Filters []FiltersInfo  `json:"filters"`
            Unfulfilled []FiltersInfo  `json:"unfulfilled"`
            Processed []FiltersInfo  `json:"processed"`
         
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

        
            RequestDetails map[string]interface{}  `json:"request_details"`
            ReportType string  `json:"report_type"`
            Status string  `json:"status"`
            ReportRequestedAt string  `json:"report_requested_at"`
            ReportID string  `json:"report_id"`
            ReportName string  `json:"report_name"`
            DisplayName string  `json:"display_name"`
            ReportCreatedAt string  `json:"report_created_at"`
            S3Key string  `json:"s3_key"`
         
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

        
            Type string  `json:"type"`
            Value string  `json:"value"`
            Message string  `json:"message"`
         
    }
    
    // JioCodeUpsertResponse used by Order
    type JioCodeUpsertResponse struct {

        
            Data []map[string]interface{}  `json:"data"`
            Identifier string  `json:"identifier"`
            Success bool  `json:"success"`
            TraceID string  `json:"trace_id"`
            Error []NestedErrorSchemaDataSet  `json:"error"`
         
    }
    
    // BulkInvoicingResponse used by Order
    type BulkInvoicingResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // BulkInvoiceLabelResponse used by Order
    type BulkInvoiceLabelResponse struct {

        
            InvoiceStatus string  `json:"invoice_status"`
            BatchID string  `json:"batch_id"`
            Data map[string]interface{}  `json:"data"`
            StoreID string  `json:"store_id"`
            StoreCode string  `json:"store_code"`
            StoreName string  `json:"store_name"`
            Label map[string]interface{}  `json:"label"`
            DoInvoiceLabelGenerated bool  `json:"do_invoice_label_generated"`
            CompanyID string  `json:"company_id"`
            Invoice map[string]interface{}  `json:"invoice"`
         
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

        
            Method string  `json:"method"`
            Upload FileUploadResponse  `json:"upload"`
            ContentType string  `json:"content_type"`
            Size float64  `json:"size"`
            FilePath string  `json:"file_path"`
            Tags []string  `json:"tags"`
            FileName string  `json:"file_name"`
            Namespace string  `json:"namespace"`
            Cdn URL  `json:"cdn"`
            Operation string  `json:"operation"`
         
    }
    
    // BulkListingPage used by Order
    type BulkListingPage struct {

        
            Type string  `json:"type"`
            HasNext bool  `json:"has_next"`
            Total float64  `json:"total"`
            Size float64  `json:"size"`
            HasPrevious bool  `json:"has_previous"`
            Current float64  `json:"current"`
         
    }
    
    // bulkListingData used by Order
    type bulkListingData struct {

        
            Total float64  `json:"total"`
            FileName string  `json:"file_name"`
            StoreName string  `json:"store_name"`
            ID string  `json:"id"`
            SuccessfulShipments []map[string]interface{}  `json:"successful_shipments"`
            Processing float64  `json:"processing"`
            UploadedOn string  `json:"uploaded_on"`
            FailedShipments []map[string]interface{}  `json:"failed_shipments"`
            CompanyID float64  `json:"company_id"`
            BatchID string  `json:"batch_id"`
            ExcelURL string  `json:"excel_url"`
            StoreID float64  `json:"store_id"`
            Status string  `json:"status"`
            Successful float64  `json:"successful"`
            ProcessingShipments []string  `json:"processing_shipments"`
            Failed float64  `json:"failed"`
            UserID string  `json:"user_id"`
            UserName string  `json:"user_name"`
            StoreCode string  `json:"store_code"`
         
    }
    
    // BulkListingResponse used by Order
    type BulkListingResponse struct {

        
            Page BulkListingPage  `json:"page"`
            Success bool  `json:"success"`
            Data []bulkListingData  `json:"data"`
            Error string  `json:"error"`
         
    }
    
    // QuestionSet used by Order
    type QuestionSet struct {

        
            DisplayName string  `json:"display_name"`
            ID float64  `json:"id"`
         
    }
    
    // Reason used by Order
    type Reason struct {

        
            QcType []string  `json:"qc_type"`
            DisplayName string  `json:"display_name"`
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

        
            Status bool  `json:"status"`
            Message string  `json:"message"`
         
    }
    
    // BulkActionDetailsDataField used by Order
    type BulkActionDetailsDataField struct {

        
            TotalShipmentsCount float64  `json:"total_shipments_count"`
            BatchID string  `json:"batch_id"`
            SuccessfulShipmentIds []string  `json:"successful_shipment_ids"`
            SuccessfulShipmentsCount float64  `json:"successful_shipments_count"`
            CompanyID string  `json:"company_id"`
            FailedShipmentsCount float64  `json:"failed_shipments_count"`
            ProcessingShipmentsCount float64  `json:"processing_shipments_count"`
         
    }
    
    // BulkActionDetailsResponse used by Order
    type BulkActionDetailsResponse struct {

        
            FailedRecords []string  `json:"failed_records"`
            Data []BulkActionDetailsDataField  `json:"data"`
            Message string  `json:"message"`
            UploadedBy string  `json:"uploaded_by"`
            Status bool  `json:"status"`
            UploadedOn string  `json:"uploaded_on"`
            Success string  `json:"success"`
            UserID string  `json:"user_id"`
            Error []string  `json:"error"`
         
    }
    
    // B2BPODetails used by Order
    type B2BPODetails struct {

        
            PoLineAmount float64  `json:"po_line_amount"`
            DockerNumber string  `json:"docker_number"`
            PartialCanRet bool  `json:"partial_can_ret"`
            TotalGstPercentage float64  `json:"total_gst_percentage"`
            PoTaxAmount float64  `json:"po_tax_amount"`
            ItemBasePrice float64  `json:"item_base_price"`
         
    }
    
    // BagMeta used by Order
    type BagMeta struct {

        
            B2bPoDetails B2BPODetails  `json:"b2b_po_details"`
         
    }
    
    // Attributes used by Order
    type Attributes struct {

        
            BrandName string  `json:"brand_name"`
            PrimaryColorHex string  `json:"primary_color_hex"`
            PrimaryColor string  `json:"primary_color"`
            Essential string  `json:"essential"`
            MarketerName string  `json:"marketer_name"`
            PrimaryMaterial string  `json:"primary_material"`
            Name string  `json:"name"`
            Gender []string  `json:"gender"`
            MarketerAddress string  `json:"marketer_address"`
         
    }
    
    // Item used by Order
    type Item struct {

        
            Meta map[string]interface{}  `json:"meta"`
            CanReturn bool  `json:"can_return"`
            DepartmentID float64  `json:"department_id"`
            Code string  `json:"code"`
            WebstoreProductURL string  `json:"webstore_product_url"`
            LastUpdatedAt string  `json:"last_updated_at"`
            Image []string  `json:"image"`
            ItemID float64  `json:"item_id"`
            L3Category float64  `json:"l3_category"`
            Brand string  `json:"brand"`
            CanCancel bool  `json:"can_cancel"`
            SlugKey string  `json:"slug_key"`
            BrandID float64  `json:"brand_id"`
            L1CategoryID float64  `json:"l1_category_id"`
            L3CategoryName string  `json:"l3_category_name"`
            Gender string  `json:"gender"`
            L2CategoryID float64  `json:"l2_category_id"`
            Color string  `json:"color"`
            Attributes Attributes  `json:"attributes"`
            L2Category []string  `json:"l2_category"`
            Size string  `json:"size"`
            Name string  `json:"name"`
            L1Category []string  `json:"l1_category"`
            BranchURL string  `json:"branch_url"`
         
    }
    
    // BagGSTDetails used by Order
    type BagGSTDetails struct {

        
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            GstFee float64  `json:"gst_fee"`
            IsDefaultHsnCode bool  `json:"is_default_hsn_code"`
            SgstTaxPercentage float64  `json:"sgst_tax_percentage"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            SgstGstFee string  `json:"sgst_gst_fee"`
            CgstTaxPercentage float64  `json:"cgst_tax_percentage"`
            GstTag string  `json:"gst_tag"`
            ValueOfGood float64  `json:"value_of_good"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            IgstTaxPercentage float64  `json:"igst_tax_percentage"`
            GstinCode string  `json:"gstin_code"`
            HsnCode string  `json:"hsn_code"`
            CgstGstFee string  `json:"cgst_gst_fee"`
            IgstGstFee string  `json:"igst_gst_fee"`
            HsnCodeID string  `json:"hsn_code_id"`
         
    }
    
    // Dates used by Order
    type Dates struct {

        
            DeliveryDate interface{}  `json:"delivery_date"`
            OrderCreated string  `json:"order_created"`
         
    }
    
    // Brand used by Order
    type Brand struct {

        
            InvoicePrefix string  `json:"invoice_prefix"`
            IsVirtualInvoice bool  `json:"is_virtual_invoice"`
            BrandName string  `json:"brand_name"`
            ModifiedOn float64  `json:"modified_on"`
            PickupLocation string  `json:"pickup_location"`
            BrandID float64  `json:"brand_id"`
            Logo string  `json:"logo"`
            Company string  `json:"company"`
            StartDate string  `json:"start_date"`
            CreditNoteExpiryDays float64  `json:"credit_note_expiry_days"`
            CreditNoteAllowed bool  `json:"credit_note_allowed"`
            CreatedOn float64  `json:"created_on"`
            ScriptLastRan string  `json:"script_last_ran"`
         
    }
    
    // BagReturnableCancelableStatus used by Order
    type BagReturnableCancelableStatus struct {

        
            EnableTracking bool  `json:"enable_tracking"`
            CanBeCancelled bool  `json:"can_be_cancelled"`
            IsActive bool  `json:"is_active"`
            IsCustomerReturnAllowed bool  `json:"is_customer_return_allowed"`
            IsReturnable bool  `json:"is_returnable"`
         
    }
    
    // ReturnConfig used by Order
    type ReturnConfig struct {

        
            Returnable bool  `json:"returnable"`
            Time float64  `json:"time"`
            Unit string  `json:"unit"`
         
    }
    
    // Weight used by Order
    type Weight struct {

        
            Shipping float64  `json:"shipping"`
            IsDefault bool  `json:"is_default"`
            Unit string  `json:"unit"`
         
    }
    
    // Article used by Order
    type Article struct {

        
            IsSet bool  `json:"is_set"`
            ReturnConfig ReturnConfig  `json:"return_config"`
            Size string  `json:"size"`
            ChildDetails map[string]interface{}  `json:"child_details"`
            Identifiers Identifier  `json:"identifiers"`
            EspModified interface{}  `json:"esp_modified"`
            SellerIdentifier string  `json:"seller_identifier"`
            Code string  `json:"code"`
            UID string  `json:"uid"`
            Dimensions Dimensions  `json:"dimensions"`
            RawMeta interface{}  `json:"raw_meta"`
            ASet map[string]interface{}  `json:"a_set"`
            ID string  `json:"_id"`
            Weight Weight  `json:"weight"`
         
    }
    
    // StoreEinvoice used by Order
    type StoreEinvoice struct {

        
            Password string  `json:"password"`
            User string  `json:"user"`
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
    
    // Document used by Order
    type Document struct {

        
            LegalName string  `json:"legal_name"`
            Value string  `json:"value"`
            URL string  `json:"url"`
            Verified bool  `json:"verified"`
            DsType string  `json:"ds_type"`
         
    }
    
    // StoreDocuments used by Order
    type StoreDocuments struct {

        
            Gst Document  `json:"gst"`
         
    }
    
    // StoreMeta used by Order
    type StoreMeta struct {

        
            Stage string  `json:"stage"`
            GstNumber string  `json:"gst_number"`
            GstCredentials StoreGstCredentials  `json:"gst_credentials"`
            NotificationEmails []string  `json:"notification_emails"`
            Timing []map[string]interface{}  `json:"timing"`
            EinvoicePortalDetails EInvoicePortalDetails  `json:"einvoice_portal_details"`
            Documents StoreDocuments  `json:"documents"`
            AdditionalContactDetails map[string]interface{}  `json:"additional_contact_details"`
            EwaybillPortalDetails map[string]interface{}  `json:"ewaybill_portal_details"`
            ProductReturnConfig map[string]interface{}  `json:"product_return_config"`
            DisplayName string  `json:"display_name"`
         
    }
    
    // StoreAddress used by Order
    type StoreAddress struct {

        
            Email string  `json:"email"`
            UpdatedAt string  `json:"updated_at"`
            City string  `json:"city"`
            Latitude float64  `json:"latitude"`
            AddressType string  `json:"address_type"`
            CountryCode string  `json:"country_code"`
            Address2 string  `json:"address2"`
            Address1 string  `json:"address1"`
            Longitude float64  `json:"longitude"`
            Country string  `json:"country"`
            AddressCategory string  `json:"address_category"`
            Pincode float64  `json:"pincode"`
            ContactPerson string  `json:"contact_person"`
            Phone string  `json:"phone"`
            CreatedAt string  `json:"created_at"`
            Landmark string  `json:"landmark"`
            Area string  `json:"area"`
            State string  `json:"state"`
            Version string  `json:"version"`
         
    }
    
    // Store used by Order
    type Store struct {

        
            Meta StoreMeta  `json:"meta"`
            IsEnabledForRecon bool  `json:"is_enabled_for_recon"`
            SID string  `json:"s_id"`
            Code string  `json:"code"`
            UpdatedAt string  `json:"updated_at"`
            City string  `json:"city"`
            StoreAddressJson StoreAddress  `json:"store_address_json"`
            Latitude float64  `json:"latitude"`
            Address2 string  `json:"address2"`
            AlohomoraUserID float64  `json:"alohomora_user_id"`
            BrandStoreTags []string  `json:"brand_store_tags"`
            IsActive bool  `json:"is_active"`
            Address1 string  `json:"address1"`
            MallArea string  `json:"mall_area"`
            StoreEmail string  `json:"store_email"`
            CompanyID float64  `json:"company_id"`
            Longitude float64  `json:"longitude"`
            Country string  `json:"country"`
            LoginUsername string  `json:"login_username"`
            BrandID interface{}  `json:"brand_id"`
            FulfillmentChannel string  `json:"fulfillment_channel"`
            VatNo string  `json:"vat_no"`
            Pincode string  `json:"pincode"`
            IsArchived bool  `json:"is_archived"`
            ContactPerson string  `json:"contact_person"`
            PackagingMaterialCount float64  `json:"packaging_material_count"`
            ParentStoreID float64  `json:"parent_store_id"`
            LocationType string  `json:"location_type"`
            Phone float64  `json:"phone"`
            CreatedAt string  `json:"created_at"`
            MallName string  `json:"mall_name"`
            OrderIntegrationID string  `json:"order_integration_id"`
            Name string  `json:"name"`
            StoreActiveFrom string  `json:"store_active_from"`
            State string  `json:"state"`
         
    }
    
    // AffiliateBagDetails used by Order
    type AffiliateBagDetails struct {

        
            AffiliateBagID string  `json:"affiliate_bag_id"`
            AffiliateMeta AffiliateMeta  `json:"affiliate_meta"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            EmployeeDiscount float64  `json:"employee_discount"`
         
    }
    
    // ArticleDetails used by Order
    type ArticleDetails struct {

        
            Status map[string]interface{}  `json:"status"`
         
    }
    
    // BagDetailsPlatformResponse used by Order
    type BagDetailsPlatformResponse struct {

        
            Meta BagMeta  `json:"meta"`
            Item Item  `json:"item"`
            Tags []string  `json:"tags"`
            ShipmentID string  `json:"shipment_id"`
            RestorePromos map[string]interface{}  `json:"restore_promos"`
            GstDetails BagGSTDetails  `json:"gst_details"`
            Dates Dates  `json:"dates"`
            DisplayName string  `json:"display_name"`
            AppliedPromos []map[string]interface{}  `json:"applied_promos"`
            OriginalBagList []float64  `json:"original_bag_list"`
            BID float64  `json:"b_id"`
            RestoreCoupon bool  `json:"restore_coupon"`
            Identifier string  `json:"identifier"`
            SellerIdentifier string  `json:"seller_identifier"`
            Brand Brand  `json:"brand"`
            CurrentStatus BagStatusHistory  `json:"current_status"`
            CurrentOperationalStatus BagStatusHistory  `json:"current_operational_status"`
            OperationalStatus string  `json:"operational_status"`
            Reasons []map[string]interface{}  `json:"reasons"`
            EntityType string  `json:"entity_type"`
            BagStatus []BagStatusHistory  `json:"bag_status"`
            BagUpdateTime float64  `json:"bag_update_time"`
            Prices Prices  `json:"prices"`
            BType string  `json:"b_type"`
            Status BagReturnableCancelableStatus  `json:"status"`
            JourneyType string  `json:"journey_type"`
            Article Article  `json:"article"`
            OrderingStore Store  `json:"ordering_store"`
            QcRequired interface{}  `json:"qc_required"`
            FinancialBreakup []FinancialBreakup  `json:"financial_breakup"`
            ParentPromoBags map[string]interface{}  `json:"parent_promo_bags"`
            OrderIntegrationID string  `json:"order_integration_id"`
            NoOfBagsOrder float64  `json:"no_of_bags_order"`
            AffiliateBagDetails AffiliateBagDetails  `json:"affiliate_bag_details"`
            AffiliateDetails AffiliateDetails  `json:"affiliate_details"`
            Quantity float64  `json:"quantity"`
            BagStatusHistory BagStatusHistory  `json:"bag_status_history"`
            ArticleDetails ArticleDetails  `json:"article_details"`
            LineNumber float64  `json:"line_number"`
         
    }
    
    // ErrorResponse used by Order
    type ErrorResponse struct {

        
            Message string  `json:"message"`
            Error string  `json:"error"`
         
    }
    
    // Page1 used by Order
    type Page1 struct {

        
            HasNext bool  `json:"has_next"`
            Size float64  `json:"size"`
            PageType string  `json:"page_type"`
            ItemTotal float64  `json:"item_total"`
            Current float64  `json:"current"`
         
    }
    
    // GetBagsPlatformResponse used by Order
    type GetBagsPlatformResponse struct {

        
            Page Page1  `json:"page"`
            Items []BagDetailsPlatformResponse  `json:"items"`
         
    }
    
    // GeneratePosOrderReceiptResponse used by Order
    type GeneratePosOrderReceiptResponse struct {

        
            Success bool  `json:"success"`
            InvoiceReceipt string  `json:"invoice_receipt"`
            OrderID string  `json:"order_id"`
            PaymentReceipt string  `json:"payment_receipt"`
         
    }
    
    // InvalidateShipmentCachePayload used by Order
    type InvalidateShipmentCachePayload struct {

        
            ShipmentIds []string  `json:"shipment_ids"`
         
    }
    
    // InvalidateShipmentCacheNestedResponse used by Order
    type InvalidateShipmentCacheNestedResponse struct {

        
            Error string  `json:"error"`
            ShipmentID string  `json:"shipment_id"`
            Message string  `json:"message"`
            Status float64  `json:"status"`
         
    }
    
    // InvalidateShipmentCacheResponse used by Order
    type InvalidateShipmentCacheResponse struct {

        
            Response []InvalidateShipmentCacheNestedResponse  `json:"response"`
         
    }
    
    // ErrorResponse1 used by Order
    type ErrorResponse1 struct {

        
            ErrorTrace string  `json:"error_trace"`
            Message string  `json:"message"`
            Status float64  `json:"status"`
         
    }
    
    // StoreReassign used by Order
    type StoreReassign struct {

        
            AffiliateBagID string  `json:"affiliate_bag_id"`
            SetID string  `json:"set_id"`
            BagID float64  `json:"bag_id"`
            AffiliateID string  `json:"affiliate_id"`
            StoreID float64  `json:"store_id"`
            MongoArticleID string  `json:"mongo_article_id"`
            ReasonIds []float64  `json:"reason_ids"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            FyndOrderID string  `json:"fynd_order_id"`
            ItemID string  `json:"item_id"`
         
    }
    
    // StoreReassignResponse used by Order
    type StoreReassignResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // Entities used by Order
    type Entities struct {

        
            AffiliateBagID string  `json:"affiliate_bag_id"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            AffiliateID string  `json:"affiliate_id"`
            ID string  `json:"id"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            ReasonText string  `json:"reason_text"`
         
    }
    
    // UpdateShipmentLockPayload used by Order
    type UpdateShipmentLockPayload struct {

        
            EntityType string  `json:"entity_type"`
            Entities []Entities  `json:"entities"`
            Action string  `json:"action"`
            ActionType string  `json:"action_type"`
         
    }
    
    // Bags used by Order
    type Bags struct {

        
            IsLocked bool  `json:"is_locked"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            BagID float64  `json:"bag_id"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
         
    }
    
    // OriginalFilter used by Order
    type OriginalFilter struct {

        
            AffiliateID string  `json:"affiliate_id"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
         
    }
    
    // CheckResponse used by Order
    type CheckResponse struct {

        
            Bags []Bags  `json:"bags"`
            OriginalFilter OriginalFilter  `json:"original_filter"`
            Status string  `json:"status"`
            IsBagLocked bool  `json:"is_bag_locked"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            AffiliateID string  `json:"affiliate_id"`
            ShipmentID string  `json:"shipment_id"`
            IsShipmentLocked bool  `json:"is_shipment_locked"`
            LockStatus bool  `json:"lock_status"`
         
    }
    
    // UpdateShipmentLockResponse used by Order
    type UpdateShipmentLockResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
            CheckResponse []CheckResponse  `json:"check_response"`
         
    }
    
    // AnnouncementResponse used by Order
    type AnnouncementResponse struct {

        
            Title string  `json:"title"`
            CreatedAt string  `json:"created_at"`
            Description string  `json:"description"`
            PlatformID string  `json:"platform_id"`
            PlatformName string  `json:"platform_name"`
            LogoURL string  `json:"logo_url"`
            CompanyID float64  `json:"company_id"`
            ID float64  `json:"id"`
            ToDatetime string  `json:"to_datetime"`
            FromDatetime string  `json:"from_datetime"`
         
    }
    
    // AnnouncementsResponse used by Order
    type AnnouncementsResponse struct {

        
            Announcements []AnnouncementResponse  `json:"announcements"`
         
    }
    
    // BaseResponse used by Order
    type BaseResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // Click2CallResponse used by Order
    type Click2CallResponse struct {

        
            CallID string  `json:"call_id"`
            Status bool  `json:"status"`
         
    }
    
    // EntitiesDataUpdates used by Order
    type EntitiesDataUpdates struct {

        
            Data map[string]interface{}  `json:"data"`
            Filters []map[string]interface{}  `json:"filters"`
         
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
            Products []ProductsDataUpdates  `json:"products"`
         
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
            Identifier string  `json:"identifier"`
            Quantity float64  `json:"quantity"`
         
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
    
    // Products used by Order
    type Products struct {

        
            LineNumber float64  `json:"line_number"`
            Identifier string  `json:"identifier"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // ShipmentsRequest used by Order
    type ShipmentsRequest struct {

        
            DataUpdates DataUpdates  `json:"data_updates"`
            Reasons ReasonsData  `json:"reasons"`
            Products []Products  `json:"products"`
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

        
            ForceTransition bool  `json:"force_transition"`
            LockAfterTransition bool  `json:"lock_after_transition"`
            Statuses []StatuesRequest  `json:"statuses"`
            Task bool  `json:"task"`
            UnlockBeforeTransition bool  `json:"unlock_before_transition"`
         
    }
    
    // ShipmentsResponse used by Order
    type ShipmentsResponse struct {

        
            Meta map[string]interface{}  `json:"meta"`
            Status float64  `json:"status"`
            FinalState map[string]interface{}  `json:"final_state"`
            Exception string  `json:"exception"`
            Code string  `json:"code"`
            Message string  `json:"message"`
            StackTrace string  `json:"stack_trace"`
            Identifier string  `json:"identifier"`
         
    }
    
    // StatuesResponse used by Order
    type StatuesResponse struct {

        
            Shipments []ShipmentsResponse  `json:"shipments"`
         
    }
    
    // UpdateShipmentStatusResponseBody used by Order
    type UpdateShipmentStatusResponseBody struct {

        
            Statuses []StatuesResponse  `json:"statuses"`
         
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
    
    // AffiliateInventoryPaymentConfig used by Order
    type AffiliateInventoryPaymentConfig struct {

        
            Source string  `json:"source"`
            ModeOfPayment string  `json:"mode_of_payment"`
         
    }
    
    // AffiliateInventoryConfig used by Order
    type AffiliateInventoryConfig struct {

        
            ArticleAssignment AffiliateInventoryArticleAssignmentConfig  `json:"article_assignment"`
            Logistics AffiliateInventoryLogisticsConfig  `json:"logistics"`
            Inventory AffiliateInventoryStoreConfig  `json:"inventory"`
            Order AffiliateInventoryOrderConfig  `json:"order"`
            Payment AffiliateInventoryPaymentConfig  `json:"payment"`
         
    }
    
    // AffiliateAppConfigMeta used by Order
    type AffiliateAppConfigMeta struct {

        
            Value string  `json:"value"`
            Name string  `json:"name"`
         
    }
    
    // AffiliateAppConfig used by Order
    type AffiliateAppConfig struct {

        
            UpdatedAt string  `json:"updated_at"`
            Meta []AffiliateAppConfigMeta  `json:"meta"`
            Description string  `json:"description"`
            Name string  `json:"name"`
            Owner string  `json:"owner"`
            ID string  `json:"id"`
            Token string  `json:"token"`
            CreatedAt string  `json:"created_at"`
            Secret string  `json:"secret"`
         
    }
    
    // AffiliateConfig used by Order
    type AffiliateConfig struct {

        
            Inventory AffiliateInventoryConfig  `json:"inventory"`
            App AffiliateAppConfig  `json:"app"`
         
    }
    
    // Affiliate used by Order
    type Affiliate struct {

        
            Config AffiliateConfig  `json:"config"`
            ID string  `json:"id"`
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
            StoreLookup string  `json:"store_lookup"`
            ArticleLookup string  `json:"article_lookup"`
            AffiliateStoreIDMapping []AffiliateStoreIdMapping  `json:"affiliate_store_id_mapping"`
            BagEndState string  `json:"bag_end_state"`
            CreateUser bool  `json:"create_user"`
         
    }
    
    // OrderUser used by Order
    type OrderUser struct {

        
            State string  `json:"state"`
            Phone float64  `json:"phone"`
            Email string  `json:"email"`
            Address2 string  `json:"address2"`
            Mobile float64  `json:"mobile"`
            Pincode string  `json:"pincode"`
            Address1 string  `json:"address1"`
            City string  `json:"city"`
            LastName string  `json:"last_name"`
            Country string  `json:"country"`
            FirstName string  `json:"first_name"`
         
    }
    
    // UserData used by Order
    type UserData struct {

        
            ShippingUser OrderUser  `json:"shipping_user"`
            BillingUser OrderUser  `json:"billing_user"`
         
    }
    
    // MarketPlacePdf used by Order
    type MarketPlacePdf struct {

        
            Label string  `json:"label"`
            Invoice string  `json:"invoice"`
         
    }
    
    // AffiliateBag used by Order
    type AffiliateBag struct {

        
            ItemSize string  `json:"item_size"`
            AffiliateMeta map[string]interface{}  `json:"affiliate_meta"`
            PdfLinks MarketPlacePdf  `json:"pdf_links"`
            AmountPaid float64  `json:"amount_paid"`
            ModifiedOn string  `json:"modified_on"`
            PriceEffective float64  `json:"price_effective"`
            SellerIdentifier string  `json:"seller_identifier"`
            Identifier map[string]interface{}  `json:"identifier"`
            DeliveryCharge float64  `json:"delivery_charge"`
            TransferPrice float64  `json:"transfer_price"`
            UnitPrice float64  `json:"unit_price"`
            HsnCodeID string  `json:"hsn_code_id"`
            Discount float64  `json:"discount"`
            PriceMarked float64  `json:"price_marked"`
            StoreID float64  `json:"store_id"`
            FyndStoreID string  `json:"fynd_store_id"`
            AvlQty float64  `json:"avl_qty"`
            ID string  `json:"_id"`
            Sku string  `json:"sku"`
            AffiliateStoreID string  `json:"affiliate_store_id"`
            CompanyID float64  `json:"company_id"`
            Quantity float64  `json:"quantity"`
            ItemID float64  `json:"item_id"`
         
    }
    
    // ArticleDetails1 used by Order
    type ArticleDetails1 struct {

        
            ID string  `json:"_id"`
            Dimension map[string]interface{}  `json:"dimension"`
            Attributes map[string]interface{}  `json:"attributes"`
            Weight map[string]interface{}  `json:"weight"`
            BrandID float64  `json:"brand_id"`
            Quantity float64  `json:"quantity"`
            Category map[string]interface{}  `json:"category"`
         
    }
    
    // LocationDetails used by Order
    type LocationDetails struct {

        
            Articles []ArticleDetails1  `json:"articles"`
            FulfillmentID float64  `json:"fulfillment_id"`
            FulfillmentType string  `json:"fulfillment_type"`
         
    }
    
    // ShipmentDetails used by Order
    type ShipmentDetails struct {

        
            Articles []ArticleDetails1  `json:"articles"`
            Meta map[string]interface{}  `json:"meta"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            DpID float64  `json:"dp_id"`
            Shipments float64  `json:"shipments"`
            BoxType string  `json:"box_type"`
            FulfillmentID float64  `json:"fulfillment_id"`
         
    }
    
    // ShipmentConfig used by Order
    type ShipmentConfig struct {

        
            LocationDetails LocationDetails  `json:"location_details"`
            Action string  `json:"action"`
            ToPincode string  `json:"to_pincode"`
            Shipment []ShipmentDetails  `json:"shipment"`
            PaymentMode string  `json:"payment_mode"`
            Journey string  `json:"journey"`
            Source string  `json:"source"`
            Identifier string  `json:"identifier"`
         
    }
    
    // ShipmentData used by Order
    type ShipmentData struct {

        
            ShipmentData ShipmentConfig  `json:"shipment_data"`
         
    }
    
    // OrderPriority used by Order
    type OrderPriority struct {

        
            FulfilmentPriorityText string  `json:"fulfilment_priority_text"`
            AffiliatePriorityCode string  `json:"affiliate_priority_code"`
            FulfilmentPriority float64  `json:"fulfilment_priority"`
         
    }
    
    // OrderInfo used by Order
    type OrderInfo struct {

        
            User UserData  `json:"user"`
            Bags []AffiliateBag  `json:"bags"`
            BillingAddress OrderUser  `json:"billing_address"`
            Items map[string]interface{}  `json:"items"`
            Coupon string  `json:"coupon"`
            DeliveryCharges float64  `json:"delivery_charges"`
            Shipment ShipmentData  `json:"shipment"`
            PaymentMode string  `json:"payment_mode"`
            CodCharges float64  `json:"cod_charges"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            ShippingAddress OrderUser  `json:"shipping_address"`
            OrderValue float64  `json:"order_value"`
            Payment map[string]interface{}  `json:"payment"`
            OrderPriority OrderPriority  `json:"order_priority"`
            Discount float64  `json:"discount"`
         
    }
    
    // CreateOrderPayload used by Order
    type CreateOrderPayload struct {

        
            OrderConfig OrderConfig  `json:"order_config"`
            OrderInfo OrderInfo  `json:"order_info"`
            AffiliateID string  `json:"affiliate_id"`
         
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

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
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
    
    // HistoryDict used by Order
    type HistoryDict struct {

        
            User string  `json:"user"`
            Type string  `json:"type"`
            Createdat string  `json:"createdat"`
            L3Detail string  `json:"l3_detail"`
            TicketURL string  `json:"ticket_url"`
            L1Detail string  `json:"l1_detail"`
            BagID float64  `json:"bag_id"`
            Message string  `json:"message"`
            L2Detail string  `json:"l2_detail"`
            TicketID string  `json:"ticket_id"`
         
    }
    
    // ShipmentHistoryResponse used by Order
    type ShipmentHistoryResponse struct {

        
            ActivityHistory []HistoryDict  `json:"activity_history"`
         
    }
    
    // ErrorDetail used by Order
    type ErrorDetail struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
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
    
    // SmsDataPayload used by Order
    type SmsDataPayload struct {

        
            CustomerName string  `json:"customer_name"`
            OrderID string  `json:"order_id"`
            AmountPaid float64  `json:"amount_paid"`
            PhoneNumber float64  `json:"phone_number"`
            PaymentMode string  `json:"payment_mode"`
            Message string  `json:"message"`
            BrandName string  `json:"brand_name"`
            ShipmentID float64  `json:"shipment_id"`
            CountryCode string  `json:"country_code"`
         
    }
    
    // SendSmsPayload used by Order
    type SendSmsPayload struct {

        
            Slug string  `json:"slug"`
            Data SmsDataPayload  `json:"data"`
            BagID float64  `json:"bag_id"`
         
    }
    
    // Meta1 used by Order
    type Meta1 struct {

        
            StateManagerUsed string  `json:"state_manager_used"`
            KafkaEmissionStatus float64  `json:"kafka_emission_status"`
         
    }
    
    // ShipmentDetail used by Order
    type ShipmentDetail struct {

        
            Meta Meta1  `json:"meta"`
            Status string  `json:"status"`
            BagList []float64  `json:"bag_list"`
            ID float64  `json:"id"`
            ShipmentID string  `json:"shipment_id"`
            Remarks string  `json:"remarks"`
         
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
            ShipmentIds []string  `json:"shipment_ids"`
            OrderType string  `json:"order_type"`
            DpID float64  `json:"dp_id"`
         
    }
    
    // ManualAssignDPToShipmentResponse used by Order
    type ManualAssignDPToShipmentResponse struct {

        
            Success string  `json:"success"`
            Errors []string  `json:"errors"`
         
    }
    
    // BillingInfo used by Order
    type BillingInfo struct {

        
            Title string  `json:"title"`
            MiddleName string  `json:"middle_name"`
            FloorNo string  `json:"floor_no"`
            LastName string  `json:"last_name"`
            Country string  `json:"country"`
            State string  `json:"state"`
            HouseNo string  `json:"house_no"`
            CustomerCode string  `json:"customer_code"`
            Address2 string  `json:"address2"`
            Pincode string  `json:"pincode"`
            FirstName string  `json:"first_name"`
            ExternalCustomerCode string  `json:"external_customer_code"`
            StateCode string  `json:"state_code"`
            PrimaryMobileNumber string  `json:"primary_mobile_number"`
            City string  `json:"city"`
            AlternateMobileNumber string  `json:"alternate_mobile_number"`
            Gender string  `json:"gender"`
            PrimaryEmail string  `json:"primary_email"`
            AlternateEmail string  `json:"alternate_email"`
            Address1 string  `json:"address1"`
            CountryCode string  `json:"country_code"`
         
    }
    
    // Tax used by Order
    type Tax struct {

        
            Amount map[string]interface{}  `json:"amount"`
            Rate float64  `json:"rate"`
            Breakup []map[string]interface{}  `json:"breakup"`
            Name string  `json:"name"`
         
    }
    
    // Charge used by Order
    type Charge struct {

        
            Amount map[string]interface{}  `json:"amount"`
            Type string  `json:"type"`
            Name string  `json:"name"`
            Tax Tax  `json:"tax"`
            Code string  `json:"code"`
         
    }
    
    // ShippingInfo used by Order
    type ShippingInfo struct {

        
            Title string  `json:"title"`
            MiddleName string  `json:"middle_name"`
            FloorNo string  `json:"floor_no"`
            LastName string  `json:"last_name"`
            Country string  `json:"country"`
            AddressType string  `json:"address_type"`
            State string  `json:"state"`
            HouseNo string  `json:"house_no"`
            CustomerCode string  `json:"customer_code"`
            Address2 string  `json:"address2"`
            Pincode string  `json:"pincode"`
            FirstName string  `json:"first_name"`
            ExternalCustomerCode string  `json:"external_customer_code"`
            StateCode string  `json:"state_code"`
            PrimaryMobileNumber string  `json:"primary_mobile_number"`
            ShippingType string  `json:"shipping_type"`
            City string  `json:"city"`
            AlternateMobileNumber string  `json:"alternate_mobile_number"`
            Landmark string  `json:"landmark"`
            Slot []map[string]interface{}  `json:"slot"`
            GeoLocation map[string]interface{}  `json:"geo_location"`
            Gender string  `json:"gender"`
            PrimaryEmail string  `json:"primary_email"`
            AlternateEmail string  `json:"alternate_email"`
            Address1 string  `json:"address1"`
            CountryCode string  `json:"country_code"`
         
    }
    
    // TaxInfo used by Order
    type TaxInfo struct {

        
            Gstin string  `json:"gstin"`
            B2bGstinNumber string  `json:"b2b_gstin_number"`
         
    }
    
    // PaymentMethod used by Order
    type PaymentMethod struct {

        
            Amount float64  `json:"amount"`
            CollectBy string  `json:"collect_by"`
            Meta map[string]interface{}  `json:"meta"`
            RefundBy string  `json:"refund_by"`
            Mode string  `json:"mode"`
            Name string  `json:"name"`
            TransactionData map[string]interface{}  `json:"transaction_data"`
         
    }
    
    // PaymentInfo used by Order
    type PaymentInfo struct {

        
            PrimaryMode string  `json:"primary_mode"`
            PaymentMethods []PaymentMethod  `json:"payment_methods"`
         
    }
    
    // LineItem used by Order
    type LineItem struct {

        
            Meta map[string]interface{}  `json:"meta"`
            CustomMessasge string  `json:"custom_messasge"`
            Charges []Charge  `json:"charges"`
            Quantity float64  `json:"quantity"`
            ExternalLineID string  `json:"external_line_id"`
            SellerIdentifier string  `json:"seller_identifier"`
         
    }
    
    // ProcessingDates used by Order
    type ProcessingDates struct {

        
            CustomerPickupSlot map[string]interface{}  `json:"customer_pickup_slot"`
            DpPickupSlot map[string]interface{}  `json:"dp_pickup_slot"`
            DispatchByDate string  `json:"dispatch_by_date"`
            DispatchAfterDate string  `json:"dispatch_after_date"`
            PackByDate string  `json:"pack_by_date"`
            ConfirmByDate string  `json:"confirm_by_date"`
         
    }
    
    // Shipment used by Order
    type Shipment struct {

        
            Meta map[string]interface{}  `json:"meta"`
            ExternalShipmentID string  `json:"external_shipment_id"`
            Priority float64  `json:"priority"`
            LineItems []LineItem  `json:"line_items"`
            ProcessingDates ProcessingDates  `json:"processing_dates"`
            LocationID float64  `json:"location_id"`
         
    }
    
    // CreateOrderAPI used by Order
    type CreateOrderAPI struct {

        
            BillingInfo BillingInfo  `json:"billing_info"`
            Meta map[string]interface{}  `json:"meta"`
            Charges []Charge  `json:"charges"`
            ShippingInfo ShippingInfo  `json:"shipping_info"`
            TaxInfo TaxInfo  `json:"tax_info"`
            PaymentInfo PaymentInfo  `json:"payment_info"`
            ExternalOrderID string  `json:"external_order_id"`
            Shipments []Shipment  `json:"shipments"`
            CurrencyInfo map[string]interface{}  `json:"currency_info"`
            ExternalCreationDate string  `json:"external_creation_date"`
         
    }
    
    // CreateOrderErrorReponse used by Order
    type CreateOrderErrorReponse struct {

        
            Meta string  `json:"meta"`
            Status float64  `json:"status"`
            RequestID string  `json:"request_id"`
            Exception string  `json:"exception"`
            Code string  `json:"code"`
            Message string  `json:"message"`
            StackTrace string  `json:"stack_trace"`
            Info interface{}  `json:"info"`
         
    }
    
    // DpConfiguration used by Order
    type DpConfiguration struct {

        
            ShippingBy string  `json:"shipping_by"`
         
    }
    
    // PaymentMethods used by Order
    type PaymentMethods struct {

        
            CollectBy string  `json:"collect_by"`
            Mode string  `json:"mode"`
            RefundBy string  `json:"refund_by"`
         
    }
    
    // CreateChannelPaymentInfo used by Order
    type CreateChannelPaymentInfo struct {

        
            Source string  `json:"source"`
            ModeOfPayment string  `json:"mode_of_payment"`
            PaymentMethods []PaymentMethods  `json:"payment_methods"`
         
    }
    
    // CreateChannelConfig used by Order
    type CreateChannelConfig struct {

        
            ShipmentAssignment string  `json:"shipment_assignment"`
            LocationReassignment bool  `json:"location_reassignment"`
            LockStates []string  `json:"lock_states"`
            DpConfiguration DpConfiguration  `json:"dp_configuration"`
            LogoURL map[string]interface{}  `json:"logo_url"`
            PaymentInfo CreateChannelPaymentInfo  `json:"payment_info"`
         
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

        
            Success bool  `json:"success"`
            Message []string  `json:"message"`
         
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
            Error string  `json:"error"`
            Status float64  `json:"status"`
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

        
            IsActive bool  `json:"is_active"`
            Result SearchKeywordResult  `json:"result"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Words []string  `json:"words"`
            AppID string  `json:"app_id"`
         
    }
    
    // GetSearchWordsData used by Catalog
    type GetSearchWordsData struct {

        
            IsActive bool  `json:"is_active"`
            UID string  `json:"uid"`
            Result map[string]interface{}  `json:"result"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Words []string  `json:"words"`
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
    
    // GetSearchWordsResponse used by Catalog
    type GetSearchWordsResponse struct {

        
            Items []GetSearchWordsData  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // Media used by Catalog
    type Media struct {

        
            URL string  `json:"url"`
            Type string  `json:"type"`
            AspectRatio string  `json:"aspect_ratio"`
         
    }
    
    // AutocompletePageAction used by Catalog
    type AutocompletePageAction struct {

        
            URL string  `json:"url"`
            Query map[string]interface{}  `json:"query"`
            Type string  `json:"type"`
            Params map[string]interface{}  `json:"params"`
         
    }
    
    // AutocompleteAction used by Catalog
    type AutocompleteAction struct {

        
            Type string  `json:"type"`
            Page AutocompletePageAction  `json:"page"`
         
    }
    
    // AutocompleteResult used by Catalog
    type AutocompleteResult struct {

        
            Logo Media  `json:"logo"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Display string  `json:"display"`
            Action AutocompleteAction  `json:"action"`
         
    }
    
    // CreateAutocompleteKeyword used by Catalog
    type CreateAutocompleteKeyword struct {

        
            IsActive bool  `json:"is_active"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Words []string  `json:"words"`
            AppID string  `json:"app_id"`
            Results []AutocompleteResult  `json:"results"`
         
    }
    
    // GetAutocompleteWordsData used by Catalog
    type GetAutocompleteWordsData struct {

        
            UID string  `json:"uid"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Words []string  `json:"words"`
            AppID string  `json:"app_id"`
            Results []map[string]interface{}  `json:"results"`
         
    }
    
    // GetAutocompleteWordsResponse used by Catalog
    type GetAutocompleteWordsResponse struct {

        
            Items []GetAutocompleteWordsData  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // CreateAutocompleteWordsResponse used by Catalog
    type CreateAutocompleteWordsResponse struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            AppID string  `json:"app_id"`
            Results []map[string]interface{}  `json:"results"`
            Words []string  `json:"words"`
         
    }
    
    // ProductBundleItem used by Catalog
    type ProductBundleItem struct {

        
            AllowRemove bool  `json:"allow_remove"`
            AutoAddToCart bool  `json:"auto_add_to_cart"`
            AutoSelect bool  `json:"auto_select"`
            MinQuantity float64  `json:"min_quantity"`
            MaxQuantity float64  `json:"max_quantity"`
            ProductUID float64  `json:"product_uid"`
         
    }
    
    // ProductBundleRequest used by Catalog
    type ProductBundleRequest struct {

        
            Logo string  `json:"logo"`
            Name string  `json:"name"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            IsActive bool  `json:"is_active"`
            Slug string  `json:"slug"`
            ModifiedOn string  `json:"modified_on"`
            CompanyID float64  `json:"company_id"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Choice string  `json:"choice"`
            PageVisibility []string  `json:"page_visibility"`
            Products []ProductBundleItem  `json:"products"`
            CreatedOn string  `json:"created_on"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // GetProductBundleCreateResponse used by Catalog
    type GetProductBundleCreateResponse struct {

        
            Logo string  `json:"logo"`
            Name string  `json:"name"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            IsActive bool  `json:"is_active"`
            Slug string  `json:"slug"`
            ModifiedOn string  `json:"modified_on"`
            CompanyID float64  `json:"company_id"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Choice string  `json:"choice"`
            PageVisibility []string  `json:"page_visibility"`
            Products []ProductBundleItem  `json:"products"`
            CreatedOn string  `json:"created_on"`
            Meta map[string]interface{}  `json:"meta"`
            ID string  `json:"id"`
         
    }
    
    // GetProductBundleListingResponse used by Catalog
    type GetProductBundleListingResponse struct {

        
            Items []GetProductBundleCreateResponse  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // ProductBundleUpdateRequest used by Catalog
    type ProductBundleUpdateRequest struct {

        
            Logo string  `json:"logo"`
            Name string  `json:"name"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            IsActive bool  `json:"is_active"`
            Slug string  `json:"slug"`
            ModifiedOn string  `json:"modified_on"`
            CompanyID float64  `json:"company_id"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Choice string  `json:"choice"`
            PageVisibility []string  `json:"page_visibility"`
            Products []ProductBundleItem  `json:"products"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // Size used by Catalog
    type Size struct {

        
            Value string  `json:"value"`
            IsAvailable bool  `json:"is_available"`
            Quantity float64  `json:"quantity"`
            Display string  `json:"display"`
         
    }
    
    // Price used by Catalog
    type Price struct {

        
            MaxMarked float64  `json:"max_marked"`
            Currency string  `json:"currency"`
            MinEffective float64  `json:"min_effective"`
            MinMarked float64  `json:"min_marked"`
            MaxEffective float64  `json:"max_effective"`
         
    }
    
    // LimitedProductData used by Catalog
    type LimitedProductData struct {

        
            Sizes []string  `json:"sizes"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
            Price map[string]interface{}  `json:"price"`
            ItemCode string  `json:"item_code"`
            CountryOfOrigin string  `json:"country_of_origin"`
            ShortDescription string  `json:"short_description"`
            Images []string  `json:"images"`
            Quantity float64  `json:"quantity"`
            Attributes map[string]interface{}  `json:"attributes"`
            Identifier map[string]interface{}  `json:"identifier"`
         
    }
    
    // GetProducts used by Catalog
    type GetProducts struct {

        
            Sizes []Size  `json:"sizes"`
            AllowRemove bool  `json:"allow_remove"`
            AutoAddToCart bool  `json:"auto_add_to_cart"`
            AutoSelect bool  `json:"auto_select"`
            MinQuantity float64  `json:"min_quantity"`
            Price Price  `json:"price"`
            ProductDetails LimitedProductData  `json:"product_details"`
            MaxQuantity float64  `json:"max_quantity"`
            ProductUID float64  `json:"product_uid"`
         
    }
    
    // GetProductBundleResponse used by Catalog
    type GetProductBundleResponse struct {

        
            Logo string  `json:"logo"`
            Name string  `json:"name"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            IsActive bool  `json:"is_active"`
            Slug string  `json:"slug"`
            CompanyID float64  `json:"company_id"`
            PageVisibility []string  `json:"page_visibility"`
            Choice string  `json:"choice"`
            Products []GetProducts  `json:"products"`
            Meta map[string]interface{}  `json:"meta"`
         
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

        
            BrandID float64  `json:"brand_id"`
            Name string  `json:"name"`
            Active bool  `json:"active"`
            ModifiedOn string  `json:"modified_on"`
            Description string  `json:"description"`
            CompanyID float64  `json:"company_id"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Subtitle string  `json:"subtitle"`
            Image string  `json:"image"`
            Title string  `json:"title"`
            CreatedOn string  `json:"created_on"`
            Tag string  `json:"tag"`
            ID string  `json:"id"`
            Guide Guide  `json:"guide"`
         
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

        
            BrandID float64  `json:"brand_id"`
            Active bool  `json:"active"`
            Name string  `json:"name"`
            ModifiedOn string  `json:"modified_on"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            CompanyID float64  `json:"company_id"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Subtitle string  `json:"subtitle"`
            Title string  `json:"title"`
            CreatedOn string  `json:"created_on"`
            Tag string  `json:"tag"`
            ID string  `json:"id"`
            Guide map[string]interface{}  `json:"guide"`
         
    }
    
    // ApplicationItemSEO used by Catalog
    type ApplicationItemSEO struct {

        
            Description string  `json:"description"`
            Title string  `json:"title"`
         
    }
    
    // ApplicationItemMOQ used by Catalog
    type ApplicationItemMOQ struct {

        
            Minimum float64  `json:"minimum"`
            Maximum float64  `json:"maximum"`
            IncrementUnit float64  `json:"increment_unit"`
         
    }
    
    // MetaFields used by Catalog
    type MetaFields struct {

        
            Value string  `json:"value"`
            Key string  `json:"key"`
         
    }
    
    // ApplicationItemMeta used by Catalog
    type ApplicationItemMeta struct {

        
            IsCod bool  `json:"is_cod"`
            Seo ApplicationItemSEO  `json:"seo"`
            Moq ApplicationItemMOQ  `json:"moq"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            CustomMeta []MetaFields  `json:"_custom_meta"`
            AltText map[string]interface{}  `json:"alt_text"`
            IsGift bool  `json:"is_gift"`
         
    }
    
    // SuccessResponse1 used by Catalog
    type SuccessResponse1 struct {

        
            Success bool  `json:"success"`
            UID float64  `json:"uid"`
         
    }
    
    // SEOData used by Catalog
    type SEOData struct {

        
            Description string  `json:"description"`
            Title string  `json:"title"`
         
    }
    
    // MOQData used by Catalog
    type MOQData struct {

        
            Minimum float64  `json:"minimum"`
            Maximum float64  `json:"maximum"`
            IncrementUnit float64  `json:"increment_unit"`
         
    }
    
    // OwnerAppItemResponse used by Catalog
    type OwnerAppItemResponse struct {

        
            IsCod bool  `json:"is_cod"`
            Seo SEOData  `json:"seo"`
            Moq MOQData  `json:"moq"`
            AltText map[string]interface{}  `json:"alt_text"`
            IsGift bool  `json:"is_gift"`
         
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
            Priority float64  `json:"priority"`
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Key string  `json:"key"`
            DisplayType string  `json:"display_type"`
            Unit string  `json:"unit"`
         
    }
    
    // AppConfigurationDetail used by Catalog
    type AppConfigurationDetail struct {

        
            Logo string  `json:"logo"`
            IsDefault bool  `json:"is_default"`
            Priority float64  `json:"priority"`
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            AppID string  `json:"app_id"`
            Attributes []AttributeDetailsGroup  `json:"attributes"`
            TemplateSlugs []string  `json:"template_slugs"`
         
    }
    
    // ConfigErrorResponse used by Catalog
    type ConfigErrorResponse struct {

        
            Message string  `json:"message"`
         
    }
    
    // PageResponseType used by Catalog
    type PageResponseType struct {

        
            HasNext bool  `json:"has_next"`
            TotalCount float64  `json:"total_count"`
            Next float64  `json:"next"`
            Current float64  `json:"current"`
         
    }
    
    // GetConfigResponse used by Catalog
    type GetConfigResponse struct {

        
            Data []map[string]interface{}  `json:"data"`
            Page PageResponseType  `json:"page"`
         
    }
    
    // ConfigSuccessResponse used by Catalog
    type ConfigSuccessResponse struct {

        
            Message string  `json:"message"`
         
    }
    
    // AppConfigurationsSort used by Catalog
    type AppConfigurationsSort struct {

        
            Logo string  `json:"logo"`
            IsDefault bool  `json:"is_default"`
            Priority float64  `json:"priority"`
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            DefaultKey string  `json:"default_key"`
            AppID string  `json:"app_id"`
            Key string  `json:"key"`
         
    }
    
    // AllowSingleRequest used by Catalog
    type AllowSingleRequest struct {

        
            AllowSingle bool  `json:"allow_single"`
         
    }
    
    // DefaultKeyRequest used by Catalog
    type DefaultKeyRequest struct {

        
            DefaultKey string  `json:"default_key"`
         
    }
    
    // ProductSortOn used by Catalog
    type ProductSortOn struct {

        
            Value string  `json:"value"`
            IsSelected bool  `json:"is_selected"`
            Name string  `json:"name"`
         
    }
    
    // ProductFiltersKey used by Catalog
    type ProductFiltersKey struct {

        
            Logo string  `json:"logo"`
            Name string  `json:"name"`
            Kind string  `json:"kind"`
            Display string  `json:"display"`
            Operators []string  `json:"operators"`
         
    }
    
    // ProductFiltersValue used by Catalog
    type ProductFiltersValue struct {

        
            SelectedMax float64  `json:"selected_max"`
            DisplayFormat string  `json:"display_format"`
            IsSelected bool  `json:"is_selected"`
            CurrencySymbol string  `json:"currency_symbol"`
            Min float64  `json:"min"`
            Value interface{}  `json:"value"`
            SelectedMin float64  `json:"selected_min"`
            Display string  `json:"display"`
            Count float64  `json:"count"`
            QueryFormat string  `json:"query_format"`
            Max float64  `json:"max"`
            CurrencyCode string  `json:"currency_code"`
         
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

        
            Value []interface{}  `json:"value"`
            Op string  `json:"op"`
            Attribute string  `json:"attribute"`
         
    }
    
    // CollectionImage used by Catalog
    type CollectionImage struct {

        
            URL string  `json:"url"`
            AspectRatio string  `json:"aspect_ratio"`
         
    }
    
    // NextSchedule used by Catalog
    type NextSchedule struct {

        
            End string  `json:"end"`
            Start string  `json:"start"`
         
    }
    
    // CollectionSchedule used by Catalog
    type CollectionSchedule struct {

        
            Cron string  `json:"cron"`
            NextSchedule []NextSchedule  `json:"next_schedule"`
            Start string  `json:"start"`
            End string  `json:"end"`
            Duration float64  `json:"duration"`
         
    }
    
    // CollectionBanner used by Catalog
    type CollectionBanner struct {

        
            Portrait CollectionImage  `json:"portrait"`
            Landscape CollectionImage  `json:"landscape"`
         
    }
    
    // UserInfo used by Catalog
    type UserInfo struct {

        
            Email string  `json:"email"`
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
            UID string  `json:"uid"`
         
    }
    
    // CreateCollection used by Catalog
    type CreateCollection struct {

        
            AllowFacets bool  `json:"allow_facets"`
            IsActive bool  `json:"is_active"`
            Priority float64  `json:"priority"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Published bool  `json:"published"`
            Seo SeoDetail  `json:"seo"`
            IsVisible bool  `json:"is_visible"`
            SortOn string  `json:"sort_on"`
            Type string  `json:"type"`
            Tags []string  `json:"tags"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Description string  `json:"description"`
            Badge CollectionBadge  `json:"badge"`
            Meta map[string]interface{}  `json:"meta"`
            Query []CollectionQuery  `json:"query"`
            Logo CollectionImage  `json:"logo"`
            Schedule CollectionSchedule  `json:"_schedule"`
            Banners CollectionBanner  `json:"banners"`
            CreatedBy UserInfo  `json:"created_by"`
            ModifiedBy UserInfo  `json:"modified_by"`
            AppID string  `json:"app_id"`
            AllowSort bool  `json:"allow_sort"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
         
    }
    
    // BannerImage used by Catalog
    type BannerImage struct {

        
            URL string  `json:"url"`
            Type string  `json:"type"`
            AspectRatio string  `json:"aspect_ratio"`
         
    }
    
    // ImageUrls used by Catalog
    type ImageUrls struct {

        
            Portrait BannerImage  `json:"portrait"`
            Landscape BannerImage  `json:"landscape"`
         
    }
    
    // CollectionCreateResponse used by Catalog
    type CollectionCreateResponse struct {

        
            AllowFacets bool  `json:"allow_facets"`
            IsActive bool  `json:"is_active"`
            Priority float64  `json:"priority"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Tag []string  `json:"tag"`
            SortOn string  `json:"sort_on"`
            Type string  `json:"type"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Description string  `json:"description"`
            Badge map[string]interface{}  `json:"badge"`
            Meta map[string]interface{}  `json:"meta"`
            Query []CollectionQuery  `json:"query"`
            Logo BannerImage  `json:"logo"`
            Schedule map[string]interface{}  `json:"_schedule"`
            Banners ImageUrls  `json:"banners"`
            Cron map[string]interface{}  `json:"cron"`
            AppID string  `json:"app_id"`
            AllowSort bool  `json:"allow_sort"`
         
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
    
    // Media1 used by Catalog
    type Media1 struct {

        
            URL string  `json:"url"`
            Type string  `json:"type"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // GetCollectionDetailNest used by Catalog
    type GetCollectionDetailNest struct {

        
            Priority float64  `json:"priority"`
            IsActive bool  `json:"is_active"`
            AllowFacets bool  `json:"allow_facets"`
            UID string  `json:"uid"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Tag []string  `json:"tag"`
            Type string  `json:"type"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Description string  `json:"description"`
            Badge map[string]interface{}  `json:"badge"`
            Action Action  `json:"action"`
            Meta map[string]interface{}  `json:"meta"`
            Query []CollectionQuery  `json:"query"`
            Logo Media1  `json:"logo"`
            Schedule map[string]interface{}  `json:"_schedule"`
            Banners ImageUrls  `json:"banners"`
            Cron map[string]interface{}  `json:"cron"`
            AppID string  `json:"app_id"`
            AllowSort bool  `json:"allow_sort"`
         
    }
    
    // CollectionListingFilterTag used by Catalog
    type CollectionListingFilterTag struct {

        
            IsSelected bool  `json:"is_selected"`
            Name string  `json:"name"`
            Display string  `json:"display"`
         
    }
    
    // CollectionListingFilterType used by Catalog
    type CollectionListingFilterType struct {

        
            IsSelected bool  `json:"is_selected"`
            Name string  `json:"name"`
            Display string  `json:"display"`
         
    }
    
    // CollectionListingFilter used by Catalog
    type CollectionListingFilter struct {

        
            Tags []CollectionListingFilterTag  `json:"tags"`
            Type []CollectionListingFilterType  `json:"type"`
         
    }
    
    // GetCollectionListingResponse used by Catalog
    type GetCollectionListingResponse struct {

        
            Items []GetCollectionDetailNest  `json:"items"`
            Page Page  `json:"page"`
            Filters CollectionListingFilter  `json:"filters"`
         
    }
    
    // CollectionDetailResponse used by Catalog
    type CollectionDetailResponse struct {

        
            Priority float64  `json:"priority"`
            IsActive bool  `json:"is_active"`
            AllowFacets bool  `json:"allow_facets"`
            UID string  `json:"uid"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Tag []string  `json:"tag"`
            Type string  `json:"type"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Description string  `json:"description"`
            Badge map[string]interface{}  `json:"badge"`
            Meta map[string]interface{}  `json:"meta"`
            Query []CollectionQuery  `json:"query"`
            Logo Media1  `json:"logo"`
            Schedule map[string]interface{}  `json:"_schedule"`
            Banners ImageUrls  `json:"banners"`
            Cron map[string]interface{}  `json:"cron"`
            AppID string  `json:"app_id"`
            AllowSort bool  `json:"allow_sort"`
         
    }
    
    // UpdateCollection used by Catalog
    type UpdateCollection struct {

        
            AllowFacets bool  `json:"allow_facets"`
            IsActive bool  `json:"is_active"`
            Priority float64  `json:"priority"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Published bool  `json:"published"`
            Seo SeoDetail  `json:"seo"`
            IsVisible bool  `json:"is_visible"`
            SortOn string  `json:"sort_on"`
            Tags []string  `json:"tags"`
            Type string  `json:"type"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Description string  `json:"description"`
            Badge CollectionBadge  `json:"badge"`
            Meta map[string]interface{}  `json:"meta"`
            Query []CollectionQuery  `json:"query"`
            Logo CollectionImage  `json:"logo"`
            Schedule CollectionSchedule  `json:"_schedule"`
            Banners CollectionBanner  `json:"banners"`
            ModifiedBy UserInfo  `json:"modified_by"`
            AllowSort bool  `json:"allow_sort"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
         
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

        
            ItemsNotUpdated []float64  `json:"items_not_updated"`
            Message string  `json:"message"`
         
    }
    
    // ProductBrand used by Catalog
    type ProductBrand struct {

        
            Logo Media1  `json:"logo"`
            Action Action  `json:"action"`
            Name string  `json:"name"`
            UID float64  `json:"uid"`
         
    }
    
    // Price1 used by Catalog
    type Price1 struct {

        
            CurrencySymbol string  `json:"currency_symbol"`
            Max float64  `json:"max"`
            CurrencyCode string  `json:"currency_code"`
            Min float64  `json:"min"`
         
    }
    
    // ProductListingPrice used by Catalog
    type ProductListingPrice struct {

        
            Marked Price1  `json:"marked"`
            Effective Price1  `json:"effective"`
         
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
    
    // ProductListingDetail used by Catalog
    type ProductListingDetail struct {

        
            HasVariant bool  `json:"has_variant"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            UID float64  `json:"uid"`
            ItemCode string  `json:"item_code"`
            Similars []string  `json:"similars"`
            ImageNature string  `json:"image_nature"`
            Brand ProductBrand  `json:"brand"`
            Highlights []string  `json:"highlights"`
            Discount string  `json:"discount"`
            Medias []Media1  `json:"medias"`
            Price ProductListingPrice  `json:"price"`
            Sellable bool  `json:"sellable"`
            Type string  `json:"type"`
            RatingCount float64  `json:"rating_count"`
            Rating float64  `json:"rating"`
            ItemType string  `json:"item_type"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Description string  `json:"description"`
            Tryouts []string  `json:"tryouts"`
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
            Color string  `json:"color"`
            PromoMeta map[string]interface{}  `json:"promo_meta"`
            ShortDescription string  `json:"short_description"`
            ProductOnlineDate string  `json:"product_online_date"`
            Attributes map[string]interface{}  `json:"attributes"`
         
    }
    
    // GetCollectionItemsResponse used by Catalog
    type GetCollectionItemsResponse struct {

        
            Items []ProductListingDetail  `json:"items"`
            SortOn []ProductSortOn  `json:"sort_on"`
            Page Page  `json:"page"`
            Filters []ProductFilters  `json:"filters"`
         
    }
    
    // CatalogInsightItem used by Catalog
    type CatalogInsightItem struct {

        
            SellableCount float64  `json:"sellable_count"`
            OutOfStockCount float64  `json:"out_of_stock_count"`
            Count float64  `json:"count"`
         
    }
    
    // CatalogInsightBrand used by Catalog
    type CatalogInsightBrand struct {

        
            Name string  `json:"name"`
            TotalSizes float64  `json:"total_sizes"`
            AvailableSizes float64  `json:"available_sizes"`
            AvailableArticles float64  `json:"available_articles"`
            ArticleFreshness float64  `json:"article_freshness"`
            TotalArticles float64  `json:"total_articles"`
         
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

        
            BrandIds []float64  `json:"brand_ids"`
            CompanyID float64  `json:"company_id"`
            OptLevel string  `json:"opt_level"`
            StoreIds []float64  `json:"store_ids"`
            Enabled bool  `json:"enabled"`
            Platform string  `json:"platform"`
         
    }
    
    // CompanyOptIn used by Catalog
    type CompanyOptIn struct {

        
            BrandIds []float64  `json:"brand_ids"`
            ModifiedOn float64  `json:"modified_on"`
            CompanyID float64  `json:"company_id"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            OptLevel string  `json:"opt_level"`
            StoreIds []float64  `json:"store_ids"`
            CreatedOn float64  `json:"created_on"`
            Enabled bool  `json:"enabled"`
            Platform string  `json:"platform"`
         
    }
    
    // GetOptInPlatform used by Catalog
    type GetOptInPlatform struct {

        
            Items []CompanyOptIn  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // OptinCompanyDetail used by Catalog
    type OptinCompanyDetail struct {

        
            BusinessType string  `json:"business_type"`
            Name string  `json:"name"`
            CompanyType string  `json:"company_type"`
            UID float64  `json:"uid"`
         
    }
    
    // CompanyBrandDetail used by Catalog
    type CompanyBrandDetail struct {

        
            CompanyID float64  `json:"company_id"`
            BrandID float64  `json:"brand_id"`
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

        
            Store float64  `json:"store"`
            Company string  `json:"company"`
            Brand float64  `json:"brand"`
         
    }
    
    // StoreDetail used by Catalog
    type StoreDetail struct {

        
            Name string  `json:"name"`
            UID float64  `json:"uid"`
            ModifiedOn string  `json:"modified_on"`
            CompanyID float64  `json:"company_id"`
            StoreCode string  `json:"store_code"`
            Manager map[string]interface{}  `json:"manager"`
            Timing map[string]interface{}  `json:"timing"`
            StoreType string  `json:"store_type"`
            CreatedOn string  `json:"created_on"`
            AdditionalContacts []map[string]interface{}  `json:"additional_contacts"`
            DisplayName string  `json:"display_name"`
            Documents []map[string]interface{}  `json:"documents"`
            Address map[string]interface{}  `json:"address"`
         
    }
    
    // OptinStoreDetails used by Catalog
    type OptinStoreDetails struct {

        
            Items []StoreDetail  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // AttributeMasterFilter used by Catalog
    type AttributeMasterFilter struct {

        
            Indexing bool  `json:"indexing"`
            Priority float64  `json:"priority"`
            DependsOn []string  `json:"depends_on"`
         
    }
    
    // AttributeSchemaRange used by Catalog
    type AttributeSchemaRange struct {

        
            Max float64  `json:"max"`
            Min float64  `json:"min"`
         
    }
    
    // AttributeMaster used by Catalog
    type AttributeMaster struct {

        
            Mandatory bool  `json:"mandatory"`
            Format string  `json:"format"`
            Range AttributeSchemaRange  `json:"range"`
            AllowedValues []string  `json:"allowed_values"`
            Type string  `json:"type"`
            Multi bool  `json:"multi"`
         
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
    
    // GenderDetail used by Catalog
    type GenderDetail struct {

        
            Logo string  `json:"logo"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Filters AttributeMasterFilter  `json:"filters"`
            Description string  `json:"description"`
            Schema AttributeMaster  `json:"schema"`
            IsNested bool  `json:"is_nested"`
            EnabledForEndConsumer bool  `json:"enabled_for_end_consumer"`
            Details AttributeMasterDetails  `json:"details"`
            Meta AttributeMasterMeta  `json:"meta"`
            Departments []string  `json:"departments"`
            ID string  `json:"id"`
         
    }
    
    // ProductTemplateCategory used by Catalog
    type ProductTemplateCategory struct {

        
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            SlugKey string  `json:"slug_key"`
            UID float64  `json:"uid"`
            TemplateSlug string  `json:"template_slug"`
         
    }
    
    // ProdcutTemplateCategoriesResponse used by Catalog
    type ProdcutTemplateCategoriesResponse struct {

        
            Items []ProductTemplateCategory  `json:"items"`
         
    }
    
    // PTErrorResponse used by Catalog
    type PTErrorResponse struct {

        
            Message string  `json:"message"`
            Errors map[string]interface{}  `json:"errors"`
            Status float64  `json:"status"`
            Meta map[string]interface{}  `json:"meta"`
            Code string  `json:"code"`
         
    }
    
    // DepartmentCreateUpdate used by Catalog
    type DepartmentCreateUpdate struct {

        
            Logo string  `json:"logo"`
            Name string  `json:"name"`
            IsActive bool  `json:"is_active"`
            Platforms map[string]interface{}  `json:"platforms"`
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            PriorityOrder float64  `json:"priority_order"`
            Synonyms []string  `json:"synonyms"`
            Cls string  `json:"_cls"`
            Tags []string  `json:"tags"`
         
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
            Username string  `json:"username"`
            ID string  `json:"_id"`
            UserID string  `json:"user_id"`
            Contact string  `json:"contact"`
         
    }
    
    // GetDepartment used by Catalog
    type GetDepartment struct {

        
            Logo string  `json:"logo"`
            ItemType string  `json:"item_type"`
            Name string  `json:"name"`
            IsActive bool  `json:"is_active"`
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
            PageNo float64  `json:"page_no"`
            CreatedBy UserSerializer  `json:"created_by"`
            ModifiedOn string  `json:"modified_on"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            PriorityOrder float64  `json:"priority_order"`
            Synonyms []string  `json:"synonyms"`
            CreatedOn string  `json:"created_on"`
            PageSize float64  `json:"page_size"`
            Search string  `json:"search"`
         
    }
    
    // DepartmentsResponse used by Catalog
    type DepartmentsResponse struct {

        
            Items []GetDepartment  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // DepartmentErrorResponse used by Catalog
    type DepartmentErrorResponse struct {

        
            Message string  `json:"message"`
            Errors map[string]interface{}  `json:"errors"`
            Status float64  `json:"status"`
            Meta map[string]interface{}  `json:"meta"`
            Code string  `json:"code"`
         
    }
    
    // UserDetail used by Catalog
    type UserDetail struct {

        
            UserID string  `json:"user_id"`
            SuperUser bool  `json:"super_user"`
            Contact string  `json:"contact"`
            Username string  `json:"username"`
         
    }
    
    // DepartmentModel used by Catalog
    type DepartmentModel struct {

        
            Logo string  `json:"logo"`
            Name string  `json:"name"`
            IsActive bool  `json:"is_active"`
            Slug string  `json:"slug"`
            ModifiedOn string  `json:"modified_on"`
            CreatedBy UserDetail  `json:"created_by"`
            UID float64  `json:"uid"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ModifiedBy UserDetail  `json:"modified_by"`
            PriorityOrder float64  `json:"priority_order"`
            Synonyms []string  `json:"synonyms"`
            VerifiedOn string  `json:"verified_on"`
            ID interface{}  `json:"_id"`
            VerifiedBy UserDetail  `json:"verified_by"`
            CreatedOn string  `json:"created_on"`
            Cls string  `json:"_cls"`
         
    }
    
    // ProductTemplate used by Catalog
    type ProductTemplate struct {

        
            Logo string  `json:"logo"`
            Name string  `json:"name"`
            IsActive bool  `json:"is_active"`
            IsPhysical bool  `json:"is_physical"`
            Slug string  `json:"slug"`
            ModifiedOn string  `json:"modified_on"`
            Description string  `json:"description"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            IsArchived bool  `json:"is_archived"`
            IsExpirable bool  `json:"is_expirable"`
            CreatedOn string  `json:"created_on"`
            Categories []string  `json:"categories"`
            Attributes []string  `json:"attributes"`
            Tag string  `json:"tag"`
            Departments []string  `json:"departments"`
         
    }
    
    // TemplatesResponse used by Catalog
    type TemplatesResponse struct {

        
            Items ProductTemplate  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // Properties used by Catalog
    type Properties struct {

        
            IsActive map[string]interface{}  `json:"is_active"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            IsDependent map[string]interface{}  `json:"is_dependent"`
            ReturnConfig map[string]interface{}  `json:"return_config"`
            ItemCode map[string]interface{}  `json:"item_code"`
            Command map[string]interface{}  `json:"command"`
            MultiSize map[string]interface{}  `json:"multi_size"`
            SizeGuide map[string]interface{}  `json:"size_guide"`
            TraderType map[string]interface{}  `json:"trader_type"`
            Trader map[string]interface{}  `json:"trader"`
            NoOfBoxes map[string]interface{}  `json:"no_of_boxes"`
            Sizes map[string]interface{}  `json:"sizes"`
            Highlights map[string]interface{}  `json:"highlights"`
            BrandUID map[string]interface{}  `json:"brand_uid"`
            Currency map[string]interface{}  `json:"currency"`
            Tags map[string]interface{}  `json:"tags"`
            Media map[string]interface{}  `json:"media"`
            ItemType map[string]interface{}  `json:"item_type"`
            Name map[string]interface{}  `json:"name"`
            Slug map[string]interface{}  `json:"slug"`
            CategorySlug map[string]interface{}  `json:"category_slug"`
            Description map[string]interface{}  `json:"description"`
            HsnCode map[string]interface{}  `json:"hsn_code"`
            Variants map[string]interface{}  `json:"variants"`
            ProductPublish map[string]interface{}  `json:"product_publish"`
            ProductGroupTag map[string]interface{}  `json:"product_group_tag"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            CountryOfOrigin map[string]interface{}  `json:"country_of_origin"`
            ShortDescription map[string]interface{}  `json:"short_description"`
         
    }
    
    // GlobalValidation used by Catalog
    type GlobalValidation struct {

        
            Description string  `json:"description"`
            Properties Properties  `json:"properties"`
            Title string  `json:"title"`
            Required []string  `json:"required"`
            Type string  `json:"type"`
            Definitions map[string]interface{}  `json:"definitions"`
         
    }
    
    // TemplateValidationData used by Catalog
    type TemplateValidationData struct {

        
            GlobalValidation GlobalValidation  `json:"global_validation"`
            TemplateValidation map[string]interface{}  `json:"template_validation"`
         
    }
    
    // TemplateDetails used by Catalog
    type TemplateDetails struct {

        
            Logo string  `json:"logo"`
            Name string  `json:"name"`
            IsActive bool  `json:"is_active"`
            IsPhysical bool  `json:"is_physical"`
            Slug string  `json:"slug"`
            Description string  `json:"description"`
            IsArchived bool  `json:"is_archived"`
            IsExpirable bool  `json:"is_expirable"`
            Categories []string  `json:"categories"`
            Attributes []string  `json:"attributes"`
            Tag string  `json:"tag"`
            Departments []string  `json:"departments"`
            ID string  `json:"id"`
         
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
    
    // VerifiedBy used by Catalog
    type VerifiedBy struct {

        
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
         
    }
    
    // ProductDownloadItemsData used by Catalog
    type ProductDownloadItemsData struct {

        
            Type string  `json:"type"`
            Templates []string  `json:"templates"`
            Brand []string  `json:"brand"`
         
    }
    
    // ProductDownloadsItems used by Catalog
    type ProductDownloadsItems struct {

        
            CreatedBy VerifiedBy  `json:"created_by"`
            Data ProductDownloadItemsData  `json:"data"`
            URL string  `json:"url"`
            SellerID float64  `json:"seller_id"`
            Status string  `json:"status"`
            TemplateTags map[string]interface{}  `json:"template_tags"`
            TriggerOn string  `json:"trigger_on"`
            CompletedOn string  `json:"completed_on"`
            TaskID string  `json:"task_id"`
            ID string  `json:"id"`
         
    }
    
    // ProductDownloadsResponse used by Catalog
    type ProductDownloadsResponse struct {

        
            Items ProductDownloadsItems  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // ProductTemplateExportFilterRequest used by Catalog
    type ProductTemplateExportFilterRequest struct {

        
            ToDate string  `json:"to_date"`
            Templates []string  `json:"templates"`
            Brands []string  `json:"brands"`
            FromDate string  `json:"from_date"`
            CatalogueTypes []string  `json:"catalogue_types"`
         
    }
    
    // ProductTemplateDownloadsExport used by Catalog
    type ProductTemplateDownloadsExport struct {

        
            NotificationEmails []string  `json:"notification_emails"`
            Type string  `json:"type"`
            Filters ProductTemplateExportFilterRequest  `json:"filters"`
         
    }
    
    // ProductConfigurationDownloads used by Catalog
    type ProductConfigurationDownloads struct {

        
            Data []map[string]interface{}  `json:"data"`
            Multivalue bool  `json:"multivalue"`
         
    }
    
    // Media2 used by Catalog
    type Media2 struct {

        
            Logo string  `json:"logo"`
            Portrait string  `json:"portrait"`
            Landscape string  `json:"landscape"`
         
    }
    
    // Hierarchy used by Catalog
    type Hierarchy struct {

        
            L1 float64  `json:"l1"`
            Department float64  `json:"department"`
            L2 float64  `json:"l2"`
         
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

        
            Media Media2  `json:"media"`
            Name string  `json:"name"`
            Priority float64  `json:"priority"`
            IsActive bool  `json:"is_active"`
            Slug string  `json:"slug"`
            Hierarchy []Hierarchy  `json:"hierarchy"`
            Marketplaces CategoryMapping  `json:"marketplaces"`
            Synonyms []string  `json:"synonyms"`
            Tryouts []string  `json:"tryouts"`
            Level float64  `json:"level"`
            Departments []float64  `json:"departments"`
         
    }
    
    // CategoryCreateResponse used by Catalog
    type CategoryCreateResponse struct {

        
            Message string  `json:"message"`
            UID float64  `json:"uid"`
         
    }
    
    // Category used by Catalog
    type Category struct {

        
            Media Media2  `json:"media"`
            Name string  `json:"name"`
            Priority float64  `json:"priority"`
            IsActive bool  `json:"is_active"`
            Slug string  `json:"slug"`
            Hierarchy []Hierarchy  `json:"hierarchy"`
            UID float64  `json:"uid"`
            Marketplaces CategoryMapping  `json:"marketplaces"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            ModifiedOn string  `json:"modified_on"`
            Synonyms []string  `json:"synonyms"`
            Tryouts []string  `json:"tryouts"`
            CreatedOn string  `json:"created_on"`
            Level float64  `json:"level"`
            Departments []float64  `json:"departments"`
            ID string  `json:"id"`
         
    }
    
    // CategoryResponse used by Catalog
    type CategoryResponse struct {

        
            Items []Category  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // CategoryUpdateResponse used by Catalog
    type CategoryUpdateResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // SingleCategoryResponse used by Catalog
    type SingleCategoryResponse struct {

        
            Data Category  `json:"data"`
         
    }
    
    // TeaserTag used by Catalog
    type TeaserTag struct {

        
            Tag string  `json:"tag"`
            URL string  `json:"url"`
         
    }
    
    // NetQuantity used by Catalog
    type NetQuantity struct {

        
            Value float64  `json:"value"`
            Unit string  `json:"unit"`
         
    }
    
    // ReturnConfig used by Catalog
    type ReturnConfig struct {

        
            Time float64  `json:"time"`
            Returnable bool  `json:"returnable"`
            Unit string  `json:"unit"`
         
    }
    
    // Trader used by Catalog
    type Trader struct {

        
            Name string  `json:"name"`
            Type string  `json:"type"`
            Address []string  `json:"address"`
         
    }
    
    // TaxIdentifier used by Catalog
    type TaxIdentifier struct {

        
            ReportingHsn string  `json:"reporting_hsn"`
            HsnCodeID string  `json:"hsn_code_id"`
            HsnCode string  `json:"hsn_code"`
         
    }
    
    // ProductPublish used by Catalog
    type ProductPublish struct {

        
            ProductOnlineDate string  `json:"product_online_date"`
            IsSet bool  `json:"is_set"`
         
    }
    
    // CustomOrder used by Catalog
    type CustomOrder struct {

        
            IsCustomOrder bool  `json:"is_custom_order"`
            ManufacturingTimeUnit string  `json:"manufacturing_time_unit"`
            ManufacturingTime float64  `json:"manufacturing_time"`
         
    }
    
    // ProductCreateUpdateSchemaV2 used by Catalog
    type ProductCreateUpdateSchemaV2 struct {

        
            IsActive bool  `json:"is_active"`
            TeaserTag TeaserTag  `json:"teaser_tag"`
            UID float64  `json:"uid"`
            CompanyID float64  `json:"company_id"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            NetQuantity NetQuantity  `json:"net_quantity"`
            IsDependent bool  `json:"is_dependent"`
            ReturnConfig ReturnConfig  `json:"return_config"`
            ItemCode string  `json:"item_code"`
            MultiSize bool  `json:"multi_size"`
            SizeGuide string  `json:"size_guide"`
            Trader []Trader  `json:"trader"`
            NoOfBoxes float64  `json:"no_of_boxes"`
            Sizes []map[string]interface{}  `json:"sizes"`
            Highlights []string  `json:"highlights"`
            BrandUID float64  `json:"brand_uid"`
            TemplateTag string  `json:"template_tag"`
            Currency string  `json:"currency"`
            IsImageLessProduct bool  `json:"is_image_less_product"`
            Tags []string  `json:"tags"`
            ChangeRequestID interface{}  `json:"change_request_id"`
            IsSet bool  `json:"is_set"`
            TaxIdentifier TaxIdentifier  `json:"tax_identifier"`
            Media []Media1  `json:"media"`
            ItemType string  `json:"item_type"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            CategorySlug string  `json:"category_slug"`
            Description string  `json:"description"`
            Action string  `json:"action"`
            Variants map[string]interface{}  `json:"variants"`
            Departments []float64  `json:"departments"`
            ProductPublish ProductPublish  `json:"product_publish"`
            ProductGroupTag []string  `json:"product_group_tag"`
            CustomOrder CustomOrder  `json:"custom_order"`
            CountryOfOrigin string  `json:"country_of_origin"`
            ShortDescription string  `json:"short_description"`
            VariantGroup map[string]interface{}  `json:"variant_group"`
            Attributes map[string]interface{}  `json:"attributes"`
            VariantMedia map[string]interface{}  `json:"variant_media"`
            Requester string  `json:"requester"`
            BulkJobID string  `json:"bulk_job_id"`
         
    }
    
    // NetQuantityResponse used by Catalog
    type NetQuantityResponse struct {

        
            Value float64  `json:"value"`
            Unit string  `json:"unit"`
         
    }
    
    // ReturnConfigResponse used by Catalog
    type ReturnConfigResponse struct {

        
            Time float64  `json:"time"`
            Returnable bool  `json:"returnable"`
            Unit string  `json:"unit"`
         
    }
    
    // Logo used by Catalog
    type Logo struct {

        
            URL string  `json:"url"`
            AspectRatioF float64  `json:"aspect_ratio_f"`
            AspectRatio string  `json:"aspect_ratio"`
            SecureURL string  `json:"secure_url"`
         
    }
    
    // Brand used by Catalog
    type Brand struct {

        
            Logo Logo  `json:"logo"`
            Name string  `json:"name"`
            UID float64  `json:"uid"`
         
    }
    
    // ProductPublish1 used by Catalog
    type ProductPublish1 struct {

        
            ProductOnlineDate string  `json:"product_online_date"`
            IsSet bool  `json:"is_set"`
         
    }
    
    // Image used by Catalog
    type Image struct {

        
            URL string  `json:"url"`
            AspectRatioF float64  `json:"aspect_ratio_f"`
            AspectRatio string  `json:"aspect_ratio"`
            SecureURL string  `json:"secure_url"`
         
    }
    
    // ProductSchemaV2 used by Catalog
    type ProductSchemaV2 struct {

        
            AllCompanyIds []float64  `json:"all_company_ids"`
            IsActive bool  `json:"is_active"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            UID float64  `json:"uid"`
            CompanyID float64  `json:"company_id"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            NetQuantity NetQuantityResponse  `json:"net_quantity"`
            IsDependent bool  `json:"is_dependent"`
            ReturnConfig ReturnConfigResponse  `json:"return_config"`
            Pending string  `json:"pending"`
            ItemCode string  `json:"item_code"`
            MultiSize bool  `json:"multi_size"`
            SizeGuide string  `json:"size_guide"`
            Trader []map[string]interface{}  `json:"trader"`
            NoOfBoxes float64  `json:"no_of_boxes"`
            Stage string  `json:"stage"`
            ID string  `json:"id"`
            Sizes []map[string]interface{}  `json:"sizes"`
            Highlights []string  `json:"highlights"`
            BrandUID float64  `json:"brand_uid"`
            ImageNature string  `json:"image_nature"`
            Brand Brand  `json:"brand"`
            Currency string  `json:"currency"`
            TemplateTag string  `json:"template_tag"`
            IsImageLessProduct bool  `json:"is_image_less_product"`
            CreatedOn string  `json:"created_on"`
            Tags []string  `json:"tags"`
            L3Mapping []string  `json:"l3_mapping"`
            AllIdentifiers []string  `json:"all_identifiers"`
            IsSet bool  `json:"is_set"`
            Media []Media1  `json:"media"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            ItemType string  `json:"item_type"`
            Name string  `json:"name"`
            PrimaryColor string  `json:"primary_color"`
            Slug string  `json:"slug"`
            CategorySlug string  `json:"category_slug"`
            Description string  `json:"description"`
            ModifiedOn string  `json:"modified_on"`
            Moq map[string]interface{}  `json:"moq"`
            AllSizes []map[string]interface{}  `json:"all_sizes"`
            IsExpirable bool  `json:"is_expirable"`
            VerifiedBy VerifiedBy  `json:"verified_by"`
            VerifiedOn string  `json:"verified_on"`
            HsnCode string  `json:"hsn_code"`
            Variants map[string]interface{}  `json:"variants"`
            Departments []float64  `json:"departments"`
            IsPhysical bool  `json:"is_physical"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            ProductPublish ProductPublish1  `json:"product_publish"`
            Color string  `json:"color"`
            CategoryUID float64  `json:"category_uid"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            CountryOfOrigin string  `json:"country_of_origin"`
            Category map[string]interface{}  `json:"category"`
            ShortDescription string  `json:"short_description"`
            VariantGroup map[string]interface{}  `json:"variant_group"`
            Images []Image  `json:"images"`
            ProductGroupTag []string  `json:"product_group_tag"`
            Attributes map[string]interface{}  `json:"attributes"`
            VariantMedia map[string]interface{}  `json:"variant_media"`
         
    }
    
    // ProductListingResponseV2 used by Catalog
    type ProductListingResponseV2 struct {

        
            Items []ProductSchemaV2  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // ProductVariants used by Catalog
    type ProductVariants struct {

        
            Media []Media1  `json:"media"`
            Name string  `json:"name"`
            BrandUID float64  `json:"brand_uid"`
            UID float64  `json:"uid"`
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

        
            Variant bool  `json:"variant"`
            Unit string  `json:"unit"`
            Filters AttributeMasterFilter  `json:"filters"`
            RawKey string  `json:"raw_key"`
            CreatedOn string  `json:"created_on"`
            EnabledForEndConsumer bool  `json:"enabled_for_end_consumer"`
            Tags []string  `json:"tags"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            ModifiedOn string  `json:"modified_on"`
            Description string  `json:"description"`
            Details AttributeMasterDetails  `json:"details"`
            Departments []string  `json:"departments"`
            Logo string  `json:"logo"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Schema AttributeMaster  `json:"schema"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Suggestion string  `json:"suggestion"`
            IsNested bool  `json:"is_nested"`
            Synonyms map[string]interface{}  `json:"synonyms"`
         
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
            ItemDimensionsUnitOfMeasure string  `json:"item_dimensions_unit_of_measure"`
            ItemWeight float64  `json:"item_weight"`
            Identifiers []ValidateIdentifier  `json:"identifiers"`
            ItemWeightUnitOfMeasure string  `json:"item_weight_unit_of_measure"`
            ItemHeight float64  `json:"item_height"`
            Size string  `json:"size"`
            ItemLength float64  `json:"item_length"`
         
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

        
            AllCompanyIds []float64  `json:"all_company_ids"`
            IsActive bool  `json:"is_active"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            UID float64  `json:"uid"`
            CompanyID float64  `json:"company_id"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            NetQuantity NetQuantityResponse  `json:"net_quantity"`
            IsDependent bool  `json:"is_dependent"`
            ReturnConfig ReturnConfigResponse  `json:"return_config"`
            Pending string  `json:"pending"`
            ItemCode string  `json:"item_code"`
            MultiSize bool  `json:"multi_size"`
            SizeGuide string  `json:"size_guide"`
            Trader []map[string]interface{}  `json:"trader"`
            NoOfBoxes float64  `json:"no_of_boxes"`
            Stage string  `json:"stage"`
            ID string  `json:"id"`
            Sizes []map[string]interface{}  `json:"sizes"`
            Highlights []string  `json:"highlights"`
            BrandUID float64  `json:"brand_uid"`
            ImageNature string  `json:"image_nature"`
            Brand Brand  `json:"brand"`
            Currency string  `json:"currency"`
            TemplateTag string  `json:"template_tag"`
            IsImageLessProduct bool  `json:"is_image_less_product"`
            CreatedOn string  `json:"created_on"`
            Tags []string  `json:"tags"`
            L3Mapping []string  `json:"l3_mapping"`
            AllIdentifiers []string  `json:"all_identifiers"`
            IsSet bool  `json:"is_set"`
            Media []Media1  `json:"media"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            ItemType string  `json:"item_type"`
            Name string  `json:"name"`
            PrimaryColor string  `json:"primary_color"`
            Slug string  `json:"slug"`
            CategorySlug string  `json:"category_slug"`
            Description string  `json:"description"`
            ModifiedOn string  `json:"modified_on"`
            Moq map[string]interface{}  `json:"moq"`
            AllSizes []map[string]interface{}  `json:"all_sizes"`
            IsExpirable bool  `json:"is_expirable"`
            VerifiedBy VerifiedBy  `json:"verified_by"`
            VerifiedOn string  `json:"verified_on"`
            HsnCode string  `json:"hsn_code"`
            Variants map[string]interface{}  `json:"variants"`
            Departments []float64  `json:"departments"`
            IsPhysical bool  `json:"is_physical"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            ProductPublish ProductPublished  `json:"product_publish"`
            Color string  `json:"color"`
            CategoryUID float64  `json:"category_uid"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            CountryOfOrigin string  `json:"country_of_origin"`
            Category map[string]interface{}  `json:"category"`
            ShortDescription string  `json:"short_description"`
            VariantGroup map[string]interface{}  `json:"variant_group"`
            Images []Image  `json:"images"`
            ProductGroupTag []string  `json:"product_group_tag"`
            Attributes map[string]interface{}  `json:"attributes"`
            VariantMedia map[string]interface{}  `json:"variant_media"`
         
    }
    
    // ProductListingResponse used by Catalog
    type ProductListingResponse struct {

        
            Items []Product  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // UserInfo1 used by Catalog
    type UserInfo1 struct {

        
            Email string  `json:"email"`
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
            UID string  `json:"uid"`
         
    }
    
    // BulkJob used by Catalog
    type BulkJob struct {

        
            TrackingURL string  `json:"tracking_url"`
            IsActive bool  `json:"is_active"`
            ModifiedOn string  `json:"modified_on"`
            CreatedBy UserInfo1  `json:"created_by"`
            CompanyID float64  `json:"company_id"`
            Total float64  `json:"total"`
            ModifiedBy UserInfo1  `json:"modified_by"`
            CancelledRecords []map[string]interface{}  `json:"cancelled_records"`
            Failed float64  `json:"failed"`
            FailedRecords []map[string]interface{}  `json:"failed_records"`
            TemplateTag string  `json:"template_tag"`
            CreatedOn string  `json:"created_on"`
            Cancelled float64  `json:"cancelled"`
            Succeed float64  `json:"succeed"`
            FilePath string  `json:"file_path"`
            Stage string  `json:"stage"`
            CustomTemplateTag string  `json:"custom_template_tag"`
         
    }
    
    // BulkResponse used by Catalog
    type BulkResponse struct {

        
            IsActive bool  `json:"is_active"`
            ModifiedOn string  `json:"modified_on"`
            CreatedBy UserInfo1  `json:"created_by"`
            BatchID string  `json:"batch_id"`
            ModifiedBy UserInfo1  `json:"modified_by"`
            CreatedOn string  `json:"created_on"`
         
    }
    
    // UserDetail1 used by Catalog
    type UserDetail1 struct {

        
            UserID string  `json:"user_id"`
            FullName string  `json:"full_name"`
            Username string  `json:"username"`
         
    }
    
    // ProductBulkRequest used by Catalog
    type ProductBulkRequest struct {

        
            IsActive bool  `json:"is_active"`
            CancelledRecords []string  `json:"cancelled_records"`
            CompanyID float64  `json:"company_id"`
            TemplateTag string  `json:"template_tag"`
            ModifiedOn string  `json:"modified_on"`
            CreatedBy UserDetail1  `json:"created_by"`
            Total float64  `json:"total"`
            Failed float64  `json:"failed"`
            FailedRecords []string  `json:"failed_records"`
            ModifiedBy UserDetail1  `json:"modified_by"`
            CreatedOn string  `json:"created_on"`
            FilePath string  `json:"file_path"`
            Succeed float64  `json:"succeed"`
            Cancelled float64  `json:"cancelled"`
            Template ProductTemplate  `json:"template"`
            Stage string  `json:"stage"`
         
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

        
            CompanyID float64  `json:"company_id"`
            URL string  `json:"url"`
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

        
            TrackingURL string  `json:"tracking_url"`
            IsActive bool  `json:"is_active"`
            ModifiedOn string  `json:"modified_on"`
            CreatedBy UserCommon  `json:"created_by"`
            CancelledRecords []string  `json:"cancelled_records"`
            Total float64  `json:"total"`
            ModifiedBy UserCommon  `json:"modified_by"`
            CompanyID float64  `json:"company_id"`
            Failed float64  `json:"failed"`
            FailedRecords []string  `json:"failed_records"`
            CreatedOn string  `json:"created_on"`
            FilePath string  `json:"file_path"`
            Succeed float64  `json:"succeed"`
            Cancelled float64  `json:"cancelled"`
            Retry float64  `json:"retry"`
            Stage string  `json:"stage"`
            ID string  `json:"id"`
         
    }
    
    // BulkAssetResponse used by Catalog
    type BulkAssetResponse struct {

        
            Items []Items  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // ProductSizeDeleteDataResponse used by Catalog
    type ProductSizeDeleteDataResponse struct {

        
            CompanyID float64  `json:"company_id"`
            Size string  `json:"size"`
            ItemID float64  `json:"item_id"`
         
    }
    
    // ProductSizeDeleteResponse used by Catalog
    type ProductSizeDeleteResponse struct {

        
            Data ProductSizeDeleteDataResponse  `json:"data"`
            Success bool  `json:"success"`
         
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

        
            GtinValue string  `json:"gtin_value"`
            Primary bool  `json:"primary"`
            GtinType string  `json:"gtin_type"`
         
    }
    
    // InvSize used by Catalog
    type InvSize struct {

        
            PriceEffective float64  `json:"price_effective"`
            ItemWidth float64  `json:"item_width"`
            Set InventorySet  `json:"set"`
            Currency string  `json:"currency"`
            ItemDimensionsUnitOfMeasure string  `json:"item_dimensions_unit_of_measure"`
            Price float64  `json:"price"`
            StoreCode string  `json:"store_code"`
            ItemWeight float64  `json:"item_weight"`
            ItemWeightUnitOfMeasure string  `json:"item_weight_unit_of_measure"`
            Identifiers []GTIN  `json:"identifiers"`
            Quantity float64  `json:"quantity"`
            ItemHeight float64  `json:"item_height"`
            Size string  `json:"size"`
            PriceTransfer float64  `json:"price_transfer"`
            ExpirationDate string  `json:"expiration_date"`
            IsSet bool  `json:"is_set"`
            ItemLength float64  `json:"item_length"`
         
    }
    
    // ItemQuery used by Catalog
    type ItemQuery struct {

        
            ItemCode string  `json:"item_code"`
            BrandUID float64  `json:"brand_uid"`
            UID float64  `json:"uid"`
         
    }
    
    // InventoryRequest used by Catalog
    type InventoryRequest struct {

        
            CompanyID float64  `json:"company_id"`
            Sizes []InvSize  `json:"sizes"`
            Item ItemQuery  `json:"item"`
         
    }
    
    // InventoryResponse used by Catalog
    type InventoryResponse struct {

        
            Size string  `json:"size"`
            Store map[string]interface{}  `json:"store"`
            UID string  `json:"uid"`
            Currency string  `json:"currency"`
            Price float64  `json:"price"`
            SellableQuantity float64  `json:"sellable_quantity"`
            InventoryUpdatedOn string  `json:"inventory_updated_on"`
            Identifiers map[string]interface{}  `json:"identifiers"`
            SellerIdentifier string  `json:"seller_identifier"`
            Quantity float64  `json:"quantity"`
            PriceEffective float64  `json:"price_effective"`
            PriceTransfer float64  `json:"price_transfer"`
            ItemID float64  `json:"item_id"`
         
    }
    
    // InventoryResponsePaginated used by Catalog
    type InventoryResponsePaginated struct {

        
            Items []InventoryResponse  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // ReturnConfig1 used by Catalog
    type ReturnConfig1 struct {

        
            Time float64  `json:"time"`
            Returnable bool  `json:"returnable"`
            Unit string  `json:"unit"`
         
    }
    
    // CompanyMeta used by Catalog
    type CompanyMeta struct {

        
            ID float64  `json:"id"`
         
    }
    
    // Trader1 used by Catalog
    type Trader1 struct {

        
            Name string  `json:"name"`
            Type string  `json:"type"`
            Address []string  `json:"address"`
         
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

        
            Transfer float64  `json:"transfer"`
            Effective float64  `json:"effective"`
            Currency string  `json:"currency"`
            Marked float64  `json:"marked"`
            UpdatedAt string  `json:"updated_at"`
            TpNotes map[string]interface{}  `json:"tp_notes"`
         
    }
    
    // ManufacturerResponse used by Catalog
    type ManufacturerResponse struct {

        
            IsDefault bool  `json:"is_default"`
            Name string  `json:"name"`
            Address string  `json:"address"`
         
    }
    
    // StoreMeta used by Catalog
    type StoreMeta struct {

        
            ID float64  `json:"id"`
         
    }
    
    // QuantityBase used by Catalog
    type QuantityBase struct {

        
            UpdatedAt string  `json:"updated_at"`
            Count float64  `json:"count"`
         
    }
    
    // Quantities used by Catalog
    type Quantities struct {

        
            OrderCommitted QuantityBase  `json:"order_committed"`
            Sellable QuantityBase  `json:"sellable"`
            Damaged QuantityBase  `json:"damaged"`
            NotAvailable QuantityBase  `json:"not_available"`
         
    }
    
    // DimensionResponse used by Catalog
    type DimensionResponse struct {

        
            IsDefault bool  `json:"is_default"`
            Length float64  `json:"length"`
            Unit string  `json:"unit"`
            Height float64  `json:"height"`
            Width float64  `json:"width"`
         
    }
    
    // InventorySellerResponse used by Catalog
    type InventorySellerResponse struct {

        
            IsActive bool  `json:"is_active"`
            UID string  `json:"uid"`
            RawMeta map[string]interface{}  `json:"raw_meta"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ReturnConfig ReturnConfig1  `json:"return_config"`
            FyndMeta map[string]interface{}  `json:"fynd_meta"`
            FyndArticleCode string  `json:"fynd_article_code"`
            Company CompanyMeta  `json:"company"`
            Trader []Trader1  `json:"trader"`
            Stage string  `json:"stage"`
            TotalQuantity float64  `json:"total_quantity"`
            Brand BrandMeta  `json:"brand"`
            Weight WeightResponse  `json:"weight"`
            Price PriceMeta  `json:"price"`
            SellerIdentifier string  `json:"seller_identifier"`
            Tags []string  `json:"tags"`
            TraceID string  `json:"trace_id"`
            Fragile bool  `json:"fragile"`
            Manufacturer ManufacturerResponse  `json:"manufacturer"`
            FyndItemCode string  `json:"fynd_item_code"`
            ItemID float64  `json:"item_id"`
            IsSet bool  `json:"is_set"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            Store StoreMeta  `json:"store"`
            Set InventorySet  `json:"set"`
            AddedOnStore string  `json:"added_on_store"`
            TrackInventory bool  `json:"track_inventory"`
            Meta map[string]interface{}  `json:"meta"`
            Identifier map[string]interface{}  `json:"identifier"`
            CreatedBy UserSerializer  `json:"created_by"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            Quantities Quantities  `json:"quantities"`
            CountryOfOrigin string  `json:"country_of_origin"`
            Dimension DimensionResponse  `json:"dimension"`
            Size string  `json:"size"`
            ExpirationDate string  `json:"expiration_date"`
         
    }
    
    // InventorySellerIdentifierResponsePaginated used by Catalog
    type InventorySellerIdentifierResponsePaginated struct {

        
            Items []InventorySellerResponse  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // ReturnConfig2 used by Catalog
    type ReturnConfig2 struct {

        
            Time float64  `json:"time"`
            Returnable bool  `json:"returnable"`
            Unit string  `json:"unit"`
         
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
    
    // DateMeta used by Catalog
    type DateMeta struct {

        
            InventoryUpdatedOn string  `json:"inventory_updated_on"`
            AddedOnStore string  `json:"added_on_store"`
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // BrandMeta1 used by Catalog
    type BrandMeta1 struct {

        
            Name string  `json:"name"`
            ID float64  `json:"id"`
         
    }
    
    // WeightResponse1 used by Catalog
    type WeightResponse1 struct {

        
            Unit string  `json:"unit"`
            Shipping float64  `json:"shipping"`
         
    }
    
    // PriceArticle used by Catalog
    type PriceArticle struct {

        
            Transfer float64  `json:"transfer"`
            Effective float64  `json:"effective"`
            Currency string  `json:"currency"`
            Marked float64  `json:"marked"`
            TpNotes map[string]interface{}  `json:"tp_notes"`
         
    }
    
    // ManufacturerResponse1 used by Catalog
    type ManufacturerResponse1 struct {

        
            IsDefault bool  `json:"is_default"`
            Name string  `json:"name"`
            Address string  `json:"address"`
         
    }
    
    // ArticleStoreResponse used by Catalog
    type ArticleStoreResponse struct {

        
            StoreType string  `json:"store_type"`
            Name string  `json:"name"`
            StoreCode string  `json:"store_code"`
            UID float64  `json:"uid"`
         
    }
    
    // Quantity used by Catalog
    type Quantity struct {

        
            Count float64  `json:"count"`
         
    }
    
    // QuantitiesArticle used by Catalog
    type QuantitiesArticle struct {

        
            OrderCommitted Quantity  `json:"order_committed"`
            Sellable Quantity  `json:"sellable"`
            Damaged Quantity  `json:"damaged"`
            NotAvailable Quantity  `json:"not_available"`
         
    }
    
    // DimensionResponse1 used by Catalog
    type DimensionResponse1 struct {

        
            Height float64  `json:"height"`
            Width float64  `json:"width"`
            Length float64  `json:"length"`
            Unit string  `json:"unit"`
         
    }
    
    // GetInventories used by Catalog
    type GetInventories struct {

        
            UID string  `json:"uid"`
            ReturnConfig ReturnConfig2  `json:"return_config"`
            Company CompanyMeta1  `json:"company"`
            Stage string  `json:"stage"`
            Trader []Trader2  `json:"trader"`
            TotalQuantity float64  `json:"total_quantity"`
            DateMeta DateMeta  `json:"date_meta"`
            ID string  `json:"id"`
            Brand BrandMeta1  `json:"brand"`
            Weight WeightResponse1  `json:"weight"`
            Price PriceArticle  `json:"price"`
            SellerIdentifier string  `json:"seller_identifier"`
            Tags []string  `json:"tags"`
            TraceID string  `json:"trace_id"`
            Manufacturer ManufacturerResponse1  `json:"manufacturer"`
            ItemID float64  `json:"item_id"`
            IsSet bool  `json:"is_set"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            Store ArticleStoreResponse  `json:"store"`
            Platforms map[string]interface{}  `json:"platforms"`
            InventoryUpdatedOn string  `json:"inventory_updated_on"`
            TrackInventory bool  `json:"track_inventory"`
            Identifier map[string]interface{}  `json:"identifier"`
            CreatedBy UserSerializer  `json:"created_by"`
            Quantities QuantitiesArticle  `json:"quantities"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            CountryOfOrigin string  `json:"country_of_origin"`
            Dimension DimensionResponse1  `json:"dimension"`
            Size string  `json:"size"`
            ExpirationDate string  `json:"expiration_date"`
         
    }
    
    // GetInventoriesResponse used by Catalog
    type GetInventoriesResponse struct {

        
            Items []GetInventories  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // BulkInventoryGetItems used by Catalog
    type BulkInventoryGetItems struct {

        
            IsActive bool  `json:"is_active"`
            CancelledRecords []string  `json:"cancelled_records"`
            CompanyID float64  `json:"company_id"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ModifiedOn string  `json:"modified_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Total float64  `json:"total"`
            Failed float64  `json:"failed"`
            FailedRecords []string  `json:"failed_records"`
            CreatedOn string  `json:"created_on"`
            Cancelled float64  `json:"cancelled"`
            FilePath string  `json:"file_path"`
            Succeed float64  `json:"succeed"`
            Stage string  `json:"stage"`
            ID string  `json:"id"`
         
    }
    
    // BulkInventoryGet used by Catalog
    type BulkInventoryGet struct {

        
            Items []BulkInventoryGetItems  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // InventoryJobPayload used by Catalog
    type InventoryJobPayload struct {

        
            PriceMarked float64  `json:"price_marked"`
            Currency string  `json:"currency"`
            ItemDimensionsUnitOfMeasure string  `json:"item_dimensions_unit_of_measure"`
            StoreCode string  `json:"store_code"`
            Price float64  `json:"price"`
            SellerIdentifier string  `json:"seller_identifier"`
            ItemWeightUnitOfMeasure string  `json:"item_weight_unit_of_measure"`
            Quantity float64  `json:"quantity"`
            Tags []string  `json:"tags"`
            TraceID string  `json:"trace_id"`
            PriceEffective float64  `json:"price_effective"`
            TotalQuantity float64  `json:"total_quantity"`
            ExpirationDate string  `json:"expiration_date"`
         
    }
    
    // InventoryBulkRequest used by Catalog
    type InventoryBulkRequest struct {

        
            CompanyID float64  `json:"company_id"`
            BatchID string  `json:"batch_id"`
            Sizes []InventoryJobPayload  `json:"sizes"`
            User map[string]interface{}  `json:"user"`
         
    }
    
    // InventoryExportRequest used by Catalog
    type InventoryExportRequest struct {

        
            Type string  `json:"type"`
            Store []float64  `json:"store"`
            Brand []float64  `json:"brand"`
         
    }
    
    // InventoryExportResponse used by Catalog
    type InventoryExportResponse struct {

        
            Filters map[string]interface{}  `json:"filters"`
            CreatedBy string  `json:"created_by"`
            ModifiedOn string  `json:"modified_on"`
            SellerID float64  `json:"seller_id"`
            NotificationEmails []string  `json:"notification_emails"`
            CreatedOn string  `json:"created_on"`
            Status string  `json:"status"`
            Type string  `json:"type"`
            TaskID string  `json:"task_id"`
         
    }
    
    // InventoryExportQuantityFilter used by Catalog
    type InventoryExportQuantityFilter struct {

        
            Operators string  `json:"operators"`
            Max float64  `json:"max"`
            Min float64  `json:"min"`
         
    }
    
    // InventoryExportAdvanceOption used by Catalog
    type InventoryExportAdvanceOption struct {

        
            BrandIds []float64  `json:"brand_ids"`
            ToDate string  `json:"to_date"`
            StoreIds []float64  `json:"store_ids"`
            Quantity InventoryExportQuantityFilter  `json:"quantity"`
            FromDate string  `json:"from_date"`
         
    }
    
    // InventoryExportJob used by Catalog
    type InventoryExportJob struct {

        
            Filters InventoryExportAdvanceOption  `json:"filters"`
            URL string  `json:"url"`
            SellerID float64  `json:"seller_id"`
            NotificationEmails []string  `json:"notification_emails"`
            Status string  `json:"status"`
            Type string  `json:"type"`
            CompletedOn string  `json:"completed_on"`
            TaskID string  `json:"task_id"`
         
    }
    
    // InventoryExportFilter used by Catalog
    type InventoryExportFilter struct {

        
            BrandIds []float64  `json:"brand_ids"`
            ToDate string  `json:"to_date"`
            StoreIds []float64  `json:"store_ids"`
            Quantity InventoryExportQuantityFilter  `json:"quantity"`
            FromDate string  `json:"from_date"`
         
    }
    
    // InventoryCreateRequest used by Catalog
    type InventoryCreateRequest struct {

        
            Data []string  `json:"data"`
            Type string  `json:"type"`
            Filters InventoryExportFilter  `json:"filters"`
            NotificationEmails []string  `json:"notification_emails"`
         
    }
    
    // InventoryJobFilters used by Catalog
    type InventoryJobFilters struct {

        
            ToDate string  `json:"to_date"`
            Stores []string  `json:"stores"`
            Quantity InventoryExportQuantityFilter  `json:"quantity"`
            Brands []string  `json:"brands"`
            FromDate string  `json:"from_date"`
         
    }
    
    // InventoryJobDetailResponse used by Catalog
    type InventoryJobDetailResponse struct {

        
            Filters InventoryJobFilters  `json:"filters"`
            CancelledBy UserDetail  `json:"cancelled_by"`
            URL string  `json:"url"`
            CreatedBy UserDetail  `json:"created_by"`
            ModifiedOn string  `json:"modified_on"`
            SellerID float64  `json:"seller_id"`
            NotificationEmails []string  `json:"notification_emails"`
            CreatedOn string  `json:"created_on"`
            Status string  `json:"status"`
            CancelledOn string  `json:"cancelled_on"`
            CompletedOn string  `json:"completed_on"`
            Type string  `json:"type"`
            TaskID string  `json:"task_id"`
            ID string  `json:"id"`
         
    }
    
    // InventoryExportJobListResponse used by Catalog
    type InventoryExportJobListResponse struct {

        
            Items InventoryJobDetailResponse  `json:"items"`
         
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
    
    // InventoryPayload used by Catalog
    type InventoryPayload struct {

        
            StoreID float64  `json:"store_id"`
            PriceMarked float64  `json:"price_marked"`
            SellerIdentifier string  `json:"seller_identifier"`
            Tags []string  `json:"tags"`
            TraceID string  `json:"trace_id"`
            PriceEffective float64  `json:"price_effective"`
            TotalQuantity float64  `json:"total_quantity"`
            ExpirationDate string  `json:"expiration_date"`
         
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

        
            IsActive bool  `json:"is_active"`
            UID float64  `json:"uid"`
            CompanyID float64  `json:"company_id"`
            Threshold1 float64  `json:"threshold1"`
            Threshold2 float64  `json:"threshold2"`
            Hs2Code string  `json:"hs2_code"`
            TaxOnMrp bool  `json:"tax_on_mrp"`
            HsnCode string  `json:"hsn_code"`
            Tax1 float64  `json:"tax1"`
            TaxOnEsp bool  `json:"tax_on_esp"`
            Tax2 float64  `json:"tax2"`
         
    }
    
    // HsnCodesObject used by Catalog
    type HsnCodesObject struct {

        
            ModifiedOn string  `json:"modified_on"`
            CompanyID float64  `json:"company_id"`
            Threshold1 float64  `json:"threshold1"`
            Threshold2 float64  `json:"threshold2"`
            Hs2Code string  `json:"hs2_code"`
            TaxOnMrp bool  `json:"tax_on_mrp"`
            HsnCode string  `json:"hsn_code"`
            Tax1 float64  `json:"tax1"`
            TaxOnEsp bool  `json:"tax_on_esp"`
            Tax2 float64  `json:"tax2"`
            ID string  `json:"id"`
         
    }
    
    // HsnCode used by Catalog
    type HsnCode struct {

        
            Data HsnCodesObject  `json:"data"`
         
    }
    
    // PageResponse used by Catalog
    type PageResponse struct {

        
            HasNext bool  `json:"has_next"`
            HasPrevious bool  `json:"has_previous"`
            Current string  `json:"current"`
            Size float64  `json:"size"`
            ItemTotal float64  `json:"item_total"`
         
    }
    
    // HsnCodesListingResponse used by Catalog
    type HsnCodesListingResponse struct {

        
            Items []HsnCodesObject  `json:"items"`
            Page PageResponse  `json:"page"`
         
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
            Cess float64  `json:"cess"`
            Threshold float64  `json:"threshold"`
            EffectiveDate string  `json:"effective_date"`
         
    }
    
    // HSNDataInsertV2 used by Catalog
    type HSNDataInsertV2 struct {

        
            Taxes []TaxSlab  `json:"taxes"`
            ModifiedOn string  `json:"modified_on"`
            Description string  `json:"description"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            ReportingHsn string  `json:"reporting_hsn"`
            CountryCode string  `json:"country_code"`
            Type string  `json:"type"`
            CreatedOn string  `json:"created_on"`
            HsnCode string  `json:"hsn_code"`
         
    }
    
    // HsnCodesListingResponseSchemaV2 used by Catalog
    type HsnCodesListingResponseSchemaV2 struct {

        
            Items []HSNDataInsertV2  `json:"items"`
            Page PageResponse  `json:"page"`
         
    }
    
    // BrandItem used by Catalog
    type BrandItem struct {

        
            Logo Media  `json:"logo"`
            Name string  `json:"name"`
            Banners ImageUrls  `json:"banners"`
            Slug string  `json:"slug"`
            Discount string  `json:"discount"`
            UID float64  `json:"uid"`
            Action Action  `json:"action"`
            Departments []string  `json:"departments"`
         
    }
    
    // BrandListingResponse used by Catalog
    type BrandListingResponse struct {

        
            Items []BrandItem  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // Department used by Catalog
    type Department struct {

        
            Logo Media  `json:"logo"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
            PriorityOrder float64  `json:"priority_order"`
         
    }
    
    // DepartmentResponse used by Catalog
    type DepartmentResponse struct {

        
            Items []Department  `json:"items"`
         
    }
    
    // ThirdLevelChild used by Catalog
    type ThirdLevelChild struct {

        
            Name string  `json:"name"`
            Banners ImageUrls  `json:"banners"`
            Childs []map[string]interface{}  `json:"childs"`
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Action Action  `json:"action"`
         
    }
    
    // SecondLevelChild used by Catalog
    type SecondLevelChild struct {

        
            Name string  `json:"name"`
            Banners ImageUrls  `json:"banners"`
            Childs []ThirdLevelChild  `json:"childs"`
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Action Action  `json:"action"`
         
    }
    
    // Child used by Catalog
    type Child struct {

        
            Name string  `json:"name"`
            Banners ImageUrls  `json:"banners"`
            Childs []SecondLevelChild  `json:"childs"`
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Action Action  `json:"action"`
         
    }
    
    // CategoryItems used by Catalog
    type CategoryItems struct {

        
            Name string  `json:"name"`
            Banners ImageUrls  `json:"banners"`
            Childs []Child  `json:"childs"`
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
            Action Action  `json:"action"`
         
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

        
            Items []ProductListingDetail  `json:"items"`
            Filters []ProductFilters  `json:"filters"`
            SortOn []ProductSortOn  `json:"sort_on"`
            Page Page  `json:"page"`
            Operators map[string]interface{}  `json:"operators"`
         
    }
    
    // ProductDetail used by Catalog
    type ProductDetail struct {

        
            HasVariant bool  `json:"has_variant"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            UID float64  `json:"uid"`
            ItemCode string  `json:"item_code"`
            Similars []string  `json:"similars"`
            ImageNature string  `json:"image_nature"`
            Brand ProductBrand  `json:"brand"`
            Highlights []string  `json:"highlights"`
            Medias []Media1  `json:"medias"`
            Type string  `json:"type"`
            RatingCount float64  `json:"rating_count"`
            Rating float64  `json:"rating"`
            ItemType string  `json:"item_type"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Description string  `json:"description"`
            Tryouts []string  `json:"tryouts"`
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
            Color string  `json:"color"`
            PromoMeta map[string]interface{}  `json:"promo_meta"`
            ShortDescription string  `json:"short_description"`
            ProductOnlineDate string  `json:"product_online_date"`
            Attributes map[string]interface{}  `json:"attributes"`
         
    }
    
    // InventoryPage used by Catalog
    type InventoryPage struct {

        
            HasNext bool  `json:"has_next"`
            HasPrevious bool  `json:"has_previous"`
            Type string  `json:"type"`
            NextID string  `json:"next_id"`
            ItemTotal float64  `json:"item_total"`
         
    }
    
    // InventoryStockResponse used by Catalog
    type InventoryStockResponse struct {

        
            Items []map[string]interface{}  `json:"items"`
            Page InventoryPage  `json:"page"`
         
    }
    
    // ArticleQuery used by Catalog
    type ArticleQuery struct {

        
            IgnoredStores []float64  `json:"ignored_stores"`
            Size string  `json:"size"`
            ItemID float64  `json:"item_id"`
         
    }
    
    // ArticleAssignment used by Catalog
    type ArticleAssignment struct {

        
            Strategy string  `json:"strategy"`
            Level string  `json:"level"`
         
    }
    
    // AssignStoreArticle used by Catalog
    type AssignStoreArticle struct {

        
            Quantity float64  `json:"quantity"`
            Meta map[string]interface{}  `json:"meta"`
            Query ArticleQuery  `json:"query"`
            ArticleAssignment ArticleAssignment  `json:"article_assignment"`
            GroupID string  `json:"group_id"`
         
    }
    
    // AssignStore used by Catalog
    type AssignStore struct {

        
            CompanyID float64  `json:"company_id"`
            StoreIds []float64  `json:"store_ids"`
            AppID string  `json:"app_id"`
            ChannelType string  `json:"channel_type"`
            ChannelIdentifier string  `json:"channel_identifier"`
            Pincode string  `json:"pincode"`
            Articles []AssignStoreArticle  `json:"articles"`
         
    }
    
    // ArticleAssignment1 used by Catalog
    type ArticleAssignment1 struct {

        
            Strategy string  `json:"strategy"`
            Level string  `json:"level"`
         
    }
    
    // StoreAssignResponse used by Catalog
    type StoreAssignResponse struct {

        
            Index float64  `json:"index"`
            StorePincode float64  `json:"store_pincode"`
            PriceEffective float64  `json:"price_effective"`
            StoreID float64  `json:"store_id"`
            PriceMarked float64  `json:"price_marked"`
            UID string  `json:"uid"`
            CompanyID float64  `json:"company_id"`
            StrategyWiseListing []map[string]interface{}  `json:"strategy_wise_listing"`
            ID string  `json:"_id"`
            Status bool  `json:"status"`
            Quantity float64  `json:"quantity"`
            Meta map[string]interface{}  `json:"meta"`
            SCity string  `json:"s_city"`
            Size string  `json:"size"`
            ItemID float64  `json:"item_id"`
            ArticleAssignment ArticleAssignment1  `json:"article_assignment"`
            GroupID string  `json:"group_id"`
         
    }
    
    // GetAddressSerializer used by Catalog
    type GetAddressSerializer struct {

        
            Address1 string  `json:"address1"`
            Longitude float64  `json:"longitude"`
            Address2 string  `json:"address2"`
            AddressType string  `json:"address_type"`
            City string  `json:"city"`
            Landmark string  `json:"landmark"`
            CountryCode string  `json:"country_code"`
            Country string  `json:"country"`
            Latitude float64  `json:"latitude"`
            Pincode float64  `json:"pincode"`
            State string  `json:"state"`
         
    }
    
    // UserSerializer1 used by Catalog
    type UserSerializer1 struct {

        
            UserID string  `json:"user_id"`
            Contact string  `json:"contact"`
            Username string  `json:"username"`
         
    }
    
    // GetCompanySerializer used by Catalog
    type GetCompanySerializer struct {

        
            Name string  `json:"name"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            Stage string  `json:"stage"`
            ModifiedOn string  `json:"modified_on"`
            UID float64  `json:"uid"`
            CreatedBy UserSerializer1  `json:"created_by"`
            ModifiedBy UserSerializer1  `json:"modified_by"`
            VerifiedOn string  `json:"verified_on"`
            VerifiedBy UserSerializer1  `json:"verified_by"`
            BusinessType string  `json:"business_type"`
            CreatedOn string  `json:"created_on"`
            RejectReason string  `json:"reject_reason"`
            CompanyType string  `json:"company_type"`
         
    }
    
    // LocationIntegrationType used by Catalog
    type LocationIntegrationType struct {

        
            Inventory string  `json:"inventory"`
            Order string  `json:"order"`
         
    }
    
    // SellerPhoneNumber used by Catalog
    type SellerPhoneNumber struct {

        
            CountryCode float64  `json:"country_code"`
            Number string  `json:"number"`
         
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
    
    // LocationTimingSerializer used by Catalog
    type LocationTimingSerializer struct {

        
            Hour float64  `json:"hour"`
            Minute float64  `json:"minute"`
         
    }
    
    // LocationDayWiseSerializer used by Catalog
    type LocationDayWiseSerializer struct {

        
            Open bool  `json:"open"`
            Closing LocationTimingSerializer  `json:"closing"`
            Weekday string  `json:"weekday"`
            Opening LocationTimingSerializer  `json:"opening"`
         
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
    
    // UserSerializer2 used by Catalog
    type UserSerializer2 struct {

        
            UserID string  `json:"user_id"`
            Contact string  `json:"contact"`
            Username string  `json:"username"`
         
    }
    
    // Document used by Catalog
    type Document struct {

        
            URL string  `json:"url"`
            Verified bool  `json:"verified"`
            Value string  `json:"value"`
            LegalName string  `json:"legal_name"`
            Type string  `json:"type"`
         
    }
    
    // GetLocationSerializer used by Catalog
    type GetLocationSerializer struct {

        
            UID float64  `json:"uid"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            StoreType string  `json:"store_type"`
            Company GetCompanySerializer  `json:"company"`
            IntegrationType LocationIntegrationType  `json:"integration_type"`
            Stage string  `json:"stage"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            GstCredentials InvoiceDetailsSerializer  `json:"gst_credentials"`
            Warnings map[string]interface{}  `json:"warnings"`
            Timing []LocationDayWiseSerializer  `json:"timing"`
            CreatedOn string  `json:"created_on"`
            Address GetAddressSerializer  `json:"address"`
            Name string  `json:"name"`
            ProductReturnConfig ProductReturnConfigSerializer  `json:"product_return_config"`
            ModifiedOn string  `json:"modified_on"`
            Manager LocationManagerSerializer  `json:"manager"`
            VerifiedOn string  `json:"verified_on"`
            VerifiedBy UserSerializer2  `json:"verified_by"`
            NotificationEmails []string  `json:"notification_emails"`
            Code string  `json:"code"`
            DisplayName string  `json:"display_name"`
            CreatedBy UserSerializer2  `json:"created_by"`
            ModifiedBy UserSerializer2  `json:"modified_by"`
            PhoneNumber string  `json:"phone_number"`
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

        
            Logo string  `json:"logo"`
            Name string  `json:"name"`
            IsActive bool  `json:"is_active"`
            UID float64  `json:"uid"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
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
    

    
    // UserSerializer used by CompanyProfile
    type UserSerializer struct {

        
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
            Contact string  `json:"contact"`
         
    }
    
    // Document used by CompanyProfile
    type Document struct {

        
            URL string  `json:"url"`
            LegalName string  `json:"legal_name"`
            Type string  `json:"type"`
            Verified bool  `json:"verified"`
            Value string  `json:"value"`
         
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

        
            EffectiveDate string  `json:"effective_date"`
            Enable bool  `json:"enable"`
            Rate float64  `json:"rate"`
         
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
    
    // GetAddressSerializer used by CompanyProfile
    type GetAddressSerializer struct {

        
            AddressType string  `json:"address_type"`
            Longitude float64  `json:"longitude"`
            Pincode float64  `json:"pincode"`
            Address1 string  `json:"address1"`
            Country string  `json:"country"`
            Address2 string  `json:"address2"`
            State string  `json:"state"`
            Landmark string  `json:"landmark"`
            CountryCode string  `json:"country_code"`
            Latitude float64  `json:"latitude"`
            City string  `json:"city"`
         
    }
    
    // BusinessCountryInfo used by CompanyProfile
    type BusinessCountryInfo struct {

        
            CountryCode string  `json:"country_code"`
            Country string  `json:"country"`
         
    }
    
    // GetCompanyProfileSerializerResponse used by CompanyProfile
    type GetCompanyProfileSerializerResponse struct {

        
            VerifiedBy UserSerializer  `json:"verified_by"`
            Documents []Document  `json:"documents"`
            ModifiedOn string  `json:"modified_on"`
            BusinessDetails BusinessDetails  `json:"business_details"`
            Taxes []CompanyTaxesSerializer  `json:"taxes"`
            ContactDetails ContactDetails  `json:"contact_details"`
            CreatedOn string  `json:"created_on"`
            VerifiedOn string  `json:"verified_on"`
            Mode string  `json:"mode"`
            Warnings map[string]interface{}  `json:"warnings"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            CompanyType string  `json:"company_type"`
            NotificationEmails []string  `json:"notification_emails"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            BusinessInfo string  `json:"business_info"`
            CreatedBy UserSerializer  `json:"created_by"`
            FranchiseEnabled bool  `json:"franchise_enabled"`
            UID float64  `json:"uid"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            Stage string  `json:"stage"`
            Name string  `json:"name"`
            BusinessCountryInfo BusinessCountryInfo  `json:"business_country_info"`
            BusinessType string  `json:"business_type"`
         
    }
    
    // ErrorResponse used by CompanyProfile
    type ErrorResponse struct {

        
            Meta map[string]interface{}  `json:"meta"`
            Message string  `json:"message"`
            Code string  `json:"code"`
            Status float64  `json:"status"`
         
    }
    
    // CompanyTaxesSerializer1 used by CompanyProfile
    type CompanyTaxesSerializer1 struct {

        
            EffectiveDate string  `json:"effective_date"`
            Enable bool  `json:"enable"`
            Rate float64  `json:"rate"`
         
    }
    
    // CreateUpdateAddressSerializer used by CompanyProfile
    type CreateUpdateAddressSerializer struct {

        
            AddressType string  `json:"address_type"`
            Longitude float64  `json:"longitude"`
            Pincode float64  `json:"pincode"`
            Address1 string  `json:"address1"`
            Country string  `json:"country"`
            State string  `json:"state"`
            Address2 string  `json:"address2"`
            Landmark string  `json:"landmark"`
            CountryCode string  `json:"country_code"`
            Latitude float64  `json:"latitude"`
            City string  `json:"city"`
         
    }
    
    // UpdateCompany used by CompanyProfile
    type UpdateCompany struct {

        
            BusinessInfo string  `json:"business_info"`
            Name string  `json:"name"`
            FranchiseEnabled bool  `json:"franchise_enabled"`
            NotificationEmails []string  `json:"notification_emails"`
            ContactDetails ContactDetails  `json:"contact_details"`
            Documents []Document  `json:"documents"`
            Taxes []CompanyTaxesSerializer1  `json:"taxes"`
            BusinessDetails BusinessDetails  `json:"business_details"`
            Addresses []CreateUpdateAddressSerializer  `json:"addresses"`
            CompanyType string  `json:"company_type"`
            BusinessType string  `json:"business_type"`
            RejectReason string  `json:"reject_reason"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Warnings map[string]interface{}  `json:"warnings"`
         
    }
    
    // ProfileSuccessResponse used by CompanyProfile
    type ProfileSuccessResponse struct {

        
            UID float64  `json:"uid"`
            Success bool  `json:"success"`
         
    }
    
    // DocumentsObj used by CompanyProfile
    type DocumentsObj struct {

        
            Pending float64  `json:"pending"`
            Verified float64  `json:"verified"`
         
    }
    
    // MetricsSerializer used by CompanyProfile
    type MetricsSerializer struct {

        
            CompanyDocuments DocumentsObj  `json:"company_documents"`
            Store DocumentsObj  `json:"store"`
            StoreDocuments DocumentsObj  `json:"store_documents"`
            UID float64  `json:"uid"`
            Product DocumentsObj  `json:"product"`
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

        
            VerifiedBy UserSerializer  `json:"verified_by"`
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
            VerifiedOn string  `json:"verified_on"`
            Mode string  `json:"mode"`
            Warnings map[string]interface{}  `json:"warnings"`
            Banner BrandBannerSerializer  `json:"banner"`
            RejectReason string  `json:"reject_reason"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            CreatedBy UserSerializer  `json:"created_by"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            UID float64  `json:"uid"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            Stage string  `json:"stage"`
            Description string  `json:"description"`
            Name string  `json:"name"`
            Synonyms []string  `json:"synonyms"`
            Logo string  `json:"logo"`
            SlugKey string  `json:"slug_key"`
         
    }
    
    // CreateUpdateBrandRequestSerializer used by CompanyProfile
    type CreateUpdateBrandRequestSerializer struct {

        
            Name string  `json:"name"`
            Synonyms []string  `json:"synonyms"`
            Logo string  `json:"logo"`
            Banner BrandBannerSerializer  `json:"banner"`
            BrandTier string  `json:"brand_tier"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            UID float64  `json:"uid"`
            CompanyID float64  `json:"company_id"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Description string  `json:"description"`
         
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
            Name string  `json:"name"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            CreatedBy UserSerializer  `json:"created_by"`
            NotificationEmails []string  `json:"notification_emails"`
            BusinessCountryInfo BusinessCountryInfo  `json:"business_country_info"`
            ModifiedOn string  `json:"modified_on"`
            UID float64  `json:"uid"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            Details CompanyDetails  `json:"details"`
            CompanyType string  `json:"company_type"`
            BusinessType string  `json:"business_type"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            CreatedOn string  `json:"created_on"`
            VerifiedOn string  `json:"verified_on"`
            RejectReason string  `json:"reject_reason"`
            MarketChannels []string  `json:"market_channels"`
            Stage string  `json:"stage"`
         
    }
    
    // CompanyBrandSerializer used by CompanyProfile
    type CompanyBrandSerializer struct {

        
            VerifiedBy UserSerializer  `json:"verified_by"`
            CreatedBy UserSerializer  `json:"created_by"`
            Company CompanySerializer  `json:"company"`
            ModifiedOn string  `json:"modified_on"`
            UID float64  `json:"uid"`
            Brand GetBrandResponseSerializer  `json:"brand"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            Stage string  `json:"stage"`
            RejectReason string  `json:"reject_reason"`
            CreatedOn string  `json:"created_on"`
            VerifiedOn string  `json:"verified_on"`
            Warnings map[string]interface{}  `json:"warnings"`
         
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
    
    // InvoiceCredSerializer used by CompanyProfile
    type InvoiceCredSerializer struct {

        
            Password string  `json:"password"`
            Username string  `json:"username"`
            Enabled bool  `json:"enabled"`
         
    }
    
    // InvoiceDetailsSerializer used by CompanyProfile
    type InvoiceDetailsSerializer struct {

        
            EInvoice InvoiceCredSerializer  `json:"e_invoice"`
            EWaybill InvoiceCredSerializer  `json:"e_waybill"`
         
    }
    
    // GetCompanySerializer used by CompanyProfile
    type GetCompanySerializer struct {

        
            VerifiedBy UserSerializer  `json:"verified_by"`
            Name string  `json:"name"`
            CreatedBy UserSerializer  `json:"created_by"`
            ModifiedOn string  `json:"modified_on"`
            UID float64  `json:"uid"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            Stage string  `json:"stage"`
            CompanyType string  `json:"company_type"`
            BusinessType string  `json:"business_type"`
            RejectReason string  `json:"reject_reason"`
            CreatedOn string  `json:"created_on"`
            VerifiedOn string  `json:"verified_on"`
            Addresses []GetAddressSerializer  `json:"addresses"`
         
    }
    
    // LocationManagerSerializer used by CompanyProfile
    type LocationManagerSerializer struct {

        
            Email string  `json:"email"`
            Name string  `json:"name"`
            MobileNo SellerPhoneNumber  `json:"mobile_no"`
         
    }
    
    // ProductReturnConfigSerializer used by CompanyProfile
    type ProductReturnConfigSerializer struct {

        
            OnSameStore bool  `json:"on_same_store"`
            StoreUID float64  `json:"store_uid"`
         
    }
    
    // LocationTimingSerializer used by CompanyProfile
    type LocationTimingSerializer struct {

        
            Minute float64  `json:"minute"`
            Hour float64  `json:"hour"`
         
    }
    
    // LocationDayWiseSerializer used by CompanyProfile
    type LocationDayWiseSerializer struct {

        
            Closing LocationTimingSerializer  `json:"closing"`
            Weekday string  `json:"weekday"`
            Opening LocationTimingSerializer  `json:"opening"`
            Open bool  `json:"open"`
         
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
    
    // GetLocationSerializer used by CompanyProfile
    type GetLocationSerializer struct {

        
            VerifiedBy UserSerializer  `json:"verified_by"`
            Documents []Document  `json:"documents"`
            ModifiedOn string  `json:"modified_on"`
            GstCredentials InvoiceDetailsSerializer  `json:"gst_credentials"`
            Company GetCompanySerializer  `json:"company"`
            Address GetAddressSerializer  `json:"address"`
            CreatedOn string  `json:"created_on"`
            VerifiedOn string  `json:"verified_on"`
            StoreType string  `json:"store_type"`
            Warnings map[string]interface{}  `json:"warnings"`
            PhoneNumber string  `json:"phone_number"`
            Code string  `json:"code"`
            NotificationEmails []string  `json:"notification_emails"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Manager LocationManagerSerializer  `json:"manager"`
            CreatedBy UserSerializer  `json:"created_by"`
            DisplayName string  `json:"display_name"`
            UID float64  `json:"uid"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            Stage string  `json:"stage"`
            ProductReturnConfig ProductReturnConfigSerializer  `json:"product_return_config"`
            Timing []LocationDayWiseSerializer  `json:"timing"`
            Name string  `json:"name"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            Holiday []HolidaySchemaSerializer  `json:"holiday"`
         
    }
    
    // LocationListSerializer used by CompanyProfile
    type LocationListSerializer struct {

        
            Items []GetLocationSerializer  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // AddressSerializer used by CompanyProfile
    type AddressSerializer struct {

        
            AddressType string  `json:"address_type"`
            Longitude float64  `json:"longitude"`
            Pincode float64  `json:"pincode"`
            Address1 string  `json:"address1"`
            Country string  `json:"country"`
            Address2 string  `json:"address2"`
            State string  `json:"state"`
            Landmark string  `json:"landmark"`
            CountryCode string  `json:"country_code"`
            Latitude float64  `json:"latitude"`
            City string  `json:"city"`
         
    }
    
    // LocationSerializer used by CompanyProfile
    type LocationSerializer struct {

        
            GstCredentials InvoiceDetailsSerializer  `json:"gst_credentials"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Name string  `json:"name"`
            Documents []Document  `json:"documents"`
            Code string  `json:"code"`
            DisplayName string  `json:"display_name"`
            Manager LocationManagerSerializer  `json:"manager"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            UID float64  `json:"uid"`
            Address AddressSerializer  `json:"address"`
            Stage string  `json:"stage"`
            ProductReturnConfig ProductReturnConfigSerializer  `json:"product_return_config"`
            NotificationEmails []string  `json:"notification_emails"`
            Holiday []HolidaySchemaSerializer  `json:"holiday"`
            Company float64  `json:"company"`
            StoreType string  `json:"store_type"`
            Timing []LocationDayWiseSerializer  `json:"timing"`
            Warnings map[string]interface{}  `json:"warnings"`
         
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

        
            Integration string  `json:"integration"`
            IntegrationData map[string]map[string]interface{}  `json:"integration_data"`
            CompanyName string  `json:"company_name"`
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

        
            Integration string  `json:"integration"`
            CompanyName string  `json:"company_name"`
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
    
    // State used by Cart
    type State struct {

        
            IsArchived bool  `json:"is_archived"`
            IsPublic bool  `json:"is_public"`
            IsDisplay bool  `json:"is_display"`
         
    }
    
    // Identifier used by Cart
    type Identifier struct {

        
            ItemID []float64  `json:"item_id"`
            ArticleID []string  `json:"article_id"`
            CategoryID []float64  `json:"category_id"`
            UserID []string  `json:"user_id"`
            CollectionID []string  `json:"collection_id"`
            StoreID []float64  `json:"store_id"`
            BrandID []float64  `json:"brand_id"`
            CompanyID []float64  `json:"company_id"`
            ExcludeBrandID []float64  `json:"exclude_brand_id"`
         
    }
    
    // DisplayMetaDict used by Cart
    type DisplayMetaDict struct {

        
            Title string  `json:"title"`
            Subtitle string  `json:"subtitle"`
         
    }
    
    // DisplayMeta used by Cart
    type DisplayMeta struct {

        
            Apply DisplayMetaDict  `json:"apply"`
            Remove DisplayMetaDict  `json:"remove"`
            Description string  `json:"description"`
            Title string  `json:"title"`
            Auto DisplayMetaDict  `json:"auto"`
            Subtitle string  `json:"subtitle"`
         
    }
    
    // CouponSchedule used by Cart
    type CouponSchedule struct {

        
            Cron string  `json:"cron"`
            Duration float64  `json:"duration"`
            Start string  `json:"start"`
            End string  `json:"end"`
            NextSchedule []map[string]interface{}  `json:"next_schedule"`
         
    }
    
    // RuleDefinition used by Cart
    type RuleDefinition struct {

        
            IsExact bool  `json:"is_exact"`
            Type string  `json:"type"`
            Scope []string  `json:"scope"`
            CurrencyCode string  `json:"currency_code"`
            ApplicableOn string  `json:"applicable_on"`
            CalculateOn string  `json:"calculate_on"`
            AutoApply bool  `json:"auto_apply"`
            ValueType string  `json:"value_type"`
         
    }
    
    // CouponAction used by Cart
    type CouponAction struct {

        
            TxnMode string  `json:"txn_mode"`
            ActionDate string  `json:"action_date"`
         
    }
    
    // Rule used by Cart
    type Rule struct {

        
            Key float64  `json:"key"`
            Max float64  `json:"max"`
            Value float64  `json:"value"`
            DiscountQty float64  `json:"discount_qty"`
            Min float64  `json:"min"`
         
    }
    
    // Validation used by Cart
    type Validation struct {

        
            UserRegisteredAfter string  `json:"user_registered_after"`
            AppID []string  `json:"app_id"`
            Anonymous bool  `json:"anonymous"`
         
    }
    
    // Validity used by Cart
    type Validity struct {

        
            Priority float64  `json:"priority"`
         
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
    
    // PostOrder used by Cart
    type PostOrder struct {

        
            CancellationAllowed bool  `json:"cancellation_allowed"`
            ReturnAllowed bool  `json:"return_allowed"`
         
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

        
            Types []string  `json:"types"`
            Networks []string  `json:"networks"`
            Uses PaymentAllowValue  `json:"uses"`
            Codes []string  `json:"codes"`
         
    }
    
    // BulkBundleRestriction used by Cart
    type BulkBundleRestriction struct {

        
            MultiStoreAllowed bool  `json:"multi_store_allowed"`
         
    }
    
    // Restrictions used by Cart
    type Restrictions struct {

        
            UserType string  `json:"user_type"`
            Uses UsesRestriction  `json:"uses"`
            PostOrder PostOrder  `json:"post_order"`
            PriceRange PriceRange  `json:"price_range"`
            Payments map[string]PaymentModes  `json:"payments"`
            BulkBundle BulkBundleRestriction  `json:"bulk_bundle"`
            CouponAllowed bool  `json:"coupon_allowed"`
            Platforms []string  `json:"platforms"`
            OrderingStores []float64  `json:"ordering_stores"`
            UserGroups []float64  `json:"user_groups"`
         
    }
    
    // CouponAuthor used by Cart
    type CouponAuthor struct {

        
            ModifiedBy string  `json:"modified_by"`
            CreatedBy string  `json:"created_by"`
         
    }
    
    // CouponAdd used by Cart
    type CouponAdd struct {

        
            Code string  `json:"code"`
            DateMeta CouponDateMeta  `json:"date_meta"`
            Ownership Ownership  `json:"ownership"`
            State State  `json:"state"`
            Identifiers Identifier  `json:"identifiers"`
            DisplayMeta DisplayMeta  `json:"display_meta"`
            Schedule CouponSchedule  `json:"_schedule"`
            RuleDefinition RuleDefinition  `json:"rule_definition"`
            Action CouponAction  `json:"action"`
            TypeSlug string  `json:"type_slug"`
            Rule []Rule  `json:"rule"`
            Validation Validation  `json:"validation"`
            Tags []string  `json:"tags"`
            Validity Validity  `json:"validity"`
            Restrictions Restrictions  `json:"restrictions"`
            Author CouponAuthor  `json:"author"`
         
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
            DateMeta CouponDateMeta  `json:"date_meta"`
            Ownership Ownership  `json:"ownership"`
            State State  `json:"state"`
            Identifiers Identifier  `json:"identifiers"`
            DisplayMeta DisplayMeta  `json:"display_meta"`
            Schedule CouponSchedule  `json:"_schedule"`
            RuleDefinition RuleDefinition  `json:"rule_definition"`
            Action CouponAction  `json:"action"`
            TypeSlug string  `json:"type_slug"`
            Rule []Rule  `json:"rule"`
            Validation Validation  `json:"validation"`
            Tags []string  `json:"tags"`
            Validity Validity  `json:"validity"`
            Restrictions Restrictions  `json:"restrictions"`
            Author CouponAuthor  `json:"author"`
         
    }
    
    // CouponPartialUpdate used by Cart
    type CouponPartialUpdate struct {

        
            Schedule CouponSchedule  `json:"schedule"`
            Archive bool  `json:"archive"`
         
    }
    
    // CompareObject used by Cart
    type CompareObject struct {

        
            GreaterThanEquals float64  `json:"greater_than_equals"`
            Equals float64  `json:"equals"`
            LessThan float64  `json:"less_than"`
            GreaterThan float64  `json:"greater_than"`
            LessThanEquals float64  `json:"less_than_equals"`
         
    }
    
    // ItemCriteria used by Cart
    type ItemCriteria struct {

        
            ItemExcludeBrand []float64  `json:"item_exclude_brand"`
            CartQuantity CompareObject  `json:"cart_quantity"`
            AvailableZones []string  `json:"available_zones"`
            ItemExcludeCompany []float64  `json:"item_exclude_company"`
            CartUniqueItemAmount CompareObject  `json:"cart_unique_item_amount"`
            ItemExcludeCategory []float64  `json:"item_exclude_category"`
            ItemSize []string  `json:"item_size"`
            ItemCompany []float64  `json:"item_company"`
            ItemID []float64  `json:"item_id"`
            ItemExcludeID []float64  `json:"item_exclude_id"`
            ItemCategory []float64  `json:"item_category"`
            ProductTags []string  `json:"product_tags"`
            ItemExcludeSku []string  `json:"item_exclude_sku"`
            CartTotal CompareObject  `json:"cart_total"`
            AllItems bool  `json:"all_items"`
            ItemBrand []float64  `json:"item_brand"`
            ItemSku []string  `json:"item_sku"`
            ItemExcludeStore []float64  `json:"item_exclude_store"`
            ItemStore []float64  `json:"item_store"`
            CartUniqueItemQuantity CompareObject  `json:"cart_unique_item_quantity"`
            BuyRules []string  `json:"buy_rules"`
         
    }
    
    // DiscountOffer used by Cart
    type DiscountOffer struct {

        
            Code string  `json:"code"`
            DiscountAmount float64  `json:"discount_amount"`
            ApportionDiscount bool  `json:"apportion_discount"`
            MinOfferQuantity float64  `json:"min_offer_quantity"`
            MaxDiscountAmount float64  `json:"max_discount_amount"`
            MaxOfferQuantity float64  `json:"max_offer_quantity"`
            PartialCanRet bool  `json:"partial_can_ret"`
            DiscountPercentage float64  `json:"discount_percentage"`
            MaxUsagePerTransaction float64  `json:"max_usage_per_transaction"`
            DiscountPrice float64  `json:"discount_price"`
         
    }
    
    // DiscountRule used by Cart
    type DiscountRule struct {

        
            ItemCriteria ItemCriteria  `json:"item_criteria"`
            BuyCondition string  `json:"buy_condition"`
            DiscountType string  `json:"discount_type"`
            Offer DiscountOffer  `json:"offer"`
         
    }
    
    // DisplayMeta1 used by Cart
    type DisplayMeta1 struct {

        
            Name string  `json:"name"`
            Description string  `json:"description"`
            OfferLabel string  `json:"offer_label"`
            OfferText string  `json:"offer_text"`
         
    }
    
    // PromotionAction used by Cart
    type PromotionAction struct {

        
            ActionDate string  `json:"action_date"`
            ActionType string  `json:"action_type"`
         
    }
    
    // PromotionSchedule used by Cart
    type PromotionSchedule struct {

        
            Cron string  `json:"cron"`
            Duration float64  `json:"duration"`
            Start string  `json:"start"`
            Published bool  `json:"published"`
            End string  `json:"end"`
            NextSchedule []map[string]interface{}  `json:"next_schedule"`
         
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
    
    // Ownership1 used by Cart
    type Ownership1 struct {

        
            PayableCategory string  `json:"payable_category"`
            PayableBy string  `json:"payable_by"`
         
    }
    
    // Visibility used by Cart
    type Visibility struct {

        
            CouponList bool  `json:"coupon_list"`
            Pdp bool  `json:"pdp"`
         
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
    
    // PostOrder1 used by Cart
    type PostOrder1 struct {

        
            CancellationAllowed bool  `json:"cancellation_allowed"`
            ReturnAllowed bool  `json:"return_allowed"`
         
    }
    
    // PaymentAllowValue1 used by Cart
    type PaymentAllowValue1 struct {

        
            Max float64  `json:"max"`
         
    }
    
    // PromotionPaymentModes used by Cart
    type PromotionPaymentModes struct {

        
            Uses PaymentAllowValue1  `json:"uses"`
            Type string  `json:"type"`
            Codes []string  `json:"codes"`
         
    }
    
    // UserRegistered used by Cart
    type UserRegistered struct {

        
            End string  `json:"end"`
            Start string  `json:"start"`
         
    }
    
    // Restrictions1 used by Cart
    type Restrictions1 struct {

        
            AnonymousUsers bool  `json:"anonymous_users"`
            Uses UsesRestriction1  `json:"uses"`
            PostOrder PostOrder1  `json:"post_order"`
            Payments []PromotionPaymentModes  `json:"payments"`
            Platforms []string  `json:"platforms"`
            UserGroups []float64  `json:"user_groups"`
            UserRegistered UserRegistered  `json:"user_registered"`
            UserID []string  `json:"user_id"`
            OrderQuantity float64  `json:"order_quantity"`
         
    }
    
    // PromotionListItem used by Cart
    type PromotionListItem struct {

        
            DiscountRules []DiscountRule  `json:"discount_rules"`
            ApplyPriority float64  `json:"apply_priority"`
            Currency string  `json:"currency"`
            PromotionType string  `json:"promotion_type"`
            DisplayMeta DisplayMeta1  `json:"display_meta"`
            PostOrderAction PromotionAction  `json:"post_order_action"`
            Schedule PromotionSchedule  `json:"_schedule"`
            CalculateOn string  `json:"calculate_on"`
            Code string  `json:"code"`
            DateMeta PromotionDateMeta  `json:"date_meta"`
            PromoGroup string  `json:"promo_group"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Mode string  `json:"mode"`
            Author PromotionAuthor  `json:"author"`
            ApplicationID string  `json:"application_id"`
            Ownership Ownership1  `json:"ownership"`
            Stackable bool  `json:"stackable"`
            ApplyExclusive string  `json:"apply_exclusive"`
            Visiblility Visibility  `json:"visiblility"`
            ApplyAllDiscount bool  `json:"apply_all_discount"`
            Restrictions Restrictions1  `json:"restrictions"`
            BuyRules map[string]ItemCriteria  `json:"buy_rules"`
         
    }
    
    // PromotionsResponse used by Cart
    type PromotionsResponse struct {

        
            Page Page  `json:"page"`
            Items PromotionListItem  `json:"items"`
         
    }
    
    // PromotionAdd used by Cart
    type PromotionAdd struct {

        
            DiscountRules []DiscountRule  `json:"discount_rules"`
            ApplyPriority float64  `json:"apply_priority"`
            Currency string  `json:"currency"`
            PromotionType string  `json:"promotion_type"`
            DisplayMeta DisplayMeta1  `json:"display_meta"`
            PostOrderAction PromotionAction  `json:"post_order_action"`
            Schedule PromotionSchedule  `json:"_schedule"`
            CalculateOn string  `json:"calculate_on"`
            Code string  `json:"code"`
            DateMeta PromotionDateMeta  `json:"date_meta"`
            PromoGroup string  `json:"promo_group"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Mode string  `json:"mode"`
            Author PromotionAuthor  `json:"author"`
            ApplicationID string  `json:"application_id"`
            Ownership Ownership1  `json:"ownership"`
            Stackable bool  `json:"stackable"`
            ApplyExclusive string  `json:"apply_exclusive"`
            Visiblility Visibility  `json:"visiblility"`
            ApplyAllDiscount bool  `json:"apply_all_discount"`
            Restrictions Restrictions1  `json:"restrictions"`
            BuyRules map[string]ItemCriteria  `json:"buy_rules"`
         
    }
    
    // PromotionUpdate used by Cart
    type PromotionUpdate struct {

        
            DiscountRules []DiscountRule  `json:"discount_rules"`
            ApplyPriority float64  `json:"apply_priority"`
            Currency string  `json:"currency"`
            PromotionType string  `json:"promotion_type"`
            DisplayMeta DisplayMeta1  `json:"display_meta"`
            PostOrderAction PromotionAction  `json:"post_order_action"`
            Schedule PromotionSchedule  `json:"_schedule"`
            CalculateOn string  `json:"calculate_on"`
            Code string  `json:"code"`
            DateMeta PromotionDateMeta  `json:"date_meta"`
            PromoGroup string  `json:"promo_group"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Mode string  `json:"mode"`
            Author PromotionAuthor  `json:"author"`
            ApplicationID string  `json:"application_id"`
            Ownership Ownership1  `json:"ownership"`
            Stackable bool  `json:"stackable"`
            ApplyExclusive string  `json:"apply_exclusive"`
            Visiblility Visibility  `json:"visiblility"`
            ApplyAllDiscount bool  `json:"apply_all_discount"`
            Restrictions Restrictions1  `json:"restrictions"`
            BuyRules map[string]ItemCriteria  `json:"buy_rules"`
         
    }
    
    // PromotionPartialUpdate used by Cart
    type PromotionPartialUpdate struct {

        
            Schedule PromotionSchedule  `json:"schedule"`
            Archive bool  `json:"archive"`
         
    }
    
    // ActivePromosResponse used by Cart
    type ActivePromosResponse struct {

        
            ModifiedOn string  `json:"modified_on"`
            EntityType string  `json:"entity_type"`
            Type string  `json:"type"`
            EntitySlug string  `json:"entity_slug"`
            CreatedOn string  `json:"created_on"`
            Description string  `json:"description"`
            Title string  `json:"title"`
            Example string  `json:"example"`
            IsHidden bool  `json:"is_hidden"`
            Subtitle string  `json:"subtitle"`
         
    }
    
    // CartItem used by Cart
    type CartItem struct {

        
            Size string  `json:"size"`
            Quantity float64  `json:"quantity"`
            ProductID string  `json:"product_id"`
         
    }
    
    // OpenapiCartDetailsRequest used by Cart
    type OpenapiCartDetailsRequest struct {

        
            CartItems CartItem  `json:"cart_items"`
         
    }
    
    // CartProductIdentifer used by Cart
    type CartProductIdentifer struct {

        
            Identifier string  `json:"identifier"`
         
    }
    
    // ProductPrice used by Cart
    type ProductPrice struct {

        
            Effective float64  `json:"effective"`
            Marked float64  `json:"marked"`
            AddOn float64  `json:"add_on"`
            CurrencyCode string  `json:"currency_code"`
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

        
            CurrencyCode string  `json:"currency_code"`
            Effective float64  `json:"effective"`
            Marked float64  `json:"marked"`
            CurrencySymbol string  `json:"currency_symbol"`
         
    }
    
    // ArticlePriceInfo used by Cart
    type ArticlePriceInfo struct {

        
            Base BasePrice  `json:"base"`
            Converted BasePrice  `json:"converted"`
         
    }
    
    // ProductArticle used by Cart
    type ProductArticle struct {

        
            UID string  `json:"uid"`
            Size string  `json:"size"`
            Type string  `json:"type"`
            Seller BaseInfo  `json:"seller"`
            Price ArticlePriceInfo  `json:"price"`
            Store BaseInfo  `json:"store"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            ProductGroupTags []string  `json:"product_group_tags"`
            Quantity float64  `json:"quantity"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
         
    }
    
    // DiscountRulesApp used by Cart
    type DiscountRulesApp struct {

        
            RawOffer map[string]interface{}  `json:"raw_offer"`
            ItemCriteria map[string]interface{}  `json:"item_criteria"`
            MatchedBuyRules []string  `json:"matched_buy_rules"`
            Offer map[string]interface{}  `json:"offer"`
         
    }
    
    // Ownership2 used by Cart
    type Ownership2 struct {

        
            PayableCategory string  `json:"payable_category"`
            PayableBy string  `json:"payable_by"`
         
    }
    
    // FreeGiftItem used by Cart
    type FreeGiftItem struct {

        
            ItemID float64  `json:"item_id"`
            ItemSlug string  `json:"item_slug"`
            ItemBrandName string  `json:"item_brand_name"`
            ItemPriceDetails map[string]interface{}  `json:"item_price_details"`
            ItemName string  `json:"item_name"`
            ItemImagesURL []string  `json:"item_images_url"`
         
    }
    
    // AppliedFreeArticles used by Cart
    type AppliedFreeArticles struct {

        
            ArticleID string  `json:"article_id"`
            FreeGiftItemDetails FreeGiftItem  `json:"free_gift_item_details"`
            Quantity float64  `json:"quantity"`
            ParentItemIdentifier string  `json:"parent_item_identifier"`
         
    }
    
    // BuyRules used by Cart
    type BuyRules struct {

        
            ItemCriteria map[string]interface{}  `json:"item_criteria"`
            CartConditions map[string]interface{}  `json:"cart_conditions"`
         
    }
    
    // AppliedPromotion used by Cart
    type AppliedPromotion struct {

        
            Amount float64  `json:"amount"`
            DiscountRules []DiscountRulesApp  `json:"discount_rules"`
            Ownership Ownership2  `json:"ownership"`
            AppliedFreeArticles []AppliedFreeArticles  `json:"applied_free_articles"`
            PromotionType string  `json:"promotion_type"`
            OfferText string  `json:"offer_text"`
            ArticleQuantity float64  `json:"article_quantity"`
            PromoID string  `json:"promo_id"`
            PromotionGroup string  `json:"promotion_group"`
            MrpPromotion bool  `json:"mrp_promotion"`
            PromotionName string  `json:"promotion_name"`
            BuyRules []BuyRules  `json:"buy_rules"`
         
    }
    
    // PromoMeta used by Cart
    type PromoMeta struct {

        
            Message string  `json:"message"`
         
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
    
    // NetQuantity used by Cart
    type NetQuantity struct {

        
            Unit string  `json:"unit"`
            Value string  `json:"value"`
         
    }
    
    // CartProduct used by Cart
    type CartProduct struct {

        
            Name string  `json:"name"`
            UID float64  `json:"uid"`
            Categories []CategoryInfo  `json:"categories"`
            Type string  `json:"type"`
            ItemCode string  `json:"item_code"`
            Images []ProductImage  `json:"images"`
            Action ProductAction  `json:"action"`
            Brand BaseInfo  `json:"brand"`
            NetQuantity NetQuantity  `json:"net_quantity"`
            Slug string  `json:"slug"`
         
    }
    
    // ProductAvailability used by Cart
    type ProductAvailability struct {

        
            OutOfStock bool  `json:"out_of_stock"`
            Sizes []string  `json:"sizes"`
            OtherStoreQuantity float64  `json:"other_store_quantity"`
            Deliverable bool  `json:"deliverable"`
            IsValid bool  `json:"is_valid"`
         
    }
    
    // CartProductInfo used by Cart
    type CartProductInfo struct {

        
            IsSet bool  `json:"is_set"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
            CouponMessage string  `json:"coupon_message"`
            Price ProductPriceInfo  `json:"price"`
            Discount string  `json:"discount"`
            Article ProductArticle  `json:"article"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            Message string  `json:"message"`
            Key string  `json:"key"`
            PromotionsApplied []AppliedPromotion  `json:"promotions_applied"`
            PromoMeta PromoMeta  `json:"promo_meta"`
            PricePerUnit ProductPriceInfo  `json:"price_per_unit"`
            Product CartProduct  `json:"product"`
            Quantity float64  `json:"quantity"`
            BulkOffer map[string]interface{}  `json:"bulk_offer"`
            Availability ProductAvailability  `json:"availability"`
         
    }
    
    // DisplayBreakup used by Cart
    type DisplayBreakup struct {

        
            Display string  `json:"display"`
            Message []string  `json:"message"`
            Key string  `json:"key"`
            CurrencyCode string  `json:"currency_code"`
            Value float64  `json:"value"`
            CurrencySymbol string  `json:"currency_symbol"`
         
    }
    
    // LoyaltyPoints used by Cart
    type LoyaltyPoints struct {

        
            Description string  `json:"description"`
            IsApplied bool  `json:"is_applied"`
            Total float64  `json:"total"`
            Applicable float64  `json:"applicable"`
         
    }
    
    // CouponBreakup used by Cart
    type CouponBreakup struct {

        
            Code string  `json:"code"`
            UID string  `json:"uid"`
            CouponValue float64  `json:"coupon_value"`
            SubTitle string  `json:"sub_title"`
            Type string  `json:"type"`
            MinimumCartValue float64  `json:"minimum_cart_value"`
            IsApplied bool  `json:"is_applied"`
            Message string  `json:"message"`
            Description string  `json:"description"`
            Title string  `json:"title"`
            MaxDiscountValue float64  `json:"max_discount_value"`
            Value float64  `json:"value"`
            CouponType string  `json:"coupon_type"`
         
    }
    
    // RawBreakup used by Cart
    type RawBreakup struct {

        
            Vog float64  `json:"vog"`
            Discount float64  `json:"discount"`
            FyndCash float64  `json:"fynd_cash"`
            DeliveryCharge float64  `json:"delivery_charge"`
            GstCharges float64  `json:"gst_charges"`
            CodCharge float64  `json:"cod_charge"`
            YouSaved float64  `json:"you_saved"`
            ConvenienceFee float64  `json:"convenience_fee"`
            Total float64  `json:"total"`
            MrpTotal float64  `json:"mrp_total"`
            Coupon float64  `json:"coupon"`
            Subtotal float64  `json:"subtotal"`
         
    }
    
    // CartBreakup used by Cart
    type CartBreakup struct {

        
            Display []DisplayBreakup  `json:"display"`
            LoyaltyPoints LoyaltyPoints  `json:"loyalty_points"`
            Coupon CouponBreakup  `json:"coupon"`
            Raw RawBreakup  `json:"raw"`
         
    }
    
    // OpenapiCartDetailsResponse used by Cart
    type OpenapiCartDetailsResponse struct {

        
            IsValid bool  `json:"is_valid"`
            Message string  `json:"message"`
            Items []CartProductInfo  `json:"items"`
            BreakupValues CartBreakup  `json:"breakup_values"`
         
    }
    
    // OpenApiErrorResponse used by Cart
    type OpenApiErrorResponse struct {

        
            Errors map[string]interface{}  `json:"errors"`
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // ShippingAddress used by Cart
    type ShippingAddress struct {

        
            Name string  `json:"name"`
            State string  `json:"state"`
            Area string  `json:"area"`
            Email string  `json:"email"`
            Address string  `json:"address"`
            City string  `json:"city"`
            AreaCode string  `json:"area_code"`
            AreaCodeSlug string  `json:"area_code_slug"`
            Meta map[string]interface{}  `json:"meta"`
            Phone float64  `json:"phone"`
            Country string  `json:"country"`
            CountryCode string  `json:"country_code"`
            Pincode float64  `json:"pincode"`
            AddressType string  `json:"address_type"`
            Landmark string  `json:"landmark"`
         
    }
    
    // OpenApiCartServiceabilityRequest used by Cart
    type OpenApiCartServiceabilityRequest struct {

        
            CartItems CartItem  `json:"cart_items"`
            ShippingAddress ShippingAddress  `json:"shipping_address"`
         
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

        
            Message string  `json:"message"`
            Items []CartProductInfo  `json:"items"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            IsValid bool  `json:"is_valid"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
         
    }
    
    // MultiTenderPaymentMeta used by Cart
    type MultiTenderPaymentMeta struct {

        
            PaymentID string  `json:"payment_id"`
            OrderID string  `json:"order_id"`
            PaymentGateway string  `json:"payment_gateway"`
            CurrentStatus string  `json:"current_status"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
         
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

        
            Values []string  `json:"values"`
            Key string  `json:"key"`
         
    }
    
    // CartItemMeta used by Cart
    type CartItemMeta struct {

        
            GroupID string  `json:"group_id"`
            PrimaryItem bool  `json:"primary_item"`
         
    }
    
    // OpenApiOrderItem used by Cart
    type OpenApiOrderItem struct {

        
            PaymentMethods []MultiTenderPaymentMethod  `json:"payment_methods"`
            ProductID float64  `json:"product_id"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            Files []OpenApiFiles  `json:"files"`
            Discount float64  `json:"discount"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            CashbackApplied float64  `json:"cashback_applied"`
            PriceMarked float64  `json:"price_marked"`
            EmployeeDiscount float64  `json:"employee_discount"`
            Meta CartItemMeta  `json:"meta"`
            AmountPaid float64  `json:"amount_paid"`
            CodCharges float64  `json:"cod_charges"`
            Quantity float64  `json:"quantity"`
            DeliveryCharges float64  `json:"delivery_charges"`
            Size string  `json:"size"`
            PriceEffective float64  `json:"price_effective"`
         
    }
    
    // OpenApiPlatformCheckoutReq used by Cart
    type OpenApiPlatformCheckoutReq struct {

        
            CouponValue float64  `json:"coupon_value"`
            CartItems []OpenApiOrderItem  `json:"cart_items"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            CartValue float64  `json:"cart_value"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            EmployeeDiscount map[string]interface{}  `json:"employee_discount"`
            CodCharges float64  `json:"cod_charges"`
            DeliveryCharges float64  `json:"delivery_charges"`
            Comment string  `json:"comment"`
            Gstin string  `json:"gstin"`
            Coupon string  `json:"coupon"`
            PaymentMethods []MultiTenderPaymentMethod  `json:"payment_methods"`
            OrderID string  `json:"order_id"`
            CouponCode string  `json:"coupon_code"`
            BillingAddress ShippingAddress  `json:"billing_address"`
            Files []OpenApiFiles  `json:"files"`
            CashbackApplied float64  `json:"cashback_applied"`
            CurrencyCode string  `json:"currency_code"`
            PaymentMode string  `json:"payment_mode"`
            ShippingAddress ShippingAddress  `json:"shipping_address"`
         
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
            Meta map[string]interface{}  `json:"meta"`
            CodCharges map[string]interface{}  `json:"cod_charges"`
            DeliveryCharges map[string]interface{}  `json:"delivery_charges"`
            Comment string  `json:"comment"`
            Coupon map[string]interface{}  `json:"coupon"`
            UserID string  `json:"user_id"`
            PaymentMethods []map[string]interface{}  `json:"payment_methods"`
            Shipments []map[string]interface{}  `json:"shipments"`
            MergeQty bool  `json:"merge_qty"`
            OrderID string  `json:"order_id"`
            Promotion map[string]interface{}  `json:"promotion"`
            Cashback map[string]interface{}  `json:"cashback"`
            Payments map[string]interface{}  `json:"payments"`
            ExpireAt string  `json:"expire_at"`
            PaymentMode string  `json:"payment_mode"`
            FcIndexMap []float64  `json:"fc_index_map"`
            BuyNow bool  `json:"buy_now"`
            AppID string  `json:"app_id"`
            UID float64  `json:"uid"`
            IsDefault bool  `json:"is_default"`
            BulkCouponDiscount float64  `json:"bulk_coupon_discount"`
            Discount float64  `json:"discount"`
            Articles []map[string]interface{}  `json:"articles"`
            CreatedOn string  `json:"created_on"`
            LastModified string  `json:"last_modified"`
            Gstin string  `json:"gstin"`
            IsArchive bool  `json:"is_archive"`
            IsActive bool  `json:"is_active"`
            FyndCredits map[string]interface{}  `json:"fynd_credits"`
            ID string  `json:"_id"`
            PickUpCustomerDetails map[string]interface{}  `json:"pick_up_customer_details"`
            CheckoutMode string  `json:"checkout_mode"`
         
    }
    
    // AbandonedCartResponse used by Cart
    type AbandonedCartResponse struct {

        
            Page Page  `json:"page"`
            Message string  `json:"message"`
            Items []AbandonedCart  `json:"items"`
            Success bool  `json:"success"`
            Result map[string]interface{}  `json:"result"`
         
    }
    
    // CartCurrency used by Cart
    type CartCurrency struct {

        
            Code string  `json:"code"`
            Symbol string  `json:"symbol"`
         
    }
    
    // PaymentSelectionLock used by Cart
    type PaymentSelectionLock struct {

        
            PaymentIdentifier string  `json:"payment_identifier"`
            Enabled bool  `json:"enabled"`
            DefaultOptions string  `json:"default_options"`
         
    }
    
    // CartDetailResponse used by Cart
    type CartDetailResponse struct {

        
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            ID string  `json:"id"`
            PanConfig map[string]interface{}  `json:"pan_config"`
            Message string  `json:"message"`
            Items []CartProductInfo  `json:"items"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            IsValid bool  `json:"is_valid"`
            Currency CartCurrency  `json:"currency"`
            Comment string  `json:"comment"`
            LastModified string  `json:"last_modified"`
            CouponText string  `json:"coupon_text"`
            PanNo string  `json:"pan_no"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            Gstin string  `json:"gstin"`
            CheckoutMode string  `json:"checkout_mode"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            BuyNow bool  `json:"buy_now"`
            RestrictCheckout bool  `json:"restrict_checkout"`
         
    }
    
    // AddProductCart used by Cart
    type AddProductCart struct {

        
            ItemID float64  `json:"item_id"`
            ArticleAssignment map[string]interface{}  `json:"article_assignment"`
            ArticleID string  `json:"article_id"`
            Display string  `json:"display"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            Pos bool  `json:"pos"`
            SellerID float64  `json:"seller_id"`
            StoreID float64  `json:"store_id"`
            ProductGroupTags []string  `json:"product_group_tags"`
            Quantity float64  `json:"quantity"`
            ItemSize string  `json:"item_size"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
         
    }
    
    // AddCartRequest used by Cart
    type AddCartRequest struct {

        
            NewCart bool  `json:"new_cart"`
            Items []AddProductCart  `json:"items"`
         
    }
    
    // AddCartDetailResponse used by Cart
    type AddCartDetailResponse struct {

        
            Cart CartDetailResponse  `json:"cart"`
            Partial bool  `json:"partial"`
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // UpdateProductCart used by Cart
    type UpdateProductCart struct {

        
            ItemID float64  `json:"item_id"`
            ArticleID string  `json:"article_id"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            ItemIndex float64  `json:"item_index"`
            Quantity float64  `json:"quantity"`
            ItemSize string  `json:"item_size"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
         
    }
    
    // UpdateCartRequest used by Cart
    type UpdateCartRequest struct {

        
            Operation string  `json:"operation"`
            Items []UpdateProductCart  `json:"items"`
         
    }
    
    // UpdateCartDetailResponse used by Cart
    type UpdateCartDetailResponse struct {

        
            Cart CartDetailResponse  `json:"cart"`
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // GetShareCartLinkRequest used by Cart
    type GetShareCartLinkRequest struct {

        
            Meta map[string]interface{}  `json:"meta"`
            ID string  `json:"id"`
         
    }
    
    // GetShareCartLinkResponse used by Cart
    type GetShareCartLinkResponse struct {

        
            Token string  `json:"token"`
            ShareURL string  `json:"share_url"`
         
    }
    
    // SharedCartDetails used by Cart
    type SharedCartDetails struct {

        
            CreatedOn string  `json:"created_on"`
            Meta map[string]interface{}  `json:"meta"`
            Source map[string]interface{}  `json:"source"`
            Token string  `json:"token"`
            User map[string]interface{}  `json:"user"`
         
    }
    
    // SharedCart used by Cart
    type SharedCart struct {

        
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            SharedCartDetails SharedCartDetails  `json:"shared_cart_details"`
            Message string  `json:"message"`
            Currency CartCurrency  `json:"currency"`
            IsValid bool  `json:"is_valid"`
            Comment string  `json:"comment"`
            Items []CartProductInfo  `json:"items"`
            CartID float64  `json:"cart_id"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            BuyNow bool  `json:"buy_now"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            UID string  `json:"uid"`
            ID string  `json:"id"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            LastModified string  `json:"last_modified"`
            Gstin string  `json:"gstin"`
            CouponText string  `json:"coupon_text"`
            CheckoutMode string  `json:"checkout_mode"`
         
    }
    
    // SharedCartResponse used by Cart
    type SharedCartResponse struct {

        
            Cart SharedCart  `json:"cart"`
            Error string  `json:"error"`
         
    }
    
    // CartList used by Cart
    type CartList struct {

        
            ItemCounts string  `json:"item_counts"`
            CartValue string  `json:"cart_value"`
            CreatedOn string  `json:"created_on"`
            CartID string  `json:"cart_id"`
            PickUpCustomerDetails string  `json:"pick_up_customer_details"`
            UserID string  `json:"user_id"`
         
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

        
            UID string  `json:"uid"`
            ModifiedOn string  `json:"modified_on"`
            FirstName string  `json:"first_name"`
            Mobile string  `json:"mobile"`
            ExternalID string  `json:"external_id"`
            CreatedAt string  `json:"created_at"`
            ID string  `json:"_id"`
            Gender string  `json:"gender"`
            LastName string  `json:"last_name"`
         
    }
    
    // UserCartMappingResponse used by Cart
    type UserCartMappingResponse struct {

        
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            Message string  `json:"message"`
            Currency CartCurrency  `json:"currency"`
            IsValid bool  `json:"is_valid"`
            Comment string  `json:"comment"`
            Items []CartProductInfo  `json:"items"`
            User UserInfo  `json:"user"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            BuyNow bool  `json:"buy_now"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            ID string  `json:"id"`
            PanConfig map[string]interface{}  `json:"pan_config"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            LastModified string  `json:"last_modified"`
            Gstin string  `json:"gstin"`
            CouponText string  `json:"coupon_text"`
            PanNo string  `json:"pan_no"`
            CheckoutMode string  `json:"checkout_mode"`
         
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
    
    // PageCoupon used by Cart
    type PageCoupon struct {

        
            Current float64  `json:"current"`
            HasPrevious bool  `json:"has_previous"`
            Total float64  `json:"total"`
            TotalItemCount float64  `json:"total_item_count"`
            HasNext bool  `json:"has_next"`
         
    }
    
    // Coupon used by Cart
    type Coupon struct {

        
            CouponValue float64  `json:"coupon_value"`
            SubTitle string  `json:"sub_title"`
            ExpiresOn string  `json:"expires_on"`
            CouponCode string  `json:"coupon_code"`
            MinimumCartValue float64  `json:"minimum_cart_value"`
            Message string  `json:"message"`
            IsApplied bool  `json:"is_applied"`
            Description string  `json:"description"`
            Title string  `json:"title"`
            MaxDiscountValue float64  `json:"max_discount_value"`
            IsApplicable bool  `json:"is_applicable"`
            CouponType string  `json:"coupon_type"`
         
    }
    
    // GetCouponResponse used by Cart
    type GetCouponResponse struct {

        
            Page PageCoupon  `json:"page"`
            AvailableCouponList []Coupon  `json:"available_coupon_list"`
         
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

        
            Name string  `json:"name"`
            GoogleMapPoint map[string]interface{}  `json:"google_map_point"`
            GeoLocation GeoLocation  `json:"geo_location"`
            AreaCodeSlug string  `json:"area_code_slug"`
            Meta map[string]interface{}  `json:"meta"`
            Tags []string  `json:"tags"`
            AddressType string  `json:"address_type"`
            UserID string  `json:"user_id"`
            Landmark string  `json:"landmark"`
            CreatedByUserID string  `json:"created_by_user_id"`
            City string  `json:"city"`
            CartID string  `json:"cart_id"`
            Phone string  `json:"phone"`
            CountryCode string  `json:"country_code"`
            State string  `json:"state"`
            ID string  `json:"id"`
            Email string  `json:"email"`
            AreaCode string  `json:"area_code"`
            Country string  `json:"country"`
            IsActive bool  `json:"is_active"`
            IsDefaultAddress bool  `json:"is_default_address"`
            Area string  `json:"area"`
            Address string  `json:"address"`
            CheckoutMode string  `json:"checkout_mode"`
         
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

        
            IsUpdated bool  `json:"is_updated"`
            IsDefaultAddress bool  `json:"is_default_address"`
            ID string  `json:"id"`
            Success bool  `json:"success"`
         
    }
    
    // DeleteAddressResponse used by Cart
    type DeleteAddressResponse struct {

        
            ID string  `json:"id"`
            IsDeleted bool  `json:"is_deleted"`
         
    }
    
    // PlatformSelectCartAddressRequest used by Cart
    type PlatformSelectCartAddressRequest struct {

        
            BillingAddressID string  `json:"billing_address_id"`
            CheckoutMode string  `json:"checkout_mode"`
            ID string  `json:"id"`
            CartID string  `json:"cart_id"`
         
    }
    
    // ShipmentResponse used by Cart
    type ShipmentResponse struct {

        
            FulfillmentType string  `json:"fulfillment_type"`
            Shipments float64  `json:"shipments"`
            FulfillmentID float64  `json:"fulfillment_id"`
            Promise ShipmentPromise  `json:"promise"`
            OrderType string  `json:"order_type"`
            Items []CartProductInfo  `json:"items"`
            DpOptions map[string]interface{}  `json:"dp_options"`
            BoxType string  `json:"box_type"`
            DpID string  `json:"dp_id"`
            ShipmentType string  `json:"shipment_type"`
         
    }
    
    // CartShipmentsResponse used by Cart
    type CartShipmentsResponse struct {

        
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            UID string  `json:"uid"`
            ID string  `json:"id"`
            Message string  `json:"message"`
            Currency CartCurrency  `json:"currency"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            IsValid bool  `json:"is_valid"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            LastModified string  `json:"last_modified"`
            Comment string  `json:"comment"`
            Gstin string  `json:"gstin"`
            Error bool  `json:"error"`
            Shipments []ShipmentResponse  `json:"shipments"`
            CartID float64  `json:"cart_id"`
            CouponText string  `json:"coupon_text"`
            CheckoutMode string  `json:"checkout_mode"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            BuyNow bool  `json:"buy_now"`
            RestrictCheckout bool  `json:"restrict_checkout"`
         
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
            PanNo string  `json:"pan_no"`
            Comment string  `json:"comment"`
            PickUpCustomerDetails map[string]interface{}  `json:"pick_up_customer_details"`
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

        
            FirstName string  `json:"first_name"`
            ID string  `json:"_id"`
            LastName string  `json:"last_name"`
            User string  `json:"user"`
            EmployeeCode string  `json:"employee_code"`
         
    }
    
    // Files used by Cart
    type Files struct {

        
            Values []string  `json:"values"`
            Key string  `json:"key"`
         
    }
    
    // PlatformCartCheckoutDetailRequest used by Cart
    type PlatformCartCheckoutDetailRequest struct {

        
            PickAtStoreUID float64  `json:"pick_at_store_uid"`
            Meta map[string]interface{}  `json:"meta"`
            CallbackURL string  `json:"callback_url"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            UserID string  `json:"user_id"`
            PaymentAutoConfirm bool  `json:"payment_auto_confirm"`
            AddressID string  `json:"address_id"`
            OrderType string  `json:"order_type"`
            Pos bool  `json:"pos"`
            BillingAddressID string  `json:"billing_address_id"`
            PaymentMode string  `json:"payment_mode"`
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
            Staff StaffCheckout  `json:"staff"`
            ID string  `json:"id"`
            PaymentParams map[string]interface{}  `json:"payment_params"`
            OrderingStore float64  `json:"ordering_store"`
            Aggregator string  `json:"aggregator"`
            DeviceID string  `json:"device_id"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
            Files []Files  `json:"files"`
            PaymentIdentifier string  `json:"payment_identifier"`
            CheckoutMode string  `json:"checkout_mode"`
            MerchantCode string  `json:"merchant_code"`
         
    }
    
    // CheckCart used by Cart
    type CheckCart struct {

        
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            UserType string  `json:"user_type"`
            Message string  `json:"message"`
            Currency CartCurrency  `json:"currency"`
            IsValid bool  `json:"is_valid"`
            CodCharges float64  `json:"cod_charges"`
            DeliveryCharges float64  `json:"delivery_charges"`
            Comment string  `json:"comment"`
            ErrorMessage string  `json:"error_message"`
            OrderID string  `json:"order_id"`
            Items []CartProductInfo  `json:"items"`
            CartID float64  `json:"cart_id"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            BuyNow bool  `json:"buy_now"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            UID string  `json:"uid"`
            ID string  `json:"id"`
            StoreEmps []map[string]interface{}  `json:"store_emps"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            Success bool  `json:"success"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            LastModified string  `json:"last_modified"`
            Gstin string  `json:"gstin"`
            DeliveryChargeOrderValue float64  `json:"delivery_charge_order_value"`
            CodMessage string  `json:"cod_message"`
            CodAvailable bool  `json:"cod_available"`
            CouponText string  `json:"coupon_text"`
            StoreCode string  `json:"store_code"`
            CheckoutMode string  `json:"checkout_mode"`
         
    }
    
    // CartCheckoutResponse used by Cart
    type CartCheckoutResponse struct {

        
            AppInterceptURL string  `json:"app_intercept_url"`
            PaymentConfirmURL string  `json:"payment_confirm_url"`
            OrderID string  `json:"order_id"`
            Message string  `json:"message"`
            Success bool  `json:"success"`
            Data map[string]interface{}  `json:"data"`
            CallbackURL string  `json:"callback_url"`
            Cart CheckCart  `json:"cart"`
         
    }
    
    // CartDeliveryModesResponse used by Cart
    type CartDeliveryModesResponse struct {

        
            AvailableModes []string  `json:"available_modes"`
            PickupStores []float64  `json:"pickup_stores"`
         
    }
    
    // PickupStoreDetail used by Cart
    type PickupStoreDetail struct {

        
            Name string  `json:"name"`
            State string  `json:"state"`
            Area string  `json:"area"`
            UID float64  `json:"uid"`
            ID float64  `json:"id"`
            Email string  `json:"email"`
            Address string  `json:"address"`
            City string  `json:"city"`
            AreaCode string  `json:"area_code"`
            AreaCodeSlug string  `json:"area_code_slug"`
            Phone string  `json:"phone"`
            Country string  `json:"country"`
            StoreCode string  `json:"store_code"`
            Pincode float64  `json:"pincode"`
            AddressType string  `json:"address_type"`
            Landmark string  `json:"landmark"`
         
    }
    
    // StoreDetailsResponse used by Cart
    type StoreDetailsResponse struct {

        
            Items []PickupStoreDetail  `json:"items"`
         
    }
    
    // UpdateCartPaymentRequest used by Cart
    type UpdateCartPaymentRequest struct {

        
            AddressID string  `json:"address_id"`
            ID string  `json:"id"`
            AggregatorName string  `json:"aggregator_name"`
            PaymentIdentifier string  `json:"payment_identifier"`
            PaymentMode string  `json:"payment_mode"`
            MerchantCode string  `json:"merchant_code"`
         
    }
    
    // CouponValidity used by Cart
    type CouponValidity struct {

        
            Code string  `json:"code"`
            DisplayMessageEn string  `json:"display_message_en"`
            Discount float64  `json:"discount"`
            Title string  `json:"title"`
            Valid bool  `json:"valid"`
         
    }
    
    // PaymentCouponValidate used by Cart
    type PaymentCouponValidate struct {

        
            Message string  `json:"message"`
            CouponValidity CouponValidity  `json:"coupon_validity"`
            Success bool  `json:"success"`
         
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
    

    
    // ServiceabilityErrorResponse used by Logistic
    type ServiceabilityErrorResponse struct {

        
            Message string  `json:"message"`
            Value string  `json:"value"`
            Type string  `json:"type"`
         
    }
    
    // ApplicationServiceabilityConfig used by Logistic
    type ApplicationServiceabilityConfig struct {

        
            ChannelID string  `json:"channel_id"`
            ServiceabilityType string  `json:"serviceability_type"`
            ChannelType string  `json:"channel_type"`
         
    }
    
    // ApplicationServiceabilityConfigResponse used by Logistic
    type ApplicationServiceabilityConfigResponse struct {

        
            Error ServiceabilityErrorResponse  `json:"error"`
            Data ApplicationServiceabilityConfig  `json:"data"`
            Success bool  `json:"success"`
         
    }
    
    // EntityRegionView_Request used by Logistic
    type EntityRegionView_Request struct {

        
            SubType []string  `json:"sub_type"`
            ParentID []string  `json:"parent_id"`
         
    }
    
    // EntityRegionView_Error used by Logistic
    type EntityRegionView_Error struct {

        
            Message string  `json:"message"`
            Value string  `json:"value"`
            Type string  `json:"type"`
         
    }
    
    // EntityRegionView_page used by Logistic
    type EntityRegionView_page struct {

        
            Current float64  `json:"current"`
            Size float64  `json:"size"`
            ItemTotal float64  `json:"item_total"`
            HasNext bool  `json:"has_next"`
            Type string  `json:"type"`
         
    }
    
    // EntityRegionView_Items used by Logistic
    type EntityRegionView_Items struct {

        
            Name string  `json:"name"`
            SubType string  `json:"sub_type"`
            UID string  `json:"uid"`
         
    }
    
    // EntityRegionView_Response used by Logistic
    type EntityRegionView_Response struct {

        
            Error EntityRegionView_Error  `json:"error"`
            Page EntityRegionView_page  `json:"page"`
            Data []EntityRegionView_Items  `json:"data"`
            Success bool  `json:"success"`
         
    }
    
    // ZoneDataItem used by Logistic
    type ZoneDataItem struct {

        
            Current float64  `json:"current"`
            Size float64  `json:"size"`
            ItemTotal float64  `json:"item_total"`
            Type string  `json:"type"`
            HasNext bool  `json:"has_next"`
         
    }
    
    // ListViewSummary used by Logistic
    type ListViewSummary struct {

        
            TotalZones float64  `json:"total_zones"`
            TotalActiveZones float64  `json:"total_active_zones"`
            TotalPincodesServed float64  `json:"total_pincodes_served"`
         
    }
    
    // ListViewProduct used by Logistic
    type ListViewProduct struct {

        
            Count float64  `json:"count"`
            Type string  `json:"type"`
         
    }
    
    // ListViewChannels used by Logistic
    type ListViewChannels struct {

        
            ChannelID string  `json:"channel_id"`
            ChannelType string  `json:"channel_type"`
         
    }
    
    // ListViewItems used by Logistic
    type ListViewItems struct {

        
            StoresCount float64  `json:"stores_count"`
            PincodesCount float64  `json:"pincodes_count"`
            IsActive bool  `json:"is_active"`
            Product ListViewProduct  `json:"product"`
            CompanyID float64  `json:"company_id"`
            Channels ListViewChannels  `json:"channels"`
            Slug string  `json:"slug"`
            Name string  `json:"name"`
            ZoneID string  `json:"zone_id"`
         
    }
    
    // ListViewResponse used by Logistic
    type ListViewResponse struct {

        
            Page []ZoneDataItem  `json:"page"`
            Summary []ListViewSummary  `json:"summary"`
            Items []ListViewItems  `json:"items"`
         
    }
    
    // CompanyStoreView_PageItems used by Logistic
    type CompanyStoreView_PageItems struct {

        
            Current float64  `json:"current"`
            Size float64  `json:"size"`
            ItemTotal float64  `json:"item_total"`
            HasNext bool  `json:"has_next"`
            Type string  `json:"type"`
         
    }
    
    // CompanyStoreView_Response used by Logistic
    type CompanyStoreView_Response struct {

        
            Page []CompanyStoreView_PageItems  `json:"page"`
            Items []map[string]interface{}  `json:"items"`
         
    }
    
    // GetZoneDataViewChannels used by Logistic
    type GetZoneDataViewChannels struct {

        
            ChannelID string  `json:"channel_id"`
            ChannelType string  `json:"channel_type"`
         
    }
    
    // ZoneProductTypes used by Logistic
    type ZoneProductTypes struct {

        
            Tags []string  `json:"tags"`
            Type string  `json:"type"`
         
    }
    
    // ZoneMappingType used by Logistic
    type ZoneMappingType struct {

        
            Country string  `json:"country"`
            Pincode []string  `json:"pincode"`
            State []string  `json:"state"`
         
    }
    
    // GetZoneDataViewItems used by Logistic
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
    
    // GetSingleZoneDataViewResponse used by Logistic
    type GetSingleZoneDataViewResponse struct {

        
            Data GetZoneDataViewItems  `json:"data"`
         
    }
    
    // UpdateZoneData used by Logistic
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
    
    // ZoneUpdateRequest used by Logistic
    type ZoneUpdateRequest struct {

        
            Identifier string  `json:"identifier"`
            Data UpdateZoneData  `json:"data"`
         
    }
    
    // ZoneSuccessResponse used by Logistic
    type ZoneSuccessResponse struct {

        
            StatusCode float64  `json:"status_code"`
            Success bool  `json:"success"`
         
    }
    
    // CreateZoneData used by Logistic
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
    
    // ZoneRequest used by Logistic
    type ZoneRequest struct {

        
            Identifier string  `json:"identifier"`
            Data CreateZoneData  `json:"data"`
         
    }
    
    // ZoneResponse used by Logistic
    type ZoneResponse struct {

        
            StatusCode float64  `json:"status_code"`
            ZoneID string  `json:"zone_id"`
            Success bool  `json:"success"`
         
    }
    
    // GetZoneFromPincodeViewRequest used by Logistic
    type GetZoneFromPincodeViewRequest struct {

        
            Pincode string  `json:"pincode"`
            Country string  `json:"country"`
         
    }
    
    // GetZoneFromPincodeViewResponse used by Logistic
    type GetZoneFromPincodeViewResponse struct {

        
            Zones []string  `json:"zones"`
            ServiceabilityType string  `json:"serviceability_type"`
         
    }
    
    // ServiceabilityPageResponse used by Logistic
    type ServiceabilityPageResponse struct {

        
            Current float64  `json:"current"`
            Size float64  `json:"size"`
            ItemTotal float64  `json:"item_total"`
            HasNext bool  `json:"has_next"`
            Type string  `json:"type"`
         
    }
    
    // DocumentsResponse used by Logistic
    type DocumentsResponse struct {

        
            Value string  `json:"value"`
            LegalName string  `json:"legal_name"`
            Verified bool  `json:"verified"`
            Type string  `json:"type"`
         
    }
    
    // AddressResponse used by Logistic
    type AddressResponse struct {

        
            Landmark string  `json:"landmark"`
            Pincode float64  `json:"pincode"`
            City string  `json:"city"`
            Country string  `json:"country"`
            Latitude float64  `json:"latitude"`
            State string  `json:"state"`
            Address1 string  `json:"address1"`
            Address2 string  `json:"address2"`
            Longitude float64  `json:"longitude"`
         
    }
    
    // ModifiedByResponse used by Logistic
    type ModifiedByResponse struct {

        
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
         
    }
    
    // CreatedByResponse used by Logistic
    type CreatedByResponse struct {

        
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
         
    }
    
    // WarningsResponse used by Logistic
    type WarningsResponse struct {

        
            StoreAddress string  `json:"store_address"`
         
    }
    
    // OpeningClosing used by Logistic
    type OpeningClosing struct {

        
            Hour float64  `json:"hour"`
            Minute float64  `json:"minute"`
         
    }
    
    // TimmingResponse used by Logistic
    type TimmingResponse struct {

        
            Closing OpeningClosing  `json:"closing"`
            Opening OpeningClosing  `json:"opening"`
            Weekday string  `json:"weekday"`
            Open bool  `json:"open"`
         
    }
    
    // ProductReturnConfigResponse used by Logistic
    type ProductReturnConfigResponse struct {

        
            OnSameStore bool  `json:"on_same_store"`
         
    }
    
    // MobileNo used by Logistic
    type MobileNo struct {

        
            CountryCode float64  `json:"country_code"`
            Number string  `json:"number"`
         
    }
    
    // ManagerResponse used by Logistic
    type ManagerResponse struct {

        
            MobileNo MobileNo  `json:"mobile_no"`
            Name string  `json:"name"`
            Email string  `json:"email"`
         
    }
    
    // IntegrationTypeResponse used by Logistic
    type IntegrationTypeResponse struct {

        
            Inventory string  `json:"inventory"`
            Order string  `json:"order"`
         
    }
    
    // ContactNumberResponse used by Logistic
    type ContactNumberResponse struct {

        
            CountryCode float64  `json:"country_code"`
            Number string  `json:"number"`
         
    }
    
    // Dp used by Logistic
    type Dp struct {

        
            FmPriority float64  `json:"fm_priority"`
            PaymentMode string  `json:"payment_mode"`
            ExternalAccountID string  `json:"external_account_id"`
            RvpPriority float64  `json:"rvp_priority"`
            AssignDpFromSb bool  `json:"assign_dp_from_sb"`
            LmPriority float64  `json:"lm_priority"`
            TransportMode string  `json:"transport_mode"`
            Operations []string  `json:"operations"`
            InternalAccountID string  `json:"internal_account_id"`
            AreaCode float64  `json:"area_code"`
         
    }
    
    // LogisticsResponse used by Logistic
    type LogisticsResponse struct {

        
            Override bool  `json:"override"`
            Dp Dp  `json:"dp"`
         
    }
    
    // EwayBillResponse used by Logistic
    type EwayBillResponse struct {

        
            Enabled bool  `json:"enabled"`
         
    }
    
    // EinvoiceResponse used by Logistic
    type EinvoiceResponse struct {

        
            Enabled bool  `json:"enabled"`
         
    }
    
    // GstCredentialsResponse used by Logistic
    type GstCredentialsResponse struct {

        
            EWaybill EwayBillResponse  `json:"e_waybill"`
            EInvoice EinvoiceResponse  `json:"e_invoice"`
         
    }
    
    // ItemResponse used by Logistic
    type ItemResponse struct {

        
            Documents []DocumentsResponse  `json:"documents"`
            Company float64  `json:"company"`
            Address AddressResponse  `json:"address"`
            DisplayName string  `json:"display_name"`
            ModifiedBy ModifiedByResponse  `json:"modified_by"`
            ModifiedOn string  `json:"modified_on"`
            Cls string  `json:"_cls"`
            CreatedBy CreatedByResponse  `json:"created_by"`
            Warnings WarningsResponse  `json:"warnings"`
            Timing []TimmingResponse  `json:"timing"`
            Code string  `json:"code"`
            CompanyID float64  `json:"company_id"`
            CreatedOn string  `json:"created_on"`
            SubType string  `json:"sub_type"`
            ProductReturnConfig ProductReturnConfigResponse  `json:"product_return_config"`
            Manager ManagerResponse  `json:"manager"`
            IntegrationType IntegrationTypeResponse  `json:"integration_type"`
            ContactNumbers []ContactNumberResponse  `json:"contact_numbers"`
            UID float64  `json:"uid"`
            Stage string  `json:"stage"`
            VerifiedBy ModifiedByResponse  `json:"verified_by"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            VerifiedOn string  `json:"verified_on"`
            Logistics LogisticsResponse  `json:"logistics"`
            StoreType string  `json:"store_type"`
            GstCredentials GstCredentialsResponse  `json:"gst_credentials"`
            Name string  `json:"name"`
            NotificationEmails []string  `json:"notification_emails"`
         
    }
    
    // GetStoresViewResponse used by Logistic
    type GetStoresViewResponse struct {

        
            Page ServiceabilityPageResponse  `json:"page"`
            Items []ItemResponse  `json:"items"`
         
    }
    
