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

        
            Created bool  `json:"created"`
            Aggregators []map[string]interface{}  `json:"aggregators"`
            AppID string  `json:"app_id"`
            ExcludedFields []string  `json:"excluded_fields"`
            Success bool  `json:"success"`
            DisplayFields []string  `json:"display_fields"`
         
    }
    
    // ErrorCodeDescription used by Payment
    type ErrorCodeDescription struct {

        
            Code string  `json:"code"`
            Description string  `json:"description"`
            Success bool  `json:"success"`
         
    }
    
    // PaymentGatewayConfig used by Payment
    type PaymentGatewayConfig struct {

        
            Key string  `json:"key"`
            Secret string  `json:"secret"`
            IsActive bool  `json:"is_active"`
            ConfigType string  `json:"config_type"`
            MerchantSalt string  `json:"merchant_salt"`
         
    }
    
    // PaymentGatewayConfigRequest used by Payment
    type PaymentGatewayConfigRequest struct {

        
            AppID string  `json:"app_id"`
            IsActive bool  `json:"is_active"`
            AggregatorName PaymentGatewayConfig  `json:"aggregator_name"`
         
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
    
    // IntentAppErrorList used by Payment
    type IntentAppErrorList struct {

        
            Code string  `json:"code"`
            PackageName string  `json:"package_name"`
         
    }
    
    // IntentApp used by Payment
    type IntentApp struct {

        
            DisplayName string  `json:"display_name"`
            Logos PaymentModeLogo  `json:"logos"`
            Code string  `json:"code"`
            PackageName string  `json:"package_name"`
         
    }
    
    // PaymentModeList used by Payment
    type PaymentModeList struct {

        
            CodLimit float64  `json:"cod_limit"`
            CardFingerprint string  `json:"card_fingerprint"`
            FyndVpa string  `json:"fynd_vpa"`
            AggregatorName string  `json:"aggregator_name"`
            CardBrandImage string  `json:"card_brand_image"`
            IntentAppErrorList []string  `json:"intent_app_error_list"`
            Nickname string  `json:"nickname"`
            DisplayName string  `json:"display_name"`
            MerchantCode string  `json:"merchant_code"`
            Timeout float64  `json:"timeout"`
            LogoURL PaymentModeLogo  `json:"logo_url"`
            RetryCount float64  `json:"retry_count"`
            Expired bool  `json:"expired"`
            DisplayPriority float64  `json:"display_priority"`
            ExpMonth float64  `json:"exp_month"`
            CardReference string  `json:"card_reference"`
            CardIsin string  `json:"card_isin"`
            CardIssuer string  `json:"card_issuer"`
            CardType string  `json:"card_type"`
            IntentFlow bool  `json:"intent_flow"`
            IntentAppErrorDictList []IntentAppErrorList  `json:"intent_app_error_dict_list"`
            RemainingLimit float64  `json:"remaining_limit"`
            CodLimitPerOrder float64  `json:"cod_limit_per_order"`
            CardID string  `json:"card_id"`
            Name string  `json:"name"`
            Code string  `json:"code"`
            IntentApp []IntentApp  `json:"intent_app"`
            CardBrand string  `json:"card_brand"`
            CardNumber string  `json:"card_number"`
            CompliantWithTokenisationGuidelines bool  `json:"compliant_with_tokenisation_guidelines"`
            ExpYear float64  `json:"exp_year"`
            CardToken string  `json:"card_token"`
            CardName string  `json:"card_name"`
         
    }
    
    // RootPaymentMode used by Payment
    type RootPaymentMode struct {

        
            List []PaymentModeList  `json:"list"`
            SaveCard bool  `json:"save_card"`
            AnonymousEnable bool  `json:"anonymous_enable"`
            DisplayName string  `json:"display_name"`
            Name string  `json:"name"`
            AddCardEnabled bool  `json:"add_card_enabled"`
            AggregatorName string  `json:"aggregator_name"`
            DisplayPriority float64  `json:"display_priority"`
            IsPayByCardPl bool  `json:"is_pay_by_card_pl"`
         
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

        
            AggregatorFundID float64  `json:"aggregator_fund_id"`
            PayoutDetailsID float64  `json:"payout_details_id"`
            AggregatorID float64  `json:"aggregator_id"`
         
    }
    
    // PayoutMoreAttributes used by Payment
    type PayoutMoreAttributes struct {

        
            AccountHolder string  `json:"account_holder"`
            AccountType string  `json:"account_type"`
            City string  `json:"city"`
            Country string  `json:"country"`
            AccountNo string  `json:"account_no"`
            BankName string  `json:"bank_name"`
            BranchName string  `json:"branch_name"`
            State string  `json:"state"`
            IfscCode string  `json:"ifsc_code"`
         
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

        
            TransferType string  `json:"transfer_type"`
            UniqueTransferNo string  `json:"unique_transfer_no"`
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
            AccountType string  `json:"account_type"`
            Pincode float64  `json:"pincode"`
            City string  `json:"city"`
            BranchName string  `json:"branch_name"`
            AccountNo string  `json:"account_no"`
            BankName string  `json:"bank_name"`
            Country string  `json:"country"`
            State string  `json:"state"`
            IfscCode string  `json:"ifsc_code"`
         
    }
    
    // PayoutRequest used by Payment
    type PayoutRequest struct {

        
            BankDetails PayoutBankDetails  `json:"bank_details"`
            TransferType string  `json:"transfer_type"`
            Users map[string]interface{}  `json:"users"`
            IsActive bool  `json:"is_active"`
            UniqueExternalID string  `json:"unique_external_id"`
            Aggregator string  `json:"aggregator"`
         
    }
    
    // PayoutResponse used by Payment
    type PayoutResponse struct {

        
            BankDetails map[string]interface{}  `json:"bank_details"`
            UniqueTransferNo string  `json:"unique_transfer_no"`
            TransferType string  `json:"transfer_type"`
            Created bool  `json:"created"`
            Payouts map[string]interface{}  `json:"payouts"`
            PaymentStatus string  `json:"payment_status"`
            Users map[string]interface{}  `json:"users"`
            IsActive bool  `json:"is_active"`
            Success bool  `json:"success"`
            Aggregator string  `json:"aggregator"`
         
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

        
            Success bool  `json:"success"`
            Data map[string]interface{}  `json:"data"`
         
    }
    
    // RefundAccountResponse used by Payment
    type RefundAccountResponse struct {

        
            IsVerifiedFlag bool  `json:"is_verified_flag"`
            Data map[string]interface{}  `json:"data"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // NotFoundResourceError used by Payment
    type NotFoundResourceError struct {

        
            Code string  `json:"code"`
            Description string  `json:"description"`
            Success bool  `json:"success"`
         
    }
    
    // BankDetailsForOTP used by Payment
    type BankDetailsForOTP struct {

        
            AccountHolder string  `json:"account_holder"`
            BranchName string  `json:"branch_name"`
            AccountNo string  `json:"account_no"`
            BankName string  `json:"bank_name"`
            IfscCode string  `json:"ifsc_code"`
         
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

        
            AccountHolder string  `json:"account_holder"`
            BranchName string  `json:"branch_name"`
            BeneficiaryID string  `json:"beneficiary_id"`
            Email string  `json:"email"`
            IfscCode string  `json:"ifsc_code"`
            AccountNo string  `json:"account_no"`
            DisplayName string  `json:"display_name"`
            IsActive bool  `json:"is_active"`
            ID float64  `json:"id"`
            DelightsUserName string  `json:"delights_user_name"`
            Address string  `json:"address"`
            CreatedOn string  `json:"created_on"`
            Mobile string  `json:"mobile"`
            Subtitle string  `json:"subtitle"`
            Title string  `json:"title"`
            ModifiedOn string  `json:"modified_on"`
            Comment string  `json:"comment"`
            BankName string  `json:"bank_name"`
            TransferMode string  `json:"transfer_mode"`
         
    }
    
    // OrderBeneficiaryResponse used by Payment
    type OrderBeneficiaryResponse struct {

        
            Beneficiaries []OrderBeneficiaryDetails  `json:"beneficiaries"`
            ShowBeneficiaryDetails bool  `json:"show_beneficiary_details"`
         
    }
    
    // MultiTenderPaymentMeta used by Payment
    type MultiTenderPaymentMeta struct {

        
            CurrentStatus string  `json:"current_status"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            OrderID string  `json:"order_id"`
            PaymentGateway string  `json:"payment_gateway"`
            PaymentID string  `json:"payment_id"`
         
    }
    
    // MultiTenderPaymentMethod used by Payment
    type MultiTenderPaymentMethod struct {

        
            Mode string  `json:"mode"`
            Name string  `json:"name"`
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
    
    // PlatformPaymentOptions used by Payment
    type PlatformPaymentOptions struct {

        
            CodCharges float64  `json:"cod_charges"`
            AnonymousCod bool  `json:"anonymous_cod"`
            Enabled bool  `json:"enabled"`
            Source string  `json:"source"`
            Methods map[string]interface{}  `json:"methods"`
            ModeOfPayment string  `json:"mode_of_payment"`
            CodAmountLimit float64  `json:"cod_amount_limit"`
            PaymentSelectionLock map[string]interface{}  `json:"payment_selection_lock"`
            CallbackURL map[string]interface{}  `json:"callback_url"`
         
    }
    
    // PlatfromPaymentConfig used by Payment
    type PlatfromPaymentConfig struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
            Data PlatformPaymentOptions  `json:"data"`
         
    }
    
    // UpdatePlatformPaymentConfig used by Payment
    type UpdatePlatformPaymentConfig struct {

        
            CodCharges float64  `json:"cod_charges"`
            AnonymousCod bool  `json:"anonymous_cod"`
            Methods map[string]interface{}  `json:"methods"`
            CodAmountLimit float64  `json:"cod_amount_limit"`
            PaymentSelectionLock map[string]interface{}  `json:"payment_selection_lock"`
         
    }
    
    // CODdata used by Payment
    type CODdata struct {

        
            UserID string  `json:"user_id"`
            Limit float64  `json:"limit"`
            IsActive bool  `json:"is_active"`
            Usages float64  `json:"usages"`
            RemainingLimit float64  `json:"remaining_limit"`
         
    }
    
    // GetUserCODLimitResponse used by Payment
    type GetUserCODLimitResponse struct {

        
            UserCodData CODdata  `json:"user_cod_data"`
            Success bool  `json:"success"`
         
    }
    
    // SetCODForUserRequest used by Payment
    type SetCODForUserRequest struct {

        
            MerchantUserID string  `json:"merchant_user_id"`
            Mobileno string  `json:"mobileno"`
            IsActive bool  `json:"is_active"`
         
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

        
            EdcModel string  `json:"edc_model"`
            AggregatorID float64  `json:"aggregator_id"`
            DeviceTag string  `json:"device_tag"`
            EdcDeviceSerialNo string  `json:"edc_device_serial_no"`
            StoreID float64  `json:"store_id"`
            TerminalSerialNo string  `json:"terminal_serial_no"`
         
    }
    
    // EdcDevice used by Payment
    type EdcDevice struct {

        
            ApplicationID string  `json:"application_id"`
            EdcModel string  `json:"edc_model"`
            AggregatorID float64  `json:"aggregator_id"`
            DeviceTag string  `json:"device_tag"`
            IsActive bool  `json:"is_active"`
            TerminalUniqueIdentifier string  `json:"terminal_unique_identifier"`
            MerchantStorePosCode string  `json:"merchant_store_pos_code"`
            EdcDeviceSerialNo string  `json:"edc_device_serial_no"`
            StoreID float64  `json:"store_id"`
            TerminalSerialNo string  `json:"terminal_serial_no"`
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

        
            EdcModel string  `json:"edc_model"`
            AggregatorID float64  `json:"aggregator_id"`
            DeviceTag string  `json:"device_tag"`
            IsActive bool  `json:"is_active"`
            MerchantStorePosCode string  `json:"merchant_store_pos_code"`
            EdcDeviceSerialNo string  `json:"edc_device_serial_no"`
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

        
            Items []EdcDevice  `json:"items"`
            Success bool  `json:"success"`
            Page Page  `json:"page"`
         
    }
    
    // PaymentInitializationRequest used by Payment
    type PaymentInitializationRequest struct {

        
            RazorpayPaymentID string  `json:"razorpay_payment_id"`
            Method string  `json:"method"`
            DeviceID string  `json:"device_id"`
            Timeout float64  `json:"timeout"`
            Vpa string  `json:"vpa"`
            Currency string  `json:"currency"`
            MerchantOrderID string  `json:"merchant_order_id"`
            Contact string  `json:"contact"`
            Amount float64  `json:"amount"`
            OrderID string  `json:"order_id"`
            CustomerID string  `json:"customer_id"`
            Email string  `json:"email"`
            Aggregator string  `json:"aggregator"`
         
    }
    
    // PaymentInitializationResponse used by Payment
    type PaymentInitializationResponse struct {

        
            RazorpayPaymentID string  `json:"razorpay_payment_id"`
            DeviceID string  `json:"device_id"`
            Currency string  `json:"currency"`
            AggregatorOrderID string  `json:"aggregator_order_id"`
            OrderID string  `json:"order_id"`
            PaymentID string  `json:"payment_id"`
            Email string  `json:"email"`
            Aggregator string  `json:"aggregator"`
            Timeout float64  `json:"timeout"`
            PollingURL string  `json:"polling_url"`
            Contact string  `json:"contact"`
            Status string  `json:"status"`
            VirtualID string  `json:"virtual_id"`
            MerchantOrderID string  `json:"merchant_order_id"`
            BqrImage string  `json:"bqr_image"`
            Amount float64  `json:"amount"`
            Method string  `json:"method"`
            Vpa string  `json:"vpa"`
            Success bool  `json:"success"`
            CustomerID string  `json:"customer_id"`
            UpiPollURL string  `json:"upi_poll_url"`
         
    }
    
    // PaymentStatusUpdateRequest used by Payment
    type PaymentStatusUpdateRequest struct {

        
            RazorpayPaymentID string  `json:"razorpay_payment_id"`
            Method string  `json:"method"`
            DeviceID string  `json:"device_id"`
            Email string  `json:"email"`
            Vpa string  `json:"vpa"`
            Currency string  `json:"currency"`
            CustomerID string  `json:"customer_id"`
            MerchantOrderID string  `json:"merchant_order_id"`
            Contact string  `json:"contact"`
            OrderID string  `json:"order_id"`
            PaymentID string  `json:"payment_id"`
            Status string  `json:"status"`
            Amount float64  `json:"amount"`
            Aggregator string  `json:"aggregator"`
         
    }
    
    // PaymentStatusUpdateResponse used by Payment
    type PaymentStatusUpdateResponse struct {

        
            Retry bool  `json:"retry"`
            Success bool  `json:"success"`
            AggregatorName string  `json:"aggregator_name"`
            RedirectURL string  `json:"redirect_url"`
            Status string  `json:"status"`
         
    }
    
    // ResendOrCancelPaymentRequest used by Payment
    type ResendOrCancelPaymentRequest struct {

        
            DeviceID string  `json:"device_id"`
            RequestType string  `json:"request_type"`
            OrderID string  `json:"order_id"`
         
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
    

    
    // ShipmentStatus used by Order
    type ShipmentStatus struct {

        
            HexCode string  `json:"hex_code"`
            OpsStatus string  `json:"ops_status"`
            Title string  `json:"title"`
            Status string  `json:"status"`
            ActualStatus string  `json:"actual_status"`
         
    }
    
    // ShipmentItemFulFillingStore used by Order
    type ShipmentItemFulFillingStore struct {

        
            ID string  `json:"id"`
            Code string  `json:"code"`
         
    }
    
    // UserDataInfo used by Order
    type UserDataInfo struct {

        
            UID float64  `json:"uid"`
            Email string  `json:"email"`
            Gender string  `json:"gender"`
            FirstName string  `json:"first_name"`
            Name string  `json:"name"`
            AvisUserID string  `json:"avis_user_id"`
            IsAnonymousUser bool  `json:"is_anonymous_user"`
            LastName string  `json:"last_name"`
            Mobile string  `json:"mobile"`
         
    }
    
    // Prices used by Order
    type Prices struct {

        
            ValueOfGood float64  `json:"value_of_good"`
            AmountPaidRoundoff float64  `json:"amount_paid_roundoff"`
            FyndCredits float64  `json:"fynd_credits"`
            Cashback float64  `json:"cashback"`
            PromotionEffectiveDiscount float64  `json:"promotion_effective_discount"`
            RefundCredit float64  `json:"refund_credit"`
            TransferPrice float64  `json:"transfer_price"`
            PriceMarked float64  `json:"price_marked"`
            PriceEffective float64  `json:"price_effective"`
            DeliveryCharge float64  `json:"delivery_charge"`
            AmountPaid float64  `json:"amount_paid"`
            Discount float64  `json:"discount"`
            RefundAmount float64  `json:"refund_amount"`
            CouponValue float64  `json:"coupon_value"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            CashbackApplied float64  `json:"cashback_applied"`
            CodCharges float64  `json:"cod_charges"`
         
    }
    
    // PlatformItem used by Order
    type PlatformItem struct {

        
            CanCancel bool  `json:"can_cancel"`
            Code string  `json:"code"`
            DepartmentID float64  `json:"department_id"`
            Images []string  `json:"images"`
            L3CategoryName string  `json:"l3_category_name"`
            Image []string  `json:"image"`
            Name string  `json:"name"`
            CanReturn bool  `json:"can_return"`
            Size string  `json:"size"`
            L3Category float64  `json:"l3_category"`
            L1Category []string  `json:"l1_category"`
            ID float64  `json:"id"`
            Color string  `json:"color"`
         
    }
    
    // GSTDetailsData used by Order
    type GSTDetailsData struct {

        
            ValueOfGood float64  `json:"value_of_good"`
            GstFee float64  `json:"gst_fee"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            GstinCode string  `json:"gstin_code"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
         
    }
    
    // BagUnit used by Order
    type BagUnit struct {

        
            CanCancel bool  `json:"can_cancel"`
            Item PlatformItem  `json:"item"`
            ShipmentID string  `json:"shipment_id"`
            TotalShipmentBags float64  `json:"total_shipment_bags"`
            Prices Prices  `json:"prices"`
            OrderingChannel string  `json:"ordering_channel"`
            CanReturn bool  `json:"can_return"`
            Status map[string]interface{}  `json:"status"`
            Gst GSTDetailsData  `json:"gst"`
            BagID float64  `json:"bag_id"`
            ItemQuantity float64  `json:"item_quantity"`
         
    }
    
    // PaymentModeInfo used by Order
    type PaymentModeInfo struct {

        
            Logo string  `json:"logo"`
            Type string  `json:"type"`
         
    }
    
    // ShipmentItem used by Order
    type ShipmentItem struct {

        
            ShipmentStatus ShipmentStatus  `json:"shipment_status"`
            FulfillingStore ShipmentItemFulFillingStore  `json:"fulfilling_store"`
            User UserDataInfo  `json:"user"`
            ShipmentID string  `json:"shipment_id"`
            Prices Prices  `json:"prices"`
            ShipmentCreatedAt string  `json:"shipment_created_at"`
            Bags []BagUnit  `json:"bags"`
            Sla map[string]interface{}  `json:"sla"`
            TotalShipmentsInOrder float64  `json:"total_shipments_in_order"`
            PaymentMethods map[string]interface{}  `json:"payment_methods"`
            Channel map[string]interface{}  `json:"channel"`
            CreatedAt string  `json:"created_at"`
            FulfillingCentre string  `json:"fulfilling_centre"`
            Application map[string]interface{}  `json:"application"`
            PaymentModeInfo PaymentModeInfo  `json:"payment_mode_info"`
            TotalBagsCount float64  `json:"total_bags_count"`
            Company map[string]interface{}  `json:"company"`
            ID string  `json:"id"`
         
    }
    
    // FilterInfoOption used by Order
    type FilterInfoOption struct {

        
            Text string  `json:"text"`
            Name string  `json:"name"`
            Value string  `json:"value"`
         
    }
    
    // FiltersInfo used by Order
    type FiltersInfo struct {

        
            Text string  `json:"text"`
            Options []FilterInfoOption  `json:"options"`
            Value string  `json:"value"`
            Type string  `json:"type"`
         
    }
    
    // ShipmentInternalPlatformViewResponse used by Order
    type ShipmentInternalPlatformViewResponse struct {

        
            Items []ShipmentItem  `json:"items"`
            Filters []FiltersInfo  `json:"filters"`
            Page map[string]interface{}  `json:"page"`
            AppliedFilters map[string]interface{}  `json:"applied_filters"`
         
    }
    
    // Error used by Order
    type Error struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // BagStateMapper used by Order
    type BagStateMapper struct {

        
            DisplayName string  `json:"display_name"`
            JourneyType string  `json:"journey_type"`
            AppDisplayName string  `json:"app_display_name"`
            StateType string  `json:"state_type"`
            NotifyCustomer bool  `json:"notify_customer"`
            Name string  `json:"name"`
            IsActive bool  `json:"is_active"`
            BsID float64  `json:"bs_id"`
            AppStateName string  `json:"app_state_name"`
            AppFacing bool  `json:"app_facing"`
         
    }
    
    // BagStatusHistory used by Order
    type BagStatusHistory struct {

        
            StoreID float64  `json:"store_id"`
            Reasons []map[string]interface{}  `json:"reasons"`
            Forward bool  `json:"forward"`
            ShipmentID string  `json:"shipment_id"`
            DisplayName string  `json:"display_name"`
            DeliveryAwbNumber string  `json:"delivery_awb_number"`
            AppDisplayName string  `json:"app_display_name"`
            KafkaSync bool  `json:"kafka_sync"`
            StateType string  `json:"state_type"`
            CreatedAt string  `json:"created_at"`
            UpdatedAt string  `json:"updated_at"`
            Status string  `json:"status"`
            DeliveryPartnerID float64  `json:"delivery_partner_id"`
            StateID float64  `json:"state_id"`
            BagID float64  `json:"bag_id"`
            BagStateMapper BagStateMapper  `json:"bag_state_mapper"`
            BshID float64  `json:"bsh_id"`
         
    }
    
    // Dimensions used by Order
    type Dimensions struct {

        
            Unit string  `json:"unit"`
            Width float64  `json:"width"`
            Length float64  `json:"length"`
            Height float64  `json:"height"`
            IsDefault bool  `json:"is_default"`
         
    }
    
    // Meta used by Order
    type Meta struct {

        
            Dimension Dimensions  `json:"dimension"`
         
    }
    
    // OrderDetailsData used by Order
    type OrderDetailsData struct {

        
            AffiliateID string  `json:"affiliate_id"`
            OrderingChannelLogo map[string]interface{}  `json:"ordering_channel_logo"`
            Source string  `json:"source"`
            FyndOrderID string  `json:"fynd_order_id"`
            OrderValue string  `json:"order_value"`
            OrderingChannel string  `json:"ordering_channel"`
            OrderDate string  `json:"order_date"`
            CodCharges string  `json:"cod_charges"`
            TaxDetails map[string]interface{}  `json:"tax_details"`
         
    }
    
    // AffiliateMeta used by Order
    type AffiliateMeta struct {

        
            ChannelOrderID string  `json:"channel_order_id"`
            BoxType string  `json:"box_type"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            SizeLevelTotalQty float64  `json:"size_level_total_qty"`
            DueDate string  `json:"due_date"`
            OrderItemID string  `json:"order_item_id"`
            IsPriority bool  `json:"is_priority"`
            ChannelShipmentID string  `json:"channel_shipment_id"`
            Quantity float64  `json:"quantity"`
            EmployeeDiscount float64  `json:"employee_discount"`
            MarketplaceInvoiceID string  `json:"marketplace_invoice_id"`
            CouponCode string  `json:"coupon_code"`
         
    }
    
    // PDFLinks used by Order
    type PDFLinks struct {

        
            Label string  `json:"label"`
            LabelType string  `json:"label_type"`
            InvoicePos string  `json:"invoice_pos"`
            LabelPos string  `json:"label_pos"`
            LabelA4 string  `json:"label_a4"`
            CreditNoteURL string  `json:"credit_note_url"`
            B2b string  `json:"b2b"`
            InvoiceType string  `json:"invoice_type"`
            PoInvoice string  `json:"po_invoice"`
            Invoice string  `json:"invoice"`
            DeliveryChallanA4 string  `json:"delivery_challan_a4"`
            LabelA6 string  `json:"label_a6"`
            InvoiceA4 string  `json:"invoice_a4"`
            InvoiceA6 string  `json:"invoice_a6"`
         
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
    
    // LockData used by Order
    type LockData struct {

        
            Mto bool  `json:"mto"`
            LockMessage string  `json:"lock_message"`
            Locked bool  `json:"locked"`
         
    }
    
    // ShipmentTimeStamp used by Order
    type ShipmentTimeStamp struct {

        
            TMax string  `json:"t_max"`
            TMin string  `json:"t_min"`
         
    }
    
    // BuyerDetails used by Order
    type BuyerDetails struct {

        
            City string  `json:"city"`
            Name string  `json:"name"`
            Address string  `json:"address"`
            Gstin string  `json:"gstin"`
            AjioSiteID string  `json:"ajio_site_id"`
            State string  `json:"state"`
            Pincode float64  `json:"pincode"`
         
    }
    
    // Formatted used by Order
    type Formatted struct {

        
            FMax string  `json:"f_max"`
            FMin string  `json:"f_min"`
         
    }
    
    // ShipmentMeta used by Order
    type ShipmentMeta struct {

        
            AutoTriggerDpAssignmentAcf bool  `json:"auto_trigger_dp_assignment_acf"`
            BoxType string  `json:"box_type"`
            EinvoiceInfo EinvoiceInfo  `json:"einvoice_info"`
            DpOptions map[string]interface{}  `json:"dp_options"`
            ReturnDetails map[string]interface{}  `json:"return_details"`
            MarketplaceStoreID string  `json:"marketplace_store_id"`
            ForwardAffiliateOrderID string  `json:"forward_affiliate_order_id"`
            OrderType string  `json:"order_type"`
            DebugInfo DebugInfo  `json:"debug_info"`
            ShipmentWeight float64  `json:"shipment_weight"`
            BagWeight map[string]interface{}  `json:"bag_weight"`
            LockData LockData  `json:"lock_data"`
            PoNumber string  `json:"po_number"`
            Weight float64  `json:"weight"`
            DpSortKey string  `json:"dp_sort_key"`
            B2cBuyerDetails map[string]interface{}  `json:"b2c_buyer_details"`
            Timestamp ShipmentTimeStamp  `json:"timestamp"`
            DpName string  `json:"dp_name"`
            ShipmentVolumetricWeight float64  `json:"shipment_volumetric_weight"`
            StoreInvoiceUpdatedDate string  `json:"store_invoice_updated_date"`
            AwbNumber string  `json:"awb_number"`
            FulfilmentPriorityText string  `json:"fulfilment_priority_text"`
            ReturnAffiliateOrderID string  `json:"return_affiliate_order_id"`
            DpID string  `json:"dp_id"`
            EwaybillInfo map[string]interface{}  `json:"ewaybill_info"`
            SameStoreAvailable bool  `json:"same_store_available"`
            DueDate string  `json:"due_date"`
            ReturnStoreNode float64  `json:"return_store_node"`
            B2bBuyerDetails BuyerDetails  `json:"b2b_buyer_details"`
            ReturnAwbNumber string  `json:"return_awb_number"`
            AssignDpFromSb bool  `json:"assign_dp_from_sb"`
            External map[string]interface{}  `json:"external"`
            PackagingName string  `json:"packaging_name"`
            ForwardAffiliateShipmentID string  `json:"forward_affiliate_shipment_id"`
            ReturnAffiliateShipmentID string  `json:"return_affiliate_shipment_id"`
            Formatted Formatted  `json:"formatted"`
         
    }
    
    // AffiliateDetails used by Order
    type AffiliateDetails struct {

        
            AffiliateID string  `json:"affiliate_id"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
            AdID string  `json:"ad_id"`
            CompanyAffiliateTag string  `json:"company_affiliate_tag"`
            AffiliateMeta AffiliateMeta  `json:"affiliate_meta"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            PdfLinks PDFLinks  `json:"pdf_links"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            ShipmentMeta ShipmentMeta  `json:"shipment_meta"`
            AffiliateStoreID string  `json:"affiliate_store_id"`
         
    }
    
    // OrderingStoreDetails used by Order
    type OrderingStoreDetails struct {

        
            Phone string  `json:"phone"`
            City string  `json:"city"`
            Code string  `json:"code"`
            Meta map[string]interface{}  `json:"meta"`
            ContactPerson string  `json:"contact_person"`
            OrderingStoreID float64  `json:"ordering_store_id"`
            Address string  `json:"address"`
            State string  `json:"state"`
            Country string  `json:"country"`
            StoreName string  `json:"store_name"`
            Pincode string  `json:"pincode"`
         
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

        
            PromotionType string  `json:"promotion_type"`
            ArticleQuantity float64  `json:"article_quantity"`
            PromoID string  `json:"promo_id"`
            DiscountRules []DiscountRules  `json:"discount_rules"`
            PromotionName string  `json:"promotion_name"`
            MrpPromotion bool  `json:"mrp_promotion"`
            BuyRules []BuyRules  `json:"buy_rules"`
            Amount float64  `json:"amount"`
         
    }
    
    // GiftCard used by Order
    type GiftCard struct {

        
            IsGiftApplied bool  `json:"is_gift_applied"`
            DisplayText string  `json:"display_text"`
            GiftMessage string  `json:"gift_message"`
            GiftPrice float64  `json:"gift_price"`
         
    }
    
    // B2BPODetails used by Order
    type B2BPODetails struct {

        
            PartialCanRet bool  `json:"partial_can_ret"`
            TotalGstPercentage float64  `json:"total_gst_percentage"`
            PoTaxAmount float64  `json:"po_tax_amount"`
            PoLineAmount float64  `json:"po_line_amount"`
            DockerNumber string  `json:"docker_number"`
            ItemBasePrice float64  `json:"item_base_price"`
         
    }
    
    // BagMeta used by Order
    type BagMeta struct {

        
            GiftCard GiftCard  `json:"gift_card"`
            CustomMessage string  `json:"custom_message"`
            DocketNumber string  `json:"docket_number"`
            B2bPoDetails B2BPODetails  `json:"b2b_po_details"`
            PartialCanRet bool  `json:"partial_can_ret"`
            GroupID string  `json:"group_id"`
            CustomJson map[string]interface{}  `json:"custom_json"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
         
    }
    
    // Identifier used by Order
    type Identifier struct {

        
            Ean string  `json:"ean"`
         
    }
    
    // FinancialBreakup used by Order
    type FinancialBreakup struct {

        
            PromotionEffectiveDiscount float64  `json:"promotion_effective_discount"`
            RefundCredit float64  `json:"refund_credit"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            HsnCode string  `json:"hsn_code"`
            TotalUnits float64  `json:"total_units"`
            FyndCredits float64  `json:"fynd_credits"`
            TransferPrice float64  `json:"transfer_price"`
            Identifiers Identifier  `json:"identifiers"`
            DeliveryCharge float64  `json:"delivery_charge"`
            AmountPaid float64  `json:"amount_paid"`
            CouponValue float64  `json:"coupon_value"`
            ItemName string  `json:"item_name"`
            PriceEffective float64  `json:"price_effective"`
            PriceMarked float64  `json:"price_marked"`
            AmountPaidRoundoff float64  `json:"amount_paid_roundoff"`
            ValueOfGood float64  `json:"value_of_good"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            GstTag string  `json:"gst_tag"`
            GstFee float64  `json:"gst_fee"`
            AddedToFyndCash bool  `json:"added_to_fynd_cash"`
            CodCharges float64  `json:"cod_charges"`
            Discount float64  `json:"discount"`
            Size string  `json:"size"`
            CashbackApplied float64  `json:"cashback_applied"`
            Cashback float64  `json:"cashback"`
         
    }
    
    // BagConfigs used by Order
    type BagConfigs struct {

        
            IsReturnable bool  `json:"is_returnable"`
            IsCustomerReturnAllowed bool  `json:"is_customer_return_allowed"`
            CanBeCancelled bool  `json:"can_be_cancelled"`
            IsActive bool  `json:"is_active"`
            AllowForceReturn bool  `json:"allow_force_return"`
            EnableTracking bool  `json:"enable_tracking"`
         
    }
    
    // BagGST used by Order
    type BagGST struct {

        
            ValueOfGood float64  `json:"value_of_good"`
            GstTag string  `json:"gst_tag"`
            GstFee float64  `json:"gst_fee"`
            GstinCode string  `json:"gstin_code"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            IsDefaultHsnCode bool  `json:"is_default_hsn_code"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            HsnCode string  `json:"hsn_code"`
         
    }
    
    // OrderBrandName used by Order
    type OrderBrandName struct {

        
            Logo string  `json:"logo"`
            ModifiedOn string  `json:"modified_on"`
            ID float64  `json:"id"`
            BrandName string  `json:"brand_name"`
            Company string  `json:"company"`
            CreatedOn string  `json:"created_on"`
         
    }
    
    // PlatformDeliveryAddress used by Order
    type PlatformDeliveryAddress struct {

        
            Phone string  `json:"phone"`
            AddressCategory string  `json:"address_category"`
            City string  `json:"city"`
            Version string  `json:"version"`
            ContactPerson string  `json:"contact_person"`
            Address1 string  `json:"address1"`
            Email string  `json:"email"`
            Address2 string  `json:"address2"`
            Longitude float64  `json:"longitude"`
            Latitude float64  `json:"latitude"`
            CreatedAt string  `json:"created_at"`
            Landmark string  `json:"landmark"`
            UpdatedAt string  `json:"updated_at"`
            Country string  `json:"country"`
            AddressType string  `json:"address_type"`
            State string  `json:"state"`
            Area string  `json:"area"`
            Pincode string  `json:"pincode"`
         
    }
    
    // CurrentStatus used by Order
    type CurrentStatus struct {

        
            StoreID float64  `json:"store_id"`
            ShipmentID string  `json:"shipment_id"`
            DeliveryAwbNumber string  `json:"delivery_awb_number"`
            KafkaSync bool  `json:"kafka_sync"`
            StateType string  `json:"state_type"`
            CurrentStatusID float64  `json:"current_status_id"`
            CreatedAt string  `json:"created_at"`
            UpdatedAt string  `json:"updated_at"`
            Status string  `json:"status"`
            DeliveryPartnerID float64  `json:"delivery_partner_id"`
            StateID float64  `json:"state_id"`
            BagID float64  `json:"bag_id"`
            BagStateMapper BagStateMapper  `json:"bag_state_mapper"`
         
    }
    
    // OrderBagArticle used by Order
    type OrderBagArticle struct {

        
            UID string  `json:"uid"`
            Identifiers map[string]interface{}  `json:"identifiers"`
            ReturnConfig map[string]interface{}  `json:"return_config"`
         
    }
    
    // OrderBags used by Order
    type OrderBags struct {

        
            AppliedPromos []AppliedPromos  `json:"applied_promos"`
            Meta BagMeta  `json:"meta"`
            GroupID string  `json:"group_id"`
            EntityType string  `json:"entity_type"`
            Quantity float64  `json:"quantity"`
            CanCancel bool  `json:"can_cancel"`
            FinancialBreakup FinancialBreakup  `json:"financial_breakup"`
            CanReturn bool  `json:"can_return"`
            BagConfigs BagConfigs  `json:"bag_configs"`
            ParentPromoBags map[string]interface{}  `json:"parent_promo_bags"`
            Item PlatformItem  `json:"item"`
            Prices Prices  `json:"prices"`
            DisplayName string  `json:"display_name"`
            GstDetails BagGST  `json:"gst_details"`
            Identifier string  `json:"identifier"`
            LineNumber float64  `json:"line_number"`
            Brand OrderBrandName  `json:"brand"`
            BagID float64  `json:"bag_id"`
            SellerIdentifier string  `json:"seller_identifier"`
            DeliveryAddress PlatformDeliveryAddress  `json:"delivery_address"`
            CurrentStatus CurrentStatus  `json:"current_status"`
            Article OrderBagArticle  `json:"article"`
         
    }
    
    // UserDetailsData used by Order
    type UserDetailsData struct {

        
            Phone string  `json:"phone"`
            City string  `json:"city"`
            Address1 string  `json:"address1"`
            Email string  `json:"email"`
            Landmark string  `json:"landmark"`
            Name string  `json:"name"`
            Address string  `json:"address"`
            Country string  `json:"country"`
            AddressType string  `json:"address_type"`
            State string  `json:"state"`
            Area string  `json:"area"`
            Pincode string  `json:"pincode"`
         
    }
    
    // ShipmentPayments used by Order
    type ShipmentPayments struct {

        
            Logo string  `json:"logo"`
            Source string  `json:"source"`
            Mode string  `json:"mode"`
         
    }
    
    // ShipmentStatusData used by Order
    type ShipmentStatusData struct {

        
            BagList []string  `json:"bag_list"`
            ShipmentID string  `json:"shipment_id"`
            CreatedAt string  `json:"created_at"`
            Status string  `json:"status"`
            ID float64  `json:"id"`
         
    }
    
    // FulfillingStore used by Order
    type FulfillingStore struct {

        
            Phone string  `json:"phone"`
            City string  `json:"city"`
            Code string  `json:"code"`
            Meta map[string]interface{}  `json:"meta"`
            ContactPerson string  `json:"contact_person"`
            FulfillmentChannel string  `json:"fulfillment_channel"`
            Address string  `json:"address"`
            State string  `json:"state"`
            Country string  `json:"country"`
            ID float64  `json:"id"`
            StoreName string  `json:"store_name"`
            Pincode string  `json:"pincode"`
         
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

        
            CompanyName string  `json:"company_name"`
            CompanyGst string  `json:"company_gst"`
            CompanyContact ContactDetails  `json:"company_contact"`
            CompanyCin string  `json:"company_cin"`
            Address map[string]interface{}  `json:"address"`
            CompanyID float64  `json:"company_id"`
         
    }
    
    // DPDetailsData used by Order
    type DPDetailsData struct {

        
            TrackURL string  `json:"track_url"`
            GstTag string  `json:"gst_tag"`
            AwbNo string  `json:"awb_no"`
            EwayBillID string  `json:"eway_bill_id"`
            Name string  `json:"name"`
            Country string  `json:"country"`
            ID float64  `json:"id"`
            Pincode string  `json:"pincode"`
         
    }
    
    // TrackingList used by Order
    type TrackingList struct {

        
            Text string  `json:"text"`
            IsPassed bool  `json:"is_passed"`
            Time string  `json:"time"`
            IsCurrent bool  `json:"is_current"`
            Status string  `json:"status"`
         
    }
    
    // InvoiceInfo used by Order
    type InvoiceInfo struct {

        
            StoreInvoiceID string  `json:"store_invoice_id"`
            InvoiceURL string  `json:"invoice_url"`
            CreditNoteID string  `json:"credit_note_id"`
            LabelURL string  `json:"label_url"`
            UpdatedDate string  `json:"updated_date"`
         
    }
    
    // PlatformShipment used by Order
    type PlatformShipment struct {

        
            BagStatusHistory []BagStatusHistory  `json:"bag_status_history"`
            ShipmentImages []string  `json:"shipment_images"`
            DeliverySlot map[string]interface{}  `json:"delivery_slot"`
            ShipmentID string  `json:"shipment_id"`
            Meta Meta  `json:"meta"`
            Order OrderDetailsData  `json:"order"`
            AffiliateDetails AffiliateDetails  `json:"affiliate_details"`
            Vertical string  `json:"vertical"`
            ShipmentQuantity float64  `json:"shipment_quantity"`
            PlatformLogo string  `json:"platform_logo"`
            Coupon map[string]interface{}  `json:"coupon"`
            ForwardShipmentID string  `json:"forward_shipment_id"`
            OperationalStatus string  `json:"operational_status"`
            ShipmentStatus string  `json:"shipment_status"`
            OrderingStore OrderingStoreDetails  `json:"ordering_store"`
            User UserDataInfo  `json:"user"`
            Bags []OrderBags  `json:"bags"`
            TotalItems float64  `json:"total_items"`
            LockStatus bool  `json:"lock_status"`
            PickedDate string  `json:"picked_date"`
            CustomMeta []map[string]interface{}  `json:"custom_meta"`
            PackagingType string  `json:"packaging_type"`
            BillingDetails UserDetailsData  `json:"billing_details"`
            EnableDpTracking bool  `json:"enable_dp_tracking"`
            UserAgent string  `json:"user_agent"`
            Prices Prices  `json:"prices"`
            GstDetails GSTDetailsData  `json:"gst_details"`
            PriorityText string  `json:"priority_text"`
            Payments ShipmentPayments  `json:"payments"`
            Status ShipmentStatusData  `json:"status"`
            PaymentMethods map[string]interface{}  `json:"payment_methods"`
            FulfillingStore FulfillingStore  `json:"fulfilling_store"`
            CompanyDetails CompanyDetails  `json:"company_details"`
            JourneyType string  `json:"journey_type"`
            DpDetails DPDetailsData  `json:"dp_details"`
            PaymentMode string  `json:"payment_mode"`
            FulfilmentPriority float64  `json:"fulfilment_priority"`
            DeliveryDetails UserDetailsData  `json:"delivery_details"`
            TrackingList []TrackingList  `json:"tracking_list"`
            TotalBags float64  `json:"total_bags"`
            Invoice InvoiceInfo  `json:"invoice"`
            InvoiceID string  `json:"invoice_id"`
         
    }
    
    // ShipmentInfoResponse used by Order
    type ShipmentInfoResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
            Shipments []PlatformShipment  `json:"shipments"`
         
    }
    
    // AssetByShipment used by Order
    type AssetByShipment struct {

        
            ExpiresIn string  `json:"expires_in"`
            Assets map[string]string  `json:"assets"`
            ShipmentID string  `json:"shipment_id"`
         
    }
    
    // ResponseGetAssetShipment used by Order
    type ResponseGetAssetShipment struct {

        
            PresignedType string  `json:"presigned_type"`
            Success bool  `json:"success"`
            Result []AssetByShipment  `json:"result"`
         
    }
    
    // TransactionData used by Order
    type TransactionData struct {

        
            TerminalID string  `json:"terminal_id"`
            PaymentID string  `json:"payment_id"`
            TransactionID string  `json:"transaction_id"`
            Currency string  `json:"currency"`
            AmountPaid string  `json:"amount_paid"`
            UniqueReferenceNumber string  `json:"unique_reference_number"`
            Status string  `json:"status"`
            Entity string  `json:"entity"`
         
    }
    
    // PlatformUserDetails used by Order
    type PlatformUserDetails struct {

        
            PlatformUserID string  `json:"platform_user_id"`
            PlatformUserEmployeeCode string  `json:"platform_user_employee_code"`
            PlatformUserLastName string  `json:"platform_user_last_name"`
            PlatformUserFirstName string  `json:"platform_user_first_name"`
         
    }
    
    // BillingStaffDetails used by Order
    type BillingStaffDetails struct {

        
            User string  `json:"user"`
            FirstName string  `json:"first_name"`
            StaffID float64  `json:"staff_id"`
            EmployeeCode string  `json:"employee_code"`
            LastName string  `json:"last_name"`
         
    }
    
    // OrderMeta used by Order
    type OrderMeta struct {

        
            OrderTags []map[string]interface{}  `json:"order_tags"`
            TransactionData TransactionData  `json:"transaction_data"`
            EmployeeID float64  `json:"employee_id"`
            Comment string  `json:"comment"`
            MongoCartID float64  `json:"mongo_cart_id"`
            OrderingStore float64  `json:"ordering_store"`
            OrderType string  `json:"order_type"`
            Staff map[string]interface{}  `json:"staff"`
            PaymentType string  `json:"payment_type"`
            CompanyLogo string  `json:"company_logo"`
            PlatformUserDetails PlatformUserDetails  `json:"platform_user_details"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            BillingStaffDetails BillingStaffDetails  `json:"billing_staff_details"`
            CustomerNote string  `json:"customer_note"`
            CurrencySymbol string  `json:"currency_symbol"`
            CartID float64  `json:"cart_id"`
            Files []map[string]interface{}  `json:"files"`
            OrderPlatform string  `json:"order_platform"`
            OrderChildEntities []string  `json:"order_child_entities"`
         
    }
    
    // TaxDetails used by Order
    type TaxDetails struct {

        
            Gstin string  `json:"gstin"`
            PanNo string  `json:"pan_no"`
         
    }
    
    // OrderDict used by Order
    type OrderDict struct {

        
            Meta OrderMeta  `json:"meta"`
            Prices Prices  `json:"prices"`
            FyndOrderID string  `json:"fynd_order_id"`
            PaymentMethods map[string]interface{}  `json:"payment_methods"`
            OrderDate string  `json:"order_date"`
            TaxDetails TaxDetails  `json:"tax_details"`
         
    }
    
    // ShipmentDetailsResponse used by Order
    type ShipmentDetailsResponse struct {

        
            Success bool  `json:"success"`
            Order OrderDict  `json:"order"`
            Shipments []PlatformShipment  `json:"shipments"`
         
    }
    
    // SubLane used by Order
    type SubLane struct {

        
            Text string  `json:"text"`
            TotalItems float64  `json:"total_items"`
            Value string  `json:"value"`
            Index float64  `json:"index"`
            Actions []map[string]interface{}  `json:"actions"`
         
    }
    
    // SuperLane used by Order
    type SuperLane struct {

        
            Text string  `json:"text"`
            TotalItems float64  `json:"total_items"`
            Options []SubLane  `json:"options"`
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

        
            Meta map[string]interface{}  `json:"meta"`
            BreakupValues []PlatformBreakupValues  `json:"breakup_values"`
            OrderCreatedTime string  `json:"order_created_time"`
            PaymentMode string  `json:"payment_mode"`
            TotalOrderValue float64  `json:"total_order_value"`
            Shipments []PlatformShipment  `json:"shipments"`
            Channel PlatformChannel  `json:"channel"`
            OrderID string  `json:"order_id"`
            OrderValue float64  `json:"order_value"`
            UserInfo UserDataInfo  `json:"user_info"`
         
    }
    
    // OrderListingResponse used by Order
    type OrderListingResponse struct {

        
            Lane string  `json:"lane"`
            Success bool  `json:"success"`
            Page Page  `json:"page"`
            TotalCount float64  `json:"total_count"`
            Message string  `json:"message"`
            Items []PlatformOrderItems  `json:"items"`
         
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
            Value float64  `json:"value"`
            Key string  `json:"key"`
         
    }
    
    // MetricCountResponse used by Order
    type MetricCountResponse struct {

        
            Items []MetricsCount  `json:"items"`
         
    }
    
    // PlatformTrack used by Order
    type PlatformTrack struct {

        
            Meta map[string]interface{}  `json:"meta"`
            LastLocationRecievedAt string  `json:"last_location_recieved_at"`
            RawStatus string  `json:"raw_status"`
            Awb string  `json:"awb"`
            AccountName string  `json:"account_name"`
            UpdatedAt string  `json:"updated_at"`
            Status string  `json:"status"`
            ShipmentType string  `json:"shipment_type"`
            Reason string  `json:"reason"`
            UpdatedTime string  `json:"updated_time"`
         
    }
    
    // PlatformShipmentTrack used by Order
    type PlatformShipmentTrack struct {

        
            Meta map[string]interface{}  `json:"meta"`
            Results []PlatformTrack  `json:"results"`
         
    }
    
    // AdvanceFilterInfo used by Order
    type AdvanceFilterInfo struct {

        
            Filters []FiltersInfo  `json:"filters"`
            Unfulfilled []FiltersInfo  `json:"unfulfilled"`
            Processed []FiltersInfo  `json:"processed"`
            ActionCentre []FiltersInfo  `json:"action_centre"`
            Returned []FiltersInfo  `json:"returned"`
         
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
            DisplayName string  `json:"display_name"`
            ReportName string  `json:"report_name"`
            ReportRequestedAt string  `json:"report_requested_at"`
            S3Key string  `json:"s3_key"`
            Status string  `json:"status"`
            ReportType string  `json:"report_type"`
            RequestDetails map[string]interface{}  `json:"request_details"`
         
    }
    
    // JioCodeUpsertDataSet used by Order
    type JioCodeUpsertDataSet struct {

        
            ItemID string  `json:"item_id"`
            ArticleID string  `json:"article_id"`
            CompanyID string  `json:"company_id"`
            JioCode string  `json:"jio_code"`
         
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

        
            Success bool  `json:"success"`
            Data []map[string]interface{}  `json:"data"`
            Identifier string  `json:"identifier"`
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

        
            DoInvoiceLabelGenerated bool  `json:"do_invoice_label_generated"`
            StoreID string  `json:"store_id"`
            StoreCode string  `json:"store_code"`
            Label map[string]interface{}  `json:"label"`
            Data map[string]interface{}  `json:"data"`
            InvoiceStatus string  `json:"invoice_status"`
            BatchID string  `json:"batch_id"`
            Invoice map[string]interface{}  `json:"invoice"`
            CompanyID string  `json:"company_id"`
            StoreName string  `json:"store_name"`
         
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
            FilePath string  `json:"file_path"`
            Namespace string  `json:"namespace"`
            Method string  `json:"method"`
            Cdn URL  `json:"cdn"`
            Tags []string  `json:"tags"`
            Operation string  `json:"operation"`
            Size float64  `json:"size"`
            FileName string  `json:"file_name"`
            ContentType string  `json:"content_type"`
         
    }
    
    // bulkListingData used by Order
    type bulkListingData struct {

        
            StoreID float64  `json:"store_id"`
            Successful float64  `json:"successful"`
            FailedShipments []map[string]interface{}  `json:"failed_shipments"`
            SuccessfulShipments []map[string]interface{}  `json:"successful_shipments"`
            FileName string  `json:"file_name"`
            CompanyID float64  `json:"company_id"`
            StoreCode string  `json:"store_code"`
            Failed float64  `json:"failed"`
            UserName string  `json:"user_name"`
            StoreName string  `json:"store_name"`
            UploadedOn string  `json:"uploaded_on"`
            ProcessingShipments []string  `json:"processing_shipments"`
            ExcelURL string  `json:"excel_url"`
            Status string  `json:"status"`
            UserID string  `json:"user_id"`
            ID string  `json:"id"`
            Total float64  `json:"total"`
            BatchID string  `json:"batch_id"`
            Processing float64  `json:"processing"`
         
    }
    
    // BulkListingPage used by Order
    type BulkListingPage struct {

        
            Type string  `json:"type"`
            Total float64  `json:"total"`
            HasPrevious bool  `json:"has_previous"`
            HasNext bool  `json:"has_next"`
            Current float64  `json:"current"`
            Size float64  `json:"size"`
         
    }
    
    // BulkListingResponse used by Order
    type BulkListingResponse struct {

        
            Success bool  `json:"success"`
            Error string  `json:"error"`
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

        
            QcType []string  `json:"qc_type"`
            QuestionSet []QuestionSet  `json:"question_set"`
            ID float64  `json:"id"`
            DisplayName string  `json:"display_name"`
         
    }
    
    // PlatformShipmentReasonsResponse used by Order
    type PlatformShipmentReasonsResponse struct {

        
            Reasons []Reason  `json:"reasons"`
            Success bool  `json:"success"`
         
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

        
            SuccessfulShipmentIds []string  `json:"successful_shipment_ids"`
            SuccessfulShipmentsCount float64  `json:"successful_shipments_count"`
            TotalShipmentsCount float64  `json:"total_shipments_count"`
            ProcessingShipmentsCount float64  `json:"processing_shipments_count"`
            BatchID string  `json:"batch_id"`
            FailedShipmentsCount float64  `json:"failed_shipments_count"`
            CompanyID string  `json:"company_id"`
         
    }
    
    // BulkActionDetailsResponse used by Order
    type BulkActionDetailsResponse struct {

        
            Success string  `json:"success"`
            Data []BulkActionDetailsDataField  `json:"data"`
            UploadedOn string  `json:"uploaded_on"`
            UploadedBy string  `json:"uploaded_by"`
            Message string  `json:"message"`
            Status bool  `json:"status"`
            UserID string  `json:"user_id"`
            Error []string  `json:"error"`
            FailedRecords []string  `json:"failed_records"`
         
    }
    
    // ArticleDetails used by Order
    type ArticleDetails struct {

        
            Status map[string]interface{}  `json:"status"`
         
    }
    
    // AffiliateBagDetails used by Order
    type AffiliateBagDetails struct {

        
            AffiliateBagID string  `json:"affiliate_bag_id"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            AffiliateMeta AffiliateMeta  `json:"affiliate_meta"`
            EmployeeDiscount float64  `json:"employee_discount"`
         
    }
    
    // Document used by Order
    type Document struct {

        
            DsType string  `json:"ds_type"`
            Value string  `json:"value"`
            LegalName string  `json:"legal_name"`
            Verified bool  `json:"verified"`
            URL string  `json:"url"`
         
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

        
            Username string  `json:"username"`
            Password string  `json:"password"`
            User string  `json:"user"`
            Enabled bool  `json:"enabled"`
         
    }
    
    // StoreGstCredentials used by Order
    type StoreGstCredentials struct {

        
            EWaybill StoreEwaybill  `json:"e_waybill"`
            EInvoice StoreEinvoice  `json:"e_invoice"`
         
    }
    
    // EInvoicePortalDetails used by Order
    type EInvoicePortalDetails struct {

        
            Username string  `json:"username"`
            Password string  `json:"password"`
            User string  `json:"user"`
         
    }
    
    // StoreMeta used by Order
    type StoreMeta struct {

        
            Documents StoreDocuments  `json:"documents"`
            ProductReturnConfig map[string]interface{}  `json:"product_return_config"`
            DisplayName string  `json:"display_name"`
            AdditionalContactDetails map[string]interface{}  `json:"additional_contact_details"`
            GstCredentials StoreGstCredentials  `json:"gst_credentials"`
            Stage string  `json:"stage"`
            NotificationEmails []string  `json:"notification_emails"`
            EwaybillPortalDetails map[string]interface{}  `json:"ewaybill_portal_details"`
            Timing []map[string]interface{}  `json:"timing"`
            EinvoicePortalDetails EInvoicePortalDetails  `json:"einvoice_portal_details"`
            GstNumber string  `json:"gst_number"`
         
    }
    
    // StoreAddress used by Order
    type StoreAddress struct {

        
            Phone string  `json:"phone"`
            AddressCategory string  `json:"address_category"`
            Address1 string  `json:"address1"`
            AddressType string  `json:"address_type"`
            State string  `json:"state"`
            Pincode float64  `json:"pincode"`
            Longitude float64  `json:"longitude"`
            UpdatedAt string  `json:"updated_at"`
            Country string  `json:"country"`
            City string  `json:"city"`
            ContactPerson string  `json:"contact_person"`
            Address2 string  `json:"address2"`
            Email string  `json:"email"`
            CreatedAt string  `json:"created_at"`
            Landmark string  `json:"landmark"`
            CountryCode string  `json:"country_code"`
            Version string  `json:"version"`
            Latitude float64  `json:"latitude"`
            Area string  `json:"area"`
         
    }
    
    // Store used by Order
    type Store struct {

        
            BrandStoreTags []string  `json:"brand_store_tags"`
            Phone float64  `json:"phone"`
            Code string  `json:"code"`
            IsEnabledForRecon bool  `json:"is_enabled_for_recon"`
            Meta StoreMeta  `json:"meta"`
            Address1 string  `json:"address1"`
            FulfillmentChannel string  `json:"fulfillment_channel"`
            IsArchived bool  `json:"is_archived"`
            CompanyID float64  `json:"company_id"`
            State string  `json:"state"`
            Pincode string  `json:"pincode"`
            OrderIntegrationID string  `json:"order_integration_id"`
            Longitude float64  `json:"longitude"`
            Name string  `json:"name"`
            UpdatedAt string  `json:"updated_at"`
            Country string  `json:"country"`
            City string  `json:"city"`
            LoginUsername string  `json:"login_username"`
            ContactPerson string  `json:"contact_person"`
            Address2 string  `json:"address2"`
            SID string  `json:"s_id"`
            CreatedAt string  `json:"created_at"`
            PackagingMaterialCount float64  `json:"packaging_material_count"`
            IsActive bool  `json:"is_active"`
            AlohomoraUserID float64  `json:"alohomora_user_id"`
            MallArea string  `json:"mall_area"`
            BrandID interface{}  `json:"brand_id"`
            StoreEmail string  `json:"store_email"`
            VatNo string  `json:"vat_no"`
            Latitude float64  `json:"latitude"`
            MallName string  `json:"mall_name"`
            ParentStoreID float64  `json:"parent_store_id"`
            StoreActiveFrom string  `json:"store_active_from"`
            LocationType string  `json:"location_type"`
            StoreAddressJson StoreAddress  `json:"store_address_json"`
         
    }
    
    // Dates used by Order
    type Dates struct {

        
            DeliveryDate interface{}  `json:"delivery_date"`
            OrderCreated string  `json:"order_created"`
         
    }
    
    // Attributes used by Order
    type Attributes struct {

        
            PrimaryMaterial string  `json:"primary_material"`
            PrimaryColor string  `json:"primary_color"`
            BrandName string  `json:"brand_name"`
            Gender []string  `json:"gender"`
            Essential string  `json:"essential"`
            MarketerAddress string  `json:"marketer_address"`
            Name string  `json:"name"`
            MarketerName string  `json:"marketer_name"`
            PrimaryColorHex string  `json:"primary_color_hex"`
         
    }
    
    // Item used by Order
    type Item struct {

        
            Meta map[string]interface{}  `json:"meta"`
            Code string  `json:"code"`
            Gender string  `json:"gender"`
            LastUpdatedAt string  `json:"last_updated_at"`
            Image []string  `json:"image"`
            WebstoreProductURL string  `json:"webstore_product_url"`
            L3Category float64  `json:"l3_category"`
            Color string  `json:"color"`
            CanCancel bool  `json:"can_cancel"`
            CanReturn bool  `json:"can_return"`
            Name string  `json:"name"`
            ItemID float64  `json:"item_id"`
            L1Category []string  `json:"l1_category"`
            BranchURL string  `json:"branch_url"`
            DepartmentID float64  `json:"department_id"`
            L2Category []string  `json:"l2_category"`
            L3CategoryName string  `json:"l3_category_name"`
            SlugKey string  `json:"slug_key"`
            Brand string  `json:"brand"`
            BrandID float64  `json:"brand_id"`
            L2CategoryID float64  `json:"l2_category_id"`
            L1CategoryID float64  `json:"l1_category_id"`
            Attributes Attributes  `json:"attributes"`
            Size string  `json:"size"`
         
    }
    
    // BagGSTDetails used by Order
    type BagGSTDetails struct {

        
            HsnCodeID string  `json:"hsn_code_id"`
            ValueOfGood float64  `json:"value_of_good"`
            GstTag string  `json:"gst_tag"`
            GstFee float64  `json:"gst_fee"`
            CgstGstFee string  `json:"cgst_gst_fee"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            IgstGstFee string  `json:"igst_gst_fee"`
            GstinCode string  `json:"gstin_code"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            IsDefaultHsnCode bool  `json:"is_default_hsn_code"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            IgstTaxPercentage float64  `json:"igst_tax_percentage"`
            SgstTaxPercentage float64  `json:"sgst_tax_percentage"`
            SgstGstFee string  `json:"sgst_gst_fee"`
            HsnCode string  `json:"hsn_code"`
            CgstTaxPercentage float64  `json:"cgst_tax_percentage"`
         
    }
    
    // Brand used by Order
    type Brand struct {

        
            IsVirtualInvoice bool  `json:"is_virtual_invoice"`
            InvoicePrefix string  `json:"invoice_prefix"`
            CreditNoteExpiryDays float64  `json:"credit_note_expiry_days"`
            Logo string  `json:"logo"`
            StartDate string  `json:"start_date"`
            ModifiedOn float64  `json:"modified_on"`
            BrandName string  `json:"brand_name"`
            PickupLocation string  `json:"pickup_location"`
            CreditNoteAllowed bool  `json:"credit_note_allowed"`
            ScriptLastRan string  `json:"script_last_ran"`
            BrandID float64  `json:"brand_id"`
            Company string  `json:"company"`
            CreatedOn float64  `json:"created_on"`
         
    }
    
    // BagReturnableCancelableStatus used by Order
    type BagReturnableCancelableStatus struct {

        
            IsReturnable bool  `json:"is_returnable"`
            IsCustomerReturnAllowed bool  `json:"is_customer_return_allowed"`
            CanBeCancelled bool  `json:"can_be_cancelled"`
            IsActive bool  `json:"is_active"`
            EnableTracking bool  `json:"enable_tracking"`
         
    }
    
    // ReturnConfig used by Order
    type ReturnConfig struct {

        
            Unit string  `json:"unit"`
            Returnable bool  `json:"returnable"`
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

        
            UID string  `json:"uid"`
            Code string  `json:"code"`
            SellerIdentifier string  `json:"seller_identifier"`
            RawMeta interface{}  `json:"raw_meta"`
            Identifiers Identifier  `json:"identifiers"`
            ReturnConfig ReturnConfig  `json:"return_config"`
            ASet map[string]interface{}  `json:"a_set"`
            Dimensions Dimensions  `json:"dimensions"`
            Weight Weight  `json:"weight"`
            ChildDetails map[string]interface{}  `json:"child_details"`
            EspModified interface{}  `json:"esp_modified"`
            IsSet bool  `json:"is_set"`
            Size string  `json:"size"`
            ID string  `json:"_id"`
         
    }
    
    // BagDetailsPlatformResponse used by Order
    type BagDetailsPlatformResponse struct {

        
            AppliedPromos []map[string]interface{}  `json:"applied_promos"`
            ArticleDetails ArticleDetails  `json:"article_details"`
            BagStatusHistory BagStatusHistory  `json:"bag_status_history"`
            Meta BagMeta  `json:"meta"`
            ShipmentID string  `json:"shipment_id"`
            NoOfBagsOrder float64  `json:"no_of_bags_order"`
            AffiliateDetails AffiliateDetails  `json:"affiliate_details"`
            BagStatus []BagStatusHistory  `json:"bag_status"`
            AffiliateBagDetails AffiliateBagDetails  `json:"affiliate_bag_details"`
            EntityType string  `json:"entity_type"`
            Quantity float64  `json:"quantity"`
            OperationalStatus string  `json:"operational_status"`
            FinancialBreakup []FinancialBreakup  `json:"financial_breakup"`
            OrderIntegrationID string  `json:"order_integration_id"`
            BType string  `json:"b_type"`
            OrderingStore Store  `json:"ordering_store"`
            QcRequired interface{}  `json:"qc_required"`
            Reasons []map[string]interface{}  `json:"reasons"`
            Dates Dates  `json:"dates"`
            CurrentOperationalStatus BagStatusHistory  `json:"current_operational_status"`
            ParentPromoBags map[string]interface{}  `json:"parent_promo_bags"`
            Item Item  `json:"item"`
            GstDetails BagGSTDetails  `json:"gst_details"`
            DisplayName string  `json:"display_name"`
            Prices Prices  `json:"prices"`
            Identifier string  `json:"identifier"`
            LineNumber float64  `json:"line_number"`
            Brand Brand  `json:"brand"`
            Status BagReturnableCancelableStatus  `json:"status"`
            BagUpdateTime float64  `json:"bag_update_time"`
            BID float64  `json:"b_id"`
            SellerIdentifier string  `json:"seller_identifier"`
            JourneyType string  `json:"journey_type"`
            Tags []string  `json:"tags"`
            RestoreCoupon bool  `json:"restore_coupon"`
            CurrentStatus BagStatusHistory  `json:"current_status"`
            RestorePromos map[string]interface{}  `json:"restore_promos"`
            OriginalBagList []float64  `json:"original_bag_list"`
            Article Article  `json:"article"`
         
    }
    
    // ErrorResponse used by Order
    type ErrorResponse struct {

        
            Error string  `json:"error"`
            Message string  `json:"message"`
         
    }
    
    // BagsPage used by Order
    type BagsPage struct {

        
            ItemTotal float64  `json:"item_total"`
            HasNext bool  `json:"has_next"`
            Current float64  `json:"current"`
            PageType string  `json:"page_type"`
            Size float64  `json:"size"`
         
    }
    
    // GetBagsPlatformResponse used by Order
    type GetBagsPlatformResponse struct {

        
            Items []BagDetailsPlatformResponse  `json:"items"`
            Page BagsPage  `json:"page"`
         
    }
    
    // GeneratePosOrderReceiptResponse used by Order
    type GeneratePosOrderReceiptResponse struct {

        
            PaymentReceipt string  `json:"payment_receipt"`
            Success bool  `json:"success"`
            InvoiceReceipt string  `json:"invoice_receipt"`
            OrderID string  `json:"order_id"`
         
    }
    
    // InvalidateShipmentCachePayload used by Order
    type InvalidateShipmentCachePayload struct {

        
            ShipmentIds []string  `json:"shipment_ids"`
         
    }
    
    // InvalidateShipmentCacheNestedResponse used by Order
    type InvalidateShipmentCacheNestedResponse struct {

        
            Message string  `json:"message"`
            ShipmentID string  `json:"shipment_id"`
            Error string  `json:"error"`
            Status float64  `json:"status"`
         
    }
    
    // InvalidateShipmentCacheResponse used by Order
    type InvalidateShipmentCacheResponse struct {

        
            Response []InvalidateShipmentCacheNestedResponse  `json:"response"`
         
    }
    
    // ErrorResponse1 used by Order
    type ErrorResponse1 struct {

        
            Message string  `json:"message"`
            ErrorTrace string  `json:"error_trace"`
            Status float64  `json:"status"`
         
    }
    
    // StoreReassign used by Order
    type StoreReassign struct {

        
            AffiliateBagID string  `json:"affiliate_bag_id"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            BagID float64  `json:"bag_id"`
            AffiliateID string  `json:"affiliate_id"`
            ReasonIds []float64  `json:"reason_ids"`
            ItemID string  `json:"item_id"`
            FyndOrderID string  `json:"fynd_order_id"`
            SetID string  `json:"set_id"`
            StoreID float64  `json:"store_id"`
            MongoArticleID string  `json:"mongo_article_id"`
         
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
            ReasonText string  `json:"reason_text"`
            ID string  `json:"id"`
            AffiliateID string  `json:"affiliate_id"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
         
    }
    
    // UpdateShipmentLockPayload used by Order
    type UpdateShipmentLockPayload struct {

        
            EntityType string  `json:"entity_type"`
            Action string  `json:"action"`
            Entities []Entities  `json:"entities"`
            ActionType string  `json:"action_type"`
         
    }
    
    // OriginalFilter used by Order
    type OriginalFilter struct {

        
            AffiliateID string  `json:"affiliate_id"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
         
    }
    
    // Bags used by Order
    type Bags struct {

        
            IsLocked bool  `json:"is_locked"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
            BagID float64  `json:"bag_id"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
         
    }
    
    // CheckResponse used by Order
    type CheckResponse struct {

        
            OriginalFilter OriginalFilter  `json:"original_filter"`
            LockStatus bool  `json:"lock_status"`
            AffiliateID string  `json:"affiliate_id"`
            IsBagLocked bool  `json:"is_bag_locked"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            Bags []Bags  `json:"bags"`
            IsShipmentLocked bool  `json:"is_shipment_locked"`
            ShipmentID string  `json:"shipment_id"`
            Status string  `json:"status"`
         
    }
    
    // UpdateShipmentLockResponse used by Order
    type UpdateShipmentLockResponse struct {

        
            Success bool  `json:"success"`
            CheckResponse []CheckResponse  `json:"check_response"`
            Message string  `json:"message"`
         
    }
    
    // AnnouncementResponse used by Order
    type AnnouncementResponse struct {

        
            PlatformName string  `json:"platform_name"`
            PlatformID string  `json:"platform_id"`
            CreatedAt string  `json:"created_at"`
            Description string  `json:"description"`
            ID float64  `json:"id"`
            Title string  `json:"title"`
            CompanyID float64  `json:"company_id"`
            FromDatetime string  `json:"from_datetime"`
            ToDatetime string  `json:"to_datetime"`
            LogoURL string  `json:"logo_url"`
         
    }
    
    // AnnouncementsResponse used by Order
    type AnnouncementsResponse struct {

        
            Announcements []AnnouncementResponse  `json:"announcements"`
         
    }
    
    // BaseResponse used by Order
    type BaseResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // Click2CallResponse used by Order
    type Click2CallResponse struct {

        
            CallID string  `json:"call_id"`
            Status bool  `json:"status"`
         
    }
    
    // ProductsReasonsFilters used by Order
    type ProductsReasonsFilters struct {

        
            Identifier string  `json:"identifier"`
            Quantity float64  `json:"quantity"`
            LineNumber float64  `json:"line_number"`
         
    }
    
    // ProductsReasonsData used by Order
    type ProductsReasonsData struct {

        
            ReasonText string  `json:"reason_text"`
            ReasonID float64  `json:"reason_id"`
         
    }
    
    // ProductsReasons used by Order
    type ProductsReasons struct {

        
            Filters []ProductsReasonsFilters  `json:"filters"`
            Data ProductsReasonsData  `json:"data"`
         
    }
    
    // EntityReasonData used by Order
    type EntityReasonData struct {

        
            ReasonText string  `json:"reason_text"`
            ReasonID float64  `json:"reason_id"`
         
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
            Quantity float64  `json:"quantity"`
            LineNumber float64  `json:"line_number"`
         
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

        
            Reasons ReasonsData  `json:"reasons"`
            Products []Products  `json:"products"`
            Identifier string  `json:"identifier"`
            DataUpdates DataUpdates  `json:"data_updates"`
         
    }
    
    // StatuesRequest used by Order
    type StatuesRequest struct {

        
            ExcludeBagsNextState string  `json:"exclude_bags_next_state"`
            Shipments []ShipmentsRequest  `json:"shipments"`
            Status string  `json:"status"`
         
    }
    
    // UpdateShipmentStatusRequest used by Order
    type UpdateShipmentStatusRequest struct {

        
            UnlockBeforeTransition bool  `json:"unlock_before_transition"`
            ForceTransition bool  `json:"force_transition"`
            Task bool  `json:"task"`
            LockAfterTransition bool  `json:"lock_after_transition"`
            Statuses []StatuesRequest  `json:"statuses"`
         
    }
    
    // ShipmentsResponse used by Order
    type ShipmentsResponse struct {

        
            Exception string  `json:"exception"`
            Message string  `json:"message"`
            Meta map[string]interface{}  `json:"meta"`
            StackTrace string  `json:"stack_trace"`
            FinalState map[string]interface{}  `json:"final_state"`
            Code string  `json:"code"`
            Identifier string  `json:"identifier"`
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

        
            State string  `json:"state"`
            Phone float64  `json:"phone"`
            Address2 string  `json:"address2"`
            Email string  `json:"email"`
            Country string  `json:"country"`
            Address1 string  `json:"address1"`
            FirstName string  `json:"first_name"`
            Mobile float64  `json:"mobile"`
            LastName string  `json:"last_name"`
            City string  `json:"city"`
            Pincode string  `json:"pincode"`
         
    }
    
    // UserData used by Order
    type UserData struct {

        
            BillingUser OrderUser  `json:"billing_user"`
            ShippingUser OrderUser  `json:"shipping_user"`
         
    }
    
    // OrderPriority used by Order
    type OrderPriority struct {

        
            FulfilmentPriorityText string  `json:"fulfilment_priority_text"`
            FulfilmentPriority float64  `json:"fulfilment_priority"`
            AffiliatePriorityCode string  `json:"affiliate_priority_code"`
         
    }
    
    // MarketPlacePdf used by Order
    type MarketPlacePdf struct {

        
            Label string  `json:"label"`
            Invoice string  `json:"invoice"`
         
    }
    
    // AffiliateBag used by Order
    type AffiliateBag struct {

        
            AffiliateMeta map[string]interface{}  `json:"affiliate_meta"`
            FyndStoreID string  `json:"fynd_store_id"`
            DeliveryCharge float64  `json:"delivery_charge"`
            AffiliateStoreID string  `json:"affiliate_store_id"`
            SellerIdentifier string  `json:"seller_identifier"`
            PriceMarked float64  `json:"price_marked"`
            ItemSize string  `json:"item_size"`
            AvlQty float64  `json:"avl_qty"`
            Quantity float64  `json:"quantity"`
            AmountPaid float64  `json:"amount_paid"`
            ItemID float64  `json:"item_id"`
            HsnCodeID string  `json:"hsn_code_id"`
            UnitPrice float64  `json:"unit_price"`
            PdfLinks MarketPlacePdf  `json:"pdf_links"`
            StoreID float64  `json:"store_id"`
            Discount float64  `json:"discount"`
            ModifiedOn string  `json:"modified_on"`
            ID string  `json:"_id"`
            CompanyID float64  `json:"company_id"`
            Sku string  `json:"sku"`
            PriceEffective float64  `json:"price_effective"`
            Identifier map[string]interface{}  `json:"identifier"`
            TransferPrice float64  `json:"transfer_price"`
         
    }
    
    // ArticleDetails1 used by Order
    type ArticleDetails1 struct {

        
            ID string  `json:"_id"`
            Attributes map[string]interface{}  `json:"attributes"`
            Quantity float64  `json:"quantity"`
            Dimension map[string]interface{}  `json:"dimension"`
            Category map[string]interface{}  `json:"category"`
            BrandID float64  `json:"brand_id"`
            Weight map[string]interface{}  `json:"weight"`
         
    }
    
    // LocationDetails used by Order
    type LocationDetails struct {

        
            Articles []ArticleDetails1  `json:"articles"`
            FulfillmentType string  `json:"fulfillment_type"`
            FulfillmentID float64  `json:"fulfillment_id"`
         
    }
    
    // ShipmentDetails used by Order
    type ShipmentDetails struct {

        
            Meta map[string]interface{}  `json:"meta"`
            BoxType string  `json:"box_type"`
            FulfillmentID float64  `json:"fulfillment_id"`
            Articles []ArticleDetails1  `json:"articles"`
            DpID float64  `json:"dp_id"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            Shipments float64  `json:"shipments"`
         
    }
    
    // ShipmentConfig used by Order
    type ShipmentConfig struct {

        
            Action string  `json:"action"`
            ToPincode string  `json:"to_pincode"`
            LocationDetails LocationDetails  `json:"location_details"`
            PaymentMode string  `json:"payment_mode"`
            Journey string  `json:"journey"`
            Identifier string  `json:"identifier"`
            Source string  `json:"source"`
            Shipment []ShipmentDetails  `json:"shipment"`
         
    }
    
    // ShipmentData used by Order
    type ShipmentData struct {

        
            ShipmentData ShipmentConfig  `json:"shipment_data"`
         
    }
    
    // OrderInfo used by Order
    type OrderInfo struct {

        
            User UserData  `json:"user"`
            Discount float64  `json:"discount"`
            OrderPriority OrderPriority  `json:"order_priority"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            Coupon string  `json:"coupon"`
            PaymentMode string  `json:"payment_mode"`
            DeliveryCharges float64  `json:"delivery_charges"`
            BillingAddress OrderUser  `json:"billing_address"`
            Payment map[string]interface{}  `json:"payment"`
            Items map[string]interface{}  `json:"items"`
            OrderValue float64  `json:"order_value"`
            Bags []AffiliateBag  `json:"bags"`
            ShippingAddress OrderUser  `json:"shipping_address"`
            Shipment ShipmentData  `json:"shipment"`
            CodCharges float64  `json:"cod_charges"`
         
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
    
    // AffiliateInventoryConfig used by Order
    type AffiliateInventoryConfig struct {

        
            Order AffiliateInventoryOrderConfig  `json:"order"`
            Inventory AffiliateInventoryStoreConfig  `json:"inventory"`
            Payment AffiliateInventoryPaymentConfig  `json:"payment"`
            Logistics AffiliateInventoryLogisticsConfig  `json:"logistics"`
            ArticleAssignment AffiliateInventoryArticleAssignmentConfig  `json:"article_assignment"`
         
    }
    
    // AffiliateAppConfigMeta used by Order
    type AffiliateAppConfigMeta struct {

        
            Name string  `json:"name"`
            Value string  `json:"value"`
         
    }
    
    // AffiliateAppConfig used by Order
    type AffiliateAppConfig struct {

        
            UpdatedAt string  `json:"updated_at"`
            Meta []AffiliateAppConfigMeta  `json:"meta"`
            Secret string  `json:"secret"`
            Token string  `json:"token"`
            CreatedAt string  `json:"created_at"`
            Description string  `json:"description"`
            ID string  `json:"id"`
            Name string  `json:"name"`
            Owner string  `json:"owner"`
         
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
            BagEndState string  `json:"bag_end_state"`
            StoreLookup string  `json:"store_lookup"`
            AffiliateStoreIDMapping []AffiliateStoreIdMapping  `json:"affiliate_store_id_mapping"`
            ArticleLookup string  `json:"article_lookup"`
            CreateUser bool  `json:"create_user"`
         
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
            DisplayText string  `json:"display_text"`
            Description string  `json:"description"`
            ID float64  `json:"id"`
         
    }
    
    // GetActionsResponse used by Order
    type GetActionsResponse struct {

        
            Permissions ActionInfo  `json:"permissions"`
         
    }
    
    // HistoryDict used by Order
    type HistoryDict struct {

        
            L2Detail string  `json:"l2_detail"`
            User string  `json:"user"`
            Message string  `json:"message"`
            TicketID string  `json:"ticket_id"`
            Createdat string  `json:"createdat"`
            L3Detail string  `json:"l3_detail"`
            BagID float64  `json:"bag_id"`
            TicketURL string  `json:"ticket_url"`
            Type string  `json:"type"`
            L1Detail string  `json:"l1_detail"`
         
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
    
    // SmsDataPayload used by Order
    type SmsDataPayload struct {

        
            Message string  `json:"message"`
            PhoneNumber float64  `json:"phone_number"`
            CountryCode string  `json:"country_code"`
            CustomerName string  `json:"customer_name"`
            AmountPaid float64  `json:"amount_paid"`
            BrandName string  `json:"brand_name"`
            PaymentMode string  `json:"payment_mode"`
            ShipmentID float64  `json:"shipment_id"`
            OrderID string  `json:"order_id"`
         
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
    
    // Meta1 used by Order
    type Meta1 struct {

        
            StateManagerUsed string  `json:"state_manager_used"`
            KafkaEmissionStatus float64  `json:"kafka_emission_status"`
         
    }
    
    // ShipmentDetail used by Order
    type ShipmentDetail struct {

        
            Meta Meta1  `json:"meta"`
            Remarks string  `json:"remarks"`
            ID float64  `json:"id"`
            ShipmentID string  `json:"shipment_id"`
            BagList []float64  `json:"bag_list"`
            Status string  `json:"status"`
         
    }
    
    // OrderStatusData used by Order
    type OrderStatusData struct {

        
            OrderDetails OrderDetails  `json:"order_details"`
            ShipmentDetails []ShipmentDetail  `json:"shipment_details"`
            Errors []string  `json:"errors"`
         
    }
    
    // OrderStatusResult used by Order
    type OrderStatusResult struct {

        
            Success string  `json:"success"`
            Result []OrderStatusData  `json:"result"`
         
    }
    
    // ManualAssignDPToShipment used by Order
    type ManualAssignDPToShipment struct {

        
            DpID float64  `json:"dp_id"`
            ShipmentIds []string  `json:"shipment_ids"`
            OrderType string  `json:"order_type"`
            QcRequired string  `json:"qc_required"`
         
    }
    
    // ManualAssignDPToShipmentResponse used by Order
    type ManualAssignDPToShipmentResponse struct {

        
            Success string  `json:"success"`
            Errors []string  `json:"errors"`
         
    }
    
    // TaxInfo used by Order
    type TaxInfo struct {

        
            Gstin string  `json:"gstin"`
            B2bGstinNumber string  `json:"b2b_gstin_number"`
         
    }
    
    // Tax used by Order
    type Tax struct {

        
            Breakup []map[string]interface{}  `json:"breakup"`
            Amount map[string]interface{}  `json:"amount"`
            Name string  `json:"name"`
            Rate float64  `json:"rate"`
         
    }
    
    // Charge used by Order
    type Charge struct {

        
            Tax Tax  `json:"tax"`
            Type string  `json:"type"`
            Code string  `json:"code"`
            Name string  `json:"name"`
            Amount map[string]interface{}  `json:"amount"`
         
    }
    
    // ShippingInfo used by Order
    type ShippingInfo struct {

        
            FloorNo string  `json:"floor_no"`
            AlternateMobileNumber string  `json:"alternate_mobile_number"`
            LastName string  `json:"last_name"`
            Slot []map[string]interface{}  `json:"slot"`
            City string  `json:"city"`
            Pincode string  `json:"pincode"`
            AddressType string  `json:"address_type"`
            Gender string  `json:"gender"`
            Landmark string  `json:"landmark"`
            State string  `json:"state"`
            Address2 string  `json:"address2"`
            PrimaryMobileNumber string  `json:"primary_mobile_number"`
            Country string  `json:"country"`
            Title string  `json:"title"`
            Address1 string  `json:"address1"`
            CustomerCode string  `json:"customer_code"`
            ExternalCustomerCode string  `json:"external_customer_code"`
            ShippingType string  `json:"shipping_type"`
            HouseNo string  `json:"house_no"`
            AlternateEmail string  `json:"alternate_email"`
            CountryCode string  `json:"country_code"`
            StateCode string  `json:"state_code"`
            PrimaryEmail string  `json:"primary_email"`
            GeoLocation map[string]interface{}  `json:"geo_location"`
            FirstName string  `json:"first_name"`
            MiddleName string  `json:"middle_name"`
         
    }
    
    // BillingInfo used by Order
    type BillingInfo struct {

        
            FloorNo string  `json:"floor_no"`
            AlternateMobileNumber string  `json:"alternate_mobile_number"`
            LastName string  `json:"last_name"`
            City string  `json:"city"`
            Pincode string  `json:"pincode"`
            Gender string  `json:"gender"`
            State string  `json:"state"`
            Address2 string  `json:"address2"`
            PrimaryMobileNumber string  `json:"primary_mobile_number"`
            Country string  `json:"country"`
            Title string  `json:"title"`
            Address1 string  `json:"address1"`
            CustomerCode string  `json:"customer_code"`
            ExternalCustomerCode string  `json:"external_customer_code"`
            HouseNo string  `json:"house_no"`
            AlternateEmail string  `json:"alternate_email"`
            CountryCode string  `json:"country_code"`
            StateCode string  `json:"state_code"`
            PrimaryEmail string  `json:"primary_email"`
            FirstName string  `json:"first_name"`
            MiddleName string  `json:"middle_name"`
         
    }
    
    // ProcessingDates used by Order
    type ProcessingDates struct {

        
            DpPickupSlot map[string]interface{}  `json:"dp_pickup_slot"`
            DispatchAfterDate string  `json:"dispatch_after_date"`
            ConfirmByDate string  `json:"confirm_by_date"`
            PackByDate string  `json:"pack_by_date"`
            DispatchByDate string  `json:"dispatch_by_date"`
            CustomerPickupSlot map[string]interface{}  `json:"customer_pickup_slot"`
         
    }
    
    // LineItem used by Order
    type LineItem struct {

        
            SellerIdentifier string  `json:"seller_identifier"`
            Meta map[string]interface{}  `json:"meta"`
            CustomMessasge string  `json:"custom_messasge"`
            Charges []Charge  `json:"charges"`
            ExternalLineID string  `json:"external_line_id"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // Shipment used by Order
    type Shipment struct {

        
            LocationID float64  `json:"location_id"`
            Meta map[string]interface{}  `json:"meta"`
            ExternalShipmentID string  `json:"external_shipment_id"`
            ProcessingDates ProcessingDates  `json:"processing_dates"`
            Priority float64  `json:"priority"`
            LineItems []LineItem  `json:"line_items"`
         
    }
    
    // PaymentMethod used by Order
    type PaymentMethod struct {

        
            TransactionData map[string]interface{}  `json:"transaction_data"`
            Meta map[string]interface{}  `json:"meta"`
            CollectBy string  `json:"collect_by"`
            Mode string  `json:"mode"`
            RefundBy string  `json:"refund_by"`
            Name string  `json:"name"`
            Amount float64  `json:"amount"`
         
    }
    
    // PaymentInfo used by Order
    type PaymentInfo struct {

        
            PrimaryMode string  `json:"primary_mode"`
            PaymentMethods []PaymentMethod  `json:"payment_methods"`
         
    }
    
    // CreateOrderAPI used by Order
    type CreateOrderAPI struct {

        
            Config map[string]interface{}  `json:"config"`
            Meta map[string]interface{}  `json:"meta"`
            TaxInfo TaxInfo  `json:"tax_info"`
            Charges []Charge  `json:"charges"`
            CurrencyInfo map[string]interface{}  `json:"currency_info"`
            ShippingInfo ShippingInfo  `json:"shipping_info"`
            BillingInfo BillingInfo  `json:"billing_info"`
            Shipments []Shipment  `json:"shipments"`
            ExternalOrderID string  `json:"external_order_id"`
            ExternalCreationDate string  `json:"external_creation_date"`
            PaymentInfo PaymentInfo  `json:"payment_info"`
         
    }
    
    // CreateOrderErrorReponse used by Order
    type CreateOrderErrorReponse struct {

        
            Exception string  `json:"exception"`
            Message string  `json:"message"`
            Meta string  `json:"meta"`
            StackTrace string  `json:"stack_trace"`
            Code string  `json:"code"`
            RequestID string  `json:"request_id"`
            Info interface{}  `json:"info"`
            Status float64  `json:"status"`
         
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

        
            ModeOfPayment string  `json:"mode_of_payment"`
            Source string  `json:"source"`
            PaymentMethods []PaymentMethods  `json:"payment_methods"`
         
    }
    
    // CreateChannelConfig used by Order
    type CreateChannelConfig struct {

        
            DpConfiguration DpConfiguration  `json:"dp_configuration"`
            ShipmentAssignment string  `json:"shipment_assignment"`
            PaymentInfo CreateChannelPaymentInfo  `json:"payment_info"`
            LockStates []string  `json:"lock_states"`
            LogoURL map[string]interface{}  `json:"logo_url"`
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
            Acknowledged bool  `json:"acknowledged"`
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
            OrderDetails []FyndOrderIdList  `json:"order_details"`
            EndDate string  `json:"end_date"`
            StartDate string  `json:"start_date"`
         
    }
    
    // FetchCreditBalanceRequestPayload used by Order
    type FetchCreditBalanceRequestPayload struct {

        
            CustomerMobileNumber string  `json:"customer_mobile_number"`
            SellerID string  `json:"seller_id"`
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
            AffiliateID string  `json:"affiliate_id"`
            FyndOrderID string  `json:"fynd_order_id"`
            OrderingChannel string  `json:"ordering_channel"`
            SellerID string  `json:"seller_id"`
         
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

        
            LastName string  `json:"last_name"`
            Mobile float64  `json:"mobile"`
            FirstName string  `json:"first_name"`
            CountryCode string  `json:"country_code"`
         
    }
    
    // AttachUserOtpData used by Order
    type AttachUserOtpData struct {

        
            RequestID string  `json:"request_id"`
            OtpCode float64  `json:"otp_code"`
         
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

        
            Mobile float64  `json:"mobile"`
            CountryCode string  `json:"country_code"`
         
    }
    
    // PointBlankOtpData used by Order
    type PointBlankOtpData struct {

        
            Mobile float64  `json:"mobile"`
            Message string  `json:"message"`
            RequestID string  `json:"request_id"`
            ResendTimer float64  `json:"resend_timer"`
         
    }
    
    // SendUserMobileOtpResponse used by Order
    type SendUserMobileOtpResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
            Data PointBlankOtpData  `json:"data"`
         
    }
    

    
    // DeleteResponse used by Catalog
    type DeleteResponse struct {

        
            Message string  `json:"message"`
         
    }
    
    // ErrorResponse used by Catalog
    type ErrorResponse struct {

        
            Code string  `json:"code"`
            Status float64  `json:"status"`
            Message string  `json:"message"`
            Error string  `json:"error"`
            Meta map[string]interface{}  `json:"meta"`
         
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
    
    // GetSearchWordsData used by Catalog
    type GetSearchWordsData struct {

        
            IsActive bool  `json:"is_active"`
            Words []string  `json:"words"`
            AppID string  `json:"app_id"`
            Result map[string]interface{}  `json:"result"`
            UID string  `json:"uid"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // GetSearchWordsDetailResponse used by Catalog
    type GetSearchWordsDetailResponse struct {

        
            Page Page  `json:"page"`
            Items GetSearchWordsData  `json:"items"`
         
    }
    
    // SearchKeywordResult used by Catalog
    type SearchKeywordResult struct {

        
            Query map[string]interface{}  `json:"query"`
            SortOn string  `json:"sort_on"`
         
    }
    
    // CreateSearchKeyword used by Catalog
    type CreateSearchKeyword struct {

        
            IsActive bool  `json:"is_active"`
            Words []string  `json:"words"`
            AppID string  `json:"app_id"`
            Result SearchKeywordResult  `json:"result"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // GetSearchWordsResponse used by Catalog
    type GetSearchWordsResponse struct {

        
            Page Page  `json:"page"`
            Items []GetSearchWordsData  `json:"items"`
         
    }
    
    // GetAutocompleteWordsData used by Catalog
    type GetAutocompleteWordsData struct {

        
            Words []string  `json:"words"`
            Results []map[string]interface{}  `json:"results"`
            AppID string  `json:"app_id"`
            UID string  `json:"uid"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // GetAutocompleteWordsResponse used by Catalog
    type GetAutocompleteWordsResponse struct {

        
            Page Page  `json:"page"`
            Items []GetAutocompleteWordsData  `json:"items"`
         
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

        
            Page AutocompletePageAction  `json:"page"`
            Type string  `json:"type"`
         
    }
    
    // Media used by Catalog
    type Media struct {

        
            Type string  `json:"type"`
            URL string  `json:"url"`
            AspectRatio string  `json:"aspect_ratio"`
         
    }
    
    // AutocompleteResult used by Catalog
    type AutocompleteResult struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Display string  `json:"display"`
            Action AutocompleteAction  `json:"action"`
            Logo Media  `json:"logo"`
         
    }
    
    // CreateAutocompleteKeyword used by Catalog
    type CreateAutocompleteKeyword struct {

        
            IsActive bool  `json:"is_active"`
            Words []string  `json:"words"`
            Results []AutocompleteResult  `json:"results"`
            AppID string  `json:"app_id"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // CreateAutocompleteWordsResponse used by Catalog
    type CreateAutocompleteWordsResponse struct {

        
            AppID string  `json:"app_id"`
            Words []string  `json:"words"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Results []map[string]interface{}  `json:"results"`
         
    }
    
    // ProductBundleItem used by Catalog
    type ProductBundleItem struct {

        
            AllowRemove bool  `json:"allow_remove"`
            ProductUID float64  `json:"product_uid"`
            MinQuantity float64  `json:"min_quantity"`
            AutoAddToCart bool  `json:"auto_add_to_cart"`
            MaxQuantity float64  `json:"max_quantity"`
            AutoSelect bool  `json:"auto_select"`
         
    }
    
    // GetProductBundleCreateResponse used by Catalog
    type GetProductBundleCreateResponse struct {

        
            IsActive bool  `json:"is_active"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ID string  `json:"id"`
            Name string  `json:"name"`
            CompanyID float64  `json:"company_id"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            Logo string  `json:"logo"`
            CreatedOn string  `json:"created_on"`
            Choice string  `json:"choice"`
            Products []ProductBundleItem  `json:"products"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Slug string  `json:"slug"`
            PageVisibility []string  `json:"page_visibility"`
            ModifiedOn string  `json:"modified_on"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // GetProductBundleListingResponse used by Catalog
    type GetProductBundleListingResponse struct {

        
            Page Page  `json:"page"`
            Items []GetProductBundleCreateResponse  `json:"items"`
         
    }
    
    // ProductBundleRequest used by Catalog
    type ProductBundleRequest struct {

        
            IsActive bool  `json:"is_active"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Name string  `json:"name"`
            CompanyID float64  `json:"company_id"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            Logo string  `json:"logo"`
            CreatedOn string  `json:"created_on"`
            Choice string  `json:"choice"`
            Products []ProductBundleItem  `json:"products"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Slug string  `json:"slug"`
            PageVisibility []string  `json:"page_visibility"`
            ModifiedOn string  `json:"modified_on"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // Price used by Catalog
    type Price struct {

        
            MaxMarked float64  `json:"max_marked"`
            Currency string  `json:"currency"`
            MaxEffective float64  `json:"max_effective"`
            MinEffective float64  `json:"min_effective"`
            MinMarked float64  `json:"min_marked"`
         
    }
    
    // Size used by Catalog
    type Size struct {

        
            Value string  `json:"value"`
            Display string  `json:"display"`
            Quantity float64  `json:"quantity"`
            IsAvailable bool  `json:"is_available"`
         
    }
    
    // LimitedProductData used by Catalog
    type LimitedProductData struct {

        
            CountryOfOrigin string  `json:"country_of_origin"`
            ItemCode string  `json:"item_code"`
            Price map[string]interface{}  `json:"price"`
            ShortDescription string  `json:"short_description"`
            Name string  `json:"name"`
            Identifier map[string]interface{}  `json:"identifier"`
            Images []string  `json:"images"`
            UID float64  `json:"uid"`
            Attributes map[string]interface{}  `json:"attributes"`
            Slug string  `json:"slug"`
            Quantity float64  `json:"quantity"`
            Sizes []string  `json:"sizes"`
         
    }
    
    // GetProducts used by Catalog
    type GetProducts struct {

        
            Price Price  `json:"price"`
            AllowRemove bool  `json:"allow_remove"`
            ProductUID float64  `json:"product_uid"`
            MinQuantity float64  `json:"min_quantity"`
            AutoAddToCart bool  `json:"auto_add_to_cart"`
            MaxQuantity float64  `json:"max_quantity"`
            Sizes []Size  `json:"sizes"`
            ProductDetails LimitedProductData  `json:"product_details"`
            AutoSelect bool  `json:"auto_select"`
         
    }
    
    // GetProductBundleResponse used by Catalog
    type GetProductBundleResponse struct {

        
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            CompanyID float64  `json:"company_id"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            Logo string  `json:"logo"`
            Choice string  `json:"choice"`
            Products []GetProducts  `json:"products"`
            Slug string  `json:"slug"`
            PageVisibility []string  `json:"page_visibility"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // ProductBundleUpdateRequest used by Catalog
    type ProductBundleUpdateRequest struct {

        
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            CompanyID float64  `json:"company_id"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            Logo string  `json:"logo"`
            Choice string  `json:"choice"`
            Products []ProductBundleItem  `json:"products"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Slug string  `json:"slug"`
            PageVisibility []string  `json:"page_visibility"`
            ModifiedOn string  `json:"modified_on"`
            Meta map[string]interface{}  `json:"meta"`
         
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

        
            CreatedBy map[string]interface{}  `json:"created_by"`
            Active bool  `json:"active"`
            Tag string  `json:"tag"`
            ID string  `json:"id"`
            Subtitle string  `json:"subtitle"`
            Name string  `json:"name"`
            CompanyID float64  `json:"company_id"`
            CreatedOn string  `json:"created_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Image string  `json:"image"`
            Title string  `json:"title"`
            BrandID float64  `json:"brand_id"`
            Guide Guide  `json:"guide"`
            ModifiedOn string  `json:"modified_on"`
            Description string  `json:"description"`
         
    }
    
    // SuccessResponse used by Catalog
    type SuccessResponse struct {

        
            Success bool  `json:"success"`
         
    }
    
    // SizeGuideResponse used by Catalog
    type SizeGuideResponse struct {

        
            CreatedBy map[string]interface{}  `json:"created_by"`
            Active bool  `json:"active"`
            Tag string  `json:"tag"`
            ID string  `json:"id"`
            Subtitle string  `json:"subtitle"`
            Name string  `json:"name"`
            CompanyID float64  `json:"company_id"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            BrandID float64  `json:"brand_id"`
            Title string  `json:"title"`
            Guide map[string]interface{}  `json:"guide"`
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
         
    }
    
    // MetaFields used by Catalog
    type MetaFields struct {

        
            Value string  `json:"value"`
            Key string  `json:"key"`
         
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

        
            CustomMeta []MetaFields  `json:"_custom_meta"`
            Moq ApplicationItemMOQ  `json:"moq"`
            IsGift bool  `json:"is_gift"`
            Seo ApplicationItemSEO  `json:"seo"`
            IsCod bool  `json:"is_cod"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            AltText map[string]interface{}  `json:"alt_text"`
         
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

        
            IsGift bool  `json:"is_gift"`
            Moq MOQData  `json:"moq"`
            Seo SEOData  `json:"seo"`
            IsCod bool  `json:"is_cod"`
            AltText map[string]interface{}  `json:"alt_text"`
         
    }
    
    // GetConfigMetadataResponse used by Catalog
    type GetConfigMetadataResponse struct {

        
            Data []map[string]interface{}  `json:"data"`
            Condition []map[string]interface{}  `json:"condition"`
            Values []map[string]interface{}  `json:"values"`
         
    }
    
    // PageResponseType used by Catalog
    type PageResponseType struct {

        
            Next float64  `json:"next"`
            TotalCount float64  `json:"total_count"`
            HasNext bool  `json:"has_next"`
            Current float64  `json:"current"`
         
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

        
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            Logo string  `json:"logo"`
            Priority float64  `json:"priority"`
            Unit string  `json:"unit"`
            DisplayType string  `json:"display_type"`
            Key string  `json:"key"`
            Slug string  `json:"slug"`
         
    }
    
    // AppConfigurationDetail used by Catalog
    type AppConfigurationDetail struct {

        
            IsActive bool  `json:"is_active"`
            IsDefault bool  `json:"is_default"`
            Name string  `json:"name"`
            TemplateSlugs []string  `json:"template_slugs"`
            Logo string  `json:"logo"`
            AppID string  `json:"app_id"`
            Priority float64  `json:"priority"`
            Attributes []AttributeDetailsGroup  `json:"attributes"`
            Slug string  `json:"slug"`
         
    }
    
    // ConfigSuccessResponse used by Catalog
    type ConfigSuccessResponse struct {

        
            Message string  `json:"message"`
         
    }
    
    // AppConfigurationsSort used by Catalog
    type AppConfigurationsSort struct {

        
            IsActive bool  `json:"is_active"`
            IsDefault bool  `json:"is_default"`
            DefaultKey string  `json:"default_key"`
            Name string  `json:"name"`
            Logo string  `json:"logo"`
            AppID string  `json:"app_id"`
            Priority float64  `json:"priority"`
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
    
    // MetaDataListingFilterMetaResponse used by Catalog
    type MetaDataListingFilterMetaResponse struct {

        
            FilterTypes []string  `json:"filter_types"`
            Display string  `json:"display"`
            Units []map[string]interface{}  `json:"units"`
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

        
            Similar map[string]interface{}  `json:"similar"`
            Detail map[string]interface{}  `json:"detail"`
            Variant map[string]interface{}  `json:"variant"`
            Compare map[string]interface{}  `json:"compare"`
         
    }
    
    // GetCatalogConfigurationMetaData used by Catalog
    type GetCatalogConfigurationMetaData struct {

        
            Listing MetaDataListingResponse  `json:"listing"`
            Product GetCatalogConfigurationDetailsProduct  `json:"product"`
         
    }
    
    // ConfigurationBucketPoints used by Catalog
    type ConfigurationBucketPoints struct {

        
            Display string  `json:"display"`
            End float64  `json:"end"`
            Start float64  `json:"start"`
         
    }
    
    // ConfigurationListingFilterValue used by Catalog
    type ConfigurationListingFilterValue struct {

        
            MapValues []map[string]interface{}  `json:"map_values"`
            Map map[string]interface{}  `json:"map"`
            Sort string  `json:"sort"`
            Condition string  `json:"condition"`
            Value string  `json:"value"`
            BucketPoints []ConfigurationBucketPoints  `json:"bucket_points"`
         
    }
    
    // ConfigurationListingFilterConfig used by Catalog
    type ConfigurationListingFilterConfig struct {

        
            DisplayName string  `json:"display_name"`
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            Logo string  `json:"logo"`
            Priority float64  `json:"priority"`
            Type string  `json:"type"`
            Key string  `json:"key"`
            ValueConfig ConfigurationListingFilterValue  `json:"value_config"`
         
    }
    
    // ConfigurationListingFilter used by Catalog
    type ConfigurationListingFilter struct {

        
            AllowSingle bool  `json:"allow_single"`
            AttributeConfig []ConfigurationListingFilterConfig  `json:"attribute_config"`
         
    }
    
    // ConfigurationListingSortConfig used by Catalog
    type ConfigurationListingSortConfig struct {

        
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            Logo string  `json:"logo"`
            Priority float64  `json:"priority"`
            Key string  `json:"key"`
         
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
    
    // ProductSize used by Catalog
    type ProductSize struct {

        
            Max float64  `json:"max"`
            Min float64  `json:"min"`
         
    }
    
    // ConfigurationProductConfig used by Catalog
    type ConfigurationProductConfig struct {

        
            IsActive bool  `json:"is_active"`
            Key string  `json:"key"`
            Size ProductSize  `json:"size"`
            Subtitle string  `json:"subtitle"`
            Logo string  `json:"logo"`
            Priority float64  `json:"priority"`
            Title string  `json:"title"`
         
    }
    
    // ConfigurationProductSimilar used by Catalog
    type ConfigurationProductSimilar struct {

        
            Config []ConfigurationProductConfig  `json:"config"`
         
    }
    
    // ConfigurationProductVariantConfig used by Catalog
    type ConfigurationProductVariantConfig struct {

        
            IsActive bool  `json:"is_active"`
            Size ProductSize  `json:"size"`
            Name string  `json:"name"`
            Logo string  `json:"logo"`
            Priority float64  `json:"priority"`
            DisplayType string  `json:"display_type"`
            Key string  `json:"key"`
         
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
            Product ConfigurationProduct  `json:"product"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ID string  `json:"id"`
            ConfigType string  `json:"config_type"`
            AppID string  `json:"app_id"`
            Type string  `json:"type"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            ModifiedOn string  `json:"modified_on"`
            ConfigID string  `json:"config_id"`
            CreatedOn string  `json:"created_on"`
         
    }
    
    // GetAppCatalogConfiguration used by Catalog
    type GetAppCatalogConfiguration struct {

        
            Data AppCatalogConfiguration  `json:"data"`
            IsDefault bool  `json:"is_default"`
         
    }
    
    // AppConfiguration used by Catalog
    type AppConfiguration struct {

        
            Listing ConfigurationListing  `json:"listing"`
            Product ConfigurationProduct  `json:"product"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ConfigType string  `json:"config_type"`
            AppID string  `json:"app_id"`
            Type string  `json:"type"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            ModifiedOn string  `json:"modified_on"`
            ConfigID string  `json:"config_id"`
            CreatedOn string  `json:"created_on"`
         
    }
    
    // GetCatalogConfigurationDetailsSchemaListing used by Catalog
    type GetCatalogConfigurationDetailsSchemaListing struct {

        
            Filter map[string]interface{}  `json:"filter"`
            Sort map[string]interface{}  `json:"sort"`
         
    }
    
    // EntityConfiguration used by Catalog
    type EntityConfiguration struct {

        
            Listing GetCatalogConfigurationDetailsSchemaListing  `json:"listing"`
            Product GetCatalogConfigurationDetailsProduct  `json:"product"`
            ID string  `json:"id"`
            ConfigType string  `json:"config_type"`
            AppID string  `json:"app_id"`
            ConfigID string  `json:"config_id"`
         
    }
    
    // GetAppCatalogEntityConfiguration used by Catalog
    type GetAppCatalogEntityConfiguration struct {

        
            Data EntityConfiguration  `json:"data"`
            IsDefault bool  `json:"is_default"`
         
    }
    
    // ProductFiltersValue used by Catalog
    type ProductFiltersValue struct {

        
            Max float64  `json:"max"`
            SelectedMin float64  `json:"selected_min"`
            QueryFormat string  `json:"query_format"`
            Min float64  `json:"min"`
            Value interface{}  `json:"value"`
            SelectedMax float64  `json:"selected_max"`
            Display string  `json:"display"`
            IsSelected bool  `json:"is_selected"`
            DisplayFormat string  `json:"display_format"`
            CurrencySymbol string  `json:"currency_symbol"`
            Count float64  `json:"count"`
            CurrencyCode string  `json:"currency_code"`
         
    }
    
    // ProductFiltersKey used by Catalog
    type ProductFiltersKey struct {

        
            Name string  `json:"name"`
            Logo string  `json:"logo"`
            Display string  `json:"display"`
            Kind string  `json:"kind"`
            Operators []string  `json:"operators"`
         
    }
    
    // ProductFilters used by Catalog
    type ProductFilters struct {

        
            Values []ProductFiltersValue  `json:"values"`
            Key ProductFiltersKey  `json:"key"`
         
    }
    
    // ProductSortOn used by Catalog
    type ProductSortOn struct {

        
            Value string  `json:"value"`
            Name string  `json:"name"`
            IsSelected bool  `json:"is_selected"`
         
    }
    
    // GetCollectionQueryOptionResponse used by Catalog
    type GetCollectionQueryOptionResponse struct {

        
            Operators map[string]string  `json:"operators"`
            Filters []ProductFilters  `json:"filters"`
            SortOn []ProductSortOn  `json:"sort_on"`
         
    }
    
    // CollectionListingFilterTag used by Catalog
    type CollectionListingFilterTag struct {

        
            Display string  `json:"display"`
            Name string  `json:"name"`
            IsSelected bool  `json:"is_selected"`
         
    }
    
    // CollectionListingFilterType used by Catalog
    type CollectionListingFilterType struct {

        
            Display string  `json:"display"`
            Name string  `json:"name"`
            IsSelected bool  `json:"is_selected"`
         
    }
    
    // CollectionListingFilter used by Catalog
    type CollectionListingFilter struct {

        
            Tags []CollectionListingFilterTag  `json:"tags"`
            Type []CollectionListingFilterType  `json:"type"`
         
    }
    
    // Media1 used by Catalog
    type Media1 struct {

        
            Type string  `json:"type"`
            URL string  `json:"url"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // CollectionQuery used by Catalog
    type CollectionQuery struct {

        
            Value []interface{}  `json:"value"`
            Op string  `json:"op"`
            Attribute string  `json:"attribute"`
         
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

        
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            Logo Media1  `json:"logo"`
            AppID string  `json:"app_id"`
            Priority float64  `json:"priority"`
            AllowSort bool  `json:"allow_sort"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Type string  `json:"type"`
            AllowFacets bool  `json:"allow_facets"`
            Query []CollectionQuery  `json:"query"`
            Slug string  `json:"slug"`
            Badge map[string]interface{}  `json:"badge"`
            Tag []string  `json:"tag"`
            Banners ImageUrls  `json:"banners"`
            Cron map[string]interface{}  `json:"cron"`
            UID string  `json:"uid"`
            Description string  `json:"description"`
            Schedule map[string]interface{}  `json:"_schedule"`
            Action Action  `json:"action"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // GetCollectionListingResponse used by Catalog
    type GetCollectionListingResponse struct {

        
            Page Page  `json:"page"`
            Filters CollectionListingFilter  `json:"filters"`
            Items []GetCollectionDetailNest  `json:"items"`
         
    }
    
    // CollectionImage used by Catalog
    type CollectionImage struct {

        
            URL string  `json:"url"`
            AspectRatio string  `json:"aspect_ratio"`
         
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
    
    // NextSchedule used by Catalog
    type NextSchedule struct {

        
            End string  `json:"end"`
            Start string  `json:"start"`
         
    }
    
    // CollectionSchedule used by Catalog
    type CollectionSchedule struct {

        
            Start string  `json:"start"`
            Cron string  `json:"cron"`
            End string  `json:"end"`
            Duration float64  `json:"duration"`
            NextSchedule []NextSchedule  `json:"next_schedule"`
         
    }
    
    // CreateCollection used by Catalog
    type CreateCollection struct {

        
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            Logo CollectionImage  `json:"logo"`
            AppID string  `json:"app_id"`
            Priority float64  `json:"priority"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            SortOn string  `json:"sort_on"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            AllowSort bool  `json:"allow_sort"`
            Seo SeoDetail  `json:"seo"`
            Type string  `json:"type"`
            AllowFacets bool  `json:"allow_facets"`
            Query []CollectionQuery  `json:"query"`
            Slug string  `json:"slug"`
            Badge CollectionBadge  `json:"badge"`
            Tags []string  `json:"tags"`
            CreatedBy UserInfo  `json:"created_by"`
            Published bool  `json:"published"`
            Banners CollectionBanner  `json:"banners"`
            Description string  `json:"description"`
            IsVisible bool  `json:"is_visible"`
            Schedule CollectionSchedule  `json:"_schedule"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            ModifiedBy UserInfo  `json:"modified_by"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // CollectionCreateResponse used by Catalog
    type CollectionCreateResponse struct {

        
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            Logo BannerImage  `json:"logo"`
            AppID string  `json:"app_id"`
            Priority float64  `json:"priority"`
            SortOn string  `json:"sort_on"`
            AllowSort bool  `json:"allow_sort"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Type string  `json:"type"`
            AllowFacets bool  `json:"allow_facets"`
            Query []CollectionQuery  `json:"query"`
            Slug string  `json:"slug"`
            Badge map[string]interface{}  `json:"badge"`
            Tag []string  `json:"tag"`
            Banners ImageUrls  `json:"banners"`
            Cron map[string]interface{}  `json:"cron"`
            Description string  `json:"description"`
            Schedule map[string]interface{}  `json:"_schedule"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // CollectionDetailResponse used by Catalog
    type CollectionDetailResponse struct {

        
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            Logo Media1  `json:"logo"`
            AppID string  `json:"app_id"`
            Priority float64  `json:"priority"`
            AllowSort bool  `json:"allow_sort"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Type string  `json:"type"`
            AllowFacets bool  `json:"allow_facets"`
            Query []CollectionQuery  `json:"query"`
            Slug string  `json:"slug"`
            Badge map[string]interface{}  `json:"badge"`
            Tag []string  `json:"tag"`
            Banners ImageUrls  `json:"banners"`
            Cron map[string]interface{}  `json:"cron"`
            UID string  `json:"uid"`
            Description string  `json:"description"`
            Schedule map[string]interface{}  `json:"_schedule"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // UpdateCollection used by Catalog
    type UpdateCollection struct {

        
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            Logo CollectionImage  `json:"logo"`
            Priority float64  `json:"priority"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            SortOn string  `json:"sort_on"`
            AllowSort bool  `json:"allow_sort"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Seo SeoDetail  `json:"seo"`
            Type string  `json:"type"`
            AllowFacets bool  `json:"allow_facets"`
            Query []CollectionQuery  `json:"query"`
            Slug string  `json:"slug"`
            Badge CollectionBadge  `json:"badge"`
            Tags []string  `json:"tags"`
            Published bool  `json:"published"`
            Banners CollectionBanner  `json:"banners"`
            Description string  `json:"description"`
            IsVisible bool  `json:"is_visible"`
            Schedule CollectionSchedule  `json:"_schedule"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            ModifiedBy UserInfo  `json:"modified_by"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // ProductBrand used by Catalog
    type ProductBrand struct {

        
            Name string  `json:"name"`
            Action Action  `json:"action"`
            UID float64  `json:"uid"`
            Logo Media1  `json:"logo"`
         
    }
    
    // Price1 used by Catalog
    type Price1 struct {

        
            Max float64  `json:"max"`
            Min float64  `json:"min"`
            CurrencySymbol string  `json:"currency_symbol"`
            CurrencyCode string  `json:"currency_code"`
         
    }
    
    // ProductListingPrice used by Catalog
    type ProductListingPrice struct {

        
            Marked Price1  `json:"marked"`
            Effective Price1  `json:"effective"`
         
    }
    
    // ProductDetailAttribute used by Catalog
    type ProductDetailAttribute struct {

        
            Value string  `json:"value"`
            Type string  `json:"type"`
            Key string  `json:"key"`
         
    }
    
    // ProductDetailGroupedAttribute used by Catalog
    type ProductDetailGroupedAttribute struct {

        
            Details []ProductDetailAttribute  `json:"details"`
            Title string  `json:"title"`
         
    }
    
    // ProductListingDetail used by Catalog
    type ProductListingDetail struct {

        
            ProductOnlineDate string  `json:"product_online_date"`
            HasVariant bool  `json:"has_variant"`
            Name string  `json:"name"`
            Attributes map[string]interface{}  `json:"attributes"`
            Rating float64  `json:"rating"`
            RatingCount float64  `json:"rating_count"`
            ImageNature string  `json:"image_nature"`
            Sellable bool  `json:"sellable"`
            ShortDescription string  `json:"short_description"`
            Discount string  `json:"discount"`
            Type string  `json:"type"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            Medias []Media1  `json:"medias"`
            PromoMeta map[string]interface{}  `json:"promo_meta"`
            Slug string  `json:"slug"`
            ItemCode string  `json:"item_code"`
            Highlights []string  `json:"highlights"`
            Similars []string  `json:"similars"`
            Color string  `json:"color"`
            UID float64  `json:"uid"`
            Brand ProductBrand  `json:"brand"`
            Description string  `json:"description"`
            Price ProductListingPrice  `json:"price"`
            ItemType string  `json:"item_type"`
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
            Tryouts []string  `json:"tryouts"`
         
    }
    
    // GetCollectionItemsResponse used by Catalog
    type GetCollectionItemsResponse struct {

        
            Page Page  `json:"page"`
            Filters []ProductFilters  `json:"filters"`
            Items []ProductListingDetail  `json:"items"`
            SortOn []ProductSortOn  `json:"sort_on"`
         
    }
    
    // CollectionItem used by Catalog
    type CollectionItem struct {

        
            Priority float64  `json:"priority"`
            Action string  `json:"action"`
            ItemID float64  `json:"item_id"`
         
    }
    
    // CollectionItemUpdate used by Catalog
    type CollectionItemUpdate struct {

        
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            AllowSort bool  `json:"allow_sort"`
            Items []CollectionItem  `json:"items"`
            Type string  `json:"type"`
            Query []CollectionQuery  `json:"query"`
            AllowFacets bool  `json:"allow_facets"`
         
    }
    
    // UpdatedResponse used by Catalog
    type UpdatedResponse struct {

        
            Message string  `json:"message"`
            ItemsNotUpdated []float64  `json:"items_not_updated"`
         
    }
    
    // CatalogInsightItem used by Catalog
    type CatalogInsightItem struct {

        
            SellableCount float64  `json:"sellable_count"`
            OutOfStockCount float64  `json:"out_of_stock_count"`
            Count float64  `json:"count"`
         
    }
    
    // CatalogInsightBrand used by Catalog
    type CatalogInsightBrand struct {

        
            AvailableArticles float64  `json:"available_articles"`
            TotalSizes float64  `json:"total_sizes"`
            AvailableSizes float64  `json:"available_sizes"`
            Name string  `json:"name"`
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

        
            Platform string  `json:"platform"`
            Enabled bool  `json:"enabled"`
            CompanyID float64  `json:"company_id"`
            OptLevel string  `json:"opt_level"`
            BrandIds []float64  `json:"brand_ids"`
            StoreIds []float64  `json:"store_ids"`
         
    }
    
    // CompanyOptIn used by Catalog
    type CompanyOptIn struct {

        
            Platform string  `json:"platform"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Enabled bool  `json:"enabled"`
            CompanyID float64  `json:"company_id"`
            CreatedOn float64  `json:"created_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            OptLevel string  `json:"opt_level"`
            ModifiedOn float64  `json:"modified_on"`
            BrandIds []float64  `json:"brand_ids"`
            StoreIds []float64  `json:"store_ids"`
         
    }
    
    // GetOptInPlatform used by Catalog
    type GetOptInPlatform struct {

        
            Page Page  `json:"page"`
            Items []CompanyOptIn  `json:"items"`
         
    }
    
    // OptinCompanyDetail used by Catalog
    type OptinCompanyDetail struct {

        
            Name string  `json:"name"`
            BusinessType string  `json:"business_type"`
            UID float64  `json:"uid"`
            CompanyType string  `json:"company_type"`
         
    }
    
    // CompanyBrandDetail used by Catalog
    type CompanyBrandDetail struct {

        
            TotalArticle float64  `json:"total_article"`
            CompanyID float64  `json:"company_id"`
            BrandID float64  `json:"brand_id"`
            BrandName string  `json:"brand_name"`
         
    }
    
    // OptinCompanyBrandDetailsView used by Catalog
    type OptinCompanyBrandDetailsView struct {

        
            Page Page  `json:"page"`
            Items []CompanyBrandDetail  `json:"items"`
         
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
            AdditionalContacts []map[string]interface{}  `json:"additional_contacts"`
            StoreType string  `json:"store_type"`
            Documents []map[string]interface{}  `json:"documents"`
            Name string  `json:"name"`
            Timing map[string]interface{}  `json:"timing"`
            CompanyID float64  `json:"company_id"`
            StoreCode string  `json:"store_code"`
            UID float64  `json:"uid"`
            Address map[string]interface{}  `json:"address"`
            Manager map[string]interface{}  `json:"manager"`
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
         
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
            Multi bool  `json:"multi"`
            AllowedValues []string  `json:"allowed_values"`
            Format string  `json:"format"`
            Range AttributeSchemaRange  `json:"range"`
            Type string  `json:"type"`
         
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
            ID string  `json:"id"`
            Name string  `json:"name"`
            Logo string  `json:"logo"`
            Filters AttributeMasterFilter  `json:"filters"`
            Schema AttributeMaster  `json:"schema"`
            Description string  `json:"description"`
            EnabledForEndConsumer bool  `json:"enabled_for_end_consumer"`
            Departments []string  `json:"departments"`
            Slug string  `json:"slug"`
            Meta AttributeMasterMeta  `json:"meta"`
         
    }
    
    // CategoriesResponse used by Catalog
    type CategoriesResponse struct {

        
            Name string  `json:"name"`
            UID float64  `json:"uid"`
            TemplateSlug string  `json:"template_slug"`
            Slug string  `json:"slug"`
            SlugKey string  `json:"slug_key"`
         
    }
    
    // ProdcutTemplateCategoriesResponse used by Catalog
    type ProdcutTemplateCategoriesResponse struct {

        
            Page Page  `json:"page"`
            Items []CategoriesResponse  `json:"items"`
         
    }
    
    // PTErrorResponse used by Catalog
    type PTErrorResponse struct {

        
            Code string  `json:"code"`
            Status float64  `json:"status"`
            Message string  `json:"message"`
            Errors map[string]interface{}  `json:"errors"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // UserSerializer used by Catalog
    type UserSerializer struct {

        
            Username string  `json:"username"`
            Contact string  `json:"contact"`
            UID string  `json:"uid"`
            ID string  `json:"_id"`
            UserID string  `json:"user_id"`
         
    }
    
    // GetDepartment used by Catalog
    type GetDepartment struct {

        
            PriorityOrder float64  `json:"priority_order"`
            IsActive bool  `json:"is_active"`
            Search string  `json:"search"`
            CreatedBy UserSerializer  `json:"created_by"`
            PageNo float64  `json:"page_no"`
            Name string  `json:"name"`
            ItemType string  `json:"item_type"`
            Logo string  `json:"logo"`
            CreatedOn string  `json:"created_on"`
            UID float64  `json:"uid"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            Slug string  `json:"slug"`
            Synonyms []string  `json:"synonyms"`
            ModifiedOn string  `json:"modified_on"`
            PageSize float64  `json:"page_size"`
         
    }
    
    // DepartmentsResponse used by Catalog
    type DepartmentsResponse struct {

        
            Page Page  `json:"page"`
            Items []GetDepartment  `json:"items"`
         
    }
    
    // DepartmentErrorResponse used by Catalog
    type DepartmentErrorResponse struct {

        
            Code string  `json:"code"`
            Status float64  `json:"status"`
            Message string  `json:"message"`
            Errors map[string]interface{}  `json:"errors"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // DepartmentCreateUpdate used by Catalog
    type DepartmentCreateUpdate struct {

        
            PriorityOrder float64  `json:"priority_order"`
            IsActive bool  `json:"is_active"`
            Platforms map[string]interface{}  `json:"platforms"`
            Name string  `json:"name"`
            Logo string  `json:"logo"`
            UID float64  `json:"uid"`
            Cls string  `json:"_cls"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Tags []string  `json:"tags"`
            Slug string  `json:"slug"`
            Synonyms []string  `json:"synonyms"`
         
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
            Contact string  `json:"contact"`
            SuperUser bool  `json:"super_user"`
            UserID string  `json:"user_id"`
         
    }
    
    // DepartmentModel used by Catalog
    type DepartmentModel struct {

        
            PriorityOrder float64  `json:"priority_order"`
            CreatedBy UserDetail  `json:"created_by"`
            IsActive bool  `json:"is_active"`
            VerifiedOn string  `json:"verified_on"`
            Name string  `json:"name"`
            VerifiedBy UserDetail  `json:"verified_by"`
            Logo string  `json:"logo"`
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
            ModifiedBy UserDetail  `json:"modified_by"`
            Cls string  `json:"_cls"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ID interface{}  `json:"_id"`
            Synonyms []string  `json:"synonyms"`
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
         
    }
    
    // ProductTemplate used by Catalog
    type ProductTemplate struct {

        
            IsActive bool  `json:"is_active"`
            IsExpirable bool  `json:"is_expirable"`
            Tag string  `json:"tag"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Categories []string  `json:"categories"`
            IsPhysical bool  `json:"is_physical"`
            Name string  `json:"name"`
            Logo string  `json:"logo"`
            CreatedOn string  `json:"created_on"`
            Attributes []string  `json:"attributes"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Departments []string  `json:"departments"`
            IsArchived bool  `json:"is_archived"`
            Slug string  `json:"slug"`
            ModifiedOn string  `json:"modified_on"`
            Description string  `json:"description"`
         
    }
    
    // TemplatesResponse used by Catalog
    type TemplatesResponse struct {

        
            Page Page  `json:"page"`
            Items ProductTemplate  `json:"items"`
         
    }
    
    // Properties used by Catalog
    type Properties struct {

        
            ReturnConfig map[string]interface{}  `json:"return_config"`
            IsActive map[string]interface{}  `json:"is_active"`
            Name map[string]interface{}  `json:"name"`
            CategorySlug map[string]interface{}  `json:"category_slug"`
            Media map[string]interface{}  `json:"media"`
            Command map[string]interface{}  `json:"command"`
            Trader map[string]interface{}  `json:"trader"`
            Sizes map[string]interface{}  `json:"sizes"`
            SizeGuide map[string]interface{}  `json:"size_guide"`
            ShortDescription map[string]interface{}  `json:"short_description"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            MultiSize map[string]interface{}  `json:"multi_size"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            Slug map[string]interface{}  `json:"slug"`
            Tags map[string]interface{}  `json:"tags"`
            ItemCode map[string]interface{}  `json:"item_code"`
            Highlights map[string]interface{}  `json:"highlights"`
            NoOfBoxes map[string]interface{}  `json:"no_of_boxes"`
            BrandUID map[string]interface{}  `json:"brand_uid"`
            ProductPublish map[string]interface{}  `json:"product_publish"`
            Currency map[string]interface{}  `json:"currency"`
            IsDependent map[string]interface{}  `json:"is_dependent"`
            TraderType map[string]interface{}  `json:"trader_type"`
            Description map[string]interface{}  `json:"description"`
            CountryOfOrigin map[string]interface{}  `json:"country_of_origin"`
            HsnCode map[string]interface{}  `json:"hsn_code"`
            Variants map[string]interface{}  `json:"variants"`
            ItemType map[string]interface{}  `json:"item_type"`
            ProductGroupTag map[string]interface{}  `json:"product_group_tag"`
         
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
    
    // TemplateDetails used by Catalog
    type TemplateDetails struct {

        
            IsActive bool  `json:"is_active"`
            IsExpirable bool  `json:"is_expirable"`
            Tag string  `json:"tag"`
            ID string  `json:"id"`
            Categories []string  `json:"categories"`
            IsPhysical bool  `json:"is_physical"`
            Name string  `json:"name"`
            Logo string  `json:"logo"`
            Attributes []string  `json:"attributes"`
            Departments []string  `json:"departments"`
            IsArchived bool  `json:"is_archived"`
            Slug string  `json:"slug"`
            Description string  `json:"description"`
         
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

        
            Username string  `json:"username"`
            Email string  `json:"email"`
            UserID string  `json:"user_id"`
            UID string  `json:"uid"`
         
    }
    
    // ProductTemplateExportResponse used by Catalog
    type ProductTemplateExportResponse struct {

        
            CreatedBy UserInfo1  `json:"created_by"`
            ModifiedOn string  `json:"modified_on"`
            Status string  `json:"status"`
            TaskID string  `json:"task_id"`
            Filters map[string]interface{}  `json:"filters"`
            Type string  `json:"type"`
            URL string  `json:"url"`
            CompletedOn string  `json:"completed_on"`
            NotificationEmails []string  `json:"notification_emails"`
            SellerID float64  `json:"seller_id"`
         
    }
    
    // ProductDownloadsResponse used by Catalog
    type ProductDownloadsResponse struct {

        
            Items []ProductTemplateExportResponse  `json:"items"`
         
    }
    
    // ProductTemplateExportFilterRequest used by Catalog
    type ProductTemplateExportFilterRequest struct {

        
            FromDate string  `json:"from_date"`
            CatalogueTypes []string  `json:"catalogue_types"`
            Templates []string  `json:"templates"`
            Brands []string  `json:"brands"`
            ToDate string  `json:"to_date"`
         
    }
    
    // ProductTemplateDownloadsExport used by Catalog
    type ProductTemplateDownloadsExport struct {

        
            Filters ProductTemplateExportFilterRequest  `json:"filters"`
            Type string  `json:"type"`
            NotificationEmails []string  `json:"notification_emails"`
         
    }
    
    // ProductConfigurationDownloads used by Catalog
    type ProductConfigurationDownloads struct {

        
            Data []map[string]interface{}  `json:"data"`
            Multivalue bool  `json:"multivalue"`
         
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
    
    // Hierarchy used by Catalog
    type Hierarchy struct {

        
            Department float64  `json:"department"`
            L2 float64  `json:"l2"`
            L1 float64  `json:"l1"`
         
    }
    
    // Media2 used by Catalog
    type Media2 struct {

        
            Portrait string  `json:"portrait"`
            Landscape string  `json:"landscape"`
            Logo string  `json:"logo"`
         
    }
    
    // Category used by Catalog
    type Category struct {

        
            Marketplaces CategoryMapping  `json:"marketplaces"`
            IsActive bool  `json:"is_active"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ID string  `json:"id"`
            Hierarchy []Hierarchy  `json:"hierarchy"`
            Level float64  `json:"level"`
            Name string  `json:"name"`
            Priority float64  `json:"priority"`
            UID float64  `json:"uid"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Media Media2  `json:"media"`
            Departments []float64  `json:"departments"`
            Tryouts []string  `json:"tryouts"`
            Slug string  `json:"slug"`
            Synonyms []string  `json:"synonyms"`
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
         
    }
    
    // CategoryResponse used by Catalog
    type CategoryResponse struct {

        
            Page Page  `json:"page"`
            Items []Category  `json:"items"`
         
    }
    
    // CategoryRequestBody used by Catalog
    type CategoryRequestBody struct {

        
            Marketplaces CategoryMapping  `json:"marketplaces"`
            IsActive bool  `json:"is_active"`
            Hierarchy []Hierarchy  `json:"hierarchy"`
            Level float64  `json:"level"`
            Name string  `json:"name"`
            Priority float64  `json:"priority"`
            Media Media2  `json:"media"`
            Departments []float64  `json:"departments"`
            Tryouts []string  `json:"tryouts"`
            Slug string  `json:"slug"`
            Synonyms []string  `json:"synonyms"`
         
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
    
    // ReturnConfigResponse used by Catalog
    type ReturnConfigResponse struct {

        
            Returnable bool  `json:"returnable"`
            Unit string  `json:"unit"`
            Time float64  `json:"time"`
         
    }
    
    // Trader used by Catalog
    type Trader struct {

        
            Address []string  `json:"address"`
            Name string  `json:"name"`
            Type string  `json:"type"`
         
    }
    
    // NetQuantityResponse used by Catalog
    type NetQuantityResponse struct {

        
            Value float64  `json:"value"`
            Unit string  `json:"unit"`
         
    }
    
    // Image used by Catalog
    type Image struct {

        
            AspectRatioF float64  `json:"aspect_ratio_f"`
            URL string  `json:"url"`
            AspectRatio string  `json:"aspect_ratio"`
            SecureURL string  `json:"secure_url"`
         
    }
    
    // ProductPublish used by Catalog
    type ProductPublish struct {

        
            ProductOnlineDate string  `json:"product_online_date"`
            IsSet bool  `json:"is_set"`
         
    }
    
    // VerifiedBy used by Catalog
    type VerifiedBy struct {

        
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
         
    }
    
    // Logo used by Catalog
    type Logo struct {

        
            AspectRatioF float64  `json:"aspect_ratio_f"`
            URL string  `json:"url"`
            AspectRatio string  `json:"aspect_ratio"`
            SecureURL string  `json:"secure_url"`
         
    }
    
    // Brand used by Catalog
    type Brand struct {

        
            Name string  `json:"name"`
            UID float64  `json:"uid"`
            Logo Logo  `json:"logo"`
         
    }
    
    // ProductSchemaV2 used by Catalog
    type ProductSchemaV2 struct {

        
            ReturnConfig ReturnConfigResponse  `json:"return_config"`
            IsActive bool  `json:"is_active"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            VariantGroup map[string]interface{}  `json:"variant_group"`
            AllIdentifiers []string  `json:"all_identifiers"`
            Name string  `json:"name"`
            CategorySlug string  `json:"category_slug"`
            Attributes map[string]interface{}  `json:"attributes"`
            Media []Media1  `json:"media"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Trader []Trader  `json:"trader"`
            ImageNature string  `json:"image_nature"`
            Sizes []map[string]interface{}  `json:"sizes"`
            SizeGuide string  `json:"size_guide"`
            CategoryUID float64  `json:"category_uid"`
            ID string  `json:"id"`
            Pending string  `json:"pending"`
            Category map[string]interface{}  `json:"category"`
            ShortDescription string  `json:"short_description"`
            AllSizes []map[string]interface{}  `json:"all_sizes"`
            CompanyID float64  `json:"company_id"`
            NetQuantity NetQuantityResponse  `json:"net_quantity"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            MultiSize bool  `json:"multi_size"`
            Images []Image  `json:"images"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            Slug string  `json:"slug"`
            Tags []string  `json:"tags"`
            CreatedOn string  `json:"created_on"`
            AllCompanyIds []float64  `json:"all_company_ids"`
            IsSet bool  `json:"is_set"`
            ItemCode string  `json:"item_code"`
            TemplateTag string  `json:"template_tag"`
            Highlights []string  `json:"highlights"`
            IsExpirable bool  `json:"is_expirable"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            NoOfBoxes float64  `json:"no_of_boxes"`
            BrandUID float64  `json:"brand_uid"`
            Color string  `json:"color"`
            ProductPublish ProductPublish  `json:"product_publish"`
            Stage string  `json:"stage"`
            PrimaryColor string  `json:"primary_color"`
            Currency string  `json:"currency"`
            UID float64  `json:"uid"`
            IsDependent bool  `json:"is_dependent"`
            L3Mapping []string  `json:"l3_mapping"`
            ModifiedOn string  `json:"modified_on"`
            Description string  `json:"description"`
            IsImageLessProduct bool  `json:"is_image_less_product"`
            CountryOfOrigin string  `json:"country_of_origin"`
            HsnCode string  `json:"hsn_code"`
            IsPhysical bool  `json:"is_physical"`
            Variants map[string]interface{}  `json:"variants"`
            ItemType string  `json:"item_type"`
            Moq map[string]interface{}  `json:"moq"`
            VerifiedBy VerifiedBy  `json:"verified_by"`
            VerifiedOn string  `json:"verified_on"`
            ProductGroupTag []string  `json:"product_group_tag"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Departments []float64  `json:"departments"`
            Brand Brand  `json:"brand"`
            VariantMedia map[string]interface{}  `json:"variant_media"`
         
    }
    
    // ProductListingResponseV2 used by Catalog
    type ProductListingResponseV2 struct {

        
            Page Page  `json:"page"`
            Items []ProductSchemaV2  `json:"items"`
         
    }
    
    // ReturnConfig used by Catalog
    type ReturnConfig struct {

        
            Returnable bool  `json:"returnable"`
            Unit string  `json:"unit"`
            Time float64  `json:"time"`
         
    }
    
    // TaxIdentifier used by Catalog
    type TaxIdentifier struct {

        
            ReportingHsn string  `json:"reporting_hsn"`
            HsnCodeID string  `json:"hsn_code_id"`
            HsnCode string  `json:"hsn_code"`
         
    }
    
    // NetQuantity used by Catalog
    type NetQuantity struct {

        
            Value float64  `json:"value"`
            Unit string  `json:"unit"`
         
    }
    
    // CustomOrder used by Catalog
    type CustomOrder struct {

        
            ManufacturingTime float64  `json:"manufacturing_time"`
            IsCustomOrder bool  `json:"is_custom_order"`
            ManufacturingTimeUnit string  `json:"manufacturing_time_unit"`
         
    }
    
    // TeaserTag used by Catalog
    type TeaserTag struct {

        
            Tag string  `json:"tag"`
            URL string  `json:"url"`
         
    }
    
    // ProductPublish1 used by Catalog
    type ProductPublish1 struct {

        
            ProductOnlineDate string  `json:"product_online_date"`
            IsSet bool  `json:"is_set"`
         
    }
    
    // ProductCreateUpdateSchemaV2 used by Catalog
    type ProductCreateUpdateSchemaV2 struct {

        
            ReturnConfig ReturnConfig  `json:"return_config"`
            IsActive bool  `json:"is_active"`
            TaxIdentifier TaxIdentifier  `json:"tax_identifier"`
            VariantGroup map[string]interface{}  `json:"variant_group"`
            Name string  `json:"name"`
            CategorySlug string  `json:"category_slug"`
            Attributes map[string]interface{}  `json:"attributes"`
            Media []Media1  `json:"media"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Trader []Trader  `json:"trader"`
            Sizes []map[string]interface{}  `json:"sizes"`
            SizeGuide string  `json:"size_guide"`
            ShortDescription string  `json:"short_description"`
            CompanyID float64  `json:"company_id"`
            NetQuantity NetQuantity  `json:"net_quantity"`
            CustomOrder CustomOrder  `json:"custom_order"`
            MultiSize bool  `json:"multi_size"`
            BulkJobID string  `json:"bulk_job_id"`
            TeaserTag TeaserTag  `json:"teaser_tag"`
            Slug string  `json:"slug"`
            Tags []string  `json:"tags"`
            ChangeRequestID interface{}  `json:"change_request_id"`
            ItemCode string  `json:"item_code"`
            IsSet bool  `json:"is_set"`
            TemplateTag string  `json:"template_tag"`
            Highlights []string  `json:"highlights"`
            NoOfBoxes float64  `json:"no_of_boxes"`
            BrandUID float64  `json:"brand_uid"`
            ProductPublish ProductPublish1  `json:"product_publish"`
            Currency string  `json:"currency"`
            UID float64  `json:"uid"`
            Requester string  `json:"requester"`
            IsDependent bool  `json:"is_dependent"`
            Description string  `json:"description"`
            IsImageLessProduct bool  `json:"is_image_less_product"`
            CountryOfOrigin string  `json:"country_of_origin"`
            Variants map[string]interface{}  `json:"variants"`
            ItemType string  `json:"item_type"`
            ProductGroupTag []string  `json:"product_group_tag"`
            Departments []float64  `json:"departments"`
            Action string  `json:"action"`
            VariantMedia map[string]interface{}  `json:"variant_media"`
         
    }
    
    // ProductVariants used by Catalog
    type ProductVariants struct {

        
            ItemCode string  `json:"item_code"`
            CategoryUID float64  `json:"category_uid"`
            Name string  `json:"name"`
            BrandUID float64  `json:"brand_uid"`
            UID float64  `json:"uid"`
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
            Suggestion string  `json:"suggestion"`
            Name string  `json:"name"`
            Logo string  `json:"logo"`
            Unit string  `json:"unit"`
            EnabledForEndConsumer bool  `json:"enabled_for_end_consumer"`
            Details AttributeMasterDetails  `json:"details"`
            Variant bool  `json:"variant"`
            Filters AttributeMasterFilter  `json:"filters"`
            Tags []string  `json:"tags"`
            Slug string  `json:"slug"`
            Synonyms map[string]interface{}  `json:"synonyms"`
            CreatedOn string  `json:"created_on"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Schema AttributeMaster  `json:"schema"`
            ModifiedOn string  `json:"modified_on"`
            Description string  `json:"description"`
            RawKey string  `json:"raw_key"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Departments []string  `json:"departments"`
         
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

        
            ItemLength float64  `json:"item_length"`
            Size string  `json:"size"`
            ItemWeightUnitOfMeasure string  `json:"item_weight_unit_of_measure"`
            ItemWeight float64  `json:"item_weight"`
            Identifiers []ValidateIdentifier  `json:"identifiers"`
            ItemHeight float64  `json:"item_height"`
            ItemDimensionsUnitOfMeasure string  `json:"item_dimensions_unit_of_measure"`
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

        
            ReturnConfig ReturnConfigResponse  `json:"return_config"`
            IsActive bool  `json:"is_active"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            VariantGroup map[string]interface{}  `json:"variant_group"`
            AllIdentifiers []string  `json:"all_identifiers"`
            Name string  `json:"name"`
            CategorySlug string  `json:"category_slug"`
            Attributes map[string]interface{}  `json:"attributes"`
            Media []Media1  `json:"media"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Trader []Trader  `json:"trader"`
            ImageNature string  `json:"image_nature"`
            Sizes []map[string]interface{}  `json:"sizes"`
            SizeGuide string  `json:"size_guide"`
            CategoryUID float64  `json:"category_uid"`
            ID string  `json:"id"`
            Pending string  `json:"pending"`
            Category map[string]interface{}  `json:"category"`
            ShortDescription string  `json:"short_description"`
            AllSizes []map[string]interface{}  `json:"all_sizes"`
            CompanyID float64  `json:"company_id"`
            NetQuantity NetQuantityResponse  `json:"net_quantity"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            MultiSize bool  `json:"multi_size"`
            Images []Image  `json:"images"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            Slug string  `json:"slug"`
            Tags []string  `json:"tags"`
            CreatedOn string  `json:"created_on"`
            AllCompanyIds []float64  `json:"all_company_ids"`
            IsSet bool  `json:"is_set"`
            ItemCode string  `json:"item_code"`
            TemplateTag string  `json:"template_tag"`
            Highlights []string  `json:"highlights"`
            IsExpirable bool  `json:"is_expirable"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            NoOfBoxes float64  `json:"no_of_boxes"`
            BrandUID float64  `json:"brand_uid"`
            Color string  `json:"color"`
            ProductPublish ProductPublished  `json:"product_publish"`
            Stage string  `json:"stage"`
            PrimaryColor string  `json:"primary_color"`
            Currency string  `json:"currency"`
            UID float64  `json:"uid"`
            IsDependent bool  `json:"is_dependent"`
            L3Mapping []string  `json:"l3_mapping"`
            ModifiedOn string  `json:"modified_on"`
            Description string  `json:"description"`
            IsImageLessProduct bool  `json:"is_image_less_product"`
            CountryOfOrigin string  `json:"country_of_origin"`
            HsnCode string  `json:"hsn_code"`
            IsPhysical bool  `json:"is_physical"`
            Variants map[string]interface{}  `json:"variants"`
            ItemType string  `json:"item_type"`
            Moq map[string]interface{}  `json:"moq"`
            VerifiedBy VerifiedBy  `json:"verified_by"`
            VerifiedOn string  `json:"verified_on"`
            ProductGroupTag []string  `json:"product_group_tag"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Departments []float64  `json:"departments"`
            Brand Brand  `json:"brand"`
            VariantMedia map[string]interface{}  `json:"variant_media"`
         
    }
    
    // ProductListingResponse used by Catalog
    type ProductListingResponse struct {

        
            Page Page  `json:"page"`
            Items []Product  `json:"items"`
         
    }
    
    // UserDetail1 used by Catalog
    type UserDetail1 struct {

        
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
            FullName string  `json:"full_name"`
         
    }
    
    // ProductBulkRequest used by Catalog
    type ProductBulkRequest struct {

        
            IsActive bool  `json:"is_active"`
            CreatedBy UserDetail1  `json:"created_by"`
            TemplateTag string  `json:"template_tag"`
            Cancelled float64  `json:"cancelled"`
            CompanyID float64  `json:"company_id"`
            Stage string  `json:"stage"`
            Total float64  `json:"total"`
            FilePath string  `json:"file_path"`
            Succeed float64  `json:"succeed"`
            ModifiedBy UserDetail1  `json:"modified_by"`
            CancelledRecords []string  `json:"cancelled_records"`
            Template ProductTemplate  `json:"template"`
            FailedRecords []string  `json:"failed_records"`
            Failed float64  `json:"failed"`
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
         
    }
    
    // ProductBulkRequestList used by Catalog
    type ProductBulkRequestList struct {

        
            Page Page  `json:"page"`
            Items ProductBulkRequest  `json:"items"`
         
    }
    
    // BulkJob used by Catalog
    type BulkJob struct {

        
            TrackingURL string  `json:"tracking_url"`
            IsActive bool  `json:"is_active"`
            CreatedBy UserInfo1  `json:"created_by"`
            Cancelled float64  `json:"cancelled"`
            TemplateTag string  `json:"template_tag"`
            CustomTemplateTag string  `json:"custom_template_tag"`
            CompanyID float64  `json:"company_id"`
            Stage string  `json:"stage"`
            Total float64  `json:"total"`
            Succeed float64  `json:"succeed"`
            FilePath string  `json:"file_path"`
            ModifiedBy UserInfo1  `json:"modified_by"`
            CancelledRecords []map[string]interface{}  `json:"cancelled_records"`
            FailedRecords []map[string]interface{}  `json:"failed_records"`
            Failed float64  `json:"failed"`
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
         
    }
    
    // BulkResponse used by Catalog
    type BulkResponse struct {

        
            IsActive bool  `json:"is_active"`
            CreatedBy UserInfo1  `json:"created_by"`
            ModifiedBy UserInfo1  `json:"modified_by"`
            BatchID string  `json:"batch_id"`
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
         
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
    
    // UserCommon used by Catalog
    type UserCommon struct {

        
            Username string  `json:"username"`
            CompanyID float64  `json:"company_id"`
            UserID string  `json:"user_id"`
         
    }
    
    // Items used by Catalog
    type Items struct {

        
            TrackingURL string  `json:"tracking_url"`
            CreatedBy UserCommon  `json:"created_by"`
            IsActive bool  `json:"is_active"`
            Retry float64  `json:"retry"`
            Cancelled float64  `json:"cancelled"`
            ID string  `json:"id"`
            CompanyID float64  `json:"company_id"`
            Stage string  `json:"stage"`
            FilePath string  `json:"file_path"`
            Succeed float64  `json:"succeed"`
            Total float64  `json:"total"`
            ModifiedBy UserCommon  `json:"modified_by"`
            CancelledRecords []string  `json:"cancelled_records"`
            FailedRecords []string  `json:"failed_records"`
            Failed float64  `json:"failed"`
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
         
    }
    
    // BulkAssetResponse used by Catalog
    type BulkAssetResponse struct {

        
            Page Page  `json:"page"`
            Items []Items  `json:"items"`
         
    }
    
    // ProductBulkAssets used by Catalog
    type ProductBulkAssets struct {

        
            User map[string]interface{}  `json:"user"`
            CompanyID float64  `json:"company_id"`
            URL string  `json:"url"`
         
    }
    
    // ProductSizeDeleteDataResponse used by Catalog
    type ProductSizeDeleteDataResponse struct {

        
            Size string  `json:"size"`
            CompanyID float64  `json:"company_id"`
            ItemID float64  `json:"item_id"`
         
    }
    
    // ProductSizeDeleteResponse used by Catalog
    type ProductSizeDeleteResponse struct {

        
            Data ProductSizeDeleteDataResponse  `json:"data"`
            Success bool  `json:"success"`
         
    }
    
    // InventoryResponse used by Catalog
    type InventoryResponse struct {

        
            Price float64  `json:"price"`
            Size string  `json:"size"`
            Store map[string]interface{}  `json:"store"`
            ItemID float64  `json:"item_id"`
            InventoryUpdatedOn string  `json:"inventory_updated_on"`
            Currency string  `json:"currency"`
            UID string  `json:"uid"`
            Identifiers map[string]interface{}  `json:"identifiers"`
            SellableQuantity float64  `json:"sellable_quantity"`
            SellerIdentifier string  `json:"seller_identifier"`
            PriceEffective float64  `json:"price_effective"`
            Quantity float64  `json:"quantity"`
            PriceTransfer float64  `json:"price_transfer"`
         
    }
    
    // InventoryResponsePaginated used by Catalog
    type InventoryResponsePaginated struct {

        
            Page Page  `json:"page"`
            Items []InventoryResponse  `json:"items"`
         
    }
    
    // ItemQuery used by Catalog
    type ItemQuery struct {

        
            ItemCode string  `json:"item_code"`
            UID float64  `json:"uid"`
            BrandUID float64  `json:"brand_uid"`
         
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
            GtinType string  `json:"gtin_type"`
            GtinValue string  `json:"gtin_value"`
         
    }
    
    // InvSize used by Catalog
    type InvSize struct {

        
            ExpirationDate string  `json:"expiration_date"`
            IsSet bool  `json:"is_set"`
            ItemLength float64  `json:"item_length"`
            Price float64  `json:"price"`
            Size string  `json:"size"`
            ItemWeightUnitOfMeasure string  `json:"item_weight_unit_of_measure"`
            ItemWeight float64  `json:"item_weight"`
            Set InventorySet  `json:"set"`
            StoreCode string  `json:"store_code"`
            Currency string  `json:"currency"`
            Identifiers []GTIN  `json:"identifiers"`
            ItemHeight float64  `json:"item_height"`
            ItemDimensionsUnitOfMeasure string  `json:"item_dimensions_unit_of_measure"`
            PriceEffective float64  `json:"price_effective"`
            Quantity float64  `json:"quantity"`
            ItemWidth float64  `json:"item_width"`
            PriceTransfer float64  `json:"price_transfer"`
         
    }
    
    // InventoryRequest used by Catalog
    type InventoryRequest struct {

        
            Item ItemQuery  `json:"item"`
            CompanyID float64  `json:"company_id"`
            Sizes []InvSize  `json:"sizes"`
         
    }
    
    // ReturnConfig1 used by Catalog
    type ReturnConfig1 struct {

        
            Returnable bool  `json:"returnable"`
            Unit string  `json:"unit"`
            Time float64  `json:"time"`
         
    }
    
    // ManufacturerResponse used by Catalog
    type ManufacturerResponse struct {

        
            Address string  `json:"address"`
            Name string  `json:"name"`
            IsDefault bool  `json:"is_default"`
         
    }
    
    // StoreMeta used by Catalog
    type StoreMeta struct {

        
            ID float64  `json:"id"`
         
    }
    
    // DimensionResponse used by Catalog
    type DimensionResponse struct {

        
            Length float64  `json:"length"`
            IsDefault bool  `json:"is_default"`
            Height float64  `json:"height"`
            Width float64  `json:"width"`
            Unit string  `json:"unit"`
         
    }
    
    // Trader1 used by Catalog
    type Trader1 struct {

        
            Address []string  `json:"address"`
            Name string  `json:"name"`
            Type string  `json:"type"`
         
    }
    
    // WeightResponse used by Catalog
    type WeightResponse struct {

        
            Shipping float64  `json:"shipping"`
            IsDefault bool  `json:"is_default"`
            Unit string  `json:"unit"`
         
    }
    
    // CompanyMeta used by Catalog
    type CompanyMeta struct {

        
            ID float64  `json:"id"`
         
    }
    
    // QuantityBase used by Catalog
    type QuantityBase struct {

        
            UpdatedAt string  `json:"updated_at"`
            Count float64  `json:"count"`
         
    }
    
    // Quantities used by Catalog
    type Quantities struct {

        
            Sellable QuantityBase  `json:"sellable"`
            NotAvailable QuantityBase  `json:"not_available"`
            OrderCommitted QuantityBase  `json:"order_committed"`
            Damaged QuantityBase  `json:"damaged"`
         
    }
    
    // PriceMeta used by Catalog
    type PriceMeta struct {

        
            Marked float64  `json:"marked"`
            Effective float64  `json:"effective"`
            TpNotes map[string]interface{}  `json:"tp_notes"`
            UpdatedAt string  `json:"updated_at"`
            Transfer float64  `json:"transfer"`
            Currency string  `json:"currency"`
         
    }
    
    // BrandMeta used by Catalog
    type BrandMeta struct {

        
            Name string  `json:"name"`
            ID float64  `json:"id"`
         
    }
    
    // InventorySellerResponse used by Catalog
    type InventorySellerResponse struct {

        
            ReturnConfig ReturnConfig1  `json:"return_config"`
            IsActive bool  `json:"is_active"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            Manufacturer ManufacturerResponse  `json:"manufacturer"`
            Store StoreMeta  `json:"store"`
            Fragile bool  `json:"fragile"`
            TraceID string  `json:"trace_id"`
            Dimension DimensionResponse  `json:"dimension"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            SellerIdentifier string  `json:"seller_identifier"`
            RawMeta map[string]interface{}  `json:"raw_meta"`
            Trader []Trader1  `json:"trader"`
            AddedOnStore string  `json:"added_on_store"`
            Meta map[string]interface{}  `json:"meta"`
            TrackInventory bool  `json:"track_inventory"`
            Weight WeightResponse  `json:"weight"`
            Tags []string  `json:"tags"`
            IsSet bool  `json:"is_set"`
            CreatedBy UserSerializer  `json:"created_by"`
            Size string  `json:"size"`
            Stage string  `json:"stage"`
            ItemID float64  `json:"item_id"`
            Set InventorySet  `json:"set"`
            UID string  `json:"uid"`
            Company CompanyMeta  `json:"company"`
            CountryOfOrigin string  `json:"country_of_origin"`
            ExpirationDate string  `json:"expiration_date"`
            Quantities Quantities  `json:"quantities"`
            TotalQuantity float64  `json:"total_quantity"`
            Price PriceMeta  `json:"price"`
            Identifier map[string]interface{}  `json:"identifier"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            FyndArticleCode string  `json:"fynd_article_code"`
            Brand BrandMeta  `json:"brand"`
            FyndItemCode string  `json:"fynd_item_code"`
            FyndMeta map[string]interface{}  `json:"fynd_meta"`
         
    }
    
    // InventorySellerIdentifierResponsePaginated used by Catalog
    type InventorySellerIdentifierResponsePaginated struct {

        
            Page Page  `json:"page"`
            Items []InventorySellerResponse  `json:"items"`
         
    }
    
    // ReturnConfig2 used by Catalog
    type ReturnConfig2 struct {

        
            Returnable bool  `json:"returnable"`
            Unit string  `json:"unit"`
            Time float64  `json:"time"`
         
    }
    
    // ManufacturerResponse1 used by Catalog
    type ManufacturerResponse1 struct {

        
            Address string  `json:"address"`
            Name string  `json:"name"`
            IsDefault bool  `json:"is_default"`
         
    }
    
    // ArticleStoreResponse used by Catalog
    type ArticleStoreResponse struct {

        
            StoreCode string  `json:"store_code"`
            Name string  `json:"name"`
            UID float64  `json:"uid"`
            StoreType string  `json:"store_type"`
         
    }
    
    // DimensionResponse1 used by Catalog
    type DimensionResponse1 struct {

        
            Length float64  `json:"length"`
            Height float64  `json:"height"`
            Unit string  `json:"unit"`
            Width float64  `json:"width"`
         
    }
    
    // Trader2 used by Catalog
    type Trader2 struct {

        
            Address []string  `json:"address"`
            Name string  `json:"name"`
            Type string  `json:"type"`
         
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
    
    // Quantity used by Catalog
    type Quantity struct {

        
            Count float64  `json:"count"`
         
    }
    
    // QuantitiesArticle used by Catalog
    type QuantitiesArticle struct {

        
            Sellable Quantity  `json:"sellable"`
            NotAvailable Quantity  `json:"not_available"`
            OrderCommitted Quantity  `json:"order_committed"`
            Damaged Quantity  `json:"damaged"`
         
    }
    
    // PriceArticle used by Catalog
    type PriceArticle struct {

        
            Marked float64  `json:"marked"`
            Effective float64  `json:"effective"`
            TpNotes map[string]interface{}  `json:"tp_notes"`
            Transfer float64  `json:"transfer"`
            Currency string  `json:"currency"`
         
    }
    
    // DateMeta used by Catalog
    type DateMeta struct {

        
            InventoryUpdatedOn string  `json:"inventory_updated_on"`
            ModifiedOn string  `json:"modified_on"`
            AddedOnStore string  `json:"added_on_store"`
            CreatedOn string  `json:"created_on"`
         
    }
    
    // BrandMeta1 used by Catalog
    type BrandMeta1 struct {

        
            Name string  `json:"name"`
            ID float64  `json:"id"`
         
    }
    
    // GetInventories used by Catalog
    type GetInventories struct {

        
            ReturnConfig ReturnConfig2  `json:"return_config"`
            Manufacturer ManufacturerResponse1  `json:"manufacturer"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            Store ArticleStoreResponse  `json:"store"`
            TraceID string  `json:"trace_id"`
            Dimension DimensionResponse1  `json:"dimension"`
            InventoryUpdatedOn string  `json:"inventory_updated_on"`
            SellerIdentifier string  `json:"seller_identifier"`
            Trader []Trader2  `json:"trader"`
            ID string  `json:"id"`
            TrackInventory bool  `json:"track_inventory"`
            Weight WeightResponse1  `json:"weight"`
            Tags []string  `json:"tags"`
            IsSet bool  `json:"is_set"`
            CreatedBy UserSerializer  `json:"created_by"`
            Size string  `json:"size"`
            Platforms map[string]interface{}  `json:"platforms"`
            Stage string  `json:"stage"`
            ItemID float64  `json:"item_id"`
            UID string  `json:"uid"`
            Company CompanyMeta1  `json:"company"`
            Quantities QuantitiesArticle  `json:"quantities"`
            ExpirationDate string  `json:"expiration_date"`
            CountryOfOrigin string  `json:"country_of_origin"`
            TotalQuantity float64  `json:"total_quantity"`
            Price PriceArticle  `json:"price"`
            DateMeta DateMeta  `json:"date_meta"`
            Identifier map[string]interface{}  `json:"identifier"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            Brand BrandMeta1  `json:"brand"`
         
    }
    
    // GetInventoriesResponse used by Catalog
    type GetInventoriesResponse struct {

        
            Page Page  `json:"page"`
            Items []GetInventories  `json:"items"`
         
    }
    
    // BulkInventoryGetItems used by Catalog
    type BulkInventoryGetItems struct {

        
            CreatedBy map[string]interface{}  `json:"created_by"`
            Cancelled float64  `json:"cancelled"`
            IsActive bool  `json:"is_active"`
            ID string  `json:"id"`
            CompanyID float64  `json:"company_id"`
            Stage string  `json:"stage"`
            FilePath string  `json:"file_path"`
            Succeed float64  `json:"succeed"`
            Total float64  `json:"total"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            CancelledRecords []string  `json:"cancelled_records"`
            FailedRecords []string  `json:"failed_records"`
            Failed float64  `json:"failed"`
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
         
    }
    
    // BulkInventoryGet used by Catalog
    type BulkInventoryGet struct {

        
            Page Page  `json:"page"`
            Items []BulkInventoryGetItems  `json:"items"`
         
    }
    
    // InventoryJobPayload used by Catalog
    type InventoryJobPayload struct {

        
            ExpirationDate string  `json:"expiration_date"`
            TotalQuantity float64  `json:"total_quantity"`
            Price float64  `json:"price"`
            ItemWeightUnitOfMeasure string  `json:"item_weight_unit_of_measure"`
            TraceID string  `json:"trace_id"`
            StoreCode string  `json:"store_code"`
            Currency string  `json:"currency"`
            Tags []string  `json:"tags"`
            SellerIdentifier string  `json:"seller_identifier"`
            ItemDimensionsUnitOfMeasure string  `json:"item_dimensions_unit_of_measure"`
            PriceEffective float64  `json:"price_effective"`
            Quantity float64  `json:"quantity"`
            PriceMarked float64  `json:"price_marked"`
         
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

        
            Max float64  `json:"max"`
            Min float64  `json:"min"`
            Operators string  `json:"operators"`
         
    }
    
    // InventoryExportAdvanceOption used by Catalog
    type InventoryExportAdvanceOption struct {

        
            FromDate string  `json:"from_date"`
            BrandIds []float64  `json:"brand_ids"`
            Quantity InventoryExportQuantityFilter  `json:"quantity"`
            ToDate string  `json:"to_date"`
            StoreIds []float64  `json:"store_ids"`
         
    }
    
    // InventoryExportJob used by Catalog
    type InventoryExportJob struct {

        
            Status string  `json:"status"`
            TaskID string  `json:"task_id"`
            Filters InventoryExportAdvanceOption  `json:"filters"`
            Type string  `json:"type"`
            URL string  `json:"url"`
            CompletedOn string  `json:"completed_on"`
            NotificationEmails []string  `json:"notification_emails"`
            SellerID float64  `json:"seller_id"`
         
    }
    
    // InventoryExportRequest used by Catalog
    type InventoryExportRequest struct {

        
            Brand []float64  `json:"brand"`
            Type string  `json:"type"`
            Store []float64  `json:"store"`
         
    }
    
    // InventoryExportResponse used by Catalog
    type InventoryExportResponse struct {

        
            CreatedBy string  `json:"created_by"`
            ModifiedOn string  `json:"modified_on"`
            Status string  `json:"status"`
            TaskID string  `json:"task_id"`
            Filters map[string]interface{}  `json:"filters"`
            Type string  `json:"type"`
            NotificationEmails []string  `json:"notification_emails"`
            SellerID float64  `json:"seller_id"`
            CreatedOn string  `json:"created_on"`
         
    }
    
    // InventoryJobFilters used by Catalog
    type InventoryJobFilters struct {

        
            Stores []string  `json:"stores"`
            FromDate string  `json:"from_date"`
            Brands []string  `json:"brands"`
            Quantity InventoryExportQuantityFilter  `json:"quantity"`
            ToDate string  `json:"to_date"`
         
    }
    
    // InventoryJobDetailResponse used by Catalog
    type InventoryJobDetailResponse struct {

        
            CreatedBy UserDetail  `json:"created_by"`
            ModifiedOn string  `json:"modified_on"`
            CancelledOn string  `json:"cancelled_on"`
            ID string  `json:"id"`
            Status string  `json:"status"`
            TaskID string  `json:"task_id"`
            Filters InventoryJobFilters  `json:"filters"`
            Type string  `json:"type"`
            URL string  `json:"url"`
            CompletedOn string  `json:"completed_on"`
            CancelledBy UserDetail  `json:"cancelled_by"`
            NotificationEmails []string  `json:"notification_emails"`
            SellerID float64  `json:"seller_id"`
            CreatedOn string  `json:"created_on"`
         
    }
    
    // InventoryExportJobListResponse used by Catalog
    type InventoryExportJobListResponse struct {

        
            Items InventoryJobDetailResponse  `json:"items"`
         
    }
    
    // InventoryExportFilter used by Catalog
    type InventoryExportFilter struct {

        
            FromDate string  `json:"from_date"`
            BrandIds []float64  `json:"brand_ids"`
            Quantity InventoryExportQuantityFilter  `json:"quantity"`
            ToDate string  `json:"to_date"`
            StoreIds []float64  `json:"store_ids"`
         
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
    
    // InventoryPayload used by Catalog
    type InventoryPayload struct {

        
            ExpirationDate string  `json:"expiration_date"`
            TotalQuantity float64  `json:"total_quantity"`
            TraceID string  `json:"trace_id"`
            StoreID float64  `json:"store_id"`
            SellerIdentifier string  `json:"seller_identifier"`
            PriceEffective float64  `json:"price_effective"`
            Tags []string  `json:"tags"`
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

        
            HsnCode string  `json:"hsn_code"`
            Tax2 float64  `json:"tax2"`
            ID string  `json:"id"`
            Hs2Code string  `json:"hs2_code"`
            TaxOnEsp bool  `json:"tax_on_esp"`
            CompanyID float64  `json:"company_id"`
            Threshold1 float64  `json:"threshold1"`
            TaxOnMrp bool  `json:"tax_on_mrp"`
            Threshold2 float64  `json:"threshold2"`
            Tax1 float64  `json:"tax1"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // HsnCode used by Catalog
    type HsnCode struct {

        
            Data HsnCodesObject  `json:"data"`
         
    }
    
    // HsnUpsert used by Catalog
    type HsnUpsert struct {

        
            IsActive bool  `json:"is_active"`
            HsnCode string  `json:"hsn_code"`
            Tax2 float64  `json:"tax2"`
            Hs2Code string  `json:"hs2_code"`
            TaxOnEsp bool  `json:"tax_on_esp"`
            CompanyID float64  `json:"company_id"`
            Threshold1 float64  `json:"threshold1"`
            TaxOnMrp bool  `json:"tax_on_mrp"`
            UID float64  `json:"uid"`
            Threshold2 float64  `json:"threshold2"`
            Tax1 float64  `json:"tax1"`
         
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

        
            Size float64  `json:"size"`
            HasNext bool  `json:"has_next"`
            HasPrevious bool  `json:"has_previous"`
            ItemTotal float64  `json:"item_total"`
            Current string  `json:"current"`
         
    }
    
    // TaxSlab used by Catalog
    type TaxSlab struct {

        
            Cess float64  `json:"cess"`
            EffectiveDate string  `json:"effective_date"`
            Threshold float64  `json:"threshold"`
            Rate float64  `json:"rate"`
         
    }
    
    // HSNDataInsertV2 used by Catalog
    type HSNDataInsertV2 struct {

        
            CreatedBy map[string]interface{}  `json:"created_by"`
            HsnCode string  `json:"hsn_code"`
            ModifiedOn string  `json:"modified_on"`
            Taxes []TaxSlab  `json:"taxes"`
            CreatedOn string  `json:"created_on"`
            Type string  `json:"type"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            HsnCodeID string  `json:"hsn_code_id"`
            ReportingHsn string  `json:"reporting_hsn"`
            CountryCode string  `json:"country_code"`
            Description string  `json:"description"`
         
    }
    
    // HsnCodesListingResponseSchemaV2 used by Catalog
    type HsnCodesListingResponseSchemaV2 struct {

        
            Page PageResponse  `json:"page"`
            Items []HSNDataInsertV2  `json:"items"`
         
    }
    
    // BrandItem used by Catalog
    type BrandItem struct {

        
            Banners ImageUrls  `json:"banners"`
            Name string  `json:"name"`
            Logo Media  `json:"logo"`
            Discount string  `json:"discount"`
            UID float64  `json:"uid"`
            Action Action  `json:"action"`
            Departments []string  `json:"departments"`
            Slug string  `json:"slug"`
         
    }
    
    // BrandListingResponse used by Catalog
    type BrandListingResponse struct {

        
            Page Page  `json:"page"`
            Items []BrandItem  `json:"items"`
         
    }
    
    // Department used by Catalog
    type Department struct {

        
            PriorityOrder float64  `json:"priority_order"`
            Name string  `json:"name"`
            Logo Media  `json:"logo"`
            UID float64  `json:"uid"`
            Slug string  `json:"slug"`
         
    }
    
    // DepartmentResponse used by Catalog
    type DepartmentResponse struct {

        
            Items []Department  `json:"items"`
         
    }
    
    // ThirdLevelChild used by Catalog
    type ThirdLevelChild struct {

        
            Childs []map[string]interface{}  `json:"childs"`
            Banners ImageUrls  `json:"banners"`
            Name string  `json:"name"`
            UID float64  `json:"uid"`
            Action Action  `json:"action"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Slug string  `json:"slug"`
         
    }
    
    // SecondLevelChild used by Catalog
    type SecondLevelChild struct {

        
            Childs []ThirdLevelChild  `json:"childs"`
            Banners ImageUrls  `json:"banners"`
            Name string  `json:"name"`
            UID float64  `json:"uid"`
            Action Action  `json:"action"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Slug string  `json:"slug"`
         
    }
    
    // Child used by Catalog
    type Child struct {

        
            Childs []SecondLevelChild  `json:"childs"`
            Banners ImageUrls  `json:"banners"`
            Name string  `json:"name"`
            UID float64  `json:"uid"`
            Action Action  `json:"action"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Slug string  `json:"slug"`
         
    }
    
    // CategoryItems used by Catalog
    type CategoryItems struct {

        
            Childs []Child  `json:"childs"`
            Banners ImageUrls  `json:"banners"`
            Name string  `json:"name"`
            UID float64  `json:"uid"`
            Action Action  `json:"action"`
            Slug string  `json:"slug"`
         
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

        
            Page Page  `json:"page"`
            Filters []ProductFilters  `json:"filters"`
            Items []ProductListingDetail  `json:"items"`
            Operators map[string]interface{}  `json:"operators"`
            SortOn []ProductSortOn  `json:"sort_on"`
         
    }
    
    // ProductDetail used by Catalog
    type ProductDetail struct {

        
            ProductOnlineDate string  `json:"product_online_date"`
            HasVariant bool  `json:"has_variant"`
            Name string  `json:"name"`
            Attributes map[string]interface{}  `json:"attributes"`
            Rating float64  `json:"rating"`
            RatingCount float64  `json:"rating_count"`
            ImageNature string  `json:"image_nature"`
            ShortDescription string  `json:"short_description"`
            Type string  `json:"type"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            Medias []Media1  `json:"medias"`
            PromoMeta map[string]interface{}  `json:"promo_meta"`
            Slug string  `json:"slug"`
            ItemCode string  `json:"item_code"`
            Highlights []string  `json:"highlights"`
            Similars []string  `json:"similars"`
            Color string  `json:"color"`
            UID float64  `json:"uid"`
            Brand ProductBrand  `json:"brand"`
            Description string  `json:"description"`
            ItemType string  `json:"item_type"`
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
            Tryouts []string  `json:"tryouts"`
         
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

        
            Page InventoryPage  `json:"page"`
            Items []map[string]interface{}  `json:"items"`
         
    }
    
    // ArticleAssignment used by Catalog
    type ArticleAssignment struct {

        
            Level string  `json:"level"`
            Strategy string  `json:"strategy"`
         
    }
    
    // ArticleQuery used by Catalog
    type ArticleQuery struct {

        
            ItemID float64  `json:"item_id"`
            IgnoredStores []float64  `json:"ignored_stores"`
            Size string  `json:"size"`
         
    }
    
    // AssignStoreArticle used by Catalog
    type AssignStoreArticle struct {

        
            GroupID string  `json:"group_id"`
            ArticleAssignment ArticleAssignment  `json:"article_assignment"`
            Query ArticleQuery  `json:"query"`
            Quantity float64  `json:"quantity"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // AssignStore used by Catalog
    type AssignStore struct {

        
            ChannelIdentifier string  `json:"channel_identifier"`
            CompanyID float64  `json:"company_id"`
            AppID string  `json:"app_id"`
            Pincode string  `json:"pincode"`
            ChannelType string  `json:"channel_type"`
            Articles []AssignStoreArticle  `json:"articles"`
            StoreIds []float64  `json:"store_ids"`
         
    }
    
    // StoreAssignError used by Catalog
    type StoreAssignError struct {

        
            Value interface{}  `json:"value"`
            Type string  `json:"type"`
            Message string  `json:"message"`
         
    }
    
    // ArticleAssignment1 used by Catalog
    type ArticleAssignment1 struct {

        
            Level string  `json:"level"`
            Strategy string  `json:"strategy"`
         
    }
    
    // StoreAssign used by Catalog
    type StoreAssign struct {

        
            PriceEffective float64  `json:"price_effective"`
            Size string  `json:"size"`
            SCity string  `json:"s_city"`
            Status bool  `json:"status"`
            CompanyID float64  `json:"company_id"`
            StorePincode float64  `json:"store_pincode"`
            ArticleAssignment ArticleAssignment1  `json:"article_assignment"`
            ItemID float64  `json:"item_id"`
            Index float64  `json:"index"`
            GroupID string  `json:"group_id"`
            UID string  `json:"uid"`
            StrategyWiseListing []map[string]interface{}  `json:"strategy_wise_listing"`
            StoreID float64  `json:"store_id"`
            ID string  `json:"_id"`
            Quantity float64  `json:"quantity"`
            PriceMarked float64  `json:"price_marked"`
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
    
    // Document used by Catalog
    type Document struct {

        
            Value string  `json:"value"`
            Type string  `json:"type"`
            URL string  `json:"url"`
            LegalName string  `json:"legal_name"`
            Verified bool  `json:"verified"`
         
    }
    
    // LocationTimingSerializer used by Catalog
    type LocationTimingSerializer struct {

        
            Hour float64  `json:"hour"`
            Minute float64  `json:"minute"`
         
    }
    
    // LocationDayWiseSerializer used by Catalog
    type LocationDayWiseSerializer struct {

        
            Closing LocationTimingSerializer  `json:"closing"`
            Opening LocationTimingSerializer  `json:"opening"`
            Weekday string  `json:"weekday"`
            Open bool  `json:"open"`
         
    }
    
    // GetAddressSerializer used by Catalog
    type GetAddressSerializer struct {

        
            Landmark string  `json:"landmark"`
            Longitude float64  `json:"longitude"`
            Latitude float64  `json:"latitude"`
            Pincode float64  `json:"pincode"`
            Address2 string  `json:"address2"`
            State string  `json:"state"`
            City string  `json:"city"`
            Country string  `json:"country"`
            Address1 string  `json:"address1"`
            CountryCode string  `json:"country_code"`
            AddressType string  `json:"address_type"`
         
    }
    
    // UserSerializer1 used by Catalog
    type UserSerializer1 struct {

        
            Username string  `json:"username"`
            Contact string  `json:"contact"`
            UserID string  `json:"user_id"`
         
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
    
    // LocationManagerSerializer used by Catalog
    type LocationManagerSerializer struct {

        
            MobileNo SellerPhoneNumber  `json:"mobile_no"`
            Email string  `json:"email"`
            Name string  `json:"name"`
         
    }
    
    // UserSerializer2 used by Catalog
    type UserSerializer2 struct {

        
            Username string  `json:"username"`
            Contact string  `json:"contact"`
            UserID string  `json:"user_id"`
         
    }
    
    // GetCompanySerializer used by Catalog
    type GetCompanySerializer struct {

        
            RejectReason string  `json:"reject_reason"`
            CreatedBy UserSerializer2  `json:"created_by"`
            VerifiedOn string  `json:"verified_on"`
            Name string  `json:"name"`
            Stage string  `json:"stage"`
            VerifiedBy UserSerializer2  `json:"verified_by"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            BusinessType string  `json:"business_type"`
            UID float64  `json:"uid"`
            ModifiedBy UserSerializer2  `json:"modified_by"`
            CompanyType string  `json:"company_type"`
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
         
    }
    
    // ProductReturnConfigSerializer used by Catalog
    type ProductReturnConfigSerializer struct {

        
            StoreUID float64  `json:"store_uid"`
            OnSameStore bool  `json:"on_same_store"`
         
    }
    
    // GetLocationSerializer used by Catalog
    type GetLocationSerializer struct {

        
            GstCredentials InvoiceDetailsSerializer  `json:"gst_credentials"`
            DisplayName string  `json:"display_name"`
            Documents []Document  `json:"documents"`
            Name string  `json:"name"`
            Timing []LocationDayWiseSerializer  `json:"timing"`
            Address GetAddressSerializer  `json:"address"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            NotificationEmails []string  `json:"notification_emails"`
            Code string  `json:"code"`
            StoreType string  `json:"store_type"`
            Warnings map[string]interface{}  `json:"warnings"`
            CreatedOn string  `json:"created_on"`
            CreatedBy UserSerializer1  `json:"created_by"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            IntegrationType LocationIntegrationType  `json:"integration_type"`
            Stage string  `json:"stage"`
            UID float64  `json:"uid"`
            Manager LocationManagerSerializer  `json:"manager"`
            ModifiedOn string  `json:"modified_on"`
            Company GetCompanySerializer  `json:"company"`
            PhoneNumber string  `json:"phone_number"`
            VerifiedOn string  `json:"verified_on"`
            VerifiedBy UserSerializer1  `json:"verified_by"`
            ProductReturnConfig ProductReturnConfigSerializer  `json:"product_return_config"`
            ModifiedBy UserSerializer1  `json:"modified_by"`
         
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

        
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            Logo string  `json:"logo"`
            AppID string  `json:"app_id"`
            UID float64  `json:"uid"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
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
    
    // Document used by CompanyProfile
    type Document struct {

        
            LegalName string  `json:"legal_name"`
            Value string  `json:"value"`
            Type string  `json:"type"`
            URL string  `json:"url"`
            Verified bool  `json:"verified"`
         
    }
    
    // CompanyTaxesSerializer used by CompanyProfile
    type CompanyTaxesSerializer struct {

        
            Rate float64  `json:"rate"`
            EffectiveDate string  `json:"effective_date"`
            Enable bool  `json:"enable"`
         
    }
    
    // GetAddressSerializer used by CompanyProfile
    type GetAddressSerializer struct {

        
            Longitude float64  `json:"longitude"`
            Address2 string  `json:"address2"`
            Latitude float64  `json:"latitude"`
            City string  `json:"city"`
            State string  `json:"state"`
            Country string  `json:"country"`
            CountryCode string  `json:"country_code"`
            Address1 string  `json:"address1"`
            AddressType string  `json:"address_type"`
            Pincode float64  `json:"pincode"`
            Landmark string  `json:"landmark"`
         
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
    
    // GetCompanyProfileSerializerResponse used by CompanyProfile
    type GetCompanyProfileSerializerResponse struct {

        
            VerifiedBy UserSerializer  `json:"verified_by"`
            Documents []Document  `json:"documents"`
            NotificationEmails []string  `json:"notification_emails"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Stage string  `json:"stage"`
            CreatedOn string  `json:"created_on"`
            Taxes []CompanyTaxesSerializer  `json:"taxes"`
            CreatedBy UserSerializer  `json:"created_by"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            FranchiseEnabled bool  `json:"franchise_enabled"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            ContactDetails ContactDetails  `json:"contact_details"`
            Name string  `json:"name"`
            CompanyType string  `json:"company_type"`
            VerifiedOn string  `json:"verified_on"`
            ModifiedOn string  `json:"modified_on"`
            Mode string  `json:"mode"`
            BusinessInfo string  `json:"business_info"`
            BusinessCountryInfo BusinessCountryInfo  `json:"business_country_info"`
            UID float64  `json:"uid"`
            Warnings map[string]interface{}  `json:"warnings"`
            BusinessType string  `json:"business_type"`
            BusinessDetails BusinessDetails  `json:"business_details"`
         
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

        
            Rate float64  `json:"rate"`
            EffectiveDate string  `json:"effective_date"`
            Enable bool  `json:"enable"`
         
    }
    
    // CreateUpdateAddressSerializer used by CompanyProfile
    type CreateUpdateAddressSerializer struct {

        
            Longitude float64  `json:"longitude"`
            Address2 string  `json:"address2"`
            Latitude float64  `json:"latitude"`
            City string  `json:"city"`
            State string  `json:"state"`
            Country string  `json:"country"`
            CountryCode string  `json:"country_code"`
            Address1 string  `json:"address1"`
            AddressType string  `json:"address_type"`
            Pincode float64  `json:"pincode"`
            Landmark string  `json:"landmark"`
         
    }
    
    // UpdateCompany used by CompanyProfile
    type UpdateCompany struct {

        
            BusinessInfo string  `json:"business_info"`
            Taxes []CompanyTaxesSerializer1  `json:"taxes"`
            FranchiseEnabled bool  `json:"franchise_enabled"`
            Warnings map[string]interface{}  `json:"warnings"`
            RejectReason string  `json:"reject_reason"`
            ContactDetails ContactDetails  `json:"contact_details"`
            Documents []Document  `json:"documents"`
            Name string  `json:"name"`
            BusinessType string  `json:"business_type"`
            CompanyType string  `json:"company_type"`
            BusinessDetails BusinessDetails  `json:"business_details"`
            NotificationEmails []string  `json:"notification_emails"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Addresses []CreateUpdateAddressSerializer  `json:"addresses"`
         
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

        
            UID float64  `json:"uid"`
            CompanyDocuments DocumentsObj  `json:"company_documents"`
            Store DocumentsObj  `json:"store"`
            Brand DocumentsObj  `json:"brand"`
            StoreDocuments DocumentsObj  `json:"store_documents"`
            Product DocumentsObj  `json:"product"`
            Stage string  `json:"stage"`
         
    }
    
    // BrandBannerSerializer used by CompanyProfile
    type BrandBannerSerializer struct {

        
            Portrait string  `json:"portrait"`
            Landscape string  `json:"landscape"`
         
    }
    
    // GetBrandResponseSerializer used by CompanyProfile
    type GetBrandResponseSerializer struct {

        
            VerifiedBy UserSerializer  `json:"verified_by"`
            Banner BrandBannerSerializer  `json:"banner"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Stage string  `json:"stage"`
            CreatedOn string  `json:"created_on"`
            CreatedBy UserSerializer  `json:"created_by"`
            Logo string  `json:"logo"`
            RejectReason string  `json:"reject_reason"`
            SlugKey string  `json:"slug_key"`
            Synonyms []string  `json:"synonyms"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            Name string  `json:"name"`
            VerifiedOn string  `json:"verified_on"`
            ModifiedOn string  `json:"modified_on"`
            Mode string  `json:"mode"`
            UID float64  `json:"uid"`
            Warnings map[string]interface{}  `json:"warnings"`
            Description string  `json:"description"`
         
    }
    
    // CreateUpdateBrandRequestSerializer used by CompanyProfile
    type CreateUpdateBrandRequestSerializer struct {

        
            BrandTier string  `json:"brand_tier"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            CompanyID float64  `json:"company_id"`
            UID float64  `json:"uid"`
            Logo string  `json:"logo"`
            Banner BrandBannerSerializer  `json:"banner"`
            Name string  `json:"name"`
            Description string  `json:"description"`
            Synonyms []string  `json:"synonyms"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
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

        
            VerifiedBy UserSerializer  `json:"verified_by"`
            CreatedBy UserSerializer  `json:"created_by"`
            BusinessCountryInfo BusinessCountryInfo  `json:"business_country_info"`
            MarketChannels []string  `json:"market_channels"`
            UID float64  `json:"uid"`
            Stage string  `json:"stage"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            RejectReason string  `json:"reject_reason"`
            Details CompanyDetails  `json:"details"`
            Name string  `json:"name"`
            CompanyType string  `json:"company_type"`
            BusinessType string  `json:"business_type"`
            VerifiedOn string  `json:"verified_on"`
            NotificationEmails []string  `json:"notification_emails"`
            CreatedOn string  `json:"created_on"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ModifiedOn string  `json:"modified_on"`
            Addresses []GetAddressSerializer  `json:"addresses"`
         
    }
    
    // CompanyBrandSerializer used by CompanyProfile
    type CompanyBrandSerializer struct {

        
            VerifiedBy UserSerializer  `json:"verified_by"`
            Company CompanySerializer  `json:"company"`
            CreatedBy UserSerializer  `json:"created_by"`
            UID float64  `json:"uid"`
            Stage string  `json:"stage"`
            Warnings map[string]interface{}  `json:"warnings"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            RejectReason string  `json:"reject_reason"`
            Brand GetBrandResponseSerializer  `json:"brand"`
            VerifiedOn string  `json:"verified_on"`
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
         
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

        
            Name string  `json:"name"`
            MobileNo SellerPhoneNumber  `json:"mobile_no"`
            Email string  `json:"email"`
         
    }
    
    // GetCompanySerializer used by CompanyProfile
    type GetCompanySerializer struct {

        
            VerifiedBy UserSerializer  `json:"verified_by"`
            CreatedBy UserSerializer  `json:"created_by"`
            UID float64  `json:"uid"`
            Stage string  `json:"stage"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            RejectReason string  `json:"reject_reason"`
            Name string  `json:"name"`
            BusinessType string  `json:"business_type"`
            CompanyType string  `json:"company_type"`
            VerifiedOn string  `json:"verified_on"`
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
            Addresses []GetAddressSerializer  `json:"addresses"`
         
    }
    
    // LocationTimingSerializer used by CompanyProfile
    type LocationTimingSerializer struct {

        
            Minute float64  `json:"minute"`
            Hour float64  `json:"hour"`
         
    }
    
    // LocationDayWiseSerializer used by CompanyProfile
    type LocationDayWiseSerializer struct {

        
            Opening LocationTimingSerializer  `json:"opening"`
            Weekday string  `json:"weekday"`
            Closing LocationTimingSerializer  `json:"closing"`
            Open bool  `json:"open"`
         
    }
    
    // ProductReturnConfigSerializer used by CompanyProfile
    type ProductReturnConfigSerializer struct {

        
            OnSameStore bool  `json:"on_same_store"`
            StoreUID float64  `json:"store_uid"`
         
    }
    
    // GetLocationSerializer used by CompanyProfile
    type GetLocationSerializer struct {

        
            VerifiedBy UserSerializer  `json:"verified_by"`
            Holiday []HolidaySchemaSerializer  `json:"holiday"`
            Documents []Document  `json:"documents"`
            GstCredentials InvoiceDetailsSerializer  `json:"gst_credentials"`
            NotificationEmails []string  `json:"notification_emails"`
            PhoneNumber string  `json:"phone_number"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Stage string  `json:"stage"`
            CreatedOn string  `json:"created_on"`
            CreatedBy UserSerializer  `json:"created_by"`
            DisplayName string  `json:"display_name"`
            Manager LocationManagerSerializer  `json:"manager"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            Name string  `json:"name"`
            VerifiedOn string  `json:"verified_on"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            ModifiedOn string  `json:"modified_on"`
            StoreType string  `json:"store_type"`
            Company GetCompanySerializer  `json:"company"`
            UID float64  `json:"uid"`
            Code string  `json:"code"`
            Timing []LocationDayWiseSerializer  `json:"timing"`
            Warnings map[string]interface{}  `json:"warnings"`
            ProductReturnConfig ProductReturnConfigSerializer  `json:"product_return_config"`
            Address GetAddressSerializer  `json:"address"`
         
    }
    
    // LocationListSerializer used by CompanyProfile
    type LocationListSerializer struct {

        
            Page Page  `json:"page"`
            Items []GetLocationSerializer  `json:"items"`
         
    }
    
    // AddressSerializer used by CompanyProfile
    type AddressSerializer struct {

        
            Longitude float64  `json:"longitude"`
            Address2 string  `json:"address2"`
            Latitude float64  `json:"latitude"`
            City string  `json:"city"`
            State string  `json:"state"`
            Country string  `json:"country"`
            CountryCode string  `json:"country_code"`
            Address1 string  `json:"address1"`
            AddressType string  `json:"address_type"`
            Pincode float64  `json:"pincode"`
            Landmark string  `json:"landmark"`
         
    }
    
    // LocationSerializer used by CompanyProfile
    type LocationSerializer struct {

        
            StoreType string  `json:"store_type"`
            Company float64  `json:"company"`
            Holiday []HolidaySchemaSerializer  `json:"holiday"`
            UID float64  `json:"uid"`
            Manager LocationManagerSerializer  `json:"manager"`
            Stage string  `json:"stage"`
            Code string  `json:"code"`
            Timing []LocationDayWiseSerializer  `json:"timing"`
            Warnings map[string]interface{}  `json:"warnings"`
            Documents []Document  `json:"documents"`
            Name string  `json:"name"`
            GstCredentials InvoiceDetailsSerializer  `json:"gst_credentials"`
            ProductReturnConfig ProductReturnConfigSerializer  `json:"product_return_config"`
            DisplayName string  `json:"display_name"`
            Address AddressSerializer  `json:"address"`
            NotificationEmails []string  `json:"notification_emails"`
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
    

    
    // CouponDateMeta used by Cart
    type CouponDateMeta struct {

        
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // CouponAction used by Cart
    type CouponAction struct {

        
            TxnMode string  `json:"txn_mode"`
            ActionDate string  `json:"action_date"`
         
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
    
    // PriceRange used by Cart
    type PriceRange struct {

        
            Min float64  `json:"min"`
            Max float64  `json:"max"`
         
    }
    
    // PaymentAllowValue used by Cart
    type PaymentAllowValue struct {

        
            Max float64  `json:"max"`
         
    }
    
    // PaymentModes used by Cart
    type PaymentModes struct {

        
            Codes []string  `json:"codes"`
            Types []string  `json:"types"`
            Uses PaymentAllowValue  `json:"uses"`
            Networks []string  `json:"networks"`
         
    }
    
    // Restrictions used by Cart
    type Restrictions struct {

        
            UserGroups []float64  `json:"user_groups"`
            PostOrder PostOrder  `json:"post_order"`
            BulkBundle BulkBundleRestriction  `json:"bulk_bundle"`
            Uses UsesRestriction  `json:"uses"`
            CouponAllowed bool  `json:"coupon_allowed"`
            PriceRange PriceRange  `json:"price_range"`
            Platforms []string  `json:"platforms"`
            Payments map[string]PaymentModes  `json:"payments"`
            OrderingStores []float64  `json:"ordering_stores"`
            UserType string  `json:"user_type"`
         
    }
    
    // CouponAuthor used by Cart
    type CouponAuthor struct {

        
            ModifiedBy string  `json:"modified_by"`
            CreatedBy string  `json:"created_by"`
         
    }
    
    // Validation used by Cart
    type Validation struct {

        
            Anonymous bool  `json:"anonymous"`
            UserRegisteredAfter string  `json:"user_registered_after"`
            AppID []string  `json:"app_id"`
         
    }
    
    // State used by Cart
    type State struct {

        
            IsDisplay bool  `json:"is_display"`
            IsArchived bool  `json:"is_archived"`
            IsPublic bool  `json:"is_public"`
         
    }
    
    // Ownership used by Cart
    type Ownership struct {

        
            PayableBy string  `json:"payable_by"`
            PayableCategory string  `json:"payable_category"`
         
    }
    
    // Rule used by Cart
    type Rule struct {

        
            DiscountQty float64  `json:"discount_qty"`
            Key float64  `json:"key"`
            Value float64  `json:"value"`
            Min float64  `json:"min"`
            Max float64  `json:"max"`
         
    }
    
    // RuleDefinition used by Cart
    type RuleDefinition struct {

        
            IsExact bool  `json:"is_exact"`
            ValueType string  `json:"value_type"`
            ApplicableOn string  `json:"applicable_on"`
            CurrencyCode string  `json:"currency_code"`
            Scope []string  `json:"scope"`
            AutoApply bool  `json:"auto_apply"`
            CalculateOn string  `json:"calculate_on"`
            Type string  `json:"type"`
         
    }
    
    // Validity used by Cart
    type Validity struct {

        
            Priority float64  `json:"priority"`
         
    }
    
    // DisplayMetaDict used by Cart
    type DisplayMetaDict struct {

        
            Subtitle string  `json:"subtitle"`
            Title string  `json:"title"`
         
    }
    
    // DisplayMeta used by Cart
    type DisplayMeta struct {

        
            Title string  `json:"title"`
            Auto DisplayMetaDict  `json:"auto"`
            Remove DisplayMetaDict  `json:"remove"`
            Subtitle string  `json:"subtitle"`
            Apply DisplayMetaDict  `json:"apply"`
            Description string  `json:"description"`
         
    }
    
    // CouponSchedule used by Cart
    type CouponSchedule struct {

        
            Start string  `json:"start"`
            Cron string  `json:"cron"`
            NextSchedule []map[string]interface{}  `json:"next_schedule"`
            Duration float64  `json:"duration"`
            End string  `json:"end"`
         
    }
    
    // Identifier used by Cart
    type Identifier struct {

        
            BrandID []float64  `json:"brand_id"`
            CollectionID []string  `json:"collection_id"`
            ExcludeBrandID []float64  `json:"exclude_brand_id"`
            CompanyID []float64  `json:"company_id"`
            StoreID []float64  `json:"store_id"`
            CategoryID []float64  `json:"category_id"`
            UserID []string  `json:"user_id"`
            ItemID []float64  `json:"item_id"`
            ArticleID []string  `json:"article_id"`
         
    }
    
    // CouponAdd used by Cart
    type CouponAdd struct {

        
            DateMeta CouponDateMeta  `json:"date_meta"`
            Action CouponAction  `json:"action"`
            Restrictions Restrictions  `json:"restrictions"`
            Author CouponAuthor  `json:"author"`
            Validation Validation  `json:"validation"`
            State State  `json:"state"`
            Ownership Ownership  `json:"ownership"`
            Code string  `json:"code"`
            Tags []string  `json:"tags"`
            Rule []Rule  `json:"rule"`
            RuleDefinition RuleDefinition  `json:"rule_definition"`
            Validity Validity  `json:"validity"`
            DisplayMeta DisplayMeta  `json:"display_meta"`
            Schedule CouponSchedule  `json:"_schedule"`
            TypeSlug string  `json:"type_slug"`
            Identifiers Identifier  `json:"identifiers"`
         
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

        
            DateMeta CouponDateMeta  `json:"date_meta"`
            Action CouponAction  `json:"action"`
            Restrictions Restrictions  `json:"restrictions"`
            Author CouponAuthor  `json:"author"`
            Validation Validation  `json:"validation"`
            State State  `json:"state"`
            Ownership Ownership  `json:"ownership"`
            Code string  `json:"code"`
            Tags []string  `json:"tags"`
            Rule []Rule  `json:"rule"`
            RuleDefinition RuleDefinition  `json:"rule_definition"`
            Validity Validity  `json:"validity"`
            DisplayMeta DisplayMeta  `json:"display_meta"`
            Schedule CouponSchedule  `json:"_schedule"`
            TypeSlug string  `json:"type_slug"`
            Identifiers Identifier  `json:"identifiers"`
         
    }
    
    // CouponPartialUpdate used by Cart
    type CouponPartialUpdate struct {

        
            Archive bool  `json:"archive"`
            Schedule CouponSchedule  `json:"schedule"`
         
    }
    
    // PromotionDateMeta used by Cart
    type PromotionDateMeta struct {

        
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // Ownership1 used by Cart
    type Ownership1 struct {

        
            PayableBy string  `json:"payable_by"`
            PayableCategory string  `json:"payable_category"`
         
    }
    
    // CompareObject used by Cart
    type CompareObject struct {

        
            LessThan float64  `json:"less_than"`
            GreaterThanEquals float64  `json:"greater_than_equals"`
            Equals float64  `json:"equals"`
            GreaterThan float64  `json:"greater_than"`
            LessThanEquals float64  `json:"less_than_equals"`
         
    }
    
    // ItemCriteria used by Cart
    type ItemCriteria struct {

        
            ItemExcludeBrand []float64  `json:"item_exclude_brand"`
            AvailableZones []string  `json:"available_zones"`
            ItemSize []string  `json:"item_size"`
            CartTotal CompareObject  `json:"cart_total"`
            ItemSku []string  `json:"item_sku"`
            ItemCompany []float64  `json:"item_company"`
            ItemBrand []float64  `json:"item_brand"`
            ItemStore []float64  `json:"item_store"`
            CartQuantity CompareObject  `json:"cart_quantity"`
            ItemExcludeSku []string  `json:"item_exclude_sku"`
            AllItems bool  `json:"all_items"`
            CartUniqueItemAmount CompareObject  `json:"cart_unique_item_amount"`
            ItemID []float64  `json:"item_id"`
            ItemExcludeID []float64  `json:"item_exclude_id"`
            ItemExcludeCompany []float64  `json:"item_exclude_company"`
            ItemExcludeStore []float64  `json:"item_exclude_store"`
            BuyRules []string  `json:"buy_rules"`
            ProductTags []string  `json:"product_tags"`
            CartUniqueItemQuantity CompareObject  `json:"cart_unique_item_quantity"`
            ItemExcludeCategory []float64  `json:"item_exclude_category"`
            ItemCategory []float64  `json:"item_category"`
         
    }
    
    // DiscountOffer used by Cart
    type DiscountOffer struct {

        
            MinOfferQuantity float64  `json:"min_offer_quantity"`
            Code string  `json:"code"`
            DiscountPercentage float64  `json:"discount_percentage"`
            ApportionDiscount bool  `json:"apportion_discount"`
            MaxDiscountAmount float64  `json:"max_discount_amount"`
            MaxUsagePerTransaction float64  `json:"max_usage_per_transaction"`
            MaxOfferQuantity float64  `json:"max_offer_quantity"`
            DiscountAmount float64  `json:"discount_amount"`
            DiscountPrice float64  `json:"discount_price"`
            PartialCanRet bool  `json:"partial_can_ret"`
         
    }
    
    // DiscountRule used by Cart
    type DiscountRule struct {

        
            DiscountType string  `json:"discount_type"`
            BuyCondition string  `json:"buy_condition"`
            ItemCriteria ItemCriteria  `json:"item_criteria"`
            Offer DiscountOffer  `json:"offer"`
         
    }
    
    // PromotionAction used by Cart
    type PromotionAction struct {

        
            ActionType string  `json:"action_type"`
            ActionDate string  `json:"action_date"`
         
    }
    
    // PostOrder1 used by Cart
    type PostOrder1 struct {

        
            CancellationAllowed bool  `json:"cancellation_allowed"`
            ReturnAllowed bool  `json:"return_allowed"`
         
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

        
            OrderQuantity float64  `json:"order_quantity"`
            UserGroups []float64  `json:"user_groups"`
            PostOrder PostOrder1  `json:"post_order"`
            UserRegistered UserRegistered  `json:"user_registered"`
            AnonymousUsers bool  `json:"anonymous_users"`
            Uses UsesRestriction1  `json:"uses"`
            Platforms []string  `json:"platforms"`
            UserID []string  `json:"user_id"`
            Payments []PromotionPaymentModes  `json:"payments"`
         
    }
    
    // PromotionAuthor used by Cart
    type PromotionAuthor struct {

        
            ModifiedBy string  `json:"modified_by"`
            CreatedBy string  `json:"created_by"`
         
    }
    
    // Visibility used by Cart
    type Visibility struct {

        
            CouponList bool  `json:"coupon_list"`
            Pdp bool  `json:"pdp"`
         
    }
    
    // DisplayMeta1 used by Cart
    type DisplayMeta1 struct {

        
            Name string  `json:"name"`
            OfferText string  `json:"offer_text"`
            OfferLabel string  `json:"offer_label"`
            Description string  `json:"description"`
         
    }
    
    // PromotionSchedule used by Cart
    type PromotionSchedule struct {

        
            Start string  `json:"start"`
            Published bool  `json:"published"`
            Cron string  `json:"cron"`
            NextSchedule []map[string]interface{}  `json:"next_schedule"`
            Duration float64  `json:"duration"`
            End string  `json:"end"`
         
    }
    
    // PromotionListItem used by Cart
    type PromotionListItem struct {

        
            DateMeta PromotionDateMeta  `json:"date_meta"`
            Ownership Ownership1  `json:"ownership"`
            DiscountRules []DiscountRule  `json:"discount_rules"`
            CalculateOn string  `json:"calculate_on"`
            PostOrderAction PromotionAction  `json:"post_order_action"`
            Currency string  `json:"currency"`
            Restrictions Restrictions1  `json:"restrictions"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ApplyExclusive string  `json:"apply_exclusive"`
            PromoGroup string  `json:"promo_group"`
            Mode string  `json:"mode"`
            Author PromotionAuthor  `json:"author"`
            PromotionType string  `json:"promotion_type"`
            Visiblility Visibility  `json:"visiblility"`
            Code string  `json:"code"`
            ApplicationID string  `json:"application_id"`
            DisplayMeta DisplayMeta1  `json:"display_meta"`
            Schedule PromotionSchedule  `json:"_schedule"`
            ApplyPriority float64  `json:"apply_priority"`
            Stackable bool  `json:"stackable"`
            BuyRules map[string]ItemCriteria  `json:"buy_rules"`
            ApplyAllDiscount bool  `json:"apply_all_discount"`
         
    }
    
    // PromotionsResponse used by Cart
    type PromotionsResponse struct {

        
            Items PromotionListItem  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // PromotionAdd used by Cart
    type PromotionAdd struct {

        
            DateMeta PromotionDateMeta  `json:"date_meta"`
            Ownership Ownership1  `json:"ownership"`
            DiscountRules []DiscountRule  `json:"discount_rules"`
            CalculateOn string  `json:"calculate_on"`
            PostOrderAction PromotionAction  `json:"post_order_action"`
            Currency string  `json:"currency"`
            Restrictions Restrictions1  `json:"restrictions"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ApplyExclusive string  `json:"apply_exclusive"`
            PromoGroup string  `json:"promo_group"`
            Mode string  `json:"mode"`
            Author PromotionAuthor  `json:"author"`
            PromotionType string  `json:"promotion_type"`
            Visiblility Visibility  `json:"visiblility"`
            Code string  `json:"code"`
            ApplicationID string  `json:"application_id"`
            DisplayMeta DisplayMeta1  `json:"display_meta"`
            Schedule PromotionSchedule  `json:"_schedule"`
            ApplyPriority float64  `json:"apply_priority"`
            Stackable bool  `json:"stackable"`
            BuyRules map[string]ItemCriteria  `json:"buy_rules"`
            ApplyAllDiscount bool  `json:"apply_all_discount"`
         
    }
    
    // PromotionUpdate used by Cart
    type PromotionUpdate struct {

        
            DateMeta PromotionDateMeta  `json:"date_meta"`
            Ownership Ownership1  `json:"ownership"`
            DiscountRules []DiscountRule  `json:"discount_rules"`
            CalculateOn string  `json:"calculate_on"`
            PostOrderAction PromotionAction  `json:"post_order_action"`
            Currency string  `json:"currency"`
            Restrictions Restrictions1  `json:"restrictions"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ApplyExclusive string  `json:"apply_exclusive"`
            PromoGroup string  `json:"promo_group"`
            Mode string  `json:"mode"`
            Author PromotionAuthor  `json:"author"`
            PromotionType string  `json:"promotion_type"`
            Visiblility Visibility  `json:"visiblility"`
            Code string  `json:"code"`
            ApplicationID string  `json:"application_id"`
            DisplayMeta DisplayMeta1  `json:"display_meta"`
            Schedule PromotionSchedule  `json:"_schedule"`
            ApplyPriority float64  `json:"apply_priority"`
            Stackable bool  `json:"stackable"`
            BuyRules map[string]ItemCriteria  `json:"buy_rules"`
            ApplyAllDiscount bool  `json:"apply_all_discount"`
         
    }
    
    // PromotionPartialUpdate used by Cart
    type PromotionPartialUpdate struct {

        
            Archive bool  `json:"archive"`
            Schedule PromotionSchedule  `json:"schedule"`
         
    }
    
    // ActivePromosResponse used by Cart
    type ActivePromosResponse struct {

        
            Title string  `json:"title"`
            Example string  `json:"example"`
            EntitySlug string  `json:"entity_slug"`
            EntityType string  `json:"entity_type"`
            Subtitle string  `json:"subtitle"`
            CreatedOn string  `json:"created_on"`
            IsHidden bool  `json:"is_hidden"`
            Description string  `json:"description"`
            Type string  `json:"type"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // CartItem used by Cart
    type CartItem struct {

        
            Size string  `json:"size"`
            Quantity float64  `json:"quantity"`
            ProductID string  `json:"product_id"`
         
    }
    
    // OpenapiCartDetailsRequest used by Cart
    type OpenapiCartDetailsRequest struct {

        
            CartItems []CartItem  `json:"cart_items"`
         
    }
    
    // ProductPrice used by Cart
    type ProductPrice struct {

        
            CurrencyCode string  `json:"currency_code"`
            AddOn float64  `json:"add_on"`
            Selling float64  `json:"selling"`
            Marked float64  `json:"marked"`
            CurrencySymbol string  `json:"currency_symbol"`
            Effective float64  `json:"effective"`
         
    }
    
    // ProductPriceInfo used by Cart
    type ProductPriceInfo struct {

        
            Converted ProductPrice  `json:"converted"`
            Base ProductPrice  `json:"base"`
         
    }
    
    // FreeGiftItem used by Cart
    type FreeGiftItem struct {

        
            ItemName string  `json:"item_name"`
            ItemPriceDetails map[string]interface{}  `json:"item_price_details"`
            ItemBrandName string  `json:"item_brand_name"`
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
    
    // BuyRules used by Cart
    type BuyRules struct {

        
            ItemCriteria map[string]interface{}  `json:"item_criteria"`
            CartConditions map[string]interface{}  `json:"cart_conditions"`
         
    }
    
    // Ownership2 used by Cart
    type Ownership2 struct {

        
            PayableBy string  `json:"payable_by"`
            PayableCategory string  `json:"payable_category"`
         
    }
    
    // DiscountRulesApp used by Cart
    type DiscountRulesApp struct {

        
            MatchedBuyRules []string  `json:"matched_buy_rules"`
            RawOffer map[string]interface{}  `json:"raw_offer"`
            ItemCriteria map[string]interface{}  `json:"item_criteria"`
            Offer map[string]interface{}  `json:"offer"`
         
    }
    
    // AppliedPromotion used by Cart
    type AppliedPromotion struct {

        
            AppliedFreeArticles []AppliedFreeArticles  `json:"applied_free_articles"`
            BuyRules []BuyRules  `json:"buy_rules"`
            Amount float64  `json:"amount"`
            PromotionType string  `json:"promotion_type"`
            Ownership Ownership2  `json:"ownership"`
            MrpPromotion bool  `json:"mrp_promotion"`
            DiscountRules []DiscountRulesApp  `json:"discount_rules"`
            PromotionGroup string  `json:"promotion_group"`
            OfferText string  `json:"offer_text"`
            ArticleQuantity float64  `json:"article_quantity"`
            PromotionName string  `json:"promotion_name"`
            PromoID string  `json:"promo_id"`
         
    }
    
    // PromoMeta used by Cart
    type PromoMeta struct {

        
            Message string  `json:"message"`
         
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
    
    // NetQuantity used by Cart
    type NetQuantity struct {

        
            Value string  `json:"value"`
            Unit string  `json:"unit"`
         
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
    
    // ProductImage used by Cart
    type ProductImage struct {

        
            URL string  `json:"url"`
            SecureURL string  `json:"secure_url"`
            AspectRatio string  `json:"aspect_ratio"`
         
    }
    
    // CartProduct used by Cart
    type CartProduct struct {

        
            Action ProductAction  `json:"action"`
            ItemCode string  `json:"item_code"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            NetQuantity NetQuantity  `json:"net_quantity"`
            Categories []CategoryInfo  `json:"categories"`
            Brand BaseInfo  `json:"brand"`
            UID float64  `json:"uid"`
            Type string  `json:"type"`
            Images []ProductImage  `json:"images"`
         
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
    
    // BasePrice used by Cart
    type BasePrice struct {

        
            CurrencyCode string  `json:"currency_code"`
            Marked float64  `json:"marked"`
            CurrencySymbol string  `json:"currency_symbol"`
            Effective float64  `json:"effective"`
         
    }
    
    // ArticlePriceInfo used by Cart
    type ArticlePriceInfo struct {

        
            Converted BasePrice  `json:"converted"`
            Base BasePrice  `json:"base"`
         
    }
    
    // ProductArticle used by Cart
    type ProductArticle struct {

        
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            Quantity float64  `json:"quantity"`
            ProductGroupTags []string  `json:"product_group_tags"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            Seller BaseInfo  `json:"seller"`
            Size string  `json:"size"`
            Store BaseInfo  `json:"store"`
            Price ArticlePriceInfo  `json:"price"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            UID string  `json:"uid"`
            Type string  `json:"type"`
         
    }
    
    // CartProductIdentifer used by Cart
    type CartProductIdentifer struct {

        
            Identifier string  `json:"identifier"`
         
    }
    
    // ProductAvailabilitySize used by Cart
    type ProductAvailabilitySize struct {

        
            IsAvailable bool  `json:"is_available"`
            Value string  `json:"value"`
            Display string  `json:"display"`
         
    }
    
    // ProductAvailability used by Cart
    type ProductAvailability struct {

        
            IsValid bool  `json:"is_valid"`
            Deliverable bool  `json:"deliverable"`
            Sizes []string  `json:"sizes"`
            OtherStoreQuantity float64  `json:"other_store_quantity"`
            AvailableSizes []ProductAvailabilitySize  `json:"available_sizes"`
            OutOfStock bool  `json:"out_of_stock"`
         
    }
    
    // CartProductInfo used by Cart
    type CartProductInfo struct {

        
            PricePerUnit ProductPriceInfo  `json:"price_per_unit"`
            PromotionsApplied []AppliedPromotion  `json:"promotions_applied"`
            Message string  `json:"message"`
            Quantity float64  `json:"quantity"`
            PromoMeta PromoMeta  `json:"promo_meta"`
            Moq map[string]interface{}  `json:"moq"`
            Key string  `json:"key"`
            Discount string  `json:"discount"`
            Product CartProduct  `json:"product"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            Article ProductArticle  `json:"article"`
            Price ProductPriceInfo  `json:"price"`
            CouponMessage string  `json:"coupon_message"`
            BulkOffer map[string]interface{}  `json:"bulk_offer"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
            Availability ProductAvailability  `json:"availability"`
            IsSet bool  `json:"is_set"`
         
    }
    
    // CouponBreakup used by Cart
    type CouponBreakup struct {

        
            Title string  `json:"title"`
            Message string  `json:"message"`
            CouponValue float64  `json:"coupon_value"`
            Value float64  `json:"value"`
            MinimumCartValue float64  `json:"minimum_cart_value"`
            Code string  `json:"code"`
            SubTitle string  `json:"sub_title"`
            MaxDiscountValue float64  `json:"max_discount_value"`
            UID string  `json:"uid"`
            IsApplied bool  `json:"is_applied"`
            Description string  `json:"description"`
            CouponType string  `json:"coupon_type"`
            Type string  `json:"type"`
         
    }
    
    // RawBreakup used by Cart
    type RawBreakup struct {

        
            FyndCash float64  `json:"fynd_cash"`
            CodCharge float64  `json:"cod_charge"`
            Subtotal float64  `json:"subtotal"`
            YouSaved float64  `json:"you_saved"`
            Coupon float64  `json:"coupon"`
            DeliveryCharge float64  `json:"delivery_charge"`
            Discount float64  `json:"discount"`
            ConvenienceFee float64  `json:"convenience_fee"`
            GstCharges float64  `json:"gst_charges"`
            Vog float64  `json:"vog"`
            MrpTotal float64  `json:"mrp_total"`
            Total float64  `json:"total"`
         
    }
    
    // LoyaltyPoints used by Cart
    type LoyaltyPoints struct {

        
            Applicable float64  `json:"applicable"`
            Total float64  `json:"total"`
            IsApplied bool  `json:"is_applied"`
            Description string  `json:"description"`
         
    }
    
    // DisplayBreakup used by Cart
    type DisplayBreakup struct {

        
            Message []string  `json:"message"`
            Key string  `json:"key"`
            Value float64  `json:"value"`
            CurrencyCode string  `json:"currency_code"`
            CurrencySymbol string  `json:"currency_symbol"`
            Display string  `json:"display"`
         
    }
    
    // CartBreakup used by Cart
    type CartBreakup struct {

        
            Coupon CouponBreakup  `json:"coupon"`
            Raw RawBreakup  `json:"raw"`
            LoyaltyPoints LoyaltyPoints  `json:"loyalty_points"`
            Display []DisplayBreakup  `json:"display"`
         
    }
    
    // OpenapiCartDetailsResponse used by Cart
    type OpenapiCartDetailsResponse struct {

        
            Items []CartProductInfo  `json:"items"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            Message string  `json:"message"`
            IsValid bool  `json:"is_valid"`
         
    }
    
    // OpenApiErrorResponse used by Cart
    type OpenApiErrorResponse struct {

        
            Errors map[string]interface{}  `json:"errors"`
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // ShippingAddress used by Cart
    type ShippingAddress struct {

        
            Area string  `json:"area"`
            Landmark string  `json:"landmark"`
            State string  `json:"state"`
            CountryPhoneCode string  `json:"country_phone_code"`
            AreaCode string  `json:"area_code"`
            Phone float64  `json:"phone"`
            Name string  `json:"name"`
            Meta map[string]interface{}  `json:"meta"`
            CountryIsoCode string  `json:"country_iso_code"`
            City string  `json:"city"`
            Pincode float64  `json:"pincode"`
            Country string  `json:"country"`
            AreaCodeSlug string  `json:"area_code_slug"`
            CountryCode string  `json:"country_code"`
            Email string  `json:"email"`
            AddressType string  `json:"address_type"`
            Address string  `json:"address"`
         
    }
    
    // OpenApiCartServiceabilityRequest used by Cart
    type OpenApiCartServiceabilityRequest struct {

        
            CartItems []CartItem  `json:"cart_items"`
            ShippingAddress ShippingAddress  `json:"shipping_address"`
         
    }
    
    // OpenApiCartServiceabilityResponse used by Cart
    type OpenApiCartServiceabilityResponse struct {

        
            Message string  `json:"message"`
            IsValid bool  `json:"is_valid"`
            Items []CartProductInfo  `json:"items"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            BreakupValues CartBreakup  `json:"breakup_values"`
         
    }
    
    // MultiTenderPaymentMeta used by Cart
    type MultiTenderPaymentMeta struct {

        
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            CurrentStatus string  `json:"current_status"`
            OrderID string  `json:"order_id"`
            PaymentID string  `json:"payment_id"`
            PaymentGateway string  `json:"payment_gateway"`
         
    }
    
    // MultiTenderPaymentMethod used by Cart
    type MultiTenderPaymentMethod struct {

        
            Amount float64  `json:"amount"`
            Meta MultiTenderPaymentMeta  `json:"meta"`
            Mode string  `json:"mode"`
            Name string  `json:"name"`
         
    }
    
    // CartItemMeta used by Cart
    type CartItemMeta struct {

        
            PrimaryItem bool  `json:"primary_item"`
            GroupID string  `json:"group_id"`
         
    }
    
    // OpenApiFiles used by Cart
    type OpenApiFiles struct {

        
            Values []string  `json:"values"`
            Key string  `json:"key"`
         
    }
    
    // OpenApiOrderItem used by Cart
    type OpenApiOrderItem struct {

        
            DeliveryCharges float64  `json:"delivery_charges"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            CashbackApplied float64  `json:"cashback_applied"`
            Quantity float64  `json:"quantity"`
            PriceEffective float64  `json:"price_effective"`
            CodCharges float64  `json:"cod_charges"`
            Discount float64  `json:"discount"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            PaymentMethods []MultiTenderPaymentMethod  `json:"payment_methods"`
            Meta CartItemMeta  `json:"meta"`
            PriceMarked float64  `json:"price_marked"`
            Size string  `json:"size"`
            EmployeeDiscount float64  `json:"employee_discount"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            ProductID float64  `json:"product_id"`
            Files []OpenApiFiles  `json:"files"`
            AmountPaid float64  `json:"amount_paid"`
         
    }
    
    // OpenApiPlatformCheckoutReq used by Cart
    type OpenApiPlatformCheckoutReq struct {

        
            CashbackApplied float64  `json:"cashback_applied"`
            CartValue float64  `json:"cart_value"`
            CouponValue float64  `json:"coupon_value"`
            CodCharges float64  `json:"cod_charges"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            CurrencyCode string  `json:"currency_code"`
            Comment string  `json:"comment"`
            EmployeeDiscount map[string]interface{}  `json:"employee_discount"`
            OrderID string  `json:"order_id"`
            PaymentMode string  `json:"payment_mode"`
            DeliveryCharges float64  `json:"delivery_charges"`
            Coupon string  `json:"coupon"`
            CouponCode string  `json:"coupon_code"`
            Gstin string  `json:"gstin"`
            CartItems []OpenApiOrderItem  `json:"cart_items"`
            ShippingAddress ShippingAddress  `json:"shipping_address"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            PaymentMethods []MultiTenderPaymentMethod  `json:"payment_methods"`
            Files []OpenApiFiles  `json:"files"`
            BillingAddress ShippingAddress  `json:"billing_address"`
         
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

        
            CodCharges map[string]interface{}  `json:"cod_charges"`
            Discount float64  `json:"discount"`
            IsDefault bool  `json:"is_default"`
            Promotion map[string]interface{}  `json:"promotion"`
            PickUpCustomerDetails map[string]interface{}  `json:"pick_up_customer_details"`
            AppID string  `json:"app_id"`
            ID string  `json:"_id"`
            MergeQty bool  `json:"merge_qty"`
            CheckoutMode string  `json:"checkout_mode"`
            PaymentMode string  `json:"payment_mode"`
            IsActive bool  `json:"is_active"`
            DeliveryCharges map[string]interface{}  `json:"delivery_charges"`
            Articles []map[string]interface{}  `json:"articles"`
            Shipments []map[string]interface{}  `json:"shipments"`
            UID float64  `json:"uid"`
            UserID string  `json:"user_id"`
            PaymentMethods []map[string]interface{}  `json:"payment_methods"`
            ExpireAt string  `json:"expire_at"`
            Cashback map[string]interface{}  `json:"cashback"`
            BuyNow bool  `json:"buy_now"`
            CartValue float64  `json:"cart_value"`
            Meta map[string]interface{}  `json:"meta"`
            Comment string  `json:"comment"`
            OrderID string  `json:"order_id"`
            BulkCouponDiscount float64  `json:"bulk_coupon_discount"`
            FyndCredits map[string]interface{}  `json:"fynd_credits"`
            Coupon map[string]interface{}  `json:"coupon"`
            Gstin string  `json:"gstin"`
            IsArchive bool  `json:"is_archive"`
            CreatedOn string  `json:"created_on"`
            FcIndexMap []float64  `json:"fc_index_map"`
            Payments map[string]interface{}  `json:"payments"`
            LastModified string  `json:"last_modified"`
         
    }
    
    // AbandonedCartResponse used by Cart
    type AbandonedCartResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
            Items []AbandonedCart  `json:"items"`
            Result map[string]interface{}  `json:"result"`
            Page Page  `json:"page"`
         
    }
    
    // CartCurrency used by Cart
    type CartCurrency struct {

        
            Symbol string  `json:"symbol"`
            Code string  `json:"code"`
         
    }
    
    // PaymentSelectionLock used by Cart
    type PaymentSelectionLock struct {

        
            Enabled bool  `json:"enabled"`
            DefaultOptions string  `json:"default_options"`
            PaymentIdentifier string  `json:"payment_identifier"`
         
    }
    
    // CartDetailResponse used by Cart
    type CartDetailResponse struct {

        
            Message string  `json:"message"`
            CouponText string  `json:"coupon_text"`
            Items []CartProductInfo  `json:"items"`
            PanNo string  `json:"pan_no"`
            Currency CartCurrency  `json:"currency"`
            ID string  `json:"id"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            CheckoutMode string  `json:"checkout_mode"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            AppliedPromoDetails []AppliedPromotion  `json:"applied_promo_details"`
            BuyNow bool  `json:"buy_now"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            PanConfig map[string]interface{}  `json:"pan_config"`
            Comment string  `json:"comment"`
            IsValid bool  `json:"is_valid"`
            Gstin string  `json:"gstin"`
            LastModified string  `json:"last_modified"`
            BreakupValues CartBreakup  `json:"breakup_values"`
         
    }
    
    // AddProductCart used by Cart
    type AddProductCart struct {

        
            ArticleAssignment map[string]interface{}  `json:"article_assignment"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            Quantity float64  `json:"quantity"`
            ProductGroupTags []string  `json:"product_group_tags"`
            Pos bool  `json:"pos"`
            SellerID float64  `json:"seller_id"`
            ItemSize string  `json:"item_size"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            StoreID float64  `json:"store_id"`
            Display string  `json:"display"`
            ItemID float64  `json:"item_id"`
            ArticleID string  `json:"article_id"`
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
            Message string  `json:"message"`
            Success bool  `json:"success"`
            Cart CartDetailResponse  `json:"cart"`
         
    }
    
    // UpdateProductCart used by Cart
    type UpdateProductCart struct {

        
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            Quantity float64  `json:"quantity"`
            ItemIndex float64  `json:"item_index"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ItemSize string  `json:"item_size"`
            ItemID float64  `json:"item_id"`
            ArticleID string  `json:"article_id"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
         
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
            Meta map[string]interface{}  `json:"meta"`
            CreatedOn string  `json:"created_on"`
            Source map[string]interface{}  `json:"source"`
         
    }
    
    // SharedCart used by Cart
    type SharedCart struct {

        
            Message string  `json:"message"`
            CouponText string  `json:"coupon_text"`
            Items []CartProductInfo  `json:"items"`
            Currency CartCurrency  `json:"currency"`
            ID string  `json:"id"`
            CheckoutMode string  `json:"checkout_mode"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            SharedCartDetails SharedCartDetails  `json:"shared_cart_details"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            UID string  `json:"uid"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            BuyNow bool  `json:"buy_now"`
            CartID float64  `json:"cart_id"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            Comment string  `json:"comment"`
            IsValid bool  `json:"is_valid"`
            Gstin string  `json:"gstin"`
            LastModified string  `json:"last_modified"`
            BreakupValues CartBreakup  `json:"breakup_values"`
         
    }
    
    // SharedCartResponse used by Cart
    type SharedCartResponse struct {

        
            Error string  `json:"error"`
            Cart SharedCart  `json:"cart"`
         
    }
    
    // CartList used by Cart
    type CartList struct {

        
            CartID string  `json:"cart_id"`
            CartValue float64  `json:"cart_value"`
            ItemCounts float64  `json:"item_counts"`
            CreatedOn string  `json:"created_on"`
            PickUpCustomerDetails map[string]interface{}  `json:"pick_up_customer_details"`
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

        
            CreatedAt string  `json:"created_at"`
            ExternalID string  `json:"external_id"`
            Mobile string  `json:"mobile"`
            Gender string  `json:"gender"`
            LastName string  `json:"last_name"`
            FirstName string  `json:"first_name"`
            UID string  `json:"uid"`
            ID string  `json:"_id"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // UserCartMappingResponse used by Cart
    type UserCartMappingResponse struct {

        
            Message string  `json:"message"`
            CouponText string  `json:"coupon_text"`
            Items []CartProductInfo  `json:"items"`
            PanNo string  `json:"pan_no"`
            Currency CartCurrency  `json:"currency"`
            ID string  `json:"id"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            CheckoutMode string  `json:"checkout_mode"`
            User UserInfo  `json:"user"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            AppliedPromoDetails []AppliedPromotion  `json:"applied_promo_details"`
            BuyNow bool  `json:"buy_now"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            PanConfig map[string]interface{}  `json:"pan_config"`
            Comment string  `json:"comment"`
            IsValid bool  `json:"is_valid"`
            Gstin string  `json:"gstin"`
            LastModified string  `json:"last_modified"`
            BreakupValues CartBreakup  `json:"breakup_values"`
         
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

        
            Title string  `json:"title"`
            Message string  `json:"message"`
            CouponValue float64  `json:"coupon_value"`
            MinimumCartValue float64  `json:"minimum_cart_value"`
            CouponCode string  `json:"coupon_code"`
            CouponType string  `json:"coupon_type"`
            SubTitle string  `json:"sub_title"`
            MaxDiscountValue float64  `json:"max_discount_value"`
            IsApplicable bool  `json:"is_applicable"`
            Description string  `json:"description"`
            IsApplied bool  `json:"is_applied"`
            ExpiresOn string  `json:"expires_on"`
         
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

        
            Landmark string  `json:"landmark"`
            IsDefaultAddress bool  `json:"is_default_address"`
            GeoLocation GeoLocation  `json:"geo_location"`
            Name string  `json:"name"`
            Email string  `json:"email"`
            ID string  `json:"id"`
            CheckoutMode string  `json:"checkout_mode"`
            CountryCode string  `json:"country_code"`
            City string  `json:"city"`
            IsActive bool  `json:"is_active"`
            GoogleMapPoint map[string]interface{}  `json:"google_map_point"`
            Tags []string  `json:"tags"`
            AreaCode string  `json:"area_code"`
            UserID string  `json:"user_id"`
            CartID string  `json:"cart_id"`
            Meta map[string]interface{}  `json:"meta"`
            Address string  `json:"address"`
            Area string  `json:"area"`
            Country string  `json:"country"`
            State string  `json:"state"`
            Phone string  `json:"phone"`
            CreatedByUserID string  `json:"created_by_user_id"`
            AreaCodeSlug string  `json:"area_code_slug"`
            AddressType string  `json:"address_type"`
         
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

        
            IsDeleted bool  `json:"is_deleted"`
            ID string  `json:"id"`
         
    }
    
    // PlatformSelectCartAddressRequest used by Cart
    type PlatformSelectCartAddressRequest struct {

        
            CartID string  `json:"cart_id"`
            BillingAddressID string  `json:"billing_address_id"`
            UserID string  `json:"user_id"`
            ID string  `json:"id"`
            CheckoutMode string  `json:"checkout_mode"`
         
    }
    
    // ShipmentResponse used by Cart
    type ShipmentResponse struct {

        
            OrderType string  `json:"order_type"`
            FulfillmentID float64  `json:"fulfillment_id"`
            ShipmentType string  `json:"shipment_type"`
            Items []CartProductInfo  `json:"items"`
            Shipments float64  `json:"shipments"`
            FulfillmentType string  `json:"fulfillment_type"`
            DpOptions map[string]interface{}  `json:"dp_options"`
            Promise ShipmentPromise  `json:"promise"`
            BoxType string  `json:"box_type"`
            DpID string  `json:"dp_id"`
         
    }
    
    // CartShipmentsResponse used by Cart
    type CartShipmentsResponse struct {

        
            Error bool  `json:"error"`
            Message string  `json:"message"`
            CouponText string  `json:"coupon_text"`
            CartID float64  `json:"cart_id"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            Comment string  `json:"comment"`
            Currency CartCurrency  `json:"currency"`
            ID string  `json:"id"`
            CheckoutMode string  `json:"checkout_mode"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            IsValid bool  `json:"is_valid"`
            Gstin string  `json:"gstin"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            Shipments []ShipmentResponse  `json:"shipments"`
            UID string  `json:"uid"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            LastModified string  `json:"last_modified"`
            BuyNow bool  `json:"buy_now"`
            BreakupValues CartBreakup  `json:"breakup_values"`
         
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

        
            User string  `json:"user"`
            EmployeeCode string  `json:"employee_code"`
            LastName string  `json:"last_name"`
            FirstName string  `json:"first_name"`
            ID string  `json:"_id"`
         
    }
    
    // Files used by Cart
    type Files struct {

        
            Values []string  `json:"values"`
            Key string  `json:"key"`
         
    }
    
    // PlatformCartCheckoutDetailRequest used by Cart
    type PlatformCartCheckoutDetailRequest struct {

        
            OrderType string  `json:"order_type"`
            Aggregator string  `json:"aggregator"`
            ID string  `json:"id"`
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
            CheckoutMode string  `json:"checkout_mode"`
            PaymentMode string  `json:"payment_mode"`
            PaymentParams map[string]interface{}  `json:"payment_params"`
            PaymentIdentifier string  `json:"payment_identifier"`
            EmployeeCode string  `json:"employee_code"`
            Pos bool  `json:"pos"`
            UserID string  `json:"user_id"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            Meta map[string]interface{}  `json:"meta"`
            Staff StaffCheckout  `json:"staff"`
            OrderingStore float64  `json:"ordering_store"`
            MerchantCode string  `json:"merchant_code"`
            BillingAddressID string  `json:"billing_address_id"`
            DeviceID string  `json:"device_id"`
            PaymentAutoConfirm bool  `json:"payment_auto_confirm"`
            CallbackURL string  `json:"callback_url"`
            Files []Files  `json:"files"`
            AddressID string  `json:"address_id"`
            PickAtStoreUID float64  `json:"pick_at_store_uid"`
         
    }
    
    // CheckCart used by Cart
    type CheckCart struct {

        
            Message string  `json:"message"`
            CouponText string  `json:"coupon_text"`
            Success bool  `json:"success"`
            CodCharges float64  `json:"cod_charges"`
            Items []CartProductInfo  `json:"items"`
            StoreEmps []map[string]interface{}  `json:"store_emps"`
            Currency CartCurrency  `json:"currency"`
            ID string  `json:"id"`
            CheckoutMode string  `json:"checkout_mode"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            DeliveryCharges float64  `json:"delivery_charges"`
            StoreCode string  `json:"store_code"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            UID string  `json:"uid"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            DeliveryChargeOrderValue float64  `json:"delivery_charge_order_value"`
            BuyNow bool  `json:"buy_now"`
            CartID float64  `json:"cart_id"`
            CodAvailable bool  `json:"cod_available"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            Comment string  `json:"comment"`
            OrderID string  `json:"order_id"`
            UserType string  `json:"user_type"`
            CodMessage string  `json:"cod_message"`
            IsValid bool  `json:"is_valid"`
            Gstin string  `json:"gstin"`
            ErrorMessage string  `json:"error_message"`
            LastModified string  `json:"last_modified"`
            BreakupValues CartBreakup  `json:"breakup_values"`
         
    }
    
    // CartCheckoutResponse used by Cart
    type CartCheckoutResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
            Data map[string]interface{}  `json:"data"`
            Cart CheckCart  `json:"cart"`
            CallbackURL string  `json:"callback_url"`
            OrderID string  `json:"order_id"`
            PaymentConfirmURL string  `json:"payment_confirm_url"`
            AppInterceptURL string  `json:"app_intercept_url"`
         
    }
    
    // CartDeliveryModesResponse used by Cart
    type CartDeliveryModesResponse struct {

        
            AvailableModes []string  `json:"available_modes"`
            PickupStores []float64  `json:"pickup_stores"`
         
    }
    
    // PickupStoreDetail used by Cart
    type PickupStoreDetail struct {

        
            Area string  `json:"area"`
            Landmark string  `json:"landmark"`
            StoreCode string  `json:"store_code"`
            State string  `json:"state"`
            AreaCode string  `json:"area_code"`
            Name string  `json:"name"`
            Phone string  `json:"phone"`
            City string  `json:"city"`
            Pincode float64  `json:"pincode"`
            Country string  `json:"country"`
            AreaCodeSlug string  `json:"area_code_slug"`
            UID float64  `json:"uid"`
            ID float64  `json:"id"`
            Email string  `json:"email"`
            AddressType string  `json:"address_type"`
            Address string  `json:"address"`
         
    }
    
    // StoreDetailsResponse used by Cart
    type StoreDetailsResponse struct {

        
            Items []PickupStoreDetail  `json:"items"`
         
    }
    
    // UpdateCartPaymentRequest used by Cart
    type UpdateCartPaymentRequest struct {

        
            MerchantCode string  `json:"merchant_code"`
            PaymentMode string  `json:"payment_mode"`
            AggregatorName string  `json:"aggregator_name"`
            ID string  `json:"id"`
            AddressID string  `json:"address_id"`
            PaymentIdentifier string  `json:"payment_identifier"`
         
    }
    
    // CouponValidity used by Cart
    type CouponValidity struct {

        
            Title string  `json:"title"`
            Discount float64  `json:"discount"`
            Valid bool  `json:"valid"`
            Code string  `json:"code"`
            DisplayMessageEn string  `json:"display_message_en"`
         
    }
    
    // PaymentCouponValidate used by Cart
    type PaymentCouponValidate struct {

        
            CouponValidity CouponValidity  `json:"coupon_validity"`
            Message string  `json:"message"`
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
    
