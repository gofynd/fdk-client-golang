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

        
            DisplayFields []string  `json:"display_fields"`
            Success bool  `json:"success"`
            AppID string  `json:"app_id"`
            Aggregators []map[string]interface{}  `json:"aggregators"`
            Created bool  `json:"created"`
            ExcludedFields []string  `json:"excluded_fields"`
         
    }
    
    // ErrorCodeDescription used by Payment
    type ErrorCodeDescription struct {

        
            Description string  `json:"description"`
            Code string  `json:"code"`
            Success bool  `json:"success"`
         
    }
    
    // PaymentGatewayConfig used by Payment
    type PaymentGatewayConfig struct {

        
            IsActive bool  `json:"is_active"`
            Secret string  `json:"secret"`
            Key string  `json:"key"`
            MerchantSalt string  `json:"merchant_salt"`
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

        
            DisplayName string  `json:"display_name"`
            Code string  `json:"code"`
            PackageName string  `json:"package_name"`
            Logos PaymentModeLogo  `json:"logos"`
         
    }
    
    // PaymentModeList used by Payment
    type PaymentModeList struct {

        
            CardIsin string  `json:"card_isin"`
            Expired bool  `json:"expired"`
            IntentAppErrorDictList []IntentAppErrorList  `json:"intent_app_error_dict_list"`
            Code string  `json:"code"`
            CardBrandImage string  `json:"card_brand_image"`
            DisplayName string  `json:"display_name"`
            CodLimitPerOrder float64  `json:"cod_limit_per_order"`
            AggregatorName string  `json:"aggregator_name"`
            CardName string  `json:"card_name"`
            CardBrand string  `json:"card_brand"`
            FyndVpa string  `json:"fynd_vpa"`
            DisplayPriority float64  `json:"display_priority"`
            RemainingLimit float64  `json:"remaining_limit"`
            CompliantWithTokenisationGuidelines bool  `json:"compliant_with_tokenisation_guidelines"`
            IntentFlow bool  `json:"intent_flow"`
            CardID string  `json:"card_id"`
            Timeout float64  `json:"timeout"`
            CardIssuer string  `json:"card_issuer"`
            CodLimit float64  `json:"cod_limit"`
            CardNumber string  `json:"card_number"`
            Nickname string  `json:"nickname"`
            CardFingerprint string  `json:"card_fingerprint"`
            ExpMonth float64  `json:"exp_month"`
            CardType string  `json:"card_type"`
            CardToken string  `json:"card_token"`
            IntentAppErrorList []string  `json:"intent_app_error_list"`
            RetryCount float64  `json:"retry_count"`
            IntentApp []IntentApp  `json:"intent_app"`
            Name string  `json:"name"`
            MerchantCode string  `json:"merchant_code"`
            LogoURL PaymentModeLogo  `json:"logo_url"`
            ExpYear float64  `json:"exp_year"`
            CardReference string  `json:"card_reference"`
         
    }
    
    // RootPaymentMode used by Payment
    type RootPaymentMode struct {

        
            IsPayByCardPl bool  `json:"is_pay_by_card_pl"`
            DisplayName string  `json:"display_name"`
            AnonymousEnable bool  `json:"anonymous_enable"`
            AggregatorName string  `json:"aggregator_name"`
            SaveCard bool  `json:"save_card"`
            AddCardEnabled bool  `json:"add_card_enabled"`
            List []PaymentModeList  `json:"list"`
            DisplayPriority float64  `json:"display_priority"`
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
    
    // PayoutCustomer used by Payment
    type PayoutCustomer struct {

        
            Email string  `json:"email"`
            UniqueExternalID string  `json:"unique_external_id"`
            Mobile string  `json:"mobile"`
            ID float64  `json:"id"`
            Name string  `json:"name"`
         
    }
    
    // PayoutMoreAttributes used by Payment
    type PayoutMoreAttributes struct {

        
            BankName string  `json:"bank_name"`
            City string  `json:"city"`
            State string  `json:"state"`
            BranchName string  `json:"branch_name"`
            Country string  `json:"country"`
            AccountType string  `json:"account_type"`
            IfscCode string  `json:"ifsc_code"`
            AccountNo string  `json:"account_no"`
            AccountHolder string  `json:"account_holder"`
         
    }
    
    // PayoutAggregator used by Payment
    type PayoutAggregator struct {

        
            PayoutDetailsID float64  `json:"payout_details_id"`
            AggregatorID float64  `json:"aggregator_id"`
            AggregatorFundID float64  `json:"aggregator_fund_id"`
         
    }
    
    // Payout used by Payment
    type Payout struct {

        
            Customers PayoutCustomer  `json:"customers"`
            MoreAttributes PayoutMoreAttributes  `json:"more_attributes"`
            IsActive bool  `json:"is_active"`
            TransferType string  `json:"transfer_type"`
            PayoutsAggregators []PayoutAggregator  `json:"payouts_aggregators"`
            IsDefault bool  `json:"is_default"`
            UniqueTransferNo string  `json:"unique_transfer_no"`
         
    }
    
    // PayoutsResponse used by Payment
    type PayoutsResponse struct {

        
            Success bool  `json:"success"`
            Items []Payout  `json:"items"`
         
    }
    
    // PayoutBankDetails used by Payment
    type PayoutBankDetails struct {

        
            BankName string  `json:"bank_name"`
            City string  `json:"city"`
            Pincode float64  `json:"pincode"`
            State string  `json:"state"`
            BranchName string  `json:"branch_name"`
            Country string  `json:"country"`
            AccountType string  `json:"account_type"`
            AccountHolder string  `json:"account_holder"`
            AccountNo string  `json:"account_no"`
            IfscCode string  `json:"ifsc_code"`
         
    }
    
    // PayoutRequest used by Payment
    type PayoutRequest struct {

        
            Users map[string]interface{}  `json:"users"`
            Aggregator string  `json:"aggregator"`
            UniqueExternalID string  `json:"unique_external_id"`
            IsActive bool  `json:"is_active"`
            TransferType string  `json:"transfer_type"`
            BankDetails PayoutBankDetails  `json:"bank_details"`
         
    }
    
    // PayoutResponse used by Payment
    type PayoutResponse struct {

        
            Success bool  `json:"success"`
            Users map[string]interface{}  `json:"users"`
            Aggregator string  `json:"aggregator"`
            IsActive bool  `json:"is_active"`
            TransferType string  `json:"transfer_type"`
            UniqueTransferNo string  `json:"unique_transfer_no"`
            BankDetails map[string]interface{}  `json:"bank_details"`
            PaymentStatus string  `json:"payment_status"`
            Payouts map[string]interface{}  `json:"payouts"`
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
            Code string  `json:"code"`
            Success bool  `json:"success"`
         
    }
    
    // BankDetailsForOTP used by Payment
    type BankDetailsForOTP struct {

        
            BankName string  `json:"bank_name"`
            BranchName string  `json:"branch_name"`
            AccountHolder string  `json:"account_holder"`
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

        
            Subtitle string  `json:"subtitle"`
            DisplayName string  `json:"display_name"`
            ModifiedOn string  `json:"modified_on"`
            TransferMode string  `json:"transfer_mode"`
            ID float64  `json:"id"`
            AccountHolder string  `json:"account_holder"`
            Email string  `json:"email"`
            BankName string  `json:"bank_name"`
            BeneficiaryID string  `json:"beneficiary_id"`
            DelightsUserName string  `json:"delights_user_name"`
            AccountNo string  `json:"account_no"`
            CreatedOn string  `json:"created_on"`
            Comment string  `json:"comment"`
            Mobile string  `json:"mobile"`
            IsActive bool  `json:"is_active"`
            BranchName string  `json:"branch_name"`
            Address string  `json:"address"`
            Title string  `json:"title"`
            IfscCode string  `json:"ifsc_code"`
         
    }
    
    // OrderBeneficiaryResponse used by Payment
    type OrderBeneficiaryResponse struct {

        
            Beneficiaries []OrderBeneficiaryDetails  `json:"beneficiaries"`
            ShowBeneficiaryDetails bool  `json:"show_beneficiary_details"`
         
    }
    
    // MultiTenderPaymentMeta used by Payment
    type MultiTenderPaymentMeta struct {

        
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            CurrentStatus string  `json:"current_status"`
            PaymentGateway string  `json:"payment_gateway"`
            OrderID string  `json:"order_id"`
            PaymentID string  `json:"payment_id"`
         
    }
    
    // MultiTenderPaymentMethod used by Payment
    type MultiTenderPaymentMethod struct {

        
            Mode string  `json:"mode"`
            Amount float64  `json:"amount"`
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
            Success bool  `json:"success"`
            OrderID string  `json:"order_id"`
         
    }
    
    // CODdata used by Payment
    type CODdata struct {

        
            UserID string  `json:"user_id"`
            RemainingLimit float64  `json:"remaining_limit"`
            IsActive bool  `json:"is_active"`
            Usages float64  `json:"usages"`
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
            Mobileno string  `json:"mobileno"`
            MerchantUserID string  `json:"merchant_user_id"`
         
    }
    
    // SetCODOptionResponse used by Payment
    type SetCODOptionResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
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

        
            Logo string  `json:"logo"`
            ChannelShipmentID string  `json:"channel_shipment_id"`
            Name string  `json:"name"`
            IsAffiliate bool  `json:"is_affiliate"`
         
    }
    
    // ShipmentStatus used by Order
    type ShipmentStatus struct {

        
            ShipmentID string  `json:"shipment_id"`
            BagList []string  `json:"bag_list"`
            Status string  `json:"status"`
            CurrentShipmentStatus string  `json:"current_shipment_status"`
            ShipmentStatusID float64  `json:"shipment_status_id"`
            StatusCreatedAt string  `json:"status_created_at"`
            CreatedAt string  `json:"created_at"`
            Title string  `json:"title"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // Prices used by Order
    type Prices struct {

        
            GiftPrice float64  `json:"gift_price"`
            ValueOfGood float64  `json:"value_of_good"`
            FyndCredits float64  `json:"fynd_credits"`
            PriceMarked float64  `json:"price_marked"`
            DeliveryCharge float64  `json:"delivery_charge"`
            CashbackApplied float64  `json:"cashback_applied"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            CouponValue float64  `json:"coupon_value"`
            PromotionEffectiveDiscount float64  `json:"promotion_effective_discount"`
            AmountPaidRoundoff float64  `json:"amount_paid_roundoff"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            AmountPaid float64  `json:"amount_paid"`
            TransferPrice float64  `json:"transfer_price"`
            RefundAmount float64  `json:"refund_amount"`
            Discount float64  `json:"discount"`
            Cashback float64  `json:"cashback"`
            PriceEffective float64  `json:"price_effective"`
            RefundCredit float64  `json:"refund_credit"`
            CodCharges float64  `json:"cod_charges"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            PmPriceSplit float64  `json:"pm_price_split"`
         
    }
    
    // PlatformDeliveryAddress used by Order
    type PlatformDeliveryAddress struct {

        
            Pincode string  `json:"pincode"`
            Version string  `json:"version"`
            AddressType string  `json:"address_type"`
            Area string  `json:"area"`
            City string  `json:"city"`
            UpdatedAt string  `json:"updated_at"`
            Latitude float64  `json:"latitude"`
            ContactPerson string  `json:"contact_person"`
            Landmark string  `json:"landmark"`
            State string  `json:"state"`
            Address1 string  `json:"address1"`
            Address2 string  `json:"address2"`
            Longitude float64  `json:"longitude"`
            CreatedAt string  `json:"created_at"`
            AddressCategory string  `json:"address_category"`
            Phone string  `json:"phone"`
            Email string  `json:"email"`
            Country string  `json:"country"`
         
    }
    
    // ShipmentItemFulFillingStore used by Order
    type ShipmentItemFulFillingStore struct {

        
            Pincode string  `json:"pincode"`
            BrandStoreTags string  `json:"brand_store_tags"`
            LocationType string  `json:"location_type"`
            Code string  `json:"code"`
            City string  `json:"city"`
            ID float64  `json:"id"`
            StoreEmail string  `json:"store_email"`
            Name string  `json:"name"`
            Address string  `json:"address"`
            State string  `json:"state"`
            Phone string  `json:"phone"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // ShipmentTags used by Order
    type ShipmentTags struct {

        
            EntityType string  `json:"entity_type"`
            Slug string  `json:"slug"`
            DisplayText string  `json:"display_text"`
         
    }
    
    // Formatted used by Order
    type Formatted struct {

        
            Min string  `json:"min"`
            Max string  `json:"max"`
         
    }
    
    // LockData used by Order
    type LockData struct {

        
            Locked bool  `json:"locked"`
            Mto bool  `json:"mto"`
            LockMessage string  `json:"lock_message"`
         
    }
    
    // ShipmentTimeStamp used by Order
    type ShipmentTimeStamp struct {

        
            Min float64  `json:"min"`
            Max float64  `json:"max"`
         
    }
    
    // ShipmentItemMeta used by Order
    type ShipmentItemMeta struct {

        
            ShipmentTags []ShipmentTags  `json:"shipment_tags"`
            ShippingZone string  `json:"shipping_zone"`
            FulfilmentPriorityText string  `json:"fulfilment_priority_text"`
            External map[string]interface{}  `json:"external"`
            SameStoreAvailable bool  `json:"same_store_available"`
            StoreInvoiceUpdatedDate string  `json:"store_invoice_updated_date"`
            Formatted Formatted  `json:"formatted"`
            DebugInfo map[string]interface{}  `json:"debug_info"`
            AssignDpFromSb bool  `json:"assign_dp_from_sb"`
            ShipmentWeight float64  `json:"shipment_weight"`
            ParentDpID string  `json:"parent_dp_id"`
            ShipmentChargeableWeight float64  `json:"shipment_chargeable_weight"`
            ShipmentVolumetricWeight float64  `json:"shipment_volumetric_weight"`
            BagWeight map[string]interface{}  `json:"bag_weight"`
            LockData LockData  `json:"lock_data"`
            Sla float64  `json:"sla"`
            DpOptions map[string]interface{}  `json:"dp_options"`
            Weight float64  `json:"weight"`
            ExistingDpList []string  `json:"existing_dp_list"`
            PackagingName string  `json:"packaging_name"`
            PdfMedia []map[string]interface{}  `json:"pdf_media"`
            DpSortKey string  `json:"dp_sort_key"`
            IsInternational bool  `json:"is_international"`
            ActivityComment string  `json:"activity_comment"`
            AutoTriggerDpAssignmentAcf bool  `json:"auto_trigger_dp_assignment_acf"`
            Timestamp ShipmentTimeStamp  `json:"timestamp"`
            OrderType string  `json:"order_type"`
            EwaybillInfo map[string]interface{}  `json:"ewaybill_info"`
            Tags []map[string]interface{}  `json:"tags"`
         
    }
    
    // BagStateMapper used by Order
    type BagStateMapper struct {

        
            IsActive bool  `json:"is_active"`
            NotifyCustomer bool  `json:"notify_customer"`
            ID float64  `json:"id"`
            AppStateName string  `json:"app_state_name"`
            Name string  `json:"name"`
            AppFacing bool  `json:"app_facing"`
            StateType string  `json:"state_type"`
            AppDisplayName string  `json:"app_display_name"`
            JourneyType string  `json:"journey_type"`
            DisplayName string  `json:"display_name"`
         
    }
    
    // BagStatusHistory used by Order
    type BagStatusHistory struct {

        
            KafkaSync bool  `json:"kafka_sync"`
            ShipmentID string  `json:"shipment_id"`
            Reasons []map[string]interface{}  `json:"reasons"`
            Forward bool  `json:"forward"`
            UpdatedAt string  `json:"updated_at"`
            BshID float64  `json:"bsh_id"`
            DeliveryAwbNumber string  `json:"delivery_awb_number"`
            Status string  `json:"status"`
            StateID float64  `json:"state_id"`
            StoreID float64  `json:"store_id"`
            CreatedAt string  `json:"created_at"`
            BagStateMapper BagStateMapper  `json:"bag_state_mapper"`
            DeliveryPartnerID float64  `json:"delivery_partner_id"`
            StateType string  `json:"state_type"`
            AppDisplayName string  `json:"app_display_name"`
            BagID float64  `json:"bag_id"`
            DisplayName string  `json:"display_name"`
         
    }
    
    // BagReturnableCancelableStatus used by Order
    type BagReturnableCancelableStatus struct {

        
            IsActive bool  `json:"is_active"`
            CanBeCancelled bool  `json:"can_be_cancelled"`
            IsReturnable bool  `json:"is_returnable"`
            EnableTracking bool  `json:"enable_tracking"`
            IsCustomerReturnAllowed bool  `json:"is_customer_return_allowed"`
         
    }
    
    // Identifier used by Order
    type Identifier struct {

        
            Isbn string  `json:"isbn"`
            Alu string  `json:"alu"`
            Upc string  `json:"upc"`
            SkuCode string  `json:"sku_code"`
            Ean string  `json:"ean"`
         
    }
    
    // FinancialBreakup used by Order
    type FinancialBreakup struct {

        
            HsnCode string  `json:"hsn_code"`
            ValueOfGood float64  `json:"value_of_good"`
            AddedToFyndCash bool  `json:"added_to_fynd_cash"`
            Size string  `json:"size"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            FyndCredits float64  `json:"fynd_credits"`
            PriceMarked float64  `json:"price_marked"`
            DeliveryCharge float64  `json:"delivery_charge"`
            CashbackApplied float64  `json:"cashback_applied"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            CouponValue float64  `json:"coupon_value"`
            PromotionEffectiveDiscount float64  `json:"promotion_effective_discount"`
            TotalUnits float64  `json:"total_units"`
            ItemName string  `json:"item_name"`
            AmountPaidRoundoff float64  `json:"amount_paid_roundoff"`
            GstFee float64  `json:"gst_fee"`
            Identifiers Identifier  `json:"identifiers"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            AmountPaid float64  `json:"amount_paid"`
            TransferPrice float64  `json:"transfer_price"`
            Discount float64  `json:"discount"`
            Cashback float64  `json:"cashback"`
            PriceEffective float64  `json:"price_effective"`
            GstTag string  `json:"gst_tag"`
            RefundCredit float64  `json:"refund_credit"`
            CodCharges float64  `json:"cod_charges"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
         
    }
    
    // ShipmentListingBrand used by Order
    type ShipmentListingBrand struct {

        
            Logo string  `json:"logo"`
            CreatedOn string  `json:"created_on"`
            Name string  `json:"name"`
            LogoBase64 string  `json:"logo_base64"`
         
    }
    
    // PlatformArticleAttributes used by Order
    type PlatformArticleAttributes struct {

        
            Currency string  `json:"currency"`
         
    }
    
    // PlatformItem used by Order
    type PlatformItem struct {

        
            L3CategoryName string  `json:"l3_category_name"`
            Code string  `json:"code"`
            Image []string  `json:"image"`
            L1Category []string  `json:"l1_category"`
            Size string  `json:"size"`
            L2Category []string  `json:"l2_category"`
            ID float64  `json:"id"`
            Brand string  `json:"brand"`
            SlugKey string  `json:"slug_key"`
            Images []string  `json:"images"`
            CanReturn bool  `json:"can_return"`
            Attributes PlatformArticleAttributes  `json:"attributes"`
            Color string  `json:"color"`
            LastUpdatedAt string  `json:"last_updated_at"`
            BrandID float64  `json:"brand_id"`
            L3Category float64  `json:"l3_category"`
            DepartmentID float64  `json:"department_id"`
            Meta map[string]interface{}  `json:"meta"`
            CanCancel bool  `json:"can_cancel"`
            BranchURL string  `json:"branch_url"`
            Name string  `json:"name"`
         
    }
    
    // ReturnConfig used by Order
    type ReturnConfig struct {

        
            Returnable bool  `json:"returnable"`
            Unit string  `json:"unit"`
            Time float64  `json:"time"`
         
    }
    
    // Weight used by Order
    type Weight struct {

        
            Shipping float64  `json:"shipping"`
            Unit string  `json:"unit"`
            IsDefault bool  `json:"is_default"`
         
    }
    
    // Dimensions used by Order
    type Dimensions struct {

        
            Unit string  `json:"unit"`
            Height float64  `json:"height"`
            Width float64  `json:"width"`
            IsDefault bool  `json:"is_default"`
            Length float64  `json:"length"`
         
    }
    
    // Article used by Order
    type Article struct {

        
            IsSet bool  `json:"is_set"`
            UID string  `json:"uid"`
            SellerIdentifier string  `json:"seller_identifier"`
            ChildDetails map[string]interface{}  `json:"child_details"`
            RawMeta string  `json:"raw_meta"`
            Code string  `json:"code"`
            ReturnConfig ReturnConfig  `json:"return_config"`
            ID string  `json:"_id"`
            Currency map[string]interface{}  `json:"currency"`
            Weight Weight  `json:"weight"`
            ASet map[string]interface{}  `json:"a_set"`
            Dimensions Dimensions  `json:"dimensions"`
            Identifiers map[string]interface{}  `json:"identifiers"`
            EspModified bool  `json:"esp_modified"`
            Size string  `json:"size"`
         
    }
    
    // Dates used by Order
    type Dates struct {

        
            OrderCreated string  `json:"order_created"`
            DeliveryDate string  `json:"delivery_date"`
         
    }
    
    // GSTDetailsData used by Order
    type GSTDetailsData struct {

        
            IsDefaultHsnCode bool  `json:"is_default_hsn_code"`
            GstTag string  `json:"gst_tag"`
            SgstTaxPercentage float64  `json:"sgst_tax_percentage"`
            SgstGstFee string  `json:"sgst_gst_fee"`
            IgstTaxPercentage float64  `json:"igst_tax_percentage"`
            CgstGstFee string  `json:"cgst_gst_fee"`
            GstinCode string  `json:"gstin_code"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            CgstTaxPercentage float64  `json:"cgst_tax_percentage"`
            IgstGstFee string  `json:"igst_gst_fee"`
            HsnCode string  `json:"hsn_code"`
            GstFee float64  `json:"gst_fee"`
            ValueOfGood float64  `json:"value_of_good"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            HsnCodeID string  `json:"hsn_code_id"`
         
    }
    
    // ReplacementDetails used by Order
    type ReplacementDetails struct {

        
            OriginalAffiliateOrderID string  `json:"original_affiliate_order_id"`
            ReplacementType string  `json:"replacement_type"`
         
    }
    
    // AffiliateMeta used by Order
    type AffiliateMeta struct {

        
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            CouponCode string  `json:"coupon_code"`
            EmployeeDiscount float64  `json:"employee_discount"`
            IsPriority bool  `json:"is_priority"`
            ChannelOrderID string  `json:"channel_order_id"`
            Quantity float64  `json:"quantity"`
            MarketplaceInvoiceID string  `json:"marketplace_invoice_id"`
            SizeLevelTotalQty float64  `json:"size_level_total_qty"`
            ReplacementDetails ReplacementDetails  `json:"replacement_details"`
            ChannelShipmentID string  `json:"channel_shipment_id"`
            OrderItemID string  `json:"order_item_id"`
            BoxType string  `json:"box_type"`
            DueDate string  `json:"due_date"`
         
    }
    
    // AffiliateBagDetails used by Order
    type AffiliateBagDetails struct {

        
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            AffiliateMeta AffiliateMeta  `json:"affiliate_meta"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            EmployeeDiscount float64  `json:"employee_discount"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
         
    }
    
    // BagUnit used by Order
    type BagUnit struct {

        
            CurrentStatus BagStatusHistory  `json:"current_status"`
            BagExpiryDate string  `json:"bag_expiry_date"`
            LineNumber float64  `json:"line_number"`
            Status BagReturnableCancelableStatus  `json:"status"`
            EntityType string  `json:"entity_type"`
            CurrentOperationalStatus BagStatusHistory  `json:"current_operational_status"`
            Prices Prices  `json:"prices"`
            BagType string  `json:"bag_type"`
            DisplayName string  `json:"display_name"`
            Size string  `json:"size"`
            FinancialBreakup []FinancialBreakup  `json:"financial_breakup"`
            ProductQuantity float64  `json:"product_quantity"`
            Brand ShipmentListingBrand  `json:"brand"`
            Item PlatformItem  `json:"item"`
            Article Article  `json:"article"`
            Reasons []map[string]interface{}  `json:"reasons"`
            CanReturn bool  `json:"can_return"`
            Dates Dates  `json:"dates"`
            BagStatus []BagStatusHistory  `json:"bag_status"`
            Meta map[string]interface{}  `json:"meta"`
            Gst GSTDetailsData  `json:"gst"`
            CanCancel bool  `json:"can_cancel"`
            AffiliateBagDetails AffiliateBagDetails  `json:"affiliate_bag_details"`
            BagID float64  `json:"bag_id"`
         
    }
    
    // UserDataInfo used by Order
    type UserDataInfo struct {

        
            UID float64  `json:"uid"`
            AvisUserID string  `json:"avis_user_id"`
            Mobile string  `json:"mobile"`
            Gender string  `json:"gender"`
            FirstName string  `json:"first_name"`
            Name string  `json:"name"`
            IsAnonymousUser bool  `json:"is_anonymous_user"`
            Email string  `json:"email"`
            LastName string  `json:"last_name"`
         
    }
    
    // ShipmentItem used by Order
    type ShipmentItem struct {

        
            Channel ShipmentListingChannel  `json:"channel"`
            CanProcess bool  `json:"can_process"`
            ShipmentStatus ShipmentStatus  `json:"shipment_status"`
            PaymentMethods map[string]interface{}  `json:"payment_methods"`
            Prices Prices  `json:"prices"`
            LockStatus bool  `json:"lock_status"`
            DisplayName string  `json:"display_name"`
            DeliveryAddress PlatformDeliveryAddress  `json:"delivery_address"`
            ShipmentCreatedAt string  `json:"shipment_created_at"`
            OrderingChannnel string  `json:"ordering_channnel"`
            StatusCreatedAt string  `json:"status_created_at"`
            EstimatedSlaTime string  `json:"estimated_sla_time"`
            FulfillingStore ShipmentItemFulFillingStore  `json:"fulfilling_store"`
            ShipmentID string  `json:"shipment_id"`
            CustomerNote string  `json:"customer_note"`
            OrderDate string  `json:"order_date"`
            TotalBags float64  `json:"total_bags"`
            Meta ShipmentItemMeta  `json:"meta"`
            PreviousShipmentID string  `json:"previous_shipment_id"`
            InvoiceID string  `json:"invoice_id"`
            Bags []BagUnit  `json:"bags"`
            PaymentMode string  `json:"payment_mode"`
            OrderID string  `json:"order_id"`
            User UserDataInfo  `json:"user"`
         
    }
    
    // ShipmentInternalPlatformViewResponse used by Order
    type ShipmentInternalPlatformViewResponse struct {

        
            Page Page  `json:"page"`
            Success bool  `json:"success"`
            Items []ShipmentItem  `json:"items"`
            TotalCount float64  `json:"total_count"`
            Message string  `json:"message"`
            Lane string  `json:"lane"`
         
    }
    
    // Error used by Order
    type Error struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // OrderDetailsData used by Order
    type OrderDetailsData struct {

        
            AffiliateID string  `json:"affiliate_id"`
            OrderingChannel string  `json:"ordering_channel"`
            OrderDate string  `json:"order_date"`
            OrderingChannelLogo map[string]interface{}  `json:"ordering_channel_logo"`
            OrderValue string  `json:"order_value"`
            TaxDetails map[string]interface{}  `json:"tax_details"`
            CodCharges string  `json:"cod_charges"`
            Source string  `json:"source"`
            FyndOrderID string  `json:"fynd_order_id"`
         
    }
    
    // TrackingList used by Order
    type TrackingList struct {

        
            IsPassed bool  `json:"is_passed"`
            IsCurrent bool  `json:"is_current"`
            Text string  `json:"text"`
            Status string  `json:"status"`
            Time string  `json:"time"`
         
    }
    
    // Meta used by Order
    type Meta struct {

        
            MerchantCode string  `json:"merchant_code"`
            LogoURL map[string]string  `json:"logo_url"`
            PaymentGateway string  `json:"payment_gateway"`
            PaymentID string  `json:"payment_id"`
            PaymentIdentifier string  `json:"payment_identifier"`
         
    }
    
    // PaymentMethod used by Order
    type PaymentMethod struct {

        
            Amount float64  `json:"amount"`
            RefundBy string  `json:"refund_by"`
            Mode string  `json:"mode"`
            CollectBy string  `json:"collect_by"`
            Name string  `json:"name"`
            Meta Meta  `json:"meta"`
         
    }
    
    // ShipmentStatusData used by Order
    type ShipmentStatusData struct {

        
            ShipmentID string  `json:"shipment_id"`
            BagList []string  `json:"bag_list"`
            ID float64  `json:"id"`
            Status string  `json:"status"`
            CreatedAt string  `json:"created_at"`
            Meta map[string]interface{}  `json:"meta"`
            DisplayName string  `json:"display_name"`
         
    }
    
    // UserDetailsData used by Order
    type UserDetailsData struct {

        
            Pincode string  `json:"pincode"`
            AddressType string  `json:"address_type"`
            Area string  `json:"area"`
            City string  `json:"city"`
            Landmark string  `json:"landmark"`
            Name string  `json:"name"`
            Address string  `json:"address"`
            State string  `json:"state"`
            Address1 string  `json:"address1"`
            Phone string  `json:"phone"`
            Email string  `json:"email"`
            Country string  `json:"country"`
         
    }
    
    // InvoiceInfo used by Order
    type InvoiceInfo struct {

        
            UpdatedDate string  `json:"updated_date"`
            InvoiceURL string  `json:"invoice_url"`
            ExternalInvoiceID string  `json:"external_invoice_id"`
            LabelURL string  `json:"label_url"`
            StoreInvoiceID string  `json:"store_invoice_id"`
            CreditNoteID string  `json:"credit_note_id"`
         
    }
    
    // PhoneDetails used by Order
    type PhoneDetails struct {

        
            Number string  `json:"number"`
            CountryCode float64  `json:"country_code"`
         
    }
    
    // ContactDetails used by Order
    type ContactDetails struct {

        
            Phone []PhoneDetails  `json:"phone"`
            Emails []string  `json:"emails"`
         
    }
    
    // CompanyDetails used by Order
    type CompanyDetails struct {

        
            CompanyContact ContactDetails  `json:"company_contact"`
            CompanyCin string  `json:"company_cin"`
            CompanyGst string  `json:"company_gst"`
            Address map[string]interface{}  `json:"address"`
            CompanyID float64  `json:"company_id"`
            CompanyName string  `json:"company_name"`
         
    }
    
    // ShipmentPayments used by Order
    type ShipmentPayments struct {

        
            Logo string  `json:"logo"`
            Mode string  `json:"mode"`
            Source string  `json:"source"`
         
    }
    
    // FulfillingStore used by Order
    type FulfillingStore struct {

        
            Pincode string  `json:"pincode"`
            FulfillmentChannel string  `json:"fulfillment_channel"`
            Code string  `json:"code"`
            ID float64  `json:"id"`
            City string  `json:"city"`
            State string  `json:"state"`
            Address string  `json:"address"`
            ContactPerson string  `json:"contact_person"`
            Phone string  `json:"phone"`
            Meta map[string]interface{}  `json:"meta"`
            StoreName string  `json:"store_name"`
            Country string  `json:"country"`
         
    }
    
    // BuyerDetails used by Order
    type BuyerDetails struct {

        
            Pincode float64  `json:"pincode"`
            Gstin string  `json:"gstin"`
            City string  `json:"city"`
            State string  `json:"state"`
            Name string  `json:"name"`
            Address string  `json:"address"`
            AjioSiteID string  `json:"ajio_site_id"`
         
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
    
    // ShipmentMeta used by Order
    type ShipmentMeta struct {

        
            ShipmentTags []ShipmentTags  `json:"shipment_tags"`
            FulfilmentPriorityText string  `json:"fulfilment_priority_text"`
            ReturnDetails map[string]interface{}  `json:"return_details"`
            External map[string]interface{}  `json:"external"`
            B2bBuyerDetails BuyerDetails  `json:"b2b_buyer_details"`
            DpName string  `json:"dp_name"`
            SameStoreAvailable bool  `json:"same_store_available"`
            StoreInvoiceUpdatedDate string  `json:"store_invoice_updated_date"`
            DueDate string  `json:"due_date"`
            Formatted Formatted  `json:"formatted"`
            DebugInfo DebugInfo  `json:"debug_info"`
            ReturnStoreNode float64  `json:"return_store_node"`
            ForwardAffiliateOrderID string  `json:"forward_affiliate_order_id"`
            AssignDpFromSb bool  `json:"assign_dp_from_sb"`
            ShipmentWeight float64  `json:"shipment_weight"`
            ReturnAwbNumber string  `json:"return_awb_number"`
            ReturnAffiliateOrderID string  `json:"return_affiliate_order_id"`
            Dimension Dimensions  `json:"dimension"`
            PoNumber string  `json:"po_number"`
            ParentDpID string  `json:"parent_dp_id"`
            ShipmentVolumetricWeight float64  `json:"shipment_volumetric_weight"`
            LockData LockData  `json:"lock_data"`
            BagWeight map[string]interface{}  `json:"bag_weight"`
            B2cBuyerDetails map[string]interface{}  `json:"b2c_buyer_details"`
            ReturnAffiliateShipmentID string  `json:"return_affiliate_shipment_id"`
            MarketplaceStoreID string  `json:"marketplace_store_id"`
            DpOptions map[string]interface{}  `json:"dp_options"`
            Weight float64  `json:"weight"`
            PackagingName string  `json:"packaging_name"`
            AwbNumber string  `json:"awb_number"`
            EinvoiceInfo EinvoiceInfo  `json:"einvoice_info"`
            DpSortKey string  `json:"dp_sort_key"`
            AutoTriggerDpAssignmentAcf bool  `json:"auto_trigger_dp_assignment_acf"`
            Timestamp ShipmentTimeStamp  `json:"timestamp"`
            BoxType string  `json:"box_type"`
            OrderType string  `json:"order_type"`
            ForwardAffiliateShipmentID string  `json:"forward_affiliate_shipment_id"`
            DpID string  `json:"dp_id"`
            EwaybillInfo map[string]interface{}  `json:"ewaybill_info"`
         
    }
    
    // PDFLinks used by Order
    type PDFLinks struct {

        
            LabelType string  `json:"label_type"`
            InvoicePos string  `json:"invoice_pos"`
            LabelExport string  `json:"label_export"`
            Label string  `json:"label"`
            DeliveryChallanA4 string  `json:"delivery_challan_a4"`
            InvoiceType string  `json:"invoice_type"`
            InvoiceA4 string  `json:"invoice_a4"`
            InvoiceA6 string  `json:"invoice_a6"`
            LabelA6 string  `json:"label_a6"`
            CreditNoteURL string  `json:"credit_note_url"`
            InvoiceExport string  `json:"invoice_export"`
            LabelPos string  `json:"label_pos"`
            PoInvoice string  `json:"po_invoice"`
            LabelA4 string  `json:"label_a4"`
            B2b string  `json:"b2b"`
            Invoice string  `json:"invoice"`
         
    }
    
    // AffiliateDetails used by Order
    type AffiliateDetails struct {

        
            AffiliateID string  `json:"affiliate_id"`
            ShipmentMeta ShipmentMeta  `json:"shipment_meta"`
            AffiliateMeta AffiliateMeta  `json:"affiliate_meta"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            CompanyAffiliateTag string  `json:"company_affiliate_tag"`
            AffiliateStoreID string  `json:"affiliate_store_id"`
            AdID string  `json:"ad_id"`
            PdfLinks PDFLinks  `json:"pdf_links"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
         
    }
    
    // ShipmentDetails used by Order
    type ShipmentDetails struct {

        
            LockStatus bool  `json:"lock_status"`
            LockMessage string  `json:"lock_message"`
            ActionToStatus map[string]interface{}  `json:"action_to_status"`
         
    }
    
    // OrderingStoreDetails used by Order
    type OrderingStoreDetails struct {

        
            Pincode string  `json:"pincode"`
            Code string  `json:"code"`
            ID float64  `json:"id"`
            City string  `json:"city"`
            State string  `json:"state"`
            Address string  `json:"address"`
            ContactPerson string  `json:"contact_person"`
            Phone string  `json:"phone"`
            Meta map[string]interface{}  `json:"meta"`
            StoreName string  `json:"store_name"`
            Country string  `json:"country"`
         
    }
    
    // DPDetailsData used by Order
    type DPDetailsData struct {

        
            Pincode string  `json:"pincode"`
            GstTag string  `json:"gst_tag"`
            TrackURL string  `json:"track_url"`
            AwbNo string  `json:"awb_no"`
            ID float64  `json:"id"`
            Name string  `json:"name"`
            EwayBillID string  `json:"eway_bill_id"`
            Country string  `json:"country"`
         
    }
    
    // CurrentStatus used by Order
    type CurrentStatus struct {

        
            KafkaSync bool  `json:"kafka_sync"`
            ShipmentID string  `json:"shipment_id"`
            DeliveryAwbNumber string  `json:"delivery_awb_number"`
            ID float64  `json:"id"`
            UpdatedAt float64  `json:"updated_at"`
            CreatedAt string  `json:"created_at"`
            Status string  `json:"status"`
            StoreID float64  `json:"store_id"`
            StateID float64  `json:"state_id"`
            BagStateMapper BagStateMapper  `json:"bag_state_mapper"`
            DeliveryPartnerID float64  `json:"delivery_partner_id"`
            StateType string  `json:"state_type"`
            BagID float64  `json:"bag_id"`
         
    }
    
    // BagGST used by Order
    type BagGST struct {

        
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            GstTag string  `json:"gst_tag"`
            SgstTaxPercentage float64  `json:"sgst_tax_percentage"`
            SgstGstFee string  `json:"sgst_gst_fee"`
            IgstTaxPercentage float64  `json:"igst_tax_percentage"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            GstinCode string  `json:"gstin_code"`
            CgstGstFee string  `json:"cgst_gst_fee"`
            CgstTaxPercentage float64  `json:"cgst_tax_percentage"`
            IgstGstFee string  `json:"igst_gst_fee"`
            HsnCode string  `json:"hsn_code"`
            GstFee float64  `json:"gst_fee"`
            ValueOfGood float64  `json:"value_of_good"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            IsDefaultHsnCode bool  `json:"is_default_hsn_code"`
            HsnCodeID string  `json:"hsn_code_id"`
         
    }
    
    // BagConfigs used by Order
    type BagConfigs struct {

        
            IsActive bool  `json:"is_active"`
            EnableTracking bool  `json:"enable_tracking"`
            CanBeCancelled bool  `json:"can_be_cancelled"`
            AllowForceReturn bool  `json:"allow_force_return"`
            IsReturnable bool  `json:"is_returnable"`
            IsCustomerReturnAllowed bool  `json:"is_customer_return_allowed"`
         
    }
    
    // OrderBrandName used by Order
    type OrderBrandName struct {

        
            Company float64  `json:"company"`
            BrandName string  `json:"brand_name"`
            ID float64  `json:"id"`
            Logo string  `json:"logo"`
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
         
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

        
            Value float64  `json:"value"`
            Type string  `json:"type"`
         
    }
    
    // AppliedPromos used by Order
    type AppliedPromos struct {

        
            Amount float64  `json:"amount"`
            BuyRules []BuyRules  `json:"buy_rules"`
            MrpPromotion bool  `json:"mrp_promotion"`
            PromotionName string  `json:"promotion_name"`
            DiscountRules []DiscountRules  `json:"discount_rules"`
            PromoID string  `json:"promo_id"`
            PromotionType string  `json:"promotion_type"`
            ArticleQuantity float64  `json:"article_quantity"`
         
    }
    
    // OrderBagArticle used by Order
    type OrderBagArticle struct {

        
            Size string  `json:"size"`
            UID string  `json:"uid"`
            Identifiers map[string]interface{}  `json:"identifiers"`
            ReturnConfig ReturnConfig  `json:"return_config"`
         
    }
    
    // B2BPODetails used by Order
    type B2BPODetails struct {

        
            PoTaxAmount float64  `json:"po_tax_amount"`
            PoLineAmount float64  `json:"po_line_amount"`
            TotalGstPercentage float64  `json:"total_gst_percentage"`
            PartialCanRet bool  `json:"partial_can_ret"`
            DockerNumber string  `json:"docker_number"`
            ItemBasePrice float64  `json:"item_base_price"`
         
    }
    
    // GiftCard used by Order
    type GiftCard struct {

        
            IsGiftApplied bool  `json:"is_gift_applied"`
            GiftPrice float64  `json:"gift_price"`
            DisplayText string  `json:"display_text"`
            GiftMessage string  `json:"gift_message"`
         
    }
    
    // BagMeta used by Order
    type BagMeta struct {

        
            CustomMessage string  `json:"custom_message"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            CustomJson map[string]interface{}  `json:"custom_json"`
            GroupID string  `json:"group_id"`
            B2bPoDetails B2BPODetails  `json:"b2b_po_details"`
            PartialCanRet bool  `json:"partial_can_ret"`
            DocketNumber string  `json:"docket_number"`
            GiftCard GiftCard  `json:"gift_card"`
         
    }
    
    // AffiliateBagsDetails used by Order
    type AffiliateBagsDetails struct {

        
            CouponCode string  `json:"coupon_code"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
         
    }
    
    // OrderBags used by Order
    type OrderBags struct {

        
            CurrentStatus CurrentStatus  `json:"current_status"`
            LineNumber float64  `json:"line_number"`
            Quantity float64  `json:"quantity"`
            Identifier string  `json:"identifier"`
            EntityType string  `json:"entity_type"`
            Prices Prices  `json:"prices"`
            ParentPromoBags map[string]interface{}  `json:"parent_promo_bags"`
            DisplayName string  `json:"display_name"`
            GstDetails BagGST  `json:"gst_details"`
            FinancialBreakup FinancialBreakup  `json:"financial_breakup"`
            DeliveryAddress PlatformDeliveryAddress  `json:"delivery_address"`
            BagConfigs BagConfigs  `json:"bag_configs"`
            Brand OrderBrandName  `json:"brand"`
            Item PlatformItem  `json:"item"`
            AppliedPromos []AppliedPromos  `json:"applied_promos"`
            Article OrderBagArticle  `json:"article"`
            CanReturn bool  `json:"can_return"`
            SellerIdentifier string  `json:"seller_identifier"`
            Meta BagMeta  `json:"meta"`
            GroupID string  `json:"group_id"`
            CanCancel bool  `json:"can_cancel"`
            AffiliateBagDetails AffiliateBagsDetails  `json:"affiliate_bag_details"`
            BagID float64  `json:"bag_id"`
         
    }
    
    // PlatformShipment used by Order
    type PlatformShipment struct {

        
            ShipmentStatus string  `json:"shipment_status"`
            Order OrderDetailsData  `json:"order"`
            TrackingList []TrackingList  `json:"tracking_list"`
            PaymentMethods map[string]PaymentMethod  `json:"payment_methods"`
            Status ShipmentStatusData  `json:"status"`
            EnableDpTracking bool  `json:"enable_dp_tracking"`
            TotalItems float64  `json:"total_items"`
            Prices Prices  `json:"prices"`
            LockStatus bool  `json:"lock_status"`
            BillingDetails UserDetailsData  `json:"billing_details"`
            GstDetails GSTDetailsData  `json:"gst_details"`
            CustomMessage string  `json:"custom_message"`
            Invoice InvoiceInfo  `json:"invoice"`
            ShipmentCreatedAt string  `json:"shipment_created_at"`
            ShipmentImages []string  `json:"shipment_images"`
            CompanyDetails CompanyDetails  `json:"company_details"`
            Payments ShipmentPayments  `json:"payments"`
            Coupon map[string]interface{}  `json:"coupon"`
            EstimatedSlaTime string  `json:"estimated_sla_time"`
            JourneyType string  `json:"journey_type"`
            FulfillingStore FulfillingStore  `json:"fulfilling_store"`
            CustomMeta []map[string]interface{}  `json:"custom_meta"`
            ShipmentID string  `json:"shipment_id"`
            DpAssignment bool  `json:"dp_assignment"`
            ForwardShipmentID string  `json:"forward_shipment_id"`
            PackagingType string  `json:"packaging_type"`
            DeliverySlot map[string]interface{}  `json:"delivery_slot"`
            AffiliateDetails AffiliateDetails  `json:"affiliate_details"`
            PlatformLogo string  `json:"platform_logo"`
            TotalBags float64  `json:"total_bags"`
            FulfilmentPriority float64  `json:"fulfilment_priority"`
            OperationalStatus string  `json:"operational_status"`
            Meta ShipmentMeta  `json:"meta"`
            PriorityText string  `json:"priority_text"`
            ShipmentDetails ShipmentDetails  `json:"shipment_details"`
            Vertical string  `json:"vertical"`
            ShipmentQuantity float64  `json:"shipment_quantity"`
            OrderingStore OrderingStoreDetails  `json:"ordering_store"`
            InvoiceID string  `json:"invoice_id"`
            DpDetails DPDetailsData  `json:"dp_details"`
            Bags []OrderBags  `json:"bags"`
            PaymentMode string  `json:"payment_mode"`
            BagStatusHistory []BagStatusHistory  `json:"bag_status_history"`
            CanUpdateDimension bool  `json:"can_update_dimension"`
            User UserDataInfo  `json:"user"`
            UserAgent string  `json:"user_agent"`
            PickedDate string  `json:"picked_date"`
            IsDpAssignEnabled bool  `json:"is_dp_assign_enabled"`
            DeliveryDetails UserDetailsData  `json:"delivery_details"`
         
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
            Success bool  `json:"success"`
            PresignedType string  `json:"presigned_type"`
            PresignedUrls map[string]string  `json:"presigned_urls"`
            ExpiresIn string  `json:"expires_in"`
         
    }
    
    // ResponseGetAssetShipment used by Order
    type ResponseGetAssetShipment struct {

        
            Success bool  `json:"success"`
            Data []AssetByShipment  `json:"data"`
            Message string  `json:"message"`
         
    }
    
    // TaxDetails used by Order
    type TaxDetails struct {

        
            PanNo string  `json:"pan_no"`
            Gstin string  `json:"gstin"`
         
    }
    
    // TransactionData used by Order
    type TransactionData struct {

        
            AmountPaid float64  `json:"amount_paid"`
            PaymentID string  `json:"payment_id"`
            TransactionID string  `json:"transaction_id"`
            TerminalID string  `json:"terminal_id"`
            Entity string  `json:"entity"`
            UniqueReferenceNumber string  `json:"unique_reference_number"`
            Status string  `json:"status"`
            Currency string  `json:"currency"`
         
    }
    
    // PlatformUserDetails used by Order
    type PlatformUserDetails struct {

        
            PlatformUserLastName string  `json:"platform_user_last_name"`
            PlatformUserFirstName string  `json:"platform_user_first_name"`
            PlatformUserEmployeeCode string  `json:"platform_user_employee_code"`
            PlatformUserID string  `json:"platform_user_id"`
         
    }
    
    // BillingStaffDetails used by Order
    type BillingStaffDetails struct {

        
            EmployeeCode string  `json:"employee_code"`
            FirstName string  `json:"first_name"`
            User string  `json:"user"`
            StaffID float64  `json:"staff_id"`
            LastName string  `json:"last_name"`
         
    }
    
    // OrderMeta used by Order
    type OrderMeta struct {

        
            TransactionData TransactionData  `json:"transaction_data"`
            CurrencySymbol string  `json:"currency_symbol"`
            Staff map[string]interface{}  `json:"staff"`
            EmployeeID float64  `json:"employee_id"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            Comment string  `json:"comment"`
            OrderTags []map[string]interface{}  `json:"order_tags"`
            MongoCartID float64  `json:"mongo_cart_id"`
            PaymentType string  `json:"payment_type"`
            PlatformUserDetails PlatformUserDetails  `json:"platform_user_details"`
            CartID float64  `json:"cart_id"`
            BillingStaffDetails BillingStaffDetails  `json:"billing_staff_details"`
            CustomerNote string  `json:"customer_note"`
            CompanyLogo string  `json:"company_logo"`
            OrderingStore float64  `json:"ordering_store"`
            OrderPlatform string  `json:"order_platform"`
            OrderType string  `json:"order_type"`
            Files []map[string]interface{}  `json:"files"`
            OrderChildEntities []string  `json:"order_child_entities"`
         
    }
    
    // OrderData used by Order
    type OrderData struct {

        
            PaymentMethods map[string]PaymentMethod  `json:"payment_methods"`
            OrderDate string  `json:"order_date"`
            TaxDetails TaxDetails  `json:"tax_details"`
            Prices Prices  `json:"prices"`
            Meta OrderMeta  `json:"meta"`
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
            Text string  `json:"text"`
            Actions []map[string]interface{}  `json:"actions"`
            TotalItems float64  `json:"total_items"`
            Value string  `json:"value"`
         
    }
    
    // SuperLane used by Order
    type SuperLane struct {

        
            Text string  `json:"text"`
            Options []SubLane  `json:"options"`
            TotalItems float64  `json:"total_items"`
            Value string  `json:"value"`
         
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
            Name string  `json:"name"`
            Value string  `json:"value"`
         
    }
    
    // PlatformOrderItems used by Order
    type PlatformOrderItems struct {

        
            Channel PlatformChannel  `json:"channel"`
            Shipments []PlatformShipment  `json:"shipments"`
            OrderCreatedTime string  `json:"order_created_time"`
            PaymentMode string  `json:"payment_mode"`
            BreakupValues []PlatformBreakupValues  `json:"breakup_values"`
            OrderID string  `json:"order_id"`
            TotalOrderValue float64  `json:"total_order_value"`
            OrderValue float64  `json:"order_value"`
            Meta map[string]interface{}  `json:"meta"`
            UserInfo UserDataInfo  `json:"user_info"`
         
    }
    
    // OrderListingResponse used by Order
    type OrderListingResponse struct {

        
            Page Page  `json:"page"`
            Success bool  `json:"success"`
            Items []PlatformOrderItems  `json:"items"`
            TotalCount float64  `json:"total_count"`
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
            Options []Options  `json:"options"`
            Key string  `json:"key"`
            Value float64  `json:"value"`
         
    }
    
    // MetricCountResponse used by Order
    type MetricCountResponse struct {

        
            Items []MetricsCount  `json:"items"`
         
    }
    
    // PlatformTrack used by Order
    type PlatformTrack struct {

        
            RawStatus string  `json:"raw_status"`
            Reason string  `json:"reason"`
            UpdatedAt string  `json:"updated_at"`
            Awb string  `json:"awb"`
            Status string  `json:"status"`
            UpdatedTime string  `json:"updated_time"`
            ShipmentType string  `json:"shipment_type"`
            Meta map[string]interface{}  `json:"meta"`
            LastLocationRecievedAt string  `json:"last_location_recieved_at"`
            AccountName string  `json:"account_name"`
         
    }
    
    // PlatformShipmentTrack used by Order
    type PlatformShipmentTrack struct {

        
            Meta map[string]interface{}  `json:"meta"`
            Results []PlatformTrack  `json:"results"`
         
    }
    
    // FilterInfoOption used by Order
    type FilterInfoOption struct {

        
            MinSearchSize float64  `json:"min_search_size"`
            ShowUI bool  `json:"show_ui"`
            Text string  `json:"text"`
            PlaceholderText string  `json:"placeholder_text"`
            Name string  `json:"name"`
            Value string  `json:"value"`
         
    }
    
    // FiltersInfo used by Order
    type FiltersInfo struct {

        
            Options []FilterInfoOption  `json:"options"`
            Required bool  `json:"required"`
            Value string  `json:"value"`
            Text string  `json:"text"`
            PlaceholderText string  `json:"placeholder_text"`
            Type string  `json:"type"`
         
    }
    
    // AdvanceFilterInfo used by Order
    type AdvanceFilterInfo struct {

        
            Returned []FiltersInfo  `json:"returned"`
            ActionCentre []FiltersInfo  `json:"action_centre"`
            Page map[string]interface{}  `json:"page"`
            Processed []FiltersInfo  `json:"processed"`
            AppliedFilters map[string]interface{}  `json:"applied_filters"`
            Filters []FiltersInfo  `json:"filters"`
            Unfulfilled []FiltersInfo  `json:"unfulfilled"`
         
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
            ReportCreatedAt string  `json:"report_created_at"`
            ReportName string  `json:"report_name"`
            ReportType string  `json:"report_type"`
            ReportID string  `json:"report_id"`
            Status string  `json:"status"`
            RequestDetails map[string]interface{}  `json:"request_details"`
            DisplayName string  `json:"display_name"`
         
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

        
            Value string  `json:"value"`
            Message string  `json:"message"`
            Type string  `json:"type"`
         
    }
    
    // JioCodeUpsertResponse used by Order
    type JioCodeUpsertResponse struct {

        
            Data []map[string]interface{}  `json:"data"`
            TraceID string  `json:"trace_id"`
            Success bool  `json:"success"`
            Identifier string  `json:"identifier"`
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
            Data map[string]interface{}  `json:"data"`
            StoreCode string  `json:"store_code"`
            Label map[string]interface{}  `json:"label"`
            StoreID string  `json:"store_id"`
            CompanyID string  `json:"company_id"`
            DoInvoiceLabelGenerated bool  `json:"do_invoice_label_generated"`
            BatchID string  `json:"batch_id"`
            StoreName string  `json:"store_name"`
            Invoice map[string]interface{}  `json:"invoice"`
         
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
    
    // bulkListingData used by Order
    type bulkListingData struct {

        
            ExcelURL string  `json:"excel_url"`
            FailedShipments []map[string]interface{}  `json:"failed_shipments"`
            Status string  `json:"status"`
            Processing float64  `json:"processing"`
            Successful float64  `json:"successful"`
            Total float64  `json:"total"`
            BatchID string  `json:"batch_id"`
            ProcessingShipments []string  `json:"processing_shipments"`
            UserID string  `json:"user_id"`
            ID string  `json:"id"`
            CompanyID float64  `json:"company_id"`
            StoreCode string  `json:"store_code"`
            UploadedOn string  `json:"uploaded_on"`
            UserName string  `json:"user_name"`
            SuccessfulShipments []map[string]interface{}  `json:"successful_shipments"`
            FileName string  `json:"file_name"`
            Failed float64  `json:"failed"`
            StoreID float64  `json:"store_id"`
            StoreName string  `json:"store_name"`
         
    }
    
    // BulkListingPage used by Order
    type BulkListingPage struct {

        
            Current float64  `json:"current"`
            HasNext bool  `json:"has_next"`
            HasPrevious bool  `json:"has_previous"`
            Type string  `json:"type"`
            Total float64  `json:"total"`
            Size float64  `json:"size"`
         
    }
    
    // BulkListingResponse used by Order
    type BulkListingResponse struct {

        
            Success bool  `json:"success"`
            Data []bulkListingData  `json:"data"`
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

        
            QcType []string  `json:"qc_type"`
            QuestionSet []QuestionSet  `json:"question_set"`
            ID float64  `json:"id"`
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

        
            Status bool  `json:"status"`
            Message string  `json:"message"`
         
    }
    
    // BulkActionDetailsDataField used by Order
    type BulkActionDetailsDataField struct {

        
            TotalShipmentsCount float64  `json:"total_shipments_count"`
            FailedShipmentsCount float64  `json:"failed_shipments_count"`
            SuccessfulShipmentIds []string  `json:"successful_shipment_ids"`
            SuccessfulShipmentsCount float64  `json:"successful_shipments_count"`
            CompanyID string  `json:"company_id"`
            BatchID string  `json:"batch_id"`
            ProcessingShipmentsCount float64  `json:"processing_shipments_count"`
         
    }
    
    // BulkActionDetailsResponse used by Order
    type BulkActionDetailsResponse struct {

        
            Data []BulkActionDetailsDataField  `json:"data"`
            UploadedOn string  `json:"uploaded_on"`
            Success string  `json:"success"`
            UploadedBy string  `json:"uploaded_by"`
            UserID string  `json:"user_id"`
            FailedRecords []string  `json:"failed_records"`
            Status bool  `json:"status"`
            Message string  `json:"message"`
            Error []string  `json:"error"`
         
    }
    
    // BagGSTDetails used by Order
    type BagGSTDetails struct {

        
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            GstTag string  `json:"gst_tag"`
            SgstTaxPercentage float64  `json:"sgst_tax_percentage"`
            SgstGstFee string  `json:"sgst_gst_fee"`
            IgstTaxPercentage float64  `json:"igst_tax_percentage"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            GstinCode string  `json:"gstin_code"`
            CgstGstFee string  `json:"cgst_gst_fee"`
            CgstTaxPercentage float64  `json:"cgst_tax_percentage"`
            IgstGstFee string  `json:"igst_gst_fee"`
            HsnCode string  `json:"hsn_code"`
            GstFee float64  `json:"gst_fee"`
            ValueOfGood float64  `json:"value_of_good"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            IsDefaultHsnCode bool  `json:"is_default_hsn_code"`
            HsnCodeID string  `json:"hsn_code_id"`
         
    }
    
    // Brand used by Order
    type Brand struct {

        
            Company string  `json:"company"`
            BrandName string  `json:"brand_name"`
            PickupLocation string  `json:"pickup_location"`
            ScriptLastRan string  `json:"script_last_ran"`
            CreditNoteExpiryDays float64  `json:"credit_note_expiry_days"`
            CreditNoteAllowed bool  `json:"credit_note_allowed"`
            BrandID float64  `json:"brand_id"`
            CreatedOn float64  `json:"created_on"`
            IsVirtualInvoice bool  `json:"is_virtual_invoice"`
            StartDate string  `json:"start_date"`
            InvoicePrefix string  `json:"invoice_prefix"`
            Logo string  `json:"logo"`
            ModifiedOn float64  `json:"modified_on"`
         
    }
    
    // Attributes used by Order
    type Attributes struct {

        
            PrimaryColorHex string  `json:"primary_color_hex"`
            BrandName string  `json:"brand_name"`
            PrimaryMaterial string  `json:"primary_material"`
            PrimaryColor string  `json:"primary_color"`
            Gender []string  `json:"gender"`
            Essential string  `json:"essential"`
            Name string  `json:"name"`
            MarketerAddress string  `json:"marketer_address"`
            MarketerName string  `json:"marketer_name"`
         
    }
    
    // Item used by Order
    type Item struct {

        
            L3CategoryName string  `json:"l3_category_name"`
            Code string  `json:"code"`
            Image []string  `json:"image"`
            L1Category []string  `json:"l1_category"`
            WebstoreProductURL string  `json:"webstore_product_url"`
            Size string  `json:"size"`
            ItemID float64  `json:"item_id"`
            L2CategoryID float64  `json:"l2_category_id"`
            L2Category []string  `json:"l2_category"`
            Gender string  `json:"gender"`
            L1CategoryID float64  `json:"l1_category_id"`
            Brand string  `json:"brand"`
            SlugKey string  `json:"slug_key"`
            CanReturn bool  `json:"can_return"`
            Attributes Attributes  `json:"attributes"`
            Color string  `json:"color"`
            LastUpdatedAt string  `json:"last_updated_at"`
            BrandID float64  `json:"brand_id"`
            L3Category float64  `json:"l3_category"`
            DepartmentID float64  `json:"department_id"`
            Meta map[string]interface{}  `json:"meta"`
            CanCancel bool  `json:"can_cancel"`
            BranchURL string  `json:"branch_url"`
            Name string  `json:"name"`
         
    }
    
    // ArticleDetails used by Order
    type ArticleDetails struct {

        
            Status map[string]interface{}  `json:"status"`
         
    }
    
    // StoreAddress used by Order
    type StoreAddress struct {

        
            Pincode float64  `json:"pincode"`
            Version string  `json:"version"`
            UpdatedAt string  `json:"updated_at"`
            City string  `json:"city"`
            Address1 string  `json:"address1"`
            CreatedAt string  `json:"created_at"`
            CountryCode string  `json:"country_code"`
            Landmark string  `json:"landmark"`
            Longitude float64  `json:"longitude"`
            State string  `json:"state"`
            ContactPerson string  `json:"contact_person"`
            Country string  `json:"country"`
            Latitude float64  `json:"latitude"`
            Address2 string  `json:"address2"`
            Email string  `json:"email"`
            AddressType string  `json:"address_type"`
            Area string  `json:"area"`
            AddressCategory string  `json:"address_category"`
            Phone string  `json:"phone"`
         
    }
    
    // Document used by Order
    type Document struct {

        
            Verified bool  `json:"verified"`
            URL string  `json:"url"`
            DsType string  `json:"ds_type"`
            Value string  `json:"value"`
            LegalName string  `json:"legal_name"`
         
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
            Enabled bool  `json:"enabled"`
            User string  `json:"user"`
            Username string  `json:"username"`
         
    }
    
    // StoreGstCredentials used by Order
    type StoreGstCredentials struct {

        
            EWaybill StoreEwaybill  `json:"e_waybill"`
            EInvoice StoreEinvoice  `json:"e_invoice"`
         
    }
    
    // EInvoicePortalDetails used by Order
    type EInvoicePortalDetails struct {

        
            Password string  `json:"password"`
            User string  `json:"user"`
            Username string  `json:"username"`
         
    }
    
    // StoreMeta used by Order
    type StoreMeta struct {

        
            ProductReturnConfig map[string]interface{}  `json:"product_return_config"`
            Documents StoreDocuments  `json:"documents"`
            GstNumber string  `json:"gst_number"`
            Stage string  `json:"stage"`
            GstCredentials StoreGstCredentials  `json:"gst_credentials"`
            EinvoicePortalDetails EInvoicePortalDetails  `json:"einvoice_portal_details"`
            Timing []map[string]interface{}  `json:"timing"`
            AdditionalContactDetails map[string]interface{}  `json:"additional_contact_details"`
            NotificationEmails []string  `json:"notification_emails"`
            EwaybillPortalDetails map[string]interface{}  `json:"ewaybill_portal_details"`
            DisplayName string  `json:"display_name"`
         
    }
    
    // Store used by Order
    type Store struct {

        
            ParentStoreID float64  `json:"parent_store_id"`
            Pincode string  `json:"pincode"`
            StoreActiveFrom string  `json:"store_active_from"`
            BrandStoreTags []string  `json:"brand_store_tags"`
            FulfillmentChannel string  `json:"fulfillment_channel"`
            LocationType string  `json:"location_type"`
            Code string  `json:"code"`
            City string  `json:"city"`
            IsArchived bool  `json:"is_archived"`
            UpdatedAt string  `json:"updated_at"`
            VatNo string  `json:"vat_no"`
            Address1 string  `json:"address1"`
            CreatedAt string  `json:"created_at"`
            PackagingMaterialCount float64  `json:"packaging_material_count"`
            IsEnabledForRecon bool  `json:"is_enabled_for_recon"`
            Longitude float64  `json:"longitude"`
            State string  `json:"state"`
            ContactPerson string  `json:"contact_person"`
            CompanyID float64  `json:"company_id"`
            SID string  `json:"s_id"`
            MallName string  `json:"mall_name"`
            Country string  `json:"country"`
            StoreAddressJson StoreAddress  `json:"store_address_json"`
            MallArea string  `json:"mall_area"`
            Latitude float64  `json:"latitude"`
            StoreEmail string  `json:"store_email"`
            BrandID interface{}  `json:"brand_id"`
            Address2 string  `json:"address2"`
            AlohomoraUserID float64  `json:"alohomora_user_id"`
            Meta StoreMeta  `json:"meta"`
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            LoginUsername string  `json:"login_username"`
            OrderIntegrationID string  `json:"order_integration_id"`
            Phone float64  `json:"phone"`
         
    }
    
    // BagDetailsPlatformResponse used by Order
    type BagDetailsPlatformResponse struct {

        
            CurrentStatus BagStatusHistory  `json:"current_status"`
            LineNumber float64  `json:"line_number"`
            Quantity float64  `json:"quantity"`
            BagUpdateTime float64  `json:"bag_update_time"`
            Status BagReturnableCancelableStatus  `json:"status"`
            Identifier string  `json:"identifier"`
            EntityType string  `json:"entity_type"`
            CurrentOperationalStatus BagStatusHistory  `json:"current_operational_status"`
            Prices Prices  `json:"prices"`
            ParentPromoBags map[string]interface{}  `json:"parent_promo_bags"`
            DisplayName string  `json:"display_name"`
            GstDetails BagGSTDetails  `json:"gst_details"`
            FinancialBreakup []FinancialBreakup  `json:"financial_breakup"`
            BType string  `json:"b_type"`
            Brand Brand  `json:"brand"`
            AppliedPromos []map[string]interface{}  `json:"applied_promos"`
            Item Item  `json:"item"`
            Article Article  `json:"article"`
            JourneyType string  `json:"journey_type"`
            OriginalBagList []float64  `json:"original_bag_list"`
            Reasons []map[string]interface{}  `json:"reasons"`
            RestoreCoupon bool  `json:"restore_coupon"`
            NoOfBagsOrder float64  `json:"no_of_bags_order"`
            ShipmentID string  `json:"shipment_id"`
            SellerIdentifier string  `json:"seller_identifier"`
            AffiliateDetails AffiliateDetails  `json:"affiliate_details"`
            Dates Dates  `json:"dates"`
            QcRequired interface{}  `json:"qc_required"`
            ArticleDetails ArticleDetails  `json:"article_details"`
            BagStatus []BagStatusHistory  `json:"bag_status"`
            Meta BagMeta  `json:"meta"`
            OperationalStatus string  `json:"operational_status"`
            OrderingStore Store  `json:"ordering_store"`
            BagStatusHistory BagStatusHistory  `json:"bag_status_history"`
            RestorePromos map[string]interface{}  `json:"restore_promos"`
            OrderIntegrationID string  `json:"order_integration_id"`
            AffiliateBagDetails AffiliateBagDetails  `json:"affiliate_bag_details"`
            BID float64  `json:"b_id"`
            Tags []string  `json:"tags"`
         
    }
    
    // ErrorResponse used by Order
    type ErrorResponse struct {

        
            Error string  `json:"error"`
            Message string  `json:"message"`
         
    }
    
    // Page1 used by Order
    type Page1 struct {

        
            Current float64  `json:"current"`
            HasNext bool  `json:"has_next"`
            Size float64  `json:"size"`
            PageType string  `json:"page_type"`
            ItemTotal float64  `json:"item_total"`
         
    }
    
    // GetBagsPlatformResponse used by Order
    type GetBagsPlatformResponse struct {

        
            Items []BagDetailsPlatformResponse  `json:"items"`
            Page Page1  `json:"page"`
         
    }
    
    // InvalidateShipmentCachePayload used by Order
    type InvalidateShipmentCachePayload struct {

        
            BagIds []string  `json:"bag_ids"`
            ShipmentIds []string  `json:"shipment_ids"`
            AffiliateBagIds []string  `json:"affiliate_bag_ids"`
         
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

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
            ErrorTrace string  `json:"error_trace"`
            Status float64  `json:"status"`
         
    }
    
    // StoreReassign used by Order
    type StoreReassign struct {

        
            MongoArticleID string  `json:"mongo_article_id"`
            AffiliateID string  `json:"affiliate_id"`
            SetID string  `json:"set_id"`
            ItemID string  `json:"item_id"`
            StoreID float64  `json:"store_id"`
            ReasonIds []float64  `json:"reason_ids"`
            FyndOrderID string  `json:"fynd_order_id"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
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
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            ID string  `json:"id"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
            AffiliateID string  `json:"affiliate_id"`
         
    }
    
    // UpdateShipmentLockPayload used by Order
    type UpdateShipmentLockPayload struct {

        
            Entities []Entities  `json:"entities"`
            ActionType string  `json:"action_type"`
            EntityType string  `json:"entity_type"`
            Action string  `json:"action"`
         
    }
    
    // Bags used by Order
    type Bags struct {

        
            IsLocked bool  `json:"is_locked"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
            BagID float64  `json:"bag_id"`
         
    }
    
    // OriginalFilter used by Order
    type OriginalFilter struct {

        
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            AffiliateID string  `json:"affiliate_id"`
         
    }
    
    // CheckResponse used by Order
    type CheckResponse struct {

        
            Bags []Bags  `json:"bags"`
            LockStatus string  `json:"lock_status"`
            IsShipmentLocked bool  `json:"is_shipment_locked"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            IsBagLocked bool  `json:"is_bag_locked"`
            Status string  `json:"status"`
            ShipmentID string  `json:"shipment_id"`
            OriginalFilter OriginalFilter  `json:"original_filter"`
            AffiliateID string  `json:"affiliate_id"`
         
    }
    
    // UpdateShipmentLockResponse used by Order
    type UpdateShipmentLockResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
            CheckResponse []CheckResponse  `json:"check_response"`
         
    }
    
    // AnnouncementResponse used by Order
    type AnnouncementResponse struct {

        
            ToDatetime string  `json:"to_datetime"`
            CompanyID float64  `json:"company_id"`
            Title string  `json:"title"`
            LogoURL string  `json:"logo_url"`
            PlatformID string  `json:"platform_id"`
            CreatedAt string  `json:"created_at"`
            ID float64  `json:"id"`
            FromDatetime string  `json:"from_datetime"`
            PlatformName string  `json:"platform_name"`
            Description string  `json:"description"`
         
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

        
            ForceTransition bool  `json:"force_transition"`
            Task bool  `json:"task"`
            UnlockBeforeTransition bool  `json:"unlock_before_transition"`
            Statuses []StatuesRequest  `json:"statuses"`
            LockAfterTransition bool  `json:"lock_after_transition"`
         
    }
    
    // ShipmentsResponse used by Order
    type ShipmentsResponse struct {

        
            Meta map[string]interface{}  `json:"meta"`
            Exception string  `json:"exception"`
            StackTrace string  `json:"stack_trace"`
            Code string  `json:"code"`
            Status float64  `json:"status"`
            FinalState map[string]interface{}  `json:"final_state"`
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
    
    // MarketPlacePdf used by Order
    type MarketPlacePdf struct {

        
            Label string  `json:"label"`
            Invoice string  `json:"invoice"`
         
    }
    
    // AffiliateBag used by Order
    type AffiliateBag struct {

        
            AvlQty float64  `json:"avl_qty"`
            CompanyID float64  `json:"company_id"`
            DeliveryCharge float64  `json:"delivery_charge"`
            ID string  `json:"_id"`
            ItemID float64  `json:"item_id"`
            PriceMarked float64  `json:"price_marked"`
            HsnCodeID string  `json:"hsn_code_id"`
            AmountPaid float64  `json:"amount_paid"`
            ItemSize string  `json:"item_size"`
            StoreID float64  `json:"store_id"`
            TransferPrice float64  `json:"transfer_price"`
            PriceEffective float64  `json:"price_effective"`
            AffiliateStoreID string  `json:"affiliate_store_id"`
            AffiliateMeta map[string]interface{}  `json:"affiliate_meta"`
            SellerIdentifier string  `json:"seller_identifier"`
            Quantity float64  `json:"quantity"`
            ModifiedOn string  `json:"modified_on"`
            Sku string  `json:"sku"`
            PdfLinks MarketPlacePdf  `json:"pdf_links"`
            Identifier map[string]interface{}  `json:"identifier"`
            FyndStoreID string  `json:"fynd_store_id"`
            UnitPrice float64  `json:"unit_price"`
            Discount float64  `json:"discount"`
         
    }
    
    // OrderUser used by Order
    type OrderUser struct {

        
            Address1 string  `json:"address1"`
            Country string  `json:"country"`
            Phone float64  `json:"phone"`
            State string  `json:"state"`
            Email string  `json:"email"`
            Mobile float64  `json:"mobile"`
            City string  `json:"city"`
            Address2 string  `json:"address2"`
            Pincode string  `json:"pincode"`
            LastName string  `json:"last_name"`
            FirstName string  `json:"first_name"`
         
    }
    
    // OrderPriority used by Order
    type OrderPriority struct {

        
            AffiliatePriorityCode string  `json:"affiliate_priority_code"`
            FulfilmentPriority float64  `json:"fulfilment_priority"`
            FulfilmentPriorityText string  `json:"fulfilment_priority_text"`
         
    }
    
    // UserData used by Order
    type UserData struct {

        
            BillingUser OrderUser  `json:"billing_user"`
            ShippingUser OrderUser  `json:"shipping_user"`
         
    }
    
    // ArticleDetails1 used by Order
    type ArticleDetails1 struct {

        
            Weight map[string]interface{}  `json:"weight"`
            Quantity float64  `json:"quantity"`
            Attributes map[string]interface{}  `json:"attributes"`
            Category map[string]interface{}  `json:"category"`
            ID string  `json:"_id"`
            BrandID float64  `json:"brand_id"`
            Dimension map[string]interface{}  `json:"dimension"`
         
    }
    
    // LocationDetails used by Order
    type LocationDetails struct {

        
            Articles []ArticleDetails1  `json:"articles"`
            FulfillmentID float64  `json:"fulfillment_id"`
            FulfillmentType string  `json:"fulfillment_type"`
         
    }
    
    // ShipmentDetails1 used by Order
    type ShipmentDetails1 struct {

        
            DpID float64  `json:"dp_id"`
            Shipments float64  `json:"shipments"`
            Meta map[string]interface{}  `json:"meta"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            FulfillmentID float64  `json:"fulfillment_id"`
            BoxType string  `json:"box_type"`
            Articles []ArticleDetails1  `json:"articles"`
         
    }
    
    // ShipmentConfig used by Order
    type ShipmentConfig struct {

        
            Journey string  `json:"journey"`
            Source string  `json:"source"`
            LocationDetails LocationDetails  `json:"location_details"`
            Shipment []ShipmentDetails1  `json:"shipment"`
            ToPincode string  `json:"to_pincode"`
            PaymentMode string  `json:"payment_mode"`
            Identifier string  `json:"identifier"`
            Action string  `json:"action"`
         
    }
    
    // ShipmentData used by Order
    type ShipmentData struct {

        
            ShipmentData ShipmentConfig  `json:"shipment_data"`
         
    }
    
    // OrderInfo used by Order
    type OrderInfo struct {

        
            Bags []AffiliateBag  `json:"bags"`
            BillingAddress OrderUser  `json:"billing_address"`
            OrderPriority OrderPriority  `json:"order_priority"`
            CodCharges float64  `json:"cod_charges"`
            Payment map[string]interface{}  `json:"payment"`
            ShippingAddress OrderUser  `json:"shipping_address"`
            Items map[string]interface{}  `json:"items"`
            User UserData  `json:"user"`
            Shipment ShipmentData  `json:"shipment"`
            PaymentMode string  `json:"payment_mode"`
            OrderValue float64  `json:"order_value"`
            DeliveryCharges float64  `json:"delivery_charges"`
            Coupon string  `json:"coupon"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            Discount float64  `json:"discount"`
         
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
    
    // AffiliateInventoryLogisticsConfig used by Order
    type AffiliateInventoryLogisticsConfig struct {

        
            DpAssignment bool  `json:"dp_assignment"`
         
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
    
    // AffiliateInventoryConfig used by Order
    type AffiliateInventoryConfig struct {

        
            Payment AffiliateInventoryPaymentConfig  `json:"payment"`
            Logistics AffiliateInventoryLogisticsConfig  `json:"logistics"`
            ArticleAssignment AffiliateInventoryArticleAssignmentConfig  `json:"article_assignment"`
            Order AffiliateInventoryOrderConfig  `json:"order"`
            Inventory AffiliateInventoryStoreConfig  `json:"inventory"`
         
    }
    
    // AffiliateAppConfigMeta used by Order
    type AffiliateAppConfigMeta struct {

        
            Value string  `json:"value"`
            Name string  `json:"name"`
         
    }
    
    // AffiliateAppConfig used by Order
    type AffiliateAppConfig struct {

        
            UpdatedAt string  `json:"updated_at"`
            Owner string  `json:"owner"`
            Meta []AffiliateAppConfigMeta  `json:"meta"`
            Name string  `json:"name"`
            Secret string  `json:"secret"`
            CreatedAt string  `json:"created_at"`
            ID string  `json:"id"`
            Description string  `json:"description"`
            Token string  `json:"token"`
         
    }
    
    // AffiliateConfig used by Order
    type AffiliateConfig struct {

        
            Inventory AffiliateInventoryConfig  `json:"inventory"`
            App AffiliateAppConfig  `json:"app"`
         
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
            BagEndState string  `json:"bag_end_state"`
            ArticleLookup string  `json:"article_lookup"`
            StoreLookup string  `json:"store_lookup"`
            CreateUser bool  `json:"create_user"`
            Affiliate Affiliate  `json:"affiliate"`
         
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

        
            Slug string  `json:"slug"`
            ID float64  `json:"id"`
            Description string  `json:"description"`
            DisplayText string  `json:"display_text"`
         
    }
    
    // GetActionsResponse used by Order
    type GetActionsResponse struct {

        
            Permissions ActionInfo  `json:"permissions"`
         
    }
    
    // HistoryReason used by Order
    type HistoryReason struct {

        
            Quantity float64  `json:"quantity"`
            State string  `json:"state"`
            Category string  `json:"category"`
            DislayName string  `json:"dislay_name"`
            Code float64  `json:"code"`
            Text string  `json:"text"`
         
    }
    
    // HistoryMeta used by Order
    type HistoryMeta struct {

        
            Billsec string  `json:"billsec"`
            Message string  `json:"message"`
            ShortLink string  `json:"short_link"`
            StoreName string  `json:"store_name"`
            Recordpath string  `json:"recordpath"`
            Status2 string  `json:"status2"`
            Duration string  `json:"duration"`
            ActivityType string  `json:"activity_type"`
            ActivityComment string  `json:"activity_comment"`
            StoreID float64  `json:"store_id"`
            StoreCode string  `json:"store_code"`
            CallID string  `json:"call_id"`
            Receiver string  `json:"receiver"`
            Reason HistoryReason  `json:"reason"`
            ChannelType string  `json:"channel_type"`
            Status1 string  `json:"status1"`
            Starttime string  `json:"starttime"`
            Callerid string  `json:"callerid"`
            Endtime string  `json:"endtime"`
            Status string  `json:"status"`
            Caller string  `json:"caller"`
            Slug string  `json:"slug"`
            Recipient string  `json:"recipient"`
         
    }
    
    // HistoryDict used by Order
    type HistoryDict struct {

        
            Type string  `json:"type"`
            L2Detail string  `json:"l2_detail"`
            AssignedAgent string  `json:"assigned_agent"`
            DisplayMessage string  `json:"display_message"`
            L3Detail string  `json:"l3_detail"`
            Meta HistoryMeta  `json:"meta"`
            TicketURL string  `json:"ticket_url"`
            Createdat string  `json:"createdat"`
            User string  `json:"user"`
            L1Detail string  `json:"l1_detail"`
            Message string  `json:"message"`
            TicketID string  `json:"ticket_id"`
            BagID float64  `json:"bag_id"`
         
    }
    
    // ShipmentHistoryResponse used by Order
    type ShipmentHistoryResponse struct {

        
            Success bool  `json:"success"`
            ActivityHistory []HistoryDict  `json:"activity_history"`
         
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
    
    // SmsDataPayload used by Order
    type SmsDataPayload struct {

        
            AmountPaid float64  `json:"amount_paid"`
            OrderID string  `json:"order_id"`
            CustomerName string  `json:"customer_name"`
            CountryCode string  `json:"country_code"`
            PaymentMode string  `json:"payment_mode"`
            Message string  `json:"message"`
            BrandName string  `json:"brand_name"`
            ShipmentID float64  `json:"shipment_id"`
            PhoneNumber float64  `json:"phone_number"`
         
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
            ID float64  `json:"id"`
            Status string  `json:"status"`
            Remarks string  `json:"remarks"`
            ShipmentID string  `json:"shipment_id"`
            BagList []float64  `json:"bag_list"`
         
    }
    
    // OrderDetails used by Order
    type OrderDetails struct {

        
            CreatedAt string  `json:"created_at"`
            FyndOrderID string  `json:"fynd_order_id"`
         
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

        
            Weight string  `json:"weight"`
            Height string  `json:"height"`
            Length float64  `json:"length"`
            PackagingType string  `json:"packaging_type"`
            Width float64  `json:"width"`
         
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
    
    // Tax used by Order
    type Tax struct {

        
            Amount map[string]interface{}  `json:"amount"`
            Name string  `json:"name"`
            Rate float64  `json:"rate"`
            Breakup []map[string]interface{}  `json:"breakup"`
         
    }
    
    // Charge used by Order
    type Charge struct {

        
            Type string  `json:"type"`
            Name string  `json:"name"`
            Amount map[string]interface{}  `json:"amount"`
            Code string  `json:"code"`
            Tax Tax  `json:"tax"`
         
    }
    
    // TaxInfo used by Order
    type TaxInfo struct {

        
            Gstin string  `json:"gstin"`
            B2bGstinNumber string  `json:"b2b_gstin_number"`
         
    }
    
    // ProcessingDates used by Order
    type ProcessingDates struct {

        
            DispatchByDate string  `json:"dispatch_by_date"`
            DpPickupSlot map[string]interface{}  `json:"dp_pickup_slot"`
            ConfirmByDate string  `json:"confirm_by_date"`
            PackByDate string  `json:"pack_by_date"`
            DispatchAfterDate string  `json:"dispatch_after_date"`
            CustomerPickupSlot map[string]interface{}  `json:"customer_pickup_slot"`
         
    }
    
    // LineItem used by Order
    type LineItem struct {

        
            CustomMessasge string  `json:"custom_messasge"`
            Charges []Charge  `json:"charges"`
            SellerIdentifier string  `json:"seller_identifier"`
            Quantity float64  `json:"quantity"`
            Meta map[string]interface{}  `json:"meta"`
            ExternalLineID string  `json:"external_line_id"`
         
    }
    
    // Shipment used by Order
    type Shipment struct {

        
            ProcessingDates ProcessingDates  `json:"processing_dates"`
            LineItems []LineItem  `json:"line_items"`
            Meta map[string]interface{}  `json:"meta"`
            LocationID float64  `json:"location_id"`
            Priority float64  `json:"priority"`
            ExternalShipmentID string  `json:"external_shipment_id"`
         
    }
    
    // BillingInfo used by Order
    type BillingInfo struct {

        
            PrimaryEmail string  `json:"primary_email"`
            State string  `json:"state"`
            Address2 string  `json:"address2"`
            HouseNo string  `json:"house_no"`
            MiddleName string  `json:"middle_name"`
            Address1 string  `json:"address1"`
            StateCode string  `json:"state_code"`
            AlternateMobileNumber string  `json:"alternate_mobile_number"`
            CustomerCode string  `json:"customer_code"`
            Pincode string  `json:"pincode"`
            Country string  `json:"country"`
            FloorNo string  `json:"floor_no"`
            Gender string  `json:"gender"`
            CountryCode string  `json:"country_code"`
            ExternalCustomerCode string  `json:"external_customer_code"`
            FirstName string  `json:"first_name"`
            Title string  `json:"title"`
            PrimaryMobileNumber string  `json:"primary_mobile_number"`
            City string  `json:"city"`
            LastName string  `json:"last_name"`
            AlternateEmail string  `json:"alternate_email"`
         
    }
    
    // ShippingInfo used by Order
    type ShippingInfo struct {

        
            PrimaryEmail string  `json:"primary_email"`
            State string  `json:"state"`
            Slot []map[string]interface{}  `json:"slot"`
            Address2 string  `json:"address2"`
            HouseNo string  `json:"house_no"`
            MiddleName string  `json:"middle_name"`
            Address1 string  `json:"address1"`
            StateCode string  `json:"state_code"`
            AlternateMobileNumber string  `json:"alternate_mobile_number"`
            GeoLocation map[string]interface{}  `json:"geo_location"`
            CustomerCode string  `json:"customer_code"`
            Landmark string  `json:"landmark"`
            Pincode string  `json:"pincode"`
            Country string  `json:"country"`
            FloorNo string  `json:"floor_no"`
            Gender string  `json:"gender"`
            CountryCode string  `json:"country_code"`
            ExternalCustomerCode string  `json:"external_customer_code"`
            FirstName string  `json:"first_name"`
            Title string  `json:"title"`
            ShippingType string  `json:"shipping_type"`
            PrimaryMobileNumber string  `json:"primary_mobile_number"`
            City string  `json:"city"`
            AddressType string  `json:"address_type"`
            LastName string  `json:"last_name"`
            AlternateEmail string  `json:"alternate_email"`
         
    }
    
    // PaymentMethod1 used by Order
    type PaymentMethod1 struct {

        
            RefundBy string  `json:"refund_by"`
            Mode string  `json:"mode"`
            Meta map[string]interface{}  `json:"meta"`
            Name string  `json:"name"`
            Amount float64  `json:"amount"`
            TransactionData map[string]interface{}  `json:"transaction_data"`
            CollectBy string  `json:"collect_by"`
         
    }
    
    // PaymentInfo used by Order
    type PaymentInfo struct {

        
            PaymentMethods []PaymentMethod1  `json:"payment_methods"`
            PrimaryMode string  `json:"primary_mode"`
         
    }
    
    // CreateOrderAPI used by Order
    type CreateOrderAPI struct {

        
            Charges []Charge  `json:"charges"`
            TaxInfo TaxInfo  `json:"tax_info"`
            Config map[string]interface{}  `json:"config"`
            Shipments []Shipment  `json:"shipments"`
            ExternalOrderID string  `json:"external_order_id"`
            Meta map[string]interface{}  `json:"meta"`
            CurrencyInfo map[string]interface{}  `json:"currency_info"`
            BillingInfo BillingInfo  `json:"billing_info"`
            ShippingInfo ShippingInfo  `json:"shipping_info"`
            PaymentInfo PaymentInfo  `json:"payment_info"`
            ExternalCreationDate string  `json:"external_creation_date"`
         
    }
    
    // CreateOrderErrorReponse used by Order
    type CreateOrderErrorReponse struct {

        
            RequestID string  `json:"request_id"`
            Meta string  `json:"meta"`
            Exception string  `json:"exception"`
            StackTrace string  `json:"stack_trace"`
            Code string  `json:"code"`
            Status float64  `json:"status"`
            Info interface{}  `json:"info"`
            Message string  `json:"message"`
         
    }
    
    // DpConfiguration used by Order
    type DpConfiguration struct {

        
            ShippingBy string  `json:"shipping_by"`
         
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
    
    // CreateChannelConfig used by Order
    type CreateChannelConfig struct {

        
            ShipmentAssignment string  `json:"shipment_assignment"`
            DpConfiguration DpConfiguration  `json:"dp_configuration"`
            LogoURL map[string]interface{}  `json:"logo_url"`
            LockStates []string  `json:"lock_states"`
            PaymentInfo CreateChannelPaymentInfo  `json:"payment_info"`
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

        
            Acknowledged bool  `json:"acknowledged"`
            IsInserted bool  `json:"is_inserted"`
            IsUpserted bool  `json:"is_upserted"`
         
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
            EndDate string  `json:"end_date"`
            StartDate string  `json:"start_date"`
            OrderDetails []FyndOrderIdList  `json:"order_details"`
         
    }
    
    // BagStateTransitionMap used by Order
    type BagStateTransitionMap struct {

        
            Affiliate map[string]interface{}  `json:"affiliate"`
            Fynd map[string]interface{}  `json:"fynd"`
         
    }
    

    
    // SearchKeywordResult used by Catalog
    type SearchKeywordResult struct {

        
            Query map[string]interface{}  `json:"query"`
            SortOn string  `json:"sort_on"`
         
    }
    
    // CreateSearchKeyword used by Catalog
    type CreateSearchKeyword struct {

        
            Result SearchKeywordResult  `json:"result"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            IsActive bool  `json:"is_active"`
            AppID string  `json:"app_id"`
            Words []string  `json:"words"`
         
    }
    
    // SearchKeywordResult1 used by Catalog
    type SearchKeywordResult1 struct {

        
            Query map[string]interface{}  `json:"query"`
            SortOn string  `json:"sort_on"`
         
    }
    
    // GetSearchWordsData used by Catalog
    type GetSearchWordsData struct {

        
            Result SearchKeywordResult1  `json:"result"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            UID string  `json:"uid"`
            IsActive bool  `json:"is_active"`
            AppID string  `json:"app_id"`
            Words []string  `json:"words"`
         
    }
    
    // SearchErrorResponse used by Catalog
    type SearchErrorResponse struct {

        
            Errors []map[string]interface{}  `json:"errors"`
            Code float64  `json:"code"`
            Message string  `json:"message"`
         
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

        
            Items []GetSearchWordsData  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // AutocompletePageAction used by Catalog
    type AutocompletePageAction struct {

        
            Query map[string]interface{}  `json:"query"`
            Params map[string]interface{}  `json:"params"`
            Type string  `json:"type"`
            URL string  `json:"url"`
         
    }
    
    // AutocompleteAction used by Catalog
    type AutocompleteAction struct {

        
            Type string  `json:"type"`
            Page AutocompletePageAction  `json:"page"`
         
    }
    
    // Media used by Catalog
    type Media struct {

        
            AspectRatio string  `json:"aspect_ratio"`
            Type string  `json:"type"`
            URL string  `json:"url"`
         
    }
    
    // AutocompleteResult used by Catalog
    type AutocompleteResult struct {

        
            Action AutocompleteAction  `json:"action"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Display string  `json:"display"`
            Logo Media  `json:"logo"`
         
    }
    
    // GetAutocompleteWordsData used by Catalog
    type GetAutocompleteWordsData struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            UID string  `json:"uid"`
            IsActive bool  `json:"is_active"`
            AppID string  `json:"app_id"`
            Words []string  `json:"words"`
            Results []AutocompleteResult  `json:"results"`
         
    }
    
    // UpdateAutocompleteWordData used by Catalog
    type UpdateAutocompleteWordData struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            UID string  `json:"uid"`
            IsActive bool  `json:"is_active"`
            AppID string  `json:"app_id"`
            Words []string  `json:"words"`
            Results []AutocompleteResult  `json:"results"`
            ID string  `json:"_id"`
         
    }
    
    // GetAutocompleteWordsResponse used by Catalog
    type GetAutocompleteWordsResponse struct {

        
            Items []GetAutocompleteWordsData  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // AutocompletePageAction1 used by Catalog
    type AutocompletePageAction1 struct {

        
            Query map[string]interface{}  `json:"query"`
            Params map[string]interface{}  `json:"params"`
            Type string  `json:"type"`
            URL string  `json:"url"`
         
    }
    
    // AutocompleteAction1 used by Catalog
    type AutocompleteAction1 struct {

        
            Type string  `json:"type"`
            Page AutocompletePageAction1  `json:"page"`
         
    }
    
    // Media1 used by Catalog
    type Media1 struct {

        
            AspectRatio string  `json:"aspect_ratio"`
            Type string  `json:"type"`
            URL string  `json:"url"`
         
    }
    
    // AutocompleteResult1 used by Catalog
    type AutocompleteResult1 struct {

        
            Action AutocompleteAction1  `json:"action"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Display string  `json:"display"`
            Logo Media1  `json:"logo"`
         
    }
    
    // CreateAutocompleteKeyword used by Catalog
    type CreateAutocompleteKeyword struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            IsActive bool  `json:"is_active"`
            AppID string  `json:"app_id"`
            Words []string  `json:"words"`
            Results []AutocompleteResult1  `json:"results"`
         
    }
    
    // RerankingAttribute used by Catalog
    type RerankingAttribute struct {

        
            AttributeValue string  `json:"attribute_value"`
            AttributeKey string  `json:"attribute_key"`
         
    }
    
    // BoostBury used by Catalog
    type BoostBury struct {

        
            Boost []RerankingAttribute  `json:"boost"`
         
    }
    
    // UserDetail used by Catalog
    type UserDetail struct {

        
            Contact string  `json:"contact"`
            UserID string  `json:"user_id"`
            SuperUser bool  `json:"super_user"`
            Username string  `json:"username"`
         
    }
    
    // SearchRerankingModel used by Catalog
    type SearchRerankingModel struct {

        
            Ranking BoostBury  `json:"ranking"`
            CreatedBy UserDetail  `json:"created_by"`
            VerifiedOn string  `json:"verified_on"`
            VerifiedBy UserDetail  `json:"verified_by"`
            ModifiedOn string  `json:"modified_on"`
            AppID string  `json:"app_id"`
            CreatedOn string  `json:"created_on"`
            Words []string  `json:"words"`
            IsActive bool  `json:"is_active"`
            ModifiedBy UserDetail  `json:"modified_by"`
            ID interface{}  `json:"_id"`
         
    }
    
    // SearchRerankListing used by Catalog
    type SearchRerankListing struct {

        
            Items []SearchRerankingModel  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // RerankingAttribute1 used by Catalog
    type RerankingAttribute1 struct {

        
            AttributeValue string  `json:"attribute_value"`
            AttributeKey string  `json:"attribute_key"`
         
    }
    
    // BoostBury1 used by Catalog
    type BoostBury1 struct {

        
            Boost []RerankingAttribute1  `json:"boost"`
         
    }
    
    // CreateSearchReranking used by Catalog
    type CreateSearchReranking struct {

        
            Ranking BoostBury1  `json:"ranking"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            IsActive bool  `json:"is_active"`
            AppID string  `json:"app_id"`
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
            Words []string  `json:"words"`
         
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
            MinQuantity float64  `json:"min_quantity"`
            ProductUID float64  `json:"product_uid"`
            AllowRemove bool  `json:"allow_remove"`
            AutoAddToCart bool  `json:"auto_add_to_cart"`
         
    }
    
    // GetProductBundleCreateResponse used by Catalog
    type GetProductBundleCreateResponse struct {

        
            CreatedBy map[string]interface{}  `json:"created_by"`
            CompanyID float64  `json:"company_id"`
            Meta map[string]interface{}  `json:"meta"`
            ID string  `json:"id"`
            Choice string  `json:"choice"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            IsActive bool  `json:"is_active"`
            ModifiedOn string  `json:"modified_on"`
            Slug string  `json:"slug"`
            PageVisibility []string  `json:"page_visibility"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            CreatedOn string  `json:"created_on"`
            Logo string  `json:"logo"`
            Products []ProductBundleItem  `json:"products"`
            Name string  `json:"name"`
         
    }
    
    // GetProductBundleListingResponse used by Catalog
    type GetProductBundleListingResponse struct {

        
            Items []GetProductBundleCreateResponse  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // ErrorResponse used by Catalog
    type ErrorResponse struct {

        
            Meta map[string]interface{}  `json:"meta"`
            Message string  `json:"message"`
            Error string  `json:"error"`
            Status float64  `json:"status"`
            Code string  `json:"code"`
         
    }
    
    // ProductBundleRequest used by Catalog
    type ProductBundleRequest struct {

        
            CreatedBy map[string]interface{}  `json:"created_by"`
            CompanyID float64  `json:"company_id"`
            Meta map[string]interface{}  `json:"meta"`
            Choice string  `json:"choice"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            IsActive bool  `json:"is_active"`
            ModifiedOn string  `json:"modified_on"`
            Slug string  `json:"slug"`
            PageVisibility []string  `json:"page_visibility"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            CreatedOn string  `json:"created_on"`
            Logo string  `json:"logo"`
            Products []ProductBundleItem  `json:"products"`
            Name string  `json:"name"`
         
    }
    
    // ProductBundleUpdateRequest used by Catalog
    type ProductBundleUpdateRequest struct {

        
            CompanyID float64  `json:"company_id"`
            Meta map[string]interface{}  `json:"meta"`
            Choice string  `json:"choice"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            IsActive bool  `json:"is_active"`
            ModifiedOn string  `json:"modified_on"`
            Slug string  `json:"slug"`
            PageVisibility []string  `json:"page_visibility"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            Logo string  `json:"logo"`
            Products []ProductBundleItem  `json:"products"`
            Name string  `json:"name"`
         
    }
    
    // LimitedProductData used by Catalog
    type LimitedProductData struct {

        
            UID float64  `json:"uid"`
            Attributes map[string]interface{}  `json:"attributes"`
            Quantity float64  `json:"quantity"`
            ShortDescription string  `json:"short_description"`
            ItemCode string  `json:"item_code"`
            Slug string  `json:"slug"`
            Sizes []string  `json:"sizes"`
            Name string  `json:"name"`
            Price map[string]interface{}  `json:"price"`
            Images []string  `json:"images"`
            CountryOfOrigin string  `json:"country_of_origin"`
            Identifier map[string]interface{}  `json:"identifier"`
         
    }
    
    // Size used by Catalog
    type Size struct {

        
            IsAvailable bool  `json:"is_available"`
            Value string  `json:"value"`
            Quantity float64  `json:"quantity"`
            Display string  `json:"display"`
         
    }
    
    // Price used by Catalog
    type Price struct {

        
            Currency string  `json:"currency"`
            MaxMarked float64  `json:"max_marked"`
            MinMarked float64  `json:"min_marked"`
            MaxEffective float64  `json:"max_effective"`
            MinEffective float64  `json:"min_effective"`
         
    }
    
    // GetProducts used by Catalog
    type GetProducts struct {

        
            ProductDetails LimitedProductData  `json:"product_details"`
            AutoSelect bool  `json:"auto_select"`
            MaxQuantity float64  `json:"max_quantity"`
            MinQuantity float64  `json:"min_quantity"`
            ProductUID float64  `json:"product_uid"`
            Sizes []Size  `json:"sizes"`
            Price Price  `json:"price"`
            AllowRemove bool  `json:"allow_remove"`
            AutoAddToCart bool  `json:"auto_add_to_cart"`
         
    }
    
    // GetProductBundleResponse used by Catalog
    type GetProductBundleResponse struct {

        
            CompanyID float64  `json:"company_id"`
            Meta map[string]interface{}  `json:"meta"`
            Choice string  `json:"choice"`
            PageVisibility []string  `json:"page_visibility"`
            IsActive bool  `json:"is_active"`
            Slug string  `json:"slug"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            Name string  `json:"name"`
            Products []GetProducts  `json:"products"`
            Logo string  `json:"logo"`
         
    }
    
    // ListSizeGuide used by Catalog
    type ListSizeGuide struct {

        
            Items []map[string]interface{}  `json:"items"`
            Page map[string]interface{}  `json:"page"`
         
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

        
            Active bool  `json:"active"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            CompanyID float64  `json:"company_id"`
            Subtitle string  `json:"subtitle"`
            BrandID float64  `json:"brand_id"`
            ID string  `json:"id"`
            Description string  `json:"description"`
            Title string  `json:"title"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
            Image string  `json:"image"`
            Tag string  `json:"tag"`
            Guide Guide  `json:"guide"`
            Name string  `json:"name"`
         
    }
    
    // SuccessResponse used by Catalog
    type SuccessResponse struct {

        
            Success bool  `json:"success"`
         
    }
    
    // SizeGuideResponse used by Catalog
    type SizeGuideResponse struct {

        
            Active bool  `json:"active"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            CompanyID float64  `json:"company_id"`
            Subtitle string  `json:"subtitle"`
            BrandID float64  `json:"brand_id"`
            ID string  `json:"id"`
            Title string  `json:"title"`
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
            Tag string  `json:"tag"`
            Guide map[string]interface{}  `json:"guide"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Name string  `json:"name"`
         
    }
    
    // MOQData used by Catalog
    type MOQData struct {

        
            Minimum float64  `json:"minimum"`
            IncrementUnit float64  `json:"increment_unit"`
            Maximum float64  `json:"maximum"`
         
    }
    
    // SEOData used by Catalog
    type SEOData struct {

        
            Description string  `json:"description"`
            Title string  `json:"title"`
         
    }
    
    // OwnerAppItemResponse used by Catalog
    type OwnerAppItemResponse struct {

        
            Moq MOQData  `json:"moq"`
            IsCod bool  `json:"is_cod"`
            AltText map[string]interface{}  `json:"alt_text"`
            IsGift bool  `json:"is_gift"`
            Seo SEOData  `json:"seo"`
         
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

        
            Description string  `json:"description"`
            Title string  `json:"title"`
         
    }
    
    // ApplicationItemMeta used by Catalog
    type ApplicationItemMeta struct {

        
            Moq ApplicationItemMOQ  `json:"moq"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            IsCod bool  `json:"is_cod"`
            AltText map[string]interface{}  `json:"alt_text"`
            CustomMeta []MetaFields  `json:"_custom_meta"`
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

        
            Values []map[string]interface{}  `json:"values"`
            Condition []map[string]interface{}  `json:"condition"`
            Data []map[string]interface{}  `json:"data"`
         
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
    
    // ConfigErrorResponse used by Catalog
    type ConfigErrorResponse struct {

        
            Message string  `json:"message"`
         
    }
    
    // AttributeDetailsGroup used by Catalog
    type AttributeDetailsGroup struct {

        
            DisplayType string  `json:"display_type"`
            IsActive bool  `json:"is_active"`
            Slug string  `json:"slug"`
            Key string  `json:"key"`
            Priority float64  `json:"priority"`
            Logo string  `json:"logo"`
            Unit string  `json:"unit"`
            Name string  `json:"name"`
         
    }
    
    // AppConfigurationDetail used by Catalog
    type AppConfigurationDetail struct {

        
            Attributes []AttributeDetailsGroup  `json:"attributes"`
            TemplateSlugs []string  `json:"template_slugs"`
            IsDefault bool  `json:"is_default"`
            AppID string  `json:"app_id"`
            Slug string  `json:"slug"`
            IsActive bool  `json:"is_active"`
            Priority float64  `json:"priority"`
            Name string  `json:"name"`
            Logo string  `json:"logo"`
         
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
            Key string  `json:"key"`
            IsActive bool  `json:"is_active"`
            Priority float64  `json:"priority"`
            Logo string  `json:"logo"`
            Name string  `json:"name"`
         
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
            Units []map[string]interface{}  `json:"units"`
            Display string  `json:"display"`
         
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
            Compare map[string]interface{}  `json:"compare"`
            Variant map[string]interface{}  `json:"variant"`
         
    }
    
    // GetCatalogConfigurationMetaData used by Catalog
    type GetCatalogConfigurationMetaData struct {

        
            Listing MetaDataListingResponse  `json:"listing"`
            Product GetCatalogConfigurationDetailsProduct  `json:"product"`
         
    }
    
    // ConfigurationListingSortConfig used by Catalog
    type ConfigurationListingSortConfig struct {

        
            IsActive bool  `json:"is_active"`
            Key string  `json:"key"`
            Priority float64  `json:"priority"`
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

        
            Map map[string]interface{}  `json:"map"`
            Condition string  `json:"condition"`
            MapValues []map[string]interface{}  `json:"map_values"`
            BucketPoints []ConfigurationBucketPoints  `json:"bucket_points"`
            Sort string  `json:"sort"`
            Value string  `json:"value"`
         
    }
    
    // ConfigurationListingFilterConfig used by Catalog
    type ConfigurationListingFilterConfig struct {

        
            Type string  `json:"type"`
            DisplayName string  `json:"display_name"`
            IsActive bool  `json:"is_active"`
            Key string  `json:"key"`
            Priority float64  `json:"priority"`
            Name string  `json:"name"`
            ValueConfig ConfigurationListingFilterValue  `json:"value_config"`
            Logo string  `json:"logo"`
         
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

        
            Min float64  `json:"min"`
            Max float64  `json:"max"`
         
    }
    
    // ConfigurationProductConfig used by Catalog
    type ConfigurationProductConfig struct {

        
            Subtitle string  `json:"subtitle"`
            Title string  `json:"title"`
            IsActive bool  `json:"is_active"`
            Key string  `json:"key"`
            Priority float64  `json:"priority"`
            Logo string  `json:"logo"`
            Size ProductSize  `json:"size"`
         
    }
    
    // ConfigurationProductSimilar used by Catalog
    type ConfigurationProductSimilar struct {

        
            Config []ConfigurationProductConfig  `json:"config"`
         
    }
    
    // ConfigurationProductVariantConfig used by Catalog
    type ConfigurationProductVariantConfig struct {

        
            Size ProductSize  `json:"size"`
            DisplayType string  `json:"display_type"`
            IsActive bool  `json:"is_active"`
            Key string  `json:"key"`
            Priority float64  `json:"priority"`
            Name string  `json:"name"`
            Logo string  `json:"logo"`
         
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

        
            Listing ConfigurationListing  `json:"listing"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ConfigID string  `json:"config_id"`
            Product ConfigurationProduct  `json:"product"`
            Type string  `json:"type"`
            ID string  `json:"id"`
            ModifiedOn string  `json:"modified_on"`
            AppID string  `json:"app_id"`
            ConfigType string  `json:"config_type"`
            CreatedOn string  `json:"created_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
         
    }
    
    // GetAppCatalogConfiguration used by Catalog
    type GetAppCatalogConfiguration struct {

        
            IsDefault bool  `json:"is_default"`
            Data AppCatalogConfiguration  `json:"data"`
         
    }
    
    // AppConfiguration used by Catalog
    type AppConfiguration struct {

        
            Listing ConfigurationListing  `json:"listing"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ConfigID string  `json:"config_id"`
            Product ConfigurationProduct  `json:"product"`
            Type string  `json:"type"`
            ModifiedOn string  `json:"modified_on"`
            AppID string  `json:"app_id"`
            ConfigType string  `json:"config_type"`
            CreatedOn string  `json:"created_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
         
    }
    
    // GetCatalogConfigurationDetailsSchemaListing used by Catalog
    type GetCatalogConfigurationDetailsSchemaListing struct {

        
            Sort map[string]interface{}  `json:"sort"`
            Filter map[string]interface{}  `json:"filter"`
         
    }
    
    // EntityConfiguration used by Catalog
    type EntityConfiguration struct {

        
            Listing GetCatalogConfigurationDetailsSchemaListing  `json:"listing"`
            ConfigID string  `json:"config_id"`
            Product GetCatalogConfigurationDetailsProduct  `json:"product"`
            ID string  `json:"id"`
            AppID string  `json:"app_id"`
            ConfigType string  `json:"config_type"`
         
    }
    
    // GetAppCatalogEntityConfiguration used by Catalog
    type GetAppCatalogEntityConfiguration struct {

        
            IsDefault bool  `json:"is_default"`
            Data EntityConfiguration  `json:"data"`
         
    }
    
    // ProductFiltersValue used by Catalog
    type ProductFiltersValue struct {

        
            DisplayFormat string  `json:"display_format"`
            CurrencySymbol string  `json:"currency_symbol"`
            Count float64  `json:"count"`
            SelectedMin float64  `json:"selected_min"`
            IsSelected bool  `json:"is_selected"`
            CurrencyCode string  `json:"currency_code"`
            Display string  `json:"display"`
            Min float64  `json:"min"`
            Value interface{}  `json:"value"`
            QueryFormat string  `json:"query_format"`
            Max float64  `json:"max"`
            SelectedMax float64  `json:"selected_max"`
         
    }
    
    // ProductFiltersKey used by Catalog
    type ProductFiltersKey struct {

        
            Kind string  `json:"kind"`
            Display string  `json:"display"`
            Operators []string  `json:"operators"`
            Logo string  `json:"logo"`
            Name string  `json:"name"`
         
    }
    
    // ProductFilters used by Catalog
    type ProductFilters struct {

        
            Values []ProductFiltersValue  `json:"values"`
            Key ProductFiltersKey  `json:"key"`
         
    }
    
    // ProductSortOn used by Catalog
    type ProductSortOn struct {

        
            IsSelected bool  `json:"is_selected"`
            Display string  `json:"display"`
            Logo string  `json:"logo"`
            Name string  `json:"name"`
            Value string  `json:"value"`
         
    }
    
    // GetCollectionQueryOptionResponse used by Catalog
    type GetCollectionQueryOptionResponse struct {

        
            Filters []ProductFilters  `json:"filters"`
            Operators map[string]string  `json:"operators"`
            SortOn []ProductSortOn  `json:"sort_on"`
         
    }
    
    // CollectionErrorResponse used by Catalog
    type CollectionErrorResponse struct {

        
            Errors map[string]interface{}  `json:"errors"`
            Code float64  `json:"code"`
            Message string  `json:"message"`
         
    }
    
    // CollectionListingFilterType used by Catalog
    type CollectionListingFilterType struct {

        
            IsSelected bool  `json:"is_selected"`
            Display string  `json:"display"`
            Name string  `json:"name"`
         
    }
    
    // CollectionListingFilterTag used by Catalog
    type CollectionListingFilterTag struct {

        
            IsSelected bool  `json:"is_selected"`
            Display string  `json:"display"`
            Name string  `json:"name"`
         
    }
    
    // CollectionListingFilter used by Catalog
    type CollectionListingFilter struct {

        
            Type []CollectionListingFilterType  `json:"type"`
            Tags []CollectionListingFilterTag  `json:"tags"`
         
    }
    
    // Media2 used by Catalog
    type Media2 struct {

        
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
    
    // BannerImage used by Catalog
    type BannerImage struct {

        
            URL string  `json:"url"`
            Type string  `json:"type"`
            AspectRatio string  `json:"aspect_ratio"`
         
    }
    
    // ImageUrls used by Catalog
    type ImageUrls struct {

        
            Landscape BannerImage  `json:"landscape"`
            Portrait BannerImage  `json:"portrait"`
         
    }
    
    // CollectionQuery used by Catalog
    type CollectionQuery struct {

        
            Op string  `json:"op"`
            Value []interface{}  `json:"value"`
            Attribute string  `json:"attribute"`
         
    }
    
    // GetCollectionDetailNest used by Catalog
    type GetCollectionDetailNest struct {

        
            Schedule map[string]interface{}  `json:"_schedule"`
            Logo Media2  `json:"logo"`
            UID string  `json:"uid"`
            SortOn string  `json:"sort_on"`
            AppID string  `json:"app_id"`
            Name string  `json:"name"`
            Action Action  `json:"action"`
            Meta map[string]interface{}  `json:"meta"`
            Description string  `json:"description"`
            Banners ImageUrls  `json:"banners"`
            Priority float64  `json:"priority"`
            AllowSort bool  `json:"allow_sort"`
            Cron map[string]interface{}  `json:"cron"`
            Badge map[string]interface{}  `json:"badge"`
            Query []CollectionQuery  `json:"query"`
            Type string  `json:"type"`
            AllowFacets bool  `json:"allow_facets"`
            IsActive bool  `json:"is_active"`
            Slug string  `json:"slug"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Tag []string  `json:"tag"`
         
    }
    
    // GetCollectionListingResponse used by Catalog
    type GetCollectionListingResponse struct {

        
            Filters CollectionListingFilter  `json:"filters"`
            Items []GetCollectionDetailNest  `json:"items"`
            Page Page  `json:"page"`
         
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
            End string  `json:"end"`
            Start string  `json:"start"`
            Cron string  `json:"cron"`
         
    }
    
    // CollectionImage used by Catalog
    type CollectionImage struct {

        
            AspectRatio string  `json:"aspect_ratio"`
            URL string  `json:"url"`
         
    }
    
    // UserInfo used by Catalog
    type UserInfo struct {

        
            UserID string  `json:"user_id"`
            UID string  `json:"uid"`
            Username string  `json:"username"`
            Email string  `json:"email"`
         
    }
    
    // CollectionBanner used by Catalog
    type CollectionBanner struct {

        
            Landscape CollectionImage  `json:"landscape"`
            Portrait CollectionImage  `json:"portrait"`
         
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
    
    // CreateCollection used by Catalog
    type CreateCollection struct {

        
            Published bool  `json:"published"`
            Schedule CollectionSchedule  `json:"_schedule"`
            Tags []string  `json:"tags"`
            Logo CollectionImage  `json:"logo"`
            CreatedBy UserInfo  `json:"created_by"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            SortOn string  `json:"sort_on"`
            AppID string  `json:"app_id"`
            Name string  `json:"name"`
            IsVisible bool  `json:"is_visible"`
            Meta map[string]interface{}  `json:"meta"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            Description string  `json:"description"`
            Banners CollectionBanner  `json:"banners"`
            Priority float64  `json:"priority"`
            ModifiedBy UserInfo  `json:"modified_by"`
            AllowSort bool  `json:"allow_sort"`
            Seo SeoDetail  `json:"seo"`
            Badge CollectionBadge  `json:"badge"`
            Query []CollectionQuery  `json:"query"`
            Type string  `json:"type"`
            AllowFacets bool  `json:"allow_facets"`
            IsActive bool  `json:"is_active"`
            Slug string  `json:"slug"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
         
    }
    
    // CollectionCreateResponse used by Catalog
    type CollectionCreateResponse struct {

        
            Schedule map[string]interface{}  `json:"_schedule"`
            Tags []string  `json:"tags"`
            Logo Media2  `json:"logo"`
            SortOn string  `json:"sort_on"`
            AppID string  `json:"app_id"`
            Name string  `json:"name"`
            Meta map[string]interface{}  `json:"meta"`
            Description string  `json:"description"`
            Banners ImageUrls  `json:"banners"`
            Priority float64  `json:"priority"`
            AllowSort bool  `json:"allow_sort"`
            Cron map[string]interface{}  `json:"cron"`
            Badge map[string]interface{}  `json:"badge"`
            Query []CollectionQuery  `json:"query"`
            Type string  `json:"type"`
            AllowFacets bool  `json:"allow_facets"`
            IsActive bool  `json:"is_active"`
            Slug string  `json:"slug"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
         
    }
    
    // CollectionDetailResponse used by Catalog
    type CollectionDetailResponse struct {

        
            Schedule map[string]interface{}  `json:"_schedule"`
            Logo Media2  `json:"logo"`
            UID string  `json:"uid"`
            SortOn string  `json:"sort_on"`
            AppID string  `json:"app_id"`
            Name string  `json:"name"`
            Meta map[string]interface{}  `json:"meta"`
            Description string  `json:"description"`
            Banners ImageUrls  `json:"banners"`
            Priority float64  `json:"priority"`
            AllowSort bool  `json:"allow_sort"`
            Cron map[string]interface{}  `json:"cron"`
            Badge map[string]interface{}  `json:"badge"`
            Query []CollectionQuery  `json:"query"`
            Type string  `json:"type"`
            AllowFacets bool  `json:"allow_facets"`
            IsActive bool  `json:"is_active"`
            Slug string  `json:"slug"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Tag []string  `json:"tag"`
         
    }
    
    // UpdateCollection used by Catalog
    type UpdateCollection struct {

        
            Published bool  `json:"published"`
            Schedule CollectionSchedule  `json:"_schedule"`
            Tags []string  `json:"tags"`
            Logo CollectionImage  `json:"logo"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            SortOn string  `json:"sort_on"`
            Name string  `json:"name"`
            IsVisible bool  `json:"is_visible"`
            Meta map[string]interface{}  `json:"meta"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            Description string  `json:"description"`
            Banners CollectionBanner  `json:"banners"`
            Priority float64  `json:"priority"`
            ModifiedBy UserInfo  `json:"modified_by"`
            AllowSort bool  `json:"allow_sort"`
            Seo SeoDetail  `json:"seo"`
            Badge CollectionBadge  `json:"badge"`
            Query []CollectionQuery  `json:"query"`
            Type string  `json:"type"`
            AllowFacets bool  `json:"allow_facets"`
            IsActive bool  `json:"is_active"`
            Slug string  `json:"slug"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
         
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

        
            Published bool  `json:"published"`
            Schedule CollectionSchedule  `json:"_schedule"`
            Tags []string  `json:"tags"`
            Logo CollectionUpdateImage  `json:"logo"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            SortOn string  `json:"sort_on"`
            AppID string  `json:"app_id"`
            Name string  `json:"name"`
            IsVisible bool  `json:"is_visible"`
            Meta map[string]interface{}  `json:"meta"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            Description string  `json:"description"`
            Banners CollectionUpdateBanner  `json:"banners"`
            Priority float64  `json:"priority"`
            ModifiedBy UserInfo  `json:"modified_by"`
            AllowSort bool  `json:"allow_sort"`
            Seo SeoDetail  `json:"seo"`
            Badge CollectionBadge  `json:"badge"`
            Query []CollectionQuery  `json:"query"`
            Type string  `json:"type"`
            AllowFacets bool  `json:"allow_facets"`
            IsActive bool  `json:"is_active"`
            Slug string  `json:"slug"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
         
    }
    
    // ProductDetailAttribute used by Catalog
    type ProductDetailAttribute struct {

        
            Value string  `json:"value"`
            Key string  `json:"key"`
            Type string  `json:"type"`
         
    }
    
    // ProductDetailGroupedAttribute used by Catalog
    type ProductDetailGroupedAttribute struct {

        
            Details []ProductDetailAttribute  `json:"details"`
            Title string  `json:"title"`
         
    }
    
    // Price1 used by Catalog
    type Price1 struct {

        
            Min float64  `json:"min"`
            CurrencyCode string  `json:"currency_code"`
            CurrencySymbol string  `json:"currency_symbol"`
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
            UID float64  `json:"uid"`
            Logo Media2  `json:"logo"`
            Name string  `json:"name"`
         
    }
    
    // ProductListingDetail used by Catalog
    type ProductListingDetail struct {

        
            Similars []string  `json:"similars"`
            ItemType string  `json:"item_type"`
            Attributes map[string]interface{}  `json:"attributes"`
            Rating float64  `json:"rating"`
            ItemCode string  `json:"item_code"`
            Medias []Media2  `json:"medias"`
            Discount string  `json:"discount"`
            UID float64  `json:"uid"`
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
            ShortDescription string  `json:"short_description"`
            PromoMeta map[string]interface{}  `json:"promo_meta"`
            Name string  `json:"name"`
            Price ProductListingPrice  `json:"price"`
            Sellable bool  `json:"sellable"`
            Tryouts []string  `json:"tryouts"`
            ImageNature string  `json:"image_nature"`
            RatingCount float64  `json:"rating_count"`
            Color string  `json:"color"`
            Highlights []string  `json:"highlights"`
            HasVariant bool  `json:"has_variant"`
            Description string  `json:"description"`
            Brand ProductBrand  `json:"brand"`
            Type string  `json:"type"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            Slug string  `json:"slug"`
            ProductOnlineDate string  `json:"product_online_date"`
         
    }
    
    // GetCollectionItemsResponse used by Catalog
    type GetCollectionItemsResponse struct {

        
            Filters []ProductFilters  `json:"filters"`
            Items []ProductListingDetail  `json:"items"`
            Page Page  `json:"page"`
            SortOn []ProductSortOn  `json:"sort_on"`
         
    }
    
    // CollectionItem used by Catalog
    type CollectionItem struct {

        
            Action string  `json:"action"`
            ItemID float64  `json:"item_id"`
            Priority float64  `json:"priority"`
         
    }
    
    // CollectionItemUpdate used by Catalog
    type CollectionItemUpdate struct {

        
            Query []CollectionQuery  `json:"query"`
            Type string  `json:"type"`
            AllowFacets bool  `json:"allow_facets"`
            Items []CollectionItem  `json:"items"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            AllowSort bool  `json:"allow_sort"`
         
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
            AvailableSizes float64  `json:"available_sizes"`
            ArticleFreshness float64  `json:"article_freshness"`
            Name string  `json:"name"`
            AvailableArticles float64  `json:"available_articles"`
         
    }
    
    // CatalogInsightResponse used by Catalog
    type CatalogInsightResponse struct {

        
            Item CatalogInsightItem  `json:"item"`
            BrandDistribution CatalogInsightBrand  `json:"brand_distribution"`
         
    }
    
    // CrossSellingDataResponse used by Catalog
    type CrossSellingDataResponse struct {

        
            Products float64  `json:"products"`
            Articles float64  `json:"articles"`
         
    }
    
    // OptInPostRequest used by Catalog
    type OptInPostRequest struct {

        
            CompanyID float64  `json:"company_id"`
            Platform string  `json:"platform"`
            BrandIds []float64  `json:"brand_ids"`
            StoreIds []float64  `json:"store_ids"`
            OptLevel string  `json:"opt_level"`
            Enabled bool  `json:"enabled"`
         
    }
    
    // CompanyOptIn used by Catalog
    type CompanyOptIn struct {

        
            CreatedBy map[string]interface{}  `json:"created_by"`
            CompanyID float64  `json:"company_id"`
            ModifiedOn float64  `json:"modified_on"`
            Platform string  `json:"platform"`
            CreatedOn float64  `json:"created_on"`
            StoreIds []float64  `json:"store_ids"`
            BrandIds []float64  `json:"brand_ids"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            OptLevel string  `json:"opt_level"`
            Enabled bool  `json:"enabled"`
         
    }
    
    // GetOptInPlatform used by Catalog
    type GetOptInPlatform struct {

        
            Items []CompanyOptIn  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // OptinCompanyDetail used by Catalog
    type OptinCompanyDetail struct {

        
            UID float64  `json:"uid"`
            CompanyType string  `json:"company_type"`
            BusinessType string  `json:"business_type"`
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

        
            Store float64  `json:"store"`
            Brand float64  `json:"brand"`
            Company string  `json:"company"`
         
    }
    
    // StoreDetail used by Catalog
    type StoreDetail struct {

        
            Timing map[string]interface{}  `json:"timing"`
            CompanyID float64  `json:"company_id"`
            UID float64  `json:"uid"`
            Address map[string]interface{}  `json:"address"`
            Manager map[string]interface{}  `json:"manager"`
            StoreType string  `json:"store_type"`
            StoreCode string  `json:"store_code"`
            DisplayName string  `json:"display_name"`
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
            Name string  `json:"name"`
            AdditionalContacts []map[string]interface{}  `json:"additional_contacts"`
            Documents []map[string]interface{}  `json:"documents"`
         
    }
    
    // OptinStoreDetails used by Catalog
    type OptinStoreDetails struct {

        
            Items []StoreDetail  `json:"items"`
            Page Page  `json:"page"`
         
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

        
            DependsOn []string  `json:"depends_on"`
            Indexing bool  `json:"indexing"`
            Priority float64  `json:"priority"`
         
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
            Mandatory bool  `json:"mandatory"`
            Range AttributeSchemaRange  `json:"range"`
            Multi bool  `json:"multi"`
            AllowedValues []string  `json:"allowed_values"`
         
    }
    
    // GenderDetail used by Catalog
    type GenderDetail struct {

        
            Departments []string  `json:"departments"`
            Meta AttributeMasterMeta  `json:"meta"`
            EnabledForEndConsumer bool  `json:"enabled_for_end_consumer"`
            Description string  `json:"description"`
            ID string  `json:"id"`
            Filters AttributeMasterFilter  `json:"filters"`
            Slug string  `json:"slug"`
            IsNested bool  `json:"is_nested"`
            Name string  `json:"name"`
            Details AttributeMasterDetails  `json:"details"`
            Logo string  `json:"logo"`
            Schema AttributeMaster  `json:"schema"`
         
    }
    
    // CategoriesResponse used by Catalog
    type CategoriesResponse struct {

        
            UID float64  `json:"uid"`
            TemplateSlug string  `json:"template_slug"`
            Slug string  `json:"slug"`
            SlugKey string  `json:"slug_key"`
            Name string  `json:"name"`
         
    }
    
    // ProdcutTemplateCategoriesResponse used by Catalog
    type ProdcutTemplateCategoriesResponse struct {

        
            Items []CategoriesResponse  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // PTErrorResponse used by Catalog
    type PTErrorResponse struct {

        
            Meta map[string]interface{}  `json:"meta"`
            Message string  `json:"message"`
            Errors map[string]interface{}  `json:"errors"`
            Status float64  `json:"status"`
            Code string  `json:"code"`
         
    }
    
    // UserSerializer used by Catalog
    type UserSerializer struct {

        
            UID string  `json:"uid"`
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
            ID string  `json:"_id"`
            Contact string  `json:"contact"`
         
    }
    
    // GetDepartment used by Catalog
    type GetDepartment struct {

        
            Synonyms []string  `json:"synonyms"`
            CreatedBy UserSerializer  `json:"created_by"`
            UID float64  `json:"uid"`
            PageSize float64  `json:"page_size"`
            Search string  `json:"search"`
            ItemType string  `json:"item_type"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            IsActive bool  `json:"is_active"`
            ModifiedOn string  `json:"modified_on"`
            Slug string  `json:"slug"`
            CreatedOn string  `json:"created_on"`
            PriorityOrder float64  `json:"priority_order"`
            Logo string  `json:"logo"`
            Name string  `json:"name"`
            PageNo float64  `json:"page_no"`
         
    }
    
    // DepartmentsResponse used by Catalog
    type DepartmentsResponse struct {

        
            Items []GetDepartment  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // DepartmentErrorResponse used by Catalog
    type DepartmentErrorResponse struct {

        
            Meta map[string]interface{}  `json:"meta"`
            Message string  `json:"message"`
            Errors map[string]interface{}  `json:"errors"`
            Status float64  `json:"status"`
            Code string  `json:"code"`
         
    }
    
    // DepartmentCreateUpdate used by Catalog
    type DepartmentCreateUpdate struct {

        
            Synonyms []string  `json:"synonyms"`
            Cls string  `json:"_cls"`
            UID float64  `json:"uid"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Platforms map[string]interface{}  `json:"platforms"`
            IsActive bool  `json:"is_active"`
            Slug string  `json:"slug"`
            Tags []string  `json:"tags"`
            PriorityOrder float64  `json:"priority_order"`
            Logo string  `json:"logo"`
            Name string  `json:"name"`
         
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

        
            CreatedBy UserDetail  `json:"created_by"`
            VerifiedOn string  `json:"verified_on"`
            UID float64  `json:"uid"`
            Synonyms []string  `json:"synonyms"`
            Cls string  `json:"_cls"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            VerifiedBy UserDetail  `json:"verified_by"`
            ModifiedOn string  `json:"modified_on"`
            IsActive bool  `json:"is_active"`
            CreatedOn string  `json:"created_on"`
            Slug string  `json:"slug"`
            PriorityOrder float64  `json:"priority_order"`
            ModifiedBy UserDetail  `json:"modified_by"`
            Name string  `json:"name"`
            ID interface{}  `json:"_id"`
            Logo string  `json:"logo"`
         
    }
    
    // ProductTemplate used by Catalog
    type ProductTemplate struct {

        
            CreatedBy map[string]interface{}  `json:"created_by"`
            Categories []string  `json:"categories"`
            Departments []string  `json:"departments"`
            Attributes []string  `json:"attributes"`
            IsArchived bool  `json:"is_archived"`
            Description string  `json:"description"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            IsActive bool  `json:"is_active"`
            ModifiedOn string  `json:"modified_on"`
            Slug string  `json:"slug"`
            IsPhysical bool  `json:"is_physical"`
            Tag string  `json:"tag"`
            CreatedOn string  `json:"created_on"`
            Name string  `json:"name"`
            Logo string  `json:"logo"`
            IsExpirable bool  `json:"is_expirable"`
         
    }
    
    // TemplatesResponse used by Catalog
    type TemplatesResponse struct {

        
            Items ProductTemplate  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // TemplateDetails used by Catalog
    type TemplateDetails struct {

        
            Categories []string  `json:"categories"`
            Departments []string  `json:"departments"`
            Attributes []string  `json:"attributes"`
            IsArchived bool  `json:"is_archived"`
            Description string  `json:"description"`
            ID string  `json:"id"`
            IsActive bool  `json:"is_active"`
            Slug string  `json:"slug"`
            IsPhysical bool  `json:"is_physical"`
            Tag string  `json:"tag"`
            Name string  `json:"name"`
            Logo string  `json:"logo"`
            IsExpirable bool  `json:"is_expirable"`
         
    }
    
    // Properties used by Catalog
    type Properties struct {

        
            ItemType map[string]interface{}  `json:"item_type"`
            ItemCode map[string]interface{}  `json:"item_code"`
            Tags map[string]interface{}  `json:"tags"`
            Variants map[string]interface{}  `json:"variants"`
            BrandUID map[string]interface{}  `json:"brand_uid"`
            Currency map[string]interface{}  `json:"currency"`
            ShortDescription map[string]interface{}  `json:"short_description"`
            TraderType map[string]interface{}  `json:"trader_type"`
            HsnCode map[string]interface{}  `json:"hsn_code"`
            CategorySlug map[string]interface{}  `json:"category_slug"`
            Name map[string]interface{}  `json:"name"`
            NoOfBoxes map[string]interface{}  `json:"no_of_boxes"`
            Highlights map[string]interface{}  `json:"highlights"`
            Description map[string]interface{}  `json:"description"`
            ProductGroupTag map[string]interface{}  `json:"product_group_tag"`
            ReturnConfig map[string]interface{}  `json:"return_config"`
            SizeGuide map[string]interface{}  `json:"size_guide"`
            Command map[string]interface{}  `json:"command"`
            ProductPublish map[string]interface{}  `json:"product_publish"`
            Trader map[string]interface{}  `json:"trader"`
            Media map[string]interface{}  `json:"media"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            IsActive map[string]interface{}  `json:"is_active"`
            Slug map[string]interface{}  `json:"slug"`
            MultiSize map[string]interface{}  `json:"multi_size"`
            Sizes map[string]interface{}  `json:"sizes"`
            CountryOfOrigin map[string]interface{}  `json:"country_of_origin"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            IsDependent map[string]interface{}  `json:"is_dependent"`
         
    }
    
    // GlobalValidation used by Catalog
    type GlobalValidation struct {

        
            Type string  `json:"type"`
            Description string  `json:"description"`
            Title string  `json:"title"`
            Required []string  `json:"required"`
            Definitions map[string]interface{}  `json:"definitions"`
            Properties Properties  `json:"properties"`
         
    }
    
    // TemplateValidationData used by Catalog
    type TemplateValidationData struct {

        
            TemplateValidation map[string]interface{}  `json:"template_validation"`
            GlobalValidation GlobalValidation  `json:"global_validation"`
         
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
            UID string  `json:"uid"`
            Username string  `json:"username"`
            Email string  `json:"email"`
         
    }
    
    // ProductTemplateExportResponse used by Catalog
    type ProductTemplateExportResponse struct {

        
            CompletedOn string  `json:"completed_on"`
            CreatedBy UserInfo1  `json:"created_by"`
            SellerID float64  `json:"seller_id"`
            Type string  `json:"type"`
            TaskID string  `json:"task_id"`
            NotificationEmails []string  `json:"notification_emails"`
            Status string  `json:"status"`
            URL string  `json:"url"`
            ModifiedOn string  `json:"modified_on"`
            Filters map[string]interface{}  `json:"filters"`
         
    }
    
    // ProductDownloadsResponse used by Catalog
    type ProductDownloadsResponse struct {

        
            Items []ProductTemplateExportResponse  `json:"items"`
         
    }
    
    // ProductTemplateExportFilterRequest used by Catalog
    type ProductTemplateExportFilterRequest struct {

        
            FromDate string  `json:"from_date"`
            Brands []string  `json:"brands"`
            Templates []string  `json:"templates"`
            CatalogueTypes []string  `json:"catalogue_types"`
            ToDate string  `json:"to_date"`
         
    }
    
    // ProductTemplateDownloadsExport used by Catalog
    type ProductTemplateDownloadsExport struct {

        
            NotificationEmails []string  `json:"notification_emails"`
            Filters ProductTemplateExportFilterRequest  `json:"filters"`
            Type string  `json:"type"`
         
    }
    
    // ProductConfigurationDownloads used by Catalog
    type ProductConfigurationDownloads struct {

        
            Multivalue bool  `json:"multivalue"`
            Data []map[string]interface{}  `json:"data"`
         
    }
    
    // CategoryMappingValues used by Catalog
    type CategoryMappingValues struct {

        
            CatalogID float64  `json:"catalog_id"`
            Name string  `json:"name"`
         
    }
    
    // CategoryMapping used by Catalog
    type CategoryMapping struct {

        
            Facebook CategoryMappingValues  `json:"facebook"`
            Ajio CategoryMappingValues  `json:"ajio"`
            Google CategoryMappingValues  `json:"google"`
         
    }
    
    // Media3 used by Catalog
    type Media3 struct {

        
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

        
            Tryouts []string  `json:"tryouts"`
            Synonyms []string  `json:"synonyms"`
            Departments []float64  `json:"departments"`
            UID float64  `json:"uid"`
            Marketplaces CategoryMapping  `json:"marketplaces"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Media Media3  `json:"media"`
            ID string  `json:"id"`
            Hierarchy []Hierarchy  `json:"hierarchy"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            IsActive bool  `json:"is_active"`
            ModifiedOn string  `json:"modified_on"`
            Slug string  `json:"slug"`
            CreatedOn string  `json:"created_on"`
            Priority float64  `json:"priority"`
            Name string  `json:"name"`
            Level float64  `json:"level"`
         
    }
    
    // CategoryResponse used by Catalog
    type CategoryResponse struct {

        
            Items []Category  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // CategoryRequestBody used by Catalog
    type CategoryRequestBody struct {

        
            Tryouts []string  `json:"tryouts"`
            Synonyms []string  `json:"synonyms"`
            Departments []float64  `json:"departments"`
            Marketplaces CategoryMapping  `json:"marketplaces"`
            Media Media3  `json:"media"`
            Hierarchy []Hierarchy  `json:"hierarchy"`
            IsActive bool  `json:"is_active"`
            Slug string  `json:"slug"`
            Priority float64  `json:"priority"`
            Name string  `json:"name"`
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
    
    // NetQuantityResponse used by Catalog
    type NetQuantityResponse struct {

        
            Unit string  `json:"unit"`
            Value float64  `json:"value"`
         
    }
    
    // VerifiedBy used by Catalog
    type VerifiedBy struct {

        
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
         
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
    
    // Trader used by Catalog
    type Trader struct {

        
            Type string  `json:"type"`
            Address []string  `json:"address"`
            Name string  `json:"name"`
         
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

        
            UID float64  `json:"uid"`
            Logo Logo  `json:"logo"`
            Name string  `json:"name"`
         
    }
    
    // Image used by Catalog
    type Image struct {

        
            URL string  `json:"url"`
            SecureURL string  `json:"secure_url"`
            AspectRatioF float64  `json:"aspect_ratio_f"`
            AspectRatio string  `json:"aspect_ratio"`
         
    }
    
    // ProductSchemaV2 used by Catalog
    type ProductSchemaV2 struct {

        
            ItemType string  `json:"item_type"`
            VerifiedOn string  `json:"verified_on"`
            Attributes map[string]interface{}  `json:"attributes"`
            ModifiedOn string  `json:"modified_on"`
            ItemCode string  `json:"item_code"`
            CreatedOn string  `json:"created_on"`
            L3Mapping []string  `json:"l3_mapping"`
            Variants map[string]interface{}  `json:"variants"`
            Tags []string  `json:"tags"`
            NetQuantity NetQuantityResponse  `json:"net_quantity"`
            PrimaryColor string  `json:"primary_color"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Departments []float64  `json:"departments"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            UID float64  `json:"uid"`
            AllIdentifiers []string  `json:"all_identifiers"`
            AllCompanyIds []float64  `json:"all_company_ids"`
            Currency string  `json:"currency"`
            Stage string  `json:"stage"`
            ShortDescription string  `json:"short_description"`
            VerifiedBy VerifiedBy  `json:"verified_by"`
            HsnCode string  `json:"hsn_code"`
            CategorySlug string  `json:"category_slug"`
            VariantGroup map[string]interface{}  `json:"variant_group"`
            Name string  `json:"name"`
            NoOfBoxes float64  `json:"no_of_boxes"`
            IsImageLessProduct bool  `json:"is_image_less_product"`
            ImageNature string  `json:"image_nature"`
            Highlights []string  `json:"highlights"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            Color string  `json:"color"`
            VariantMedia map[string]interface{}  `json:"variant_media"`
            Description string  `json:"description"`
            ProductGroupTag []string  `json:"product_group_tag"`
            IsPhysical bool  `json:"is_physical"`
            ReturnConfig ReturnConfigResponse  `json:"return_config"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            SizeGuide string  `json:"size_guide"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            ProductPublish ProductPublish  `json:"product_publish"`
            Pending string  `json:"pending"`
            Category map[string]interface{}  `json:"category"`
            IsExpirable bool  `json:"is_expirable"`
            CategoryUID float64  `json:"category_uid"`
            Trader []Trader  `json:"trader"`
            Brand Brand  `json:"brand"`
            Moq map[string]interface{}  `json:"moq"`
            CompanyID float64  `json:"company_id"`
            Media []Media2  `json:"media"`
            ID string  `json:"id"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            TemplateTag string  `json:"template_tag"`
            AllSizes []map[string]interface{}  `json:"all_sizes"`
            IsActive bool  `json:"is_active"`
            Slug string  `json:"slug"`
            MultiSize bool  `json:"multi_size"`
            Sizes []map[string]interface{}  `json:"sizes"`
            IsSet bool  `json:"is_set"`
            CountryOfOrigin string  `json:"country_of_origin"`
            Images []Image  `json:"images"`
            BrandUID float64  `json:"brand_uid"`
            IsDependent bool  `json:"is_dependent"`
         
    }
    
    // ProductListingResponseV2 used by Catalog
    type ProductListingResponseV2 struct {

        
            Items []ProductSchemaV2  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // NetQuantity used by Catalog
    type NetQuantity struct {

        
            Unit string  `json:"unit"`
            Value float64  `json:"value"`
         
    }
    
    // TaxIdentifier used by Catalog
    type TaxIdentifier struct {

        
            HsnCodeID string  `json:"hsn_code_id"`
            HsnCode string  `json:"hsn_code"`
            ReportingHsn string  `json:"reporting_hsn"`
         
    }
    
    // ReturnConfig used by Catalog
    type ReturnConfig struct {

        
            Time float64  `json:"time"`
            Unit string  `json:"unit"`
            Returnable bool  `json:"returnable"`
         
    }
    
    // CustomOrder used by Catalog
    type CustomOrder struct {

        
            ManufacturingTime float64  `json:"manufacturing_time"`
            IsCustomOrder bool  `json:"is_custom_order"`
            ManufacturingTimeUnit string  `json:"manufacturing_time_unit"`
         
    }
    
    // ProductPublish1 used by Catalog
    type ProductPublish1 struct {

        
            ProductOnlineDate string  `json:"product_online_date"`
            IsSet bool  `json:"is_set"`
         
    }
    
    // TeaserTag used by Catalog
    type TeaserTag struct {

        
            Tag string  `json:"tag"`
            URL string  `json:"url"`
         
    }
    
    // ProductCreateUpdateSchemaV2 used by Catalog
    type ProductCreateUpdateSchemaV2 struct {

        
            ItemType string  `json:"item_type"`
            Attributes map[string]interface{}  `json:"attributes"`
            ItemCode string  `json:"item_code"`
            Tags []string  `json:"tags"`
            Variants map[string]interface{}  `json:"variants"`
            NetQuantity NetQuantity  `json:"net_quantity"`
            Requester string  `json:"requester"`
            Departments []float64  `json:"departments"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            UID float64  `json:"uid"`
            Currency string  `json:"currency"`
            BulkJobID string  `json:"bulk_job_id"`
            ShortDescription string  `json:"short_description"`
            VariantGroup map[string]interface{}  `json:"variant_group"`
            CategorySlug string  `json:"category_slug"`
            Name string  `json:"name"`
            Action string  `json:"action"`
            NoOfBoxes float64  `json:"no_of_boxes"`
            IsImageLessProduct bool  `json:"is_image_less_product"`
            Highlights []string  `json:"highlights"`
            TaxIdentifier TaxIdentifier  `json:"tax_identifier"`
            VariantMedia map[string]interface{}  `json:"variant_media"`
            Description string  `json:"description"`
            ProductGroupTag []string  `json:"product_group_tag"`
            ReturnConfig ReturnConfig  `json:"return_config"`
            CustomOrder CustomOrder  `json:"custom_order"`
            SizeGuide string  `json:"size_guide"`
            ProductPublish ProductPublish1  `json:"product_publish"`
            CompanyID float64  `json:"company_id"`
            Trader []Trader  `json:"trader"`
            Media []Media2  `json:"media"`
            ChangeRequestID interface{}  `json:"change_request_id"`
            TeaserTag TeaserTag  `json:"teaser_tag"`
            TemplateTag string  `json:"template_tag"`
            IsActive bool  `json:"is_active"`
            Slug string  `json:"slug"`
            MultiSize bool  `json:"multi_size"`
            Sizes []map[string]interface{}  `json:"sizes"`
            IsSet bool  `json:"is_set"`
            CountryOfOrigin string  `json:"country_of_origin"`
            BrandUID float64  `json:"brand_uid"`
            IsDependent bool  `json:"is_dependent"`
         
    }
    
    // ProductVariants used by Catalog
    type ProductVariants struct {

        
            UID float64  `json:"uid"`
            Media []Media2  `json:"media"`
            ItemCode string  `json:"item_code"`
            Name string  `json:"name"`
            BrandUID float64  `json:"brand_uid"`
            CategoryUID float64  `json:"category_uid"`
         
    }
    
    // ProductVariantsResponse used by Catalog
    type ProductVariantsResponse struct {

        
            Variants []ProductVariants  `json:"variants"`
            Page Page  `json:"page"`
         
    }
    
    // AttributeMasterSerializer used by Catalog
    type AttributeMasterSerializer struct {

        
            Synonyms map[string]interface{}  `json:"synonyms"`
            Variant bool  `json:"variant"`
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
            Tags []string  `json:"tags"`
            Logo string  `json:"logo"`
            RawKey string  `json:"raw_key"`
            Schema AttributeMaster  `json:"schema"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Departments []string  `json:"departments"`
            Name string  `json:"name"`
            Description string  `json:"description"`
            IsNested bool  `json:"is_nested"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Unit string  `json:"unit"`
            EnabledForEndConsumer bool  `json:"enabled_for_end_consumer"`
            Filters AttributeMasterFilter  `json:"filters"`
            Slug string  `json:"slug"`
            Details AttributeMasterDetails  `json:"details"`
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

        
            ItemWeight float64  `json:"item_weight"`
            ItemWidth float64  `json:"item_width"`
            ItemLength float64  `json:"item_length"`
            Identifiers []ValidateIdentifier  `json:"identifiers"`
            ItemWeightUnitOfMeasure string  `json:"item_weight_unit_of_measure"`
            ItemDimensionsUnitOfMeasure string  `json:"item_dimensions_unit_of_measure"`
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

        
            ItemType string  `json:"item_type"`
            VerifiedOn string  `json:"verified_on"`
            Attributes map[string]interface{}  `json:"attributes"`
            ModifiedOn string  `json:"modified_on"`
            ItemCode string  `json:"item_code"`
            CreatedOn string  `json:"created_on"`
            L3Mapping []string  `json:"l3_mapping"`
            Variants map[string]interface{}  `json:"variants"`
            Tags []string  `json:"tags"`
            NetQuantity NetQuantityResponse  `json:"net_quantity"`
            PrimaryColor string  `json:"primary_color"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Departments []float64  `json:"departments"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            UID float64  `json:"uid"`
            AllIdentifiers []string  `json:"all_identifiers"`
            AllCompanyIds []float64  `json:"all_company_ids"`
            Currency string  `json:"currency"`
            Stage string  `json:"stage"`
            ShortDescription string  `json:"short_description"`
            VerifiedBy VerifiedBy  `json:"verified_by"`
            HsnCode string  `json:"hsn_code"`
            CategorySlug string  `json:"category_slug"`
            VariantGroup map[string]interface{}  `json:"variant_group"`
            Name string  `json:"name"`
            NoOfBoxes float64  `json:"no_of_boxes"`
            IsImageLessProduct bool  `json:"is_image_less_product"`
            ImageNature string  `json:"image_nature"`
            Highlights []string  `json:"highlights"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            Color string  `json:"color"`
            VariantMedia map[string]interface{}  `json:"variant_media"`
            Description string  `json:"description"`
            ProductGroupTag []string  `json:"product_group_tag"`
            IsPhysical bool  `json:"is_physical"`
            ReturnConfig ReturnConfigResponse  `json:"return_config"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            SizeGuide string  `json:"size_guide"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            ProductPublish ProductPublished  `json:"product_publish"`
            Pending string  `json:"pending"`
            Category map[string]interface{}  `json:"category"`
            IsExpirable bool  `json:"is_expirable"`
            CategoryUID float64  `json:"category_uid"`
            Trader []Trader  `json:"trader"`
            Brand Brand  `json:"brand"`
            Moq map[string]interface{}  `json:"moq"`
            CompanyID float64  `json:"company_id"`
            Media []Media2  `json:"media"`
            ID string  `json:"id"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            TemplateTag string  `json:"template_tag"`
            AllSizes []map[string]interface{}  `json:"all_sizes"`
            IsActive bool  `json:"is_active"`
            Slug string  `json:"slug"`
            MultiSize bool  `json:"multi_size"`
            Sizes []map[string]interface{}  `json:"sizes"`
            IsSet bool  `json:"is_set"`
            CountryOfOrigin string  `json:"country_of_origin"`
            Images []Image  `json:"images"`
            BrandUID float64  `json:"brand_uid"`
            IsDependent bool  `json:"is_dependent"`
         
    }
    
    // ProductListingResponse used by Catalog
    type ProductListingResponse struct {

        
            Items []Product  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // UserDetail1 used by Catalog
    type UserDetail1 struct {

        
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
            FullName string  `json:"full_name"`
         
    }
    
    // ProductBulkRequest used by Catalog
    type ProductBulkRequest struct {

        
            Template ProductTemplate  `json:"template"`
            CreatedBy UserDetail1  `json:"created_by"`
            CompanyID float64  `json:"company_id"`
            FilePath string  `json:"file_path"`
            Stage string  `json:"stage"`
            ModifiedBy UserDetail1  `json:"modified_by"`
            TemplateTag string  `json:"template_tag"`
            Failed float64  `json:"failed"`
            IsActive bool  `json:"is_active"`
            CancelledRecords []string  `json:"cancelled_records"`
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
            Cancelled float64  `json:"cancelled"`
            FailedRecords []string  `json:"failed_records"`
            Succeed float64  `json:"succeed"`
            Total float64  `json:"total"`
         
    }
    
    // ProductBulkRequestList used by Catalog
    type ProductBulkRequestList struct {

        
            Items ProductBulkRequest  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // BulkJob used by Catalog
    type BulkJob struct {

        
            CreatedBy UserInfo1  `json:"created_by"`
            CompanyID float64  `json:"company_id"`
            FilePath string  `json:"file_path"`
            Stage string  `json:"stage"`
            FailedRecords []map[string]interface{}  `json:"failed_records"`
            Cancelled float64  `json:"cancelled"`
            Failed float64  `json:"failed"`
            IsActive bool  `json:"is_active"`
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
            CancelledRecords []map[string]interface{}  `json:"cancelled_records"`
            TemplateTag string  `json:"template_tag"`
            TrackingURL string  `json:"tracking_url"`
            ModifiedBy UserInfo1  `json:"modified_by"`
            Succeed float64  `json:"succeed"`
            Total float64  `json:"total"`
            CustomTemplateTag string  `json:"custom_template_tag"`
         
    }
    
    // BulkResponse used by Catalog
    type BulkResponse struct {

        
            CreatedBy UserInfo1  `json:"created_by"`
            IsActive bool  `json:"is_active"`
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
            BatchID string  `json:"batch_id"`
            ModifiedBy UserInfo1  `json:"modified_by"`
         
    }
    
    // BulkProductRequest used by Catalog
    type BulkProductRequest struct {

        
            Data []map[string]interface{}  `json:"data"`
            CompanyID float64  `json:"company_id"`
            BatchID string  `json:"batch_id"`
            TemplateTag string  `json:"template_tag"`
         
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

        
            CreatedBy UserCommon  `json:"created_by"`
            CompanyID float64  `json:"company_id"`
            FilePath string  `json:"file_path"`
            ID string  `json:"id"`
            Stage string  `json:"stage"`
            FailedRecords []string  `json:"failed_records"`
            Retry float64  `json:"retry"`
            Failed float64  `json:"failed"`
            ModifiedOn string  `json:"modified_on"`
            CancelledRecords []string  `json:"cancelled_records"`
            CreatedOn string  `json:"created_on"`
            IsActive bool  `json:"is_active"`
            Cancelled float64  `json:"cancelled"`
            TrackingURL string  `json:"tracking_url"`
            ModifiedBy UserCommon  `json:"modified_by"`
            Succeed float64  `json:"succeed"`
            Total float64  `json:"total"`
         
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
            URL string  `json:"url"`
         
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

        
            InventoryUpdatedOn string  `json:"inventory_updated_on"`
            PriceTransfer float64  `json:"price_transfer"`
            UID string  `json:"uid"`
            Identifiers map[string]interface{}  `json:"identifiers"`
            Currency string  `json:"currency"`
            SellerIdentifier string  `json:"seller_identifier"`
            Store map[string]interface{}  `json:"store"`
            Quantity float64  `json:"quantity"`
            SellableQuantity float64  `json:"sellable_quantity"`
            Price float64  `json:"price"`
            PriceEffective float64  `json:"price_effective"`
            ItemID float64  `json:"item_id"`
            Size string  `json:"size"`
         
    }
    
    // InventoryResponsePaginated used by Catalog
    type InventoryResponsePaginated struct {

        
            Items []InventoryResponse  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // ItemQuery used by Catalog
    type ItemQuery struct {

        
            ItemCode string  `json:"item_code"`
            UID float64  `json:"uid"`
            BrandUID float64  `json:"brand_uid"`
         
    }
    
    // GTIN used by Catalog
    type GTIN struct {

        
            Primary bool  `json:"primary"`
            GtinType string  `json:"gtin_type"`
            GtinValue string  `json:"gtin_value"`
         
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

        
            ItemWeight float64  `json:"item_weight"`
            PriceTransfer float64  `json:"price_transfer"`
            ItemWidth float64  `json:"item_width"`
            ItemLength float64  `json:"item_length"`
            Identifiers []GTIN  `json:"identifiers"`
            Currency string  `json:"currency"`
            ItemWeightUnitOfMeasure string  `json:"item_weight_unit_of_measure"`
            Quantity float64  `json:"quantity"`
            StoreCode string  `json:"store_code"`
            ItemDimensionsUnitOfMeasure string  `json:"item_dimensions_unit_of_measure"`
            Set InventorySet  `json:"set"`
            IsSet bool  `json:"is_set"`
            Price float64  `json:"price"`
            ItemHeight float64  `json:"item_height"`
            PriceEffective float64  `json:"price_effective"`
            Size string  `json:"size"`
            ExpirationDate string  `json:"expiration_date"`
         
    }
    
    // InventoryRequest used by Catalog
    type InventoryRequest struct {

        
            Item ItemQuery  `json:"item"`
            CompanyID float64  `json:"company_id"`
            Sizes []InvSize  `json:"sizes"`
         
    }
    
    // WeightResponse used by Catalog
    type WeightResponse struct {

        
            Unit string  `json:"unit"`
            IsDefault bool  `json:"is_default"`
            Shipping float64  `json:"shipping"`
         
    }
    
    // CompanyMeta used by Catalog
    type CompanyMeta struct {

        
            ID float64  `json:"id"`
         
    }
    
    // PriceMeta used by Catalog
    type PriceMeta struct {

        
            Transfer float64  `json:"transfer"`
            UpdatedAt string  `json:"updated_at"`
            Currency string  `json:"currency"`
            TpNotes map[string]interface{}  `json:"tp_notes"`
            Effective float64  `json:"effective"`
            Marked float64  `json:"marked"`
         
    }
    
    // QuantityBase used by Catalog
    type QuantityBase struct {

        
            Count float64  `json:"count"`
            UpdatedAt string  `json:"updated_at"`
         
    }
    
    // Quantities used by Catalog
    type Quantities struct {

        
            NotAvailable QuantityBase  `json:"not_available"`
            OrderCommitted QuantityBase  `json:"order_committed"`
            Sellable QuantityBase  `json:"sellable"`
            Damaged QuantityBase  `json:"damaged"`
         
    }
    
    // StoreMeta used by Catalog
    type StoreMeta struct {

        
            ID float64  `json:"id"`
         
    }
    
    // ReturnConfig1 used by Catalog
    type ReturnConfig1 struct {

        
            Time float64  `json:"time"`
            Unit string  `json:"unit"`
            Returnable bool  `json:"returnable"`
         
    }
    
    // ManufacturerResponse used by Catalog
    type ManufacturerResponse struct {

        
            IsDefault bool  `json:"is_default"`
            Address string  `json:"address"`
            Name string  `json:"name"`
         
    }
    
    // Trader1 used by Catalog
    type Trader1 struct {

        
            Type string  `json:"type"`
            Address []string  `json:"address"`
            Name string  `json:"name"`
         
    }
    
    // BrandMeta used by Catalog
    type BrandMeta struct {

        
            ID float64  `json:"id"`
            Name string  `json:"name"`
         
    }
    
    // DimensionResponse used by Catalog
    type DimensionResponse struct {

        
            Width float64  `json:"width"`
            IsDefault bool  `json:"is_default"`
            Height float64  `json:"height"`
            Unit string  `json:"unit"`
            Length float64  `json:"length"`
         
    }
    
    // InventorySellerResponse used by Catalog
    type InventorySellerResponse struct {

        
            RawMeta map[string]interface{}  `json:"raw_meta"`
            FyndMeta map[string]interface{}  `json:"fynd_meta"`
            Weight WeightResponse  `json:"weight"`
            Company CompanyMeta  `json:"company"`
            Tags []string  `json:"tags"`
            Identifier map[string]interface{}  `json:"identifier"`
            TrackInventory bool  `json:"track_inventory"`
            CreatedBy UserSerializer  `json:"created_by"`
            UID string  `json:"uid"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Stage string  `json:"stage"`
            Set InventorySet  `json:"set"`
            Price PriceMeta  `json:"price"`
            Quantities Quantities  `json:"quantities"`
            FyndArticleCode string  `json:"fynd_article_code"`
            Size string  `json:"size"`
            Meta map[string]interface{}  `json:"meta"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            Store StoreMeta  `json:"store"`
            TraceID string  `json:"trace_id"`
            TotalQuantity float64  `json:"total_quantity"`
            AddedOnStore string  `json:"added_on_store"`
            ReturnConfig ReturnConfig1  `json:"return_config"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            ItemID float64  `json:"item_id"`
            Fragile bool  `json:"fragile"`
            ExpirationDate string  `json:"expiration_date"`
            Manufacturer ManufacturerResponse  `json:"manufacturer"`
            Trader []Trader1  `json:"trader"`
            Brand BrandMeta  `json:"brand"`
            SellerIdentifier string  `json:"seller_identifier"`
            IsActive bool  `json:"is_active"`
            IsSet bool  `json:"is_set"`
            CountryOfOrigin string  `json:"country_of_origin"`
            FyndItemCode string  `json:"fynd_item_code"`
            Dimension DimensionResponse  `json:"dimension"`
         
    }
    
    // InventorySellerIdentifierResponsePaginated used by Catalog
    type InventorySellerIdentifierResponsePaginated struct {

        
            Items []InventorySellerResponse  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // WeightResponse1 used by Catalog
    type WeightResponse1 struct {

        
            Unit string  `json:"unit"`
            Shipping float64  `json:"shipping"`
         
    }
    
    // CompanyMeta1 used by Catalog
    type CompanyMeta1 struct {

        
            ID float64  `json:"id"`
         
    }
    
    // PriceArticle used by Catalog
    type PriceArticle struct {

        
            Transfer float64  `json:"transfer"`
            Currency string  `json:"currency"`
            TpNotes map[string]interface{}  `json:"tp_notes"`
            Effective float64  `json:"effective"`
            Marked float64  `json:"marked"`
         
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
    
    // DateMeta used by Catalog
    type DateMeta struct {

        
            ModifiedOn string  `json:"modified_on"`
            InventoryUpdatedOn string  `json:"inventory_updated_on"`
            CreatedOn string  `json:"created_on"`
            AddedOnStore string  `json:"added_on_store"`
         
    }
    
    // ArticleStoreResponse used by Catalog
    type ArticleStoreResponse struct {

        
            StoreType string  `json:"store_type"`
            UID float64  `json:"uid"`
            StoreCode string  `json:"store_code"`
            Name string  `json:"name"`
         
    }
    
    // ReturnConfig2 used by Catalog
    type ReturnConfig2 struct {

        
            Time float64  `json:"time"`
            Unit string  `json:"unit"`
            Returnable bool  `json:"returnable"`
         
    }
    
    // ManufacturerResponse1 used by Catalog
    type ManufacturerResponse1 struct {

        
            IsDefault bool  `json:"is_default"`
            Address string  `json:"address"`
            Name string  `json:"name"`
         
    }
    
    // Trader2 used by Catalog
    type Trader2 struct {

        
            Type string  `json:"type"`
            Address []string  `json:"address"`
            Name string  `json:"name"`
         
    }
    
    // BrandMeta1 used by Catalog
    type BrandMeta1 struct {

        
            ID float64  `json:"id"`
            Name string  `json:"name"`
         
    }
    
    // DimensionResponse1 used by Catalog
    type DimensionResponse1 struct {

        
            Unit string  `json:"unit"`
            Width float64  `json:"width"`
            Height float64  `json:"height"`
            Length float64  `json:"length"`
         
    }
    
    // GetInventories used by Catalog
    type GetInventories struct {

        
            Platforms map[string]interface{}  `json:"platforms"`
            Weight WeightResponse1  `json:"weight"`
            Company CompanyMeta1  `json:"company"`
            Tags []string  `json:"tags"`
            Identifier map[string]interface{}  `json:"identifier"`
            TrackInventory bool  `json:"track_inventory"`
            CreatedBy UserSerializer  `json:"created_by"`
            InventoryUpdatedOn string  `json:"inventory_updated_on"`
            UID string  `json:"uid"`
            Stage string  `json:"stage"`
            Price PriceArticle  `json:"price"`
            Quantities QuantitiesArticle  `json:"quantities"`
            Size string  `json:"size"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            DateMeta DateMeta  `json:"date_meta"`
            Store ArticleStoreResponse  `json:"store"`
            TraceID string  `json:"trace_id"`
            TotalQuantity float64  `json:"total_quantity"`
            ReturnConfig ReturnConfig2  `json:"return_config"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            ItemID float64  `json:"item_id"`
            ExpirationDate string  `json:"expiration_date"`
            Manufacturer ManufacturerResponse1  `json:"manufacturer"`
            Trader []Trader2  `json:"trader"`
            Brand BrandMeta1  `json:"brand"`
            ID string  `json:"id"`
            SellerIdentifier string  `json:"seller_identifier"`
            IsSet bool  `json:"is_set"`
            CountryOfOrigin string  `json:"country_of_origin"`
            Dimension DimensionResponse1  `json:"dimension"`
         
    }
    
    // GetInventoriesResponse used by Catalog
    type GetInventoriesResponse struct {

        
            Items []GetInventories  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // BulkInventoryGetItems used by Catalog
    type BulkInventoryGetItems struct {

        
            CreatedBy map[string]interface{}  `json:"created_by"`
            CompanyID float64  `json:"company_id"`
            FilePath string  `json:"file_path"`
            ID string  `json:"id"`
            Stage string  `json:"stage"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Cancelled float64  `json:"cancelled"`
            Failed float64  `json:"failed"`
            IsActive bool  `json:"is_active"`
            CancelledRecords []string  `json:"cancelled_records"`
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
            FailedRecords []string  `json:"failed_records"`
            Succeed float64  `json:"succeed"`
            Total float64  `json:"total"`
         
    }
    
    // BulkInventoryGet used by Catalog
    type BulkInventoryGet struct {

        
            Items []BulkInventoryGetItems  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // InventoryJobPayload used by Catalog
    type InventoryJobPayload struct {

        
            Currency string  `json:"currency"`
            SellerIdentifier string  `json:"seller_identifier"`
            TraceID string  `json:"trace_id"`
            Quantity float64  `json:"quantity"`
            StoreCode string  `json:"store_code"`
            TotalQuantity float64  `json:"total_quantity"`
            ItemDimensionsUnitOfMeasure string  `json:"item_dimensions_unit_of_measure"`
            PriceMarked float64  `json:"price_marked"`
            Tags []string  `json:"tags"`
            Price float64  `json:"price"`
            PriceEffective float64  `json:"price_effective"`
            ItemWeightUnitOfMeasure string  `json:"item_weight_unit_of_measure"`
            ExpirationDate string  `json:"expiration_date"`
         
    }
    
    // InventoryBulkRequest used by Catalog
    type InventoryBulkRequest struct {

        
            User map[string]interface{}  `json:"user"`
            BatchID string  `json:"batch_id"`
            CompanyID float64  `json:"company_id"`
            Sizes []InventoryJobPayload  `json:"sizes"`
         
    }
    
    // InventoryExportQuantityFilter used by Catalog
    type InventoryExportQuantityFilter struct {

        
            Min float64  `json:"min"`
            Max float64  `json:"max"`
            Operators string  `json:"operators"`
         
    }
    
    // InventoryExportAdvanceOption used by Catalog
    type InventoryExportAdvanceOption struct {

        
            FromDate string  `json:"from_date"`
            Quantity InventoryExportQuantityFilter  `json:"quantity"`
            StoreIds []float64  `json:"store_ids"`
            BrandIds []float64  `json:"brand_ids"`
            ToDate string  `json:"to_date"`
         
    }
    
    // InventoryExportJob used by Catalog
    type InventoryExportJob struct {

        
            CompletedOn string  `json:"completed_on"`
            SellerID float64  `json:"seller_id"`
            Type string  `json:"type"`
            TaskID string  `json:"task_id"`
            NotificationEmails []string  `json:"notification_emails"`
            Status string  `json:"status"`
            URL string  `json:"url"`
            Filters InventoryExportAdvanceOption  `json:"filters"`
         
    }
    
    // InventoryExportRequest used by Catalog
    type InventoryExportRequest struct {

        
            Store []float64  `json:"store"`
            Brand []float64  `json:"brand"`
            Type string  `json:"type"`
         
    }
    
    // InventoryExportResponse used by Catalog
    type InventoryExportResponse struct {

        
            CreatedBy string  `json:"created_by"`
            SellerID float64  `json:"seller_id"`
            Type string  `json:"type"`
            TaskID string  `json:"task_id"`
            NotificationEmails []string  `json:"notification_emails"`
            Status string  `json:"status"`
            ModifiedOn string  `json:"modified_on"`
            Filters map[string]interface{}  `json:"filters"`
            CreatedOn string  `json:"created_on"`
         
    }
    
    // InventoryJobFilters used by Catalog
    type InventoryJobFilters struct {

        
            FromDate string  `json:"from_date"`
            Quantity InventoryExportQuantityFilter  `json:"quantity"`
            Brands []string  `json:"brands"`
            Stores []string  `json:"stores"`
            ToDate string  `json:"to_date"`
         
    }
    
    // InventoryJobDetailResponse used by Catalog
    type InventoryJobDetailResponse struct {

        
            CompletedOn string  `json:"completed_on"`
            CreatedBy UserDetail  `json:"created_by"`
            CancelledOn string  `json:"cancelled_on"`
            SellerID float64  `json:"seller_id"`
            CancelledBy UserDetail  `json:"cancelled_by"`
            Type string  `json:"type"`
            TaskID string  `json:"task_id"`
            NotificationEmails []string  `json:"notification_emails"`
            ID string  `json:"id"`
            Status string  `json:"status"`
            URL string  `json:"url"`
            ModifiedOn string  `json:"modified_on"`
            Filters InventoryJobFilters  `json:"filters"`
            CreatedOn string  `json:"created_on"`
         
    }
    
    // InventoryExportJobListResponse used by Catalog
    type InventoryExportJobListResponse struct {

        
            Items InventoryJobDetailResponse  `json:"items"`
         
    }
    
    // InventoryExportFilter used by Catalog
    type InventoryExportFilter struct {

        
            FromDate string  `json:"from_date"`
            Quantity InventoryExportQuantityFilter  `json:"quantity"`
            StoreIds []float64  `json:"store_ids"`
            BrandIds []float64  `json:"brand_ids"`
            ToDate string  `json:"to_date"`
         
    }
    
    // InventoryCreateRequest used by Catalog
    type InventoryCreateRequest struct {

        
            NotificationEmails []string  `json:"notification_emails"`
            Filters InventoryExportFilter  `json:"filters"`
            Type string  `json:"type"`
            Data []string  `json:"data"`
         
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
            TraceID string  `json:"trace_id"`
            TotalQuantity float64  `json:"total_quantity"`
            PriceMarked float64  `json:"price_marked"`
            Tags []string  `json:"tags"`
            PriceEffective float64  `json:"price_effective"`
            StoreID float64  `json:"store_id"`
            ExpirationDate string  `json:"expiration_date"`
         
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

        
            Threshold2 float64  `json:"threshold2"`
            CompanyID float64  `json:"company_id"`
            UID float64  `json:"uid"`
            TaxOnMrp bool  `json:"tax_on_mrp"`
            Threshold1 float64  `json:"threshold1"`
            Tax1 float64  `json:"tax1"`
            TaxOnEsp bool  `json:"tax_on_esp"`
            IsActive bool  `json:"is_active"`
            HsnCode string  `json:"hsn_code"`
            Tax2 float64  `json:"tax2"`
            Hs2Code string  `json:"hs2_code"`
         
    }
    
    // HsnCodesObject used by Catalog
    type HsnCodesObject struct {

        
            Threshold2 float64  `json:"threshold2"`
            CompanyID float64  `json:"company_id"`
            TaxOnMrp bool  `json:"tax_on_mrp"`
            Threshold1 float64  `json:"threshold1"`
            ID string  `json:"id"`
            Tax1 float64  `json:"tax1"`
            TaxOnEsp bool  `json:"tax_on_esp"`
            ModifiedOn string  `json:"modified_on"`
            HsnCode string  `json:"hsn_code"`
            Tax2 float64  `json:"tax2"`
            Hs2Code string  `json:"hs2_code"`
         
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

        
            EffectiveDate string  `json:"effective_date"`
            Rate float64  `json:"rate"`
            Threshold float64  `json:"threshold"`
            Cess float64  `json:"cess"`
         
    }
    
    // HSNDataInsertV2 used by Catalog
    type HSNDataInsertV2 struct {

        
            CreatedBy map[string]interface{}  `json:"created_by"`
            Type string  `json:"type"`
            Taxes []TaxSlab  `json:"taxes"`
            CountryCode string  `json:"country_code"`
            Description string  `json:"description"`
            ModifiedOn string  `json:"modified_on"`
            HsnCodeID string  `json:"hsn_code_id"`
            HsnCode string  `json:"hsn_code"`
            CreatedOn string  `json:"created_on"`
            ReportingHsn string  `json:"reporting_hsn"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
         
    }
    
    // PageResponse used by Catalog
    type PageResponse struct {

        
            HasPrevious bool  `json:"has_previous"`
            HasNext bool  `json:"has_next"`
            ItemTotal float64  `json:"item_total"`
            Current string  `json:"current"`
            Size float64  `json:"size"`
         
    }
    
    // HsnCodesListingResponseSchemaV2 used by Catalog
    type HsnCodesListingResponseSchemaV2 struct {

        
            Items []HSNDataInsertV2  `json:"items"`
            Page PageResponse  `json:"page"`
         
    }
    
    // BrandItem used by Catalog
    type BrandItem struct {

        
            Discount string  `json:"discount"`
            Departments []string  `json:"departments"`
            UID float64  `json:"uid"`
            Banners ImageUrls  `json:"banners"`
            Slug string  `json:"slug"`
            Logo Media1  `json:"logo"`
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

        
            UID float64  `json:"uid"`
            Slug string  `json:"slug"`
            PriorityOrder float64  `json:"priority_order"`
            Logo Media1  `json:"logo"`
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

        
            UID float64  `json:"uid"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Banners ImageUrls  `json:"banners"`
            Slug string  `json:"slug"`
            Childs []map[string]interface{}  `json:"childs"`
            Name string  `json:"name"`
            Action Action  `json:"action"`
         
    }
    
    // SecondLevelChild used by Catalog
    type SecondLevelChild struct {

        
            UID float64  `json:"uid"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Banners ImageUrls  `json:"banners"`
            Slug string  `json:"slug"`
            Childs []ThirdLevelChild  `json:"childs"`
            Name string  `json:"name"`
            Action Action  `json:"action"`
         
    }
    
    // Child used by Catalog
    type Child struct {

        
            UID float64  `json:"uid"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Banners ImageUrls  `json:"banners"`
            Slug string  `json:"slug"`
            Childs []SecondLevelChild  `json:"childs"`
            Name string  `json:"name"`
            Action Action  `json:"action"`
         
    }
    
    // CategoryItems used by Catalog
    type CategoryItems struct {

        
            UID float64  `json:"uid"`
            Banners ImageUrls  `json:"banners"`
            Slug string  `json:"slug"`
            Childs []Child  `json:"childs"`
            Name string  `json:"name"`
            Action Action  `json:"action"`
         
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

        
            SortOn []ProductSortOn  `json:"sort_on"`
            Filters []ProductFilters  `json:"filters"`
            Items []ProductListingDetail  `json:"items"`
            Page Page  `json:"page"`
            Operators map[string]interface{}  `json:"operators"`
         
    }
    
    // ProductDetail used by Catalog
    type ProductDetail struct {

        
            Similars []string  `json:"similars"`
            ItemType string  `json:"item_type"`
            Attributes map[string]interface{}  `json:"attributes"`
            Rating float64  `json:"rating"`
            ItemCode string  `json:"item_code"`
            Medias []Media2  `json:"medias"`
            UID float64  `json:"uid"`
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
            ShortDescription string  `json:"short_description"`
            PromoMeta map[string]interface{}  `json:"promo_meta"`
            Name string  `json:"name"`
            Tryouts []string  `json:"tryouts"`
            ImageNature string  `json:"image_nature"`
            RatingCount float64  `json:"rating_count"`
            Color string  `json:"color"`
            Highlights []string  `json:"highlights"`
            HasVariant bool  `json:"has_variant"`
            Description string  `json:"description"`
            Brand ProductBrand  `json:"brand"`
            Type string  `json:"type"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            Slug string  `json:"slug"`
            ProductOnlineDate string  `json:"product_online_date"`
         
    }
    
    // InventoryPage used by Catalog
    type InventoryPage struct {

        
            Type string  `json:"type"`
            HasPrevious bool  `json:"has_previous"`
            HasNext bool  `json:"has_next"`
            ItemTotal float64  `json:"item_total"`
            NextID string  `json:"next_id"`
         
    }
    
    // InventoryStockResponse used by Catalog
    type InventoryStockResponse struct {

        
            Items []map[string]interface{}  `json:"items"`
            Page InventoryPage  `json:"page"`
         
    }
    
    // ArticleQuery used by Catalog
    type ArticleQuery struct {

        
            ItemID float64  `json:"item_id"`
            Size string  `json:"size"`
            IgnoredStores []float64  `json:"ignored_stores"`
         
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

        
            Pincode string  `json:"pincode"`
            ChannelIdentifier string  `json:"channel_identifier"`
            CompanyID float64  `json:"company_id"`
            Articles []AssignStoreArticle  `json:"articles"`
            AppID string  `json:"app_id"`
            StoreIds []float64  `json:"store_ids"`
            ChannelType string  `json:"channel_type"`
         
    }
    
    // ArticleAssignment1 used by Catalog
    type ArticleAssignment1 struct {

        
            Level string  `json:"level"`
            Strategy string  `json:"strategy"`
         
    }
    
    // StoreAssignResponse used by Catalog
    type StoreAssignResponse struct {

        
            GroupID string  `json:"group_id"`
            CompanyID float64  `json:"company_id"`
            Index float64  `json:"index"`
            Meta map[string]interface{}  `json:"meta"`
            UID string  `json:"uid"`
            StrategyWiseListing []map[string]interface{}  `json:"strategy_wise_listing"`
            Quantity float64  `json:"quantity"`
            Status bool  `json:"status"`
            PriceEffective float64  `json:"price_effective"`
            ArticleAssignment ArticleAssignment1  `json:"article_assignment"`
            PriceMarked float64  `json:"price_marked"`
            SCity string  `json:"s_city"`
            StorePincode float64  `json:"store_pincode"`
            ID string  `json:"_id"`
            ItemID float64  `json:"item_id"`
            Size string  `json:"size"`
            StoreID float64  `json:"store_id"`
         
    }
    
    // SellerPhoneNumber used by Catalog
    type SellerPhoneNumber struct {

        
            CountryCode float64  `json:"country_code"`
            Number string  `json:"number"`
         
    }
    
    // LocationManagerSerializer used by Catalog
    type LocationManagerSerializer struct {

        
            MobileNo SellerPhoneNumber  `json:"mobile_no"`
            Email string  `json:"email"`
            Name string  `json:"name"`
         
    }
    
    // UserSerializer1 used by Catalog
    type UserSerializer1 struct {

        
            Contact string  `json:"contact"`
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
         
    }
    
    // GetAddressSerializer used by Catalog
    type GetAddressSerializer struct {

        
            Pincode float64  `json:"pincode"`
            AddressType string  `json:"address_type"`
            State string  `json:"state"`
            City string  `json:"city"`
            Latitude float64  `json:"latitude"`
            Address2 string  `json:"address2"`
            Longitude float64  `json:"longitude"`
            CountryCode string  `json:"country_code"`
            Address1 string  `json:"address1"`
            Country string  `json:"country"`
            Landmark string  `json:"landmark"`
         
    }
    
    // GetCompanySerializer used by Catalog
    type GetCompanySerializer struct {

        
            CreatedBy UserSerializer1  `json:"created_by"`
            VerifiedOn string  `json:"verified_on"`
            UID float64  `json:"uid"`
            BusinessType string  `json:"business_type"`
            Stage string  `json:"stage"`
            VerifiedBy UserSerializer1  `json:"verified_by"`
            RejectReason string  `json:"reject_reason"`
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            ModifiedBy UserSerializer1  `json:"modified_by"`
            Name string  `json:"name"`
            CompanyType string  `json:"company_type"`
         
    }
    
    // UserSerializer2 used by Catalog
    type UserSerializer2 struct {

        
            Contact string  `json:"contact"`
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
         
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
            Closing LocationTimingSerializer  `json:"closing"`
            Opening LocationTimingSerializer  `json:"opening"`
         
    }
    
    // ProductReturnConfigSerializer used by Catalog
    type ProductReturnConfigSerializer struct {

        
            OnSameStore bool  `json:"on_same_store"`
            StoreUID float64  `json:"store_uid"`
         
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
    
    // Document used by Catalog
    type Document struct {

        
            Type string  `json:"type"`
            LegalName string  `json:"legal_name"`
            URL string  `json:"url"`
            Verified bool  `json:"verified"`
            Value string  `json:"value"`
         
    }
    
    // LocationIntegrationType used by Catalog
    type LocationIntegrationType struct {

        
            Inventory string  `json:"inventory"`
            Order string  `json:"order"`
         
    }
    
    // GetLocationSerializer used by Catalog
    type GetLocationSerializer struct {

        
            VerifiedOn string  `json:"verified_on"`
            Manager LocationManagerSerializer  `json:"manager"`
            DisplayName string  `json:"display_name"`
            Company GetCompanySerializer  `json:"company"`
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
            Code string  `json:"code"`
            CreatedBy UserSerializer2  `json:"created_by"`
            Timing []LocationDayWiseSerializer  `json:"timing"`
            UID float64  `json:"uid"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ProductReturnConfig ProductReturnConfigSerializer  `json:"product_return_config"`
            GstCredentials InvoiceDetailsSerializer  `json:"gst_credentials"`
            Stage string  `json:"stage"`
            Warnings map[string]interface{}  `json:"warnings"`
            VerifiedBy UserSerializer2  `json:"verified_by"`
            Name string  `json:"name"`
            Documents []Document  `json:"documents"`
            Address GetAddressSerializer  `json:"address"`
            NotificationEmails []string  `json:"notification_emails"`
            ModifiedBy UserSerializer2  `json:"modified_by"`
            PhoneNumber string  `json:"phone_number"`
            IntegrationType LocationIntegrationType  `json:"integration_type"`
            StoreType string  `json:"store_type"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
         
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
            CustomJson map[string]interface{}  `json:"_custom_json"`
            IsActive bool  `json:"is_active"`
            AppID string  `json:"app_id"`
            Logo string  `json:"logo"`
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
    

    
    // GetAddressSerializer used by CompanyProfile
    type GetAddressSerializer struct {

        
            AddressType string  `json:"address_type"`
            Pincode float64  `json:"pincode"`
            State string  `json:"state"`
            CountryCode string  `json:"country_code"`
            Country string  `json:"country"`
            Latitude float64  `json:"latitude"`
            Landmark string  `json:"landmark"`
            Address2 string  `json:"address2"`
            Address1 string  `json:"address1"`
            City string  `json:"city"`
            Longitude float64  `json:"longitude"`
         
    }
    
    // BusinessCountryInfo used by CompanyProfile
    type BusinessCountryInfo struct {

        
            Country string  `json:"country"`
            CountryCode string  `json:"country_code"`
         
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
    
    // Document used by CompanyProfile
    type Document struct {

        
            Value string  `json:"value"`
            Type string  `json:"type"`
            URL string  `json:"url"`
            LegalName string  `json:"legal_name"`
            Verified bool  `json:"verified"`
         
    }
    
    // UserSerializer used by CompanyProfile
    type UserSerializer struct {

        
            Contact string  `json:"contact"`
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
         
    }
    
    // CompanyTaxesSerializer used by CompanyProfile
    type CompanyTaxesSerializer struct {

        
            Enable bool  `json:"enable"`
            EffectiveDate string  `json:"effective_date"`
            Rate float64  `json:"rate"`
         
    }
    
    // Website used by CompanyProfile
    type Website struct {

        
            URL string  `json:"url"`
         
    }
    
    // BusinessDetails used by CompanyProfile
    type BusinessDetails struct {

        
            Website Website  `json:"website"`
         
    }
    
    // GetCompanyProfileSerializerResponse used by CompanyProfile
    type GetCompanyProfileSerializerResponse struct {

        
            Addresses []GetAddressSerializer  `json:"addresses"`
            BusinessCountryInfo BusinessCountryInfo  `json:"business_country_info"`
            ContactDetails ContactDetails  `json:"contact_details"`
            Documents []Document  `json:"documents"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            NotificationEmails []string  `json:"notification_emails"`
            Stage string  `json:"stage"`
            BusinessType string  `json:"business_type"`
            Taxes []CompanyTaxesSerializer  `json:"taxes"`
            CompanyType string  `json:"company_type"`
            Name string  `json:"name"`
            Warnings map[string]interface{}  `json:"warnings"`
            CreatedOn string  `json:"created_on"`
            Mode string  `json:"mode"`
            UID float64  `json:"uid"`
            BusinessDetails BusinessDetails  `json:"business_details"`
            FranchiseEnabled bool  `json:"franchise_enabled"`
            VerifiedOn string  `json:"verified_on"`
            BusinessInfo string  `json:"business_info"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ModifiedOn string  `json:"modified_on"`
            CreatedBy UserSerializer  `json:"created_by"`
         
    }
    
    // ErrorResponse used by CompanyProfile
    type ErrorResponse struct {

        
            Status float64  `json:"status"`
            Meta map[string]interface{}  `json:"meta"`
            Message string  `json:"message"`
            Code string  `json:"code"`
         
    }
    
    // CreateUpdateAddressSerializer used by CompanyProfile
    type CreateUpdateAddressSerializer struct {

        
            AddressType string  `json:"address_type"`
            Pincode float64  `json:"pincode"`
            State string  `json:"state"`
            CountryCode string  `json:"country_code"`
            Country string  `json:"country"`
            Latitude float64  `json:"latitude"`
            Landmark string  `json:"landmark"`
            Address1 string  `json:"address1"`
            Address2 string  `json:"address2"`
            City string  `json:"city"`
            Longitude float64  `json:"longitude"`
         
    }
    
    // CompanyTaxesSerializer1 used by CompanyProfile
    type CompanyTaxesSerializer1 struct {

        
            Enable bool  `json:"enable"`
            EffectiveDate string  `json:"effective_date"`
            Rate float64  `json:"rate"`
         
    }
    
    // UpdateCompany used by CompanyProfile
    type UpdateCompany struct {

        
            NotificationEmails []string  `json:"notification_emails"`
            Addresses []CreateUpdateAddressSerializer  `json:"addresses"`
            RejectReason string  `json:"reject_reason"`
            Documents []Document  `json:"documents"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            BusinessType string  `json:"business_type"`
            ContactDetails ContactDetails  `json:"contact_details"`
            BusinessDetails BusinessDetails  `json:"business_details"`
            FranchiseEnabled bool  `json:"franchise_enabled"`
            Taxes []CompanyTaxesSerializer1  `json:"taxes"`
            CompanyType string  `json:"company_type"`
            Warnings map[string]interface{}  `json:"warnings"`
            Name string  `json:"name"`
            BusinessInfo string  `json:"business_info"`
         
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

        
            UID float64  `json:"uid"`
            Store DocumentsObj  `json:"store"`
            Brand DocumentsObj  `json:"brand"`
            Stage string  `json:"stage"`
            Product DocumentsObj  `json:"product"`
            StoreDocuments DocumentsObj  `json:"store_documents"`
            CompanyDocuments DocumentsObj  `json:"company_documents"`
         
    }
    
    // BrandBannerSerializer used by CompanyProfile
    type BrandBannerSerializer struct {

        
            Landscape string  `json:"landscape"`
            Portrait string  `json:"portrait"`
         
    }
    
    // GetBrandResponseSerializer used by CompanyProfile
    type GetBrandResponseSerializer struct {

        
            ModifiedBy UserSerializer  `json:"modified_by"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            Stage string  `json:"stage"`
            Banner BrandBannerSerializer  `json:"banner"`
            Synonyms []string  `json:"synonyms"`
            Name string  `json:"name"`
            Warnings map[string]interface{}  `json:"warnings"`
            CreatedOn string  `json:"created_on"`
            Mode string  `json:"mode"`
            UID float64  `json:"uid"`
            RejectReason string  `json:"reject_reason"`
            Logo string  `json:"logo"`
            SlugKey string  `json:"slug_key"`
            VerifiedOn string  `json:"verified_on"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ModifiedOn string  `json:"modified_on"`
            CreatedBy UserSerializer  `json:"created_by"`
            Description string  `json:"description"`
         
    }
    
    // CreateUpdateBrandRequestSerializer used by CompanyProfile
    type CreateUpdateBrandRequestSerializer struct {

        
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            UID float64  `json:"uid"`
            BrandTier string  `json:"brand_tier"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Logo string  `json:"logo"`
            Description string  `json:"description"`
            Banner BrandBannerSerializer  `json:"banner"`
            CompanyID float64  `json:"company_id"`
            Synonyms []string  `json:"synonyms"`
            Name string  `json:"name"`
         
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

        
            Socials []CompanySocialAccounts  `json:"socials"`
            WebsiteURL string  `json:"website_url"`
         
    }
    
    // CompanySerializer used by CompanyProfile
    type CompanySerializer struct {

        
            ModifiedBy UserSerializer  `json:"modified_by"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            CompanyType string  `json:"company_type"`
            NotificationEmails []string  `json:"notification_emails"`
            UID float64  `json:"uid"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            RejectReason string  `json:"reject_reason"`
            Stage string  `json:"stage"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            BusinessType string  `json:"business_type"`
            MarketChannels []string  `json:"market_channels"`
            BusinessCountryInfo BusinessCountryInfo  `json:"business_country_info"`
            ModifiedOn string  `json:"modified_on"`
            CreatedBy UserSerializer  `json:"created_by"`
            CreatedOn string  `json:"created_on"`
            Name string  `json:"name"`
            VerifiedOn string  `json:"verified_on"`
            Details CompanyDetails  `json:"details"`
         
    }
    
    // CompanyBrandSerializer used by CompanyProfile
    type CompanyBrandSerializer struct {

        
            ModifiedBy UserSerializer  `json:"modified_by"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            UID float64  `json:"uid"`
            RejectReason string  `json:"reject_reason"`
            Brand GetBrandResponseSerializer  `json:"brand"`
            Stage string  `json:"stage"`
            ModifiedOn string  `json:"modified_on"`
            Company CompanySerializer  `json:"company"`
            CreatedBy UserSerializer  `json:"created_by"`
            CreatedOn string  `json:"created_on"`
            Warnings map[string]interface{}  `json:"warnings"`
            VerifiedOn string  `json:"verified_on"`
         
    }
    
    // CompanyBrandListSerializer used by CompanyProfile
    type CompanyBrandListSerializer struct {

        
            Page Page  `json:"page"`
            Items []CompanyBrandSerializer  `json:"items"`
         
    }
    
    // CompanyBrandPostRequestSerializer used by CompanyProfile
    type CompanyBrandPostRequestSerializer struct {

        
            Brands []float64  `json:"brands"`
            UID float64  `json:"uid"`
            Company float64  `json:"company"`
         
    }
    
    // LocationTimingSerializer used by CompanyProfile
    type LocationTimingSerializer struct {

        
            Hour float64  `json:"hour"`
            Minute float64  `json:"minute"`
         
    }
    
    // LocationDayWiseSerializer used by CompanyProfile
    type LocationDayWiseSerializer struct {

        
            Open bool  `json:"open"`
            Opening LocationTimingSerializer  `json:"opening"`
            Closing LocationTimingSerializer  `json:"closing"`
            Weekday string  `json:"weekday"`
         
    }
    
    // InvoiceCredSerializer used by CompanyProfile
    type InvoiceCredSerializer struct {

        
            Username string  `json:"username"`
            Password string  `json:"password"`
            Enabled bool  `json:"enabled"`
         
    }
    
    // InvoiceDetailsSerializer used by CompanyProfile
    type InvoiceDetailsSerializer struct {

        
            EWaybill InvoiceCredSerializer  `json:"e_waybill"`
            EInvoice InvoiceCredSerializer  `json:"e_invoice"`
         
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
    
    // ProductReturnConfigSerializer used by CompanyProfile
    type ProductReturnConfigSerializer struct {

        
            StoreUID float64  `json:"store_uid"`
            OnSameStore bool  `json:"on_same_store"`
         
    }
    
    // GetCompanySerializer used by CompanyProfile
    type GetCompanySerializer struct {

        
            ModifiedBy UserSerializer  `json:"modified_by"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            CompanyType string  `json:"company_type"`
            UID float64  `json:"uid"`
            RejectReason string  `json:"reject_reason"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            Stage string  `json:"stage"`
            BusinessType string  `json:"business_type"`
            ModifiedOn string  `json:"modified_on"`
            CreatedBy UserSerializer  `json:"created_by"`
            CreatedOn string  `json:"created_on"`
            Name string  `json:"name"`
            VerifiedOn string  `json:"verified_on"`
         
    }
    
    // LocationManagerSerializer used by CompanyProfile
    type LocationManagerSerializer struct {

        
            MobileNo SellerPhoneNumber  `json:"mobile_no"`
            Email string  `json:"email"`
            Name string  `json:"name"`
         
    }
    
    // GetLocationSerializer used by CompanyProfile
    type GetLocationSerializer struct {

        
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            Timing []LocationDayWiseSerializer  `json:"timing"`
            StoreType string  `json:"store_type"`
            PhoneNumber string  `json:"phone_number"`
            Documents []Document  `json:"documents"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            NotificationEmails []string  `json:"notification_emails"`
            Stage string  `json:"stage"`
            DisplayName string  `json:"display_name"`
            AutoInvoice bool  `json:"auto_invoice"`
            Name string  `json:"name"`
            Warnings map[string]interface{}  `json:"warnings"`
            CreatedOn string  `json:"created_on"`
            UID float64  `json:"uid"`
            Code string  `json:"code"`
            GstCredentials InvoiceDetailsSerializer  `json:"gst_credentials"`
            VerifiedOn string  `json:"verified_on"`
            Holiday []HolidaySchemaSerializer  `json:"holiday"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            CreditNote bool  `json:"credit_note"`
            ProductReturnConfig ProductReturnConfigSerializer  `json:"product_return_config"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ModifiedOn string  `json:"modified_on"`
            Address GetAddressSerializer  `json:"address"`
            Company GetCompanySerializer  `json:"company"`
            CreatedBy UserSerializer  `json:"created_by"`
            Manager LocationManagerSerializer  `json:"manager"`
         
    }
    
    // LocationListSerializer used by CompanyProfile
    type LocationListSerializer struct {

        
            Page Page  `json:"page"`
            Items []GetLocationSerializer  `json:"items"`
         
    }
    
    // AddressSerializer used by CompanyProfile
    type AddressSerializer struct {

        
            AddressType string  `json:"address_type"`
            Pincode float64  `json:"pincode"`
            State string  `json:"state"`
            CountryCode string  `json:"country_code"`
            Country string  `json:"country"`
            Latitude float64  `json:"latitude"`
            Landmark string  `json:"landmark"`
            Address2 string  `json:"address2"`
            Address1 string  `json:"address1"`
            City string  `json:"city"`
            Longitude float64  `json:"longitude"`
         
    }
    
    // LocationSerializer used by CompanyProfile
    type LocationSerializer struct {

        
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            Timing []LocationDayWiseSerializer  `json:"timing"`
            StoreType string  `json:"store_type"`
            Documents []Document  `json:"documents"`
            NotificationEmails []string  `json:"notification_emails"`
            Stage string  `json:"stage"`
            DisplayName string  `json:"display_name"`
            AutoInvoice bool  `json:"auto_invoice"`
            Name string  `json:"name"`
            Warnings map[string]interface{}  `json:"warnings"`
            UID float64  `json:"uid"`
            Code string  `json:"code"`
            GstCredentials InvoiceDetailsSerializer  `json:"gst_credentials"`
            Holiday []HolidaySchemaSerializer  `json:"holiday"`
            CreditNote bool  `json:"credit_note"`
            ProductReturnConfig ProductReturnConfigSerializer  `json:"product_return_config"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Address AddressSerializer  `json:"address"`
            Company float64  `json:"company"`
            Manager LocationManagerSerializer  `json:"manager"`
         
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
    
    // Ownership used by Cart
    type Ownership struct {

        
            PayableCategory string  `json:"payable_category"`
            PayableBy string  `json:"payable_by"`
         
    }
    
    // Identifier used by Cart
    type Identifier struct {

        
            BrandID []float64  `json:"brand_id"`
            CompanyID []float64  `json:"company_id"`
            ExcludeBrandID []float64  `json:"exclude_brand_id"`
            CollectionID []string  `json:"collection_id"`
            StoreID []float64  `json:"store_id"`
            CategoryID []float64  `json:"category_id"`
            ArticleID []string  `json:"article_id"`
            ItemID []float64  `json:"item_id"`
            UserID []string  `json:"user_id"`
         
    }
    
    // Validity used by Cart
    type Validity struct {

        
            Priority float64  `json:"priority"`
         
    }
    
    // Validation used by Cart
    type Validation struct {

        
            Anonymous bool  `json:"anonymous"`
            UserRegisteredAfter string  `json:"user_registered_after"`
            AppID []string  `json:"app_id"`
         
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
    
    // PriceRange used by Cart
    type PriceRange struct {

        
            Max float64  `json:"max"`
            Min float64  `json:"min"`
         
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

        
            Codes []string  `json:"codes"`
            Networks []string  `json:"networks"`
            Types []string  `json:"types"`
            Uses PaymentAllowValue  `json:"uses"`
         
    }
    
    // Restrictions used by Cart
    type Restrictions struct {

        
            OrderingStores []float64  `json:"ordering_stores"`
            UserGroups []float64  `json:"user_groups"`
            CouponAllowed bool  `json:"coupon_allowed"`
            PostOrder PostOrder  `json:"post_order"`
            BulkBundle BulkBundleRestriction  `json:"bulk_bundle"`
            Platforms []string  `json:"platforms"`
            PriceRange PriceRange  `json:"price_range"`
            Uses UsesRestriction  `json:"uses"`
            Payments map[string]PaymentModes  `json:"payments"`
         
    }
    
    // CouponSchedule used by Cart
    type CouponSchedule struct {

        
            NextSchedule []map[string]interface{}  `json:"next_schedule"`
            Cron string  `json:"cron"`
            Start string  `json:"start"`
            Duration float64  `json:"duration"`
            End string  `json:"end"`
         
    }
    
    // DisplayMetaDict used by Cart
    type DisplayMetaDict struct {

        
            Title string  `json:"title"`
            Subtitle string  `json:"subtitle"`
         
    }
    
    // DisplayMeta used by Cart
    type DisplayMeta struct {

        
            Apply DisplayMetaDict  `json:"apply"`
            Title string  `json:"title"`
            Subtitle string  `json:"subtitle"`
            Remove DisplayMetaDict  `json:"remove"`
            Description string  `json:"description"`
            Auto DisplayMetaDict  `json:"auto"`
         
    }
    
    // CouponDateMeta used by Cart
    type CouponDateMeta struct {

        
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
         
    }
    
    // CouponAction used by Cart
    type CouponAction struct {

        
            ActionDate string  `json:"action_date"`
            TxnMode string  `json:"txn_mode"`
         
    }
    
    // Rule used by Cart
    type Rule struct {

        
            Max float64  `json:"max"`
            Key float64  `json:"key"`
            Min float64  `json:"min"`
            Value float64  `json:"value"`
            DiscountQty float64  `json:"discount_qty"`
         
    }
    
    // State used by Cart
    type State struct {

        
            IsDisplay bool  `json:"is_display"`
            IsArchived bool  `json:"is_archived"`
            IsPublic bool  `json:"is_public"`
         
    }
    
    // CouponAuthor used by Cart
    type CouponAuthor struct {

        
            ModifiedBy string  `json:"modified_by"`
            CreatedBy string  `json:"created_by"`
         
    }
    
    // RuleDefinition used by Cart
    type RuleDefinition struct {

        
            CalculateOn string  `json:"calculate_on"`
            CurrencyCode string  `json:"currency_code"`
            Scope []string  `json:"scope"`
            Type string  `json:"type"`
            IsExact bool  `json:"is_exact"`
            ValueType string  `json:"value_type"`
            AutoApply bool  `json:"auto_apply"`
            ApplicableOn string  `json:"applicable_on"`
         
    }
    
    // CouponAdd used by Cart
    type CouponAdd struct {

        
            Tags []string  `json:"tags"`
            Ownership Ownership  `json:"ownership"`
            Identifiers Identifier  `json:"identifiers"`
            Validity Validity  `json:"validity"`
            Validation Validation  `json:"validation"`
            TypeSlug string  `json:"type_slug"`
            Restrictions Restrictions  `json:"restrictions"`
            Schedule CouponSchedule  `json:"_schedule"`
            Code string  `json:"code"`
            DisplayMeta DisplayMeta  `json:"display_meta"`
            DateMeta CouponDateMeta  `json:"date_meta"`
            Action CouponAction  `json:"action"`
            Rule []Rule  `json:"rule"`
            State State  `json:"state"`
            Author CouponAuthor  `json:"author"`
            RuleDefinition RuleDefinition  `json:"rule_definition"`
         
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

        
            Tags []string  `json:"tags"`
            Ownership Ownership  `json:"ownership"`
            Identifiers Identifier  `json:"identifiers"`
            Validity Validity  `json:"validity"`
            Validation Validation  `json:"validation"`
            TypeSlug string  `json:"type_slug"`
            Restrictions Restrictions  `json:"restrictions"`
            Schedule CouponSchedule  `json:"_schedule"`
            Code string  `json:"code"`
            DisplayMeta DisplayMeta  `json:"display_meta"`
            DateMeta CouponDateMeta  `json:"date_meta"`
            Action CouponAction  `json:"action"`
            Rule []Rule  `json:"rule"`
            State State  `json:"state"`
            Author CouponAuthor  `json:"author"`
            RuleDefinition RuleDefinition  `json:"rule_definition"`
         
    }
    
    // CouponPartialUpdate used by Cart
    type CouponPartialUpdate struct {

        
            Archive bool  `json:"archive"`
            Schedule CouponSchedule  `json:"schedule"`
         
    }
    
    // Ownership1 used by Cart
    type Ownership1 struct {

        
            PayableCategory string  `json:"payable_category"`
            PayableBy string  `json:"payable_by"`
         
    }
    
    // PromotionDateMeta used by Cart
    type PromotionDateMeta struct {

        
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
         
    }
    
    // PromotionSchedule used by Cart
    type PromotionSchedule struct {

        
            NextSchedule []map[string]interface{}  `json:"next_schedule"`
            Cron string  `json:"cron"`
            Start string  `json:"start"`
            Duration float64  `json:"duration"`
            End string  `json:"end"`
            Published bool  `json:"published"`
         
    }
    
    // PromotionAction used by Cart
    type PromotionAction struct {

        
            ActionType string  `json:"action_type"`
            ActionDate string  `json:"action_date"`
         
    }
    
    // PromotionAuthor used by Cart
    type PromotionAuthor struct {

        
            ModifiedBy string  `json:"modified_by"`
            CreatedBy string  `json:"created_by"`
         
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

        
            Total float64  `json:"total"`
            User float64  `json:"user"`
         
    }
    
    // UsesRestriction1 used by Cart
    type UsesRestriction1 struct {

        
            Maximum UsesRemaining1  `json:"maximum"`
            Remaining UsesRemaining1  `json:"remaining"`
         
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
    
    // Restrictions1 used by Cart
    type Restrictions1 struct {

        
            OrderingStores []float64  `json:"ordering_stores"`
            UserGroups []float64  `json:"user_groups"`
            PostOrder PostOrder1  `json:"post_order"`
            OrderQuantity float64  `json:"order_quantity"`
            UserRegistered UserRegistered  `json:"user_registered"`
            Platforms []string  `json:"platforms"`
            AnonymousUsers bool  `json:"anonymous_users"`
            Uses UsesRestriction1  `json:"uses"`
            Payments []PromotionPaymentModes  `json:"payments"`
            UserID []string  `json:"user_id"`
         
    }
    
    // Visibility used by Cart
    type Visibility struct {

        
            CouponList bool  `json:"coupon_list"`
            Pdp bool  `json:"pdp"`
         
    }
    
    // CompareObject used by Cart
    type CompareObject struct {

        
            Equals float64  `json:"equals"`
            LessThanEquals float64  `json:"less_than_equals"`
            GreaterThan float64  `json:"greater_than"`
            LessThan float64  `json:"less_than"`
            GreaterThanEquals float64  `json:"greater_than_equals"`
         
    }
    
    // ItemCriteria used by Cart
    type ItemCriteria struct {

        
            ItemExcludeCompany []float64  `json:"item_exclude_company"`
            ItemCompany []float64  `json:"item_company"`
            ItemExcludeSku []string  `json:"item_exclude_sku"`
            ItemExcludeCategory []float64  `json:"item_exclude_category"`
            ItemExcludeBrand []float64  `json:"item_exclude_brand"`
            ItemStore []float64  `json:"item_store"`
            CartUniqueItemAmount CompareObject  `json:"cart_unique_item_amount"`
            ItemExcludeStore []float64  `json:"item_exclude_store"`
            ItemBrand []float64  `json:"item_brand"`
            ItemSize []string  `json:"item_size"`
            ItemCategory []float64  `json:"item_category"`
            CartUniqueItemQuantity CompareObject  `json:"cart_unique_item_quantity"`
            ItemExcludeID []float64  `json:"item_exclude_id"`
            ItemID []float64  `json:"item_id"`
            CartTotal CompareObject  `json:"cart_total"`
            ItemSku []string  `json:"item_sku"`
            BuyRules []string  `json:"buy_rules"`
            CartQuantity CompareObject  `json:"cart_quantity"`
            AvailableZones []string  `json:"available_zones"`
            AllItems bool  `json:"all_items"`
         
    }
    
    // DiscountOffer used by Cart
    type DiscountOffer struct {

        
            PartialCanRet bool  `json:"partial_can_ret"`
            ApportionDiscount bool  `json:"apportion_discount"`
            MaxUsagePerTransaction float64  `json:"max_usage_per_transaction"`
            MaxOfferQuantity float64  `json:"max_offer_quantity"`
            Code string  `json:"code"`
            DiscountPrice float64  `json:"discount_price"`
            DiscountPercentage float64  `json:"discount_percentage"`
            DiscountAmount float64  `json:"discount_amount"`
            MaxDiscountAmount float64  `json:"max_discount_amount"`
            MinOfferQuantity float64  `json:"min_offer_quantity"`
         
    }
    
    // DiscountRule used by Cart
    type DiscountRule struct {

        
            BuyCondition string  `json:"buy_condition"`
            ItemCriteria ItemCriteria  `json:"item_criteria"`
            Offer DiscountOffer  `json:"offer"`
            DiscountType string  `json:"discount_type"`
         
    }
    
    // DisplayMeta1 used by Cart
    type DisplayMeta1 struct {

        
            OfferText string  `json:"offer_text"`
            Description string  `json:"description"`
            Name string  `json:"name"`
         
    }
    
    // PromotionListItem used by Cart
    type PromotionListItem struct {

        
            Ownership Ownership1  `json:"ownership"`
            Mode string  `json:"mode"`
            DateMeta PromotionDateMeta  `json:"date_meta"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            PromotionType string  `json:"promotion_type"`
            ApplyPriority float64  `json:"apply_priority"`
            Schedule PromotionSchedule  `json:"_schedule"`
            ApplyExclusive string  `json:"apply_exclusive"`
            PostOrderAction PromotionAction  `json:"post_order_action"`
            Author PromotionAuthor  `json:"author"`
            PromoGroup string  `json:"promo_group"`
            CalculateOn string  `json:"calculate_on"`
            Restrictions Restrictions1  `json:"restrictions"`
            Code string  `json:"code"`
            ApplyAllDiscount bool  `json:"apply_all_discount"`
            Visiblility Visibility  `json:"visiblility"`
            Currency string  `json:"currency"`
            DiscountRules []DiscountRule  `json:"discount_rules"`
            BuyRules map[string]ItemCriteria  `json:"buy_rules"`
            DisplayMeta DisplayMeta1  `json:"display_meta"`
            Stackable bool  `json:"stackable"`
            ApplicationID string  `json:"application_id"`
         
    }
    
    // PromotionsResponse used by Cart
    type PromotionsResponse struct {

        
            Page Page  `json:"page"`
            Items PromotionListItem  `json:"items"`
         
    }
    
    // PromotionAdd used by Cart
    type PromotionAdd struct {

        
            Ownership Ownership1  `json:"ownership"`
            Mode string  `json:"mode"`
            DateMeta PromotionDateMeta  `json:"date_meta"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            PromotionType string  `json:"promotion_type"`
            ApplyPriority float64  `json:"apply_priority"`
            Schedule PromotionSchedule  `json:"_schedule"`
            ApplyExclusive string  `json:"apply_exclusive"`
            PostOrderAction PromotionAction  `json:"post_order_action"`
            Author PromotionAuthor  `json:"author"`
            PromoGroup string  `json:"promo_group"`
            CalculateOn string  `json:"calculate_on"`
            Restrictions Restrictions1  `json:"restrictions"`
            Code string  `json:"code"`
            ApplyAllDiscount bool  `json:"apply_all_discount"`
            Visiblility Visibility  `json:"visiblility"`
            Currency string  `json:"currency"`
            DiscountRules []DiscountRule  `json:"discount_rules"`
            BuyRules map[string]ItemCriteria  `json:"buy_rules"`
            DisplayMeta DisplayMeta1  `json:"display_meta"`
            Stackable bool  `json:"stackable"`
            ApplicationID string  `json:"application_id"`
         
    }
    
    // PromotionUpdate used by Cart
    type PromotionUpdate struct {

        
            Ownership Ownership1  `json:"ownership"`
            Mode string  `json:"mode"`
            DateMeta PromotionDateMeta  `json:"date_meta"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            PromotionType string  `json:"promotion_type"`
            ApplyPriority float64  `json:"apply_priority"`
            Schedule PromotionSchedule  `json:"_schedule"`
            ApplyExclusive string  `json:"apply_exclusive"`
            PostOrderAction PromotionAction  `json:"post_order_action"`
            Author PromotionAuthor  `json:"author"`
            PromoGroup string  `json:"promo_group"`
            CalculateOn string  `json:"calculate_on"`
            Restrictions Restrictions1  `json:"restrictions"`
            Code string  `json:"code"`
            ApplyAllDiscount bool  `json:"apply_all_discount"`
            Visiblility Visibility  `json:"visiblility"`
            Currency string  `json:"currency"`
            DiscountRules []DiscountRule  `json:"discount_rules"`
            BuyRules map[string]ItemCriteria  `json:"buy_rules"`
            DisplayMeta DisplayMeta1  `json:"display_meta"`
            Stackable bool  `json:"stackable"`
            ApplicationID string  `json:"application_id"`
         
    }
    
    // PromotionPartialUpdate used by Cart
    type PromotionPartialUpdate struct {

        
            Archive bool  `json:"archive"`
            Schedule PromotionSchedule  `json:"schedule"`
         
    }
    
    // ActivePromosResponse used by Cart
    type ActivePromosResponse struct {

        
            CreatedOn string  `json:"created_on"`
            Type string  `json:"type"`
            Title string  `json:"title"`
            Subtitle string  `json:"subtitle"`
            IsHidden bool  `json:"is_hidden"`
            ModifiedOn string  `json:"modified_on"`
            EntitySlug string  `json:"entity_slug"`
            Description string  `json:"description"`
            EntityType string  `json:"entity_type"`
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

        
            GiftDisplayText string  `json:"gift_display_text"`
            MaxCartItems float64  `json:"max_cart_items"`
            MinCartValue float64  `json:"min_cart_value"`
            Enabled bool  `json:"enabled"`
            DeliveryCharges DeliveryCharges  `json:"delivery_charges"`
            RevenueEngineCoupon bool  `json:"revenue_engine_coupon"`
            BulkCoupons bool  `json:"bulk_coupons"`
            GiftPricing float64  `json:"gift_pricing"`
         
    }
    
    // CartMetaConfigAdd used by Cart
    type CartMetaConfigAdd struct {

        
            GiftDisplayText string  `json:"gift_display_text"`
            MaxCartItems float64  `json:"max_cart_items"`
            MinCartValue float64  `json:"min_cart_value"`
            Enabled bool  `json:"enabled"`
            DeliveryCharges DeliveryCharges  `json:"delivery_charges"`
            RevenueEngineCoupon bool  `json:"revenue_engine_coupon"`
            BulkCoupons bool  `json:"bulk_coupons"`
            GiftPricing float64  `json:"gift_pricing"`
         
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
    
    // DisplayBreakup used by Cart
    type DisplayBreakup struct {

        
            CurrencyCode string  `json:"currency_code"`
            Key string  `json:"key"`
            CurrencySymbol string  `json:"currency_symbol"`
            Display string  `json:"display"`
            Value float64  `json:"value"`
            Message []string  `json:"message"`
         
    }
    
    // RawBreakup used by Cart
    type RawBreakup struct {

        
            YouSaved float64  `json:"you_saved"`
            Discount float64  `json:"discount"`
            MrpTotal float64  `json:"mrp_total"`
            FyndCash float64  `json:"fynd_cash"`
            CodCharge float64  `json:"cod_charge"`
            GstCharges float64  `json:"gst_charges"`
            Subtotal float64  `json:"subtotal"`
            Vog float64  `json:"vog"`
            Total float64  `json:"total"`
            ConvenienceFee float64  `json:"convenience_fee"`
            Coupon float64  `json:"coupon"`
            DeliveryCharge float64  `json:"delivery_charge"`
         
    }
    
    // CouponBreakup used by Cart
    type CouponBreakup struct {

        
            Type string  `json:"type"`
            MaxDiscountValue float64  `json:"max_discount_value"`
            CouponValue float64  `json:"coupon_value"`
            UID string  `json:"uid"`
            Title string  `json:"title"`
            Value float64  `json:"value"`
            Code string  `json:"code"`
            MinimumCartValue float64  `json:"minimum_cart_value"`
            CouponType string  `json:"coupon_type"`
            Message string  `json:"message"`
            SubTitle string  `json:"sub_title"`
            Description string  `json:"description"`
            IsApplied bool  `json:"is_applied"`
         
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

        
            Display []DisplayBreakup  `json:"display"`
            Raw RawBreakup  `json:"raw"`
            Coupon CouponBreakup  `json:"coupon"`
            LoyaltyPoints LoyaltyPoints  `json:"loyalty_points"`
         
    }
    
    // BaseInfo used by Cart
    type BaseInfo struct {

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // ProductImage used by Cart
    type ProductImage struct {

        
            AspectRatio string  `json:"aspect_ratio"`
            URL string  `json:"url"`
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

        
            Query ActionQuery  `json:"query"`
            URL string  `json:"url"`
            Type string  `json:"type"`
         
    }
    
    // CartProduct used by Cart
    type CartProduct struct {

        
            Slug string  `json:"slug"`
            Type string  `json:"type"`
            UID float64  `json:"uid"`
            Brand BaseInfo  `json:"brand"`
            Name string  `json:"name"`
            Images []ProductImage  `json:"images"`
            Categories []CategoryInfo  `json:"categories"`
            Action ProductAction  `json:"action"`
         
    }
    
    // ProductAvailabilitySize used by Cart
    type ProductAvailabilitySize struct {

        
            IsAvailable bool  `json:"is_available"`
            Value string  `json:"value"`
            Display string  `json:"display"`
         
    }
    
    // ProductAvailability used by Cart
    type ProductAvailability struct {

        
            OtherStoreQuantity float64  `json:"other_store_quantity"`
            Deliverable bool  `json:"deliverable"`
            IsValid bool  `json:"is_valid"`
            Sizes []string  `json:"sizes"`
            AvailableSizes []ProductAvailabilitySize  `json:"available_sizes"`
            OutOfStock bool  `json:"out_of_stock"`
         
    }
    
    // PromoMeta used by Cart
    type PromoMeta struct {

        
            Message string  `json:"message"`
         
    }
    
    // CartProductIdentifer used by Cart
    type CartProductIdentifer struct {

        
            Identifier string  `json:"identifier"`
         
    }
    
    // ProductPrice used by Cart
    type ProductPrice struct {

        
            CurrencyCode string  `json:"currency_code"`
            Selling float64  `json:"selling"`
            CurrencySymbol string  `json:"currency_symbol"`
            AddOn float64  `json:"add_on"`
            Effective float64  `json:"effective"`
            Marked float64  `json:"marked"`
         
    }
    
    // ProductPriceInfo used by Cart
    type ProductPriceInfo struct {

        
            Converted ProductPrice  `json:"converted"`
            Base ProductPrice  `json:"base"`
         
    }
    
    // FreeGiftItem used by Cart
    type FreeGiftItem struct {

        
            ItemImagesURL []string  `json:"item_images_url"`
            ItemPriceDetails map[string]interface{}  `json:"item_price_details"`
            ItemBrandName string  `json:"item_brand_name"`
            ItemSlug string  `json:"item_slug"`
            ItemName string  `json:"item_name"`
            ItemID float64  `json:"item_id"`
         
    }
    
    // AppliedFreeArticles used by Cart
    type AppliedFreeArticles struct {

        
            ParentItemIdentifier string  `json:"parent_item_identifier"`
            FreeGiftItemDetails FreeGiftItem  `json:"free_gift_item_details"`
            Quantity float64  `json:"quantity"`
            ArticleID string  `json:"article_id"`
         
    }
    
    // Ownership2 used by Cart
    type Ownership2 struct {

        
            PayableCategory string  `json:"payable_category"`
            PayableBy string  `json:"payable_by"`
         
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

        
            ItemCriteria map[string]interface{}  `json:"item_criteria"`
            CartConditions map[string]interface{}  `json:"cart_conditions"`
         
    }
    
    // AppliedPromotion used by Cart
    type AppliedPromotion struct {

        
            OfferText string  `json:"offer_text"`
            AppliedFreeArticles []AppliedFreeArticles  `json:"applied_free_articles"`
            Ownership Ownership2  `json:"ownership"`
            DiscountRules []DiscountRulesApp  `json:"discount_rules"`
            MrpPromotion bool  `json:"mrp_promotion"`
            PromoID string  `json:"promo_id"`
            BuyRules []BuyRules  `json:"buy_rules"`
            ArticleQuantity float64  `json:"article_quantity"`
            PromotionName string  `json:"promotion_name"`
            PromotionGroup string  `json:"promotion_group"`
            PromotionType string  `json:"promotion_type"`
            Amount float64  `json:"amount"`
         
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

        
            Converted BasePrice  `json:"converted"`
            Base BasePrice  `json:"base"`
         
    }
    
    // ProductArticle used by Cart
    type ProductArticle struct {

        
            Store BaseInfo  `json:"store"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            Price ArticlePriceInfo  `json:"price"`
            Type string  `json:"type"`
            Seller BaseInfo  `json:"seller"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            UID string  `json:"uid"`
            ProductGroupTags []string  `json:"product_group_tags"`
            Identifier map[string]interface{}  `json:"identifier"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            SellerIdentifier string  `json:"seller_identifier"`
            Quantity float64  `json:"quantity"`
            Size string  `json:"size"`
         
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
    
    // CartProductInfo used by Cart
    type CartProductInfo struct {

        
            Product CartProduct  `json:"product"`
            Availability ProductAvailability  `json:"availability"`
            PromoMeta PromoMeta  `json:"promo_meta"`
            Key string  `json:"key"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
            Price ProductPriceInfo  `json:"price"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            Discount string  `json:"discount"`
            Message string  `json:"message"`
            PromotionsApplied []AppliedPromotion  `json:"promotions_applied"`
            Moq map[string]interface{}  `json:"moq"`
            CouponMessage string  `json:"coupon_message"`
            IsSet bool  `json:"is_set"`
            Quantity float64  `json:"quantity"`
            Article ProductArticle  `json:"article"`
            PricePerUnit ProductPriceInfo  `json:"price_per_unit"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            BulkOffer map[string]interface{}  `json:"bulk_offer"`
         
    }
    
    // OpenapiCartDetailsResponse used by Cart
    type OpenapiCartDetailsResponse struct {

        
            Message string  `json:"message"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            IsValid bool  `json:"is_valid"`
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

        
            State string  `json:"state"`
            Address string  `json:"address"`
            AreaCode string  `json:"area_code"`
            Country string  `json:"country"`
            Pincode float64  `json:"pincode"`
            Area string  `json:"area"`
            Landmark string  `json:"landmark"`
            AreaCodeSlug string  `json:"area_code_slug"`
            Name string  `json:"name"`
            AddressType string  `json:"address_type"`
            CountryIsoCode string  `json:"country_iso_code"`
            City string  `json:"city"`
            CountryCode string  `json:"country_code"`
            Email string  `json:"email"`
            Phone float64  `json:"phone"`
            Meta map[string]interface{}  `json:"meta"`
            CountryPhoneCode string  `json:"country_phone_code"`
         
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

        
            Key string  `json:"key"`
            Values []string  `json:"values"`
         
    }
    
    // MultiTenderPaymentMeta used by Cart
    type MultiTenderPaymentMeta struct {

        
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            OrderID string  `json:"order_id"`
            CurrentStatus string  `json:"current_status"`
            PaymentID string  `json:"payment_id"`
            PaymentGateway string  `json:"payment_gateway"`
         
    }
    
    // MultiTenderPaymentMethod used by Cart
    type MultiTenderPaymentMethod struct {

        
            Amount float64  `json:"amount"`
            Meta MultiTenderPaymentMeta  `json:"meta"`
            Name string  `json:"name"`
            Mode string  `json:"mode"`
         
    }
    
    // CartItemMeta used by Cart
    type CartItemMeta struct {

        
            PrimaryItem bool  `json:"primary_item"`
            GroupID string  `json:"group_id"`
         
    }
    
    // OpenApiOrderItem used by Cart
    type OpenApiOrderItem struct {

        
            PriceEffective float64  `json:"price_effective"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            AmountPaid float64  `json:"amount_paid"`
            Files []OpenApiFiles  `json:"files"`
            Discount float64  `json:"discount"`
            Size string  `json:"size"`
            CashbackApplied float64  `json:"cashback_applied"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            CodCharges float64  `json:"cod_charges"`
            DeliveryCharges float64  `json:"delivery_charges"`
            Quantity float64  `json:"quantity"`
            PriceMarked float64  `json:"price_marked"`
            PaymentMethods []MultiTenderPaymentMethod  `json:"payment_methods"`
            ProductID float64  `json:"product_id"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            Meta CartItemMeta  `json:"meta"`
            EmployeeDiscount float64  `json:"employee_discount"`
         
    }
    
    // OpenApiPlatformCheckoutReq used by Cart
    type OpenApiPlatformCheckoutReq struct {

        
            CurrencyCode string  `json:"currency_code"`
            Gstin string  `json:"gstin"`
            Files []OpenApiFiles  `json:"files"`
            CartItems []OpenApiOrderItem  `json:"cart_items"`
            DeliveryCharges float64  `json:"delivery_charges"`
            CashbackApplied float64  `json:"cashback_applied"`
            CodCharges float64  `json:"cod_charges"`
            PaymentMethods []MultiTenderPaymentMethod  `json:"payment_methods"`
            Comment string  `json:"comment"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            Coupon string  `json:"coupon"`
            EmployeeDiscount map[string]interface{}  `json:"employee_discount"`
            OrderID string  `json:"order_id"`
            CouponValue float64  `json:"coupon_value"`
            CouponCode string  `json:"coupon_code"`
            BillingAddress ShippingAddress  `json:"billing_address"`
            CartValue float64  `json:"cart_value"`
            ShippingAddress ShippingAddress  `json:"shipping_address"`
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

        
            IsArchive bool  `json:"is_archive"`
            CreatedOn string  `json:"created_on"`
            Discount float64  `json:"discount"`
            DeliveryCharges map[string]interface{}  `json:"delivery_charges"`
            IsDefault bool  `json:"is_default"`
            PaymentMethods []map[string]interface{}  `json:"payment_methods"`
            Comment string  `json:"comment"`
            Coupon map[string]interface{}  `json:"coupon"`
            UserID string  `json:"user_id"`
            ExpireAt string  `json:"expire_at"`
            LastModified string  `json:"last_modified"`
            Shipments []map[string]interface{}  `json:"shipments"`
            FyndCredits map[string]interface{}  `json:"fynd_credits"`
            MergeQty bool  `json:"merge_qty"`
            Payments map[string]interface{}  `json:"payments"`
            Cashback map[string]interface{}  `json:"cashback"`
            Gstin string  `json:"gstin"`
            UID float64  `json:"uid"`
            PickUpCustomerDetails map[string]interface{}  `json:"pick_up_customer_details"`
            Promotion map[string]interface{}  `json:"promotion"`
            CheckoutMode string  `json:"checkout_mode"`
            CodCharges map[string]interface{}  `json:"cod_charges"`
            IsActive bool  `json:"is_active"`
            Articles []map[string]interface{}  `json:"articles"`
            Meta map[string]interface{}  `json:"meta"`
            BulkCouponDiscount float64  `json:"bulk_coupon_discount"`
            OrderID string  `json:"order_id"`
            BuyNow bool  `json:"buy_now"`
            ID string  `json:"_id"`
            AppID string  `json:"app_id"`
            FcIndexMap []float64  `json:"fc_index_map"`
            CartValue float64  `json:"cart_value"`
            PaymentMode string  `json:"payment_mode"`
         
    }
    
    // AbandonedCartResponse used by Cart
    type AbandonedCartResponse struct {

        
            Result map[string]interface{}  `json:"result"`
            Success bool  `json:"success"`
            Items []AbandonedCart  `json:"items"`
            Message string  `json:"message"`
            Page Page  `json:"page"`
         
    }
    
    // CartCurrency used by Cart
    type CartCurrency struct {

        
            Code string  `json:"code"`
            Symbol string  `json:"symbol"`
         
    }
    
    // PaymentSelectionLock used by Cart
    type PaymentSelectionLock struct {

        
            Enabled bool  `json:"enabled"`
            DefaultOptions string  `json:"default_options"`
            PaymentIdentifier string  `json:"payment_identifier"`
         
    }
    
    // CartDetailResponse used by Cart
    type CartDetailResponse struct {

        
            BreakupValues CartBreakup  `json:"breakup_values"`
            Gstin string  `json:"gstin"`
            Currency CartCurrency  `json:"currency"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            BuyNow bool  `json:"buy_now"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            ID string  `json:"id"`
            IsValid bool  `json:"is_valid"`
            LastModified string  `json:"last_modified"`
            Items []CartProductInfo  `json:"items"`
            CheckoutMode string  `json:"checkout_mode"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            AppliedPromoDetails []AppliedPromotion  `json:"applied_promo_details"`
            Comment string  `json:"comment"`
            Message string  `json:"message"`
            CouponText string  `json:"coupon_text"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
         
    }
    
    // AddProductCart used by Cart
    type AddProductCart struct {

        
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            SellerID float64  `json:"seller_id"`
            Display string  `json:"display"`
            Pos bool  `json:"pos"`
            ParentItemIdentifiers []map[string]string  `json:"parent_item_identifiers"`
            ProductGroupTags []string  `json:"product_group_tags"`
            StoreID float64  `json:"store_id"`
            ArticleAssignment map[string]interface{}  `json:"article_assignment"`
            ItemSize string  `json:"item_size"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Quantity float64  `json:"quantity"`
            ArticleID string  `json:"article_id"`
            ItemID float64  `json:"item_id"`
         
    }
    
    // AddCartRequest used by Cart
    type AddCartRequest struct {

        
            Items []AddProductCart  `json:"items"`
         
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

        
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            ItemSize string  `json:"item_size"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Quantity float64  `json:"quantity"`
            ItemIndex float64  `json:"item_index"`
            ArticleID string  `json:"article_id"`
            ItemID float64  `json:"item_id"`
         
    }
    
    // UpdateCartRequest used by Cart
    type UpdateCartRequest struct {

        
            Operation string  `json:"operation"`
            Items []UpdateProductCart  `json:"items"`
         
    }
    
    // UpdateCartDetailResponse used by Cart
    type UpdateCartDetailResponse struct {

        
            Message string  `json:"message"`
            Cart CartDetailResponse  `json:"cart"`
            Success bool  `json:"success"`
         
    }
    
    // OverrideCartItemPromo used by Cart
    type OverrideCartItemPromo struct {

        
            PromoID string  `json:"promo_id"`
            RwrdTndr string  `json:"rwrd_tndr"`
            PromoAmount string  `json:"promo_amount"`
            PromoDesc string  `json:"promo_desc"`
            ItemList []map[string]interface{}  `json:"item_list"`
         
    }
    
    // OverrideCartItem used by Cart
    type OverrideCartItem struct {

        
            PriceEffective float64  `json:"price_effective"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            AmountPaid float64  `json:"amount_paid"`
            Discount float64  `json:"discount"`
            Size string  `json:"size"`
            PriceMarked float64  `json:"price_marked"`
            SellerIdentifier string  `json:"seller_identifier"`
            PromoList []OverrideCartItemPromo  `json:"promo_list"`
            Quantity float64  `json:"quantity"`
            ItemID float64  `json:"item_id"`
         
    }
    
    // OverrideCheckoutReq used by Cart
    type OverrideCheckoutReq struct {

        
            CurrencyCode string  `json:"currency_code"`
            MerchantCode string  `json:"merchant_code"`
            CartID string  `json:"cart_id"`
            OrderingStore float64  `json:"ordering_store"`
            CartItems []OverrideCartItem  `json:"cart_items"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
            Aggregator string  `json:"aggregator"`
            OrderType string  `json:"order_type"`
            PaymentIdentifier string  `json:"payment_identifier"`
            ShippingAddress map[string]interface{}  `json:"shipping_address"`
            PaymentMode string  `json:"payment_mode"`
         
    }
    
    // OverrideCheckoutResponse used by Cart
    type OverrideCheckoutResponse struct {

        
            Cart map[string]interface{}  `json:"cart"`
            OrderID string  `json:"order_id"`
            Success string  `json:"success"`
            Data map[string]interface{}  `json:"data"`
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
    
