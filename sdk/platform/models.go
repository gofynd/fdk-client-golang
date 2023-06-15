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
            PlanData map[string]interface{}  `json:"plan_data"`
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

        
            Aggregators []map[string]interface{}  `json:"aggregators"`
            DisplayFields []string  `json:"display_fields"`
            Success bool  `json:"success"`
            AppID string  `json:"app_id"`
            ExcludedFields []string  `json:"excluded_fields"`
            Created bool  `json:"created"`
         
    }
    
    // ErrorCodeDescription used by Payment
    type ErrorCodeDescription struct {

        
            Code string  `json:"code"`
            Success bool  `json:"success"`
            Description string  `json:"description"`
         
    }
    
    // PaymentGatewayConfig used by Payment
    type PaymentGatewayConfig struct {

        
            Secret string  `json:"secret"`
            IsActive bool  `json:"is_active"`
            MerchantSalt string  `json:"merchant_salt"`
            ConfigType string  `json:"config_type"`
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

        
            Code string  `json:"code"`
            Description string  `json:"description"`
         
    }
    
    // HttpErrorCodeAndResponse used by Payment
    type HttpErrorCodeAndResponse struct {

        
            Success bool  `json:"success"`
            Error ErrorCodeAndDescription  `json:"error"`
         
    }
    
    // IntentAppErrorList used by Payment
    type IntentAppErrorList struct {

        
            Code string  `json:"code"`
            PackageName string  `json:"package_name"`
         
    }
    
    // PaymentModeLogo used by Payment
    type PaymentModeLogo struct {

        
            Small string  `json:"small"`
            Large string  `json:"large"`
         
    }
    
    // IntentApp used by Payment
    type IntentApp struct {

        
            Code string  `json:"code"`
            DisplayName string  `json:"display_name"`
            PackageName string  `json:"package_name"`
            Logos PaymentModeLogo  `json:"logos"`
         
    }
    
    // PaymentModeList used by Payment
    type PaymentModeList struct {

        
            MerchantCode string  `json:"merchant_code"`
            CardBrand string  `json:"card_brand"`
            RetryCount float64  `json:"retry_count"`
            RemainingLimit float64  `json:"remaining_limit"`
            CardID string  `json:"card_id"`
            CardIssuer string  `json:"card_issuer"`
            DisplayName string  `json:"display_name"`
            CardBrandImage string  `json:"card_brand_image"`
            Expired bool  `json:"expired"`
            FyndVpa string  `json:"fynd_vpa"`
            IntentAppErrorDictList []IntentAppErrorList  `json:"intent_app_error_dict_list"`
            Timeout float64  `json:"timeout"`
            LogoURL PaymentModeLogo  `json:"logo_url"`
            Name string  `json:"name"`
            IntentApp []IntentApp  `json:"intent_app"`
            CardName string  `json:"card_name"`
            Code string  `json:"code"`
            CardFingerprint string  `json:"card_fingerprint"`
            ExpMonth float64  `json:"exp_month"`
            IntentFlow bool  `json:"intent_flow"`
            CardType string  `json:"card_type"`
            CardNumber string  `json:"card_number"`
            CompliantWithTokenisationGuidelines bool  `json:"compliant_with_tokenisation_guidelines"`
            ExpYear float64  `json:"exp_year"`
            AggregatorName string  `json:"aggregator_name"`
            CardReference string  `json:"card_reference"`
            Nickname string  `json:"nickname"`
            CodLimitPerOrder float64  `json:"cod_limit_per_order"`
            CardToken string  `json:"card_token"`
            IntentAppErrorList []string  `json:"intent_app_error_list"`
            CardIsin string  `json:"card_isin"`
            DisplayPriority float64  `json:"display_priority"`
            CodLimit float64  `json:"cod_limit"`
         
    }
    
    // RootPaymentMode used by Payment
    type RootPaymentMode struct {

        
            AggregatorName string  `json:"aggregator_name"`
            List []PaymentModeList  `json:"list"`
            AnonymousEnable bool  `json:"anonymous_enable"`
            IsPayByCardPl bool  `json:"is_pay_by_card_pl"`
            AddCardEnabled bool  `json:"add_card_enabled"`
            SaveCard bool  `json:"save_card"`
            DisplayName string  `json:"display_name"`
            Name string  `json:"name"`
            DisplayPriority float64  `json:"display_priority"`
         
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

        
            AccountHolder string  `json:"account_holder"`
            BankName string  `json:"bank_name"`
            BranchName string  `json:"branch_name"`
            State string  `json:"state"`
            Country string  `json:"country"`
            AccountNo string  `json:"account_no"`
            City string  `json:"city"`
            AccountType string  `json:"account_type"`
            IfscCode string  `json:"ifsc_code"`
         
    }
    
    // PayoutCustomer used by Payment
    type PayoutCustomer struct {

        
            ID float64  `json:"id"`
            Email string  `json:"email"`
            UniqueExternalID string  `json:"unique_external_id"`
            Name string  `json:"name"`
            Mobile string  `json:"mobile"`
         
    }
    
    // Payout used by Payment
    type Payout struct {

        
            UniqueTransferNo string  `json:"unique_transfer_no"`
            TransferType string  `json:"transfer_type"`
            IsActive bool  `json:"is_active"`
            PayoutsAggregators []PayoutAggregator  `json:"payouts_aggregators"`
            MoreAttributes PayoutMoreAttributes  `json:"more_attributes"`
            IsDefault bool  `json:"is_default"`
            Customers PayoutCustomer  `json:"customers"`
         
    }
    
    // PayoutsResponse used by Payment
    type PayoutsResponse struct {

        
            Items []Payout  `json:"items"`
            Success bool  `json:"success"`
         
    }
    
    // PayoutBankDetails used by Payment
    type PayoutBankDetails struct {

        
            AccountHolder string  `json:"account_holder"`
            BranchName string  `json:"branch_name"`
            BankName string  `json:"bank_name"`
            State string  `json:"state"`
            Country string  `json:"country"`
            AccountNo string  `json:"account_no"`
            Pincode float64  `json:"pincode"`
            City string  `json:"city"`
            AccountType string  `json:"account_type"`
            IfscCode string  `json:"ifsc_code"`
         
    }
    
    // PayoutRequest used by Payment
    type PayoutRequest struct {

        
            Aggregator string  `json:"aggregator"`
            TransferType string  `json:"transfer_type"`
            IsActive bool  `json:"is_active"`
            BankDetails PayoutBankDetails  `json:"bank_details"`
            Users map[string]interface{}  `json:"users"`
            UniqueExternalID string  `json:"unique_external_id"`
         
    }
    
    // PayoutResponse used by Payment
    type PayoutResponse struct {

        
            Aggregator string  `json:"aggregator"`
            UniqueTransferNo string  `json:"unique_transfer_no"`
            TransferType string  `json:"transfer_type"`
            Payouts map[string]interface{}  `json:"payouts"`
            PaymentStatus string  `json:"payment_status"`
            IsActive bool  `json:"is_active"`
            BankDetails map[string]interface{}  `json:"bank_details"`
            Success bool  `json:"success"`
            Users map[string]interface{}  `json:"users"`
            Created bool  `json:"created"`
         
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
            Success bool  `json:"success"`
            IsVerifiedFlag bool  `json:"is_verified_flag"`
            Message string  `json:"message"`
         
    }
    
    // NotFoundResourceError used by Payment
    type NotFoundResourceError struct {

        
            Code string  `json:"code"`
            Success bool  `json:"success"`
            Description string  `json:"description"`
         
    }
    
    // BankDetailsForOTP used by Payment
    type BankDetailsForOTP struct {

        
            AccountHolder string  `json:"account_holder"`
            BranchName string  `json:"branch_name"`
            BankName string  `json:"bank_name"`
            AccountNo string  `json:"account_no"`
            IfscCode string  `json:"ifsc_code"`
         
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
            CreatedOn string  `json:"created_on"`
            BranchName string  `json:"branch_name"`
            AccountNo string  `json:"account_no"`
            Comment string  `json:"comment"`
            DisplayName string  `json:"display_name"`
            Mobile string  `json:"mobile"`
            Address string  `json:"address"`
            Subtitle string  `json:"subtitle"`
            DelightsUserName string  `json:"delights_user_name"`
            TransferMode string  `json:"transfer_mode"`
            Email string  `json:"email"`
            ModifiedOn string  `json:"modified_on"`
            IfscCode string  `json:"ifsc_code"`
            AccountHolder string  `json:"account_holder"`
            Title string  `json:"title"`
            BeneficiaryID string  `json:"beneficiary_id"`
            BankName string  `json:"bank_name"`
            ID float64  `json:"id"`
         
    }
    
    // OrderBeneficiaryResponse used by Payment
    type OrderBeneficiaryResponse struct {

        
            ShowBeneficiaryDetails bool  `json:"show_beneficiary_details"`
            Beneficiaries []OrderBeneficiaryDetails  `json:"beneficiaries"`
         
    }
    
    // MultiTenderPaymentMeta used by Payment
    type MultiTenderPaymentMeta struct {

        
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            OrderID string  `json:"order_id"`
            PaymentGateway string  `json:"payment_gateway"`
            PaymentID string  `json:"payment_id"`
            CurrentStatus string  `json:"current_status"`
         
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
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // CODdata used by Payment
    type CODdata struct {

        
            UserID string  `json:"user_id"`
            Usages float64  `json:"usages"`
            IsActive bool  `json:"is_active"`
            RemainingLimit float64  `json:"remaining_limit"`
            Limit float64  `json:"limit"`
         
    }
    
    // GetUserCODLimitResponse used by Payment
    type GetUserCODLimitResponse struct {

        
            UserCodData CODdata  `json:"user_cod_data"`
            Success bool  `json:"success"`
         
    }
    
    // SetCODForUserRequest used by Payment
    type SetCODForUserRequest struct {

        
            IsActive bool  `json:"is_active"`
            MerchantUserID string  `json:"merchant_user_id"`
            Mobileno string  `json:"mobileno"`
         
    }
    
    // SetCODOptionResponse used by Payment
    type SetCODOptionResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
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

        
            Pincode string  `json:"pincode"`
            AddressType string  `json:"address_type"`
            City string  `json:"city"`
            Address1 string  `json:"address1"`
            Address2 string  `json:"address2"`
            State string  `json:"state"`
            Country string  `json:"country"`
            AddressCategory string  `json:"address_category"`
            UpdatedAt string  `json:"updated_at"`
            Landmark string  `json:"landmark"`
            Area string  `json:"area"`
            Latitude float64  `json:"latitude"`
            Longitude float64  `json:"longitude"`
            CreatedAt string  `json:"created_at"`
            Version string  `json:"version"`
            Phone string  `json:"phone"`
            ContactPerson string  `json:"contact_person"`
            Email string  `json:"email"`
         
    }
    
    // Prices used by Order
    type Prices struct {

        
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            PriceEffective float64  `json:"price_effective"`
            RefundCredit float64  `json:"refund_credit"`
            CouponValue float64  `json:"coupon_value"`
            RefundAmount float64  `json:"refund_amount"`
            Cashback float64  `json:"cashback"`
            AmountPaidRoundoff float64  `json:"amount_paid_roundoff"`
            CodCharges float64  `json:"cod_charges"`
            TransferPrice float64  `json:"transfer_price"`
            PriceMarked float64  `json:"price_marked"`
            ValueOfGood float64  `json:"value_of_good"`
            FyndCredits float64  `json:"fynd_credits"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            PmPriceSplit float64  `json:"pm_price_split"`
            GiftPrice float64  `json:"gift_price"`
            AmountPaid float64  `json:"amount_paid"`
            PromotionEffectiveDiscount float64  `json:"promotion_effective_discount"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            DeliveryCharge float64  `json:"delivery_charge"`
            Discount float64  `json:"discount"`
            CashbackApplied float64  `json:"cashback_applied"`
         
    }
    
    // ShipmentStatus used by Order
    type ShipmentStatus struct {

        
            ShipmentID string  `json:"shipment_id"`
            ShipmentStatusID float64  `json:"shipment_status_id"`
            CurrentShipmentStatus string  `json:"current_shipment_status"`
            BagList []string  `json:"bag_list"`
            Meta map[string]interface{}  `json:"meta"`
            StatusCreatedAt string  `json:"status_created_at"`
            Title string  `json:"title"`
            CreatedAt string  `json:"created_at"`
            Status string  `json:"status"`
         
    }
    
    // LockData used by Order
    type LockData struct {

        
            LockMessage string  `json:"lock_message"`
            Locked bool  `json:"locked"`
            Mto bool  `json:"mto"`
         
    }
    
    // ShipmentTimeStamp used by Order
    type ShipmentTimeStamp struct {

        
            Max float64  `json:"max"`
            Min float64  `json:"min"`
         
    }
    
    // Formatted used by Order
    type Formatted struct {

        
            Max string  `json:"max"`
            Min string  `json:"min"`
         
    }
    
    // ShipmentTags used by Order
    type ShipmentTags struct {

        
            EntityType string  `json:"entity_type"`
            Slug string  `json:"slug"`
            DisplayText string  `json:"display_text"`
         
    }
    
    // ShipmentItemMeta used by Order
    type ShipmentItemMeta struct {

        
            EwaybillInfo map[string]interface{}  `json:"ewaybill_info"`
            LockData LockData  `json:"lock_data"`
            External map[string]interface{}  `json:"external"`
            StoreInvoiceUpdatedDate string  `json:"store_invoice_updated_date"`
            PackagingName string  `json:"packaging_name"`
            SameStoreAvailable bool  `json:"same_store_available"`
            DpOptions map[string]interface{}  `json:"dp_options"`
            Timestamp ShipmentTimeStamp  `json:"timestamp"`
            Formatted Formatted  `json:"formatted"`
            ShipmentVolumetricWeight float64  `json:"shipment_volumetric_weight"`
            DpSortKey string  `json:"dp_sort_key"`
            FulfilmentPriorityText string  `json:"fulfilment_priority_text"`
            DebugInfo map[string]interface{}  `json:"debug_info"`
            ExistingDpList []string  `json:"existing_dp_list"`
            ShipmentWeight float64  `json:"shipment_weight"`
            Weight float64  `json:"weight"`
            Sla float64  `json:"sla"`
            ActivityComment string  `json:"activity_comment"`
            PdfMedia []map[string]interface{}  `json:"pdf_media"`
            ShipmentChargeableWeight float64  `json:"shipment_chargeable_weight"`
            ParentDpID string  `json:"parent_dp_id"`
            AutoTriggerDpAssignmentAcf bool  `json:"auto_trigger_dp_assignment_acf"`
            BagWeight map[string]interface{}  `json:"bag_weight"`
            IsInternational bool  `json:"is_international"`
            ShippingZone string  `json:"shipping_zone"`
            OrderType string  `json:"order_type"`
            AssignDpFromSb bool  `json:"assign_dp_from_sb"`
            Tags []map[string]interface{}  `json:"tags"`
            ShipmentTags []ShipmentTags  `json:"shipment_tags"`
         
    }
    
    // ShipmentItemFulFillingStore used by Order
    type ShipmentItemFulFillingStore struct {

        
            Pincode string  `json:"pincode"`
            LocationType string  `json:"location_type"`
            Code string  `json:"code"`
            StoreEmail string  `json:"store_email"`
            City string  `json:"city"`
            Meta map[string]interface{}  `json:"meta"`
            Name string  `json:"name"`
            State string  `json:"state"`
            BrandStoreTags string  `json:"brand_store_tags"`
            Address string  `json:"address"`
            Phone string  `json:"phone"`
            ID float64  `json:"id"`
         
    }
    
    // UserDataInfo used by Order
    type UserDataInfo struct {

        
            AvisUserID string  `json:"avis_user_id"`
            Mobile string  `json:"mobile"`
            Gender string  `json:"gender"`
            UID float64  `json:"uid"`
            IsAnonymousUser bool  `json:"is_anonymous_user"`
            LastName string  `json:"last_name"`
            FirstName string  `json:"first_name"`
            Name string  `json:"name"`
            Email string  `json:"email"`
         
    }
    
    // ShipmentListingChannel used by Order
    type ShipmentListingChannel struct {

        
            IsAffiliate bool  `json:"is_affiliate"`
            Logo string  `json:"logo"`
            Name string  `json:"name"`
            ChannelShipmentID string  `json:"channel_shipment_id"`
         
    }
    
    // BagStateMapper used by Order
    type BagStateMapper struct {

        
            StateType string  `json:"state_type"`
            AppFacing bool  `json:"app_facing"`
            JourneyType string  `json:"journey_type"`
            NotifyCustomer bool  `json:"notify_customer"`
            Name string  `json:"name"`
            DisplayName string  `json:"display_name"`
            AppDisplayName string  `json:"app_display_name"`
            AppStateName string  `json:"app_state_name"`
            ID float64  `json:"id"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // BagStatusHistory used by Order
    type BagStatusHistory struct {

        
            KafkaSync bool  `json:"kafka_sync"`
            ShipmentID string  `json:"shipment_id"`
            StoreID float64  `json:"store_id"`
            BagID float64  `json:"bag_id"`
            DeliveryPartnerID float64  `json:"delivery_partner_id"`
            StateType string  `json:"state_type"`
            StateID float64  `json:"state_id"`
            Forward bool  `json:"forward"`
            DisplayName string  `json:"display_name"`
            UpdatedAt string  `json:"updated_at"`
            CreatedAt string  `json:"created_at"`
            DeliveryAwbNumber string  `json:"delivery_awb_number"`
            BagStateMapper BagStateMapper  `json:"bag_state_mapper"`
            Status string  `json:"status"`
            AppDisplayName string  `json:"app_display_name"`
            BshID float64  `json:"bsh_id"`
            Reasons []map[string]interface{}  `json:"reasons"`
         
    }
    
    // ShipmentListingBrand used by Order
    type ShipmentListingBrand struct {

        
            Logo string  `json:"logo"`
            LogoBase64 string  `json:"logo_base64"`
            Name string  `json:"name"`
            CreatedOn string  `json:"created_on"`
         
    }
    
    // PlatformArticleAttributes used by Order
    type PlatformArticleAttributes struct {

        
            Currency string  `json:"currency"`
         
    }
    
    // PlatformItem used by Order
    type PlatformItem struct {

        
            Code string  `json:"code"`
            L3CategoryName string  `json:"l3_category_name"`
            Size string  `json:"size"`
            Attributes PlatformArticleAttributes  `json:"attributes"`
            Brand string  `json:"brand"`
            Color string  `json:"color"`
            BranchURL string  `json:"branch_url"`
            Images []string  `json:"images"`
            Image []string  `json:"image"`
            CanCancel bool  `json:"can_cancel"`
            CanReturn bool  `json:"can_return"`
            DepartmentID float64  `json:"department_id"`
            Meta map[string]interface{}  `json:"meta"`
            BrandID float64  `json:"brand_id"`
            ID float64  `json:"id"`
            Name string  `json:"name"`
            L3Category float64  `json:"l3_category"`
            LastUpdatedAt string  `json:"last_updated_at"`
            L1Category []string  `json:"l1_category"`
            L2Category []string  `json:"l2_category"`
            SlugKey string  `json:"slug_key"`
         
    }
    
    // Dates used by Order
    type Dates struct {

        
            OrderCreated string  `json:"order_created"`
            DeliveryDate string  `json:"delivery_date"`
         
    }
    
    // Identifier used by Order
    type Identifier struct {

        
            Isbn string  `json:"isbn"`
            SkuCode string  `json:"sku_code"`
            Ean string  `json:"ean"`
            Alu string  `json:"alu"`
            Upc string  `json:"upc"`
         
    }
    
    // FinancialBreakup used by Order
    type FinancialBreakup struct {

        
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            PriceEffective float64  `json:"price_effective"`
            RefundCredit float64  `json:"refund_credit"`
            CouponValue float64  `json:"coupon_value"`
            Cashback float64  `json:"cashback"`
            AmountPaidRoundoff float64  `json:"amount_paid_roundoff"`
            Size string  `json:"size"`
            Identifiers Identifier  `json:"identifiers"`
            GstTag string  `json:"gst_tag"`
            CodCharges float64  `json:"cod_charges"`
            GstFee float64  `json:"gst_fee"`
            TotalUnits float64  `json:"total_units"`
            TransferPrice float64  `json:"transfer_price"`
            PriceMarked float64  `json:"price_marked"`
            ValueOfGood float64  `json:"value_of_good"`
            FyndCredits float64  `json:"fynd_credits"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            AmountPaid float64  `json:"amount_paid"`
            PromotionEffectiveDiscount float64  `json:"promotion_effective_discount"`
            HsnCode string  `json:"hsn_code"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            ItemName string  `json:"item_name"`
            DeliveryCharge float64  `json:"delivery_charge"`
            Discount float64  `json:"discount"`
            CashbackApplied float64  `json:"cashback_applied"`
            AddedToFyndCash bool  `json:"added_to_fynd_cash"`
         
    }
    
    // Dimensions used by Order
    type Dimensions struct {

        
            IsDefault bool  `json:"is_default"`
            Unit string  `json:"unit"`
            Width float64  `json:"width"`
            Height float64  `json:"height"`
            Length float64  `json:"length"`
         
    }
    
    // ReturnConfig used by Order
    type ReturnConfig struct {

        
            Returnable bool  `json:"returnable"`
            Unit string  `json:"unit"`
            Time float64  `json:"time"`
         
    }
    
    // Weight used by Order
    type Weight struct {

        
            IsDefault bool  `json:"is_default"`
            Unit string  `json:"unit"`
            Shipping float64  `json:"shipping"`
         
    }
    
    // Article used by Order
    type Article struct {

        
            ASet map[string]interface{}  `json:"a_set"`
            IsSet bool  `json:"is_set"`
            RawMeta string  `json:"raw_meta"`
            Code string  `json:"code"`
            Currency map[string]interface{}  `json:"currency"`
            EspModified bool  `json:"esp_modified"`
            Identifiers map[string]interface{}  `json:"identifiers"`
            ID string  `json:"_id"`
            UID string  `json:"uid"`
            Dimensions Dimensions  `json:"dimensions"`
            ReturnConfig ReturnConfig  `json:"return_config"`
            SellerIdentifier string  `json:"seller_identifier"`
            ChildDetails map[string]interface{}  `json:"child_details"`
            Weight Weight  `json:"weight"`
            Size string  `json:"size"`
         
    }
    
    // ReplacementDetails used by Order
    type ReplacementDetails struct {

        
            ReplacementType string  `json:"replacement_type"`
            OriginalAffiliateOrderID string  `json:"original_affiliate_order_id"`
         
    }
    
    // AffiliateMeta used by Order
    type AffiliateMeta struct {

        
            IsPriority bool  `json:"is_priority"`
            OrderItemID string  `json:"order_item_id"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            DueDate string  `json:"due_date"`
            Quantity float64  `json:"quantity"`
            BoxType string  `json:"box_type"`
            MarketplaceInvoiceID string  `json:"marketplace_invoice_id"`
            EmployeeDiscount float64  `json:"employee_discount"`
            ReplacementDetails ReplacementDetails  `json:"replacement_details"`
            ChannelOrderID string  `json:"channel_order_id"`
            ChannelShipmentID string  `json:"channel_shipment_id"`
            SizeLevelTotalQty float64  `json:"size_level_total_qty"`
            CouponCode string  `json:"coupon_code"`
         
    }
    
    // AffiliateBagDetails used by Order
    type AffiliateBagDetails struct {

        
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            AffiliateMeta AffiliateMeta  `json:"affiliate_meta"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
            EmployeeDiscount float64  `json:"employee_discount"`
         
    }
    
    // GSTDetailsData used by Order
    type GSTDetailsData struct {

        
            SgstTaxPercentage float64  `json:"sgst_tax_percentage"`
            ValueOfGood float64  `json:"value_of_good"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            HsnCode string  `json:"hsn_code"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            CgstTaxPercentage float64  `json:"cgst_tax_percentage"`
            GstTag string  `json:"gst_tag"`
            IgstTaxPercentage float64  `json:"igst_tax_percentage"`
            SgstGstFee string  `json:"sgst_gst_fee"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            IsDefaultHsnCode bool  `json:"is_default_hsn_code"`
            GstFee float64  `json:"gst_fee"`
            GstinCode string  `json:"gstin_code"`
            CgstGstFee string  `json:"cgst_gst_fee"`
            HsnCodeID string  `json:"hsn_code_id"`
            IgstGstFee string  `json:"igst_gst_fee"`
         
    }
    
    // BagReturnableCancelableStatus used by Order
    type BagReturnableCancelableStatus struct {

        
            IsCustomerReturnAllowed bool  `json:"is_customer_return_allowed"`
            EnableTracking bool  `json:"enable_tracking"`
            IsReturnable bool  `json:"is_returnable"`
            CanBeCancelled bool  `json:"can_be_cancelled"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // BagUnit used by Order
    type BagUnit struct {

        
            BagID float64  `json:"bag_id"`
            CurrentOperationalStatus BagStatusHistory  `json:"current_operational_status"`
            Prices Prices  `json:"prices"`
            ProductQuantity float64  `json:"product_quantity"`
            Size string  `json:"size"`
            BagStatus []BagStatusHistory  `json:"bag_status"`
            EntityType string  `json:"entity_type"`
            Brand ShipmentListingBrand  `json:"brand"`
            DisplayName string  `json:"display_name"`
            LineNumber float64  `json:"line_number"`
            CanCancel bool  `json:"can_cancel"`
            Item PlatformItem  `json:"item"`
            Dates Dates  `json:"dates"`
            BagType string  `json:"bag_type"`
            CanReturn bool  `json:"can_return"`
            Reasons []map[string]interface{}  `json:"reasons"`
            BagExpiryDate string  `json:"bag_expiry_date"`
            FinancialBreakup []FinancialBreakup  `json:"financial_breakup"`
            Meta map[string]interface{}  `json:"meta"`
            Article Article  `json:"article"`
            AffiliateBagDetails AffiliateBagDetails  `json:"affiliate_bag_details"`
            Gst GSTDetailsData  `json:"gst"`
            CurrentStatus BagStatusHistory  `json:"current_status"`
            Status BagReturnableCancelableStatus  `json:"status"`
         
    }
    
    // ShipmentItem used by Order
    type ShipmentItem struct {

        
            DeliveryAddress PlatformDeliveryAddress  `json:"delivery_address"`
            Prices Prices  `json:"prices"`
            LockStatus bool  `json:"lock_status"`
            InvoiceID string  `json:"invoice_id"`
            ShipmentStatus ShipmentStatus  `json:"shipment_status"`
            ShipmentCreatedAt string  `json:"shipment_created_at"`
            DisplayName string  `json:"display_name"`
            PaymentMethods map[string]interface{}  `json:"payment_methods"`
            TotalBags float64  `json:"total_bags"`
            EstimatedSlaTime string  `json:"estimated_sla_time"`
            OrderDate string  `json:"order_date"`
            Meta ShipmentItemMeta  `json:"meta"`
            FulfillingStore ShipmentItemFulFillingStore  `json:"fulfilling_store"`
            OrderID string  `json:"order_id"`
            StatusCreatedAt string  `json:"status_created_at"`
            User UserDataInfo  `json:"user"`
            CustomerNote string  `json:"customer_note"`
            PreviousShipmentID string  `json:"previous_shipment_id"`
            Channel ShipmentListingChannel  `json:"channel"`
            ShipmentID string  `json:"shipment_id"`
            CanProcess bool  `json:"can_process"`
            OrderingChannnel string  `json:"ordering_channnel"`
            Bags []BagUnit  `json:"bags"`
            PaymentMode string  `json:"payment_mode"`
         
    }
    
    // ShipmentInternalPlatformViewResponse used by Order
    type ShipmentInternalPlatformViewResponse struct {

        
            Lane string  `json:"lane"`
            Page Page  `json:"page"`
            Message string  `json:"message"`
            TotalCount float64  `json:"total_count"`
            Items []ShipmentItem  `json:"items"`
            Success bool  `json:"success"`
         
    }
    
    // Error used by Order
    type Error struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // TrackingList used by Order
    type TrackingList struct {

        
            IsPassed bool  `json:"is_passed"`
            Text string  `json:"text"`
            Time string  `json:"time"`
            Status string  `json:"status"`
            IsCurrent bool  `json:"is_current"`
         
    }
    
    // UserDetailsData used by Order
    type UserDetailsData struct {

        
            Pincode string  `json:"pincode"`
            AddressType string  `json:"address_type"`
            City string  `json:"city"`
            Address1 string  `json:"address1"`
            State string  `json:"state"`
            Country string  `json:"country"`
            Landmark string  `json:"landmark"`
            Area string  `json:"area"`
            Address string  `json:"address"`
            Phone string  `json:"phone"`
            Name string  `json:"name"`
            Email string  `json:"email"`
         
    }
    
    // ShipmentDetails used by Order
    type ShipmentDetails struct {

        
            LockMessage string  `json:"lock_message"`
            ActionToStatus map[string]interface{}  `json:"action_to_status"`
            LockStatus bool  `json:"lock_status"`
         
    }
    
    // PDFLinks used by Order
    type PDFLinks struct {

        
            InvoiceA4 string  `json:"invoice_a4"`
            DeliveryChallanA4 string  `json:"delivery_challan_a4"`
            LabelA4 string  `json:"label_a4"`
            LabelExport string  `json:"label_export"`
            CreditNoteURL string  `json:"credit_note_url"`
            Label string  `json:"label"`
            LabelA6 string  `json:"label_a6"`
            B2b string  `json:"b2b"`
            Invoice string  `json:"invoice"`
            InvoiceExport string  `json:"invoice_export"`
            LabelPos string  `json:"label_pos"`
            InvoiceA6 string  `json:"invoice_a6"`
            InvoiceType string  `json:"invoice_type"`
            LabelType string  `json:"label_type"`
            PoInvoice string  `json:"po_invoice"`
            InvoicePos string  `json:"invoice_pos"`
         
    }
    
    // DebugInfo used by Order
    type DebugInfo struct {

        
            StormbreakerUUID string  `json:"stormbreaker_uuid"`
         
    }
    
    // EinvoiceInfo used by Order
    type EinvoiceInfo struct {

        
            CreditNote map[string]interface{}  `json:"credit_note"`
            Invoice map[string]interface{}  `json:"invoice"`
         
    }
    
    // BuyerDetails used by Order
    type BuyerDetails struct {

        
            Pincode float64  `json:"pincode"`
            AjioSiteID string  `json:"ajio_site_id"`
            City string  `json:"city"`
            State string  `json:"state"`
            Gstin string  `json:"gstin"`
            Address string  `json:"address"`
            Name string  `json:"name"`
         
    }
    
    // ShipmentMeta used by Order
    type ShipmentMeta struct {

        
            ReturnStoreNode float64  `json:"return_store_node"`
            EwaybillInfo map[string]interface{}  `json:"ewaybill_info"`
            PoNumber string  `json:"po_number"`
            Dimension Dimensions  `json:"dimension"`
            B2cBuyerDetails map[string]interface{}  `json:"b2c_buyer_details"`
            BoxType string  `json:"box_type"`
            LockData LockData  `json:"lock_data"`
            External map[string]interface{}  `json:"external"`
            StoreInvoiceUpdatedDate string  `json:"store_invoice_updated_date"`
            PackagingName string  `json:"packaging_name"`
            SameStoreAvailable bool  `json:"same_store_available"`
            DpOptions map[string]interface{}  `json:"dp_options"`
            Timestamp ShipmentTimeStamp  `json:"timestamp"`
            ForwardAffiliateShipmentID string  `json:"forward_affiliate_shipment_id"`
            DpName string  `json:"dp_name"`
            ReturnDetails map[string]interface{}  `json:"return_details"`
            Formatted Formatted  `json:"formatted"`
            AwbNumber string  `json:"awb_number"`
            DueDate string  `json:"due_date"`
            ShipmentVolumetricWeight float64  `json:"shipment_volumetric_weight"`
            DpSortKey string  `json:"dp_sort_key"`
            FulfilmentPriorityText string  `json:"fulfilment_priority_text"`
            DebugInfo DebugInfo  `json:"debug_info"`
            ShipmentWeight float64  `json:"shipment_weight"`
            Weight float64  `json:"weight"`
            ShipmentTags []ShipmentTags  `json:"shipment_tags"`
            MarketplaceStoreID string  `json:"marketplace_store_id"`
            ReturnAffiliateShipmentID string  `json:"return_affiliate_shipment_id"`
            ForwardAffiliateOrderID string  `json:"forward_affiliate_order_id"`
            DpID string  `json:"dp_id"`
            ReturnAwbNumber string  `json:"return_awb_number"`
            ParentDpID string  `json:"parent_dp_id"`
            BagWeight map[string]interface{}  `json:"bag_weight"`
            AutoTriggerDpAssignmentAcf bool  `json:"auto_trigger_dp_assignment_acf"`
            OrderType string  `json:"order_type"`
            AssignDpFromSb bool  `json:"assign_dp_from_sb"`
            ReturnAffiliateOrderID string  `json:"return_affiliate_order_id"`
            EinvoiceInfo EinvoiceInfo  `json:"einvoice_info"`
            B2bBuyerDetails BuyerDetails  `json:"b2b_buyer_details"`
         
    }
    
    // AffiliateDetails used by Order
    type AffiliateDetails struct {

        
            AdID string  `json:"ad_id"`
            AffiliateMeta AffiliateMeta  `json:"affiliate_meta"`
            AffiliateID string  `json:"affiliate_id"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
            CompanyAffiliateTag string  `json:"company_affiliate_tag"`
            PdfLinks PDFLinks  `json:"pdf_links"`
            ShipmentMeta ShipmentMeta  `json:"shipment_meta"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            AffiliateStoreID string  `json:"affiliate_store_id"`
         
    }
    
    // Meta used by Order
    type Meta struct {

        
            MerchantCode string  `json:"merchant_code"`
            PaymentID string  `json:"payment_id"`
            LogoURL map[string]string  `json:"logo_url"`
            PaymentGateway string  `json:"payment_gateway"`
            PaymentIdentifier string  `json:"payment_identifier"`
         
    }
    
    // PaymentMethod used by Order
    type PaymentMethod struct {

        
            Mode string  `json:"mode"`
            Meta Meta  `json:"meta"`
            RefundBy string  `json:"refund_by"`
            CollectBy string  `json:"collect_by"`
            Amount float64  `json:"amount"`
            Name string  `json:"name"`
         
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

        
            CompanyGst string  `json:"company_gst"`
            CompanyContact ContactDetails  `json:"company_contact"`
            CompanyCin string  `json:"company_cin"`
            CompanyID float64  `json:"company_id"`
            Address map[string]interface{}  `json:"address"`
            CompanyName string  `json:"company_name"`
         
    }
    
    // InvoiceInfo used by Order
    type InvoiceInfo struct {

        
            InvoiceURL string  `json:"invoice_url"`
            StoreInvoiceID string  `json:"store_invoice_id"`
            CreditNoteID string  `json:"credit_note_id"`
            ExternalInvoiceID string  `json:"external_invoice_id"`
            LabelURL string  `json:"label_url"`
            UpdatedDate string  `json:"updated_date"`
         
    }
    
    // FulfillingStore used by Order
    type FulfillingStore struct {

        
            Pincode string  `json:"pincode"`
            Code string  `json:"code"`
            City string  `json:"city"`
            Meta map[string]interface{}  `json:"meta"`
            FulfillmentChannel string  `json:"fulfillment_channel"`
            State string  `json:"state"`
            Country string  `json:"country"`
            StoreName string  `json:"store_name"`
            Address string  `json:"address"`
            Phone string  `json:"phone"`
            ID float64  `json:"id"`
            ContactPerson string  `json:"contact_person"`
         
    }
    
    // OrderingStoreDetails used by Order
    type OrderingStoreDetails struct {

        
            Pincode string  `json:"pincode"`
            Code string  `json:"code"`
            City string  `json:"city"`
            Meta map[string]interface{}  `json:"meta"`
            State string  `json:"state"`
            Country string  `json:"country"`
            StoreName string  `json:"store_name"`
            Address string  `json:"address"`
            Phone string  `json:"phone"`
            ID float64  `json:"id"`
            ContactPerson string  `json:"contact_person"`
         
    }
    
    // DPDetailsData used by Order
    type DPDetailsData struct {

        
            Pincode string  `json:"pincode"`
            ID float64  `json:"id"`
            GstTag string  `json:"gst_tag"`
            Country string  `json:"country"`
            EwayBillID string  `json:"eway_bill_id"`
            TrackURL string  `json:"track_url"`
            AwbNo string  `json:"awb_no"`
            Name string  `json:"name"`
         
    }
    
    // ShipmentPayments used by Order
    type ShipmentPayments struct {

        
            Mode string  `json:"mode"`
            Source string  `json:"source"`
            Logo string  `json:"logo"`
         
    }
    
    // ShipmentStatusData used by Order
    type ShipmentStatusData struct {

        
            ShipmentID string  `json:"shipment_id"`
            BagList []string  `json:"bag_list"`
            Meta map[string]interface{}  `json:"meta"`
            DisplayName string  `json:"display_name"`
            CreatedAt string  `json:"created_at"`
            Status string  `json:"status"`
            ID float64  `json:"id"`
         
    }
    
    // OrderDetailsData used by Order
    type OrderDetailsData struct {

        
            OrderingChannelLogo map[string]interface{}  `json:"ordering_channel_logo"`
            OrderDate string  `json:"order_date"`
            AffiliateID string  `json:"affiliate_id"`
            CodCharges string  `json:"cod_charges"`
            OrderingChannel string  `json:"ordering_channel"`
            FyndOrderID string  `json:"fynd_order_id"`
            Source string  `json:"source"`
            OrderValue string  `json:"order_value"`
            TaxDetails map[string]interface{}  `json:"tax_details"`
         
    }
    
    // BagGST used by Order
    type BagGST struct {

        
            SgstTaxPercentage float64  `json:"sgst_tax_percentage"`
            ValueOfGood float64  `json:"value_of_good"`
            HsnCode string  `json:"hsn_code"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            CgstTaxPercentage float64  `json:"cgst_tax_percentage"`
            GstTag string  `json:"gst_tag"`
            IgstTaxPercentage float64  `json:"igst_tax_percentage"`
            SgstGstFee string  `json:"sgst_gst_fee"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            IsDefaultHsnCode bool  `json:"is_default_hsn_code"`
            GstFee float64  `json:"gst_fee"`
            GstinCode string  `json:"gstin_code"`
            CgstGstFee string  `json:"cgst_gst_fee"`
            HsnCodeID string  `json:"hsn_code_id"`
            IgstGstFee string  `json:"igst_gst_fee"`
         
    }
    
    // OrderBrandName used by Order
    type OrderBrandName struct {

        
            ModifiedOn string  `json:"modified_on"`
            BrandName string  `json:"brand_name"`
            CreatedOn string  `json:"created_on"`
            Logo string  `json:"logo"`
            Company float64  `json:"company"`
            ID float64  `json:"id"`
         
    }
    
    // BagConfigs used by Order
    type BagConfigs struct {

        
            IsCustomerReturnAllowed bool  `json:"is_customer_return_allowed"`
            EnableTracking bool  `json:"enable_tracking"`
            IsReturnable bool  `json:"is_returnable"`
            CanBeCancelled bool  `json:"can_be_cancelled"`
            AllowForceReturn bool  `json:"allow_force_return"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // B2BPODetails used by Order
    type B2BPODetails struct {

        
            PoLineAmount float64  `json:"po_line_amount"`
            DockerNumber string  `json:"docker_number"`
            TotalGstPercentage float64  `json:"total_gst_percentage"`
            ItemBasePrice float64  `json:"item_base_price"`
            PartialCanRet bool  `json:"partial_can_ret"`
            PoTaxAmount float64  `json:"po_tax_amount"`
         
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

        
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            CustomMessage string  `json:"custom_message"`
            B2bPoDetails B2BPODetails  `json:"b2b_po_details"`
            GiftCard GiftCard  `json:"gift_card"`
            GroupID string  `json:"group_id"`
            DocketNumber string  `json:"docket_number"`
            PartialCanRet bool  `json:"partial_can_ret"`
            CustomJson map[string]interface{}  `json:"custom_json"`
         
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
    
    // DiscountRules used by Order
    type DiscountRules struct {

        
            Type string  `json:"type"`
            Value float64  `json:"value"`
         
    }
    
    // AppliedPromos used by Order
    type AppliedPromos struct {

        
            PromoID string  `json:"promo_id"`
            MrpPromotion bool  `json:"mrp_promotion"`
            BuyRules []BuyRules  `json:"buy_rules"`
            DiscountRules []DiscountRules  `json:"discount_rules"`
            PromotionType string  `json:"promotion_type"`
            ArticleQuantity float64  `json:"article_quantity"`
            PromotionName string  `json:"promotion_name"`
            Amount float64  `json:"amount"`
         
    }
    
    // OrderBagArticle used by Order
    type OrderBagArticle struct {

        
            UID string  `json:"uid"`
            Size string  `json:"size"`
            ReturnConfig ReturnConfig  `json:"return_config"`
            Identifiers map[string]interface{}  `json:"identifiers"`
         
    }
    
    // AffiliateBagsDetails used by Order
    type AffiliateBagsDetails struct {

        
            AffiliateBagID string  `json:"affiliate_bag_id"`
            CouponCode string  `json:"coupon_code"`
         
    }
    
    // CurrentStatus used by Order
    type CurrentStatus struct {

        
            KafkaSync bool  `json:"kafka_sync"`
            StoreID float64  `json:"store_id"`
            BagID float64  `json:"bag_id"`
            StateType string  `json:"state_type"`
            DeliveryPartnerID float64  `json:"delivery_partner_id"`
            ShipmentID string  `json:"shipment_id"`
            StateID float64  `json:"state_id"`
            UpdatedAt float64  `json:"updated_at"`
            CreatedAt string  `json:"created_at"`
            DeliveryAwbNumber string  `json:"delivery_awb_number"`
            BagStateMapper BagStateMapper  `json:"bag_state_mapper"`
            Status string  `json:"status"`
            ID float64  `json:"id"`
         
    }
    
    // OrderBags used by Order
    type OrderBags struct {

        
            BagID float64  `json:"bag_id"`
            Identifier string  `json:"identifier"`
            Quantity float64  `json:"quantity"`
            GstDetails BagGST  `json:"gst_details"`
            DeliveryAddress PlatformDeliveryAddress  `json:"delivery_address"`
            Prices Prices  `json:"prices"`
            GroupID string  `json:"group_id"`
            EntityType string  `json:"entity_type"`
            Brand OrderBrandName  `json:"brand"`
            DisplayName string  `json:"display_name"`
            LineNumber float64  `json:"line_number"`
            CanCancel bool  `json:"can_cancel"`
            Item PlatformItem  `json:"item"`
            CanReturn bool  `json:"can_return"`
            BagConfigs BagConfigs  `json:"bag_configs"`
            FinancialBreakup FinancialBreakup  `json:"financial_breakup"`
            Meta BagMeta  `json:"meta"`
            AppliedPromos []AppliedPromos  `json:"applied_promos"`
            Article OrderBagArticle  `json:"article"`
            AffiliateBagDetails AffiliateBagsDetails  `json:"affiliate_bag_details"`
            SellerIdentifier string  `json:"seller_identifier"`
            CurrentStatus CurrentStatus  `json:"current_status"`
            ParentPromoBags map[string]interface{}  `json:"parent_promo_bags"`
         
    }
    
    // PlatformShipment used by Order
    type PlatformShipment struct {

        
            CanUpdateDimension bool  `json:"can_update_dimension"`
            GstDetails GSTDetailsData  `json:"gst_details"`
            ShipmentImages []string  `json:"shipment_images"`
            Prices Prices  `json:"prices"`
            Coupon map[string]interface{}  `json:"coupon"`
            LockStatus bool  `json:"lock_status"`
            ShipmentQuantity float64  `json:"shipment_quantity"`
            InvoiceID string  `json:"invoice_id"`
            FulfilmentPriority float64  `json:"fulfilment_priority"`
            PriorityText string  `json:"priority_text"`
            CustomMeta []map[string]interface{}  `json:"custom_meta"`
            ShipmentStatus string  `json:"shipment_status"`
            TrackingList []TrackingList  `json:"tracking_list"`
            CustomMessage string  `json:"custom_message"`
            BillingDetails UserDetailsData  `json:"billing_details"`
            JourneyType string  `json:"journey_type"`
            PackagingType string  `json:"packaging_type"`
            ShipmentCreatedAt string  `json:"shipment_created_at"`
            ShipmentDetails ShipmentDetails  `json:"shipment_details"`
            PlatformLogo string  `json:"platform_logo"`
            AffiliateDetails AffiliateDetails  `json:"affiliate_details"`
            PaymentMethods map[string]PaymentMethod  `json:"payment_methods"`
            TotalBags float64  `json:"total_bags"`
            CompanyDetails CompanyDetails  `json:"company_details"`
            EstimatedSlaTime string  `json:"estimated_sla_time"`
            PickedDate string  `json:"picked_date"`
            DeliveryDetails UserDetailsData  `json:"delivery_details"`
            Meta ShipmentMeta  `json:"meta"`
            TotalItems float64  `json:"total_items"`
            Invoice InvoiceInfo  `json:"invoice"`
            FulfillingStore FulfillingStore  `json:"fulfilling_store"`
            User UserDataInfo  `json:"user"`
            OrderingStore OrderingStoreDetails  `json:"ordering_store"`
            OperationalStatus string  `json:"operational_status"`
            UserAgent string  `json:"user_agent"`
            ShipmentID string  `json:"shipment_id"`
            DpDetails DPDetailsData  `json:"dp_details"`
            IsDpAssignEnabled bool  `json:"is_dp_assign_enabled"`
            Vertical string  `json:"vertical"`
            DpAssignment bool  `json:"dp_assignment"`
            ForwardShipmentID string  `json:"forward_shipment_id"`
            EnableDpTracking bool  `json:"enable_dp_tracking"`
            BagStatusHistory []BagStatusHistory  `json:"bag_status_history"`
            DeliverySlot map[string]interface{}  `json:"delivery_slot"`
            Payments ShipmentPayments  `json:"payments"`
            Status ShipmentStatusData  `json:"status"`
            Order OrderDetailsData  `json:"order"`
            Bags []OrderBags  `json:"bags"`
            PaymentMode string  `json:"payment_mode"`
         
    }
    
    // ShipmentInfoResponse used by Order
    type ShipmentInfoResponse struct {

        
            Shipments []PlatformShipment  `json:"shipments"`
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // AssetByShipment used by Order
    type AssetByShipment struct {

        
            ShipmentID string  `json:"shipment_id"`
            PresignedType string  `json:"presigned_type"`
            PresignedUrls map[string]string  `json:"presigned_urls"`
            ExpiresIn string  `json:"expires_in"`
            Success bool  `json:"success"`
         
    }
    
    // ResponseGetAssetShipment used by Order
    type ResponseGetAssetShipment struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
            Data []AssetByShipment  `json:"data"`
         
    }
    
    // TransactionData used by Order
    type TransactionData struct {

        
            Currency string  `json:"currency"`
            TerminalID string  `json:"terminal_id"`
            TransactionID string  `json:"transaction_id"`
            PaymentID string  `json:"payment_id"`
            UniqueReferenceNumber string  `json:"unique_reference_number"`
            AmountPaid float64  `json:"amount_paid"`
            Entity string  `json:"entity"`
            Status string  `json:"status"`
         
    }
    
    // BillingStaffDetails used by Order
    type BillingStaffDetails struct {

        
            User string  `json:"user"`
            StaffID float64  `json:"staff_id"`
            LastName string  `json:"last_name"`
            FirstName string  `json:"first_name"`
            EmployeeCode string  `json:"employee_code"`
         
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

        
            MongoCartID float64  `json:"mongo_cart_id"`
            TransactionData TransactionData  `json:"transaction_data"`
            EmployeeID float64  `json:"employee_id"`
            CartID float64  `json:"cart_id"`
            BillingStaffDetails BillingStaffDetails  `json:"billing_staff_details"`
            Staff map[string]interface{}  `json:"staff"`
            Comment string  `json:"comment"`
            OrderingStore float64  `json:"ordering_store"`
            CustomerNote string  `json:"customer_note"`
            OrderPlatform string  `json:"order_platform"`
            CompanyLogo string  `json:"company_logo"`
            OrderTags []map[string]interface{}  `json:"order_tags"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            OrderType string  `json:"order_type"`
            OrderChildEntities []string  `json:"order_child_entities"`
            CurrencySymbol string  `json:"currency_symbol"`
            PlatformUserDetails PlatformUserDetails  `json:"platform_user_details"`
            PaymentType string  `json:"payment_type"`
            Files []map[string]interface{}  `json:"files"`
         
    }
    
    // TaxDetails used by Order
    type TaxDetails struct {

        
            PanNo string  `json:"pan_no"`
            Gstin string  `json:"gstin"`
         
    }
    
    // OrderData used by Order
    type OrderData struct {

        
            OrderDate string  `json:"order_date"`
            Meta OrderMeta  `json:"meta"`
            Prices Prices  `json:"prices"`
            FyndOrderID string  `json:"fynd_order_id"`
            PaymentMethods map[string]PaymentMethod  `json:"payment_methods"`
            TaxDetails TaxDetails  `json:"tax_details"`
         
    }
    
    // OrderDetailsResponse used by Order
    type OrderDetailsResponse struct {

        
            Shipments []PlatformShipment  `json:"shipments"`
            Success bool  `json:"success"`
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
            TotalItems float64  `json:"total_items"`
            Value string  `json:"value"`
            Options []SubLane  `json:"options"`
         
    }
    
    // LaneConfigResponse used by Order
    type LaneConfigResponse struct {

        
            SuperLanes []SuperLane  `json:"super_lanes"`
         
    }
    
    // PlatformBreakupValues used by Order
    type PlatformBreakupValues struct {

        
            Value string  `json:"value"`
            Name string  `json:"name"`
            Display string  `json:"display"`
         
    }
    
    // PlatformChannel used by Order
    type PlatformChannel struct {

        
            Logo string  `json:"logo"`
            Name string  `json:"name"`
         
    }
    
    // PlatformOrderItems used by Order
    type PlatformOrderItems struct {

        
            TotalOrderValue float64  `json:"total_order_value"`
            UserInfo UserDataInfo  `json:"user_info"`
            OrderCreatedTime string  `json:"order_created_time"`
            Meta map[string]interface{}  `json:"meta"`
            Shipments []PlatformShipment  `json:"shipments"`
            BreakupValues []PlatformBreakupValues  `json:"breakup_values"`
            OrderID string  `json:"order_id"`
            OrderValue float64  `json:"order_value"`
            Channel PlatformChannel  `json:"channel"`
            PaymentMode string  `json:"payment_mode"`
         
    }
    
    // OrderListingResponse used by Order
    type OrderListingResponse struct {

        
            Lane string  `json:"lane"`
            Page Page  `json:"page"`
            Message string  `json:"message"`
            TotalCount float64  `json:"total_count"`
            Items []PlatformOrderItems  `json:"items"`
            Success bool  `json:"success"`
         
    }
    
    // Options used by Order
    type Options struct {

        
            Text string  `json:"text"`
            Value float64  `json:"value"`
         
    }
    
    // MetricsCount used by Order
    type MetricsCount struct {

        
            Options []Options  `json:"options"`
            Text string  `json:"text"`
            Key string  `json:"key"`
            Value float64  `json:"value"`
         
    }
    
    // MetricCountResponse used by Order
    type MetricCountResponse struct {

        
            Items []MetricsCount  `json:"items"`
         
    }
    
    // PlatformTrack used by Order
    type PlatformTrack struct {

        
            UpdatedTime string  `json:"updated_time"`
            Reason string  `json:"reason"`
            AccountName string  `json:"account_name"`
            Meta map[string]interface{}  `json:"meta"`
            UpdatedAt string  `json:"updated_at"`
            LastLocationRecievedAt string  `json:"last_location_recieved_at"`
            Status string  `json:"status"`
            ShipmentType string  `json:"shipment_type"`
            Awb string  `json:"awb"`
            RawStatus string  `json:"raw_status"`
         
    }
    
    // PlatformShipmentTrack used by Order
    type PlatformShipmentTrack struct {

        
            Meta map[string]interface{}  `json:"meta"`
            Results []PlatformTrack  `json:"results"`
         
    }
    
    // FilterInfoOption used by Order
    type FilterInfoOption struct {

        
            ShowUI bool  `json:"show_ui"`
            MinSearchSize float64  `json:"min_search_size"`
            Text string  `json:"text"`
            Value string  `json:"value"`
            PlaceholderText string  `json:"placeholder_text"`
            Name string  `json:"name"`
         
    }
    
    // FiltersInfo used by Order
    type FiltersInfo struct {

        
            Text string  `json:"text"`
            Options []FilterInfoOption  `json:"options"`
            Value string  `json:"value"`
            PlaceholderText string  `json:"placeholder_text"`
            Type string  `json:"type"`
            Required bool  `json:"required"`
         
    }
    
    // AdvanceFilterInfo used by Order
    type AdvanceFilterInfo struct {

        
            Returned []FiltersInfo  `json:"returned"`
            Page map[string]interface{}  `json:"page"`
            AppliedFilters map[string]interface{}  `json:"applied_filters"`
            Processed []FiltersInfo  `json:"processed"`
            ActionCentre []FiltersInfo  `json:"action_centre"`
            Unfulfilled []FiltersInfo  `json:"unfulfilled"`
            Filters []FiltersInfo  `json:"filters"`
         
    }
    
    // FiltersResponse used by Order
    type FiltersResponse struct {

        
            AdvanceFilter AdvanceFilterInfo  `json:"advance_filter"`
            GlobalFilter []FiltersInfo  `json:"global_filter"`
         
    }
    
    // Success used by Order
    type Success struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // OmsReports used by Order
    type OmsReports struct {

        
            ReportRequestedAt string  `json:"report_requested_at"`
            ReportID string  `json:"report_id"`
            S3Key string  `json:"s3_key"`
            ReportType string  `json:"report_type"`
            DisplayName string  `json:"display_name"`
            ReportCreatedAt string  `json:"report_created_at"`
            ReportName string  `json:"report_name"`
            Status string  `json:"status"`
            RequestDetails map[string]interface{}  `json:"request_details"`
         
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

        
            Type string  `json:"type"`
            Message string  `json:"message"`
            Value string  `json:"value"`
         
    }
    
    // JioCodeUpsertResponse used by Order
    type JioCodeUpsertResponse struct {

        
            Identifier string  `json:"identifier"`
            Data []map[string]interface{}  `json:"data"`
            TraceID string  `json:"trace_id"`
            Error []NestedErrorSchemaDataSet  `json:"error"`
            Success bool  `json:"success"`
         
    }
    
    // BulkInvoicingResponse used by Order
    type BulkInvoicingResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // BulkInvoiceLabelResponse used by Order
    type BulkInvoiceLabelResponse struct {

        
            BatchID string  `json:"batch_id"`
            StoreID string  `json:"store_id"`
            Label map[string]interface{}  `json:"label"`
            DoInvoiceLabelGenerated bool  `json:"do_invoice_label_generated"`
            StoreCode string  `json:"store_code"`
            Data map[string]interface{}  `json:"data"`
            Invoice map[string]interface{}  `json:"invoice"`
            CompanyID string  `json:"company_id"`
            StoreName string  `json:"store_name"`
            InvoiceStatus string  `json:"invoice_status"`
         
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
    
    // BulkListingPage used by Order
    type BulkListingPage struct {

        
            Size float64  `json:"size"`
            HasNext bool  `json:"has_next"`
            Total float64  `json:"total"`
            HasPrevious bool  `json:"has_previous"`
            Type string  `json:"type"`
            Current float64  `json:"current"`
         
    }
    
    // bulkListingData used by Order
    type bulkListingData struct {

        
            BatchID string  `json:"batch_id"`
            Successful float64  `json:"successful"`
            UploadedOn string  `json:"uploaded_on"`
            ProcessingShipments []string  `json:"processing_shipments"`
            StoreCode string  `json:"store_code"`
            Processing float64  `json:"processing"`
            SuccessfulShipments []map[string]interface{}  `json:"successful_shipments"`
            UserID string  `json:"user_id"`
            Failed float64  `json:"failed"`
            StoreID float64  `json:"store_id"`
            ExcelURL string  `json:"excel_url"`
            CompanyID float64  `json:"company_id"`
            UserName string  `json:"user_name"`
            StoreName string  `json:"store_name"`
            ID string  `json:"id"`
            Total float64  `json:"total"`
            FailedShipments []map[string]interface{}  `json:"failed_shipments"`
            FileName string  `json:"file_name"`
            Status string  `json:"status"`
         
    }
    
    // BulkListingResponse used by Order
    type BulkListingResponse struct {

        
            Page BulkListingPage  `json:"page"`
            Data []bulkListingData  `json:"data"`
            Success bool  `json:"success"`
            Error string  `json:"error"`
         
    }
    
    // QuestionSet used by Order
    type QuestionSet struct {

        
            ID float64  `json:"id"`
            DisplayName string  `json:"display_name"`
         
    }
    
    // Reason used by Order
    type Reason struct {

        
            QuestionSet []QuestionSet  `json:"question_set"`
            ID float64  `json:"id"`
            QcType []string  `json:"qc_type"`
            DisplayName string  `json:"display_name"`
         
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

        
            BatchID string  `json:"batch_id"`
            TotalShipmentsCount float64  `json:"total_shipments_count"`
            CompanyID string  `json:"company_id"`
            FailedShipmentsCount float64  `json:"failed_shipments_count"`
            SuccessfulShipmentsCount float64  `json:"successful_shipments_count"`
            SuccessfulShipmentIds []string  `json:"successful_shipment_ids"`
            ProcessingShipmentsCount float64  `json:"processing_shipments_count"`
         
    }
    
    // BulkActionDetailsResponse used by Order
    type BulkActionDetailsResponse struct {

        
            UploadedBy string  `json:"uploaded_by"`
            Data []BulkActionDetailsDataField  `json:"data"`
            UploadedOn string  `json:"uploaded_on"`
            Message string  `json:"message"`
            UserID string  `json:"user_id"`
            Status bool  `json:"status"`
            Error []string  `json:"error"`
            Success string  `json:"success"`
            FailedRecords []string  `json:"failed_records"`
         
    }
    
    // BagGSTDetails used by Order
    type BagGSTDetails struct {

        
            SgstTaxPercentage float64  `json:"sgst_tax_percentage"`
            ValueOfGood float64  `json:"value_of_good"`
            HsnCode string  `json:"hsn_code"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            GstTag string  `json:"gst_tag"`
            IgstTaxPercentage float64  `json:"igst_tax_percentage"`
            CgstTaxPercentage float64  `json:"cgst_tax_percentage"`
            SgstGstFee string  `json:"sgst_gst_fee"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            IsDefaultHsnCode bool  `json:"is_default_hsn_code"`
            GstFee float64  `json:"gst_fee"`
            GstinCode string  `json:"gstin_code"`
            HsnCodeID string  `json:"hsn_code_id"`
            CgstGstFee string  `json:"cgst_gst_fee"`
            IgstGstFee string  `json:"igst_gst_fee"`
         
    }
    
    // Brand used by Order
    type Brand struct {

        
            ModifiedOn float64  `json:"modified_on"`
            BrandName string  `json:"brand_name"`
            CreatedOn float64  `json:"created_on"`
            StartDate string  `json:"start_date"`
            InvoicePrefix string  `json:"invoice_prefix"`
            CreditNoteExpiryDays float64  `json:"credit_note_expiry_days"`
            ScriptLastRan string  `json:"script_last_ran"`
            Logo string  `json:"logo"`
            CreditNoteAllowed bool  `json:"credit_note_allowed"`
            BrandID float64  `json:"brand_id"`
            Company string  `json:"company"`
            IsVirtualInvoice bool  `json:"is_virtual_invoice"`
            PickupLocation string  `json:"pickup_location"`
         
    }
    
    // Attributes used by Order
    type Attributes struct {

        
            PrimaryColorHex string  `json:"primary_color_hex"`
            BrandName string  `json:"brand_name"`
            Gender []string  `json:"gender"`
            PrimaryMaterial string  `json:"primary_material"`
            MarketerName string  `json:"marketer_name"`
            MarketerAddress string  `json:"marketer_address"`
            PrimaryColor string  `json:"primary_color"`
            Essential string  `json:"essential"`
            Name string  `json:"name"`
         
    }
    
    // Item used by Order
    type Item struct {

        
            Code string  `json:"code"`
            L3CategoryName string  `json:"l3_category_name"`
            Size string  `json:"size"`
            Attributes Attributes  `json:"attributes"`
            ItemID float64  `json:"item_id"`
            Brand string  `json:"brand"`
            Color string  `json:"color"`
            BranchURL string  `json:"branch_url"`
            Image []string  `json:"image"`
            CanCancel bool  `json:"can_cancel"`
            CanReturn bool  `json:"can_return"`
            L2CategoryID float64  `json:"l2_category_id"`
            DepartmentID float64  `json:"department_id"`
            L1CategoryID float64  `json:"l1_category_id"`
            Gender string  `json:"gender"`
            Meta map[string]interface{}  `json:"meta"`
            BrandID float64  `json:"brand_id"`
            WebstoreProductURL string  `json:"webstore_product_url"`
            Name string  `json:"name"`
            L3Category float64  `json:"l3_category"`
            LastUpdatedAt string  `json:"last_updated_at"`
            L1Category []string  `json:"l1_category"`
            L2Category []string  `json:"l2_category"`
            SlugKey string  `json:"slug_key"`
         
    }
    
    // Document used by Order
    type Document struct {

        
            DsType string  `json:"ds_type"`
            URL string  `json:"url"`
            Value string  `json:"value"`
            LegalName string  `json:"legal_name"`
            Verified bool  `json:"verified"`
         
    }
    
    // StoreDocuments used by Order
    type StoreDocuments struct {

        
            Gst Document  `json:"gst"`
         
    }
    
    // StoreEwaybill used by Order
    type StoreEwaybill struct {

        
            Enabled bool  `json:"enabled"`
         
    }
    
    // StoreEinvoice used by Order
    type StoreEinvoice struct {

        
            Password string  `json:"password"`
            User string  `json:"user"`
            Username string  `json:"username"`
            Enabled bool  `json:"enabled"`
         
    }
    
    // StoreGstCredentials used by Order
    type StoreGstCredentials struct {

        
            EWaybill StoreEwaybill  `json:"e_waybill"`
            EInvoice StoreEinvoice  `json:"e_invoice"`
         
    }
    
    // EInvoicePortalDetails used by Order
    type EInvoicePortalDetails struct {

        
            User string  `json:"user"`
            Password string  `json:"password"`
            Username string  `json:"username"`
         
    }
    
    // StoreMeta used by Order
    type StoreMeta struct {

        
            ProductReturnConfig map[string]interface{}  `json:"product_return_config"`
            NotificationEmails []string  `json:"notification_emails"`
            Documents StoreDocuments  `json:"documents"`
            GstCredentials StoreGstCredentials  `json:"gst_credentials"`
            DisplayName string  `json:"display_name"`
            EinvoicePortalDetails EInvoicePortalDetails  `json:"einvoice_portal_details"`
            EwaybillPortalDetails map[string]interface{}  `json:"ewaybill_portal_details"`
            Timing []map[string]interface{}  `json:"timing"`
            AdditionalContactDetails map[string]interface{}  `json:"additional_contact_details"`
            Stage string  `json:"stage"`
            GstNumber string  `json:"gst_number"`
         
    }
    
    // StoreAddress used by Order
    type StoreAddress struct {

        
            State string  `json:"state"`
            UpdatedAt string  `json:"updated_at"`
            CreatedAt string  `json:"created_at"`
            CountryCode string  `json:"country_code"`
            Phone string  `json:"phone"`
            City string  `json:"city"`
            Country string  `json:"country"`
            Landmark string  `json:"landmark"`
            Longitude float64  `json:"longitude"`
            Email string  `json:"email"`
            AddressType string  `json:"address_type"`
            Address1 string  `json:"address1"`
            Latitude float64  `json:"latitude"`
            Area string  `json:"area"`
            Address2 string  `json:"address2"`
            Pincode float64  `json:"pincode"`
            AddressCategory string  `json:"address_category"`
            Version string  `json:"version"`
            ContactPerson string  `json:"contact_person"`
         
    }
    
    // Store used by Order
    type Store struct {

        
            Code string  `json:"code"`
            StoreEmail string  `json:"store_email"`
            MallName string  `json:"mall_name"`
            PackagingMaterialCount float64  `json:"packaging_material_count"`
            State string  `json:"state"`
            ParentStoreID float64  `json:"parent_store_id"`
            UpdatedAt string  `json:"updated_at"`
            CreatedAt string  `json:"created_at"`
            Phone float64  `json:"phone"`
            IsEnabledForRecon bool  `json:"is_enabled_for_recon"`
            IsActive bool  `json:"is_active"`
            LocationType string  `json:"location_type"`
            City string  `json:"city"`
            Country string  `json:"country"`
            StoreActiveFrom string  `json:"store_active_from"`
            Longitude float64  `json:"longitude"`
            VatNo string  `json:"vat_no"`
            SID string  `json:"s_id"`
            IsArchived bool  `json:"is_archived"`
            AlohomoraUserID float64  `json:"alohomora_user_id"`
            MallArea string  `json:"mall_area"`
            Address1 string  `json:"address1"`
            Meta StoreMeta  `json:"meta"`
            LoginUsername string  `json:"login_username"`
            CompanyID float64  `json:"company_id"`
            Latitude float64  `json:"latitude"`
            BrandID interface{}  `json:"brand_id"`
            Name string  `json:"name"`
            StoreAddressJson StoreAddress  `json:"store_address_json"`
            Address2 string  `json:"address2"`
            Pincode string  `json:"pincode"`
            FulfillmentChannel string  `json:"fulfillment_channel"`
            BrandStoreTags []string  `json:"brand_store_tags"`
            OrderIntegrationID string  `json:"order_integration_id"`
            ContactPerson string  `json:"contact_person"`
         
    }
    
    // ArticleDetails used by Order
    type ArticleDetails struct {

        
            Status map[string]interface{}  `json:"status"`
         
    }
    
    // BagDetailsPlatformResponse used by Order
    type BagDetailsPlatformResponse struct {

        
            Identifier string  `json:"identifier"`
            BID float64  `json:"b_id"`
            CurrentOperationalStatus BagStatusHistory  `json:"current_operational_status"`
            Quantity float64  `json:"quantity"`
            RestoreCoupon bool  `json:"restore_coupon"`
            GstDetails BagGSTDetails  `json:"gst_details"`
            Prices Prices  `json:"prices"`
            OriginalBagList []float64  `json:"original_bag_list"`
            ParentPromoBags map[string]interface{}  `json:"parent_promo_bags"`
            BagStatus []BagStatusHistory  `json:"bag_status"`
            EntityType string  `json:"entity_type"`
            Brand Brand  `json:"brand"`
            RestorePromos map[string]interface{}  `json:"restore_promos"`
            JourneyType string  `json:"journey_type"`
            DisplayName string  `json:"display_name"`
            LineNumber float64  `json:"line_number"`
            BType string  `json:"b_type"`
            AffiliateDetails AffiliateDetails  `json:"affiliate_details"`
            Item Item  `json:"item"`
            Dates Dates  `json:"dates"`
            Reasons []map[string]interface{}  `json:"reasons"`
            NoOfBagsOrder float64  `json:"no_of_bags_order"`
            FinancialBreakup []FinancialBreakup  `json:"financial_breakup"`
            Meta BagMeta  `json:"meta"`
            BagUpdateTime float64  `json:"bag_update_time"`
            OrderingStore Store  `json:"ordering_store"`
            QcRequired interface{}  `json:"qc_required"`
            OperationalStatus string  `json:"operational_status"`
            AppliedPromos []map[string]interface{}  `json:"applied_promos"`
            ShipmentID string  `json:"shipment_id"`
            Article Article  `json:"article"`
            ArticleDetails ArticleDetails  `json:"article_details"`
            BagStatusHistory BagStatusHistory  `json:"bag_status_history"`
            AffiliateBagDetails AffiliateBagDetails  `json:"affiliate_bag_details"`
            SellerIdentifier string  `json:"seller_identifier"`
            Status BagReturnableCancelableStatus  `json:"status"`
            CurrentStatus BagStatusHistory  `json:"current_status"`
            Tags []string  `json:"tags"`
            OrderIntegrationID string  `json:"order_integration_id"`
         
    }
    
    // ErrorResponse used by Order
    type ErrorResponse struct {

        
            Error string  `json:"error"`
            Message string  `json:"message"`
         
    }
    
    // Page1 used by Order
    type Page1 struct {

        
            Size float64  `json:"size"`
            HasNext bool  `json:"has_next"`
            PageType string  `json:"page_type"`
            ItemTotal float64  `json:"item_total"`
            Current float64  `json:"current"`
         
    }
    
    // GetBagsPlatformResponse used by Order
    type GetBagsPlatformResponse struct {

        
            Items []BagDetailsPlatformResponse  `json:"items"`
            Page Page1  `json:"page"`
         
    }
    
    // InvalidateShipmentCachePayload used by Order
    type InvalidateShipmentCachePayload struct {

        
            AffiliateBagIds []string  `json:"affiliate_bag_ids"`
            BagIds []string  `json:"bag_ids"`
            ShipmentIds []string  `json:"shipment_ids"`
         
    }
    
    // InvalidateShipmentCacheNestedResponse used by Order
    type InvalidateShipmentCacheNestedResponse struct {

        
            Message string  `json:"message"`
            Error string  `json:"error"`
            ShipmentID string  `json:"shipment_id"`
            Status float64  `json:"status"`
         
    }
    
    // InvalidateShipmentCacheResponse used by Order
    type InvalidateShipmentCacheResponse struct {

        
            Response []InvalidateShipmentCacheNestedResponse  `json:"response"`
         
    }
    
    // ErrorResponse1 used by Order
    type ErrorResponse1 struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
            ErrorTrace string  `json:"error_trace"`
            Status float64  `json:"status"`
         
    }
    
    // StoreReassign used by Order
    type StoreReassign struct {

        
            ReasonIds []float64  `json:"reason_ids"`
            SetID string  `json:"set_id"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
            StoreID float64  `json:"store_id"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            MongoArticleID string  `json:"mongo_article_id"`
            ItemID string  `json:"item_id"`
            FyndOrderID string  `json:"fynd_order_id"`
            AffiliateID string  `json:"affiliate_id"`
            BagID float64  `json:"bag_id"`
         
    }
    
    // StoreReassignResponse used by Order
    type StoreReassignResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // Entities used by Order
    type Entities struct {

        
            ReasonText string  `json:"reason_text"`
            ID string  `json:"id"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            AffiliateID string  `json:"affiliate_id"`
         
    }
    
    // UpdateShipmentLockPayload used by Order
    type UpdateShipmentLockPayload struct {

        
            ActionType string  `json:"action_type"`
            Action string  `json:"action"`
            Entities []Entities  `json:"entities"`
            EntityType string  `json:"entity_type"`
         
    }
    
    // OriginalFilter used by Order
    type OriginalFilter struct {

        
            AffiliateID string  `json:"affiliate_id"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
         
    }
    
    // Bags used by Order
    type Bags struct {

        
            AffiliateOrderID string  `json:"affiliate_order_id"`
            BagID float64  `json:"bag_id"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
            IsLocked bool  `json:"is_locked"`
         
    }
    
    // CheckResponse used by Order
    type CheckResponse struct {

        
            LockStatus string  `json:"lock_status"`
            OriginalFilter OriginalFilter  `json:"original_filter"`
            IsBagLocked bool  `json:"is_bag_locked"`
            Bags []Bags  `json:"bags"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            IsShipmentLocked bool  `json:"is_shipment_locked"`
            ShipmentID string  `json:"shipment_id"`
            AffiliateID string  `json:"affiliate_id"`
            Status string  `json:"status"`
         
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
            Title string  `json:"title"`
            PlatformID string  `json:"platform_id"`
            CompanyID float64  `json:"company_id"`
            CreatedAt string  `json:"created_at"`
            ID float64  `json:"id"`
            PlatformName string  `json:"platform_name"`
            Description string  `json:"description"`
            ToDatetime string  `json:"to_datetime"`
         
    }
    
    // AnnouncementsResponse used by Order
    type AnnouncementsResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
            Announcements []AnnouncementResponse  `json:"announcements"`
         
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
    
    // ProductsReasonsFilters used by Order
    type ProductsReasonsFilters struct {

        
            Identifier string  `json:"identifier"`
            LineNumber float64  `json:"line_number"`
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
    
    // Products used by Order
    type Products struct {

        
            Identifier string  `json:"identifier"`
            LineNumber float64  `json:"line_number"`
            Quantity float64  `json:"quantity"`
         
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
    
    // ShipmentsRequest used by Order
    type ShipmentsRequest struct {

        
            Identifier string  `json:"identifier"`
            Reasons ReasonsData  `json:"reasons"`
            Products []Products  `json:"products"`
            DataUpdates DataUpdates  `json:"data_updates"`
         
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
            ForceTransition bool  `json:"force_transition"`
            Task bool  `json:"task"`
            Statuses []StatuesRequest  `json:"statuses"`
            LockAfterTransition bool  `json:"lock_after_transition"`
         
    }
    
    // ShipmentsResponse used by Order
    type ShipmentsResponse struct {

        
            Identifier string  `json:"identifier"`
            Code string  `json:"code"`
            FinalState map[string]interface{}  `json:"final_state"`
            Message string  `json:"message"`
            Exception string  `json:"exception"`
            StackTrace string  `json:"stack_trace"`
            Meta map[string]interface{}  `json:"meta"`
            Status float64  `json:"status"`
         
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

        
            Email string  `json:"email"`
            Country string  `json:"country"`
            State string  `json:"state"`
            Address1 string  `json:"address1"`
            Pincode string  `json:"pincode"`
            Mobile float64  `json:"mobile"`
            FirstName string  `json:"first_name"`
            City string  `json:"city"`
            LastName string  `json:"last_name"`
            Address2 string  `json:"address2"`
            Phone float64  `json:"phone"`
         
    }
    
    // UserData used by Order
    type UserData struct {

        
            ShippingUser OrderUser  `json:"shipping_user"`
            BillingUser OrderUser  `json:"billing_user"`
         
    }
    
    // ArticleDetails1 used by Order
    type ArticleDetails1 struct {

        
            Category map[string]interface{}  `json:"category"`
            BrandID float64  `json:"brand_id"`
            Quantity float64  `json:"quantity"`
            Dimension map[string]interface{}  `json:"dimension"`
            Weight map[string]interface{}  `json:"weight"`
            Attributes map[string]interface{}  `json:"attributes"`
            ID string  `json:"_id"`
         
    }
    
    // ShipmentDetails1 used by Order
    type ShipmentDetails1 struct {

        
            Articles []ArticleDetails1  `json:"articles"`
            BoxType string  `json:"box_type"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            Shipments float64  `json:"shipments"`
            DpID float64  `json:"dp_id"`
            FulfillmentID float64  `json:"fulfillment_id"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // LocationDetails used by Order
    type LocationDetails struct {

        
            Articles []ArticleDetails1  `json:"articles"`
            FulfillmentID float64  `json:"fulfillment_id"`
            FulfillmentType string  `json:"fulfillment_type"`
         
    }
    
    // ShipmentConfig used by Order
    type ShipmentConfig struct {

        
            Identifier string  `json:"identifier"`
            PaymentMode string  `json:"payment_mode"`
            Journey string  `json:"journey"`
            Source string  `json:"source"`
            Shipment []ShipmentDetails1  `json:"shipment"`
            Action string  `json:"action"`
            ToPincode string  `json:"to_pincode"`
            LocationDetails LocationDetails  `json:"location_details"`
         
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

        
            SellerIdentifier string  `json:"seller_identifier"`
            TransferPrice float64  `json:"transfer_price"`
            StoreID float64  `json:"store_id"`
            PriceMarked float64  `json:"price_marked"`
            Identifier map[string]interface{}  `json:"identifier"`
            DeliveryCharge float64  `json:"delivery_charge"`
            AmountPaid float64  `json:"amount_paid"`
            PriceEffective float64  `json:"price_effective"`
            PdfLinks MarketPlacePdf  `json:"pdf_links"`
            AffiliateMeta map[string]interface{}  `json:"affiliate_meta"`
            AffiliateStoreID string  `json:"affiliate_store_id"`
            Quantity float64  `json:"quantity"`
            AvlQty float64  `json:"avl_qty"`
            CompanyID float64  `json:"company_id"`
            ItemSize string  `json:"item_size"`
            HsnCodeID string  `json:"hsn_code_id"`
            ModifiedOn string  `json:"modified_on"`
            ItemID float64  `json:"item_id"`
            Discount float64  `json:"discount"`
            ID string  `json:"_id"`
            Sku string  `json:"sku"`
            FyndStoreID string  `json:"fynd_store_id"`
            UnitPrice float64  `json:"unit_price"`
         
    }
    
    // OrderPriority used by Order
    type OrderPriority struct {

        
            FulfilmentPriorityText string  `json:"fulfilment_priority_text"`
            AffiliatePriorityCode string  `json:"affiliate_priority_code"`
            FulfilmentPriority float64  `json:"fulfilment_priority"`
         
    }
    
    // OrderInfo used by Order
    type OrderInfo struct {

        
            PaymentMode string  `json:"payment_mode"`
            BillingAddress OrderUser  `json:"billing_address"`
            DeliveryCharges float64  `json:"delivery_charges"`
            User UserData  `json:"user"`
            Shipment ShipmentData  `json:"shipment"`
            Bags []AffiliateBag  `json:"bags"`
            ShippingAddress OrderUser  `json:"shipping_address"`
            Items map[string]interface{}  `json:"items"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            Coupon string  `json:"coupon"`
            Payment map[string]interface{}  `json:"payment"`
            Discount float64  `json:"discount"`
            OrderPriority OrderPriority  `json:"order_priority"`
            OrderValue float64  `json:"order_value"`
            CodCharges float64  `json:"cod_charges"`
         
    }
    
    // AffiliateStoreIdMapping used by Order
    type AffiliateStoreIdMapping struct {

        
            MarketplaceStoreID string  `json:"marketplace_store_id"`
            StoreID float64  `json:"store_id"`
         
    }
    
    // AffiliateAppConfigMeta used by Order
    type AffiliateAppConfigMeta struct {

        
            Name string  `json:"name"`
            Value string  `json:"value"`
         
    }
    
    // AffiliateAppConfig used by Order
    type AffiliateAppConfig struct {

        
            Token string  `json:"token"`
            Name string  `json:"name"`
            CreatedAt string  `json:"created_at"`
            ID string  `json:"id"`
            Owner string  `json:"owner"`
            UpdatedAt string  `json:"updated_at"`
            Secret string  `json:"secret"`
            Meta []AffiliateAppConfigMeta  `json:"meta"`
            Description string  `json:"description"`
         
    }
    
    // AffiliateInventoryLogisticsConfig used by Order
    type AffiliateInventoryLogisticsConfig struct {

        
            DpAssignment bool  `json:"dp_assignment"`
         
    }
    
    // AffiliateInventoryArticleAssignmentConfig used by Order
    type AffiliateInventoryArticleAssignmentConfig struct {

        
            PostOrderReassignment bool  `json:"post_order_reassignment"`
         
    }
    
    // AffiliateInventoryStoreConfig used by Order
    type AffiliateInventoryStoreConfig struct {

        
            Store map[string]interface{}  `json:"store"`
         
    }
    
    // AffiliateInventoryPaymentConfig used by Order
    type AffiliateInventoryPaymentConfig struct {

        
            Source string  `json:"source"`
            ModeOfPayment string  `json:"mode_of_payment"`
         
    }
    
    // AffiliateInventoryOrderConfig used by Order
    type AffiliateInventoryOrderConfig struct {

        
            ForceReassignment bool  `json:"force_reassignment"`
         
    }
    
    // AffiliateInventoryConfig used by Order
    type AffiliateInventoryConfig struct {

        
            Logistics AffiliateInventoryLogisticsConfig  `json:"logistics"`
            ArticleAssignment AffiliateInventoryArticleAssignmentConfig  `json:"article_assignment"`
            Inventory AffiliateInventoryStoreConfig  `json:"inventory"`
            Payment AffiliateInventoryPaymentConfig  `json:"payment"`
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
    
    // OrderConfig used by Order
    type OrderConfig struct {

        
            CreateUser bool  `json:"create_user"`
            AffiliateStoreIDMapping []AffiliateStoreIdMapping  `json:"affiliate_store_id_mapping"`
            ArticleLookup string  `json:"article_lookup"`
            Affiliate Affiliate  `json:"affiliate"`
            BagEndState string  `json:"bag_end_state"`
            StoreLookup string  `json:"store_lookup"`
         
    }
    
    // CreateOrderPayload used by Order
    type CreateOrderPayload struct {

        
            OrderInfo OrderInfo  `json:"order_info"`
            OrderConfig OrderConfig  `json:"order_config"`
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

        
            DisplayText string  `json:"display_text"`
            ID float64  `json:"id"`
            Description string  `json:"description"`
            Slug string  `json:"slug"`
         
    }
    
    // GetActionsResponse used by Order
    type GetActionsResponse struct {

        
            Permissions []ActionInfo  `json:"permissions"`
         
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
    
    // HistoryReason used by Order
    type HistoryReason struct {

        
            DislayName string  `json:"dislay_name"`
            Category string  `json:"category"`
            Quantity float64  `json:"quantity"`
            State string  `json:"state"`
            Code float64  `json:"code"`
            Text string  `json:"text"`
         
    }
    
    // HistoryMeta used by Order
    type HistoryMeta struct {

        
            ShortLink string  `json:"short_link"`
            Reason HistoryReason  `json:"reason"`
            StoreName string  `json:"store_name"`
            StoreID float64  `json:"store_id"`
            Callerid string  `json:"callerid"`
            Recipient string  `json:"recipient"`
            Status1 string  `json:"status1"`
            Billsec string  `json:"billsec"`
            ActivityType string  `json:"activity_type"`
            Message string  `json:"message"`
            Caller string  `json:"caller"`
            Endtime string  `json:"endtime"`
            ChannelType string  `json:"channel_type"`
            Status string  `json:"status"`
            Starttime string  `json:"starttime"`
            Status2 string  `json:"status2"`
            Slug string  `json:"slug"`
            StoreCode string  `json:"store_code"`
            Recordpath string  `json:"recordpath"`
            Duration string  `json:"duration"`
            CallID string  `json:"call_id"`
            Receiver string  `json:"receiver"`
            ActivityComment string  `json:"activity_comment"`
         
    }
    
    // HistoryDict used by Order
    type HistoryDict struct {

        
            L1Detail string  `json:"l1_detail"`
            Type string  `json:"type"`
            TicketURL string  `json:"ticket_url"`
            User string  `json:"user"`
            L3Detail string  `json:"l3_detail"`
            DisplayMessage string  `json:"display_message"`
            Createdat string  `json:"createdat"`
            AssignedAgent string  `json:"assigned_agent"`
            Message string  `json:"message"`
            TicketID string  `json:"ticket_id"`
            Meta HistoryMeta  `json:"meta"`
            L2Detail string  `json:"l2_detail"`
            BagID float64  `json:"bag_id"`
         
    }
    
    // ShipmentHistoryResponse used by Order
    type ShipmentHistoryResponse struct {

        
            Success bool  `json:"success"`
            ActivityHistory []HistoryDict  `json:"activity_history"`
         
    }
    
    // SmsDataPayload used by Order
    type SmsDataPayload struct {

        
            PaymentMode string  `json:"payment_mode"`
            AmountPaid float64  `json:"amount_paid"`
            CustomerName string  `json:"customer_name"`
            OrderID string  `json:"order_id"`
            PhoneNumber float64  `json:"phone_number"`
            Message string  `json:"message"`
            ShipmentID float64  `json:"shipment_id"`
            CountryCode string  `json:"country_code"`
            BrandName string  `json:"brand_name"`
         
    }
    
    // SendSmsPayload used by Order
    type SendSmsPayload struct {

        
            BagID float64  `json:"bag_id"`
            Slug string  `json:"slug"`
            Data SmsDataPayload  `json:"data"`
         
    }
    
    // Meta1 used by Order
    type Meta1 struct {

        
            KafkaEmissionStatus float64  `json:"kafka_emission_status"`
            StateManagerUsed string  `json:"state_manager_used"`
         
    }
    
    // ShipmentDetail used by Order
    type ShipmentDetail struct {

        
            Remarks string  `json:"remarks"`
            ID float64  `json:"id"`
            ShipmentID string  `json:"shipment_id"`
            Status string  `json:"status"`
            Meta Meta1  `json:"meta"`
            BagList []float64  `json:"bag_list"`
         
    }
    
    // OrderDetails used by Order
    type OrderDetails struct {

        
            CreatedAt string  `json:"created_at"`
            FyndOrderID string  `json:"fynd_order_id"`
         
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
    
    // Dimension used by Order
    type Dimension struct {

        
            PackagingType string  `json:"packaging_type"`
            Height string  `json:"height"`
            Width float64  `json:"width"`
            Length float64  `json:"length"`
            Weight string  `json:"weight"`
         
    }
    
    // UpdatePackagingDimensionsPayload used by Order
    type UpdatePackagingDimensionsPayload struct {

        
            CurrentStatus string  `json:"current_status"`
            ShipmentID string  `json:"shipment_id"`
            Dimension []Dimension  `json:"dimension"`
         
    }
    
    // UpdatePackagingDimensionsResponse used by Order
    type UpdatePackagingDimensionsResponse struct {

        
            Message string  `json:"message"`
         
    }
    
    // BillingInfo used by Order
    type BillingInfo struct {

        
            HouseNo string  `json:"house_no"`
            State string  `json:"state"`
            Address1 string  `json:"address1"`
            Title string  `json:"title"`
            MiddleName string  `json:"middle_name"`
            Pincode string  `json:"pincode"`
            FloorNo string  `json:"floor_no"`
            City string  `json:"city"`
            CountryCode string  `json:"country_code"`
            PrimaryEmail string  `json:"primary_email"`
            StateCode string  `json:"state_code"`
            Gender string  `json:"gender"`
            Address2 string  `json:"address2"`
            AlternateEmail string  `json:"alternate_email"`
            Country string  `json:"country"`
            AlternateMobileNumber string  `json:"alternate_mobile_number"`
            CustomerCode string  `json:"customer_code"`
            FirstName string  `json:"first_name"`
            ExternalCustomerCode string  `json:"external_customer_code"`
            LastName string  `json:"last_name"`
            PrimaryMobileNumber string  `json:"primary_mobile_number"`
         
    }
    
    // Tax used by Order
    type Tax struct {

        
            Name string  `json:"name"`
            Amount map[string]interface{}  `json:"amount"`
            Breakup []map[string]interface{}  `json:"breakup"`
            Rate float64  `json:"rate"`
         
    }
    
    // Charge used by Order
    type Charge struct {

        
            Type string  `json:"type"`
            Name string  `json:"name"`
            Code string  `json:"code"`
            Amount map[string]interface{}  `json:"amount"`
            Tax Tax  `json:"tax"`
         
    }
    
    // ShippingInfo used by Order
    type ShippingInfo struct {

        
            HouseNo string  `json:"house_no"`
            State string  `json:"state"`
            Address1 string  `json:"address1"`
            Title string  `json:"title"`
            MiddleName string  `json:"middle_name"`
            Pincode string  `json:"pincode"`
            Slot []map[string]interface{}  `json:"slot"`
            Landmark string  `json:"landmark"`
            FloorNo string  `json:"floor_no"`
            City string  `json:"city"`
            ShippingType string  `json:"shipping_type"`
            CountryCode string  `json:"country_code"`
            PrimaryEmail string  `json:"primary_email"`
            AddressType string  `json:"address_type"`
            StateCode string  `json:"state_code"`
            Gender string  `json:"gender"`
            Address2 string  `json:"address2"`
            AlternateEmail string  `json:"alternate_email"`
            Country string  `json:"country"`
            AlternateMobileNumber string  `json:"alternate_mobile_number"`
            GeoLocation map[string]interface{}  `json:"geo_location"`
            CustomerCode string  `json:"customer_code"`
            FirstName string  `json:"first_name"`
            ExternalCustomerCode string  `json:"external_customer_code"`
            LastName string  `json:"last_name"`
            PrimaryMobileNumber string  `json:"primary_mobile_number"`
         
    }
    
    // PaymentMethod1 used by Order
    type PaymentMethod1 struct {

        
            Name string  `json:"name"`
            RefundBy string  `json:"refund_by"`
            Mode string  `json:"mode"`
            CollectBy string  `json:"collect_by"`
            Amount float64  `json:"amount"`
            TransactionData map[string]interface{}  `json:"transaction_data"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // PaymentInfo used by Order
    type PaymentInfo struct {

        
            PaymentMethods []PaymentMethod1  `json:"payment_methods"`
            PrimaryMode string  `json:"primary_mode"`
         
    }
    
    // LineItem used by Order
    type LineItem struct {

        
            Charges []Charge  `json:"charges"`
            Quantity float64  `json:"quantity"`
            SellerIdentifier string  `json:"seller_identifier"`
            CustomMessasge string  `json:"custom_messasge"`
            ExternalLineID string  `json:"external_line_id"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // ProcessingDates used by Order
    type ProcessingDates struct {

        
            PackByDate string  `json:"pack_by_date"`
            CustomerPickupSlot map[string]interface{}  `json:"customer_pickup_slot"`
            DispatchAfterDate string  `json:"dispatch_after_date"`
            DpPickupSlot map[string]interface{}  `json:"dp_pickup_slot"`
            ConfirmByDate string  `json:"confirm_by_date"`
            DispatchByDate string  `json:"dispatch_by_date"`
         
    }
    
    // Shipment used by Order
    type Shipment struct {

        
            LocationID float64  `json:"location_id"`
            ExternalShipmentID string  `json:"external_shipment_id"`
            LineItems []LineItem  `json:"line_items"`
            Priority float64  `json:"priority"`
            Meta map[string]interface{}  `json:"meta"`
            ProcessingDates ProcessingDates  `json:"processing_dates"`
         
    }
    
    // TaxInfo used by Order
    type TaxInfo struct {

        
            B2bGstinNumber string  `json:"b2b_gstin_number"`
            Gstin string  `json:"gstin"`
         
    }
    
    // CreateOrderAPI used by Order
    type CreateOrderAPI struct {

        
            BillingInfo BillingInfo  `json:"billing_info"`
            Charges []Charge  `json:"charges"`
            ShippingInfo ShippingInfo  `json:"shipping_info"`
            CurrencyInfo map[string]interface{}  `json:"currency_info"`
            Config map[string]interface{}  `json:"config"`
            ExternalOrderID string  `json:"external_order_id"`
            ExternalCreationDate string  `json:"external_creation_date"`
            PaymentInfo PaymentInfo  `json:"payment_info"`
            Shipments []Shipment  `json:"shipments"`
            Meta map[string]interface{}  `json:"meta"`
            TaxInfo TaxInfo  `json:"tax_info"`
         
    }
    
    // CreateOrderErrorReponse used by Order
    type CreateOrderErrorReponse struct {

        
            RequestID string  `json:"request_id"`
            Code string  `json:"code"`
            Info interface{}  `json:"info"`
            Message string  `json:"message"`
            Exception string  `json:"exception"`
            StackTrace string  `json:"stack_trace"`
            Meta string  `json:"meta"`
            Status float64  `json:"status"`
         
    }
    
    // DpConfiguration used by Order
    type DpConfiguration struct {

        
            ShippingBy string  `json:"shipping_by"`
         
    }
    
    // PaymentMethods used by Order
    type PaymentMethods struct {

        
            CollectBy string  `json:"collect_by"`
            RefundBy string  `json:"refund_by"`
            Mode string  `json:"mode"`
         
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
            LogoURL map[string]interface{}  `json:"logo_url"`
            LocationReassignment bool  `json:"location_reassignment"`
            LockStates []string  `json:"lock_states"`
            ShipmentAssignment string  `json:"shipment_assignment"`
            PaymentInfo CreateChannelPaymentInfo  `json:"payment_info"`
         
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

        
            Mobile float64  `json:"mobile"`
            OrderDetails []FyndOrderIdList  `json:"order_details"`
            EndDate string  `json:"end_date"`
            StartDate string  `json:"start_date"`
         
    }
    
    // BagStateTransitionMap used by Order
    type BagStateTransitionMap struct {

        
            Affiliate map[string]interface{}  `json:"affiliate"`
            Fynd map[string]interface{}  `json:"fynd"`
         
    }
    

    
    // DeleteResponse used by Catalog
    type DeleteResponse struct {

        
            Message string  `json:"message"`
         
    }
    
    // BaseErrorResponse used by Catalog
    type BaseErrorResponse struct {

        
            Errors map[string]interface{}  `json:"errors"`
            Code float64  `json:"code"`
            Message string  `json:"message"`
         
    }
    
    // SearchKeywordResult used by Catalog
    type SearchKeywordResult struct {

        
            Query map[string]interface{}  `json:"query"`
            SortOn string  `json:"sort_on"`
         
    }
    
    // GetSearchWordsData used by Catalog
    type GetSearchWordsData struct {

        
            Result SearchKeywordResult  `json:"result"`
            IsActive bool  `json:"is_active"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Words []string  `json:"words"`
            AppID string  `json:"app_id"`
            UID string  `json:"uid"`
         
    }
    
    // SearchErrorResponse used by Catalog
    type SearchErrorResponse struct {

        
            Errors []map[string]interface{}  `json:"errors"`
            Code float64  `json:"code"`
            Message string  `json:"message"`
         
    }
    
    // SearchKeywordResult1 used by Catalog
    type SearchKeywordResult1 struct {

        
            Query map[string]interface{}  `json:"query"`
            SortOn string  `json:"sort_on"`
         
    }
    
    // CreateSearchKeyword used by Catalog
    type CreateSearchKeyword struct {

        
            Result SearchKeywordResult1  `json:"result"`
            IsActive bool  `json:"is_active"`
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
    
    // GetSearchWordsResponse used by Catalog
    type GetSearchWordsResponse struct {

        
            Page Page  `json:"page"`
            Items []GetSearchWordsData  `json:"items"`
         
    }
    
    // AutocompletePageAction used by Catalog
    type AutocompletePageAction struct {

        
            Type string  `json:"type"`
            Query map[string]interface{}  `json:"query"`
            Params map[string]interface{}  `json:"params"`
            URL string  `json:"url"`
         
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
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Action AutocompleteAction  `json:"action"`
            Logo Media  `json:"logo"`
         
    }
    
    // GetAutocompleteWordsData used by Catalog
    type GetAutocompleteWordsData struct {

        
            IsActive bool  `json:"is_active"`
            Results []AutocompleteResult  `json:"results"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Words []string  `json:"words"`
            AppID string  `json:"app_id"`
            UID string  `json:"uid"`
         
    }
    
    // UpdateAutocompleteWordData used by Catalog
    type UpdateAutocompleteWordData struct {

        
            ID string  `json:"_id"`
            IsActive bool  `json:"is_active"`
            Results []AutocompleteResult  `json:"results"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Words []string  `json:"words"`
            AppID string  `json:"app_id"`
            UID string  `json:"uid"`
         
    }
    
    // GetAutocompleteWordsResponse used by Catalog
    type GetAutocompleteWordsResponse struct {

        
            Page Page  `json:"page"`
            Items []GetAutocompleteWordsData  `json:"items"`
         
    }
    
    // AutocompletePageAction1 used by Catalog
    type AutocompletePageAction1 struct {

        
            Type string  `json:"type"`
            Query map[string]interface{}  `json:"query"`
            Params map[string]interface{}  `json:"params"`
            URL string  `json:"url"`
         
    }
    
    // AutocompleteAction1 used by Catalog
    type AutocompleteAction1 struct {

        
            Type string  `json:"type"`
            Page AutocompletePageAction1  `json:"page"`
         
    }
    
    // Media1 used by Catalog
    type Media1 struct {

        
            Type string  `json:"type"`
            URL string  `json:"url"`
            AspectRatio string  `json:"aspect_ratio"`
         
    }
    
    // AutocompleteResult1 used by Catalog
    type AutocompleteResult1 struct {

        
            Display string  `json:"display"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Action AutocompleteAction1  `json:"action"`
            Logo Media1  `json:"logo"`
         
    }
    
    // CreateAutocompleteKeyword used by Catalog
    type CreateAutocompleteKeyword struct {

        
            IsActive bool  `json:"is_active"`
            Results []AutocompleteResult1  `json:"results"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Words []string  `json:"words"`
            AppID string  `json:"app_id"`
         
    }
    
    // UserDetail used by Catalog
    type UserDetail struct {

        
            UserID string  `json:"user_id"`
            Contact string  `json:"contact"`
            Username string  `json:"username"`
            SuperUser bool  `json:"super_user"`
         
    }
    
    // RerankingAttribute used by Catalog
    type RerankingAttribute struct {

        
            AttributeKey string  `json:"attribute_key"`
            AttributeValue string  `json:"attribute_value"`
         
    }
    
    // BoostBury used by Catalog
    type BoostBury struct {

        
            Boost []RerankingAttribute  `json:"boost"`
         
    }
    
    // SearchRerankingModel used by Catalog
    type SearchRerankingModel struct {

        
            ID interface{}  `json:"_id"`
            ModifiedBy UserDetail  `json:"modified_by"`
            CreatedBy UserDetail  `json:"created_by"`
            VerifiedBy UserDetail  `json:"verified_by"`
            CreatedOn string  `json:"created_on"`
            IsActive bool  `json:"is_active"`
            Ranking BoostBury  `json:"ranking"`
            Words []string  `json:"words"`
            AppID string  `json:"app_id"`
            VerifiedOn string  `json:"verified_on"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // SearchRerankListing used by Catalog
    type SearchRerankListing struct {

        
            Page Page  `json:"page"`
            Items []SearchRerankingModel  `json:"items"`
         
    }
    
    // RerankingAttribute1 used by Catalog
    type RerankingAttribute1 struct {

        
            AttributeKey string  `json:"attribute_key"`
            AttributeValue string  `json:"attribute_value"`
         
    }
    
    // BoostBury1 used by Catalog
    type BoostBury1 struct {

        
            Boost []RerankingAttribute1  `json:"boost"`
         
    }
    
    // CreateSearchReranking used by Catalog
    type CreateSearchReranking struct {

        
            Ranking BoostBury1  `json:"ranking"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            IsActive bool  `json:"is_active"`
            CreatedOn string  `json:"created_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Words []string  `json:"words"`
            AppID string  `json:"app_id"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // DeleteRerankResponse used by Catalog
    type DeleteRerankResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // ProductBundleItem used by Catalog
    type ProductBundleItem struct {

        
            AutoSelect bool  `json:"auto_select"`
            MaxQuantity float64  `json:"max_quantity"`
            ProductUID float64  `json:"product_uid"`
            MinQuantity float64  `json:"min_quantity"`
            AllowRemove bool  `json:"allow_remove"`
            AutoAddToCart bool  `json:"auto_add_to_cart"`
         
    }
    
    // GetProductBundleCreateResponse used by Catalog
    type GetProductBundleCreateResponse struct {

        
            ModifiedOn string  `json:"modified_on"`
            CompanyID float64  `json:"company_id"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            Name string  `json:"name"`
            Products []ProductBundleItem  `json:"products"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Slug string  `json:"slug"`
            IsActive bool  `json:"is_active"`
            Logo string  `json:"logo"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            CreatedOn string  `json:"created_on"`
            ID string  `json:"id"`
            Choice string  `json:"choice"`
            Meta map[string]interface{}  `json:"meta"`
            PageVisibility []string  `json:"page_visibility"`
         
    }
    
    // GetProductBundleListingResponse used by Catalog
    type GetProductBundleListingResponse struct {

        
            Page Page  `json:"page"`
            Items []GetProductBundleCreateResponse  `json:"items"`
         
    }
    
    // ErrorResponse used by Catalog
    type ErrorResponse struct {

        
            Code string  `json:"code"`
            Status float64  `json:"status"`
            Error string  `json:"error"`
            Meta map[string]interface{}  `json:"meta"`
            Message string  `json:"message"`
         
    }
    
    // ProductBundleRequest used by Catalog
    type ProductBundleRequest struct {

        
            ModifiedOn string  `json:"modified_on"`
            CompanyID float64  `json:"company_id"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            Name string  `json:"name"`
            Products []ProductBundleItem  `json:"products"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Slug string  `json:"slug"`
            IsActive bool  `json:"is_active"`
            Logo string  `json:"logo"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            CreatedOn string  `json:"created_on"`
            Choice string  `json:"choice"`
            Meta map[string]interface{}  `json:"meta"`
            PageVisibility []string  `json:"page_visibility"`
         
    }
    
    // ProductBundleUpdateRequest used by Catalog
    type ProductBundleUpdateRequest struct {

        
            ModifiedOn string  `json:"modified_on"`
            CompanyID float64  `json:"company_id"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            Name string  `json:"name"`
            Products []ProductBundleItem  `json:"products"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Slug string  `json:"slug"`
            IsActive bool  `json:"is_active"`
            Logo string  `json:"logo"`
            Choice string  `json:"choice"`
            Meta map[string]interface{}  `json:"meta"`
            PageVisibility []string  `json:"page_visibility"`
         
    }
    
    // Price used by Catalog
    type Price struct {

        
            MaxEffective float64  `json:"max_effective"`
            Currency string  `json:"currency"`
            MaxMarked float64  `json:"max_marked"`
            MinEffective float64  `json:"min_effective"`
            MinMarked float64  `json:"min_marked"`
         
    }
    
    // Size used by Catalog
    type Size struct {

        
            Display string  `json:"display"`
            Quantity float64  `json:"quantity"`
            Value string  `json:"value"`
            IsAvailable bool  `json:"is_available"`
         
    }
    
    // LimitedProductData used by Catalog
    type LimitedProductData struct {

        
            Price map[string]interface{}  `json:"price"`
            Name string  `json:"name"`
            Quantity float64  `json:"quantity"`
            Slug string  `json:"slug"`
            ItemCode string  `json:"item_code"`
            Sizes []string  `json:"sizes"`
            ShortDescription string  `json:"short_description"`
            Images []string  `json:"images"`
            UID float64  `json:"uid"`
            Attributes map[string]interface{}  `json:"attributes"`
            CountryOfOrigin string  `json:"country_of_origin"`
            Identifier map[string]interface{}  `json:"identifier"`
         
    }
    
    // GetProducts used by Catalog
    type GetProducts struct {

        
            Price Price  `json:"price"`
            AutoSelect bool  `json:"auto_select"`
            MaxQuantity float64  `json:"max_quantity"`
            ProductUID float64  `json:"product_uid"`
            AutoAddToCart bool  `json:"auto_add_to_cart"`
            Sizes []Size  `json:"sizes"`
            AllowRemove bool  `json:"allow_remove"`
            MinQuantity float64  `json:"min_quantity"`
            ProductDetails LimitedProductData  `json:"product_details"`
         
    }
    
    // GetProductBundleResponse used by Catalog
    type GetProductBundleResponse struct {

        
            CompanyID float64  `json:"company_id"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            Name string  `json:"name"`
            Products []GetProducts  `json:"products"`
            Slug string  `json:"slug"`
            IsActive bool  `json:"is_active"`
            Logo string  `json:"logo"`
            Choice string  `json:"choice"`
            Meta map[string]interface{}  `json:"meta"`
            PageVisibility []string  `json:"page_visibility"`
         
    }
    
    // ListSizeGuide used by Catalog
    type ListSizeGuide struct {

        
            Page map[string]interface{}  `json:"page"`
            Items []map[string]interface{}  `json:"items"`
         
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

        
            CompanyID float64  `json:"company_id"`
            Name string  `json:"name"`
            Active bool  `json:"active"`
            Image string  `json:"image"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            BrandID float64  `json:"brand_id"`
            CreatedOn string  `json:"created_on"`
            ID string  `json:"id"`
            Title string  `json:"title"`
            Subtitle string  `json:"subtitle"`
            Tag string  `json:"tag"`
            ModifiedOn string  `json:"modified_on"`
            Description string  `json:"description"`
            Guide Guide  `json:"guide"`
         
    }
    
    // SuccessResponse used by Catalog
    type SuccessResponse struct {

        
            Success bool  `json:"success"`
         
    }
    
    // SizeGuideResponse used by Catalog
    type SizeGuideResponse struct {

        
            CompanyID float64  `json:"company_id"`
            Name string  `json:"name"`
            Active bool  `json:"active"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            BrandID float64  `json:"brand_id"`
            CreatedOn string  `json:"created_on"`
            ID string  `json:"id"`
            Title string  `json:"title"`
            Subtitle string  `json:"subtitle"`
            Tag string  `json:"tag"`
            ModifiedOn string  `json:"modified_on"`
            Guide map[string]interface{}  `json:"guide"`
         
    }
    
    // ApplicationItemSEO used by Catalog
    type ApplicationItemSEO struct {

        
            Title string  `json:"title"`
            Description string  `json:"description"`
         
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

        
            IsGift bool  `json:"is_gift"`
            Seo ApplicationItemSEO  `json:"seo"`
            Moq ApplicationItemMOQ  `json:"moq"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            AltText map[string]interface{}  `json:"alt_text"`
            IsCod bool  `json:"is_cod"`
            CustomMeta []MetaFields  `json:"_custom_meta"`
         
    }
    
    // SuccessResponse1 used by Catalog
    type SuccessResponse1 struct {

        
            Success bool  `json:"success"`
            UID float64  `json:"uid"`
         
    }
    
    // SEOData used by Catalog
    type SEOData struct {

        
            Title string  `json:"title"`
            Description string  `json:"description"`
         
    }
    
    // MOQData used by Catalog
    type MOQData struct {

        
            Minimum float64  `json:"minimum"`
            Maximum float64  `json:"maximum"`
            IncrementUnit float64  `json:"increment_unit"`
         
    }
    
    // OwnerAppItemResponse used by Catalog
    type OwnerAppItemResponse struct {

        
            IsGift bool  `json:"is_gift"`
            Seo SEOData  `json:"seo"`
            Moq MOQData  `json:"moq"`
            AltText map[string]interface{}  `json:"alt_text"`
            IsCod bool  `json:"is_cod"`
         
    }
    
    // GetConfigMetadataResponse used by Catalog
    type GetConfigMetadataResponse struct {

        
            Condition []map[string]interface{}  `json:"condition"`
            Values []map[string]interface{}  `json:"values"`
            Data []map[string]interface{}  `json:"data"`
         
    }
    
    // PageResponseType used by Catalog
    type PageResponseType struct {

        
            Next float64  `json:"next"`
            HasNext bool  `json:"has_next"`
            Current float64  `json:"current"`
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

        
            Priority float64  `json:"priority"`
            DisplayType string  `json:"display_type"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            IsActive bool  `json:"is_active"`
            Logo string  `json:"logo"`
            Unit string  `json:"unit"`
            Key string  `json:"key"`
         
    }
    
    // AppConfigurationDetail used by Catalog
    type AppConfigurationDetail struct {

        
            Priority float64  `json:"priority"`
            IsDefault bool  `json:"is_default"`
            TemplateSlugs []string  `json:"template_slugs"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            IsActive bool  `json:"is_active"`
            Logo string  `json:"logo"`
            AppID string  `json:"app_id"`
            Attributes []AttributeDetailsGroup  `json:"attributes"`
         
    }
    
    // ConfigSuccessResponse used by Catalog
    type ConfigSuccessResponse struct {

        
            Message string  `json:"message"`
         
    }
    
    // AppConfigurationsSort used by Catalog
    type AppConfigurationsSort struct {

        
            Priority float64  `json:"priority"`
            IsDefault bool  `json:"is_default"`
            Name string  `json:"name"`
            IsActive bool  `json:"is_active"`
            Logo string  `json:"logo"`
            AppID string  `json:"app_id"`
            DefaultKey string  `json:"default_key"`
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
            Units []map[string]interface{}  `json:"units"`
            Key string  `json:"key"`
            FilterTypes []string  `json:"filter_types"`
         
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
            Variant map[string]interface{}  `json:"variant"`
            Detail map[string]interface{}  `json:"detail"`
            Compare map[string]interface{}  `json:"compare"`
         
    }
    
    // GetCatalogConfigurationMetaData used by Catalog
    type GetCatalogConfigurationMetaData struct {

        
            Listing MetaDataListingResponse  `json:"listing"`
            Product GetCatalogConfigurationDetailsProduct  `json:"product"`
         
    }
    
    // ConfigurationListingSortConfig used by Catalog
    type ConfigurationListingSortConfig struct {

        
            Priority float64  `json:"priority"`
            Name string  `json:"name"`
            IsActive bool  `json:"is_active"`
            Logo string  `json:"logo"`
            Key string  `json:"key"`
         
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

        
            Map map[string]interface{}  `json:"map"`
            Condition string  `json:"condition"`
            Value string  `json:"value"`
            BucketPoints []ConfigurationBucketPoints  `json:"bucket_points"`
            Sort string  `json:"sort"`
            MapValues []map[string]interface{}  `json:"map_values"`
         
    }
    
    // ConfigurationListingFilterConfig used by Catalog
    type ConfigurationListingFilterConfig struct {

        
            Type string  `json:"type"`
            Priority float64  `json:"priority"`
            Name string  `json:"name"`
            DisplayName string  `json:"display_name"`
            IsActive bool  `json:"is_active"`
            Logo string  `json:"logo"`
            ValueConfig ConfigurationListingFilterValue  `json:"value_config"`
            Key string  `json:"key"`
         
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

        
            Priority float64  `json:"priority"`
            IsActive bool  `json:"is_active"`
            Logo string  `json:"logo"`
            Title string  `json:"title"`
            Subtitle string  `json:"subtitle"`
            Key string  `json:"key"`
            Size ProductSize  `json:"size"`
         
    }
    
    // ConfigurationProductSimilar used by Catalog
    type ConfigurationProductSimilar struct {

        
            Config []ConfigurationProductConfig  `json:"config"`
         
    }
    
    // ConfigurationProductVariantConfig used by Catalog
    type ConfigurationProductVariantConfig struct {

        
            Priority float64  `json:"priority"`
            DisplayType string  `json:"display_type"`
            Name string  `json:"name"`
            IsActive bool  `json:"is_active"`
            Logo string  `json:"logo"`
            Key string  `json:"key"`
            Size ProductSize  `json:"size"`
         
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

        
            Type string  `json:"type"`
            Listing ConfigurationListing  `json:"listing"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            CreatedOn string  `json:"created_on"`
            AppID string  `json:"app_id"`
            ID string  `json:"id"`
            ConfigID string  `json:"config_id"`
            ConfigType string  `json:"config_type"`
            ModifiedOn string  `json:"modified_on"`
            Product ConfigurationProduct  `json:"product"`
         
    }
    
    // GetAppCatalogConfiguration used by Catalog
    type GetAppCatalogConfiguration struct {

        
            IsDefault bool  `json:"is_default"`
            Data AppCatalogConfiguration  `json:"data"`
         
    }
    
    // AppConfiguration used by Catalog
    type AppConfiguration struct {

        
            Type string  `json:"type"`
            Listing ConfigurationListing  `json:"listing"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            CreatedOn string  `json:"created_on"`
            AppID string  `json:"app_id"`
            ConfigID string  `json:"config_id"`
            ConfigType string  `json:"config_type"`
            ModifiedOn string  `json:"modified_on"`
            Product ConfigurationProduct  `json:"product"`
         
    }
    
    // GetCatalogConfigurationDetailsSchemaListing used by Catalog
    type GetCatalogConfigurationDetailsSchemaListing struct {

        
            Sort map[string]interface{}  `json:"sort"`
            Filter map[string]interface{}  `json:"filter"`
         
    }
    
    // EntityConfiguration used by Catalog
    type EntityConfiguration struct {

        
            Listing GetCatalogConfigurationDetailsSchemaListing  `json:"listing"`
            AppID string  `json:"app_id"`
            ID string  `json:"id"`
            ConfigID string  `json:"config_id"`
            ConfigType string  `json:"config_type"`
            Product GetCatalogConfigurationDetailsProduct  `json:"product"`
         
    }
    
    // GetAppCatalogEntityConfiguration used by Catalog
    type GetAppCatalogEntityConfiguration struct {

        
            IsDefault bool  `json:"is_default"`
            Data EntityConfiguration  `json:"data"`
         
    }
    
    // ProductSortOn used by Catalog
    type ProductSortOn struct {

        
            Display string  `json:"display"`
            Name string  `json:"name"`
            Value string  `json:"value"`
            Logo string  `json:"logo"`
            IsSelected bool  `json:"is_selected"`
         
    }
    
    // ProductFiltersValue used by Catalog
    type ProductFiltersValue struct {

        
            Display string  `json:"display"`
            CurrencySymbol string  `json:"currency_symbol"`
            SelectedMax float64  `json:"selected_max"`
            CurrencyCode string  `json:"currency_code"`
            Value interface{}  `json:"value"`
            QueryFormat string  `json:"query_format"`
            Max float64  `json:"max"`
            SelectedMin float64  `json:"selected_min"`
            IsSelected bool  `json:"is_selected"`
            DisplayFormat string  `json:"display_format"`
            Min float64  `json:"min"`
            Count float64  `json:"count"`
         
    }
    
    // ProductFiltersKey used by Catalog
    type ProductFiltersKey struct {

        
            Display string  `json:"display"`
            Name string  `json:"name"`
            Kind string  `json:"kind"`
            Logo string  `json:"logo"`
            Operators []string  `json:"operators"`
         
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
    
    // CollectionErrorResponse used by Catalog
    type CollectionErrorResponse struct {

        
            Errors map[string]interface{}  `json:"errors"`
            Code float64  `json:"code"`
            Message string  `json:"message"`
         
    }
    
    // CollectionListingFilterType used by Catalog
    type CollectionListingFilterType struct {

        
            Display string  `json:"display"`
            IsSelected bool  `json:"is_selected"`
            Name string  `json:"name"`
         
    }
    
    // CollectionListingFilterTag used by Catalog
    type CollectionListingFilterTag struct {

        
            Display string  `json:"display"`
            IsSelected bool  `json:"is_selected"`
            Name string  `json:"name"`
         
    }
    
    // CollectionListingFilter used by Catalog
    type CollectionListingFilter struct {

        
            Type []CollectionListingFilterType  `json:"type"`
            Tags []CollectionListingFilterTag  `json:"tags"`
         
    }
    
    // CollectionQuery used by Catalog
    type CollectionQuery struct {

        
            Attribute string  `json:"attribute"`
            Op string  `json:"op"`
            Value []interface{}  `json:"value"`
         
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
    
    // Media2 used by Catalog
    type Media2 struct {

        
            Type string  `json:"type"`
            URL string  `json:"url"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // GetCollectionDetailNest used by Catalog
    type GetCollectionDetailNest struct {

        
            Name string  `json:"name"`
            Badge map[string]interface{}  `json:"badge"`
            Query []CollectionQuery  `json:"query"`
            Banners ImageUrls  `json:"banners"`
            Cron map[string]interface{}  `json:"cron"`
            AllowFacets bool  `json:"allow_facets"`
            Meta map[string]interface{}  `json:"meta"`
            Description string  `json:"description"`
            Type string  `json:"type"`
            Priority float64  `json:"priority"`
            Action Action  `json:"action"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            SortOn string  `json:"sort_on"`
            Slug string  `json:"slug"`
            IsActive bool  `json:"is_active"`
            Logo Media2  `json:"logo"`
            AllowSort bool  `json:"allow_sort"`
            UID string  `json:"uid"`
            Schedule map[string]interface{}  `json:"_schedule"`
            Tag []string  `json:"tag"`
            AppID string  `json:"app_id"`
         
    }
    
    // GetCollectionListingResponse used by Catalog
    type GetCollectionListingResponse struct {

        
            Page Page  `json:"page"`
            Filters CollectionListingFilter  `json:"filters"`
            Items []GetCollectionDetailNest  `json:"items"`
         
    }
    
    // CollectionBadge used by Catalog
    type CollectionBadge struct {

        
            Color string  `json:"color"`
            Text string  `json:"text"`
         
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
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
            Email string  `json:"email"`
         
    }
    
    // SeoDetail used by Catalog
    type SeoDetail struct {

        
            Title string  `json:"title"`
            Description string  `json:"description"`
         
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
            Duration float64  `json:"duration"`
            NextSchedule []NextSchedule  `json:"next_schedule"`
            End string  `json:"end"`
         
    }
    
    // CreateCollection used by Catalog
    type CreateCollection struct {

        
            Tags []string  `json:"tags"`
            Name string  `json:"name"`
            Badge CollectionBadge  `json:"badge"`
            Query []CollectionQuery  `json:"query"`
            Banners CollectionBanner  `json:"banners"`
            AllowFacets bool  `json:"allow_facets"`
            Meta map[string]interface{}  `json:"meta"`
            Description string  `json:"description"`
            Type string  `json:"type"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            Priority float64  `json:"priority"`
            ModifiedBy UserInfo  `json:"modified_by"`
            CreatedBy UserInfo  `json:"created_by"`
            Seo SeoDetail  `json:"seo"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            IsVisible bool  `json:"is_visible"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            SortOn string  `json:"sort_on"`
            Slug string  `json:"slug"`
            IsActive bool  `json:"is_active"`
            Logo CollectionImage  `json:"logo"`
            AllowSort bool  `json:"allow_sort"`
            Schedule CollectionSchedule  `json:"_schedule"`
            Published bool  `json:"published"`
            AppID string  `json:"app_id"`
         
    }
    
    // CollectionCreateResponse used by Catalog
    type CollectionCreateResponse struct {

        
            Tags []string  `json:"tags"`
            Name string  `json:"name"`
            Badge map[string]interface{}  `json:"badge"`
            Query []CollectionQuery  `json:"query"`
            Banners ImageUrls  `json:"banners"`
            Cron map[string]interface{}  `json:"cron"`
            AllowFacets bool  `json:"allow_facets"`
            Meta map[string]interface{}  `json:"meta"`
            Description string  `json:"description"`
            Type string  `json:"type"`
            Priority float64  `json:"priority"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            SortOn string  `json:"sort_on"`
            Slug string  `json:"slug"`
            IsActive bool  `json:"is_active"`
            Logo Media2  `json:"logo"`
            AllowSort bool  `json:"allow_sort"`
            Schedule map[string]interface{}  `json:"_schedule"`
            AppID string  `json:"app_id"`
         
    }
    
    // CollectionDetailResponse used by Catalog
    type CollectionDetailResponse struct {

        
            Name string  `json:"name"`
            Badge map[string]interface{}  `json:"badge"`
            Query []CollectionQuery  `json:"query"`
            Banners ImageUrls  `json:"banners"`
            Cron map[string]interface{}  `json:"cron"`
            AllowFacets bool  `json:"allow_facets"`
            Meta map[string]interface{}  `json:"meta"`
            Description string  `json:"description"`
            Type string  `json:"type"`
            Priority float64  `json:"priority"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            SortOn string  `json:"sort_on"`
            Slug string  `json:"slug"`
            IsActive bool  `json:"is_active"`
            Logo Media2  `json:"logo"`
            AllowSort bool  `json:"allow_sort"`
            UID string  `json:"uid"`
            Schedule map[string]interface{}  `json:"_schedule"`
            Tag []string  `json:"tag"`
            AppID string  `json:"app_id"`
         
    }
    
    // UpdateCollection used by Catalog
    type UpdateCollection struct {

        
            Tags []string  `json:"tags"`
            Name string  `json:"name"`
            Badge CollectionBadge  `json:"badge"`
            Query []CollectionQuery  `json:"query"`
            Banners CollectionBanner  `json:"banners"`
            AllowFacets bool  `json:"allow_facets"`
            Meta map[string]interface{}  `json:"meta"`
            Description string  `json:"description"`
            Type string  `json:"type"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            Priority float64  `json:"priority"`
            ModifiedBy UserInfo  `json:"modified_by"`
            Seo SeoDetail  `json:"seo"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            IsVisible bool  `json:"is_visible"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            SortOn string  `json:"sort_on"`
            Slug string  `json:"slug"`
            IsActive bool  `json:"is_active"`
            Logo CollectionImage  `json:"logo"`
            AllowSort bool  `json:"allow_sort"`
            Schedule CollectionSchedule  `json:"_schedule"`
            Published bool  `json:"published"`
         
    }
    
    // CollectionUpdateImage used by Catalog
    type CollectionUpdateImage struct {

        
            Type string  `json:"type"`
            URL string  `json:"url"`
         
    }
    
    // CollectionUpdateBanner used by Catalog
    type CollectionUpdateBanner struct {

        
            Landscape CollectionUpdateImage  `json:"landscape"`
            Portrait CollectionUpdateImage  `json:"portrait"`
         
    }
    
    // CollectionUpdateResponse used by Catalog
    type CollectionUpdateResponse struct {

        
            Tags []string  `json:"tags"`
            Name string  `json:"name"`
            Badge CollectionBadge  `json:"badge"`
            Query []CollectionQuery  `json:"query"`
            Banners CollectionUpdateBanner  `json:"banners"`
            AllowFacets bool  `json:"allow_facets"`
            Meta map[string]interface{}  `json:"meta"`
            Description string  `json:"description"`
            Type string  `json:"type"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            Priority float64  `json:"priority"`
            ModifiedBy UserInfo  `json:"modified_by"`
            Seo SeoDetail  `json:"seo"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            IsVisible bool  `json:"is_visible"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            SortOn string  `json:"sort_on"`
            Slug string  `json:"slug"`
            IsActive bool  `json:"is_active"`
            Logo CollectionUpdateImage  `json:"logo"`
            AllowSort bool  `json:"allow_sort"`
            Schedule CollectionSchedule  `json:"_schedule"`
            Published bool  `json:"published"`
            AppID string  `json:"app_id"`
         
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

        
            Effective Price1  `json:"effective"`
            Marked Price1  `json:"marked"`
         
    }
    
    // ProductBrand used by Catalog
    type ProductBrand struct {

        
            Action Action  `json:"action"`
            Name string  `json:"name"`
            UID float64  `json:"uid"`
            Logo Media2  `json:"logo"`
         
    }
    
    // ProductDetailAttribute used by Catalog
    type ProductDetailAttribute struct {

        
            Type string  `json:"type"`
            Value string  `json:"value"`
            Key string  `json:"key"`
         
    }
    
    // ProductDetailGroupedAttribute used by Catalog
    type ProductDetailGroupedAttribute struct {

        
            Details []ProductDetailAttribute  `json:"details"`
            Title string  `json:"title"`
         
    }
    
    // ProductListingDetail used by Catalog
    type ProductListingDetail struct {

        
            Price ProductListingPrice  `json:"price"`
            Brand ProductBrand  `json:"brand"`
            Name string  `json:"name"`
            Rating float64  `json:"rating"`
            ProductOnlineDate string  `json:"product_online_date"`
            Color string  `json:"color"`
            ItemCode string  `json:"item_code"`
            PromoMeta map[string]interface{}  `json:"promo_meta"`
            Description string  `json:"description"`
            Type string  `json:"type"`
            RatingCount float64  `json:"rating_count"`
            Tryouts []string  `json:"tryouts"`
            ShortDescription string  `json:"short_description"`
            Sellable bool  `json:"sellable"`
            Highlights []string  `json:"highlights"`
            ItemType string  `json:"item_type"`
            Slug string  `json:"slug"`
            HasVariant bool  `json:"has_variant"`
            UID float64  `json:"uid"`
            Attributes map[string]interface{}  `json:"attributes"`
            Discount string  `json:"discount"`
            ImageNature string  `json:"image_nature"`
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            Medias []Media2  `json:"medias"`
            Similars []string  `json:"similars"`
         
    }
    
    // GetCollectionItemsResponse used by Catalog
    type GetCollectionItemsResponse struct {

        
            Page Page  `json:"page"`
            Filters []ProductFilters  `json:"filters"`
            SortOn []ProductSortOn  `json:"sort_on"`
            Items []ProductListingDetail  `json:"items"`
         
    }
    
    // CollectionItem used by Catalog
    type CollectionItem struct {

        
            ItemID float64  `json:"item_id"`
            Priority float64  `json:"priority"`
            Action string  `json:"action"`
         
    }
    
    // CollectionItemUpdate used by Catalog
    type CollectionItemUpdate struct {

        
            Type string  `json:"type"`
            Items []CollectionItem  `json:"items"`
            Query []CollectionQuery  `json:"query"`
            AllowSort bool  `json:"allow_sort"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            AllowFacets bool  `json:"allow_facets"`
         
    }
    
    // UpdatedResponse used by Catalog
    type UpdatedResponse struct {

        
            ItemsNotUpdated []float64  `json:"items_not_updated"`
            Message string  `json:"message"`
         
    }
    
    // CatalogInsightBrand used by Catalog
    type CatalogInsightBrand struct {

        
            AvailableSizes float64  `json:"available_sizes"`
            TotalArticles float64  `json:"total_articles"`
            Name string  `json:"name"`
            TotalSizes float64  `json:"total_sizes"`
            AvailableArticles float64  `json:"available_articles"`
            ArticleFreshness float64  `json:"article_freshness"`
         
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
            BrandIds []float64  `json:"brand_ids"`
            Enabled bool  `json:"enabled"`
            StoreIds []float64  `json:"store_ids"`
            Platform string  `json:"platform"`
            OptLevel string  `json:"opt_level"`
         
    }
    
    // CompanyOptIn used by Catalog
    type CompanyOptIn struct {

        
            CompanyID float64  `json:"company_id"`
            StoreIds []float64  `json:"store_ids"`
            BrandIds []float64  `json:"brand_ids"`
            Enabled bool  `json:"enabled"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            CreatedOn float64  `json:"created_on"`
            Platform string  `json:"platform"`
            OptLevel string  `json:"opt_level"`
            ModifiedOn float64  `json:"modified_on"`
         
    }
    
    // GetOptInPlatform used by Catalog
    type GetOptInPlatform struct {

        
            Page Page  `json:"page"`
            Items []CompanyOptIn  `json:"items"`
         
    }
    
    // OptinCompanyDetail used by Catalog
    type OptinCompanyDetail struct {

        
            CompanyType string  `json:"company_type"`
            Name string  `json:"name"`
            UID float64  `json:"uid"`
            BusinessType string  `json:"business_type"`
         
    }
    
    // CompanyBrandDetail used by Catalog
    type CompanyBrandDetail struct {

        
            TotalArticle float64  `json:"total_article"`
            CompanyID float64  `json:"company_id"`
            BrandName string  `json:"brand_name"`
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
            Brand float64  `json:"brand"`
            Store float64  `json:"store"`
         
    }
    
    // StoreDetail used by Catalog
    type StoreDetail struct {

        
            StoreType string  `json:"store_type"`
            Manager map[string]interface{}  `json:"manager"`
            CompanyID float64  `json:"company_id"`
            Documents []map[string]interface{}  `json:"documents"`
            Name string  `json:"name"`
            DisplayName string  `json:"display_name"`
            Address map[string]interface{}  `json:"address"`
            StoreCode string  `json:"store_code"`
            CreatedOn string  `json:"created_on"`
            UID float64  `json:"uid"`
            AdditionalContacts []map[string]interface{}  `json:"additional_contacts"`
            ModifiedOn string  `json:"modified_on"`
            Timing map[string]interface{}  `json:"timing"`
         
    }
    
    // OptinStoreDetails used by Catalog
    type OptinStoreDetails struct {

        
            Page Page  `json:"page"`
            Items []StoreDetail  `json:"items"`
         
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

        
            Type string  `json:"type"`
            Format string  `json:"format"`
            Multi bool  `json:"multi"`
            Mandatory bool  `json:"mandatory"`
            AllowedValues []string  `json:"allowed_values"`
            Range AttributeSchemaRange  `json:"range"`
         
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
    
    // GenderDetail used by Catalog
    type GenderDetail struct {

        
            Details AttributeMasterDetails  `json:"details"`
            IsNested bool  `json:"is_nested"`
            Departments []string  `json:"departments"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Schema AttributeMaster  `json:"schema"`
            Logo string  `json:"logo"`
            Filters AttributeMasterFilter  `json:"filters"`
            ID string  `json:"id"`
            EnabledForEndConsumer bool  `json:"enabled_for_end_consumer"`
            Meta AttributeMasterMeta  `json:"meta"`
            Description string  `json:"description"`
         
    }
    
    // CategoriesResponse used by Catalog
    type CategoriesResponse struct {

        
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
            TemplateSlug string  `json:"template_slug"`
            SlugKey string  `json:"slug_key"`
         
    }
    
    // ProdcutTemplateCategoriesResponse used by Catalog
    type ProdcutTemplateCategoriesResponse struct {

        
            Page Page  `json:"page"`
            Items []CategoriesResponse  `json:"items"`
         
    }
    
    // PTErrorResponse used by Catalog
    type PTErrorResponse struct {

        
            Errors map[string]interface{}  `json:"errors"`
            Code string  `json:"code"`
            Status float64  `json:"status"`
            Meta map[string]interface{}  `json:"meta"`
            Message string  `json:"message"`
         
    }
    
    // UserSerializer used by Catalog
    type UserSerializer struct {

        
            Contact string  `json:"contact"`
            Username string  `json:"username"`
            ID string  `json:"_id"`
            UserID string  `json:"user_id"`
            UID string  `json:"uid"`
         
    }
    
    // GetDepartment used by Catalog
    type GetDepartment struct {

        
            PageSize float64  `json:"page_size"`
            Name string  `json:"name"`
            PriorityOrder float64  `json:"priority_order"`
            ItemType string  `json:"item_type"`
            Synonyms []string  `json:"synonyms"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            Slug string  `json:"slug"`
            IsActive bool  `json:"is_active"`
            Logo string  `json:"logo"`
            Search string  `json:"search"`
            CreatedOn string  `json:"created_on"`
            CreatedBy UserSerializer  `json:"created_by"`
            UID float64  `json:"uid"`
            ModifiedOn string  `json:"modified_on"`
            PageNo float64  `json:"page_no"`
         
    }
    
    // DepartmentsResponse used by Catalog
    type DepartmentsResponse struct {

        
            Page Page  `json:"page"`
            Items []GetDepartment  `json:"items"`
         
    }
    
    // DepartmentErrorResponse used by Catalog
    type DepartmentErrorResponse struct {

        
            Errors map[string]interface{}  `json:"errors"`
            Code string  `json:"code"`
            Status float64  `json:"status"`
            Meta map[string]interface{}  `json:"meta"`
            Message string  `json:"message"`
         
    }
    
    // DepartmentCreateUpdate used by Catalog
    type DepartmentCreateUpdate struct {

        
            Tags []string  `json:"tags"`
            Name string  `json:"name"`
            PriorityOrder float64  `json:"priority_order"`
            Synonyms []string  `json:"synonyms"`
            Cls string  `json:"_cls"`
            Slug string  `json:"slug"`
            IsActive bool  `json:"is_active"`
            Logo string  `json:"logo"`
            Platforms map[string]interface{}  `json:"platforms"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            UID float64  `json:"uid"`
         
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
    
    // DepartmentModel used by Catalog
    type DepartmentModel struct {

        
            Name string  `json:"name"`
            PriorityOrder float64  `json:"priority_order"`
            ID interface{}  `json:"_id"`
            Synonyms []string  `json:"synonyms"`
            Cls string  `json:"_cls"`
            ModifiedBy UserDetail  `json:"modified_by"`
            CreatedBy UserDetail  `json:"created_by"`
            VerifiedBy UserDetail  `json:"verified_by"`
            CreatedOn string  `json:"created_on"`
            Slug string  `json:"slug"`
            Logo string  `json:"logo"`
            IsActive bool  `json:"is_active"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            UID float64  `json:"uid"`
            VerifiedOn string  `json:"verified_on"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // ProductTemplate used by Catalog
    type ProductTemplate struct {

        
            Departments []string  `json:"departments"`
            Name string  `json:"name"`
            IsExpirable bool  `json:"is_expirable"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Slug string  `json:"slug"`
            IsActive bool  `json:"is_active"`
            IsArchived bool  `json:"is_archived"`
            Categories []string  `json:"categories"`
            Logo string  `json:"logo"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            CreatedOn string  `json:"created_on"`
            IsPhysical bool  `json:"is_physical"`
            Attributes []string  `json:"attributes"`
            Tag string  `json:"tag"`
            ModifiedOn string  `json:"modified_on"`
            Description string  `json:"description"`
         
    }
    
    // TemplatesResponse used by Catalog
    type TemplatesResponse struct {

        
            Page Page  `json:"page"`
            Items ProductTemplate  `json:"items"`
         
    }
    
    // TemplateDetails used by Catalog
    type TemplateDetails struct {

        
            Departments []string  `json:"departments"`
            Name string  `json:"name"`
            IsExpirable bool  `json:"is_expirable"`
            Slug string  `json:"slug"`
            IsActive bool  `json:"is_active"`
            IsArchived bool  `json:"is_archived"`
            Categories []string  `json:"categories"`
            Logo string  `json:"logo"`
            IsPhysical bool  `json:"is_physical"`
            ID string  `json:"id"`
            Attributes []string  `json:"attributes"`
            Tag string  `json:"tag"`
            Description string  `json:"description"`
         
    }
    
    // Properties used by Catalog
    type Properties struct {

        
            Tags map[string]interface{}  `json:"tags"`
            NoOfBoxes map[string]interface{}  `json:"no_of_boxes"`
            Name map[string]interface{}  `json:"name"`
            Variants map[string]interface{}  `json:"variants"`
            BrandUID map[string]interface{}  `json:"brand_uid"`
            ItemCode map[string]interface{}  `json:"item_code"`
            IsDependent map[string]interface{}  `json:"is_dependent"`
            Command map[string]interface{}  `json:"command"`
            Description map[string]interface{}  `json:"description"`
            ProductPublish map[string]interface{}  `json:"product_publish"`
            Currency map[string]interface{}  `json:"currency"`
            ShortDescription map[string]interface{}  `json:"short_description"`
            TraderType map[string]interface{}  `json:"trader_type"`
            Highlights map[string]interface{}  `json:"highlights"`
            Trader map[string]interface{}  `json:"trader"`
            MultiSize map[string]interface{}  `json:"multi_size"`
            Media map[string]interface{}  `json:"media"`
            ItemType map[string]interface{}  `json:"item_type"`
            Slug map[string]interface{}  `json:"slug"`
            IsActive map[string]interface{}  `json:"is_active"`
            HsnCode map[string]interface{}  `json:"hsn_code"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            SizeGuide map[string]interface{}  `json:"size_guide"`
            ReturnConfig map[string]interface{}  `json:"return_config"`
            CategorySlug map[string]interface{}  `json:"category_slug"`
            Sizes map[string]interface{}  `json:"sizes"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            ProductGroupTag map[string]interface{}  `json:"product_group_tag"`
            CountryOfOrigin map[string]interface{}  `json:"country_of_origin"`
         
    }
    
    // GlobalValidation used by Catalog
    type GlobalValidation struct {

        
            Type string  `json:"type"`
            Properties Properties  `json:"properties"`
            Required []string  `json:"required"`
            Title string  `json:"title"`
            Definitions map[string]interface{}  `json:"definitions"`
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

        
            UID string  `json:"uid"`
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
            Email string  `json:"email"`
         
    }
    
    // ProductTemplateExportResponse used by Catalog
    type ProductTemplateExportResponse struct {

        
            Type string  `json:"type"`
            CompletedOn string  `json:"completed_on"`
            TaskID string  `json:"task_id"`
            CreatedBy UserInfo1  `json:"created_by"`
            NotificationEmails []string  `json:"notification_emails"`
            Status string  `json:"status"`
            URL string  `json:"url"`
            Filters map[string]interface{}  `json:"filters"`
            SellerID float64  `json:"seller_id"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // ProductDownloadsResponse used by Catalog
    type ProductDownloadsResponse struct {

        
            Items []ProductTemplateExportResponse  `json:"items"`
         
    }
    
    // ProductTemplateExportFilterRequest used by Catalog
    type ProductTemplateExportFilterRequest struct {

        
            Templates []string  `json:"templates"`
            Brands []string  `json:"brands"`
            CatalogueTypes []string  `json:"catalogue_types"`
            FromDate string  `json:"from_date"`
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
    
    // Hierarchy used by Catalog
    type Hierarchy struct {

        
            L2 float64  `json:"l2"`
            Department float64  `json:"department"`
            L1 float64  `json:"l1"`
         
    }
    
    // Media3 used by Catalog
    type Media3 struct {

        
            Portrait string  `json:"portrait"`
            Landscape string  `json:"landscape"`
            Logo string  `json:"logo"`
         
    }
    
    // CategoryMappingValues used by Catalog
    type CategoryMappingValues struct {

        
            CatalogID float64  `json:"catalog_id"`
            Name string  `json:"name"`
         
    }
    
    // CategoryMapping used by Catalog
    type CategoryMapping struct {

        
            Google CategoryMappingValues  `json:"google"`
            Facebook CategoryMappingValues  `json:"facebook"`
            Ajio CategoryMappingValues  `json:"ajio"`
         
    }
    
    // Category used by Catalog
    type Category struct {

        
            Hierarchy []Hierarchy  `json:"hierarchy"`
            Priority float64  `json:"priority"`
            Departments []float64  `json:"departments"`
            Name string  `json:"name"`
            Media Media3  `json:"media"`
            Synonyms []string  `json:"synonyms"`
            Tryouts []string  `json:"tryouts"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Slug string  `json:"slug"`
            IsActive bool  `json:"is_active"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            CreatedOn string  `json:"created_on"`
            Marketplaces CategoryMapping  `json:"marketplaces"`
            ID string  `json:"id"`
            UID float64  `json:"uid"`
            Level float64  `json:"level"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // CategoryResponse used by Catalog
    type CategoryResponse struct {

        
            Page Page  `json:"page"`
            Items []Category  `json:"items"`
         
    }
    
    // CategoryRequestBody used by Catalog
    type CategoryRequestBody struct {

        
            Hierarchy []Hierarchy  `json:"hierarchy"`
            Priority float64  `json:"priority"`
            Departments []float64  `json:"departments"`
            Media Media3  `json:"media"`
            Name string  `json:"name"`
            Synonyms []string  `json:"synonyms"`
            Tryouts []string  `json:"tryouts"`
            Slug string  `json:"slug"`
            IsActive bool  `json:"is_active"`
            Marketplaces CategoryMapping  `json:"marketplaces"`
            Level float64  `json:"level"`
         
    }
    
    // CategoryCreateResponse used by Catalog
    type CategoryCreateResponse struct {

        
            UID float64  `json:"uid"`
            Message string  `json:"message"`
         
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
    
    // Logo used by Catalog
    type Logo struct {

        
            SecureURL string  `json:"secure_url"`
            URL string  `json:"url"`
            AspectRatioF float64  `json:"aspect_ratio_f"`
            AspectRatio string  `json:"aspect_ratio"`
         
    }
    
    // Brand used by Catalog
    type Brand struct {

        
            Name string  `json:"name"`
            UID float64  `json:"uid"`
            Logo Logo  `json:"logo"`
         
    }
    
    // ProductPublish used by Catalog
    type ProductPublish struct {

        
            IsSet bool  `json:"is_set"`
            ProductOnlineDate string  `json:"product_online_date"`
         
    }
    
    // VerifiedBy used by Catalog
    type VerifiedBy struct {

        
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
         
    }
    
    // Image used by Catalog
    type Image struct {

        
            SecureURL string  `json:"secure_url"`
            URL string  `json:"url"`
            AspectRatioF float64  `json:"aspect_ratio_f"`
            AspectRatio string  `json:"aspect_ratio"`
         
    }
    
    // Trader used by Catalog
    type Trader struct {

        
            Type string  `json:"type"`
            Address []string  `json:"address"`
            Name string  `json:"name"`
         
    }
    
    // ReturnConfigResponse used by Catalog
    type ReturnConfigResponse struct {

        
            Returnable bool  `json:"returnable"`
            Time float64  `json:"time"`
            Unit string  `json:"unit"`
         
    }
    
    // NetQuantityResponse used by Catalog
    type NetQuantityResponse struct {

        
            Value float64  `json:"value"`
            Unit string  `json:"unit"`
         
    }
    
    // ProductSchemaV2 used by Catalog
    type ProductSchemaV2 struct {

        
            Brand Brand  `json:"brand"`
            Tags []string  `json:"tags"`
            NoOfBoxes float64  `json:"no_of_boxes"`
            Name string  `json:"name"`
            CompanyID float64  `json:"company_id"`
            Stage string  `json:"stage"`
            AllIdentifiers []string  `json:"all_identifiers"`
            AllSizes []map[string]interface{}  `json:"all_sizes"`
            Variants map[string]interface{}  `json:"variants"`
            Moq map[string]interface{}  `json:"moq"`
            Color string  `json:"color"`
            BrandUID float64  `json:"brand_uid"`
            IsPhysical bool  `json:"is_physical"`
            ID string  `json:"id"`
            ItemCode string  `json:"item_code"`
            IsDependent bool  `json:"is_dependent"`
            Description string  `json:"description"`
            IsExpirable bool  `json:"is_expirable"`
            ProductPublish ProductPublish  `json:"product_publish"`
            Currency string  `json:"currency"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            VerifiedBy VerifiedBy  `json:"verified_by"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ShortDescription string  `json:"short_description"`
            Images []Image  `json:"images"`
            IsSet bool  `json:"is_set"`
            Highlights []string  `json:"highlights"`
            Trader []Trader  `json:"trader"`
            Category map[string]interface{}  `json:"category"`
            TemplateTag string  `json:"template_tag"`
            Media []Media2  `json:"media"`
            ItemType string  `json:"item_type"`
            CategoryUID float64  `json:"category_uid"`
            MultiSize bool  `json:"multi_size"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            VariantGroup map[string]interface{}  `json:"variant_group"`
            Slug string  `json:"slug"`
            IsActive bool  `json:"is_active"`
            CreatedOn string  `json:"created_on"`
            UID float64  `json:"uid"`
            HsnCode string  `json:"hsn_code"`
            Attributes map[string]interface{}  `json:"attributes"`
            VariantMedia map[string]interface{}  `json:"variant_media"`
            L3Mapping []string  `json:"l3_mapping"`
            Pending string  `json:"pending"`
            Departments []float64  `json:"departments"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            IsImageLessProduct bool  `json:"is_image_less_product"`
            SizeGuide string  `json:"size_guide"`
            CategorySlug string  `json:"category_slug"`
            ImageNature string  `json:"image_nature"`
            Sizes []map[string]interface{}  `json:"sizes"`
            ReturnConfig ReturnConfigResponse  `json:"return_config"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            NetQuantity NetQuantityResponse  `json:"net_quantity"`
            ProductGroupTag []string  `json:"product_group_tag"`
            CountryOfOrigin string  `json:"country_of_origin"`
            VerifiedOn string  `json:"verified_on"`
            PrimaryColor string  `json:"primary_color"`
            ModifiedOn string  `json:"modified_on"`
            AllCompanyIds []float64  `json:"all_company_ids"`
         
    }
    
    // ProductListingResponseV2 used by Catalog
    type ProductListingResponseV2 struct {

        
            Page Page  `json:"page"`
            Items []ProductSchemaV2  `json:"items"`
         
    }
    
    // ProductPublish1 used by Catalog
    type ProductPublish1 struct {

        
            IsSet bool  `json:"is_set"`
            ProductOnlineDate string  `json:"product_online_date"`
         
    }
    
    // TaxIdentifier used by Catalog
    type TaxIdentifier struct {

        
            ReportingHsn string  `json:"reporting_hsn"`
            HsnCodeID string  `json:"hsn_code_id"`
            HsnCode string  `json:"hsn_code"`
         
    }
    
    // CustomOrder used by Catalog
    type CustomOrder struct {

        
            ManufacturingTime float64  `json:"manufacturing_time"`
            IsCustomOrder bool  `json:"is_custom_order"`
            ManufacturingTimeUnit string  `json:"manufacturing_time_unit"`
         
    }
    
    // ReturnConfig used by Catalog
    type ReturnConfig struct {

        
            Returnable bool  `json:"returnable"`
            Time float64  `json:"time"`
            Unit string  `json:"unit"`
         
    }
    
    // TeaserTag used by Catalog
    type TeaserTag struct {

        
            URL string  `json:"url"`
            Tag string  `json:"tag"`
         
    }
    
    // NetQuantity used by Catalog
    type NetQuantity struct {

        
            Value float64  `json:"value"`
            Unit string  `json:"unit"`
         
    }
    
    // ProductCreateUpdateSchemaV2 used by Catalog
    type ProductCreateUpdateSchemaV2 struct {

        
            CompanyID float64  `json:"company_id"`
            NoOfBoxes float64  `json:"no_of_boxes"`
            Tags []string  `json:"tags"`
            Name string  `json:"name"`
            Variants map[string]interface{}  `json:"variants"`
            BrandUID float64  `json:"brand_uid"`
            ChangeRequestID interface{}  `json:"change_request_id"`
            ItemCode string  `json:"item_code"`
            IsDependent bool  `json:"is_dependent"`
            Description string  `json:"description"`
            Action string  `json:"action"`
            ProductPublish ProductPublish1  `json:"product_publish"`
            Currency string  `json:"currency"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ShortDescription string  `json:"short_description"`
            IsSet bool  `json:"is_set"`
            Highlights []string  `json:"highlights"`
            Trader []Trader  `json:"trader"`
            TemplateTag string  `json:"template_tag"`
            TaxIdentifier TaxIdentifier  `json:"tax_identifier"`
            VariantGroup map[string]interface{}  `json:"variant_group"`
            Media []Media2  `json:"media"`
            Requester string  `json:"requester"`
            MultiSize bool  `json:"multi_size"`
            ItemType string  `json:"item_type"`
            Slug string  `json:"slug"`
            IsActive bool  `json:"is_active"`
            BulkJobID string  `json:"bulk_job_id"`
            UID float64  `json:"uid"`
            Attributes map[string]interface{}  `json:"attributes"`
            VariantMedia map[string]interface{}  `json:"variant_media"`
            Departments []float64  `json:"departments"`
            CustomOrder CustomOrder  `json:"custom_order"`
            SizeGuide string  `json:"size_guide"`
            CategorySlug string  `json:"category_slug"`
            ReturnConfig ReturnConfig  `json:"return_config"`
            Sizes []map[string]interface{}  `json:"sizes"`
            TeaserTag TeaserTag  `json:"teaser_tag"`
            ProductGroupTag []string  `json:"product_group_tag"`
            NetQuantity NetQuantity  `json:"net_quantity"`
            CountryOfOrigin string  `json:"country_of_origin"`
            IsImageLessProduct bool  `json:"is_image_less_product"`
         
    }
    
    // ProductVariants used by Catalog
    type ProductVariants struct {

        
            Media []Media2  `json:"media"`
            Name string  `json:"name"`
            CategoryUID float64  `json:"category_uid"`
            BrandUID float64  `json:"brand_uid"`
            UID float64  `json:"uid"`
            ItemCode string  `json:"item_code"`
         
    }
    
    // ProductVariantsResponse used by Catalog
    type ProductVariantsResponse struct {

        
            Page Page  `json:"page"`
            Variants []ProductVariants  `json:"variants"`
         
    }
    
    // AttributeMasterSerializer used by Catalog
    type AttributeMasterSerializer struct {

        
            IsNested bool  `json:"is_nested"`
            Tags []string  `json:"tags"`
            Name string  `json:"name"`
            Suggestion string  `json:"suggestion"`
            Description string  `json:"description"`
            Details AttributeMasterDetails  `json:"details"`
            RawKey string  `json:"raw_key"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            EnabledForEndConsumer bool  `json:"enabled_for_end_consumer"`
            Variant bool  `json:"variant"`
            Synonyms map[string]interface{}  `json:"synonyms"`
            Slug string  `json:"slug"`
            Unit string  `json:"unit"`
            Logo string  `json:"logo"`
            CreatedOn string  `json:"created_on"`
            Departments []string  `json:"departments"`
            Schema AttributeMaster  `json:"schema"`
            Filters AttributeMasterFilter  `json:"filters"`
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
            Identifiers []ValidateIdentifier  `json:"identifiers"`
            ItemWidth float64  `json:"item_width"`
            ItemHeight float64  `json:"item_height"`
            ItemDimensionsUnitOfMeasure string  `json:"item_dimensions_unit_of_measure"`
            ItemLength float64  `json:"item_length"`
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

        
            IsSet bool  `json:"is_set"`
            ProductOnlineDate float64  `json:"product_online_date"`
         
    }
    
    // Product used by Catalog
    type Product struct {

        
            Brand Brand  `json:"brand"`
            Tags []string  `json:"tags"`
            NoOfBoxes float64  `json:"no_of_boxes"`
            Name string  `json:"name"`
            CompanyID float64  `json:"company_id"`
            Stage string  `json:"stage"`
            AllIdentifiers []string  `json:"all_identifiers"`
            AllSizes []map[string]interface{}  `json:"all_sizes"`
            Variants map[string]interface{}  `json:"variants"`
            Moq map[string]interface{}  `json:"moq"`
            Color string  `json:"color"`
            BrandUID float64  `json:"brand_uid"`
            IsPhysical bool  `json:"is_physical"`
            ID string  `json:"id"`
            ItemCode string  `json:"item_code"`
            IsDependent bool  `json:"is_dependent"`
            Description string  `json:"description"`
            IsExpirable bool  `json:"is_expirable"`
            ProductPublish ProductPublished  `json:"product_publish"`
            Currency string  `json:"currency"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            VerifiedBy VerifiedBy  `json:"verified_by"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ShortDescription string  `json:"short_description"`
            Images []Image  `json:"images"`
            IsSet bool  `json:"is_set"`
            Highlights []string  `json:"highlights"`
            Trader []Trader  `json:"trader"`
            Category map[string]interface{}  `json:"category"`
            TemplateTag string  `json:"template_tag"`
            Media []Media2  `json:"media"`
            ItemType string  `json:"item_type"`
            CategoryUID float64  `json:"category_uid"`
            MultiSize bool  `json:"multi_size"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            VariantGroup map[string]interface{}  `json:"variant_group"`
            Slug string  `json:"slug"`
            IsActive bool  `json:"is_active"`
            CreatedOn string  `json:"created_on"`
            UID float64  `json:"uid"`
            HsnCode string  `json:"hsn_code"`
            Attributes map[string]interface{}  `json:"attributes"`
            VariantMedia map[string]interface{}  `json:"variant_media"`
            L3Mapping []string  `json:"l3_mapping"`
            Pending string  `json:"pending"`
            Departments []float64  `json:"departments"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            IsImageLessProduct bool  `json:"is_image_less_product"`
            SizeGuide string  `json:"size_guide"`
            CategorySlug string  `json:"category_slug"`
            ImageNature string  `json:"image_nature"`
            Sizes []map[string]interface{}  `json:"sizes"`
            ReturnConfig ReturnConfigResponse  `json:"return_config"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            NetQuantity NetQuantityResponse  `json:"net_quantity"`
            ProductGroupTag []string  `json:"product_group_tag"`
            CountryOfOrigin string  `json:"country_of_origin"`
            VerifiedOn string  `json:"verified_on"`
            PrimaryColor string  `json:"primary_color"`
            ModifiedOn string  `json:"modified_on"`
            AllCompanyIds []float64  `json:"all_company_ids"`
         
    }
    
    // ProductListingResponse used by Catalog
    type ProductListingResponse struct {

        
            Page Page  `json:"page"`
            Items []Product  `json:"items"`
         
    }
    
    // UserDetail1 used by Catalog
    type UserDetail1 struct {

        
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
            FullName string  `json:"full_name"`
         
    }
    
    // ProductBulkRequest used by Catalog
    type ProductBulkRequest struct {

        
            CompanyID float64  `json:"company_id"`
            Stage string  `json:"stage"`
            Failed float64  `json:"failed"`
            TemplateTag string  `json:"template_tag"`
            Succeed float64  `json:"succeed"`
            Template ProductTemplate  `json:"template"`
            ModifiedBy UserDetail1  `json:"modified_by"`
            CreatedBy UserDetail1  `json:"created_by"`
            IsActive bool  `json:"is_active"`
            CreatedOn string  `json:"created_on"`
            Total float64  `json:"total"`
            FilePath string  `json:"file_path"`
            Cancelled float64  `json:"cancelled"`
            FailedRecords []string  `json:"failed_records"`
            ModifiedOn string  `json:"modified_on"`
            CancelledRecords []string  `json:"cancelled_records"`
         
    }
    
    // ProductBulkRequestList used by Catalog
    type ProductBulkRequestList struct {

        
            Page Page  `json:"page"`
            Items ProductBulkRequest  `json:"items"`
         
    }
    
    // BulkJob used by Catalog
    type BulkJob struct {

        
            CompanyID float64  `json:"company_id"`
            FailedRecords []map[string]interface{}  `json:"failed_records"`
            Failed float64  `json:"failed"`
            Stage string  `json:"stage"`
            TemplateTag string  `json:"template_tag"`
            TrackingURL string  `json:"tracking_url"`
            Succeed float64  `json:"succeed"`
            ModifiedBy UserInfo1  `json:"modified_by"`
            CreatedBy UserInfo1  `json:"created_by"`
            IsActive bool  `json:"is_active"`
            CreatedOn string  `json:"created_on"`
            Total float64  `json:"total"`
            FilePath string  `json:"file_path"`
            Cancelled float64  `json:"cancelled"`
            CustomTemplateTag string  `json:"custom_template_tag"`
            ModifiedOn string  `json:"modified_on"`
            CancelledRecords []map[string]interface{}  `json:"cancelled_records"`
         
    }
    
    // BulkResponse used by Catalog
    type BulkResponse struct {

        
            ModifiedBy UserInfo1  `json:"modified_by"`
            CreatedBy UserInfo1  `json:"created_by"`
            IsActive bool  `json:"is_active"`
            CreatedOn string  `json:"created_on"`
            BatchID string  `json:"batch_id"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // BulkProductRequest used by Catalog
    type BulkProductRequest struct {

        
            BatchID string  `json:"batch_id"`
            CompanyID float64  `json:"company_id"`
            TemplateTag string  `json:"template_tag"`
            Data []map[string]interface{}  `json:"data"`
         
    }
    
    // ProductTagsViewResponse used by Catalog
    type ProductTagsViewResponse struct {

        
            Items []string  `json:"items"`
         
    }
    
    // UserCommon used by Catalog
    type UserCommon struct {

        
            UserID string  `json:"user_id"`
            CompanyID float64  `json:"company_id"`
            Username string  `json:"username"`
         
    }
    
    // Items used by Catalog
    type Items struct {

        
            CompanyID float64  `json:"company_id"`
            TrackingURL string  `json:"tracking_url"`
            Failed float64  `json:"failed"`
            Stage string  `json:"stage"`
            Succeed float64  `json:"succeed"`
            ModifiedBy UserCommon  `json:"modified_by"`
            CreatedBy UserCommon  `json:"created_by"`
            IsActive bool  `json:"is_active"`
            CreatedOn string  `json:"created_on"`
            Retry float64  `json:"retry"`
            Total float64  `json:"total"`
            ID string  `json:"id"`
            FilePath string  `json:"file_path"`
            Cancelled float64  `json:"cancelled"`
            FailedRecords []string  `json:"failed_records"`
            ModifiedOn string  `json:"modified_on"`
            CancelledRecords []string  `json:"cancelled_records"`
         
    }
    
    // BulkAssetResponse used by Catalog
    type BulkAssetResponse struct {

        
            Page Page  `json:"page"`
            Items []Items  `json:"items"`
         
    }
    
    // ProductBulkAssets used by Catalog
    type ProductBulkAssets struct {

        
            URL string  `json:"url"`
            CompanyID float64  `json:"company_id"`
            User map[string]interface{}  `json:"user"`
         
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

        
            Price float64  `json:"price"`
            Store map[string]interface{}  `json:"store"`
            Quantity float64  `json:"quantity"`
            SellerIdentifier string  `json:"seller_identifier"`
            PriceTransfer float64  `json:"price_transfer"`
            Identifiers map[string]interface{}  `json:"identifiers"`
            InventoryUpdatedOn string  `json:"inventory_updated_on"`
            Currency string  `json:"currency"`
            PriceEffective float64  `json:"price_effective"`
            SellableQuantity float64  `json:"sellable_quantity"`
            UID string  `json:"uid"`
            ItemID float64  `json:"item_id"`
            Size string  `json:"size"`
         
    }
    
    // InventoryResponsePaginated used by Catalog
    type InventoryResponsePaginated struct {

        
            Page Page  `json:"page"`
            Items []InventoryResponse  `json:"items"`
         
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

        
            Price float64  `json:"price"`
            ItemWeightUnitOfMeasure string  `json:"item_weight_unit_of_measure"`
            Quantity float64  `json:"quantity"`
            ExpirationDate string  `json:"expiration_date"`
            PriceTransfer float64  `json:"price_transfer"`
            Set InventorySet  `json:"set"`
            Identifiers []GTIN  `json:"identifiers"`
            PriceEffective float64  `json:"price_effective"`
            StoreCode string  `json:"store_code"`
            Currency string  `json:"currency"`
            ItemHeight float64  `json:"item_height"`
            ItemWidth float64  `json:"item_width"`
            ItemDimensionsUnitOfMeasure string  `json:"item_dimensions_unit_of_measure"`
            ItemLength float64  `json:"item_length"`
            IsSet bool  `json:"is_set"`
            ItemWeight float64  `json:"item_weight"`
            Size string  `json:"size"`
         
    }
    
    // ItemQuery used by Catalog
    type ItemQuery struct {

        
            BrandUID float64  `json:"brand_uid"`
            UID float64  `json:"uid"`
            ItemCode string  `json:"item_code"`
         
    }
    
    // InventoryRequest used by Catalog
    type InventoryRequest struct {

        
            Sizes []InvSize  `json:"sizes"`
            CompanyID float64  `json:"company_id"`
            Item ItemQuery  `json:"item"`
         
    }
    
    // PriceMeta used by Catalog
    type PriceMeta struct {

        
            TpNotes map[string]interface{}  `json:"tp_notes"`
            Currency string  `json:"currency"`
            Effective float64  `json:"effective"`
            Transfer float64  `json:"transfer"`
            Marked float64  `json:"marked"`
            UpdatedAt string  `json:"updated_at"`
         
    }
    
    // BrandMeta used by Catalog
    type BrandMeta struct {

        
            ID float64  `json:"id"`
            Name string  `json:"name"`
         
    }
    
    // ManufacturerResponse used by Catalog
    type ManufacturerResponse struct {

        
            IsDefault bool  `json:"is_default"`
            Address string  `json:"address"`
            Name string  `json:"name"`
         
    }
    
    // QuantityBase used by Catalog
    type QuantityBase struct {

        
            UpdatedAt string  `json:"updated_at"`
            Count float64  `json:"count"`
         
    }
    
    // Quantities used by Catalog
    type Quantities struct {

        
            Sellable QuantityBase  `json:"sellable"`
            Damaged QuantityBase  `json:"damaged"`
            NotAvailable QuantityBase  `json:"not_available"`
            OrderCommitted QuantityBase  `json:"order_committed"`
         
    }
    
    // StoreMeta used by Catalog
    type StoreMeta struct {

        
            ID float64  `json:"id"`
         
    }
    
    // CompanyMeta used by Catalog
    type CompanyMeta struct {

        
            ID float64  `json:"id"`
         
    }
    
    // Trader1 used by Catalog
    type Trader1 struct {

        
            Type string  `json:"type"`
            Address []string  `json:"address"`
            Name string  `json:"name"`
         
    }
    
    // ReturnConfig1 used by Catalog
    type ReturnConfig1 struct {

        
            Returnable bool  `json:"returnable"`
            Time float64  `json:"time"`
            Unit string  `json:"unit"`
         
    }
    
    // WeightResponse used by Catalog
    type WeightResponse struct {

        
            IsDefault bool  `json:"is_default"`
            Shipping float64  `json:"shipping"`
            Unit string  `json:"unit"`
         
    }
    
    // DimensionResponse used by Catalog
    type DimensionResponse struct {

        
            IsDefault bool  `json:"is_default"`
            Height float64  `json:"height"`
            Length float64  `json:"length"`
            Unit string  `json:"unit"`
            Width float64  `json:"width"`
         
    }
    
    // InventorySellerResponse used by Catalog
    type InventorySellerResponse struct {

        
            FyndItemCode string  `json:"fynd_item_code"`
            Price PriceMeta  `json:"price"`
            Brand BrandMeta  `json:"brand"`
            AddedOnStore string  `json:"added_on_store"`
            Stage string  `json:"stage"`
            Tags []string  `json:"tags"`
            FyndArticleCode string  `json:"fynd_article_code"`
            Set InventorySet  `json:"set"`
            Manufacturer ManufacturerResponse  `json:"manufacturer"`
            Quantities Quantities  `json:"quantities"`
            Identifier map[string]interface{}  `json:"identifier"`
            Meta map[string]interface{}  `json:"meta"`
            Store StoreMeta  `json:"store"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            CreatedBy UserSerializer  `json:"created_by"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            TraceID string  `json:"trace_id"`
            IsSet bool  `json:"is_set"`
            Company CompanyMeta  `json:"company"`
            TotalQuantity float64  `json:"total_quantity"`
            Trader []Trader1  `json:"trader"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            ExpirationDate string  `json:"expiration_date"`
            IsActive bool  `json:"is_active"`
            RawMeta map[string]interface{}  `json:"raw_meta"`
            UID string  `json:"uid"`
            ItemID float64  `json:"item_id"`
            Size string  `json:"size"`
            Fragile bool  `json:"fragile"`
            SellerIdentifier string  `json:"seller_identifier"`
            FyndMeta map[string]interface{}  `json:"fynd_meta"`
            TrackInventory bool  `json:"track_inventory"`
            ReturnConfig ReturnConfig1  `json:"return_config"`
            Weight WeightResponse  `json:"weight"`
            CountryOfOrigin string  `json:"country_of_origin"`
            Dimension DimensionResponse  `json:"dimension"`
         
    }
    
    // InventorySellerIdentifierResponsePaginated used by Catalog
    type InventorySellerIdentifierResponsePaginated struct {

        
            Page Page  `json:"page"`
            Items []InventorySellerResponse  `json:"items"`
         
    }
    
    // PriceArticle used by Catalog
    type PriceArticle struct {

        
            TpNotes map[string]interface{}  `json:"tp_notes"`
            Currency string  `json:"currency"`
            Effective float64  `json:"effective"`
            Transfer float64  `json:"transfer"`
            Marked float64  `json:"marked"`
         
    }
    
    // BrandMeta1 used by Catalog
    type BrandMeta1 struct {

        
            ID float64  `json:"id"`
            Name string  `json:"name"`
         
    }
    
    // ManufacturerResponse1 used by Catalog
    type ManufacturerResponse1 struct {

        
            IsDefault bool  `json:"is_default"`
            Address string  `json:"address"`
            Name string  `json:"name"`
         
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
    
    // ArticleStoreResponse used by Catalog
    type ArticleStoreResponse struct {

        
            StoreType string  `json:"store_type"`
            Name string  `json:"name"`
            UID float64  `json:"uid"`
            StoreCode string  `json:"store_code"`
         
    }
    
    // CompanyMeta1 used by Catalog
    type CompanyMeta1 struct {

        
            ID float64  `json:"id"`
         
    }
    
    // Trader2 used by Catalog
    type Trader2 struct {

        
            Type string  `json:"type"`
            Address []string  `json:"address"`
            Name string  `json:"name"`
         
    }
    
    // DateMeta used by Catalog
    type DateMeta struct {

        
            AddedOnStore string  `json:"added_on_store"`
            ModifiedOn string  `json:"modified_on"`
            InventoryUpdatedOn string  `json:"inventory_updated_on"`
            CreatedOn string  `json:"created_on"`
         
    }
    
    // ReturnConfig2 used by Catalog
    type ReturnConfig2 struct {

        
            Returnable bool  `json:"returnable"`
            Time float64  `json:"time"`
            Unit string  `json:"unit"`
         
    }
    
    // WeightResponse1 used by Catalog
    type WeightResponse1 struct {

        
            Shipping float64  `json:"shipping"`
            Unit string  `json:"unit"`
         
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

        
            Price PriceArticle  `json:"price"`
            Brand BrandMeta1  `json:"brand"`
            Tags []string  `json:"tags"`
            Stage string  `json:"stage"`
            Manufacturer ManufacturerResponse1  `json:"manufacturer"`
            Quantities QuantitiesArticle  `json:"quantities"`
            ID string  `json:"id"`
            Identifier map[string]interface{}  `json:"identifier"`
            Store ArticleStoreResponse  `json:"store"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            CreatedBy UserSerializer  `json:"created_by"`
            Platforms map[string]interface{}  `json:"platforms"`
            TraceID string  `json:"trace_id"`
            IsSet bool  `json:"is_set"`
            Company CompanyMeta1  `json:"company"`
            TotalQuantity float64  `json:"total_quantity"`
            Trader []Trader2  `json:"trader"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            ExpirationDate string  `json:"expiration_date"`
            InventoryUpdatedOn string  `json:"inventory_updated_on"`
            UID string  `json:"uid"`
            ItemID float64  `json:"item_id"`
            Size string  `json:"size"`
            SellerIdentifier string  `json:"seller_identifier"`
            TrackInventory bool  `json:"track_inventory"`
            DateMeta DateMeta  `json:"date_meta"`
            ReturnConfig ReturnConfig2  `json:"return_config"`
            Weight WeightResponse1  `json:"weight"`
            CountryOfOrigin string  `json:"country_of_origin"`
            Dimension DimensionResponse1  `json:"dimension"`
         
    }
    
    // GetInventoriesResponse used by Catalog
    type GetInventoriesResponse struct {

        
            Page Page  `json:"page"`
            Items []GetInventories  `json:"items"`
         
    }
    
    // BulkInventoryGetItems used by Catalog
    type BulkInventoryGetItems struct {

        
            CompanyID float64  `json:"company_id"`
            FailedRecords []string  `json:"failed_records"`
            Failed float64  `json:"failed"`
            Stage string  `json:"stage"`
            Succeed float64  `json:"succeed"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            IsActive bool  `json:"is_active"`
            CreatedOn string  `json:"created_on"`
            Total float64  `json:"total"`
            ID string  `json:"id"`
            FilePath string  `json:"file_path"`
            Cancelled float64  `json:"cancelled"`
            ModifiedOn string  `json:"modified_on"`
            CancelledRecords []string  `json:"cancelled_records"`
         
    }
    
    // BulkInventoryGet used by Catalog
    type BulkInventoryGet struct {

        
            Page Page  `json:"page"`
            Items []BulkInventoryGetItems  `json:"items"`
         
    }
    
    // InventoryJobPayload used by Catalog
    type InventoryJobPayload struct {

        
            Price float64  `json:"price"`
            Tags []string  `json:"tags"`
            ItemWeightUnitOfMeasure string  `json:"item_weight_unit_of_measure"`
            SellerIdentifier string  `json:"seller_identifier"`
            ExpirationDate string  `json:"expiration_date"`
            Quantity float64  `json:"quantity"`
            Currency string  `json:"currency"`
            PriceEffective float64  `json:"price_effective"`
            PriceMarked float64  `json:"price_marked"`
            StoreCode string  `json:"store_code"`
            TraceID string  `json:"trace_id"`
            TotalQuantity float64  `json:"total_quantity"`
            ItemDimensionsUnitOfMeasure string  `json:"item_dimensions_unit_of_measure"`
         
    }
    
    // InventoryBulkRequest used by Catalog
    type InventoryBulkRequest struct {

        
            Sizes []InventoryJobPayload  `json:"sizes"`
            BatchID string  `json:"batch_id"`
            CompanyID float64  `json:"company_id"`
            User map[string]interface{}  `json:"user"`
         
    }
    
    // InventoryExportQuantityFilter used by Catalog
    type InventoryExportQuantityFilter struct {

        
            Min float64  `json:"min"`
            Operators string  `json:"operators"`
            Max float64  `json:"max"`
         
    }
    
    // InventoryExportAdvanceOption used by Catalog
    type InventoryExportAdvanceOption struct {

        
            StoreIds []float64  `json:"store_ids"`
            BrandIds []float64  `json:"brand_ids"`
            Quantity InventoryExportQuantityFilter  `json:"quantity"`
            FromDate string  `json:"from_date"`
            ToDate string  `json:"to_date"`
         
    }
    
    // InventoryExportJob used by Catalog
    type InventoryExportJob struct {

        
            Type string  `json:"type"`
            CompletedOn string  `json:"completed_on"`
            TaskID string  `json:"task_id"`
            NotificationEmails []string  `json:"notification_emails"`
            Status string  `json:"status"`
            URL string  `json:"url"`
            Filters InventoryExportAdvanceOption  `json:"filters"`
            SellerID float64  `json:"seller_id"`
         
    }
    
    // InventoryExportRequest used by Catalog
    type InventoryExportRequest struct {

        
            Type string  `json:"type"`
            Brand []float64  `json:"brand"`
            Store []float64  `json:"store"`
         
    }
    
    // InventoryExportResponse used by Catalog
    type InventoryExportResponse struct {

        
            Type string  `json:"type"`
            TaskID string  `json:"task_id"`
            CreatedBy string  `json:"created_by"`
            NotificationEmails []string  `json:"notification_emails"`
            CreatedOn string  `json:"created_on"`
            Status string  `json:"status"`
            Filters map[string]interface{}  `json:"filters"`
            SellerID float64  `json:"seller_id"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // InventoryJobFilters used by Catalog
    type InventoryJobFilters struct {

        
            Quantity InventoryExportQuantityFilter  `json:"quantity"`
            Brands []string  `json:"brands"`
            Stores []string  `json:"stores"`
            FromDate string  `json:"from_date"`
            ToDate string  `json:"to_date"`
         
    }
    
    // InventoryJobDetailResponse used by Catalog
    type InventoryJobDetailResponse struct {

        
            Type string  `json:"type"`
            CompletedOn string  `json:"completed_on"`
            CancelledBy UserDetail  `json:"cancelled_by"`
            TaskID string  `json:"task_id"`
            CreatedBy UserDetail  `json:"created_by"`
            NotificationEmails []string  `json:"notification_emails"`
            CreatedOn string  `json:"created_on"`
            Status string  `json:"status"`
            URL string  `json:"url"`
            Filters InventoryJobFilters  `json:"filters"`
            ID string  `json:"id"`
            SellerID float64  `json:"seller_id"`
            ModifiedOn string  `json:"modified_on"`
            CancelledOn string  `json:"cancelled_on"`
         
    }
    
    // InventoryExportJobListResponse used by Catalog
    type InventoryExportJobListResponse struct {

        
            Items InventoryJobDetailResponse  `json:"items"`
         
    }
    
    // InventoryExportFilter used by Catalog
    type InventoryExportFilter struct {

        
            StoreIds []float64  `json:"store_ids"`
            BrandIds []float64  `json:"brand_ids"`
            Quantity InventoryExportQuantityFilter  `json:"quantity"`
            FromDate string  `json:"from_date"`
            ToDate string  `json:"to_date"`
         
    }
    
    // InventoryCreateRequest used by Catalog
    type InventoryCreateRequest struct {

        
            Type string  `json:"type"`
            Filters InventoryExportFilter  `json:"filters"`
            NotificationEmails []string  `json:"notification_emails"`
            Data []string  `json:"data"`
         
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

        
            Tags []string  `json:"tags"`
            SellerIdentifier string  `json:"seller_identifier"`
            ExpirationDate string  `json:"expiration_date"`
            PriceEffective float64  `json:"price_effective"`
            PriceMarked float64  `json:"price_marked"`
            StoreID float64  `json:"store_id"`
            TraceID string  `json:"trace_id"`
            TotalQuantity float64  `json:"total_quantity"`
         
    }
    
    // InventoryRequestSchemaV2 used by Catalog
    type InventoryRequestSchemaV2 struct {

        
            CompanyID float64  `json:"company_id"`
            Meta map[string]interface{}  `json:"meta"`
            Payload []InventoryPayload  `json:"payload"`
         
    }
    
    // InventoryFailedReason used by Catalog
    type InventoryFailedReason struct {

        
            Errors string  `json:"errors"`
            Message string  `json:"message"`
         
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

        
            Tax1 float64  `json:"tax1"`
            Tax2 float64  `json:"tax2"`
            CompanyID float64  `json:"company_id"`
            Threshold1 float64  `json:"threshold1"`
            Hs2Code string  `json:"hs2_code"`
            Threshold2 float64  `json:"threshold2"`
            TaxOnMrp bool  `json:"tax_on_mrp"`
            IsActive bool  `json:"is_active"`
            TaxOnEsp bool  `json:"tax_on_esp"`
            UID float64  `json:"uid"`
            HsnCode string  `json:"hsn_code"`
         
    }
    
    // HsnCodesObject used by Catalog
    type HsnCodesObject struct {

        
            Tax1 float64  `json:"tax1"`
            Tax2 float64  `json:"tax2"`
            CompanyID float64  `json:"company_id"`
            Threshold1 float64  `json:"threshold1"`
            Hs2Code string  `json:"hs2_code"`
            Threshold2 float64  `json:"threshold2"`
            TaxOnMrp bool  `json:"tax_on_mrp"`
            TaxOnEsp bool  `json:"tax_on_esp"`
            ID string  `json:"id"`
            HsnCode string  `json:"hsn_code"`
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
    
    // PageResponse used by Catalog
    type PageResponse struct {

        
            HasPrevious bool  `json:"has_previous"`
            ItemTotal float64  `json:"item_total"`
            HasNext bool  `json:"has_next"`
            Current string  `json:"current"`
            Size float64  `json:"size"`
         
    }
    
    // TaxSlab used by Catalog
    type TaxSlab struct {

        
            Rate float64  `json:"rate"`
            EffectiveDate string  `json:"effective_date"`
            Cess float64  `json:"cess"`
            Threshold float64  `json:"threshold"`
         
    }
    
    // HSNDataInsertV2 used by Catalog
    type HSNDataInsertV2 struct {

        
            Type string  `json:"type"`
            HsnCodeID string  `json:"hsn_code_id"`
            Taxes []TaxSlab  `json:"taxes"`
            ReportingHsn string  `json:"reporting_hsn"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            CreatedOn string  `json:"created_on"`
            CountryCode string  `json:"country_code"`
            HsnCode string  `json:"hsn_code"`
            ModifiedOn string  `json:"modified_on"`
            Description string  `json:"description"`
         
    }
    
    // HsnCodesListingResponseSchemaV2 used by Catalog
    type HsnCodesListingResponseSchemaV2 struct {

        
            Page PageResponse  `json:"page"`
            Items []HSNDataInsertV2  `json:"items"`
         
    }
    
    // BrandItem used by Catalog
    type BrandItem struct {

        
            Departments []string  `json:"departments"`
            Action Action  `json:"action"`
            Name string  `json:"name"`
            Banners ImageUrls  `json:"banners"`
            Slug string  `json:"slug"`
            Discount string  `json:"discount"`
            Logo Media1  `json:"logo"`
            UID float64  `json:"uid"`
         
    }
    
    // BrandListingResponse used by Catalog
    type BrandListingResponse struct {

        
            Page Page  `json:"page"`
            Items []BrandItem  `json:"items"`
         
    }
    
    // Department used by Catalog
    type Department struct {

        
            Name string  `json:"name"`
            PriorityOrder float64  `json:"priority_order"`
            Slug string  `json:"slug"`
            Logo Media1  `json:"logo"`
            UID float64  `json:"uid"`
         
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

        
            Action Action  `json:"action"`
            Name string  `json:"name"`
            Banners ImageUrls  `json:"banners"`
            Childs []map[string]interface{}  `json:"childs"`
            Slug string  `json:"slug"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            UID float64  `json:"uid"`
         
    }
    
    // SecondLevelChild used by Catalog
    type SecondLevelChild struct {

        
            Action Action  `json:"action"`
            Name string  `json:"name"`
            Banners ImageUrls  `json:"banners"`
            Childs []ThirdLevelChild  `json:"childs"`
            Slug string  `json:"slug"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            UID float64  `json:"uid"`
         
    }
    
    // Child used by Catalog
    type Child struct {

        
            Action Action  `json:"action"`
            Name string  `json:"name"`
            Banners ImageUrls  `json:"banners"`
            Childs []SecondLevelChild  `json:"childs"`
            Slug string  `json:"slug"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            UID float64  `json:"uid"`
         
    }
    
    // CategoryItems used by Catalog
    type CategoryItems struct {

        
            Action Action  `json:"action"`
            Name string  `json:"name"`
            Banners ImageUrls  `json:"banners"`
            Childs []Child  `json:"childs"`
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
         
    }
    
    // DepartmentCategoryTree used by Catalog
    type DepartmentCategoryTree struct {

        
            Department string  `json:"department"`
            Items []CategoryItems  `json:"items"`
         
    }
    
    // CategoryListingResponse used by Catalog
    type CategoryListingResponse struct {

        
            Departments []DepartmentIdentifier  `json:"departments"`
            Data []DepartmentCategoryTree  `json:"data"`
         
    }
    
    // ApplicationProductListingResponse used by Catalog
    type ApplicationProductListingResponse struct {

        
            Items []ProductListingDetail  `json:"items"`
            Page Page  `json:"page"`
            SortOn []ProductSortOn  `json:"sort_on"`
            Filters []ProductFilters  `json:"filters"`
            Operators map[string]interface{}  `json:"operators"`
         
    }
    
    // ProductDetail used by Catalog
    type ProductDetail struct {

        
            Brand ProductBrand  `json:"brand"`
            Name string  `json:"name"`
            Rating float64  `json:"rating"`
            ProductOnlineDate string  `json:"product_online_date"`
            Color string  `json:"color"`
            ItemCode string  `json:"item_code"`
            PromoMeta map[string]interface{}  `json:"promo_meta"`
            Description string  `json:"description"`
            Type string  `json:"type"`
            RatingCount float64  `json:"rating_count"`
            Tryouts []string  `json:"tryouts"`
            ShortDescription string  `json:"short_description"`
            Highlights []string  `json:"highlights"`
            ItemType string  `json:"item_type"`
            Slug string  `json:"slug"`
            HasVariant bool  `json:"has_variant"`
            UID float64  `json:"uid"`
            Attributes map[string]interface{}  `json:"attributes"`
            ImageNature string  `json:"image_nature"`
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            Medias []Media2  `json:"medias"`
            Similars []string  `json:"similars"`
         
    }
    
    // InventoryPage used by Catalog
    type InventoryPage struct {

        
            Type string  `json:"type"`
            NextID string  `json:"next_id"`
            HasPrevious bool  `json:"has_previous"`
            ItemTotal float64  `json:"item_total"`
            HasNext bool  `json:"has_next"`
         
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

        
            Query ArticleQuery  `json:"query"`
            Quantity float64  `json:"quantity"`
            ArticleAssignment ArticleAssignment  `json:"article_assignment"`
            GroupID string  `json:"group_id"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // AssignStore used by Catalog
    type AssignStore struct {

        
            CompanyID float64  `json:"company_id"`
            StoreIds []float64  `json:"store_ids"`
            Articles []AssignStoreArticle  `json:"articles"`
            Pincode string  `json:"pincode"`
            AppID string  `json:"app_id"`
            ChannelType string  `json:"channel_type"`
            ChannelIdentifier string  `json:"channel_identifier"`
         
    }
    
    // ArticleAssignment1 used by Catalog
    type ArticleAssignment1 struct {

        
            Strategy string  `json:"strategy"`
            Level string  `json:"level"`
         
    }
    
    // StoreAssignResponse used by Catalog
    type StoreAssignResponse struct {

        
            StrategyWiseListing []map[string]interface{}  `json:"strategy_wise_listing"`
            CompanyID float64  `json:"company_id"`
            SCity string  `json:"s_city"`
            StorePincode float64  `json:"store_pincode"`
            ID string  `json:"_id"`
            Quantity float64  `json:"quantity"`
            ArticleAssignment ArticleAssignment1  `json:"article_assignment"`
            PriceEffective float64  `json:"price_effective"`
            PriceMarked float64  `json:"price_marked"`
            Index float64  `json:"index"`
            Status bool  `json:"status"`
            StoreID float64  `json:"store_id"`
            UID string  `json:"uid"`
            GroupID string  `json:"group_id"`
            ItemID float64  `json:"item_id"`
            Meta map[string]interface{}  `json:"meta"`
            Size string  `json:"size"`
         
    }
    
    // Document used by Catalog
    type Document struct {

        
            Type string  `json:"type"`
            Verified bool  `json:"verified"`
            Value string  `json:"value"`
            LegalName string  `json:"legal_name"`
            URL string  `json:"url"`
         
    }
    
    // SellerPhoneNumber used by Catalog
    type SellerPhoneNumber struct {

        
            CountryCode float64  `json:"country_code"`
            Number string  `json:"number"`
         
    }
    
    // GetAddressSerializer used by Catalog
    type GetAddressSerializer struct {

        
            Landmark string  `json:"landmark"`
            Longitude float64  `json:"longitude"`
            City string  `json:"city"`
            AddressType string  `json:"address_type"`
            Address1 string  `json:"address1"`
            Pincode float64  `json:"pincode"`
            Address2 string  `json:"address2"`
            CountryCode string  `json:"country_code"`
            State string  `json:"state"`
            Latitude float64  `json:"latitude"`
            Country string  `json:"country"`
         
    }
    
    // UserSerializer1 used by Catalog
    type UserSerializer1 struct {

        
            UserID string  `json:"user_id"`
            Contact string  `json:"contact"`
            Username string  `json:"username"`
         
    }
    
    // UserSerializer2 used by Catalog
    type UserSerializer2 struct {

        
            UserID string  `json:"user_id"`
            Contact string  `json:"contact"`
            Username string  `json:"username"`
         
    }
    
    // GetCompanySerializer used by Catalog
    type GetCompanySerializer struct {

        
            Stage string  `json:"stage"`
            Name string  `json:"name"`
            CompanyType string  `json:"company_type"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            ModifiedBy UserSerializer2  `json:"modified_by"`
            CreatedBy UserSerializer2  `json:"created_by"`
            VerifiedBy UserSerializer2  `json:"verified_by"`
            CreatedOn string  `json:"created_on"`
            UID float64  `json:"uid"`
            BusinessType string  `json:"business_type"`
            RejectReason string  `json:"reject_reason"`
            VerifiedOn string  `json:"verified_on"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // LocationManagerSerializer used by Catalog
    type LocationManagerSerializer struct {

        
            MobileNo SellerPhoneNumber  `json:"mobile_no"`
            Name string  `json:"name"`
            Email string  `json:"email"`
         
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

        
            Minute float64  `json:"minute"`
            Hour float64  `json:"hour"`
         
    }
    
    // LocationDayWiseSerializer used by Catalog
    type LocationDayWiseSerializer struct {

        
            Opening LocationTimingSerializer  `json:"opening"`
            Open bool  `json:"open"`
            Closing LocationTimingSerializer  `json:"closing"`
            Weekday string  `json:"weekday"`
         
    }
    
    // LocationIntegrationType used by Catalog
    type LocationIntegrationType struct {

        
            Inventory string  `json:"inventory"`
            Order string  `json:"order"`
         
    }
    
    // ProductReturnConfigSerializer used by Catalog
    type ProductReturnConfigSerializer struct {

        
            StoreUID float64  `json:"store_uid"`
            OnSameStore bool  `json:"on_same_store"`
         
    }
    
    // GetLocationSerializer used by Catalog
    type GetLocationSerializer struct {

        
            Documents []Document  `json:"documents"`
            Stage string  `json:"stage"`
            Name string  `json:"name"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            Address GetAddressSerializer  `json:"address"`
            ModifiedBy UserSerializer1  `json:"modified_by"`
            CreatedBy UserSerializer1  `json:"created_by"`
            VerifiedBy UserSerializer1  `json:"verified_by"`
            Code string  `json:"code"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Company GetCompanySerializer  `json:"company"`
            Manager LocationManagerSerializer  `json:"manager"`
            DisplayName string  `json:"display_name"`
            NotificationEmails []string  `json:"notification_emails"`
            CreatedOn string  `json:"created_on"`
            Warnings map[string]interface{}  `json:"warnings"`
            GstCredentials InvoiceDetailsSerializer  `json:"gst_credentials"`
            UID float64  `json:"uid"`
            Timing []LocationDayWiseSerializer  `json:"timing"`
            StoreType string  `json:"store_type"`
            IntegrationType LocationIntegrationType  `json:"integration_type"`
            PhoneNumber string  `json:"phone_number"`
            VerifiedOn string  `json:"verified_on"`
            ModifiedOn string  `json:"modified_on"`
            ProductReturnConfig ProductReturnConfigSerializer  `json:"product_return_config"`
         
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

        
            Name string  `json:"name"`
            IsActive bool  `json:"is_active"`
            Logo string  `json:"logo"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            AppID string  `json:"app_id"`
            UID float64  `json:"uid"`
         
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

        
            Type string  `json:"type"`
            Value string  `json:"value"`
            LegalName string  `json:"legal_name"`
            URL string  `json:"url"`
            Verified bool  `json:"verified"`
         
    }
    
    // BusinessCountryInfo used by CompanyProfile
    type BusinessCountryInfo struct {

        
            Country string  `json:"country"`
            CountryCode string  `json:"country_code"`
         
    }
    
    // CompanyTaxesSerializer used by CompanyProfile
    type CompanyTaxesSerializer struct {

        
            Rate float64  `json:"rate"`
            EffectiveDate string  `json:"effective_date"`
            Enable bool  `json:"enable"`
         
    }
    
    // GetAddressSerializer used by CompanyProfile
    type GetAddressSerializer struct {

        
            AddressType string  `json:"address_type"`
            Address2 string  `json:"address2"`
            Pincode float64  `json:"pincode"`
            Country string  `json:"country"`
            State string  `json:"state"`
            Latitude float64  `json:"latitude"`
            Landmark string  `json:"landmark"`
            City string  `json:"city"`
            CountryCode string  `json:"country_code"`
            Longitude float64  `json:"longitude"`
            Address1 string  `json:"address1"`
         
    }
    
    // GetCompanyProfileSerializerResponse used by CompanyProfile
    type GetCompanyProfileSerializerResponse struct {

        
            ModifiedBy UserSerializer  `json:"modified_by"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            ContactDetails ContactDetails  `json:"contact_details"`
            CreatedOn string  `json:"created_on"`
            Mode string  `json:"mode"`
            BusinessDetails BusinessDetails  `json:"business_details"`
            BusinessInfo string  `json:"business_info"`
            ModifiedOn string  `json:"modified_on"`
            UID float64  `json:"uid"`
            Warnings map[string]interface{}  `json:"warnings"`
            CreatedBy UserSerializer  `json:"created_by"`
            Stage string  `json:"stage"`
            Documents []Document  `json:"documents"`
            BusinessCountryInfo BusinessCountryInfo  `json:"business_country_info"`
            BusinessType string  `json:"business_type"`
            Taxes []CompanyTaxesSerializer  `json:"taxes"`
            VerifiedOn string  `json:"verified_on"`
            NotificationEmails []string  `json:"notification_emails"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            Name string  `json:"name"`
            CompanyType string  `json:"company_type"`
            FranchiseEnabled bool  `json:"franchise_enabled"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // ErrorResponse used by CompanyProfile
    type ErrorResponse struct {

        
            Code string  `json:"code"`
            Status float64  `json:"status"`
            Message string  `json:"message"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // CreateUpdateAddressSerializer used by CompanyProfile
    type CreateUpdateAddressSerializer struct {

        
            AddressType string  `json:"address_type"`
            Address2 string  `json:"address2"`
            Pincode float64  `json:"pincode"`
            Country string  `json:"country"`
            State string  `json:"state"`
            Latitude float64  `json:"latitude"`
            Landmark string  `json:"landmark"`
            Address1 string  `json:"address1"`
            Longitude float64  `json:"longitude"`
            City string  `json:"city"`
            CountryCode string  `json:"country_code"`
         
    }
    
    // CompanyTaxesSerializer1 used by CompanyProfile
    type CompanyTaxesSerializer1 struct {

        
            Rate float64  `json:"rate"`
            EffectiveDate string  `json:"effective_date"`
            Enable bool  `json:"enable"`
         
    }
    
    // UpdateCompany used by CompanyProfile
    type UpdateCompany struct {

        
            Documents []Document  `json:"documents"`
            Addresses []CreateUpdateAddressSerializer  `json:"addresses"`
            Warnings map[string]interface{}  `json:"warnings"`
            ContactDetails ContactDetails  `json:"contact_details"`
            BusinessType string  `json:"business_type"`
            Name string  `json:"name"`
            CompanyType string  `json:"company_type"`
            Taxes []CompanyTaxesSerializer1  `json:"taxes"`
            FranchiseEnabled bool  `json:"franchise_enabled"`
            RejectReason string  `json:"reject_reason"`
            BusinessDetails BusinessDetails  `json:"business_details"`
            BusinessInfo string  `json:"business_info"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            NotificationEmails []string  `json:"notification_emails"`
         
    }
    
    // ProfileSuccessResponse used by CompanyProfile
    type ProfileSuccessResponse struct {

        
            Success bool  `json:"success"`
            UID float64  `json:"uid"`
            Message string  `json:"message"`
         
    }
    
    // DocumentsObj used by CompanyProfile
    type DocumentsObj struct {

        
            Verified float64  `json:"verified"`
            Pending float64  `json:"pending"`
         
    }
    
    // MetricsSerializer used by CompanyProfile
    type MetricsSerializer struct {

        
            UID float64  `json:"uid"`
            Store DocumentsObj  `json:"store"`
            CompanyDocuments DocumentsObj  `json:"company_documents"`
            StoreDocuments DocumentsObj  `json:"store_documents"`
            Brand DocumentsObj  `json:"brand"`
            Stage string  `json:"stage"`
            Product DocumentsObj  `json:"product"`
         
    }
    
    // BrandBannerSerializer used by CompanyProfile
    type BrandBannerSerializer struct {

        
            Landscape string  `json:"landscape"`
            Portrait string  `json:"portrait"`
         
    }
    
    // GetBrandResponseSerializer used by CompanyProfile
    type GetBrandResponseSerializer struct {

        
            ModifiedBy UserSerializer  `json:"modified_by"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            CreatedOn string  `json:"created_on"`
            Synonyms []string  `json:"synonyms"`
            Mode string  `json:"mode"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            Banner BrandBannerSerializer  `json:"banner"`
            ModifiedOn string  `json:"modified_on"`
            UID float64  `json:"uid"`
            Warnings map[string]interface{}  `json:"warnings"`
            CreatedBy UserSerializer  `json:"created_by"`
            Stage string  `json:"stage"`
            SlugKey string  `json:"slug_key"`
            Logo string  `json:"logo"`
            RejectReason string  `json:"reject_reason"`
            VerifiedOn string  `json:"verified_on"`
            Description string  `json:"description"`
            Name string  `json:"name"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // CreateUpdateBrandRequestSerializer used by CompanyProfile
    type CreateUpdateBrandRequestSerializer struct {

        
            Description string  `json:"description"`
            UID float64  `json:"uid"`
            Logo string  `json:"logo"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            Name string  `json:"name"`
            Synonyms []string  `json:"synonyms"`
            BrandTier string  `json:"brand_tier"`
            CompanyID float64  `json:"company_id"`
            Banner BrandBannerSerializer  `json:"banner"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // CompanySocialAccounts used by CompanyProfile
    type CompanySocialAccounts struct {

        
            Name string  `json:"name"`
            URL string  `json:"url"`
         
    }
    
    // CompanyDetails used by CompanyProfile
    type CompanyDetails struct {

        
            Socials []CompanySocialAccounts  `json:"socials"`
            WebsiteURL string  `json:"website_url"`
         
    }
    
    // CompanySerializer used by CompanyProfile
    type CompanySerializer struct {

        
            UID float64  `json:"uid"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            BusinessCountryInfo BusinessCountryInfo  `json:"business_country_info"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            BusinessType string  `json:"business_type"`
            MarketChannels []string  `json:"market_channels"`
            CreatedOn string  `json:"created_on"`
            CreatedBy UserSerializer  `json:"created_by"`
            Name string  `json:"name"`
            CompanyType string  `json:"company_type"`
            Details CompanyDetails  `json:"details"`
            Stage string  `json:"stage"`
            RejectReason string  `json:"reject_reason"`
            VerifiedOn string  `json:"verified_on"`
            ModifiedOn string  `json:"modified_on"`
            NotificationEmails []string  `json:"notification_emails"`
         
    }
    
    // CompanyBrandSerializer used by CompanyProfile
    type CompanyBrandSerializer struct {

        
            UID float64  `json:"uid"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            Company CompanySerializer  `json:"company"`
            Warnings map[string]interface{}  `json:"warnings"`
            CreatedOn string  `json:"created_on"`
            CreatedBy UserSerializer  `json:"created_by"`
            Brand GetBrandResponseSerializer  `json:"brand"`
            Stage string  `json:"stage"`
            RejectReason string  `json:"reject_reason"`
            VerifiedOn string  `json:"verified_on"`
            ModifiedOn string  `json:"modified_on"`
         
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

        
            UID float64  `json:"uid"`
            Company float64  `json:"company"`
            Brands []float64  `json:"brands"`
         
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
    
    // LocationManagerSerializer used by CompanyProfile
    type LocationManagerSerializer struct {

        
            Name string  `json:"name"`
            Email string  `json:"email"`
            MobileNo SellerPhoneNumber  `json:"mobile_no"`
         
    }
    
    // GetCompanySerializer used by CompanyProfile
    type GetCompanySerializer struct {

        
            UID float64  `json:"uid"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            BusinessType string  `json:"business_type"`
            CreatedOn string  `json:"created_on"`
            CreatedBy UserSerializer  `json:"created_by"`
            Name string  `json:"name"`
            CompanyType string  `json:"company_type"`
            Stage string  `json:"stage"`
            RejectReason string  `json:"reject_reason"`
            VerifiedOn string  `json:"verified_on"`
            ModifiedOn string  `json:"modified_on"`
         
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
    
    // ProductReturnConfigSerializer used by CompanyProfile
    type ProductReturnConfigSerializer struct {

        
            OnSameStore bool  `json:"on_same_store"`
            StoreUID float64  `json:"store_uid"`
         
    }
    
    // LocationTimingSerializer used by CompanyProfile
    type LocationTimingSerializer struct {

        
            Hour float64  `json:"hour"`
            Minute float64  `json:"minute"`
         
    }
    
    // LocationDayWiseSerializer used by CompanyProfile
    type LocationDayWiseSerializer struct {

        
            Closing LocationTimingSerializer  `json:"closing"`
            Opening LocationTimingSerializer  `json:"opening"`
            Open bool  `json:"open"`
            Weekday string  `json:"weekday"`
         
    }
    
    // GetLocationSerializer used by CompanyProfile
    type GetLocationSerializer struct {

        
            DisplayName string  `json:"display_name"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
            StoreType string  `json:"store_type"`
            AutoInvoice bool  `json:"auto_invoice"`
            UID float64  `json:"uid"`
            Warnings map[string]interface{}  `json:"warnings"`
            GstCredentials InvoiceDetailsSerializer  `json:"gst_credentials"`
            CreatedBy UserSerializer  `json:"created_by"`
            Stage string  `json:"stage"`
            Manager LocationManagerSerializer  `json:"manager"`
            CreditNote bool  `json:"credit_note"`
            Code string  `json:"code"`
            PhoneNumber string  `json:"phone_number"`
            Documents []Document  `json:"documents"`
            Company GetCompanySerializer  `json:"company"`
            Address GetAddressSerializer  `json:"address"`
            VerifiedOn string  `json:"verified_on"`
            NotificationEmails []string  `json:"notification_emails"`
            Holiday []HolidaySchemaSerializer  `json:"holiday"`
            ProductReturnConfig ProductReturnConfigSerializer  `json:"product_return_config"`
            Timing []LocationDayWiseSerializer  `json:"timing"`
            Name string  `json:"name"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
         
    }
    
    // LocationListSerializer used by CompanyProfile
    type LocationListSerializer struct {

        
            Items []GetLocationSerializer  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // AddressSerializer used by CompanyProfile
    type AddressSerializer struct {

        
            AddressType string  `json:"address_type"`
            Address2 string  `json:"address2"`
            Pincode float64  `json:"pincode"`
            Country string  `json:"country"`
            State string  `json:"state"`
            Latitude float64  `json:"latitude"`
            Landmark string  `json:"landmark"`
            City string  `json:"city"`
            CountryCode string  `json:"country_code"`
            Longitude float64  `json:"longitude"`
            Address1 string  `json:"address1"`
         
    }
    
    // LocationSerializer used by CompanyProfile
    type LocationSerializer struct {

        
            DisplayName string  `json:"display_name"`
            AutoInvoice bool  `json:"auto_invoice"`
            StoreType string  `json:"store_type"`
            UID float64  `json:"uid"`
            Warnings map[string]interface{}  `json:"warnings"`
            GstCredentials InvoiceDetailsSerializer  `json:"gst_credentials"`
            Stage string  `json:"stage"`
            Manager LocationManagerSerializer  `json:"manager"`
            CreditNote bool  `json:"credit_note"`
            Code string  `json:"code"`
            Documents []Document  `json:"documents"`
            Company float64  `json:"company"`
            Address AddressSerializer  `json:"address"`
            NotificationEmails []string  `json:"notification_emails"`
            Holiday []HolidaySchemaSerializer  `json:"holiday"`
            ProductReturnConfig ProductReturnConfigSerializer  `json:"product_return_config"`
            Timing []LocationDayWiseSerializer  `json:"timing"`
            Name string  `json:"name"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
         
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
            CodThresholdAmount float64  `json:"cod_threshold_amount"`
            OnlineThresholdAmount float64  `json:"online_threshold_amount"`
         
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
    
    // Identifier used by Cart
    type Identifier struct {

        
            StoreID []float64  `json:"store_id"`
            ExcludeBrandID []float64  `json:"exclude_brand_id"`
            CategoryID []float64  `json:"category_id"`
            ArticleID []string  `json:"article_id"`
            UserID []string  `json:"user_id"`
            CollectionID []string  `json:"collection_id"`
            ItemID []float64  `json:"item_id"`
            BrandID []float64  `json:"brand_id"`
            CompanyID []float64  `json:"company_id"`
         
    }
    
    // Rule used by Cart
    type Rule struct {

        
            Key float64  `json:"key"`
            Max float64  `json:"max"`
            DiscountQty float64  `json:"discount_qty"`
            Value float64  `json:"value"`
            Min float64  `json:"min"`
         
    }
    
    // Validity used by Cart
    type Validity struct {

        
            Priority float64  `json:"priority"`
         
    }
    
    // Ownership used by Cart
    type Ownership struct {

        
            PayableBy string  `json:"payable_by"`
            PayableCategory string  `json:"payable_category"`
         
    }
    
    // RuleDefinition used by Cart
    type RuleDefinition struct {

        
            AutoApply bool  `json:"auto_apply"`
            ValueType string  `json:"value_type"`
            Type string  `json:"type"`
            CurrencyCode string  `json:"currency_code"`
            IsExact bool  `json:"is_exact"`
            ApplicableOn string  `json:"applicable_on"`
            Scope []string  `json:"scope"`
            CalculateOn string  `json:"calculate_on"`
         
    }
    
    // CouponAuthor used by Cart
    type CouponAuthor struct {

        
            ModifiedBy string  `json:"modified_by"`
            CreatedBy string  `json:"created_by"`
         
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
    
    // PaymentAllowValue used by Cart
    type PaymentAllowValue struct {

        
            Max float64  `json:"max"`
         
    }
    
    // PaymentModes used by Cart
    type PaymentModes struct {

        
            Types []string  `json:"types"`
            Codes []string  `json:"codes"`
            Networks []string  `json:"networks"`
            Uses PaymentAllowValue  `json:"uses"`
         
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
    
    // Restrictions used by Cart
    type Restrictions struct {

        
            BulkBundle BulkBundleRestriction  `json:"bulk_bundle"`
            UserGroups []float64  `json:"user_groups"`
            Uses UsesRestriction  `json:"uses"`
            CouponAllowed bool  `json:"coupon_allowed"`
            Payments map[string]PaymentModes  `json:"payments"`
            PriceRange PriceRange  `json:"price_range"`
            OrderingStores []float64  `json:"ordering_stores"`
            PostOrder PostOrder  `json:"post_order"`
            Platforms []string  `json:"platforms"`
         
    }
    
    // Validation used by Cart
    type Validation struct {

        
            AppID []string  `json:"app_id"`
            Anonymous bool  `json:"anonymous"`
            UserRegisteredAfter string  `json:"user_registered_after"`
         
    }
    
    // State used by Cart
    type State struct {

        
            IsDisplay bool  `json:"is_display"`
            IsPublic bool  `json:"is_public"`
            IsArchived bool  `json:"is_archived"`
         
    }
    
    // CouponAction used by Cart
    type CouponAction struct {

        
            ActionDate string  `json:"action_date"`
            TxnMode string  `json:"txn_mode"`
         
    }
    
    // DisplayMetaDict used by Cart
    type DisplayMetaDict struct {

        
            Subtitle string  `json:"subtitle"`
            Title string  `json:"title"`
         
    }
    
    // DisplayMeta used by Cart
    type DisplayMeta struct {

        
            Auto DisplayMetaDict  `json:"auto"`
            Subtitle string  `json:"subtitle"`
            Description string  `json:"description"`
            Remove DisplayMetaDict  `json:"remove"`
            Title string  `json:"title"`
            Apply DisplayMetaDict  `json:"apply"`
         
    }
    
    // CouponSchedule used by Cart
    type CouponSchedule struct {

        
            Duration float64  `json:"duration"`
            NextSchedule []map[string]interface{}  `json:"next_schedule"`
            End string  `json:"end"`
            Cron string  `json:"cron"`
            Start string  `json:"start"`
         
    }
    
    // CouponAdd used by Cart
    type CouponAdd struct {

        
            Identifiers Identifier  `json:"identifiers"`
            Rule []Rule  `json:"rule"`
            Validity Validity  `json:"validity"`
            Ownership Ownership  `json:"ownership"`
            RuleDefinition RuleDefinition  `json:"rule_definition"`
            Author CouponAuthor  `json:"author"`
            Code string  `json:"code"`
            DateMeta CouponDateMeta  `json:"date_meta"`
            Restrictions Restrictions  `json:"restrictions"`
            Validation Validation  `json:"validation"`
            State State  `json:"state"`
            TypeSlug string  `json:"type_slug"`
            Action CouponAction  `json:"action"`
            DisplayMeta DisplayMeta  `json:"display_meta"`
            Schedule CouponSchedule  `json:"_schedule"`
            Tags []string  `json:"tags"`
         
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

        
            Identifiers Identifier  `json:"identifiers"`
            Rule []Rule  `json:"rule"`
            Validity Validity  `json:"validity"`
            Ownership Ownership  `json:"ownership"`
            RuleDefinition RuleDefinition  `json:"rule_definition"`
            Author CouponAuthor  `json:"author"`
            Code string  `json:"code"`
            DateMeta CouponDateMeta  `json:"date_meta"`
            Restrictions Restrictions  `json:"restrictions"`
            Validation Validation  `json:"validation"`
            State State  `json:"state"`
            TypeSlug string  `json:"type_slug"`
            Action CouponAction  `json:"action"`
            DisplayMeta DisplayMeta  `json:"display_meta"`
            Schedule CouponSchedule  `json:"_schedule"`
            Tags []string  `json:"tags"`
         
    }
    
    // CouponPartialUpdate used by Cart
    type CouponPartialUpdate struct {

        
            Schedule CouponSchedule  `json:"schedule"`
            Archive bool  `json:"archive"`
         
    }
    
    // PromotionDateMeta used by Cart
    type PromotionDateMeta struct {

        
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // CompareObject used by Cart
    type CompareObject struct {

        
            LessThan float64  `json:"less_than"`
            Equals float64  `json:"equals"`
            LessThanEquals float64  `json:"less_than_equals"`
            GreaterThan float64  `json:"greater_than"`
            GreaterThanEquals float64  `json:"greater_than_equals"`
         
    }
    
    // ItemCriteria used by Cart
    type ItemCriteria struct {

        
            ItemExcludeCategory []float64  `json:"item_exclude_category"`
            ItemExcludeBrand []float64  `json:"item_exclude_brand"`
            ItemExcludeCompany []float64  `json:"item_exclude_company"`
            ItemExcludeStore []float64  `json:"item_exclude_store"`
            ItemExcludeSku []string  `json:"item_exclude_sku"`
            BuyRules []string  `json:"buy_rules"`
            ItemCategory []float64  `json:"item_category"`
            ItemStore []float64  `json:"item_store"`
            AvailableZones []string  `json:"available_zones"`
            ItemSku []string  `json:"item_sku"`
            ItemID []float64  `json:"item_id"`
            CartQuantity CompareObject  `json:"cart_quantity"`
            ItemCompany []float64  `json:"item_company"`
            CartUniqueItemAmount CompareObject  `json:"cart_unique_item_amount"`
            AllItems bool  `json:"all_items"`
            CartUniqueItemQuantity CompareObject  `json:"cart_unique_item_quantity"`
            CartTotal CompareObject  `json:"cart_total"`
            ItemExcludeID []float64  `json:"item_exclude_id"`
            ItemSize []string  `json:"item_size"`
            ItemBrand []float64  `json:"item_brand"`
         
    }
    
    // PromotionSchedule used by Cart
    type PromotionSchedule struct {

        
            Duration float64  `json:"duration"`
            NextSchedule []map[string]interface{}  `json:"next_schedule"`
            End string  `json:"end"`
            Published bool  `json:"published"`
            Cron string  `json:"cron"`
            Start string  `json:"start"`
         
    }
    
    // DiscountOffer used by Cart
    type DiscountOffer struct {

        
            MinOfferQuantity float64  `json:"min_offer_quantity"`
            Code string  `json:"code"`
            DiscountPercentage float64  `json:"discount_percentage"`
            MaxDiscountAmount float64  `json:"max_discount_amount"`
            MaxUsagePerTransaction float64  `json:"max_usage_per_transaction"`
            DiscountPrice float64  `json:"discount_price"`
            PartialCanRet bool  `json:"partial_can_ret"`
            DiscountAmount float64  `json:"discount_amount"`
            ApportionDiscount bool  `json:"apportion_discount"`
            MaxOfferQuantity float64  `json:"max_offer_quantity"`
         
    }
    
    // DiscountRule used by Cart
    type DiscountRule struct {

        
            BuyCondition string  `json:"buy_condition"`
            DiscountType string  `json:"discount_type"`
            Offer DiscountOffer  `json:"offer"`
            ItemCriteria ItemCriteria  `json:"item_criteria"`
         
    }
    
    // Visibility used by Cart
    type Visibility struct {

        
            Pdp bool  `json:"pdp"`
            CouponList bool  `json:"coupon_list"`
         
    }
    
    // Ownership1 used by Cart
    type Ownership1 struct {

        
            PayableBy string  `json:"payable_by"`
            PayableCategory string  `json:"payable_category"`
         
    }
    
    // PromotionAuthor used by Cart
    type PromotionAuthor struct {

        
            ModifiedBy string  `json:"modified_by"`
            CreatedBy string  `json:"created_by"`
         
    }
    
    // PromotionAction used by Cart
    type PromotionAction struct {

        
            ActionDate string  `json:"action_date"`
            ActionType string  `json:"action_type"`
         
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
            Type string  `json:"type"`
            Uses PaymentAllowValue1  `json:"uses"`
         
    }
    
    // PostOrder1 used by Cart
    type PostOrder1 struct {

        
            CancellationAllowed bool  `json:"cancellation_allowed"`
            ReturnAllowed bool  `json:"return_allowed"`
         
    }
    
    // Restrictions1 used by Cart
    type Restrictions1 struct {

        
            UserGroups []float64  `json:"user_groups"`
            Uses UsesRestriction1  `json:"uses"`
            AnonymousUsers bool  `json:"anonymous_users"`
            UserRegistered UserRegistered  `json:"user_registered"`
            Payments []PromotionPaymentModes  `json:"payments"`
            UserID []string  `json:"user_id"`
            OrderQuantity float64  `json:"order_quantity"`
            OrderingStores []float64  `json:"ordering_stores"`
            PostOrder PostOrder1  `json:"post_order"`
            Platforms []string  `json:"platforms"`
         
    }
    
    // DisplayMeta1 used by Cart
    type DisplayMeta1 struct {

        
            OfferText string  `json:"offer_text"`
            Description string  `json:"description"`
            Name string  `json:"name"`
         
    }
    
    // PromotionListItem used by Cart
    type PromotionListItem struct {

        
            DateMeta PromotionDateMeta  `json:"date_meta"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            BuyRules map[string]ItemCriteria  `json:"buy_rules"`
            Schedule PromotionSchedule  `json:"_schedule"`
            DiscountRules []DiscountRule  `json:"discount_rules"`
            Visiblility Visibility  `json:"visiblility"`
            CalculateOn string  `json:"calculate_on"`
            Ownership Ownership1  `json:"ownership"`
            PromoGroup string  `json:"promo_group"`
            ApplyAllDiscount bool  `json:"apply_all_discount"`
            Author PromotionAuthor  `json:"author"`
            PostOrderAction PromotionAction  `json:"post_order_action"`
            Currency string  `json:"currency"`
            ApplyExclusive string  `json:"apply_exclusive"`
            Mode string  `json:"mode"`
            Restrictions Restrictions1  `json:"restrictions"`
            Stackable bool  `json:"stackable"`
            PromotionType string  `json:"promotion_type"`
            ApplicationID string  `json:"application_id"`
            ApplyPriority float64  `json:"apply_priority"`
            Code string  `json:"code"`
            DisplayMeta DisplayMeta1  `json:"display_meta"`
         
    }
    
    // PromotionsResponse used by Cart
    type PromotionsResponse struct {

        
            Page Page  `json:"page"`
            Items PromotionListItem  `json:"items"`
         
    }
    
    // PromotionAdd used by Cart
    type PromotionAdd struct {

        
            DateMeta PromotionDateMeta  `json:"date_meta"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            BuyRules map[string]ItemCriteria  `json:"buy_rules"`
            Schedule PromotionSchedule  `json:"_schedule"`
            DiscountRules []DiscountRule  `json:"discount_rules"`
            Visiblility Visibility  `json:"visiblility"`
            CalculateOn string  `json:"calculate_on"`
            Ownership Ownership1  `json:"ownership"`
            PromoGroup string  `json:"promo_group"`
            ApplyAllDiscount bool  `json:"apply_all_discount"`
            Author PromotionAuthor  `json:"author"`
            PostOrderAction PromotionAction  `json:"post_order_action"`
            Currency string  `json:"currency"`
            ApplyExclusive string  `json:"apply_exclusive"`
            Mode string  `json:"mode"`
            Restrictions Restrictions1  `json:"restrictions"`
            Stackable bool  `json:"stackable"`
            PromotionType string  `json:"promotion_type"`
            ApplicationID string  `json:"application_id"`
            ApplyPriority float64  `json:"apply_priority"`
            Code string  `json:"code"`
            DisplayMeta DisplayMeta1  `json:"display_meta"`
         
    }
    
    // PromotionUpdate used by Cart
    type PromotionUpdate struct {

        
            DateMeta PromotionDateMeta  `json:"date_meta"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            BuyRules map[string]ItemCriteria  `json:"buy_rules"`
            Schedule PromotionSchedule  `json:"_schedule"`
            DiscountRules []DiscountRule  `json:"discount_rules"`
            Visiblility Visibility  `json:"visiblility"`
            CalculateOn string  `json:"calculate_on"`
            Ownership Ownership1  `json:"ownership"`
            PromoGroup string  `json:"promo_group"`
            ApplyAllDiscount bool  `json:"apply_all_discount"`
            Author PromotionAuthor  `json:"author"`
            PostOrderAction PromotionAction  `json:"post_order_action"`
            Currency string  `json:"currency"`
            ApplyExclusive string  `json:"apply_exclusive"`
            Mode string  `json:"mode"`
            Restrictions Restrictions1  `json:"restrictions"`
            Stackable bool  `json:"stackable"`
            PromotionType string  `json:"promotion_type"`
            ApplicationID string  `json:"application_id"`
            ApplyPriority float64  `json:"apply_priority"`
            Code string  `json:"code"`
            DisplayMeta DisplayMeta1  `json:"display_meta"`
         
    }
    
    // PromotionPartialUpdate used by Cart
    type PromotionPartialUpdate struct {

        
            Schedule PromotionSchedule  `json:"schedule"`
            Archive bool  `json:"archive"`
         
    }
    
    // ActivePromosResponse used by Cart
    type ActivePromosResponse struct {

        
            CreatedOn string  `json:"created_on"`
            Subtitle string  `json:"subtitle"`
            Description string  `json:"description"`
            IsHidden bool  `json:"is_hidden"`
            Type string  `json:"type"`
            Title string  `json:"title"`
            EntitySlug string  `json:"entity_slug"`
            EntityType string  `json:"entity_type"`
            Example string  `json:"example"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // Charges used by Cart
    type Charges struct {

        
            Charges float64  `json:"charges"`
            Threshold float64  `json:"threshold"`
         
    }
    
    // DeliveryCharges used by Cart
    type DeliveryCharges struct {

        
            Charges []Charges  `json:"charges"`
            Enabled bool  `json:"enabled"`
         
    }
    
    // CartMetaConfigUpdate used by Cart
    type CartMetaConfigUpdate struct {

        
            DeliveryCharges DeliveryCharges  `json:"delivery_charges"`
            RevenueEngineCoupon bool  `json:"revenue_engine_coupon"`
            GiftPricing float64  `json:"gift_pricing"`
            MinCartValue float64  `json:"min_cart_value"`
            MaxCartItems float64  `json:"max_cart_items"`
            Enabled bool  `json:"enabled"`
            BulkCoupons bool  `json:"bulk_coupons"`
            GiftDisplayText string  `json:"gift_display_text"`
         
    }
    
    // CartMetaConfigAdd used by Cart
    type CartMetaConfigAdd struct {

        
            DeliveryCharges DeliveryCharges  `json:"delivery_charges"`
            RevenueEngineCoupon bool  `json:"revenue_engine_coupon"`
            GiftPricing float64  `json:"gift_pricing"`
            MinCartValue float64  `json:"min_cart_value"`
            MaxCartItems float64  `json:"max_cart_items"`
            Enabled bool  `json:"enabled"`
            BulkCoupons bool  `json:"bulk_coupons"`
            GiftDisplayText string  `json:"gift_display_text"`
         
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
    
    // PromoMeta used by Cart
    type PromoMeta struct {

        
            Message string  `json:"message"`
         
    }
    
    // CartProductIdentifer used by Cart
    type CartProductIdentifer struct {

        
            Identifier string  `json:"identifier"`
         
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
    
    // ProductPrice used by Cart
    type ProductPrice struct {

        
            Effective float64  `json:"effective"`
            Marked float64  `json:"marked"`
            CurrencySymbol string  `json:"currency_symbol"`
            Selling float64  `json:"selling"`
            AddOn float64  `json:"add_on"`
            CurrencyCode string  `json:"currency_code"`
         
    }
    
    // ProductPriceInfo used by Cart
    type ProductPriceInfo struct {

        
            Converted ProductPrice  `json:"converted"`
            Base ProductPrice  `json:"base"`
         
    }
    
    // CategoryInfo used by Cart
    type CategoryInfo struct {

        
            Name string  `json:"name"`
            UID float64  `json:"uid"`
         
    }
    
    // BaseInfo used by Cart
    type BaseInfo struct {

        
            Name string  `json:"name"`
            UID float64  `json:"uid"`
         
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
    
    // ProductImage used by Cart
    type ProductImage struct {

        
            URL string  `json:"url"`
            SecureURL string  `json:"secure_url"`
            AspectRatio string  `json:"aspect_ratio"`
         
    }
    
    // CartProduct used by Cart
    type CartProduct struct {

        
            UID float64  `json:"uid"`
            Categories []CategoryInfo  `json:"categories"`
            Type string  `json:"type"`
            Slug string  `json:"slug"`
            Name string  `json:"name"`
            Brand BaseInfo  `json:"brand"`
            Action ProductAction  `json:"action"`
            Images []ProductImage  `json:"images"`
         
    }
    
    // BasePrice used by Cart
    type BasePrice struct {

        
            Effective float64  `json:"effective"`
            Marked float64  `json:"marked"`
            CurrencyCode string  `json:"currency_code"`
            CurrencySymbol string  `json:"currency_symbol"`
         
    }
    
    // ArticlePriceInfo used by Cart
    type ArticlePriceInfo struct {

        
            Converted BasePrice  `json:"converted"`
            Base BasePrice  `json:"base"`
         
    }
    
    // ProductArticle used by Cart
    type ProductArticle struct {

        
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            Size string  `json:"size"`
            Identifier map[string]interface{}  `json:"identifier"`
            UID string  `json:"uid"`
            SellerIdentifier string  `json:"seller_identifier"`
            Store BaseInfo  `json:"store"`
            Type string  `json:"type"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            Seller BaseInfo  `json:"seller"`
            Quantity float64  `json:"quantity"`
            Price ArticlePriceInfo  `json:"price"`
            ProductGroupTags []string  `json:"product_group_tags"`
         
    }
    
    // Ownership2 used by Cart
    type Ownership2 struct {

        
            PayableBy string  `json:"payable_by"`
            PayableCategory string  `json:"payable_category"`
         
    }
    
    // FreeGiftItem used by Cart
    type FreeGiftItem struct {

        
            ItemPriceDetails map[string]interface{}  `json:"item_price_details"`
            ItemName string  `json:"item_name"`
            ItemImagesURL []string  `json:"item_images_url"`
            ItemBrandName string  `json:"item_brand_name"`
            ItemID float64  `json:"item_id"`
            ItemSlug string  `json:"item_slug"`
         
    }
    
    // AppliedFreeArticles used by Cart
    type AppliedFreeArticles struct {

        
            ArticleID string  `json:"article_id"`
            Quantity float64  `json:"quantity"`
            ParentItemIdentifier string  `json:"parent_item_identifier"`
            FreeGiftItemDetails FreeGiftItem  `json:"free_gift_item_details"`
         
    }
    
    // BuyRules used by Cart
    type BuyRules struct {

        
            ItemCriteria map[string]interface{}  `json:"item_criteria"`
            CartConditions map[string]interface{}  `json:"cart_conditions"`
         
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

        
            PromoID string  `json:"promo_id"`
            PromotionGroup string  `json:"promotion_group"`
            ArticleQuantity float64  `json:"article_quantity"`
            Ownership Ownership2  `json:"ownership"`
            Amount float64  `json:"amount"`
            PromotionName string  `json:"promotion_name"`
            AppliedFreeArticles []AppliedFreeArticles  `json:"applied_free_articles"`
            PromotionType string  `json:"promotion_type"`
            BuyRules []BuyRules  `json:"buy_rules"`
            MrpPromotion bool  `json:"mrp_promotion"`
            OfferText string  `json:"offer_text"`
            DiscountRules []DiscountRulesApp  `json:"discount_rules"`
         
    }
    
    // ProductAvailabilitySize used by Cart
    type ProductAvailabilitySize struct {

        
            Display string  `json:"display"`
            IsAvailable bool  `json:"is_available"`
            Value string  `json:"value"`
         
    }
    
    // ProductAvailability used by Cart
    type ProductAvailability struct {

        
            OutOfStock bool  `json:"out_of_stock"`
            OtherStoreQuantity float64  `json:"other_store_quantity"`
            Deliverable bool  `json:"deliverable"`
            Sizes []string  `json:"sizes"`
            AvailableSizes []ProductAvailabilitySize  `json:"available_sizes"`
            IsValid bool  `json:"is_valid"`
         
    }
    
    // CartProductInfo used by Cart
    type CartProductInfo struct {

        
            PromoMeta PromoMeta  `json:"promo_meta"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
            Key string  `json:"key"`
            BulkOffer map[string]interface{}  `json:"bulk_offer"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            Moq map[string]interface{}  `json:"moq"`
            PricePerUnit ProductPriceInfo  `json:"price_per_unit"`
            Product CartProduct  `json:"product"`
            Article ProductArticle  `json:"article"`
            CouponMessage string  `json:"coupon_message"`
            Message string  `json:"message"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            PromotionsApplied []AppliedPromotion  `json:"promotions_applied"`
            Discount string  `json:"discount"`
            IsSet bool  `json:"is_set"`
            Quantity float64  `json:"quantity"`
            Price ProductPriceInfo  `json:"price"`
            Availability ProductAvailability  `json:"availability"`
         
    }
    
    // RawBreakup used by Cart
    type RawBreakup struct {

        
            YouSaved float64  `json:"you_saved"`
            DeliveryCharge float64  `json:"delivery_charge"`
            ConvenienceFee float64  `json:"convenience_fee"`
            FyndCash float64  `json:"fynd_cash"`
            GstCharges float64  `json:"gst_charges"`
            Vog float64  `json:"vog"`
            CodCharge float64  `json:"cod_charge"`
            Discount float64  `json:"discount"`
            MrpTotal float64  `json:"mrp_total"`
            Coupon float64  `json:"coupon"`
            Subtotal float64  `json:"subtotal"`
            Total float64  `json:"total"`
         
    }
    
    // CouponBreakup used by Cart
    type CouponBreakup struct {

        
            IsApplied bool  `json:"is_applied"`
            CouponValue float64  `json:"coupon_value"`
            UID string  `json:"uid"`
            Value float64  `json:"value"`
            MinimumCartValue float64  `json:"minimum_cart_value"`
            Description string  `json:"description"`
            SubTitle string  `json:"sub_title"`
            CouponType string  `json:"coupon_type"`
            Code string  `json:"code"`
            Type string  `json:"type"`
            Title string  `json:"title"`
            Message string  `json:"message"`
            MaxDiscountValue float64  `json:"max_discount_value"`
         
    }
    
    // DisplayBreakup used by Cart
    type DisplayBreakup struct {

        
            Key string  `json:"key"`
            CurrencySymbol string  `json:"currency_symbol"`
            Value float64  `json:"value"`
            Message []string  `json:"message"`
            CurrencyCode string  `json:"currency_code"`
            Display string  `json:"display"`
         
    }
    
    // LoyaltyPoints used by Cart
    type LoyaltyPoints struct {

        
            Description string  `json:"description"`
            IsApplied bool  `json:"is_applied"`
            Applicable float64  `json:"applicable"`
            Total float64  `json:"total"`
         
    }
    
    // CartBreakup used by Cart
    type CartBreakup struct {

        
            Raw RawBreakup  `json:"raw"`
            Coupon CouponBreakup  `json:"coupon"`
            Display []DisplayBreakup  `json:"display"`
            LoyaltyPoints LoyaltyPoints  `json:"loyalty_points"`
         
    }
    
    // OpenapiCartDetailsResponse used by Cart
    type OpenapiCartDetailsResponse struct {

        
            Message string  `json:"message"`
            Items []CartProductInfo  `json:"items"`
            IsValid bool  `json:"is_valid"`
            BreakupValues CartBreakup  `json:"breakup_values"`
         
    }
    
    // OpenApiErrorResponse used by Cart
    type OpenApiErrorResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
            Errors map[string]interface{}  `json:"errors"`
         
    }
    
    // ShippingAddress used by Cart
    type ShippingAddress struct {

        
            Landmark string  `json:"landmark"`
            CountryCode string  `json:"country_code"`
            Address string  `json:"address"`
            Meta map[string]interface{}  `json:"meta"`
            CountryPhoneCode string  `json:"country_phone_code"`
            Phone float64  `json:"phone"`
            Area string  `json:"area"`
            CountryIsoCode string  `json:"country_iso_code"`
            AreaCode string  `json:"area_code"`
            State string  `json:"state"`
            AreaCodeSlug string  `json:"area_code_slug"`
            Name string  `json:"name"`
            Email string  `json:"email"`
            AddressType string  `json:"address_type"`
            City string  `json:"city"`
            Pincode float64  `json:"pincode"`
            Country string  `json:"country"`
         
    }
    
    // OpenApiCartServiceabilityRequest used by Cart
    type OpenApiCartServiceabilityRequest struct {

        
            CartItems []CartItem  `json:"cart_items"`
            ShippingAddress ShippingAddress  `json:"shipping_address"`
         
    }
    
    // OpenApiCartServiceabilityResponse used by Cart
    type OpenApiCartServiceabilityResponse struct {

        
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            Message string  `json:"message"`
            IsValid bool  `json:"is_valid"`
            Items []CartProductInfo  `json:"items"`
            BreakupValues CartBreakup  `json:"breakup_values"`
         
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
    
    // MultiTenderPaymentMeta used by Cart
    type MultiTenderPaymentMeta struct {

        
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            PaymentID string  `json:"payment_id"`
            PaymentGateway string  `json:"payment_gateway"`
            OrderID string  `json:"order_id"`
            CurrentStatus string  `json:"current_status"`
         
    }
    
    // MultiTenderPaymentMethod used by Cart
    type MultiTenderPaymentMethod struct {

        
            Meta MultiTenderPaymentMeta  `json:"meta"`
            Mode string  `json:"mode"`
            Amount float64  `json:"amount"`
            Name string  `json:"name"`
         
    }
    
    // OpenApiOrderItem used by Cart
    type OpenApiOrderItem struct {

        
            PriceMarked float64  `json:"price_marked"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            PriceEffective float64  `json:"price_effective"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            DeliveryCharges float64  `json:"delivery_charges"`
            AmountPaid float64  `json:"amount_paid"`
            Meta CartItemMeta  `json:"meta"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            CashbackApplied float64  `json:"cashback_applied"`
            Files []OpenApiFiles  `json:"files"`
            ProductID float64  `json:"product_id"`
            Discount float64  `json:"discount"`
            EmployeeDiscount float64  `json:"employee_discount"`
            Quantity float64  `json:"quantity"`
            Size string  `json:"size"`
            PaymentMethods []MultiTenderPaymentMethod  `json:"payment_methods"`
            CodCharges float64  `json:"cod_charges"`
         
    }
    
    // OpenApiPlatformCheckoutReq used by Cart
    type OpenApiPlatformCheckoutReq struct {

        
            CouponValue float64  `json:"coupon_value"`
            DeliveryCharges float64  `json:"delivery_charges"`
            CouponCode string  `json:"coupon_code"`
            CurrencyCode string  `json:"currency_code"`
            Files []OpenApiFiles  `json:"files"`
            ShippingAddress ShippingAddress  `json:"shipping_address"`
            Coupon string  `json:"coupon"`
            CodCharges float64  `json:"cod_charges"`
            Comment string  `json:"comment"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            CashbackApplied float64  `json:"cashback_applied"`
            PaymentMode string  `json:"payment_mode"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            CartItems []OpenApiOrderItem  `json:"cart_items"`
            EmployeeDiscount map[string]interface{}  `json:"employee_discount"`
            Gstin string  `json:"gstin"`
            PaymentMethods []MultiTenderPaymentMethod  `json:"payment_methods"`
            OrderID string  `json:"order_id"`
            BillingAddress ShippingAddress  `json:"billing_address"`
            CartValue float64  `json:"cart_value"`
         
    }
    
    // OpenApiCheckoutResponse used by Cart
    type OpenApiCheckoutResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
            OrderID string  `json:"order_id"`
            OrderRefID string  `json:"order_ref_id"`
         
    }
    
    // AbandonedCart used by Cart
    type AbandonedCart struct {

        
            UID float64  `json:"uid"`
            DeliveryCharges map[string]interface{}  `json:"delivery_charges"`
            Discount float64  `json:"discount"`
            Coupon map[string]interface{}  `json:"coupon"`
            ID string  `json:"_id"`
            BuyNow bool  `json:"buy_now"`
            Comment string  `json:"comment"`
            ExpireAt string  `json:"expire_at"`
            BulkCouponDiscount float64  `json:"bulk_coupon_discount"`
            Shipments []map[string]interface{}  `json:"shipments"`
            Articles []map[string]interface{}  `json:"articles"`
            IsDefault bool  `json:"is_default"`
            AppID string  `json:"app_id"`
            PaymentMethods []map[string]interface{}  `json:"payment_methods"`
            IsArchive bool  `json:"is_archive"`
            FcIndexMap []float64  `json:"fc_index_map"`
            CreatedOn string  `json:"created_on"`
            Cashback map[string]interface{}  `json:"cashback"`
            Meta map[string]interface{}  `json:"meta"`
            LastModified string  `json:"last_modified"`
            IsActive bool  `json:"is_active"`
            Promotion map[string]interface{}  `json:"promotion"`
            Payments map[string]interface{}  `json:"payments"`
            CheckoutMode string  `json:"checkout_mode"`
            CodCharges map[string]interface{}  `json:"cod_charges"`
            MergeQty bool  `json:"merge_qty"`
            PaymentMode string  `json:"payment_mode"`
            UserID string  `json:"user_id"`
            Gstin string  `json:"gstin"`
            PickUpCustomerDetails map[string]interface{}  `json:"pick_up_customer_details"`
            OrderID string  `json:"order_id"`
            FyndCredits map[string]interface{}  `json:"fynd_credits"`
            CartValue float64  `json:"cart_value"`
         
    }
    
    // AbandonedCartResponse used by Cart
    type AbandonedCartResponse struct {

        
            Success bool  `json:"success"`
            Page Page  `json:"page"`
            Message string  `json:"message"`
            Result map[string]interface{}  `json:"result"`
            Items []AbandonedCart  `json:"items"`
         
    }
    
    // CartCurrency used by Cart
    type CartCurrency struct {

        
            Code string  `json:"code"`
            Symbol string  `json:"symbol"`
         
    }
    
    // PaymentSelectionLock used by Cart
    type PaymentSelectionLock struct {

        
            DefaultOptions string  `json:"default_options"`
            Enabled bool  `json:"enabled"`
            PaymentIdentifier string  `json:"payment_identifier"`
         
    }
    
    // CartDetailResponse used by Cart
    type CartDetailResponse struct {

        
            Currency CartCurrency  `json:"currency"`
            BuyNow bool  `json:"buy_now"`
            LastModified string  `json:"last_modified"`
            Comment string  `json:"comment"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            AppliedPromoDetails []AppliedPromotion  `json:"applied_promo_details"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            Message string  `json:"message"`
            CouponText string  `json:"coupon_text"`
            Gstin string  `json:"gstin"`
            Items []CartProductInfo  `json:"items"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            IsValid bool  `json:"is_valid"`
            CheckoutMode string  `json:"checkout_mode"`
            ID string  `json:"id"`
            RestrictCheckout bool  `json:"restrict_checkout"`
         
    }
    
    // AddProductCart used by Cart
    type AddProductCart struct {

        
            SellerID float64  `json:"seller_id"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            StoreID float64  `json:"store_id"`
            Pos bool  `json:"pos"`
            ArticleID string  `json:"article_id"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ArticleAssignment map[string]interface{}  `json:"article_assignment"`
            ParentItemIdentifiers []map[string]string  `json:"parent_item_identifiers"`
            ItemID float64  `json:"item_id"`
            Display string  `json:"display"`
            Quantity float64  `json:"quantity"`
            ItemSize string  `json:"item_size"`
            ProductGroupTags []string  `json:"product_group_tags"`
         
    }
    
    // AddCartRequest used by Cart
    type AddCartRequest struct {

        
            Items []AddProductCart  `json:"items"`
         
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

        
            Identifiers CartProductIdentifer  `json:"identifiers"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            ArticleID string  `json:"article_id"`
            ItemIndex float64  `json:"item_index"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            ItemID float64  `json:"item_id"`
            Quantity float64  `json:"quantity"`
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
    
    // OverrideCartItemPromo used by Cart
    type OverrideCartItemPromo struct {

        
            PromoDesc string  `json:"promo_desc"`
            PromoID string  `json:"promo_id"`
            RwrdTndr string  `json:"rwrd_tndr"`
            ItemList []map[string]interface{}  `json:"item_list"`
            PromoAmount string  `json:"promo_amount"`
         
    }
    
    // OverrideCartItem used by Cart
    type OverrideCartItem struct {

        
            PriceMarked float64  `json:"price_marked"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            PriceEffective float64  `json:"price_effective"`
            AmountPaid float64  `json:"amount_paid"`
            PromoList []OverrideCartItemPromo  `json:"promo_list"`
            SellerIdentifier string  `json:"seller_identifier"`
            ItemID float64  `json:"item_id"`
            Discount float64  `json:"discount"`
            Quantity float64  `json:"quantity"`
            Size string  `json:"size"`
         
    }
    
    // OverrideCheckoutReq used by Cart
    type OverrideCheckoutReq struct {

        
            MerchantCode string  `json:"merchant_code"`
            Aggregator string  `json:"aggregator"`
            CartID string  `json:"cart_id"`
            PaymentIdentifier string  `json:"payment_identifier"`
            OrderType string  `json:"order_type"`
            OrderingStore float64  `json:"ordering_store"`
            PaymentMode string  `json:"payment_mode"`
            CurrencyCode string  `json:"currency_code"`
            CartItems []OverrideCartItem  `json:"cart_items"`
            ShippingAddress map[string]interface{}  `json:"shipping_address"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
         
    }
    
    // OverrideCheckoutResponse used by Cart
    type OverrideCheckoutResponse struct {

        
            Success string  `json:"success"`
            Data map[string]interface{}  `json:"data"`
            Message string  `json:"message"`
            Cart map[string]interface{}  `json:"cart"`
            OrderID string  `json:"order_id"`
         
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
    

    
    // ExtensionResponse used by Partner
    type ExtensionResponse struct {

        
            Items []ExtensionListItems  `json:"items"`
            Page Pagination  `json:"page"`
         
    }
    
    // ExtensionListItems used by Partner
    type ExtensionListItems struct {

        
            BaseURL string  `json:"base_url"`
            Callbacks Callback  `json:"callbacks"`
            ContactEmail string  `json:"contact_email"`
            CurrentStatus string  `json:"current_status"`
            Description string  `json:"description"`
            DevelopedByName string  `json:"developed_by_name"`
            ExtVersion string  `json:"ext_version"`
            ExtentionType string  `json:"extention_type"`
            IsApplicationLevel bool  `json:"is_application_level"`
            IsComingSoon bool  `json:"is_coming_soon"`
            IsSaleschannel bool  `json:"is_saleschannel"`
            Logo Logo  `json:"logo"`
            Name string  `json:"name"`
            Scope []string  `json:"scope"`
            CreatedAt string  `json:"created_at"`
            IsHidden bool  `json:"is_hidden"`
            ModifiedAt string  `json:"modified_at"`
            OrganizationID string  `json:"organization_id"`
            WhitelistedUrls []string  `json:"whitelisted_urls"`
            ID string  `json:"_id"`
         
    }
    
    // ExtensionCommon used by Partner
    type ExtensionCommon struct {

        
            BaseURL string  `json:"base_url"`
            Callbacks Callback  `json:"callbacks"`
            ContactEmail string  `json:"contact_email"`
            CurrentStatus string  `json:"current_status"`
            Description string  `json:"description"`
            DevelopedByName string  `json:"developed_by_name"`
            ExtVersion string  `json:"ext_version"`
            ExtentionType string  `json:"extention_type"`
            IsApplicationLevel bool  `json:"is_application_level"`
            IsComingSoon bool  `json:"is_coming_soon"`
            IsSaleschannel bool  `json:"is_saleschannel"`
            Logo Logo  `json:"logo"`
            Name string  `json:"name"`
            Scope []string  `json:"scope"`
         
    }
    
    // ExtensionList used by Partner
    type ExtensionList struct {

        
            Items []ExtensionItems  `json:"items"`
            Page Pagination  `json:"page"`
         
    }
    
    // ExtensionItems used by Partner
    type ExtensionItems struct {

        
            BaseURL string  `json:"base_url"`
            Company string  `json:"company"`
            Description string  `json:"description"`
            DevelopedByName string  `json:"developed_by_name"`
            ExtVersion string  `json:"ext_version"`
            ExtensionID string  `json:"extension_id"`
            ExtentionType string  `json:"extention_type"`
            Installed bool  `json:"installed"`
            IsSaleschannel bool  `json:"is_saleschannel"`
            LaunchURL string  `json:"launch_url"`
            Name string  `json:"name"`
            Logo Logo  `json:"logo"`
            Scope []Scope  `json:"scope"`
         
    }
    
    // Scope used by Partner
    type Scope struct {

        
            Description string  `json:"description"`
            Name string  `json:"name"`
         
    }
    
    // Pagination used by Partner
    type Pagination struct {

        
            Current float64  `json:"current"`
            HasNext bool  `json:"has_next"`
            ItemTotal float64  `json:"item_total"`
            Size float64  `json:"size"`
            Type string  `json:"type"`
         
    }
    
    // ExtensionSuggestionList used by Partner
    type ExtensionSuggestionList struct {

        
            Items []ExtensionSuggestion  `json:"items"`
         
    }
    
    // OrganizationBasicInfo used by Partner
    type OrganizationBasicInfo struct {

        
            Name string  `json:"name"`
            Slug string  `json:"slug"`
         
    }
    
    // ExtensionSuggestion used by Partner
    type ExtensionSuggestion struct {

        
            ListingInfo ListingInfo  `json:"listing_info"`
            Organization OrganizationBasicInfo  `json:"organization"`
            OrganizationID string  `json:"organization_id"`
            Plans []Plan  `json:"plans"`
            Slug string  `json:"slug"`
         
    }
    
    // Plan used by Partner
    type Plan struct {

        
            AdditionalCharges string  `json:"additional_charges"`
            Features string  `json:"features"`
            Name string  `json:"name"`
            TrialDays float64  `json:"trial_days"`
            Type string  `json:"type"`
            Price Price  `json:"price"`
         
    }
    
    // ListingInfo used by Partner
    type ListingInfo struct {

        
            Icon string  `json:"icon"`
            Name string  `json:"name"`
            Tagline string  `json:"tagline"`
            Keywords []string  `json:"keywords"`
         
    }
    
    // Callback used by Partner
    type Callback struct {

        
            Auth string  `json:"auth"`
            AutoInstall string  `json:"auto_install"`
            Install string  `json:"install"`
            Setup string  `json:"setup"`
            Uninstall string  `json:"uninstall"`
         
    }
    
    // Logo used by Partner
    type Logo struct {

        
            Large string  `json:"large"`
            Small string  `json:"small"`
         
    }
    
    // Category used by Partner
    type Category struct {

        
            CategoryL1 []CategoryL1  `json:"category_l1"`
            CategoryL2 []CategoryL2  `json:"category_l2"`
         
    }
    
    // CommingSoon used by Partner
    type CommingSoon struct {

        
            IsComingSoon bool  `json:"is_coming_soon"`
            UpvoteCount float64  `json:"upvote_count"`
         
    }
    
    // ContactInfo used by Partner
    type ContactInfo struct {

        
            Support Support  `json:"support"`
         
    }
    
    // Benefits used by Partner
    type Benefits struct {

        
            Title string  `json:"title"`
            Description string  `json:"description"`
         
    }
    
    // Screenshots used by Partner
    type Screenshots struct {

        
            Desktop []string  `json:"desktop"`
            Mobile []string  `json:"mobile"`
         
    }
    
    // ExtensionDetails used by Partner
    type ExtensionDetails struct {

        
            Benefits []Benefits  `json:"benefits"`
            DemoURL string  `json:"demo_url"`
            Description string  `json:"description"`
            Integration []string  `json:"integration"`
            VideoURL []map[string]interface{}  `json:"video_url"`
            Youtube []string  `json:"youtube"`
            Screenshots Screenshots  `json:"screenshots"`
         
    }
    
    // Plans used by Partner
    type Plans struct {

        
            AdditionalCharges string  `json:"additional_charges"`
            Features string  `json:"features"`
            Name string  `json:"name"`
            Price Price  `json:"price"`
            TrialDays float64  `json:"trial_days"`
            Type string  `json:"type"`
         
    }
    
    // PublicExtension used by Partner
    type PublicExtension struct {

        
            Category Category  `json:"category"`
            ComingSoon CommingSoon  `json:"coming_soon"`
            ContactInfo ContactInfo  `json:"contact_info"`
            CreatedAt string  `json:"created_at"`
            CurrentStatus string  `json:"current_status"`
            Details ExtensionDetails  `json:"details"`
            ExtensionID string  `json:"extension_id"`
            IsComingSoon bool  `json:"is_coming_soon"`
            ListingInfo ListingInfo  `json:"listing_info"`
            ModifiedAt string  `json:"modified_at"`
            Organization OrganizationBasicInfo  `json:"organization"`
            OrganizationID string  `json:"organization_id"`
            PlanType string  `json:"plan_type"`
            Plans []Plans  `json:"plans"`
            PlansURL string  `json:"plans_url"`
            ReviewInstructions string  `json:"review_instructions"`
            Scope []string  `json:"scope"`
            Slug string  `json:"slug"`
            ID string  `json:"_id"`
         
    }
    
    // CategoryL1 used by Partner
    type CategoryL1 struct {

        
            Description string  `json:"description"`
            Display string  `json:"display"`
            Level float64  `json:"level"`
            Logo string  `json:"logo"`
            Slug string  `json:"slug"`
            Value string  `json:"value"`
            ID string  `json:"_id"`
         
    }
    
    // CategoryL2 used by Partner
    type CategoryL2 struct {

        
            Parent string  `json:"parent"`
            Display string  `json:"display"`
            Level float64  `json:"level"`
            Slug string  `json:"slug"`
            Value string  `json:"value"`
            ID string  `json:"_id"`
         
    }
    
    // Support used by Partner
    type Support struct {

        
            Email string  `json:"email"`
            FaqURL string  `json:"faq_url"`
            Phone string  `json:"phone"`
            PrivacyPolicyURL string  `json:"privacy_policy_url"`
            WebsiteURL string  `json:"website_url"`
         
    }
    
    // Price used by Partner
    type Price struct {

        
            Amount float64  `json:"amount"`
            Currency string  `json:"currency"`
         
    }
    
    // UninstallExtension used by Partner
    type UninstallExtension struct {

        
            Success bool  `json:"success"`
         
    }
    
    // SubscriptionRequest used by Partner
    type SubscriptionRequest struct {

        
            Approved string  `json:"approved"`
            ClientID string  `json:"client_id"`
            CompanyID string  `json:"company_id"`
            CreditBalance string  `json:"credit_balance"`
            RequestID string  `json:"request_id"`
         
    }
    
    // SubscriptionRes used by Partner
    type SubscriptionRes struct {

        
            RedirectURL string  `json:"redirect_url"`
         
    }
    
    // PartnerInviteDetails used by Partner
    type PartnerInviteDetails struct {

        
            AccountType string  `json:"account_type"`
            ApprovedPermissions ApprovedPermissions  `json:"approved_permissions"`
            Comment string  `json:"comment"`
            CompanyID float64  `json:"company_id"`
            CompanyName string  `json:"company_name"`
            CreatedAt string  `json:"created_at"`
            ModifiedAt string  `json:"modified_at"`
            OrganizationID string  `json:"organization_id"`
            OrganizationName string  `json:"organization_name"`
            RequestStatus string  `json:"request_status"`
            RequestedPermissions RequestedPermissions  `json:"requested_permissions"`
            UserID string  `json:"user_id"`
            ID string  `json:"_id"`
         
    }
    
    // ApprovedPermissions used by Partner
    type ApprovedPermissions struct {

        
            ApplicationRole []string  `json:"application_role"`
            CompanyPermissions []string  `json:"company_permissions"`
            CompanyRole []string  `json:"company_role"`
         
    }
    
    // RequestedPermissions used by Partner
    type RequestedPermissions struct {

        
            ApplicationPermissions []string  `json:"application_permissions"`
            ApplicationRole []string  `json:"application_role"`
            CompanyPermissions []string  `json:"company_permissions"`
            CompanyRole []string  `json:"company_role"`
         
    }
    
    // ModifyPartnerReq used by Partner
    type ModifyPartnerReq struct {

        
            AccountType string  `json:"account_type"`
            ApprovedPermissions ApprovedPermissionsInfo  `json:"approved_permissions"`
            Comment string  `json:"comment"`
            CompanyID float64  `json:"company_id"`
            CompanyName string  `json:"company_name"`
            OrganizationID string  `json:"organization_id"`
            OrganizationName string  `json:"organization_name"`
            RequestStatus string  `json:"request_status"`
            RequestedPermissions RequestedPermissions  `json:"requested_permissions"`
            UserID string  `json:"user_id"`
            ID string  `json:"_id"`
         
    }
    
    // ApprovedPermissionsInfo used by Partner
    type ApprovedPermissionsInfo struct {

        
            ApplicationPermissions map[string]ApplicationPermissions  `json:"application_permissions"`
            CompanyPermissions []string  `json:"company_permissions"`
         
    }
    
    // ApplicationPermissions used by Partner
    type ApplicationPermissions struct {

        
            Permissions []string  `json:"permissions"`
            Roles []string  `json:"roles"`
         
    }
    
    // PartnerRequestList used by Partner
    type PartnerRequestList struct {

        
            Items PartnerList  `json:"items"`
            Page Pagination  `json:"page"`
         
    }
    
    // PartnerList used by Partner
    type PartnerList struct {

        
            AccountType string  `json:"account_type"`
            ApprovedPermissions ApprovedPermissionsInfo  `json:"approved_permissions"`
            ApproverID string  `json:"approver_id"`
            Comment string  `json:"comment"`
            CompanyID float64  `json:"company_id"`
            CompanyName string  `json:"company_name"`
            CreatedAt string  `json:"created_at"`
            ModifiedAt string  `json:"modified_at"`
            OrganizationID string  `json:"organization_id"`
            OrganizationName string  `json:"organization_name"`
            RequestStatus string  `json:"request_status"`
            RequestedPermissions RequestedPermissions  `json:"requested_permissions"`
            UserID string  `json:"user_id"`
            ID string  `json:"_id"`
         
    }
    
    // SetupProductRes used by Partner
    type SetupProductRes struct {

        
            Message string  `json:"message"`
            RequestID string  `json:"request_id"`
            NextStep float64  `json:"next_step"`
            CliWaitTime float64  `json:"cli_wait_time"`
         
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
    
    // getProxyPathRes used by Partner
    type getProxyPathRes struct {

        
            Page Pagination  `json:"page"`
            Items []AddProxyResponse  `json:"items"`
         
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
    

    
    // ServiceabilityrErrorResponse used by Serviceability
    type ServiceabilityrErrorResponse struct {

        
            Message string  `json:"message"`
            Value string  `json:"value"`
            Type string  `json:"type"`
         
    }
    
    // ApplicationServiceabilityConfig used by Serviceability
    type ApplicationServiceabilityConfig struct {

        
            ServiceabilityType string  `json:"serviceability_type"`
            ChannelID string  `json:"channel_id"`
            ChannelType string  `json:"channel_type"`
         
    }
    
    // ApplicationServiceabilityConfigResponse used by Serviceability
    type ApplicationServiceabilityConfigResponse struct {

        
            Error ServiceabilityrErrorResponse  `json:"error"`
            Data ApplicationServiceabilityConfig  `json:"data"`
            Success bool  `json:"success"`
         
    }
    
    // EntityRegionView_Request used by Serviceability
    type EntityRegionView_Request struct {

        
            ParentID []string  `json:"parent_id"`
            SubType []string  `json:"sub_type"`
         
    }
    
    // EntityRegionView_Items used by Serviceability
    type EntityRegionView_Items struct {

        
            UID string  `json:"uid"`
            SubType string  `json:"sub_type"`
            Name string  `json:"name"`
         
    }
    
    // EntityRegionView_page used by Serviceability
    type EntityRegionView_page struct {

        
            Size float64  `json:"size"`
            Current float64  `json:"current"`
            HasNext bool  `json:"has_next"`
            ItemTotal float64  `json:"item_total"`
            Type string  `json:"type"`
         
    }
    
    // EntityRegionView_Error used by Serviceability
    type EntityRegionView_Error struct {

        
            Message string  `json:"message"`
            Value string  `json:"value"`
            Type string  `json:"type"`
         
    }
    
    // EntityRegionView_Response used by Serviceability
    type EntityRegionView_Response struct {

        
            Success bool  `json:"success"`
            Data []EntityRegionView_Items  `json:"data"`
            Page EntityRegionView_page  `json:"page"`
            Error EntityRegionView_Error  `json:"error"`
         
    }
    
    // ListViewSummary used by Serviceability
    type ListViewSummary struct {

        
            TotalActiveZones float64  `json:"total_active_zones"`
            TotalZones float64  `json:"total_zones"`
            TotalPincodesServed float64  `json:"total_pincodes_served"`
         
    }
    
    // ZoneDataItem used by Serviceability
    type ZoneDataItem struct {

        
            Size float64  `json:"size"`
            Type string  `json:"type"`
            Current float64  `json:"current"`
            HasNext bool  `json:"has_next"`
            ItemTotal float64  `json:"item_total"`
         
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

        
            Name string  `json:"name"`
            StoresCount float64  `json:"stores_count"`
            Channels ListViewChannels  `json:"channels"`
            Slug string  `json:"slug"`
            ZoneID string  `json:"zone_id"`
            IsActive bool  `json:"is_active"`
            CompanyID float64  `json:"company_id"`
            Product ListViewProduct  `json:"product"`
            PincodesCount float64  `json:"pincodes_count"`
         
    }
    
    // ListViewResponse used by Serviceability
    type ListViewResponse struct {

        
            Summary []ListViewSummary  `json:"summary"`
            Page []ZoneDataItem  `json:"page"`
            Items []ListViewItems  `json:"items"`
         
    }
    
    // ServiceabilityErrorResponse used by Serviceability
    type ServiceabilityErrorResponse struct {

        
            Message string  `json:"message"`
            Value string  `json:"value"`
            Type string  `json:"type"`
         
    }
    
    // CompanyStoreView_PageItems used by Serviceability
    type CompanyStoreView_PageItems struct {

        
            Size float64  `json:"size"`
            Current float64  `json:"current"`
            HasNext bool  `json:"has_next"`
            ItemTotal float64  `json:"item_total"`
            Type string  `json:"type"`
         
    }
    
    // CompanyStoreView_Response used by Serviceability
    type CompanyStoreView_Response struct {

        
            Page []CompanyStoreView_PageItems  `json:"page"`
            Items []map[string]interface{}  `json:"items"`
         
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

        
            State []string  `json:"state"`
            Pincode []string  `json:"pincode"`
            Country string  `json:"country"`
         
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

        
            Data UpdateZoneData  `json:"data"`
            Identifier string  `json:"identifier"`
         
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

        
            Data CreateZoneData  `json:"data"`
            Identifier string  `json:"identifier"`
         
    }
    
    // ZoneResponse used by Serviceability
    type ZoneResponse struct {

        
            StatusCode float64  `json:"status_code"`
            ZoneID string  `json:"zone_id"`
            Success bool  `json:"success"`
         
    }
    
    // GetZoneFromApplicationIdViewResponse used by Serviceability
    type GetZoneFromApplicationIdViewResponse struct {

        
            Page []ZoneDataItem  `json:"page"`
            Items []ListViewItems  `json:"items"`
         
    }
    
    // GetZoneFromPincodeViewRequest used by Serviceability
    type GetZoneFromPincodeViewRequest struct {

        
            Pincode string  `json:"pincode"`
            Country string  `json:"country"`
         
    }
    
    // GetZoneFromPincodeViewResponse used by Serviceability
    type GetZoneFromPincodeViewResponse struct {

        
            ServiceabilityType string  `json:"serviceability_type"`
            Zones []string  `json:"zones"`
         
    }
    
    // ServiceabilityPageResponse used by Serviceability
    type ServiceabilityPageResponse struct {

        
            Size float64  `json:"size"`
            Current float64  `json:"current"`
            HasNext bool  `json:"has_next"`
            ItemTotal float64  `json:"item_total"`
            Type string  `json:"type"`
         
    }
    
    // AddressResponse used by Serviceability
    type AddressResponse struct {

        
            Latitude float64  `json:"latitude"`
            Landmark string  `json:"landmark"`
            Address2 string  `json:"address2"`
            State string  `json:"state"`
            Pincode float64  `json:"pincode"`
            Address1 string  `json:"address1"`
            City string  `json:"city"`
            Longitude float64  `json:"longitude"`
            Country string  `json:"country"`
         
    }
    
    // Dp used by Serviceability
    type Dp struct {

        
            TransportMode string  `json:"transport_mode"`
            ExternalAccountID string  `json:"external_account_id"`
            AreaCode float64  `json:"area_code"`
            RvpPriority float64  `json:"rvp_priority"`
            Operations []string  `json:"operations"`
            FmPriority float64  `json:"fm_priority"`
            InternalAccountID string  `json:"internal_account_id"`
            LmPriority float64  `json:"lm_priority"`
            PaymentMode string  `json:"payment_mode"`
            AssignDpFromSb bool  `json:"assign_dp_from_sb"`
         
    }
    
    // LogisticsResponse used by Serviceability
    type LogisticsResponse struct {

        
            Override bool  `json:"override"`
            Dp Dp  `json:"dp"`
         
    }
    
    // CreatedByResponse used by Serviceability
    type CreatedByResponse struct {

        
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
         
    }
    
    // DocumentsResponse used by Serviceability
    type DocumentsResponse struct {

        
            LegalName string  `json:"legal_name"`
            Verified bool  `json:"verified"`
            Value string  `json:"value"`
            Type string  `json:"type"`
         
    }
    
    // ModifiedByResponse used by Serviceability
    type ModifiedByResponse struct {

        
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
         
    }
    
    // MobileNo used by Serviceability
    type MobileNo struct {

        
            Number string  `json:"number"`
            CountryCode float64  `json:"country_code"`
         
    }
    
    // ManagerResponse used by Serviceability
    type ManagerResponse struct {

        
            Email string  `json:"email"`
            Name string  `json:"name"`
            MobileNo MobileNo  `json:"mobile_no"`
         
    }
    
    // WarningsResponse used by Serviceability
    type WarningsResponse struct {

        
            StoreAddress string  `json:"store_address"`
         
    }
    
    // IntegrationTypeResponse used by Serviceability
    type IntegrationTypeResponse struct {

        
            Order string  `json:"order"`
            Inventory string  `json:"inventory"`
         
    }
    
    // ContactNumberResponse used by Serviceability
    type ContactNumberResponse struct {

        
            Number string  `json:"number"`
            CountryCode float64  `json:"country_code"`
         
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
            Weekday string  `json:"weekday"`
            Open bool  `json:"open"`
            Closing OpeningClosing  `json:"closing"`
         
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
    
    // ItemResponse used by Serviceability
    type ItemResponse struct {

        
            UID float64  `json:"uid"`
            DisplayName string  `json:"display_name"`
            Cls string  `json:"_cls"`
            Address AddressResponse  `json:"address"`
            Logistics LogisticsResponse  `json:"logistics"`
            NotificationEmails []string  `json:"notification_emails"`
            CreatedBy CreatedByResponse  `json:"created_by"`
            Documents []DocumentsResponse  `json:"documents"`
            StoreType string  `json:"store_type"`
            VerifiedBy ModifiedByResponse  `json:"verified_by"`
            Manager ManagerResponse  `json:"manager"`
            CompanyID float64  `json:"company_id"`
            ModifiedOn string  `json:"modified_on"`
            Warnings WarningsResponse  `json:"warnings"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            IntegrationType IntegrationTypeResponse  `json:"integration_type"`
            Stage string  `json:"stage"`
            ModifiedBy ModifiedByResponse  `json:"modified_by"`
            ContactNumbers []ContactNumberResponse  `json:"contact_numbers"`
            VerifiedOn string  `json:"verified_on"`
            Company float64  `json:"company"`
            Name string  `json:"name"`
            CreatedOn string  `json:"created_on"`
            ProductReturnConfig ProductReturnConfigResponse  `json:"product_return_config"`
            Timing []TimmingResponse  `json:"timing"`
            GstCredentials GstCredentialsResponse  `json:"gst_credentials"`
            Code string  `json:"code"`
            SubType string  `json:"sub_type"`
         
    }
    
    // GetStoresViewResponse used by Serviceability
    type GetStoresViewResponse struct {

        
            Page ServiceabilityPageResponse  `json:"page"`
            Items []ItemResponse  `json:"items"`
         
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

        
            Size float64  `json:"size"`
            Current float64  `json:"current"`
            HasNext bool  `json:"has_next"`
            ItemTotal float64  `json:"item_total"`
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
    
