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

        
            Created bool  `json:"created"`
            AppID string  `json:"app_id"`
            DisplayFields []string  `json:"display_fields"`
            ExcludedFields []string  `json:"excluded_fields"`
            Aggregators []map[string]interface{}  `json:"aggregators"`
            Success bool  `json:"success"`
         
    }
    
    // ErrorCodeDescription used by Payment
    type ErrorCodeDescription struct {

        
            Description string  `json:"description"`
            Success bool  `json:"success"`
            Code string  `json:"code"`
         
    }
    
    // PaymentGatewayConfig used by Payment
    type PaymentGatewayConfig struct {

        
            Secret string  `json:"secret"`
            ConfigType string  `json:"config_type"`
            IsActive bool  `json:"is_active"`
            MerchantSalt string  `json:"merchant_salt"`
            Key string  `json:"key"`
         
    }
    
    // PaymentGatewayConfigRequest used by Payment
    type PaymentGatewayConfigRequest struct {

        
            AggregatorName PaymentGatewayConfig  `json:"aggregator_name"`
            IsActive bool  `json:"is_active"`
            AppID string  `json:"app_id"`
         
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

        
            Small string  `json:"small"`
            Large string  `json:"large"`
         
    }
    
    // IntentApp used by Payment
    type IntentApp struct {

        
            PackageName string  `json:"package_name"`
            Logos PaymentModeLogo  `json:"logos"`
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

        
            Expired bool  `json:"expired"`
            Timeout float64  `json:"timeout"`
            MerchantCode string  `json:"merchant_code"`
            CardBrand string  `json:"card_brand"`
            IntentApp []IntentApp  `json:"intent_app"`
            IntentAppErrorList []string  `json:"intent_app_error_list"`
            DisplayPriority float64  `json:"display_priority"`
            ExpYear float64  `json:"exp_year"`
            CodLimit float64  `json:"cod_limit"`
            AggregatorName string  `json:"aggregator_name"`
            CardType string  `json:"card_type"`
            DisplayName string  `json:"display_name"`
            CardID string  `json:"card_id"`
            CardName string  `json:"card_name"`
            Name string  `json:"name"`
            CardReference string  `json:"card_reference"`
            CompliantWithTokenisationGuidelines bool  `json:"compliant_with_tokenisation_guidelines"`
            CardIssuer string  `json:"card_issuer"`
            CardNumber string  `json:"card_number"`
            IntentAppErrorDictList []IntentAppErrorList  `json:"intent_app_error_dict_list"`
            CardIsin string  `json:"card_isin"`
            Nickname string  `json:"nickname"`
            CodLimitPerOrder float64  `json:"cod_limit_per_order"`
            FyndVpa string  `json:"fynd_vpa"`
            RemainingLimit float64  `json:"remaining_limit"`
            LogoURL PaymentModeLogo  `json:"logo_url"`
            CardToken string  `json:"card_token"`
            IntentFlow bool  `json:"intent_flow"`
            RetryCount float64  `json:"retry_count"`
            Code string  `json:"code"`
            CardFingerprint string  `json:"card_fingerprint"`
            CardBrandImage string  `json:"card_brand_image"`
            ExpMonth float64  `json:"exp_month"`
         
    }
    
    // RootPaymentMode used by Payment
    type RootPaymentMode struct {

        
            AddCardEnabled bool  `json:"add_card_enabled"`
            AnonymousEnable bool  `json:"anonymous_enable"`
            Name string  `json:"name"`
            IsPayByCardPl bool  `json:"is_pay_by_card_pl"`
            DisplayPriority float64  `json:"display_priority"`
            SaveCard bool  `json:"save_card"`
            List []PaymentModeList  `json:"list"`
            AggregatorName string  `json:"aggregator_name"`
            DisplayName string  `json:"display_name"`
         
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
            State string  `json:"state"`
            AccountType string  `json:"account_type"`
            City string  `json:"city"`
            BankName string  `json:"bank_name"`
            IfscCode string  `json:"ifsc_code"`
            BranchName string  `json:"branch_name"`
            AccountNo string  `json:"account_no"`
            AccountHolder string  `json:"account_holder"`
         
    }
    
    // PayoutCustomer used by Payment
    type PayoutCustomer struct {

        
            UniqueExternalID string  `json:"unique_external_id"`
            ID float64  `json:"id"`
            Name string  `json:"name"`
            Mobile string  `json:"mobile"`
         
    }
    
    // PayoutAggregator used by Payment
    type PayoutAggregator struct {

        
            AggregatorID float64  `json:"aggregator_id"`
            AggregatorFundID float64  `json:"aggregator_fund_id"`
            PayoutDetailsID float64  `json:"payout_details_id"`
         
    }
    
    // Payout used by Payment
    type Payout struct {

        
            TransferType string  `json:"transfer_type"`
            IsActive bool  `json:"is_active"`
            MoreAttributes PayoutMoreAttributes  `json:"more_attributes"`
            Customers PayoutCustomer  `json:"customers"`
            PayoutsAggregators []PayoutAggregator  `json:"payouts_aggregators"`
            IsDefault bool  `json:"is_default"`
            UniqueTransferNo string  `json:"unique_transfer_no"`
         
    }
    
    // PayoutsResponse used by Payment
    type PayoutsResponse struct {

        
            Items []Payout  `json:"items"`
            Success bool  `json:"success"`
         
    }
    
    // PayoutBankDetails used by Payment
    type PayoutBankDetails struct {

        
            Country string  `json:"country"`
            State string  `json:"state"`
            AccountType string  `json:"account_type"`
            City string  `json:"city"`
            BankName string  `json:"bank_name"`
            IfscCode string  `json:"ifsc_code"`
            Pincode float64  `json:"pincode"`
            BranchName string  `json:"branch_name"`
            AccountNo string  `json:"account_no"`
            AccountHolder string  `json:"account_holder"`
         
    }
    
    // PayoutRequest used by Payment
    type PayoutRequest struct {

        
            BankDetails PayoutBankDetails  `json:"bank_details"`
            TransferType string  `json:"transfer_type"`
            IsActive bool  `json:"is_active"`
            UniqueExternalID string  `json:"unique_external_id"`
            Aggregator string  `json:"aggregator"`
            Users map[string]interface{}  `json:"users"`
         
    }
    
    // PayoutResponse used by Payment
    type PayoutResponse struct {

        
            Created bool  `json:"created"`
            Success bool  `json:"success"`
            BankDetails map[string]interface{}  `json:"bank_details"`
            TransferType string  `json:"transfer_type"`
            IsActive bool  `json:"is_active"`
            Payouts map[string]interface{}  `json:"payouts"`
            Aggregator string  `json:"aggregator"`
            Users map[string]interface{}  `json:"users"`
            PaymentStatus string  `json:"payment_status"`
            UniqueTransferNo string  `json:"unique_transfer_no"`
         
    }
    
    // UpdatePayoutResponse used by Payment
    type UpdatePayoutResponse struct {

        
            IsDefault bool  `json:"is_default"`
            IsActive bool  `json:"is_active"`
            Success bool  `json:"success"`
         
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

        
            Config map[string]interface{}  `json:"config"`
            Aggregator string  `json:"aggregator"`
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
            BranchName string  `json:"branch_name"`
            AccountNo string  `json:"account_no"`
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
            Title string  `json:"title"`
            Mobile string  `json:"mobile"`
            TransferMode string  `json:"transfer_mode"`
            DelightsUserName string  `json:"delights_user_name"`
            Address string  `json:"address"`
            ModifiedOn string  `json:"modified_on"`
            Comment string  `json:"comment"`
            Email string  `json:"email"`
            BankName string  `json:"bank_name"`
            BranchName string  `json:"branch_name"`
            Subtitle string  `json:"subtitle"`
            AccountNo string  `json:"account_no"`
            DisplayName string  `json:"display_name"`
            AccountHolder string  `json:"account_holder"`
            BeneficiaryID string  `json:"beneficiary_id"`
            IsActive bool  `json:"is_active"`
            IfscCode string  `json:"ifsc_code"`
            ID float64  `json:"id"`
         
    }
    
    // OrderBeneficiaryResponse used by Payment
    type OrderBeneficiaryResponse struct {

        
            ShowBeneficiaryDetails bool  `json:"show_beneficiary_details"`
            Beneficiaries []OrderBeneficiaryDetails  `json:"beneficiaries"`
         
    }
    
    // MultiTenderPaymentMeta used by Payment
    type MultiTenderPaymentMeta struct {

        
            OrderID string  `json:"order_id"`
            CurrentStatus string  `json:"current_status"`
            PaymentID string  `json:"payment_id"`
            PaymentGateway string  `json:"payment_gateway"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
         
    }
    
    // MultiTenderPaymentMethod used by Payment
    type MultiTenderPaymentMethod struct {

        
            Amount float64  `json:"amount"`
            Mode string  `json:"mode"`
            Meta MultiTenderPaymentMeta  `json:"meta"`
            Name string  `json:"name"`
         
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

        
            Methods map[string]interface{}  `json:"methods"`
            CallbackURL map[string]interface{}  `json:"callback_url"`
            CodAmountLimit float64  `json:"cod_amount_limit"`
            ModeOfPayment string  `json:"mode_of_payment"`
            AnonymousCod bool  `json:"anonymous_cod"`
            Enabled bool  `json:"enabled"`
            Source string  `json:"source"`
            PaymentSelectionLock map[string]interface{}  `json:"payment_selection_lock"`
            CodCharges float64  `json:"cod_charges"`
         
    }
    
    // PlatfromPaymentConfig used by Payment
    type PlatfromPaymentConfig struct {

        
            Message string  `json:"message"`
            Data PlatformPaymentOptions  `json:"data"`
            Success bool  `json:"success"`
         
    }
    
    // UpdatePlatformPaymentConfig used by Payment
    type UpdatePlatformPaymentConfig struct {

        
            Methods map[string]interface{}  `json:"methods"`
            CodAmountLimit float64  `json:"cod_amount_limit"`
            AnonymousCod bool  `json:"anonymous_cod"`
            CodCharges float64  `json:"cod_charges"`
            PaymentSelectionLock map[string]interface{}  `json:"payment_selection_lock"`
         
    }
    
    // CODdata used by Payment
    type CODdata struct {

        
            Usages float64  `json:"usages"`
            IsActive bool  `json:"is_active"`
            UserID string  `json:"user_id"`
            Limit float64  `json:"limit"`
            RemainingLimit float64  `json:"remaining_limit"`
         
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

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // EdcModelData used by Payment
    type EdcModelData struct {

        
            AggregatorID float64  `json:"aggregator_id"`
            Aggregator string  `json:"aggregator"`
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

        
            Statistics StatisticsData  `json:"statistics"`
            Success bool  `json:"success"`
         
    }
    
    // EdcAddRequest used by Payment
    type EdcAddRequest struct {

        
            EdcModel string  `json:"edc_model"`
            StoreID float64  `json:"store_id"`
            DeviceTag string  `json:"device_tag"`
            TerminalSerialNo string  `json:"terminal_serial_no"`
            AggregatorID float64  `json:"aggregator_id"`
            EdcDeviceSerialNo string  `json:"edc_device_serial_no"`
         
    }
    
    // EdcDevice used by Payment
    type EdcDevice struct {

        
            TerminalUniqueIdentifier string  `json:"terminal_unique_identifier"`
            MerchantStorePosCode string  `json:"merchant_store_pos_code"`
            EdcModel string  `json:"edc_model"`
            ApplicationID string  `json:"application_id"`
            StoreID float64  `json:"store_id"`
            DeviceTag string  `json:"device_tag"`
            IsActive bool  `json:"is_active"`
            TerminalSerialNo string  `json:"terminal_serial_no"`
            AggregatorID float64  `json:"aggregator_id"`
            EdcDeviceSerialNo string  `json:"edc_device_serial_no"`
            AggregatorName string  `json:"aggregator_name"`
         
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

        
            MerchantStorePosCode string  `json:"merchant_store_pos_code"`
            EdcModel string  `json:"edc_model"`
            StoreID float64  `json:"store_id"`
            IsActive bool  `json:"is_active"`
            DeviceTag string  `json:"device_tag"`
            AggregatorID float64  `json:"aggregator_id"`
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

        
            Items []EdcDevice  `json:"items"`
            Page Page  `json:"page"`
            Success bool  `json:"success"`
         
    }
    
    // PaymentInitializationRequest used by Payment
    type PaymentInitializationRequest struct {

        
            Timeout float64  `json:"timeout"`
            CustomerID string  `json:"customer_id"`
            Contact string  `json:"contact"`
            OrderID string  `json:"order_id"`
            Currency string  `json:"currency"`
            Email string  `json:"email"`
            DeviceID string  `json:"device_id"`
            Amount float64  `json:"amount"`
            MerchantOrderID string  `json:"merchant_order_id"`
            RazorpayPaymentID string  `json:"razorpay_payment_id"`
            Aggregator string  `json:"aggregator"`
            Vpa string  `json:"vpa"`
            Method string  `json:"method"`
         
    }
    
    // PaymentInitializationResponse used by Payment
    type PaymentInitializationResponse struct {

        
            Timeout float64  `json:"timeout"`
            Contact string  `json:"contact"`
            VirtualID string  `json:"virtual_id"`
            Aggregator string  `json:"aggregator"`
            Status string  `json:"status"`
            PaymentID string  `json:"payment_id"`
            UpiPollURL string  `json:"upi_poll_url"`
            DeviceID string  `json:"device_id"`
            Success bool  `json:"success"`
            CustomerID string  `json:"customer_id"`
            Email string  `json:"email"`
            Amount float64  `json:"amount"`
            MerchantOrderID string  `json:"merchant_order_id"`
            BqrImage string  `json:"bqr_image"`
            Currency string  `json:"currency"`
            RazorpayPaymentID string  `json:"razorpay_payment_id"`
            PollingURL string  `json:"polling_url"`
            Method string  `json:"method"`
            OrderID string  `json:"order_id"`
            AggregatorOrderID string  `json:"aggregator_order_id"`
            Vpa string  `json:"vpa"`
         
    }
    
    // PaymentStatusUpdateRequest used by Payment
    type PaymentStatusUpdateRequest struct {

        
            MerchantTransactionID string  `json:"merchant_transaction_id"`
            CustomerID string  `json:"customer_id"`
            Contact string  `json:"contact"`
            OrderID string  `json:"order_id"`
            Currency string  `json:"currency"`
            Email string  `json:"email"`
            DeviceID string  `json:"device_id"`
            Amount float64  `json:"amount"`
            MerchantOrderID string  `json:"merchant_order_id"`
            RazorpayPaymentID string  `json:"razorpay_payment_id"`
            Aggregator string  `json:"aggregator"`
            Status string  `json:"status"`
            PaymentID string  `json:"payment_id"`
            Vpa string  `json:"vpa"`
            Method string  `json:"method"`
         
    }
    
    // PaymentStatusUpdateResponse used by Payment
    type PaymentStatusUpdateResponse struct {

        
            Retry bool  `json:"retry"`
            RedirectURL string  `json:"redirect_url"`
            Status string  `json:"status"`
            AggregatorName string  `json:"aggregator_name"`
            Success bool  `json:"success"`
         
    }
    
    // ResendOrCancelPaymentRequest used by Payment
    type ResendOrCancelPaymentRequest struct {

        
            RequestType string  `json:"request_type"`
            DeviceID string  `json:"device_id"`
            OrderID string  `json:"order_id"`
         
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

        
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
            AggregatorPaymentObject map[string]interface{}  `json:"aggregator_payment_object"`
            RefundObject map[string]interface{}  `json:"refund_object"`
            Currency string  `json:"currency"`
            PaymentModeIdentifier string  `json:"payment_mode_identifier"`
            PaymentGateway string  `json:"payment_gateway"`
            ApplicationID string  `json:"application_id"`
            AllStatus []string  `json:"all_status"`
            RefundedBy string  `json:"refunded_by"`
            AmountInPaisa string  `json:"amount_in_paisa"`
            CollectedBy string  `json:"collected_by"`
            UserObject map[string]interface{}  `json:"user_object"`
            CurrentStatus string  `json:"current_status"`
            CompanyID string  `json:"company_id"`
            ID string  `json:"id"`
            PaymentID string  `json:"payment_id"`
            PaymentMode string  `json:"payment_mode"`
         
    }
    
    // PaymentStatusObject used by Payment
    type PaymentStatusObject struct {

        
            MerchantOrderID string  `json:"merchant_order_id"`
            PaymentObjectList []PaymentObjectListSerializer  `json:"payment_object_list"`
         
    }
    
    // PaymentStatusBulkHandlerResponse used by Payment
    type PaymentStatusBulkHandlerResponse struct {

        
            Count float64  `json:"count"`
            Error string  `json:"error"`
            Data []PaymentStatusObject  `json:"data"`
            Status float64  `json:"status"`
            Success string  `json:"success"`
         
    }
    
    // RevokeOAuthToken used by Payment
    type RevokeOAuthToken struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    

    
    // Prices used by Order
    type Prices struct {

        
            CashbackApplied float64  `json:"cashback_applied"`
            PmPriceSplit float64  `json:"pm_price_split"`
            PriceEffective float64  `json:"price_effective"`
            Cashback float64  `json:"cashback"`
            Discount float64  `json:"discount"`
            RefundAmount float64  `json:"refund_amount"`
            PromotionEffectiveDiscount float64  `json:"promotion_effective_discount"`
            FyndCredits float64  `json:"fynd_credits"`
            TransferPrice float64  `json:"transfer_price"`
            AmountPaid float64  `json:"amount_paid"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            CodCharges float64  `json:"cod_charges"`
            DeliveryCharge float64  `json:"delivery_charge"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            AmountPaidRoundoff float64  `json:"amount_paid_roundoff"`
            PriceMarked float64  `json:"price_marked"`
            RefundCredit float64  `json:"refund_credit"`
            CouponValue float64  `json:"coupon_value"`
            ValueOfGood float64  `json:"value_of_good"`
         
    }
    
    // BagReturnableCancelableStatus used by Order
    type BagReturnableCancelableStatus struct {

        
            IsReturnable bool  `json:"is_returnable"`
            CanBeCancelled bool  `json:"can_be_cancelled"`
            IsActive bool  `json:"is_active"`
            IsCustomerReturnAllowed bool  `json:"is_customer_return_allowed"`
            EnableTracking bool  `json:"enable_tracking"`
         
    }
    
    // Identifier used by Order
    type Identifier struct {

        
            Ean string  `json:"ean"`
         
    }
    
    // FinancialBreakup used by Order
    type FinancialBreakup struct {

        
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            HsnCode string  `json:"hsn_code"`
            CashbackApplied float64  `json:"cashback_applied"`
            Identifiers Identifier  `json:"identifiers"`
            PriceEffective float64  `json:"price_effective"`
            Cashback float64  `json:"cashback"`
            GstTag string  `json:"gst_tag"`
            Discount float64  `json:"discount"`
            PromotionEffectiveDiscount float64  `json:"promotion_effective_discount"`
            TotalUnits float64  `json:"total_units"`
            FyndCredits float64  `json:"fynd_credits"`
            TransferPrice float64  `json:"transfer_price"`
            AmountPaid float64  `json:"amount_paid"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            AddedToFyndCash bool  `json:"added_to_fynd_cash"`
            CodCharges float64  `json:"cod_charges"`
            DeliveryCharge float64  `json:"delivery_charge"`
            ItemName string  `json:"item_name"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            GstFee float64  `json:"gst_fee"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            AmountPaidRoundoff float64  `json:"amount_paid_roundoff"`
            PriceMarked float64  `json:"price_marked"`
            RefundCredit float64  `json:"refund_credit"`
            Size string  `json:"size"`
            CouponValue float64  `json:"coupon_value"`
            ValueOfGood float64  `json:"value_of_good"`
         
    }
    
    // PlatformArticleAttributes used by Order
    type PlatformArticleAttributes struct {

        
            Currency string  `json:"currency"`
         
    }
    
    // PlatformItem used by Order
    type PlatformItem struct {

        
            Image []string  `json:"image"`
            CanCancel bool  `json:"can_cancel"`
            L1Category []string  `json:"l1_category"`
            BrandID float64  `json:"brand_id"`
            L3Category float64  `json:"l3_category"`
            Name string  `json:"name"`
            CanReturn bool  `json:"can_return"`
            Brand string  `json:"brand"`
            L2Category []string  `json:"l2_category"`
            BranchURL string  `json:"branch_url"`
            Color string  `json:"color"`
            SlugKey string  `json:"slug_key"`
            Code string  `json:"code"`
            LastUpdatedAt string  `json:"last_updated_at"`
            DepartmentID float64  `json:"department_id"`
            Attributes PlatformArticleAttributes  `json:"attributes"`
            ID float64  `json:"id"`
            Size string  `json:"size"`
            L3CategoryName string  `json:"l3_category_name"`
            Images []string  `json:"images"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // BagStateMapper used by Order
    type BagStateMapper struct {

        
            AppStateName string  `json:"app_state_name"`
            NotifyCustomer bool  `json:"notify_customer"`
            IsActive bool  `json:"is_active"`
            JourneyType string  `json:"journey_type"`
            StateType string  `json:"state_type"`
            AppFacing bool  `json:"app_facing"`
            Name string  `json:"name"`
            AppDisplayName string  `json:"app_display_name"`
            BsID float64  `json:"bs_id"`
            DisplayName string  `json:"display_name"`
         
    }
    
    // BagStatusHistory used by Order
    type BagStatusHistory struct {

        
            Status string  `json:"status"`
            AppDisplayName string  `json:"app_display_name"`
            Reasons []map[string]interface{}  `json:"reasons"`
            BagStateMapper BagStateMapper  `json:"bag_state_mapper"`
            StoreID float64  `json:"store_id"`
            BshID float64  `json:"bsh_id"`
            Forward bool  `json:"forward"`
            DeliveryPartnerID float64  `json:"delivery_partner_id"`
            KafkaSync bool  `json:"kafka_sync"`
            StateType string  `json:"state_type"`
            UpdatedAt string  `json:"updated_at"`
            ShipmentID string  `json:"shipment_id"`
            BagID float64  `json:"bag_id"`
            CreatedAt string  `json:"created_at"`
            DeliveryAwbNumber string  `json:"delivery_awb_number"`
            DisplayName string  `json:"display_name"`
            StateID float64  `json:"state_id"`
         
    }
    
    // ShipmentListingBrand used by Order
    type ShipmentListingBrand struct {

        
            Name string  `json:"name"`
            LogoBase64 string  `json:"logo_base64"`
            CreatedOn string  `json:"created_on"`
            Logo string  `json:"logo"`
         
    }
    
    // AffiliateMeta used by Order
    type AffiliateMeta struct {

        
            SizeLevelTotalQty float64  `json:"size_level_total_qty"`
            CouponCode string  `json:"coupon_code"`
            BoxType string  `json:"box_type"`
            EmployeeDiscount float64  `json:"employee_discount"`
            Quantity float64  `json:"quantity"`
            IsPriority bool  `json:"is_priority"`
            DueDate string  `json:"due_date"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            OrderItemID string  `json:"order_item_id"`
            ChannelShipmentID string  `json:"channel_shipment_id"`
            ChannelOrderID string  `json:"channel_order_id"`
            MarketplaceInvoiceID string  `json:"marketplace_invoice_id"`
         
    }
    
    // AffiliateBagDetails used by Order
    type AffiliateBagDetails struct {

        
            AffiliateOrderID string  `json:"affiliate_order_id"`
            EmployeeDiscount float64  `json:"employee_discount"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
            AffiliateMeta AffiliateMeta  `json:"affiliate_meta"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
         
    }
    
    // GSTDetailsData used by Order
    type GSTDetailsData struct {

        
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            SgstGstFee string  `json:"sgst_gst_fee"`
            CgstGstFee string  `json:"cgst_gst_fee"`
            HsnCode string  `json:"hsn_code"`
            HsnCodeID string  `json:"hsn_code_id"`
            GstTag string  `json:"gst_tag"`
            GstinCode string  `json:"gstin_code"`
            IgstGstFee string  `json:"igst_gst_fee"`
            SgstTaxPercentage float64  `json:"sgst_tax_percentage"`
            CgstTaxPercentage float64  `json:"cgst_tax_percentage"`
            GstFee float64  `json:"gst_fee"`
            IgstTaxPercentage float64  `json:"igst_tax_percentage"`
            IsDefaultHsnCode bool  `json:"is_default_hsn_code"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            ValueOfGood float64  `json:"value_of_good"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
         
    }
    
    // Dates used by Order
    type Dates struct {

        
            DeliveryDate string  `json:"delivery_date"`
            OrderCreated string  `json:"order_created"`
         
    }
    
    // Dimensions used by Order
    type Dimensions struct {

        
            Unit string  `json:"unit"`
            Height float64  `json:"height"`
            Length float64  `json:"length"`
            Width float64  `json:"width"`
            IsDefault bool  `json:"is_default"`
         
    }
    
    // Weight used by Order
    type Weight struct {

        
            Shipping float64  `json:"shipping"`
            Unit string  `json:"unit"`
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

        
            Code string  `json:"code"`
            Currency map[string]interface{}  `json:"currency"`
            Dimensions Dimensions  `json:"dimensions"`
            Weight Weight  `json:"weight"`
            RawMeta string  `json:"raw_meta"`
            Size string  `json:"size"`
            ReturnConfig ReturnConfig  `json:"return_config"`
            ID string  `json:"_id"`
            UID string  `json:"uid"`
            IsSet bool  `json:"is_set"`
            ASet map[string]interface{}  `json:"a_set"`
            SellerIdentifier string  `json:"seller_identifier"`
            EspModified bool  `json:"esp_modified"`
            ChildDetails map[string]interface{}  `json:"child_details"`
            Identifiers Identifier  `json:"identifiers"`
         
    }
    
    // BagUnit used by Order
    type BagUnit struct {

        
            Status BagReturnableCancelableStatus  `json:"status"`
            CanCancel bool  `json:"can_cancel"`
            EntityType string  `json:"entity_type"`
            Prices Prices  `json:"prices"`
            CanReturn bool  `json:"can_return"`
            BagType string  `json:"bag_type"`
            FinancialBreakup []FinancialBreakup  `json:"financial_breakup"`
            Item PlatformItem  `json:"item"`
            BagStatus []BagStatusHistory  `json:"bag_status"`
            Brand ShipmentListingBrand  `json:"brand"`
            BagExpiryDate string  `json:"bag_expiry_date"`
            ProductQuantity float64  `json:"product_quantity"`
            AffiliateBagDetails AffiliateBagDetails  `json:"affiliate_bag_details"`
            DisplayName string  `json:"display_name"`
            CurrentStatus BagStatusHistory  `json:"current_status"`
            Gst GSTDetailsData  `json:"gst"`
            BagID float64  `json:"bag_id"`
            Reasons []map[string]interface{}  `json:"reasons"`
            Size string  `json:"size"`
            Dates Dates  `json:"dates"`
            CurrentOperationalStatus BagStatusHistory  `json:"current_operational_status"`
            Article Article  `json:"article"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // UserDataInfo used by Order
    type UserDataInfo struct {

        
            Gender string  `json:"gender"`
            IsAnonymousUser bool  `json:"is_anonymous_user"`
            AvisUserID string  `json:"avis_user_id"`
            LastName string  `json:"last_name"`
            UID float64  `json:"uid"`
            FirstName string  `json:"first_name"`
            Mobile string  `json:"mobile"`
            Name string  `json:"name"`
            Email string  `json:"email"`
         
    }
    
    // ShipmentListingChannel used by Order
    type ShipmentListingChannel struct {

        
            Name string  `json:"name"`
            ChannelShipmentID string  `json:"channel_shipment_id"`
            IsAffiliate bool  `json:"is_affiliate"`
            Logo string  `json:"logo"`
         
    }
    
    // PlatformDeliveryAddress used by Order
    type PlatformDeliveryAddress struct {

        
            Pincode string  `json:"pincode"`
            AddressType string  `json:"address_type"`
            City string  `json:"city"`
            Area string  `json:"area"`
            State string  `json:"state"`
            Version string  `json:"version"`
            AddressCategory string  `json:"address_category"`
            Longitude float64  `json:"longitude"`
            Phone string  `json:"phone"`
            Latitude float64  `json:"latitude"`
            UpdatedAt string  `json:"updated_at"`
            ContactPerson string  `json:"contact_person"`
            Address1 string  `json:"address1"`
            Landmark string  `json:"landmark"`
            Country string  `json:"country"`
            Address2 string  `json:"address2"`
            Email string  `json:"email"`
            CreatedAt string  `json:"created_at"`
         
    }
    
    // ShipmentItemFulFillingStore used by Order
    type ShipmentItemFulFillingStore struct {

        
            Code string  `json:"code"`
            Pincode string  `json:"pincode"`
            City string  `json:"city"`
            State string  `json:"state"`
            FulfillingStoreID float64  `json:"fulfilling_store_id"`
            Phone string  `json:"phone"`
            Name string  `json:"name"`
            StoreEmail string  `json:"store_email"`
            Address string  `json:"address"`
            LocationType string  `json:"location_type"`
            BrandStoreTags string  `json:"brand_store_tags"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // ShipmentStatus used by Order
    type ShipmentStatus struct {

        
            CurrentShipmentStatus string  `json:"current_shipment_status"`
            Status string  `json:"status"`
            Title string  `json:"title"`
            ShipmentStatusID float64  `json:"shipment_status_id"`
            ShipmentID string  `json:"shipment_id"`
            CreatedAt string  `json:"created_at"`
            BagList []string  `json:"bag_list"`
            StatusCreatedAt string  `json:"status_created_at"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // ShipmentItem used by Order
    type ShipmentItem struct {

        
            PreviousShipmentID string  `json:"previous_shipment_id"`
            OrderID string  `json:"order_id"`
            TotalBags float64  `json:"total_bags"`
            Prices Prices  `json:"prices"`
            Bags []BagUnit  `json:"bags"`
            ShipmentID string  `json:"shipment_id"`
            User UserDataInfo  `json:"user"`
            LockStatus bool  `json:"lock_status"`
            CustomerNote string  `json:"customer_note"`
            Channel ShipmentListingChannel  `json:"channel"`
            DeliveryAddress PlatformDeliveryAddress  `json:"delivery_address"`
            FulfillingStore ShipmentItemFulFillingStore  `json:"fulfilling_store"`
            PaymentMethods map[string]interface{}  `json:"payment_methods"`
            CanProcess bool  `json:"can_process"`
            DisplayName string  `json:"display_name"`
            StatusCreatedAt string  `json:"status_created_at"`
            ShipmentStatus ShipmentStatus  `json:"shipment_status"`
            PaymentMode string  `json:"payment_mode"`
            OrderDate string  `json:"order_date"`
            EstimatedSlaTime string  `json:"estimated_sla_time"`
            InvoiceID string  `json:"invoice_id"`
            ShipmentCreatedAt string  `json:"shipment_created_at"`
            OrderingChannnel string  `json:"ordering_channnel"`
            Meta map[string]interface{}  `json:"meta"`
         
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

        
            Items []ShipmentItem  `json:"items"`
            Lane string  `json:"lane"`
            TotalCount float64  `json:"total_count"`
            Message string  `json:"message"`
            Page Page  `json:"page"`
            Success bool  `json:"success"`
         
    }
    
    // Error used by Order
    type Error struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // ShipmentStatusData used by Order
    type ShipmentStatusData struct {

        
            Status string  `json:"status"`
            ID float64  `json:"id"`
            ShipmentID string  `json:"shipment_id"`
            BagList []string  `json:"bag_list"`
            CreatedAt string  `json:"created_at"`
         
    }
    
    // UserDetailsData used by Order
    type UserDetailsData struct {

        
            Pincode string  `json:"pincode"`
            AddressType string  `json:"address_type"`
            City string  `json:"city"`
            Area string  `json:"area"`
            State string  `json:"state"`
            Phone string  `json:"phone"`
            Name string  `json:"name"`
            Address string  `json:"address"`
            Address1 string  `json:"address1"`
            Landmark string  `json:"landmark"`
            Country string  `json:"country"`
            Email string  `json:"email"`
         
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

        
            ItemCriteria ItemCriterias  `json:"item_criteria"`
            CartConditions map[string]interface{}  `json:"cart_conditions"`
         
    }
    
    // AppliedPromos used by Order
    type AppliedPromos struct {

        
            ArticleQuantity float64  `json:"article_quantity"`
            PromotionType string  `json:"promotion_type"`
            MrpPromotion bool  `json:"mrp_promotion"`
            DiscountRules []DiscountRules  `json:"discount_rules"`
            PromotionName string  `json:"promotion_name"`
            BuyRules []BuyRules  `json:"buy_rules"`
            PromoID string  `json:"promo_id"`
            Amount float64  `json:"amount"`
         
    }
    
    // BagGST used by Order
    type BagGST struct {

        
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            HsnCode string  `json:"hsn_code"`
            GstTag string  `json:"gst_tag"`
            GstinCode string  `json:"gstin_code"`
            IsDefaultHsnCode bool  `json:"is_default_hsn_code"`
            GstFee float64  `json:"gst_fee"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            ValueOfGood float64  `json:"value_of_good"`
         
    }
    
    // OrderBrandName used by Order
    type OrderBrandName struct {

        
            ID float64  `json:"id"`
            BrandName string  `json:"brand_name"`
            Company string  `json:"company"`
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
            Logo string  `json:"logo"`
         
    }
    
    // CurrentStatus used by Order
    type CurrentStatus struct {

        
            Status string  `json:"status"`
            BagStateMapper BagStateMapper  `json:"bag_state_mapper"`
            CurrentStatusID float64  `json:"current_status_id"`
            StoreID float64  `json:"store_id"`
            DeliveryPartnerID float64  `json:"delivery_partner_id"`
            KafkaSync bool  `json:"kafka_sync"`
            StateType string  `json:"state_type"`
            UpdatedAt string  `json:"updated_at"`
            ShipmentID string  `json:"shipment_id"`
            BagID float64  `json:"bag_id"`
            DeliveryAwbNumber string  `json:"delivery_awb_number"`
            StateID float64  `json:"state_id"`
            CreatedAt string  `json:"created_at"`
         
    }
    
    // BagConfigs used by Order
    type BagConfigs struct {

        
            IsReturnable bool  `json:"is_returnable"`
            CanBeCancelled bool  `json:"can_be_cancelled"`
            AllowForceReturn bool  `json:"allow_force_return"`
            IsActive bool  `json:"is_active"`
            IsCustomerReturnAllowed bool  `json:"is_customer_return_allowed"`
            EnableTracking bool  `json:"enable_tracking"`
         
    }
    
    // OrderBagArticle used by Order
    type OrderBagArticle struct {

        
            ReturnConfig map[string]interface{}  `json:"return_config"`
            Identifiers map[string]interface{}  `json:"identifiers"`
            UID string  `json:"uid"`
         
    }
    
    // GiftCard used by Order
    type GiftCard struct {

        
            GiftMessage string  `json:"gift_message"`
            IsGiftApplied bool  `json:"is_gift_applied"`
            DisplayText string  `json:"display_text"`
            GiftPrice float64  `json:"gift_price"`
         
    }
    
    // B2BPODetails used by Order
    type B2BPODetails struct {

        
            ItemBasePrice float64  `json:"item_base_price"`
            PartialCanRet bool  `json:"partial_can_ret"`
            PoLineAmount float64  `json:"po_line_amount"`
            TotalGstPercentage float64  `json:"total_gst_percentage"`
            PoTaxAmount float64  `json:"po_tax_amount"`
            DockerNumber string  `json:"docker_number"`
         
    }
    
    // BagMeta used by Order
    type BagMeta struct {

        
            GiftCard GiftCard  `json:"gift_card"`
            GroupID string  `json:"group_id"`
            PartialCanRet bool  `json:"partial_can_ret"`
            CustomMessage string  `json:"custom_message"`
            B2bPoDetails B2BPODetails  `json:"b2b_po_details"`
            CustomJson map[string]interface{}  `json:"custom_json"`
            DocketNumber string  `json:"docket_number"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
         
    }
    
    // OrderBags used by Order
    type OrderBags struct {

        
            CanCancel bool  `json:"can_cancel"`
            AppliedPromos []AppliedPromos  `json:"applied_promos"`
            EntityType string  `json:"entity_type"`
            GstDetails BagGST  `json:"gst_details"`
            Prices Prices  `json:"prices"`
            CanReturn bool  `json:"can_return"`
            FinancialBreakup FinancialBreakup  `json:"financial_breakup"`
            Item PlatformItem  `json:"item"`
            Brand OrderBrandName  `json:"brand"`
            LineNumber float64  `json:"line_number"`
            ParentPromoBags map[string]interface{}  `json:"parent_promo_bags"`
            Identifier string  `json:"identifier"`
            DeliveryAddress PlatformDeliveryAddress  `json:"delivery_address"`
            SellerIdentifier string  `json:"seller_identifier"`
            DisplayName string  `json:"display_name"`
            GroupID string  `json:"group_id"`
            CurrentStatus CurrentStatus  `json:"current_status"`
            BagID float64  `json:"bag_id"`
            Quantity float64  `json:"quantity"`
            BagConfigs BagConfigs  `json:"bag_configs"`
            Article OrderBagArticle  `json:"article"`
            Meta BagMeta  `json:"meta"`
         
    }
    
    // ShipmentPayments used by Order
    type ShipmentPayments struct {

        
            Source string  `json:"source"`
            Mode string  `json:"mode"`
            Logo string  `json:"logo"`
         
    }
    
    // TrackingList used by Order
    type TrackingList struct {

        
            Status string  `json:"status"`
            IsCurrent bool  `json:"is_current"`
            IsPassed bool  `json:"is_passed"`
            Text string  `json:"text"`
            Time string  `json:"time"`
         
    }
    
    // OrderingStoreDetails used by Order
    type OrderingStoreDetails struct {

        
            Code string  `json:"code"`
            Pincode string  `json:"pincode"`
            City string  `json:"city"`
            State string  `json:"state"`
            OrderingStoreID float64  `json:"ordering_store_id"`
            Phone string  `json:"phone"`
            ContactPerson string  `json:"contact_person"`
            StoreName string  `json:"store_name"`
            Address string  `json:"address"`
            Country string  `json:"country"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // FulfillingStore used by Order
    type FulfillingStore struct {

        
            Code string  `json:"code"`
            ID float64  `json:"id"`
            FulfillmentChannel string  `json:"fulfillment_channel"`
            Pincode string  `json:"pincode"`
            City string  `json:"city"`
            State string  `json:"state"`
            Phone string  `json:"phone"`
            ContactPerson string  `json:"contact_person"`
            StoreName string  `json:"store_name"`
            Address string  `json:"address"`
            Country string  `json:"country"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // DPDetailsData used by Order
    type DPDetailsData struct {

        
            TrackURL string  `json:"track_url"`
            ID float64  `json:"id"`
            Pincode string  `json:"pincode"`
            AwbNo string  `json:"awb_no"`
            GstTag string  `json:"gst_tag"`
            EwayBillID string  `json:"eway_bill_id"`
            Name string  `json:"name"`
            Country string  `json:"country"`
         
    }
    
    // PDFLinks used by Order
    type PDFLinks struct {

        
            CreditNoteURL string  `json:"credit_note_url"`
            PoInvoice string  `json:"po_invoice"`
            LabelType string  `json:"label_type"`
            DeliveryChallanA4 string  `json:"delivery_challan_a4"`
            InvoiceA4 string  `json:"invoice_a4"`
            InvoiceType string  `json:"invoice_type"`
            LabelA6 string  `json:"label_a6"`
            LabelA4 string  `json:"label_a4"`
            InvoiceA6 string  `json:"invoice_a6"`
            B2b string  `json:"b2b"`
            LabelPos string  `json:"label_pos"`
            InvoicePos string  `json:"invoice_pos"`
            Label string  `json:"label"`
            Invoice string  `json:"invoice"`
         
    }
    
    // LockData used by Order
    type LockData struct {

        
            Mto bool  `json:"mto"`
            LockMessage string  `json:"lock_message"`
            Locked bool  `json:"locked"`
         
    }
    
    // Formatted used by Order
    type Formatted struct {

        
            FMin string  `json:"f_min"`
            FMax string  `json:"f_max"`
         
    }
    
    // ShipmentTimeStamp used by Order
    type ShipmentTimeStamp struct {

        
            TMin string  `json:"t_min"`
            TMax string  `json:"t_max"`
         
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

        
            Pincode float64  `json:"pincode"`
            City string  `json:"city"`
            State string  `json:"state"`
            Name string  `json:"name"`
            AjioSiteID string  `json:"ajio_site_id"`
            Address string  `json:"address"`
            Gstin string  `json:"gstin"`
         
    }
    
    // ShipmentMeta used by Order
    type ShipmentMeta struct {

        
            SameStoreAvailable bool  `json:"same_store_available"`
            AutoTriggerDpAssignmentAcf bool  `json:"auto_trigger_dp_assignment_acf"`
            ReturnAffiliateOrderID string  `json:"return_affiliate_order_id"`
            DpSortKey string  `json:"dp_sort_key"`
            BagWeight map[string]interface{}  `json:"bag_weight"`
            LockData LockData  `json:"lock_data"`
            MarketplaceStoreID string  `json:"marketplace_store_id"`
            AwbNumber string  `json:"awb_number"`
            Weight float64  `json:"weight"`
            B2cBuyerDetails map[string]interface{}  `json:"b2c_buyer_details"`
            ShipmentVolumetricWeight float64  `json:"shipment_volumetric_weight"`
            ReturnDetails map[string]interface{}  `json:"return_details"`
            PackagingName string  `json:"packaging_name"`
            PoNumber string  `json:"po_number"`
            EwaybillInfo map[string]interface{}  `json:"ewaybill_info"`
            BoxType string  `json:"box_type"`
            DpOptions map[string]interface{}  `json:"dp_options"`
            ReturnAffiliateShipmentID string  `json:"return_affiliate_shipment_id"`
            Formatted Formatted  `json:"formatted"`
            ShipmentWeight float64  `json:"shipment_weight"`
            External map[string]interface{}  `json:"external"`
            FulfilmentPriorityText string  `json:"fulfilment_priority_text"`
            ReturnAwbNumber string  `json:"return_awb_number"`
            Timestamp ShipmentTimeStamp  `json:"timestamp"`
            DpID string  `json:"dp_id"`
            ReturnStoreNode float64  `json:"return_store_node"`
            EinvoiceInfo EinvoiceInfo  `json:"einvoice_info"`
            DebugInfo DebugInfo  `json:"debug_info"`
            OrderType string  `json:"order_type"`
            DpName string  `json:"dp_name"`
            ForwardAffiliateShipmentID string  `json:"forward_affiliate_shipment_id"`
            ForwardAffiliateOrderID string  `json:"forward_affiliate_order_id"`
            DueDate string  `json:"due_date"`
            B2bBuyerDetails BuyerDetails  `json:"b2b_buyer_details"`
            AssignDpFromSb bool  `json:"assign_dp_from_sb"`
            StoreInvoiceUpdatedDate string  `json:"store_invoice_updated_date"`
         
    }
    
    // AffiliateDetails used by Order
    type AffiliateDetails struct {

        
            PdfLinks PDFLinks  `json:"pdf_links"`
            AffiliateID string  `json:"affiliate_id"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            ShipmentMeta ShipmentMeta  `json:"shipment_meta"`
            AffiliateStoreID string  `json:"affiliate_store_id"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
            AffiliateMeta AffiliateMeta  `json:"affiliate_meta"`
            AdID string  `json:"ad_id"`
            CompanyAffiliateTag string  `json:"company_affiliate_tag"`
         
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

        
            CompanyID float64  `json:"company_id"`
            CompanyName string  `json:"company_name"`
            CompanyCin string  `json:"company_cin"`
            Address map[string]interface{}  `json:"address"`
            CompanyContact ContactDetails  `json:"company_contact"`
            CompanyGst string  `json:"company_gst"`
         
    }
    
    // OrderDetailsData used by Order
    type OrderDetailsData struct {

        
            AffiliateID string  `json:"affiliate_id"`
            TaxDetails map[string]interface{}  `json:"tax_details"`
            CodCharges string  `json:"cod_charges"`
            OrderingChannelLogo map[string]interface{}  `json:"ordering_channel_logo"`
            OrderDate string  `json:"order_date"`
            OrderValue string  `json:"order_value"`
            FyndOrderID string  `json:"fynd_order_id"`
            Source string  `json:"source"`
            OrderingChannel string  `json:"ordering_channel"`
         
    }
    
    // InvoiceInfo used by Order
    type InvoiceInfo struct {

        
            CreditNoteID string  `json:"credit_note_id"`
            StoreInvoiceID string  `json:"store_invoice_id"`
            InvoiceURL string  `json:"invoice_url"`
            UpdatedDate string  `json:"updated_date"`
            LabelURL string  `json:"label_url"`
         
    }
    
    // Meta used by Order
    type Meta struct {

        
            Dimension Dimensions  `json:"dimension"`
         
    }
    
    // PlatformShipment used by Order
    type PlatformShipment struct {

        
            Status ShipmentStatusData  `json:"status"`
            DeliveryDetails UserDetailsData  `json:"delivery_details"`
            DeliverySlot map[string]interface{}  `json:"delivery_slot"`
            EnableDpTracking bool  `json:"enable_dp_tracking"`
            BillingDetails UserDetailsData  `json:"billing_details"`
            JourneyType string  `json:"journey_type"`
            GstDetails GSTDetailsData  `json:"gst_details"`
            ForwardShipmentID string  `json:"forward_shipment_id"`
            TotalBags float64  `json:"total_bags"`
            ShipmentID string  `json:"shipment_id"`
            Bags []OrderBags  `json:"bags"`
            Prices Prices  `json:"prices"`
            PackagingType string  `json:"packaging_type"`
            FulfilmentPriority float64  `json:"fulfilment_priority"`
            User UserDataInfo  `json:"user"`
            Coupon map[string]interface{}  `json:"coupon"`
            PickedDate string  `json:"picked_date"`
            LockStatus bool  `json:"lock_status"`
            Vertical string  `json:"vertical"`
            Payments ShipmentPayments  `json:"payments"`
            TrackingList []TrackingList  `json:"tracking_list"`
            OperationalStatus string  `json:"operational_status"`
            OrderingStore OrderingStoreDetails  `json:"ordering_store"`
            FulfillingStore FulfillingStore  `json:"fulfilling_store"`
            PaymentMethods map[string]interface{}  `json:"payment_methods"`
            BagStatusHistory []BagStatusHistory  `json:"bag_status_history"`
            InvoiceID string  `json:"invoice_id"`
            ShipmentImages []string  `json:"shipment_images"`
            DpDetails DPDetailsData  `json:"dp_details"`
            TotalItems float64  `json:"total_items"`
            CustomMeta []map[string]interface{}  `json:"custom_meta"`
            UserAgent string  `json:"user_agent"`
            ShipmentQuantity float64  `json:"shipment_quantity"`
            AffiliateDetails AffiliateDetails  `json:"affiliate_details"`
            ShipmentStatus string  `json:"shipment_status"`
            PlatformLogo string  `json:"platform_logo"`
            PaymentMode string  `json:"payment_mode"`
            PriorityText string  `json:"priority_text"`
            CompanyDetails CompanyDetails  `json:"company_details"`
            Order OrderDetailsData  `json:"order"`
            Invoice InvoiceInfo  `json:"invoice"`
            Meta Meta  `json:"meta"`
         
    }
    
    // ShipmentInfoResponse used by Order
    type ShipmentInfoResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
            Shipments []PlatformShipment  `json:"shipments"`
         
    }
    
    // AssetByShipment used by Order
    type AssetByShipment struct {

        
            Assets map[string]string  `json:"assets"`
            ExpiresIn string  `json:"expires_in"`
            ShipmentID string  `json:"shipment_id"`
         
    }
    
    // ResponseGetAssetShipment used by Order
    type ResponseGetAssetShipment struct {

        
            PresignedType string  `json:"presigned_type"`
            Result []AssetByShipment  `json:"result"`
            Success bool  `json:"success"`
         
    }
    
    // TaxDetails used by Order
    type TaxDetails struct {

        
            PanNo string  `json:"pan_no"`
            Gstin string  `json:"gstin"`
         
    }
    
    // PlatformUserDetails used by Order
    type PlatformUserDetails struct {

        
            PlatformUserLastName string  `json:"platform_user_last_name"`
            PlatformUserID string  `json:"platform_user_id"`
            PlatformUserEmployeeCode string  `json:"platform_user_employee_code"`
            PlatformUserFirstName string  `json:"platform_user_first_name"`
         
    }
    
    // BillingStaffDetails used by Order
    type BillingStaffDetails struct {

        
            StaffID float64  `json:"staff_id"`
            LastName string  `json:"last_name"`
            FirstName string  `json:"first_name"`
            EmployeeCode string  `json:"employee_code"`
            User string  `json:"user"`
         
    }
    
    // TransactionData used by Order
    type TransactionData struct {

        
            Status string  `json:"status"`
            Currency string  `json:"currency"`
            TerminalID string  `json:"terminal_id"`
            PaymentID string  `json:"payment_id"`
            TransactionID string  `json:"transaction_id"`
            Entity string  `json:"entity"`
            AmountPaid string  `json:"amount_paid"`
            UniqueReferenceNumber string  `json:"unique_reference_number"`
         
    }
    
    // OrderMeta used by Order
    type OrderMeta struct {

        
            Files []map[string]interface{}  `json:"files"`
            Staff map[string]interface{}  `json:"staff"`
            CurrencySymbol string  `json:"currency_symbol"`
            PlatformUserDetails PlatformUserDetails  `json:"platform_user_details"`
            CustomerNote string  `json:"customer_note"`
            OrderingStore float64  `json:"ordering_store"`
            OrderPlatform string  `json:"order_platform"`
            OrderChildEntities []string  `json:"order_child_entities"`
            CartID float64  `json:"cart_id"`
            BillingStaffDetails BillingStaffDetails  `json:"billing_staff_details"`
            TransactionData TransactionData  `json:"transaction_data"`
            CompanyLogo string  `json:"company_logo"`
            Comment string  `json:"comment"`
            OrderType string  `json:"order_type"`
            MongoCartID float64  `json:"mongo_cart_id"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            PaymentType string  `json:"payment_type"`
            OrderTags []map[string]interface{}  `json:"order_tags"`
            EmployeeID float64  `json:"employee_id"`
         
    }
    
    // OrderDict used by Order
    type OrderDict struct {

        
            TaxDetails TaxDetails  `json:"tax_details"`
            OrderDate string  `json:"order_date"`
            PaymentMethods map[string]interface{}  `json:"payment_methods"`
            FyndOrderID string  `json:"fynd_order_id"`
            Prices Prices  `json:"prices"`
            Meta OrderMeta  `json:"meta"`
         
    }
    
    // ShipmentDetailsResponse used by Order
    type ShipmentDetailsResponse struct {

        
            Order OrderDict  `json:"order"`
            Success bool  `json:"success"`
            Shipments []PlatformShipment  `json:"shipments"`
         
    }
    
    // SubLane used by Order
    type SubLane struct {

        
            Actions []map[string]interface{}  `json:"actions"`
            Value string  `json:"value"`
            Text string  `json:"text"`
            Index float64  `json:"index"`
            TotalItems float64  `json:"total_items"`
         
    }
    
    // SuperLane used by Order
    type SuperLane struct {

        
            Value string  `json:"value"`
            Text string  `json:"text"`
            Options []SubLane  `json:"options"`
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

        
            Value string  `json:"value"`
            Display string  `json:"display"`
            Name string  `json:"name"`
         
    }
    
    // PlatformOrderItems used by Order
    type PlatformOrderItems struct {

        
            OrderCreatedTime string  `json:"order_created_time"`
            PaymentMode string  `json:"payment_mode"`
            Shipments []PlatformShipment  `json:"shipments"`
            TotalOrderValue float64  `json:"total_order_value"`
            Channel PlatformChannel  `json:"channel"`
            OrderValue float64  `json:"order_value"`
            OrderID string  `json:"order_id"`
            UserInfo UserDataInfo  `json:"user_info"`
            BreakupValues []PlatformBreakupValues  `json:"breakup_values"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // OrderListingResponse used by Order
    type OrderListingResponse struct {

        
            Items []PlatformOrderItems  `json:"items"`
            Lane string  `json:"lane"`
            TotalCount float64  `json:"total_count"`
            Message string  `json:"message"`
            Page Page  `json:"page"`
            Success bool  `json:"success"`
         
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

        
            Status string  `json:"status"`
            RawStatus string  `json:"raw_status"`
            Awb string  `json:"awb"`
            AccountName string  `json:"account_name"`
            Reason string  `json:"reason"`
            ShipmentType string  `json:"shipment_type"`
            UpdatedAt string  `json:"updated_at"`
            Meta map[string]interface{}  `json:"meta"`
            LastLocationRecievedAt string  `json:"last_location_recieved_at"`
            UpdatedTime string  `json:"updated_time"`
         
    }
    
    // PlatformShipmentTrack used by Order
    type PlatformShipmentTrack struct {

        
            Results []PlatformTrack  `json:"results"`
            Meta map[string]interface{}  `json:"meta"`
         
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
            Text string  `json:"text"`
            Options []FilterInfoOption  `json:"options"`
            Type string  `json:"type"`
         
    }
    
    // AdvanceFilterInfo used by Order
    type AdvanceFilterInfo struct {

        
            ActionCentre []FiltersInfo  `json:"action_centre"`
            Returned []FiltersInfo  `json:"returned"`
            Unfulfilled []FiltersInfo  `json:"unfulfilled"`
            Processed []FiltersInfo  `json:"processed"`
            Filters []FiltersInfo  `json:"filters"`
         
    }
    
    // FiltersResponse used by Order
    type FiltersResponse struct {

        
            GlobalFilter []FiltersInfo  `json:"global_filter"`
            AdvanceFilter AdvanceFilterInfo  `json:"advance_filter"`
         
    }
    
    // Success used by Order
    type Success struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // OmsReports used by Order
    type OmsReports struct {

        
            Status string  `json:"status"`
            ReportName string  `json:"report_name"`
            ReportRequestedAt string  `json:"report_requested_at"`
            ReportType string  `json:"report_type"`
            ReportID string  `json:"report_id"`
            S3Key string  `json:"s3_key"`
            ReportCreatedAt string  `json:"report_created_at"`
            RequestDetails map[string]interface{}  `json:"request_details"`
            DisplayName string  `json:"display_name"`
         
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

        
            TraceID string  `json:"trace_id"`
            Data []map[string]interface{}  `json:"data"`
            Identifier string  `json:"identifier"`
            Error []NestedErrorSchemaDataSet  `json:"error"`
            Success bool  `json:"success"`
         
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

        
            FilePath string  `json:"file_path"`
            Upload FileUploadResponse  `json:"upload"`
            FileName string  `json:"file_name"`
            Size float64  `json:"size"`
            Method string  `json:"method"`
            Tags []string  `json:"tags"`
            Operation string  `json:"operation"`
            Namespace string  `json:"namespace"`
            Cdn URL  `json:"cdn"`
            ContentType string  `json:"content_type"`
         
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

        
            ID float64  `json:"id"`
            QcType []string  `json:"qc_type"`
            QuestionSet []QuestionSet  `json:"question_set"`
            DisplayName string  `json:"display_name"`
         
    }
    
    // PlatformShipmentReasonsResponse used by Order
    type PlatformShipmentReasonsResponse struct {

        
            Success bool  `json:"success"`
            Reasons []Reason  `json:"reasons"`
         
    }
    
    // BagReturnableCancelableStatus1 used by Order
    type BagReturnableCancelableStatus1 struct {

        
            IsReturnable bool  `json:"is_returnable"`
            CanBeCancelled bool  `json:"can_be_cancelled"`
            IsActive bool  `json:"is_active"`
            IsCustomerReturnAllowed bool  `json:"is_customer_return_allowed"`
            EnableTracking bool  `json:"enable_tracking"`
         
    }
    
    // BagGSTDetails used by Order
    type BagGSTDetails struct {

        
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            SgstGstFee string  `json:"sgst_gst_fee"`
            HsnCode string  `json:"hsn_code"`
            CgstGstFee string  `json:"cgst_gst_fee"`
            HsnCodeID string  `json:"hsn_code_id"`
            GstTag string  `json:"gst_tag"`
            GstinCode string  `json:"gstin_code"`
            SgstTaxPercentage float64  `json:"sgst_tax_percentage"`
            IgstGstFee string  `json:"igst_gst_fee"`
            IsDefaultHsnCode bool  `json:"is_default_hsn_code"`
            GstFee float64  `json:"gst_fee"`
            IgstTaxPercentage float64  `json:"igst_tax_percentage"`
            CgstTaxPercentage float64  `json:"cgst_tax_percentage"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            ValueOfGood float64  `json:"value_of_good"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
         
    }
    
    // Attributes used by Order
    type Attributes struct {

        
            PrimaryColorHex string  `json:"primary_color_hex"`
            Gender []string  `json:"gender"`
            BrandName string  `json:"brand_name"`
            Essential string  `json:"essential"`
            Name string  `json:"name"`
            PrimaryColor string  `json:"primary_color"`
            PrimaryMaterial string  `json:"primary_material"`
            MarketerName string  `json:"marketer_name"`
            MarketerAddress string  `json:"marketer_address"`
         
    }
    
    // Item used by Order
    type Item struct {

        
            L2CategoryID float64  `json:"l2_category_id"`
            CanCancel bool  `json:"can_cancel"`
            Image []string  `json:"image"`
            L1Category []string  `json:"l1_category"`
            L1CategoryID float64  `json:"l1_category_id"`
            BrandID float64  `json:"brand_id"`
            L3Category float64  `json:"l3_category"`
            Name string  `json:"name"`
            CanReturn bool  `json:"can_return"`
            ItemID float64  `json:"item_id"`
            Brand string  `json:"brand"`
            Gender string  `json:"gender"`
            L2Category []string  `json:"l2_category"`
            BranchURL string  `json:"branch_url"`
            SlugKey string  `json:"slug_key"`
            Color string  `json:"color"`
            Code string  `json:"code"`
            LastUpdatedAt string  `json:"last_updated_at"`
            WebstoreProductURL string  `json:"webstore_product_url"`
            DepartmentID float64  `json:"department_id"`
            Attributes Attributes  `json:"attributes"`
            Size string  `json:"size"`
            L3CategoryName string  `json:"l3_category_name"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // Brand used by Order
    type Brand struct {

        
            BrandName string  `json:"brand_name"`
            Company string  `json:"company"`
            IsVirtualInvoice bool  `json:"is_virtual_invoice"`
            CreatedOn float64  `json:"created_on"`
            StartDate string  `json:"start_date"`
            ModifiedOn float64  `json:"modified_on"`
            BrandID float64  `json:"brand_id"`
            PickupLocation string  `json:"pickup_location"`
            CreditNoteExpiryDays float64  `json:"credit_note_expiry_days"`
            ScriptLastRan string  `json:"script_last_ran"`
            Logo string  `json:"logo"`
            InvoicePrefix string  `json:"invoice_prefix"`
            CreditNoteAllowed bool  `json:"credit_note_allowed"`
         
    }
    
    // StoreAddress used by Order
    type StoreAddress struct {

        
            AddressType string  `json:"address_type"`
            Landmark string  `json:"landmark"`
            CreatedAt string  `json:"created_at"`
            Version string  `json:"version"`
            CountryCode string  `json:"country_code"`
            State string  `json:"state"`
            Phone string  `json:"phone"`
            Country string  `json:"country"`
            Pincode float64  `json:"pincode"`
            City string  `json:"city"`
            Area string  `json:"area"`
            Longitude float64  `json:"longitude"`
            Latitude float64  `json:"latitude"`
            Address2 string  `json:"address2"`
            AddressCategory string  `json:"address_category"`
            UpdatedAt string  `json:"updated_at"`
            ContactPerson string  `json:"contact_person"`
            Address1 string  `json:"address1"`
            Email string  `json:"email"`
         
    }
    
    // StoreEinvoice used by Order
    type StoreEinvoice struct {

        
            Password string  `json:"password"`
            Enabled bool  `json:"enabled"`
            Username string  `json:"username"`
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

        
            Password string  `json:"password"`
            User string  `json:"user"`
            Username string  `json:"username"`
         
    }
    
    // Document used by Order
    type Document struct {

        
            URL string  `json:"url"`
            Verified bool  `json:"verified"`
            Value string  `json:"value"`
            LegalName string  `json:"legal_name"`
            DsType string  `json:"ds_type"`
         
    }
    
    // StoreDocuments used by Order
    type StoreDocuments struct {

        
            Gst Document  `json:"gst"`
         
    }
    
    // StoreMeta used by Order
    type StoreMeta struct {

        
            AdditionalContactDetails map[string]interface{}  `json:"additional_contact_details"`
            Timing []map[string]interface{}  `json:"timing"`
            Stage string  `json:"stage"`
            EwaybillPortalDetails map[string]interface{}  `json:"ewaybill_portal_details"`
            GstCredentials StoreGstCredentials  `json:"gst_credentials"`
            ProductReturnConfig map[string]interface{}  `json:"product_return_config"`
            GstNumber string  `json:"gst_number"`
            NotificationEmails []string  `json:"notification_emails"`
            EinvoicePortalDetails EInvoicePortalDetails  `json:"einvoice_portal_details"`
            Documents StoreDocuments  `json:"documents"`
            DisplayName string  `json:"display_name"`
         
    }
    
    // Store used by Order
    type Store struct {

        
            ParentStoreID float64  `json:"parent_store_id"`
            BrandID interface{}  `json:"brand_id"`
            IsEnabledForRecon bool  `json:"is_enabled_for_recon"`
            SID string  `json:"s_id"`
            Name string  `json:"name"`
            MallArea string  `json:"mall_area"`
            CreatedAt string  `json:"created_at"`
            State string  `json:"state"`
            VatNo string  `json:"vat_no"`
            Phone float64  `json:"phone"`
            StoreEmail string  `json:"store_email"`
            PackagingMaterialCount float64  `json:"packaging_material_count"`
            Country string  `json:"country"`
            LoginUsername string  `json:"login_username"`
            Code string  `json:"code"`
            StoreActiveFrom string  `json:"store_active_from"`
            OrderIntegrationID string  `json:"order_integration_id"`
            Pincode string  `json:"pincode"`
            City string  `json:"city"`
            IsArchived bool  `json:"is_archived"`
            Longitude float64  `json:"longitude"`
            StoreAddressJson StoreAddress  `json:"store_address_json"`
            Latitude float64  `json:"latitude"`
            LocationType string  `json:"location_type"`
            AlohomoraUserID float64  `json:"alohomora_user_id"`
            BrandStoreTags []string  `json:"brand_store_tags"`
            Address2 string  `json:"address2"`
            FulfillmentChannel string  `json:"fulfillment_channel"`
            CompanyID float64  `json:"company_id"`
            IsActive bool  `json:"is_active"`
            MallName string  `json:"mall_name"`
            UpdatedAt string  `json:"updated_at"`
            ContactPerson string  `json:"contact_person"`
            Address1 string  `json:"address1"`
            Meta StoreMeta  `json:"meta"`
         
    }
    
    // ArticleDetails used by Order
    type ArticleDetails struct {

        
            Status map[string]interface{}  `json:"status"`
         
    }
    
    // BagDetailsPlatformResponse used by Order
    type BagDetailsPlatformResponse struct {

        
            Status BagReturnableCancelableStatus1  `json:"status"`
            AppliedPromos []map[string]interface{}  `json:"applied_promos"`
            EntityType string  `json:"entity_type"`
            BagUpdateTime float64  `json:"bag_update_time"`
            GstDetails BagGSTDetails  `json:"gst_details"`
            JourneyType string  `json:"journey_type"`
            Prices Prices  `json:"prices"`
            ShipmentID string  `json:"shipment_id"`
            FinancialBreakup []FinancialBreakup  `json:"financial_breakup"`
            Item Item  `json:"item"`
            BagStatus []BagStatusHistory  `json:"bag_status"`
            BID float64  `json:"b_id"`
            Brand Brand  `json:"brand"`
            RestorePromos map[string]interface{}  `json:"restore_promos"`
            LineNumber float64  `json:"line_number"`
            OperationalStatus string  `json:"operational_status"`
            OrderingStore Store  `json:"ordering_store"`
            Identifier string  `json:"identifier"`
            ParentPromoBags map[string]interface{}  `json:"parent_promo_bags"`
            BagStatusHistory BagStatusHistory  `json:"bag_status_history"`
            SellerIdentifier string  `json:"seller_identifier"`
            NoOfBagsOrder float64  `json:"no_of_bags_order"`
            BType string  `json:"b_type"`
            AffiliateBagDetails AffiliateBagDetails  `json:"affiliate_bag_details"`
            DisplayName string  `json:"display_name"`
            OrderIntegrationID string  `json:"order_integration_id"`
            RestoreCoupon bool  `json:"restore_coupon"`
            CurrentStatus BagStatusHistory  `json:"current_status"`
            Tags []string  `json:"tags"`
            AffiliateDetails AffiliateDetails  `json:"affiliate_details"`
            ArticleDetails ArticleDetails  `json:"article_details"`
            OriginalBagList []float64  `json:"original_bag_list"`
            Reasons []map[string]interface{}  `json:"reasons"`
            Dates Dates  `json:"dates"`
            QcRequired interface{}  `json:"qc_required"`
            Quantity float64  `json:"quantity"`
            CurrentOperationalStatus BagStatusHistory  `json:"current_operational_status"`
            Article Article  `json:"article"`
            Meta BagMeta  `json:"meta"`
         
    }
    
    // ErrorResponse used by Order
    type ErrorResponse struct {

        
            Message string  `json:"message"`
            Error string  `json:"error"`
         
    }
    
    // BagsPage used by Order
    type BagsPage struct {

        
            ItemTotal float64  `json:"item_total"`
            PageType string  `json:"page_type"`
            Size float64  `json:"size"`
            HasNext bool  `json:"has_next"`
            Current float64  `json:"current"`
         
    }
    
    // GetBagsPlatformResponse used by Order
    type GetBagsPlatformResponse struct {

        
            Page BagsPage  `json:"page"`
            Items []BagDetailsPlatformResponse  `json:"items"`
         
    }
    
    // GeneratePosOrderReceiptResponse used by Order
    type GeneratePosOrderReceiptResponse struct {

        
            PaymentReceipt string  `json:"payment_receipt"`
            InvoiceReceipt string  `json:"invoice_receipt"`
            Success bool  `json:"success"`
            OrderID string  `json:"order_id"`
         
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
            Success bool  `json:"success"`
            Message string  `json:"message"`
            ErrorTrace string  `json:"error_trace"`
         
    }
    
    // StoreReassign used by Order
    type StoreReassign struct {

        
            ReasonIds []float64  `json:"reason_ids"`
            MongoArticleID string  `json:"mongo_article_id"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
            FyndOrderID string  `json:"fynd_order_id"`
            SetID string  `json:"set_id"`
            AffiliateID string  `json:"affiliate_id"`
            ItemID string  `json:"item_id"`
            BagID float64  `json:"bag_id"`
            StoreID float64  `json:"store_id"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
         
    }
    
    // StoreReassignResponse used by Order
    type StoreReassignResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // Entities used by Order
    type Entities struct {

        
            AffiliateBagID string  `json:"affiliate_bag_id"`
            AffiliateID string  `json:"affiliate_id"`
            ReasonText string  `json:"reason_text"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            ID string  `json:"id"`
         
    }
    
    // UpdateShipmentLockPayload used by Order
    type UpdateShipmentLockPayload struct {

        
            EntityType string  `json:"entity_type"`
            ActionType string  `json:"action_type"`
            Action string  `json:"action"`
            Entities []Entities  `json:"entities"`
         
    }
    
    // Bags used by Order
    type Bags struct {

        
            AffiliateOrderID string  `json:"affiliate_order_id"`
            IsLocked bool  `json:"is_locked"`
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

        
            Status string  `json:"status"`
            ShipmentID string  `json:"shipment_id"`
            LockStatus string  `json:"lock_status"`
            IsShipmentLocked bool  `json:"is_shipment_locked"`
            IsBagLocked bool  `json:"is_bag_locked"`
            AffiliateID string  `json:"affiliate_id"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            Bags []Bags  `json:"bags"`
            OriginalFilter OriginalFilter  `json:"original_filter"`
         
    }
    
    // UpdateShipmentLockResponse used by Order
    type UpdateShipmentLockResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
            CheckResponse []CheckResponse  `json:"check_response"`
         
    }
    
    // AnnouncementResponse used by Order
    type AnnouncementResponse struct {

        
            LogoURL string  `json:"logo_url"`
            FromDatetime string  `json:"from_datetime"`
            PlatformID string  `json:"platform_id"`
            PlatformName string  `json:"platform_name"`
            ToDatetime string  `json:"to_datetime"`
            Description string  `json:"description"`
            CompanyID float64  `json:"company_id"`
            CreatedAt string  `json:"created_at"`
            Title string  `json:"title"`
            ID float64  `json:"id"`
         
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
            Identifier string  `json:"identifier"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // ShipmentsRequest used by Order
    type ShipmentsRequest struct {

        
            Reasons ReasonsData  `json:"reasons"`
            Identifier string  `json:"identifier"`
            DataUpdates DataUpdates  `json:"data_updates"`
            Products []Products  `json:"products"`
         
    }
    
    // StatuesRequest used by Order
    type StatuesRequest struct {

        
            Status string  `json:"status"`
            Shipments []ShipmentsRequest  `json:"shipments"`
            ExcludeBagsNextState string  `json:"exclude_bags_next_state"`
         
    }
    
    // UpdateShipmentStatusRequest used by Order
    type UpdateShipmentStatusRequest struct {

        
            Statuses []StatuesRequest  `json:"statuses"`
            Task bool  `json:"task"`
            ForceTransition bool  `json:"force_transition"`
            LockAfterTransition bool  `json:"lock_after_transition"`
            UnlockBeforeTransition bool  `json:"unlock_before_transition"`
         
    }
    
    // ShipmentsResponse used by Order
    type ShipmentsResponse struct {

        
            Meta map[string]interface{}  `json:"meta"`
            Message string  `json:"message"`
            StackTrace string  `json:"stack_trace"`
            Status float64  `json:"status"`
            Identifier string  `json:"identifier"`
            Code string  `json:"code"`
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

        
            Meta []AffiliateAppConfigMeta  `json:"meta"`
            Secret string  `json:"secret"`
            Token string  `json:"token"`
            UpdatedAt string  `json:"updated_at"`
            Owner string  `json:"owner"`
            Name string  `json:"name"`
            Description string  `json:"description"`
            CreatedAt string  `json:"created_at"`
            ID string  `json:"id"`
         
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
    
    // AffiliateInventoryOrderConfig used by Order
    type AffiliateInventoryOrderConfig struct {

        
            ForceReassignment bool  `json:"force_reassignment"`
         
    }
    
    // AffiliateInventoryArticleAssignmentConfig used by Order
    type AffiliateInventoryArticleAssignmentConfig struct {

        
            PostOrderReassignment bool  `json:"post_order_reassignment"`
         
    }
    
    // AffiliateInventoryLogisticsConfig used by Order
    type AffiliateInventoryLogisticsConfig struct {

        
            DpAssignment bool  `json:"dp_assignment"`
         
    }
    
    // AffiliateInventoryConfig used by Order
    type AffiliateInventoryConfig struct {

        
            Payment AffiliateInventoryPaymentConfig  `json:"payment"`
            Inventory AffiliateInventoryStoreConfig  `json:"inventory"`
            Order AffiliateInventoryOrderConfig  `json:"order"`
            ArticleAssignment AffiliateInventoryArticleAssignmentConfig  `json:"article_assignment"`
            Logistics AffiliateInventoryLogisticsConfig  `json:"logistics"`
         
    }
    
    // AffiliateConfig used by Order
    type AffiliateConfig struct {

        
            App AffiliateAppConfig  `json:"app"`
            Inventory AffiliateInventoryConfig  `json:"inventory"`
         
    }
    
    // Affiliate used by Order
    type Affiliate struct {

        
            Config AffiliateConfig  `json:"config"`
            Token string  `json:"token"`
            ID string  `json:"id"`
         
    }
    
    // OrderConfig used by Order
    type OrderConfig struct {

        
            BagEndState string  `json:"bag_end_state"`
            CreateUser bool  `json:"create_user"`
            ArticleLookup string  `json:"article_lookup"`
            StoreLookup string  `json:"store_lookup"`
            AffiliateStoreIDMapping []AffiliateStoreIdMapping  `json:"affiliate_store_id_mapping"`
            Affiliate Affiliate  `json:"affiliate"`
         
    }
    
    // OrderUser used by Order
    type OrderUser struct {

        
            Address1 string  `json:"address1"`
            Mobile float64  `json:"mobile"`
            City string  `json:"city"`
            LastName string  `json:"last_name"`
            FirstName string  `json:"first_name"`
            Country string  `json:"country"`
            Pincode string  `json:"pincode"`
            Address2 string  `json:"address2"`
            Email string  `json:"email"`
            State string  `json:"state"`
            Phone float64  `json:"phone"`
         
    }
    
    // ArticleDetails1 used by Order
    type ArticleDetails1 struct {

        
            BrandID float64  `json:"brand_id"`
            Dimension map[string]interface{}  `json:"dimension"`
            Quantity float64  `json:"quantity"`
            ID string  `json:"_id"`
            Category map[string]interface{}  `json:"category"`
            Weight map[string]interface{}  `json:"weight"`
            Attributes map[string]interface{}  `json:"attributes"`
         
    }
    
    // ShipmentDetails used by Order
    type ShipmentDetails struct {

        
            Meta map[string]interface{}  `json:"meta"`
            Shipments float64  `json:"shipments"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            BoxType string  `json:"box_type"`
            Articles []ArticleDetails1  `json:"articles"`
            DpID float64  `json:"dp_id"`
            FulfillmentID float64  `json:"fulfillment_id"`
         
    }
    
    // LocationDetails used by Order
    type LocationDetails struct {

        
            FulfillmentType string  `json:"fulfillment_type"`
            Articles []ArticleDetails1  `json:"articles"`
            FulfillmentID float64  `json:"fulfillment_id"`
         
    }
    
    // ShipmentConfig used by Order
    type ShipmentConfig struct {

        
            Action string  `json:"action"`
            Shipment []ShipmentDetails  `json:"shipment"`
            LocationDetails LocationDetails  `json:"location_details"`
            Journey string  `json:"journey"`
            Identifier string  `json:"identifier"`
            Source string  `json:"source"`
            ToPincode string  `json:"to_pincode"`
            PaymentMode string  `json:"payment_mode"`
         
    }
    
    // ShipmentData used by Order
    type ShipmentData struct {

        
            ShipmentData ShipmentConfig  `json:"shipment_data"`
         
    }
    
    // UserData used by Order
    type UserData struct {

        
            ShippingUser OrderUser  `json:"shipping_user"`
            BillingUser OrderUser  `json:"billing_user"`
         
    }
    
    // OrderPriority used by Order
    type OrderPriority struct {

        
            FulfilmentPriority float64  `json:"fulfilment_priority"`
            AffiliatePriorityCode string  `json:"affiliate_priority_code"`
            FulfilmentPriorityText string  `json:"fulfilment_priority_text"`
         
    }
    
    // MarketPlacePdf used by Order
    type MarketPlacePdf struct {

        
            Invoice string  `json:"invoice"`
            Label string  `json:"label"`
         
    }
    
    // AffiliateBag used by Order
    type AffiliateBag struct {

        
            UnitPrice float64  `json:"unit_price"`
            ItemSize string  `json:"item_size"`
            Discount float64  `json:"discount"`
            ID string  `json:"_id"`
            PriceEffective float64  `json:"price_effective"`
            Quantity float64  `json:"quantity"`
            AffiliateStoreID string  `json:"affiliate_store_id"`
            TransferPrice float64  `json:"transfer_price"`
            CompanyID float64  `json:"company_id"`
            PdfLinks MarketPlacePdf  `json:"pdf_links"`
            PriceMarked float64  `json:"price_marked"`
            ItemID float64  `json:"item_id"`
            AffiliateMeta map[string]interface{}  `json:"affiliate_meta"`
            SellerIdentifier string  `json:"seller_identifier"`
            HsnCodeID string  `json:"hsn_code_id"`
            Identifier map[string]interface{}  `json:"identifier"`
            Sku string  `json:"sku"`
            DeliveryCharge float64  `json:"delivery_charge"`
            StoreID float64  `json:"store_id"`
            AvlQty float64  `json:"avl_qty"`
            ModifiedOn string  `json:"modified_on"`
            FyndStoreID string  `json:"fynd_store_id"`
            AmountPaid float64  `json:"amount_paid"`
         
    }
    
    // OrderInfo used by Order
    type OrderInfo struct {

        
            BillingAddress OrderUser  `json:"billing_address"`
            Shipment ShipmentData  `json:"shipment"`
            ShippingAddress OrderUser  `json:"shipping_address"`
            Coupon string  `json:"coupon"`
            Discount float64  `json:"discount"`
            CodCharges float64  `json:"cod_charges"`
            Payment map[string]interface{}  `json:"payment"`
            User UserData  `json:"user"`
            DeliveryCharges float64  `json:"delivery_charges"`
            OrderPriority OrderPriority  `json:"order_priority"`
            OrderValue float64  `json:"order_value"`
            Bags []AffiliateBag  `json:"bags"`
            Items map[string]interface{}  `json:"items"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            PaymentMode string  `json:"payment_mode"`
         
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
            DisplayText string  `json:"display_text"`
            ID float64  `json:"id"`
         
    }
    
    // GetActionsResponse used by Order
    type GetActionsResponse struct {

        
            Permissions ActionInfo  `json:"permissions"`
         
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
    
    // HistoryDict used by Order
    type HistoryDict struct {

        
            Meta map[string]interface{}  `json:"meta"`
            Message string  `json:"message"`
            Createdat string  `json:"createdat"`
            AssignedAgent string  `json:"assigned_agent"`
            TicketURL string  `json:"ticket_url"`
            Type string  `json:"type"`
            L3Detail string  `json:"l3_detail"`
            DisplayMessage string  `json:"display_message"`
            User string  `json:"user"`
            L1Detail string  `json:"l1_detail"`
            L2Detail string  `json:"l2_detail"`
            BagID float64  `json:"bag_id"`
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
    
    // SmsDataPayload used by Order
    type SmsDataPayload struct {

        
            Message string  `json:"message"`
            CountryCode string  `json:"country_code"`
            ShipmentID float64  `json:"shipment_id"`
            PhoneNumber float64  `json:"phone_number"`
            CustomerName string  `json:"customer_name"`
            OrderID string  `json:"order_id"`
            AmountPaid float64  `json:"amount_paid"`
            BrandName string  `json:"brand_name"`
            PaymentMode string  `json:"payment_mode"`
         
    }
    
    // SendSmsPayload used by Order
    type SendSmsPayload struct {

        
            Slug string  `json:"slug"`
            BagID float64  `json:"bag_id"`
            Data SmsDataPayload  `json:"data"`
         
    }
    
    // OrderDetails used by Order
    type OrderDetails struct {

        
            FyndOrderID string  `json:"fynd_order_id"`
            CreatedAt string  `json:"created_at"`
         
    }
    
    // Meta1 used by Order
    type Meta1 struct {

        
            KafkaEmissionStatus float64  `json:"kafka_emission_status"`
            StateManagerUsed string  `json:"state_manager_used"`
         
    }
    
    // ShipmentDetail used by Order
    type ShipmentDetail struct {

        
            Meta Meta1  `json:"meta"`
            Status string  `json:"status"`
            ShipmentID string  `json:"shipment_id"`
            BagList []float64  `json:"bag_list"`
            Remarks string  `json:"remarks"`
            ID float64  `json:"id"`
         
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
    
    // ManualAssignDPToShipment used by Order
    type ManualAssignDPToShipment struct {

        
            OrderType string  `json:"order_type"`
            QcRequired string  `json:"qc_required"`
            ShipmentIds []string  `json:"shipment_ids"`
            DpID float64  `json:"dp_id"`
         
    }
    
    // ManualAssignDPToShipmentResponse used by Order
    type ManualAssignDPToShipmentResponse struct {

        
            Success string  `json:"success"`
            Errors []string  `json:"errors"`
         
    }
    
    // TaxInfo used by Order
    type TaxInfo struct {

        
            B2bGstinNumber string  `json:"b2b_gstin_number"`
            Gstin string  `json:"gstin"`
         
    }
    
    // PaymentMethod used by Order
    type PaymentMethod struct {

        
            Meta map[string]interface{}  `json:"meta"`
            Amount float64  `json:"amount"`
            RefundBy string  `json:"refund_by"`
            TransactionData map[string]interface{}  `json:"transaction_data"`
            Name string  `json:"name"`
            CollectBy string  `json:"collect_by"`
            Mode string  `json:"mode"`
         
    }
    
    // PaymentInfo used by Order
    type PaymentInfo struct {

        
            PrimaryMode string  `json:"primary_mode"`
            PaymentMethods []PaymentMethod  `json:"payment_methods"`
         
    }
    
    // BillingInfo used by Order
    type BillingInfo struct {

        
            Address1 string  `json:"address1"`
            LastName string  `json:"last_name"`
            Country string  `json:"country"`
            AlternateMobileNumber string  `json:"alternate_mobile_number"`
            HouseNo string  `json:"house_no"`
            FloorNo string  `json:"floor_no"`
            CustomerCode string  `json:"customer_code"`
            AlternateEmail string  `json:"alternate_email"`
            MiddleName string  `json:"middle_name"`
            Pincode string  `json:"pincode"`
            PrimaryMobileNumber string  `json:"primary_mobile_number"`
            State string  `json:"state"`
            Gender string  `json:"gender"`
            Title string  `json:"title"`
            PrimaryEmail string  `json:"primary_email"`
            StateCode string  `json:"state_code"`
            ExternalCustomerCode string  `json:"external_customer_code"`
            City string  `json:"city"`
            FirstName string  `json:"first_name"`
            Address2 string  `json:"address2"`
            CountryCode string  `json:"country_code"`
         
    }
    
    // ProcessingDates used by Order
    type ProcessingDates struct {

        
            DispatchAfterDate string  `json:"dispatch_after_date"`
            DpPickupSlot map[string]interface{}  `json:"dp_pickup_slot"`
            CustomerPickupSlot map[string]interface{}  `json:"customer_pickup_slot"`
            DispatchByDate string  `json:"dispatch_by_date"`
            PackByDate string  `json:"pack_by_date"`
            ConfirmByDate string  `json:"confirm_by_date"`
         
    }
    
    // Tax used by Order
    type Tax struct {

        
            Breakup []map[string]interface{}  `json:"breakup"`
            Rate float64  `json:"rate"`
            Name string  `json:"name"`
            Amount map[string]interface{}  `json:"amount"`
         
    }
    
    // Charge used by Order
    type Charge struct {

        
            Amount map[string]interface{}  `json:"amount"`
            Type string  `json:"type"`
            Name string  `json:"name"`
            Tax Tax  `json:"tax"`
            Code string  `json:"code"`
         
    }
    
    // LineItem used by Order
    type LineItem struct {

        
            ExternalLineID string  `json:"external_line_id"`
            Meta map[string]interface{}  `json:"meta"`
            CustomMessasge string  `json:"custom_messasge"`
            SellerIdentifier string  `json:"seller_identifier"`
            Quantity float64  `json:"quantity"`
            Charges []Charge  `json:"charges"`
         
    }
    
    // Shipment used by Order
    type Shipment struct {

        
            Meta map[string]interface{}  `json:"meta"`
            ExternalShipmentID string  `json:"external_shipment_id"`
            ProcessingDates ProcessingDates  `json:"processing_dates"`
            LineItems []LineItem  `json:"line_items"`
            Priority float64  `json:"priority"`
            LocationID float64  `json:"location_id"`
         
    }
    
    // ShippingInfo used by Order
    type ShippingInfo struct {

        
            Address1 string  `json:"address1"`
            LastName string  `json:"last_name"`
            Slot []map[string]interface{}  `json:"slot"`
            Country string  `json:"country"`
            ShippingType string  `json:"shipping_type"`
            AlternateMobileNumber string  `json:"alternate_mobile_number"`
            AddressType string  `json:"address_type"`
            HouseNo string  `json:"house_no"`
            FloorNo string  `json:"floor_no"`
            CustomerCode string  `json:"customer_code"`
            AlternateEmail string  `json:"alternate_email"`
            MiddleName string  `json:"middle_name"`
            Pincode string  `json:"pincode"`
            PrimaryMobileNumber string  `json:"primary_mobile_number"`
            State string  `json:"state"`
            Gender string  `json:"gender"`
            Title string  `json:"title"`
            PrimaryEmail string  `json:"primary_email"`
            StateCode string  `json:"state_code"`
            GeoLocation map[string]interface{}  `json:"geo_location"`
            ExternalCustomerCode string  `json:"external_customer_code"`
            City string  `json:"city"`
            FirstName string  `json:"first_name"`
            Address2 string  `json:"address2"`
            Landmark string  `json:"landmark"`
            CountryCode string  `json:"country_code"`
         
    }
    
    // CreateOrderAPI used by Order
    type CreateOrderAPI struct {

        
            Meta map[string]interface{}  `json:"meta"`
            TaxInfo TaxInfo  `json:"tax_info"`
            PaymentInfo PaymentInfo  `json:"payment_info"`
            ExternalOrderID string  `json:"external_order_id"`
            BillingInfo BillingInfo  `json:"billing_info"`
            ExternalCreationDate string  `json:"external_creation_date"`
            Shipments []Shipment  `json:"shipments"`
            Charges []Charge  `json:"charges"`
            ShippingInfo ShippingInfo  `json:"shipping_info"`
            CurrencyInfo map[string]interface{}  `json:"currency_info"`
            Config map[string]interface{}  `json:"config"`
         
    }
    
    // CreateOrderErrorReponse used by Order
    type CreateOrderErrorReponse struct {

        
            Meta string  `json:"meta"`
            Message string  `json:"message"`
            StackTrace string  `json:"stack_trace"`
            Status float64  `json:"status"`
            Info interface{}  `json:"info"`
            Code string  `json:"code"`
            Exception string  `json:"exception"`
            RequestID string  `json:"request_id"`
         
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
            PaymentMethods []PaymentMethods  `json:"payment_methods"`
            ModeOfPayment string  `json:"mode_of_payment"`
         
    }
    
    // CreateChannelConfig used by Order
    type CreateChannelConfig struct {

        
            DpConfiguration DpConfiguration  `json:"dp_configuration"`
            PaymentInfo CreateChannelPaymentInfo  `json:"payment_info"`
            LockStates []string  `json:"lock_states"`
            LogoURL map[string]interface{}  `json:"logo_url"`
            LocationReassignment bool  `json:"location_reassignment"`
            ShipmentAssignment string  `json:"shipment_assignment"`
         
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
            OrderDetails []FyndOrderIdList  `json:"order_details"`
            Mobile float64  `json:"mobile"`
            StartDate string  `json:"start_date"`
         
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

        
            OrderingChannel string  `json:"ordering_channel"`
            SellerID string  `json:"seller_id"`
            FyndOrderID string  `json:"fynd_order_id"`
            AffiliateID string  `json:"affiliate_id"`
            CustomerMobileNumber string  `json:"customer_mobile_number"`
         
    }
    
    // SingleRefundModeInfo used by Order
    type SingleRefundModeInfo struct {

        
            IsActive bool  `json:"is_active"`
            Slug string  `json:"slug"`
            DisplayName string  `json:"display_name"`
         
    }
    
    // RefundModeInfo used by Order
    type RefundModeInfo struct {

        
            RefundModeName SingleRefundModeInfo  `json:"refund_mode_name"`
         
    }
    
    // RefundModeConfigResponsePayload used by Order
    type RefundModeConfigResponsePayload struct {

        
            Success bool  `json:"success"`
            Data RefundModeInfo  `json:"data"`
         
    }
    
    // AttachUserInfo used by Order
    type AttachUserInfo struct {

        
            CountryCode string  `json:"country_code"`
            LastName string  `json:"last_name"`
            FirstName string  `json:"first_name"`
            Mobile float64  `json:"mobile"`
         
    }
    
    // AttachUserOtpData used by Order
    type AttachUserOtpData struct {

        
            RequestID string  `json:"request_id"`
         
    }
    
    // AttachOrderUser used by Order
    type AttachOrderUser struct {

        
            FyndOrderID string  `json:"fynd_order_id"`
            UserInfo AttachUserInfo  `json:"user_info"`
            OtpData AttachUserOtpData  `json:"otp_data"`
         
    }
    
    // AttachOrderUserResponse used by Order
    type AttachOrderUserResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // SendUserMobileOTP used by Order
    type SendUserMobileOTP struct {

        
            CountryCode string  `json:"country_code"`
            Mobile float64  `json:"mobile"`
         
    }
    
    // PointBlankOtpData used by Order
    type PointBlankOtpData struct {

        
            ResendTimer float64  `json:"resend_timer"`
            Message string  `json:"message"`
            Mobile float64  `json:"mobile"`
            RequestID string  `json:"request_id"`
         
    }
    
    // SendUserMobileOtpResponse used by Order
    type SendUserMobileOtpResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
            Data PointBlankOtpData  `json:"data"`
         
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
    

    
    // GetSearchWordsData used by Catalog
    type GetSearchWordsData struct {

        
            AppID string  `json:"app_id"`
            UID string  `json:"uid"`
            Result map[string]interface{}  `json:"result"`
            IsActive bool  `json:"is_active"`
            Words []string  `json:"words"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
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
    
    // ErrorResponse used by Catalog
    type ErrorResponse struct {

        
            Error string  `json:"error"`
            Status float64  `json:"status"`
            Message string  `json:"message"`
            Code string  `json:"code"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // DeleteResponse used by Catalog
    type DeleteResponse struct {

        
            Message string  `json:"message"`
         
    }
    
    // SearchKeywordResult used by Catalog
    type SearchKeywordResult struct {

        
            SortOn string  `json:"sort_on"`
            Query map[string]interface{}  `json:"query"`
         
    }
    
    // CreateSearchKeyword used by Catalog
    type CreateSearchKeyword struct {

        
            AppID string  `json:"app_id"`
            Result SearchKeywordResult  `json:"result"`
            IsActive bool  `json:"is_active"`
            Words []string  `json:"words"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // GetSearchWordsResponse used by Catalog
    type GetSearchWordsResponse struct {

        
            Items []GetSearchWordsData  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // GetAutocompleteWordsData used by Catalog
    type GetAutocompleteWordsData struct {

        
            Results []map[string]interface{}  `json:"results"`
            AppID string  `json:"app_id"`
            UID string  `json:"uid"`
            Words []string  `json:"words"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // GetAutocompleteWordsResponse used by Catalog
    type GetAutocompleteWordsResponse struct {

        
            Items []GetAutocompleteWordsData  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // Media used by Catalog
    type Media struct {

        
            Type string  `json:"type"`
            AspectRatio string  `json:"aspect_ratio"`
            URL string  `json:"url"`
         
    }
    
    // AutocompletePageAction used by Catalog
    type AutocompletePageAction struct {

        
            Type string  `json:"type"`
            Params map[string]interface{}  `json:"params"`
            URL string  `json:"url"`
            Query map[string]interface{}  `json:"query"`
         
    }
    
    // AutocompleteAction used by Catalog
    type AutocompleteAction struct {

        
            Type string  `json:"type"`
            Page AutocompletePageAction  `json:"page"`
         
    }
    
    // AutocompleteResult used by Catalog
    type AutocompleteResult struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Display string  `json:"display"`
            Logo Media  `json:"logo"`
            Action AutocompleteAction  `json:"action"`
         
    }
    
    // CreateAutocompleteKeyword used by Catalog
    type CreateAutocompleteKeyword struct {

        
            Results []AutocompleteResult  `json:"results"`
            AppID string  `json:"app_id"`
            IsActive bool  `json:"is_active"`
            Words []string  `json:"words"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // CreateAutocompleteWordsResponse used by Catalog
    type CreateAutocompleteWordsResponse struct {

        
            Words []string  `json:"words"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Results []map[string]interface{}  `json:"results"`
            AppID string  `json:"app_id"`
         
    }
    
    // ProductBundleItem used by Catalog
    type ProductBundleItem struct {

        
            MaxQuantity float64  `json:"max_quantity"`
            AllowRemove bool  `json:"allow_remove"`
            MinQuantity float64  `json:"min_quantity"`
            AutoSelect bool  `json:"auto_select"`
            ProductUID float64  `json:"product_uid"`
            AutoAddToCart bool  `json:"auto_add_to_cart"`
         
    }
    
    // ProductBundleRequest used by Catalog
    type ProductBundleRequest struct {

        
            Slug string  `json:"slug"`
            Choice string  `json:"choice"`
            CreatedOn string  `json:"created_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Logo string  `json:"logo"`
            IsActive bool  `json:"is_active"`
            Meta map[string]interface{}  `json:"meta"`
            PageVisibility []string  `json:"page_visibility"`
            Products []ProductBundleItem  `json:"products"`
            Name string  `json:"name"`
            ModifiedOn string  `json:"modified_on"`
            CompanyID float64  `json:"company_id"`
         
    }
    
    // GetProductBundleCreateResponse used by Catalog
    type GetProductBundleCreateResponse struct {

        
            Slug string  `json:"slug"`
            Choice string  `json:"choice"`
            CreatedOn string  `json:"created_on"`
            ID string  `json:"id"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Logo string  `json:"logo"`
            IsActive bool  `json:"is_active"`
            Meta map[string]interface{}  `json:"meta"`
            PageVisibility []string  `json:"page_visibility"`
            Products []ProductBundleItem  `json:"products"`
            Name string  `json:"name"`
            ModifiedOn string  `json:"modified_on"`
            CompanyID float64  `json:"company_id"`
         
    }
    
    // GetProductBundleListingResponse used by Catalog
    type GetProductBundleListingResponse struct {

        
            Items []GetProductBundleCreateResponse  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // LimitedProductData used by Catalog
    type LimitedProductData struct {

        
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
            Price map[string]interface{}  `json:"price"`
            Sizes []string  `json:"sizes"`
            Images []string  `json:"images"`
            ItemCode string  `json:"item_code"`
            Quantity float64  `json:"quantity"`
            Attributes map[string]interface{}  `json:"attributes"`
            Identifier map[string]interface{}  `json:"identifier"`
            Name string  `json:"name"`
            CountryOfOrigin string  `json:"country_of_origin"`
            ShortDescription string  `json:"short_description"`
         
    }
    
    // Price used by Catalog
    type Price struct {

        
            MaxEffective float64  `json:"max_effective"`
            MaxMarked float64  `json:"max_marked"`
            MinMarked float64  `json:"min_marked"`
            MinEffective float64  `json:"min_effective"`
            Currency string  `json:"currency"`
         
    }
    
    // Size used by Catalog
    type Size struct {

        
            IsAvailable bool  `json:"is_available"`
            Quantity float64  `json:"quantity"`
            Display string  `json:"display"`
            Value string  `json:"value"`
         
    }
    
    // GetProducts used by Catalog
    type GetProducts struct {

        
            ProductDetails LimitedProductData  `json:"product_details"`
            MaxQuantity float64  `json:"max_quantity"`
            Price Price  `json:"price"`
            AllowRemove bool  `json:"allow_remove"`
            Sizes []Size  `json:"sizes"`
            MinQuantity float64  `json:"min_quantity"`
            AutoSelect bool  `json:"auto_select"`
            ProductUID float64  `json:"product_uid"`
            AutoAddToCart bool  `json:"auto_add_to_cart"`
         
    }
    
    // GetProductBundleResponse used by Catalog
    type GetProductBundleResponse struct {

        
            Slug string  `json:"slug"`
            Choice string  `json:"choice"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            Logo string  `json:"logo"`
            IsActive bool  `json:"is_active"`
            Meta map[string]interface{}  `json:"meta"`
            PageVisibility []string  `json:"page_visibility"`
            Products []GetProducts  `json:"products"`
            Name string  `json:"name"`
            CompanyID float64  `json:"company_id"`
         
    }
    
    // ProductBundleUpdateRequest used by Catalog
    type ProductBundleUpdateRequest struct {

        
            Slug string  `json:"slug"`
            Choice string  `json:"choice"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            Logo string  `json:"logo"`
            IsActive bool  `json:"is_active"`
            Meta map[string]interface{}  `json:"meta"`
            PageVisibility []string  `json:"page_visibility"`
            Products []ProductBundleItem  `json:"products"`
            Name string  `json:"name"`
            ModifiedOn string  `json:"modified_on"`
            CompanyID float64  `json:"company_id"`
         
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

        
            Subtitle string  `json:"subtitle"`
            Description string  `json:"description"`
            Guide Guide  `json:"guide"`
            CreatedOn string  `json:"created_on"`
            ID string  `json:"id"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Active bool  `json:"active"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Image string  `json:"image"`
            Tag string  `json:"tag"`
            BrandID float64  `json:"brand_id"`
            Name string  `json:"name"`
            ModifiedOn string  `json:"modified_on"`
            CompanyID float64  `json:"company_id"`
            Title string  `json:"title"`
         
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

        
            Subtitle string  `json:"subtitle"`
            CreatedOn string  `json:"created_on"`
            Guide map[string]interface{}  `json:"guide"`
            ID string  `json:"id"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Active bool  `json:"active"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Tag string  `json:"tag"`
            BrandID float64  `json:"brand_id"`
            Name string  `json:"name"`
            ModifiedOn string  `json:"modified_on"`
            CompanyID float64  `json:"company_id"`
            Title string  `json:"title"`
         
    }
    
    // ApplicationItemMOQ used by Catalog
    type ApplicationItemMOQ struct {

        
            Maximum float64  `json:"maximum"`
            Minimum float64  `json:"minimum"`
            IncrementUnit float64  `json:"increment_unit"`
         
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

        
            Moq ApplicationItemMOQ  `json:"moq"`
            CustomMeta []MetaFields  `json:"_custom_meta"`
            Seo ApplicationItemSEO  `json:"seo"`
            IsCod bool  `json:"is_cod"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            AltText map[string]interface{}  `json:"alt_text"`
            IsGift bool  `json:"is_gift"`
         
    }
    
    // SuccessResponse1 used by Catalog
    type SuccessResponse1 struct {

        
            UID float64  `json:"uid"`
            Success bool  `json:"success"`
         
    }
    
    // MOQData used by Catalog
    type MOQData struct {

        
            Maximum float64  `json:"maximum"`
            Minimum float64  `json:"minimum"`
            IncrementUnit float64  `json:"increment_unit"`
         
    }
    
    // SEOData used by Catalog
    type SEOData struct {

        
            Description interface{}  `json:"description"`
            Title interface{}  `json:"title"`
         
    }
    
    // OwnerAppItemResponse used by Catalog
    type OwnerAppItemResponse struct {

        
            Moq MOQData  `json:"moq"`
            Seo SEOData  `json:"seo"`
            IsCod bool  `json:"is_cod"`
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

        
            Slug string  `json:"slug"`
            Key string  `json:"key"`
            Logo string  `json:"logo"`
            DisplayType string  `json:"display_type"`
            Unit string  `json:"unit"`
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            Priority float64  `json:"priority"`
         
    }
    
    // AppConfigurationDetail used by Catalog
    type AppConfigurationDetail struct {

        
            Slug string  `json:"slug"`
            AppID string  `json:"app_id"`
            Logo string  `json:"logo"`
            IsDefault bool  `json:"is_default"`
            TemplateSlugs []string  `json:"template_slugs"`
            IsActive bool  `json:"is_active"`
            Attributes []AttributeDetailsGroup  `json:"attributes"`
            Name string  `json:"name"`
            Priority float64  `json:"priority"`
         
    }
    
    // ConfigErrorResponse used by Catalog
    type ConfigErrorResponse struct {

        
            Message string  `json:"message"`
         
    }
    
    // PageResponseType used by Catalog
    type PageResponseType struct {

        
            TotalCount float64  `json:"total_count"`
            Next float64  `json:"next"`
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
            Key string  `json:"key"`
            Logo string  `json:"logo"`
            IsDefault bool  `json:"is_default"`
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
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
    
    // GetCatalogConfigurationDetailsProduct used by Catalog
    type GetCatalogConfigurationDetailsProduct struct {

        
            Detail map[string]interface{}  `json:"detail"`
            Variant map[string]interface{}  `json:"variant"`
            Compare map[string]interface{}  `json:"compare"`
            Similar map[string]interface{}  `json:"similar"`
         
    }
    
    // MetaDataListingFilterMetaResponse used by Catalog
    type MetaDataListingFilterMetaResponse struct {

        
            Key string  `json:"key"`
            Units []map[string]interface{}  `json:"units"`
            Display string  `json:"display"`
            FilterTypes []string  `json:"filter_types"`
         
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
    
    // ConfigurationBucketPoints used by Catalog
    type ConfigurationBucketPoints struct {

        
            Start float64  `json:"start"`
            End float64  `json:"end"`
            Display string  `json:"display"`
         
    }
    
    // ConfigurationListingFilterValue used by Catalog
    type ConfigurationListingFilterValue struct {

        
            Map map[string]interface{}  `json:"map"`
            Condition string  `json:"condition"`
            Sort string  `json:"sort"`
            Value string  `json:"value"`
            BucketPoints []ConfigurationBucketPoints  `json:"bucket_points"`
            MapValues []map[string]interface{}  `json:"map_values"`
         
    }
    
    // ConfigurationListingFilterConfig used by Catalog
    type ConfigurationListingFilterConfig struct {

        
            DisplayName string  `json:"display_name"`
            Key string  `json:"key"`
            Logo string  `json:"logo"`
            Type string  `json:"type"`
            IsActive bool  `json:"is_active"`
            ValueConfig ConfigurationListingFilterValue  `json:"value_config"`
            Name string  `json:"name"`
            Priority float64  `json:"priority"`
         
    }
    
    // ConfigurationListingFilter used by Catalog
    type ConfigurationListingFilter struct {

        
            AttributeConfig []ConfigurationListingFilterConfig  `json:"attribute_config"`
            AllowSingle bool  `json:"allow_single"`
         
    }
    
    // ConfigurationListingSortConfig used by Catalog
    type ConfigurationListingSortConfig struct {

        
            Key string  `json:"key"`
            Logo string  `json:"logo"`
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            Priority float64  `json:"priority"`
         
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
    
    // ProductSize used by Catalog
    type ProductSize struct {

        
            Max float64  `json:"max"`
            Min float64  `json:"min"`
         
    }
    
    // ConfigurationProductVariantConfig used by Catalog
    type ConfigurationProductVariantConfig struct {

        
            Key string  `json:"key"`
            Logo string  `json:"logo"`
            DisplayType string  `json:"display_type"`
            IsActive bool  `json:"is_active"`
            Size ProductSize  `json:"size"`
            Name string  `json:"name"`
            Priority float64  `json:"priority"`
         
    }
    
    // ConfigurationProductVariant used by Catalog
    type ConfigurationProductVariant struct {

        
            Config []ConfigurationProductVariantConfig  `json:"config"`
         
    }
    
    // ConfigurationProductConfig used by Catalog
    type ConfigurationProductConfig struct {

        
            Subtitle string  `json:"subtitle"`
            Key string  `json:"key"`
            Logo string  `json:"logo"`
            IsActive bool  `json:"is_active"`
            Size ProductSize  `json:"size"`
            Priority float64  `json:"priority"`
            Title string  `json:"title"`
         
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
    
    // AppConfiguration used by Catalog
    type AppConfiguration struct {

        
            CreatedOn string  `json:"created_on"`
            Listing ConfigurationListing  `json:"listing"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            AppID string  `json:"app_id"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Type string  `json:"type"`
            ConfigID string  `json:"config_id"`
            Product ConfigurationProduct  `json:"product"`
            ModifiedOn string  `json:"modified_on"`
            ConfigType string  `json:"config_type"`
         
    }
    
    // AppCatalogConfiguration used by Catalog
    type AppCatalogConfiguration struct {

        
            CreatedOn string  `json:"created_on"`
            Listing ConfigurationListing  `json:"listing"`
            ID string  `json:"id"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            AppID string  `json:"app_id"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Type string  `json:"type"`
            ConfigID string  `json:"config_id"`
            Product ConfigurationProduct  `json:"product"`
            ModifiedOn string  `json:"modified_on"`
            ConfigType string  `json:"config_type"`
         
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

        
            Listing GetCatalogConfigurationDetailsSchemaListing  `json:"listing"`
            ID string  `json:"id"`
            AppID string  `json:"app_id"`
            ConfigID string  `json:"config_id"`
            Product GetCatalogConfigurationDetailsProduct  `json:"product"`
            ConfigType string  `json:"config_type"`
         
    }
    
    // GetAppCatalogEntityConfiguration used by Catalog
    type GetAppCatalogEntityConfiguration struct {

        
            Data EntityConfiguration  `json:"data"`
            IsDefault bool  `json:"is_default"`
         
    }
    
    // ProductFiltersKey used by Catalog
    type ProductFiltersKey struct {

        
            Logo string  `json:"logo"`
            Kind string  `json:"kind"`
            Name string  `json:"name"`
            Display string  `json:"display"`
            Operators []string  `json:"operators"`
         
    }
    
    // ProductFiltersValue used by Catalog
    type ProductFiltersValue struct {

        
            DisplayFormat string  `json:"display_format"`
            IsSelected bool  `json:"is_selected"`
            Max float64  `json:"max"`
            Count float64  `json:"count"`
            Min float64  `json:"min"`
            QueryFormat string  `json:"query_format"`
            Value interface{}  `json:"value"`
            SelectedMin float64  `json:"selected_min"`
            CurrencyCode string  `json:"currency_code"`
            SelectedMax float64  `json:"selected_max"`
            Display string  `json:"display"`
            CurrencySymbol string  `json:"currency_symbol"`
         
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
    
    // SeoDetail used by Catalog
    type SeoDetail struct {

        
            Description string  `json:"description"`
            Title string  `json:"title"`
         
    }
    
    // CollectionImage used by Catalog
    type CollectionImage struct {

        
            AspectRatio string  `json:"aspect_ratio"`
            URL string  `json:"url"`
         
    }
    
    // UserInfo used by Catalog
    type UserInfo struct {

        
            UID string  `json:"uid"`
            Email string  `json:"email"`
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
         
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
    
    // CollectionBanner used by Catalog
    type CollectionBanner struct {

        
            Portrait CollectionImage  `json:"portrait"`
            Landscape CollectionImage  `json:"landscape"`
         
    }
    
    // NextSchedule used by Catalog
    type NextSchedule struct {

        
            Start string  `json:"start"`
            End string  `json:"end"`
         
    }
    
    // CollectionSchedule used by Catalog
    type CollectionSchedule struct {

        
            End string  `json:"end"`
            NextSchedule []NextSchedule  `json:"next_schedule"`
            Start string  `json:"start"`
            Cron string  `json:"cron"`
            Duration float64  `json:"duration"`
         
    }
    
    // CreateCollection used by Catalog
    type CreateCollection struct {

        
            Description string  `json:"description"`
            Seo SeoDetail  `json:"seo"`
            Logo CollectionImage  `json:"logo"`
            Type string  `json:"type"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Published bool  `json:"published"`
            ModifiedBy UserInfo  `json:"modified_by"`
            IsActive bool  `json:"is_active"`
            Meta map[string]interface{}  `json:"meta"`
            AllowSort bool  `json:"allow_sort"`
            Badge CollectionBadge  `json:"badge"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            Tags []string  `json:"tags"`
            AllowFacets bool  `json:"allow_facets"`
            Query []CollectionQuery  `json:"query"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Banners CollectionBanner  `json:"banners"`
            IsVisible bool  `json:"is_visible"`
            AppID string  `json:"app_id"`
            CreatedBy UserInfo  `json:"created_by"`
            Schedule CollectionSchedule  `json:"_schedule"`
            Priority float64  `json:"priority"`
            SortOn string  `json:"sort_on"`
         
    }
    
    // BannerImage used by Catalog
    type BannerImage struct {

        
            Type string  `json:"type"`
            AspectRatio string  `json:"aspect_ratio"`
            URL string  `json:"url"`
         
    }
    
    // ImageUrls used by Catalog
    type ImageUrls struct {

        
            Portrait BannerImage  `json:"portrait"`
            Landscape BannerImage  `json:"landscape"`
         
    }
    
    // CollectionCreateResponse used by Catalog
    type CollectionCreateResponse struct {

        
            Description string  `json:"description"`
            Logo BannerImage  `json:"logo"`
            Type string  `json:"type"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            IsActive bool  `json:"is_active"`
            Cron map[string]interface{}  `json:"cron"`
            Meta map[string]interface{}  `json:"meta"`
            AllowSort bool  `json:"allow_sort"`
            Badge map[string]interface{}  `json:"badge"`
            AllowFacets bool  `json:"allow_facets"`
            Query []CollectionQuery  `json:"query"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Tag []string  `json:"tag"`
            Banners ImageUrls  `json:"banners"`
            AppID string  `json:"app_id"`
            Schedule map[string]interface{}  `json:"_schedule"`
            Priority float64  `json:"priority"`
            SortOn string  `json:"sort_on"`
         
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
    
    // Media1 used by Catalog
    type Media1 struct {

        
            Type string  `json:"type"`
            Meta map[string]interface{}  `json:"meta"`
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

        
            Description string  `json:"description"`
            Logo Media1  `json:"logo"`
            Type string  `json:"type"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            IsActive bool  `json:"is_active"`
            Cron map[string]interface{}  `json:"cron"`
            Meta map[string]interface{}  `json:"meta"`
            AllowSort bool  `json:"allow_sort"`
            Badge map[string]interface{}  `json:"badge"`
            AllowFacets bool  `json:"allow_facets"`
            Query []CollectionQuery  `json:"query"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Tag []string  `json:"tag"`
            Action Action  `json:"action"`
            Banners ImageUrls  `json:"banners"`
            AppID string  `json:"app_id"`
            UID string  `json:"uid"`
            Schedule map[string]interface{}  `json:"_schedule"`
            Priority float64  `json:"priority"`
         
    }
    
    // GetCollectionListingResponse used by Catalog
    type GetCollectionListingResponse struct {

        
            Filters CollectionListingFilter  `json:"filters"`
            Items []GetCollectionDetailNest  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // CollectionDetailResponse used by Catalog
    type CollectionDetailResponse struct {

        
            Description string  `json:"description"`
            Logo Media1  `json:"logo"`
            Type string  `json:"type"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            IsActive bool  `json:"is_active"`
            Cron map[string]interface{}  `json:"cron"`
            Meta map[string]interface{}  `json:"meta"`
            AllowSort bool  `json:"allow_sort"`
            Badge map[string]interface{}  `json:"badge"`
            AllowFacets bool  `json:"allow_facets"`
            Query []CollectionQuery  `json:"query"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Tag []string  `json:"tag"`
            Banners ImageUrls  `json:"banners"`
            AppID string  `json:"app_id"`
            UID string  `json:"uid"`
            Schedule map[string]interface{}  `json:"_schedule"`
            Priority float64  `json:"priority"`
         
    }
    
    // UpdateCollection used by Catalog
    type UpdateCollection struct {

        
            Description string  `json:"description"`
            Seo SeoDetail  `json:"seo"`
            Logo CollectionImage  `json:"logo"`
            Type string  `json:"type"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Published bool  `json:"published"`
            ModifiedBy UserInfo  `json:"modified_by"`
            IsActive bool  `json:"is_active"`
            Meta map[string]interface{}  `json:"meta"`
            AllowSort bool  `json:"allow_sort"`
            Badge CollectionBadge  `json:"badge"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            Tags []string  `json:"tags"`
            AllowFacets bool  `json:"allow_facets"`
            Query []CollectionQuery  `json:"query"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Banners CollectionBanner  `json:"banners"`
            IsVisible bool  `json:"is_visible"`
            Schedule CollectionSchedule  `json:"_schedule"`
            Priority float64  `json:"priority"`
            SortOn string  `json:"sort_on"`
         
    }
    
    // CollectionItem used by Catalog
    type CollectionItem struct {

        
            Priority float64  `json:"priority"`
            ItemID float64  `json:"item_id"`
            Action string  `json:"action"`
         
    }
    
    // CollectionItemUpdate used by Catalog
    type CollectionItemUpdate struct {

        
            Items []CollectionItem  `json:"items"`
            AllowFacets bool  `json:"allow_facets"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Type string  `json:"type"`
            Query []CollectionQuery  `json:"query"`
            AllowSort bool  `json:"allow_sort"`
         
    }
    
    // UpdatedResponse used by Catalog
    type UpdatedResponse struct {

        
            ItemsNotUpdated []float64  `json:"items_not_updated"`
            Message string  `json:"message"`
         
    }
    
    // ProductBrand used by Catalog
    type ProductBrand struct {

        
            Name string  `json:"name"`
            Logo Media1  `json:"logo"`
            UID float64  `json:"uid"`
            Action Action  `json:"action"`
         
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

        
            Effective Price1  `json:"effective"`
            Marked Price1  `json:"marked"`
         
    }
    
    // ProductDetailAttribute used by Catalog
    type ProductDetailAttribute struct {

        
            Type string  `json:"type"`
            Key string  `json:"key"`
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
            Description string  `json:"description"`
            PromoMeta map[string]interface{}  `json:"promo_meta"`
            HasVariant bool  `json:"has_variant"`
            Type string  `json:"type"`
            Sellable bool  `json:"sellable"`
            Name string  `json:"name"`
            Color string  `json:"color"`
            ShortDescription string  `json:"short_description"`
            Slug string  `json:"slug"`
            RatingCount float64  `json:"rating_count"`
            ImageNature string  `json:"image_nature"`
            Similars []string  `json:"similars"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            Medias []Media1  `json:"medias"`
            Discount string  `json:"discount"`
            ProductOnlineDate string  `json:"product_online_date"`
            Price ProductListingPrice  `json:"price"`
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
            Highlights []string  `json:"highlights"`
            Attributes map[string]interface{}  `json:"attributes"`
            ItemType string  `json:"item_type"`
            UID float64  `json:"uid"`
            ItemCode string  `json:"item_code"`
            Rating float64  `json:"rating"`
            Tryouts []string  `json:"tryouts"`
         
    }
    
    // GetCollectionItemsResponse used by Catalog
    type GetCollectionItemsResponse struct {

        
            Filters []ProductFilters  `json:"filters"`
            Items []ProductListingDetail  `json:"items"`
            Page Page  `json:"page"`
            SortOn []ProductSortOn  `json:"sort_on"`
         
    }
    
    // CatalogInsightItem used by Catalog
    type CatalogInsightItem struct {

        
            SellableCount float64  `json:"sellable_count"`
            Count float64  `json:"count"`
            OutOfStockCount float64  `json:"out_of_stock_count"`
         
    }
    
    // CatalogInsightBrand used by Catalog
    type CatalogInsightBrand struct {

        
            AvailableSizes float64  `json:"available_sizes"`
            TotalSizes float64  `json:"total_sizes"`
            ArticleFreshness float64  `json:"article_freshness"`
            TotalArticles float64  `json:"total_articles"`
            AvailableArticles float64  `json:"available_articles"`
            Name string  `json:"name"`
         
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

        
            Data CrossSellingData  `json:"data"`
            BrandDistribution CatalogInsightBrand  `json:"brand_distribution"`
         
    }
    
    // OptInPostRequest used by Catalog
    type OptInPostRequest struct {

        
            Enabled bool  `json:"enabled"`
            OptLevel string  `json:"opt_level"`
            Platform string  `json:"platform"`
            StoreIds []float64  `json:"store_ids"`
            CompanyID float64  `json:"company_id"`
            BrandIds []float64  `json:"brand_ids"`
         
    }
    
    // CompanyOptIn used by Catalog
    type CompanyOptIn struct {

        
            Enabled bool  `json:"enabled"`
            CreatedOn float64  `json:"created_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            OptLevel string  `json:"opt_level"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Platform string  `json:"platform"`
            StoreIds []float64  `json:"store_ids"`
            ModifiedOn float64  `json:"modified_on"`
            CompanyID float64  `json:"company_id"`
            BrandIds []float64  `json:"brand_ids"`
         
    }
    
    // GetOptInPlatform used by Catalog
    type GetOptInPlatform struct {

        
            Items []CompanyOptIn  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // OptinCompanyDetail used by Catalog
    type OptinCompanyDetail struct {

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            BusinessType string  `json:"business_type"`
            CompanyType string  `json:"company_type"`
         
    }
    
    // CompanyBrandDetail used by Catalog
    type CompanyBrandDetail struct {

        
            BrandName string  `json:"brand_name"`
            TotalArticle float64  `json:"total_article"`
            CompanyID float64  `json:"company_id"`
            BrandID float64  `json:"brand_id"`
         
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

        
            DisplayName string  `json:"display_name"`
            CreatedOn string  `json:"created_on"`
            Manager map[string]interface{}  `json:"manager"`
            AdditionalContacts []map[string]interface{}  `json:"additional_contacts"`
            UID float64  `json:"uid"`
            Timing map[string]interface{}  `json:"timing"`
            Documents []map[string]interface{}  `json:"documents"`
            StoreType string  `json:"store_type"`
            Name string  `json:"name"`
            ModifiedOn string  `json:"modified_on"`
            CompanyID float64  `json:"company_id"`
            Address map[string]interface{}  `json:"address"`
            StoreCode string  `json:"store_code"`
         
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

        
            Range AttributeSchemaRange  `json:"range"`
            Format string  `json:"format"`
            AllowedValues []string  `json:"allowed_values"`
            Multi bool  `json:"multi"`
            Type string  `json:"type"`
            Mandatory bool  `json:"mandatory"`
         
    }
    
    // AttributeMasterDetails used by Catalog
    type AttributeMasterDetails struct {

        
            DisplayType string  `json:"display_type"`
         
    }
    
    // AttributeMasterFilter used by Catalog
    type AttributeMasterFilter struct {

        
            Indexing bool  `json:"indexing"`
            Priority float64  `json:"priority"`
            DependsOn []string  `json:"depends_on"`
         
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

        
            Slug string  `json:"slug"`
            Description string  `json:"description"`
            Schema AttributeMaster  `json:"schema"`
            Details AttributeMasterDetails  `json:"details"`
            IsNested bool  `json:"is_nested"`
            Departments []string  `json:"departments"`
            ID string  `json:"id"`
            Filters AttributeMasterFilter  `json:"filters"`
            Logo string  `json:"logo"`
            Meta AttributeMasterMeta  `json:"meta"`
            EnabledForEndConsumer bool  `json:"enabled_for_end_consumer"`
            Name string  `json:"name"`
         
    }
    
    // CategoriesResponse used by Catalog
    type CategoriesResponse struct {

        
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
            TemplateSlug string  `json:"template_slug"`
            Name string  `json:"name"`
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
            Message string  `json:"message"`
            Errors map[string]interface{}  `json:"errors"`
            Code string  `json:"code"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // DepartmentCreateUpdate used by Catalog
    type DepartmentCreateUpdate struct {

        
            Slug string  `json:"slug"`
            PriorityOrder float64  `json:"priority_order"`
            UID float64  `json:"uid"`
            Logo string  `json:"logo"`
            IsActive bool  `json:"is_active"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Platforms map[string]interface{}  `json:"platforms"`
            Synonyms []string  `json:"synonyms"`
            Name string  `json:"name"`
            Cls string  `json:"_cls"`
            Tags []string  `json:"tags"`
         
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

        
            UID string  `json:"uid"`
            UserID string  `json:"user_id"`
            Contact string  `json:"contact"`
            Username string  `json:"username"`
            ID string  `json:"_id"`
         
    }
    
    // GetDepartment used by Catalog
    type GetDepartment struct {

        
            Slug string  `json:"slug"`
            CreatedOn string  `json:"created_on"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            PriorityOrder float64  `json:"priority_order"`
            PageNo float64  `json:"page_no"`
            UID float64  `json:"uid"`
            PageSize float64  `json:"page_size"`
            CreatedBy UserSerializer  `json:"created_by"`
            Logo string  `json:"logo"`
            IsActive bool  `json:"is_active"`
            ItemType string  `json:"item_type"`
            Synonyms []string  `json:"synonyms"`
            Name string  `json:"name"`
            ModifiedOn string  `json:"modified_on"`
            Search string  `json:"search"`
         
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
            Code string  `json:"code"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // UserDetail used by Catalog
    type UserDetail struct {

        
            UserID string  `json:"user_id"`
            SuperUser bool  `json:"super_user"`
            Username string  `json:"username"`
            Contact string  `json:"contact"`
         
    }
    
    // DepartmentModel used by Catalog
    type DepartmentModel struct {

        
            Slug interface{}  `json:"slug"`
            CreatedOn string  `json:"created_on"`
            Cls interface{}  `json:"_cls"`
            ModifiedBy UserDetail  `json:"modified_by"`
            PriorityOrder float64  `json:"priority_order"`
            UID float64  `json:"uid"`
            CreatedBy UserDetail  `json:"created_by"`
            Logo string  `json:"logo"`
            Synonyms []interface{}  `json:"synonyms"`
            IsActive bool  `json:"is_active"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ID interface{}  `json:"_id"`
            Name interface{}  `json:"name"`
            ModifiedOn string  `json:"modified_on"`
            VerifiedBy UserDetail  `json:"verified_by"`
            VerifiedOn string  `json:"verified_on"`
         
    }
    
    // ProductTemplate used by Catalog
    type ProductTemplate struct {

        
            Slug string  `json:"slug"`
            Description string  `json:"description"`
            CreatedOn string  `json:"created_on"`
            Categories []string  `json:"categories"`
            Departments []string  `json:"departments"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            IsExpirable bool  `json:"is_expirable"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            IsArchived bool  `json:"is_archived"`
            Logo string  `json:"logo"`
            IsPhysical bool  `json:"is_physical"`
            Tag string  `json:"tag"`
            IsActive bool  `json:"is_active"`
            Attributes []string  `json:"attributes"`
            Name string  `json:"name"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // TemplatesResponse used by Catalog
    type TemplatesResponse struct {

        
            Items ProductTemplate  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // Properties used by Catalog
    type Properties struct {

        
            TraderType map[string]interface{}  `json:"trader_type"`
            Description map[string]interface{}  `json:"description"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            ProductGroupTag map[string]interface{}  `json:"product_group_tag"`
            HsnCode map[string]interface{}  `json:"hsn_code"`
            BrandUID map[string]interface{}  `json:"brand_uid"`
            Name map[string]interface{}  `json:"name"`
            ShortDescription map[string]interface{}  `json:"short_description"`
            Slug map[string]interface{}  `json:"slug"`
            Command map[string]interface{}  `json:"command"`
            IsActive map[string]interface{}  `json:"is_active"`
            CountryOfOrigin map[string]interface{}  `json:"country_of_origin"`
            CategorySlug map[string]interface{}  `json:"category_slug"`
            IsDependent map[string]interface{}  `json:"is_dependent"`
            Tags map[string]interface{}  `json:"tags"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            MultiSize map[string]interface{}  `json:"multi_size"`
            Sizes map[string]interface{}  `json:"sizes"`
            ProductPublish map[string]interface{}  `json:"product_publish"`
            Highlights map[string]interface{}  `json:"highlights"`
            ItemType map[string]interface{}  `json:"item_type"`
            Trader map[string]interface{}  `json:"trader"`
            Media map[string]interface{}  `json:"media"`
            Currency map[string]interface{}  `json:"currency"`
            NoOfBoxes map[string]interface{}  `json:"no_of_boxes"`
            ReturnConfig map[string]interface{}  `json:"return_config"`
            SizeGuide map[string]interface{}  `json:"size_guide"`
            ItemCode map[string]interface{}  `json:"item_code"`
            Variants map[string]interface{}  `json:"variants"`
         
    }
    
    // GlobalValidation used by Catalog
    type GlobalValidation struct {

        
            Description string  `json:"description"`
            Properties Properties  `json:"properties"`
            Required []string  `json:"required"`
            Definitions map[string]interface{}  `json:"definitions"`
            Type string  `json:"type"`
            Title string  `json:"title"`
         
    }
    
    // TemplateValidationData used by Catalog
    type TemplateValidationData struct {

        
            TemplateValidation map[string]interface{}  `json:"template_validation"`
            GlobalValidation GlobalValidation  `json:"global_validation"`
         
    }
    
    // TemplateDetails used by Catalog
    type TemplateDetails struct {

        
            Slug string  `json:"slug"`
            Description string  `json:"description"`
            ID string  `json:"id"`
            Categories []string  `json:"categories"`
            Departments []string  `json:"departments"`
            IsExpirable bool  `json:"is_expirable"`
            IsArchived bool  `json:"is_archived"`
            Logo string  `json:"logo"`
            IsPhysical bool  `json:"is_physical"`
            Tag string  `json:"tag"`
            IsActive bool  `json:"is_active"`
            Attributes []string  `json:"attributes"`
            Name string  `json:"name"`
         
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

        
            UID string  `json:"uid"`
            Email string  `json:"email"`
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
         
    }
    
    // ProductTemplateExportResponse used by Catalog
    type ProductTemplateExportResponse struct {

        
            TaskID string  `json:"task_id"`
            Filters map[string]interface{}  `json:"filters"`
            Status string  `json:"status"`
            NotificationEmails []string  `json:"notification_emails"`
            CreatedBy UserInfo1  `json:"created_by"`
            SellerID float64  `json:"seller_id"`
            URL string  `json:"url"`
            Type string  `json:"type"`
            CompletedOn string  `json:"completed_on"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // ProductDownloadsResponse used by Catalog
    type ProductDownloadsResponse struct {

        
            Items []ProductTemplateExportResponse  `json:"items"`
         
    }
    
    // ProductTemplateExportFilterRequest used by Catalog
    type ProductTemplateExportFilterRequest struct {

        
            Brands []string  `json:"brands"`
            Templates []string  `json:"templates"`
            FromDate string  `json:"from_date"`
            ToDate string  `json:"to_date"`
            CatalogueTypes []string  `json:"catalogue_types"`
         
    }
    
    // ProductTemplateDownloadsExport used by Catalog
    type ProductTemplateDownloadsExport struct {

        
            Type string  `json:"type"`
            Filters ProductTemplateExportFilterRequest  `json:"filters"`
            NotificationEmails []string  `json:"notification_emails"`
         
    }
    
    // ProductConfigurationDownloads used by Catalog
    type ProductConfigurationDownloads struct {

        
            Data []map[string]interface{}  `json:"data"`
            Multivalue bool  `json:"multivalue"`
         
    }
    
    // Media2 used by Catalog
    type Media2 struct {

        
            Portrait string  `json:"portrait"`
            Landscape string  `json:"landscape"`
            Logo string  `json:"logo"`
         
    }
    
    // Hierarchy used by Catalog
    type Hierarchy struct {

        
            L1 float64  `json:"l1"`
            L2 float64  `json:"l2"`
            Department float64  `json:"department"`
         
    }
    
    // CategoryMappingValues used by Catalog
    type CategoryMappingValues struct {

        
            Name string  `json:"name"`
            CatalogID float64  `json:"catalog_id"`
         
    }
    
    // CategoryMapping used by Catalog
    type CategoryMapping struct {

        
            Facebook CategoryMappingValues  `json:"facebook"`
            Google CategoryMappingValues  `json:"google"`
            Ajio CategoryMappingValues  `json:"ajio"`
         
    }
    
    // CategoryRequestBody used by Catalog
    type CategoryRequestBody struct {

        
            Slug string  `json:"slug"`
            Tryouts []string  `json:"tryouts"`
            Departments []float64  `json:"departments"`
            Media Media2  `json:"media"`
            IsActive bool  `json:"is_active"`
            Level float64  `json:"level"`
            Hierarchy []Hierarchy  `json:"hierarchy"`
            Synonyms []string  `json:"synonyms"`
            Marketplaces CategoryMapping  `json:"marketplaces"`
            Name string  `json:"name"`
            Priority float64  `json:"priority"`
         
    }
    
    // CategoryCreateResponse used by Catalog
    type CategoryCreateResponse struct {

        
            UID float64  `json:"uid"`
            Message string  `json:"message"`
         
    }
    
    // Category used by Catalog
    type Category struct {

        
            Slug string  `json:"slug"`
            Tryouts []string  `json:"tryouts"`
            CreatedOn string  `json:"created_on"`
            ID string  `json:"id"`
            Departments []float64  `json:"departments"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Media Media2  `json:"media"`
            UID float64  `json:"uid"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ModifiedOn string  `json:"modified_on"`
            IsActive bool  `json:"is_active"`
            Level float64  `json:"level"`
            Hierarchy []Hierarchy  `json:"hierarchy"`
            Synonyms []string  `json:"synonyms"`
            Marketplaces CategoryMapping  `json:"marketplaces"`
            Name string  `json:"name"`
            Priority float64  `json:"priority"`
         
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

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // CustomOrder used by Catalog
    type CustomOrder struct {

        
            ManufacturingTime float64  `json:"manufacturing_time"`
            ManufacturingTimeUnit string  `json:"manufacturing_time_unit"`
            IsCustomOrder bool  `json:"is_custom_order"`
         
    }
    
    // TeaserTag used by Catalog
    type TeaserTag struct {

        
            Tag string  `json:"tag"`
            URL string  `json:"url"`
         
    }
    
    // TaxIdentifier used by Catalog
    type TaxIdentifier struct {

        
            HsnCodeID string  `json:"hsn_code_id"`
            ReportingHsn string  `json:"reporting_hsn"`
            HsnCode string  `json:"hsn_code"`
         
    }
    
    // ProductPublish used by Catalog
    type ProductPublish struct {

        
            IsSet bool  `json:"is_set"`
            ProductOnlineDate string  `json:"product_online_date"`
         
    }
    
    // NetQuantity used by Catalog
    type NetQuantity struct {

        
            Value float64  `json:"value"`
            Unit interface{}  `json:"unit"`
         
    }
    
    // Trader used by Catalog
    type Trader struct {

        
            Type string  `json:"type"`
            Name interface{}  `json:"name"`
            Address []string  `json:"address"`
         
    }
    
    // ReturnConfig used by Catalog
    type ReturnConfig struct {

        
            Time float64  `json:"time"`
            Returnable bool  `json:"returnable"`
            Unit string  `json:"unit"`
         
    }
    
    // ProductCreateUpdateSchemaV2 used by Catalog
    type ProductCreateUpdateSchemaV2 struct {

        
            Description string  `json:"description"`
            Requester string  `json:"requester"`
            IsSet bool  `json:"is_set"`
            CustomOrder CustomOrder  `json:"custom_order"`
            ProductGroupTag []string  `json:"product_group_tag"`
            BrandUID float64  `json:"brand_uid"`
            Name string  `json:"name"`
            ShortDescription string  `json:"short_description"`
            Slug string  `json:"slug"`
            VariantMedia map[string]interface{}  `json:"variant_media"`
            IsActive bool  `json:"is_active"`
            CountryOfOrigin string  `json:"country_of_origin"`
            CategorySlug string  `json:"category_slug"`
            IsDependent bool  `json:"is_dependent"`
            Tags []string  `json:"tags"`
            TeaserTag TeaserTag  `json:"teaser_tag"`
            TemplateTag string  `json:"template_tag"`
            Departments []float64  `json:"departments"`
            VariantGroup map[string]interface{}  `json:"variant_group"`
            MultiSize bool  `json:"multi_size"`
            IsImageLessProduct bool  `json:"is_image_less_product"`
            TaxIdentifier TaxIdentifier  `json:"tax_identifier"`
            Sizes []map[string]interface{}  `json:"sizes"`
            ProductPublish ProductPublish  `json:"product_publish"`
            NetQuantity NetQuantity  `json:"net_quantity"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            BulkJobID string  `json:"bulk_job_id"`
            Highlights []string  `json:"highlights"`
            Attributes map[string]interface{}  `json:"attributes"`
            ItemType string  `json:"item_type"`
            Trader []Trader  `json:"trader"`
            CompanyID float64  `json:"company_id"`
            Action string  `json:"action"`
            Media []Media1  `json:"media"`
            UID float64  `json:"uid"`
            Currency string  `json:"currency"`
            ChangeRequestID interface{}  `json:"change_request_id"`
            NoOfBoxes float64  `json:"no_of_boxes"`
            SizeGuide string  `json:"size_guide"`
            ReturnConfig ReturnConfig  `json:"return_config"`
            ItemCode string  `json:"item_code"`
            Variants map[string]interface{}  `json:"variants"`
         
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

        
            Name string  `json:"name"`
            Logo Logo  `json:"logo"`
            UID float64  `json:"uid"`
         
    }
    
    // Image used by Catalog
    type Image struct {

        
            SecureURL string  `json:"secure_url"`
            AspectRatioF float64  `json:"aspect_ratio_f"`
            AspectRatio string  `json:"aspect_ratio"`
            URL string  `json:"url"`
         
    }
    
    // VerifiedBy used by Catalog
    type VerifiedBy struct {

        
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
         
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
    
    // ReturnConfigResponse used by Catalog
    type ReturnConfigResponse struct {

        
            Time float64  `json:"time"`
            Returnable bool  `json:"returnable"`
            Unit string  `json:"unit"`
         
    }
    
    // ProductSchemaV2 used by Catalog
    type ProductSchemaV2 struct {

        
            Brand Brand  `json:"brand"`
            Description string  `json:"description"`
            ID string  `json:"id"`
            L3Mapping []string  `json:"l3_mapping"`
            Moq map[string]interface{}  `json:"moq"`
            CategoryUID float64  `json:"category_uid"`
            IsSet bool  `json:"is_set"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            Images []Image  `json:"images"`
            ProductGroupTag []string  `json:"product_group_tag"`
            HsnCode string  `json:"hsn_code"`
            BrandUID float64  `json:"brand_uid"`
            Name string  `json:"name"`
            VerifiedBy VerifiedBy  `json:"verified_by"`
            Color string  `json:"color"`
            ShortDescription string  `json:"short_description"`
            Slug string  `json:"slug"`
            VerifiedOn string  `json:"verified_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            VariantMedia map[string]interface{}  `json:"variant_media"`
            IsPhysical bool  `json:"is_physical"`
            AllSizes []map[string]interface{}  `json:"all_sizes"`
            ImageNature string  `json:"image_nature"`
            IsActive bool  `json:"is_active"`
            CountryOfOrigin string  `json:"country_of_origin"`
            AllIdentifiers []string  `json:"all_identifiers"`
            CategorySlug string  `json:"category_slug"`
            IsDependent bool  `json:"is_dependent"`
            Stage string  `json:"stage"`
            Tags []string  `json:"tags"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            TemplateTag string  `json:"template_tag"`
            CreatedOn string  `json:"created_on"`
            Pending string  `json:"pending"`
            Departments []float64  `json:"departments"`
            Category map[string]interface{}  `json:"category"`
            VariantGroup map[string]interface{}  `json:"variant_group"`
            MultiSize bool  `json:"multi_size"`
            IsImageLessProduct bool  `json:"is_image_less_product"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            Sizes []map[string]interface{}  `json:"sizes"`
            ProductPublish ProductPublish1  `json:"product_publish"`
            NetQuantity NetQuantityResponse  `json:"net_quantity"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Highlights []string  `json:"highlights"`
            PrimaryColor string  `json:"primary_color"`
            Attributes map[string]interface{}  `json:"attributes"`
            ItemType string  `json:"item_type"`
            Trader []Trader  `json:"trader"`
            CompanyID float64  `json:"company_id"`
            Media []Media1  `json:"media"`
            UID float64  `json:"uid"`
            IsExpirable bool  `json:"is_expirable"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            AllCompanyIds []float64  `json:"all_company_ids"`
            Currency string  `json:"currency"`
            SizeGuide string  `json:"size_guide"`
            ReturnConfig ReturnConfigResponse  `json:"return_config"`
            NoOfBoxes float64  `json:"no_of_boxes"`
            ItemCode string  `json:"item_code"`
            ModifiedOn string  `json:"modified_on"`
            Variants map[string]interface{}  `json:"variants"`
         
    }
    
    // ProductListingResponseV2 used by Catalog
    type ProductListingResponseV2 struct {

        
            Items []ProductSchemaV2  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // ProductVariants used by Catalog
    type ProductVariants struct {

        
            Media []Media1  `json:"media"`
            UID float64  `json:"uid"`
            CategoryUID float64  `json:"category_uid"`
            ItemCode string  `json:"item_code"`
            BrandUID float64  `json:"brand_uid"`
            Name string  `json:"name"`
         
    }
    
    // ProductVariantsResponse used by Catalog
    type ProductVariantsResponse struct {

        
            Variants []ProductVariants  `json:"variants"`
            Page Page  `json:"page"`
         
    }
    
    // AttributeMasterSerializer used by Catalog
    type AttributeMasterSerializer struct {

        
            Description string  `json:"description"`
            Variant bool  `json:"variant"`
            Filters AttributeMasterFilter  `json:"filters"`
            Logo string  `json:"logo"`
            Name string  `json:"name"`
            RawKey string  `json:"raw_key"`
            Slug string  `json:"slug"`
            Schema AttributeMaster  `json:"schema"`
            Details AttributeMasterDetails  `json:"details"`
            IsNested bool  `json:"is_nested"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Unit string  `json:"unit"`
            Synonyms map[string]interface{}  `json:"synonyms"`
            Tags []string  `json:"tags"`
            CreatedOn string  `json:"created_on"`
            Departments []string  `json:"departments"`
            EnabledForEndConsumer bool  `json:"enabled_for_end_consumer"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Suggestion string  `json:"suggestion"`
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

        
            GtinValue string  `json:"gtin_value"`
            GtinType string  `json:"gtin_type"`
            Primary bool  `json:"primary"`
         
    }
    
    // AllSizes used by Catalog
    type AllSizes struct {

        
            ItemHeight float64  `json:"item_height"`
            ItemWeight float64  `json:"item_weight"`
            ItemLength float64  `json:"item_length"`
            ItemWidth float64  `json:"item_width"`
            Identifiers []ValidateIdentifier  `json:"identifiers"`
            Size interface{}  `json:"size"`
            ItemWeightUnitOfMeasure interface{}  `json:"item_weight_unit_of_measure"`
            ItemDimensionsUnitOfMeasure string  `json:"item_dimensions_unit_of_measure"`
         
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

        
            Brand Brand  `json:"brand"`
            Description string  `json:"description"`
            ID string  `json:"id"`
            L3Mapping []string  `json:"l3_mapping"`
            Moq map[string]interface{}  `json:"moq"`
            CategoryUID float64  `json:"category_uid"`
            IsSet bool  `json:"is_set"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            Images []Image  `json:"images"`
            ProductGroupTag []string  `json:"product_group_tag"`
            HsnCode string  `json:"hsn_code"`
            BrandUID float64  `json:"brand_uid"`
            Name string  `json:"name"`
            VerifiedBy VerifiedBy  `json:"verified_by"`
            Color string  `json:"color"`
            ShortDescription string  `json:"short_description"`
            Slug string  `json:"slug"`
            VerifiedOn string  `json:"verified_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            VariantMedia map[string]interface{}  `json:"variant_media"`
            IsPhysical bool  `json:"is_physical"`
            AllSizes []map[string]interface{}  `json:"all_sizes"`
            ImageNature string  `json:"image_nature"`
            IsActive bool  `json:"is_active"`
            CountryOfOrigin string  `json:"country_of_origin"`
            AllIdentifiers []string  `json:"all_identifiers"`
            CategorySlug string  `json:"category_slug"`
            IsDependent bool  `json:"is_dependent"`
            Stage string  `json:"stage"`
            Tags []string  `json:"tags"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            TemplateTag string  `json:"template_tag"`
            CreatedOn string  `json:"created_on"`
            Pending string  `json:"pending"`
            Departments []float64  `json:"departments"`
            Category map[string]interface{}  `json:"category"`
            VariantGroup map[string]interface{}  `json:"variant_group"`
            MultiSize bool  `json:"multi_size"`
            IsImageLessProduct bool  `json:"is_image_less_product"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            Sizes []map[string]interface{}  `json:"sizes"`
            ProductPublish ProductPublished  `json:"product_publish"`
            NetQuantity NetQuantityResponse  `json:"net_quantity"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Highlights []string  `json:"highlights"`
            PrimaryColor string  `json:"primary_color"`
            Attributes map[string]interface{}  `json:"attributes"`
            ItemType string  `json:"item_type"`
            Trader []Trader  `json:"trader"`
            CompanyID float64  `json:"company_id"`
            Media []Media1  `json:"media"`
            UID float64  `json:"uid"`
            IsExpirable bool  `json:"is_expirable"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            AllCompanyIds []float64  `json:"all_company_ids"`
            Currency string  `json:"currency"`
            SizeGuide string  `json:"size_guide"`
            ReturnConfig ReturnConfigResponse  `json:"return_config"`
            NoOfBoxes float64  `json:"no_of_boxes"`
            ItemCode string  `json:"item_code"`
            ModifiedOn string  `json:"modified_on"`
            Variants map[string]interface{}  `json:"variants"`
         
    }
    
    // ProductListingResponse used by Catalog
    type ProductListingResponse struct {

        
            Items []Product  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // BulkJob used by Catalog
    type BulkJob struct {

        
            TemplateTag string  `json:"template_tag"`
            CreatedOn string  `json:"created_on"`
            TrackingURL string  `json:"tracking_url"`
            ModifiedBy UserInfo1  `json:"modified_by"`
            CancelledRecords []map[string]interface{}  `json:"cancelled_records"`
            CreatedBy UserInfo1  `json:"created_by"`
            CustomTemplateTag string  `json:"custom_template_tag"`
            IsActive bool  `json:"is_active"`
            Cancelled float64  `json:"cancelled"`
            Failed float64  `json:"failed"`
            FailedRecords []map[string]interface{}  `json:"failed_records"`
            Total float64  `json:"total"`
            FilePath string  `json:"file_path"`
            Succeed float64  `json:"succeed"`
            ModifiedOn string  `json:"modified_on"`
            CompanyID float64  `json:"company_id"`
            Stage string  `json:"stage"`
         
    }
    
    // BulkResponse used by Catalog
    type BulkResponse struct {

        
            CreatedOn string  `json:"created_on"`
            ModifiedBy UserInfo1  `json:"modified_by"`
            BatchID string  `json:"batch_id"`
            CreatedBy UserInfo1  `json:"created_by"`
            IsActive bool  `json:"is_active"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // UserDetail1 used by Catalog
    type UserDetail1 struct {

        
            UserID string  `json:"user_id"`
            FullName string  `json:"full_name"`
            Username string  `json:"username"`
         
    }
    
    // ProductBulkRequest used by Catalog
    type ProductBulkRequest struct {

        
            TemplateTag string  `json:"template_tag"`
            CreatedOn string  `json:"created_on"`
            Template ProductTemplate  `json:"template"`
            ModifiedBy UserDetail1  `json:"modified_by"`
            CancelledRecords []string  `json:"cancelled_records"`
            CreatedBy UserDetail1  `json:"created_by"`
            IsActive bool  `json:"is_active"`
            Cancelled float64  `json:"cancelled"`
            Failed float64  `json:"failed"`
            FailedRecords []string  `json:"failed_records"`
            Total float64  `json:"total"`
            FilePath string  `json:"file_path"`
            Succeed float64  `json:"succeed"`
            ModifiedOn string  `json:"modified_on"`
            CompanyID float64  `json:"company_id"`
            Stage string  `json:"stage"`
         
    }
    
    // ProductBulkRequestList used by Catalog
    type ProductBulkRequestList struct {

        
            Items ProductBulkRequest  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // BulkProductRequest used by Catalog
    type BulkProductRequest struct {

        
            TemplateTag string  `json:"template_tag"`
            Data []map[string]interface{}  `json:"data"`
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

        
            CreatedOn string  `json:"created_on"`
            TrackingURL string  `json:"tracking_url"`
            ID string  `json:"id"`
            Retry float64  `json:"retry"`
            ModifiedBy UserCommon  `json:"modified_by"`
            CancelledRecords []string  `json:"cancelled_records"`
            CreatedBy UserCommon  `json:"created_by"`
            IsActive bool  `json:"is_active"`
            Cancelled float64  `json:"cancelled"`
            Failed float64  `json:"failed"`
            FailedRecords []string  `json:"failed_records"`
            Total float64  `json:"total"`
            FilePath string  `json:"file_path"`
            Succeed float64  `json:"succeed"`
            ModifiedOn string  `json:"modified_on"`
            CompanyID float64  `json:"company_id"`
            Stage string  `json:"stage"`
         
    }
    
    // BulkAssetResponse used by Catalog
    type BulkAssetResponse struct {

        
            Items []Items  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // ProductSizeDeleteDataResponse used by Catalog
    type ProductSizeDeleteDataResponse struct {

        
            ItemID float64  `json:"item_id"`
            Size string  `json:"size"`
            CompanyID float64  `json:"company_id"`
         
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

        
            GtinValue interface{}  `json:"gtin_value"`
            GtinType string  `json:"gtin_type"`
            Primary bool  `json:"primary"`
         
    }
    
    // InvSize used by Catalog
    type InvSize struct {

        
            ItemHeight float64  `json:"item_height"`
            ExpirationDate string  `json:"expiration_date"`
            Price float64  `json:"price"`
            Set InventorySet  `json:"set"`
            Currency string  `json:"currency"`
            IsSet bool  `json:"is_set"`
            ItemWeight float64  `json:"item_weight"`
            ItemWidth float64  `json:"item_width"`
            ItemLength float64  `json:"item_length"`
            Size interface{}  `json:"size"`
            ItemWeightUnitOfMeasure string  `json:"item_weight_unit_of_measure"`
            PriceTransfer float64  `json:"price_transfer"`
            Identifiers []GTIN  `json:"identifiers"`
            Quantity float64  `json:"quantity"`
            StoreCode string  `json:"store_code"`
            PriceEffective float64  `json:"price_effective"`
            ItemDimensionsUnitOfMeasure string  `json:"item_dimensions_unit_of_measure"`
         
    }
    
    // InventoryRequest used by Catalog
    type InventoryRequest struct {

        
            Item ItemQuery  `json:"item"`
            CompanyID float64  `json:"company_id"`
            Sizes []InvSize  `json:"sizes"`
         
    }
    
    // InventoryResponse used by Catalog
    type InventoryResponse struct {

        
            InventoryUpdatedOn string  `json:"inventory_updated_on"`
            Price float64  `json:"price"`
            UID string  `json:"uid"`
            ItemID float64  `json:"item_id"`
            SellerIdentifier string  `json:"seller_identifier"`
            Currency string  `json:"currency"`
            Identifiers map[string]interface{}  `json:"identifiers"`
            Quantity float64  `json:"quantity"`
            Size string  `json:"size"`
            PriceTransfer float64  `json:"price_transfer"`
            Store map[string]interface{}  `json:"store"`
            PriceEffective float64  `json:"price_effective"`
            SellableQuantity float64  `json:"sellable_quantity"`
         
    }
    
    // InventoryResponsePaginated used by Catalog
    type InventoryResponsePaginated struct {

        
            Items []InventoryResponse  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // BrandMeta used by Catalog
    type BrandMeta struct {

        
            Name string  `json:"name"`
            ID float64  `json:"id"`
         
    }
    
    // QuantityBase used by Catalog
    type QuantityBase struct {

        
            Count float64  `json:"count"`
            UpdatedAt string  `json:"updated_at"`
         
    }
    
    // Quantities used by Catalog
    type Quantities struct {

        
            Damaged QuantityBase  `json:"damaged"`
            OrderCommitted QuantityBase  `json:"order_committed"`
            Sellable QuantityBase  `json:"sellable"`
            NotAvailable QuantityBase  `json:"not_available"`
         
    }
    
    // ManufacturerResponse used by Catalog
    type ManufacturerResponse struct {

        
            Address string  `json:"address"`
            Name string  `json:"name"`
            IsDefault bool  `json:"is_default"`
         
    }
    
    // DimensionResponse used by Catalog
    type DimensionResponse struct {

        
            IsDefault bool  `json:"is_default"`
            Unit string  `json:"unit"`
            Width float64  `json:"width"`
            Height float64  `json:"height"`
            Length float64  `json:"length"`
         
    }
    
    // PriceMeta used by Catalog
    type PriceMeta struct {

        
            Transfer float64  `json:"transfer"`
            Currency string  `json:"currency"`
            Effective float64  `json:"effective"`
            UpdatedAt string  `json:"updated_at"`
            TpNotes map[string]interface{}  `json:"tp_notes"`
            Marked float64  `json:"marked"`
         
    }
    
    // WeightResponse used by Catalog
    type WeightResponse struct {

        
            Shipping float64  `json:"shipping"`
            IsDefault bool  `json:"is_default"`
            Unit string  `json:"unit"`
         
    }
    
    // StoreMeta used by Catalog
    type StoreMeta struct {

        
            ID float64  `json:"id"`
         
    }
    
    // Trader1 used by Catalog
    type Trader1 struct {

        
            Type string  `json:"type"`
            Name string  `json:"name"`
            Address []string  `json:"address"`
         
    }
    
    // CompanyMeta used by Catalog
    type CompanyMeta struct {

        
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

        
            Brand BrandMeta  `json:"brand"`
            Quantities Quantities  `json:"quantities"`
            ItemID float64  `json:"item_id"`
            Set InventorySet  `json:"set"`
            IsSet bool  `json:"is_set"`
            Manufacturer ManufacturerResponse  `json:"manufacturer"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            RawMeta map[string]interface{}  `json:"raw_meta"`
            Dimension DimensionResponse  `json:"dimension"`
            IsActive bool  `json:"is_active"`
            Meta map[string]interface{}  `json:"meta"`
            CountryOfOrigin string  `json:"country_of_origin"`
            Identifier map[string]interface{}  `json:"identifier"`
            FyndItemCode string  `json:"fynd_item_code"`
            TraceID string  `json:"trace_id"`
            Stage string  `json:"stage"`
            Tags []string  `json:"tags"`
            ExpirationDate string  `json:"expiration_date"`
            Price PriceMeta  `json:"price"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Weight WeightResponse  `json:"weight"`
            Store StoreMeta  `json:"store"`
            Trader []Trader1  `json:"trader"`
            SellerIdentifier string  `json:"seller_identifier"`
            TrackInventory bool  `json:"track_inventory"`
            FyndMeta map[string]interface{}  `json:"fynd_meta"`
            TotalQuantity float64  `json:"total_quantity"`
            AddedOnStore string  `json:"added_on_store"`
            Fragile bool  `json:"fragile"`
            UID string  `json:"uid"`
            Company CompanyMeta  `json:"company"`
            CreatedBy UserSerializer  `json:"created_by"`
            ReturnConfig ReturnConfig1  `json:"return_config"`
            FyndArticleCode string  `json:"fynd_article_code"`
            Size string  `json:"size"`
         
    }
    
    // InventorySellerIdentifierResponsePaginated used by Catalog
    type InventorySellerIdentifierResponsePaginated struct {

        
            Items []InventorySellerResponse  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // BrandMeta1 used by Catalog
    type BrandMeta1 struct {

        
            Name string  `json:"name"`
            ID float64  `json:"id"`
         
    }
    
    // Quantity used by Catalog
    type Quantity struct {

        
            Count float64  `json:"count"`
         
    }
    
    // QuantitiesArticle used by Catalog
    type QuantitiesArticle struct {

        
            Damaged Quantity  `json:"damaged"`
            OrderCommitted Quantity  `json:"order_committed"`
            Sellable Quantity  `json:"sellable"`
            NotAvailable Quantity  `json:"not_available"`
         
    }
    
    // ManufacturerResponse1 used by Catalog
    type ManufacturerResponse1 struct {

        
            Address string  `json:"address"`
            Name string  `json:"name"`
            IsDefault bool  `json:"is_default"`
         
    }
    
    // DimensionResponse1 used by Catalog
    type DimensionResponse1 struct {

        
            Height float64  `json:"height"`
            Width float64  `json:"width"`
            Length float64  `json:"length"`
            Unit string  `json:"unit"`
         
    }
    
    // PriceArticle used by Catalog
    type PriceArticle struct {

        
            Transfer float64  `json:"transfer"`
            Currency string  `json:"currency"`
            Effective float64  `json:"effective"`
            TpNotes map[string]interface{}  `json:"tp_notes"`
            Marked float64  `json:"marked"`
         
    }
    
    // WeightResponse1 used by Catalog
    type WeightResponse1 struct {

        
            Shipping float64  `json:"shipping"`
            Unit string  `json:"unit"`
         
    }
    
    // ArticleStoreResponse used by Catalog
    type ArticleStoreResponse struct {

        
            StoreType string  `json:"store_type"`
            Name string  `json:"name"`
            StoreCode string  `json:"store_code"`
            UID float64  `json:"uid"`
         
    }
    
    // Trader2 used by Catalog
    type Trader2 struct {

        
            Type string  `json:"type"`
            Name string  `json:"name"`
            Address []string  `json:"address"`
         
    }
    
    // CompanyMeta1 used by Catalog
    type CompanyMeta1 struct {

        
            ID float64  `json:"id"`
         
    }
    
    // DateMeta used by Catalog
    type DateMeta struct {

        
            AddedOnStore string  `json:"added_on_store"`
            CreatedOn string  `json:"created_on"`
            InventoryUpdatedOn string  `json:"inventory_updated_on"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // ReturnConfig2 used by Catalog
    type ReturnConfig2 struct {

        
            Time float64  `json:"time"`
            Returnable bool  `json:"returnable"`
            Unit string  `json:"unit"`
         
    }
    
    // GetInventories used by Catalog
    type GetInventories struct {

        
            Brand BrandMeta1  `json:"brand"`
            ID string  `json:"id"`
            Quantities QuantitiesArticle  `json:"quantities"`
            ItemID float64  `json:"item_id"`
            IsSet bool  `json:"is_set"`
            Platforms map[string]interface{}  `json:"platforms"`
            Manufacturer ManufacturerResponse1  `json:"manufacturer"`
            InventoryUpdatedOn string  `json:"inventory_updated_on"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            Dimension DimensionResponse1  `json:"dimension"`
            CountryOfOrigin string  `json:"country_of_origin"`
            Identifier map[string]interface{}  `json:"identifier"`
            TraceID string  `json:"trace_id"`
            Stage string  `json:"stage"`
            Tags []string  `json:"tags"`
            ExpirationDate string  `json:"expiration_date"`
            Price PriceArticle  `json:"price"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            Weight WeightResponse1  `json:"weight"`
            Store ArticleStoreResponse  `json:"store"`
            Trader []Trader2  `json:"trader"`
            SellerIdentifier string  `json:"seller_identifier"`
            TrackInventory bool  `json:"track_inventory"`
            TotalQuantity float64  `json:"total_quantity"`
            Company CompanyMeta1  `json:"company"`
            UID string  `json:"uid"`
            DateMeta DateMeta  `json:"date_meta"`
            CreatedBy UserSerializer  `json:"created_by"`
            ReturnConfig ReturnConfig2  `json:"return_config"`
            Size string  `json:"size"`
         
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
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            CancelledRecords []string  `json:"cancelled_records"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Cancelled float64  `json:"cancelled"`
            IsActive bool  `json:"is_active"`
            Failed float64  `json:"failed"`
            FailedRecords []string  `json:"failed_records"`
            Total float64  `json:"total"`
            FilePath string  `json:"file_path"`
            Succeed float64  `json:"succeed"`
            ModifiedOn string  `json:"modified_on"`
            CompanyID float64  `json:"company_id"`
            Stage string  `json:"stage"`
         
    }
    
    // BulkInventoryGet used by Catalog
    type BulkInventoryGet struct {

        
            Items []BulkInventoryGetItems  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // InventoryJobPayload used by Catalog
    type InventoryJobPayload struct {

        
            TotalQuantity float64  `json:"total_quantity"`
            ExpirationDate string  `json:"expiration_date"`
            ItemDimensionsUnitOfMeasure string  `json:"item_dimensions_unit_of_measure"`
            PriceMarked float64  `json:"price_marked"`
            Price float64  `json:"price"`
            Currency string  `json:"currency"`
            ItemWeightUnitOfMeasure string  `json:"item_weight_unit_of_measure"`
            Quantity float64  `json:"quantity"`
            PriceEffective float64  `json:"price_effective"`
            TraceID string  `json:"trace_id"`
            StoreCode string  `json:"store_code"`
            SellerIdentifier string  `json:"seller_identifier"`
            Tags []string  `json:"tags"`
         
    }
    
    // InventoryBulkRequest used by Catalog
    type InventoryBulkRequest struct {

        
            BatchID string  `json:"batch_id"`
            CompanyID float64  `json:"company_id"`
            User map[string]interface{}  `json:"user"`
            Sizes []InventoryJobPayload  `json:"sizes"`
         
    }
    
    // InventoryExportRequest used by Catalog
    type InventoryExportRequest struct {

        
            Type string  `json:"type"`
            Brand []float64  `json:"brand"`
            Store []float64  `json:"store"`
         
    }
    
    // InventoryExportResponse used by Catalog
    type InventoryExportResponse struct {

        
            CreatedOn string  `json:"created_on"`
            TaskID string  `json:"task_id"`
            Filters map[string]interface{}  `json:"filters"`
            Status string  `json:"status"`
            NotificationEmails []string  `json:"notification_emails"`
            SellerID float64  `json:"seller_id"`
            CreatedBy string  `json:"created_by"`
            Type string  `json:"type"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // InventoryExportQuantityFilter used by Catalog
    type InventoryExportQuantityFilter struct {

        
            Max float64  `json:"max"`
            Operators string  `json:"operators"`
            Min float64  `json:"min"`
         
    }
    
    // InventoryExportAdvanceOption used by Catalog
    type InventoryExportAdvanceOption struct {

        
            StoreIds []float64  `json:"store_ids"`
            FromDate string  `json:"from_date"`
            Quantity InventoryExportQuantityFilter  `json:"quantity"`
            ToDate string  `json:"to_date"`
            BrandIds []float64  `json:"brand_ids"`
         
    }
    
    // InventoryExportJob used by Catalog
    type InventoryExportJob struct {

        
            TaskID string  `json:"task_id"`
            Filters InventoryExportAdvanceOption  `json:"filters"`
            Status string  `json:"status"`
            SellerID float64  `json:"seller_id"`
            URL string  `json:"url"`
            Type string  `json:"type"`
            CompletedOn string  `json:"completed_on"`
            NotificationEmails []string  `json:"notification_emails"`
         
    }
    
    // InventoryExportFilter used by Catalog
    type InventoryExportFilter struct {

        
            StoreIds []float64  `json:"store_ids"`
            FromDate string  `json:"from_date"`
            Quantity InventoryExportQuantityFilter  `json:"quantity"`
            ToDate string  `json:"to_date"`
            BrandIds []float64  `json:"brand_ids"`
         
    }
    
    // InventoryCreateRequest used by Catalog
    type InventoryCreateRequest struct {

        
            Type string  `json:"type"`
            Filters InventoryExportFilter  `json:"filters"`
            NotificationEmails []string  `json:"notification_emails"`
            Data []string  `json:"data"`
         
    }
    
    // InventoryJobFilters used by Catalog
    type InventoryJobFilters struct {

        
            Brands []string  `json:"brands"`
            Stores []string  `json:"stores"`
            FromDate string  `json:"from_date"`
            Quantity InventoryExportQuantityFilter  `json:"quantity"`
            ToDate string  `json:"to_date"`
         
    }
    
    // InventoryJobDetailResponse used by Catalog
    type InventoryJobDetailResponse struct {

        
            CreatedOn string  `json:"created_on"`
            TaskID string  `json:"task_id"`
            ID string  `json:"id"`
            Filters InventoryJobFilters  `json:"filters"`
            Status interface{}  `json:"status"`
            NotificationEmails []string  `json:"notification_emails"`
            SellerID float64  `json:"seller_id"`
            CreatedBy UserDetail  `json:"created_by"`
            URL string  `json:"url"`
            CancelledBy UserDetail  `json:"cancelled_by"`
            Type string  `json:"type"`
            CompletedOn string  `json:"completed_on"`
            CancelledOn string  `json:"cancelled_on"`
            ModifiedOn string  `json:"modified_on"`
         
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

        
            TotalQuantity float64  `json:"total_quantity"`
            ExpirationDate string  `json:"expiration_date"`
            PriceMarked float64  `json:"price_marked"`
            PriceEffective float64  `json:"price_effective"`
            TraceID string  `json:"trace_id"`
            StoreID float64  `json:"store_id"`
            SellerIdentifier string  `json:"seller_identifier"`
            Tags []string  `json:"tags"`
         
    }
    
    // InventoryRequestSchemaV2 used by Catalog
    type InventoryRequestSchemaV2 struct {

        
            Meta map[string]interface{}  `json:"meta"`
            CompanyID float64  `json:"company_id"`
            Payload []InventoryPayload  `json:"payload"`
         
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
    
    // HsnCodesObject used by Catalog
    type HsnCodesObject struct {

        
            Threshold1 float64  `json:"threshold1"`
            Hs2Code string  `json:"hs2_code"`
            Tax2 float64  `json:"tax2"`
            ID string  `json:"id"`
            Tax1 float64  `json:"tax1"`
            TaxOnMrp bool  `json:"tax_on_mrp"`
            HsnCode string  `json:"hsn_code"`
            Threshold2 float64  `json:"threshold2"`
            ModifiedOn string  `json:"modified_on"`
            CompanyID float64  `json:"company_id"`
            TaxOnEsp bool  `json:"tax_on_esp"`
         
    }
    
    // HsnCode used by Catalog
    type HsnCode struct {

        
            Data HsnCodesObject  `json:"data"`
         
    }
    
    // HsnUpsert used by Catalog
    type HsnUpsert struct {

        
            Threshold1 float64  `json:"threshold1"`
            Hs2Code string  `json:"hs2_code"`
            Tax2 float64  `json:"tax2"`
            UID float64  `json:"uid"`
            Tax1 float64  `json:"tax1"`
            IsActive bool  `json:"is_active"`
            TaxOnMrp bool  `json:"tax_on_mrp"`
            HsnCode string  `json:"hsn_code"`
            Threshold2 float64  `json:"threshold2"`
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
    
    // TaxSlab used by Catalog
    type TaxSlab struct {

        
            Threshold float64  `json:"threshold"`
            Rate float64  `json:"rate"`
            EffectiveDate string  `json:"effective_date"`
            Cess float64  `json:"cess"`
         
    }
    
    // HSNDataInsertV2 used by Catalog
    type HSNDataInsertV2 struct {

        
            Description string  `json:"description"`
            CreatedOn string  `json:"created_on"`
            HsnCodeID string  `json:"hsn_code_id"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Taxes []TaxSlab  `json:"taxes"`
            CountryCode string  `json:"country_code"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Type string  `json:"type"`
            HsnCode string  `json:"hsn_code"`
            ModifiedOn string  `json:"modified_on"`
            ReportingHsn string  `json:"reporting_hsn"`
         
    }
    
    // PageResponse used by Catalog
    type PageResponse struct {

        
            HasNext bool  `json:"has_next"`
            ItemTotal float64  `json:"item_total"`
            Size float64  `json:"size"`
            Current string  `json:"current"`
            HasPrevious bool  `json:"has_previous"`
         
    }
    
    // HsnCodesListingResponseSchemaV2 used by Catalog
    type HsnCodesListingResponseSchemaV2 struct {

        
            Items []HSNDataInsertV2  `json:"items"`
            Page PageResponse  `json:"page"`
         
    }
    
    // BrandItem used by Catalog
    type BrandItem struct {

        
            Banners ImageUrls  `json:"banners"`
            Slug string  `json:"slug"`
            Discount string  `json:"discount"`
            Departments []string  `json:"departments"`
            UID float64  `json:"uid"`
            Logo Media  `json:"logo"`
            Name string  `json:"name"`
            Action Action  `json:"action"`
         
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
            UID float64  `json:"uid"`
            Logo Media  `json:"logo"`
            Name string  `json:"name"`
         
    }
    
    // DepartmentResponse used by Catalog
    type DepartmentResponse struct {

        
            Items []Department  `json:"items"`
         
    }
    
    // ThirdLevelChild used by Catalog
    type ThirdLevelChild struct {

        
            Banners ImageUrls  `json:"banners"`
            Slug string  `json:"slug"`
            Childs []map[string]interface{}  `json:"childs"`
            UID float64  `json:"uid"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Name string  `json:"name"`
            Action Action  `json:"action"`
         
    }
    
    // SecondLevelChild used by Catalog
    type SecondLevelChild struct {

        
            Banners ImageUrls  `json:"banners"`
            Slug string  `json:"slug"`
            Childs []ThirdLevelChild  `json:"childs"`
            UID float64  `json:"uid"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Name string  `json:"name"`
            Action Action  `json:"action"`
         
    }
    
    // Child used by Catalog
    type Child struct {

        
            Banners ImageUrls  `json:"banners"`
            Slug string  `json:"slug"`
            Childs []SecondLevelChild  `json:"childs"`
            UID float64  `json:"uid"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Name string  `json:"name"`
            Action Action  `json:"action"`
         
    }
    
    // CategoryItems used by Catalog
    type CategoryItems struct {

        
            Banners ImageUrls  `json:"banners"`
            Slug string  `json:"slug"`
            Childs []Child  `json:"childs"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
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
            Page Page  `json:"page"`
            Operators map[string]interface{}  `json:"operators"`
            SortOn []ProductSortOn  `json:"sort_on"`
         
    }
    
    // ProductDetail used by Catalog
    type ProductDetail struct {

        
            Brand ProductBrand  `json:"brand"`
            Description string  `json:"description"`
            PromoMeta map[string]interface{}  `json:"promo_meta"`
            HasVariant bool  `json:"has_variant"`
            Type string  `json:"type"`
            Name string  `json:"name"`
            Color string  `json:"color"`
            ShortDescription string  `json:"short_description"`
            Slug string  `json:"slug"`
            RatingCount float64  `json:"rating_count"`
            ImageNature string  `json:"image_nature"`
            Similars []string  `json:"similars"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            Medias []Media1  `json:"medias"`
            ProductOnlineDate string  `json:"product_online_date"`
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
            Highlights []string  `json:"highlights"`
            Attributes map[string]interface{}  `json:"attributes"`
            ItemType string  `json:"item_type"`
            UID float64  `json:"uid"`
            ItemCode string  `json:"item_code"`
            Rating float64  `json:"rating"`
            Tryouts []string  `json:"tryouts"`
         
    }
    
    // InventoryPage used by Catalog
    type InventoryPage struct {

        
            HasNext bool  `json:"has_next"`
            ItemTotal float64  `json:"item_total"`
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

        
            Level string  `json:"level"`
            Strategy string  `json:"strategy"`
         
    }
    
    // AssignStoreArticle used by Catalog
    type AssignStoreArticle struct {

        
            GroupID string  `json:"group_id"`
            Query ArticleQuery  `json:"query"`
            Meta map[string]interface{}  `json:"meta"`
            Quantity float64  `json:"quantity"`
            ArticleAssignment ArticleAssignment  `json:"article_assignment"`
         
    }
    
    // AssignStore used by Catalog
    type AssignStore struct {

        
            Articles []AssignStoreArticle  `json:"articles"`
            AppID string  `json:"app_id"`
            Pincode string  `json:"pincode"`
            ChannelType string  `json:"channel_type"`
            StoreIds []float64  `json:"store_ids"`
            CompanyID float64  `json:"company_id"`
            ChannelIdentifier string  `json:"channel_identifier"`
         
    }
    
    // ArticleAssignment1 used by Catalog
    type ArticleAssignment1 struct {

        
            Level string  `json:"level"`
            Strategy string  `json:"strategy"`
         
    }
    
    // StoreAssignResponse used by Catalog
    type StoreAssignResponse struct {

        
            Index float64  `json:"index"`
            GroupID string  `json:"group_id"`
            PriceMarked float64  `json:"price_marked"`
            Status bool  `json:"status"`
            UID string  `json:"uid"`
            ItemID float64  `json:"item_id"`
            StorePincode float64  `json:"store_pincode"`
            Meta map[string]interface{}  `json:"meta"`
            Quantity float64  `json:"quantity"`
            Size string  `json:"size"`
            PriceEffective float64  `json:"price_effective"`
            ID string  `json:"_id"`
            ArticleAssignment ArticleAssignment1  `json:"article_assignment"`
            SCity string  `json:"s_city"`
            StrategyWiseListing []map[string]interface{}  `json:"strategy_wise_listing"`
            CompanyID float64  `json:"company_id"`
            StoreID float64  `json:"store_id"`
         
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
    
    // UserSerializer1 used by Catalog
    type UserSerializer1 struct {

        
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
            Contact string  `json:"contact"`
         
    }
    
    // Document used by Catalog
    type Document struct {

        
            Verified bool  `json:"verified"`
            LegalName string  `json:"legal_name"`
            Value string  `json:"value"`
            URL string  `json:"url"`
            Type string  `json:"type"`
         
    }
    
    // LocationIntegrationType used by Catalog
    type LocationIntegrationType struct {

        
            Order string  `json:"order"`
            Inventory string  `json:"inventory"`
         
    }
    
    // ProductReturnConfigSerializer used by Catalog
    type ProductReturnConfigSerializer struct {

        
            OnSameStore bool  `json:"on_same_store"`
            StoreUID float64  `json:"store_uid"`
         
    }
    
    // SellerPhoneNumber used by Catalog
    type SellerPhoneNumber struct {

        
            CountryCode float64  `json:"country_code"`
            Number string  `json:"number"`
         
    }
    
    // LocationManagerSerializer used by Catalog
    type LocationManagerSerializer struct {

        
            MobileNo SellerPhoneNumber  `json:"mobile_no"`
            Name string  `json:"name"`
            Email string  `json:"email"`
         
    }
    
    // GetAddressSerializer used by Catalog
    type GetAddressSerializer struct {

        
            Address1 string  `json:"address1"`
            Latitude float64  `json:"latitude"`
            City string  `json:"city"`
            CountryCode string  `json:"country_code"`
            Longitude float64  `json:"longitude"`
            Pincode float64  `json:"pincode"`
            Landmark string  `json:"landmark"`
            AddressType string  `json:"address_type"`
            State string  `json:"state"`
            Address2 string  `json:"address2"`
            Country string  `json:"country"`
         
    }
    
    // UserSerializer2 used by Catalog
    type UserSerializer2 struct {

        
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
            Contact string  `json:"contact"`
         
    }
    
    // GetCompanySerializer used by Catalog
    type GetCompanySerializer struct {

        
            Addresses []GetAddressSerializer  `json:"addresses"`
            CreatedOn string  `json:"created_on"`
            BusinessType string  `json:"business_type"`
            ModifiedBy UserSerializer2  `json:"modified_by"`
            UID float64  `json:"uid"`
            CreatedBy UserSerializer2  `json:"created_by"`
            CompanyType string  `json:"company_type"`
            VerifiedBy UserSerializer2  `json:"verified_by"`
            RejectReason string  `json:"reject_reason"`
            Name string  `json:"name"`
            ModifiedOn string  `json:"modified_on"`
            Stage string  `json:"stage"`
            VerifiedOn string  `json:"verified_on"`
         
    }
    
    // LocationTimingSerializer used by Catalog
    type LocationTimingSerializer struct {

        
            Minute float64  `json:"minute"`
            Hour float64  `json:"hour"`
         
    }
    
    // LocationDayWiseSerializer used by Catalog
    type LocationDayWiseSerializer struct {

        
            Closing LocationTimingSerializer  `json:"closing"`
            Opening LocationTimingSerializer  `json:"opening"`
            Open bool  `json:"open"`
            Weekday string  `json:"weekday"`
         
    }
    
    // GetLocationSerializer used by Catalog
    type GetLocationSerializer struct {

        
            GstCredentials InvoiceDetailsSerializer  `json:"gst_credentials"`
            VerifiedBy UserSerializer1  `json:"verified_by"`
            Documents []Document  `json:"documents"`
            Name string  `json:"name"`
            PhoneNumber string  `json:"phone_number"`
            VerifiedOn string  `json:"verified_on"`
            ModifiedBy UserSerializer1  `json:"modified_by"`
            NotificationEmails []string  `json:"notification_emails"`
            IntegrationType LocationIntegrationType  `json:"integration_type"`
            Stage string  `json:"stage"`
            ProductReturnConfig ProductReturnConfigSerializer  `json:"product_return_config"`
            DisplayName string  `json:"display_name"`
            CreatedOn string  `json:"created_on"`
            Warnings map[string]interface{}  `json:"warnings"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            StoreType string  `json:"store_type"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            Manager LocationManagerSerializer  `json:"manager"`
            UID float64  `json:"uid"`
            Company GetCompanySerializer  `json:"company"`
            CreatedBy UserSerializer1  `json:"created_by"`
            Code string  `json:"code"`
            Timing []LocationDayWiseSerializer  `json:"timing"`
            Address GetAddressSerializer  `json:"address"`
            ModifiedOn string  `json:"modified_on"`
         
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
            UID float64  `json:"uid"`
            Logo string  `json:"logo"`
            IsActive bool  `json:"is_active"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Name string  `json:"name"`
         
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

        
            Contact string  `json:"contact"`
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
         
    }
    
    // BusinessCountryInfo used by CompanyProfile
    type BusinessCountryInfo struct {

        
            Country string  `json:"country"`
            CountryCode string  `json:"country_code"`
         
    }
    
    // CompanyTaxesSerializer used by CompanyProfile
    type CompanyTaxesSerializer struct {

        
            EffectiveDate string  `json:"effective_date"`
            Rate float64  `json:"rate"`
            Enable bool  `json:"enable"`
         
    }
    
    // GetAddressSerializer used by CompanyProfile
    type GetAddressSerializer struct {

        
            Address1 string  `json:"address1"`
            Landmark string  `json:"landmark"`
            Address2 string  `json:"address2"`
            AddressType string  `json:"address_type"`
            Latitude float64  `json:"latitude"`
            Longitude float64  `json:"longitude"`
            Pincode float64  `json:"pincode"`
            Country string  `json:"country"`
            City string  `json:"city"`
            CountryCode string  `json:"country_code"`
            State string  `json:"state"`
         
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

        
            LegalName string  `json:"legal_name"`
            URL string  `json:"url"`
            Type string  `json:"type"`
            Verified bool  `json:"verified"`
            Value string  `json:"value"`
         
    }
    
    // GetCompanyProfileSerializerResponse used by CompanyProfile
    type GetCompanyProfileSerializerResponse struct {

        
            CreatedOn string  `json:"created_on"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            Mode string  `json:"mode"`
            Name string  `json:"name"`
            CompanyType string  `json:"company_type"`
            VerifiedOn string  `json:"verified_on"`
            BusinessCountryInfo BusinessCountryInfo  `json:"business_country_info"`
            FranchiseEnabled bool  `json:"franchise_enabled"`
            Stage string  `json:"stage"`
            Taxes []CompanyTaxesSerializer  `json:"taxes"`
            BusinessInfo string  `json:"business_info"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            CreatedBy UserSerializer  `json:"created_by"`
            ContactDetails ContactDetails  `json:"contact_details"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            BusinessDetails BusinessDetails  `json:"business_details"`
            BusinessType string  `json:"business_type"`
            Warnings map[string]interface{}  `json:"warnings"`
            ModifiedOn string  `json:"modified_on"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            NotificationEmails []string  `json:"notification_emails"`
            Documents []Document  `json:"documents"`
            UID float64  `json:"uid"`
         
    }
    
    // ErrorResponse used by CompanyProfile
    type ErrorResponse struct {

        
            Code string  `json:"code"`
            Status float64  `json:"status"`
            Meta map[string]interface{}  `json:"meta"`
            Message string  `json:"message"`
         
    }
    
    // CompanyTaxesSerializer1 used by CompanyProfile
    type CompanyTaxesSerializer1 struct {

        
            EffectiveDate string  `json:"effective_date"`
            Rate float64  `json:"rate"`
            Enable bool  `json:"enable"`
         
    }
    
    // CreateUpdateAddressSerializer used by CompanyProfile
    type CreateUpdateAddressSerializer struct {

        
            Address1 string  `json:"address1"`
            Landmark string  `json:"landmark"`
            Address2 string  `json:"address2"`
            AddressType string  `json:"address_type"`
            Latitude float64  `json:"latitude"`
            Longitude float64  `json:"longitude"`
            Pincode float64  `json:"pincode"`
            Country string  `json:"country"`
            City string  `json:"city"`
            CountryCode string  `json:"country_code"`
            State string  `json:"state"`
         
    }
    
    // UpdateCompany used by CompanyProfile
    type UpdateCompany struct {

        
            BusinessType string  `json:"business_type"`
            Warnings map[string]interface{}  `json:"warnings"`
            FranchiseEnabled bool  `json:"franchise_enabled"`
            BusinessDetails BusinessDetails  `json:"business_details"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            NotificationEmails []string  `json:"notification_emails"`
            Taxes []CompanyTaxesSerializer1  `json:"taxes"`
            Addresses []CreateUpdateAddressSerializer  `json:"addresses"`
            BusinessInfo string  `json:"business_info"`
            Name string  `json:"name"`
            CompanyType string  `json:"company_type"`
            RejectReason string  `json:"reject_reason"`
            ContactDetails ContactDetails  `json:"contact_details"`
            Documents []Document  `json:"documents"`
         
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

        
            CompanyDocuments DocumentsObj  `json:"company_documents"`
            StoreDocuments DocumentsObj  `json:"store_documents"`
            Stage string  `json:"stage"`
            Product DocumentsObj  `json:"product"`
            Store DocumentsObj  `json:"store"`
            Brand DocumentsObj  `json:"brand"`
            UID float64  `json:"uid"`
         
    }
    
    // BrandBannerSerializer used by CompanyProfile
    type BrandBannerSerializer struct {

        
            Portrait string  `json:"portrait"`
            Landscape string  `json:"landscape"`
         
    }
    
    // GetBrandResponseSerializer used by CompanyProfile
    type GetBrandResponseSerializer struct {

        
            Warnings map[string]interface{}  `json:"warnings"`
            CreatedOn string  `json:"created_on"`
            Synonyms []string  `json:"synonyms"`
            ModifiedOn string  `json:"modified_on"`
            Stage string  `json:"stage"`
            SlugKey string  `json:"slug_key"`
            Logo string  `json:"logo"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            Banner BrandBannerSerializer  `json:"banner"`
            RejectReason string  `json:"reject_reason"`
            Mode string  `json:"mode"`
            Name string  `json:"name"`
            VerifiedOn string  `json:"verified_on"`
            Description string  `json:"description"`
            UID float64  `json:"uid"`
         
    }
    
    // CreateUpdateBrandRequestSerializer used by CompanyProfile
    type CreateUpdateBrandRequestSerializer struct {

        
            Synonyms []string  `json:"synonyms"`
            CompanyID float64  `json:"company_id"`
            Logo string  `json:"logo"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            Banner BrandBannerSerializer  `json:"banner"`
            Name string  `json:"name"`
            Description string  `json:"description"`
            UID float64  `json:"uid"`
            BrandTier string  `json:"brand_tier"`
         
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

        
            UID float64  `json:"uid"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            BusinessType string  `json:"business_type"`
            CreatedOn string  `json:"created_on"`
            Details CompanyDetails  `json:"details"`
            BusinessCountryInfo BusinessCountryInfo  `json:"business_country_info"`
            ModifiedOn string  `json:"modified_on"`
            Stage string  `json:"stage"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            NotificationEmails []string  `json:"notification_emails"`
            Name string  `json:"name"`
            MarketChannels []string  `json:"market_channels"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            RejectReason string  `json:"reject_reason"`
            CreatedBy UserSerializer  `json:"created_by"`
            CompanyType string  `json:"company_type"`
            VerifiedOn string  `json:"verified_on"`
         
    }
    
    // CompanyBrandSerializer used by CompanyProfile
    type CompanyBrandSerializer struct {

        
            UID float64  `json:"uid"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            Warnings map[string]interface{}  `json:"warnings"`
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
            Stage string  `json:"stage"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            RejectReason string  `json:"reject_reason"`
            Company CompanySerializer  `json:"company"`
            CreatedBy UserSerializer  `json:"created_by"`
            Brand GetBrandResponseSerializer  `json:"brand"`
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

        
            Brands []float64  `json:"brands"`
            UID float64  `json:"uid"`
            Company float64  `json:"company"`
         
    }
    
    // ProductReturnConfigSerializer used by CompanyProfile
    type ProductReturnConfigSerializer struct {

        
            OnSameStore bool  `json:"on_same_store"`
            StoreUID float64  `json:"store_uid"`
         
    }
    
    // HolidayDateSerializer used by CompanyProfile
    type HolidayDateSerializer struct {

        
            EndDate string  `json:"end_date"`
            StartDate string  `json:"start_date"`
         
    }
    
    // HolidaySchemaSerializer used by CompanyProfile
    type HolidaySchemaSerializer struct {

        
            HolidayType string  `json:"holiday_type"`
            Date HolidayDateSerializer  `json:"date"`
            Title string  `json:"title"`
         
    }
    
    // LocationTimingSerializer used by CompanyProfile
    type LocationTimingSerializer struct {

        
            Minute float64  `json:"minute"`
            Hour float64  `json:"hour"`
         
    }
    
    // LocationDayWiseSerializer used by CompanyProfile
    type LocationDayWiseSerializer struct {

        
            Weekday string  `json:"weekday"`
            Closing LocationTimingSerializer  `json:"closing"`
            Opening LocationTimingSerializer  `json:"opening"`
            Open bool  `json:"open"`
         
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
    
    // LocationManagerSerializer used by CompanyProfile
    type LocationManagerSerializer struct {

        
            Name string  `json:"name"`
            MobileNo SellerPhoneNumber  `json:"mobile_no"`
            Email string  `json:"email"`
         
    }
    
    // GetCompanySerializer used by CompanyProfile
    type GetCompanySerializer struct {

        
            UID float64  `json:"uid"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            BusinessType string  `json:"business_type"`
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
            Stage string  `json:"stage"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            Name string  `json:"name"`
            RejectReason string  `json:"reject_reason"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            CreatedBy UserSerializer  `json:"created_by"`
            CompanyType string  `json:"company_type"`
            VerifiedOn string  `json:"verified_on"`
         
    }
    
    // GetLocationSerializer used by CompanyProfile
    type GetLocationSerializer struct {

        
            CreatedOn string  `json:"created_on"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            ProductReturnConfig ProductReturnConfigSerializer  `json:"product_return_config"`
            Address GetAddressSerializer  `json:"address"`
            Name string  `json:"name"`
            Holiday []HolidaySchemaSerializer  `json:"holiday"`
            VerifiedOn string  `json:"verified_on"`
            DisplayName string  `json:"display_name"`
            Stage string  `json:"stage"`
            CreatedBy UserSerializer  `json:"created_by"`
            Timing []LocationDayWiseSerializer  `json:"timing"`
            GstCredentials InvoiceDetailsSerializer  `json:"gst_credentials"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            PhoneNumber string  `json:"phone_number"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            Warnings map[string]interface{}  `json:"warnings"`
            StoreType string  `json:"store_type"`
            ModifiedOn string  `json:"modified_on"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            NotificationEmails []string  `json:"notification_emails"`
            Manager LocationManagerSerializer  `json:"manager"`
            Company GetCompanySerializer  `json:"company"`
            Code string  `json:"code"`
            Documents []Document  `json:"documents"`
            UID float64  `json:"uid"`
         
    }
    
    // LocationListSerializer used by CompanyProfile
    type LocationListSerializer struct {

        
            Items []GetLocationSerializer  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // AddressSerializer used by CompanyProfile
    type AddressSerializer struct {

        
            Address1 string  `json:"address1"`
            Landmark string  `json:"landmark"`
            Address2 string  `json:"address2"`
            AddressType string  `json:"address_type"`
            Latitude float64  `json:"latitude"`
            Longitude float64  `json:"longitude"`
            Pincode float64  `json:"pincode"`
            Country string  `json:"country"`
            City string  `json:"city"`
            CountryCode string  `json:"country_code"`
            State string  `json:"state"`
         
    }
    
    // LocationSerializer used by CompanyProfile
    type LocationSerializer struct {

        
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            Warnings map[string]interface{}  `json:"warnings"`
            Holiday []HolidaySchemaSerializer  `json:"holiday"`
            StoreType string  `json:"store_type"`
            DisplayName string  `json:"display_name"`
            Stage string  `json:"stage"`
            ProductReturnConfig ProductReturnConfigSerializer  `json:"product_return_config"`
            Timing []LocationDayWiseSerializer  `json:"timing"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            NotificationEmails []string  `json:"notification_emails"`
            Manager LocationManagerSerializer  `json:"manager"`
            Address AddressSerializer  `json:"address"`
            Company float64  `json:"company"`
            Name string  `json:"name"`
            Code string  `json:"code"`
            Documents []Document  `json:"documents"`
            UID float64  `json:"uid"`
            GstCredentials InvoiceDetailsSerializer  `json:"gst_credentials"`
         
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
    

    
    // RuleSchema used by Cart
    type RuleSchema struct {

        
            Key float64  `json:"key"`
            Value float64  `json:"value"`
            Max float64  `json:"max"`
            Min float64  `json:"min"`
            DiscountQty float64  `json:"discount_qty"`
         
    }
    
    // IdentifierSchema used by Cart
    type IdentifierSchema struct {

        
            StoreID []float64  `json:"store_id"`
            CompanyID []float64  `json:"company_id"`
            CategoryID []float64  `json:"category_id"`
            BrandID []float64  `json:"brand_id"`
            ItemID []float64  `json:"item_id"`
            CollectionID []string  `json:"collection_id"`
            EmailDomain []string  `json:"email_domain"`
            ArticleID []string  `json:"article_id"`
            UserID []string  `json:"user_id"`
            ExcludeBrandID []float64  `json:"exclude_brand_id"`
         
    }
    
    // RuleDefinitionSchema used by Cart
    type RuleDefinitionSchema struct {

        
            CalculateOn string  `json:"calculate_on"`
            IsExact bool  `json:"is_exact"`
            ApplicableOn string  `json:"applicable_on"`
            CurrencyCode string  `json:"currency_code"`
            ValueType string  `json:"value_type"`
            Scope []string  `json:"scope"`
            Type string  `json:"type"`
            AutoApply bool  `json:"auto_apply"`
         
    }
    
    // ValiditySchema used by Cart
    type ValiditySchema struct {

        
            Priority float64  `json:"priority"`
         
    }
    
    // OwnershipSchema used by Cart
    type OwnershipSchema struct {

        
            PayableBy string  `json:"payable_by"`
            PayableCategory string  `json:"payable_category"`
         
    }
    
    // CouponAuthorSchema used by Cart
    type CouponAuthorSchema struct {

        
            CreatedBy string  `json:"created_by"`
            ModifiedBy string  `json:"modified_by"`
         
    }
    
    // UsesRemainingSchema used by Cart
    type UsesRemainingSchema struct {

        
            User float64  `json:"user"`
            App float64  `json:"app"`
            Total float64  `json:"total"`
         
    }
    
    // UsesRestrictionSchema used by Cart
    type UsesRestrictionSchema struct {

        
            Maximum UsesRemainingSchema  `json:"maximum"`
            Remaining UsesRemainingSchema  `json:"remaining"`
         
    }
    
    // PriceRangeSchema used by Cart
    type PriceRangeSchema struct {

        
            Max float64  `json:"max"`
            Min float64  `json:"min"`
         
    }
    
    // BulkBundleRestriction used by Cart
    type BulkBundleRestriction struct {

        
            MultiStoreAllowed bool  `json:"multi_store_allowed"`
         
    }
    
    // PostOrderSchema used by Cart
    type PostOrderSchema struct {

        
            CancellationAllowed bool  `json:"cancellation_allowed"`
            ReturnAllowed bool  `json:"return_allowed"`
         
    }
    
    // PaymentAllowValueSchema used by Cart
    type PaymentAllowValueSchema struct {

        
            Max float64  `json:"max"`
         
    }
    
    // PaymentModesSchema used by Cart
    type PaymentModesSchema struct {

        
            Types []string  `json:"types"`
            Uses PaymentAllowValueSchema  `json:"uses"`
            Networks []string  `json:"networks"`
            Codes []string  `json:"codes"`
         
    }
    
    // RestrictionsSchema used by Cart
    type RestrictionsSchema struct {

        
            UserType string  `json:"user_type"`
            Uses UsesRestrictionSchema  `json:"uses"`
            PriceRange PriceRangeSchema  `json:"price_range"`
            Platforms []string  `json:"platforms"`
            BulkBundle BulkBundleRestriction  `json:"bulk_bundle"`
            PostOrder PostOrderSchema  `json:"post_order"`
            CouponAllowed bool  `json:"coupon_allowed"`
            OrderingStores []float64  `json:"ordering_stores"`
            UserGroups []float64  `json:"user_groups"`
            Payments map[string]PaymentModesSchema  `json:"payments"`
         
    }
    
    // StateSchema used by Cart
    type StateSchema struct {

        
            IsPublic bool  `json:"is_public"`
            IsArchived bool  `json:"is_archived"`
            IsDisplay bool  `json:"is_display"`
         
    }
    
    // CouponActionSchema used by Cart
    type CouponActionSchema struct {

        
            ActionDate string  `json:"action_date"`
            TxnMode string  `json:"txn_mode"`
         
    }
    
    // ValidationSchema used by Cart
    type ValidationSchema struct {

        
            AppID []string  `json:"app_id"`
            UserRegisteredAfter string  `json:"user_registered_after"`
            Anonymous bool  `json:"anonymous"`
         
    }
    
    // CouponDateMetaSchema used by Cart
    type CouponDateMetaSchema struct {

        
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // CouponScheduleSchema used by Cart
    type CouponScheduleSchema struct {

        
            End string  `json:"end"`
            Start string  `json:"start"`
            NextSchedule []map[string]interface{}  `json:"next_schedule"`
            Duration float64  `json:"duration"`
            Cron string  `json:"cron"`
         
    }
    
    // DisplayMetaDictSchema used by Cart
    type DisplayMetaDictSchema struct {

        
            Title string  `json:"title"`
            Subtitle string  `json:"subtitle"`
         
    }
    
    // DisplayMetaSchema used by Cart
    type DisplayMetaSchema struct {

        
            Subtitle string  `json:"subtitle"`
            Auto DisplayMetaDictSchema  `json:"auto"`
            Title string  `json:"title"`
            Description string  `json:"description"`
            Apply DisplayMetaDictSchema  `json:"apply"`
            Remove DisplayMetaDictSchema  `json:"remove"`
         
    }
    
    // CouponAddSchema used by Cart
    type CouponAddSchema struct {

        
            Rule []RuleSchema  `json:"rule"`
            Identifiers IdentifierSchema  `json:"identifiers"`
            RuleDefinition RuleDefinitionSchema  `json:"rule_definition"`
            Validity ValiditySchema  `json:"validity"`
            Ownership OwnershipSchema  `json:"ownership"`
            Author CouponAuthorSchema  `json:"author"`
            Tags []string  `json:"tags"`
            Restrictions RestrictionsSchema  `json:"restrictions"`
            State StateSchema  `json:"state"`
            Action CouponActionSchema  `json:"action"`
            Validation ValidationSchema  `json:"validation"`
            TypeSlug string  `json:"type_slug"`
            DateMeta CouponDateMetaSchema  `json:"date_meta"`
            Code string  `json:"code"`
            Schedule CouponScheduleSchema  `json:"_schedule"`
            DisplayMeta DisplayMetaSchema  `json:"display_meta"`
         
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

        
            Items CouponAddSchema  `json:"items"`
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
    
    // CouponUpdateSchema used by Cart
    type CouponUpdateSchema struct {

        
            Rule []RuleSchema  `json:"rule"`
            Identifiers IdentifierSchema  `json:"identifiers"`
            RuleDefinition RuleDefinitionSchema  `json:"rule_definition"`
            Validity ValiditySchema  `json:"validity"`
            Ownership OwnershipSchema  `json:"ownership"`
            Author CouponAuthorSchema  `json:"author"`
            Tags []string  `json:"tags"`
            Restrictions RestrictionsSchema  `json:"restrictions"`
            State StateSchema  `json:"state"`
            Action CouponActionSchema  `json:"action"`
            Validation ValidationSchema  `json:"validation"`
            TypeSlug string  `json:"type_slug"`
            DateMeta CouponDateMetaSchema  `json:"date_meta"`
            Code string  `json:"code"`
            Schedule CouponScheduleSchema  `json:"_schedule"`
            DisplayMeta DisplayMetaSchema  `json:"display_meta"`
         
    }
    
    // CouponPartialUpdate used by Cart
    type CouponPartialUpdate struct {

        
            Archive bool  `json:"archive"`
            Schedule CouponScheduleSchema  `json:"schedule"`
         
    }
    
    // CompareObjectSchema used by Cart
    type CompareObjectSchema struct {

        
            GreaterThanEquals float64  `json:"greater_than_equals"`
            GreaterThan float64  `json:"greater_than"`
            LessThan float64  `json:"less_than"`
            Equals float64  `json:"equals"`
            LessThanEquals float64  `json:"less_than_equals"`
         
    }
    
    // ItemCriteriaSchema used by Cart
    type ItemCriteriaSchema struct {

        
            ItemL1Category []float64  `json:"item_l1_category"`
            ItemStore []float64  `json:"item_store"`
            ItemExcludeL2Category []float64  `json:"item_exclude_l2_category"`
            CartTotal CompareObjectSchema  `json:"cart_total"`
            ItemExcludeDepartment []float64  `json:"item_exclude_department"`
            BuyRules []string  `json:"buy_rules"`
            ItemExcludeL1Category []float64  `json:"item_exclude_l1_category"`
            ItemL2Category []float64  `json:"item_l2_category"`
            ItemExcludeCompany []float64  `json:"item_exclude_company"`
            ItemExcludeCategory []float64  `json:"item_exclude_category"`
            ItemTags []string  `json:"item_tags"`
            ItemSize []string  `json:"item_size"`
            CartUniqueItemAmount CompareObjectSchema  `json:"cart_unique_item_amount"`
            CartQuantity CompareObjectSchema  `json:"cart_quantity"`
            ItemCompany []float64  `json:"item_company"`
            ItemExcludeSku []string  `json:"item_exclude_sku"`
            AvailableZones []string  `json:"available_zones"`
            ItemCategory []float64  `json:"item_category"`
            ProductTags []string  `json:"product_tags"`
            CartUniqueItemQuantity CompareObjectSchema  `json:"cart_unique_item_quantity"`
            ItemExcludeBrand []float64  `json:"item_exclude_brand"`
            ItemID []float64  `json:"item_id"`
            ItemExcludeStore []float64  `json:"item_exclude_store"`
            ItemDepartment []float64  `json:"item_department"`
            AllItems bool  `json:"all_items"`
            ItemBrand []float64  `json:"item_brand"`
            ItemSku []string  `json:"item_sku"`
            ItemExcludeID []float64  `json:"item_exclude_id"`
         
    }
    
    // PromotionActionSchema used by Cart
    type PromotionActionSchema struct {

        
            ActionDate string  `json:"action_date"`
            ActionType string  `json:"action_type"`
         
    }
    
    // PromotionDateMetaSchema used by Cart
    type PromotionDateMetaSchema struct {

        
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // DiscountOfferSchema used by Cart
    type DiscountOfferSchema struct {

        
            ApportionDiscount bool  `json:"apportion_discount"`
            PartialCanRet bool  `json:"partial_can_ret"`
            MaxOfferQuantity float64  `json:"max_offer_quantity"`
            MaxDiscountAmount float64  `json:"max_discount_amount"`
            DiscountAmount float64  `json:"discount_amount"`
            MinOfferQuantity float64  `json:"min_offer_quantity"`
            MaxUsagePerTransaction float64  `json:"max_usage_per_transaction"`
            DiscountPrice float64  `json:"discount_price"`
            DiscountPercentage float64  `json:"discount_percentage"`
            Code string  `json:"code"`
         
    }
    
    // DiscountRuleSchema used by Cart
    type DiscountRuleSchema struct {

        
            ItemCriteria ItemCriteriaSchema  `json:"item_criteria"`
            DiscountType string  `json:"discount_type"`
            Offer DiscountOfferSchema  `json:"offer"`
            BuyCondition string  `json:"buy_condition"`
         
    }
    
    // DisplayMetaSchema1 used by Cart
    type DisplayMetaSchema1 struct {

        
            OfferLabel string  `json:"offer_label"`
            Description string  `json:"description"`
            OfferText string  `json:"offer_text"`
            Name string  `json:"name"`
         
    }
    
    // PromotionScheduleSchema used by Cart
    type PromotionScheduleSchema struct {

        
            End string  `json:"end"`
            Start string  `json:"start"`
            NextSchedule []map[string]interface{}  `json:"next_schedule"`
            Published bool  `json:"published"`
            Duration float64  `json:"duration"`
            Cron string  `json:"cron"`
         
    }
    
    // VisibilitySchema used by Cart
    type VisibilitySchema struct {

        
            CouponList bool  `json:"coupon_list"`
            Pdp bool  `json:"pdp"`
         
    }
    
    // PromotionAuthorSchema used by Cart
    type PromotionAuthorSchema struct {

        
            CreatedBy string  `json:"created_by"`
            ModifiedBy string  `json:"modified_by"`
         
    }
    
    // UserRegisteredSchema used by Cart
    type UserRegisteredSchema struct {

        
            Start string  `json:"start"`
            End string  `json:"end"`
         
    }
    
    // UsesRemainingSchema1 used by Cart
    type UsesRemainingSchema1 struct {

        
            User float64  `json:"user"`
            Total float64  `json:"total"`
         
    }
    
    // UsesRestrictionSchema1 used by Cart
    type UsesRestrictionSchema1 struct {

        
            Maximum UsesRemainingSchema1  `json:"maximum"`
            Remaining UsesRemainingSchema1  `json:"remaining"`
         
    }
    
    // PostOrderSchema1 used by Cart
    type PostOrderSchema1 struct {

        
            CancellationAllowed bool  `json:"cancellation_allowed"`
            ReturnAllowed bool  `json:"return_allowed"`
         
    }
    
    // PaymentAllowValueSchema1 used by Cart
    type PaymentAllowValueSchema1 struct {

        
            Max float64  `json:"max"`
         
    }
    
    // PromotionPaymentModesSchema used by Cart
    type PromotionPaymentModesSchema struct {

        
            Uses PaymentAllowValueSchema1  `json:"uses"`
            Type string  `json:"type"`
            Codes []string  `json:"codes"`
         
    }
    
    // RestrictionsSchema1 used by Cart
    type RestrictionsSchema1 struct {

        
            UserRegistered UserRegisteredSchema  `json:"user_registered"`
            Uses UsesRestrictionSchema1  `json:"uses"`
            Platforms []string  `json:"platforms"`
            PostOrder PostOrderSchema1  `json:"post_order"`
            AnonymousUsers bool  `json:"anonymous_users"`
            OrderQuantity float64  `json:"order_quantity"`
            UserGroups []float64  `json:"user_groups"`
            UserID []string  `json:"user_id"`
            Payments []PromotionPaymentModesSchema  `json:"payments"`
         
    }
    
    // OwnershipSchema1 used by Cart
    type OwnershipSchema1 struct {

        
            PayableBy string  `json:"payable_by"`
            PayableCategory string  `json:"payable_category"`
         
    }
    
    // PromotionListItem used by Cart
    type PromotionListItem struct {

        
            BuyRules map[string]ItemCriteriaSchema  `json:"buy_rules"`
            PostOrderAction PromotionActionSchema  `json:"post_order_action"`
            DateMeta PromotionDateMetaSchema  `json:"date_meta"`
            PromoGroup string  `json:"promo_group"`
            DiscountRules []DiscountRuleSchema  `json:"discount_rules"`
            ApplyPriority float64  `json:"apply_priority"`
            Code string  `json:"code"`
            ApplyExclusive string  `json:"apply_exclusive"`
            DisplayMeta DisplayMetaSchema1  `json:"display_meta"`
            Schedule PromotionScheduleSchema  `json:"_schedule"`
            Visiblility VisibilitySchema  `json:"visiblility"`
            Author PromotionAuthorSchema  `json:"author"`
            Stackable bool  `json:"stackable"`
            Restrictions RestrictionsSchema1  `json:"restrictions"`
            ApplyAllDiscount bool  `json:"apply_all_discount"`
            CalculateOn string  `json:"calculate_on"`
            Ownership OwnershipSchema1  `json:"ownership"`
            ApplicationID string  `json:"application_id"`
            Currency string  `json:"currency"`
            PromotionType string  `json:"promotion_type"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Mode string  `json:"mode"`
         
    }
    
    // PromotionsResponse used by Cart
    type PromotionsResponse struct {

        
            Items PromotionListItem  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // PromotionAddSchema used by Cart
    type PromotionAddSchema struct {

        
            BuyRules map[string]ItemCriteriaSchema  `json:"buy_rules"`
            PostOrderAction PromotionActionSchema  `json:"post_order_action"`
            DateMeta PromotionDateMetaSchema  `json:"date_meta"`
            PromoGroup string  `json:"promo_group"`
            DiscountRules []DiscountRuleSchema  `json:"discount_rules"`
            ApplyPriority float64  `json:"apply_priority"`
            Code string  `json:"code"`
            ApplyExclusive string  `json:"apply_exclusive"`
            DisplayMeta DisplayMetaSchema1  `json:"display_meta"`
            Schedule PromotionScheduleSchema  `json:"_schedule"`
            Visiblility VisibilitySchema  `json:"visiblility"`
            Author PromotionAuthorSchema  `json:"author"`
            Stackable bool  `json:"stackable"`
            Restrictions RestrictionsSchema1  `json:"restrictions"`
            ApplyAllDiscount bool  `json:"apply_all_discount"`
            CalculateOn string  `json:"calculate_on"`
            Ownership OwnershipSchema1  `json:"ownership"`
            ApplicationID string  `json:"application_id"`
            Currency string  `json:"currency"`
            PromotionType string  `json:"promotion_type"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Mode string  `json:"mode"`
         
    }
    
    // PromotionUpdateSchema used by Cart
    type PromotionUpdateSchema struct {

        
            BuyRules map[string]ItemCriteriaSchema  `json:"buy_rules"`
            PostOrderAction PromotionActionSchema  `json:"post_order_action"`
            DateMeta PromotionDateMetaSchema  `json:"date_meta"`
            PromoGroup string  `json:"promo_group"`
            DiscountRules []DiscountRuleSchema  `json:"discount_rules"`
            ApplyPriority float64  `json:"apply_priority"`
            Code string  `json:"code"`
            ApplyExclusive string  `json:"apply_exclusive"`
            DisplayMeta DisplayMetaSchema1  `json:"display_meta"`
            Schedule PromotionScheduleSchema  `json:"_schedule"`
            Visiblility VisibilitySchema  `json:"visiblility"`
            Author PromotionAuthorSchema  `json:"author"`
            Stackable bool  `json:"stackable"`
            Restrictions RestrictionsSchema1  `json:"restrictions"`
            ApplyAllDiscount bool  `json:"apply_all_discount"`
            CalculateOn string  `json:"calculate_on"`
            Ownership OwnershipSchema1  `json:"ownership"`
            ApplicationID string  `json:"application_id"`
            Currency string  `json:"currency"`
            PromotionType string  `json:"promotion_type"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Mode string  `json:"mode"`
         
    }
    
    // PromotionPartialUpdate used by Cart
    type PromotionPartialUpdate struct {

        
            Archive bool  `json:"archive"`
            Schedule PromotionScheduleSchema  `json:"schedule"`
         
    }
    
    // ActivePromosResponse used by Cart
    type ActivePromosResponse struct {

        
            ModifiedOn string  `json:"modified_on"`
            Subtitle string  `json:"subtitle"`
            IsHidden bool  `json:"is_hidden"`
            Title string  `json:"title"`
            Description string  `json:"description"`
            EntitySlug string  `json:"entity_slug"`
            Example string  `json:"example"`
            CreatedOn string  `json:"created_on"`
            Type string  `json:"type"`
            EntityType string  `json:"entity_type"`
         
    }
    
    // CartItemSchema used by Cart
    type CartItemSchema struct {

        
            Quantity float64  `json:"quantity"`
            Size string  `json:"size"`
            ProductID string  `json:"product_id"`
         
    }
    
    // OpenapiCartDetailsRequest used by Cart
    type OpenapiCartDetailsRequest struct {

        
            CartItems []CartItemSchema  `json:"cart_items"`
         
    }
    
    // PromoMeta used by Cart
    type PromoMeta struct {

        
            Message string  `json:"message"`
         
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
    
    // FreeGiftItemSchema used by Cart
    type FreeGiftItemSchema struct {

        
            ItemPriceDetails map[string]interface{}  `json:"item_price_details"`
            ItemID float64  `json:"item_id"`
            ItemSlug string  `json:"item_slug"`
            ItemName string  `json:"item_name"`
            ItemBrandName string  `json:"item_brand_name"`
            ItemImagesURL []string  `json:"item_images_url"`
         
    }
    
    // AppliedFreeArticlesSchema used by Cart
    type AppliedFreeArticlesSchema struct {

        
            FreeGiftItemDetails FreeGiftItemSchema  `json:"free_gift_item_details"`
            Quantity float64  `json:"quantity"`
            ArticleID string  `json:"article_id"`
            ParentItemIdentifier string  `json:"parent_item_identifier"`
         
    }
    
    // DiscountRulesAppSchema used by Cart
    type DiscountRulesAppSchema struct {

        
            ItemCriteria map[string]interface{}  `json:"item_criteria"`
            MatchedBuyRules []string  `json:"matched_buy_rules"`
            RawOffer map[string]interface{}  `json:"raw_offer"`
            Offer map[string]interface{}  `json:"offer"`
         
    }
    
    // BuyRulesSchema used by Cart
    type BuyRulesSchema struct {

        
            ItemCriteria map[string]interface{}  `json:"item_criteria"`
            CartConditions map[string]interface{}  `json:"cart_conditions"`
         
    }
    
    // Ownership used by Cart
    type Ownership struct {

        
            PayableBy string  `json:"payable_by"`
            PayableCategory string  `json:"payable_category"`
         
    }
    
    // AppliedPromotion used by Cart
    type AppliedPromotion struct {

        
            MrpPromotion bool  `json:"mrp_promotion"`
            AppliedFreeArticles []AppliedFreeArticlesSchema  `json:"applied_free_articles"`
            DiscountRules []DiscountRulesAppSchema  `json:"discount_rules"`
            BuyRules []BuyRulesSchema  `json:"buy_rules"`
            Ownership Ownership  `json:"ownership"`
            Amount float64  `json:"amount"`
            PromotionGroup string  `json:"promotion_group"`
            OfferText string  `json:"offer_text"`
            PromotionType string  `json:"promotion_type"`
            PromoID string  `json:"promo_id"`
            PromotionName string  `json:"promotion_name"`
            ArticleQuantity float64  `json:"article_quantity"`
         
    }
    
    // ProductPrice used by Cart
    type ProductPrice struct {

        
            AddOn float64  `json:"add_on"`
            CurrencySymbol string  `json:"currency_symbol"`
            Effective float64  `json:"effective"`
            Marked float64  `json:"marked"`
            CurrencyCode string  `json:"currency_code"`
            Selling float64  `json:"selling"`
         
    }
    
    // ProductPriceInfo used by Cart
    type ProductPriceInfo struct {

        
            Converted ProductPrice  `json:"converted"`
            Base ProductPrice  `json:"base"`
         
    }
    
    // BaseInfo used by Cart
    type BaseInfo struct {

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // NetQuantity used by Cart
    type NetQuantity struct {

        
            Value string  `json:"value"`
            Unit string  `json:"unit"`
         
    }
    
    // ProductImage used by Cart
    type ProductImage struct {

        
            SecureURL string  `json:"secure_url"`
            AspectRatio string  `json:"aspect_ratio"`
            URL string  `json:"url"`
         
    }
    
    // Tags used by Cart
    type Tags struct {

        
            Tags map[string]interface{}  `json:"tags"`
         
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
    
    // CategoryInfo used by Cart
    type CategoryInfo struct {

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // CartProduct used by Cart
    type CartProduct struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Brand BaseInfo  `json:"brand"`
            NetQuantity NetQuantity  `json:"net_quantity"`
            ItemCode string  `json:"item_code"`
            Images []ProductImage  `json:"images"`
            Tags []string  `json:"tags"`
            TeaserTag Tags  `json:"teaser_tag"`
            Action ProductAction  `json:"action"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
            Type string  `json:"type"`
            Categories []CategoryInfo  `json:"categories"`
         
    }
    
    // CouponDetails used by Cart
    type CouponDetails struct {

        
            DiscountTotalQuantity float64  `json:"discount_total_quantity"`
            Code string  `json:"code"`
            DiscountSingleQuantity float64  `json:"discount_single_quantity"`
         
    }
    
    // CartProductIdentifer used by Cart
    type CartProductIdentifer struct {

        
            Identifier string  `json:"identifier"`
         
    }
    
    // BasePriceSchema used by Cart
    type BasePriceSchema struct {

        
            Marked float64  `json:"marked"`
            CurrencyCode string  `json:"currency_code"`
            CurrencySymbol string  `json:"currency_symbol"`
            Effective float64  `json:"effective"`
         
    }
    
    // ArticlePriceInfo used by Cart
    type ArticlePriceInfo struct {

        
            Converted BasePriceSchema  `json:"converted"`
            Base BasePriceSchema  `json:"base"`
         
    }
    
    // ProductArticle used by Cart
    type ProductArticle struct {

        
            Quantity float64  `json:"quantity"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            GiftCard map[string]interface{}  `json:"gift_card"`
            IsGiftVisible bool  `json:"is_gift_visible"`
            Identifier map[string]interface{}  `json:"identifier"`
            Price ArticlePriceInfo  `json:"price"`
            ProductGroupTags []string  `json:"product_group_tags"`
            Store BaseInfo  `json:"store"`
            SellerIdentifier string  `json:"seller_identifier"`
            CartItemMeta map[string]interface{}  `json:"cart_item_meta"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            UID string  `json:"uid"`
            Type string  `json:"type"`
            Size string  `json:"size"`
            Seller BaseInfo  `json:"seller"`
         
    }
    
    // ProductAvailabilitySize used by Cart
    type ProductAvailabilitySize struct {

        
            Display string  `json:"display"`
            IsAvailable bool  `json:"is_available"`
            Value string  `json:"value"`
         
    }
    
    // ProductAvailability used by Cart
    type ProductAvailability struct {

        
            Deliverable bool  `json:"deliverable"`
            AvailableSizes []ProductAvailabilitySize  `json:"available_sizes"`
            Sizes []string  `json:"sizes"`
            OtherStoreQuantity float64  `json:"other_store_quantity"`
            IsValid bool  `json:"is_valid"`
            OutOfStock bool  `json:"out_of_stock"`
         
    }
    
    // CartProductInfo used by Cart
    type CartProductInfo struct {

        
            Quantity float64  `json:"quantity"`
            PromoMeta PromoMeta  `json:"promo_meta"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            PromotionsApplied []AppliedPromotion  `json:"promotions_applied"`
            Price ProductPriceInfo  `json:"price"`
            Moq map[string]interface{}  `json:"moq"`
            IsSet bool  `json:"is_set"`
            CouponMessage string  `json:"coupon_message"`
            BulkOffer map[string]interface{}  `json:"bulk_offer"`
            Key string  `json:"key"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            Message string  `json:"message"`
            Product CartProduct  `json:"product"`
            Coupon CouponDetails  `json:"coupon"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            PricePerUnit ProductPriceInfo  `json:"price_per_unit"`
            Discount string  `json:"discount"`
            Article ProductArticle  `json:"article"`
            Availability ProductAvailability  `json:"availability"`
         
    }
    
    // RawBreakupSchema used by Cart
    type RawBreakupSchema struct {

        
            GstCharges float64  `json:"gst_charges"`
            GiftCard float64  `json:"gift_card"`
            Total float64  `json:"total"`
            Coupon float64  `json:"coupon"`
            YouSaved float64  `json:"you_saved"`
            Discount float64  `json:"discount"`
            MrpTotal float64  `json:"mrp_total"`
            DeliveryCharge float64  `json:"delivery_charge"`
            ConvenienceFee float64  `json:"convenience_fee"`
            CodCharge float64  `json:"cod_charge"`
            FyndCash float64  `json:"fynd_cash"`
            Subtotal float64  `json:"subtotal"`
            Vog float64  `json:"vog"`
         
    }
    
    // LoyaltyPoints used by Cart
    type LoyaltyPoints struct {

        
            IsApplied bool  `json:"is_applied"`
            Description string  `json:"description"`
            Total float64  `json:"total"`
            Applicable float64  `json:"applicable"`
         
    }
    
    // DisplayBreakupSchema used by Cart
    type DisplayBreakupSchema struct {

        
            Key string  `json:"key"`
            Display string  `json:"display"`
            Value float64  `json:"value"`
            CurrencySymbol string  `json:"currency_symbol"`
            Message []string  `json:"message"`
            CurrencyCode string  `json:"currency_code"`
         
    }
    
    // CouponBreakupSchema used by Cart
    type CouponBreakupSchema struct {

        
            Value float64  `json:"value"`
            CouponType string  `json:"coupon_type"`
            IsApplied bool  `json:"is_applied"`
            Message string  `json:"message"`
            Title string  `json:"title"`
            Description string  `json:"description"`
            SubTitle string  `json:"sub_title"`
            MaxDiscountValue float64  `json:"max_discount_value"`
            MinimumCartValue float64  `json:"minimum_cart_value"`
            CouponValue float64  `json:"coupon_value"`
            UID string  `json:"uid"`
            Type string  `json:"type"`
            Code string  `json:"code"`
         
    }
    
    // CartBreakupSchema used by Cart
    type CartBreakupSchema struct {

        
            Raw RawBreakupSchema  `json:"raw"`
            LoyaltyPoints LoyaltyPoints  `json:"loyalty_points"`
            Display []DisplayBreakupSchema  `json:"display"`
            Coupon CouponBreakupSchema  `json:"coupon"`
         
    }
    
    // OpenapiCartDetailsResponse used by Cart
    type OpenapiCartDetailsResponse struct {

        
            Message string  `json:"message"`
            Items []CartProductInfo  `json:"items"`
            BreakupValues CartBreakupSchema  `json:"breakup_values"`
            IsValid bool  `json:"is_valid"`
         
    }
    
    // OpenApiErrorResponse used by Cart
    type OpenApiErrorResponse struct {

        
            Message string  `json:"message"`
            Errors map[string]interface{}  `json:"errors"`
            Success bool  `json:"success"`
         
    }
    
    // ShippingAddressSchema used by Cart
    type ShippingAddressSchema struct {

        
            Address string  `json:"address"`
            Pincode float64  `json:"pincode"`
            Meta map[string]interface{}  `json:"meta"`
            CountryCode string  `json:"country_code"`
            Landmark string  `json:"landmark"`
            Area string  `json:"area"`
            AddressType string  `json:"address_type"`
            Phone float64  `json:"phone"`
            CountryIsoCode string  `json:"country_iso_code"`
            AreaCodeSlug string  `json:"area_code_slug"`
            CountryPhoneCode string  `json:"country_phone_code"`
            City string  `json:"city"`
            State string  `json:"state"`
            Name string  `json:"name"`
            Country string  `json:"country"`
            Email string  `json:"email"`
            AreaCode string  `json:"area_code"`
         
    }
    
    // OpenApiCartServiceabilityRequest used by Cart
    type OpenApiCartServiceabilityRequest struct {

        
            CartItems []CartItemSchema  `json:"cart_items"`
            ShippingAddress ShippingAddressSchema  `json:"shipping_address"`
         
    }
    
    // OpenApiCartServiceabilityResponse used by Cart
    type OpenApiCartServiceabilityResponse struct {

        
            BreakupValues CartBreakupSchema  `json:"breakup_values"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            Message string  `json:"message"`
            Items []CartProductInfo  `json:"items"`
            IsValid bool  `json:"is_valid"`
         
    }
    
    // OpenApiFilesSchema used by Cart
    type OpenApiFilesSchema struct {

        
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

        
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            OrderID string  `json:"order_id"`
            PaymentGateway string  `json:"payment_gateway"`
            PaymentID string  `json:"payment_id"`
            CurrentStatus string  `json:"current_status"`
         
    }
    
    // MultiTenderPaymentMethod used by Cart
    type MultiTenderPaymentMethod struct {

        
            Amount float64  `json:"amount"`
            Meta MultiTenderPaymentMeta  `json:"meta"`
            Mode string  `json:"mode"`
            Name string  `json:"name"`
         
    }
    
    // OpenApiOrderItemSchema used by Cart
    type OpenApiOrderItemSchema struct {

        
            Quantity float64  `json:"quantity"`
            EmployeeDiscount float64  `json:"employee_discount"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            Meta CartItemMeta  `json:"meta"`
            PriceEffective float64  `json:"price_effective"`
            Files []OpenApiFilesSchema  `json:"files"`
            ProductID float64  `json:"product_id"`
            CashbackApplied float64  `json:"cashback_applied"`
            Discount float64  `json:"discount"`
            AmountPaid float64  `json:"amount_paid"`
            PriceMarked float64  `json:"price_marked"`
            PaymentMethods []MultiTenderPaymentMethod  `json:"payment_methods"`
            DeliveryCharges float64  `json:"delivery_charges"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            Size string  `json:"size"`
            CodCharges float64  `json:"cod_charges"`
         
    }
    
    // OpenApiPlatformCheckoutReqSchema used by Cart
    type OpenApiPlatformCheckoutReqSchema struct {

        
            AffiliateOrderID string  `json:"affiliate_order_id"`
            Files []OpenApiFilesSchema  `json:"files"`
            Coupon string  `json:"coupon"`
            CashbackApplied float64  `json:"cashback_applied"`
            DeliveryCharges float64  `json:"delivery_charges"`
            CurrencyCode string  `json:"currency_code"`
            CartItems []OpenApiOrderItemSchema  `json:"cart_items"`
            CartValue float64  `json:"cart_value"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            EmployeeDiscount map[string]interface{}  `json:"employee_discount"`
            OrderID string  `json:"order_id"`
            CouponCode string  `json:"coupon_code"`
            CodCharges float64  `json:"cod_charges"`
            ShippingAddress ShippingAddressSchema  `json:"shipping_address"`
            CouponValue float64  `json:"coupon_value"`
            Comment string  `json:"comment"`
            Gstin string  `json:"gstin"`
            PaymentMode string  `json:"payment_mode"`
            BillingAddress ShippingAddressSchema  `json:"billing_address"`
            PaymentMethods []MultiTenderPaymentMethod  `json:"payment_methods"`
         
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

        
            LastModified string  `json:"last_modified"`
            Articles []map[string]interface{}  `json:"articles"`
            ExpireAt string  `json:"expire_at"`
            PaymentMethods []map[string]interface{}  `json:"payment_methods"`
            Cashback map[string]interface{}  `json:"cashback"`
            CheckoutMode string  `json:"checkout_mode"`
            ID string  `json:"_id"`
            UID float64  `json:"uid"`
            FyndCredits map[string]interface{}  `json:"fynd_credits"`
            BulkCouponDiscount float64  `json:"bulk_coupon_discount"`
            Gstin string  `json:"gstin"`
            PaymentMode string  `json:"payment_mode"`
            UserID string  `json:"user_id"`
            IsDefault bool  `json:"is_default"`
            IsActive bool  `json:"is_active"`
            CartValue float64  `json:"cart_value"`
            Coupon map[string]interface{}  `json:"coupon"`
            Shipments []map[string]interface{}  `json:"shipments"`
            BuyNow bool  `json:"buy_now"`
            Promotion map[string]interface{}  `json:"promotion"`
            PickUpCustomerDetails map[string]interface{}  `json:"pick_up_customer_details"`
            DeliveryCharges map[string]interface{}  `json:"delivery_charges"`
            FcIndexMap []float64  `json:"fc_index_map"`
            Payments map[string]interface{}  `json:"payments"`
            AppID string  `json:"app_id"`
            Meta map[string]interface{}  `json:"meta"`
            Discount float64  `json:"discount"`
            OrderID string  `json:"order_id"`
            CreatedOn string  `json:"created_on"`
            Comment string  `json:"comment"`
            IsArchive bool  `json:"is_archive"`
            MergeQty bool  `json:"merge_qty"`
            CodCharges map[string]interface{}  `json:"cod_charges"`
         
    }
    
    // AbandonedCartResponseSchema used by Cart
    type AbandonedCartResponseSchema struct {

        
            Result map[string]interface{}  `json:"result"`
            Success bool  `json:"success"`
            Page Page  `json:"page"`
            Message string  `json:"message"`
            Items []AbandonedCart  `json:"items"`
         
    }
    
    // PaymentSelectionLock used by Cart
    type PaymentSelectionLock struct {

        
            Enabled bool  `json:"enabled"`
            PaymentIdentifier string  `json:"payment_identifier"`
            DefaultOptions string  `json:"default_options"`
         
    }
    
    // CartCurrencySchema used by Cart
    type CartCurrencySchema struct {

        
            Symbol string  `json:"symbol"`
            Code string  `json:"code"`
         
    }
    
    // CartDetailResponse used by Cart
    type CartDetailResponse struct {

        
            ID string  `json:"id"`
            LastModified string  `json:"last_modified"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            IsValid bool  `json:"is_valid"`
            CheckoutMode string  `json:"checkout_mode"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            BreakupValues CartBreakupSchema  `json:"breakup_values"`
            Message string  `json:"message"`
            Gstin string  `json:"gstin"`
            AppliedPromoDetails []AppliedPromotion  `json:"applied_promo_details"`
            BuyNow bool  `json:"buy_now"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            PanNo string  `json:"pan_no"`
            PanConfig map[string]interface{}  `json:"pan_config"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            Items []CartProductInfo  `json:"items"`
            Comment string  `json:"comment"`
            CouponText string  `json:"coupon_text"`
            Currency CartCurrencySchema  `json:"currency"`
         
    }
    
    // AddProductCart used by Cart
    type AddProductCart struct {

        
            Quantity float64  `json:"quantity"`
            ParentItemIdentifiers []map[string]string  `json:"parent_item_identifiers"`
            StoreID float64  `json:"store_id"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            ArticleAssignment map[string]interface{}  `json:"article_assignment"`
            ItemID float64  `json:"item_id"`
            Display string  `json:"display"`
            ProductGroupTags []string  `json:"product_group_tags"`
            SellerID float64  `json:"seller_id"`
            Pos bool  `json:"pos"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ArticleID string  `json:"article_id"`
            ItemSize string  `json:"item_size"`
         
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

        
            Quantity float64  `json:"quantity"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
            ItemID float64  `json:"item_id"`
            ItemIndex float64  `json:"item_index"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ArticleID string  `json:"article_id"`
            ItemSize string  `json:"item_size"`
         
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

        
            Source map[string]interface{}  `json:"source"`
            Meta map[string]interface{}  `json:"meta"`
            Token string  `json:"token"`
            CreatedOn string  `json:"created_on"`
            User map[string]interface{}  `json:"user"`
         
    }
    
    // SharedCart used by Cart
    type SharedCart struct {

        
            ID string  `json:"id"`
            LastModified string  `json:"last_modified"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            IsValid bool  `json:"is_valid"`
            CartID float64  `json:"cart_id"`
            CheckoutMode string  `json:"checkout_mode"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            BreakupValues CartBreakupSchema  `json:"breakup_values"`
            Message string  `json:"message"`
            Gstin string  `json:"gstin"`
            BuyNow bool  `json:"buy_now"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            SharedCartDetails SharedCartDetails  `json:"shared_cart_details"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            Items []CartProductInfo  `json:"items"`
            Comment string  `json:"comment"`
            CouponText string  `json:"coupon_text"`
            Currency CartCurrencySchema  `json:"currency"`
            UID string  `json:"uid"`
         
    }
    
    // SharedCartResponse used by Cart
    type SharedCartResponse struct {

        
            Error string  `json:"error"`
            Cart SharedCart  `json:"cart"`
         
    }
    
    // CartListSchema used by Cart
    type CartListSchema struct {

        
            ItemCounts float64  `json:"item_counts"`
            PickUpCustomerDetails map[string]interface{}  `json:"pick_up_customer_details"`
            CreatedOn string  `json:"created_on"`
            CartID string  `json:"cart_id"`
            UserID string  `json:"user_id"`
            CartValue float64  `json:"cart_value"`
         
    }
    
    // MultiCartResponseSchema used by Cart
    type MultiCartResponseSchema struct {

        
            Data []CartListSchema  `json:"data"`
            Success bool  `json:"success"`
         
    }
    
    // UpdateUserCartMapping used by Cart
    type UpdateUserCartMapping struct {

        
            UserID string  `json:"user_id"`
         
    }
    
    // UserInfoSchema used by Cart
    type UserInfoSchema struct {

        
            ExternalID string  `json:"external_id"`
            FirstName string  `json:"first_name"`
            ModifiedOn string  `json:"modified_on"`
            Mobile string  `json:"mobile"`
            LastName string  `json:"last_name"`
            CreatedAt string  `json:"created_at"`
            UID string  `json:"uid"`
            Gender string  `json:"gender"`
            ID string  `json:"_id"`
         
    }
    
    // UserCartMappingResponse used by Cart
    type UserCartMappingResponse struct {

        
            ID string  `json:"id"`
            LastModified string  `json:"last_modified"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            IsValid bool  `json:"is_valid"`
            CheckoutMode string  `json:"checkout_mode"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            BreakupValues CartBreakupSchema  `json:"breakup_values"`
            Message string  `json:"message"`
            Gstin string  `json:"gstin"`
            User UserInfoSchema  `json:"user"`
            AppliedPromoDetails []AppliedPromotion  `json:"applied_promo_details"`
            BuyNow bool  `json:"buy_now"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            PanNo string  `json:"pan_no"`
            PanConfig map[string]interface{}  `json:"pan_config"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            Items []CartProductInfo  `json:"items"`
            Comment string  `json:"comment"`
            CouponText string  `json:"coupon_text"`
            Currency CartCurrencySchema  `json:"currency"`
         
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

        
            IsApplicable bool  `json:"is_applicable"`
            CouponCode string  `json:"coupon_code"`
            CouponType string  `json:"coupon_type"`
            IsApplied bool  `json:"is_applied"`
            Message string  `json:"message"`
            Title string  `json:"title"`
            CouponValue float64  `json:"coupon_value"`
            MinimumCartValue float64  `json:"minimum_cart_value"`
            Description string  `json:"description"`
            SubTitle string  `json:"sub_title"`
            MaxDiscountValue float64  `json:"max_discount_value"`
            ExpiresOn string  `json:"expires_on"`
         
    }
    
    // PageCouponSchema used by Cart
    type PageCouponSchema struct {

        
            Total float64  `json:"total"`
            TotalItemCount float64  `json:"total_item_count"`
            HasNext bool  `json:"has_next"`
            HasPrevious bool  `json:"has_previous"`
            Current float64  `json:"current"`
         
    }
    
    // GetCouponResponse used by Cart
    type GetCouponResponse struct {

        
            AvailableCouponList []Coupon  `json:"available_coupon_list"`
            Page PageCouponSchema  `json:"page"`
         
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

        
            ID string  `json:"id"`
            CartID string  `json:"cart_id"`
            Country string  `json:"country"`
            Email string  `json:"email"`
            CheckoutMode string  `json:"checkout_mode"`
            Landmark string  `json:"landmark"`
            City string  `json:"city"`
            AddressType string  `json:"address_type"`
            CountryCode string  `json:"country_code"`
            IsDefaultAddress bool  `json:"is_default_address"`
            CreatedByUserID string  `json:"created_by_user_id"`
            Tags []string  `json:"tags"`
            UserID string  `json:"user_id"`
            AreaCode string  `json:"area_code"`
            IsActive bool  `json:"is_active"`
            AreaCodeSlug string  `json:"area_code_slug"`
            GeoLocation GeoLocation  `json:"geo_location"`
            State string  `json:"state"`
            Name string  `json:"name"`
            Address string  `json:"address"`
            Meta map[string]interface{}  `json:"meta"`
            Area string  `json:"area"`
            Phone string  `json:"phone"`
            GoogleMapPoint map[string]interface{}  `json:"google_map_point"`
         
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

        
            ID string  `json:"id"`
            BillingAddressID string  `json:"billing_address_id"`
            CartID string  `json:"cart_id"`
            UserID string  `json:"user_id"`
            CheckoutMode string  `json:"checkout_mode"`
         
    }
    
    // ShipmentResponse used by Cart
    type ShipmentResponse struct {

        
            DpOptions map[string]interface{}  `json:"dp_options"`
            Shipments float64  `json:"shipments"`
            FulfillmentType string  `json:"fulfillment_type"`
            Promise ShipmentPromise  `json:"promise"`
            FulfillmentID float64  `json:"fulfillment_id"`
            Items []CartProductInfo  `json:"items"`
            OrderType string  `json:"order_type"`
            ShipmentType string  `json:"shipment_type"`
            DpID string  `json:"dp_id"`
            BoxType string  `json:"box_type"`
         
    }
    
    // CartShipmentsResponse used by Cart
    type CartShipmentsResponse struct {

        
            ID string  `json:"id"`
            LastModified string  `json:"last_modified"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            BuyNow bool  `json:"buy_now"`
            Shipments []ShipmentResponse  `json:"shipments"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            Error bool  `json:"error"`
            IsValid bool  `json:"is_valid"`
            CartID float64  `json:"cart_id"`
            CheckoutMode string  `json:"checkout_mode"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            BreakupValues CartBreakupSchema  `json:"breakup_values"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            Message string  `json:"message"`
            Comment string  `json:"comment"`
            CouponText string  `json:"coupon_text"`
            Currency CartCurrencySchema  `json:"currency"`
            Gstin string  `json:"gstin"`
            UID string  `json:"uid"`
         
    }
    
    // UpdateCartShipmentItem used by Cart
    type UpdateCartShipmentItem struct {

        
            ArticleUID string  `json:"article_uid"`
            Quantity float64  `json:"quantity"`
            ShipmentType string  `json:"shipment_type"`
         
    }
    
    // UpdateCartShipmentRequest used by Cart
    type UpdateCartShipmentRequest struct {

        
            Shipments []UpdateCartShipmentItem  `json:"shipments"`
         
    }
    
    // GiftDetailSchema used by Cart
    type GiftDetailSchema struct {

        
            IsGiftApplied bool  `json:"is_gift_applied"`
            GiftMessage string  `json:"gift_message"`
         
    }
    
    // ArticleGiftDetailSchema used by Cart
    type ArticleGiftDetailSchema struct {

        
            ArticleID GiftDetailSchema  `json:"article_id"`
         
    }
    
    // PlatformCartMetaRequest used by Cart
    type PlatformCartMetaRequest struct {

        
            PickUpCustomerDetails map[string]interface{}  `json:"pick_up_customer_details"`
            PanNo string  `json:"pan_no"`
            Comment string  `json:"comment"`
            GiftDetails ArticleGiftDetailSchema  `json:"gift_details"`
            Gstin string  `json:"gstin"`
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
    
    // FilesSchema used by Cart
    type FilesSchema struct {

        
            Key string  `json:"key"`
            Values []string  `json:"values"`
         
    }
    
    // StaffCheckoutSchema used by Cart
    type StaffCheckoutSchema struct {

        
            FirstName string  `json:"first_name"`
            EmployeeCode string  `json:"employee_code"`
            LastName string  `json:"last_name"`
            User string  `json:"user"`
            ID string  `json:"_id"`
         
    }
    
    // PlatformCartCheckoutDetailRequest used by Cart
    type PlatformCartCheckoutDetailRequest struct {

        
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            ID string  `json:"id"`
            CallbackURL string  `json:"callback_url"`
            BillingAddressID string  `json:"billing_address_id"`
            Pos bool  `json:"pos"`
            PaymentIdentifier string  `json:"payment_identifier"`
            Aggregator string  `json:"aggregator"`
            CheckoutMode string  `json:"checkout_mode"`
            PaymentParams map[string]interface{}  `json:"payment_params"`
            EmployeeCode string  `json:"employee_code"`
            PaymentMode string  `json:"payment_mode"`
            UserID string  `json:"user_id"`
            PickAtStoreUID float64  `json:"pick_at_store_uid"`
            Files []FilesSchema  `json:"files"`
            DeviceID string  `json:"device_id"`
            OrderingStore float64  `json:"ordering_store"`
            PaymentAutoConfirm bool  `json:"payment_auto_confirm"`
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
            Staff StaffCheckoutSchema  `json:"staff"`
            MerchantCode string  `json:"merchant_code"`
            Meta map[string]interface{}  `json:"meta"`
            AddressID string  `json:"address_id"`
            OrderType string  `json:"order_type"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
         
    }
    
    // CheckCartSchema used by Cart
    type CheckCartSchema struct {

        
            ID string  `json:"id"`
            LastModified string  `json:"last_modified"`
            CodMessage string  `json:"cod_message"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            DeliveryChargeOrderValue float64  `json:"delivery_charge_order_value"`
            IsValid bool  `json:"is_valid"`
            CartID float64  `json:"cart_id"`
            StoreCode string  `json:"store_code"`
            CheckoutMode string  `json:"checkout_mode"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            BreakupValues CartBreakupSchema  `json:"breakup_values"`
            Message string  `json:"message"`
            Gstin string  `json:"gstin"`
            UserType string  `json:"user_type"`
            BuyNow bool  `json:"buy_now"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            Success bool  `json:"success"`
            DeliveryCharges float64  `json:"delivery_charges"`
            ErrorMessage string  `json:"error_message"`
            StoreEmps []map[string]interface{}  `json:"store_emps"`
            OrderID string  `json:"order_id"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            Items []CartProductInfo  `json:"items"`
            Comment string  `json:"comment"`
            CouponText string  `json:"coupon_text"`
            Currency CartCurrencySchema  `json:"currency"`
            UID string  `json:"uid"`
            CodAvailable bool  `json:"cod_available"`
            CodCharges float64  `json:"cod_charges"`
         
    }
    
    // CartCheckoutResponseSchema used by Cart
    type CartCheckoutResponseSchema struct {

        
            Data map[string]interface{}  `json:"data"`
            CallbackURL string  `json:"callback_url"`
            OrderID string  `json:"order_id"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
            PaymentConfirmURL string  `json:"payment_confirm_url"`
            Cart CheckCartSchema  `json:"cart"`
            AppInterceptURL string  `json:"app_intercept_url"`
         
    }
    
    // CartDeliveryModesResponse used by Cart
    type CartDeliveryModesResponse struct {

        
            PickupStores []float64  `json:"pickup_stores"`
            AvailableModes []string  `json:"available_modes"`
         
    }
    
    // PickupStoreDetail used by Cart
    type PickupStoreDetail struct {

        
            Address string  `json:"address"`
            Pincode float64  `json:"pincode"`
            ID float64  `json:"id"`
            Landmark string  `json:"landmark"`
            Area string  `json:"area"`
            AddressType string  `json:"address_type"`
            Phone string  `json:"phone"`
            City string  `json:"city"`
            AreaCodeSlug string  `json:"area_code_slug"`
            State string  `json:"state"`
            Name string  `json:"name"`
            Country string  `json:"country"`
            UID float64  `json:"uid"`
            Email string  `json:"email"`
            StoreCode string  `json:"store_code"`
            AreaCode string  `json:"area_code"`
         
    }
    
    // StoreDetailsResponse used by Cart
    type StoreDetailsResponse struct {

        
            Items []PickupStoreDetail  `json:"items"`
         
    }
    
    // UpdateCartPaymentRequest used by Cart
    type UpdateCartPaymentRequest struct {

        
            MerchantCode string  `json:"merchant_code"`
            ID string  `json:"id"`
            AddressID string  `json:"address_id"`
            PaymentIdentifier string  `json:"payment_identifier"`
            PaymentMode string  `json:"payment_mode"`
            AggregatorName string  `json:"aggregator_name"`
         
    }
    
    // CouponValidity used by Cart
    type CouponValidity struct {

        
            Valid bool  `json:"valid"`
            Discount float64  `json:"discount"`
            Title string  `json:"title"`
            DisplayMessageEn string  `json:"display_message_en"`
            Code string  `json:"code"`
         
    }
    
    // PaymentCouponValidateSchema used by Cart
    type PaymentCouponValidateSchema struct {

        
            Message string  `json:"message"`
            CouponValidity CouponValidity  `json:"coupon_validity"`
            Success bool  `json:"success"`
         
    }
    
    // PaymentMetaSchema used by Cart
    type PaymentMetaSchema struct {

        
            MerchantCode string  `json:"merchant_code"`
            PaymentGateway string  `json:"payment_gateway"`
            PaymentIdentifier string  `json:"payment_identifier"`
            Type string  `json:"type"`
         
    }
    
    // PaymentMethodSchema used by Cart
    type PaymentMethodSchema struct {

        
            Amount float64  `json:"amount"`
            Payment string  `json:"payment"`
            PaymentMeta PaymentMetaSchema  `json:"payment_meta"`
            Name string  `json:"name"`
            Mode string  `json:"mode"`
         
    }
    
    // PlatformCartCheckoutDetailV2Request used by Cart
    type PlatformCartCheckoutDetailV2Request struct {

        
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            ID string  `json:"id"`
            CallbackURL string  `json:"callback_url"`
            BillingAddressID string  `json:"billing_address_id"`
            Pos bool  `json:"pos"`
            PaymentIdentifier string  `json:"payment_identifier"`
            Aggregator string  `json:"aggregator"`
            CheckoutMode string  `json:"checkout_mode"`
            PaymentParams map[string]interface{}  `json:"payment_params"`
            CustomMeta map[string]interface{}  `json:"custom_meta"`
            EmployeeCode string  `json:"employee_code"`
            PaymentMode string  `json:"payment_mode"`
            UserID string  `json:"user_id"`
            PickAtStoreUID float64  `json:"pick_at_store_uid"`
            Files []FilesSchema  `json:"files"`
            DeviceID string  `json:"device_id"`
            OrderingStore float64  `json:"ordering_store"`
            PaymentAutoConfirm bool  `json:"payment_auto_confirm"`
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
            Staff StaffCheckoutSchema  `json:"staff"`
            MerchantCode string  `json:"merchant_code"`
            Meta map[string]interface{}  `json:"meta"`
            AddressID string  `json:"address_id"`
            OrderType string  `json:"order_type"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
            PaymentMethods []PaymentMethodSchema  `json:"payment_methods"`
         
    }
    
    // UpdateCartPaymentRequestV2 used by Cart
    type UpdateCartPaymentRequestV2 struct {

        
            MerchantCode string  `json:"merchant_code"`
            ID string  `json:"id"`
            AddressID string  `json:"address_id"`
            PaymentIdentifier string  `json:"payment_identifier"`
            PaymentMode string  `json:"payment_mode"`
            AggregatorName string  `json:"aggregator_name"`
            PaymentMethods []PaymentMethodSchema  `json:"payment_methods"`
         
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
    
