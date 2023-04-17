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
         
    }
    
    // SubscriptionStatus used by Billing
    type SubscriptionStatus struct {

        
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
            Aggregators []map[string]interface{}  `json:"aggregators"`
            Success bool  `json:"success"`
            Created bool  `json:"created"`
            AppID string  `json:"app_id"`
            ExcludedFields []string  `json:"excluded_fields"`
         
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
            MerchantSalt string  `json:"merchant_salt"`
            IsActive bool  `json:"is_active"`
            Key string  `json:"key"`
            ConfigType string  `json:"config_type"`
         
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

        
            Large string  `json:"large"`
            Small string  `json:"small"`
         
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

        
            RemainingLimit float64  `json:"remaining_limit"`
            CardBrandImage string  `json:"card_brand_image"`
            AggregatorName string  `json:"aggregator_name"`
            FyndVpa string  `json:"fynd_vpa"`
            CodLimitPerOrder float64  `json:"cod_limit_per_order"`
            CardNumber string  `json:"card_number"`
            IntentAppErrorDictList []IntentAppErrorList  `json:"intent_app_error_dict_list"`
            CardReference string  `json:"card_reference"`
            LogoURL PaymentModeLogo  `json:"logo_url"`
            CardName string  `json:"card_name"`
            Nickname string  `json:"nickname"`
            Timeout float64  `json:"timeout"`
            CardFingerprint string  `json:"card_fingerprint"`
            ExpYear float64  `json:"exp_year"`
            CardID string  `json:"card_id"`
            DisplayName string  `json:"display_name"`
            CardIsin string  `json:"card_isin"`
            Name string  `json:"name"`
            Code string  `json:"code"`
            MerchantCode string  `json:"merchant_code"`
            IntentFlow bool  `json:"intent_flow"`
            RetryCount float64  `json:"retry_count"`
            DisplayPriority float64  `json:"display_priority"`
            ExpMonth float64  `json:"exp_month"`
            CompliantWithTokenisationGuidelines bool  `json:"compliant_with_tokenisation_guidelines"`
            CardIssuer string  `json:"card_issuer"`
            CardBrand string  `json:"card_brand"`
            IntentAppErrorList []string  `json:"intent_app_error_list"`
            CardToken string  `json:"card_token"`
            CardType string  `json:"card_type"`
            Expired bool  `json:"expired"`
            CodLimit float64  `json:"cod_limit"`
            IntentApp []IntentApp  `json:"intent_app"`
         
    }
    
    // RootPaymentMode used by Payment
    type RootPaymentMode struct {

        
            List []PaymentModeList  `json:"list"`
            DisplayName string  `json:"display_name"`
            IsPayByCardPl bool  `json:"is_pay_by_card_pl"`
            SaveCard bool  `json:"save_card"`
            AggregatorName string  `json:"aggregator_name"`
            Name string  `json:"name"`
            AddCardEnabled bool  `json:"add_card_enabled"`
            AnonymousEnable bool  `json:"anonymous_enable"`
            DisplayPriority float64  `json:"display_priority"`
         
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

        
            UniqueTransferNo map[string]interface{}  `json:"unique_transfer_no"`
            Customers map[string]interface{}  `json:"customers"`
            IsActive bool  `json:"is_active"`
            TransferType string  `json:"transfer_type"`
            IsDefault bool  `json:"is_default"`
            PayoutsAggregators []map[string]interface{}  `json:"payouts_aggregators"`
            MoreAttributes map[string]interface{}  `json:"more_attributes"`
         
    }
    
    // PayoutBankDetails used by Payment
    type PayoutBankDetails struct {

        
            Country string  `json:"country"`
            BranchName string  `json:"branch_name"`
            Pincode float64  `json:"pincode"`
            AccountHolder string  `json:"account_holder"`
            AccountType string  `json:"account_type"`
            IfscCode string  `json:"ifsc_code"`
            City string  `json:"city"`
            AccountNo string  `json:"account_no"`
            State string  `json:"state"`
            BankName string  `json:"bank_name"`
         
    }
    
    // PayoutRequest used by Payment
    type PayoutRequest struct {

        
            Aggregator string  `json:"aggregator"`
            UniqueExternalID string  `json:"unique_external_id"`
            IsActive bool  `json:"is_active"`
            BankDetails PayoutBankDetails  `json:"bank_details"`
            TransferType string  `json:"transfer_type"`
            Users map[string]interface{}  `json:"users"`
         
    }
    
    // PayoutResponse used by Payment
    type PayoutResponse struct {

        
            UniqueTransferNo string  `json:"unique_transfer_no"`
            Aggregator string  `json:"aggregator"`
            Payouts map[string]interface{}  `json:"payouts"`
            IsActive bool  `json:"is_active"`
            BankDetails map[string]interface{}  `json:"bank_details"`
            TransferType string  `json:"transfer_type"`
            Success bool  `json:"success"`
            Created bool  `json:"created"`
            Users map[string]interface{}  `json:"users"`
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
            Aggregator string  `json:"aggregator"`
            Config map[string]interface{}  `json:"config"`
         
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
            IsVerifiedFlag bool  `json:"is_verified_flag"`
            Message string  `json:"message"`
            Data map[string]interface{}  `json:"data"`
         
    }
    
    // NotFoundResourceError used by Payment
    type NotFoundResourceError struct {

        
            Code string  `json:"code"`
            Success bool  `json:"success"`
            Description string  `json:"description"`
         
    }
    
    // BankDetailsForOTP used by Payment
    type BankDetailsForOTP struct {

        
            BranchName string  `json:"branch_name"`
            AccountHolder string  `json:"account_holder"`
            IfscCode string  `json:"ifsc_code"`
            AccountNo string  `json:"account_no"`
            BankName string  `json:"bank_name"`
         
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

        
            Email string  `json:"email"`
            BranchName string  `json:"branch_name"`
            Address string  `json:"address"`
            IsActive bool  `json:"is_active"`
            AccountHolder string  `json:"account_holder"`
            IfscCode string  `json:"ifsc_code"`
            AccountNo string  `json:"account_no"`
            Comment string  `json:"comment"`
            BankName string  `json:"bank_name"`
            Title string  `json:"title"`
            Mobile string  `json:"mobile"`
            BeneficiaryID string  `json:"beneficiary_id"`
            CreatedOn string  `json:"created_on"`
            TransferMode string  `json:"transfer_mode"`
            DisplayName string  `json:"display_name"`
            Subtitle string  `json:"subtitle"`
            ModifiedOn string  `json:"modified_on"`
            DelightsUserName string  `json:"delights_user_name"`
            ID float64  `json:"id"`
         
    }
    
    // OrderBeneficiaryResponse used by Payment
    type OrderBeneficiaryResponse struct {

        
            Beneficiaries []OrderBeneficiaryDetails  `json:"beneficiaries"`
            ShowBeneficiaryDetails bool  `json:"show_beneficiary_details"`
         
    }
    
    // MultiTenderPaymentMeta used by Payment
    type MultiTenderPaymentMeta struct {

        
            OrderID string  `json:"order_id"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            CurrentStatus string  `json:"current_status"`
            PaymentID string  `json:"payment_id"`
            PaymentGateway string  `json:"payment_gateway"`
         
    }
    
    // MultiTenderPaymentMethod used by Payment
    type MultiTenderPaymentMethod struct {

        
            Amount float64  `json:"amount"`
            Meta MultiTenderPaymentMeta  `json:"meta"`
            Mode string  `json:"mode"`
            Name string  `json:"name"`
         
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

        
            Usages float64  `json:"usages"`
            RemainingLimit float64  `json:"remaining_limit"`
            UserID string  `json:"user_id"`
            IsActive bool  `json:"is_active"`
            Limit float64  `json:"limit"`
         
    }
    
    // GetUserCODLimitResponse used by Payment
    type GetUserCODLimitResponse struct {

        
            Success bool  `json:"success"`
            UserCodData CODdata  `json:"user_cod_data"`
         
    }
    
    // SetCODForUserRequest used by Payment
    type SetCODForUserRequest struct {

        
            MerchantUserID string  `json:"merchant_user_id"`
            IsActive bool  `json:"is_active"`
            Mobileno string  `json:"mobileno"`
         
    }
    
    // SetCODOptionResponse used by Payment
    type SetCODOptionResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    

    
    // PlatformItem used by Order
    type PlatformItem struct {

        
            Code string  `json:"code"`
            L3Category float64  `json:"l3_category"`
            DepartmentID float64  `json:"department_id"`
            L1Category []string  `json:"l1_category"`
            Images []string  `json:"images"`
            ID float64  `json:"id"`
            Size string  `json:"size"`
            CanCancel bool  `json:"can_cancel"`
            Color string  `json:"color"`
            CanReturn bool  `json:"can_return"`
            Name string  `json:"name"`
            L3CategoryName string  `json:"l3_category_name"`
            Image []string  `json:"image"`
         
    }
    
    // GSTDetailsData used by Order
    type GSTDetailsData struct {

        
            ValueOfGood float64  `json:"value_of_good"`
            GstFee float64  `json:"gst_fee"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            GstinCode string  `json:"gstin_code"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
         
    }
    
    // Prices used by Order
    type Prices struct {

        
            TransferPrice float64  `json:"transfer_price"`
            CouponValue float64  `json:"coupon_value"`
            ValueOfGood float64  `json:"value_of_good"`
            CashbackApplied float64  `json:"cashback_applied"`
            RefundAmount float64  `json:"refund_amount"`
            DeliveryCharge float64  `json:"delivery_charge"`
            PriceEffective float64  `json:"price_effective"`
            Discount float64  `json:"discount"`
            CodCharges float64  `json:"cod_charges"`
            Cashback float64  `json:"cashback"`
            FyndCredits float64  `json:"fynd_credits"`
            RefundCredit float64  `json:"refund_credit"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            AmountPaid float64  `json:"amount_paid"`
            PriceMarked float64  `json:"price_marked"`
            PromotionEffectiveDiscount float64  `json:"promotion_effective_discount"`
            AmountPaidRoundoff float64  `json:"amount_paid_roundoff"`
         
    }
    
    // BagUnit used by Order
    type BagUnit struct {

        
            Item PlatformItem  `json:"item"`
            OrderingChannel string  `json:"ordering_channel"`
            ShipmentID string  `json:"shipment_id"`
            Status map[string]interface{}  `json:"status"`
            BagID float64  `json:"bag_id"`
            ItemQuantity float64  `json:"item_quantity"`
            Gst GSTDetailsData  `json:"gst"`
            CanCancel bool  `json:"can_cancel"`
            CanReturn bool  `json:"can_return"`
            Prices Prices  `json:"prices"`
            TotalShipmentBags float64  `json:"total_shipment_bags"`
         
    }
    
    // PaymentModeInfo used by Order
    type PaymentModeInfo struct {

        
            Logo string  `json:"logo"`
            Type string  `json:"type"`
         
    }
    
    // UserDataInfo used by Order
    type UserDataInfo struct {

        
            FirstName string  `json:"first_name"`
            Mobile string  `json:"mobile"`
            AvisUserID string  `json:"avis_user_id"`
            Gender string  `json:"gender"`
            LastName string  `json:"last_name"`
            Email string  `json:"email"`
            IsAnonymousUser bool  `json:"is_anonymous_user"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // ShipmentStatus used by Order
    type ShipmentStatus struct {

        
            Status string  `json:"status"`
            HexCode string  `json:"hex_code"`
            Title string  `json:"title"`
            OpsStatus string  `json:"ops_status"`
            ActualStatus string  `json:"actual_status"`
         
    }
    
    // ShipmentItemFulFillingStore used by Order
    type ShipmentItemFulFillingStore struct {

        
            Code string  `json:"code"`
            ID string  `json:"id"`
         
    }
    
    // ShipmentItem used by Order
    type ShipmentItem struct {

        
            Sla map[string]interface{}  `json:"sla"`
            Bags []BagUnit  `json:"bags"`
            FulfillingCentre string  `json:"fulfilling_centre"`
            PaymentModeInfo PaymentModeInfo  `json:"payment_mode_info"`
            ID string  `json:"id"`
            ShipmentID string  `json:"shipment_id"`
            ShipmentCreatedAt string  `json:"shipment_created_at"`
            Channel map[string]interface{}  `json:"channel"`
            User UserDataInfo  `json:"user"`
            ShipmentStatus ShipmentStatus  `json:"shipment_status"`
            FulfillingStore ShipmentItemFulFillingStore  `json:"fulfilling_store"`
            TotalBagsCount float64  `json:"total_bags_count"`
            PaymentMethods map[string]interface{}  `json:"payment_methods"`
            CreatedAt string  `json:"created_at"`
            Prices Prices  `json:"prices"`
            Application map[string]interface{}  `json:"application"`
            TotalShipmentsInOrder float64  `json:"total_shipments_in_order"`
         
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
            Type string  `json:"type"`
            Value string  `json:"value"`
         
    }
    
    // ShipmentInternalPlatformViewResponse used by Order
    type ShipmentInternalPlatformViewResponse struct {

        
            Items []ShipmentItem  `json:"items"`
            Filters []FiltersInfo  `json:"filters"`
            AppliedFilters map[string]interface{}  `json:"applied_filters"`
            Page map[string]interface{}  `json:"page"`
         
    }
    
    // Error used by Order
    type Error struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // ShipmentPayments used by Order
    type ShipmentPayments struct {

        
            Mode string  `json:"mode"`
            Logo string  `json:"logo"`
            Source string  `json:"source"`
         
    }
    
    // BagStateMapper used by Order
    type BagStateMapper struct {

        
            BsID float64  `json:"bs_id"`
            StateType string  `json:"state_type"`
            AppDisplayName string  `json:"app_display_name"`
            DisplayName string  `json:"display_name"`
            NotifyCustomer bool  `json:"notify_customer"`
            AppStateName string  `json:"app_state_name"`
            AppFacing bool  `json:"app_facing"`
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            JourneyType string  `json:"journey_type"`
         
    }
    
    // BagStatusHistory used by Order
    type BagStatusHistory struct {

        
            Forward bool  `json:"forward"`
            StateType string  `json:"state_type"`
            AppDisplayName string  `json:"app_display_name"`
            ShipmentID string  `json:"shipment_id"`
            Status string  `json:"status"`
            Reasons []map[string]interface{}  `json:"reasons"`
            BagID float64  `json:"bag_id"`
            DeliveryPartnerID float64  `json:"delivery_partner_id"`
            CreatedAt string  `json:"created_at"`
            DeliveryAwbNumber string  `json:"delivery_awb_number"`
            BshID float64  `json:"bsh_id"`
            DisplayName string  `json:"display_name"`
            StoreID float64  `json:"store_id"`
            StateID float64  `json:"state_id"`
            BagStateMapper BagStateMapper  `json:"bag_state_mapper"`
            UpdatedAt string  `json:"updated_at"`
            KafkaSync bool  `json:"kafka_sync"`
         
    }
    
    // FulfillingStore used by Order
    type FulfillingStore struct {

        
            Code string  `json:"code"`
            Address string  `json:"address"`
            ID float64  `json:"id"`
            Pincode string  `json:"pincode"`
            ContactPerson string  `json:"contact_person"`
            FulfillmentChannel string  `json:"fulfillment_channel"`
            State string  `json:"state"`
            Country string  `json:"country"`
            StoreName string  `json:"store_name"`
            City string  `json:"city"`
            Meta map[string]interface{}  `json:"meta"`
            Phone string  `json:"phone"`
         
    }
    
    // PhoneDetails used by Order
    type PhoneDetails struct {

        
            CountryCode float64  `json:"country_code"`
            MobileNumber string  `json:"mobile_number"`
         
    }
    
    // ContactDetails used by Order
    type ContactDetails struct {

        
            Emails []string  `json:"emails"`
            Phone []PhoneDetails  `json:"phone"`
         
    }
    
    // CompanyDetails used by Order
    type CompanyDetails struct {

        
            CompanyID float64  `json:"company_id"`
            Address map[string]interface{}  `json:"address"`
            CompanyGst string  `json:"company_gst"`
            CompanyName string  `json:"company_name"`
            CompanyContact ContactDetails  `json:"company_contact"`
            CompanyCin string  `json:"company_cin"`
         
    }
    
    // TrackingList used by Order
    type TrackingList struct {

        
            Text string  `json:"text"`
            IsPassed bool  `json:"is_passed"`
            Status string  `json:"status"`
            Time string  `json:"time"`
            IsCurrent bool  `json:"is_current"`
         
    }
    
    // OrderDetailsData used by Order
    type OrderDetailsData struct {

        
            AffiliateID string  `json:"affiliate_id"`
            OrderingChannel string  `json:"ordering_channel"`
            Source string  `json:"source"`
            OrderDate string  `json:"order_date"`
            FyndOrderID string  `json:"fynd_order_id"`
            OrderingChannelLogo map[string]interface{}  `json:"ordering_channel_logo"`
            CodCharges string  `json:"cod_charges"`
            OrderValue string  `json:"order_value"`
            TaxDetails map[string]interface{}  `json:"tax_details"`
         
    }
    
    // DPDetailsData used by Order
    type DPDetailsData struct {

        
            EwayBillID string  `json:"eway_bill_id"`
            TrackURL string  `json:"track_url"`
            ID float64  `json:"id"`
            Pincode string  `json:"pincode"`
            GstTag string  `json:"gst_tag"`
            Country string  `json:"country"`
            AwbNo string  `json:"awb_no"`
            Name string  `json:"name"`
         
    }
    
    // PlatformDeliveryAddress used by Order
    type PlatformDeliveryAddress struct {

        
            Address1 string  `json:"address1"`
            AddressType string  `json:"address_type"`
            AddressCategory string  `json:"address_category"`
            Longitude float64  `json:"longitude"`
            Pincode string  `json:"pincode"`
            ContactPerson string  `json:"contact_person"`
            Area string  `json:"area"`
            State string  `json:"state"`
            Latitude float64  `json:"latitude"`
            Country string  `json:"country"`
            UpdatedAt string  `json:"updated_at"`
            Email string  `json:"email"`
            Landmark string  `json:"landmark"`
            CreatedAt string  `json:"created_at"`
            Version string  `json:"version"`
            City string  `json:"city"`
            Address2 string  `json:"address2"`
            Phone string  `json:"phone"`
         
    }
    
    // CurrentStatus used by Order
    type CurrentStatus struct {

        
            StateID float64  `json:"state_id"`
            StateType string  `json:"state_type"`
            ShipmentID string  `json:"shipment_id"`
            Status string  `json:"status"`
            BagID float64  `json:"bag_id"`
            DeliveryPartnerID float64  `json:"delivery_partner_id"`
            DeliveryAwbNumber string  `json:"delivery_awb_number"`
            CurrentStatusID float64  `json:"current_status_id"`
            StoreID float64  `json:"store_id"`
            CreatedAt string  `json:"created_at"`
            BagStateMapper BagStateMapper  `json:"bag_state_mapper"`
            UpdatedAt string  `json:"updated_at"`
            KafkaSync bool  `json:"kafka_sync"`
         
    }
    
    // BagGST used by Order
    type BagGST struct {

        
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            IsDefaultHsnCode bool  `json:"is_default_hsn_code"`
            ValueOfGood float64  `json:"value_of_good"`
            GstTag string  `json:"gst_tag"`
            GstFee float64  `json:"gst_fee"`
            GstinCode string  `json:"gstin_code"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            HsnCode string  `json:"hsn_code"`
         
    }
    
    // Identifier used by Order
    type Identifier struct {

        
            Ean string  `json:"ean"`
         
    }
    
    // FinancialBreakup used by Order
    type FinancialBreakup struct {

        
            Identifiers Identifier  `json:"identifiers"`
            TransferPrice float64  `json:"transfer_price"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            CashbackApplied float64  `json:"cashback_applied"`
            GstTag string  `json:"gst_tag"`
            GstFee float64  `json:"gst_fee"`
            RefundCredit float64  `json:"refund_credit"`
            Cashback float64  `json:"cashback"`
            AmountPaid float64  `json:"amount_paid"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            TotalUnits float64  `json:"total_units"`
            Size string  `json:"size"`
            DeliveryCharge float64  `json:"delivery_charge"`
            FyndCredits float64  `json:"fynd_credits"`
            HsnCode string  `json:"hsn_code"`
            AddedToFyndCash bool  `json:"added_to_fynd_cash"`
            CouponValue float64  `json:"coupon_value"`
            PriceEffective float64  `json:"price_effective"`
            Discount float64  `json:"discount"`
            CodCharges float64  `json:"cod_charges"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            PromotionEffectiveDiscount float64  `json:"promotion_effective_discount"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            PriceMarked float64  `json:"price_marked"`
            AmountPaidRoundoff float64  `json:"amount_paid_roundoff"`
            ValueOfGood float64  `json:"value_of_good"`
            ItemName string  `json:"item_name"`
         
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

        
            Amount float64  `json:"amount"`
            PromotionName string  `json:"promotion_name"`
            PromotionType string  `json:"promotion_type"`
            MrpPromotion bool  `json:"mrp_promotion"`
            DiscountRules []DiscountRules  `json:"discount_rules"`
            BuyRules []BuyRules  `json:"buy_rules"`
            ArticleQuantity float64  `json:"article_quantity"`
            PromoID string  `json:"promo_id"`
         
    }
    
    // OrderBrandName used by Order
    type OrderBrandName struct {

        
            BrandName string  `json:"brand_name"`
            ID float64  `json:"id"`
            Logo string  `json:"logo"`
            CreatedOn string  `json:"created_on"`
            Company string  `json:"company"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // OrderBagArticle used by Order
    type OrderBagArticle struct {

        
            Identifiers map[string]interface{}  `json:"identifiers"`
            ReturnConfig map[string]interface{}  `json:"return_config"`
            UID string  `json:"uid"`
         
    }
    
    // BagConfigs used by Order
    type BagConfigs struct {

        
            IsCustomerReturnAllowed bool  `json:"is_customer_return_allowed"`
            CanBeCancelled bool  `json:"can_be_cancelled"`
            AllowForceReturn bool  `json:"allow_force_return"`
            IsReturnable bool  `json:"is_returnable"`
            IsActive bool  `json:"is_active"`
            EnableTracking bool  `json:"enable_tracking"`
         
    }
    
    // OrderBags used by Order
    type OrderBags struct {

        
            Item PlatformItem  `json:"item"`
            Quantity float64  `json:"quantity"`
            LineNumber float64  `json:"line_number"`
            DeliveryAddress PlatformDeliveryAddress  `json:"delivery_address"`
            CurrentStatus CurrentStatus  `json:"current_status"`
            DisplayName string  `json:"display_name"`
            BagID float64  `json:"bag_id"`
            CanReturn bool  `json:"can_return"`
            GstDetails BagGST  `json:"gst_details"`
            FinancialBreakup FinancialBreakup  `json:"financial_breakup"`
            Prices Prices  `json:"prices"`
            EntityType string  `json:"entity_type"`
            AppliedPromos []AppliedPromos  `json:"applied_promos"`
            Brand OrderBrandName  `json:"brand"`
            Identifier string  `json:"identifier"`
            Article OrderBagArticle  `json:"article"`
            SellerIdentifier string  `json:"seller_identifier"`
            ParentPromoBags map[string]interface{}  `json:"parent_promo_bags"`
            CanCancel bool  `json:"can_cancel"`
            BagConfigs BagConfigs  `json:"bag_configs"`
         
    }
    
    // PDFLinks used by Order
    type PDFLinks struct {

        
            CreditNoteURL string  `json:"credit_note_url"`
            InvoiceA4 string  `json:"invoice_a4"`
            DeliveryChallanA4 string  `json:"delivery_challan_a4"`
            LabelA6 string  `json:"label_a6"`
            Invoice string  `json:"invoice"`
            B2b string  `json:"b2b"`
            Label string  `json:"label"`
            InvoiceA6 string  `json:"invoice_a6"`
            LabelA4 string  `json:"label_a4"`
            InvoiceType string  `json:"invoice_type"`
            InvoicePos string  `json:"invoice_pos"`
            LabelType string  `json:"label_type"`
            LabelPos string  `json:"label_pos"`
            PoInvoice string  `json:"po_invoice"`
         
    }
    
    // AffiliateMeta used by Order
    type AffiliateMeta struct {

        
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            BoxType string  `json:"box_type"`
            SizeLevelTotalQty float64  `json:"size_level_total_qty"`
            EmployeeDiscount float64  `json:"employee_discount"`
            Quantity float64  `json:"quantity"`
            CouponCode string  `json:"coupon_code"`
            ChannelOrderID string  `json:"channel_order_id"`
            DueDate string  `json:"due_date"`
            ChannelShipmentID string  `json:"channel_shipment_id"`
            IsPriority bool  `json:"is_priority"`
            OrderItemID string  `json:"order_item_id"`
         
    }
    
    // Formatted used by Order
    type Formatted struct {

        
            FMax string  `json:"f_max"`
            FMin string  `json:"f_min"`
         
    }
    
    // DebugInfo used by Order
    type DebugInfo struct {

        
            StormbreakerUUID string  `json:"stormbreaker_uuid"`
         
    }
    
    // ShipmentTimeStamp used by Order
    type ShipmentTimeStamp struct {

        
            TMax string  `json:"t_max"`
            TMin string  `json:"t_min"`
         
    }
    
    // BuyerDetails used by Order
    type BuyerDetails struct {

        
            Address string  `json:"address"`
            AjioSiteID string  `json:"ajio_site_id"`
            Pincode float64  `json:"pincode"`
            State string  `json:"state"`
            Gstin string  `json:"gstin"`
            City string  `json:"city"`
            Name string  `json:"name"`
         
    }
    
    // EinvoiceInfo used by Order
    type EinvoiceInfo struct {

        
            CreditNote map[string]interface{}  `json:"credit_note"`
            Invoice map[string]interface{}  `json:"invoice"`
         
    }
    
    // LockData used by Order
    type LockData struct {

        
            LockMessage string  `json:"lock_message"`
            Locked bool  `json:"locked"`
            Mto bool  `json:"mto"`
         
    }
    
    // ShipmentMeta used by Order
    type ShipmentMeta struct {

        
            ReturnAwbNumber string  `json:"return_awb_number"`
            OrderType string  `json:"order_type"`
            Formatted Formatted  `json:"formatted"`
            AutoTriggerDpAssignmentAcf bool  `json:"auto_trigger_dp_assignment_acf"`
            MarketplaceStoreID string  `json:"marketplace_store_id"`
            PoNumber string  `json:"po_number"`
            BagWeight map[string]interface{}  `json:"bag_weight"`
            DebugInfo DebugInfo  `json:"debug_info"`
            DpName string  `json:"dp_name"`
            Timestamp ShipmentTimeStamp  `json:"timestamp"`
            DpOptions map[string]interface{}  `json:"dp_options"`
            ShipmentWeight float64  `json:"shipment_weight"`
            PackagingName string  `json:"packaging_name"`
            ReturnAffiliateOrderID string  `json:"return_affiliate_order_id"`
            ReturnStoreNode float64  `json:"return_store_node"`
            DpID string  `json:"dp_id"`
            DueDate string  `json:"due_date"`
            AssignDpFromSb bool  `json:"assign_dp_from_sb"`
            SameStoreAvailable bool  `json:"same_store_available"`
            External map[string]interface{}  `json:"external"`
            B2cBuyerDetails map[string]interface{}  `json:"b2c_buyer_details"`
            FulfilmentPriorityText string  `json:"fulfilment_priority_text"`
            ReturnAffiliateShipmentID string  `json:"return_affiliate_shipment_id"`
            B2bBuyerDetails BuyerDetails  `json:"b2b_buyer_details"`
            BoxType string  `json:"box_type"`
            StoreInvoiceUpdatedDate string  `json:"store_invoice_updated_date"`
            AwbNumber string  `json:"awb_number"`
            ForwardAffiliateOrderID string  `json:"forward_affiliate_order_id"`
            ReturnDetails map[string]interface{}  `json:"return_details"`
            ShipmentVolumetricWeight float64  `json:"shipment_volumetric_weight"`
            ForwardAffiliateShipmentID string  `json:"forward_affiliate_shipment_id"`
            DpSortKey string  `json:"dp_sort_key"`
            EinvoiceInfo EinvoiceInfo  `json:"einvoice_info"`
            EwaybillInfo map[string]interface{}  `json:"ewaybill_info"`
            Weight float64  `json:"weight"`
            LockData LockData  `json:"lock_data"`
         
    }
    
    // AffiliateDetails used by Order
    type AffiliateDetails struct {

        
            CompanyAffiliateTag string  `json:"company_affiliate_tag"`
            AffiliateID string  `json:"affiliate_id"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            PdfLinks PDFLinks  `json:"pdf_links"`
            AdID string  `json:"ad_id"`
            AffiliateStoreID string  `json:"affiliate_store_id"`
            AffiliateMeta AffiliateMeta  `json:"affiliate_meta"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
            ShipmentMeta ShipmentMeta  `json:"shipment_meta"`
         
    }
    
    // InvoiceInfo used by Order
    type InvoiceInfo struct {

        
            InvoiceURL string  `json:"invoice_url"`
            LabelURL string  `json:"label_url"`
            StoreInvoiceID string  `json:"store_invoice_id"`
            UpdatedDate string  `json:"updated_date"`
            CreditNoteID string  `json:"credit_note_id"`
         
    }
    
    // Dimensions used by Order
    type Dimensions struct {

        
            Width float64  `json:"width"`
            Unit string  `json:"unit"`
            Height float64  `json:"height"`
            IsDefault bool  `json:"is_default"`
            Length float64  `json:"length"`
         
    }
    
    // Meta used by Order
    type Meta struct {

        
            Dimension Dimensions  `json:"dimension"`
         
    }
    
    // UserDetailsData used by Order
    type UserDetailsData struct {

        
            Address1 string  `json:"address1"`
            AddressType string  `json:"address_type"`
            Address string  `json:"address"`
            Pincode string  `json:"pincode"`
            Area string  `json:"area"`
            State string  `json:"state"`
            Country string  `json:"country"`
            Email string  `json:"email"`
            Landmark string  `json:"landmark"`
            City string  `json:"city"`
            Name string  `json:"name"`
            Phone string  `json:"phone"`
         
    }
    
    // ShipmentStatusData used by Order
    type ShipmentStatusData struct {

        
            ShipmentID string  `json:"shipment_id"`
            BagList []string  `json:"bag_list"`
            Status string  `json:"status"`
            ID float64  `json:"id"`
            CreatedAt string  `json:"created_at"`
         
    }
    
    // OrderingStoreDetails used by Order
    type OrderingStoreDetails struct {

        
            Code string  `json:"code"`
            Address string  `json:"address"`
            OrderingStoreID float64  `json:"ordering_store_id"`
            Pincode string  `json:"pincode"`
            ContactPerson string  `json:"contact_person"`
            State string  `json:"state"`
            Country string  `json:"country"`
            StoreName string  `json:"store_name"`
            City string  `json:"city"`
            Meta map[string]interface{}  `json:"meta"`
            Phone string  `json:"phone"`
         
    }
    
    // PlatformShipment used by Order
    type PlatformShipment struct {

        
            TotalBags float64  `json:"total_bags"`
            LockStatus bool  `json:"lock_status"`
            Payments ShipmentPayments  `json:"payments"`
            ShipmentID string  `json:"shipment_id"`
            CustomMeta []map[string]interface{}  `json:"custom_meta"`
            OperationalStatus string  `json:"operational_status"`
            ShipmentStatus string  `json:"shipment_status"`
            ShipmentQuantity float64  `json:"shipment_quantity"`
            DeliverySlot map[string]interface{}  `json:"delivery_slot"`
            PlatformLogo string  `json:"platform_logo"`
            PriorityText string  `json:"priority_text"`
            BagStatusHistory []BagStatusHistory  `json:"bag_status_history"`
            FulfillingStore FulfillingStore  `json:"fulfilling_store"`
            CompanyDetails CompanyDetails  `json:"company_details"`
            TrackingList []TrackingList  `json:"tracking_list"`
            InvoiceID string  `json:"invoice_id"`
            GstDetails GSTDetailsData  `json:"gst_details"`
            Vertical string  `json:"vertical"`
            Order OrderDetailsData  `json:"order"`
            Prices Prices  `json:"prices"`
            ShipmentImages []string  `json:"shipment_images"`
            DpDetails DPDetailsData  `json:"dp_details"`
            PickedDate string  `json:"picked_date"`
            Bags []OrderBags  `json:"bags"`
            EnableDpTracking bool  `json:"enable_dp_tracking"`
            AffiliateDetails AffiliateDetails  `json:"affiliate_details"`
            Invoice InvoiceInfo  `json:"invoice"`
            User UserDataInfo  `json:"user"`
            UserAgent string  `json:"user_agent"`
            PackagingType string  `json:"packaging_type"`
            ForwardShipmentID string  `json:"forward_shipment_id"`
            PaymentMethods map[string]interface{}  `json:"payment_methods"`
            Meta Meta  `json:"meta"`
            BillingDetails UserDetailsData  `json:"billing_details"`
            FulfilmentPriority float64  `json:"fulfilment_priority"`
            TotalItems float64  `json:"total_items"`
            Status ShipmentStatusData  `json:"status"`
            OrderingStore OrderingStoreDetails  `json:"ordering_store"`
            DeliveryDetails UserDetailsData  `json:"delivery_details"`
            JourneyType string  `json:"journey_type"`
            Coupon map[string]interface{}  `json:"coupon"`
            PaymentMode string  `json:"payment_mode"`
         
    }
    
    // ShipmentInfoResponse used by Order
    type ShipmentInfoResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
            Shipments []PlatformShipment  `json:"shipments"`
         
    }
    
    // PlatformUserDetails used by Order
    type PlatformUserDetails struct {

        
            PlatformUserID string  `json:"platform_user_id"`
            PlatformUserFirstName string  `json:"platform_user_first_name"`
            PlatformUserLastName string  `json:"platform_user_last_name"`
            PlatformUserEmployeeCode string  `json:"platform_user_employee_code"`
         
    }
    
    // TransactionData used by Order
    type TransactionData struct {

        
            UniqueReferenceNumber string  `json:"unique_reference_number"`
            TransactionID string  `json:"transaction_id"`
            TerminalID string  `json:"terminal_id"`
            Status string  `json:"status"`
            PaymentID string  `json:"payment_id"`
            Currency string  `json:"currency"`
            AmountPaid string  `json:"amount_paid"`
            Entity string  `json:"entity"`
         
    }
    
    // BillingStaffDetails used by Order
    type BillingStaffDetails struct {

        
            FirstName string  `json:"first_name"`
            LastName string  `json:"last_name"`
            User string  `json:"user"`
            StaffID float64  `json:"staff_id"`
            EmployeeCode string  `json:"employee_code"`
         
    }
    
    // OrderMeta used by Order
    type OrderMeta struct {

        
            Comment string  `json:"comment"`
            OrderType string  `json:"order_type"`
            CompanyLogo string  `json:"company_logo"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            PaymentType string  `json:"payment_type"`
            PlatformUserDetails PlatformUserDetails  `json:"platform_user_details"`
            Staff map[string]interface{}  `json:"staff"`
            OrderPlatform string  `json:"order_platform"`
            OrderChildEntities []string  `json:"order_child_entities"`
            CartID float64  `json:"cart_id"`
            OrderTags []map[string]interface{}  `json:"order_tags"`
            CustomerNote string  `json:"customer_note"`
            CurrencySymbol string  `json:"currency_symbol"`
            MongoCartID float64  `json:"mongo_cart_id"`
            TransactionData TransactionData  `json:"transaction_data"`
            BillingStaffDetails BillingStaffDetails  `json:"billing_staff_details"`
            Files []map[string]interface{}  `json:"files"`
            OrderingStore float64  `json:"ordering_store"`
            EmployeeID float64  `json:"employee_id"`
         
    }
    
    // TaxDetails used by Order
    type TaxDetails struct {

        
            PanNo string  `json:"pan_no"`
            Gstin string  `json:"gstin"`
         
    }
    
    // OrderDict used by Order
    type OrderDict struct {

        
            OrderDate string  `json:"order_date"`
            FyndOrderID string  `json:"fynd_order_id"`
            Meta OrderMeta  `json:"meta"`
            PaymentMethods map[string]interface{}  `json:"payment_methods"`
            Prices Prices  `json:"prices"`
            TaxDetails TaxDetails  `json:"tax_details"`
         
    }
    
    // ShipmentDetailsResponse used by Order
    type ShipmentDetailsResponse struct {

        
            Order OrderDict  `json:"order"`
            Shipments []PlatformShipment  `json:"shipments"`
            Success bool  `json:"success"`
         
    }
    
    // SubLane used by Order
    type SubLane struct {

        
            Text string  `json:"text"`
            Index float64  `json:"index"`
            TotalItems float64  `json:"total_items"`
            Actions []map[string]interface{}  `json:"actions"`
            Value string  `json:"value"`
         
    }
    
    // SuperLane used by Order
    type SuperLane struct {

        
            TotalItems float64  `json:"total_items"`
            Text string  `json:"text"`
            Options []SubLane  `json:"options"`
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

        
            Value string  `json:"value"`
            Name string  `json:"name"`
            Display string  `json:"display"`
         
    }
    
    // PlatformOrderItems used by Order
    type PlatformOrderItems struct {

        
            Shipments []PlatformShipment  `json:"shipments"`
            OrderID string  `json:"order_id"`
            Meta map[string]interface{}  `json:"meta"`
            Channel PlatformChannel  `json:"channel"`
            TotalOrderValue float64  `json:"total_order_value"`
            UserInfo UserDataInfo  `json:"user_info"`
            BreakupValues []PlatformBreakupValues  `json:"breakup_values"`
            PaymentMode string  `json:"payment_mode"`
            OrderValue float64  `json:"order_value"`
            OrderCreatedTime string  `json:"order_created_time"`
         
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
    
    // OrderListingResponse used by Order
    type OrderListingResponse struct {

        
            Message string  `json:"message"`
            Lane string  `json:"lane"`
            Items []PlatformOrderItems  `json:"items"`
            TotalCount float64  `json:"total_count"`
            Page Page  `json:"page"`
            Success bool  `json:"success"`
         
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

        
            Reason string  `json:"reason"`
            Status string  `json:"status"`
            Awb string  `json:"awb"`
            Meta map[string]interface{}  `json:"meta"`
            LastLocationRecievedAt string  `json:"last_location_recieved_at"`
            AccountName string  `json:"account_name"`
            RawStatus string  `json:"raw_status"`
            UpdatedTime string  `json:"updated_time"`
            UpdatedAt string  `json:"updated_at"`
            ShipmentType string  `json:"shipment_type"`
         
    }
    
    // PlatformShipmentTrack used by Order
    type PlatformShipmentTrack struct {

        
            Results []PlatformTrack  `json:"results"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // FiltersResponse used by Order
    type FiltersResponse struct {

        
            Advance []map[string]interface{}  `json:"advance"`
         
    }
    
    // Success used by Order
    type Success struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // OmsReports used by Order
    type OmsReports struct {

        
            ReportRequestedAt string  `json:"report_requested_at"`
            S3Key string  `json:"s3_key"`
            ReportCreatedAt string  `json:"report_created_at"`
            DisplayName string  `json:"display_name"`
            Status string  `json:"status"`
            ReportID string  `json:"report_id"`
            ReportType string  `json:"report_type"`
            RequestDetails map[string]interface{}  `json:"request_details"`
            ReportName string  `json:"report_name"`
         
    }
    
    // JioCodeUpsertDataSet used by Order
    type JioCodeUpsertDataSet struct {

        
            ArticleID string  `json:"article_id"`
            ItemID string  `json:"item_id"`
            CompanyID string  `json:"company_id"`
            JioCode string  `json:"jio_code"`
         
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

        
            Data []map[string]interface{}  `json:"data"`
            Success bool  `json:"success"`
            TraceID string  `json:"trace_id"`
            Identifier string  `json:"identifier"`
            Error []NestedErrorSchemaDataSet  `json:"error"`
         
    }
    
    // BulkInvoicingResponse used by Order
    type BulkInvoicingResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // BulkInvoiceLabelResponse used by Order
    type BulkInvoiceLabelResponse struct {

        
            CompanyID string  `json:"company_id"`
            Invoice map[string]interface{}  `json:"invoice"`
            Data map[string]interface{}  `json:"data"`
            InvoiceStatus string  `json:"invoice_status"`
            StoreCode string  `json:"store_code"`
            StoreID string  `json:"store_id"`
            BatchID string  `json:"batch_id"`
            Label map[string]interface{}  `json:"label"`
            StoreName string  `json:"store_name"`
            DoInvoiceLabelGenerated bool  `json:"do_invoice_label_generated"`
         
    }
    
    // URL used by Order
    type URL struct {

        
            URL string  `json:"url"`
         
    }
    
    // FileUploadResponse used by Order
    type FileUploadResponse struct {

        
            URL string  `json:"url"`
            Expiry float64  `json:"expiry"`
         
    }
    
    // FileResponse used by Order
    type FileResponse struct {

        
            FilePath string  `json:"file_path"`
            Size float64  `json:"size"`
            Method string  `json:"method"`
            FileName string  `json:"file_name"`
            Operation string  `json:"operation"`
            ContentType string  `json:"content_type"`
            Tags []string  `json:"tags"`
            Namespace string  `json:"namespace"`
            Cdn URL  `json:"cdn"`
            Upload FileUploadResponse  `json:"upload"`
         
    }
    
    // bulkListingData used by Order
    type bulkListingData struct {

        
            FailedShipments []map[string]interface{}  `json:"failed_shipments"`
            ID string  `json:"id"`
            ProcessingShipments []string  `json:"processing_shipments"`
            ExcelURL string  `json:"excel_url"`
            CompanyID float64  `json:"company_id"`
            StoreCode string  `json:"store_code"`
            SuccessfulShipments []map[string]interface{}  `json:"successful_shipments"`
            Successful float64  `json:"successful"`
            Processing float64  `json:"processing"`
            FileName string  `json:"file_name"`
            Total float64  `json:"total"`
            UploadedOn string  `json:"uploaded_on"`
            UserID string  `json:"user_id"`
            Status string  `json:"status"`
            Failed float64  `json:"failed"`
            StoreID float64  `json:"store_id"`
            BatchID string  `json:"batch_id"`
            StoreName string  `json:"store_name"`
            UserName string  `json:"user_name"`
         
    }
    
    // BulkListingPage used by Order
    type BulkListingPage struct {

        
            Current float64  `json:"current"`
            Type string  `json:"type"`
            Size float64  `json:"size"`
            HasPrevious bool  `json:"has_previous"`
            Total float64  `json:"total"`
            HasNext bool  `json:"has_next"`
         
    }
    
    // BulkListingResponse used by Order
    type BulkListingResponse struct {

        
            Data []bulkListingData  `json:"data"`
            Success bool  `json:"success"`
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

        
            ID float64  `json:"id"`
            DisplayName string  `json:"display_name"`
            QcType []string  `json:"qc_type"`
            QuestionSet []QuestionSet  `json:"question_set"`
         
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

        
            CompanyID string  `json:"company_id"`
            SuccessfulShipmentsCount float64  `json:"successful_shipments_count"`
            ProcessingShipmentsCount float64  `json:"processing_shipments_count"`
            TotalShipmentsCount float64  `json:"total_shipments_count"`
            BatchID string  `json:"batch_id"`
            SuccessfulShipmentIds []string  `json:"successful_shipment_ids"`
            FailedShipmentsCount float64  `json:"failed_shipments_count"`
         
    }
    
    // BulkActionDetailsResponse used by Order
    type BulkActionDetailsResponse struct {

        
            Message string  `json:"message"`
            UploadedOn string  `json:"uploaded_on"`
            Data []BulkActionDetailsDataField  `json:"data"`
            UploadedBy string  `json:"uploaded_by"`
            UserID string  `json:"user_id"`
            Status bool  `json:"status"`
            Success string  `json:"success"`
            FailedRecords []string  `json:"failed_records"`
            Error []string  `json:"error"`
         
    }
    
    // Attributes used by Order
    type Attributes struct {

        
            BrandName string  `json:"brand_name"`
            MarketerName string  `json:"marketer_name"`
            Gender []string  `json:"gender"`
            PrimaryMaterial string  `json:"primary_material"`
            Essential string  `json:"essential"`
            MarketerAddress string  `json:"marketer_address"`
            PrimaryColor string  `json:"primary_color"`
            PrimaryColorHex string  `json:"primary_color_hex"`
            Name string  `json:"name"`
         
    }
    
    // Item used by Order
    type Item struct {

        
            WebstoreProductURL string  `json:"webstore_product_url"`
            SlugKey string  `json:"slug_key"`
            LastUpdatedAt string  `json:"last_updated_at"`
            L2Category []string  `json:"l2_category"`
            Gender string  `json:"gender"`
            Size string  `json:"size"`
            Color string  `json:"color"`
            L2CategoryID float64  `json:"l2_category_id"`
            CanReturn bool  `json:"can_return"`
            ItemID float64  `json:"item_id"`
            Attributes Attributes  `json:"attributes"`
            DepartmentID float64  `json:"department_id"`
            Brand string  `json:"brand"`
            BrandID float64  `json:"brand_id"`
            L3CategoryName string  `json:"l3_category_name"`
            Meta map[string]interface{}  `json:"meta"`
            L1CategoryID float64  `json:"l1_category_id"`
            Code string  `json:"code"`
            L3Category float64  `json:"l3_category"`
            L1Category []string  `json:"l1_category"`
            BranchURL string  `json:"branch_url"`
            CanCancel bool  `json:"can_cancel"`
            Name string  `json:"name"`
            Image []string  `json:"image"`
         
    }
    
    // Dates used by Order
    type Dates struct {

        
            DeliveryDate interface{}  `json:"delivery_date"`
            OrderCreated string  `json:"order_created"`
         
    }
    
    // AffiliateBagDetails used by Order
    type AffiliateBagDetails struct {

        
            AffiliateOrderID string  `json:"affiliate_order_id"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            EmployeeDiscount float64  `json:"employee_discount"`
            AffiliateMeta AffiliateMeta  `json:"affiliate_meta"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
         
    }
    
    // BagGSTDetails used by Order
    type BagGSTDetails struct {

        
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            HsnCodeID string  `json:"hsn_code_id"`
            IgstTaxPercentage float64  `json:"igst_tax_percentage"`
            IsDefaultHsnCode bool  `json:"is_default_hsn_code"`
            ValueOfGood float64  `json:"value_of_good"`
            CgstGstFee string  `json:"cgst_gst_fee"`
            GstTag string  `json:"gst_tag"`
            GstFee float64  `json:"gst_fee"`
            IgstGstFee string  `json:"igst_gst_fee"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            SgstTaxPercentage float64  `json:"sgst_tax_percentage"`
            SgstGstFee string  `json:"sgst_gst_fee"`
            CgstTaxPercentage float64  `json:"cgst_tax_percentage"`
            GstinCode string  `json:"gstin_code"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            HsnCode string  `json:"hsn_code"`
         
    }
    
    // ArticleDetails used by Order
    type ArticleDetails struct {

        
            Status map[string]interface{}  `json:"status"`
         
    }
    
    // Brand used by Order
    type Brand struct {

        
            IsVirtualInvoice bool  `json:"is_virtual_invoice"`
            BrandName string  `json:"brand_name"`
            StartDate string  `json:"start_date"`
            CreditNoteExpiryDays float64  `json:"credit_note_expiry_days"`
            InvoicePrefix string  `json:"invoice_prefix"`
            Logo string  `json:"logo"`
            Company string  `json:"company"`
            CreatedOn float64  `json:"created_on"`
            BrandID float64  `json:"brand_id"`
            ModifiedOn float64  `json:"modified_on"`
            ScriptLastRan string  `json:"script_last_ran"`
            CreditNoteAllowed bool  `json:"credit_note_allowed"`
            PickupLocation string  `json:"pickup_location"`
         
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
            Identifiers Identifier  `json:"identifiers"`
            EspModified interface{}  `json:"esp_modified"`
            IsSet bool  `json:"is_set"`
            SellerIdentifier string  `json:"seller_identifier"`
            ReturnConfig ReturnConfig  `json:"return_config"`
            Size string  `json:"size"`
            RawMeta interface{}  `json:"raw_meta"`
            ChildDetails map[string]interface{}  `json:"child_details"`
            Dimensions Dimensions  `json:"dimensions"`
            Weight Weight  `json:"weight"`
            UID string  `json:"uid"`
            ID string  `json:"_id"`
            ASet map[string]interface{}  `json:"a_set"`
         
    }
    
    // B2BPODetails used by Order
    type B2BPODetails struct {

        
            ItemBasePrice float64  `json:"item_base_price"`
            PoTaxAmount float64  `json:"po_tax_amount"`
            PoLineAmount float64  `json:"po_line_amount"`
            DockerNumber string  `json:"docker_number"`
            PartialCanRet bool  `json:"partial_can_ret"`
            TotalGstPercentage float64  `json:"total_gst_percentage"`
         
    }
    
    // BagMeta used by Order
    type BagMeta struct {

        
            B2bPoDetails B2BPODetails  `json:"b2b_po_details"`
         
    }
    
    // BagReturnableCancelableStatus used by Order
    type BagReturnableCancelableStatus struct {

        
            IsCustomerReturnAllowed bool  `json:"is_customer_return_allowed"`
            IsReturnable bool  `json:"is_returnable"`
            CanBeCancelled bool  `json:"can_be_cancelled"`
            IsActive bool  `json:"is_active"`
            EnableTracking bool  `json:"enable_tracking"`
         
    }
    
    // StoreAddress used by Order
    type StoreAddress struct {

        
            AddressCategory string  `json:"address_category"`
            Longitude float64  `json:"longitude"`
            Area string  `json:"area"`
            State string  `json:"state"`
            Country string  `json:"country"`
            CountryCode string  `json:"country_code"`
            AddressType string  `json:"address_type"`
            Latitude float64  `json:"latitude"`
            Email string  `json:"email"`
            Landmark string  `json:"landmark"`
            Address2 string  `json:"address2"`
            Phone string  `json:"phone"`
            Address1 string  `json:"address1"`
            Pincode float64  `json:"pincode"`
            ContactPerson string  `json:"contact_person"`
            CreatedAt string  `json:"created_at"`
            Version string  `json:"version"`
            City string  `json:"city"`
            UpdatedAt string  `json:"updated_at"`
         
    }
    
    // Document used by Order
    type Document struct {

        
            DsType string  `json:"ds_type"`
            LegalName string  `json:"legal_name"`
            Verified bool  `json:"verified"`
            URL string  `json:"url"`
            Value string  `json:"value"`
         
    }
    
    // StoreDocuments used by Order
    type StoreDocuments struct {

        
            Gst Document  `json:"gst"`
         
    }
    
    // EInvoicePortalDetails used by Order
    type EInvoicePortalDetails struct {

        
            Password string  `json:"password"`
            Username string  `json:"username"`
            User string  `json:"user"`
         
    }
    
    // StoreEwaybill used by Order
    type StoreEwaybill struct {

        
            Enabled bool  `json:"enabled"`
         
    }
    
    // StoreEinvoice used by Order
    type StoreEinvoice struct {

        
            Password string  `json:"password"`
            Enabled bool  `json:"enabled"`
            Username string  `json:"username"`
            User string  `json:"user"`
         
    }
    
    // StoreGstCredentials used by Order
    type StoreGstCredentials struct {

        
            EWaybill StoreEwaybill  `json:"e_waybill"`
            EInvoice StoreEinvoice  `json:"e_invoice"`
         
    }
    
    // StoreMeta used by Order
    type StoreMeta struct {

        
            Timing []map[string]interface{}  `json:"timing"`
            EwaybillPortalDetails map[string]interface{}  `json:"ewaybill_portal_details"`
            Documents StoreDocuments  `json:"documents"`
            AdditionalContactDetails map[string]interface{}  `json:"additional_contact_details"`
            NotificationEmails []string  `json:"notification_emails"`
            ProductReturnConfig map[string]interface{}  `json:"product_return_config"`
            DisplayName string  `json:"display_name"`
            Stage string  `json:"stage"`
            EinvoicePortalDetails EInvoicePortalDetails  `json:"einvoice_portal_details"`
            GstCredentials StoreGstCredentials  `json:"gst_credentials"`
            GstNumber string  `json:"gst_number"`
         
    }
    
    // Store used by Order
    type Store struct {

        
            Longitude float64  `json:"longitude"`
            AlohomoraUserID float64  `json:"alohomora_user_id"`
            StoreAddressJson StoreAddress  `json:"store_address_json"`
            MallName string  `json:"mall_name"`
            CompanyID float64  `json:"company_id"`
            IsArchived bool  `json:"is_archived"`
            PackagingMaterialCount float64  `json:"packaging_material_count"`
            MallArea string  `json:"mall_area"`
            StoreActiveFrom string  `json:"store_active_from"`
            VatNo string  `json:"vat_no"`
            State string  `json:"state"`
            Country string  `json:"country"`
            LoginUsername string  `json:"login_username"`
            OrderIntegrationID string  `json:"order_integration_id"`
            ParentStoreID float64  `json:"parent_store_id"`
            LocationType string  `json:"location_type"`
            FulfillmentChannel string  `json:"fulfillment_channel"`
            Latitude float64  `json:"latitude"`
            BrandID interface{}  `json:"brand_id"`
            IsActive bool  `json:"is_active"`
            IsEnabledForRecon bool  `json:"is_enabled_for_recon"`
            Address2 string  `json:"address2"`
            Meta StoreMeta  `json:"meta"`
            Phone float64  `json:"phone"`
            Code string  `json:"code"`
            Address1 string  `json:"address1"`
            BrandStoreTags []string  `json:"brand_store_tags"`
            SID string  `json:"s_id"`
            Pincode string  `json:"pincode"`
            ContactPerson string  `json:"contact_person"`
            UpdatedAt string  `json:"updated_at"`
            CreatedAt string  `json:"created_at"`
            City string  `json:"city"`
            Name string  `json:"name"`
            StoreEmail string  `json:"store_email"`
         
    }
    
    // BagDetailsPlatformResponse used by Order
    type BagDetailsPlatformResponse struct {

        
            Item Item  `json:"item"`
            OriginalBagList []float64  `json:"original_bag_list"`
            ShipmentID string  `json:"shipment_id"`
            OperationalStatus string  `json:"operational_status"`
            Quantity float64  `json:"quantity"`
            Dates Dates  `json:"dates"`
            Tags []string  `json:"tags"`
            CurrentOperationalStatus BagStatusHistory  `json:"current_operational_status"`
            LineNumber float64  `json:"line_number"`
            BagUpdateTime float64  `json:"bag_update_time"`
            AffiliateBagDetails AffiliateBagDetails  `json:"affiliate_bag_details"`
            CurrentStatus BagStatusHistory  `json:"current_status"`
            DisplayName string  `json:"display_name"`
            BagStatusHistory BagStatusHistory  `json:"bag_status_history"`
            OrderIntegrationID string  `json:"order_integration_id"`
            RestorePromos map[string]interface{}  `json:"restore_promos"`
            GstDetails BagGSTDetails  `json:"gst_details"`
            FinancialBreakup []FinancialBreakup  `json:"financial_breakup"`
            Prices Prices  `json:"prices"`
            ArticleDetails ArticleDetails  `json:"article_details"`
            EntityType string  `json:"entity_type"`
            AffiliateDetails AffiliateDetails  `json:"affiliate_details"`
            AppliedPromos []map[string]interface{}  `json:"applied_promos"`
            BID float64  `json:"b_id"`
            QcRequired interface{}  `json:"qc_required"`
            Brand Brand  `json:"brand"`
            Identifier string  `json:"identifier"`
            Article Article  `json:"article"`
            Meta BagMeta  `json:"meta"`
            BType string  `json:"b_type"`
            NoOfBagsOrder float64  `json:"no_of_bags_order"`
            SellerIdentifier string  `json:"seller_identifier"`
            Reasons []map[string]interface{}  `json:"reasons"`
            Status BagReturnableCancelableStatus  `json:"status"`
            ParentPromoBags map[string]interface{}  `json:"parent_promo_bags"`
            RestoreCoupon bool  `json:"restore_coupon"`
            OrderingStore Store  `json:"ordering_store"`
            JourneyType string  `json:"journey_type"`
            BagStatus []BagStatusHistory  `json:"bag_status"`
         
    }
    
    // ErrorResponse used by Order
    type ErrorResponse struct {

        
            Message string  `json:"message"`
            Error string  `json:"error"`
         
    }
    
    // Page1 used by Order
    type Page1 struct {

        
            PageType string  `json:"page_type"`
            Current float64  `json:"current"`
            Size float64  `json:"size"`
            ItemTotal float64  `json:"item_total"`
            HasNext bool  `json:"has_next"`
         
    }
    
    // GetBagsPlatformResponse used by Order
    type GetBagsPlatformResponse struct {

        
            Items []BagDetailsPlatformResponse  `json:"items"`
            Page Page1  `json:"page"`
         
    }
    
    // InvalidateShipmentCachePayload used by Order
    type InvalidateShipmentCachePayload struct {

        
            ShipmentIds []string  `json:"shipment_ids"`
         
    }
    
    // InvalidateShipmentCacheNestedResponse used by Order
    type InvalidateShipmentCacheNestedResponse struct {

        
            Message string  `json:"message"`
            ShipmentID string  `json:"shipment_id"`
            Status float64  `json:"status"`
            Error string  `json:"error"`
         
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

        
            AffiliateOrderID string  `json:"affiliate_order_id"`
            ItemID string  `json:"item_id"`
            AffiliateID string  `json:"affiliate_id"`
            BagID float64  `json:"bag_id"`
            MongoArticleID string  `json:"mongo_article_id"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
            StoreID float64  `json:"store_id"`
            FyndOrderID string  `json:"fynd_order_id"`
            SetID string  `json:"set_id"`
            ReasonIds []float64  `json:"reason_ids"`
         
    }
    
    // StoreReassignResponse used by Order
    type StoreReassignResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // Entities used by Order
    type Entities struct {

        
            AffiliateOrderID string  `json:"affiliate_order_id"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            AffiliateID string  `json:"affiliate_id"`
            ReasonText string  `json:"reason_text"`
            ID string  `json:"id"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
         
    }
    
    // UpdateShipmentLockPayload used by Order
    type UpdateShipmentLockPayload struct {

        
            ActionType string  `json:"action_type"`
            EntityType string  `json:"entity_type"`
            Action string  `json:"action"`
            Entities []Entities  `json:"entities"`
         
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

        
            ShipmentID string  `json:"shipment_id"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            LockStatus bool  `json:"lock_status"`
            AffiliateID string  `json:"affiliate_id"`
            IsBagLocked bool  `json:"is_bag_locked"`
            IsShipmentLocked bool  `json:"is_shipment_locked"`
            Status string  `json:"status"`
            Bags []Bags  `json:"bags"`
            OriginalFilter OriginalFilter  `json:"original_filter"`
         
    }
    
    // UpdateShipmentLockResponse used by Order
    type UpdateShipmentLockResponse struct {

        
            Message string  `json:"message"`
            CheckResponse []CheckResponse  `json:"check_response"`
            Success bool  `json:"success"`
         
    }
    
    // AnnouncementResponse used by Order
    type AnnouncementResponse struct {

        
            FromDatetime string  `json:"from_datetime"`
            Title string  `json:"title"`
            PlatformID string  `json:"platform_id"`
            LogoURL string  `json:"logo_url"`
            ID float64  `json:"id"`
            CompanyID float64  `json:"company_id"`
            PlatformName string  `json:"platform_name"`
            Description string  `json:"description"`
            ToDatetime string  `json:"to_datetime"`
            CreatedAt string  `json:"created_at"`
         
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

        
            Status bool  `json:"status"`
            CallID string  `json:"call_id"`
         
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

        
            ReasonText string  `json:"reason_text"`
            ReasonID float64  `json:"reason_id"`
         
    }
    
    // ProductsReasonsFilters used by Order
    type ProductsReasonsFilters struct {

        
            LineNumber float64  `json:"line_number"`
            Identifier string  `json:"identifier"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // ProductsReasons used by Order
    type ProductsReasons struct {

        
            Data ProductsReasonsData  `json:"data"`
            Filters []ProductsReasonsFilters  `json:"filters"`
         
    }
    
    // EntityReasonData used by Order
    type EntityReasonData struct {

        
            ReasonText string  `json:"reason_text"`
            ReasonID float64  `json:"reason_id"`
         
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
    
    // Products used by Order
    type Products struct {

        
            LineNumber float64  `json:"line_number"`
            Identifier string  `json:"identifier"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // ShipmentsRequest used by Order
    type ShipmentsRequest struct {

        
            DataUpdates DataUpdates  `json:"data_updates"`
            Identifier string  `json:"identifier"`
            Reasons ReasonsData  `json:"reasons"`
            Products []Products  `json:"products"`
         
    }
    
    // StatuesRequest used by Order
    type StatuesRequest struct {

        
            Status string  `json:"status"`
            ExcludeBagsNextState string  `json:"exclude_bags_next_state"`
            Shipments []ShipmentsRequest  `json:"shipments"`
         
    }
    
    // UpdateShipmentStatusRequest used by Order
    type UpdateShipmentStatusRequest struct {

        
            LockAfterTransition bool  `json:"lock_after_transition"`
            Statuses []StatuesRequest  `json:"statuses"`
            Task bool  `json:"task"`
            UnlockBeforeTransition bool  `json:"unlock_before_transition"`
            ForceTransition bool  `json:"force_transition"`
         
    }
    
    // ShipmentsResponse used by Order
    type ShipmentsResponse struct {

        
            FinalState map[string]interface{}  `json:"final_state"`
            Message string  `json:"message"`
            Identifier string  `json:"identifier"`
            Meta map[string]interface{}  `json:"meta"`
            StackTrace string  `json:"stack_trace"`
            Code string  `json:"code"`
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
    
    // AffiliateInventoryOrderConfig used by Order
    type AffiliateInventoryOrderConfig struct {

        
            ForceReassignment bool  `json:"force_reassignment"`
         
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
            Payment AffiliateInventoryPaymentConfig  `json:"payment"`
            Logistics AffiliateInventoryLogisticsConfig  `json:"logistics"`
            ArticleAssignment AffiliateInventoryArticleAssignmentConfig  `json:"article_assignment"`
            Inventory AffiliateInventoryStoreConfig  `json:"inventory"`
         
    }
    
    // AffiliateAppConfigMeta used by Order
    type AffiliateAppConfigMeta struct {

        
            Value string  `json:"value"`
            Name string  `json:"name"`
         
    }
    
    // AffiliateAppConfig used by Order
    type AffiliateAppConfig struct {

        
            Owner string  `json:"owner"`
            Name string  `json:"name"`
            Secret string  `json:"secret"`
            Meta []AffiliateAppConfigMeta  `json:"meta"`
            ID string  `json:"id"`
            UpdatedAt string  `json:"updated_at"`
            Token string  `json:"token"`
            Description string  `json:"description"`
            CreatedAt string  `json:"created_at"`
         
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

        
            MarketplaceStoreID string  `json:"marketplace_store_id"`
            StoreID float64  `json:"store_id"`
         
    }
    
    // OrderConfig used by Order
    type OrderConfig struct {

        
            BagEndState string  `json:"bag_end_state"`
            StoreLookup string  `json:"store_lookup"`
            Affiliate Affiliate  `json:"affiliate"`
            CreateUser bool  `json:"create_user"`
            ArticleLookup string  `json:"article_lookup"`
            AffiliateStoreIDMapping []AffiliateStoreIdMapping  `json:"affiliate_store_id_mapping"`
         
    }
    
    // ArticleDetails1 used by Order
    type ArticleDetails1 struct {

        
            BrandID float64  `json:"brand_id"`
            Attributes map[string]interface{}  `json:"attributes"`
            Category map[string]interface{}  `json:"category"`
            Quantity float64  `json:"quantity"`
            Weight map[string]interface{}  `json:"weight"`
            ID string  `json:"_id"`
            Dimension map[string]interface{}  `json:"dimension"`
         
    }
    
    // ShipmentDetails used by Order
    type ShipmentDetails struct {

        
            DpID float64  `json:"dp_id"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            BoxType string  `json:"box_type"`
            Shipments float64  `json:"shipments"`
            Meta map[string]interface{}  `json:"meta"`
            Articles []ArticleDetails1  `json:"articles"`
            FulfillmentID float64  `json:"fulfillment_id"`
         
    }
    
    // LocationDetails used by Order
    type LocationDetails struct {

        
            Articles []ArticleDetails1  `json:"articles"`
            FulfillmentID float64  `json:"fulfillment_id"`
            FulfillmentType string  `json:"fulfillment_type"`
         
    }
    
    // ShipmentConfig used by Order
    type ShipmentConfig struct {

        
            Shipment []ShipmentDetails  `json:"shipment"`
            PaymentMode string  `json:"payment_mode"`
            Identifier string  `json:"identifier"`
            LocationDetails LocationDetails  `json:"location_details"`
            ToPincode string  `json:"to_pincode"`
            Journey string  `json:"journey"`
            Source string  `json:"source"`
            Action string  `json:"action"`
         
    }
    
    // ShipmentData used by Order
    type ShipmentData struct {

        
            ShipmentData ShipmentConfig  `json:"shipment_data"`
         
    }
    
    // OrderPriority used by Order
    type OrderPriority struct {

        
            FulfilmentPriorityText string  `json:"fulfilment_priority_text"`
            FulfilmentPriority float64  `json:"fulfilment_priority"`
            AffiliatePriorityCode string  `json:"affiliate_priority_code"`
         
    }
    
    // OrderUser used by Order
    type OrderUser struct {

        
            Phone float64  `json:"phone"`
            Pincode string  `json:"pincode"`
            FirstName string  `json:"first_name"`
            City string  `json:"city"`
            State string  `json:"state"`
            Address2 string  `json:"address2"`
            LastName string  `json:"last_name"`
            Country string  `json:"country"`
            Address1 string  `json:"address1"`
            Email string  `json:"email"`
            Mobile float64  `json:"mobile"`
         
    }
    
    // UserData used by Order
    type UserData struct {

        
            ShippingUser OrderUser  `json:"shipping_user"`
            BillingUser OrderUser  `json:"billing_user"`
         
    }
    
    // MarketPlacePdf used by Order
    type MarketPlacePdf struct {

        
            Label string  `json:"label"`
            Invoice string  `json:"invoice"`
         
    }
    
    // AffiliateBag used by Order
    type AffiliateBag struct {

        
            AmountPaid float64  `json:"amount_paid"`
            UnitPrice float64  `json:"unit_price"`
            ModifiedOn string  `json:"modified_on"`
            FyndStoreID string  `json:"fynd_store_id"`
            Quantity float64  `json:"quantity"`
            PriceEffective float64  `json:"price_effective"`
            PdfLinks MarketPlacePdf  `json:"pdf_links"`
            SellerIdentifier string  `json:"seller_identifier"`
            ItemID float64  `json:"item_id"`
            PriceMarked float64  `json:"price_marked"`
            TransferPrice float64  `json:"transfer_price"`
            StoreID float64  `json:"store_id"`
            Identifier map[string]interface{}  `json:"identifier"`
            Discount float64  `json:"discount"`
            Sku string  `json:"sku"`
            HsnCodeID string  `json:"hsn_code_id"`
            DeliveryCharge float64  `json:"delivery_charge"`
            CompanyID float64  `json:"company_id"`
            ItemSize string  `json:"item_size"`
            ID string  `json:"_id"`
            AvlQty float64  `json:"avl_qty"`
            AffiliateStoreID string  `json:"affiliate_store_id"`
            AffiliateMeta map[string]interface{}  `json:"affiliate_meta"`
         
    }
    
    // OrderInfo used by Order
    type OrderInfo struct {

        
            Shipment ShipmentData  `json:"shipment"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            Payment map[string]interface{}  `json:"payment"`
            DeliveryCharges float64  `json:"delivery_charges"`
            OrderPriority OrderPriority  `json:"order_priority"`
            OrderValue float64  `json:"order_value"`
            PaymentMode string  `json:"payment_mode"`
            ShippingAddress OrderUser  `json:"shipping_address"`
            BillingAddress OrderUser  `json:"billing_address"`
            Discount float64  `json:"discount"`
            Items map[string]interface{}  `json:"items"`
            CodCharges float64  `json:"cod_charges"`
            User UserData  `json:"user"`
            Coupon string  `json:"coupon"`
            Bags []AffiliateBag  `json:"bags"`
         
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

        
            ID float64  `json:"id"`
            Slug string  `json:"slug"`
            DisplayText string  `json:"display_text"`
            Description string  `json:"description"`
         
    }
    
    // GetActionsResponse used by Order
    type GetActionsResponse struct {

        
            Permissions ActionInfo  `json:"permissions"`
         
    }
    
    // HistoryDict used by Order
    type HistoryDict struct {

        
            Type string  `json:"type"`
            L3Detail string  `json:"l3_detail"`
            Message string  `json:"message"`
            L2Detail string  `json:"l2_detail"`
            TicketID string  `json:"ticket_id"`
            BagID float64  `json:"bag_id"`
            L1Detail string  `json:"l1_detail"`
            User string  `json:"user"`
            TicketURL string  `json:"ticket_url"`
            Createdat string  `json:"createdat"`
         
    }
    
    // ShipmentHistoryResponse used by Order
    type ShipmentHistoryResponse struct {

        
            ActivityHistory []HistoryDict  `json:"activity_history"`
         
    }
    
    // ErrorDetail used by Order
    type ErrorDetail struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // PostHistoryData used by Order
    type PostHistoryData struct {

        
            Message string  `json:"message"`
            UserName string  `json:"user_name"`
         
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
    
    // SmsDataPayload used by Order
    type SmsDataPayload struct {

        
            AmountPaid float64  `json:"amount_paid"`
            OrderID string  `json:"order_id"`
            ShipmentID float64  `json:"shipment_id"`
            CustomerName string  `json:"customer_name"`
            BrandName string  `json:"brand_name"`
            Message string  `json:"message"`
            PaymentMode string  `json:"payment_mode"`
            PhoneNumber float64  `json:"phone_number"`
            CountryCode string  `json:"country_code"`
         
    }
    
    // SendSmsPayload used by Order
    type SendSmsPayload struct {

        
            Data SmsDataPayload  `json:"data"`
            Slug string  `json:"slug"`
            BagID float64  `json:"bag_id"`
         
    }
    
    // Meta1 used by Order
    type Meta1 struct {

        
            KafkaEmissionStatus float64  `json:"kafka_emission_status"`
            StateManagerUsed string  `json:"state_manager_used"`
         
    }
    
    // ShipmentDetail used by Order
    type ShipmentDetail struct {

        
            ShipmentID string  `json:"shipment_id"`
            Meta Meta1  `json:"meta"`
            ID float64  `json:"id"`
            Remarks string  `json:"remarks"`
            BagList []float64  `json:"bag_list"`
            Status string  `json:"status"`
         
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
    
    // ManualAssignDPToShipment used by Order
    type ManualAssignDPToShipment struct {

        
            DpID float64  `json:"dp_id"`
            OrderType string  `json:"order_type"`
            QcRequired string  `json:"qc_required"`
            ShipmentIds []string  `json:"shipment_ids"`
         
    }
    
    // ManualAssignDPToShipmentResponse used by Order
    type ManualAssignDPToShipmentResponse struct {

        
            Errors []string  `json:"errors"`
            Success string  `json:"success"`
         
    }
    
    // Tax used by Order
    type Tax struct {

        
            Rate float64  `json:"rate"`
            Name string  `json:"name"`
            Amount map[string]interface{}  `json:"amount"`
            Breakup []map[string]interface{}  `json:"breakup"`
         
    }
    
    // Charge used by Order
    type Charge struct {

        
            Type string  `json:"type"`
            Tax Tax  `json:"tax"`
            Name string  `json:"name"`
            Code string  `json:"code"`
            Amount map[string]interface{}  `json:"amount"`
         
    }
    
    // LineItem used by Order
    type LineItem struct {

        
            Charges []Charge  `json:"charges"`
            SellerIdentifier string  `json:"seller_identifier"`
            Meta map[string]interface{}  `json:"meta"`
            ExternalLineID string  `json:"external_line_id"`
            Quantity float64  `json:"quantity"`
            CustomMessasge string  `json:"custom_messasge"`
         
    }
    
    // ProcessingDates used by Order
    type ProcessingDates struct {

        
            CustomerPickupSlot map[string]interface{}  `json:"customer_pickup_slot"`
            ConfirmByDate string  `json:"confirm_by_date"`
            DispatchByDate string  `json:"dispatch_by_date"`
            PackByDate string  `json:"pack_by_date"`
            DispatchAfterDate string  `json:"dispatch_after_date"`
            DpPickupSlot map[string]interface{}  `json:"dp_pickup_slot"`
         
    }
    
    // Shipment used by Order
    type Shipment struct {

        
            Priority float64  `json:"priority"`
            Meta map[string]interface{}  `json:"meta"`
            LocationID float64  `json:"location_id"`
            LineItems []LineItem  `json:"line_items"`
            ExternalShipmentID string  `json:"external_shipment_id"`
            ProcessingDates ProcessingDates  `json:"processing_dates"`
         
    }
    
    // BillingInfo used by Order
    type BillingInfo struct {

        
            AlternateEmail string  `json:"alternate_email"`
            Pincode string  `json:"pincode"`
            FloorNo string  `json:"floor_no"`
            Title string  `json:"title"`
            AlternateMobileNumber string  `json:"alternate_mobile_number"`
            PrimaryEmail string  `json:"primary_email"`
            Gender string  `json:"gender"`
            CustomerCode string  `json:"customer_code"`
            MiddleName string  `json:"middle_name"`
            LastName string  `json:"last_name"`
            ExternalCustomerCode string  `json:"external_customer_code"`
            StateCode string  `json:"state_code"`
            FirstName string  `json:"first_name"`
            City string  `json:"city"`
            PrimaryMobileNumber string  `json:"primary_mobile_number"`
            HouseNo string  `json:"house_no"`
            Address1 string  `json:"address1"`
            CountryCode string  `json:"country_code"`
            State string  `json:"state"`
            Address2 string  `json:"address2"`
            Country string  `json:"country"`
         
    }
    
    // PaymentMethod used by Order
    type PaymentMethod struct {

        
            Name string  `json:"name"`
            Meta map[string]interface{}  `json:"meta"`
            RefundBy string  `json:"refund_by"`
            CollectBy string  `json:"collect_by"`
            TransactionData map[string]interface{}  `json:"transaction_data"`
            Amount float64  `json:"amount"`
            Mode string  `json:"mode"`
         
    }
    
    // PaymentInfo used by Order
    type PaymentInfo struct {

        
            PaymentMethods []PaymentMethod  `json:"payment_methods"`
            PrimaryMode string  `json:"primary_mode"`
         
    }
    
    // ShippingInfo used by Order
    type ShippingInfo struct {

        
            AlternateEmail string  `json:"alternate_email"`
            Pincode string  `json:"pincode"`
            FloorNo string  `json:"floor_no"`
            Title string  `json:"title"`
            Landmark string  `json:"landmark"`
            Slot []map[string]interface{}  `json:"slot"`
            AlternateMobileNumber string  `json:"alternate_mobile_number"`
            PrimaryEmail string  `json:"primary_email"`
            Gender string  `json:"gender"`
            CustomerCode string  `json:"customer_code"`
            MiddleName string  `json:"middle_name"`
            LastName string  `json:"last_name"`
            ExternalCustomerCode string  `json:"external_customer_code"`
            StateCode string  `json:"state_code"`
            AddressType string  `json:"address_type"`
            FirstName string  `json:"first_name"`
            City string  `json:"city"`
            PrimaryMobileNumber string  `json:"primary_mobile_number"`
            GeoLocation map[string]interface{}  `json:"geo_location"`
            HouseNo string  `json:"house_no"`
            Address1 string  `json:"address1"`
            CountryCode string  `json:"country_code"`
            ShippingType string  `json:"shipping_type"`
            State string  `json:"state"`
            Address2 string  `json:"address2"`
            Country string  `json:"country"`
         
    }
    
    // TaxInfo used by Order
    type TaxInfo struct {

        
            Gstin string  `json:"gstin"`
            B2bGstinNumber string  `json:"b2b_gstin_number"`
         
    }
    
    // CreateOrderAPI used by Order
    type CreateOrderAPI struct {

        
            Charges []Charge  `json:"charges"`
            ExternalOrderID string  `json:"external_order_id"`
            Shipments []Shipment  `json:"shipments"`
            BillingInfo BillingInfo  `json:"billing_info"`
            Meta map[string]interface{}  `json:"meta"`
            Config map[string]interface{}  `json:"config"`
            PaymentInfo PaymentInfo  `json:"payment_info"`
            ShippingInfo ShippingInfo  `json:"shipping_info"`
            TaxInfo TaxInfo  `json:"tax_info"`
            ExternalCreationDate string  `json:"external_creation_date"`
            CurrencyInfo map[string]interface{}  `json:"currency_info"`
         
    }
    
    // CreateOrderErrorReponse used by Order
    type CreateOrderErrorReponse struct {

        
            RequestID string  `json:"request_id"`
            Message string  `json:"message"`
            Meta string  `json:"meta"`
            StackTrace string  `json:"stack_trace"`
            Info interface{}  `json:"info"`
            Code string  `json:"code"`
            Status float64  `json:"status"`
            Exception string  `json:"exception"`
         
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

        
            PaymentMethods []PaymentMethods  `json:"payment_methods"`
            ModeOfPayment string  `json:"mode_of_payment"`
            Source string  `json:"source"`
         
    }
    
    // CreateChannelConfig used by Order
    type CreateChannelConfig struct {

        
            LockStates []string  `json:"lock_states"`
            ShipmentAssignment string  `json:"shipment_assignment"`
            LocationReassignment bool  `json:"location_reassignment"`
            DpConfiguration DpConfiguration  `json:"dp_configuration"`
            LogoURL map[string]interface{}  `json:"logo_url"`
            PaymentInfo CreateChannelPaymentInfo  `json:"payment_info"`
         
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
            IsUpserted bool  `json:"is_upserted"`
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

        
            Message []string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // FyndOrderIdList used by Order
    type FyndOrderIdList struct {

        
            FyndOrderID []string  `json:"fynd_order_id"`
         
    }
    
    // OrderStatus used by Order
    type OrderStatus struct {

        
            OrderDetails []FyndOrderIdList  `json:"order_details"`
            Mobile float64  `json:"mobile"`
            StartDate string  `json:"start_date"`
            EndDate string  `json:"end_date"`
         
    }
    

    
    // DeleteResponse used by Catalog
    type DeleteResponse struct {

        
            Message string  `json:"message"`
         
    }
    
    // ErrorResponse used by Catalog
    type ErrorResponse struct {

        
            Message string  `json:"message"`
            Code string  `json:"code"`
            Error string  `json:"error"`
            Status float64  `json:"status"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // GetSearchWordsData used by Catalog
    type GetSearchWordsData struct {

        
            AppID string  `json:"app_id"`
            Words []string  `json:"words"`
            IsActive bool  `json:"is_active"`
            Result map[string]interface{}  `json:"result"`
            UID string  `json:"uid"`
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
    
    // SearchKeywordResult used by Catalog
    type SearchKeywordResult struct {

        
            Query map[string]interface{}  `json:"query"`
            SortOn string  `json:"sort_on"`
         
    }
    
    // CreateSearchKeyword used by Catalog
    type CreateSearchKeyword struct {

        
            AppID string  `json:"app_id"`
            Words []string  `json:"words"`
            IsActive bool  `json:"is_active"`
            Result SearchKeywordResult  `json:"result"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // GetSearchWordsResponse used by Catalog
    type GetSearchWordsResponse struct {

        
            Items []GetSearchWordsData  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // GetAutocompleteWordsData used by Catalog
    type GetAutocompleteWordsData struct {

        
            AppID string  `json:"app_id"`
            Words []string  `json:"words"`
            Results []map[string]interface{}  `json:"results"`
            UID string  `json:"uid"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // GetAutocompleteWordsResponse used by Catalog
    type GetAutocompleteWordsResponse struct {

        
            Items []GetAutocompleteWordsData  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // AutocompletePageAction used by Catalog
    type AutocompletePageAction struct {

        
            Params map[string]interface{}  `json:"params"`
            Type string  `json:"type"`
            Query map[string]interface{}  `json:"query"`
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
            Action AutocompleteAction  `json:"action"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Logo Media  `json:"logo"`
         
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

        
            Results []map[string]interface{}  `json:"results"`
            Words []string  `json:"words"`
            AppID string  `json:"app_id"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // ProductBundleItem used by Catalog
    type ProductBundleItem struct {

        
            ProductUID float64  `json:"product_uid"`
            AutoAddToCart bool  `json:"auto_add_to_cart"`
            MaxQuantity float64  `json:"max_quantity"`
            AutoSelect bool  `json:"auto_select"`
            AllowRemove bool  `json:"allow_remove"`
            MinQuantity float64  `json:"min_quantity"`
         
    }
    
    // ProductBundleRequest used by Catalog
    type ProductBundleRequest struct {

        
            CompanyID float64  `json:"company_id"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            CreatedOn string  `json:"created_on"`
            Slug string  `json:"slug"`
            Logo string  `json:"logo"`
            IsActive bool  `json:"is_active"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Choice string  `json:"choice"`
            Products []ProductBundleItem  `json:"products"`
            ModifiedOn string  `json:"modified_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Meta map[string]interface{}  `json:"meta"`
            PageVisibility []string  `json:"page_visibility"`
            Name string  `json:"name"`
         
    }
    
    // GetProductBundleCreateResponse used by Catalog
    type GetProductBundleCreateResponse struct {

        
            ID string  `json:"id"`
            CompanyID float64  `json:"company_id"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            CreatedOn string  `json:"created_on"`
            Slug string  `json:"slug"`
            Logo string  `json:"logo"`
            IsActive bool  `json:"is_active"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Choice string  `json:"choice"`
            Products []ProductBundleItem  `json:"products"`
            ModifiedOn string  `json:"modified_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Meta map[string]interface{}  `json:"meta"`
            PageVisibility []string  `json:"page_visibility"`
            Name string  `json:"name"`
         
    }
    
    // GetProductBundleListingResponse used by Catalog
    type GetProductBundleListingResponse struct {

        
            Items []GetProductBundleCreateResponse  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // Size used by Catalog
    type Size struct {

        
            IsAvailable bool  `json:"is_available"`
            Display string  `json:"display"`
            Value string  `json:"value"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // LimitedProductData used by Catalog
    type LimitedProductData struct {

        
            Sizes []string  `json:"sizes"`
            Identifier map[string]interface{}  `json:"identifier"`
            Slug string  `json:"slug"`
            Quantity float64  `json:"quantity"`
            Images []string  `json:"images"`
            ItemCode string  `json:"item_code"`
            Price map[string]interface{}  `json:"price"`
            ShortDescription string  `json:"short_description"`
            Attributes map[string]interface{}  `json:"attributes"`
            CountryOfOrigin string  `json:"country_of_origin"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // Price used by Catalog
    type Price struct {

        
            Currency string  `json:"currency"`
            MinEffective float64  `json:"min_effective"`
            MaxMarked float64  `json:"max_marked"`
            MaxEffective float64  `json:"max_effective"`
            MinMarked float64  `json:"min_marked"`
         
    }
    
    // GetProducts used by Catalog
    type GetProducts struct {

        
            Sizes []Size  `json:"sizes"`
            ProductUID float64  `json:"product_uid"`
            AutoAddToCart bool  `json:"auto_add_to_cart"`
            ProductDetails LimitedProductData  `json:"product_details"`
            MaxQuantity float64  `json:"max_quantity"`
            Price Price  `json:"price"`
            AutoSelect bool  `json:"auto_select"`
            AllowRemove bool  `json:"allow_remove"`
            MinQuantity float64  `json:"min_quantity"`
         
    }
    
    // GetProductBundleResponse used by Catalog
    type GetProductBundleResponse struct {

        
            CompanyID float64  `json:"company_id"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            Slug string  `json:"slug"`
            Logo string  `json:"logo"`
            IsActive bool  `json:"is_active"`
            Products []GetProducts  `json:"products"`
            Choice string  `json:"choice"`
            Meta map[string]interface{}  `json:"meta"`
            PageVisibility []string  `json:"page_visibility"`
            Name string  `json:"name"`
         
    }
    
    // ProductBundleUpdateRequest used by Catalog
    type ProductBundleUpdateRequest struct {

        
            CompanyID float64  `json:"company_id"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            Slug string  `json:"slug"`
            Logo string  `json:"logo"`
            IsActive bool  `json:"is_active"`
            Choice string  `json:"choice"`
            Products []ProductBundleItem  `json:"products"`
            ModifiedOn string  `json:"modified_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Meta map[string]interface{}  `json:"meta"`
            PageVisibility []string  `json:"page_visibility"`
            Name string  `json:"name"`
         
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
            Description string  `json:"description"`
            CompanyID float64  `json:"company_id"`
            Guide Guide  `json:"guide"`
            Name string  `json:"name"`
            CreatedOn string  `json:"created_on"`
            Subtitle string  `json:"subtitle"`
            ModifiedOn string  `json:"modified_on"`
            Title string  `json:"title"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Active bool  `json:"active"`
            Image string  `json:"image"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            BrandID float64  `json:"brand_id"`
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

        
            ID string  `json:"id"`
            CompanyID float64  `json:"company_id"`
            Guide map[string]interface{}  `json:"guide"`
            CreatedOn string  `json:"created_on"`
            Subtitle string  `json:"subtitle"`
            Title string  `json:"title"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Active bool  `json:"active"`
            ModifiedOn string  `json:"modified_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            BrandID float64  `json:"brand_id"`
            Tag string  `json:"tag"`
            Name string  `json:"name"`
         
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

        
            Moq MOQData  `json:"moq"`
            AltText map[string]interface{}  `json:"alt_text"`
            Seo SEOData  `json:"seo"`
            IsCod bool  `json:"is_cod"`
            IsGift bool  `json:"is_gift"`
         
    }
    
    // ApplicationItemMOQ used by Catalog
    type ApplicationItemMOQ struct {

        
            IncrementUnit float64  `json:"increment_unit"`
            Maximum float64  `json:"maximum"`
            Minimum float64  `json:"minimum"`
         
    }
    
    // MetaFields used by Catalog
    type MetaFields struct {

        
            Value interface{}  `json:"value"`
            Key interface{}  `json:"key"`
         
    }
    
    // ApplicationItemSEO used by Catalog
    type ApplicationItemSEO struct {

        
            Description interface{}  `json:"description"`
            Title interface{}  `json:"title"`
         
    }
    
    // ApplicationItemMeta used by Catalog
    type ApplicationItemMeta struct {

        
            IsGift bool  `json:"is_gift"`
            Moq ApplicationItemMOQ  `json:"moq"`
            CustomMeta []MetaFields  `json:"_custom_meta"`
            Seo ApplicationItemSEO  `json:"seo"`
            AltText map[string]interface{}  `json:"alt_text"`
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

        
            Data []map[string]interface{}  `json:"data"`
            Condition []map[string]interface{}  `json:"condition"`
            Values []map[string]interface{}  `json:"values"`
         
    }
    
    // AttributeDetailsGroup used by Catalog
    type AttributeDetailsGroup struct {

        
            Priority float64  `json:"priority"`
            Unit string  `json:"unit"`
            Key string  `json:"key"`
            DisplayType string  `json:"display_type"`
            Slug string  `json:"slug"`
            Logo string  `json:"logo"`
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
         
    }
    
    // AppConfigurationDetail used by Catalog
    type AppConfigurationDetail struct {

        
            Priority float64  `json:"priority"`
            AppID string  `json:"app_id"`
            IsDefault bool  `json:"is_default"`
            Slug string  `json:"slug"`
            IsActive bool  `json:"is_active"`
            Logo string  `json:"logo"`
            Attributes []AttributeDetailsGroup  `json:"attributes"`
            TemplateSlugs []string  `json:"template_slugs"`
            Name string  `json:"name"`
         
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

        
            Priority float64  `json:"priority"`
            AppID string  `json:"app_id"`
            Key string  `json:"key"`
            IsDefault bool  `json:"is_default"`
            DefaultKey string  `json:"default_key"`
            Logo string  `json:"logo"`
            IsActive bool  `json:"is_active"`
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
    
    // MetaDataListingFilterMetaResponse used by Catalog
    type MetaDataListingFilterMetaResponse struct {

        
            FilterTypes []string  `json:"filter_types"`
            Display string  `json:"display"`
            Key string  `json:"key"`
            Units []map[string]interface{}  `json:"units"`
         
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

        
            Variant map[string]interface{}  `json:"variant"`
            Detail map[string]interface{}  `json:"detail"`
            Similar map[string]interface{}  `json:"similar"`
            Compare map[string]interface{}  `json:"compare"`
         
    }
    
    // GetCatalogConfigurationMetaData used by Catalog
    type GetCatalogConfigurationMetaData struct {

        
            Listing MetaDataListingResponse  `json:"listing"`
            Product GetCatalogConfigurationDetailsProduct  `json:"product"`
         
    }
    
    // ConfigurationBucketPoints used by Catalog
    type ConfigurationBucketPoints struct {

        
            Start float64  `json:"start"`
            Display string  `json:"display"`
            End float64  `json:"end"`
         
    }
    
    // ConfigurationListingFilterValue used by Catalog
    type ConfigurationListingFilterValue struct {

        
            Map map[string]interface{}  `json:"map"`
            Value string  `json:"value"`
            Condition string  `json:"condition"`
            BucketPoints []ConfigurationBucketPoints  `json:"bucket_points"`
            Sort string  `json:"sort"`
            MapValues []map[string]interface{}  `json:"map_values"`
         
    }
    
    // ConfigurationListingFilterConfig used by Catalog
    type ConfigurationListingFilterConfig struct {

        
            Priority float64  `json:"priority"`
            Type string  `json:"type"`
            DisplayName string  `json:"display_name"`
            Key string  `json:"key"`
            Logo string  `json:"logo"`
            IsActive bool  `json:"is_active"`
            ValueConfig ConfigurationListingFilterValue  `json:"value_config"`
            Name string  `json:"name"`
         
    }
    
    // ConfigurationListingFilter used by Catalog
    type ConfigurationListingFilter struct {

        
            AllowSingle bool  `json:"allow_single"`
            AttributeConfig []ConfigurationListingFilterConfig  `json:"attribute_config"`
         
    }
    
    // ConfigurationListingSortConfig used by Catalog
    type ConfigurationListingSortConfig struct {

        
            Priority float64  `json:"priority"`
            Key string  `json:"key"`
            Logo string  `json:"logo"`
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
         
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

        
            Priority float64  `json:"priority"`
            Key string  `json:"key"`
            DisplayType string  `json:"display_type"`
            Logo string  `json:"logo"`
            IsActive bool  `json:"is_active"`
            Size ProductSize  `json:"size"`
            Name string  `json:"name"`
         
    }
    
    // ConfigurationProductVariant used by Catalog
    type ConfigurationProductVariant struct {

        
            Config []ConfigurationProductVariantConfig  `json:"config"`
         
    }
    
    // ConfigurationProductConfig used by Catalog
    type ConfigurationProductConfig struct {

        
            Priority float64  `json:"priority"`
            Key string  `json:"key"`
            Subtitle string  `json:"subtitle"`
            Title string  `json:"title"`
            Logo string  `json:"logo"`
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
    
    // AppConfiguration used by Catalog
    type AppConfiguration struct {

        
            Type string  `json:"type"`
            AppID string  `json:"app_id"`
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
            ConfigID string  `json:"config_id"`
            Listing ConfigurationListing  `json:"listing"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Product ConfigurationProduct  `json:"product"`
            ConfigType string  `json:"config_type"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
         
    }
    
    // AppCatalogConfiguration used by Catalog
    type AppCatalogConfiguration struct {

        
            Type string  `json:"type"`
            ID string  `json:"id"`
            AppID string  `json:"app_id"`
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
            ConfigID string  `json:"config_id"`
            Listing ConfigurationListing  `json:"listing"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Product ConfigurationProduct  `json:"product"`
            ConfigType string  `json:"config_type"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
         
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

        
            ID string  `json:"id"`
            AppID string  `json:"app_id"`
            ConfigID string  `json:"config_id"`
            Listing GetCatalogConfigurationDetailsSchemaListing  `json:"listing"`
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

        
            Operators []string  `json:"operators"`
            Logo string  `json:"logo"`
            Kind string  `json:"kind"`
            Display string  `json:"display"`
            Name string  `json:"name"`
         
    }
    
    // ProductFiltersValue used by Catalog
    type ProductFiltersValue struct {

        
            Min float64  `json:"min"`
            CurrencySymbol string  `json:"currency_symbol"`
            Value interface{}  `json:"value"`
            SelectedMax float64  `json:"selected_max"`
            IsSelected bool  `json:"is_selected"`
            DisplayFormat string  `json:"display_format"`
            Count float64  `json:"count"`
            Max float64  `json:"max"`
            CurrencyCode string  `json:"currency_code"`
            Display string  `json:"display"`
            SelectedMin float64  `json:"selected_min"`
            QueryFormat string  `json:"query_format"`
         
    }
    
    // ProductFilters used by Catalog
    type ProductFilters struct {

        
            Key ProductFiltersKey  `json:"key"`
            Values []ProductFiltersValue  `json:"values"`
         
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
            SortOn []ProductSortOn  `json:"sort_on"`
            Operators map[string]string  `json:"operators"`
         
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
            Start string  `json:"start"`
            End string  `json:"end"`
         
    }
    
    // CollectionQuery used by Catalog
    type CollectionQuery struct {

        
            Value []interface{}  `json:"value"`
            Op string  `json:"op"`
            Attribute string  `json:"attribute"`
         
    }
    
    // CollectionImage used by Catalog
    type CollectionImage struct {

        
            AspectRatio string  `json:"aspect_ratio"`
            URL string  `json:"url"`
         
    }
    
    // CollectionBanner used by Catalog
    type CollectionBanner struct {

        
            Portrait CollectionImage  `json:"portrait"`
            Landscape CollectionImage  `json:"landscape"`
         
    }
    
    // UserInfo used by Catalog
    type UserInfo struct {

        
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
            UID string  `json:"uid"`
            Email string  `json:"email"`
         
    }
    
    // SeoDetail used by Catalog
    type SeoDetail struct {

        
            Description string  `json:"description"`
            Title string  `json:"title"`
         
    }
    
    // CreateCollection used by Catalog
    type CreateCollection struct {

        
            IsVisible bool  `json:"is_visible"`
            Badge CollectionBadge  `json:"badge"`
            Schedule CollectionSchedule  `json:"_schedule"`
            Query []CollectionQuery  `json:"query"`
            AllowFacets bool  `json:"allow_facets"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Published bool  `json:"published"`
            Priority float64  `json:"priority"`
            Slug string  `json:"slug"`
            SortOn string  `json:"sort_on"`
            Banners CollectionBanner  `json:"banners"`
            IsActive bool  `json:"is_active"`
            CreatedBy UserInfo  `json:"created_by"`
            Seo SeoDetail  `json:"seo"`
            Type string  `json:"type"`
            AppID string  `json:"app_id"`
            Logo CollectionImage  `json:"logo"`
            ModifiedBy UserInfo  `json:"modified_by"`
            Meta map[string]interface{}  `json:"meta"`
            Tags []string  `json:"tags"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Description string  `json:"description"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            AllowSort bool  `json:"allow_sort"`
            Name string  `json:"name"`
         
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

        
            Badge map[string]interface{}  `json:"badge"`
            Schedule map[string]interface{}  `json:"_schedule"`
            Query []CollectionQuery  `json:"query"`
            AllowFacets bool  `json:"allow_facets"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Cron map[string]interface{}  `json:"cron"`
            Priority float64  `json:"priority"`
            Slug string  `json:"slug"`
            SortOn string  `json:"sort_on"`
            Banners ImageUrls  `json:"banners"`
            IsActive bool  `json:"is_active"`
            Type string  `json:"type"`
            AppID string  `json:"app_id"`
            Logo BannerImage  `json:"logo"`
            Meta map[string]interface{}  `json:"meta"`
            Tag []string  `json:"tag"`
            Description string  `json:"description"`
            AllowSort bool  `json:"allow_sort"`
            Name string  `json:"name"`
         
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
    
    // Media1 used by Catalog
    type Media1 struct {

        
            Type string  `json:"type"`
            Meta map[string]interface{}  `json:"meta"`
            URL string  `json:"url"`
         
    }
    
    // GetCollectionDetailNest used by Catalog
    type GetCollectionDetailNest struct {

        
            Badge map[string]interface{}  `json:"badge"`
            Schedule map[string]interface{}  `json:"_schedule"`
            Query []CollectionQuery  `json:"query"`
            AllowFacets bool  `json:"allow_facets"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Cron map[string]interface{}  `json:"cron"`
            Priority float64  `json:"priority"`
            Slug string  `json:"slug"`
            Action Action  `json:"action"`
            Banners ImageUrls  `json:"banners"`
            IsActive bool  `json:"is_active"`
            Type string  `json:"type"`
            AppID string  `json:"app_id"`
            Logo Media1  `json:"logo"`
            Meta map[string]interface{}  `json:"meta"`
            Tag []string  `json:"tag"`
            Description string  `json:"description"`
            AllowSort bool  `json:"allow_sort"`
            UID string  `json:"uid"`
            Name string  `json:"name"`
         
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
    
    // GetCollectionListingResponse used by Catalog
    type GetCollectionListingResponse struct {

        
            Items []GetCollectionDetailNest  `json:"items"`
            Filters CollectionListingFilter  `json:"filters"`
            Page Page  `json:"page"`
         
    }
    
    // CollectionDetailResponse used by Catalog
    type CollectionDetailResponse struct {

        
            Badge map[string]interface{}  `json:"badge"`
            Schedule map[string]interface{}  `json:"_schedule"`
            Query []CollectionQuery  `json:"query"`
            AllowFacets bool  `json:"allow_facets"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Cron map[string]interface{}  `json:"cron"`
            Priority float64  `json:"priority"`
            Slug string  `json:"slug"`
            Banners ImageUrls  `json:"banners"`
            IsActive bool  `json:"is_active"`
            Type string  `json:"type"`
            AppID string  `json:"app_id"`
            Logo Media1  `json:"logo"`
            Meta map[string]interface{}  `json:"meta"`
            Tag []string  `json:"tag"`
            Description string  `json:"description"`
            AllowSort bool  `json:"allow_sort"`
            UID string  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // UpdateCollection used by Catalog
    type UpdateCollection struct {

        
            IsVisible bool  `json:"is_visible"`
            Badge CollectionBadge  `json:"badge"`
            Schedule CollectionSchedule  `json:"_schedule"`
            Query []CollectionQuery  `json:"query"`
            AllowFacets bool  `json:"allow_facets"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Published bool  `json:"published"`
            Priority float64  `json:"priority"`
            Slug string  `json:"slug"`
            SortOn string  `json:"sort_on"`
            Banners CollectionBanner  `json:"banners"`
            IsActive bool  `json:"is_active"`
            Seo SeoDetail  `json:"seo"`
            Type string  `json:"type"`
            Logo CollectionImage  `json:"logo"`
            ModifiedBy UserInfo  `json:"modified_by"`
            Meta map[string]interface{}  `json:"meta"`
            Tags []string  `json:"tags"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Description string  `json:"description"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            AllowSort bool  `json:"allow_sort"`
            Name string  `json:"name"`
         
    }
    
    // CollectionItem used by Catalog
    type CollectionItem struct {

        
            Priority float64  `json:"priority"`
            Action string  `json:"action"`
            ItemID float64  `json:"item_id"`
         
    }
    
    // CollectionItemUpdate used by Catalog
    type CollectionItemUpdate struct {

        
            Type string  `json:"type"`
            Query []CollectionQuery  `json:"query"`
            AllowFacets bool  `json:"allow_facets"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            AllowSort bool  `json:"allow_sort"`
            Items []CollectionItem  `json:"items"`
         
    }
    
    // UpdatedResponse used by Catalog
    type UpdatedResponse struct {

        
            ItemsNotUpdated []float64  `json:"items_not_updated"`
            Message string  `json:"message"`
         
    }
    
    // ProductBrand used by Catalog
    type ProductBrand struct {

        
            Action Action  `json:"action"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            Logo Media1  `json:"logo"`
         
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

        
            Highlights []string  `json:"highlights"`
            ItemCode string  `json:"item_code"`
            Color string  `json:"color"`
            Brand ProductBrand  `json:"brand"`
            Rating float64  `json:"rating"`
            Sellable bool  `json:"sellable"`
            Slug string  `json:"slug"`
            ImageNature string  `json:"image_nature"`
            Price ProductListingPrice  `json:"price"`
            Type string  `json:"type"`
            Discount string  `json:"discount"`
            PromoMeta map[string]interface{}  `json:"promo_meta"`
            HasVariant bool  `json:"has_variant"`
            Tryouts []string  `json:"tryouts"`
            Description string  `json:"description"`
            Medias []Media1  `json:"medias"`
            RatingCount float64  `json:"rating_count"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            ItemType string  `json:"item_type"`
            ProductOnlineDate string  `json:"product_online_date"`
            Similars []string  `json:"similars"`
            ShortDescription string  `json:"short_description"`
            Attributes map[string]interface{}  `json:"attributes"`
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // GetCollectionItemsResponse used by Catalog
    type GetCollectionItemsResponse struct {

        
            Items []ProductListingDetail  `json:"items"`
            Filters []ProductFilters  `json:"filters"`
            SortOn []ProductSortOn  `json:"sort_on"`
            Page Page  `json:"page"`
         
    }
    
    // CatalogInsightBrand used by Catalog
    type CatalogInsightBrand struct {

        
            AvailableArticles float64  `json:"available_articles"`
            ArticleFreshness float64  `json:"article_freshness"`
            TotalArticles float64  `json:"total_articles"`
            TotalSizes float64  `json:"total_sizes"`
            AvailableSizes float64  `json:"available_sizes"`
            Name string  `json:"name"`
         
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

        
            Data CrossSellingData  `json:"data"`
            BrandDistribution CatalogInsightBrand  `json:"brand_distribution"`
         
    }
    
    // OptInPostRequest used by Catalog
    type OptInPostRequest struct {

        
            Enabled bool  `json:"enabled"`
            StoreIds []float64  `json:"store_ids"`
            CompanyID float64  `json:"company_id"`
            Platform string  `json:"platform"`
            OptLevel string  `json:"opt_level"`
            BrandIds []float64  `json:"brand_ids"`
         
    }
    
    // CompanyOptIn used by Catalog
    type CompanyOptIn struct {

        
            CompanyID float64  `json:"company_id"`
            StoreIds []float64  `json:"store_ids"`
            Enabled bool  `json:"enabled"`
            Platform string  `json:"platform"`
            CreatedOn float64  `json:"created_on"`
            OptLevel string  `json:"opt_level"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ModifiedOn float64  `json:"modified_on"`
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
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            CompanyType string  `json:"company_type"`
         
    }
    
    // CompanyBrandDetail used by Catalog
    type CompanyBrandDetail struct {

        
            BrandID float64  `json:"brand_id"`
            CompanyID float64  `json:"company_id"`
            BrandName string  `json:"brand_name"`
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

        
            CompanyID float64  `json:"company_id"`
            Documents []map[string]interface{}  `json:"documents"`
            DisplayName string  `json:"display_name"`
            StoreCode string  `json:"store_code"`
            CreatedOn string  `json:"created_on"`
            AdditionalContacts []map[string]interface{}  `json:"additional_contacts"`
            Manager map[string]interface{}  `json:"manager"`
            ModifiedOn string  `json:"modified_on"`
            Address map[string]interface{}  `json:"address"`
            Timing map[string]interface{}  `json:"timing"`
            StoreType string  `json:"store_type"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
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

        
            Type string  `json:"type"`
            Range AttributeSchemaRange  `json:"range"`
            AllowedValues []string  `json:"allowed_values"`
            Mandatory bool  `json:"mandatory"`
            Format string  `json:"format"`
            Multi bool  `json:"multi"`
         
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

        
            IsNested bool  `json:"is_nested"`
            Description string  `json:"description"`
            Schema AttributeMaster  `json:"schema"`
            ID string  `json:"id"`
            EnabledForEndConsumer bool  `json:"enabled_for_end_consumer"`
            Slug string  `json:"slug"`
            Details AttributeMasterDetails  `json:"details"`
            Logo string  `json:"logo"`
            Filters AttributeMasterFilter  `json:"filters"`
            Departments []string  `json:"departments"`
            Meta AttributeMasterMeta  `json:"meta"`
            Name string  `json:"name"`
         
    }
    
    // CategoriesResponse used by Catalog
    type CategoriesResponse struct {

        
            Slug string  `json:"slug"`
            SlugKey string  `json:"slug_key"`
            TemplateSlug string  `json:"template_slug"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // ProdcutTemplateCategoriesResponse used by Catalog
    type ProdcutTemplateCategoriesResponse struct {

        
            Items []CategoriesResponse  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // PTErrorResponse used by Catalog
    type PTErrorResponse struct {

        
            Message string  `json:"message"`
            Code string  `json:"code"`
            Errors map[string]interface{}  `json:"errors"`
            Status float64  `json:"status"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // DepartmentCreateUpdate used by Catalog
    type DepartmentCreateUpdate struct {

        
            Synonyms []string  `json:"synonyms"`
            Slug string  `json:"slug"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            PriorityOrder float64  `json:"priority_order"`
            IsActive bool  `json:"is_active"`
            Logo string  `json:"logo"`
            Platforms map[string]interface{}  `json:"platforms"`
            Cls string  `json:"_cls"`
            Tags []string  `json:"tags"`
            UID float64  `json:"uid"`
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
    
    // UserSerializer used by Catalog
    type UserSerializer struct {

        
            Username string  `json:"username"`
            Contact string  `json:"contact"`
            ID string  `json:"_id"`
            UserID string  `json:"user_id"`
            UID string  `json:"uid"`
         
    }
    
    // GetDepartment used by Catalog
    type GetDepartment struct {

        
            CreatedOn string  `json:"created_on"`
            Synonyms []string  `json:"synonyms"`
            ModifiedOn string  `json:"modified_on"`
            Slug string  `json:"slug"`
            PageSize float64  `json:"page_size"`
            ItemType string  `json:"item_type"`
            IsActive bool  `json:"is_active"`
            Logo string  `json:"logo"`
            PriorityOrder float64  `json:"priority_order"`
            Search string  `json:"search"`
            CreatedBy UserSerializer  `json:"created_by"`
            PageNo float64  `json:"page_no"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // DepartmentsResponse used by Catalog
    type DepartmentsResponse struct {

        
            Items []GetDepartment  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // DepartmentErrorResponse used by Catalog
    type DepartmentErrorResponse struct {

        
            Message string  `json:"message"`
            Code string  `json:"code"`
            Errors map[string]interface{}  `json:"errors"`
            Status float64  `json:"status"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // UserDetail used by Catalog
    type UserDetail struct {

        
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
            Contact string  `json:"contact"`
            SuperUser bool  `json:"super_user"`
         
    }
    
    // DepartmentModel used by Catalog
    type DepartmentModel struct {

        
            VerifiedBy UserDetail  `json:"verified_by"`
            CreatedOn string  `json:"created_on"`
            Synonyms []interface{}  `json:"synonyms"`
            Slug interface{}  `json:"slug"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Logo string  `json:"logo"`
            IsActive bool  `json:"is_active"`
            UID float64  `json:"uid"`
            PriorityOrder float64  `json:"priority_order"`
            CreatedBy UserDetail  `json:"created_by"`
            ModifiedBy UserDetail  `json:"modified_by"`
            Cls interface{}  `json:"_cls"`
            ModifiedOn string  `json:"modified_on"`
            ID interface{}  `json:"_id"`
            VerifiedOn string  `json:"verified_on"`
            Name interface{}  `json:"name"`
         
    }
    
    // ProductTemplate used by Catalog
    type ProductTemplate struct {

        
            Description string  `json:"description"`
            IsArchived bool  `json:"is_archived"`
            IsExpirable bool  `json:"is_expirable"`
            IsPhysical bool  `json:"is_physical"`
            Categories []string  `json:"categories"`
            Slug string  `json:"slug"`
            CreatedOn string  `json:"created_on"`
            IsActive bool  `json:"is_active"`
            Logo string  `json:"logo"`
            Departments []string  `json:"departments"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Attributes []string  `json:"attributes"`
            ModifiedOn string  `json:"modified_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Tag string  `json:"tag"`
            Name string  `json:"name"`
         
    }
    
    // TemplatesResponse used by Catalog
    type TemplatesResponse struct {

        
            Items ProductTemplate  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // Properties used by Catalog
    type Properties struct {

        
            Sizes map[string]interface{}  `json:"sizes"`
            Currency map[string]interface{}  `json:"currency"`
            Highlights map[string]interface{}  `json:"highlights"`
            TraderType map[string]interface{}  `json:"trader_type"`
            ProductPublish map[string]interface{}  `json:"product_publish"`
            ItemCode map[string]interface{}  `json:"item_code"`
            Command map[string]interface{}  `json:"command"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            ProductGroupTag map[string]interface{}  `json:"product_group_tag"`
            Media map[string]interface{}  `json:"media"`
            Slug map[string]interface{}  `json:"slug"`
            IsActive map[string]interface{}  `json:"is_active"`
            IsDependent map[string]interface{}  `json:"is_dependent"`
            MultiSize map[string]interface{}  `json:"multi_size"`
            ReturnConfig map[string]interface{}  `json:"return_config"`
            HsnCode map[string]interface{}  `json:"hsn_code"`
            CountryOfOrigin map[string]interface{}  `json:"country_of_origin"`
            Variants map[string]interface{}  `json:"variants"`
            SizeGuide map[string]interface{}  `json:"size_guide"`
            Trader map[string]interface{}  `json:"trader"`
            Tags map[string]interface{}  `json:"tags"`
            Description map[string]interface{}  `json:"description"`
            CategorySlug map[string]interface{}  `json:"category_slug"`
            ItemType map[string]interface{}  `json:"item_type"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            BrandUID map[string]interface{}  `json:"brand_uid"`
            ShortDescription map[string]interface{}  `json:"short_description"`
            NoOfBoxes map[string]interface{}  `json:"no_of_boxes"`
            Name map[string]interface{}  `json:"name"`
         
    }
    
    // GlobalValidation used by Catalog
    type GlobalValidation struct {

        
            Type string  `json:"type"`
            Description string  `json:"description"`
            Properties Properties  `json:"properties"`
            Required []string  `json:"required"`
            Title string  `json:"title"`
            Definitions map[string]interface{}  `json:"definitions"`
         
    }
    
    // TemplateValidationData used by Catalog
    type TemplateValidationData struct {

        
            GlobalValidation GlobalValidation  `json:"global_validation"`
            TemplateValidation map[string]interface{}  `json:"template_validation"`
         
    }
    
    // TemplateDetails used by Catalog
    type TemplateDetails struct {

        
            ID string  `json:"id"`
            Description string  `json:"description"`
            IsArchived bool  `json:"is_archived"`
            IsExpirable bool  `json:"is_expirable"`
            IsPhysical bool  `json:"is_physical"`
            Categories []string  `json:"categories"`
            Slug string  `json:"slug"`
            IsActive bool  `json:"is_active"`
            Logo string  `json:"logo"`
            Departments []string  `json:"departments"`
            Attributes []string  `json:"attributes"`
            Tag string  `json:"tag"`
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

        
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
            UID string  `json:"uid"`
            Email string  `json:"email"`
         
    }
    
    // ProductTemplateExportResponse used by Catalog
    type ProductTemplateExportResponse struct {

        
            Type string  `json:"type"`
            CompletedOn string  `json:"completed_on"`
            NotificationEmails []string  `json:"notification_emails"`
            ModifiedOn string  `json:"modified_on"`
            Filters map[string]interface{}  `json:"filters"`
            CreatedBy UserInfo1  `json:"created_by"`
            SellerID float64  `json:"seller_id"`
            URL string  `json:"url"`
            Status string  `json:"status"`
            TaskID string  `json:"task_id"`
         
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
            NotificationEmails []string  `json:"notification_emails"`
            Filters ProductTemplateExportFilterRequest  `json:"filters"`
         
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

        
            L2 float64  `json:"l2"`
            L1 float64  `json:"l1"`
            Department float64  `json:"department"`
         
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
    
    // CategoryRequestBody used by Catalog
    type CategoryRequestBody struct {

        
            Priority float64  `json:"priority"`
            Media Media2  `json:"media"`
            Synonyms []string  `json:"synonyms"`
            Slug string  `json:"slug"`
            Level float64  `json:"level"`
            IsActive bool  `json:"is_active"`
            Departments []float64  `json:"departments"`
            Hierarchy []Hierarchy  `json:"hierarchy"`
            Tryouts []string  `json:"tryouts"`
            Marketplaces CategoryMapping  `json:"marketplaces"`
            Name string  `json:"name"`
         
    }
    
    // CategoryCreateResponse used by Catalog
    type CategoryCreateResponse struct {

        
            UID float64  `json:"uid"`
            Message string  `json:"message"`
         
    }
    
    // Category used by Catalog
    type Category struct {

        
            Priority float64  `json:"priority"`
            Media Media2  `json:"media"`
            ID string  `json:"id"`
            CreatedOn string  `json:"created_on"`
            Synonyms []string  `json:"synonyms"`
            ModifiedOn string  `json:"modified_on"`
            Slug string  `json:"slug"`
            Level float64  `json:"level"`
            IsActive bool  `json:"is_active"`
            Departments []float64  `json:"departments"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Hierarchy []Hierarchy  `json:"hierarchy"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Tryouts []string  `json:"tryouts"`
            Marketplaces CategoryMapping  `json:"marketplaces"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
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
    
    // CustomOrder used by Catalog
    type CustomOrder struct {

        
            ManufacturingTime float64  `json:"manufacturing_time"`
            IsCustomOrder bool  `json:"is_custom_order"`
            ManufacturingTimeUnit string  `json:"manufacturing_time_unit"`
         
    }
    
    // ReturnConfig used by Catalog
    type ReturnConfig struct {

        
            Unit string  `json:"unit"`
            Time float64  `json:"time"`
            Returnable bool  `json:"returnable"`
         
    }
    
    // TaxIdentifier used by Catalog
    type TaxIdentifier struct {

        
            ReportingHsn string  `json:"reporting_hsn"`
            HsnCodeID string  `json:"hsn_code_id"`
            HsnCode string  `json:"hsn_code"`
         
    }
    
    // Trader used by Catalog
    type Trader struct {

        
            Address []string  `json:"address"`
            Type string  `json:"type"`
            Name interface{}  `json:"name"`
         
    }
    
    // TeaserTag used by Catalog
    type TeaserTag struct {

        
            URL string  `json:"url"`
            Tag string  `json:"tag"`
         
    }
    
    // ProductCreateUpdateSchemaV2 used by Catalog
    type ProductCreateUpdateSchemaV2 struct {

        
            Sizes []map[string]interface{}  `json:"sizes"`
            Currency string  `json:"currency"`
            Highlights []string  `json:"highlights"`
            IsImageLessProduct bool  `json:"is_image_less_product"`
            ProductPublish ProductPublish  `json:"product_publish"`
            ItemCode string  `json:"item_code"`
            ChangeRequestID interface{}  `json:"change_request_id"`
            NetQuantity NetQuantity  `json:"net_quantity"`
            CustomOrder CustomOrder  `json:"custom_order"`
            IsSet bool  `json:"is_set"`
            Requester string  `json:"requester"`
            ProductGroupTag []string  `json:"product_group_tag"`
            VariantMedia map[string]interface{}  `json:"variant_media"`
            Media []Media1  `json:"media"`
            Slug string  `json:"slug"`
            Action string  `json:"action"`
            TemplateTag string  `json:"template_tag"`
            IsActive bool  `json:"is_active"`
            IsDependent bool  `json:"is_dependent"`
            MultiSize bool  `json:"multi_size"`
            ReturnConfig ReturnConfig  `json:"return_config"`
            CountryOfOrigin string  `json:"country_of_origin"`
            Variants map[string]interface{}  `json:"variants"`
            TaxIdentifier TaxIdentifier  `json:"tax_identifier"`
            SizeGuide string  `json:"size_guide"`
            Trader []Trader  `json:"trader"`
            Tags []string  `json:"tags"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            CompanyID float64  `json:"company_id"`
            Description string  `json:"description"`
            CategorySlug string  `json:"category_slug"`
            ItemType string  `json:"item_type"`
            TeaserTag TeaserTag  `json:"teaser_tag"`
            BrandUID float64  `json:"brand_uid"`
            VariantGroup map[string]interface{}  `json:"variant_group"`
            Departments []float64  `json:"departments"`
            ShortDescription string  `json:"short_description"`
            NoOfBoxes float64  `json:"no_of_boxes"`
            Attributes map[string]interface{}  `json:"attributes"`
            BulkJobID string  `json:"bulk_job_id"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // ProductPublish1 used by Catalog
    type ProductPublish1 struct {

        
            ProductOnlineDate string  `json:"product_online_date"`
            IsSet bool  `json:"is_set"`
         
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

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            Logo Logo  `json:"logo"`
         
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
            Time float64  `json:"time"`
            Returnable bool  `json:"returnable"`
         
    }
    
    // Image used by Catalog
    type Image struct {

        
            SecureURL string  `json:"secure_url"`
            AspectRatioF float64  `json:"aspect_ratio_f"`
            AspectRatio string  `json:"aspect_ratio"`
            URL string  `json:"url"`
         
    }
    
    // ProductSchemaV2 used by Catalog
    type ProductSchemaV2 struct {

        
            Sizes []map[string]interface{}  `json:"sizes"`
            ID string  `json:"id"`
            Category map[string]interface{}  `json:"category"`
            Currency string  `json:"currency"`
            Highlights []string  `json:"highlights"`
            IsExpirable bool  `json:"is_expirable"`
            CreatedOn string  `json:"created_on"`
            IsImageLessProduct bool  `json:"is_image_less_product"`
            ProductPublish ProductPublish1  `json:"product_publish"`
            ItemCode string  `json:"item_code"`
            Color string  `json:"color"`
            Brand Brand  `json:"brand"`
            NetQuantity NetQuantityResponse  `json:"net_quantity"`
            AllCompanyIds []float64  `json:"all_company_ids"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            IsSet bool  `json:"is_set"`
            ProductGroupTag []string  `json:"product_group_tag"`
            VariantMedia map[string]interface{}  `json:"variant_media"`
            Media []Media1  `json:"media"`
            VerifiedBy VerifiedBy  `json:"verified_by"`
            IsPhysical bool  `json:"is_physical"`
            Slug string  `json:"slug"`
            ImageNature string  `json:"image_nature"`
            PrimaryColor string  `json:"primary_color"`
            TemplateTag string  `json:"template_tag"`
            IsActive bool  `json:"is_active"`
            Moq map[string]interface{}  `json:"moq"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Pending string  `json:"pending"`
            IsDependent bool  `json:"is_dependent"`
            AllSizes []map[string]interface{}  `json:"all_sizes"`
            MultiSize bool  `json:"multi_size"`
            ReturnConfig ReturnConfigResponse  `json:"return_config"`
            HsnCode string  `json:"hsn_code"`
            CountryOfOrigin string  `json:"country_of_origin"`
            VerifiedOn string  `json:"verified_on"`
            Variants map[string]interface{}  `json:"variants"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            AllIdentifiers []string  `json:"all_identifiers"`
            Images []Image  `json:"images"`
            CategoryUID float64  `json:"category_uid"`
            Stage string  `json:"stage"`
            L3Mapping []string  `json:"l3_mapping"`
            SizeGuide string  `json:"size_guide"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            ModifiedOn string  `json:"modified_on"`
            Tags []string  `json:"tags"`
            Trader []Trader  `json:"trader"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Description string  `json:"description"`
            CompanyID float64  `json:"company_id"`
            CategorySlug string  `json:"category_slug"`
            ItemType string  `json:"item_type"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            BrandUID float64  `json:"brand_uid"`
            VariantGroup map[string]interface{}  `json:"variant_group"`
            Departments []float64  `json:"departments"`
            ShortDescription string  `json:"short_description"`
            NoOfBoxes float64  `json:"no_of_boxes"`
            Attributes map[string]interface{}  `json:"attributes"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // ProductListingResponseV2 used by Catalog
    type ProductListingResponseV2 struct {

        
            Items []ProductSchemaV2  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // ProductVariants used by Catalog
    type ProductVariants struct {

        
            Media []Media1  `json:"media"`
            BrandUID float64  `json:"brand_uid"`
            ItemCode string  `json:"item_code"`
            CategoryUID float64  `json:"category_uid"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // ProductVariantsResponse used by Catalog
    type ProductVariantsResponse struct {

        
            Page Page  `json:"page"`
            Variants []ProductVariants  `json:"variants"`
         
    }
    
    // AttributeMasterSerializer used by Catalog
    type AttributeMasterSerializer struct {

        
            CreatedOn string  `json:"created_on"`
            Synonyms map[string]interface{}  `json:"synonyms"`
            Variant bool  `json:"variant"`
            Schema AttributeMaster  `json:"schema"`
            Slug string  `json:"slug"`
            Details AttributeMasterDetails  `json:"details"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            RawKey string  `json:"raw_key"`
            Unit string  `json:"unit"`
            EnabledForEndConsumer bool  `json:"enabled_for_end_consumer"`
            Suggestion string  `json:"suggestion"`
            Logo string  `json:"logo"`
            ModifiedOn string  `json:"modified_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Tags []string  `json:"tags"`
            IsNested bool  `json:"is_nested"`
            Description string  `json:"description"`
            Filters AttributeMasterFilter  `json:"filters"`
            Departments []string  `json:"departments"`
            Name string  `json:"name"`
         
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
            Primary bool  `json:"primary"`
            GtinValue string  `json:"gtin_value"`
         
    }
    
    // AllSizes used by Catalog
    type AllSizes struct {

        
            ItemLength float64  `json:"item_length"`
            Identifiers []ValidateIdentifier  `json:"identifiers"`
            ItemHeight float64  `json:"item_height"`
            ItemWidth float64  `json:"item_width"`
            ItemDimensionsUnitOfMeasure string  `json:"item_dimensions_unit_of_measure"`
            ItemWeightUnitOfMeasure interface{}  `json:"item_weight_unit_of_measure"`
            ItemWeight float64  `json:"item_weight"`
            Size interface{}  `json:"size"`
         
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

        
            Sizes []map[string]interface{}  `json:"sizes"`
            ID string  `json:"id"`
            Category map[string]interface{}  `json:"category"`
            Currency string  `json:"currency"`
            Highlights []string  `json:"highlights"`
            IsExpirable bool  `json:"is_expirable"`
            CreatedOn string  `json:"created_on"`
            IsImageLessProduct bool  `json:"is_image_less_product"`
            ProductPublish ProductPublished  `json:"product_publish"`
            ItemCode string  `json:"item_code"`
            Color string  `json:"color"`
            Brand Brand  `json:"brand"`
            NetQuantity NetQuantityResponse  `json:"net_quantity"`
            AllCompanyIds []float64  `json:"all_company_ids"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            IsSet bool  `json:"is_set"`
            ProductGroupTag []string  `json:"product_group_tag"`
            VariantMedia map[string]interface{}  `json:"variant_media"`
            Media []Media1  `json:"media"`
            VerifiedBy VerifiedBy  `json:"verified_by"`
            IsPhysical bool  `json:"is_physical"`
            Slug string  `json:"slug"`
            ImageNature string  `json:"image_nature"`
            PrimaryColor string  `json:"primary_color"`
            TemplateTag string  `json:"template_tag"`
            IsActive bool  `json:"is_active"`
            Moq map[string]interface{}  `json:"moq"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Pending string  `json:"pending"`
            IsDependent bool  `json:"is_dependent"`
            AllSizes []map[string]interface{}  `json:"all_sizes"`
            MultiSize bool  `json:"multi_size"`
            ReturnConfig ReturnConfigResponse  `json:"return_config"`
            HsnCode string  `json:"hsn_code"`
            CountryOfOrigin string  `json:"country_of_origin"`
            VerifiedOn string  `json:"verified_on"`
            Variants map[string]interface{}  `json:"variants"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            AllIdentifiers []string  `json:"all_identifiers"`
            Images []Image  `json:"images"`
            CategoryUID float64  `json:"category_uid"`
            Stage string  `json:"stage"`
            L3Mapping []string  `json:"l3_mapping"`
            SizeGuide string  `json:"size_guide"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            ModifiedOn string  `json:"modified_on"`
            Tags []string  `json:"tags"`
            Trader []Trader  `json:"trader"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Description string  `json:"description"`
            CompanyID float64  `json:"company_id"`
            CategorySlug string  `json:"category_slug"`
            ItemType string  `json:"item_type"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            BrandUID float64  `json:"brand_uid"`
            VariantGroup map[string]interface{}  `json:"variant_group"`
            Departments []float64  `json:"departments"`
            ShortDescription string  `json:"short_description"`
            NoOfBoxes float64  `json:"no_of_boxes"`
            Attributes map[string]interface{}  `json:"attributes"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // ProductListingResponse used by Catalog
    type ProductListingResponse struct {

        
            Items []Product  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // BulkJob used by Catalog
    type BulkJob struct {

        
            Cancelled float64  `json:"cancelled"`
            CompanyID float64  `json:"company_id"`
            Succeed float64  `json:"succeed"`
            CreatedOn string  `json:"created_on"`
            Total float64  `json:"total"`
            CancelledRecords []map[string]interface{}  `json:"cancelled_records"`
            FailedRecords []map[string]interface{}  `json:"failed_records"`
            FilePath string  `json:"file_path"`
            CustomTemplateTag string  `json:"custom_template_tag"`
            TemplateTag string  `json:"template_tag"`
            IsActive bool  `json:"is_active"`
            TrackingURL string  `json:"tracking_url"`
            CreatedBy UserInfo1  `json:"created_by"`
            ModifiedOn string  `json:"modified_on"`
            ModifiedBy UserInfo1  `json:"modified_by"`
            Failed float64  `json:"failed"`
            Stage string  `json:"stage"`
         
    }
    
    // BulkResponse used by Catalog
    type BulkResponse struct {

        
            CreatedOn string  `json:"created_on"`
            IsActive bool  `json:"is_active"`
            BatchID string  `json:"batch_id"`
            CreatedBy UserInfo1  `json:"created_by"`
            ModifiedOn string  `json:"modified_on"`
            ModifiedBy UserInfo1  `json:"modified_by"`
         
    }
    
    // UserDetail1 used by Catalog
    type UserDetail1 struct {

        
            Username string  `json:"username"`
            FullName string  `json:"full_name"`
            UserID string  `json:"user_id"`
         
    }
    
    // ProductBulkRequest used by Catalog
    type ProductBulkRequest struct {

        
            Cancelled float64  `json:"cancelled"`
            Template ProductTemplate  `json:"template"`
            CompanyID float64  `json:"company_id"`
            Succeed float64  `json:"succeed"`
            CreatedOn string  `json:"created_on"`
            CancelledRecords []string  `json:"cancelled_records"`
            Total float64  `json:"total"`
            FailedRecords []string  `json:"failed_records"`
            FilePath string  `json:"file_path"`
            TemplateTag string  `json:"template_tag"`
            IsActive bool  `json:"is_active"`
            CreatedBy UserDetail1  `json:"created_by"`
            ModifiedOn string  `json:"modified_on"`
            ModifiedBy UserDetail1  `json:"modified_by"`
            Failed float64  `json:"failed"`
            Stage string  `json:"stage"`
         
    }
    
    // ProductBulkRequestList used by Catalog
    type ProductBulkRequestList struct {

        
            Items ProductBulkRequest  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // BulkProductRequest used by Catalog
    type BulkProductRequest struct {

        
            Data []map[string]interface{}  `json:"data"`
            BatchID string  `json:"batch_id"`
            CompanyID float64  `json:"company_id"`
            TemplateTag string  `json:"template_tag"`
         
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

        
            Username string  `json:"username"`
            CompanyID float64  `json:"company_id"`
            UserID string  `json:"user_id"`
         
    }
    
    // Items used by Catalog
    type Items struct {

        
            Cancelled float64  `json:"cancelled"`
            ID string  `json:"id"`
            CompanyID float64  `json:"company_id"`
            Succeed float64  `json:"succeed"`
            CreatedOn string  `json:"created_on"`
            FilePath string  `json:"file_path"`
            CancelledRecords []string  `json:"cancelled_records"`
            FailedRecords []string  `json:"failed_records"`
            Total float64  `json:"total"`
            IsActive bool  `json:"is_active"`
            TrackingURL string  `json:"tracking_url"`
            CreatedBy UserCommon  `json:"created_by"`
            Retry float64  `json:"retry"`
            ModifiedOn string  `json:"modified_on"`
            ModifiedBy UserCommon  `json:"modified_by"`
            Failed float64  `json:"failed"`
            Stage string  `json:"stage"`
         
    }
    
    // BulkAssetResponse used by Catalog
    type BulkAssetResponse struct {

        
            Items []Items  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // ProductSizeDeleteDataResponse used by Catalog
    type ProductSizeDeleteDataResponse struct {

        
            CompanyID float64  `json:"company_id"`
            Size string  `json:"size"`
            ItemID float64  `json:"item_id"`
         
    }
    
    // ProductSizeDeleteResponse used by Catalog
    type ProductSizeDeleteResponse struct {

        
            Data ProductSizeDeleteDataResponse  `json:"data"`
            Success bool  `json:"success"`
         
    }
    
    // GTIN used by Catalog
    type GTIN struct {

        
            GtinType string  `json:"gtin_type"`
            Primary bool  `json:"primary"`
            GtinValue interface{}  `json:"gtin_value"`
         
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
    
    // InvSize used by Catalog
    type InvSize struct {

        
            ItemLength float64  `json:"item_length"`
            Currency string  `json:"currency"`
            StoreCode string  `json:"store_code"`
            Quantity float64  `json:"quantity"`
            Identifiers []GTIN  `json:"identifiers"`
            PriceTransfer float64  `json:"price_transfer"`
            Set InventorySet  `json:"set"`
            ItemHeight float64  `json:"item_height"`
            ItemWidth float64  `json:"item_width"`
            Price float64  `json:"price"`
            IsSet bool  `json:"is_set"`
            ExpirationDate string  `json:"expiration_date"`
            ItemDimensionsUnitOfMeasure string  `json:"item_dimensions_unit_of_measure"`
            ItemWeightUnitOfMeasure string  `json:"item_weight_unit_of_measure"`
            PriceEffective float64  `json:"price_effective"`
            ItemWeight float64  `json:"item_weight"`
            Size interface{}  `json:"size"`
         
    }
    
    // ItemQuery used by Catalog
    type ItemQuery struct {

        
            ItemCode string  `json:"item_code"`
            UID float64  `json:"uid"`
            BrandUID float64  `json:"brand_uid"`
         
    }
    
    // InventoryRequest used by Catalog
    type InventoryRequest struct {

        
            Sizes []InvSize  `json:"sizes"`
            CompanyID float64  `json:"company_id"`
            Item ItemQuery  `json:"item"`
         
    }
    
    // InventoryResponse used by Catalog
    type InventoryResponse struct {

        
            Currency string  `json:"currency"`
            Quantity float64  `json:"quantity"`
            Identifiers map[string]interface{}  `json:"identifiers"`
            InventoryUpdatedOn string  `json:"inventory_updated_on"`
            PriceTransfer float64  `json:"price_transfer"`
            SellableQuantity float64  `json:"sellable_quantity"`
            Store map[string]interface{}  `json:"store"`
            UID string  `json:"uid"`
            Price float64  `json:"price"`
            ItemID float64  `json:"item_id"`
            SellerIdentifier string  `json:"seller_identifier"`
            PriceEffective float64  `json:"price_effective"`
            Size string  `json:"size"`
         
    }
    
    // InventoryResponsePaginated used by Catalog
    type InventoryResponsePaginated struct {

        
            Items []InventoryResponse  `json:"items"`
            Page Page  `json:"page"`
         
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
    
    // DimensionResponse used by Catalog
    type DimensionResponse struct {

        
            Unit string  `json:"unit"`
            IsDefault bool  `json:"is_default"`
            Height float64  `json:"height"`
            Length float64  `json:"length"`
            Width float64  `json:"width"`
         
    }
    
    // PriceMeta used by Catalog
    type PriceMeta struct {

        
            Currency string  `json:"currency"`
            Transfer float64  `json:"transfer"`
            Marked float64  `json:"marked"`
            TpNotes map[string]interface{}  `json:"tp_notes"`
            UpdatedAt string  `json:"updated_at"`
            Effective float64  `json:"effective"`
         
    }
    
    // ReturnConfig1 used by Catalog
    type ReturnConfig1 struct {

        
            Unit string  `json:"unit"`
            Time float64  `json:"time"`
            Returnable bool  `json:"returnable"`
         
    }
    
    // StoreMeta used by Catalog
    type StoreMeta struct {

        
            ID float64  `json:"id"`
         
    }
    
    // WeightResponse used by Catalog
    type WeightResponse struct {

        
            Unit string  `json:"unit"`
            IsDefault bool  `json:"is_default"`
            Shipping float64  `json:"shipping"`
         
    }
    
    // ManufacturerResponse used by Catalog
    type ManufacturerResponse struct {

        
            Address string  `json:"address"`
            IsDefault bool  `json:"is_default"`
            Name string  `json:"name"`
         
    }
    
    // Trader1 used by Catalog
    type Trader1 struct {

        
            Address []string  `json:"address"`
            Type string  `json:"type"`
            Name string  `json:"name"`
         
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
    
    // InventorySellerResponse used by Catalog
    type InventorySellerResponse struct {

        
            TotalQuantity float64  `json:"total_quantity"`
            RawMeta map[string]interface{}  `json:"raw_meta"`
            Brand BrandMeta  `json:"brand"`
            Set InventorySet  `json:"set"`
            UID string  `json:"uid"`
            Company CompanyMeta  `json:"company"`
            IsSet bool  `json:"is_set"`
            TraceID string  `json:"trace_id"`
            SellerIdentifier string  `json:"seller_identifier"`
            Identifier map[string]interface{}  `json:"identifier"`
            Dimension DimensionResponse  `json:"dimension"`
            IsActive bool  `json:"is_active"`
            CreatedBy UserSerializer  `json:"created_by"`
            Price PriceMeta  `json:"price"`
            FyndArticleCode string  `json:"fynd_article_code"`
            ItemID float64  `json:"item_id"`
            ReturnConfig ReturnConfig1  `json:"return_config"`
            Store StoreMeta  `json:"store"`
            Weight WeightResponse  `json:"weight"`
            CountryOfOrigin string  `json:"country_of_origin"`
            AddedOnStore string  `json:"added_on_store"`
            Manufacturer ManufacturerResponse  `json:"manufacturer"`
            TrackInventory bool  `json:"track_inventory"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            ExpirationDate string  `json:"expiration_date"`
            Stage string  `json:"stage"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            Meta map[string]interface{}  `json:"meta"`
            Trader []Trader1  `json:"trader"`
            Fragile bool  `json:"fragile"`
            Tags []string  `json:"tags"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Quantities Quantities  `json:"quantities"`
            FyndMeta map[string]interface{}  `json:"fynd_meta"`
            Size string  `json:"size"`
            FyndItemCode string  `json:"fynd_item_code"`
         
    }
    
    // InventorySellerIdentifierResponsePaginated used by Catalog
    type InventorySellerIdentifierResponsePaginated struct {

        
            Items []InventorySellerResponse  `json:"items"`
            Page Page  `json:"page"`
         
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
    
    // DateMeta used by Catalog
    type DateMeta struct {

        
            CreatedOn string  `json:"created_on"`
            AddedOnStore string  `json:"added_on_store"`
            InventoryUpdatedOn string  `json:"inventory_updated_on"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // DimensionResponse1 used by Catalog
    type DimensionResponse1 struct {

        
            Unit string  `json:"unit"`
            Length float64  `json:"length"`
            Height float64  `json:"height"`
            Width float64  `json:"width"`
         
    }
    
    // PriceArticle used by Catalog
    type PriceArticle struct {

        
            Currency string  `json:"currency"`
            Transfer float64  `json:"transfer"`
            Marked float64  `json:"marked"`
            TpNotes map[string]interface{}  `json:"tp_notes"`
            Effective float64  `json:"effective"`
         
    }
    
    // ReturnConfig2 used by Catalog
    type ReturnConfig2 struct {

        
            Unit string  `json:"unit"`
            Time float64  `json:"time"`
            Returnable bool  `json:"returnable"`
         
    }
    
    // ArticleStoreResponse used by Catalog
    type ArticleStoreResponse struct {

        
            StoreCode string  `json:"store_code"`
            StoreType string  `json:"store_type"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // WeightResponse1 used by Catalog
    type WeightResponse1 struct {

        
            Unit string  `json:"unit"`
            Shipping float64  `json:"shipping"`
         
    }
    
    // ManufacturerResponse1 used by Catalog
    type ManufacturerResponse1 struct {

        
            Address string  `json:"address"`
            IsDefault bool  `json:"is_default"`
            Name string  `json:"name"`
         
    }
    
    // Trader2 used by Catalog
    type Trader2 struct {

        
            Address []string  `json:"address"`
            Type string  `json:"type"`
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
    
    // GetInventories used by Catalog
    type GetInventories struct {

        
            TotalQuantity float64  `json:"total_quantity"`
            ID string  `json:"id"`
            Brand BrandMeta1  `json:"brand"`
            UID string  `json:"uid"`
            Company CompanyMeta1  `json:"company"`
            IsSet bool  `json:"is_set"`
            TraceID string  `json:"trace_id"`
            SellerIdentifier string  `json:"seller_identifier"`
            Identifier map[string]interface{}  `json:"identifier"`
            DateMeta DateMeta  `json:"date_meta"`
            Dimension DimensionResponse1  `json:"dimension"`
            CreatedBy UserSerializer  `json:"created_by"`
            Price PriceArticle  `json:"price"`
            ItemID float64  `json:"item_id"`
            ReturnConfig ReturnConfig2  `json:"return_config"`
            Store ArticleStoreResponse  `json:"store"`
            Weight WeightResponse1  `json:"weight"`
            CountryOfOrigin string  `json:"country_of_origin"`
            Manufacturer ManufacturerResponse1  `json:"manufacturer"`
            TrackInventory bool  `json:"track_inventory"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            InventoryUpdatedOn string  `json:"inventory_updated_on"`
            ExpirationDate string  `json:"expiration_date"`
            Stage string  `json:"stage"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            Tags []string  `json:"tags"`
            Trader []Trader2  `json:"trader"`
            Quantities QuantitiesArticle  `json:"quantities"`
            Platforms map[string]interface{}  `json:"platforms"`
            Size string  `json:"size"`
         
    }
    
    // GetInventoriesResponse used by Catalog
    type GetInventoriesResponse struct {

        
            Items []GetInventories  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // BulkInventoryGetItems used by Catalog
    type BulkInventoryGetItems struct {

        
            Cancelled float64  `json:"cancelled"`
            ID string  `json:"id"`
            CompanyID float64  `json:"company_id"`
            Succeed float64  `json:"succeed"`
            CreatedOn string  `json:"created_on"`
            CancelledRecords []string  `json:"cancelled_records"`
            FilePath string  `json:"file_path"`
            FailedRecords []string  `json:"failed_records"`
            Total float64  `json:"total"`
            IsActive bool  `json:"is_active"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ModifiedOn string  `json:"modified_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Failed float64  `json:"failed"`
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
            ItemWeightUnitOfMeasure string  `json:"item_weight_unit_of_measure"`
            PriceMarked float64  `json:"price_marked"`
            Currency string  `json:"currency"`
            StoreCode string  `json:"store_code"`
            Quantity float64  `json:"quantity"`
            Price float64  `json:"price"`
            ExpirationDate string  `json:"expiration_date"`
            TraceID string  `json:"trace_id"`
            ItemDimensionsUnitOfMeasure string  `json:"item_dimensions_unit_of_measure"`
            SellerIdentifier string  `json:"seller_identifier"`
            PriceEffective float64  `json:"price_effective"`
            Tags []string  `json:"tags"`
         
    }
    
    // InventoryBulkRequest used by Catalog
    type InventoryBulkRequest struct {

        
            Sizes []InventoryJobPayload  `json:"sizes"`
            BatchID string  `json:"batch_id"`
            CompanyID float64  `json:"company_id"`
            User map[string]interface{}  `json:"user"`
         
    }
    
    // InventoryExportRequest used by Catalog
    type InventoryExportRequest struct {

        
            Type string  `json:"type"`
            Store []float64  `json:"store"`
            Brand []float64  `json:"brand"`
         
    }
    
    // InventoryExportResponse used by Catalog
    type InventoryExportResponse struct {

        
            Type string  `json:"type"`
            CreatedOn string  `json:"created_on"`
            NotificationEmails []string  `json:"notification_emails"`
            Filters map[string]interface{}  `json:"filters"`
            CreatedBy string  `json:"created_by"`
            SellerID float64  `json:"seller_id"`
            Status string  `json:"status"`
            TaskID string  `json:"task_id"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // InventoryExportQuantityFilter used by Catalog
    type InventoryExportQuantityFilter struct {

        
            Max float64  `json:"max"`
            Min float64  `json:"min"`
            Operators string  `json:"operators"`
         
    }
    
    // InventoryExportAdvanceOption used by Catalog
    type InventoryExportAdvanceOption struct {

        
            StoreIds []float64  `json:"store_ids"`
            Quantity InventoryExportQuantityFilter  `json:"quantity"`
            FromDate string  `json:"from_date"`
            ToDate string  `json:"to_date"`
            BrandIds []float64  `json:"brand_ids"`
         
    }
    
    // InventoryExportJob used by Catalog
    type InventoryExportJob struct {

        
            Type string  `json:"type"`
            CompletedOn string  `json:"completed_on"`
            NotificationEmails []string  `json:"notification_emails"`
            Filters InventoryExportAdvanceOption  `json:"filters"`
            SellerID float64  `json:"seller_id"`
            URL string  `json:"url"`
            Status string  `json:"status"`
            TaskID string  `json:"task_id"`
         
    }
    
    // InventoryExportFilter used by Catalog
    type InventoryExportFilter struct {

        
            StoreIds []float64  `json:"store_ids"`
            Quantity InventoryExportQuantityFilter  `json:"quantity"`
            FromDate string  `json:"from_date"`
            ToDate string  `json:"to_date"`
            BrandIds []float64  `json:"brand_ids"`
         
    }
    
    // InventoryCreateRequest used by Catalog
    type InventoryCreateRequest struct {

        
            Data []string  `json:"data"`
            Type string  `json:"type"`
            NotificationEmails []string  `json:"notification_emails"`
            Filters InventoryExportFilter  `json:"filters"`
         
    }
    
    // InventoryJobFilters used by Catalog
    type InventoryJobFilters struct {

        
            Brands []string  `json:"brands"`
            Stores []string  `json:"stores"`
            Quantity InventoryExportQuantityFilter  `json:"quantity"`
            FromDate string  `json:"from_date"`
            ToDate string  `json:"to_date"`
         
    }
    
    // InventoryJobDetailResponse used by Catalog
    type InventoryJobDetailResponse struct {

        
            CancelledOn string  `json:"cancelled_on"`
            CompletedOn string  `json:"completed_on"`
            Type string  `json:"type"`
            ID string  `json:"id"`
            CreatedOn string  `json:"created_on"`
            NotificationEmails []string  `json:"notification_emails"`
            ModifiedOn string  `json:"modified_on"`
            CancelledBy UserDetail  `json:"cancelled_by"`
            Filters InventoryJobFilters  `json:"filters"`
            CreatedBy UserDetail  `json:"created_by"`
            SellerID float64  `json:"seller_id"`
            URL string  `json:"url"`
            Status interface{}  `json:"status"`
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

        
            TotalQuantity float64  `json:"total_quantity"`
            PriceMarked float64  `json:"price_marked"`
            ExpirationDate string  `json:"expiration_date"`
            TraceID string  `json:"trace_id"`
            SellerIdentifier string  `json:"seller_identifier"`
            StoreID float64  `json:"store_id"`
            PriceEffective float64  `json:"price_effective"`
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
            ID string  `json:"id"`
            CompanyID float64  `json:"company_id"`
            Hs2Code string  `json:"hs2_code"`
            TaxOnEsp bool  `json:"tax_on_esp"`
            ModifiedOn string  `json:"modified_on"`
            Tax1 float64  `json:"tax1"`
            Tax2 float64  `json:"tax2"`
            Threshold1 float64  `json:"threshold1"`
            TaxOnMrp bool  `json:"tax_on_mrp"`
            HsnCode string  `json:"hsn_code"`
         
    }
    
    // HsnCode used by Catalog
    type HsnCode struct {

        
            Data HsnCodesObject  `json:"data"`
         
    }
    
    // HsnUpsert used by Catalog
    type HsnUpsert struct {

        
            Threshold2 float64  `json:"threshold2"`
            CompanyID float64  `json:"company_id"`
            Hs2Code string  `json:"hs2_code"`
            TaxOnEsp bool  `json:"tax_on_esp"`
            Tax1 float64  `json:"tax1"`
            IsActive bool  `json:"is_active"`
            Tax2 float64  `json:"tax2"`
            Threshold1 float64  `json:"threshold1"`
            TaxOnMrp bool  `json:"tax_on_mrp"`
            HsnCode string  `json:"hsn_code"`
            UID float64  `json:"uid"`
         
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

        
            Rate float64  `json:"rate"`
            Cess float64  `json:"cess"`
            Threshold float64  `json:"threshold"`
            EffectiveDate string  `json:"effective_date"`
         
    }
    
    // HSNDataInsertV2 used by Catalog
    type HSNDataInsertV2 struct {

        
            Type string  `json:"type"`
            Description string  `json:"description"`
            ReportingHsn string  `json:"reporting_hsn"`
            CreatedOn string  `json:"created_on"`
            HsnCodeID string  `json:"hsn_code_id"`
            CountryCode string  `json:"country_code"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Taxes []TaxSlab  `json:"taxes"`
            ModifiedOn string  `json:"modified_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            HsnCode string  `json:"hsn_code"`
         
    }
    
    // PageResponse used by Catalog
    type PageResponse struct {

        
            ItemTotal float64  `json:"item_total"`
            HasPrevious bool  `json:"has_previous"`
            Current string  `json:"current"`
            HasNext bool  `json:"has_next"`
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
            Action Action  `json:"action"`
            Slug string  `json:"slug"`
            Banners ImageUrls  `json:"banners"`
            Logo Media  `json:"logo"`
            Departments []string  `json:"departments"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
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
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // DepartmentResponse used by Catalog
    type DepartmentResponse struct {

        
            Items []Department  `json:"items"`
         
    }
    
    // ThirdLevelChild used by Catalog
    type ThirdLevelChild struct {

        
            Childs []map[string]interface{}  `json:"childs"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Action Action  `json:"action"`
            Banners ImageUrls  `json:"banners"`
            UID float64  `json:"uid"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // SecondLevelChild used by Catalog
    type SecondLevelChild struct {

        
            Childs []ThirdLevelChild  `json:"childs"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Action Action  `json:"action"`
            Banners ImageUrls  `json:"banners"`
            UID float64  `json:"uid"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // Child used by Catalog
    type Child struct {

        
            Childs []SecondLevelChild  `json:"childs"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Action Action  `json:"action"`
            Banners ImageUrls  `json:"banners"`
            UID float64  `json:"uid"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // CategoryItems used by Catalog
    type CategoryItems struct {

        
            Childs []Child  `json:"childs"`
            Slug string  `json:"slug"`
            Action Action  `json:"action"`
            Banners ImageUrls  `json:"banners"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
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

        
            SortOn []ProductSortOn  `json:"sort_on"`
            Page Page  `json:"page"`
            Operators map[string]interface{}  `json:"operators"`
            Items []ProductListingDetail  `json:"items"`
            Filters []ProductFilters  `json:"filters"`
         
    }
    
    // ProductDetail used by Catalog
    type ProductDetail struct {

        
            Highlights []string  `json:"highlights"`
            ItemCode string  `json:"item_code"`
            Color string  `json:"color"`
            Brand ProductBrand  `json:"brand"`
            Rating float64  `json:"rating"`
            Slug string  `json:"slug"`
            ImageNature string  `json:"image_nature"`
            Type string  `json:"type"`
            PromoMeta map[string]interface{}  `json:"promo_meta"`
            HasVariant bool  `json:"has_variant"`
            Tryouts []string  `json:"tryouts"`
            Description string  `json:"description"`
            Medias []Media1  `json:"medias"`
            RatingCount float64  `json:"rating_count"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            ItemType string  `json:"item_type"`
            ProductOnlineDate string  `json:"product_online_date"`
            Similars []string  `json:"similars"`
            ShortDescription string  `json:"short_description"`
            Attributes map[string]interface{}  `json:"attributes"`
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // InventoryPage used by Catalog
    type InventoryPage struct {

        
            Type string  `json:"type"`
            ItemTotal float64  `json:"item_total"`
            NextID string  `json:"next_id"`
            HasPrevious bool  `json:"has_previous"`
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
            ArticleAssignment ArticleAssignment  `json:"article_assignment"`
            GroupID string  `json:"group_id"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // AssignStore used by Catalog
    type AssignStore struct {

        
            CompanyID float64  `json:"company_id"`
            StoreIds []float64  `json:"store_ids"`
            AppID string  `json:"app_id"`
            ChannelType string  `json:"channel_type"`
            Articles []AssignStoreArticle  `json:"articles"`
            Pincode string  `json:"pincode"`
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
            Index float64  `json:"index"`
            CompanyID float64  `json:"company_id"`
            PriceMarked float64  `json:"price_marked"`
            SCity string  `json:"s_city"`
            Quantity float64  `json:"quantity"`
            StorePincode float64  `json:"store_pincode"`
            ArticleAssignment ArticleAssignment1  `json:"article_assignment"`
            UID string  `json:"uid"`
            Status bool  `json:"status"`
            GroupID string  `json:"group_id"`
            ItemID float64  `json:"item_id"`
            ID string  `json:"_id"`
            Meta map[string]interface{}  `json:"meta"`
            StoreID float64  `json:"store_id"`
            PriceEffective float64  `json:"price_effective"`
            Size string  `json:"size"`
         
    }
    
    // ProductReturnConfigSerializer used by Catalog
    type ProductReturnConfigSerializer struct {

        
            OnSameStore bool  `json:"on_same_store"`
            StoreUID float64  `json:"store_uid"`
         
    }
    
    // UserSerializer1 used by Catalog
    type UserSerializer1 struct {

        
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
            Contact string  `json:"contact"`
         
    }
    
    // GetAddressSerializer used by Catalog
    type GetAddressSerializer struct {

        
            State string  `json:"state"`
            Landmark string  `json:"landmark"`
            AddressType string  `json:"address_type"`
            Address2 string  `json:"address2"`
            Address1 string  `json:"address1"`
            Country string  `json:"country"`
            CountryCode string  `json:"country_code"`
            Latitude float64  `json:"latitude"`
            Pincode float64  `json:"pincode"`
            City string  `json:"city"`
            Longitude float64  `json:"longitude"`
         
    }
    
    // GetCompanySerializer used by Catalog
    type GetCompanySerializer struct {

        
            VerifiedBy UserSerializer1  `json:"verified_by"`
            CreatedOn string  `json:"created_on"`
            BusinessType string  `json:"business_type"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            UID float64  `json:"uid"`
            CompanyType string  `json:"company_type"`
            CreatedBy UserSerializer1  `json:"created_by"`
            Stage string  `json:"stage"`
            ModifiedOn string  `json:"modified_on"`
            ModifiedBy UserSerializer1  `json:"modified_by"`
            RejectReason string  `json:"reject_reason"`
            VerifiedOn string  `json:"verified_on"`
            Name string  `json:"name"`
         
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
    
    // Document used by Catalog
    type Document struct {

        
            Type string  `json:"type"`
            LegalName string  `json:"legal_name"`
            Value string  `json:"value"`
            Verified bool  `json:"verified"`
            URL string  `json:"url"`
         
    }
    
    // UserSerializer2 used by Catalog
    type UserSerializer2 struct {

        
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
            Contact string  `json:"contact"`
         
    }
    
    // LocationIntegrationType used by Catalog
    type LocationIntegrationType struct {

        
            Inventory string  `json:"inventory"`
            Order string  `json:"order"`
         
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
    
    // InvoiceCredSerializer used by Catalog
    type InvoiceCredSerializer struct {

        
            Username string  `json:"username"`
            Enabled bool  `json:"enabled"`
            Password string  `json:"password"`
         
    }
    
    // InvoiceDetailsSerializer used by Catalog
    type InvoiceDetailsSerializer struct {

        
            EWaybill InvoiceCredSerializer  `json:"e_waybill"`
            EInvoice InvoiceCredSerializer  `json:"e_invoice"`
         
    }
    
    // GetLocationSerializer used by Catalog
    type GetLocationSerializer struct {

        
            DisplayName string  `json:"display_name"`
            CreatedOn string  `json:"created_on"`
            Code string  `json:"code"`
            ProductReturnConfig ProductReturnConfigSerializer  `json:"product_return_config"`
            Company GetCompanySerializer  `json:"company"`
            Manager LocationManagerSerializer  `json:"manager"`
            Warnings map[string]interface{}  `json:"warnings"`
            Documents []Document  `json:"documents"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            VerifiedBy UserSerializer2  `json:"verified_by"`
            IntegrationType LocationIntegrationType  `json:"integration_type"`
            NotificationEmails []string  `json:"notification_emails"`
            CreatedBy UserSerializer2  `json:"created_by"`
            Timing []LocationDayWiseSerializer  `json:"timing"`
            VerifiedOn string  `json:"verified_on"`
            GstCredentials InvoiceDetailsSerializer  `json:"gst_credentials"`
            Stage string  `json:"stage"`
            ModifiedOn string  `json:"modified_on"`
            ModifiedBy UserSerializer2  `json:"modified_by"`
            Address GetAddressSerializer  `json:"address"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            StoreType string  `json:"store_type"`
            PhoneNumber string  `json:"phone_number"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
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
            IsActive bool  `json:"is_active"`
            Logo string  `json:"logo"`
            UID float64  `json:"uid"`
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
    

    
    // Document used by CompanyProfile
    type Document struct {

        
            Verified bool  `json:"verified"`
            Type string  `json:"type"`
            URL string  `json:"url"`
            LegalName string  `json:"legal_name"`
            Value string  `json:"value"`
         
    }
    
    // UserSerializer used by CompanyProfile
    type UserSerializer struct {

        
            Contact string  `json:"contact"`
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
         
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
            Address2 string  `json:"address2"`
            Landmark string  `json:"landmark"`
            Country string  `json:"country"`
            CountryCode string  `json:"country_code"`
            State string  `json:"state"`
            City string  `json:"city"`
            Address1 string  `json:"address1"`
            AddressType string  `json:"address_type"`
            Pincode float64  `json:"pincode"`
            Longitude float64  `json:"longitude"`
         
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

        
            ModifiedOn string  `json:"modified_on"`
            UID float64  `json:"uid"`
            Documents []Document  `json:"documents"`
            FranchiseEnabled bool  `json:"franchise_enabled"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            Warnings map[string]interface{}  `json:"warnings"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            CreatedOn string  `json:"created_on"`
            NotificationEmails []string  `json:"notification_emails"`
            CompanyType string  `json:"company_type"`
            Taxes []CompanyTaxesSerializer  `json:"taxes"`
            CreatedBy UserSerializer  `json:"created_by"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            VerifiedOn string  `json:"verified_on"`
            BusinessCountryInfo BusinessCountryInfo  `json:"business_country_info"`
            BusinessDetails BusinessDetails  `json:"business_details"`
            BusinessType string  `json:"business_type"`
            Mode string  `json:"mode"`
            Stage string  `json:"stage"`
            ContactDetails ContactDetails  `json:"contact_details"`
            BusinessInfo string  `json:"business_info"`
            Name string  `json:"name"`
            VerifiedBy UserSerializer  `json:"verified_by"`
         
    }
    
    // ErrorResponse used by CompanyProfile
    type ErrorResponse struct {

        
            Message string  `json:"message"`
            Status float64  `json:"status"`
            Code string  `json:"code"`
            Meta map[string]interface{}  `json:"meta"`
         
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
            Address2 string  `json:"address2"`
            Landmark string  `json:"landmark"`
            Country string  `json:"country"`
            CountryCode string  `json:"country_code"`
            State string  `json:"state"`
            City string  `json:"city"`
            Address1 string  `json:"address1"`
            AddressType string  `json:"address_type"`
            Pincode float64  `json:"pincode"`
            Longitude float64  `json:"longitude"`
         
    }
    
    // UpdateCompany used by CompanyProfile
    type UpdateCompany struct {

        
            RejectReason string  `json:"reject_reason"`
            CompanyType string  `json:"company_type"`
            NotificationEmails []string  `json:"notification_emails"`
            Documents []Document  `json:"documents"`
            FranchiseEnabled bool  `json:"franchise_enabled"`
            ContactDetails ContactDetails  `json:"contact_details"`
            BusinessDetails BusinessDetails  `json:"business_details"`
            BusinessType string  `json:"business_type"`
            Warnings map[string]interface{}  `json:"warnings"`
            BusinessInfo string  `json:"business_info"`
            Taxes []CompanyTaxesSerializer1  `json:"taxes"`
            Name string  `json:"name"`
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

        
            Verified float64  `json:"verified"`
            Pending float64  `json:"pending"`
         
    }
    
    // MetricsSerializer used by CompanyProfile
    type MetricsSerializer struct {

        
            CompanyDocuments DocumentsObj  `json:"company_documents"`
            UID float64  `json:"uid"`
            Brand DocumentsObj  `json:"brand"`
            StoreDocuments DocumentsObj  `json:"store_documents"`
            Store DocumentsObj  `json:"store"`
            Stage string  `json:"stage"`
            Product DocumentsObj  `json:"product"`
         
    }
    
    // BrandBannerSerializer used by CompanyProfile
    type BrandBannerSerializer struct {

        
            Portrait string  `json:"portrait"`
            Landscape string  `json:"landscape"`
         
    }
    
    // GetBrandResponseSerializer used by CompanyProfile
    type GetBrandResponseSerializer struct {

        
            ModifiedOn string  `json:"modified_on"`
            UID float64  `json:"uid"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            Warnings map[string]interface{}  `json:"warnings"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Logo string  `json:"logo"`
            CreatedOn string  `json:"created_on"`
            Synonyms []string  `json:"synonyms"`
            Banner BrandBannerSerializer  `json:"banner"`
            Description string  `json:"description"`
            CreatedBy UserSerializer  `json:"created_by"`
            VerifiedOn string  `json:"verified_on"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            Mode string  `json:"mode"`
            Stage string  `json:"stage"`
            SlugKey string  `json:"slug_key"`
            RejectReason string  `json:"reject_reason"`
            Name string  `json:"name"`
            VerifiedBy UserSerializer  `json:"verified_by"`
         
    }
    
    // CreateUpdateBrandRequestSerializer used by CompanyProfile
    type CreateUpdateBrandRequestSerializer struct {

        
            Logo string  `json:"logo"`
            UID float64  `json:"uid"`
            Banner BrandBannerSerializer  `json:"banner"`
            BrandTier string  `json:"brand_tier"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            Synonyms []string  `json:"synonyms"`
            CompanyID float64  `json:"company_id"`
            Name string  `json:"name"`
            Description string  `json:"description"`
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

        
            ModifiedOn string  `json:"modified_on"`
            UID float64  `json:"uid"`
            VerifiedOn string  `json:"verified_on"`
            CreatedOn string  `json:"created_on"`
            NotificationEmails []string  `json:"notification_emails"`
            CompanyType string  `json:"company_type"`
            RejectReason string  `json:"reject_reason"`
            BusinessCountryInfo BusinessCountryInfo  `json:"business_country_info"`
            Details CompanyDetails  `json:"details"`
            BusinessType string  `json:"business_type"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            MarketChannels []string  `json:"market_channels"`
            Stage string  `json:"stage"`
            Name string  `json:"name"`
            CreatedBy UserSerializer  `json:"created_by"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Addresses []GetAddressSerializer  `json:"addresses"`
         
    }
    
    // CompanyBrandSerializer used by CompanyProfile
    type CompanyBrandSerializer struct {

        
            ModifiedOn string  `json:"modified_on"`
            UID float64  `json:"uid"`
            VerifiedOn string  `json:"verified_on"`
            CreatedOn string  `json:"created_on"`
            RejectReason string  `json:"reject_reason"`
            Brand GetBrandResponseSerializer  `json:"brand"`
            Company CompanySerializer  `json:"company"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            Stage string  `json:"stage"`
            Warnings map[string]interface{}  `json:"warnings"`
            CreatedBy UserSerializer  `json:"created_by"`
            VerifiedBy UserSerializer  `json:"verified_by"`
         
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
    
    // GetCompanySerializer used by CompanyProfile
    type GetCompanySerializer struct {

        
            ModifiedOn string  `json:"modified_on"`
            UID float64  `json:"uid"`
            VerifiedOn string  `json:"verified_on"`
            CreatedOn string  `json:"created_on"`
            RejectReason string  `json:"reject_reason"`
            CompanyType string  `json:"company_type"`
            BusinessType string  `json:"business_type"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            Stage string  `json:"stage"`
            Name string  `json:"name"`
            CreatedBy UserSerializer  `json:"created_by"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            Addresses []GetAddressSerializer  `json:"addresses"`
         
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

        
            Date HolidayDateSerializer  `json:"date"`
            Title string  `json:"title"`
            HolidayType string  `json:"holiday_type"`
         
    }
    
    // InvoiceCredSerializer used by CompanyProfile
    type InvoiceCredSerializer struct {

        
            Password string  `json:"password"`
            Enabled bool  `json:"enabled"`
            Username string  `json:"username"`
         
    }
    
    // InvoiceDetailsSerializer used by CompanyProfile
    type InvoiceDetailsSerializer struct {

        
            EWaybill InvoiceCredSerializer  `json:"e_waybill"`
            EInvoice InvoiceCredSerializer  `json:"e_invoice"`
         
    }
    
    // LocationTimingSerializer used by CompanyProfile
    type LocationTimingSerializer struct {

        
            Hour float64  `json:"hour"`
            Minute float64  `json:"minute"`
         
    }
    
    // LocationDayWiseSerializer used by CompanyProfile
    type LocationDayWiseSerializer struct {

        
            Opening LocationTimingSerializer  `json:"opening"`
            Open bool  `json:"open"`
            Weekday string  `json:"weekday"`
            Closing LocationTimingSerializer  `json:"closing"`
         
    }
    
    // LocationManagerSerializer used by CompanyProfile
    type LocationManagerSerializer struct {

        
            Name string  `json:"name"`
            MobileNo SellerPhoneNumber  `json:"mobile_no"`
            Email string  `json:"email"`
         
    }
    
    // GetLocationSerializer used by CompanyProfile
    type GetLocationSerializer struct {

        
            ModifiedOn string  `json:"modified_on"`
            UID float64  `json:"uid"`
            Documents []Document  `json:"documents"`
            Company GetCompanySerializer  `json:"company"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            Warnings map[string]interface{}  `json:"warnings"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ProductReturnConfig ProductReturnConfigSerializer  `json:"product_return_config"`
            CreatedOn string  `json:"created_on"`
            StoreType string  `json:"store_type"`
            NotificationEmails []string  `json:"notification_emails"`
            PhoneNumber string  `json:"phone_number"`
            Holiday []HolidaySchemaSerializer  `json:"holiday"`
            GstCredentials InvoiceDetailsSerializer  `json:"gst_credentials"`
            DisplayName string  `json:"display_name"`
            Address GetAddressSerializer  `json:"address"`
            CreatedBy UserSerializer  `json:"created_by"`
            Timing []LocationDayWiseSerializer  `json:"timing"`
            VerifiedOn string  `json:"verified_on"`
            Code string  `json:"code"`
            Manager LocationManagerSerializer  `json:"manager"`
            Stage string  `json:"stage"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            Name string  `json:"name"`
            VerifiedBy UserSerializer  `json:"verified_by"`
         
    }
    
    // LocationListSerializer used by CompanyProfile
    type LocationListSerializer struct {

        
            Page Page  `json:"page"`
            Items []GetLocationSerializer  `json:"items"`
         
    }
    
    // AddressSerializer used by CompanyProfile
    type AddressSerializer struct {

        
            Latitude float64  `json:"latitude"`
            Address2 string  `json:"address2"`
            Landmark string  `json:"landmark"`
            Country string  `json:"country"`
            CountryCode string  `json:"country_code"`
            State string  `json:"state"`
            City string  `json:"city"`
            Address1 string  `json:"address1"`
            AddressType string  `json:"address_type"`
            Pincode float64  `json:"pincode"`
            Longitude float64  `json:"longitude"`
         
    }
    
    // LocationSerializer used by CompanyProfile
    type LocationSerializer struct {

        
            Timing []LocationDayWiseSerializer  `json:"timing"`
            UID float64  `json:"uid"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            StoreType string  `json:"store_type"`
            NotificationEmails []string  `json:"notification_emails"`
            Documents []Document  `json:"documents"`
            Holiday []HolidaySchemaSerializer  `json:"holiday"`
            Code string  `json:"code"`
            Company float64  `json:"company"`
            Manager LocationManagerSerializer  `json:"manager"`
            GstCredentials InvoiceDetailsSerializer  `json:"gst_credentials"`
            Warnings map[string]interface{}  `json:"warnings"`
            Stage string  `json:"stage"`
            DisplayName string  `json:"display_name"`
            Address AddressSerializer  `json:"address"`
            Name string  `json:"name"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ProductReturnConfig ProductReturnConfigSerializer  `json:"product_return_config"`
         
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
    
    // Rule used by Cart
    type Rule struct {

        
            Max float64  `json:"max"`
            Key float64  `json:"key"`
            DiscountQty float64  `json:"discount_qty"`
            Min float64  `json:"min"`
            Value float64  `json:"value"`
         
    }
    
    // DisplayMetaDict used by Cart
    type DisplayMetaDict struct {

        
            Title string  `json:"title"`
            Subtitle string  `json:"subtitle"`
         
    }
    
    // DisplayMeta used by Cart
    type DisplayMeta struct {

        
            Remove DisplayMetaDict  `json:"remove"`
            Auto DisplayMetaDict  `json:"auto"`
            Apply DisplayMetaDict  `json:"apply"`
            Description string  `json:"description"`
            Title string  `json:"title"`
            Subtitle string  `json:"subtitle"`
         
    }
    
    // Ownership used by Cart
    type Ownership struct {

        
            PayableCategory string  `json:"payable_category"`
            PayableBy string  `json:"payable_by"`
         
    }
    
    // CouponAuthor used by Cart
    type CouponAuthor struct {

        
            CreatedBy string  `json:"created_by"`
            ModifiedBy string  `json:"modified_by"`
         
    }
    
    // Identifier used by Cart
    type Identifier struct {

        
            CategoryID []float64  `json:"category_id"`
            ExcludeBrandID []float64  `json:"exclude_brand_id"`
            BrandID []float64  `json:"brand_id"`
            ArticleID []string  `json:"article_id"`
            StoreID []float64  `json:"store_id"`
            UserID []string  `json:"user_id"`
            CollectionID []string  `json:"collection_id"`
            CompanyID []float64  `json:"company_id"`
            ItemID []float64  `json:"item_id"`
         
    }
    
    // RuleDefinition used by Cart
    type RuleDefinition struct {

        
            ApplicableOn string  `json:"applicable_on"`
            Type string  `json:"type"`
            CalculateOn string  `json:"calculate_on"`
            IsExact bool  `json:"is_exact"`
            Scope []string  `json:"scope"`
            AutoApply bool  `json:"auto_apply"`
            CurrencyCode string  `json:"currency_code"`
            ValueType string  `json:"value_type"`
         
    }
    
    // PriceRange used by Cart
    type PriceRange struct {

        
            Max float64  `json:"max"`
            Min float64  `json:"min"`
         
    }
    
    // UsesRemaining used by Cart
    type UsesRemaining struct {

        
            Total float64  `json:"total"`
            User float64  `json:"user"`
            App float64  `json:"app"`
         
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

        
            Networks []string  `json:"networks"`
            Codes []string  `json:"codes"`
            Uses PaymentAllowValue  `json:"uses"`
            Types []string  `json:"types"`
         
    }
    
    // BulkBundleRestriction used by Cart
    type BulkBundleRestriction struct {

        
            MultiStoreAllowed bool  `json:"multi_store_allowed"`
         
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
            Platforms []string  `json:"platforms"`
            Uses UsesRestriction  `json:"uses"`
            Payments map[string]PaymentModes  `json:"payments"`
            BulkBundle BulkBundleRestriction  `json:"bulk_bundle"`
            CouponAllowed bool  `json:"coupon_allowed"`
            OrderingStores []float64  `json:"ordering_stores"`
            PostOrder PostOrder  `json:"post_order"`
         
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
            UserRegisteredAfter string  `json:"user_registered_after"`
            Anonymous bool  `json:"anonymous"`
         
    }
    
    // CouponAction used by Cart
    type CouponAction struct {

        
            TxnMode string  `json:"txn_mode"`
            ActionDate string  `json:"action_date"`
         
    }
    
    // CouponSchedule used by Cart
    type CouponSchedule struct {

        
            End string  `json:"end"`
            Duration float64  `json:"duration"`
            NextSchedule []map[string]interface{}  `json:"next_schedule"`
            Cron string  `json:"cron"`
            Start string  `json:"start"`
         
    }
    
    // CouponAdd used by Cart
    type CouponAdd struct {

        
            DateMeta CouponDateMeta  `json:"date_meta"`
            Validity Validity  `json:"validity"`
            Rule []Rule  `json:"rule"`
            DisplayMeta DisplayMeta  `json:"display_meta"`
            Ownership Ownership  `json:"ownership"`
            Author CouponAuthor  `json:"author"`
            Identifiers Identifier  `json:"identifiers"`
            RuleDefinition RuleDefinition  `json:"rule_definition"`
            Restrictions Restrictions  `json:"restrictions"`
            State State  `json:"state"`
            Validation Validation  `json:"validation"`
            Action CouponAction  `json:"action"`
            Tags []string  `json:"tags"`
            Schedule CouponSchedule  `json:"_schedule"`
            TypeSlug string  `json:"type_slug"`
            Code string  `json:"code"`
         
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
            Rule []Rule  `json:"rule"`
            DisplayMeta DisplayMeta  `json:"display_meta"`
            Ownership Ownership  `json:"ownership"`
            Author CouponAuthor  `json:"author"`
            Identifiers Identifier  `json:"identifiers"`
            RuleDefinition RuleDefinition  `json:"rule_definition"`
            Restrictions Restrictions  `json:"restrictions"`
            State State  `json:"state"`
            Validation Validation  `json:"validation"`
            Action CouponAction  `json:"action"`
            Tags []string  `json:"tags"`
            Schedule CouponSchedule  `json:"_schedule"`
            TypeSlug string  `json:"type_slug"`
            Code string  `json:"code"`
         
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

        
            End string  `json:"end"`
            Start string  `json:"start"`
         
    }
    
    // PaymentAllowValue1 used by Cart
    type PaymentAllowValue1 struct {

        
            Max float64  `json:"max"`
         
    }
    
    // PromotionPaymentModes used by Cart
    type PromotionPaymentModes struct {

        
            Type string  `json:"type"`
            Codes []string  `json:"codes"`
            Uses PaymentAllowValue1  `json:"uses"`
         
    }
    
    // PostOrder1 used by Cart
    type PostOrder1 struct {

        
            ReturnAllowed bool  `json:"return_allowed"`
            CancellationAllowed bool  `json:"cancellation_allowed"`
         
    }
    
    // Restrictions1 used by Cart
    type Restrictions1 struct {

        
            UserGroups []float64  `json:"user_groups"`
            Platforms []string  `json:"platforms"`
            Uses UsesRestriction1  `json:"uses"`
            UserRegistered UserRegistered  `json:"user_registered"`
            Payments []PromotionPaymentModes  `json:"payments"`
            UserID []string  `json:"user_id"`
            AnonymousUsers bool  `json:"anonymous_users"`
            OrderQuantity float64  `json:"order_quantity"`
            PostOrder PostOrder1  `json:"post_order"`
         
    }
    
    // CompareObject used by Cart
    type CompareObject struct {

        
            GreaterThanEquals float64  `json:"greater_than_equals"`
            Equals float64  `json:"equals"`
            LessThan float64  `json:"less_than"`
            LessThanEquals float64  `json:"less_than_equals"`
            GreaterThan float64  `json:"greater_than"`
         
    }
    
    // ItemCriteria used by Cart
    type ItemCriteria struct {

        
            ItemExcludeCategory []float64  `json:"item_exclude_category"`
            AllItems bool  `json:"all_items"`
            ItemStore []float64  `json:"item_store"`
            ItemExcludeBrand []float64  `json:"item_exclude_brand"`
            ItemExcludeSku []string  `json:"item_exclude_sku"`
            ItemExcludeCompany []float64  `json:"item_exclude_company"`
            ItemCategory []float64  `json:"item_category"`
            ItemBrand []float64  `json:"item_brand"`
            ItemCompany []float64  `json:"item_company"`
            BuyRules []string  `json:"buy_rules"`
            ItemSize []string  `json:"item_size"`
            CartQuantity CompareObject  `json:"cart_quantity"`
            ItemSku []string  `json:"item_sku"`
            ItemExcludeStore []float64  `json:"item_exclude_store"`
            ItemID []float64  `json:"item_id"`
            CartTotal CompareObject  `json:"cart_total"`
         
    }
    
    // Visibility used by Cart
    type Visibility struct {

        
            CouponList bool  `json:"coupon_list"`
            Pdp bool  `json:"pdp"`
         
    }
    
    // Ownership1 used by Cart
    type Ownership1 struct {

        
            PayableCategory string  `json:"payable_category"`
            PayableBy string  `json:"payable_by"`
         
    }
    
    // PromotionDateMeta used by Cart
    type PromotionDateMeta struct {

        
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // PromotionAuthor used by Cart
    type PromotionAuthor struct {

        
            CreatedBy string  `json:"created_by"`
            ModifiedBy string  `json:"modified_by"`
         
    }
    
    // PromotionAction used by Cart
    type PromotionAction struct {

        
            ActionType string  `json:"action_type"`
            ActionDate string  `json:"action_date"`
         
    }
    
    // PromotionSchedule used by Cart
    type PromotionSchedule struct {

        
            End string  `json:"end"`
            Duration float64  `json:"duration"`
            Published bool  `json:"published"`
            NextSchedule []map[string]interface{}  `json:"next_schedule"`
            Cron string  `json:"cron"`
            Start string  `json:"start"`
         
    }
    
    // DiscountOffer used by Cart
    type DiscountOffer struct {

        
            MinOfferQuantity float64  `json:"min_offer_quantity"`
            DiscountPercentage float64  `json:"discount_percentage"`
            DiscountAmount float64  `json:"discount_amount"`
            MaxOfferQuantity float64  `json:"max_offer_quantity"`
            MaxDiscountAmount float64  `json:"max_discount_amount"`
            DiscountPrice float64  `json:"discount_price"`
            Code string  `json:"code"`
         
    }
    
    // DiscountRule used by Cart
    type DiscountRule struct {

        
            Offer DiscountOffer  `json:"offer"`
            BuyCondition string  `json:"buy_condition"`
            ItemCriteria ItemCriteria  `json:"item_criteria"`
            DiscountType string  `json:"discount_type"`
         
    }
    
    // PromotionListItem used by Cart
    type PromotionListItem struct {

        
            ApplyExclusive string  `json:"apply_exclusive"`
            DisplayMeta DisplayMeta1  `json:"display_meta"`
            PromoGroup string  `json:"promo_group"`
            Restrictions Restrictions1  `json:"restrictions"`
            ApplyAllDiscount bool  `json:"apply_all_discount"`
            BuyRules map[string]ItemCriteria  `json:"buy_rules"`
            Visiblility Visibility  `json:"visiblility"`
            Mode string  `json:"mode"`
            ApplyPriority float64  `json:"apply_priority"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Ownership Ownership1  `json:"ownership"`
            Stackable bool  `json:"stackable"`
            ApplicationID string  `json:"application_id"`
            DateMeta PromotionDateMeta  `json:"date_meta"`
            PromotionType string  `json:"promotion_type"`
            Author PromotionAuthor  `json:"author"`
            PostOrderAction PromotionAction  `json:"post_order_action"`
            Currency string  `json:"currency"`
            Schedule PromotionSchedule  `json:"_schedule"`
            DiscountRules []DiscountRule  `json:"discount_rules"`
            Code string  `json:"code"`
         
    }
    
    // PromotionsResponse used by Cart
    type PromotionsResponse struct {

        
            Page Page  `json:"page"`
            Items PromotionListItem  `json:"items"`
         
    }
    
    // PromotionAdd used by Cart
    type PromotionAdd struct {

        
            ApplyExclusive string  `json:"apply_exclusive"`
            DisplayMeta DisplayMeta1  `json:"display_meta"`
            PromoGroup string  `json:"promo_group"`
            Restrictions Restrictions1  `json:"restrictions"`
            ApplyAllDiscount bool  `json:"apply_all_discount"`
            BuyRules map[string]ItemCriteria  `json:"buy_rules"`
            Visiblility Visibility  `json:"visiblility"`
            Mode string  `json:"mode"`
            ApplyPriority float64  `json:"apply_priority"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Ownership Ownership1  `json:"ownership"`
            Stackable bool  `json:"stackable"`
            ApplicationID string  `json:"application_id"`
            DateMeta PromotionDateMeta  `json:"date_meta"`
            PromotionType string  `json:"promotion_type"`
            Author PromotionAuthor  `json:"author"`
            PostOrderAction PromotionAction  `json:"post_order_action"`
            Currency string  `json:"currency"`
            Schedule PromotionSchedule  `json:"_schedule"`
            DiscountRules []DiscountRule  `json:"discount_rules"`
            Code string  `json:"code"`
         
    }
    
    // PromotionUpdate used by Cart
    type PromotionUpdate struct {

        
            ApplyExclusive string  `json:"apply_exclusive"`
            DisplayMeta DisplayMeta1  `json:"display_meta"`
            PromoGroup string  `json:"promo_group"`
            Restrictions Restrictions1  `json:"restrictions"`
            ApplyAllDiscount bool  `json:"apply_all_discount"`
            BuyRules map[string]ItemCriteria  `json:"buy_rules"`
            Visiblility Visibility  `json:"visiblility"`
            Mode string  `json:"mode"`
            ApplyPriority float64  `json:"apply_priority"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Ownership Ownership1  `json:"ownership"`
            Stackable bool  `json:"stackable"`
            ApplicationID string  `json:"application_id"`
            DateMeta PromotionDateMeta  `json:"date_meta"`
            PromotionType string  `json:"promotion_type"`
            Author PromotionAuthor  `json:"author"`
            PostOrderAction PromotionAction  `json:"post_order_action"`
            Currency string  `json:"currency"`
            Schedule PromotionSchedule  `json:"_schedule"`
            DiscountRules []DiscountRule  `json:"discount_rules"`
            Code string  `json:"code"`
         
    }
    
    // PromotionPartialUpdate used by Cart
    type PromotionPartialUpdate struct {

        
            Schedule PromotionSchedule  `json:"schedule"`
            Archive bool  `json:"archive"`
         
    }
    
    // CartItem used by Cart
    type CartItem struct {

        
            Size string  `json:"size"`
            ProductID string  `json:"product_id"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // OpenapiCartDetailsRequest used by Cart
    type OpenapiCartDetailsRequest struct {

        
            CartItems CartItem  `json:"cart_items"`
         
    }
    
    // ProductPrice used by Cart
    type ProductPrice struct {

        
            Marked float64  `json:"marked"`
            CurrencySymbol string  `json:"currency_symbol"`
            AddOn float64  `json:"add_on"`
            Effective float64  `json:"effective"`
            CurrencyCode string  `json:"currency_code"`
            Selling float64  `json:"selling"`
         
    }
    
    // ProductPriceInfo used by Cart
    type ProductPriceInfo struct {

        
            Converted ProductPrice  `json:"converted"`
            Base ProductPrice  `json:"base"`
         
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
    
    // CartProductIdentifer used by Cart
    type CartProductIdentifer struct {

        
            Identifier string  `json:"identifier"`
         
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
            URL string  `json:"url"`
            SecureURL string  `json:"secure_url"`
         
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
    
    // CartProduct used by Cart
    type CartProduct struct {

        
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Action ProductAction  `json:"action"`
            Type string  `json:"type"`
            Images []ProductImage  `json:"images"`
            Categories []CategoryInfo  `json:"categories"`
            Brand BaseInfo  `json:"brand"`
            UID float64  `json:"uid"`
         
    }
    
    // BasePrice used by Cart
    type BasePrice struct {

        
            Marked float64  `json:"marked"`
            CurrencySymbol string  `json:"currency_symbol"`
            CurrencyCode string  `json:"currency_code"`
            Effective float64  `json:"effective"`
         
    }
    
    // ArticlePriceInfo used by Cart
    type ArticlePriceInfo struct {

        
            Converted BasePrice  `json:"converted"`
            Base BasePrice  `json:"base"`
         
    }
    
    // ProductArticle used by Cart
    type ProductArticle struct {

        
            Size string  `json:"size"`
            Price ArticlePriceInfo  `json:"price"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            Type string  `json:"type"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            Store BaseInfo  `json:"store"`
            Seller BaseInfo  `json:"seller"`
            ProductGroupTags []string  `json:"product_group_tags"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            UID string  `json:"uid"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // Ownership2 used by Cart
    type Ownership2 struct {

        
            PayableCategory string  `json:"payable_category"`
            PayableBy string  `json:"payable_by"`
         
    }
    
    // AppliedPromotion used by Cart
    type AppliedPromotion struct {

        
            PromoID string  `json:"promo_id"`
            PromotionType string  `json:"promotion_type"`
            OfferText string  `json:"offer_text"`
            ArticleQuantity float64  `json:"article_quantity"`
            Amount float64  `json:"amount"`
            MrpPromotion bool  `json:"mrp_promotion"`
            Ownership Ownership2  `json:"ownership"`
         
    }
    
    // PromoMeta used by Cart
    type PromoMeta struct {

        
            Message string  `json:"message"`
         
    }
    
    // ProductAvailabilitySize used by Cart
    type ProductAvailabilitySize struct {

        
            Display string  `json:"display"`
            IsAvailable bool  `json:"is_available"`
            Value string  `json:"value"`
         
    }
    
    // ProductAvailability used by Cart
    type ProductAvailability struct {

        
            OtherStoreQuantity float64  `json:"other_store_quantity"`
            IsValid bool  `json:"is_valid"`
            Deliverable bool  `json:"deliverable"`
            AvailableSizes []ProductAvailabilitySize  `json:"available_sizes"`
            Sizes []string  `json:"sizes"`
            OutOfStock bool  `json:"out_of_stock"`
         
    }
    
    // CartProductInfo used by Cart
    type CartProductInfo struct {

        
            Discount string  `json:"discount"`
            Price ProductPriceInfo  `json:"price"`
            PricePerUnit ProductPriceInfo  `json:"price_per_unit"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
            Product CartProduct  `json:"product"`
            CouponMessage string  `json:"coupon_message"`
            Article ProductArticle  `json:"article"`
            BulkOffer map[string]interface{}  `json:"bulk_offer"`
            IsSet bool  `json:"is_set"`
            PromotionsApplied []AppliedPromotion  `json:"promotions_applied"`
            Moq map[string]interface{}  `json:"moq"`
            Key string  `json:"key"`
            Message string  `json:"message"`
            PromoMeta PromoMeta  `json:"promo_meta"`
            Availability ProductAvailability  `json:"availability"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // RawBreakup used by Cart
    type RawBreakup struct {

        
            MrpTotal float64  `json:"mrp_total"`
            Discount float64  `json:"discount"`
            FyndCash float64  `json:"fynd_cash"`
            GstCharges float64  `json:"gst_charges"`
            ConvenienceFee float64  `json:"convenience_fee"`
            DeliveryCharge float64  `json:"delivery_charge"`
            Vog float64  `json:"vog"`
            YouSaved float64  `json:"you_saved"`
            Coupon float64  `json:"coupon"`
            CodCharge float64  `json:"cod_charge"`
            Subtotal float64  `json:"subtotal"`
            Total float64  `json:"total"`
         
    }
    
    // LoyaltyPoints used by Cart
    type LoyaltyPoints struct {

        
            Description string  `json:"description"`
            Total float64  `json:"total"`
            IsApplied bool  `json:"is_applied"`
            Applicable float64  `json:"applicable"`
         
    }
    
    // DisplayBreakup used by Cart
    type DisplayBreakup struct {

        
            CurrencySymbol string  `json:"currency_symbol"`
            Display string  `json:"display"`
            CurrencyCode string  `json:"currency_code"`
            Key string  `json:"key"`
            Message []string  `json:"message"`
            Value float64  `json:"value"`
         
    }
    
    // CouponBreakup used by Cart
    type CouponBreakup struct {

        
            IsApplied bool  `json:"is_applied"`
            Type string  `json:"type"`
            Code string  `json:"code"`
            Message string  `json:"message"`
            UID string  `json:"uid"`
            Value float64  `json:"value"`
         
    }
    
    // CartBreakup used by Cart
    type CartBreakup struct {

        
            Raw RawBreakup  `json:"raw"`
            LoyaltyPoints LoyaltyPoints  `json:"loyalty_points"`
            Display []DisplayBreakup  `json:"display"`
            Coupon CouponBreakup  `json:"coupon"`
         
    }
    
    // OpenapiCartDetailsResponse used by Cart
    type OpenapiCartDetailsResponse struct {

        
            Message string  `json:"message"`
            IsValid bool  `json:"is_valid"`
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
            Pincode float64  `json:"pincode"`
            Area string  `json:"area"`
            Phone float64  `json:"phone"`
            State string  `json:"state"`
            AddressType string  `json:"address_type"`
            City string  `json:"city"`
            Email string  `json:"email"`
            Meta map[string]interface{}  `json:"meta"`
            CountryCode string  `json:"country_code"`
            AreaCode string  `json:"area_code"`
            Country string  `json:"country"`
            Landmark string  `json:"landmark"`
            Address string  `json:"address"`
            AreaCodeSlug string  `json:"area_code_slug"`
         
    }
    
    // OpenApiCartServiceabilityRequest used by Cart
    type OpenApiCartServiceabilityRequest struct {

        
            ShippingAddress ShippingAddress  `json:"shipping_address"`
            CartItems CartItem  `json:"cart_items"`
         
    }
    
    // OpenApiCartServiceabilityResponse used by Cart
    type OpenApiCartServiceabilityResponse struct {

        
            IsValid bool  `json:"is_valid"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            Message string  `json:"message"`
            Items []CartProductInfo  `json:"items"`
         
    }
    
    // MultiTenderPaymentMeta used by Cart
    type MultiTenderPaymentMeta struct {

        
            PaymentGateway string  `json:"payment_gateway"`
            PaymentID string  `json:"payment_id"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            OrderID string  `json:"order_id"`
            CurrentStatus string  `json:"current_status"`
         
    }
    
    // MultiTenderPaymentMethod used by Cart
    type MultiTenderPaymentMethod struct {

        
            Name string  `json:"name"`
            Meta MultiTenderPaymentMeta  `json:"meta"`
            Amount float64  `json:"amount"`
            Mode string  `json:"mode"`
         
    }
    
    // OpenApiFiles used by Cart
    type OpenApiFiles struct {

        
            Values []string  `json:"values"`
            Key string  `json:"key"`
         
    }
    
    // CartItemMeta used by Cart
    type CartItemMeta struct {

        
            PrimaryItem bool  `json:"primary_item"`
            GroupID string  `json:"group_id"`
         
    }
    
    // OpenApiOrderItem used by Cart
    type OpenApiOrderItem struct {

        
            Discount float64  `json:"discount"`
            CashbackApplied float64  `json:"cashback_applied"`
            PaymentMethods []MultiTenderPaymentMethod  `json:"payment_methods"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            PriceMarked float64  `json:"price_marked"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            Files []OpenApiFiles  `json:"files"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            Size string  `json:"size"`
            DeliveryCharges float64  `json:"delivery_charges"`
            Meta CartItemMeta  `json:"meta"`
            PriceEffective float64  `json:"price_effective"`
            EmployeeDiscount float64  `json:"employee_discount"`
            AmountPaid float64  `json:"amount_paid"`
            CodCharges float64  `json:"cod_charges"`
            ProductID float64  `json:"product_id"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // OpenApiPlatformCheckoutReq used by Cart
    type OpenApiPlatformCheckoutReq struct {

        
            CartValue float64  `json:"cart_value"`
            CartItems []OpenApiOrderItem  `json:"cart_items"`
            PaymentMethods []MultiTenderPaymentMethod  `json:"payment_methods"`
            PaymentMode string  `json:"payment_mode"`
            CashbackApplied float64  `json:"cashback_applied"`
            Gstin string  `json:"gstin"`
            OrderID string  `json:"order_id"`
            CouponValue float64  `json:"coupon_value"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            EmployeeDiscount map[string]interface{}  `json:"employee_discount"`
            BillingAddress ShippingAddress  `json:"billing_address"`
            ShippingAddress ShippingAddress  `json:"shipping_address"`
            Coupon string  `json:"coupon"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            DeliveryCharges float64  `json:"delivery_charges"`
            CodCharges float64  `json:"cod_charges"`
            Comment string  `json:"comment"`
            CurrencyCode string  `json:"currency_code"`
            Files []OpenApiFiles  `json:"files"`
            CouponCode string  `json:"coupon_code"`
         
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

        
            CartValue float64  `json:"cart_value"`
            PaymentMethods []map[string]interface{}  `json:"payment_methods"`
            PaymentMode string  `json:"payment_mode"`
            Articles []map[string]interface{}  `json:"articles"`
            UserID string  `json:"user_id"`
            Meta map[string]interface{}  `json:"meta"`
            BulkCouponDiscount float64  `json:"bulk_coupon_discount"`
            AppID string  `json:"app_id"`
            IsActive bool  `json:"is_active"`
            Coupon map[string]interface{}  `json:"coupon"`
            Promotion map[string]interface{}  `json:"promotion"`
            CodCharges map[string]interface{}  `json:"cod_charges"`
            Comment string  `json:"comment"`
            PickUpCustomerDetails map[string]interface{}  `json:"pick_up_customer_details"`
            CreatedOn string  `json:"created_on"`
            Discount float64  `json:"discount"`
            LastModified string  `json:"last_modified"`
            Cashback map[string]interface{}  `json:"cashback"`
            Gstin string  `json:"gstin"`
            OrderID string  `json:"order_id"`
            ExpireAt string  `json:"expire_at"`
            FcIndexMap []float64  `json:"fc_index_map"`
            ID string  `json:"_id"`
            CheckoutMode string  `json:"checkout_mode"`
            UID float64  `json:"uid"`
            FyndCredits map[string]interface{}  `json:"fynd_credits"`
            IsDefault bool  `json:"is_default"`
            Payments map[string]interface{}  `json:"payments"`
            MergeQty bool  `json:"merge_qty"`
            BuyNow bool  `json:"buy_now"`
            DeliveryCharges map[string]interface{}  `json:"delivery_charges"`
            Shipments []map[string]interface{}  `json:"shipments"`
            IsArchive bool  `json:"is_archive"`
         
    }
    
    // AbandonedCartResponse used by Cart
    type AbandonedCartResponse struct {

        
            Page Page  `json:"page"`
            Result map[string]interface{}  `json:"result"`
            Message string  `json:"message"`
            Success bool  `json:"success"`
            Items []AbandonedCart  `json:"items"`
         
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

        
            IsValid bool  `json:"is_valid"`
            LastModified string  `json:"last_modified"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            ID string  `json:"id"`
            Gstin string  `json:"gstin"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            CouponText string  `json:"coupon_text"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            BuyNow bool  `json:"buy_now"`
            Currency CartCurrency  `json:"currency"`
            Comment string  `json:"comment"`
            CheckoutMode string  `json:"checkout_mode"`
            Message string  `json:"message"`
            AppliedPromoDetails []AppliedPromotion  `json:"applied_promo_details"`
            Items []CartProductInfo  `json:"items"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
         
    }
    
    // AddProductCart used by Cart
    type AddProductCart struct {

        
            ArticleID string  `json:"article_id"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            SellerID float64  `json:"seller_id"`
            StoreID float64  `json:"store_id"`
            ProductGroupTags []string  `json:"product_group_tags"`
            Display string  `json:"display"`
            ItemSize string  `json:"item_size"`
            ArticleAssignment map[string]interface{}  `json:"article_assignment"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Pos bool  `json:"pos"`
            ItemID float64  `json:"item_id"`
            Quantity float64  `json:"quantity"`
         
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

        
            ArticleID string  `json:"article_id"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            ItemIndex float64  `json:"item_index"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
            ItemSize string  `json:"item_size"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ItemID float64  `json:"item_id"`
            Quantity float64  `json:"quantity"`
         
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
    
