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
    

    
    // PaymentGatewayConfigResponse used by Payment
    type PaymentGatewayConfigResponse struct {

        
            DisplayFields []string  `json:"display_fields"`
            Aggregators []map[string]interface{}  `json:"aggregators"`
            Success bool  `json:"success"`
            ExcludedFields []string  `json:"excluded_fields"`
            Created bool  `json:"created"`
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

        
            IsActive bool  `json:"is_active"`
            MerchantSalt string  `json:"merchant_salt"`
            Secret string  `json:"secret"`
            ConfigType string  `json:"config_type"`
            Key string  `json:"key"`
         
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

        
            Logos PaymentModeLogo  `json:"logos"`
            PackageName string  `json:"package_name"`
            DisplayName string  `json:"display_name"`
            Code string  `json:"code"`
         
    }
    
    // IntentAppErrorList used by Payment
    type IntentAppErrorList struct {

        
            PackageName string  `json:"package_name"`
            Code string  `json:"code"`
         
    }
    
    // PaymentModeList used by Payment
    type PaymentModeList struct {

        
            RemainingLimit float64  `json:"remaining_limit"`
            CardToken string  `json:"card_token"`
            DisplayName string  `json:"display_name"`
            LogoURL PaymentModeLogo  `json:"logo_url"`
            CardType string  `json:"card_type"`
            FyndVpa string  `json:"fynd_vpa"`
            ExpYear float64  `json:"exp_year"`
            CardBrandImage string  `json:"card_brand_image"`
            CardNumber string  `json:"card_number"`
            CardID string  `json:"card_id"`
            CardIsin string  `json:"card_isin"`
            Timeout float64  `json:"timeout"`
            CardName string  `json:"card_name"`
            IntentApp []IntentApp  `json:"intent_app"`
            CardFingerprint string  `json:"card_fingerprint"`
            CompliantWithTokenisationGuidelines bool  `json:"compliant_with_tokenisation_guidelines"`
            AggregatorName string  `json:"aggregator_name"`
            MerchantCode string  `json:"merchant_code"`
            Code string  `json:"code"`
            Expired bool  `json:"expired"`
            IntentAppErrorList []string  `json:"intent_app_error_list"`
            CardBrand string  `json:"card_brand"`
            CodLimit float64  `json:"cod_limit"`
            DisplayPriority float64  `json:"display_priority"`
            ExpMonth float64  `json:"exp_month"`
            Nickname string  `json:"nickname"`
            IntentFlow bool  `json:"intent_flow"`
            CardReference string  `json:"card_reference"`
            CodLimitPerOrder float64  `json:"cod_limit_per_order"`
            IntentAppErrorDictList []IntentAppErrorList  `json:"intent_app_error_dict_list"`
            CardIssuer string  `json:"card_issuer"`
            Name string  `json:"name"`
            RetryCount float64  `json:"retry_count"`
         
    }
    
    // RootPaymentMode used by Payment
    type RootPaymentMode struct {

        
            IsPayByCardPl bool  `json:"is_pay_by_card_pl"`
            AddCardEnabled bool  `json:"add_card_enabled"`
            Name string  `json:"name"`
            AnonymousEnable bool  `json:"anonymous_enable"`
            DisplayName string  `json:"display_name"`
            List []PaymentModeList  `json:"list"`
            AggregatorName string  `json:"aggregator_name"`
            DisplayPriority float64  `json:"display_priority"`
            SaveCard bool  `json:"save_card"`
         
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
    
    // PayoutAggregator used by Payment
    type PayoutAggregator struct {

        
            AggregatorID float64  `json:"aggregator_id"`
            PayoutDetailsID float64  `json:"payout_details_id"`
            AggregatorFundID float64  `json:"aggregator_fund_id"`
         
    }
    
    // PayoutMoreAttributes used by Payment
    type PayoutMoreAttributes struct {

        
            State string  `json:"state"`
            AccountNo string  `json:"account_no"`
            AccountType string  `json:"account_type"`
            BankName string  `json:"bank_name"`
            AccountHolder string  `json:"account_holder"`
            IfscCode string  `json:"ifsc_code"`
            City string  `json:"city"`
            Country string  `json:"country"`
            BranchName string  `json:"branch_name"`
         
    }
    
    // PayoutCustomer used by Payment
    type PayoutCustomer struct {

        
            Name string  `json:"name"`
            UniqueExternalID string  `json:"unique_external_id"`
            ID float64  `json:"id"`
            Mobile string  `json:"mobile"`
         
    }
    
    // Payout used by Payment
    type Payout struct {

        
            IsDefault bool  `json:"is_default"`
            IsActive bool  `json:"is_active"`
            PayoutsAggregators []PayoutAggregator  `json:"payouts_aggregators"`
            TransferType string  `json:"transfer_type"`
            MoreAttributes PayoutMoreAttributes  `json:"more_attributes"`
            UniqueTransferNo string  `json:"unique_transfer_no"`
            Customers PayoutCustomer  `json:"customers"`
         
    }
    
    // PayoutsResponse used by Payment
    type PayoutsResponse struct {

        
            Items []Payout  `json:"items"`
            Success bool  `json:"success"`
         
    }
    
    // PayoutBankDetails used by Payment
    type PayoutBankDetails struct {

        
            AccountNo string  `json:"account_no"`
            City string  `json:"city"`
            Pincode float64  `json:"pincode"`
            AccountType string  `json:"account_type"`
            AccountHolder string  `json:"account_holder"`
            BankName string  `json:"bank_name"`
            IfscCode string  `json:"ifsc_code"`
            State string  `json:"state"`
            Country string  `json:"country"`
            BranchName string  `json:"branch_name"`
         
    }
    
    // PayoutRequest used by Payment
    type PayoutRequest struct {

        
            IsActive bool  `json:"is_active"`
            TransferType string  `json:"transfer_type"`
            Users map[string]interface{}  `json:"users"`
            UniqueExternalID string  `json:"unique_external_id"`
            Aggregator string  `json:"aggregator"`
            BankDetails PayoutBankDetails  `json:"bank_details"`
         
    }
    
    // PayoutResponse used by Payment
    type PayoutResponse struct {

        
            IsActive bool  `json:"is_active"`
            Success bool  `json:"success"`
            TransferType string  `json:"transfer_type"`
            UniqueTransferNo string  `json:"unique_transfer_no"`
            Users map[string]interface{}  `json:"users"`
            Payouts map[string]interface{}  `json:"payouts"`
            PaymentStatus string  `json:"payment_status"`
            Aggregator string  `json:"aggregator"`
            Created bool  `json:"created"`
            BankDetails map[string]interface{}  `json:"bank_details"`
         
    }
    
    // UpdatePayoutResponse used by Payment
    type UpdatePayoutResponse struct {

        
            Success bool  `json:"success"`
            IsDefault bool  `json:"is_default"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // UpdatePayoutRequest used by Payment
    type UpdatePayoutRequest struct {

        
            UniqueExternalID string  `json:"unique_external_id"`
            IsDefault bool  `json:"is_default"`
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

        
            Data map[string]interface{}  `json:"data"`
            Success bool  `json:"success"`
            IsVerifiedFlag bool  `json:"is_verified_flag"`
            Message string  `json:"message"`
         
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
            BankName string  `json:"bank_name"`
            AccountHolder string  `json:"account_holder"`
            IfscCode string  `json:"ifsc_code"`
            BranchName string  `json:"branch_name"`
         
    }
    
    // AddBeneficiaryDetailsOTPRequest used by Payment
    type AddBeneficiaryDetailsOTPRequest struct {

        
            OrderID string  `json:"order_id"`
            Details BankDetailsForOTP  `json:"details"`
         
    }
    
    // IfscCodeResponse used by Payment
    type IfscCodeResponse struct {

        
            BranchName string  `json:"branch_name"`
            Success bool  `json:"success"`
            BankName string  `json:"bank_name"`
         
    }
    
    // OrderBeneficiaryDetails used by Payment
    type OrderBeneficiaryDetails struct {

        
            IsActive bool  `json:"is_active"`
            Address string  `json:"address"`
            BeneficiaryID string  `json:"beneficiary_id"`
            AccountHolder string  `json:"account_holder"`
            BankName string  `json:"bank_name"`
            DisplayName string  `json:"display_name"`
            ID float64  `json:"id"`
            Email string  `json:"email"`
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
            AccountNo string  `json:"account_no"`
            Title string  `json:"title"`
            IfscCode string  `json:"ifsc_code"`
            Subtitle string  `json:"subtitle"`
            BranchName string  `json:"branch_name"`
            Comment string  `json:"comment"`
            DelightsUserName string  `json:"delights_user_name"`
            Mobile string  `json:"mobile"`
            TransferMode string  `json:"transfer_mode"`
         
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
            PaymentID string  `json:"payment_id"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            PaymentGateway string  `json:"payment_gateway"`
         
    }
    
    // MultiTenderPaymentMethod used by Payment
    type MultiTenderPaymentMethod struct {

        
            Meta MultiTenderPaymentMeta  `json:"meta"`
            Name string  `json:"name"`
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
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // PlatformPaymentOptions used by Payment
    type PlatformPaymentOptions struct {

        
            Methods map[string]interface{}  `json:"methods"`
            CodCharges float64  `json:"cod_charges"`
            PaymentSelectionLock map[string]interface{}  `json:"payment_selection_lock"`
            Enabled bool  `json:"enabled"`
            AnonymousCod bool  `json:"anonymous_cod"`
            ModeOfPayment string  `json:"mode_of_payment"`
            Source string  `json:"source"`
            CodAmountLimit float64  `json:"cod_amount_limit"`
            CallbackURL map[string]interface{}  `json:"callback_url"`
         
    }
    
    // PlatfromPaymentConfig used by Payment
    type PlatfromPaymentConfig struct {

        
            Data PlatformPaymentOptions  `json:"data"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // UpdatePlatformPaymentConfig used by Payment
    type UpdatePlatformPaymentConfig struct {

        
            Methods map[string]interface{}  `json:"methods"`
            CodCharges float64  `json:"cod_charges"`
            PaymentSelectionLock map[string]interface{}  `json:"payment_selection_lock"`
            AnonymousCod bool  `json:"anonymous_cod"`
            CodAmountLimit float64  `json:"cod_amount_limit"`
         
    }
    
    // CODdata used by Payment
    type CODdata struct {

        
            UserID string  `json:"user_id"`
            Limit float64  `json:"limit"`
            IsActive bool  `json:"is_active"`
            RemainingLimit float64  `json:"remaining_limit"`
            Usages float64  `json:"usages"`
         
    }
    
    // GetUserCODLimitResponse used by Payment
    type GetUserCODLimitResponse struct {

        
            UserCodData CODdata  `json:"user_cod_data"`
            Success bool  `json:"success"`
         
    }
    
    // SetCODForUserRequest used by Payment
    type SetCODForUserRequest struct {

        
            Mobileno string  `json:"mobileno"`
            MerchantUserID string  `json:"merchant_user_id"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // SetCODOptionResponse used by Payment
    type SetCODOptionResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
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

        
            AggregatorID float64  `json:"aggregator_id"`
            DeviceTag string  `json:"device_tag"`
            EdcModel string  `json:"edc_model"`
            StoreID float64  `json:"store_id"`
            EdcDeviceSerialNo string  `json:"edc_device_serial_no"`
            TerminalSerialNo string  `json:"terminal_serial_no"`
         
    }
    
    // EdcDevice used by Payment
    type EdcDevice struct {

        
            AggregatorID float64  `json:"aggregator_id"`
            DeviceTag string  `json:"device_tag"`
            MerchantStorePosCode string  `json:"merchant_store_pos_code"`
            ApplicationID string  `json:"application_id"`
            IsActive bool  `json:"is_active"`
            EdcModel string  `json:"edc_model"`
            StoreID float64  `json:"store_id"`
            TerminalUniqueIdentifier string  `json:"terminal_unique_identifier"`
            AggregatorName string  `json:"aggregator_name"`
            EdcDeviceSerialNo string  `json:"edc_device_serial_no"`
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

        
            AggregatorID float64  `json:"aggregator_id"`
            DeviceTag string  `json:"device_tag"`
            IsActive bool  `json:"is_active"`
            MerchantStorePosCode string  `json:"merchant_store_pos_code"`
            EdcModel string  `json:"edc_model"`
            StoreID float64  `json:"store_id"`
            EdcDeviceSerialNo string  `json:"edc_device_serial_no"`
         
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

        
            CustomerID string  `json:"customer_id"`
            Contact string  `json:"contact"`
            Amount float64  `json:"amount"`
            OrderID string  `json:"order_id"`
            MerchantOrderID string  `json:"merchant_order_id"`
            Email string  `json:"email"`
            Method string  `json:"method"`
            DeviceID string  `json:"device_id"`
            Timeout float64  `json:"timeout"`
            Aggregator string  `json:"aggregator"`
            Currency string  `json:"currency"`
            Vpa string  `json:"vpa"`
            RazorpayPaymentID string  `json:"razorpay_payment_id"`
         
    }
    
    // PaymentInitializationResponse used by Payment
    type PaymentInitializationResponse struct {

        
            Contact string  `json:"contact"`
            Amount float64  `json:"amount"`
            Success bool  `json:"success"`
            DeviceID string  `json:"device_id"`
            Aggregator string  `json:"aggregator"`
            Vpa string  `json:"vpa"`
            VirtualID string  `json:"virtual_id"`
            CustomerID string  `json:"customer_id"`
            PollingURL string  `json:"polling_url"`
            Email string  `json:"email"`
            Timeout float64  `json:"timeout"`
            PaymentID string  `json:"payment_id"`
            OrderID string  `json:"order_id"`
            BqrImage string  `json:"bqr_image"`
            MerchantOrderID string  `json:"merchant_order_id"`
            Method string  `json:"method"`
            UpiPollURL string  `json:"upi_poll_url"`
            Status string  `json:"status"`
            RazorpayPaymentID string  `json:"razorpay_payment_id"`
            AggregatorOrderID string  `json:"aggregator_order_id"`
            Currency string  `json:"currency"`
         
    }
    
    // PaymentStatusUpdateRequest used by Payment
    type PaymentStatusUpdateRequest struct {

        
            CustomerID string  `json:"customer_id"`
            Contact string  `json:"contact"`
            Amount float64  `json:"amount"`
            OrderID string  `json:"order_id"`
            MerchantOrderID string  `json:"merchant_order_id"`
            MerchantTransactionID string  `json:"merchant_transaction_id"`
            Email string  `json:"email"`
            Method string  `json:"method"`
            DeviceID string  `json:"device_id"`
            PaymentID string  `json:"payment_id"`
            Aggregator string  `json:"aggregator"`
            Currency string  `json:"currency"`
            Status string  `json:"status"`
            Vpa string  `json:"vpa"`
            RazorpayPaymentID string  `json:"razorpay_payment_id"`
         
    }
    
    // PaymentStatusUpdateResponse used by Payment
    type PaymentStatusUpdateResponse struct {

        
            RedirectURL string  `json:"redirect_url"`
            Retry bool  `json:"retry"`
            AggregatorName string  `json:"aggregator_name"`
            Success bool  `json:"success"`
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

        
            CreatedOn string  `json:"created_on"`
            RefundedBy string  `json:"refunded_by"`
            ID string  `json:"id"`
            AggregatorPaymentObject map[string]interface{}  `json:"aggregator_payment_object"`
            CollectedBy string  `json:"collected_by"`
            ApplicationID string  `json:"application_id"`
            CurrentStatus string  `json:"current_status"`
            CompanyID string  `json:"company_id"`
            ModifiedOn string  `json:"modified_on"`
            PaymentMode string  `json:"payment_mode"`
            PaymentID string  `json:"payment_id"`
            PaymentModeIdentifier string  `json:"payment_mode_identifier"`
            RefundObject map[string]interface{}  `json:"refund_object"`
            PaymentGateway string  `json:"payment_gateway"`
            Currency string  `json:"currency"`
            AmountInPaisa string  `json:"amount_in_paisa"`
            UserObject map[string]interface{}  `json:"user_object"`
            AllStatus []string  `json:"all_status"`
         
    }
    
    // PaymentStatusObject used by Payment
    type PaymentStatusObject struct {

        
            PaymentObjectList []PaymentObjectListSerializer  `json:"payment_object_list"`
            MerchantOrderID string  `json:"merchant_order_id"`
         
    }
    
    // PaymentStatusBulkHandlerResponse used by Payment
    type PaymentStatusBulkHandlerResponse struct {

        
            Count float64  `json:"count"`
            Error string  `json:"error"`
            Data []PaymentStatusObject  `json:"data"`
            Success string  `json:"success"`
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
    
    // ValidateCustomerRequest used by Payment
    type ValidateCustomerRequest struct {

        
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
            TransactionAmountInPaise float64  `json:"transaction_amount_in_paise"`
            Payload string  `json:"payload"`
            OrderItems []map[string]interface{}  `json:"order_items"`
            MerchantParams map[string]interface{}  `json:"merchant_params"`
            Aggregator string  `json:"aggregator"`
            PhoneNumber string  `json:"phone_number"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
         
    }
    
    // ValidateCustomerResponse used by Payment
    type ValidateCustomerResponse struct {

        
            Error map[string]interface{}  `json:"error"`
            Data map[string]interface{}  `json:"data"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // GetPaymentLinkResponse used by Payment
    type GetPaymentLinkResponse struct {

        
            PaymentLinkCurrentStatus string  `json:"payment_link_current_status"`
            Message string  `json:"message"`
            Amount float64  `json:"amount"`
            PaymentLinkURL string  `json:"payment_link_url"`
            PollingTimeout float64  `json:"polling_timeout"`
            Success bool  `json:"success"`
            StatusCode float64  `json:"status_code"`
            ExternalOrderID string  `json:"external_order_id"`
            MerchantName string  `json:"merchant_name"`
         
    }
    
    // ErrorDescription used by Payment
    type ErrorDescription struct {

        
            Amount float64  `json:"amount"`
            InvalidID bool  `json:"invalid_id"`
            Cancelled bool  `json:"cancelled"`
            MerchantOrderID string  `json:"merchant_order_id"`
            Expired bool  `json:"expired"`
            Msg string  `json:"msg"`
            PaymentTransactionID string  `json:"payment_transaction_id"`
            MerchantName string  `json:"merchant_name"`
         
    }
    
    // ErrorResponse used by Payment
    type ErrorResponse struct {

        
            Error ErrorDescription  `json:"error"`
            StatusCode float64  `json:"status_code"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
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

        
            Meta CreatePaymentLinkMeta  `json:"meta"`
            Amount float64  `json:"amount"`
            Description string  `json:"description"`
            Email string  `json:"email"`
            ExternalOrderID string  `json:"external_order_id"`
            MobileNumber string  `json:"mobile_number"`
         
    }
    
    // CreatePaymentLinkResponse used by Payment
    type CreatePaymentLinkResponse struct {

        
            Message string  `json:"message"`
            PollingTimeout float64  `json:"polling_timeout"`
            PaymentLinkURL string  `json:"payment_link_url"`
            Success bool  `json:"success"`
            StatusCode float64  `json:"status_code"`
            PaymentLinkID string  `json:"payment_link_id"`
         
    }
    
    // PollingPaymentLinkResponse used by Payment
    type PollingPaymentLinkResponse struct {

        
            Message string  `json:"message"`
            RedirectURL string  `json:"redirect_url"`
            Amount float64  `json:"amount"`
            OrderID string  `json:"order_id"`
            AggregatorName string  `json:"aggregator_name"`
            Success bool  `json:"success"`
            HttpStatus float64  `json:"http_status"`
            StatusCode float64  `json:"status_code"`
            Status string  `json:"status"`
            PaymentLinkID string  `json:"payment_link_id"`
         
    }
    
    // CancelOrResendPaymentLinkRequest used by Payment
    type CancelOrResendPaymentLinkRequest struct {

        
            PaymentLinkID string  `json:"payment_link_id"`
         
    }
    
    // ResendPaymentLinkResponse used by Payment
    type ResendPaymentLinkResponse struct {

        
            StatusCode float64  `json:"status_code"`
            PollingTimeout float64  `json:"polling_timeout"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // CancelPaymentLinkResponse used by Payment
    type CancelPaymentLinkResponse struct {

        
            StatusCode float64  `json:"status_code"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    

    
    // UserDataInfo used by Order
    type UserDataInfo struct {

        
            Email string  `json:"email"`
            Name string  `json:"name"`
            Gender string  `json:"gender"`
            IsAnonymousUser bool  `json:"is_anonymous_user"`
            AvisUserID string  `json:"avis_user_id"`
            UID float64  `json:"uid"`
            FirstName string  `json:"first_name"`
            Mobile string  `json:"mobile"`
            LastName string  `json:"last_name"`
         
    }
    
    // ShipmentTimeStamp used by Order
    type ShipmentTimeStamp struct {

        
            TMax string  `json:"t_max"`
            TMin string  `json:"t_min"`
         
    }
    
    // LockData used by Order
    type LockData struct {

        
            LockMessage string  `json:"lock_message"`
            Locked bool  `json:"locked"`
            Mto bool  `json:"mto"`
         
    }
    
    // Formatted used by Order
    type Formatted struct {

        
            FMax string  `json:"f_max"`
            FMin string  `json:"f_min"`
         
    }
    
    // ShipmentTags used by Order
    type ShipmentTags struct {

        
            Slug string  `json:"slug"`
            DisplayText string  `json:"display_text"`
            EntityType string  `json:"entity_type"`
         
    }
    
    // ShipmentItemMeta used by Order
    type ShipmentItemMeta struct {

        
            PackagingName string  `json:"packaging_name"`
            ShippingZone string  `json:"shipping_zone"`
            Timestamp ShipmentTimeStamp  `json:"timestamp"`
            Tags []map[string]interface{}  `json:"tags"`
            ShipmentWeight float64  `json:"shipment_weight"`
            OrderType string  `json:"order_type"`
            DpOptions map[string]interface{}  `json:"dp_options"`
            Weight float64  `json:"weight"`
            ShipmentChargeableWeight float64  `json:"shipment_chargeable_weight"`
            DpSortKey string  `json:"dp_sort_key"`
            LockData LockData  `json:"lock_data"`
            AssignDpFromSb bool  `json:"assign_dp_from_sb"`
            ParentDpID string  `json:"parent_dp_id"`
            DebugInfo map[string]interface{}  `json:"debug_info"`
            External map[string]interface{}  `json:"external"`
            BagWeight map[string]interface{}  `json:"bag_weight"`
            ActivityComment string  `json:"activity_comment"`
            EwaybillInfo map[string]interface{}  `json:"ewaybill_info"`
            Formatted Formatted  `json:"formatted"`
            ShipmentVolumetricWeight float64  `json:"shipment_volumetric_weight"`
            SameStoreAvailable bool  `json:"same_store_available"`
            FulfilmentPriorityText string  `json:"fulfilment_priority_text"`
            ShipmentTags []ShipmentTags  `json:"shipment_tags"`
            IsInternational bool  `json:"is_international"`
            PdfMedia []map[string]interface{}  `json:"pdf_media"`
            Sla float64  `json:"sla"`
            StoreInvoiceUpdatedDate string  `json:"store_invoice_updated_date"`
            AutoTriggerDpAssignmentAcf bool  `json:"auto_trigger_dp_assignment_acf"`
            ExistingDpList []string  `json:"existing_dp_list"`
         
    }
    
    // PlatformDeliveryAddress used by Order
    type PlatformDeliveryAddress struct {

        
            Latitude float64  `json:"latitude"`
            Longitude float64  `json:"longitude"`
            Area string  `json:"area"`
            State string  `json:"state"`
            Email string  `json:"email"`
            Pincode string  `json:"pincode"`
            Version string  `json:"version"`
            UpdatedAt string  `json:"updated_at"`
            CreatedAt string  `json:"created_at"`
            City string  `json:"city"`
            Address1 string  `json:"address1"`
            Address2 string  `json:"address2"`
            ContactPerson string  `json:"contact_person"`
            AddressCategory string  `json:"address_category"`
            Phone string  `json:"phone"`
            AddressType string  `json:"address_type"`
            Country string  `json:"country"`
            Landmark string  `json:"landmark"`
         
    }
    
    // Prices used by Order
    type Prices struct {

        
            FyndCredits float64  `json:"fynd_credits"`
            TransferPrice float64  `json:"transfer_price"`
            RefundCredit float64  `json:"refund_credit"`
            RefundAmount float64  `json:"refund_amount"`
            CodCharges float64  `json:"cod_charges"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            GiftPrice float64  `json:"gift_price"`
            DeliveryCharge float64  `json:"delivery_charge"`
            Cashback float64  `json:"cashback"`
            CashbackApplied float64  `json:"cashback_applied"`
            PriceMarked float64  `json:"price_marked"`
            PriceEffective float64  `json:"price_effective"`
            ValueOfGood float64  `json:"value_of_good"`
            PromotionEffectiveDiscount float64  `json:"promotion_effective_discount"`
            PmPriceSplit float64  `json:"pm_price_split"`
            Discount float64  `json:"discount"`
            AmountPaidRoundoff float64  `json:"amount_paid_roundoff"`
            CouponValue float64  `json:"coupon_value"`
            AmountPaid float64  `json:"amount_paid"`
         
    }
    
    // ShipmentStatus used by Order
    type ShipmentStatus struct {

        
            BagList []string  `json:"bag_list"`
            CreatedAt string  `json:"created_at"`
            Title string  `json:"title"`
            ShipmentID string  `json:"shipment_id"`
            CurrentShipmentStatus string  `json:"current_shipment_status"`
            Meta map[string]interface{}  `json:"meta"`
            Status string  `json:"status"`
            StatusCreatedAt string  `json:"status_created_at"`
            ShipmentStatusID float64  `json:"shipment_status_id"`
         
    }
    
    // BagStateMapper used by Order
    type BagStateMapper struct {

        
            DisplayName string  `json:"display_name"`
            BsID float64  `json:"bs_id"`
            AppStateName string  `json:"app_state_name"`
            Name string  `json:"name"`
            AppDisplayName string  `json:"app_display_name"`
            AppFacing bool  `json:"app_facing"`
            IsActive bool  `json:"is_active"`
            NotifyCustomer bool  `json:"notify_customer"`
            StateType string  `json:"state_type"`
            JourneyType string  `json:"journey_type"`
         
    }
    
    // BagStatusHistory used by Order
    type BagStatusHistory struct {

        
            UpdatedAt string  `json:"updated_at"`
            DisplayName string  `json:"display_name"`
            BagStateMapper BagStateMapper  `json:"bag_state_mapper"`
            DeliveryPartnerID float64  `json:"delivery_partner_id"`
            CreatedAt string  `json:"created_at"`
            AppDisplayName string  `json:"app_display_name"`
            BagID float64  `json:"bag_id"`
            Forward bool  `json:"forward"`
            KafkaSync bool  `json:"kafka_sync"`
            Reasons []map[string]interface{}  `json:"reasons"`
            ShipmentID string  `json:"shipment_id"`
            StateType string  `json:"state_type"`
            Status string  `json:"status"`
            StoreID float64  `json:"store_id"`
            StateID float64  `json:"state_id"`
            BshID float64  `json:"bsh_id"`
            DeliveryAwbNumber string  `json:"delivery_awb_number"`
         
    }
    
    // GSTDetailsData used by Order
    type GSTDetailsData struct {

        
            GstFee float64  `json:"gst_fee"`
            GstinCode string  `json:"gstin_code"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            CgstTaxPercentage float64  `json:"cgst_tax_percentage"`
            IsDefaultHsnCode bool  `json:"is_default_hsn_code"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            SgstTaxPercentage float64  `json:"sgst_tax_percentage"`
            ValueOfGood float64  `json:"value_of_good"`
            IgstGstFee string  `json:"igst_gst_fee"`
            SgstGstFee string  `json:"sgst_gst_fee"`
            CgstGstFee string  `json:"cgst_gst_fee"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            GstTag string  `json:"gst_tag"`
            HsnCode string  `json:"hsn_code"`
            HsnCodeID string  `json:"hsn_code_id"`
            IgstTaxPercentage float64  `json:"igst_tax_percentage"`
         
    }
    
    // BagReturnableCancelableStatus used by Order
    type BagReturnableCancelableStatus struct {

        
            CanBeCancelled bool  `json:"can_be_cancelled"`
            IsCustomerReturnAllowed bool  `json:"is_customer_return_allowed"`
            IsActive bool  `json:"is_active"`
            EnableTracking bool  `json:"enable_tracking"`
            IsReturnable bool  `json:"is_returnable"`
         
    }
    
    // Identifier used by Order
    type Identifier struct {

        
            Ean string  `json:"ean"`
         
    }
    
    // FinancialBreakup used by Order
    type FinancialBreakup struct {

        
            AddedToFyndCash bool  `json:"added_to_fynd_cash"`
            FyndCredits float64  `json:"fynd_credits"`
            TransferPrice float64  `json:"transfer_price"`
            RefundCredit float64  `json:"refund_credit"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            HsnCode string  `json:"hsn_code"`
            CodCharges float64  `json:"cod_charges"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            ItemName string  `json:"item_name"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            DeliveryCharge float64  `json:"delivery_charge"`
            Cashback float64  `json:"cashback"`
            Identifiers Identifier  `json:"identifiers"`
            CashbackApplied float64  `json:"cashback_applied"`
            PriceEffective float64  `json:"price_effective"`
            ValueOfGood float64  `json:"value_of_good"`
            PriceMarked float64  `json:"price_marked"`
            PromotionEffectiveDiscount float64  `json:"promotion_effective_discount"`
            GstTag string  `json:"gst_tag"`
            Discount float64  `json:"discount"`
            GstFee float64  `json:"gst_fee"`
            TotalUnits float64  `json:"total_units"`
            Size string  `json:"size"`
            AmountPaidRoundoff float64  `json:"amount_paid_roundoff"`
            CouponValue float64  `json:"coupon_value"`
            AmountPaid float64  `json:"amount_paid"`
         
    }
    
    // ReturnConfig used by Order
    type ReturnConfig struct {

        
            Time float64  `json:"time"`
            Unit string  `json:"unit"`
            Returnable bool  `json:"returnable"`
         
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

        
            Unit string  `json:"unit"`
            Shipping float64  `json:"shipping"`
            IsDefault bool  `json:"is_default"`
         
    }
    
    // Article used by Order
    type Article struct {

        
            ASet map[string]interface{}  `json:"a_set"`
            Size string  `json:"size"`
            SellerIdentifier string  `json:"seller_identifier"`
            ChildDetails map[string]interface{}  `json:"child_details"`
            Identifiers Identifier  `json:"identifiers"`
            ReturnConfig ReturnConfig  `json:"return_config"`
            IsSet bool  `json:"is_set"`
            ID string  `json:"_id"`
            UID string  `json:"uid"`
            EspModified bool  `json:"esp_modified"`
            Dimensions Dimensions  `json:"dimensions"`
            Currency map[string]interface{}  `json:"currency"`
            RawMeta string  `json:"raw_meta"`
            Weight Weight  `json:"weight"`
            Code string  `json:"code"`
         
    }
    
    // AffiliateMeta used by Order
    type AffiliateMeta struct {

        
            DueDate string  `json:"due_date"`
            BoxType string  `json:"box_type"`
            IsPriority bool  `json:"is_priority"`
            MarketplaceInvoiceID string  `json:"marketplace_invoice_id"`
            CouponCode string  `json:"coupon_code"`
            Quantity float64  `json:"quantity"`
            ChannelShipmentID string  `json:"channel_shipment_id"`
            SizeLevelTotalQty float64  `json:"size_level_total_qty"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            OrderItemID string  `json:"order_item_id"`
            ChannelOrderID string  `json:"channel_order_id"`
            EmployeeDiscount float64  `json:"employee_discount"`
         
    }
    
    // AffiliateBagDetails used by Order
    type AffiliateBagDetails struct {

        
            AffiliateMeta AffiliateMeta  `json:"affiliate_meta"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            EmployeeDiscount float64  `json:"employee_discount"`
         
    }
    
    // ShipmentListingBrand used by Order
    type ShipmentListingBrand struct {

        
            Name string  `json:"name"`
            LogoBase64 string  `json:"logo_base64"`
            CreatedOn string  `json:"created_on"`
            Logo string  `json:"logo"`
         
    }
    
    // PlatformArticleAttributes used by Order
    type PlatformArticleAttributes struct {

        
            Currency string  `json:"currency"`
         
    }
    
    // PlatformItem used by Order
    type PlatformItem struct {

        
            Color string  `json:"color"`
            LastUpdatedAt string  `json:"last_updated_at"`
            Code string  `json:"code"`
            Meta map[string]interface{}  `json:"meta"`
            L3CategoryName string  `json:"l3_category_name"`
            L1Category []string  `json:"l1_category"`
            CanReturn bool  `json:"can_return"`
            Name string  `json:"name"`
            Image []string  `json:"image"`
            DepartmentID float64  `json:"department_id"`
            CanCancel bool  `json:"can_cancel"`
            BrandID float64  `json:"brand_id"`
            L3Category float64  `json:"l3_category"`
            Attributes PlatformArticleAttributes  `json:"attributes"`
            Brand string  `json:"brand"`
            BranchURL string  `json:"branch_url"`
            Images []string  `json:"images"`
            Size string  `json:"size"`
            SlugKey string  `json:"slug_key"`
            L2Category []string  `json:"l2_category"`
            ID float64  `json:"id"`
         
    }
    
    // Dates used by Order
    type Dates struct {

        
            DeliveryDate string  `json:"delivery_date"`
            OrderCreated string  `json:"order_created"`
         
    }
    
    // BagUnit used by Order
    type BagUnit struct {

        
            DisplayName string  `json:"display_name"`
            CurrentOperationalStatus BagStatusHistory  `json:"current_operational_status"`
            Meta map[string]interface{}  `json:"meta"`
            BagStatus []BagStatusHistory  `json:"bag_status"`
            CanReturn bool  `json:"can_return"`
            BagID float64  `json:"bag_id"`
            BagType string  `json:"bag_type"`
            Prices Prices  `json:"prices"`
            CanCancel bool  `json:"can_cancel"`
            Gst GSTDetailsData  `json:"gst"`
            CurrentStatus BagStatusHistory  `json:"current_status"`
            Status BagReturnableCancelableStatus  `json:"status"`
            FinancialBreakup []FinancialBreakup  `json:"financial_breakup"`
            Article Article  `json:"article"`
            AffiliateBagDetails AffiliateBagDetails  `json:"affiliate_bag_details"`
            EntityType string  `json:"entity_type"`
            BagExpiryDate string  `json:"bag_expiry_date"`
            Reasons []map[string]interface{}  `json:"reasons"`
            Brand ShipmentListingBrand  `json:"brand"`
            Item PlatformItem  `json:"item"`
            ProductQuantity float64  `json:"product_quantity"`
            Size string  `json:"size"`
            Dates Dates  `json:"dates"`
         
    }
    
    // ShipmentItemFulFillingStore used by Order
    type ShipmentItemFulFillingStore struct {

        
            StoreEmail string  `json:"store_email"`
            Pincode string  `json:"pincode"`
            State string  `json:"state"`
            Name string  `json:"name"`
            City string  `json:"city"`
            FulfillingStoreID float64  `json:"fulfilling_store_id"`
            LocationType string  `json:"location_type"`
            Meta map[string]interface{}  `json:"meta"`
            Phone string  `json:"phone"`
            BrandStoreTags string  `json:"brand_store_tags"`
            Code string  `json:"code"`
            Address string  `json:"address"`
         
    }
    
    // ShipmentListingChannel used by Order
    type ShipmentListingChannel struct {

        
            Name string  `json:"name"`
            ChannelShipmentID string  `json:"channel_shipment_id"`
            Logo string  `json:"logo"`
            IsAffiliate bool  `json:"is_affiliate"`
         
    }
    
    // ShipmentItem used by Order
    type ShipmentItem struct {

        
            DisplayName string  `json:"display_name"`
            User UserDataInfo  `json:"user"`
            Meta ShipmentItemMeta  `json:"meta"`
            PaymentMode string  `json:"payment_mode"`
            DeliveryAddress PlatformDeliveryAddress  `json:"delivery_address"`
            OrderID string  `json:"order_id"`
            StatusCreatedAt string  `json:"status_created_at"`
            LockStatus bool  `json:"lock_status"`
            OrderDate string  `json:"order_date"`
            Prices Prices  `json:"prices"`
            PaymentMethods map[string]interface{}  `json:"payment_methods"`
            CustomerNote string  `json:"customer_note"`
            ShipmentCreatedAt string  `json:"shipment_created_at"`
            OrderingChannnel string  `json:"ordering_channnel"`
            ShipmentStatus ShipmentStatus  `json:"shipment_status"`
            Bags []BagUnit  `json:"bags"`
            FulfillingStore ShipmentItemFulFillingStore  `json:"fulfilling_store"`
            PreviousShipmentID string  `json:"previous_shipment_id"`
            Channel ShipmentListingChannel  `json:"channel"`
            InvoiceID string  `json:"invoice_id"`
            EstimatedSlaTime string  `json:"estimated_sla_time"`
            TotalBags float64  `json:"total_bags"`
            ShipmentID string  `json:"shipment_id"`
            CanProcess bool  `json:"can_process"`
         
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
            Success bool  `json:"success"`
            TotalCount float64  `json:"total_count"`
            Lane string  `json:"lane"`
            Items []ShipmentItem  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // Error used by Order
    type Error struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // UserDetailsData used by Order
    type UserDetailsData struct {

        
            Area string  `json:"area"`
            State string  `json:"state"`
            Email string  `json:"email"`
            Pincode string  `json:"pincode"`
            Name string  `json:"name"`
            City string  `json:"city"`
            Address1 string  `json:"address1"`
            Phone string  `json:"phone"`
            AddressType string  `json:"address_type"`
            Country string  `json:"country"`
            Address string  `json:"address"`
            Landmark string  `json:"landmark"`
         
    }
    
    // DebugInfo used by Order
    type DebugInfo struct {

        
            StormbreakerUUID string  `json:"stormbreaker_uuid"`
         
    }
    
    // BuyerDetails used by Order
    type BuyerDetails struct {

        
            State string  `json:"state"`
            Pincode float64  `json:"pincode"`
            Name string  `json:"name"`
            AjioSiteID string  `json:"ajio_site_id"`
            City string  `json:"city"`
            Address string  `json:"address"`
            Gstin string  `json:"gstin"`
         
    }
    
    // EinvoiceInfo used by Order
    type EinvoiceInfo struct {

        
            CreditNote map[string]interface{}  `json:"credit_note"`
            Invoice map[string]interface{}  `json:"invoice"`
         
    }
    
    // ShipmentMeta used by Order
    type ShipmentMeta struct {

        
            AwbNumber string  `json:"awb_number"`
            PackagingName string  `json:"packaging_name"`
            Timestamp ShipmentTimeStamp  `json:"timestamp"`
            ShipmentWeight float64  `json:"shipment_weight"`
            OrderType string  `json:"order_type"`
            DpOptions map[string]interface{}  `json:"dp_options"`
            Weight float64  `json:"weight"`
            ReturnDetails map[string]interface{}  `json:"return_details"`
            DpName string  `json:"dp_name"`
            ReturnAffiliateShipmentID string  `json:"return_affiliate_shipment_id"`
            DpSortKey string  `json:"dp_sort_key"`
            ForwardAffiliateOrderID string  `json:"forward_affiliate_order_id"`
            ReturnAwbNumber string  `json:"return_awb_number"`
            LockData LockData  `json:"lock_data"`
            AssignDpFromSb bool  `json:"assign_dp_from_sb"`
            ParentDpID string  `json:"parent_dp_id"`
            DebugInfo DebugInfo  `json:"debug_info"`
            External map[string]interface{}  `json:"external"`
            DpID string  `json:"dp_id"`
            BagWeight map[string]interface{}  `json:"bag_weight"`
            B2bBuyerDetails BuyerDetails  `json:"b2b_buyer_details"`
            ForwardAffiliateShipmentID string  `json:"forward_affiliate_shipment_id"`
            DueDate string  `json:"due_date"`
            EwaybillInfo map[string]interface{}  `json:"ewaybill_info"`
            Formatted Formatted  `json:"formatted"`
            ShipmentVolumetricWeight float64  `json:"shipment_volumetric_weight"`
            PoNumber string  `json:"po_number"`
            SameStoreAvailable bool  `json:"same_store_available"`
            B2cBuyerDetails map[string]interface{}  `json:"b2c_buyer_details"`
            FulfilmentPriorityText string  `json:"fulfilment_priority_text"`
            ShipmentTags []ShipmentTags  `json:"shipment_tags"`
            BoxType string  `json:"box_type"`
            ReturnStoreNode float64  `json:"return_store_node"`
            StoreInvoiceUpdatedDate string  `json:"store_invoice_updated_date"`
            AutoTriggerDpAssignmentAcf bool  `json:"auto_trigger_dp_assignment_acf"`
            MarketplaceStoreID string  `json:"marketplace_store_id"`
            EinvoiceInfo EinvoiceInfo  `json:"einvoice_info"`
            Dimension Dimensions  `json:"dimension"`
            ReturnAffiliateOrderID string  `json:"return_affiliate_order_id"`
         
    }
    
    // InvoiceInfo used by Order
    type InvoiceInfo struct {

        
            InvoiceURL string  `json:"invoice_url"`
            UpdatedDate string  `json:"updated_date"`
            StoreInvoiceID string  `json:"store_invoice_id"`
            LabelURL string  `json:"label_url"`
            CreditNoteID string  `json:"credit_note_id"`
            ExternalInvoiceID string  `json:"external_invoice_id"`
         
    }
    
    // OrderDetailsData used by Order
    type OrderDetailsData struct {

        
            Source string  `json:"source"`
            AffiliateID string  `json:"affiliate_id"`
            OrderValue string  `json:"order_value"`
            OrderDate string  `json:"order_date"`
            OrderingChannelLogo map[string]interface{}  `json:"ordering_channel_logo"`
            OrderingChannel string  `json:"ordering_channel"`
            TaxDetails map[string]interface{}  `json:"tax_details"`
            FyndOrderID string  `json:"fynd_order_id"`
            CodCharges string  `json:"cod_charges"`
         
    }
    
    // DPDetailsData used by Order
    type DPDetailsData struct {

        
            AwbNo string  `json:"awb_no"`
            ID float64  `json:"id"`
            Pincode string  `json:"pincode"`
            Name string  `json:"name"`
            Country string  `json:"country"`
            TrackURL string  `json:"track_url"`
            GstTag string  `json:"gst_tag"`
            EwayBillID string  `json:"eway_bill_id"`
         
    }
    
    // ShipmentStatusData used by Order
    type ShipmentStatusData struct {

        
            BagList []string  `json:"bag_list"`
            DisplayName string  `json:"display_name"`
            CreatedAt string  `json:"created_at"`
            ShipmentID string  `json:"shipment_id"`
            Meta map[string]interface{}  `json:"meta"`
            Status string  `json:"status"`
            ID float64  `json:"id"`
         
    }
    
    // OrderingStoreDetails used by Order
    type OrderingStoreDetails struct {

        
            StoreName string  `json:"store_name"`
            State string  `json:"state"`
            Pincode string  `json:"pincode"`
            City string  `json:"city"`
            ContactPerson string  `json:"contact_person"`
            Meta map[string]interface{}  `json:"meta"`
            Phone string  `json:"phone"`
            OrderingStoreID float64  `json:"ordering_store_id"`
            Country string  `json:"country"`
            Code string  `json:"code"`
            Address string  `json:"address"`
         
    }
    
    // PhoneDetails used by Order
    type PhoneDetails struct {

        
            MobileNumber string  `json:"mobile_number"`
            CountryCode float64  `json:"country_code"`
         
    }
    
    // ContactDetails used by Order
    type ContactDetails struct {

        
            Phone []PhoneDetails  `json:"phone"`
            Emails []string  `json:"emails"`
         
    }
    
    // CompanyDetails used by Order
    type CompanyDetails struct {

        
            CompanyGst string  `json:"company_gst"`
            CompanyID float64  `json:"company_id"`
            CompanyCin string  `json:"company_cin"`
            CompanyName string  `json:"company_name"`
            Address map[string]interface{}  `json:"address"`
            CompanyContact ContactDetails  `json:"company_contact"`
         
    }
    
    // ShipmentPayments used by Order
    type ShipmentPayments struct {

        
            Mode string  `json:"mode"`
            Source string  `json:"source"`
            Logo string  `json:"logo"`
         
    }
    
    // PDFLinks used by Order
    type PDFLinks struct {

        
            LabelPos string  `json:"label_pos"`
            LabelA4 string  `json:"label_a4"`
            LabelA6 string  `json:"label_a6"`
            InvoiceType string  `json:"invoice_type"`
            InvoicePos string  `json:"invoice_pos"`
            LabelType string  `json:"label_type"`
            InvoiceA4 string  `json:"invoice_a4"`
            Label string  `json:"label"`
            DeliveryChallanA4 string  `json:"delivery_challan_a4"`
            CreditNoteURL string  `json:"credit_note_url"`
            InvoiceA6 string  `json:"invoice_a6"`
            B2b string  `json:"b2b"`
            PoInvoice string  `json:"po_invoice"`
            Invoice string  `json:"invoice"`
         
    }
    
    // AffiliateDetails used by Order
    type AffiliateDetails struct {

        
            AffiliateStoreID string  `json:"affiliate_store_id"`
            ShipmentMeta ShipmentMeta  `json:"shipment_meta"`
            AffiliateID string  `json:"affiliate_id"`
            CompanyAffiliateTag string  `json:"company_affiliate_tag"`
            AffiliateMeta AffiliateMeta  `json:"affiliate_meta"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            AdID string  `json:"ad_id"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            PdfLinks PDFLinks  `json:"pdf_links"`
         
    }
    
    // ShipmentDetails used by Order
    type ShipmentDetails struct {

        
            LockMessage string  `json:"lock_message"`
            ActionToStatus map[string]interface{}  `json:"action_to_status"`
            LockStatus bool  `json:"lock_status"`
         
    }
    
    // B2BPODetails used by Order
    type B2BPODetails struct {

        
            ItemBasePrice float64  `json:"item_base_price"`
            PartialCanRet bool  `json:"partial_can_ret"`
            DockerNumber string  `json:"docker_number"`
            TotalGstPercentage float64  `json:"total_gst_percentage"`
            PoTaxAmount float64  `json:"po_tax_amount"`
            PoLineAmount float64  `json:"po_line_amount"`
         
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

        
            DocketNumber string  `json:"docket_number"`
            CustomMessage string  `json:"custom_message"`
            CustomJson map[string]interface{}  `json:"custom_json"`
            PartialCanRet bool  `json:"partial_can_ret"`
            B2bPoDetails B2BPODetails  `json:"b2b_po_details"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            GroupID string  `json:"group_id"`
            GiftCard GiftCard  `json:"gift_card"`
         
    }
    
    // CurrentStatus used by Order
    type CurrentStatus struct {

        
            UpdatedAt string  `json:"updated_at"`
            BagStateMapper BagStateMapper  `json:"bag_state_mapper"`
            DeliveryPartnerID float64  `json:"delivery_partner_id"`
            CreatedAt string  `json:"created_at"`
            BagID float64  `json:"bag_id"`
            KafkaSync bool  `json:"kafka_sync"`
            ShipmentID string  `json:"shipment_id"`
            StateType string  `json:"state_type"`
            Status string  `json:"status"`
            StoreID float64  `json:"store_id"`
            StateID float64  `json:"state_id"`
            CurrentStatusID float64  `json:"current_status_id"`
            DeliveryAwbNumber string  `json:"delivery_awb_number"`
         
    }
    
    // OrderBagArticle used by Order
    type OrderBagArticle struct {

        
            Identifiers map[string]interface{}  `json:"identifiers"`
            ReturnConfig map[string]interface{}  `json:"return_config"`
            UID string  `json:"uid"`
         
    }
    
    // BagGST used by Order
    type BagGST struct {

        
            GstFee float64  `json:"gst_fee"`
            GstinCode string  `json:"gstin_code"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            IsDefaultHsnCode bool  `json:"is_default_hsn_code"`
            CgstTaxPercentage float64  `json:"cgst_tax_percentage"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            SgstTaxPercentage float64  `json:"sgst_tax_percentage"`
            ValueOfGood float64  `json:"value_of_good"`
            IgstGstFee string  `json:"igst_gst_fee"`
            SgstGstFee string  `json:"sgst_gst_fee"`
            CgstGstFee string  `json:"cgst_gst_fee"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            GstTag string  `json:"gst_tag"`
            HsnCode string  `json:"hsn_code"`
            HsnCodeID string  `json:"hsn_code_id"`
            IgstTaxPercentage string  `json:"igst_tax_percentage"`
         
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

        
            PromoID string  `json:"promo_id"`
            DiscountRules []DiscountRules  `json:"discount_rules"`
            Amount float64  `json:"amount"`
            BuyRules []BuyRules  `json:"buy_rules"`
            ArticleQuantity float64  `json:"article_quantity"`
            MrpPromotion bool  `json:"mrp_promotion"`
            PromotionName string  `json:"promotion_name"`
            PromotionType string  `json:"promotion_type"`
         
    }
    
    // OrderBrandName used by Order
    type OrderBrandName struct {

        
            ID float64  `json:"id"`
            Company string  `json:"company"`
            CreatedOn string  `json:"created_on"`
            Logo string  `json:"logo"`
            BrandName string  `json:"brand_name"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // BagConfigs used by Order
    type BagConfigs struct {

        
            CanBeCancelled bool  `json:"can_be_cancelled"`
            AllowForceReturn bool  `json:"allow_force_return"`
            IsCustomerReturnAllowed bool  `json:"is_customer_return_allowed"`
            IsActive bool  `json:"is_active"`
            EnableTracking bool  `json:"enable_tracking"`
            IsReturnable bool  `json:"is_returnable"`
         
    }
    
    // OrderBags used by Order
    type OrderBags struct {

        
            DisplayName string  `json:"display_name"`
            Meta BagMeta  `json:"meta"`
            DeliveryAddress PlatformDeliveryAddress  `json:"delivery_address"`
            CanReturn bool  `json:"can_return"`
            SellerIdentifier string  `json:"seller_identifier"`
            BagID float64  `json:"bag_id"`
            Prices Prices  `json:"prices"`
            CanCancel bool  `json:"can_cancel"`
            CurrentStatus CurrentStatus  `json:"current_status"`
            Identifier string  `json:"identifier"`
            FinancialBreakup FinancialBreakup  `json:"financial_breakup"`
            EntityType string  `json:"entity_type"`
            Article OrderBagArticle  `json:"article"`
            GstDetails BagGST  `json:"gst_details"`
            LineNumber float64  `json:"line_number"`
            AppliedPromos []AppliedPromos  `json:"applied_promos"`
            Brand OrderBrandName  `json:"brand"`
            BagConfigs BagConfigs  `json:"bag_configs"`
            Item PlatformItem  `json:"item"`
            GroupID string  `json:"group_id"`
            Quantity float64  `json:"quantity"`
            ParentPromoBags map[string]interface{}  `json:"parent_promo_bags"`
         
    }
    
    // TrackingList used by Order
    type TrackingList struct {

        
            Text string  `json:"text"`
            IsCurrent bool  `json:"is_current"`
            IsPassed bool  `json:"is_passed"`
            Time string  `json:"time"`
            Status string  `json:"status"`
         
    }
    
    // FulfillingStore used by Order
    type FulfillingStore struct {

        
            StoreName string  `json:"store_name"`
            State string  `json:"state"`
            Pincode string  `json:"pincode"`
            City string  `json:"city"`
            ContactPerson string  `json:"contact_person"`
            FulfillmentChannel string  `json:"fulfillment_channel"`
            Meta map[string]interface{}  `json:"meta"`
            Phone string  `json:"phone"`
            Country string  `json:"country"`
            Code string  `json:"code"`
            Address string  `json:"address"`
            ID float64  `json:"id"`
         
    }
    
    // PlatformShipment used by Order
    type PlatformShipment struct {

        
            FulfilmentPriority float64  `json:"fulfilment_priority"`
            BillingDetails UserDetailsData  `json:"billing_details"`
            User UserDataInfo  `json:"user"`
            EnableDpTracking bool  `json:"enable_dp_tracking"`
            DeliverySlot map[string]interface{}  `json:"delivery_slot"`
            UserAgent string  `json:"user_agent"`
            Meta ShipmentMeta  `json:"meta"`
            PaymentMode string  `json:"payment_mode"`
            PickedDate string  `json:"picked_date"`
            Invoice InvoiceInfo  `json:"invoice"`
            CustomMessage string  `json:"custom_message"`
            Order OrderDetailsData  `json:"order"`
            Coupon map[string]interface{}  `json:"coupon"`
            Vertical string  `json:"vertical"`
            LockStatus bool  `json:"lock_status"`
            PackagingType string  `json:"packaging_type"`
            DpDetails DPDetailsData  `json:"dp_details"`
            Prices Prices  `json:"prices"`
            PaymentMethods map[string]interface{}  `json:"payment_methods"`
            TotalItems float64  `json:"total_items"`
            CustomMeta []map[string]interface{}  `json:"custom_meta"`
            Status ShipmentStatusData  `json:"status"`
            OrderingStore OrderingStoreDetails  `json:"ordering_store"`
            JourneyType string  `json:"journey_type"`
            ShipmentCreatedAt string  `json:"shipment_created_at"`
            CompanyDetails CompanyDetails  `json:"company_details"`
            GstDetails GSTDetailsData  `json:"gst_details"`
            Payments ShipmentPayments  `json:"payments"`
            DeliveryDetails UserDetailsData  `json:"delivery_details"`
            AffiliateDetails AffiliateDetails  `json:"affiliate_details"`
            BagStatusHistory []BagStatusHistory  `json:"bag_status_history"`
            ShipmentImages []string  `json:"shipment_images"`
            ShipmentDetails ShipmentDetails  `json:"shipment_details"`
            ShipmentStatus string  `json:"shipment_status"`
            OperationalStatus string  `json:"operational_status"`
            Bags []OrderBags  `json:"bags"`
            TrackingList []TrackingList  `json:"tracking_list"`
            FulfillingStore FulfillingStore  `json:"fulfilling_store"`
            PlatformLogo string  `json:"platform_logo"`
            InvoiceID string  `json:"invoice_id"`
            EstimatedSlaTime string  `json:"estimated_sla_time"`
            TotalBags float64  `json:"total_bags"`
            ShipmentID string  `json:"shipment_id"`
            ForwardShipmentID string  `json:"forward_shipment_id"`
            ShipmentQuantity float64  `json:"shipment_quantity"`
            PriorityText string  `json:"priority_text"`
         
    }
    
    // ShipmentInfoResponse used by Order
    type ShipmentInfoResponse struct {

        
            Success bool  `json:"success"`
            Shipments []PlatformShipment  `json:"shipments"`
            Message string  `json:"message"`
         
    }
    
    // AssetByShipment used by Order
    type AssetByShipment struct {

        
            ShipmentID string  `json:"shipment_id"`
            Assets map[string]string  `json:"assets"`
            ExpiresIn string  `json:"expires_in"`
         
    }
    
    // ResponseGetAssetShipment used by Order
    type ResponseGetAssetShipment struct {

        
            Success bool  `json:"success"`
            PresignedType string  `json:"presigned_type"`
            Result []AssetByShipment  `json:"result"`
         
    }
    
    // TransactionData used by Order
    type TransactionData struct {

        
            TransactionID string  `json:"transaction_id"`
            TerminalID string  `json:"terminal_id"`
            Currency string  `json:"currency"`
            PaymentID string  `json:"payment_id"`
            Status string  `json:"status"`
            UniqueReferenceNumber string  `json:"unique_reference_number"`
            AmountPaid string  `json:"amount_paid"`
            Entity string  `json:"entity"`
         
    }
    
    // BillingStaffDetails used by Order
    type BillingStaffDetails struct {

        
            EmployeeCode string  `json:"employee_code"`
            User string  `json:"user"`
            StaffID float64  `json:"staff_id"`
            FirstName string  `json:"first_name"`
            LastName string  `json:"last_name"`
         
    }
    
    // PlatformUserDetails used by Order
    type PlatformUserDetails struct {

        
            PlatformUserID string  `json:"platform_user_id"`
            PlatformUserEmployeeCode string  `json:"platform_user_employee_code"`
            PlatformUserFirstName string  `json:"platform_user_first_name"`
            PlatformUserLastName string  `json:"platform_user_last_name"`
         
    }
    
    // OrderMeta used by Order
    type OrderMeta struct {

        
            OrderChildEntities []string  `json:"order_child_entities"`
            OrderType string  `json:"order_type"`
            Staff map[string]interface{}  `json:"staff"`
            PaymentType string  `json:"payment_type"`
            MongoCartID float64  `json:"mongo_cart_id"`
            Comment string  `json:"comment"`
            OrderTags []map[string]interface{}  `json:"order_tags"`
            CurrencySymbol string  `json:"currency_symbol"`
            CustomerNote string  `json:"customer_note"`
            OrderingStore float64  `json:"ordering_store"`
            CompanyLogo string  `json:"company_logo"`
            TransactionData TransactionData  `json:"transaction_data"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            BillingStaffDetails BillingStaffDetails  `json:"billing_staff_details"`
            Files []map[string]interface{}  `json:"files"`
            PlatformUserDetails PlatformUserDetails  `json:"platform_user_details"`
            OrderPlatform string  `json:"order_platform"`
            EmployeeID float64  `json:"employee_id"`
            CartID float64  `json:"cart_id"`
         
    }
    
    // TaxDetails used by Order
    type TaxDetails struct {

        
            PanNo string  `json:"pan_no"`
            Gstin string  `json:"gstin"`
         
    }
    
    // OrderDict used by Order
    type OrderDict struct {

        
            OrderDate string  `json:"order_date"`
            Prices Prices  `json:"prices"`
            PaymentMethods map[string]interface{}  `json:"payment_methods"`
            Meta OrderMeta  `json:"meta"`
            TaxDetails TaxDetails  `json:"tax_details"`
            FyndOrderID string  `json:"fynd_order_id"`
         
    }
    
    // ShipmentDetailsResponse used by Order
    type ShipmentDetailsResponse struct {

        
            Success bool  `json:"success"`
            Order OrderDict  `json:"order"`
            Shipments []PlatformShipment  `json:"shipments"`
         
    }
    
    // SubLane used by Order
    type SubLane struct {

        
            Actions []map[string]interface{}  `json:"actions"`
            Text string  `json:"text"`
            Index float64  `json:"index"`
            TotalItems float64  `json:"total_items"`
            Value string  `json:"value"`
         
    }
    
    // SuperLane used by Order
    type SuperLane struct {

        
            Options []SubLane  `json:"options"`
            Value string  `json:"value"`
            Text string  `json:"text"`
            TotalItems float64  `json:"total_items"`
         
    }
    
    // LaneConfigResponse used by Order
    type LaneConfigResponse struct {

        
            SuperLanes []SuperLane  `json:"super_lanes"`
         
    }
    
    // PlatformBreakupValues used by Order
    type PlatformBreakupValues struct {

        
            Value string  `json:"value"`
            Display string  `json:"display"`
            Name string  `json:"name"`
         
    }
    
    // PlatformChannel used by Order
    type PlatformChannel struct {

        
            Logo string  `json:"logo"`
            Name string  `json:"name"`
         
    }
    
    // PlatformOrderItems used by Order
    type PlatformOrderItems struct {

        
            BreakupValues []PlatformBreakupValues  `json:"breakup_values"`
            Channel PlatformChannel  `json:"channel"`
            OrderCreatedTime string  `json:"order_created_time"`
            TotalOrderValue float64  `json:"total_order_value"`
            Shipments []PlatformShipment  `json:"shipments"`
            OrderValue float64  `json:"order_value"`
            Meta map[string]interface{}  `json:"meta"`
            PaymentMode string  `json:"payment_mode"`
            UserInfo UserDataInfo  `json:"user_info"`
            OrderID string  `json:"order_id"`
         
    }
    
    // OrderListingResponse used by Order
    type OrderListingResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
            TotalCount float64  `json:"total_count"`
            Lane string  `json:"lane"`
            Items []PlatformOrderItems  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // Options used by Order
    type Options struct {

        
            Value float64  `json:"value"`
            Text string  `json:"text"`
         
    }
    
    // MetricsCount used by Order
    type MetricsCount struct {

        
            Options []Options  `json:"options"`
            Key string  `json:"key"`
            Value float64  `json:"value"`
            Text string  `json:"text"`
         
    }
    
    // MetricCountResponse used by Order
    type MetricCountResponse struct {

        
            Items []MetricsCount  `json:"items"`
         
    }
    
    // PlatformTrack used by Order
    type PlatformTrack struct {

        
            UpdatedAt string  `json:"updated_at"`
            AccountName string  `json:"account_name"`
            ShipmentType string  `json:"shipment_type"`
            Reason string  `json:"reason"`
            RawStatus string  `json:"raw_status"`
            Meta map[string]interface{}  `json:"meta"`
            Status string  `json:"status"`
            UpdatedTime string  `json:"updated_time"`
            Awb string  `json:"awb"`
            LastLocationRecievedAt string  `json:"last_location_recieved_at"`
         
    }
    
    // PlatformShipmentTrack used by Order
    type PlatformShipmentTrack struct {

        
            Meta map[string]interface{}  `json:"meta"`
            Results []PlatformTrack  `json:"results"`
         
    }
    
    // FilterInfoOption used by Order
    type FilterInfoOption struct {

        
            Value string  `json:"value"`
            Text string  `json:"text"`
            Name string  `json:"name"`
         
    }
    
    // FiltersInfo used by Order
    type FiltersInfo struct {

        
            Type string  `json:"type"`
            Options []FilterInfoOption  `json:"options"`
            Value string  `json:"value"`
            Text string  `json:"text"`
         
    }
    
    // AdvanceFilterInfo used by Order
    type AdvanceFilterInfo struct {

        
            ActionCentre []FiltersInfo  `json:"action_centre"`
            Returned []FiltersInfo  `json:"returned"`
            Filters []FiltersInfo  `json:"filters"`
            AppliedFilters map[string]interface{}  `json:"applied_filters"`
            Processed []FiltersInfo  `json:"processed"`
            Unfulfilled []FiltersInfo  `json:"unfulfilled"`
            Page map[string]interface{}  `json:"page"`
         
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

        
            ReportCreatedAt string  `json:"report_created_at"`
            DisplayName string  `json:"display_name"`
            ReportID string  `json:"report_id"`
            RequestDetails map[string]interface{}  `json:"request_details"`
            ReportType string  `json:"report_type"`
            ReportRequestedAt string  `json:"report_requested_at"`
            ReportName string  `json:"report_name"`
            S3Key string  `json:"s3_key"`
            Status string  `json:"status"`
         
    }
    
    // JioCodeUpsertDataSet used by Order
    type JioCodeUpsertDataSet struct {

        
            ItemID string  `json:"item_id"`
            CompanyID string  `json:"company_id"`
            JioCode string  `json:"jio_code"`
            ArticleID string  `json:"article_id"`
         
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

        
            Identifier string  `json:"identifier"`
            Success bool  `json:"success"`
            Error []NestedErrorSchemaDataSet  `json:"error"`
            Data []map[string]interface{}  `json:"data"`
            TraceID string  `json:"trace_id"`
         
    }
    
    // FileUploadResponse used by Order
    type FileUploadResponse struct {

        
            Expiry float64  `json:"expiry"`
            URL string  `json:"url"`
         
    }
    
    // URL used by Order
    type URL struct {

        
            URL string  `json:"url"`
         
    }
    
    // FileResponse used by Order
    type FileResponse struct {

        
            Upload FileUploadResponse  `json:"upload"`
            Namespace string  `json:"namespace"`
            Size float64  `json:"size"`
            FileName string  `json:"file_name"`
            FilePath string  `json:"file_path"`
            Tags []string  `json:"tags"`
            Cdn URL  `json:"cdn"`
            Operation string  `json:"operation"`
            ContentType string  `json:"content_type"`
            Method string  `json:"method"`
         
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
    
    // BagReturnableCancelableStatus1 used by Order
    type BagReturnableCancelableStatus1 struct {

        
            CanBeCancelled bool  `json:"can_be_cancelled"`
            IsCustomerReturnAllowed bool  `json:"is_customer_return_allowed"`
            IsActive bool  `json:"is_active"`
            EnableTracking bool  `json:"enable_tracking"`
            IsReturnable bool  `json:"is_returnable"`
         
    }
    
    // EInvoicePortalDetails used by Order
    type EInvoicePortalDetails struct {

        
            Username string  `json:"username"`
            User string  `json:"user"`
            Password string  `json:"password"`
         
    }
    
    // Document used by Order
    type Document struct {

        
            DsType string  `json:"ds_type"`
            URL string  `json:"url"`
            Verified bool  `json:"verified"`
            Value string  `json:"value"`
            LegalName string  `json:"legal_name"`
         
    }
    
    // StoreDocuments used by Order
    type StoreDocuments struct {

        
            Gst Document  `json:"gst"`
         
    }
    
    // StoreEinvoice used by Order
    type StoreEinvoice struct {

        
            Enabled bool  `json:"enabled"`
            Username string  `json:"username"`
            Password string  `json:"password"`
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
    
    // StoreMeta used by Order
    type StoreMeta struct {

        
            DisplayName string  `json:"display_name"`
            ProductReturnConfig map[string]interface{}  `json:"product_return_config"`
            EinvoicePortalDetails EInvoicePortalDetails  `json:"einvoice_portal_details"`
            Documents StoreDocuments  `json:"documents"`
            NotificationEmails []string  `json:"notification_emails"`
            Stage string  `json:"stage"`
            GstCredentials StoreGstCredentials  `json:"gst_credentials"`
            AdditionalContactDetails map[string]interface{}  `json:"additional_contact_details"`
            GstNumber string  `json:"gst_number"`
            EwaybillPortalDetails map[string]interface{}  `json:"ewaybill_portal_details"`
            Timing []map[string]interface{}  `json:"timing"`
         
    }
    
    // StoreAddress used by Order
    type StoreAddress struct {

        
            Latitude float64  `json:"latitude"`
            Area string  `json:"area"`
            Address1 string  `json:"address1"`
            AddressCategory string  `json:"address_category"`
            AddressType string  `json:"address_type"`
            CountryCode string  `json:"country_code"`
            UpdatedAt string  `json:"updated_at"`
            Longitude float64  `json:"longitude"`
            Pincode float64  `json:"pincode"`
            CreatedAt string  `json:"created_at"`
            City string  `json:"city"`
            Phone string  `json:"phone"`
            State string  `json:"state"`
            Email string  `json:"email"`
            Version string  `json:"version"`
            Address2 string  `json:"address2"`
            ContactPerson string  `json:"contact_person"`
            Country string  `json:"country"`
            Landmark string  `json:"landmark"`
         
    }
    
    // Store used by Order
    type Store struct {

        
            Latitude float64  `json:"latitude"`
            IsArchived bool  `json:"is_archived"`
            OrderIntegrationID string  `json:"order_integration_id"`
            IsEnabledForRecon bool  `json:"is_enabled_for_recon"`
            MallArea string  `json:"mall_area"`
            IsActive bool  `json:"is_active"`
            Code string  `json:"code"`
            Meta StoreMeta  `json:"meta"`
            MallName string  `json:"mall_name"`
            Name string  `json:"name"`
            ParentStoreID float64  `json:"parent_store_id"`
            Address1 string  `json:"address1"`
            LoginUsername string  `json:"login_username"`
            VatNo string  `json:"vat_no"`
            BrandID interface{}  `json:"brand_id"`
            UpdatedAt string  `json:"updated_at"`
            Longitude float64  `json:"longitude"`
            Pincode string  `json:"pincode"`
            CreatedAt string  `json:"created_at"`
            City string  `json:"city"`
            CompanyID float64  `json:"company_id"`
            SID string  `json:"s_id"`
            LocationType string  `json:"location_type"`
            Phone float64  `json:"phone"`
            BrandStoreTags []string  `json:"brand_store_tags"`
            StoreEmail string  `json:"store_email"`
            State string  `json:"state"`
            StoreActiveFrom string  `json:"store_active_from"`
            PackagingMaterialCount float64  `json:"packaging_material_count"`
            AlohomoraUserID float64  `json:"alohomora_user_id"`
            Address2 string  `json:"address2"`
            ContactPerson string  `json:"contact_person"`
            FulfillmentChannel string  `json:"fulfillment_channel"`
            StoreAddressJson StoreAddress  `json:"store_address_json"`
            Country string  `json:"country"`
         
    }
    
    // BagGSTDetails used by Order
    type BagGSTDetails struct {

        
            GstFee float64  `json:"gst_fee"`
            GstinCode string  `json:"gstin_code"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            IsDefaultHsnCode bool  `json:"is_default_hsn_code"`
            CgstTaxPercentage float64  `json:"cgst_tax_percentage"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            SgstTaxPercentage float64  `json:"sgst_tax_percentage"`
            ValueOfGood float64  `json:"value_of_good"`
            IgstGstFee string  `json:"igst_gst_fee"`
            SgstGstFee string  `json:"sgst_gst_fee"`
            CgstGstFee string  `json:"cgst_gst_fee"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            GstTag string  `json:"gst_tag"`
            HsnCode string  `json:"hsn_code"`
            HsnCodeID string  `json:"hsn_code_id"`
            IgstTaxPercentage float64  `json:"igst_tax_percentage"`
         
    }
    
    // Brand used by Order
    type Brand struct {

        
            ScriptLastRan string  `json:"script_last_ran"`
            Company string  `json:"company"`
            InvoicePrefix string  `json:"invoice_prefix"`
            CreatedOn float64  `json:"created_on"`
            CreditNoteExpiryDays float64  `json:"credit_note_expiry_days"`
            StartDate string  `json:"start_date"`
            Logo string  `json:"logo"`
            BrandName string  `json:"brand_name"`
            IsVirtualInvoice bool  `json:"is_virtual_invoice"`
            PickupLocation string  `json:"pickup_location"`
            BrandID float64  `json:"brand_id"`
            CreditNoteAllowed bool  `json:"credit_note_allowed"`
            ModifiedOn float64  `json:"modified_on"`
         
    }
    
    // Attributes used by Order
    type Attributes struct {

        
            Gender []string  `json:"gender"`
            Name string  `json:"name"`
            PrimaryColorHex string  `json:"primary_color_hex"`
            MarketerName string  `json:"marketer_name"`
            Essential string  `json:"essential"`
            BrandName string  `json:"brand_name"`
            MarketerAddress string  `json:"marketer_address"`
            PrimaryMaterial string  `json:"primary_material"`
            PrimaryColor string  `json:"primary_color"`
         
    }
    
    // Item used by Order
    type Item struct {

        
            Color string  `json:"color"`
            Gender string  `json:"gender"`
            LastUpdatedAt string  `json:"last_updated_at"`
            L1CategoryID float64  `json:"l1_category_id"`
            Code string  `json:"code"`
            Meta map[string]interface{}  `json:"meta"`
            L3CategoryName string  `json:"l3_category_name"`
            L2CategoryID float64  `json:"l2_category_id"`
            L1Category []string  `json:"l1_category"`
            CanReturn bool  `json:"can_return"`
            Name string  `json:"name"`
            Image []string  `json:"image"`
            WebstoreProductURL string  `json:"webstore_product_url"`
            DepartmentID float64  `json:"department_id"`
            CanCancel bool  `json:"can_cancel"`
            BrandID float64  `json:"brand_id"`
            ItemID float64  `json:"item_id"`
            L3Category float64  `json:"l3_category"`
            Attributes Attributes  `json:"attributes"`
            Brand string  `json:"brand"`
            BranchURL string  `json:"branch_url"`
            Size string  `json:"size"`
            SlugKey string  `json:"slug_key"`
            L2Category []string  `json:"l2_category"`
         
    }
    
    // ArticleDetails used by Order
    type ArticleDetails struct {

        
            Status map[string]interface{}  `json:"status"`
         
    }
    
    // BagDetailsPlatformResponse used by Order
    type BagDetailsPlatformResponse struct {

        
            DisplayName string  `json:"display_name"`
            OrderIntegrationID string  `json:"order_integration_id"`
            CurrentOperationalStatus BagStatusHistory  `json:"current_operational_status"`
            Tags []string  `json:"tags"`
            Meta BagMeta  `json:"meta"`
            BagStatus []BagStatusHistory  `json:"bag_status"`
            RestoreCoupon bool  `json:"restore_coupon"`
            SellerIdentifier string  `json:"seller_identifier"`
            BID float64  `json:"b_id"`
            Prices Prices  `json:"prices"`
            CurrentStatus BagStatusHistory  `json:"current_status"`
            Status BagReturnableCancelableStatus1  `json:"status"`
            RestorePromos map[string]interface{}  `json:"restore_promos"`
            OrderingStore Store  `json:"ordering_store"`
            JourneyType string  `json:"journey_type"`
            Identifier string  `json:"identifier"`
            FinancialBreakup []FinancialBreakup  `json:"financial_breakup"`
            Article Article  `json:"article"`
            AffiliateBagDetails AffiliateBagDetails  `json:"affiliate_bag_details"`
            EntityType string  `json:"entity_type"`
            GstDetails BagGSTDetails  `json:"gst_details"`
            LineNumber float64  `json:"line_number"`
            AppliedPromos []map[string]interface{}  `json:"applied_promos"`
            OriginalBagList []float64  `json:"original_bag_list"`
            AffiliateDetails AffiliateDetails  `json:"affiliate_details"`
            Reasons []map[string]interface{}  `json:"reasons"`
            BagStatusHistory BagStatusHistory  `json:"bag_status_history"`
            NoOfBagsOrder float64  `json:"no_of_bags_order"`
            Brand Brand  `json:"brand"`
            Item Item  `json:"item"`
            OperationalStatus string  `json:"operational_status"`
            BagUpdateTime float64  `json:"bag_update_time"`
            BType string  `json:"b_type"`
            ArticleDetails ArticleDetails  `json:"article_details"`
            Quantity float64  `json:"quantity"`
            ShipmentID string  `json:"shipment_id"`
            ParentPromoBags map[string]interface{}  `json:"parent_promo_bags"`
            QcRequired interface{}  `json:"qc_required"`
            Dates Dates  `json:"dates"`
         
    }
    
    // ErrorResponse used by Order
    type ErrorResponse struct {

        
            Message string  `json:"message"`
            Error string  `json:"error"`
         
    }
    
    // BagsPage used by Order
    type BagsPage struct {

        
            ItemTotal float64  `json:"item_total"`
            Size float64  `json:"size"`
            PageType string  `json:"page_type"`
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

        
            Success bool  `json:"success"`
            PaymentReceipt string  `json:"payment_receipt"`
            InvoiceReceipt string  `json:"invoice_receipt"`
            OrderID string  `json:"order_id"`
         
    }
    
    // InvalidateShipmentCachePayload used by Order
    type InvalidateShipmentCachePayload struct {

        
            BagIds []string  `json:"bag_ids"`
            ShipmentIds []string  `json:"shipment_ids"`
            AffiliateBagIds []string  `json:"affiliate_bag_ids"`
         
    }
    
    // InvalidateShipmentCacheNestedResponse used by Order
    type InvalidateShipmentCacheNestedResponse struct {

        
            Error string  `json:"error"`
            Message string  `json:"message"`
            Status float64  `json:"status"`
            ShipmentID string  `json:"shipment_id"`
         
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
            Success bool  `json:"success"`
         
    }
    
    // StoreReassign used by Order
    type StoreReassign struct {

        
            ReasonIds []float64  `json:"reason_ids"`
            BagID float64  `json:"bag_id"`
            StoreID float64  `json:"store_id"`
            AffiliateID string  `json:"affiliate_id"`
            MongoArticleID string  `json:"mongo_article_id"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            FyndOrderID string  `json:"fynd_order_id"`
            SetID string  `json:"set_id"`
            ItemID string  `json:"item_id"`
         
    }
    
    // StoreReassignResponse used by Order
    type StoreReassignResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // Entities used by Order
    type Entities struct {

        
            ID string  `json:"id"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            AffiliateID string  `json:"affiliate_id"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            ReasonText string  `json:"reason_text"`
         
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
            AffiliateOrderID string  `json:"affiliate_order_id"`
            BagID float64  `json:"bag_id"`
            IsLocked bool  `json:"is_locked"`
         
    }
    
    // OriginalFilter used by Order
    type OriginalFilter struct {

        
            AffiliateID string  `json:"affiliate_id"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
         
    }
    
    // CheckResponse used by Order
    type CheckResponse struct {

        
            IsBagLocked bool  `json:"is_bag_locked"`
            Status string  `json:"status"`
            ShipmentID string  `json:"shipment_id"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            Bags []Bags  `json:"bags"`
            AffiliateID string  `json:"affiliate_id"`
            IsShipmentLocked bool  `json:"is_shipment_locked"`
            OriginalFilter OriginalFilter  `json:"original_filter"`
            LockStatus string  `json:"lock_status"`
         
    }
    
    // UpdateShipmentLockResponse used by Order
    type UpdateShipmentLockResponse struct {

        
            CheckResponse []CheckResponse  `json:"check_response"`
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // AnnouncementResponse used by Order
    type AnnouncementResponse struct {

        
            ID float64  `json:"id"`
            FromDatetime string  `json:"from_datetime"`
            PlatformID string  `json:"platform_id"`
            ToDatetime string  `json:"to_datetime"`
            Description string  `json:"description"`
            CreatedAt string  `json:"created_at"`
            CompanyID float64  `json:"company_id"`
            PlatformName string  `json:"platform_name"`
            Title string  `json:"title"`
            LogoURL string  `json:"logo_url"`
         
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

        
            Products []ProductsDataUpdates  `json:"products"`
            Entities []EntitiesDataUpdates  `json:"entities"`
         
    }
    
    // Products used by Order
    type Products struct {

        
            LineNumber float64  `json:"line_number"`
            Quantity float64  `json:"quantity"`
            Identifier string  `json:"identifier"`
         
    }
    
    // ProductsReasonsData used by Order
    type ProductsReasonsData struct {

        
            ReasonID float64  `json:"reason_id"`
            ReasonText string  `json:"reason_text"`
         
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
    
    // EntityReasonData used by Order
    type EntityReasonData struct {

        
            ReasonID float64  `json:"reason_id"`
            ReasonText string  `json:"reason_text"`
         
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
    
    // ShipmentsRequest used by Order
    type ShipmentsRequest struct {

        
            DataUpdates DataUpdates  `json:"data_updates"`
            Products []Products  `json:"products"`
            Identifier string  `json:"identifier"`
            Reasons ReasonsData  `json:"reasons"`
         
    }
    
    // StatuesRequest used by Order
    type StatuesRequest struct {

        
            ExcludeBagsNextState string  `json:"exclude_bags_next_state"`
            Shipments []ShipmentsRequest  `json:"shipments"`
            Status string  `json:"status"`
         
    }
    
    // UpdateShipmentStatusRequest used by Order
    type UpdateShipmentStatusRequest struct {

        
            Statuses []StatuesRequest  `json:"statuses"`
            UnlockBeforeTransition bool  `json:"unlock_before_transition"`
            Task bool  `json:"task"`
            LockAfterTransition bool  `json:"lock_after_transition"`
            ForceTransition bool  `json:"force_transition"`
         
    }
    
    // ShipmentsResponse used by Order
    type ShipmentsResponse struct {

        
            Message string  `json:"message"`
            Status float64  `json:"status"`
            Exception string  `json:"exception"`
            StackTrace string  `json:"stack_trace"`
            FinalState map[string]interface{}  `json:"final_state"`
            Meta map[string]interface{}  `json:"meta"`
            Code string  `json:"code"`
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
    
    // AffiliateStoreIdMapping used by Order
    type AffiliateStoreIdMapping struct {

        
            MarketplaceStoreID string  `json:"marketplace_store_id"`
            StoreID float64  `json:"store_id"`
         
    }
    
    // AffiliateInventoryPaymentConfig used by Order
    type AffiliateInventoryPaymentConfig struct {

        
            Source string  `json:"source"`
            ModeOfPayment string  `json:"mode_of_payment"`
         
    }
    
    // AffiliateInventoryStoreConfig used by Order
    type AffiliateInventoryStoreConfig struct {

        
            Store map[string]interface{}  `json:"store"`
         
    }
    
    // AffiliateInventoryArticleAssignmentConfig used by Order
    type AffiliateInventoryArticleAssignmentConfig struct {

        
            PostOrderReassignment bool  `json:"post_order_reassignment"`
         
    }
    
    // AffiliateInventoryOrderConfig used by Order
    type AffiliateInventoryOrderConfig struct {

        
            ForceReassignment bool  `json:"force_reassignment"`
         
    }
    
    // AffiliateInventoryLogisticsConfig used by Order
    type AffiliateInventoryLogisticsConfig struct {

        
            DpAssignment bool  `json:"dp_assignment"`
         
    }
    
    // AffiliateInventoryConfig used by Order
    type AffiliateInventoryConfig struct {

        
            Payment AffiliateInventoryPaymentConfig  `json:"payment"`
            Inventory AffiliateInventoryStoreConfig  `json:"inventory"`
            ArticleAssignment AffiliateInventoryArticleAssignmentConfig  `json:"article_assignment"`
            Order AffiliateInventoryOrderConfig  `json:"order"`
            Logistics AffiliateInventoryLogisticsConfig  `json:"logistics"`
         
    }
    
    // AffiliateAppConfigMeta used by Order
    type AffiliateAppConfigMeta struct {

        
            Value string  `json:"value"`
            Name string  `json:"name"`
         
    }
    
    // AffiliateAppConfig used by Order
    type AffiliateAppConfig struct {

        
            Token string  `json:"token"`
            ID string  `json:"id"`
            Secret string  `json:"secret"`
            Name string  `json:"name"`
            CreatedAt string  `json:"created_at"`
            Description string  `json:"description"`
            Owner string  `json:"owner"`
            Meta []AffiliateAppConfigMeta  `json:"meta"`
            UpdatedAt string  `json:"updated_at"`
         
    }
    
    // AffiliateConfig used by Order
    type AffiliateConfig struct {

        
            Inventory AffiliateInventoryConfig  `json:"inventory"`
            App AffiliateAppConfig  `json:"app"`
         
    }
    
    // Affiliate used by Order
    type Affiliate struct {

        
            Token string  `json:"token"`
            ID string  `json:"id"`
            Config AffiliateConfig  `json:"config"`
         
    }
    
    // OrderConfig used by Order
    type OrderConfig struct {

        
            CreateUser bool  `json:"create_user"`
            ArticleLookup string  `json:"article_lookup"`
            AffiliateStoreIDMapping []AffiliateStoreIdMapping  `json:"affiliate_store_id_mapping"`
            Affiliate Affiliate  `json:"affiliate"`
            StoreLookup string  `json:"store_lookup"`
            BagEndState string  `json:"bag_end_state"`
         
    }
    
    // ArticleDetails1 used by Order
    type ArticleDetails1 struct {

        
            ID string  `json:"_id"`
            BrandID float64  `json:"brand_id"`
            Quantity float64  `json:"quantity"`
            Weight map[string]interface{}  `json:"weight"`
            Category map[string]interface{}  `json:"category"`
            Attributes map[string]interface{}  `json:"attributes"`
            Dimension map[string]interface{}  `json:"dimension"`
         
    }
    
    // ShipmentDetails1 used by Order
    type ShipmentDetails1 struct {

        
            Articles []ArticleDetails1  `json:"articles"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            FulfillmentID float64  `json:"fulfillment_id"`
            BoxType string  `json:"box_type"`
            Shipments float64  `json:"shipments"`
            Meta map[string]interface{}  `json:"meta"`
            DpID float64  `json:"dp_id"`
         
    }
    
    // LocationDetails used by Order
    type LocationDetails struct {

        
            Articles []ArticleDetails1  `json:"articles"`
            FulfillmentID float64  `json:"fulfillment_id"`
            FulfillmentType string  `json:"fulfillment_type"`
         
    }
    
    // ShipmentConfig used by Order
    type ShipmentConfig struct {

        
            Shipment []ShipmentDetails1  `json:"shipment"`
            Source string  `json:"source"`
            PaymentMode string  `json:"payment_mode"`
            LocationDetails LocationDetails  `json:"location_details"`
            Journey string  `json:"journey"`
            ToPincode string  `json:"to_pincode"`
            Action string  `json:"action"`
            Identifier string  `json:"identifier"`
         
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

        
            AvlQty float64  `json:"avl_qty"`
            CompanyID float64  `json:"company_id"`
            UnitPrice float64  `json:"unit_price"`
            ModifiedOn string  `json:"modified_on"`
            Sku string  `json:"sku"`
            PdfLinks MarketPlacePdf  `json:"pdf_links"`
            ID string  `json:"_id"`
            HsnCodeID string  `json:"hsn_code_id"`
            FyndStoreID string  `json:"fynd_store_id"`
            AmountPaid float64  `json:"amount_paid"`
            PriceMarked float64  `json:"price_marked"`
            Identifier map[string]interface{}  `json:"identifier"`
            ItemSize string  `json:"item_size"`
            Quantity float64  `json:"quantity"`
            SellerIdentifier string  `json:"seller_identifier"`
            TransferPrice float64  `json:"transfer_price"`
            Discount float64  `json:"discount"`
            AffiliateMeta map[string]interface{}  `json:"affiliate_meta"`
            StoreID float64  `json:"store_id"`
            AffiliateStoreID string  `json:"affiliate_store_id"`
            PriceEffective float64  `json:"price_effective"`
            DeliveryCharge float64  `json:"delivery_charge"`
            ItemID float64  `json:"item_id"`
         
    }
    
    // OrderUser used by Order
    type OrderUser struct {

        
            Address1 string  `json:"address1"`
            FirstName string  `json:"first_name"`
            Mobile float64  `json:"mobile"`
            Pincode string  `json:"pincode"`
            Phone float64  `json:"phone"`
            LastName string  `json:"last_name"`
            State string  `json:"state"`
            City string  `json:"city"`
            Email string  `json:"email"`
            Country string  `json:"country"`
            Address2 string  `json:"address2"`
         
    }
    
    // UserData used by Order
    type UserData struct {

        
            BillingUser OrderUser  `json:"billing_user"`
            ShippingUser OrderUser  `json:"shipping_user"`
         
    }
    
    // OrderPriority used by Order
    type OrderPriority struct {

        
            FulfilmentPriority float64  `json:"fulfilment_priority"`
            AffiliatePriorityCode string  `json:"affiliate_priority_code"`
            FulfilmentPriorityText string  `json:"fulfilment_priority_text"`
         
    }
    
    // OrderInfo used by Order
    type OrderInfo struct {

        
            DeliveryCharges float64  `json:"delivery_charges"`
            Shipment ShipmentData  `json:"shipment"`
            OrderValue float64  `json:"order_value"`
            CodCharges float64  `json:"cod_charges"`
            Coupon string  `json:"coupon"`
            Payment map[string]interface{}  `json:"payment"`
            PaymentMode string  `json:"payment_mode"`
            Bags []AffiliateBag  `json:"bags"`
            BillingAddress OrderUser  `json:"billing_address"`
            ShippingAddress OrderUser  `json:"shipping_address"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            Discount float64  `json:"discount"`
            User UserData  `json:"user"`
            OrderPriority OrderPriority  `json:"order_priority"`
            Items map[string]interface{}  `json:"items"`
         
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

        
            DisplayText string  `json:"display_text"`
            ID float64  `json:"id"`
            Slug string  `json:"slug"`
            Description string  `json:"description"`
         
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

        
            Message string  `json:"message"`
            L3Detail string  `json:"l3_detail"`
            BagID float64  `json:"bag_id"`
            DisplayMessage string  `json:"display_message"`
            TicketID string  `json:"ticket_id"`
            Createdat string  `json:"createdat"`
            AssignedAgent string  `json:"assigned_agent"`
            TicketURL string  `json:"ticket_url"`
            L2Detail string  `json:"l2_detail"`
            Meta map[string]interface{}  `json:"meta"`
            User string  `json:"user"`
            Type string  `json:"type"`
            L1Detail string  `json:"l1_detail"`
         
    }
    
    // ShipmentHistoryResponse used by Order
    type ShipmentHistoryResponse struct {

        
            ActivityHistory []HistoryDict  `json:"activity_history"`
            Success bool  `json:"success"`
         
    }
    
    // SmsDataPayload used by Order
    type SmsDataPayload struct {

        
            CustomerName string  `json:"customer_name"`
            Message string  `json:"message"`
            BrandName string  `json:"brand_name"`
            PhoneNumber float64  `json:"phone_number"`
            PaymentMode string  `json:"payment_mode"`
            ShipmentID float64  `json:"shipment_id"`
            OrderID string  `json:"order_id"`
            CountryCode string  `json:"country_code"`
            AmountPaid float64  `json:"amount_paid"`
         
    }
    
    // SendSmsPayload used by Order
    type SendSmsPayload struct {

        
            Data SmsDataPayload  `json:"data"`
            Slug string  `json:"slug"`
            BagID float64  `json:"bag_id"`
         
    }
    
    // Meta used by Order
    type Meta struct {

        
            StateManagerUsed string  `json:"state_manager_used"`
            KafkaEmissionStatus float64  `json:"kafka_emission_status"`
         
    }
    
    // ShipmentDetail used by Order
    type ShipmentDetail struct {

        
            ID float64  `json:"id"`
            Status string  `json:"status"`
            ShipmentID string  `json:"shipment_id"`
            BagList []float64  `json:"bag_list"`
            Meta Meta  `json:"meta"`
            Remarks string  `json:"remarks"`
         
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

        
            Result []OrderStatusData  `json:"result"`
            Success string  `json:"success"`
         
    }
    
    // Dimension used by Order
    type Dimension struct {

        
            Length float64  `json:"length"`
            PackagingType string  `json:"packaging_type"`
            Weight string  `json:"weight"`
            Width float64  `json:"width"`
            Height string  `json:"height"`
         
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
    
    // TaxInfo used by Order
    type TaxInfo struct {

        
            Gstin string  `json:"gstin"`
            B2bGstinNumber string  `json:"b2b_gstin_number"`
         
    }
    
    // ShippingInfo used by Order
    type ShippingInfo struct {

        
            FirstName string  `json:"first_name"`
            PrimaryMobileNumber string  `json:"primary_mobile_number"`
            CountryCode string  `json:"country_code"`
            CustomerCode string  `json:"customer_code"`
            FloorNo string  `json:"floor_no"`
            Landmark string  `json:"landmark"`
            Title string  `json:"title"`
            Country string  `json:"country"`
            City string  `json:"city"`
            Address1 string  `json:"address1"`
            AddressType string  `json:"address_type"`
            State string  `json:"state"`
            AlternateEmail string  `json:"alternate_email"`
            StateCode string  `json:"state_code"`
            ShippingType string  `json:"shipping_type"`
            GeoLocation map[string]interface{}  `json:"geo_location"`
            LastName string  `json:"last_name"`
            MiddleName string  `json:"middle_name"`
            PrimaryEmail string  `json:"primary_email"`
            Slot []map[string]interface{}  `json:"slot"`
            Gender string  `json:"gender"`
            ExternalCustomerCode string  `json:"external_customer_code"`
            AlternateMobileNumber string  `json:"alternate_mobile_number"`
            Pincode string  `json:"pincode"`
            HouseNo string  `json:"house_no"`
            Address2 string  `json:"address2"`
         
    }
    
    // Tax used by Order
    type Tax struct {

        
            Rate float64  `json:"rate"`
            Amount map[string]interface{}  `json:"amount"`
            Breakup []map[string]interface{}  `json:"breakup"`
            Name string  `json:"name"`
         
    }
    
    // Charge used by Order
    type Charge struct {

        
            Name string  `json:"name"`
            Tax Tax  `json:"tax"`
            Code string  `json:"code"`
            Amount map[string]interface{}  `json:"amount"`
            Type string  `json:"type"`
         
    }
    
    // PaymentMethod used by Order
    type PaymentMethod struct {

        
            CollectBy string  `json:"collect_by"`
            Name string  `json:"name"`
            Mode string  `json:"mode"`
            TransactionData map[string]interface{}  `json:"transaction_data"`
            Meta map[string]interface{}  `json:"meta"`
            RefundBy string  `json:"refund_by"`
            Amount float64  `json:"amount"`
         
    }
    
    // PaymentInfo used by Order
    type PaymentInfo struct {

        
            PaymentMethods []PaymentMethod  `json:"payment_methods"`
            PrimaryMode string  `json:"primary_mode"`
         
    }
    
    // BillingInfo used by Order
    type BillingInfo struct {

        
            FirstName string  `json:"first_name"`
            PrimaryMobileNumber string  `json:"primary_mobile_number"`
            CountryCode string  `json:"country_code"`
            CustomerCode string  `json:"customer_code"`
            FloorNo string  `json:"floor_no"`
            Title string  `json:"title"`
            Country string  `json:"country"`
            City string  `json:"city"`
            Address1 string  `json:"address1"`
            State string  `json:"state"`
            AlternateEmail string  `json:"alternate_email"`
            StateCode string  `json:"state_code"`
            LastName string  `json:"last_name"`
            MiddleName string  `json:"middle_name"`
            PrimaryEmail string  `json:"primary_email"`
            Gender string  `json:"gender"`
            ExternalCustomerCode string  `json:"external_customer_code"`
            AlternateMobileNumber string  `json:"alternate_mobile_number"`
            Pincode string  `json:"pincode"`
            HouseNo string  `json:"house_no"`
            Address2 string  `json:"address2"`
         
    }
    
    // ProcessingDates used by Order
    type ProcessingDates struct {

        
            PackByDate string  `json:"pack_by_date"`
            DispatchByDate string  `json:"dispatch_by_date"`
            CustomerPickupSlot map[string]interface{}  `json:"customer_pickup_slot"`
            DpPickupSlot map[string]interface{}  `json:"dp_pickup_slot"`
            ConfirmByDate string  `json:"confirm_by_date"`
            DispatchAfterDate string  `json:"dispatch_after_date"`
         
    }
    
    // LineItem used by Order
    type LineItem struct {

        
            CustomMessasge string  `json:"custom_messasge"`
            Charges []Charge  `json:"charges"`
            Quantity float64  `json:"quantity"`
            Meta map[string]interface{}  `json:"meta"`
            SellerIdentifier string  `json:"seller_identifier"`
            ExternalLineID string  `json:"external_line_id"`
         
    }
    
    // Shipment used by Order
    type Shipment struct {

        
            ExternalShipmentID string  `json:"external_shipment_id"`
            LocationID float64  `json:"location_id"`
            Priority float64  `json:"priority"`
            Meta map[string]interface{}  `json:"meta"`
            ProcessingDates ProcessingDates  `json:"processing_dates"`
            LineItems []LineItem  `json:"line_items"`
         
    }
    
    // CreateOrderAPI used by Order
    type CreateOrderAPI struct {

        
            CurrencyInfo map[string]interface{}  `json:"currency_info"`
            TaxInfo TaxInfo  `json:"tax_info"`
            ShippingInfo ShippingInfo  `json:"shipping_info"`
            Charges []Charge  `json:"charges"`
            PaymentInfo PaymentInfo  `json:"payment_info"`
            BillingInfo BillingInfo  `json:"billing_info"`
            ExternalCreationDate string  `json:"external_creation_date"`
            Shipments []Shipment  `json:"shipments"`
            Meta map[string]interface{}  `json:"meta"`
            Config map[string]interface{}  `json:"config"`
            ExternalOrderID string  `json:"external_order_id"`
         
    }
    
    // CreateOrderErrorReponse used by Order
    type CreateOrderErrorReponse struct {

        
            Message string  `json:"message"`
            Status float64  `json:"status"`
            Exception string  `json:"exception"`
            Info interface{}  `json:"info"`
            RequestID string  `json:"request_id"`
            StackTrace string  `json:"stack_trace"`
            Meta string  `json:"meta"`
            Code string  `json:"code"`
         
    }
    
    // PaymentMethods used by Order
    type PaymentMethods struct {

        
            RefundBy string  `json:"refund_by"`
            CollectBy string  `json:"collect_by"`
            Mode string  `json:"mode"`
         
    }
    
    // CreateChannelPaymentInfo used by Order
    type CreateChannelPaymentInfo struct {

        
            Source string  `json:"source"`
            PaymentMethods []PaymentMethods  `json:"payment_methods"`
            ModeOfPayment string  `json:"mode_of_payment"`
         
    }
    
    // DpConfiguration used by Order
    type DpConfiguration struct {

        
            ShippingBy string  `json:"shipping_by"`
         
    }
    
    // CreateChannelConfig used by Order
    type CreateChannelConfig struct {

        
            PaymentInfo CreateChannelPaymentInfo  `json:"payment_info"`
            ShipmentAssignment string  `json:"shipment_assignment"`
            DpConfiguration DpConfiguration  `json:"dp_configuration"`
            LockStates []string  `json:"lock_states"`
            LocationReassignment bool  `json:"location_reassignment"`
            LogoURL map[string]interface{}  `json:"logo_url"`
         
    }
    
    // CreateChannelConfigData used by Order
    type CreateChannelConfigData struct {

        
            ConfigData CreateChannelConfig  `json:"config_data"`
         
    }
    
    // CreateChannelConfigResponse used by Order
    type CreateChannelConfigResponse struct {

        
            Acknowledged bool  `json:"acknowledged"`
            IsUpserted bool  `json:"is_upserted"`
            IsInserted bool  `json:"is_inserted"`
         
    }
    
    // CreateChannelConifgErrorResponse used by Order
    type CreateChannelConifgErrorResponse struct {

        
            Error string  `json:"error"`
         
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

        
            EndDate string  `json:"end_date"`
            OrderDetails []FyndOrderIdList  `json:"order_details"`
            Mobile float64  `json:"mobile"`
            StartDate string  `json:"start_date"`
         
    }
    
    // BagStateTransitionMap used by Order
    type BagStateTransitionMap struct {

        
            Affiliate map[string]interface{}  `json:"affiliate"`
            Fynd map[string]interface{}  `json:"fynd"`
         
    }
    
    // FetchCreditBalanceRequestPayload used by Order
    type FetchCreditBalanceRequestPayload struct {

        
            SellerID string  `json:"seller_id"`
            AffiliateID string  `json:"affiliate_id"`
            CustomerMobileNumber string  `json:"customer_mobile_number"`
         
    }
    
    // CreditBalanceInfo used by Order
    type CreditBalanceInfo struct {

        
            Reason string  `json:"reason"`
            TotalCreditedBalance string  `json:"total_credited_balance"`
            CustomerMobileNumber string  `json:"customer_mobile_number"`
         
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
            AffiliateID string  `json:"affiliate_id"`
            OrderingChannel string  `json:"ordering_channel"`
            FyndOrderID string  `json:"fynd_order_id"`
         
    }
    
    // SingleRefundModeInfo used by Order
    type SingleRefundModeInfo struct {

        
            DisplayName string  `json:"display_name"`
            IsActive bool  `json:"is_active"`
            Slug string  `json:"slug"`
         
    }
    
    // RefundModeInfo used by Order
    type RefundModeInfo struct {

        
            RefundModeName SingleRefundModeInfo  `json:"refund_mode_name"`
         
    }
    
    // RefundModeConfigResponsePayload used by Order
    type RefundModeConfigResponsePayload struct {

        
            Data RefundModeInfo  `json:"data"`
            Success bool  `json:"success"`
         
    }
    
    // AttachUserOtpData used by Order
    type AttachUserOtpData struct {

        
            RequestID string  `json:"request_id"`
         
    }
    
    // AttachUserInfo used by Order
    type AttachUserInfo struct {

        
            FirstName string  `json:"first_name"`
            LastName string  `json:"last_name"`
            Mobile float64  `json:"mobile"`
            CountryCode string  `json:"country_code"`
         
    }
    
    // AttachOrderUser used by Order
    type AttachOrderUser struct {

        
            OtpData AttachUserOtpData  `json:"otp_data"`
            FyndOrderID string  `json:"fynd_order_id"`
            UserInfo AttachUserInfo  `json:"user_info"`
         
    }
    
    // AttachOrderUserResponse used by Order
    type AttachOrderUserResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // SendUserMobileOTP used by Order
    type SendUserMobileOTP struct {

        
            Mobile float64  `json:"mobile"`
            CountryCode string  `json:"country_code"`
         
    }
    
    // PointBlankOtpData used by Order
    type PointBlankOtpData struct {

        
            Message string  `json:"message"`
            ResendTimer float64  `json:"resend_timer"`
            Mobile float64  `json:"mobile"`
            RequestID string  `json:"request_id"`
         
    }
    
    // SendUserMobileOtpResponse used by Order
    type SendUserMobileOtpResponse struct {

        
            Message string  `json:"message"`
            Data PointBlankOtpData  `json:"data"`
            Success bool  `json:"success"`
         
    }
    
    // VerifyOtpData used by Order
    type VerifyOtpData struct {

        
            OtpCode float64  `json:"otp_code"`
            RequestID string  `json:"request_id"`
         
    }
    
    // VerifyMobileOTP used by Order
    type VerifyMobileOTP struct {

        
            OtpData VerifyOtpData  `json:"otp_data"`
         
    }
    
    // EInvoiceRetryShipmentData used by Order
    type EInvoiceRetryShipmentData struct {

        
            EinvoiceType string  `json:"einvoice_type"`
            ShipmentID string  `json:"shipment_id"`
         
    }
    
    // EInvoiceRetry used by Order
    type EInvoiceRetry struct {

        
            ShipmentsData []EInvoiceRetryShipmentData  `json:"shipments_data"`
         
    }
    
    // EinvoiceResponseData used by Order
    type EinvoiceResponseData struct {

        
            Message string  `json:"message"`
         
    }
    
    // EInvoiceRetryResponse used by Order
    type EInvoiceRetryResponse struct {

        
            ResponseData []EinvoiceResponseData  `json:"response_data"`
         
    }
    

    
    // DeleteResponse used by Catalog
    type DeleteResponse struct {

        
            Message string  `json:"message"`
         
    }
    
    // ErrorResponse used by Catalog
    type ErrorResponse struct {

        
            Error string  `json:"error"`
            Code string  `json:"code"`
            Message string  `json:"message"`
            Status float64  `json:"status"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // GetSearchWordsData used by Catalog
    type GetSearchWordsData struct {

        
            AppID string  `json:"app_id"`
            Result map[string]interface{}  `json:"result"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Words []string  `json:"words"`
            IsActive bool  `json:"is_active"`
            UID string  `json:"uid"`
         
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

        
            Query map[string]interface{}  `json:"query"`
            SortOn string  `json:"sort_on"`
         
    }
    
    // CreateSearchKeyword used by Catalog
    type CreateSearchKeyword struct {

        
            AppID string  `json:"app_id"`
            Result SearchKeywordResult  `json:"result"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Words []string  `json:"words"`
            IsActive bool  `json:"is_active"`
         
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
            Results []map[string]interface{}  `json:"results"`
            Words []string  `json:"words"`
            UID string  `json:"uid"`
         
    }
    
    // GetAutocompleteWordsResponse used by Catalog
    type GetAutocompleteWordsResponse struct {

        
            Items []GetAutocompleteWordsData  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // AutocompletePageAction used by Catalog
    type AutocompletePageAction struct {

        
            Query map[string]interface{}  `json:"query"`
            Type string  `json:"type"`
            URL string  `json:"url"`
            Params map[string]interface{}  `json:"params"`
         
    }
    
    // AutocompleteAction used by Catalog
    type AutocompleteAction struct {

        
            Type string  `json:"type"`
            Page AutocompletePageAction  `json:"page"`
         
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
            Action AutocompleteAction  `json:"action"`
            Logo Media  `json:"logo"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // CreateAutocompleteKeyword used by Catalog
    type CreateAutocompleteKeyword struct {

        
            AppID string  `json:"app_id"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Results []AutocompleteResult  `json:"results"`
            Words []string  `json:"words"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // CreateAutocompleteWordsResponse used by Catalog
    type CreateAutocompleteWordsResponse struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Results []map[string]interface{}  `json:"results"`
            AppID string  `json:"app_id"`
            Words []string  `json:"words"`
         
    }
    
    // ProductBundleItem used by Catalog
    type ProductBundleItem struct {

        
            AutoAddToCart bool  `json:"auto_add_to_cart"`
            MinQuantity float64  `json:"min_quantity"`
            AllowRemove bool  `json:"allow_remove"`
            AutoSelect bool  `json:"auto_select"`
            MaxQuantity float64  `json:"max_quantity"`
            ProductUID float64  `json:"product_uid"`
         
    }
    
    // ProductBundleRequest used by Catalog
    type ProductBundleRequest struct {

        
            Slug string  `json:"slug"`
            PageVisibility []string  `json:"page_visibility"`
            CompanyID float64  `json:"company_id"`
            Products []ProductBundleItem  `json:"products"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            ModifiedOn string  `json:"modified_on"`
            Logo string  `json:"logo"`
            Choice string  `json:"choice"`
            CreatedOn string  `json:"created_on"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            Name string  `json:"name"`
            IsActive bool  `json:"is_active"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // GetProductBundleCreateResponse used by Catalog
    type GetProductBundleCreateResponse struct {

        
            Slug string  `json:"slug"`
            PageVisibility []string  `json:"page_visibility"`
            CompanyID float64  `json:"company_id"`
            Products []ProductBundleItem  `json:"products"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            ModifiedOn string  `json:"modified_on"`
            Logo string  `json:"logo"`
            Choice string  `json:"choice"`
            ID string  `json:"id"`
            CreatedOn string  `json:"created_on"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            Name string  `json:"name"`
            IsActive bool  `json:"is_active"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // GetProductBundleListingResponse used by Catalog
    type GetProductBundleListingResponse struct {

        
            Items []GetProductBundleCreateResponse  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // Size used by Catalog
    type Size struct {

        
            Quantity float64  `json:"quantity"`
            Display string  `json:"display"`
            IsAvailable bool  `json:"is_available"`
            Value string  `json:"value"`
         
    }
    
    // LimitedProductData used by Catalog
    type LimitedProductData struct {

        
            Slug string  `json:"slug"`
            ItemCode string  `json:"item_code"`
            Price map[string]interface{}  `json:"price"`
            Sizes []string  `json:"sizes"`
            Attributes map[string]interface{}  `json:"attributes"`
            Quantity float64  `json:"quantity"`
            Name string  `json:"name"`
            ShortDescription string  `json:"short_description"`
            Images []string  `json:"images"`
            Identifier map[string]interface{}  `json:"identifier"`
            CountryOfOrigin string  `json:"country_of_origin"`
            UID float64  `json:"uid"`
         
    }
    
    // Price used by Catalog
    type Price struct {

        
            Currency string  `json:"currency"`
            MinMarked float64  `json:"min_marked"`
            MaxEffective float64  `json:"max_effective"`
            MaxMarked float64  `json:"max_marked"`
            MinEffective float64  `json:"min_effective"`
         
    }
    
    // GetProducts used by Catalog
    type GetProducts struct {

        
            AutoAddToCart bool  `json:"auto_add_to_cart"`
            Sizes []Size  `json:"sizes"`
            MinQuantity float64  `json:"min_quantity"`
            AllowRemove bool  `json:"allow_remove"`
            ProductDetails LimitedProductData  `json:"product_details"`
            AutoSelect bool  `json:"auto_select"`
            MaxQuantity float64  `json:"max_quantity"`
            ProductUID float64  `json:"product_uid"`
            Price Price  `json:"price"`
         
    }
    
    // GetProductBundleResponse used by Catalog
    type GetProductBundleResponse struct {

        
            Slug string  `json:"slug"`
            PageVisibility []string  `json:"page_visibility"`
            CompanyID float64  `json:"company_id"`
            Products []GetProducts  `json:"products"`
            Logo string  `json:"logo"`
            Choice string  `json:"choice"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            Name string  `json:"name"`
            IsActive bool  `json:"is_active"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // ProductBundleUpdateRequest used by Catalog
    type ProductBundleUpdateRequest struct {

        
            Slug string  `json:"slug"`
            PageVisibility []string  `json:"page_visibility"`
            CompanyID float64  `json:"company_id"`
            Products []ProductBundleItem  `json:"products"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            ModifiedOn string  `json:"modified_on"`
            Logo string  `json:"logo"`
            Choice string  `json:"choice"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            Name string  `json:"name"`
            IsActive bool  `json:"is_active"`
            Meta map[string]interface{}  `json:"meta"`
         
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

        
            Title string  `json:"title"`
            Description string  `json:"description"`
            Guide Guide  `json:"guide"`
            CompanyID float64  `json:"company_id"`
            Subtitle string  `json:"subtitle"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Active bool  `json:"active"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            ModifiedOn string  `json:"modified_on"`
            BrandID float64  `json:"brand_id"`
            CreatedOn string  `json:"created_on"`
            ID string  `json:"id"`
            Image string  `json:"image"`
            Name string  `json:"name"`
            Tag string  `json:"tag"`
         
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

        
            Title string  `json:"title"`
            CompanyID float64  `json:"company_id"`
            Subtitle string  `json:"subtitle"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Name string  `json:"name"`
            Active bool  `json:"active"`
            ModifiedOn string  `json:"modified_on"`
            BrandID float64  `json:"brand_id"`
            ID string  `json:"id"`
            CreatedOn string  `json:"created_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Tag string  `json:"tag"`
            Guide map[string]interface{}  `json:"guide"`
         
    }
    
    // MOQData used by Catalog
    type MOQData struct {

        
            Minimum float64  `json:"minimum"`
            Maximum float64  `json:"maximum"`
            IncrementUnit float64  `json:"increment_unit"`
         
    }
    
    // SEOData used by Catalog
    type SEOData struct {

        
            Title string  `json:"title"`
            Description string  `json:"description"`
         
    }
    
    // OwnerAppItemResponse used by Catalog
    type OwnerAppItemResponse struct {

        
            IsCod bool  `json:"is_cod"`
            AltText map[string]interface{}  `json:"alt_text"`
            Moq MOQData  `json:"moq"`
            IsGift bool  `json:"is_gift"`
            Seo SEOData  `json:"seo"`
         
    }
    
    // MetaFields used by Catalog
    type MetaFields struct {

        
            Key string  `json:"key"`
            Value string  `json:"value"`
         
    }
    
    // ApplicationItemMOQ used by Catalog
    type ApplicationItemMOQ struct {

        
            Minimum float64  `json:"minimum"`
            Maximum float64  `json:"maximum"`
            IncrementUnit float64  `json:"increment_unit"`
         
    }
    
    // ApplicationItemSEO used by Catalog
    type ApplicationItemSEO struct {

        
            Title string  `json:"title"`
            Description string  `json:"description"`
         
    }
    
    // ApplicationItemMeta used by Catalog
    type ApplicationItemMeta struct {

        
            IsCod bool  `json:"is_cod"`
            AltText map[string]interface{}  `json:"alt_text"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            CustomMeta []MetaFields  `json:"_custom_meta"`
            Moq ApplicationItemMOQ  `json:"moq"`
            IsGift bool  `json:"is_gift"`
            Seo ApplicationItemSEO  `json:"seo"`
         
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
    
    // AttributeDetailsGroup used by Catalog
    type AttributeDetailsGroup struct {

        
            Slug string  `json:"slug"`
            DisplayType string  `json:"display_type"`
            Priority float64  `json:"priority"`
            Unit string  `json:"unit"`
            Key string  `json:"key"`
            Logo string  `json:"logo"`
            Name string  `json:"name"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // AppConfigurationDetail used by Catalog
    type AppConfigurationDetail struct {

        
            Slug string  `json:"slug"`
            TemplateSlugs []string  `json:"template_slugs"`
            AppID string  `json:"app_id"`
            Attributes []AttributeDetailsGroup  `json:"attributes"`
            Priority float64  `json:"priority"`
            Logo string  `json:"logo"`
            IsDefault bool  `json:"is_default"`
            Name string  `json:"name"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // ConfigErrorResponse used by Catalog
    type ConfigErrorResponse struct {

        
            Message string  `json:"message"`
         
    }
    
    // PageResponseType used by Catalog
    type PageResponseType struct {

        
            Next float64  `json:"next"`
            TotalCount float64  `json:"total_count"`
            Current float64  `json:"current"`
            HasNext bool  `json:"has_next"`
         
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

        
            DefaultKey string  `json:"default_key"`
            AppID string  `json:"app_id"`
            Priority float64  `json:"priority"`
            Key string  `json:"key"`
            Logo string  `json:"logo"`
            IsDefault bool  `json:"is_default"`
            Name string  `json:"name"`
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
            Similar map[string]interface{}  `json:"similar"`
            Compare map[string]interface{}  `json:"compare"`
            Detail map[string]interface{}  `json:"detail"`
         
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
    
    // ConfigurationProductConfig used by Catalog
    type ConfigurationProductConfig struct {

        
            Title string  `json:"title"`
            Subtitle string  `json:"subtitle"`
            Size ProductSize  `json:"size"`
            Priority float64  `json:"priority"`
            Key string  `json:"key"`
            Logo string  `json:"logo"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // ConfigurationProductSimilar used by Catalog
    type ConfigurationProductSimilar struct {

        
            Config []ConfigurationProductConfig  `json:"config"`
         
    }
    
    // ConfigurationProductVariantConfig used by Catalog
    type ConfigurationProductVariantConfig struct {

        
            Size ProductSize  `json:"size"`
            DisplayType string  `json:"display_type"`
            Priority float64  `json:"priority"`
            Key string  `json:"key"`
            Logo string  `json:"logo"`
            Name string  `json:"name"`
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
    
    // ConfigurationBucketPoints used by Catalog
    type ConfigurationBucketPoints struct {

        
            Display string  `json:"display"`
            Start float64  `json:"start"`
            End float64  `json:"end"`
         
    }
    
    // ConfigurationListingFilterValue used by Catalog
    type ConfigurationListingFilterValue struct {

        
            MapValues []map[string]interface{}  `json:"map_values"`
            BucketPoints []ConfigurationBucketPoints  `json:"bucket_points"`
            Condition string  `json:"condition"`
            Sort string  `json:"sort"`
            Map map[string]interface{}  `json:"map"`
            Value string  `json:"value"`
         
    }
    
    // ConfigurationListingFilterConfig used by Catalog
    type ConfigurationListingFilterConfig struct {

        
            DisplayName string  `json:"display_name"`
            ValueConfig ConfigurationListingFilterValue  `json:"value_config"`
            Type string  `json:"type"`
            Priority float64  `json:"priority"`
            Key string  `json:"key"`
            Logo string  `json:"logo"`
            Name string  `json:"name"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // ConfigurationListingFilter used by Catalog
    type ConfigurationListingFilter struct {

        
            AttributeConfig []ConfigurationListingFilterConfig  `json:"attribute_config"`
            AllowSingle bool  `json:"allow_single"`
         
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

        
            Config []ConfigurationListingSortConfig  `json:"config"`
            DefaultKey string  `json:"default_key"`
         
    }
    
    // ConfigurationListing used by Catalog
    type ConfigurationListing struct {

        
            Filter ConfigurationListingFilter  `json:"filter"`
            Sort ConfigurationListingSort  `json:"sort"`
         
    }
    
    // AppConfiguration used by Catalog
    type AppConfiguration struct {

        
            Product ConfigurationProduct  `json:"product"`
            ConfigID string  `json:"config_id"`
            AppID string  `json:"app_id"`
            Type string  `json:"type"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
            ConfigType string  `json:"config_type"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Listing ConfigurationListing  `json:"listing"`
         
    }
    
    // AppCatalogConfiguration used by Catalog
    type AppCatalogConfiguration struct {

        
            Product ConfigurationProduct  `json:"product"`
            ConfigID string  `json:"config_id"`
            AppID string  `json:"app_id"`
            Type string  `json:"type"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ModifiedOn string  `json:"modified_on"`
            ID string  `json:"id"`
            CreatedOn string  `json:"created_on"`
            ConfigType string  `json:"config_type"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Listing ConfigurationListing  `json:"listing"`
         
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

        
            Product GetCatalogConfigurationDetailsProduct  `json:"product"`
            ConfigID string  `json:"config_id"`
            AppID string  `json:"app_id"`
            ID string  `json:"id"`
            ConfigType string  `json:"config_type"`
            Listing GetCatalogConfigurationDetailsSchemaListing  `json:"listing"`
         
    }
    
    // GetAppCatalogEntityConfiguration used by Catalog
    type GetAppCatalogEntityConfiguration struct {

        
            Data EntityConfiguration  `json:"data"`
            IsDefault bool  `json:"is_default"`
         
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

        
            SelectedMin float64  `json:"selected_min"`
            Max float64  `json:"max"`
            CurrencySymbol string  `json:"currency_symbol"`
            QueryFormat string  `json:"query_format"`
            Display string  `json:"display"`
            Count float64  `json:"count"`
            DisplayFormat string  `json:"display_format"`
            Min float64  `json:"min"`
            CurrencyCode string  `json:"currency_code"`
            SelectedMax float64  `json:"selected_max"`
            IsSelected bool  `json:"is_selected"`
            Value interface{}  `json:"value"`
         
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

        
            Operators map[string]string  `json:"operators"`
            Filters []ProductFilters  `json:"filters"`
            SortOn []ProductSortOn  `json:"sort_on"`
         
    }
    
    // CollectionQuery used by Catalog
    type CollectionQuery struct {

        
            Op string  `json:"op"`
            Attribute string  `json:"attribute"`
            Value []interface{}  `json:"value"`
         
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

        
            Email string  `json:"email"`
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
            UID string  `json:"uid"`
         
    }
    
    // CollectionBadge used by Catalog
    type CollectionBadge struct {

        
            Color string  `json:"color"`
            Text string  `json:"text"`
         
    }
    
    // NextSchedule used by Catalog
    type NextSchedule struct {

        
            Start string  `json:"start"`
            End string  `json:"end"`
         
    }
    
    // CollectionSchedule used by Catalog
    type CollectionSchedule struct {

        
            NextSchedule []NextSchedule  `json:"next_schedule"`
            Duration float64  `json:"duration"`
            Start string  `json:"start"`
            Cron string  `json:"cron"`
            End string  `json:"end"`
         
    }
    
    // SeoDetail used by Catalog
    type SeoDetail struct {

        
            Title string  `json:"title"`
            Description string  `json:"description"`
         
    }
    
    // CreateCollection used by Catalog
    type CreateCollection struct {

        
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Query []CollectionQuery  `json:"query"`
            Type string  `json:"type"`
            IsVisible bool  `json:"is_visible"`
            SortOn string  `json:"sort_on"`
            Name string  `json:"name"`
            Banners CollectionBanner  `json:"banners"`
            AllowFacets bool  `json:"allow_facets"`
            Logo CollectionImage  `json:"logo"`
            ModifiedBy UserInfo  `json:"modified_by"`
            Meta map[string]interface{}  `json:"meta"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            AllowSort bool  `json:"allow_sort"`
            Priority float64  `json:"priority"`
            Published bool  `json:"published"`
            IsActive bool  `json:"is_active"`
            Badge CollectionBadge  `json:"badge"`
            Slug string  `json:"slug"`
            Description string  `json:"description"`
            AppID string  `json:"app_id"`
            Tags []string  `json:"tags"`
            Schedule CollectionSchedule  `json:"_schedule"`
            CreatedBy UserInfo  `json:"created_by"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Seo SeoDetail  `json:"seo"`
         
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
    
    // Media1 used by Catalog
    type Media1 struct {

        
            Type string  `json:"type"`
            URL string  `json:"url"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // CollectionCreateResponse used by Catalog
    type CollectionCreateResponse struct {

        
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Query []CollectionQuery  `json:"query"`
            Type string  `json:"type"`
            SortOn string  `json:"sort_on"`
            Name string  `json:"name"`
            Banners ImageUrls  `json:"banners"`
            AllowFacets bool  `json:"allow_facets"`
            Logo Media1  `json:"logo"`
            Cron map[string]interface{}  `json:"cron"`
            Meta map[string]interface{}  `json:"meta"`
            AllowSort bool  `json:"allow_sort"`
            Priority float64  `json:"priority"`
            IsActive bool  `json:"is_active"`
            Badge map[string]interface{}  `json:"badge"`
            Description string  `json:"description"`
            Slug string  `json:"slug"`
            AppID string  `json:"app_id"`
            Tags []string  `json:"tags"`
            Schedule map[string]interface{}  `json:"_schedule"`
         
    }
    
    // CollectionListingFilterType used by Catalog
    type CollectionListingFilterType struct {

        
            Display string  `json:"display"`
            Name string  `json:"name"`
            IsSelected bool  `json:"is_selected"`
         
    }
    
    // CollectionListingFilterTag used by Catalog
    type CollectionListingFilterTag struct {

        
            Display string  `json:"display"`
            Name string  `json:"name"`
            IsSelected bool  `json:"is_selected"`
         
    }
    
    // CollectionListingFilter used by Catalog
    type CollectionListingFilter struct {

        
            Type []CollectionListingFilterType  `json:"type"`
            Tags []CollectionListingFilterTag  `json:"tags"`
         
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

        
            Action Action  `json:"action"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Query []CollectionQuery  `json:"query"`
            Type string  `json:"type"`
            SortOn string  `json:"sort_on"`
            Name string  `json:"name"`
            Banners ImageUrls  `json:"banners"`
            UID string  `json:"uid"`
            AllowFacets bool  `json:"allow_facets"`
            Logo Media1  `json:"logo"`
            Cron map[string]interface{}  `json:"cron"`
            Meta map[string]interface{}  `json:"meta"`
            AllowSort bool  `json:"allow_sort"`
            Priority float64  `json:"priority"`
            Tag []string  `json:"tag"`
            IsActive bool  `json:"is_active"`
            Badge map[string]interface{}  `json:"badge"`
            Description string  `json:"description"`
            Slug string  `json:"slug"`
            AppID string  `json:"app_id"`
            Schedule map[string]interface{}  `json:"_schedule"`
         
    }
    
    // GetCollectionListingResponse used by Catalog
    type GetCollectionListingResponse struct {

        
            Filters CollectionListingFilter  `json:"filters"`
            Items []GetCollectionDetailNest  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // CollectionDetailResponse used by Catalog
    type CollectionDetailResponse struct {

        
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Query []CollectionQuery  `json:"query"`
            Type string  `json:"type"`
            SortOn string  `json:"sort_on"`
            Name string  `json:"name"`
            Banners ImageUrls  `json:"banners"`
            UID string  `json:"uid"`
            AllowFacets bool  `json:"allow_facets"`
            Logo Media1  `json:"logo"`
            Cron map[string]interface{}  `json:"cron"`
            Meta map[string]interface{}  `json:"meta"`
            AllowSort bool  `json:"allow_sort"`
            Priority float64  `json:"priority"`
            Tag []string  `json:"tag"`
            IsActive bool  `json:"is_active"`
            Badge map[string]interface{}  `json:"badge"`
            Description string  `json:"description"`
            Slug string  `json:"slug"`
            AppID string  `json:"app_id"`
            Schedule map[string]interface{}  `json:"_schedule"`
         
    }
    
    // UpdateCollection used by Catalog
    type UpdateCollection struct {

        
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Query []CollectionQuery  `json:"query"`
            Type string  `json:"type"`
            IsVisible bool  `json:"is_visible"`
            SortOn string  `json:"sort_on"`
            Name string  `json:"name"`
            Banners CollectionBanner  `json:"banners"`
            AllowFacets bool  `json:"allow_facets"`
            Logo CollectionImage  `json:"logo"`
            ModifiedBy UserInfo  `json:"modified_by"`
            Meta map[string]interface{}  `json:"meta"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            AllowSort bool  `json:"allow_sort"`
            Priority float64  `json:"priority"`
            Published bool  `json:"published"`
            IsActive bool  `json:"is_active"`
            Badge CollectionBadge  `json:"badge"`
            Slug string  `json:"slug"`
            Description string  `json:"description"`
            Tags []string  `json:"tags"`
            Schedule CollectionSchedule  `json:"_schedule"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Seo SeoDetail  `json:"seo"`
         
    }
    
    // CollectionUpdateResponse used by Catalog
    type CollectionUpdateResponse struct {

        
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Query []CollectionQuery  `json:"query"`
            Type string  `json:"type"`
            IsVisible bool  `json:"is_visible"`
            SortOn string  `json:"sort_on"`
            Name string  `json:"name"`
            Banners CollectionBanner  `json:"banners"`
            AllowFacets bool  `json:"allow_facets"`
            Logo CollectionImage  `json:"logo"`
            ModifiedBy UserInfo  `json:"modified_by"`
            Meta map[string]interface{}  `json:"meta"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            AllowSort bool  `json:"allow_sort"`
            Priority float64  `json:"priority"`
            Published bool  `json:"published"`
            IsActive bool  `json:"is_active"`
            Badge CollectionBadge  `json:"badge"`
            Slug string  `json:"slug"`
            Description string  `json:"description"`
            Tags []string  `json:"tags"`
            AppID string  `json:"app_id"`
            Schedule CollectionSchedule  `json:"_schedule"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Seo SeoDetail  `json:"seo"`
         
    }
    
    // CollectionItem used by Catalog
    type CollectionItem struct {

        
            Action string  `json:"action"`
            ItemID float64  `json:"item_id"`
            Priority float64  `json:"priority"`
         
    }
    
    // CollectionItemUpdate used by Catalog
    type CollectionItemUpdate struct {

        
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Items []CollectionItem  `json:"items"`
            Query []CollectionQuery  `json:"query"`
            Type string  `json:"type"`
            AllowSort bool  `json:"allow_sort"`
            AllowFacets bool  `json:"allow_facets"`
         
    }
    
    // UpdatedResponse used by Catalog
    type UpdatedResponse struct {

        
            Message string  `json:"message"`
            ItemsNotUpdated []float64  `json:"items_not_updated"`
         
    }
    
    // ProductBrand used by Catalog
    type ProductBrand struct {

        
            Logo Media1  `json:"logo"`
            Action Action  `json:"action"`
            Name string  `json:"name"`
            UID float64  `json:"uid"`
         
    }
    
    // ProductDetailAttribute used by Catalog
    type ProductDetailAttribute struct {

        
            Key string  `json:"key"`
            Type string  `json:"type"`
            Value string  `json:"value"`
         
    }
    
    // ProductDetailGroupedAttribute used by Catalog
    type ProductDetailGroupedAttribute struct {

        
            Title string  `json:"title"`
            Details []ProductDetailAttribute  `json:"details"`
         
    }
    
    // Price1 used by Catalog
    type Price1 struct {

        
            Min float64  `json:"min"`
            CurrencyCode string  `json:"currency_code"`
            Max float64  `json:"max"`
            CurrencySymbol string  `json:"currency_symbol"`
         
    }
    
    // ProductListingPrice used by Catalog
    type ProductListingPrice struct {

        
            Marked Price1  `json:"marked"`
            Effective Price1  `json:"effective"`
         
    }
    
    // ProductListingDetail used by Catalog
    type ProductListingDetail struct {

        
            ItemCode string  `json:"item_code"`
            ImageNature string  `json:"image_nature"`
            Type string  `json:"type"`
            Name string  `json:"name"`
            Sellable bool  `json:"sellable"`
            HasVariant bool  `json:"has_variant"`
            UID float64  `json:"uid"`
            Medias []Media1  `json:"medias"`
            RatingCount float64  `json:"rating_count"`
            Brand ProductBrand  `json:"brand"`
            PromoMeta map[string]interface{}  `json:"promo_meta"`
            Similars []string  `json:"similars"`
            Color string  `json:"color"`
            ShortDescription string  `json:"short_description"`
            Highlights []string  `json:"highlights"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            Rating float64  `json:"rating"`
            Slug string  `json:"slug"`
            Description string  `json:"description"`
            Discount string  `json:"discount"`
            Attributes map[string]interface{}  `json:"attributes"`
            ProductOnlineDate string  `json:"product_online_date"`
            ItemType string  `json:"item_type"`
            Tryouts []string  `json:"tryouts"`
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
            Price ProductListingPrice  `json:"price"`
         
    }
    
    // GetCollectionItemsResponse used by Catalog
    type GetCollectionItemsResponse struct {

        
            Items []ProductListingDetail  `json:"items"`
            Filters []ProductFilters  `json:"filters"`
            Page Page  `json:"page"`
            SortOn []ProductSortOn  `json:"sort_on"`
         
    }
    
    // CatalogInsightBrand used by Catalog
    type CatalogInsightBrand struct {

        
            TotalSizes float64  `json:"total_sizes"`
            AvailableSizes float64  `json:"available_sizes"`
            TotalArticles float64  `json:"total_articles"`
            ArticleFreshness float64  `json:"article_freshness"`
            AvailableArticles float64  `json:"available_articles"`
            Name string  `json:"name"`
         
    }
    
    // CatalogInsightItem used by Catalog
    type CatalogInsightItem struct {

        
            OutOfStockCount float64  `json:"out_of_stock_count"`
            Count float64  `json:"count"`
            SellableCount float64  `json:"sellable_count"`
         
    }
    
    // CatalogInsightResponse used by Catalog
    type CatalogInsightResponse struct {

        
            BrandDistribution CatalogInsightBrand  `json:"brand_distribution"`
            Item CatalogInsightItem  `json:"item"`
         
    }
    
    // CrossSellingDataResponse used by Catalog
    type CrossSellingDataResponse struct {

        
            Products float64  `json:"products"`
            Articles float64  `json:"articles"`
         
    }
    
    // OptInPostRequest used by Catalog
    type OptInPostRequest struct {

        
            CompanyID float64  `json:"company_id"`
            StoreIds []float64  `json:"store_ids"`
            Enabled bool  `json:"enabled"`
            OptLevel string  `json:"opt_level"`
            Platform string  `json:"platform"`
            BrandIds []float64  `json:"brand_ids"`
         
    }
    
    // CompanyOptIn used by Catalog
    type CompanyOptIn struct {

        
            CompanyID float64  `json:"company_id"`
            StoreIds []float64  `json:"store_ids"`
            Enabled bool  `json:"enabled"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ModifiedOn float64  `json:"modified_on"`
            OptLevel string  `json:"opt_level"`
            Platform string  `json:"platform"`
            CreatedOn float64  `json:"created_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            BrandIds []float64  `json:"brand_ids"`
         
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

        
            TotalArticle float64  `json:"total_article"`
            BrandName string  `json:"brand_name"`
            BrandID float64  `json:"brand_id"`
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

        
            StoreCode string  `json:"store_code"`
            DisplayName string  `json:"display_name"`
            AdditionalContacts []map[string]interface{}  `json:"additional_contacts"`
            Timing map[string]interface{}  `json:"timing"`
            CompanyID float64  `json:"company_id"`
            Documents []map[string]interface{}  `json:"documents"`
            Address map[string]interface{}  `json:"address"`
            StoreType string  `json:"store_type"`
            ModifiedOn string  `json:"modified_on"`
            Manager map[string]interface{}  `json:"manager"`
            CreatedOn string  `json:"created_on"`
            Name string  `json:"name"`
            UID float64  `json:"uid"`
         
    }
    
    // OptinStoreDetails used by Catalog
    type OptinStoreDetails struct {

        
            Items []StoreDetail  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // AttributeMasterDetails used by Catalog
    type AttributeMasterDetails struct {

        
            DisplayType string  `json:"display_type"`
         
    }
    
    // AttributeMasterFilter used by Catalog
    type AttributeMasterFilter struct {

        
            Indexing bool  `json:"indexing"`
            DependsOn []string  `json:"depends_on"`
            Priority float64  `json:"priority"`
         
    }
    
    // AttributeSchemaRange used by Catalog
    type AttributeSchemaRange struct {

        
            Min float64  `json:"min"`
            Max float64  `json:"max"`
         
    }
    
    // AttributeMaster used by Catalog
    type AttributeMaster struct {

        
            Multi bool  `json:"multi"`
            Range AttributeSchemaRange  `json:"range"`
            Type string  `json:"type"`
            Format string  `json:"format"`
            Mandatory bool  `json:"mandatory"`
            AllowedValues []string  `json:"allowed_values"`
         
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

        
            Description string  `json:"description"`
            Slug string  `json:"slug"`
            Details AttributeMasterDetails  `json:"details"`
            IsNested bool  `json:"is_nested"`
            Departments []string  `json:"departments"`
            Logo string  `json:"logo"`
            Filters AttributeMasterFilter  `json:"filters"`
            ID string  `json:"id"`
            Schema AttributeMaster  `json:"schema"`
            Name string  `json:"name"`
            EnabledForEndConsumer bool  `json:"enabled_for_end_consumer"`
            Meta AttributeMasterMeta  `json:"meta"`
         
    }
    
    // CategoriesResponse used by Catalog
    type CategoriesResponse struct {

        
            Slug string  `json:"slug"`
            TemplateSlug string  `json:"template_slug"`
            SlugKey string  `json:"slug_key"`
            Name string  `json:"name"`
            UID float64  `json:"uid"`
         
    }
    
    // ProdcutTemplateCategoriesResponse used by Catalog
    type ProdcutTemplateCategoriesResponse struct {

        
            Items []CategoriesResponse  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // PTErrorResponse used by Catalog
    type PTErrorResponse struct {

        
            Errors map[string]interface{}  `json:"errors"`
            Code string  `json:"code"`
            Message string  `json:"message"`
            Status float64  `json:"status"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // UserDetail used by Catalog
    type UserDetail struct {

        
            UserID string  `json:"user_id"`
            Contact string  `json:"contact"`
            SuperUser bool  `json:"super_user"`
            Username string  `json:"username"`
         
    }
    
    // DepartmentModel used by Catalog
    type DepartmentModel struct {

        
            Slug string  `json:"slug"`
            Synonyms []string  `json:"synonyms"`
            Cls string  `json:"_cls"`
            VerifiedBy UserDetail  `json:"verified_by"`
            IsActive bool  `json:"is_active"`
            CreatedBy UserDetail  `json:"created_by"`
            PriorityOrder float64  `json:"priority_order"`
            Name string  `json:"name"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ModifiedOn string  `json:"modified_on"`
            Logo string  `json:"logo"`
            ID interface{}  `json:"_id"`
            CreatedOn string  `json:"created_on"`
            ModifiedBy UserDetail  `json:"modified_by"`
            VerifiedOn string  `json:"verified_on"`
            UID float64  `json:"uid"`
         
    }
    
    // DepartmentPageResponse used by Catalog
    type DepartmentPageResponse struct {

        
            Type string  `json:"type"`
            Size float64  `json:"size"`
            ItemTotal float64  `json:"item_total"`
            HasPrevious bool  `json:"has_previous"`
            Current float64  `json:"current"`
            NextID string  `json:"next_id"`
            HasNext bool  `json:"has_next"`
         
    }
    
    // DepartmentListResponse used by Catalog
    type DepartmentListResponse struct {

        
            Items []DepartmentModel  `json:"items"`
            Page DepartmentPageResponse  `json:"page"`
         
    }
    
    // DepartmentCreateError used by Catalog
    type DepartmentCreateError struct {

        
            Error string  `json:"error"`
         
    }
    
    // DepartmentUpdate used by Catalog
    type DepartmentUpdate struct {

        
            Synonyms []string  `json:"synonyms"`
            Cls string  `json:"_cls"`
            Slug string  `json:"slug"`
            Tags []string  `json:"tags"`
            PriorityOrder float64  `json:"priority_order"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Logo string  `json:"logo"`
            Platform string  `json:"platform"`
            Name string  `json:"name"`
            IsActive bool  `json:"is_active"`
            UID float64  `json:"uid"`
         
    }
    
    // ProductTemplate used by Catalog
    type ProductTemplate struct {

        
            Slug string  `json:"slug"`
            Description string  `json:"description"`
            Categories []string  `json:"categories"`
            Attributes []string  `json:"attributes"`
            IsArchived bool  `json:"is_archived"`
            IsExpirable bool  `json:"is_expirable"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            ModifiedOn string  `json:"modified_on"`
            Departments []string  `json:"departments"`
            Logo string  `json:"logo"`
            CreatedOn string  `json:"created_on"`
            Name string  `json:"name"`
            Tag string  `json:"tag"`
            IsActive bool  `json:"is_active"`
            IsPhysical bool  `json:"is_physical"`
         
    }
    
    // TemplatesResponse used by Catalog
    type TemplatesResponse struct {

        
            Items ProductTemplate  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // Properties used by Catalog
    type Properties struct {

        
            ItemCode map[string]interface{}  `json:"item_code"`
            TraderType map[string]interface{}  `json:"trader_type"`
            ReturnConfig map[string]interface{}  `json:"return_config"`
            Media map[string]interface{}  `json:"media"`
            BrandUID map[string]interface{}  `json:"brand_uid"`
            Name map[string]interface{}  `json:"name"`
            CountryOfOrigin map[string]interface{}  `json:"country_of_origin"`
            MultiSize map[string]interface{}  `json:"multi_size"`
            HsnCode map[string]interface{}  `json:"hsn_code"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            ProductGroupTag map[string]interface{}  `json:"product_group_tag"`
            CategorySlug map[string]interface{}  `json:"category_slug"`
            IsDependent map[string]interface{}  `json:"is_dependent"`
            SizeGuide map[string]interface{}  `json:"size_guide"`
            Command map[string]interface{}  `json:"command"`
            Sizes map[string]interface{}  `json:"sizes"`
            NoOfBoxes map[string]interface{}  `json:"no_of_boxes"`
            Currency map[string]interface{}  `json:"currency"`
            ProductPublish map[string]interface{}  `json:"product_publish"`
            ShortDescription map[string]interface{}  `json:"short_description"`
            Variants map[string]interface{}  `json:"variants"`
            Highlights map[string]interface{}  `json:"highlights"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            Trader map[string]interface{}  `json:"trader"`
            IsActive map[string]interface{}  `json:"is_active"`
            Slug map[string]interface{}  `json:"slug"`
            Description map[string]interface{}  `json:"description"`
            Tags map[string]interface{}  `json:"tags"`
            ItemType map[string]interface{}  `json:"item_type"`
         
    }
    
    // GlobalValidation used by Catalog
    type GlobalValidation struct {

        
            Title string  `json:"title"`
            Description string  `json:"description"`
            Properties Properties  `json:"properties"`
            Definitions map[string]interface{}  `json:"definitions"`
            Type string  `json:"type"`
            Required []string  `json:"required"`
         
    }
    
    // TemplateValidationData used by Catalog
    type TemplateValidationData struct {

        
            GlobalValidation GlobalValidation  `json:"global_validation"`
            TemplateValidation map[string]interface{}  `json:"template_validation"`
         
    }
    
    // TemplateDetails used by Catalog
    type TemplateDetails struct {

        
            Slug string  `json:"slug"`
            Description string  `json:"description"`
            Categories []string  `json:"categories"`
            Attributes []string  `json:"attributes"`
            IsArchived bool  `json:"is_archived"`
            IsExpirable bool  `json:"is_expirable"`
            Departments []string  `json:"departments"`
            Logo string  `json:"logo"`
            ID string  `json:"id"`
            Name string  `json:"name"`
            Tag string  `json:"tag"`
            IsActive bool  `json:"is_active"`
            IsPhysical bool  `json:"is_physical"`
         
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
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
            UID string  `json:"uid"`
         
    }
    
    // ProductTemplateExportResponse used by Catalog
    type ProductTemplateExportResponse struct {

        
            SellerID float64  `json:"seller_id"`
            CompletedOn string  `json:"completed_on"`
            Type string  `json:"type"`
            CreatedBy UserInfo1  `json:"created_by"`
            URL string  `json:"url"`
            ModifiedOn string  `json:"modified_on"`
            TaskID string  `json:"task_id"`
            Filters map[string]interface{}  `json:"filters"`
            NotificationEmails []string  `json:"notification_emails"`
            Status string  `json:"status"`
         
    }
    
    // ProductDownloadsResponse used by Catalog
    type ProductDownloadsResponse struct {

        
            Items []ProductTemplateExportResponse  `json:"items"`
         
    }
    
    // ProductTemplateExportFilterRequest used by Catalog
    type ProductTemplateExportFilterRequest struct {

        
            Brands []string  `json:"brands"`
            Templates []string  `json:"templates"`
            CatalogueTypes []string  `json:"catalogue_types"`
            ToDate string  `json:"to_date"`
            FromDate string  `json:"from_date"`
         
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
    
    // Hierarchy used by Catalog
    type Hierarchy struct {

        
            L2 float64  `json:"l2"`
            L1 float64  `json:"l1"`
            Department float64  `json:"department"`
         
    }
    
    // Media2 used by Catalog
    type Media2 struct {

        
            Landscape string  `json:"landscape"`
            Logo string  `json:"logo"`
            Portrait string  `json:"portrait"`
         
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
    
    // CategoryRequestBody used by Catalog
    type CategoryRequestBody struct {

        
            Slug string  `json:"slug"`
            Synonyms []string  `json:"synonyms"`
            Hierarchy []Hierarchy  `json:"hierarchy"`
            Media Media2  `json:"media"`
            Priority float64  `json:"priority"`
            Departments []float64  `json:"departments"`
            Tryouts []string  `json:"tryouts"`
            Marketplaces CategoryMapping  `json:"marketplaces"`
            Name string  `json:"name"`
            IsActive bool  `json:"is_active"`
            Level float64  `json:"level"`
         
    }
    
    // CategoryCreateResponse used by Catalog
    type CategoryCreateResponse struct {

        
            Message string  `json:"message"`
            UID float64  `json:"uid"`
         
    }
    
    // Category used by Catalog
    type Category struct {

        
            Slug string  `json:"slug"`
            Synonyms []string  `json:"synonyms"`
            Hierarchy []Hierarchy  `json:"hierarchy"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Media Media2  `json:"media"`
            Priority float64  `json:"priority"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Departments []float64  `json:"departments"`
            ModifiedOn string  `json:"modified_on"`
            Tryouts []string  `json:"tryouts"`
            Marketplaces CategoryMapping  `json:"marketplaces"`
            ID string  `json:"id"`
            CreatedOn string  `json:"created_on"`
            Name string  `json:"name"`
            IsActive bool  `json:"is_active"`
            Level float64  `json:"level"`
            UID float64  `json:"uid"`
         
    }
    
    // CategoryResponse used by Catalog
    type CategoryResponse struct {

        
            Items []Category  `json:"items"`
            Page Page  `json:"page"`
         
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
    
    // ReturnConfig used by Catalog
    type ReturnConfig struct {

        
            Time float64  `json:"time"`
            Returnable bool  `json:"returnable"`
            Unit string  `json:"unit"`
         
    }
    
    // CustomOrder used by Catalog
    type CustomOrder struct {

        
            ManufacturingTimeUnit string  `json:"manufacturing_time_unit"`
            IsCustomOrder bool  `json:"is_custom_order"`
            ManufacturingTime float64  `json:"manufacturing_time"`
         
    }
    
    // ProductPublish used by Catalog
    type ProductPublish struct {

        
            IsSet bool  `json:"is_set"`
            ProductOnlineDate string  `json:"product_online_date"`
         
    }
    
    // TeaserTag used by Catalog
    type TeaserTag struct {

        
            Tag string  `json:"tag"`
            URL string  `json:"url"`
         
    }
    
    // Trader used by Catalog
    type Trader struct {

        
            Address []string  `json:"address"`
            Type string  `json:"type"`
            Name string  `json:"name"`
         
    }
    
    // NetQuantity used by Catalog
    type NetQuantity struct {

        
            Value float64  `json:"value"`
            Unit string  `json:"unit"`
         
    }
    
    // TaxIdentifier used by Catalog
    type TaxIdentifier struct {

        
            HsnCodeID string  `json:"hsn_code_id"`
            HsnCode string  `json:"hsn_code"`
            ReportingHsn string  `json:"reporting_hsn"`
         
    }
    
    // ProductCreateUpdateSchemaV2 used by Catalog
    type ProductCreateUpdateSchemaV2 struct {

        
            Action string  `json:"action"`
            ItemCode string  `json:"item_code"`
            BulkJobID string  `json:"bulk_job_id"`
            ReturnConfig ReturnConfig  `json:"return_config"`
            IsImageLessProduct bool  `json:"is_image_less_product"`
            Media []Media1  `json:"media"`
            BrandUID float64  `json:"brand_uid"`
            Departments []float64  `json:"departments"`
            Name string  `json:"name"`
            CountryOfOrigin string  `json:"country_of_origin"`
            UID float64  `json:"uid"`
            MultiSize bool  `json:"multi_size"`
            CustomOrder CustomOrder  `json:"custom_order"`
            ProductGroupTag []string  `json:"product_group_tag"`
            ChangeRequestID interface{}  `json:"change_request_id"`
            CategorySlug string  `json:"category_slug"`
            IsDependent bool  `json:"is_dependent"`
            SizeGuide string  `json:"size_guide"`
            Sizes []map[string]interface{}  `json:"sizes"`
            NoOfBoxes float64  `json:"no_of_boxes"`
            Currency string  `json:"currency"`
            ProductPublish ProductPublish  `json:"product_publish"`
            ShortDescription string  `json:"short_description"`
            TemplateTag string  `json:"template_tag"`
            Variants map[string]interface{}  `json:"variants"`
            IsSet bool  `json:"is_set"`
            Highlights []string  `json:"highlights"`
            TeaserTag TeaserTag  `json:"teaser_tag"`
            Trader []Trader  `json:"trader"`
            IsActive bool  `json:"is_active"`
            NetQuantity NetQuantity  `json:"net_quantity"`
            Slug string  `json:"slug"`
            Description string  `json:"description"`
            CompanyID float64  `json:"company_id"`
            Tags []string  `json:"tags"`
            VariantMedia map[string]interface{}  `json:"variant_media"`
            VariantGroup map[string]interface{}  `json:"variant_group"`
            Attributes map[string]interface{}  `json:"attributes"`
            Requester string  `json:"requester"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            TaxIdentifier TaxIdentifier  `json:"tax_identifier"`
            ItemType string  `json:"item_type"`
         
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
    
    // Logo used by Catalog
    type Logo struct {

        
            AspectRatioF float64  `json:"aspect_ratio_f"`
            URL string  `json:"url"`
            SecureURL string  `json:"secure_url"`
            AspectRatio string  `json:"aspect_ratio"`
         
    }
    
    // Brand used by Catalog
    type Brand struct {

        
            Logo Logo  `json:"logo"`
            Name string  `json:"name"`
            UID float64  `json:"uid"`
         
    }
    
    // Image used by Catalog
    type Image struct {

        
            AspectRatioF float64  `json:"aspect_ratio_f"`
            URL string  `json:"url"`
            SecureURL string  `json:"secure_url"`
            AspectRatio string  `json:"aspect_ratio"`
         
    }
    
    // ProductPublish1 used by Catalog
    type ProductPublish1 struct {

        
            IsSet bool  `json:"is_set"`
            ProductOnlineDate string  `json:"product_online_date"`
         
    }
    
    // NetQuantityResponse used by Catalog
    type NetQuantityResponse struct {

        
            Value float64  `json:"value"`
            Unit string  `json:"unit"`
         
    }
    
    // ProductSchemaV2 used by Catalog
    type ProductSchemaV2 struct {

        
            ItemCode string  `json:"item_code"`
            ImageNature string  `json:"image_nature"`
            ReturnConfig ReturnConfigResponse  `json:"return_config"`
            IsImageLessProduct bool  `json:"is_image_less_product"`
            Media []Media1  `json:"media"`
            Moq map[string]interface{}  `json:"moq"`
            BrandUID float64  `json:"brand_uid"`
            Departments []float64  `json:"departments"`
            AllIdentifiers []string  `json:"all_identifiers"`
            Name string  `json:"name"`
            CountryOfOrigin string  `json:"country_of_origin"`
            UID float64  `json:"uid"`
            CategoryUID float64  `json:"category_uid"`
            MultiSize bool  `json:"multi_size"`
            HsnCode string  `json:"hsn_code"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            ProductGroupTag []string  `json:"product_group_tag"`
            Pending string  `json:"pending"`
            PrimaryColor string  `json:"primary_color"`
            IsExpirable bool  `json:"is_expirable"`
            VerifiedBy VerifiedBy  `json:"verified_by"`
            ModifiedOn string  `json:"modified_on"`
            CategorySlug string  `json:"category_slug"`
            Brand Brand  `json:"brand"`
            IsDependent bool  `json:"is_dependent"`
            Images []Image  `json:"images"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            CreatedOn string  `json:"created_on"`
            Stage string  `json:"stage"`
            IsPhysical bool  `json:"is_physical"`
            SizeGuide string  `json:"size_guide"`
            AllSizes []map[string]interface{}  `json:"all_sizes"`
            Sizes []map[string]interface{}  `json:"sizes"`
            NoOfBoxes float64  `json:"no_of_boxes"`
            Color string  `json:"color"`
            Currency string  `json:"currency"`
            Category map[string]interface{}  `json:"category"`
            ProductPublish ProductPublish1  `json:"product_publish"`
            ShortDescription string  `json:"short_description"`
            TemplateTag string  `json:"template_tag"`
            Highlights []string  `json:"highlights"`
            IsSet bool  `json:"is_set"`
            Variants map[string]interface{}  `json:"variants"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            Trader []Trader  `json:"trader"`
            IsActive bool  `json:"is_active"`
            NetQuantity NetQuantityResponse  `json:"net_quantity"`
            Description string  `json:"description"`
            Slug string  `json:"slug"`
            Tags []string  `json:"tags"`
            CompanyID float64  `json:"company_id"`
            VariantMedia map[string]interface{}  `json:"variant_media"`
            VariantGroup map[string]interface{}  `json:"variant_group"`
            Attributes map[string]interface{}  `json:"attributes"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            ItemType string  `json:"item_type"`
            ID string  `json:"id"`
            L3Mapping []string  `json:"l3_mapping"`
            AllCompanyIds []float64  `json:"all_company_ids"`
            VerifiedOn string  `json:"verified_on"`
         
    }
    
    // ProductListingResponseV2 used by Catalog
    type ProductListingResponseV2 struct {

        
            Items []ProductSchemaV2  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // ProductVariants used by Catalog
    type ProductVariants struct {

        
            ItemCode string  `json:"item_code"`
            CategoryUID float64  `json:"category_uid"`
            Media []Media1  `json:"media"`
            BrandUID float64  `json:"brand_uid"`
            Name string  `json:"name"`
            UID float64  `json:"uid"`
         
    }
    
    // ProductVariantsResponse used by Catalog
    type ProductVariantsResponse struct {

        
            Page Page  `json:"page"`
            Variants []ProductVariants  `json:"variants"`
         
    }
    
    // AttributeMasterSerializer used by Catalog
    type AttributeMasterSerializer struct {

        
            Details AttributeMasterDetails  `json:"details"`
            Departments []string  `json:"departments"`
            Name string  `json:"name"`
            Variant bool  `json:"variant"`
            Suggestion string  `json:"suggestion"`
            Unit string  `json:"unit"`
            ModifiedOn string  `json:"modified_on"`
            Logo string  `json:"logo"`
            CreatedOn string  `json:"created_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            RawKey string  `json:"raw_key"`
            EnabledForEndConsumer bool  `json:"enabled_for_end_consumer"`
            IsNested bool  `json:"is_nested"`
            Filters AttributeMasterFilter  `json:"filters"`
            Slug string  `json:"slug"`
            Description string  `json:"description"`
            Synonyms map[string]interface{}  `json:"synonyms"`
            Tags []string  `json:"tags"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Schema AttributeMaster  `json:"schema"`
         
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

        
            ItemDimensionsUnitOfMeasure string  `json:"item_dimensions_unit_of_measure"`
            ItemLength float64  `json:"item_length"`
            ItemWeight float64  `json:"item_weight"`
            ItemHeight float64  `json:"item_height"`
            Size string  `json:"size"`
            Identifiers []ValidateIdentifier  `json:"identifiers"`
            ItemWeightUnitOfMeasure string  `json:"item_weight_unit_of_measure"`
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

        
            IsSet bool  `json:"is_set"`
            ProductOnlineDate float64  `json:"product_online_date"`
         
    }
    
    // Product used by Catalog
    type Product struct {

        
            ItemCode string  `json:"item_code"`
            ImageNature string  `json:"image_nature"`
            ReturnConfig ReturnConfigResponse  `json:"return_config"`
            IsImageLessProduct bool  `json:"is_image_less_product"`
            Media []Media1  `json:"media"`
            Moq map[string]interface{}  `json:"moq"`
            BrandUID float64  `json:"brand_uid"`
            Departments []float64  `json:"departments"`
            AllIdentifiers []string  `json:"all_identifiers"`
            Name string  `json:"name"`
            CountryOfOrigin string  `json:"country_of_origin"`
            UID float64  `json:"uid"`
            CategoryUID float64  `json:"category_uid"`
            MultiSize bool  `json:"multi_size"`
            HsnCode string  `json:"hsn_code"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            ProductGroupTag []string  `json:"product_group_tag"`
            Pending string  `json:"pending"`
            PrimaryColor string  `json:"primary_color"`
            IsExpirable bool  `json:"is_expirable"`
            VerifiedBy VerifiedBy  `json:"verified_by"`
            ModifiedOn string  `json:"modified_on"`
            CategorySlug string  `json:"category_slug"`
            Brand Brand  `json:"brand"`
            IsDependent bool  `json:"is_dependent"`
            Images []Image  `json:"images"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            CreatedOn string  `json:"created_on"`
            Stage string  `json:"stage"`
            IsPhysical bool  `json:"is_physical"`
            SizeGuide string  `json:"size_guide"`
            AllSizes []map[string]interface{}  `json:"all_sizes"`
            Sizes []map[string]interface{}  `json:"sizes"`
            NoOfBoxes float64  `json:"no_of_boxes"`
            Color string  `json:"color"`
            Currency string  `json:"currency"`
            Category map[string]interface{}  `json:"category"`
            ProductPublish ProductPublished  `json:"product_publish"`
            ShortDescription string  `json:"short_description"`
            TemplateTag string  `json:"template_tag"`
            Highlights []string  `json:"highlights"`
            IsSet bool  `json:"is_set"`
            Variants map[string]interface{}  `json:"variants"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            Trader []Trader  `json:"trader"`
            IsActive bool  `json:"is_active"`
            NetQuantity NetQuantityResponse  `json:"net_quantity"`
            Description string  `json:"description"`
            Slug string  `json:"slug"`
            Tags []string  `json:"tags"`
            CompanyID float64  `json:"company_id"`
            VariantMedia map[string]interface{}  `json:"variant_media"`
            VariantGroup map[string]interface{}  `json:"variant_group"`
            Attributes map[string]interface{}  `json:"attributes"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            ItemType string  `json:"item_type"`
            ID string  `json:"id"`
            L3Mapping []string  `json:"l3_mapping"`
            AllCompanyIds []float64  `json:"all_company_ids"`
            VerifiedOn string  `json:"verified_on"`
         
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
            CancelledRecords []map[string]interface{}  `json:"cancelled_records"`
            Failed float64  `json:"failed"`
            CustomTemplateTag string  `json:"custom_template_tag"`
            CreatedBy UserInfo1  `json:"created_by"`
            FilePath string  `json:"file_path"`
            ModifiedOn string  `json:"modified_on"`
            Succeed float64  `json:"succeed"`
            TemplateTag string  `json:"template_tag"`
            FailedRecords []map[string]interface{}  `json:"failed_records"`
            Stage string  `json:"stage"`
            CreatedOn string  `json:"created_on"`
            Cancelled float64  `json:"cancelled"`
            ModifiedBy UserInfo1  `json:"modified_by"`
            IsActive bool  `json:"is_active"`
            Total float64  `json:"total"`
         
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

        
            UserID string  `json:"user_id"`
            FullName string  `json:"full_name"`
            Username string  `json:"username"`
         
    }
    
    // ProductBulkRequest used by Catalog
    type ProductBulkRequest struct {

        
            Total float64  `json:"total"`
            CompanyID float64  `json:"company_id"`
            CancelledRecords []string  `json:"cancelled_records"`
            Template ProductTemplate  `json:"template"`
            Failed float64  `json:"failed"`
            CreatedBy UserDetail1  `json:"created_by"`
            FilePath string  `json:"file_path"`
            ModifiedBy UserDetail1  `json:"modified_by"`
            ModifiedOn string  `json:"modified_on"`
            Succeed float64  `json:"succeed"`
            TemplateTag string  `json:"template_tag"`
            Stage string  `json:"stage"`
            CreatedOn string  `json:"created_on"`
            Cancelled float64  `json:"cancelled"`
            IsActive bool  `json:"is_active"`
            FailedRecords []string  `json:"failed_records"`
         
    }
    
    // ProductBulkRequestList used by Catalog
    type ProductBulkRequestList struct {

        
            Items ProductBulkRequest  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // BulkProductRequest used by Catalog
    type BulkProductRequest struct {

        
            Data []map[string]interface{}  `json:"data"`
            TemplateTag string  `json:"template_tag"`
            CompanyID float64  `json:"company_id"`
            BatchID string  `json:"batch_id"`
         
    }
    
    // ProductTagsViewResponse used by Catalog
    type ProductTagsViewResponse struct {

        
            Items []string  `json:"items"`
         
    }
    
    // ProductBulkAssets used by Catalog
    type ProductBulkAssets struct {

        
            CompanyID float64  `json:"company_id"`
            User map[string]interface{}  `json:"user"`
            URL string  `json:"url"`
         
    }
    
    // UserCommon used by Catalog
    type UserCommon struct {

        
            UserID string  `json:"user_id"`
            CompanyID float64  `json:"company_id"`
            Username string  `json:"username"`
         
    }
    
    // Items used by Catalog
    type Items struct {

        
            Total float64  `json:"total"`
            CompanyID float64  `json:"company_id"`
            TrackingURL string  `json:"tracking_url"`
            CancelledRecords []string  `json:"cancelled_records"`
            Failed float64  `json:"failed"`
            CreatedBy UserCommon  `json:"created_by"`
            FilePath string  `json:"file_path"`
            ModifiedOn string  `json:"modified_on"`
            Retry float64  `json:"retry"`
            FailedRecords []string  `json:"failed_records"`
            ID string  `json:"id"`
            Stage string  `json:"stage"`
            CreatedOn string  `json:"created_on"`
            Cancelled float64  `json:"cancelled"`
            ModifiedBy UserCommon  `json:"modified_by"`
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

        
            Size string  `json:"size"`
            ItemID float64  `json:"item_id"`
            CompanyID float64  `json:"company_id"`
         
    }
    
    // ProductSizeDeleteResponse used by Catalog
    type ProductSizeDeleteResponse struct {

        
            Data ProductSizeDeleteDataResponse  `json:"data"`
            Success bool  `json:"success"`
         
    }
    
    // ItemQuery used by Catalog
    type ItemQuery struct {

        
            BrandUID float64  `json:"brand_uid"`
            ItemCode string  `json:"item_code"`
            UID float64  `json:"uid"`
         
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

        
            ItemDimensionsUnitOfMeasure string  `json:"item_dimensions_unit_of_measure"`
            ItemLength float64  `json:"item_length"`
            StoreCode string  `json:"store_code"`
            Set InventorySet  `json:"set"`
            PriceTransfer float64  `json:"price_transfer"`
            PriceEffective float64  `json:"price_effective"`
            ItemWeight float64  `json:"item_weight"`
            ItemHeight float64  `json:"item_height"`
            Size string  `json:"size"`
            Quantity float64  `json:"quantity"`
            Currency string  `json:"currency"`
            ItemWeightUnitOfMeasure string  `json:"item_weight_unit_of_measure"`
            Identifiers []GTIN  `json:"identifiers"`
            IsSet bool  `json:"is_set"`
            ExpirationDate string  `json:"expiration_date"`
            ItemWidth float64  `json:"item_width"`
            Price float64  `json:"price"`
         
    }
    
    // InventoryRequest used by Catalog
    type InventoryRequest struct {

        
            Item ItemQuery  `json:"item"`
            CompanyID float64  `json:"company_id"`
            Sizes []InvSize  `json:"sizes"`
         
    }
    
    // InventoryResponse used by Catalog
    type InventoryResponse struct {

        
            UID string  `json:"uid"`
            SellableQuantity float64  `json:"sellable_quantity"`
            PriceTransfer float64  `json:"price_transfer"`
            Quantity float64  `json:"quantity"`
            InventoryUpdatedOn string  `json:"inventory_updated_on"`
            Size string  `json:"size"`
            Currency string  `json:"currency"`
            SellerIdentifier string  `json:"seller_identifier"`
            Identifiers map[string]interface{}  `json:"identifiers"`
            Store map[string]interface{}  `json:"store"`
            PriceEffective float64  `json:"price_effective"`
            ItemID float64  `json:"item_id"`
            Price float64  `json:"price"`
         
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
    
    // StoreMeta used by Catalog
    type StoreMeta struct {

        
            ID float64  `json:"id"`
         
    }
    
    // WeightResponse used by Catalog
    type WeightResponse struct {

        
            IsDefault bool  `json:"is_default"`
            Shipping float64  `json:"shipping"`
            Unit string  `json:"unit"`
         
    }
    
    // BrandMeta used by Catalog
    type BrandMeta struct {

        
            Name string  `json:"name"`
            ID float64  `json:"id"`
         
    }
    
    // UserSerializer used by Catalog
    type UserSerializer struct {

        
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
            Contact string  `json:"contact"`
            ID string  `json:"_id"`
            UID string  `json:"uid"`
         
    }
    
    // DimensionResponse used by Catalog
    type DimensionResponse struct {

        
            Height float64  `json:"height"`
            Width float64  `json:"width"`
            Unit string  `json:"unit"`
            IsDefault bool  `json:"is_default"`
            Length float64  `json:"length"`
         
    }
    
    // QuantityBase used by Catalog
    type QuantityBase struct {

        
            Count float64  `json:"count"`
            UpdatedAt string  `json:"updated_at"`
         
    }
    
    // Quantities used by Catalog
    type Quantities struct {

        
            Sellable QuantityBase  `json:"sellable"`
            Damaged QuantityBase  `json:"damaged"`
            NotAvailable QuantityBase  `json:"not_available"`
            OrderCommitted QuantityBase  `json:"order_committed"`
         
    }
    
    // Trader1 used by Catalog
    type Trader1 struct {

        
            Address []string  `json:"address"`
            Type string  `json:"type"`
            Name string  `json:"name"`
         
    }
    
    // ManufacturerResponse used by Catalog
    type ManufacturerResponse struct {

        
            Name string  `json:"name"`
            IsDefault bool  `json:"is_default"`
            Address string  `json:"address"`
         
    }
    
    // CompanyMeta used by Catalog
    type CompanyMeta struct {

        
            ID float64  `json:"id"`
         
    }
    
    // PriceMeta used by Catalog
    type PriceMeta struct {

        
            Currency string  `json:"currency"`
            UpdatedAt string  `json:"updated_at"`
            TpNotes map[string]interface{}  `json:"tp_notes"`
            Marked float64  `json:"marked"`
            Transfer float64  `json:"transfer"`
            Effective float64  `json:"effective"`
         
    }
    
    // InventorySellerResponse used by Catalog
    type InventorySellerResponse struct {

        
            FyndItemCode string  `json:"fynd_item_code"`
            Set InventorySet  `json:"set"`
            ReturnConfig ReturnConfig1  `json:"return_config"`
            Fragile bool  `json:"fragile"`
            Store StoreMeta  `json:"store"`
            CountryOfOrigin string  `json:"country_of_origin"`
            UID string  `json:"uid"`
            Weight WeightResponse  `json:"weight"`
            TrackInventory bool  `json:"track_inventory"`
            SellerIdentifier string  `json:"seller_identifier"`
            Brand BrandMeta  `json:"brand"`
            Stage string  `json:"stage"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            Meta map[string]interface{}  `json:"meta"`
            TraceID string  `json:"trace_id"`
            AddedOnStore string  `json:"added_on_store"`
            Dimension DimensionResponse  `json:"dimension"`
            FyndArticleCode string  `json:"fynd_article_code"`
            Size string  `json:"size"`
            RawMeta map[string]interface{}  `json:"raw_meta"`
            Quantities Quantities  `json:"quantities"`
            IsSet bool  `json:"is_set"`
            ExpirationDate string  `json:"expiration_date"`
            Trader []Trader1  `json:"trader"`
            IsActive bool  `json:"is_active"`
            FyndMeta map[string]interface{}  `json:"fynd_meta"`
            Tags []string  `json:"tags"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            CreatedBy UserSerializer  `json:"created_by"`
            Manufacturer ManufacturerResponse  `json:"manufacturer"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            TotalQuantity float64  `json:"total_quantity"`
            Identifier map[string]interface{}  `json:"identifier"`
            Company CompanyMeta  `json:"company"`
            ItemID float64  `json:"item_id"`
            Price PriceMeta  `json:"price"`
         
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
    
    // DateMeta used by Catalog
    type DateMeta struct {

        
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
            InventoryUpdatedOn string  `json:"inventory_updated_on"`
            AddedOnStore string  `json:"added_on_store"`
         
    }
    
    // ArticleStoreResponse used by Catalog
    type ArticleStoreResponse struct {

        
            StoreCode string  `json:"store_code"`
            Name string  `json:"name"`
            StoreType string  `json:"store_type"`
            UID float64  `json:"uid"`
         
    }
    
    // WeightResponse1 used by Catalog
    type WeightResponse1 struct {

        
            Shipping float64  `json:"shipping"`
            Unit string  `json:"unit"`
         
    }
    
    // BrandMeta1 used by Catalog
    type BrandMeta1 struct {

        
            Name string  `json:"name"`
            ID float64  `json:"id"`
         
    }
    
    // DimensionResponse1 used by Catalog
    type DimensionResponse1 struct {

        
            Height float64  `json:"height"`
            Length float64  `json:"length"`
            Width float64  `json:"width"`
            Unit string  `json:"unit"`
         
    }
    
    // Quantity used by Catalog
    type Quantity struct {

        
            Count float64  `json:"count"`
         
    }
    
    // QuantitiesArticle used by Catalog
    type QuantitiesArticle struct {

        
            Sellable Quantity  `json:"sellable"`
            Damaged Quantity  `json:"damaged"`
            NotAvailable Quantity  `json:"not_available"`
            OrderCommitted Quantity  `json:"order_committed"`
         
    }
    
    // Trader2 used by Catalog
    type Trader2 struct {

        
            Address []string  `json:"address"`
            Type string  `json:"type"`
            Name string  `json:"name"`
         
    }
    
    // ManufacturerResponse1 used by Catalog
    type ManufacturerResponse1 struct {

        
            Name string  `json:"name"`
            IsDefault bool  `json:"is_default"`
            Address string  `json:"address"`
         
    }
    
    // CompanyMeta1 used by Catalog
    type CompanyMeta1 struct {

        
            ID float64  `json:"id"`
         
    }
    
    // PriceArticle used by Catalog
    type PriceArticle struct {

        
            Currency string  `json:"currency"`
            TpNotes map[string]interface{}  `json:"tp_notes"`
            Marked float64  `json:"marked"`
            Transfer float64  `json:"transfer"`
            Effective float64  `json:"effective"`
         
    }
    
    // GetInventories used by Catalog
    type GetInventories struct {

        
            Platforms map[string]interface{}  `json:"platforms"`
            ReturnConfig ReturnConfig2  `json:"return_config"`
            DateMeta DateMeta  `json:"date_meta"`
            Store ArticleStoreResponse  `json:"store"`
            CountryOfOrigin string  `json:"country_of_origin"`
            UID string  `json:"uid"`
            Weight WeightResponse1  `json:"weight"`
            TrackInventory bool  `json:"track_inventory"`
            SellerIdentifier string  `json:"seller_identifier"`
            Brand BrandMeta1  `json:"brand"`
            Stage string  `json:"stage"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            TraceID string  `json:"trace_id"`
            Dimension DimensionResponse1  `json:"dimension"`
            InventoryUpdatedOn string  `json:"inventory_updated_on"`
            Size string  `json:"size"`
            Quantities QuantitiesArticle  `json:"quantities"`
            IsSet bool  `json:"is_set"`
            ExpirationDate string  `json:"expiration_date"`
            Trader []Trader2  `json:"trader"`
            Tags []string  `json:"tags"`
            CreatedBy UserSerializer  `json:"created_by"`
            Manufacturer ManufacturerResponse1  `json:"manufacturer"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            TotalQuantity float64  `json:"total_quantity"`
            ID string  `json:"id"`
            Identifier map[string]interface{}  `json:"identifier"`
            Company CompanyMeta1  `json:"company"`
            ItemID float64  `json:"item_id"`
            Price PriceArticle  `json:"price"`
         
    }
    
    // GetInventoriesResponse used by Catalog
    type GetInventoriesResponse struct {

        
            Items []GetInventories  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // BulkInventoryGetItems used by Catalog
    type BulkInventoryGetItems struct {

        
            Total float64  `json:"total"`
            CompanyID float64  `json:"company_id"`
            CancelledRecords []string  `json:"cancelled_records"`
            Failed float64  `json:"failed"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            FilePath string  `json:"file_path"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            ModifiedOn string  `json:"modified_on"`
            Succeed float64  `json:"succeed"`
            ID string  `json:"id"`
            Stage string  `json:"stage"`
            CreatedOn string  `json:"created_on"`
            Cancelled float64  `json:"cancelled"`
            IsActive bool  `json:"is_active"`
            FailedRecords []string  `json:"failed_records"`
         
    }
    
    // BulkInventoryGet used by Catalog
    type BulkInventoryGet struct {

        
            Items []BulkInventoryGetItems  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // InventoryJobPayload used by Catalog
    type InventoryJobPayload struct {

        
            StoreCode string  `json:"store_code"`
            ItemDimensionsUnitOfMeasure string  `json:"item_dimensions_unit_of_measure"`
            TraceID string  `json:"trace_id"`
            Tags []string  `json:"tags"`
            Quantity float64  `json:"quantity"`
            Currency string  `json:"currency"`
            SellerIdentifier string  `json:"seller_identifier"`
            TotalQuantity float64  `json:"total_quantity"`
            ItemWeightUnitOfMeasure string  `json:"item_weight_unit_of_measure"`
            PriceEffective float64  `json:"price_effective"`
            ExpirationDate string  `json:"expiration_date"`
            PriceMarked float64  `json:"price_marked"`
            Price float64  `json:"price"`
         
    }
    
    // InventoryBulkRequest used by Catalog
    type InventoryBulkRequest struct {

        
            User map[string]interface{}  `json:"user"`
            CompanyID float64  `json:"company_id"`
            BatchID string  `json:"batch_id"`
            Sizes []InventoryJobPayload  `json:"sizes"`
         
    }
    
    // InventoryExportRequest used by Catalog
    type InventoryExportRequest struct {

        
            Type string  `json:"type"`
            Store []float64  `json:"store"`
            Brand []float64  `json:"brand"`
         
    }
    
    // InventoryExportResponse used by Catalog
    type InventoryExportResponse struct {

        
            NotificationEmails []string  `json:"notification_emails"`
            SellerID float64  `json:"seller_id"`
            Type string  `json:"type"`
            CreatedBy string  `json:"created_by"`
            ModifiedOn string  `json:"modified_on"`
            TaskID string  `json:"task_id"`
            Filters map[string]interface{}  `json:"filters"`
            CreatedOn string  `json:"created_on"`
            Status string  `json:"status"`
         
    }
    
    // InventoryExportQuantityFilter used by Catalog
    type InventoryExportQuantityFilter struct {

        
            Min float64  `json:"min"`
            Max float64  `json:"max"`
            Operators string  `json:"operators"`
         
    }
    
    // InventoryExportAdvanceOption used by Catalog
    type InventoryExportAdvanceOption struct {

        
            StoreIds []float64  `json:"store_ids"`
            Quantity InventoryExportQuantityFilter  `json:"quantity"`
            ToDate string  `json:"to_date"`
            BrandIds []float64  `json:"brand_ids"`
            FromDate string  `json:"from_date"`
         
    }
    
    // InventoryExportJob used by Catalog
    type InventoryExportJob struct {

        
            SellerID float64  `json:"seller_id"`
            CompletedOn string  `json:"completed_on"`
            Type string  `json:"type"`
            URL string  `json:"url"`
            TaskID string  `json:"task_id"`
            Filters InventoryExportAdvanceOption  `json:"filters"`
            NotificationEmails []string  `json:"notification_emails"`
            Status string  `json:"status"`
         
    }
    
    // InventoryExportFilter used by Catalog
    type InventoryExportFilter struct {

        
            StoreIds []float64  `json:"store_ids"`
            Quantity InventoryExportQuantityFilter  `json:"quantity"`
            ToDate string  `json:"to_date"`
            BrandIds []float64  `json:"brand_ids"`
            FromDate string  `json:"from_date"`
         
    }
    
    // InventoryCreateRequest used by Catalog
    type InventoryCreateRequest struct {

        
            NotificationEmails []string  `json:"notification_emails"`
            Type string  `json:"type"`
            Filters InventoryExportFilter  `json:"filters"`
            Data []string  `json:"data"`
         
    }
    
    // InventoryJobFilters used by Catalog
    type InventoryJobFilters struct {

        
            Brands []string  `json:"brands"`
            Quantity InventoryExportQuantityFilter  `json:"quantity"`
            ToDate string  `json:"to_date"`
            Stores []string  `json:"stores"`
            FromDate string  `json:"from_date"`
         
    }
    
    // InventoryJobDetailResponse used by Catalog
    type InventoryJobDetailResponse struct {

        
            NotificationEmails []string  `json:"notification_emails"`
            SellerID float64  `json:"seller_id"`
            CompletedOn string  `json:"completed_on"`
            Type string  `json:"type"`
            URL string  `json:"url"`
            CreatedBy UserDetail  `json:"created_by"`
            ModifiedOn string  `json:"modified_on"`
            TaskID string  `json:"task_id"`
            CancelledBy UserDetail  `json:"cancelled_by"`
            Filters InventoryJobFilters  `json:"filters"`
            ID string  `json:"id"`
            CancelledOn string  `json:"cancelled_on"`
            CreatedOn string  `json:"created_on"`
            Status string  `json:"status"`
         
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

        
            TraceID string  `json:"trace_id"`
            Tags []string  `json:"tags"`
            SellerIdentifier string  `json:"seller_identifier"`
            TotalQuantity float64  `json:"total_quantity"`
            StoreID float64  `json:"store_id"`
            PriceEffective float64  `json:"price_effective"`
            ExpirationDate string  `json:"expiration_date"`
            PriceMarked float64  `json:"price_marked"`
         
    }
    
    // InventoryRequestSchemaV2 used by Catalog
    type InventoryRequestSchemaV2 struct {

        
            Payload []InventoryPayload  `json:"payload"`
            CompanyID float64  `json:"company_id"`
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

        
            Message string  `json:"message"`
            Items []InventoryResponseItem  `json:"items"`
         
    }
    
    // HsnCodesObject used by Catalog
    type HsnCodesObject struct {

        
            Threshold1 float64  `json:"threshold1"`
            CompanyID float64  `json:"company_id"`
            HsnCode string  `json:"hsn_code"`
            Hs2Code string  `json:"hs2_code"`
            ModifiedOn string  `json:"modified_on"`
            TaxOnMrp bool  `json:"tax_on_mrp"`
            ID string  `json:"id"`
            Threshold2 float64  `json:"threshold2"`
            Tax1 float64  `json:"tax1"`
            TaxOnEsp bool  `json:"tax_on_esp"`
            Tax2 float64  `json:"tax2"`
         
    }
    
    // HsnCode used by Catalog
    type HsnCode struct {

        
            Data HsnCodesObject  `json:"data"`
         
    }
    
    // HsnUpsert used by Catalog
    type HsnUpsert struct {

        
            Threshold1 float64  `json:"threshold1"`
            UID float64  `json:"uid"`
            CompanyID float64  `json:"company_id"`
            HsnCode string  `json:"hsn_code"`
            Hs2Code string  `json:"hs2_code"`
            TaxOnMrp bool  `json:"tax_on_mrp"`
            Threshold2 float64  `json:"threshold2"`
            Tax1 float64  `json:"tax1"`
            IsActive bool  `json:"is_active"`
            TaxOnEsp bool  `json:"tax_on_esp"`
            Tax2 float64  `json:"tax2"`
         
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

        
            Description string  `json:"description"`
            CountryCode string  `json:"country_code"`
            HsnCode string  `json:"hsn_code"`
            ReportingHsn string  `json:"reporting_hsn"`
            Type string  `json:"type"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ModifiedOn string  `json:"modified_on"`
            HsnCodeID string  `json:"hsn_code_id"`
            CreatedOn string  `json:"created_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Taxes []TaxSlab  `json:"taxes"`
         
    }
    
    // PageResponse used by Catalog
    type PageResponse struct {

        
            Size float64  `json:"size"`
            ItemTotal float64  `json:"item_total"`
            HasPrevious bool  `json:"has_previous"`
            Current string  `json:"current"`
            HasNext bool  `json:"has_next"`
         
    }
    
    // HsnCodesListingResponseSchemaV2 used by Catalog
    type HsnCodesListingResponseSchemaV2 struct {

        
            Items []HSNDataInsertV2  `json:"items"`
            Page PageResponse  `json:"page"`
         
    }
    
    // BrandItem used by Catalog
    type BrandItem struct {

        
            Action Action  `json:"action"`
            Slug string  `json:"slug"`
            Discount string  `json:"discount"`
            Logo Media  `json:"logo"`
            Departments []string  `json:"departments"`
            Name string  `json:"name"`
            Banners ImageUrls  `json:"banners"`
            UID float64  `json:"uid"`
         
    }
    
    // BrandListingResponse used by Catalog
    type BrandListingResponse struct {

        
            Items []BrandItem  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // Department used by Catalog
    type Department struct {

        
            Slug string  `json:"slug"`
            PriorityOrder float64  `json:"priority_order"`
            Logo Media  `json:"logo"`
            Name string  `json:"name"`
            UID float64  `json:"uid"`
         
    }
    
    // DepartmentResponse used by Catalog
    type DepartmentResponse struct {

        
            Items []Department  `json:"items"`
         
    }
    
    // ThirdLevelChild used by Catalog
    type ThirdLevelChild struct {

        
            Action Action  `json:"action"`
            Slug string  `json:"slug"`
            Childs []map[string]interface{}  `json:"childs"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Name string  `json:"name"`
            Banners ImageUrls  `json:"banners"`
            UID float64  `json:"uid"`
         
    }
    
    // SecondLevelChild used by Catalog
    type SecondLevelChild struct {

        
            Action Action  `json:"action"`
            Slug string  `json:"slug"`
            Childs []ThirdLevelChild  `json:"childs"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Name string  `json:"name"`
            Banners ImageUrls  `json:"banners"`
            UID float64  `json:"uid"`
         
    }
    
    // Child used by Catalog
    type Child struct {

        
            Action Action  `json:"action"`
            Slug string  `json:"slug"`
            Childs []SecondLevelChild  `json:"childs"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Name string  `json:"name"`
            Banners ImageUrls  `json:"banners"`
            UID float64  `json:"uid"`
         
    }
    
    // CategoryItems used by Catalog
    type CategoryItems struct {

        
            Action Action  `json:"action"`
            Slug string  `json:"slug"`
            Childs []Child  `json:"childs"`
            Name string  `json:"name"`
            Banners ImageUrls  `json:"banners"`
            UID float64  `json:"uid"`
         
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
            Page Page  `json:"page"`
            SortOn []ProductSortOn  `json:"sort_on"`
            Operators map[string]interface{}  `json:"operators"`
         
    }
    
    // ProductDetail used by Catalog
    type ProductDetail struct {

        
            ItemCode string  `json:"item_code"`
            ImageNature string  `json:"image_nature"`
            Type string  `json:"type"`
            Name string  `json:"name"`
            HasVariant bool  `json:"has_variant"`
            UID float64  `json:"uid"`
            Medias []Media1  `json:"medias"`
            RatingCount float64  `json:"rating_count"`
            Brand ProductBrand  `json:"brand"`
            PromoMeta map[string]interface{}  `json:"promo_meta"`
            Similars []string  `json:"similars"`
            Color string  `json:"color"`
            ShortDescription string  `json:"short_description"`
            Highlights []string  `json:"highlights"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            Rating float64  `json:"rating"`
            Slug string  `json:"slug"`
            Description string  `json:"description"`
            Attributes map[string]interface{}  `json:"attributes"`
            ProductOnlineDate string  `json:"product_online_date"`
            ItemType string  `json:"item_type"`
            Tryouts []string  `json:"tryouts"`
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
         
    }
    
    // InventoryPage used by Catalog
    type InventoryPage struct {

        
            Type string  `json:"type"`
            ItemTotal float64  `json:"item_total"`
            HasPrevious bool  `json:"has_previous"`
            NextID string  `json:"next_id"`
            HasNext bool  `json:"has_next"`
         
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

        
            Query ArticleQuery  `json:"query"`
            Quantity float64  `json:"quantity"`
            GroupID string  `json:"group_id"`
            ArticleAssignment ArticleAssignment  `json:"article_assignment"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // AssignStore used by Catalog
    type AssignStore struct {

        
            ChannelType string  `json:"channel_type"`
            CompanyID float64  `json:"company_id"`
            AppID string  `json:"app_id"`
            StoreIds []float64  `json:"store_ids"`
            ChannelIdentifier string  `json:"channel_identifier"`
            Articles []AssignStoreArticle  `json:"articles"`
            Pincode string  `json:"pincode"`
         
    }
    
    // StoreAssignError used by Catalog
    type StoreAssignError struct {

        
            Message string  `json:"message"`
            Type string  `json:"type"`
            Value interface{}  `json:"value"`
         
    }
    
    // ArticleAssignment1 used by Catalog
    type ArticleAssignment1 struct {

        
            Strategy string  `json:"strategy"`
            Level string  `json:"level"`
         
    }
    
    // StoreAssign used by Catalog
    type StoreAssign struct {

        
            UID string  `json:"uid"`
            PriceMarked float64  `json:"price_marked"`
            CompanyID float64  `json:"company_id"`
            PriceEffective float64  `json:"price_effective"`
            SCity string  `json:"s_city"`
            Quantity float64  `json:"quantity"`
            Size string  `json:"size"`
            StorePincode float64  `json:"store_pincode"`
            GroupID string  `json:"group_id"`
            ArticleAssignment ArticleAssignment1  `json:"article_assignment"`
            ID string  `json:"_id"`
            StoreID float64  `json:"store_id"`
            StrategyWiseListing []map[string]interface{}  `json:"strategy_wise_listing"`
            Index float64  `json:"index"`
            ItemID float64  `json:"item_id"`
            Status bool  `json:"status"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // StoreAssignResponse used by Catalog
    type StoreAssignResponse struct {

        
            Error StoreAssignError  `json:"error"`
            Items []StoreAssign  `json:"items"`
            Success bool  `json:"success"`
         
    }
    
    // InvoiceCredSerializer used by Catalog
    type InvoiceCredSerializer struct {

        
            Password string  `json:"password"`
            Username string  `json:"username"`
            Enabled bool  `json:"enabled"`
         
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
    
    // GetAddressSerializer used by Catalog
    type GetAddressSerializer struct {

        
            AddressType string  `json:"address_type"`
            Landmark string  `json:"landmark"`
            CountryCode string  `json:"country_code"`
            Address1 string  `json:"address1"`
            State string  `json:"state"`
            Longitude float64  `json:"longitude"`
            Address2 string  `json:"address2"`
            City string  `json:"city"`
            Country string  `json:"country"`
            Latitude float64  `json:"latitude"`
            Pincode float64  `json:"pincode"`
         
    }
    
    // UserSerializer1 used by Catalog
    type UserSerializer1 struct {

        
            UserID string  `json:"user_id"`
            Contact string  `json:"contact"`
            Username string  `json:"username"`
         
    }
    
    // SellerPhoneNumber used by Catalog
    type SellerPhoneNumber struct {

        
            Number string  `json:"number"`
            CountryCode float64  `json:"country_code"`
         
    }
    
    // LocationIntegrationType used by Catalog
    type LocationIntegrationType struct {

        
            Order string  `json:"order"`
            Inventory string  `json:"inventory"`
         
    }
    
    // Document used by Catalog
    type Document struct {

        
            LegalName string  `json:"legal_name"`
            Verified bool  `json:"verified"`
            Type string  `json:"type"`
            URL string  `json:"url"`
            Value string  `json:"value"`
         
    }
    
    // LocationManagerSerializer used by Catalog
    type LocationManagerSerializer struct {

        
            Email string  `json:"email"`
            Name string  `json:"name"`
            MobileNo SellerPhoneNumber  `json:"mobile_no"`
         
    }
    
    // LocationTimingSerializer used by Catalog
    type LocationTimingSerializer struct {

        
            Hour float64  `json:"hour"`
            Minute float64  `json:"minute"`
         
    }
    
    // LocationDayWiseSerializer used by Catalog
    type LocationDayWiseSerializer struct {

        
            Opening LocationTimingSerializer  `json:"opening"`
            Closing LocationTimingSerializer  `json:"closing"`
            Weekday string  `json:"weekday"`
            Open bool  `json:"open"`
         
    }
    
    // UserSerializer2 used by Catalog
    type UserSerializer2 struct {

        
            UserID string  `json:"user_id"`
            Contact string  `json:"contact"`
            Username string  `json:"username"`
         
    }
    
    // GetCompanySerializer used by Catalog
    type GetCompanySerializer struct {

        
            RejectReason string  `json:"reject_reason"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            BusinessType string  `json:"business_type"`
            VerifiedBy UserSerializer2  `json:"verified_by"`
            CreatedBy UserSerializer2  `json:"created_by"`
            Name string  `json:"name"`
            ModifiedOn string  `json:"modified_on"`
            Stage string  `json:"stage"`
            CreatedOn string  `json:"created_on"`
            ModifiedBy UserSerializer2  `json:"modified_by"`
            CompanyType string  `json:"company_type"`
            VerifiedOn string  `json:"verified_on"`
            UID float64  `json:"uid"`
         
    }
    
    // GetLocationSerializer used by Catalog
    type GetLocationSerializer struct {

        
            PhoneNumber string  `json:"phone_number"`
            GstCredentials InvoiceDetailsSerializer  `json:"gst_credentials"`
            Name string  `json:"name"`
            ProductReturnConfig ProductReturnConfigSerializer  `json:"product_return_config"`
            UID float64  `json:"uid"`
            Address GetAddressSerializer  `json:"address"`
            VerifiedBy UserSerializer1  `json:"verified_by"`
            StoreType string  `json:"store_type"`
            ModifiedOn string  `json:"modified_on"`
            Stage string  `json:"stage"`
            CreatedOn string  `json:"created_on"`
            NotificationEmails []string  `json:"notification_emails"`
            ModifiedBy UserSerializer1  `json:"modified_by"`
            Warnings map[string]interface{}  `json:"warnings"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            DisplayName string  `json:"display_name"`
            IntegrationType LocationIntegrationType  `json:"integration_type"`
            Documents []Document  `json:"documents"`
            Manager LocationManagerSerializer  `json:"manager"`
            Timing []LocationDayWiseSerializer  `json:"timing"`
            CreatedBy UserSerializer1  `json:"created_by"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Code string  `json:"code"`
            Company GetCompanySerializer  `json:"company"`
            VerifiedOn string  `json:"verified_on"`
         
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
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Logo string  `json:"logo"`
            Name string  `json:"name"`
            IsActive bool  `json:"is_active"`
            UID float64  `json:"uid"`
         
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
    
    // GetAddressSerializer used by CompanyProfile
    type GetAddressSerializer struct {

        
            AddressType string  `json:"address_type"`
            CountryCode string  `json:"country_code"`
            Address1 string  `json:"address1"`
            Latitude float64  `json:"latitude"`
            Country string  `json:"country"`
            Longitude float64  `json:"longitude"`
            State string  `json:"state"`
            Address2 string  `json:"address2"`
            Landmark string  `json:"landmark"`
            Pincode float64  `json:"pincode"`
            City string  `json:"city"`
         
    }
    
    // UserSerializer used by CompanyProfile
    type UserSerializer struct {

        
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
            Contact string  `json:"contact"`
         
    }
    
    // Document used by CompanyProfile
    type Document struct {

        
            LegalName string  `json:"legal_name"`
            Value string  `json:"value"`
            Type string  `json:"type"`
            URL string  `json:"url"`
            Verified bool  `json:"verified"`
         
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
    
    // CompanyTaxesSerializer used by CompanyProfile
    type CompanyTaxesSerializer struct {

        
            Rate float64  `json:"rate"`
            Enable bool  `json:"enable"`
            EffectiveDate string  `json:"effective_date"`
         
    }
    
    // GetCompanyProfileSerializerResponse used by CompanyProfile
    type GetCompanyProfileSerializerResponse struct {

        
            BusinessCountryInfo BusinessCountryInfo  `json:"business_country_info"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            CompanyType string  `json:"company_type"`
            ModifiedOn string  `json:"modified_on"`
            BusinessDetails BusinessDetails  `json:"business_details"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            Documents []Document  `json:"documents"`
            BusinessType string  `json:"business_type"`
            Stage string  `json:"stage"`
            Warnings map[string]interface{}  `json:"warnings"`
            UID float64  `json:"uid"`
            CreatedOn string  `json:"created_on"`
            CreatedBy UserSerializer  `json:"created_by"`
            ContactDetails ContactDetails  `json:"contact_details"`
            Name string  `json:"name"`
            VerifiedOn string  `json:"verified_on"`
            FranchiseEnabled bool  `json:"franchise_enabled"`
            BusinessInfo string  `json:"business_info"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            NotificationEmails []string  `json:"notification_emails"`
            Mode string  `json:"mode"`
            Taxes []CompanyTaxesSerializer  `json:"taxes"`
         
    }
    
    // ErrorResponse used by CompanyProfile
    type ErrorResponse struct {

        
            Code string  `json:"code"`
            Meta map[string]interface{}  `json:"meta"`
            Status float64  `json:"status"`
            Message string  `json:"message"`
         
    }
    
    // CompanyTaxesSerializer1 used by CompanyProfile
    type CompanyTaxesSerializer1 struct {

        
            Rate float64  `json:"rate"`
            Enable bool  `json:"enable"`
            EffectiveDate string  `json:"effective_date"`
         
    }
    
    // CreateUpdateAddressSerializer used by CompanyProfile
    type CreateUpdateAddressSerializer struct {

        
            CountryCode string  `json:"country_code"`
            AddressType string  `json:"address_type"`
            Address1 string  `json:"address1"`
            Latitude float64  `json:"latitude"`
            Country string  `json:"country"`
            Longitude float64  `json:"longitude"`
            State string  `json:"state"`
            Address2 string  `json:"address2"`
            Landmark string  `json:"landmark"`
            Pincode float64  `json:"pincode"`
            City string  `json:"city"`
         
    }
    
    // UpdateCompany used by CompanyProfile
    type UpdateCompany struct {

        
            BusinessInfo string  `json:"business_info"`
            BusinessDetails BusinessDetails  `json:"business_details"`
            ContactDetails ContactDetails  `json:"contact_details"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Taxes []CompanyTaxesSerializer1  `json:"taxes"`
            RejectReason string  `json:"reject_reason"`
            CompanyType string  `json:"company_type"`
            Slug string  `json:"slug"`
            Name string  `json:"name"`
            Addresses []CreateUpdateAddressSerializer  `json:"addresses"`
            NotificationEmails []string  `json:"notification_emails"`
            Documents []Document  `json:"documents"`
            BusinessType string  `json:"business_type"`
            FranchiseEnabled bool  `json:"franchise_enabled"`
            Warnings map[string]interface{}  `json:"warnings"`
         
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

        
            Product DocumentsObj  `json:"product"`
            Brand DocumentsObj  `json:"brand"`
            Store DocumentsObj  `json:"store"`
            CompanyDocuments DocumentsObj  `json:"company_documents"`
            Stage string  `json:"stage"`
            UID float64  `json:"uid"`
            StoreDocuments DocumentsObj  `json:"store_documents"`
         
    }
    
    // BrandBannerSerializer used by CompanyProfile
    type BrandBannerSerializer struct {

        
            Portrait string  `json:"portrait"`
            Landscape string  `json:"landscape"`
         
    }
    
    // GetBrandResponseSerializer used by CompanyProfile
    type GetBrandResponseSerializer struct {

        
            CreatedOn string  `json:"created_on"`
            SlugKey string  `json:"slug_key"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            RejectReason string  `json:"reject_reason"`
            Description string  `json:"description"`
            Name string  `json:"name"`
            Logo string  `json:"logo"`
            VerifiedOn string  `json:"verified_on"`
            Banner BrandBannerSerializer  `json:"banner"`
            ModifiedOn string  `json:"modified_on"`
            Mode string  `json:"mode"`
            Synonyms []string  `json:"synonyms"`
            Stage string  `json:"stage"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            Warnings map[string]interface{}  `json:"warnings"`
            UID float64  `json:"uid"`
         
    }
    
    // CreateUpdateBrandRequestSerializer used by CompanyProfile
    type CreateUpdateBrandRequestSerializer struct {

        
            BrandTier string  `json:"brand_tier"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Description string  `json:"description"`
            Name string  `json:"name"`
            CompanyID float64  `json:"company_id"`
            Logo string  `json:"logo"`
            Banner BrandBannerSerializer  `json:"banner"`
            Synonyms []string  `json:"synonyms"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            UID float64  `json:"uid"`
         
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

        
            BusinessCountryInfo BusinessCountryInfo  `json:"business_country_info"`
            CreatedOn string  `json:"created_on"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            CreatedBy UserSerializer  `json:"created_by"`
            RejectReason string  `json:"reject_reason"`
            CompanyType string  `json:"company_type"`
            Name string  `json:"name"`
            VerifiedOn string  `json:"verified_on"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            MarketChannels []string  `json:"market_channels"`
            NotificationEmails []string  `json:"notification_emails"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            ModifiedOn string  `json:"modified_on"`
            BusinessType string  `json:"business_type"`
            Stage string  `json:"stage"`
            UID float64  `json:"uid"`
            Details CompanyDetails  `json:"details"`
         
    }
    
    // CompanyBrandSerializer used by CompanyProfile
    type CompanyBrandSerializer struct {

        
            CreatedOn string  `json:"created_on"`
            CreatedBy UserSerializer  `json:"created_by"`
            Company CompanySerializer  `json:"company"`
            Brand GetBrandResponseSerializer  `json:"brand"`
            RejectReason string  `json:"reject_reason"`
            VerifiedOn string  `json:"verified_on"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            ModifiedOn string  `json:"modified_on"`
            Stage string  `json:"stage"`
            Warnings map[string]interface{}  `json:"warnings"`
            UID float64  `json:"uid"`
         
    }
    
    // CompanyBrandListSerializer used by CompanyProfile
    type CompanyBrandListSerializer struct {

        
            Page Page  `json:"page"`
            Items []CompanyBrandSerializer  `json:"items"`
         
    }
    
    // CompanyBrandPostRequestSerializer used by CompanyProfile
    type CompanyBrandPostRequestSerializer struct {

        
            UID float64  `json:"uid"`
            Brands []float64  `json:"brands"`
            Company float64  `json:"company"`
         
    }
    
    // ProductReturnConfigSerializer used by CompanyProfile
    type ProductReturnConfigSerializer struct {

        
            StoreUID float64  `json:"store_uid"`
            OnSameStore bool  `json:"on_same_store"`
         
    }
    
    // LocationTimingSerializer used by CompanyProfile
    type LocationTimingSerializer struct {

        
            Hour float64  `json:"hour"`
            Minute float64  `json:"minute"`
         
    }
    
    // LocationDayWiseSerializer used by CompanyProfile
    type LocationDayWiseSerializer struct {

        
            Weekday string  `json:"weekday"`
            Opening LocationTimingSerializer  `json:"opening"`
            Closing LocationTimingSerializer  `json:"closing"`
            Open bool  `json:"open"`
         
    }
    
    // GetCompanySerializer used by CompanyProfile
    type GetCompanySerializer struct {

        
            CreatedOn string  `json:"created_on"`
            CreatedBy UserSerializer  `json:"created_by"`
            RejectReason string  `json:"reject_reason"`
            CompanyType string  `json:"company_type"`
            Name string  `json:"name"`
            VerifiedOn string  `json:"verified_on"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            ModifiedOn string  `json:"modified_on"`
            BusinessType string  `json:"business_type"`
            Stage string  `json:"stage"`
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
            Title string  `json:"title"`
            HolidayType string  `json:"holiday_type"`
         
    }
    
    // LocationManagerSerializer used by CompanyProfile
    type LocationManagerSerializer struct {

        
            Email string  `json:"email"`
            MobileNo SellerPhoneNumber  `json:"mobile_no"`
            Name string  `json:"name"`
         
    }
    
    // InvoiceCredSerializer used by CompanyProfile
    type InvoiceCredSerializer struct {

        
            Username string  `json:"username"`
            Enabled bool  `json:"enabled"`
            Password string  `json:"password"`
         
    }
    
    // InvoiceDetailsSerializer used by CompanyProfile
    type InvoiceDetailsSerializer struct {

        
            EWaybill InvoiceCredSerializer  `json:"e_waybill"`
            EInvoice InvoiceCredSerializer  `json:"e_invoice"`
         
    }
    
    // GetLocationSerializer used by CompanyProfile
    type GetLocationSerializer struct {

        
            DisplayName string  `json:"display_name"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            PhoneNumber string  `json:"phone_number"`
            ModifiedOn string  `json:"modified_on"`
            ProductReturnConfig ProductReturnConfigSerializer  `json:"product_return_config"`
            AutoInvoice bool  `json:"auto_invoice"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            Documents []Document  `json:"documents"`
            Stage string  `json:"stage"`
            Address GetAddressSerializer  `json:"address"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            UID float64  `json:"uid"`
            Warnings map[string]interface{}  `json:"warnings"`
            CreditNote bool  `json:"credit_note"`
            Timing []LocationDayWiseSerializer  `json:"timing"`
            CreatedOn string  `json:"created_on"`
            CreatedBy UserSerializer  `json:"created_by"`
            Company GetCompanySerializer  `json:"company"`
            Name string  `json:"name"`
            StoreType string  `json:"store_type"`
            VerifiedOn string  `json:"verified_on"`
            Holiday []HolidaySchemaSerializer  `json:"holiday"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            Manager LocationManagerSerializer  `json:"manager"`
            GstCredentials InvoiceDetailsSerializer  `json:"gst_credentials"`
            Code string  `json:"code"`
            NotificationEmails []string  `json:"notification_emails"`
         
    }
    
    // LocationListSerializer used by CompanyProfile
    type LocationListSerializer struct {

        
            Page Page  `json:"page"`
            Items []GetLocationSerializer  `json:"items"`
         
    }
    
    // AddressSerializer used by CompanyProfile
    type AddressSerializer struct {

        
            AddressType string  `json:"address_type"`
            CountryCode string  `json:"country_code"`
            Address1 string  `json:"address1"`
            Latitude float64  `json:"latitude"`
            Country string  `json:"country"`
            Longitude float64  `json:"longitude"`
            State string  `json:"state"`
            Address2 string  `json:"address2"`
            Landmark string  `json:"landmark"`
            Pincode float64  `json:"pincode"`
            City string  `json:"city"`
         
    }
    
    // LocationSerializer used by CompanyProfile
    type LocationSerializer struct {

        
            DisplayName string  `json:"display_name"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Slug string  `json:"slug"`
            ProductReturnConfig ProductReturnConfigSerializer  `json:"product_return_config"`
            AutoInvoice bool  `json:"auto_invoice"`
            Documents []Document  `json:"documents"`
            Stage string  `json:"stage"`
            Address AddressSerializer  `json:"address"`
            Warnings map[string]interface{}  `json:"warnings"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            UID float64  `json:"uid"`
            CreditNote bool  `json:"credit_note"`
            Timing []LocationDayWiseSerializer  `json:"timing"`
            Company float64  `json:"company"`
            Name string  `json:"name"`
            StoreType string  `json:"store_type"`
            Holiday []HolidaySchemaSerializer  `json:"holiday"`
            Manager LocationManagerSerializer  `json:"manager"`
            GstCredentials InvoiceDetailsSerializer  `json:"gst_credentials"`
            Code string  `json:"code"`
            NotificationEmails []string  `json:"notification_emails"`
         
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

        
            AppID []string  `json:"app_id"`
            Anonymous bool  `json:"anonymous"`
            UserRegisteredAfter string  `json:"user_registered_after"`
         
    }
    
    // Rule used by Cart
    type Rule struct {

        
            Min float64  `json:"min"`
            Max float64  `json:"max"`
            Key float64  `json:"key"`
            Value float64  `json:"value"`
            DiscountQty float64  `json:"discount_qty"`
         
    }
    
    // Ownership used by Cart
    type Ownership struct {

        
            PayableCategory string  `json:"payable_category"`
            PayableBy string  `json:"payable_by"`
         
    }
    
    // State used by Cart
    type State struct {

        
            IsArchived bool  `json:"is_archived"`
            IsDisplay bool  `json:"is_display"`
            IsPublic bool  `json:"is_public"`
         
    }
    
    // Identifier used by Cart
    type Identifier struct {

        
            CategoryID []float64  `json:"category_id"`
            BrandID []float64  `json:"brand_id"`
            EmailDomain []string  `json:"email_domain"`
            ItemID []float64  `json:"item_id"`
            ExcludeBrandID []float64  `json:"exclude_brand_id"`
            CompanyID []float64  `json:"company_id"`
            CollectionID []string  `json:"collection_id"`
            StoreID []float64  `json:"store_id"`
            ArticleID []string  `json:"article_id"`
            UserID []string  `json:"user_id"`
         
    }
    
    // DisplayMetaDict used by Cart
    type DisplayMetaDict struct {

        
            Title string  `json:"title"`
            Subtitle string  `json:"subtitle"`
         
    }
    
    // DisplayMeta used by Cart
    type DisplayMeta struct {

        
            Title string  `json:"title"`
            Auto DisplayMetaDict  `json:"auto"`
            Description string  `json:"description"`
            Subtitle string  `json:"subtitle"`
            Apply DisplayMetaDict  `json:"apply"`
            Remove DisplayMetaDict  `json:"remove"`
         
    }
    
    // CouponSchedule used by Cart
    type CouponSchedule struct {

        
            Start string  `json:"start"`
            Cron string  `json:"cron"`
            Duration float64  `json:"duration"`
            NextSchedule []map[string]interface{}  `json:"next_schedule"`
            End string  `json:"end"`
         
    }
    
    // CouponDateMeta used by Cart
    type CouponDateMeta struct {

        
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
         
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
    
    // PaymentAllowValue used by Cart
    type PaymentAllowValue struct {

        
            Max float64  `json:"max"`
         
    }
    
    // PaymentModes used by Cart
    type PaymentModes struct {

        
            Types []string  `json:"types"`
            Uses PaymentAllowValue  `json:"uses"`
            Networks []string  `json:"networks"`
            Codes []string  `json:"codes"`
         
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
    
    // PriceRange used by Cart
    type PriceRange struct {

        
            Min float64  `json:"min"`
            Max float64  `json:"max"`
         
    }
    
    // Restrictions used by Cart
    type Restrictions struct {

        
            UserGroups []float64  `json:"user_groups"`
            PostOrder PostOrder  `json:"post_order"`
            BulkBundle BulkBundleRestriction  `json:"bulk_bundle"`
            Platforms []string  `json:"platforms"`
            OrderingStores []float64  `json:"ordering_stores"`
            Payments map[string]PaymentModes  `json:"payments"`
            UserType string  `json:"user_type"`
            Uses UsesRestriction  `json:"uses"`
            PriceRange PriceRange  `json:"price_range"`
            CouponAllowed bool  `json:"coupon_allowed"`
         
    }
    
    // CouponAuthor used by Cart
    type CouponAuthor struct {

        
            CreatedBy string  `json:"created_by"`
            ModifiedBy string  `json:"modified_by"`
         
    }
    
    // CouponAction used by Cart
    type CouponAction struct {

        
            TxnMode string  `json:"txn_mode"`
            ActionDate string  `json:"action_date"`
         
    }
    
    // RuleDefinition used by Cart
    type RuleDefinition struct {

        
            CurrencyCode string  `json:"currency_code"`
            Scope []string  `json:"scope"`
            IsExact bool  `json:"is_exact"`
            CalculateOn string  `json:"calculate_on"`
            ApplicableOn string  `json:"applicable_on"`
            Type string  `json:"type"`
            ValueType string  `json:"value_type"`
            AutoApply bool  `json:"auto_apply"`
         
    }
    
    // Validity used by Cart
    type Validity struct {

        
            Priority float64  `json:"priority"`
         
    }
    
    // CouponAdd used by Cart
    type CouponAdd struct {

        
            Validation Validation  `json:"validation"`
            Rule []Rule  `json:"rule"`
            Ownership Ownership  `json:"ownership"`
            TypeSlug string  `json:"type_slug"`
            State State  `json:"state"`
            Tags []string  `json:"tags"`
            Identifiers Identifier  `json:"identifiers"`
            DisplayMeta DisplayMeta  `json:"display_meta"`
            Schedule CouponSchedule  `json:"_schedule"`
            DateMeta CouponDateMeta  `json:"date_meta"`
            Restrictions Restrictions  `json:"restrictions"`
            Code string  `json:"code"`
            Author CouponAuthor  `json:"author"`
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

        
            Validation Validation  `json:"validation"`
            Rule []Rule  `json:"rule"`
            Ownership Ownership  `json:"ownership"`
            TypeSlug string  `json:"type_slug"`
            State State  `json:"state"`
            Tags []string  `json:"tags"`
            Identifiers Identifier  `json:"identifiers"`
            DisplayMeta DisplayMeta  `json:"display_meta"`
            Schedule CouponSchedule  `json:"_schedule"`
            DateMeta CouponDateMeta  `json:"date_meta"`
            Restrictions Restrictions  `json:"restrictions"`
            Code string  `json:"code"`
            Author CouponAuthor  `json:"author"`
            Action CouponAction  `json:"action"`
            RuleDefinition RuleDefinition  `json:"rule_definition"`
            Validity Validity  `json:"validity"`
         
    }
    
    // CouponPartialUpdate used by Cart
    type CouponPartialUpdate struct {

        
            Schedule CouponSchedule  `json:"schedule"`
            Archive bool  `json:"archive"`
         
    }
    
    // PostOrder1 used by Cart
    type PostOrder1 struct {

        
            ReturnAllowed bool  `json:"return_allowed"`
            CancellationAllowed bool  `json:"cancellation_allowed"`
         
    }
    
    // PaymentAllowValue1 used by Cart
    type PaymentAllowValue1 struct {

        
            Max float64  `json:"max"`
         
    }
    
    // PromotionPaymentModes used by Cart
    type PromotionPaymentModes struct {

        
            Uses PaymentAllowValue1  `json:"uses"`
            Codes []string  `json:"codes"`
            Type string  `json:"type"`
         
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

        
            UserGroups []float64  `json:"user_groups"`
            AnonymousUsers bool  `json:"anonymous_users"`
            PostOrder PostOrder1  `json:"post_order"`
            OrderQuantity float64  `json:"order_quantity"`
            Platforms []string  `json:"platforms"`
            Payments []PromotionPaymentModes  `json:"payments"`
            UserRegistered UserRegistered  `json:"user_registered"`
            UserID []string  `json:"user_id"`
            Uses UsesRestriction1  `json:"uses"`
         
    }
    
    // PromotionAuthor used by Cart
    type PromotionAuthor struct {

        
            CreatedBy string  `json:"created_by"`
            ModifiedBy string  `json:"modified_by"`
         
    }
    
    // PromotionAction used by Cart
    type PromotionAction struct {

        
            ActionDate string  `json:"action_date"`
            ActionType string  `json:"action_type"`
         
    }
    
    // Visibility used by Cart
    type Visibility struct {

        
            Pdp bool  `json:"pdp"`
            CouponList bool  `json:"coupon_list"`
         
    }
    
    // PromotionDateMeta used by Cart
    type PromotionDateMeta struct {

        
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
         
    }
    
    // Ownership1 used by Cart
    type Ownership1 struct {

        
            PayableCategory string  `json:"payable_category"`
            PayableBy string  `json:"payable_by"`
         
    }
    
    // DisplayMeta1 used by Cart
    type DisplayMeta1 struct {

        
            OfferLabel string  `json:"offer_label"`
            Name string  `json:"name"`
            Description string  `json:"description"`
            OfferText string  `json:"offer_text"`
         
    }
    
    // DiscountOffer used by Cart
    type DiscountOffer struct {

        
            MaxUsagePerTransaction float64  `json:"max_usage_per_transaction"`
            MaxOfferQuantity float64  `json:"max_offer_quantity"`
            DiscountPrice float64  `json:"discount_price"`
            MaxDiscountAmount float64  `json:"max_discount_amount"`
            MinOfferQuantity float64  `json:"min_offer_quantity"`
            DiscountAmount float64  `json:"discount_amount"`
            ApportionDiscount bool  `json:"apportion_discount"`
            PartialCanRet bool  `json:"partial_can_ret"`
            Code string  `json:"code"`
            DiscountPercentage float64  `json:"discount_percentage"`
         
    }
    
    // CompareObject used by Cart
    type CompareObject struct {

        
            GreaterThan float64  `json:"greater_than"`
            Equals float64  `json:"equals"`
            GreaterThanEquals float64  `json:"greater_than_equals"`
            LessThan float64  `json:"less_than"`
            LessThanEquals float64  `json:"less_than_equals"`
         
    }
    
    // ItemCriteria used by Cart
    type ItemCriteria struct {

        
            ItemExcludeCategory []float64  `json:"item_exclude_category"`
            AllItems bool  `json:"all_items"`
            AvailableZones []string  `json:"available_zones"`
            ItemDepartment []float64  `json:"item_department"`
            ItemStore []float64  `json:"item_store"`
            ItemSize []string  `json:"item_size"`
            ItemExcludeBrand []float64  `json:"item_exclude_brand"`
            ProductTags []string  `json:"product_tags"`
            ItemExcludeCompany []float64  `json:"item_exclude_company"`
            ItemBrand []float64  `json:"item_brand"`
            ItemExcludeDepartment []float64  `json:"item_exclude_department"`
            CartUniqueItemQuantity CompareObject  `json:"cart_unique_item_quantity"`
            ItemID []float64  `json:"item_id"`
            ItemExcludeL1Category []float64  `json:"item_exclude_l1_category"`
            ItemExcludeSku []string  `json:"item_exclude_sku"`
            ItemExcludeL2Category []float64  `json:"item_exclude_l2_category"`
            ItemL2Category []float64  `json:"item_l2_category"`
            CartUniqueItemAmount CompareObject  `json:"cart_unique_item_amount"`
            CartTotal CompareObject  `json:"cart_total"`
            CartQuantity CompareObject  `json:"cart_quantity"`
            ItemTags []string  `json:"item_tags"`
            ItemExcludeStore []float64  `json:"item_exclude_store"`
            BuyRules []string  `json:"buy_rules"`
            ItemExcludeID []float64  `json:"item_exclude_id"`
            ItemCategory []float64  `json:"item_category"`
            ItemSku []string  `json:"item_sku"`
            ItemCompany []float64  `json:"item_company"`
            ItemL1Category []float64  `json:"item_l1_category"`
         
    }
    
    // DiscountRule used by Cart
    type DiscountRule struct {

        
            Offer DiscountOffer  `json:"offer"`
            BuyCondition string  `json:"buy_condition"`
            ItemCriteria ItemCriteria  `json:"item_criteria"`
            DiscountType string  `json:"discount_type"`
         
    }
    
    // PromotionSchedule used by Cart
    type PromotionSchedule struct {

        
            Start string  `json:"start"`
            Published bool  `json:"published"`
            Cron string  `json:"cron"`
            Duration float64  `json:"duration"`
            NextSchedule []map[string]interface{}  `json:"next_schedule"`
            End string  `json:"end"`
         
    }
    
    // PromotionListItem used by Cart
    type PromotionListItem struct {

        
            ApplyExclusive string  `json:"apply_exclusive"`
            Restrictions Restrictions1  `json:"restrictions"`
            Code string  `json:"code"`
            Author PromotionAuthor  `json:"author"`
            PromotionType string  `json:"promotion_type"`
            Stackable bool  `json:"stackable"`
            PostOrderAction PromotionAction  `json:"post_order_action"`
            Visiblility Visibility  `json:"visiblility"`
            Currency string  `json:"currency"`
            DateMeta PromotionDateMeta  `json:"date_meta"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Ownership Ownership1  `json:"ownership"`
            DisplayMeta DisplayMeta1  `json:"display_meta"`
            DiscountRules []DiscountRule  `json:"discount_rules"`
            Schedule PromotionSchedule  `json:"_schedule"`
            Mode string  `json:"mode"`
            ApplicationID string  `json:"application_id"`
            PromoGroup string  `json:"promo_group"`
            ApplyPriority float64  `json:"apply_priority"`
            ApplyAllDiscount bool  `json:"apply_all_discount"`
            BuyRules map[string]ItemCriteria  `json:"buy_rules"`
            CalculateOn string  `json:"calculate_on"`
         
    }
    
    // PromotionsResponse used by Cart
    type PromotionsResponse struct {

        
            Items PromotionListItem  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // PromotionAdd used by Cart
    type PromotionAdd struct {

        
            ApplyExclusive string  `json:"apply_exclusive"`
            Restrictions Restrictions1  `json:"restrictions"`
            Code string  `json:"code"`
            Author PromotionAuthor  `json:"author"`
            PromotionType string  `json:"promotion_type"`
            Stackable bool  `json:"stackable"`
            PostOrderAction PromotionAction  `json:"post_order_action"`
            Visiblility Visibility  `json:"visiblility"`
            Currency string  `json:"currency"`
            DateMeta PromotionDateMeta  `json:"date_meta"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Ownership Ownership1  `json:"ownership"`
            DisplayMeta DisplayMeta1  `json:"display_meta"`
            DiscountRules []DiscountRule  `json:"discount_rules"`
            Schedule PromotionSchedule  `json:"_schedule"`
            Mode string  `json:"mode"`
            ApplicationID string  `json:"application_id"`
            PromoGroup string  `json:"promo_group"`
            ApplyPriority float64  `json:"apply_priority"`
            ApplyAllDiscount bool  `json:"apply_all_discount"`
            BuyRules map[string]ItemCriteria  `json:"buy_rules"`
            CalculateOn string  `json:"calculate_on"`
         
    }
    
    // PromotionUpdate used by Cart
    type PromotionUpdate struct {

        
            ApplyExclusive string  `json:"apply_exclusive"`
            Restrictions Restrictions1  `json:"restrictions"`
            Code string  `json:"code"`
            Author PromotionAuthor  `json:"author"`
            PromotionType string  `json:"promotion_type"`
            Stackable bool  `json:"stackable"`
            PostOrderAction PromotionAction  `json:"post_order_action"`
            Visiblility Visibility  `json:"visiblility"`
            Currency string  `json:"currency"`
            DateMeta PromotionDateMeta  `json:"date_meta"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Ownership Ownership1  `json:"ownership"`
            DisplayMeta DisplayMeta1  `json:"display_meta"`
            DiscountRules []DiscountRule  `json:"discount_rules"`
            Schedule PromotionSchedule  `json:"_schedule"`
            Mode string  `json:"mode"`
            ApplicationID string  `json:"application_id"`
            PromoGroup string  `json:"promo_group"`
            ApplyPriority float64  `json:"apply_priority"`
            ApplyAllDiscount bool  `json:"apply_all_discount"`
            BuyRules map[string]ItemCriteria  `json:"buy_rules"`
            CalculateOn string  `json:"calculate_on"`
         
    }
    
    // PromotionPartialUpdate used by Cart
    type PromotionPartialUpdate struct {

        
            Schedule PromotionSchedule  `json:"schedule"`
            Archive bool  `json:"archive"`
         
    }
    
    // ActivePromosResponse used by Cart
    type ActivePromosResponse struct {

        
            EntityType string  `json:"entity_type"`
            IsHidden bool  `json:"is_hidden"`
            EntitySlug string  `json:"entity_slug"`
            Example string  `json:"example"`
            Description string  `json:"description"`
            CreatedOn string  `json:"created_on"`
            Subtitle string  `json:"subtitle"`
            Type string  `json:"type"`
            ModifiedOn string  `json:"modified_on"`
            Title string  `json:"title"`
         
    }
    
    // Charges used by Cart
    type Charges struct {

        
            Charges float64  `json:"charges"`
            Threshold float64  `json:"threshold"`
         
    }
    
    // DeliveryCharges used by Cart
    type DeliveryCharges struct {

        
            Enabled bool  `json:"enabled"`
            Charges []Charges  `json:"charges"`
         
    }
    
    // CartMetaConfigUpdate used by Cart
    type CartMetaConfigUpdate struct {

        
            MaxCartItems float64  `json:"max_cart_items"`
            Enabled bool  `json:"enabled"`
            GiftDisplayText string  `json:"gift_display_text"`
            MinCartValue float64  `json:"min_cart_value"`
            BulkCoupons bool  `json:"bulk_coupons"`
            RevenueEngineCoupon bool  `json:"revenue_engine_coupon"`
            GiftPricing float64  `json:"gift_pricing"`
            DeliveryCharges DeliveryCharges  `json:"delivery_charges"`
         
    }
    
    // CartMetaConfigAdd used by Cart
    type CartMetaConfigAdd struct {

        
            MaxCartItems float64  `json:"max_cart_items"`
            Enabled bool  `json:"enabled"`
            GiftDisplayText string  `json:"gift_display_text"`
            MinCartValue float64  `json:"min_cart_value"`
            BulkCoupons bool  `json:"bulk_coupons"`
            RevenueEngineCoupon bool  `json:"revenue_engine_coupon"`
            GiftPricing float64  `json:"gift_pricing"`
            DeliveryCharges DeliveryCharges  `json:"delivery_charges"`
         
    }
    
    // CartItem used by Cart
    type CartItem struct {

        
            ProductID string  `json:"product_id"`
            Size string  `json:"size"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // OpenapiCartDetailsRequest used by Cart
    type OpenapiCartDetailsRequest struct {

        
            CartItems []CartItem  `json:"cart_items"`
         
    }
    
    // RawBreakup used by Cart
    type RawBreakup struct {

        
            MrpTotal float64  `json:"mrp_total"`
            GiftCard float64  `json:"gift_card"`
            Coupon float64  `json:"coupon"`
            YouSaved float64  `json:"you_saved"`
            Subtotal float64  `json:"subtotal"`
            Total float64  `json:"total"`
            Discount float64  `json:"discount"`
            FyndCash float64  `json:"fynd_cash"`
            DeliveryCharge float64  `json:"delivery_charge"`
            CodCharge float64  `json:"cod_charge"`
            GstCharges float64  `json:"gst_charges"`
            Vog float64  `json:"vog"`
            ConvenienceFee float64  `json:"convenience_fee"`
         
    }
    
    // DisplayBreakup used by Cart
    type DisplayBreakup struct {

        
            Display string  `json:"display"`
            Key string  `json:"key"`
            CurrencyCode string  `json:"currency_code"`
            Value float64  `json:"value"`
            Message []string  `json:"message"`
            CurrencySymbol string  `json:"currency_symbol"`
         
    }
    
    // CouponBreakup used by Cart
    type CouponBreakup struct {

        
            MinimumCartValue float64  `json:"minimum_cart_value"`
            CouponType string  `json:"coupon_type"`
            IsApplied bool  `json:"is_applied"`
            CouponValue float64  `json:"coupon_value"`
            Value float64  `json:"value"`
            Description string  `json:"description"`
            Code string  `json:"code"`
            Type string  `json:"type"`
            Message string  `json:"message"`
            SubTitle string  `json:"sub_title"`
            MaxDiscountValue float64  `json:"max_discount_value"`
            Title string  `json:"title"`
            UID string  `json:"uid"`
         
    }
    
    // LoyaltyPoints used by Cart
    type LoyaltyPoints struct {

        
            Applicable float64  `json:"applicable"`
            Description string  `json:"description"`
            Total float64  `json:"total"`
            IsApplied bool  `json:"is_applied"`
         
    }
    
    // CartBreakup used by Cart
    type CartBreakup struct {

        
            Raw RawBreakup  `json:"raw"`
            Display []DisplayBreakup  `json:"display"`
            Coupon CouponBreakup  `json:"coupon"`
            LoyaltyPoints LoyaltyPoints  `json:"loyalty_points"`
         
    }
    
    // CouponDetails used by Cart
    type CouponDetails struct {

        
            DiscountSingleQuantity float64  `json:"discount_single_quantity"`
            DiscountTotalQuantity float64  `json:"discount_total_quantity"`
            Code string  `json:"code"`
         
    }
    
    // CartProductIdentifer used by Cart
    type CartProductIdentifer struct {

        
            Identifier string  `json:"identifier"`
         
    }
    
    // ProductPrice used by Cart
    type ProductPrice struct {

        
            Marked float64  `json:"marked"`
            CurrencyCode string  `json:"currency_code"`
            Selling float64  `json:"selling"`
            AddOn float64  `json:"add_on"`
            Effective float64  `json:"effective"`
            CurrencySymbol string  `json:"currency_symbol"`
         
    }
    
    // ProductPriceInfo used by Cart
    type ProductPriceInfo struct {

        
            Base ProductPrice  `json:"base"`
            Converted ProductPrice  `json:"converted"`
         
    }
    
    // ProductAvailabilitySize used by Cart
    type ProductAvailabilitySize struct {

        
            Display string  `json:"display"`
            Value string  `json:"value"`
            IsAvailable bool  `json:"is_available"`
         
    }
    
    // ProductAvailability used by Cart
    type ProductAvailability struct {

        
            IsValid bool  `json:"is_valid"`
            OtherStoreQuantity float64  `json:"other_store_quantity"`
            OutOfStock bool  `json:"out_of_stock"`
            Sizes []string  `json:"sizes"`
            AvailableSizes []ProductAvailabilitySize  `json:"available_sizes"`
            Deliverable bool  `json:"deliverable"`
         
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
    
    // ProductArticle used by Cart
    type ProductArticle struct {

        
            SellerIdentifier string  `json:"seller_identifier"`
            GiftCard map[string]interface{}  `json:"gift_card"`
            Size string  `json:"size"`
            IsGiftVisible bool  `json:"is_gift_visible"`
            ProductGroupTags []string  `json:"product_group_tags"`
            Identifier map[string]interface{}  `json:"identifier"`
            Price ArticlePriceInfo  `json:"price"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Quantity float64  `json:"quantity"`
            Type string  `json:"type"`
            CartItemMeta map[string]interface{}  `json:"cart_item_meta"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            Seller BaseInfo  `json:"seller"`
            Store StoreInfo  `json:"store"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            UID string  `json:"uid"`
         
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
            URL string  `json:"url"`
            Type string  `json:"type"`
         
    }
    
    // NetQuantity used by Cart
    type NetQuantity struct {

        
            Unit string  `json:"unit"`
            Value string  `json:"value"`
         
    }
    
    // CartProduct used by Cart
    type CartProduct struct {

        
            Brand BaseInfo  `json:"brand"`
            Tags []string  `json:"tags"`
            Slug string  `json:"slug"`
            ItemCode string  `json:"item_code"`
            Name string  `json:"name"`
            TeaserTag Tags  `json:"teaser_tag"`
            Categories []CategoryInfo  `json:"categories"`
            Images []ProductImage  `json:"images"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Type string  `json:"type"`
            Action ProductAction  `json:"action"`
            NetQuantity NetQuantity  `json:"net_quantity"`
            UID float64  `json:"uid"`
         
    }
    
    // Ownership2 used by Cart
    type Ownership2 struct {

        
            PayableCategory string  `json:"payable_category"`
            PayableBy string  `json:"payable_by"`
         
    }
    
    // FreeGiftItem used by Cart
    type FreeGiftItem struct {

        
            ItemName string  `json:"item_name"`
            ItemPriceDetails map[string]interface{}  `json:"item_price_details"`
            ItemID float64  `json:"item_id"`
            ItemBrandName string  `json:"item_brand_name"`
            ItemSlug string  `json:"item_slug"`
            ItemImagesURL []string  `json:"item_images_url"`
         
    }
    
    // AppliedFreeArticles used by Cart
    type AppliedFreeArticles struct {

        
            ParentItemIdentifier string  `json:"parent_item_identifier"`
            FreeGiftItemDetails FreeGiftItem  `json:"free_gift_item_details"`
            Quantity float64  `json:"quantity"`
            ArticleID string  `json:"article_id"`
         
    }
    
    // DiscountRulesApp used by Cart
    type DiscountRulesApp struct {

        
            MatchedBuyRules []string  `json:"matched_buy_rules"`
            ItemCriteria map[string]interface{}  `json:"item_criteria"`
            Offer map[string]interface{}  `json:"offer"`
            RawOffer map[string]interface{}  `json:"raw_offer"`
         
    }
    
    // BuyRules used by Cart
    type BuyRules struct {

        
            CartConditions map[string]interface{}  `json:"cart_conditions"`
            ItemCriteria map[string]interface{}  `json:"item_criteria"`
         
    }
    
    // AppliedPromotion used by Cart
    type AppliedPromotion struct {

        
            Amount float64  `json:"amount"`
            PromotionName string  `json:"promotion_name"`
            Ownership Ownership2  `json:"ownership"`
            ArticleQuantity float64  `json:"article_quantity"`
            AppliedFreeArticles []AppliedFreeArticles  `json:"applied_free_articles"`
            OfferText string  `json:"offer_text"`
            DiscountRules []DiscountRulesApp  `json:"discount_rules"`
            BuyRules []BuyRules  `json:"buy_rules"`
            PromotionType string  `json:"promotion_type"`
            MrpPromotion bool  `json:"mrp_promotion"`
            PromotionGroup string  `json:"promotion_group"`
            PromoID string  `json:"promo_id"`
         
    }
    
    // PromoMeta used by Cart
    type PromoMeta struct {

        
            Message string  `json:"message"`
         
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
    
    // CartProductInfo used by Cart
    type CartProductInfo struct {

        
            IsSet bool  `json:"is_set"`
            Moq map[string]interface{}  `json:"moq"`
            Coupon CouponDetails  `json:"coupon"`
            Message string  `json:"message"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
            CouponMessage string  `json:"coupon_message"`
            Price ProductPriceInfo  `json:"price"`
            Availability ProductAvailability  `json:"availability"`
            BulkOffer map[string]interface{}  `json:"bulk_offer"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            Article ProductArticle  `json:"article"`
            Key string  `json:"key"`
            Quantity float64  `json:"quantity"`
            Product CartProduct  `json:"product"`
            PromotionsApplied []AppliedPromotion  `json:"promotions_applied"`
            PricePerUnit ProductPriceInfo  `json:"price_per_unit"`
            PromoMeta PromoMeta  `json:"promo_meta"`
            Discount string  `json:"discount"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
         
    }
    
    // OpenapiCartDetailsResponse used by Cart
    type OpenapiCartDetailsResponse struct {

        
            BreakupValues CartBreakup  `json:"breakup_values"`
            IsValid bool  `json:"is_valid"`
            Items []CartProductInfo  `json:"items"`
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

        
            Landmark string  `json:"landmark"`
            City string  `json:"city"`
            Phone float64  `json:"phone"`
            State string  `json:"state"`
            Country string  `json:"country"`
            Name string  `json:"name"`
            Address string  `json:"address"`
            AddressType string  `json:"address_type"`
            CountryPhoneCode string  `json:"country_phone_code"`
            Email string  `json:"email"`
            CountryIsoCode string  `json:"country_iso_code"`
            Area string  `json:"area"`
            Pincode float64  `json:"pincode"`
            AreaCodeSlug string  `json:"area_code_slug"`
            AreaCode string  `json:"area_code"`
            CountryCode string  `json:"country_code"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // OpenApiCartServiceabilityRequest used by Cart
    type OpenApiCartServiceabilityRequest struct {

        
            CartItems []CartItem  `json:"cart_items"`
            ShippingAddress ShippingAddress  `json:"shipping_address"`
         
    }
    
    // OpenApiCartServiceabilityResponse used by Cart
    type OpenApiCartServiceabilityResponse struct {

        
            Items []CartProductInfo  `json:"items"`
            IsValid bool  `json:"is_valid"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            Message string  `json:"message"`
            BreakupValues CartBreakup  `json:"breakup_values"`
         
    }
    
    // OpenApiFiles used by Cart
    type OpenApiFiles struct {

        
            Key string  `json:"key"`
            Values []string  `json:"values"`
         
    }
    
    // MultiTenderPaymentMeta used by Cart
    type MultiTenderPaymentMeta struct {

        
            PaymentID string  `json:"payment_id"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            PaymentGateway string  `json:"payment_gateway"`
            CurrentStatus string  `json:"current_status"`
            OrderID string  `json:"order_id"`
         
    }
    
    // MultiTenderPaymentMethod used by Cart
    type MultiTenderPaymentMethod struct {

        
            Mode string  `json:"mode"`
            Amount float64  `json:"amount"`
            Name string  `json:"name"`
            Meta MultiTenderPaymentMeta  `json:"meta"`
         
    }
    
    // CartItemMeta used by Cart
    type CartItemMeta struct {

        
            PrimaryItem bool  `json:"primary_item"`
            GroupID string  `json:"group_id"`
         
    }
    
    // OpenApiOrderItem used by Cart
    type OpenApiOrderItem struct {

        
            Files []OpenApiFiles  `json:"files"`
            PaymentMethods []MultiTenderPaymentMethod  `json:"payment_methods"`
            CodCharges float64  `json:"cod_charges"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            AmountPaid float64  `json:"amount_paid"`
            ProductID float64  `json:"product_id"`
            Discount float64  `json:"discount"`
            PriceEffective float64  `json:"price_effective"`
            PriceMarked float64  `json:"price_marked"`
            CashbackApplied float64  `json:"cashback_applied"`
            DeliveryCharges float64  `json:"delivery_charges"`
            Quantity float64  `json:"quantity"`
            Size string  `json:"size"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            Meta CartItemMeta  `json:"meta"`
            EmployeeDiscount float64  `json:"employee_discount"`
         
    }
    
    // OpenApiPlatformCheckoutReq used by Cart
    type OpenApiPlatformCheckoutReq struct {

        
            Files []OpenApiFiles  `json:"files"`
            PaymentMode string  `json:"payment_mode"`
            CodCharges float64  `json:"cod_charges"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            Coupon string  `json:"coupon"`
            BillingAddress ShippingAddress  `json:"billing_address"`
            CouponCode string  `json:"coupon_code"`
            CurrencyCode string  `json:"currency_code"`
            ShippingAddress ShippingAddress  `json:"shipping_address"`
            CashbackApplied float64  `json:"cashback_applied"`
            OrderID string  `json:"order_id"`
            CartItems []OpenApiOrderItem  `json:"cart_items"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            CartValue float64  `json:"cart_value"`
            Gstin string  `json:"gstin"`
            PaymentMethods []MultiTenderPaymentMethod  `json:"payment_methods"`
            CouponValue float64  `json:"coupon_value"`
            Comment string  `json:"comment"`
            DeliveryCharges float64  `json:"delivery_charges"`
            EmployeeDiscount map[string]interface{}  `json:"employee_discount"`
         
    }
    
    // OpenApiCheckoutResponse used by Cart
    type OpenApiCheckoutResponse struct {

        
            OrderID string  `json:"order_id"`
            OrderRefID string  `json:"order_ref_id"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // AbandonedCart used by Cart
    type AbandonedCart struct {

        
            PaymentMode string  `json:"payment_mode"`
            CodCharges map[string]interface{}  `json:"cod_charges"`
            LastModified string  `json:"last_modified"`
            Coupon map[string]interface{}  `json:"coupon"`
            Payments map[string]interface{}  `json:"payments"`
            BulkCouponDiscount float64  `json:"bulk_coupon_discount"`
            Shipments []map[string]interface{}  `json:"shipments"`
            ID string  `json:"_id"`
            IsArchive bool  `json:"is_archive"`
            UID float64  `json:"uid"`
            Meta map[string]interface{}  `json:"meta"`
            Gstin string  `json:"gstin"`
            ExpireAt string  `json:"expire_at"`
            FyndCredits map[string]interface{}  `json:"fynd_credits"`
            DeliveryCharges map[string]interface{}  `json:"delivery_charges"`
            CreatedOn string  `json:"created_on"`
            AppID string  `json:"app_id"`
            IsActive bool  `json:"is_active"`
            MergeQty bool  `json:"merge_qty"`
            FcIndexMap []float64  `json:"fc_index_map"`
            BuyNow bool  `json:"buy_now"`
            OrderID string  `json:"order_id"`
            Articles []map[string]interface{}  `json:"articles"`
            Cashback map[string]interface{}  `json:"cashback"`
            CheckoutMode string  `json:"checkout_mode"`
            CartValue float64  `json:"cart_value"`
            PaymentMethods []map[string]interface{}  `json:"payment_methods"`
            Promotion map[string]interface{}  `json:"promotion"`
            IsDefault bool  `json:"is_default"`
            Discount float64  `json:"discount"`
            Comment string  `json:"comment"`
            PickUpCustomerDetails map[string]interface{}  `json:"pick_up_customer_details"`
            UserID string  `json:"user_id"`
         
    }
    
    // AbandonedCartResponse used by Cart
    type AbandonedCartResponse struct {

        
            Page Page  `json:"page"`
            Items []AbandonedCart  `json:"items"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
            Result map[string]interface{}  `json:"result"`
         
    }
    
    // CartCurrency used by Cart
    type CartCurrency struct {

        
            Symbol string  `json:"symbol"`
            Code string  `json:"code"`
         
    }
    
    // PaymentSelectionLock used by Cart
    type PaymentSelectionLock struct {

        
            DefaultOptions string  `json:"default_options"`
            Enabled bool  `json:"enabled"`
            PaymentIdentifier string  `json:"payment_identifier"`
         
    }
    
    // CartDetailResponse used by Cart
    type CartDetailResponse struct {

        
            AppliedPromoDetails []AppliedPromotion  `json:"applied_promo_details"`
            LastModified string  `json:"last_modified"`
            Message string  `json:"message"`
            ID string  `json:"id"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            Gstin string  `json:"gstin"`
            CouponText string  `json:"coupon_text"`
            Items []CartProductInfo  `json:"items"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            PanConfig map[string]interface{}  `json:"pan_config"`
            Currency CartCurrency  `json:"currency"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            BuyNow bool  `json:"buy_now"`
            PanNo string  `json:"pan_no"`
            CheckoutMode string  `json:"checkout_mode"`
            IsValid bool  `json:"is_valid"`
            Comment string  `json:"comment"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
         
    }
    
    // AddProductCart used by Cart
    type AddProductCart struct {

        
            Display string  `json:"display"`
            ArticleAssignment map[string]interface{}  `json:"article_assignment"`
            ItemID float64  `json:"item_id"`
            SellerID float64  `json:"seller_id"`
            Pos bool  `json:"pos"`
            ProductGroupTags []string  `json:"product_group_tags"`
            StoreID float64  `json:"store_id"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Quantity float64  `json:"quantity"`
            ArticleID string  `json:"article_id"`
            ItemSize string  `json:"item_size"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            ParentItemIdentifiers []map[string]string  `json:"parent_item_identifiers"`
         
    }
    
    // AddCartRequest used by Cart
    type AddCartRequest struct {

        
            Items []AddProductCart  `json:"items"`
            NewCart bool  `json:"new_cart"`
         
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

        
            ItemIndex float64  `json:"item_index"`
            ItemID float64  `json:"item_id"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ArticleID string  `json:"article_id"`
            Quantity float64  `json:"quantity"`
            ItemSize string  `json:"item_size"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
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

        
            User map[string]interface{}  `json:"user"`
            Token string  `json:"token"`
            Source map[string]interface{}  `json:"source"`
            CreatedOn string  `json:"created_on"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // SharedCart used by Cart
    type SharedCart struct {

        
            LastModified string  `json:"last_modified"`
            Message string  `json:"message"`
            ID string  `json:"id"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            UID string  `json:"uid"`
            Gstin string  `json:"gstin"`
            CouponText string  `json:"coupon_text"`
            Items []CartProductInfo  `json:"items"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            Currency CartCurrency  `json:"currency"`
            SharedCartDetails SharedCartDetails  `json:"shared_cart_details"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            BuyNow bool  `json:"buy_now"`
            CheckoutMode string  `json:"checkout_mode"`
            IsValid bool  `json:"is_valid"`
            Comment string  `json:"comment"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            CartID float64  `json:"cart_id"`
         
    }
    
    // SharedCartResponse used by Cart
    type SharedCartResponse struct {

        
            Error string  `json:"error"`
            Cart SharedCart  `json:"cart"`
         
    }
    
    // CartList used by Cart
    type CartList struct {

        
            ItemCounts float64  `json:"item_counts"`
            CartValue float64  `json:"cart_value"`
            CreatedOn string  `json:"created_on"`
            PickUpCustomerDetails map[string]interface{}  `json:"pick_up_customer_details"`
            CartID string  `json:"cart_id"`
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

        
            LastName string  `json:"last_name"`
            ExternalID string  `json:"external_id"`
            CreatedAt string  `json:"created_at"`
            Gender string  `json:"gender"`
            FirstName string  `json:"first_name"`
            ModifiedOn string  `json:"modified_on"`
            Mobile string  `json:"mobile"`
            ID string  `json:"_id"`
            UID string  `json:"uid"`
         
    }
    
    // UserCartMappingResponse used by Cart
    type UserCartMappingResponse struct {

        
            AppliedPromoDetails []AppliedPromotion  `json:"applied_promo_details"`
            LastModified string  `json:"last_modified"`
            Message string  `json:"message"`
            ID string  `json:"id"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            Gstin string  `json:"gstin"`
            CouponText string  `json:"coupon_text"`
            Items []CartProductInfo  `json:"items"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            PanConfig map[string]interface{}  `json:"pan_config"`
            Currency CartCurrency  `json:"currency"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            BuyNow bool  `json:"buy_now"`
            PanNo string  `json:"pan_no"`
            CheckoutMode string  `json:"checkout_mode"`
            User UserInfo  `json:"user"`
            IsValid bool  `json:"is_valid"`
            Comment string  `json:"comment"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
         
    }
    
    // DeleteCartRequest used by Cart
    type DeleteCartRequest struct {

        
            CartIDList []string  `json:"cart_id_list"`
         
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
    
    // PageCoupon used by Cart
    type PageCoupon struct {

        
            HasPrevious bool  `json:"has_previous"`
            HasNext bool  `json:"has_next"`
            Total float64  `json:"total"`
            TotalItemCount float64  `json:"total_item_count"`
            Current float64  `json:"current"`
         
    }
    
    // Coupon used by Cart
    type Coupon struct {

        
            Title string  `json:"title"`
            MinimumCartValue float64  `json:"minimum_cart_value"`
            IsApplied bool  `json:"is_applied"`
            CouponValue float64  `json:"coupon_value"`
            CouponCode string  `json:"coupon_code"`
            IsApplicable bool  `json:"is_applicable"`
            ExpiresOn string  `json:"expires_on"`
            Description string  `json:"description"`
            Message string  `json:"message"`
            SubTitle string  `json:"sub_title"`
            MaxDiscountValue float64  `json:"max_discount_value"`
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

        
            Longitude float64  `json:"longitude"`
            Latitude float64  `json:"latitude"`
         
    }
    
    // PlatformAddress used by Cart
    type PlatformAddress struct {

        
            Country string  `json:"country"`
            AreaCodeSlug string  `json:"area_code_slug"`
            ID string  `json:"id"`
            Meta map[string]interface{}  `json:"meta"`
            Landmark string  `json:"landmark"`
            City string  `json:"city"`
            CreatedByUserID string  `json:"created_by_user_id"`
            Address string  `json:"address"`
            Email string  `json:"email"`
            Phone string  `json:"phone"`
            Tags []string  `json:"tags"`
            Name string  `json:"name"`
            AddressType string  `json:"address_type"`
            IsActive bool  `json:"is_active"`
            Area string  `json:"area"`
            GeoLocation GeoLocation  `json:"geo_location"`
            CountryCode string  `json:"country_code"`
            AreaCode string  `json:"area_code"`
            CheckoutMode string  `json:"checkout_mode"`
            GoogleMapPoint map[string]interface{}  `json:"google_map_point"`
            IsDefaultAddress bool  `json:"is_default_address"`
            State string  `json:"state"`
            CartID string  `json:"cart_id"`
            UserID string  `json:"user_id"`
         
    }
    
    // PlatformGetAddressesResponse used by Cart
    type PlatformGetAddressesResponse struct {

        
            Address []PlatformAddress  `json:"address"`
         
    }
    
    // SaveAddressResponse used by Cart
    type SaveAddressResponse struct {

        
            ID string  `json:"id"`
            IsDefaultAddress bool  `json:"is_default_address"`
            Success bool  `json:"success"`
         
    }
    
    // UpdateAddressResponse used by Cart
    type UpdateAddressResponse struct {

        
            ID string  `json:"id"`
            IsDefaultAddress bool  `json:"is_default_address"`
            IsUpdated bool  `json:"is_updated"`
            Success bool  `json:"success"`
         
    }
    
    // DeleteAddressResponse used by Cart
    type DeleteAddressResponse struct {

        
            ID string  `json:"id"`
            IsDeleted bool  `json:"is_deleted"`
         
    }
    
    // PlatformSelectCartAddressRequest used by Cart
    type PlatformSelectCartAddressRequest struct {

        
            CartID string  `json:"cart_id"`
            ID string  `json:"id"`
            UserID string  `json:"user_id"`
            BillingAddressID string  `json:"billing_address_id"`
            CheckoutMode string  `json:"checkout_mode"`
         
    }
    
    // ShipmentArticle used by Cart
    type ShipmentArticle struct {

        
            Quantity string  `json:"quantity"`
            ArticleID string  `json:"article_id"`
            Meta string  `json:"meta"`
         
    }
    
    // PlatformShipmentResponse used by Cart
    type PlatformShipmentResponse struct {

        
            ShipmentType string  `json:"shipment_type"`
            DpID string  `json:"dp_id"`
            Promise ShipmentPromise  `json:"promise"`
            FulfillmentType string  `json:"fulfillment_type"`
            Items []CartProductInfo  `json:"items"`
            BoxType string  `json:"box_type"`
            FulfillmentID float64  `json:"fulfillment_id"`
            OrderType string  `json:"order_type"`
            Shipments float64  `json:"shipments"`
            DpOptions map[string]interface{}  `json:"dp_options"`
            Articles []ShipmentArticle  `json:"articles"`
         
    }
    
    // PlatformCartShipmentsResponse used by Cart
    type PlatformCartShipmentsResponse struct {

        
            AppliedPromoDetails []AppliedPromotion  `json:"applied_promo_details"`
            LastModified string  `json:"last_modified"`
            Error bool  `json:"error"`
            Message string  `json:"message"`
            ID string  `json:"id"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            Shipments []PlatformShipmentResponse  `json:"shipments"`
            Gstin string  `json:"gstin"`
            CouponText string  `json:"coupon_text"`
            Items []CartProductInfo  `json:"items"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            PanConfig map[string]interface{}  `json:"pan_config"`
            Currency CartCurrency  `json:"currency"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            BuyNow bool  `json:"buy_now"`
            PanNo string  `json:"pan_no"`
            CheckoutMode string  `json:"checkout_mode"`
            IsValid bool  `json:"is_valid"`
            Comment string  `json:"comment"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
         
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
            Comment string  `json:"comment"`
            GiftDetails map[string]interface{}  `json:"gift_details"`
            PickUpCustomerDetails map[string]interface{}  `json:"pick_up_customer_details"`
            PanNo string  `json:"pan_no"`
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

        
            User string  `json:"user"`
            LastName string  `json:"last_name"`
            FirstName string  `json:"first_name"`
            EmployeeCode string  `json:"employee_code"`
            ID string  `json:"_id"`
         
    }
    
    // Files used by Cart
    type Files struct {

        
            Key string  `json:"key"`
            Values []string  `json:"values"`
         
    }
    
    // PlatformCartCheckoutDetailRequest used by Cart
    type PlatformCartCheckoutDetailRequest struct {

        
            PaymentMode string  `json:"payment_mode"`
            Pos bool  `json:"pos"`
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
            Aggregator string  `json:"aggregator"`
            DeviceID string  `json:"device_id"`
            MerchantCode string  `json:"merchant_code"`
            Staff StaffCheckout  `json:"staff"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            ID string  `json:"id"`
            Meta map[string]interface{}  `json:"meta"`
            PaymentIdentifier string  `json:"payment_identifier"`
            PickAtStoreUID float64  `json:"pick_at_store_uid"`
            CallbackURL string  `json:"callback_url"`
            OrderingStore float64  `json:"ordering_store"`
            PaymentAutoConfirm bool  `json:"payment_auto_confirm"`
            AddressID string  `json:"address_id"`
            Files []Files  `json:"files"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
            CheckoutMode string  `json:"checkout_mode"`
            PaymentParams map[string]interface{}  `json:"payment_params"`
            OrderType string  `json:"order_type"`
            EmployeeCode string  `json:"employee_code"`
            UserID string  `json:"user_id"`
            BillingAddressID string  `json:"billing_address_id"`
         
    }
    
    // CheckCart used by Cart
    type CheckCart struct {

        
            CodCharges float64  `json:"cod_charges"`
            LastModified string  `json:"last_modified"`
            ErrorMessage string  `json:"error_message"`
            DeliveryChargeOrderValue float64  `json:"delivery_charge_order_value"`
            Success bool  `json:"success"`
            CodMessage string  `json:"cod_message"`
            StoreEmps []map[string]interface{}  `json:"store_emps"`
            Message string  `json:"message"`
            ID string  `json:"id"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            UserType string  `json:"user_type"`
            UID string  `json:"uid"`
            Gstin string  `json:"gstin"`
            CouponText string  `json:"coupon_text"`
            Items []CartProductInfo  `json:"items"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            Currency CartCurrency  `json:"currency"`
            DeliveryCharges float64  `json:"delivery_charges"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            BuyNow bool  `json:"buy_now"`
            OrderID string  `json:"order_id"`
            CheckoutMode string  `json:"checkout_mode"`
            CodAvailable bool  `json:"cod_available"`
            StoreCode string  `json:"store_code"`
            IsValid bool  `json:"is_valid"`
            Comment string  `json:"comment"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            CartID float64  `json:"cart_id"`
         
    }
    
    // CartCheckoutResponse used by Cart
    type CartCheckoutResponse struct {

        
            AppInterceptURL string  `json:"app_intercept_url"`
            Data map[string]interface{}  `json:"data"`
            PaymentConfirmURL string  `json:"payment_confirm_url"`
            Success bool  `json:"success"`
            Cart CheckCart  `json:"cart"`
            CallbackURL string  `json:"callback_url"`
            Message string  `json:"message"`
            OrderID string  `json:"order_id"`
         
    }
    
    // CartDeliveryModesResponse used by Cart
    type CartDeliveryModesResponse struct {

        
            PickupStores []float64  `json:"pickup_stores"`
            AvailableModes []string  `json:"available_modes"`
         
    }
    
    // PickupStoreDetail used by Cart
    type PickupStoreDetail struct {

        
            Landmark string  `json:"landmark"`
            City string  `json:"city"`
            Phone string  `json:"phone"`
            Country string  `json:"country"`
            State string  `json:"state"`
            Name string  `json:"name"`
            StoreCode string  `json:"store_code"`
            Address string  `json:"address"`
            AddressType string  `json:"address_type"`
            Email string  `json:"email"`
            Area string  `json:"area"`
            Pincode float64  `json:"pincode"`
            AreaCodeSlug string  `json:"area_code_slug"`
            AreaCode string  `json:"area_code"`
            ID float64  `json:"id"`
            UID float64  `json:"uid"`
         
    }
    
    // StoreDetailsResponse used by Cart
    type StoreDetailsResponse struct {

        
            Items []PickupStoreDetail  `json:"items"`
         
    }
    
    // UpdateCartPaymentRequest used by Cart
    type UpdateCartPaymentRequest struct {

        
            PaymentMode string  `json:"payment_mode"`
            PaymentIdentifier string  `json:"payment_identifier"`
            MerchantCode string  `json:"merchant_code"`
            ID string  `json:"id"`
            AddressID string  `json:"address_id"`
            AggregatorName string  `json:"aggregator_name"`
         
    }
    
    // CouponValidity used by Cart
    type CouponValidity struct {

        
            Valid bool  `json:"valid"`
            Discount float64  `json:"discount"`
            DisplayMessageEn string  `json:"display_message_en"`
            Code string  `json:"code"`
            Title string  `json:"title"`
         
    }
    
    // PaymentCouponValidate used by Cart
    type PaymentCouponValidate struct {

        
            CouponValidity CouponValidity  `json:"coupon_validity"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // PaymentMeta used by Cart
    type PaymentMeta struct {

        
            PaymentIdentifier string  `json:"payment_identifier"`
            PaymentGateway string  `json:"payment_gateway"`
            MerchantCode string  `json:"merchant_code"`
            Type string  `json:"type"`
         
    }
    
    // PaymentMethod used by Cart
    type PaymentMethod struct {

        
            Amount float64  `json:"amount"`
            Payment string  `json:"payment"`
            PaymentMeta PaymentMeta  `json:"payment_meta"`
            Name string  `json:"name"`
            Mode string  `json:"mode"`
         
    }
    
    // PlatformCartCheckoutDetailV2Request used by Cart
    type PlatformCartCheckoutDetailV2Request struct {

        
            PaymentMode string  `json:"payment_mode"`
            Pos bool  `json:"pos"`
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
            Aggregator string  `json:"aggregator"`
            DeviceID string  `json:"device_id"`
            MerchantCode string  `json:"merchant_code"`
            Staff StaffCheckout  `json:"staff"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            ID string  `json:"id"`
            CustomMeta map[string]interface{}  `json:"custom_meta"`
            Meta map[string]interface{}  `json:"meta"`
            PaymentIdentifier string  `json:"payment_identifier"`
            PickAtStoreUID float64  `json:"pick_at_store_uid"`
            CallbackURL string  `json:"callback_url"`
            OrderingStore float64  `json:"ordering_store"`
            PaymentAutoConfirm bool  `json:"payment_auto_confirm"`
            AddressID string  `json:"address_id"`
            Files []Files  `json:"files"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
            CheckoutMode string  `json:"checkout_mode"`
            PaymentParams map[string]interface{}  `json:"payment_params"`
            PaymentMethods []PaymentMethod  `json:"payment_methods"`
            OrderType string  `json:"order_type"`
            EmployeeCode string  `json:"employee_code"`
            UserID string  `json:"user_id"`
            BillingAddressID string  `json:"billing_address_id"`
         
    }
    
    // UpdateCartPaymentRequestV2 used by Cart
    type UpdateCartPaymentRequestV2 struct {

        
            PaymentMode string  `json:"payment_mode"`
            PaymentMethods []PaymentMethod  `json:"payment_methods"`
            PaymentIdentifier string  `json:"payment_identifier"`
            MerchantCode string  `json:"merchant_code"`
            ID string  `json:"id"`
            AddressID string  `json:"address_id"`
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
    

    
    // ApplicationServiceabilityConfig used by Logistics
    type ApplicationServiceabilityConfig struct {

        
            ChannelType string  `json:"channel_type"`
            ServiceabilityType string  `json:"serviceability_type"`
            ChannelID string  `json:"channel_id"`
         
    }
    
    // ServiceabilityrErrorResponse used by Logistics
    type ServiceabilityrErrorResponse struct {

        
            Value string  `json:"value"`
            Message string  `json:"message"`
            Type string  `json:"type"`
         
    }
    
    // ApplicationServiceabilityConfigResponse used by Logistics
    type ApplicationServiceabilityConfigResponse struct {

        
            Data ApplicationServiceabilityConfig  `json:"data"`
            Error ServiceabilityrErrorResponse  `json:"error"`
            Success bool  `json:"success"`
         
    }
    
    // EntityRegionView_Request used by Logistics
    type EntityRegionView_Request struct {

        
            ParentID []string  `json:"parent_id"`
            SubType []string  `json:"sub_type"`
         
    }
    
    // EntityRegionView_Items used by Logistics
    type EntityRegionView_Items struct {

        
            UID string  `json:"uid"`
            Name string  `json:"name"`
            SubType string  `json:"sub_type"`
         
    }
    
    // EntityRegionView_Error used by Logistics
    type EntityRegionView_Error struct {

        
            Value string  `json:"value"`
            Message string  `json:"message"`
            Type string  `json:"type"`
         
    }
    
    // EntityRegionView_page used by Logistics
    type EntityRegionView_page struct {

        
            ItemTotal float64  `json:"item_total"`
            HasNext bool  `json:"has_next"`
            Size float64  `json:"size"`
            Type string  `json:"type"`
            Current float64  `json:"current"`
         
    }
    
    // EntityRegionView_Response used by Logistics
    type EntityRegionView_Response struct {

        
            Data []EntityRegionView_Items  `json:"data"`
            Success bool  `json:"success"`
            Error EntityRegionView_Error  `json:"error"`
            Page EntityRegionView_page  `json:"page"`
         
    }
    
    // ZoneDataItem used by Logistics
    type ZoneDataItem struct {

        
            ItemTotal float64  `json:"item_total"`
            HasNext bool  `json:"has_next"`
            Size float64  `json:"size"`
            Type string  `json:"type"`
            Current float64  `json:"current"`
         
    }
    
    // ListViewChannels used by Logistics
    type ListViewChannels struct {

        
            ChannelType string  `json:"channel_type"`
            ChannelID string  `json:"channel_id"`
         
    }
    
    // ListViewProduct used by Logistics
    type ListViewProduct struct {

        
            Type string  `json:"type"`
            Count float64  `json:"count"`
         
    }
    
    // ListViewItems used by Logistics
    type ListViewItems struct {

        
            Name string  `json:"name"`
            IsActive bool  `json:"is_active"`
            ZoneID string  `json:"zone_id"`
            Slug string  `json:"slug"`
            Channels ListViewChannels  `json:"channels"`
            StoresCount float64  `json:"stores_count"`
            Product ListViewProduct  `json:"product"`
            CompanyID float64  `json:"company_id"`
            PincodesCount float64  `json:"pincodes_count"`
         
    }
    
    // ListViewSummary used by Logistics
    type ListViewSummary struct {

        
            TotalActiveZones float64  `json:"total_active_zones"`
            TotalPincodesServed float64  `json:"total_pincodes_served"`
            TotalZones float64  `json:"total_zones"`
         
    }
    
    // ListViewResponse used by Logistics
    type ListViewResponse struct {

        
            Page []ZoneDataItem  `json:"page"`
            Items []ListViewItems  `json:"items"`
            Summary []ListViewSummary  `json:"summary"`
         
    }
    
    // CompanyStoreView_PageItems used by Logistics
    type CompanyStoreView_PageItems struct {

        
            ItemTotal float64  `json:"item_total"`
            HasNext bool  `json:"has_next"`
            Size float64  `json:"size"`
            Type string  `json:"type"`
            Current float64  `json:"current"`
         
    }
    
    // CompanyStoreView_Response used by Logistics
    type CompanyStoreView_Response struct {

        
            Page []CompanyStoreView_PageItems  `json:"page"`
            Items []map[string]interface{}  `json:"items"`
         
    }
    
    // GetZoneDataViewChannels used by Logistics
    type GetZoneDataViewChannels struct {

        
            ChannelType string  `json:"channel_type"`
            ChannelID string  `json:"channel_id"`
         
    }
    
    // ZoneProductTypes used by Logistics
    type ZoneProductTypes struct {

        
            Tags []string  `json:"tags"`
            Type string  `json:"type"`
         
    }
    
    // ZoneMappingType used by Logistics
    type ZoneMappingType struct {

        
            Pincode []string  `json:"pincode"`
            State []string  `json:"state"`
            Country string  `json:"country"`
         
    }
    
    // UpdateZoneData used by Logistics
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
    
    // ZoneUpdateRequest used by Logistics
    type ZoneUpdateRequest struct {

        
            Identifier string  `json:"identifier"`
            Data UpdateZoneData  `json:"data"`
         
    }
    
    // ZoneSuccessResponse used by Logistics
    type ZoneSuccessResponse struct {

        
            StatusCode float64  `json:"status_code"`
            Success bool  `json:"success"`
         
    }
    
    // GetZoneDataViewItems used by Logistics
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
    
    // GetSingleZoneDataViewResponse used by Logistics
    type GetSingleZoneDataViewResponse struct {

        
            Data GetZoneDataViewItems  `json:"data"`
         
    }
    
    // CreateZoneData used by Logistics
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
    
    // ZoneRequest used by Logistics
    type ZoneRequest struct {

        
            Identifier string  `json:"identifier"`
            Data CreateZoneData  `json:"data"`
         
    }
    
    // ZoneResponse used by Logistics
    type ZoneResponse struct {

        
            StatusCode float64  `json:"status_code"`
            Success bool  `json:"success"`
            ZoneID string  `json:"zone_id"`
         
    }
    
    // GetZoneFromPincodeViewRequest used by Logistics
    type GetZoneFromPincodeViewRequest struct {

        
            Pincode string  `json:"pincode"`
            Country string  `json:"country"`
         
    }
    
    // GetZoneFromPincodeViewResponse used by Logistics
    type GetZoneFromPincodeViewResponse struct {

        
            ServiceabilityType string  `json:"serviceability_type"`
            Zones []string  `json:"zones"`
         
    }
    
    // GetZoneFromApplicationIdViewResponse used by Logistics
    type GetZoneFromApplicationIdViewResponse struct {

        
            Page []ZoneDataItem  `json:"page"`
            Items []ListViewItems  `json:"items"`
         
    }
    
    // ServiceabilityErrorResponse used by Logistics
    type ServiceabilityErrorResponse struct {

        
            Value string  `json:"value"`
            Message string  `json:"message"`
            Type string  `json:"type"`
         
    }
    
    // PageResponse used by Logistics
    type PageResponse struct {

        
            ItemTotal float64  `json:"item_total"`
            HasNext bool  `json:"has_next"`
            Size float64  `json:"size"`
            Type string  `json:"type"`
            Current float64  `json:"current"`
         
    }
    
    // WarningsResponse used by Logistics
    type WarningsResponse struct {

        
            StoreAddress string  `json:"store_address"`
         
    }
    
    // Dp used by Logistics
    type Dp struct {

        
            ExternalAccountID string  `json:"external_account_id"`
            PaymentMode string  `json:"payment_mode"`
            InternalAccountID string  `json:"internal_account_id"`
            TransportMode string  `json:"transport_mode"`
            FmPriority float64  `json:"fm_priority"`
            LmPriority float64  `json:"lm_priority"`
            AssignDpFromSb bool  `json:"assign_dp_from_sb"`
            AreaCode float64  `json:"area_code"`
            Operations []string  `json:"operations"`
            RvpPriority float64  `json:"rvp_priority"`
         
    }
    
    // LogisticsResponse used by Logistics
    type LogisticsResponse struct {

        
            Dp Dp  `json:"dp"`
            Override bool  `json:"override"`
         
    }
    
    // ContactNumberResponse used by Logistics
    type ContactNumberResponse struct {

        
            CountryCode float64  `json:"country_code"`
            Number string  `json:"number"`
         
    }
    
    // ProductReturnConfigResponse used by Logistics
    type ProductReturnConfigResponse struct {

        
            OnSameStore bool  `json:"on_same_store"`
         
    }
    
    // CreatedByResponse used by Logistics
    type CreatedByResponse struct {

        
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
         
    }
    
    // EinvoiceResponse used by Logistics
    type EinvoiceResponse struct {

        
            Enabled bool  `json:"enabled"`
         
    }
    
    // EwayBillResponse used by Logistics
    type EwayBillResponse struct {

        
            Enabled bool  `json:"enabled"`
         
    }
    
    // GstCredentialsResponse used by Logistics
    type GstCredentialsResponse struct {

        
            EInvoice EinvoiceResponse  `json:"e_invoice"`
            EWaybill EwayBillResponse  `json:"e_waybill"`
         
    }
    
    // MobileNo used by Logistics
    type MobileNo struct {

        
            CountryCode float64  `json:"country_code"`
            Number string  `json:"number"`
         
    }
    
    // ManagerResponse used by Logistics
    type ManagerResponse struct {

        
            Name string  `json:"name"`
            MobileNo MobileNo  `json:"mobile_no"`
            Email string  `json:"email"`
         
    }
    
    // IntegrationTypeResponse used by Logistics
    type IntegrationTypeResponse struct {

        
            Order string  `json:"order"`
            Inventory string  `json:"inventory"`
         
    }
    
    // ModifiedByResponse used by Logistics
    type ModifiedByResponse struct {

        
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
         
    }
    
    // OpeningClosing used by Logistics
    type OpeningClosing struct {

        
            Minute float64  `json:"minute"`
            Hour float64  `json:"hour"`
         
    }
    
    // TimmingResponse used by Logistics
    type TimmingResponse struct {

        
            Open bool  `json:"open"`
            Opening OpeningClosing  `json:"opening"`
            Weekday string  `json:"weekday"`
            Closing OpeningClosing  `json:"closing"`
         
    }
    
    // DocumentsResponse used by Logistics
    type DocumentsResponse struct {

        
            Verified bool  `json:"verified"`
            Value string  `json:"value"`
            LegalName string  `json:"legal_name"`
            Type string  `json:"type"`
         
    }
    
    // AddressResponse used by Logistics
    type AddressResponse struct {

        
            City string  `json:"city"`
            State string  `json:"state"`
            Longitude float64  `json:"longitude"`
            Country string  `json:"country"`
            Latitude float64  `json:"latitude"`
            Landmark string  `json:"landmark"`
            Pincode float64  `json:"pincode"`
            Address2 string  `json:"address2"`
            Address1 string  `json:"address1"`
         
    }
    
    // ItemResponse used by Logistics
    type ItemResponse struct {

        
            DisplayName string  `json:"display_name"`
            Warnings WarningsResponse  `json:"warnings"`
            SubType string  `json:"sub_type"`
            Company float64  `json:"company"`
            Logistics LogisticsResponse  `json:"logistics"`
            NotificationEmails []string  `json:"notification_emails"`
            ContactNumbers []ContactNumberResponse  `json:"contact_numbers"`
            Cls string  `json:"_cls"`
            ProductReturnConfig ProductReturnConfigResponse  `json:"product_return_config"`
            CreatedBy CreatedByResponse  `json:"created_by"`
            VerifiedOn string  `json:"verified_on"`
            Name string  `json:"name"`
            GstCredentials GstCredentialsResponse  `json:"gst_credentials"`
            Manager ManagerResponse  `json:"manager"`
            ModifiedOn string  `json:"modified_on"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Stage string  `json:"stage"`
            IntegrationType IntegrationTypeResponse  `json:"integration_type"`
            CompanyID float64  `json:"company_id"`
            VerifiedBy ModifiedByResponse  `json:"verified_by"`
            Timing []TimmingResponse  `json:"timing"`
            UID float64  `json:"uid"`
            StoreType string  `json:"store_type"`
            Documents []DocumentsResponse  `json:"documents"`
            Code string  `json:"code"`
            ModifiedBy ModifiedByResponse  `json:"modified_by"`
            Address AddressResponse  `json:"address"`
            CreatedOn string  `json:"created_on"`
         
    }
    
    // GetStoresViewResponse used by Logistics
    type GetStoresViewResponse struct {

        
            Page PageResponse  `json:"page"`
            Items []ItemResponse  `json:"items"`
         
    }
    
    // PincodeMopData used by Logistics
    type PincodeMopData struct {

        
            Pincodes []float64  `json:"pincodes"`
            Country string  `json:"country"`
            Action string  `json:"action"`
         
    }
    
    // PincodeMopUpdateResponse used by Logistics
    type PincodeMopUpdateResponse struct {

        
            Pincode float64  `json:"pincode"`
            ChannelID string  `json:"channel_id"`
            Country string  `json:"country"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // PincodeMOPresponse used by Logistics
    type PincodeMOPresponse struct {

        
            Success bool  `json:"success"`
            StatusCode float64  `json:"status_code"`
            BatchID string  `json:"batch_id"`
            Country string  `json:"country"`
            Action string  `json:"action"`
            Pincodes []float64  `json:"pincodes"`
            UpdatedPincodes []PincodeMopUpdateResponse  `json:"updated_pincodes"`
         
    }
    
    // CommonError used by Logistics
    type CommonError struct {

        
            StatusCode string  `json:"status_code"`
            Success string  `json:"success"`
            Error interface{}  `json:"error"`
         
    }
    
    // PincodeMopBulkData used by Logistics
    type PincodeMopBulkData struct {

        
            BatchID string  `json:"batch_id"`
            S3URL string  `json:"s3_url"`
         
    }
    
    // PincodeBulkViewResponse used by Logistics
    type PincodeBulkViewResponse struct {

        
            BatchID string  `json:"batch_id"`
            S3URL string  `json:"s3_url"`
         
    }
    
    // PincodeCodStatusListingRequest used by Logistics
    type PincodeCodStatusListingRequest struct {

        
            Country string  `json:"country"`
            IsActive bool  `json:"is_active"`
            Pincode float64  `json:"pincode"`
            Current float64  `json:"current"`
            PageSize float64  `json:"page_size"`
         
    }
    
    // PincodeCodStatusListingResponse used by Logistics
    type PincodeCodStatusListingResponse struct {

        
            Country string  `json:"country"`
            Data []PincodeCodStatusListingResponse  `json:"data"`
            Success bool  `json:"success"`
            Errors []Error  `json:"errors"`
            Page PincodeCodStatusListingPage  `json:"page"`
            Summary PincodeCodStatusListingSummary  `json:"summary"`
         
    }
    
    // Error used by Logistics
    type Error struct {

        
            Type string  `json:"type"`
            Value string  `json:"value"`
            Message string  `json:"message"`
         
    }
    
    // PincodeCodStatusListingPage used by Logistics
    type PincodeCodStatusListingPage struct {

        
            ItemTotal float64  `json:"item_total"`
            HasNext bool  `json:"has_next"`
            Size float64  `json:"size"`
            Type string  `json:"type"`
            Current float64  `json:"current"`
         
    }
    
    // PincodeCodStatusListingSummary used by Logistics
    type PincodeCodStatusListingSummary struct {

        
            TotalActivePincodes float64  `json:"total_active_pincodes"`
            TotalInactivePincodes float64  `json:"total_inactive_pincodes"`
         
    }
    
    // PincodeMopUpdateAuditHistoryRequest used by Logistics
    type PincodeMopUpdateAuditHistoryRequest struct {

        
            EntityType string  `json:"entity_type"`
            FileName string  `json:"file_name"`
         
    }
    
    // PincodeMopUpdateAuditHistoryPaging used by Logistics
    type PincodeMopUpdateAuditHistoryPaging struct {

        
            Type string  `json:"type"`
            Size float64  `json:"size"`
            Current float64  `json:"current"`
            HasNext bool  `json:"has_next"`
            ItemTotal float64  `json:"item_total"`
         
    }
    
    // PincodeMopUpdateAuditHistoryResponse used by Logistics
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
    
    // PincodeMopUpdateAuditHistoryResponseData used by Logistics
    type PincodeMopUpdateAuditHistoryResponseData struct {

        
            EntityType string  `json:"entity_type"`
            Page PincodeMopUpdateAuditHistoryPaging  `json:"page"`
            Data []PincodeMopUpdateAuditHistoryResponse  `json:"data"`
         
    }
    
