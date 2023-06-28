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
            Success bool  `json:"success"`
            Aggregators []map[string]interface{}  `json:"aggregators"`
            DisplayFields []string  `json:"display_fields"`
            ExcludedFields []string  `json:"excluded_fields"`
            AppID string  `json:"app_id"`
         
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
            Key string  `json:"key"`
            MerchantSalt string  `json:"merchant_salt"`
            Secret string  `json:"secret"`
         
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

        
            Code string  `json:"code"`
            Description string  `json:"description"`
         
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
            Code string  `json:"code"`
            PackageName string  `json:"package_name"`
            DisplayName string  `json:"display_name"`
         
    }
    
    // IntentAppErrorList used by Payment
    type IntentAppErrorList struct {

        
            Code string  `json:"code"`
            PackageName string  `json:"package_name"`
         
    }
    
    // PaymentModeList used by Payment
    type PaymentModeList struct {

        
            FyndVpa string  `json:"fynd_vpa"`
            DisplayName string  `json:"display_name"`
            AggregatorName string  `json:"aggregator_name"`
            CardType string  `json:"card_type"`
            RemainingLimit float64  `json:"remaining_limit"`
            ExpMonth float64  `json:"exp_month"`
            LogoURL PaymentModeLogo  `json:"logo_url"`
            CardBrandImage string  `json:"card_brand_image"`
            CardID string  `json:"card_id"`
            Nickname string  `json:"nickname"`
            Expired bool  `json:"expired"`
            Code string  `json:"code"`
            RetryCount float64  `json:"retry_count"`
            MerchantCode string  `json:"merchant_code"`
            CardReference string  `json:"card_reference"`
            DisplayPriority float64  `json:"display_priority"`
            CompliantWithTokenisationGuidelines bool  `json:"compliant_with_tokenisation_guidelines"`
            CardToken string  `json:"card_token"`
            CardIssuer string  `json:"card_issuer"`
            CardFingerprint string  `json:"card_fingerprint"`
            CodLimitPerOrder float64  `json:"cod_limit_per_order"`
            Timeout float64  `json:"timeout"`
            CardBrand string  `json:"card_brand"`
            IntentFlow bool  `json:"intent_flow"`
            CardName string  `json:"card_name"`
            Name string  `json:"name"`
            IntentApp []IntentApp  `json:"intent_app"`
            CardIsin string  `json:"card_isin"`
            CodLimit float64  `json:"cod_limit"`
            CardNumber string  `json:"card_number"`
            IntentAppErrorList []string  `json:"intent_app_error_list"`
            IntentAppErrorDictList []IntentAppErrorList  `json:"intent_app_error_dict_list"`
            ExpYear float64  `json:"exp_year"`
         
    }
    
    // RootPaymentMode used by Payment
    type RootPaymentMode struct {

        
            Name string  `json:"name"`
            SaveCard bool  `json:"save_card"`
            DisplayName string  `json:"display_name"`
            DisplayPriority float64  `json:"display_priority"`
            AggregatorName string  `json:"aggregator_name"`
            AnonymousEnable bool  `json:"anonymous_enable"`
            AddCardEnabled bool  `json:"add_card_enabled"`
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
    
    // PayoutsResponse used by Payment
    type PayoutsResponse struct {

        
            TransferType string  `json:"transfer_type"`
            IsActive bool  `json:"is_active"`
            PayoutsAggregators []map[string]interface{}  `json:"payouts_aggregators"`
            IsDefault bool  `json:"is_default"`
            MoreAttributes map[string]interface{}  `json:"more_attributes"`
            Customers map[string]interface{}  `json:"customers"`
            UniqueTransferNo map[string]interface{}  `json:"unique_transfer_no"`
         
    }
    
    // PayoutBankDetails used by Payment
    type PayoutBankDetails struct {

        
            City string  `json:"city"`
            AccountNo string  `json:"account_no"`
            BranchName string  `json:"branch_name"`
            Pincode float64  `json:"pincode"`
            IfscCode string  `json:"ifsc_code"`
            AccountHolder string  `json:"account_holder"`
            Country string  `json:"country"`
            AccountType string  `json:"account_type"`
            BankName string  `json:"bank_name"`
            State string  `json:"state"`
         
    }
    
    // PayoutRequest used by Payment
    type PayoutRequest struct {

        
            Users map[string]interface{}  `json:"users"`
            BankDetails PayoutBankDetails  `json:"bank_details"`
            TransferType string  `json:"transfer_type"`
            IsActive bool  `json:"is_active"`
            Aggregator string  `json:"aggregator"`
            UniqueExternalID string  `json:"unique_external_id"`
         
    }
    
    // PayoutResponse used by Payment
    type PayoutResponse struct {

        
            Users map[string]interface{}  `json:"users"`
            BankDetails map[string]interface{}  `json:"bank_details"`
            TransferType string  `json:"transfer_type"`
            Created bool  `json:"created"`
            IsActive bool  `json:"is_active"`
            Success bool  `json:"success"`
            Aggregator string  `json:"aggregator"`
            UniqueTransferNo string  `json:"unique_transfer_no"`
            PaymentStatus string  `json:"payment_status"`
            Payouts map[string]interface{}  `json:"payouts"`
         
    }
    
    // UpdatePayoutResponse used by Payment
    type UpdatePayoutResponse struct {

        
            Success bool  `json:"success"`
            IsActive bool  `json:"is_active"`
            IsDefault bool  `json:"is_default"`
         
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

        
            Success bool  `json:"success"`
            Description string  `json:"description"`
            Code string  `json:"code"`
         
    }
    
    // BankDetailsForOTP used by Payment
    type BankDetailsForOTP struct {

        
            AccountNo string  `json:"account_no"`
            BranchName string  `json:"branch_name"`
            IfscCode string  `json:"ifsc_code"`
            AccountHolder string  `json:"account_holder"`
            BankName string  `json:"bank_name"`
         
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

        
            AccountNo string  `json:"account_no"`
            BranchName string  `json:"branch_name"`
            Title string  `json:"title"`
            DisplayName string  `json:"display_name"`
            AccountHolder string  `json:"account_holder"`
            ID float64  `json:"id"`
            Address string  `json:"address"`
            Email string  `json:"email"`
            IfscCode string  `json:"ifsc_code"`
            DelightsUserName string  `json:"delights_user_name"`
            CreatedOn string  `json:"created_on"`
            IsActive bool  `json:"is_active"`
            Subtitle string  `json:"subtitle"`
            BeneficiaryID string  `json:"beneficiary_id"`
            BankName string  `json:"bank_name"`
            Comment string  `json:"comment"`
            ModifiedOn string  `json:"modified_on"`
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

        
            PaymentGateway string  `json:"payment_gateway"`
            OrderID string  `json:"order_id"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            CurrentStatus string  `json:"current_status"`
            PaymentID string  `json:"payment_id"`
         
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

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
            OrderID string  `json:"order_id"`
         
    }
    
    // CODdata used by Payment
    type CODdata struct {

        
            IsActive bool  `json:"is_active"`
            Limit float64  `json:"limit"`
            RemainingLimit float64  `json:"remaining_limit"`
            UserID string  `json:"user_id"`
            Usages float64  `json:"usages"`
         
    }
    
    // GetUserCODLimitResponse used by Payment
    type GetUserCODLimitResponse struct {

        
            Success bool  `json:"success"`
            UserCodData CODdata  `json:"user_cod_data"`
         
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
    

    
    // PlatformDeliveryAddress used by Order
    type PlatformDeliveryAddress struct {

        
            Address1 string  `json:"address1"`
            City string  `json:"city"`
            Phone string  `json:"phone"`
            Latitude float64  `json:"latitude"`
            Email string  `json:"email"`
            Country string  `json:"country"`
            AddressType string  `json:"address_type"`
            Landmark string  `json:"landmark"`
            State string  `json:"state"`
            Area string  `json:"area"`
            Address2 string  `json:"address2"`
            ContactPerson string  `json:"contact_person"`
            CreatedAt string  `json:"created_at"`
            AddressCategory string  `json:"address_category"`
            Version string  `json:"version"`
            Longitude float64  `json:"longitude"`
            UpdatedAt string  `json:"updated_at"`
            Pincode string  `json:"pincode"`
         
    }
    
    // ShipmentListingChannel used by Order
    type ShipmentListingChannel struct {

        
            ChannelShipmentID string  `json:"channel_shipment_id"`
            Logo string  `json:"logo"`
            Name string  `json:"name"`
            IsAffiliate bool  `json:"is_affiliate"`
         
    }
    
    // ShipmentItemFulFillingStore used by Order
    type ShipmentItemFulFillingStore struct {

        
            Meta map[string]interface{}  `json:"meta"`
            City string  `json:"city"`
            Phone string  `json:"phone"`
            ID float64  `json:"id"`
            Name string  `json:"name"`
            BrandStoreTags string  `json:"brand_store_tags"`
            Address string  `json:"address"`
            State string  `json:"state"`
            Code string  `json:"code"`
            LocationType string  `json:"location_type"`
            StoreEmail string  `json:"store_email"`
            Pincode string  `json:"pincode"`
         
    }
    
    // ReplacementDetails used by Order
    type ReplacementDetails struct {

        
            OriginalAffiliateOrderID string  `json:"original_affiliate_order_id"`
            ReplacementType string  `json:"replacement_type"`
         
    }
    
    // AffiliateMeta used by Order
    type AffiliateMeta struct {

        
            ReplacementDetails ReplacementDetails  `json:"replacement_details"`
            ChannelOrderID string  `json:"channel_order_id"`
            ChannelShipmentID string  `json:"channel_shipment_id"`
            MarketplaceInvoiceID string  `json:"marketplace_invoice_id"`
            Quantity float64  `json:"quantity"`
            IsPriority bool  `json:"is_priority"`
            EmployeeDiscount float64  `json:"employee_discount"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            BoxType string  `json:"box_type"`
            DueDate string  `json:"due_date"`
            CouponCode string  `json:"coupon_code"`
            OrderItemID string  `json:"order_item_id"`
            SizeLevelTotalQty float64  `json:"size_level_total_qty"`
         
    }
    
    // AffiliateBagDetails used by Order
    type AffiliateBagDetails struct {

        
            AffiliateMeta AffiliateMeta  `json:"affiliate_meta"`
            EmployeeDiscount float64  `json:"employee_discount"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
         
    }
    
    // BagReturnableCancelableStatus used by Order
    type BagReturnableCancelableStatus struct {

        
            CanBeCancelled bool  `json:"can_be_cancelled"`
            IsActive bool  `json:"is_active"`
            EnableTracking bool  `json:"enable_tracking"`
            IsCustomerReturnAllowed bool  `json:"is_customer_return_allowed"`
            IsReturnable bool  `json:"is_returnable"`
         
    }
    
    // BagStateMapper used by Order
    type BagStateMapper struct {

        
            AppStateName string  `json:"app_state_name"`
            StateType string  `json:"state_type"`
            ID float64  `json:"id"`
            DisplayName string  `json:"display_name"`
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            NotifyCustomer bool  `json:"notify_customer"`
            AppDisplayName string  `json:"app_display_name"`
            JourneyType string  `json:"journey_type"`
            AppFacing bool  `json:"app_facing"`
         
    }
    
    // BagStatusHistory used by Order
    type BagStatusHistory struct {

        
            Reasons []map[string]interface{}  `json:"reasons"`
            BshID float64  `json:"bsh_id"`
            StateType string  `json:"state_type"`
            BagID float64  `json:"bag_id"`
            Status string  `json:"status"`
            StoreID float64  `json:"store_id"`
            StateID float64  `json:"state_id"`
            ShipmentID string  `json:"shipment_id"`
            DisplayName string  `json:"display_name"`
            DeliveryAwbNumber string  `json:"delivery_awb_number"`
            KafkaSync bool  `json:"kafka_sync"`
            BagStateMapper BagStateMapper  `json:"bag_state_mapper"`
            CreatedAt string  `json:"created_at"`
            AppDisplayName string  `json:"app_display_name"`
            DeliveryPartnerID float64  `json:"delivery_partner_id"`
            UpdatedAt string  `json:"updated_at"`
            Forward bool  `json:"forward"`
         
    }
    
    // PlatformArticleAttributes used by Order
    type PlatformArticleAttributes struct {

        
            Currency string  `json:"currency"`
         
    }
    
    // PlatformItem used by Order
    type PlatformItem struct {

        
            ID float64  `json:"id"`
            L2Category []string  `json:"l2_category"`
            CanReturn bool  `json:"can_return"`
            LastUpdatedAt string  `json:"last_updated_at"`
            BrandID float64  `json:"brand_id"`
            Image []string  `json:"image"`
            L1Category []string  `json:"l1_category"`
            Attributes PlatformArticleAttributes  `json:"attributes"`
            SlugKey string  `json:"slug_key"`
            L3Category float64  `json:"l3_category"`
            Brand string  `json:"brand"`
            Code string  `json:"code"`
            BranchURL string  `json:"branch_url"`
            Meta map[string]interface{}  `json:"meta"`
            Images []string  `json:"images"`
            Size string  `json:"size"`
            DepartmentID float64  `json:"department_id"`
            L3CategoryName string  `json:"l3_category_name"`
            Name string  `json:"name"`
            CanCancel bool  `json:"can_cancel"`
            Color string  `json:"color"`
         
    }
    
    // GSTDetailsData used by Order
    type GSTDetailsData struct {

        
            IgstTaxPercentage float64  `json:"igst_tax_percentage"`
            ValueOfGood float64  `json:"value_of_good"`
            CgstTaxPercentage float64  `json:"cgst_tax_percentage"`
            SgstTaxPercentage float64  `json:"sgst_tax_percentage"`
            IgstGstFee string  `json:"igst_gst_fee"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            IsDefaultHsnCode bool  `json:"is_default_hsn_code"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            HsnCode string  `json:"hsn_code"`
            SgstGstFee string  `json:"sgst_gst_fee"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            HsnCodeID string  `json:"hsn_code_id"`
            GstinCode string  `json:"gstin_code"`
            CgstGstFee string  `json:"cgst_gst_fee"`
            GstFee float64  `json:"gst_fee"`
            GstTag string  `json:"gst_tag"`
         
    }
    
    // Prices used by Order
    type Prices struct {

        
            ValueOfGood float64  `json:"value_of_good"`
            Discount float64  `json:"discount"`
            CodCharges float64  `json:"cod_charges"`
            CouponValue float64  `json:"coupon_value"`
            PriceEffective float64  `json:"price_effective"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            PromotionEffectiveDiscount float64  `json:"promotion_effective_discount"`
            PmPriceSplit float64  `json:"pm_price_split"`
            AmountPaid float64  `json:"amount_paid"`
            RefundCredit float64  `json:"refund_credit"`
            DeliveryCharge float64  `json:"delivery_charge"`
            TransferPrice float64  `json:"transfer_price"`
            FyndCredits float64  `json:"fynd_credits"`
            PriceMarked float64  `json:"price_marked"`
            RefundAmount float64  `json:"refund_amount"`
            CashbackApplied float64  `json:"cashback_applied"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            Cashback float64  `json:"cashback"`
            GiftPrice float64  `json:"gift_price"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            AmountPaidRoundoff float64  `json:"amount_paid_roundoff"`
         
    }
    
    // ShipmentListingBrand used by Order
    type ShipmentListingBrand struct {

        
            Logo string  `json:"logo"`
            CreatedOn string  `json:"created_on"`
            LogoBase64 string  `json:"logo_base64"`
            Name string  `json:"name"`
         
    }
    
    // Identifier used by Order
    type Identifier struct {

        
            Ean string  `json:"ean"`
            Upc string  `json:"upc"`
            Isbn string  `json:"isbn"`
            Alu string  `json:"alu"`
            SkuCode string  `json:"sku_code"`
         
    }
    
    // FinancialBreakup used by Order
    type FinancialBreakup struct {

        
            ValueOfGood float64  `json:"value_of_good"`
            Discount float64  `json:"discount"`
            HsnCode string  `json:"hsn_code"`
            CodCharges float64  `json:"cod_charges"`
            GstTag string  `json:"gst_tag"`
            CouponValue float64  `json:"coupon_value"`
            PriceEffective float64  `json:"price_effective"`
            TotalUnits float64  `json:"total_units"`
            ItemName string  `json:"item_name"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            PromotionEffectiveDiscount float64  `json:"promotion_effective_discount"`
            AmountPaid float64  `json:"amount_paid"`
            RefundCredit float64  `json:"refund_credit"`
            DeliveryCharge float64  `json:"delivery_charge"`
            TransferPrice float64  `json:"transfer_price"`
            FyndCredits float64  `json:"fynd_credits"`
            PriceMarked float64  `json:"price_marked"`
            CashbackApplied float64  `json:"cashback_applied"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            Size string  `json:"size"`
            Cashback float64  `json:"cashback"`
            Identifiers Identifier  `json:"identifiers"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            AmountPaidRoundoff float64  `json:"amount_paid_roundoff"`
            GstFee float64  `json:"gst_fee"`
            AddedToFyndCash bool  `json:"added_to_fynd_cash"`
         
    }
    
    // Dates used by Order
    type Dates struct {

        
            OrderCreated string  `json:"order_created"`
            DeliveryDate string  `json:"delivery_date"`
         
    }
    
    // ReturnConfig used by Order
    type ReturnConfig struct {

        
            Unit string  `json:"unit"`
            Returnable bool  `json:"returnable"`
            Time float64  `json:"time"`
         
    }
    
    // Weight used by Order
    type Weight struct {

        
            Unit string  `json:"unit"`
            IsDefault bool  `json:"is_default"`
            Shipping float64  `json:"shipping"`
         
    }
    
    // Dimensions used by Order
    type Dimensions struct {

        
            Length float64  `json:"length"`
            IsDefault bool  `json:"is_default"`
            Unit string  `json:"unit"`
            Width float64  `json:"width"`
            Height float64  `json:"height"`
         
    }
    
    // Article used by Order
    type Article struct {

        
            SellerIdentifier string  `json:"seller_identifier"`
            Size string  `json:"size"`
            UID string  `json:"uid"`
            ReturnConfig ReturnConfig  `json:"return_config"`
            Identifiers map[string]interface{}  `json:"identifiers"`
            EspModified bool  `json:"esp_modified"`
            IsSet bool  `json:"is_set"`
            ASet map[string]interface{}  `json:"a_set"`
            RawMeta string  `json:"raw_meta"`
            Weight Weight  `json:"weight"`
            Code string  `json:"code"`
            ID string  `json:"_id"`
            Dimensions Dimensions  `json:"dimensions"`
            Currency map[string]interface{}  `json:"currency"`
            ChildDetails map[string]interface{}  `json:"child_details"`
         
    }
    
    // BagUnit used by Order
    type BagUnit struct {

        
            AffiliateBagDetails AffiliateBagDetails  `json:"affiliate_bag_details"`
            Status BagReturnableCancelableStatus  `json:"status"`
            DisplayName string  `json:"display_name"`
            LineNumber float64  `json:"line_number"`
            BagStatus []BagStatusHistory  `json:"bag_status"`
            CanReturn bool  `json:"can_return"`
            BagType string  `json:"bag_type"`
            Item PlatformItem  `json:"item"`
            ProductQuantity float64  `json:"product_quantity"`
            Gst GSTDetailsData  `json:"gst"`
            CurrentOperationalStatus BagStatusHistory  `json:"current_operational_status"`
            Prices Prices  `json:"prices"`
            BagID float64  `json:"bag_id"`
            Brand ShipmentListingBrand  `json:"brand"`
            FinancialBreakup []FinancialBreakup  `json:"financial_breakup"`
            Dates Dates  `json:"dates"`
            Meta map[string]interface{}  `json:"meta"`
            Reasons []map[string]interface{}  `json:"reasons"`
            Size string  `json:"size"`
            Article Article  `json:"article"`
            BagExpiryDate string  `json:"bag_expiry_date"`
            CanCancel bool  `json:"can_cancel"`
            CurrentStatus BagStatusHistory  `json:"current_status"`
            EntityType string  `json:"entity_type"`
         
    }
    
    // ShipmentStatus used by Order
    type ShipmentStatus struct {

        
            Meta map[string]interface{}  `json:"meta"`
            StatusCreatedAt string  `json:"status_created_at"`
            Status string  `json:"status"`
            ShipmentID string  `json:"shipment_id"`
            ShipmentStatusID float64  `json:"shipment_status_id"`
            CreatedAt string  `json:"created_at"`
            BagList []string  `json:"bag_list"`
            Title string  `json:"title"`
            CurrentShipmentStatus string  `json:"current_shipment_status"`
         
    }
    
    // UserDataInfo used by Order
    type UserDataInfo struct {

        
            Mobile string  `json:"mobile"`
            LastName string  `json:"last_name"`
            UID float64  `json:"uid"`
            Email string  `json:"email"`
            Gender string  `json:"gender"`
            Name string  `json:"name"`
            AvisUserID string  `json:"avis_user_id"`
            IsAnonymousUser bool  `json:"is_anonymous_user"`
            FirstName string  `json:"first_name"`
         
    }
    
    // Formatted used by Order
    type Formatted struct {

        
            Max string  `json:"max"`
            Min string  `json:"min"`
         
    }
    
    // ShipmentTags used by Order
    type ShipmentTags struct {

        
            DisplayText string  `json:"display_text"`
            EntityType string  `json:"entity_type"`
            Slug string  `json:"slug"`
         
    }
    
    // ShipmentTimeStamp used by Order
    type ShipmentTimeStamp struct {

        
            Max float64  `json:"max"`
            Min float64  `json:"min"`
         
    }
    
    // LockData used by Order
    type LockData struct {

        
            LockMessage string  `json:"lock_message"`
            Locked bool  `json:"locked"`
            Mto bool  `json:"mto"`
         
    }
    
    // ShipmentItemMeta used by Order
    type ShipmentItemMeta struct {

        
            AssignDpFromSb bool  `json:"assign_dp_from_sb"`
            ExistingDpList []string  `json:"existing_dp_list"`
            FulfilmentPriorityText string  `json:"fulfilment_priority_text"`
            ParentDpID string  `json:"parent_dp_id"`
            Weight float64  `json:"weight"`
            DebugInfo map[string]interface{}  `json:"debug_info"`
            OrderType string  `json:"order_type"`
            StoreInvoiceUpdatedDate string  `json:"store_invoice_updated_date"`
            ActivityComment string  `json:"activity_comment"`
            PdfMedia []map[string]interface{}  `json:"pdf_media"`
            Formatted Formatted  `json:"formatted"`
            ShipmentWeight float64  `json:"shipment_weight"`
            IsInternational bool  `json:"is_international"`
            ShipmentVolumetricWeight float64  `json:"shipment_volumetric_weight"`
            AutoTriggerDpAssignmentAcf bool  `json:"auto_trigger_dp_assignment_acf"`
            Sla float64  `json:"sla"`
            EwaybillInfo map[string]interface{}  `json:"ewaybill_info"`
            BagWeight map[string]interface{}  `json:"bag_weight"`
            ShippingZone string  `json:"shipping_zone"`
            SameStoreAvailable bool  `json:"same_store_available"`
            ShipmentTags []ShipmentTags  `json:"shipment_tags"`
            Timestamp ShipmentTimeStamp  `json:"timestamp"`
            LockData LockData  `json:"lock_data"`
            DpSortKey string  `json:"dp_sort_key"`
            PackagingName string  `json:"packaging_name"`
            ShipmentChargeableWeight float64  `json:"shipment_chargeable_weight"`
            Tags []map[string]interface{}  `json:"tags"`
            External map[string]interface{}  `json:"external"`
            DpOptions map[string]interface{}  `json:"dp_options"`
         
    }
    
    // ShipmentItem used by Order
    type ShipmentItem struct {

        
            DisplayName string  `json:"display_name"`
            PreviousShipmentID string  `json:"previous_shipment_id"`
            DeliveryAddress PlatformDeliveryAddress  `json:"delivery_address"`
            OrderID string  `json:"order_id"`
            Channel ShipmentListingChannel  `json:"channel"`
            ShipmentCreatedAt string  `json:"shipment_created_at"`
            FulfillingStore ShipmentItemFulFillingStore  `json:"fulfilling_store"`
            LockStatus bool  `json:"lock_status"`
            TotalBags float64  `json:"total_bags"`
            InvoiceID string  `json:"invoice_id"`
            Bags []BagUnit  `json:"bags"`
            PaymentMode string  `json:"payment_mode"`
            Prices Prices  `json:"prices"`
            OrderingChannnel string  `json:"ordering_channnel"`
            ShipmentStatus ShipmentStatus  `json:"shipment_status"`
            ShipmentID string  `json:"shipment_id"`
            PaymentMethods map[string]interface{}  `json:"payment_methods"`
            User UserDataInfo  `json:"user"`
            Meta ShipmentItemMeta  `json:"meta"`
            StatusCreatedAt string  `json:"status_created_at"`
            CanProcess bool  `json:"can_process"`
            CustomerNote string  `json:"customer_note"`
            OrderDate string  `json:"order_date"`
            EstimatedSlaTime string  `json:"estimated_sla_time"`
         
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
            Success bool  `json:"success"`
            Page Page  `json:"page"`
            Message string  `json:"message"`
            Lane string  `json:"lane"`
            TotalCount float64  `json:"total_count"`
         
    }
    
    // Error used by Order
    type Error struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // ShipmentPayments used by Order
    type ShipmentPayments struct {

        
            Logo string  `json:"logo"`
            Mode string  `json:"mode"`
            Source string  `json:"source"`
         
    }
    
    // UserDetailsData used by Order
    type UserDetailsData struct {

        
            Address1 string  `json:"address1"`
            Phone string  `json:"phone"`
            City string  `json:"city"`
            Email string  `json:"email"`
            Country string  `json:"country"`
            AddressType string  `json:"address_type"`
            Landmark string  `json:"landmark"`
            Name string  `json:"name"`
            State string  `json:"state"`
            Address string  `json:"address"`
            Area string  `json:"area"`
            Pincode string  `json:"pincode"`
         
    }
    
    // ShipmentStatusData used by Order
    type ShipmentStatusData struct {

        
            Meta map[string]interface{}  `json:"meta"`
            Status string  `json:"status"`
            ID float64  `json:"id"`
            ShipmentID string  `json:"shipment_id"`
            DisplayName string  `json:"display_name"`
            CreatedAt string  `json:"created_at"`
            BagList []string  `json:"bag_list"`
         
    }
    
    // FulfillingStore used by Order
    type FulfillingStore struct {

        
            Meta map[string]interface{}  `json:"meta"`
            Phone string  `json:"phone"`
            City string  `json:"city"`
            StoreName string  `json:"store_name"`
            ID float64  `json:"id"`
            FulfillmentChannel string  `json:"fulfillment_channel"`
            Country string  `json:"country"`
            State string  `json:"state"`
            Address string  `json:"address"`
            ContactPerson string  `json:"contact_person"`
            Code string  `json:"code"`
            Pincode string  `json:"pincode"`
         
    }
    
    // ShipmentDetails used by Order
    type ShipmentDetails struct {

        
            LockStatus bool  `json:"lock_status"`
            ActionToStatus map[string]interface{}  `json:"action_to_status"`
            LockMessage string  `json:"lock_message"`
         
    }
    
    // OrderDetailsData used by Order
    type OrderDetailsData struct {

        
            OrderingChannelLogo map[string]interface{}  `json:"ordering_channel_logo"`
            Source string  `json:"source"`
            AffiliateID string  `json:"affiliate_id"`
            OrderValue string  `json:"order_value"`
            FyndOrderID string  `json:"fynd_order_id"`
            OrderDate string  `json:"order_date"`
            OrderingChannel string  `json:"ordering_channel"`
            CodCharges string  `json:"cod_charges"`
            TaxDetails map[string]interface{}  `json:"tax_details"`
         
    }
    
    // AffiliateBagsDetails used by Order
    type AffiliateBagsDetails struct {

        
            CouponCode string  `json:"coupon_code"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
         
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

        
            MrpPromotion bool  `json:"mrp_promotion"`
            PromoID string  `json:"promo_id"`
            DiscountRules []DiscountRules  `json:"discount_rules"`
            BuyRules []BuyRules  `json:"buy_rules"`
            PromotionName string  `json:"promotion_name"`
            ArticleQuantity float64  `json:"article_quantity"`
            PromotionType string  `json:"promotion_type"`
            Amount float64  `json:"amount"`
         
    }
    
    // OrderBrandName used by Order
    type OrderBrandName struct {

        
            Company float64  `json:"company"`
            CreatedOn string  `json:"created_on"`
            BrandName string  `json:"brand_name"`
            ID float64  `json:"id"`
            Logo string  `json:"logo"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // B2BPODetails used by Order
    type B2BPODetails struct {

        
            DocketNumber string  `json:"docket_number"`
            ItemBasePrice float64  `json:"item_base_price"`
            PartialCanRet bool  `json:"partial_can_ret"`
            PoTaxAmount float64  `json:"po_tax_amount"`
            PoLineAmount float64  `json:"po_line_amount"`
            TotalGstPercentage float64  `json:"total_gst_percentage"`
         
    }
    
    // GiftCard used by Order
    type GiftCard struct {

        
            DisplayText string  `json:"display_text"`
            GiftMessage string  `json:"gift_message"`
            IsGiftApplied bool  `json:"is_gift_applied"`
            GiftPrice float64  `json:"gift_price"`
         
    }
    
    // BagMeta used by Order
    type BagMeta struct {

        
            CustomJson map[string]interface{}  `json:"custom_json"`
            CustomMessage string  `json:"custom_message"`
            B2bPoDetails B2BPODetails  `json:"b2b_po_details"`
            DocketNumber string  `json:"docket_number"`
            GroupID string  `json:"group_id"`
            GiftCard GiftCard  `json:"gift_card"`
            PartialCanRet bool  `json:"partial_can_ret"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
         
    }
    
    // ReturnConfig1 used by Order
    type ReturnConfig1 struct {

        
            Unit string  `json:"unit"`
            Returnable bool  `json:"returnable"`
            Time float64  `json:"time"`
         
    }
    
    // OrderBagArticle used by Order
    type OrderBagArticle struct {

        
            Size string  `json:"size"`
            UID string  `json:"uid"`
            ReturnConfig ReturnConfig1  `json:"return_config"`
            Identifiers map[string]interface{}  `json:"identifiers"`
         
    }
    
    // BagConfigs used by Order
    type BagConfigs struct {

        
            CanBeCancelled bool  `json:"can_be_cancelled"`
            AllowForceReturn bool  `json:"allow_force_return"`
            IsActive bool  `json:"is_active"`
            EnableTracking bool  `json:"enable_tracking"`
            IsCustomerReturnAllowed bool  `json:"is_customer_return_allowed"`
            IsReturnable bool  `json:"is_returnable"`
         
    }
    
    // CurrentStatus used by Order
    type CurrentStatus struct {

        
            StateType string  `json:"state_type"`
            BagID float64  `json:"bag_id"`
            Status string  `json:"status"`
            ID float64  `json:"id"`
            StoreID float64  `json:"store_id"`
            StateID float64  `json:"state_id"`
            ShipmentID string  `json:"shipment_id"`
            DeliveryAwbNumber string  `json:"delivery_awb_number"`
            KafkaSync bool  `json:"kafka_sync"`
            BagStateMapper BagStateMapper  `json:"bag_state_mapper"`
            CreatedAt string  `json:"created_at"`
            DeliveryPartnerID float64  `json:"delivery_partner_id"`
            UpdatedAt float64  `json:"updated_at"`
         
    }
    
    // BagGST used by Order
    type BagGST struct {

        
            IgstTaxPercentage float64  `json:"igst_tax_percentage"`
            ValueOfGood float64  `json:"value_of_good"`
            CgstTaxPercentage float64  `json:"cgst_tax_percentage"`
            SgstTaxPercentage float64  `json:"sgst_tax_percentage"`
            IsDefaultHsnCode bool  `json:"is_default_hsn_code"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            IgstGstFee string  `json:"igst_gst_fee"`
            HsnCode string  `json:"hsn_code"`
            SgstGstFee string  `json:"sgst_gst_fee"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            HsnCodeID string  `json:"hsn_code_id"`
            GstinCode string  `json:"gstin_code"`
            CgstGstFee string  `json:"cgst_gst_fee"`
            GstFee float64  `json:"gst_fee"`
            GstTag string  `json:"gst_tag"`
         
    }
    
    // OrderBags used by Order
    type OrderBags struct {

        
            AffiliateBagDetails AffiliateBagsDetails  `json:"affiliate_bag_details"`
            DisplayName string  `json:"display_name"`
            ParentPromoBags map[string]interface{}  `json:"parent_promo_bags"`
            LineNumber float64  `json:"line_number"`
            DeliveryAddress PlatformDeliveryAddress  `json:"delivery_address"`
            CanReturn bool  `json:"can_return"`
            Quantity float64  `json:"quantity"`
            Identifier string  `json:"identifier"`
            Item PlatformItem  `json:"item"`
            GroupID string  `json:"group_id"`
            Prices Prices  `json:"prices"`
            BagID float64  `json:"bag_id"`
            AppliedPromos []AppliedPromos  `json:"applied_promos"`
            Brand OrderBrandName  `json:"brand"`
            FinancialBreakup FinancialBreakup  `json:"financial_breakup"`
            Meta BagMeta  `json:"meta"`
            SellerIdentifier string  `json:"seller_identifier"`
            Article OrderBagArticle  `json:"article"`
            BagConfigs BagConfigs  `json:"bag_configs"`
            CanCancel bool  `json:"can_cancel"`
            CurrentStatus CurrentStatus  `json:"current_status"`
            EntityType string  `json:"entity_type"`
            GstDetails BagGST  `json:"gst_details"`
         
    }
    
    // OrderingStoreDetails used by Order
    type OrderingStoreDetails struct {

        
            Meta map[string]interface{}  `json:"meta"`
            Phone string  `json:"phone"`
            City string  `json:"city"`
            StoreName string  `json:"store_name"`
            ID float64  `json:"id"`
            Country string  `json:"country"`
            State string  `json:"state"`
            Address string  `json:"address"`
            ContactPerson string  `json:"contact_person"`
            Code string  `json:"code"`
            Pincode string  `json:"pincode"`
         
    }
    
    // InvoiceInfo used by Order
    type InvoiceInfo struct {

        
            ExternalInvoiceID string  `json:"external_invoice_id"`
            UpdatedDate string  `json:"updated_date"`
            CreditNoteID string  `json:"credit_note_id"`
            LabelURL string  `json:"label_url"`
            InvoiceURL string  `json:"invoice_url"`
            StoreInvoiceID string  `json:"store_invoice_id"`
         
    }
    
    // DPDetailsData used by Order
    type DPDetailsData struct {

        
            EwayBillID string  `json:"eway_bill_id"`
            ID float64  `json:"id"`
            TrackURL string  `json:"track_url"`
            AwbNo string  `json:"awb_no"`
            Country string  `json:"country"`
            Name string  `json:"name"`
            GstTag string  `json:"gst_tag"`
            Pincode string  `json:"pincode"`
         
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

        
            City string  `json:"city"`
            Gstin string  `json:"gstin"`
            AjioSiteID string  `json:"ajio_site_id"`
            Name string  `json:"name"`
            State string  `json:"state"`
            Address string  `json:"address"`
            Pincode float64  `json:"pincode"`
         
    }
    
    // ShipmentMeta used by Order
    type ShipmentMeta struct {

        
            AssignDpFromSb bool  `json:"assign_dp_from_sb"`
            ForwardAffiliateShipmentID string  `json:"forward_affiliate_shipment_id"`
            FulfilmentPriorityText string  `json:"fulfilment_priority_text"`
            EinvoiceInfo EinvoiceInfo  `json:"einvoice_info"`
            ParentDpID string  `json:"parent_dp_id"`
            Weight float64  `json:"weight"`
            DebugInfo DebugInfo  `json:"debug_info"`
            B2bBuyerDetails BuyerDetails  `json:"b2b_buyer_details"`
            ForwardAffiliateOrderID string  `json:"forward_affiliate_order_id"`
            OrderType string  `json:"order_type"`
            StoreInvoiceUpdatedDate string  `json:"store_invoice_updated_date"`
            DueDate string  `json:"due_date"`
            ReturnDetails map[string]interface{}  `json:"return_details"`
            ReturnAwbNumber string  `json:"return_awb_number"`
            Formatted Formatted  `json:"formatted"`
            ShipmentWeight float64  `json:"shipment_weight"`
            ShipmentVolumetricWeight float64  `json:"shipment_volumetric_weight"`
            B2cBuyerDetails map[string]interface{}  `json:"b2c_buyer_details"`
            AutoTriggerDpAssignmentAcf bool  `json:"auto_trigger_dp_assignment_acf"`
            ReturnAffiliateOrderID string  `json:"return_affiliate_order_id"`
            Dimension Dimensions  `json:"dimension"`
            EwaybillInfo map[string]interface{}  `json:"ewaybill_info"`
            BagWeight map[string]interface{}  `json:"bag_weight"`
            PoNumber string  `json:"po_number"`
            SameStoreAvailable bool  `json:"same_store_available"`
            ShipmentTags []ShipmentTags  `json:"shipment_tags"`
            Timestamp ShipmentTimeStamp  `json:"timestamp"`
            ReturnAffiliateShipmentID string  `json:"return_affiliate_shipment_id"`
            ReturnStoreNode float64  `json:"return_store_node"`
            LockData LockData  `json:"lock_data"`
            DpName string  `json:"dp_name"`
            DpSortKey string  `json:"dp_sort_key"`
            PackagingName string  `json:"packaging_name"`
            AwbNumber string  `json:"awb_number"`
            MarketplaceStoreID string  `json:"marketplace_store_id"`
            BoxType string  `json:"box_type"`
            External map[string]interface{}  `json:"external"`
            DpOptions map[string]interface{}  `json:"dp_options"`
            DpID string  `json:"dp_id"`
         
    }
    
    // TrackingList used by Order
    type TrackingList struct {

        
            Status string  `json:"status"`
            Text string  `json:"text"`
            Time string  `json:"time"`
            IsPassed bool  `json:"is_passed"`
            IsCurrent bool  `json:"is_current"`
         
    }
    
    // PDFLinks used by Order
    type PDFLinks struct {

        
            Invoice string  `json:"invoice"`
            LabelType string  `json:"label_type"`
            CreditNoteURL string  `json:"credit_note_url"`
            InvoiceA6 string  `json:"invoice_a6"`
            LabelExport string  `json:"label_export"`
            InvoiceA4 string  `json:"invoice_a4"`
            LabelA4 string  `json:"label_a4"`
            Label string  `json:"label"`
            B2b string  `json:"b2b"`
            InvoiceType string  `json:"invoice_type"`
            DeliveryChallanA4 string  `json:"delivery_challan_a4"`
            PoInvoice string  `json:"po_invoice"`
            LabelA6 string  `json:"label_a6"`
            InvoiceExport string  `json:"invoice_export"`
            LabelPos string  `json:"label_pos"`
            InvoicePos string  `json:"invoice_pos"`
         
    }
    
    // AffiliateDetails used by Order
    type AffiliateDetails struct {

        
            AffiliateMeta AffiliateMeta  `json:"affiliate_meta"`
            AffiliateID string  `json:"affiliate_id"`
            ShipmentMeta ShipmentMeta  `json:"shipment_meta"`
            AffiliateStoreID string  `json:"affiliate_store_id"`
            PdfLinks PDFLinks  `json:"pdf_links"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            CompanyAffiliateTag string  `json:"company_affiliate_tag"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            AdID string  `json:"ad_id"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
         
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

        
            CompanyCin string  `json:"company_cin"`
            CompanyName string  `json:"company_name"`
            Address map[string]interface{}  `json:"address"`
            CompanyContact ContactDetails  `json:"company_contact"`
            CompanyGst string  `json:"company_gst"`
            CompanyID float64  `json:"company_id"`
         
    }
    
    // PlatformShipment used by Order
    type PlatformShipment struct {

        
            Payments ShipmentPayments  `json:"payments"`
            BillingDetails UserDetailsData  `json:"billing_details"`
            Status ShipmentStatusData  `json:"status"`
            TotalItems float64  `json:"total_items"`
            CustomMeta []map[string]interface{}  `json:"custom_meta"`
            PriorityText string  `json:"priority_text"`
            DpAssignment bool  `json:"dp_assignment"`
            OperationalStatus string  `json:"operational_status"`
            ShipmentCreatedAt string  `json:"shipment_created_at"`
            FulfillingStore FulfillingStore  `json:"fulfilling_store"`
            ShipmentDetails ShipmentDetails  `json:"shipment_details"`
            PackagingType string  `json:"packaging_type"`
            ForwardShipmentID string  `json:"forward_shipment_id"`
            TotalBags float64  `json:"total_bags"`
            Order OrderDetailsData  `json:"order"`
            Vertical string  `json:"vertical"`
            EnableDpTracking bool  `json:"enable_dp_tracking"`
            InvoiceID string  `json:"invoice_id"`
            LockStatus bool  `json:"lock_status"`
            Bags []OrderBags  `json:"bags"`
            PickedDate string  `json:"picked_date"`
            OrderingStore OrderingStoreDetails  `json:"ordering_store"`
            PaymentMode string  `json:"payment_mode"`
            Prices Prices  `json:"prices"`
            ShipmentStatus string  `json:"shipment_status"`
            Invoice InvoiceInfo  `json:"invoice"`
            UserAgent string  `json:"user_agent"`
            ShipmentImages []string  `json:"shipment_images"`
            ShipmentQuantity float64  `json:"shipment_quantity"`
            ShipmentID string  `json:"shipment_id"`
            PaymentMethods map[string]interface{}  `json:"payment_methods"`
            DeliveryDetails UserDetailsData  `json:"delivery_details"`
            DpDetails DPDetailsData  `json:"dp_details"`
            User UserDataInfo  `json:"user"`
            Meta ShipmentMeta  `json:"meta"`
            TrackingList []TrackingList  `json:"tracking_list"`
            FulfilmentPriority float64  `json:"fulfilment_priority"`
            CustomMessage string  `json:"custom_message"`
            CanUpdateDimension bool  `json:"can_update_dimension"`
            AffiliateDetails AffiliateDetails  `json:"affiliate_details"`
            BagStatusHistory []BagStatusHistory  `json:"bag_status_history"`
            IsDpAssignEnabled bool  `json:"is_dp_assign_enabled"`
            DeliverySlot map[string]interface{}  `json:"delivery_slot"`
            PlatformLogo string  `json:"platform_logo"`
            JourneyType string  `json:"journey_type"`
            CompanyDetails CompanyDetails  `json:"company_details"`
            Coupon map[string]interface{}  `json:"coupon"`
            EstimatedSlaTime string  `json:"estimated_sla_time"`
            GstDetails GSTDetailsData  `json:"gst_details"`
         
    }
    
    // ShipmentInfoResponse used by Order
    type ShipmentInfoResponse struct {

        
            Shipments []PlatformShipment  `json:"shipments"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // PlatformUserDetails used by Order
    type PlatformUserDetails struct {

        
            PlatformUserFirstName string  `json:"platform_user_first_name"`
            PlatformUserLastName string  `json:"platform_user_last_name"`
            PlatformUserEmployeeCode string  `json:"platform_user_employee_code"`
            PlatformUserID string  `json:"platform_user_id"`
         
    }
    
    // BillingStaffDetails used by Order
    type BillingStaffDetails struct {

        
            StaffID float64  `json:"staff_id"`
            LastName string  `json:"last_name"`
            EmployeeCode string  `json:"employee_code"`
            User string  `json:"user"`
            FirstName string  `json:"first_name"`
         
    }
    
    // TransactionData used by Order
    type TransactionData struct {

        
            Status string  `json:"status"`
            TerminalID string  `json:"terminal_id"`
            TransactionID string  `json:"transaction_id"`
            AmountPaid float64  `json:"amount_paid"`
            PaymentID string  `json:"payment_id"`
            UniqueReferenceNumber string  `json:"unique_reference_number"`
            Currency string  `json:"currency"`
            Entity string  `json:"entity"`
         
    }
    
    // OrderMeta used by Order
    type OrderMeta struct {

        
            OrderType string  `json:"order_type"`
            OrderChildEntities []string  `json:"order_child_entities"`
            Comment string  `json:"comment"`
            PlatformUserDetails PlatformUserDetails  `json:"platform_user_details"`
            PaymentType string  `json:"payment_type"`
            OrderingStore float64  `json:"ordering_store"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            CompanyLogo string  `json:"company_logo"`
            OrderPlatform string  `json:"order_platform"`
            CurrencySymbol string  `json:"currency_symbol"`
            Staff map[string]interface{}  `json:"staff"`
            Files []map[string]interface{}  `json:"files"`
            BillingStaffDetails BillingStaffDetails  `json:"billing_staff_details"`
            CartID float64  `json:"cart_id"`
            MongoCartID float64  `json:"mongo_cart_id"`
            OrderTags []map[string]interface{}  `json:"order_tags"`
            CustomerNote string  `json:"customer_note"`
            TransactionData TransactionData  `json:"transaction_data"`
            EmployeeID float64  `json:"employee_id"`
         
    }
    
    // TaxDetails used by Order
    type TaxDetails struct {

        
            Gstin string  `json:"gstin"`
            PanNo string  `json:"pan_no"`
         
    }
    
    // OrderData used by Order
    type OrderData struct {

        
            Meta OrderMeta  `json:"meta"`
            PaymentMethods map[string]interface{}  `json:"payment_methods"`
            FyndOrderID string  `json:"fynd_order_id"`
            OrderDate string  `json:"order_date"`
            TaxDetails TaxDetails  `json:"tax_details"`
            Prices Prices  `json:"prices"`
         
    }
    
    // OrderDetailsResponse used by Order
    type OrderDetailsResponse struct {

        
            Shipments []PlatformShipment  `json:"shipments"`
            Success bool  `json:"success"`
            Order OrderData  `json:"order"`
         
    }
    
    // SubLane used by Order
    type SubLane struct {

        
            Text string  `json:"text"`
            TotalItems float64  `json:"total_items"`
            Index float64  `json:"index"`
            Actions []map[string]interface{}  `json:"actions"`
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

        
            Meta map[string]interface{}  `json:"meta"`
            OrderCreatedTime string  `json:"order_created_time"`
            OrderValue float64  `json:"order_value"`
            Shipments []PlatformShipment  `json:"shipments"`
            UserInfo UserDataInfo  `json:"user_info"`
            OrderID string  `json:"order_id"`
            TotalOrderValue float64  `json:"total_order_value"`
            BreakupValues []PlatformBreakupValues  `json:"breakup_values"`
            Channel PlatformChannel  `json:"channel"`
            PaymentMode string  `json:"payment_mode"`
         
    }
    
    // OrderListingResponse used by Order
    type OrderListingResponse struct {

        
            Items []PlatformOrderItems  `json:"items"`
            Success bool  `json:"success"`
            Page Page  `json:"page"`
            Message string  `json:"message"`
            Lane string  `json:"lane"`
            TotalCount float64  `json:"total_count"`
         
    }
    
    // PlatformTrack used by Order
    type PlatformTrack struct {

        
            Meta map[string]interface{}  `json:"meta"`
            ShipmentType string  `json:"shipment_type"`
            UpdatedTime string  `json:"updated_time"`
            Status string  `json:"status"`
            AccountName string  `json:"account_name"`
            RawStatus string  `json:"raw_status"`
            Reason string  `json:"reason"`
            LastLocationRecievedAt string  `json:"last_location_recieved_at"`
            Awb string  `json:"awb"`
            UpdatedAt string  `json:"updated_at"`
         
    }
    
    // PlatformShipmentTrack used by Order
    type PlatformShipmentTrack struct {

        
            Meta map[string]interface{}  `json:"meta"`
            Results []PlatformTrack  `json:"results"`
         
    }
    
    // FilterInfoOption used by Order
    type FilterInfoOption struct {

        
            PlaceholderText string  `json:"placeholder_text"`
            Text string  `json:"text"`
            Name string  `json:"name"`
            MinSearchSize float64  `json:"min_search_size"`
            ShowUI bool  `json:"show_ui"`
            Value string  `json:"value"`
         
    }
    
    // FiltersInfo used by Order
    type FiltersInfo struct {

        
            PlaceholderText string  `json:"placeholder_text"`
            Text string  `json:"text"`
            Type string  `json:"type"`
            Required bool  `json:"required"`
            Value string  `json:"value"`
            Options []FilterInfoOption  `json:"options"`
         
    }
    
    // AdvanceFilterInfo used by Order
    type AdvanceFilterInfo struct {

        
            Returned []FiltersInfo  `json:"returned"`
            Processed []FiltersInfo  `json:"processed"`
            Filters []FiltersInfo  `json:"filters"`
            AppliedFilters map[string]interface{}  `json:"applied_filters"`
            Unfulfilled []FiltersInfo  `json:"unfulfilled"`
            Page map[string]interface{}  `json:"page"`
            ActionCentre []FiltersInfo  `json:"action_centre"`
         
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

        
            Cdn URL  `json:"cdn"`
            FileName string  `json:"file_name"`
         
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

        
            ID float64  `json:"id"`
            DisplayName string  `json:"display_name"`
            QcType []string  `json:"qc_type"`
            QuestionSet []QuestionSet  `json:"question_set"`
         
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
            Reasons []ShipmentResponseReasons  `json:"reasons"`
            Message string  `json:"message"`
         
    }
    
    // BagReturnableCancelableStatus1 used by Order
    type BagReturnableCancelableStatus1 struct {

        
            CanBeCancelled bool  `json:"can_be_cancelled"`
            IsActive bool  `json:"is_active"`
            EnableTracking bool  `json:"enable_tracking"`
            IsCustomerReturnAllowed bool  `json:"is_customer_return_allowed"`
            IsReturnable bool  `json:"is_returnable"`
         
    }
    
    // Attributes used by Order
    type Attributes struct {

        
            PrimaryColor string  `json:"primary_color"`
            BrandName string  `json:"brand_name"`
            Gender []string  `json:"gender"`
            Name string  `json:"name"`
            MarketerAddress string  `json:"marketer_address"`
            MarketerName string  `json:"marketer_name"`
            Essential string  `json:"essential"`
            PrimaryMaterial string  `json:"primary_material"`
            PrimaryColorHex string  `json:"primary_color_hex"`
         
    }
    
    // Item used by Order
    type Item struct {

        
            L2Category []string  `json:"l2_category"`
            CanReturn bool  `json:"can_return"`
            LastUpdatedAt string  `json:"last_updated_at"`
            BrandID float64  `json:"brand_id"`
            Image []string  `json:"image"`
            L1Category []string  `json:"l1_category"`
            Attributes Attributes  `json:"attributes"`
            SlugKey string  `json:"slug_key"`
            WebstoreProductURL string  `json:"webstore_product_url"`
            L3Category float64  `json:"l3_category"`
            Brand string  `json:"brand"`
            Code string  `json:"code"`
            BranchURL string  `json:"branch_url"`
            L1CategoryID float64  `json:"l1_category_id"`
            ItemID float64  `json:"item_id"`
            Meta map[string]interface{}  `json:"meta"`
            Size string  `json:"size"`
            Gender string  `json:"gender"`
            DepartmentID float64  `json:"department_id"`
            L3CategoryName string  `json:"l3_category_name"`
            Name string  `json:"name"`
            CanCancel bool  `json:"can_cancel"`
            L2CategoryID float64  `json:"l2_category_id"`
            Color string  `json:"color"`
         
    }
    
    // ArticleDetails used by Order
    type ArticleDetails struct {

        
            Status map[string]interface{}  `json:"status"`
         
    }
    
    // StoreAddress used by Order
    type StoreAddress struct {

        
            Address1 string  `json:"address1"`
            City string  `json:"city"`
            Phone string  `json:"phone"`
            State string  `json:"state"`
            ContactPerson string  `json:"contact_person"`
            AddressCategory string  `json:"address_category"`
            Longitude float64  `json:"longitude"`
            Email string  `json:"email"`
            AddressType string  `json:"address_type"`
            Country string  `json:"country"`
            UpdatedAt string  `json:"updated_at"`
            Latitude float64  `json:"latitude"`
            Address2 string  `json:"address2"`
            CountryCode string  `json:"country_code"`
            Landmark string  `json:"landmark"`
            Area string  `json:"area"`
            CreatedAt string  `json:"created_at"`
            Version string  `json:"version"`
            Pincode float64  `json:"pincode"`
         
    }
    
    // StoreEinvoice used by Order
    type StoreEinvoice struct {

        
            User string  `json:"user"`
            Username string  `json:"username"`
            Enabled bool  `json:"enabled"`
            Password string  `json:"password"`
         
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
            URL string  `json:"url"`
            DsType string  `json:"ds_type"`
            Verified bool  `json:"verified"`
            Value string  `json:"value"`
         
    }
    
    // StoreDocuments used by Order
    type StoreDocuments struct {

        
            Gst Document  `json:"gst"`
         
    }
    
    // EInvoicePortalDetails used by Order
    type EInvoicePortalDetails struct {

        
            User string  `json:"user"`
            Password string  `json:"password"`
            Username string  `json:"username"`
         
    }
    
    // StoreMeta used by Order
    type StoreMeta struct {

        
            GstCredentials StoreGstCredentials  `json:"gst_credentials"`
            GstNumber string  `json:"gst_number"`
            ProductReturnConfig map[string]interface{}  `json:"product_return_config"`
            Stage string  `json:"stage"`
            DisplayName string  `json:"display_name"`
            EwaybillPortalDetails map[string]interface{}  `json:"ewaybill_portal_details"`
            Documents StoreDocuments  `json:"documents"`
            AdditionalContactDetails map[string]interface{}  `json:"additional_contact_details"`
            NotificationEmails []string  `json:"notification_emails"`
            Timing []map[string]interface{}  `json:"timing"`
            EinvoicePortalDetails EInvoicePortalDetails  `json:"einvoice_portal_details"`
         
    }
    
    // Store used by Order
    type Store struct {

        
            Address1 string  `json:"address1"`
            City string  `json:"city"`
            Phone float64  `json:"phone"`
            FulfillmentChannel string  `json:"fulfillment_channel"`
            IsActive bool  `json:"is_active"`
            BrandStoreTags []string  `json:"brand_store_tags"`
            IsArchived bool  `json:"is_archived"`
            ContactPerson string  `json:"contact_person"`
            PackagingMaterialCount float64  `json:"packaging_material_count"`
            State string  `json:"state"`
            LocationType string  `json:"location_type"`
            Longitude float64  `json:"longitude"`
            VatNo string  `json:"vat_no"`
            CompanyID float64  `json:"company_id"`
            BrandID interface{}  `json:"brand_id"`
            Country string  `json:"country"`
            StoreAddressJson StoreAddress  `json:"store_address_json"`
            AlohomoraUserID float64  `json:"alohomora_user_id"`
            UpdatedAt string  `json:"updated_at"`
            Latitude float64  `json:"latitude"`
            OrderIntegrationID string  `json:"order_integration_id"`
            StoreActiveFrom string  `json:"store_active_from"`
            MallArea string  `json:"mall_area"`
            Address2 string  `json:"address2"`
            Code string  `json:"code"`
            ParentStoreID float64  `json:"parent_store_id"`
            IsEnabledForRecon bool  `json:"is_enabled_for_recon"`
            Meta StoreMeta  `json:"meta"`
            LoginUsername string  `json:"login_username"`
            Name string  `json:"name"`
            SID string  `json:"s_id"`
            CreatedAt string  `json:"created_at"`
            StoreEmail string  `json:"store_email"`
            MallName string  `json:"mall_name"`
            Pincode string  `json:"pincode"`
         
    }
    
    // Brand used by Order
    type Brand struct {

        
            InvoicePrefix string  `json:"invoice_prefix"`
            Company string  `json:"company"`
            IsVirtualInvoice bool  `json:"is_virtual_invoice"`
            BrandID float64  `json:"brand_id"`
            CreatedOn float64  `json:"created_on"`
            BrandName string  `json:"brand_name"`
            StartDate string  `json:"start_date"`
            CreditNoteExpiryDays float64  `json:"credit_note_expiry_days"`
            PickupLocation string  `json:"pickup_location"`
            CreditNoteAllowed bool  `json:"credit_note_allowed"`
            Logo string  `json:"logo"`
            ScriptLastRan string  `json:"script_last_ran"`
            ModifiedOn float64  `json:"modified_on"`
         
    }
    
    // BagGSTDetails used by Order
    type BagGSTDetails struct {

        
            IgstTaxPercentage float64  `json:"igst_tax_percentage"`
            ValueOfGood float64  `json:"value_of_good"`
            CgstTaxPercentage float64  `json:"cgst_tax_percentage"`
            SgstTaxPercentage float64  `json:"sgst_tax_percentage"`
            IsDefaultHsnCode bool  `json:"is_default_hsn_code"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            IgstGstFee string  `json:"igst_gst_fee"`
            HsnCode string  `json:"hsn_code"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            SgstGstFee string  `json:"sgst_gst_fee"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            HsnCodeID string  `json:"hsn_code_id"`
            GstinCode string  `json:"gstin_code"`
            CgstGstFee string  `json:"cgst_gst_fee"`
            GstFee float64  `json:"gst_fee"`
            GstTag string  `json:"gst_tag"`
         
    }
    
    // BagDetailsPlatformResponse used by Order
    type BagDetailsPlatformResponse struct {

        
            AffiliateBagDetails AffiliateBagDetails  `json:"affiliate_bag_details"`
            Status BagReturnableCancelableStatus1  `json:"status"`
            ID float64  `json:"id"`
            DisplayName string  `json:"display_name"`
            LineNumber float64  `json:"line_number"`
            ParentPromoBags map[string]interface{}  `json:"parent_promo_bags"`
            BagStatus []BagStatusHistory  `json:"bag_status"`
            Quantity float64  `json:"quantity"`
            NoOfBagsOrder float64  `json:"no_of_bags_order"`
            BagUpdateTime float64  `json:"bag_update_time"`
            Identifier string  `json:"identifier"`
            OperationalStatus string  `json:"operational_status"`
            OriginalBagList []float64  `json:"original_bag_list"`
            QcRequired interface{}  `json:"qc_required"`
            RestorePromos map[string]interface{}  `json:"restore_promos"`
            Item Item  `json:"item"`
            ArticleDetails ArticleDetails  `json:"article_details"`
            CurrentOperationalStatus BagStatusHistory  `json:"current_operational_status"`
            OrderingStore Store  `json:"ordering_store"`
            RestoreCoupon bool  `json:"restore_coupon"`
            Prices Prices  `json:"prices"`
            OrderIntegrationID string  `json:"order_integration_id"`
            AppliedPromos []map[string]interface{}  `json:"applied_promos"`
            ShipmentID string  `json:"shipment_id"`
            Brand Brand  `json:"brand"`
            FinancialBreakup []FinancialBreakup  `json:"financial_breakup"`
            Dates Dates  `json:"dates"`
            Meta BagMeta  `json:"meta"`
            Reasons []map[string]interface{}  `json:"reasons"`
            SellerIdentifier string  `json:"seller_identifier"`
            Article Article  `json:"article"`
            AffiliateDetails AffiliateDetails  `json:"affiliate_details"`
            Type string  `json:"type"`
            BagStatusHistory BagStatusHistory  `json:"bag_status_history"`
            CurrentStatus BagStatusHistory  `json:"current_status"`
            EntityType string  `json:"entity_type"`
            JourneyType string  `json:"journey_type"`
            Tags []string  `json:"tags"`
            GstDetails BagGSTDetails  `json:"gst_details"`
         
    }
    
    // ErrorResponse used by Order
    type ErrorResponse struct {

        
            Error string  `json:"error"`
            Message string  `json:"message"`
         
    }
    
    // BagsPage used by Order
    type BagsPage struct {

        
            Size float64  `json:"size"`
            HasNext bool  `json:"has_next"`
            PageType string  `json:"page_type"`
            ItemTotal float64  `json:"item_total"`
            Current float64  `json:"current"`
         
    }
    
    // GetBagsPlatformResponse used by Order
    type GetBagsPlatformResponse struct {

        
            Items []BagDetailsPlatformResponse  `json:"items"`
            Page BagsPage  `json:"page"`
         
    }
    
    // InvalidateShipmentCachePayload used by Order
    type InvalidateShipmentCachePayload struct {

        
            AffiliateBagIds []string  `json:"affiliate_bag_ids"`
            BagIds []string  `json:"bag_ids"`
            ShipmentIds []string  `json:"shipment_ids"`
         
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

        
            Message string  `json:"message"`
            Status float64  `json:"status"`
            Success bool  `json:"success"`
            ErrorTrace string  `json:"error_trace"`
         
    }
    
    // StoreReassign used by Order
    type StoreReassign struct {

        
            AffiliateBagID string  `json:"affiliate_bag_id"`
            StoreID float64  `json:"store_id"`
            SetID string  `json:"set_id"`
            BagID float64  `json:"bag_id"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            ItemID string  `json:"item_id"`
            FyndOrderID string  `json:"fynd_order_id"`
            ReasonIds []float64  `json:"reason_ids"`
            AffiliateID string  `json:"affiliate_id"`
            MongoArticleID string  `json:"mongo_article_id"`
         
    }
    
    // StoreReassignResponse used by Order
    type StoreReassignResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // Entities used by Order
    type Entities struct {

        
            AffiliateBagID string  `json:"affiliate_bag_id"`
            ID string  `json:"id"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            ReasonText string  `json:"reason_text"`
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

        
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            AffiliateID string  `json:"affiliate_id"`
         
    }
    
    // Bags used by Order
    type Bags struct {

        
            AffiliateBagID string  `json:"affiliate_bag_id"`
            IsLocked bool  `json:"is_locked"`
            BagID float64  `json:"bag_id"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
         
    }
    
    // CheckResponse used by Order
    type CheckResponse struct {

        
            IsShipmentLocked bool  `json:"is_shipment_locked"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            OriginalFilter OriginalFilter  `json:"original_filter"`
            LockStatus string  `json:"lock_status"`
            IsBagLocked bool  `json:"is_bag_locked"`
            Bags []Bags  `json:"bags"`
            ShipmentID string  `json:"shipment_id"`
            Status string  `json:"status"`
            AffiliateID string  `json:"affiliate_id"`
         
    }
    
    // UpdateShipmentLockResponse used by Order
    type UpdateShipmentLockResponse struct {

        
            CheckResponse []CheckResponse  `json:"check_response"`
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // AnnouncementResponse used by Order
    type AnnouncementResponse struct {

        
            Title string  `json:"title"`
            CreatedAt string  `json:"created_at"`
            ID float64  `json:"id"`
            LogoURL string  `json:"logo_url"`
            PlatformID string  `json:"platform_id"`
            ToDatetime string  `json:"to_datetime"`
            Description string  `json:"description"`
            FromDatetime string  `json:"from_datetime"`
            PlatformName string  `json:"platform_name"`
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
    
    // Products used by Order
    type Products struct {

        
            LineNumber float64  `json:"line_number"`
            Quantity float64  `json:"quantity"`
            Identifier string  `json:"identifier"`
         
    }
    
    // ProductsReasonsFilters used by Order
    type ProductsReasonsFilters struct {

        
            LineNumber float64  `json:"line_number"`
            Quantity float64  `json:"quantity"`
            Identifier string  `json:"identifier"`
         
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

        
            Products []Products  `json:"products"`
            Reasons ReasonsData  `json:"reasons"`
            DataUpdates DataUpdates  `json:"data_updates"`
            Identifier string  `json:"identifier"`
         
    }
    
    // StatuesRequest used by Order
    type StatuesRequest struct {

        
            Shipments []ShipmentsRequest  `json:"shipments"`
            Status string  `json:"status"`
            SplitShipment bool  `json:"split_shipment"`
            ExcludeBagsNextState string  `json:"exclude_bags_next_state"`
         
    }
    
    // UpdateShipmentStatusRequest used by Order
    type UpdateShipmentStatusRequest struct {

        
            Task bool  `json:"task"`
            Statuses []StatuesRequest  `json:"statuses"`
            UnlockBeforeTransition bool  `json:"unlock_before_transition"`
            ForceTransition bool  `json:"force_transition"`
            LockAfterTransition bool  `json:"lock_after_transition"`
         
    }
    
    // ShipmentsResponse used by Order
    type ShipmentsResponse struct {

        
            Meta map[string]interface{}  `json:"meta"`
            Identifier string  `json:"identifier"`
            Message string  `json:"message"`
            Code string  `json:"code"`
            FinalState map[string]interface{}  `json:"final_state"`
            StackTrace string  `json:"stack_trace"`
            Status float64  `json:"status"`
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
    
    // ArticleDetails1 used by Order
    type ArticleDetails1 struct {

        
            Category map[string]interface{}  `json:"category"`
            Attributes map[string]interface{}  `json:"attributes"`
            BrandID float64  `json:"brand_id"`
            ID string  `json:"_id"`
            Dimension map[string]interface{}  `json:"dimension"`
            Quantity float64  `json:"quantity"`
            Weight map[string]interface{}  `json:"weight"`
         
    }
    
    // ShipmentDetails1 used by Order
    type ShipmentDetails1 struct {

        
            Meta map[string]interface{}  `json:"meta"`
            Shipments float64  `json:"shipments"`
            Articles []ArticleDetails1  `json:"articles"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            DpID float64  `json:"dp_id"`
            BoxType string  `json:"box_type"`
            FulfillmentID float64  `json:"fulfillment_id"`
         
    }
    
    // LocationDetails used by Order
    type LocationDetails struct {

        
            Articles []ArticleDetails1  `json:"articles"`
            FulfillmentType string  `json:"fulfillment_type"`
            FulfillmentID float64  `json:"fulfillment_id"`
         
    }
    
    // ShipmentConfig used by Order
    type ShipmentConfig struct {

        
            Shipment []ShipmentDetails1  `json:"shipment"`
            Identifier string  `json:"identifier"`
            Action string  `json:"action"`
            ToPincode string  `json:"to_pincode"`
            Journey string  `json:"journey"`
            LocationDetails LocationDetails  `json:"location_details"`
            PaymentMode string  `json:"payment_mode"`
            Source string  `json:"source"`
         
    }
    
    // ShipmentData used by Order
    type ShipmentData struct {

        
            ShipmentData ShipmentConfig  `json:"shipment_data"`
         
    }
    
    // OrderUser used by Order
    type OrderUser struct {

        
            State string  `json:"state"`
            Address2 string  `json:"address2"`
            Country string  `json:"country"`
            LastName string  `json:"last_name"`
            FirstName string  `json:"first_name"`
            Pincode string  `json:"pincode"`
            Mobile float64  `json:"mobile"`
            City string  `json:"city"`
            Phone float64  `json:"phone"`
            Email string  `json:"email"`
            Address1 string  `json:"address1"`
         
    }
    
    // MarketPlacePdf used by Order
    type MarketPlacePdf struct {

        
            Label string  `json:"label"`
            Invoice string  `json:"invoice"`
         
    }
    
    // AffiliateBag used by Order
    type AffiliateBag struct {

        
            TransferPrice float64  `json:"transfer_price"`
            UnitPrice float64  `json:"unit_price"`
            StoreID float64  `json:"store_id"`
            AvlQty float64  `json:"avl_qty"`
            FyndStoreID string  `json:"fynd_store_id"`
            AffiliateStoreID string  `json:"affiliate_store_id"`
            ItemID float64  `json:"item_id"`
            Quantity float64  `json:"quantity"`
            PriceEffective float64  `json:"price_effective"`
            Sku string  `json:"sku"`
            Identifier map[string]interface{}  `json:"identifier"`
            HsnCodeID string  `json:"hsn_code_id"`
            ModifiedOn string  `json:"modified_on"`
            PdfLinks MarketPlacePdf  `json:"pdf_links"`
            CompanyID float64  `json:"company_id"`
            Discount float64  `json:"discount"`
            ID string  `json:"_id"`
            SellerIdentifier string  `json:"seller_identifier"`
            DeliveryCharge float64  `json:"delivery_charge"`
            AmountPaid float64  `json:"amount_paid"`
            ItemSize string  `json:"item_size"`
            PriceMarked float64  `json:"price_marked"`
            AffiliateMeta map[string]interface{}  `json:"affiliate_meta"`
         
    }
    
    // UserData used by Order
    type UserData struct {

        
            ShippingUser OrderUser  `json:"shipping_user"`
            BillingUser OrderUser  `json:"billing_user"`
         
    }
    
    // OrderPriority used by Order
    type OrderPriority struct {

        
            FulfilmentPriority float64  `json:"fulfilment_priority"`
            FulfilmentPriorityText string  `json:"fulfilment_priority_text"`
            AffiliatePriorityCode string  `json:"affiliate_priority_code"`
         
    }
    
    // OrderInfo used by Order
    type OrderInfo struct {

        
            Shipment ShipmentData  `json:"shipment"`
            BillingAddress OrderUser  `json:"billing_address"`
            Payment map[string]interface{}  `json:"payment"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            Coupon string  `json:"coupon"`
            Bags []AffiliateBag  `json:"bags"`
            OrderValue float64  `json:"order_value"`
            User UserData  `json:"user"`
            Items map[string]interface{}  `json:"items"`
            CodCharges float64  `json:"cod_charges"`
            ShippingAddress OrderUser  `json:"shipping_address"`
            PaymentMode string  `json:"payment_mode"`
            Discount float64  `json:"discount"`
            OrderPriority OrderPriority  `json:"order_priority"`
            DeliveryCharges float64  `json:"delivery_charges"`
         
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
            Owner string  `json:"owner"`
            CreatedAt string  `json:"created_at"`
            ID string  `json:"id"`
            UpdatedAt string  `json:"updated_at"`
            Description string  `json:"description"`
            Secret string  `json:"secret"`
            Name string  `json:"name"`
            Token string  `json:"token"`
         
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

        
            ModeOfPayment string  `json:"mode_of_payment"`
            Source string  `json:"source"`
         
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
            Inventory AffiliateInventoryStoreConfig  `json:"inventory"`
            Payment AffiliateInventoryPaymentConfig  `json:"payment"`
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

        
            Config AffiliateConfig  `json:"config"`
            ID string  `json:"id"`
            Token string  `json:"token"`
         
    }
    
    // OrderConfig used by Order
    type OrderConfig struct {

        
            BagEndState string  `json:"bag_end_state"`
            AffiliateStoreIDMapping []AffiliateStoreIdMapping  `json:"affiliate_store_id_mapping"`
            ArticleLookup string  `json:"article_lookup"`
            StoreLookup string  `json:"store_lookup"`
            Affiliate Affiliate  `json:"affiliate"`
            CreateUser bool  `json:"create_user"`
         
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
    
    // PostHistoryFilters used by Order
    type PostHistoryFilters struct {

        
            LineNumber string  `json:"line_number"`
            ShipmentID string  `json:"shipment_id"`
            Identifier string  `json:"identifier"`
         
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

        
            Category string  `json:"category"`
            Text string  `json:"text"`
            State string  `json:"state"`
            Code float64  `json:"code"`
            DislayName string  `json:"dislay_name"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // HistoryMeta used by Order
    type HistoryMeta struct {

        
            Message string  `json:"message"`
            StoreID float64  `json:"store_id"`
            StoreName string  `json:"store_name"`
            Recordpath string  `json:"recordpath"`
            Endtime string  `json:"endtime"`
            Starttime string  `json:"starttime"`
            Caller string  `json:"caller"`
            Status1 string  `json:"status1"`
            Billsec string  `json:"billsec"`
            ActivityType string  `json:"activity_type"`
            CallID string  `json:"call_id"`
            Recipient string  `json:"recipient"`
            Status string  `json:"status"`
            Status2 string  `json:"status2"`
            Slug string  `json:"slug"`
            Callerid string  `json:"callerid"`
            ShortLink string  `json:"short_link"`
            ChannelType string  `json:"channel_type"`
            ActivityComment string  `json:"activity_comment"`
            Reason HistoryReason  `json:"reason"`
            StoreCode string  `json:"store_code"`
            Receiver string  `json:"receiver"`
            Duration string  `json:"duration"`
         
    }
    
    // HistoryDict used by Order
    type HistoryDict struct {

        
            Meta HistoryMeta  `json:"meta"`
            Type string  `json:"type"`
            DisplayMessage string  `json:"display_message"`
            AssignedAgent string  `json:"assigned_agent"`
            Message string  `json:"message"`
            BagID float64  `json:"bag_id"`
            L3Detail string  `json:"l3_detail"`
            TicketURL string  `json:"ticket_url"`
            User string  `json:"user"`
            L2Detail string  `json:"l2_detail"`
            L1Detail string  `json:"l1_detail"`
            Createdat string  `json:"createdat"`
            TicketID string  `json:"ticket_id"`
         
    }
    
    // ShipmentHistoryResponse used by Order
    type ShipmentHistoryResponse struct {

        
            Success bool  `json:"success"`
            ActivityHistory []HistoryDict  `json:"activity_history"`
         
    }
    
    // SmsDataPayload used by Order
    type SmsDataPayload struct {

        
            CustomerName string  `json:"customer_name"`
            AmountPaid float64  `json:"amount_paid"`
            OrderID string  `json:"order_id"`
            CountryCode string  `json:"country_code"`
            Message string  `json:"message"`
            PhoneNumber float64  `json:"phone_number"`
            BrandName string  `json:"brand_name"`
            ShipmentID float64  `json:"shipment_id"`
            PaymentMode string  `json:"payment_mode"`
         
    }
    
    // SendSmsPayload used by Order
    type SendSmsPayload struct {

        
            Slug string  `json:"slug"`
            Data SmsDataPayload  `json:"data"`
            BagID float64  `json:"bag_id"`
         
    }
    
    // OrderDetails used by Order
    type OrderDetails struct {

        
            CreatedAt string  `json:"created_at"`
            FyndOrderID string  `json:"fynd_order_id"`
         
    }
    
    // Meta used by Order
    type Meta struct {

        
            KafkaEmissionStatus float64  `json:"kafka_emission_status"`
            StateManagerUsed string  `json:"state_manager_used"`
         
    }
    
    // ShipmentDetail used by Order
    type ShipmentDetail struct {

        
            Meta Meta  `json:"meta"`
            BagList []float64  `json:"bag_list"`
            ID float64  `json:"id"`
            Remarks string  `json:"remarks"`
            ShipmentID string  `json:"shipment_id"`
            Status string  `json:"status"`
         
    }
    
    // OrderStatusData used by Order
    type OrderStatusData struct {

        
            Errors []string  `json:"errors"`
            OrderDetails OrderDetails  `json:"order_details"`
            ShipmentDetails []ShipmentDetail  `json:"shipment_details"`
         
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
            Height string  `json:"height"`
            Weight string  `json:"weight"`
         
    }
    
    // UpdatePackagingDimensionsPayload used by Order
    type UpdatePackagingDimensionsPayload struct {

        
            Dimension []Dimension  `json:"dimension"`
            CurrentStatus string  `json:"current_status"`
            ShipmentID string  `json:"shipment_id"`
         
    }
    
    // UpdatePackagingDimensionsResponse used by Order
    type UpdatePackagingDimensionsResponse struct {

        
            Message string  `json:"message"`
         
    }
    
    // Tax used by Order
    type Tax struct {

        
            Breakup []map[string]interface{}  `json:"breakup"`
            Name string  `json:"name"`
            Rate float64  `json:"rate"`
            Amount map[string]interface{}  `json:"amount"`
         
    }
    
    // Charge used by Order
    type Charge struct {

        
            Type string  `json:"type"`
            Tax Tax  `json:"tax"`
            Code string  `json:"code"`
            Amount map[string]interface{}  `json:"amount"`
            Name string  `json:"name"`
         
    }
    
    // ProcessingDates used by Order
    type ProcessingDates struct {

        
            DispatchAfterDate string  `json:"dispatch_after_date"`
            CustomerPickupSlot map[string]interface{}  `json:"customer_pickup_slot"`
            DispatchByDate string  `json:"dispatch_by_date"`
            PackByDate string  `json:"pack_by_date"`
            DpPickupSlot map[string]interface{}  `json:"dp_pickup_slot"`
            ConfirmByDate string  `json:"confirm_by_date"`
         
    }
    
    // LineItem used by Order
    type LineItem struct {

        
            Meta map[string]interface{}  `json:"meta"`
            Charges []Charge  `json:"charges"`
            ExternalLineID string  `json:"external_line_id"`
            Quantity float64  `json:"quantity"`
            CustomMessasge string  `json:"custom_messasge"`
            SellerIdentifier string  `json:"seller_identifier"`
         
    }
    
    // Shipment used by Order
    type Shipment struct {

        
            Meta map[string]interface{}  `json:"meta"`
            LocationID float64  `json:"location_id"`
            ExternalShipmentID string  `json:"external_shipment_id"`
            Priority float64  `json:"priority"`
            ProcessingDates ProcessingDates  `json:"processing_dates"`
            LineItems []LineItem  `json:"line_items"`
         
    }
    
    // ShippingInfo used by Order
    type ShippingInfo struct {

        
            AlternateMobileNumber string  `json:"alternate_mobile_number"`
            CountryCode string  `json:"country_code"`
            Slot []map[string]interface{}  `json:"slot"`
            Country string  `json:"country"`
            Landmark string  `json:"landmark"`
            CustomerCode string  `json:"customer_code"`
            AlternateEmail string  `json:"alternate_email"`
            StateCode string  `json:"state_code"`
            FirstName string  `json:"first_name"`
            LastName string  `json:"last_name"`
            Pincode string  `json:"pincode"`
            PrimaryEmail string  `json:"primary_email"`
            State string  `json:"state"`
            Address2 string  `json:"address2"`
            MiddleName string  `json:"middle_name"`
            ExternalCustomerCode string  `json:"external_customer_code"`
            GeoLocation map[string]interface{}  `json:"geo_location"`
            PrimaryMobileNumber string  `json:"primary_mobile_number"`
            HouseNo string  `json:"house_no"`
            Title string  `json:"title"`
            Gender string  `json:"gender"`
            FloorNo string  `json:"floor_no"`
            ShippingType string  `json:"shipping_type"`
            City string  `json:"city"`
            AddressType string  `json:"address_type"`
            Address1 string  `json:"address1"`
         
    }
    
    // PaymentMethod used by Order
    type PaymentMethod struct {

        
            Meta map[string]interface{}  `json:"meta"`
            RefundBy string  `json:"refund_by"`
            TransactionData map[string]interface{}  `json:"transaction_data"`
            Amount float64  `json:"amount"`
            CollectBy string  `json:"collect_by"`
            Mode string  `json:"mode"`
            Name string  `json:"name"`
         
    }
    
    // PaymentInfo used by Order
    type PaymentInfo struct {

        
            PrimaryMode string  `json:"primary_mode"`
            PaymentMethods []PaymentMethod  `json:"payment_methods"`
         
    }
    
    // BillingInfo used by Order
    type BillingInfo struct {

        
            AlternateMobileNumber string  `json:"alternate_mobile_number"`
            CountryCode string  `json:"country_code"`
            Country string  `json:"country"`
            CustomerCode string  `json:"customer_code"`
            AlternateEmail string  `json:"alternate_email"`
            StateCode string  `json:"state_code"`
            FirstName string  `json:"first_name"`
            LastName string  `json:"last_name"`
            Pincode string  `json:"pincode"`
            PrimaryEmail string  `json:"primary_email"`
            State string  `json:"state"`
            Address2 string  `json:"address2"`
            MiddleName string  `json:"middle_name"`
            ExternalCustomerCode string  `json:"external_customer_code"`
            PrimaryMobileNumber string  `json:"primary_mobile_number"`
            HouseNo string  `json:"house_no"`
            Title string  `json:"title"`
            Gender string  `json:"gender"`
            FloorNo string  `json:"floor_no"`
            City string  `json:"city"`
            Address1 string  `json:"address1"`
         
    }
    
    // TaxInfo used by Order
    type TaxInfo struct {

        
            Gstin string  `json:"gstin"`
            B2bGstinNumber string  `json:"b2b_gstin_number"`
         
    }
    
    // CreateOrderAPI used by Order
    type CreateOrderAPI struct {

        
            Meta map[string]interface{}  `json:"meta"`
            Charges []Charge  `json:"charges"`
            Shipments []Shipment  `json:"shipments"`
            ExternalOrderID string  `json:"external_order_id"`
            ShippingInfo ShippingInfo  `json:"shipping_info"`
            PaymentInfo PaymentInfo  `json:"payment_info"`
            ExternalCreationDate string  `json:"external_creation_date"`
            BillingInfo BillingInfo  `json:"billing_info"`
            Config map[string]interface{}  `json:"config"`
            CurrencyInfo map[string]interface{}  `json:"currency_info"`
            TaxInfo TaxInfo  `json:"tax_info"`
         
    }
    
    // CreateOrderErrorReponse used by Order
    type CreateOrderErrorReponse struct {

        
            Meta string  `json:"meta"`
            Message string  `json:"message"`
            RequestID string  `json:"request_id"`
            Info interface{}  `json:"info"`
            Code string  `json:"code"`
            StackTrace string  `json:"stack_trace"`
            Status float64  `json:"status"`
            Exception string  `json:"exception"`
         
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
            PaymentMethods []PaymentMethods  `json:"payment_methods"`
            Source string  `json:"source"`
         
    }
    
    // DpConfiguration used by Order
    type DpConfiguration struct {

        
            ShippingBy string  `json:"shipping_by"`
         
    }
    
    // CreateChannelConfig used by Order
    type CreateChannelConfig struct {

        
            ShipmentAssignment string  `json:"shipment_assignment"`
            LogoURL map[string]interface{}  `json:"logo_url"`
            PaymentInfo CreateChannelPaymentInfo  `json:"payment_info"`
            LockStates []string  `json:"lock_states"`
            LocationReassignment bool  `json:"location_reassignment"`
            DpConfiguration DpConfiguration  `json:"dp_configuration"`
         
    }
    
    // CreateChannelConfigData used by Order
    type CreateChannelConfigData struct {

        
            ConfigData CreateChannelConfig  `json:"config_data"`
         
    }
    
    // CreateChannelConfigResponse used by Order
    type CreateChannelConfigResponse struct {

        
            IsInserted bool  `json:"is_inserted"`
            Acknowledged bool  `json:"acknowledged"`
            IsUpserted bool  `json:"is_upserted"`
         
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
            OrderDetails []FyndOrderIdList  `json:"order_details"`
            EndDate string  `json:"end_date"`
            StartDate string  `json:"start_date"`
         
    }
    
    // BagStateTransitionMap used by Order
    type BagStateTransitionMap struct {

        
            Fynd map[string]interface{}  `json:"fynd"`
            Affiliate map[string]interface{}  `json:"affiliate"`
         
    }
    

    
    // SearchKeywordResult used by Catalog
    type SearchKeywordResult struct {

        
            Query map[string]interface{}  `json:"query"`
            SortOn string  `json:"sort_on"`
         
    }
    
    // CreateSearchKeyword used by Catalog
    type CreateSearchKeyword struct {

        
            IsActive bool  `json:"is_active"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Result SearchKeywordResult  `json:"result"`
            AppID string  `json:"app_id"`
            Words []string  `json:"words"`
         
    }
    
    // GetSearchWordsData used by Catalog
    type GetSearchWordsData struct {

        
            IsActive bool  `json:"is_active"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Result map[string]interface{}  `json:"result"`
            AppID string  `json:"app_id"`
            UID string  `json:"uid"`
            Words []string  `json:"words"`
         
    }
    
    // ErrorResponse used by Catalog
    type ErrorResponse struct {

        
            Code string  `json:"code"`
            Status float64  `json:"status"`
            Meta map[string]interface{}  `json:"meta"`
            Message string  `json:"message"`
            Error string  `json:"error"`
         
    }
    
    // DeleteResponse used by Catalog
    type DeleteResponse struct {

        
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
    
    // AutocompletePageAction used by Catalog
    type AutocompletePageAction struct {

        
            Query map[string]interface{}  `json:"query"`
            Params map[string]interface{}  `json:"params"`
            URL string  `json:"url"`
            Type string  `json:"type"`
         
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

        
            Display string  `json:"display"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Action AutocompleteAction  `json:"action"`
            Logo AutoCompleteMedia  `json:"logo"`
         
    }
    
    // CreateAutocompleteKeyword used by Catalog
    type CreateAutocompleteKeyword struct {

        
            IsActive bool  `json:"is_active"`
            Results []AutocompleteResult  `json:"results"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            AppID string  `json:"app_id"`
            Words []string  `json:"words"`
         
    }
    
    // GetAutocompleteWordsData used by Catalog
    type GetAutocompleteWordsData struct {

        
            Results []map[string]interface{}  `json:"results"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            AppID string  `json:"app_id"`
            UID string  `json:"uid"`
            Words []string  `json:"words"`
         
    }
    
    // GetAutocompleteWordsResponse used by Catalog
    type GetAutocompleteWordsResponse struct {

        
            Items []GetAutocompleteWordsData  `json:"items"`
            Page Page  `json:"page"`
         
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

        
            ProductUID float64  `json:"product_uid"`
            AutoAddToCart bool  `json:"auto_add_to_cart"`
            AutoSelect bool  `json:"auto_select"`
            MinQuantity float64  `json:"min_quantity"`
            MaxQuantity float64  `json:"max_quantity"`
            AllowRemove bool  `json:"allow_remove"`
         
    }
    
    // ProductBundleRequest used by Catalog
    type ProductBundleRequest struct {

        
            Name string  `json:"name"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            Meta map[string]interface{}  `json:"meta"`
            Products []ProductBundleItem  `json:"products"`
            CompanyID float64  `json:"company_id"`
            IsActive bool  `json:"is_active"`
            Slug string  `json:"slug"`
            Choice string  `json:"choice"`
            ModifiedOn string  `json:"modified_on"`
            Logo string  `json:"logo"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            CreatedOn string  `json:"created_on"`
            PageVisibility []string  `json:"page_visibility"`
         
    }
    
    // GetProductBundleCreateResponse used by Catalog
    type GetProductBundleCreateResponse struct {

        
            Name string  `json:"name"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            Meta map[string]interface{}  `json:"meta"`
            Products []ProductBundleItem  `json:"products"`
            CompanyID float64  `json:"company_id"`
            ID string  `json:"id"`
            IsActive bool  `json:"is_active"`
            Slug string  `json:"slug"`
            Choice string  `json:"choice"`
            ModifiedOn string  `json:"modified_on"`
            Logo string  `json:"logo"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            CreatedOn string  `json:"created_on"`
            PageVisibility []string  `json:"page_visibility"`
         
    }
    
    // GetProductBundleListingResponse used by Catalog
    type GetProductBundleListingResponse struct {

        
            Items []GetProductBundleCreateResponse  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // ProductBundleUpdateRequest used by Catalog
    type ProductBundleUpdateRequest struct {

        
            Name string  `json:"name"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            Meta map[string]interface{}  `json:"meta"`
            Products []ProductBundleItem  `json:"products"`
            CompanyID float64  `json:"company_id"`
            IsActive bool  `json:"is_active"`
            Slug string  `json:"slug"`
            Choice string  `json:"choice"`
            ModifiedOn string  `json:"modified_on"`
            Logo string  `json:"logo"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            PageVisibility []string  `json:"page_visibility"`
         
    }
    
    // LimitedProductData used by Catalog
    type LimitedProductData struct {

        
            Quantity float64  `json:"quantity"`
            Name string  `json:"name"`
            Identifier map[string]interface{}  `json:"identifier"`
            ItemCode string  `json:"item_code"`
            Images []string  `json:"images"`
            Price map[string]interface{}  `json:"price"`
            Slug string  `json:"slug"`
            Attributes map[string]interface{}  `json:"attributes"`
            UID float64  `json:"uid"`
            ShortDescription string  `json:"short_description"`
            Sizes []string  `json:"sizes"`
            CountryOfOrigin string  `json:"country_of_origin"`
         
    }
    
    // Size used by Catalog
    type Size struct {

        
            Display string  `json:"display"`
            IsAvailable bool  `json:"is_available"`
            Value string  `json:"value"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // Price used by Catalog
    type Price struct {

        
            MaxMarked float64  `json:"max_marked"`
            MinMarked float64  `json:"min_marked"`
            MinEffective float64  `json:"min_effective"`
            MaxEffective float64  `json:"max_effective"`
            Currency string  `json:"currency"`
         
    }
    
    // GetProducts used by Catalog
    type GetProducts struct {

        
            ProductDetails LimitedProductData  `json:"product_details"`
            ProductUID float64  `json:"product_uid"`
            Sizes []Size  `json:"sizes"`
            AutoAddToCart bool  `json:"auto_add_to_cart"`
            AutoSelect bool  `json:"auto_select"`
            Price Price  `json:"price"`
            MinQuantity float64  `json:"min_quantity"`
            MaxQuantity float64  `json:"max_quantity"`
            AllowRemove bool  `json:"allow_remove"`
         
    }
    
    // GetProductBundleResponse used by Catalog
    type GetProductBundleResponse struct {

        
            Name string  `json:"name"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            Products []GetProducts  `json:"products"`
            Meta map[string]interface{}  `json:"meta"`
            CompanyID float64  `json:"company_id"`
            IsActive bool  `json:"is_active"`
            Slug string  `json:"slug"`
            Choice string  `json:"choice"`
            Logo string  `json:"logo"`
            PageVisibility []string  `json:"page_visibility"`
         
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

        
            Name string  `json:"name"`
            Title string  `json:"title"`
            Image string  `json:"image"`
            Guide Guide  `json:"guide"`
            CompanyID float64  `json:"company_id"`
            ID string  `json:"id"`
            CreatedOn string  `json:"created_on"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Subtitle string  `json:"subtitle"`
            BrandID float64  `json:"brand_id"`
            ModifiedOn string  `json:"modified_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Active bool  `json:"active"`
            Tag string  `json:"tag"`
            Description string  `json:"description"`
         
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

        
            Name string  `json:"name"`
            Guide map[string]interface{}  `json:"guide"`
            Title string  `json:"title"`
            CompanyID float64  `json:"company_id"`
            ID string  `json:"id"`
            CreatedOn string  `json:"created_on"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ModifiedOn string  `json:"modified_on"`
            BrandID float64  `json:"brand_id"`
            Subtitle string  `json:"subtitle"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Active bool  `json:"active"`
            Tag string  `json:"tag"`
         
    }
    
    // MetaFields used by Catalog
    type MetaFields struct {

        
            Value interface{}  `json:"value"`
            Key interface{}  `json:"key"`
         
    }
    
    // ApplicationItemSEO used by Catalog
    type ApplicationItemSEO struct {

        
            Title interface{}  `json:"title"`
            Description interface{}  `json:"description"`
         
    }
    
    // ApplicationItemMOQ used by Catalog
    type ApplicationItemMOQ struct {

        
            IncrementUnit float64  `json:"increment_unit"`
            Minimum float64  `json:"minimum"`
            Maximum float64  `json:"maximum"`
         
    }
    
    // ApplicationItemMeta used by Catalog
    type ApplicationItemMeta struct {

        
            CustomMeta []MetaFields  `json:"_custom_meta"`
            Seo ApplicationItemSEO  `json:"seo"`
            Moq ApplicationItemMOQ  `json:"moq"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            AltText map[string]interface{}  `json:"alt_text"`
            IsGift bool  `json:"is_gift"`
            IsCod bool  `json:"is_cod"`
         
    }
    
    // SuccessResponse1 used by Catalog
    type SuccessResponse1 struct {

        
            Success bool  `json:"success"`
            UID float64  `json:"uid"`
         
    }
    
    // SEOData used by Catalog
    type SEOData struct {

        
            Title interface{}  `json:"title"`
            Description interface{}  `json:"description"`
         
    }
    
    // MOQData used by Catalog
    type MOQData struct {

        
            IncrementUnit float64  `json:"increment_unit"`
            Minimum float64  `json:"minimum"`
            Maximum float64  `json:"maximum"`
         
    }
    
    // OwnerAppItemResponse used by Catalog
    type OwnerAppItemResponse struct {

        
            Seo SEOData  `json:"seo"`
            Moq MOQData  `json:"moq"`
            AltText map[string]interface{}  `json:"alt_text"`
            IsGift bool  `json:"is_gift"`
            IsCod bool  `json:"is_cod"`
         
    }
    
    // GetConfigMetadataResponse used by Catalog
    type GetConfigMetadataResponse struct {

        
            Condition []map[string]interface{}  `json:"condition"`
            Data []map[string]interface{}  `json:"data"`
            Values []map[string]interface{}  `json:"values"`
         
    }
    
    // AttributeDetailsGroup used by Catalog
    type AttributeDetailsGroup struct {

        
            Name string  `json:"name"`
            Priority float64  `json:"priority"`
            IsActive bool  `json:"is_active"`
            Unit string  `json:"unit"`
            Key string  `json:"key"`
            DisplayType string  `json:"display_type"`
            Slug string  `json:"slug"`
            Logo string  `json:"logo"`
         
    }
    
    // AppConfigurationDetail used by Catalog
    type AppConfigurationDetail struct {

        
            Name string  `json:"name"`
            Priority float64  `json:"priority"`
            IsActive bool  `json:"is_active"`
            TemplateSlugs []string  `json:"template_slugs"`
            IsDefault bool  `json:"is_default"`
            Slug string  `json:"slug"`
            Logo string  `json:"logo"`
            AppID string  `json:"app_id"`
            Attributes []AttributeDetailsGroup  `json:"attributes"`
         
    }
    
    // ConfigErrorResponse used by Catalog
    type ConfigErrorResponse struct {

        
            Message string  `json:"message"`
         
    }
    
    // PageResponseType used by Catalog
    type PageResponseType struct {

        
            Next float64  `json:"next"`
            Current float64  `json:"current"`
            HasNext bool  `json:"has_next"`
            TotalCount float64  `json:"total_count"`
         
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

        
            Name string  `json:"name"`
            Priority float64  `json:"priority"`
            IsActive bool  `json:"is_active"`
            Key string  `json:"key"`
            IsDefault bool  `json:"is_default"`
            Logo string  `json:"logo"`
            AppID string  `json:"app_id"`
            DefaultKey string  `json:"default_key"`
         
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

        
            Similar map[string]interface{}  `json:"similar"`
            Compare map[string]interface{}  `json:"compare"`
            Variant map[string]interface{}  `json:"variant"`
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
    
    // ConfigurationProductConfig used by Catalog
    type ConfigurationProductConfig struct {

        
            Title string  `json:"title"`
            Size ProductSize  `json:"size"`
            Priority float64  `json:"priority"`
            IsActive bool  `json:"is_active"`
            Key string  `json:"key"`
            Subtitle string  `json:"subtitle"`
            Logo string  `json:"logo"`
         
    }
    
    // ConfigurationProductSimilar used by Catalog
    type ConfigurationProductSimilar struct {

        
            Config []ConfigurationProductConfig  `json:"config"`
         
    }
    
    // ConfigurationProductVariantConfig used by Catalog
    type ConfigurationProductVariantConfig struct {

        
            Name string  `json:"name"`
            Size ProductSize  `json:"size"`
            Priority float64  `json:"priority"`
            IsActive bool  `json:"is_active"`
            Key string  `json:"key"`
            DisplayType string  `json:"display_type"`
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
    
    // ConfigurationListingSortConfig used by Catalog
    type ConfigurationListingSortConfig struct {

        
            Name string  `json:"name"`
            Priority float64  `json:"priority"`
            IsActive bool  `json:"is_active"`
            Key string  `json:"key"`
            Logo string  `json:"logo"`
         
    }
    
    // ConfigurationListingSort used by Catalog
    type ConfigurationListingSort struct {

        
            DefaultKey string  `json:"default_key"`
            Config []ConfigurationListingSortConfig  `json:"config"`
         
    }
    
    // ConfigurationBucketPoints used by Catalog
    type ConfigurationBucketPoints struct {

        
            Display string  `json:"display"`
            Start float64  `json:"start"`
            End float64  `json:"end"`
         
    }
    
    // ConfigurationListingFilterValue used by Catalog
    type ConfigurationListingFilterValue struct {

        
            BucketPoints []ConfigurationBucketPoints  `json:"bucket_points"`
            Map map[string]interface{}  `json:"map"`
            Sort string  `json:"sort"`
            Condition string  `json:"condition"`
            Value string  `json:"value"`
            MapValues []map[string]interface{}  `json:"map_values"`
         
    }
    
    // ConfigurationListingFilterConfig used by Catalog
    type ConfigurationListingFilterConfig struct {

        
            Name string  `json:"name"`
            Priority float64  `json:"priority"`
            IsActive bool  `json:"is_active"`
            ValueConfig ConfigurationListingFilterValue  `json:"value_config"`
            Key string  `json:"key"`
            Type string  `json:"type"`
            DisplayName string  `json:"display_name"`
            Logo string  `json:"logo"`
         
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
    
    // AppConfiguration used by Catalog
    type AppConfiguration struct {

        
            ConfigID string  `json:"config_id"`
            Product ConfigurationProduct  `json:"product"`
            ConfigType string  `json:"config_type"`
            Listing ConfigurationListing  `json:"listing"`
            Type string  `json:"type"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            AppID string  `json:"app_id"`
         
    }
    
    // AppCatalogConfiguration used by Catalog
    type AppCatalogConfiguration struct {

        
            ConfigID string  `json:"config_id"`
            Product ConfigurationProduct  `json:"product"`
            ConfigType string  `json:"config_type"`
            ID string  `json:"id"`
            Listing ConfigurationListing  `json:"listing"`
            Type string  `json:"type"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            AppID string  `json:"app_id"`
         
    }
    
    // GetAppCatalogConfiguration used by Catalog
    type GetAppCatalogConfiguration struct {

        
            IsDefault bool  `json:"is_default"`
            Data AppCatalogConfiguration  `json:"data"`
         
    }
    
    // GetCatalogConfigurationDetailsSchemaListing used by Catalog
    type GetCatalogConfigurationDetailsSchemaListing struct {

        
            Sort map[string]interface{}  `json:"sort"`
            Filter map[string]interface{}  `json:"filter"`
         
    }
    
    // EntityConfiguration used by Catalog
    type EntityConfiguration struct {

        
            ConfigID string  `json:"config_id"`
            Product GetCatalogConfigurationDetailsProduct  `json:"product"`
            ConfigType string  `json:"config_type"`
            ID string  `json:"id"`
            Listing GetCatalogConfigurationDetailsSchemaListing  `json:"listing"`
            AppID string  `json:"app_id"`
         
    }
    
    // GetAppCatalogEntityConfiguration used by Catalog
    type GetAppCatalogEntityConfiguration struct {

        
            IsDefault bool  `json:"is_default"`
            Data EntityConfiguration  `json:"data"`
         
    }
    
    // ProductFiltersValue used by Catalog
    type ProductFiltersValue struct {

        
            Min float64  `json:"min"`
            SelectedMax float64  `json:"selected_max"`
            Display string  `json:"display"`
            DisplayFormat string  `json:"display_format"`
            Max float64  `json:"max"`
            CurrencyCode string  `json:"currency_code"`
            Value interface{}  `json:"value"`
            QueryFormat string  `json:"query_format"`
            CurrencySymbol string  `json:"currency_symbol"`
            IsSelected bool  `json:"is_selected"`
            Count float64  `json:"count"`
            SelectedMin float64  `json:"selected_min"`
         
    }
    
    // ProductFiltersKey used by Catalog
    type ProductFiltersKey struct {

        
            Name string  `json:"name"`
            Display string  `json:"display"`
            Kind string  `json:"kind"`
            Operators []string  `json:"operators"`
            Logo string  `json:"logo"`
         
    }
    
    // ProductFilters used by Catalog
    type ProductFilters struct {

        
            Values []ProductFiltersValue  `json:"values"`
            Key ProductFiltersKey  `json:"key"`
         
    }
    
    // ProductSortOn used by Catalog
    type ProductSortOn struct {

        
            Name string  `json:"name"`
            Value string  `json:"value"`
            IsSelected bool  `json:"is_selected"`
         
    }
    
    // GetCollectionQueryOptionResponse used by Catalog
    type GetCollectionQueryOptionResponse struct {

        
            Filters []ProductFilters  `json:"filters"`
            SortOn []ProductSortOn  `json:"sort_on"`
            Operators map[string]string  `json:"operators"`
         
    }
    
    // SeoDetail used by Catalog
    type SeoDetail struct {

        
            Title string  `json:"title"`
            Description string  `json:"description"`
         
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
    
    // UserInfo used by Catalog
    type UserInfo struct {

        
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
            Email string  `json:"email"`
            UID string  `json:"uid"`
         
    }
    
    // NextSchedule used by Catalog
    type NextSchedule struct {

        
            Start string  `json:"start"`
            End string  `json:"end"`
         
    }
    
    // CollectionSchedule used by Catalog
    type CollectionSchedule struct {

        
            Duration float64  `json:"duration"`
            Start string  `json:"start"`
            End string  `json:"end"`
            NextSchedule []NextSchedule  `json:"next_schedule"`
            Cron string  `json:"cron"`
         
    }
    
    // CollectionQuery used by Catalog
    type CollectionQuery struct {

        
            Value []interface{}  `json:"value"`
            Op string  `json:"op"`
            Attribute string  `json:"attribute"`
         
    }
    
    // CollectionBadge used by Catalog
    type CollectionBadge struct {

        
            Text string  `json:"text"`
            Color string  `json:"color"`
         
    }
    
    // CreateCollection used by Catalog
    type CreateCollection struct {

        
            Meta map[string]interface{}  `json:"meta"`
            Seo SeoDetail  `json:"seo"`
            Banners CollectionBanner  `json:"banners"`
            Tags []string  `json:"tags"`
            Published bool  `json:"published"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ModifiedBy UserInfo  `json:"modified_by"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            SortOn string  `json:"sort_on"`
            Name string  `json:"name"`
            Schedule CollectionSchedule  `json:"_schedule"`
            Query []CollectionQuery  `json:"query"`
            Type string  `json:"type"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            IsVisible bool  `json:"is_visible"`
            Badge CollectionBadge  `json:"badge"`
            Priority float64  `json:"priority"`
            IsActive bool  `json:"is_active"`
            Slug string  `json:"slug"`
            CreatedBy UserInfo  `json:"created_by"`
            Logo CollectionImage  `json:"logo"`
            AllowSort bool  `json:"allow_sort"`
            AppID string  `json:"app_id"`
            AllowFacets bool  `json:"allow_facets"`
            Description string  `json:"description"`
         
    }
    
    // BannerImage used by Catalog
    type BannerImage struct {

        
            Type string  `json:"type"`
            AspectRatio string  `json:"aspect_ratio"`
            URL string  `json:"url"`
         
    }
    
    // ImageUrls used by Catalog
    type ImageUrls struct {

        
            Landscape BannerImage  `json:"landscape"`
            Portrait BannerImage  `json:"portrait"`
         
    }
    
    // CollectionCreateResponse used by Catalog
    type CollectionCreateResponse struct {

        
            Meta map[string]interface{}  `json:"meta"`
            Banners ImageUrls  `json:"banners"`
            Cron map[string]interface{}  `json:"cron"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            SortOn string  `json:"sort_on"`
            Tag []string  `json:"tag"`
            Name string  `json:"name"`
            Schedule map[string]interface{}  `json:"_schedule"`
            Query []CollectionQuery  `json:"query"`
            Type string  `json:"type"`
            Badge map[string]interface{}  `json:"badge"`
            Priority float64  `json:"priority"`
            IsActive bool  `json:"is_active"`
            Slug string  `json:"slug"`
            Logo BannerImage  `json:"logo"`
            AllowSort bool  `json:"allow_sort"`
            AppID string  `json:"app_id"`
            AllowFacets bool  `json:"allow_facets"`
            Description string  `json:"description"`
         
    }
    
    // CollectionListingFilterTag used by Catalog
    type CollectionListingFilterTag struct {

        
            Name string  `json:"name"`
            Display string  `json:"display"`
            IsSelected bool  `json:"is_selected"`
         
    }
    
    // CollectionListingFilterType used by Catalog
    type CollectionListingFilterType struct {

        
            Name string  `json:"name"`
            Display string  `json:"display"`
            IsSelected bool  `json:"is_selected"`
         
    }
    
    // CollectionListingFilter used by Catalog
    type CollectionListingFilter struct {

        
            Tags []CollectionListingFilterTag  `json:"tags"`
            Type []CollectionListingFilterType  `json:"type"`
         
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
    
    // Media used by Catalog
    type Media struct {

        
            Type string  `json:"type"`
            Meta map[string]interface{}  `json:"meta"`
            URL string  `json:"url"`
         
    }
    
    // GetCollectionDetailNest used by Catalog
    type GetCollectionDetailNest struct {

        
            Meta map[string]interface{}  `json:"meta"`
            Banners ImageUrls  `json:"banners"`
            Action Action  `json:"action"`
            Cron map[string]interface{}  `json:"cron"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Tag []string  `json:"tag"`
            Name string  `json:"name"`
            Schedule map[string]interface{}  `json:"_schedule"`
            Query []CollectionQuery  `json:"query"`
            Type string  `json:"type"`
            Badge map[string]interface{}  `json:"badge"`
            UID string  `json:"uid"`
            Priority float64  `json:"priority"`
            IsActive bool  `json:"is_active"`
            Slug string  `json:"slug"`
            Logo Media  `json:"logo"`
            AllowSort bool  `json:"allow_sort"`
            AppID string  `json:"app_id"`
            AllowFacets bool  `json:"allow_facets"`
            Description string  `json:"description"`
         
    }
    
    // GetCollectionListingResponse used by Catalog
    type GetCollectionListingResponse struct {

        
            Filters CollectionListingFilter  `json:"filters"`
            Items []GetCollectionDetailNest  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // CollectionDetailResponse used by Catalog
    type CollectionDetailResponse struct {

        
            Meta map[string]interface{}  `json:"meta"`
            Banners ImageUrls  `json:"banners"`
            Cron map[string]interface{}  `json:"cron"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Tag []string  `json:"tag"`
            Name string  `json:"name"`
            Schedule map[string]interface{}  `json:"_schedule"`
            Query []CollectionQuery  `json:"query"`
            Type string  `json:"type"`
            Badge map[string]interface{}  `json:"badge"`
            UID string  `json:"uid"`
            Priority float64  `json:"priority"`
            IsActive bool  `json:"is_active"`
            Slug string  `json:"slug"`
            Logo Media  `json:"logo"`
            AllowSort bool  `json:"allow_sort"`
            AppID string  `json:"app_id"`
            AllowFacets bool  `json:"allow_facets"`
            Description string  `json:"description"`
         
    }
    
    // UpdateCollection used by Catalog
    type UpdateCollection struct {

        
            Meta map[string]interface{}  `json:"meta"`
            Seo SeoDetail  `json:"seo"`
            Banners CollectionBanner  `json:"banners"`
            Tags []string  `json:"tags"`
            Published bool  `json:"published"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ModifiedBy UserInfo  `json:"modified_by"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            SortOn string  `json:"sort_on"`
            Name string  `json:"name"`
            Schedule CollectionSchedule  `json:"_schedule"`
            Query []CollectionQuery  `json:"query"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            Type string  `json:"type"`
            IsVisible bool  `json:"is_visible"`
            Badge CollectionBadge  `json:"badge"`
            Priority float64  `json:"priority"`
            IsActive bool  `json:"is_active"`
            Slug string  `json:"slug"`
            Logo CollectionImage  `json:"logo"`
            AllowSort bool  `json:"allow_sort"`
            AllowFacets bool  `json:"allow_facets"`
            Description string  `json:"description"`
         
    }
    
    // CollectionItem used by Catalog
    type CollectionItem struct {

        
            Priority float64  `json:"priority"`
            Action string  `json:"action"`
            ItemID float64  `json:"item_id"`
         
    }
    
    // CollectionItemUpdate used by Catalog
    type CollectionItemUpdate struct {

        
            Items []CollectionItem  `json:"items"`
            Query []CollectionQuery  `json:"query"`
            Type string  `json:"type"`
            AllowSort bool  `json:"allow_sort"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            AllowFacets bool  `json:"allow_facets"`
         
    }
    
    // UpdatedResponse used by Catalog
    type UpdatedResponse struct {

        
            ItemsNotUpdated []float64  `json:"items_not_updated"`
            Message string  `json:"message"`
         
    }
    
    // Price1 used by Catalog
    type Price1 struct {

        
            Min float64  `json:"min"`
            CurrencySymbol string  `json:"currency_symbol"`
            Max float64  `json:"max"`
            CurrencyCode string  `json:"currency_code"`
         
    }
    
    // ProductListingPrice used by Catalog
    type ProductListingPrice struct {

        
            Effective Price1  `json:"effective"`
            Marked Price1  `json:"marked"`
         
    }
    
    // ProductDetailAttribute used by Catalog
    type ProductDetailAttribute struct {

        
            Value string  `json:"value"`
            Type string  `json:"type"`
            Key string  `json:"key"`
         
    }
    
    // ProductDetailGroupedAttribute used by Catalog
    type ProductDetailGroupedAttribute struct {

        
            Title string  `json:"title"`
            Details []ProductDetailAttribute  `json:"details"`
         
    }
    
    // ProductBrand used by Catalog
    type ProductBrand struct {

        
            Name string  `json:"name"`
            UID float64  `json:"uid"`
            Action Action  `json:"action"`
            Logo Media  `json:"logo"`
         
    }
    
    // ProductListingDetail used by Catalog
    type ProductListingDetail struct {

        
            Discount string  `json:"discount"`
            RatingCount float64  `json:"rating_count"`
            Price ProductListingPrice  `json:"price"`
            Medias []Media  `json:"medias"`
            Tryouts []string  `json:"tryouts"`
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
            Highlights []string  `json:"highlights"`
            ItemType string  `json:"item_type"`
            Color string  `json:"color"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            HasVariant bool  `json:"has_variant"`
            Name string  `json:"name"`
            ImageNature string  `json:"image_nature"`
            ProductOnlineDate string  `json:"product_online_date"`
            Type string  `json:"type"`
            Rating float64  `json:"rating"`
            Sellable bool  `json:"sellable"`
            UID float64  `json:"uid"`
            PromoMeta map[string]interface{}  `json:"promo_meta"`
            ItemCode string  `json:"item_code"`
            Brand ProductBrand  `json:"brand"`
            Slug string  `json:"slug"`
            Similars []string  `json:"similars"`
            Attributes map[string]interface{}  `json:"attributes"`
            ShortDescription string  `json:"short_description"`
            Description string  `json:"description"`
         
    }
    
    // GetCollectionItemsResponse used by Catalog
    type GetCollectionItemsResponse struct {

        
            Filters []ProductFilters  `json:"filters"`
            Items []ProductListingDetail  `json:"items"`
            SortOn []ProductSortOn  `json:"sort_on"`
            Page Page  `json:"page"`
         
    }
    
    // CatalogInsightItem used by Catalog
    type CatalogInsightItem struct {

        
            Count float64  `json:"count"`
            OutOfStockCount float64  `json:"out_of_stock_count"`
            SellableCount float64  `json:"sellable_count"`
         
    }
    
    // CatalogInsightBrand used by Catalog
    type CatalogInsightBrand struct {

        
            Name string  `json:"name"`
            AvailableArticles float64  `json:"available_articles"`
            TotalArticles float64  `json:"total_articles"`
            AvailableSizes float64  `json:"available_sizes"`
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
            CompanyID float64  `json:"company_id"`
            Enabled bool  `json:"enabled"`
            BrandIds []float64  `json:"brand_ids"`
            StoreIds []float64  `json:"store_ids"`
            OptLevel string  `json:"opt_level"`
         
    }
    
    // CompanyOptIn used by Catalog
    type CompanyOptIn struct {

        
            StoreIds []float64  `json:"store_ids"`
            Platform string  `json:"platform"`
            CompanyID float64  `json:"company_id"`
            Enabled bool  `json:"enabled"`
            CreatedOn float64  `json:"created_on"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ModifiedOn float64  `json:"modified_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            BrandIds []float64  `json:"brand_ids"`
            OptLevel string  `json:"opt_level"`
         
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
            CompanyType string  `json:"company_type"`
            BusinessType string  `json:"business_type"`
         
    }
    
    // CompanyBrandDetail used by Catalog
    type CompanyBrandDetail struct {

        
            BrandName string  `json:"brand_name"`
            TotalArticle float64  `json:"total_article"`
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
            Brand float64  `json:"brand"`
            Store float64  `json:"store"`
         
    }
    
    // StoreDetail used by Catalog
    type StoreDetail struct {

        
            Name string  `json:"name"`
            StoreType string  `json:"store_type"`
            CompanyID float64  `json:"company_id"`
            Address map[string]interface{}  `json:"address"`
            Manager map[string]interface{}  `json:"manager"`
            AdditionalContacts []map[string]interface{}  `json:"additional_contacts"`
            CreatedOn string  `json:"created_on"`
            Documents []map[string]interface{}  `json:"documents"`
            ModifiedOn string  `json:"modified_on"`
            DisplayName string  `json:"display_name"`
            StoreCode string  `json:"store_code"`
            UID float64  `json:"uid"`
            Timing map[string]interface{}  `json:"timing"`
         
    }
    
    // OptinStoreDetails used by Catalog
    type OptinStoreDetails struct {

        
            Items []StoreDetail  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // AttributeSchemaRange used by Catalog
    type AttributeSchemaRange struct {

        
            Min float64  `json:"min"`
            Max float64  `json:"max"`
         
    }
    
    // AttributeMaster used by Catalog
    type AttributeMaster struct {

        
            Range AttributeSchemaRange  `json:"range"`
            AllowedValues []string  `json:"allowed_values"`
            Multi bool  `json:"multi"`
            Mandatory bool  `json:"mandatory"`
            Type string  `json:"type"`
            Format string  `json:"format"`
         
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
    
    // AttributeMasterDetails used by Catalog
    type AttributeMasterDetails struct {

        
            DisplayType string  `json:"display_type"`
         
    }
    
    // AttributeMasterFilter used by Catalog
    type AttributeMasterFilter struct {

        
            Priority float64  `json:"priority"`
            DependsOn []string  `json:"depends_on"`
            Indexing bool  `json:"indexing"`
         
    }
    
    // GenderDetail used by Catalog
    type GenderDetail struct {

        
            Name string  `json:"name"`
            EnabledForEndConsumer bool  `json:"enabled_for_end_consumer"`
            Schema AttributeMaster  `json:"schema"`
            Meta AttributeMasterMeta  `json:"meta"`
            Details AttributeMasterDetails  `json:"details"`
            ID string  `json:"id"`
            Slug string  `json:"slug"`
            IsNested bool  `json:"is_nested"`
            Logo string  `json:"logo"`
            Filters AttributeMasterFilter  `json:"filters"`
            Departments []string  `json:"departments"`
            Description string  `json:"description"`
         
    }
    
    // CategoriesResponse used by Catalog
    type CategoriesResponse struct {

        
            Name string  `json:"name"`
            SlugKey string  `json:"slug_key"`
            Slug string  `json:"slug"`
            TemplateSlug string  `json:"template_slug"`
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
            Status float64  `json:"status"`
            Meta map[string]interface{}  `json:"meta"`
            Message string  `json:"message"`
         
    }
    
    // DepartmentCreateUpdate used by Catalog
    type DepartmentCreateUpdate struct {

        
            Name string  `json:"name"`
            Cls string  `json:"_cls"`
            Tags []string  `json:"tags"`
            IsActive bool  `json:"is_active"`
            Platforms map[string]interface{}  `json:"platforms"`
            Slug string  `json:"slug"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Synonyms []string  `json:"synonyms"`
            Logo string  `json:"logo"`
            PriorityOrder float64  `json:"priority_order"`
            UID float64  `json:"uid"`
         
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
            UserID string  `json:"user_id"`
            ID string  `json:"_id"`
            Contact string  `json:"contact"`
            UID string  `json:"uid"`
         
    }
    
    // GetDepartment used by Catalog
    type GetDepartment struct {

        
            Name string  `json:"name"`
            PageNo float64  `json:"page_no"`
            IsActive bool  `json:"is_active"`
            PageSize float64  `json:"page_size"`
            ItemType string  `json:"item_type"`
            Search string  `json:"search"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            Slug string  `json:"slug"`
            CreatedOn string  `json:"created_on"`
            Synonyms []string  `json:"synonyms"`
            Logo string  `json:"logo"`
            CreatedBy UserSerializer  `json:"created_by"`
            ModifiedOn string  `json:"modified_on"`
            PriorityOrder float64  `json:"priority_order"`
            UID float64  `json:"uid"`
         
    }
    
    // DepartmentsResponse used by Catalog
    type DepartmentsResponse struct {

        
            Items []GetDepartment  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // DepartmentErrorResponse used by Catalog
    type DepartmentErrorResponse struct {

        
            Errors map[string]interface{}  `json:"errors"`
            Code string  `json:"code"`
            Status float64  `json:"status"`
            Meta map[string]interface{}  `json:"meta"`
            Message string  `json:"message"`
         
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

        
            Name interface{}  `json:"name"`
            Cls interface{}  `json:"_cls"`
            IsActive bool  `json:"is_active"`
            ID interface{}  `json:"_id"`
            VerifiedOn string  `json:"verified_on"`
            CreatedOn string  `json:"created_on"`
            CreatedBy UserDetail  `json:"created_by"`
            ModifiedOn string  `json:"modified_on"`
            Slug interface{}  `json:"slug"`
            Logo interface{}  `json:"logo"`
            ModifiedBy UserDetail  `json:"modified_by"`
            Synonyms []interface{}  `json:"synonyms"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            PriorityOrder float64  `json:"priority_order"`
            UID float64  `json:"uid"`
            VerifiedBy UserDetail  `json:"verified_by"`
         
    }
    
    // ProductTemplate used by Catalog
    type ProductTemplate struct {

        
            Name string  `json:"name"`
            Description string  `json:"description"`
            IsActive bool  `json:"is_active"`
            IsExpirable bool  `json:"is_expirable"`
            IsArchived bool  `json:"is_archived"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Slug string  `json:"slug"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            CreatedOn string  `json:"created_on"`
            Logo string  `json:"logo"`
            ModifiedOn string  `json:"modified_on"`
            IsPhysical bool  `json:"is_physical"`
            Attributes []string  `json:"attributes"`
            Departments []string  `json:"departments"`
            Tag string  `json:"tag"`
            Categories []string  `json:"categories"`
         
    }
    
    // TemplatesResponse used by Catalog
    type TemplatesResponse struct {

        
            Items ProductTemplate  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // Properties used by Catalog
    type Properties struct {

        
            Trader map[string]interface{}  `json:"trader"`
            CategorySlug map[string]interface{}  `json:"category_slug"`
            Highlights map[string]interface{}  `json:"highlights"`
            MultiSize map[string]interface{}  `json:"multi_size"`
            Tags map[string]interface{}  `json:"tags"`
            ReturnConfig map[string]interface{}  `json:"return_config"`
            Command map[string]interface{}  `json:"command"`
            ItemType map[string]interface{}  `json:"item_type"`
            BrandUID map[string]interface{}  `json:"brand_uid"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            HsnCode map[string]interface{}  `json:"hsn_code"`
            Currency map[string]interface{}  `json:"currency"`
            Name map[string]interface{}  `json:"name"`
            Media map[string]interface{}  `json:"media"`
            NoOfBoxes map[string]interface{}  `json:"no_of_boxes"`
            IsDependent map[string]interface{}  `json:"is_dependent"`
            ProductGroupTag map[string]interface{}  `json:"product_group_tag"`
            TraderType map[string]interface{}  `json:"trader_type"`
            SizeGuide map[string]interface{}  `json:"size_guide"`
            Sizes map[string]interface{}  `json:"sizes"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            ItemCode map[string]interface{}  `json:"item_code"`
            Variants map[string]interface{}  `json:"variants"`
            IsActive map[string]interface{}  `json:"is_active"`
            ProductPublish map[string]interface{}  `json:"product_publish"`
            Slug map[string]interface{}  `json:"slug"`
            ShortDescription map[string]interface{}  `json:"short_description"`
            Description map[string]interface{}  `json:"description"`
            CountryOfOrigin map[string]interface{}  `json:"country_of_origin"`
         
    }
    
    // GlobalValidation used by Catalog
    type GlobalValidation struct {

        
            Title string  `json:"title"`
            Required []string  `json:"required"`
            Properties Properties  `json:"properties"`
            Type string  `json:"type"`
            Definitions map[string]interface{}  `json:"definitions"`
            Description string  `json:"description"`
         
    }
    
    // TemplateValidationData used by Catalog
    type TemplateValidationData struct {

        
            GlobalValidation GlobalValidation  `json:"global_validation"`
            TemplateValidation map[string]interface{}  `json:"template_validation"`
         
    }
    
    // TemplateDetails used by Catalog
    type TemplateDetails struct {

        
            Name string  `json:"name"`
            ID string  `json:"id"`
            Description string  `json:"description"`
            IsActive bool  `json:"is_active"`
            IsExpirable bool  `json:"is_expirable"`
            IsArchived bool  `json:"is_archived"`
            Slug string  `json:"slug"`
            Logo string  `json:"logo"`
            IsPhysical bool  `json:"is_physical"`
            Attributes []string  `json:"attributes"`
            Departments []string  `json:"departments"`
            Tag string  `json:"tag"`
            Categories []string  `json:"categories"`
         
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

        
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
            Email string  `json:"email"`
            UID string  `json:"uid"`
         
    }
    
    // ProductTemplateExportResponse used by Catalog
    type ProductTemplateExportResponse struct {

        
            Status string  `json:"status"`
            URL string  `json:"url"`
            SellerID float64  `json:"seller_id"`
            CompletedOn string  `json:"completed_on"`
            Type string  `json:"type"`
            CreatedBy UserInfo1  `json:"created_by"`
            ModifiedOn string  `json:"modified_on"`
            Filters map[string]interface{}  `json:"filters"`
            NotificationEmails []string  `json:"notification_emails"`
            TaskID string  `json:"task_id"`
         
    }
    
    // ProductDownloadsResponse used by Catalog
    type ProductDownloadsResponse struct {

        
            Items []ProductTemplateExportResponse  `json:"items"`
         
    }
    
    // ProductTemplateExportFilterRequest used by Catalog
    type ProductTemplateExportFilterRequest struct {

        
            ToDate string  `json:"to_date"`
            Brands []string  `json:"brands"`
            Templates []string  `json:"templates"`
            FromDate string  `json:"from_date"`
            CatalogueTypes []string  `json:"catalogue_types"`
         
    }
    
    // ProductTemplateDownloadsExport used by Catalog
    type ProductTemplateDownloadsExport struct {

        
            Filters ProductTemplateExportFilterRequest  `json:"filters"`
            NotificationEmails []string  `json:"notification_emails"`
            Type string  `json:"type"`
         
    }
    
    // ProductConfigurationDownloads used by Catalog
    type ProductConfigurationDownloads struct {

        
            Data []map[string]interface{}  `json:"data"`
            Multivalue bool  `json:"multivalue"`
         
    }
    
    // Media1 used by Catalog
    type Media1 struct {

        
            Landscape string  `json:"landscape"`
            Logo string  `json:"logo"`
            Portrait string  `json:"portrait"`
         
    }
    
    // CategoryMappingValues used by Catalog
    type CategoryMappingValues struct {

        
            Name string  `json:"name"`
            CatalogID float64  `json:"catalog_id"`
         
    }
    
    // CategoryMapping used by Catalog
    type CategoryMapping struct {

        
            Facebook CategoryMappingValues  `json:"facebook"`
            Ajio CategoryMappingValues  `json:"ajio"`
            Google CategoryMappingValues  `json:"google"`
         
    }
    
    // Hierarchy used by Catalog
    type Hierarchy struct {

        
            Department float64  `json:"department"`
            L2 float64  `json:"l2"`
            L1 float64  `json:"l1"`
         
    }
    
    // CategoryRequestBody used by Catalog
    type CategoryRequestBody struct {

        
            Name string  `json:"name"`
            Media Media1  `json:"media"`
            Marketplaces CategoryMapping  `json:"marketplaces"`
            Priority float64  `json:"priority"`
            IsActive bool  `json:"is_active"`
            Hierarchy []Hierarchy  `json:"hierarchy"`
            Slug string  `json:"slug"`
            Level float64  `json:"level"`
            Tryouts []string  `json:"tryouts"`
            Synonyms []string  `json:"synonyms"`
            Departments []float64  `json:"departments"`
         
    }
    
    // CategoryCreateResponse used by Catalog
    type CategoryCreateResponse struct {

        
            Message string  `json:"message"`
            UID float64  `json:"uid"`
         
    }
    
    // Category used by Catalog
    type Category struct {

        
            Name string  `json:"name"`
            Media Media1  `json:"media"`
            Marketplaces CategoryMapping  `json:"marketplaces"`
            Priority float64  `json:"priority"`
            IsActive bool  `json:"is_active"`
            ID string  `json:"id"`
            Hierarchy []Hierarchy  `json:"hierarchy"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Slug string  `json:"slug"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Level float64  `json:"level"`
            Tryouts []string  `json:"tryouts"`
            Synonyms []string  `json:"synonyms"`
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
            Departments []float64  `json:"departments"`
            UID float64  `json:"uid"`
         
    }
    
    // CategoryResponse used by Catalog
    type CategoryResponse struct {

        
            Items []Category  `json:"items"`
            Page Page  `json:"page"`
         
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
    
    // Trader used by Catalog
    type Trader struct {

        
            Name interface{}  `json:"name"`
            Type string  `json:"type"`
            Address []string  `json:"address"`
         
    }
    
    // NetQuantity used by Catalog
    type NetQuantity struct {

        
            Value float64  `json:"value"`
            Unit interface{}  `json:"unit"`
         
    }
    
    // TaxIdentifier used by Catalog
    type TaxIdentifier struct {

        
            HsnCode string  `json:"hsn_code"`
            ReportingHsn string  `json:"reporting_hsn"`
            HsnCodeID string  `json:"hsn_code_id"`
         
    }
    
    // ReturnConfig used by Catalog
    type ReturnConfig struct {

        
            Time float64  `json:"time"`
            Returnable bool  `json:"returnable"`
            Unit string  `json:"unit"`
         
    }
    
    // TeaserTag used by Catalog
    type TeaserTag struct {

        
            Tag string  `json:"tag"`
            URL string  `json:"url"`
         
    }
    
    // CustomOrder used by Catalog
    type CustomOrder struct {

        
            ManufacturingTime float64  `json:"manufacturing_time"`
            IsCustomOrder bool  `json:"is_custom_order"`
            ManufacturingTimeUnit string  `json:"manufacturing_time_unit"`
         
    }
    
    // ProductPublish used by Catalog
    type ProductPublish struct {

        
            ProductOnlineDate string  `json:"product_online_date"`
            IsSet bool  `json:"is_set"`
         
    }
    
    // ProductCreateUpdateSchemaV2 used by Catalog
    type ProductCreateUpdateSchemaV2 struct {

        
            Requester string  `json:"requester"`
            Action string  `json:"action"`
            Trader []Trader  `json:"trader"`
            NetQuantity NetQuantity  `json:"net_quantity"`
            CategorySlug string  `json:"category_slug"`
            Departments []float64  `json:"departments"`
            TaxIdentifier TaxIdentifier  `json:"tax_identifier"`
            Highlights []string  `json:"highlights"`
            IsSet bool  `json:"is_set"`
            CompanyID float64  `json:"company_id"`
            MultiSize bool  `json:"multi_size"`
            Tags []string  `json:"tags"`
            ReturnConfig ReturnConfig  `json:"return_config"`
            ItemType string  `json:"item_type"`
            BrandUID float64  `json:"brand_uid"`
            VariantGroup map[string]interface{}  `json:"variant_group"`
            TeaserTag TeaserTag  `json:"teaser_tag"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Currency string  `json:"currency"`
            VariantMedia map[string]interface{}  `json:"variant_media"`
            Name string  `json:"name"`
            Media []Media  `json:"media"`
            BulkJobID string  `json:"bulk_job_id"`
            NoOfBoxes float64  `json:"no_of_boxes"`
            IsDependent bool  `json:"is_dependent"`
            ProductGroupTag []string  `json:"product_group_tag"`
            UID float64  `json:"uid"`
            SizeGuide string  `json:"size_guide"`
            Sizes []map[string]interface{}  `json:"sizes"`
            CustomOrder CustomOrder  `json:"custom_order"`
            ItemCode string  `json:"item_code"`
            Variants map[string]interface{}  `json:"variants"`
            IsActive bool  `json:"is_active"`
            ProductPublish ProductPublish  `json:"product_publish"`
            Slug string  `json:"slug"`
            TemplateTag string  `json:"template_tag"`
            ChangeRequestID interface{}  `json:"change_request_id"`
            Attributes map[string]interface{}  `json:"attributes"`
            IsImageLessProduct bool  `json:"is_image_less_product"`
            ShortDescription string  `json:"short_description"`
            Description string  `json:"description"`
            CountryOfOrigin string  `json:"country_of_origin"`
         
    }
    
    // NetQuantityResponse used by Catalog
    type NetQuantityResponse struct {

        
            Value float64  `json:"value"`
            Unit string  `json:"unit"`
         
    }
    
    // Image used by Catalog
    type Image struct {

        
            SecureURL string  `json:"secure_url"`
            AspectRatio string  `json:"aspect_ratio"`
            AspectRatioF float64  `json:"aspect_ratio_f"`
            URL string  `json:"url"`
         
    }
    
    // VerifiedBy used by Catalog
    type VerifiedBy struct {

        
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
         
    }
    
    // ReturnConfigResponse used by Catalog
    type ReturnConfigResponse struct {

        
            Time float64  `json:"time"`
            Returnable bool  `json:"returnable"`
            Unit string  `json:"unit"`
         
    }
    
    // Logo used by Catalog
    type Logo struct {

        
            SecureURL string  `json:"secure_url"`
            AspectRatio string  `json:"aspect_ratio"`
            AspectRatioF float64  `json:"aspect_ratio_f"`
            URL string  `json:"url"`
         
    }
    
    // Brand used by Catalog
    type Brand struct {

        
            Name string  `json:"name"`
            UID float64  `json:"uid"`
            Logo Logo  `json:"logo"`
         
    }
    
    // ProductPublish1 used by Catalog
    type ProductPublish1 struct {

        
            ProductOnlineDate string  `json:"product_online_date"`
            IsSet bool  `json:"is_set"`
         
    }
    
    // ProductSchemaV2 used by Catalog
    type ProductSchemaV2 struct {

        
            Trader []Trader  `json:"trader"`
            NetQuantity NetQuantityResponse  `json:"net_quantity"`
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
            CategorySlug string  `json:"category_slug"`
            AllIdentifiers []string  `json:"all_identifiers"`
            IsPhysical bool  `json:"is_physical"`
            Departments []float64  `json:"departments"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            Images []Image  `json:"images"`
            VerifiedBy VerifiedBy  `json:"verified_by"`
            Category map[string]interface{}  `json:"category"`
            Highlights []string  `json:"highlights"`
            IsSet bool  `json:"is_set"`
            CompanyID float64  `json:"company_id"`
            MultiSize bool  `json:"multi_size"`
            AllSizes []map[string]interface{}  `json:"all_sizes"`
            Tags []string  `json:"tags"`
            ReturnConfig ReturnConfigResponse  `json:"return_config"`
            ItemType string  `json:"item_type"`
            VariantGroup map[string]interface{}  `json:"variant_group"`
            Color string  `json:"color"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            HsnCode string  `json:"hsn_code"`
            AllCompanyIds []float64  `json:"all_company_ids"`
            IsExpirable bool  `json:"is_expirable"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Currency string  `json:"currency"`
            VariantMedia map[string]interface{}  `json:"variant_media"`
            Name string  `json:"name"`
            ImageNature string  `json:"image_nature"`
            Media []Media  `json:"media"`
            ID string  `json:"id"`
            NoOfBoxes float64  `json:"no_of_boxes"`
            VerifiedOn string  `json:"verified_on"`
            IsDependent bool  `json:"is_dependent"`
            ProductGroupTag []string  `json:"product_group_tag"`
            Pending string  `json:"pending"`
            UID float64  `json:"uid"`
            ShortDescription string  `json:"short_description"`
            SizeGuide string  `json:"size_guide"`
            L3Mapping []string  `json:"l3_mapping"`
            Sizes []map[string]interface{}  `json:"sizes"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            ItemCode string  `json:"item_code"`
            Variants map[string]interface{}  `json:"variants"`
            Brand Brand  `json:"brand"`
            IsActive bool  `json:"is_active"`
            CategoryUID float64  `json:"category_uid"`
            PrimaryColor string  `json:"primary_color"`
            Moq map[string]interface{}  `json:"moq"`
            ProductPublish ProductPublish1  `json:"product_publish"`
            Slug string  `json:"slug"`
            TemplateTag string  `json:"template_tag"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Stage string  `json:"stage"`
            Attributes map[string]interface{}  `json:"attributes"`
            IsImageLessProduct bool  `json:"is_image_less_product"`
            BrandUID float64  `json:"brand_uid"`
            Description string  `json:"description"`
            CountryOfOrigin string  `json:"country_of_origin"`
         
    }
    
    // ProductListingResponseV2 used by Catalog
    type ProductListingResponseV2 struct {

        
            Items []ProductSchemaV2  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // ProductVariants used by Catalog
    type ProductVariants struct {

        
            Name string  `json:"name"`
            Media []Media  `json:"media"`
            ItemCode string  `json:"item_code"`
            CategoryUID float64  `json:"category_uid"`
            UID float64  `json:"uid"`
            BrandUID float64  `json:"brand_uid"`
         
    }
    
    // ProductVariantsResponse used by Catalog
    type ProductVariantsResponse struct {

        
            Variants []ProductVariants  `json:"variants"`
            Page Page  `json:"page"`
         
    }
    
    // AttributeMasterSerializer used by Catalog
    type AttributeMasterSerializer struct {

        
            Variant bool  `json:"variant"`
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
            Departments []string  `json:"departments"`
            Suggestion string  `json:"suggestion"`
            RawKey string  `json:"raw_key"`
            EnabledForEndConsumer bool  `json:"enabled_for_end_consumer"`
            Details AttributeMasterDetails  `json:"details"`
            Tags []string  `json:"tags"`
            Unit string  `json:"unit"`
            Synonyms map[string]interface{}  `json:"synonyms"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Example string  `json:"example"`
            Name string  `json:"name"`
            IsNested bool  `json:"is_nested"`
            Schema AttributeMaster  `json:"schema"`
            Slug string  `json:"slug"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Logo string  `json:"logo"`
            Filters AttributeMasterFilter  `json:"filters"`
            Description string  `json:"description"`
         
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
            Size interface{}  `json:"size"`
            Identifiers []ValidateIdentifier  `json:"identifiers"`
            ItemWeightUnitOfMeasure interface{}  `json:"item_weight_unit_of_measure"`
            ItemHeight float64  `json:"item_height"`
            ItemLength float64  `json:"item_length"`
            ItemWeight float64  `json:"item_weight"`
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

        
            Trader []Trader  `json:"trader"`
            NetQuantity NetQuantityResponse  `json:"net_quantity"`
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
            CategorySlug string  `json:"category_slug"`
            AllIdentifiers []string  `json:"all_identifiers"`
            IsPhysical bool  `json:"is_physical"`
            Departments []float64  `json:"departments"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            Images []Image  `json:"images"`
            VerifiedBy VerifiedBy  `json:"verified_by"`
            Category map[string]interface{}  `json:"category"`
            Highlights []string  `json:"highlights"`
            IsSet bool  `json:"is_set"`
            CompanyID float64  `json:"company_id"`
            MultiSize bool  `json:"multi_size"`
            AllSizes []map[string]interface{}  `json:"all_sizes"`
            Tags []string  `json:"tags"`
            ReturnConfig ReturnConfigResponse  `json:"return_config"`
            ItemType string  `json:"item_type"`
            VariantGroup map[string]interface{}  `json:"variant_group"`
            Color string  `json:"color"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            HsnCode string  `json:"hsn_code"`
            AllCompanyIds []float64  `json:"all_company_ids"`
            IsExpirable bool  `json:"is_expirable"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Currency string  `json:"currency"`
            VariantMedia map[string]interface{}  `json:"variant_media"`
            Name string  `json:"name"`
            ImageNature string  `json:"image_nature"`
            Media []Media  `json:"media"`
            ID string  `json:"id"`
            NoOfBoxes float64  `json:"no_of_boxes"`
            VerifiedOn string  `json:"verified_on"`
            IsDependent bool  `json:"is_dependent"`
            ProductGroupTag []string  `json:"product_group_tag"`
            Pending string  `json:"pending"`
            UID float64  `json:"uid"`
            ShortDescription string  `json:"short_description"`
            SizeGuide string  `json:"size_guide"`
            L3Mapping []string  `json:"l3_mapping"`
            Sizes []map[string]interface{}  `json:"sizes"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            ItemCode string  `json:"item_code"`
            Variants map[string]interface{}  `json:"variants"`
            Brand Brand  `json:"brand"`
            IsActive bool  `json:"is_active"`
            CategoryUID float64  `json:"category_uid"`
            PrimaryColor string  `json:"primary_color"`
            Moq map[string]interface{}  `json:"moq"`
            ProductPublish ProductPublished  `json:"product_publish"`
            Slug string  `json:"slug"`
            TemplateTag string  `json:"template_tag"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Stage string  `json:"stage"`
            Attributes map[string]interface{}  `json:"attributes"`
            IsImageLessProduct bool  `json:"is_image_less_product"`
            BrandUID float64  `json:"brand_uid"`
            Description string  `json:"description"`
            CountryOfOrigin string  `json:"country_of_origin"`
         
    }
    
    // ProductListingResponse used by Catalog
    type ProductListingResponse struct {

        
            Items []Product  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // BulkJob used by Catalog
    type BulkJob struct {

        
            CancelledRecords []map[string]interface{}  `json:"cancelled_records"`
            CompanyID float64  `json:"company_id"`
            IsActive bool  `json:"is_active"`
            Total float64  `json:"total"`
            Cancelled float64  `json:"cancelled"`
            FailedRecords []map[string]interface{}  `json:"failed_records"`
            FilePath string  `json:"file_path"`
            CreatedOn string  `json:"created_on"`
            CreatedBy UserInfo1  `json:"created_by"`
            ModifiedOn string  `json:"modified_on"`
            Stage string  `json:"stage"`
            TemplateTag string  `json:"template_tag"`
            ModifiedBy UserInfo1  `json:"modified_by"`
            Failed float64  `json:"failed"`
            CustomTemplateTag string  `json:"custom_template_tag"`
            TrackingURL string  `json:"tracking_url"`
            Succeed float64  `json:"succeed"`
         
    }
    
    // BulkResponse used by Catalog
    type BulkResponse struct {

        
            IsActive bool  `json:"is_active"`
            BatchID string  `json:"batch_id"`
            CreatedOn string  `json:"created_on"`
            CreatedBy UserInfo1  `json:"created_by"`
            ModifiedOn string  `json:"modified_on"`
            ModifiedBy UserInfo1  `json:"modified_by"`
         
    }
    
    // UserDetail1 used by Catalog
    type UserDetail1 struct {

        
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
            FullName string  `json:"full_name"`
         
    }
    
    // ProductBulkRequest used by Catalog
    type ProductBulkRequest struct {

        
            CancelledRecords []string  `json:"cancelled_records"`
            CompanyID float64  `json:"company_id"`
            IsActive bool  `json:"is_active"`
            Total float64  `json:"total"`
            Cancelled float64  `json:"cancelled"`
            FilePath string  `json:"file_path"`
            Stage string  `json:"stage"`
            TemplateTag string  `json:"template_tag"`
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
            CreatedBy UserDetail1  `json:"created_by"`
            FailedRecords []string  `json:"failed_records"`
            Failed float64  `json:"failed"`
            ModifiedBy UserDetail1  `json:"modified_by"`
            Template ProductTemplate  `json:"template"`
            Succeed float64  `json:"succeed"`
         
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
            BatchID string  `json:"batch_id"`
         
    }
    
    // UserCommon used by Catalog
    type UserCommon struct {

        
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
            CompanyID float64  `json:"company_id"`
         
    }
    
    // Items used by Catalog
    type Items struct {

        
            CancelledRecords []string  `json:"cancelled_records"`
            CompanyID float64  `json:"company_id"`
            ID string  `json:"id"`
            IsActive bool  `json:"is_active"`
            TrackingURL string  `json:"tracking_url"`
            Total float64  `json:"total"`
            FailedRecords []string  `json:"failed_records"`
            Cancelled float64  `json:"cancelled"`
            FilePath string  `json:"file_path"`
            CreatedOn string  `json:"created_on"`
            CreatedBy UserCommon  `json:"created_by"`
            ModifiedOn string  `json:"modified_on"`
            Stage string  `json:"stage"`
            ModifiedBy UserCommon  `json:"modified_by"`
            Failed float64  `json:"failed"`
            Retry float64  `json:"retry"`
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
            CompanyID float64  `json:"company_id"`
            ItemID float64  `json:"item_id"`
         
    }
    
    // ProductSizeDeleteResponse used by Catalog
    type ProductSizeDeleteResponse struct {

        
            Success bool  `json:"success"`
            Data ProductSizeDeleteDataResponse  `json:"data"`
         
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

        
            Quantity float64  `json:"quantity"`
            SizeDistribution SizeDistribution  `json:"size_distribution"`
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

        
            ItemDimensionsUnitOfMeasure string  `json:"item_dimensions_unit_of_measure"`
            Quantity float64  `json:"quantity"`
            Set InventorySet  `json:"set"`
            StoreCode string  `json:"store_code"`
            IsSet bool  `json:"is_set"`
            Size interface{}  `json:"size"`
            Identifiers []GTIN  `json:"identifiers"`
            PriceEffective float64  `json:"price_effective"`
            PriceTransfer float64  `json:"price_transfer"`
            ItemWeightUnitOfMeasure string  `json:"item_weight_unit_of_measure"`
            Price float64  `json:"price"`
            ItemHeight float64  `json:"item_height"`
            ItemLength float64  `json:"item_length"`
            ExpirationDate string  `json:"expiration_date"`
            ItemWeight float64  `json:"item_weight"`
            ItemWidth float64  `json:"item_width"`
            Currency string  `json:"currency"`
         
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
            Quantity float64  `json:"quantity"`
            Size string  `json:"size"`
            Identifiers map[string]interface{}  `json:"identifiers"`
            PriceEffective float64  `json:"price_effective"`
            PriceTransfer float64  `json:"price_transfer"`
            ItemID float64  `json:"item_id"`
            Price float64  `json:"price"`
            SellableQuantity float64  `json:"sellable_quantity"`
            SellerIdentifier string  `json:"seller_identifier"`
            Currency string  `json:"currency"`
            UID string  `json:"uid"`
            Store map[string]interface{}  `json:"store"`
         
    }
    
    // InventoryResponsePaginated used by Catalog
    type InventoryResponsePaginated struct {

        
            Items []InventoryResponse  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // Trader1 used by Catalog
    type Trader1 struct {

        
            Name string  `json:"name"`
            Type string  `json:"type"`
            Address []string  `json:"address"`
         
    }
    
    // PriceMeta used by Catalog
    type PriceMeta struct {

        
            Effective float64  `json:"effective"`
            Marked float64  `json:"marked"`
            TpNotes map[string]interface{}  `json:"tp_notes"`
            Transfer float64  `json:"transfer"`
            UpdatedAt string  `json:"updated_at"`
            Currency string  `json:"currency"`
         
    }
    
    // DimensionResponse used by Catalog
    type DimensionResponse struct {

        
            Unit string  `json:"unit"`
            Length float64  `json:"length"`
            Height float64  `json:"height"`
            IsDefault bool  `json:"is_default"`
            Width float64  `json:"width"`
         
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
    
    // WeightResponse used by Catalog
    type WeightResponse struct {

        
            IsDefault bool  `json:"is_default"`
            Unit string  `json:"unit"`
            Shipping float64  `json:"shipping"`
         
    }
    
    // ManufacturerResponse used by Catalog
    type ManufacturerResponse struct {

        
            Name string  `json:"name"`
            IsDefault bool  `json:"is_default"`
            Address string  `json:"address"`
         
    }
    
    // QuantityBase used by Catalog
    type QuantityBase struct {

        
            Count float64  `json:"count"`
            UpdatedAt string  `json:"updated_at"`
         
    }
    
    // Quantities used by Catalog
    type Quantities struct {

        
            NotAvailable QuantityBase  `json:"not_available"`
            Sellable QuantityBase  `json:"sellable"`
            OrderCommitted QuantityBase  `json:"order_committed"`
            Damaged QuantityBase  `json:"damaged"`
         
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
    
    // InventorySellerResponse used by Catalog
    type InventorySellerResponse struct {

        
            Meta map[string]interface{}  `json:"meta"`
            Size string  `json:"size"`
            Trader []Trader1  `json:"trader"`
            Price PriceMeta  `json:"price"`
            FyndItemCode string  `json:"fynd_item_code"`
            TrackInventory bool  `json:"track_inventory"`
            TotalQuantity float64  `json:"total_quantity"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            Dimension DimensionResponse  `json:"dimension"`
            Company CompanyMeta  `json:"company"`
            Identifier map[string]interface{}  `json:"identifier"`
            IsSet bool  `json:"is_set"`
            Tags []string  `json:"tags"`
            ReturnConfig ReturnConfig1  `json:"return_config"`
            ItemID float64  `json:"item_id"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            FyndMeta map[string]interface{}  `json:"fynd_meta"`
            Fragile bool  `json:"fragile"`
            RawMeta map[string]interface{}  `json:"raw_meta"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            Weight WeightResponse  `json:"weight"`
            Set InventorySet  `json:"set"`
            TraceID string  `json:"trace_id"`
            Manufacturer ManufacturerResponse  `json:"manufacturer"`
            AddedOnStore string  `json:"added_on_store"`
            Quantities Quantities  `json:"quantities"`
            SellerIdentifier string  `json:"seller_identifier"`
            UID string  `json:"uid"`
            Store StoreMeta  `json:"store"`
            IsActive bool  `json:"is_active"`
            Brand BrandMeta  `json:"brand"`
            Stage string  `json:"stage"`
            CreatedBy UserSerializer  `json:"created_by"`
            ExpirationDate string  `json:"expiration_date"`
            FyndArticleCode string  `json:"fynd_article_code"`
            CountryOfOrigin string  `json:"country_of_origin"`
         
    }
    
    // InventorySellerIdentifierResponsePaginated used by Catalog
    type InventorySellerIdentifierResponsePaginated struct {

        
            Items []InventorySellerResponse  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // Trader2 used by Catalog
    type Trader2 struct {

        
            Name string  `json:"name"`
            Type string  `json:"type"`
            Address []string  `json:"address"`
         
    }
    
    // PriceArticle used by Catalog
    type PriceArticle struct {

        
            Effective float64  `json:"effective"`
            Marked float64  `json:"marked"`
            TpNotes map[string]interface{}  `json:"tp_notes"`
            Transfer float64  `json:"transfer"`
            Currency string  `json:"currency"`
         
    }
    
    // DimensionResponse1 used by Catalog
    type DimensionResponse1 struct {

        
            Width float64  `json:"width"`
            Unit string  `json:"unit"`
            Length float64  `json:"length"`
            Height float64  `json:"height"`
         
    }
    
    // CompanyMeta1 used by Catalog
    type CompanyMeta1 struct {

        
            ID float64  `json:"id"`
         
    }
    
    // ReturnConfig2 used by Catalog
    type ReturnConfig2 struct {

        
            Time float64  `json:"time"`
            Returnable bool  `json:"returnable"`
            Unit string  `json:"unit"`
         
    }
    
    // WeightResponse1 used by Catalog
    type WeightResponse1 struct {

        
            Unit string  `json:"unit"`
            Shipping float64  `json:"shipping"`
         
    }
    
    // ManufacturerResponse1 used by Catalog
    type ManufacturerResponse1 struct {

        
            Name string  `json:"name"`
            IsDefault bool  `json:"is_default"`
            Address string  `json:"address"`
         
    }
    
    // Quantity used by Catalog
    type Quantity struct {

        
            Count float64  `json:"count"`
         
    }
    
    // QuantitiesArticle used by Catalog
    type QuantitiesArticle struct {

        
            NotAvailable Quantity  `json:"not_available"`
            Sellable Quantity  `json:"sellable"`
            OrderCommitted Quantity  `json:"order_committed"`
            Damaged Quantity  `json:"damaged"`
         
    }
    
    // DateMeta used by Catalog
    type DateMeta struct {

        
            InventoryUpdatedOn string  `json:"inventory_updated_on"`
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
            AddedOnStore string  `json:"added_on_store"`
         
    }
    
    // ArticleStoreResponse used by Catalog
    type ArticleStoreResponse struct {

        
            Name string  `json:"name"`
            StoreCode string  `json:"store_code"`
            UID float64  `json:"uid"`
            StoreType string  `json:"store_type"`
         
    }
    
    // BrandMeta1 used by Catalog
    type BrandMeta1 struct {

        
            ID float64  `json:"id"`
            Name string  `json:"name"`
         
    }
    
    // GetInventories used by Catalog
    type GetInventories struct {

        
            Size string  `json:"size"`
            Platforms map[string]interface{}  `json:"platforms"`
            Trader []Trader2  `json:"trader"`
            Price PriceArticle  `json:"price"`
            TrackInventory bool  `json:"track_inventory"`
            TotalQuantity float64  `json:"total_quantity"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            Dimension DimensionResponse1  `json:"dimension"`
            Company CompanyMeta1  `json:"company"`
            Identifier map[string]interface{}  `json:"identifier"`
            IsSet bool  `json:"is_set"`
            Tags []string  `json:"tags"`
            ReturnConfig ReturnConfig2  `json:"return_config"`
            ItemID float64  `json:"item_id"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            Weight WeightResponse1  `json:"weight"`
            TraceID string  `json:"trace_id"`
            ID string  `json:"id"`
            Manufacturer ManufacturerResponse1  `json:"manufacturer"`
            Quantities QuantitiesArticle  `json:"quantities"`
            DateMeta DateMeta  `json:"date_meta"`
            SellerIdentifier string  `json:"seller_identifier"`
            UID string  `json:"uid"`
            Store ArticleStoreResponse  `json:"store"`
            InventoryUpdatedOn string  `json:"inventory_updated_on"`
            Brand BrandMeta1  `json:"brand"`
            Stage string  `json:"stage"`
            CreatedBy UserSerializer  `json:"created_by"`
            ExpirationDate string  `json:"expiration_date"`
            CountryOfOrigin string  `json:"country_of_origin"`
         
    }
    
    // GetInventoriesResponse used by Catalog
    type GetInventoriesResponse struct {

        
            Items []GetInventories  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // BulkInventoryGetItems used by Catalog
    type BulkInventoryGetItems struct {

        
            CancelledRecords []string  `json:"cancelled_records"`
            CompanyID float64  `json:"company_id"`
            ID string  `json:"id"`
            IsActive bool  `json:"is_active"`
            Cancelled float64  `json:"cancelled"`
            Total float64  `json:"total"`
            FilePath string  `json:"file_path"`
            CreatedOn string  `json:"created_on"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            ModifiedOn string  `json:"modified_on"`
            Stage string  `json:"stage"`
            FailedRecords []string  `json:"failed_records"`
            Failed float64  `json:"failed"`
            Succeed float64  `json:"succeed"`
         
    }
    
    // BulkInventoryGet used by Catalog
    type BulkInventoryGet struct {

        
            Items []BulkInventoryGetItems  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // InventoryJobPayload used by Catalog
    type InventoryJobPayload struct {

        
            ItemDimensionsUnitOfMeasure string  `json:"item_dimensions_unit_of_measure"`
            Quantity float64  `json:"quantity"`
            Currency string  `json:"currency"`
            TraceID string  `json:"trace_id"`
            Tags []string  `json:"tags"`
            PriceEffective float64  `json:"price_effective"`
            TotalQuantity float64  `json:"total_quantity"`
            ItemWeightUnitOfMeasure string  `json:"item_weight_unit_of_measure"`
            Price float64  `json:"price"`
            ExpirationDate string  `json:"expiration_date"`
            StoreCode string  `json:"store_code"`
            SellerIdentifier string  `json:"seller_identifier"`
            PriceMarked float64  `json:"price_marked"`
         
    }
    
    // InventoryBulkRequest used by Catalog
    type InventoryBulkRequest struct {

        
            BatchID string  `json:"batch_id"`
            User map[string]interface{}  `json:"user"`
            CompanyID float64  `json:"company_id"`
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

        
            Status string  `json:"status"`
            SellerID float64  `json:"seller_id"`
            CreatedOn string  `json:"created_on"`
            CreatedBy string  `json:"created_by"`
            ModifiedOn string  `json:"modified_on"`
            Type string  `json:"type"`
            Filters map[string]interface{}  `json:"filters"`
            NotificationEmails []string  `json:"notification_emails"`
            TaskID string  `json:"task_id"`
         
    }
    
    // InventoryExportQuantityFilter used by Catalog
    type InventoryExportQuantityFilter struct {

        
            Min float64  `json:"min"`
            Max float64  `json:"max"`
            Operators string  `json:"operators"`
         
    }
    
    // InventoryExportAdvanceOption used by Catalog
    type InventoryExportAdvanceOption struct {

        
            Quantity InventoryExportQuantityFilter  `json:"quantity"`
            ToDate string  `json:"to_date"`
            StoreIds []float64  `json:"store_ids"`
            BrandIds []float64  `json:"brand_ids"`
            FromDate string  `json:"from_date"`
         
    }
    
    // InventoryExportJob used by Catalog
    type InventoryExportJob struct {

        
            Status string  `json:"status"`
            URL string  `json:"url"`
            SellerID float64  `json:"seller_id"`
            CompletedOn string  `json:"completed_on"`
            Type string  `json:"type"`
            Filters InventoryExportAdvanceOption  `json:"filters"`
            NotificationEmails []string  `json:"notification_emails"`
            TaskID string  `json:"task_id"`
         
    }
    
    // InventoryExportFilter used by Catalog
    type InventoryExportFilter struct {

        
            Quantity InventoryExportQuantityFilter  `json:"quantity"`
            ToDate string  `json:"to_date"`
            StoreIds []float64  `json:"store_ids"`
            BrandIds []float64  `json:"brand_ids"`
            FromDate string  `json:"from_date"`
         
    }
    
    // InventoryCreateRequest used by Catalog
    type InventoryCreateRequest struct {

        
            Filters InventoryExportFilter  `json:"filters"`
            NotificationEmails []string  `json:"notification_emails"`
            Type string  `json:"type"`
            Data []string  `json:"data"`
         
    }
    
    // InventoryJobFilters used by Catalog
    type InventoryJobFilters struct {

        
            Quantity InventoryExportQuantityFilter  `json:"quantity"`
            ToDate string  `json:"to_date"`
            Brands []string  `json:"brands"`
            FromDate string  `json:"from_date"`
            Stores []string  `json:"stores"`
         
    }
    
    // InventoryJobDetailResponse used by Catalog
    type InventoryJobDetailResponse struct {

        
            Status interface{}  `json:"status"`
            CancelledOn string  `json:"cancelled_on"`
            ID string  `json:"id"`
            CancelledBy UserDetail  `json:"cancelled_by"`
            URL string  `json:"url"`
            SellerID float64  `json:"seller_id"`
            CompletedOn string  `json:"completed_on"`
            CreatedOn string  `json:"created_on"`
            CreatedBy UserDetail  `json:"created_by"`
            ModifiedOn string  `json:"modified_on"`
            Type string  `json:"type"`
            Filters InventoryJobFilters  `json:"filters"`
            NotificationEmails []string  `json:"notification_emails"`
            TaskID string  `json:"task_id"`
         
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

        
            StoreID float64  `json:"store_id"`
            TraceID string  `json:"trace_id"`
            Tags []string  `json:"tags"`
            PriceEffective float64  `json:"price_effective"`
            TotalQuantity float64  `json:"total_quantity"`
            ExpirationDate string  `json:"expiration_date"`
            SellerIdentifier string  `json:"seller_identifier"`
            PriceMarked float64  `json:"price_marked"`
         
    }
    
    // InventoryRequestSchemaV2 used by Catalog
    type InventoryRequestSchemaV2 struct {

        
            Payload []InventoryPayload  `json:"payload"`
            Meta map[string]interface{}  `json:"meta"`
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

        
            Items []InventoryResponseItem  `json:"items"`
            Message string  `json:"message"`
         
    }
    
    // HsnUpsert used by Catalog
    type HsnUpsert struct {

        
            Tax2 float64  `json:"tax2"`
            CompanyID float64  `json:"company_id"`
            UID float64  `json:"uid"`
            IsActive bool  `json:"is_active"`
            Threshold2 float64  `json:"threshold2"`
            Hs2Code string  `json:"hs2_code"`
            HsnCode string  `json:"hsn_code"`
            TaxOnEsp bool  `json:"tax_on_esp"`
            Tax1 float64  `json:"tax1"`
            TaxOnMrp bool  `json:"tax_on_mrp"`
            Threshold1 float64  `json:"threshold1"`
         
    }
    
    // HsnCodesObject used by Catalog
    type HsnCodesObject struct {

        
            Tax2 float64  `json:"tax2"`
            CompanyID float64  `json:"company_id"`
            ID string  `json:"id"`
            Threshold2 float64  `json:"threshold2"`
            Hs2Code string  `json:"hs2_code"`
            HsnCode string  `json:"hsn_code"`
            ModifiedOn string  `json:"modified_on"`
            TaxOnEsp bool  `json:"tax_on_esp"`
            Tax1 float64  `json:"tax1"`
            TaxOnMrp bool  `json:"tax_on_mrp"`
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
    
    // TaxSlab used by Catalog
    type TaxSlab struct {

        
            Cess float64  `json:"cess"`
            Rate float64  `json:"rate"`
            Threshold float64  `json:"threshold"`
            EffectiveDate string  `json:"effective_date"`
         
    }
    
    // HSNDataInsertV2 used by Catalog
    type HSNDataInsertV2 struct {

        
            CountryCode string  `json:"country_code"`
            ReportingHsn string  `json:"reporting_hsn"`
            Type string  `json:"type"`
            HsnCode string  `json:"hsn_code"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Taxes []TaxSlab  `json:"taxes"`
            HsnCodeID string  `json:"hsn_code_id"`
            Description string  `json:"description"`
         
    }
    
    // PageResponse used by Catalog
    type PageResponse struct {

        
            HasPrevious bool  `json:"has_previous"`
            Size float64  `json:"size"`
            HasNext bool  `json:"has_next"`
            ItemTotal float64  `json:"item_total"`
            Current string  `json:"current"`
         
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

        
            Name string  `json:"name"`
            Discount string  `json:"discount"`
            Banners ImageUrls  `json:"banners"`
            Action Action  `json:"action"`
            Slug string  `json:"slug"`
            Logo Media2  `json:"logo"`
            Departments []string  `json:"departments"`
            UID float64  `json:"uid"`
         
    }
    
    // BrandListingResponse used by Catalog
    type BrandListingResponse struct {

        
            Items []BrandItem  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // Department used by Catalog
    type Department struct {

        
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Logo Media2  `json:"logo"`
            PriorityOrder float64  `json:"priority_order"`
            UID float64  `json:"uid"`
         
    }
    
    // DepartmentResponse used by Catalog
    type DepartmentResponse struct {

        
            Items []Department  `json:"items"`
         
    }
    
    // ThirdLevelChild used by Catalog
    type ThirdLevelChild struct {

        
            Name string  `json:"name"`
            Banners ImageUrls  `json:"banners"`
            Action Action  `json:"action"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
            Childs []map[string]interface{}  `json:"childs"`
         
    }
    
    // SecondLevelChild used by Catalog
    type SecondLevelChild struct {

        
            Name string  `json:"name"`
            Banners ImageUrls  `json:"banners"`
            Action Action  `json:"action"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
            Childs []ThirdLevelChild  `json:"childs"`
         
    }
    
    // Child used by Catalog
    type Child struct {

        
            Name string  `json:"name"`
            Banners ImageUrls  `json:"banners"`
            Action Action  `json:"action"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
            Childs []SecondLevelChild  `json:"childs"`
         
    }
    
    // CategoryItems used by Catalog
    type CategoryItems struct {

        
            Name string  `json:"name"`
            Banners ImageUrls  `json:"banners"`
            Action Action  `json:"action"`
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
            Childs []Child  `json:"childs"`
         
    }
    
    // DepartmentCategoryTree used by Catalog
    type DepartmentCategoryTree struct {

        
            Department string  `json:"department"`
            Items []CategoryItems  `json:"items"`
         
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
            Page Page  `json:"page"`
            Operators map[string]interface{}  `json:"operators"`
            Filters []ProductFilters  `json:"filters"`
            SortOn []ProductSortOn  `json:"sort_on"`
         
    }
    
    // ProductDetail used by Catalog
    type ProductDetail struct {

        
            RatingCount float64  `json:"rating_count"`
            Medias []Media  `json:"medias"`
            Tryouts []string  `json:"tryouts"`
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
            Highlights []string  `json:"highlights"`
            ItemType string  `json:"item_type"`
            Color string  `json:"color"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            HasVariant bool  `json:"has_variant"`
            Name string  `json:"name"`
            ImageNature string  `json:"image_nature"`
            ProductOnlineDate string  `json:"product_online_date"`
            Type string  `json:"type"`
            Rating float64  `json:"rating"`
            UID float64  `json:"uid"`
            PromoMeta map[string]interface{}  `json:"promo_meta"`
            ItemCode string  `json:"item_code"`
            Brand ProductBrand  `json:"brand"`
            Slug string  `json:"slug"`
            Similars []string  `json:"similars"`
            Attributes map[string]interface{}  `json:"attributes"`
            ShortDescription string  `json:"short_description"`
            Description string  `json:"description"`
         
    }
    
    // InventoryPage used by Catalog
    type InventoryPage struct {

        
            HasPrevious bool  `json:"has_previous"`
            NextID string  `json:"next_id"`
            HasNext bool  `json:"has_next"`
            ItemTotal float64  `json:"item_total"`
            Type string  `json:"type"`
         
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
            GroupID string  `json:"group_id"`
            Query ArticleQuery  `json:"query"`
            ArticleAssignment ArticleAssignment  `json:"article_assignment"`
         
    }
    
    // AssignStore used by Catalog
    type AssignStore struct {

        
            CompanyID float64  `json:"company_id"`
            Articles []AssignStoreArticle  `json:"articles"`
            ChannelType string  `json:"channel_type"`
            StoreIds []float64  `json:"store_ids"`
            AppID string  `json:"app_id"`
            ChannelIdentifier string  `json:"channel_identifier"`
            Pincode string  `json:"pincode"`
         
    }
    
    // ArticleAssignment1 used by Catalog
    type ArticleAssignment1 struct {

        
            Strategy string  `json:"strategy"`
            Level string  `json:"level"`
         
    }
    
    // StoreAssignResponse used by Catalog
    type StoreAssignResponse struct {

        
            Quantity float64  `json:"quantity"`
            StrategyWiseListing []map[string]interface{}  `json:"strategy_wise_listing"`
            Status bool  `json:"status"`
            Meta map[string]interface{}  `json:"meta"`
            StoreID float64  `json:"store_id"`
            Size string  `json:"size"`
            CompanyID float64  `json:"company_id"`
            GroupID string  `json:"group_id"`
            PriceEffective float64  `json:"price_effective"`
            Index float64  `json:"index"`
            ID string  `json:"_id"`
            ItemID float64  `json:"item_id"`
            SCity string  `json:"s_city"`
            StorePincode float64  `json:"store_pincode"`
            PriceMarked float64  `json:"price_marked"`
            ArticleAssignment ArticleAssignment1  `json:"article_assignment"`
            UID string  `json:"uid"`
         
    }
    
    // InvoiceCredSerializer used by Catalog
    type InvoiceCredSerializer struct {

        
            Username string  `json:"username"`
            Enabled bool  `json:"enabled"`
            Password string  `json:"password"`
         
    }
    
    // InvoiceDetailsSerializer used by Catalog
    type InvoiceDetailsSerializer struct {

        
            EInvoice InvoiceCredSerializer  `json:"e_invoice"`
            EWaybill InvoiceCredSerializer  `json:"e_waybill"`
         
    }
    
    // Document used by Catalog
    type Document struct {

        
            Verified bool  `json:"verified"`
            LegalName string  `json:"legal_name"`
            URL string  `json:"url"`
            Type string  `json:"type"`
            Value string  `json:"value"`
         
    }
    
    // SellerPhoneNumber used by Catalog
    type SellerPhoneNumber struct {

        
            CountryCode float64  `json:"country_code"`
            Number string  `json:"number"`
         
    }
    
    // UserSerializer1 used by Catalog
    type UserSerializer1 struct {

        
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
            Contact string  `json:"contact"`
         
    }
    
    // GetAddressSerializer used by Catalog
    type GetAddressSerializer struct {

        
            Longitude float64  `json:"longitude"`
            Address2 string  `json:"address2"`
            Latitude float64  `json:"latitude"`
            State string  `json:"state"`
            Landmark string  `json:"landmark"`
            CountryCode string  `json:"country_code"`
            Address1 string  `json:"address1"`
            City string  `json:"city"`
            Country string  `json:"country"`
            AddressType string  `json:"address_type"`
            Pincode float64  `json:"pincode"`
         
    }
    
    // UserSerializer2 used by Catalog
    type UserSerializer2 struct {

        
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
            Contact string  `json:"contact"`
         
    }
    
    // GetCompanySerializer used by Catalog
    type GetCompanySerializer struct {

        
            Name string  `json:"name"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            RejectReason string  `json:"reject_reason"`
            VerifiedOn string  `json:"verified_on"`
            BusinessType string  `json:"business_type"`
            CreatedOn string  `json:"created_on"`
            CreatedBy UserSerializer2  `json:"created_by"`
            ModifiedOn string  `json:"modified_on"`
            CompanyType string  `json:"company_type"`
            Stage string  `json:"stage"`
            ModifiedBy UserSerializer2  `json:"modified_by"`
            UID float64  `json:"uid"`
            VerifiedBy UserSerializer2  `json:"verified_by"`
         
    }
    
    // ProductReturnConfigSerializer used by Catalog
    type ProductReturnConfigSerializer struct {

        
            StoreUID float64  `json:"store_uid"`
            OnSameStore bool  `json:"on_same_store"`
         
    }
    
    // LocationTimingSerializer used by Catalog
    type LocationTimingSerializer struct {

        
            Hour float64  `json:"hour"`
            Minute float64  `json:"minute"`
         
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

        
            Order string  `json:"order"`
            Inventory string  `json:"inventory"`
         
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
            CreatedOn string  `json:"created_on"`
            Documents []Document  `json:"documents"`
            ModifiedOn string  `json:"modified_on"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            VerifiedBy UserSerializer1  `json:"verified_by"`
            Code string  `json:"code"`
            Company GetCompanySerializer  `json:"company"`
            ProductReturnConfig ProductReturnConfigSerializer  `json:"product_return_config"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ModifiedBy UserSerializer1  `json:"modified_by"`
            Timing []LocationDayWiseSerializer  `json:"timing"`
            Name string  `json:"name"`
            StoreType string  `json:"store_type"`
            Address GetAddressSerializer  `json:"address"`
            VerifiedOn string  `json:"verified_on"`
            DisplayName string  `json:"display_name"`
            IntegrationType LocationIntegrationType  `json:"integration_type"`
            NotificationEmails []string  `json:"notification_emails"`
            UID float64  `json:"uid"`
            Warnings map[string]interface{}  `json:"warnings"`
            PhoneNumber string  `json:"phone_number"`
            Manager LocationManagerSerializer  `json:"manager"`
            Stage string  `json:"stage"`
            CreatedBy UserSerializer1  `json:"created_by"`
         
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

        
            Name string  `json:"name"`
            IsActive bool  `json:"is_active"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Logo string  `json:"logo"`
            AppID string  `json:"app_id"`
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
    

    
    // Website used by CompanyProfile
    type Website struct {

        
            URL string  `json:"url"`
         
    }
    
    // BusinessDetails used by CompanyProfile
    type BusinessDetails struct {

        
            Website Website  `json:"website"`
         
    }
    
    // UserSerializer used by CompanyProfile
    type UserSerializer struct {

        
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
            Contact string  `json:"contact"`
         
    }
    
    // BusinessCountryInfo used by CompanyProfile
    type BusinessCountryInfo struct {

        
            Country string  `json:"country"`
            CountryCode string  `json:"country_code"`
         
    }
    
    // GetAddressSerializer used by CompanyProfile
    type GetAddressSerializer struct {

        
            Longitude float64  `json:"longitude"`
            CountryCode string  `json:"country_code"`
            Address1 string  `json:"address1"`
            Address2 string  `json:"address2"`
            Landmark string  `json:"landmark"`
            Country string  `json:"country"`
            AddressType string  `json:"address_type"`
            City string  `json:"city"`
            State string  `json:"state"`
            Pincode float64  `json:"pincode"`
            Latitude float64  `json:"latitude"`
         
    }
    
    // Document used by CompanyProfile
    type Document struct {

        
            Type string  `json:"type"`
            Value string  `json:"value"`
            URL string  `json:"url"`
            LegalName string  `json:"legal_name"`
            Verified bool  `json:"verified"`
         
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

        
            Enable bool  `json:"enable"`
            Rate float64  `json:"rate"`
            EffectiveDate string  `json:"effective_date"`
         
    }
    
    // GetCompanyProfileSerializerResponse used by CompanyProfile
    type GetCompanyProfileSerializerResponse struct {

        
            BusinessDetails BusinessDetails  `json:"business_details"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            BusinessCountryInfo BusinessCountryInfo  `json:"business_country_info"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Warnings map[string]interface{}  `json:"warnings"`
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
            VerifiedOn string  `json:"verified_on"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            Stage string  `json:"stage"`
            NotificationEmails []string  `json:"notification_emails"`
            BusinessType string  `json:"business_type"`
            Mode string  `json:"mode"`
            CreatedBy UserSerializer  `json:"created_by"`
            Documents []Document  `json:"documents"`
            ContactDetails ContactDetails  `json:"contact_details"`
            CompanyType string  `json:"company_type"`
            FranchiseEnabled bool  `json:"franchise_enabled"`
            Taxes []CompanyTaxesSerializer  `json:"taxes"`
            Name string  `json:"name"`
            UID float64  `json:"uid"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            BusinessInfo string  `json:"business_info"`
         
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

        
            Enable bool  `json:"enable"`
            Rate float64  `json:"rate"`
            EffectiveDate string  `json:"effective_date"`
         
    }
    
    // CreateUpdateAddressSerializer used by CompanyProfile
    type CreateUpdateAddressSerializer struct {

        
            Longitude float64  `json:"longitude"`
            CountryCode string  `json:"country_code"`
            Address1 string  `json:"address1"`
            Address2 string  `json:"address2"`
            Landmark string  `json:"landmark"`
            Country string  `json:"country"`
            State string  `json:"state"`
            City string  `json:"city"`
            AddressType string  `json:"address_type"`
            Pincode float64  `json:"pincode"`
            Latitude float64  `json:"latitude"`
         
    }
    
    // UpdateCompany used by CompanyProfile
    type UpdateCompany struct {

        
            FranchiseEnabled bool  `json:"franchise_enabled"`
            BusinessDetails BusinessDetails  `json:"business_details"`
            Taxes []CompanyTaxesSerializer1  `json:"taxes"`
            NotificationEmails []string  `json:"notification_emails"`
            RejectReason string  `json:"reject_reason"`
            Name string  `json:"name"`
            BusinessType string  `json:"business_type"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            BusinessInfo string  `json:"business_info"`
            Warnings map[string]interface{}  `json:"warnings"`
            Documents []Document  `json:"documents"`
            ContactDetails ContactDetails  `json:"contact_details"`
            CompanyType string  `json:"company_type"`
            Addresses []CreateUpdateAddressSerializer  `json:"addresses"`
         
    }
    
    // ProfileSuccessResponse used by CompanyProfile
    type ProfileSuccessResponse struct {

        
            UID float64  `json:"uid"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // DocumentsObj used by CompanyProfile
    type DocumentsObj struct {

        
            Verified float64  `json:"verified"`
            Pending float64  `json:"pending"`
         
    }
    
    // MetricsSerializer used by CompanyProfile
    type MetricsSerializer struct {

        
            Stage string  `json:"stage"`
            Store DocumentsObj  `json:"store"`
            StoreDocuments DocumentsObj  `json:"store_documents"`
            UID float64  `json:"uid"`
            Brand DocumentsObj  `json:"brand"`
            CompanyDocuments DocumentsObj  `json:"company_documents"`
            Product DocumentsObj  `json:"product"`
         
    }
    
    // BrandBannerSerializer used by CompanyProfile
    type BrandBannerSerializer struct {

        
            Portrait string  `json:"portrait"`
            Landscape string  `json:"landscape"`
         
    }
    
    // GetBrandResponseSerializer used by CompanyProfile
    type GetBrandResponseSerializer struct {

        
            VerifiedBy UserSerializer  `json:"verified_by"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Warnings map[string]interface{}  `json:"warnings"`
            ModifiedOn string  `json:"modified_on"`
            SlugKey string  `json:"slug_key"`
            RejectReason string  `json:"reject_reason"`
            CreatedOn string  `json:"created_on"`
            VerifiedOn string  `json:"verified_on"`
            Stage string  `json:"stage"`
            Description string  `json:"description"`
            Mode string  `json:"mode"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            CreatedBy UserSerializer  `json:"created_by"`
            Synonyms []string  `json:"synonyms"`
            Logo string  `json:"logo"`
            Name string  `json:"name"`
            UID float64  `json:"uid"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            Banner BrandBannerSerializer  `json:"banner"`
         
    }
    
    // CreateUpdateBrandRequestSerializer used by CompanyProfile
    type CreateUpdateBrandRequestSerializer struct {

        
            Synonyms []string  `json:"synonyms"`
            Logo string  `json:"logo"`
            BrandTier string  `json:"brand_tier"`
            Name string  `json:"name"`
            Description string  `json:"description"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            CompanyID float64  `json:"company_id"`
            UID float64  `json:"uid"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            Banner BrandBannerSerializer  `json:"banner"`
         
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

        
            MarketChannels []string  `json:"market_channels"`
            ModifiedOn string  `json:"modified_on"`
            Stage string  `json:"stage"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            NotificationEmails []string  `json:"notification_emails"`
            Details CompanyDetails  `json:"details"`
            Name string  `json:"name"`
            BusinessType string  `json:"business_type"`
            RejectReason string  `json:"reject_reason"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            BusinessCountryInfo BusinessCountryInfo  `json:"business_country_info"`
            UID float64  `json:"uid"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            CreatedBy UserSerializer  `json:"created_by"`
            CreatedOn string  `json:"created_on"`
            CompanyType string  `json:"company_type"`
            VerifiedOn string  `json:"verified_on"`
            Addresses []GetAddressSerializer  `json:"addresses"`
         
    }
    
    // CompanyBrandSerializer used by CompanyProfile
    type CompanyBrandSerializer struct {

        
            ModifiedOn string  `json:"modified_on"`
            Stage string  `json:"stage"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            RejectReason string  `json:"reject_reason"`
            Company CompanySerializer  `json:"company"`
            UID float64  `json:"uid"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            CreatedBy UserSerializer  `json:"created_by"`
            Brand GetBrandResponseSerializer  `json:"brand"`
            Warnings map[string]interface{}  `json:"warnings"`
            CreatedOn string  `json:"created_on"`
            VerifiedOn string  `json:"verified_on"`
         
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
    
    // InvoiceCredSerializer used by CompanyProfile
    type InvoiceCredSerializer struct {

        
            Enabled bool  `json:"enabled"`
            Password string  `json:"password"`
            Username string  `json:"username"`
         
    }
    
    // InvoiceDetailsSerializer used by CompanyProfile
    type InvoiceDetailsSerializer struct {

        
            EInvoice InvoiceCredSerializer  `json:"e_invoice"`
            EWaybill InvoiceCredSerializer  `json:"e_waybill"`
         
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

        
            HolidayType string  `json:"holiday_type"`
            Title string  `json:"title"`
            Date HolidayDateSerializer  `json:"date"`
         
    }
    
    // LocationManagerSerializer used by CompanyProfile
    type LocationManagerSerializer struct {

        
            MobileNo SellerPhoneNumber  `json:"mobile_no"`
            Email string  `json:"email"`
            Name string  `json:"name"`
         
    }
    
    // GetCompanySerializer used by CompanyProfile
    type GetCompanySerializer struct {

        
            ModifiedOn string  `json:"modified_on"`
            Stage string  `json:"stage"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            RejectReason string  `json:"reject_reason"`
            Name string  `json:"name"`
            BusinessType string  `json:"business_type"`
            UID float64  `json:"uid"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            CreatedBy UserSerializer  `json:"created_by"`
            CreatedOn string  `json:"created_on"`
            CompanyType string  `json:"company_type"`
            VerifiedOn string  `json:"verified_on"`
            Addresses []GetAddressSerializer  `json:"addresses"`
         
    }
    
    // LocationTimingSerializer used by CompanyProfile
    type LocationTimingSerializer struct {

        
            Minute float64  `json:"minute"`
            Hour float64  `json:"hour"`
         
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

        
            VerifiedBy UserSerializer  `json:"verified_by"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Warnings map[string]interface{}  `json:"warnings"`
            CreditNote bool  `json:"credit_note"`
            ModifiedOn string  `json:"modified_on"`
            PhoneNumber string  `json:"phone_number"`
            StoreType string  `json:"store_type"`
            Code string  `json:"code"`
            DisplayName string  `json:"display_name"`
            CreatedOn string  `json:"created_on"`
            VerifiedOn string  `json:"verified_on"`
            GstCredentials InvoiceDetailsSerializer  `json:"gst_credentials"`
            Stage string  `json:"stage"`
            NotificationEmails []string  `json:"notification_emails"`
            ProductReturnConfig ProductReturnConfigSerializer  `json:"product_return_config"`
            CreatedBy UserSerializer  `json:"created_by"`
            Documents []Document  `json:"documents"`
            Holiday []HolidaySchemaSerializer  `json:"holiday"`
            Manager LocationManagerSerializer  `json:"manager"`
            Address GetAddressSerializer  `json:"address"`
            Name string  `json:"name"`
            Company GetCompanySerializer  `json:"company"`
            UID float64  `json:"uid"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            Timing []LocationDayWiseSerializer  `json:"timing"`
            AutoInvoice bool  `json:"auto_invoice"`
         
    }
    
    // LocationListSerializer used by CompanyProfile
    type LocationListSerializer struct {

        
            Page Page  `json:"page"`
            Items []GetLocationSerializer  `json:"items"`
         
    }
    
    // AddressSerializer used by CompanyProfile
    type AddressSerializer struct {

        
            Longitude float64  `json:"longitude"`
            CountryCode string  `json:"country_code"`
            Address1 string  `json:"address1"`
            Address2 string  `json:"address2"`
            Landmark string  `json:"landmark"`
            Country string  `json:"country"`
            AddressType string  `json:"address_type"`
            City string  `json:"city"`
            State string  `json:"state"`
            Pincode float64  `json:"pincode"`
            Latitude float64  `json:"latitude"`
         
    }
    
    // LocationSerializer used by CompanyProfile
    type LocationSerializer struct {

        
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Warnings map[string]interface{}  `json:"warnings"`
            CreditNote bool  `json:"credit_note"`
            StoreType string  `json:"store_type"`
            Code string  `json:"code"`
            DisplayName string  `json:"display_name"`
            GstCredentials InvoiceDetailsSerializer  `json:"gst_credentials"`
            Stage string  `json:"stage"`
            NotificationEmails []string  `json:"notification_emails"`
            ProductReturnConfig ProductReturnConfigSerializer  `json:"product_return_config"`
            Documents []Document  `json:"documents"`
            Holiday []HolidaySchemaSerializer  `json:"holiday"`
            Manager LocationManagerSerializer  `json:"manager"`
            Address AddressSerializer  `json:"address"`
            Name string  `json:"name"`
            Company float64  `json:"company"`
            UID float64  `json:"uid"`
            Timing []LocationDayWiseSerializer  `json:"timing"`
            AutoInvoice bool  `json:"auto_invoice"`
         
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
    
    // UsesRemaining used by Cart
    type UsesRemaining struct {

        
            User float64  `json:"user"`
            Total float64  `json:"total"`
            App float64  `json:"app"`
         
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
            Networks []string  `json:"networks"`
            Types []string  `json:"types"`
            Codes []string  `json:"codes"`
         
    }
    
    // BulkBundleRestriction used by Cart
    type BulkBundleRestriction struct {

        
            MultiStoreAllowed bool  `json:"multi_store_allowed"`
         
    }
    
    // Restrictions used by Cart
    type Restrictions struct {

        
            Uses UsesRestriction  `json:"uses"`
            PriceRange PriceRange  `json:"price_range"`
            CouponAllowed bool  `json:"coupon_allowed"`
            UserGroups []float64  `json:"user_groups"`
            Platforms []string  `json:"platforms"`
            PostOrder PostOrder  `json:"post_order"`
            Payments map[string]PaymentModes  `json:"payments"`
            OrderingStores []float64  `json:"ordering_stores"`
            BulkBundle BulkBundleRestriction  `json:"bulk_bundle"`
         
    }
    
    // Validity used by Cart
    type Validity struct {

        
            Priority float64  `json:"priority"`
         
    }
    
    // State used by Cart
    type State struct {

        
            IsArchived bool  `json:"is_archived"`
            IsDisplay bool  `json:"is_display"`
            IsPublic bool  `json:"is_public"`
         
    }
    
    // Validation used by Cart
    type Validation struct {

        
            AppID []string  `json:"app_id"`
            Anonymous bool  `json:"anonymous"`
            UserRegisteredAfter string  `json:"user_registered_after"`
         
    }
    
    // CouponSchedule used by Cart
    type CouponSchedule struct {

        
            Start string  `json:"start"`
            End string  `json:"end"`
            NextSchedule []map[string]interface{}  `json:"next_schedule"`
            Cron string  `json:"cron"`
            Duration float64  `json:"duration"`
         
    }
    
    // Identifier used by Cart
    type Identifier struct {

        
            ExcludeBrandID []float64  `json:"exclude_brand_id"`
            CompanyID []float64  `json:"company_id"`
            UserID []string  `json:"user_id"`
            CategoryID []float64  `json:"category_id"`
            StoreID []float64  `json:"store_id"`
            ArticleID []string  `json:"article_id"`
            CollectionID []string  `json:"collection_id"`
            BrandID []float64  `json:"brand_id"`
            ItemID []float64  `json:"item_id"`
         
    }
    
    // RuleDefinition used by Cart
    type RuleDefinition struct {

        
            CurrencyCode string  `json:"currency_code"`
            Type string  `json:"type"`
            ValueType string  `json:"value_type"`
            AutoApply bool  `json:"auto_apply"`
            IsExact bool  `json:"is_exact"`
            CalculateOn string  `json:"calculate_on"`
            Scope []string  `json:"scope"`
            ApplicableOn string  `json:"applicable_on"`
         
    }
    
    // CouponAction used by Cart
    type CouponAction struct {

        
            TxnMode string  `json:"txn_mode"`
            ActionDate string  `json:"action_date"`
         
    }
    
    // CouponAuthor used by Cart
    type CouponAuthor struct {

        
            ModifiedBy string  `json:"modified_by"`
            CreatedBy string  `json:"created_by"`
         
    }
    
    // Rule used by Cart
    type Rule struct {

        
            Value float64  `json:"value"`
            Min float64  `json:"min"`
            Max float64  `json:"max"`
            DiscountQty float64  `json:"discount_qty"`
            Key float64  `json:"key"`
         
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
    
    // CouponAdd used by Cart
    type CouponAdd struct {

        
            TypeSlug string  `json:"type_slug"`
            Code string  `json:"code"`
            DateMeta CouponDateMeta  `json:"date_meta"`
            Ownership Ownership  `json:"ownership"`
            Restrictions Restrictions  `json:"restrictions"`
            Validity Validity  `json:"validity"`
            State State  `json:"state"`
            Validation Validation  `json:"validation"`
            Schedule CouponSchedule  `json:"_schedule"`
            Identifiers Identifier  `json:"identifiers"`
            RuleDefinition RuleDefinition  `json:"rule_definition"`
            Action CouponAction  `json:"action"`
            Tags []string  `json:"tags"`
            Author CouponAuthor  `json:"author"`
            Rule []Rule  `json:"rule"`
            DisplayMeta DisplayMeta  `json:"display_meta"`
         
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

        
            TypeSlug string  `json:"type_slug"`
            Code string  `json:"code"`
            DateMeta CouponDateMeta  `json:"date_meta"`
            Ownership Ownership  `json:"ownership"`
            Restrictions Restrictions  `json:"restrictions"`
            Validity Validity  `json:"validity"`
            State State  `json:"state"`
            Validation Validation  `json:"validation"`
            Schedule CouponSchedule  `json:"_schedule"`
            Identifiers Identifier  `json:"identifiers"`
            RuleDefinition RuleDefinition  `json:"rule_definition"`
            Action CouponAction  `json:"action"`
            Tags []string  `json:"tags"`
            Author CouponAuthor  `json:"author"`
            Rule []Rule  `json:"rule"`
            DisplayMeta DisplayMeta  `json:"display_meta"`
         
    }
    
    // CouponPartialUpdate used by Cart
    type CouponPartialUpdate struct {

        
            Archive bool  `json:"archive"`
            Schedule CouponSchedule  `json:"schedule"`
         
    }
    
    // CompareObject used by Cart
    type CompareObject struct {

        
            GreaterThanEquals float64  `json:"greater_than_equals"`
            LessThan float64  `json:"less_than"`
            GreaterThan float64  `json:"greater_than"`
            Equals float64  `json:"equals"`
            LessThanEquals float64  `json:"less_than_equals"`
         
    }
    
    // ItemCriteria used by Cart
    type ItemCriteria struct {

        
            ItemExcludeCompany []float64  `json:"item_exclude_company"`
            ItemExcludeSku []string  `json:"item_exclude_sku"`
            ItemCompany []float64  `json:"item_company"`
            ItemBrand []float64  `json:"item_brand"`
            CartQuantity CompareObject  `json:"cart_quantity"`
            ItemSize []string  `json:"item_size"`
            ItemExcludeID []float64  `json:"item_exclude_id"`
            CartUniqueItemAmount CompareObject  `json:"cart_unique_item_amount"`
            ItemSku []string  `json:"item_sku"`
            CartUniqueItemQuantity CompareObject  `json:"cart_unique_item_quantity"`
            ItemStore []float64  `json:"item_store"`
            CartTotal CompareObject  `json:"cart_total"`
            BuyRules []string  `json:"buy_rules"`
            ItemCategory []float64  `json:"item_category"`
            AllItems bool  `json:"all_items"`
            ItemExcludeBrand []float64  `json:"item_exclude_brand"`
            ItemExcludeStore []float64  `json:"item_exclude_store"`
            ItemExcludeCategory []float64  `json:"item_exclude_category"`
            AvailableZones []string  `json:"available_zones"`
            ItemID []float64  `json:"item_id"`
         
    }
    
    // DiscountOffer used by Cart
    type DiscountOffer struct {

        
            PartialCanRet bool  `json:"partial_can_ret"`
            DiscountPrice float64  `json:"discount_price"`
            Code string  `json:"code"`
            MaxDiscountAmount float64  `json:"max_discount_amount"`
            DiscountPercentage float64  `json:"discount_percentage"`
            MinOfferQuantity float64  `json:"min_offer_quantity"`
            MaxUsagePerTransaction float64  `json:"max_usage_per_transaction"`
            DiscountAmount float64  `json:"discount_amount"`
            MaxOfferQuantity float64  `json:"max_offer_quantity"`
            ApportionDiscount bool  `json:"apportion_discount"`
         
    }
    
    // DiscountRule used by Cart
    type DiscountRule struct {

        
            BuyCondition string  `json:"buy_condition"`
            DiscountType string  `json:"discount_type"`
            ItemCriteria ItemCriteria  `json:"item_criteria"`
            Offer DiscountOffer  `json:"offer"`
         
    }
    
    // DisplayMeta1 used by Cart
    type DisplayMeta1 struct {

        
            Description string  `json:"description"`
            Name string  `json:"name"`
            OfferText string  `json:"offer_text"`
         
    }
    
    // PromotionDateMeta used by Cart
    type PromotionDateMeta struct {

        
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
         
    }
    
    // PromotionAction used by Cart
    type PromotionAction struct {

        
            ActionDate string  `json:"action_date"`
            ActionType string  `json:"action_type"`
         
    }
    
    // Ownership1 used by Cart
    type Ownership1 struct {

        
            PayableCategory string  `json:"payable_category"`
            PayableBy string  `json:"payable_by"`
         
    }
    
    // UsesRemaining1 used by Cart
    type UsesRemaining1 struct {

        
            User float64  `json:"user"`
            Total float64  `json:"total"`
         
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
            Type string  `json:"type"`
            Codes []string  `json:"codes"`
         
    }
    
    // Restrictions1 used by Cart
    type Restrictions1 struct {

        
            Uses UsesRestriction1  `json:"uses"`
            OrderQuantity float64  `json:"order_quantity"`
            AnonymousUsers bool  `json:"anonymous_users"`
            UserGroups []float64  `json:"user_groups"`
            UserID []string  `json:"user_id"`
            Platforms []string  `json:"platforms"`
            PostOrder PostOrder1  `json:"post_order"`
            UserRegistered UserRegistered  `json:"user_registered"`
            Payments []PromotionPaymentModes  `json:"payments"`
            OrderingStores []float64  `json:"ordering_stores"`
         
    }
    
    // PromotionSchedule used by Cart
    type PromotionSchedule struct {

        
            Start string  `json:"start"`
            End string  `json:"end"`
            NextSchedule []map[string]interface{}  `json:"next_schedule"`
            Published bool  `json:"published"`
            Cron string  `json:"cron"`
            Duration float64  `json:"duration"`
         
    }
    
    // PromotionAuthor used by Cart
    type PromotionAuthor struct {

        
            ModifiedBy string  `json:"modified_by"`
            CreatedBy string  `json:"created_by"`
         
    }
    
    // Visibility used by Cart
    type Visibility struct {

        
            Pdp bool  `json:"pdp"`
            CouponList bool  `json:"coupon_list"`
         
    }
    
    // PromotionListItem used by Cart
    type PromotionListItem struct {

        
            PromotionType string  `json:"promotion_type"`
            ApplyAllDiscount bool  `json:"apply_all_discount"`
            Code string  `json:"code"`
            DiscountRules []DiscountRule  `json:"discount_rules"`
            ApplyExclusive string  `json:"apply_exclusive"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            DisplayMeta DisplayMeta1  `json:"display_meta"`
            Currency string  `json:"currency"`
            ApplyPriority float64  `json:"apply_priority"`
            Stackable bool  `json:"stackable"`
            DateMeta PromotionDateMeta  `json:"date_meta"`
            PostOrderAction PromotionAction  `json:"post_order_action"`
            Mode string  `json:"mode"`
            BuyRules map[string]ItemCriteria  `json:"buy_rules"`
            Ownership Ownership1  `json:"ownership"`
            Restrictions Restrictions1  `json:"restrictions"`
            Schedule PromotionSchedule  `json:"_schedule"`
            CalculateOn string  `json:"calculate_on"`
            Author PromotionAuthor  `json:"author"`
            PromoGroup string  `json:"promo_group"`
            Visiblility Visibility  `json:"visiblility"`
            ApplicationID string  `json:"application_id"`
         
    }
    
    // PromotionsResponse used by Cart
    type PromotionsResponse struct {

        
            Page Page  `json:"page"`
            Items PromotionListItem  `json:"items"`
         
    }
    
    // PromotionAdd used by Cart
    type PromotionAdd struct {

        
            PromotionType string  `json:"promotion_type"`
            ApplyAllDiscount bool  `json:"apply_all_discount"`
            Code string  `json:"code"`
            DiscountRules []DiscountRule  `json:"discount_rules"`
            ApplyExclusive string  `json:"apply_exclusive"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            DisplayMeta DisplayMeta1  `json:"display_meta"`
            Currency string  `json:"currency"`
            ApplyPriority float64  `json:"apply_priority"`
            Stackable bool  `json:"stackable"`
            DateMeta PromotionDateMeta  `json:"date_meta"`
            PostOrderAction PromotionAction  `json:"post_order_action"`
            Mode string  `json:"mode"`
            BuyRules map[string]ItemCriteria  `json:"buy_rules"`
            Ownership Ownership1  `json:"ownership"`
            Restrictions Restrictions1  `json:"restrictions"`
            Schedule PromotionSchedule  `json:"_schedule"`
            CalculateOn string  `json:"calculate_on"`
            Author PromotionAuthor  `json:"author"`
            PromoGroup string  `json:"promo_group"`
            Visiblility Visibility  `json:"visiblility"`
            ApplicationID string  `json:"application_id"`
         
    }
    
    // PromotionUpdate used by Cart
    type PromotionUpdate struct {

        
            PromotionType string  `json:"promotion_type"`
            ApplyAllDiscount bool  `json:"apply_all_discount"`
            Code string  `json:"code"`
            DiscountRules []DiscountRule  `json:"discount_rules"`
            ApplyExclusive string  `json:"apply_exclusive"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            DisplayMeta DisplayMeta1  `json:"display_meta"`
            Currency string  `json:"currency"`
            ApplyPriority float64  `json:"apply_priority"`
            Stackable bool  `json:"stackable"`
            DateMeta PromotionDateMeta  `json:"date_meta"`
            PostOrderAction PromotionAction  `json:"post_order_action"`
            Mode string  `json:"mode"`
            BuyRules map[string]ItemCriteria  `json:"buy_rules"`
            Ownership Ownership1  `json:"ownership"`
            Restrictions Restrictions1  `json:"restrictions"`
            Schedule PromotionSchedule  `json:"_schedule"`
            CalculateOn string  `json:"calculate_on"`
            Author PromotionAuthor  `json:"author"`
            PromoGroup string  `json:"promo_group"`
            Visiblility Visibility  `json:"visiblility"`
            ApplicationID string  `json:"application_id"`
         
    }
    
    // PromotionPartialUpdate used by Cart
    type PromotionPartialUpdate struct {

        
            Archive bool  `json:"archive"`
            Schedule PromotionSchedule  `json:"schedule"`
         
    }
    
    // ActivePromosResponse used by Cart
    type ActivePromosResponse struct {

        
            Type string  `json:"type"`
            Subtitle string  `json:"subtitle"`
            CreatedOn string  `json:"created_on"`
            EntityType string  `json:"entity_type"`
            Example string  `json:"example"`
            ModifiedOn string  `json:"modified_on"`
            EntitySlug string  `json:"entity_slug"`
            Description string  `json:"description"`
            IsHidden bool  `json:"is_hidden"`
            Title string  `json:"title"`
         
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

        
            BulkCoupons bool  `json:"bulk_coupons"`
            GiftPricing float64  `json:"gift_pricing"`
            DeliveryCharges DeliveryCharges  `json:"delivery_charges"`
            RevenueEngineCoupon bool  `json:"revenue_engine_coupon"`
            MinCartValue float64  `json:"min_cart_value"`
            GiftDisplayText string  `json:"gift_display_text"`
            Enabled bool  `json:"enabled"`
            MaxCartItems float64  `json:"max_cart_items"`
         
    }
    
    // CartMetaConfigAdd used by Cart
    type CartMetaConfigAdd struct {

        
            BulkCoupons bool  `json:"bulk_coupons"`
            GiftPricing float64  `json:"gift_pricing"`
            DeliveryCharges DeliveryCharges  `json:"delivery_charges"`
            RevenueEngineCoupon bool  `json:"revenue_engine_coupon"`
            MinCartValue float64  `json:"min_cart_value"`
            GiftDisplayText string  `json:"gift_display_text"`
            Enabled bool  `json:"enabled"`
            MaxCartItems float64  `json:"max_cart_items"`
         
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
    
    // CouponBreakup used by Cart
    type CouponBreakup struct {

        
            Type string  `json:"type"`
            Code string  `json:"code"`
            Value float64  `json:"value"`
            Message string  `json:"message"`
            IsApplied bool  `json:"is_applied"`
            CouponValue float64  `json:"coupon_value"`
            UID string  `json:"uid"`
            MaxDiscountValue float64  `json:"max_discount_value"`
            CouponType string  `json:"coupon_type"`
            MinimumCartValue float64  `json:"minimum_cart_value"`
            SubTitle string  `json:"sub_title"`
            Description string  `json:"description"`
            Title string  `json:"title"`
         
    }
    
    // RawBreakup used by Cart
    type RawBreakup struct {

        
            CodCharge float64  `json:"cod_charge"`
            Vog float64  `json:"vog"`
            GstCharges float64  `json:"gst_charges"`
            Coupon float64  `json:"coupon"`
            Discount float64  `json:"discount"`
            Total float64  `json:"total"`
            YouSaved float64  `json:"you_saved"`
            Subtotal float64  `json:"subtotal"`
            FyndCash float64  `json:"fynd_cash"`
            ConvenienceFee float64  `json:"convenience_fee"`
            MrpTotal float64  `json:"mrp_total"`
            DeliveryCharge float64  `json:"delivery_charge"`
         
    }
    
    // LoyaltyPoints used by Cart
    type LoyaltyPoints struct {

        
            Description string  `json:"description"`
            IsApplied bool  `json:"is_applied"`
            Applicable float64  `json:"applicable"`
            Total float64  `json:"total"`
         
    }
    
    // DisplayBreakup used by Cart
    type DisplayBreakup struct {

        
            CurrencySymbol string  `json:"currency_symbol"`
            CurrencyCode string  `json:"currency_code"`
            Value float64  `json:"value"`
            Message []string  `json:"message"`
            Display string  `json:"display"`
            Key string  `json:"key"`
         
    }
    
    // CartBreakup used by Cart
    type CartBreakup struct {

        
            Coupon CouponBreakup  `json:"coupon"`
            Raw RawBreakup  `json:"raw"`
            LoyaltyPoints LoyaltyPoints  `json:"loyalty_points"`
            Display []DisplayBreakup  `json:"display"`
         
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

        
            URL string  `json:"url"`
            SecureURL string  `json:"secure_url"`
            AspectRatio string  `json:"aspect_ratio"`
         
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
    
    // CartProduct used by Cart
    type CartProduct struct {

        
            Name string  `json:"name"`
            Type string  `json:"type"`
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
            Categories []CategoryInfo  `json:"categories"`
            Images []ProductImage  `json:"images"`
            Brand BaseInfo  `json:"brand"`
            Action ProductAction  `json:"action"`
         
    }
    
    // ProductPrice used by Cart
    type ProductPrice struct {

        
            CurrencySymbol string  `json:"currency_symbol"`
            CurrencyCode string  `json:"currency_code"`
            Selling float64  `json:"selling"`
            Marked float64  `json:"marked"`
            Effective float64  `json:"effective"`
            AddOn float64  `json:"add_on"`
         
    }
    
    // ProductPriceInfo used by Cart
    type ProductPriceInfo struct {

        
            Converted ProductPrice  `json:"converted"`
            Base ProductPrice  `json:"base"`
         
    }
    
    // ProductAvailabilitySize used by Cart
    type ProductAvailabilitySize struct {

        
            Value string  `json:"value"`
            IsAvailable bool  `json:"is_available"`
            Display string  `json:"display"`
         
    }
    
    // ProductAvailability used by Cart
    type ProductAvailability struct {

        
            Deliverable bool  `json:"deliverable"`
            OutOfStock bool  `json:"out_of_stock"`
            IsValid bool  `json:"is_valid"`
            AvailableSizes []ProductAvailabilitySize  `json:"available_sizes"`
            OtherStoreQuantity float64  `json:"other_store_quantity"`
            Sizes []string  `json:"sizes"`
         
    }
    
    // CartProductIdentifer used by Cart
    type CartProductIdentifer struct {

        
            Identifier string  `json:"identifier"`
         
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
    
    // ProductArticle used by Cart
    type ProductArticle struct {

        
            ProductGroupTags []string  `json:"product_group_tags"`
            Type string  `json:"type"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            Size string  `json:"size"`
            SellerIdentifier string  `json:"seller_identifier"`
            UID string  `json:"uid"`
            Quantity float64  `json:"quantity"`
            Store BaseInfo  `json:"store"`
            Seller BaseInfo  `json:"seller"`
            Price ArticlePriceInfo  `json:"price"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            Identifier map[string]interface{}  `json:"identifier"`
         
    }
    
    // BuyRules used by Cart
    type BuyRules struct {

        
            CartConditions map[string]interface{}  `json:"cart_conditions"`
            ItemCriteria map[string]interface{}  `json:"item_criteria"`
         
    }
    
    // FreeGiftItem used by Cart
    type FreeGiftItem struct {

        
            ItemBrandName string  `json:"item_brand_name"`
            ItemName string  `json:"item_name"`
            ItemPriceDetails map[string]interface{}  `json:"item_price_details"`
            ItemImagesURL []string  `json:"item_images_url"`
            ItemID float64  `json:"item_id"`
            ItemSlug string  `json:"item_slug"`
         
    }
    
    // AppliedFreeArticles used by Cart
    type AppliedFreeArticles struct {

        
            ParentItemIdentifier string  `json:"parent_item_identifier"`
            FreeGiftItemDetails FreeGiftItem  `json:"free_gift_item_details"`
            ArticleID string  `json:"article_id"`
            Quantity float64  `json:"quantity"`
         
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
            RawOffer map[string]interface{}  `json:"raw_offer"`
            Offer map[string]interface{}  `json:"offer"`
         
    }
    
    // AppliedPromotion used by Cart
    type AppliedPromotion struct {

        
            PromotionType string  `json:"promotion_type"`
            BuyRules []BuyRules  `json:"buy_rules"`
            ArticleQuantity float64  `json:"article_quantity"`
            PromoID string  `json:"promo_id"`
            AppliedFreeArticles []AppliedFreeArticles  `json:"applied_free_articles"`
            Ownership Ownership2  `json:"ownership"`
            DiscountRules []DiscountRulesApp  `json:"discount_rules"`
            PromotionGroup string  `json:"promotion_group"`
            PromotionName string  `json:"promotion_name"`
            Amount float64  `json:"amount"`
            OfferText string  `json:"offer_text"`
            MrpPromotion bool  `json:"mrp_promotion"`
         
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
    
    // CartProductInfo used by Cart
    type CartProductInfo struct {

        
            PromoMeta PromoMeta  `json:"promo_meta"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            BulkOffer map[string]interface{}  `json:"bulk_offer"`
            Product CartProduct  `json:"product"`
            Message string  `json:"message"`
            CouponMessage string  `json:"coupon_message"`
            Quantity float64  `json:"quantity"`
            PricePerUnit ProductPriceInfo  `json:"price_per_unit"`
            Availability ProductAvailability  `json:"availability"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
            Article ProductArticle  `json:"article"`
            Moq map[string]interface{}  `json:"moq"`
            PromotionsApplied []AppliedPromotion  `json:"promotions_applied"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            Discount string  `json:"discount"`
            Price ProductPriceInfo  `json:"price"`
            Key string  `json:"key"`
            IsSet bool  `json:"is_set"`
         
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
            Success bool  `json:"success"`
            Errors map[string]interface{}  `json:"errors"`
         
    }
    
    // ShippingAddress used by Cart
    type ShippingAddress struct {

        
            Name string  `json:"name"`
            Pincode float64  `json:"pincode"`
            CountryPhoneCode string  `json:"country_phone_code"`
            AreaCode string  `json:"area_code"`
            Phone float64  `json:"phone"`
            Email string  `json:"email"`
            CountryIsoCode string  `json:"country_iso_code"`
            State string  `json:"state"`
            Landmark string  `json:"landmark"`
            AreaCodeSlug string  `json:"area_code_slug"`
            AddressType string  `json:"address_type"`
            Country string  `json:"country"`
            CountryCode string  `json:"country_code"`
            Area string  `json:"area"`
            City string  `json:"city"`
            Address string  `json:"address"`
            Meta map[string]interface{}  `json:"meta"`
         
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
            BreakupValues CartBreakup  `json:"breakup_values"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            Items []CartProductInfo  `json:"items"`
         
    }
    
    // OpenApiFiles used by Cart
    type OpenApiFiles struct {

        
            Values []string  `json:"values"`
            Key string  `json:"key"`
         
    }
    
    // MultiTenderPaymentMeta used by Cart
    type MultiTenderPaymentMeta struct {

        
            PaymentID string  `json:"payment_id"`
            CurrentStatus string  `json:"current_status"`
            PaymentGateway string  `json:"payment_gateway"`
            OrderID string  `json:"order_id"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
         
    }
    
    // MultiTenderPaymentMethod used by Cart
    type MultiTenderPaymentMethod struct {

        
            Amount float64  `json:"amount"`
            Name string  `json:"name"`
            Meta MultiTenderPaymentMeta  `json:"meta"`
            Mode string  `json:"mode"`
         
    }
    
    // CartItemMeta used by Cart
    type CartItemMeta struct {

        
            PrimaryItem bool  `json:"primary_item"`
            GroupID string  `json:"group_id"`
         
    }
    
    // OpenApiOrderItem used by Cart
    type OpenApiOrderItem struct {

        
            CodCharges float64  `json:"cod_charges"`
            Size string  `json:"size"`
            CashbackApplied float64  `json:"cashback_applied"`
            DeliveryCharges float64  `json:"delivery_charges"`
            PriceMarked float64  `json:"price_marked"`
            Files []OpenApiFiles  `json:"files"`
            PriceEffective float64  `json:"price_effective"`
            Quantity float64  `json:"quantity"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            AmountPaid float64  `json:"amount_paid"`
            Discount float64  `json:"discount"`
            ProductID float64  `json:"product_id"`
            PaymentMethods []MultiTenderPaymentMethod  `json:"payment_methods"`
            EmployeeDiscount float64  `json:"employee_discount"`
            Meta CartItemMeta  `json:"meta"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
         
    }
    
    // OpenApiPlatformCheckoutReq used by Cart
    type OpenApiPlatformCheckoutReq struct {

        
            CodCharges float64  `json:"cod_charges"`
            CurrencyCode string  `json:"currency_code"`
            Gstin string  `json:"gstin"`
            PaymentMode string  `json:"payment_mode"`
            CashbackApplied float64  `json:"cashback_applied"`
            CouponValue float64  `json:"coupon_value"`
            DeliveryCharges float64  `json:"delivery_charges"`
            OrderID string  `json:"order_id"`
            BillingAddress ShippingAddress  `json:"billing_address"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            Comment string  `json:"comment"`
            CartItems []OpenApiOrderItem  `json:"cart_items"`
            PaymentMethods []MultiTenderPaymentMethod  `json:"payment_methods"`
            CartValue float64  `json:"cart_value"`
            Coupon string  `json:"coupon"`
            Files []OpenApiFiles  `json:"files"`
            CouponCode string  `json:"coupon_code"`
            ShippingAddress ShippingAddress  `json:"shipping_address"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            EmployeeDiscount map[string]interface{}  `json:"employee_discount"`
         
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

        
            ID string  `json:"_id"`
            CodCharges map[string]interface{}  `json:"cod_charges"`
            BuyNow bool  `json:"buy_now"`
            Gstin string  `json:"gstin"`
            PaymentMode string  `json:"payment_mode"`
            UserID string  `json:"user_id"`
            Comment string  `json:"comment"`
            Payments map[string]interface{}  `json:"payments"`
            PickUpCustomerDetails map[string]interface{}  `json:"pick_up_customer_details"`
            PaymentMethods []map[string]interface{}  `json:"payment_methods"`
            CartValue float64  `json:"cart_value"`
            FyndCredits map[string]interface{}  `json:"fynd_credits"`
            BulkCouponDiscount float64  `json:"bulk_coupon_discount"`
            Articles []map[string]interface{}  `json:"articles"`
            IsDefault bool  `json:"is_default"`
            IsArchive bool  `json:"is_archive"`
            LastModified string  `json:"last_modified"`
            FcIndexMap []float64  `json:"fc_index_map"`
            IsActive bool  `json:"is_active"`
            AppID string  `json:"app_id"`
            DeliveryCharges map[string]interface{}  `json:"delivery_charges"`
            Promotion map[string]interface{}  `json:"promotion"`
            OrderID string  `json:"order_id"`
            UID float64  `json:"uid"`
            ExpireAt string  `json:"expire_at"`
            MergeQty bool  `json:"merge_qty"`
            Cashback map[string]interface{}  `json:"cashback"`
            CheckoutMode string  `json:"checkout_mode"`
            Meta map[string]interface{}  `json:"meta"`
            CreatedOn string  `json:"created_on"`
            Shipments []map[string]interface{}  `json:"shipments"`
            Coupon map[string]interface{}  `json:"coupon"`
            Discount float64  `json:"discount"`
         
    }
    
    // AbandonedCartResponse used by Cart
    type AbandonedCartResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
            Page Page  `json:"page"`
            Result map[string]interface{}  `json:"result"`
            Items []AbandonedCart  `json:"items"`
         
    }
    
    // CartCurrency used by Cart
    type CartCurrency struct {

        
            Symbol string  `json:"symbol"`
            Code string  `json:"code"`
         
    }
    
    // PaymentSelectionLock used by Cart
    type PaymentSelectionLock struct {

        
            DefaultOptions string  `json:"default_options"`
            PaymentIdentifier string  `json:"payment_identifier"`
            Enabled bool  `json:"enabled"`
         
    }
    
    // CartDetailResponse used by Cart
    type CartDetailResponse struct {

        
            Currency CartCurrency  `json:"currency"`
            BuyNow bool  `json:"buy_now"`
            Gstin string  `json:"gstin"`
            ID string  `json:"id"`
            Message string  `json:"message"`
            IsValid bool  `json:"is_valid"`
            Comment string  `json:"comment"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            AppliedPromoDetails []AppliedPromotion  `json:"applied_promo_details"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            LastModified string  `json:"last_modified"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            CouponText string  `json:"coupon_text"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            CheckoutMode string  `json:"checkout_mode"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            Items []CartProductInfo  `json:"items"`
         
    }
    
    // AddProductCart used by Cart
    type AddProductCart struct {

        
            Pos bool  `json:"pos"`
            ProductGroupTags []string  `json:"product_group_tags"`
            ItemSize string  `json:"item_size"`
            ParentItemIdentifiers []map[string]string  `json:"parent_item_identifiers"`
            Quantity float64  `json:"quantity"`
            Display string  `json:"display"`
            StoreID float64  `json:"store_id"`
            ArticleID string  `json:"article_id"`
            SellerID float64  `json:"seller_id"`
            ItemID float64  `json:"item_id"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            ArticleAssignment map[string]interface{}  `json:"article_assignment"`
         
    }
    
    // AddCartRequest used by Cart
    type AddCartRequest struct {

        
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

        
            ItemSize string  `json:"item_size"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            Quantity float64  `json:"quantity"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
            ArticleID string  `json:"article_id"`
            ItemID float64  `json:"item_id"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
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

        
            Message string  `json:"message"`
            Cart CartDetailResponse  `json:"cart"`
            Success bool  `json:"success"`
         
    }
    
    // OverrideCartItemPromo used by Cart
    type OverrideCartItemPromo struct {

        
            PromoID string  `json:"promo_id"`
            PromoAmount string  `json:"promo_amount"`
            PromoDesc string  `json:"promo_desc"`
            ItemList []map[string]interface{}  `json:"item_list"`
            RwrdTndr string  `json:"rwrd_tndr"`
         
    }
    
    // OverrideCartItem used by Cart
    type OverrideCartItem struct {

        
            Size string  `json:"size"`
            PriceMarked float64  `json:"price_marked"`
            PriceEffective float64  `json:"price_effective"`
            Quantity float64  `json:"quantity"`
            SellerIdentifier string  `json:"seller_identifier"`
            PromoList []OverrideCartItemPromo  `json:"promo_list"`
            AmountPaid float64  `json:"amount_paid"`
            Discount float64  `json:"discount"`
            ItemID float64  `json:"item_id"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
         
    }
    
    // OverrideCheckoutReq used by Cart
    type OverrideCheckoutReq struct {

        
            CurrencyCode string  `json:"currency_code"`
            PaymentIdentifier string  `json:"payment_identifier"`
            PaymentMode string  `json:"payment_mode"`
            MerchantCode string  `json:"merchant_code"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
            CartItems []OverrideCartItem  `json:"cart_items"`
            OrderType string  `json:"order_type"`
            ShippingAddress map[string]interface{}  `json:"shipping_address"`
            CartID string  `json:"cart_id"`
            OrderingStore float64  `json:"ordering_store"`
            Aggregator string  `json:"aggregator"`
         
    }
    
    // OverrideCheckoutResponse used by Cart
    type OverrideCheckoutResponse struct {

        
            Data map[string]interface{}  `json:"data"`
            Cart map[string]interface{}  `json:"cart"`
            Message string  `json:"message"`
            Success string  `json:"success"`
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
    

    
    // GenerateReportMeta used by Finance
    type GenerateReportMeta struct {

        
            Company string  `json:"company"`
            Brand string  `json:"brand"`
            Channel string  `json:"channel"`
         
    }
    
    // GenerateReportFilters used by Finance
    type GenerateReportFilters struct {

        
            Company []string  `json:"company"`
            Brand []string  `json:"brand"`
            Channel []string  `json:"channel"`
         
    }
    
    // GenerateReportPlatform used by Finance
    type GenerateReportPlatform struct {

        
            Meta GenerateReportMeta  `json:"meta"`
            EndDate string  `json:"end_date"`
            ReportID string  `json:"report_id"`
            Filters GenerateReportFilters  `json:"filters"`
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
            Page Page  `json:"page"`
            EndDate string  `json:"end_date"`
            StartDate string  `json:"start_date"`
            Items [][]string  `json:"items"`
            Headers []string  `json:"headers"`
         
    }
    
    // Error used by Finance
    type Error struct {

        
            Success bool  `json:"success"`
            Reason string  `json:"reason"`
         
    }
    
    // DownloadReport used by Finance
    type DownloadReport struct {

        
            Pagesize float64  `json:"pagesize"`
            EndDate string  `json:"end_date"`
            StartDate string  `json:"start_date"`
            Page float64  `json:"page"`
         
    }
    
    // DownloadReportItems used by Finance
    type DownloadReportItems struct {

        
            Meta GenerateReportMeta  `json:"meta"`
            TypeOfRequest string  `json:"type_of_request"`
            EndDate string  `json:"end_date"`
            ReportID string  `json:"report_id"`
            Filters GenerateReportFilters  `json:"filters"`
            StartDate string  `json:"start_date"`
         
    }
    
    // DownloadReportList used by Finance
    type DownloadReportList struct {

        
            Items []DownloadReportItems  `json:"items"`
            Page Page  `json:"page"`
            ItemCount float64  `json:"item_count"`
         
    }
    
    // GetEngineFilters used by Finance
    type GetEngineFilters struct {

        
            ConfigField string  `json:"config_field"`
         
    }
    
    // GetEngineData used by Finance
    type GetEngineData struct {

        
            Filters GetEngineFilters  `json:"filters"`
            TableName string  `json:"table_name"`
            Project []string  `json:"project"`
         
    }
    
    // GetEngineRequest used by Finance
    type GetEngineRequest struct {

        
            Data GetEngineData  `json:"data"`
         
    }
    
    // GetEngineResponse used by Finance
    type GetEngineResponse struct {

        
            Success bool  `json:"success"`
            Items []map[string]interface{}  `json:"items"`
            Page Page  `json:"page"`
            ItemCount float64  `json:"item_count"`
         
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

        
            Docs []map[string]interface{}  `json:"docs"`
            Items []map[string]interface{}  `json:"items"`
         
    }
    
    // GetReasonResponse used by Finance
    type GetReasonResponse struct {

        
            Success bool  `json:"success"`
            Data GetDocs  `json:"data"`
         
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

        
            PdfS3URL string  `json:"pdf_s3_url"`
            ID string  `json:"id"`
         
    }
    
    // DownloadCreditDebitNoteResponse used by Finance
    type DownloadCreditDebitNoteResponse struct {

        
            Success bool  `json:"success"`
            Data []DownloadCreditDebitNoteResponseData  `json:"data"`
         
    }
    
    // PaymentProcessPayload used by Finance
    type PaymentProcessPayload struct {

        
            Meta map[string]interface{}  `json:"meta"`
            SourceReference string  `json:"source_reference"`
            Platform string  `json:"platform"`
            Currency string  `json:"currency"`
            TransactionType string  `json:"transaction_type"`
            SellerID string  `json:"seller_id"`
            ModeOfPayment string  `json:"mode_of_payment"`
            InvoiceNumber string  `json:"invoice_number"`
            Amount string  `json:"amount"`
            TotalAmount string  `json:"total_amount"`
         
    }
    
    // PaymentProcessRequest used by Finance
    type PaymentProcessRequest struct {

        
            Data PaymentProcessPayload  `json:"data"`
         
    }
    
    // PaymentProcessResponse used by Finance
    type PaymentProcessResponse struct {

        
            Meta map[string]interface{}  `json:"meta"`
            Message string  `json:"message"`
            RedirectURL string  `json:"redirect_url"`
            TransactionID string  `json:"transaction_id"`
            Code float64  `json:"code"`
         
    }
    
    // CreditlineDataPlatformPayload used by Finance
    type CreditlineDataPlatformPayload struct {

        
            EndEnd string  `json:"end_end"`
            Page float64  `json:"page"`
            SellerID string  `json:"seller_id"`
            Pagesize float64  `json:"pagesize"`
            StartEnd string  `json:"start_end"`
         
    }
    
    // CreditlineDataPlatformRequest used by Finance
    type CreditlineDataPlatformRequest struct {

        
            Data CreditlineDataPlatformPayload  `json:"data"`
         
    }
    
    // CreditlineDataPlatformResponse used by Finance
    type CreditlineDataPlatformResponse struct {

        
            ItemCount float64  `json:"item_count"`
            Message string  `json:"message"`
            Page map[string]interface{}  `json:"page"`
            Code float64  `json:"code"`
            ShowMr bool  `json:"show_mr"`
            Items []map[string]interface{}  `json:"items"`
            Headers []string  `json:"headers"`
         
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

        
            Code float64  `json:"code"`
            IsCreditlineOpted bool  `json:"is_creditline_opted"`
         
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

        
            Text string  `json:"text"`
            Value string  `json:"value"`
         
    }
    
    // InvoiceTypeResponse used by Finance
    type InvoiceTypeResponse struct {

        
            InvoiceTypeList []InvoiceTypeResponseItems  `json:"invoice_type_list"`
            Success bool  `json:"success"`
            PaymentStatusList []InvoiceTypeResponseItems  `json:"payment_status_list"`
         
    }
    
    // InoviceListingPayloadDataFilters used by Finance
    type InoviceListingPayloadDataFilters struct {

        
            PaymentStatus []string  `json:"payment_status"`
            CompanyID []string  `json:"company_id"`
            InvoiceType []string  `json:"invoice_type"`
         
    }
    
    // InvoiceListingPayloadData used by Finance
    type InvoiceListingPayloadData struct {

        
            PageSize float64  `json:"page_size"`
            Page float64  `json:"page"`
            Search string  `json:"search"`
            EndDate string  `json:"end_date"`
            Filters InoviceListingPayloadDataFilters  `json:"filters"`
            StartDate string  `json:"start_date"`
         
    }
    
    // InvoiceListingRequest used by Finance
    type InvoiceListingRequest struct {

        
            Data InvoiceListingPayloadData  `json:"data"`
         
    }
    
    // UnpaidInvoiceDataItems used by Finance
    type UnpaidInvoiceDataItems struct {

        
            TotalUnpaidInvoiceCount float64  `json:"total_unpaid_invoice_count"`
            Currency string  `json:"currency"`
            TotalUnpaidAmount float64  `json:"total_unpaid_amount"`
         
    }
    
    // InvoiceListingResponseItems used by Finance
    type InvoiceListingResponseItems struct {

        
            Company string  `json:"company"`
            Amount string  `json:"amount"`
            Period string  `json:"period"`
            IsDownloadable bool  `json:"is_downloadable"`
            DueDate string  `json:"due_date"`
            InvoiceID string  `json:"invoice_id"`
            InvoiceType string  `json:"invoice_type"`
            InvoiceNumber string  `json:"invoice_number"`
            Status string  `json:"status"`
            InvoiceDate string  `json:"invoice_date"`
         
    }
    
    // InvoiceListingResponse used by Finance
    type InvoiceListingResponse struct {

        
            UnpaidInvoiceData UnpaidInvoiceDataItems  `json:"unpaid_invoice_data"`
            Items []InvoiceListingResponseItems  `json:"items"`
            Page Page  `json:"page"`
            ItemCount float64  `json:"item_count"`
         
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

        
            Success bool  `json:"success"`
            Data []string  `json:"data"`
            Error []string  `json:"error"`
         
    }
    
