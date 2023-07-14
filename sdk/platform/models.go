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
            AppID string  `json:"app_id"`
            Created bool  `json:"created"`
            ExcludedFields []string  `json:"excluded_fields"`
            DisplayFields []string  `json:"display_fields"`
         
    }
    
    // ErrorCodeDescription used by Payment
    type ErrorCodeDescription struct {

        
            Description string  `json:"description"`
            Code string  `json:"code"`
            Success bool  `json:"success"`
         
    }
    
    // PaymentGatewayConfig used by Payment
    type PaymentGatewayConfig struct {

        
            MerchantSalt string  `json:"merchant_salt"`
            IsActive bool  `json:"is_active"`
            Secret string  `json:"secret"`
            Key string  `json:"key"`
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

        
            PackageName string  `json:"package_name"`
            Code string  `json:"code"`
         
    }
    
    // IntentApp used by Payment
    type IntentApp struct {

        
            Logos PaymentModeLogo  `json:"logos"`
            DisplayName string  `json:"display_name"`
            PackageName string  `json:"package_name"`
            Code string  `json:"code"`
         
    }
    
    // PaymentModeList used by Payment
    type PaymentModeList struct {

        
            CardBrandImage string  `json:"card_brand_image"`
            CodLimitPerOrder float64  `json:"cod_limit_per_order"`
            CardID string  `json:"card_id"`
            Nickname string  `json:"nickname"`
            CardReference string  `json:"card_reference"`
            Name string  `json:"name"`
            Code string  `json:"code"`
            AggregatorName string  `json:"aggregator_name"`
            CardIssuer string  `json:"card_issuer"`
            LogoURL PaymentModeLogo  `json:"logo_url"`
            IntentAppErrorList []string  `json:"intent_app_error_list"`
            RemainingLimit float64  `json:"remaining_limit"`
            CardType string  `json:"card_type"`
            CardBrand string  `json:"card_brand"`
            DisplayName string  `json:"display_name"`
            CardIsin string  `json:"card_isin"`
            IntentFlow bool  `json:"intent_flow"`
            IntentAppErrorDictList []IntentAppErrorList  `json:"intent_app_error_dict_list"`
            RetryCount float64  `json:"retry_count"`
            CodLimit float64  `json:"cod_limit"`
            ExpYear float64  `json:"exp_year"`
            CardFingerprint string  `json:"card_fingerprint"`
            IntentApp []IntentApp  `json:"intent_app"`
            CardToken string  `json:"card_token"`
            CompliantWithTokenisationGuidelines bool  `json:"compliant_with_tokenisation_guidelines"`
            Timeout float64  `json:"timeout"`
            FyndVpa string  `json:"fynd_vpa"`
            DisplayPriority float64  `json:"display_priority"`
            MerchantCode string  `json:"merchant_code"`
            CardNumber string  `json:"card_number"`
            ExpMonth float64  `json:"exp_month"`
            Expired bool  `json:"expired"`
            CardName string  `json:"card_name"`
         
    }
    
    // RootPaymentMode used by Payment
    type RootPaymentMode struct {

        
            AddCardEnabled bool  `json:"add_card_enabled"`
            AggregatorName string  `json:"aggregator_name"`
            List []PaymentModeList  `json:"list"`
            AnonymousEnable bool  `json:"anonymous_enable"`
            DisplayPriority float64  `json:"display_priority"`
            DisplayName string  `json:"display_name"`
            Name string  `json:"name"`
            IsPayByCardPl bool  `json:"is_pay_by_card_pl"`
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
    
    // PayoutsResponse used by Payment
    type PayoutsResponse struct {

        
            IsActive bool  `json:"is_active"`
            UniqueTransferNo map[string]interface{}  `json:"unique_transfer_no"`
            TransferType string  `json:"transfer_type"`
            MoreAttributes map[string]interface{}  `json:"more_attributes"`
            Customers map[string]interface{}  `json:"customers"`
            PayoutsAggregators []map[string]interface{}  `json:"payouts_aggregators"`
            IsDefault bool  `json:"is_default"`
         
    }
    
    // PayoutBankDetails used by Payment
    type PayoutBankDetails struct {

        
            Pincode float64  `json:"pincode"`
            BranchName string  `json:"branch_name"`
            IfscCode string  `json:"ifsc_code"`
            City string  `json:"city"`
            AccountNo string  `json:"account_no"`
            AccountType string  `json:"account_type"`
            State string  `json:"state"`
            BankName string  `json:"bank_name"`
            AccountHolder string  `json:"account_holder"`
            Country string  `json:"country"`
         
    }
    
    // PayoutRequest used by Payment
    type PayoutRequest struct {

        
            Users map[string]interface{}  `json:"users"`
            UniqueExternalID string  `json:"unique_external_id"`
            IsActive bool  `json:"is_active"`
            TransferType string  `json:"transfer_type"`
            BankDetails PayoutBankDetails  `json:"bank_details"`
            Aggregator string  `json:"aggregator"`
         
    }
    
    // PayoutResponse used by Payment
    type PayoutResponse struct {

        
            Users map[string]interface{}  `json:"users"`
            Created bool  `json:"created"`
            IsActive bool  `json:"is_active"`
            UniqueTransferNo string  `json:"unique_transfer_no"`
            PaymentStatus string  `json:"payment_status"`
            TransferType string  `json:"transfer_type"`
            Success bool  `json:"success"`
            BankDetails map[string]interface{}  `json:"bank_details"`
            Payouts map[string]interface{}  `json:"payouts"`
            Aggregator string  `json:"aggregator"`
         
    }
    
    // UpdatePayoutResponse used by Payment
    type UpdatePayoutResponse struct {

        
            IsActive bool  `json:"is_active"`
            IsDefault bool  `json:"is_default"`
            Success bool  `json:"success"`
         
    }
    
    // UpdatePayoutRequest used by Payment
    type UpdatePayoutRequest struct {

        
            UniqueExternalID string  `json:"unique_external_id"`
            IsActive bool  `json:"is_active"`
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

        
            IsVerifiedFlag bool  `json:"is_verified_flag"`
            Data map[string]interface{}  `json:"data"`
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // NotFoundResourceError used by Payment
    type NotFoundResourceError struct {

        
            Description string  `json:"description"`
            Code string  `json:"code"`
            Success bool  `json:"success"`
         
    }
    
    // BankDetailsForOTP used by Payment
    type BankDetailsForOTP struct {

        
            BranchName string  `json:"branch_name"`
            IfscCode string  `json:"ifsc_code"`
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

        
            BankName string  `json:"bank_name"`
            BranchName string  `json:"branch_name"`
            Success bool  `json:"success"`
         
    }
    
    // OrderBeneficiaryDetails used by Payment
    type OrderBeneficiaryDetails struct {

        
            IsActive bool  `json:"is_active"`
            DelightsUserName string  `json:"delights_user_name"`
            ModifiedOn string  `json:"modified_on"`
            Comment string  `json:"comment"`
            Address string  `json:"address"`
            DisplayName string  `json:"display_name"`
            Email string  `json:"email"`
            TransferMode string  `json:"transfer_mode"`
            IfscCode string  `json:"ifsc_code"`
            Subtitle string  `json:"subtitle"`
            AccountHolder string  `json:"account_holder"`
            Mobile string  `json:"mobile"`
            BranchName string  `json:"branch_name"`
            Title string  `json:"title"`
            AccountNo string  `json:"account_no"`
            ID float64  `json:"id"`
            BeneficiaryID string  `json:"beneficiary_id"`
            BankName string  `json:"bank_name"`
            CreatedOn string  `json:"created_on"`
         
    }
    
    // OrderBeneficiaryResponse used by Payment
    type OrderBeneficiaryResponse struct {

        
            Beneficiaries []OrderBeneficiaryDetails  `json:"beneficiaries"`
            ShowBeneficiaryDetails bool  `json:"show_beneficiary_details"`
         
    }
    
    // MultiTenderPaymentMeta used by Payment
    type MultiTenderPaymentMeta struct {

        
            PaymentID string  `json:"payment_id"`
            CurrentStatus string  `json:"current_status"`
            OrderID string  `json:"order_id"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            PaymentGateway string  `json:"payment_gateway"`
         
    }
    
    // MultiTenderPaymentMethod used by Payment
    type MultiTenderPaymentMethod struct {

        
            Meta MultiTenderPaymentMeta  `json:"meta"`
            Mode string  `json:"mode"`
            Amount float64  `json:"amount"`
            Name string  `json:"name"`
         
    }
    
    // PaymentConfirmationRequest used by Payment
    type PaymentConfirmationRequest struct {

        
            PaymentMethods []MultiTenderPaymentMethod  `json:"payment_methods"`
            OrderID string  `json:"order_id"`
         
    }
    
    // PaymentConfirmationResponse used by Payment
    type PaymentConfirmationResponse struct {

        
            OrderID string  `json:"order_id"`
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // CODdata used by Payment
    type CODdata struct {

        
            IsActive bool  `json:"is_active"`
            UserID string  `json:"user_id"`
            Usages float64  `json:"usages"`
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

        
            Mobileno string  `json:"mobileno"`
            IsActive bool  `json:"is_active"`
            MerchantUserID string  `json:"merchant_user_id"`
         
    }
    
    // SetCODOptionResponse used by Payment
    type SetCODOptionResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // EdcModelData used by Payment
    type EdcModelData struct {

        
            Aggregator string  `json:"aggregator"`
            Models []string  `json:"models"`
            AggregatorID float64  `json:"aggregator_id"`
         
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

        
            EdcDeviceSerialNo string  `json:"edc_device_serial_no"`
            DeviceTag string  `json:"device_tag"`
            AggregatorID float64  `json:"aggregator_id"`
            EdcModel string  `json:"edc_model"`
            StoreID float64  `json:"store_id"`
            TerminalSerialNo string  `json:"terminal_serial_no"`
         
    }
    
    // EdcDevice used by Payment
    type EdcDevice struct {

        
            IsActive bool  `json:"is_active"`
            EdcDeviceSerialNo string  `json:"edc_device_serial_no"`
            DeviceTag string  `json:"device_tag"`
            AggregatorName string  `json:"aggregator_name"`
            TerminalUniqueIdentifier string  `json:"terminal_unique_identifier"`
            AggregatorID float64  `json:"aggregator_id"`
            EdcModel string  `json:"edc_model"`
            StoreID float64  `json:"store_id"`
            TerminalSerialNo string  `json:"terminal_serial_no"`
            MerchantStorePosCode string  `json:"merchant_store_pos_code"`
            ApplicationID string  `json:"application_id"`
         
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

        
            IsActive bool  `json:"is_active"`
            EdcDeviceSerialNo string  `json:"edc_device_serial_no"`
            DeviceTag string  `json:"device_tag"`
            AggregatorID float64  `json:"aggregator_id"`
            EdcModel string  `json:"edc_model"`
            StoreID float64  `json:"store_id"`
            MerchantStorePosCode string  `json:"merchant_store_pos_code"`
         
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

        
            CustomerID string  `json:"customer_id"`
            DeviceID string  `json:"device_id"`
            RazorpayPaymentID string  `json:"razorpay_payment_id"`
            Currency string  `json:"currency"`
            OrderID string  `json:"order_id"`
            Method string  `json:"method"`
            Email string  `json:"email"`
            Vpa string  `json:"vpa"`
            Aggregator string  `json:"aggregator"`
            Amount float64  `json:"amount"`
            Contact string  `json:"contact"`
            Timeout float64  `json:"timeout"`
            MerchantOrderID string  `json:"merchant_order_id"`
         
    }
    
    // PaymentInitializationResponse used by Payment
    type PaymentInitializationResponse struct {

        
            Status string  `json:"status"`
            AggregatorOrderID string  `json:"aggregator_order_id"`
            CustomerID string  `json:"customer_id"`
            UpiPollURL string  `json:"upi_poll_url"`
            PollingURL string  `json:"polling_url"`
            Success bool  `json:"success"`
            DeviceID string  `json:"device_id"`
            Currency string  `json:"currency"`
            RazorpayPaymentID string  `json:"razorpay_payment_id"`
            VirtualID string  `json:"virtual_id"`
            Method string  `json:"method"`
            BqrImage string  `json:"bqr_image"`
            Vpa string  `json:"vpa"`
            Aggregator string  `json:"aggregator"`
            Amount float64  `json:"amount"`
            Timeout float64  `json:"timeout"`
            MerchantOrderID string  `json:"merchant_order_id"`
         
    }
    
    // PaymentStatusUpdateRequest used by Payment
    type PaymentStatusUpdateRequest struct {

        
            Status string  `json:"status"`
            CustomerID string  `json:"customer_id"`
            MerchantTransactionID string  `json:"merchant_transaction_id"`
            DeviceID string  `json:"device_id"`
            Currency string  `json:"currency"`
            OrderID string  `json:"order_id"`
            Method string  `json:"method"`
            Email string  `json:"email"`
            Vpa string  `json:"vpa"`
            Aggregator string  `json:"aggregator"`
            Amount float64  `json:"amount"`
            Contact string  `json:"contact"`
            MerchantOrderID string  `json:"merchant_order_id"`
         
    }
    
    // PaymentStatusUpdateResponse used by Payment
    type PaymentStatusUpdateResponse struct {

        
            Status string  `json:"status"`
            RedirectURL string  `json:"redirect_url"`
            Success bool  `json:"success"`
            AggregatorName string  `json:"aggregator_name"`
            Retry bool  `json:"retry"`
         
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

        
            Data LinkStatus  `json:"data"`
            Success bool  `json:"success"`
         
    }
    
    // PaymentStatusBulkHandlerRequest used by Payment
    type PaymentStatusBulkHandlerRequest struct {

        
            MerchantOrderID []string  `json:"merchant_order_id"`
         
    }
    
    // PaymentObjectListSerializer used by Payment
    type PaymentObjectListSerializer struct {

        
            PaymentID string  `json:"payment_id"`
            RefundedBy string  `json:"refunded_by"`
            CollectedBy string  `json:"collected_by"`
            UserObject map[string]interface{}  `json:"user_object"`
            Currency string  `json:"currency"`
            CompanyID string  `json:"company_id"`
            ApplicationID string  `json:"application_id"`
            CreatedOn string  `json:"created_on"`
            PaymentModeIdentifier string  `json:"payment_mode_identifier"`
            CurrentStatus string  `json:"current_status"`
            ModifiedOn string  `json:"modified_on"`
            AggregatorPaymentObject map[string]interface{}  `json:"aggregator_payment_object"`
            AllStatus []string  `json:"all_status"`
            ID string  `json:"id"`
            RefundObject map[string]interface{}  `json:"refund_object"`
            AmountInPaisa string  `json:"amount_in_paisa"`
            PaymentMode string  `json:"payment_mode"`
            PaymentGateway string  `json:"payment_gateway"`
         
    }
    
    // PaymentStatusObject used by Payment
    type PaymentStatusObject struct {

        
            PaymentObjectList []PaymentObjectListSerializer  `json:"payment_object_list"`
            MerchantOrderID string  `json:"merchant_order_id"`
         
    }
    
    // PaymentStatusBulkHandlerResponse used by Payment
    type PaymentStatusBulkHandlerResponse struct {

        
            Status float64  `json:"status"`
            Success string  `json:"success"`
            Error string  `json:"error"`
            Count float64  `json:"count"`
            Data []PaymentStatusObject  `json:"data"`
         
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

        
            AggregatorOrderID string  `json:"aggregator_order_id"`
            OutstandingDetailsID float64  `json:"outstanding_details_id"`
            PaymentModeIdentifier string  `json:"payment_mode_identifier"`
            CurrentStatus string  `json:"current_status"`
            AggregatorTransactionID string  `json:"aggregator_transaction_id"`
            FwdShipmentID string  `json:"fwd_shipment_id"`
            Aggregator string  `json:"aggregator"`
            Amount float64  `json:"amount"`
            PaymentMode string  `json:"payment_mode"`
            MerchantOrderID string  `json:"merchant_order_id"`
         
    }
    
    // RepaymentDetailsSerialiserPayAll used by Payment
    type RepaymentDetailsSerialiserPayAll struct {

        
            AggregatorOrderID string  `json:"aggregator_order_id"`
            ExtensionOrderID string  `json:"extension_order_id"`
            AggregatorTransactionID string  `json:"aggregator_transaction_id"`
            ShipmentDetails []RepaymentRequestDetails  `json:"shipment_details"`
            TotalAmount float64  `json:"total_amount"`
         
    }
    
    // RepaymentResponse used by Payment
    type RepaymentResponse struct {

        
            Data map[string]interface{}  `json:"data"`
            Success bool  `json:"success"`
         
    }
    
    // MerchantOnBoardingRequest used by Payment
    type MerchantOnBoardingRequest struct {

        
            Status string  `json:"status"`
            UserID string  `json:"user_id"`
            AppID string  `json:"app_id"`
            Aggregator string  `json:"aggregator"`
            CreditLineID string  `json:"credit_line_id"`
         
    }
    
    // MerchantOnBoardingResponse used by Payment
    type MerchantOnBoardingResponse struct {

        
            Data map[string]interface{}  `json:"data"`
            Success bool  `json:"success"`
         
    }
    
    // ValidateCustomerRequest used by Payment
    type ValidateCustomerRequest struct {

        
            Payload string  `json:"payload"`
            MerchantParams map[string]interface{}  `json:"merchant_params"`
            TransactionAmountInPaise float64  `json:"transaction_amount_in_paise"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
            PhoneNumber string  `json:"phone_number"`
            OrderItems []map[string]interface{}  `json:"order_items"`
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
            Aggregator string  `json:"aggregator"`
         
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
            PollingTimeout float64  `json:"polling_timeout"`
            Success bool  `json:"success"`
            MerchantName string  `json:"merchant_name"`
            PaymentLinkCurrentStatus string  `json:"payment_link_current_status"`
            StatusCode float64  `json:"status_code"`
            Amount float64  `json:"amount"`
            PaymentLinkURL string  `json:"payment_link_url"`
            Message string  `json:"message"`
         
    }
    
    // ErrorDescription used by Payment
    type ErrorDescription struct {

        
            MerchantName string  `json:"merchant_name"`
            Msg string  `json:"msg"`
            InvalidID bool  `json:"invalid_id"`
            Cancelled bool  `json:"cancelled"`
            Expired bool  `json:"expired"`
            Amount float64  `json:"amount"`
            PaymentTransactionID string  `json:"payment_transaction_id"`
            MerchantOrderID string  `json:"merchant_order_id"`
         
    }
    
    // ErrorResponse used by Payment
    type ErrorResponse struct {

        
            Success bool  `json:"success"`
            StatusCode float64  `json:"status_code"`
            Error ErrorDescription  `json:"error"`
            Message string  `json:"message"`
         
    }
    
    // CreatePaymentLinkMeta used by Payment
    type CreatePaymentLinkMeta struct {

        
            Pincode string  `json:"pincode"`
            CheckoutMode string  `json:"checkout_mode"`
            Amount string  `json:"amount"`
            AssignCardID string  `json:"assign_card_id"`
            CartID string  `json:"cart_id"`
         
    }
    
    // CreatePaymentLinkRequest used by Payment
    type CreatePaymentLinkRequest struct {

        
            Meta CreatePaymentLinkMeta  `json:"meta"`
            ExternalOrderID string  `json:"external_order_id"`
            MobileNumber string  `json:"mobile_number"`
            Email string  `json:"email"`
            Description string  `json:"description"`
            Amount float64  `json:"amount"`
         
    }
    
    // CreatePaymentLinkResponse used by Payment
    type CreatePaymentLinkResponse struct {

        
            PollingTimeout float64  `json:"polling_timeout"`
            Success bool  `json:"success"`
            StatusCode float64  `json:"status_code"`
            PaymentLinkURL string  `json:"payment_link_url"`
            PaymentLinkID string  `json:"payment_link_id"`
            Message string  `json:"message"`
         
    }
    
    // PollingPaymentLinkResponse used by Payment
    type PollingPaymentLinkResponse struct {

        
            Status string  `json:"status"`
            RedirectURL string  `json:"redirect_url"`
            AggregatorName string  `json:"aggregator_name"`
            HttpStatus float64  `json:"http_status"`
            Success bool  `json:"success"`
            OrderID string  `json:"order_id"`
            StatusCode float64  `json:"status_code"`
            Amount float64  `json:"amount"`
            PaymentLinkID string  `json:"payment_link_id"`
            Message string  `json:"message"`
         
    }
    
    // CancelOrResendPaymentLinkRequest used by Payment
    type CancelOrResendPaymentLinkRequest struct {

        
            PaymentLinkID string  `json:"payment_link_id"`
         
    }
    
    // ResendPaymentLinkResponse used by Payment
    type ResendPaymentLinkResponse struct {

        
            StatusCode float64  `json:"status_code"`
            PollingTimeout float64  `json:"polling_timeout"`
            Message string  `json:"message"`
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

        
            MerchantCode string  `json:"merchant_code"`
            Name string  `json:"name"`
            Code string  `json:"code"`
         
    }
    
    // PaymentCode used by Payment
    type PaymentCode struct {

        
            Types string  `json:"types"`
            Networks string  `json:"networks"`
            Codes Code  `json:"codes"`
            Name string  `json:"name"`
         
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
    

    
    // ShipmentTimeStamp used by Order
    type ShipmentTimeStamp struct {

        
            Min float64  `json:"min"`
            Max float64  `json:"max"`
         
    }
    
    // LockData used by Order
    type LockData struct {

        
            LockMessage string  `json:"lock_message"`
            Locked bool  `json:"locked"`
            Mto bool  `json:"mto"`
         
    }
    
    // ShipmentTags used by Order
    type ShipmentTags struct {

        
            EntityType string  `json:"entity_type"`
            DisplayText string  `json:"display_text"`
            Slug string  `json:"slug"`
         
    }
    
    // Formatted used by Order
    type Formatted struct {

        
            Min string  `json:"min"`
            Max string  `json:"max"`
         
    }
    
    // ShipmentItemMeta used by Order
    type ShipmentItemMeta struct {

        
            External map[string]interface{}  `json:"external"`
            ShippingZone string  `json:"shipping_zone"`
            ShipmentWeight float64  `json:"shipment_weight"`
            IsInternational bool  `json:"is_international"`
            Weight float64  `json:"weight"`
            PdfMedia []map[string]interface{}  `json:"pdf_media"`
            ParentDpID string  `json:"parent_dp_id"`
            Sla float64  `json:"sla"`
            BagWeight map[string]interface{}  `json:"bag_weight"`
            EwaybillInfo map[string]interface{}  `json:"ewaybill_info"`
            ExistingDpList []string  `json:"existing_dp_list"`
            Timestamp ShipmentTimeStamp  `json:"timestamp"`
            Tags []map[string]interface{}  `json:"tags"`
            OrderType string  `json:"order_type"`
            LockData LockData  `json:"lock_data"`
            StoreInvoiceUpdatedDate string  `json:"store_invoice_updated_date"`
            DebugInfo map[string]interface{}  `json:"debug_info"`
            ShipmentTags []ShipmentTags  `json:"shipment_tags"`
            DpOptions map[string]interface{}  `json:"dp_options"`
            PackagingName string  `json:"packaging_name"`
            AssignDpFromSb bool  `json:"assign_dp_from_sb"`
            DpSortKey string  `json:"dp_sort_key"`
            SameStoreAvailable bool  `json:"same_store_available"`
            ActivityComment string  `json:"activity_comment"`
            ShipmentVolumetricWeight float64  `json:"shipment_volumetric_weight"`
            ShipmentChargeableWeight float64  `json:"shipment_chargeable_weight"`
            FulfilmentPriorityText string  `json:"fulfilment_priority_text"`
            Formatted Formatted  `json:"formatted"`
            AutoTriggerDpAssignmentAcf bool  `json:"auto_trigger_dp_assignment_acf"`
         
    }
    
    // Prices used by Order
    type Prices struct {

        
            RefundCredit float64  `json:"refund_credit"`
            TransferPrice float64  `json:"transfer_price"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            CodCharges float64  `json:"cod_charges"`
            DeliveryCharge float64  `json:"delivery_charge"`
            PmPriceSplit float64  `json:"pm_price_split"`
            PromotionEffectiveDiscount float64  `json:"promotion_effective_discount"`
            Discount float64  `json:"discount"`
            FyndCredits float64  `json:"fynd_credits"`
            GiftPrice float64  `json:"gift_price"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            Cashback float64  `json:"cashback"`
            PriceMarked float64  `json:"price_marked"`
            CouponValue float64  `json:"coupon_value"`
            AmountPaid float64  `json:"amount_paid"`
            RefundAmount float64  `json:"refund_amount"`
            ValueOfGood float64  `json:"value_of_good"`
            AmountPaidRoundoff float64  `json:"amount_paid_roundoff"`
            CashbackApplied float64  `json:"cashback_applied"`
            PriceEffective float64  `json:"price_effective"`
         
    }
    
    // ShipmentItemFulFillingStore used by Order
    type ShipmentItemFulFillingStore struct {

        
            City string  `json:"city"`
            BrandStoreTags string  `json:"brand_store_tags"`
            Meta map[string]interface{}  `json:"meta"`
            LocationType string  `json:"location_type"`
            Name string  `json:"name"`
            State string  `json:"state"`
            Address string  `json:"address"`
            StoreEmail string  `json:"store_email"`
            Pincode string  `json:"pincode"`
            ID float64  `json:"id"`
            Code string  `json:"code"`
            Phone string  `json:"phone"`
         
    }
    
    // UserDataInfo used by Order
    type UserDataInfo struct {

        
            Mobile string  `json:"mobile"`
            Email string  `json:"email"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            AvisUserID string  `json:"avis_user_id"`
            LastName string  `json:"last_name"`
            FirstName string  `json:"first_name"`
            IsAnonymousUser bool  `json:"is_anonymous_user"`
            Gender string  `json:"gender"`
         
    }
    
    // PlatformDeliveryAddress used by Order
    type PlatformDeliveryAddress struct {

        
            Landmark string  `json:"landmark"`
            Area string  `json:"area"`
            Version string  `json:"version"`
            Latitude float64  `json:"latitude"`
            AddressCategory string  `json:"address_category"`
            Email string  `json:"email"`
            City string  `json:"city"`
            Address1 string  `json:"address1"`
            Address2 string  `json:"address2"`
            Longitude float64  `json:"longitude"`
            AddressType string  `json:"address_type"`
            State string  `json:"state"`
            Pincode string  `json:"pincode"`
            CreatedAt string  `json:"created_at"`
            UpdatedAt string  `json:"updated_at"`
            ContactPerson string  `json:"contact_person"`
            Phone string  `json:"phone"`
            Country string  `json:"country"`
         
    }
    
    // Dates used by Order
    type Dates struct {

        
            OrderCreated string  `json:"order_created"`
            DeliveryDate string  `json:"delivery_date"`
         
    }
    
    // BagReturnableCancelableStatus used by Order
    type BagReturnableCancelableStatus struct {

        
            IsCustomerReturnAllowed bool  `json:"is_customer_return_allowed"`
            IsReturnable bool  `json:"is_returnable"`
            EnableTracking bool  `json:"enable_tracking"`
            CanBeCancelled bool  `json:"can_be_cancelled"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // BagStateMapper used by Order
    type BagStateMapper struct {

        
            DisplayName string  `json:"display_name"`
            JourneyType string  `json:"journey_type"`
            Name string  `json:"name"`
            AppDisplayName string  `json:"app_display_name"`
            StateType string  `json:"state_type"`
            AppFacing bool  `json:"app_facing"`
            ID float64  `json:"id"`
            IsActive bool  `json:"is_active"`
            AppStateName string  `json:"app_state_name"`
            NotifyCustomer bool  `json:"notify_customer"`
         
    }
    
    // BagStatusHistory used by Order
    type BagStatusHistory struct {

        
            DisplayName string  `json:"display_name"`
            Reasons []map[string]interface{}  `json:"reasons"`
            ShipmentID string  `json:"shipment_id"`
            Forward bool  `json:"forward"`
            StoreID float64  `json:"store_id"`
            DeliveryAwbNumber string  `json:"delivery_awb_number"`
            AppDisplayName string  `json:"app_display_name"`
            BshID float64  `json:"bsh_id"`
            KafkaSync bool  `json:"kafka_sync"`
            DeliveryPartnerID float64  `json:"delivery_partner_id"`
            UpdatedAt string  `json:"updated_at"`
            StateType string  `json:"state_type"`
            CreatedAt string  `json:"created_at"`
            BagID float64  `json:"bag_id"`
            BagStateMapper BagStateMapper  `json:"bag_state_mapper"`
            Status string  `json:"status"`
            StateID float64  `json:"state_id"`
         
    }
    
    // Identifier used by Order
    type Identifier struct {

        
            Alu string  `json:"alu"`
            Isbn string  `json:"isbn"`
            Upc string  `json:"upc"`
            Ean string  `json:"ean"`
            SkuCode string  `json:"sku_code"`
         
    }
    
    // FinancialBreakup used by Order
    type FinancialBreakup struct {

        
            RefundCredit float64  `json:"refund_credit"`
            TotalUnits float64  `json:"total_units"`
            TransferPrice float64  `json:"transfer_price"`
            ItemName string  `json:"item_name"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            CodCharges float64  `json:"cod_charges"`
            DeliveryCharge float64  `json:"delivery_charge"`
            PromotionEffectiveDiscount float64  `json:"promotion_effective_discount"`
            GstFee float64  `json:"gst_fee"`
            Discount float64  `json:"discount"`
            FyndCredits float64  `json:"fynd_credits"`
            Identifiers Identifier  `json:"identifiers"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            Cashback float64  `json:"cashback"`
            Size string  `json:"size"`
            PriceMarked float64  `json:"price_marked"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            CouponValue float64  `json:"coupon_value"`
            AddedToFyndCash bool  `json:"added_to_fynd_cash"`
            AmountPaid float64  `json:"amount_paid"`
            GstTag string  `json:"gst_tag"`
            HsnCode string  `json:"hsn_code"`
            ValueOfGood float64  `json:"value_of_good"`
            AmountPaidRoundoff float64  `json:"amount_paid_roundoff"`
            CashbackApplied float64  `json:"cashback_applied"`
            PriceEffective float64  `json:"price_effective"`
         
    }
    
    // PlatformArticleAttributes used by Order
    type PlatformArticleAttributes struct {

        
            Currency string  `json:"currency"`
         
    }
    
    // PlatformItem used by Order
    type PlatformItem struct {

        
            Meta map[string]interface{}  `json:"meta"`
            SlugKey string  `json:"slug_key"`
            Image []string  `json:"image"`
            Color string  `json:"color"`
            CanReturn bool  `json:"can_return"`
            L3Category float64  `json:"l3_category"`
            ID float64  `json:"id"`
            L3CategoryName string  `json:"l3_category_name"`
            CanCancel bool  `json:"can_cancel"`
            L1Category []string  `json:"l1_category"`
            Size string  `json:"size"`
            Attributes PlatformArticleAttributes  `json:"attributes"`
            LastUpdatedAt string  `json:"last_updated_at"`
            DepartmentID float64  `json:"department_id"`
            Name string  `json:"name"`
            Brand string  `json:"brand"`
            BranchURL string  `json:"branch_url"`
            Code string  `json:"code"`
            L2Category []string  `json:"l2_category"`
            BrandID float64  `json:"brand_id"`
            Images []string  `json:"images"`
         
    }
    
    // ShipmentListingBrand used by Order
    type ShipmentListingBrand struct {

        
            Name string  `json:"name"`
            CreatedOn string  `json:"created_on"`
            Logo string  `json:"logo"`
            LogoBase64 string  `json:"logo_base64"`
         
    }
    
    // Dimensions used by Order
    type Dimensions struct {

        
            IsDefault bool  `json:"is_default"`
            Length float64  `json:"length"`
            Unit string  `json:"unit"`
            Width float64  `json:"width"`
            Height float64  `json:"height"`
         
    }
    
    // Weight used by Order
    type Weight struct {

        
            Shipping float64  `json:"shipping"`
            IsDefault bool  `json:"is_default"`
            Unit string  `json:"unit"`
         
    }
    
    // ReturnConfig used by Order
    type ReturnConfig struct {

        
            Time float64  `json:"time"`
            Returnable bool  `json:"returnable"`
            Unit string  `json:"unit"`
         
    }
    
    // Article used by Order
    type Article struct {

        
            Size string  `json:"size"`
            Dimensions Dimensions  `json:"dimensions"`
            EspModified bool  `json:"esp_modified"`
            ID string  `json:"_id"`
            UID string  `json:"uid"`
            Currency map[string]interface{}  `json:"currency"`
            SellerIdentifier string  `json:"seller_identifier"`
            RawMeta string  `json:"raw_meta"`
            IsSet bool  `json:"is_set"`
            ASet map[string]interface{}  `json:"a_set"`
            ChildDetails map[string]interface{}  `json:"child_details"`
            Identifiers map[string]interface{}  `json:"identifiers"`
            Weight Weight  `json:"weight"`
            Code string  `json:"code"`
            ReturnConfig ReturnConfig  `json:"return_config"`
         
    }
    
    // ReplacementDetails used by Order
    type ReplacementDetails struct {

        
            ReplacementType string  `json:"replacement_type"`
            OriginalAffiliateOrderID string  `json:"original_affiliate_order_id"`
         
    }
    
    // AffiliateMeta used by Order
    type AffiliateMeta struct {

        
            IsPriority bool  `json:"is_priority"`
            Quantity float64  `json:"quantity"`
            DueDate string  `json:"due_date"`
            SizeLevelTotalQty float64  `json:"size_level_total_qty"`
            ReplacementDetails ReplacementDetails  `json:"replacement_details"`
            ChannelShipmentID string  `json:"channel_shipment_id"`
            OrderItemID string  `json:"order_item_id"`
            MarketplaceInvoiceID string  `json:"marketplace_invoice_id"`
            EmployeeDiscount float64  `json:"employee_discount"`
            BoxType string  `json:"box_type"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            ChannelOrderID string  `json:"channel_order_id"`
            CouponCode string  `json:"coupon_code"`
         
    }
    
    // AffiliateBagDetails used by Order
    type AffiliateBagDetails struct {

        
            AffiliateOrderID string  `json:"affiliate_order_id"`
            AffiliateMeta AffiliateMeta  `json:"affiliate_meta"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
            EmployeeDiscount float64  `json:"employee_discount"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
         
    }
    
    // GSTDetailsData used by Order
    type GSTDetailsData struct {

        
            IsDefaultHsnCode bool  `json:"is_default_hsn_code"`
            CgstGstFee string  `json:"cgst_gst_fee"`
            HsnCode string  `json:"hsn_code"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            ValueOfGood float64  `json:"value_of_good"`
            SgstTaxPercentage float64  `json:"sgst_tax_percentage"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            CgstTaxPercentage float64  `json:"cgst_tax_percentage"`
            IgstTaxPercentage float64  `json:"igst_tax_percentage"`
            IgstGstFee string  `json:"igst_gst_fee"`
            SgstGstFee string  `json:"sgst_gst_fee"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            GstFee float64  `json:"gst_fee"`
            GstTag string  `json:"gst_tag"`
            GstinCode string  `json:"gstin_code"`
            HsnCodeID string  `json:"hsn_code_id"`
         
    }
    
    // BagUnit used by Order
    type BagUnit struct {

        
            DisplayName string  `json:"display_name"`
            Reasons []map[string]interface{}  `json:"reasons"`
            Meta map[string]interface{}  `json:"meta"`
            Dates Dates  `json:"dates"`
            Prices Prices  `json:"prices"`
            ProductQuantity float64  `json:"product_quantity"`
            BagID float64  `json:"bag_id"`
            Status BagReturnableCancelableStatus  `json:"status"`
            CurrentOperationalStatus BagStatusHistory  `json:"current_operational_status"`
            CanReturn bool  `json:"can_return"`
            FinancialBreakup []FinancialBreakup  `json:"financial_breakup"`
            Item PlatformItem  `json:"item"`
            EntityType string  `json:"entity_type"`
            CanCancel bool  `json:"can_cancel"`
            Size string  `json:"size"`
            BagType string  `json:"bag_type"`
            CurrentStatus BagStatusHistory  `json:"current_status"`
            Brand ShipmentListingBrand  `json:"brand"`
            BagExpiryDate string  `json:"bag_expiry_date"`
            Article Article  `json:"article"`
            AffiliateBagDetails AffiliateBagDetails  `json:"affiliate_bag_details"`
            Gst GSTDetailsData  `json:"gst"`
            LineNumber float64  `json:"line_number"`
            BagStatus []BagStatusHistory  `json:"bag_status"`
         
    }
    
    // ShipmentListingChannel used by Order
    type ShipmentListingChannel struct {

        
            ChannelShipmentID string  `json:"channel_shipment_id"`
            Name string  `json:"name"`
            IsAffiliate bool  `json:"is_affiliate"`
            Logo string  `json:"logo"`
         
    }
    
    // ShipmentStatus used by Order
    type ShipmentStatus struct {

        
            ShipmentStatusID float64  `json:"shipment_status_id"`
            StatusCreatedAt string  `json:"status_created_at"`
            BagList []string  `json:"bag_list"`
            ShipmentID string  `json:"shipment_id"`
            Meta map[string]interface{}  `json:"meta"`
            CreatedAt string  `json:"created_at"`
            Status string  `json:"status"`
            Title string  `json:"title"`
            CurrentShipmentStatus string  `json:"current_shipment_status"`
         
    }
    
    // ShipmentItem used by Order
    type ShipmentItem struct {

        
            LockStatus bool  `json:"lock_status"`
            DisplayName string  `json:"display_name"`
            StatusCreatedAt string  `json:"status_created_at"`
            Meta ShipmentItemMeta  `json:"meta"`
            Prices Prices  `json:"prices"`
            PaymentMethods map[string]interface{}  `json:"payment_methods"`
            FulfillingStore ShipmentItemFulFillingStore  `json:"fulfilling_store"`
            User UserDataInfo  `json:"user"`
            EstimatedSlaTime string  `json:"estimated_sla_time"`
            CustomerNote string  `json:"customer_note"`
            DeliveryAddress PlatformDeliveryAddress  `json:"delivery_address"`
            Bags []BagUnit  `json:"bags"`
            PaymentMode string  `json:"payment_mode"`
            PreviousShipmentID string  `json:"previous_shipment_id"`
            OrderingChannnel string  `json:"ordering_channnel"`
            CanProcess bool  `json:"can_process"`
            ShipmentCreatedAt string  `json:"shipment_created_at"`
            InvoiceID string  `json:"invoice_id"`
            TotalBags float64  `json:"total_bags"`
            ShipmentID string  `json:"shipment_id"`
            Channel ShipmentListingChannel  `json:"channel"`
            OrderDate string  `json:"order_date"`
            ShipmentStatus ShipmentStatus  `json:"shipment_status"`
            OrderID string  `json:"order_id"`
         
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

        
            Success bool  `json:"success"`
            TotalCount float64  `json:"total_count"`
            Lane string  `json:"lane"`
            Items []ShipmentItem  `json:"items"`
            Page Page  `json:"page"`
            Message string  `json:"message"`
         
    }
    
    // Error used by Order
    type Error struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // BuyerDetails used by Order
    type BuyerDetails struct {

        
            City string  `json:"city"`
            Name string  `json:"name"`
            State string  `json:"state"`
            Address string  `json:"address"`
            Pincode float64  `json:"pincode"`
            AjioSiteID string  `json:"ajio_site_id"`
            Gstin string  `json:"gstin"`
         
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
    
    // CreditNoteData used by Order
    type CreditNoteData struct {

        
            DateOfIssue string  `json:"date_of_issue"`
            ExpiryDate string  `json:"expiry_date"`
            TotalBalance float64  `json:"total_balance"`
            Validity float64  `json:"validity"`
         
    }
    
    // ShipmentMeta used by Order
    type ShipmentMeta struct {

        
            External map[string]interface{}  `json:"external"`
            DpID string  `json:"dp_id"`
            PoNumber string  `json:"po_number"`
            AwbNumber string  `json:"awb_number"`
            B2bBuyerDetails BuyerDetails  `json:"b2b_buyer_details"`
            ReturnAffiliateShipmentID string  `json:"return_affiliate_shipment_id"`
            ShipmentWeight float64  `json:"shipment_weight"`
            Weight float64  `json:"weight"`
            ReturnAwbNumber string  `json:"return_awb_number"`
            ParentDpID string  `json:"parent_dp_id"`
            ForwardAffiliateShipmentID string  `json:"forward_affiliate_shipment_id"`
            RefundTo string  `json:"refund_to"`
            ReturnAffiliateOrderID string  `json:"return_affiliate_order_id"`
            BagWeight map[string]interface{}  `json:"bag_weight"`
            EwaybillInfo map[string]interface{}  `json:"ewaybill_info"`
            Timestamp ShipmentTimeStamp  `json:"timestamp"`
            Dimension Dimensions  `json:"dimension"`
            OrderType string  `json:"order_type"`
            LockData LockData  `json:"lock_data"`
            StoreInvoiceUpdatedDate string  `json:"store_invoice_updated_date"`
            DebugInfo DebugInfo  `json:"debug_info"`
            ShipmentTags []ShipmentTags  `json:"shipment_tags"`
            BoxType string  `json:"box_type"`
            ReturnStoreNode float64  `json:"return_store_node"`
            DpOptions map[string]interface{}  `json:"dp_options"`
            MarketplaceStoreID string  `json:"marketplace_store_id"`
            PackagingName string  `json:"packaging_name"`
            ForwardAffiliateOrderID string  `json:"forward_affiliate_order_id"`
            AssignDpFromSb bool  `json:"assign_dp_from_sb"`
            DpSortKey string  `json:"dp_sort_key"`
            DpName string  `json:"dp_name"`
            SameStoreAvailable bool  `json:"same_store_available"`
            EinvoiceInfo EinvoiceInfo  `json:"einvoice_info"`
            ReturnDetails map[string]interface{}  `json:"return_details"`
            DueDate string  `json:"due_date"`
            CreditNoteData CreditNoteData  `json:"credit_note_data"`
            ShipmentVolumetricWeight float64  `json:"shipment_volumetric_weight"`
            FulfilmentPriorityText string  `json:"fulfilment_priority_text"`
            Formatted Formatted  `json:"formatted"`
            B2cBuyerDetails map[string]interface{}  `json:"b2c_buyer_details"`
            AutoTriggerDpAssignmentAcf bool  `json:"auto_trigger_dp_assignment_acf"`
         
    }
    
    // TrackingList used by Order
    type TrackingList struct {

        
            Text string  `json:"text"`
            IsCurrent bool  `json:"is_current"`
            IsPassed bool  `json:"is_passed"`
            Status string  `json:"status"`
            Time string  `json:"time"`
         
    }
    
    // OrderingStoreDetails used by Order
    type OrderingStoreDetails struct {

        
            StoreName string  `json:"store_name"`
            City string  `json:"city"`
            Meta map[string]interface{}  `json:"meta"`
            State string  `json:"state"`
            Address string  `json:"address"`
            Pincode string  `json:"pincode"`
            ID float64  `json:"id"`
            ContactPerson string  `json:"contact_person"`
            Code string  `json:"code"`
            Phone string  `json:"phone"`
            Country string  `json:"country"`
         
    }
    
    // UserDetailsData used by Order
    type UserDetailsData struct {

        
            Landmark string  `json:"landmark"`
            Area string  `json:"area"`
            Email string  `json:"email"`
            City string  `json:"city"`
            Address1 string  `json:"address1"`
            Name string  `json:"name"`
            AddressType string  `json:"address_type"`
            State string  `json:"state"`
            Address string  `json:"address"`
            Pincode string  `json:"pincode"`
            Phone string  `json:"phone"`
            Country string  `json:"country"`
         
    }
    
    // ShipmentStatusData used by Order
    type ShipmentStatusData struct {

        
            DisplayName string  `json:"display_name"`
            BagList []string  `json:"bag_list"`
            ShipmentID string  `json:"shipment_id"`
            Meta map[string]interface{}  `json:"meta"`
            CreatedAt string  `json:"created_at"`
            ID float64  `json:"id"`
            Status string  `json:"status"`
         
    }
    
    // InvoiceInfo used by Order
    type InvoiceInfo struct {

        
            StoreInvoiceID string  `json:"store_invoice_id"`
            LabelURL string  `json:"label_url"`
            CreditNoteID string  `json:"credit_note_id"`
            ExternalInvoiceID string  `json:"external_invoice_id"`
            InvoiceURL string  `json:"invoice_url"`
            UpdatedDate string  `json:"updated_date"`
         
    }
    
    // FulfillingStore used by Order
    type FulfillingStore struct {

        
            StoreName string  `json:"store_name"`
            City string  `json:"city"`
            Meta map[string]interface{}  `json:"meta"`
            State string  `json:"state"`
            Address string  `json:"address"`
            Pincode string  `json:"pincode"`
            ID float64  `json:"id"`
            FulfillmentChannel string  `json:"fulfillment_channel"`
            ContactPerson string  `json:"contact_person"`
            Code string  `json:"code"`
            Phone string  `json:"phone"`
            Country string  `json:"country"`
         
    }
    
    // GiftCard used by Order
    type GiftCard struct {

        
            DisplayText string  `json:"display_text"`
            IsGiftApplied bool  `json:"is_gift_applied"`
            GiftPrice float64  `json:"gift_price"`
            GiftMessage string  `json:"gift_message"`
         
    }
    
    // B2BPODetails used by Order
    type B2BPODetails struct {

        
            TotalGstPercentage float64  `json:"total_gst_percentage"`
            PoTaxAmount float64  `json:"po_tax_amount"`
            DockerNumber string  `json:"docker_number"`
            ItemBasePrice float64  `json:"item_base_price"`
            PartialCanRet bool  `json:"partial_can_ret"`
            PoLineAmount float64  `json:"po_line_amount"`
         
    }
    
    // BagMeta used by Order
    type BagMeta struct {

        
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            GroupID string  `json:"group_id"`
            CustomMessage string  `json:"custom_message"`
            CustomJson map[string]interface{}  `json:"custom_json"`
            GiftCard GiftCard  `json:"gift_card"`
            B2bPoDetails B2BPODetails  `json:"b2b_po_details"`
            PartialCanRet bool  `json:"partial_can_ret"`
            DocketNumber string  `json:"docket_number"`
         
    }
    
    // BagPaymentMethods used by Order
    type BagPaymentMethods struct {

        
            Amount float64  `json:"amount"`
            Mode string  `json:"mode"`
         
    }
    
    // BagConfigs used by Order
    type BagConfigs struct {

        
            IsCustomerReturnAllowed bool  `json:"is_customer_return_allowed"`
            IsReturnable bool  `json:"is_returnable"`
            AllowForceReturn bool  `json:"allow_force_return"`
            EnableTracking bool  `json:"enable_tracking"`
            CanBeCancelled bool  `json:"can_be_cancelled"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // CurrentStatus used by Order
    type CurrentStatus struct {

        
            ShipmentID string  `json:"shipment_id"`
            StoreID float64  `json:"store_id"`
            DeliveryAwbNumber string  `json:"delivery_awb_number"`
            KafkaSync bool  `json:"kafka_sync"`
            StateType string  `json:"state_type"`
            CreatedAt string  `json:"created_at"`
            DeliveryPartnerID float64  `json:"delivery_partner_id"`
            ID float64  `json:"id"`
            UpdatedAt float64  `json:"updated_at"`
            BagID float64  `json:"bag_id"`
            BagStateMapper BagStateMapper  `json:"bag_state_mapper"`
            Status string  `json:"status"`
            StateID float64  `json:"state_id"`
         
    }
    
    // OrderBrandName used by Order
    type OrderBrandName struct {

        
            Company float64  `json:"company"`
            CreatedOn string  `json:"created_on"`
            BrandName string  `json:"brand_name"`
            ModifiedOn string  `json:"modified_on"`
            ID float64  `json:"id"`
            Logo string  `json:"logo"`
         
    }
    
    // ReturnConfig1 used by Order
    type ReturnConfig1 struct {

        
            Time float64  `json:"time"`
            Returnable bool  `json:"returnable"`
            Unit string  `json:"unit"`
         
    }
    
    // OrderBagArticle used by Order
    type OrderBagArticle struct {

        
            UID string  `json:"uid"`
            Size string  `json:"size"`
            Identifiers map[string]interface{}  `json:"identifiers"`
            ReturnConfig ReturnConfig1  `json:"return_config"`
         
    }
    
    // AffiliateBagsDetails used by Order
    type AffiliateBagsDetails struct {

        
            AffiliateBagID string  `json:"affiliate_bag_id"`
            CouponCode string  `json:"coupon_code"`
         
    }
    
    // BagGST used by Order
    type BagGST struct {

        
            IsDefaultHsnCode bool  `json:"is_default_hsn_code"`
            CgstGstFee string  `json:"cgst_gst_fee"`
            HsnCode string  `json:"hsn_code"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            ValueOfGood float64  `json:"value_of_good"`
            SgstTaxPercentage float64  `json:"sgst_tax_percentage"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            CgstTaxPercentage float64  `json:"cgst_tax_percentage"`
            IgstTaxPercentage float64  `json:"igst_tax_percentage"`
            IgstGstFee string  `json:"igst_gst_fee"`
            SgstGstFee string  `json:"sgst_gst_fee"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            GstFee float64  `json:"gst_fee"`
            GstTag string  `json:"gst_tag"`
            GstinCode string  `json:"gstin_code"`
            HsnCodeID string  `json:"hsn_code_id"`
         
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

        
            MrpPromotion bool  `json:"mrp_promotion"`
            PromotionName string  `json:"promotion_name"`
            ArticleQuantity float64  `json:"article_quantity"`
            DiscountRules []DiscountRules  `json:"discount_rules"`
            Amount float64  `json:"amount"`
            PromoID string  `json:"promo_id"`
            BuyRules []BuyRules  `json:"buy_rules"`
            PromotionType string  `json:"promotion_type"`
         
    }
    
    // OrderBags used by Order
    type OrderBags struct {

        
            GroupID string  `json:"group_id"`
            DisplayName string  `json:"display_name"`
            Quantity float64  `json:"quantity"`
            Meta BagMeta  `json:"meta"`
            Prices Prices  `json:"prices"`
            PaymentMethods []BagPaymentMethods  `json:"payment_methods"`
            BagID float64  `json:"bag_id"`
            CanReturn bool  `json:"can_return"`
            ParentPromoBags map[string]interface{}  `json:"parent_promo_bags"`
            FinancialBreakup FinancialBreakup  `json:"financial_breakup"`
            Item PlatformItem  `json:"item"`
            DeliveryAddress PlatformDeliveryAddress  `json:"delivery_address"`
            EntityType string  `json:"entity_type"`
            BagConfigs BagConfigs  `json:"bag_configs"`
            CanCancel bool  `json:"can_cancel"`
            CurrentStatus CurrentStatus  `json:"current_status"`
            Brand OrderBrandName  `json:"brand"`
            Article OrderBagArticle  `json:"article"`
            AffiliateBagDetails AffiliateBagsDetails  `json:"affiliate_bag_details"`
            GstDetails BagGST  `json:"gst_details"`
            Identifier string  `json:"identifier"`
            SellerIdentifier string  `json:"seller_identifier"`
            LineNumber float64  `json:"line_number"`
            AppliedPromos []AppliedPromos  `json:"applied_promos"`
         
    }
    
    // ShipmentPayments used by Order
    type ShipmentPayments struct {

        
            Source string  `json:"source"`
            Logo string  `json:"logo"`
            Mode string  `json:"mode"`
         
    }
    
    // PDFLinks used by Order
    type PDFLinks struct {

        
            LabelA4 string  `json:"label_a4"`
            InvoiceExport string  `json:"invoice_export"`
            DeliveryChallanA4 string  `json:"delivery_challan_a4"`
            LabelType string  `json:"label_type"`
            Label string  `json:"label"`
            LabelA6 string  `json:"label_a6"`
            CreditNoteURL string  `json:"credit_note_url"`
            InvoiceA4 string  `json:"invoice_a4"`
            LabelPos string  `json:"label_pos"`
            InvoiceType string  `json:"invoice_type"`
            InvoicePos string  `json:"invoice_pos"`
            LabelExport string  `json:"label_export"`
            Invoice string  `json:"invoice"`
            InvoiceA6 string  `json:"invoice_a6"`
            PoInvoice string  `json:"po_invoice"`
            B2b string  `json:"b2b"`
         
    }
    
    // AffiliateDetails used by Order
    type AffiliateDetails struct {

        
            AffiliateOrderID string  `json:"affiliate_order_id"`
            PdfLinks PDFLinks  `json:"pdf_links"`
            AdID string  `json:"ad_id"`
            CompanyAffiliateTag string  `json:"company_affiliate_tag"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
            ShipmentMeta ShipmentMeta  `json:"shipment_meta"`
            AffiliateID string  `json:"affiliate_id"`
            AffiliateMeta AffiliateMeta  `json:"affiliate_meta"`
            AffiliateStoreID string  `json:"affiliate_store_id"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
         
    }
    
    // PDFLink used by Order
    type PDFLink struct {

        
            InvoiceA4 string  `json:"invoice_a4"`
            CreditNoteA4 string  `json:"credit_note_a4"`
         
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

        
            CompanyID float64  `json:"company_id"`
            CompanyGst string  `json:"company_gst"`
            CompanyName string  `json:"company_name"`
            Address map[string]interface{}  `json:"address"`
            CompanyContact ContactDetails  `json:"company_contact"`
            CompanyCin string  `json:"company_cin"`
         
    }
    
    // OrderDetailsData used by Order
    type OrderDetailsData struct {

        
            OrderingChannel string  `json:"ordering_channel"`
            OrderDate string  `json:"order_date"`
            TaxDetails map[string]interface{}  `json:"tax_details"`
            OrderValue string  `json:"order_value"`
            Source string  `json:"source"`
            CodCharges string  `json:"cod_charges"`
            OrderingChannelLogo map[string]interface{}  `json:"ordering_channel_logo"`
            AffiliateID string  `json:"affiliate_id"`
            FyndOrderID string  `json:"fynd_order_id"`
         
    }
    
    // DPDetailsData used by Order
    type DPDetailsData struct {

        
            Name string  `json:"name"`
            EwayBillID string  `json:"eway_bill_id"`
            TrackURL string  `json:"track_url"`
            AwbNo string  `json:"awb_no"`
            Pincode string  `json:"pincode"`
            ID float64  `json:"id"`
            GstTag string  `json:"gst_tag"`
            Country string  `json:"country"`
         
    }
    
    // ShipmentDetails used by Order
    type ShipmentDetails struct {

        
            LockMessage string  `json:"lock_message"`
            LockStatus bool  `json:"lock_status"`
            ActionToStatus map[string]interface{}  `json:"action_to_status"`
         
    }
    
    // PlatformShipment used by Order
    type PlatformShipment struct {

        
            LockStatus bool  `json:"lock_status"`
            Meta ShipmentMeta  `json:"meta"`
            TrackingList []TrackingList  `json:"tracking_list"`
            ShipmentImages []string  `json:"shipment_images"`
            DeliverySlot map[string]interface{}  `json:"delivery_slot"`
            Coupon map[string]interface{}  `json:"coupon"`
            Prices Prices  `json:"prices"`
            OrderingStore OrderingStoreDetails  `json:"ordering_store"`
            DeliveryDetails UserDetailsData  `json:"delivery_details"`
            PackagingType string  `json:"packaging_type"`
            PickedDate string  `json:"picked_date"`
            Status ShipmentStatusData  `json:"status"`
            EnableDpTracking bool  `json:"enable_dp_tracking"`
            Invoice InvoiceInfo  `json:"invoice"`
            PaymentMethods map[string]interface{}  `json:"payment_methods"`
            FulfillingStore FulfillingStore  `json:"fulfilling_store"`
            ForwardShipmentID string  `json:"forward_shipment_id"`
            User UserDataInfo  `json:"user"`
            EstimatedSlaTime string  `json:"estimated_sla_time"`
            UserAgent string  `json:"user_agent"`
            Bags []OrderBags  `json:"bags"`
            PlatformLogo string  `json:"platform_logo"`
            Vertical string  `json:"vertical"`
            Payments ShipmentPayments  `json:"payments"`
            PaymentMode string  `json:"payment_mode"`
            FulfilmentPriority float64  `json:"fulfilment_priority"`
            CanUpdateDimension bool  `json:"can_update_dimension"`
            JourneyType string  `json:"journey_type"`
            AffiliateDetails AffiliateDetails  `json:"affiliate_details"`
            PdfLinks PDFLink  `json:"pdf_links"`
            CompanyDetails CompanyDetails  `json:"company_details"`
            Order OrderDetailsData  `json:"order"`
            ShipmentQuantity float64  `json:"shipment_quantity"`
            TotalItems float64  `json:"total_items"`
            ShipmentCreatedAt string  `json:"shipment_created_at"`
            TotalBags float64  `json:"total_bags"`
            InvoiceID string  `json:"invoice_id"`
            GstDetails GSTDetailsData  `json:"gst_details"`
            CustomMeta []map[string]interface{}  `json:"custom_meta"`
            DpAssignment bool  `json:"dp_assignment"`
            PriorityText string  `json:"priority_text"`
            DpDetails DPDetailsData  `json:"dp_details"`
            BagStatusHistory []BagStatusHistory  `json:"bag_status_history"`
            ShipmentID string  `json:"shipment_id"`
            BillingDetails UserDetailsData  `json:"billing_details"`
            ShipmentDetails ShipmentDetails  `json:"shipment_details"`
            OperationalStatus string  `json:"operational_status"`
            IsDpAssignEnabled bool  `json:"is_dp_assign_enabled"`
            ShipmentStatus string  `json:"shipment_status"`
            CustomMessage string  `json:"custom_message"`
         
    }
    
    // ShipmentInfoResponse used by Order
    type ShipmentInfoResponse struct {

        
            Message string  `json:"message"`
            Shipments []PlatformShipment  `json:"shipments"`
            Success bool  `json:"success"`
         
    }
    
    // PlatformUserDetails used by Order
    type PlatformUserDetails struct {

        
            PlatformUserEmployeeCode string  `json:"platform_user_employee_code"`
            PlatformUserFirstName string  `json:"platform_user_first_name"`
            PlatformUserLastName string  `json:"platform_user_last_name"`
            PlatformUserID string  `json:"platform_user_id"`
         
    }
    
    // BillingStaffDetails used by Order
    type BillingStaffDetails struct {

        
            User string  `json:"user"`
            EmployeeCode string  `json:"employee_code"`
            StaffID float64  `json:"staff_id"`
            LastName string  `json:"last_name"`
            FirstName string  `json:"first_name"`
         
    }
    
    // TransactionData used by Order
    type TransactionData struct {

        
            Entity string  `json:"entity"`
            TransactionID string  `json:"transaction_id"`
            PaymentID string  `json:"payment_id"`
            Currency string  `json:"currency"`
            TerminalID string  `json:"terminal_id"`
            AmountPaid float64  `json:"amount_paid"`
            UniqueReferenceNumber string  `json:"unique_reference_number"`
            Status string  `json:"status"`
         
    }
    
    // OrderMeta used by Order
    type OrderMeta struct {

        
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            Staff map[string]interface{}  `json:"staff"`
            PlatformUserDetails PlatformUserDetails  `json:"platform_user_details"`
            OrderingStore float64  `json:"ordering_store"`
            BillingStaffDetails BillingStaffDetails  `json:"billing_staff_details"`
            TransactionData TransactionData  `json:"transaction_data"`
            CompanyLogo string  `json:"company_logo"`
            CustomerNote string  `json:"customer_note"`
            OrderType string  `json:"order_type"`
            OrderChildEntities []string  `json:"order_child_entities"`
            OrderTags []map[string]interface{}  `json:"order_tags"`
            CartID float64  `json:"cart_id"`
            Comment string  `json:"comment"`
            PaymentType string  `json:"payment_type"`
            OrderPlatform string  `json:"order_platform"`
            MongoCartID float64  `json:"mongo_cart_id"`
            CurrencySymbol string  `json:"currency_symbol"`
            EmployeeID float64  `json:"employee_id"`
            Files []map[string]interface{}  `json:"files"`
         
    }
    
    // TaxDetails used by Order
    type TaxDetails struct {

        
            PanNo string  `json:"pan_no"`
            Gstin string  `json:"gstin"`
         
    }
    
    // OrderData used by Order
    type OrderData struct {

        
            Meta OrderMeta  `json:"meta"`
            OrderDate string  `json:"order_date"`
            TaxDetails TaxDetails  `json:"tax_details"`
            Prices Prices  `json:"prices"`
            PaymentMethods map[string]interface{}  `json:"payment_methods"`
            FyndOrderID string  `json:"fynd_order_id"`
         
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
            TotalItems float64  `json:"total_items"`
            Value string  `json:"value"`
            Index float64  `json:"index"`
            Actions []map[string]interface{}  `json:"actions"`
         
    }
    
    // SuperLane used by Order
    type SuperLane struct {

        
            Options []SubLane  `json:"options"`
            Text string  `json:"text"`
            TotalItems float64  `json:"total_items"`
            Value string  `json:"value"`
         
    }
    
    // LaneConfigResponse used by Order
    type LaneConfigResponse struct {

        
            SuperLanes []SuperLane  `json:"super_lanes"`
         
    }
    
    // PlatformBreakupValues used by Order
    type PlatformBreakupValues struct {

        
            Name string  `json:"name"`
            Display string  `json:"display"`
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
            BreakupValues []PlatformBreakupValues  `json:"breakup_values"`
            Meta map[string]interface{}  `json:"meta"`
            Channel PlatformChannel  `json:"channel"`
            UserInfo UserDataInfo  `json:"user_info"`
            OrderCreatedTime string  `json:"order_created_time"`
            OrderValue float64  `json:"order_value"`
            Shipments []PlatformShipment  `json:"shipments"`
            OrderID string  `json:"order_id"`
            PaymentMode string  `json:"payment_mode"`
         
    }
    
    // OrderListingResponse used by Order
    type OrderListingResponse struct {

        
            Success bool  `json:"success"`
            TotalCount float64  `json:"total_count"`
            Lane string  `json:"lane"`
            Items []PlatformOrderItems  `json:"items"`
            Page Page  `json:"page"`
            Message string  `json:"message"`
         
    }
    
    // PlatformTrack used by Order
    type PlatformTrack struct {

        
            Reason string  `json:"reason"`
            Meta map[string]interface{}  `json:"meta"`
            Awb string  `json:"awb"`
            RawStatus string  `json:"raw_status"`
            UpdatedAt string  `json:"updated_at"`
            AccountName string  `json:"account_name"`
            ShipmentType string  `json:"shipment_type"`
            Status string  `json:"status"`
            UpdatedTime string  `json:"updated_time"`
            LastLocationRecievedAt string  `json:"last_location_recieved_at"`
         
    }
    
    // PlatformShipmentTrack used by Order
    type PlatformShipmentTrack struct {

        
            Results []PlatformTrack  `json:"results"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // FilterInfoOption used by Order
    type FilterInfoOption struct {

        
            PlaceholderText string  `json:"placeholder_text"`
            Text string  `json:"text"`
            MinSearchSize float64  `json:"min_search_size"`
            Name string  `json:"name"`
            Value string  `json:"value"`
            ShowUI bool  `json:"show_ui"`
         
    }
    
    // FiltersInfo used by Order
    type FiltersInfo struct {

        
            PlaceholderText string  `json:"placeholder_text"`
            Text string  `json:"text"`
            Value string  `json:"value"`
            Required bool  `json:"required"`
            Type string  `json:"type"`
            Options []FilterInfoOption  `json:"options"`
         
    }
    
    // AdvanceFilterInfo used by Order
    type AdvanceFilterInfo struct {

        
            Unfulfilled []FiltersInfo  `json:"unfulfilled"`
            ActionCentre []FiltersInfo  `json:"action_centre"`
            Processed []FiltersInfo  `json:"processed"`
            Returned []FiltersInfo  `json:"returned"`
            Filters []FiltersInfo  `json:"filters"`
            AppliedFilters map[string]interface{}  `json:"applied_filters"`
            Page map[string]interface{}  `json:"page"`
         
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

        
            Text string  `json:"text"`
            Value string  `json:"value"`
         
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

        
            QcType []string  `json:"qc_type"`
            ID float64  `json:"id"`
            QuestionSet []QuestionSet  `json:"question_set"`
            DisplayName string  `json:"display_name"`
         
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

        
            Message string  `json:"message"`
            Reasons []ShipmentResponseReasons  `json:"reasons"`
            Success bool  `json:"success"`
         
    }
    
    // StoreEinvoice used by Order
    type StoreEinvoice struct {

        
            User string  `json:"user"`
            Username string  `json:"username"`
            Password string  `json:"password"`
            Enabled bool  `json:"enabled"`
         
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

        
            DsType string  `json:"ds_type"`
            Verified bool  `json:"verified"`
            Value string  `json:"value"`
            LegalName string  `json:"legal_name"`
            URL string  `json:"url"`
         
    }
    
    // StoreDocuments used by Order
    type StoreDocuments struct {

        
            Gst Document  `json:"gst"`
         
    }
    
    // EInvoicePortalDetails used by Order
    type EInvoicePortalDetails struct {

        
            User string  `json:"user"`
            Username string  `json:"username"`
            Password string  `json:"password"`
         
    }
    
    // StoreMeta used by Order
    type StoreMeta struct {

        
            DisplayName string  `json:"display_name"`
            GstCredentials StoreGstCredentials  `json:"gst_credentials"`
            Documents StoreDocuments  `json:"documents"`
            Stage string  `json:"stage"`
            GstNumber string  `json:"gst_number"`
            NotificationEmails []string  `json:"notification_emails"`
            EwaybillPortalDetails map[string]interface{}  `json:"ewaybill_portal_details"`
            Timing []map[string]interface{}  `json:"timing"`
            AdditionalContactDetails map[string]interface{}  `json:"additional_contact_details"`
            EinvoicePortalDetails EInvoicePortalDetails  `json:"einvoice_portal_details"`
            ProductReturnConfig map[string]interface{}  `json:"product_return_config"`
         
    }
    
    // StoreAddress used by Order
    type StoreAddress struct {

        
            CreatedAt string  `json:"created_at"`
            ContactPerson string  `json:"contact_person"`
            Phone string  `json:"phone"`
            Address2 string  `json:"address2"`
            Landmark string  `json:"landmark"`
            Longitude float64  `json:"longitude"`
            Pincode float64  `json:"pincode"`
            AddressCategory string  `json:"address_category"`
            CountryCode string  `json:"country_code"`
            UpdatedAt string  `json:"updated_at"`
            Area string  `json:"area"`
            Version string  `json:"version"`
            Latitude float64  `json:"latitude"`
            Address1 string  `json:"address1"`
            City string  `json:"city"`
            Email string  `json:"email"`
            AddressType string  `json:"address_type"`
            State string  `json:"state"`
            Country string  `json:"country"`
         
    }
    
    // Store used by Order
    type Store struct {

        
            BrandStoreTags []string  `json:"brand_store_tags"`
            Meta StoreMeta  `json:"meta"`
            LocationType string  `json:"location_type"`
            CreatedAt string  `json:"created_at"`
            ContactPerson string  `json:"contact_person"`
            Phone float64  `json:"phone"`
            IsEnabledForRecon bool  `json:"is_enabled_for_recon"`
            Address2 string  `json:"address2"`
            ParentStoreID float64  `json:"parent_store_id"`
            Longitude float64  `json:"longitude"`
            LoginUsername string  `json:"login_username"`
            SID string  `json:"s_id"`
            Pincode string  `json:"pincode"`
            MallArea string  `json:"mall_area"`
            MallName string  `json:"mall_name"`
            IsArchived bool  `json:"is_archived"`
            PackagingMaterialCount float64  `json:"packaging_material_count"`
            AlohomoraUserID float64  `json:"alohomora_user_id"`
            CompanyID float64  `json:"company_id"`
            Name string  `json:"name"`
            StoreEmail string  `json:"store_email"`
            UpdatedAt string  `json:"updated_at"`
            OrderIntegrationID string  `json:"order_integration_id"`
            IsActive bool  `json:"is_active"`
            Code string  `json:"code"`
            Latitude float64  `json:"latitude"`
            VatNo string  `json:"vat_no"`
            Address1 string  `json:"address1"`
            City string  `json:"city"`
            StoreAddressJson StoreAddress  `json:"store_address_json"`
            BrandID interface{}  `json:"brand_id"`
            State string  `json:"state"`
            StoreActiveFrom string  `json:"store_active_from"`
            FulfillmentChannel string  `json:"fulfillment_channel"`
            Country string  `json:"country"`
         
    }
    
    // BagReturnableCancelableStatus1 used by Order
    type BagReturnableCancelableStatus1 struct {

        
            IsCustomerReturnAllowed bool  `json:"is_customer_return_allowed"`
            IsReturnable bool  `json:"is_returnable"`
            EnableTracking bool  `json:"enable_tracking"`
            CanBeCancelled bool  `json:"can_be_cancelled"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // Attributes used by Order
    type Attributes struct {

        
            MarketerName string  `json:"marketer_name"`
            Name string  `json:"name"`
            PrimaryColor string  `json:"primary_color"`
            BrandName string  `json:"brand_name"`
            PrimaryMaterial string  `json:"primary_material"`
            Essential string  `json:"essential"`
            Gender []string  `json:"gender"`
            PrimaryColorHex string  `json:"primary_color_hex"`
            MarketerAddress string  `json:"marketer_address"`
         
    }
    
    // Item used by Order
    type Item struct {

        
            Meta map[string]interface{}  `json:"meta"`
            SlugKey string  `json:"slug_key"`
            Image []string  `json:"image"`
            Color string  `json:"color"`
            ItemID float64  `json:"item_id"`
            CanReturn bool  `json:"can_return"`
            L3Category float64  `json:"l3_category"`
            WebstoreProductURL string  `json:"webstore_product_url"`
            L3CategoryName string  `json:"l3_category_name"`
            CanCancel bool  `json:"can_cancel"`
            L1Category []string  `json:"l1_category"`
            Size string  `json:"size"`
            Attributes Attributes  `json:"attributes"`
            LastUpdatedAt string  `json:"last_updated_at"`
            DepartmentID float64  `json:"department_id"`
            Name string  `json:"name"`
            Brand string  `json:"brand"`
            BranchURL string  `json:"branch_url"`
            Code string  `json:"code"`
            L2Category []string  `json:"l2_category"`
            L2CategoryID float64  `json:"l2_category_id"`
            BrandID float64  `json:"brand_id"`
            Gender string  `json:"gender"`
            L1CategoryID float64  `json:"l1_category_id"`
         
    }
    
    // Brand used by Order
    type Brand struct {

        
            IsVirtualInvoice bool  `json:"is_virtual_invoice"`
            StartDate string  `json:"start_date"`
            Company string  `json:"company"`
            CreditNoteExpiryDays float64  `json:"credit_note_expiry_days"`
            ScriptLastRan string  `json:"script_last_ran"`
            BrandID float64  `json:"brand_id"`
            CreatedOn float64  `json:"created_on"`
            PickupLocation string  `json:"pickup_location"`
            BrandName string  `json:"brand_name"`
            ModifiedOn float64  `json:"modified_on"`
            InvoicePrefix string  `json:"invoice_prefix"`
            Logo string  `json:"logo"`
            CreditNoteAllowed bool  `json:"credit_note_allowed"`
         
    }
    
    // BagGSTDetails used by Order
    type BagGSTDetails struct {

        
            IsDefaultHsnCode bool  `json:"is_default_hsn_code"`
            CgstGstFee string  `json:"cgst_gst_fee"`
            HsnCode string  `json:"hsn_code"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            ValueOfGood float64  `json:"value_of_good"`
            SgstTaxPercentage float64  `json:"sgst_tax_percentage"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            CgstTaxPercentage float64  `json:"cgst_tax_percentage"`
            IgstTaxPercentage float64  `json:"igst_tax_percentage"`
            IgstGstFee string  `json:"igst_gst_fee"`
            SgstGstFee string  `json:"sgst_gst_fee"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            GstFee float64  `json:"gst_fee"`
            GstTag string  `json:"gst_tag"`
            GstinCode string  `json:"gstin_code"`
            HsnCodeID string  `json:"hsn_code_id"`
         
    }
    
    // ArticleDetails used by Order
    type ArticleDetails struct {

        
            Status map[string]interface{}  `json:"status"`
         
    }
    
    // BagDetailsPlatformResponse used by Order
    type BagDetailsPlatformResponse struct {

        
            DisplayName string  `json:"display_name"`
            Reasons []map[string]interface{}  `json:"reasons"`
            OriginalBagList []float64  `json:"original_bag_list"`
            Meta BagMeta  `json:"meta"`
            Quantity float64  `json:"quantity"`
            RestorePromos map[string]interface{}  `json:"restore_promos"`
            Dates Dates  `json:"dates"`
            Prices Prices  `json:"prices"`
            OrderingStore Store  `json:"ordering_store"`
            Status BagReturnableCancelableStatus1  `json:"status"`
            CurrentOperationalStatus BagStatusHistory  `json:"current_operational_status"`
            ParentPromoBags map[string]interface{}  `json:"parent_promo_bags"`
            FinancialBreakup []FinancialBreakup  `json:"financial_breakup"`
            Tags []string  `json:"tags"`
            Item Item  `json:"item"`
            ID float64  `json:"id"`
            EntityType string  `json:"entity_type"`
            BagUpdateTime float64  `json:"bag_update_time"`
            JourneyType string  `json:"journey_type"`
            AffiliateDetails AffiliateDetails  `json:"affiliate_details"`
            RestoreCoupon bool  `json:"restore_coupon"`
            CurrentStatus BagStatusHistory  `json:"current_status"`
            Brand Brand  `json:"brand"`
            OrderIntegrationID string  `json:"order_integration_id"`
            Article Article  `json:"article"`
            AffiliateBagDetails AffiliateBagDetails  `json:"affiliate_bag_details"`
            GstDetails BagGSTDetails  `json:"gst_details"`
            NoOfBagsOrder float64  `json:"no_of_bags_order"`
            BagStatusHistory BagStatusHistory  `json:"bag_status_history"`
            ArticleDetails ArticleDetails  `json:"article_details"`
            Identifier string  `json:"identifier"`
            ShipmentID string  `json:"shipment_id"`
            SellerIdentifier string  `json:"seller_identifier"`
            QcRequired interface{}  `json:"qc_required"`
            OperationalStatus string  `json:"operational_status"`
            Type string  `json:"type"`
            LineNumber float64  `json:"line_number"`
            BagStatus []BagStatusHistory  `json:"bag_status"`
            AppliedPromos []map[string]interface{}  `json:"applied_promos"`
         
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
            HasNext bool  `json:"has_next"`
            Current float64  `json:"current"`
            ItemTotal float64  `json:"item_total"`
         
    }
    
    // GetBagsPlatformResponse used by Order
    type GetBagsPlatformResponse struct {

        
            Items []BagDetailsPlatformResponse  `json:"items"`
            Page BagsPage  `json:"page"`
         
    }
    
    // GeneratePosOrderReceiptResponse used by Order
    type GeneratePosOrderReceiptResponse struct {

        
            Success bool  `json:"success"`
            InvoiceReceipt string  `json:"invoice_receipt"`
            CustomerCnReceipt string  `json:"customer_cn_receipt"`
            MerchantCnReceipt string  `json:"merchant_cn_receipt"`
            PaymentReceipt string  `json:"payment_receipt"`
            OrderID string  `json:"order_id"`
         
    }
    
    // InvalidateShipmentCachePayload used by Order
    type InvalidateShipmentCachePayload struct {

        
            BagIds []string  `json:"bag_ids"`
            AffiliateBagIds []string  `json:"affiliate_bag_ids"`
            ShipmentIds []string  `json:"shipment_ids"`
         
    }
    
    // InvalidateShipmentCacheNestedResponse used by Order
    type InvalidateShipmentCacheNestedResponse struct {

        
            Status float64  `json:"status"`
            ShipmentID string  `json:"shipment_id"`
            Error string  `json:"error"`
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
            AffiliateID string  `json:"affiliate_id"`
            StoreID float64  `json:"store_id"`
            FyndOrderID string  `json:"fynd_order_id"`
            MongoArticleID string  `json:"mongo_article_id"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
            ReasonIds []float64  `json:"reason_ids"`
            BagID float64  `json:"bag_id"`
            SetID string  `json:"set_id"`
         
    }
    
    // StoreReassignResponse used by Order
    type StoreReassignResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // Entities used by Order
    type Entities struct {

        
            ReasonText string  `json:"reason_text"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            ID string  `json:"id"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            AffiliateID string  `json:"affiliate_id"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
         
    }
    
    // UpdateShipmentLockPayload used by Order
    type UpdateShipmentLockPayload struct {

        
            Action string  `json:"action"`
            ActionType string  `json:"action_type"`
            EntityType string  `json:"entity_type"`
            Entities []Entities  `json:"entities"`
         
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

        
            LockStatus string  `json:"lock_status"`
            OriginalFilter OriginalFilter  `json:"original_filter"`
            Bags []Bags  `json:"bags"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            AffiliateID string  `json:"affiliate_id"`
            Status string  `json:"status"`
            IsShipmentLocked bool  `json:"is_shipment_locked"`
            ShipmentID string  `json:"shipment_id"`
            IsBagLocked bool  `json:"is_bag_locked"`
         
    }
    
    // UpdateShipmentLockResponse used by Order
    type UpdateShipmentLockResponse struct {

        
            CheckResponse []CheckResponse  `json:"check_response"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // AnnouncementResponse used by Order
    type AnnouncementResponse struct {

        
            FromDatetime string  `json:"from_datetime"`
            ToDatetime string  `json:"to_datetime"`
            ID float64  `json:"id"`
            CompanyID float64  `json:"company_id"`
            CreatedAt string  `json:"created_at"`
            Title string  `json:"title"`
            PlatformName string  `json:"platform_name"`
            LogoURL string  `json:"logo_url"`
            Description string  `json:"description"`
            PlatformID string  `json:"platform_id"`
         
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
    
    // Products used by Order
    type Products struct {

        
            Quantity float64  `json:"quantity"`
            Identifier string  `json:"identifier"`
            LineNumber float64  `json:"line_number"`
         
    }
    
    // ProductsReasonsFilters used by Order
    type ProductsReasonsFilters struct {

        
            Quantity float64  `json:"quantity"`
            Identifier string  `json:"identifier"`
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
    
    // OrderItemDataUpdates used by Order
    type OrderItemDataUpdates struct {

        
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
            OrderItemStatus []OrderItemDataUpdates  `json:"order_item_status"`
            Entities []EntitiesDataUpdates  `json:"entities"`
         
    }
    
    // ShipmentsRequest used by Order
    type ShipmentsRequest struct {

        
            Products []Products  `json:"products"`
            Reasons ReasonsData  `json:"reasons"`
            DataUpdates DataUpdates  `json:"data_updates"`
            Identifier string  `json:"identifier"`
         
    }
    
    // StatuesRequest used by Order
    type StatuesRequest struct {

        
            Shipments []ShipmentsRequest  `json:"shipments"`
            ExcludeBagsNextState string  `json:"exclude_bags_next_state"`
            SplitShipment bool  `json:"split_shipment"`
            Status string  `json:"status"`
         
    }
    
    // UpdateShipmentStatusRequest used by Order
    type UpdateShipmentStatusRequest struct {

        
            Task bool  `json:"task"`
            Statuses []StatuesRequest  `json:"statuses"`
            ForceTransition bool  `json:"force_transition"`
            LockAfterTransition bool  `json:"lock_after_transition"`
            UnlockBeforeTransition bool  `json:"unlock_before_transition"`
         
    }
    
    // ShipmentsResponse used by Order
    type ShipmentsResponse struct {

        
            Meta map[string]interface{}  `json:"meta"`
            Exception string  `json:"exception"`
            Code string  `json:"code"`
            Message string  `json:"message"`
            Status float64  `json:"status"`
            StackTrace string  `json:"stack_trace"`
            Identifier string  `json:"identifier"`
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
    
    // AffiliateAppConfigMeta used by Order
    type AffiliateAppConfigMeta struct {

        
            Name string  `json:"name"`
            Value string  `json:"value"`
         
    }
    
    // AffiliateAppConfig used by Order
    type AffiliateAppConfig struct {

        
            Meta []AffiliateAppConfigMeta  `json:"meta"`
            ID string  `json:"id"`
            CreatedAt string  `json:"created_at"`
            UpdatedAt string  `json:"updated_at"`
            Owner string  `json:"owner"`
            Secret string  `json:"secret"`
            Name string  `json:"name"`
            Description string  `json:"description"`
            Token string  `json:"token"`
         
    }
    
    // AffiliateInventoryOrderConfig used by Order
    type AffiliateInventoryOrderConfig struct {

        
            ForceReassignment bool  `json:"force_reassignment"`
         
    }
    
    // AffiliateInventoryLogisticsConfig used by Order
    type AffiliateInventoryLogisticsConfig struct {

        
            DpAssignment bool  `json:"dp_assignment"`
         
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
    
    // AffiliateInventoryArticleAssignmentConfig used by Order
    type AffiliateInventoryArticleAssignmentConfig struct {

        
            PostOrderReassignment bool  `json:"post_order_reassignment"`
         
    }
    
    // AffiliateInventoryConfig used by Order
    type AffiliateInventoryConfig struct {

        
            Order AffiliateInventoryOrderConfig  `json:"order"`
            Logistics AffiliateInventoryLogisticsConfig  `json:"logistics"`
            Inventory AffiliateInventoryStoreConfig  `json:"inventory"`
            Payment AffiliateInventoryPaymentConfig  `json:"payment"`
            ArticleAssignment AffiliateInventoryArticleAssignmentConfig  `json:"article_assignment"`
         
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

        
            StoreID float64  `json:"store_id"`
            MarketplaceStoreID string  `json:"marketplace_store_id"`
         
    }
    
    // OrderConfig used by Order
    type OrderConfig struct {

        
            Affiliate Affiliate  `json:"affiliate"`
            BagEndState string  `json:"bag_end_state"`
            AffiliateStoreIDMapping []AffiliateStoreIdMapping  `json:"affiliate_store_id_mapping"`
            CreateUser bool  `json:"create_user"`
            StoreLookup string  `json:"store_lookup"`
            ArticleLookup string  `json:"article_lookup"`
         
    }
    
    // OrderUser used by Order
    type OrderUser struct {

        
            Address1 string  `json:"address1"`
            LastName string  `json:"last_name"`
            Mobile float64  `json:"mobile"`
            Address2 string  `json:"address2"`
            Phone float64  `json:"phone"`
            Country string  `json:"country"`
            FirstName string  `json:"first_name"`
            Pincode string  `json:"pincode"`
            State string  `json:"state"`
            Email string  `json:"email"`
            City string  `json:"city"`
         
    }
    
    // UserData used by Order
    type UserData struct {

        
            BillingUser OrderUser  `json:"billing_user"`
            ShippingUser OrderUser  `json:"shipping_user"`
         
    }
    
    // MarketPlacePdf used by Order
    type MarketPlacePdf struct {

        
            Invoice string  `json:"invoice"`
            Label string  `json:"label"`
         
    }
    
    // AffiliateBag used by Order
    type AffiliateBag struct {

        
            Sku string  `json:"sku"`
            AffiliateMeta map[string]interface{}  `json:"affiliate_meta"`
            AvlQty float64  `json:"avl_qty"`
            TransferPrice float64  `json:"transfer_price"`
            SellerIdentifier string  `json:"seller_identifier"`
            UnitPrice float64  `json:"unit_price"`
            FyndStoreID string  `json:"fynd_store_id"`
            PdfLinks MarketPlacePdf  `json:"pdf_links"`
            ItemID float64  `json:"item_id"`
            CompanyID float64  `json:"company_id"`
            AmountPaid float64  `json:"amount_paid"`
            ID string  `json:"_id"`
            StoreID float64  `json:"store_id"`
            Identifier map[string]interface{}  `json:"identifier"`
            Discount float64  `json:"discount"`
            ModifiedOn string  `json:"modified_on"`
            ItemSize string  `json:"item_size"`
            HsnCodeID string  `json:"hsn_code_id"`
            AffiliateStoreID string  `json:"affiliate_store_id"`
            Quantity float64  `json:"quantity"`
            DeliveryCharge float64  `json:"delivery_charge"`
            PriceEffective float64  `json:"price_effective"`
            PriceMarked float64  `json:"price_marked"`
         
    }
    
    // OrderPriority used by Order
    type OrderPriority struct {

        
            AffiliatePriorityCode string  `json:"affiliate_priority_code"`
            FulfilmentPriorityText string  `json:"fulfilment_priority_text"`
            FulfilmentPriority float64  `json:"fulfilment_priority"`
         
    }
    
    // ArticleDetails1 used by Order
    type ArticleDetails1 struct {

        
            BrandID float64  `json:"brand_id"`
            Quantity float64  `json:"quantity"`
            Category map[string]interface{}  `json:"category"`
            ID string  `json:"_id"`
            Dimension map[string]interface{}  `json:"dimension"`
            Attributes map[string]interface{}  `json:"attributes"`
            Weight map[string]interface{}  `json:"weight"`
         
    }
    
    // LocationDetails used by Order
    type LocationDetails struct {

        
            Articles []ArticleDetails1  `json:"articles"`
            FulfillmentType string  `json:"fulfillment_type"`
            FulfillmentID float64  `json:"fulfillment_id"`
         
    }
    
    // ShipmentDetails1 used by Order
    type ShipmentDetails1 struct {

        
            Shipments float64  `json:"shipments"`
            Meta map[string]interface{}  `json:"meta"`
            DpID float64  `json:"dp_id"`
            FulfillmentID float64  `json:"fulfillment_id"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            Articles []ArticleDetails1  `json:"articles"`
            BoxType string  `json:"box_type"`
         
    }
    
    // ShipmentConfig used by Order
    type ShipmentConfig struct {

        
            Journey string  `json:"journey"`
            PaymentMode string  `json:"payment_mode"`
            LocationDetails LocationDetails  `json:"location_details"`
            Action string  `json:"action"`
            ToPincode string  `json:"to_pincode"`
            Identifier string  `json:"identifier"`
            Source string  `json:"source"`
            Shipment []ShipmentDetails1  `json:"shipment"`
         
    }
    
    // ShipmentData used by Order
    type ShipmentData struct {

        
            ShipmentData ShipmentConfig  `json:"shipment_data"`
         
    }
    
    // OrderInfo used by Order
    type OrderInfo struct {

        
            Items map[string]interface{}  `json:"items"`
            PaymentMode string  `json:"payment_mode"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            OrderValue float64  `json:"order_value"`
            DeliveryCharges float64  `json:"delivery_charges"`
            ShippingAddress OrderUser  `json:"shipping_address"`
            User UserData  `json:"user"`
            Coupon string  `json:"coupon"`
            Bags []AffiliateBag  `json:"bags"`
            BillingAddress OrderUser  `json:"billing_address"`
            Payment map[string]interface{}  `json:"payment"`
            OrderPriority OrderPriority  `json:"order_priority"`
            Discount float64  `json:"discount"`
            CodCharges float64  `json:"cod_charges"`
            Shipment ShipmentData  `json:"shipment"`
         
    }
    
    // CreateOrderPayload used by Order
    type CreateOrderPayload struct {

        
            AffiliateID string  `json:"affiliate_id"`
            OrderConfig OrderConfig  `json:"order_config"`
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

        
            ID float64  `json:"id"`
            Description string  `json:"description"`
            Slug string  `json:"slug"`
            DisplayText string  `json:"display_text"`
         
    }
    
    // GetActionsResponse used by Order
    type GetActionsResponse struct {

        
            Permissions []ActionInfo  `json:"permissions"`
         
    }
    
    // PostHistoryFilters used by Order
    type PostHistoryFilters struct {

        
            ShipmentID string  `json:"shipment_id"`
            Identifier string  `json:"identifier"`
            LineNumber string  `json:"line_number"`
         
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
    
    // HistoryReason used by Order
    type HistoryReason struct {

        
            DislayName string  `json:"dislay_name"`
            Category string  `json:"category"`
            Quantity float64  `json:"quantity"`
            Code float64  `json:"code"`
            Text string  `json:"text"`
            State string  `json:"state"`
         
    }
    
    // HistoryMeta used by Order
    type HistoryMeta struct {

        
            Recordpath string  `json:"recordpath"`
            Billsec string  `json:"billsec"`
            Receiver string  `json:"receiver"`
            Recipient string  `json:"recipient"`
            Status string  `json:"status"`
            ActivityType string  `json:"activity_type"`
            ChannelType string  `json:"channel_type"`
            Callerid string  `json:"callerid"`
            Endtime string  `json:"endtime"`
            ShortLink string  `json:"short_link"`
            ActivityComment string  `json:"activity_comment"`
            Status2 string  `json:"status2"`
            Slug string  `json:"slug"`
            Message string  `json:"message"`
            StoreCode string  `json:"store_code"`
            StoreID float64  `json:"store_id"`
            Status1 string  `json:"status1"`
            StoreName string  `json:"store_name"`
            Starttime string  `json:"starttime"`
            Duration string  `json:"duration"`
            CallID string  `json:"call_id"`
            Caller string  `json:"caller"`
            Reason HistoryReason  `json:"reason"`
         
    }
    
    // HistoryDict used by Order
    type HistoryDict struct {

        
            Createdat string  `json:"createdat"`
            L1Detail string  `json:"l1_detail"`
            Meta HistoryMeta  `json:"meta"`
            AssignedAgent string  `json:"assigned_agent"`
            User string  `json:"user"`
            L3Detail string  `json:"l3_detail"`
            L2Detail string  `json:"l2_detail"`
            Message string  `json:"message"`
            BagID float64  `json:"bag_id"`
            DisplayMessage string  `json:"display_message"`
            Type string  `json:"type"`
            TicketURL string  `json:"ticket_url"`
            TicketID string  `json:"ticket_id"`
         
    }
    
    // ShipmentHistoryResponse used by Order
    type ShipmentHistoryResponse struct {

        
            Success bool  `json:"success"`
            ActivityHistory []HistoryDict  `json:"activity_history"`
         
    }
    
    // SmsDataPayload used by Order
    type SmsDataPayload struct {

        
            PaymentMode string  `json:"payment_mode"`
            BrandName string  `json:"brand_name"`
            AmountPaid float64  `json:"amount_paid"`
            Message string  `json:"message"`
            CustomerName string  `json:"customer_name"`
            OrderID string  `json:"order_id"`
            CountryCode string  `json:"country_code"`
            PhoneNumber float64  `json:"phone_number"`
            ShipmentID float64  `json:"shipment_id"`
         
    }
    
    // SendSmsPayload used by Order
    type SendSmsPayload struct {

        
            Data SmsDataPayload  `json:"data"`
            BagID float64  `json:"bag_id"`
            Slug string  `json:"slug"`
         
    }
    
    // Meta used by Order
    type Meta struct {

        
            KafkaEmissionStatus float64  `json:"kafka_emission_status"`
            StateManagerUsed string  `json:"state_manager_used"`
         
    }
    
    // ShipmentDetail used by Order
    type ShipmentDetail struct {

        
            Meta Meta  `json:"meta"`
            ID float64  `json:"id"`
            Status string  `json:"status"`
            BagList []float64  `json:"bag_list"`
            Remarks string  `json:"remarks"`
            ShipmentID string  `json:"shipment_id"`
         
    }
    
    // OrderDetails used by Order
    type OrderDetails struct {

        
            FyndOrderID string  `json:"fynd_order_id"`
            CreatedAt string  `json:"created_at"`
         
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

        
            Length float64  `json:"length"`
            Width float64  `json:"width"`
            PackagingType string  `json:"packaging_type"`
            Weight string  `json:"weight"`
            Height string  `json:"height"`
         
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
    
    // ProcessingDates used by Order
    type ProcessingDates struct {

        
            ConfirmByDate string  `json:"confirm_by_date"`
            DpPickupSlot map[string]interface{}  `json:"dp_pickup_slot"`
            CustomerPickupSlot map[string]interface{}  `json:"customer_pickup_slot"`
            PackByDate string  `json:"pack_by_date"`
            DispatchByDate string  `json:"dispatch_by_date"`
            DispatchAfterDate string  `json:"dispatch_after_date"`
         
    }
    
    // Tax used by Order
    type Tax struct {

        
            Name string  `json:"name"`
            Rate float64  `json:"rate"`
            Amount map[string]interface{}  `json:"amount"`
            Breakup []map[string]interface{}  `json:"breakup"`
         
    }
    
    // Charge used by Order
    type Charge struct {

        
            Tax Tax  `json:"tax"`
            Code string  `json:"code"`
            Amount map[string]interface{}  `json:"amount"`
            Type string  `json:"type"`
            Name string  `json:"name"`
         
    }
    
    // LineItem used by Order
    type LineItem struct {

        
            ExternalLineID string  `json:"external_line_id"`
            Charges []Charge  `json:"charges"`
            Quantity float64  `json:"quantity"`
            Meta map[string]interface{}  `json:"meta"`
            SellerIdentifier string  `json:"seller_identifier"`
            CustomMessasge string  `json:"custom_messasge"`
         
    }
    
    // Shipment used by Order
    type Shipment struct {

        
            ExternalShipmentID string  `json:"external_shipment_id"`
            Meta map[string]interface{}  `json:"meta"`
            ProcessingDates ProcessingDates  `json:"processing_dates"`
            Priority float64  `json:"priority"`
            LocationID float64  `json:"location_id"`
            LineItems []LineItem  `json:"line_items"`
         
    }
    
    // ShippingInfo used by Order
    type ShippingInfo struct {

        
            FloorNo string  `json:"floor_no"`
            ExternalCustomerCode string  `json:"external_customer_code"`
            AddressType string  `json:"address_type"`
            CustomerCode string  `json:"customer_code"`
            Country string  `json:"country"`
            Gender string  `json:"gender"`
            CountryCode string  `json:"country_code"`
            Landmark string  `json:"landmark"`
            Address2 string  `json:"address2"`
            StateCode string  `json:"state_code"`
            ShippingType string  `json:"shipping_type"`
            State string  `json:"state"`
            City string  `json:"city"`
            Address1 string  `json:"address1"`
            PrimaryEmail string  `json:"primary_email"`
            AlternateEmail string  `json:"alternate_email"`
            FirstName string  `json:"first_name"`
            MiddleName string  `json:"middle_name"`
            GeoLocation map[string]interface{}  `json:"geo_location"`
            Pincode string  `json:"pincode"`
            HouseNo string  `json:"house_no"`
            Slot []map[string]interface{}  `json:"slot"`
            AlternateMobileNumber string  `json:"alternate_mobile_number"`
            LastName string  `json:"last_name"`
            Title string  `json:"title"`
            PrimaryMobileNumber string  `json:"primary_mobile_number"`
         
    }
    
    // TaxInfo used by Order
    type TaxInfo struct {

        
            B2bGstinNumber string  `json:"b2b_gstin_number"`
            Gstin string  `json:"gstin"`
         
    }
    
    // PaymentMethod used by Order
    type PaymentMethod struct {

        
            Meta map[string]interface{}  `json:"meta"`
            CollectBy string  `json:"collect_by"`
            Amount float64  `json:"amount"`
            TransactionData map[string]interface{}  `json:"transaction_data"`
            Mode string  `json:"mode"`
            RefundBy string  `json:"refund_by"`
            Name string  `json:"name"`
         
    }
    
    // PaymentInfo used by Order
    type PaymentInfo struct {

        
            PaymentMethods []PaymentMethod  `json:"payment_methods"`
            PrimaryMode string  `json:"primary_mode"`
         
    }
    
    // BillingInfo used by Order
    type BillingInfo struct {

        
            FloorNo string  `json:"floor_no"`
            ExternalCustomerCode string  `json:"external_customer_code"`
            CustomerCode string  `json:"customer_code"`
            Country string  `json:"country"`
            Gender string  `json:"gender"`
            CountryCode string  `json:"country_code"`
            Address2 string  `json:"address2"`
            StateCode string  `json:"state_code"`
            State string  `json:"state"`
            City string  `json:"city"`
            Address1 string  `json:"address1"`
            PrimaryEmail string  `json:"primary_email"`
            AlternateEmail string  `json:"alternate_email"`
            FirstName string  `json:"first_name"`
            MiddleName string  `json:"middle_name"`
            Pincode string  `json:"pincode"`
            HouseNo string  `json:"house_no"`
            AlternateMobileNumber string  `json:"alternate_mobile_number"`
            LastName string  `json:"last_name"`
            Title string  `json:"title"`
            PrimaryMobileNumber string  `json:"primary_mobile_number"`
         
    }
    
    // CreateOrderAPI used by Order
    type CreateOrderAPI struct {

        
            Shipments []Shipment  `json:"shipments"`
            ExternalOrderID string  `json:"external_order_id"`
            ShippingInfo ShippingInfo  `json:"shipping_info"`
            TaxInfo TaxInfo  `json:"tax_info"`
            Charges []Charge  `json:"charges"`
            PaymentInfo PaymentInfo  `json:"payment_info"`
            Config map[string]interface{}  `json:"config"`
            Meta map[string]interface{}  `json:"meta"`
            CurrencyInfo map[string]interface{}  `json:"currency_info"`
            BillingInfo BillingInfo  `json:"billing_info"`
            ExternalCreationDate string  `json:"external_creation_date"`
         
    }
    
    // CreateOrderErrorReponse used by Order
    type CreateOrderErrorReponse struct {

        
            Meta string  `json:"meta"`
            Exception string  `json:"exception"`
            Code string  `json:"code"`
            Message string  `json:"message"`
            Status float64  `json:"status"`
            Info interface{}  `json:"info"`
            RequestID string  `json:"request_id"`
            StackTrace string  `json:"stack_trace"`
         
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

        
            PaymentMethods []PaymentMethods  `json:"payment_methods"`
            Source string  `json:"source"`
            ModeOfPayment string  `json:"mode_of_payment"`
         
    }
    
    // CreateChannelConfig used by Order
    type CreateChannelConfig struct {

        
            DpConfiguration DpConfiguration  `json:"dp_configuration"`
            LockStates []string  `json:"lock_states"`
            PaymentInfo CreateChannelPaymentInfo  `json:"payment_info"`
            ShipmentAssignment string  `json:"shipment_assignment"`
            LocationReassignment bool  `json:"location_reassignment"`
            LogoURL map[string]interface{}  `json:"logo_url"`
         
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

        
            Mobile float64  `json:"mobile"`
            StartDate string  `json:"start_date"`
            EndDate string  `json:"end_date"`
            OrderDetails []FyndOrderIdList  `json:"order_details"`
         
    }
    
    // BagStateTransitionMap used by Order
    type BagStateTransitionMap struct {

        
            Fynd map[string]interface{}  `json:"fynd"`
            Affiliate map[string]interface{}  `json:"affiliate"`
         
    }
    
    // FetchCreditBalanceRequestPayload used by Order
    type FetchCreditBalanceRequestPayload struct {

        
            AffiliateID string  `json:"affiliate_id"`
            SellerID string  `json:"seller_id"`
            CustomerMobileNumber string  `json:"customer_mobile_number"`
         
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

        
            SellerID float64  `json:"seller_id"`
            FyndOrderID string  `json:"fynd_order_id"`
            CustomerMobileNumber string  `json:"customer_mobile_number"`
            AffiliateID string  `json:"affiliate_id"`
         
    }
    
    // RefundOption used by Order
    type RefundOption struct {

        
            Type string  `json:"type"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Value string  `json:"value"`
         
    }
    
    // RefundModeInfo used by Order
    type RefundModeInfo struct {

        
            DisplayName string  `json:"display_name"`
            Options []RefundOption  `json:"options"`
            IsActive bool  `json:"is_active"`
            Slug string  `json:"slug"`
         
    }
    
    // RefundModeConfigResponsePayload used by Order
    type RefundModeConfigResponsePayload struct {

        
            Success bool  `json:"success"`
            Data []RefundModeInfo  `json:"data"`
         
    }
    
    // AttachUserInfo used by Order
    type AttachUserInfo struct {

        
            LastName string  `json:"last_name"`
            Mobile string  `json:"mobile"`
            FirstName string  `json:"first_name"`
            CountryCode string  `json:"country_code"`
         
    }
    
    // AttachUserOtpData used by Order
    type AttachUserOtpData struct {

        
            RequestID string  `json:"request_id"`
         
    }
    
    // AttachOrderUser used by Order
    type AttachOrderUser struct {

        
            UserInfo AttachUserInfo  `json:"user_info"`
            FyndOrderID string  `json:"fynd_order_id"`
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

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
            Data PointBlankOtpData  `json:"data"`
         
    }
    
    // VerifyOtpData used by Order
    type VerifyOtpData struct {

        
            RequestID string  `json:"request_id"`
            Mobile string  `json:"mobile"`
            OtpCode float64  `json:"otp_code"`
         
    }
    
    // VerifyMobileOTP used by Order
    type VerifyMobileOTP struct {

        
            FyndOrderID string  `json:"fynd_order_id"`
            OtpData VerifyOtpData  `json:"otp_data"`
         
    }
    
    // VerifyOtpResponseData used by Order
    type VerifyOtpResponseData struct {

        
            CountryCode string  `json:"country_code"`
            FyndOrderID string  `json:"fynd_order_id"`
            Message string  `json:"message"`
            Mobile string  `json:"mobile"`
         
    }
    
    // VerifyOtpResponse used by Order
    type VerifyOtpResponse struct {

        
            Data VerifyOtpResponseData  `json:"data"`
            Status float64  `json:"status"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    

    
    // GetSearchWordsData used by Catalog
    type GetSearchWordsData struct {

        
            UID string  `json:"uid"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Words []string  `json:"words"`
            IsActive bool  `json:"is_active"`
            Result map[string]interface{}  `json:"result"`
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

        
            Query map[string]interface{}  `json:"query"`
            SortOn string  `json:"sort_on"`
         
    }
    
    // CreateSearchKeyword used by Catalog
    type CreateSearchKeyword struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Words []string  `json:"words"`
            IsActive bool  `json:"is_active"`
            Result SearchKeywordResult  `json:"result"`
            AppID string  `json:"app_id"`
         
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
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Words []string  `json:"words"`
            Results []map[string]interface{}  `json:"results"`
            AppID string  `json:"app_id"`
         
    }
    
    // GetAutocompleteWordsResponse used by Catalog
    type GetAutocompleteWordsResponse struct {

        
            Items []GetAutocompleteWordsData  `json:"items"`
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
    
    // AutoCompleteMedia used by Catalog
    type AutoCompleteMedia struct {

        
            AspectRatio string  `json:"aspect_ratio"`
            Type string  `json:"type"`
            URL string  `json:"url"`
         
    }
    
    // AutocompleteResult used by Catalog
    type AutocompleteResult struct {

        
            Action AutocompleteAction  `json:"action"`
            Logo AutoCompleteMedia  `json:"logo"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Display string  `json:"display"`
         
    }
    
    // CreateAutocompleteKeyword used by Catalog
    type CreateAutocompleteKeyword struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Words []string  `json:"words"`
            IsActive bool  `json:"is_active"`
            Results []AutocompleteResult  `json:"results"`
            AppID string  `json:"app_id"`
         
    }
    
    // CreateAutocompleteWordsResponse used by Catalog
    type CreateAutocompleteWordsResponse struct {

        
            Results []map[string]interface{}  `json:"results"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            AppID string  `json:"app_id"`
            Words []string  `json:"words"`
         
    }
    
    // ProductBundleItem used by Catalog
    type ProductBundleItem struct {

        
            ProductUID float64  `json:"product_uid"`
            AutoAddToCart bool  `json:"auto_add_to_cart"`
            MinQuantity float64  `json:"min_quantity"`
            AutoSelect bool  `json:"auto_select"`
            MaxQuantity float64  `json:"max_quantity"`
            AllowRemove bool  `json:"allow_remove"`
         
    }
    
    // GetProductBundleCreateResponse used by Catalog
    type GetProductBundleCreateResponse struct {

        
            PageVisibility []string  `json:"page_visibility"`
            ID string  `json:"id"`
            CompanyID float64  `json:"company_id"`
            Products []ProductBundleItem  `json:"products"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Choice string  `json:"choice"`
            Slug string  `json:"slug"`
            CreatedOn string  `json:"created_on"`
            IsActive bool  `json:"is_active"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            ModifiedOn string  `json:"modified_on"`
            Meta map[string]interface{}  `json:"meta"`
            Name string  `json:"name"`
            Logo string  `json:"logo"`
         
    }
    
    // GetProductBundleListingResponse used by Catalog
    type GetProductBundleListingResponse struct {

        
            Items []GetProductBundleCreateResponse  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // ProductBundleRequest used by Catalog
    type ProductBundleRequest struct {

        
            PageVisibility []string  `json:"page_visibility"`
            CompanyID float64  `json:"company_id"`
            Products []ProductBundleItem  `json:"products"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Choice string  `json:"choice"`
            Slug string  `json:"slug"`
            CreatedOn string  `json:"created_on"`
            IsActive bool  `json:"is_active"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            ModifiedOn string  `json:"modified_on"`
            Meta map[string]interface{}  `json:"meta"`
            Name string  `json:"name"`
            Logo string  `json:"logo"`
         
    }
    
    // LimitedProductData used by Catalog
    type LimitedProductData struct {

        
            Sizes []string  `json:"sizes"`
            UID float64  `json:"uid"`
            Identifier map[string]interface{}  `json:"identifier"`
            Slug string  `json:"slug"`
            CountryOfOrigin string  `json:"country_of_origin"`
            Quantity float64  `json:"quantity"`
            Images []string  `json:"images"`
            Name string  `json:"name"`
            Attributes map[string]interface{}  `json:"attributes"`
            ItemCode string  `json:"item_code"`
            ShortDescription string  `json:"short_description"`
            Price map[string]interface{}  `json:"price"`
         
    }
    
    // Size used by Catalog
    type Size struct {

        
            IsAvailable bool  `json:"is_available"`
            Value string  `json:"value"`
            Display string  `json:"display"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // Price used by Catalog
    type Price struct {

        
            Currency string  `json:"currency"`
            MaxEffective float64  `json:"max_effective"`
            MinEffective float64  `json:"min_effective"`
            MinMarked float64  `json:"min_marked"`
            MaxMarked float64  `json:"max_marked"`
         
    }
    
    // GetProducts used by Catalog
    type GetProducts struct {

        
            ProductUID float64  `json:"product_uid"`
            ProductDetails LimitedProductData  `json:"product_details"`
            AutoAddToCart bool  `json:"auto_add_to_cart"`
            MinQuantity float64  `json:"min_quantity"`
            AutoSelect bool  `json:"auto_select"`
            MaxQuantity float64  `json:"max_quantity"`
            Sizes []Size  `json:"sizes"`
            AllowRemove bool  `json:"allow_remove"`
            Price Price  `json:"price"`
         
    }
    
    // GetProductBundleResponse used by Catalog
    type GetProductBundleResponse struct {

        
            PageVisibility []string  `json:"page_visibility"`
            CompanyID float64  `json:"company_id"`
            Products []GetProducts  `json:"products"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            Choice string  `json:"choice"`
            Slug string  `json:"slug"`
            IsActive bool  `json:"is_active"`
            Meta map[string]interface{}  `json:"meta"`
            Name string  `json:"name"`
            Logo string  `json:"logo"`
         
    }
    
    // ProductBundleUpdateRequest used by Catalog
    type ProductBundleUpdateRequest struct {

        
            PageVisibility []string  `json:"page_visibility"`
            CompanyID float64  `json:"company_id"`
            Products []ProductBundleItem  `json:"products"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            Choice string  `json:"choice"`
            Slug string  `json:"slug"`
            IsActive bool  `json:"is_active"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            ModifiedOn string  `json:"modified_on"`
            Meta map[string]interface{}  `json:"meta"`
            Name string  `json:"name"`
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
            Headers map[string]interface{}  `json:"headers"`
            Values []map[string]interface{}  `json:"values"`
         
    }
    
    // Guide used by Catalog
    type Guide struct {

        
            Meta Meta  `json:"meta"`
         
    }
    
    // ValidateSizeGuide used by Catalog
    type ValidateSizeGuide struct {

        
            ID string  `json:"id"`
            CompanyID float64  `json:"company_id"`
            Image string  `json:"image"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            CreatedOn string  `json:"created_on"`
            Title string  `json:"title"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Active bool  `json:"active"`
            Tag string  `json:"tag"`
            BrandID float64  `json:"brand_id"`
            ModifiedOn string  `json:"modified_on"`
            Name string  `json:"name"`
            Description string  `json:"description"`
            Guide Guide  `json:"guide"`
            Subtitle string  `json:"subtitle"`
         
    }
    
    // SuccessResponse used by Catalog
    type SuccessResponse struct {

        
            Success bool  `json:"success"`
         
    }
    
    // SizeGuideResponse used by Catalog
    type SizeGuideResponse struct {

        
            ID string  `json:"id"`
            CompanyID float64  `json:"company_id"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Subtitle string  `json:"subtitle"`
            CreatedOn string  `json:"created_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            BrandID float64  `json:"brand_id"`
            ModifiedOn string  `json:"modified_on"`
            Tag string  `json:"tag"`
            Name string  `json:"name"`
            Active bool  `json:"active"`
            Guide map[string]interface{}  `json:"guide"`
            Title string  `json:"title"`
         
    }
    
    // MOQData used by Catalog
    type MOQData struct {

        
            IncrementUnit float64  `json:"increment_unit"`
            Maximum float64  `json:"maximum"`
            Minimum float64  `json:"minimum"`
         
    }
    
    // SEOData used by Catalog
    type SEOData struct {

        
            Description interface{}  `json:"description"`
            Title interface{}  `json:"title"`
         
    }
    
    // OwnerAppItemResponse used by Catalog
    type OwnerAppItemResponse struct {

        
            IsCod bool  `json:"is_cod"`
            Moq MOQData  `json:"moq"`
            IsGift bool  `json:"is_gift"`
            AltText map[string]interface{}  `json:"alt_text"`
            Seo SEOData  `json:"seo"`
         
    }
    
    // MetaFields used by Catalog
    type MetaFields struct {

        
            Key interface{}  `json:"key"`
            Value interface{}  `json:"value"`
         
    }
    
    // ApplicationItemMOQ used by Catalog
    type ApplicationItemMOQ struct {

        
            IncrementUnit float64  `json:"increment_unit"`
            Maximum float64  `json:"maximum"`
            Minimum float64  `json:"minimum"`
         
    }
    
    // ApplicationItemSEO used by Catalog
    type ApplicationItemSEO struct {

        
            Description interface{}  `json:"description"`
            Title interface{}  `json:"title"`
         
    }
    
    // ApplicationItemMeta used by Catalog
    type ApplicationItemMeta struct {

        
            CustomMeta []MetaFields  `json:"_custom_meta"`
            IsCod bool  `json:"is_cod"`
            Moq ApplicationItemMOQ  `json:"moq"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            IsGift bool  `json:"is_gift"`
            AltText map[string]interface{}  `json:"alt_text"`
            Seo ApplicationItemSEO  `json:"seo"`
         
    }
    
    // SuccessResponse1 used by Catalog
    type SuccessResponse1 struct {

        
            UID float64  `json:"uid"`
            Success bool  `json:"success"`
         
    }
    
    // GetConfigMetadataResponse used by Catalog
    type GetConfigMetadataResponse struct {

        
            Values []map[string]interface{}  `json:"values"`
            Condition []map[string]interface{}  `json:"condition"`
            Data []map[string]interface{}  `json:"data"`
         
    }
    
    // PageResponseType used by Catalog
    type PageResponseType struct {

        
            Next float64  `json:"next"`
            Current float64  `json:"current"`
            TotalCount float64  `json:"total_count"`
            HasNext bool  `json:"has_next"`
         
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

        
            Key string  `json:"key"`
            Slug string  `json:"slug"`
            IsActive bool  `json:"is_active"`
            DisplayType string  `json:"display_type"`
            Unit string  `json:"unit"`
            Name string  `json:"name"`
            Logo string  `json:"logo"`
            Priority float64  `json:"priority"`
         
    }
    
    // AppConfigurationDetail used by Catalog
    type AppConfigurationDetail struct {

        
            Attributes []AttributeDetailsGroup  `json:"attributes"`
            IsDefault bool  `json:"is_default"`
            Slug string  `json:"slug"`
            Priority float64  `json:"priority"`
            IsActive bool  `json:"is_active"`
            TemplateSlugs []string  `json:"template_slugs"`
            Name string  `json:"name"`
            Logo string  `json:"logo"`
            AppID string  `json:"app_id"`
         
    }
    
    // ConfigSuccessResponse used by Catalog
    type ConfigSuccessResponse struct {

        
            Message string  `json:"message"`
         
    }
    
    // AppConfigurationsSort used by Catalog
    type AppConfigurationsSort struct {

        
            Key string  `json:"key"`
            IsDefault bool  `json:"is_default"`
            DefaultKey string  `json:"default_key"`
            Priority float64  `json:"priority"`
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            Logo string  `json:"logo"`
            AppID string  `json:"app_id"`
         
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

        
            Key string  `json:"key"`
            Units []map[string]interface{}  `json:"units"`
            FilterTypes []string  `json:"filter_types"`
            Display string  `json:"display"`
         
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

        
            Similar map[string]interface{}  `json:"similar"`
            Detail map[string]interface{}  `json:"detail"`
            Compare map[string]interface{}  `json:"compare"`
            Variant map[string]interface{}  `json:"variant"`
         
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

        
            Key string  `json:"key"`
            Size ProductSize  `json:"size"`
            Subtitle string  `json:"subtitle"`
            IsActive bool  `json:"is_active"`
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

        
            Key string  `json:"key"`
            Size ProductSize  `json:"size"`
            IsActive bool  `json:"is_active"`
            DisplayType string  `json:"display_type"`
            Name string  `json:"name"`
            Logo string  `json:"logo"`
            Priority float64  `json:"priority"`
         
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

        
            End float64  `json:"end"`
            Display string  `json:"display"`
            Start float64  `json:"start"`
         
    }
    
    // ConfigurationListingFilterValue used by Catalog
    type ConfigurationListingFilterValue struct {

        
            Condition string  `json:"condition"`
            BucketPoints []ConfigurationBucketPoints  `json:"bucket_points"`
            MapValues []map[string]interface{}  `json:"map_values"`
            Sort string  `json:"sort"`
            Map map[string]interface{}  `json:"map"`
            Priority []string  `json:"priority"`
            Value string  `json:"value"`
         
    }
    
    // ConfigurationListingFilterConfig used by Catalog
    type ConfigurationListingFilterConfig struct {

        
            Key string  `json:"key"`
            ValueConfig ConfigurationListingFilterValue  `json:"value_config"`
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            Logo string  `json:"logo"`
            DisplayName string  `json:"display_name"`
            Priority float64  `json:"priority"`
            Type string  `json:"type"`
         
    }
    
    // ConfigurationListingFilter used by Catalog
    type ConfigurationListingFilter struct {

        
            AttributeConfig []ConfigurationListingFilterConfig  `json:"attribute_config"`
            AllowSingle bool  `json:"allow_single"`
         
    }
    
    // ConfigurationListingSortConfig used by Catalog
    type ConfigurationListingSortConfig struct {

        
            Key string  `json:"key"`
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            Logo string  `json:"logo"`
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
    
    // AppCatalogConfiguration used by Catalog
    type AppCatalogConfiguration struct {

        
            ID string  `json:"id"`
            Product ConfigurationProduct  `json:"product"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ConfigType string  `json:"config_type"`
            Listing ConfigurationListing  `json:"listing"`
            CreatedOn string  `json:"created_on"`
            ConfigID string  `json:"config_id"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            ModifiedOn string  `json:"modified_on"`
            AppID string  `json:"app_id"`
            Type string  `json:"type"`
         
    }
    
    // GetAppCatalogConfiguration used by Catalog
    type GetAppCatalogConfiguration struct {

        
            IsDefault bool  `json:"is_default"`
            Data AppCatalogConfiguration  `json:"data"`
         
    }
    
    // AppConfiguration used by Catalog
    type AppConfiguration struct {

        
            Product ConfigurationProduct  `json:"product"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ConfigType string  `json:"config_type"`
            Listing ConfigurationListing  `json:"listing"`
            CreatedOn string  `json:"created_on"`
            ConfigID string  `json:"config_id"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            ModifiedOn string  `json:"modified_on"`
            AppID string  `json:"app_id"`
            Type string  `json:"type"`
         
    }
    
    // GetCatalogConfigurationDetailsSchemaListing used by Catalog
    type GetCatalogConfigurationDetailsSchemaListing struct {

        
            Filter map[string]interface{}  `json:"filter"`
            Sort map[string]interface{}  `json:"sort"`
         
    }
    
    // EntityConfiguration used by Catalog
    type EntityConfiguration struct {

        
            ID string  `json:"id"`
            Product GetCatalogConfigurationDetailsProduct  `json:"product"`
            ConfigType string  `json:"config_type"`
            Listing GetCatalogConfigurationDetailsSchemaListing  `json:"listing"`
            ConfigID string  `json:"config_id"`
            AppID string  `json:"app_id"`
         
    }
    
    // GetAppCatalogEntityConfiguration used by Catalog
    type GetAppCatalogEntityConfiguration struct {

        
            IsDefault bool  `json:"is_default"`
            Data EntityConfiguration  `json:"data"`
         
    }
    
    // ProductSortOn used by Catalog
    type ProductSortOn struct {

        
            Name string  `json:"name"`
            IsSelected bool  `json:"is_selected"`
            Value string  `json:"value"`
         
    }
    
    // ProductFiltersKey used by Catalog
    type ProductFiltersKey struct {

        
            Display string  `json:"display"`
            Operators []string  `json:"operators"`
            Kind string  `json:"kind"`
            Name string  `json:"name"`
            Logo string  `json:"logo"`
         
    }
    
    // ProductFiltersValue used by Catalog
    type ProductFiltersValue struct {

        
            Max float64  `json:"max"`
            SelectedMin float64  `json:"selected_min"`
            Min float64  `json:"min"`
            IsSelected bool  `json:"is_selected"`
            Count float64  `json:"count"`
            Display string  `json:"display"`
            DisplayFormat string  `json:"display_format"`
            CurrencySymbol string  `json:"currency_symbol"`
            QueryFormat string  `json:"query_format"`
            SelectedMax float64  `json:"selected_max"`
            CurrencyCode string  `json:"currency_code"`
            Value interface{}  `json:"value"`
         
    }
    
    // ProductFilters used by Catalog
    type ProductFilters struct {

        
            Key ProductFiltersKey  `json:"key"`
            Values []ProductFiltersValue  `json:"values"`
         
    }
    
    // GetCollectionQueryOptionResponse used by Catalog
    type GetCollectionQueryOptionResponse struct {

        
            SortOn []ProductSortOn  `json:"sort_on"`
            Filters []ProductFilters  `json:"filters"`
            Operators map[string]string  `json:"operators"`
         
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
    
    // CollectionQuery used by Catalog
    type CollectionQuery struct {

        
            Attribute string  `json:"attribute"`
            Value []interface{}  `json:"value"`
            Op string  `json:"op"`
         
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

        
            AspectRatio string  `json:"aspect_ratio"`
            Type string  `json:"type"`
            URL string  `json:"url"`
         
    }
    
    // ImageUrls used by Catalog
    type ImageUrls struct {

        
            Landscape BannerImage  `json:"landscape"`
            Portrait BannerImage  `json:"portrait"`
         
    }
    
    // Media used by Catalog
    type Media struct {

        
            Meta map[string]interface{}  `json:"meta"`
            Type string  `json:"type"`
            URL string  `json:"url"`
         
    }
    
    // GetCollectionDetailNest used by Catalog
    type GetCollectionDetailNest struct {

        
            Badge map[string]interface{}  `json:"badge"`
            Query []CollectionQuery  `json:"query"`
            AllowFacets bool  `json:"allow_facets"`
            Name string  `json:"name"`
            AppID string  `json:"app_id"`
            Slug string  `json:"slug"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            IsActive bool  `json:"is_active"`
            Action Action  `json:"action"`
            Tag []string  `json:"tag"`
            Schedule map[string]interface{}  `json:"_schedule"`
            AllowSort bool  `json:"allow_sort"`
            Cron map[string]interface{}  `json:"cron"`
            Banners ImageUrls  `json:"banners"`
            Meta map[string]interface{}  `json:"meta"`
            Description string  `json:"description"`
            UID string  `json:"uid"`
            Logo Media  `json:"logo"`
            Priority float64  `json:"priority"`
            Type string  `json:"type"`
         
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
    
    // UserInfo used by Catalog
    type UserInfo struct {

        
            Email string  `json:"email"`
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
            UID string  `json:"uid"`
         
    }
    
    // SeoDetail used by Catalog
    type SeoDetail struct {

        
            Description string  `json:"description"`
            Title string  `json:"title"`
         
    }
    
    // NextSchedule used by Catalog
    type NextSchedule struct {

        
            End string  `json:"end"`
            Start string  `json:"start"`
         
    }
    
    // CollectionSchedule used by Catalog
    type CollectionSchedule struct {

        
            Cron string  `json:"cron"`
            End string  `json:"end"`
            NextSchedule []NextSchedule  `json:"next_schedule"`
            Duration float64  `json:"duration"`
            Start string  `json:"start"`
         
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
    
    // CreateCollection used by Catalog
    type CreateCollection struct {

        
            Badge CollectionBadge  `json:"badge"`
            CreatedBy UserInfo  `json:"created_by"`
            Query []CollectionQuery  `json:"query"`
            AllowFacets bool  `json:"allow_facets"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            SortOn string  `json:"sort_on"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            IsActive bool  `json:"is_active"`
            ModifiedBy UserInfo  `json:"modified_by"`
            Seo SeoDetail  `json:"seo"`
            Priority float64  `json:"priority"`
            Schedule CollectionSchedule  `json:"_schedule"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            AllowSort bool  `json:"allow_sort"`
            Banners CollectionBanner  `json:"banners"`
            Meta map[string]interface{}  `json:"meta"`
            IsVisible bool  `json:"is_visible"`
            Description string  `json:"description"`
            Published bool  `json:"published"`
            Tags []string  `json:"tags"`
            Logo CollectionImage  `json:"logo"`
            AppID string  `json:"app_id"`
            Type string  `json:"type"`
         
    }
    
    // CollectionCreateResponse used by Catalog
    type CollectionCreateResponse struct {

        
            Badge map[string]interface{}  `json:"badge"`
            Query []CollectionQuery  `json:"query"`
            AllowFacets bool  `json:"allow_facets"`
            Name string  `json:"name"`
            AppID string  `json:"app_id"`
            Slug string  `json:"slug"`
            SortOn string  `json:"sort_on"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            IsActive bool  `json:"is_active"`
            Tag []string  `json:"tag"`
            Schedule map[string]interface{}  `json:"_schedule"`
            AllowSort bool  `json:"allow_sort"`
            Cron map[string]interface{}  `json:"cron"`
            Banners ImageUrls  `json:"banners"`
            Meta map[string]interface{}  `json:"meta"`
            Description string  `json:"description"`
            Logo BannerImage  `json:"logo"`
            Priority float64  `json:"priority"`
            Type string  `json:"type"`
         
    }
    
    // CollectionDetailResponse used by Catalog
    type CollectionDetailResponse struct {

        
            Badge map[string]interface{}  `json:"badge"`
            Query []CollectionQuery  `json:"query"`
            AllowFacets bool  `json:"allow_facets"`
            Name string  `json:"name"`
            AppID string  `json:"app_id"`
            Slug string  `json:"slug"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            IsActive bool  `json:"is_active"`
            Tag []string  `json:"tag"`
            Schedule map[string]interface{}  `json:"_schedule"`
            AllowSort bool  `json:"allow_sort"`
            Cron map[string]interface{}  `json:"cron"`
            Banners ImageUrls  `json:"banners"`
            Meta map[string]interface{}  `json:"meta"`
            Description string  `json:"description"`
            UID string  `json:"uid"`
            Logo Media  `json:"logo"`
            Priority float64  `json:"priority"`
            Type string  `json:"type"`
         
    }
    
    // UpdateCollection used by Catalog
    type UpdateCollection struct {

        
            Badge CollectionBadge  `json:"badge"`
            Query []CollectionQuery  `json:"query"`
            AllowFacets bool  `json:"allow_facets"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            SortOn string  `json:"sort_on"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            IsActive bool  `json:"is_active"`
            ModifiedBy UserInfo  `json:"modified_by"`
            Seo SeoDetail  `json:"seo"`
            Schedule CollectionSchedule  `json:"_schedule"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            AllowSort bool  `json:"allow_sort"`
            Banners CollectionBanner  `json:"banners"`
            Meta map[string]interface{}  `json:"meta"`
            IsVisible bool  `json:"is_visible"`
            Description string  `json:"description"`
            Published bool  `json:"published"`
            Tags []string  `json:"tags"`
            Logo CollectionImage  `json:"logo"`
            Priority float64  `json:"priority"`
            Type string  `json:"type"`
         
    }
    
    // ProductDetailAttribute used by Catalog
    type ProductDetailAttribute struct {

        
            Key string  `json:"key"`
            Value string  `json:"value"`
            Type string  `json:"type"`
         
    }
    
    // ProductDetailGroupedAttribute used by Catalog
    type ProductDetailGroupedAttribute struct {

        
            Details []ProductDetailAttribute  `json:"details"`
            Title string  `json:"title"`
         
    }
    
    // Price1 used by Catalog
    type Price1 struct {

        
            Max float64  `json:"max"`
            CurrencySymbol string  `json:"currency_symbol"`
            Min float64  `json:"min"`
            CurrencyCode string  `json:"currency_code"`
         
    }
    
    // ProductListingPrice used by Catalog
    type ProductListingPrice struct {

        
            Effective Price1  `json:"effective"`
            Marked Price1  `json:"marked"`
         
    }
    
    // ProductBrand used by Catalog
    type ProductBrand struct {

        
            Name string  `json:"name"`
            Logo Media  `json:"logo"`
            Action Action  `json:"action"`
            UID float64  `json:"uid"`
         
    }
    
    // ProductListingDetail used by Catalog
    type ProductListingDetail struct {

        
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            ImageNature string  `json:"image_nature"`
            Name string  `json:"name"`
            Attributes map[string]interface{}  `json:"attributes"`
            Price ProductListingPrice  `json:"price"`
            Brand ProductBrand  `json:"brand"`
            Rating float64  `json:"rating"`
            Slug string  `json:"slug"`
            Tryouts []string  `json:"tryouts"`
            ProductOnlineDate string  `json:"product_online_date"`
            PromoMeta map[string]interface{}  `json:"promo_meta"`
            ItemCode string  `json:"item_code"`
            Similars []string  `json:"similars"`
            Medias []Media  `json:"medias"`
            RatingCount float64  `json:"rating_count"`
            Color string  `json:"color"`
            ItemType string  `json:"item_type"`
            Description string  `json:"description"`
            ShortDescription string  `json:"short_description"`
            UID float64  `json:"uid"`
            Highlights []string  `json:"highlights"`
            HasVariant bool  `json:"has_variant"`
            Discount string  `json:"discount"`
            Sellable bool  `json:"sellable"`
            Type string  `json:"type"`
         
    }
    
    // GetCollectionItemsResponse used by Catalog
    type GetCollectionItemsResponse struct {

        
            SortOn []ProductSortOn  `json:"sort_on"`
            Filters []ProductFilters  `json:"filters"`
            Items []ProductListingDetail  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // CollectionItem used by Catalog
    type CollectionItem struct {

        
            Action string  `json:"action"`
            Priority float64  `json:"priority"`
            ItemID float64  `json:"item_id"`
         
    }
    
    // CollectionItemUpdate used by Catalog
    type CollectionItemUpdate struct {

        
            AllowSort bool  `json:"allow_sort"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Query []CollectionQuery  `json:"query"`
            Items []CollectionItem  `json:"items"`
            AllowFacets bool  `json:"allow_facets"`
            Type string  `json:"type"`
         
    }
    
    // UpdatedResponse used by Catalog
    type UpdatedResponse struct {

        
            Message string  `json:"message"`
            ItemsNotUpdated []float64  `json:"items_not_updated"`
         
    }
    
    // CatalogInsightItem used by Catalog
    type CatalogInsightItem struct {

        
            OutOfStockCount float64  `json:"out_of_stock_count"`
            Count float64  `json:"count"`
            SellableCount float64  `json:"sellable_count"`
         
    }
    
    // CatalogInsightBrand used by Catalog
    type CatalogInsightBrand struct {

        
            TotalArticles float64  `json:"total_articles"`
            ArticleFreshness float64  `json:"article_freshness"`
            TotalSizes float64  `json:"total_sizes"`
            AvailableSizes float64  `json:"available_sizes"`
            Name string  `json:"name"`
            AvailableArticles float64  `json:"available_articles"`
         
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

        
            Enabled bool  `json:"enabled"`
            CompanyID float64  `json:"company_id"`
            Platform string  `json:"platform"`
            OptLevel string  `json:"opt_level"`
            StoreIds []float64  `json:"store_ids"`
            BrandIds []float64  `json:"brand_ids"`
         
    }
    
    // CompanyOptIn used by Catalog
    type CompanyOptIn struct {

        
            CompanyID float64  `json:"company_id"`
            Enabled bool  `json:"enabled"`
            Platform string  `json:"platform"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            OptLevel string  `json:"opt_level"`
            CreatedOn float64  `json:"created_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            ModifiedOn float64  `json:"modified_on"`
            StoreIds []float64  `json:"store_ids"`
            BrandIds []float64  `json:"brand_ids"`
         
    }
    
    // GetOptInPlatform used by Catalog
    type GetOptInPlatform struct {

        
            Items []CompanyOptIn  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // OptinCompanyDetail used by Catalog
    type OptinCompanyDetail struct {

        
            Name string  `json:"name"`
            UID float64  `json:"uid"`
            BusinessType string  `json:"business_type"`
            CompanyType string  `json:"company_type"`
         
    }
    
    // CompanyBrandDetail used by Catalog
    type CompanyBrandDetail struct {

        
            BrandID float64  `json:"brand_id"`
            CompanyID float64  `json:"company_id"`
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

        
            StoreCode string  `json:"store_code"`
            CompanyID float64  `json:"company_id"`
            UID float64  `json:"uid"`
            AdditionalContacts []map[string]interface{}  `json:"additional_contacts"`
            CreatedOn string  `json:"created_on"`
            StoreType string  `json:"store_type"`
            Address map[string]interface{}  `json:"address"`
            ModifiedOn string  `json:"modified_on"`
            Timing map[string]interface{}  `json:"timing"`
            Name string  `json:"name"`
            Documents []map[string]interface{}  `json:"documents"`
            DisplayName string  `json:"display_name"`
            Manager map[string]interface{}  `json:"manager"`
         
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

        
            Format string  `json:"format"`
            AllowedValues []string  `json:"allowed_values"`
            Mandatory bool  `json:"mandatory"`
            Range AttributeSchemaRange  `json:"range"`
            Multi bool  `json:"multi"`
            Type string  `json:"type"`
         
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

        
            ID string  `json:"id"`
            Schema AttributeMaster  `json:"schema"`
            Slug string  `json:"slug"`
            EnabledForEndConsumer bool  `json:"enabled_for_end_consumer"`
            Details AttributeMasterDetails  `json:"details"`
            IsNested bool  `json:"is_nested"`
            Filters AttributeMasterFilter  `json:"filters"`
            Departments []string  `json:"departments"`
            Meta AttributeMasterMeta  `json:"meta"`
            Name string  `json:"name"`
            Logo string  `json:"logo"`
            Description string  `json:"description"`
         
    }
    
    // CategoriesResponse used by Catalog
    type CategoriesResponse struct {

        
            UID float64  `json:"uid"`
            Slug string  `json:"slug"`
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

        
            Code string  `json:"code"`
            Errors map[string]interface{}  `json:"errors"`
            Message string  `json:"message"`
            Meta map[string]interface{}  `json:"meta"`
            Status float64  `json:"status"`
         
    }
    
    // UserSerializer used by Catalog
    type UserSerializer struct {

        
            UID string  `json:"uid"`
            Contact string  `json:"contact"`
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
            ID string  `json:"_id"`
         
    }
    
    // GetDepartment used by Catalog
    type GetDepartment struct {

        
            UID float64  `json:"uid"`
            Synonyms []string  `json:"synonyms"`
            CreatedBy UserSerializer  `json:"created_by"`
            Slug string  `json:"slug"`
            PriorityOrder float64  `json:"priority_order"`
            PageSize float64  `json:"page_size"`
            CreatedOn string  `json:"created_on"`
            IsActive bool  `json:"is_active"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            ItemType string  `json:"item_type"`
            ModifiedOn string  `json:"modified_on"`
            Search string  `json:"search"`
            Name string  `json:"name"`
            Logo string  `json:"logo"`
            PageNo float64  `json:"page_no"`
         
    }
    
    // DepartmentsResponse used by Catalog
    type DepartmentsResponse struct {

        
            Items []GetDepartment  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // DepartmentErrorResponse used by Catalog
    type DepartmentErrorResponse struct {

        
            Code string  `json:"code"`
            Errors map[string]interface{}  `json:"errors"`
            Message string  `json:"message"`
            Meta map[string]interface{}  `json:"meta"`
            Status float64  `json:"status"`
         
    }
    
    // DepartmentCreateUpdate used by Catalog
    type DepartmentCreateUpdate struct {

        
            Cls string  `json:"_cls"`
            UID float64  `json:"uid"`
            Synonyms []string  `json:"synonyms"`
            PriorityOrder float64  `json:"priority_order"`
            Slug string  `json:"slug"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            IsActive bool  `json:"is_active"`
            Tags []string  `json:"tags"`
            Name string  `json:"name"`
            Logo string  `json:"logo"`
            Platforms map[string]interface{}  `json:"platforms"`
         
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
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
         
    }
    
    // DepartmentModel used by Catalog
    type DepartmentModel struct {

        
            Cls interface{}  `json:"_cls"`
            UID float64  `json:"uid"`
            CreatedBy UserDetail  `json:"created_by"`
            Synonyms []interface{}  `json:"synonyms"`
            Slug interface{}  `json:"slug"`
            PriorityOrder float64  `json:"priority_order"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            CreatedOn string  `json:"created_on"`
            IsActive bool  `json:"is_active"`
            ModifiedBy UserDetail  `json:"modified_by"`
            ModifiedOn string  `json:"modified_on"`
            VerifiedBy UserDetail  `json:"verified_by"`
            VerifiedOn string  `json:"verified_on"`
            Name interface{}  `json:"name"`
            Logo interface{}  `json:"logo"`
            ID interface{}  `json:"_id"`
         
    }
    
    // ProductTemplate used by Catalog
    type ProductTemplate struct {

        
            Categories []string  `json:"categories"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Slug string  `json:"slug"`
            Logo string  `json:"logo"`
            CreatedOn string  `json:"created_on"`
            IsActive bool  `json:"is_active"`
            IsArchived bool  `json:"is_archived"`
            IsPhysical bool  `json:"is_physical"`
            Tag string  `json:"tag"`
            Departments []string  `json:"departments"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            ModifiedOn string  `json:"modified_on"`
            Name string  `json:"name"`
            Attributes []string  `json:"attributes"`
            Description string  `json:"description"`
            IsExpirable bool  `json:"is_expirable"`
         
    }
    
    // TemplatesResponse used by Catalog
    type TemplatesResponse struct {

        
            Items ProductTemplate  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // TemplateDetails used by Catalog
    type TemplateDetails struct {

        
            ID string  `json:"id"`
            Categories []string  `json:"categories"`
            Slug string  `json:"slug"`
            Logo string  `json:"logo"`
            IsActive bool  `json:"is_active"`
            IsArchived bool  `json:"is_archived"`
            IsPhysical bool  `json:"is_physical"`
            Tag string  `json:"tag"`
            Departments []string  `json:"departments"`
            Name string  `json:"name"`
            Attributes []string  `json:"attributes"`
            Description string  `json:"description"`
            IsExpirable bool  `json:"is_expirable"`
         
    }
    
    // Properties used by Catalog
    type Properties struct {

        
            IsDependent map[string]interface{}  `json:"is_dependent"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            ReturnConfig map[string]interface{}  `json:"return_config"`
            Name map[string]interface{}  `json:"name"`
            Sizes map[string]interface{}  `json:"sizes"`
            SizeGuide map[string]interface{}  `json:"size_guide"`
            Currency map[string]interface{}  `json:"currency"`
            Slug map[string]interface{}  `json:"slug"`
            HsnCode map[string]interface{}  `json:"hsn_code"`
            CategorySlug map[string]interface{}  `json:"category_slug"`
            IsActive map[string]interface{}  `json:"is_active"`
            ProductGroupTag map[string]interface{}  `json:"product_group_tag"`
            ProductPublish map[string]interface{}  `json:"product_publish"`
            ItemCode map[string]interface{}  `json:"item_code"`
            BrandUID map[string]interface{}  `json:"brand_uid"`
            ItemType map[string]interface{}  `json:"item_type"`
            Media map[string]interface{}  `json:"media"`
            TraderType map[string]interface{}  `json:"trader_type"`
            Description map[string]interface{}  `json:"description"`
            Variants map[string]interface{}  `json:"variants"`
            ShortDescription map[string]interface{}  `json:"short_description"`
            Highlights map[string]interface{}  `json:"highlights"`
            Command map[string]interface{}  `json:"command"`
            NoOfBoxes map[string]interface{}  `json:"no_of_boxes"`
            CountryOfOrigin map[string]interface{}  `json:"country_of_origin"`
            MultiSize map[string]interface{}  `json:"multi_size"`
            Tags map[string]interface{}  `json:"tags"`
            Trader map[string]interface{}  `json:"trader"`
         
    }
    
    // GlobalValidation used by Catalog
    type GlobalValidation struct {

        
            Title string  `json:"title"`
            Description string  `json:"description"`
            Definitions map[string]interface{}  `json:"definitions"`
            Required []string  `json:"required"`
            Type string  `json:"type"`
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

        
            Email string  `json:"email"`
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
            UID string  `json:"uid"`
         
    }
    
    // ProductTemplateExportResponse used by Catalog
    type ProductTemplateExportResponse struct {

        
            CreatedBy UserInfo1  `json:"created_by"`
            URL string  `json:"url"`
            Filters map[string]interface{}  `json:"filters"`
            CompletedOn string  `json:"completed_on"`
            NotificationEmails []string  `json:"notification_emails"`
            SellerID float64  `json:"seller_id"`
            ModifiedOn string  `json:"modified_on"`
            TaskID string  `json:"task_id"`
            Type string  `json:"type"`
            Status string  `json:"status"`
         
    }
    
    // ProductDownloadsResponse used by Catalog
    type ProductDownloadsResponse struct {

        
            Items []ProductTemplateExportResponse  `json:"items"`
         
    }
    
    // ProductTemplateExportFilterRequest used by Catalog
    type ProductTemplateExportFilterRequest struct {

        
            Templates []string  `json:"templates"`
            CatalogueTypes []string  `json:"catalogue_types"`
            FromDate string  `json:"from_date"`
            Brands []string  `json:"brands"`
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
    
    // Hierarchy used by Catalog
    type Hierarchy struct {

        
            Department float64  `json:"department"`
            L2 float64  `json:"l2"`
            L1 float64  `json:"l1"`
         
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
    
    // Media1 used by Catalog
    type Media1 struct {

        
            Logo string  `json:"logo"`
            Landscape string  `json:"landscape"`
            Portrait string  `json:"portrait"`
         
    }
    
    // Category used by Catalog
    type Category struct {

        
            ID string  `json:"id"`
            UID float64  `json:"uid"`
            Level float64  `json:"level"`
            Synonyms []string  `json:"synonyms"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Slug string  `json:"slug"`
            Tryouts []string  `json:"tryouts"`
            CreatedOn string  `json:"created_on"`
            IsActive bool  `json:"is_active"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Hierarchy []Hierarchy  `json:"hierarchy"`
            Departments []float64  `json:"departments"`
            Marketplaces CategoryMapping  `json:"marketplaces"`
            ModifiedOn string  `json:"modified_on"`
            Media Media1  `json:"media"`
            Name string  `json:"name"`
            Priority float64  `json:"priority"`
         
    }
    
    // CategoryResponse used by Catalog
    type CategoryResponse struct {

        
            Items []Category  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // CategoryRequestBody used by Catalog
    type CategoryRequestBody struct {

        
            Level float64  `json:"level"`
            Synonyms []string  `json:"synonyms"`
            Slug string  `json:"slug"`
            Tryouts []string  `json:"tryouts"`
            IsActive bool  `json:"is_active"`
            Hierarchy []Hierarchy  `json:"hierarchy"`
            Departments []float64  `json:"departments"`
            Marketplaces CategoryMapping  `json:"marketplaces"`
            Media Media1  `json:"media"`
            Name string  `json:"name"`
            Priority float64  `json:"priority"`
         
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

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // ReturnConfigResponse used by Catalog
    type ReturnConfigResponse struct {

        
            Time float64  `json:"time"`
            Returnable bool  `json:"returnable"`
            Unit string  `json:"unit"`
         
    }
    
    // NetQuantityResponse used by Catalog
    type NetQuantityResponse struct {

        
            Unit string  `json:"unit"`
            Value float64  `json:"value"`
         
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
    
    // ProductPublish used by Catalog
    type ProductPublish struct {

        
            ProductOnlineDate string  `json:"product_online_date"`
            IsSet bool  `json:"is_set"`
         
    }
    
    // Image used by Catalog
    type Image struct {

        
            AspectRatio string  `json:"aspect_ratio"`
            URL string  `json:"url"`
            AspectRatioF float64  `json:"aspect_ratio_f"`
            SecureURL string  `json:"secure_url"`
         
    }
    
    // VerifiedBy used by Catalog
    type VerifiedBy struct {

        
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
         
    }
    
    // Trader used by Catalog
    type Trader struct {

        
            Name interface{}  `json:"name"`
            Address []string  `json:"address"`
            Type string  `json:"type"`
         
    }
    
    // ProductSchemaV2 used by Catalog
    type ProductSchemaV2 struct {

        
            CreatedBy map[string]interface{}  `json:"created_by"`
            IsDependent bool  `json:"is_dependent"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            ReturnConfig ReturnConfigResponse  `json:"return_config"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            NetQuantity NetQuantityResponse  `json:"net_quantity"`
            ImageNature string  `json:"image_nature"`
            Name string  `json:"name"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            Sizes []map[string]interface{}  `json:"sizes"`
            Attributes map[string]interface{}  `json:"attributes"`
            Brand Brand  `json:"brand"`
            SizeGuide string  `json:"size_guide"`
            Currency string  `json:"currency"`
            PrimaryColor string  `json:"primary_color"`
            Slug string  `json:"slug"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            CategorySlug string  `json:"category_slug"`
            HsnCode string  `json:"hsn_code"`
            CreatedOn string  `json:"created_on"`
            IsActive bool  `json:"is_active"`
            IsPhysical bool  `json:"is_physical"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            L3Mapping []string  `json:"l3_mapping"`
            Departments []float64  `json:"departments"`
            ModifiedOn string  `json:"modified_on"`
            AllCompanyIds []float64  `json:"all_company_ids"`
            ProductGroupTag []string  `json:"product_group_tag"`
            TemplateTag string  `json:"template_tag"`
            ProductPublish ProductPublish  `json:"product_publish"`
            AllIdentifiers []string  `json:"all_identifiers"`
            ItemCode string  `json:"item_code"`
            Stage string  `json:"stage"`
            ID string  `json:"id"`
            IsImageLessProduct bool  `json:"is_image_less_product"`
            Pending string  `json:"pending"`
            IsSet bool  `json:"is_set"`
            Images []Image  `json:"images"`
            Color string  `json:"color"`
            BrandUID float64  `json:"brand_uid"`
            ItemType string  `json:"item_type"`
            AllSizes []map[string]interface{}  `json:"all_sizes"`
            Media []Media  `json:"media"`
            VerifiedOn string  `json:"verified_on"`
            Description string  `json:"description"`
            ShortDescription string  `json:"short_description"`
            Variants map[string]interface{}  `json:"variants"`
            IsExpirable bool  `json:"is_expirable"`
            VariantGroup map[string]interface{}  `json:"variant_group"`
            CompanyID float64  `json:"company_id"`
            Highlights []string  `json:"highlights"`
            UID float64  `json:"uid"`
            VariantMedia map[string]interface{}  `json:"variant_media"`
            NoOfBoxes float64  `json:"no_of_boxes"`
            CountryOfOrigin string  `json:"country_of_origin"`
            Moq map[string]interface{}  `json:"moq"`
            Category map[string]interface{}  `json:"category"`
            MultiSize bool  `json:"multi_size"`
            VerifiedBy VerifiedBy  `json:"verified_by"`
            Tags []string  `json:"tags"`
            Trader []Trader  `json:"trader"`
            CategoryUID float64  `json:"category_uid"`
         
    }
    
    // ProductListingResponseV2 used by Catalog
    type ProductListingResponseV2 struct {

        
            Items []ProductSchemaV2  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // TeaserTag used by Catalog
    type TeaserTag struct {

        
            Tag string  `json:"tag"`
            URL string  `json:"url"`
         
    }
    
    // CustomOrder used by Catalog
    type CustomOrder struct {

        
            ManufacturingTime float64  `json:"manufacturing_time"`
            ManufacturingTimeUnit string  `json:"manufacturing_time_unit"`
            IsCustomOrder bool  `json:"is_custom_order"`
         
    }
    
    // ReturnConfig used by Catalog
    type ReturnConfig struct {

        
            Time float64  `json:"time"`
            Returnable bool  `json:"returnable"`
            Unit string  `json:"unit"`
         
    }
    
    // NetQuantity used by Catalog
    type NetQuantity struct {

        
            Unit interface{}  `json:"unit"`
            Value float64  `json:"value"`
         
    }
    
    // TaxIdentifier used by Catalog
    type TaxIdentifier struct {

        
            HsnCode string  `json:"hsn_code"`
            ReportingHsn string  `json:"reporting_hsn"`
            HsnCodeID string  `json:"hsn_code_id"`
         
    }
    
    // ProductPublish1 used by Catalog
    type ProductPublish1 struct {

        
            ProductOnlineDate string  `json:"product_online_date"`
            IsSet bool  `json:"is_set"`
         
    }
    
    // ProductCreateUpdateSchemaV2 used by Catalog
    type ProductCreateUpdateSchemaV2 struct {

        
            IsDependent bool  `json:"is_dependent"`
            ChangeRequestID interface{}  `json:"change_request_id"`
            TeaserTag TeaserTag  `json:"teaser_tag"`
            CustomOrder CustomOrder  `json:"custom_order"`
            ReturnConfig ReturnConfig  `json:"return_config"`
            NetQuantity NetQuantity  `json:"net_quantity"`
            Name string  `json:"name"`
            TaxIdentifier TaxIdentifier  `json:"tax_identifier"`
            Attributes map[string]interface{}  `json:"attributes"`
            Sizes []map[string]interface{}  `json:"sizes"`
            SizeGuide string  `json:"size_guide"`
            Currency string  `json:"currency"`
            Slug string  `json:"slug"`
            CategorySlug string  `json:"category_slug"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            BulkJobID string  `json:"bulk_job_id"`
            IsActive bool  `json:"is_active"`
            Action string  `json:"action"`
            Departments []float64  `json:"departments"`
            ProductGroupTag []string  `json:"product_group_tag"`
            TemplateTag string  `json:"template_tag"`
            ProductPublish ProductPublish1  `json:"product_publish"`
            ItemCode string  `json:"item_code"`
            IsImageLessProduct bool  `json:"is_image_less_product"`
            IsSet bool  `json:"is_set"`
            Requester string  `json:"requester"`
            BrandUID float64  `json:"brand_uid"`
            ItemType string  `json:"item_type"`
            Media []Media  `json:"media"`
            VariantGroup map[string]interface{}  `json:"variant_group"`
            Description string  `json:"description"`
            ShortDescription string  `json:"short_description"`
            Variants map[string]interface{}  `json:"variants"`
            CompanyID float64  `json:"company_id"`
            Highlights []string  `json:"highlights"`
            UID float64  `json:"uid"`
            VariantMedia map[string]interface{}  `json:"variant_media"`
            NoOfBoxes float64  `json:"no_of_boxes"`
            CountryOfOrigin string  `json:"country_of_origin"`
            MultiSize bool  `json:"multi_size"`
            Tags []string  `json:"tags"`
            Trader []Trader  `json:"trader"`
         
    }
    
    // ProductVariants used by Catalog
    type ProductVariants struct {

        
            UID float64  `json:"uid"`
            BrandUID float64  `json:"brand_uid"`
            Media []Media  `json:"media"`
            Name string  `json:"name"`
            ItemCode string  `json:"item_code"`
            CategoryUID float64  `json:"category_uid"`
         
    }
    
    // ProductVariantsResponse used by Catalog
    type ProductVariantsResponse struct {

        
            Variants []ProductVariants  `json:"variants"`
            Page Page  `json:"page"`
         
    }
    
    // AttributeMasterSerializer used by Catalog
    type AttributeMasterSerializer struct {

        
            CreatedBy map[string]interface{}  `json:"created_by"`
            Variant bool  `json:"variant"`
            Name string  `json:"name"`
            Unit string  `json:"unit"`
            Slug string  `json:"slug"`
            CreatedOn string  `json:"created_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Departments []string  `json:"departments"`
            ModifiedOn string  `json:"modified_on"`
            RawKey string  `json:"raw_key"`
            EnabledForEndConsumer bool  `json:"enabled_for_end_consumer"`
            Details AttributeMasterDetails  `json:"details"`
            IsNested bool  `json:"is_nested"`
            Description string  `json:"description"`
            Schema AttributeMaster  `json:"schema"`
            Example string  `json:"example"`
            Synonyms map[string]interface{}  `json:"synonyms"`
            Filters AttributeMasterFilter  `json:"filters"`
            Tags []string  `json:"tags"`
            Suggestion string  `json:"suggestion"`
            Logo string  `json:"logo"`
         
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
            Size interface{}  `json:"size"`
            ItemHeight float64  `json:"item_height"`
            ItemDimensionsUnitOfMeasure string  `json:"item_dimensions_unit_of_measure"`
            ItemWeightUnitOfMeasure interface{}  `json:"item_weight_unit_of_measure"`
            ItemLength float64  `json:"item_length"`
            ItemWidth float64  `json:"item_width"`
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

        
            CreatedBy map[string]interface{}  `json:"created_by"`
            IsDependent bool  `json:"is_dependent"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            ReturnConfig ReturnConfigResponse  `json:"return_config"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            NetQuantity NetQuantityResponse  `json:"net_quantity"`
            ImageNature string  `json:"image_nature"`
            Name string  `json:"name"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            Sizes []map[string]interface{}  `json:"sizes"`
            Attributes map[string]interface{}  `json:"attributes"`
            Brand Brand  `json:"brand"`
            SizeGuide string  `json:"size_guide"`
            Currency string  `json:"currency"`
            PrimaryColor string  `json:"primary_color"`
            Slug string  `json:"slug"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            CategorySlug string  `json:"category_slug"`
            HsnCode string  `json:"hsn_code"`
            CreatedOn string  `json:"created_on"`
            IsActive bool  `json:"is_active"`
            IsPhysical bool  `json:"is_physical"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            L3Mapping []string  `json:"l3_mapping"`
            Departments []float64  `json:"departments"`
            ModifiedOn string  `json:"modified_on"`
            AllCompanyIds []float64  `json:"all_company_ids"`
            ProductGroupTag []string  `json:"product_group_tag"`
            TemplateTag string  `json:"template_tag"`
            ProductPublish ProductPublished  `json:"product_publish"`
            AllIdentifiers []string  `json:"all_identifiers"`
            ItemCode string  `json:"item_code"`
            Stage string  `json:"stage"`
            ID string  `json:"id"`
            IsImageLessProduct bool  `json:"is_image_less_product"`
            Pending string  `json:"pending"`
            IsSet bool  `json:"is_set"`
            Images []Image  `json:"images"`
            Color string  `json:"color"`
            BrandUID float64  `json:"brand_uid"`
            ItemType string  `json:"item_type"`
            AllSizes []map[string]interface{}  `json:"all_sizes"`
            Media []Media  `json:"media"`
            VerifiedOn string  `json:"verified_on"`
            Description string  `json:"description"`
            ShortDescription string  `json:"short_description"`
            Variants map[string]interface{}  `json:"variants"`
            IsExpirable bool  `json:"is_expirable"`
            VariantGroup map[string]interface{}  `json:"variant_group"`
            CompanyID float64  `json:"company_id"`
            Highlights []string  `json:"highlights"`
            UID float64  `json:"uid"`
            VariantMedia map[string]interface{}  `json:"variant_media"`
            NoOfBoxes float64  `json:"no_of_boxes"`
            CountryOfOrigin string  `json:"country_of_origin"`
            Moq map[string]interface{}  `json:"moq"`
            Category map[string]interface{}  `json:"category"`
            MultiSize bool  `json:"multi_size"`
            VerifiedBy VerifiedBy  `json:"verified_by"`
            Tags []string  `json:"tags"`
            Trader []Trader  `json:"trader"`
            CategoryUID float64  `json:"category_uid"`
         
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

        
            FailedRecords []string  `json:"failed_records"`
            CompanyID float64  `json:"company_id"`
            CancelledRecords []string  `json:"cancelled_records"`
            CreatedBy UserDetail1  `json:"created_by"`
            Total float64  `json:"total"`
            Template ProductTemplate  `json:"template"`
            CreatedOn string  `json:"created_on"`
            IsActive bool  `json:"is_active"`
            FilePath string  `json:"file_path"`
            Succeed float64  `json:"succeed"`
            ModifiedBy UserDetail1  `json:"modified_by"`
            ModifiedOn string  `json:"modified_on"`
            TemplateTag string  `json:"template_tag"`
            Failed float64  `json:"failed"`
            Cancelled float64  `json:"cancelled"`
            Stage string  `json:"stage"`
         
    }
    
    // ProductBulkRequestList used by Catalog
    type ProductBulkRequestList struct {

        
            Items ProductBulkRequest  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // BulkJob used by Catalog
    type BulkJob struct {

        
            FailedRecords []map[string]interface{}  `json:"failed_records"`
            CompanyID float64  `json:"company_id"`
            CancelledRecords []map[string]interface{}  `json:"cancelled_records"`
            CustomTemplateTag string  `json:"custom_template_tag"`
            CreatedBy UserInfo1  `json:"created_by"`
            Total float64  `json:"total"`
            CreatedOn string  `json:"created_on"`
            IsActive bool  `json:"is_active"`
            ModifiedBy UserInfo1  `json:"modified_by"`
            Succeed float64  `json:"succeed"`
            FilePath string  `json:"file_path"`
            ModifiedOn string  `json:"modified_on"`
            TrackingURL string  `json:"tracking_url"`
            TemplateTag string  `json:"template_tag"`
            Failed float64  `json:"failed"`
            Cancelled float64  `json:"cancelled"`
            Stage string  `json:"stage"`
         
    }
    
    // BulkResponse used by Catalog
    type BulkResponse struct {

        
            CreatedBy UserInfo1  `json:"created_by"`
            BatchID string  `json:"batch_id"`
            CreatedOn string  `json:"created_on"`
            IsActive bool  `json:"is_active"`
            ModifiedBy UserInfo1  `json:"modified_by"`
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

        
            CompanyID float64  `json:"company_id"`
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
         
    }
    
    // Items used by Catalog
    type Items struct {

        
            FailedRecords []string  `json:"failed_records"`
            ID string  `json:"id"`
            CompanyID float64  `json:"company_id"`
            Retry float64  `json:"retry"`
            CancelledRecords []string  `json:"cancelled_records"`
            CreatedBy UserCommon  `json:"created_by"`
            Total float64  `json:"total"`
            CreatedOn string  `json:"created_on"`
            FilePath string  `json:"file_path"`
            ModifiedBy UserCommon  `json:"modified_by"`
            Succeed float64  `json:"succeed"`
            IsActive bool  `json:"is_active"`
            ModifiedOn string  `json:"modified_on"`
            TrackingURL string  `json:"tracking_url"`
            Failed float64  `json:"failed"`
            Cancelled float64  `json:"cancelled"`
            Stage string  `json:"stage"`
         
    }
    
    // BulkAssetResponse used by Catalog
    type BulkAssetResponse struct {

        
            Items []Items  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // ProductBulkAssets used by Catalog
    type ProductBulkAssets struct {

        
            CompanyID float64  `json:"company_id"`
            User map[string]interface{}  `json:"user"`
            BatchID string  `json:"batch_id"`
            URL string  `json:"url"`
         
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
    
    // InventoryResponse used by Catalog
    type InventoryResponse struct {

        
            Currency string  `json:"currency"`
            UID string  `json:"uid"`
            Size string  `json:"size"`
            SellableQuantity float64  `json:"sellable_quantity"`
            SellerIdentifier string  `json:"seller_identifier"`
            Store map[string]interface{}  `json:"store"`
            Quantity float64  `json:"quantity"`
            InventoryUpdatedOn string  `json:"inventory_updated_on"`
            PriceEffective float64  `json:"price_effective"`
            PriceTransfer float64  `json:"price_transfer"`
            ItemID float64  `json:"item_id"`
            Identifiers map[string]interface{}  `json:"identifiers"`
            Price float64  `json:"price"`
         
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

        
            StoreCode string  `json:"store_code"`
            Set InventorySet  `json:"set"`
            ItemWeight float64  `json:"item_weight"`
            Currency string  `json:"currency"`
            Size interface{}  `json:"size"`
            ItemHeight float64  `json:"item_height"`
            ItemDimensionsUnitOfMeasure string  `json:"item_dimensions_unit_of_measure"`
            ItemWeightUnitOfMeasure string  `json:"item_weight_unit_of_measure"`
            Quantity float64  `json:"quantity"`
            IsSet bool  `json:"is_set"`
            ItemLength float64  `json:"item_length"`
            PriceEffective float64  `json:"price_effective"`
            PriceTransfer float64  `json:"price_transfer"`
            ItemWidth float64  `json:"item_width"`
            Identifiers []GTIN  `json:"identifiers"`
            ExpirationDate string  `json:"expiration_date"`
            Price float64  `json:"price"`
         
    }
    
    // ItemQuery used by Catalog
    type ItemQuery struct {

        
            BrandUID float64  `json:"brand_uid"`
            ItemCode string  `json:"item_code"`
            UID float64  `json:"uid"`
         
    }
    
    // InventoryRequest used by Catalog
    type InventoryRequest struct {

        
            Sizes []InvSize  `json:"sizes"`
            Item ItemQuery  `json:"item"`
            CompanyID float64  `json:"company_id"`
         
    }
    
    // ManufacturerResponse used by Catalog
    type ManufacturerResponse struct {

        
            IsDefault bool  `json:"is_default"`
            Address string  `json:"address"`
            Name string  `json:"name"`
         
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
    
    // PriceMeta used by Catalog
    type PriceMeta struct {

        
            Currency string  `json:"currency"`
            Effective float64  `json:"effective"`
            TpNotes map[string]interface{}  `json:"tp_notes"`
            UpdatedAt string  `json:"updated_at"`
            Transfer float64  `json:"transfer"`
            Marked float64  `json:"marked"`
         
    }
    
    // BrandMeta used by Catalog
    type BrandMeta struct {

        
            ID float64  `json:"id"`
            Name string  `json:"name"`
         
    }
    
    // StoreMeta used by Catalog
    type StoreMeta struct {

        
            ID float64  `json:"id"`
         
    }
    
    // QuantityBase used by Catalog
    type QuantityBase struct {

        
            Count float64  `json:"count"`
            UpdatedAt string  `json:"updated_at"`
         
    }
    
    // Quantities used by Catalog
    type Quantities struct {

        
            Sellable QuantityBase  `json:"sellable"`
            NotAvailable QuantityBase  `json:"not_available"`
            OrderCommitted QuantityBase  `json:"order_committed"`
            Damaged QuantityBase  `json:"damaged"`
         
    }
    
    // DimensionResponse used by Catalog
    type DimensionResponse struct {

        
            IsDefault bool  `json:"is_default"`
            Width float64  `json:"width"`
            Length float64  `json:"length"`
            Height float64  `json:"height"`
            Unit string  `json:"unit"`
         
    }
    
    // WeightResponse used by Catalog
    type WeightResponse struct {

        
            IsDefault bool  `json:"is_default"`
            Unit string  `json:"unit"`
            Shipping float64  `json:"shipping"`
         
    }
    
    // Trader1 used by Catalog
    type Trader1 struct {

        
            Name string  `json:"name"`
            Address []string  `json:"address"`
            Type string  `json:"type"`
         
    }
    
    // InventorySellerResponse used by Catalog
    type InventorySellerResponse struct {

        
            TrackInventory bool  `json:"track_inventory"`
            CreatedBy UserSerializer  `json:"created_by"`
            Manufacturer ManufacturerResponse  `json:"manufacturer"`
            ReturnConfig ReturnConfig1  `json:"return_config"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            TotalQuantity float64  `json:"total_quantity"`
            Company CompanyMeta  `json:"company"`
            Price PriceMeta  `json:"price"`
            ExpirationDate string  `json:"expiration_date"`
            Brand BrandMeta  `json:"brand"`
            Size string  `json:"size"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Store StoreMeta  `json:"store"`
            Quantities Quantities  `json:"quantities"`
            IsActive bool  `json:"is_active"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            TraceID string  `json:"trace_id"`
            Dimension DimensionResponse  `json:"dimension"`
            FyndItemCode string  `json:"fynd_item_code"`
            FyndMeta map[string]interface{}  `json:"fynd_meta"`
            Stage string  `json:"stage"`
            Set InventorySet  `json:"set"`
            Identifier map[string]interface{}  `json:"identifier"`
            IsSet bool  `json:"is_set"`
            AddedOnStore string  `json:"added_on_store"`
            Weight WeightResponse  `json:"weight"`
            Meta map[string]interface{}  `json:"meta"`
            ItemID float64  `json:"item_id"`
            SellerIdentifier string  `json:"seller_identifier"`
            UID string  `json:"uid"`
            RawMeta map[string]interface{}  `json:"raw_meta"`
            CountryOfOrigin string  `json:"country_of_origin"`
            FyndArticleCode string  `json:"fynd_article_code"`
            Tags []string  `json:"tags"`
            Fragile bool  `json:"fragile"`
            Trader []Trader1  `json:"trader"`
         
    }
    
    // InventorySellerIdentifierResponsePaginated used by Catalog
    type InventorySellerIdentifierResponsePaginated struct {

        
            Items []InventorySellerResponse  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // ManufacturerResponse1 used by Catalog
    type ManufacturerResponse1 struct {

        
            IsDefault bool  `json:"is_default"`
            Address string  `json:"address"`
            Name string  `json:"name"`
         
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
    
    // PriceArticle used by Catalog
    type PriceArticle struct {

        
            Currency string  `json:"currency"`
            Effective float64  `json:"effective"`
            TpNotes map[string]interface{}  `json:"tp_notes"`
            Transfer float64  `json:"transfer"`
            Marked float64  `json:"marked"`
         
    }
    
    // BrandMeta1 used by Catalog
    type BrandMeta1 struct {

        
            ID float64  `json:"id"`
            Name string  `json:"name"`
         
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
    
    // ArticleStoreResponse used by Catalog
    type ArticleStoreResponse struct {

        
            Name string  `json:"name"`
            StoreCode string  `json:"store_code"`
            UID float64  `json:"uid"`
            StoreType string  `json:"store_type"`
         
    }
    
    // DimensionResponse1 used by Catalog
    type DimensionResponse1 struct {

        
            Length float64  `json:"length"`
            Unit string  `json:"unit"`
            Width float64  `json:"width"`
            Height float64  `json:"height"`
         
    }
    
    // WeightResponse1 used by Catalog
    type WeightResponse1 struct {

        
            Unit string  `json:"unit"`
            Shipping float64  `json:"shipping"`
         
    }
    
    // DateMeta used by Catalog
    type DateMeta struct {

        
            AddedOnStore string  `json:"added_on_store"`
            ModifiedOn string  `json:"modified_on"`
            InventoryUpdatedOn string  `json:"inventory_updated_on"`
            CreatedOn string  `json:"created_on"`
         
    }
    
    // Trader2 used by Catalog
    type Trader2 struct {

        
            Name string  `json:"name"`
            Address []string  `json:"address"`
            Type string  `json:"type"`
         
    }
    
    // GetInventories used by Catalog
    type GetInventories struct {

        
            TrackInventory bool  `json:"track_inventory"`
            CreatedBy UserSerializer  `json:"created_by"`
            Manufacturer ManufacturerResponse1  `json:"manufacturer"`
            ReturnConfig ReturnConfig2  `json:"return_config"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            TotalQuantity float64  `json:"total_quantity"`
            ExpirationDate string  `json:"expiration_date"`
            Company CompanyMeta1  `json:"company"`
            Price PriceArticle  `json:"price"`
            Brand BrandMeta1  `json:"brand"`
            Size string  `json:"size"`
            Quantities QuantitiesArticle  `json:"quantities"`
            Store ArticleStoreResponse  `json:"store"`
            InventoryUpdatedOn string  `json:"inventory_updated_on"`
            TraceID string  `json:"trace_id"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            Dimension DimensionResponse1  `json:"dimension"`
            Stage string  `json:"stage"`
            ID string  `json:"id"`
            Identifier map[string]interface{}  `json:"identifier"`
            IsSet bool  `json:"is_set"`
            Weight WeightResponse1  `json:"weight"`
            ItemID float64  `json:"item_id"`
            SellerIdentifier string  `json:"seller_identifier"`
            UID string  `json:"uid"`
            CountryOfOrigin string  `json:"country_of_origin"`
            DateMeta DateMeta  `json:"date_meta"`
            Tags []string  `json:"tags"`
            Trader []Trader2  `json:"trader"`
            Platforms map[string]interface{}  `json:"platforms"`
         
    }
    
    // GetInventoriesResponse used by Catalog
    type GetInventoriesResponse struct {

        
            Items []GetInventories  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // BulkInventoryGetItems used by Catalog
    type BulkInventoryGetItems struct {

        
            FailedRecords []string  `json:"failed_records"`
            ID string  `json:"id"`
            CompanyID float64  `json:"company_id"`
            CancelledRecords []string  `json:"cancelled_records"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Total float64  `json:"total"`
            CreatedOn string  `json:"created_on"`
            FilePath string  `json:"file_path"`
            IsActive bool  `json:"is_active"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Succeed float64  `json:"succeed"`
            ModifiedOn string  `json:"modified_on"`
            Failed float64  `json:"failed"`
            Cancelled float64  `json:"cancelled"`
            Stage string  `json:"stage"`
         
    }
    
    // BulkInventoryGet used by Catalog
    type BulkInventoryGet struct {

        
            Items []BulkInventoryGetItems  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // InventoryJobPayload used by Catalog
    type InventoryJobPayload struct {

        
            StoreCode string  `json:"store_code"`
            Currency string  `json:"currency"`
            ItemDimensionsUnitOfMeasure string  `json:"item_dimensions_unit_of_measure"`
            ItemWeightUnitOfMeasure string  `json:"item_weight_unit_of_measure"`
            Quantity float64  `json:"quantity"`
            TotalQuantity float64  `json:"total_quantity"`
            TraceID string  `json:"trace_id"`
            PriceEffective float64  `json:"price_effective"`
            PriceMarked float64  `json:"price_marked"`
            Tags []string  `json:"tags"`
            SellerIdentifier string  `json:"seller_identifier"`
            ExpirationDate string  `json:"expiration_date"`
            Price float64  `json:"price"`
         
    }
    
    // InventoryBulkRequest used by Catalog
    type InventoryBulkRequest struct {

        
            Sizes []InventoryJobPayload  `json:"sizes"`
            CompanyID float64  `json:"company_id"`
            BatchID string  `json:"batch_id"`
            User map[string]interface{}  `json:"user"`
         
    }
    
    // InventoryExportQuantityFilter used by Catalog
    type InventoryExportQuantityFilter struct {

        
            Max float64  `json:"max"`
            Min float64  `json:"min"`
            Operators string  `json:"operators"`
         
    }
    
    // InventoryExportAdvanceOption used by Catalog
    type InventoryExportAdvanceOption struct {

        
            Quantity InventoryExportQuantityFilter  `json:"quantity"`
            ToDate string  `json:"to_date"`
            FromDate string  `json:"from_date"`
            StoreIds []float64  `json:"store_ids"`
            BrandIds []float64  `json:"brand_ids"`
         
    }
    
    // InventoryExportJob used by Catalog
    type InventoryExportJob struct {

        
            URL string  `json:"url"`
            Filters InventoryExportAdvanceOption  `json:"filters"`
            CompletedOn string  `json:"completed_on"`
            NotificationEmails []string  `json:"notification_emails"`
            SellerID float64  `json:"seller_id"`
            TaskID string  `json:"task_id"`
            Type string  `json:"type"`
            Status string  `json:"status"`
         
    }
    
    // InventoryExportRequest used by Catalog
    type InventoryExportRequest struct {

        
            Store []float64  `json:"store"`
            Type string  `json:"type"`
            Brand []float64  `json:"brand"`
         
    }
    
    // InventoryExportResponse used by Catalog
    type InventoryExportResponse struct {

        
            CreatedBy string  `json:"created_by"`
            Filters map[string]interface{}  `json:"filters"`
            CreatedOn string  `json:"created_on"`
            NotificationEmails []string  `json:"notification_emails"`
            SellerID float64  `json:"seller_id"`
            ModifiedOn string  `json:"modified_on"`
            TaskID string  `json:"task_id"`
            Type string  `json:"type"`
            Status string  `json:"status"`
         
    }
    
    // InventoryJobFilters used by Catalog
    type InventoryJobFilters struct {

        
            Stores []string  `json:"stores"`
            Quantity InventoryExportQuantityFilter  `json:"quantity"`
            FromDate string  `json:"from_date"`
            Brands []string  `json:"brands"`
            ToDate string  `json:"to_date"`
         
    }
    
    // InventoryJobDetailResponse used by Catalog
    type InventoryJobDetailResponse struct {

        
            ID string  `json:"id"`
            CreatedBy UserDetail  `json:"created_by"`
            URL string  `json:"url"`
            CancelledBy UserDetail  `json:"cancelled_by"`
            Filters InventoryJobFilters  `json:"filters"`
            CreatedOn string  `json:"created_on"`
            CompletedOn string  `json:"completed_on"`
            NotificationEmails []string  `json:"notification_emails"`
            SellerID float64  `json:"seller_id"`
            ModifiedOn string  `json:"modified_on"`
            TaskID string  `json:"task_id"`
            CancelledOn string  `json:"cancelled_on"`
            Type string  `json:"type"`
            Status interface{}  `json:"status"`
         
    }
    
    // InventoryExportJobListResponse used by Catalog
    type InventoryExportJobListResponse struct {

        
            Items InventoryJobDetailResponse  `json:"items"`
         
    }
    
    // InventoryExportFilter used by Catalog
    type InventoryExportFilter struct {

        
            Quantity InventoryExportQuantityFilter  `json:"quantity"`
            ToDate string  `json:"to_date"`
            FromDate string  `json:"from_date"`
            StoreIds []float64  `json:"store_ids"`
            BrandIds []float64  `json:"brand_ids"`
         
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

        
            StoreID float64  `json:"store_id"`
            TotalQuantity float64  `json:"total_quantity"`
            TraceID string  `json:"trace_id"`
            PriceEffective float64  `json:"price_effective"`
            PriceMarked float64  `json:"price_marked"`
            Tags []string  `json:"tags"`
            SellerIdentifier string  `json:"seller_identifier"`
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

        
            Message string  `json:"message"`
            Items []InventoryResponseItem  `json:"items"`
         
    }
    
    // HsnCodesObject used by Catalog
    type HsnCodesObject struct {

        
            ID string  `json:"id"`
            CompanyID float64  `json:"company_id"`
            Tax1 float64  `json:"tax1"`
            TaxOnMrp bool  `json:"tax_on_mrp"`
            TaxOnEsp bool  `json:"tax_on_esp"`
            HsnCode string  `json:"hsn_code"`
            Tax2 float64  `json:"tax2"`
            ModifiedOn string  `json:"modified_on"`
            Threshold2 float64  `json:"threshold2"`
            Hs2Code string  `json:"hs2_code"`
            Threshold1 float64  `json:"threshold1"`
         
    }
    
    // HsnCode used by Catalog
    type HsnCode struct {

        
            Data HsnCodesObject  `json:"data"`
         
    }
    
    // HsnUpsert used by Catalog
    type HsnUpsert struct {

        
            CompanyID float64  `json:"company_id"`
            UID float64  `json:"uid"`
            Tax1 float64  `json:"tax1"`
            TaxOnMrp bool  `json:"tax_on_mrp"`
            TaxOnEsp bool  `json:"tax_on_esp"`
            HsnCode string  `json:"hsn_code"`
            Tax2 float64  `json:"tax2"`
            IsActive bool  `json:"is_active"`
            Threshold2 float64  `json:"threshold2"`
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

        
            Threshold float64  `json:"threshold"`
            Rate float64  `json:"rate"`
            EffectiveDate string  `json:"effective_date"`
            Cess float64  `json:"cess"`
         
    }
    
    // HSNDataInsertV2 used by Catalog
    type HSNDataInsertV2 struct {

        
            Taxes []TaxSlab  `json:"taxes"`
            HsnCodeID string  `json:"hsn_code_id"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            HsnCode string  `json:"hsn_code"`
            CreatedOn string  `json:"created_on"`
            CountryCode string  `json:"country_code"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            ModifiedOn string  `json:"modified_on"`
            ReportingHsn string  `json:"reporting_hsn"`
            Description string  `json:"description"`
            Type string  `json:"type"`
         
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
    
    // Media2 used by Catalog
    type Media2 struct {

        
            AspectRatio string  `json:"aspect_ratio"`
            Type string  `json:"type"`
            URL string  `json:"url"`
         
    }
    
    // BrandItem used by Catalog
    type BrandItem struct {

        
            UID float64  `json:"uid"`
            Slug string  `json:"slug"`
            Banners ImageUrls  `json:"banners"`
            Action Action  `json:"action"`
            Discount string  `json:"discount"`
            Departments []string  `json:"departments"`
            Name string  `json:"name"`
            Logo Media2  `json:"logo"`
         
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
            Name string  `json:"name"`
            Logo Media2  `json:"logo"`
         
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
            Slug string  `json:"slug"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Banners ImageUrls  `json:"banners"`
            Action Action  `json:"action"`
            Childs []map[string]interface{}  `json:"childs"`
            Name string  `json:"name"`
         
    }
    
    // SecondLevelChild used by Catalog
    type SecondLevelChild struct {

        
            UID float64  `json:"uid"`
            Slug string  `json:"slug"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Banners ImageUrls  `json:"banners"`
            Action Action  `json:"action"`
            Childs []ThirdLevelChild  `json:"childs"`
            Name string  `json:"name"`
         
    }
    
    // Child used by Catalog
    type Child struct {

        
            UID float64  `json:"uid"`
            Slug string  `json:"slug"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Banners ImageUrls  `json:"banners"`
            Action Action  `json:"action"`
            Childs []SecondLevelChild  `json:"childs"`
            Name string  `json:"name"`
         
    }
    
    // CategoryItems used by Catalog
    type CategoryItems struct {

        
            UID float64  `json:"uid"`
            Slug string  `json:"slug"`
            Banners ImageUrls  `json:"banners"`
            Action Action  `json:"action"`
            Childs []Child  `json:"childs"`
            Name string  `json:"name"`
         
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

        
            Operators map[string]interface{}  `json:"operators"`
            SortOn []ProductSortOn  `json:"sort_on"`
            Filters []ProductFilters  `json:"filters"`
            Page Page  `json:"page"`
            Items []ProductListingDetail  `json:"items"`
         
    }
    
    // ProductDetail used by Catalog
    type ProductDetail struct {

        
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            ImageNature string  `json:"image_nature"`
            Name string  `json:"name"`
            Attributes map[string]interface{}  `json:"attributes"`
            Brand ProductBrand  `json:"brand"`
            Rating float64  `json:"rating"`
            Slug string  `json:"slug"`
            Tryouts []string  `json:"tryouts"`
            ProductOnlineDate string  `json:"product_online_date"`
            PromoMeta map[string]interface{}  `json:"promo_meta"`
            ItemCode string  `json:"item_code"`
            Similars []string  `json:"similars"`
            Medias []Media  `json:"medias"`
            RatingCount float64  `json:"rating_count"`
            Color string  `json:"color"`
            ItemType string  `json:"item_type"`
            Description string  `json:"description"`
            ShortDescription string  `json:"short_description"`
            UID float64  `json:"uid"`
            Highlights []string  `json:"highlights"`
            HasVariant bool  `json:"has_variant"`
            Type string  `json:"type"`
         
    }
    
    // InventoryPage used by Catalog
    type InventoryPage struct {

        
            ItemTotal float64  `json:"item_total"`
            HasPrevious bool  `json:"has_previous"`
            NextID string  `json:"next_id"`
            Type string  `json:"type"`
            HasNext bool  `json:"has_next"`
         
    }
    
    // InventoryStockResponse used by Catalog
    type InventoryStockResponse struct {

        
            Items []map[string]interface{}  `json:"items"`
            Page InventoryPage  `json:"page"`
         
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

        
            GroupID string  `json:"group_id"`
            Quantity float64  `json:"quantity"`
            ArticleAssignment ArticleAssignment  `json:"article_assignment"`
            Query ArticleQuery  `json:"query"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // AssignStore used by Catalog
    type AssignStore struct {

        
            CompanyID float64  `json:"company_id"`
            Pincode string  `json:"pincode"`
            ChannelType string  `json:"channel_type"`
            Articles []AssignStoreArticle  `json:"articles"`
            StoreIds []float64  `json:"store_ids"`
            AppID string  `json:"app_id"`
            ChannelIdentifier string  `json:"channel_identifier"`
         
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
            UID string  `json:"uid"`
            Size string  `json:"size"`
            StoreID float64  `json:"store_id"`
            SCity string  `json:"s_city"`
            StrategyWiseListing []map[string]interface{}  `json:"strategy_wise_listing"`
            GroupID string  `json:"group_id"`
            Quantity float64  `json:"quantity"`
            ArticleAssignment ArticleAssignment1  `json:"article_assignment"`
            PriceEffective float64  `json:"price_effective"`
            Meta map[string]interface{}  `json:"meta"`
            Index float64  `json:"index"`
            ItemID float64  `json:"item_id"`
            PriceMarked float64  `json:"price_marked"`
            ID string  `json:"_id"`
            Status bool  `json:"status"`
         
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
    
    // UserSerializer1 used by Catalog
    type UserSerializer1 struct {

        
            Contact string  `json:"contact"`
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
         
    }
    
    // ProductReturnConfigSerializer used by Catalog
    type ProductReturnConfigSerializer struct {

        
            StoreUID float64  `json:"store_uid"`
            OnSameStore bool  `json:"on_same_store"`
         
    }
    
    // UserSerializer2 used by Catalog
    type UserSerializer2 struct {

        
            Contact string  `json:"contact"`
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
         
    }
    
    // GetAddressSerializer used by Catalog
    type GetAddressSerializer struct {

        
            Address2 string  `json:"address2"`
            Pincode float64  `json:"pincode"`
            City string  `json:"city"`
            AddressType string  `json:"address_type"`
            Longitude float64  `json:"longitude"`
            Landmark string  `json:"landmark"`
            State string  `json:"state"`
            Latitude float64  `json:"latitude"`
            CountryCode string  `json:"country_code"`
            Address1 string  `json:"address1"`
            Country string  `json:"country"`
         
    }
    
    // GetCompanySerializer used by Catalog
    type GetCompanySerializer struct {

        
            UID float64  `json:"uid"`
            CreatedBy UserSerializer2  `json:"created_by"`
            CompanyType string  `json:"company_type"`
            RejectReason string  `json:"reject_reason"`
            BusinessType string  `json:"business_type"`
            CreatedOn string  `json:"created_on"`
            ModifiedBy UserSerializer2  `json:"modified_by"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            ModifiedOn string  `json:"modified_on"`
            VerifiedBy UserSerializer2  `json:"verified_by"`
            VerifiedOn string  `json:"verified_on"`
            Name string  `json:"name"`
            Stage string  `json:"stage"`
         
    }
    
    // Document used by Catalog
    type Document struct {

        
            Verified bool  `json:"verified"`
            URL string  `json:"url"`
            LegalName string  `json:"legal_name"`
            Value string  `json:"value"`
            Type string  `json:"type"`
         
    }
    
    // LocationIntegrationType used by Catalog
    type LocationIntegrationType struct {

        
            Order string  `json:"order"`
            Inventory string  `json:"inventory"`
         
    }
    
    // SellerPhoneNumber used by Catalog
    type SellerPhoneNumber struct {

        
            CountryCode float64  `json:"country_code"`
            Number string  `json:"number"`
         
    }
    
    // LocationTimingSerializer used by Catalog
    type LocationTimingSerializer struct {

        
            Hour float64  `json:"hour"`
            Minute float64  `json:"minute"`
         
    }
    
    // LocationDayWiseSerializer used by Catalog
    type LocationDayWiseSerializer struct {

        
            Open bool  `json:"open"`
            Weekday string  `json:"weekday"`
            Opening LocationTimingSerializer  `json:"opening"`
            Closing LocationTimingSerializer  `json:"closing"`
         
    }
    
    // LocationManagerSerializer used by Catalog
    type LocationManagerSerializer struct {

        
            Name string  `json:"name"`
            MobileNo SellerPhoneNumber  `json:"mobile_no"`
            Email string  `json:"email"`
         
    }
    
    // GetLocationSerializer used by Catalog
    type GetLocationSerializer struct {

        
            GstCredentials InvoiceDetailsSerializer  `json:"gst_credentials"`
            CreatedBy UserSerializer1  `json:"created_by"`
            StoreType string  `json:"store_type"`
            ProductReturnConfig ProductReturnConfigSerializer  `json:"product_return_config"`
            Name string  `json:"name"`
            Company GetCompanySerializer  `json:"company"`
            Code string  `json:"code"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            CreatedOn string  `json:"created_on"`
            ModifiedBy UserSerializer1  `json:"modified_by"`
            PhoneNumber string  `json:"phone_number"`
            ModifiedOn string  `json:"modified_on"`
            Documents []Document  `json:"documents"`
            Stage string  `json:"stage"`
            IntegrationType LocationIntegrationType  `json:"integration_type"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            Address GetAddressSerializer  `json:"address"`
            NotificationEmails []string  `json:"notification_emails"`
            VerifiedOn string  `json:"verified_on"`
            DisplayName string  `json:"display_name"`
            UID float64  `json:"uid"`
            Warnings map[string]interface{}  `json:"warnings"`
            VerifiedBy UserSerializer1  `json:"verified_by"`
            Timing []LocationDayWiseSerializer  `json:"timing"`
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

        
            UID float64  `json:"uid"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            Logo string  `json:"logo"`
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

        
            UserID string  `json:"user_id"`
            Contact string  `json:"contact"`
            Username string  `json:"username"`
         
    }
    
    // BusinessCountryInfo used by CompanyProfile
    type BusinessCountryInfo struct {

        
            CountryCode string  `json:"country_code"`
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
            Value string  `json:"value"`
            Type string  `json:"type"`
            Verified bool  `json:"verified"`
            URL string  `json:"url"`
         
    }
    
    // GetAddressSerializer used by CompanyProfile
    type GetAddressSerializer struct {

        
            CountryCode string  `json:"country_code"`
            Landmark string  `json:"landmark"`
            Country string  `json:"country"`
            AddressType string  `json:"address_type"`
            Address2 string  `json:"address2"`
            Pincode float64  `json:"pincode"`
            Address1 string  `json:"address1"`
            State string  `json:"state"`
            Longitude float64  `json:"longitude"`
            Latitude float64  `json:"latitude"`
            City string  `json:"city"`
         
    }
    
    // GetCompanyProfileSerializerResponse used by CompanyProfile
    type GetCompanyProfileSerializerResponse struct {

        
            Stage string  `json:"stage"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            NotificationEmails []string  `json:"notification_emails"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            BusinessCountryInfo BusinessCountryInfo  `json:"business_country_info"`
            ContactDetails ContactDetails  `json:"contact_details"`
            Taxes []CompanyTaxesSerializer  `json:"taxes"`
            CreatedOn string  `json:"created_on"`
            Warnings map[string]interface{}  `json:"warnings"`
            CompanyType string  `json:"company_type"`
            ModifiedOn string  `json:"modified_on"`
            Name string  `json:"name"`
            VerifiedOn string  `json:"verified_on"`
            BusinessType string  `json:"business_type"`
            CreatedBy UserSerializer  `json:"created_by"`
            BusinessInfo string  `json:"business_info"`
            Mode string  `json:"mode"`
            BusinessDetails BusinessDetails  `json:"business_details"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            Documents []Document  `json:"documents"`
            FranchiseEnabled bool  `json:"franchise_enabled"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            UID float64  `json:"uid"`
         
    }
    
    // ErrorResponse used by CompanyProfile
    type ErrorResponse struct {

        
            Status float64  `json:"status"`
            Meta map[string]interface{}  `json:"meta"`
            Message string  `json:"message"`
            Code string  `json:"code"`
         
    }
    
    // CompanyTaxesSerializer1 used by CompanyProfile
    type CompanyTaxesSerializer1 struct {

        
            EffectiveDate string  `json:"effective_date"`
            Rate float64  `json:"rate"`
            Enable bool  `json:"enable"`
         
    }
    
    // CreateUpdateAddressSerializer used by CompanyProfile
    type CreateUpdateAddressSerializer struct {

        
            CountryCode string  `json:"country_code"`
            Landmark string  `json:"landmark"`
            Country string  `json:"country"`
            Pincode float64  `json:"pincode"`
            Address2 string  `json:"address2"`
            AddressType string  `json:"address_type"`
            Address1 string  `json:"address1"`
            State string  `json:"state"`
            Longitude float64  `json:"longitude"`
            Latitude float64  `json:"latitude"`
            City string  `json:"city"`
         
    }
    
    // UpdateCompany used by CompanyProfile
    type UpdateCompany struct {

        
            Taxes []CompanyTaxesSerializer1  `json:"taxes"`
            ContactDetails ContactDetails  `json:"contact_details"`
            RejectReason string  `json:"reject_reason"`
            BusinessDetails BusinessDetails  `json:"business_details"`
            BusinessType string  `json:"business_type"`
            Warnings map[string]interface{}  `json:"warnings"`
            CompanyType string  `json:"company_type"`
            NotificationEmails []string  `json:"notification_emails"`
            FranchiseEnabled bool  `json:"franchise_enabled"`
            Documents []Document  `json:"documents"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Addresses []CreateUpdateAddressSerializer  `json:"addresses"`
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

        
            Stage string  `json:"stage"`
            CompanyDocuments DocumentsObj  `json:"company_documents"`
            Product DocumentsObj  `json:"product"`
            Store DocumentsObj  `json:"store"`
            Brand DocumentsObj  `json:"brand"`
            UID float64  `json:"uid"`
            StoreDocuments DocumentsObj  `json:"store_documents"`
         
    }
    
    // BrandBannerSerializer used by CompanyProfile
    type BrandBannerSerializer struct {

        
            Landscape string  `json:"landscape"`
            Portrait string  `json:"portrait"`
         
    }
    
    // GetBrandResponseSerializer used by CompanyProfile
    type GetBrandResponseSerializer struct {

        
            Synonyms []string  `json:"synonyms"`
            Stage string  `json:"stage"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            SlugKey string  `json:"slug_key"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            Logo string  `json:"logo"`
            CreatedOn string  `json:"created_on"`
            Warnings map[string]interface{}  `json:"warnings"`
            ModifiedOn string  `json:"modified_on"`
            Name string  `json:"name"`
            VerifiedOn string  `json:"verified_on"`
            RejectReason string  `json:"reject_reason"`
            Description string  `json:"description"`
            CreatedBy UserSerializer  `json:"created_by"`
            Mode string  `json:"mode"`
            Banner BrandBannerSerializer  `json:"banner"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            UID float64  `json:"uid"`
         
    }
    
    // CreateUpdateBrandRequestSerializer used by CompanyProfile
    type CreateUpdateBrandRequestSerializer struct {

        
            Synonyms []string  `json:"synonyms"`
            Banner BrandBannerSerializer  `json:"banner"`
            Description string  `json:"description"`
            BrandTier string  `json:"brand_tier"`
            CompanyID float64  `json:"company_id"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Name string  `json:"name"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            UID float64  `json:"uid"`
            Logo string  `json:"logo"`
         
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

        
            BusinessCountryInfo BusinessCountryInfo  `json:"business_country_info"`
            RejectReason string  `json:"reject_reason"`
            Stage string  `json:"stage"`
            CreatedBy UserSerializer  `json:"created_by"`
            BusinessType string  `json:"business_type"`
            CreatedOn string  `json:"created_on"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            CompanyType string  `json:"company_type"`
            NotificationEmails []string  `json:"notification_emails"`
            ModifiedOn string  `json:"modified_on"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            MarketChannels []string  `json:"market_channels"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            Name string  `json:"name"`
            UID float64  `json:"uid"`
            Details CompanyDetails  `json:"details"`
            VerifiedOn string  `json:"verified_on"`
         
    }
    
    // CompanyBrandSerializer used by CompanyProfile
    type CompanyBrandSerializer struct {

        
            RejectReason string  `json:"reject_reason"`
            Stage string  `json:"stage"`
            CreatedBy UserSerializer  `json:"created_by"`
            CreatedOn string  `json:"created_on"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            Warnings map[string]interface{}  `json:"warnings"`
            ModifiedOn string  `json:"modified_on"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            Brand GetBrandResponseSerializer  `json:"brand"`
            UID float64  `json:"uid"`
            VerifiedOn string  `json:"verified_on"`
            Company CompanySerializer  `json:"company"`
         
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
    
    // LocationTimingSerializer used by CompanyProfile
    type LocationTimingSerializer struct {

        
            Hour float64  `json:"hour"`
            Minute float64  `json:"minute"`
         
    }
    
    // LocationDayWiseSerializer used by CompanyProfile
    type LocationDayWiseSerializer struct {

        
            Closing LocationTimingSerializer  `json:"closing"`
            Opening LocationTimingSerializer  `json:"opening"`
            Weekday string  `json:"weekday"`
            Open bool  `json:"open"`
         
    }
    
    // ProductReturnConfigSerializer used by CompanyProfile
    type ProductReturnConfigSerializer struct {

        
            StoreUID float64  `json:"store_uid"`
            OnSameStore bool  `json:"on_same_store"`
         
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
            Name string  `json:"name"`
            MobileNo SellerPhoneNumber  `json:"mobile_no"`
         
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
    
    // GetCompanySerializer used by CompanyProfile
    type GetCompanySerializer struct {

        
            RejectReason string  `json:"reject_reason"`
            Stage string  `json:"stage"`
            CreatedBy UserSerializer  `json:"created_by"`
            BusinessType string  `json:"business_type"`
            CreatedOn string  `json:"created_on"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            CompanyType string  `json:"company_type"`
            ModifiedOn string  `json:"modified_on"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            Name string  `json:"name"`
            UID float64  `json:"uid"`
            VerifiedOn string  `json:"verified_on"`
         
    }
    
    // GetLocationSerializer used by CompanyProfile
    type GetLocationSerializer struct {

        
            Stage string  `json:"stage"`
            Timing []LocationDayWiseSerializer  `json:"timing"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            NotificationEmails []string  `json:"notification_emails"`
            Address GetAddressSerializer  `json:"address"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            CreatedOn string  `json:"created_on"`
            Warnings map[string]interface{}  `json:"warnings"`
            CreditNote bool  `json:"credit_note"`
            ModifiedOn string  `json:"modified_on"`
            AutoInvoice bool  `json:"auto_invoice"`
            Name string  `json:"name"`
            VerifiedOn string  `json:"verified_on"`
            PhoneNumber string  `json:"phone_number"`
            Code string  `json:"code"`
            ProductReturnConfig ProductReturnConfigSerializer  `json:"product_return_config"`
            Holiday []HolidaySchemaSerializer  `json:"holiday"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            CreatedBy UserSerializer  `json:"created_by"`
            DisplayName string  `json:"display_name"`
            StoreType string  `json:"store_type"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            Manager LocationManagerSerializer  `json:"manager"`
            Documents []Document  `json:"documents"`
            UID float64  `json:"uid"`
            GstCredentials InvoiceDetailsSerializer  `json:"gst_credentials"`
            Company GetCompanySerializer  `json:"company"`
         
    }
    
    // LocationListSerializer used by CompanyProfile
    type LocationListSerializer struct {

        
            Items []GetLocationSerializer  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // AddressSerializer used by CompanyProfile
    type AddressSerializer struct {

        
            CountryCode string  `json:"country_code"`
            Landmark string  `json:"landmark"`
            Country string  `json:"country"`
            AddressType string  `json:"address_type"`
            Address2 string  `json:"address2"`
            Pincode float64  `json:"pincode"`
            Address1 string  `json:"address1"`
            State string  `json:"state"`
            Longitude float64  `json:"longitude"`
            Latitude float64  `json:"latitude"`
            City string  `json:"city"`
         
    }
    
    // LocationSerializer used by CompanyProfile
    type LocationSerializer struct {

        
            Stage string  `json:"stage"`
            Timing []LocationDayWiseSerializer  `json:"timing"`
            NotificationEmails []string  `json:"notification_emails"`
            Address AddressSerializer  `json:"address"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Warnings map[string]interface{}  `json:"warnings"`
            CreditNote bool  `json:"credit_note"`
            AutoInvoice bool  `json:"auto_invoice"`
            Name string  `json:"name"`
            Code string  `json:"code"`
            ProductReturnConfig ProductReturnConfigSerializer  `json:"product_return_config"`
            Holiday []HolidaySchemaSerializer  `json:"holiday"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            DisplayName string  `json:"display_name"`
            StoreType string  `json:"store_type"`
            Documents []Document  `json:"documents"`
            Manager LocationManagerSerializer  `json:"manager"`
            UID float64  `json:"uid"`
            GstCredentials InvoiceDetailsSerializer  `json:"gst_credentials"`
            Company float64  `json:"company"`
         
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
    

    
    // CouponDateMeta used by Cart
    type CouponDateMeta struct {

        
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // Identifier used by Cart
    type Identifier struct {

        
            ItemID []float64  `json:"item_id"`
            CollectionID []string  `json:"collection_id"`
            ExcludeBrandID []float64  `json:"exclude_brand_id"`
            ArticleID []string  `json:"article_id"`
            BrandID []float64  `json:"brand_id"`
            UserID []string  `json:"user_id"`
            EmailDomain []string  `json:"email_domain"`
            CategoryID []float64  `json:"category_id"`
            CompanyID []float64  `json:"company_id"`
            StoreID []float64  `json:"store_id"`
         
    }
    
    // DisplayMetaDict used by Cart
    type DisplayMetaDict struct {

        
            Title string  `json:"title"`
            Subtitle string  `json:"subtitle"`
         
    }
    
    // DisplayMeta used by Cart
    type DisplayMeta struct {

        
            Description string  `json:"description"`
            Title string  `json:"title"`
            Subtitle string  `json:"subtitle"`
            Auto DisplayMetaDict  `json:"auto"`
            Remove DisplayMetaDict  `json:"remove"`
            Apply DisplayMetaDict  `json:"apply"`
         
    }
    
    // Ownership used by Cart
    type Ownership struct {

        
            PayableCategory string  `json:"payable_category"`
            PayableBy string  `json:"payable_by"`
         
    }
    
    // RuleDefinition used by Cart
    type RuleDefinition struct {

        
            ApplicableOn string  `json:"applicable_on"`
            IsExact bool  `json:"is_exact"`
            AutoApply bool  `json:"auto_apply"`
            Scope []string  `json:"scope"`
            CurrencyCode string  `json:"currency_code"`
            Type string  `json:"type"`
            ValueType string  `json:"value_type"`
            CalculateOn string  `json:"calculate_on"`
         
    }
    
    // Rule used by Cart
    type Rule struct {

        
            Max float64  `json:"max"`
            Min float64  `json:"min"`
            DiscountQty float64  `json:"discount_qty"`
            Key float64  `json:"key"`
            Value float64  `json:"value"`
         
    }
    
    // Validation used by Cart
    type Validation struct {

        
            UserRegisteredAfter string  `json:"user_registered_after"`
            Anonymous bool  `json:"anonymous"`
            AppID []string  `json:"app_id"`
         
    }
    
    // CouponAuthor used by Cart
    type CouponAuthor struct {

        
            CreatedBy string  `json:"created_by"`
            ModifiedBy string  `json:"modified_by"`
         
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
    
    // PriceRange used by Cart
    type PriceRange struct {

        
            Max float64  `json:"max"`
            Min float64  `json:"min"`
         
    }
    
    // PostOrder used by Cart
    type PostOrder struct {

        
            ReturnAllowed bool  `json:"return_allowed"`
            CancellationAllowed bool  `json:"cancellation_allowed"`
         
    }
    
    // PaymentAllowValue used by Cart
    type PaymentAllowValue struct {

        
            Max float64  `json:"max"`
         
    }
    
    // PaymentModes used by Cart
    type PaymentModes struct {

        
            Uses PaymentAllowValue  `json:"uses"`
            Types []string  `json:"types"`
            Networks []string  `json:"networks"`
            Iins []string  `json:"iins"`
            Codes []string  `json:"codes"`
         
    }
    
    // BulkBundleRestriction used by Cart
    type BulkBundleRestriction struct {

        
            MultiStoreAllowed bool  `json:"multi_store_allowed"`
         
    }
    
    // Restrictions used by Cart
    type Restrictions struct {

        
            UserGroups []float64  `json:"user_groups"`
            Uses UsesRestriction  `json:"uses"`
            Platforms []string  `json:"platforms"`
            PriceRange PriceRange  `json:"price_range"`
            OrderingStores []float64  `json:"ordering_stores"`
            PostOrder PostOrder  `json:"post_order"`
            CouponAllowed bool  `json:"coupon_allowed"`
            UserType string  `json:"user_type"`
            Payments map[string]PaymentModes  `json:"payments"`
            BulkBundle BulkBundleRestriction  `json:"bulk_bundle"`
         
    }
    
    // CouponAction used by Cart
    type CouponAction struct {

        
            TxnMode string  `json:"txn_mode"`
            ActionDate string  `json:"action_date"`
         
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
    
    // CouponSchedule used by Cart
    type CouponSchedule struct {

        
            Start string  `json:"start"`
            NextSchedule []map[string]interface{}  `json:"next_schedule"`
            Duration float64  `json:"duration"`
            Cron string  `json:"cron"`
            End string  `json:"end"`
         
    }
    
    // CouponAdd used by Cart
    type CouponAdd struct {

        
            DateMeta CouponDateMeta  `json:"date_meta"`
            Identifiers Identifier  `json:"identifiers"`
            Code string  `json:"code"`
            DisplayMeta DisplayMeta  `json:"display_meta"`
            Ownership Ownership  `json:"ownership"`
            RuleDefinition RuleDefinition  `json:"rule_definition"`
            Rule []Rule  `json:"rule"`
            Validation Validation  `json:"validation"`
            Tags []string  `json:"tags"`
            TypeSlug string  `json:"type_slug"`
            Author CouponAuthor  `json:"author"`
            Restrictions Restrictions  `json:"restrictions"`
            Action CouponAction  `json:"action"`
            Validity Validity  `json:"validity"`
            State State  `json:"state"`
            Schedule CouponSchedule  `json:"_schedule"`
         
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

        
            DateMeta CouponDateMeta  `json:"date_meta"`
            Identifiers Identifier  `json:"identifiers"`
            Code string  `json:"code"`
            DisplayMeta DisplayMeta  `json:"display_meta"`
            Ownership Ownership  `json:"ownership"`
            RuleDefinition RuleDefinition  `json:"rule_definition"`
            Rule []Rule  `json:"rule"`
            Validation Validation  `json:"validation"`
            Tags []string  `json:"tags"`
            TypeSlug string  `json:"type_slug"`
            Author CouponAuthor  `json:"author"`
            Restrictions Restrictions  `json:"restrictions"`
            Action CouponAction  `json:"action"`
            Validity Validity  `json:"validity"`
            State State  `json:"state"`
            Schedule CouponSchedule  `json:"_schedule"`
         
    }
    
    // CouponPartialUpdate used by Cart
    type CouponPartialUpdate struct {

        
            Schedule CouponSchedule  `json:"schedule"`
            Archive bool  `json:"archive"`
         
    }
    
    // DisplayMeta1 used by Cart
    type DisplayMeta1 struct {

        
            Name string  `json:"name"`
            Description string  `json:"description"`
            OfferText string  `json:"offer_text"`
            OfferLabel string  `json:"offer_label"`
         
    }
    
    // CompareObject used by Cart
    type CompareObject struct {

        
            LessThanEquals float64  `json:"less_than_equals"`
            Equals float64  `json:"equals"`
            GreaterThan float64  `json:"greater_than"`
            GreaterThanEquals float64  `json:"greater_than_equals"`
            LessThan float64  `json:"less_than"`
         
    }
    
    // ItemCriteria used by Cart
    type ItemCriteria struct {

        
            ItemExcludeStore []float64  `json:"item_exclude_store"`
            ItemExcludeSku []string  `json:"item_exclude_sku"`
            ItemExcludeCategory []float64  `json:"item_exclude_category"`
            AvailableZones []string  `json:"available_zones"`
            CartUniqueItemQuantity CompareObject  `json:"cart_unique_item_quantity"`
            ItemCompany []float64  `json:"item_company"`
            ItemSize []string  `json:"item_size"`
            ItemExcludeL2Category []float64  `json:"item_exclude_l2_category"`
            ItemID []float64  `json:"item_id"`
            ItemDepartment []float64  `json:"item_department"`
            BuyRules []string  `json:"buy_rules"`
            CartQuantity CompareObject  `json:"cart_quantity"`
            ItemExcludeBrand []float64  `json:"item_exclude_brand"`
            ItemL1Category []float64  `json:"item_l1_category"`
            ItemSku []string  `json:"item_sku"`
            ItemTags []string  `json:"item_tags"`
            ItemStore []float64  `json:"item_store"`
            CartTotal CompareObject  `json:"cart_total"`
            ItemBrand []float64  `json:"item_brand"`
            ItemExcludeCompany []float64  `json:"item_exclude_company"`
            AllItems bool  `json:"all_items"`
            ItemCategory []float64  `json:"item_category"`
            ItemExcludeDepartment []float64  `json:"item_exclude_department"`
            ItemL2Category []float64  `json:"item_l2_category"`
            ItemExcludeID []float64  `json:"item_exclude_id"`
            ItemExcludeL1Category []float64  `json:"item_exclude_l1_category"`
            CartUniqueItemAmount CompareObject  `json:"cart_unique_item_amount"`
         
    }
    
    // DiscountOffer used by Cart
    type DiscountOffer struct {

        
            PartialCanRet bool  `json:"partial_can_ret"`
            MinOfferQuantity float64  `json:"min_offer_quantity"`
            MaxUsagePerTransaction float64  `json:"max_usage_per_transaction"`
            Code string  `json:"code"`
            MaxDiscountAmount float64  `json:"max_discount_amount"`
            MaxOfferQuantity float64  `json:"max_offer_quantity"`
            DiscountAmount float64  `json:"discount_amount"`
            DiscountPrice float64  `json:"discount_price"`
            DiscountPercentage float64  `json:"discount_percentage"`
            ApportionDiscount bool  `json:"apportion_discount"`
         
    }
    
    // DiscountRule used by Cart
    type DiscountRule struct {

        
            ItemCriteria ItemCriteria  `json:"item_criteria"`
            Offer DiscountOffer  `json:"offer"`
            BuyCondition string  `json:"buy_condition"`
            DiscountType string  `json:"discount_type"`
         
    }
    
    // PromotionDateMeta used by Cart
    type PromotionDateMeta struct {

        
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
         
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
    
    // PromotionAction used by Cart
    type PromotionAction struct {

        
            ActionType string  `json:"action_type"`
            ActionDate string  `json:"action_date"`
         
    }
    
    // PromotionAuthor used by Cart
    type PromotionAuthor struct {

        
            CreatedBy string  `json:"created_by"`
            ModifiedBy string  `json:"modified_by"`
         
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

        
            ReturnAllowed bool  `json:"return_allowed"`
            CancellationAllowed bool  `json:"cancellation_allowed"`
         
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

        
            Uses PaymentAllowValue1  `json:"uses"`
            Type string  `json:"type"`
            Codes []string  `json:"codes"`
         
    }
    
    // Restrictions1 used by Cart
    type Restrictions1 struct {

        
            UserGroups []float64  `json:"user_groups"`
            Uses UsesRestriction1  `json:"uses"`
            Platforms []string  `json:"platforms"`
            OrderingStores []float64  `json:"ordering_stores"`
            OrderQuantity float64  `json:"order_quantity"`
            PostOrder PostOrder1  `json:"post_order"`
            UserID []string  `json:"user_id"`
            UserRegistered UserRegistered  `json:"user_registered"`
            AnonymousUsers bool  `json:"anonymous_users"`
            Payments []PromotionPaymentModes  `json:"payments"`
         
    }
    
    // PromotionSchedule used by Cart
    type PromotionSchedule struct {

        
            Start string  `json:"start"`
            Published bool  `json:"published"`
            NextSchedule []map[string]interface{}  `json:"next_schedule"`
            Duration float64  `json:"duration"`
            Cron string  `json:"cron"`
            End string  `json:"end"`
         
    }
    
    // PromotionListItem used by Cart
    type PromotionListItem struct {

        
            DisplayMeta DisplayMeta1  `json:"display_meta"`
            ApplyPriority float64  `json:"apply_priority"`
            Mode string  `json:"mode"`
            PromoGroup string  `json:"promo_group"`
            BuyRules map[string]ItemCriteria  `json:"buy_rules"`
            ApplyAllDiscount bool  `json:"apply_all_discount"`
            Currency string  `json:"currency"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            PromotionType string  `json:"promotion_type"`
            CalculateOn string  `json:"calculate_on"`
            DiscountRules []DiscountRule  `json:"discount_rules"`
            Code string  `json:"code"`
            DateMeta PromotionDateMeta  `json:"date_meta"`
            ApplicationID string  `json:"application_id"`
            Stackable bool  `json:"stackable"`
            Ownership Ownership1  `json:"ownership"`
            Visiblility Visibility  `json:"visiblility"`
            PostOrderAction PromotionAction  `json:"post_order_action"`
            Author PromotionAuthor  `json:"author"`
            Restrictions Restrictions1  `json:"restrictions"`
            ApplyExclusive string  `json:"apply_exclusive"`
            Schedule PromotionSchedule  `json:"_schedule"`
         
    }
    
    // PromotionsResponse used by Cart
    type PromotionsResponse struct {

        
            Items PromotionListItem  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // PromotionAdd used by Cart
    type PromotionAdd struct {

        
            DisplayMeta DisplayMeta1  `json:"display_meta"`
            ApplyPriority float64  `json:"apply_priority"`
            Mode string  `json:"mode"`
            PromoGroup string  `json:"promo_group"`
            BuyRules map[string]ItemCriteria  `json:"buy_rules"`
            ApplyAllDiscount bool  `json:"apply_all_discount"`
            Currency string  `json:"currency"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            PromotionType string  `json:"promotion_type"`
            CalculateOn string  `json:"calculate_on"`
            DiscountRules []DiscountRule  `json:"discount_rules"`
            Code string  `json:"code"`
            DateMeta PromotionDateMeta  `json:"date_meta"`
            ApplicationID string  `json:"application_id"`
            Stackable bool  `json:"stackable"`
            Ownership Ownership1  `json:"ownership"`
            Visiblility Visibility  `json:"visiblility"`
            PostOrderAction PromotionAction  `json:"post_order_action"`
            Author PromotionAuthor  `json:"author"`
            Restrictions Restrictions1  `json:"restrictions"`
            ApplyExclusive string  `json:"apply_exclusive"`
            Schedule PromotionSchedule  `json:"_schedule"`
         
    }
    
    // PromotionUpdate used by Cart
    type PromotionUpdate struct {

        
            DisplayMeta DisplayMeta1  `json:"display_meta"`
            ApplyPriority float64  `json:"apply_priority"`
            Mode string  `json:"mode"`
            PromoGroup string  `json:"promo_group"`
            BuyRules map[string]ItemCriteria  `json:"buy_rules"`
            ApplyAllDiscount bool  `json:"apply_all_discount"`
            Currency string  `json:"currency"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            PromotionType string  `json:"promotion_type"`
            CalculateOn string  `json:"calculate_on"`
            DiscountRules []DiscountRule  `json:"discount_rules"`
            Code string  `json:"code"`
            DateMeta PromotionDateMeta  `json:"date_meta"`
            ApplicationID string  `json:"application_id"`
            Stackable bool  `json:"stackable"`
            Ownership Ownership1  `json:"ownership"`
            Visiblility Visibility  `json:"visiblility"`
            PostOrderAction PromotionAction  `json:"post_order_action"`
            Author PromotionAuthor  `json:"author"`
            Restrictions Restrictions1  `json:"restrictions"`
            ApplyExclusive string  `json:"apply_exclusive"`
            Schedule PromotionSchedule  `json:"_schedule"`
         
    }
    
    // PromotionPartialUpdate used by Cart
    type PromotionPartialUpdate struct {

        
            Schedule PromotionSchedule  `json:"schedule"`
            Archive bool  `json:"archive"`
         
    }
    
    // ActivePromosResponse used by Cart
    type ActivePromosResponse struct {

        
            Description string  `json:"description"`
            IsHidden bool  `json:"is_hidden"`
            EntitySlug string  `json:"entity_slug"`
            CreatedOn string  `json:"created_on"`
            Title string  `json:"title"`
            EntityType string  `json:"entity_type"`
            Subtitle string  `json:"subtitle"`
            Example string  `json:"example"`
            Type string  `json:"type"`
            ModifiedOn string  `json:"modified_on"`
         
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

        
            MinCartValue float64  `json:"min_cart_value"`
            Enabled bool  `json:"enabled"`
            MaxCartItems float64  `json:"max_cart_items"`
            GiftPricing float64  `json:"gift_pricing"`
            DeliveryCharges DeliveryCharges  `json:"delivery_charges"`
            GiftDisplayText string  `json:"gift_display_text"`
            RevenueEngineCoupon bool  `json:"revenue_engine_coupon"`
            BulkCoupons bool  `json:"bulk_coupons"`
         
    }
    
    // CartMetaConfigAdd used by Cart
    type CartMetaConfigAdd struct {

        
            MinCartValue float64  `json:"min_cart_value"`
            Enabled bool  `json:"enabled"`
            MaxCartItems float64  `json:"max_cart_items"`
            GiftPricing float64  `json:"gift_pricing"`
            DeliveryCharges DeliveryCharges  `json:"delivery_charges"`
            GiftDisplayText string  `json:"gift_display_text"`
            RevenueEngineCoupon bool  `json:"revenue_engine_coupon"`
            BulkCoupons bool  `json:"bulk_coupons"`
         
    }
    
    // Article used by Cart
    type Article struct {

        
            Code string  `json:"code"`
            ArticleID string  `json:"article_id"`
            Meta map[string]interface{}  `json:"meta"`
            Type string  `json:"type"`
            Value float64  `json:"value"`
         
    }
    
    // Collecttion used by Cart
    type Collecttion struct {

        
            CollectedBy string  `json:"collected_by"`
            RefundBy string  `json:"refund_by"`
         
    }
    
    // CartDynamicInjectionUpdate used by Cart
    type CartDynamicInjectionUpdate struct {

        
            ApplyExpiry string  `json:"apply_expiry"`
            ArticleIds []Article  `json:"article_ids"`
            Collection Collecttion  `json:"collection"`
            IsAuthenticated bool  `json:"is_authenticated"`
            CartValue float64  `json:"cart_value"`
            Meta map[string]interface{}  `json:"meta"`
            AllowedRefund bool  `json:"allowed_refund"`
            CartID string  `json:"cart_id"`
            Type string  `json:"type"`
            UserID string  `json:"user_id"`
            Value float64  `json:"value"`
            ArticleLevelDistribution bool  `json:"article_level_distribution"`
            Message string  `json:"message"`
         
    }
    
    // CartDynamicInjection used by Cart
    type CartDynamicInjection struct {

        
            ApplyExpiry string  `json:"apply_expiry"`
            ArticleIds []Article  `json:"article_ids"`
            Collection Collecttion  `json:"collection"`
            IsAuthenticated bool  `json:"is_authenticated"`
            InjectionID string  `json:"injection_id"`
            CartValue float64  `json:"cart_value"`
            Meta map[string]interface{}  `json:"meta"`
            AllowedRefund bool  `json:"allowed_refund"`
            CartID string  `json:"cart_id"`
            Type string  `json:"type"`
            UserID string  `json:"user_id"`
            Value float64  `json:"value"`
            ArticleLevelDistribution bool  `json:"article_level_distribution"`
            Message string  `json:"message"`
         
    }
    
    // CartDynamicInjectionResponse used by Cart
    type CartDynamicInjectionResponse struct {

        
            Data CartDynamicInjection  `json:"data"`
         
    }
    
    // CartDynamicInjectionAdd used by Cart
    type CartDynamicInjectionAdd struct {

        
            ApplyExpiry string  `json:"apply_expiry"`
            ArticleIds []Article  `json:"article_ids"`
            Collection Collecttion  `json:"collection"`
            IsAuthenticated bool  `json:"is_authenticated"`
            CartValue float64  `json:"cart_value"`
            Meta map[string]interface{}  `json:"meta"`
            AllowedRefund bool  `json:"allowed_refund"`
            CartID string  `json:"cart_id"`
            Type string  `json:"type"`
            UserID string  `json:"user_id"`
            Value float64  `json:"value"`
            ArticleLevelDistribution bool  `json:"article_level_distribution"`
            Message string  `json:"message"`
         
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

        
            DeliveryCharge float64  `json:"delivery_charge"`
            FyndCash float64  `json:"fynd_cash"`
            Discount float64  `json:"discount"`
            ConvenienceFee float64  `json:"convenience_fee"`
            GstCharges float64  `json:"gst_charges"`
            MrpTotal float64  `json:"mrp_total"`
            Vog float64  `json:"vog"`
            CodCharge float64  `json:"cod_charge"`
            Coupon float64  `json:"coupon"`
            Total float64  `json:"total"`
            YouSaved float64  `json:"you_saved"`
            Subtotal float64  `json:"subtotal"`
            GiftCard float64  `json:"gift_card"`
         
    }
    
    // DisplayBreakup used by Cart
    type DisplayBreakup struct {

        
            CurrencyCode string  `json:"currency_code"`
            Key string  `json:"key"`
            CurrencySymbol string  `json:"currency_symbol"`
            Value float64  `json:"value"`
            Display string  `json:"display"`
            Message []string  `json:"message"`
         
    }
    
    // CouponBreakup used by Cart
    type CouponBreakup struct {

        
            UID string  `json:"uid"`
            Description string  `json:"description"`
            IsApplied bool  `json:"is_applied"`
            Code string  `json:"code"`
            CouponValue float64  `json:"coupon_value"`
            Title string  `json:"title"`
            SubTitle string  `json:"sub_title"`
            MaxDiscountValue float64  `json:"max_discount_value"`
            Type string  `json:"type"`
            Value float64  `json:"value"`
            MinimumCartValue float64  `json:"minimum_cart_value"`
            CouponType string  `json:"coupon_type"`
            Message string  `json:"message"`
         
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
            Display []DisplayBreakup  `json:"display"`
            Coupon CouponBreakup  `json:"coupon"`
            LoyaltyPoints LoyaltyPoints  `json:"loyalty_points"`
         
    }
    
    // PromoMeta used by Cart
    type PromoMeta struct {

        
            Message string  `json:"message"`
         
    }
    
    // ProductAvailabilitySize used by Cart
    type ProductAvailabilitySize struct {

        
            IsAvailable bool  `json:"is_available"`
            Display string  `json:"display"`
            Value string  `json:"value"`
         
    }
    
    // ProductAvailability used by Cart
    type ProductAvailability struct {

        
            Sizes []string  `json:"sizes"`
            AvailableSizes []ProductAvailabilitySize  `json:"available_sizes"`
            IsValid bool  `json:"is_valid"`
            OutOfStock bool  `json:"out_of_stock"`
            OtherStoreQuantity float64  `json:"other_store_quantity"`
            Deliverable bool  `json:"deliverable"`
         
    }
    
    // BasePrice used by Cart
    type BasePrice struct {

        
            CurrencySymbol string  `json:"currency_symbol"`
            CurrencyCode string  `json:"currency_code"`
            Effective float64  `json:"effective"`
            Marked float64  `json:"marked"`
         
    }
    
    // ArticlePriceInfo used by Cart
    type ArticlePriceInfo struct {

        
            Converted BasePrice  `json:"converted"`
            Base BasePrice  `json:"base"`
         
    }
    
    // StoreInfo used by Cart
    type StoreInfo struct {

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            StoreCode string  `json:"store_code"`
         
    }
    
    // BaseInfo used by Cart
    type BaseInfo struct {

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // ProductArticle used by Cart
    type ProductArticle struct {

        
            UID string  `json:"uid"`
            Identifier map[string]interface{}  `json:"identifier"`
            Quantity float64  `json:"quantity"`
            Price ArticlePriceInfo  `json:"price"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            IsGiftVisible bool  `json:"is_gift_visible"`
            Size string  `json:"size"`
            Meta map[string]interface{}  `json:"meta"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            ProductGroupTags []string  `json:"product_group_tags"`
            Store StoreInfo  `json:"store"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            Type string  `json:"type"`
            Seller BaseInfo  `json:"seller"`
            MtoQuantity float64  `json:"mto_quantity"`
            CartItemMeta map[string]interface{}  `json:"cart_item_meta"`
            SellerIdentifier string  `json:"seller_identifier"`
            GiftCard map[string]interface{}  `json:"gift_card"`
         
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
    
    // CouponDetails used by Cart
    type CouponDetails struct {

        
            DiscountTotalQuantity float64  `json:"discount_total_quantity"`
            DiscountSingleQuantity float64  `json:"discount_single_quantity"`
            Code string  `json:"code"`
         
    }
    
    // CategoryInfo used by Cart
    type CategoryInfo struct {

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
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
    
    // ActionQuery used by Cart
    type ActionQuery struct {

        
            ProductSlug []string  `json:"product_slug"`
         
    }
    
    // ProductAction used by Cart
    type ProductAction struct {

        
            URL string  `json:"url"`
            Type string  `json:"type"`
            Query ActionQuery  `json:"query"`
         
    }
    
    // CartProduct used by Cart
    type CartProduct struct {

        
            UID float64  `json:"uid"`
            Categories []CategoryInfo  `json:"categories"`
            Images []ProductImage  `json:"images"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            TeaserTag Tags  `json:"teaser_tag"`
            Tags []string  `json:"tags"`
            Type string  `json:"type"`
            ItemCode string  `json:"item_code"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Action ProductAction  `json:"action"`
            Brand BaseInfo  `json:"brand"`
         
    }
    
    // ProductPrice used by Cart
    type ProductPrice struct {

        
            Marked float64  `json:"marked"`
            CurrencyCode string  `json:"currency_code"`
            CurrencySymbol string  `json:"currency_symbol"`
            AddOn float64  `json:"add_on"`
            Selling float64  `json:"selling"`
            Effective float64  `json:"effective"`
         
    }
    
    // ProductPriceInfo used by Cart
    type ProductPriceInfo struct {

        
            Converted ProductPrice  `json:"converted"`
            Base ProductPrice  `json:"base"`
         
    }
    
    // CartProductIdentifer used by Cart
    type CartProductIdentifer struct {

        
            Identifier string  `json:"identifier"`
         
    }
    
    // DiscountRulesApp used by Cart
    type DiscountRulesApp struct {

        
            ItemCriteria map[string]interface{}  `json:"item_criteria"`
            MatchedBuyRules []string  `json:"matched_buy_rules"`
            RawOffer map[string]interface{}  `json:"raw_offer"`
            Offer map[string]interface{}  `json:"offer"`
         
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

        
            ItemImagesURL []string  `json:"item_images_url"`
            ItemID float64  `json:"item_id"`
            ItemSlug string  `json:"item_slug"`
            ItemBrandName string  `json:"item_brand_name"`
            ItemName string  `json:"item_name"`
            ItemPriceDetails map[string]interface{}  `json:"item_price_details"`
         
    }
    
    // AppliedFreeArticles used by Cart
    type AppliedFreeArticles struct {

        
            ParentItemIdentifier string  `json:"parent_item_identifier"`
            Quantity float64  `json:"quantity"`
            ArticleID string  `json:"article_id"`
            FreeGiftItemDetails FreeGiftItem  `json:"free_gift_item_details"`
         
    }
    
    // AppliedPromotion used by Cart
    type AppliedPromotion struct {

        
            ArticleQuantity float64  `json:"article_quantity"`
            PromotionName string  `json:"promotion_name"`
            MrpPromotion bool  `json:"mrp_promotion"`
            OfferText string  `json:"offer_text"`
            DiscountRules []DiscountRulesApp  `json:"discount_rules"`
            PromotionGroup string  `json:"promotion_group"`
            Ownership Ownership2  `json:"ownership"`
            BuyRules []BuyRules  `json:"buy_rules"`
            Amount float64  `json:"amount"`
            PromoID string  `json:"promo_id"`
            AppliedFreeArticles []AppliedFreeArticles  `json:"applied_free_articles"`
            PromotionType string  `json:"promotion_type"`
         
    }
    
    // CartProductInfo used by Cart
    type CartProductInfo struct {

        
            PromoMeta PromoMeta  `json:"promo_meta"`
            Availability ProductAvailability  `json:"availability"`
            Article ProductArticle  `json:"article"`
            CouponMessage string  `json:"coupon_message"`
            BulkOffer map[string]interface{}  `json:"bulk_offer"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            Coupon CouponDetails  `json:"coupon"`
            Product CartProduct  `json:"product"`
            Quantity float64  `json:"quantity"`
            Key string  `json:"key"`
            IsSet bool  `json:"is_set"`
            PricePerUnit ProductPriceInfo  `json:"price_per_unit"`
            Moq map[string]interface{}  `json:"moq"`
            Message string  `json:"message"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
            Price ProductPriceInfo  `json:"price"`
            Discount string  `json:"discount"`
            PromotionsApplied []AppliedPromotion  `json:"promotions_applied"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
         
    }
    
    // OpenapiCartDetailsResponse used by Cart
    type OpenapiCartDetailsResponse struct {

        
            BreakupValues CartBreakup  `json:"breakup_values"`
            IsValid bool  `json:"is_valid"`
            Message string  `json:"message"`
            Items []CartProductInfo  `json:"items"`
         
    }
    
    // OpenApiErrorResponse used by Cart
    type OpenApiErrorResponse struct {

        
            Errors map[string]interface{}  `json:"errors"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // ShippingAddress used by Cart
    type ShippingAddress struct {

        
            Country string  `json:"country"`
            CountryCode string  `json:"country_code"`
            AddressType string  `json:"address_type"`
            City string  `json:"city"`
            Pincode float64  `json:"pincode"`
            Address string  `json:"address"`
            Landmark string  `json:"landmark"`
            Email string  `json:"email"`
            Meta map[string]interface{}  `json:"meta"`
            CountryIsoCode string  `json:"country_iso_code"`
            AreaCodeSlug string  `json:"area_code_slug"`
            CountryPhoneCode string  `json:"country_phone_code"`
            Area string  `json:"area"`
            Name string  `json:"name"`
            Phone float64  `json:"phone"`
            State string  `json:"state"`
            AreaCode string  `json:"area_code"`
         
    }
    
    // OpenApiCartServiceabilityRequest used by Cart
    type OpenApiCartServiceabilityRequest struct {

        
            CartItems []CartItem  `json:"cart_items"`
            ShippingAddress ShippingAddress  `json:"shipping_address"`
         
    }
    
    // OpenApiCartServiceabilityResponse used by Cart
    type OpenApiCartServiceabilityResponse struct {

        
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            Items []CartProductInfo  `json:"items"`
            IsValid bool  `json:"is_valid"`
            Message string  `json:"message"`
         
    }
    
    // MultiTenderPaymentMeta used by Cart
    type MultiTenderPaymentMeta struct {

        
            CurrentStatus string  `json:"current_status"`
            PaymentID string  `json:"payment_id"`
            PaymentGateway string  `json:"payment_gateway"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            OrderID string  `json:"order_id"`
         
    }
    
    // MultiTenderPaymentMethod used by Cart
    type MultiTenderPaymentMethod struct {

        
            Mode string  `json:"mode"`
            Name string  `json:"name"`
            Amount float64  `json:"amount"`
            Meta MultiTenderPaymentMeta  `json:"meta"`
         
    }
    
    // CartItemMeta used by Cart
    type CartItemMeta struct {

        
            PrimaryItem bool  `json:"primary_item"`
            GroupID string  `json:"group_id"`
         
    }
    
    // OpenApiFiles used by Cart
    type OpenApiFiles struct {

        
            Key string  `json:"key"`
            Values []string  `json:"values"`
         
    }
    
    // OpenApiOrderItem used by Cart
    type OpenApiOrderItem struct {

        
            Size string  `json:"size"`
            AmountPaid float64  `json:"amount_paid"`
            ProductID float64  `json:"product_id"`
            Quantity float64  `json:"quantity"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            Discount float64  `json:"discount"`
            PaymentMethods []MultiTenderPaymentMethod  `json:"payment_methods"`
            EmployeeDiscount float64  `json:"employee_discount"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            Meta CartItemMeta  `json:"meta"`
            CodCharges float64  `json:"cod_charges"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            DeliveryCharges float64  `json:"delivery_charges"`
            PriceMarked float64  `json:"price_marked"`
            Files []OpenApiFiles  `json:"files"`
            PriceEffective float64  `json:"price_effective"`
            CashbackApplied float64  `json:"cashback_applied"`
         
    }
    
    // OpenApiPlatformCheckoutReq used by Cart
    type OpenApiPlatformCheckoutReq struct {

        
            Gstin string  `json:"gstin"`
            CouponValue float64  `json:"coupon_value"`
            CurrencyCode string  `json:"currency_code"`
            ShippingAddress ShippingAddress  `json:"shipping_address"`
            PaymentMode string  `json:"payment_mode"`
            CodCharges float64  `json:"cod_charges"`
            DeliveryCharges float64  `json:"delivery_charges"`
            Comment string  `json:"comment"`
            BillingAddress ShippingAddress  `json:"billing_address"`
            OrderID string  `json:"order_id"`
            CouponCode string  `json:"coupon_code"`
            PaymentMethods []MultiTenderPaymentMethod  `json:"payment_methods"`
            CartItems []OpenApiOrderItem  `json:"cart_items"`
            EmployeeDiscount map[string]interface{}  `json:"employee_discount"`
            CartValue float64  `json:"cart_value"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            Coupon string  `json:"coupon"`
            Files []OpenApiFiles  `json:"files"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            CashbackApplied float64  `json:"cashback_applied"`
         
    }
    
    // OpenApiCheckoutResponse used by Cart
    type OpenApiCheckoutResponse struct {

        
            OrderRefID string  `json:"order_ref_id"`
            OrderID string  `json:"order_id"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // AbandonedCart used by Cart
    type AbandonedCart struct {

        
            Gstin string  `json:"gstin"`
            BulkCouponDiscount float64  `json:"bulk_coupon_discount"`
            Shipments []map[string]interface{}  `json:"shipments"`
            IsArchive bool  `json:"is_archive"`
            DeliveryCharges map[string]interface{}  `json:"delivery_charges"`
            OrderID string  `json:"order_id"`
            MergeQty bool  `json:"merge_qty"`
            LastModified string  `json:"last_modified"`
            CreatedOn string  `json:"created_on"`
            IsActive bool  `json:"is_active"`
            IsDefault bool  `json:"is_default"`
            Cashback map[string]interface{}  `json:"cashback"`
            Coupon map[string]interface{}  `json:"coupon"`
            UserID string  `json:"user_id"`
            ExpireAt string  `json:"expire_at"`
            CheckoutMode string  `json:"checkout_mode"`
            AppID string  `json:"app_id"`
            BuyNow bool  `json:"buy_now"`
            UID float64  `json:"uid"`
            FyndCredits map[string]interface{}  `json:"fynd_credits"`
            ID string  `json:"_id"`
            PaymentMode string  `json:"payment_mode"`
            CodCharges map[string]interface{}  `json:"cod_charges"`
            Comment string  `json:"comment"`
            FcIndexMap []float64  `json:"fc_index_map"`
            Payments map[string]interface{}  `json:"payments"`
            PickUpCustomerDetails map[string]interface{}  `json:"pick_up_customer_details"`
            Discount float64  `json:"discount"`
            PaymentMethods []map[string]interface{}  `json:"payment_methods"`
            CartValue float64  `json:"cart_value"`
            Articles []map[string]interface{}  `json:"articles"`
            Meta map[string]interface{}  `json:"meta"`
            Promotion map[string]interface{}  `json:"promotion"`
         
    }
    
    // AbandonedCartResponse used by Cart
    type AbandonedCartResponse struct {

        
            Result map[string]interface{}  `json:"result"`
            Page Page  `json:"page"`
            Items []AbandonedCart  `json:"items"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
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

        
            Gstin string  `json:"gstin"`
            PanNo string  `json:"pan_no"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            IsValid bool  `json:"is_valid"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            ID string  `json:"id"`
            LastModified string  `json:"last_modified"`
            Currency CartCurrency  `json:"currency"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            PanConfig map[string]interface{}  `json:"pan_config"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            CheckoutMode string  `json:"checkout_mode"`
            BuyNow bool  `json:"buy_now"`
            CouponText string  `json:"coupon_text"`
            Comment string  `json:"comment"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            AppliedPromoDetails []AppliedPromotion  `json:"applied_promo_details"`
            Message string  `json:"message"`
            Items []CartProductInfo  `json:"items"`
         
    }
    
    // AddProductCart used by Cart
    type AddProductCart struct {

        
            ItemSize string  `json:"item_size"`
            ItemID float64  `json:"item_id"`
            Quantity float64  `json:"quantity"`
            ArticleAssignment map[string]interface{}  `json:"article_assignment"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Pos bool  `json:"pos"`
            ArticleID string  `json:"article_id"`
            ParentItemIdentifiers []map[string]string  `json:"parent_item_identifiers"`
            ProductGroupTags []string  `json:"product_group_tags"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            Meta map[string]interface{}  `json:"meta"`
            Display string  `json:"display"`
            SellerID float64  `json:"seller_id"`
            StoreID float64  `json:"store_id"`
         
    }
    
    // AddCartRequest used by Cart
    type AddCartRequest struct {

        
            Items []AddProductCart  `json:"items"`
            NewCart bool  `json:"new_cart"`
         
    }
    
    // AddCartDetailResponse used by Cart
    type AddCartDetailResponse struct {

        
            Partial bool  `json:"partial"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
            Cart CartDetailResponse  `json:"cart"`
         
    }
    
    // UpdateProductCart used by Cart
    type UpdateProductCart struct {

        
            ItemSize string  `json:"item_size"`
            ItemID float64  `json:"item_id"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
            Quantity float64  `json:"quantity"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ArticleID string  `json:"article_id"`
            Meta map[string]interface{}  `json:"meta"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            ItemIndex float64  `json:"item_index"`
         
    }
    
    // UpdateCartRequest used by Cart
    type UpdateCartRequest struct {

        
            Items []UpdateProductCart  `json:"items"`
            Operation string  `json:"operation"`
         
    }
    
    // UpdateCartDetailResponse used by Cart
    type UpdateCartDetailResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
            Cart CartDetailResponse  `json:"cart"`
         
    }
    
    // OverrideCartItemPromo used by Cart
    type OverrideCartItemPromo struct {

        
            ItemList []map[string]interface{}  `json:"item_list"`
            PromoAmount string  `json:"promo_amount"`
            PromoDesc string  `json:"promo_desc"`
            PromoID string  `json:"promo_id"`
            RwrdTndr string  `json:"rwrd_tndr"`
         
    }
    
    // OverrideCartItem used by Cart
    type OverrideCartItem struct {

        
            Size string  `json:"size"`
            PromoList []OverrideCartItemPromo  `json:"promo_list"`
            ItemID float64  `json:"item_id"`
            AmountPaid float64  `json:"amount_paid"`
            Quantity float64  `json:"quantity"`
            Discount float64  `json:"discount"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            PriceMarked float64  `json:"price_marked"`
            PriceEffective float64  `json:"price_effective"`
            SellerIdentifier string  `json:"seller_identifier"`
         
    }
    
    // OverrideCheckoutReq used by Cart
    type OverrideCheckoutReq struct {

        
            Aggregator string  `json:"aggregator"`
            MerchantCode string  `json:"merchant_code"`
            OrderingStore float64  `json:"ordering_store"`
            CartItems []OverrideCartItem  `json:"cart_items"`
            ShippingAddress map[string]interface{}  `json:"shipping_address"`
            PaymentMode string  `json:"payment_mode"`
            PaymentIdentifier string  `json:"payment_identifier"`
            CurrencyCode string  `json:"currency_code"`
            CartID string  `json:"cart_id"`
            OrderType string  `json:"order_type"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
         
    }
    
    // OverrideCheckoutResponse used by Cart
    type OverrideCheckoutResponse struct {

        
            Data map[string]interface{}  `json:"data"`
            Cart map[string]interface{}  `json:"cart"`
            Success string  `json:"success"`
            OrderID string  `json:"order_id"`
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

        
            Source map[string]interface{}  `json:"source"`
            CreatedOn string  `json:"created_on"`
            Meta map[string]interface{}  `json:"meta"`
            User map[string]interface{}  `json:"user"`
            Token string  `json:"token"`
         
    }
    
    // SharedCart used by Cart
    type SharedCart struct {

        
            Gstin string  `json:"gstin"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            IsValid bool  `json:"is_valid"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            ID string  `json:"id"`
            SharedCartDetails SharedCartDetails  `json:"shared_cart_details"`
            LastModified string  `json:"last_modified"`
            Currency CartCurrency  `json:"currency"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            CheckoutMode string  `json:"checkout_mode"`
            BuyNow bool  `json:"buy_now"`
            UID string  `json:"uid"`
            CouponText string  `json:"coupon_text"`
            Comment string  `json:"comment"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            Message string  `json:"message"`
            CartID float64  `json:"cart_id"`
            Items []CartProductInfo  `json:"items"`
         
    }
    
    // SharedCartResponse used by Cart
    type SharedCartResponse struct {

        
            Error string  `json:"error"`
            Cart SharedCart  `json:"cart"`
         
    }
    
    // CartList used by Cart
    type CartList struct {

        
            ItemCounts float64  `json:"item_counts"`
            CreatedOn string  `json:"created_on"`
            CartValue float64  `json:"cart_value"`
            CartID string  `json:"cart_id"`
            UserID string  `json:"user_id"`
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

        
            UID string  `json:"uid"`
            LastName string  `json:"last_name"`
            Mobile string  `json:"mobile"`
            ID string  `json:"_id"`
            Gender string  `json:"gender"`
            ExternalID string  `json:"external_id"`
            ModifiedOn string  `json:"modified_on"`
            CreatedAt string  `json:"created_at"`
            FirstName string  `json:"first_name"`
         
    }
    
    // UserCartMappingResponse used by Cart
    type UserCartMappingResponse struct {

        
            Gstin string  `json:"gstin"`
            PanNo string  `json:"pan_no"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            IsValid bool  `json:"is_valid"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            ID string  `json:"id"`
            LastModified string  `json:"last_modified"`
            Currency CartCurrency  `json:"currency"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            PanConfig map[string]interface{}  `json:"pan_config"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            CheckoutMode string  `json:"checkout_mode"`
            BuyNow bool  `json:"buy_now"`
            CouponText string  `json:"coupon_text"`
            Comment string  `json:"comment"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            AppliedPromoDetails []AppliedPromotion  `json:"applied_promo_details"`
            Message string  `json:"message"`
            User UserInfo  `json:"user"`
            Items []CartProductInfo  `json:"items"`
         
    }
    
    // PlatformAddCartRequest used by Cart
    type PlatformAddCartRequest struct {

        
            Items []AddProductCart  `json:"items"`
            UserID string  `json:"user_id"`
            NewCart bool  `json:"new_cart"`
         
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

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // CartItemCountResponse used by Cart
    type CartItemCountResponse struct {

        
            UserCartItemsCount float64  `json:"user_cart_items_count"`
         
    }
    
    // PageCoupon used by Cart
    type PageCoupon struct {

        
            Current float64  `json:"current"`
            HasNext bool  `json:"has_next"`
            HasPrevious bool  `json:"has_previous"`
            Total float64  `json:"total"`
            TotalItemCount float64  `json:"total_item_count"`
         
    }
    
    // Coupon used by Cart
    type Coupon struct {

        
            Description string  `json:"description"`
            IsApplied bool  `json:"is_applied"`
            CouponValue float64  `json:"coupon_value"`
            Title string  `json:"title"`
            Message string  `json:"message"`
            SubTitle string  `json:"sub_title"`
            MaxDiscountValue float64  `json:"max_discount_value"`
            ExpiresOn string  `json:"expires_on"`
            IsApplicable bool  `json:"is_applicable"`
            MinimumCartValue float64  `json:"minimum_cart_value"`
            CouponType string  `json:"coupon_type"`
            CouponCode string  `json:"coupon_code"`
         
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

        
            GoogleMapPoint map[string]interface{}  `json:"google_map_point"`
            GeoLocation GeoLocation  `json:"geo_location"`
            ID string  `json:"id"`
            IsActive bool  `json:"is_active"`
            City string  `json:"city"`
            AddressType string  `json:"address_type"`
            Email string  `json:"email"`
            Tags []string  `json:"tags"`
            CheckoutMode string  `json:"checkout_mode"`
            UserID string  `json:"user_id"`
            Area string  `json:"area"`
            Country string  `json:"country"`
            CountryCode string  `json:"country_code"`
            CreatedByUserID string  `json:"created_by_user_id"`
            IsDefaultAddress bool  `json:"is_default_address"`
            Phone string  `json:"phone"`
            State string  `json:"state"`
            AreaCode string  `json:"area_code"`
            Address string  `json:"address"`
            Landmark string  `json:"landmark"`
            Meta map[string]interface{}  `json:"meta"`
            AreaCodeSlug string  `json:"area_code_slug"`
            CartID string  `json:"cart_id"`
            Name string  `json:"name"`
         
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

        
            ID string  `json:"id"`
            BillingAddressID string  `json:"billing_address_id"`
            CartID string  `json:"cart_id"`
            UserID string  `json:"user_id"`
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

        
            Shipments float64  `json:"shipments"`
            DpOptions map[string]interface{}  `json:"dp_options"`
            ShipmentType string  `json:"shipment_type"`
            Promise ShipmentPromise  `json:"promise"`
            Articles []ShipmentArticle  `json:"articles"`
            BoxType string  `json:"box_type"`
            DpID string  `json:"dp_id"`
            OrderType string  `json:"order_type"`
            FulfillmentType string  `json:"fulfillment_type"`
            Items []CartProductInfo  `json:"items"`
            FulfillmentID float64  `json:"fulfillment_id"`
         
    }
    
    // PlatformCartShipmentsResponse used by Cart
    type PlatformCartShipmentsResponse struct {

        
            Gstin string  `json:"gstin"`
            PanNo string  `json:"pan_no"`
            Shipments []PlatformShipmentResponse  `json:"shipments"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            IsValid bool  `json:"is_valid"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            ID string  `json:"id"`
            LastModified string  `json:"last_modified"`
            Currency CartCurrency  `json:"currency"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            PanConfig map[string]interface{}  `json:"pan_config"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            CheckoutMode string  `json:"checkout_mode"`
            BuyNow bool  `json:"buy_now"`
            CouponText string  `json:"coupon_text"`
            Comment string  `json:"comment"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            AppliedPromoDetails []AppliedPromotion  `json:"applied_promo_details"`
            Message string  `json:"message"`
            Error bool  `json:"error"`
            StaffUserID string  `json:"staff_user_id"`
            Items []CartProductInfo  `json:"items"`
         
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

        
            Gstin string  `json:"gstin"`
            PanNo string  `json:"pan_no"`
            GiftDetails map[string]interface{}  `json:"gift_details"`
            Comment string  `json:"comment"`
            CheckoutMode string  `json:"checkout_mode"`
            PickUpCustomerDetails map[string]interface{}  `json:"pick_up_customer_details"`
            StaffUserID string  `json:"staff_user_id"`
         
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

        
            LastName string  `json:"last_name"`
            ID string  `json:"_id"`
            EmployeeCode string  `json:"employee_code"`
            User string  `json:"user"`
            FirstName string  `json:"first_name"`
         
    }
    
    // Files used by Cart
    type Files struct {

        
            Key string  `json:"key"`
            Values []string  `json:"values"`
         
    }
    
    // PlatformCartCheckoutDetailRequest used by Cart
    type PlatformCartCheckoutDetailRequest struct {

        
            MerchantCode string  `json:"merchant_code"`
            BillingAddressID string  `json:"billing_address_id"`
            PaymentParams map[string]interface{}  `json:"payment_params"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
            ID string  `json:"id"`
            AddressID string  `json:"address_id"`
            Pos bool  `json:"pos"`
            UserID string  `json:"user_id"`
            OrderType string  `json:"order_type"`
            CheckoutMode string  `json:"checkout_mode"`
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
            PickAtStoreUID float64  `json:"pick_at_store_uid"`
            EmployeeCode string  `json:"employee_code"`
            OrderingStore float64  `json:"ordering_store"`
            CallbackURL string  `json:"callback_url"`
            PaymentMode string  `json:"payment_mode"`
            PaymentAutoConfirm bool  `json:"payment_auto_confirm"`
            Staff StaffCheckout  `json:"staff"`
            Aggregator string  `json:"aggregator"`
            PaymentIdentifier string  `json:"payment_identifier"`
            Meta map[string]interface{}  `json:"meta"`
            DeviceID string  `json:"device_id"`
            Files []Files  `json:"files"`
         
    }
    
    // CheckCart used by Cart
    type CheckCart struct {

        
            Gstin string  `json:"gstin"`
            StoreEmps []map[string]interface{}  `json:"store_emps"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            DeliveryCharges float64  `json:"delivery_charges"`
            IsValid bool  `json:"is_valid"`
            OrderID string  `json:"order_id"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            CodMessage string  `json:"cod_message"`
            ID string  `json:"id"`
            LastModified string  `json:"last_modified"`
            Currency CartCurrency  `json:"currency"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            ErrorMessage string  `json:"error_message"`
            CodAvailable bool  `json:"cod_available"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            CheckoutMode string  `json:"checkout_mode"`
            UserType string  `json:"user_type"`
            BuyNow bool  `json:"buy_now"`
            UID string  `json:"uid"`
            CouponText string  `json:"coupon_text"`
            CodCharges float64  `json:"cod_charges"`
            Comment string  `json:"comment"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            StoreCode string  `json:"store_code"`
            Message string  `json:"message"`
            DeliveryChargeOrderValue float64  `json:"delivery_charge_order_value"`
            CartID float64  `json:"cart_id"`
            Items []CartProductInfo  `json:"items"`
            Success bool  `json:"success"`
         
    }
    
    // CartCheckoutResponse used by Cart
    type CartCheckoutResponse struct {

        
            Data map[string]interface{}  `json:"data"`
            PaymentConfirmURL string  `json:"payment_confirm_url"`
            CallbackURL string  `json:"callback_url"`
            Cart CheckCart  `json:"cart"`
            AppInterceptURL string  `json:"app_intercept_url"`
            OrderID string  `json:"order_id"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // CartDeliveryModesResponse used by Cart
    type CartDeliveryModesResponse struct {

        
            AvailableModes []string  `json:"available_modes"`
            PickupStores []float64  `json:"pickup_stores"`
         
    }
    
    // PickupStoreDetail used by Cart
    type PickupStoreDetail struct {

        
            UID float64  `json:"uid"`
            ID float64  `json:"id"`
            Country string  `json:"country"`
            AddressType string  `json:"address_type"`
            City string  `json:"city"`
            Pincode float64  `json:"pincode"`
            Address string  `json:"address"`
            Landmark string  `json:"landmark"`
            Email string  `json:"email"`
            AreaCodeSlug string  `json:"area_code_slug"`
            Area string  `json:"area"`
            Name string  `json:"name"`
            Phone string  `json:"phone"`
            StoreCode string  `json:"store_code"`
            State string  `json:"state"`
            AreaCode string  `json:"area_code"`
         
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
            PaymentMode string  `json:"payment_mode"`
            PaymentIdentifier string  `json:"payment_identifier"`
            AggregatorName string  `json:"aggregator_name"`
         
    }
    
    // CouponValidity used by Cart
    type CouponValidity struct {

        
            Code string  `json:"code"`
            Discount float64  `json:"discount"`
            Title string  `json:"title"`
            NextValidationRequired bool  `json:"next_validation_required"`
            DisplayMessageEn string  `json:"display_message_en"`
            Valid bool  `json:"valid"`
         
    }
    
    // PaymentCouponValidate used by Cart
    type PaymentCouponValidate struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
            CouponValidity CouponValidity  `json:"coupon_validity"`
         
    }
    
    // PaymentMeta used by Cart
    type PaymentMeta struct {

        
            Type string  `json:"type"`
            PaymentGateway string  `json:"payment_gateway"`
            MerchantCode string  `json:"merchant_code"`
            PaymentIdentifier string  `json:"payment_identifier"`
         
    }
    
    // PaymentMethod used by Cart
    type PaymentMethod struct {

        
            Mode string  `json:"mode"`
            PaymentMeta PaymentMeta  `json:"payment_meta"`
            Amount float64  `json:"amount"`
            Payment string  `json:"payment"`
            Name string  `json:"name"`
         
    }
    
    // PlatformCartCheckoutDetailV2Request used by Cart
    type PlatformCartCheckoutDetailV2Request struct {

        
            MerchantCode string  `json:"merchant_code"`
            BillingAddressID string  `json:"billing_address_id"`
            PaymentParams map[string]interface{}  `json:"payment_params"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
            ID string  `json:"id"`
            AddressID string  `json:"address_id"`
            Pos bool  `json:"pos"`
            UserID string  `json:"user_id"`
            OrderType string  `json:"order_type"`
            CheckoutMode string  `json:"checkout_mode"`
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
            PickAtStoreUID float64  `json:"pick_at_store_uid"`
            EmployeeCode string  `json:"employee_code"`
            OrderingStore float64  `json:"ordering_store"`
            CallbackURL string  `json:"callback_url"`
            PaymentMode string  `json:"payment_mode"`
            PaymentAutoConfirm bool  `json:"payment_auto_confirm"`
            Staff StaffCheckout  `json:"staff"`
            Aggregator string  `json:"aggregator"`
            PaymentMethods []PaymentMethod  `json:"payment_methods"`
            CustomMeta map[string]interface{}  `json:"custom_meta"`
            PaymentIdentifier string  `json:"payment_identifier"`
            Meta map[string]interface{}  `json:"meta"`
            DeviceID string  `json:"device_id"`
            Files []Files  `json:"files"`
         
    }
    
    // UpdateCartPaymentRequestV2 used by Cart
    type UpdateCartPaymentRequestV2 struct {

        
            ID string  `json:"id"`
            MerchantCode string  `json:"merchant_code"`
            AddressID string  `json:"address_id"`
            PaymentMethods []PaymentMethod  `json:"payment_methods"`
            PaymentMode string  `json:"payment_mode"`
            PaymentIdentifier string  `json:"payment_identifier"`
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
    

    
    // ServiceabilityErrorResponse used by Serviceability
    type ServiceabilityErrorResponse struct {

        
            Value string  `json:"value"`
            Message string  `json:"message"`
            Type string  `json:"type"`
         
    }
    
    // ApplicationServiceabilityConfig used by Serviceability
    type ApplicationServiceabilityConfig struct {

        
            ChannelID string  `json:"channel_id"`
            ServiceabilityType string  `json:"serviceability_type"`
            ChannelType string  `json:"channel_type"`
         
    }
    
    // ApplicationServiceabilityConfigResponse used by Serviceability
    type ApplicationServiceabilityConfigResponse struct {

        
            Error ServiceabilityErrorResponse  `json:"error"`
            Data ApplicationServiceabilityConfig  `json:"data"`
            Success bool  `json:"success"`
         
    }
    
    // EntityRegionView_Request used by Serviceability
    type EntityRegionView_Request struct {

        
            ParentID []string  `json:"parent_id"`
            SubType []string  `json:"sub_type"`
         
    }
    
    // EntityRegionView_Error used by Serviceability
    type EntityRegionView_Error struct {

        
            Value string  `json:"value"`
            Message string  `json:"message"`
            Type string  `json:"type"`
         
    }
    
    // EntityRegionView_page used by Serviceability
    type EntityRegionView_page struct {

        
            Current float64  `json:"current"`
            Size float64  `json:"size"`
            ItemTotal float64  `json:"item_total"`
            HasNext bool  `json:"has_next"`
            Type string  `json:"type"`
         
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
            Page EntityRegionView_page  `json:"page"`
            Data []EntityRegionView_Items  `json:"data"`
            Success bool  `json:"success"`
         
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

        
            Channels ListViewChannels  `json:"channels"`
            IsActive bool  `json:"is_active"`
            PincodesCount float64  `json:"pincodes_count"`
            CompanyID float64  `json:"company_id"`
            Product ListViewProduct  `json:"product"`
            Name string  `json:"name"`
            ZoneID string  `json:"zone_id"`
            StoresCount float64  `json:"stores_count"`
            Slug string  `json:"slug"`
         
    }
    
    // ListViewSummary used by Serviceability
    type ListViewSummary struct {

        
            TotalZones float64  `json:"total_zones"`
            TotalPincodesServed float64  `json:"total_pincodes_served"`
            TotalActiveZones float64  `json:"total_active_zones"`
         
    }
    
    // ZoneDataItem used by Serviceability
    type ZoneDataItem struct {

        
            Current float64  `json:"current"`
            Size float64  `json:"size"`
            ItemTotal float64  `json:"item_total"`
            Type string  `json:"type"`
            HasNext bool  `json:"has_next"`
         
    }
    
    // ListViewResponse used by Serviceability
    type ListViewResponse struct {

        
            Items []ListViewItems  `json:"items"`
            Summary []ListViewSummary  `json:"summary"`
            Page []ZoneDataItem  `json:"page"`
         
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

        
            Items []map[string]interface{}  `json:"items"`
            Page []CompanyStoreView_PageItems  `json:"page"`
         
    }
    
    // GetZoneDataViewChannels used by Serviceability
    type GetZoneDataViewChannels struct {

        
            ChannelID string  `json:"channel_id"`
            ChannelType string  `json:"channel_type"`
         
    }
    
    // ZoneProductTypes used by Serviceability
    type ZoneProductTypes struct {

        
            Tags []string  `json:"tags"`
            Type string  `json:"type"`
         
    }
    
    // ZoneMappingType used by Serviceability
    type ZoneMappingType struct {

        
            State []string  `json:"state"`
            Pincode []string  `json:"pincode"`
            Country string  `json:"country"`
         
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

        
            Pincode string  `json:"pincode"`
            Country string  `json:"country"`
         
    }
    
    // Zone used by Serviceability
    type Zone struct {

        
            ZoneID string  `json:"zone_id"`
            Tags []string  `json:"tags"`
            AssignmentPreference string  `json:"assignment_preference"`
            IsActive bool  `json:"is_active"`
            Type string  `json:"type"`
            Name string  `json:"name"`
            StoreIds []float64  `json:"store_ids"`
            Slug string  `json:"slug"`
         
    }
    
    // GetZoneFromPincodeViewResponse used by Serviceability
    type GetZoneFromPincodeViewResponse struct {

        
            ServiceabilityType string  `json:"serviceability_type"`
            Zones []Zone  `json:"zones"`
         
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
    
    // DocumentsResponse used by Serviceability
    type DocumentsResponse struct {

        
            Value string  `json:"value"`
            Verified bool  `json:"verified"`
            LegalName string  `json:"legal_name"`
            Type string  `json:"type"`
         
    }
    
    // IntegrationTypeResponse used by Serviceability
    type IntegrationTypeResponse struct {

        
            Order string  `json:"order"`
            Inventory string  `json:"inventory"`
         
    }
    
    // ProductReturnConfigResponse used by Serviceability
    type ProductReturnConfigResponse struct {

        
            OnSameStore bool  `json:"on_same_store"`
         
    }
    
    // OpeningClosing used by Serviceability
    type OpeningClosing struct {

        
            Minute float64  `json:"minute"`
            Hour float64  `json:"hour"`
         
    }
    
    // TimmingResponse used by Serviceability
    type TimmingResponse struct {

        
            Weekday string  `json:"weekday"`
            Opening OpeningClosing  `json:"opening"`
            Open bool  `json:"open"`
            Closing OpeningClosing  `json:"closing"`
         
    }
    
    // ModifiedByResponse used by Serviceability
    type ModifiedByResponse struct {

        
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
         
    }
    
    // ContactNumberResponse used by Serviceability
    type ContactNumberResponse struct {

        
            CountryCode float64  `json:"country_code"`
            Number string  `json:"number"`
         
    }
    
    // AddressResponse used by Serviceability
    type AddressResponse struct {

        
            Address1 string  `json:"address1"`
            Address2 string  `json:"address2"`
            Pincode float64  `json:"pincode"`
            Country string  `json:"country"`
            Longitude float64  `json:"longitude"`
            City string  `json:"city"`
            Latitude float64  `json:"latitude"`
            Landmark string  `json:"landmark"`
            State string  `json:"state"`
         
    }
    
    // CreatedByResponse used by Serviceability
    type CreatedByResponse struct {

        
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
         
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
    
    // Dp used by Serviceability
    type Dp struct {

        
            FmPriority float64  `json:"fm_priority"`
            AssignDpFromSb bool  `json:"assign_dp_from_sb"`
            ExternalAccountID string  `json:"external_account_id"`
            RvpPriority float64  `json:"rvp_priority"`
            Operations []string  `json:"operations"`
            AreaCode float64  `json:"area_code"`
            InternalAccountID string  `json:"internal_account_id"`
            TransportMode string  `json:"transport_mode"`
            LmPriority float64  `json:"lm_priority"`
            PaymentMode string  `json:"payment_mode"`
         
    }
    
    // LogisticsResponse used by Serviceability
    type LogisticsResponse struct {

        
            Dp Dp  `json:"dp"`
            Override bool  `json:"override"`
         
    }
    
    // WarningsResponse used by Serviceability
    type WarningsResponse struct {

        
            StoreAddress string  `json:"store_address"`
         
    }
    
    // ItemResponse used by Serviceability
    type ItemResponse struct {

        
            Manager ManagerResponse  `json:"manager"`
            Documents []DocumentsResponse  `json:"documents"`
            ModifiedOn string  `json:"modified_on"`
            DisplayName string  `json:"display_name"`
            UID float64  `json:"uid"`
            IntegrationType IntegrationTypeResponse  `json:"integration_type"`
            ProductReturnConfig ProductReturnConfigResponse  `json:"product_return_config"`
            Timing []TimmingResponse  `json:"timing"`
            VerifiedBy ModifiedByResponse  `json:"verified_by"`
            ModifiedBy ModifiedByResponse  `json:"modified_by"`
            Stage string  `json:"stage"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Name string  `json:"name"`
            Cls string  `json:"_cls"`
            ContactNumbers []ContactNumberResponse  `json:"contact_numbers"`
            CompanyID float64  `json:"company_id"`
            Address AddressResponse  `json:"address"`
            VerifiedOn string  `json:"verified_on"`
            NotificationEmails []string  `json:"notification_emails"`
            StoreType string  `json:"store_type"`
            CreatedBy CreatedByResponse  `json:"created_by"`
            SubType string  `json:"sub_type"`
            GstCredentials GstCredentialsResponse  `json:"gst_credentials"`
            Logistics LogisticsResponse  `json:"logistics"`
            Code string  `json:"code"`
            Company float64  `json:"company"`
            CreatedOn string  `json:"created_on"`
            Warnings WarningsResponse  `json:"warnings"`
         
    }
    
    // ServiceabilityPageResponse used by Serviceability
    type ServiceabilityPageResponse struct {

        
            Current float64  `json:"current"`
            Size float64  `json:"size"`
            ItemTotal float64  `json:"item_total"`
            HasNext bool  `json:"has_next"`
            Type string  `json:"type"`
         
    }
    
    // GetStoresViewResponse used by Serviceability
    type GetStoresViewResponse struct {

        
            Items []ItemResponse  `json:"items"`
            Page ServiceabilityPageResponse  `json:"page"`
         
    }
    
    // ReAssignStoreRequest used by Serviceability
    type ReAssignStoreRequest struct {

        
            Articles []map[string]interface{}  `json:"articles"`
            Identifier string  `json:"identifier"`
            ToPincode string  `json:"to_pincode"`
            IgnoredLocations []string  `json:"ignored_locations"`
            Configuration map[string]interface{}  `json:"configuration"`
         
    }
    
    // ReAssignStoreResponse used by Serviceability
    type ReAssignStoreResponse struct {

        
            Error map[string]interface{}  `json:"error"`
            ToPincode string  `json:"to_pincode"`
            Articles []map[string]interface{}  `json:"articles"`
            Success bool  `json:"success"`
         
    }
    
    // ApplicationCompanyDpViewRequest used by Serviceability
    type ApplicationCompanyDpViewRequest struct {

        
            DpID string  `json:"dp_id"`
         
    }
    
    // ApplicationCompanyDpViewResponse used by Serviceability
    type ApplicationCompanyDpViewResponse struct {

        
            CourierPartnerID float64  `json:"courier_partner_id"`
            ApplicationID string  `json:"application_id"`
            CompanyID float64  `json:"company_id"`
            Success bool  `json:"success"`
         
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

        
            Error interface{}  `json:"error"`
            StatusCode string  `json:"status_code"`
            Success string  `json:"success"`
         
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

        
            PlanRules map[string]interface{}  `json:"plan_rules"`
            IsSelfShip bool  `json:"is_self_ship"`
            Stage string  `json:"stage"`
            DpID string  `json:"dp_id"`
            Name string  `json:"name"`
            PlanID string  `json:"plan_id"`
            AccountID string  `json:"account_id"`
         
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

        
            Items []Dp1  `json:"items"`
            Page Page  `json:"page"`
            Success bool  `json:"success"`
         
    }
    
    // ErrorResponse used by Serviceability
    type ErrorResponse struct {

        
            Value string  `json:"value"`
            Message string  `json:"message"`
            Type string  `json:"type"`
         
    }
    
    // DpAccountFailureResponse used by Serviceability
    type DpAccountFailureResponse struct {

        
            Error []ErrorResponse  `json:"error"`
            StatusCode float64  `json:"status_code"`
            Success bool  `json:"success"`
         
    }
    
    // CompanyDpAccountRequest used by Serviceability
    type CompanyDpAccountRequest struct {

        
            Data []Dp1  `json:"data"`
         
    }
    
    // CompanyDpAccountResponse used by Serviceability
    type CompanyDpAccountResponse struct {

        
            Success bool  `json:"success"`
         
    }
    
    // DpRulesUpdateRequest used by Serviceability
    type DpRulesUpdateRequest struct {

        
            DpIds map[string]map[string]interface{}  `json:"dp_ids"`
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            Conditions []map[string]interface{}  `json:"conditions"`
         
    }
    
    // DpRuleResponse used by Serviceability
    type DpRuleResponse struct {

        
            ModifiedOn string  `json:"modified_on"`
            DpIds map[string]interface{}  `json:"dp_ids"`
            IsActive bool  `json:"is_active"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            UID string  `json:"uid"`
            Name string  `json:"name"`
            CreatedOn string  `json:"created_on"`
            Conditions []string  `json:"conditions"`
            CompanyID float64  `json:"company_id"`
         
    }
    
    // DpRuleUpdateSuccessResponse used by Serviceability
    type DpRuleUpdateSuccessResponse struct {

        
            Data DpRuleResponse  `json:"data"`
            StatusCode float64  `json:"status_code"`
            Success bool  `json:"success"`
         
    }
    
    // FailureResponse used by Serviceability
    type FailureResponse struct {

        
            Error []ErrorResponse  `json:"error"`
            StatusCode float64  `json:"status_code"`
            Success bool  `json:"success"`
         
    }
    
    // DpSchemaInRuleListing used by Serviceability
    type DpSchemaInRuleListing struct {

        
            PlanRules map[string]interface{}  `json:"plan_rules"`
            IsSelfShip bool  `json:"is_self_ship"`
            Stage string  `json:"stage"`
            DpID string  `json:"dp_id"`
            Name string  `json:"name"`
            PlanID string  `json:"plan_id"`
            AccountID string  `json:"account_id"`
            Priority float64  `json:"priority"`
         
    }
    
    // DpRule used by Serviceability
    type DpRule struct {

        
            DpIds map[string]DpSchemaInRuleListing  `json:"dp_ids"`
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            Conditions []map[string]interface{}  `json:"conditions"`
            CompanyID float64  `json:"company_id"`
         
    }
    
    // DpRuleSuccessResponse used by Serviceability
    type DpRuleSuccessResponse struct {

        
            Data DpRule  `json:"data"`
            StatusCode float64  `json:"status_code"`
            Success bool  `json:"success"`
         
    }
    
    // DpMultipleRuleSuccessResponse used by Serviceability
    type DpMultipleRuleSuccessResponse struct {

        
            Items []DpRule  `json:"items"`
            Page Page  `json:"page"`
            Success bool  `json:"success"`
         
    }
    
    // DpIds used by Serviceability
    type DpIds struct {

        
            Enabled bool  `json:"enabled"`
            Priority float64  `json:"priority"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // DpRuleRequest used by Serviceability
    type DpRuleRequest struct {

        
            DpIds map[string]DpIds  `json:"dp_ids"`
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            Conditions []map[string]interface{}  `json:"conditions"`
            CompanyID float64  `json:"company_id"`
         
    }
    
    // DPCompanyRuleRequest used by Serviceability
    type DPCompanyRuleRequest struct {

        
            RuleIds []string  `json:"rule_ids"`
         
    }
    
    // DPCompanyRuleResponse used by Serviceability
    type DPCompanyRuleResponse struct {

        
            Data []DpRuleResponse  `json:"data"`
            StatusCode float64  `json:"status_code"`
            Success bool  `json:"success"`
         
    }
    
    // DPApplicationRuleRequest used by Serviceability
    type DPApplicationRuleRequest struct {

        
            ShippingRules []string  `json:"shipping_rules"`
         
    }
    
    // DPApplicationRuleResponse used by Serviceability
    type DPApplicationRuleResponse struct {

        
            Data []DpRuleResponse  `json:"data"`
            StatusCode bool  `json:"status_code"`
            Success bool  `json:"success"`
         
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

        
            Error ServiceabilityErrorResponse  `json:"error"`
            Data ApplicationSelfShipConfig  `json:"data"`
            Success bool  `json:"success"`
         
    }
    

    
    // GenerateReportFilters used by Finance
    type GenerateReportFilters struct {

        
            Company []string  `json:"company"`
            Channel []string  `json:"channel"`
            Brand []string  `json:"brand"`
         
    }
    
    // GenerateReportMeta used by Finance
    type GenerateReportMeta struct {

        
            Company string  `json:"company"`
            Channel string  `json:"channel"`
            Brand string  `json:"brand"`
         
    }
    
    // GenerateReportPlatform used by Finance
    type GenerateReportPlatform struct {

        
            ReportID string  `json:"report_id"`
            Filters GenerateReportFilters  `json:"filters"`
            EndDate string  `json:"end_date"`
            Meta GenerateReportMeta  `json:"meta"`
            StartDate string  `json:"start_date"`
         
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

        
            ItemCount float64  `json:"item_count"`
            Headers []string  `json:"headers"`
            EndDate string  `json:"end_date"`
            StartDate string  `json:"start_date"`
            Page Page  `json:"page"`
            Items [][]string  `json:"items"`
         
    }
    
    // Error used by Finance
    type Error struct {

        
            Reason string  `json:"reason"`
            Success bool  `json:"success"`
         
    }
    
    // DownloadReport used by Finance
    type DownloadReport struct {

        
            StartDate string  `json:"start_date"`
            Pagesize float64  `json:"pagesize"`
            Page float64  `json:"page"`
            EndDate string  `json:"end_date"`
         
    }
    
    // DownloadReportItems used by Finance
    type DownloadReportItems struct {

        
            ReportID string  `json:"report_id"`
            Filters GenerateReportFilters  `json:"filters"`
            EndDate string  `json:"end_date"`
            TypeOfRequest string  `json:"type_of_request"`
            Meta GenerateReportMeta  `json:"meta"`
            StartDate string  `json:"start_date"`
         
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

        
            TableName string  `json:"table_name"`
            Project []string  `json:"project"`
            Filters GetEngineFilters  `json:"filters"`
         
    }
    
    // GetEngineRequest used by Finance
    type GetEngineRequest struct {

        
            Data GetEngineData  `json:"data"`
         
    }
    
    // GetEngineResponse used by Finance
    type GetEngineResponse struct {

        
            ItemCount float64  `json:"item_count"`
            Page Page  `json:"page"`
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

        
            Data GetDocs  `json:"data"`
            Success bool  `json:"success"`
         
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

        
            ID string  `json:"id"`
            PdfS3URL string  `json:"pdf_s3_url"`
         
    }
    
    // DownloadCreditDebitNoteResponse used by Finance
    type DownloadCreditDebitNoteResponse struct {

        
            Data []DownloadCreditDebitNoteResponseData  `json:"data"`
            Success bool  `json:"success"`
         
    }
    
    // PaymentProcessPayload used by Finance
    type PaymentProcessPayload struct {

        
            InvoiceNumber string  `json:"invoice_number"`
            SellerID string  `json:"seller_id"`
            Currency string  `json:"currency"`
            Amount string  `json:"amount"`
            TransactionType string  `json:"transaction_type"`
            SourceReference string  `json:"source_reference"`
            Meta map[string]interface{}  `json:"meta"`
            Platform string  `json:"platform"`
            TotalAmount string  `json:"total_amount"`
            ModeOfPayment string  `json:"mode_of_payment"`
         
    }
    
    // PaymentProcessRequest used by Finance
    type PaymentProcessRequest struct {

        
            Data PaymentProcessPayload  `json:"data"`
         
    }
    
    // PaymentProcessResponse used by Finance
    type PaymentProcessResponse struct {

        
            TransactionID string  `json:"transaction_id"`
            Code float64  `json:"code"`
            Message string  `json:"message"`
            RedirectURL string  `json:"redirect_url"`
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
            ItemCount float64  `json:"item_count"`
            Headers []string  `json:"headers"`
            Message string  `json:"message"`
            ShowMr bool  `json:"show_mr"`
            Page map[string]interface{}  `json:"page"`
            Items []map[string]interface{}  `json:"items"`
         
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

        
            IsCreditlineOpted bool  `json:"is_creditline_opted"`
            Code float64  `json:"code"`
         
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
            PaymentStatusList []InvoiceTypeResponseItems  `json:"payment_status_list"`
            Success bool  `json:"success"`
         
    }
    
    // InoviceListingPayloadDataFilters used by Finance
    type InoviceListingPayloadDataFilters struct {

        
            InvoiceType []string  `json:"invoice_type"`
            CompanyID []string  `json:"company_id"`
            PaymentStatus []string  `json:"payment_status"`
         
    }
    
    // InvoiceListingPayloadData used by Finance
    type InvoiceListingPayloadData struct {

        
            Filters InoviceListingPayloadDataFilters  `json:"filters"`
            EndDate string  `json:"end_date"`
            Search string  `json:"search"`
            StartDate string  `json:"start_date"`
            Page float64  `json:"page"`
            PageSize float64  `json:"page_size"`
         
    }
    
    // InvoiceListingRequest used by Finance
    type InvoiceListingRequest struct {

        
            Data InvoiceListingPayloadData  `json:"data"`
         
    }
    
    // UnpaidInvoiceDataItems used by Finance
    type UnpaidInvoiceDataItems struct {

        
            TotalUnpaidInvoiceCount float64  `json:"total_unpaid_invoice_count"`
            TotalUnpaidAmount float64  `json:"total_unpaid_amount"`
            Currency string  `json:"currency"`
         
    }
    
    // InvoiceListingResponseItems used by Finance
    type InvoiceListingResponseItems struct {

        
            InvoiceNumber string  `json:"invoice_number"`
            Amount string  `json:"amount"`
            InvoiceID string  `json:"invoice_id"`
            InvoiceType string  `json:"invoice_type"`
            Status string  `json:"status"`
            DueDate string  `json:"due_date"`
            Company string  `json:"company"`
            IsDownloadable bool  `json:"is_downloadable"`
            Period string  `json:"period"`
            InvoiceDate string  `json:"invoice_date"`
         
    }
    
    // InvoiceListingResponse used by Finance
    type InvoiceListingResponse struct {

        
            ItemCount float64  `json:"item_count"`
            Page Page  `json:"page"`
            UnpaidInvoiceData UnpaidInvoiceDataItems  `json:"unpaid_invoice_data"`
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
            Data []string  `json:"data"`
            Success bool  `json:"success"`
         
    }
    
