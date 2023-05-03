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

        
            Aggregators []map[string]interface{}  `json:"aggregators"`
            ExcludedFields []string  `json:"excluded_fields"`
            AppID string  `json:"app_id"`
            Success bool  `json:"success"`
            Created bool  `json:"created"`
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

        
            MerchantSalt string  `json:"merchant_salt"`
            Key string  `json:"key"`
            ConfigType string  `json:"config_type"`
            Secret string  `json:"secret"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // PaymentGatewayConfigRequest used by Payment
    type PaymentGatewayConfigRequest struct {

        
            AppID string  `json:"app_id"`
            AggregatorName PaymentGatewayConfig  `json:"aggregator_name"`
            IsActive bool  `json:"is_active"`
         
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
    
    // IntentAppErrorList used by Payment
    type IntentAppErrorList struct {

        
            Code string  `json:"code"`
            PackageName string  `json:"package_name"`
         
    }
    
    // PaymentModeList used by Payment
    type PaymentModeList struct {

        
            CardBrand string  `json:"card_brand"`
            CodLimitPerOrder float64  `json:"cod_limit_per_order"`
            ExpYear float64  `json:"exp_year"`
            ExpMonth float64  `json:"exp_month"`
            DisplayName string  `json:"display_name"`
            DisplayPriority float64  `json:"display_priority"`
            RemainingLimit float64  `json:"remaining_limit"`
            CardType string  `json:"card_type"`
            CardIsin string  `json:"card_isin"`
            CompliantWithTokenisationGuidelines bool  `json:"compliant_with_tokenisation_guidelines"`
            CardToken string  `json:"card_token"`
            IntentApp []IntentApp  `json:"intent_app"`
            CardIssuer string  `json:"card_issuer"`
            AggregatorName string  `json:"aggregator_name"`
            IntentAppErrorDictList []IntentAppErrorList  `json:"intent_app_error_dict_list"`
            Expired bool  `json:"expired"`
            CodLimit float64  `json:"cod_limit"`
            CardBrandImage string  `json:"card_brand_image"`
            CardName string  `json:"card_name"`
            CardReference string  `json:"card_reference"`
            CardFingerprint string  `json:"card_fingerprint"`
            IntentFlow bool  `json:"intent_flow"`
            MerchantCode string  `json:"merchant_code"`
            IntentAppErrorList []string  `json:"intent_app_error_list"`
            CardID string  `json:"card_id"`
            Nickname string  `json:"nickname"`
            FyndVpa string  `json:"fynd_vpa"`
            Code string  `json:"code"`
            CardNumber string  `json:"card_number"`
            Name string  `json:"name"`
            LogoURL PaymentModeLogo  `json:"logo_url"`
            RetryCount float64  `json:"retry_count"`
            Timeout float64  `json:"timeout"`
         
    }
    
    // RootPaymentMode used by Payment
    type RootPaymentMode struct {

        
            AnonymousEnable bool  `json:"anonymous_enable"`
            DisplayName string  `json:"display_name"`
            SaveCard bool  `json:"save_card"`
            DisplayPriority float64  `json:"display_priority"`
            Name string  `json:"name"`
            AddCardEnabled bool  `json:"add_card_enabled"`
            AggregatorName string  `json:"aggregator_name"`
            List []PaymentModeList  `json:"list"`
            IsPayByCardPl bool  `json:"is_pay_by_card_pl"`
         
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
    
    // PayoutCustomer used by Payment
    type PayoutCustomer struct {

        
            Name string  `json:"name"`
            UniqueExternalID string  `json:"unique_external_id"`
            Mobile string  `json:"mobile"`
            ID float64  `json:"id"`
         
    }
    
    // PayoutMoreAttributes used by Payment
    type PayoutMoreAttributes struct {

        
            State string  `json:"state"`
            BankName string  `json:"bank_name"`
            City string  `json:"city"`
            AccountNo string  `json:"account_no"`
            BranchName string  `json:"branch_name"`
            AccountHolder string  `json:"account_holder"`
            IfscCode string  `json:"ifsc_code"`
            AccountType string  `json:"account_type"`
            Country string  `json:"country"`
         
    }
    
    // PayoutAggregator used by Payment
    type PayoutAggregator struct {

        
            AggregatorFundID float64  `json:"aggregator_fund_id"`
            AggregatorID float64  `json:"aggregator_id"`
            PayoutDetailsID float64  `json:"payout_details_id"`
         
    }
    
    // Payout used by Payment
    type Payout struct {

        
            IsDefault bool  `json:"is_default"`
            Customers PayoutCustomer  `json:"customers"`
            MoreAttributes PayoutMoreAttributes  `json:"more_attributes"`
            PayoutsAggregators []PayoutAggregator  `json:"payouts_aggregators"`
            UniqueTransferNo string  `json:"unique_transfer_no"`
            TransferType string  `json:"transfer_type"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // PayoutsResponse used by Payment
    type PayoutsResponse struct {

        
            Items []Payout  `json:"items"`
            Success bool  `json:"success"`
         
    }
    
    // PayoutBankDetails used by Payment
    type PayoutBankDetails struct {

        
            State string  `json:"state"`
            AccountHolder string  `json:"account_holder"`
            BankName string  `json:"bank_name"`
            AccountNo string  `json:"account_no"`
            BranchName string  `json:"branch_name"`
            City string  `json:"city"`
            IfscCode string  `json:"ifsc_code"`
            AccountType string  `json:"account_type"`
            Pincode float64  `json:"pincode"`
            Country string  `json:"country"`
         
    }
    
    // PayoutRequest used by Payment
    type PayoutRequest struct {

        
            Users map[string]interface{}  `json:"users"`
            BankDetails PayoutBankDetails  `json:"bank_details"`
            TransferType string  `json:"transfer_type"`
            Aggregator string  `json:"aggregator"`
            UniqueExternalID string  `json:"unique_external_id"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // PayoutResponse used by Payment
    type PayoutResponse struct {

        
            PaymentStatus string  `json:"payment_status"`
            Users map[string]interface{}  `json:"users"`
            UniqueTransferNo string  `json:"unique_transfer_no"`
            BankDetails map[string]interface{}  `json:"bank_details"`
            Success bool  `json:"success"`
            TransferType string  `json:"transfer_type"`
            Created bool  `json:"created"`
            Aggregator string  `json:"aggregator"`
            Payouts map[string]interface{}  `json:"payouts"`
            IsActive bool  `json:"is_active"`
         
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
            AccountHolder string  `json:"account_holder"`
            BranchName string  `json:"branch_name"`
            AccountNo string  `json:"account_no"`
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

        
            ID float64  `json:"id"`
            DisplayName string  `json:"display_name"`
            AccountNo string  `json:"account_no"`
            DelightsUserName string  `json:"delights_user_name"`
            BankName string  `json:"bank_name"`
            TransferMode string  `json:"transfer_mode"`
            Title string  `json:"title"`
            AccountHolder string  `json:"account_holder"`
            BranchName string  `json:"branch_name"`
            Address string  `json:"address"`
            IfscCode string  `json:"ifsc_code"`
            ModifiedOn string  `json:"modified_on"`
            Subtitle string  `json:"subtitle"`
            Mobile string  `json:"mobile"`
            IsActive bool  `json:"is_active"`
            BeneficiaryID string  `json:"beneficiary_id"`
            CreatedOn string  `json:"created_on"`
            Email string  `json:"email"`
            Comment string  `json:"comment"`
         
    }
    
    // OrderBeneficiaryResponse used by Payment
    type OrderBeneficiaryResponse struct {

        
            ShowBeneficiaryDetails bool  `json:"show_beneficiary_details"`
            Beneficiaries []OrderBeneficiaryDetails  `json:"beneficiaries"`
         
    }
    
    // MultiTenderPaymentMeta used by Payment
    type MultiTenderPaymentMeta struct {

        
            OrderID string  `json:"order_id"`
            PaymentGateway string  `json:"payment_gateway"`
            PaymentID string  `json:"payment_id"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            CurrentStatus string  `json:"current_status"`
         
    }
    
    // MultiTenderPaymentMethod used by Payment
    type MultiTenderPaymentMethod struct {

        
            Name string  `json:"name"`
            Meta MultiTenderPaymentMeta  `json:"meta"`
            Amount float64  `json:"amount"`
            Mode string  `json:"mode"`
         
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
            CallbackURL map[string]interface{}  `json:"callback_url"`
            Source string  `json:"source"`
            Methods map[string]interface{}  `json:"methods"`
            Enabled bool  `json:"enabled"`
            PaymentSelectionLock map[string]interface{}  `json:"payment_selection_lock"`
            CodCharges float64  `json:"cod_charges"`
            AnonymousCod bool  `json:"anonymous_cod"`
            CodAmountLimit float64  `json:"cod_amount_limit"`
         
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
            PaymentSelectionLock map[string]interface{}  `json:"payment_selection_lock"`
            CodCharges float64  `json:"cod_charges"`
            AnonymousCod bool  `json:"anonymous_cod"`
            CodAmountLimit float64  `json:"cod_amount_limit"`
         
    }
    
    // CODdata used by Payment
    type CODdata struct {

        
            UserID string  `json:"user_id"`
            Usages float64  `json:"usages"`
            Limit float64  `json:"limit"`
            RemainingLimit float64  `json:"remaining_limit"`
            IsActive bool  `json:"is_active"`
         
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

        
            Models []string  `json:"models"`
            AggregatorID float64  `json:"aggregator_id"`
            Aggregator string  `json:"aggregator"`
         
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

        
            Success bool  `json:"success"`
            Statistics StatisticsData  `json:"statistics"`
         
    }
    
    // EdcAddRequest used by Payment
    type EdcAddRequest struct {

        
            EdcModel string  `json:"edc_model"`
            DeviceTag string  `json:"device_tag"`
            AggregatorID float64  `json:"aggregator_id"`
            StoreID float64  `json:"store_id"`
            EdcDeviceSerialNo string  `json:"edc_device_serial_no"`
            TerminalSerialNo string  `json:"terminal_serial_no"`
         
    }
    
    // EdcDevice used by Payment
    type EdcDevice struct {

        
            EdcModel string  `json:"edc_model"`
            DeviceTag string  `json:"device_tag"`
            AggregatorID float64  `json:"aggregator_id"`
            StoreID float64  `json:"store_id"`
            ApplicationID string  `json:"application_id"`
            AggregatorName string  `json:"aggregator_name"`
            MerchantStorePosCode string  `json:"merchant_store_pos_code"`
            TerminalUniqueIdentifier string  `json:"terminal_unique_identifier"`
            EdcDeviceSerialNo string  `json:"edc_device_serial_no"`
            TerminalSerialNo string  `json:"terminal_serial_no"`
            IsActive bool  `json:"is_active"`
         
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
            DeviceTag string  `json:"device_tag"`
            AggregatorID float64  `json:"aggregator_id"`
            StoreID float64  `json:"store_id"`
            MerchantStorePosCode string  `json:"merchant_store_pos_code"`
            EdcDeviceSerialNo string  `json:"edc_device_serial_no"`
            IsActive bool  `json:"is_active"`
         
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

        
            MerchantOrderID string  `json:"merchant_order_id"`
            RazorpayPaymentID string  `json:"razorpay_payment_id"`
            DeviceID string  `json:"device_id"`
            CustomerID string  `json:"customer_id"`
            Email string  `json:"email"`
            Amount float64  `json:"amount"`
            Method string  `json:"method"`
            Vpa string  `json:"vpa"`
            Aggregator string  `json:"aggregator"`
            Timeout float64  `json:"timeout"`
            OrderID string  `json:"order_id"`
            Currency string  `json:"currency"`
            Contact string  `json:"contact"`
         
    }
    
    // PaymentInitializationResponse used by Payment
    type PaymentInitializationResponse struct {

        
            DeviceID string  `json:"device_id"`
            Success bool  `json:"success"`
            OrderID string  `json:"order_id"`
            Currency string  `json:"currency"`
            UpiPollURL string  `json:"upi_poll_url"`
            PaymentID string  `json:"payment_id"`
            Method string  `json:"method"`
            PollingURL string  `json:"polling_url"`
            RazorpayPaymentID string  `json:"razorpay_payment_id"`
            CustomerID string  `json:"customer_id"`
            AggregatorOrderID string  `json:"aggregator_order_id"`
            Aggregator string  `json:"aggregator"`
            MerchantOrderID string  `json:"merchant_order_id"`
            Status string  `json:"status"`
            VirtualID string  `json:"virtual_id"`
            BqrImage string  `json:"bqr_image"`
            Email string  `json:"email"`
            Amount float64  `json:"amount"`
            Vpa string  `json:"vpa"`
            Timeout float64  `json:"timeout"`
            Contact string  `json:"contact"`
         
    }
    
    // PaymentStatusUpdateRequest used by Payment
    type PaymentStatusUpdateRequest struct {

        
            MerchantOrderID string  `json:"merchant_order_id"`
            DeviceID string  `json:"device_id"`
            RazorpayPaymentID string  `json:"razorpay_payment_id"`
            Status string  `json:"status"`
            PaymentID string  `json:"payment_id"`
            CustomerID string  `json:"customer_id"`
            Email string  `json:"email"`
            Amount float64  `json:"amount"`
            Method string  `json:"method"`
            Vpa string  `json:"vpa"`
            Aggregator string  `json:"aggregator"`
            OrderID string  `json:"order_id"`
            Currency string  `json:"currency"`
            Contact string  `json:"contact"`
         
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
            OrderID string  `json:"order_id"`
            DeviceID string  `json:"device_id"`
         
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
    

    
    // FilterInfoOption used by Order
    type FilterInfoOption struct {

        
            Name string  `json:"name"`
            Text string  `json:"text"`
            Value string  `json:"value"`
         
    }
    
    // FiltersInfo used by Order
    type FiltersInfo struct {

        
            Options []FilterInfoOption  `json:"options"`
            Type string  `json:"type"`
            Text string  `json:"text"`
            Value string  `json:"value"`
         
    }
    
    // GSTDetailsData used by Order
    type GSTDetailsData struct {

        
            GstinCode string  `json:"gstin_code"`
            ValueOfGood float64  `json:"value_of_good"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            GstFee float64  `json:"gst_fee"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
         
    }
    
    // PlatformItem used by Order
    type PlatformItem struct {

        
            L3Category float64  `json:"l3_category"`
            Code string  `json:"code"`
            ID float64  `json:"id"`
            Color string  `json:"color"`
            L1Category []string  `json:"l1_category"`
            DepartmentID float64  `json:"department_id"`
            L3CategoryName string  `json:"l3_category_name"`
            CanCancel bool  `json:"can_cancel"`
            CanReturn bool  `json:"can_return"`
            Size string  `json:"size"`
            Image []string  `json:"image"`
            Name string  `json:"name"`
            Images []string  `json:"images"`
         
    }
    
    // Prices used by Order
    type Prices struct {

        
            TransferPrice float64  `json:"transfer_price"`
            CouponValue float64  `json:"coupon_value"`
            ValueOfGood float64  `json:"value_of_good"`
            PromotionEffectiveDiscount float64  `json:"promotion_effective_discount"`
            FyndCredits float64  `json:"fynd_credits"`
            CodCharges float64  `json:"cod_charges"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            AmountPaidRoundoff float64  `json:"amount_paid_roundoff"`
            CashbackApplied float64  `json:"cashback_applied"`
            RefundAmount float64  `json:"refund_amount"`
            DeliveryCharge float64  `json:"delivery_charge"`
            PriceMarked float64  `json:"price_marked"`
            Discount float64  `json:"discount"`
            PriceEffective float64  `json:"price_effective"`
            Cashback float64  `json:"cashback"`
            AmountPaid float64  `json:"amount_paid"`
            RefundCredit float64  `json:"refund_credit"`
         
    }
    
    // BagUnit used by Order
    type BagUnit struct {

        
            OrderingChannel string  `json:"ordering_channel"`
            Gst GSTDetailsData  `json:"gst"`
            Status map[string]interface{}  `json:"status"`
            CanReturn bool  `json:"can_return"`
            CanCancel bool  `json:"can_cancel"`
            ItemQuantity float64  `json:"item_quantity"`
            Item PlatformItem  `json:"item"`
            TotalShipmentBags float64  `json:"total_shipment_bags"`
            BagID float64  `json:"bag_id"`
            Prices Prices  `json:"prices"`
            ShipmentID string  `json:"shipment_id"`
         
    }
    
    // ShipmentItemFulFillingStore used by Order
    type ShipmentItemFulFillingStore struct {

        
            Code string  `json:"code"`
            ID string  `json:"id"`
         
    }
    
    // ShipmentStatus used by Order
    type ShipmentStatus struct {

        
            HexCode string  `json:"hex_code"`
            Title string  `json:"title"`
            OpsStatus string  `json:"ops_status"`
            Status string  `json:"status"`
            ActualStatus string  `json:"actual_status"`
         
    }
    
    // PaymentModeInfo used by Order
    type PaymentModeInfo struct {

        
            Logo string  `json:"logo"`
            Type string  `json:"type"`
         
    }
    
    // UserDataInfo used by Order
    type UserDataInfo struct {

        
            Gender string  `json:"gender"`
            Email string  `json:"email"`
            FirstName string  `json:"first_name"`
            AvisUserID string  `json:"avis_user_id"`
            Mobile string  `json:"mobile"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            IsAnonymousUser bool  `json:"is_anonymous_user"`
            LastName string  `json:"last_name"`
         
    }
    
    // ShipmentItem used by Order
    type ShipmentItem struct {

        
            TotalBagsCount float64  `json:"total_bags_count"`
            ID string  `json:"id"`
            FulfillingCentre string  `json:"fulfilling_centre"`
            ShipmentCreatedAt string  `json:"shipment_created_at"`
            Company map[string]interface{}  `json:"company"`
            Channel map[string]interface{}  `json:"channel"`
            CreatedAt string  `json:"created_at"`
            PaymentMethods map[string]interface{}  `json:"payment_methods"`
            Bags []BagUnit  `json:"bags"`
            Application map[string]interface{}  `json:"application"`
            FulfillingStore ShipmentItemFulFillingStore  `json:"fulfilling_store"`
            TotalShipmentsInOrder float64  `json:"total_shipments_in_order"`
            ShipmentStatus ShipmentStatus  `json:"shipment_status"`
            Sla map[string]interface{}  `json:"sla"`
            PaymentModeInfo PaymentModeInfo  `json:"payment_mode_info"`
            Prices Prices  `json:"prices"`
            ShipmentID string  `json:"shipment_id"`
            User UserDataInfo  `json:"user"`
         
    }
    
    // ShipmentInternalPlatformViewResponse used by Order
    type ShipmentInternalPlatformViewResponse struct {

        
            Page map[string]interface{}  `json:"page"`
            Filters []FiltersInfo  `json:"filters"`
            Items []ShipmentItem  `json:"items"`
            AppliedFilters map[string]interface{}  `json:"applied_filters"`
         
    }
    
    // Error used by Order
    type Error struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // UserDetailsData used by Order
    type UserDetailsData struct {

        
            Pincode string  `json:"pincode"`
            Country string  `json:"country"`
            Area string  `json:"area"`
            Email string  `json:"email"`
            Phone string  `json:"phone"`
            Address string  `json:"address"`
            AddressType string  `json:"address_type"`
            Name string  `json:"name"`
            Landmark string  `json:"landmark"`
            State string  `json:"state"`
            Address1 string  `json:"address1"`
            City string  `json:"city"`
         
    }
    
    // InvoiceInfo used by Order
    type InvoiceInfo struct {

        
            StoreInvoiceID string  `json:"store_invoice_id"`
            LabelURL string  `json:"label_url"`
            InvoiceURL string  `json:"invoice_url"`
            CreditNoteID string  `json:"credit_note_id"`
            UpdatedDate string  `json:"updated_date"`
         
    }
    
    // BagStateMapper used by Order
    type BagStateMapper struct {

        
            IsActive bool  `json:"is_active"`
            AppFacing bool  `json:"app_facing"`
            NotifyCustomer bool  `json:"notify_customer"`
            DisplayName string  `json:"display_name"`
            AppStateName string  `json:"app_state_name"`
            BsID float64  `json:"bs_id"`
            Name string  `json:"name"`
            StateType string  `json:"state_type"`
            JourneyType string  `json:"journey_type"`
            AppDisplayName string  `json:"app_display_name"`
         
    }
    
    // BagStatusHistory used by Order
    type BagStatusHistory struct {

        
            Forward bool  `json:"forward"`
            StoreID float64  `json:"store_id"`
            AppDisplayName string  `json:"app_display_name"`
            BshID float64  `json:"bsh_id"`
            BagStateMapper BagStateMapper  `json:"bag_state_mapper"`
            KafkaSync bool  `json:"kafka_sync"`
            UpdatedAt string  `json:"updated_at"`
            CreatedAt string  `json:"created_at"`
            DeliveryAwbNumber string  `json:"delivery_awb_number"`
            Status string  `json:"status"`
            DisplayName string  `json:"display_name"`
            StateID float64  `json:"state_id"`
            Reasons []map[string]interface{}  `json:"reasons"`
            BagID float64  `json:"bag_id"`
            DeliveryPartnerID float64  `json:"delivery_partner_id"`
            StateType string  `json:"state_type"`
            ShipmentID string  `json:"shipment_id"`
         
    }
    
    // CurrentStatus used by Order
    type CurrentStatus struct {

        
            StoreID float64  `json:"store_id"`
            BagStateMapper BagStateMapper  `json:"bag_state_mapper"`
            KafkaSync bool  `json:"kafka_sync"`
            UpdatedAt string  `json:"updated_at"`
            CreatedAt string  `json:"created_at"`
            DeliveryAwbNumber string  `json:"delivery_awb_number"`
            Status string  `json:"status"`
            StateID float64  `json:"state_id"`
            CurrentStatusID float64  `json:"current_status_id"`
            BagID float64  `json:"bag_id"`
            DeliveryPartnerID float64  `json:"delivery_partner_id"`
            StateType string  `json:"state_type"`
            ShipmentID string  `json:"shipment_id"`
         
    }
    
    // OrderBagArticle used by Order
    type OrderBagArticle struct {

        
            Identifiers map[string]interface{}  `json:"identifiers"`
            ReturnConfig map[string]interface{}  `json:"return_config"`
            UID string  `json:"uid"`
         
    }
    
    // BagConfigs used by Order
    type BagConfigs struct {

        
            IsActive bool  `json:"is_active"`
            IsReturnable bool  `json:"is_returnable"`
            EnableTracking bool  `json:"enable_tracking"`
            CanBeCancelled bool  `json:"can_be_cancelled"`
            AllowForceReturn bool  `json:"allow_force_return"`
            IsCustomerReturnAllowed bool  `json:"is_customer_return_allowed"`
         
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

        
            Amount float64  `json:"amount"`
            DiscountRules []DiscountRules  `json:"discount_rules"`
            MrpPromotion bool  `json:"mrp_promotion"`
            PromotionType string  `json:"promotion_type"`
            PromotionName string  `json:"promotion_name"`
            ArticleQuantity float64  `json:"article_quantity"`
            PromoID string  `json:"promo_id"`
            BuyRules []BuyRules  `json:"buy_rules"`
         
    }
    
    // GiftCard used by Order
    type GiftCard struct {

        
            DisplayText string  `json:"display_text"`
            GiftMessage string  `json:"gift_message"`
            IsGiftApplied bool  `json:"is_gift_applied"`
            GiftPrice float64  `json:"gift_price"`
         
    }
    
    // B2BPODetails used by Order
    type B2BPODetails struct {

        
            TotalGstPercentage float64  `json:"total_gst_percentage"`
            DockerNumber string  `json:"docker_number"`
            PartialCanRet bool  `json:"partial_can_ret"`
            PoTaxAmount float64  `json:"po_tax_amount"`
            PoLineAmount float64  `json:"po_line_amount"`
            ItemBasePrice float64  `json:"item_base_price"`
         
    }
    
    // BagMeta used by Order
    type BagMeta struct {

        
            DocketNumber string  `json:"docket_number"`
            GroupID string  `json:"group_id"`
            CustomMessage string  `json:"custom_message"`
            PartialCanRet bool  `json:"partial_can_ret"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            GiftCard GiftCard  `json:"gift_card"`
            B2bPoDetails B2BPODetails  `json:"b2b_po_details"`
            CustomJson map[string]interface{}  `json:"custom_json"`
         
    }
    
    // PlatformDeliveryAddress used by Order
    type PlatformDeliveryAddress struct {

        
            ContactPerson string  `json:"contact_person"`
            Address2 string  `json:"address2"`
            Pincode string  `json:"pincode"`
            Country string  `json:"country"`
            Area string  `json:"area"`
            Email string  `json:"email"`
            Latitude float64  `json:"latitude"`
            CreatedAt string  `json:"created_at"`
            Phone string  `json:"phone"`
            UpdatedAt string  `json:"updated_at"`
            AddressType string  `json:"address_type"`
            AddressCategory string  `json:"address_category"`
            Landmark string  `json:"landmark"`
            Version string  `json:"version"`
            State string  `json:"state"`
            Address1 string  `json:"address1"`
            City string  `json:"city"`
            Longitude float64  `json:"longitude"`
         
    }
    
    // BagGST used by Order
    type BagGST struct {

        
            GstinCode string  `json:"gstin_code"`
            ValueOfGood float64  `json:"value_of_good"`
            IsDefaultHsnCode bool  `json:"is_default_hsn_code"`
            GstTag string  `json:"gst_tag"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            GstFee float64  `json:"gst_fee"`
            HsnCode string  `json:"hsn_code"`
         
    }
    
    // Identifier used by Order
    type Identifier struct {

        
            Ean string  `json:"ean"`
         
    }
    
    // FinancialBreakup used by Order
    type FinancialBreakup struct {

        
            ValueOfGood float64  `json:"value_of_good"`
            FyndCredits float64  `json:"fynd_credits"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            DeliveryCharge float64  `json:"delivery_charge"`
            TransferPrice float64  `json:"transfer_price"`
            Discount float64  `json:"discount"`
            TotalUnits float64  `json:"total_units"`
            CouponValue float64  `json:"coupon_value"`
            Identifiers Identifier  `json:"identifiers"`
            CashbackApplied float64  `json:"cashback_applied"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            PriceEffective float64  `json:"price_effective"`
            PriceMarked float64  `json:"price_marked"`
            AmountPaid float64  `json:"amount_paid"`
            RefundCredit float64  `json:"refund_credit"`
            PromotionEffectiveDiscount float64  `json:"promotion_effective_discount"`
            GstFee float64  `json:"gst_fee"`
            AddedToFyndCash bool  `json:"added_to_fynd_cash"`
            CodCharges float64  `json:"cod_charges"`
            AmountPaidRoundoff float64  `json:"amount_paid_roundoff"`
            Size string  `json:"size"`
            ItemName string  `json:"item_name"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            GstTag string  `json:"gst_tag"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            HsnCode string  `json:"hsn_code"`
            Cashback float64  `json:"cashback"`
         
    }
    
    // OrderBrandName used by Order
    type OrderBrandName struct {

        
            ID float64  `json:"id"`
            Company string  `json:"company"`
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
            Logo string  `json:"logo"`
            BrandName string  `json:"brand_name"`
         
    }
    
    // OrderBags used by Order
    type OrderBags struct {

        
            Item PlatformItem  `json:"item"`
            CurrentStatus CurrentStatus  `json:"current_status"`
            Article OrderBagArticle  `json:"article"`
            BagID float64  `json:"bag_id"`
            Quantity float64  `json:"quantity"`
            Prices Prices  `json:"prices"`
            BagConfigs BagConfigs  `json:"bag_configs"`
            AppliedPromos []AppliedPromos  `json:"applied_promos"`
            Identifier string  `json:"identifier"`
            Meta BagMeta  `json:"meta"`
            EntityType string  `json:"entity_type"`
            DisplayName string  `json:"display_name"`
            CanReturn bool  `json:"can_return"`
            LineNumber float64  `json:"line_number"`
            DeliveryAddress PlatformDeliveryAddress  `json:"delivery_address"`
            SellerIdentifier string  `json:"seller_identifier"`
            GstDetails BagGST  `json:"gst_details"`
            FinancialBreakup FinancialBreakup  `json:"financial_breakup"`
            CanCancel bool  `json:"can_cancel"`
            ParentPromoBags map[string]interface{}  `json:"parent_promo_bags"`
            GroupID string  `json:"group_id"`
            Brand OrderBrandName  `json:"brand"`
         
    }
    
    // DPDetailsData used by Order
    type DPDetailsData struct {

        
            ID float64  `json:"id"`
            EwayBillID string  `json:"eway_bill_id"`
            Pincode string  `json:"pincode"`
            Country string  `json:"country"`
            GstTag string  `json:"gst_tag"`
            TrackURL string  `json:"track_url"`
            AwbNo string  `json:"awb_no"`
            Name string  `json:"name"`
         
    }
    
    // PhoneDetails used by Order
    type PhoneDetails struct {

        
            MobileNumber string  `json:"mobile_number"`
            CountryCode float64  `json:"country_code"`
         
    }
    
    // ContactDetails used by Order
    type ContactDetails struct {

        
            Emails []string  `json:"emails"`
            Phone []PhoneDetails  `json:"phone"`
         
    }
    
    // CompanyDetails used by Order
    type CompanyDetails struct {

        
            CompanyCin string  `json:"company_cin"`
            CompanyContact ContactDetails  `json:"company_contact"`
            Address map[string]interface{}  `json:"address"`
            CompanyGst string  `json:"company_gst"`
            CompanyID float64  `json:"company_id"`
            CompanyName string  `json:"company_name"`
         
    }
    
    // OrderDetailsData used by Order
    type OrderDetailsData struct {

        
            OrderingChannel string  `json:"ordering_channel"`
            Source string  `json:"source"`
            FyndOrderID string  `json:"fynd_order_id"`
            AffiliateID string  `json:"affiliate_id"`
            CodCharges string  `json:"cod_charges"`
            TaxDetails map[string]interface{}  `json:"tax_details"`
            OrderDate string  `json:"order_date"`
            OrderValue string  `json:"order_value"`
            OrderingChannelLogo map[string]interface{}  `json:"ordering_channel_logo"`
         
    }
    
    // Dimensions used by Order
    type Dimensions struct {

        
            Unit string  `json:"unit"`
            Height float64  `json:"height"`
            IsDefault bool  `json:"is_default"`
            Length float64  `json:"length"`
            Width float64  `json:"width"`
         
    }
    
    // Meta used by Order
    type Meta struct {

        
            Dimension Dimensions  `json:"dimension"`
         
    }
    
    // TrackingList used by Order
    type TrackingList struct {

        
            Text string  `json:"text"`
            Status string  `json:"status"`
            Time string  `json:"time"`
            IsCurrent bool  `json:"is_current"`
            IsPassed bool  `json:"is_passed"`
         
    }
    
    // ShipmentStatusData used by Order
    type ShipmentStatusData struct {

        
            ID float64  `json:"id"`
            CreatedAt string  `json:"created_at"`
            Status string  `json:"status"`
            BagList []string  `json:"bag_list"`
            ShipmentID string  `json:"shipment_id"`
         
    }
    
    // LockData used by Order
    type LockData struct {

        
            Locked bool  `json:"locked"`
            LockMessage string  `json:"lock_message"`
            Mto bool  `json:"mto"`
         
    }
    
    // Formatted used by Order
    type Formatted struct {

        
            FMin string  `json:"f_min"`
            FMax string  `json:"f_max"`
         
    }
    
    // EinvoiceInfo used by Order
    type EinvoiceInfo struct {

        
            CreditNote map[string]interface{}  `json:"credit_note"`
            Invoice map[string]interface{}  `json:"invoice"`
         
    }
    
    // DebugInfo used by Order
    type DebugInfo struct {

        
            StormbreakerUUID string  `json:"stormbreaker_uuid"`
         
    }
    
    // BuyerDetails used by Order
    type BuyerDetails struct {

        
            Gstin string  `json:"gstin"`
            Pincode float64  `json:"pincode"`
            AjioSiteID string  `json:"ajio_site_id"`
            Address string  `json:"address"`
            Name string  `json:"name"`
            State string  `json:"state"`
            City string  `json:"city"`
         
    }
    
    // ShipmentTimeStamp used by Order
    type ShipmentTimeStamp struct {

        
            TMax string  `json:"t_max"`
            TMin string  `json:"t_min"`
         
    }
    
    // ShipmentMeta used by Order
    type ShipmentMeta struct {

        
            DueDate string  `json:"due_date"`
            PoNumber string  `json:"po_number"`
            DpName string  `json:"dp_name"`
            ReturnAffiliateOrderID string  `json:"return_affiliate_order_id"`
            External map[string]interface{}  `json:"external"`
            PackagingName string  `json:"packaging_name"`
            DpOptions map[string]interface{}  `json:"dp_options"`
            ReturnDetails map[string]interface{}  `json:"return_details"`
            MarketplaceStoreID string  `json:"marketplace_store_id"`
            AssignDpFromSb bool  `json:"assign_dp_from_sb"`
            LockData LockData  `json:"lock_data"`
            AwbNumber string  `json:"awb_number"`
            ForwardAffiliateOrderID string  `json:"forward_affiliate_order_id"`
            Formatted Formatted  `json:"formatted"`
            DpID string  `json:"dp_id"`
            OrderType string  `json:"order_type"`
            EinvoiceInfo EinvoiceInfo  `json:"einvoice_info"`
            EwaybillInfo map[string]interface{}  `json:"ewaybill_info"`
            B2cBuyerDetails map[string]interface{}  `json:"b2c_buyer_details"`
            ReturnAwbNumber string  `json:"return_awb_number"`
            ShipmentVolumetricWeight float64  `json:"shipment_volumetric_weight"`
            StoreInvoiceUpdatedDate string  `json:"store_invoice_updated_date"`
            DebugInfo DebugInfo  `json:"debug_info"`
            ForwardAffiliateShipmentID string  `json:"forward_affiliate_shipment_id"`
            BoxType string  `json:"box_type"`
            ReturnAffiliateShipmentID string  `json:"return_affiliate_shipment_id"`
            BagWeight map[string]interface{}  `json:"bag_weight"`
            FulfilmentPriorityText string  `json:"fulfilment_priority_text"`
            ShipmentWeight float64  `json:"shipment_weight"`
            DpSortKey string  `json:"dp_sort_key"`
            ReturnStoreNode float64  `json:"return_store_node"`
            AutoTriggerDpAssignmentAcf bool  `json:"auto_trigger_dp_assignment_acf"`
            Weight float64  `json:"weight"`
            B2bBuyerDetails BuyerDetails  `json:"b2b_buyer_details"`
            SameStoreAvailable bool  `json:"same_store_available"`
            Timestamp ShipmentTimeStamp  `json:"timestamp"`
         
    }
    
    // PDFLinks used by Order
    type PDFLinks struct {

        
            InvoiceA4 string  `json:"invoice_a4"`
            B2b string  `json:"b2b"`
            Invoice string  `json:"invoice"`
            LabelA4 string  `json:"label_a4"`
            InvoiceA6 string  `json:"invoice_a6"`
            LabelPos string  `json:"label_pos"`
            Label string  `json:"label"`
            InvoiceType string  `json:"invoice_type"`
            DeliveryChallanA4 string  `json:"delivery_challan_a4"`
            CreditNoteURL string  `json:"credit_note_url"`
            InvoicePos string  `json:"invoice_pos"`
            LabelA6 string  `json:"label_a6"`
            PoInvoice string  `json:"po_invoice"`
            LabelType string  `json:"label_type"`
         
    }
    
    // AffiliateMeta used by Order
    type AffiliateMeta struct {

        
            DueDate string  `json:"due_date"`
            ChannelShipmentID string  `json:"channel_shipment_id"`
            CouponCode string  `json:"coupon_code"`
            SizeLevelTotalQty float64  `json:"size_level_total_qty"`
            IsPriority bool  `json:"is_priority"`
            BoxType string  `json:"box_type"`
            OrderItemID string  `json:"order_item_id"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            EmployeeDiscount float64  `json:"employee_discount"`
            Quantity float64  `json:"quantity"`
            ChannelOrderID string  `json:"channel_order_id"`
            MarketplaceInvoiceID string  `json:"marketplace_invoice_id"`
         
    }
    
    // AffiliateDetails used by Order
    type AffiliateDetails struct {

        
            CompanyAffiliateTag string  `json:"company_affiliate_tag"`
            AffiliateID string  `json:"affiliate_id"`
            AdID string  `json:"ad_id"`
            ShipmentMeta ShipmentMeta  `json:"shipment_meta"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
            PdfLinks PDFLinks  `json:"pdf_links"`
            AffiliateMeta AffiliateMeta  `json:"affiliate_meta"`
            AffiliateStoreID string  `json:"affiliate_store_id"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
         
    }
    
    // OrderingStoreDetails used by Order
    type OrderingStoreDetails struct {

        
            OrderingStoreID float64  `json:"ordering_store_id"`
            Code string  `json:"code"`
            ContactPerson string  `json:"contact_person"`
            Pincode string  `json:"pincode"`
            Country string  `json:"country"`
            Meta map[string]interface{}  `json:"meta"`
            Phone string  `json:"phone"`
            Address string  `json:"address"`
            State string  `json:"state"`
            City string  `json:"city"`
            StoreName string  `json:"store_name"`
         
    }
    
    // ShipmentPayments used by Order
    type ShipmentPayments struct {

        
            Mode string  `json:"mode"`
            Source string  `json:"source"`
            Logo string  `json:"logo"`
         
    }
    
    // FulfillingStore used by Order
    type FulfillingStore struct {

        
            Code string  `json:"code"`
            ContactPerson string  `json:"contact_person"`
            ID float64  `json:"id"`
            FulfillmentChannel string  `json:"fulfillment_channel"`
            Pincode string  `json:"pincode"`
            Country string  `json:"country"`
            Meta map[string]interface{}  `json:"meta"`
            Phone string  `json:"phone"`
            Address string  `json:"address"`
            State string  `json:"state"`
            City string  `json:"city"`
            StoreName string  `json:"store_name"`
         
    }
    
    // PlatformShipment used by Order
    type PlatformShipment struct {

        
            DeliveryDetails UserDetailsData  `json:"delivery_details"`
            Invoice InvoiceInfo  `json:"invoice"`
            TotalBags float64  `json:"total_bags"`
            BagStatusHistory []BagStatusHistory  `json:"bag_status_history"`
            Bags []OrderBags  `json:"bags"`
            DpDetails DPDetailsData  `json:"dp_details"`
            Vertical string  `json:"vertical"`
            FulfilmentPriority float64  `json:"fulfilment_priority"`
            PlatformLogo string  `json:"platform_logo"`
            EnableDpTracking bool  `json:"enable_dp_tracking"`
            InvoiceID string  `json:"invoice_id"`
            Prices Prices  `json:"prices"`
            CustomMeta []map[string]interface{}  `json:"custom_meta"`
            Coupon map[string]interface{}  `json:"coupon"`
            DeliverySlot map[string]interface{}  `json:"delivery_slot"`
            CompanyDetails CompanyDetails  `json:"company_details"`
            ShipmentStatus string  `json:"shipment_status"`
            LockStatus bool  `json:"lock_status"`
            ForwardShipmentID string  `json:"forward_shipment_id"`
            Order OrderDetailsData  `json:"order"`
            OperationalStatus string  `json:"operational_status"`
            Meta Meta  `json:"meta"`
            ShipmentImages []string  `json:"shipment_images"`
            TrackingList []TrackingList  `json:"tracking_list"`
            PaymentMethods map[string]interface{}  `json:"payment_methods"`
            Status ShipmentStatusData  `json:"status"`
            PaymentMode string  `json:"payment_mode"`
            TotalItems float64  `json:"total_items"`
            PickedDate string  `json:"picked_date"`
            AffiliateDetails AffiliateDetails  `json:"affiliate_details"`
            GstDetails GSTDetailsData  `json:"gst_details"`
            OrderingStore OrderingStoreDetails  `json:"ordering_store"`
            JourneyType string  `json:"journey_type"`
            PriorityText string  `json:"priority_text"`
            ShipmentQuantity float64  `json:"shipment_quantity"`
            UserAgent string  `json:"user_agent"`
            Payments ShipmentPayments  `json:"payments"`
            FulfillingStore FulfillingStore  `json:"fulfilling_store"`
            PackagingType string  `json:"packaging_type"`
            BillingDetails UserDetailsData  `json:"billing_details"`
            ShipmentID string  `json:"shipment_id"`
            User UserDataInfo  `json:"user"`
         
    }
    
    // ShipmentInfoResponse used by Order
    type ShipmentInfoResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
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
            Success bool  `json:"success"`
            Result []AssetByShipment  `json:"result"`
         
    }
    
    // BillingStaffDetails used by Order
    type BillingStaffDetails struct {

        
            FirstName string  `json:"first_name"`
            EmployeeCode string  `json:"employee_code"`
            StaffID float64  `json:"staff_id"`
            LastName string  `json:"last_name"`
            User string  `json:"user"`
         
    }
    
    // PlatformUserDetails used by Order
    type PlatformUserDetails struct {

        
            PlatformUserFirstName string  `json:"platform_user_first_name"`
            PlatformUserLastName string  `json:"platform_user_last_name"`
            PlatformUserID string  `json:"platform_user_id"`
            PlatformUserEmployeeCode string  `json:"platform_user_employee_code"`
         
    }
    
    // TransactionData used by Order
    type TransactionData struct {

        
            TransactionID string  `json:"transaction_id"`
            UniqueReferenceNumber string  `json:"unique_reference_number"`
            Status string  `json:"status"`
            Entity string  `json:"entity"`
            Currency string  `json:"currency"`
            TerminalID string  `json:"terminal_id"`
            PaymentID string  `json:"payment_id"`
            AmountPaid string  `json:"amount_paid"`
         
    }
    
    // OrderMeta used by Order
    type OrderMeta struct {

        
            CurrencySymbol string  `json:"currency_symbol"`
            Staff map[string]interface{}  `json:"staff"`
            OrderType string  `json:"order_type"`
            OrderChildEntities []string  `json:"order_child_entities"`
            MongoCartID float64  `json:"mongo_cart_id"`
            OrderTags []map[string]interface{}  `json:"order_tags"`
            Comment string  `json:"comment"`
            CustomerNote string  `json:"customer_note"`
            BillingStaffDetails BillingStaffDetails  `json:"billing_staff_details"`
            OrderPlatform string  `json:"order_platform"`
            CompanyLogo string  `json:"company_logo"`
            OrderingStore float64  `json:"ordering_store"`
            PlatformUserDetails PlatformUserDetails  `json:"platform_user_details"`
            TransactionData TransactionData  `json:"transaction_data"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            EmployeeID float64  `json:"employee_id"`
            Files []map[string]interface{}  `json:"files"`
            PaymentType string  `json:"payment_type"`
            CartID float64  `json:"cart_id"`
         
    }
    
    // TaxDetails used by Order
    type TaxDetails struct {

        
            Gstin string  `json:"gstin"`
            PanNo string  `json:"pan_no"`
         
    }
    
    // OrderDict used by Order
    type OrderDict struct {

        
            FyndOrderID string  `json:"fynd_order_id"`
            Meta OrderMeta  `json:"meta"`
            PaymentMethods map[string]interface{}  `json:"payment_methods"`
            TaxDetails TaxDetails  `json:"tax_details"`
            OrderDate string  `json:"order_date"`
            Prices Prices  `json:"prices"`
         
    }
    
    // ShipmentDetailsResponse used by Order
    type ShipmentDetailsResponse struct {

        
            Shipments []PlatformShipment  `json:"shipments"`
            Success bool  `json:"success"`
            Order OrderDict  `json:"order"`
         
    }
    
    // SubLane used by Order
    type SubLane struct {

        
            Text string  `json:"text"`
            Index float64  `json:"index"`
            Value string  `json:"value"`
            TotalItems float64  `json:"total_items"`
            Actions []map[string]interface{}  `json:"actions"`
         
    }
    
    // SuperLane used by Order
    type SuperLane struct {

        
            Options []SubLane  `json:"options"`
            TotalItems float64  `json:"total_items"`
            Text string  `json:"text"`
            Value string  `json:"value"`
         
    }
    
    // LaneConfigResponse used by Order
    type LaneConfigResponse struct {

        
            SuperLanes []SuperLane  `json:"super_lanes"`
         
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
    
    // PlatformBreakupValues used by Order
    type PlatformBreakupValues struct {

        
            Name string  `json:"name"`
            Display string  `json:"display"`
            Value string  `json:"value"`
         
    }
    
    // PlatformChannel used by Order
    type PlatformChannel struct {

        
            Logo string  `json:"logo"`
            Name string  `json:"name"`
         
    }
    
    // PlatformOrderItems used by Order
    type PlatformOrderItems struct {

        
            BreakupValues []PlatformBreakupValues  `json:"breakup_values"`
            TotalOrderValue float64  `json:"total_order_value"`
            Meta map[string]interface{}  `json:"meta"`
            OrderID string  `json:"order_id"`
            Channel PlatformChannel  `json:"channel"`
            UserInfo UserDataInfo  `json:"user_info"`
            Shipments []PlatformShipment  `json:"shipments"`
            OrderCreatedTime string  `json:"order_created_time"`
            PaymentMode string  `json:"payment_mode"`
            OrderValue float64  `json:"order_value"`
         
    }
    
    // OrderListingResponse used by Order
    type OrderListingResponse struct {

        
            Lane string  `json:"lane"`
            Success bool  `json:"success"`
            Page Page  `json:"page"`
            Message string  `json:"message"`
            Items []PlatformOrderItems  `json:"items"`
            TotalCount float64  `json:"total_count"`
         
    }
    
    // Options used by Order
    type Options struct {

        
            Text string  `json:"text"`
            Value float64  `json:"value"`
         
    }
    
    // MetricsCount used by Order
    type MetricsCount struct {

        
            Key string  `json:"key"`
            Options []Options  `json:"options"`
            Value float64  `json:"value"`
            Text string  `json:"text"`
         
    }
    
    // MetricCountResponse used by Order
    type MetricCountResponse struct {

        
            Items []MetricsCount  `json:"items"`
         
    }
    
    // PlatformTrack used by Order
    type PlatformTrack struct {

        
            UpdatedTime string  `json:"updated_time"`
            ShipmentType string  `json:"shipment_type"`
            AccountName string  `json:"account_name"`
            Meta map[string]interface{}  `json:"meta"`
            UpdatedAt string  `json:"updated_at"`
            Status string  `json:"status"`
            RawStatus string  `json:"raw_status"`
            Awb string  `json:"awb"`
            Reason string  `json:"reason"`
            LastLocationRecievedAt string  `json:"last_location_recieved_at"`
         
    }
    
    // PlatformShipmentTrack used by Order
    type PlatformShipmentTrack struct {

        
            Meta map[string]interface{}  `json:"meta"`
            Results []PlatformTrack  `json:"results"`
         
    }
    
    // AdvanceFilterInfo used by Order
    type AdvanceFilterInfo struct {

        
            Returned []FiltersInfo  `json:"returned"`
            Unfulfilled []FiltersInfo  `json:"unfulfilled"`
            Processed []FiltersInfo  `json:"processed"`
            ActionCentre []FiltersInfo  `json:"action_centre"`
            Filters []FiltersInfo  `json:"filters"`
         
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

        
            ReportID string  `json:"report_id"`
            ReportType string  `json:"report_type"`
            DisplayName string  `json:"display_name"`
            Status string  `json:"status"`
            ReportCreatedAt string  `json:"report_created_at"`
            RequestDetails map[string]interface{}  `json:"request_details"`
            ReportRequestedAt string  `json:"report_requested_at"`
            S3Key string  `json:"s3_key"`
            ReportName string  `json:"report_name"`
         
    }
    
    // JioCodeUpsertDataSet used by Order
    type JioCodeUpsertDataSet struct {

        
            JioCode string  `json:"jio_code"`
            ItemID string  `json:"item_id"`
            CompanyID string  `json:"company_id"`
            ArticleID string  `json:"article_id"`
         
    }
    
    // JioCodeUpsertPayload used by Order
    type JioCodeUpsertPayload struct {

        
            Data []JioCodeUpsertDataSet  `json:"data"`
         
    }
    
    // NestedErrorSchemaDataSet used by Order
    type NestedErrorSchemaDataSet struct {

        
            Message string  `json:"message"`
            Value string  `json:"value"`
            Type string  `json:"type"`
         
    }
    
    // JioCodeUpsertResponse used by Order
    type JioCodeUpsertResponse struct {

        
            Success bool  `json:"success"`
            Data []map[string]interface{}  `json:"data"`
            Identifier string  `json:"identifier"`
            Error []NestedErrorSchemaDataSet  `json:"error"`
            TraceID string  `json:"trace_id"`
         
    }
    
    // BulkInvoicingResponse used by Order
    type BulkInvoicingResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // BulkInvoiceLabelResponse used by Order
    type BulkInvoiceLabelResponse struct {

        
            InvoiceStatus string  `json:"invoice_status"`
            StoreID string  `json:"store_id"`
            Data map[string]interface{}  `json:"data"`
            DoInvoiceLabelGenerated bool  `json:"do_invoice_label_generated"`
            Invoice map[string]interface{}  `json:"invoice"`
            StoreCode string  `json:"store_code"`
            Label map[string]interface{}  `json:"label"`
            BatchID string  `json:"batch_id"`
            CompanyID string  `json:"company_id"`
            StoreName string  `json:"store_name"`
         
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
            ContentType string  `json:"content_type"`
            Tags []string  `json:"tags"`
            Namespace string  `json:"namespace"`
            Upload FileUploadResponse  `json:"upload"`
            Cdn URL  `json:"cdn"`
            Size float64  `json:"size"`
            FilePath string  `json:"file_path"`
            Method string  `json:"method"`
            Operation string  `json:"operation"`
         
    }
    
    // bulkListingData used by Order
    type bulkListingData struct {

        
            ExcelURL string  `json:"excel_url"`
            StoreID float64  `json:"store_id"`
            SuccessfulShipments []map[string]interface{}  `json:"successful_shipments"`
            FailedShipments []map[string]interface{}  `json:"failed_shipments"`
            ProcessingShipments []string  `json:"processing_shipments"`
            Processing float64  `json:"processing"`
            FileName string  `json:"file_name"`
            Successful float64  `json:"successful"`
            CompanyID float64  `json:"company_id"`
            StoreCode string  `json:"store_code"`
            Status string  `json:"status"`
            Failed float64  `json:"failed"`
            ID string  `json:"id"`
            UserID string  `json:"user_id"`
            UploadedOn string  `json:"uploaded_on"`
            BatchID string  `json:"batch_id"`
            Total float64  `json:"total"`
            UserName string  `json:"user_name"`
            StoreName string  `json:"store_name"`
         
    }
    
    // BulkListingPage used by Order
    type BulkListingPage struct {

        
            Current float64  `json:"current"`
            HasPrevious bool  `json:"has_previous"`
            Size float64  `json:"size"`
            HasNext bool  `json:"has_next"`
            Total float64  `json:"total"`
            Type string  `json:"type"`
         
    }
    
    // BulkListingResponse used by Order
    type BulkListingResponse struct {

        
            Error string  `json:"error"`
            Success bool  `json:"success"`
            Data []bulkListingData  `json:"data"`
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
            QcType []string  `json:"qc_type"`
            QuestionSet []QuestionSet  `json:"question_set"`
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

        
            TotalShipmentsCount float64  `json:"total_shipments_count"`
            SuccessfulShipmentIds []string  `json:"successful_shipment_ids"`
            BatchID string  `json:"batch_id"`
            FailedShipmentsCount float64  `json:"failed_shipments_count"`
            SuccessfulShipmentsCount float64  `json:"successful_shipments_count"`
            ProcessingShipmentsCount float64  `json:"processing_shipments_count"`
            CompanyID string  `json:"company_id"`
         
    }
    
    // BulkActionDetailsResponse used by Order
    type BulkActionDetailsResponse struct {

        
            Success string  `json:"success"`
            Data []BulkActionDetailsDataField  `json:"data"`
            UserID string  `json:"user_id"`
            Error []string  `json:"error"`
            FailedRecords []string  `json:"failed_records"`
            Status bool  `json:"status"`
            UploadedOn string  `json:"uploaded_on"`
            UploadedBy string  `json:"uploaded_by"`
            Message string  `json:"message"`
         
    }
    
    // Attributes used by Order
    type Attributes struct {

        
            PrimaryColor string  `json:"primary_color"`
            MarketerName string  `json:"marketer_name"`
            Gender []string  `json:"gender"`
            Essential string  `json:"essential"`
            PrimaryMaterial string  `json:"primary_material"`
            BrandName string  `json:"brand_name"`
            PrimaryColorHex string  `json:"primary_color_hex"`
            MarketerAddress string  `json:"marketer_address"`
            Name string  `json:"name"`
         
    }
    
    // Item used by Order
    type Item struct {

        
            SlugKey string  `json:"slug_key"`
            BrandID float64  `json:"brand_id"`
            Name string  `json:"name"`
            LastUpdatedAt string  `json:"last_updated_at"`
            WebstoreProductURL string  `json:"webstore_product_url"`
            Image []string  `json:"image"`
            L2CategoryID float64  `json:"l2_category_id"`
            ItemID float64  `json:"item_id"`
            L2Category []string  `json:"l2_category"`
            L3Category float64  `json:"l3_category"`
            Code string  `json:"code"`
            Meta map[string]interface{}  `json:"meta"`
            L1Category []string  `json:"l1_category"`
            L3CategoryName string  `json:"l3_category_name"`
            CanReturn bool  `json:"can_return"`
            Color string  `json:"color"`
            Gender string  `json:"gender"`
            DepartmentID float64  `json:"department_id"`
            CanCancel bool  `json:"can_cancel"`
            Size string  `json:"size"`
            L1CategoryID float64  `json:"l1_category_id"`
            BranchURL string  `json:"branch_url"`
            Attributes Attributes  `json:"attributes"`
            Brand string  `json:"brand"`
         
    }
    
    // ReturnConfig used by Order
    type ReturnConfig struct {

        
            Returnable bool  `json:"returnable"`
            Unit string  `json:"unit"`
            Time float64  `json:"time"`
         
    }
    
    // Weight used by Order
    type Weight struct {

        
            Unit string  `json:"unit"`
            Shipping float64  `json:"shipping"`
            IsDefault bool  `json:"is_default"`
         
    }
    
    // Article used by Order
    type Article struct {

        
            Code string  `json:"code"`
            RawMeta interface{}  `json:"raw_meta"`
            Identifiers Identifier  `json:"identifiers"`
            ASet map[string]interface{}  `json:"a_set"`
            Size string  `json:"size"`
            SellerIdentifier string  `json:"seller_identifier"`
            ReturnConfig ReturnConfig  `json:"return_config"`
            EspModified interface{}  `json:"esp_modified"`
            UID string  `json:"uid"`
            IsSet bool  `json:"is_set"`
            ChildDetails map[string]interface{}  `json:"child_details"`
            Weight Weight  `json:"weight"`
            Dimensions Dimensions  `json:"dimensions"`
            ID string  `json:"_id"`
         
    }
    
    // AffiliateBagDetails used by Order
    type AffiliateBagDetails struct {

        
            AffiliateBagID string  `json:"affiliate_bag_id"`
            EmployeeDiscount float64  `json:"employee_discount"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            AffiliateMeta AffiliateMeta  `json:"affiliate_meta"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
         
    }
    
    // ArticleDetails used by Order
    type ArticleDetails struct {

        
            Status map[string]interface{}  `json:"status"`
         
    }
    
    // BagReturnableCancelableStatus used by Order
    type BagReturnableCancelableStatus struct {

        
            IsActive bool  `json:"is_active"`
            IsReturnable bool  `json:"is_returnable"`
            EnableTracking bool  `json:"enable_tracking"`
            CanBeCancelled bool  `json:"can_be_cancelled"`
            IsCustomerReturnAllowed bool  `json:"is_customer_return_allowed"`
         
    }
    
    // BagGSTDetails used by Order
    type BagGSTDetails struct {

        
            GstinCode string  `json:"gstin_code"`
            ValueOfGood float64  `json:"value_of_good"`
            IsDefaultHsnCode bool  `json:"is_default_hsn_code"`
            IgstGstFee string  `json:"igst_gst_fee"`
            SgstGstFee string  `json:"sgst_gst_fee"`
            GstTag string  `json:"gst_tag"`
            CgstTaxPercentage float64  `json:"cgst_tax_percentage"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            GstFee float64  `json:"gst_fee"`
            CgstGstFee string  `json:"cgst_gst_fee"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            HsnCode string  `json:"hsn_code"`
            SgstTaxPercentage float64  `json:"sgst_tax_percentage"`
            HsnCodeID string  `json:"hsn_code_id"`
            IgstTaxPercentage float64  `json:"igst_tax_percentage"`
         
    }
    
    // StoreAddress used by Order
    type StoreAddress struct {

        
            Pincode float64  `json:"pincode"`
            Country string  `json:"country"`
            Email string  `json:"email"`
            Area string  `json:"area"`
            Latitude float64  `json:"latitude"`
            Landmark string  `json:"landmark"`
            CountryCode string  `json:"country_code"`
            AddressCategory string  `json:"address_category"`
            AddressType string  `json:"address_type"`
            Phone string  `json:"phone"`
            Version string  `json:"version"`
            State string  `json:"state"`
            City string  `json:"city"`
            ContactPerson string  `json:"contact_person"`
            Address2 string  `json:"address2"`
            UpdatedAt string  `json:"updated_at"`
            CreatedAt string  `json:"created_at"`
            Address1 string  `json:"address1"`
            Longitude float64  `json:"longitude"`
         
    }
    
    // Document used by Order
    type Document struct {

        
            DsType string  `json:"ds_type"`
            Value string  `json:"value"`
            URL string  `json:"url"`
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
            Username string  `json:"username"`
            Enabled bool  `json:"enabled"`
            User string  `json:"user"`
         
    }
    
    // StoreGstCredentials used by Order
    type StoreGstCredentials struct {

        
            EWaybill StoreEwaybill  `json:"e_waybill"`
            EInvoice StoreEinvoice  `json:"e_invoice"`
         
    }
    
    // EInvoicePortalDetails used by Order
    type EInvoicePortalDetails struct {

        
            Password string  `json:"password"`
            Username string  `json:"username"`
            User string  `json:"user"`
         
    }
    
    // StoreMeta used by Order
    type StoreMeta struct {

        
            AdditionalContactDetails map[string]interface{}  `json:"additional_contact_details"`
            EwaybillPortalDetails map[string]interface{}  `json:"ewaybill_portal_details"`
            NotificationEmails []string  `json:"notification_emails"`
            DisplayName string  `json:"display_name"`
            Documents StoreDocuments  `json:"documents"`
            Stage string  `json:"stage"`
            Timing []map[string]interface{}  `json:"timing"`
            ProductReturnConfig map[string]interface{}  `json:"product_return_config"`
            GstCredentials StoreGstCredentials  `json:"gst_credentials"`
            GstNumber string  `json:"gst_number"`
            EinvoicePortalDetails EInvoicePortalDetails  `json:"einvoice_portal_details"`
         
    }
    
    // Store used by Order
    type Store struct {

        
            SID string  `json:"s_id"`
            Pincode string  `json:"pincode"`
            Country string  `json:"country"`
            BrandID interface{}  `json:"brand_id"`
            StoreAddressJson StoreAddress  `json:"store_address_json"`
            Name string  `json:"name"`
            Longitude float64  `json:"longitude"`
            StoreEmail string  `json:"store_email"`
            Latitude float64  `json:"latitude"`
            ParentStoreID float64  `json:"parent_store_id"`
            StoreActiveFrom string  `json:"store_active_from"`
            OrderIntegrationID string  `json:"order_integration_id"`
            MallName string  `json:"mall_name"`
            IsActive bool  `json:"is_active"`
            Code string  `json:"code"`
            VatNo string  `json:"vat_no"`
            Meta StoreMeta  `json:"meta"`
            IsArchived bool  `json:"is_archived"`
            Phone float64  `json:"phone"`
            State string  `json:"state"`
            AlohomoraUserID float64  `json:"alohomora_user_id"`
            City string  `json:"city"`
            IsEnabledForRecon bool  `json:"is_enabled_for_recon"`
            BrandStoreTags []string  `json:"brand_store_tags"`
            FulfillmentChannel string  `json:"fulfillment_channel"`
            ContactPerson string  `json:"contact_person"`
            Address2 string  `json:"address2"`
            PackagingMaterialCount float64  `json:"packaging_material_count"`
            UpdatedAt string  `json:"updated_at"`
            CreatedAt string  `json:"created_at"`
            Address1 string  `json:"address1"`
            LocationType string  `json:"location_type"`
            LoginUsername string  `json:"login_username"`
            MallArea string  `json:"mall_area"`
            CompanyID float64  `json:"company_id"`
         
    }
    
    // Dates used by Order
    type Dates struct {

        
            DeliveryDate interface{}  `json:"delivery_date"`
            OrderCreated string  `json:"order_created"`
         
    }
    
    // Brand used by Order
    type Brand struct {

        
            CreditNoteAllowed bool  `json:"credit_note_allowed"`
            StartDate string  `json:"start_date"`
            IsVirtualInvoice bool  `json:"is_virtual_invoice"`
            Company string  `json:"company"`
            ScriptLastRan string  `json:"script_last_ran"`
            CreatedOn float64  `json:"created_on"`
            ModifiedOn float64  `json:"modified_on"`
            PickupLocation string  `json:"pickup_location"`
            Logo string  `json:"logo"`
            CreditNoteExpiryDays float64  `json:"credit_note_expiry_days"`
            BrandID float64  `json:"brand_id"`
            BrandName string  `json:"brand_name"`
            InvoicePrefix string  `json:"invoice_prefix"`
         
    }
    
    // BagDetailsPlatformResponse used by Order
    type BagDetailsPlatformResponse struct {

        
            Tags []string  `json:"tags"`
            BagStatus []BagStatusHistory  `json:"bag_status"`
            BagStatusHistory BagStatusHistory  `json:"bag_status_history"`
            Item Item  `json:"item"`
            CurrentStatus BagStatusHistory  `json:"current_status"`
            Article Article  `json:"article"`
            Prices Prices  `json:"prices"`
            OriginalBagList []float64  `json:"original_bag_list"`
            Quantity float64  `json:"quantity"`
            QcRequired interface{}  `json:"qc_required"`
            RestoreCoupon bool  `json:"restore_coupon"`
            OrderIntegrationID string  `json:"order_integration_id"`
            AppliedPromos []map[string]interface{}  `json:"applied_promos"`
            BagUpdateTime float64  `json:"bag_update_time"`
            Meta BagMeta  `json:"meta"`
            Identifier string  `json:"identifier"`
            AffiliateBagDetails AffiliateBagDetails  `json:"affiliate_bag_details"`
            OperationalStatus string  `json:"operational_status"`
            ArticleDetails ArticleDetails  `json:"article_details"`
            EntityType string  `json:"entity_type"`
            DisplayName string  `json:"display_name"`
            Status BagReturnableCancelableStatus  `json:"status"`
            LineNumber float64  `json:"line_number"`
            AffiliateDetails AffiliateDetails  `json:"affiliate_details"`
            CurrentOperationalStatus BagStatusHistory  `json:"current_operational_status"`
            Reasons []map[string]interface{}  `json:"reasons"`
            NoOfBagsOrder float64  `json:"no_of_bags_order"`
            RestorePromos map[string]interface{}  `json:"restore_promos"`
            SellerIdentifier string  `json:"seller_identifier"`
            GstDetails BagGSTDetails  `json:"gst_details"`
            OrderingStore Store  `json:"ordering_store"`
            JourneyType string  `json:"journey_type"`
            FinancialBreakup []FinancialBreakup  `json:"financial_breakup"`
            Dates Dates  `json:"dates"`
            ParentPromoBags map[string]interface{}  `json:"parent_promo_bags"`
            BType string  `json:"b_type"`
            Brand Brand  `json:"brand"`
            ShipmentID string  `json:"shipment_id"`
            BID float64  `json:"b_id"`
         
    }
    
    // ErrorResponse used by Order
    type ErrorResponse struct {

        
            Error string  `json:"error"`
            Message string  `json:"message"`
         
    }
    
    // BagsPage used by Order
    type BagsPage struct {

        
            PageType string  `json:"page_type"`
            Current float64  `json:"current"`
            Size float64  `json:"size"`
            HasNext bool  `json:"has_next"`
            ItemTotal float64  `json:"item_total"`
         
    }
    
    // GetBagsPlatformResponse used by Order
    type GetBagsPlatformResponse struct {

        
            Page BagsPage  `json:"page"`
            Items []BagDetailsPlatformResponse  `json:"items"`
         
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

        
            ShipmentID string  `json:"shipment_id"`
            Error string  `json:"error"`
            Message string  `json:"message"`
            Status float64  `json:"status"`
         
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
         
    }
    
    // StoreReassign used by Order
    type StoreReassign struct {

        
            AffiliateBagID string  `json:"affiliate_bag_id"`
            StoreID float64  `json:"store_id"`
            MongoArticleID string  `json:"mongo_article_id"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            FyndOrderID string  `json:"fynd_order_id"`
            BagID float64  `json:"bag_id"`
            ItemID string  `json:"item_id"`
            ReasonIds []float64  `json:"reason_ids"`
            SetID string  `json:"set_id"`
            AffiliateID string  `json:"affiliate_id"`
         
    }
    
    // StoreReassignResponse used by Order
    type StoreReassignResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // Entities used by Order
    type Entities struct {

        
            AffiliateBagID string  `json:"affiliate_bag_id"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            ID string  `json:"id"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            ReasonText string  `json:"reason_text"`
            AffiliateID string  `json:"affiliate_id"`
         
    }
    
    // UpdateShipmentLockPayload used by Order
    type UpdateShipmentLockPayload struct {

        
            Entities []Entities  `json:"entities"`
            Action string  `json:"action"`
            EntityType string  `json:"entity_type"`
            ActionType string  `json:"action_type"`
         
    }
    
    // Bags used by Order
    type Bags struct {

        
            BagID float64  `json:"bag_id"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            IsLocked bool  `json:"is_locked"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
         
    }
    
    // OriginalFilter used by Order
    type OriginalFilter struct {

        
            AffiliateID string  `json:"affiliate_id"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
         
    }
    
    // CheckResponse used by Order
    type CheckResponse struct {

        
            IsBagLocked bool  `json:"is_bag_locked"`
            ShipmentID string  `json:"shipment_id"`
            Bags []Bags  `json:"bags"`
            OriginalFilter OriginalFilter  `json:"original_filter"`
            Status string  `json:"status"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            LockStatus bool  `json:"lock_status"`
            AffiliateID string  `json:"affiliate_id"`
            IsShipmentLocked bool  `json:"is_shipment_locked"`
         
    }
    
    // UpdateShipmentLockResponse used by Order
    type UpdateShipmentLockResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
            CheckResponse []CheckResponse  `json:"check_response"`
         
    }
    
    // AnnouncementResponse used by Order
    type AnnouncementResponse struct {

        
            CompanyID float64  `json:"company_id"`
            Description string  `json:"description"`
            PlatformID string  `json:"platform_id"`
            Title string  `json:"title"`
            FromDatetime string  `json:"from_datetime"`
            ID float64  `json:"id"`
            ToDatetime string  `json:"to_datetime"`
            CreatedAt string  `json:"created_at"`
            LogoURL string  `json:"logo_url"`
            PlatformName string  `json:"platform_name"`
         
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
    
    // ProductsReasonsData used by Order
    type ProductsReasonsData struct {

        
            ReasonID float64  `json:"reason_id"`
            ReasonText string  `json:"reason_text"`
         
    }
    
    // ProductsReasonsFilters used by Order
    type ProductsReasonsFilters struct {

        
            Quantity float64  `json:"quantity"`
            LineNumber float64  `json:"line_number"`
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

        
            Products []Products  `json:"products"`
            DataUpdates DataUpdates  `json:"data_updates"`
            Reasons ReasonsData  `json:"reasons"`
            Identifier string  `json:"identifier"`
         
    }
    
    // StatuesRequest used by Order
    type StatuesRequest struct {

        
            ExcludeBagsNextState string  `json:"exclude_bags_next_state"`
            Shipments []ShipmentsRequest  `json:"shipments"`
            Status string  `json:"status"`
         
    }
    
    // UpdateShipmentStatusRequest used by Order
    type UpdateShipmentStatusRequest struct {

        
            Task bool  `json:"task"`
            ForceTransition bool  `json:"force_transition"`
            LockAfterTransition bool  `json:"lock_after_transition"`
            UnlockBeforeTransition bool  `json:"unlock_before_transition"`
            Statuses []StatuesRequest  `json:"statuses"`
         
    }
    
    // ShipmentsResponse used by Order
    type ShipmentsResponse struct {

        
            Exception string  `json:"exception"`
            StackTrace string  `json:"stack_trace"`
            Code string  `json:"code"`
            FinalState map[string]interface{}  `json:"final_state"`
            Status float64  `json:"status"`
            Meta map[string]interface{}  `json:"meta"`
            Message string  `json:"message"`
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
    
    // AffiliateInventoryOrderConfig used by Order
    type AffiliateInventoryOrderConfig struct {

        
            ForceReassignment bool  `json:"force_reassignment"`
         
    }
    
    // AffiliateInventoryStoreConfig used by Order
    type AffiliateInventoryStoreConfig struct {

        
            Store map[string]interface{}  `json:"store"`
         
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
    
    // AffiliateInventoryConfig used by Order
    type AffiliateInventoryConfig struct {

        
            ArticleAssignment AffiliateInventoryArticleAssignmentConfig  `json:"article_assignment"`
            Order AffiliateInventoryOrderConfig  `json:"order"`
            Inventory AffiliateInventoryStoreConfig  `json:"inventory"`
            Payment AffiliateInventoryPaymentConfig  `json:"payment"`
            Logistics AffiliateInventoryLogisticsConfig  `json:"logistics"`
         
    }
    
    // AffiliateAppConfigMeta used by Order
    type AffiliateAppConfigMeta struct {

        
            Value string  `json:"value"`
            Name string  `json:"name"`
         
    }
    
    // AffiliateAppConfig used by Order
    type AffiliateAppConfig struct {

        
            Description string  `json:"description"`
            Owner string  `json:"owner"`
            Secret string  `json:"secret"`
            Meta []AffiliateAppConfigMeta  `json:"meta"`
            ID string  `json:"id"`
            CreatedAt string  `json:"created_at"`
            Token string  `json:"token"`
            UpdatedAt string  `json:"updated_at"`
            Name string  `json:"name"`
         
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
    
    // AffiliateStoreIdMapping used by Order
    type AffiliateStoreIdMapping struct {

        
            StoreID float64  `json:"store_id"`
            MarketplaceStoreID string  `json:"marketplace_store_id"`
         
    }
    
    // OrderConfig used by Order
    type OrderConfig struct {

        
            StoreLookup string  `json:"store_lookup"`
            Affiliate Affiliate  `json:"affiliate"`
            BagEndState string  `json:"bag_end_state"`
            AffiliateStoreIDMapping []AffiliateStoreIdMapping  `json:"affiliate_store_id_mapping"`
            ArticleLookup string  `json:"article_lookup"`
            CreateUser bool  `json:"create_user"`
         
    }
    
    // OrderPriority used by Order
    type OrderPriority struct {

        
            AffiliatePriorityCode string  `json:"affiliate_priority_code"`
            FulfilmentPriority float64  `json:"fulfilment_priority"`
            FulfilmentPriorityText string  `json:"fulfilment_priority_text"`
         
    }
    
    // OrderUser used by Order
    type OrderUser struct {

        
            LastName string  `json:"last_name"`
            Pincode string  `json:"pincode"`
            FirstName string  `json:"first_name"`
            Country string  `json:"country"`
            Address1 string  `json:"address1"`
            Email string  `json:"email"`
            Address2 string  `json:"address2"`
            State string  `json:"state"`
            City string  `json:"city"`
            Mobile float64  `json:"mobile"`
            Phone float64  `json:"phone"`
         
    }
    
    // UserData used by Order
    type UserData struct {

        
            ShippingUser OrderUser  `json:"shipping_user"`
            BillingUser OrderUser  `json:"billing_user"`
         
    }
    
    // MarketPlacePdf used by Order
    type MarketPlacePdf struct {

        
            Invoice string  `json:"invoice"`
            Label string  `json:"label"`
         
    }
    
    // AffiliateBag used by Order
    type AffiliateBag struct {

        
            StoreID float64  `json:"store_id"`
            DeliveryCharge float64  `json:"delivery_charge"`
            Sku string  `json:"sku"`
            AffiliateMeta map[string]interface{}  `json:"affiliate_meta"`
            Quantity float64  `json:"quantity"`
            FyndStoreID string  `json:"fynd_store_id"`
            ID string  `json:"_id"`
            SellerIdentifier string  `json:"seller_identifier"`
            UnitPrice float64  `json:"unit_price"`
            Identifier map[string]interface{}  `json:"identifier"`
            PdfLinks MarketPlacePdf  `json:"pdf_links"`
            CompanyID float64  `json:"company_id"`
            ItemSize string  `json:"item_size"`
            AffiliateStoreID string  `json:"affiliate_store_id"`
            ModifiedOn string  `json:"modified_on"`
            HsnCodeID string  `json:"hsn_code_id"`
            AmountPaid float64  `json:"amount_paid"`
            Discount float64  `json:"discount"`
            PriceEffective float64  `json:"price_effective"`
            ItemID float64  `json:"item_id"`
            TransferPrice float64  `json:"transfer_price"`
            PriceMarked float64  `json:"price_marked"`
            AvlQty float64  `json:"avl_qty"`
         
    }
    
    // ArticleDetails1 used by Order
    type ArticleDetails1 struct {

        
            Dimension map[string]interface{}  `json:"dimension"`
            Attributes map[string]interface{}  `json:"attributes"`
            Category map[string]interface{}  `json:"category"`
            Quantity float64  `json:"quantity"`
            ID string  `json:"_id"`
            Weight map[string]interface{}  `json:"weight"`
            BrandID float64  `json:"brand_id"`
         
    }
    
    // ShipmentDetails used by Order
    type ShipmentDetails struct {

        
            Shipments float64  `json:"shipments"`
            FulfillmentID float64  `json:"fulfillment_id"`
            DpID float64  `json:"dp_id"`
            Meta map[string]interface{}  `json:"meta"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            Articles []ArticleDetails1  `json:"articles"`
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

        
            ToPincode string  `json:"to_pincode"`
            Source string  `json:"source"`
            PaymentMode string  `json:"payment_mode"`
            Journey string  `json:"journey"`
            Shipment []ShipmentDetails  `json:"shipment"`
            Action string  `json:"action"`
            LocationDetails LocationDetails  `json:"location_details"`
            Identifier string  `json:"identifier"`
         
    }
    
    // ShipmentData used by Order
    type ShipmentData struct {

        
            ShipmentData ShipmentConfig  `json:"shipment_data"`
         
    }
    
    // OrderInfo used by Order
    type OrderInfo struct {

        
            Items map[string]interface{}  `json:"items"`
            DeliveryCharges float64  `json:"delivery_charges"`
            OrderPriority OrderPriority  `json:"order_priority"`
            User UserData  `json:"user"`
            Bags []AffiliateBag  `json:"bags"`
            PaymentMode string  `json:"payment_mode"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            OrderValue float64  `json:"order_value"`
            Discount float64  `json:"discount"`
            Coupon string  `json:"coupon"`
            CodCharges float64  `json:"cod_charges"`
            ShippingAddress OrderUser  `json:"shipping_address"`
            Shipment ShipmentData  `json:"shipment"`
            Payment map[string]interface{}  `json:"payment"`
            BillingAddress OrderUser  `json:"billing_address"`
         
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

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // ActionInfo used by Order
    type ActionInfo struct {

        
            Slug string  `json:"slug"`
            ID float64  `json:"id"`
            DisplayText string  `json:"display_text"`
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
            ShipmentID string  `json:"shipment_id"`
            Identifier string  `json:"identifier"`
         
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

        
            User string  `json:"user"`
            Createdat string  `json:"createdat"`
            TicketURL string  `json:"ticket_url"`
            L2Detail string  `json:"l2_detail"`
            BagID float64  `json:"bag_id"`
            L1Detail string  `json:"l1_detail"`
            TicketID string  `json:"ticket_id"`
            Message string  `json:"message"`
            Type string  `json:"type"`
            L3Detail string  `json:"l3_detail"`
         
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

        
            CountryCode string  `json:"country_code"`
            ShipmentID float64  `json:"shipment_id"`
            CustomerName string  `json:"customer_name"`
            AmountPaid float64  `json:"amount_paid"`
            BrandName string  `json:"brand_name"`
            PaymentMode string  `json:"payment_mode"`
            PhoneNumber float64  `json:"phone_number"`
            Message string  `json:"message"`
            OrderID string  `json:"order_id"`
         
    }
    
    // SendSmsPayload used by Order
    type SendSmsPayload struct {

        
            BagID float64  `json:"bag_id"`
            Data SmsDataPayload  `json:"data"`
            Slug string  `json:"slug"`
         
    }
    
    // Meta1 used by Order
    type Meta1 struct {

        
            KafkaEmissionStatus float64  `json:"kafka_emission_status"`
            StateManagerUsed string  `json:"state_manager_used"`
         
    }
    
    // ShipmentDetail used by Order
    type ShipmentDetail struct {

        
            Remarks string  `json:"remarks"`
            ShipmentID string  `json:"shipment_id"`
            BagList []float64  `json:"bag_list"`
            Status string  `json:"status"`
            Meta Meta1  `json:"meta"`
            ID float64  `json:"id"`
         
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

        
            Result []OrderStatusData  `json:"result"`
            Success string  `json:"success"`
         
    }
    
    // ManualAssignDPToShipment used by Order
    type ManualAssignDPToShipment struct {

        
            OrderType string  `json:"order_type"`
            ShipmentIds []string  `json:"shipment_ids"`
            QcRequired string  `json:"qc_required"`
            DpID float64  `json:"dp_id"`
         
    }
    
    // ManualAssignDPToShipmentResponse used by Order
    type ManualAssignDPToShipmentResponse struct {

        
            Success string  `json:"success"`
            Errors []string  `json:"errors"`
         
    }
    
    // ShippingInfo used by Order
    type ShippingInfo struct {

        
            FirstName string  `json:"first_name"`
            ExternalCustomerCode string  `json:"external_customer_code"`
            City string  `json:"city"`
            Landmark string  `json:"landmark"`
            CountryCode string  `json:"country_code"`
            Slot []map[string]interface{}  `json:"slot"`
            Address2 string  `json:"address2"`
            Title string  `json:"title"`
            MiddleName string  `json:"middle_name"`
            PrimaryMobileNumber string  `json:"primary_mobile_number"`
            AlternateEmail string  `json:"alternate_email"`
            Address1 string  `json:"address1"`
            PrimaryEmail string  `json:"primary_email"`
            State string  `json:"state"`
            AlternateMobileNumber string  `json:"alternate_mobile_number"`
            AddressType string  `json:"address_type"`
            GeoLocation map[string]interface{}  `json:"geo_location"`
            CustomerCode string  `json:"customer_code"`
            FloorNo string  `json:"floor_no"`
            LastName string  `json:"last_name"`
            Gender string  `json:"gender"`
            Pincode string  `json:"pincode"`
            ShippingType string  `json:"shipping_type"`
            Country string  `json:"country"`
            HouseNo string  `json:"house_no"`
            StateCode string  `json:"state_code"`
         
    }
    
    // ProcessingDates used by Order
    type ProcessingDates struct {

        
            CustomerPickupSlot map[string]interface{}  `json:"customer_pickup_slot"`
            DispatchAfterDate string  `json:"dispatch_after_date"`
            DpPickupSlot map[string]interface{}  `json:"dp_pickup_slot"`
            DispatchByDate string  `json:"dispatch_by_date"`
            PackByDate string  `json:"pack_by_date"`
            ConfirmByDate string  `json:"confirm_by_date"`
         
    }
    
    // Tax used by Order
    type Tax struct {

        
            Rate float64  `json:"rate"`
            Breakup []map[string]interface{}  `json:"breakup"`
            Name string  `json:"name"`
            Amount map[string]interface{}  `json:"amount"`
         
    }
    
    // Charge used by Order
    type Charge struct {

        
            Code string  `json:"code"`
            Tax Tax  `json:"tax"`
            Type string  `json:"type"`
            Name string  `json:"name"`
            Amount map[string]interface{}  `json:"amount"`
         
    }
    
    // LineItem used by Order
    type LineItem struct {

        
            CustomMessasge string  `json:"custom_messasge"`
            Quantity float64  `json:"quantity"`
            Meta map[string]interface{}  `json:"meta"`
            Charges []Charge  `json:"charges"`
            SellerIdentifier string  `json:"seller_identifier"`
            ExternalLineID string  `json:"external_line_id"`
         
    }
    
    // Shipment used by Order
    type Shipment struct {

        
            LocationID float64  `json:"location_id"`
            Meta map[string]interface{}  `json:"meta"`
            ExternalShipmentID string  `json:"external_shipment_id"`
            ProcessingDates ProcessingDates  `json:"processing_dates"`
            LineItems []LineItem  `json:"line_items"`
            Priority float64  `json:"priority"`
         
    }
    
    // PaymentMethod used by Order
    type PaymentMethod struct {

        
            Mode string  `json:"mode"`
            CollectBy string  `json:"collect_by"`
            Meta map[string]interface{}  `json:"meta"`
            RefundBy string  `json:"refund_by"`
            TransactionData map[string]interface{}  `json:"transaction_data"`
            Name string  `json:"name"`
            Amount float64  `json:"amount"`
         
    }
    
    // PaymentInfo used by Order
    type PaymentInfo struct {

        
            PrimaryMode string  `json:"primary_mode"`
            PaymentMethods []PaymentMethod  `json:"payment_methods"`
         
    }
    
    // BillingInfo used by Order
    type BillingInfo struct {

        
            FirstName string  `json:"first_name"`
            ExternalCustomerCode string  `json:"external_customer_code"`
            City string  `json:"city"`
            CountryCode string  `json:"country_code"`
            Address2 string  `json:"address2"`
            Title string  `json:"title"`
            MiddleName string  `json:"middle_name"`
            PrimaryMobileNumber string  `json:"primary_mobile_number"`
            AlternateEmail string  `json:"alternate_email"`
            Address1 string  `json:"address1"`
            PrimaryEmail string  `json:"primary_email"`
            State string  `json:"state"`
            AlternateMobileNumber string  `json:"alternate_mobile_number"`
            CustomerCode string  `json:"customer_code"`
            FloorNo string  `json:"floor_no"`
            LastName string  `json:"last_name"`
            Gender string  `json:"gender"`
            Pincode string  `json:"pincode"`
            Country string  `json:"country"`
            HouseNo string  `json:"house_no"`
            StateCode string  `json:"state_code"`
         
    }
    
    // TaxInfo used by Order
    type TaxInfo struct {

        
            Gstin string  `json:"gstin"`
            B2bGstinNumber string  `json:"b2b_gstin_number"`
         
    }
    
    // CreateOrderAPI used by Order
    type CreateOrderAPI struct {

        
            CurrencyInfo map[string]interface{}  `json:"currency_info"`
            ShippingInfo ShippingInfo  `json:"shipping_info"`
            Shipments []Shipment  `json:"shipments"`
            Meta map[string]interface{}  `json:"meta"`
            Charges []Charge  `json:"charges"`
            PaymentInfo PaymentInfo  `json:"payment_info"`
            Config map[string]interface{}  `json:"config"`
            ExternalOrderID string  `json:"external_order_id"`
            BillingInfo BillingInfo  `json:"billing_info"`
            TaxInfo TaxInfo  `json:"tax_info"`
            ExternalCreationDate string  `json:"external_creation_date"`
         
    }
    
    // CreateOrderErrorReponse used by Order
    type CreateOrderErrorReponse struct {

        
            Exception string  `json:"exception"`
            StackTrace string  `json:"stack_trace"`
            Code string  `json:"code"`
            Info interface{}  `json:"info"`
            RequestID string  `json:"request_id"`
            Status float64  `json:"status"`
            Meta string  `json:"meta"`
            Message string  `json:"message"`
         
    }
    
    // PaymentMethods used by Order
    type PaymentMethods struct {

        
            Mode string  `json:"mode"`
            CollectBy string  `json:"collect_by"`
            RefundBy string  `json:"refund_by"`
         
    }
    
    // CreateChannelPaymentInfo used by Order
    type CreateChannelPaymentInfo struct {

        
            Source string  `json:"source"`
            ModeOfPayment string  `json:"mode_of_payment"`
            PaymentMethods []PaymentMethods  `json:"payment_methods"`
         
    }
    
    // DpConfiguration used by Order
    type DpConfiguration struct {

        
            ShippingBy string  `json:"shipping_by"`
         
    }
    
    // CreateChannelConfig used by Order
    type CreateChannelConfig struct {

        
            LogoURL map[string]interface{}  `json:"logo_url"`
            PaymentInfo CreateChannelPaymentInfo  `json:"payment_info"`
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
            IsUpserted bool  `json:"is_upserted"`
            Acknowledged bool  `json:"acknowledged"`
         
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
            OrderDetails []FyndOrderIdList  `json:"order_details"`
            Mobile float64  `json:"mobile"`
         
    }
    
    // FetchCreditBalanceRequestPayload used by Order
    type FetchCreditBalanceRequestPayload struct {

        
            SellerID string  `json:"seller_id"`
            CustomerMobileNumber string  `json:"customer_mobile_number"`
            AffiliateID string  `json:"affiliate_id"`
         
    }
    
    // CreditBalanceInfo used by Order
    type CreditBalanceInfo struct {

        
            Reason string  `json:"reason"`
            CustomerMobileNumber string  `json:"customer_mobile_number"`
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
            OrderingChannel string  `json:"ordering_channel"`
            FyndOrderID string  `json:"fynd_order_id"`
            SellerID string  `json:"seller_id"`
            AffiliateID string  `json:"affiliate_id"`
         
    }
    
    // SingleRefundModeInfo used by Order
    type SingleRefundModeInfo struct {

        
            IsActive bool  `json:"is_active"`
            DisplayName string  `json:"display_name"`
            Slug string  `json:"slug"`
         
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
    
    // AttachUserOtpData used by Order
    type AttachUserOtpData struct {

        
            RequestID string  `json:"request_id"`
            OtpCode float64  `json:"otp_code"`
         
    }
    
    // AttachUserInfo used by Order
    type AttachUserInfo struct {

        
            LastName string  `json:"last_name"`
            CountryCode string  `json:"country_code"`
            FirstName string  `json:"first_name"`
            Mobile float64  `json:"mobile"`
         
    }
    
    // AttachOrderUser used by Order
    type AttachOrderUser struct {

        
            OtpData AttachUserOtpData  `json:"otp_data"`
            UserInfo AttachUserInfo  `json:"user_info"`
            FyndOrderID string  `json:"fynd_order_id"`
         
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

        
            RequestID string  `json:"request_id"`
            Mobile float64  `json:"mobile"`
            Message string  `json:"message"`
            ResendTimer float64  `json:"resend_timer"`
         
    }
    
    // SendUserMobileOtpResponse used by Order
    type SendUserMobileOtpResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
            Data PointBlankOtpData  `json:"data"`
         
    }
    

    
    // GetSearchWordsData used by Catalog
    type GetSearchWordsData struct {

        
            Result map[string]interface{}  `json:"result"`
            UID string  `json:"uid"`
            Words []string  `json:"words"`
            AppID string  `json:"app_id"`
            IsActive bool  `json:"is_active"`
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

        
            Status float64  `json:"status"`
            Message string  `json:"message"`
            Error string  `json:"error"`
            Code string  `json:"code"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // SearchKeywordResult used by Catalog
    type SearchKeywordResult struct {

        
            Query map[string]interface{}  `json:"query"`
            SortOn string  `json:"sort_on"`
         
    }
    
    // CreateSearchKeyword used by Catalog
    type CreateSearchKeyword struct {

        
            Result SearchKeywordResult  `json:"result"`
            Words []string  `json:"words"`
            AppID string  `json:"app_id"`
            IsActive bool  `json:"is_active"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // DeleteResponse used by Catalog
    type DeleteResponse struct {

        
            Message string  `json:"message"`
         
    }
    
    // GetSearchWordsResponse used by Catalog
    type GetSearchWordsResponse struct {

        
            Items []GetSearchWordsData  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // GetAutocompleteWordsData used by Catalog
    type GetAutocompleteWordsData struct {

        
            UID string  `json:"uid"`
            Words []string  `json:"words"`
            AppID string  `json:"app_id"`
            Results []map[string]interface{}  `json:"results"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // GetAutocompleteWordsResponse used by Catalog
    type GetAutocompleteWordsResponse struct {

        
            Items []GetAutocompleteWordsData  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // Media used by Catalog
    type Media struct {

        
            URL string  `json:"url"`
            AspectRatio string  `json:"aspect_ratio"`
            Type string  `json:"type"`
         
    }
    
    // AutocompletePageAction used by Catalog
    type AutocompletePageAction struct {

        
            Params map[string]interface{}  `json:"params"`
            Query map[string]interface{}  `json:"query"`
            URL string  `json:"url"`
            Type string  `json:"type"`
         
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
            AppID string  `json:"app_id"`
            Results []AutocompleteResult  `json:"results"`
            IsActive bool  `json:"is_active"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
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

        
            MinQuantity float64  `json:"min_quantity"`
            MaxQuantity float64  `json:"max_quantity"`
            ProductUID float64  `json:"product_uid"`
            AllowRemove bool  `json:"allow_remove"`
            AutoSelect bool  `json:"auto_select"`
            AutoAddToCart bool  `json:"auto_add_to_cart"`
         
    }
    
    // GetProductBundleCreateResponse used by Catalog
    type GetProductBundleCreateResponse struct {

        
            Slug string  `json:"slug"`
            Choice string  `json:"choice"`
            CreatedOn string  `json:"created_on"`
            Name string  `json:"name"`
            PageVisibility []string  `json:"page_visibility"`
            Logo string  `json:"logo"`
            Products []ProductBundleItem  `json:"products"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            CompanyID float64  `json:"company_id"`
            IsActive bool  `json:"is_active"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ID string  `json:"id"`
            ModifiedOn string  `json:"modified_on"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // GetProductBundleListingResponse used by Catalog
    type GetProductBundleListingResponse struct {

        
            Items []GetProductBundleCreateResponse  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // ProductBundleRequest used by Catalog
    type ProductBundleRequest struct {

        
            Slug string  `json:"slug"`
            Choice string  `json:"choice"`
            CreatedOn string  `json:"created_on"`
            Name string  `json:"name"`
            PageVisibility []string  `json:"page_visibility"`
            Logo string  `json:"logo"`
            Products []ProductBundleItem  `json:"products"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            CompanyID float64  `json:"company_id"`
            IsActive bool  `json:"is_active"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ModifiedOn string  `json:"modified_on"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // Size used by Catalog
    type Size struct {

        
            Value string  `json:"value"`
            Quantity float64  `json:"quantity"`
            IsAvailable bool  `json:"is_available"`
            Display string  `json:"display"`
         
    }
    
    // LimitedProductData used by Catalog
    type LimitedProductData struct {

        
            Slug string  `json:"slug"`
            Sizes []string  `json:"sizes"`
            CountryOfOrigin string  `json:"country_of_origin"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            ShortDescription string  `json:"short_description"`
            Identifier map[string]interface{}  `json:"identifier"`
            ItemCode string  `json:"item_code"`
            Price map[string]interface{}  `json:"price"`
            Quantity float64  `json:"quantity"`
            Attributes map[string]interface{}  `json:"attributes"`
            Images []string  `json:"images"`
         
    }
    
    // Price used by Catalog
    type Price struct {

        
            MinMarked float64  `json:"min_marked"`
            MaxMarked float64  `json:"max_marked"`
            MaxEffective float64  `json:"max_effective"`
            MinEffective float64  `json:"min_effective"`
            Currency string  `json:"currency"`
         
    }
    
    // GetProducts used by Catalog
    type GetProducts struct {

        
            MinQuantity float64  `json:"min_quantity"`
            Sizes []Size  `json:"sizes"`
            MaxQuantity float64  `json:"max_quantity"`
            ProductUID float64  `json:"product_uid"`
            AllowRemove bool  `json:"allow_remove"`
            ProductDetails LimitedProductData  `json:"product_details"`
            Price Price  `json:"price"`
            AutoSelect bool  `json:"auto_select"`
            AutoAddToCart bool  `json:"auto_add_to_cart"`
         
    }
    
    // GetProductBundleResponse used by Catalog
    type GetProductBundleResponse struct {

        
            Slug string  `json:"slug"`
            Choice string  `json:"choice"`
            PageVisibility []string  `json:"page_visibility"`
            Name string  `json:"name"`
            Logo string  `json:"logo"`
            Products []GetProducts  `json:"products"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            CompanyID float64  `json:"company_id"`
            IsActive bool  `json:"is_active"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // ProductBundleUpdateRequest used by Catalog
    type ProductBundleUpdateRequest struct {

        
            Slug string  `json:"slug"`
            Choice string  `json:"choice"`
            Name string  `json:"name"`
            PageVisibility []string  `json:"page_visibility"`
            Logo string  `json:"logo"`
            Products []ProductBundleItem  `json:"products"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            CompanyID float64  `json:"company_id"`
            IsActive bool  `json:"is_active"`
            ModifiedOn string  `json:"modified_on"`
            Meta map[string]interface{}  `json:"meta"`
         
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

        
            Guide Guide  `json:"guide"`
            Image string  `json:"image"`
            Description string  `json:"description"`
            CreatedOn string  `json:"created_on"`
            Name string  `json:"name"`
            Title string  `json:"title"`
            BrandID float64  `json:"brand_id"`
            Subtitle string  `json:"subtitle"`
            ModifiedOn string  `json:"modified_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            CompanyID float64  `json:"company_id"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ID string  `json:"id"`
            Active bool  `json:"active"`
            Tag string  `json:"tag"`
         
    }
    
    // SuccessResponse used by Catalog
    type SuccessResponse struct {

        
            Success bool  `json:"success"`
         
    }
    
    // SizeGuideResponse used by Catalog
    type SizeGuideResponse struct {

        
            CreatedOn string  `json:"created_on"`
            Guide map[string]interface{}  `json:"guide"`
            Active bool  `json:"active"`
            Name string  `json:"name"`
            BrandID float64  `json:"brand_id"`
            Title string  `json:"title"`
            Subtitle string  `json:"subtitle"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            CompanyID float64  `json:"company_id"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ID string  `json:"id"`
            ModifiedOn string  `json:"modified_on"`
            Tag string  `json:"tag"`
         
    }
    
    // SEOData used by Catalog
    type SEOData struct {

        
            Description string  `json:"description"`
            Title string  `json:"title"`
         
    }
    
    // MOQData used by Catalog
    type MOQData struct {

        
            Minimum float64  `json:"minimum"`
            IncrementUnit float64  `json:"increment_unit"`
            Maximum float64  `json:"maximum"`
         
    }
    
    // OwnerAppItemResponse used by Catalog
    type OwnerAppItemResponse struct {

        
            AltText map[string]interface{}  `json:"alt_text"`
            Seo SEOData  `json:"seo"`
            IsGift bool  `json:"is_gift"`
            Moq MOQData  `json:"moq"`
            IsCod bool  `json:"is_cod"`
         
    }
    
    // ApplicationItemSEO used by Catalog
    type ApplicationItemSEO struct {

        
            Description string  `json:"description"`
            Title string  `json:"title"`
         
    }
    
    // MetaFields used by Catalog
    type MetaFields struct {

        
            Value string  `json:"value"`
            Key string  `json:"key"`
         
    }
    
    // ApplicationItemMOQ used by Catalog
    type ApplicationItemMOQ struct {

        
            Minimum float64  `json:"minimum"`
            IncrementUnit float64  `json:"increment_unit"`
            Maximum float64  `json:"maximum"`
         
    }
    
    // ApplicationItemMeta used by Catalog
    type ApplicationItemMeta struct {

        
            AltText map[string]interface{}  `json:"alt_text"`
            Seo ApplicationItemSEO  `json:"seo"`
            CustomMeta []MetaFields  `json:"_custom_meta"`
            IsGift bool  `json:"is_gift"`
            Moq ApplicationItemMOQ  `json:"moq"`
            IsCod bool  `json:"is_cod"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
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
    
    // PageResponseType used by Catalog
    type PageResponseType struct {

        
            HasNext bool  `json:"has_next"`
            Current float64  `json:"current"`
            TotalCount float64  `json:"total_count"`
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

        
            Slug string  `json:"slug"`
            Name string  `json:"name"`
            Key string  `json:"key"`
            Logo string  `json:"logo"`
            Priority float64  `json:"priority"`
            IsActive bool  `json:"is_active"`
            Unit string  `json:"unit"`
            DisplayType string  `json:"display_type"`
         
    }
    
    // AppConfigurationDetail used by Catalog
    type AppConfigurationDetail struct {

        
            Slug string  `json:"slug"`
            Name string  `json:"name"`
            IsDefault bool  `json:"is_default"`
            Logo string  `json:"logo"`
            Priority float64  `json:"priority"`
            AppID string  `json:"app_id"`
            IsActive bool  `json:"is_active"`
            TemplateSlugs []string  `json:"template_slugs"`
            Attributes []AttributeDetailsGroup  `json:"attributes"`
         
    }
    
    // ConfigSuccessResponse used by Catalog
    type ConfigSuccessResponse struct {

        
            Message string  `json:"message"`
         
    }
    
    // AppConfigurationsSort used by Catalog
    type AppConfigurationsSort struct {

        
            DefaultKey string  `json:"default_key"`
            Name string  `json:"name"`
            Key string  `json:"key"`
            IsDefault bool  `json:"is_default"`
            Logo string  `json:"logo"`
            Priority float64  `json:"priority"`
            AppID string  `json:"app_id"`
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
            Units []map[string]interface{}  `json:"units"`
            FilterTypes []string  `json:"filter_types"`
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
            Variant map[string]interface{}  `json:"variant"`
            Similar map[string]interface{}  `json:"similar"`
            Compare map[string]interface{}  `json:"compare"`
         
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
            Logo string  `json:"logo"`
            Priority float64  `json:"priority"`
            IsActive bool  `json:"is_active"`
         
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

        
            BucketPoints []ConfigurationBucketPoints  `json:"bucket_points"`
            MapValues []map[string]interface{}  `json:"map_values"`
            Map map[string]interface{}  `json:"map"`
            Sort string  `json:"sort"`
            Value string  `json:"value"`
            Condition string  `json:"condition"`
         
    }
    
    // ConfigurationListingFilterConfig used by Catalog
    type ConfigurationListingFilterConfig struct {

        
            Name string  `json:"name"`
            Key string  `json:"key"`
            Logo string  `json:"logo"`
            Priority float64  `json:"priority"`
            IsActive bool  `json:"is_active"`
            Type string  `json:"type"`
            DisplayName string  `json:"display_name"`
            ValueConfig ConfigurationListingFilterValue  `json:"value_config"`
         
    }
    
    // ConfigurationListingFilter used by Catalog
    type ConfigurationListingFilter struct {

        
            AttributeConfig []ConfigurationListingFilterConfig  `json:"attribute_config"`
            AllowSingle bool  `json:"allow_single"`
         
    }
    
    // ConfigurationListing used by Catalog
    type ConfigurationListing struct {

        
            Sort ConfigurationListingSort  `json:"sort"`
            Filter ConfigurationListingFilter  `json:"filter"`
         
    }
    
    // ProductSize used by Catalog
    type ProductSize struct {

        
            Max float64  `json:"max"`
            Min float64  `json:"min"`
         
    }
    
    // ConfigurationProductVariantConfig used by Catalog
    type ConfigurationProductVariantConfig struct {

        
            Name string  `json:"name"`
            Key string  `json:"key"`
            Logo string  `json:"logo"`
            Priority float64  `json:"priority"`
            IsActive bool  `json:"is_active"`
            Size ProductSize  `json:"size"`
            DisplayType string  `json:"display_type"`
         
    }
    
    // ConfigurationProductVariant used by Catalog
    type ConfigurationProductVariant struct {

        
            Config []ConfigurationProductVariantConfig  `json:"config"`
         
    }
    
    // ConfigurationProductConfig used by Catalog
    type ConfigurationProductConfig struct {

        
            Title string  `json:"title"`
            Key string  `json:"key"`
            Logo string  `json:"logo"`
            Subtitle string  `json:"subtitle"`
            Priority float64  `json:"priority"`
            IsActive bool  `json:"is_active"`
            Size ProductSize  `json:"size"`
         
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
    
    // AppCatalogConfiguration used by Catalog
    type AppCatalogConfiguration struct {

        
            Listing ConfigurationListing  `json:"listing"`
            CreatedOn string  `json:"created_on"`
            ConfigType string  `json:"config_type"`
            Product ConfigurationProduct  `json:"product"`
            AppID string  `json:"app_id"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Type string  `json:"type"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ID string  `json:"id"`
            ModifiedOn string  `json:"modified_on"`
            ConfigID string  `json:"config_id"`
         
    }
    
    // GetAppCatalogConfiguration used by Catalog
    type GetAppCatalogConfiguration struct {

        
            IsDefault bool  `json:"is_default"`
            Data AppCatalogConfiguration  `json:"data"`
         
    }
    
    // AppConfiguration used by Catalog
    type AppConfiguration struct {

        
            Listing ConfigurationListing  `json:"listing"`
            CreatedOn string  `json:"created_on"`
            ConfigType string  `json:"config_type"`
            Product ConfigurationProduct  `json:"product"`
            AppID string  `json:"app_id"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Type string  `json:"type"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ModifiedOn string  `json:"modified_on"`
            ConfigID string  `json:"config_id"`
         
    }
    
    // GetCatalogConfigurationDetailsSchemaListing used by Catalog
    type GetCatalogConfigurationDetailsSchemaListing struct {

        
            Sort map[string]interface{}  `json:"sort"`
            Filter map[string]interface{}  `json:"filter"`
         
    }
    
    // EntityConfiguration used by Catalog
    type EntityConfiguration struct {

        
            Listing GetCatalogConfigurationDetailsSchemaListing  `json:"listing"`
            ConfigType string  `json:"config_type"`
            Product GetCatalogConfigurationDetailsProduct  `json:"product"`
            AppID string  `json:"app_id"`
            ID string  `json:"id"`
            ConfigID string  `json:"config_id"`
         
    }
    
    // GetAppCatalogEntityConfiguration used by Catalog
    type GetAppCatalogEntityConfiguration struct {

        
            IsDefault bool  `json:"is_default"`
            Data EntityConfiguration  `json:"data"`
         
    }
    
    // ProductSortOn used by Catalog
    type ProductSortOn struct {

        
            IsSelected bool  `json:"is_selected"`
            Name string  `json:"name"`
            Value string  `json:"value"`
         
    }
    
    // ProductFiltersValue used by Catalog
    type ProductFiltersValue struct {

        
            Count float64  `json:"count"`
            Min float64  `json:"min"`
            IsSelected bool  `json:"is_selected"`
            SelectedMax float64  `json:"selected_max"`
            DisplayFormat string  `json:"display_format"`
            Max float64  `json:"max"`
            SelectedMin float64  `json:"selected_min"`
            QueryFormat string  `json:"query_format"`
            CurrencySymbol string  `json:"currency_symbol"`
            Value interface{}  `json:"value"`
            CurrencyCode string  `json:"currency_code"`
            Display string  `json:"display"`
         
    }
    
    // ProductFiltersKey used by Catalog
    type ProductFiltersKey struct {

        
            Name string  `json:"name"`
            Logo string  `json:"logo"`
            Kind string  `json:"kind"`
            Operators []string  `json:"operators"`
            Display string  `json:"display"`
         
    }
    
    // ProductFilters used by Catalog
    type ProductFilters struct {

        
            Values []ProductFiltersValue  `json:"values"`
            Key ProductFiltersKey  `json:"key"`
         
    }
    
    // GetCollectionQueryOptionResponse used by Catalog
    type GetCollectionQueryOptionResponse struct {

        
            SortOn []ProductSortOn  `json:"sort_on"`
            Operators map[string]string  `json:"operators"`
            Filters []ProductFilters  `json:"filters"`
         
    }
    
    // Media1 used by Catalog
    type Media1 struct {

        
            Type string  `json:"type"`
            URL string  `json:"url"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // BannerImage used by Catalog
    type BannerImage struct {

        
            Type string  `json:"type"`
            URL string  `json:"url"`
            AspectRatio string  `json:"aspect_ratio"`
         
    }
    
    // ImageUrls used by Catalog
    type ImageUrls struct {

        
            Portrait BannerImage  `json:"portrait"`
            Landscape BannerImage  `json:"landscape"`
         
    }
    
    // CollectionQuery used by Catalog
    type CollectionQuery struct {

        
            Attribute string  `json:"attribute"`
            Op string  `json:"op"`
            Value []interface{}  `json:"value"`
         
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

        
            AllowSort bool  `json:"allow_sort"`
            Logo Media1  `json:"logo"`
            Banners ImageUrls  `json:"banners"`
            AppID string  `json:"app_id"`
            Cron map[string]interface{}  `json:"cron"`
            Query []CollectionQuery  `json:"query"`
            Tag []string  `json:"tag"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            UID string  `json:"uid"`
            IsActive bool  `json:"is_active"`
            Action Action  `json:"action"`
            Type string  `json:"type"`
            Slug string  `json:"slug"`
            Description string  `json:"description"`
            Name string  `json:"name"`
            AllowFacets bool  `json:"allow_facets"`
            Priority float64  `json:"priority"`
            Badge map[string]interface{}  `json:"badge"`
            Schedule map[string]interface{}  `json:"_schedule"`
            Meta map[string]interface{}  `json:"meta"`
         
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

        
            UserID string  `json:"user_id"`
            Email string  `json:"email"`
            UID string  `json:"uid"`
            Username string  `json:"username"`
         
    }
    
    // SeoDetail used by Catalog
    type SeoDetail struct {

        
            Description string  `json:"description"`
            Title string  `json:"title"`
         
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

        
            Start string  `json:"start"`
            Duration float64  `json:"duration"`
            End string  `json:"end"`
            Cron string  `json:"cron"`
            NextSchedule []NextSchedule  `json:"next_schedule"`
         
    }
    
    // CreateCollection used by Catalog
    type CreateCollection struct {

        
            AllowSort bool  `json:"allow_sort"`
            IsVisible bool  `json:"is_visible"`
            Logo CollectionImage  `json:"logo"`
            Banners CollectionBanner  `json:"banners"`
            AppID string  `json:"app_id"`
            CreatedBy UserInfo  `json:"created_by"`
            Query []CollectionQuery  `json:"query"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Published bool  `json:"published"`
            ModifiedBy UserInfo  `json:"modified_by"`
            Tags []string  `json:"tags"`
            IsActive bool  `json:"is_active"`
            Type string  `json:"type"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            SortOn string  `json:"sort_on"`
            Slug string  `json:"slug"`
            Description string  `json:"description"`
            Name string  `json:"name"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            AllowFacets bool  `json:"allow_facets"`
            Seo SeoDetail  `json:"seo"`
            Priority float64  `json:"priority"`
            Badge CollectionBadge  `json:"badge"`
            Schedule CollectionSchedule  `json:"_schedule"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // CollectionCreateResponse used by Catalog
    type CollectionCreateResponse struct {

        
            AllowSort bool  `json:"allow_sort"`
            Logo BannerImage  `json:"logo"`
            Banners ImageUrls  `json:"banners"`
            AppID string  `json:"app_id"`
            Cron map[string]interface{}  `json:"cron"`
            Query []CollectionQuery  `json:"query"`
            Tag []string  `json:"tag"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            IsActive bool  `json:"is_active"`
            Type string  `json:"type"`
            SortOn string  `json:"sort_on"`
            Slug string  `json:"slug"`
            Description string  `json:"description"`
            Name string  `json:"name"`
            AllowFacets bool  `json:"allow_facets"`
            Priority float64  `json:"priority"`
            Badge map[string]interface{}  `json:"badge"`
            Schedule map[string]interface{}  `json:"_schedule"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // CollectionDetailResponse used by Catalog
    type CollectionDetailResponse struct {

        
            AllowSort bool  `json:"allow_sort"`
            Logo Media1  `json:"logo"`
            Banners ImageUrls  `json:"banners"`
            AppID string  `json:"app_id"`
            Cron map[string]interface{}  `json:"cron"`
            Query []CollectionQuery  `json:"query"`
            Tag []string  `json:"tag"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            UID string  `json:"uid"`
            IsActive bool  `json:"is_active"`
            Type string  `json:"type"`
            Slug string  `json:"slug"`
            Description string  `json:"description"`
            Name string  `json:"name"`
            AllowFacets bool  `json:"allow_facets"`
            Priority float64  `json:"priority"`
            Badge map[string]interface{}  `json:"badge"`
            Schedule map[string]interface{}  `json:"_schedule"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // UpdateCollection used by Catalog
    type UpdateCollection struct {

        
            AllowSort bool  `json:"allow_sort"`
            IsVisible bool  `json:"is_visible"`
            Logo CollectionImage  `json:"logo"`
            Banners CollectionBanner  `json:"banners"`
            Query []CollectionQuery  `json:"query"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Published bool  `json:"published"`
            ModifiedBy UserInfo  `json:"modified_by"`
            Tags []string  `json:"tags"`
            IsActive bool  `json:"is_active"`
            Type string  `json:"type"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            SortOn string  `json:"sort_on"`
            Slug string  `json:"slug"`
            Description string  `json:"description"`
            Name string  `json:"name"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            AllowFacets bool  `json:"allow_facets"`
            Seo SeoDetail  `json:"seo"`
            Priority float64  `json:"priority"`
            Badge CollectionBadge  `json:"badge"`
            Schedule CollectionSchedule  `json:"_schedule"`
            Meta map[string]interface{}  `json:"meta"`
         
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
    
    // ProductBrand used by Catalog
    type ProductBrand struct {

        
            Logo Media1  `json:"logo"`
            UID float64  `json:"uid"`
            Action Action  `json:"action"`
            Name string  `json:"name"`
         
    }
    
    // Price1 used by Catalog
    type Price1 struct {

        
            CurrencyCode string  `json:"currency_code"`
            Max float64  `json:"max"`
            Min float64  `json:"min"`
            CurrencySymbol string  `json:"currency_symbol"`
         
    }
    
    // ProductListingPrice used by Catalog
    type ProductListingPrice struct {

        
            Marked Price1  `json:"marked"`
            Effective Price1  `json:"effective"`
         
    }
    
    // ProductListingDetail used by Catalog
    type ProductListingDetail struct {

        
            HasVariant bool  `json:"has_variant"`
            Color string  `json:"color"`
            Similars []string  `json:"similars"`
            Sellable bool  `json:"sellable"`
            Discount string  `json:"discount"`
            UID float64  `json:"uid"`
            Rating float64  `json:"rating"`
            Highlights []string  `json:"highlights"`
            Type string  `json:"type"`
            RatingCount float64  `json:"rating_count"`
            Attributes map[string]interface{}  `json:"attributes"`
            Slug string  `json:"slug"`
            ShortDescription string  `json:"short_description"`
            ItemType string  `json:"item_type"`
            PromoMeta map[string]interface{}  `json:"promo_meta"`
            ItemCode string  `json:"item_code"`
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
            Tryouts []string  `json:"tryouts"`
            Description string  `json:"description"`
            ProductOnlineDate string  `json:"product_online_date"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            Name string  `json:"name"`
            Brand ProductBrand  `json:"brand"`
            Price ProductListingPrice  `json:"price"`
            Medias []Media1  `json:"medias"`
            ImageNature string  `json:"image_nature"`
         
    }
    
    // GetCollectionItemsResponse used by Catalog
    type GetCollectionItemsResponse struct {

        
            Items []ProductListingDetail  `json:"items"`
            Page Page  `json:"page"`
            SortOn []ProductSortOn  `json:"sort_on"`
            Filters []ProductFilters  `json:"filters"`
         
    }
    
    // CollectionItem used by Catalog
    type CollectionItem struct {

        
            ItemID float64  `json:"item_id"`
            Priority float64  `json:"priority"`
            Action string  `json:"action"`
         
    }
    
    // CollectionItemUpdate used by Catalog
    type CollectionItemUpdate struct {

        
            AllowSort bool  `json:"allow_sort"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            AllowFacets bool  `json:"allow_facets"`
            Items []CollectionItem  `json:"items"`
            Type string  `json:"type"`
            Query []CollectionQuery  `json:"query"`
         
    }
    
    // UpdatedResponse used by Catalog
    type UpdatedResponse struct {

        
            ItemsNotUpdated []float64  `json:"items_not_updated"`
            Message string  `json:"message"`
         
    }
    
    // CatalogInsightBrand used by Catalog
    type CatalogInsightBrand struct {

        
            TotalArticles float64  `json:"total_articles"`
            Name string  `json:"name"`
            ArticleFreshness float64  `json:"article_freshness"`
            TotalSizes float64  `json:"total_sizes"`
            AvailableArticles float64  `json:"available_articles"`
            AvailableSizes float64  `json:"available_sizes"`
         
    }
    
    // CatalogInsightItem used by Catalog
    type CatalogInsightItem struct {

        
            Count float64  `json:"count"`
            OutOfStockCount float64  `json:"out_of_stock_count"`
            SellableCount float64  `json:"sellable_count"`
         
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
            Platform string  `json:"platform"`
            OptLevel string  `json:"opt_level"`
            Enabled bool  `json:"enabled"`
            StoreIds []float64  `json:"store_ids"`
            CompanyID float64  `json:"company_id"`
         
    }
    
    // CompanyOptIn used by Catalog
    type CompanyOptIn struct {

        
            BrandIds []float64  `json:"brand_ids"`
            CreatedOn float64  `json:"created_on"`
            Platform string  `json:"platform"`
            OptLevel string  `json:"opt_level"`
            Enabled bool  `json:"enabled"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            StoreIds []float64  `json:"store_ids"`
            CompanyID float64  `json:"company_id"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ModifiedOn float64  `json:"modified_on"`
         
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
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // CompanyBrandDetail used by Catalog
    type CompanyBrandDetail struct {

        
            BrandID float64  `json:"brand_id"`
            BrandName string  `json:"brand_name"`
            CompanyID float64  `json:"company_id"`
            TotalArticle float64  `json:"total_article"`
         
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
            CreatedOn string  `json:"created_on"`
            StoreType string  `json:"store_type"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            AdditionalContacts []map[string]interface{}  `json:"additional_contacts"`
            Address map[string]interface{}  `json:"address"`
            Timing map[string]interface{}  `json:"timing"`
            CompanyID float64  `json:"company_id"`
            Manager map[string]interface{}  `json:"manager"`
            DisplayName string  `json:"display_name"`
            ModifiedOn string  `json:"modified_on"`
            Documents []map[string]interface{}  `json:"documents"`
         
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
    
    // AttributeSchemaRange used by Catalog
    type AttributeSchemaRange struct {

        
            Max float64  `json:"max"`
            Min float64  `json:"min"`
         
    }
    
    // AttributeMaster used by Catalog
    type AttributeMaster struct {

        
            Range AttributeSchemaRange  `json:"range"`
            Multi bool  `json:"multi"`
            Format string  `json:"format"`
            AllowedValues []string  `json:"allowed_values"`
            Type string  `json:"type"`
            Mandatory bool  `json:"mandatory"`
         
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

        
            Slug string  `json:"slug"`
            IsNested bool  `json:"is_nested"`
            EnabledForEndConsumer bool  `json:"enabled_for_end_consumer"`
            Description string  `json:"description"`
            Name string  `json:"name"`
            Departments []string  `json:"departments"`
            Details AttributeMasterDetails  `json:"details"`
            Schema AttributeMaster  `json:"schema"`
            Logo string  `json:"logo"`
            Filters AttributeMasterFilter  `json:"filters"`
            ID string  `json:"id"`
            Meta AttributeMasterMeta  `json:"meta"`
         
    }
    
    // CategoriesResponse used by Catalog
    type CategoriesResponse struct {

        
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            SlugKey string  `json:"slug_key"`
            TemplateSlug string  `json:"template_slug"`
         
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
            Code string  `json:"code"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // UserSerializer used by Catalog
    type UserSerializer struct {

        
            UserID string  `json:"user_id"`
            UID string  `json:"uid"`
            Contact string  `json:"contact"`
            Username string  `json:"username"`
            ID string  `json:"_id"`
         
    }
    
    // GetDepartment used by Catalog
    type GetDepartment struct {

        
            Slug string  `json:"slug"`
            CreatedOn string  `json:"created_on"`
            Search string  `json:"search"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            PriorityOrder float64  `json:"priority_order"`
            Logo string  `json:"logo"`
            CreatedBy UserSerializer  `json:"created_by"`
            PageNo float64  `json:"page_no"`
            ItemType string  `json:"item_type"`
            Synonyms []string  `json:"synonyms"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            IsActive bool  `json:"is_active"`
            PageSize float64  `json:"page_size"`
            ModifiedOn string  `json:"modified_on"`
         
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
            Code string  `json:"code"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // DepartmentCreateUpdate used by Catalog
    type DepartmentCreateUpdate struct {

        
            Slug string  `json:"slug"`
            PriorityOrder float64  `json:"priority_order"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            Logo string  `json:"logo"`
            Cls string  `json:"_cls"`
            Synonyms []string  `json:"synonyms"`
            Tags []string  `json:"tags"`
            IsActive bool  `json:"is_active"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
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

        
            SuperUser bool  `json:"super_user"`
            UserID string  `json:"user_id"`
            Contact string  `json:"contact"`
            Username string  `json:"username"`
         
    }
    
    // DepartmentModel used by Catalog
    type DepartmentModel struct {

        
            Slug string  `json:"slug"`
            CreatedOn string  `json:"created_on"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            PriorityOrder float64  `json:"priority_order"`
            Logo string  `json:"logo"`
            Cls string  `json:"_cls"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            VerifiedBy UserDetail  `json:"verified_by"`
            Synonyms []string  `json:"synonyms"`
            ModifiedBy UserDetail  `json:"modified_by"`
            IsActive bool  `json:"is_active"`
            CreatedBy UserDetail  `json:"created_by"`
            ID interface{}  `json:"_id"`
            ModifiedOn string  `json:"modified_on"`
            VerifiedOn string  `json:"verified_on"`
         
    }
    
    // ProductTemplate used by Catalog
    type ProductTemplate struct {

        
            Slug string  `json:"slug"`
            Categories []string  `json:"categories"`
            IsArchived bool  `json:"is_archived"`
            Description string  `json:"description"`
            IsExpirable bool  `json:"is_expirable"`
            CreatedOn string  `json:"created_on"`
            Name string  `json:"name"`
            Departments []string  `json:"departments"`
            Logo string  `json:"logo"`
            ModifiedOn string  `json:"modified_on"`
            IsPhysical bool  `json:"is_physical"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            IsActive bool  `json:"is_active"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Tag string  `json:"tag"`
            Attributes []string  `json:"attributes"`
         
    }
    
    // TemplatesResponse used by Catalog
    type TemplatesResponse struct {

        
            Items ProductTemplate  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // TemplateDetails used by Catalog
    type TemplateDetails struct {

        
            Slug string  `json:"slug"`
            Categories []string  `json:"categories"`
            IsArchived bool  `json:"is_archived"`
            Description string  `json:"description"`
            IsExpirable bool  `json:"is_expirable"`
            Name string  `json:"name"`
            Departments []string  `json:"departments"`
            Logo string  `json:"logo"`
            IsPhysical bool  `json:"is_physical"`
            IsActive bool  `json:"is_active"`
            ID string  `json:"id"`
            Tag string  `json:"tag"`
            Attributes []string  `json:"attributes"`
         
    }
    
    // Properties used by Catalog
    type Properties struct {

        
            CountryOfOrigin map[string]interface{}  `json:"country_of_origin"`
            ReturnConfig map[string]interface{}  `json:"return_config"`
            HsnCode map[string]interface{}  `json:"hsn_code"`
            Command map[string]interface{}  `json:"command"`
            ProductPublish map[string]interface{}  `json:"product_publish"`
            CategorySlug map[string]interface{}  `json:"category_slug"`
            TraderType map[string]interface{}  `json:"trader_type"`
            Highlights map[string]interface{}  `json:"highlights"`
            Tags map[string]interface{}  `json:"tags"`
            IsActive map[string]interface{}  `json:"is_active"`
            BrandUID map[string]interface{}  `json:"brand_uid"`
            MultiSize map[string]interface{}  `json:"multi_size"`
            Slug map[string]interface{}  `json:"slug"`
            Sizes map[string]interface{}  `json:"sizes"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            Trader map[string]interface{}  `json:"trader"`
            ShortDescription map[string]interface{}  `json:"short_description"`
            ItemType map[string]interface{}  `json:"item_type"`
            Variants map[string]interface{}  `json:"variants"`
            ProductGroupTag map[string]interface{}  `json:"product_group_tag"`
            ItemCode map[string]interface{}  `json:"item_code"`
            NoOfBoxes map[string]interface{}  `json:"no_of_boxes"`
            Description map[string]interface{}  `json:"description"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            Name map[string]interface{}  `json:"name"`
            IsDependent map[string]interface{}  `json:"is_dependent"`
            SizeGuide map[string]interface{}  `json:"size_guide"`
            Currency map[string]interface{}  `json:"currency"`
            Media map[string]interface{}  `json:"media"`
         
    }
    
    // GlobalValidation used by Catalog
    type GlobalValidation struct {

        
            Required []string  `json:"required"`
            Properties Properties  `json:"properties"`
            Description string  `json:"description"`
            Title string  `json:"title"`
            Type string  `json:"type"`
            Definitions map[string]interface{}  `json:"definitions"`
         
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

        
            UserID string  `json:"user_id"`
            Email string  `json:"email"`
            UID string  `json:"uid"`
            Username string  `json:"username"`
         
    }
    
    // ProductTemplateExportResponse used by Catalog
    type ProductTemplateExportResponse struct {

        
            Status string  `json:"status"`
            CompletedOn string  `json:"completed_on"`
            Type string  `json:"type"`
            Filters map[string]interface{}  `json:"filters"`
            SellerID float64  `json:"seller_id"`
            TaskID string  `json:"task_id"`
            NotificationEmails []string  `json:"notification_emails"`
            CreatedBy UserInfo1  `json:"created_by"`
            URL string  `json:"url"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // ProductDownloadsResponse used by Catalog
    type ProductDownloadsResponse struct {

        
            Items []ProductTemplateExportResponse  `json:"items"`
         
    }
    
    // ProductTemplateExportFilterRequest used by Catalog
    type ProductTemplateExportFilterRequest struct {

        
            CatalogueTypes []string  `json:"catalogue_types"`
            Brands []string  `json:"brands"`
            Templates []string  `json:"templates"`
            ToDate string  `json:"to_date"`
            FromDate string  `json:"from_date"`
         
    }
    
    // ProductTemplateDownloadsExport used by Catalog
    type ProductTemplateDownloadsExport struct {

        
            Filters ProductTemplateExportFilterRequest  `json:"filters"`
            NotificationEmails []string  `json:"notification_emails"`
            Type string  `json:"type"`
         
    }
    
    // ProductConfigurationDownloads used by Catalog
    type ProductConfigurationDownloads struct {

        
            Multivalue bool  `json:"multivalue"`
            Data []map[string]interface{}  `json:"data"`
         
    }
    
    // Hierarchy used by Catalog
    type Hierarchy struct {

        
            L1 float64  `json:"l1"`
            L2 float64  `json:"l2"`
            Department float64  `json:"department"`
         
    }
    
    // Media2 used by Catalog
    type Media2 struct {

        
            Logo string  `json:"logo"`
            Portrait string  `json:"portrait"`
            Landscape string  `json:"landscape"`
         
    }
    
    // CategoryMappingValues used by Catalog
    type CategoryMappingValues struct {

        
            CatalogID float64  `json:"catalog_id"`
            Name string  `json:"name"`
         
    }
    
    // CategoryMapping used by Catalog
    type CategoryMapping struct {

        
            Google CategoryMappingValues  `json:"google"`
            Ajio CategoryMappingValues  `json:"ajio"`
            Facebook CategoryMappingValues  `json:"facebook"`
         
    }
    
    // Category used by Catalog
    type Category struct {

        
            Slug string  `json:"slug"`
            CreatedOn string  `json:"created_on"`
            ID string  `json:"id"`
            Departments []float64  `json:"departments"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            Priority float64  `json:"priority"`
            Synonyms []string  `json:"synonyms"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            IsActive bool  `json:"is_active"`
            Hierarchy []Hierarchy  `json:"hierarchy"`
            Tryouts []string  `json:"tryouts"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Level float64  `json:"level"`
            ModifiedOn string  `json:"modified_on"`
            Media Media2  `json:"media"`
            Marketplaces CategoryMapping  `json:"marketplaces"`
         
    }
    
    // CategoryResponse used by Catalog
    type CategoryResponse struct {

        
            Items []Category  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // CategoryRequestBody used by Catalog
    type CategoryRequestBody struct {

        
            Slug string  `json:"slug"`
            Departments []float64  `json:"departments"`
            Name string  `json:"name"`
            Priority float64  `json:"priority"`
            Synonyms []string  `json:"synonyms"`
            IsActive bool  `json:"is_active"`
            Hierarchy []Hierarchy  `json:"hierarchy"`
            Tryouts []string  `json:"tryouts"`
            Level float64  `json:"level"`
            Media Media2  `json:"media"`
            Marketplaces CategoryMapping  `json:"marketplaces"`
         
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

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // ReturnConfigResponse used by Catalog
    type ReturnConfigResponse struct {

        
            Time float64  `json:"time"`
            Unit string  `json:"unit"`
            Returnable bool  `json:"returnable"`
         
    }
    
    // ProductPublish used by Catalog
    type ProductPublish struct {

        
            ProductOnlineDate string  `json:"product_online_date"`
            IsSet bool  `json:"is_set"`
         
    }
    
    // VerifiedBy used by Catalog
    type VerifiedBy struct {

        
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
         
    }
    
    // Image used by Catalog
    type Image struct {

        
            URL string  `json:"url"`
            SecureURL string  `json:"secure_url"`
            AspectRatioF float64  `json:"aspect_ratio_f"`
            AspectRatio string  `json:"aspect_ratio"`
         
    }
    
    // Trader used by Catalog
    type Trader struct {

        
            Address []string  `json:"address"`
            Name string  `json:"name"`
            Type string  `json:"type"`
         
    }
    
    // NetQuantityResponse used by Catalog
    type NetQuantityResponse struct {

        
            Unit string  `json:"unit"`
            Value float64  `json:"value"`
         
    }
    
    // Logo used by Catalog
    type Logo struct {

        
            URL string  `json:"url"`
            SecureURL string  `json:"secure_url"`
            AspectRatioF float64  `json:"aspect_ratio_f"`
            AspectRatio string  `json:"aspect_ratio"`
         
    }
    
    // Brand used by Catalog
    type Brand struct {

        
            Logo Logo  `json:"logo"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // ProductSchemaV2 used by Catalog
    type ProductSchemaV2 struct {

        
            CountryOfOrigin string  `json:"country_of_origin"`
            Color string  `json:"color"`
            ReturnConfig ReturnConfigResponse  `json:"return_config"`
            HsnCode string  `json:"hsn_code"`
            ProductPublish ProductPublish  `json:"product_publish"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            AllSizes []map[string]interface{}  `json:"all_sizes"`
            Category map[string]interface{}  `json:"category"`
            VerifiedOn string  `json:"verified_on"`
            IsExpirable bool  `json:"is_expirable"`
            CreatedOn string  `json:"created_on"`
            Pending string  `json:"pending"`
            Departments []float64  `json:"departments"`
            UID float64  `json:"uid"`
            AllIdentifiers []string  `json:"all_identifiers"`
            CategorySlug string  `json:"category_slug"`
            VariantGroup map[string]interface{}  `json:"variant_group"`
            VariantMedia map[string]interface{}  `json:"variant_media"`
            VerifiedBy VerifiedBy  `json:"verified_by"`
            Highlights []string  `json:"highlights"`
            Tags []string  `json:"tags"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            AllCompanyIds []float64  `json:"all_company_ids"`
            IsActive bool  `json:"is_active"`
            BrandUID float64  `json:"brand_uid"`
            TemplateTag string  `json:"template_tag"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            L3Mapping []string  `json:"l3_mapping"`
            MultiSize bool  `json:"multi_size"`
            Images []Image  `json:"images"`
            Slug string  `json:"slug"`
            Sizes []map[string]interface{}  `json:"sizes"`
            Attributes map[string]interface{}  `json:"attributes"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            Trader []Trader  `json:"trader"`
            ShortDescription string  `json:"short_description"`
            ItemType string  `json:"item_type"`
            Variants map[string]interface{}  `json:"variants"`
            IsPhysical bool  `json:"is_physical"`
            ProductGroupTag []string  `json:"product_group_tag"`
            ItemCode string  `json:"item_code"`
            IsImageLessProduct bool  `json:"is_image_less_product"`
            Stage string  `json:"stage"`
            PrimaryColor string  `json:"primary_color"`
            Moq map[string]interface{}  `json:"moq"`
            CompanyID float64  `json:"company_id"`
            ID string  `json:"id"`
            ModifiedOn string  `json:"modified_on"`
            CategoryUID float64  `json:"category_uid"`
            NoOfBoxes float64  `json:"no_of_boxes"`
            Description string  `json:"description"`
            NetQuantity NetQuantityResponse  `json:"net_quantity"`
            Name string  `json:"name"`
            IsDependent bool  `json:"is_dependent"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            Brand Brand  `json:"brand"`
            SizeGuide string  `json:"size_guide"`
            IsSet bool  `json:"is_set"`
            Currency string  `json:"currency"`
            Media []Media1  `json:"media"`
            ImageNature string  `json:"image_nature"`
         
    }
    
    // ProductListingResponseV2 used by Catalog
    type ProductListingResponseV2 struct {

        
            Items []ProductSchemaV2  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // ReturnConfig used by Catalog
    type ReturnConfig struct {

        
            Time float64  `json:"time"`
            Unit string  `json:"unit"`
            Returnable bool  `json:"returnable"`
         
    }
    
    // ProductPublish1 used by Catalog
    type ProductPublish1 struct {

        
            ProductOnlineDate string  `json:"product_online_date"`
            IsSet bool  `json:"is_set"`
         
    }
    
    // CustomOrder used by Catalog
    type CustomOrder struct {

        
            ManufacturingTime float64  `json:"manufacturing_time"`
            ManufacturingTimeUnit string  `json:"manufacturing_time_unit"`
            IsCustomOrder bool  `json:"is_custom_order"`
         
    }
    
    // TeaserTag used by Catalog
    type TeaserTag struct {

        
            URL string  `json:"url"`
            Tag string  `json:"tag"`
         
    }
    
    // NetQuantity used by Catalog
    type NetQuantity struct {

        
            Unit string  `json:"unit"`
            Value float64  `json:"value"`
         
    }
    
    // TaxIdentifier used by Catalog
    type TaxIdentifier struct {

        
            ReportingHsn string  `json:"reporting_hsn"`
            HsnCodeID string  `json:"hsn_code_id"`
            HsnCode string  `json:"hsn_code"`
         
    }
    
    // ProductCreateUpdateSchemaV2 used by Catalog
    type ProductCreateUpdateSchemaV2 struct {

        
            CountryOfOrigin string  `json:"country_of_origin"`
            ReturnConfig ReturnConfig  `json:"return_config"`
            ProductPublish ProductPublish1  `json:"product_publish"`
            UID float64  `json:"uid"`
            Departments []float64  `json:"departments"`
            Requester string  `json:"requester"`
            CategorySlug string  `json:"category_slug"`
            VariantGroup map[string]interface{}  `json:"variant_group"`
            VariantMedia map[string]interface{}  `json:"variant_media"`
            Highlights []string  `json:"highlights"`
            Tags []string  `json:"tags"`
            IsActive bool  `json:"is_active"`
            BrandUID float64  `json:"brand_uid"`
            Action string  `json:"action"`
            TemplateTag string  `json:"template_tag"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            MultiSize bool  `json:"multi_size"`
            Attributes map[string]interface{}  `json:"attributes"`
            Slug string  `json:"slug"`
            Sizes []map[string]interface{}  `json:"sizes"`
            CustomOrder CustomOrder  `json:"custom_order"`
            Trader []Trader  `json:"trader"`
            BulkJobID string  `json:"bulk_job_id"`
            ShortDescription string  `json:"short_description"`
            ChangeRequestID interface{}  `json:"change_request_id"`
            ItemType string  `json:"item_type"`
            Variants map[string]interface{}  `json:"variants"`
            ProductGroupTag []string  `json:"product_group_tag"`
            ItemCode string  `json:"item_code"`
            IsImageLessProduct bool  `json:"is_image_less_product"`
            CompanyID float64  `json:"company_id"`
            NoOfBoxes float64  `json:"no_of_boxes"`
            Description string  `json:"description"`
            TeaserTag TeaserTag  `json:"teaser_tag"`
            Name string  `json:"name"`
            IsDependent bool  `json:"is_dependent"`
            NetQuantity NetQuantity  `json:"net_quantity"`
            TaxIdentifier TaxIdentifier  `json:"tax_identifier"`
            SizeGuide string  `json:"size_guide"`
            IsSet bool  `json:"is_set"`
            Currency string  `json:"currency"`
            Media []Media1  `json:"media"`
         
    }
    
    // ProductVariants used by Catalog
    type ProductVariants struct {

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            ItemCode string  `json:"item_code"`
            BrandUID float64  `json:"brand_uid"`
            CategoryUID float64  `json:"category_uid"`
            Media []Media1  `json:"media"`
         
    }
    
    // ProductVariantsResponse used by Catalog
    type ProductVariantsResponse struct {

        
            Page Page  `json:"page"`
            Variants []ProductVariants  `json:"variants"`
         
    }
    
    // AttributeMasterSerializer used by Catalog
    type AttributeMasterSerializer struct {

        
            IsNested bool  `json:"is_nested"`
            Logo string  `json:"logo"`
            Details AttributeMasterDetails  `json:"details"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            RawKey string  `json:"raw_key"`
            CreatedOn string  `json:"created_on"`
            Departments []string  `json:"departments"`
            Schema AttributeMaster  `json:"schema"`
            Variant bool  `json:"variant"`
            Filters AttributeMasterFilter  `json:"filters"`
            Tags []string  `json:"tags"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Unit string  `json:"unit"`
            Slug string  `json:"slug"`
            Synonyms map[string]interface{}  `json:"synonyms"`
            ModifiedOn string  `json:"modified_on"`
            EnabledForEndConsumer bool  `json:"enabled_for_end_consumer"`
            Description string  `json:"description"`
            Name string  `json:"name"`
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

        
            GtinType string  `json:"gtin_type"`
            GtinValue string  `json:"gtin_value"`
            Primary bool  `json:"primary"`
         
    }
    
    // AllSizes used by Catalog
    type AllSizes struct {

        
            ItemWeight float64  `json:"item_weight"`
            ItemDimensionsUnitOfMeasure string  `json:"item_dimensions_unit_of_measure"`
            Identifiers []ValidateIdentifier  `json:"identifiers"`
            ItemWidth float64  `json:"item_width"`
            ItemWeightUnitOfMeasure string  `json:"item_weight_unit_of_measure"`
            ItemLength float64  `json:"item_length"`
            ItemHeight float64  `json:"item_height"`
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

        
            CountryOfOrigin string  `json:"country_of_origin"`
            Color string  `json:"color"`
            ReturnConfig ReturnConfigResponse  `json:"return_config"`
            HsnCode string  `json:"hsn_code"`
            ProductPublish ProductPublished  `json:"product_publish"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            AllSizes []map[string]interface{}  `json:"all_sizes"`
            Category map[string]interface{}  `json:"category"`
            VerifiedOn string  `json:"verified_on"`
            IsExpirable bool  `json:"is_expirable"`
            CreatedOn string  `json:"created_on"`
            Pending string  `json:"pending"`
            Departments []float64  `json:"departments"`
            UID float64  `json:"uid"`
            AllIdentifiers []string  `json:"all_identifiers"`
            CategorySlug string  `json:"category_slug"`
            VariantGroup map[string]interface{}  `json:"variant_group"`
            VariantMedia map[string]interface{}  `json:"variant_media"`
            VerifiedBy VerifiedBy  `json:"verified_by"`
            Highlights []string  `json:"highlights"`
            Tags []string  `json:"tags"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            AllCompanyIds []float64  `json:"all_company_ids"`
            IsActive bool  `json:"is_active"`
            BrandUID float64  `json:"brand_uid"`
            TemplateTag string  `json:"template_tag"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            L3Mapping []string  `json:"l3_mapping"`
            MultiSize bool  `json:"multi_size"`
            Images []Image  `json:"images"`
            Slug string  `json:"slug"`
            Sizes []map[string]interface{}  `json:"sizes"`
            Attributes map[string]interface{}  `json:"attributes"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            Trader []Trader  `json:"trader"`
            ShortDescription string  `json:"short_description"`
            ItemType string  `json:"item_type"`
            Variants map[string]interface{}  `json:"variants"`
            IsPhysical bool  `json:"is_physical"`
            ProductGroupTag []string  `json:"product_group_tag"`
            ItemCode string  `json:"item_code"`
            IsImageLessProduct bool  `json:"is_image_less_product"`
            Stage string  `json:"stage"`
            PrimaryColor string  `json:"primary_color"`
            Moq map[string]interface{}  `json:"moq"`
            CompanyID float64  `json:"company_id"`
            ID string  `json:"id"`
            ModifiedOn string  `json:"modified_on"`
            CategoryUID float64  `json:"category_uid"`
            NoOfBoxes float64  `json:"no_of_boxes"`
            Description string  `json:"description"`
            NetQuantity NetQuantityResponse  `json:"net_quantity"`
            Name string  `json:"name"`
            IsDependent bool  `json:"is_dependent"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            Brand Brand  `json:"brand"`
            SizeGuide string  `json:"size_guide"`
            IsSet bool  `json:"is_set"`
            Currency string  `json:"currency"`
            Media []Media1  `json:"media"`
            ImageNature string  `json:"image_nature"`
         
    }
    
    // ProductListingResponse used by Catalog
    type ProductListingResponse struct {

        
            Items []Product  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // UserDetail1 used by Catalog
    type UserDetail1 struct {

        
            FullName string  `json:"full_name"`
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
         
    }
    
    // ProductBulkRequest used by Catalog
    type ProductBulkRequest struct {

        
            Template ProductTemplate  `json:"template"`
            CreatedOn string  `json:"created_on"`
            CancelledRecords []string  `json:"cancelled_records"`
            CreatedBy UserDetail1  `json:"created_by"`
            Total float64  `json:"total"`
            Succeed float64  `json:"succeed"`
            Stage string  `json:"stage"`
            ModifiedBy UserDetail1  `json:"modified_by"`
            Failed float64  `json:"failed"`
            IsActive bool  `json:"is_active"`
            CompanyID float64  `json:"company_id"`
            Cancelled float64  `json:"cancelled"`
            TemplateTag string  `json:"template_tag"`
            ModifiedOn string  `json:"modified_on"`
            FilePath string  `json:"file_path"`
            FailedRecords []string  `json:"failed_records"`
         
    }
    
    // ProductBulkRequestList used by Catalog
    type ProductBulkRequestList struct {

        
            Items ProductBulkRequest  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // BulkJob used by Catalog
    type BulkJob struct {

        
            CreatedOn string  `json:"created_on"`
            CancelledRecords []map[string]interface{}  `json:"cancelled_records"`
            Total float64  `json:"total"`
            Succeed float64  `json:"succeed"`
            TemplateTag string  `json:"template_tag"`
            Stage string  `json:"stage"`
            ModifiedBy UserInfo1  `json:"modified_by"`
            Failed float64  `json:"failed"`
            TrackingURL string  `json:"tracking_url"`
            IsActive bool  `json:"is_active"`
            CompanyID float64  `json:"company_id"`
            Cancelled float64  `json:"cancelled"`
            CreatedBy UserInfo1  `json:"created_by"`
            CustomTemplateTag string  `json:"custom_template_tag"`
            ModifiedOn string  `json:"modified_on"`
            FilePath string  `json:"file_path"`
            FailedRecords []map[string]interface{}  `json:"failed_records"`
         
    }
    
    // BulkResponse used by Catalog
    type BulkResponse struct {

        
            CreatedOn string  `json:"created_on"`
            BatchID string  `json:"batch_id"`
            ModifiedBy UserInfo1  `json:"modified_by"`
            IsActive bool  `json:"is_active"`
            CreatedBy UserInfo1  `json:"created_by"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // BulkProductRequest used by Catalog
    type BulkProductRequest struct {

        
            TemplateTag string  `json:"template_tag"`
            CompanyID float64  `json:"company_id"`
            BatchID string  `json:"batch_id"`
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

        
            CreatedOn string  `json:"created_on"`
            CancelledRecords []string  `json:"cancelled_records"`
            Retry float64  `json:"retry"`
            Total float64  `json:"total"`
            Succeed float64  `json:"succeed"`
            Stage string  `json:"stage"`
            ModifiedBy UserCommon  `json:"modified_by"`
            Failed float64  `json:"failed"`
            TrackingURL string  `json:"tracking_url"`
            IsActive bool  `json:"is_active"`
            CompanyID float64  `json:"company_id"`
            Cancelled float64  `json:"cancelled"`
            CreatedBy UserCommon  `json:"created_by"`
            ID string  `json:"id"`
            ModifiedOn string  `json:"modified_on"`
            FilePath string  `json:"file_path"`
            FailedRecords []string  `json:"failed_records"`
         
    }
    
    // BulkAssetResponse used by Catalog
    type BulkAssetResponse struct {

        
            Items []Items  `json:"items"`
            Page Page  `json:"page"`
         
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
            Size string  `json:"size"`
            CompanyID float64  `json:"company_id"`
         
    }
    
    // ProductSizeDeleteResponse used by Catalog
    type ProductSizeDeleteResponse struct {

        
            Success bool  `json:"success"`
            Data ProductSizeDeleteDataResponse  `json:"data"`
         
    }
    
    // InventoryResponse used by Catalog
    type InventoryResponse struct {

        
            ItemID float64  `json:"item_id"`
            Quantity float64  `json:"quantity"`
            PriceTransfer float64  `json:"price_transfer"`
            Identifiers map[string]interface{}  `json:"identifiers"`
            UID string  `json:"uid"`
            SellerIdentifier string  `json:"seller_identifier"`
            Store map[string]interface{}  `json:"store"`
            Price float64  `json:"price"`
            SellableQuantity float64  `json:"sellable_quantity"`
            Size string  `json:"size"`
            InventoryUpdatedOn string  `json:"inventory_updated_on"`
            Currency string  `json:"currency"`
            PriceEffective float64  `json:"price_effective"`
         
    }
    
    // InventoryResponsePaginated used by Catalog
    type InventoryResponsePaginated struct {

        
            Items []InventoryResponse  `json:"items"`
            Page Page  `json:"page"`
         
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

        
            Quantity float64  `json:"quantity"`
            SizeDistribution SizeDistribution  `json:"size_distribution"`
            Name string  `json:"name"`
         
    }
    
    // GTIN used by Catalog
    type GTIN struct {

        
            GtinType string  `json:"gtin_type"`
            GtinValue string  `json:"gtin_value"`
            Primary bool  `json:"primary"`
         
    }
    
    // InvSize used by Catalog
    type InvSize struct {

        
            ItemDimensionsUnitOfMeasure string  `json:"item_dimensions_unit_of_measure"`
            ItemWeight float64  `json:"item_weight"`
            StoreCode string  `json:"store_code"`
            PriceTransfer float64  `json:"price_transfer"`
            Set InventorySet  `json:"set"`
            ItemWidth float64  `json:"item_width"`
            Identifiers []GTIN  `json:"identifiers"`
            ItemWeightUnitOfMeasure string  `json:"item_weight_unit_of_measure"`
            Price float64  `json:"price"`
            ItemLength float64  `json:"item_length"`
            ExpirationDate string  `json:"expiration_date"`
            ItemHeight float64  `json:"item_height"`
            IsSet bool  `json:"is_set"`
            Size string  `json:"size"`
            Quantity float64  `json:"quantity"`
            Currency string  `json:"currency"`
            PriceEffective float64  `json:"price_effective"`
         
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
    
    // ReturnConfig1 used by Catalog
    type ReturnConfig1 struct {

        
            Time float64  `json:"time"`
            Unit string  `json:"unit"`
            Returnable bool  `json:"returnable"`
         
    }
    
    // QuantityBase used by Catalog
    type QuantityBase struct {

        
            Count float64  `json:"count"`
            UpdatedAt string  `json:"updated_at"`
         
    }
    
    // Quantities used by Catalog
    type Quantities struct {

        
            Sellable QuantityBase  `json:"sellable"`
            OrderCommitted QuantityBase  `json:"order_committed"`
            NotAvailable QuantityBase  `json:"not_available"`
            Damaged QuantityBase  `json:"damaged"`
         
    }
    
    // DimensionResponse used by Catalog
    type DimensionResponse struct {

        
            Width float64  `json:"width"`
            IsDefault bool  `json:"is_default"`
            Length float64  `json:"length"`
            Height float64  `json:"height"`
            Unit string  `json:"unit"`
         
    }
    
    // Trader1 used by Catalog
    type Trader1 struct {

        
            Address []string  `json:"address"`
            Name string  `json:"name"`
            Type string  `json:"type"`
         
    }
    
    // ManufacturerResponse used by Catalog
    type ManufacturerResponse struct {

        
            IsDefault bool  `json:"is_default"`
            Address string  `json:"address"`
            Name string  `json:"name"`
         
    }
    
    // WeightResponse used by Catalog
    type WeightResponse struct {

        
            IsDefault bool  `json:"is_default"`
            Unit string  `json:"unit"`
            Shipping float64  `json:"shipping"`
         
    }
    
    // StoreMeta used by Catalog
    type StoreMeta struct {

        
            ID float64  `json:"id"`
         
    }
    
    // BrandMeta used by Catalog
    type BrandMeta struct {

        
            ID float64  `json:"id"`
            Name string  `json:"name"`
         
    }
    
    // CompanyMeta used by Catalog
    type CompanyMeta struct {

        
            ID float64  `json:"id"`
         
    }
    
    // PriceMeta used by Catalog
    type PriceMeta struct {

        
            UpdatedAt string  `json:"updated_at"`
            Transfer float64  `json:"transfer"`
            Effective float64  `json:"effective"`
            Marked float64  `json:"marked"`
            TpNotes map[string]interface{}  `json:"tp_notes"`
            Currency string  `json:"currency"`
         
    }
    
    // InventorySellerResponse used by Catalog
    type InventorySellerResponse struct {

        
            CountryOfOrigin string  `json:"country_of_origin"`
            ReturnConfig ReturnConfig1  `json:"return_config"`
            Fragile bool  `json:"fragile"`
            CreatedBy UserSerializer  `json:"created_by"`
            TrackInventory bool  `json:"track_inventory"`
            FyndMeta map[string]interface{}  `json:"fynd_meta"`
            Quantities Quantities  `json:"quantities"`
            Set InventorySet  `json:"set"`
            UID string  `json:"uid"`
            Identifier map[string]interface{}  `json:"identifier"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            Tags []string  `json:"tags"`
            IsActive bool  `json:"is_active"`
            TraceID string  `json:"trace_id"`
            FyndArticleCode string  `json:"fynd_article_code"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Dimension DimensionResponse  `json:"dimension"`
            TotalQuantity float64  `json:"total_quantity"`
            Trader []Trader1  `json:"trader"`
            RawMeta map[string]interface{}  `json:"raw_meta"`
            Manufacturer ManufacturerResponse  `json:"manufacturer"`
            SellerIdentifier string  `json:"seller_identifier"`
            FyndItemCode string  `json:"fynd_item_code"`
            Stage string  `json:"stage"`
            Size string  `json:"size"`
            ItemID float64  `json:"item_id"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            Weight WeightResponse  `json:"weight"`
            Store StoreMeta  `json:"store"`
            Brand BrandMeta  `json:"brand"`
            Company CompanyMeta  `json:"company"`
            AddedOnStore string  `json:"added_on_store"`
            Price PriceMeta  `json:"price"`
            ExpirationDate string  `json:"expiration_date"`
            IsSet bool  `json:"is_set"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // InventorySellerIdentifierResponsePaginated used by Catalog
    type InventorySellerIdentifierResponsePaginated struct {

        
            Items []InventorySellerResponse  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // ReturnConfig2 used by Catalog
    type ReturnConfig2 struct {

        
            Time float64  `json:"time"`
            Unit string  `json:"unit"`
            Returnable bool  `json:"returnable"`
         
    }
    
    // DateMeta used by Catalog
    type DateMeta struct {

        
            InventoryUpdatedOn string  `json:"inventory_updated_on"`
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
            AddedOnStore string  `json:"added_on_store"`
         
    }
    
    // Quantity used by Catalog
    type Quantity struct {

        
            Count float64  `json:"count"`
         
    }
    
    // QuantitiesArticle used by Catalog
    type QuantitiesArticle struct {

        
            Sellable Quantity  `json:"sellable"`
            OrderCommitted Quantity  `json:"order_committed"`
            NotAvailable Quantity  `json:"not_available"`
            Damaged Quantity  `json:"damaged"`
         
    }
    
    // DimensionResponse1 used by Catalog
    type DimensionResponse1 struct {

        
            Length float64  `json:"length"`
            Width float64  `json:"width"`
            Unit string  `json:"unit"`
            Height float64  `json:"height"`
         
    }
    
    // Trader2 used by Catalog
    type Trader2 struct {

        
            Address []string  `json:"address"`
            Name string  `json:"name"`
            Type string  `json:"type"`
         
    }
    
    // ManufacturerResponse1 used by Catalog
    type ManufacturerResponse1 struct {

        
            IsDefault bool  `json:"is_default"`
            Address string  `json:"address"`
            Name string  `json:"name"`
         
    }
    
    // WeightResponse1 used by Catalog
    type WeightResponse1 struct {

        
            Unit string  `json:"unit"`
            Shipping float64  `json:"shipping"`
         
    }
    
    // ArticleStoreResponse used by Catalog
    type ArticleStoreResponse struct {

        
            StoreCode string  `json:"store_code"`
            UID float64  `json:"uid"`
            StoreType string  `json:"store_type"`
            Name string  `json:"name"`
         
    }
    
    // BrandMeta1 used by Catalog
    type BrandMeta1 struct {

        
            ID float64  `json:"id"`
            Name string  `json:"name"`
         
    }
    
    // CompanyMeta1 used by Catalog
    type CompanyMeta1 struct {

        
            ID float64  `json:"id"`
         
    }
    
    // PriceArticle used by Catalog
    type PriceArticle struct {

        
            Transfer float64  `json:"transfer"`
            Effective float64  `json:"effective"`
            Marked float64  `json:"marked"`
            TpNotes map[string]interface{}  `json:"tp_notes"`
            Currency string  `json:"currency"`
         
    }
    
    // GetInventories used by Catalog
    type GetInventories struct {

        
            CountryOfOrigin string  `json:"country_of_origin"`
            ReturnConfig ReturnConfig2  `json:"return_config"`
            DateMeta DateMeta  `json:"date_meta"`
            CreatedBy UserSerializer  `json:"created_by"`
            TrackInventory bool  `json:"track_inventory"`
            Quantities QuantitiesArticle  `json:"quantities"`
            UID string  `json:"uid"`
            Identifier map[string]interface{}  `json:"identifier"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            Tags []string  `json:"tags"`
            TraceID string  `json:"trace_id"`
            Platforms map[string]interface{}  `json:"platforms"`
            Dimension DimensionResponse1  `json:"dimension"`
            TotalQuantity float64  `json:"total_quantity"`
            Trader []Trader2  `json:"trader"`
            Manufacturer ManufacturerResponse1  `json:"manufacturer"`
            SellerIdentifier string  `json:"seller_identifier"`
            Stage string  `json:"stage"`
            Size string  `json:"size"`
            ID string  `json:"id"`
            InventoryUpdatedOn string  `json:"inventory_updated_on"`
            ItemID float64  `json:"item_id"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            Weight WeightResponse1  `json:"weight"`
            Store ArticleStoreResponse  `json:"store"`
            Brand BrandMeta1  `json:"brand"`
            Company CompanyMeta1  `json:"company"`
            Price PriceArticle  `json:"price"`
            ExpirationDate string  `json:"expiration_date"`
            IsSet bool  `json:"is_set"`
         
    }
    
    // GetInventoriesResponse used by Catalog
    type GetInventoriesResponse struct {

        
            Items []GetInventories  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // BulkInventoryGetItems used by Catalog
    type BulkInventoryGetItems struct {

        
            CreatedOn string  `json:"created_on"`
            CancelledRecords []string  `json:"cancelled_records"`
            Total float64  `json:"total"`
            Succeed float64  `json:"succeed"`
            Stage string  `json:"stage"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Failed float64  `json:"failed"`
            CompanyID float64  `json:"company_id"`
            IsActive bool  `json:"is_active"`
            Cancelled float64  `json:"cancelled"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ID string  `json:"id"`
            ModifiedOn string  `json:"modified_on"`
            FilePath string  `json:"file_path"`
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
            TotalQuantity float64  `json:"total_quantity"`
            ItemWeightUnitOfMeasure string  `json:"item_weight_unit_of_measure"`
            SellerIdentifier string  `json:"seller_identifier"`
            PriceMarked float64  `json:"price_marked"`
            ExpirationDate string  `json:"expiration_date"`
            Price float64  `json:"price"`
            Tags []string  `json:"tags"`
            TraceID string  `json:"trace_id"`
            Quantity float64  `json:"quantity"`
            Currency string  `json:"currency"`
            PriceEffective float64  `json:"price_effective"`
         
    }
    
    // InventoryBulkRequest used by Catalog
    type InventoryBulkRequest struct {

        
            Sizes []InventoryJobPayload  `json:"sizes"`
            User map[string]interface{}  `json:"user"`
            CompanyID float64  `json:"company_id"`
            BatchID string  `json:"batch_id"`
         
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
            ToDate string  `json:"to_date"`
            Quantity InventoryExportQuantityFilter  `json:"quantity"`
            FromDate string  `json:"from_date"`
         
    }
    
    // InventoryExportJob used by Catalog
    type InventoryExportJob struct {

        
            Status string  `json:"status"`
            CompletedOn string  `json:"completed_on"`
            Filters InventoryExportAdvanceOption  `json:"filters"`
            SellerID float64  `json:"seller_id"`
            TaskID string  `json:"task_id"`
            Type string  `json:"type"`
            URL string  `json:"url"`
            NotificationEmails []string  `json:"notification_emails"`
         
    }
    
    // InventoryExportRequest used by Catalog
    type InventoryExportRequest struct {

        
            Store []float64  `json:"store"`
            Brand []float64  `json:"brand"`
            Type string  `json:"type"`
         
    }
    
    // InventoryExportResponse used by Catalog
    type InventoryExportResponse struct {

        
            CreatedOn string  `json:"created_on"`
            Status string  `json:"status"`
            Type string  `json:"type"`
            Filters map[string]interface{}  `json:"filters"`
            SellerID float64  `json:"seller_id"`
            TaskID string  `json:"task_id"`
            NotificationEmails []string  `json:"notification_emails"`
            CreatedBy string  `json:"created_by"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // InventoryJobFilters used by Catalog
    type InventoryJobFilters struct {

        
            Brands []string  `json:"brands"`
            ToDate string  `json:"to_date"`
            Stores []string  `json:"stores"`
            Quantity InventoryExportQuantityFilter  `json:"quantity"`
            FromDate string  `json:"from_date"`
         
    }
    
    // InventoryJobDetailResponse used by Catalog
    type InventoryJobDetailResponse struct {

        
            CreatedOn string  `json:"created_on"`
            Status string  `json:"status"`
            CompletedOn string  `json:"completed_on"`
            Type string  `json:"type"`
            CancelledOn string  `json:"cancelled_on"`
            CancelledBy UserDetail  `json:"cancelled_by"`
            Filters InventoryJobFilters  `json:"filters"`
            SellerID float64  `json:"seller_id"`
            TaskID string  `json:"task_id"`
            NotificationEmails []string  `json:"notification_emails"`
            CreatedBy UserDetail  `json:"created_by"`
            ID string  `json:"id"`
            URL string  `json:"url"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // InventoryExportJobListResponse used by Catalog
    type InventoryExportJobListResponse struct {

        
            Items InventoryJobDetailResponse  `json:"items"`
         
    }
    
    // InventoryExportFilter used by Catalog
    type InventoryExportFilter struct {

        
            BrandIds []float64  `json:"brand_ids"`
            StoreIds []float64  `json:"store_ids"`
            ToDate string  `json:"to_date"`
            Quantity InventoryExportQuantityFilter  `json:"quantity"`
            FromDate string  `json:"from_date"`
         
    }
    
    // InventoryCreateRequest used by Catalog
    type InventoryCreateRequest struct {

        
            Filters InventoryExportFilter  `json:"filters"`
            NotificationEmails []string  `json:"notification_emails"`
            Data []string  `json:"data"`
            Type string  `json:"type"`
         
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
            StoreID float64  `json:"store_id"`
            SellerIdentifier string  `json:"seller_identifier"`
            PriceMarked float64  `json:"price_marked"`
            ExpirationDate string  `json:"expiration_date"`
            Tags []string  `json:"tags"`
            TraceID string  `json:"trace_id"`
            PriceEffective float64  `json:"price_effective"`
         
    }
    
    // InventoryRequestSchemaV2 used by Catalog
    type InventoryRequestSchemaV2 struct {

        
            Payload []InventoryPayload  `json:"payload"`
            CompanyID float64  `json:"company_id"`
            Meta map[string]interface{}  `json:"meta"`
         
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
    
    // HsnCodesObject used by Catalog
    type HsnCodesObject struct {

        
            Threshold2 float64  `json:"threshold2"`
            TaxOnEsp bool  `json:"tax_on_esp"`
            ModifiedOn string  `json:"modified_on"`
            TaxOnMrp bool  `json:"tax_on_mrp"`
            Tax2 float64  `json:"tax2"`
            CompanyID float64  `json:"company_id"`
            HsnCode string  `json:"hsn_code"`
            ID string  `json:"id"`
            Tax1 float64  `json:"tax1"`
            Hs2Code string  `json:"hs2_code"`
            Threshold1 float64  `json:"threshold1"`
         
    }
    
    // HsnCode used by Catalog
    type HsnCode struct {

        
            Data HsnCodesObject  `json:"data"`
         
    }
    
    // HsnUpsert used by Catalog
    type HsnUpsert struct {

        
            Threshold2 float64  `json:"threshold2"`
            TaxOnEsp bool  `json:"tax_on_esp"`
            UID float64  `json:"uid"`
            TaxOnMrp bool  `json:"tax_on_mrp"`
            Tax2 float64  `json:"tax2"`
            CompanyID float64  `json:"company_id"`
            HsnCode string  `json:"hsn_code"`
            IsActive bool  `json:"is_active"`
            Tax1 float64  `json:"tax1"`
            Hs2Code string  `json:"hs2_code"`
            Threshold1 float64  `json:"threshold1"`
         
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

        
            EffectiveDate string  `json:"effective_date"`
            Threshold float64  `json:"threshold"`
            Cess float64  `json:"cess"`
            Rate float64  `json:"rate"`
         
    }
    
    // HSNDataInsertV2 used by Catalog
    type HSNDataInsertV2 struct {

        
            Taxes []TaxSlab  `json:"taxes"`
            CreatedOn string  `json:"created_on"`
            Description string  `json:"description"`
            ReportingHsn string  `json:"reporting_hsn"`
            HsnCodeID string  `json:"hsn_code_id"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            CountryCode string  `json:"country_code"`
            HsnCode string  `json:"hsn_code"`
            Type string  `json:"type"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // PageResponse used by Catalog
    type PageResponse struct {

        
            ItemTotal float64  `json:"item_total"`
            HasPrevious bool  `json:"has_previous"`
            HasNext bool  `json:"has_next"`
            Size float64  `json:"size"`
            Current string  `json:"current"`
         
    }
    
    // HsnCodesListingResponseSchemaV2 used by Catalog
    type HsnCodesListingResponseSchemaV2 struct {

        
            Items []HSNDataInsertV2  `json:"items"`
            Page PageResponse  `json:"page"`
         
    }
    
    // BrandItem used by Catalog
    type BrandItem struct {

        
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            Departments []string  `json:"departments"`
            Logo Media  `json:"logo"`
            Banners ImageUrls  `json:"banners"`
            Discount string  `json:"discount"`
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
            Name string  `json:"name"`
            Logo Media  `json:"logo"`
         
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
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            Childs []map[string]interface{}  `json:"childs"`
            Banners ImageUrls  `json:"banners"`
            Action Action  `json:"action"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // SecondLevelChild used by Catalog
    type SecondLevelChild struct {

        
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            Childs []ThirdLevelChild  `json:"childs"`
            Banners ImageUrls  `json:"banners"`
            Action Action  `json:"action"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // Child used by Catalog
    type Child struct {

        
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            Childs []SecondLevelChild  `json:"childs"`
            Banners ImageUrls  `json:"banners"`
            Action Action  `json:"action"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // CategoryItems used by Catalog
    type CategoryItems struct {

        
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            Childs []Child  `json:"childs"`
            Banners ImageUrls  `json:"banners"`
            Action Action  `json:"action"`
         
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
            Items []ProductListingDetail  `json:"items"`
            SortOn []ProductSortOn  `json:"sort_on"`
            Operators map[string]interface{}  `json:"operators"`
         
    }
    
    // ProductDetail used by Catalog
    type ProductDetail struct {

        
            HasVariant bool  `json:"has_variant"`
            Color string  `json:"color"`
            Similars []string  `json:"similars"`
            UID float64  `json:"uid"`
            Rating float64  `json:"rating"`
            Highlights []string  `json:"highlights"`
            Type string  `json:"type"`
            RatingCount float64  `json:"rating_count"`
            Attributes map[string]interface{}  `json:"attributes"`
            Slug string  `json:"slug"`
            ShortDescription string  `json:"short_description"`
            ItemType string  `json:"item_type"`
            PromoMeta map[string]interface{}  `json:"promo_meta"`
            ItemCode string  `json:"item_code"`
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
            Tryouts []string  `json:"tryouts"`
            Description string  `json:"description"`
            ProductOnlineDate string  `json:"product_online_date"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            Name string  `json:"name"`
            Brand ProductBrand  `json:"brand"`
            Medias []Media1  `json:"medias"`
            ImageNature string  `json:"image_nature"`
         
    }
    
    // InventoryPage used by Catalog
    type InventoryPage struct {

        
            ItemTotal float64  `json:"item_total"`
            HasPrevious bool  `json:"has_previous"`
            Type string  `json:"type"`
            HasNext bool  `json:"has_next"`
            NextID string  `json:"next_id"`
         
    }
    
    // InventoryStockResponse used by Catalog
    type InventoryStockResponse struct {

        
            Items []map[string]interface{}  `json:"items"`
            Page InventoryPage  `json:"page"`
         
    }
    
    // ArticleAssignment used by Catalog
    type ArticleAssignment struct {

        
            Level string  `json:"level"`
            Strategy string  `json:"strategy"`
         
    }
    
    // ArticleQuery used by Catalog
    type ArticleQuery struct {

        
            Size string  `json:"size"`
            ItemID float64  `json:"item_id"`
            IgnoredStores []float64  `json:"ignored_stores"`
         
    }
    
    // AssignStoreArticle used by Catalog
    type AssignStoreArticle struct {

        
            ArticleAssignment ArticleAssignment  `json:"article_assignment"`
            GroupID string  `json:"group_id"`
            Quantity float64  `json:"quantity"`
            Query ArticleQuery  `json:"query"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // AssignStore used by Catalog
    type AssignStore struct {

        
            ChannelType string  `json:"channel_type"`
            Pincode string  `json:"pincode"`
            AppID string  `json:"app_id"`
            StoreIds []float64  `json:"store_ids"`
            CompanyID float64  `json:"company_id"`
            Articles []AssignStoreArticle  `json:"articles"`
            ChannelIdentifier string  `json:"channel_identifier"`
         
    }
    
    // ArticleAssignment1 used by Catalog
    type ArticleAssignment1 struct {

        
            Level string  `json:"level"`
            Strategy string  `json:"strategy"`
         
    }
    
    // StoreAssign used by Catalog
    type StoreAssign struct {

        
            ItemID float64  `json:"item_id"`
            Index float64  `json:"index"`
            Status bool  `json:"status"`
            UID string  `json:"uid"`
            StoreID float64  `json:"store_id"`
            PriceMarked float64  `json:"price_marked"`
            StrategyWiseListing []map[string]interface{}  `json:"strategy_wise_listing"`
            ArticleAssignment ArticleAssignment1  `json:"article_assignment"`
            CompanyID float64  `json:"company_id"`
            GroupID string  `json:"group_id"`
            Size string  `json:"size"`
            Quantity float64  `json:"quantity"`
            ID string  `json:"_id"`
            PriceEffective float64  `json:"price_effective"`
            StorePincode float64  `json:"store_pincode"`
            SCity string  `json:"s_city"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // StoreAssignError used by Catalog
    type StoreAssignError struct {

        
            Value interface{}  `json:"value"`
            Message string  `json:"message"`
            Type string  `json:"type"`
         
    }
    
    // StoreAssignResponse used by Catalog
    type StoreAssignResponse struct {

        
            Items []StoreAssign  `json:"items"`
            Success bool  `json:"success"`
            Error StoreAssignError  `json:"error"`
         
    }
    
    // GetAddressSerializer used by Catalog
    type GetAddressSerializer struct {

        
            Country string  `json:"country"`
            Longitude float64  `json:"longitude"`
            Address2 string  `json:"address2"`
            Pincode float64  `json:"pincode"`
            State string  `json:"state"`
            Landmark string  `json:"landmark"`
            CountryCode string  `json:"country_code"`
            Address1 string  `json:"address1"`
            Latitude float64  `json:"latitude"`
            City string  `json:"city"`
            AddressType string  `json:"address_type"`
         
    }
    
    // LocationTimingSerializer used by Catalog
    type LocationTimingSerializer struct {

        
            Hour float64  `json:"hour"`
            Minute float64  `json:"minute"`
         
    }
    
    // LocationDayWiseSerializer used by Catalog
    type LocationDayWiseSerializer struct {

        
            Closing LocationTimingSerializer  `json:"closing"`
            Open bool  `json:"open"`
            Weekday string  `json:"weekday"`
            Opening LocationTimingSerializer  `json:"opening"`
         
    }
    
    // UserSerializer1 used by Catalog
    type UserSerializer1 struct {

        
            UserID string  `json:"user_id"`
            Contact string  `json:"contact"`
            Username string  `json:"username"`
         
    }
    
    // ProductReturnConfigSerializer used by Catalog
    type ProductReturnConfigSerializer struct {

        
            StoreUID float64  `json:"store_uid"`
            OnSameStore bool  `json:"on_same_store"`
         
    }
    
    // InvoiceCredSerializer used by Catalog
    type InvoiceCredSerializer struct {

        
            Password string  `json:"password"`
            Enabled bool  `json:"enabled"`
            Username string  `json:"username"`
         
    }
    
    // InvoiceDetailsSerializer used by Catalog
    type InvoiceDetailsSerializer struct {

        
            EWaybill InvoiceCredSerializer  `json:"e_waybill"`
            EInvoice InvoiceCredSerializer  `json:"e_invoice"`
         
    }
    
    // Document used by Catalog
    type Document struct {

        
            LegalName string  `json:"legal_name"`
            Type string  `json:"type"`
            Verified bool  `json:"verified"`
            Value string  `json:"value"`
            URL string  `json:"url"`
         
    }
    
    // SellerPhoneNumber used by Catalog
    type SellerPhoneNumber struct {

        
            CountryCode float64  `json:"country_code"`
            Number string  `json:"number"`
         
    }
    
    // LocationIntegrationType used by Catalog
    type LocationIntegrationType struct {

        
            Inventory string  `json:"inventory"`
            Order string  `json:"order"`
         
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
    
    // GetCompanySerializer used by Catalog
    type GetCompanySerializer struct {

        
            CreatedOn string  `json:"created_on"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            RejectReason string  `json:"reject_reason"`
            VerifiedBy UserSerializer2  `json:"verified_by"`
            CompanyType string  `json:"company_type"`
            Stage string  `json:"stage"`
            ModifiedBy UserSerializer2  `json:"modified_by"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            CreatedBy UserSerializer2  `json:"created_by"`
            ModifiedOn string  `json:"modified_on"`
            VerifiedOn string  `json:"verified_on"`
            BusinessType string  `json:"business_type"`
         
    }
    
    // GetLocationSerializer used by Catalog
    type GetLocationSerializer struct {

        
            StoreType string  `json:"store_type"`
            Address GetAddressSerializer  `json:"address"`
            Timing []LocationDayWiseSerializer  `json:"timing"`
            PhoneNumber string  `json:"phone_number"`
            CreatedBy UserSerializer1  `json:"created_by"`
            DisplayName string  `json:"display_name"`
            VerifiedOn string  `json:"verified_on"`
            Warnings map[string]interface{}  `json:"warnings"`
            CreatedOn string  `json:"created_on"`
            UID float64  `json:"uid"`
            ProductReturnConfig ProductReturnConfigSerializer  `json:"product_return_config"`
            VerifiedBy UserSerializer1  `json:"verified_by"`
            GstCredentials InvoiceDetailsSerializer  `json:"gst_credentials"`
            ModifiedBy UserSerializer1  `json:"modified_by"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Documents []Document  `json:"documents"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            Stage string  `json:"stage"`
            IntegrationType LocationIntegrationType  `json:"integration_type"`
            Manager LocationManagerSerializer  `json:"manager"`
            ModifiedOn string  `json:"modified_on"`
            Name string  `json:"name"`
            Company GetCompanySerializer  `json:"company"`
            NotificationEmails []string  `json:"notification_emails"`
            Code string  `json:"code"`
         
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

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            Logo string  `json:"logo"`
            AppID string  `json:"app_id"`
            IsActive bool  `json:"is_active"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
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
    

    
    // Document used by CompanyProfile
    type Document struct {

        
            Type string  `json:"type"`
            Verified bool  `json:"verified"`
            Value string  `json:"value"`
            URL string  `json:"url"`
            LegalName string  `json:"legal_name"`
         
    }
    
    // CompanyTaxesSerializer used by CompanyProfile
    type CompanyTaxesSerializer struct {

        
            Enable bool  `json:"enable"`
            EffectiveDate string  `json:"effective_date"`
            Rate float64  `json:"rate"`
         
    }
    
    // GetAddressSerializer used by CompanyProfile
    type GetAddressSerializer struct {

        
            Latitude float64  `json:"latitude"`
            Landmark string  `json:"landmark"`
            AddressType string  `json:"address_type"`
            Address1 string  `json:"address1"`
            Longitude float64  `json:"longitude"`
            State string  `json:"state"`
            Country string  `json:"country"`
            Pincode float64  `json:"pincode"`
            CountryCode string  `json:"country_code"`
            City string  `json:"city"`
            Address2 string  `json:"address2"`
         
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
    
    // UserSerializer used by CompanyProfile
    type UserSerializer struct {

        
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
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
    
    // BusinessCountryInfo used by CompanyProfile
    type BusinessCountryInfo struct {

        
            Country string  `json:"country"`
            CountryCode string  `json:"country_code"`
         
    }
    
    // GetCompanyProfileSerializerResponse used by CompanyProfile
    type GetCompanyProfileSerializerResponse struct {

        
            Documents []Document  `json:"documents"`
            Taxes []CompanyTaxesSerializer  `json:"taxes"`
            FranchiseEnabled bool  `json:"franchise_enabled"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            Name string  `json:"name"`
            BusinessInfo string  `json:"business_info"`
            CompanyType string  `json:"company_type"`
            Warnings map[string]interface{}  `json:"warnings"`
            ContactDetails ContactDetails  `json:"contact_details"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            NotificationEmails []string  `json:"notification_emails"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            UID float64  `json:"uid"`
            CreatedOn string  `json:"created_on"`
            Stage string  `json:"stage"`
            BusinessType string  `json:"business_type"`
            VerifiedOn string  `json:"verified_on"`
            ModifiedOn string  `json:"modified_on"`
            BusinessDetails BusinessDetails  `json:"business_details"`
            BusinessCountryInfo BusinessCountryInfo  `json:"business_country_info"`
            Mode string  `json:"mode"`
            CreatedBy UserSerializer  `json:"created_by"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // ErrorResponse used by CompanyProfile
    type ErrorResponse struct {

        
            Meta map[string]interface{}  `json:"meta"`
            Status float64  `json:"status"`
            Message string  `json:"message"`
            Code string  `json:"code"`
         
    }
    
    // CompanyTaxesSerializer1 used by CompanyProfile
    type CompanyTaxesSerializer1 struct {

        
            Enable bool  `json:"enable"`
            EffectiveDate string  `json:"effective_date"`
            Rate float64  `json:"rate"`
         
    }
    
    // CreateUpdateAddressSerializer used by CompanyProfile
    type CreateUpdateAddressSerializer struct {

        
            Latitude float64  `json:"latitude"`
            Landmark string  `json:"landmark"`
            Address1 string  `json:"address1"`
            AddressType string  `json:"address_type"`
            Longitude float64  `json:"longitude"`
            State string  `json:"state"`
            Country string  `json:"country"`
            Pincode float64  `json:"pincode"`
            CountryCode string  `json:"country_code"`
            City string  `json:"city"`
            Address2 string  `json:"address2"`
         
    }
    
    // UpdateCompany used by CompanyProfile
    type UpdateCompany struct {

        
            NotificationEmails []string  `json:"notification_emails"`
            Name string  `json:"name"`
            Documents []Document  `json:"documents"`
            BusinessDetails BusinessDetails  `json:"business_details"`
            BusinessInfo string  `json:"business_info"`
            Taxes []CompanyTaxesSerializer1  `json:"taxes"`
            CompanyType string  `json:"company_type"`
            RejectReason string  `json:"reject_reason"`
            Warnings map[string]interface{}  `json:"warnings"`
            ContactDetails ContactDetails  `json:"contact_details"`
            FranchiseEnabled bool  `json:"franchise_enabled"`
            Addresses []CreateUpdateAddressSerializer  `json:"addresses"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            BusinessType string  `json:"business_type"`
         
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
            Product DocumentsObj  `json:"product"`
            Brand DocumentsObj  `json:"brand"`
            CompanyDocuments DocumentsObj  `json:"company_documents"`
            UID float64  `json:"uid"`
            Stage string  `json:"stage"`
         
    }
    
    // BrandBannerSerializer used by CompanyProfile
    type BrandBannerSerializer struct {

        
            Portrait string  `json:"portrait"`
            Landscape string  `json:"landscape"`
         
    }
    
    // GetBrandResponseSerializer used by CompanyProfile
    type GetBrandResponseSerializer struct {

        
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            SlugKey string  `json:"slug_key"`
            RejectReason string  `json:"reject_reason"`
            Description string  `json:"description"`
            Name string  `json:"name"`
            Warnings map[string]interface{}  `json:"warnings"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            Banner BrandBannerSerializer  `json:"banner"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            UID float64  `json:"uid"`
            CreatedOn string  `json:"created_on"`
            Synonyms []string  `json:"synonyms"`
            Logo string  `json:"logo"`
            Stage string  `json:"stage"`
            VerifiedOn string  `json:"verified_on"`
            ModifiedOn string  `json:"modified_on"`
            Mode string  `json:"mode"`
            CreatedBy UserSerializer  `json:"created_by"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // CreateUpdateBrandRequestSerializer used by CompanyProfile
    type CreateUpdateBrandRequestSerializer struct {

        
            Name string  `json:"name"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            Synonyms []string  `json:"synonyms"`
            Banner BrandBannerSerializer  `json:"banner"`
            BrandTier string  `json:"brand_tier"`
            UID float64  `json:"uid"`
            CompanyID float64  `json:"company_id"`
            Description string  `json:"description"`
            Logo string  `json:"logo"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
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

        
            VerifiedBy UserSerializer  `json:"verified_by"`
            Details CompanyDetails  `json:"details"`
            NotificationEmails []string  `json:"notification_emails"`
            Name string  `json:"name"`
            BusinessType string  `json:"business_type"`
            BusinessCountryInfo BusinessCountryInfo  `json:"business_country_info"`
            CompanyType string  `json:"company_type"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            RejectReason string  `json:"reject_reason"`
            MarketChannels []string  `json:"market_channels"`
            UID float64  `json:"uid"`
            CreatedBy UserSerializer  `json:"created_by"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            CreatedOn string  `json:"created_on"`
            Stage string  `json:"stage"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            VerifiedOn string  `json:"verified_on"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // CompanyBrandSerializer used by CompanyProfile
    type CompanyBrandSerializer struct {

        
            VerifiedBy UserSerializer  `json:"verified_by"`
            Brand GetBrandResponseSerializer  `json:"brand"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            RejectReason string  `json:"reject_reason"`
            Company CompanySerializer  `json:"company"`
            Warnings map[string]interface{}  `json:"warnings"`
            UID float64  `json:"uid"`
            CreatedBy UserSerializer  `json:"created_by"`
            CreatedOn string  `json:"created_on"`
            Stage string  `json:"stage"`
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
            Brands []float64  `json:"brands"`
            Company float64  `json:"company"`
         
    }
    
    // GetCompanySerializer used by CompanyProfile
    type GetCompanySerializer struct {

        
            VerifiedBy UserSerializer  `json:"verified_by"`
            Name string  `json:"name"`
            BusinessType string  `json:"business_type"`
            CompanyType string  `json:"company_type"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            RejectReason string  `json:"reject_reason"`
            UID float64  `json:"uid"`
            CreatedBy UserSerializer  `json:"created_by"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            CreatedOn string  `json:"created_on"`
            Stage string  `json:"stage"`
            VerifiedOn string  `json:"verified_on"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // InvoiceCredSerializer used by CompanyProfile
    type InvoiceCredSerializer struct {

        
            Username string  `json:"username"`
            Password string  `json:"password"`
            Enabled bool  `json:"enabled"`
         
    }
    
    // InvoiceDetailsSerializer used by CompanyProfile
    type InvoiceDetailsSerializer struct {

        
            EInvoice InvoiceCredSerializer  `json:"e_invoice"`
            EWaybill InvoiceCredSerializer  `json:"e_waybill"`
         
    }
    
    // LocationTimingSerializer used by CompanyProfile
    type LocationTimingSerializer struct {

        
            Hour float64  `json:"hour"`
            Minute float64  `json:"minute"`
         
    }
    
    // LocationDayWiseSerializer used by CompanyProfile
    type LocationDayWiseSerializer struct {

        
            Open bool  `json:"open"`
            Weekday string  `json:"weekday"`
            Opening LocationTimingSerializer  `json:"opening"`
            Closing LocationTimingSerializer  `json:"closing"`
         
    }
    
    // LocationManagerSerializer used by CompanyProfile
    type LocationManagerSerializer struct {

        
            MobileNo SellerPhoneNumber  `json:"mobile_no"`
            Email string  `json:"email"`
            Name string  `json:"name"`
         
    }
    
    // HolidayDateSerializer used by CompanyProfile
    type HolidayDateSerializer struct {

        
            StartDate string  `json:"start_date"`
            EndDate string  `json:"end_date"`
         
    }
    
    // HolidaySchemaSerializer used by CompanyProfile
    type HolidaySchemaSerializer struct {

        
            Title string  `json:"title"`
            Date HolidayDateSerializer  `json:"date"`
            HolidayType string  `json:"holiday_type"`
         
    }
    
    // ProductReturnConfigSerializer used by CompanyProfile
    type ProductReturnConfigSerializer struct {

        
            OnSameStore bool  `json:"on_same_store"`
            StoreUID float64  `json:"store_uid"`
         
    }
    
    // GetLocationSerializer used by CompanyProfile
    type GetLocationSerializer struct {

        
            Documents []Document  `json:"documents"`
            Company GetCompanySerializer  `json:"company"`
            DisplayName string  `json:"display_name"`
            GstCredentials InvoiceDetailsSerializer  `json:"gst_credentials"`
            Name string  `json:"name"`
            Code string  `json:"code"`
            Timing []LocationDayWiseSerializer  `json:"timing"`
            Warnings map[string]interface{}  `json:"warnings"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            Manager LocationManagerSerializer  `json:"manager"`
            NotificationEmails []string  `json:"notification_emails"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            StoreType string  `json:"store_type"`
            Holiday []HolidaySchemaSerializer  `json:"holiday"`
            UID float64  `json:"uid"`
            CreatedOn string  `json:"created_on"`
            ProductReturnConfig ProductReturnConfigSerializer  `json:"product_return_config"`
            Stage string  `json:"stage"`
            VerifiedOn string  `json:"verified_on"`
            ModifiedOn string  `json:"modified_on"`
            PhoneNumber string  `json:"phone_number"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            CreatedBy UserSerializer  `json:"created_by"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Address GetAddressSerializer  `json:"address"`
         
    }
    
    // LocationListSerializer used by CompanyProfile
    type LocationListSerializer struct {

        
            Items []GetLocationSerializer  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // AddressSerializer used by CompanyProfile
    type AddressSerializer struct {

        
            Latitude float64  `json:"latitude"`
            Landmark string  `json:"landmark"`
            AddressType string  `json:"address_type"`
            Address1 string  `json:"address1"`
            Longitude float64  `json:"longitude"`
            State string  `json:"state"`
            Country string  `json:"country"`
            Pincode float64  `json:"pincode"`
            CountryCode string  `json:"country_code"`
            City string  `json:"city"`
            Address2 string  `json:"address2"`
         
    }
    
    // LocationSerializer used by CompanyProfile
    type LocationSerializer struct {

        
            Documents []Document  `json:"documents"`
            Name string  `json:"name"`
            Manager LocationManagerSerializer  `json:"manager"`
            NotificationEmails []string  `json:"notification_emails"`
            Code string  `json:"code"`
            Timing []LocationDayWiseSerializer  `json:"timing"`
            Company float64  `json:"company"`
            StoreType string  `json:"store_type"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            DisplayName string  `json:"display_name"`
            UID float64  `json:"uid"`
            Holiday []HolidaySchemaSerializer  `json:"holiday"`
            Warnings map[string]interface{}  `json:"warnings"`
            ProductReturnConfig ProductReturnConfigSerializer  `json:"product_return_config"`
            GstCredentials InvoiceDetailsSerializer  `json:"gst_credentials"`
            Stage string  `json:"stage"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Address AddressSerializer  `json:"address"`
         
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
    
    // ValiditySchema used by Cart
    type ValiditySchema struct {

        
            Priority float64  `json:"priority"`
         
    }
    
    // StateSchema used by Cart
    type StateSchema struct {

        
            IsDisplay bool  `json:"is_display"`
            IsArchived bool  `json:"is_archived"`
            IsPublic bool  `json:"is_public"`
         
    }
    
    // CouponActionSchema used by Cart
    type CouponActionSchema struct {

        
            ActionDate string  `json:"action_date"`
            TxnMode string  `json:"txn_mode"`
         
    }
    
    // IdentifierSchema used by Cart
    type IdentifierSchema struct {

        
            CompanyID []float64  `json:"company_id"`
            ExcludeBrandID []float64  `json:"exclude_brand_id"`
            ItemID []float64  `json:"item_id"`
            BrandID []float64  `json:"brand_id"`
            CollectionID []string  `json:"collection_id"`
            UserID []string  `json:"user_id"`
            StoreID []float64  `json:"store_id"`
            EmailDomain []string  `json:"email_domain"`
            CategoryID []float64  `json:"category_id"`
            ArticleID []string  `json:"article_id"`
         
    }
    
    // RuleDefinitionSchema used by Cart
    type RuleDefinitionSchema struct {

        
            CalculateOn string  `json:"calculate_on"`
            CurrencyCode string  `json:"currency_code"`
            IsExact bool  `json:"is_exact"`
            AutoApply bool  `json:"auto_apply"`
            Scope []string  `json:"scope"`
            Type string  `json:"type"`
            ApplicableOn string  `json:"applicable_on"`
            ValueType string  `json:"value_type"`
         
    }
    
    // DisplayMetaDictSchema used by Cart
    type DisplayMetaDictSchema struct {

        
            Title string  `json:"title"`
            Subtitle string  `json:"subtitle"`
         
    }
    
    // DisplayMetaSchema used by Cart
    type DisplayMetaSchema struct {

        
            Description string  `json:"description"`
            Title string  `json:"title"`
            Apply DisplayMetaDictSchema  `json:"apply"`
            Auto DisplayMetaDictSchema  `json:"auto"`
            Subtitle string  `json:"subtitle"`
            Remove DisplayMetaDictSchema  `json:"remove"`
         
    }
    
    // CouponAuthorSchema used by Cart
    type CouponAuthorSchema struct {

        
            ModifiedBy string  `json:"modified_by"`
            CreatedBy string  `json:"created_by"`
         
    }
    
    // BulkBundleRestriction used by Cart
    type BulkBundleRestriction struct {

        
            MultiStoreAllowed bool  `json:"multi_store_allowed"`
         
    }
    
    // PriceRangeSchema used by Cart
    type PriceRangeSchema struct {

        
            Min float64  `json:"min"`
            Max float64  `json:"max"`
         
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

        
            Codes []string  `json:"codes"`
            Types []string  `json:"types"`
            Uses PaymentAllowValueSchema  `json:"uses"`
            Networks []string  `json:"networks"`
         
    }
    
    // UsesRemainingSchema used by Cart
    type UsesRemainingSchema struct {

        
            Total float64  `json:"total"`
            App float64  `json:"app"`
            User float64  `json:"user"`
         
    }
    
    // UsesRestrictionSchema used by Cart
    type UsesRestrictionSchema struct {

        
            Remaining UsesRemainingSchema  `json:"remaining"`
            Maximum UsesRemainingSchema  `json:"maximum"`
         
    }
    
    // RestrictionsSchema used by Cart
    type RestrictionsSchema struct {

        
            BulkBundle BulkBundleRestriction  `json:"bulk_bundle"`
            PriceRange PriceRangeSchema  `json:"price_range"`
            UserGroups []float64  `json:"user_groups"`
            OrderingStores []float64  `json:"ordering_stores"`
            Platforms []string  `json:"platforms"`
            PostOrder PostOrderSchema  `json:"post_order"`
            Payments map[string]PaymentModesSchema  `json:"payments"`
            Uses UsesRestrictionSchema  `json:"uses"`
            CouponAllowed bool  `json:"coupon_allowed"`
            UserType string  `json:"user_type"`
         
    }
    
    // ValidationSchema used by Cart
    type ValidationSchema struct {

        
            UserRegisteredAfter string  `json:"user_registered_after"`
            AppID []string  `json:"app_id"`
            Anonymous bool  `json:"anonymous"`
         
    }
    
    // OwnershipSchema used by Cart
    type OwnershipSchema struct {

        
            PayableCategory string  `json:"payable_category"`
            PayableBy string  `json:"payable_by"`
         
    }
    
    // CouponScheduleSchema used by Cart
    type CouponScheduleSchema struct {

        
            NextSchedule []map[string]interface{}  `json:"next_schedule"`
            End string  `json:"end"`
            Duration float64  `json:"duration"`
            Start string  `json:"start"`
            Cron string  `json:"cron"`
         
    }
    
    // RuleSchema used by Cart
    type RuleSchema struct {

        
            Min float64  `json:"min"`
            Max float64  `json:"max"`
            Key float64  `json:"key"`
            Value float64  `json:"value"`
            DiscountQty float64  `json:"discount_qty"`
         
    }
    
    // CouponDateMetaSchema used by Cart
    type CouponDateMetaSchema struct {

        
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
         
    }
    
    // CouponAddSchema used by Cart
    type CouponAddSchema struct {

        
            Validity ValiditySchema  `json:"validity"`
            State StateSchema  `json:"state"`
            Action CouponActionSchema  `json:"action"`
            Identifiers IdentifierSchema  `json:"identifiers"`
            RuleDefinition RuleDefinitionSchema  `json:"rule_definition"`
            DisplayMeta DisplayMetaSchema  `json:"display_meta"`
            Author CouponAuthorSchema  `json:"author"`
            Restrictions RestrictionsSchema  `json:"restrictions"`
            Validation ValidationSchema  `json:"validation"`
            Ownership OwnershipSchema  `json:"ownership"`
            Schedule CouponScheduleSchema  `json:"_schedule"`
            Tags []string  `json:"tags"`
            Rule []RuleSchema  `json:"rule"`
            TypeSlug string  `json:"type_slug"`
            Code string  `json:"code"`
            DateMeta CouponDateMetaSchema  `json:"date_meta"`
         
    }
    
    // CouponsResponse used by Cart
    type CouponsResponse struct {

        
            Page Page  `json:"page"`
            Items CouponAddSchema  `json:"items"`
         
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
    
    // CouponUpdateSchema used by Cart
    type CouponUpdateSchema struct {

        
            Validity ValiditySchema  `json:"validity"`
            State StateSchema  `json:"state"`
            Action CouponActionSchema  `json:"action"`
            Identifiers IdentifierSchema  `json:"identifiers"`
            RuleDefinition RuleDefinitionSchema  `json:"rule_definition"`
            DisplayMeta DisplayMetaSchema  `json:"display_meta"`
            Author CouponAuthorSchema  `json:"author"`
            Restrictions RestrictionsSchema  `json:"restrictions"`
            Validation ValidationSchema  `json:"validation"`
            Ownership OwnershipSchema  `json:"ownership"`
            Schedule CouponScheduleSchema  `json:"_schedule"`
            Tags []string  `json:"tags"`
            Rule []RuleSchema  `json:"rule"`
            TypeSlug string  `json:"type_slug"`
            Code string  `json:"code"`
            DateMeta CouponDateMetaSchema  `json:"date_meta"`
         
    }
    
    // CouponPartialUpdate used by Cart
    type CouponPartialUpdate struct {

        
            Archive bool  `json:"archive"`
            Schedule CouponScheduleSchema  `json:"schedule"`
         
    }
    
    // DiscountOfferSchema used by Cart
    type DiscountOfferSchema struct {

        
            MinOfferQuantity float64  `json:"min_offer_quantity"`
            DiscountPercentage float64  `json:"discount_percentage"`
            MaxDiscountAmount float64  `json:"max_discount_amount"`
            MaxOfferQuantity float64  `json:"max_offer_quantity"`
            DiscountAmount float64  `json:"discount_amount"`
            ApportionDiscount bool  `json:"apportion_discount"`
            DiscountPrice float64  `json:"discount_price"`
            Code string  `json:"code"`
            PartialCanRet bool  `json:"partial_can_ret"`
            MaxUsagePerTransaction float64  `json:"max_usage_per_transaction"`
         
    }
    
    // CompareObjectSchema used by Cart
    type CompareObjectSchema struct {

        
            LessThan float64  `json:"less_than"`
            LessThanEquals float64  `json:"less_than_equals"`
            GreaterThan float64  `json:"greater_than"`
            GreaterThanEquals float64  `json:"greater_than_equals"`
            Equals float64  `json:"equals"`
         
    }
    
    // ItemCriteriaSchema used by Cart
    type ItemCriteriaSchema struct {

        
            AllItems bool  `json:"all_items"`
            ItemExcludeCategory []float64  `json:"item_exclude_category"`
            ProductTags []string  `json:"product_tags"`
            ItemSize []string  `json:"item_size"`
            ItemExcludeL1Category []float64  `json:"item_exclude_l1_category"`
            ItemSku []string  `json:"item_sku"`
            ItemDepartment []float64  `json:"item_department"`
            ItemTags []string  `json:"item_tags"`
            ItemExcludeCompany []float64  `json:"item_exclude_company"`
            ItemStore []float64  `json:"item_store"`
            ItemExcludeStore []float64  `json:"item_exclude_store"`
            ItemExcludeDepartment []float64  `json:"item_exclude_department"`
            ItemExcludeID []float64  `json:"item_exclude_id"`
            ItemCategory []float64  `json:"item_category"`
            ItemExcludeL2Category []float64  `json:"item_exclude_l2_category"`
            ItemExcludeBrand []float64  `json:"item_exclude_brand"`
            AvailableZones []string  `json:"available_zones"`
            CartQuantity CompareObjectSchema  `json:"cart_quantity"`
            BuyRules []string  `json:"buy_rules"`
            CartUniqueItemQuantity CompareObjectSchema  `json:"cart_unique_item_quantity"`
            ItemL1Category []float64  `json:"item_l1_category"`
            ItemBrand []float64  `json:"item_brand"`
            ItemID []float64  `json:"item_id"`
            CartUniqueItemAmount CompareObjectSchema  `json:"cart_unique_item_amount"`
            ItemL2Category []float64  `json:"item_l2_category"`
            ItemCompany []float64  `json:"item_company"`
            CartTotal CompareObjectSchema  `json:"cart_total"`
            ItemExcludeSku []string  `json:"item_exclude_sku"`
         
    }
    
    // DiscountRuleSchema used by Cart
    type DiscountRuleSchema struct {

        
            DiscountType string  `json:"discount_type"`
            BuyCondition string  `json:"buy_condition"`
            Offer DiscountOfferSchema  `json:"offer"`
            ItemCriteria ItemCriteriaSchema  `json:"item_criteria"`
         
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

        
            Codes []string  `json:"codes"`
            Uses PaymentAllowValueSchema1  `json:"uses"`
            Type string  `json:"type"`
         
    }
    
    // UsesRemainingSchema1 used by Cart
    type UsesRemainingSchema1 struct {

        
            Total float64  `json:"total"`
            User float64  `json:"user"`
         
    }
    
    // UsesRestrictionSchema1 used by Cart
    type UsesRestrictionSchema1 struct {

        
            Remaining UsesRemainingSchema1  `json:"remaining"`
            Maximum UsesRemainingSchema1  `json:"maximum"`
         
    }
    
    // UserRegisteredSchema used by Cart
    type UserRegisteredSchema struct {

        
            End string  `json:"end"`
            Start string  `json:"start"`
         
    }
    
    // RestrictionsSchema1 used by Cart
    type RestrictionsSchema1 struct {

        
            UserGroups []float64  `json:"user_groups"`
            Platforms []string  `json:"platforms"`
            PostOrder PostOrderSchema1  `json:"post_order"`
            UserID []string  `json:"user_id"`
            Payments []PromotionPaymentModesSchema  `json:"payments"`
            Uses UsesRestrictionSchema1  `json:"uses"`
            AnonymousUsers bool  `json:"anonymous_users"`
            OrderQuantity float64  `json:"order_quantity"`
            UserRegistered UserRegisteredSchema  `json:"user_registered"`
         
    }
    
    // PromotionAuthorSchema used by Cart
    type PromotionAuthorSchema struct {

        
            ModifiedBy string  `json:"modified_by"`
            CreatedBy string  `json:"created_by"`
         
    }
    
    // OwnershipSchema1 used by Cart
    type OwnershipSchema1 struct {

        
            PayableCategory string  `json:"payable_category"`
            PayableBy string  `json:"payable_by"`
         
    }
    
    // VisibilitySchema used by Cart
    type VisibilitySchema struct {

        
            Pdp bool  `json:"pdp"`
            CouponList bool  `json:"coupon_list"`
         
    }
    
    // PromotionDateMetaSchema used by Cart
    type PromotionDateMetaSchema struct {

        
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
         
    }
    
    // DisplayMetaSchema1 used by Cart
    type DisplayMetaSchema1 struct {

        
            OfferLabel string  `json:"offer_label"`
            OfferText string  `json:"offer_text"`
            Description string  `json:"description"`
            Name string  `json:"name"`
         
    }
    
    // PromotionScheduleSchema used by Cart
    type PromotionScheduleSchema struct {

        
            NextSchedule []map[string]interface{}  `json:"next_schedule"`
            End string  `json:"end"`
            Duration float64  `json:"duration"`
            Start string  `json:"start"`
            Cron string  `json:"cron"`
            Published bool  `json:"published"`
         
    }
    
    // PromotionActionSchema used by Cart
    type PromotionActionSchema struct {

        
            ActionDate string  `json:"action_date"`
            ActionType string  `json:"action_type"`
         
    }
    
    // PromotionListItem used by Cart
    type PromotionListItem struct {

        
            DiscountRules []DiscountRuleSchema  `json:"discount_rules"`
            Restrictions RestrictionsSchema1  `json:"restrictions"`
            Code string  `json:"code"`
            PromoGroup string  `json:"promo_group"`
            Currency string  `json:"currency"`
            PromotionType string  `json:"promotion_type"`
            ApplyExclusive string  `json:"apply_exclusive"`
            ApplyPriority float64  `json:"apply_priority"`
            ApplyAllDiscount bool  `json:"apply_all_discount"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Author PromotionAuthorSchema  `json:"author"`
            Mode string  `json:"mode"`
            Ownership OwnershipSchema1  `json:"ownership"`
            Visiblility VisibilitySchema  `json:"visiblility"`
            ApplicationID string  `json:"application_id"`
            DateMeta PromotionDateMetaSchema  `json:"date_meta"`
            Stackable bool  `json:"stackable"`
            CalculateOn string  `json:"calculate_on"`
            BuyRules map[string]ItemCriteriaSchema  `json:"buy_rules"`
            DisplayMeta DisplayMetaSchema1  `json:"display_meta"`
            Schedule PromotionScheduleSchema  `json:"_schedule"`
            PostOrderAction PromotionActionSchema  `json:"post_order_action"`
         
    }
    
    // PromotionsResponse used by Cart
    type PromotionsResponse struct {

        
            Page Page  `json:"page"`
            Items PromotionListItem  `json:"items"`
         
    }
    
    // PromotionAddSchema used by Cart
    type PromotionAddSchema struct {

        
            DiscountRules []DiscountRuleSchema  `json:"discount_rules"`
            Restrictions RestrictionsSchema1  `json:"restrictions"`
            Code string  `json:"code"`
            PromoGroup string  `json:"promo_group"`
            Currency string  `json:"currency"`
            PromotionType string  `json:"promotion_type"`
            ApplyExclusive string  `json:"apply_exclusive"`
            ApplyPriority float64  `json:"apply_priority"`
            ApplyAllDiscount bool  `json:"apply_all_discount"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Author PromotionAuthorSchema  `json:"author"`
            Mode string  `json:"mode"`
            Ownership OwnershipSchema1  `json:"ownership"`
            Visiblility VisibilitySchema  `json:"visiblility"`
            ApplicationID string  `json:"application_id"`
            DateMeta PromotionDateMetaSchema  `json:"date_meta"`
            Stackable bool  `json:"stackable"`
            CalculateOn string  `json:"calculate_on"`
            BuyRules map[string]ItemCriteriaSchema  `json:"buy_rules"`
            DisplayMeta DisplayMetaSchema1  `json:"display_meta"`
            Schedule PromotionScheduleSchema  `json:"_schedule"`
            PostOrderAction PromotionActionSchema  `json:"post_order_action"`
         
    }
    
    // PromotionUpdateSchema used by Cart
    type PromotionUpdateSchema struct {

        
            DiscountRules []DiscountRuleSchema  `json:"discount_rules"`
            Restrictions RestrictionsSchema1  `json:"restrictions"`
            Code string  `json:"code"`
            PromoGroup string  `json:"promo_group"`
            Currency string  `json:"currency"`
            PromotionType string  `json:"promotion_type"`
            ApplyExclusive string  `json:"apply_exclusive"`
            ApplyPriority float64  `json:"apply_priority"`
            ApplyAllDiscount bool  `json:"apply_all_discount"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Author PromotionAuthorSchema  `json:"author"`
            Mode string  `json:"mode"`
            Ownership OwnershipSchema1  `json:"ownership"`
            Visiblility VisibilitySchema  `json:"visiblility"`
            ApplicationID string  `json:"application_id"`
            DateMeta PromotionDateMetaSchema  `json:"date_meta"`
            Stackable bool  `json:"stackable"`
            CalculateOn string  `json:"calculate_on"`
            BuyRules map[string]ItemCriteriaSchema  `json:"buy_rules"`
            DisplayMeta DisplayMetaSchema1  `json:"display_meta"`
            Schedule PromotionScheduleSchema  `json:"_schedule"`
            PostOrderAction PromotionActionSchema  `json:"post_order_action"`
         
    }
    
    // PromotionPartialUpdate used by Cart
    type PromotionPartialUpdate struct {

        
            Archive bool  `json:"archive"`
            Schedule PromotionScheduleSchema  `json:"schedule"`
         
    }
    
    // ActivePromosResponse used by Cart
    type ActivePromosResponse struct {

        
            IsHidden bool  `json:"is_hidden"`
            Example string  `json:"example"`
            Description string  `json:"description"`
            EntityType string  `json:"entity_type"`
            Title string  `json:"title"`
            Type string  `json:"type"`
            CreatedOn string  `json:"created_on"`
            EntitySlug string  `json:"entity_slug"`
            Subtitle string  `json:"subtitle"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // CartItemSchema used by Cart
    type CartItemSchema struct {

        
            Size string  `json:"size"`
            Quantity float64  `json:"quantity"`
            ProductID string  `json:"product_id"`
         
    }
    
    // OpenapiCartDetailsRequest used by Cart
    type OpenapiCartDetailsRequest struct {

        
            CartItems []CartItemSchema  `json:"cart_items"`
         
    }
    
    // CouponBreakupSchema used by Cart
    type CouponBreakupSchema struct {

        
            UID string  `json:"uid"`
            MaxDiscountValue float64  `json:"max_discount_value"`
            CouponType string  `json:"coupon_type"`
            Message string  `json:"message"`
            Description string  `json:"description"`
            MinimumCartValue float64  `json:"minimum_cart_value"`
            CouponValue float64  `json:"coupon_value"`
            Title string  `json:"title"`
            Value float64  `json:"value"`
            Type string  `json:"type"`
            SubTitle string  `json:"sub_title"`
            IsApplied bool  `json:"is_applied"`
            Code string  `json:"code"`
         
    }
    
    // DisplayBreakupSchema used by Cart
    type DisplayBreakupSchema struct {

        
            CurrencyCode string  `json:"currency_code"`
            Key string  `json:"key"`
            Display string  `json:"display"`
            Message []string  `json:"message"`
            CurrencySymbol string  `json:"currency_symbol"`
            Value float64  `json:"value"`
         
    }
    
    // LoyaltyPoints used by Cart
    type LoyaltyPoints struct {

        
            IsApplied bool  `json:"is_applied"`
            Total float64  `json:"total"`
            Description string  `json:"description"`
            Applicable float64  `json:"applicable"`
         
    }
    
    // RawBreakupSchema used by Cart
    type RawBreakupSchema struct {

        
            ConvenienceFee float64  `json:"convenience_fee"`
            YouSaved float64  `json:"you_saved"`
            FyndCash float64  `json:"fynd_cash"`
            Total float64  `json:"total"`
            DeliveryCharge float64  `json:"delivery_charge"`
            CodCharge float64  `json:"cod_charge"`
            Subtotal float64  `json:"subtotal"`
            MrpTotal float64  `json:"mrp_total"`
            Coupon float64  `json:"coupon"`
            GstCharges float64  `json:"gst_charges"`
            GiftCard float64  `json:"gift_card"`
            Discount float64  `json:"discount"`
            Vog float64  `json:"vog"`
         
    }
    
    // CartBreakupSchema used by Cart
    type CartBreakupSchema struct {

        
            Coupon CouponBreakupSchema  `json:"coupon"`
            Display []DisplayBreakupSchema  `json:"display"`
            LoyaltyPoints LoyaltyPoints  `json:"loyalty_points"`
            Raw RawBreakupSchema  `json:"raw"`
         
    }
    
    // PromoMeta used by Cart
    type PromoMeta struct {

        
            Message string  `json:"message"`
         
    }
    
    // ProductPrice used by Cart
    type ProductPrice struct {

        
            CurrencyCode string  `json:"currency_code"`
            Effective float64  `json:"effective"`
            Marked float64  `json:"marked"`
            CurrencySymbol string  `json:"currency_symbol"`
            Selling float64  `json:"selling"`
            AddOn float64  `json:"add_on"`
         
    }
    
    // ProductPriceInfo used by Cart
    type ProductPriceInfo struct {

        
            Base ProductPrice  `json:"base"`
            Converted ProductPrice  `json:"converted"`
         
    }
    
    // CouponDetails used by Cart
    type CouponDetails struct {

        
            DiscountTotalQuantity float64  `json:"discount_total_quantity"`
            DiscountSingleQuantity float64  `json:"discount_single_quantity"`
            Code string  `json:"code"`
         
    }
    
    // BaseInfo used by Cart
    type BaseInfo struct {

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // BasePriceSchema used by Cart
    type BasePriceSchema struct {

        
            CurrencySymbol string  `json:"currency_symbol"`
            Marked float64  `json:"marked"`
            CurrencyCode string  `json:"currency_code"`
            Effective float64  `json:"effective"`
         
    }
    
    // ArticlePriceInfo used by Cart
    type ArticlePriceInfo struct {

        
            Base BasePriceSchema  `json:"base"`
            Converted BasePriceSchema  `json:"converted"`
         
    }
    
    // ProductArticle used by Cart
    type ProductArticle struct {

        
            UID string  `json:"uid"`
            ProductGroupTags []string  `json:"product_group_tags"`
            Identifier map[string]interface{}  `json:"identifier"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            Quantity float64  `json:"quantity"`
            CartItemMeta map[string]interface{}  `json:"cart_item_meta"`
            Seller BaseInfo  `json:"seller"`
            IsGiftVisible bool  `json:"is_gift_visible"`
            Size string  `json:"size"`
            Price ArticlePriceInfo  `json:"price"`
            Type string  `json:"type"`
            GiftCard map[string]interface{}  `json:"gift_card"`
            SellerIdentifier string  `json:"seller_identifier"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Store BaseInfo  `json:"store"`
         
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
            Deliverable bool  `json:"deliverable"`
            AvailableSizes []ProductAvailabilitySize  `json:"available_sizes"`
            OutOfStock bool  `json:"out_of_stock"`
            Sizes []string  `json:"sizes"`
            OtherStoreQuantity float64  `json:"other_store_quantity"`
         
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
    
    // CartProductIdentifer used by Cart
    type CartProductIdentifer struct {

        
            Identifier string  `json:"identifier"`
         
    }
    
    // BuyRulesSchema used by Cart
    type BuyRulesSchema struct {

        
            CartConditions map[string]interface{}  `json:"cart_conditions"`
            ItemCriteria map[string]interface{}  `json:"item_criteria"`
         
    }
    
    // DiscountRulesAppSchema used by Cart
    type DiscountRulesAppSchema struct {

        
            RawOffer map[string]interface{}  `json:"raw_offer"`
            Offer map[string]interface{}  `json:"offer"`
            MatchedBuyRules []string  `json:"matched_buy_rules"`
            ItemCriteria map[string]interface{}  `json:"item_criteria"`
         
    }
    
    // FreeGiftItemSchema used by Cart
    type FreeGiftItemSchema struct {

        
            ItemBrandName string  `json:"item_brand_name"`
            ItemID float64  `json:"item_id"`
            ItemSlug string  `json:"item_slug"`
            ItemPriceDetails map[string]interface{}  `json:"item_price_details"`
            ItemImagesURL []string  `json:"item_images_url"`
            ItemName string  `json:"item_name"`
         
    }
    
    // AppliedFreeArticlesSchema used by Cart
    type AppliedFreeArticlesSchema struct {

        
            Quantity float64  `json:"quantity"`
            ParentItemIdentifier string  `json:"parent_item_identifier"`
            ArticleID string  `json:"article_id"`
            FreeGiftItemDetails FreeGiftItemSchema  `json:"free_gift_item_details"`
         
    }
    
    // Ownership used by Cart
    type Ownership struct {

        
            PayableCategory string  `json:"payable_category"`
            PayableBy string  `json:"payable_by"`
         
    }
    
    // AppliedPromotion used by Cart
    type AppliedPromotion struct {

        
            ArticleQuantity float64  `json:"article_quantity"`
            PromotionType string  `json:"promotion_type"`
            BuyRules []BuyRulesSchema  `json:"buy_rules"`
            DiscountRules []DiscountRulesAppSchema  `json:"discount_rules"`
            AppliedFreeArticles []AppliedFreeArticlesSchema  `json:"applied_free_articles"`
            PromoID string  `json:"promo_id"`
            PromotionName string  `json:"promotion_name"`
            Ownership Ownership  `json:"ownership"`
            MrpPromotion bool  `json:"mrp_promotion"`
            PromotionGroup string  `json:"promotion_group"`
            Amount float64  `json:"amount"`
            OfferText string  `json:"offer_text"`
         
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

        
            UID float64  `json:"uid"`
            Categories []CategoryInfo  `json:"categories"`
            Action ProductAction  `json:"action"`
            Name string  `json:"name"`
            ItemCode string  `json:"item_code"`
            Brand BaseInfo  `json:"brand"`
            Type string  `json:"type"`
            Slug string  `json:"slug"`
            Tags []string  `json:"tags"`
            Images []ProductImage  `json:"images"`
            TeaserTag Tags  `json:"teaser_tag"`
            NetQuantity NetQuantity  `json:"net_quantity"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // CartProductInfo used by Cart
    type CartProductInfo struct {

        
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            Discount string  `json:"discount"`
            Message string  `json:"message"`
            PromoMeta PromoMeta  `json:"promo_meta"`
            Price ProductPriceInfo  `json:"price"`
            BulkOffer map[string]interface{}  `json:"bulk_offer"`
            Coupon CouponDetails  `json:"coupon"`
            CouponMessage string  `json:"coupon_message"`
            PricePerUnit ProductPriceInfo  `json:"price_per_unit"`
            Article ProductArticle  `json:"article"`
            Availability ProductAvailability  `json:"availability"`
            Moq map[string]interface{}  `json:"moq"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            Key string  `json:"key"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
            Quantity float64  `json:"quantity"`
            PromotionsApplied []AppliedPromotion  `json:"promotions_applied"`
            Product CartProduct  `json:"product"`
            IsSet bool  `json:"is_set"`
         
    }
    
    // OpenapiCartDetailsResponse used by Cart
    type OpenapiCartDetailsResponse struct {

        
            Message string  `json:"message"`
            BreakupValues CartBreakupSchema  `json:"breakup_values"`
            IsValid bool  `json:"is_valid"`
            Items []CartProductInfo  `json:"items"`
         
    }
    
    // OpenApiErrorResponse used by Cart
    type OpenApiErrorResponse struct {

        
            Success bool  `json:"success"`
            Errors map[string]interface{}  `json:"errors"`
            Message string  `json:"message"`
         
    }
    
    // ShippingAddressSchema used by Cart
    type ShippingAddressSchema struct {

        
            State string  `json:"state"`
            Address string  `json:"address"`
            Meta map[string]interface{}  `json:"meta"`
            CountryCode string  `json:"country_code"`
            Email string  `json:"email"`
            AddressType string  `json:"address_type"`
            Area string  `json:"area"`
            Pincode float64  `json:"pincode"`
            Name string  `json:"name"`
            AreaCode string  `json:"area_code"`
            City string  `json:"city"`
            CountryIsoCode string  `json:"country_iso_code"`
            Landmark string  `json:"landmark"`
            AreaCodeSlug string  `json:"area_code_slug"`
            CountryPhoneCode string  `json:"country_phone_code"`
            Phone float64  `json:"phone"`
            Country string  `json:"country"`
         
    }
    
    // OpenApiCartServiceabilityRequest used by Cart
    type OpenApiCartServiceabilityRequest struct {

        
            ShippingAddress ShippingAddressSchema  `json:"shipping_address"`
            CartItems []CartItemSchema  `json:"cart_items"`
         
    }
    
    // OpenApiCartServiceabilityResponse used by Cart
    type OpenApiCartServiceabilityResponse struct {

        
            IsValid bool  `json:"is_valid"`
            Items []CartProductInfo  `json:"items"`
            Message string  `json:"message"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            BreakupValues CartBreakupSchema  `json:"breakup_values"`
         
    }
    
    // MultiTenderPaymentMeta used by Cart
    type MultiTenderPaymentMeta struct {

        
            CurrentStatus string  `json:"current_status"`
            PaymentID string  `json:"payment_id"`
            OrderID string  `json:"order_id"`
            PaymentGateway string  `json:"payment_gateway"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
         
    }
    
    // MultiTenderPaymentMethod used by Cart
    type MultiTenderPaymentMethod struct {

        
            Meta MultiTenderPaymentMeta  `json:"meta"`
            Amount float64  `json:"amount"`
            Name string  `json:"name"`
            Mode string  `json:"mode"`
         
    }
    
    // OpenApiFilesSchema used by Cart
    type OpenApiFilesSchema struct {

        
            Values []string  `json:"values"`
            Key string  `json:"key"`
         
    }
    
    // CartItemMeta used by Cart
    type CartItemMeta struct {

        
            GroupID string  `json:"group_id"`
            PrimaryItem bool  `json:"primary_item"`
         
    }
    
    // OpenApiOrderItemSchema used by Cart
    type OpenApiOrderItemSchema struct {

        
            Files []OpenApiFilesSchema  `json:"files"`
            PriceMarked float64  `json:"price_marked"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            Quantity float64  `json:"quantity"`
            ProductID float64  `json:"product_id"`
            Meta CartItemMeta  `json:"meta"`
            PaymentMethods []MultiTenderPaymentMethod  `json:"payment_methods"`
            CashbackApplied float64  `json:"cashback_applied"`
            CodCharges float64  `json:"cod_charges"`
            DeliveryCharges float64  `json:"delivery_charges"`
            Size string  `json:"size"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            AmountPaid float64  `json:"amount_paid"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            Discount float64  `json:"discount"`
            PriceEffective float64  `json:"price_effective"`
            EmployeeDiscount float64  `json:"employee_discount"`
         
    }
    
    // OpenApiPlatformCheckoutReqSchema used by Cart
    type OpenApiPlatformCheckoutReqSchema struct {

        
            CouponCode string  `json:"coupon_code"`
            Gstin string  `json:"gstin"`
            OrderID string  `json:"order_id"`
            PaymentMethods []MultiTenderPaymentMethod  `json:"payment_methods"`
            CashbackApplied float64  `json:"cashback_applied"`
            CodCharges float64  `json:"cod_charges"`
            Comment string  `json:"comment"`
            BillingAddress ShippingAddressSchema  `json:"billing_address"`
            EmployeeDiscount map[string]interface{}  `json:"employee_discount"`
            CurrencyCode string  `json:"currency_code"`
            Files []OpenApiFilesSchema  `json:"files"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            ShippingAddress ShippingAddressSchema  `json:"shipping_address"`
            DeliveryCharges float64  `json:"delivery_charges"`
            CouponValue float64  `json:"coupon_value"`
            Coupon string  `json:"coupon"`
            CartItems []OpenApiOrderItemSchema  `json:"cart_items"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            CartValue float64  `json:"cart_value"`
            PaymentMode string  `json:"payment_mode"`
         
    }
    
    // OpenApiCheckoutResponse used by Cart
    type OpenApiCheckoutResponse struct {

        
            Success bool  `json:"success"`
            OrderRefID string  `json:"order_ref_id"`
            Message string  `json:"message"`
            OrderID string  `json:"order_id"`
         
    }
    
    // AbandonedCart used by Cart
    type AbandonedCart struct {

        
            BulkCouponDiscount float64  `json:"bulk_coupon_discount"`
            OrderID string  `json:"order_id"`
            Meta map[string]interface{}  `json:"meta"`
            AppID string  `json:"app_id"`
            LastModified string  `json:"last_modified"`
            Comment string  `json:"comment"`
            Cashback map[string]interface{}  `json:"cashback"`
            Shipments []map[string]interface{}  `json:"shipments"`
            DeliveryCharges map[string]interface{}  `json:"delivery_charges"`
            Coupon map[string]interface{}  `json:"coupon"`
            Articles []map[string]interface{}  `json:"articles"`
            ID string  `json:"_id"`
            CheckoutMode string  `json:"checkout_mode"`
            PaymentMode string  `json:"payment_mode"`
            IsActive bool  `json:"is_active"`
            UID float64  `json:"uid"`
            IsArchive bool  `json:"is_archive"`
            FyndCredits map[string]interface{}  `json:"fynd_credits"`
            Gstin string  `json:"gstin"`
            PaymentMethods []map[string]interface{}  `json:"payment_methods"`
            CodCharges map[string]interface{}  `json:"cod_charges"`
            UserID string  `json:"user_id"`
            MergeQty bool  `json:"merge_qty"`
            PickUpCustomerDetails map[string]interface{}  `json:"pick_up_customer_details"`
            CreatedOn string  `json:"created_on"`
            ExpireAt string  `json:"expire_at"`
            FcIndexMap []float64  `json:"fc_index_map"`
            IsDefault bool  `json:"is_default"`
            Payments map[string]interface{}  `json:"payments"`
            BuyNow bool  `json:"buy_now"`
            Promotion map[string]interface{}  `json:"promotion"`
            Discount float64  `json:"discount"`
            CartValue float64  `json:"cart_value"`
         
    }
    
    // AbandonedCartResponseSchema used by Cart
    type AbandonedCartResponseSchema struct {

        
            Items []AbandonedCart  `json:"items"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
            Result map[string]interface{}  `json:"result"`
            Page Page  `json:"page"`
         
    }
    
    // CartCurrencySchema used by Cart
    type CartCurrencySchema struct {

        
            Code string  `json:"code"`
            Symbol string  `json:"symbol"`
         
    }
    
    // PaymentSelectionLock used by Cart
    type PaymentSelectionLock struct {

        
            Enabled bool  `json:"enabled"`
            PaymentIdentifier string  `json:"payment_identifier"`
            DefaultOptions string  `json:"default_options"`
         
    }
    
    // CartDetailResponse used by Cart
    type CartDetailResponse struct {

        
            Items []CartProductInfo  `json:"items"`
            Message string  `json:"message"`
            LastModified string  `json:"last_modified"`
            Comment string  `json:"comment"`
            PanNo string  `json:"pan_no"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            Currency CartCurrencySchema  `json:"currency"`
            AppliedPromoDetails []AppliedPromotion  `json:"applied_promo_details"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            BreakupValues CartBreakupSchema  `json:"breakup_values"`
            ID string  `json:"id"`
            IsValid bool  `json:"is_valid"`
            Gstin string  `json:"gstin"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            PanConfig map[string]interface{}  `json:"pan_config"`
            BuyNow bool  `json:"buy_now"`
            CouponText string  `json:"coupon_text"`
            CheckoutMode string  `json:"checkout_mode"`
         
    }
    
    // AddProductCart used by Cart
    type AddProductCart struct {

        
            ProductGroupTags []string  `json:"product_group_tags"`
            ParentItemIdentifiers []map[string]string  `json:"parent_item_identifiers"`
            Quantity float64  `json:"quantity"`
            ItemID float64  `json:"item_id"`
            Display string  `json:"display"`
            Pos bool  `json:"pos"`
            SellerID float64  `json:"seller_id"`
            ItemSize string  `json:"item_size"`
            StoreID float64  `json:"store_id"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            ArticleAssignment map[string]interface{}  `json:"article_assignment"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ArticleID string  `json:"article_id"`
         
    }
    
    // AddCartRequest used by Cart
    type AddCartRequest struct {

        
            NewCart bool  `json:"new_cart"`
            Items []AddProductCart  `json:"items"`
         
    }
    
    // AddCartDetailResponse used by Cart
    type AddCartDetailResponse struct {

        
            Success bool  `json:"success"`
            Cart CartDetailResponse  `json:"cart"`
            Partial bool  `json:"partial"`
            Message string  `json:"message"`
         
    }
    
    // UpdateProductCart used by Cart
    type UpdateProductCart struct {

        
            Identifiers CartProductIdentifer  `json:"identifiers"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            Quantity float64  `json:"quantity"`
            ItemID float64  `json:"item_id"`
            ItemSize string  `json:"item_size"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ArticleID string  `json:"article_id"`
            ItemIndex float64  `json:"item_index"`
         
    }
    
    // UpdateCartRequest used by Cart
    type UpdateCartRequest struct {

        
            Operation string  `json:"operation"`
            Items []UpdateProductCart  `json:"items"`
         
    }
    
    // UpdateCartDetailResponse used by Cart
    type UpdateCartDetailResponse struct {

        
            Success bool  `json:"success"`
            Cart CartDetailResponse  `json:"cart"`
            Message string  `json:"message"`
         
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

        
            User map[string]interface{}  `json:"user"`
            Meta map[string]interface{}  `json:"meta"`
            Token string  `json:"token"`
            Source map[string]interface{}  `json:"source"`
            CreatedOn string  `json:"created_on"`
         
    }
    
    // SharedCart used by Cart
    type SharedCart struct {

        
            Items []CartProductInfo  `json:"items"`
            Message string  `json:"message"`
            LastModified string  `json:"last_modified"`
            Comment string  `json:"comment"`
            Currency CartCurrencySchema  `json:"currency"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            BreakupValues CartBreakupSchema  `json:"breakup_values"`
            ID string  `json:"id"`
            UID string  `json:"uid"`
            IsValid bool  `json:"is_valid"`
            Gstin string  `json:"gstin"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            CartID float64  `json:"cart_id"`
            SharedCartDetails SharedCartDetails  `json:"shared_cart_details"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            BuyNow bool  `json:"buy_now"`
            CouponText string  `json:"coupon_text"`
            CheckoutMode string  `json:"checkout_mode"`
         
    }
    
    // SharedCartResponse used by Cart
    type SharedCartResponse struct {

        
            Cart SharedCart  `json:"cart"`
            Error string  `json:"error"`
         
    }
    
    // CartListSchema used by Cart
    type CartListSchema struct {

        
            ItemCounts float64  `json:"item_counts"`
            UserID string  `json:"user_id"`
            PickUpCustomerDetails map[string]interface{}  `json:"pick_up_customer_details"`
            CreatedOn string  `json:"created_on"`
            CartValue float64  `json:"cart_value"`
            CartID string  `json:"cart_id"`
         
    }
    
    // MultiCartResponseSchema used by Cart
    type MultiCartResponseSchema struct {

        
            Success bool  `json:"success"`
            Data []CartListSchema  `json:"data"`
         
    }
    
    // UpdateUserCartMapping used by Cart
    type UpdateUserCartMapping struct {

        
            UserID string  `json:"user_id"`
         
    }
    
    // UserInfoSchema used by Cart
    type UserInfoSchema struct {

        
            ExternalID string  `json:"external_id"`
            UID string  `json:"uid"`
            LastName string  `json:"last_name"`
            ID string  `json:"_id"`
            CreatedAt string  `json:"created_at"`
            FirstName string  `json:"first_name"`
            Gender string  `json:"gender"`
            Mobile string  `json:"mobile"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // UserCartMappingResponse used by Cart
    type UserCartMappingResponse struct {

        
            User UserInfoSchema  `json:"user"`
            Items []CartProductInfo  `json:"items"`
            Message string  `json:"message"`
            LastModified string  `json:"last_modified"`
            Comment string  `json:"comment"`
            PanNo string  `json:"pan_no"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            Currency CartCurrencySchema  `json:"currency"`
            AppliedPromoDetails []AppliedPromotion  `json:"applied_promo_details"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            BreakupValues CartBreakupSchema  `json:"breakup_values"`
            ID string  `json:"id"`
            IsValid bool  `json:"is_valid"`
            Gstin string  `json:"gstin"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            PanConfig map[string]interface{}  `json:"pan_config"`
            BuyNow bool  `json:"buy_now"`
            CouponText string  `json:"coupon_text"`
            CheckoutMode string  `json:"checkout_mode"`
         
    }
    
    // DeleteCartRequest used by Cart
    type DeleteCartRequest struct {

        
            CartIDList []string  `json:"cart_id_list"`
            ModifiedOn string  `json:"modified_on"`
         
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
    
    // PageCouponSchema used by Cart
    type PageCouponSchema struct {

        
            TotalItemCount float64  `json:"total_item_count"`
            HasNext bool  `json:"has_next"`
            Total float64  `json:"total"`
            Current float64  `json:"current"`
            HasPrevious bool  `json:"has_previous"`
         
    }
    
    // Coupon used by Cart
    type Coupon struct {

        
            CouponCode string  `json:"coupon_code"`
            MaxDiscountValue float64  `json:"max_discount_value"`
            CouponType string  `json:"coupon_type"`
            Description string  `json:"description"`
            Message string  `json:"message"`
            MinimumCartValue float64  `json:"minimum_cart_value"`
            CouponValue float64  `json:"coupon_value"`
            Title string  `json:"title"`
            ExpiresOn string  `json:"expires_on"`
            IsApplicable bool  `json:"is_applicable"`
            SubTitle string  `json:"sub_title"`
            IsApplied bool  `json:"is_applied"`
         
    }
    
    // GetCouponResponse used by Cart
    type GetCouponResponse struct {

        
            Page PageCouponSchema  `json:"page"`
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

        
            Meta map[string]interface{}  `json:"meta"`
            CreatedByUserID string  `json:"created_by_user_id"`
            Country string  `json:"country"`
            GoogleMapPoint map[string]interface{}  `json:"google_map_point"`
            Address string  `json:"address"`
            Tags []string  `json:"tags"`
            IsDefaultAddress bool  `json:"is_default_address"`
            IsActive bool  `json:"is_active"`
            ID string  `json:"id"`
            State string  `json:"state"`
            CountryCode string  `json:"country_code"`
            AreaCode string  `json:"area_code"`
            UserID string  `json:"user_id"`
            City string  `json:"city"`
            Landmark string  `json:"landmark"`
            CartID string  `json:"cart_id"`
            Email string  `json:"email"`
            Area string  `json:"area"`
            AddressType string  `json:"address_type"`
            Name string  `json:"name"`
            GeoLocation GeoLocation  `json:"geo_location"`
            AreaCodeSlug string  `json:"area_code_slug"`
            CheckoutMode string  `json:"checkout_mode"`
            Phone string  `json:"phone"`
         
    }
    
    // PlatformGetAddressesResponse used by Cart
    type PlatformGetAddressesResponse struct {

        
            Address []PlatformAddress  `json:"address"`
         
    }
    
    // SaveAddressResponse used by Cart
    type SaveAddressResponse struct {

        
            IsDefaultAddress bool  `json:"is_default_address"`
            Success bool  `json:"success"`
            ID string  `json:"id"`
         
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
            UserID string  `json:"user_id"`
            BillingAddressID string  `json:"billing_address_id"`
            CheckoutMode string  `json:"checkout_mode"`
            CartID string  `json:"cart_id"`
         
    }
    
    // ShipmentResponse used by Cart
    type ShipmentResponse struct {

        
            Promise ShipmentPromise  `json:"promise"`
            DpID string  `json:"dp_id"`
            FulfillmentID float64  `json:"fulfillment_id"`
            Items []CartProductInfo  `json:"items"`
            BoxType string  `json:"box_type"`
            OrderType string  `json:"order_type"`
            DpOptions map[string]interface{}  `json:"dp_options"`
            FulfillmentType string  `json:"fulfillment_type"`
            ShipmentType string  `json:"shipment_type"`
            Shipments float64  `json:"shipments"`
         
    }
    
    // CartShipmentsResponse used by Cart
    type CartShipmentsResponse struct {

        
            ID string  `json:"id"`
            UID string  `json:"uid"`
            IsValid bool  `json:"is_valid"`
            Gstin string  `json:"gstin"`
            Message string  `json:"message"`
            LastModified string  `json:"last_modified"`
            Comment string  `json:"comment"`
            Error bool  `json:"error"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            Shipments []ShipmentResponse  `json:"shipments"`
            CartID float64  `json:"cart_id"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            Currency CartCurrencySchema  `json:"currency"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            BuyNow bool  `json:"buy_now"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            CouponText string  `json:"coupon_text"`
            CheckoutMode string  `json:"checkout_mode"`
            BreakupValues CartBreakupSchema  `json:"breakup_values"`
         
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

        
            PanNo string  `json:"pan_no"`
            Gstin string  `json:"gstin"`
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
    
    // StaffCheckoutSchema used by Cart
    type StaffCheckoutSchema struct {

        
            EmployeeCode string  `json:"employee_code"`
            User string  `json:"user"`
            LastName string  `json:"last_name"`
            ID string  `json:"_id"`
            FirstName string  `json:"first_name"`
         
    }
    
    // FilesSchema used by Cart
    type FilesSchema struct {

        
            Values []string  `json:"values"`
            Key string  `json:"key"`
         
    }
    
    // PlatformCartCheckoutDetailRequest used by Cart
    type PlatformCartCheckoutDetailRequest struct {

        
            EmployeeCode string  `json:"employee_code"`
            Meta map[string]interface{}  `json:"meta"`
            AddressID string  `json:"address_id"`
            BillingAddressID string  `json:"billing_address_id"`
            MerchantCode string  `json:"merchant_code"`
            PaymentParams map[string]interface{}  `json:"payment_params"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            PaymentMode string  `json:"payment_mode"`
            PaymentIdentifier string  `json:"payment_identifier"`
            ID string  `json:"id"`
            PaymentAutoConfirm bool  `json:"payment_auto_confirm"`
            CallbackURL string  `json:"callback_url"`
            PickAtStoreUID float64  `json:"pick_at_store_uid"`
            OrderType string  `json:"order_type"`
            Aggregator string  `json:"aggregator"`
            UserID string  `json:"user_id"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
            DeviceID string  `json:"device_id"`
            Staff StaffCheckoutSchema  `json:"staff"`
            OrderingStore float64  `json:"ordering_store"`
            Files []FilesSchema  `json:"files"`
            Pos bool  `json:"pos"`
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
            CheckoutMode string  `json:"checkout_mode"`
         
    }
    
    // CheckCartSchema used by Cart
    type CheckCartSchema struct {

        
            Items []CartProductInfo  `json:"items"`
            OrderID string  `json:"order_id"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
            LastModified string  `json:"last_modified"`
            Comment string  `json:"comment"`
            StoreEmps []map[string]interface{}  `json:"store_emps"`
            CodAvailable bool  `json:"cod_available"`
            Currency CartCurrencySchema  `json:"currency"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            ErrorMessage string  `json:"error_message"`
            CodMessage string  `json:"cod_message"`
            DeliveryCharges float64  `json:"delivery_charges"`
            DeliveryChargeOrderValue float64  `json:"delivery_charge_order_value"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            BreakupValues CartBreakupSchema  `json:"breakup_values"`
            ID string  `json:"id"`
            UID string  `json:"uid"`
            IsValid bool  `json:"is_valid"`
            Gstin string  `json:"gstin"`
            CodCharges float64  `json:"cod_charges"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            CartID float64  `json:"cart_id"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            StoreCode string  `json:"store_code"`
            BuyNow bool  `json:"buy_now"`
            CouponText string  `json:"coupon_text"`
            CheckoutMode string  `json:"checkout_mode"`
            UserType string  `json:"user_type"`
         
    }
    
    // CartCheckoutResponseSchema used by Cart
    type CartCheckoutResponseSchema struct {

        
            CallbackURL string  `json:"callback_url"`
            OrderID string  `json:"order_id"`
            Success bool  `json:"success"`
            PaymentConfirmURL string  `json:"payment_confirm_url"`
            Message string  `json:"message"`
            AppInterceptURL string  `json:"app_intercept_url"`
            Cart CheckCartSchema  `json:"cart"`
            Data map[string]interface{}  `json:"data"`
         
    }
    
    // CartDeliveryModesResponse used by Cart
    type CartDeliveryModesResponse struct {

        
            AvailableModes []string  `json:"available_modes"`
            PickupStores []float64  `json:"pickup_stores"`
         
    }
    
    // PickupStoreDetail used by Cart
    type PickupStoreDetail struct {

        
            ID float64  `json:"id"`
            UID float64  `json:"uid"`
            State string  `json:"state"`
            Address string  `json:"address"`
            Email string  `json:"email"`
            AddressType string  `json:"address_type"`
            Area string  `json:"area"`
            Pincode float64  `json:"pincode"`
            Name string  `json:"name"`
            AreaCode string  `json:"area_code"`
            City string  `json:"city"`
            StoreCode string  `json:"store_code"`
            Landmark string  `json:"landmark"`
            AreaCodeSlug string  `json:"area_code_slug"`
            Phone string  `json:"phone"`
            Country string  `json:"country"`
         
    }
    
    // StoreDetailsResponse used by Cart
    type StoreDetailsResponse struct {

        
            Items []PickupStoreDetail  `json:"items"`
         
    }
    
    // UpdateCartPaymentRequest used by Cart
    type UpdateCartPaymentRequest struct {

        
            ID string  `json:"id"`
            AddressID string  `json:"address_id"`
            MerchantCode string  `json:"merchant_code"`
            AggregatorName string  `json:"aggregator_name"`
            PaymentMode string  `json:"payment_mode"`
            PaymentIdentifier string  `json:"payment_identifier"`
         
    }
    
    // CouponValidity used by Cart
    type CouponValidity struct {

        
            Valid bool  `json:"valid"`
            DisplayMessageEn string  `json:"display_message_en"`
            Title string  `json:"title"`
            Discount float64  `json:"discount"`
            Code string  `json:"code"`
         
    }
    
    // PaymentCouponValidateSchema used by Cart
    type PaymentCouponValidateSchema struct {

        
            Success bool  `json:"success"`
            CouponValidity CouponValidity  `json:"coupon_validity"`
            Message string  `json:"message"`
         
    }
    
    // PaymentMetaSchema used by Cart
    type PaymentMetaSchema struct {

        
            PaymentGateway string  `json:"payment_gateway"`
            MerchantCode string  `json:"merchant_code"`
            Type string  `json:"type"`
            PaymentIdentifier string  `json:"payment_identifier"`
         
    }
    
    // PaymentMethodSchema used by Cart
    type PaymentMethodSchema struct {

        
            Payment string  `json:"payment"`
            PaymentMeta PaymentMetaSchema  `json:"payment_meta"`
            Name string  `json:"name"`
            Mode string  `json:"mode"`
            Amount float64  `json:"amount"`
         
    }
    
    // PlatformCartCheckoutDetailV2Request used by Cart
    type PlatformCartCheckoutDetailV2Request struct {

        
            EmployeeCode string  `json:"employee_code"`
            Meta map[string]interface{}  `json:"meta"`
            AddressID string  `json:"address_id"`
            BillingAddressID string  `json:"billing_address_id"`
            MerchantCode string  `json:"merchant_code"`
            PaymentParams map[string]interface{}  `json:"payment_params"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            PaymentMode string  `json:"payment_mode"`
            PaymentIdentifier string  `json:"payment_identifier"`
            ID string  `json:"id"`
            PaymentAutoConfirm bool  `json:"payment_auto_confirm"`
            CustomMeta map[string]interface{}  `json:"custom_meta"`
            CallbackURL string  `json:"callback_url"`
            PickAtStoreUID float64  `json:"pick_at_store_uid"`
            OrderType string  `json:"order_type"`
            PaymentMethods []PaymentMethodSchema  `json:"payment_methods"`
            Aggregator string  `json:"aggregator"`
            UserID string  `json:"user_id"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
            DeviceID string  `json:"device_id"`
            Staff StaffCheckoutSchema  `json:"staff"`
            OrderingStore float64  `json:"ordering_store"`
            Files []FilesSchema  `json:"files"`
            Pos bool  `json:"pos"`
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
            CheckoutMode string  `json:"checkout_mode"`
         
    }
    
    // UpdateCartPaymentRequestV2 used by Cart
    type UpdateCartPaymentRequestV2 struct {

        
            ID string  `json:"id"`
            PaymentMethods []PaymentMethodSchema  `json:"payment_methods"`
            AddressID string  `json:"address_id"`
            MerchantCode string  `json:"merchant_code"`
            AggregatorName string  `json:"aggregator_name"`
            PaymentMode string  `json:"payment_mode"`
            PaymentIdentifier string  `json:"payment_identifier"`
         
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
    
