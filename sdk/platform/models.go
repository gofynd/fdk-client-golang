package platform



    
    // ApplicationResponse used by Common
    type ApplicationResponse struct {

        
            Application Application  `json:"application"`
         
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
    
    // Application used by Common
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
    
    // LocationCountry used by Common
    type LocationCountry struct {

        
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
         
    }
    
    // Locations used by Common
    type Locations struct {

        
            Items []map[string]interface{}  `json:"items"`
         
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

        
            AppID string  `json:"app_id"`
            DisplayFields []string  `json:"display_fields"`
            Aggregators []map[string]interface{}  `json:"aggregators"`
            Created bool  `json:"created"`
            Success bool  `json:"success"`
            ExcludedFields []string  `json:"excluded_fields"`
         
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
            IsActive bool  `json:"is_active"`
            ConfigType string  `json:"config_type"`
            Secret string  `json:"secret"`
            MerchantSalt string  `json:"merchant_salt"`
         
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
            Logos PaymentModeLogo  `json:"logos"`
            Code string  `json:"code"`
            PackageName string  `json:"package_name"`
         
    }
    
    // PaymentModeList used by Payment
    type PaymentModeList struct {

        
            AggregatorName string  `json:"aggregator_name"`
            DisplayPriority float64  `json:"display_priority"`
            MerchantCode string  `json:"merchant_code"`
            FyndVpa string  `json:"fynd_vpa"`
            CardToken string  `json:"card_token"`
            CardReference string  `json:"card_reference"`
            Timeout float64  `json:"timeout"`
            CardIssuer string  `json:"card_issuer"`
            CardBrandImage string  `json:"card_brand_image"`
            DisplayName string  `json:"display_name"`
            CardName string  `json:"card_name"`
            CompliantWithTokenisationGuidelines bool  `json:"compliant_with_tokenisation_guidelines"`
            Code string  `json:"code"`
            CardNumber string  `json:"card_number"`
            ExpYear float64  `json:"exp_year"`
            ExpMonth float64  `json:"exp_month"`
            IntentAppErrorDictList []IntentAppErrorList  `json:"intent_app_error_dict_list"`
            Name string  `json:"name"`
            CodLimit float64  `json:"cod_limit"`
            LogoURL PaymentModeLogo  `json:"logo_url"`
            CardType string  `json:"card_type"`
            RemainingLimit float64  `json:"remaining_limit"`
            Expired bool  `json:"expired"`
            CardID string  `json:"card_id"`
            RetryCount float64  `json:"retry_count"`
            IntentApp []IntentApp  `json:"intent_app"`
            IntentAppErrorList []string  `json:"intent_app_error_list"`
            CodLimitPerOrder float64  `json:"cod_limit_per_order"`
            CardBrand string  `json:"card_brand"`
            CardIsin string  `json:"card_isin"`
            CardFingerprint string  `json:"card_fingerprint"`
            IntentFlow bool  `json:"intent_flow"`
            Nickname string  `json:"nickname"`
         
    }
    
    // RootPaymentMode used by Payment
    type RootPaymentMode struct {

        
            AggregatorName string  `json:"aggregator_name"`
            DisplayName string  `json:"display_name"`
            AnonymousEnable bool  `json:"anonymous_enable"`
            Name string  `json:"name"`
            DisplayPriority float64  `json:"display_priority"`
            List []PaymentModeList  `json:"list"`
            AddCardEnabled bool  `json:"add_card_enabled"`
            SaveCard bool  `json:"save_card"`
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
    
    // PayoutsResponse used by Payment
    type PayoutsResponse struct {

        
            MoreAttributes map[string]interface{}  `json:"more_attributes"`
            IsActive bool  `json:"is_active"`
            PayoutsAggregators []map[string]interface{}  `json:"payouts_aggregators"`
            TransferType string  `json:"transfer_type"`
            IsDefault bool  `json:"is_default"`
            Customers map[string]interface{}  `json:"customers"`
            UniqueTransferNo map[string]interface{}  `json:"unique_transfer_no"`
         
    }
    
    // PayoutBankDetails used by Payment
    type PayoutBankDetails struct {

        
            State string  `json:"state"`
            Pincode float64  `json:"pincode"`
            IfscCode string  `json:"ifsc_code"`
            BranchName string  `json:"branch_name"`
            Country string  `json:"country"`
            AccountNo string  `json:"account_no"`
            AccountType string  `json:"account_type"`
            BankName string  `json:"bank_name"`
            City string  `json:"city"`
            AccountHolder string  `json:"account_holder"`
         
    }
    
    // PayoutRequest used by Payment
    type PayoutRequest struct {

        
            IsActive bool  `json:"is_active"`
            BankDetails PayoutBankDetails  `json:"bank_details"`
            Users map[string]interface{}  `json:"users"`
            TransferType string  `json:"transfer_type"`
            UniqueExternalID string  `json:"unique_external_id"`
            Aggregator string  `json:"aggregator"`
         
    }
    
    // PayoutResponse used by Payment
    type PayoutResponse struct {

        
            BankDetails map[string]interface{}  `json:"bank_details"`
            IsActive bool  `json:"is_active"`
            Users map[string]interface{}  `json:"users"`
            Payouts map[string]interface{}  `json:"payouts"`
            TransferType string  `json:"transfer_type"`
            PaymentStatus string  `json:"payment_status"`
            Created bool  `json:"created"`
            Aggregator string  `json:"aggregator"`
            Success bool  `json:"success"`
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

        
            IsDefault bool  `json:"is_default"`
            UniqueExternalID string  `json:"unique_external_id"`
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
            Message string  `json:"message"`
            IsVerifiedFlag bool  `json:"is_verified_flag"`
            Success bool  `json:"success"`
         
    }
    
    // NotFoundResourceError used by Payment
    type NotFoundResourceError struct {

        
            Description string  `json:"description"`
            Success bool  `json:"success"`
            Code string  `json:"code"`
         
    }
    
    // BankDetailsForOTP used by Payment
    type BankDetailsForOTP struct {

        
            IfscCode string  `json:"ifsc_code"`
            BranchName string  `json:"branch_name"`
            AccountNo string  `json:"account_no"`
            BankName string  `json:"bank_name"`
            AccountHolder string  `json:"account_holder"`
         
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

        
            IsActive bool  `json:"is_active"`
            DelightsUserName string  `json:"delights_user_name"`
            AccountHolder string  `json:"account_holder"`
            Email string  `json:"email"`
            ModifiedOn string  `json:"modified_on"`
            Subtitle string  `json:"subtitle"`
            DisplayName string  `json:"display_name"`
            IfscCode string  `json:"ifsc_code"`
            AccountNo string  `json:"account_no"`
            Mobile string  `json:"mobile"`
            Title string  `json:"title"`
            CreatedOn string  `json:"created_on"`
            Address string  `json:"address"`
            BankName string  `json:"bank_name"`
            Comment string  `json:"comment"`
            ID float64  `json:"id"`
            BeneficiaryID string  `json:"beneficiary_id"`
            BranchName string  `json:"branch_name"`
            TransferMode string  `json:"transfer_mode"`
         
    }
    
    // OrderBeneficiaryResponse used by Payment
    type OrderBeneficiaryResponse struct {

        
            Beneficiaries []OrderBeneficiaryDetails  `json:"beneficiaries"`
            ShowBeneficiaryDetails bool  `json:"show_beneficiary_details"`
         
    }
    
    // MultiTenderPaymentMeta used by Payment
    type MultiTenderPaymentMeta struct {

        
            PaymentID string  `json:"payment_id"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            OrderID string  `json:"order_id"`
            CurrentStatus string  `json:"current_status"`
            PaymentGateway string  `json:"payment_gateway"`
         
    }
    
    // MultiTenderPaymentMethod used by Payment
    type MultiTenderPaymentMethod struct {

        
            Name string  `json:"name"`
            Amount float64  `json:"amount"`
            Meta MultiTenderPaymentMeta  `json:"meta"`
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
            Success bool  `json:"success"`
            OrderID string  `json:"order_id"`
         
    }
    
    // CODdata used by Payment
    type CODdata struct {

        
            IsActive bool  `json:"is_active"`
            Usages float64  `json:"usages"`
            UserID string  `json:"user_id"`
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
            Mobileno string  `json:"mobileno"`
            MerchantUserID string  `json:"merchant_user_id"`
         
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

        
            DeviceTag string  `json:"device_tag"`
            StoreID float64  `json:"store_id"`
            TerminalSerialNo string  `json:"terminal_serial_no"`
            EdcDeviceSerialNo string  `json:"edc_device_serial_no"`
            AggregatorID float64  `json:"aggregator_id"`
            EdcModel string  `json:"edc_model"`
         
    }
    
    // EdcDevice used by Payment
    type EdcDevice struct {

        
            AggregatorName string  `json:"aggregator_name"`
            DeviceTag string  `json:"device_tag"`
            TerminalUniqueIdentifier string  `json:"terminal_unique_identifier"`
            IsActive bool  `json:"is_active"`
            StoreID float64  `json:"store_id"`
            MerchantStorePosCode string  `json:"merchant_store_pos_code"`
            TerminalSerialNo string  `json:"terminal_serial_no"`
            EdcDeviceSerialNo string  `json:"edc_device_serial_no"`
            AggregatorID float64  `json:"aggregator_id"`
            ApplicationID string  `json:"application_id"`
            EdcModel string  `json:"edc_model"`
         
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

        
            DeviceTag string  `json:"device_tag"`
            IsActive bool  `json:"is_active"`
            StoreID float64  `json:"store_id"`
            MerchantStorePosCode string  `json:"merchant_store_pos_code"`
            EdcDeviceSerialNo string  `json:"edc_device_serial_no"`
            AggregatorID float64  `json:"aggregator_id"`
            EdcModel string  `json:"edc_model"`
         
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

        
            Currency string  `json:"currency"`
            CustomerID string  `json:"customer_id"`
            Amount float64  `json:"amount"`
            MerchantOrderID string  `json:"merchant_order_id"`
            OrderID string  `json:"order_id"`
            Timeout float64  `json:"timeout"`
            Method string  `json:"method"`
            Vpa string  `json:"vpa"`
            RazorpayPaymentID string  `json:"razorpay_payment_id"`
            Contact string  `json:"contact"`
            Email string  `json:"email"`
            DeviceID string  `json:"device_id"`
            Aggregator string  `json:"aggregator"`
         
    }
    
    // PaymentInitializationResponse used by Payment
    type PaymentInitializationResponse struct {

        
            Status string  `json:"status"`
            Currency string  `json:"currency"`
            VirtualID string  `json:"virtual_id"`
            CustomerID string  `json:"customer_id"`
            PollingURL string  `json:"polling_url"`
            Amount float64  `json:"amount"`
            MerchantOrderID string  `json:"merchant_order_id"`
            AggregatorOrderID string  `json:"aggregator_order_id"`
            UpiPollURL string  `json:"upi_poll_url"`
            DeviceID string  `json:"device_id"`
            Method string  `json:"method"`
            BqrImage string  `json:"bqr_image"`
            Vpa string  `json:"vpa"`
            RazorpayPaymentID string  `json:"razorpay_payment_id"`
            Aggregator string  `json:"aggregator"`
            Success bool  `json:"success"`
            Timeout float64  `json:"timeout"`
         
    }
    
    // PaymentStatusUpdateRequest used by Payment
    type PaymentStatusUpdateRequest struct {

        
            Status string  `json:"status"`
            Currency string  `json:"currency"`
            CustomerID string  `json:"customer_id"`
            Amount float64  `json:"amount"`
            MerchantOrderID string  `json:"merchant_order_id"`
            OrderID string  `json:"order_id"`
            Contact string  `json:"contact"`
            Method string  `json:"method"`
            Vpa string  `json:"vpa"`
            Aggregator string  `json:"aggregator"`
            Email string  `json:"email"`
            DeviceID string  `json:"device_id"`
            MerchantTransactionID string  `json:"merchant_transaction_id"`
         
    }
    
    // PaymentStatusUpdateResponse used by Payment
    type PaymentStatusUpdateResponse struct {

        
            AggregatorName string  `json:"aggregator_name"`
            Status string  `json:"status"`
            RedirectURL string  `json:"redirect_url"`
            Retry bool  `json:"retry"`
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

        
            AmountInPaisa string  `json:"amount_in_paisa"`
            PaymentModeIdentifier string  `json:"payment_mode_identifier"`
            PaymentID string  `json:"payment_id"`
            Currency string  `json:"currency"`
            AllStatus []string  `json:"all_status"`
            ApplicationID string  `json:"application_id"`
            RefundedBy string  `json:"refunded_by"`
            CreatedOn string  `json:"created_on"`
            UserObject map[string]interface{}  `json:"user_object"`
            CurrentStatus string  `json:"current_status"`
            AggregatorPaymentObject map[string]interface{}  `json:"aggregator_payment_object"`
            CollectedBy string  `json:"collected_by"`
            CompanyID string  `json:"company_id"`
            ModifiedOn string  `json:"modified_on"`
            PaymentGateway string  `json:"payment_gateway"`
            PaymentMode string  `json:"payment_mode"`
            ID string  `json:"id"`
            RefundObject map[string]interface{}  `json:"refund_object"`
         
    }
    
    // PaymentStatusObject used by Payment
    type PaymentStatusObject struct {

        
            PaymentObjectList []PaymentObjectListSerializer  `json:"payment_object_list"`
            MerchantOrderID string  `json:"merchant_order_id"`
         
    }
    
    // PaymentStatusBulkHandlerResponse used by Payment
    type PaymentStatusBulkHandlerResponse struct {

        
            Status float64  `json:"status"`
            Error string  `json:"error"`
            Data []PaymentStatusObject  `json:"data"`
            Success string  `json:"success"`
            Count float64  `json:"count"`
         
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

        
            PaymentModeIdentifier string  `json:"payment_mode_identifier"`
            Amount float64  `json:"amount"`
            MerchantOrderID string  `json:"merchant_order_id"`
            AggregatorOrderID string  `json:"aggregator_order_id"`
            FwdShipmentID string  `json:"fwd_shipment_id"`
            CurrentStatus string  `json:"current_status"`
            Aggregator string  `json:"aggregator"`
            OutstandingDetailsID float64  `json:"outstanding_details_id"`
            AggregatorTransactionID string  `json:"aggregator_transaction_id"`
            PaymentMode string  `json:"payment_mode"`
         
    }
    
    // RepaymentDetailsSerialiserPayAll used by Payment
    type RepaymentDetailsSerialiserPayAll struct {

        
            TotalAmount float64  `json:"total_amount"`
            ExtensionOrderID string  `json:"extension_order_id"`
            AggregatorOrderID string  `json:"aggregator_order_id"`
            ShipmentDetails []RepaymentRequestDetails  `json:"shipment_details"`
            AggregatorTransactionID string  `json:"aggregator_transaction_id"`
         
    }
    
    // RepaymentResponse used by Payment
    type RepaymentResponse struct {

        
            Data map[string]interface{}  `json:"data"`
            Success bool  `json:"success"`
         
    }
    
    // MerchantOnBoardingRequest used by Payment
    type MerchantOnBoardingRequest struct {

        
            AppID string  `json:"app_id"`
            Status string  `json:"status"`
            UserID string  `json:"user_id"`
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

        
            TransactionAmountInPaise float64  `json:"transaction_amount_in_paise"`
            OrderItems []map[string]interface{}  `json:"order_items"`
            PhoneNumber string  `json:"phone_number"`
            Payload string  `json:"payload"`
            MerchantParams map[string]interface{}  `json:"merchant_params"`
            Aggregator string  `json:"aggregator"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
         
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
            PollingTimeout float64  `json:"polling_timeout"`
            Message string  `json:"message"`
            Amount float64  `json:"amount"`
            StatusCode float64  `json:"status_code"`
            ExternalOrderID string  `json:"external_order_id"`
            MerchantName string  `json:"merchant_name"`
            Success bool  `json:"success"`
            PaymentLinkCurrentStatus string  `json:"payment_link_current_status"`
         
    }
    
    // ErrorDescription used by Payment
    type ErrorDescription struct {

        
            PaymentTransactionID string  `json:"payment_transaction_id"`
            Msg string  `json:"msg"`
            Amount float64  `json:"amount"`
            MerchantOrderID string  `json:"merchant_order_id"`
            Cancelled bool  `json:"cancelled"`
            MerchantName string  `json:"merchant_name"`
            Expired bool  `json:"expired"`
            InvalidID bool  `json:"invalid_id"`
         
    }
    
    // ErrorResponse used by Payment
    type ErrorResponse struct {

        
            Message string  `json:"message"`
            Error ErrorDescription  `json:"error"`
            Success bool  `json:"success"`
            StatusCode float64  `json:"status_code"`
         
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

        
            Amount float64  `json:"amount"`
            ExternalOrderID string  `json:"external_order_id"`
            MobileNumber string  `json:"mobile_number"`
            Email string  `json:"email"`
            Description string  `json:"description"`
            Meta CreatePaymentLinkMeta  `json:"meta"`
         
    }
    
    // CreatePaymentLinkResponse used by Payment
    type CreatePaymentLinkResponse struct {

        
            PaymentLinkURL string  `json:"payment_link_url"`
            PollingTimeout float64  `json:"polling_timeout"`
            Message string  `json:"message"`
            StatusCode float64  `json:"status_code"`
            PaymentLinkID string  `json:"payment_link_id"`
            Success bool  `json:"success"`
         
    }
    
    // PollingPaymentLinkResponse used by Payment
    type PollingPaymentLinkResponse struct {

        
            AggregatorName string  `json:"aggregator_name"`
            HttpStatus float64  `json:"http_status"`
            Status string  `json:"status"`
            RedirectURL string  `json:"redirect_url"`
            Message string  `json:"message"`
            Amount float64  `json:"amount"`
            StatusCode float64  `json:"status_code"`
            OrderID string  `json:"order_id"`
            PaymentLinkID string  `json:"payment_link_id"`
            Success bool  `json:"success"`
         
    }
    
    // CancelOrResendPaymentLinkRequest used by Payment
    type CancelOrResendPaymentLinkRequest struct {

        
            PaymentLinkID string  `json:"payment_link_id"`
         
    }
    
    // ResendPaymentLinkResponse used by Payment
    type ResendPaymentLinkResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
            StatusCode float64  `json:"status_code"`
            PollingTimeout float64  `json:"polling_timeout"`
         
    }
    
    // CancelPaymentLinkResponse used by Payment
    type CancelPaymentLinkResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
            StatusCode float64  `json:"status_code"`
         
    }
    
    // Code used by Payment
    type Code struct {

        
            Name string  `json:"name"`
            MerchantCode string  `json:"merchant_code"`
            Code string  `json:"code"`
         
    }
    
    // PaymentCode used by Payment
    type PaymentCode struct {

        
            Name string  `json:"name"`
            Codes Code  `json:"codes"`
            Networks string  `json:"networks"`
            Types string  `json:"types"`
         
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
    
    // ShipmentItemFulFillingStore used by Order
    type ShipmentItemFulFillingStore struct {

        
            Code string  `json:"code"`
            LocationType string  `json:"location_type"`
            ID float64  `json:"id"`
            Meta map[string]interface{}  `json:"meta"`
            State string  `json:"state"`
            BrandStoreTags string  `json:"brand_store_tags"`
            Pincode string  `json:"pincode"`
            Phone string  `json:"phone"`
            City string  `json:"city"`
            Name string  `json:"name"`
            Address string  `json:"address"`
            StoreEmail string  `json:"store_email"`
         
    }
    
    // UserDataInfo used by Order
    type UserDataInfo struct {

        
            IsAnonymousUser bool  `json:"is_anonymous_user"`
            UID float64  `json:"uid"`
            Gender string  `json:"gender"`
            Mobile string  `json:"mobile"`
            AvisUserID string  `json:"avis_user_id"`
            LastName string  `json:"last_name"`
            Email string  `json:"email"`
            Name string  `json:"name"`
            FirstName string  `json:"first_name"`
         
    }
    
    // PlatformDeliveryAddress used by Order
    type PlatformDeliveryAddress struct {

        
            ContactPerson string  `json:"contact_person"`
            Latitude float64  `json:"latitude"`
            State string  `json:"state"`
            AddressCategory string  `json:"address_category"`
            Pincode string  `json:"pincode"`
            Landmark string  `json:"landmark"`
            Longitude float64  `json:"longitude"`
            UpdatedAt string  `json:"updated_at"`
            Phone string  `json:"phone"`
            Email string  `json:"email"`
            City string  `json:"city"`
            Version string  `json:"version"`
            Area string  `json:"area"`
            CreatedAt string  `json:"created_at"`
            Address1 string  `json:"address1"`
            AddressType string  `json:"address_type"`
            Country string  `json:"country"`
            Address2 string  `json:"address2"`
         
    }
    
    // Prices used by Order
    type Prices struct {

        
            PmPriceSplit float64  `json:"pm_price_split"`
            CashbackApplied float64  `json:"cashback_applied"`
            PriceEffective float64  `json:"price_effective"`
            PromotionEffectiveDiscount float64  `json:"promotion_effective_discount"`
            PriceMarked float64  `json:"price_marked"`
            RefundCredit float64  `json:"refund_credit"`
            Cashback float64  `json:"cashback"`
            GiftPrice float64  `json:"gift_price"`
            DeliveryCharge float64  `json:"delivery_charge"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            AmountPaid float64  `json:"amount_paid"`
            FyndCredits float64  `json:"fynd_credits"`
            CodCharges float64  `json:"cod_charges"`
            RefundAmount float64  `json:"refund_amount"`
            TransferPrice float64  `json:"transfer_price"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            CouponValue float64  `json:"coupon_value"`
            Discount float64  `json:"discount"`
            ValueOfGood float64  `json:"value_of_good"`
            AmountPaidRoundoff float64  `json:"amount_paid_roundoff"`
         
    }
    
    // ShipmentListingChannel used by Order
    type ShipmentListingChannel struct {

        
            Name string  `json:"name"`
            ChannelShipmentID string  `json:"channel_shipment_id"`
            Logo string  `json:"logo"`
            IsAffiliate bool  `json:"is_affiliate"`
         
    }
    
    // BagStateMapper used by Order
    type BagStateMapper struct {

        
            NotifyCustomer bool  `json:"notify_customer"`
            DisplayName string  `json:"display_name"`
            StateType string  `json:"state_type"`
            ID float64  `json:"id"`
            IsActive bool  `json:"is_active"`
            AppFacing bool  `json:"app_facing"`
            JourneyType string  `json:"journey_type"`
            AppDisplayName string  `json:"app_display_name"`
            AppStateName string  `json:"app_state_name"`
            Name string  `json:"name"`
         
    }
    
    // BagStatusHistory used by Order
    type BagStatusHistory struct {

        
            BshID float64  `json:"bsh_id"`
            StateType string  `json:"state_type"`
            DisplayName string  `json:"display_name"`
            Forward bool  `json:"forward"`
            KafkaSync bool  `json:"kafka_sync"`
            DeliveryAwbNumber string  `json:"delivery_awb_number"`
            UpdatedAt string  `json:"updated_at"`
            BagStateMapper BagStateMapper  `json:"bag_state_mapper"`
            ShipmentID string  `json:"shipment_id"`
            Reasons []map[string]interface{}  `json:"reasons"`
            Status string  `json:"status"`
            BagID float64  `json:"bag_id"`
            CreatedAt string  `json:"created_at"`
            AppDisplayName string  `json:"app_display_name"`
            StateID float64  `json:"state_id"`
            StoreID float64  `json:"store_id"`
            DeliveryPartnerID float64  `json:"delivery_partner_id"`
         
    }
    
    // PlatformArticleAttributes used by Order
    type PlatformArticleAttributes struct {

        
            Currency string  `json:"currency"`
         
    }
    
    // PlatformItem used by Order
    type PlatformItem struct {

        
            BrandID float64  `json:"brand_id"`
            DepartmentID float64  `json:"department_id"`
            L2Category []string  `json:"l2_category"`
            SlugKey string  `json:"slug_key"`
            Image []string  `json:"image"`
            L3CategoryName string  `json:"l3_category_name"`
            LastUpdatedAt string  `json:"last_updated_at"`
            Images []string  `json:"images"`
            CanCancel bool  `json:"can_cancel"`
            Size string  `json:"size"`
            Color string  `json:"color"`
            ID float64  `json:"id"`
            CanReturn bool  `json:"can_return"`
            BranchURL string  `json:"branch_url"`
            Name string  `json:"name"`
            Code string  `json:"code"`
            Meta map[string]interface{}  `json:"meta"`
            L1Category []string  `json:"l1_category"`
            Attributes PlatformArticleAttributes  `json:"attributes"`
            Brand string  `json:"brand"`
            L3Category float64  `json:"l3_category"`
         
    }
    
    // Dates used by Order
    type Dates struct {

        
            OrderCreated string  `json:"order_created"`
            DeliveryDate string  `json:"delivery_date"`
         
    }
    
    // Identifier used by Order
    type Identifier struct {

        
            Alu string  `json:"alu"`
            Upc string  `json:"upc"`
            Ean string  `json:"ean"`
            SkuCode string  `json:"sku_code"`
            Isbn string  `json:"isbn"`
         
    }
    
    // FinancialBreakup used by Order
    type FinancialBreakup struct {

        
            TotalUnits float64  `json:"total_units"`
            CashbackApplied float64  `json:"cashback_applied"`
            PriceEffective float64  `json:"price_effective"`
            PromotionEffectiveDiscount float64  `json:"promotion_effective_discount"`
            PriceMarked float64  `json:"price_marked"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            RefundCredit float64  `json:"refund_credit"`
            ItemName string  `json:"item_name"`
            Cashback float64  `json:"cashback"`
            Size string  `json:"size"`
            HsnCode string  `json:"hsn_code"`
            AddedToFyndCash bool  `json:"added_to_fynd_cash"`
            DeliveryCharge float64  `json:"delivery_charge"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            AmountPaid float64  `json:"amount_paid"`
            FyndCredits float64  `json:"fynd_credits"`
            GstFee float64  `json:"gst_fee"`
            CodCharges float64  `json:"cod_charges"`
            TransferPrice float64  `json:"transfer_price"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            CouponValue float64  `json:"coupon_value"`
            Identifiers Identifier  `json:"identifiers"`
            Discount float64  `json:"discount"`
            ValueOfGood float64  `json:"value_of_good"`
            AmountPaidRoundoff float64  `json:"amount_paid_roundoff"`
            GstTag string  `json:"gst_tag"`
         
    }
    
    // BagReturnableCancelableStatus used by Order
    type BagReturnableCancelableStatus struct {

        
            CanBeCancelled bool  `json:"can_be_cancelled"`
            IsCustomerReturnAllowed bool  `json:"is_customer_return_allowed"`
            EnableTracking bool  `json:"enable_tracking"`
            IsActive bool  `json:"is_active"`
            IsReturnable bool  `json:"is_returnable"`
         
    }
    
    // Weight used by Order
    type Weight struct {

        
            Unit string  `json:"unit"`
            Shipping float64  `json:"shipping"`
            IsDefault bool  `json:"is_default"`
         
    }
    
    // Dimensions used by Order
    type Dimensions struct {

        
            Height float64  `json:"height"`
            IsDefault bool  `json:"is_default"`
            Width float64  `json:"width"`
            Length float64  `json:"length"`
            Unit string  `json:"unit"`
         
    }
    
    // ReturnConfig used by Order
    type ReturnConfig struct {

        
            Unit string  `json:"unit"`
            Time float64  `json:"time"`
            Returnable bool  `json:"returnable"`
         
    }
    
    // Article used by Order
    type Article struct {

        
            Code string  `json:"code"`
            RawMeta string  `json:"raw_meta"`
            ID string  `json:"_id"`
            UID string  `json:"uid"`
            ChildDetails map[string]interface{}  `json:"child_details"`
            Weight Weight  `json:"weight"`
            Currency map[string]interface{}  `json:"currency"`
            Dimensions Dimensions  `json:"dimensions"`
            Identifiers map[string]interface{}  `json:"identifiers"`
            EspModified bool  `json:"esp_modified"`
            IsSet bool  `json:"is_set"`
            SellerIdentifier string  `json:"seller_identifier"`
            ASet map[string]interface{}  `json:"a_set"`
            Size string  `json:"size"`
            ReturnConfig ReturnConfig  `json:"return_config"`
         
    }
    
    // ReplacementDetails used by Order
    type ReplacementDetails struct {

        
            OriginalAffiliateOrderID string  `json:"original_affiliate_order_id"`
            ReplacementType string  `json:"replacement_type"`
         
    }
    
    // AffiliateMeta used by Order
    type AffiliateMeta struct {

        
            BoxType string  `json:"box_type"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            SizeLevelTotalQty float64  `json:"size_level_total_qty"`
            IsPriority bool  `json:"is_priority"`
            EmployeeDiscount float64  `json:"employee_discount"`
            OrderItemID string  `json:"order_item_id"`
            MarketplaceInvoiceID string  `json:"marketplace_invoice_id"`
            ReplacementDetails ReplacementDetails  `json:"replacement_details"`
            DueDate string  `json:"due_date"`
            ChannelShipmentID string  `json:"channel_shipment_id"`
            Quantity float64  `json:"quantity"`
            CouponCode string  `json:"coupon_code"`
            ChannelOrderID string  `json:"channel_order_id"`
         
    }
    
    // AffiliateBagDetails used by Order
    type AffiliateBagDetails struct {

        
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
            EmployeeDiscount float64  `json:"employee_discount"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            AffiliateMeta AffiliateMeta  `json:"affiliate_meta"`
         
    }
    
    // GSTDetailsData used by Order
    type GSTDetailsData struct {

        
            CgstGstFee string  `json:"cgst_gst_fee"`
            HsnCode string  `json:"hsn_code"`
            SgstGstFee string  `json:"sgst_gst_fee"`
            GstinCode string  `json:"gstin_code"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            HsnCodeID string  `json:"hsn_code_id"`
            IgstTaxPercentage float64  `json:"igst_tax_percentage"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            IgstGstFee string  `json:"igst_gst_fee"`
            IsDefaultHsnCode bool  `json:"is_default_hsn_code"`
            GstFee float64  `json:"gst_fee"`
            CgstTaxPercentage float64  `json:"cgst_tax_percentage"`
            ValueOfGood float64  `json:"value_of_good"`
            SgstTaxPercentage float64  `json:"sgst_tax_percentage"`
            GstTag string  `json:"gst_tag"`
         
    }
    
    // ShipmentListingBrand used by Order
    type ShipmentListingBrand struct {

        
            Logo string  `json:"logo"`
            Name string  `json:"name"`
            CreatedOn string  `json:"created_on"`
            LogoBase64 string  `json:"logo_base64"`
         
    }
    
    // BagUnit used by Order
    type BagUnit struct {

        
            ProductQuantity float64  `json:"product_quantity"`
            CurrentOperationalStatus BagStatusHistory  `json:"current_operational_status"`
            Item PlatformItem  `json:"item"`
            Dates Dates  `json:"dates"`
            FinancialBreakup []FinancialBreakup  `json:"financial_breakup"`
            CurrentStatus BagStatusHistory  `json:"current_status"`
            CanCancel bool  `json:"can_cancel"`
            BagType string  `json:"bag_type"`
            Reasons []map[string]interface{}  `json:"reasons"`
            BagID float64  `json:"bag_id"`
            Size string  `json:"size"`
            Status BagReturnableCancelableStatus  `json:"status"`
            Article Article  `json:"article"`
            DisplayName string  `json:"display_name"`
            CanReturn bool  `json:"can_return"`
            Prices Prices  `json:"prices"`
            AffiliateBagDetails AffiliateBagDetails  `json:"affiliate_bag_details"`
            Gst GSTDetailsData  `json:"gst"`
            LineNumber float64  `json:"line_number"`
            EntityType string  `json:"entity_type"`
            Meta map[string]interface{}  `json:"meta"`
            BagStatus []BagStatusHistory  `json:"bag_status"`
            BagExpiryDate string  `json:"bag_expiry_date"`
            Brand ShipmentListingBrand  `json:"brand"`
         
    }
    
    // ShipmentStatus used by Order
    type ShipmentStatus struct {

        
            Meta map[string]interface{}  `json:"meta"`
            Title string  `json:"title"`
            ShipmentID string  `json:"shipment_id"`
            BagList []string  `json:"bag_list"`
            CreatedAt string  `json:"created_at"`
            CurrentShipmentStatus string  `json:"current_shipment_status"`
            StatusCreatedAt string  `json:"status_created_at"`
            ShipmentStatusID float64  `json:"shipment_status_id"`
            Status string  `json:"status"`
         
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

        
            DisplayText string  `json:"display_text"`
            Slug string  `json:"slug"`
            EntityType string  `json:"entity_type"`
         
    }
    
    // LockData used by Order
    type LockData struct {

        
            Locked bool  `json:"locked"`
            Mto bool  `json:"mto"`
            LockMessage string  `json:"lock_message"`
         
    }
    
    // ShipmentItemMeta used by Order
    type ShipmentItemMeta struct {

        
            OrderType string  `json:"order_type"`
            PdfMedia []map[string]interface{}  `json:"pdf_media"`
            ExistingDpList []string  `json:"existing_dp_list"`
            Formatted Formatted  `json:"formatted"`
            Timestamp ShipmentTimeStamp  `json:"timestamp"`
            PackagingName string  `json:"packaging_name"`
            FulfilmentPriorityText string  `json:"fulfilment_priority_text"`
            ShipmentTags []ShipmentTags  `json:"shipment_tags"`
            AutoTriggerDpAssignmentAcf bool  `json:"auto_trigger_dp_assignment_acf"`
            DpSortKey string  `json:"dp_sort_key"`
            External map[string]interface{}  `json:"external"`
            ShipmentVolumetricWeight float64  `json:"shipment_volumetric_weight"`
            Sla float64  `json:"sla"`
            ParentDpID string  `json:"parent_dp_id"`
            Tags []map[string]interface{}  `json:"tags"`
            LockData LockData  `json:"lock_data"`
            ShippingZone string  `json:"shipping_zone"`
            BagWeight map[string]interface{}  `json:"bag_weight"`
            DpOptions map[string]interface{}  `json:"dp_options"`
            EwaybillInfo map[string]interface{}  `json:"ewaybill_info"`
            Weight float64  `json:"weight"`
            ShipmentWeight float64  `json:"shipment_weight"`
            ActivityComment string  `json:"activity_comment"`
            AssignDpFromSb bool  `json:"assign_dp_from_sb"`
            StoreInvoiceUpdatedDate string  `json:"store_invoice_updated_date"`
            SameStoreAvailable bool  `json:"same_store_available"`
            ShipmentChargeableWeight float64  `json:"shipment_chargeable_weight"`
            DebugInfo map[string]interface{}  `json:"debug_info"`
            IsInternational bool  `json:"is_international"`
         
    }
    
    // ShipmentItem used by Order
    type ShipmentItem struct {

        
            ShipmentCreatedAt string  `json:"shipment_created_at"`
            CustomerNote string  `json:"customer_note"`
            FulfillingStore ShipmentItemFulFillingStore  `json:"fulfilling_store"`
            OrderDate string  `json:"order_date"`
            PaymentMethods map[string]interface{}  `json:"payment_methods"`
            User UserDataInfo  `json:"user"`
            ShipmentID string  `json:"shipment_id"`
            StatusCreatedAt string  `json:"status_created_at"`
            OrderID string  `json:"order_id"`
            PreviousShipmentID string  `json:"previous_shipment_id"`
            DisplayName string  `json:"display_name"`
            DeliveryAddress PlatformDeliveryAddress  `json:"delivery_address"`
            Prices Prices  `json:"prices"`
            TotalBags float64  `json:"total_bags"`
            Channel ShipmentListingChannel  `json:"channel"`
            Bags []BagUnit  `json:"bags"`
            ShipmentStatus ShipmentStatus  `json:"shipment_status"`
            Meta ShipmentItemMeta  `json:"meta"`
            InvoiceID string  `json:"invoice_id"`
            EstimatedSlaTime string  `json:"estimated_sla_time"`
            PaymentMode string  `json:"payment_mode"`
            LockStatus bool  `json:"lock_status"`
            CanProcess bool  `json:"can_process"`
            OrderingChannnel string  `json:"ordering_channnel"`
         
    }
    
    // ShipmentInternalPlatformViewResponse used by Order
    type ShipmentInternalPlatformViewResponse struct {

        
            Lane string  `json:"lane"`
            TotalCount float64  `json:"total_count"`
            Message string  `json:"message"`
            Success bool  `json:"success"`
            Page Page  `json:"page"`
            Items []ShipmentItem  `json:"items"`
         
    }
    
    // Error used by Order
    type Error struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // FulfillingStore used by Order
    type FulfillingStore struct {

        
            Code string  `json:"code"`
            ContactPerson string  `json:"contact_person"`
            Meta map[string]interface{}  `json:"meta"`
            ID float64  `json:"id"`
            FulfillmentChannel string  `json:"fulfillment_channel"`
            State string  `json:"state"`
            Pincode string  `json:"pincode"`
            Phone string  `json:"phone"`
            City string  `json:"city"`
            Address string  `json:"address"`
            StoreName string  `json:"store_name"`
            Country string  `json:"country"`
         
    }
    
    // InvoiceInfo used by Order
    type InvoiceInfo struct {

        
            InvoiceURL string  `json:"invoice_url"`
            UpdatedDate string  `json:"updated_date"`
            LabelURL string  `json:"label_url"`
            CreditNoteID string  `json:"credit_note_id"`
            StoreInvoiceID string  `json:"store_invoice_id"`
            ExternalInvoiceID string  `json:"external_invoice_id"`
         
    }
    
    // UserDetailsData used by Order
    type UserDetailsData struct {

        
            State string  `json:"state"`
            Pincode string  `json:"pincode"`
            Landmark string  `json:"landmark"`
            Phone string  `json:"phone"`
            Email string  `json:"email"`
            Name string  `json:"name"`
            City string  `json:"city"`
            Area string  `json:"area"`
            Address string  `json:"address"`
            Address1 string  `json:"address1"`
            AddressType string  `json:"address_type"`
            Country string  `json:"country"`
         
    }
    
    // ShipmentPayments used by Order
    type ShipmentPayments struct {

        
            Source string  `json:"source"`
            Logo string  `json:"logo"`
            Mode string  `json:"mode"`
         
    }
    
    // DPDetailsData used by Order
    type DPDetailsData struct {

        
            ID float64  `json:"id"`
            Pincode string  `json:"pincode"`
            TrackURL string  `json:"track_url"`
            Name string  `json:"name"`
            EwayBillID string  `json:"eway_bill_id"`
            Country string  `json:"country"`
            AwbNo string  `json:"awb_no"`
            GstTag string  `json:"gst_tag"`
         
    }
    
    // ShipmentStatusData used by Order
    type ShipmentStatusData struct {

        
            DisplayName string  `json:"display_name"`
            Meta map[string]interface{}  `json:"meta"`
            ID float64  `json:"id"`
            ShipmentID string  `json:"shipment_id"`
            BagList []string  `json:"bag_list"`
            CreatedAt string  `json:"created_at"`
            Status string  `json:"status"`
         
    }
    
    // OrderingStoreDetails used by Order
    type OrderingStoreDetails struct {

        
            Code string  `json:"code"`
            ContactPerson string  `json:"contact_person"`
            Meta map[string]interface{}  `json:"meta"`
            ID float64  `json:"id"`
            State string  `json:"state"`
            Pincode string  `json:"pincode"`
            Phone string  `json:"phone"`
            City string  `json:"city"`
            Address string  `json:"address"`
            StoreName string  `json:"store_name"`
            Country string  `json:"country"`
         
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

        
            CompanyGst string  `json:"company_gst"`
            CompanyID float64  `json:"company_id"`
            CompanyContact ContactDetails  `json:"company_contact"`
            Address map[string]interface{}  `json:"address"`
            CompanyName string  `json:"company_name"`
            CompanyCin string  `json:"company_cin"`
         
    }
    
    // PDFLinks used by Order
    type PDFLinks struct {

        
            InvoiceExport string  `json:"invoice_export"`
            DeliveryChallanA4 string  `json:"delivery_challan_a4"`
            LabelA6 string  `json:"label_a6"`
            LabelExport string  `json:"label_export"`
            InvoicePos string  `json:"invoice_pos"`
            LabelType string  `json:"label_type"`
            InvoiceType string  `json:"invoice_type"`
            CreditNoteURL string  `json:"credit_note_url"`
            LabelA4 string  `json:"label_a4"`
            Label string  `json:"label"`
            Invoice string  `json:"invoice"`
            PoInvoice string  `json:"po_invoice"`
            B2b string  `json:"b2b"`
            InvoiceA6 string  `json:"invoice_a6"`
            LabelPos string  `json:"label_pos"`
            InvoiceA4 string  `json:"invoice_a4"`
         
    }
    
    // EinvoiceInfo used by Order
    type EinvoiceInfo struct {

        
            Invoice map[string]interface{}  `json:"invoice"`
            CreditNote map[string]interface{}  `json:"credit_note"`
         
    }
    
    // BuyerDetails used by Order
    type BuyerDetails struct {

        
            Gstin string  `json:"gstin"`
            State string  `json:"state"`
            Pincode float64  `json:"pincode"`
            Name string  `json:"name"`
            City string  `json:"city"`
            AjioSiteID string  `json:"ajio_site_id"`
            Address string  `json:"address"`
         
    }
    
    // DebugInfo used by Order
    type DebugInfo struct {

        
            StormbreakerUUID string  `json:"stormbreaker_uuid"`
         
    }
    
    // ShipmentMeta used by Order
    type ShipmentMeta struct {

        
            B2cBuyerDetails map[string]interface{}  `json:"b2c_buyer_details"`
            OrderType string  `json:"order_type"`
            ReturnDetails map[string]interface{}  `json:"return_details"`
            EinvoiceInfo EinvoiceInfo  `json:"einvoice_info"`
            B2bBuyerDetails BuyerDetails  `json:"b2b_buyer_details"`
            MarketplaceStoreID string  `json:"marketplace_store_id"`
            ForwardAffiliateOrderID string  `json:"forward_affiliate_order_id"`
            AwbNumber string  `json:"awb_number"`
            Formatted Formatted  `json:"formatted"`
            Timestamp ShipmentTimeStamp  `json:"timestamp"`
            PackagingName string  `json:"packaging_name"`
            ReturnAffiliateOrderID string  `json:"return_affiliate_order_id"`
            FulfilmentPriorityText string  `json:"fulfilment_priority_text"`
            ShipmentTags []ShipmentTags  `json:"shipment_tags"`
            AutoTriggerDpAssignmentAcf bool  `json:"auto_trigger_dp_assignment_acf"`
            ReturnAffiliateShipmentID string  `json:"return_affiliate_shipment_id"`
            PoNumber string  `json:"po_number"`
            DpSortKey string  `json:"dp_sort_key"`
            External map[string]interface{}  `json:"external"`
            DpName string  `json:"dp_name"`
            ReturnAwbNumber string  `json:"return_awb_number"`
            ShipmentVolumetricWeight float64  `json:"shipment_volumetric_weight"`
            DueDate string  `json:"due_date"`
            DpID string  `json:"dp_id"`
            ParentDpID string  `json:"parent_dp_id"`
            LockData LockData  `json:"lock_data"`
            BoxType string  `json:"box_type"`
            ForwardAffiliateShipmentID string  `json:"forward_affiliate_shipment_id"`
            BagWeight map[string]interface{}  `json:"bag_weight"`
            ReturnStoreNode float64  `json:"return_store_node"`
            DpOptions map[string]interface{}  `json:"dp_options"`
            Weight float64  `json:"weight"`
            EwaybillInfo map[string]interface{}  `json:"ewaybill_info"`
            ShipmentWeight float64  `json:"shipment_weight"`
            AssignDpFromSb bool  `json:"assign_dp_from_sb"`
            StoreInvoiceUpdatedDate string  `json:"store_invoice_updated_date"`
            Dimension Dimensions  `json:"dimension"`
            SameStoreAvailable bool  `json:"same_store_available"`
            DebugInfo DebugInfo  `json:"debug_info"`
         
    }
    
    // AffiliateDetails used by Order
    type AffiliateDetails struct {

        
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            AffiliateID string  `json:"affiliate_id"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            AffiliateStoreID string  `json:"affiliate_store_id"`
            AdID string  `json:"ad_id"`
            PdfLinks PDFLinks  `json:"pdf_links"`
            ShipmentMeta ShipmentMeta  `json:"shipment_meta"`
            CompanyAffiliateTag string  `json:"company_affiliate_tag"`
            AffiliateMeta AffiliateMeta  `json:"affiliate_meta"`
         
    }
    
    // BagPaymentMethods used by Order
    type BagPaymentMethods struct {

        
            Amount float64  `json:"amount"`
            Mode string  `json:"mode"`
         
    }
    
    // CurrentStatus used by Order
    type CurrentStatus struct {

        
            StateType string  `json:"state_type"`
            ID float64  `json:"id"`
            DeliveryAwbNumber string  `json:"delivery_awb_number"`
            KafkaSync bool  `json:"kafka_sync"`
            UpdatedAt float64  `json:"updated_at"`
            BagStateMapper BagStateMapper  `json:"bag_state_mapper"`
            ShipmentID string  `json:"shipment_id"`
            CreatedAt string  `json:"created_at"`
            Status string  `json:"status"`
            BagID float64  `json:"bag_id"`
            StateID float64  `json:"state_id"`
            StoreID float64  `json:"store_id"`
            DeliveryPartnerID float64  `json:"delivery_partner_id"`
         
    }
    
    // OrderBagArticle used by Order
    type OrderBagArticle struct {

        
            Identifiers map[string]interface{}  `json:"identifiers"`
            ReturnConfig ReturnConfig  `json:"return_config"`
            Size string  `json:"size"`
            UID string  `json:"uid"`
         
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

        
            PromotionName string  `json:"promotion_name"`
            PromoID string  `json:"promo_id"`
            Amount float64  `json:"amount"`
            BuyRules []BuyRules  `json:"buy_rules"`
            MrpPromotion bool  `json:"mrp_promotion"`
            PromotionType string  `json:"promotion_type"`
            ArticleQuantity float64  `json:"article_quantity"`
            DiscountRules []DiscountRules  `json:"discount_rules"`
         
    }
    
    // AffiliateBagsDetails used by Order
    type AffiliateBagsDetails struct {

        
            AffiliateBagID string  `json:"affiliate_bag_id"`
            CouponCode string  `json:"coupon_code"`
         
    }
    
    // BagGST used by Order
    type BagGST struct {

        
            HsnCode string  `json:"hsn_code"`
            CgstGstFee string  `json:"cgst_gst_fee"`
            SgstGstFee string  `json:"sgst_gst_fee"`
            GstinCode string  `json:"gstin_code"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            HsnCodeID string  `json:"hsn_code_id"`
            IgstTaxPercentage float64  `json:"igst_tax_percentage"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            IgstGstFee string  `json:"igst_gst_fee"`
            IsDefaultHsnCode bool  `json:"is_default_hsn_code"`
            GstFee float64  `json:"gst_fee"`
            CgstTaxPercentage float64  `json:"cgst_tax_percentage"`
            ValueOfGood float64  `json:"value_of_good"`
            SgstTaxPercentage float64  `json:"sgst_tax_percentage"`
            GstTag string  `json:"gst_tag"`
         
    }
    
    // BagConfigs used by Order
    type BagConfigs struct {

        
            CanBeCancelled bool  `json:"can_be_cancelled"`
            IsCustomerReturnAllowed bool  `json:"is_customer_return_allowed"`
            EnableTracking bool  `json:"enable_tracking"`
            IsActive bool  `json:"is_active"`
            IsReturnable bool  `json:"is_returnable"`
            AllowForceReturn bool  `json:"allow_force_return"`
         
    }
    
    // GiftCard used by Order
    type GiftCard struct {

        
            IsGiftApplied bool  `json:"is_gift_applied"`
            GiftMessage string  `json:"gift_message"`
            DisplayText string  `json:"display_text"`
            GiftPrice float64  `json:"gift_price"`
         
    }
    
    // B2BPODetails used by Order
    type B2BPODetails struct {

        
            DockerNumber string  `json:"docker_number"`
            ItemBasePrice float64  `json:"item_base_price"`
            PoLineAmount float64  `json:"po_line_amount"`
            PoTaxAmount float64  `json:"po_tax_amount"`
            TotalGstPercentage float64  `json:"total_gst_percentage"`
            PartialCanRet bool  `json:"partial_can_ret"`
         
    }
    
    // BagMeta used by Order
    type BagMeta struct {

        
            GiftCard GiftCard  `json:"gift_card"`
            B2bPoDetails B2BPODetails  `json:"b2b_po_details"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            PartialCanRet bool  `json:"partial_can_ret"`
            CustomJson map[string]interface{}  `json:"custom_json"`
            CustomMessage string  `json:"custom_message"`
            GroupID string  `json:"group_id"`
            DocketNumber string  `json:"docket_number"`
         
    }
    
    // OrderBrandName used by Order
    type OrderBrandName struct {

        
            BrandName string  `json:"brand_name"`
            ID float64  `json:"id"`
            CreatedOn string  `json:"created_on"`
            Company float64  `json:"company"`
            ModifiedOn string  `json:"modified_on"`
            Logo string  `json:"logo"`
         
    }
    
    // OrderBags used by Order
    type OrderBags struct {

        
            Item PlatformItem  `json:"item"`
            PaymentMethods []BagPaymentMethods  `json:"payment_methods"`
            ParentPromoBags map[string]interface{}  `json:"parent_promo_bags"`
            FinancialBreakup FinancialBreakup  `json:"financial_breakup"`
            CurrentStatus CurrentStatus  `json:"current_status"`
            Identifier string  `json:"identifier"`
            CanCancel bool  `json:"can_cancel"`
            BagID float64  `json:"bag_id"`
            GroupID string  `json:"group_id"`
            Article OrderBagArticle  `json:"article"`
            DisplayName string  `json:"display_name"`
            DeliveryAddress PlatformDeliveryAddress  `json:"delivery_address"`
            Prices Prices  `json:"prices"`
            CanReturn bool  `json:"can_return"`
            AppliedPromos []AppliedPromos  `json:"applied_promos"`
            AffiliateBagDetails AffiliateBagsDetails  `json:"affiliate_bag_details"`
            GstDetails BagGST  `json:"gst_details"`
            BagConfigs BagConfigs  `json:"bag_configs"`
            LineNumber float64  `json:"line_number"`
            EntityType string  `json:"entity_type"`
            Meta BagMeta  `json:"meta"`
            Quantity float64  `json:"quantity"`
            SellerIdentifier string  `json:"seller_identifier"`
            Brand OrderBrandName  `json:"brand"`
         
    }
    
    // ShipmentDetails used by Order
    type ShipmentDetails struct {

        
            LockMessage string  `json:"lock_message"`
            LockStatus bool  `json:"lock_status"`
            ActionToStatus map[string]interface{}  `json:"action_to_status"`
         
    }
    
    // OrderDetailsData used by Order
    type OrderDetailsData struct {

        
            Source string  `json:"source"`
            OrderingChannel string  `json:"ordering_channel"`
            AffiliateID string  `json:"affiliate_id"`
            OrderValue string  `json:"order_value"`
            OrderingChannelLogo map[string]interface{}  `json:"ordering_channel_logo"`
            OrderDate string  `json:"order_date"`
            TaxDetails map[string]interface{}  `json:"tax_details"`
            CodCharges string  `json:"cod_charges"`
            FyndOrderID string  `json:"fynd_order_id"`
         
    }
    
    // TrackingList used by Order
    type TrackingList struct {

        
            IsPassed bool  `json:"is_passed"`
            Text string  `json:"text"`
            IsCurrent bool  `json:"is_current"`
            Status string  `json:"status"`
            Time string  `json:"time"`
         
    }
    
    // PlatformShipment used by Order
    type PlatformShipment struct {

        
            ShipmentCreatedAt string  `json:"shipment_created_at"`
            FulfillingStore FulfillingStore  `json:"fulfilling_store"`
            JourneyType string  `json:"journey_type"`
            ShipmentImages []string  `json:"shipment_images"`
            PaymentMethods map[string]interface{}  `json:"payment_methods"`
            Invoice InvoiceInfo  `json:"invoice"`
            FulfilmentPriority float64  `json:"fulfilment_priority"`
            PlatformLogo string  `json:"platform_logo"`
            User UserDataInfo  `json:"user"`
            DeliveryDetails UserDetailsData  `json:"delivery_details"`
            ShipmentID string  `json:"shipment_id"`
            UserAgent string  `json:"user_agent"`
            Payments ShipmentPayments  `json:"payments"`
            DpDetails DPDetailsData  `json:"dp_details"`
            Status ShipmentStatusData  `json:"status"`
            Vertical string  `json:"vertical"`
            Coupon map[string]interface{}  `json:"coupon"`
            OrderingStore OrderingStoreDetails  `json:"ordering_store"`
            TotalBags float64  `json:"total_bags"`
            Prices Prices  `json:"prices"`
            CompanyDetails CompanyDetails  `json:"company_details"`
            AffiliateDetails AffiliateDetails  `json:"affiliate_details"`
            DpAssignment bool  `json:"dp_assignment"`
            PickedDate string  `json:"picked_date"`
            CustomMessage string  `json:"custom_message"`
            BagStatusHistory []BagStatusHistory  `json:"bag_status_history"`
            Bags []OrderBags  `json:"bags"`
            GstDetails GSTDetailsData  `json:"gst_details"`
            ShipmentStatus string  `json:"shipment_status"`
            OperationalStatus string  `json:"operational_status"`
            CustomMeta []map[string]interface{}  `json:"custom_meta"`
            CanUpdateDimension bool  `json:"can_update_dimension"`
            Meta ShipmentMeta  `json:"meta"`
            DeliverySlot map[string]interface{}  `json:"delivery_slot"`
            PackagingType string  `json:"packaging_type"`
            TotalItems float64  `json:"total_items"`
            ShipmentQuantity float64  `json:"shipment_quantity"`
            PriorityText string  `json:"priority_text"`
            InvoiceID string  `json:"invoice_id"`
            ShipmentDetails ShipmentDetails  `json:"shipment_details"`
            IsDpAssignEnabled bool  `json:"is_dp_assign_enabled"`
            BillingDetails UserDetailsData  `json:"billing_details"`
            PdfLinks map[string]interface{}  `json:"pdf_links"`
            Order OrderDetailsData  `json:"order"`
            PaymentMode string  `json:"payment_mode"`
            EstimatedSlaTime string  `json:"estimated_sla_time"`
            ForwardShipmentID string  `json:"forward_shipment_id"`
            LockStatus bool  `json:"lock_status"`
            EnableDpTracking bool  `json:"enable_dp_tracking"`
            TrackingList []TrackingList  `json:"tracking_list"`
         
    }
    
    // ShipmentInfoResponse used by Order
    type ShipmentInfoResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
            Shipments []PlatformShipment  `json:"shipments"`
         
    }
    
    // BillingStaffDetails used by Order
    type BillingStaffDetails struct {

        
            User string  `json:"user"`
            LastName string  `json:"last_name"`
            StaffID float64  `json:"staff_id"`
            EmployeeCode string  `json:"employee_code"`
            FirstName string  `json:"first_name"`
         
    }
    
    // PlatformUserDetails used by Order
    type PlatformUserDetails struct {

        
            PlatformUserFirstName string  `json:"platform_user_first_name"`
            PlatformUserEmployeeCode string  `json:"platform_user_employee_code"`
            PlatformUserLastName string  `json:"platform_user_last_name"`
            PlatformUserID string  `json:"platform_user_id"`
         
    }
    
    // TransactionData used by Order
    type TransactionData struct {

        
            Entity string  `json:"entity"`
            TerminalID string  `json:"terminal_id"`
            TransactionID string  `json:"transaction_id"`
            AmountPaid float64  `json:"amount_paid"`
            PaymentID string  `json:"payment_id"`
            Status string  `json:"status"`
            Currency string  `json:"currency"`
            UniqueReferenceNumber string  `json:"unique_reference_number"`
         
    }
    
    // OrderMeta used by Order
    type OrderMeta struct {

        
            CustomerNote string  `json:"customer_note"`
            OrderType string  `json:"order_type"`
            OrderChildEntities []string  `json:"order_child_entities"`
            BillingStaffDetails BillingStaffDetails  `json:"billing_staff_details"`
            Staff map[string]interface{}  `json:"staff"`
            Comment string  `json:"comment"`
            EmployeeID float64  `json:"employee_id"`
            CartID float64  `json:"cart_id"`
            PaymentType string  `json:"payment_type"`
            PlatformUserDetails PlatformUserDetails  `json:"platform_user_details"`
            OrderPlatform string  `json:"order_platform"`
            Files []map[string]interface{}  `json:"files"`
            OrderingStore float64  `json:"ordering_store"`
            CompanyLogo string  `json:"company_logo"`
            CurrencySymbol string  `json:"currency_symbol"`
            TransactionData TransactionData  `json:"transaction_data"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            OrderTags []map[string]interface{}  `json:"order_tags"`
            MongoCartID float64  `json:"mongo_cart_id"`
         
    }
    
    // TaxDetails used by Order
    type TaxDetails struct {

        
            PanNo string  `json:"pan_no"`
            Gstin string  `json:"gstin"`
         
    }
    
    // OrderData used by Order
    type OrderData struct {

        
            Meta OrderMeta  `json:"meta"`
            Prices Prices  `json:"prices"`
            OrderDate string  `json:"order_date"`
            PaymentMethods map[string]interface{}  `json:"payment_methods"`
            TaxDetails TaxDetails  `json:"tax_details"`
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
            TotalItems float64  `json:"total_items"`
            Index float64  `json:"index"`
            Text string  `json:"text"`
            Actions []map[string]interface{}  `json:"actions"`
         
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
    
    // PlatformBreakupValues used by Order
    type PlatformBreakupValues struct {

        
            Display string  `json:"display"`
            Name string  `json:"name"`
            Value string  `json:"value"`
         
    }
    
    // PlatformChannel used by Order
    type PlatformChannel struct {

        
            Name string  `json:"name"`
            Logo string  `json:"logo"`
         
    }
    
    // PlatformOrderItems used by Order
    type PlatformOrderItems struct {

        
            TotalOrderValue float64  `json:"total_order_value"`
            Meta map[string]interface{}  `json:"meta"`
            BreakupValues []PlatformBreakupValues  `json:"breakup_values"`
            Channel PlatformChannel  `json:"channel"`
            OrderValue float64  `json:"order_value"`
            OrderCreatedTime string  `json:"order_created_time"`
            Shipments []PlatformShipment  `json:"shipments"`
            OrderID string  `json:"order_id"`
            PaymentMode string  `json:"payment_mode"`
            UserInfo UserDataInfo  `json:"user_info"`
         
    }
    
    // OrderListingResponse used by Order
    type OrderListingResponse struct {

        
            Lane string  `json:"lane"`
            TotalCount float64  `json:"total_count"`
            Message string  `json:"message"`
            Success bool  `json:"success"`
            Page Page  `json:"page"`
            Items []PlatformOrderItems  `json:"items"`
         
    }
    
    // PlatformTrack used by Order
    type PlatformTrack struct {

        
            LastLocationRecievedAt string  `json:"last_location_recieved_at"`
            Reason string  `json:"reason"`
            Meta map[string]interface{}  `json:"meta"`
            RawStatus string  `json:"raw_status"`
            AccountName string  `json:"account_name"`
            UpdatedAt string  `json:"updated_at"`
            Awb string  `json:"awb"`
            Status string  `json:"status"`
            ShipmentType string  `json:"shipment_type"`
            UpdatedTime string  `json:"updated_time"`
         
    }
    
    // PlatformShipmentTrack used by Order
    type PlatformShipmentTrack struct {

        
            Meta map[string]interface{}  `json:"meta"`
            Results []PlatformTrack  `json:"results"`
         
    }
    
    // FilterInfoOption used by Order
    type FilterInfoOption struct {

        
            PlaceholderText string  `json:"placeholder_text"`
            Value string  `json:"value"`
            ShowUI bool  `json:"show_ui"`
            Text string  `json:"text"`
            Name string  `json:"name"`
            MinSearchSize float64  `json:"min_search_size"`
         
    }
    
    // FiltersInfo used by Order
    type FiltersInfo struct {

        
            PlaceholderText string  `json:"placeholder_text"`
            Value string  `json:"value"`
            Options []FilterInfoOption  `json:"options"`
            Type string  `json:"type"`
            Text string  `json:"text"`
            Required bool  `json:"required"`
         
    }
    
    // AdvanceFilterInfo used by Order
    type AdvanceFilterInfo struct {

        
            Unfulfilled []FiltersInfo  `json:"unfulfilled"`
            Processed []FiltersInfo  `json:"processed"`
            Returned []FiltersInfo  `json:"returned"`
            Filters []FiltersInfo  `json:"filters"`
            ActionCentre []FiltersInfo  `json:"action_centre"`
            Page map[string]interface{}  `json:"page"`
            AppliedFilters map[string]interface{}  `json:"applied_filters"`
         
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

        
            FileName string  `json:"file_name"`
            Cdn URL  `json:"cdn"`
         
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
    
    // Attributes used by Order
    type Attributes struct {

        
            BrandName string  `json:"brand_name"`
            Gender []string  `json:"gender"`
            MarketerAddress string  `json:"marketer_address"`
            MarketerName string  `json:"marketer_name"`
            Name string  `json:"name"`
            Essential string  `json:"essential"`
            PrimaryColorHex string  `json:"primary_color_hex"`
            PrimaryMaterial string  `json:"primary_material"`
            PrimaryColor string  `json:"primary_color"`
         
    }
    
    // Item used by Order
    type Item struct {

        
            ItemID float64  `json:"item_id"`
            BrandID float64  `json:"brand_id"`
            DepartmentID float64  `json:"department_id"`
            L2Category []string  `json:"l2_category"`
            SlugKey string  `json:"slug_key"`
            Image []string  `json:"image"`
            L1CategoryID float64  `json:"l1_category_id"`
            L3CategoryName string  `json:"l3_category_name"`
            LastUpdatedAt string  `json:"last_updated_at"`
            WebstoreProductURL string  `json:"webstore_product_url"`
            Gender string  `json:"gender"`
            CanCancel bool  `json:"can_cancel"`
            L2CategoryID float64  `json:"l2_category_id"`
            Size string  `json:"size"`
            Color string  `json:"color"`
            CanReturn bool  `json:"can_return"`
            BranchURL string  `json:"branch_url"`
            Name string  `json:"name"`
            Code string  `json:"code"`
            Meta map[string]interface{}  `json:"meta"`
            L1Category []string  `json:"l1_category"`
            Attributes Attributes  `json:"attributes"`
            Brand string  `json:"brand"`
            L3Category float64  `json:"l3_category"`
         
    }
    
    // ArticleDetails used by Order
    type ArticleDetails struct {

        
            Status map[string]interface{}  `json:"status"`
         
    }
    
    // StoreAddress used by Order
    type StoreAddress struct {

        
            Latitude float64  `json:"latitude"`
            Email string  `json:"email"`
            City string  `json:"city"`
            AddressCategory string  `json:"address_category"`
            Pincode float64  `json:"pincode"`
            Phone string  `json:"phone"`
            CreatedAt string  `json:"created_at"`
            Version string  `json:"version"`
            Area string  `json:"area"`
            AddressType string  `json:"address_type"`
            Address2 string  `json:"address2"`
            ContactPerson string  `json:"contact_person"`
            State string  `json:"state"`
            UpdatedAt string  `json:"updated_at"`
            CountryCode string  `json:"country_code"`
            Landmark string  `json:"landmark"`
            Longitude float64  `json:"longitude"`
            Address1 string  `json:"address1"`
            Country string  `json:"country"`
         
    }
    
    // EInvoicePortalDetails used by Order
    type EInvoicePortalDetails struct {

        
            User string  `json:"user"`
            Password string  `json:"password"`
            Username string  `json:"username"`
         
    }
    
    // StoreEwaybill used by Order
    type StoreEwaybill struct {

        
            Enabled bool  `json:"enabled"`
         
    }
    
    // StoreEinvoice used by Order
    type StoreEinvoice struct {

        
            User string  `json:"user"`
            Password string  `json:"password"`
            Username string  `json:"username"`
            Enabled bool  `json:"enabled"`
         
    }
    
    // StoreGstCredentials used by Order
    type StoreGstCredentials struct {

        
            EWaybill StoreEwaybill  `json:"e_waybill"`
            EInvoice StoreEinvoice  `json:"e_invoice"`
         
    }
    
    // Document used by Order
    type Document struct {

        
            Value string  `json:"value"`
            URL string  `json:"url"`
            Verified bool  `json:"verified"`
            LegalName string  `json:"legal_name"`
            DsType string  `json:"ds_type"`
         
    }
    
    // StoreDocuments used by Order
    type StoreDocuments struct {

        
            Gst Document  `json:"gst"`
         
    }
    
    // StoreMeta used by Order
    type StoreMeta struct {

        
            ProductReturnConfig map[string]interface{}  `json:"product_return_config"`
            EinvoicePortalDetails EInvoicePortalDetails  `json:"einvoice_portal_details"`
            Timing []map[string]interface{}  `json:"timing"`
            DisplayName string  `json:"display_name"`
            GstNumber string  `json:"gst_number"`
            AdditionalContactDetails map[string]interface{}  `json:"additional_contact_details"`
            NotificationEmails []string  `json:"notification_emails"`
            GstCredentials StoreGstCredentials  `json:"gst_credentials"`
            Stage string  `json:"stage"`
            Documents StoreDocuments  `json:"documents"`
            EwaybillPortalDetails map[string]interface{}  `json:"ewaybill_portal_details"`
         
    }
    
    // Store used by Order
    type Store struct {

        
            LocationType string  `json:"location_type"`
            Latitude float64  `json:"latitude"`
            BrandID interface{}  `json:"brand_id"`
            City string  `json:"city"`
            MallName string  `json:"mall_name"`
            SID string  `json:"s_id"`
            IsActive bool  `json:"is_active"`
            Pincode string  `json:"pincode"`
            CompanyID float64  `json:"company_id"`
            CreatedAt string  `json:"created_at"`
            Phone float64  `json:"phone"`
            VatNo string  `json:"vat_no"`
            IsArchived bool  `json:"is_archived"`
            Address2 string  `json:"address2"`
            MallArea string  `json:"mall_area"`
            ContactPerson string  `json:"contact_person"`
            State string  `json:"state"`
            BrandStoreTags []string  `json:"brand_store_tags"`
            LoginUsername string  `json:"login_username"`
            UpdatedAt string  `json:"updated_at"`
            IsEnabledForRecon bool  `json:"is_enabled_for_recon"`
            AlohomoraUserID float64  `json:"alohomora_user_id"`
            Name string  `json:"name"`
            StoreAddressJson StoreAddress  `json:"store_address_json"`
            Code string  `json:"code"`
            Meta StoreMeta  `json:"meta"`
            FulfillmentChannel string  `json:"fulfillment_channel"`
            OrderIntegrationID string  `json:"order_integration_id"`
            PackagingMaterialCount float64  `json:"packaging_material_count"`
            Longitude float64  `json:"longitude"`
            StoreActiveFrom string  `json:"store_active_from"`
            ParentStoreID float64  `json:"parent_store_id"`
            Address1 string  `json:"address1"`
            StoreEmail string  `json:"store_email"`
            Country string  `json:"country"`
         
    }
    
    // BagGSTDetails used by Order
    type BagGSTDetails struct {

        
            HsnCode string  `json:"hsn_code"`
            CgstGstFee string  `json:"cgst_gst_fee"`
            SgstGstFee string  `json:"sgst_gst_fee"`
            GstinCode string  `json:"gstin_code"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            HsnCodeID string  `json:"hsn_code_id"`
            IgstTaxPercentage float64  `json:"igst_tax_percentage"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            IgstGstFee string  `json:"igst_gst_fee"`
            IsDefaultHsnCode bool  `json:"is_default_hsn_code"`
            GstFee float64  `json:"gst_fee"`
            CgstTaxPercentage float64  `json:"cgst_tax_percentage"`
            ValueOfGood float64  `json:"value_of_good"`
            SgstTaxPercentage float64  `json:"sgst_tax_percentage"`
            GstTag string  `json:"gst_tag"`
         
    }
    
    // Brand used by Order
    type Brand struct {

        
            BrandName string  `json:"brand_name"`
            InvoicePrefix string  `json:"invoice_prefix"`
            ScriptLastRan string  `json:"script_last_ran"`
            BrandID float64  `json:"brand_id"`
            CreatedOn float64  `json:"created_on"`
            IsVirtualInvoice bool  `json:"is_virtual_invoice"`
            Company string  `json:"company"`
            CreditNoteExpiryDays float64  `json:"credit_note_expiry_days"`
            PickupLocation string  `json:"pickup_location"`
            CreditNoteAllowed bool  `json:"credit_note_allowed"`
            StartDate string  `json:"start_date"`
            ModifiedOn float64  `json:"modified_on"`
            Logo string  `json:"logo"`
         
    }
    
    // BagDetailsPlatformResponse used by Order
    type BagDetailsPlatformResponse struct {

        
            CurrentOperationalStatus BagStatusHistory  `json:"current_operational_status"`
            Item Item  `json:"item"`
            JourneyType string  `json:"journey_type"`
            BagUpdateTime float64  `json:"bag_update_time"`
            OriginalBagList []float64  `json:"original_bag_list"`
            Dates Dates  `json:"dates"`
            RestoreCoupon bool  `json:"restore_coupon"`
            ParentPromoBags map[string]interface{}  `json:"parent_promo_bags"`
            FinancialBreakup []FinancialBreakup  `json:"financial_breakup"`
            CurrentStatus BagStatusHistory  `json:"current_status"`
            Identifier string  `json:"identifier"`
            Reasons []map[string]interface{}  `json:"reasons"`
            ShipmentID string  `json:"shipment_id"`
            Status BagReturnableCancelableStatus  `json:"status"`
            Article Article  `json:"article"`
            QcRequired interface{}  `json:"qc_required"`
            DisplayName string  `json:"display_name"`
            ID float64  `json:"id"`
            ArticleDetails ArticleDetails  `json:"article_details"`
            OrderingStore Store  `json:"ordering_store"`
            Prices Prices  `json:"prices"`
            AffiliateDetails AffiliateDetails  `json:"affiliate_details"`
            AppliedPromos []map[string]interface{}  `json:"applied_promos"`
            AffiliateBagDetails AffiliateBagDetails  `json:"affiliate_bag_details"`
            BagStatusHistory BagStatusHistory  `json:"bag_status_history"`
            GstDetails BagGSTDetails  `json:"gst_details"`
            Tags []string  `json:"tags"`
            OperationalStatus string  `json:"operational_status"`
            LineNumber float64  `json:"line_number"`
            NoOfBagsOrder float64  `json:"no_of_bags_order"`
            EntityType string  `json:"entity_type"`
            Meta BagMeta  `json:"meta"`
            OrderIntegrationID string  `json:"order_integration_id"`
            RestorePromos map[string]interface{}  `json:"restore_promos"`
            BagStatus []BagStatusHistory  `json:"bag_status"`
            Type string  `json:"type"`
            Quantity float64  `json:"quantity"`
            SellerIdentifier string  `json:"seller_identifier"`
            Brand Brand  `json:"brand"`
         
    }
    
    // ErrorResponse used by Order
    type ErrorResponse struct {

        
            Error string  `json:"error"`
            Message string  `json:"message"`
         
    }
    
    // BagsPage used by Order
    type BagsPage struct {

        
            ItemTotal float64  `json:"item_total"`
            PageType string  `json:"page_type"`
            Size float64  `json:"size"`
            Current float64  `json:"current"`
            HasNext bool  `json:"has_next"`
         
    }
    
    // GetBagsPlatformResponse used by Order
    type GetBagsPlatformResponse struct {

        
            Page BagsPage  `json:"page"`
            Items []BagDetailsPlatformResponse  `json:"items"`
         
    }
    
    // GeneratePosOrderReceiptResponse used by Order
    type GeneratePosOrderReceiptResponse struct {

        
            CustomerCnReceipt string  `json:"customer_cn_receipt"`
            PaymentReceipt string  `json:"payment_receipt"`
            OrderID string  `json:"order_id"`
            InvoiceReceipt string  `json:"invoice_receipt"`
            MerchantCnReceipt string  `json:"merchant_cn_receipt"`
            Success bool  `json:"success"`
         
    }
    
    // InvalidateShipmentCachePayload used by Order
    type InvalidateShipmentCachePayload struct {

        
            AffiliateBagIds []string  `json:"affiliate_bag_ids"`
            ShipmentIds []string  `json:"shipment_ids"`
            BagIds []string  `json:"bag_ids"`
         
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

        
            Success bool  `json:"success"`
            Status float64  `json:"status"`
            ErrorTrace string  `json:"error_trace"`
            Message string  `json:"message"`
         
    }
    
    // StoreReassign used by Order
    type StoreReassign struct {

        
            ItemID string  `json:"item_id"`
            MongoArticleID string  `json:"mongo_article_id"`
            SetID string  `json:"set_id"`
            FyndOrderID string  `json:"fynd_order_id"`
            ReasonIds []float64  `json:"reason_ids"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
            StoreID float64  `json:"store_id"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            BagID float64  `json:"bag_id"`
            AffiliateID string  `json:"affiliate_id"`
         
    }
    
    // StoreReassignResponse used by Order
    type StoreReassignResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // Entities used by Order
    type Entities struct {

        
            ID string  `json:"id"`
            ReasonText string  `json:"reason_text"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            AffiliateID string  `json:"affiliate_id"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
         
    }
    
    // UpdateShipmentLockPayload used by Order
    type UpdateShipmentLockPayload struct {

        
            Entities []Entities  `json:"entities"`
            EntityType string  `json:"entity_type"`
            ActionType string  `json:"action_type"`
            Action string  `json:"action"`
         
    }
    
    // OriginalFilter used by Order
    type OriginalFilter struct {

        
            AffiliateID string  `json:"affiliate_id"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
         
    }
    
    // Bags used by Order
    type Bags struct {

        
            AffiliateBagID string  `json:"affiliate_bag_id"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            BagID float64  `json:"bag_id"`
            IsLocked bool  `json:"is_locked"`
         
    }
    
    // CheckResponse used by Order
    type CheckResponse struct {

        
            ShipmentID string  `json:"shipment_id"`
            OriginalFilter OriginalFilter  `json:"original_filter"`
            Bags []Bags  `json:"bags"`
            IsShipmentLocked bool  `json:"is_shipment_locked"`
            IsBagLocked bool  `json:"is_bag_locked"`
            LockStatus string  `json:"lock_status"`
            Status string  `json:"status"`
            AffiliateID string  `json:"affiliate_id"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
         
    }
    
    // UpdateShipmentLockResponse used by Order
    type UpdateShipmentLockResponse struct {

        
            Success bool  `json:"success"`
            CheckResponse []CheckResponse  `json:"check_response"`
            Message string  `json:"message"`
         
    }
    
    // AnnouncementResponse used by Order
    type AnnouncementResponse struct {

        
            Title string  `json:"title"`
            PlatformName string  `json:"platform_name"`
            CreatedAt string  `json:"created_at"`
            CompanyID float64  `json:"company_id"`
            ID float64  `json:"id"`
            PlatformID string  `json:"platform_id"`
            LogoURL string  `json:"logo_url"`
            ToDatetime string  `json:"to_datetime"`
            FromDatetime string  `json:"from_datetime"`
            Description string  `json:"description"`
         
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
    
    // OrderItemDataUpdates used by Order
    type OrderItemDataUpdates struct {

        
            Data map[string]interface{}  `json:"data"`
         
    }
    
    // DataUpdates used by Order
    type DataUpdates struct {

        
            Products []ProductsDataUpdates  `json:"products"`
            Entities []EntitiesDataUpdates  `json:"entities"`
            OrderItemStatus []OrderItemDataUpdates  `json:"order_item_status"`
         
    }
    
    // ShipmentsRequest used by Order
    type ShipmentsRequest struct {

        
            Products []Products  `json:"products"`
            Reasons ReasonsData  `json:"reasons"`
            Identifier string  `json:"identifier"`
            DataUpdates DataUpdates  `json:"data_updates"`
         
    }
    
    // StatuesRequest used by Order
    type StatuesRequest struct {

        
            SplitShipment bool  `json:"split_shipment"`
            ExcludeBagsNextState string  `json:"exclude_bags_next_state"`
            Shipments []ShipmentsRequest  `json:"shipments"`
            Status string  `json:"status"`
         
    }
    
    // UpdateShipmentStatusRequest used by Order
    type UpdateShipmentStatusRequest struct {

        
            Statuses []StatuesRequest  `json:"statuses"`
            LockAfterTransition bool  `json:"lock_after_transition"`
            ForceTransition bool  `json:"force_transition"`
            UnlockBeforeTransition bool  `json:"unlock_before_transition"`
            Task bool  `json:"task"`
         
    }
    
    // ShipmentsResponse used by Order
    type ShipmentsResponse struct {

        
            Message string  `json:"message"`
            StackTrace string  `json:"stack_trace"`
            Code string  `json:"code"`
            Exception string  `json:"exception"`
            Identifier string  `json:"identifier"`
            Status float64  `json:"status"`
            Meta map[string]interface{}  `json:"meta"`
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
    
    // OrderUser used by Order
    type OrderUser struct {

        
            Address1 string  `json:"address1"`
            City string  `json:"city"`
            Pincode string  `json:"pincode"`
            Phone float64  `json:"phone"`
            LastName string  `json:"last_name"`
            FirstName string  `json:"first_name"`
            Address2 string  `json:"address2"`
            State string  `json:"state"`
            Mobile float64  `json:"mobile"`
            Email string  `json:"email"`
            Country string  `json:"country"`
         
    }
    
    // MarketPlacePdf used by Order
    type MarketPlacePdf struct {

        
            Invoice string  `json:"invoice"`
            Label string  `json:"label"`
         
    }
    
    // AffiliateBag used by Order
    type AffiliateBag struct {

        
            PriceEffective float64  `json:"price_effective"`
            Discount float64  `json:"discount"`
            Sku string  `json:"sku"`
            AvlQty float64  `json:"avl_qty"`
            TransferPrice float64  `json:"transfer_price"`
            AmountPaid float64  `json:"amount_paid"`
            StoreID float64  `json:"store_id"`
            AffiliateMeta map[string]interface{}  `json:"affiliate_meta"`
            FyndStoreID string  `json:"fynd_store_id"`
            CompanyID float64  `json:"company_id"`
            AffiliateStoreID string  `json:"affiliate_store_id"`
            ID string  `json:"_id"`
            UnitPrice float64  `json:"unit_price"`
            PriceMarked float64  `json:"price_marked"`
            ItemID float64  `json:"item_id"`
            DeliveryCharge float64  `json:"delivery_charge"`
            Identifier map[string]interface{}  `json:"identifier"`
            PdfLinks MarketPlacePdf  `json:"pdf_links"`
            HsnCodeID string  `json:"hsn_code_id"`
            SellerIdentifier string  `json:"seller_identifier"`
            ModifiedOn string  `json:"modified_on"`
            ItemSize string  `json:"item_size"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // UserData used by Order
    type UserData struct {

        
            BillingUser OrderUser  `json:"billing_user"`
            ShippingUser OrderUser  `json:"shipping_user"`
         
    }
    
    // ArticleDetails1 used by Order
    type ArticleDetails1 struct {

        
            Attributes map[string]interface{}  `json:"attributes"`
            Dimension map[string]interface{}  `json:"dimension"`
            Category map[string]interface{}  `json:"category"`
            ID string  `json:"_id"`
            Weight map[string]interface{}  `json:"weight"`
            Quantity float64  `json:"quantity"`
            BrandID float64  `json:"brand_id"`
         
    }
    
    // LocationDetails used by Order
    type LocationDetails struct {

        
            Articles []ArticleDetails1  `json:"articles"`
            FulfillmentType string  `json:"fulfillment_type"`
            FulfillmentID float64  `json:"fulfillment_id"`
         
    }
    
    // ShipmentDetails1 used by Order
    type ShipmentDetails1 struct {

        
            Articles []ArticleDetails1  `json:"articles"`
            BoxType string  `json:"box_type"`
            FulfillmentID float64  `json:"fulfillment_id"`
            DpID float64  `json:"dp_id"`
            Shipments float64  `json:"shipments"`
            Meta map[string]interface{}  `json:"meta"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
         
    }
    
    // ShipmentConfig used by Order
    type ShipmentConfig struct {

        
            ToPincode string  `json:"to_pincode"`
            PaymentMode string  `json:"payment_mode"`
            Action string  `json:"action"`
            LocationDetails LocationDetails  `json:"location_details"`
            Shipment []ShipmentDetails1  `json:"shipment"`
            Source string  `json:"source"`
            Identifier string  `json:"identifier"`
            Journey string  `json:"journey"`
         
    }
    
    // ShipmentData used by Order
    type ShipmentData struct {

        
            ShipmentData ShipmentConfig  `json:"shipment_data"`
         
    }
    
    // OrderPriority used by Order
    type OrderPriority struct {

        
            AffiliatePriorityCode string  `json:"affiliate_priority_code"`
            FulfilmentPriority float64  `json:"fulfilment_priority"`
            FulfilmentPriorityText string  `json:"fulfilment_priority_text"`
         
    }
    
    // OrderInfo used by Order
    type OrderInfo struct {

        
            ShippingAddress OrderUser  `json:"shipping_address"`
            Discount float64  `json:"discount"`
            PaymentMode string  `json:"payment_mode"`
            Coupon string  `json:"coupon"`
            Bags []AffiliateBag  `json:"bags"`
            User UserData  `json:"user"`
            BillingAddress OrderUser  `json:"billing_address"`
            Payment map[string]interface{}  `json:"payment"`
            DeliveryCharges float64  `json:"delivery_charges"`
            Shipment ShipmentData  `json:"shipment"`
            OrderPriority OrderPriority  `json:"order_priority"`
            Items map[string]interface{}  `json:"items"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            CodCharges float64  `json:"cod_charges"`
            OrderValue float64  `json:"order_value"`
         
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

        
            CreatedAt string  `json:"created_at"`
            ID string  `json:"id"`
            Name string  `json:"name"`
            UpdatedAt string  `json:"updated_at"`
            Description string  `json:"description"`
            Token string  `json:"token"`
            Meta []AffiliateAppConfigMeta  `json:"meta"`
            Secret string  `json:"secret"`
            Owner string  `json:"owner"`
         
    }
    
    // AffiliateInventoryArticleAssignmentConfig used by Order
    type AffiliateInventoryArticleAssignmentConfig struct {

        
            PostOrderReassignment bool  `json:"post_order_reassignment"`
         
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
    
    // AffiliateInventoryLogisticsConfig used by Order
    type AffiliateInventoryLogisticsConfig struct {

        
            DpAssignment bool  `json:"dp_assignment"`
         
    }
    
    // AffiliateInventoryConfig used by Order
    type AffiliateInventoryConfig struct {

        
            ArticleAssignment AffiliateInventoryArticleAssignmentConfig  `json:"article_assignment"`
            Payment AffiliateInventoryPaymentConfig  `json:"payment"`
            Inventory AffiliateInventoryStoreConfig  `json:"inventory"`
            Order AffiliateInventoryOrderConfig  `json:"order"`
            Logistics AffiliateInventoryLogisticsConfig  `json:"logistics"`
         
    }
    
    // AffiliateConfig used by Order
    type AffiliateConfig struct {

        
            App AffiliateAppConfig  `json:"app"`
            Inventory AffiliateInventoryConfig  `json:"inventory"`
         
    }
    
    // Affiliate used by Order
    type Affiliate struct {

        
            ID string  `json:"id"`
            Token string  `json:"token"`
            Config AffiliateConfig  `json:"config"`
         
    }
    
    // OrderConfig used by Order
    type OrderConfig struct {

        
            BagEndState string  `json:"bag_end_state"`
            CreateUser bool  `json:"create_user"`
            AffiliateStoreIDMapping []AffiliateStoreIdMapping  `json:"affiliate_store_id_mapping"`
            ArticleLookup string  `json:"article_lookup"`
            Affiliate Affiliate  `json:"affiliate"`
            StoreLookup string  `json:"store_lookup"`
         
    }
    
    // CreateOrderPayload used by Order
    type CreateOrderPayload struct {

        
            OrderInfo OrderInfo  `json:"order_info"`
            AffiliateID string  `json:"affiliate_id"`
            OrderConfig OrderConfig  `json:"order_config"`
         
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

        
            ID float64  `json:"id"`
            DisplayText string  `json:"display_text"`
            Description string  `json:"description"`
            Slug string  `json:"slug"`
         
    }
    
    // GetActionsResponse used by Order
    type GetActionsResponse struct {

        
            Permissions []ActionInfo  `json:"permissions"`
         
    }
    
    // HistoryReason used by Order
    type HistoryReason struct {

        
            Text string  `json:"text"`
            Code float64  `json:"code"`
            State string  `json:"state"`
            DislayName string  `json:"dislay_name"`
            Category string  `json:"category"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // HistoryMeta used by Order
    type HistoryMeta struct {

        
            ActivityType string  `json:"activity_type"`
            Slug string  `json:"slug"`
            Duration string  `json:"duration"`
            Status1 string  `json:"status1"`
            StoreName string  `json:"store_name"`
            Receiver string  `json:"receiver"`
            StoreID float64  `json:"store_id"`
            ActivityComment string  `json:"activity_comment"`
            CallID string  `json:"call_id"`
            Starttime string  `json:"starttime"`
            Message string  `json:"message"`
            Callerid string  `json:"callerid"`
            StoreCode string  `json:"store_code"`
            Status2 string  `json:"status2"`
            ShortLink string  `json:"short_link"`
            Billsec string  `json:"billsec"`
            Caller string  `json:"caller"`
            Recipient string  `json:"recipient"`
            Endtime string  `json:"endtime"`
            Reason HistoryReason  `json:"reason"`
            ChannelType string  `json:"channel_type"`
            Status string  `json:"status"`
            Recordpath string  `json:"recordpath"`
         
    }
    
    // HistoryDict used by Order
    type HistoryDict struct {

        
            Meta HistoryMeta  `json:"meta"`
            L2Detail string  `json:"l2_detail"`
            Type string  `json:"type"`
            Message string  `json:"message"`
            L1Detail string  `json:"l1_detail"`
            BagID float64  `json:"bag_id"`
            User string  `json:"user"`
            TicketURL string  `json:"ticket_url"`
            AssignedAgent string  `json:"assigned_agent"`
            DisplayMessage string  `json:"display_message"`
            Createdat string  `json:"createdat"`
            L3Detail string  `json:"l3_detail"`
            TicketID string  `json:"ticket_id"`
         
    }
    
    // ShipmentHistoryResponse used by Order
    type ShipmentHistoryResponse struct {

        
            Success bool  `json:"success"`
            ActivityHistory []HistoryDict  `json:"activity_history"`
         
    }
    
    // PostHistoryData used by Order
    type PostHistoryData struct {

        
            UserName string  `json:"user_name"`
            Message string  `json:"message"`
         
    }
    
    // PostHistoryFilters used by Order
    type PostHistoryFilters struct {

        
            ShipmentID string  `json:"shipment_id"`
            LineNumber string  `json:"line_number"`
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
    
    // SmsDataPayload used by Order
    type SmsDataPayload struct {

        
            ShipmentID float64  `json:"shipment_id"`
            Message string  `json:"message"`
            PaymentMode string  `json:"payment_mode"`
            CustomerName string  `json:"customer_name"`
            BrandName string  `json:"brand_name"`
            AmountPaid float64  `json:"amount_paid"`
            OrderID string  `json:"order_id"`
            PhoneNumber float64  `json:"phone_number"`
            CountryCode string  `json:"country_code"`
         
    }
    
    // SendSmsPayload used by Order
    type SendSmsPayload struct {

        
            Slug string  `json:"slug"`
            Data SmsDataPayload  `json:"data"`
            BagID float64  `json:"bag_id"`
         
    }
    
    // Meta used by Order
    type Meta struct {

        
            KafkaEmissionStatus float64  `json:"kafka_emission_status"`
            StateManagerUsed string  `json:"state_manager_used"`
         
    }
    
    // ShipmentDetail used by Order
    type ShipmentDetail struct {

        
            ShipmentID string  `json:"shipment_id"`
            ID float64  `json:"id"`
            BagList []float64  `json:"bag_list"`
            Status string  `json:"status"`
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
            Errors []string  `json:"errors"`
            OrderDetails OrderDetails  `json:"order_details"`
         
    }
    
    // OrderStatusResult used by Order
    type OrderStatusResult struct {

        
            Success string  `json:"success"`
            Result []OrderStatusData  `json:"result"`
         
    }
    
    // Dimension used by Order
    type Dimension struct {

        
            Height string  `json:"height"`
            Length float64  `json:"length"`
            Weight string  `json:"weight"`
            Width float64  `json:"width"`
            PackagingType string  `json:"packaging_type"`
         
    }
    
    // UpdatePackagingDimensionsPayload used by Order
    type UpdatePackagingDimensionsPayload struct {

        
            ShipmentID string  `json:"shipment_id"`
            CurrentStatus string  `json:"current_status"`
            Dimension []Dimension  `json:"dimension"`
         
    }
    
    // UpdatePackagingDimensionsResponse used by Order
    type UpdatePackagingDimensionsResponse struct {

        
            Message string  `json:"message"`
         
    }
    
    // ShippingInfo used by Order
    type ShippingInfo struct {

        
            Title string  `json:"title"`
            Pincode string  `json:"pincode"`
            MiddleName string  `json:"middle_name"`
            ExternalCustomerCode string  `json:"external_customer_code"`
            PrimaryMobileNumber string  `json:"primary_mobile_number"`
            FloorNo string  `json:"floor_no"`
            PrimaryEmail string  `json:"primary_email"`
            AlternateMobileNumber string  `json:"alternate_mobile_number"`
            Country string  `json:"country"`
            StateCode string  `json:"state_code"`
            Address1 string  `json:"address1"`
            Slot []map[string]interface{}  `json:"slot"`
            LastName string  `json:"last_name"`
            Address2 string  `json:"address2"`
            State string  `json:"state"`
            HouseNo string  `json:"house_no"`
            City string  `json:"city"`
            AddressType string  `json:"address_type"`
            FirstName string  `json:"first_name"`
            CustomerCode string  `json:"customer_code"`
            CountryCode string  `json:"country_code"`
            AlternateEmail string  `json:"alternate_email"`
            GeoLocation map[string]interface{}  `json:"geo_location"`
            Landmark string  `json:"landmark"`
            Gender string  `json:"gender"`
            ShippingType string  `json:"shipping_type"`
         
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
            Amount map[string]interface{}  `json:"amount"`
            Code string  `json:"code"`
            Name string  `json:"name"`
            Tax Tax  `json:"tax"`
         
    }
    
    // BillingInfo used by Order
    type BillingInfo struct {

        
            Title string  `json:"title"`
            Pincode string  `json:"pincode"`
            MiddleName string  `json:"middle_name"`
            ExternalCustomerCode string  `json:"external_customer_code"`
            PrimaryMobileNumber string  `json:"primary_mobile_number"`
            FloorNo string  `json:"floor_no"`
            PrimaryEmail string  `json:"primary_email"`
            AlternateMobileNumber string  `json:"alternate_mobile_number"`
            Country string  `json:"country"`
            StateCode string  `json:"state_code"`
            Address1 string  `json:"address1"`
            LastName string  `json:"last_name"`
            Address2 string  `json:"address2"`
            State string  `json:"state"`
            HouseNo string  `json:"house_no"`
            City string  `json:"city"`
            FirstName string  `json:"first_name"`
            CustomerCode string  `json:"customer_code"`
            CountryCode string  `json:"country_code"`
            AlternateEmail string  `json:"alternate_email"`
            Gender string  `json:"gender"`
         
    }
    
    // TaxInfo used by Order
    type TaxInfo struct {

        
            Gstin string  `json:"gstin"`
            B2bGstinNumber string  `json:"b2b_gstin_number"`
         
    }
    
    // ProcessingDates used by Order
    type ProcessingDates struct {

        
            CustomerPickupSlot map[string]interface{}  `json:"customer_pickup_slot"`
            PackByDate string  `json:"pack_by_date"`
            ConfirmByDate string  `json:"confirm_by_date"`
            DispatchByDate string  `json:"dispatch_by_date"`
            DispatchAfterDate string  `json:"dispatch_after_date"`
            DpPickupSlot map[string]interface{}  `json:"dp_pickup_slot"`
         
    }
    
    // LineItem used by Order
    type LineItem struct {

        
            Meta map[string]interface{}  `json:"meta"`
            Charges []Charge  `json:"charges"`
            ExternalLineID string  `json:"external_line_id"`
            SellerIdentifier string  `json:"seller_identifier"`
            Quantity float64  `json:"quantity"`
            CustomMessasge string  `json:"custom_messasge"`
         
    }
    
    // Shipment used by Order
    type Shipment struct {

        
            ProcessingDates ProcessingDates  `json:"processing_dates"`
            Priority float64  `json:"priority"`
            ExternalShipmentID string  `json:"external_shipment_id"`
            LocationID float64  `json:"location_id"`
            LineItems []LineItem  `json:"line_items"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // PaymentMethod used by Order
    type PaymentMethod struct {

        
            TransactionData map[string]interface{}  `json:"transaction_data"`
            RefundBy string  `json:"refund_by"`
            Amount float64  `json:"amount"`
            CollectBy string  `json:"collect_by"`
            Mode string  `json:"mode"`
            Name string  `json:"name"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // PaymentInfo used by Order
    type PaymentInfo struct {

        
            PrimaryMode string  `json:"primary_mode"`
            PaymentMethods []PaymentMethod  `json:"payment_methods"`
         
    }
    
    // CreateOrderAPI used by Order
    type CreateOrderAPI struct {

        
            ShippingInfo ShippingInfo  `json:"shipping_info"`
            ExternalCreationDate string  `json:"external_creation_date"`
            Charges []Charge  `json:"charges"`
            CurrencyInfo map[string]interface{}  `json:"currency_info"`
            BillingInfo BillingInfo  `json:"billing_info"`
            TaxInfo TaxInfo  `json:"tax_info"`
            ExternalOrderID string  `json:"external_order_id"`
            Shipments []Shipment  `json:"shipments"`
            Config map[string]interface{}  `json:"config"`
            PaymentInfo PaymentInfo  `json:"payment_info"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // CreateOrderErrorReponse used by Order
    type CreateOrderErrorReponse struct {

        
            Message string  `json:"message"`
            Info interface{}  `json:"info"`
            RequestID string  `json:"request_id"`
            Code string  `json:"code"`
            StackTrace string  `json:"stack_trace"`
            Exception string  `json:"exception"`
            Status float64  `json:"status"`
            Meta string  `json:"meta"`
         
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

        
            PaymentMethods []PaymentMethods  `json:"payment_methods"`
            Source string  `json:"source"`
            ModeOfPayment string  `json:"mode_of_payment"`
         
    }
    
    // CreateChannelConfig used by Order
    type CreateChannelConfig struct {

        
            ShipmentAssignment string  `json:"shipment_assignment"`
            DpConfiguration DpConfiguration  `json:"dp_configuration"`
            LockStates []string  `json:"lock_states"`
            LogoURL map[string]interface{}  `json:"logo_url"`
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

        
            IsInserted bool  `json:"is_inserted"`
            Acknowledged bool  `json:"acknowledged"`
            IsUpserted bool  `json:"is_upserted"`
         
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
            StartDate string  `json:"start_date"`
            OrderDetails []FyndOrderIdList  `json:"order_details"`
         
    }
    
    // BagStateTransitionMap used by Order
    type BagStateTransitionMap struct {

        
            Affiliate map[string]interface{}  `json:"affiliate"`
            Fynd map[string]interface{}  `json:"fynd"`
         
    }
    
    // FetchCreditBalanceRequestPayload used by Order
    type FetchCreditBalanceRequestPayload struct {

        
            CustomerMobileNumber string  `json:"customer_mobile_number"`
            SellerID string  `json:"seller_id"`
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
            FyndOrderID string  `json:"fynd_order_id"`
            SellerID float64  `json:"seller_id"`
            AffiliateID string  `json:"affiliate_id"`
         
    }
    
    // RefundOption used by Order
    type RefundOption struct {

        
            Slug string  `json:"slug"`
            Name string  `json:"name"`
            Value string  `json:"value"`
            Type string  `json:"type"`
         
    }
    
    // RefundModeInfo used by Order
    type RefundModeInfo struct {

        
            IsActive bool  `json:"is_active"`
            Options []RefundOption  `json:"options"`
            DisplayName string  `json:"display_name"`
            Slug string  `json:"slug"`
         
    }
    
    // RefundModeConfigResponsePayload used by Order
    type RefundModeConfigResponsePayload struct {

        
            Data []RefundModeInfo  `json:"data"`
            Success bool  `json:"success"`
         
    }
    
    // AttachUserOtpData used by Order
    type AttachUserOtpData struct {

        
            RequestID string  `json:"request_id"`
         
    }
    
    // AttachUserInfo used by Order
    type AttachUserInfo struct {

        
            LastName string  `json:"last_name"`
            CountryCode string  `json:"country_code"`
            Mobile string  `json:"mobile"`
            FirstName string  `json:"first_name"`
         
    }
    
    // AttachOrderUser used by Order
    type AttachOrderUser struct {

        
            OtpData AttachUserOtpData  `json:"otp_data"`
            FyndOrderID string  `json:"fynd_order_id"`
            UserInfo AttachUserInfo  `json:"user_info"`
         
    }
    
    // AttachOrderUserResponse used by Order
    type AttachOrderUserResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // SendUserMobileOTP used by Order
    type SendUserMobileOTP struct {

        
            CountryCode string  `json:"country_code"`
            Mobile string  `json:"mobile"`
         
    }
    
    // PointBlankOtpData used by Order
    type PointBlankOtpData struct {

        
            ResendTimer float64  `json:"resend_timer"`
            Mobile float64  `json:"mobile"`
            RequestID string  `json:"request_id"`
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

        
            OtpCode float64  `json:"otp_code"`
            Mobile string  `json:"mobile"`
            RequestID string  `json:"request_id"`
         
    }
    
    // VerifyMobileOTP used by Order
    type VerifyMobileOTP struct {

        
            OtpData VerifyOtpData  `json:"otp_data"`
            FyndOrderID string  `json:"fynd_order_id"`
         
    }
    
    // VerifyOtpResponseData used by Order
    type VerifyOtpResponseData struct {

        
            FyndOrderID string  `json:"fynd_order_id"`
            CountryCode string  `json:"country_code"`
            Mobile string  `json:"mobile"`
            Message string  `json:"message"`
         
    }
    
    // VerifyOtpResponse used by Order
    type VerifyOtpResponse struct {

        
            Data VerifyOtpResponseData  `json:"data"`
            Success bool  `json:"success"`
            Status float64  `json:"status"`
            Message string  `json:"message"`
         
    }
    

    
    // GetSearchWordsData used by Catalog
    type GetSearchWordsData struct {

        
            AppID string  `json:"app_id"`
            UID string  `json:"uid"`
            Result map[string]interface{}  `json:"result"`
            Words []string  `json:"words"`
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

        
            Code string  `json:"code"`
            Message string  `json:"message"`
            Error string  `json:"error"`
            Meta map[string]interface{}  `json:"meta"`
            Status float64  `json:"status"`
         
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
            Words []string  `json:"words"`
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

        
            AppID string  `json:"app_id"`
            UID string  `json:"uid"`
            Words []string  `json:"words"`
            Results []map[string]interface{}  `json:"results"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // GetAutocompleteWordsResponse used by Catalog
    type GetAutocompleteWordsResponse struct {

        
            Items []GetAutocompleteWordsData  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // AutoCompleteMedia used by Catalog
    type AutoCompleteMedia struct {

        
            AspectRatio string  `json:"aspect_ratio"`
            URL string  `json:"url"`
            Type string  `json:"type"`
         
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

        
            Type string  `json:"type"`
            Page AutocompletePageAction  `json:"page"`
         
    }
    
    // AutocompleteResult used by Catalog
    type AutocompleteResult struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Logo AutoCompleteMedia  `json:"logo"`
            Display string  `json:"display"`
            Action AutocompleteAction  `json:"action"`
         
    }
    
    // CreateAutocompleteKeyword used by Catalog
    type CreateAutocompleteKeyword struct {

        
            AppID string  `json:"app_id"`
            Words []string  `json:"words"`
            IsActive bool  `json:"is_active"`
            Results []AutocompleteResult  `json:"results"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // CreateAutocompleteWordsResponse used by Catalog
    type CreateAutocompleteWordsResponse struct {

        
            AppID string  `json:"app_id"`
            Words []string  `json:"words"`
            Results []map[string]interface{}  `json:"results"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // ProductBundleItem used by Catalog
    type ProductBundleItem struct {

        
            AllowRemove bool  `json:"allow_remove"`
            MinQuantity float64  `json:"min_quantity"`
            AutoSelect bool  `json:"auto_select"`
            AutoAddToCart bool  `json:"auto_add_to_cart"`
            ProductUID float64  `json:"product_uid"`
            MaxQuantity float64  `json:"max_quantity"`
         
    }
    
    // GetProductBundleCreateResponse used by Catalog
    type GetProductBundleCreateResponse struct {

        
            SameStoreAssignment bool  `json:"same_store_assignment"`
            ModifiedOn string  `json:"modified_on"`
            Products []ProductBundleItem  `json:"products"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Choice string  `json:"choice"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            ID string  `json:"id"`
            Logo string  `json:"logo"`
            IsActive bool  `json:"is_active"`
            PageVisibility []string  `json:"page_visibility"`
            CreatedOn string  `json:"created_on"`
            CompanyID float64  `json:"company_id"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // GetProductBundleListingResponse used by Catalog
    type GetProductBundleListingResponse struct {

        
            Items []GetProductBundleCreateResponse  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // ProductBundleRequest used by Catalog
    type ProductBundleRequest struct {

        
            SameStoreAssignment bool  `json:"same_store_assignment"`
            ModifiedOn string  `json:"modified_on"`
            Products []ProductBundleItem  `json:"products"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Choice string  `json:"choice"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Logo string  `json:"logo"`
            IsActive bool  `json:"is_active"`
            PageVisibility []string  `json:"page_visibility"`
            CreatedOn string  `json:"created_on"`
            CompanyID float64  `json:"company_id"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // Price used by Catalog
    type Price struct {

        
            Currency string  `json:"currency"`
            MinEffective float64  `json:"min_effective"`
            MaxMarked float64  `json:"max_marked"`
            MinMarked float64  `json:"min_marked"`
            MaxEffective float64  `json:"max_effective"`
         
    }
    
    // LimitedProductData used by Catalog
    type LimitedProductData struct {

        
            UID float64  `json:"uid"`
            Attributes map[string]interface{}  `json:"attributes"`
            ItemCode string  `json:"item_code"`
            Quantity float64  `json:"quantity"`
            Identifier map[string]interface{}  `json:"identifier"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            CountryOfOrigin string  `json:"country_of_origin"`
            Price map[string]interface{}  `json:"price"`
            ShortDescription string  `json:"short_description"`
            Images []string  `json:"images"`
            Sizes []string  `json:"sizes"`
         
    }
    
    // Size used by Catalog
    type Size struct {

        
            IsAvailable bool  `json:"is_available"`
            Quantity float64  `json:"quantity"`
            Value string  `json:"value"`
            Display string  `json:"display"`
         
    }
    
    // GetProducts used by Catalog
    type GetProducts struct {

        
            AllowRemove bool  `json:"allow_remove"`
            MinQuantity float64  `json:"min_quantity"`
            AutoSelect bool  `json:"auto_select"`
            AutoAddToCart bool  `json:"auto_add_to_cart"`
            Price Price  `json:"price"`
            ProductUID float64  `json:"product_uid"`
            MaxQuantity float64  `json:"max_quantity"`
            ProductDetails LimitedProductData  `json:"product_details"`
            Sizes []Size  `json:"sizes"`
         
    }
    
    // GetProductBundleResponse used by Catalog
    type GetProductBundleResponse struct {

        
            SameStoreAssignment bool  `json:"same_store_assignment"`
            Products []GetProducts  `json:"products"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Choice string  `json:"choice"`
            Logo string  `json:"logo"`
            PageVisibility []string  `json:"page_visibility"`
            IsActive bool  `json:"is_active"`
            CompanyID float64  `json:"company_id"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // ProductBundleUpdateRequest used by Catalog
    type ProductBundleUpdateRequest struct {

        
            SameStoreAssignment bool  `json:"same_store_assignment"`
            ModifiedOn string  `json:"modified_on"`
            Products []ProductBundleItem  `json:"products"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Choice string  `json:"choice"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Logo string  `json:"logo"`
            IsActive bool  `json:"is_active"`
            PageVisibility []string  `json:"page_visibility"`
            CompanyID float64  `json:"company_id"`
            Meta map[string]interface{}  `json:"meta"`
         
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

        
            Image string  `json:"image"`
            Subtitle string  `json:"subtitle"`
            Active bool  `json:"active"`
            Tag string  `json:"tag"`
            ModifiedOn string  `json:"modified_on"`
            Name string  `json:"name"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ID string  `json:"id"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            BrandID float64  `json:"brand_id"`
            Guide Guide  `json:"guide"`
            CreatedOn string  `json:"created_on"`
            CompanyID float64  `json:"company_id"`
            Description string  `json:"description"`
            Title string  `json:"title"`
         
    }
    
    // SuccessResponse used by Catalog
    type SuccessResponse struct {

        
            Success bool  `json:"success"`
         
    }
    
    // SizeGuideResponse used by Catalog
    type SizeGuideResponse struct {

        
            Active bool  `json:"active"`
            ModifiedOn string  `json:"modified_on"`
            Tag string  `json:"tag"`
            Subtitle string  `json:"subtitle"`
            Name string  `json:"name"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ID string  `json:"id"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            BrandID float64  `json:"brand_id"`
            Guide map[string]interface{}  `json:"guide"`
            CreatedOn string  `json:"created_on"`
            CompanyID float64  `json:"company_id"`
            Title string  `json:"title"`
         
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

        
            IsGift bool  `json:"is_gift"`
            IsCod bool  `json:"is_cod"`
            Moq MOQData  `json:"moq"`
            AltText map[string]interface{}  `json:"alt_text"`
            Seo SEOData  `json:"seo"`
         
    }
    
    // ApplicationItemMOQ used by Catalog
    type ApplicationItemMOQ struct {

        
            Maximum float64  `json:"maximum"`
            Minimum float64  `json:"minimum"`
            IncrementUnit float64  `json:"increment_unit"`
         
    }
    
    // ApplicationItemSEO used by Catalog
    type ApplicationItemSEO struct {

        
            Description interface{}  `json:"description"`
            Title interface{}  `json:"title"`
         
    }
    
    // MetaFields used by Catalog
    type MetaFields struct {

        
            Value interface{}  `json:"value"`
            Key interface{}  `json:"key"`
         
    }
    
    // ApplicationItemMeta used by Catalog
    type ApplicationItemMeta struct {

        
            IsGift bool  `json:"is_gift"`
            IsCod bool  `json:"is_cod"`
            Moq ApplicationItemMOQ  `json:"moq"`
            AltText map[string]interface{}  `json:"alt_text"`
            Seo ApplicationItemSEO  `json:"seo"`
            CustomMeta []MetaFields  `json:"_custom_meta"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // SuccessResponse1 used by Catalog
    type SuccessResponse1 struct {

        
            UID float64  `json:"uid"`
            Success bool  `json:"success"`
         
    }
    
    // GetConfigMetadataResponse used by Catalog
    type GetConfigMetadataResponse struct {

        
            Values []map[string]interface{}  `json:"values"`
            Data []map[string]interface{}  `json:"data"`
            Condition []map[string]interface{}  `json:"condition"`
         
    }
    
    // PageResponseType used by Catalog
    type PageResponseType struct {

        
            Current float64  `json:"current"`
            TotalCount float64  `json:"total_count"`
            HasNext bool  `json:"has_next"`
            Next float64  `json:"next"`
         
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

        
            Unit string  `json:"unit"`
            Key string  `json:"key"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Priority float64  `json:"priority"`
            Logo string  `json:"logo"`
            DisplayType string  `json:"display_type"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // AppConfigurationDetail used by Catalog
    type AppConfigurationDetail struct {

        
            AppID string  `json:"app_id"`
            Attributes []AttributeDetailsGroup  `json:"attributes"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Priority float64  `json:"priority"`
            Logo string  `json:"logo"`
            TemplateSlugs []string  `json:"template_slugs"`
            IsActive bool  `json:"is_active"`
            IsDefault bool  `json:"is_default"`
         
    }
    
    // ConfigSuccessResponse used by Catalog
    type ConfigSuccessResponse struct {

        
            Message string  `json:"message"`
         
    }
    
    // AppConfigurationsSort used by Catalog
    type AppConfigurationsSort struct {

        
            AppID string  `json:"app_id"`
            DefaultKey string  `json:"default_key"`
            Key string  `json:"key"`
            Name string  `json:"name"`
            Priority float64  `json:"priority"`
            Logo string  `json:"logo"`
            IsActive bool  `json:"is_active"`
            IsDefault bool  `json:"is_default"`
         
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
            Detail map[string]interface{}  `json:"detail"`
            Compare map[string]interface{}  `json:"compare"`
            Similar map[string]interface{}  `json:"similar"`
         
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
            Units []map[string]interface{}  `json:"units"`
            Key string  `json:"key"`
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
    
    // ConfigurationProductVariantConfig used by Catalog
    type ConfigurationProductVariantConfig struct {

        
            Size ProductSize  `json:"size"`
            Key string  `json:"key"`
            Name string  `json:"name"`
            Priority float64  `json:"priority"`
            Logo string  `json:"logo"`
            DisplayType string  `json:"display_type"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // ConfigurationProductVariant used by Catalog
    type ConfigurationProductVariant struct {

        
            Config []ConfigurationProductVariantConfig  `json:"config"`
         
    }
    
    // ConfigurationProductConfig used by Catalog
    type ConfigurationProductConfig struct {

        
            Size ProductSize  `json:"size"`
            Subtitle string  `json:"subtitle"`
            Key string  `json:"key"`
            Priority float64  `json:"priority"`
            Logo string  `json:"logo"`
            IsActive bool  `json:"is_active"`
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
    
    // ConfigurationListingSortConfig used by Catalog
    type ConfigurationListingSortConfig struct {

        
            Key string  `json:"key"`
            Name string  `json:"name"`
            Priority float64  `json:"priority"`
            Logo string  `json:"logo"`
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

        
            MapValues []map[string]interface{}  `json:"map_values"`
            Map map[string]interface{}  `json:"map"`
            Sort string  `json:"sort"`
            Value string  `json:"value"`
            Condition string  `json:"condition"`
            BucketPoints []ConfigurationBucketPoints  `json:"bucket_points"`
         
    }
    
    // ConfigurationListingFilterConfig used by Catalog
    type ConfigurationListingFilterConfig struct {

        
            DisplayName string  `json:"display_name"`
            Key string  `json:"key"`
            Name string  `json:"name"`
            Priority float64  `json:"priority"`
            Logo string  `json:"logo"`
            Type string  `json:"type"`
            ValueConfig ConfigurationListingFilterValue  `json:"value_config"`
            IsActive bool  `json:"is_active"`
         
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
    
    // AppCatalogConfiguration used by Catalog
    type AppCatalogConfiguration struct {

        
            AppID string  `json:"app_id"`
            ConfigID string  `json:"config_id"`
            ModifiedOn string  `json:"modified_on"`
            Product ConfigurationProduct  `json:"product"`
            ID string  `json:"id"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Listing ConfigurationListing  `json:"listing"`
            Type string  `json:"type"`
            CreatedOn string  `json:"created_on"`
            ConfigType string  `json:"config_type"`
         
    }
    
    // GetAppCatalogConfiguration used by Catalog
    type GetAppCatalogConfiguration struct {

        
            Data AppCatalogConfiguration  `json:"data"`
            IsDefault bool  `json:"is_default"`
         
    }
    
    // AppConfiguration used by Catalog
    type AppConfiguration struct {

        
            AppID string  `json:"app_id"`
            ConfigID string  `json:"config_id"`
            ModifiedOn string  `json:"modified_on"`
            Product ConfigurationProduct  `json:"product"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Listing ConfigurationListing  `json:"listing"`
            Type string  `json:"type"`
            CreatedOn string  `json:"created_on"`
            ConfigType string  `json:"config_type"`
         
    }
    
    // GetCatalogConfigurationDetailsSchemaListing used by Catalog
    type GetCatalogConfigurationDetailsSchemaListing struct {

        
            Sort map[string]interface{}  `json:"sort"`
            Filter map[string]interface{}  `json:"filter"`
         
    }
    
    // EntityConfiguration used by Catalog
    type EntityConfiguration struct {

        
            AppID string  `json:"app_id"`
            ConfigID string  `json:"config_id"`
            Product GetCatalogConfigurationDetailsProduct  `json:"product"`
            ID string  `json:"id"`
            Listing GetCatalogConfigurationDetailsSchemaListing  `json:"listing"`
            ConfigType string  `json:"config_type"`
         
    }
    
    // GetAppCatalogEntityConfiguration used by Catalog
    type GetAppCatalogEntityConfiguration struct {

        
            Data EntityConfiguration  `json:"data"`
            IsDefault bool  `json:"is_default"`
         
    }
    
    // ProductFiltersValue used by Catalog
    type ProductFiltersValue struct {

        
            CurrencyCode string  `json:"currency_code"`
            Min float64  `json:"min"`
            SelectedMin float64  `json:"selected_min"`
            SelectedMax float64  `json:"selected_max"`
            Display string  `json:"display"`
            DisplayFormat string  `json:"display_format"`
            IsSelected bool  `json:"is_selected"`
            QueryFormat string  `json:"query_format"`
            Value interface{}  `json:"value"`
            Max float64  `json:"max"`
            Count float64  `json:"count"`
            CurrencySymbol string  `json:"currency_symbol"`
         
    }
    
    // ProductFiltersKey used by Catalog
    type ProductFiltersKey struct {

        
            Kind string  `json:"kind"`
            Display string  `json:"display"`
            Name string  `json:"name"`
            Logo string  `json:"logo"`
            Operators []string  `json:"operators"`
         
    }
    
    // ProductFilters used by Catalog
    type ProductFilters struct {

        
            Values []ProductFiltersValue  `json:"values"`
            Key ProductFiltersKey  `json:"key"`
         
    }
    
    // ProductSortOn used by Catalog
    type ProductSortOn struct {

        
            IsSelected bool  `json:"is_selected"`
            Value string  `json:"value"`
            Name string  `json:"name"`
         
    }
    
    // GetCollectionQueryOptionResponse used by Catalog
    type GetCollectionQueryOptionResponse struct {

        
            Filters []ProductFilters  `json:"filters"`
            Operators map[string]string  `json:"operators"`
            SortOn []ProductSortOn  `json:"sort_on"`
         
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
    
    // CollectionQuery used by Catalog
    type CollectionQuery struct {

        
            Attribute string  `json:"attribute"`
            Op string  `json:"op"`
            Value []interface{}  `json:"value"`
         
    }
    
    // Media used by Catalog
    type Media struct {

        
            URL string  `json:"url"`
            Meta map[string]interface{}  `json:"meta"`
            Type string  `json:"type"`
         
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
    
    // GetCollectionDetailNest used by Catalog
    type GetCollectionDetailNest struct {

        
            UID string  `json:"uid"`
            Cron map[string]interface{}  `json:"cron"`
            Slug string  `json:"slug"`
            Action Action  `json:"action"`
            Meta map[string]interface{}  `json:"meta"`
            Description string  `json:"description"`
            Badge map[string]interface{}  `json:"badge"`
            AppID string  `json:"app_id"`
            Tag []string  `json:"tag"`
            Query []CollectionQuery  `json:"query"`
            AllowSort bool  `json:"allow_sort"`
            Schedule map[string]interface{}  `json:"_schedule"`
            AllowFacets bool  `json:"allow_facets"`
            Name string  `json:"name"`
            Priority float64  `json:"priority"`
            Logo Media  `json:"logo"`
            Banners ImageUrls  `json:"banners"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Type string  `json:"type"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // GetCollectionListingResponse used by Catalog
    type GetCollectionListingResponse struct {

        
            Filters CollectionListingFilter  `json:"filters"`
            Items []GetCollectionDetailNest  `json:"items"`
            Page Page  `json:"page"`
         
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

        
            NextSchedule []NextSchedule  `json:"next_schedule"`
            Duration float64  `json:"duration"`
            Cron string  `json:"cron"`
            End string  `json:"end"`
            Start string  `json:"start"`
         
    }
    
    // UserInfo used by Catalog
    type UserInfo struct {

        
            Email string  `json:"email"`
            UID string  `json:"uid"`
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
         
    }
    
    // CollectionImage used by Catalog
    type CollectionImage struct {

        
            AspectRatio string  `json:"aspect_ratio"`
            URL string  `json:"url"`
         
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
    
    // CreateCollection used by Catalog
    type CreateCollection struct {

        
            Slug string  `json:"slug"`
            IsVisible bool  `json:"is_visible"`
            Meta map[string]interface{}  `json:"meta"`
            Description string  `json:"description"`
            Badge CollectionBadge  `json:"badge"`
            AppID string  `json:"app_id"`
            Query []CollectionQuery  `json:"query"`
            AllowSort bool  `json:"allow_sort"`
            Schedule CollectionSchedule  `json:"_schedule"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            AllowFacets bool  `json:"allow_facets"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Tags []string  `json:"tags"`
            SortOn string  `json:"sort_on"`
            Name string  `json:"name"`
            ModifiedBy UserInfo  `json:"modified_by"`
            Priority float64  `json:"priority"`
            Logo CollectionImage  `json:"logo"`
            Banners CollectionBanner  `json:"banners"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            CreatedBy UserInfo  `json:"created_by"`
            Published bool  `json:"published"`
            Type string  `json:"type"`
            Seo SeoDetail  `json:"seo"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // CollectionCreateResponse used by Catalog
    type CollectionCreateResponse struct {

        
            Cron map[string]interface{}  `json:"cron"`
            Slug string  `json:"slug"`
            Meta map[string]interface{}  `json:"meta"`
            Description string  `json:"description"`
            Badge map[string]interface{}  `json:"badge"`
            AppID string  `json:"app_id"`
            Tag []string  `json:"tag"`
            Query []CollectionQuery  `json:"query"`
            AllowSort bool  `json:"allow_sort"`
            Schedule map[string]interface{}  `json:"_schedule"`
            AllowFacets bool  `json:"allow_facets"`
            SortOn string  `json:"sort_on"`
            Name string  `json:"name"`
            Priority float64  `json:"priority"`
            Logo BannerImage  `json:"logo"`
            Banners ImageUrls  `json:"banners"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Type string  `json:"type"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // CollectionDetailResponse used by Catalog
    type CollectionDetailResponse struct {

        
            UID string  `json:"uid"`
            Cron map[string]interface{}  `json:"cron"`
            Slug string  `json:"slug"`
            Meta map[string]interface{}  `json:"meta"`
            Description string  `json:"description"`
            Badge map[string]interface{}  `json:"badge"`
            AppID string  `json:"app_id"`
            Tag []string  `json:"tag"`
            Query []CollectionQuery  `json:"query"`
            AllowSort bool  `json:"allow_sort"`
            Schedule map[string]interface{}  `json:"_schedule"`
            AllowFacets bool  `json:"allow_facets"`
            Name string  `json:"name"`
            Priority float64  `json:"priority"`
            Logo Media  `json:"logo"`
            Banners ImageUrls  `json:"banners"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Type string  `json:"type"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // UpdateCollection used by Catalog
    type UpdateCollection struct {

        
            Slug string  `json:"slug"`
            IsVisible bool  `json:"is_visible"`
            Meta map[string]interface{}  `json:"meta"`
            Description string  `json:"description"`
            Badge CollectionBadge  `json:"badge"`
            Query []CollectionQuery  `json:"query"`
            Schedule CollectionSchedule  `json:"_schedule"`
            AllowSort bool  `json:"allow_sort"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            AllowFacets bool  `json:"allow_facets"`
            Tags []string  `json:"tags"`
            SortOn string  `json:"sort_on"`
            Name string  `json:"name"`
            ModifiedBy UserInfo  `json:"modified_by"`
            Priority float64  `json:"priority"`
            Logo CollectionImage  `json:"logo"`
            Banners CollectionBanner  `json:"banners"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Published bool  `json:"published"`
            Type string  `json:"type"`
            Seo SeoDetail  `json:"seo"`
            IsActive bool  `json:"is_active"`
         
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
    
    // ProductBrand used by Catalog
    type ProductBrand struct {

        
            UID float64  `json:"uid"`
            Action Action  `json:"action"`
            Logo Media  `json:"logo"`
            Name string  `json:"name"`
         
    }
    
    // Price1 used by Catalog
    type Price1 struct {

        
            CurrencyCode string  `json:"currency_code"`
            Min float64  `json:"min"`
            Max float64  `json:"max"`
            CurrencySymbol string  `json:"currency_symbol"`
         
    }
    
    // ProductListingPrice used by Catalog
    type ProductListingPrice struct {

        
            Effective Price1  `json:"effective"`
            Marked Price1  `json:"marked"`
         
    }
    
    // ProductListingDetail used by Catalog
    type ProductListingDetail struct {

        
            UID float64  `json:"uid"`
            HasVariant bool  `json:"has_variant"`
            ItemCode string  `json:"item_code"`
            Medias []Media  `json:"medias"`
            Slug string  `json:"slug"`
            RatingCount float64  `json:"rating_count"`
            PromoMeta map[string]interface{}  `json:"promo_meta"`
            ShortDescription string  `json:"short_description"`
            Description string  `json:"description"`
            Rating float64  `json:"rating"`
            ProductOnlineDate string  `json:"product_online_date"`
            Similars []string  `json:"similars"`
            Discount string  `json:"discount"`
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
            Attributes map[string]interface{}  `json:"attributes"`
            Name string  `json:"name"`
            Brand ProductBrand  `json:"brand"`
            ItemType string  `json:"item_type"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            Sellable bool  `json:"sellable"`
            ImageNature string  `json:"image_nature"`
            Price ProductListingPrice  `json:"price"`
            Highlights []string  `json:"highlights"`
            Type string  `json:"type"`
            Tryouts []string  `json:"tryouts"`
            Color string  `json:"color"`
         
    }
    
    // GetCollectionItemsResponse used by Catalog
    type GetCollectionItemsResponse struct {

        
            Filters []ProductFilters  `json:"filters"`
            Items []ProductListingDetail  `json:"items"`
            SortOn []ProductSortOn  `json:"sort_on"`
            Page Page  `json:"page"`
         
    }
    
    // CollectionItem used by Catalog
    type CollectionItem struct {

        
            ItemID float64  `json:"item_id"`
            Priority float64  `json:"priority"`
            Action string  `json:"action"`
         
    }
    
    // CollectionItemUpdate used by Catalog
    type CollectionItemUpdate struct {

        
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Items []CollectionItem  `json:"items"`
            Query []CollectionQuery  `json:"query"`
            AllowSort bool  `json:"allow_sort"`
            Type string  `json:"type"`
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
            Count float64  `json:"count"`
            OutOfStockCount float64  `json:"out_of_stock_count"`
         
    }
    
    // CatalogInsightBrand used by Catalog
    type CatalogInsightBrand struct {

        
            AvailableArticles float64  `json:"available_articles"`
            Name string  `json:"name"`
            TotalArticles float64  `json:"total_articles"`
            TotalSizes float64  `json:"total_sizes"`
            AvailableSizes float64  `json:"available_sizes"`
            ArticleFreshness float64  `json:"article_freshness"`
         
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

        
            Enabled bool  `json:"enabled"`
            OptLevel string  `json:"opt_level"`
            BrandIds []float64  `json:"brand_ids"`
            Platform string  `json:"platform"`
            CompanyID float64  `json:"company_id"`
            StoreIds []float64  `json:"store_ids"`
         
    }
    
    // CompanyOptIn used by Catalog
    type CompanyOptIn struct {

        
            ModifiedOn float64  `json:"modified_on"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Enabled bool  `json:"enabled"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            OptLevel string  `json:"opt_level"`
            BrandIds []float64  `json:"brand_ids"`
            Platform string  `json:"platform"`
            CreatedOn float64  `json:"created_on"`
            CompanyID float64  `json:"company_id"`
            StoreIds []float64  `json:"store_ids"`
         
    }
    
    // GetOptInPlatform used by Catalog
    type GetOptInPlatform struct {

        
            Items []CompanyOptIn  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // OptinCompanyDetail used by Catalog
    type OptinCompanyDetail struct {

        
            UID float64  `json:"uid"`
            BusinessType string  `json:"business_type"`
            CompanyType string  `json:"company_type"`
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

        
            UID float64  `json:"uid"`
            Address map[string]interface{}  `json:"address"`
            ModifiedOn string  `json:"modified_on"`
            DisplayName string  `json:"display_name"`
            StoreCode string  `json:"store_code"`
            Name string  `json:"name"`
            Timing map[string]interface{}  `json:"timing"`
            AdditionalContacts []map[string]interface{}  `json:"additional_contacts"`
            StoreType string  `json:"store_type"`
            CreatedOn string  `json:"created_on"`
            CompanyID float64  `json:"company_id"`
            Documents []map[string]interface{}  `json:"documents"`
            Manager map[string]interface{}  `json:"manager"`
         
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

        
            Priority float64  `json:"priority"`
            Indexing bool  `json:"indexing"`
            DependsOn []string  `json:"depends_on"`
         
    }
    
    // AttributeSchemaRange used by Catalog
    type AttributeSchemaRange struct {

        
            Min float64  `json:"min"`
            Max float64  `json:"max"`
         
    }
    
    // AttributeMaster used by Catalog
    type AttributeMaster struct {

        
            AllowedValues []string  `json:"allowed_values"`
            Format string  `json:"format"`
            Multi bool  `json:"multi"`
            Type string  `json:"type"`
            Range AttributeSchemaRange  `json:"range"`
            Mandatory bool  `json:"mandatory"`
         
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

        
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Details AttributeMasterDetails  `json:"details"`
            IsNested bool  `json:"is_nested"`
            ID string  `json:"id"`
            Logo string  `json:"logo"`
            Departments []string  `json:"departments"`
            Filters AttributeMasterFilter  `json:"filters"`
            Schema AttributeMaster  `json:"schema"`
            EnabledForEndConsumer bool  `json:"enabled_for_end_consumer"`
            Meta AttributeMasterMeta  `json:"meta"`
            Description string  `json:"description"`
         
    }
    
    // CategoriesResponse used by Catalog
    type CategoriesResponse struct {

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
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

        
            Code string  `json:"code"`
            Message string  `json:"message"`
            Meta map[string]interface{}  `json:"meta"`
            Errors map[string]interface{}  `json:"errors"`
            Status float64  `json:"status"`
         
    }
    
    // UserSerializer used by Catalog
    type UserSerializer struct {

        
            UID string  `json:"uid"`
            Username string  `json:"username"`
            Contact string  `json:"contact"`
            UserID string  `json:"user_id"`
            ID string  `json:"_id"`
         
    }
    
    // GetDepartment used by Catalog
    type GetDepartment struct {

        
            UID float64  `json:"uid"`
            PriorityOrder float64  `json:"priority_order"`
            ModifiedOn string  `json:"modified_on"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            CreatedBy UserSerializer  `json:"created_by"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            Logo string  `json:"logo"`
            PageNo float64  `json:"page_no"`
            Search string  `json:"search"`
            ItemType string  `json:"item_type"`
            CreatedOn string  `json:"created_on"`
            Synonyms []string  `json:"synonyms"`
            PageSize float64  `json:"page_size"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // DepartmentsResponse used by Catalog
    type DepartmentsResponse struct {

        
            Items []GetDepartment  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // DepartmentErrorResponse used by Catalog
    type DepartmentErrorResponse struct {

        
            Code string  `json:"code"`
            Message string  `json:"message"`
            Meta map[string]interface{}  `json:"meta"`
            Errors map[string]interface{}  `json:"errors"`
            Status float64  `json:"status"`
         
    }
    
    // DepartmentCreateUpdate used by Catalog
    type DepartmentCreateUpdate struct {

        
            UID float64  `json:"uid"`
            PriorityOrder float64  `json:"priority_order"`
            Tags []string  `json:"tags"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Cls string  `json:"_cls"`
            Logo string  `json:"logo"`
            Synonyms []string  `json:"synonyms"`
            Platforms map[string]interface{}  `json:"platforms"`
            IsActive bool  `json:"is_active"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
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

        
            SuperUser bool  `json:"super_user"`
            Contact string  `json:"contact"`
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
         
    }
    
    // DepartmentModel used by Catalog
    type DepartmentModel struct {

        
            UID float64  `json:"uid"`
            PriorityOrder float64  `json:"priority_order"`
            ModifiedOn string  `json:"modified_on"`
            VerifiedBy UserDetail  `json:"verified_by"`
            Name interface{}  `json:"name"`
            CreatedBy UserDetail  `json:"created_by"`
            Slug interface{}  `json:"slug"`
            ModifiedBy UserDetail  `json:"modified_by"`
            Cls interface{}  `json:"_cls"`
            VerifiedOn string  `json:"verified_on"`
            Logo interface{}  `json:"logo"`
            CreatedOn string  `json:"created_on"`
            Synonyms []interface{}  `json:"synonyms"`
            ID interface{}  `json:"_id"`
            IsActive bool  `json:"is_active"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // ProductTemplate used by Catalog
    type ProductTemplate struct {

        
            Attributes []string  `json:"attributes"`
            ModifiedOn string  `json:"modified_on"`
            Tag string  `json:"tag"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Categories []string  `json:"categories"`
            Logo string  `json:"logo"`
            IsExpirable bool  `json:"is_expirable"`
            Departments []string  `json:"departments"`
            IsArchived bool  `json:"is_archived"`
            CreatedOn string  `json:"created_on"`
            IsActive bool  `json:"is_active"`
            Description string  `json:"description"`
            IsPhysical bool  `json:"is_physical"`
         
    }
    
    // TemplatesResponse used by Catalog
    type TemplatesResponse struct {

        
            Items ProductTemplate  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // TemplateDetails used by Catalog
    type TemplateDetails struct {

        
            Attributes []string  `json:"attributes"`
            Tag string  `json:"tag"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            ID string  `json:"id"`
            Categories []string  `json:"categories"`
            Logo string  `json:"logo"`
            IsExpirable bool  `json:"is_expirable"`
            Departments []string  `json:"departments"`
            IsArchived bool  `json:"is_archived"`
            IsActive bool  `json:"is_active"`
            Description string  `json:"description"`
            IsPhysical bool  `json:"is_physical"`
         
    }
    
    // Properties used by Catalog
    type Properties struct {

        
            ItemCode map[string]interface{}  `json:"item_code"`
            MultiSize map[string]interface{}  `json:"multi_size"`
            SizeGuide map[string]interface{}  `json:"size_guide"`
            Slug map[string]interface{}  `json:"slug"`
            CategorySlug map[string]interface{}  `json:"category_slug"`
            ShortDescription map[string]interface{}  `json:"short_description"`
            TraderType map[string]interface{}  `json:"trader_type"`
            Command map[string]interface{}  `json:"command"`
            HsnCode map[string]interface{}  `json:"hsn_code"`
            Description map[string]interface{}  `json:"description"`
            Trader map[string]interface{}  `json:"trader"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            BrandUID map[string]interface{}  `json:"brand_uid"`
            Currency map[string]interface{}  `json:"currency"`
            IsDependent map[string]interface{}  `json:"is_dependent"`
            NoOfBoxes map[string]interface{}  `json:"no_of_boxes"`
            Variants map[string]interface{}  `json:"variants"`
            Tags map[string]interface{}  `json:"tags"`
            Name map[string]interface{}  `json:"name"`
            ProductGroupTag map[string]interface{}  `json:"product_group_tag"`
            CountryOfOrigin map[string]interface{}  `json:"country_of_origin"`
            ItemType map[string]interface{}  `json:"item_type"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            Media map[string]interface{}  `json:"media"`
            ReturnConfig map[string]interface{}  `json:"return_config"`
            Highlights map[string]interface{}  `json:"highlights"`
            IsActive map[string]interface{}  `json:"is_active"`
            ProductPublish map[string]interface{}  `json:"product_publish"`
            Sizes map[string]interface{}  `json:"sizes"`
         
    }
    
    // GlobalValidation used by Catalog
    type GlobalValidation struct {

        
            Definitions map[string]interface{}  `json:"definitions"`
            Type string  `json:"type"`
            Description string  `json:"description"`
            Required []string  `json:"required"`
            Properties Properties  `json:"properties"`
            Title string  `json:"title"`
         
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

        
            Email string  `json:"email"`
            UID string  `json:"uid"`
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
         
    }
    
    // ProductTemplateExportResponse used by Catalog
    type ProductTemplateExportResponse struct {

        
            ModifiedOn string  `json:"modified_on"`
            TaskID string  `json:"task_id"`
            CreatedBy UserInfo1  `json:"created_by"`
            NotificationEmails []string  `json:"notification_emails"`
            Filters map[string]interface{}  `json:"filters"`
            URL string  `json:"url"`
            CompletedOn string  `json:"completed_on"`
            Type string  `json:"type"`
            SellerID float64  `json:"seller_id"`
            Status string  `json:"status"`
         
    }
    
    // ProductDownloadsResponse used by Catalog
    type ProductDownloadsResponse struct {

        
            Items []ProductTemplateExportResponse  `json:"items"`
         
    }
    
    // ProductTemplateExportFilterRequest used by Catalog
    type ProductTemplateExportFilterRequest struct {

        
            ToDate string  `json:"to_date"`
            CatalogueTypes []string  `json:"catalogue_types"`
            Templates []string  `json:"templates"`
            FromDate string  `json:"from_date"`
            Brands []string  `json:"brands"`
         
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
    
    // Hierarchy used by Catalog
    type Hierarchy struct {

        
            L1 float64  `json:"l1"`
            Department float64  `json:"department"`
            L2 float64  `json:"l2"`
         
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
    
    // Media1 used by Catalog
    type Media1 struct {

        
            Portrait string  `json:"portrait"`
            Landscape string  `json:"landscape"`
            Logo string  `json:"logo"`
         
    }
    
    // Category used by Catalog
    type Category struct {

        
            UID float64  `json:"uid"`
            Level float64  `json:"level"`
            ModifiedOn string  `json:"modified_on"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Priority float64  `json:"priority"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            ID string  `json:"id"`
            Departments []float64  `json:"departments"`
            Hierarchy []Hierarchy  `json:"hierarchy"`
            CreatedOn string  `json:"created_on"`
            Tryouts []string  `json:"tryouts"`
            Marketplaces CategoryMapping  `json:"marketplaces"`
            Synonyms []string  `json:"synonyms"`
            IsActive bool  `json:"is_active"`
            Media Media1  `json:"media"`
         
    }
    
    // CategoryResponse used by Catalog
    type CategoryResponse struct {

        
            Items []Category  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // CategoryRequestBody used by Catalog
    type CategoryRequestBody struct {

        
            Level float64  `json:"level"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Priority float64  `json:"priority"`
            Departments []float64  `json:"departments"`
            Hierarchy []Hierarchy  `json:"hierarchy"`
            Tryouts []string  `json:"tryouts"`
            Marketplaces CategoryMapping  `json:"marketplaces"`
            Synonyms []string  `json:"synonyms"`
            IsActive bool  `json:"is_active"`
            Media Media1  `json:"media"`
         
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
    
    // Trader used by Catalog
    type Trader struct {

        
            Address []string  `json:"address"`
            Type string  `json:"type"`
            Name interface{}  `json:"name"`
         
    }
    
    // Logo used by Catalog
    type Logo struct {

        
            AspectRatio string  `json:"aspect_ratio"`
            AspectRatioF float64  `json:"aspect_ratio_f"`
            SecureURL string  `json:"secure_url"`
            URL string  `json:"url"`
         
    }
    
    // Brand used by Catalog
    type Brand struct {

        
            UID float64  `json:"uid"`
            Logo Logo  `json:"logo"`
            Name string  `json:"name"`
         
    }
    
    // Image used by Catalog
    type Image struct {

        
            AspectRatio string  `json:"aspect_ratio"`
            AspectRatioF float64  `json:"aspect_ratio_f"`
            SecureURL string  `json:"secure_url"`
            URL string  `json:"url"`
         
    }
    
    // NetQuantityResponse used by Catalog
    type NetQuantityResponse struct {

        
            Unit string  `json:"unit"`
            Value float64  `json:"value"`
         
    }
    
    // VerifiedBy used by Catalog
    type VerifiedBy struct {

        
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
         
    }
    
    // ReturnConfigResponse used by Catalog
    type ReturnConfigResponse struct {

        
            Unit string  `json:"unit"`
            Returnable bool  `json:"returnable"`
            Time float64  `json:"time"`
         
    }
    
    // ProductPublish used by Catalog
    type ProductPublish struct {

        
            IsSet bool  `json:"is_set"`
            ProductOnlineDate string  `json:"product_online_date"`
         
    }
    
    // ProductSchemaV2 used by Catalog
    type ProductSchemaV2 struct {

        
            UID float64  `json:"uid"`
            ItemCode string  `json:"item_code"`
            MultiSize bool  `json:"multi_size"`
            SizeGuide string  `json:"size_guide"`
            Slug string  `json:"slug"`
            CategorySlug string  `json:"category_slug"`
            PrimaryColor string  `json:"primary_color"`
            ShortDescription string  `json:"short_description"`
            HsnCode string  `json:"hsn_code"`
            IsSet bool  `json:"is_set"`
            Description string  `json:"description"`
            Trader []Trader  `json:"trader"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            BrandUID float64  `json:"brand_uid"`
            Currency string  `json:"currency"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            IsDependent bool  `json:"is_dependent"`
            Pending string  `json:"pending"`
            Moq map[string]interface{}  `json:"moq"`
            VariantMedia map[string]interface{}  `json:"variant_media"`
            NoOfBoxes float64  `json:"no_of_boxes"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Attributes map[string]interface{}  `json:"attributes"`
            Variants map[string]interface{}  `json:"variants"`
            Tags []string  `json:"tags"`
            Name string  `json:"name"`
            ID string  `json:"id"`
            TemplateTag string  `json:"template_tag"`
            CountryOfOrigin string  `json:"country_of_origin"`
            ProductGroupTag []string  `json:"product_group_tag"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Brand Brand  `json:"brand"`
            Departments []float64  `json:"departments"`
            VariantGroup map[string]interface{}  `json:"variant_group"`
            ItemType string  `json:"item_type"`
            L3Mapping []string  `json:"l3_mapping"`
            VerifiedOn string  `json:"verified_on"`
            IsImageLessProduct bool  `json:"is_image_less_product"`
            Images []Image  `json:"images"`
            IsExpirable bool  `json:"is_expirable"`
            Media []Media  `json:"media"`
            IsPhysical bool  `json:"is_physical"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            ModifiedOn string  `json:"modified_on"`
            Stage string  `json:"stage"`
            NetQuantity NetQuantityResponse  `json:"net_quantity"`
            VerifiedBy VerifiedBy  `json:"verified_by"`
            ReturnConfig ReturnConfigResponse  `json:"return_config"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ImageNature string  `json:"image_nature"`
            AllSizes []map[string]interface{}  `json:"all_sizes"`
            Highlights []string  `json:"highlights"`
            CreatedOn string  `json:"created_on"`
            CategoryUID float64  `json:"category_uid"`
            Color string  `json:"color"`
            Category map[string]interface{}  `json:"category"`
            AllCompanyIds []float64  `json:"all_company_ids"`
            CompanyID float64  `json:"company_id"`
            IsActive bool  `json:"is_active"`
            AllIdentifiers []string  `json:"all_identifiers"`
            ProductPublish ProductPublish  `json:"product_publish"`
            Sizes []map[string]interface{}  `json:"sizes"`
         
    }
    
    // ProductListingResponseV2 used by Catalog
    type ProductListingResponseV2 struct {

        
            Items []ProductSchemaV2  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // CustomOrder used by Catalog
    type CustomOrder struct {

        
            ManufacturingTime float64  `json:"manufacturing_time"`
            ManufacturingTimeUnit string  `json:"manufacturing_time_unit"`
            IsCustomOrder bool  `json:"is_custom_order"`
         
    }
    
    // TaxIdentifier used by Catalog
    type TaxIdentifier struct {

        
            HsnCode string  `json:"hsn_code"`
            HsnCodeID string  `json:"hsn_code_id"`
            ReportingHsn string  `json:"reporting_hsn"`
         
    }
    
    // TeaserTag used by Catalog
    type TeaserTag struct {

        
            URL string  `json:"url"`
            Tag string  `json:"tag"`
         
    }
    
    // NetQuantity used by Catalog
    type NetQuantity struct {

        
            Unit interface{}  `json:"unit"`
            Value float64  `json:"value"`
         
    }
    
    // ReturnConfig used by Catalog
    type ReturnConfig struct {

        
            Unit string  `json:"unit"`
            Returnable bool  `json:"returnable"`
            Time float64  `json:"time"`
         
    }
    
    // ProductPublish1 used by Catalog
    type ProductPublish1 struct {

        
            IsSet bool  `json:"is_set"`
            ProductOnlineDate string  `json:"product_online_date"`
         
    }
    
    // ProductCreateUpdateSchemaV2 used by Catalog
    type ProductCreateUpdateSchemaV2 struct {

        
            UID float64  `json:"uid"`
            ItemCode string  `json:"item_code"`
            MultiSize bool  `json:"multi_size"`
            SizeGuide string  `json:"size_guide"`
            Slug string  `json:"slug"`
            Requester string  `json:"requester"`
            CategorySlug string  `json:"category_slug"`
            BulkJobID string  `json:"bulk_job_id"`
            Action string  `json:"action"`
            ChangeRequestID interface{}  `json:"change_request_id"`
            ShortDescription string  `json:"short_description"`
            IsSet bool  `json:"is_set"`
            Description string  `json:"description"`
            Trader []Trader  `json:"trader"`
            CustomOrder CustomOrder  `json:"custom_order"`
            BrandUID float64  `json:"brand_uid"`
            Currency string  `json:"currency"`
            TaxIdentifier TaxIdentifier  `json:"tax_identifier"`
            IsDependent bool  `json:"is_dependent"`
            VariantMedia map[string]interface{}  `json:"variant_media"`
            NoOfBoxes float64  `json:"no_of_boxes"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Attributes map[string]interface{}  `json:"attributes"`
            Variants map[string]interface{}  `json:"variants"`
            Tags []string  `json:"tags"`
            Name string  `json:"name"`
            TemplateTag string  `json:"template_tag"`
            CountryOfOrigin string  `json:"country_of_origin"`
            ProductGroupTag []string  `json:"product_group_tag"`
            Departments []float64  `json:"departments"`
            VariantGroup map[string]interface{}  `json:"variant_group"`
            ItemType string  `json:"item_type"`
            IsImageLessProduct bool  `json:"is_image_less_product"`
            TeaserTag TeaserTag  `json:"teaser_tag"`
            Media []Media  `json:"media"`
            NetQuantity NetQuantity  `json:"net_quantity"`
            ReturnConfig ReturnConfig  `json:"return_config"`
            Highlights []string  `json:"highlights"`
            CompanyID float64  `json:"company_id"`
            IsActive bool  `json:"is_active"`
            ProductPublish ProductPublish1  `json:"product_publish"`
            Sizes []map[string]interface{}  `json:"sizes"`
         
    }
    
    // ProductVariants used by Catalog
    type ProductVariants struct {

        
            BrandUID float64  `json:"brand_uid"`
            UID float64  `json:"uid"`
            ItemCode string  `json:"item_code"`
            Name string  `json:"name"`
            CategoryUID float64  `json:"category_uid"`
            Media []Media  `json:"media"`
         
    }
    
    // ProductVariantsResponse used by Catalog
    type ProductVariantsResponse struct {

        
            Variants []ProductVariants  `json:"variants"`
            Page Page  `json:"page"`
         
    }
    
    // AttributeMasterSerializer used by Catalog
    type AttributeMasterSerializer struct {

        
            Slug string  `json:"slug"`
            Filters AttributeMasterFilter  `json:"filters"`
            Synonyms map[string]interface{}  `json:"synonyms"`
            Description string  `json:"description"`
            Suggestion string  `json:"suggestion"`
            Unit string  `json:"unit"`
            Tags []string  `json:"tags"`
            Name string  `json:"name"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Logo string  `json:"logo"`
            Departments []string  `json:"departments"`
            Example string  `json:"example"`
            Schema AttributeMaster  `json:"schema"`
            Variant bool  `json:"variant"`
            RawKey string  `json:"raw_key"`
            ModifiedOn string  `json:"modified_on"`
            Details AttributeMasterDetails  `json:"details"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            IsNested bool  `json:"is_nested"`
            CreatedOn string  `json:"created_on"`
            EnabledForEndConsumer bool  `json:"enabled_for_end_consumer"`
         
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

        
            Size interface{}  `json:"size"`
            ItemWeightUnitOfMeasure interface{}  `json:"item_weight_unit_of_measure"`
            ItemWeight float64  `json:"item_weight"`
            ItemDimensionsUnitOfMeasure string  `json:"item_dimensions_unit_of_measure"`
            ItemHeight float64  `json:"item_height"`
            ItemWidth float64  `json:"item_width"`
            Identifiers []ValidateIdentifier  `json:"identifiers"`
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

        
            IsSet bool  `json:"is_set"`
            ProductOnlineDate float64  `json:"product_online_date"`
         
    }
    
    // Product used by Catalog
    type Product struct {

        
            UID float64  `json:"uid"`
            ItemCode string  `json:"item_code"`
            MultiSize bool  `json:"multi_size"`
            SizeGuide string  `json:"size_guide"`
            Slug string  `json:"slug"`
            CategorySlug string  `json:"category_slug"`
            PrimaryColor string  `json:"primary_color"`
            ShortDescription string  `json:"short_description"`
            HsnCode string  `json:"hsn_code"`
            IsSet bool  `json:"is_set"`
            Description string  `json:"description"`
            Trader []Trader  `json:"trader"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            BrandUID float64  `json:"brand_uid"`
            Currency string  `json:"currency"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            IsDependent bool  `json:"is_dependent"`
            Pending string  `json:"pending"`
            Moq map[string]interface{}  `json:"moq"`
            VariantMedia map[string]interface{}  `json:"variant_media"`
            NoOfBoxes float64  `json:"no_of_boxes"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Attributes map[string]interface{}  `json:"attributes"`
            Variants map[string]interface{}  `json:"variants"`
            Tags []string  `json:"tags"`
            Name string  `json:"name"`
            ID string  `json:"id"`
            TemplateTag string  `json:"template_tag"`
            CountryOfOrigin string  `json:"country_of_origin"`
            ProductGroupTag []string  `json:"product_group_tag"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Brand Brand  `json:"brand"`
            Departments []float64  `json:"departments"`
            VariantGroup map[string]interface{}  `json:"variant_group"`
            ItemType string  `json:"item_type"`
            L3Mapping []string  `json:"l3_mapping"`
            VerifiedOn string  `json:"verified_on"`
            IsImageLessProduct bool  `json:"is_image_less_product"`
            Images []Image  `json:"images"`
            IsExpirable bool  `json:"is_expirable"`
            Media []Media  `json:"media"`
            IsPhysical bool  `json:"is_physical"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            ModifiedOn string  `json:"modified_on"`
            Stage string  `json:"stage"`
            NetQuantity NetQuantityResponse  `json:"net_quantity"`
            VerifiedBy VerifiedBy  `json:"verified_by"`
            ReturnConfig ReturnConfigResponse  `json:"return_config"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ImageNature string  `json:"image_nature"`
            AllSizes []map[string]interface{}  `json:"all_sizes"`
            Highlights []string  `json:"highlights"`
            CreatedOn string  `json:"created_on"`
            CategoryUID float64  `json:"category_uid"`
            Color string  `json:"color"`
            Category map[string]interface{}  `json:"category"`
            AllCompanyIds []float64  `json:"all_company_ids"`
            CompanyID float64  `json:"company_id"`
            IsActive bool  `json:"is_active"`
            AllIdentifiers []string  `json:"all_identifiers"`
            ProductPublish ProductPublished  `json:"product_publish"`
            Sizes []map[string]interface{}  `json:"sizes"`
         
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

        
            FailedRecords []string  `json:"failed_records"`
            Failed float64  `json:"failed"`
            Stage string  `json:"stage"`
            ModifiedOn string  `json:"modified_on"`
            Cancelled float64  `json:"cancelled"`
            CreatedBy UserDetail1  `json:"created_by"`
            TemplateTag string  `json:"template_tag"`
            Succeed float64  `json:"succeed"`
            FilePath string  `json:"file_path"`
            ModifiedBy UserDetail1  `json:"modified_by"`
            CreatedOn string  `json:"created_on"`
            CompanyID float64  `json:"company_id"`
            Template ProductTemplate  `json:"template"`
            IsActive bool  `json:"is_active"`
            Total float64  `json:"total"`
            CancelledRecords []string  `json:"cancelled_records"`
         
    }
    
    // ProductBulkRequestList used by Catalog
    type ProductBulkRequestList struct {

        
            Items ProductBulkRequest  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // BulkJob used by Catalog
    type BulkJob struct {

        
            FailedRecords []map[string]interface{}  `json:"failed_records"`
            Failed float64  `json:"failed"`
            ModifiedOn string  `json:"modified_on"`
            Stage string  `json:"stage"`
            Cancelled float64  `json:"cancelled"`
            CreatedBy UserInfo1  `json:"created_by"`
            ModifiedBy UserInfo1  `json:"modified_by"`
            Succeed float64  `json:"succeed"`
            FilePath string  `json:"file_path"`
            TemplateTag string  `json:"template_tag"`
            CustomTemplateTag string  `json:"custom_template_tag"`
            CreatedOn string  `json:"created_on"`
            CompanyID float64  `json:"company_id"`
            TrackingURL string  `json:"tracking_url"`
            IsActive bool  `json:"is_active"`
            Total float64  `json:"total"`
            CancelledRecords []map[string]interface{}  `json:"cancelled_records"`
         
    }
    
    // BulkResponse used by Catalog
    type BulkResponse struct {

        
            ModifiedOn string  `json:"modified_on"`
            CreatedBy UserInfo1  `json:"created_by"`
            ModifiedBy UserInfo1  `json:"modified_by"`
            BatchID string  `json:"batch_id"`
            CreatedOn string  `json:"created_on"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // BulkProductRequest used by Catalog
    type BulkProductRequest struct {

        
            TemplateTag string  `json:"template_tag"`
            BatchID string  `json:"batch_id"`
            CompanyID float64  `json:"company_id"`
            Data []map[string]interface{}  `json:"data"`
         
    }
    
    // ProductTagsViewResponse used by Catalog
    type ProductTagsViewResponse struct {

        
            Items []string  `json:"items"`
         
    }
    
    // UserCommon used by Catalog
    type UserCommon struct {

        
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
            CompanyID float64  `json:"company_id"`
         
    }
    
    // Items used by Catalog
    type Items struct {

        
            Retry float64  `json:"retry"`
            Failed float64  `json:"failed"`
            ModifiedOn string  `json:"modified_on"`
            Stage string  `json:"stage"`
            Cancelled float64  `json:"cancelled"`
            ID string  `json:"id"`
            CreatedBy UserCommon  `json:"created_by"`
            ModifiedBy UserCommon  `json:"modified_by"`
            Succeed float64  `json:"succeed"`
            FilePath string  `json:"file_path"`
            IsActive bool  `json:"is_active"`
            CreatedOn string  `json:"created_on"`
            CompanyID float64  `json:"company_id"`
            TrackingURL string  `json:"tracking_url"`
            FailedRecords []string  `json:"failed_records"`
            Total float64  `json:"total"`
            CancelledRecords []string  `json:"cancelled_records"`
         
    }
    
    // BulkAssetResponse used by Catalog
    type BulkAssetResponse struct {

        
            Items []Items  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // ProductBulkAssets used by Catalog
    type ProductBulkAssets struct {

        
            User map[string]interface{}  `json:"user"`
            URL string  `json:"url"`
            CompanyID float64  `json:"company_id"`
            BatchID string  `json:"batch_id"`
         
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
    
    // InventoryResponse used by Catalog
    type InventoryResponse struct {

        
            UID string  `json:"uid"`
            Size string  `json:"size"`
            Quantity float64  `json:"quantity"`
            ItemID float64  `json:"item_id"`
            Currency string  `json:"currency"`
            Price float64  `json:"price"`
            InventoryUpdatedOn string  `json:"inventory_updated_on"`
            PriceTransfer float64  `json:"price_transfer"`
            Identifiers map[string]interface{}  `json:"identifiers"`
            SellableQuantity float64  `json:"sellable_quantity"`
            SellerIdentifier string  `json:"seller_identifier"`
            PriceEffective float64  `json:"price_effective"`
            Store map[string]interface{}  `json:"store"`
         
    }
    
    // InventoryResponsePaginated used by Catalog
    type InventoryResponsePaginated struct {

        
            Items []InventoryResponse  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // ItemQuery used by Catalog
    type ItemQuery struct {

        
            UID float64  `json:"uid"`
            ItemCode string  `json:"item_code"`
            BrandUID float64  `json:"brand_uid"`
         
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
    
    // GTIN used by Catalog
    type GTIN struct {

        
            Primary bool  `json:"primary"`
            GtinType string  `json:"gtin_type"`
            GtinValue interface{}  `json:"gtin_value"`
         
    }
    
    // InvSize used by Catalog
    type InvSize struct {

        
            Size interface{}  `json:"size"`
            Quantity float64  `json:"quantity"`
            Set InventorySet  `json:"set"`
            StoreCode string  `json:"store_code"`
            PriceEffective float64  `json:"price_effective"`
            ItemWeightUnitOfMeasure string  `json:"item_weight_unit_of_measure"`
            ItemWeight float64  `json:"item_weight"`
            Currency string  `json:"currency"`
            ItemDimensionsUnitOfMeasure string  `json:"item_dimensions_unit_of_measure"`
            Price float64  `json:"price"`
            ItemHeight float64  `json:"item_height"`
            ItemWidth float64  `json:"item_width"`
            PriceTransfer float64  `json:"price_transfer"`
            Identifiers []GTIN  `json:"identifiers"`
            ItemLength float64  `json:"item_length"`
            IsSet bool  `json:"is_set"`
            ExpirationDate string  `json:"expiration_date"`
         
    }
    
    // InventoryRequest used by Catalog
    type InventoryRequest struct {

        
            Item ItemQuery  `json:"item"`
            CompanyID float64  `json:"company_id"`
            Sizes []InvSize  `json:"sizes"`
         
    }
    
    // DimensionResponse used by Catalog
    type DimensionResponse struct {

        
            Unit string  `json:"unit"`
            Width float64  `json:"width"`
            Length float64  `json:"length"`
            IsDefault bool  `json:"is_default"`
            Height float64  `json:"height"`
         
    }
    
    // QuantityBase used by Catalog
    type QuantityBase struct {

        
            Count float64  `json:"count"`
            UpdatedAt string  `json:"updated_at"`
         
    }
    
    // Quantities used by Catalog
    type Quantities struct {

        
            Damaged QuantityBase  `json:"damaged"`
            Sellable QuantityBase  `json:"sellable"`
            NotAvailable QuantityBase  `json:"not_available"`
            OrderCommitted QuantityBase  `json:"order_committed"`
         
    }
    
    // Trader1 used by Catalog
    type Trader1 struct {

        
            Address []string  `json:"address"`
            Type string  `json:"type"`
            Name string  `json:"name"`
         
    }
    
    // CompanyMeta used by Catalog
    type CompanyMeta struct {

        
            ID float64  `json:"id"`
         
    }
    
    // ManufacturerResponse used by Catalog
    type ManufacturerResponse struct {

        
            Address string  `json:"address"`
            IsDefault bool  `json:"is_default"`
            Name string  `json:"name"`
         
    }
    
    // StoreMeta used by Catalog
    type StoreMeta struct {

        
            ID float64  `json:"id"`
         
    }
    
    // BrandMeta used by Catalog
    type BrandMeta struct {

        
            Name string  `json:"name"`
            ID float64  `json:"id"`
         
    }
    
    // WeightResponse used by Catalog
    type WeightResponse struct {

        
            Unit string  `json:"unit"`
            IsDefault bool  `json:"is_default"`
            Shipping float64  `json:"shipping"`
         
    }
    
    // ReturnConfig1 used by Catalog
    type ReturnConfig1 struct {

        
            Unit string  `json:"unit"`
            Returnable bool  `json:"returnable"`
            Time float64  `json:"time"`
         
    }
    
    // PriceMeta used by Catalog
    type PriceMeta struct {

        
            Currency string  `json:"currency"`
            Marked float64  `json:"marked"`
            Effective float64  `json:"effective"`
            TpNotes map[string]interface{}  `json:"tp_notes"`
            Transfer float64  `json:"transfer"`
            UpdatedAt string  `json:"updated_at"`
         
    }
    
    // InventorySellerResponse used by Catalog
    type InventorySellerResponse struct {

        
            UID string  `json:"uid"`
            FyndMeta map[string]interface{}  `json:"fynd_meta"`
            TotalQuantity float64  `json:"total_quantity"`
            RawMeta map[string]interface{}  `json:"raw_meta"`
            Dimension DimensionResponse  `json:"dimension"`
            IsSet bool  `json:"is_set"`
            Meta map[string]interface{}  `json:"meta"`
            Quantities Quantities  `json:"quantities"`
            Trader []Trader1  `json:"trader"`
            Fragile bool  `json:"fragile"`
            Company CompanyMeta  `json:"company"`
            Set InventorySet  `json:"set"`
            ItemID float64  `json:"item_id"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            Manufacturer ManufacturerResponse  `json:"manufacturer"`
            Store StoreMeta  `json:"store"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Size string  `json:"size"`
            TraceID string  `json:"trace_id"`
            Tags []string  `json:"tags"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            CountryOfOrigin string  `json:"country_of_origin"`
            Brand BrandMeta  `json:"brand"`
            AddedOnStore string  `json:"added_on_store"`
            SellerIdentifier string  `json:"seller_identifier"`
            ExpirationDate string  `json:"expiration_date"`
            Identifier map[string]interface{}  `json:"identifier"`
            Weight WeightResponse  `json:"weight"`
            Stage string  `json:"stage"`
            FyndItemCode string  `json:"fynd_item_code"`
            ReturnConfig ReturnConfig1  `json:"return_config"`
            CreatedBy UserSerializer  `json:"created_by"`
            TrackInventory bool  `json:"track_inventory"`
            Price PriceMeta  `json:"price"`
            FyndArticleCode string  `json:"fynd_article_code"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // InventorySellerIdentifierResponsePaginated used by Catalog
    type InventorySellerIdentifierResponsePaginated struct {

        
            Items []InventorySellerResponse  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // DimensionResponse1 used by Catalog
    type DimensionResponse1 struct {

        
            Unit string  `json:"unit"`
            Width float64  `json:"width"`
            Length float64  `json:"length"`
            Height float64  `json:"height"`
         
    }
    
    // Quantity used by Catalog
    type Quantity struct {

        
            Count float64  `json:"count"`
         
    }
    
    // QuantitiesArticle used by Catalog
    type QuantitiesArticle struct {

        
            Damaged Quantity  `json:"damaged"`
            Sellable Quantity  `json:"sellable"`
            NotAvailable Quantity  `json:"not_available"`
            OrderCommitted Quantity  `json:"order_committed"`
         
    }
    
    // Trader2 used by Catalog
    type Trader2 struct {

        
            Address []string  `json:"address"`
            Type string  `json:"type"`
            Name string  `json:"name"`
         
    }
    
    // CompanyMeta1 used by Catalog
    type CompanyMeta1 struct {

        
            ID float64  `json:"id"`
         
    }
    
    // DateMeta used by Catalog
    type DateMeta struct {

        
            AddedOnStore string  `json:"added_on_store"`
            ModifiedOn string  `json:"modified_on"`
            InventoryUpdatedOn string  `json:"inventory_updated_on"`
            CreatedOn string  `json:"created_on"`
         
    }
    
    // ManufacturerResponse1 used by Catalog
    type ManufacturerResponse1 struct {

        
            Address string  `json:"address"`
            IsDefault bool  `json:"is_default"`
            Name string  `json:"name"`
         
    }
    
    // ArticleStoreResponse used by Catalog
    type ArticleStoreResponse struct {

        
            UID float64  `json:"uid"`
            StoreType string  `json:"store_type"`
            StoreCode string  `json:"store_code"`
            Name string  `json:"name"`
         
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
    
    // ReturnConfig2 used by Catalog
    type ReturnConfig2 struct {

        
            Unit string  `json:"unit"`
            Returnable bool  `json:"returnable"`
            Time float64  `json:"time"`
         
    }
    
    // PriceArticle used by Catalog
    type PriceArticle struct {

        
            Currency string  `json:"currency"`
            Marked float64  `json:"marked"`
            Effective float64  `json:"effective"`
            TpNotes map[string]interface{}  `json:"tp_notes"`
            Transfer float64  `json:"transfer"`
         
    }
    
    // GetInventories used by Catalog
    type GetInventories struct {

        
            UID string  `json:"uid"`
            TotalQuantity float64  `json:"total_quantity"`
            Dimension DimensionResponse1  `json:"dimension"`
            IsSet bool  `json:"is_set"`
            Quantities QuantitiesArticle  `json:"quantities"`
            Trader []Trader2  `json:"trader"`
            Company CompanyMeta1  `json:"company"`
            ItemID float64  `json:"item_id"`
            DateMeta DateMeta  `json:"date_meta"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            InventoryUpdatedOn string  `json:"inventory_updated_on"`
            Platforms map[string]interface{}  `json:"platforms"`
            Manufacturer ManufacturerResponse1  `json:"manufacturer"`
            Store ArticleStoreResponse  `json:"store"`
            Size string  `json:"size"`
            TraceID string  `json:"trace_id"`
            Tags []string  `json:"tags"`
            ID string  `json:"id"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            CountryOfOrigin string  `json:"country_of_origin"`
            Brand BrandMeta1  `json:"brand"`
            SellerIdentifier string  `json:"seller_identifier"`
            ExpirationDate string  `json:"expiration_date"`
            Identifier map[string]interface{}  `json:"identifier"`
            Weight WeightResponse1  `json:"weight"`
            Stage string  `json:"stage"`
            ReturnConfig ReturnConfig2  `json:"return_config"`
            CreatedBy UserSerializer  `json:"created_by"`
            TrackInventory bool  `json:"track_inventory"`
            Price PriceArticle  `json:"price"`
         
    }
    
    // GetInventoriesResponse used by Catalog
    type GetInventoriesResponse struct {

        
            Items []GetInventories  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // BulkInventoryGetItems used by Catalog
    type BulkInventoryGetItems struct {

        
            Failed float64  `json:"failed"`
            ModifiedOn string  `json:"modified_on"`
            Stage string  `json:"stage"`
            Cancelled float64  `json:"cancelled"`
            ID string  `json:"id"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Succeed float64  `json:"succeed"`
            FilePath string  `json:"file_path"`
            IsActive bool  `json:"is_active"`
            CreatedOn string  `json:"created_on"`
            CompanyID float64  `json:"company_id"`
            FailedRecords []string  `json:"failed_records"`
            Total float64  `json:"total"`
            CancelledRecords []string  `json:"cancelled_records"`
         
    }
    
    // BulkInventoryGet used by Catalog
    type BulkInventoryGet struct {

        
            Items []BulkInventoryGetItems  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // InventoryJobPayload used by Catalog
    type InventoryJobPayload struct {

        
            TraceID string  `json:"trace_id"`
            Quantity float64  `json:"quantity"`
            StoreCode string  `json:"store_code"`
            Tags []string  `json:"tags"`
            PriceMarked float64  `json:"price_marked"`
            Currency string  `json:"currency"`
            TotalQuantity float64  `json:"total_quantity"`
            ItemWeightUnitOfMeasure string  `json:"item_weight_unit_of_measure"`
            ItemDimensionsUnitOfMeasure string  `json:"item_dimensions_unit_of_measure"`
            Price float64  `json:"price"`
            SellerIdentifier string  `json:"seller_identifier"`
            PriceEffective float64  `json:"price_effective"`
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
            Operators string  `json:"operators"`
            Max float64  `json:"max"`
         
    }
    
    // InventoryExportAdvanceOption used by Catalog
    type InventoryExportAdvanceOption struct {

        
            ToDate string  `json:"to_date"`
            Quantity InventoryExportQuantityFilter  `json:"quantity"`
            FromDate string  `json:"from_date"`
            BrandIds []float64  `json:"brand_ids"`
            StoreIds []float64  `json:"store_ids"`
         
    }
    
    // InventoryExportJob used by Catalog
    type InventoryExportJob struct {

        
            TaskID string  `json:"task_id"`
            NotificationEmails []string  `json:"notification_emails"`
            Filters InventoryExportAdvanceOption  `json:"filters"`
            URL string  `json:"url"`
            CompletedOn string  `json:"completed_on"`
            Type string  `json:"type"`
            SellerID float64  `json:"seller_id"`
            Status string  `json:"status"`
         
    }
    
    // InventoryExportRequest used by Catalog
    type InventoryExportRequest struct {

        
            Type string  `json:"type"`
            Brand []float64  `json:"brand"`
            Store []float64  `json:"store"`
         
    }
    
    // InventoryExportResponse used by Catalog
    type InventoryExportResponse struct {

        
            ModifiedOn string  `json:"modified_on"`
            TaskID string  `json:"task_id"`
            CreatedBy string  `json:"created_by"`
            NotificationEmails []string  `json:"notification_emails"`
            Filters map[string]interface{}  `json:"filters"`
            CreatedOn string  `json:"created_on"`
            Type string  `json:"type"`
            SellerID float64  `json:"seller_id"`
            Status string  `json:"status"`
         
    }
    
    // InventoryJobFilters used by Catalog
    type InventoryJobFilters struct {

        
            ToDate string  `json:"to_date"`
            Quantity InventoryExportQuantityFilter  `json:"quantity"`
            FromDate string  `json:"from_date"`
            Stores []string  `json:"stores"`
            Brands []string  `json:"brands"`
         
    }
    
    // InventoryJobDetailResponse used by Catalog
    type InventoryJobDetailResponse struct {

        
            ModifiedOn string  `json:"modified_on"`
            TaskID string  `json:"task_id"`
            ID string  `json:"id"`
            CreatedBy UserDetail  `json:"created_by"`
            NotificationEmails []string  `json:"notification_emails"`
            CancelledOn string  `json:"cancelled_on"`
            Filters InventoryJobFilters  `json:"filters"`
            URL string  `json:"url"`
            CompletedOn string  `json:"completed_on"`
            CreatedOn string  `json:"created_on"`
            Type string  `json:"type"`
            CancelledBy UserDetail  `json:"cancelled_by"`
            SellerID float64  `json:"seller_id"`
            Status interface{}  `json:"status"`
         
    }
    
    // InventoryExportJobListResponse used by Catalog
    type InventoryExportJobListResponse struct {

        
            Items InventoryJobDetailResponse  `json:"items"`
         
    }
    
    // InventoryExportFilter used by Catalog
    type InventoryExportFilter struct {

        
            ToDate string  `json:"to_date"`
            Quantity InventoryExportQuantityFilter  `json:"quantity"`
            FromDate string  `json:"from_date"`
            BrandIds []float64  `json:"brand_ids"`
            StoreIds []float64  `json:"store_ids"`
         
    }
    
    // InventoryCreateRequest used by Catalog
    type InventoryCreateRequest struct {

        
            Filters InventoryExportFilter  `json:"filters"`
            Data []string  `json:"data"`
            Type string  `json:"type"`
            NotificationEmails []string  `json:"notification_emails"`
         
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

        
            TraceID string  `json:"trace_id"`
            Tags []string  `json:"tags"`
            PriceMarked float64  `json:"price_marked"`
            TotalQuantity float64  `json:"total_quantity"`
            StoreID float64  `json:"store_id"`
            SellerIdentifier string  `json:"seller_identifier"`
            PriceEffective float64  `json:"price_effective"`
            ExpirationDate string  `json:"expiration_date"`
         
    }
    
    // InventoryRequestSchemaV2 used by Catalog
    type InventoryRequestSchemaV2 struct {

        
            Meta map[string]interface{}  `json:"meta"`
            Payload []InventoryPayload  `json:"payload"`
            CompanyID float64  `json:"company_id"`
         
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

        
            Message string  `json:"message"`
            Items []InventoryResponseItem  `json:"items"`
         
    }
    
    // HsnCodesObject used by Catalog
    type HsnCodesObject struct {

        
            Threshold2 float64  `json:"threshold2"`
            ModifiedOn string  `json:"modified_on"`
            Hs2Code string  `json:"hs2_code"`
            ID string  `json:"id"`
            Threshold1 float64  `json:"threshold1"`
            CompanyID float64  `json:"company_id"`
            Tax1 float64  `json:"tax1"`
            HsnCode string  `json:"hsn_code"`
            Tax2 float64  `json:"tax2"`
            TaxOnEsp bool  `json:"tax_on_esp"`
            TaxOnMrp bool  `json:"tax_on_mrp"`
         
    }
    
    // HsnCode used by Catalog
    type HsnCode struct {

        
            Data HsnCodesObject  `json:"data"`
         
    }
    
    // HsnUpsert used by Catalog
    type HsnUpsert struct {

        
            UID float64  `json:"uid"`
            Threshold2 float64  `json:"threshold2"`
            Hs2Code string  `json:"hs2_code"`
            Threshold1 float64  `json:"threshold1"`
            IsActive bool  `json:"is_active"`
            CompanyID float64  `json:"company_id"`
            Tax1 float64  `json:"tax1"`
            HsnCode string  `json:"hsn_code"`
            Tax2 float64  `json:"tax2"`
            TaxOnEsp bool  `json:"tax_on_esp"`
            TaxOnMrp bool  `json:"tax_on_mrp"`
         
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
            EffectiveDate string  `json:"effective_date"`
            Threshold float64  `json:"threshold"`
         
    }
    
    // HSNDataInsertV2 used by Catalog
    type HSNDataInsertV2 struct {

        
            ModifiedOn string  `json:"modified_on"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            CountryCode string  `json:"country_code"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Type string  `json:"type"`
            Taxes []TaxSlab  `json:"taxes"`
            CreatedOn string  `json:"created_on"`
            HsnCode string  `json:"hsn_code"`
            Description string  `json:"description"`
            ReportingHsn string  `json:"reporting_hsn"`
            HsnCodeID string  `json:"hsn_code_id"`
         
    }
    
    // PageResponse used by Catalog
    type PageResponse struct {

        
            Size float64  `json:"size"`
            ItemTotal float64  `json:"item_total"`
            HasNext bool  `json:"has_next"`
            Current string  `json:"current"`
            HasPrevious bool  `json:"has_previous"`
         
    }
    
    // HsnCodesListingResponseSchemaV2 used by Catalog
    type HsnCodesListingResponseSchemaV2 struct {

        
            Items []HSNDataInsertV2  `json:"items"`
            Page PageResponse  `json:"page"`
         
    }
    
    // Media2 used by Catalog
    type Media2 struct {

        
            AspectRatio string  `json:"aspect_ratio"`
            URL string  `json:"url"`
            Type string  `json:"type"`
         
    }
    
    // BrandItem used by Catalog
    type BrandItem struct {

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Logo Media2  `json:"logo"`
            Action Action  `json:"action"`
            Departments []string  `json:"departments"`
            Discount string  `json:"discount"`
            Banners ImageUrls  `json:"banners"`
         
    }
    
    // BrandListingResponse used by Catalog
    type BrandListingResponse struct {

        
            Items []BrandItem  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // Department used by Catalog
    type Department struct {

        
            UID float64  `json:"uid"`
            PriorityOrder float64  `json:"priority_order"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Logo Media2  `json:"logo"`
         
    }
    
    // DepartmentResponse used by Catalog
    type DepartmentResponse struct {

        
            Items []Department  `json:"items"`
         
    }
    
    // ThirdLevelChild used by Catalog
    type ThirdLevelChild struct {

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Banners ImageUrls  `json:"banners"`
            Action Action  `json:"action"`
            Childs []map[string]interface{}  `json:"childs"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // SecondLevelChild used by Catalog
    type SecondLevelChild struct {

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Banners ImageUrls  `json:"banners"`
            Action Action  `json:"action"`
            Childs []ThirdLevelChild  `json:"childs"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // Child used by Catalog
    type Child struct {

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Banners ImageUrls  `json:"banners"`
            Action Action  `json:"action"`
            Childs []SecondLevelChild  `json:"childs"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // CategoryItems used by Catalog
    type CategoryItems struct {

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Action Action  `json:"action"`
            Childs []Child  `json:"childs"`
            Banners ImageUrls  `json:"banners"`
         
    }
    
    // DepartmentCategoryTree used by Catalog
    type DepartmentCategoryTree struct {

        
            Items []CategoryItems  `json:"items"`
            Department string  `json:"department"`
         
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

        
            Items []ProductListingDetail  `json:"items"`
            SortOn []ProductSortOn  `json:"sort_on"`
            Filters []ProductFilters  `json:"filters"`
            Operators map[string]interface{}  `json:"operators"`
            Page Page  `json:"page"`
         
    }
    
    // ProductDetail used by Catalog
    type ProductDetail struct {

        
            UID float64  `json:"uid"`
            HasVariant bool  `json:"has_variant"`
            ItemCode string  `json:"item_code"`
            Medias []Media  `json:"medias"`
            Slug string  `json:"slug"`
            RatingCount float64  `json:"rating_count"`
            PromoMeta map[string]interface{}  `json:"promo_meta"`
            ShortDescription string  `json:"short_description"`
            Description string  `json:"description"`
            Rating float64  `json:"rating"`
            ProductOnlineDate string  `json:"product_online_date"`
            Similars []string  `json:"similars"`
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
            Attributes map[string]interface{}  `json:"attributes"`
            Name string  `json:"name"`
            Brand ProductBrand  `json:"brand"`
            ItemType string  `json:"item_type"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            ImageNature string  `json:"image_nature"`
            Highlights []string  `json:"highlights"`
            Type string  `json:"type"`
            Tryouts []string  `json:"tryouts"`
            Color string  `json:"color"`
         
    }
    
    // InventoryPage used by Catalog
    type InventoryPage struct {

        
            ItemTotal float64  `json:"item_total"`
            HasNext bool  `json:"has_next"`
            Type string  `json:"type"`
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

        
            Size string  `json:"size"`
            IgnoredStores []float64  `json:"ignored_stores"`
            ItemID float64  `json:"item_id"`
         
    }
    
    // ArticleAssignment used by Catalog
    type ArticleAssignment struct {

        
            Level string  `json:"level"`
            Strategy string  `json:"strategy"`
         
    }
    
    // AssignStoreArticle used by Catalog
    type AssignStoreArticle struct {

        
            Quantity float64  `json:"quantity"`
            Query ArticleQuery  `json:"query"`
            GroupID string  `json:"group_id"`
            Meta map[string]interface{}  `json:"meta"`
            ArticleAssignment ArticleAssignment  `json:"article_assignment"`
         
    }
    
    // AssignStore used by Catalog
    type AssignStore struct {

        
            AppID string  `json:"app_id"`
            ChannelType string  `json:"channel_type"`
            ChannelIdentifier string  `json:"channel_identifier"`
            Articles []AssignStoreArticle  `json:"articles"`
            CompanyID float64  `json:"company_id"`
            StoreIds []float64  `json:"store_ids"`
            Pincode string  `json:"pincode"`
         
    }
    
    // ArticleAssignment1 used by Catalog
    type ArticleAssignment1 struct {

        
            Level string  `json:"level"`
            Strategy string  `json:"strategy"`
         
    }
    
    // StoreAssignResponse used by Catalog
    type StoreAssignResponse struct {

        
            UID string  `json:"uid"`
            Size string  `json:"size"`
            Quantity float64  `json:"quantity"`
            PriceEffective float64  `json:"price_effective"`
            PriceMarked float64  `json:"price_marked"`
            ItemID float64  `json:"item_id"`
            StrategyWiseListing []map[string]interface{}  `json:"strategy_wise_listing"`
            Index float64  `json:"index"`
            GroupID string  `json:"group_id"`
            StoreID float64  `json:"store_id"`
            StorePincode float64  `json:"store_pincode"`
            CompanyID float64  `json:"company_id"`
            SCity string  `json:"s_city"`
            ID string  `json:"_id"`
            Meta map[string]interface{}  `json:"meta"`
            ArticleAssignment ArticleAssignment1  `json:"article_assignment"`
            Status bool  `json:"status"`
         
    }
    
    // GetAddressSerializer used by Catalog
    type GetAddressSerializer struct {

        
            City string  `json:"city"`
            AddressType string  `json:"address_type"`
            State string  `json:"state"`
            Latitude float64  `json:"latitude"`
            CountryCode string  `json:"country_code"`
            Country string  `json:"country"`
            Address2 string  `json:"address2"`
            Longitude float64  `json:"longitude"`
            Address1 string  `json:"address1"`
            Landmark string  `json:"landmark"`
            Pincode float64  `json:"pincode"`
         
    }
    
    // UserSerializer1 used by Catalog
    type UserSerializer1 struct {

        
            Contact string  `json:"contact"`
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
         
    }
    
    // GetCompanySerializer used by Catalog
    type GetCompanySerializer struct {

        
            UID float64  `json:"uid"`
            ModifiedOn string  `json:"modified_on"`
            VerifiedBy UserSerializer1  `json:"verified_by"`
            CompanyType string  `json:"company_type"`
            Stage string  `json:"stage"`
            Name string  `json:"name"`
            CreatedBy UserSerializer1  `json:"created_by"`
            ModifiedBy UserSerializer1  `json:"modified_by"`
            BusinessType string  `json:"business_type"`
            VerifiedOn string  `json:"verified_on"`
            CreatedOn string  `json:"created_on"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            RejectReason string  `json:"reject_reason"`
         
    }
    
    // ProductReturnConfigSerializer used by Catalog
    type ProductReturnConfigSerializer struct {

        
            StoreUID float64  `json:"store_uid"`
            OnSameStore bool  `json:"on_same_store"`
         
    }
    
    // Document used by Catalog
    type Document struct {

        
            Verified bool  `json:"verified"`
            LegalName string  `json:"legal_name"`
            URL string  `json:"url"`
            Value string  `json:"value"`
            Type string  `json:"type"`
         
    }
    
    // UserSerializer2 used by Catalog
    type UserSerializer2 struct {

        
            Contact string  `json:"contact"`
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
         
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
    
    // SellerPhoneNumber used by Catalog
    type SellerPhoneNumber struct {

        
            CountryCode float64  `json:"country_code"`
            Number string  `json:"number"`
         
    }
    
    // LocationManagerSerializer used by Catalog
    type LocationManagerSerializer struct {

        
            Email string  `json:"email"`
            MobileNo SellerPhoneNumber  `json:"mobile_no"`
            Name string  `json:"name"`
         
    }
    
    // LocationIntegrationType used by Catalog
    type LocationIntegrationType struct {

        
            Inventory string  `json:"inventory"`
            Order string  `json:"order"`
         
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
            Opening LocationTimingSerializer  `json:"opening"`
            Closing LocationTimingSerializer  `json:"closing"`
         
    }
    
    // GetLocationSerializer used by Catalog
    type GetLocationSerializer struct {

        
            UID float64  `json:"uid"`
            Address GetAddressSerializer  `json:"address"`
            NotificationEmails []string  `json:"notification_emails"`
            PhoneNumber string  `json:"phone_number"`
            Company GetCompanySerializer  `json:"company"`
            ProductReturnConfig ProductReturnConfigSerializer  `json:"product_return_config"`
            StoreType string  `json:"store_type"`
            Documents []Document  `json:"documents"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            DisplayName string  `json:"display_name"`
            Name string  `json:"name"`
            ModifiedBy UserSerializer2  `json:"modified_by"`
            VerifiedOn string  `json:"verified_on"`
            Warnings map[string]interface{}  `json:"warnings"`
            GstCredentials InvoiceDetailsSerializer  `json:"gst_credentials"`
            Manager LocationManagerSerializer  `json:"manager"`
            ModifiedOn string  `json:"modified_on"`
            VerifiedBy UserSerializer2  `json:"verified_by"`
            Stage string  `json:"stage"`
            Code string  `json:"code"`
            CreatedBy UserSerializer2  `json:"created_by"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            IntegrationType LocationIntegrationType  `json:"integration_type"`
            Timing []LocationDayWiseSerializer  `json:"timing"`
            CreatedOn string  `json:"created_on"`
         
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
            Name string  `json:"name"`
            Logo string  `json:"logo"`
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
    

    
    // BusinessCountryInfo used by CompanyProfile
    type BusinessCountryInfo struct {

        
            CountryCode string  `json:"country_code"`
            Country string  `json:"country"`
         
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
            Value string  `json:"value"`
            Type string  `json:"type"`
            LegalName string  `json:"legal_name"`
            Verified bool  `json:"verified"`
         
    }
    
    // CompanyTaxesSerializer used by CompanyProfile
    type CompanyTaxesSerializer struct {

        
            Rate float64  `json:"rate"`
            EffectiveDate string  `json:"effective_date"`
            Enable bool  `json:"enable"`
         
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

        
            Latitude float64  `json:"latitude"`
            AddressType string  `json:"address_type"`
            Landmark string  `json:"landmark"`
            State string  `json:"state"`
            Longitude float64  `json:"longitude"`
            CountryCode string  `json:"country_code"`
            Address2 string  `json:"address2"`
            City string  `json:"city"`
            Address1 string  `json:"address1"`
            Pincode float64  `json:"pincode"`
            Country string  `json:"country"`
         
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
    
    // GetCompanyProfileSerializerResponse used by CompanyProfile
    type GetCompanyProfileSerializerResponse struct {

        
            UID float64  `json:"uid"`
            BusinessCountryInfo BusinessCountryInfo  `json:"business_country_info"`
            Warnings map[string]interface{}  `json:"warnings"`
            CompanyType string  `json:"company_type"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            NotificationEmails []string  `json:"notification_emails"`
            Mode string  `json:"mode"`
            CreatedBy UserSerializer  `json:"created_by"`
            Documents []Document  `json:"documents"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            Taxes []CompanyTaxesSerializer  `json:"taxes"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            CreatedOn string  `json:"created_on"`
            Name string  `json:"name"`
            FranchiseEnabled bool  `json:"franchise_enabled"`
            VerifiedOn string  `json:"verified_on"`
            BusinessInfo string  `json:"business_info"`
            BusinessDetails BusinessDetails  `json:"business_details"`
            Stage string  `json:"stage"`
            ModifiedOn string  `json:"modified_on"`
            BusinessType string  `json:"business_type"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            ContactDetails ContactDetails  `json:"contact_details"`
         
    }
    
    // ErrorResponse used by CompanyProfile
    type ErrorResponse struct {

        
            Status float64  `json:"status"`
            Meta map[string]interface{}  `json:"meta"`
            Code string  `json:"code"`
            Message string  `json:"message"`
         
    }
    
    // CreateUpdateAddressSerializer used by CompanyProfile
    type CreateUpdateAddressSerializer struct {

        
            Latitude float64  `json:"latitude"`
            AddressType string  `json:"address_type"`
            Landmark string  `json:"landmark"`
            State string  `json:"state"`
            Longitude float64  `json:"longitude"`
            CountryCode string  `json:"country_code"`
            Address2 string  `json:"address2"`
            City string  `json:"city"`
            Address1 string  `json:"address1"`
            Pincode float64  `json:"pincode"`
            Country string  `json:"country"`
         
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
            BusinessDetails BusinessDetails  `json:"business_details"`
            Name string  `json:"name"`
            NotificationEmails []string  `json:"notification_emails"`
            Warnings map[string]interface{}  `json:"warnings"`
            CompanyType string  `json:"company_type"`
            Addresses []CreateUpdateAddressSerializer  `json:"addresses"`
            BusinessType string  `json:"business_type"`
            ContactDetails ContactDetails  `json:"contact_details"`
            FranchiseEnabled bool  `json:"franchise_enabled"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            RejectReason string  `json:"reject_reason"`
            Taxes []CompanyTaxesSerializer1  `json:"taxes"`
            BusinessInfo string  `json:"business_info"`
         
    }
    
    // ProfileSuccessResponse used by CompanyProfile
    type ProfileSuccessResponse struct {

        
            UID float64  `json:"uid"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // DocumentsObj used by CompanyProfile
    type DocumentsObj struct {

        
            Pending float64  `json:"pending"`
            Verified float64  `json:"verified"`
         
    }
    
    // MetricsSerializer used by CompanyProfile
    type MetricsSerializer struct {

        
            UID float64  `json:"uid"`
            StoreDocuments DocumentsObj  `json:"store_documents"`
            Stage string  `json:"stage"`
            Brand DocumentsObj  `json:"brand"`
            CompanyDocuments DocumentsObj  `json:"company_documents"`
            Product DocumentsObj  `json:"product"`
            Store DocumentsObj  `json:"store"`
         
    }
    
    // BrandBannerSerializer used by CompanyProfile
    type BrandBannerSerializer struct {

        
            Landscape string  `json:"landscape"`
            Portrait string  `json:"portrait"`
         
    }
    
    // GetBrandResponseSerializer used by CompanyProfile
    type GetBrandResponseSerializer struct {

        
            UID float64  `json:"uid"`
            Warnings map[string]interface{}  `json:"warnings"`
            Logo string  `json:"logo"`
            Synonyms []string  `json:"synonyms"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            Mode string  `json:"mode"`
            CreatedBy UserSerializer  `json:"created_by"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            Banner BrandBannerSerializer  `json:"banner"`
            CreatedOn string  `json:"created_on"`
            Name string  `json:"name"`
            VerifiedOn string  `json:"verified_on"`
            Stage string  `json:"stage"`
            Description string  `json:"description"`
            ModifiedOn string  `json:"modified_on"`
            SlugKey string  `json:"slug_key"`
            RejectReason string  `json:"reject_reason"`
         
    }
    
    // CreateUpdateBrandRequestSerializer used by CompanyProfile
    type CreateUpdateBrandRequestSerializer struct {

        
            UID float64  `json:"uid"`
            Banner BrandBannerSerializer  `json:"banner"`
            BrandTier string  `json:"brand_tier"`
            CompanyID float64  `json:"company_id"`
            Description string  `json:"description"`
            Name string  `json:"name"`
            Logo string  `json:"logo"`
            Synonyms []string  `json:"synonyms"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
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

        
            Socials []CompanySocialAccounts  `json:"socials"`
            WebsiteURL string  `json:"website_url"`
         
    }
    
    // CompanySerializer used by CompanyProfile
    type CompanySerializer struct {

        
            UID float64  `json:"uid"`
            Details CompanyDetails  `json:"details"`
            Stage string  `json:"stage"`
            BusinessCountryInfo BusinessCountryInfo  `json:"business_country_info"`
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
            Name string  `json:"name"`
            CreatedBy UserSerializer  `json:"created_by"`
            CompanyType string  `json:"company_type"`
            MarketChannels []string  `json:"market_channels"`
            BusinessType string  `json:"business_type"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            RejectReason string  `json:"reject_reason"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            NotificationEmails []string  `json:"notification_emails"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            VerifiedOn string  `json:"verified_on"`
            ModifiedBy UserSerializer  `json:"modified_by"`
         
    }
    
    // CompanyBrandSerializer used by CompanyProfile
    type CompanyBrandSerializer struct {

        
            UID float64  `json:"uid"`
            Stage string  `json:"stage"`
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
            Company CompanySerializer  `json:"company"`
            Warnings map[string]interface{}  `json:"warnings"`
            CreatedBy UserSerializer  `json:"created_by"`
            Brand GetBrandResponseSerializer  `json:"brand"`
            RejectReason string  `json:"reject_reason"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            VerifiedOn string  `json:"verified_on"`
            ModifiedBy UserSerializer  `json:"modified_by"`
         
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
    
    // LocationManagerSerializer used by CompanyProfile
    type LocationManagerSerializer struct {

        
            Name string  `json:"name"`
            MobileNo SellerPhoneNumber  `json:"mobile_no"`
            Email string  `json:"email"`
         
    }
    
    // ProductReturnConfigSerializer used by CompanyProfile
    type ProductReturnConfigSerializer struct {

        
            StoreUID float64  `json:"store_uid"`
            OnSameStore bool  `json:"on_same_store"`
         
    }
    
    // GetCompanySerializer used by CompanyProfile
    type GetCompanySerializer struct {

        
            UID float64  `json:"uid"`
            Stage string  `json:"stage"`
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
            Name string  `json:"name"`
            CreatedBy UserSerializer  `json:"created_by"`
            CompanyType string  `json:"company_type"`
            BusinessType string  `json:"business_type"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            RejectReason string  `json:"reject_reason"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            VerifiedOn string  `json:"verified_on"`
            ModifiedBy UserSerializer  `json:"modified_by"`
         
    }
    
    // HolidayDateSerializer used by CompanyProfile
    type HolidayDateSerializer struct {

        
            EndDate string  `json:"end_date"`
            StartDate string  `json:"start_date"`
         
    }
    
    // HolidaySchemaSerializer used by CompanyProfile
    type HolidaySchemaSerializer struct {

        
            Date HolidayDateSerializer  `json:"date"`
            Title string  `json:"title"`
            HolidayType string  `json:"holiday_type"`
         
    }
    
    // LocationTimingSerializer used by CompanyProfile
    type LocationTimingSerializer struct {

        
            Minute float64  `json:"minute"`
            Hour float64  `json:"hour"`
         
    }
    
    // LocationDayWiseSerializer used by CompanyProfile
    type LocationDayWiseSerializer struct {

        
            Weekday string  `json:"weekday"`
            Open bool  `json:"open"`
            Closing LocationTimingSerializer  `json:"closing"`
            Opening LocationTimingSerializer  `json:"opening"`
         
    }
    
    // GetLocationSerializer used by CompanyProfile
    type GetLocationSerializer struct {

        
            UID float64  `json:"uid"`
            StoreType string  `json:"store_type"`
            Address GetAddressSerializer  `json:"address"`
            Warnings map[string]interface{}  `json:"warnings"`
            AutoInvoice bool  `json:"auto_invoice"`
            GstCredentials InvoiceDetailsSerializer  `json:"gst_credentials"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            NotificationEmails []string  `json:"notification_emails"`
            Manager LocationManagerSerializer  `json:"manager"`
            ProductReturnConfig ProductReturnConfigSerializer  `json:"product_return_config"`
            Company GetCompanySerializer  `json:"company"`
            CreatedBy UserSerializer  `json:"created_by"`
            DisplayName string  `json:"display_name"`
            Documents []Document  `json:"documents"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            PhoneNumber string  `json:"phone_number"`
            CreatedOn string  `json:"created_on"`
            Name string  `json:"name"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            Holiday []HolidaySchemaSerializer  `json:"holiday"`
            Timing []LocationDayWiseSerializer  `json:"timing"`
            VerifiedOn string  `json:"verified_on"`
            Code string  `json:"code"`
            Stage string  `json:"stage"`
            ModifiedOn string  `json:"modified_on"`
            CreditNote bool  `json:"credit_note"`
         
    }
    
    // LocationListSerializer used by CompanyProfile
    type LocationListSerializer struct {

        
            Page Page  `json:"page"`
            Items []GetLocationSerializer  `json:"items"`
         
    }
    
    // AddressSerializer used by CompanyProfile
    type AddressSerializer struct {

        
            Latitude float64  `json:"latitude"`
            AddressType string  `json:"address_type"`
            Landmark string  `json:"landmark"`
            State string  `json:"state"`
            Longitude float64  `json:"longitude"`
            CountryCode string  `json:"country_code"`
            Address2 string  `json:"address2"`
            City string  `json:"city"`
            Address1 string  `json:"address1"`
            Pincode float64  `json:"pincode"`
            Country string  `json:"country"`
         
    }
    
    // LocationSerializer used by CompanyProfile
    type LocationSerializer struct {

        
            UID float64  `json:"uid"`
            StoreType string  `json:"store_type"`
            Address AddressSerializer  `json:"address"`
            Warnings map[string]interface{}  `json:"warnings"`
            AutoInvoice bool  `json:"auto_invoice"`
            GstCredentials InvoiceDetailsSerializer  `json:"gst_credentials"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            NotificationEmails []string  `json:"notification_emails"`
            Manager LocationManagerSerializer  `json:"manager"`
            ProductReturnConfig ProductReturnConfigSerializer  `json:"product_return_config"`
            Company float64  `json:"company"`
            DisplayName string  `json:"display_name"`
            Documents []Document  `json:"documents"`
            Holiday []HolidaySchemaSerializer  `json:"holiday"`
            Name string  `json:"name"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            Timing []LocationDayWiseSerializer  `json:"timing"`
            Code string  `json:"code"`
            Stage string  `json:"stage"`
            CreditNote bool  `json:"credit_note"`
         
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
            Phone InformationPhone  `json:"phone"`
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

        
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // Validity used by Cart
    type Validity struct {

        
            Priority float64  `json:"priority"`
         
    }
    
    // State used by Cart
    type State struct {

        
            IsDisplay bool  `json:"is_display"`
            IsPublic bool  `json:"is_public"`
            IsArchived bool  `json:"is_archived"`
         
    }
    
    // PriceRange used by Cart
    type PriceRange struct {

        
            Min float64  `json:"min"`
            Max float64  `json:"max"`
         
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
            Iins []string  `json:"iins"`
            Codes []string  `json:"codes"`
            Types []string  `json:"types"`
            Networks []string  `json:"networks"`
         
    }
    
    // PostOrder used by Cart
    type PostOrder struct {

        
            ReturnAllowed bool  `json:"return_allowed"`
            CancellationAllowed bool  `json:"cancellation_allowed"`
         
    }
    
    // Restrictions used by Cart
    type Restrictions struct {

        
            PriceRange PriceRange  `json:"price_range"`
            UserGroups []float64  `json:"user_groups"`
            OrderingStores []float64  `json:"ordering_stores"`
            Uses UsesRestriction  `json:"uses"`
            BulkBundle BulkBundleRestriction  `json:"bulk_bundle"`
            Payments map[string]PaymentModes  `json:"payments"`
            CouponAllowed bool  `json:"coupon_allowed"`
            Platforms []string  `json:"platforms"`
            UserType string  `json:"user_type"`
            PostOrder PostOrder  `json:"post_order"`
         
    }
    
    // Rule used by Cart
    type Rule struct {

        
            Min float64  `json:"min"`
            Key float64  `json:"key"`
            DiscountQty float64  `json:"discount_qty"`
            Value float64  `json:"value"`
            Max float64  `json:"max"`
         
    }
    
    // Ownership used by Cart
    type Ownership struct {

        
            PayableBy string  `json:"payable_by"`
            PayableCategory string  `json:"payable_category"`
         
    }
    
    // CouponSchedule used by Cart
    type CouponSchedule struct {

        
            Duration float64  `json:"duration"`
            Start string  `json:"start"`
            End string  `json:"end"`
            Cron string  `json:"cron"`
            NextSchedule []map[string]interface{}  `json:"next_schedule"`
         
    }
    
    // Validation used by Cart
    type Validation struct {

        
            AppID []string  `json:"app_id"`
            Anonymous bool  `json:"anonymous"`
            UserRegisteredAfter string  `json:"user_registered_after"`
         
    }
    
    // CouponAuthor used by Cart
    type CouponAuthor struct {

        
            ModifiedBy string  `json:"modified_by"`
            CreatedBy string  `json:"created_by"`
         
    }
    
    // CouponAction used by Cart
    type CouponAction struct {

        
            ActionDate string  `json:"action_date"`
            TxnMode string  `json:"txn_mode"`
         
    }
    
    // Identifier used by Cart
    type Identifier struct {

        
            CompanyID []float64  `json:"company_id"`
            ExcludeBrandID []float64  `json:"exclude_brand_id"`
            ArticleID []string  `json:"article_id"`
            BrandID []float64  `json:"brand_id"`
            StoreID []float64  `json:"store_id"`
            ItemID []float64  `json:"item_id"`
            CollectionID []string  `json:"collection_id"`
            CategoryID []float64  `json:"category_id"`
            UserID []string  `json:"user_id"`
            EmailDomain []string  `json:"email_domain"`
         
    }
    
    // DisplayMetaDict used by Cart
    type DisplayMetaDict struct {

        
            Title string  `json:"title"`
            Subtitle string  `json:"subtitle"`
         
    }
    
    // DisplayMeta used by Cart
    type DisplayMeta struct {

        
            Description string  `json:"description"`
            Auto DisplayMetaDict  `json:"auto"`
            Remove DisplayMetaDict  `json:"remove"`
            Title string  `json:"title"`
            Apply DisplayMetaDict  `json:"apply"`
            Subtitle string  `json:"subtitle"`
         
    }
    
    // RuleDefinition used by Cart
    type RuleDefinition struct {

        
            Scope []string  `json:"scope"`
            Type string  `json:"type"`
            IsExact bool  `json:"is_exact"`
            AutoApply bool  `json:"auto_apply"`
            CurrencyCode string  `json:"currency_code"`
            ValueType string  `json:"value_type"`
            CalculateOn string  `json:"calculate_on"`
            ApplicableOn string  `json:"applicable_on"`
         
    }
    
    // CouponAdd used by Cart
    type CouponAdd struct {

        
            DateMeta CouponDateMeta  `json:"date_meta"`
            Validity Validity  `json:"validity"`
            State State  `json:"state"`
            Restrictions Restrictions  `json:"restrictions"`
            Rule []Rule  `json:"rule"`
            Ownership Ownership  `json:"ownership"`
            Schedule CouponSchedule  `json:"_schedule"`
            Validation Validation  `json:"validation"`
            TypeSlug string  `json:"type_slug"`
            Author CouponAuthor  `json:"author"`
            Action CouponAction  `json:"action"`
            Code string  `json:"code"`
            Identifiers Identifier  `json:"identifiers"`
            DisplayMeta DisplayMeta  `json:"display_meta"`
            RuleDefinition RuleDefinition  `json:"rule_definition"`
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

        
            DateMeta CouponDateMeta  `json:"date_meta"`
            Validity Validity  `json:"validity"`
            State State  `json:"state"`
            Restrictions Restrictions  `json:"restrictions"`
            Rule []Rule  `json:"rule"`
            Ownership Ownership  `json:"ownership"`
            Schedule CouponSchedule  `json:"_schedule"`
            Validation Validation  `json:"validation"`
            TypeSlug string  `json:"type_slug"`
            Author CouponAuthor  `json:"author"`
            Action CouponAction  `json:"action"`
            Code string  `json:"code"`
            Identifiers Identifier  `json:"identifiers"`
            DisplayMeta DisplayMeta  `json:"display_meta"`
            RuleDefinition RuleDefinition  `json:"rule_definition"`
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

        
            GreaterThan float64  `json:"greater_than"`
            GreaterThanEquals float64  `json:"greater_than_equals"`
            LessThanEquals float64  `json:"less_than_equals"`
            Equals float64  `json:"equals"`
            LessThan float64  `json:"less_than"`
         
    }
    
    // ItemCriteria used by Cart
    type ItemCriteria struct {

        
            CartUniqueItemQuantity CompareObject  `json:"cart_unique_item_quantity"`
            ItemDepartment []float64  `json:"item_department"`
            ItemTags []string  `json:"item_tags"`
            ItemL1Category []float64  `json:"item_l1_category"`
            ItemSku []string  `json:"item_sku"`
            ItemCategory []float64  `json:"item_category"`
            ItemExcludeBrand []float64  `json:"item_exclude_brand"`
            ItemExcludeDepartment []float64  `json:"item_exclude_department"`
            ItemExcludeL1Category []float64  `json:"item_exclude_l1_category"`
            ItemCompany []float64  `json:"item_company"`
            ItemExcludeCategory []float64  `json:"item_exclude_category"`
            ItemL2Category []float64  `json:"item_l2_category"`
            BuyRules []string  `json:"buy_rules"`
            ItemSize []string  `json:"item_size"`
            ItemExcludeCompany []float64  `json:"item_exclude_company"`
            ItemStore []float64  `json:"item_store"`
            ItemBrand []float64  `json:"item_brand"`
            ItemExcludeSku []string  `json:"item_exclude_sku"`
            ItemID []float64  `json:"item_id"`
            CartQuantity CompareObject  `json:"cart_quantity"`
            CartTotal CompareObject  `json:"cart_total"`
            ItemExcludeID []float64  `json:"item_exclude_id"`
            AvailableZones []string  `json:"available_zones"`
            CartUniqueItemAmount CompareObject  `json:"cart_unique_item_amount"`
            ItemExcludeStore []float64  `json:"item_exclude_store"`
            ItemExcludeL2Category []float64  `json:"item_exclude_l2_category"`
            AllItems bool  `json:"all_items"`
         
    }
    
    // DiscountOffer used by Cart
    type DiscountOffer struct {

        
            MaxOfferQuantity float64  `json:"max_offer_quantity"`
            MaxUsagePerTransaction float64  `json:"max_usage_per_transaction"`
            DiscountPercentage float64  `json:"discount_percentage"`
            MaxDiscountAmount float64  `json:"max_discount_amount"`
            DiscountAmount float64  `json:"discount_amount"`
            PartialCanRet bool  `json:"partial_can_ret"`
            ApportionDiscount bool  `json:"apportion_discount"`
            Code string  `json:"code"`
            DiscountPrice float64  `json:"discount_price"`
            MinOfferQuantity float64  `json:"min_offer_quantity"`
         
    }
    
    // DiscountRule used by Cart
    type DiscountRule struct {

        
            DiscountType string  `json:"discount_type"`
            BuyCondition string  `json:"buy_condition"`
            ItemCriteria ItemCriteria  `json:"item_criteria"`
            Offer DiscountOffer  `json:"offer"`
         
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
    
    // UserRegistered used by Cart
    type UserRegistered struct {

        
            End string  `json:"end"`
            Start string  `json:"start"`
         
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
    
    // PostOrder1 used by Cart
    type PostOrder1 struct {

        
            ReturnAllowed bool  `json:"return_allowed"`
            CancellationAllowed bool  `json:"cancellation_allowed"`
         
    }
    
    // Restrictions1 used by Cart
    type Restrictions1 struct {

        
            UserGroups []float64  `json:"user_groups"`
            OrderingStores []float64  `json:"ordering_stores"`
            Uses UsesRestriction1  `json:"uses"`
            AnonymousUsers bool  `json:"anonymous_users"`
            UserID []string  `json:"user_id"`
            OrderQuantity float64  `json:"order_quantity"`
            UserRegistered UserRegistered  `json:"user_registered"`
            Payments []PromotionPaymentModes  `json:"payments"`
            Platforms []string  `json:"platforms"`
            PostOrder PostOrder1  `json:"post_order"`
         
    }
    
    // Ownership1 used by Cart
    type Ownership1 struct {

        
            PayableBy string  `json:"payable_by"`
            PayableCategory string  `json:"payable_category"`
         
    }
    
    // Visibility used by Cart
    type Visibility struct {

        
            CouponList bool  `json:"coupon_list"`
            Pdp bool  `json:"pdp"`
         
    }
    
    // PromotionAction used by Cart
    type PromotionAction struct {

        
            ActionDate string  `json:"action_date"`
            ActionType string  `json:"action_type"`
         
    }
    
    // PromotionSchedule used by Cart
    type PromotionSchedule struct {

        
            Published bool  `json:"published"`
            Duration float64  `json:"duration"`
            Start string  `json:"start"`
            End string  `json:"end"`
            Cron string  `json:"cron"`
            NextSchedule []map[string]interface{}  `json:"next_schedule"`
         
    }
    
    // PromotionAuthor used by Cart
    type PromotionAuthor struct {

        
            ModifiedBy string  `json:"modified_by"`
            CreatedBy string  `json:"created_by"`
         
    }
    
    // DisplayMeta1 used by Cart
    type DisplayMeta1 struct {

        
            OfferLabel string  `json:"offer_label"`
            OfferText string  `json:"offer_text"`
            Description string  `json:"description"`
            Name string  `json:"name"`
         
    }
    
    // PromotionListItem used by Cart
    type PromotionListItem struct {

        
            DateMeta PromotionDateMeta  `json:"date_meta"`
            CalculateOn string  `json:"calculate_on"`
            DiscountRules []DiscountRule  `json:"discount_rules"`
            Currency string  `json:"currency"`
            BuyRules map[string]ItemCriteria  `json:"buy_rules"`
            ApplyPriority float64  `json:"apply_priority"`
            PromotionType string  `json:"promotion_type"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Restrictions Restrictions1  `json:"restrictions"`
            Ownership Ownership1  `json:"ownership"`
            Visiblility Visibility  `json:"visiblility"`
            ApplicationID string  `json:"application_id"`
            PostOrderAction PromotionAction  `json:"post_order_action"`
            Stackable bool  `json:"stackable"`
            Code string  `json:"code"`
            PromoGroup string  `json:"promo_group"`
            Schedule PromotionSchedule  `json:"_schedule"`
            Mode string  `json:"mode"`
            ApplyAllDiscount bool  `json:"apply_all_discount"`
            Author PromotionAuthor  `json:"author"`
            ApplyExclusive string  `json:"apply_exclusive"`
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
            CalculateOn string  `json:"calculate_on"`
            DiscountRules []DiscountRule  `json:"discount_rules"`
            Currency string  `json:"currency"`
            BuyRules map[string]ItemCriteria  `json:"buy_rules"`
            ApplyPriority float64  `json:"apply_priority"`
            PromotionType string  `json:"promotion_type"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Restrictions Restrictions1  `json:"restrictions"`
            Ownership Ownership1  `json:"ownership"`
            Visiblility Visibility  `json:"visiblility"`
            ApplicationID string  `json:"application_id"`
            PostOrderAction PromotionAction  `json:"post_order_action"`
            Stackable bool  `json:"stackable"`
            Code string  `json:"code"`
            PromoGroup string  `json:"promo_group"`
            Schedule PromotionSchedule  `json:"_schedule"`
            Mode string  `json:"mode"`
            ApplyAllDiscount bool  `json:"apply_all_discount"`
            Author PromotionAuthor  `json:"author"`
            ApplyExclusive string  `json:"apply_exclusive"`
            DisplayMeta DisplayMeta1  `json:"display_meta"`
         
    }
    
    // PromotionUpdate used by Cart
    type PromotionUpdate struct {

        
            DateMeta PromotionDateMeta  `json:"date_meta"`
            CalculateOn string  `json:"calculate_on"`
            DiscountRules []DiscountRule  `json:"discount_rules"`
            Currency string  `json:"currency"`
            BuyRules map[string]ItemCriteria  `json:"buy_rules"`
            ApplyPriority float64  `json:"apply_priority"`
            PromotionType string  `json:"promotion_type"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Restrictions Restrictions1  `json:"restrictions"`
            Ownership Ownership1  `json:"ownership"`
            Visiblility Visibility  `json:"visiblility"`
            ApplicationID string  `json:"application_id"`
            PostOrderAction PromotionAction  `json:"post_order_action"`
            Stackable bool  `json:"stackable"`
            Code string  `json:"code"`
            PromoGroup string  `json:"promo_group"`
            Schedule PromotionSchedule  `json:"_schedule"`
            Mode string  `json:"mode"`
            ApplyAllDiscount bool  `json:"apply_all_discount"`
            Author PromotionAuthor  `json:"author"`
            ApplyExclusive string  `json:"apply_exclusive"`
            DisplayMeta DisplayMeta1  `json:"display_meta"`
         
    }
    
    // PromotionPartialUpdate used by Cart
    type PromotionPartialUpdate struct {

        
            Schedule PromotionSchedule  `json:"schedule"`
            Archive bool  `json:"archive"`
         
    }
    
    // ActivePromosResponse used by Cart
    type ActivePromosResponse struct {

        
            Type string  `json:"type"`
            EntitySlug string  `json:"entity_slug"`
            IsHidden bool  `json:"is_hidden"`
            EntityType string  `json:"entity_type"`
            Example string  `json:"example"`
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
            Title string  `json:"title"`
            Description string  `json:"description"`
            Subtitle string  `json:"subtitle"`
         
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

        
            MaxCartItems float64  `json:"max_cart_items"`
            Enabled bool  `json:"enabled"`
            MinCartValue float64  `json:"min_cart_value"`
            DeliveryCharges DeliveryCharges  `json:"delivery_charges"`
            BulkCoupons bool  `json:"bulk_coupons"`
            GiftDisplayText string  `json:"gift_display_text"`
            RevenueEngineCoupon bool  `json:"revenue_engine_coupon"`
            GiftPricing float64  `json:"gift_pricing"`
         
    }
    
    // CartMetaConfigAdd used by Cart
    type CartMetaConfigAdd struct {

        
            MaxCartItems float64  `json:"max_cart_items"`
            Enabled bool  `json:"enabled"`
            MinCartValue float64  `json:"min_cart_value"`
            DeliveryCharges DeliveryCharges  `json:"delivery_charges"`
            BulkCoupons bool  `json:"bulk_coupons"`
            GiftDisplayText string  `json:"gift_display_text"`
            RevenueEngineCoupon bool  `json:"revenue_engine_coupon"`
            GiftPricing float64  `json:"gift_pricing"`
         
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
    
    // BasePrice used by Cart
    type BasePrice struct {

        
            Effective float64  `json:"effective"`
            CurrencyCode string  `json:"currency_code"`
            CurrencySymbol string  `json:"currency_symbol"`
            Marked float64  `json:"marked"`
         
    }
    
    // ArticlePriceInfo used by Cart
    type ArticlePriceInfo struct {

        
            Converted BasePrice  `json:"converted"`
            Base BasePrice  `json:"base"`
         
    }
    
    // BaseInfo used by Cart
    type BaseInfo struct {

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // StoreInfo used by Cart
    type StoreInfo struct {

        
            StoreCode string  `json:"store_code"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // ProductArticle used by Cart
    type ProductArticle struct {

        
            ProductGroupTags []string  `json:"product_group_tags"`
            MtoQuantity float64  `json:"mto_quantity"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            Identifier map[string]interface{}  `json:"identifier"`
            GiftCard map[string]interface{}  `json:"gift_card"`
            Type string  `json:"type"`
            Price ArticlePriceInfo  `json:"price"`
            SellerIdentifier string  `json:"seller_identifier"`
            Seller BaseInfo  `json:"seller"`
            IsGiftVisible bool  `json:"is_gift_visible"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            Size string  `json:"size"`
            CartItemMeta map[string]interface{}  `json:"cart_item_meta"`
            Meta map[string]interface{}  `json:"meta"`
            Store StoreInfo  `json:"store"`
            UID string  `json:"uid"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // ProductPrice used by Cart
    type ProductPrice struct {

        
            Effective float64  `json:"effective"`
            AddOn float64  `json:"add_on"`
            Marked float64  `json:"marked"`
            CurrencyCode string  `json:"currency_code"`
            CurrencySymbol string  `json:"currency_symbol"`
            Selling float64  `json:"selling"`
         
    }
    
    // ProductPriceInfo used by Cart
    type ProductPriceInfo struct {

        
            Converted ProductPrice  `json:"converted"`
            Base ProductPrice  `json:"base"`
         
    }
    
    // ProductAvailabilitySize used by Cart
    type ProductAvailabilitySize struct {

        
            Display string  `json:"display"`
            Value string  `json:"value"`
            IsAvailable bool  `json:"is_available"`
         
    }
    
    // ProductAvailability used by Cart
    type ProductAvailability struct {

        
            OtherStoreQuantity float64  `json:"other_store_quantity"`
            OutOfStock bool  `json:"out_of_stock"`
            Deliverable bool  `json:"deliverable"`
            AvailableSizes []ProductAvailabilitySize  `json:"available_sizes"`
            IsValid bool  `json:"is_valid"`
            Sizes []string  `json:"sizes"`
         
    }
    
    // PromiseFormatted used by Cart
    type PromiseFormatted struct {

        
            Min string  `json:"min"`
            Max string  `json:"max"`
         
    }
    
    // PromiseTimestamp used by Cart
    type PromiseTimestamp struct {

        
            Min float64  `json:"min"`
            Max float64  `json:"max"`
         
    }
    
    // ShipmentPromise used by Cart
    type ShipmentPromise struct {

        
            Formatted PromiseFormatted  `json:"formatted"`
            Timestamp PromiseTimestamp  `json:"timestamp"`
         
    }
    
    // CartProductIdentifer used by Cart
    type CartProductIdentifer struct {

        
            Identifier string  `json:"identifier"`
         
    }
    
    // PromoMeta used by Cart
    type PromoMeta struct {

        
            Message string  `json:"message"`
         
    }
    
    // CategoryInfo used by Cart
    type CategoryInfo struct {

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
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
    
    // CartProduct used by Cart
    type CartProduct struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Type string  `json:"type"`
            Categories []CategoryInfo  `json:"categories"`
            TeaserTag Tags  `json:"teaser_tag"`
            Action ProductAction  `json:"action"`
            Name string  `json:"name"`
            Images []ProductImage  `json:"images"`
            Brand BaseInfo  `json:"brand"`
            UID float64  `json:"uid"`
            ItemCode string  `json:"item_code"`
            Slug string  `json:"slug"`
            Tags []string  `json:"tags"`
         
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
    
    // FreeGiftItem used by Cart
    type FreeGiftItem struct {

        
            ItemPriceDetails map[string]interface{}  `json:"item_price_details"`
            ItemName string  `json:"item_name"`
            ItemID float64  `json:"item_id"`
            ItemBrandName string  `json:"item_brand_name"`
            ItemSlug string  `json:"item_slug"`
            ItemImagesURL []string  `json:"item_images_url"`
         
    }
    
    // AppliedFreeArticles used by Cart
    type AppliedFreeArticles struct {

        
            ParentItemIdentifier string  `json:"parent_item_identifier"`
            Quantity float64  `json:"quantity"`
            FreeGiftItemDetails FreeGiftItem  `json:"free_gift_item_details"`
            ArticleID string  `json:"article_id"`
         
    }
    
    // BuyRules used by Cart
    type BuyRules struct {

        
            ItemCriteria map[string]interface{}  `json:"item_criteria"`
            CartConditions map[string]interface{}  `json:"cart_conditions"`
         
    }
    
    // AppliedPromotion used by Cart
    type AppliedPromotion struct {

        
            Amount float64  `json:"amount"`
            PromoID string  `json:"promo_id"`
            PromotionType string  `json:"promotion_type"`
            PromotionName string  `json:"promotion_name"`
            Ownership Ownership2  `json:"ownership"`
            ArticleQuantity float64  `json:"article_quantity"`
            PromotionGroup string  `json:"promotion_group"`
            DiscountRules []DiscountRulesApp  `json:"discount_rules"`
            MrpPromotion bool  `json:"mrp_promotion"`
            AppliedFreeArticles []AppliedFreeArticles  `json:"applied_free_articles"`
            BuyRules []BuyRules  `json:"buy_rules"`
            OfferText string  `json:"offer_text"`
         
    }
    
    // CouponDetails used by Cart
    type CouponDetails struct {

        
            DiscountTotalQuantity float64  `json:"discount_total_quantity"`
            DiscountSingleQuantity float64  `json:"discount_single_quantity"`
            Code string  `json:"code"`
         
    }
    
    // CartProductInfo used by Cart
    type CartProductInfo struct {

        
            Message string  `json:"message"`
            Article ProductArticle  `json:"article"`
            Price ProductPriceInfo  `json:"price"`
            Availability ProductAvailability  `json:"availability"`
            IsSet bool  `json:"is_set"`
            PricePerUnit ProductPriceInfo  `json:"price_per_unit"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
            PromoMeta PromoMeta  `json:"promo_meta"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            Product CartProduct  `json:"product"`
            CouponMessage string  `json:"coupon_message"`
            Key string  `json:"key"`
            Moq map[string]interface{}  `json:"moq"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            BulkOffer map[string]interface{}  `json:"bulk_offer"`
            PromotionsApplied []AppliedPromotion  `json:"promotions_applied"`
            Discount string  `json:"discount"`
            Coupon CouponDetails  `json:"coupon"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // DisplayBreakup used by Cart
    type DisplayBreakup struct {

        
            Message []string  `json:"message"`
            CurrencyCode string  `json:"currency_code"`
            Key string  `json:"key"`
            Display string  `json:"display"`
            Value float64  `json:"value"`
            CurrencySymbol string  `json:"currency_symbol"`
         
    }
    
    // CouponBreakup used by Cart
    type CouponBreakup struct {

        
            Message string  `json:"message"`
            MinimumCartValue float64  `json:"minimum_cart_value"`
            Type string  `json:"type"`
            CouponValue float64  `json:"coupon_value"`
            Value float64  `json:"value"`
            MaxDiscountValue float64  `json:"max_discount_value"`
            Code string  `json:"code"`
            CouponType string  `json:"coupon_type"`
            IsApplied bool  `json:"is_applied"`
            UID string  `json:"uid"`
            Description string  `json:"description"`
            SubTitle string  `json:"sub_title"`
            Title string  `json:"title"`
         
    }
    
    // RawBreakup used by Cart
    type RawBreakup struct {

        
            CodCharge float64  `json:"cod_charge"`
            Vog float64  `json:"vog"`
            GiftCard float64  `json:"gift_card"`
            FyndCash float64  `json:"fynd_cash"`
            GstCharges float64  `json:"gst_charges"`
            DeliveryCharge float64  `json:"delivery_charge"`
            MrpTotal float64  `json:"mrp_total"`
            Subtotal float64  `json:"subtotal"`
            Total float64  `json:"total"`
            Discount float64  `json:"discount"`
            Coupon float64  `json:"coupon"`
            YouSaved float64  `json:"you_saved"`
            ConvenienceFee float64  `json:"convenience_fee"`
         
    }
    
    // LoyaltyPoints used by Cart
    type LoyaltyPoints struct {

        
            Description string  `json:"description"`
            Applicable float64  `json:"applicable"`
            IsApplied bool  `json:"is_applied"`
            Total float64  `json:"total"`
         
    }
    
    // CartBreakup used by Cart
    type CartBreakup struct {

        
            Display []DisplayBreakup  `json:"display"`
            Coupon CouponBreakup  `json:"coupon"`
            Raw RawBreakup  `json:"raw"`
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

        
            CountryPhoneCode string  `json:"country_phone_code"`
            Pincode float64  `json:"pincode"`
            State string  `json:"state"`
            Address string  `json:"address"`
            City string  `json:"city"`
            AreaCodeSlug string  `json:"area_code_slug"`
            Email string  `json:"email"`
            Area string  `json:"area"`
            Meta map[string]interface{}  `json:"meta"`
            AreaCode string  `json:"area_code"`
            Name string  `json:"name"`
            Country string  `json:"country"`
            Landmark string  `json:"landmark"`
            AddressType string  `json:"address_type"`
            CountryCode string  `json:"country_code"`
            CountryIsoCode string  `json:"country_iso_code"`
            Phone float64  `json:"phone"`
         
    }
    
    // OpenApiCartServiceabilityRequest used by Cart
    type OpenApiCartServiceabilityRequest struct {

        
            CartItems []CartItem  `json:"cart_items"`
            ShippingAddress ShippingAddress  `json:"shipping_address"`
         
    }
    
    // OpenApiCartServiceabilityResponse used by Cart
    type OpenApiCartServiceabilityResponse struct {

        
            Message string  `json:"message"`
            Items []CartProductInfo  `json:"items"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            IsValid bool  `json:"is_valid"`
            BreakupValues CartBreakup  `json:"breakup_values"`
         
    }
    
    // OpenApiFiles used by Cart
    type OpenApiFiles struct {

        
            Values []string  `json:"values"`
            Key string  `json:"key"`
         
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
    
    // OpenApiOrderItem used by Cart
    type OpenApiOrderItem struct {

        
            PriceEffective float64  `json:"price_effective"`
            CashbackApplied float64  `json:"cashback_applied"`
            PriceMarked float64  `json:"price_marked"`
            AmountPaid float64  `json:"amount_paid"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            Size string  `json:"size"`
            ProductID float64  `json:"product_id"`
            CodCharges float64  `json:"cod_charges"`
            DeliveryCharges float64  `json:"delivery_charges"`
            Files []OpenApiFiles  `json:"files"`
            Meta CartItemMeta  `json:"meta"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            Discount float64  `json:"discount"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            Quantity float64  `json:"quantity"`
            EmployeeDiscount float64  `json:"employee_discount"`
            PaymentMethods []MultiTenderPaymentMethod  `json:"payment_methods"`
         
    }
    
    // OpenApiPlatformCheckoutReq used by Cart
    type OpenApiPlatformCheckoutReq struct {

        
            CashbackApplied float64  `json:"cashback_applied"`
            CouponValue float64  `json:"coupon_value"`
            Files []OpenApiFiles  `json:"files"`
            BillingAddress ShippingAddress  `json:"billing_address"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            ShippingAddress ShippingAddress  `json:"shipping_address"`
            PaymentMethods []MultiTenderPaymentMethod  `json:"payment_methods"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            CurrencyCode string  `json:"currency_code"`
            OrderID string  `json:"order_id"`
            CartItems []OpenApiOrderItem  `json:"cart_items"`
            CodCharges float64  `json:"cod_charges"`
            Gstin string  `json:"gstin"`
            DeliveryCharges float64  `json:"delivery_charges"`
            Comment string  `json:"comment"`
            CouponCode string  `json:"coupon_code"`
            CartValue float64  `json:"cart_value"`
            PaymentMode string  `json:"payment_mode"`
            Coupon string  `json:"coupon"`
            EmployeeDiscount map[string]interface{}  `json:"employee_discount"`
         
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

        
            Promotion map[string]interface{}  `json:"promotion"`
            CheckoutMode string  `json:"checkout_mode"`
            Articles []map[string]interface{}  `json:"articles"`
            BulkCouponDiscount float64  `json:"bulk_coupon_discount"`
            FyndCredits map[string]interface{}  `json:"fynd_credits"`
            PaymentMethods []map[string]interface{}  `json:"payment_methods"`
            PickUpCustomerDetails map[string]interface{}  `json:"pick_up_customer_details"`
            ExpireAt string  `json:"expire_at"`
            CodCharges map[string]interface{}  `json:"cod_charges"`
            Gstin string  `json:"gstin"`
            DeliveryCharges map[string]interface{}  `json:"delivery_charges"`
            Cashback map[string]interface{}  `json:"cashback"`
            PaymentMode string  `json:"payment_mode"`
            FcIndexMap []float64  `json:"fc_index_map"`
            LastModified string  `json:"last_modified"`
            UserID string  `json:"user_id"`
            CreatedOn string  `json:"created_on"`
            IsArchive bool  `json:"is_archive"`
            MergeQty bool  `json:"merge_qty"`
            IsDefault bool  `json:"is_default"`
            ID string  `json:"_id"`
            IsActive bool  `json:"is_active"`
            OrderID string  `json:"order_id"`
            BuyNow bool  `json:"buy_now"`
            Payments map[string]interface{}  `json:"payments"`
            Meta map[string]interface{}  `json:"meta"`
            Comment string  `json:"comment"`
            Discount float64  `json:"discount"`
            Shipments []map[string]interface{}  `json:"shipments"`
            CartValue float64  `json:"cart_value"`
            AppID string  `json:"app_id"`
            Coupon map[string]interface{}  `json:"coupon"`
            UID float64  `json:"uid"`
         
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

        
            PaymentIdentifier string  `json:"payment_identifier"`
            DefaultOptions string  `json:"default_options"`
            Enabled bool  `json:"enabled"`
         
    }
    
    // CartDetailResponse used by Cart
    type CartDetailResponse struct {

        
            Message string  `json:"message"`
            CheckoutMode string  `json:"checkout_mode"`
            Items []CartProductInfo  `json:"items"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            AppliedPromoDetails []AppliedPromotion  `json:"applied_promo_details"`
            Gstin string  `json:"gstin"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            Currency CartCurrency  `json:"currency"`
            PanNo string  `json:"pan_no"`
            LastModified string  `json:"last_modified"`
            IsValid bool  `json:"is_valid"`
            ID string  `json:"id"`
            PanConfig map[string]interface{}  `json:"pan_config"`
            BuyNow bool  `json:"buy_now"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            Comment string  `json:"comment"`
            CouponText string  `json:"coupon_text"`
         
    }
    
    // AddProductCart used by Cart
    type AddProductCart struct {

        
            ProductGroupTags []string  `json:"product_group_tags"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ArticleID string  `json:"article_id"`
            ItemSize string  `json:"item_size"`
            StoreID float64  `json:"store_id"`
            ItemID float64  `json:"item_id"`
            ParentItemIdentifiers []map[string]string  `json:"parent_item_identifiers"`
            SellerID float64  `json:"seller_id"`
            ArticleAssignment map[string]interface{}  `json:"article_assignment"`
            Pos bool  `json:"pos"`
            Display string  `json:"display"`
            Meta map[string]interface{}  `json:"meta"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // AddCartRequest used by Cart
    type AddCartRequest struct {

        
            NewCart bool  `json:"new_cart"`
            Items []AddProductCart  `json:"items"`
         
    }
    
    // AddCartDetailResponse used by Cart
    type AddCartDetailResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
            Partial bool  `json:"partial"`
            Cart CartDetailResponse  `json:"cart"`
         
    }
    
    // UpdateProductCart used by Cart
    type UpdateProductCart struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ArticleID string  `json:"article_id"`
            ItemSize string  `json:"item_size"`
            ItemID float64  `json:"item_id"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            Meta map[string]interface{}  `json:"meta"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            Quantity float64  `json:"quantity"`
            ItemIndex float64  `json:"item_index"`
         
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

        
            PromoID string  `json:"promo_id"`
            PromoAmount string  `json:"promo_amount"`
            RwrdTndr string  `json:"rwrd_tndr"`
            ItemList []map[string]interface{}  `json:"item_list"`
            PromoDesc string  `json:"promo_desc"`
         
    }
    
    // OverrideCartItem used by Cart
    type OverrideCartItem struct {

        
            PriceEffective float64  `json:"price_effective"`
            PriceMarked float64  `json:"price_marked"`
            AmountPaid float64  `json:"amount_paid"`
            Size string  `json:"size"`
            SellerIdentifier string  `json:"seller_identifier"`
            ItemID float64  `json:"item_id"`
            Discount float64  `json:"discount"`
            PromoList []OverrideCartItemPromo  `json:"promo_list"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // OverrideCheckoutReq used by Cart
    type OverrideCheckoutReq struct {

        
            CartID string  `json:"cart_id"`
            MerchantCode string  `json:"merchant_code"`
            CurrencyCode string  `json:"currency_code"`
            CartItems []OverrideCartItem  `json:"cart_items"`
            OrderType string  `json:"order_type"`
            PaymentIdentifier string  `json:"payment_identifier"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
            Aggregator string  `json:"aggregator"`
            PaymentMode string  `json:"payment_mode"`
            ShippingAddress map[string]interface{}  `json:"shipping_address"`
            OrderingStore float64  `json:"ordering_store"`
         
    }
    
    // OverrideCheckoutResponse used by Cart
    type OverrideCheckoutResponse struct {

        
            Message string  `json:"message"`
            Data map[string]interface{}  `json:"data"`
            Success string  `json:"success"`
            OrderID string  `json:"order_id"`
            Cart map[string]interface{}  `json:"cart"`
         
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

        
            Source map[string]interface{}  `json:"source"`
            User map[string]interface{}  `json:"user"`
            Meta map[string]interface{}  `json:"meta"`
            CreatedOn string  `json:"created_on"`
            Token string  `json:"token"`
         
    }
    
    // SharedCart used by Cart
    type SharedCart struct {

        
            Message string  `json:"message"`
            CheckoutMode string  `json:"checkout_mode"`
            Items []CartProductInfo  `json:"items"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            CartID float64  `json:"cart_id"`
            Gstin string  `json:"gstin"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            Currency CartCurrency  `json:"currency"`
            LastModified string  `json:"last_modified"`
            IsValid bool  `json:"is_valid"`
            ID string  `json:"id"`
            BuyNow bool  `json:"buy_now"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            Comment string  `json:"comment"`
            CouponText string  `json:"coupon_text"`
            UID string  `json:"uid"`
            SharedCartDetails SharedCartDetails  `json:"shared_cart_details"`
         
    }
    
    // SharedCartResponse used by Cart
    type SharedCartResponse struct {

        
            Error string  `json:"error"`
            Cart SharedCart  `json:"cart"`
         
    }
    
    // CartList used by Cart
    type CartList struct {

        
            ItemCounts float64  `json:"item_counts"`
            CartID string  `json:"cart_id"`
            UserID string  `json:"user_id"`
            CreatedOn string  `json:"created_on"`
            CartValue float64  `json:"cart_value"`
            PickUpCustomerDetails map[string]interface{}  `json:"pick_up_customer_details"`
         
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

        
            FirstName string  `json:"first_name"`
            CreatedAt string  `json:"created_at"`
            LastName string  `json:"last_name"`
            ExternalID string  `json:"external_id"`
            Mobile string  `json:"mobile"`
            Gender string  `json:"gender"`
            ModifiedOn string  `json:"modified_on"`
            UID string  `json:"uid"`
            ID string  `json:"_id"`
         
    }
    
    // UserCartMappingResponse used by Cart
    type UserCartMappingResponse struct {

        
            Message string  `json:"message"`
            CheckoutMode string  `json:"checkout_mode"`
            Items []CartProductInfo  `json:"items"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            AppliedPromoDetails []AppliedPromotion  `json:"applied_promo_details"`
            Gstin string  `json:"gstin"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            Currency CartCurrency  `json:"currency"`
            PanNo string  `json:"pan_no"`
            LastModified string  `json:"last_modified"`
            IsValid bool  `json:"is_valid"`
            ID string  `json:"id"`
            PanConfig map[string]interface{}  `json:"pan_config"`
            BuyNow bool  `json:"buy_now"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            User UserInfo  `json:"user"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            Comment string  `json:"comment"`
            CouponText string  `json:"coupon_text"`
         
    }
    
    // PlatformAddCartRequest used by Cart
    type PlatformAddCartRequest struct {

        
            NewCart bool  `json:"new_cart"`
            Items []AddProductCart  `json:"items"`
            UserID string  `json:"user_id"`
         
    }
    
    // PlatformUpdateCartRequest used by Cart
    type PlatformUpdateCartRequest struct {

        
            Items []UpdateProductCart  `json:"items"`
            Operation string  `json:"operation"`
            UserID string  `json:"user_id"`
         
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

        
            Message string  `json:"message"`
            ExpiresOn string  `json:"expires_on"`
            MinimumCartValue float64  `json:"minimum_cart_value"`
            IsApplicable bool  `json:"is_applicable"`
            CouponValue float64  `json:"coupon_value"`
            MaxDiscountValue float64  `json:"max_discount_value"`
            CouponCode string  `json:"coupon_code"`
            CouponType string  `json:"coupon_type"`
            IsApplied bool  `json:"is_applied"`
            Title string  `json:"title"`
            Description string  `json:"description"`
            SubTitle string  `json:"sub_title"`
         
    }
    
    // PageCoupon used by Cart
    type PageCoupon struct {

        
            HasNext bool  `json:"has_next"`
            HasPrevious bool  `json:"has_previous"`
            Total float64  `json:"total"`
            Current float64  `json:"current"`
            TotalItemCount float64  `json:"total_item_count"`
         
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

        
            CheckoutMode string  `json:"checkout_mode"`
            AreaCode string  `json:"area_code"`
            Phone string  `json:"phone"`
            IsDefaultAddress bool  `json:"is_default_address"`
            AddressType string  `json:"address_type"`
            CartID string  `json:"cart_id"`
            Address string  `json:"address"`
            Area string  `json:"area"`
            Name string  `json:"name"`
            CreatedByUserID string  `json:"created_by_user_id"`
            Landmark string  `json:"landmark"`
            GoogleMapPoint map[string]interface{}  `json:"google_map_point"`
            State string  `json:"state"`
            City string  `json:"city"`
            UserID string  `json:"user_id"`
            CountryCode string  `json:"country_code"`
            ID string  `json:"id"`
            GeoLocation GeoLocation  `json:"geo_location"`
            Tags []string  `json:"tags"`
            IsActive bool  `json:"is_active"`
            AreaCodeSlug string  `json:"area_code_slug"`
            Email string  `json:"email"`
            Meta map[string]interface{}  `json:"meta"`
            Country string  `json:"country"`
         
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

        
            IsUpdated bool  `json:"is_updated"`
            IsDefaultAddress bool  `json:"is_default_address"`
            Success bool  `json:"success"`
            ID string  `json:"id"`
         
    }
    
    // DeleteAddressResponse used by Cart
    type DeleteAddressResponse struct {

        
            IsDeleted bool  `json:"is_deleted"`
            ID string  `json:"id"`
         
    }
    
    // PlatformSelectCartAddressRequest used by Cart
    type PlatformSelectCartAddressRequest struct {

        
            CheckoutMode string  `json:"checkout_mode"`
            CartID string  `json:"cart_id"`
            UserID string  `json:"user_id"`
            BillingAddressID string  `json:"billing_address_id"`
            ID string  `json:"id"`
         
    }
    
    // ShipmentArticle used by Cart
    type ShipmentArticle struct {

        
            Meta string  `json:"meta"`
            Quantity string  `json:"quantity"`
            ArticleID string  `json:"article_id"`
         
    }
    
    // PlatformShipmentResponse used by Cart
    type PlatformShipmentResponse struct {

        
            DpOptions map[string]interface{}  `json:"dp_options"`
            FulfillmentID float64  `json:"fulfillment_id"`
            Items []CartProductInfo  `json:"items"`
            ShipmentType string  `json:"shipment_type"`
            Articles []ShipmentArticle  `json:"articles"`
            OrderType string  `json:"order_type"`
            FulfillmentType string  `json:"fulfillment_type"`
            Shipments float64  `json:"shipments"`
            Promise ShipmentPromise  `json:"promise"`
            BoxType string  `json:"box_type"`
            DpID string  `json:"dp_id"`
         
    }
    
    // PlatformCartShipmentsResponse used by Cart
    type PlatformCartShipmentsResponse struct {

        
            Message string  `json:"message"`
            CheckoutMode string  `json:"checkout_mode"`
            Items []CartProductInfo  `json:"items"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            Error bool  `json:"error"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            AppliedPromoDetails []AppliedPromotion  `json:"applied_promo_details"`
            Gstin string  `json:"gstin"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            Currency CartCurrency  `json:"currency"`
            PanNo string  `json:"pan_no"`
            LastModified string  `json:"last_modified"`
            IsValid bool  `json:"is_valid"`
            ID string  `json:"id"`
            StaffUserID string  `json:"staff_user_id"`
            PanConfig map[string]interface{}  `json:"pan_config"`
            BuyNow bool  `json:"buy_now"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            Shipments []PlatformShipmentResponse  `json:"shipments"`
            Comment string  `json:"comment"`
            CouponText string  `json:"coupon_text"`
         
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

        
            CheckoutMode string  `json:"checkout_mode"`
            Gstin string  `json:"gstin"`
            PanNo string  `json:"pan_no"`
            GiftDetails map[string]interface{}  `json:"gift_details"`
            Comment string  `json:"comment"`
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

        
            FirstName string  `json:"first_name"`
            User string  `json:"user"`
            EmployeeCode string  `json:"employee_code"`
            LastName string  `json:"last_name"`
            ID string  `json:"_id"`
         
    }
    
    // Files used by Cart
    type Files struct {

        
            Values []string  `json:"values"`
            Key string  `json:"key"`
         
    }
    
    // PlatformCartCheckoutDetailRequest used by Cart
    type PlatformCartCheckoutDetailRequest struct {

        
            CallbackURL string  `json:"callback_url"`
            MerchantCode string  `json:"merchant_code"`
            CheckoutMode string  `json:"checkout_mode"`
            DeviceID string  `json:"device_id"`
            OrderType string  `json:"order_type"`
            Pos bool  `json:"pos"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
            Aggregator string  `json:"aggregator"`
            BillingAddressID string  `json:"billing_address_id"`
            PaymentAutoConfirm bool  `json:"payment_auto_confirm"`
            PaymentParams map[string]interface{}  `json:"payment_params"`
            PaymentMode string  `json:"payment_mode"`
            UserID string  `json:"user_id"`
            Staff StaffCheckout  `json:"staff"`
            PaymentIdentifier string  `json:"payment_identifier"`
            OrderingStore float64  `json:"ordering_store"`
            Files []Files  `json:"files"`
            ID string  `json:"id"`
            AddressID string  `json:"address_id"`
            EmployeeCode string  `json:"employee_code"`
            Meta map[string]interface{}  `json:"meta"`
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            PickAtStoreUID float64  `json:"pick_at_store_uid"`
         
    }
    
    // CheckCart used by Cart
    type CheckCart struct {

        
            Message string  `json:"message"`
            CheckoutMode string  `json:"checkout_mode"`
            Items []CartProductInfo  `json:"items"`
            StoreEmps []map[string]interface{}  `json:"store_emps"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            CodMessage string  `json:"cod_message"`
            ErrorMessage string  `json:"error_message"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            Success bool  `json:"success"`
            CartID float64  `json:"cart_id"`
            Gstin string  `json:"gstin"`
            CodCharges float64  `json:"cod_charges"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            DeliveryCharges float64  `json:"delivery_charges"`
            StoreCode string  `json:"store_code"`
            Currency CartCurrency  `json:"currency"`
            LastModified string  `json:"last_modified"`
            IsValid bool  `json:"is_valid"`
            ID string  `json:"id"`
            DeliveryChargeOrderValue float64  `json:"delivery_charge_order_value"`
            CodAvailable bool  `json:"cod_available"`
            OrderID string  `json:"order_id"`
            BuyNow bool  `json:"buy_now"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            Comment string  `json:"comment"`
            CouponText string  `json:"coupon_text"`
            UserType string  `json:"user_type"`
            UID string  `json:"uid"`
         
    }
    
    // CartCheckoutResponse used by Cart
    type CartCheckoutResponse struct {

        
            Message string  `json:"message"`
            Data map[string]interface{}  `json:"data"`
            CallbackURL string  `json:"callback_url"`
            Success bool  `json:"success"`
            AppInterceptURL string  `json:"app_intercept_url"`
            OrderID string  `json:"order_id"`
            PaymentConfirmURL string  `json:"payment_confirm_url"`
            Cart CheckCart  `json:"cart"`
         
    }
    
    // CartDeliveryModesResponse used by Cart
    type CartDeliveryModesResponse struct {

        
            AvailableModes []string  `json:"available_modes"`
            PickupStores []float64  `json:"pickup_stores"`
         
    }
    
    // PickupStoreDetail used by Cart
    type PickupStoreDetail struct {

        
            Pincode float64  `json:"pincode"`
            State string  `json:"state"`
            Address string  `json:"address"`
            City string  `json:"city"`
            AreaCodeSlug string  `json:"area_code_slug"`
            Email string  `json:"email"`
            Area string  `json:"area"`
            AreaCode string  `json:"area_code"`
            StoreCode string  `json:"store_code"`
            Name string  `json:"name"`
            Country string  `json:"country"`
            Landmark string  `json:"landmark"`
            Phone string  `json:"phone"`
            AddressType string  `json:"address_type"`
            UID float64  `json:"uid"`
            ID float64  `json:"id"`
         
    }
    
    // StoreDetailsResponse used by Cart
    type StoreDetailsResponse struct {

        
            Items []PickupStoreDetail  `json:"items"`
         
    }
    
    // UpdateCartPaymentRequest used by Cart
    type UpdateCartPaymentRequest struct {

        
            AddressID string  `json:"address_id"`
            MerchantCode string  `json:"merchant_code"`
            AggregatorName string  `json:"aggregator_name"`
            PaymentIdentifier string  `json:"payment_identifier"`
            PaymentMode string  `json:"payment_mode"`
            ID string  `json:"id"`
         
    }
    
    // CouponValidity used by Cart
    type CouponValidity struct {

        
            NextValidationRequired bool  `json:"next_validation_required"`
            DisplayMessageEn string  `json:"display_message_en"`
            Discount float64  `json:"discount"`
            Code string  `json:"code"`
            Valid bool  `json:"valid"`
            Title string  `json:"title"`
         
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
            PaymentIdentifier string  `json:"payment_identifier"`
            MerchantCode string  `json:"merchant_code"`
         
    }
    
    // PaymentMethod used by Cart
    type PaymentMethod struct {

        
            Amount float64  `json:"amount"`
            PaymentMeta PaymentMeta  `json:"payment_meta"`
            Mode string  `json:"mode"`
            Name string  `json:"name"`
            Payment string  `json:"payment"`
         
    }
    
    // PlatformCartCheckoutDetailV2Request used by Cart
    type PlatformCartCheckoutDetailV2Request struct {

        
            CallbackURL string  `json:"callback_url"`
            MerchantCode string  `json:"merchant_code"`
            CheckoutMode string  `json:"checkout_mode"`
            DeviceID string  `json:"device_id"`
            OrderType string  `json:"order_type"`
            Pos bool  `json:"pos"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
            Aggregator string  `json:"aggregator"`
            BillingAddressID string  `json:"billing_address_id"`
            PaymentMethods []PaymentMethod  `json:"payment_methods"`
            PaymentAutoConfirm bool  `json:"payment_auto_confirm"`
            PaymentParams map[string]interface{}  `json:"payment_params"`
            PaymentMode string  `json:"payment_mode"`
            CustomMeta map[string]interface{}  `json:"custom_meta"`
            UserID string  `json:"user_id"`
            Staff StaffCheckout  `json:"staff"`
            PaymentIdentifier string  `json:"payment_identifier"`
            OrderingStore float64  `json:"ordering_store"`
            Files []Files  `json:"files"`
            ID string  `json:"id"`
            AddressID string  `json:"address_id"`
            EmployeeCode string  `json:"employee_code"`
            Meta map[string]interface{}  `json:"meta"`
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            PickAtStoreUID float64  `json:"pick_at_store_uid"`
         
    }
    
    // UpdateCartPaymentRequestV2 used by Cart
    type UpdateCartPaymentRequestV2 struct {

        
            AddressID string  `json:"address_id"`
            MerchantCode string  `json:"merchant_code"`
            AggregatorName string  `json:"aggregator_name"`
            PaymentIdentifier string  `json:"payment_identifier"`
            PaymentMode string  `json:"payment_mode"`
            ID string  `json:"id"`
            PaymentMethods []PaymentMethod  `json:"payment_methods"`
         
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
    

    
    // ApplicationServiceabilityConfig used by Serviceability
    type ApplicationServiceabilityConfig struct {

        
            ChannelType string  `json:"channel_type"`
            ChannelID string  `json:"channel_id"`
            ServiceabilityType string  `json:"serviceability_type"`
         
    }
    
    // ServiceabilityErrorResponse used by Serviceability
    type ServiceabilityErrorResponse struct {

        
            Message string  `json:"message"`
            Value string  `json:"value"`
            Type string  `json:"type"`
         
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
    
    // EntityRegionView_Items used by Serviceability
    type EntityRegionView_Items struct {

        
            SubType string  `json:"sub_type"`
            UID string  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // EntityRegionView_page used by Serviceability
    type EntityRegionView_page struct {

        
            Current float64  `json:"current"`
            Size float64  `json:"size"`
            ItemTotal float64  `json:"item_total"`
            HasNext bool  `json:"has_next"`
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

        
            Data []EntityRegionView_Items  `json:"data"`
            Page EntityRegionView_page  `json:"page"`
            Success bool  `json:"success"`
            Error EntityRegionView_Error  `json:"error"`
         
    }
    
    // ZoneDataItem used by Serviceability
    type ZoneDataItem struct {

        
            Current float64  `json:"current"`
            Type string  `json:"type"`
            Size float64  `json:"size"`
            ItemTotal float64  `json:"item_total"`
            HasNext bool  `json:"has_next"`
         
    }
    
    // ListViewSummary used by Serviceability
    type ListViewSummary struct {

        
            TotalZones float64  `json:"total_zones"`
            TotalPincodesServed float64  `json:"total_pincodes_served"`
            TotalActiveZones float64  `json:"total_active_zones"`
         
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

        
            PincodesCount float64  `json:"pincodes_count"`
            ZoneID string  `json:"zone_id"`
            IsActive bool  `json:"is_active"`
            Product ListViewProduct  `json:"product"`
            CompanyID float64  `json:"company_id"`
            Name string  `json:"name"`
            Channels ListViewChannels  `json:"channels"`
            Slug string  `json:"slug"`
            StoresCount float64  `json:"stores_count"`
         
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
            Size float64  `json:"size"`
            ItemTotal float64  `json:"item_total"`
            HasNext bool  `json:"has_next"`
            Type string  `json:"type"`
         
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

        
            Data CreateZoneData  `json:"data"`
            Identifier string  `json:"identifier"`
         
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
            Size float64  `json:"size"`
            ItemTotal float64  `json:"item_total"`
            HasNext bool  `json:"has_next"`
            Type string  `json:"type"`
         
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
    
    // ModifiedByResponse used by Serviceability
    type ModifiedByResponse struct {

        
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
         
    }
    
    // Dp used by Serviceability
    type Dp struct {

        
            ExternalAccountID string  `json:"external_account_id"`
            AssignDpFromSb bool  `json:"assign_dp_from_sb"`
            RvpPriority float64  `json:"rvp_priority"`
            LmPriority float64  `json:"lm_priority"`
            InternalAccountID string  `json:"internal_account_id"`
            TransportMode string  `json:"transport_mode"`
            FmPriority float64  `json:"fm_priority"`
            Operations []string  `json:"operations"`
            PaymentMode string  `json:"payment_mode"`
            AreaCode float64  `json:"area_code"`
         
    }
    
    // LogisticsResponse used by Serviceability
    type LogisticsResponse struct {

        
            Override bool  `json:"override"`
            Dp Dp  `json:"dp"`
         
    }
    
    // WarningsResponse used by Serviceability
    type WarningsResponse struct {

        
            StoreAddress string  `json:"store_address"`
         
    }
    
    // AddressResponse used by Serviceability
    type AddressResponse struct {

        
            Latitude float64  `json:"latitude"`
            State string  `json:"state"`
            Longitude float64  `json:"longitude"`
            Pincode float64  `json:"pincode"`
            Address2 string  `json:"address2"`
            Address1 string  `json:"address1"`
            Country string  `json:"country"`
            City string  `json:"city"`
            Landmark string  `json:"landmark"`
         
    }
    
    // IntegrationTypeResponse used by Serviceability
    type IntegrationTypeResponse struct {

        
            Inventory string  `json:"inventory"`
            Order string  `json:"order"`
         
    }
    
    // DocumentsResponse used by Serviceability
    type DocumentsResponse struct {

        
            Value string  `json:"value"`
            LegalName string  `json:"legal_name"`
            Verified bool  `json:"verified"`
            Type string  `json:"type"`
         
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
    
    // OpeningClosing used by Serviceability
    type OpeningClosing struct {

        
            Minute float64  `json:"minute"`
            Hour float64  `json:"hour"`
         
    }
    
    // TimmingResponse used by Serviceability
    type TimmingResponse struct {

        
            Opening OpeningClosing  `json:"opening"`
            Weekday string  `json:"weekday"`
            Open bool  `json:"open"`
            Closing OpeningClosing  `json:"closing"`
         
    }
    
    // CreatedByResponse used by Serviceability
    type CreatedByResponse struct {

        
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
         
    }
    
    // ProductReturnConfigResponse used by Serviceability
    type ProductReturnConfigResponse struct {

        
            OnSameStore bool  `json:"on_same_store"`
         
    }
    
    // ContactNumberResponse used by Serviceability
    type ContactNumberResponse struct {

        
            Number string  `json:"number"`
            CountryCode float64  `json:"country_code"`
         
    }
    
    // ItemResponse used by Serviceability
    type ItemResponse struct {

        
            UID float64  `json:"uid"`
            CreatedOn string  `json:"created_on"`
            Code string  `json:"code"`
            StoreType string  `json:"store_type"`
            GstCredentials GstCredentialsResponse  `json:"gst_credentials"`
            Name string  `json:"name"`
            SubType string  `json:"sub_type"`
            VerifiedBy ModifiedByResponse  `json:"verified_by"`
            ModifiedOn string  `json:"modified_on"`
            Logistics LogisticsResponse  `json:"logistics"`
            DisplayName string  `json:"display_name"`
            Stage string  `json:"stage"`
            NotificationEmails []string  `json:"notification_emails"`
            Warnings WarningsResponse  `json:"warnings"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Address AddressResponse  `json:"address"`
            Company float64  `json:"company"`
            ModifiedBy ModifiedByResponse  `json:"modified_by"`
            CompanyID float64  `json:"company_id"`
            IntegrationType IntegrationTypeResponse  `json:"integration_type"`
            Documents []DocumentsResponse  `json:"documents"`
            Cls string  `json:"_cls"`
            Manager ManagerResponse  `json:"manager"`
            Timing []TimmingResponse  `json:"timing"`
            VerifiedOn string  `json:"verified_on"`
            CreatedBy CreatedByResponse  `json:"created_by"`
            ProductReturnConfig ProductReturnConfigResponse  `json:"product_return_config"`
            ContactNumbers []ContactNumberResponse  `json:"contact_numbers"`
         
    }
    
    // GetStoresViewResponse used by Serviceability
    type GetStoresViewResponse struct {

        
            Page ServiceabilityPageResponse  `json:"page"`
            Items []ItemResponse  `json:"items"`
         
    }
    
    // ReAssignStoreRequest used by Serviceability
    type ReAssignStoreRequest struct {

        
            Configuration map[string]interface{}  `json:"configuration"`
            IgnoredLocations []string  `json:"ignored_locations"`
            Articles []map[string]interface{}  `json:"articles"`
            ToPincode string  `json:"to_pincode"`
            Identifier string  `json:"identifier"`
         
    }
    
    // ReAssignStoreResponse used by Serviceability
    type ReAssignStoreResponse struct {

        
            Success bool  `json:"success"`
            Articles []map[string]interface{}  `json:"articles"`
            ToPincode string  `json:"to_pincode"`
            Error map[string]interface{}  `json:"error"`
         
    }
    
    // ApplicationCompanyDpViewResponse used by Serviceability
    type ApplicationCompanyDpViewResponse struct {

        
            ApplicationID string  `json:"application_id"`
            CompanyID float64  `json:"company_id"`
            Success bool  `json:"success"`
            CourierPartnerID float64  `json:"courier_partner_id"`
         
    }
    
    // ApplicationCompanyDpViewRequest used by Serviceability
    type ApplicationCompanyDpViewRequest struct {

        
            DpID string  `json:"dp_id"`
         
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
            Size float64  `json:"size"`
            ItemTotal float64  `json:"item_total"`
            HasNext bool  `json:"has_next"`
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
    
    // Dp1 used by Serviceability
    type Dp1 struct {

        
            AccountID string  `json:"account_id"`
            PlanRules map[string]interface{}  `json:"plan_rules"`
            DpID string  `json:"dp_id"`
            Name string  `json:"name"`
            PlanID string  `json:"plan_id"`
            Stage string  `json:"stage"`
            IsSelfShip bool  `json:"is_self_ship"`
         
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
            Value string  `json:"value"`
            Type string  `json:"type"`
         
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

        
            Page Page  `json:"page"`
            Success bool  `json:"success"`
            Items []Dp1  `json:"items"`
         
    }
    
    // DpSchemaInRuleListing used by Serviceability
    type DpSchemaInRuleListing struct {

        
            AccountID string  `json:"account_id"`
            PlanRules map[string]interface{}  `json:"plan_rules"`
            DpID string  `json:"dp_id"`
            Name string  `json:"name"`
            Priority float64  `json:"priority"`
            PlanID string  `json:"plan_id"`
            Stage string  `json:"stage"`
            IsSelfShip bool  `json:"is_self_ship"`
         
    }
    
    // DpRule used by Serviceability
    type DpRule struct {

        
            IsActive bool  `json:"is_active"`
            Conditions []map[string]interface{}  `json:"conditions"`
            DpIds map[string]DpSchemaInRuleListing  `json:"dp_ids"`
            CompanyID float64  `json:"company_id"`
            Name string  `json:"name"`
         
    }
    
    // DpRuleSuccessResponse used by Serviceability
    type DpRuleSuccessResponse struct {

        
            Data DpRule  `json:"data"`
            Success bool  `json:"success"`
            StatusCode float64  `json:"status_code"`
         
    }
    
    // FailureResponse used by Serviceability
    type FailureResponse struct {

        
            Success bool  `json:"success"`
            StatusCode float64  `json:"status_code"`
            Error []ErrorResponse  `json:"error"`
         
    }
    
    // DpRulesUpdateRequest used by Serviceability
    type DpRulesUpdateRequest struct {

        
            IsActive bool  `json:"is_active"`
            Conditions []map[string]interface{}  `json:"conditions"`
            Name string  `json:"name"`
            DpIds map[string]map[string]interface{}  `json:"dp_ids"`
         
    }
    
    // DpRuleResponse used by Serviceability
    type DpRuleResponse struct {

        
            UID string  `json:"uid"`
            CreatedOn string  `json:"created_on"`
            IsActive bool  `json:"is_active"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Conditions []string  `json:"conditions"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            DpIds map[string]interface{}  `json:"dp_ids"`
            ModifiedOn string  `json:"modified_on"`
            CompanyID float64  `json:"company_id"`
            Name string  `json:"name"`
         
    }
    
    // DpRuleUpdateSuccessResponse used by Serviceability
    type DpRuleUpdateSuccessResponse struct {

        
            Data DpRuleResponse  `json:"data"`
            Success bool  `json:"success"`
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
            Conditions []map[string]interface{}  `json:"conditions"`
            DpIds map[string]DpIds  `json:"dp_ids"`
            CompanyID float64  `json:"company_id"`
            Name string  `json:"name"`
         
    }
    
    // DpMultipleRuleSuccessResponse used by Serviceability
    type DpMultipleRuleSuccessResponse struct {

        
            Page Page  `json:"page"`
            Success bool  `json:"success"`
            Items []DpRule  `json:"items"`
         
    }
    
    // DPCompanyRuleResponse used by Serviceability
    type DPCompanyRuleResponse struct {

        
            Data []DpRuleResponse  `json:"data"`
            Success bool  `json:"success"`
            StatusCode float64  `json:"status_code"`
         
    }
    
    // DPCompanyRuleRequest used by Serviceability
    type DPCompanyRuleRequest struct {

        
            RuleIds []string  `json:"rule_ids"`
         
    }
    
    // DPApplicationRuleResponse used by Serviceability
    type DPApplicationRuleResponse struct {

        
            Data []DpRuleResponse  `json:"data"`
            Success bool  `json:"success"`
            StatusCode bool  `json:"status_code"`
         
    }
    
    // DPApplicationRuleRequest used by Serviceability
    type DPApplicationRuleRequest struct {

        
            ShippingRules []string  `json:"shipping_rules"`
         
    }
    
    // SelfShipResponse used by Serviceability
    type SelfShipResponse struct {

        
            IsActive bool  `json:"is_active"`
            Tat float64  `json:"tat"`
         
    }
    
    // ApplicationSelfShipConfig used by Serviceability
    type ApplicationSelfShipConfig struct {

        
            SelfShip SelfShipResponse  `json:"self_ship"`
         
    }
    
    // ApplicationSelfShipConfigResponse used by Serviceability
    type ApplicationSelfShipConfigResponse struct {

        
            Data ApplicationSelfShipConfig  `json:"data"`
            Success bool  `json:"success"`
            Error ServiceabilityErrorResponse  `json:"error"`
         
    }
    
