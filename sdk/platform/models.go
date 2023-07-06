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

        
            Success bool  `json:"success"`
            Aggregators []map[string]interface{}  `json:"aggregators"`
            Created bool  `json:"created"`
            AppID string  `json:"app_id"`
            ExcludedFields []string  `json:"excluded_fields"`
            DisplayFields []string  `json:"display_fields"`
         
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
            MerchantSalt string  `json:"merchant_salt"`
            Key string  `json:"key"`
            Secret string  `json:"secret"`
         
    }
    
    // PaymentGatewayConfigRequest used by Payment
    type PaymentGatewayConfigRequest struct {

        
            AppID string  `json:"app_id"`
            IsActive bool  `json:"is_active"`
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

        
            PackageName string  `json:"package_name"`
            Code string  `json:"code"`
         
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
    
    // PaymentModeList used by Payment
    type PaymentModeList struct {

        
            CardNumber string  `json:"card_number"`
            IntentAppErrorDictList []IntentAppErrorList  `json:"intent_app_error_dict_list"`
            RemainingLimit float64  `json:"remaining_limit"`
            Expired bool  `json:"expired"`
            AggregatorName string  `json:"aggregator_name"`
            ExpMonth float64  `json:"exp_month"`
            CardIsin string  `json:"card_isin"`
            FyndVpa string  `json:"fynd_vpa"`
            CodLimit float64  `json:"cod_limit"`
            CardBrand string  `json:"card_brand"`
            CardType string  `json:"card_type"`
            CardName string  `json:"card_name"`
            Name string  `json:"name"`
            CardFingerprint string  `json:"card_fingerprint"`
            CodLimitPerOrder float64  `json:"cod_limit_per_order"`
            IntentFlow bool  `json:"intent_flow"`
            Timeout float64  `json:"timeout"`
            CompliantWithTokenisationGuidelines bool  `json:"compliant_with_tokenisation_guidelines"`
            IntentAppErrorList []string  `json:"intent_app_error_list"`
            IntentApp []IntentApp  `json:"intent_app"`
            CardBrandImage string  `json:"card_brand_image"`
            CardToken string  `json:"card_token"`
            ExpYear float64  `json:"exp_year"`
            LogoURL PaymentModeLogo  `json:"logo_url"`
            Code string  `json:"code"`
            Nickname string  `json:"nickname"`
            MerchantCode string  `json:"merchant_code"`
            DisplayName string  `json:"display_name"`
            CardID string  `json:"card_id"`
            RetryCount float64  `json:"retry_count"`
            CardReference string  `json:"card_reference"`
            DisplayPriority float64  `json:"display_priority"`
            CardIssuer string  `json:"card_issuer"`
         
    }
    
    // RootPaymentMode used by Payment
    type RootPaymentMode struct {

        
            List []PaymentModeList  `json:"list"`
            SaveCard bool  `json:"save_card"`
            AnonymousEnable bool  `json:"anonymous_enable"`
            DisplayName string  `json:"display_name"`
            AddCardEnabled bool  `json:"add_card_enabled"`
            Name string  `json:"name"`
            DisplayPriority float64  `json:"display_priority"`
            IsPayByCardPl bool  `json:"is_pay_by_card_pl"`
            AggregatorName string  `json:"aggregator_name"`
         
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

        
            PayoutsAggregators []map[string]interface{}  `json:"payouts_aggregators"`
            UniqueTransferNo map[string]interface{}  `json:"unique_transfer_no"`
            IsActive bool  `json:"is_active"`
            MoreAttributes map[string]interface{}  `json:"more_attributes"`
            IsDefault bool  `json:"is_default"`
            Customers map[string]interface{}  `json:"customers"`
            TransferType string  `json:"transfer_type"`
         
    }
    
    // PayoutBankDetails used by Payment
    type PayoutBankDetails struct {

        
            AccountNo string  `json:"account_no"`
            State string  `json:"state"`
            Country string  `json:"country"`
            BankName string  `json:"bank_name"`
            AccountHolder string  `json:"account_holder"`
            IfscCode string  `json:"ifsc_code"`
            AccountType string  `json:"account_type"`
            Pincode float64  `json:"pincode"`
            City string  `json:"city"`
            BranchName string  `json:"branch_name"`
         
    }
    
    // PayoutRequest used by Payment
    type PayoutRequest struct {

        
            Users map[string]interface{}  `json:"users"`
            BankDetails PayoutBankDetails  `json:"bank_details"`
            IsActive bool  `json:"is_active"`
            TransferType string  `json:"transfer_type"`
            Aggregator string  `json:"aggregator"`
            UniqueExternalID string  `json:"unique_external_id"`
         
    }
    
    // PayoutResponse used by Payment
    type PayoutResponse struct {

        
            Success bool  `json:"success"`
            Users map[string]interface{}  `json:"users"`
            BankDetails map[string]interface{}  `json:"bank_details"`
            UniqueTransferNo string  `json:"unique_transfer_no"`
            Created bool  `json:"created"`
            IsActive bool  `json:"is_active"`
            Payouts map[string]interface{}  `json:"payouts"`
            TransferType string  `json:"transfer_type"`
            Aggregator string  `json:"aggregator"`
            PaymentStatus string  `json:"payment_status"`
         
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
            Config map[string]interface{}  `json:"config"`
            Aggregator string  `json:"aggregator"`
         
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
            Message string  `json:"message"`
            Data map[string]interface{}  `json:"data"`
            IsVerifiedFlag bool  `json:"is_verified_flag"`
         
    }
    
    // NotFoundResourceError used by Payment
    type NotFoundResourceError struct {

        
            Success bool  `json:"success"`
            Description string  `json:"description"`
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

        
            Success bool  `json:"success"`
            BranchName string  `json:"branch_name"`
            BankName string  `json:"bank_name"`
         
    }
    
    // OrderBeneficiaryDetails used by Payment
    type OrderBeneficiaryDetails struct {

        
            AccountNo string  `json:"account_no"`
            Mobile string  `json:"mobile"`
            CreatedOn string  `json:"created_on"`
            BeneficiaryID string  `json:"beneficiary_id"`
            TransferMode string  `json:"transfer_mode"`
            BankName string  `json:"bank_name"`
            AccountHolder string  `json:"account_holder"`
            ID float64  `json:"id"`
            BranchName string  `json:"branch_name"`
            Comment string  `json:"comment"`
            Subtitle string  `json:"subtitle"`
            Email string  `json:"email"`
            Title string  `json:"title"`
            IsActive bool  `json:"is_active"`
            DelightsUserName string  `json:"delights_user_name"`
            ModifiedOn string  `json:"modified_on"`
            Address string  `json:"address"`
            IfscCode string  `json:"ifsc_code"`
            DisplayName string  `json:"display_name"`
         
    }
    
    // OrderBeneficiaryResponse used by Payment
    type OrderBeneficiaryResponse struct {

        
            ShowBeneficiaryDetails bool  `json:"show_beneficiary_details"`
            Beneficiaries []OrderBeneficiaryDetails  `json:"beneficiaries"`
         
    }
    
    // MultiTenderPaymentMeta used by Payment
    type MultiTenderPaymentMeta struct {

        
            CurrentStatus string  `json:"current_status"`
            PaymentGateway string  `json:"payment_gateway"`
            PaymentID string  `json:"payment_id"`
            OrderID string  `json:"order_id"`
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

        
            PaymentMethods []MultiTenderPaymentMethod  `json:"payment_methods"`
            OrderID string  `json:"order_id"`
         
    }
    
    // PaymentConfirmationResponse used by Payment
    type PaymentConfirmationResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
            OrderID string  `json:"order_id"`
         
    }
    
    // CODdata used by Payment
    type CODdata struct {

        
            UserID string  `json:"user_id"`
            Usages float64  `json:"usages"`
            Limit float64  `json:"limit"`
            IsActive bool  `json:"is_active"`
            RemainingLimit float64  `json:"remaining_limit"`
         
    }
    
    // GetUserCODLimitResponse used by Payment
    type GetUserCODLimitResponse struct {

        
            Success bool  `json:"success"`
            UserCodData CODdata  `json:"user_cod_data"`
         
    }
    
    // SetCODForUserRequest used by Payment
    type SetCODForUserRequest struct {

        
            IsActive bool  `json:"is_active"`
            Mobileno string  `json:"mobileno"`
            MerchantUserID string  `json:"merchant_user_id"`
         
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

        
            ActiveDeviceCount float64  `json:"active_device_count"`
            InactiveDeviceCount float64  `json:"inactive_device_count"`
         
    }
    
    // EdcDeviceStatsResponse used by Payment
    type EdcDeviceStatsResponse struct {

        
            Success bool  `json:"success"`
            Statistics StatisticsData  `json:"statistics"`
         
    }
    
    // EdcAddRequest used by Payment
    type EdcAddRequest struct {

        
            DeviceTag string  `json:"device_tag"`
            EdcModel string  `json:"edc_model"`
            StoreID float64  `json:"store_id"`
            TerminalSerialNo string  `json:"terminal_serial_no"`
            AggregatorID float64  `json:"aggregator_id"`
            EdcDeviceSerialNo string  `json:"edc_device_serial_no"`
         
    }
    
    // EdcDevice used by Payment
    type EdcDevice struct {

        
            TerminalUniqueIdentifier string  `json:"terminal_unique_identifier"`
            DeviceTag string  `json:"device_tag"`
            ApplicationID string  `json:"application_id"`
            AggregatorName string  `json:"aggregator_name"`
            IsActive bool  `json:"is_active"`
            EdcModel string  `json:"edc_model"`
            StoreID float64  `json:"store_id"`
            TerminalSerialNo string  `json:"terminal_serial_no"`
            MerchantStorePosCode string  `json:"merchant_store_pos_code"`
            AggregatorID float64  `json:"aggregator_id"`
            EdcDeviceSerialNo string  `json:"edc_device_serial_no"`
         
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

        
            DeviceTag string  `json:"device_tag"`
            IsActive bool  `json:"is_active"`
            EdcModel string  `json:"edc_model"`
            StoreID float64  `json:"store_id"`
            MerchantStorePosCode string  `json:"merchant_store_pos_code"`
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

        
            Success bool  `json:"success"`
            Page Page  `json:"page"`
            Items []EdcDevice  `json:"items"`
         
    }
    
    // PaymentInitializationRequest used by Payment
    type PaymentInitializationRequest struct {

        
            Email string  `json:"email"`
            Timeout float64  `json:"timeout"`
            Currency string  `json:"currency"`
            DeviceID string  `json:"device_id"`
            Method string  `json:"method"`
            Vpa string  `json:"vpa"`
            Aggregator string  `json:"aggregator"`
            RazorpayPaymentID string  `json:"razorpay_payment_id"`
            MerchantOrderID string  `json:"merchant_order_id"`
            Contact string  `json:"contact"`
            CustomerID string  `json:"customer_id"`
            OrderID string  `json:"order_id"`
            Amount float64  `json:"amount"`
         
    }
    
    // PaymentInitializationResponse used by Payment
    type PaymentInitializationResponse struct {

        
            Success bool  `json:"success"`
            VirtualID string  `json:"virtual_id"`
            Status string  `json:"status"`
            BqrImage string  `json:"bqr_image"`
            Timeout float64  `json:"timeout"`
            Currency string  `json:"currency"`
            DeviceID string  `json:"device_id"`
            Method string  `json:"method"`
            Vpa string  `json:"vpa"`
            RazorpayPaymentID string  `json:"razorpay_payment_id"`
            UpiPollURL string  `json:"upi_poll_url"`
            MerchantOrderID string  `json:"merchant_order_id"`
            Aggregator string  `json:"aggregator"`
            CustomerID string  `json:"customer_id"`
            PollingURL string  `json:"polling_url"`
            AggregatorOrderID string  `json:"aggregator_order_id"`
            Amount float64  `json:"amount"`
         
    }
    
    // PaymentStatusUpdateRequest used by Payment
    type PaymentStatusUpdateRequest struct {

        
            Status string  `json:"status"`
            MerchantTransactionID string  `json:"merchant_transaction_id"`
            Email string  `json:"email"`
            Currency string  `json:"currency"`
            DeviceID string  `json:"device_id"`
            Method string  `json:"method"`
            Vpa string  `json:"vpa"`
            Aggregator string  `json:"aggregator"`
            MerchantOrderID string  `json:"merchant_order_id"`
            Contact string  `json:"contact"`
            CustomerID string  `json:"customer_id"`
            OrderID string  `json:"order_id"`
            Amount float64  `json:"amount"`
         
    }
    
    // PaymentStatusUpdateResponse used by Payment
    type PaymentStatusUpdateResponse struct {

        
            Success bool  `json:"success"`
            Status string  `json:"status"`
            RedirectURL string  `json:"redirect_url"`
            Retry bool  `json:"retry"`
            AggregatorName string  `json:"aggregator_name"`
         
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

        
            Success bool  `json:"success"`
            Data LinkStatus  `json:"data"`
         
    }
    
    // PaymentStatusBulkHandlerRequest used by Payment
    type PaymentStatusBulkHandlerRequest struct {

        
            MerchantOrderID []string  `json:"merchant_order_id"`
         
    }
    
    // PaymentObjectListSerializer used by Payment
    type PaymentObjectListSerializer struct {

        
            ModifiedOn string  `json:"modified_on"`
            AggregatorPaymentObject map[string]interface{}  `json:"aggregator_payment_object"`
            ApplicationID string  `json:"application_id"`
            AllStatus []string  `json:"all_status"`
            RefundedBy string  `json:"refunded_by"`
            RefundObject map[string]interface{}  `json:"refund_object"`
            Currency string  `json:"currency"`
            AmountInPaisa string  `json:"amount_in_paisa"`
            CurrentStatus string  `json:"current_status"`
            CollectedBy string  `json:"collected_by"`
            UserObject map[string]interface{}  `json:"user_object"`
            ID string  `json:"id"`
            PaymentModeIdentifier string  `json:"payment_mode_identifier"`
            CompanyID string  `json:"company_id"`
            PaymentGateway string  `json:"payment_gateway"`
            PaymentMode string  `json:"payment_mode"`
            CreatedOn string  `json:"created_on"`
            PaymentID string  `json:"payment_id"`
         
    }
    
    // PaymentStatusObject used by Payment
    type PaymentStatusObject struct {

        
            PaymentObjectList []PaymentObjectListSerializer  `json:"payment_object_list"`
            MerchantOrderID string  `json:"merchant_order_id"`
         
    }
    
    // PaymentStatusBulkHandlerResponse used by Payment
    type PaymentStatusBulkHandlerResponse struct {

        
            Success string  `json:"success"`
            Status float64  `json:"status"`
            Error string  `json:"error"`
            Count float64  `json:"count"`
            Data []PaymentStatusObject  `json:"data"`
         
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

        
            OutstandingDetailsID float64  `json:"outstanding_details_id"`
            FwdShipmentID string  `json:"fwd_shipment_id"`
            CurrentStatus string  `json:"current_status"`
            PaymentModeIdentifier string  `json:"payment_mode_identifier"`
            AggregatorTransactionID string  `json:"aggregator_transaction_id"`
            Aggregator string  `json:"aggregator"`
            MerchantOrderID string  `json:"merchant_order_id"`
            PaymentMode string  `json:"payment_mode"`
            AggregatorOrderID string  `json:"aggregator_order_id"`
            Amount float64  `json:"amount"`
         
    }
    
    // RepaymentDetailsSerialiserPayAll used by Payment
    type RepaymentDetailsSerialiserPayAll struct {

        
            ShipmentDetails []RepaymentRequestDetails  `json:"shipment_details"`
            ExtensionOrderID string  `json:"extension_order_id"`
            AggregatorTransactionID string  `json:"aggregator_transaction_id"`
            TotalAmount float64  `json:"total_amount"`
            AggregatorOrderID string  `json:"aggregator_order_id"`
         
    }
    
    // RepaymentResponse used by Payment
    type RepaymentResponse struct {

        
            Success bool  `json:"success"`
            Data map[string]interface{}  `json:"data"`
         
    }
    
    // MerchantOnBoardingRequest used by Payment
    type MerchantOnBoardingRequest struct {

        
            UserID string  `json:"user_id"`
            Status string  `json:"status"`
            AppID string  `json:"app_id"`
            Aggregator string  `json:"aggregator"`
            CreditLineID string  `json:"credit_line_id"`
         
    }
    
    // MerchantOnBoardingResponse used by Payment
    type MerchantOnBoardingResponse struct {

        
            Success bool  `json:"success"`
            Data map[string]interface{}  `json:"data"`
         
    }
    
    // ValidateCustomerRequest used by Payment
    type ValidateCustomerRequest struct {

        
            TransactionAmountInPaise float64  `json:"transaction_amount_in_paise"`
            PhoneNumber string  `json:"phone_number"`
            Payload string  `json:"payload"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
            Aggregator string  `json:"aggregator"`
            OrderItems []map[string]interface{}  `json:"order_items"`
            MerchantParams map[string]interface{}  `json:"merchant_params"`
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
         
    }
    
    // ValidateCustomerResponse used by Payment
    type ValidateCustomerResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
            Data map[string]interface{}  `json:"data"`
         
    }
    
    // GetPaymentLinkResponse used by Payment
    type GetPaymentLinkResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
            ExternalOrderID string  `json:"external_order_id"`
            PaymentLinkURL string  `json:"payment_link_url"`
            PaymentLinkCurrentStatus string  `json:"payment_link_current_status"`
            MerchantName string  `json:"merchant_name"`
            StatusCode float64  `json:"status_code"`
            PollingTimeout float64  `json:"polling_timeout"`
            Amount float64  `json:"amount"`
         
    }
    
    // ErrorDescription used by Payment
    type ErrorDescription struct {

        
            InvalidID bool  `json:"invalid_id"`
            Cancelled bool  `json:"cancelled"`
            Msg string  `json:"msg"`
            MerchantName string  `json:"merchant_name"`
            MerchantOrderID string  `json:"merchant_order_id"`
            Expired bool  `json:"expired"`
            PaymentTransactionID string  `json:"payment_transaction_id"`
            Amount float64  `json:"amount"`
         
    }
    
    // ErrorResponse used by Payment
    type ErrorResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
            StatusCode float64  `json:"status_code"`
            Error ErrorDescription  `json:"error"`
         
    }
    
    // CreatePaymentLinkMeta used by Payment
    type CreatePaymentLinkMeta struct {

        
            CartID string  `json:"cart_id"`
            CheckoutMode string  `json:"checkout_mode"`
            Pincode string  `json:"pincode"`
            AssignCardID string  `json:"assign_card_id"`
            Amount string  `json:"amount"`
         
    }
    
    // CreatePaymentLinkRequest used by Payment
    type CreatePaymentLinkRequest struct {

        
            ExternalOrderID string  `json:"external_order_id"`
            Email string  `json:"email"`
            Description string  `json:"description"`
            Meta CreatePaymentLinkMeta  `json:"meta"`
            Amount float64  `json:"amount"`
            MobileNumber string  `json:"mobile_number"`
         
    }
    
    // CreatePaymentLinkResponse used by Payment
    type CreatePaymentLinkResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
            PaymentLinkID string  `json:"payment_link_id"`
            PaymentLinkURL string  `json:"payment_link_url"`
            StatusCode float64  `json:"status_code"`
            PollingTimeout float64  `json:"polling_timeout"`
         
    }
    
    // PollingPaymentLinkResponse used by Payment
    type PollingPaymentLinkResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
            Status string  `json:"status"`
            RedirectURL string  `json:"redirect_url"`
            PaymentLinkID string  `json:"payment_link_id"`
            HttpStatus float64  `json:"http_status"`
            StatusCode float64  `json:"status_code"`
            OrderID string  `json:"order_id"`
            AggregatorName string  `json:"aggregator_name"`
            Amount float64  `json:"amount"`
         
    }
    
    // CancelOrResendPaymentLinkRequest used by Payment
    type CancelOrResendPaymentLinkRequest struct {

        
            PaymentLinkID string  `json:"payment_link_id"`
         
    }
    
    // ResendPaymentLinkResponse used by Payment
    type ResendPaymentLinkResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
            StatusCode float64  `json:"status_code"`
            PollingTimeout float64  `json:"polling_timeout"`
         
    }
    
    // CancelPaymentLinkResponse used by Payment
    type CancelPaymentLinkResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
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

        
            Networks string  `json:"networks"`
            Codes Code  `json:"codes"`
            Types string  `json:"types"`
            Name string  `json:"name"`
         
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
    

    
    // ShipmentStatus used by Order
    type ShipmentStatus struct {

        
            CurrentShipmentStatus string  `json:"current_shipment_status"`
            BagList []string  `json:"bag_list"`
            ShipmentID string  `json:"shipment_id"`
            Meta map[string]interface{}  `json:"meta"`
            Title string  `json:"title"`
            ShipmentStatusID float64  `json:"shipment_status_id"`
            Status string  `json:"status"`
            CreatedAt string  `json:"created_at"`
            StatusCreatedAt string  `json:"status_created_at"`
         
    }
    
    // LockData used by Order
    type LockData struct {

        
            LockMessage string  `json:"lock_message"`
            Locked bool  `json:"locked"`
            Mto bool  `json:"mto"`
         
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

        
            Slug string  `json:"slug"`
            DisplayText string  `json:"display_text"`
            EntityType string  `json:"entity_type"`
         
    }
    
    // ShipmentItemMeta used by Order
    type ShipmentItemMeta struct {

        
            IsInternational bool  `json:"is_international"`
            ShipmentVolumetricWeight float64  `json:"shipment_volumetric_weight"`
            Weight float64  `json:"weight"`
            ShipmentWeight float64  `json:"shipment_weight"`
            StoreInvoiceUpdatedDate string  `json:"store_invoice_updated_date"`
            DpOptions map[string]interface{}  `json:"dp_options"`
            LockData LockData  `json:"lock_data"`
            FulfilmentPriorityText string  `json:"fulfilment_priority_text"`
            SameStoreAvailable bool  `json:"same_store_available"`
            EwaybillInfo map[string]interface{}  `json:"ewaybill_info"`
            DebugInfo map[string]interface{}  `json:"debug_info"`
            OrderType string  `json:"order_type"`
            ActivityComment string  `json:"activity_comment"`
            ShippingZone string  `json:"shipping_zone"`
            DpSortKey string  `json:"dp_sort_key"`
            Sla float64  `json:"sla"`
            ExistingDpList []string  `json:"existing_dp_list"`
            AssignDpFromSb bool  `json:"assign_dp_from_sb"`
            ShipmentChargeableWeight float64  `json:"shipment_chargeable_weight"`
            PackagingName string  `json:"packaging_name"`
            Tags []map[string]interface{}  `json:"tags"`
            External map[string]interface{}  `json:"external"`
            ParentDpID string  `json:"parent_dp_id"`
            AutoTriggerDpAssignmentAcf bool  `json:"auto_trigger_dp_assignment_acf"`
            BagWeight map[string]interface{}  `json:"bag_weight"`
            Formatted Formatted  `json:"formatted"`
            Timestamp ShipmentTimeStamp  `json:"timestamp"`
            ShipmentTags []ShipmentTags  `json:"shipment_tags"`
            PdfMedia []map[string]interface{}  `json:"pdf_media"`
         
    }
    
    // BagStateMapper used by Order
    type BagStateMapper struct {

        
            StateType string  `json:"state_type"`
            IsActive bool  `json:"is_active"`
            AppStateName string  `json:"app_state_name"`
            AppFacing bool  `json:"app_facing"`
            ID float64  `json:"id"`
            DisplayName string  `json:"display_name"`
            JourneyType string  `json:"journey_type"`
            Name string  `json:"name"`
            AppDisplayName string  `json:"app_display_name"`
            NotifyCustomer bool  `json:"notify_customer"`
         
    }
    
    // BagStatusHistory used by Order
    type BagStatusHistory struct {

        
            UpdatedAt string  `json:"updated_at"`
            StateType string  `json:"state_type"`
            StateID float64  `json:"state_id"`
            BshID float64  `json:"bsh_id"`
            ShipmentID string  `json:"shipment_id"`
            DeliveryAwbNumber string  `json:"delivery_awb_number"`
            DisplayName string  `json:"display_name"`
            Status string  `json:"status"`
            KafkaSync bool  `json:"kafka_sync"`
            Forward bool  `json:"forward"`
            Reasons []map[string]interface{}  `json:"reasons"`
            CreatedAt string  `json:"created_at"`
            BagStateMapper BagStateMapper  `json:"bag_state_mapper"`
            DeliveryPartnerID float64  `json:"delivery_partner_id"`
            AppDisplayName string  `json:"app_display_name"`
            StoreID float64  `json:"store_id"`
            BagID float64  `json:"bag_id"`
         
    }
    
    // Prices used by Order
    type Prices struct {

        
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            Discount float64  `json:"discount"`
            RefundAmount float64  `json:"refund_amount"`
            PriceMarked float64  `json:"price_marked"`
            CouponValue float64  `json:"coupon_value"`
            Cashback float64  `json:"cashback"`
            RefundCredit float64  `json:"refund_credit"`
            PromotionEffectiveDiscount float64  `json:"promotion_effective_discount"`
            ValueOfGood float64  `json:"value_of_good"`
            AmountPaid float64  `json:"amount_paid"`
            GiftPrice float64  `json:"gift_price"`
            FyndCredits float64  `json:"fynd_credits"`
            AmountPaidRoundoff float64  `json:"amount_paid_roundoff"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            CashbackApplied float64  `json:"cashback_applied"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            PmPriceSplit float64  `json:"pm_price_split"`
            DeliveryCharge float64  `json:"delivery_charge"`
            CodCharges float64  `json:"cod_charges"`
            TransferPrice float64  `json:"transfer_price"`
            PriceEffective float64  `json:"price_effective"`
         
    }
    
    // Dates used by Order
    type Dates struct {

        
            OrderCreated string  `json:"order_created"`
            DeliveryDate string  `json:"delivery_date"`
         
    }
    
    // GSTDetailsData used by Order
    type GSTDetailsData struct {

        
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            IsDefaultHsnCode bool  `json:"is_default_hsn_code"`
            CgstGstFee string  `json:"cgst_gst_fee"`
            GstinCode string  `json:"gstin_code"`
            HsnCodeID string  `json:"hsn_code_id"`
            IgstTaxPercentage float64  `json:"igst_tax_percentage"`
            IgstGstFee string  `json:"igst_gst_fee"`
            SgstGstFee string  `json:"sgst_gst_fee"`
            HsnCode string  `json:"hsn_code"`
            GstFee float64  `json:"gst_fee"`
            CgstTaxPercentage float64  `json:"cgst_tax_percentage"`
            SgstTaxPercentage float64  `json:"sgst_tax_percentage"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            ValueOfGood float64  `json:"value_of_good"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            GstTag string  `json:"gst_tag"`
         
    }
    
    // BagReturnableCancelableStatus used by Order
    type BagReturnableCancelableStatus struct {

        
            IsActive bool  `json:"is_active"`
            EnableTracking bool  `json:"enable_tracking"`
            IsCustomerReturnAllowed bool  `json:"is_customer_return_allowed"`
            CanBeCancelled bool  `json:"can_be_cancelled"`
            IsReturnable bool  `json:"is_returnable"`
         
    }
    
    // ReplacementDetails used by Order
    type ReplacementDetails struct {

        
            OriginalAffiliateOrderID string  `json:"original_affiliate_order_id"`
            ReplacementType string  `json:"replacement_type"`
         
    }
    
    // AffiliateMeta used by Order
    type AffiliateMeta struct {

        
            ReplacementDetails ReplacementDetails  `json:"replacement_details"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            EmployeeDiscount float64  `json:"employee_discount"`
            ChannelShipmentID string  `json:"channel_shipment_id"`
            MarketplaceInvoiceID string  `json:"marketplace_invoice_id"`
            DueDate string  `json:"due_date"`
            SizeLevelTotalQty float64  `json:"size_level_total_qty"`
            Quantity float64  `json:"quantity"`
            ChannelOrderID string  `json:"channel_order_id"`
            IsPriority bool  `json:"is_priority"`
            OrderItemID string  `json:"order_item_id"`
            CouponCode string  `json:"coupon_code"`
            BoxType string  `json:"box_type"`
         
    }
    
    // AffiliateBagDetails used by Order
    type AffiliateBagDetails struct {

        
            AffiliateOrderID string  `json:"affiliate_order_id"`
            EmployeeDiscount float64  `json:"employee_discount"`
            AffiliateMeta AffiliateMeta  `json:"affiliate_meta"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
         
    }
    
    // PlatformArticleAttributes used by Order
    type PlatformArticleAttributes struct {

        
            Currency string  `json:"currency"`
         
    }
    
    // PlatformItem used by Order
    type PlatformItem struct {

        
            Image []string  `json:"image"`
            BrandID float64  `json:"brand_id"`
            L3Category float64  `json:"l3_category"`
            L1Category []string  `json:"l1_category"`
            Meta map[string]interface{}  `json:"meta"`
            SlugKey string  `json:"slug_key"`
            Attributes PlatformArticleAttributes  `json:"attributes"`
            L3CategoryName string  `json:"l3_category_name"`
            CanReturn bool  `json:"can_return"`
            ID float64  `json:"id"`
            CanCancel bool  `json:"can_cancel"`
            L2Category []string  `json:"l2_category"`
            BranchURL string  `json:"branch_url"`
            Size string  `json:"size"`
            LastUpdatedAt string  `json:"last_updated_at"`
            Color string  `json:"color"`
            Name string  `json:"name"`
            Code string  `json:"code"`
            Images []string  `json:"images"`
            Brand string  `json:"brand"`
            DepartmentID float64  `json:"department_id"`
         
    }
    
    // ShipmentListingBrand used by Order
    type ShipmentListingBrand struct {

        
            CreatedOn string  `json:"created_on"`
            LogoBase64 string  `json:"logo_base64"`
            Logo string  `json:"logo"`
            Name string  `json:"name"`
         
    }
    
    // Identifier used by Order
    type Identifier struct {

        
            Ean string  `json:"ean"`
            Alu string  `json:"alu"`
            SkuCode string  `json:"sku_code"`
            Upc string  `json:"upc"`
            Isbn string  `json:"isbn"`
         
    }
    
    // FinancialBreakup used by Order
    type FinancialBreakup struct {

        
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            Discount float64  `json:"discount"`
            PriceMarked float64  `json:"price_marked"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            RefundCredit float64  `json:"refund_credit"`
            Cashback float64  `json:"cashback"`
            CouponValue float64  `json:"coupon_value"`
            PromotionEffectiveDiscount float64  `json:"promotion_effective_discount"`
            HsnCode string  `json:"hsn_code"`
            Identifiers Identifier  `json:"identifiers"`
            ValueOfGood float64  `json:"value_of_good"`
            AmountPaid float64  `json:"amount_paid"`
            ItemName string  `json:"item_name"`
            TotalUnits float64  `json:"total_units"`
            Size string  `json:"size"`
            FyndCredits float64  `json:"fynd_credits"`
            AmountPaidRoundoff float64  `json:"amount_paid_roundoff"`
            GstFee float64  `json:"gst_fee"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            GstTag string  `json:"gst_tag"`
            CashbackApplied float64  `json:"cashback_applied"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            AddedToFyndCash bool  `json:"added_to_fynd_cash"`
            DeliveryCharge float64  `json:"delivery_charge"`
            CodCharges float64  `json:"cod_charges"`
            TransferPrice float64  `json:"transfer_price"`
            PriceEffective float64  `json:"price_effective"`
         
    }
    
    // Weight used by Order
    type Weight struct {

        
            Unit string  `json:"unit"`
            IsDefault bool  `json:"is_default"`
            Shipping float64  `json:"shipping"`
         
    }
    
    // ReturnConfig used by Order
    type ReturnConfig struct {

        
            Unit string  `json:"unit"`
            Time float64  `json:"time"`
            Returnable bool  `json:"returnable"`
         
    }
    
    // Dimensions used by Order
    type Dimensions struct {

        
            IsDefault bool  `json:"is_default"`
            Unit string  `json:"unit"`
            Width float64  `json:"width"`
            Length float64  `json:"length"`
            Height float64  `json:"height"`
         
    }
    
    // Article used by Order
    type Article struct {

        
            SellerIdentifier string  `json:"seller_identifier"`
            ChildDetails map[string]interface{}  `json:"child_details"`
            Currency map[string]interface{}  `json:"currency"`
            Size string  `json:"size"`
            Weight Weight  `json:"weight"`
            Identifiers map[string]interface{}  `json:"identifiers"`
            IsSet bool  `json:"is_set"`
            ReturnConfig ReturnConfig  `json:"return_config"`
            Dimensions Dimensions  `json:"dimensions"`
            ASet map[string]interface{}  `json:"a_set"`
            UID string  `json:"uid"`
            EspModified bool  `json:"esp_modified"`
            RawMeta string  `json:"raw_meta"`
            Code string  `json:"code"`
            ID string  `json:"_id"`
         
    }
    
    // BagUnit used by Order
    type BagUnit struct {

        
            CurrentOperationalStatus BagStatusHistory  `json:"current_operational_status"`
            CurrentStatus BagStatusHistory  `json:"current_status"`
            Meta map[string]interface{}  `json:"meta"`
            BagID float64  `json:"bag_id"`
            CanReturn bool  `json:"can_return"`
            BagExpiryDate string  `json:"bag_expiry_date"`
            CanCancel bool  `json:"can_cancel"`
            Prices Prices  `json:"prices"`
            ProductQuantity float64  `json:"product_quantity"`
            Dates Dates  `json:"dates"`
            LineNumber float64  `json:"line_number"`
            Reasons []map[string]interface{}  `json:"reasons"`
            BagStatus []BagStatusHistory  `json:"bag_status"`
            EntityType string  `json:"entity_type"`
            Size string  `json:"size"`
            Gst GSTDetailsData  `json:"gst"`
            Status BagReturnableCancelableStatus  `json:"status"`
            AffiliateBagDetails AffiliateBagDetails  `json:"affiliate_bag_details"`
            Item PlatformItem  `json:"item"`
            Brand ShipmentListingBrand  `json:"brand"`
            FinancialBreakup []FinancialBreakup  `json:"financial_breakup"`
            BagType string  `json:"bag_type"`
            DisplayName string  `json:"display_name"`
            Article Article  `json:"article"`
         
    }
    
    // UserDataInfo used by Order
    type UserDataInfo struct {

        
            Mobile string  `json:"mobile"`
            IsAnonymousUser bool  `json:"is_anonymous_user"`
            LastName string  `json:"last_name"`
            FirstName string  `json:"first_name"`
            UID float64  `json:"uid"`
            Gender string  `json:"gender"`
            Name string  `json:"name"`
            Email string  `json:"email"`
            AvisUserID string  `json:"avis_user_id"`
         
    }
    
    // ShipmentListingChannel used by Order
    type ShipmentListingChannel struct {

        
            ChannelShipmentID string  `json:"channel_shipment_id"`
            IsAffiliate bool  `json:"is_affiliate"`
            Logo string  `json:"logo"`
            Name string  `json:"name"`
         
    }
    
    // ShipmentItemFulFillingStore used by Order
    type ShipmentItemFulFillingStore struct {

        
            StoreEmail string  `json:"store_email"`
            State string  `json:"state"`
            BrandStoreTags string  `json:"brand_store_tags"`
            ID float64  `json:"id"`
            Address string  `json:"address"`
            Meta map[string]interface{}  `json:"meta"`
            Pincode string  `json:"pincode"`
            Phone string  `json:"phone"`
            City string  `json:"city"`
            Name string  `json:"name"`
            LocationType string  `json:"location_type"`
            Code string  `json:"code"`
         
    }
    
    // PlatformDeliveryAddress used by Order
    type PlatformDeliveryAddress struct {

        
            AddressCategory string  `json:"address_category"`
            UpdatedAt string  `json:"updated_at"`
            State string  `json:"state"`
            Area string  `json:"area"`
            Latitude float64  `json:"latitude"`
            AddressType string  `json:"address_type"`
            Address2 string  `json:"address2"`
            Country string  `json:"country"`
            Pincode string  `json:"pincode"`
            Phone string  `json:"phone"`
            City string  `json:"city"`
            ContactPerson string  `json:"contact_person"`
            CreatedAt string  `json:"created_at"`
            Email string  `json:"email"`
            Address1 string  `json:"address1"`
            Landmark string  `json:"landmark"`
            Longitude float64  `json:"longitude"`
            Version string  `json:"version"`
         
    }
    
    // ShipmentItem used by Order
    type ShipmentItem struct {

        
            ShipmentStatus ShipmentStatus  `json:"shipment_status"`
            OrderDate string  `json:"order_date"`
            InvoiceID string  `json:"invoice_id"`
            LockStatus bool  `json:"lock_status"`
            Meta ShipmentItemMeta  `json:"meta"`
            PaymentMethods map[string]interface{}  `json:"payment_methods"`
            PreviousShipmentID string  `json:"previous_shipment_id"`
            EstimatedSlaTime string  `json:"estimated_sla_time"`
            CanProcess bool  `json:"can_process"`
            Bags []BagUnit  `json:"bags"`
            Prices Prices  `json:"prices"`
            PaymentMode string  `json:"payment_mode"`
            User UserDataInfo  `json:"user"`
            ShipmentID string  `json:"shipment_id"`
            TotalBags float64  `json:"total_bags"`
            StatusCreatedAt string  `json:"status_created_at"`
            OrderID string  `json:"order_id"`
            DisplayName string  `json:"display_name"`
            Channel ShipmentListingChannel  `json:"channel"`
            CustomerNote string  `json:"customer_note"`
            OrderingChannnel string  `json:"ordering_channnel"`
            ShipmentCreatedAt string  `json:"shipment_created_at"`
            FulfillingStore ShipmentItemFulFillingStore  `json:"fulfilling_store"`
            DeliveryAddress PlatformDeliveryAddress  `json:"delivery_address"`
         
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
            TotalCount float64  `json:"total_count"`
            Lane string  `json:"lane"`
            Page Page  `json:"page"`
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // Error used by Order
    type Error struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // TrackingList used by Order
    type TrackingList struct {

        
            Text string  `json:"text"`
            IsPassed bool  `json:"is_passed"`
            Status string  `json:"status"`
            IsCurrent bool  `json:"is_current"`
            Time string  `json:"time"`
         
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
    
    // BuyerDetails used by Order
    type BuyerDetails struct {

        
            State string  `json:"state"`
            Address string  `json:"address"`
            AjioSiteID string  `json:"ajio_site_id"`
            Pincode float64  `json:"pincode"`
            Gstin string  `json:"gstin"`
            City string  `json:"city"`
            Name string  `json:"name"`
         
    }
    
    // ShipmentMeta used by Order
    type ShipmentMeta struct {

        
            ShipmentVolumetricWeight float64  `json:"shipment_volumetric_weight"`
            Weight float64  `json:"weight"`
            DpID string  `json:"dp_id"`
            ShipmentWeight float64  `json:"shipment_weight"`
            ReturnAffiliateOrderID string  `json:"return_affiliate_order_id"`
            DpName string  `json:"dp_name"`
            StoreInvoiceUpdatedDate string  `json:"store_invoice_updated_date"`
            DpOptions map[string]interface{}  `json:"dp_options"`
            LockData LockData  `json:"lock_data"`
            B2cBuyerDetails map[string]interface{}  `json:"b2c_buyer_details"`
            FulfilmentPriorityText string  `json:"fulfilment_priority_text"`
            SameStoreAvailable bool  `json:"same_store_available"`
            BoxType string  `json:"box_type"`
            EwaybillInfo map[string]interface{}  `json:"ewaybill_info"`
            Dimension Dimensions  `json:"dimension"`
            DebugInfo DebugInfo  `json:"debug_info"`
            DueDate string  `json:"due_date"`
            OrderType string  `json:"order_type"`
            ReturnAffiliateShipmentID string  `json:"return_affiliate_shipment_id"`
            ReturnStoreNode float64  `json:"return_store_node"`
            DpSortKey string  `json:"dp_sort_key"`
            ForwardAffiliateShipmentID string  `json:"forward_affiliate_shipment_id"`
            AssignDpFromSb bool  `json:"assign_dp_from_sb"`
            PoNumber string  `json:"po_number"`
            PackagingName string  `json:"packaging_name"`
            ReturnDetails map[string]interface{}  `json:"return_details"`
            AwbNumber string  `json:"awb_number"`
            ForwardAffiliateOrderID string  `json:"forward_affiliate_order_id"`
            External map[string]interface{}  `json:"external"`
            ReturnAwbNumber string  `json:"return_awb_number"`
            MarketplaceStoreID string  `json:"marketplace_store_id"`
            ParentDpID string  `json:"parent_dp_id"`
            AutoTriggerDpAssignmentAcf bool  `json:"auto_trigger_dp_assignment_acf"`
            BagWeight map[string]interface{}  `json:"bag_weight"`
            EinvoiceInfo EinvoiceInfo  `json:"einvoice_info"`
            Formatted Formatted  `json:"formatted"`
            Timestamp ShipmentTimeStamp  `json:"timestamp"`
            B2bBuyerDetails BuyerDetails  `json:"b2b_buyer_details"`
            ShipmentTags []ShipmentTags  `json:"shipment_tags"`
         
    }
    
    // UserDetailsData used by Order
    type UserDetailsData struct {

        
            State string  `json:"state"`
            Area string  `json:"area"`
            AddressType string  `json:"address_type"`
            Address string  `json:"address"`
            Country string  `json:"country"`
            Pincode string  `json:"pincode"`
            Phone string  `json:"phone"`
            City string  `json:"city"`
            Name string  `json:"name"`
            Email string  `json:"email"`
            Address1 string  `json:"address1"`
            Landmark string  `json:"landmark"`
         
    }
    
    // PDFLinks used by Order
    type PDFLinks struct {

        
            InvoiceType string  `json:"invoice_type"`
            InvoiceA6 string  `json:"invoice_a6"`
            LabelExport string  `json:"label_export"`
            Label string  `json:"label"`
            InvoiceExport string  `json:"invoice_export"`
            DeliveryChallanA4 string  `json:"delivery_challan_a4"`
            PoInvoice string  `json:"po_invoice"`
            InvoicePos string  `json:"invoice_pos"`
            Invoice string  `json:"invoice"`
            InvoiceA4 string  `json:"invoice_a4"`
            LabelPos string  `json:"label_pos"`
            LabelA6 string  `json:"label_a6"`
            LabelType string  `json:"label_type"`
            CreditNoteURL string  `json:"credit_note_url"`
            LabelA4 string  `json:"label_a4"`
            B2b string  `json:"b2b"`
         
    }
    
    // AffiliateDetails used by Order
    type AffiliateDetails struct {

        
            AffiliateOrderID string  `json:"affiliate_order_id"`
            AffiliateMeta AffiliateMeta  `json:"affiliate_meta"`
            AffiliateID string  `json:"affiliate_id"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            AffiliateStoreID string  `json:"affiliate_store_id"`
            ShipmentMeta ShipmentMeta  `json:"shipment_meta"`
            CompanyAffiliateTag string  `json:"company_affiliate_tag"`
            AdID string  `json:"ad_id"`
            PdfLinks PDFLinks  `json:"pdf_links"`
         
    }
    
    // PhoneDetails used by Order
    type PhoneDetails struct {

        
            Number string  `json:"number"`
            CountryCode float64  `json:"country_code"`
         
    }
    
    // ContactDetails used by Order
    type ContactDetails struct {

        
            Emails []string  `json:"emails"`
            Phone []PhoneDetails  `json:"phone"`
         
    }
    
    // CompanyDetails used by Order
    type CompanyDetails struct {

        
            Address map[string]interface{}  `json:"address"`
            CompanyID float64  `json:"company_id"`
            CompanyCin string  `json:"company_cin"`
            CompanyContact ContactDetails  `json:"company_contact"`
            CompanyName string  `json:"company_name"`
            CompanyGst string  `json:"company_gst"`
         
    }
    
    // CurrentStatus used by Order
    type CurrentStatus struct {

        
            UpdatedAt float64  `json:"updated_at"`
            StateType string  `json:"state_type"`
            DeliveryPartnerID float64  `json:"delivery_partner_id"`
            StateID float64  `json:"state_id"`
            ShipmentID string  `json:"shipment_id"`
            DeliveryAwbNumber string  `json:"delivery_awb_number"`
            ID float64  `json:"id"`
            Status string  `json:"status"`
            KafkaSync bool  `json:"kafka_sync"`
            BagStateMapper BagStateMapper  `json:"bag_state_mapper"`
            CreatedAt string  `json:"created_at"`
            StoreID float64  `json:"store_id"`
            BagID float64  `json:"bag_id"`
         
    }
    
    // B2BPODetails used by Order
    type B2BPODetails struct {

        
            PartialCanRet bool  `json:"partial_can_ret"`
            TotalGstPercentage float64  `json:"total_gst_percentage"`
            PoTaxAmount float64  `json:"po_tax_amount"`
            DocketNumber string  `json:"docket_number"`
            ItemBasePrice float64  `json:"item_base_price"`
            PoLineAmount float64  `json:"po_line_amount"`
         
    }
    
    // GiftCard used by Order
    type GiftCard struct {

        
            GiftPrice float64  `json:"gift_price"`
            GiftMessage string  `json:"gift_message"`
            DisplayText string  `json:"display_text"`
            IsGiftApplied bool  `json:"is_gift_applied"`
         
    }
    
    // BagMeta used by Order
    type BagMeta struct {

        
            GroupID string  `json:"group_id"`
            CustomJson map[string]interface{}  `json:"custom_json"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            PartialCanRet bool  `json:"partial_can_ret"`
            DocketNumber string  `json:"docket_number"`
            B2bPoDetails B2BPODetails  `json:"b2b_po_details"`
            GiftCard GiftCard  `json:"gift_card"`
            CustomMessage string  `json:"custom_message"`
         
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

        
            CartConditions map[string]interface{}  `json:"cart_conditions"`
            ItemCriteria ItemCriterias  `json:"item_criteria"`
         
    }
    
    // AppliedPromos used by Order
    type AppliedPromos struct {

        
            PromotionType string  `json:"promotion_type"`
            MrpPromotion bool  `json:"mrp_promotion"`
            DiscountRules []DiscountRules  `json:"discount_rules"`
            PromoID string  `json:"promo_id"`
            Amount float64  `json:"amount"`
            BuyRules []BuyRules  `json:"buy_rules"`
            PromotionName string  `json:"promotion_name"`
            ArticleQuantity float64  `json:"article_quantity"`
         
    }
    
    // BagGST used by Order
    type BagGST struct {

        
            IsDefaultHsnCode bool  `json:"is_default_hsn_code"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            CgstGstFee string  `json:"cgst_gst_fee"`
            GstinCode string  `json:"gstin_code"`
            HsnCodeID string  `json:"hsn_code_id"`
            IgstTaxPercentage float64  `json:"igst_tax_percentage"`
            IgstGstFee string  `json:"igst_gst_fee"`
            SgstGstFee string  `json:"sgst_gst_fee"`
            HsnCode string  `json:"hsn_code"`
            GstFee float64  `json:"gst_fee"`
            CgstTaxPercentage float64  `json:"cgst_tax_percentage"`
            SgstTaxPercentage float64  `json:"sgst_tax_percentage"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            ValueOfGood float64  `json:"value_of_good"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            GstTag string  `json:"gst_tag"`
         
    }
    
    // AffiliateBagsDetails used by Order
    type AffiliateBagsDetails struct {

        
            CouponCode string  `json:"coupon_code"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
         
    }
    
    // BagConfigs used by Order
    type BagConfigs struct {

        
            IsActive bool  `json:"is_active"`
            AllowForceReturn bool  `json:"allow_force_return"`
            EnableTracking bool  `json:"enable_tracking"`
            IsCustomerReturnAllowed bool  `json:"is_customer_return_allowed"`
            CanBeCancelled bool  `json:"can_be_cancelled"`
            IsReturnable bool  `json:"is_returnable"`
         
    }
    
    // OrderBrandName used by Order
    type OrderBrandName struct {

        
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
            BrandName string  `json:"brand_name"`
            ID float64  `json:"id"`
            Company float64  `json:"company"`
            Logo string  `json:"logo"`
         
    }
    
    // ReturnConfig1 used by Order
    type ReturnConfig1 struct {

        
            Unit string  `json:"unit"`
            Time float64  `json:"time"`
            Returnable bool  `json:"returnable"`
         
    }
    
    // OrderBagArticle used by Order
    type OrderBagArticle struct {

        
            ReturnConfig ReturnConfig1  `json:"return_config"`
            Size string  `json:"size"`
            Identifiers map[string]interface{}  `json:"identifiers"`
            UID string  `json:"uid"`
         
    }
    
    // OrderBags used by Order
    type OrderBags struct {

        
            SellerIdentifier string  `json:"seller_identifier"`
            CurrentStatus CurrentStatus  `json:"current_status"`
            Meta BagMeta  `json:"meta"`
            AppliedPromos []AppliedPromos  `json:"applied_promos"`
            BagID float64  `json:"bag_id"`
            CanReturn bool  `json:"can_return"`
            CanCancel bool  `json:"can_cancel"`
            Prices Prices  `json:"prices"`
            LineNumber float64  `json:"line_number"`
            GstDetails BagGST  `json:"gst_details"`
            EntityType string  `json:"entity_type"`
            ParentPromoBags map[string]interface{}  `json:"parent_promo_bags"`
            AffiliateBagDetails AffiliateBagsDetails  `json:"affiliate_bag_details"`
            GroupID string  `json:"group_id"`
            Item PlatformItem  `json:"item"`
            BagConfigs BagConfigs  `json:"bag_configs"`
            Brand OrderBrandName  `json:"brand"`
            FinancialBreakup FinancialBreakup  `json:"financial_breakup"`
            DisplayName string  `json:"display_name"`
            Identifier string  `json:"identifier"`
            Quantity float64  `json:"quantity"`
            Article OrderBagArticle  `json:"article"`
            DeliveryAddress PlatformDeliveryAddress  `json:"delivery_address"`
         
    }
    
    // InvoiceInfo used by Order
    type InvoiceInfo struct {

        
            LabelURL string  `json:"label_url"`
            CreditNoteID string  `json:"credit_note_id"`
            InvoiceURL string  `json:"invoice_url"`
            ExternalInvoiceID string  `json:"external_invoice_id"`
            UpdatedDate string  `json:"updated_date"`
            StoreInvoiceID string  `json:"store_invoice_id"`
         
    }
    
    // DPDetailsData used by Order
    type DPDetailsData struct {

        
            AwbNo string  `json:"awb_no"`
            EwayBillID string  `json:"eway_bill_id"`
            ID float64  `json:"id"`
            Country string  `json:"country"`
            TrackURL string  `json:"track_url"`
            Pincode string  `json:"pincode"`
            Name string  `json:"name"`
            GstTag string  `json:"gst_tag"`
         
    }
    
    // ShipmentStatusData used by Order
    type ShipmentStatusData struct {

        
            BagList []string  `json:"bag_list"`
            ShipmentID string  `json:"shipment_id"`
            ID float64  `json:"id"`
            DisplayName string  `json:"display_name"`
            Status string  `json:"status"`
            Meta map[string]interface{}  `json:"meta"`
            CreatedAt string  `json:"created_at"`
         
    }
    
    // OrderDetailsData used by Order
    type OrderDetailsData struct {

        
            OrderValue string  `json:"order_value"`
            AffiliateID string  `json:"affiliate_id"`
            OrderDate string  `json:"order_date"`
            FyndOrderID string  `json:"fynd_order_id"`
            CodCharges string  `json:"cod_charges"`
            OrderingChannelLogo map[string]interface{}  `json:"ordering_channel_logo"`
            OrderingChannel string  `json:"ordering_channel"`
            Source string  `json:"source"`
            TaxDetails map[string]interface{}  `json:"tax_details"`
         
    }
    
    // OrderingStoreDetails used by Order
    type OrderingStoreDetails struct {

        
            State string  `json:"state"`
            ID float64  `json:"id"`
            Address string  `json:"address"`
            Country string  `json:"country"`
            Meta map[string]interface{}  `json:"meta"`
            Pincode string  `json:"pincode"`
            Phone string  `json:"phone"`
            City string  `json:"city"`
            Code string  `json:"code"`
            StoreName string  `json:"store_name"`
            ContactPerson string  `json:"contact_person"`
         
    }
    
    // ShipmentDetails used by Order
    type ShipmentDetails struct {

        
            ActionToStatus map[string]interface{}  `json:"action_to_status"`
            LockMessage string  `json:"lock_message"`
            LockStatus bool  `json:"lock_status"`
         
    }
    
    // FulfillingStore used by Order
    type FulfillingStore struct {

        
            State string  `json:"state"`
            ID float64  `json:"id"`
            Address string  `json:"address"`
            Country string  `json:"country"`
            Meta map[string]interface{}  `json:"meta"`
            Pincode string  `json:"pincode"`
            Phone string  `json:"phone"`
            City string  `json:"city"`
            Code string  `json:"code"`
            StoreName string  `json:"store_name"`
            ContactPerson string  `json:"contact_person"`
            FulfillmentChannel string  `json:"fulfillment_channel"`
         
    }
    
    // ShipmentPayments used by Order
    type ShipmentPayments struct {

        
            Mode string  `json:"mode"`
            Logo string  `json:"logo"`
            Source string  `json:"source"`
         
    }
    
    // PlatformShipment used by Order
    type PlatformShipment struct {

        
            DeliverySlot map[string]interface{}  `json:"delivery_slot"`
            UserAgent string  `json:"user_agent"`
            ShipmentStatus string  `json:"shipment_status"`
            TrackingList []TrackingList  `json:"tracking_list"`
            InvoiceID string  `json:"invoice_id"`
            Meta ShipmentMeta  `json:"meta"`
            LockStatus bool  `json:"lock_status"`
            DeliveryDetails UserDetailsData  `json:"delivery_details"`
            PaymentMethods map[string]interface{}  `json:"payment_methods"`
            BagStatusHistory []BagStatusHistory  `json:"bag_status_history"`
            AffiliateDetails AffiliateDetails  `json:"affiliate_details"`
            EstimatedSlaTime string  `json:"estimated_sla_time"`
            Vertical string  `json:"vertical"`
            CompanyDetails CompanyDetails  `json:"company_details"`
            PickedDate string  `json:"picked_date"`
            OperationalStatus string  `json:"operational_status"`
            ShipmentQuantity float64  `json:"shipment_quantity"`
            Bags []OrderBags  `json:"bags"`
            Invoice InvoiceInfo  `json:"invoice"`
            DpAssignment bool  `json:"dp_assignment"`
            JourneyType string  `json:"journey_type"`
            Prices Prices  `json:"prices"`
            PaymentMode string  `json:"payment_mode"`
            User UserDataInfo  `json:"user"`
            CustomMessage string  `json:"custom_message"`
            DpDetails DPDetailsData  `json:"dp_details"`
            FulfilmentPriority float64  `json:"fulfilment_priority"`
            PdfLinks map[string]interface{}  `json:"pdf_links"`
            ForwardShipmentID string  `json:"forward_shipment_id"`
            GstDetails GSTDetailsData  `json:"gst_details"`
            IsDpAssignEnabled bool  `json:"is_dp_assign_enabled"`
            ShipmentID string  `json:"shipment_id"`
            TotalBags float64  `json:"total_bags"`
            ShipmentImages []string  `json:"shipment_images"`
            Status ShipmentStatusData  `json:"status"`
            Order OrderDetailsData  `json:"order"`
            Coupon map[string]interface{}  `json:"coupon"`
            PackagingType string  `json:"packaging_type"`
            OrderingStore OrderingStoreDetails  `json:"ordering_store"`
            CustomMeta []map[string]interface{}  `json:"custom_meta"`
            PriorityText string  `json:"priority_text"`
            ShipmentDetails ShipmentDetails  `json:"shipment_details"`
            PlatformLogo string  `json:"platform_logo"`
            ShipmentCreatedAt string  `json:"shipment_created_at"`
            EnableDpTracking bool  `json:"enable_dp_tracking"`
            FulfillingStore FulfillingStore  `json:"fulfilling_store"`
            Payments ShipmentPayments  `json:"payments"`
            CanUpdateDimension bool  `json:"can_update_dimension"`
            BillingDetails UserDetailsData  `json:"billing_details"`
            TotalItems float64  `json:"total_items"`
         
    }
    
    // ShipmentInfoResponse used by Order
    type ShipmentInfoResponse struct {

        
            Shipments []PlatformShipment  `json:"shipments"`
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // BillingStaffDetails used by Order
    type BillingStaffDetails struct {

        
            LastName string  `json:"last_name"`
            StaffID float64  `json:"staff_id"`
            User string  `json:"user"`
            FirstName string  `json:"first_name"`
            EmployeeCode string  `json:"employee_code"`
         
    }
    
    // TransactionData used by Order
    type TransactionData struct {

        
            PaymentID string  `json:"payment_id"`
            TerminalID string  `json:"terminal_id"`
            Currency string  `json:"currency"`
            UniqueReferenceNumber string  `json:"unique_reference_number"`
            Status string  `json:"status"`
            TransactionID string  `json:"transaction_id"`
            Entity string  `json:"entity"`
            AmountPaid float64  `json:"amount_paid"`
         
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
            Files []map[string]interface{}  `json:"files"`
            CurrencySymbol string  `json:"currency_symbol"`
            MongoCartID float64  `json:"mongo_cart_id"`
            OrderType string  `json:"order_type"`
            PaymentType string  `json:"payment_type"`
            OrderTags []map[string]interface{}  `json:"order_tags"`
            CompanyLogo string  `json:"company_logo"`
            OrderPlatform string  `json:"order_platform"`
            BillingStaffDetails BillingStaffDetails  `json:"billing_staff_details"`
            Comment string  `json:"comment"`
            CartID float64  `json:"cart_id"`
            OrderingStore float64  `json:"ordering_store"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            TransactionData TransactionData  `json:"transaction_data"`
            CustomerNote string  `json:"customer_note"`
            Staff map[string]interface{}  `json:"staff"`
            PlatformUserDetails PlatformUserDetails  `json:"platform_user_details"`
            EmployeeID float64  `json:"employee_id"`
         
    }
    
    // TaxDetails used by Order
    type TaxDetails struct {

        
            Gstin string  `json:"gstin"`
            PanNo string  `json:"pan_no"`
         
    }
    
    // OrderData used by Order
    type OrderData struct {

        
            OrderDate string  `json:"order_date"`
            FyndOrderID string  `json:"fynd_order_id"`
            Meta OrderMeta  `json:"meta"`
            PaymentMethods map[string]interface{}  `json:"payment_methods"`
            Prices Prices  `json:"prices"`
            TaxDetails TaxDetails  `json:"tax_details"`
         
    }
    
    // OrderDetailsResponse used by Order
    type OrderDetailsResponse struct {

        
            Order OrderData  `json:"order"`
            Shipments []PlatformShipment  `json:"shipments"`
            Success bool  `json:"success"`
         
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
            Value string  `json:"value"`
            Text string  `json:"text"`
         
    }
    
    // LaneConfigResponse used by Order
    type LaneConfigResponse struct {

        
            SuperLanes []SuperLane  `json:"super_lanes"`
         
    }
    
    // PlatformBreakupValues used by Order
    type PlatformBreakupValues struct {

        
            Display string  `json:"display"`
            Value string  `json:"value"`
            Name string  `json:"name"`
         
    }
    
    // PlatformChannel used by Order
    type PlatformChannel struct {

        
            Logo string  `json:"logo"`
            Name string  `json:"name"`
         
    }
    
    // PlatformOrderItems used by Order
    type PlatformOrderItems struct {

        
            UserInfo UserDataInfo  `json:"user_info"`
            OrderValue float64  `json:"order_value"`
            OrderID string  `json:"order_id"`
            BreakupValues []PlatformBreakupValues  `json:"breakup_values"`
            Meta map[string]interface{}  `json:"meta"`
            Channel PlatformChannel  `json:"channel"`
            Shipments []PlatformShipment  `json:"shipments"`
            PaymentMode string  `json:"payment_mode"`
            TotalOrderValue float64  `json:"total_order_value"`
            OrderCreatedTime string  `json:"order_created_time"`
         
    }
    
    // OrderListingResponse used by Order
    type OrderListingResponse struct {

        
            Items []PlatformOrderItems  `json:"items"`
            TotalCount float64  `json:"total_count"`
            Lane string  `json:"lane"`
            Page Page  `json:"page"`
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // PlatformTrack used by Order
    type PlatformTrack struct {

        
            UpdatedAt string  `json:"updated_at"`
            Reason string  `json:"reason"`
            AccountName string  `json:"account_name"`
            LastLocationRecievedAt string  `json:"last_location_recieved_at"`
            UpdatedTime string  `json:"updated_time"`
            Awb string  `json:"awb"`
            RawStatus string  `json:"raw_status"`
            Meta map[string]interface{}  `json:"meta"`
            ShipmentType string  `json:"shipment_type"`
            Status string  `json:"status"`
         
    }
    
    // PlatformShipmentTrack used by Order
    type PlatformShipmentTrack struct {

        
            Results []PlatformTrack  `json:"results"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // FilterInfoOption used by Order
    type FilterInfoOption struct {

        
            PlaceholderText string  `json:"placeholder_text"`
            ShowUI bool  `json:"show_ui"`
            Text string  `json:"text"`
            MinSearchSize float64  `json:"min_search_size"`
            Name string  `json:"name"`
            Value string  `json:"value"`
         
    }
    
    // FiltersInfo used by Order
    type FiltersInfo struct {

        
            PlaceholderText string  `json:"placeholder_text"`
            Text string  `json:"text"`
            Required bool  `json:"required"`
            Type string  `json:"type"`
            Value string  `json:"value"`
            Options []FilterInfoOption  `json:"options"`
         
    }
    
    // AdvanceFilterInfo used by Order
    type AdvanceFilterInfo struct {

        
            Processed []FiltersInfo  `json:"processed"`
            ActionCentre []FiltersInfo  `json:"action_centre"`
            Filters []FiltersInfo  `json:"filters"`
            AppliedFilters map[string]interface{}  `json:"applied_filters"`
            Unfulfilled []FiltersInfo  `json:"unfulfilled"`
            Page map[string]interface{}  `json:"page"`
            Returned []FiltersInfo  `json:"returned"`
         
    }
    
    // FiltersResponse used by Order
    type FiltersResponse struct {

        
            AdvanceFilter AdvanceFilterInfo  `json:"advance_filter"`
            GlobalFilter []FiltersInfo  `json:"global_filter"`
         
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

        
            DisplayName string  `json:"display_name"`
            QuestionSet []QuestionSet  `json:"question_set"`
            QcType []string  `json:"qc_type"`
            ID float64  `json:"id"`
         
    }
    
    // PlatformShipmentReasonsResponse used by Order
    type PlatformShipmentReasonsResponse struct {

        
            Reasons []Reason  `json:"reasons"`
            Success bool  `json:"success"`
         
    }
    
    // ShipmentResponseReasons used by Order
    type ShipmentResponseReasons struct {

        
            Reason string  `json:"reason"`
            ReasonID float64  `json:"reason_id"`
         
    }
    
    // ShipmentReasonsResponse used by Order
    type ShipmentReasonsResponse struct {

        
            Reasons []ShipmentResponseReasons  `json:"reasons"`
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // ArticleDetails used by Order
    type ArticleDetails struct {

        
            Status map[string]interface{}  `json:"status"`
         
    }
    
    // BagGSTDetails used by Order
    type BagGSTDetails struct {

        
            IsDefaultHsnCode bool  `json:"is_default_hsn_code"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            CgstGstFee string  `json:"cgst_gst_fee"`
            GstinCode string  `json:"gstin_code"`
            HsnCodeID string  `json:"hsn_code_id"`
            IgstTaxPercentage float64  `json:"igst_tax_percentage"`
            IgstGstFee string  `json:"igst_gst_fee"`
            SgstGstFee string  `json:"sgst_gst_fee"`
            HsnCode string  `json:"hsn_code"`
            GstFee float64  `json:"gst_fee"`
            CgstTaxPercentage float64  `json:"cgst_tax_percentage"`
            SgstTaxPercentage float64  `json:"sgst_tax_percentage"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            ValueOfGood float64  `json:"value_of_good"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            GstTag string  `json:"gst_tag"`
         
    }
    
    // BagReturnableCancelableStatus1 used by Order
    type BagReturnableCancelableStatus1 struct {

        
            IsActive bool  `json:"is_active"`
            EnableTracking bool  `json:"enable_tracking"`
            IsCustomerReturnAllowed bool  `json:"is_customer_return_allowed"`
            CanBeCancelled bool  `json:"can_be_cancelled"`
            IsReturnable bool  `json:"is_returnable"`
         
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
    
    // Document used by Order
    type Document struct {

        
            LegalName string  `json:"legal_name"`
            Verified bool  `json:"verified"`
            DsType string  `json:"ds_type"`
            URL string  `json:"url"`
            Value string  `json:"value"`
         
    }
    
    // StoreDocuments used by Order
    type StoreDocuments struct {

        
            Gst Document  `json:"gst"`
         
    }
    
    // EInvoicePortalDetails used by Order
    type EInvoicePortalDetails struct {

        
            Username string  `json:"username"`
            Password string  `json:"password"`
            User string  `json:"user"`
         
    }
    
    // StoreMeta used by Order
    type StoreMeta struct {

        
            EwaybillPortalDetails map[string]interface{}  `json:"ewaybill_portal_details"`
            GstCredentials StoreGstCredentials  `json:"gst_credentials"`
            Timing []map[string]interface{}  `json:"timing"`
            NotificationEmails []string  `json:"notification_emails"`
            GstNumber string  `json:"gst_number"`
            AdditionalContactDetails map[string]interface{}  `json:"additional_contact_details"`
            DisplayName string  `json:"display_name"`
            Stage string  `json:"stage"`
            Documents StoreDocuments  `json:"documents"`
            ProductReturnConfig map[string]interface{}  `json:"product_return_config"`
            EinvoicePortalDetails EInvoicePortalDetails  `json:"einvoice_portal_details"`
         
    }
    
    // StoreAddress used by Order
    type StoreAddress struct {

        
            AddressType string  `json:"address_type"`
            Email string  `json:"email"`
            Address1 string  `json:"address1"`
            Landmark string  `json:"landmark"`
            AddressCategory string  `json:"address_category"`
            State string  `json:"state"`
            Address2 string  `json:"address2"`
            Country string  `json:"country"`
            Phone string  `json:"phone"`
            CreatedAt string  `json:"created_at"`
            ContactPerson string  `json:"contact_person"`
            City string  `json:"city"`
            Latitude float64  `json:"latitude"`
            Longitude float64  `json:"longitude"`
            Version string  `json:"version"`
            UpdatedAt string  `json:"updated_at"`
            Area string  `json:"area"`
            Pincode float64  `json:"pincode"`
            CountryCode string  `json:"country_code"`
         
    }
    
    // Store used by Order
    type Store struct {

        
            VatNo string  `json:"vat_no"`
            BrandID interface{}  `json:"brand_id"`
            Meta StoreMeta  `json:"meta"`
            ParentStoreID float64  `json:"parent_store_id"`
            LocationType string  `json:"location_type"`
            Address1 string  `json:"address1"`
            FulfillmentChannel string  `json:"fulfillment_channel"`
            LoginUsername string  `json:"login_username"`
            State string  `json:"state"`
            StoreEmail string  `json:"store_email"`
            Address2 string  `json:"address2"`
            StoreActiveFrom string  `json:"store_active_from"`
            Country string  `json:"country"`
            CompanyID float64  `json:"company_id"`
            Phone float64  `json:"phone"`
            CreatedAt string  `json:"created_at"`
            IsEnabledForRecon bool  `json:"is_enabled_for_recon"`
            MallName string  `json:"mall_name"`
            ContactPerson string  `json:"contact_person"`
            OrderIntegrationID string  `json:"order_integration_id"`
            PackagingMaterialCount float64  `json:"packaging_material_count"`
            IsActive bool  `json:"is_active"`
            IsArchived bool  `json:"is_archived"`
            City string  `json:"city"`
            Name string  `json:"name"`
            Code string  `json:"code"`
            Latitude float64  `json:"latitude"`
            Longitude float64  `json:"longitude"`
            UpdatedAt string  `json:"updated_at"`
            BrandStoreTags []string  `json:"brand_store_tags"`
            MallArea string  `json:"mall_area"`
            StoreAddressJson StoreAddress  `json:"store_address_json"`
            SID string  `json:"s_id"`
            Pincode string  `json:"pincode"`
            AlohomoraUserID float64  `json:"alohomora_user_id"`
         
    }
    
    // Attributes used by Order
    type Attributes struct {

        
            PrimaryColorHex string  `json:"primary_color_hex"`
            PrimaryMaterial string  `json:"primary_material"`
            Essential string  `json:"essential"`
            MarketerName string  `json:"marketer_name"`
            BrandName string  `json:"brand_name"`
            PrimaryColor string  `json:"primary_color"`
            MarketerAddress string  `json:"marketer_address"`
            Gender []string  `json:"gender"`
            Name string  `json:"name"`
         
    }
    
    // Item used by Order
    type Item struct {

        
            ItemID float64  `json:"item_id"`
            Image []string  `json:"image"`
            BrandID float64  `json:"brand_id"`
            L3Category float64  `json:"l3_category"`
            L1Category []string  `json:"l1_category"`
            Meta map[string]interface{}  `json:"meta"`
            SlugKey string  `json:"slug_key"`
            Attributes Attributes  `json:"attributes"`
            WebstoreProductURL string  `json:"webstore_product_url"`
            L3CategoryName string  `json:"l3_category_name"`
            CanReturn bool  `json:"can_return"`
            CanCancel bool  `json:"can_cancel"`
            L2Category []string  `json:"l2_category"`
            Gender string  `json:"gender"`
            L2CategoryID float64  `json:"l2_category_id"`
            BranchURL string  `json:"branch_url"`
            Size string  `json:"size"`
            LastUpdatedAt string  `json:"last_updated_at"`
            Color string  `json:"color"`
            Name string  `json:"name"`
            Code string  `json:"code"`
            Brand string  `json:"brand"`
            DepartmentID float64  `json:"department_id"`
            L1CategoryID float64  `json:"l1_category_id"`
         
    }
    
    // Brand used by Order
    type Brand struct {

        
            ModifiedOn float64  `json:"modified_on"`
            CreatedOn float64  `json:"created_on"`
            PickupLocation string  `json:"pickup_location"`
            BrandID float64  `json:"brand_id"`
            BrandName string  `json:"brand_name"`
            Logo string  `json:"logo"`
            Company string  `json:"company"`
            IsVirtualInvoice bool  `json:"is_virtual_invoice"`
            ScriptLastRan string  `json:"script_last_ran"`
            CreditNoteExpiryDays float64  `json:"credit_note_expiry_days"`
            InvoicePrefix string  `json:"invoice_prefix"`
            CreditNoteAllowed bool  `json:"credit_note_allowed"`
            StartDate string  `json:"start_date"`
         
    }
    
    // BagDetailsPlatformResponse used by Order
    type BagDetailsPlatformResponse struct {

        
            SellerIdentifier string  `json:"seller_identifier"`
            CurrentOperationalStatus BagStatusHistory  `json:"current_operational_status"`
            ArticleDetails ArticleDetails  `json:"article_details"`
            CurrentStatus BagStatusHistory  `json:"current_status"`
            Meta BagMeta  `json:"meta"`
            BagStatusHistory BagStatusHistory  `json:"bag_status_history"`
            AppliedPromos []map[string]interface{}  `json:"applied_promos"`
            AffiliateDetails AffiliateDetails  `json:"affiliate_details"`
            BagUpdateTime float64  `json:"bag_update_time"`
            Type string  `json:"type"`
            OriginalBagList []float64  `json:"original_bag_list"`
            OperationalStatus string  `json:"operational_status"`
            ID float64  `json:"id"`
            RestoreCoupon bool  `json:"restore_coupon"`
            JourneyType string  `json:"journey_type"`
            Prices Prices  `json:"prices"`
            Dates Dates  `json:"dates"`
            LineNumber float64  `json:"line_number"`
            OrderIntegrationID string  `json:"order_integration_id"`
            Reasons []map[string]interface{}  `json:"reasons"`
            BagStatus []BagStatusHistory  `json:"bag_status"`
            EntityType string  `json:"entity_type"`
            GstDetails BagGSTDetails  `json:"gst_details"`
            ParentPromoBags map[string]interface{}  `json:"parent_promo_bags"`
            ShipmentID string  `json:"shipment_id"`
            Status BagReturnableCancelableStatus1  `json:"status"`
            QcRequired interface{}  `json:"qc_required"`
            Tags []string  `json:"tags"`
            AffiliateBagDetails AffiliateBagDetails  `json:"affiliate_bag_details"`
            OrderingStore Store  `json:"ordering_store"`
            Item Item  `json:"item"`
            Brand Brand  `json:"brand"`
            FinancialBreakup []FinancialBreakup  `json:"financial_breakup"`
            RestorePromos map[string]interface{}  `json:"restore_promos"`
            DisplayName string  `json:"display_name"`
            Identifier string  `json:"identifier"`
            Quantity float64  `json:"quantity"`
            Article Article  `json:"article"`
            NoOfBagsOrder float64  `json:"no_of_bags_order"`
         
    }
    
    // ErrorResponse used by Order
    type ErrorResponse struct {

        
            Error string  `json:"error"`
            Message string  `json:"message"`
         
    }
    
    // BagsPage used by Order
    type BagsPage struct {

        
            PageType string  `json:"page_type"`
            Size float64  `json:"size"`
            Current float64  `json:"current"`
            ItemTotal float64  `json:"item_total"`
            HasNext bool  `json:"has_next"`
         
    }
    
    // GetBagsPlatformResponse used by Order
    type GetBagsPlatformResponse struct {

        
            Page BagsPage  `json:"page"`
            Items []BagDetailsPlatformResponse  `json:"items"`
         
    }
    
    // InvalidateShipmentCachePayload used by Order
    type InvalidateShipmentCachePayload struct {

        
            BagIds []string  `json:"bag_ids"`
            AffiliateBagIds []string  `json:"affiliate_bag_ids"`
            ShipmentIds []string  `json:"shipment_ids"`
         
    }
    
    // InvalidateShipmentCacheNestedResponse used by Order
    type InvalidateShipmentCacheNestedResponse struct {

        
            ShipmentID string  `json:"shipment_id"`
            Status float64  `json:"status"`
            Message string  `json:"message"`
            Error string  `json:"error"`
         
    }
    
    // InvalidateShipmentCacheResponse used by Order
    type InvalidateShipmentCacheResponse struct {

        
            Response []InvalidateShipmentCacheNestedResponse  `json:"response"`
         
    }
    
    // ErrorResponse1 used by Order
    type ErrorResponse1 struct {

        
            Status float64  `json:"status"`
            Message string  `json:"message"`
            ErrorTrace string  `json:"error_trace"`
            Success bool  `json:"success"`
         
    }
    
    // StoreReassign used by Order
    type StoreReassign struct {

        
            AffiliateOrderID string  `json:"affiliate_order_id"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
            ItemID string  `json:"item_id"`
            FyndOrderID string  `json:"fynd_order_id"`
            StoreID float64  `json:"store_id"`
            SetID string  `json:"set_id"`
            ReasonIds []float64  `json:"reason_ids"`
            BagID float64  `json:"bag_id"`
            MongoArticleID string  `json:"mongo_article_id"`
            AffiliateID string  `json:"affiliate_id"`
         
    }
    
    // StoreReassignResponse used by Order
    type StoreReassignResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // Entities used by Order
    type Entities struct {

        
            AffiliateOrderID string  `json:"affiliate_order_id"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
            ID string  `json:"id"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            AffiliateID string  `json:"affiliate_id"`
            ReasonText string  `json:"reason_text"`
         
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

        
            AffiliateBagID string  `json:"affiliate_bag_id"`
            BagID float64  `json:"bag_id"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            IsLocked bool  `json:"is_locked"`
         
    }
    
    // OriginalFilter used by Order
    type OriginalFilter struct {

        
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            AffiliateID string  `json:"affiliate_id"`
         
    }
    
    // CheckResponse used by Order
    type CheckResponse struct {

        
            Bags []Bags  `json:"bags"`
            IsBagLocked bool  `json:"is_bag_locked"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            OriginalFilter OriginalFilter  `json:"original_filter"`
            LockStatus string  `json:"lock_status"`
            IsShipmentLocked bool  `json:"is_shipment_locked"`
            ShipmentID string  `json:"shipment_id"`
            Status string  `json:"status"`
            AffiliateID string  `json:"affiliate_id"`
         
    }
    
    // UpdateShipmentLockResponse used by Order
    type UpdateShipmentLockResponse struct {

        
            Message string  `json:"message"`
            CheckResponse []CheckResponse  `json:"check_response"`
            Success bool  `json:"success"`
         
    }
    
    // AnnouncementResponse used by Order
    type AnnouncementResponse struct {

        
            Title string  `json:"title"`
            ID float64  `json:"id"`
            CreatedAt string  `json:"created_at"`
            PlatformName string  `json:"platform_name"`
            ToDatetime string  `json:"to_datetime"`
            FromDatetime string  `json:"from_datetime"`
            Description string  `json:"description"`
            PlatformID string  `json:"platform_id"`
            LogoURL string  `json:"logo_url"`
            CompanyID float64  `json:"company_id"`
         
    }
    
    // AnnouncementsResponse used by Order
    type AnnouncementsResponse struct {

        
            Announcements []AnnouncementResponse  `json:"announcements"`
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // BaseResponse used by Order
    type BaseResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // Click2CallResponse used by Order
    type Click2CallResponse struct {

        
            CallID string  `json:"call_id"`
            Success bool  `json:"success"`
         
    }
    
    // ErrorDetail used by Order
    type ErrorDetail struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // EntitiesDataUpdates used by Order
    type EntitiesDataUpdates struct {

        
            Filters []map[string]interface{}  `json:"filters"`
            Data map[string]interface{}  `json:"data"`
         
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
    
    // DataUpdates used by Order
    type DataUpdates struct {

        
            Entities []EntitiesDataUpdates  `json:"entities"`
            Products []ProductsDataUpdates  `json:"products"`
         
    }
    
    // Products used by Order
    type Products struct {

        
            Identifier string  `json:"identifier"`
            LineNumber float64  `json:"line_number"`
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
            ExcludeBagsNextState string  `json:"exclude_bags_next_state"`
            Status string  `json:"status"`
            SplitShipment bool  `json:"split_shipment"`
         
    }
    
    // UpdateShipmentStatusRequest used by Order
    type UpdateShipmentStatusRequest struct {

        
            ForceTransition bool  `json:"force_transition"`
            LockAfterTransition bool  `json:"lock_after_transition"`
            Statuses []StatuesRequest  `json:"statuses"`
            UnlockBeforeTransition bool  `json:"unlock_before_transition"`
            Task bool  `json:"task"`
         
    }
    
    // ShipmentsResponse used by Order
    type ShipmentsResponse struct {

        
            FinalState map[string]interface{}  `json:"final_state"`
            Identifier string  `json:"identifier"`
            Code string  `json:"code"`
            Message string  `json:"message"`
            Exception string  `json:"exception"`
            Status float64  `json:"status"`
            StackTrace string  `json:"stack_trace"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // StatuesResponse used by Order
    type StatuesResponse struct {

        
            Shipments []ShipmentsResponse  `json:"shipments"`
         
    }
    
    // UpdateShipmentStatusResponseBody used by Order
    type UpdateShipmentStatusResponseBody struct {

        
            Statuses []StatuesResponse  `json:"statuses"`
         
    }
    
    // AffiliateAppConfigMeta used by Order
    type AffiliateAppConfigMeta struct {

        
            Value string  `json:"value"`
            Name string  `json:"name"`
         
    }
    
    // AffiliateAppConfig used by Order
    type AffiliateAppConfig struct {

        
            Name string  `json:"name"`
            ID string  `json:"id"`
            Secret string  `json:"secret"`
            UpdatedAt string  `json:"updated_at"`
            Token string  `json:"token"`
            Owner string  `json:"owner"`
            Description string  `json:"description"`
            CreatedAt string  `json:"created_at"`
            Meta []AffiliateAppConfigMeta  `json:"meta"`
         
    }
    
    // AffiliateInventoryOrderConfig used by Order
    type AffiliateInventoryOrderConfig struct {

        
            ForceReassignment bool  `json:"force_reassignment"`
         
    }
    
    // AffiliateInventoryLogisticsConfig used by Order
    type AffiliateInventoryLogisticsConfig struct {

        
            DpAssignment bool  `json:"dp_assignment"`
         
    }
    
    // AffiliateInventoryPaymentConfig used by Order
    type AffiliateInventoryPaymentConfig struct {

        
            Source string  `json:"source"`
            ModeOfPayment string  `json:"mode_of_payment"`
         
    }
    
    // AffiliateInventoryArticleAssignmentConfig used by Order
    type AffiliateInventoryArticleAssignmentConfig struct {

        
            PostOrderReassignment bool  `json:"post_order_reassignment"`
         
    }
    
    // AffiliateInventoryStoreConfig used by Order
    type AffiliateInventoryStoreConfig struct {

        
            Store map[string]interface{}  `json:"store"`
         
    }
    
    // AffiliateInventoryConfig used by Order
    type AffiliateInventoryConfig struct {

        
            Order AffiliateInventoryOrderConfig  `json:"order"`
            Logistics AffiliateInventoryLogisticsConfig  `json:"logistics"`
            Payment AffiliateInventoryPaymentConfig  `json:"payment"`
            ArticleAssignment AffiliateInventoryArticleAssignmentConfig  `json:"article_assignment"`
            Inventory AffiliateInventoryStoreConfig  `json:"inventory"`
         
    }
    
    // AffiliateConfig used by Order
    type AffiliateConfig struct {

        
            App AffiliateAppConfig  `json:"app"`
            Inventory AffiliateInventoryConfig  `json:"inventory"`
         
    }
    
    // Affiliate used by Order
    type Affiliate struct {

        
            Config AffiliateConfig  `json:"config"`
            ID string  `json:"id"`
            Token string  `json:"token"`
         
    }
    
    // AffiliateStoreIdMapping used by Order
    type AffiliateStoreIdMapping struct {

        
            MarketplaceStoreID string  `json:"marketplace_store_id"`
            StoreID float64  `json:"store_id"`
         
    }
    
    // OrderConfig used by Order
    type OrderConfig struct {

        
            Affiliate Affiliate  `json:"affiliate"`
            ArticleLookup string  `json:"article_lookup"`
            CreateUser bool  `json:"create_user"`
            BagEndState string  `json:"bag_end_state"`
            AffiliateStoreIDMapping []AffiliateStoreIdMapping  `json:"affiliate_store_id_mapping"`
            StoreLookup string  `json:"store_lookup"`
         
    }
    
    // MarketPlacePdf used by Order
    type MarketPlacePdf struct {

        
            Invoice string  `json:"invoice"`
            Label string  `json:"label"`
         
    }
    
    // AffiliateBag used by Order
    type AffiliateBag struct {

        
            SellerIdentifier string  `json:"seller_identifier"`
            StoreID float64  `json:"store_id"`
            DeliveryCharge float64  `json:"delivery_charge"`
            FyndStoreID string  `json:"fynd_store_id"`
            UnitPrice float64  `json:"unit_price"`
            HsnCodeID string  `json:"hsn_code_id"`
            PriceEffective float64  `json:"price_effective"`
            Sku string  `json:"sku"`
            ModifiedOn string  `json:"modified_on"`
            Quantity float64  `json:"quantity"`
            ItemID float64  `json:"item_id"`
            AmountPaid float64  `json:"amount_paid"`
            TransferPrice float64  `json:"transfer_price"`
            Discount float64  `json:"discount"`
            Identifier map[string]interface{}  `json:"identifier"`
            AvlQty float64  `json:"avl_qty"`
            AffiliateMeta map[string]interface{}  `json:"affiliate_meta"`
            PdfLinks MarketPlacePdf  `json:"pdf_links"`
            CompanyID float64  `json:"company_id"`
            ItemSize string  `json:"item_size"`
            ID string  `json:"_id"`
            PriceMarked float64  `json:"price_marked"`
            AffiliateStoreID string  `json:"affiliate_store_id"`
         
    }
    
    // ArticleDetails1 used by Order
    type ArticleDetails1 struct {

        
            BrandID float64  `json:"brand_id"`
            ID string  `json:"_id"`
            Attributes map[string]interface{}  `json:"attributes"`
            Category map[string]interface{}  `json:"category"`
            Weight map[string]interface{}  `json:"weight"`
            Quantity float64  `json:"quantity"`
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

        
            Shipments float64  `json:"shipments"`
            Articles []ArticleDetails1  `json:"articles"`
            DpID float64  `json:"dp_id"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            FulfillmentID float64  `json:"fulfillment_id"`
            BoxType string  `json:"box_type"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // ShipmentConfig used by Order
    type ShipmentConfig struct {

        
            Journey string  `json:"journey"`
            PaymentMode string  `json:"payment_mode"`
            LocationDetails LocationDetails  `json:"location_details"`
            Source string  `json:"source"`
            Shipment []ShipmentDetails1  `json:"shipment"`
            ToPincode string  `json:"to_pincode"`
            Identifier string  `json:"identifier"`
            Action string  `json:"action"`
         
    }
    
    // ShipmentData used by Order
    type ShipmentData struct {

        
            ShipmentData ShipmentConfig  `json:"shipment_data"`
         
    }
    
    // OrderUser used by Order
    type OrderUser struct {

        
            Mobile float64  `json:"mobile"`
            LastName string  `json:"last_name"`
            Phone float64  `json:"phone"`
            Email string  `json:"email"`
            Pincode string  `json:"pincode"`
            Country string  `json:"country"`
            Address2 string  `json:"address2"`
            City string  `json:"city"`
            State string  `json:"state"`
            Address1 string  `json:"address1"`
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
    
    // OrderInfo used by Order
    type OrderInfo struct {

        
            AffiliateOrderID string  `json:"affiliate_order_id"`
            Bags []AffiliateBag  `json:"bags"`
            OrderValue float64  `json:"order_value"`
            Coupon string  `json:"coupon"`
            PaymentMode string  `json:"payment_mode"`
            Shipment ShipmentData  `json:"shipment"`
            Items map[string]interface{}  `json:"items"`
            Discount float64  `json:"discount"`
            ShippingAddress OrderUser  `json:"shipping_address"`
            DeliveryCharges float64  `json:"delivery_charges"`
            Payment map[string]interface{}  `json:"payment"`
            OrderPriority OrderPriority  `json:"order_priority"`
            User UserData  `json:"user"`
            BillingAddress OrderUser  `json:"billing_address"`
            CodCharges float64  `json:"cod_charges"`
         
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
            State string  `json:"state"`
            Category string  `json:"category"`
            Code float64  `json:"code"`
            Text string  `json:"text"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // HistoryMeta used by Order
    type HistoryMeta struct {

        
            Reason HistoryReason  `json:"reason"`
            StoreID float64  `json:"store_id"`
            Recordpath string  `json:"recordpath"`
            ChannelType string  `json:"channel_type"`
            Caller string  `json:"caller"`
            Billsec string  `json:"billsec"`
            StoreCode string  `json:"store_code"`
            Recipient string  `json:"recipient"`
            Slug string  `json:"slug"`
            Status2 string  `json:"status2"`
            Duration string  `json:"duration"`
            CallID string  `json:"call_id"`
            Status string  `json:"status"`
            ActivityComment string  `json:"activity_comment"`
            ActivityType string  `json:"activity_type"`
            Callerid string  `json:"callerid"`
            Starttime string  `json:"starttime"`
            Status1 string  `json:"status1"`
            Message string  `json:"message"`
            ShortLink string  `json:"short_link"`
            StoreName string  `json:"store_name"`
            Receiver string  `json:"receiver"`
            Endtime string  `json:"endtime"`
         
    }
    
    // HistoryDict used by Order
    type HistoryDict struct {

        
            Type string  `json:"type"`
            DisplayMessage string  `json:"display_message"`
            BagID float64  `json:"bag_id"`
            L1Detail string  `json:"l1_detail"`
            TicketURL string  `json:"ticket_url"`
            TicketID string  `json:"ticket_id"`
            Message string  `json:"message"`
            User string  `json:"user"`
            L2Detail string  `json:"l2_detail"`
            L3Detail string  `json:"l3_detail"`
            Createdat string  `json:"createdat"`
            AssignedAgent string  `json:"assigned_agent"`
            Meta HistoryMeta  `json:"meta"`
         
    }
    
    // ShipmentHistoryResponse used by Order
    type ShipmentHistoryResponse struct {

        
            ActivityHistory []HistoryDict  `json:"activity_history"`
            Success bool  `json:"success"`
         
    }
    
    // PostHistoryFilters used by Order
    type PostHistoryFilters struct {

        
            Identifier string  `json:"identifier"`
            LineNumber string  `json:"line_number"`
            ShipmentID string  `json:"shipment_id"`
         
    }
    
    // PostHistoryData used by Order
    type PostHistoryData struct {

        
            UserName string  `json:"user_name"`
            Message string  `json:"message"`
         
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

        
            CustomerName string  `json:"customer_name"`
            PaymentMode string  `json:"payment_mode"`
            AmountPaid float64  `json:"amount_paid"`
            OrderID string  `json:"order_id"`
            PhoneNumber float64  `json:"phone_number"`
            CountryCode string  `json:"country_code"`
            BrandName string  `json:"brand_name"`
            Message string  `json:"message"`
            ShipmentID float64  `json:"shipment_id"`
         
    }
    
    // SendSmsPayload used by Order
    type SendSmsPayload struct {

        
            BagID float64  `json:"bag_id"`
            Data SmsDataPayload  `json:"data"`
            Slug string  `json:"slug"`
         
    }
    
    // Meta used by Order
    type Meta struct {

        
            KafkaEmissionStatus float64  `json:"kafka_emission_status"`
            StateManagerUsed string  `json:"state_manager_used"`
         
    }
    
    // ShipmentDetail used by Order
    type ShipmentDetail struct {

        
            ID float64  `json:"id"`
            Remarks string  `json:"remarks"`
            BagList []float64  `json:"bag_list"`
            ShipmentID string  `json:"shipment_id"`
            Status string  `json:"status"`
            Meta Meta  `json:"meta"`
         
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

        
            Result []OrderStatusData  `json:"result"`
            Success string  `json:"success"`
         
    }
    
    // Dimension used by Order
    type Dimension struct {

        
            Height string  `json:"height"`
            PackagingType string  `json:"packaging_type"`
            Width float64  `json:"width"`
            Length float64  `json:"length"`
            Weight string  `json:"weight"`
         
    }
    
    // UpdatePackagingDimensionsPayload used by Order
    type UpdatePackagingDimensionsPayload struct {

        
            ShipmentID string  `json:"shipment_id"`
            Dimension []Dimension  `json:"dimension"`
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
    
    // LineItem used by Order
    type LineItem struct {

        
            ExternalLineID string  `json:"external_line_id"`
            CustomMessasge string  `json:"custom_messasge"`
            SellerIdentifier string  `json:"seller_identifier"`
            Charges []Charge  `json:"charges"`
            Quantity float64  `json:"quantity"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // ProcessingDates used by Order
    type ProcessingDates struct {

        
            ConfirmByDate string  `json:"confirm_by_date"`
            CustomerPickupSlot map[string]interface{}  `json:"customer_pickup_slot"`
            DpPickupSlot map[string]interface{}  `json:"dp_pickup_slot"`
            DispatchAfterDate string  `json:"dispatch_after_date"`
            PackByDate string  `json:"pack_by_date"`
            DispatchByDate string  `json:"dispatch_by_date"`
         
    }
    
    // Shipment used by Order
    type Shipment struct {

        
            Priority float64  `json:"priority"`
            ExternalShipmentID string  `json:"external_shipment_id"`
            LineItems []LineItem  `json:"line_items"`
            LocationID float64  `json:"location_id"`
            ProcessingDates ProcessingDates  `json:"processing_dates"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // PaymentMethod used by Order
    type PaymentMethod struct {

        
            TransactionData map[string]interface{}  `json:"transaction_data"`
            Name string  `json:"name"`
            RefundBy string  `json:"refund_by"`
            Amount float64  `json:"amount"`
            Mode string  `json:"mode"`
            CollectBy string  `json:"collect_by"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // PaymentInfo used by Order
    type PaymentInfo struct {

        
            PrimaryMode string  `json:"primary_mode"`
            PaymentMethods []PaymentMethod  `json:"payment_methods"`
         
    }
    
    // BillingInfo used by Order
    type BillingInfo struct {

        
            LastName string  `json:"last_name"`
            StateCode string  `json:"state_code"`
            FirstName string  `json:"first_name"`
            Gender string  `json:"gender"`
            Pincode string  `json:"pincode"`
            Country string  `json:"country"`
            ExternalCustomerCode string  `json:"external_customer_code"`
            PrimaryMobileNumber string  `json:"primary_mobile_number"`
            HouseNo string  `json:"house_no"`
            Title string  `json:"title"`
            FloorNo string  `json:"floor_no"`
            Address2 string  `json:"address2"`
            City string  `json:"city"`
            CustomerCode string  `json:"customer_code"`
            PrimaryEmail string  `json:"primary_email"`
            CountryCode string  `json:"country_code"`
            AlternateEmail string  `json:"alternate_email"`
            MiddleName string  `json:"middle_name"`
            State string  `json:"state"`
            Address1 string  `json:"address1"`
            AlternateMobileNumber string  `json:"alternate_mobile_number"`
         
    }
    
    // TaxInfo used by Order
    type TaxInfo struct {

        
            B2bGstinNumber string  `json:"b2b_gstin_number"`
            Gstin string  `json:"gstin"`
         
    }
    
    // ShippingInfo used by Order
    type ShippingInfo struct {

        
            LastName string  `json:"last_name"`
            StateCode string  `json:"state_code"`
            FirstName string  `json:"first_name"`
            Gender string  `json:"gender"`
            Landmark string  `json:"landmark"`
            Pincode string  `json:"pincode"`
            Country string  `json:"country"`
            ExternalCustomerCode string  `json:"external_customer_code"`
            PrimaryMobileNumber string  `json:"primary_mobile_number"`
            AddressType string  `json:"address_type"`
            HouseNo string  `json:"house_no"`
            Title string  `json:"title"`
            FloorNo string  `json:"floor_no"`
            Slot []map[string]interface{}  `json:"slot"`
            Address2 string  `json:"address2"`
            City string  `json:"city"`
            CustomerCode string  `json:"customer_code"`
            PrimaryEmail string  `json:"primary_email"`
            ShippingType string  `json:"shipping_type"`
            CountryCode string  `json:"country_code"`
            AlternateEmail string  `json:"alternate_email"`
            GeoLocation map[string]interface{}  `json:"geo_location"`
            MiddleName string  `json:"middle_name"`
            State string  `json:"state"`
            Address1 string  `json:"address1"`
            AlternateMobileNumber string  `json:"alternate_mobile_number"`
         
    }
    
    // CreateOrderAPI used by Order
    type CreateOrderAPI struct {

        
            Shipments []Shipment  `json:"shipments"`
            ExternalOrderID string  `json:"external_order_id"`
            PaymentInfo PaymentInfo  `json:"payment_info"`
            BillingInfo BillingInfo  `json:"billing_info"`
            ExternalCreationDate string  `json:"external_creation_date"`
            Charges []Charge  `json:"charges"`
            CurrencyInfo map[string]interface{}  `json:"currency_info"`
            Config map[string]interface{}  `json:"config"`
            TaxInfo TaxInfo  `json:"tax_info"`
            ShippingInfo ShippingInfo  `json:"shipping_info"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // CreateOrderErrorReponse used by Order
    type CreateOrderErrorReponse struct {

        
            RequestID string  `json:"request_id"`
            Info interface{}  `json:"info"`
            Code string  `json:"code"`
            Message string  `json:"message"`
            Exception string  `json:"exception"`
            Status float64  `json:"status"`
            StackTrace string  `json:"stack_trace"`
            Meta string  `json:"meta"`
         
    }
    
    // PaymentMethods used by Order
    type PaymentMethods struct {

        
            Mode string  `json:"mode"`
            RefundBy string  `json:"refund_by"`
            CollectBy string  `json:"collect_by"`
         
    }
    
    // CreateChannelPaymentInfo used by Order
    type CreateChannelPaymentInfo struct {

        
            ModeOfPayment string  `json:"mode_of_payment"`
            Source string  `json:"source"`
            PaymentMethods []PaymentMethods  `json:"payment_methods"`
         
    }
    
    // DpConfiguration used by Order
    type DpConfiguration struct {

        
            ShippingBy string  `json:"shipping_by"`
         
    }
    
    // CreateChannelConfig used by Order
    type CreateChannelConfig struct {

        
            LocationReassignment bool  `json:"location_reassignment"`
            PaymentInfo CreateChannelPaymentInfo  `json:"payment_info"`
            LogoURL map[string]interface{}  `json:"logo_url"`
            LockStates []string  `json:"lock_states"`
            ShipmentAssignment string  `json:"shipment_assignment"`
            DpConfiguration DpConfiguration  `json:"dp_configuration"`
         
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

        
            AppID string  `json:"app_id"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Words []string  `json:"words"`
            IsActive bool  `json:"is_active"`
            Result SearchKeywordResult  `json:"result"`
         
    }
    
    // GetSearchWordsData used by Catalog
    type GetSearchWordsData struct {

        
            AppID string  `json:"app_id"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Words []string  `json:"words"`
            UID string  `json:"uid"`
            IsActive bool  `json:"is_active"`
            Result map[string]interface{}  `json:"result"`
         
    }
    
    // ErrorResponse used by Catalog
    type ErrorResponse struct {

        
            Error string  `json:"error"`
            Status float64  `json:"status"`
            Message string  `json:"message"`
            Code string  `json:"code"`
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
    
    // AutoCompleteMedia used by Catalog
    type AutoCompleteMedia struct {

        
            AspectRatio string  `json:"aspect_ratio"`
            Type string  `json:"type"`
            URL string  `json:"url"`
         
    }
    
    // AutocompletePageAction used by Catalog
    type AutocompletePageAction struct {

        
            Params map[string]interface{}  `json:"params"`
            Query map[string]interface{}  `json:"query"`
            Type string  `json:"type"`
            URL string  `json:"url"`
         
    }
    
    // AutocompleteAction used by Catalog
    type AutocompleteAction struct {

        
            Page AutocompletePageAction  `json:"page"`
            Type string  `json:"type"`
         
    }
    
    // AutocompleteResult used by Catalog
    type AutocompleteResult struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Display string  `json:"display"`
            Logo AutoCompleteMedia  `json:"logo"`
            Action AutocompleteAction  `json:"action"`
         
    }
    
    // CreateAutocompleteKeyword used by Catalog
    type CreateAutocompleteKeyword struct {

        
            AppID string  `json:"app_id"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Results []AutocompleteResult  `json:"results"`
            Words []string  `json:"words"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // GetAutocompleteWordsData used by Catalog
    type GetAutocompleteWordsData struct {

        
            AppID string  `json:"app_id"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Results []map[string]interface{}  `json:"results"`
            UID string  `json:"uid"`
            Words []string  `json:"words"`
         
    }
    
    // GetAutocompleteWordsResponse used by Catalog
    type GetAutocompleteWordsResponse struct {

        
            Page Page  `json:"page"`
            Items []GetAutocompleteWordsData  `json:"items"`
         
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

        
            AutoAddToCart bool  `json:"auto_add_to_cart"`
            MaxQuantity float64  `json:"max_quantity"`
            AllowRemove bool  `json:"allow_remove"`
            AutoSelect bool  `json:"auto_select"`
            MinQuantity float64  `json:"min_quantity"`
            ProductUID float64  `json:"product_uid"`
         
    }
    
    // ProductBundleRequest used by Catalog
    type ProductBundleRequest struct {

        
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            CompanyID float64  `json:"company_id"`
            Logo string  `json:"logo"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            Slug string  `json:"slug"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Name string  `json:"name"`
            PageVisibility []string  `json:"page_visibility"`
            Products []ProductBundleItem  `json:"products"`
            ModifiedOn string  `json:"modified_on"`
            IsActive bool  `json:"is_active"`
            Choice string  `json:"choice"`
            CreatedOn string  `json:"created_on"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // GetProductBundleCreateResponse used by Catalog
    type GetProductBundleCreateResponse struct {

        
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            CompanyID float64  `json:"company_id"`
            Logo string  `json:"logo"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            Slug string  `json:"slug"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Name string  `json:"name"`
            PageVisibility []string  `json:"page_visibility"`
            Products []ProductBundleItem  `json:"products"`
            ID string  `json:"id"`
            ModifiedOn string  `json:"modified_on"`
            IsActive bool  `json:"is_active"`
            Choice string  `json:"choice"`
            CreatedOn string  `json:"created_on"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // GetProductBundleListingResponse used by Catalog
    type GetProductBundleListingResponse struct {

        
            Page Page  `json:"page"`
            Items []GetProductBundleCreateResponse  `json:"items"`
         
    }
    
    // ProductBundleUpdateRequest used by Catalog
    type ProductBundleUpdateRequest struct {

        
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            CompanyID float64  `json:"company_id"`
            Logo string  `json:"logo"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            Slug string  `json:"slug"`
            Name string  `json:"name"`
            PageVisibility []string  `json:"page_visibility"`
            Products []ProductBundleItem  `json:"products"`
            ModifiedOn string  `json:"modified_on"`
            IsActive bool  `json:"is_active"`
            Choice string  `json:"choice"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // Price used by Catalog
    type Price struct {

        
            MaxMarked float64  `json:"max_marked"`
            MaxEffective float64  `json:"max_effective"`
            MinEffective float64  `json:"min_effective"`
            Currency string  `json:"currency"`
            MinMarked float64  `json:"min_marked"`
         
    }
    
    // Size used by Catalog
    type Size struct {

        
            IsAvailable bool  `json:"is_available"`
            Display string  `json:"display"`
            Quantity float64  `json:"quantity"`
            Value string  `json:"value"`
         
    }
    
    // LimitedProductData used by Catalog
    type LimitedProductData struct {

        
            Price map[string]interface{}  `json:"price"`
            Identifier map[string]interface{}  `json:"identifier"`
            Images []string  `json:"images"`
            Quantity float64  `json:"quantity"`
            Slug string  `json:"slug"`
            ItemCode string  `json:"item_code"`
            Name string  `json:"name"`
            Sizes []string  `json:"sizes"`
            UID float64  `json:"uid"`
            ShortDescription string  `json:"short_description"`
            Attributes map[string]interface{}  `json:"attributes"`
            CountryOfOrigin string  `json:"country_of_origin"`
         
    }
    
    // GetProducts used by Catalog
    type GetProducts struct {

        
            AutoAddToCart bool  `json:"auto_add_to_cart"`
            Price Price  `json:"price"`
            Sizes []Size  `json:"sizes"`
            MaxQuantity float64  `json:"max_quantity"`
            AllowRemove bool  `json:"allow_remove"`
            MinQuantity float64  `json:"min_quantity"`
            AutoSelect bool  `json:"auto_select"`
            ProductDetails LimitedProductData  `json:"product_details"`
            ProductUID float64  `json:"product_uid"`
         
    }
    
    // GetProductBundleResponse used by Catalog
    type GetProductBundleResponse struct {

        
            CompanyID float64  `json:"company_id"`
            Logo string  `json:"logo"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            Slug string  `json:"slug"`
            Name string  `json:"name"`
            PageVisibility []string  `json:"page_visibility"`
            Products []GetProducts  `json:"products"`
            IsActive bool  `json:"is_active"`
            Choice string  `json:"choice"`
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

        
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            CompanyID float64  `json:"company_id"`
            Guide Guide  `json:"guide"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Title string  `json:"title"`
            Name string  `json:"name"`
            Description string  `json:"description"`
            ID string  `json:"id"`
            Image string  `json:"image"`
            CreatedOn string  `json:"created_on"`
            BrandID float64  `json:"brand_id"`
            Active bool  `json:"active"`
            ModifiedOn string  `json:"modified_on"`
            Tag string  `json:"tag"`
            Subtitle string  `json:"subtitle"`
         
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

        
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            CompanyID float64  `json:"company_id"`
            Tag string  `json:"tag"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Title string  `json:"title"`
            Name string  `json:"name"`
            ID string  `json:"id"`
            ModifiedOn string  `json:"modified_on"`
            BrandID float64  `json:"brand_id"`
            Active bool  `json:"active"`
            Subtitle string  `json:"subtitle"`
            CreatedOn string  `json:"created_on"`
            Guide map[string]interface{}  `json:"guide"`
         
    }
    
    // ApplicationItemMOQ used by Catalog
    type ApplicationItemMOQ struct {

        
            IncrementUnit float64  `json:"increment_unit"`
            Minimum float64  `json:"minimum"`
            Maximum float64  `json:"maximum"`
         
    }
    
    // MetaFields used by Catalog
    type MetaFields struct {

        
            Key interface{}  `json:"key"`
            Value interface{}  `json:"value"`
         
    }
    
    // ApplicationItemSEO used by Catalog
    type ApplicationItemSEO struct {

        
            Title interface{}  `json:"title"`
            Description interface{}  `json:"description"`
         
    }
    
    // ApplicationItemMeta used by Catalog
    type ApplicationItemMeta struct {

        
            Moq ApplicationItemMOQ  `json:"moq"`
            IsGift bool  `json:"is_gift"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            IsCod bool  `json:"is_cod"`
            CustomMeta []MetaFields  `json:"_custom_meta"`
            Seo ApplicationItemSEO  `json:"seo"`
            AltText map[string]interface{}  `json:"alt_text"`
         
    }
    
    // SuccessResponse1 used by Catalog
    type SuccessResponse1 struct {

        
            Success bool  `json:"success"`
            UID float64  `json:"uid"`
         
    }
    
    // MOQData used by Catalog
    type MOQData struct {

        
            IncrementUnit float64  `json:"increment_unit"`
            Minimum float64  `json:"minimum"`
            Maximum float64  `json:"maximum"`
         
    }
    
    // SEOData used by Catalog
    type SEOData struct {

        
            Title interface{}  `json:"title"`
            Description interface{}  `json:"description"`
         
    }
    
    // OwnerAppItemResponse used by Catalog
    type OwnerAppItemResponse struct {

        
            Moq MOQData  `json:"moq"`
            IsGift bool  `json:"is_gift"`
            IsCod bool  `json:"is_cod"`
            Seo SEOData  `json:"seo"`
            AltText map[string]interface{}  `json:"alt_text"`
         
    }
    
    // GetConfigMetadataResponse used by Catalog
    type GetConfigMetadataResponse struct {

        
            Condition []map[string]interface{}  `json:"condition"`
            Values []map[string]interface{}  `json:"values"`
            Data []map[string]interface{}  `json:"data"`
         
    }
    
    // AttributeDetailsGroup used by Catalog
    type AttributeDetailsGroup struct {

        
            Unit string  `json:"unit"`
            Logo string  `json:"logo"`
            Slug string  `json:"slug"`
            Name string  `json:"name"`
            Priority float64  `json:"priority"`
            Key string  `json:"key"`
            DisplayType string  `json:"display_type"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // AppConfigurationDetail used by Catalog
    type AppConfigurationDetail struct {

        
            AppID string  `json:"app_id"`
            Logo string  `json:"logo"`
            TemplateSlugs []string  `json:"template_slugs"`
            Slug string  `json:"slug"`
            Name string  `json:"name"`
            Priority float64  `json:"priority"`
            IsDefault bool  `json:"is_default"`
            Attributes []AttributeDetailsGroup  `json:"attributes"`
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

        
            AppID string  `json:"app_id"`
            Logo string  `json:"logo"`
            DefaultKey string  `json:"default_key"`
            Name string  `json:"name"`
            Priority float64  `json:"priority"`
            IsDefault bool  `json:"is_default"`
            Key string  `json:"key"`
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
    
    // MetaDataListingFilterMetaResponse used by Catalog
    type MetaDataListingFilterMetaResponse struct {

        
            Units []map[string]interface{}  `json:"units"`
            Key string  `json:"key"`
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
    
    // GetCatalogConfigurationDetailsProduct used by Catalog
    type GetCatalogConfigurationDetailsProduct struct {

        
            Compare map[string]interface{}  `json:"compare"`
            Variant map[string]interface{}  `json:"variant"`
            Similar map[string]interface{}  `json:"similar"`
            Detail map[string]interface{}  `json:"detail"`
         
    }
    
    // GetCatalogConfigurationMetaData used by Catalog
    type GetCatalogConfigurationMetaData struct {

        
            Listing MetaDataListingResponse  `json:"listing"`
            Product GetCatalogConfigurationDetailsProduct  `json:"product"`
         
    }
    
    // ProductSize used by Catalog
    type ProductSize struct {

        
            Min float64  `json:"min"`
            Max float64  `json:"max"`
         
    }
    
    // ConfigurationProductVariantConfig used by Catalog
    type ConfigurationProductVariantConfig struct {

        
            Logo string  `json:"logo"`
            Name string  `json:"name"`
            Priority float64  `json:"priority"`
            Size ProductSize  `json:"size"`
            Key string  `json:"key"`
            DisplayType string  `json:"display_type"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // ConfigurationProductVariant used by Catalog
    type ConfigurationProductVariant struct {

        
            Config []ConfigurationProductVariantConfig  `json:"config"`
         
    }
    
    // ConfigurationProductConfig used by Catalog
    type ConfigurationProductConfig struct {

        
            Logo string  `json:"logo"`
            Title string  `json:"title"`
            Priority float64  `json:"priority"`
            Size ProductSize  `json:"size"`
            Key string  `json:"key"`
            IsActive bool  `json:"is_active"`
            Subtitle string  `json:"subtitle"`
         
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
    
    // ConfigurationBucketPoints used by Catalog
    type ConfigurationBucketPoints struct {

        
            Display string  `json:"display"`
            Start float64  `json:"start"`
            End float64  `json:"end"`
         
    }
    
    // ConfigurationListingFilterValue used by Catalog
    type ConfigurationListingFilterValue struct {

        
            Map map[string]interface{}  `json:"map"`
            Value string  `json:"value"`
            MapValues []map[string]interface{}  `json:"map_values"`
            Condition string  `json:"condition"`
            BucketPoints []ConfigurationBucketPoints  `json:"bucket_points"`
            Sort string  `json:"sort"`
         
    }
    
    // ConfigurationListingFilterConfig used by Catalog
    type ConfigurationListingFilterConfig struct {

        
            ValueConfig ConfigurationListingFilterValue  `json:"value_config"`
            Logo string  `json:"logo"`
            Type string  `json:"type"`
            Name string  `json:"name"`
            Priority float64  `json:"priority"`
            Key string  `json:"key"`
            DisplayName string  `json:"display_name"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // ConfigurationListingFilter used by Catalog
    type ConfigurationListingFilter struct {

        
            AttributeConfig []ConfigurationListingFilterConfig  `json:"attribute_config"`
            AllowSingle bool  `json:"allow_single"`
         
    }
    
    // ConfigurationListingSortConfig used by Catalog
    type ConfigurationListingSortConfig struct {

        
            Logo string  `json:"logo"`
            Name string  `json:"name"`
            Priority float64  `json:"priority"`
            Key string  `json:"key"`
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

        
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            AppID string  `json:"app_id"`
            Type string  `json:"type"`
            Product ConfigurationProduct  `json:"product"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Listing ConfigurationListing  `json:"listing"`
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
            ConfigType string  `json:"config_type"`
            ConfigID string  `json:"config_id"`
         
    }
    
    // AppCatalogConfiguration used by Catalog
    type AppCatalogConfiguration struct {

        
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            AppID string  `json:"app_id"`
            Type string  `json:"type"`
            Product ConfigurationProduct  `json:"product"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Listing ConfigurationListing  `json:"listing"`
            ID string  `json:"id"`
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
            ConfigType string  `json:"config_type"`
            ConfigID string  `json:"config_id"`
         
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

        
            AppID string  `json:"app_id"`
            Product GetCatalogConfigurationDetailsProduct  `json:"product"`
            Listing GetCatalogConfigurationDetailsSchemaListing  `json:"listing"`
            ID string  `json:"id"`
            ConfigType string  `json:"config_type"`
            ConfigID string  `json:"config_id"`
         
    }
    
    // GetAppCatalogEntityConfiguration used by Catalog
    type GetAppCatalogEntityConfiguration struct {

        
            Data EntityConfiguration  `json:"data"`
            IsDefault bool  `json:"is_default"`
         
    }
    
    // ProductFiltersValue used by Catalog
    type ProductFiltersValue struct {

        
            QueryFormat string  `json:"query_format"`
            CurrencyCode string  `json:"currency_code"`
            DisplayFormat string  `json:"display_format"`
            Max float64  `json:"max"`
            Value interface{}  `json:"value"`
            Count float64  `json:"count"`
            Display string  `json:"display"`
            SelectedMin float64  `json:"selected_min"`
            CurrencySymbol string  `json:"currency_symbol"`
            Min float64  `json:"min"`
            IsSelected bool  `json:"is_selected"`
            SelectedMax float64  `json:"selected_max"`
         
    }
    
    // ProductFiltersKey used by Catalog
    type ProductFiltersKey struct {

        
            Logo string  `json:"logo"`
            Operators []string  `json:"operators"`
            Name string  `json:"name"`
            Display string  `json:"display"`
            Kind string  `json:"kind"`
         
    }
    
    // ProductFilters used by Catalog
    type ProductFilters struct {

        
            Values []ProductFiltersValue  `json:"values"`
            Key ProductFiltersKey  `json:"key"`
         
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
    
    // UserInfo used by Catalog
    type UserInfo struct {

        
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
            Email string  `json:"email"`
            UID string  `json:"uid"`
         
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
    
    // CollectionQuery used by Catalog
    type CollectionQuery struct {

        
            Op string  `json:"op"`
            Attribute string  `json:"attribute"`
            Value []interface{}  `json:"value"`
         
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

        
            End string  `json:"end"`
            NextSchedule []NextSchedule  `json:"next_schedule"`
            Cron string  `json:"cron"`
            Duration float64  `json:"duration"`
            Start string  `json:"start"`
         
    }
    
    // SeoDetail used by Catalog
    type SeoDetail struct {

        
            Title string  `json:"title"`
            Description string  `json:"description"`
         
    }
    
    // CreateCollection used by Catalog
    type CreateCollection struct {

        
            CreatedBy UserInfo  `json:"created_by"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            IsActive bool  `json:"is_active"`
            Published bool  `json:"published"`
            Banners CollectionBanner  `json:"banners"`
            Query []CollectionQuery  `json:"query"`
            Badge CollectionBadge  `json:"badge"`
            Schedule CollectionSchedule  `json:"_schedule"`
            AllowFacets bool  `json:"allow_facets"`
            AppID string  `json:"app_id"`
            Type string  `json:"type"`
            Slug string  `json:"slug"`
            SortOn string  `json:"sort_on"`
            Name string  `json:"name"`
            Description string  `json:"description"`
            Priority float64  `json:"priority"`
            IsVisible bool  `json:"is_visible"`
            AllowSort bool  `json:"allow_sort"`
            Seo SeoDetail  `json:"seo"`
            Meta map[string]interface{}  `json:"meta"`
            ModifiedBy UserInfo  `json:"modified_by"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            Logo CollectionImage  `json:"logo"`
            Tags []string  `json:"tags"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
         
    }
    
    // BannerImage used by Catalog
    type BannerImage struct {

        
            AspectRatio string  `json:"aspect_ratio"`
            Type string  `json:"type"`
            URL string  `json:"url"`
         
    }
    
    // ImageUrls used by Catalog
    type ImageUrls struct {

        
            Landscape BannerImage  `json:"landscape"`
            Portrait BannerImage  `json:"portrait"`
         
    }
    
    // CollectionCreateResponse used by Catalog
    type CollectionCreateResponse struct {

        
            IsActive bool  `json:"is_active"`
            Banners ImageUrls  `json:"banners"`
            Query []CollectionQuery  `json:"query"`
            Badge map[string]interface{}  `json:"badge"`
            Schedule map[string]interface{}  `json:"_schedule"`
            AllowFacets bool  `json:"allow_facets"`
            AppID string  `json:"app_id"`
            Type string  `json:"type"`
            Slug string  `json:"slug"`
            SortOn string  `json:"sort_on"`
            Name string  `json:"name"`
            Description string  `json:"description"`
            Priority float64  `json:"priority"`
            Cron map[string]interface{}  `json:"cron"`
            AllowSort bool  `json:"allow_sort"`
            Meta map[string]interface{}  `json:"meta"`
            Logo BannerImage  `json:"logo"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Tag []string  `json:"tag"`
         
    }
    
    // Media used by Catalog
    type Media struct {

        
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

        
            UID string  `json:"uid"`
            IsActive bool  `json:"is_active"`
            Banners ImageUrls  `json:"banners"`
            Query []CollectionQuery  `json:"query"`
            Badge map[string]interface{}  `json:"badge"`
            Schedule map[string]interface{}  `json:"_schedule"`
            AllowFacets bool  `json:"allow_facets"`
            AppID string  `json:"app_id"`
            Type string  `json:"type"`
            Slug string  `json:"slug"`
            Name string  `json:"name"`
            Description string  `json:"description"`
            Priority float64  `json:"priority"`
            Cron map[string]interface{}  `json:"cron"`
            AllowSort bool  `json:"allow_sort"`
            Meta map[string]interface{}  `json:"meta"`
            Logo Media  `json:"logo"`
            Action Action  `json:"action"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Tag []string  `json:"tag"`
         
    }
    
    // CollectionListingFilterTag used by Catalog
    type CollectionListingFilterTag struct {

        
            Name string  `json:"name"`
            IsSelected bool  `json:"is_selected"`
            Display string  `json:"display"`
         
    }
    
    // CollectionListingFilterType used by Catalog
    type CollectionListingFilterType struct {

        
            Name string  `json:"name"`
            IsSelected bool  `json:"is_selected"`
            Display string  `json:"display"`
         
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

        
            UID string  `json:"uid"`
            IsActive bool  `json:"is_active"`
            Banners ImageUrls  `json:"banners"`
            Query []CollectionQuery  `json:"query"`
            Badge map[string]interface{}  `json:"badge"`
            Schedule map[string]interface{}  `json:"_schedule"`
            AllowFacets bool  `json:"allow_facets"`
            AppID string  `json:"app_id"`
            Type string  `json:"type"`
            Slug string  `json:"slug"`
            Name string  `json:"name"`
            Description string  `json:"description"`
            Priority float64  `json:"priority"`
            Cron map[string]interface{}  `json:"cron"`
            AllowSort bool  `json:"allow_sort"`
            Meta map[string]interface{}  `json:"meta"`
            Logo Media  `json:"logo"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Tag []string  `json:"tag"`
         
    }
    
    // UpdateCollection used by Catalog
    type UpdateCollection struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            IsActive bool  `json:"is_active"`
            Published bool  `json:"published"`
            Banners CollectionBanner  `json:"banners"`
            Query []CollectionQuery  `json:"query"`
            Badge CollectionBadge  `json:"badge"`
            Schedule CollectionSchedule  `json:"_schedule"`
            AllowFacets bool  `json:"allow_facets"`
            Type string  `json:"type"`
            Slug string  `json:"slug"`
            SortOn string  `json:"sort_on"`
            Name string  `json:"name"`
            Description string  `json:"description"`
            Priority float64  `json:"priority"`
            IsVisible bool  `json:"is_visible"`
            AllowSort bool  `json:"allow_sort"`
            Seo SeoDetail  `json:"seo"`
            Meta map[string]interface{}  `json:"meta"`
            ModifiedBy UserInfo  `json:"modified_by"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            Logo CollectionImage  `json:"logo"`
            Tags []string  `json:"tags"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
         
    }
    
    // CollectionItem used by Catalog
    type CollectionItem struct {

        
            Priority float64  `json:"priority"`
            Action string  `json:"action"`
            ItemID float64  `json:"item_id"`
         
    }
    
    // CollectionItemUpdate used by Catalog
    type CollectionItemUpdate struct {

        
            Query []CollectionQuery  `json:"query"`
            Type string  `json:"type"`
            AllowFacets bool  `json:"allow_facets"`
            AllowSort bool  `json:"allow_sort"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Items []CollectionItem  `json:"items"`
         
    }
    
    // UpdatedResponse used by Catalog
    type UpdatedResponse struct {

        
            ItemsNotUpdated []float64  `json:"items_not_updated"`
            Message string  `json:"message"`
         
    }
    
    // ProductBrand used by Catalog
    type ProductBrand struct {

        
            Name string  `json:"name"`
            Logo Media  `json:"logo"`
            UID float64  `json:"uid"`
            Action Action  `json:"action"`
         
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

        
            CurrencyCode string  `json:"currency_code"`
            Min float64  `json:"min"`
            CurrencySymbol string  `json:"currency_symbol"`
            Max float64  `json:"max"`
         
    }
    
    // ProductListingPrice used by Catalog
    type ProductListingPrice struct {

        
            Effective Price1  `json:"effective"`
            Marked Price1  `json:"marked"`
         
    }
    
    // ProductListingDetail used by Catalog
    type ProductListingDetail struct {

        
            PromoMeta map[string]interface{}  `json:"promo_meta"`
            ItemType string  `json:"item_type"`
            Highlights []string  `json:"highlights"`
            UID float64  `json:"uid"`
            Brand ProductBrand  `json:"brand"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
            ProductOnlineDate string  `json:"product_online_date"`
            Medias []Media  `json:"medias"`
            Attributes map[string]interface{}  `json:"attributes"`
            Type string  `json:"type"`
            Slug string  `json:"slug"`
            Name string  `json:"name"`
            Description string  `json:"description"`
            Discount string  `json:"discount"`
            Color string  `json:"color"`
            Rating float64  `json:"rating"`
            Similars []string  `json:"similars"`
            HasVariant bool  `json:"has_variant"`
            Tryouts []string  `json:"tryouts"`
            Price ProductListingPrice  `json:"price"`
            ImageNature string  `json:"image_nature"`
            ItemCode string  `json:"item_code"`
            Sellable bool  `json:"sellable"`
            ShortDescription string  `json:"short_description"`
            RatingCount float64  `json:"rating_count"`
         
    }
    
    // GetCollectionItemsResponse used by Catalog
    type GetCollectionItemsResponse struct {

        
            Page Page  `json:"page"`
            Items []ProductListingDetail  `json:"items"`
            Filters []ProductFilters  `json:"filters"`
            SortOn []ProductSortOn  `json:"sort_on"`
         
    }
    
    // CatalogInsightItem used by Catalog
    type CatalogInsightItem struct {

        
            Count float64  `json:"count"`
            SellableCount float64  `json:"sellable_count"`
            OutOfStockCount float64  `json:"out_of_stock_count"`
         
    }
    
    // CatalogInsightBrand used by Catalog
    type CatalogInsightBrand struct {

        
            AvailableSizes float64  `json:"available_sizes"`
            Name string  `json:"name"`
            AvailableArticles float64  `json:"available_articles"`
            TotalArticles float64  `json:"total_articles"`
            TotalSizes float64  `json:"total_sizes"`
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

        
            CompanyID float64  `json:"company_id"`
            Platform string  `json:"platform"`
            StoreIds []float64  `json:"store_ids"`
            OptLevel string  `json:"opt_level"`
            Enabled bool  `json:"enabled"`
            BrandIds []float64  `json:"brand_ids"`
         
    }
    
    // CompanyOptIn used by Catalog
    type CompanyOptIn struct {

        
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            CompanyID float64  `json:"company_id"`
            Platform string  `json:"platform"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            StoreIds []float64  `json:"store_ids"`
            ModifiedOn float64  `json:"modified_on"`
            OptLevel string  `json:"opt_level"`
            Enabled bool  `json:"enabled"`
            CreatedOn float64  `json:"created_on"`
            BrandIds []float64  `json:"brand_ids"`
         
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
            BusinessType string  `json:"business_type"`
            UID float64  `json:"uid"`
         
    }
    
    // CompanyBrandDetail used by Catalog
    type CompanyBrandDetail struct {

        
            CompanyID float64  `json:"company_id"`
            TotalArticle float64  `json:"total_article"`
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

        
            Store float64  `json:"store"`
            Company string  `json:"company"`
            Brand float64  `json:"brand"`
         
    }
    
    // StoreDetail used by Catalog
    type StoreDetail struct {

        
            CompanyID float64  `json:"company_id"`
            AdditionalContacts []map[string]interface{}  `json:"additional_contacts"`
            Name string  `json:"name"`
            StoreCode string  `json:"store_code"`
            Documents []map[string]interface{}  `json:"documents"`
            Address map[string]interface{}  `json:"address"`
            Timing map[string]interface{}  `json:"timing"`
            ModifiedOn string  `json:"modified_on"`
            UID float64  `json:"uid"`
            Manager map[string]interface{}  `json:"manager"`
            DisplayName string  `json:"display_name"`
            StoreType string  `json:"store_type"`
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

        
            Min float64  `json:"min"`
            Max float64  `json:"max"`
         
    }
    
    // AttributeMaster used by Catalog
    type AttributeMaster struct {

        
            Type string  `json:"type"`
            Multi bool  `json:"multi"`
            Format string  `json:"format"`
            Range AttributeSchemaRange  `json:"range"`
            AllowedValues []string  `json:"allowed_values"`
            Mandatory bool  `json:"mandatory"`
         
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

        
            Details AttributeMasterDetails  `json:"details"`
            Logo string  `json:"logo"`
            Slug string  `json:"slug"`
            Name string  `json:"name"`
            Description string  `json:"description"`
            ID string  `json:"id"`
            Filters AttributeMasterFilter  `json:"filters"`
            Departments []string  `json:"departments"`
            Schema AttributeMaster  `json:"schema"`
            IsNested bool  `json:"is_nested"`
            Meta AttributeMasterMeta  `json:"meta"`
            EnabledForEndConsumer bool  `json:"enabled_for_end_consumer"`
         
    }
    
    // CategoriesResponse used by Catalog
    type CategoriesResponse struct {

        
            Slug string  `json:"slug"`
            TemplateSlug string  `json:"template_slug"`
            Name string  `json:"name"`
            UID float64  `json:"uid"`
            SlugKey string  `json:"slug_key"`
         
    }
    
    // ProdcutTemplateCategoriesResponse used by Catalog
    type ProdcutTemplateCategoriesResponse struct {

        
            Page Page  `json:"page"`
            Items []CategoriesResponse  `json:"items"`
         
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

        
            Logo string  `json:"logo"`
            Slug string  `json:"slug"`
            Name string  `json:"name"`
            Tags []string  `json:"tags"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            PriorityOrder float64  `json:"priority_order"`
            Platforms map[string]interface{}  `json:"platforms"`
            Cls string  `json:"_cls"`
            UID float64  `json:"uid"`
            IsActive bool  `json:"is_active"`
            Synonyms []string  `json:"synonyms"`
         
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

        
            Username string  `json:"username"`
            Contact string  `json:"contact"`
            UserID string  `json:"user_id"`
            ID string  `json:"_id"`
            UID string  `json:"uid"`
         
    }
    
    // GetDepartment used by Catalog
    type GetDepartment struct {

        
            ModifiedBy UserSerializer  `json:"modified_by"`
            Logo string  `json:"logo"`
            CreatedBy UserSerializer  `json:"created_by"`
            Slug string  `json:"slug"`
            PageSize float64  `json:"page_size"`
            Name string  `json:"name"`
            ItemType string  `json:"item_type"`
            PriorityOrder float64  `json:"priority_order"`
            PageNo float64  `json:"page_no"`
            ModifiedOn string  `json:"modified_on"`
            UID float64  `json:"uid"`
            Search string  `json:"search"`
            IsActive bool  `json:"is_active"`
            CreatedOn string  `json:"created_on"`
            Synonyms []string  `json:"synonyms"`
         
    }
    
    // DepartmentsResponse used by Catalog
    type DepartmentsResponse struct {

        
            Page Page  `json:"page"`
            Items []GetDepartment  `json:"items"`
         
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
            Username string  `json:"username"`
            SuperUser bool  `json:"super_user"`
            Contact string  `json:"contact"`
         
    }
    
    // DepartmentModel used by Catalog
    type DepartmentModel struct {

        
            ModifiedBy UserDetail  `json:"modified_by"`
            Logo interface{}  `json:"logo"`
            CreatedBy UserDetail  `json:"created_by"`
            Slug interface{}  `json:"slug"`
            Name interface{}  `json:"name"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            VerifiedOn string  `json:"verified_on"`
            PriorityOrder float64  `json:"priority_order"`
            ModifiedOn string  `json:"modified_on"`
            ID interface{}  `json:"_id"`
            UID float64  `json:"uid"`
            Cls interface{}  `json:"_cls"`
            VerifiedBy UserDetail  `json:"verified_by"`
            IsActive bool  `json:"is_active"`
            CreatedOn string  `json:"created_on"`
            Synonyms []interface{}  `json:"synonyms"`
         
    }
    
    // ProductTemplate used by Catalog
    type ProductTemplate struct {

        
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Logo string  `json:"logo"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Slug string  `json:"slug"`
            Name string  `json:"name"`
            IsArchived bool  `json:"is_archived"`
            Categories []string  `json:"categories"`
            Description string  `json:"description"`
            Departments []string  `json:"departments"`
            IsPhysical bool  `json:"is_physical"`
            CreatedOn string  `json:"created_on"`
            IsExpirable bool  `json:"is_expirable"`
            ModifiedOn string  `json:"modified_on"`
            Attributes []string  `json:"attributes"`
            IsActive bool  `json:"is_active"`
            Tag string  `json:"tag"`
         
    }
    
    // TemplatesResponse used by Catalog
    type TemplatesResponse struct {

        
            Page Page  `json:"page"`
            Items ProductTemplate  `json:"items"`
         
    }
    
    // Properties used by Catalog
    type Properties struct {

        
            ReturnConfig map[string]interface{}  `json:"return_config"`
            ItemType map[string]interface{}  `json:"item_type"`
            Highlights map[string]interface{}  `json:"highlights"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            IsActive map[string]interface{}  `json:"is_active"`
            Trader map[string]interface{}  `json:"trader"`
            ProductGroupTag map[string]interface{}  `json:"product_group_tag"`
            HsnCode map[string]interface{}  `json:"hsn_code"`
            TraderType map[string]interface{}  `json:"trader_type"`
            ProductPublish map[string]interface{}  `json:"product_publish"`
            Currency map[string]interface{}  `json:"currency"`
            SizeGuide map[string]interface{}  `json:"size_guide"`
            Slug map[string]interface{}  `json:"slug"`
            Name map[string]interface{}  `json:"name"`
            Description map[string]interface{}  `json:"description"`
            NoOfBoxes map[string]interface{}  `json:"no_of_boxes"`
            Media map[string]interface{}  `json:"media"`
            Sizes map[string]interface{}  `json:"sizes"`
            Command map[string]interface{}  `json:"command"`
            IsDependent map[string]interface{}  `json:"is_dependent"`
            BrandUID map[string]interface{}  `json:"brand_uid"`
            MultiSize map[string]interface{}  `json:"multi_size"`
            ItemCode map[string]interface{}  `json:"item_code"`
            Tags map[string]interface{}  `json:"tags"`
            ShortDescription map[string]interface{}  `json:"short_description"`
            Variants map[string]interface{}  `json:"variants"`
            CountryOfOrigin map[string]interface{}  `json:"country_of_origin"`
            CategorySlug map[string]interface{}  `json:"category_slug"`
         
    }
    
    // GlobalValidation used by Catalog
    type GlobalValidation struct {

        
            Type string  `json:"type"`
            Title string  `json:"title"`
            Properties Properties  `json:"properties"`
            Description string  `json:"description"`
            Required []string  `json:"required"`
            Definitions map[string]interface{}  `json:"definitions"`
         
    }
    
    // TemplateValidationData used by Catalog
    type TemplateValidationData struct {

        
            TemplateValidation map[string]interface{}  `json:"template_validation"`
            GlobalValidation GlobalValidation  `json:"global_validation"`
         
    }
    
    // TemplateDetails used by Catalog
    type TemplateDetails struct {

        
            Logo string  `json:"logo"`
            Slug string  `json:"slug"`
            Name string  `json:"name"`
            IsArchived bool  `json:"is_archived"`
            Categories []string  `json:"categories"`
            Description string  `json:"description"`
            ID string  `json:"id"`
            Departments []string  `json:"departments"`
            IsPhysical bool  `json:"is_physical"`
            IsExpirable bool  `json:"is_expirable"`
            Attributes []string  `json:"attributes"`
            IsActive bool  `json:"is_active"`
            Tag string  `json:"tag"`
         
    }
    
    // TemplatesValidationResponse used by Catalog
    type TemplatesValidationResponse struct {

        
            Data TemplateValidationData  `json:"data"`
            TemplateDetails TemplateDetails  `json:"template_details"`
         
    }
    
    // InventoryValidationResponse used by Catalog
    type InventoryValidationResponse struct {

        
            Message string  `json:"message"`
            Data map[string]interface{}  `json:"data"`
         
    }
    
    // HSNData used by Catalog
    type HSNData struct {

        
            HsnCode []string  `json:"hsn_code"`
            CountryOfOrigin []string  `json:"country_of_origin"`
         
    }
    
    // HSNCodesResponse used by Catalog
    type HSNCodesResponse struct {

        
            Message string  `json:"message"`
            Data HSNData  `json:"data"`
         
    }
    
    // UserInfo1 used by Catalog
    type UserInfo1 struct {

        
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
            Email string  `json:"email"`
            UID string  `json:"uid"`
         
    }
    
    // ProductTemplateExportResponse used by Catalog
    type ProductTemplateExportResponse struct {

        
            NotificationEmails []string  `json:"notification_emails"`
            Type string  `json:"type"`
            CreatedBy UserInfo1  `json:"created_by"`
            TaskID string  `json:"task_id"`
            Status string  `json:"status"`
            SellerID float64  `json:"seller_id"`
            ModifiedOn string  `json:"modified_on"`
            CompletedOn string  `json:"completed_on"`
            Filters map[string]interface{}  `json:"filters"`
            URL string  `json:"url"`
         
    }
    
    // ProductDownloadsResponse used by Catalog
    type ProductDownloadsResponse struct {

        
            Items []ProductTemplateExportResponse  `json:"items"`
         
    }
    
    // ProductTemplateExportFilterRequest used by Catalog
    type ProductTemplateExportFilterRequest struct {

        
            ToDate string  `json:"to_date"`
            FromDate string  `json:"from_date"`
            Brands []string  `json:"brands"`
            Templates []string  `json:"templates"`
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
    
    // CategoryMappingValues used by Catalog
    type CategoryMappingValues struct {

        
            Name string  `json:"name"`
            CatalogID float64  `json:"catalog_id"`
         
    }
    
    // CategoryMapping used by Catalog
    type CategoryMapping struct {

        
            Ajio CategoryMappingValues  `json:"ajio"`
            Facebook CategoryMappingValues  `json:"facebook"`
            Google CategoryMappingValues  `json:"google"`
         
    }
    
    // Hierarchy used by Catalog
    type Hierarchy struct {

        
            L1 float64  `json:"l1"`
            Department float64  `json:"department"`
            L2 float64  `json:"l2"`
         
    }
    
    // Media1 used by Catalog
    type Media1 struct {

        
            Landscape string  `json:"landscape"`
            Portrait string  `json:"portrait"`
            Logo string  `json:"logo"`
         
    }
    
    // CategoryRequestBody used by Catalog
    type CategoryRequestBody struct {

        
            Slug string  `json:"slug"`
            Name string  `json:"name"`
            Marketplaces CategoryMapping  `json:"marketplaces"`
            Priority float64  `json:"priority"`
            Synonyms []string  `json:"synonyms"`
            Hierarchy []Hierarchy  `json:"hierarchy"`
            Media Media1  `json:"media"`
            Level float64  `json:"level"`
            Departments []float64  `json:"departments"`
            IsActive bool  `json:"is_active"`
            Tryouts []string  `json:"tryouts"`
         
    }
    
    // CategoryCreateResponse used by Catalog
    type CategoryCreateResponse struct {

        
            Message string  `json:"message"`
            UID float64  `json:"uid"`
         
    }
    
    // Category used by Catalog
    type Category struct {

        
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Slug string  `json:"slug"`
            Name string  `json:"name"`
            Marketplaces CategoryMapping  `json:"marketplaces"`
            Priority float64  `json:"priority"`
            Synonyms []string  `json:"synonyms"`
            Hierarchy []Hierarchy  `json:"hierarchy"`
            Level float64  `json:"level"`
            Media Media1  `json:"media"`
            ID string  `json:"id"`
            Departments []float64  `json:"departments"`
            ModifiedOn string  `json:"modified_on"`
            UID float64  `json:"uid"`
            IsActive bool  `json:"is_active"`
            CreatedOn string  `json:"created_on"`
            Tryouts []string  `json:"tryouts"`
         
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
    
    // ReturnConfig used by Catalog
    type ReturnConfig struct {

        
            Time float64  `json:"time"`
            Unit string  `json:"unit"`
            Returnable bool  `json:"returnable"`
         
    }
    
    // TeaserTag used by Catalog
    type TeaserTag struct {

        
            Tag string  `json:"tag"`
            URL string  `json:"url"`
         
    }
    
    // CustomOrder used by Catalog
    type CustomOrder struct {

        
            ManufacturingTimeUnit string  `json:"manufacturing_time_unit"`
            ManufacturingTime float64  `json:"manufacturing_time"`
            IsCustomOrder bool  `json:"is_custom_order"`
         
    }
    
    // Trader used by Catalog
    type Trader struct {

        
            Address []string  `json:"address"`
            Name interface{}  `json:"name"`
            Type string  `json:"type"`
         
    }
    
    // TaxIdentifier used by Catalog
    type TaxIdentifier struct {

        
            HsnCodeID string  `json:"hsn_code_id"`
            HsnCode string  `json:"hsn_code"`
            ReportingHsn string  `json:"reporting_hsn"`
         
    }
    
    // ProductPublish used by Catalog
    type ProductPublish struct {

        
            ProductOnlineDate string  `json:"product_online_date"`
            IsSet bool  `json:"is_set"`
         
    }
    
    // NetQuantity used by Catalog
    type NetQuantity struct {

        
            Unit interface{}  `json:"unit"`
            Value float64  `json:"value"`
         
    }
    
    // ProductCreateUpdateSchemaV2 used by Catalog
    type ProductCreateUpdateSchemaV2 struct {

        
            CompanyID float64  `json:"company_id"`
            ReturnConfig ReturnConfig  `json:"return_config"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            TemplateTag string  `json:"template_tag"`
            ItemType string  `json:"item_type"`
            Highlights []string  `json:"highlights"`
            UID float64  `json:"uid"`
            TeaserTag TeaserTag  `json:"teaser_tag"`
            CustomOrder CustomOrder  `json:"custom_order"`
            IsActive bool  `json:"is_active"`
            Trader []Trader  `json:"trader"`
            ProductGroupTag []string  `json:"product_group_tag"`
            TaxIdentifier TaxIdentifier  `json:"tax_identifier"`
            IsImageLessProduct bool  `json:"is_image_less_product"`
            BulkJobID string  `json:"bulk_job_id"`
            ProductPublish ProductPublish  `json:"product_publish"`
            Currency string  `json:"currency"`
            ChangeRequestID interface{}  `json:"change_request_id"`
            Requester string  `json:"requester"`
            Attributes map[string]interface{}  `json:"attributes"`
            NetQuantity NetQuantity  `json:"net_quantity"`
            SizeGuide string  `json:"size_guide"`
            Slug string  `json:"slug"`
            Name string  `json:"name"`
            Description string  `json:"description"`
            NoOfBoxes float64  `json:"no_of_boxes"`
            IsSet bool  `json:"is_set"`
            Media []Media  `json:"media"`
            Sizes []map[string]interface{}  `json:"sizes"`
            VariantGroup map[string]interface{}  `json:"variant_group"`
            IsDependent bool  `json:"is_dependent"`
            BrandUID float64  `json:"brand_uid"`
            MultiSize bool  `json:"multi_size"`
            ItemCode string  `json:"item_code"`
            Tags []string  `json:"tags"`
            VariantMedia map[string]interface{}  `json:"variant_media"`
            Departments []float64  `json:"departments"`
            Action string  `json:"action"`
            Variants map[string]interface{}  `json:"variants"`
            ShortDescription string  `json:"short_description"`
            CountryOfOrigin string  `json:"country_of_origin"`
            CategorySlug string  `json:"category_slug"`
         
    }
    
    // ReturnConfigResponse used by Catalog
    type ReturnConfigResponse struct {

        
            Time float64  `json:"time"`
            Unit string  `json:"unit"`
            Returnable bool  `json:"returnable"`
         
    }
    
    // Logo used by Catalog
    type Logo struct {

        
            AspectRatio string  `json:"aspect_ratio"`
            SecureURL string  `json:"secure_url"`
            AspectRatioF float64  `json:"aspect_ratio_f"`
            URL string  `json:"url"`
         
    }
    
    // Brand used by Catalog
    type Brand struct {

        
            Name string  `json:"name"`
            Logo Logo  `json:"logo"`
            UID float64  `json:"uid"`
         
    }
    
    // ProductPublish1 used by Catalog
    type ProductPublish1 struct {

        
            ProductOnlineDate string  `json:"product_online_date"`
            IsSet bool  `json:"is_set"`
         
    }
    
    // VerifiedBy used by Catalog
    type VerifiedBy struct {

        
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
         
    }
    
    // NetQuantityResponse used by Catalog
    type NetQuantityResponse struct {

        
            Unit string  `json:"unit"`
            Value float64  `json:"value"`
         
    }
    
    // Image used by Catalog
    type Image struct {

        
            AspectRatio string  `json:"aspect_ratio"`
            SecureURL string  `json:"secure_url"`
            AspectRatioF float64  `json:"aspect_ratio_f"`
            URL string  `json:"url"`
         
    }
    
    // ProductSchemaV2 used by Catalog
    type ProductSchemaV2 struct {

        
            AllSizes []map[string]interface{}  `json:"all_sizes"`
            CompanyID float64  `json:"company_id"`
            ReturnConfig ReturnConfigResponse  `json:"return_config"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            CategoryUID float64  `json:"category_uid"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            TemplateTag string  `json:"template_tag"`
            Category map[string]interface{}  `json:"category"`
            ItemType string  `json:"item_type"`
            ID string  `json:"id"`
            Highlights []string  `json:"highlights"`
            IsExpirable bool  `json:"is_expirable"`
            ModifiedOn string  `json:"modified_on"`
            UID float64  `json:"uid"`
            Brand Brand  `json:"brand"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            IsActive bool  `json:"is_active"`
            Trader []Trader  `json:"trader"`
            PrimaryColor string  `json:"primary_color"`
            ProductGroupTag []string  `json:"product_group_tag"`
            HsnCode string  `json:"hsn_code"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            IsImageLessProduct bool  `json:"is_image_less_product"`
            AllCompanyIds []float64  `json:"all_company_ids"`
            ProductPublish ProductPublish1  `json:"product_publish"`
            IsPhysical bool  `json:"is_physical"`
            CreatedOn string  `json:"created_on"`
            L3Mapping []string  `json:"l3_mapping"`
            Currency string  `json:"currency"`
            Attributes map[string]interface{}  `json:"attributes"`
            VerifiedBy VerifiedBy  `json:"verified_by"`
            Stage string  `json:"stage"`
            NetQuantity NetQuantityResponse  `json:"net_quantity"`
            SizeGuide string  `json:"size_guide"`
            Slug string  `json:"slug"`
            Name string  `json:"name"`
            Description string  `json:"description"`
            NoOfBoxes float64  `json:"no_of_boxes"`
            IsSet bool  `json:"is_set"`
            Media []Media  `json:"media"`
            Sizes []map[string]interface{}  `json:"sizes"`
            Color string  `json:"color"`
            Pending string  `json:"pending"`
            VariantGroup map[string]interface{}  `json:"variant_group"`
            IsDependent bool  `json:"is_dependent"`
            BrandUID float64  `json:"brand_uid"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            MultiSize bool  `json:"multi_size"`
            Images []Image  `json:"images"`
            Moq map[string]interface{}  `json:"moq"`
            ImageNature string  `json:"image_nature"`
            ItemCode string  `json:"item_code"`
            Tags []string  `json:"tags"`
            AllIdentifiers []string  `json:"all_identifiers"`
            VerifiedOn string  `json:"verified_on"`
            VariantMedia map[string]interface{}  `json:"variant_media"`
            Departments []float64  `json:"departments"`
            ShortDescription string  `json:"short_description"`
            Variants map[string]interface{}  `json:"variants"`
            CountryOfOrigin string  `json:"country_of_origin"`
            CategorySlug string  `json:"category_slug"`
         
    }
    
    // ProductListingResponseV2 used by Catalog
    type ProductListingResponseV2 struct {

        
            Page Page  `json:"page"`
            Items []ProductSchemaV2  `json:"items"`
         
    }
    
    // ProductVariants used by Catalog
    type ProductVariants struct {

        
            BrandUID float64  `json:"brand_uid"`
            CategoryUID float64  `json:"category_uid"`
            ItemCode string  `json:"item_code"`
            Name string  `json:"name"`
            Media []Media  `json:"media"`
            UID float64  `json:"uid"`
         
    }
    
    // ProductVariantsResponse used by Catalog
    type ProductVariantsResponse struct {

        
            Page Page  `json:"page"`
            Variants []ProductVariants  `json:"variants"`
         
    }
    
    // AttributeMasterSerializer used by Catalog
    type AttributeMasterSerializer struct {

        
            CreatedBy map[string]interface{}  `json:"created_by"`
            Suggestion string  `json:"suggestion"`
            ModifiedOn string  `json:"modified_on"`
            Schema AttributeMaster  `json:"schema"`
            EnabledForEndConsumer bool  `json:"enabled_for_end_consumer"`
            Details AttributeMasterDetails  `json:"details"`
            Filters AttributeMasterFilter  `json:"filters"`
            Slug string  `json:"slug"`
            Name string  `json:"name"`
            Description string  `json:"description"`
            Variant bool  `json:"variant"`
            Synonyms map[string]interface{}  `json:"synonyms"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            RawKey string  `json:"raw_key"`
            Unit string  `json:"unit"`
            Logo string  `json:"logo"`
            Tags []string  `json:"tags"`
            IsNested bool  `json:"is_nested"`
            Example string  `json:"example"`
            Departments []string  `json:"departments"`
            CreatedOn string  `json:"created_on"`
         
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

        
            ItemHeight float64  `json:"item_height"`
            ItemWeightUnitOfMeasure interface{}  `json:"item_weight_unit_of_measure"`
            ItemDimensionsUnitOfMeasure string  `json:"item_dimensions_unit_of_measure"`
            ItemLength float64  `json:"item_length"`
            ItemWidth float64  `json:"item_width"`
            Size interface{}  `json:"size"`
            ItemWeight float64  `json:"item_weight"`
            Identifiers []ValidateIdentifier  `json:"identifiers"`
         
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

        
            AllSizes []map[string]interface{}  `json:"all_sizes"`
            CompanyID float64  `json:"company_id"`
            ReturnConfig ReturnConfigResponse  `json:"return_config"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            CategoryUID float64  `json:"category_uid"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            TemplateTag string  `json:"template_tag"`
            Category map[string]interface{}  `json:"category"`
            ItemType string  `json:"item_type"`
            ID string  `json:"id"`
            Highlights []string  `json:"highlights"`
            IsExpirable bool  `json:"is_expirable"`
            ModifiedOn string  `json:"modified_on"`
            UID float64  `json:"uid"`
            Brand Brand  `json:"brand"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            IsActive bool  `json:"is_active"`
            Trader []Trader  `json:"trader"`
            PrimaryColor string  `json:"primary_color"`
            ProductGroupTag []string  `json:"product_group_tag"`
            HsnCode string  `json:"hsn_code"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            IsImageLessProduct bool  `json:"is_image_less_product"`
            AllCompanyIds []float64  `json:"all_company_ids"`
            ProductPublish ProductPublished  `json:"product_publish"`
            IsPhysical bool  `json:"is_physical"`
            CreatedOn string  `json:"created_on"`
            L3Mapping []string  `json:"l3_mapping"`
            Currency string  `json:"currency"`
            Attributes map[string]interface{}  `json:"attributes"`
            VerifiedBy VerifiedBy  `json:"verified_by"`
            Stage string  `json:"stage"`
            NetQuantity NetQuantityResponse  `json:"net_quantity"`
            SizeGuide string  `json:"size_guide"`
            Slug string  `json:"slug"`
            Name string  `json:"name"`
            Description string  `json:"description"`
            NoOfBoxes float64  `json:"no_of_boxes"`
            IsSet bool  `json:"is_set"`
            Media []Media  `json:"media"`
            Sizes []map[string]interface{}  `json:"sizes"`
            Color string  `json:"color"`
            Pending string  `json:"pending"`
            VariantGroup map[string]interface{}  `json:"variant_group"`
            IsDependent bool  `json:"is_dependent"`
            BrandUID float64  `json:"brand_uid"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            MultiSize bool  `json:"multi_size"`
            Images []Image  `json:"images"`
            Moq map[string]interface{}  `json:"moq"`
            ImageNature string  `json:"image_nature"`
            ItemCode string  `json:"item_code"`
            Tags []string  `json:"tags"`
            AllIdentifiers []string  `json:"all_identifiers"`
            VerifiedOn string  `json:"verified_on"`
            VariantMedia map[string]interface{}  `json:"variant_media"`
            Departments []float64  `json:"departments"`
            ShortDescription string  `json:"short_description"`
            Variants map[string]interface{}  `json:"variants"`
            CountryOfOrigin string  `json:"country_of_origin"`
            CategorySlug string  `json:"category_slug"`
         
    }
    
    // ProductListingResponse used by Catalog
    type ProductListingResponse struct {

        
            Page Page  `json:"page"`
            Items []Product  `json:"items"`
         
    }
    
    // BulkJob used by Catalog
    type BulkJob struct {

        
            Stage string  `json:"stage"`
            ModifiedBy UserInfo1  `json:"modified_by"`
            CompanyID float64  `json:"company_id"`
            CustomTemplateTag string  `json:"custom_template_tag"`
            CreatedBy UserInfo1  `json:"created_by"`
            TemplateTag string  `json:"template_tag"`
            TrackingURL string  `json:"tracking_url"`
            FailedRecords []map[string]interface{}  `json:"failed_records"`
            Cancelled float64  `json:"cancelled"`
            Total float64  `json:"total"`
            CancelledRecords []map[string]interface{}  `json:"cancelled_records"`
            Failed float64  `json:"failed"`
            ModifiedOn string  `json:"modified_on"`
            FilePath string  `json:"file_path"`
            IsActive bool  `json:"is_active"`
            Succeed float64  `json:"succeed"`
            CreatedOn string  `json:"created_on"`
         
    }
    
    // BulkResponse used by Catalog
    type BulkResponse struct {

        
            ModifiedBy UserInfo1  `json:"modified_by"`
            BatchID string  `json:"batch_id"`
            CreatedBy UserInfo1  `json:"created_by"`
            ModifiedOn string  `json:"modified_on"`
            IsActive bool  `json:"is_active"`
            CreatedOn string  `json:"created_on"`
         
    }
    
    // UserDetail1 used by Catalog
    type UserDetail1 struct {

        
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
            FullName string  `json:"full_name"`
         
    }
    
    // ProductBulkRequest used by Catalog
    type ProductBulkRequest struct {

        
            Stage string  `json:"stage"`
            ModifiedBy UserDetail1  `json:"modified_by"`
            CompanyID float64  `json:"company_id"`
            CreatedBy UserDetail1  `json:"created_by"`
            TemplateTag string  `json:"template_tag"`
            FailedRecords []string  `json:"failed_records"`
            Cancelled float64  `json:"cancelled"`
            CancelledRecords []string  `json:"cancelled_records"`
            Total float64  `json:"total"`
            Failed float64  `json:"failed"`
            ModifiedOn string  `json:"modified_on"`
            FilePath string  `json:"file_path"`
            Template ProductTemplate  `json:"template"`
            IsActive bool  `json:"is_active"`
            Succeed float64  `json:"succeed"`
            CreatedOn string  `json:"created_on"`
         
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

        
            User map[string]interface{}  `json:"user"`
            CompanyID float64  `json:"company_id"`
            BatchID string  `json:"batch_id"`
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

        
            Stage string  `json:"stage"`
            ModifiedBy UserCommon  `json:"modified_by"`
            CompanyID float64  `json:"company_id"`
            Retry float64  `json:"retry"`
            CreatedBy UserCommon  `json:"created_by"`
            TrackingURL string  `json:"tracking_url"`
            FailedRecords []string  `json:"failed_records"`
            ID string  `json:"id"`
            Cancelled float64  `json:"cancelled"`
            CancelledRecords []string  `json:"cancelled_records"`
            FilePath string  `json:"file_path"`
            Failed float64  `json:"failed"`
            ModifiedOn string  `json:"modified_on"`
            Total float64  `json:"total"`
            IsActive bool  `json:"is_active"`
            Succeed float64  `json:"succeed"`
            CreatedOn string  `json:"created_on"`
         
    }
    
    // BulkAssetResponse used by Catalog
    type BulkAssetResponse struct {

        
            Page Page  `json:"page"`
            Items []Items  `json:"items"`
         
    }
    
    // ProductSizeDeleteDataResponse used by Catalog
    type ProductSizeDeleteDataResponse struct {

        
            CompanyID float64  `json:"company_id"`
            Size string  `json:"size"`
            ItemID float64  `json:"item_id"`
         
    }
    
    // ProductSizeDeleteResponse used by Catalog
    type ProductSizeDeleteResponse struct {

        
            Success bool  `json:"success"`
            Data ProductSizeDeleteDataResponse  `json:"data"`
         
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

        
            GtinType string  `json:"gtin_type"`
            GtinValue interface{}  `json:"gtin_value"`
            Primary bool  `json:"primary"`
         
    }
    
    // InvSize used by Catalog
    type InvSize struct {

        
            Price float64  `json:"price"`
            ItemHeight float64  `json:"item_height"`
            ExpirationDate string  `json:"expiration_date"`
            PriceTransfer float64  `json:"price_transfer"`
            ItemWeightUnitOfMeasure string  `json:"item_weight_unit_of_measure"`
            ItemDimensionsUnitOfMeasure string  `json:"item_dimensions_unit_of_measure"`
            ItemLength float64  `json:"item_length"`
            ItemWidth float64  `json:"item_width"`
            Size interface{}  `json:"size"`
            PriceEffective float64  `json:"price_effective"`
            Quantity float64  `json:"quantity"`
            StoreCode string  `json:"store_code"`
            IsSet bool  `json:"is_set"`
            Currency string  `json:"currency"`
            ItemWeight float64  `json:"item_weight"`
            Set InventorySet  `json:"set"`
            Identifiers []GTIN  `json:"identifiers"`
         
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
    
    // InventoryResponse used by Catalog
    type InventoryResponse struct {

        
            Price float64  `json:"price"`
            PriceTransfer float64  `json:"price_transfer"`
            Quantity float64  `json:"quantity"`
            SellerIdentifier string  `json:"seller_identifier"`
            Size string  `json:"size"`
            PriceEffective float64  `json:"price_effective"`
            InventoryUpdatedOn string  `json:"inventory_updated_on"`
            SellableQuantity float64  `json:"sellable_quantity"`
            Store map[string]interface{}  `json:"store"`
            UID string  `json:"uid"`
            Currency string  `json:"currency"`
            Identifiers map[string]interface{}  `json:"identifiers"`
            ItemID float64  `json:"item_id"`
         
    }
    
    // InventoryResponsePaginated used by Catalog
    type InventoryResponsePaginated struct {

        
            Page Page  `json:"page"`
            Items []InventoryResponse  `json:"items"`
         
    }
    
    // ReturnConfig1 used by Catalog
    type ReturnConfig1 struct {

        
            Time float64  `json:"time"`
            Unit string  `json:"unit"`
            Returnable bool  `json:"returnable"`
         
    }
    
    // BrandMeta used by Catalog
    type BrandMeta struct {

        
            Name string  `json:"name"`
            ID float64  `json:"id"`
         
    }
    
    // Trader1 used by Catalog
    type Trader1 struct {

        
            Address []string  `json:"address"`
            Name string  `json:"name"`
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
            Sellable QuantityBase  `json:"sellable"`
            NotAvailable QuantityBase  `json:"not_available"`
            Damaged QuantityBase  `json:"damaged"`
         
    }
    
    // ManufacturerResponse used by Catalog
    type ManufacturerResponse struct {

        
            Address string  `json:"address"`
            Name string  `json:"name"`
            IsDefault bool  `json:"is_default"`
         
    }
    
    // DimensionResponse used by Catalog
    type DimensionResponse struct {

        
            Length float64  `json:"length"`
            Unit string  `json:"unit"`
            Width float64  `json:"width"`
            IsDefault bool  `json:"is_default"`
            Height float64  `json:"height"`
         
    }
    
    // PriceMeta used by Catalog
    type PriceMeta struct {

        
            UpdatedAt string  `json:"updated_at"`
            Marked float64  `json:"marked"`
            Currency string  `json:"currency"`
            Effective float64  `json:"effective"`
            TpNotes map[string]interface{}  `json:"tp_notes"`
            Transfer float64  `json:"transfer"`
         
    }
    
    // WeightResponse used by Catalog
    type WeightResponse struct {

        
            Shipping float64  `json:"shipping"`
            Unit string  `json:"unit"`
            IsDefault bool  `json:"is_default"`
         
    }
    
    // StoreMeta used by Catalog
    type StoreMeta struct {

        
            ID float64  `json:"id"`
         
    }
    
    // CompanyMeta used by Catalog
    type CompanyMeta struct {

        
            ID float64  `json:"id"`
         
    }
    
    // InventorySellerResponse used by Catalog
    type InventorySellerResponse struct {

        
            RawMeta map[string]interface{}  `json:"raw_meta"`
            ReturnConfig ReturnConfig1  `json:"return_config"`
            CreatedBy UserSerializer  `json:"created_by"`
            SellerIdentifier string  `json:"seller_identifier"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            TraceID string  `json:"trace_id"`
            UID string  `json:"uid"`
            Brand BrandMeta  `json:"brand"`
            TrackInventory bool  `json:"track_inventory"`
            Fragile bool  `json:"fragile"`
            IsActive bool  `json:"is_active"`
            ItemID float64  `json:"item_id"`
            Trader []Trader1  `json:"trader"`
            Identifier map[string]interface{}  `json:"identifier"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            FyndItemCode string  `json:"fynd_item_code"`
            Size string  `json:"size"`
            TotalQuantity float64  `json:"total_quantity"`
            Quantities Quantities  `json:"quantities"`
            Manufacturer ManufacturerResponse  `json:"manufacturer"`
            Set InventorySet  `json:"set"`
            Stage string  `json:"stage"`
            ExpirationDate string  `json:"expiration_date"`
            IsSet bool  `json:"is_set"`
            Dimension DimensionResponse  `json:"dimension"`
            AddedOnStore string  `json:"added_on_store"`
            Meta map[string]interface{}  `json:"meta"`
            Price PriceMeta  `json:"price"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            FyndArticleCode string  `json:"fynd_article_code"`
            FyndMeta map[string]interface{}  `json:"fynd_meta"`
            Weight WeightResponse  `json:"weight"`
            Tags []string  `json:"tags"`
            Store StoreMeta  `json:"store"`
            Company CompanyMeta  `json:"company"`
            CountryOfOrigin string  `json:"country_of_origin"`
         
    }
    
    // InventorySellerIdentifierResponsePaginated used by Catalog
    type InventorySellerIdentifierResponsePaginated struct {

        
            Page Page  `json:"page"`
            Items []InventorySellerResponse  `json:"items"`
         
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
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
            AddedOnStore string  `json:"added_on_store"`
         
    }
    
    // BrandMeta1 used by Catalog
    type BrandMeta1 struct {

        
            Name string  `json:"name"`
            ID float64  `json:"id"`
         
    }
    
    // Trader2 used by Catalog
    type Trader2 struct {

        
            Address []string  `json:"address"`
            Name string  `json:"name"`
            Type string  `json:"type"`
         
    }
    
    // Quantity used by Catalog
    type Quantity struct {

        
            Count float64  `json:"count"`
         
    }
    
    // QuantitiesArticle used by Catalog
    type QuantitiesArticle struct {

        
            OrderCommitted Quantity  `json:"order_committed"`
            Sellable Quantity  `json:"sellable"`
            NotAvailable Quantity  `json:"not_available"`
            Damaged Quantity  `json:"damaged"`
         
    }
    
    // ManufacturerResponse1 used by Catalog
    type ManufacturerResponse1 struct {

        
            Address string  `json:"address"`
            Name string  `json:"name"`
            IsDefault bool  `json:"is_default"`
         
    }
    
    // DimensionResponse1 used by Catalog
    type DimensionResponse1 struct {

        
            Length float64  `json:"length"`
            Height float64  `json:"height"`
            Unit string  `json:"unit"`
            Width float64  `json:"width"`
         
    }
    
    // PriceArticle used by Catalog
    type PriceArticle struct {

        
            Marked float64  `json:"marked"`
            Currency string  `json:"currency"`
            Effective float64  `json:"effective"`
            TpNotes map[string]interface{}  `json:"tp_notes"`
            Transfer float64  `json:"transfer"`
         
    }
    
    // WeightResponse1 used by Catalog
    type WeightResponse1 struct {

        
            Shipping float64  `json:"shipping"`
            Unit string  `json:"unit"`
         
    }
    
    // ArticleStoreResponse used by Catalog
    type ArticleStoreResponse struct {

        
            Name string  `json:"name"`
            StoreType string  `json:"store_type"`
            StoreCode string  `json:"store_code"`
            UID float64  `json:"uid"`
         
    }
    
    // CompanyMeta1 used by Catalog
    type CompanyMeta1 struct {

        
            ID float64  `json:"id"`
         
    }
    
    // GetInventories used by Catalog
    type GetInventories struct {

        
            ReturnConfig ReturnConfig2  `json:"return_config"`
            CreatedBy UserSerializer  `json:"created_by"`
            SellerIdentifier string  `json:"seller_identifier"`
            TraceID string  `json:"trace_id"`
            DateMeta DateMeta  `json:"date_meta"`
            ID string  `json:"id"`
            Brand BrandMeta1  `json:"brand"`
            TrackInventory bool  `json:"track_inventory"`
            UID string  `json:"uid"`
            ItemID float64  `json:"item_id"`
            Trader []Trader2  `json:"trader"`
            Identifier map[string]interface{}  `json:"identifier"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            Size string  `json:"size"`
            TotalQuantity float64  `json:"total_quantity"`
            InventoryUpdatedOn string  `json:"inventory_updated_on"`
            Quantities QuantitiesArticle  `json:"quantities"`
            Manufacturer ManufacturerResponse1  `json:"manufacturer"`
            Stage string  `json:"stage"`
            ExpirationDate string  `json:"expiration_date"`
            IsSet bool  `json:"is_set"`
            Dimension DimensionResponse1  `json:"dimension"`
            Platforms map[string]interface{}  `json:"platforms"`
            Price PriceArticle  `json:"price"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            Weight WeightResponse1  `json:"weight"`
            Tags []string  `json:"tags"`
            Store ArticleStoreResponse  `json:"store"`
            Company CompanyMeta1  `json:"company"`
            CountryOfOrigin string  `json:"country_of_origin"`
         
    }
    
    // GetInventoriesResponse used by Catalog
    type GetInventoriesResponse struct {

        
            Page Page  `json:"page"`
            Items []GetInventories  `json:"items"`
         
    }
    
    // BulkInventoryGetItems used by Catalog
    type BulkInventoryGetItems struct {

        
            Stage string  `json:"stage"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            CompanyID float64  `json:"company_id"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            FailedRecords []string  `json:"failed_records"`
            ID string  `json:"id"`
            Cancelled float64  `json:"cancelled"`
            CancelledRecords []string  `json:"cancelled_records"`
            FilePath string  `json:"file_path"`
            Failed float64  `json:"failed"`
            ModifiedOn string  `json:"modified_on"`
            Total float64  `json:"total"`
            IsActive bool  `json:"is_active"`
            Succeed float64  `json:"succeed"`
            CreatedOn string  `json:"created_on"`
         
    }
    
    // BulkInventoryGet used by Catalog
    type BulkInventoryGet struct {

        
            Page Page  `json:"page"`
            Items []BulkInventoryGetItems  `json:"items"`
         
    }
    
    // InventoryJobPayload used by Catalog
    type InventoryJobPayload struct {

        
            Price float64  `json:"price"`
            ExpirationDate string  `json:"expiration_date"`
            Quantity float64  `json:"quantity"`
            ItemWeightUnitOfMeasure string  `json:"item_weight_unit_of_measure"`
            ItemDimensionsUnitOfMeasure string  `json:"item_dimensions_unit_of_measure"`
            SellerIdentifier string  `json:"seller_identifier"`
            Tags []string  `json:"tags"`
            StoreCode string  `json:"store_code"`
            PriceEffective float64  `json:"price_effective"`
            TotalQuantity float64  `json:"total_quantity"`
            TraceID string  `json:"trace_id"`
            PriceMarked float64  `json:"price_marked"`
            Currency string  `json:"currency"`
         
    }
    
    // InventoryBulkRequest used by Catalog
    type InventoryBulkRequest struct {

        
            Sizes []InventoryJobPayload  `json:"sizes"`
            User map[string]interface{}  `json:"user"`
            CompanyID float64  `json:"company_id"`
            BatchID string  `json:"batch_id"`
         
    }
    
    // InventoryExportRequest used by Catalog
    type InventoryExportRequest struct {

        
            Store []float64  `json:"store"`
            Type string  `json:"type"`
            Brand []float64  `json:"brand"`
         
    }
    
    // InventoryExportResponse used by Catalog
    type InventoryExportResponse struct {

        
            NotificationEmails []string  `json:"notification_emails"`
            Type string  `json:"type"`
            CreatedBy string  `json:"created_by"`
            TaskID string  `json:"task_id"`
            Status string  `json:"status"`
            SellerID float64  `json:"seller_id"`
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
            Filters map[string]interface{}  `json:"filters"`
         
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
            FromDate string  `json:"from_date"`
            Quantity InventoryExportQuantityFilter  `json:"quantity"`
            StoreIds []float64  `json:"store_ids"`
            BrandIds []float64  `json:"brand_ids"`
         
    }
    
    // InventoryExportJob used by Catalog
    type InventoryExportJob struct {

        
            NotificationEmails []string  `json:"notification_emails"`
            Type string  `json:"type"`
            TaskID string  `json:"task_id"`
            Status string  `json:"status"`
            SellerID float64  `json:"seller_id"`
            CompletedOn string  `json:"completed_on"`
            Filters InventoryExportAdvanceOption  `json:"filters"`
            URL string  `json:"url"`
         
    }
    
    // InventoryExportFilter used by Catalog
    type InventoryExportFilter struct {

        
            ToDate string  `json:"to_date"`
            FromDate string  `json:"from_date"`
            Quantity InventoryExportQuantityFilter  `json:"quantity"`
            StoreIds []float64  `json:"store_ids"`
            BrandIds []float64  `json:"brand_ids"`
         
    }
    
    // InventoryCreateRequest used by Catalog
    type InventoryCreateRequest struct {

        
            NotificationEmails []string  `json:"notification_emails"`
            Data []string  `json:"data"`
            Type string  `json:"type"`
            Filters InventoryExportFilter  `json:"filters"`
         
    }
    
    // InventoryJobFilters used by Catalog
    type InventoryJobFilters struct {

        
            Stores []string  `json:"stores"`
            ToDate string  `json:"to_date"`
            FromDate string  `json:"from_date"`
            Brands []string  `json:"brands"`
            Quantity InventoryExportQuantityFilter  `json:"quantity"`
         
    }
    
    // InventoryJobDetailResponse used by Catalog
    type InventoryJobDetailResponse struct {

        
            NotificationEmails []string  `json:"notification_emails"`
            Type string  `json:"type"`
            CreatedBy UserDetail  `json:"created_by"`
            TaskID string  `json:"task_id"`
            ID string  `json:"id"`
            Status interface{}  `json:"status"`
            SellerID float64  `json:"seller_id"`
            CancelledBy UserDetail  `json:"cancelled_by"`
            ModifiedOn string  `json:"modified_on"`
            CancelledOn string  `json:"cancelled_on"`
            CompletedOn string  `json:"completed_on"`
            CreatedOn string  `json:"created_on"`
            Filters InventoryJobFilters  `json:"filters"`
            URL string  `json:"url"`
         
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

        
            ExpirationDate string  `json:"expiration_date"`
            StoreID float64  `json:"store_id"`
            SellerIdentifier string  `json:"seller_identifier"`
            Tags []string  `json:"tags"`
            PriceEffective float64  `json:"price_effective"`
            TotalQuantity float64  `json:"total_quantity"`
            TraceID string  `json:"trace_id"`
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

        
            Message string  `json:"message"`
            Items []InventoryResponseItem  `json:"items"`
         
    }
    
    // HsnUpsert used by Catalog
    type HsnUpsert struct {

        
            CompanyID float64  `json:"company_id"`
            HsnCode string  `json:"hsn_code"`
            Tax1 float64  `json:"tax1"`
            Threshold2 float64  `json:"threshold2"`
            Tax2 float64  `json:"tax2"`
            TaxOnEsp bool  `json:"tax_on_esp"`
            TaxOnMrp bool  `json:"tax_on_mrp"`
            UID float64  `json:"uid"`
            Hs2Code string  `json:"hs2_code"`
            IsActive bool  `json:"is_active"`
            Threshold1 float64  `json:"threshold1"`
         
    }
    
    // HsnCodesObject used by Catalog
    type HsnCodesObject struct {

        
            CompanyID float64  `json:"company_id"`
            HsnCode string  `json:"hsn_code"`
            Tax1 float64  `json:"tax1"`
            Threshold2 float64  `json:"threshold2"`
            Tax2 float64  `json:"tax2"`
            TaxOnEsp bool  `json:"tax_on_esp"`
            TaxOnMrp bool  `json:"tax_on_mrp"`
            ID string  `json:"id"`
            ModifiedOn string  `json:"modified_on"`
            Hs2Code string  `json:"hs2_code"`
            Threshold1 float64  `json:"threshold1"`
         
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

        
            ItemTotal float64  `json:"item_total"`
            Current string  `json:"current"`
            Size float64  `json:"size"`
            HasNext bool  `json:"has_next"`
            HasPrevious bool  `json:"has_previous"`
         
    }
    
    // TaxSlab used by Catalog
    type TaxSlab struct {

        
            Cess float64  `json:"cess"`
            Threshold float64  `json:"threshold"`
            EffectiveDate string  `json:"effective_date"`
            Rate float64  `json:"rate"`
         
    }
    
    // HSNDataInsertV2 used by Catalog
    type HSNDataInsertV2 struct {

        
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            HsnCode string  `json:"hsn_code"`
            Type string  `json:"type"`
            ReportingHsn string  `json:"reporting_hsn"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Description string  `json:"description"`
            CountryCode string  `json:"country_code"`
            ModifiedOn string  `json:"modified_on"`
            HsnCodeID string  `json:"hsn_code_id"`
            CreatedOn string  `json:"created_on"`
            Taxes []TaxSlab  `json:"taxes"`
         
    }
    
    // HsnCodesListingResponseSchemaV2 used by Catalog
    type HsnCodesListingResponseSchemaV2 struct {

        
            Page PageResponse  `json:"page"`
            Items []HSNDataInsertV2  `json:"items"`
         
    }
    
    // Media2 used by Catalog
    type Media2 struct {

        
            AspectRatio string  `json:"aspect_ratio"`
            Type string  `json:"type"`
            URL string  `json:"url"`
         
    }
    
    // BrandItem used by Catalog
    type BrandItem struct {

        
            Logo Media2  `json:"logo"`
            Slug string  `json:"slug"`
            Name string  `json:"name"`
            Discount string  `json:"discount"`
            Departments []string  `json:"departments"`
            Action Action  `json:"action"`
            UID float64  `json:"uid"`
            Banners ImageUrls  `json:"banners"`
         
    }
    
    // BrandListingResponse used by Catalog
    type BrandListingResponse struct {

        
            Page Page  `json:"page"`
            Items []BrandItem  `json:"items"`
         
    }
    
    // Department used by Catalog
    type Department struct {

        
            Logo Media2  `json:"logo"`
            Slug string  `json:"slug"`
            Name string  `json:"name"`
            PriorityOrder float64  `json:"priority_order"`
            UID float64  `json:"uid"`
         
    }
    
    // DepartmentResponse used by Catalog
    type DepartmentResponse struct {

        
            Items []Department  `json:"items"`
         
    }
    
    // ThirdLevelChild used by Catalog
    type ThirdLevelChild struct {

        
            Slug string  `json:"slug"`
            Name string  `json:"name"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Action Action  `json:"action"`
            UID float64  `json:"uid"`
            Childs []map[string]interface{}  `json:"childs"`
            Banners ImageUrls  `json:"banners"`
         
    }
    
    // SecondLevelChild used by Catalog
    type SecondLevelChild struct {

        
            Slug string  `json:"slug"`
            Name string  `json:"name"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Action Action  `json:"action"`
            UID float64  `json:"uid"`
            Childs []ThirdLevelChild  `json:"childs"`
            Banners ImageUrls  `json:"banners"`
         
    }
    
    // Child used by Catalog
    type Child struct {

        
            Slug string  `json:"slug"`
            Name string  `json:"name"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Action Action  `json:"action"`
            UID float64  `json:"uid"`
            Childs []SecondLevelChild  `json:"childs"`
            Banners ImageUrls  `json:"banners"`
         
    }
    
    // CategoryItems used by Catalog
    type CategoryItems struct {

        
            Slug string  `json:"slug"`
            Name string  `json:"name"`
            Action Action  `json:"action"`
            UID float64  `json:"uid"`
            Childs []Child  `json:"childs"`
            Banners ImageUrls  `json:"banners"`
         
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

        
            Page Page  `json:"page"`
            Operators map[string]interface{}  `json:"operators"`
            SortOn []ProductSortOn  `json:"sort_on"`
            Items []ProductListingDetail  `json:"items"`
            Filters []ProductFilters  `json:"filters"`
         
    }
    
    // ProductDetail used by Catalog
    type ProductDetail struct {

        
            PromoMeta map[string]interface{}  `json:"promo_meta"`
            ItemType string  `json:"item_type"`
            Highlights []string  `json:"highlights"`
            UID float64  `json:"uid"`
            Brand ProductBrand  `json:"brand"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
            ProductOnlineDate string  `json:"product_online_date"`
            Medias []Media  `json:"medias"`
            Attributes map[string]interface{}  `json:"attributes"`
            Type string  `json:"type"`
            Slug string  `json:"slug"`
            Name string  `json:"name"`
            Description string  `json:"description"`
            Color string  `json:"color"`
            Rating float64  `json:"rating"`
            Similars []string  `json:"similars"`
            HasVariant bool  `json:"has_variant"`
            Tryouts []string  `json:"tryouts"`
            ImageNature string  `json:"image_nature"`
            ItemCode string  `json:"item_code"`
            ShortDescription string  `json:"short_description"`
            RatingCount float64  `json:"rating_count"`
         
    }
    
    // InventoryPage used by Catalog
    type InventoryPage struct {

        
            NextID string  `json:"next_id"`
            Type string  `json:"type"`
            ItemTotal float64  `json:"item_total"`
            HasNext bool  `json:"has_next"`
            HasPrevious bool  `json:"has_previous"`
         
    }
    
    // InventoryStockResponse used by Catalog
    type InventoryStockResponse struct {

        
            Page InventoryPage  `json:"page"`
            Items []map[string]interface{}  `json:"items"`
         
    }
    
    // ArticleAssignment used by Catalog
    type ArticleAssignment struct {

        
            Strategy string  `json:"strategy"`
            Level string  `json:"level"`
         
    }
    
    // ArticleQuery used by Catalog
    type ArticleQuery struct {

        
            IgnoredStores []float64  `json:"ignored_stores"`
            Size string  `json:"size"`
            ItemID float64  `json:"item_id"`
         
    }
    
    // AssignStoreArticle used by Catalog
    type AssignStoreArticle struct {

        
            ArticleAssignment ArticleAssignment  `json:"article_assignment"`
            Query ArticleQuery  `json:"query"`
            Quantity float64  `json:"quantity"`
            GroupID string  `json:"group_id"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // AssignStore used by Catalog
    type AssignStore struct {

        
            CompanyID float64  `json:"company_id"`
            AppID string  `json:"app_id"`
            Pincode string  `json:"pincode"`
            Articles []AssignStoreArticle  `json:"articles"`
            ChannelIdentifier string  `json:"channel_identifier"`
            StoreIds []float64  `json:"store_ids"`
            ChannelType string  `json:"channel_type"`
         
    }
    
    // ArticleAssignment1 used by Catalog
    type ArticleAssignment1 struct {

        
            Strategy string  `json:"strategy"`
            Level string  `json:"level"`
         
    }
    
    // StoreAssignResponse used by Catalog
    type StoreAssignResponse struct {

        
            ArticleAssignment ArticleAssignment1  `json:"article_assignment"`
            CompanyID float64  `json:"company_id"`
            StrategyWiseListing []map[string]interface{}  `json:"strategy_wise_listing"`
            SCity string  `json:"s_city"`
            Quantity float64  `json:"quantity"`
            StoreID float64  `json:"store_id"`
            Size string  `json:"size"`
            ItemID float64  `json:"item_id"`
            Index float64  `json:"index"`
            PriceEffective float64  `json:"price_effective"`
            StorePincode float64  `json:"store_pincode"`
            Status bool  `json:"status"`
            ID string  `json:"_id"`
            UID string  `json:"uid"`
            PriceMarked float64  `json:"price_marked"`
            GroupID string  `json:"group_id"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // UserSerializer1 used by Catalog
    type UserSerializer1 struct {

        
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
            Contact string  `json:"contact"`
         
    }
    
    // Document used by Catalog
    type Document struct {

        
            Type string  `json:"type"`
            Value string  `json:"value"`
            Verified bool  `json:"verified"`
            LegalName string  `json:"legal_name"`
            URL string  `json:"url"`
         
    }
    
    // GetAddressSerializer used by Catalog
    type GetAddressSerializer struct {

        
            Address2 string  `json:"address2"`
            Longitude float64  `json:"longitude"`
            Pincode float64  `json:"pincode"`
            State string  `json:"state"`
            CountryCode string  `json:"country_code"`
            Latitude float64  `json:"latitude"`
            AddressType string  `json:"address_type"`
            City string  `json:"city"`
            Landmark string  `json:"landmark"`
            Country string  `json:"country"`
            Address1 string  `json:"address1"`
         
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
    
    // LocationTimingSerializer used by Catalog
    type LocationTimingSerializer struct {

        
            Minute float64  `json:"minute"`
            Hour float64  `json:"hour"`
         
    }
    
    // LocationDayWiseSerializer used by Catalog
    type LocationDayWiseSerializer struct {

        
            Open bool  `json:"open"`
            Opening LocationTimingSerializer  `json:"opening"`
            Closing LocationTimingSerializer  `json:"closing"`
            Weekday string  `json:"weekday"`
         
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

        
            OnSameStore bool  `json:"on_same_store"`
            StoreUID float64  `json:"store_uid"`
         
    }
    
    // UserSerializer2 used by Catalog
    type UserSerializer2 struct {

        
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
            Contact string  `json:"contact"`
         
    }
    
    // GetCompanySerializer used by Catalog
    type GetCompanySerializer struct {

        
            CompanyType string  `json:"company_type"`
            ModifiedBy UserSerializer2  `json:"modified_by"`
            Stage string  `json:"stage"`
            CreatedBy UserSerializer2  `json:"created_by"`
            Name string  `json:"name"`
            BusinessType string  `json:"business_type"`
            VerifiedOn string  `json:"verified_on"`
            RejectReason string  `json:"reject_reason"`
            ModifiedOn string  `json:"modified_on"`
            UID float64  `json:"uid"`
            VerifiedBy UserSerializer2  `json:"verified_by"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            CreatedOn string  `json:"created_on"`
         
    }
    
    // LocationManagerSerializer used by Catalog
    type LocationManagerSerializer struct {

        
            Name string  `json:"name"`
            MobileNo SellerPhoneNumber  `json:"mobile_no"`
            Email string  `json:"email"`
         
    }
    
    // GetLocationSerializer used by Catalog
    type GetLocationSerializer struct {

        
            CreatedBy UserSerializer1  `json:"created_by"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            PhoneNumber string  `json:"phone_number"`
            Warnings map[string]interface{}  `json:"warnings"`
            ModifiedOn string  `json:"modified_on"`
            UID float64  `json:"uid"`
            Code string  `json:"code"`
            NotificationEmails []string  `json:"notification_emails"`
            Documents []Document  `json:"documents"`
            Address GetAddressSerializer  `json:"address"`
            IntegrationType LocationIntegrationType  `json:"integration_type"`
            VerifiedBy UserSerializer1  `json:"verified_by"`
            Stage string  `json:"stage"`
            Name string  `json:"name"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            Timing []LocationDayWiseSerializer  `json:"timing"`
            StoreType string  `json:"store_type"`
            GstCredentials InvoiceDetailsSerializer  `json:"gst_credentials"`
            ProductReturnConfig ProductReturnConfigSerializer  `json:"product_return_config"`
            ModifiedBy UserSerializer1  `json:"modified_by"`
            VerifiedOn string  `json:"verified_on"`
            Company GetCompanySerializer  `json:"company"`
            Manager LocationManagerSerializer  `json:"manager"`
            DisplayName string  `json:"display_name"`
            CreatedOn string  `json:"created_on"`
         
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

        
            AppID string  `json:"app_id"`
            Logo string  `json:"logo"`
            Name string  `json:"name"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            UID float64  `json:"uid"`
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
    

    
    // UserSerializer used by CompanyProfile
    type UserSerializer struct {

        
            Contact string  `json:"contact"`
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
         
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

        
            URL string  `json:"url"`
            Verified bool  `json:"verified"`
            LegalName string  `json:"legal_name"`
            Type string  `json:"type"`
            Value string  `json:"value"`
         
    }
    
    // GetAddressSerializer used by CompanyProfile
    type GetAddressSerializer struct {

        
            City string  `json:"city"`
            Address1 string  `json:"address1"`
            CountryCode string  `json:"country_code"`
            AddressType string  `json:"address_type"`
            Pincode float64  `json:"pincode"`
            Address2 string  `json:"address2"`
            Latitude float64  `json:"latitude"`
            Landmark string  `json:"landmark"`
            State string  `json:"state"`
            Longitude float64  `json:"longitude"`
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
    
    // CompanyTaxesSerializer used by CompanyProfile
    type CompanyTaxesSerializer struct {

        
            EffectiveDate string  `json:"effective_date"`
            Rate float64  `json:"rate"`
            Enable bool  `json:"enable"`
         
    }
    
    // GetCompanyProfileSerializerResponse used by CompanyProfile
    type GetCompanyProfileSerializerResponse struct {

        
            Stage string  `json:"stage"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            NotificationEmails []string  `json:"notification_emails"`
            VerifiedOn string  `json:"verified_on"`
            BusinessCountryInfo BusinessCountryInfo  `json:"business_country_info"`
            FranchiseEnabled bool  `json:"franchise_enabled"`
            BusinessDetails BusinessDetails  `json:"business_details"`
            Mode string  `json:"mode"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            BusinessType string  `json:"business_type"`
            CreatedOn string  `json:"created_on"`
            Warnings map[string]interface{}  `json:"warnings"`
            Documents []Document  `json:"documents"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            CompanyType string  `json:"company_type"`
            BusinessInfo string  `json:"business_info"`
            ContactDetails ContactDetails  `json:"contact_details"`
            ModifiedOn string  `json:"modified_on"`
            CreatedBy UserSerializer  `json:"created_by"`
            Taxes []CompanyTaxesSerializer  `json:"taxes"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // ErrorResponse used by CompanyProfile
    type ErrorResponse struct {

        
            Meta map[string]interface{}  `json:"meta"`
            Code string  `json:"code"`
            Message string  `json:"message"`
            Status float64  `json:"status"`
         
    }
    
    // CompanyTaxesSerializer1 used by CompanyProfile
    type CompanyTaxesSerializer1 struct {

        
            EffectiveDate string  `json:"effective_date"`
            Rate float64  `json:"rate"`
            Enable bool  `json:"enable"`
         
    }
    
    // CreateUpdateAddressSerializer used by CompanyProfile
    type CreateUpdateAddressSerializer struct {

        
            City string  `json:"city"`
            Address1 string  `json:"address1"`
            CountryCode string  `json:"country_code"`
            Latitude float64  `json:"latitude"`
            Pincode float64  `json:"pincode"`
            Address2 string  `json:"address2"`
            Landmark string  `json:"landmark"`
            AddressType string  `json:"address_type"`
            State string  `json:"state"`
            Longitude float64  `json:"longitude"`
            Country string  `json:"country"`
         
    }
    
    // UpdateCompany used by CompanyProfile
    type UpdateCompany struct {

        
            NotificationEmails []string  `json:"notification_emails"`
            Warnings map[string]interface{}  `json:"warnings"`
            Documents []Document  `json:"documents"`
            Taxes []CompanyTaxesSerializer1  `json:"taxes"`
            Addresses []CreateUpdateAddressSerializer  `json:"addresses"`
            CompanyType string  `json:"company_type"`
            BusinessInfo string  `json:"business_info"`
            ContactDetails ContactDetails  `json:"contact_details"`
            Name string  `json:"name"`
            RejectReason string  `json:"reject_reason"`
            FranchiseEnabled bool  `json:"franchise_enabled"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            BusinessType string  `json:"business_type"`
            BusinessDetails BusinessDetails  `json:"business_details"`
         
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

        
            Stage string  `json:"stage"`
            Store DocumentsObj  `json:"store"`
            CompanyDocuments DocumentsObj  `json:"company_documents"`
            Product DocumentsObj  `json:"product"`
            StoreDocuments DocumentsObj  `json:"store_documents"`
            UID float64  `json:"uid"`
            Brand DocumentsObj  `json:"brand"`
         
    }
    
    // BrandBannerSerializer used by CompanyProfile
    type BrandBannerSerializer struct {

        
            Portrait string  `json:"portrait"`
            Landscape string  `json:"landscape"`
         
    }
    
    // GetBrandResponseSerializer used by CompanyProfile
    type GetBrandResponseSerializer struct {

        
            Stage string  `json:"stage"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            VerifiedOn string  `json:"verified_on"`
            Banner BrandBannerSerializer  `json:"banner"`
            Mode string  `json:"mode"`
            Description string  `json:"description"`
            RejectReason string  `json:"reject_reason"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            CreatedOn string  `json:"created_on"`
            Logo string  `json:"logo"`
            Warnings map[string]interface{}  `json:"warnings"`
            Synonyms []string  `json:"synonyms"`
            ModifiedOn string  `json:"modified_on"`
            CreatedBy UserSerializer  `json:"created_by"`
            UID float64  `json:"uid"`
            SlugKey string  `json:"slug_key"`
            Name string  `json:"name"`
         
    }
    
    // CreateUpdateBrandRequestSerializer used by CompanyProfile
    type CreateUpdateBrandRequestSerializer struct {

        
            Logo string  `json:"logo"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            Banner BrandBannerSerializer  `json:"banner"`
            Description string  `json:"description"`
            CompanyID float64  `json:"company_id"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            BrandTier string  `json:"brand_tier"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Synonyms []string  `json:"synonyms"`
         
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

        
            WebsiteURL string  `json:"website_url"`
            Socials []CompanySocialAccounts  `json:"socials"`
         
    }
    
    // CompanySerializer used by CompanyProfile
    type CompanySerializer struct {

        
            Stage string  `json:"stage"`
            CreatedBy UserSerializer  `json:"created_by"`
            Details CompanyDetails  `json:"details"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            NotificationEmails []string  `json:"notification_emails"`
            VerifiedOn string  `json:"verified_on"`
            MarketChannels []string  `json:"market_channels"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            CompanyType string  `json:"company_type"`
            UID float64  `json:"uid"`
            BusinessCountryInfo BusinessCountryInfo  `json:"business_country_info"`
            Name string  `json:"name"`
            RejectReason string  `json:"reject_reason"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            BusinessType string  `json:"business_type"`
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // CompanyBrandSerializer used by CompanyProfile
    type CompanyBrandSerializer struct {

        
            Stage string  `json:"stage"`
            CreatedBy UserSerializer  `json:"created_by"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            Warnings map[string]interface{}  `json:"warnings"`
            VerifiedOn string  `json:"verified_on"`
            Company CompanySerializer  `json:"company"`
            UID float64  `json:"uid"`
            RejectReason string  `json:"reject_reason"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            Brand GetBrandResponseSerializer  `json:"brand"`
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

        
            UID float64  `json:"uid"`
            Company float64  `json:"company"`
            Brands []float64  `json:"brands"`
         
    }
    
    // GetCompanySerializer used by CompanyProfile
    type GetCompanySerializer struct {

        
            Stage string  `json:"stage"`
            CreatedBy UserSerializer  `json:"created_by"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            VerifiedOn string  `json:"verified_on"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            CompanyType string  `json:"company_type"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            RejectReason string  `json:"reject_reason"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            BusinessType string  `json:"business_type"`
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // LocationManagerSerializer used by CompanyProfile
    type LocationManagerSerializer struct {

        
            MobileNo SellerPhoneNumber  `json:"mobile_no"`
            Name string  `json:"name"`
            Email string  `json:"email"`
         
    }
    
    // HolidayDateSerializer used by CompanyProfile
    type HolidayDateSerializer struct {

        
            StartDate string  `json:"start_date"`
            EndDate string  `json:"end_date"`
         
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

        
            Open bool  `json:"open"`
            Opening LocationTimingSerializer  `json:"opening"`
            Closing LocationTimingSerializer  `json:"closing"`
            Weekday string  `json:"weekday"`
         
    }
    
    // ProductReturnConfigSerializer used by CompanyProfile
    type ProductReturnConfigSerializer struct {

        
            OnSameStore bool  `json:"on_same_store"`
            StoreUID float64  `json:"store_uid"`
         
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

        
            Stage string  `json:"stage"`
            CreditNote bool  `json:"credit_note"`
            Address GetAddressSerializer  `json:"address"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            NotificationEmails []string  `json:"notification_emails"`
            VerifiedOn string  `json:"verified_on"`
            Company GetCompanySerializer  `json:"company"`
            Manager LocationManagerSerializer  `json:"manager"`
            Holiday []HolidaySchemaSerializer  `json:"holiday"`
            DisplayName string  `json:"display_name"`
            Timing []LocationDayWiseSerializer  `json:"timing"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            CreatedOn string  `json:"created_on"`
            Code string  `json:"code"`
            Warnings map[string]interface{}  `json:"warnings"`
            Documents []Document  `json:"documents"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            ModifiedOn string  `json:"modified_on"`
            CreatedBy UserSerializer  `json:"created_by"`
            ProductReturnConfig ProductReturnConfigSerializer  `json:"product_return_config"`
            PhoneNumber string  `json:"phone_number"`
            AutoInvoice bool  `json:"auto_invoice"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            StoreType string  `json:"store_type"`
            GstCredentials InvoiceDetailsSerializer  `json:"gst_credentials"`
         
    }
    
    // LocationListSerializer used by CompanyProfile
    type LocationListSerializer struct {

        
            Page Page  `json:"page"`
            Items []GetLocationSerializer  `json:"items"`
         
    }
    
    // AddressSerializer used by CompanyProfile
    type AddressSerializer struct {

        
            City string  `json:"city"`
            Address1 string  `json:"address1"`
            CountryCode string  `json:"country_code"`
            AddressType string  `json:"address_type"`
            Pincode float64  `json:"pincode"`
            Address2 string  `json:"address2"`
            Latitude float64  `json:"latitude"`
            Landmark string  `json:"landmark"`
            State string  `json:"state"`
            Longitude float64  `json:"longitude"`
            Country string  `json:"country"`
         
    }
    
    // LocationSerializer used by CompanyProfile
    type LocationSerializer struct {

        
            Stage string  `json:"stage"`
            CreditNote bool  `json:"credit_note"`
            Address AddressSerializer  `json:"address"`
            NotificationEmails []string  `json:"notification_emails"`
            Company float64  `json:"company"`
            Manager LocationManagerSerializer  `json:"manager"`
            Holiday []HolidaySchemaSerializer  `json:"holiday"`
            DisplayName string  `json:"display_name"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Timing []LocationDayWiseSerializer  `json:"timing"`
            Code string  `json:"code"`
            Warnings map[string]interface{}  `json:"warnings"`
            Documents []Document  `json:"documents"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            ProductReturnConfig ProductReturnConfigSerializer  `json:"product_return_config"`
            AutoInvoice bool  `json:"auto_invoice"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            StoreType string  `json:"store_type"`
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
    

    
    // Ownership used by Cart
    type Ownership struct {

        
            PayableBy string  `json:"payable_by"`
            PayableCategory string  `json:"payable_category"`
         
    }
    
    // CouponAuthor used by Cart
    type CouponAuthor struct {

        
            ModifiedBy string  `json:"modified_by"`
            CreatedBy string  `json:"created_by"`
         
    }
    
    // Identifier used by Cart
    type Identifier struct {

        
            CompanyID []float64  `json:"company_id"`
            ExcludeBrandID []float64  `json:"exclude_brand_id"`
            CategoryID []float64  `json:"category_id"`
            EmailDomain []string  `json:"email_domain"`
            ArticleID []string  `json:"article_id"`
            UserID []string  `json:"user_id"`
            StoreID []float64  `json:"store_id"`
            BrandID []float64  `json:"brand_id"`
            ItemID []float64  `json:"item_id"`
            CollectionID []string  `json:"collection_id"`
         
    }
    
    // Rule used by Cart
    type Rule struct {

        
            Key float64  `json:"key"`
            Min float64  `json:"min"`
            Max float64  `json:"max"`
            DiscountQty float64  `json:"discount_qty"`
            Value float64  `json:"value"`
         
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

        
            Networks []string  `json:"networks"`
            Types []string  `json:"types"`
            Codes []string  `json:"codes"`
            Iins []string  `json:"iins"`
            Uses PaymentAllowValue  `json:"uses"`
         
    }
    
    // BulkBundleRestriction used by Cart
    type BulkBundleRestriction struct {

        
            MultiStoreAllowed bool  `json:"multi_store_allowed"`
         
    }
    
    // PostOrder used by Cart
    type PostOrder struct {

        
            CancellationAllowed bool  `json:"cancellation_allowed"`
            ReturnAllowed bool  `json:"return_allowed"`
         
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
    
    // Restrictions used by Cart
    type Restrictions struct {

        
            PriceRange PriceRange  `json:"price_range"`
            Payments map[string]PaymentModes  `json:"payments"`
            OrderingStores []float64  `json:"ordering_stores"`
            BulkBundle BulkBundleRestriction  `json:"bulk_bundle"`
            UserGroups []float64  `json:"user_groups"`
            Platforms []string  `json:"platforms"`
            CouponAllowed bool  `json:"coupon_allowed"`
            PostOrder PostOrder  `json:"post_order"`
            Uses UsesRestriction  `json:"uses"`
            UserType string  `json:"user_type"`
         
    }
    
    // State used by Cart
    type State struct {

        
            IsPublic bool  `json:"is_public"`
            IsDisplay bool  `json:"is_display"`
            IsArchived bool  `json:"is_archived"`
         
    }
    
    // CouponSchedule used by Cart
    type CouponSchedule struct {

        
            NextSchedule []map[string]interface{}  `json:"next_schedule"`
            End string  `json:"end"`
            Cron string  `json:"cron"`
            Duration float64  `json:"duration"`
            Start string  `json:"start"`
         
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

        
            Apply DisplayMetaDict  `json:"apply"`
            Description string  `json:"description"`
            Subtitle string  `json:"subtitle"`
            Title string  `json:"title"`
            Remove DisplayMetaDict  `json:"remove"`
            Auto DisplayMetaDict  `json:"auto"`
         
    }
    
    // CouponAction used by Cart
    type CouponAction struct {

        
            TxnMode string  `json:"txn_mode"`
            ActionDate string  `json:"action_date"`
         
    }
    
    // RuleDefinition used by Cart
    type RuleDefinition struct {

        
            Type string  `json:"type"`
            CurrencyCode string  `json:"currency_code"`
            IsExact bool  `json:"is_exact"`
            CalculateOn string  `json:"calculate_on"`
            AutoApply bool  `json:"auto_apply"`
            ApplicableOn string  `json:"applicable_on"`
            Scope []string  `json:"scope"`
            ValueType string  `json:"value_type"`
         
    }
    
    // Validation used by Cart
    type Validation struct {

        
            Anonymous bool  `json:"anonymous"`
            AppID []string  `json:"app_id"`
            UserRegisteredAfter string  `json:"user_registered_after"`
         
    }
    
    // CouponDateMeta used by Cart
    type CouponDateMeta struct {

        
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // CouponAdd used by Cart
    type CouponAdd struct {

        
            Ownership Ownership  `json:"ownership"`
            Author CouponAuthor  `json:"author"`
            Identifiers Identifier  `json:"identifiers"`
            Rule []Rule  `json:"rule"`
            Restrictions Restrictions  `json:"restrictions"`
            State State  `json:"state"`
            Schedule CouponSchedule  `json:"_schedule"`
            TypeSlug string  `json:"type_slug"`
            Code string  `json:"code"`
            Validity Validity  `json:"validity"`
            DisplayMeta DisplayMeta  `json:"display_meta"`
            Tags []string  `json:"tags"`
            Action CouponAction  `json:"action"`
            RuleDefinition RuleDefinition  `json:"rule_definition"`
            Validation Validation  `json:"validation"`
            DateMeta CouponDateMeta  `json:"date_meta"`
         
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

        
            Ownership Ownership  `json:"ownership"`
            Author CouponAuthor  `json:"author"`
            Identifiers Identifier  `json:"identifiers"`
            Rule []Rule  `json:"rule"`
            Restrictions Restrictions  `json:"restrictions"`
            State State  `json:"state"`
            Schedule CouponSchedule  `json:"_schedule"`
            TypeSlug string  `json:"type_slug"`
            Code string  `json:"code"`
            Validity Validity  `json:"validity"`
            DisplayMeta DisplayMeta  `json:"display_meta"`
            Tags []string  `json:"tags"`
            Action CouponAction  `json:"action"`
            RuleDefinition RuleDefinition  `json:"rule_definition"`
            Validation Validation  `json:"validation"`
            DateMeta CouponDateMeta  `json:"date_meta"`
         
    }
    
    // CouponPartialUpdate used by Cart
    type CouponPartialUpdate struct {

        
            Schedule CouponSchedule  `json:"schedule"`
            Archive bool  `json:"archive"`
         
    }
    
    // Ownership1 used by Cart
    type Ownership1 struct {

        
            PayableBy string  `json:"payable_by"`
            PayableCategory string  `json:"payable_category"`
         
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
    
    // PostOrder1 used by Cart
    type PostOrder1 struct {

        
            CancellationAllowed bool  `json:"cancellation_allowed"`
            ReturnAllowed bool  `json:"return_allowed"`
         
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

        
            Payments []PromotionPaymentModes  `json:"payments"`
            OrderingStores []float64  `json:"ordering_stores"`
            UserID []string  `json:"user_id"`
            UserGroups []float64  `json:"user_groups"`
            OrderQuantity float64  `json:"order_quantity"`
            AnonymousUsers bool  `json:"anonymous_users"`
            UserRegistered UserRegistered  `json:"user_registered"`
            Platforms []string  `json:"platforms"`
            PostOrder PostOrder1  `json:"post_order"`
            Uses UsesRestriction1  `json:"uses"`
         
    }
    
    // PromotionAuthor used by Cart
    type PromotionAuthor struct {

        
            ModifiedBy string  `json:"modified_by"`
            CreatedBy string  `json:"created_by"`
         
    }
    
    // DisplayMeta1 used by Cart
    type DisplayMeta1 struct {

        
            Name string  `json:"name"`
            OfferText string  `json:"offer_text"`
            Description string  `json:"description"`
            OfferLabel string  `json:"offer_label"`
         
    }
    
    // PromotionDateMeta used by Cart
    type PromotionDateMeta struct {

        
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // Visibility used by Cart
    type Visibility struct {

        
            CouponList bool  `json:"coupon_list"`
            Pdp bool  `json:"pdp"`
         
    }
    
    // CompareObject used by Cart
    type CompareObject struct {

        
            GreaterThanEquals float64  `json:"greater_than_equals"`
            GreaterThan float64  `json:"greater_than"`
            Equals float64  `json:"equals"`
            LessThanEquals float64  `json:"less_than_equals"`
            LessThan float64  `json:"less_than"`
         
    }
    
    // ItemCriteria used by Cart
    type ItemCriteria struct {

        
            AllItems bool  `json:"all_items"`
            ItemCategory []float64  `json:"item_category"`
            ItemL1Category []float64  `json:"item_l1_category"`
            ItemExcludeCategory []float64  `json:"item_exclude_category"`
            ItemID []float64  `json:"item_id"`
            ItemExcludeL1Category []float64  `json:"item_exclude_l1_category"`
            CartUniqueItemAmount CompareObject  `json:"cart_unique_item_amount"`
            ItemSize []string  `json:"item_size"`
            ItemExcludeDepartment []float64  `json:"item_exclude_department"`
            CartQuantity CompareObject  `json:"cart_quantity"`
            ItemSku []string  `json:"item_sku"`
            ItemTags []string  `json:"item_tags"`
            ItemBrand []float64  `json:"item_brand"`
            ItemDepartment []float64  `json:"item_department"`
            ItemL2Category []float64  `json:"item_l2_category"`
            BuyRules []string  `json:"buy_rules"`
            ItemExcludeCompany []float64  `json:"item_exclude_company"`
            ItemCompany []float64  `json:"item_company"`
            CartUniqueItemQuantity CompareObject  `json:"cart_unique_item_quantity"`
            ItemExcludeL2Category []float64  `json:"item_exclude_l2_category"`
            ItemExcludeBrand []float64  `json:"item_exclude_brand"`
            ItemExcludeID []float64  `json:"item_exclude_id"`
            ItemExcludeStore []float64  `json:"item_exclude_store"`
            CartTotal CompareObject  `json:"cart_total"`
            ItemStore []float64  `json:"item_store"`
            AvailableZones []string  `json:"available_zones"`
            ItemExcludeSku []string  `json:"item_exclude_sku"`
         
    }
    
    // DiscountOffer used by Cart
    type DiscountOffer struct {

        
            DiscountPercentage float64  `json:"discount_percentage"`
            ApportionDiscount bool  `json:"apportion_discount"`
            MaxOfferQuantity float64  `json:"max_offer_quantity"`
            DiscountAmount float64  `json:"discount_amount"`
            MaxUsagePerTransaction float64  `json:"max_usage_per_transaction"`
            DiscountPrice float64  `json:"discount_price"`
            Code string  `json:"code"`
            MaxDiscountAmount float64  `json:"max_discount_amount"`
            PartialCanRet bool  `json:"partial_can_ret"`
            MinOfferQuantity float64  `json:"min_offer_quantity"`
         
    }
    
    // DiscountRule used by Cart
    type DiscountRule struct {

        
            ItemCriteria ItemCriteria  `json:"item_criteria"`
            DiscountType string  `json:"discount_type"`
            Offer DiscountOffer  `json:"offer"`
            BuyCondition string  `json:"buy_condition"`
         
    }
    
    // PromotionSchedule used by Cart
    type PromotionSchedule struct {

        
            NextSchedule []map[string]interface{}  `json:"next_schedule"`
            End string  `json:"end"`
            Published bool  `json:"published"`
            Cron string  `json:"cron"`
            Duration float64  `json:"duration"`
            Start string  `json:"start"`
         
    }
    
    // PromotionAction used by Cart
    type PromotionAction struct {

        
            ActionType string  `json:"action_type"`
            ActionDate string  `json:"action_date"`
         
    }
    
    // PromotionListItem used by Cart
    type PromotionListItem struct {

        
            Ownership Ownership1  `json:"ownership"`
            ApplyExclusive string  `json:"apply_exclusive"`
            Restrictions Restrictions1  `json:"restrictions"`
            Stackable bool  `json:"stackable"`
            Author PromotionAuthor  `json:"author"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            DisplayMeta DisplayMeta1  `json:"display_meta"`
            Mode string  `json:"mode"`
            ApplicationID string  `json:"application_id"`
            DateMeta PromotionDateMeta  `json:"date_meta"`
            Visiblility Visibility  `json:"visiblility"`
            BuyRules map[string]ItemCriteria  `json:"buy_rules"`
            ApplyPriority float64  `json:"apply_priority"`
            DiscountRules []DiscountRule  `json:"discount_rules"`
            CalculateOn string  `json:"calculate_on"`
            PromotionType string  `json:"promotion_type"`
            Currency string  `json:"currency"`
            Schedule PromotionSchedule  `json:"_schedule"`
            Code string  `json:"code"`
            PromoGroup string  `json:"promo_group"`
            ApplyAllDiscount bool  `json:"apply_all_discount"`
            PostOrderAction PromotionAction  `json:"post_order_action"`
         
    }
    
    // PromotionsResponse used by Cart
    type PromotionsResponse struct {

        
            Items PromotionListItem  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // PromotionAdd used by Cart
    type PromotionAdd struct {

        
            Ownership Ownership1  `json:"ownership"`
            ApplyExclusive string  `json:"apply_exclusive"`
            Restrictions Restrictions1  `json:"restrictions"`
            Stackable bool  `json:"stackable"`
            Author PromotionAuthor  `json:"author"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            DisplayMeta DisplayMeta1  `json:"display_meta"`
            Mode string  `json:"mode"`
            ApplicationID string  `json:"application_id"`
            DateMeta PromotionDateMeta  `json:"date_meta"`
            Visiblility Visibility  `json:"visiblility"`
            BuyRules map[string]ItemCriteria  `json:"buy_rules"`
            ApplyPriority float64  `json:"apply_priority"`
            DiscountRules []DiscountRule  `json:"discount_rules"`
            CalculateOn string  `json:"calculate_on"`
            PromotionType string  `json:"promotion_type"`
            Currency string  `json:"currency"`
            Schedule PromotionSchedule  `json:"_schedule"`
            Code string  `json:"code"`
            PromoGroup string  `json:"promo_group"`
            ApplyAllDiscount bool  `json:"apply_all_discount"`
            PostOrderAction PromotionAction  `json:"post_order_action"`
         
    }
    
    // PromotionUpdate used by Cart
    type PromotionUpdate struct {

        
            Ownership Ownership1  `json:"ownership"`
            ApplyExclusive string  `json:"apply_exclusive"`
            Restrictions Restrictions1  `json:"restrictions"`
            Stackable bool  `json:"stackable"`
            Author PromotionAuthor  `json:"author"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            DisplayMeta DisplayMeta1  `json:"display_meta"`
            Mode string  `json:"mode"`
            ApplicationID string  `json:"application_id"`
            DateMeta PromotionDateMeta  `json:"date_meta"`
            Visiblility Visibility  `json:"visiblility"`
            BuyRules map[string]ItemCriteria  `json:"buy_rules"`
            ApplyPriority float64  `json:"apply_priority"`
            DiscountRules []DiscountRule  `json:"discount_rules"`
            CalculateOn string  `json:"calculate_on"`
            PromotionType string  `json:"promotion_type"`
            Currency string  `json:"currency"`
            Schedule PromotionSchedule  `json:"_schedule"`
            Code string  `json:"code"`
            PromoGroup string  `json:"promo_group"`
            ApplyAllDiscount bool  `json:"apply_all_discount"`
            PostOrderAction PromotionAction  `json:"post_order_action"`
         
    }
    
    // PromotionPartialUpdate used by Cart
    type PromotionPartialUpdate struct {

        
            Schedule PromotionSchedule  `json:"schedule"`
            Archive bool  `json:"archive"`
         
    }
    
    // ActivePromosResponse used by Cart
    type ActivePromosResponse struct {

        
            Type string  `json:"type"`
            IsHidden bool  `json:"is_hidden"`
            Example string  `json:"example"`
            EntitySlug string  `json:"entity_slug"`
            ModifiedOn string  `json:"modified_on"`
            Description string  `json:"description"`
            CreatedOn string  `json:"created_on"`
            Subtitle string  `json:"subtitle"`
            Title string  `json:"title"`
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

        
            DeliveryCharges DeliveryCharges  `json:"delivery_charges"`
            RevenueEngineCoupon bool  `json:"revenue_engine_coupon"`
            MaxCartItems float64  `json:"max_cart_items"`
            MinCartValue float64  `json:"min_cart_value"`
            GiftPricing float64  `json:"gift_pricing"`
            GiftDisplayText string  `json:"gift_display_text"`
            Enabled bool  `json:"enabled"`
            BulkCoupons bool  `json:"bulk_coupons"`
         
    }
    
    // CartMetaConfigAdd used by Cart
    type CartMetaConfigAdd struct {

        
            DeliveryCharges DeliveryCharges  `json:"delivery_charges"`
            RevenueEngineCoupon bool  `json:"revenue_engine_coupon"`
            MaxCartItems float64  `json:"max_cart_items"`
            MinCartValue float64  `json:"min_cart_value"`
            GiftPricing float64  `json:"gift_pricing"`
            GiftDisplayText string  `json:"gift_display_text"`
            Enabled bool  `json:"enabled"`
            BulkCoupons bool  `json:"bulk_coupons"`
         
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
    
    // BaseInfo used by Cart
    type BaseInfo struct {

        
            Name string  `json:"name"`
            UID float64  `json:"uid"`
         
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
            URL string  `json:"url"`
            Type string  `json:"type"`
         
    }
    
    // Tags used by Cart
    type Tags struct {

        
            Tags map[string]interface{}  `json:"tags"`
         
    }
    
    // ProductImage used by Cart
    type ProductImage struct {

        
            SecureURL string  `json:"secure_url"`
            URL string  `json:"url"`
            AspectRatio string  `json:"aspect_ratio"`
         
    }
    
    // CartProduct used by Cart
    type CartProduct struct {

        
            Name string  `json:"name"`
            UID float64  `json:"uid"`
            Type string  `json:"type"`
            ItemCode string  `json:"item_code"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Brand BaseInfo  `json:"brand"`
            Categories []CategoryInfo  `json:"categories"`
            Tags []string  `json:"tags"`
            Action ProductAction  `json:"action"`
            TeaserTag Tags  `json:"teaser_tag"`
            Images []ProductImage  `json:"images"`
            Slug string  `json:"slug"`
         
    }
    
    // ProductPrice used by Cart
    type ProductPrice struct {

        
            CurrencySymbol string  `json:"currency_symbol"`
            Effective float64  `json:"effective"`
            CurrencyCode string  `json:"currency_code"`
            Marked float64  `json:"marked"`
            AddOn float64  `json:"add_on"`
            Selling float64  `json:"selling"`
         
    }
    
    // ProductPriceInfo used by Cart
    type ProductPriceInfo struct {

        
            Base ProductPrice  `json:"base"`
            Converted ProductPrice  `json:"converted"`
         
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
    
    // ProductAvailabilitySize used by Cart
    type ProductAvailabilitySize struct {

        
            Value string  `json:"value"`
            Display string  `json:"display"`
            IsAvailable bool  `json:"is_available"`
         
    }
    
    // ProductAvailability used by Cart
    type ProductAvailability struct {

        
            OutOfStock bool  `json:"out_of_stock"`
            Sizes []string  `json:"sizes"`
            IsValid bool  `json:"is_valid"`
            OtherStoreQuantity float64  `json:"other_store_quantity"`
            Deliverable bool  `json:"deliverable"`
            AvailableSizes []ProductAvailabilitySize  `json:"available_sizes"`
         
    }
    
    // Ownership2 used by Cart
    type Ownership2 struct {

        
            PayableBy string  `json:"payable_by"`
            PayableCategory string  `json:"payable_category"`
         
    }
    
    // BuyRules used by Cart
    type BuyRules struct {

        
            CartConditions map[string]interface{}  `json:"cart_conditions"`
            ItemCriteria map[string]interface{}  `json:"item_criteria"`
         
    }
    
    // DiscountRulesApp used by Cart
    type DiscountRulesApp struct {

        
            RawOffer map[string]interface{}  `json:"raw_offer"`
            ItemCriteria map[string]interface{}  `json:"item_criteria"`
            Offer map[string]interface{}  `json:"offer"`
            MatchedBuyRules []string  `json:"matched_buy_rules"`
         
    }
    
    // FreeGiftItem used by Cart
    type FreeGiftItem struct {

        
            ItemPriceDetails map[string]interface{}  `json:"item_price_details"`
            ItemID float64  `json:"item_id"`
            ItemImagesURL []string  `json:"item_images_url"`
            ItemBrandName string  `json:"item_brand_name"`
            ItemSlug string  `json:"item_slug"`
            ItemName string  `json:"item_name"`
         
    }
    
    // AppliedFreeArticles used by Cart
    type AppliedFreeArticles struct {

        
            ArticleID string  `json:"article_id"`
            Quantity float64  `json:"quantity"`
            ParentItemIdentifier string  `json:"parent_item_identifier"`
            FreeGiftItemDetails FreeGiftItem  `json:"free_gift_item_details"`
         
    }
    
    // AppliedPromotion used by Cart
    type AppliedPromotion struct {

        
            Ownership Ownership2  `json:"ownership"`
            ArticleQuantity float64  `json:"article_quantity"`
            BuyRules []BuyRules  `json:"buy_rules"`
            DiscountRules []DiscountRulesApp  `json:"discount_rules"`
            PromotionGroup string  `json:"promotion_group"`
            Amount float64  `json:"amount"`
            OfferText string  `json:"offer_text"`
            AppliedFreeArticles []AppliedFreeArticles  `json:"applied_free_articles"`
            PromotionName string  `json:"promotion_name"`
            PromoID string  `json:"promo_id"`
            MrpPromotion bool  `json:"mrp_promotion"`
            PromotionType string  `json:"promotion_type"`
         
    }
    
    // CartProductIdentifer used by Cart
    type CartProductIdentifer struct {

        
            Identifier string  `json:"identifier"`
         
    }
    
    // PromoMeta used by Cart
    type PromoMeta struct {

        
            Message string  `json:"message"`
         
    }
    
    // StoreInfo used by Cart
    type StoreInfo struct {

        
            Name string  `json:"name"`
            StoreCode string  `json:"store_code"`
            UID float64  `json:"uid"`
         
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
    
    // ProductArticle used by Cart
    type ProductArticle struct {

        
            Size string  `json:"size"`
            ProductGroupTags []string  `json:"product_group_tags"`
            Meta map[string]interface{}  `json:"meta"`
            UID string  `json:"uid"`
            Type string  `json:"type"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Seller BaseInfo  `json:"seller"`
            Quantity float64  `json:"quantity"`
            Identifier map[string]interface{}  `json:"identifier"`
            MtoQuantity float64  `json:"mto_quantity"`
            IsGiftVisible bool  `json:"is_gift_visible"`
            SellerIdentifier string  `json:"seller_identifier"`
            Store StoreInfo  `json:"store"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            GiftCard map[string]interface{}  `json:"gift_card"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            CartItemMeta map[string]interface{}  `json:"cart_item_meta"`
            Price ArticlePriceInfo  `json:"price"`
         
    }
    
    // CouponDetails used by Cart
    type CouponDetails struct {

        
            DiscountTotalQuantity float64  `json:"discount_total_quantity"`
            Code string  `json:"code"`
            DiscountSingleQuantity float64  `json:"discount_single_quantity"`
         
    }
    
    // CartProductInfo used by Cart
    type CartProductInfo struct {

        
            Key string  `json:"key"`
            Product CartProduct  `json:"product"`
            PricePerUnit ProductPriceInfo  `json:"price_per_unit"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            Price ProductPriceInfo  `json:"price"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            Availability ProductAvailability  `json:"availability"`
            Message string  `json:"message"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            Discount string  `json:"discount"`
            PromotionsApplied []AppliedPromotion  `json:"promotions_applied"`
            BulkOffer map[string]interface{}  `json:"bulk_offer"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
            PromoMeta PromoMeta  `json:"promo_meta"`
            Quantity float64  `json:"quantity"`
            Article ProductArticle  `json:"article"`
            IsSet bool  `json:"is_set"`
            Coupon CouponDetails  `json:"coupon"`
            CouponMessage string  `json:"coupon_message"`
            Moq map[string]interface{}  `json:"moq"`
         
    }
    
    // RawBreakup used by Cart
    type RawBreakup struct {

        
            YouSaved float64  `json:"you_saved"`
            GiftCard float64  `json:"gift_card"`
            Subtotal float64  `json:"subtotal"`
            Vog float64  `json:"vog"`
            Discount float64  `json:"discount"`
            DeliveryCharge float64  `json:"delivery_charge"`
            CodCharge float64  `json:"cod_charge"`
            GstCharges float64  `json:"gst_charges"`
            MrpTotal float64  `json:"mrp_total"`
            Total float64  `json:"total"`
            ConvenienceFee float64  `json:"convenience_fee"`
            FyndCash float64  `json:"fynd_cash"`
            Coupon float64  `json:"coupon"`
         
    }
    
    // DisplayBreakup used by Cart
    type DisplayBreakup struct {

        
            Key string  `json:"key"`
            Display string  `json:"display"`
            CurrencySymbol string  `json:"currency_symbol"`
            Message []string  `json:"message"`
            CurrencyCode string  `json:"currency_code"`
            Value float64  `json:"value"`
         
    }
    
    // CouponBreakup used by Cart
    type CouponBreakup struct {

        
            CouponType string  `json:"coupon_type"`
            UID string  `json:"uid"`
            Type string  `json:"type"`
            CouponValue float64  `json:"coupon_value"`
            Message string  `json:"message"`
            MinimumCartValue float64  `json:"minimum_cart_value"`
            MaxDiscountValue float64  `json:"max_discount_value"`
            Code string  `json:"code"`
            Description string  `json:"description"`
            IsApplied bool  `json:"is_applied"`
            Title string  `json:"title"`
            SubTitle string  `json:"sub_title"`
            Value float64  `json:"value"`
         
    }
    
    // LoyaltyPoints used by Cart
    type LoyaltyPoints struct {

        
            Applicable float64  `json:"applicable"`
            Description string  `json:"description"`
            IsApplied bool  `json:"is_applied"`
            Total float64  `json:"total"`
         
    }
    
    // CartBreakup used by Cart
    type CartBreakup struct {

        
            Raw RawBreakup  `json:"raw"`
            Display []DisplayBreakup  `json:"display"`
            Coupon CouponBreakup  `json:"coupon"`
            LoyaltyPoints LoyaltyPoints  `json:"loyalty_points"`
         
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

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
            Errors map[string]interface{}  `json:"errors"`
         
    }
    
    // ShippingAddress used by Cart
    type ShippingAddress struct {

        
            Name string  `json:"name"`
            Phone float64  `json:"phone"`
            Email string  `json:"email"`
            Meta map[string]interface{}  `json:"meta"`
            CountryCode string  `json:"country_code"`
            State string  `json:"state"`
            CountryIsoCode string  `json:"country_iso_code"`
            Landmark string  `json:"landmark"`
            City string  `json:"city"`
            Country string  `json:"country"`
            AreaCodeSlug string  `json:"area_code_slug"`
            Pincode float64  `json:"pincode"`
            CountryPhoneCode string  `json:"country_phone_code"`
            AreaCode string  `json:"area_code"`
            Area string  `json:"area"`
            AddressType string  `json:"address_type"`
            Address string  `json:"address"`
         
    }
    
    // OpenApiCartServiceabilityRequest used by Cart
    type OpenApiCartServiceabilityRequest struct {

        
            ShippingAddress ShippingAddress  `json:"shipping_address"`
            CartItems []CartItem  `json:"cart_items"`
         
    }
    
    // OpenApiCartServiceabilityResponse used by Cart
    type OpenApiCartServiceabilityResponse struct {

        
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            Message string  `json:"message"`
            IsValid bool  `json:"is_valid"`
            Items []CartProductInfo  `json:"items"`
         
    }
    
    // CartItemMeta used by Cart
    type CartItemMeta struct {

        
            GroupID string  `json:"group_id"`
            PrimaryItem bool  `json:"primary_item"`
         
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
            CurrentStatus string  `json:"current_status"`
            OrderID string  `json:"order_id"`
            PaymentGateway string  `json:"payment_gateway"`
         
    }
    
    // MultiTenderPaymentMethod used by Cart
    type MultiTenderPaymentMethod struct {

        
            Name string  `json:"name"`
            Amount float64  `json:"amount"`
            Mode string  `json:"mode"`
            Meta MultiTenderPaymentMeta  `json:"meta"`
         
    }
    
    // OpenApiOrderItem used by Cart
    type OpenApiOrderItem struct {

        
            DeliveryCharges float64  `json:"delivery_charges"`
            PriceMarked float64  `json:"price_marked"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            Meta CartItemMeta  `json:"meta"`
            CashbackApplied float64  `json:"cashback_applied"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            CodCharges float64  `json:"cod_charges"`
            Quantity float64  `json:"quantity"`
            Size string  `json:"size"`
            Discount float64  `json:"discount"`
            EmployeeDiscount float64  `json:"employee_discount"`
            ProductID float64  `json:"product_id"`
            PriceEffective float64  `json:"price_effective"`
            AmountPaid float64  `json:"amount_paid"`
            Files []OpenApiFiles  `json:"files"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            PaymentMethods []MultiTenderPaymentMethod  `json:"payment_methods"`
         
    }
    
    // OpenApiPlatformCheckoutReq used by Cart
    type OpenApiPlatformCheckoutReq struct {

        
            DeliveryCharges float64  `json:"delivery_charges"`
            ShippingAddress ShippingAddress  `json:"shipping_address"`
            CashbackApplied float64  `json:"cashback_applied"`
            CartItems []OpenApiOrderItem  `json:"cart_items"`
            CouponValue float64  `json:"coupon_value"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            CodCharges float64  `json:"cod_charges"`
            EmployeeDiscount map[string]interface{}  `json:"employee_discount"`
            PaymentMode string  `json:"payment_mode"`
            PaymentMethods []MultiTenderPaymentMethod  `json:"payment_methods"`
            Files []OpenApiFiles  `json:"files"`
            CartValue float64  `json:"cart_value"`
            BillingAddress ShippingAddress  `json:"billing_address"`
            Coupon string  `json:"coupon"`
            Comment string  `json:"comment"`
            CurrencyCode string  `json:"currency_code"`
            Gstin string  `json:"gstin"`
            CouponCode string  `json:"coupon_code"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            OrderID string  `json:"order_id"`
         
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

        
            DeliveryCharges map[string]interface{}  `json:"delivery_charges"`
            AppID string  `json:"app_id"`
            Articles []map[string]interface{}  `json:"articles"`
            CodCharges map[string]interface{}  `json:"cod_charges"`
            CreatedOn string  `json:"created_on"`
            FcIndexMap []float64  `json:"fc_index_map"`
            PaymentMode string  `json:"payment_mode"`
            PaymentMethods []map[string]interface{}  `json:"payment_methods"`
            ExpireAt string  `json:"expire_at"`
            Meta map[string]interface{}  `json:"meta"`
            IsActive bool  `json:"is_active"`
            UID float64  `json:"uid"`
            Comment string  `json:"comment"`
            UserID string  `json:"user_id"`
            IsDefault bool  `json:"is_default"`
            Discount float64  `json:"discount"`
            LastModified string  `json:"last_modified"`
            CheckoutMode string  `json:"checkout_mode"`
            BuyNow bool  `json:"buy_now"`
            Shipments []map[string]interface{}  `json:"shipments"`
            IsArchive bool  `json:"is_archive"`
            BulkCouponDiscount float64  `json:"bulk_coupon_discount"`
            FyndCredits map[string]interface{}  `json:"fynd_credits"`
            Payments map[string]interface{}  `json:"payments"`
            ID string  `json:"_id"`
            Promotion map[string]interface{}  `json:"promotion"`
            PickUpCustomerDetails map[string]interface{}  `json:"pick_up_customer_details"`
            CartValue float64  `json:"cart_value"`
            Cashback map[string]interface{}  `json:"cashback"`
            Coupon map[string]interface{}  `json:"coupon"`
            Gstin string  `json:"gstin"`
            OrderID string  `json:"order_id"`
            MergeQty bool  `json:"merge_qty"`
         
    }
    
    // AbandonedCartResponse used by Cart
    type AbandonedCartResponse struct {

        
            Result map[string]interface{}  `json:"result"`
            Page Page  `json:"page"`
            Message string  `json:"message"`
            Success bool  `json:"success"`
            Items []AbandonedCart  `json:"items"`
         
    }
    
    // PaymentSelectionLock used by Cart
    type PaymentSelectionLock struct {

        
            PaymentIdentifier string  `json:"payment_identifier"`
            Enabled bool  `json:"enabled"`
            DefaultOptions string  `json:"default_options"`
         
    }
    
    // CartCurrency used by Cart
    type CartCurrency struct {

        
            Code string  `json:"code"`
            Symbol string  `json:"symbol"`
         
    }
    
    // CartDetailResponse used by Cart
    type CartDetailResponse struct {

        
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            AppliedPromoDetails []AppliedPromotion  `json:"applied_promo_details"`
            CouponText string  `json:"coupon_text"`
            IsValid bool  `json:"is_valid"`
            PanNo string  `json:"pan_no"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            ID string  `json:"id"`
            Comment string  `json:"comment"`
            Message string  `json:"message"`
            LastModified string  `json:"last_modified"`
            CheckoutMode string  `json:"checkout_mode"`
            BuyNow bool  `json:"buy_now"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            PanConfig map[string]interface{}  `json:"pan_config"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            Currency CartCurrency  `json:"currency"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            Gstin string  `json:"gstin"`
            Items []CartProductInfo  `json:"items"`
         
    }
    
    // AddProductCart used by Cart
    type AddProductCart struct {

        
            ProductGroupTags []string  `json:"product_group_tags"`
            Display string  `json:"display"`
            Meta map[string]interface{}  `json:"meta"`
            Pos bool  `json:"pos"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            SellerID float64  `json:"seller_id"`
            ArticleID string  `json:"article_id"`
            ItemSize string  `json:"item_size"`
            Quantity float64  `json:"quantity"`
            StoreID float64  `json:"store_id"`
            ItemID float64  `json:"item_id"`
            ArticleAssignment map[string]interface{}  `json:"article_assignment"`
            ParentItemIdentifiers []map[string]string  `json:"parent_item_identifiers"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
         
    }
    
    // AddCartRequest used by Cart
    type AddCartRequest struct {

        
            NewCart bool  `json:"new_cart"`
            Items []AddProductCart  `json:"items"`
         
    }
    
    // AddCartDetailResponse used by Cart
    type AddCartDetailResponse struct {

        
            Message string  `json:"message"`
            Cart CartDetailResponse  `json:"cart"`
            Partial bool  `json:"partial"`
            Success bool  `json:"success"`
         
    }
    
    // UpdateProductCart used by Cart
    type UpdateProductCart struct {

        
            Meta map[string]interface{}  `json:"meta"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
            ItemIndex float64  `json:"item_index"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ArticleID string  `json:"article_id"`
            Quantity float64  `json:"quantity"`
            ItemSize string  `json:"item_size"`
            ItemID float64  `json:"item_id"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
         
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

        
            RwrdTndr string  `json:"rwrd_tndr"`
            PromoAmount string  `json:"promo_amount"`
            ItemList []map[string]interface{}  `json:"item_list"`
            PromoDesc string  `json:"promo_desc"`
            PromoID string  `json:"promo_id"`
         
    }
    
    // OverrideCartItem used by Cart
    type OverrideCartItem struct {

        
            PriceMarked float64  `json:"price_marked"`
            Quantity float64  `json:"quantity"`
            Size string  `json:"size"`
            PromoList []OverrideCartItemPromo  `json:"promo_list"`
            Discount float64  `json:"discount"`
            ItemID float64  `json:"item_id"`
            SellerIdentifier string  `json:"seller_identifier"`
            PriceEffective float64  `json:"price_effective"`
            AmountPaid float64  `json:"amount_paid"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
         
    }
    
    // OverrideCheckoutReq used by Cart
    type OverrideCheckoutReq struct {

        
            ShippingAddress map[string]interface{}  `json:"shipping_address"`
            Aggregator string  `json:"aggregator"`
            CartItems []OverrideCartItem  `json:"cart_items"`
            OrderType string  `json:"order_type"`
            MerchantCode string  `json:"merchant_code"`
            CurrencyCode string  `json:"currency_code"`
            OrderingStore float64  `json:"ordering_store"`
            PaymentMode string  `json:"payment_mode"`
            CartID string  `json:"cart_id"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
            PaymentIdentifier string  `json:"payment_identifier"`
         
    }
    
    // OverrideCheckoutResponse used by Cart
    type OverrideCheckoutResponse struct {

        
            Message string  `json:"message"`
            Success string  `json:"success"`
            Cart map[string]interface{}  `json:"cart"`
            Data map[string]interface{}  `json:"data"`
            OrderID string  `json:"order_id"`
         
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

        
            Meta map[string]interface{}  `json:"meta"`
            Source map[string]interface{}  `json:"source"`
            Token string  `json:"token"`
            CreatedOn string  `json:"created_on"`
            User map[string]interface{}  `json:"user"`
         
    }
    
    // SharedCart used by Cart
    type SharedCart struct {

        
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            CouponText string  `json:"coupon_text"`
            IsValid bool  `json:"is_valid"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            ID string  `json:"id"`
            UID string  `json:"uid"`
            Comment string  `json:"comment"`
            Message string  `json:"message"`
            SharedCartDetails SharedCartDetails  `json:"shared_cart_details"`
            LastModified string  `json:"last_modified"`
            CheckoutMode string  `json:"checkout_mode"`
            BuyNow bool  `json:"buy_now"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            CartID float64  `json:"cart_id"`
            Currency CartCurrency  `json:"currency"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            Gstin string  `json:"gstin"`
            Items []CartProductInfo  `json:"items"`
         
    }
    
    // SharedCartResponse used by Cart
    type SharedCartResponse struct {

        
            Error string  `json:"error"`
            Cart SharedCart  `json:"cart"`
         
    }
    
    // CartList used by Cart
    type CartList struct {

        
            UserID string  `json:"user_id"`
            CreatedOn string  `json:"created_on"`
            PickUpCustomerDetails map[string]interface{}  `json:"pick_up_customer_details"`
            CartID string  `json:"cart_id"`
            CartValue float64  `json:"cart_value"`
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

        
            ID string  `json:"_id"`
            FirstName string  `json:"first_name"`
            Gender string  `json:"gender"`
            CreatedAt string  `json:"created_at"`
            UID string  `json:"uid"`
            LastName string  `json:"last_name"`
            Mobile string  `json:"mobile"`
            ExternalID string  `json:"external_id"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // UserCartMappingResponse used by Cart
    type UserCartMappingResponse struct {

        
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            AppliedPromoDetails []AppliedPromotion  `json:"applied_promo_details"`
            CouponText string  `json:"coupon_text"`
            IsValid bool  `json:"is_valid"`
            PanNo string  `json:"pan_no"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            ID string  `json:"id"`
            Comment string  `json:"comment"`
            Message string  `json:"message"`
            LastModified string  `json:"last_modified"`
            CheckoutMode string  `json:"checkout_mode"`
            BuyNow bool  `json:"buy_now"`
            User UserInfo  `json:"user"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            PanConfig map[string]interface{}  `json:"pan_config"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            Currency CartCurrency  `json:"currency"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            Gstin string  `json:"gstin"`
            Items []CartProductInfo  `json:"items"`
         
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
            Items []UpdateProductCart  `json:"items"`
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

        
            CouponType string  `json:"coupon_type"`
            CouponValue float64  `json:"coupon_value"`
            Message string  `json:"message"`
            IsApplicable bool  `json:"is_applicable"`
            MinimumCartValue float64  `json:"minimum_cart_value"`
            MaxDiscountValue float64  `json:"max_discount_value"`
            Description string  `json:"description"`
            IsApplied bool  `json:"is_applied"`
            CouponCode string  `json:"coupon_code"`
            Title string  `json:"title"`
            SubTitle string  `json:"sub_title"`
            ExpiresOn string  `json:"expires_on"`
         
    }
    
    // PageCoupon used by Cart
    type PageCoupon struct {

        
            HasNext bool  `json:"has_next"`
            Current float64  `json:"current"`
            HasPrevious bool  `json:"has_previous"`
            TotalItemCount float64  `json:"total_item_count"`
            Total float64  `json:"total"`
         
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

        
            Longitude float64  `json:"longitude"`
            Latitude float64  `json:"latitude"`
         
    }
    
    // PlatformAddress used by Cart
    type PlatformAddress struct {

        
            CountryCode string  `json:"country_code"`
            Landmark string  `json:"landmark"`
            AreaCodeSlug string  `json:"area_code_slug"`
            Address string  `json:"address"`
            Email string  `json:"email"`
            Meta map[string]interface{}  `json:"meta"`
            IsActive bool  `json:"is_active"`
            ID string  `json:"id"`
            UserID string  `json:"user_id"`
            IsDefaultAddress bool  `json:"is_default_address"`
            CheckoutMode string  `json:"checkout_mode"`
            GoogleMapPoint map[string]interface{}  `json:"google_map_point"`
            AddressType string  `json:"address_type"`
            Phone string  `json:"phone"`
            State string  `json:"state"`
            Country string  `json:"country"`
            Area string  `json:"area"`
            CartID string  `json:"cart_id"`
            Name string  `json:"name"`
            GeoLocation GeoLocation  `json:"geo_location"`
            City string  `json:"city"`
            CreatedByUserID string  `json:"created_by_user_id"`
            Tags []string  `json:"tags"`
            AreaCode string  `json:"area_code"`
         
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

        
            IsDefaultAddress bool  `json:"is_default_address"`
            IsUpdated bool  `json:"is_updated"`
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

        
            ID string  `json:"id"`
            UserID string  `json:"user_id"`
            CheckoutMode string  `json:"checkout_mode"`
            CartID string  `json:"cart_id"`
            BillingAddressID string  `json:"billing_address_id"`
         
    }
    
    // ShipmentArticle used by Cart
    type ShipmentArticle struct {

        
            ArticleID string  `json:"article_id"`
            Meta string  `json:"meta"`
            Quantity string  `json:"quantity"`
         
    }
    
    // PlatformShipmentResponse used by Cart
    type PlatformShipmentResponse struct {

        
            FulfillmentID float64  `json:"fulfillment_id"`
            DpID string  `json:"dp_id"`
            Articles []ShipmentArticle  `json:"articles"`
            OrderType string  `json:"order_type"`
            FulfillmentType string  `json:"fulfillment_type"`
            Promise ShipmentPromise  `json:"promise"`
            DpOptions map[string]interface{}  `json:"dp_options"`
            Items []CartProductInfo  `json:"items"`
            Shipments float64  `json:"shipments"`
            BoxType string  `json:"box_type"`
            ShipmentType string  `json:"shipment_type"`
         
    }
    
    // PlatformCartShipmentsResponse used by Cart
    type PlatformCartShipmentsResponse struct {

        
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            AppliedPromoDetails []AppliedPromotion  `json:"applied_promo_details"`
            CouponText string  `json:"coupon_text"`
            IsValid bool  `json:"is_valid"`
            PanNo string  `json:"pan_no"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            ID string  `json:"id"`
            Comment string  `json:"comment"`
            Message string  `json:"message"`
            LastModified string  `json:"last_modified"`
            CheckoutMode string  `json:"checkout_mode"`
            BuyNow bool  `json:"buy_now"`
            Shipments []PlatformShipmentResponse  `json:"shipments"`
            Error bool  `json:"error"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            PanConfig map[string]interface{}  `json:"pan_config"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            Currency CartCurrency  `json:"currency"`
            StaffUserID string  `json:"staff_user_id"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            Gstin string  `json:"gstin"`
            Items []CartProductInfo  `json:"items"`
         
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
    
    // GiftDetail used by Cart
    type GiftDetail struct {

        
            IsGiftApplied bool  `json:"is_gift_applied"`
            GiftMessage string  `json:"gift_message"`
         
    }
    
    // ArticleGiftDetail used by Cart
    type ArticleGiftDetail struct {

        
            ArticleID GiftDetail  `json:"article_id"`
         
    }
    
    // PlatformCartMetaRequest used by Cart
    type PlatformCartMetaRequest struct {

        
            StaffUserID string  `json:"staff_user_id"`
            GiftDetails ArticleGiftDetail  `json:"gift_details"`
            Comment string  `json:"comment"`
            Gstin string  `json:"gstin"`
            CheckoutMode string  `json:"checkout_mode"`
            PickUpCustomerDetails map[string]interface{}  `json:"pick_up_customer_details"`
            PanNo string  `json:"pan_no"`
         
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

        
            ID string  `json:"_id"`
            FirstName string  `json:"first_name"`
            LastName string  `json:"last_name"`
            EmployeeCode string  `json:"employee_code"`
            User string  `json:"user"`
         
    }
    
    // Files used by Cart
    type Files struct {

        
            Key string  `json:"key"`
            Values []string  `json:"values"`
         
    }
    
    // PlatformCartCheckoutDetailRequest used by Cart
    type PlatformCartCheckoutDetailRequest struct {

        
            Pos bool  `json:"pos"`
            CallbackURL string  `json:"callback_url"`
            PaymentMode string  `json:"payment_mode"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
            Meta map[string]interface{}  `json:"meta"`
            Aggregator string  `json:"aggregator"`
            ID string  `json:"id"`
            Staff StaffCheckout  `json:"staff"`
            UserID string  `json:"user_id"`
            PickAtStoreUID float64  `json:"pick_at_store_uid"`
            CheckoutMode string  `json:"checkout_mode"`
            BillingAddressID string  `json:"billing_address_id"`
            PaymentIdentifier string  `json:"payment_identifier"`
            OrderType string  `json:"order_type"`
            PaymentAutoConfirm bool  `json:"payment_auto_confirm"`
            Files []Files  `json:"files"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            EmployeeCode string  `json:"employee_code"`
            PaymentParams map[string]interface{}  `json:"payment_params"`
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
            MerchantCode string  `json:"merchant_code"`
            AddressID string  `json:"address_id"`
            OrderingStore float64  `json:"ordering_store"`
            DeviceID string  `json:"device_id"`
         
    }
    
    // CheckCart used by Cart
    type CheckCart struct {

        
            DeliveryCharges float64  `json:"delivery_charges"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            CodCharges float64  `json:"cod_charges"`
            ErrorMessage string  `json:"error_message"`
            CouponText string  `json:"coupon_text"`
            IsValid bool  `json:"is_valid"`
            CodMessage string  `json:"cod_message"`
            UserType string  `json:"user_type"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            ID string  `json:"id"`
            UID string  `json:"uid"`
            Comment string  `json:"comment"`
            Message string  `json:"message"`
            LastModified string  `json:"last_modified"`
            CheckoutMode string  `json:"checkout_mode"`
            BuyNow bool  `json:"buy_now"`
            StoreCode string  `json:"store_code"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            CodAvailable bool  `json:"cod_available"`
            Success bool  `json:"success"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            CartID float64  `json:"cart_id"`
            Currency CartCurrency  `json:"currency"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            DeliveryChargeOrderValue float64  `json:"delivery_charge_order_value"`
            Gstin string  `json:"gstin"`
            StoreEmps []map[string]interface{}  `json:"store_emps"`
            Items []CartProductInfo  `json:"items"`
            OrderID string  `json:"order_id"`
         
    }
    
    // CartCheckoutResponse used by Cart
    type CartCheckoutResponse struct {

        
            Message string  `json:"message"`
            PaymentConfirmURL string  `json:"payment_confirm_url"`
            Success bool  `json:"success"`
            Cart CheckCart  `json:"cart"`
            CallbackURL string  `json:"callback_url"`
            Data map[string]interface{}  `json:"data"`
            OrderID string  `json:"order_id"`
            AppInterceptURL string  `json:"app_intercept_url"`
         
    }
    
    // CartDeliveryModesResponse used by Cart
    type CartDeliveryModesResponse struct {

        
            AvailableModes []string  `json:"available_modes"`
            PickupStores []float64  `json:"pickup_stores"`
         
    }
    
    // PickupStoreDetail used by Cart
    type PickupStoreDetail struct {

        
            Name string  `json:"name"`
            Email string  `json:"email"`
            Phone string  `json:"phone"`
            ID float64  `json:"id"`
            State string  `json:"state"`
            UID float64  `json:"uid"`
            Country string  `json:"country"`
            Landmark string  `json:"landmark"`
            City string  `json:"city"`
            AreaCodeSlug string  `json:"area_code_slug"`
            Pincode float64  `json:"pincode"`
            AreaCode string  `json:"area_code"`
            StoreCode string  `json:"store_code"`
            Area string  `json:"area"`
            AddressType string  `json:"address_type"`
            Address string  `json:"address"`
         
    }
    
    // StoreDetailsResponse used by Cart
    type StoreDetailsResponse struct {

        
            Items []PickupStoreDetail  `json:"items"`
         
    }
    
    // UpdateCartPaymentRequest used by Cart
    type UpdateCartPaymentRequest struct {

        
            ID string  `json:"id"`
            MerchantCode string  `json:"merchant_code"`
            AddressID string  `json:"address_id"`
            AggregatorName string  `json:"aggregator_name"`
            PaymentMode string  `json:"payment_mode"`
            PaymentIdentifier string  `json:"payment_identifier"`
         
    }
    
    // CouponValidity used by Cart
    type CouponValidity struct {

        
            NextValidationRequired bool  `json:"next_validation_required"`
            DisplayMessageEn string  `json:"display_message_en"`
            Discount float64  `json:"discount"`
            Code string  `json:"code"`
            Title string  `json:"title"`
            Valid bool  `json:"valid"`
         
    }
    
    // PaymentCouponValidate used by Cart
    type PaymentCouponValidate struct {

        
            Message string  `json:"message"`
            CouponValidity CouponValidity  `json:"coupon_validity"`
            Success bool  `json:"success"`
         
    }
    
    // PaymentMeta used by Cart
    type PaymentMeta struct {

        
            MerchantCode string  `json:"merchant_code"`
            PaymentGateway string  `json:"payment_gateway"`
            Type string  `json:"type"`
            PaymentIdentifier string  `json:"payment_identifier"`
         
    }
    
    // PaymentMethod used by Cart
    type PaymentMethod struct {

        
            Name string  `json:"name"`
            Amount float64  `json:"amount"`
            Payment string  `json:"payment"`
            PaymentMeta PaymentMeta  `json:"payment_meta"`
            Mode string  `json:"mode"`
         
    }
    
    // PlatformCartCheckoutDetailV2Request used by Cart
    type PlatformCartCheckoutDetailV2Request struct {

        
            Pos bool  `json:"pos"`
            CallbackURL string  `json:"callback_url"`
            PaymentMode string  `json:"payment_mode"`
            PaymentMethods []PaymentMethod  `json:"payment_methods"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
            Meta map[string]interface{}  `json:"meta"`
            Aggregator string  `json:"aggregator"`
            ID string  `json:"id"`
            Staff StaffCheckout  `json:"staff"`
            UserID string  `json:"user_id"`
            PickAtStoreUID float64  `json:"pick_at_store_uid"`
            CheckoutMode string  `json:"checkout_mode"`
            BillingAddressID string  `json:"billing_address_id"`
            PaymentIdentifier string  `json:"payment_identifier"`
            OrderType string  `json:"order_type"`
            PaymentAutoConfirm bool  `json:"payment_auto_confirm"`
            Files []Files  `json:"files"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            EmployeeCode string  `json:"employee_code"`
            PaymentParams map[string]interface{}  `json:"payment_params"`
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
            CustomMeta map[string]interface{}  `json:"custom_meta"`
            MerchantCode string  `json:"merchant_code"`
            AddressID string  `json:"address_id"`
            OrderingStore float64  `json:"ordering_store"`
            DeviceID string  `json:"device_id"`
         
    }
    
    // UpdateCartPaymentRequestV2 used by Cart
    type UpdateCartPaymentRequestV2 struct {

        
            ID string  `json:"id"`
            MerchantCode string  `json:"merchant_code"`
            AddressID string  `json:"address_id"`
            AggregatorName string  `json:"aggregator_name"`
            PaymentMode string  `json:"payment_mode"`
            PaymentMethods []PaymentMethod  `json:"payment_methods"`
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
    
