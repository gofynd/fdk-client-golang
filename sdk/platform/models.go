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

        
            Success bool  `json:"success"`
            DisplayFields []string  `json:"display_fields"`
            Created bool  `json:"created"`
            Aggregators []map[string]interface{}  `json:"aggregators"`
            AppID string  `json:"app_id"`
            ExcludedFields []string  `json:"excluded_fields"`
         
    }
    
    // ErrorCodeDescription used by Payment
    type ErrorCodeDescription struct {

        
            Code string  `json:"code"`
            Description string  `json:"description"`
            Success bool  `json:"success"`
         
    }
    
    // PaymentGatewayConfig used by Payment
    type PaymentGatewayConfig struct {

        
            ConfigType string  `json:"config_type"`
            Key string  `json:"key"`
            Secret string  `json:"secret"`
            IsActive bool  `json:"is_active"`
            MerchantSalt string  `json:"merchant_salt"`
         
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

        
            Small string  `json:"small"`
            Large string  `json:"large"`
         
    }
    
    // IntentApp used by Payment
    type IntentApp struct {

        
            Code string  `json:"code"`
            Logos PaymentModeLogo  `json:"logos"`
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

        
            CompliantWithTokenisationGuidelines bool  `json:"compliant_with_tokenisation_guidelines"`
            CardToken string  `json:"card_token"`
            Name string  `json:"name"`
            DisplayName string  `json:"display_name"`
            CardName string  `json:"card_name"`
            LogoURL PaymentModeLogo  `json:"logo_url"`
            CodLimitPerOrder float64  `json:"cod_limit_per_order"`
            CardReference string  `json:"card_reference"`
            CardIssuer string  `json:"card_issuer"`
            CardFingerprint string  `json:"card_fingerprint"`
            IntentAppErrorList []string  `json:"intent_app_error_list"`
            IntentApp []IntentApp  `json:"intent_app"`
            IntentAppErrorDictList []IntentAppErrorList  `json:"intent_app_error_dict_list"`
            DisplayPriority float64  `json:"display_priority"`
            MerchantCode string  `json:"merchant_code"`
            RemainingLimit float64  `json:"remaining_limit"`
            CardNumber string  `json:"card_number"`
            IntentFlow bool  `json:"intent_flow"`
            ExpYear float64  `json:"exp_year"`
            CardBrandImage string  `json:"card_brand_image"`
            RetryCount float64  `json:"retry_count"`
            Timeout float64  `json:"timeout"`
            Expired bool  `json:"expired"`
            CardID string  `json:"card_id"`
            AggregatorName string  `json:"aggregator_name"`
            Nickname string  `json:"nickname"`
            CardBrand string  `json:"card_brand"`
            Code string  `json:"code"`
            CardType string  `json:"card_type"`
            CodLimit float64  `json:"cod_limit"`
            CardIsin string  `json:"card_isin"`
            FyndVpa string  `json:"fynd_vpa"`
            ExpMonth float64  `json:"exp_month"`
         
    }
    
    // RootPaymentMode used by Payment
    type RootPaymentMode struct {

        
            Name string  `json:"name"`
            SaveCard bool  `json:"save_card"`
            AddCardEnabled bool  `json:"add_card_enabled"`
            DisplayName string  `json:"display_name"`
            DisplayPriority float64  `json:"display_priority"`
            List []PaymentModeList  `json:"list"`
            IsPayByCardPl bool  `json:"is_pay_by_card_pl"`
            AggregatorName string  `json:"aggregator_name"`
            AnonymousEnable bool  `json:"anonymous_enable"`
         
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
            Customers map[string]interface{}  `json:"customers"`
            UniqueTransferNo map[string]interface{}  `json:"unique_transfer_no"`
            IsActive bool  `json:"is_active"`
            TransferType string  `json:"transfer_type"`
            IsDefault bool  `json:"is_default"`
            MoreAttributes map[string]interface{}  `json:"more_attributes"`
         
    }
    
    // PayoutBankDetails used by Payment
    type PayoutBankDetails struct {

        
            AccountNo string  `json:"account_no"`
            City string  `json:"city"`
            Pincode float64  `json:"pincode"`
            Country string  `json:"country"`
            IfscCode string  `json:"ifsc_code"`
            AccountHolder string  `json:"account_holder"`
            BranchName string  `json:"branch_name"`
            BankName string  `json:"bank_name"`
            AccountType string  `json:"account_type"`
            State string  `json:"state"`
         
    }
    
    // PayoutRequest used by Payment
    type PayoutRequest struct {

        
            UniqueExternalID string  `json:"unique_external_id"`
            IsActive bool  `json:"is_active"`
            Users map[string]interface{}  `json:"users"`
            TransferType string  `json:"transfer_type"`
            BankDetails PayoutBankDetails  `json:"bank_details"`
            Aggregator string  `json:"aggregator"`
         
    }
    
    // PayoutResponse used by Payment
    type PayoutResponse struct {

        
            Success bool  `json:"success"`
            UniqueTransferNo string  `json:"unique_transfer_no"`
            Payouts map[string]interface{}  `json:"payouts"`
            Created bool  `json:"created"`
            IsActive bool  `json:"is_active"`
            Users map[string]interface{}  `json:"users"`
            TransferType string  `json:"transfer_type"`
            PaymentStatus string  `json:"payment_status"`
            BankDetails map[string]interface{}  `json:"bank_details"`
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

        
            Success bool  `json:"success"`
            Data map[string]interface{}  `json:"data"`
         
    }
    
    // RefundAccountResponse used by Payment
    type RefundAccountResponse struct {

        
            Message string  `json:"message"`
            IsVerifiedFlag bool  `json:"is_verified_flag"`
            Success bool  `json:"success"`
            Data map[string]interface{}  `json:"data"`
         
    }
    
    // NotFoundResourceError used by Payment
    type NotFoundResourceError struct {

        
            Code string  `json:"code"`
            Description string  `json:"description"`
            Success bool  `json:"success"`
         
    }
    
    // BankDetailsForOTP used by Payment
    type BankDetailsForOTP struct {

        
            AccountNo string  `json:"account_no"`
            IfscCode string  `json:"ifsc_code"`
            AccountHolder string  `json:"account_holder"`
            BranchName string  `json:"branch_name"`
            BankName string  `json:"bank_name"`
         
    }
    
    // AddBeneficiaryDetailsOTPRequest used by Payment
    type AddBeneficiaryDetailsOTPRequest struct {

        
            Details BankDetailsForOTP  `json:"details"`
            OrderID string  `json:"order_id"`
         
    }
    
    // IfscCodeResponse used by Payment
    type IfscCodeResponse struct {

        
            Success bool  `json:"success"`
            BranchName string  `json:"branch_name"`
            BankName string  `json:"bank_name"`
         
    }
    
    // OrderBeneficiaryDetails used by Payment
    type OrderBeneficiaryDetails struct {

        
            CreatedOn string  `json:"created_on"`
            Comment string  `json:"comment"`
            Title string  `json:"title"`
            DisplayName string  `json:"display_name"`
            IsActive bool  `json:"is_active"`
            Subtitle string  `json:"subtitle"`
            TransferMode string  `json:"transfer_mode"`
            BankName string  `json:"bank_name"`
            Email string  `json:"email"`
            IfscCode string  `json:"ifsc_code"`
            BeneficiaryID string  `json:"beneficiary_id"`
            ModifiedOn string  `json:"modified_on"`
            AccountHolder string  `json:"account_holder"`
            BranchName string  `json:"branch_name"`
            Mobile string  `json:"mobile"`
            AccountNo string  `json:"account_no"`
            ID float64  `json:"id"`
            DelightsUserName string  `json:"delights_user_name"`
            Address string  `json:"address"`
         
    }
    
    // OrderBeneficiaryResponse used by Payment
    type OrderBeneficiaryResponse struct {

        
            ShowBeneficiaryDetails bool  `json:"show_beneficiary_details"`
            Beneficiaries []OrderBeneficiaryDetails  `json:"beneficiaries"`
         
    }
    
    // MultiTenderPaymentMeta used by Payment
    type MultiTenderPaymentMeta struct {

        
            CurrentStatus string  `json:"current_status"`
            PaymentID string  `json:"payment_id"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            OrderID string  `json:"order_id"`
            PaymentGateway string  `json:"payment_gateway"`
         
    }
    
    // MultiTenderPaymentMethod used by Payment
    type MultiTenderPaymentMethod struct {

        
            Mode string  `json:"mode"`
            Meta MultiTenderPaymentMeta  `json:"meta"`
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

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
            OrderID string  `json:"order_id"`
         
    }
    
    // CODdata used by Payment
    type CODdata struct {

        
            UserID string  `json:"user_id"`
            Limit float64  `json:"limit"`
            IsActive bool  `json:"is_active"`
            RemainingLimit float64  `json:"remaining_limit"`
            Usages float64  `json:"usages"`
         
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
    

    
    // ShipmentItemFulFillingStore used by Order
    type ShipmentItemFulFillingStore struct {

        
            Code string  `json:"code"`
            ID string  `json:"id"`
         
    }
    
    // UserDataInfo used by Order
    type UserDataInfo struct {

        
            FirstName string  `json:"first_name"`
            Mobile string  `json:"mobile"`
            LastName string  `json:"last_name"`
            AvisUserID string  `json:"avis_user_id"`
            IsAnonymousUser bool  `json:"is_anonymous_user"`
            Gender string  `json:"gender"`
            Name string  `json:"name"`
            Email string  `json:"email"`
            UID float64  `json:"uid"`
         
    }
    
    // PaymentModeInfo used by Order
    type PaymentModeInfo struct {

        
            Logo string  `json:"logo"`
            Type string  `json:"type"`
         
    }
    
    // PlatformItem used by Order
    type PlatformItem struct {

        
            DepartmentID float64  `json:"department_id"`
            Image []string  `json:"image"`
            CanReturn bool  `json:"can_return"`
            L3Category float64  `json:"l3_category"`
            L1Category []string  `json:"l1_category"`
            Code string  `json:"code"`
            CanCancel bool  `json:"can_cancel"`
            L3CategoryName string  `json:"l3_category_name"`
            Size string  `json:"size"`
            Color string  `json:"color"`
            Images []string  `json:"images"`
            Name string  `json:"name"`
            ID float64  `json:"id"`
         
    }
    
    // GSTDetailsData used by Order
    type GSTDetailsData struct {

        
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            GstFee float64  `json:"gst_fee"`
            GstinCode string  `json:"gstin_code"`
            ValueOfGood float64  `json:"value_of_good"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
         
    }
    
    // Prices used by Order
    type Prices struct {

        
            AmountPaid float64  `json:"amount_paid"`
            PriceEffective float64  `json:"price_effective"`
            AmountPaidRoundoff float64  `json:"amount_paid_roundoff"`
            Discount float64  `json:"discount"`
            Cashback float64  `json:"cashback"`
            FyndCredits float64  `json:"fynd_credits"`
            CashbackApplied float64  `json:"cashback_applied"`
            RefundCredit float64  `json:"refund_credit"`
            RefundAmount float64  `json:"refund_amount"`
            CodCharges float64  `json:"cod_charges"`
            PriceMarked float64  `json:"price_marked"`
            CouponValue float64  `json:"coupon_value"`
            ValueOfGood float64  `json:"value_of_good"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            TransferPrice float64  `json:"transfer_price"`
            PromotionEffectiveDiscount float64  `json:"promotion_effective_discount"`
            DeliveryCharge float64  `json:"delivery_charge"`
         
    }
    
    // BagUnit used by Order
    type BagUnit struct {

        
            Item PlatformItem  `json:"item"`
            CanReturn bool  `json:"can_return"`
            BagID float64  `json:"bag_id"`
            CanCancel bool  `json:"can_cancel"`
            OrderingChannel string  `json:"ordering_channel"`
            TotalShipmentBags float64  `json:"total_shipment_bags"`
            ItemQuantity float64  `json:"item_quantity"`
            ShipmentID string  `json:"shipment_id"`
            Gst GSTDetailsData  `json:"gst"`
            Prices Prices  `json:"prices"`
            Status map[string]interface{}  `json:"status"`
         
    }
    
    // ShipmentStatus used by Order
    type ShipmentStatus struct {

        
            ActualStatus string  `json:"actual_status"`
            OpsStatus string  `json:"ops_status"`
            HexCode string  `json:"hex_code"`
            Title string  `json:"title"`
            Status string  `json:"status"`
         
    }
    
    // ShipmentItem used by Order
    type ShipmentItem struct {

        
            ShipmentCreatedAt string  `json:"shipment_created_at"`
            FulfillingStore ShipmentItemFulFillingStore  `json:"fulfilling_store"`
            Channel map[string]interface{}  `json:"channel"`
            TotalBagsCount float64  `json:"total_bags_count"`
            User UserDataInfo  `json:"user"`
            Sla map[string]interface{}  `json:"sla"`
            PaymentModeInfo PaymentModeInfo  `json:"payment_mode_info"`
            PaymentMethods map[string]interface{}  `json:"payment_methods"`
            Application map[string]interface{}  `json:"application"`
            Bags []BagUnit  `json:"bags"`
            CreatedAt string  `json:"created_at"`
            ShipmentID string  `json:"shipment_id"`
            TotalShipmentsInOrder float64  `json:"total_shipments_in_order"`
            FulfillingCentre string  `json:"fulfilling_centre"`
            ID string  `json:"id"`
            ShipmentStatus ShipmentStatus  `json:"shipment_status"`
            Prices Prices  `json:"prices"`
         
    }
    
    // FilterInfoOption used by Order
    type FilterInfoOption struct {

        
            Text string  `json:"text"`
            Name string  `json:"name"`
            Value string  `json:"value"`
         
    }
    
    // FiltersInfo used by Order
    type FiltersInfo struct {

        
            Options []FilterInfoOption  `json:"options"`
            Text string  `json:"text"`
            Type string  `json:"type"`
            Value string  `json:"value"`
         
    }
    
    // ShipmentInternalPlatformViewResponse used by Order
    type ShipmentInternalPlatformViewResponse struct {

        
            Items []ShipmentItem  `json:"items"`
            Page map[string]interface{}  `json:"page"`
            AppliedFilters map[string]interface{}  `json:"applied_filters"`
            Filters []FiltersInfo  `json:"filters"`
         
    }
    
    // Error used by Order
    type Error struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // TrackingList used by Order
    type TrackingList struct {

        
            IsPassed bool  `json:"is_passed"`
            IsCurrent bool  `json:"is_current"`
            Text string  `json:"text"`
            Time string  `json:"time"`
            Status string  `json:"status"`
         
    }
    
    // BagStateMapper used by Order
    type BagStateMapper struct {

        
            DisplayName string  `json:"display_name"`
            NotifyCustomer bool  `json:"notify_customer"`
            StateType string  `json:"state_type"`
            BsID float64  `json:"bs_id"`
            JourneyType string  `json:"journey_type"`
            AppFacing bool  `json:"app_facing"`
            Name string  `json:"name"`
            AppStateName string  `json:"app_state_name"`
            IsActive bool  `json:"is_active"`
            AppDisplayName string  `json:"app_display_name"`
         
    }
    
    // BagStatusHistory used by Order
    type BagStatusHistory struct {

        
            StoreID float64  `json:"store_id"`
            Forward bool  `json:"forward"`
            DisplayName string  `json:"display_name"`
            KafkaSync bool  `json:"kafka_sync"`
            BagStateMapper BagStateMapper  `json:"bag_state_mapper"`
            StateType string  `json:"state_type"`
            BagID float64  `json:"bag_id"`
            CreatedAt string  `json:"created_at"`
            ShipmentID string  `json:"shipment_id"`
            BshID float64  `json:"bsh_id"`
            StateID float64  `json:"state_id"`
            UpdatedAt string  `json:"updated_at"`
            DeliveryPartnerID float64  `json:"delivery_partner_id"`
            DeliveryAwbNumber string  `json:"delivery_awb_number"`
            Reasons []map[string]interface{}  `json:"reasons"`
            AppDisplayName string  `json:"app_display_name"`
            Status string  `json:"status"`
         
    }
    
    // InvoiceInfo used by Order
    type InvoiceInfo struct {

        
            LabelURL string  `json:"label_url"`
            CreditNoteID string  `json:"credit_note_id"`
            UpdatedDate string  `json:"updated_date"`
            StoreInvoiceID string  `json:"store_invoice_id"`
            InvoiceURL string  `json:"invoice_url"`
         
    }
    
    // ShipmentStatusData used by Order
    type ShipmentStatusData struct {

        
            BagList []string  `json:"bag_list"`
            CreatedAt string  `json:"created_at"`
            ShipmentID string  `json:"shipment_id"`
            ID float64  `json:"id"`
            Status string  `json:"status"`
         
    }
    
    // DPDetailsData used by Order
    type DPDetailsData struct {

        
            EwayBillID string  `json:"eway_bill_id"`
            GstTag string  `json:"gst_tag"`
            Pincode string  `json:"pincode"`
            TrackURL string  `json:"track_url"`
            Country string  `json:"country"`
            Name string  `json:"name"`
            AwbNo string  `json:"awb_no"`
            ID float64  `json:"id"`
         
    }
    
    // FulfillingStore used by Order
    type FulfillingStore struct {

        
            Meta map[string]interface{}  `json:"meta"`
            Pincode string  `json:"pincode"`
            Code string  `json:"code"`
            City string  `json:"city"`
            Country string  `json:"country"`
            Address string  `json:"address"`
            ContactPerson string  `json:"contact_person"`
            FulfillmentChannel string  `json:"fulfillment_channel"`
            StoreName string  `json:"store_name"`
            ID float64  `json:"id"`
            State string  `json:"state"`
            Phone string  `json:"phone"`
         
    }
    
    // OrderDetailsData used by Order
    type OrderDetailsData struct {

        
            Source string  `json:"source"`
            OrderingChannelLogo map[string]interface{}  `json:"ordering_channel_logo"`
            OrderingChannel string  `json:"ordering_channel"`
            TaxDetails map[string]interface{}  `json:"tax_details"`
            FyndOrderID string  `json:"fynd_order_id"`
            CodCharges string  `json:"cod_charges"`
            OrderDate string  `json:"order_date"`
            AffiliateID string  `json:"affiliate_id"`
            OrderValue string  `json:"order_value"`
         
    }
    
    // UserDetailsData used by Order
    type UserDetailsData struct {

        
            AddressType string  `json:"address_type"`
            Pincode string  `json:"pincode"`
            Area string  `json:"area"`
            City string  `json:"city"`
            Country string  `json:"country"`
            Address string  `json:"address"`
            Landmark string  `json:"landmark"`
            Name string  `json:"name"`
            Address1 string  `json:"address1"`
            Email string  `json:"email"`
            State string  `json:"state"`
            Phone string  `json:"phone"`
         
    }
    
    // ReplacementDetails used by Order
    type ReplacementDetails struct {

        
            OriginalAffiliateOrderID string  `json:"original_affiliate_order_id"`
            ReplacementType string  `json:"replacement_type"`
         
    }
    
    // AffiliateMeta used by Order
    type AffiliateMeta struct {

        
            CouponCode string  `json:"coupon_code"`
            ChannelOrderID string  `json:"channel_order_id"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            ChannelShipmentID string  `json:"channel_shipment_id"`
            OrderItemID string  `json:"order_item_id"`
            BoxType string  `json:"box_type"`
            Quantity float64  `json:"quantity"`
            DueDate string  `json:"due_date"`
            IsPriority bool  `json:"is_priority"`
            SizeLevelTotalQty float64  `json:"size_level_total_qty"`
            MarketplaceInvoiceID string  `json:"marketplace_invoice_id"`
            ReplacementDetails ReplacementDetails  `json:"replacement_details"`
            EmployeeDiscount float64  `json:"employee_discount"`
         
    }
    
    // PDFLinks used by Order
    type PDFLinks struct {

        
            LabelA4 string  `json:"label_a4"`
            InvoiceA6 string  `json:"invoice_a6"`
            B2b string  `json:"b2b"`
            InvoiceType string  `json:"invoice_type"`
            Invoice string  `json:"invoice"`
            LabelPos string  `json:"label_pos"`
            LabelType string  `json:"label_type"`
            LabelA6 string  `json:"label_a6"`
            PoInvoice string  `json:"po_invoice"`
            Label string  `json:"label"`
            InvoicePos string  `json:"invoice_pos"`
            CreditNoteURL string  `json:"credit_note_url"`
            DeliveryChallanA4 string  `json:"delivery_challan_a4"`
            InvoiceA4 string  `json:"invoice_a4"`
         
    }
    
    // BuyerDetails used by Order
    type BuyerDetails struct {

        
            Gstin string  `json:"gstin"`
            Pincode float64  `json:"pincode"`
            City string  `json:"city"`
            Address string  `json:"address"`
            AjioSiteID string  `json:"ajio_site_id"`
            Name string  `json:"name"`
            State string  `json:"state"`
         
    }
    
    // DebugInfo used by Order
    type DebugInfo struct {

        
            StormbreakerUUID string  `json:"stormbreaker_uuid"`
         
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
    
    // ShipmentTimeStamp used by Order
    type ShipmentTimeStamp struct {

        
            TMin string  `json:"t_min"`
            TMax string  `json:"t_max"`
         
    }
    
    // LockData used by Order
    type LockData struct {

        
            Mto bool  `json:"mto"`
            Locked bool  `json:"locked"`
            LockMessage string  `json:"lock_message"`
         
    }
    
    // ShipmentMeta used by Order
    type ShipmentMeta struct {

        
            Weight float64  `json:"weight"`
            PoNumber string  `json:"po_number"`
            ReturnDetails map[string]interface{}  `json:"return_details"`
            ReturnStoreNode float64  `json:"return_store_node"`
            B2bBuyerDetails BuyerDetails  `json:"b2b_buyer_details"`
            DebugInfo DebugInfo  `json:"debug_info"`
            SameStoreAvailable bool  `json:"same_store_available"`
            BoxType string  `json:"box_type"`
            ShipmentVolumetricWeight float64  `json:"shipment_volumetric_weight"`
            ReturnAwbNumber string  `json:"return_awb_number"`
            DpSortKey string  `json:"dp_sort_key"`
            BagWeight map[string]interface{}  `json:"bag_weight"`
            DpID string  `json:"dp_id"`
            ShipmentWeight float64  `json:"shipment_weight"`
            AwbNumber string  `json:"awb_number"`
            DpOptions map[string]interface{}  `json:"dp_options"`
            ReturnAffiliateShipmentID string  `json:"return_affiliate_shipment_id"`
            AutoTriggerDpAssignmentAcf bool  `json:"auto_trigger_dp_assignment_acf"`
            ForwardAffiliateShipmentID string  `json:"forward_affiliate_shipment_id"`
            Formatted Formatted  `json:"formatted"`
            EinvoiceInfo EinvoiceInfo  `json:"einvoice_info"`
            ReturnAffiliateOrderID string  `json:"return_affiliate_order_id"`
            DpName string  `json:"dp_name"`
            DueDate string  `json:"due_date"`
            StoreInvoiceUpdatedDate string  `json:"store_invoice_updated_date"`
            PackagingName string  `json:"packaging_name"`
            B2cBuyerDetails map[string]interface{}  `json:"b2c_buyer_details"`
            AssignDpFromSb bool  `json:"assign_dp_from_sb"`
            ForwardAffiliateOrderID string  `json:"forward_affiliate_order_id"`
            Timestamp ShipmentTimeStamp  `json:"timestamp"`
            MarketplaceStoreID string  `json:"marketplace_store_id"`
            OrderType string  `json:"order_type"`
            External map[string]interface{}  `json:"external"`
            FulfilmentPriorityText string  `json:"fulfilment_priority_text"`
            EwaybillInfo map[string]interface{}  `json:"ewaybill_info"`
            LockData LockData  `json:"lock_data"`
         
    }
    
    // AffiliateDetails used by Order
    type AffiliateDetails struct {

        
            AffiliateStoreID string  `json:"affiliate_store_id"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            AffiliateMeta AffiliateMeta  `json:"affiliate_meta"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            PdfLinks PDFLinks  `json:"pdf_links"`
            AdID string  `json:"ad_id"`
            CompanyAffiliateTag string  `json:"company_affiliate_tag"`
            AffiliateID string  `json:"affiliate_id"`
            ShipmentMeta ShipmentMeta  `json:"shipment_meta"`
         
    }
    
    // OrderingStoreDetails used by Order
    type OrderingStoreDetails struct {

        
            Meta map[string]interface{}  `json:"meta"`
            OrderingStoreID float64  `json:"ordering_store_id"`
            Pincode string  `json:"pincode"`
            Code string  `json:"code"`
            City string  `json:"city"`
            Country string  `json:"country"`
            Address string  `json:"address"`
            ContactPerson string  `json:"contact_person"`
            StoreName string  `json:"store_name"`
            State string  `json:"state"`
            Phone string  `json:"phone"`
         
    }
    
    // ShipmentPayments used by Order
    type ShipmentPayments struct {

        
            Mode string  `json:"mode"`
            Logo string  `json:"logo"`
            Source string  `json:"source"`
         
    }
    
    // Dimensions used by Order
    type Dimensions struct {

        
            Height float64  `json:"height"`
            IsDefault bool  `json:"is_default"`
            Width float64  `json:"width"`
            Unit string  `json:"unit"`
            Length float64  `json:"length"`
         
    }
    
    // Meta used by Order
    type Meta struct {

        
            Dimension Dimensions  `json:"dimension"`
         
    }
    
    // PlatformDeliveryAddress used by Order
    type PlatformDeliveryAddress struct {

        
            AddressType string  `json:"address_type"`
            Pincode string  `json:"pincode"`
            Area string  `json:"area"`
            City string  `json:"city"`
            Country string  `json:"country"`
            Address2 string  `json:"address2"`
            Version string  `json:"version"`
            Landmark string  `json:"landmark"`
            Latitude float64  `json:"latitude"`
            Longitude float64  `json:"longitude"`
            CreatedAt string  `json:"created_at"`
            ContactPerson string  `json:"contact_person"`
            Address1 string  `json:"address1"`
            AddressCategory string  `json:"address_category"`
            UpdatedAt string  `json:"updated_at"`
            Email string  `json:"email"`
            State string  `json:"state"`
            Phone string  `json:"phone"`
         
    }
    
    // Identifier used by Order
    type Identifier struct {

        
            Ean string  `json:"ean"`
         
    }
    
    // FinancialBreakup used by Order
    type FinancialBreakup struct {

        
            AmountPaid float64  `json:"amount_paid"`
            AmountPaidRoundoff float64  `json:"amount_paid_roundoff"`
            TotalUnits float64  `json:"total_units"`
            PriceMarked float64  `json:"price_marked"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            PromotionEffectiveDiscount float64  `json:"promotion_effective_discount"`
            GstTag string  `json:"gst_tag"`
            PriceEffective float64  `json:"price_effective"`
            FyndCredits float64  `json:"fynd_credits"`
            RefundCredit float64  `json:"refund_credit"`
            Size string  `json:"size"`
            Identifiers Identifier  `json:"identifiers"`
            CodCharges float64  `json:"cod_charges"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            Discount float64  `json:"discount"`
            Cashback float64  `json:"cashback"`
            AddedToFyndCash bool  `json:"added_to_fynd_cash"`
            CashbackApplied float64  `json:"cashback_applied"`
            GstFee float64  `json:"gst_fee"`
            ValueOfGood float64  `json:"value_of_good"`
            CouponValue float64  `json:"coupon_value"`
            HsnCode string  `json:"hsn_code"`
            ItemName string  `json:"item_name"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            TransferPrice float64  `json:"transfer_price"`
            DeliveryCharge float64  `json:"delivery_charge"`
         
    }
    
    // OrderBagArticle used by Order
    type OrderBagArticle struct {

        
            Identifiers map[string]interface{}  `json:"identifiers"`
            UID string  `json:"uid"`
            ReturnConfig map[string]interface{}  `json:"return_config"`
         
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
    
    // DiscountRules used by Order
    type DiscountRules struct {

        
            Type string  `json:"type"`
            Value float64  `json:"value"`
         
    }
    
    // AppliedPromos used by Order
    type AppliedPromos struct {

        
            PromotionName string  `json:"promotion_name"`
            Amount float64  `json:"amount"`
            BuyRules []BuyRules  `json:"buy_rules"`
            PromoID string  `json:"promo_id"`
            MrpPromotion bool  `json:"mrp_promotion"`
            PromotionType string  `json:"promotion_type"`
            ArticleQuantity float64  `json:"article_quantity"`
            DiscountRules []DiscountRules  `json:"discount_rules"`
         
    }
    
    // BagConfigs used by Order
    type BagConfigs struct {

        
            IsReturnable bool  `json:"is_returnable"`
            EnableTracking bool  `json:"enable_tracking"`
            CanBeCancelled bool  `json:"can_be_cancelled"`
            IsActive bool  `json:"is_active"`
            IsCustomerReturnAllowed bool  `json:"is_customer_return_allowed"`
            AllowForceReturn bool  `json:"allow_force_return"`
         
    }
    
    // OrderBrandName used by Order
    type OrderBrandName struct {

        
            Company string  `json:"company"`
            ModifiedOn string  `json:"modified_on"`
            Logo string  `json:"logo"`
            CreatedOn string  `json:"created_on"`
            BrandName string  `json:"brand_name"`
            ID float64  `json:"id"`
         
    }
    
    // CurrentStatus used by Order
    type CurrentStatus struct {

        
            StoreID float64  `json:"store_id"`
            KafkaSync bool  `json:"kafka_sync"`
            Status string  `json:"status"`
            BagStateMapper BagStateMapper  `json:"bag_state_mapper"`
            StateType string  `json:"state_type"`
            BagID float64  `json:"bag_id"`
            CreatedAt string  `json:"created_at"`
            ShipmentID string  `json:"shipment_id"`
            StateID float64  `json:"state_id"`
            UpdatedAt string  `json:"updated_at"`
            DeliveryPartnerID float64  `json:"delivery_partner_id"`
            CurrentStatusID float64  `json:"current_status_id"`
            DeliveryAwbNumber string  `json:"delivery_awb_number"`
         
    }
    
    // BagGST used by Order
    type BagGST struct {

        
            GstTag string  `json:"gst_tag"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            IsDefaultHsnCode bool  `json:"is_default_hsn_code"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            GstFee float64  `json:"gst_fee"`
            GstinCode string  `json:"gstin_code"`
            ValueOfGood float64  `json:"value_of_good"`
            HsnCode string  `json:"hsn_code"`
         
    }
    
    // OrderBags used by Order
    type OrderBags struct {

        
            DeliveryAddress PlatformDeliveryAddress  `json:"delivery_address"`
            DisplayName string  `json:"display_name"`
            ParentPromoBags map[string]interface{}  `json:"parent_promo_bags"`
            BagID float64  `json:"bag_id"`
            FinancialBreakup FinancialBreakup  `json:"financial_breakup"`
            Quantity float64  `json:"quantity"`
            Identifier string  `json:"identifier"`
            Item PlatformItem  `json:"item"`
            CanReturn bool  `json:"can_return"`
            SellerIdentifier string  `json:"seller_identifier"`
            Article OrderBagArticle  `json:"article"`
            EntityType string  `json:"entity_type"`
            CanCancel bool  `json:"can_cancel"`
            AppliedPromos []AppliedPromos  `json:"applied_promos"`
            BagConfigs BagConfigs  `json:"bag_configs"`
            Brand OrderBrandName  `json:"brand"`
            LineNumber float64  `json:"line_number"`
            CurrentStatus CurrentStatus  `json:"current_status"`
            GstDetails BagGST  `json:"gst_details"`
            Prices Prices  `json:"prices"`
         
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

        
            CompanyCin string  `json:"company_cin"`
            Address map[string]interface{}  `json:"address"`
            CompanyContact ContactDetails  `json:"company_contact"`
            CompanyGst string  `json:"company_gst"`
            CompanyID float64  `json:"company_id"`
            CompanyName string  `json:"company_name"`
         
    }
    
    // PlatformShipment used by Order
    type PlatformShipment struct {

        
            TrackingList []TrackingList  `json:"tracking_list"`
            TotalBags float64  `json:"total_bags"`
            BagStatusHistory []BagStatusHistory  `json:"bag_status_history"`
            PriorityText string  `json:"priority_text"`
            User UserDataInfo  `json:"user"`
            InvoiceID string  `json:"invoice_id"`
            Invoice InvoiceInfo  `json:"invoice"`
            FulfilmentPriority float64  `json:"fulfilment_priority"`
            DeliverySlot map[string]interface{}  `json:"delivery_slot"`
            PlatformLogo string  `json:"platform_logo"`
            Coupon map[string]interface{}  `json:"coupon"`
            Status ShipmentStatusData  `json:"status"`
            DpDetails DPDetailsData  `json:"dp_details"`
            PackagingType string  `json:"packaging_type"`
            TotalItems float64  `json:"total_items"`
            LockStatus bool  `json:"lock_status"`
            ShipmentStatus string  `json:"shipment_status"`
            FulfillingStore FulfillingStore  `json:"fulfilling_store"`
            Order OrderDetailsData  `json:"order"`
            DeliveryDetails UserDetailsData  `json:"delivery_details"`
            ForwardShipmentID string  `json:"forward_shipment_id"`
            AffiliateDetails AffiliateDetails  `json:"affiliate_details"`
            JourneyType string  `json:"journey_type"`
            PickedDate string  `json:"picked_date"`
            PaymentMode string  `json:"payment_mode"`
            OrderingStore OrderingStoreDetails  `json:"ordering_store"`
            ShipmentQuantity float64  `json:"shipment_quantity"`
            Payments ShipmentPayments  `json:"payments"`
            Meta Meta  `json:"meta"`
            BillingDetails UserDetailsData  `json:"billing_details"`
            EnableDpTracking bool  `json:"enable_dp_tracking"`
            PaymentMethods map[string]interface{}  `json:"payment_methods"`
            OperationalStatus string  `json:"operational_status"`
            Bags []OrderBags  `json:"bags"`
            ShipmentImages []string  `json:"shipment_images"`
            ShipmentID string  `json:"shipment_id"`
            Vertical string  `json:"vertical"`
            CompanyDetails CompanyDetails  `json:"company_details"`
            UserAgent string  `json:"user_agent"`
            GstDetails GSTDetailsData  `json:"gst_details"`
            CustomMeta []map[string]interface{}  `json:"custom_meta"`
            Prices Prices  `json:"prices"`
         
    }
    
    // ShipmentInfoResponse used by Order
    type ShipmentInfoResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
            Shipments []PlatformShipment  `json:"shipments"`
         
    }
    
    // BillingStaffDetails used by Order
    type BillingStaffDetails struct {

        
            FirstName string  `json:"first_name"`
            LastName string  `json:"last_name"`
            User string  `json:"user"`
            EmployeeCode string  `json:"employee_code"`
            StaffID float64  `json:"staff_id"`
         
    }
    
    // PlatformUserDetails used by Order
    type PlatformUserDetails struct {

        
            PlatformUserLastName string  `json:"platform_user_last_name"`
            PlatformUserFirstName string  `json:"platform_user_first_name"`
            PlatformUserID string  `json:"platform_user_id"`
            PlatformUserEmployeeCode string  `json:"platform_user_employee_code"`
         
    }
    
    // TransactionData used by Order
    type TransactionData struct {

        
            AmountPaid string  `json:"amount_paid"`
            Currency string  `json:"currency"`
            TransactionID string  `json:"transaction_id"`
            PaymentID string  `json:"payment_id"`
            TerminalID string  `json:"terminal_id"`
            UniqueReferenceNumber string  `json:"unique_reference_number"`
            Entity string  `json:"entity"`
            Status string  `json:"status"`
         
    }
    
    // OrderMeta used by Order
    type OrderMeta struct {

        
            BillingStaffDetails BillingStaffDetails  `json:"billing_staff_details"`
            CompanyLogo string  `json:"company_logo"`
            CartID float64  `json:"cart_id"`
            PlatformUserDetails PlatformUserDetails  `json:"platform_user_details"`
            PaymentType string  `json:"payment_type"`
            CurrencySymbol string  `json:"currency_symbol"`
            Staff map[string]interface{}  `json:"staff"`
            MongoCartID float64  `json:"mongo_cart_id"`
            Comment string  `json:"comment"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            TransactionData TransactionData  `json:"transaction_data"`
            OrderChildEntities []string  `json:"order_child_entities"`
            OrderingStore float64  `json:"ordering_store"`
            OrderPlatform string  `json:"order_platform"`
            Files []map[string]interface{}  `json:"files"`
            OrderType string  `json:"order_type"`
            CustomerNote string  `json:"customer_note"`
            OrderTags []map[string]interface{}  `json:"order_tags"`
            EmployeeID float64  `json:"employee_id"`
         
    }
    
    // TaxDetails used by Order
    type TaxDetails struct {

        
            PanNo string  `json:"pan_no"`
            Gstin string  `json:"gstin"`
         
    }
    
    // OrderDict used by Order
    type OrderDict struct {

        
            Meta OrderMeta  `json:"meta"`
            PaymentMethods map[string]interface{}  `json:"payment_methods"`
            TaxDetails TaxDetails  `json:"tax_details"`
            FyndOrderID string  `json:"fynd_order_id"`
            OrderDate string  `json:"order_date"`
            Prices Prices  `json:"prices"`
         
    }
    
    // ShipmentDetailsResponse used by Order
    type ShipmentDetailsResponse struct {

        
            Success bool  `json:"success"`
            Shipments []PlatformShipment  `json:"shipments"`
            Order OrderDict  `json:"order"`
         
    }
    
    // SubLane used by Order
    type SubLane struct {

        
            Actions []map[string]interface{}  `json:"actions"`
            TotalItems float64  `json:"total_items"`
            Value string  `json:"value"`
            Text string  `json:"text"`
            Index float64  `json:"index"`
         
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
            Display string  `json:"display"`
            Name string  `json:"name"`
         
    }
    
    // PlatformOrderItems used by Order
    type PlatformOrderItems struct {

        
            Channel PlatformChannel  `json:"channel"`
            Meta map[string]interface{}  `json:"meta"`
            UserInfo UserDataInfo  `json:"user_info"`
            BreakupValues []PlatformBreakupValues  `json:"breakup_values"`
            TotalOrderValue float64  `json:"total_order_value"`
            PaymentMode string  `json:"payment_mode"`
            OrderCreatedTime string  `json:"order_created_time"`
            OrderValue float64  `json:"order_value"`
            Shipments []PlatformShipment  `json:"shipments"`
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
    
    // OrderListingResponse used by Order
    type OrderListingResponse struct {

        
            Message string  `json:"message"`
            Items []PlatformOrderItems  `json:"items"`
            Lane string  `json:"lane"`
            Page Page  `json:"page"`
            TotalCount float64  `json:"total_count"`
            Success bool  `json:"success"`
         
    }
    
    // PlatformTrack used by Order
    type PlatformTrack struct {

        
            Status string  `json:"status"`
            AccountName string  `json:"account_name"`
            Meta map[string]interface{}  `json:"meta"`
            Reason string  `json:"reason"`
            UpdatedTime string  `json:"updated_time"`
            RawStatus string  `json:"raw_status"`
            Awb string  `json:"awb"`
            UpdatedAt string  `json:"updated_at"`
            ShipmentType string  `json:"shipment_type"`
            LastLocationRecievedAt string  `json:"last_location_recieved_at"`
         
    }
    
    // PlatformShipmentTrack used by Order
    type PlatformShipmentTrack struct {

        
            Meta map[string]interface{}  `json:"meta"`
            Results []PlatformTrack  `json:"results"`
         
    }
    
    // FiltersResponse used by Order
    type FiltersResponse struct {

        
            Advance []map[string]interface{}  `json:"advance"`
         
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

        
            ContentType string  `json:"content_type"`
            FilePath string  `json:"file_path"`
            Size float64  `json:"size"`
            Upload FileUploadResponse  `json:"upload"`
            Cdn URL  `json:"cdn"`
            FileName string  `json:"file_name"`
            Operation string  `json:"operation"`
            Method string  `json:"method"`
            Namespace string  `json:"namespace"`
            Tags []string  `json:"tags"`
         
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

        
            DisplayName string  `json:"display_name"`
            ID float64  `json:"id"`
         
    }
    
    // Reason used by Order
    type Reason struct {

        
            QcType []string  `json:"qc_type"`
            DisplayName string  `json:"display_name"`
            ID float64  `json:"id"`
            QuestionSet []QuestionSet  `json:"question_set"`
         
    }
    
    // PlatformShipmentReasonsResponse used by Order
    type PlatformShipmentReasonsResponse struct {

        
            Success bool  `json:"success"`
            Reasons []Reason  `json:"reasons"`
         
    }
    
    // AffiliateBagDetails used by Order
    type AffiliateBagDetails struct {

        
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            AffiliateMeta AffiliateMeta  `json:"affiliate_meta"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            EmployeeDiscount float64  `json:"employee_discount"`
         
    }
    
    // BagReturnableCancelableStatus used by Order
    type BagReturnableCancelableStatus struct {

        
            IsReturnable bool  `json:"is_returnable"`
            EnableTracking bool  `json:"enable_tracking"`
            CanBeCancelled bool  `json:"can_be_cancelled"`
            IsActive bool  `json:"is_active"`
            IsCustomerReturnAllowed bool  `json:"is_customer_return_allowed"`
         
    }
    
    // Attributes used by Order
    type Attributes struct {

        
            PrimaryColor string  `json:"primary_color"`
            Gender []string  `json:"gender"`
            Essential string  `json:"essential"`
            BrandName string  `json:"brand_name"`
            MarketerName string  `json:"marketer_name"`
            MarketerAddress string  `json:"marketer_address"`
            PrimaryMaterial string  `json:"primary_material"`
            PrimaryColorHex string  `json:"primary_color_hex"`
            Name string  `json:"name"`
         
    }
    
    // Item used by Order
    type Item struct {

        
            ItemID float64  `json:"item_id"`
            Color string  `json:"color"`
            Attributes Attributes  `json:"attributes"`
            WebstoreProductURL string  `json:"webstore_product_url"`
            BranchURL string  `json:"branch_url"`
            CanReturn bool  `json:"can_return"`
            L3Category float64  `json:"l3_category"`
            L1Category []string  `json:"l1_category"`
            Code string  `json:"code"`
            BrandID float64  `json:"brand_id"`
            Size string  `json:"size"`
            Name string  `json:"name"`
            LastUpdatedAt string  `json:"last_updated_at"`
            Image []string  `json:"image"`
            CanCancel bool  `json:"can_cancel"`
            L2Category []string  `json:"l2_category"`
            SlugKey string  `json:"slug_key"`
            DepartmentID float64  `json:"department_id"`
            Brand string  `json:"brand"`
            Meta map[string]interface{}  `json:"meta"`
            L2CategoryID float64  `json:"l2_category_id"`
            L1CategoryID float64  `json:"l1_category_id"`
            Gender string  `json:"gender"`
            L3CategoryName string  `json:"l3_category_name"`
         
    }
    
    // Weight used by Order
    type Weight struct {

        
            IsDefault bool  `json:"is_default"`
            Unit string  `json:"unit"`
            Shipping float64  `json:"shipping"`
         
    }
    
    // ReturnConfig used by Order
    type ReturnConfig struct {

        
            Returnable bool  `json:"returnable"`
            Unit string  `json:"unit"`
            Time float64  `json:"time"`
         
    }
    
    // Article used by Order
    type Article struct {

        
            Weight Weight  `json:"weight"`
            RawMeta interface{}  `json:"raw_meta"`
            EspModified interface{}  `json:"esp_modified"`
            SellerIdentifier string  `json:"seller_identifier"`
            Code string  `json:"code"`
            ChildDetails map[string]interface{}  `json:"child_details"`
            ReturnConfig ReturnConfig  `json:"return_config"`
            Size string  `json:"size"`
            Identifiers Identifier  `json:"identifiers"`
            ID string  `json:"_id"`
            Dimensions Dimensions  `json:"dimensions"`
            IsSet bool  `json:"is_set"`
            UID string  `json:"uid"`
            ASet map[string]interface{}  `json:"a_set"`
         
    }
    
    // Dates used by Order
    type Dates struct {

        
            DeliveryDate interface{}  `json:"delivery_date"`
            OrderCreated string  `json:"order_created"`
         
    }
    
    // ArticleDetails used by Order
    type ArticleDetails struct {

        
            Status map[string]interface{}  `json:"status"`
         
    }
    
    // StoreAddress used by Order
    type StoreAddress struct {

        
            Version string  `json:"version"`
            Address1 string  `json:"address1"`
            AddressCategory string  `json:"address_category"`
            State string  `json:"state"`
            Phone string  `json:"phone"`
            Area string  `json:"area"`
            Country string  `json:"country"`
            CreatedAt string  `json:"created_at"`
            ContactPerson string  `json:"contact_person"`
            UpdatedAt string  `json:"updated_at"`
            Latitude float64  `json:"latitude"`
            Email string  `json:"email"`
            AddressType string  `json:"address_type"`
            Pincode float64  `json:"pincode"`
            City string  `json:"city"`
            Address2 string  `json:"address2"`
            Longitude float64  `json:"longitude"`
            Landmark string  `json:"landmark"`
            CountryCode string  `json:"country_code"`
         
    }
    
    // EInvoicePortalDetails used by Order
    type EInvoicePortalDetails struct {

        
            Username string  `json:"username"`
            Password string  `json:"password"`
            User string  `json:"user"`
         
    }
    
    // Document used by Order
    type Document struct {

        
            LegalName string  `json:"legal_name"`
            Value string  `json:"value"`
            DsType string  `json:"ds_type"`
            Verified bool  `json:"verified"`
            URL string  `json:"url"`
         
    }
    
    // StoreDocuments used by Order
    type StoreDocuments struct {

        
            Gst Document  `json:"gst"`
         
    }
    
    // StoreEinvoice used by Order
    type StoreEinvoice struct {

        
            Username string  `json:"username"`
            Password string  `json:"password"`
            Enabled bool  `json:"enabled"`
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
    
    // StoreMeta used by Order
    type StoreMeta struct {

        
            Timing []map[string]interface{}  `json:"timing"`
            DisplayName string  `json:"display_name"`
            EinvoicePortalDetails EInvoicePortalDetails  `json:"einvoice_portal_details"`
            GstNumber string  `json:"gst_number"`
            Stage string  `json:"stage"`
            EwaybillPortalDetails map[string]interface{}  `json:"ewaybill_portal_details"`
            Documents StoreDocuments  `json:"documents"`
            ProductReturnConfig map[string]interface{}  `json:"product_return_config"`
            NotificationEmails []string  `json:"notification_emails"`
            GstCredentials StoreGstCredentials  `json:"gst_credentials"`
            AdditionalContactDetails map[string]interface{}  `json:"additional_contact_details"`
         
    }
    
    // Store used by Order
    type Store struct {

        
            MallName string  `json:"mall_name"`
            LoginUsername string  `json:"login_username"`
            StoreAddressJson StoreAddress  `json:"store_address_json"`
            SID string  `json:"s_id"`
            LocationType string  `json:"location_type"`
            Address1 string  `json:"address1"`
            MallArea string  `json:"mall_area"`
            State string  `json:"state"`
            VatNo string  `json:"vat_no"`
            AlohomoraUserID float64  `json:"alohomora_user_id"`
            Code string  `json:"code"`
            BrandID interface{}  `json:"brand_id"`
            PackagingMaterialCount float64  `json:"packaging_material_count"`
            IsArchived bool  `json:"is_archived"`
            Name string  `json:"name"`
            CompanyID float64  `json:"company_id"`
            OrderIntegrationID string  `json:"order_integration_id"`
            Phone float64  `json:"phone"`
            StoreEmail string  `json:"store_email"`
            StoreActiveFrom string  `json:"store_active_from"`
            ParentStoreID float64  `json:"parent_store_id"`
            Country string  `json:"country"`
            CreatedAt string  `json:"created_at"`
            ContactPerson string  `json:"contact_person"`
            BrandStoreTags []string  `json:"brand_store_tags"`
            UpdatedAt string  `json:"updated_at"`
            Latitude float64  `json:"latitude"`
            IsActive bool  `json:"is_active"`
            Meta StoreMeta  `json:"meta"`
            Pincode string  `json:"pincode"`
            IsEnabledForRecon bool  `json:"is_enabled_for_recon"`
            City string  `json:"city"`
            Address2 string  `json:"address2"`
            Longitude float64  `json:"longitude"`
            FulfillmentChannel string  `json:"fulfillment_channel"`
         
    }
    
    // Brand used by Order
    type Brand struct {

        
            Company string  `json:"company"`
            IsVirtualInvoice bool  `json:"is_virtual_invoice"`
            ModifiedOn float64  `json:"modified_on"`
            ScriptLastRan string  `json:"script_last_ran"`
            CreditNoteAllowed bool  `json:"credit_note_allowed"`
            BrandID float64  `json:"brand_id"`
            CreatedOn float64  `json:"created_on"`
            Logo string  `json:"logo"`
            BrandName string  `json:"brand_name"`
            CreditNoteExpiryDays float64  `json:"credit_note_expiry_days"`
            PickupLocation string  `json:"pickup_location"`
            InvoicePrefix string  `json:"invoice_prefix"`
            StartDate string  `json:"start_date"`
         
    }
    
    // B2BPODetails used by Order
    type B2BPODetails struct {

        
            DockerNumber string  `json:"docker_number"`
            PoLineAmount float64  `json:"po_line_amount"`
            TotalGstPercentage float64  `json:"total_gst_percentage"`
            PartialCanRet bool  `json:"partial_can_ret"`
            PoTaxAmount float64  `json:"po_tax_amount"`
            ItemBasePrice float64  `json:"item_base_price"`
         
    }
    
    // BagMeta used by Order
    type BagMeta struct {

        
            B2bPoDetails B2BPODetails  `json:"b2b_po_details"`
         
    }
    
    // BagGSTDetails used by Order
    type BagGSTDetails struct {

        
            GstTag string  `json:"gst_tag"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            IsDefaultHsnCode bool  `json:"is_default_hsn_code"`
            CgstGstFee string  `json:"cgst_gst_fee"`
            IgstTaxPercentage float64  `json:"igst_tax_percentage"`
            SgstGstFee string  `json:"sgst_gst_fee"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            HsnCodeID string  `json:"hsn_code_id"`
            GstFee float64  `json:"gst_fee"`
            GstinCode string  `json:"gstin_code"`
            ValueOfGood float64  `json:"value_of_good"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            CgstTaxPercentage float64  `json:"cgst_tax_percentage"`
            IgstGstFee string  `json:"igst_gst_fee"`
            HsnCode string  `json:"hsn_code"`
            SgstTaxPercentage float64  `json:"sgst_tax_percentage"`
         
    }
    
    // BagDetailsPlatformResponse used by Order
    type BagDetailsPlatformResponse struct {

        
            DisplayName string  `json:"display_name"`
            BagStatusHistory BagStatusHistory  `json:"bag_status_history"`
            ParentPromoBags map[string]interface{}  `json:"parent_promo_bags"`
            AffiliateBagDetails AffiliateBagDetails  `json:"affiliate_bag_details"`
            FinancialBreakup []FinancialBreakup  `json:"financial_breakup"`
            Quantity float64  `json:"quantity"`
            Status BagReturnableCancelableStatus  `json:"status"`
            BagStatus []BagStatusHistory  `json:"bag_status"`
            Identifier string  `json:"identifier"`
            Item Item  `json:"item"`
            SellerIdentifier string  `json:"seller_identifier"`
            NoOfBagsOrder float64  `json:"no_of_bags_order"`
            OrderIntegrationID string  `json:"order_integration_id"`
            Article Article  `json:"article"`
            Dates Dates  `json:"dates"`
            EntityType string  `json:"entity_type"`
            Reasons []map[string]interface{}  `json:"reasons"`
            ArticleDetails ArticleDetails  `json:"article_details"`
            RestoreCoupon bool  `json:"restore_coupon"`
            BagUpdateTime float64  `json:"bag_update_time"`
            AffiliateDetails AffiliateDetails  `json:"affiliate_details"`
            JourneyType string  `json:"journey_type"`
            BType string  `json:"b_type"`
            AppliedPromos []map[string]interface{}  `json:"applied_promos"`
            QcRequired interface{}  `json:"qc_required"`
            RestorePromos map[string]interface{}  `json:"restore_promos"`
            OrderingStore Store  `json:"ordering_store"`
            Brand Brand  `json:"brand"`
            OriginalBagList []float64  `json:"original_bag_list"`
            Meta BagMeta  `json:"meta"`
            LineNumber float64  `json:"line_number"`
            CurrentStatus BagStatusHistory  `json:"current_status"`
            OperationalStatus string  `json:"operational_status"`
            ShipmentID string  `json:"shipment_id"`
            BID float64  `json:"b_id"`
            CurrentOperationalStatus BagStatusHistory  `json:"current_operational_status"`
            GstDetails BagGSTDetails  `json:"gst_details"`
            Tags []string  `json:"tags"`
            Prices Prices  `json:"prices"`
         
    }
    
    // ErrorResponse used by Order
    type ErrorResponse struct {

        
            Message string  `json:"message"`
            Error string  `json:"error"`
         
    }
    
    // Page1 used by Order
    type Page1 struct {

        
            ItemTotal float64  `json:"item_total"`
            PageType string  `json:"page_type"`
            Size float64  `json:"size"`
            HasNext bool  `json:"has_next"`
            Current float64  `json:"current"`
         
    }
    
    // GetBagsPlatformResponse used by Order
    type GetBagsPlatformResponse struct {

        
            Items []BagDetailsPlatformResponse  `json:"items"`
            Page Page1  `json:"page"`
         
    }
    
    // InvalidateShipmentCachePayload used by Order
    type InvalidateShipmentCachePayload struct {

        
            ShipmentIds []string  `json:"shipment_ids"`
            BagIds []string  `json:"bag_ids"`
            AffiliateBagIds []string  `json:"affiliate_bag_ids"`
         
    }
    
    // InvalidateShipmentCacheNestedResponse used by Order
    type InvalidateShipmentCacheNestedResponse struct {

        
            Status float64  `json:"status"`
            ShipmentID string  `json:"shipment_id"`
            Message string  `json:"message"`
            Error string  `json:"error"`
         
    }
    
    // InvalidateShipmentCacheResponse used by Order
    type InvalidateShipmentCacheResponse struct {

        
            Response []InvalidateShipmentCacheNestedResponse  `json:"response"`
         
    }
    
    // ErrorResponse1 used by Order
    type ErrorResponse1 struct {

        
            ErrorTrace string  `json:"error_trace"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
            Status float64  `json:"status"`
         
    }
    
    // StoreReassign used by Order
    type StoreReassign struct {

        
            StoreID float64  `json:"store_id"`
            FyndOrderID string  `json:"fynd_order_id"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
            MongoArticleID string  `json:"mongo_article_id"`
            AffiliateID string  `json:"affiliate_id"`
            ReasonIds []float64  `json:"reason_ids"`
            BagID float64  `json:"bag_id"`
            SetID string  `json:"set_id"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            ItemID string  `json:"item_id"`
         
    }
    
    // StoreReassignResponse used by Order
    type StoreReassignResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // Entities used by Order
    type Entities struct {

        
            ID string  `json:"id"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
            ReasonText string  `json:"reason_text"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            AffiliateID string  `json:"affiliate_id"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
         
    }
    
    // UpdateShipmentLockPayload used by Order
    type UpdateShipmentLockPayload struct {

        
            Action string  `json:"action"`
            ActionType string  `json:"action_type"`
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

        
            BagID float64  `json:"bag_id"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
            IsLocked bool  `json:"is_locked"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
         
    }
    
    // CheckResponse used by Order
    type CheckResponse struct {

        
            OriginalFilter OriginalFilter  `json:"original_filter"`
            IsBagLocked bool  `json:"is_bag_locked"`
            LockStatus string  `json:"lock_status"`
            IsShipmentLocked bool  `json:"is_shipment_locked"`
            Status string  `json:"status"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            AffiliateID string  `json:"affiliate_id"`
            ShipmentID string  `json:"shipment_id"`
            Bags []Bags  `json:"bags"`
         
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
            ToDatetime string  `json:"to_datetime"`
            Description string  `json:"description"`
            CompanyID float64  `json:"company_id"`
            ID float64  `json:"id"`
            FromDatetime string  `json:"from_datetime"`
            CreatedAt string  `json:"created_at"`
            Title string  `json:"title"`
            PlatformID string  `json:"platform_id"`
            PlatformName string  `json:"platform_name"`
         
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

        
            Status bool  `json:"status"`
            CallID string  `json:"call_id"`
         
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
    
    // Products used by Order
    type Products struct {

        
            LineNumber float64  `json:"line_number"`
            Quantity float64  `json:"quantity"`
            Identifier string  `json:"identifier"`
         
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

        
            Reasons ReasonsData  `json:"reasons"`
            Products []Products  `json:"products"`
            DataUpdates DataUpdates  `json:"data_updates"`
            Identifier string  `json:"identifier"`
         
    }
    
    // StatuesRequest used by Order
    type StatuesRequest struct {

        
            Status string  `json:"status"`
            Shipments []ShipmentsRequest  `json:"shipments"`
            ExcludeBagsNextState string  `json:"exclude_bags_next_state"`
         
    }
    
    // UpdateShipmentStatusRequest used by Order
    type UpdateShipmentStatusRequest struct {

        
            Task bool  `json:"task"`
            ForceTransition bool  `json:"force_transition"`
            UnlockBeforeTransition bool  `json:"unlock_before_transition"`
            Statuses []StatuesRequest  `json:"statuses"`
            LockAfterTransition bool  `json:"lock_after_transition"`
         
    }
    
    // ShipmentsResponse used by Order
    type ShipmentsResponse struct {

        
            Code string  `json:"code"`
            Identifier string  `json:"identifier"`
            Status float64  `json:"status"`
            FinalState map[string]interface{}  `json:"final_state"`
            Message string  `json:"message"`
            Meta map[string]interface{}  `json:"meta"`
            Exception string  `json:"exception"`
            StackTrace string  `json:"stack_trace"`
         
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

        
            MarketplaceStoreID string  `json:"marketplace_store_id"`
            StoreID float64  `json:"store_id"`
         
    }
    
    // AffiliateAppConfigMeta used by Order
    type AffiliateAppConfigMeta struct {

        
            Value string  `json:"value"`
            Name string  `json:"name"`
         
    }
    
    // AffiliateAppConfig used by Order
    type AffiliateAppConfig struct {

        
            Description string  `json:"description"`
            Name string  `json:"name"`
            ID string  `json:"id"`
            CreatedAt string  `json:"created_at"`
            Owner string  `json:"owner"`
            Secret string  `json:"secret"`
            Token string  `json:"token"`
            Meta []AffiliateAppConfigMeta  `json:"meta"`
            UpdatedAt string  `json:"updated_at"`
         
    }
    
    // AffiliateInventoryPaymentConfig used by Order
    type AffiliateInventoryPaymentConfig struct {

        
            ModeOfPayment string  `json:"mode_of_payment"`
            Source string  `json:"source"`
         
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
    
    // AffiliateInventoryArticleAssignmentConfig used by Order
    type AffiliateInventoryArticleAssignmentConfig struct {

        
            PostOrderReassignment bool  `json:"post_order_reassignment"`
         
    }
    
    // AffiliateInventoryConfig used by Order
    type AffiliateInventoryConfig struct {

        
            Payment AffiliateInventoryPaymentConfig  `json:"payment"`
            Inventory AffiliateInventoryStoreConfig  `json:"inventory"`
            Order AffiliateInventoryOrderConfig  `json:"order"`
            Logistics AffiliateInventoryLogisticsConfig  `json:"logistics"`
            ArticleAssignment AffiliateInventoryArticleAssignmentConfig  `json:"article_assignment"`
         
    }
    
    // AffiliateConfig used by Order
    type AffiliateConfig struct {

        
            App AffiliateAppConfig  `json:"app"`
            Inventory AffiliateInventoryConfig  `json:"inventory"`
         
    }
    
    // Affiliate used by Order
    type Affiliate struct {

        
            Token string  `json:"token"`
            Config AffiliateConfig  `json:"config"`
            ID string  `json:"id"`
         
    }
    
    // OrderConfig used by Order
    type OrderConfig struct {

        
            ArticleLookup string  `json:"article_lookup"`
            CreateUser bool  `json:"create_user"`
            StoreLookup string  `json:"store_lookup"`
            AffiliateStoreIDMapping []AffiliateStoreIdMapping  `json:"affiliate_store_id_mapping"`
            Affiliate Affiliate  `json:"affiliate"`
            BagEndState string  `json:"bag_end_state"`
         
    }
    
    // ArticleDetails1 used by Order
    type ArticleDetails1 struct {

        
            Dimension map[string]interface{}  `json:"dimension"`
            Quantity float64  `json:"quantity"`
            Category map[string]interface{}  `json:"category"`
            BrandID float64  `json:"brand_id"`
            Weight map[string]interface{}  `json:"weight"`
            ID string  `json:"_id"`
            Attributes map[string]interface{}  `json:"attributes"`
         
    }
    
    // ShipmentDetails used by Order
    type ShipmentDetails struct {

        
            DpID float64  `json:"dp_id"`
            Articles []ArticleDetails1  `json:"articles"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            BoxType string  `json:"box_type"`
            Meta map[string]interface{}  `json:"meta"`
            FulfillmentID float64  `json:"fulfillment_id"`
            Shipments float64  `json:"shipments"`
         
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
            ToPincode string  `json:"to_pincode"`
            Journey string  `json:"journey"`
            Identifier string  `json:"identifier"`
            LocationDetails LocationDetails  `json:"location_details"`
            Action string  `json:"action"`
            PaymentMode string  `json:"payment_mode"`
            Source string  `json:"source"`
         
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

        
            Address2 string  `json:"address2"`
            Email string  `json:"email"`
            Country string  `json:"country"`
            State string  `json:"state"`
            FirstName string  `json:"first_name"`
            LastName string  `json:"last_name"`
            City string  `json:"city"`
            Mobile float64  `json:"mobile"`
            Pincode string  `json:"pincode"`
            Phone float64  `json:"phone"`
            Address1 string  `json:"address1"`
         
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

        
            StoreID float64  `json:"store_id"`
            HsnCodeID string  `json:"hsn_code_id"`
            SellerIdentifier string  `json:"seller_identifier"`
            DeliveryCharge float64  `json:"delivery_charge"`
            AffiliateMeta map[string]interface{}  `json:"affiliate_meta"`
            ItemID float64  `json:"item_id"`
            Identifier map[string]interface{}  `json:"identifier"`
            AvlQty float64  `json:"avl_qty"`
            ItemSize string  `json:"item_size"`
            PdfLinks MarketPlacePdf  `json:"pdf_links"`
            Discount float64  `json:"discount"`
            AffiliateStoreID string  `json:"affiliate_store_id"`
            CompanyID float64  `json:"company_id"`
            UnitPrice float64  `json:"unit_price"`
            PriceEffective float64  `json:"price_effective"`
            ModifiedOn string  `json:"modified_on"`
            TransferPrice float64  `json:"transfer_price"`
            Sku string  `json:"sku"`
            PriceMarked float64  `json:"price_marked"`
            AmountPaid float64  `json:"amount_paid"`
            Quantity float64  `json:"quantity"`
            ID string  `json:"_id"`
            FyndStoreID string  `json:"fynd_store_id"`
         
    }
    
    // OrderInfo used by Order
    type OrderInfo struct {

        
            Payment map[string]interface{}  `json:"payment"`
            Shipment ShipmentData  `json:"shipment"`
            Coupon string  `json:"coupon"`
            Discount float64  `json:"discount"`
            OrderPriority OrderPriority  `json:"order_priority"`
            ShippingAddress OrderUser  `json:"shipping_address"`
            OrderValue float64  `json:"order_value"`
            User UserData  `json:"user"`
            PaymentMode string  `json:"payment_mode"`
            Items map[string]interface{}  `json:"items"`
            BillingAddress OrderUser  `json:"billing_address"`
            CodCharges float64  `json:"cod_charges"`
            Bags []AffiliateBag  `json:"bags"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            DeliveryCharges float64  `json:"delivery_charges"`
         
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

        
            Description string  `json:"description"`
            Slug string  `json:"slug"`
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
    
    // HistoryDict used by Order
    type HistoryDict struct {

        
            L2Detail string  `json:"l2_detail"`
            L3Detail string  `json:"l3_detail"`
            L1Detail string  `json:"l1_detail"`
            Type string  `json:"type"`
            TicketURL string  `json:"ticket_url"`
            User string  `json:"user"`
            Message string  `json:"message"`
            Meta map[string]interface{}  `json:"meta"`
            DisplayMessage string  `json:"display_message"`
            TicketID string  `json:"ticket_id"`
            BagID float64  `json:"bag_id"`
            Createdat string  `json:"createdat"`
            AssignedAgent string  `json:"assigned_agent"`
         
    }
    
    // ShipmentHistoryResponse used by Order
    type ShipmentHistoryResponse struct {

        
            Success bool  `json:"success"`
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
            AmountPaid float64  `json:"amount_paid"`
            BrandName string  `json:"brand_name"`
            OrderID string  `json:"order_id"`
            Message string  `json:"message"`
            PaymentMode string  `json:"payment_mode"`
            PhoneNumber float64  `json:"phone_number"`
            ShipmentID float64  `json:"shipment_id"`
            CustomerName string  `json:"customer_name"`
         
    }
    
    // SendSmsPayload used by Order
    type SendSmsPayload struct {

        
            BagID float64  `json:"bag_id"`
            Slug string  `json:"slug"`
            Data SmsDataPayload  `json:"data"`
         
    }
    
    // OrderDetails used by Order
    type OrderDetails struct {

        
            FyndOrderID string  `json:"fynd_order_id"`
            CreatedAt string  `json:"created_at"`
         
    }
    
    // Meta1 used by Order
    type Meta1 struct {

        
            StateManagerUsed string  `json:"state_manager_used"`
            KafkaEmissionStatus float64  `json:"kafka_emission_status"`
         
    }
    
    // ShipmentDetail used by Order
    type ShipmentDetail struct {

        
            BagList []float64  `json:"bag_list"`
            Remarks string  `json:"remarks"`
            ID float64  `json:"id"`
            Status string  `json:"status"`
            Meta Meta1  `json:"meta"`
            ShipmentID string  `json:"shipment_id"`
         
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
    
    // ManualAssignDPToShipment used by Order
    type ManualAssignDPToShipment struct {

        
            QcRequired string  `json:"qc_required"`
            OrderType string  `json:"order_type"`
            ShipmentIds []string  `json:"shipment_ids"`
            DpID float64  `json:"dp_id"`
         
    }
    
    // ManualAssignDPToShipmentResponse used by Order
    type ManualAssignDPToShipmentResponse struct {

        
            Success string  `json:"success"`
            Errors []string  `json:"errors"`
         
    }
    
    // PaymentMethod used by Order
    type PaymentMethod struct {

        
            Name string  `json:"name"`
            TransactionData map[string]interface{}  `json:"transaction_data"`
            Amount float64  `json:"amount"`
            Meta map[string]interface{}  `json:"meta"`
            Mode string  `json:"mode"`
            RefundBy string  `json:"refund_by"`
            CollectBy string  `json:"collect_by"`
         
    }
    
    // PaymentInfo used by Order
    type PaymentInfo struct {

        
            PrimaryMode string  `json:"primary_mode"`
            PaymentMethods []PaymentMethod  `json:"payment_methods"`
         
    }
    
    // ShippingInfo used by Order
    type ShippingInfo struct {

        
            Country string  `json:"country"`
            FloorNo string  `json:"floor_no"`
            FirstName string  `json:"first_name"`
            Pincode string  `json:"pincode"`
            PrimaryMobileNumber string  `json:"primary_mobile_number"`
            Address1 string  `json:"address1"`
            AlternateMobileNumber string  `json:"alternate_mobile_number"`
            Slot []map[string]interface{}  `json:"slot"`
            LastName string  `json:"last_name"`
            PrimaryEmail string  `json:"primary_email"`
            CustomerCode string  `json:"customer_code"`
            Gender string  `json:"gender"`
            MiddleName string  `json:"middle_name"`
            CountryCode string  `json:"country_code"`
            State string  `json:"state"`
            ShippingType string  `json:"shipping_type"`
            City string  `json:"city"`
            Title string  `json:"title"`
            StateCode string  `json:"state_code"`
            HouseNo string  `json:"house_no"`
            GeoLocation map[string]interface{}  `json:"geo_location"`
            Address2 string  `json:"address2"`
            ExternalCustomerCode string  `json:"external_customer_code"`
            Landmark string  `json:"landmark"`
            AddressType string  `json:"address_type"`
            AlternateEmail string  `json:"alternate_email"`
         
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

        
            Tax Tax  `json:"tax"`
            Code string  `json:"code"`
            Name string  `json:"name"`
            Type string  `json:"type"`
            Amount map[string]interface{}  `json:"amount"`
         
    }
    
    // TaxInfo used by Order
    type TaxInfo struct {

        
            B2bGstinNumber string  `json:"b2b_gstin_number"`
            Gstin string  `json:"gstin"`
         
    }
    
    // ProcessingDates used by Order
    type ProcessingDates struct {

        
            CustomerPickupSlot map[string]interface{}  `json:"customer_pickup_slot"`
            PackByDate string  `json:"pack_by_date"`
            ConfirmByDate string  `json:"confirm_by_date"`
            DispatchAfterDate string  `json:"dispatch_after_date"`
            DispatchByDate string  `json:"dispatch_by_date"`
            DpPickupSlot map[string]interface{}  `json:"dp_pickup_slot"`
         
    }
    
    // LineItem used by Order
    type LineItem struct {

        
            SellerIdentifier string  `json:"seller_identifier"`
            Charges []Charge  `json:"charges"`
            ExternalLineID string  `json:"external_line_id"`
            CustomMessasge string  `json:"custom_messasge"`
            Quantity float64  `json:"quantity"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // Shipment used by Order
    type Shipment struct {

        
            ProcessingDates ProcessingDates  `json:"processing_dates"`
            LocationID float64  `json:"location_id"`
            Priority float64  `json:"priority"`
            Meta map[string]interface{}  `json:"meta"`
            LineItems []LineItem  `json:"line_items"`
            ExternalShipmentID string  `json:"external_shipment_id"`
         
    }
    
    // BillingInfo used by Order
    type BillingInfo struct {

        
            Country string  `json:"country"`
            FloorNo string  `json:"floor_no"`
            FirstName string  `json:"first_name"`
            Pincode string  `json:"pincode"`
            PrimaryMobileNumber string  `json:"primary_mobile_number"`
            Address1 string  `json:"address1"`
            AlternateMobileNumber string  `json:"alternate_mobile_number"`
            LastName string  `json:"last_name"`
            PrimaryEmail string  `json:"primary_email"`
            CustomerCode string  `json:"customer_code"`
            Gender string  `json:"gender"`
            MiddleName string  `json:"middle_name"`
            CountryCode string  `json:"country_code"`
            State string  `json:"state"`
            City string  `json:"city"`
            Title string  `json:"title"`
            StateCode string  `json:"state_code"`
            HouseNo string  `json:"house_no"`
            Address2 string  `json:"address2"`
            ExternalCustomerCode string  `json:"external_customer_code"`
            AlternateEmail string  `json:"alternate_email"`
         
    }
    
    // CreateOrderAPI used by Order
    type CreateOrderAPI struct {

        
            ExternalCreationDate string  `json:"external_creation_date"`
            PaymentInfo PaymentInfo  `json:"payment_info"`
            Config map[string]interface{}  `json:"config"`
            ShippingInfo ShippingInfo  `json:"shipping_info"`
            Charges []Charge  `json:"charges"`
            ExternalOrderID string  `json:"external_order_id"`
            Meta map[string]interface{}  `json:"meta"`
            CurrencyInfo map[string]interface{}  `json:"currency_info"`
            TaxInfo TaxInfo  `json:"tax_info"`
            Shipments []Shipment  `json:"shipments"`
            BillingInfo BillingInfo  `json:"billing_info"`
         
    }
    
    // CreateOrderErrorReponse used by Order
    type CreateOrderErrorReponse struct {

        
            Code string  `json:"code"`
            Status float64  `json:"status"`
            Message string  `json:"message"`
            Info interface{}  `json:"info"`
            Meta string  `json:"meta"`
            Exception string  `json:"exception"`
            RequestID string  `json:"request_id"`
            StackTrace string  `json:"stack_trace"`
         
    }
    
    // PaymentMethods used by Order
    type PaymentMethods struct {

        
            CollectBy string  `json:"collect_by"`
            Mode string  `json:"mode"`
            RefundBy string  `json:"refund_by"`
         
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
    
    // CreateChannelConfigResponse used by Order
    type CreateChannelConfigResponse struct {

        
            IsUpserted bool  `json:"is_upserted"`
            IsInserted bool  `json:"is_inserted"`
            Acknowledged bool  `json:"acknowledged"`
         
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
            EndDate string  `json:"end_date"`
            OrderDetails []FyndOrderIdList  `json:"order_details"`
            StartDate string  `json:"start_date"`
         
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
            CustomJson map[string]interface{}  `json:"_custom_json"`
            IsActive bool  `json:"is_active"`
            AppID string  `json:"app_id"`
         
    }
    
    // GetSearchWordsData used by Catalog
    type GetSearchWordsData struct {

        
            Result map[string]interface{}  `json:"result"`
            Words []string  `json:"words"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            UID string  `json:"uid"`
            IsActive bool  `json:"is_active"`
            AppID string  `json:"app_id"`
         
    }
    
    // ErrorResponse used by Catalog
    type ErrorResponse struct {

        
            Code string  `json:"code"`
            Error string  `json:"error"`
            Status float64  `json:"status"`
            Meta map[string]interface{}  `json:"meta"`
            Message string  `json:"message"`
         
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

        
            Page Page  `json:"page"`
            Items GetSearchWordsData  `json:"items"`
         
    }
    
    // GetSearchWordsResponse used by Catalog
    type GetSearchWordsResponse struct {

        
            Page Page  `json:"page"`
            Items []GetSearchWordsData  `json:"items"`
         
    }
    
    // AutocompletePageAction used by Catalog
    type AutocompletePageAction struct {

        
            Query map[string]interface{}  `json:"query"`
            Type string  `json:"type"`
            Params map[string]interface{}  `json:"params"`
            URL string  `json:"url"`
         
    }
    
    // AutocompleteAction used by Catalog
    type AutocompleteAction struct {

        
            Page AutocompletePageAction  `json:"page"`
            Type string  `json:"type"`
         
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
            Display string  `json:"display"`
            Logo Media  `json:"logo"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // CreateAutocompleteKeyword used by Catalog
    type CreateAutocompleteKeyword struct {

        
            Words []string  `json:"words"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Results []AutocompleteResult  `json:"results"`
            IsActive bool  `json:"is_active"`
            AppID string  `json:"app_id"`
         
    }
    
    // GetAutocompleteWordsData used by Catalog
    type GetAutocompleteWordsData struct {

        
            Words []string  `json:"words"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Results []map[string]interface{}  `json:"results"`
            UID string  `json:"uid"`
            AppID string  `json:"app_id"`
         
    }
    
    // GetAutocompleteWordsResponse used by Catalog
    type GetAutocompleteWordsResponse struct {

        
            Page Page  `json:"page"`
            Items []GetAutocompleteWordsData  `json:"items"`
         
    }
    
    // CreateAutocompleteWordsResponse used by Catalog
    type CreateAutocompleteWordsResponse struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Results []map[string]interface{}  `json:"results"`
            AppID string  `json:"app_id"`
            Words []string  `json:"words"`
         
    }
    
    // ProductBundleItem used by Catalog
    type ProductBundleItem struct {

        
            MinQuantity float64  `json:"min_quantity"`
            ProductUID float64  `json:"product_uid"`
            AutoSelect bool  `json:"auto_select"`
            AutoAddToCart bool  `json:"auto_add_to_cart"`
            MaxQuantity float64  `json:"max_quantity"`
            AllowRemove bool  `json:"allow_remove"`
         
    }
    
    // ProductBundleRequest used by Catalog
    type ProductBundleRequest struct {

        
            Products []ProductBundleItem  `json:"products"`
            Logo string  `json:"logo"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ModifiedOn string  `json:"modified_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            CreatedOn string  `json:"created_on"`
            Slug string  `json:"slug"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            Name string  `json:"name"`
            CompanyID float64  `json:"company_id"`
            Choice string  `json:"choice"`
            Meta map[string]interface{}  `json:"meta"`
            PageVisibility []string  `json:"page_visibility"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // GetProductBundleCreateResponse used by Catalog
    type GetProductBundleCreateResponse struct {

        
            Products []ProductBundleItem  `json:"products"`
            Logo string  `json:"logo"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ModifiedOn string  `json:"modified_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            CreatedOn string  `json:"created_on"`
            Slug string  `json:"slug"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            Name string  `json:"name"`
            CompanyID float64  `json:"company_id"`
            Choice string  `json:"choice"`
            Meta map[string]interface{}  `json:"meta"`
            PageVisibility []string  `json:"page_visibility"`
            ID string  `json:"id"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // GetProductBundleListingResponse used by Catalog
    type GetProductBundleListingResponse struct {

        
            Page Page  `json:"page"`
            Items []GetProductBundleCreateResponse  `json:"items"`
         
    }
    
    // ProductBundleUpdateRequest used by Catalog
    type ProductBundleUpdateRequest struct {

        
            Products []ProductBundleItem  `json:"products"`
            Logo string  `json:"logo"`
            ModifiedOn string  `json:"modified_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Slug string  `json:"slug"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            Name string  `json:"name"`
            CompanyID float64  `json:"company_id"`
            Choice string  `json:"choice"`
            Meta map[string]interface{}  `json:"meta"`
            PageVisibility []string  `json:"page_visibility"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // LimitedProductData used by Catalog
    type LimitedProductData struct {

        
            Images []string  `json:"images"`
            Slug string  `json:"slug"`
            Sizes []string  `json:"sizes"`
            CountryOfOrigin string  `json:"country_of_origin"`
            Identifier map[string]interface{}  `json:"identifier"`
            ItemCode string  `json:"item_code"`
            UID float64  `json:"uid"`
            Price map[string]interface{}  `json:"price"`
            Name string  `json:"name"`
            ShortDescription string  `json:"short_description"`
            Attributes map[string]interface{}  `json:"attributes"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // Size used by Catalog
    type Size struct {

        
            Display string  `json:"display"`
            IsAvailable bool  `json:"is_available"`
            Quantity float64  `json:"quantity"`
            Value string  `json:"value"`
         
    }
    
    // Price used by Catalog
    type Price struct {

        
            Currency string  `json:"currency"`
            MaxMarked float64  `json:"max_marked"`
            MinMarked float64  `json:"min_marked"`
            MinEffective float64  `json:"min_effective"`
            MaxEffective float64  `json:"max_effective"`
         
    }
    
    // GetProducts used by Catalog
    type GetProducts struct {

        
            MinQuantity float64  `json:"min_quantity"`
            ProductDetails LimitedProductData  `json:"product_details"`
            ProductUID float64  `json:"product_uid"`
            AutoSelect bool  `json:"auto_select"`
            AutoAddToCart bool  `json:"auto_add_to_cart"`
            Sizes []Size  `json:"sizes"`
            Price Price  `json:"price"`
            MaxQuantity float64  `json:"max_quantity"`
            AllowRemove bool  `json:"allow_remove"`
         
    }
    
    // GetProductBundleResponse used by Catalog
    type GetProductBundleResponse struct {

        
            Products []GetProducts  `json:"products"`
            Logo string  `json:"logo"`
            Slug string  `json:"slug"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            Name string  `json:"name"`
            CompanyID float64  `json:"company_id"`
            Choice string  `json:"choice"`
            Meta map[string]interface{}  `json:"meta"`
            PageVisibility []string  `json:"page_visibility"`
            IsActive bool  `json:"is_active"`
         
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

        
            Title string  `json:"title"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ModifiedOn string  `json:"modified_on"`
            Image string  `json:"image"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            CreatedOn string  `json:"created_on"`
            Subtitle string  `json:"subtitle"`
            Active bool  `json:"active"`
            BrandID float64  `json:"brand_id"`
            Name string  `json:"name"`
            CompanyID float64  `json:"company_id"`
            Tag string  `json:"tag"`
            Guide Guide  `json:"guide"`
            ID string  `json:"id"`
            Description string  `json:"description"`
         
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

        
            Title string  `json:"title"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ModifiedOn string  `json:"modified_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            CreatedOn string  `json:"created_on"`
            Subtitle string  `json:"subtitle"`
            Active bool  `json:"active"`
            BrandID float64  `json:"brand_id"`
            Name string  `json:"name"`
            CompanyID float64  `json:"company_id"`
            Tag string  `json:"tag"`
            Guide map[string]interface{}  `json:"guide"`
            ID string  `json:"id"`
         
    }
    
    // ApplicationItemSEO used by Catalog
    type ApplicationItemSEO struct {

        
            Title interface{}  `json:"title"`
            Description interface{}  `json:"description"`
         
    }
    
    // ApplicationItemMOQ used by Catalog
    type ApplicationItemMOQ struct {

        
            Minimum float64  `json:"minimum"`
            IncrementUnit float64  `json:"increment_unit"`
            Maximum float64  `json:"maximum"`
         
    }
    
    // MetaFields used by Catalog
    type MetaFields struct {

        
            Value interface{}  `json:"value"`
            Key interface{}  `json:"key"`
         
    }
    
    // ApplicationItemMeta used by Catalog
    type ApplicationItemMeta struct {

        
            Seo ApplicationItemSEO  `json:"seo"`
            IsGift bool  `json:"is_gift"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            IsCod bool  `json:"is_cod"`
            Moq ApplicationItemMOQ  `json:"moq"`
            AltText map[string]interface{}  `json:"alt_text"`
            CustomMeta []MetaFields  `json:"_custom_meta"`
         
    }
    
    // SuccessResponse1 used by Catalog
    type SuccessResponse1 struct {

        
            UID float64  `json:"uid"`
            Success bool  `json:"success"`
         
    }
    
    // SEOData used by Catalog
    type SEOData struct {

        
            Title interface{}  `json:"title"`
            Description interface{}  `json:"description"`
         
    }
    
    // MOQData used by Catalog
    type MOQData struct {

        
            Minimum float64  `json:"minimum"`
            IncrementUnit float64  `json:"increment_unit"`
            Maximum float64  `json:"maximum"`
         
    }
    
    // OwnerAppItemResponse used by Catalog
    type OwnerAppItemResponse struct {

        
            Seo SEOData  `json:"seo"`
            IsGift bool  `json:"is_gift"`
            IsCod bool  `json:"is_cod"`
            Moq MOQData  `json:"moq"`
            AltText map[string]interface{}  `json:"alt_text"`
         
    }
    
    // GetConfigMetadataResponse used by Catalog
    type GetConfigMetadataResponse struct {

        
            Values []map[string]interface{}  `json:"values"`
            Condition []map[string]interface{}  `json:"condition"`
            Data []map[string]interface{}  `json:"data"`
         
    }
    
    // AttributeDetailsGroup used by Catalog
    type AttributeDetailsGroup struct {

        
            Logo string  `json:"logo"`
            Unit string  `json:"unit"`
            DisplayType string  `json:"display_type"`
            Key string  `json:"key"`
            Slug string  `json:"slug"`
            Priority float64  `json:"priority"`
            Name string  `json:"name"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // AppConfigurationDetail used by Catalog
    type AppConfigurationDetail struct {

        
            Logo string  `json:"logo"`
            Slug string  `json:"slug"`
            IsDefault bool  `json:"is_default"`
            Priority float64  `json:"priority"`
            Name string  `json:"name"`
            TemplateSlugs []string  `json:"template_slugs"`
            Attributes []AttributeDetailsGroup  `json:"attributes"`
            IsActive bool  `json:"is_active"`
            AppID string  `json:"app_id"`
         
    }
    
    // ConfigErrorResponse used by Catalog
    type ConfigErrorResponse struct {

        
            Message string  `json:"message"`
         
    }
    
    // PageResponseType used by Catalog
    type PageResponseType struct {

        
            Current float64  `json:"current"`
            TotalCount float64  `json:"total_count"`
            Next float64  `json:"next"`
            HasNext bool  `json:"has_next"`
         
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

        
            Logo string  `json:"logo"`
            Key string  `json:"key"`
            IsDefault bool  `json:"is_default"`
            Priority float64  `json:"priority"`
            Name string  `json:"name"`
            DefaultKey string  `json:"default_key"`
            IsActive bool  `json:"is_active"`
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

        
            Display string  `json:"display"`
            Units []map[string]interface{}  `json:"units"`
            FilterTypes []string  `json:"filter_types"`
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

        
            Variant map[string]interface{}  `json:"variant"`
            Compare map[string]interface{}  `json:"compare"`
            Detail map[string]interface{}  `json:"detail"`
            Similar map[string]interface{}  `json:"similar"`
         
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
    
    // ConfigurationProductVariantConfig used by Catalog
    type ConfigurationProductVariantConfig struct {

        
            Logo string  `json:"logo"`
            DisplayType string  `json:"display_type"`
            Key string  `json:"key"`
            Priority float64  `json:"priority"`
            Name string  `json:"name"`
            Size ProductSize  `json:"size"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // ConfigurationProductVariant used by Catalog
    type ConfigurationProductVariant struct {

        
            Config []ConfigurationProductVariantConfig  `json:"config"`
         
    }
    
    // ConfigurationProductConfig used by Catalog
    type ConfigurationProductConfig struct {

        
            Title string  `json:"title"`
            Logo string  `json:"logo"`
            Key string  `json:"key"`
            Subtitle string  `json:"subtitle"`
            Priority float64  `json:"priority"`
            Size ProductSize  `json:"size"`
            IsActive bool  `json:"is_active"`
         
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

        
            Condition string  `json:"condition"`
            Value string  `json:"value"`
            BucketPoints []ConfigurationBucketPoints  `json:"bucket_points"`
            MapValues []map[string]interface{}  `json:"map_values"`
            Map map[string]interface{}  `json:"map"`
            Sort string  `json:"sort"`
         
    }
    
    // ConfigurationListingFilterConfig used by Catalog
    type ConfigurationListingFilterConfig struct {

        
            Logo string  `json:"logo"`
            ValueConfig ConfigurationListingFilterValue  `json:"value_config"`
            Key string  `json:"key"`
            Priority float64  `json:"priority"`
            Name string  `json:"name"`
            Type string  `json:"type"`
            DisplayName string  `json:"display_name"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // ConfigurationListingFilter used by Catalog
    type ConfigurationListingFilter struct {

        
            AllowSingle bool  `json:"allow_single"`
            AttributeConfig []ConfigurationListingFilterConfig  `json:"attribute_config"`
         
    }
    
    // ConfigurationListingSortConfig used by Catalog
    type ConfigurationListingSortConfig struct {

        
            Logo string  `json:"logo"`
            Key string  `json:"key"`
            Priority float64  `json:"priority"`
            Name string  `json:"name"`
            IsActive bool  `json:"is_active"`
         
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
    
    // AppConfiguration used by Catalog
    type AppConfiguration struct {

        
            CreatedBy map[string]interface{}  `json:"created_by"`
            ConfigID string  `json:"config_id"`
            ModifiedOn string  `json:"modified_on"`
            ConfigType string  `json:"config_type"`
            Product ConfigurationProduct  `json:"product"`
            CreatedOn string  `json:"created_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Listing ConfigurationListing  `json:"listing"`
            Type string  `json:"type"`
            AppID string  `json:"app_id"`
         
    }
    
    // AppCatalogConfiguration used by Catalog
    type AppCatalogConfiguration struct {

        
            CreatedBy map[string]interface{}  `json:"created_by"`
            ConfigID string  `json:"config_id"`
            ModifiedOn string  `json:"modified_on"`
            ConfigType string  `json:"config_type"`
            Product ConfigurationProduct  `json:"product"`
            CreatedOn string  `json:"created_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Listing ConfigurationListing  `json:"listing"`
            Type string  `json:"type"`
            ID string  `json:"id"`
            AppID string  `json:"app_id"`
         
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

        
            ConfigID string  `json:"config_id"`
            ConfigType string  `json:"config_type"`
            Product GetCatalogConfigurationDetailsProduct  `json:"product"`
            Listing GetCatalogConfigurationDetailsSchemaListing  `json:"listing"`
            ID string  `json:"id"`
            AppID string  `json:"app_id"`
         
    }
    
    // GetAppCatalogEntityConfiguration used by Catalog
    type GetAppCatalogEntityConfiguration struct {

        
            Data EntityConfiguration  `json:"data"`
            IsDefault bool  `json:"is_default"`
         
    }
    
    // ProductFiltersValue used by Catalog
    type ProductFiltersValue struct {

        
            DisplayFormat string  `json:"display_format"`
            Min float64  `json:"min"`
            CurrencySymbol string  `json:"currency_symbol"`
            QueryFormat string  `json:"query_format"`
            Value interface{}  `json:"value"`
            SelectedMin float64  `json:"selected_min"`
            IsSelected bool  `json:"is_selected"`
            CurrencyCode string  `json:"currency_code"`
            SelectedMax float64  `json:"selected_max"`
            Display string  `json:"display"`
            Count float64  `json:"count"`
            Max float64  `json:"max"`
         
    }
    
    // ProductFiltersKey used by Catalog
    type ProductFiltersKey struct {

        
            Logo string  `json:"logo"`
            Name string  `json:"name"`
            Operators []string  `json:"operators"`
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
    
    // CollectionImage used by Catalog
    type CollectionImage struct {

        
            AspectRatio string  `json:"aspect_ratio"`
            URL string  `json:"url"`
         
    }
    
    // SeoDetail used by Catalog
    type SeoDetail struct {

        
            Title string  `json:"title"`
            Description string  `json:"description"`
         
    }
    
    // UserInfo used by Catalog
    type UserInfo struct {

        
            Email string  `json:"email"`
            UID string  `json:"uid"`
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
         
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
            Start string  `json:"start"`
            End string  `json:"end"`
            Duration float64  `json:"duration"`
            Cron string  `json:"cron"`
         
    }
    
    // CreateCollection used by Catalog
    type CreateCollection struct {

        
            Logo CollectionImage  `json:"logo"`
            AllowSort bool  `json:"allow_sort"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Slug string  `json:"slug"`
            Meta map[string]interface{}  `json:"meta"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Tags []string  `json:"tags"`
            AppID string  `json:"app_id"`
            Seo SeoDetail  `json:"seo"`
            CreatedBy UserInfo  `json:"created_by"`
            AllowFacets bool  `json:"allow_facets"`
            IsVisible bool  `json:"is_visible"`
            ModifiedBy UserInfo  `json:"modified_by"`
            SortOn string  `json:"sort_on"`
            Type string  `json:"type"`
            Published bool  `json:"published"`
            Query []CollectionQuery  `json:"query"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            Name string  `json:"name"`
            Banners CollectionBanner  `json:"banners"`
            Description string  `json:"description"`
            IsActive bool  `json:"is_active"`
            Badge CollectionBadge  `json:"badge"`
            Schedule CollectionSchedule  `json:"_schedule"`
            Priority float64  `json:"priority"`
         
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

        
            Logo BannerImage  `json:"logo"`
            AllowSort bool  `json:"allow_sort"`
            Slug string  `json:"slug"`
            Cron map[string]interface{}  `json:"cron"`
            Meta map[string]interface{}  `json:"meta"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            AppID string  `json:"app_id"`
            AllowFacets bool  `json:"allow_facets"`
            SortOn string  `json:"sort_on"`
            Type string  `json:"type"`
            Query []CollectionQuery  `json:"query"`
            Name string  `json:"name"`
            Tag []string  `json:"tag"`
            Banners ImageUrls  `json:"banners"`
            Description string  `json:"description"`
            IsActive bool  `json:"is_active"`
            Badge map[string]interface{}  `json:"badge"`
            Schedule map[string]interface{}  `json:"_schedule"`
            Priority float64  `json:"priority"`
         
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

        
            Logo Media1  `json:"logo"`
            AllowSort bool  `json:"allow_sort"`
            Slug string  `json:"slug"`
            Cron map[string]interface{}  `json:"cron"`
            Meta map[string]interface{}  `json:"meta"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            AppID string  `json:"app_id"`
            AllowFacets bool  `json:"allow_facets"`
            UID string  `json:"uid"`
            Type string  `json:"type"`
            Query []CollectionQuery  `json:"query"`
            Name string  `json:"name"`
            Tag []string  `json:"tag"`
            Banners ImageUrls  `json:"banners"`
            Action Action  `json:"action"`
            Description string  `json:"description"`
            IsActive bool  `json:"is_active"`
            Badge map[string]interface{}  `json:"badge"`
            Schedule map[string]interface{}  `json:"_schedule"`
            Priority float64  `json:"priority"`
         
    }
    
    // GetCollectionListingResponse used by Catalog
    type GetCollectionListingResponse struct {

        
            Page Page  `json:"page"`
            Filters CollectionListingFilter  `json:"filters"`
            Items []GetCollectionDetailNest  `json:"items"`
         
    }
    
    // CollectionDetailResponse used by Catalog
    type CollectionDetailResponse struct {

        
            Logo Media1  `json:"logo"`
            AllowSort bool  `json:"allow_sort"`
            Slug string  `json:"slug"`
            Cron map[string]interface{}  `json:"cron"`
            Meta map[string]interface{}  `json:"meta"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            AppID string  `json:"app_id"`
            AllowFacets bool  `json:"allow_facets"`
            UID string  `json:"uid"`
            Type string  `json:"type"`
            Query []CollectionQuery  `json:"query"`
            Name string  `json:"name"`
            Tag []string  `json:"tag"`
            Banners ImageUrls  `json:"banners"`
            Description string  `json:"description"`
            IsActive bool  `json:"is_active"`
            Badge map[string]interface{}  `json:"badge"`
            Schedule map[string]interface{}  `json:"_schedule"`
            Priority float64  `json:"priority"`
         
    }
    
    // UpdateCollection used by Catalog
    type UpdateCollection struct {

        
            Logo CollectionImage  `json:"logo"`
            AllowSort bool  `json:"allow_sort"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Slug string  `json:"slug"`
            Meta map[string]interface{}  `json:"meta"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Tags []string  `json:"tags"`
            Seo SeoDetail  `json:"seo"`
            AllowFacets bool  `json:"allow_facets"`
            IsVisible bool  `json:"is_visible"`
            ModifiedBy UserInfo  `json:"modified_by"`
            SortOn string  `json:"sort_on"`
            Type string  `json:"type"`
            Published bool  `json:"published"`
            Query []CollectionQuery  `json:"query"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            Name string  `json:"name"`
            Banners CollectionBanner  `json:"banners"`
            Description string  `json:"description"`
            IsActive bool  `json:"is_active"`
            Badge CollectionBadge  `json:"badge"`
            Schedule CollectionSchedule  `json:"_schedule"`
            Priority float64  `json:"priority"`
         
    }
    
    // CollectionItem used by Catalog
    type CollectionItem struct {

        
            ItemID float64  `json:"item_id"`
            Priority float64  `json:"priority"`
            Action string  `json:"action"`
         
    }
    
    // CollectionItemUpdate used by Catalog
    type CollectionItemUpdate struct {

        
            AllowFacets bool  `json:"allow_facets"`
            AllowSort bool  `json:"allow_sort"`
            Items []CollectionItem  `json:"items"`
            Type string  `json:"type"`
            Query []CollectionQuery  `json:"query"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
         
    }
    
    // UpdatedResponse used by Catalog
    type UpdatedResponse struct {

        
            Message string  `json:"message"`
            ItemsNotUpdated []float64  `json:"items_not_updated"`
         
    }
    
    // Price1 used by Catalog
    type Price1 struct {

        
            CurrencySymbol string  `json:"currency_symbol"`
            CurrencyCode string  `json:"currency_code"`
            Max float64  `json:"max"`
            Min float64  `json:"min"`
         
    }
    
    // ProductListingPrice used by Catalog
    type ProductListingPrice struct {

        
            Marked Price1  `json:"marked"`
            Effective Price1  `json:"effective"`
         
    }
    
    // ProductBrand used by Catalog
    type ProductBrand struct {

        
            Action Action  `json:"action"`
            Logo Media1  `json:"logo"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
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
    
    // ProductListingDetail used by Catalog
    type ProductListingDetail struct {

        
            ProductOnlineDate string  `json:"product_online_date"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            Tryouts []string  `json:"tryouts"`
            Slug string  `json:"slug"`
            Sellable bool  `json:"sellable"`
            HasVariant bool  `json:"has_variant"`
            UID float64  `json:"uid"`
            ItemType string  `json:"item_type"`
            Type string  `json:"type"`
            ShortDescription string  `json:"short_description"`
            Price ProductListingPrice  `json:"price"`
            Attributes map[string]interface{}  `json:"attributes"`
            PromoMeta map[string]interface{}  `json:"promo_meta"`
            Similars []string  `json:"similars"`
            Color string  `json:"color"`
            ItemCode string  `json:"item_code"`
            Name string  `json:"name"`
            ImageNature string  `json:"image_nature"`
            Description string  `json:"description"`
            Medias []Media1  `json:"medias"`
            Brand ProductBrand  `json:"brand"`
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
            Highlights []string  `json:"highlights"`
            RatingCount float64  `json:"rating_count"`
            Rating float64  `json:"rating"`
            Discount string  `json:"discount"`
         
    }
    
    // GetCollectionItemsResponse used by Catalog
    type GetCollectionItemsResponse struct {

        
            Page Page  `json:"page"`
            Filters []ProductFilters  `json:"filters"`
            SortOn []ProductSortOn  `json:"sort_on"`
            Items []ProductListingDetail  `json:"items"`
         
    }
    
    // CatalogInsightItem used by Catalog
    type CatalogInsightItem struct {

        
            Count float64  `json:"count"`
            OutOfStockCount float64  `json:"out_of_stock_count"`
            SellableCount float64  `json:"sellable_count"`
         
    }
    
    // CatalogInsightBrand used by Catalog
    type CatalogInsightBrand struct {

        
            ArticleFreshness float64  `json:"article_freshness"`
            TotalArticles float64  `json:"total_articles"`
            Name string  `json:"name"`
            AvailableArticles float64  `json:"available_articles"`
            TotalSizes float64  `json:"total_sizes"`
            AvailableSizes float64  `json:"available_sizes"`
         
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

        
            Platform string  `json:"platform"`
            OptLevel string  `json:"opt_level"`
            StoreIds []float64  `json:"store_ids"`
            BrandIds []float64  `json:"brand_ids"`
            CompanyID float64  `json:"company_id"`
            Enabled bool  `json:"enabled"`
         
    }
    
    // CompanyOptIn used by Catalog
    type CompanyOptIn struct {

        
            Platform string  `json:"platform"`
            OptLevel string  `json:"opt_level"`
            StoreIds []float64  `json:"store_ids"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ModifiedOn float64  `json:"modified_on"`
            BrandIds []float64  `json:"brand_ids"`
            CreatedOn float64  `json:"created_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            CompanyID float64  `json:"company_id"`
            Enabled bool  `json:"enabled"`
         
    }
    
    // GetOptInPlatform used by Catalog
    type GetOptInPlatform struct {

        
            Page Page  `json:"page"`
            Items []CompanyOptIn  `json:"items"`
         
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
            TotalArticle float64  `json:"total_article"`
            CompanyID float64  `json:"company_id"`
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
            Brand float64  `json:"brand"`
            Company string  `json:"company"`
         
    }
    
    // StoreDetail used by Catalog
    type StoreDetail struct {

        
            Timing map[string]interface{}  `json:"timing"`
            ModifiedOn string  `json:"modified_on"`
            Manager map[string]interface{}  `json:"manager"`
            Address map[string]interface{}  `json:"address"`
            CreatedOn string  `json:"created_on"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            StoreType string  `json:"store_type"`
            CompanyID float64  `json:"company_id"`
            AdditionalContacts []map[string]interface{}  `json:"additional_contacts"`
            DisplayName string  `json:"display_name"`
            Documents []map[string]interface{}  `json:"documents"`
            StoreCode string  `json:"store_code"`
         
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

        
            Max float64  `json:"max"`
            Min float64  `json:"min"`
         
    }
    
    // AttributeMaster used by Catalog
    type AttributeMaster struct {

        
            Mandatory bool  `json:"mandatory"`
            Format string  `json:"format"`
            Range AttributeSchemaRange  `json:"range"`
            Type string  `json:"type"`
            Multi bool  `json:"multi"`
            AllowedValues []string  `json:"allowed_values"`
         
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
    
    // AttributeMasterFilter used by Catalog
    type AttributeMasterFilter struct {

        
            DependsOn []string  `json:"depends_on"`
            Priority float64  `json:"priority"`
            Indexing bool  `json:"indexing"`
         
    }
    
    // GenderDetail used by Catalog
    type GenderDetail struct {

        
            Details AttributeMasterDetails  `json:"details"`
            Logo string  `json:"logo"`
            IsNested bool  `json:"is_nested"`
            Departments []string  `json:"departments"`
            Schema AttributeMaster  `json:"schema"`
            Slug string  `json:"slug"`
            Name string  `json:"name"`
            Meta AttributeMasterMeta  `json:"meta"`
            EnabledForEndConsumer bool  `json:"enabled_for_end_consumer"`
            Filters AttributeMasterFilter  `json:"filters"`
            ID string  `json:"id"`
            Description string  `json:"description"`
         
    }
    
    // CategoriesResponse used by Catalog
    type CategoriesResponse struct {

        
            SlugKey string  `json:"slug_key"`
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            TemplateSlug string  `json:"template_slug"`
         
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
            Meta map[string]interface{}  `json:"meta"`
            Message string  `json:"message"`
            Errors map[string]interface{}  `json:"errors"`
         
    }
    
    // DepartmentCreateUpdate used by Catalog
    type DepartmentCreateUpdate struct {

        
            PriorityOrder float64  `json:"priority_order"`
            Logo string  `json:"logo"`
            Synonyms []string  `json:"synonyms"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Cls string  `json:"_cls"`
            Platforms map[string]interface{}  `json:"platforms"`
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            Tags []string  `json:"tags"`
            IsActive bool  `json:"is_active"`
         
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

        
            UID string  `json:"uid"`
            Contact string  `json:"contact"`
            Username string  `json:"username"`
            ID string  `json:"_id"`
            UserID string  `json:"user_id"`
         
    }
    
    // GetDepartment used by Catalog
    type GetDepartment struct {

        
            PriorityOrder float64  `json:"priority_order"`
            Logo string  `json:"logo"`
            Synonyms []string  `json:"synonyms"`
            CreatedBy UserSerializer  `json:"created_by"`
            PageNo float64  `json:"page_no"`
            ModifiedOn string  `json:"modified_on"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            CreatedOn string  `json:"created_on"`
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            ItemType string  `json:"item_type"`
            Search string  `json:"search"`
            IsActive bool  `json:"is_active"`
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
            Meta map[string]interface{}  `json:"meta"`
            Message string  `json:"message"`
            Errors map[string]interface{}  `json:"errors"`
         
    }
    
    // UserDetail used by Catalog
    type UserDetail struct {

        
            Contact string  `json:"contact"`
            SuperUser bool  `json:"super_user"`
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
         
    }
    
    // DepartmentModel used by Catalog
    type DepartmentModel struct {

        
            PriorityOrder float64  `json:"priority_order"`
            Logo interface{}  `json:"logo"`
            CreatedBy UserDetail  `json:"created_by"`
            Synonyms []interface{}  `json:"synonyms"`
            ModifiedOn string  `json:"modified_on"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ModifiedBy UserDetail  `json:"modified_by"`
            CreatedOn string  `json:"created_on"`
            VerifiedOn string  `json:"verified_on"`
            Slug interface{}  `json:"slug"`
            Cls interface{}  `json:"_cls"`
            UID float64  `json:"uid"`
            Name interface{}  `json:"name"`
            VerifiedBy UserDetail  `json:"verified_by"`
            ID interface{}  `json:"_id"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // ProductTemplate used by Catalog
    type ProductTemplate struct {

        
            IsActive bool  `json:"is_active"`
            Logo string  `json:"logo"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ModifiedOn string  `json:"modified_on"`
            Departments []string  `json:"departments"`
            IsExpirable bool  `json:"is_expirable"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            CreatedOn string  `json:"created_on"`
            Slug string  `json:"slug"`
            Name string  `json:"name"`
            Attributes []string  `json:"attributes"`
            Tag string  `json:"tag"`
            Categories []string  `json:"categories"`
            IsArchived bool  `json:"is_archived"`
            Description string  `json:"description"`
            IsPhysical bool  `json:"is_physical"`
         
    }
    
    // TemplatesResponse used by Catalog
    type TemplatesResponse struct {

        
            Page Page  `json:"page"`
            Items ProductTemplate  `json:"items"`
         
    }
    
    // TemplateDetails used by Catalog
    type TemplateDetails struct {

        
            IsActive bool  `json:"is_active"`
            Logo string  `json:"logo"`
            Departments []string  `json:"departments"`
            IsExpirable bool  `json:"is_expirable"`
            Slug string  `json:"slug"`
            Name string  `json:"name"`
            Attributes []string  `json:"attributes"`
            Tag string  `json:"tag"`
            Categories []string  `json:"categories"`
            IsArchived bool  `json:"is_archived"`
            ID string  `json:"id"`
            Description string  `json:"description"`
            IsPhysical bool  `json:"is_physical"`
         
    }
    
    // Properties used by Catalog
    type Properties struct {

        
            ReturnConfig map[string]interface{}  `json:"return_config"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            Slug map[string]interface{}  `json:"slug"`
            BrandUID map[string]interface{}  `json:"brand_uid"`
            Tags map[string]interface{}  `json:"tags"`
            SizeGuide map[string]interface{}  `json:"size_guide"`
            Media map[string]interface{}  `json:"media"`
            CategorySlug map[string]interface{}  `json:"category_slug"`
            NoOfBoxes map[string]interface{}  `json:"no_of_boxes"`
            ProductGroupTag map[string]interface{}  `json:"product_group_tag"`
            Command map[string]interface{}  `json:"command"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            ItemType map[string]interface{}  `json:"item_type"`
            ShortDescription map[string]interface{}  `json:"short_description"`
            IsDependent map[string]interface{}  `json:"is_dependent"`
            ProductPublish map[string]interface{}  `json:"product_publish"`
            MultiSize map[string]interface{}  `json:"multi_size"`
            Currency map[string]interface{}  `json:"currency"`
            Variants map[string]interface{}  `json:"variants"`
            TraderType map[string]interface{}  `json:"trader_type"`
            CountryOfOrigin map[string]interface{}  `json:"country_of_origin"`
            ItemCode map[string]interface{}  `json:"item_code"`
            Name map[string]interface{}  `json:"name"`
            Trader map[string]interface{}  `json:"trader"`
            HsnCode map[string]interface{}  `json:"hsn_code"`
            Description map[string]interface{}  `json:"description"`
            IsActive map[string]interface{}  `json:"is_active"`
            Highlights map[string]interface{}  `json:"highlights"`
            Sizes map[string]interface{}  `json:"sizes"`
         
    }
    
    // GlobalValidation used by Catalog
    type GlobalValidation struct {

        
            Required []string  `json:"required"`
            Title string  `json:"title"`
            Properties Properties  `json:"properties"`
            Type string  `json:"type"`
            Definitions map[string]interface{}  `json:"definitions"`
            Description string  `json:"description"`
         
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

        
            Email string  `json:"email"`
            UID string  `json:"uid"`
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
         
    }
    
    // ProductTemplateExportResponse used by Catalog
    type ProductTemplateExportResponse struct {

        
            CreatedBy UserInfo1  `json:"created_by"`
            ModifiedOn string  `json:"modified_on"`
            URL string  `json:"url"`
            NotificationEmails []string  `json:"notification_emails"`
            TaskID string  `json:"task_id"`
            SellerID float64  `json:"seller_id"`
            Type string  `json:"type"`
            Status string  `json:"status"`
            Filters map[string]interface{}  `json:"filters"`
            CompletedOn string  `json:"completed_on"`
         
    }
    
    // ProductDownloadsResponse used by Catalog
    type ProductDownloadsResponse struct {

        
            Items []ProductTemplateExportResponse  `json:"items"`
         
    }
    
    // ProductTemplateExportFilterRequest used by Catalog
    type ProductTemplateExportFilterRequest struct {

        
            Brands []string  `json:"brands"`
            FromDate string  `json:"from_date"`
            CatalogueTypes []string  `json:"catalogue_types"`
            Templates []string  `json:"templates"`
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

        
            CatalogID float64  `json:"catalog_id"`
            Name string  `json:"name"`
         
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

        
            Logo string  `json:"logo"`
            Landscape string  `json:"landscape"`
            Portrait string  `json:"portrait"`
         
    }
    
    // CategoryRequestBody used by Catalog
    type CategoryRequestBody struct {

        
            Marketplaces CategoryMapping  `json:"marketplaces"`
            Synonyms []string  `json:"synonyms"`
            Hierarchy []Hierarchy  `json:"hierarchy"`
            Departments []float64  `json:"departments"`
            Level float64  `json:"level"`
            Media Media2  `json:"media"`
            Tryouts []string  `json:"tryouts"`
            Slug string  `json:"slug"`
            Priority float64  `json:"priority"`
            Name string  `json:"name"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // CategoryCreateResponse used by Catalog
    type CategoryCreateResponse struct {

        
            Message string  `json:"message"`
            UID float64  `json:"uid"`
         
    }
    
    // Category used by Catalog
    type Category struct {

        
            Marketplaces CategoryMapping  `json:"marketplaces"`
            Synonyms []string  `json:"synonyms"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Hierarchy []Hierarchy  `json:"hierarchy"`
            ModifiedOn string  `json:"modified_on"`
            Departments []float64  `json:"departments"`
            Level float64  `json:"level"`
            Media Media2  `json:"media"`
            Tryouts []string  `json:"tryouts"`
            CreatedOn string  `json:"created_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Slug string  `json:"slug"`
            Priority float64  `json:"priority"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            ID string  `json:"id"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // CategoryResponse used by Catalog
    type CategoryResponse struct {

        
            Page Page  `json:"page"`
            Items []Category  `json:"items"`
         
    }
    
    // CategoryUpdateResponse used by Catalog
    type CategoryUpdateResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
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

        
            URL string  `json:"url"`
            Tag string  `json:"tag"`
         
    }
    
    // NetQuantity used by Catalog
    type NetQuantity struct {

        
            Unit interface{}  `json:"unit"`
            Value float64  `json:"value"`
         
    }
    
    // CustomOrder used by Catalog
    type CustomOrder struct {

        
            ManufacturingTimeUnit string  `json:"manufacturing_time_unit"`
            IsCustomOrder bool  `json:"is_custom_order"`
            ManufacturingTime float64  `json:"manufacturing_time"`
         
    }
    
    // ProductPublish used by Catalog
    type ProductPublish struct {

        
            IsSet bool  `json:"is_set"`
            ProductOnlineDate string  `json:"product_online_date"`
         
    }
    
    // Trader used by Catalog
    type Trader struct {

        
            Address []string  `json:"address"`
            Name interface{}  `json:"name"`
            Type string  `json:"type"`
         
    }
    
    // TaxIdentifier used by Catalog
    type TaxIdentifier struct {

        
            ReportingHsn string  `json:"reporting_hsn"`
            HsnCodeID string  `json:"hsn_code_id"`
            HsnCode string  `json:"hsn_code"`
         
    }
    
    // ProductCreateUpdateSchemaV2 used by Catalog
    type ProductCreateUpdateSchemaV2 struct {

        
            ReturnConfig ReturnConfig  `json:"return_config"`
            TeaserTag TeaserTag  `json:"teaser_tag"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Slug string  `json:"slug"`
            IsSet bool  `json:"is_set"`
            CompanyID float64  `json:"company_id"`
            BrandUID float64  `json:"brand_uid"`
            NetQuantity NetQuantity  `json:"net_quantity"`
            Requester string  `json:"requester"`
            Tags []string  `json:"tags"`
            ChangeRequestID interface{}  `json:"change_request_id"`
            Departments []float64  `json:"departments"`
            SizeGuide string  `json:"size_guide"`
            Media []Media1  `json:"media"`
            CategorySlug string  `json:"category_slug"`
            NoOfBoxes float64  `json:"no_of_boxes"`
            ProductGroupTag []string  `json:"product_group_tag"`
            UID float64  `json:"uid"`
            CustomOrder CustomOrder  `json:"custom_order"`
            IsImageLessProduct bool  `json:"is_image_less_product"`
            ShortDescription string  `json:"short_description"`
            ItemType string  `json:"item_type"`
            IsDependent bool  `json:"is_dependent"`
            Attributes map[string]interface{}  `json:"attributes"`
            VariantMedia map[string]interface{}  `json:"variant_media"`
            ProductPublish ProductPublish  `json:"product_publish"`
            MultiSize bool  `json:"multi_size"`
            Currency string  `json:"currency"`
            Variants map[string]interface{}  `json:"variants"`
            CountryOfOrigin string  `json:"country_of_origin"`
            ItemCode string  `json:"item_code"`
            Name string  `json:"name"`
            Trader []Trader  `json:"trader"`
            Action string  `json:"action"`
            Description string  `json:"description"`
            IsActive bool  `json:"is_active"`
            BulkJobID string  `json:"bulk_job_id"`
            TemplateTag string  `json:"template_tag"`
            Highlights []string  `json:"highlights"`
            TaxIdentifier TaxIdentifier  `json:"tax_identifier"`
            Sizes []map[string]interface{}  `json:"sizes"`
            VariantGroup map[string]interface{}  `json:"variant_group"`
         
    }
    
    // ReturnConfigResponse used by Catalog
    type ReturnConfigResponse struct {

        
            Time float64  `json:"time"`
            Unit string  `json:"unit"`
            Returnable bool  `json:"returnable"`
         
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
    
    // ProductPublish1 used by Catalog
    type ProductPublish1 struct {

        
            IsSet bool  `json:"is_set"`
            ProductOnlineDate string  `json:"product_online_date"`
         
    }
    
    // VerifiedBy used by Catalog
    type VerifiedBy struct {

        
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
         
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

        
            Logo Logo  `json:"logo"`
            Name string  `json:"name"`
            UID float64  `json:"uid"`
         
    }
    
    // ProductSchemaV2 used by Catalog
    type ProductSchemaV2 struct {

        
            ReturnConfig ReturnConfigResponse  `json:"return_config"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            Stage string  `json:"stage"`
            AllCompanyIds []float64  `json:"all_company_ids"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Slug string  `json:"slug"`
            IsSet bool  `json:"is_set"`
            Moq map[string]interface{}  `json:"moq"`
            CompanyID float64  `json:"company_id"`
            BrandUID float64  `json:"brand_uid"`
            NetQuantity NetQuantityResponse  `json:"net_quantity"`
            Tags []string  `json:"tags"`
            IsPhysical bool  `json:"is_physical"`
            Category map[string]interface{}  `json:"category"`
            Images []Image  `json:"images"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ModifiedOn string  `json:"modified_on"`
            Departments []float64  `json:"departments"`
            SizeGuide string  `json:"size_guide"`
            Media []Media1  `json:"media"`
            CategoryUID float64  `json:"category_uid"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            CreatedOn string  `json:"created_on"`
            CategorySlug string  `json:"category_slug"`
            VerifiedOn string  `json:"verified_on"`
            NoOfBoxes float64  `json:"no_of_boxes"`
            ProductGroupTag []string  `json:"product_group_tag"`
            UID float64  `json:"uid"`
            ItemType string  `json:"item_type"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            ShortDescription string  `json:"short_description"`
            IsImageLessProduct bool  `json:"is_image_less_product"`
            IsDependent bool  `json:"is_dependent"`
            Attributes map[string]interface{}  `json:"attributes"`
            VariantMedia map[string]interface{}  `json:"variant_media"`
            AllIdentifiers []string  `json:"all_identifiers"`
            ProductPublish ProductPublish1  `json:"product_publish"`
            MultiSize bool  `json:"multi_size"`
            Currency string  `json:"currency"`
            IsExpirable bool  `json:"is_expirable"`
            Variants map[string]interface{}  `json:"variants"`
            Color string  `json:"color"`
            CountryOfOrigin string  `json:"country_of_origin"`
            Pending string  `json:"pending"`
            ItemCode string  `json:"item_code"`
            Name string  `json:"name"`
            Trader []Trader  `json:"trader"`
            HsnCode string  `json:"hsn_code"`
            VerifiedBy VerifiedBy  `json:"verified_by"`
            ImageNature string  `json:"image_nature"`
            Description string  `json:"description"`
            IsActive bool  `json:"is_active"`
            ID string  `json:"id"`
            PrimaryColor string  `json:"primary_color"`
            Brand Brand  `json:"brand"`
            TemplateTag string  `json:"template_tag"`
            Highlights []string  `json:"highlights"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            Sizes []map[string]interface{}  `json:"sizes"`
            L3Mapping []string  `json:"l3_mapping"`
            VariantGroup map[string]interface{}  `json:"variant_group"`
            AllSizes []map[string]interface{}  `json:"all_sizes"`
         
    }
    
    // ProductListingResponseV2 used by Catalog
    type ProductListingResponseV2 struct {

        
            Page Page  `json:"page"`
            Items []ProductSchemaV2  `json:"items"`
         
    }
    
    // ProductVariants used by Catalog
    type ProductVariants struct {

        
            Media []Media1  `json:"media"`
            CategoryUID float64  `json:"category_uid"`
            ItemCode string  `json:"item_code"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            BrandUID float64  `json:"brand_uid"`
         
    }
    
    // ProductVariantsResponse used by Catalog
    type ProductVariantsResponse struct {

        
            Page Page  `json:"page"`
            Variants []ProductVariants  `json:"variants"`
         
    }
    
    // AttributeMasterSerializer used by Catalog
    type AttributeMasterSerializer struct {

        
            Synonyms map[string]interface{}  `json:"synonyms"`
            Logo string  `json:"logo"`
            Unit string  `json:"unit"`
            Schema AttributeMaster  `json:"schema"`
            Slug string  `json:"slug"`
            EnabledForEndConsumer bool  `json:"enabled_for_end_consumer"`
            Tags []string  `json:"tags"`
            Details AttributeMasterDetails  `json:"details"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ModifiedOn string  `json:"modified_on"`
            Departments []string  `json:"departments"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            CreatedOn string  `json:"created_on"`
            RawKey string  `json:"raw_key"`
            Suggestion string  `json:"suggestion"`
            IsNested bool  `json:"is_nested"`
            Variant bool  `json:"variant"`
            Name string  `json:"name"`
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

        
            GtinValue string  `json:"gtin_value"`
            Primary bool  `json:"primary"`
            GtinType string  `json:"gtin_type"`
         
    }
    
    // AllSizes used by Catalog
    type AllSizes struct {

        
            ItemHeight float64  `json:"item_height"`
            ItemWeight float64  `json:"item_weight"`
            Identifiers []ValidateIdentifier  `json:"identifiers"`
            ItemLength float64  `json:"item_length"`
            ItemWeightUnitOfMeasure interface{}  `json:"item_weight_unit_of_measure"`
            ItemWidth float64  `json:"item_width"`
            ItemDimensionsUnitOfMeasure string  `json:"item_dimensions_unit_of_measure"`
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

        
            IsSet bool  `json:"is_set"`
            ProductOnlineDate float64  `json:"product_online_date"`
         
    }
    
    // Product used by Catalog
    type Product struct {

        
            ReturnConfig ReturnConfigResponse  `json:"return_config"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            Stage string  `json:"stage"`
            AllCompanyIds []float64  `json:"all_company_ids"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Slug string  `json:"slug"`
            IsSet bool  `json:"is_set"`
            Moq map[string]interface{}  `json:"moq"`
            CompanyID float64  `json:"company_id"`
            BrandUID float64  `json:"brand_uid"`
            NetQuantity NetQuantityResponse  `json:"net_quantity"`
            Tags []string  `json:"tags"`
            IsPhysical bool  `json:"is_physical"`
            Category map[string]interface{}  `json:"category"`
            Images []Image  `json:"images"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ModifiedOn string  `json:"modified_on"`
            Departments []float64  `json:"departments"`
            SizeGuide string  `json:"size_guide"`
            Media []Media1  `json:"media"`
            CategoryUID float64  `json:"category_uid"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            CreatedOn string  `json:"created_on"`
            CategorySlug string  `json:"category_slug"`
            VerifiedOn string  `json:"verified_on"`
            NoOfBoxes float64  `json:"no_of_boxes"`
            ProductGroupTag []string  `json:"product_group_tag"`
            UID float64  `json:"uid"`
            ItemType string  `json:"item_type"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            ShortDescription string  `json:"short_description"`
            IsImageLessProduct bool  `json:"is_image_less_product"`
            IsDependent bool  `json:"is_dependent"`
            Attributes map[string]interface{}  `json:"attributes"`
            VariantMedia map[string]interface{}  `json:"variant_media"`
            AllIdentifiers []string  `json:"all_identifiers"`
            ProductPublish ProductPublished  `json:"product_publish"`
            MultiSize bool  `json:"multi_size"`
            Currency string  `json:"currency"`
            IsExpirable bool  `json:"is_expirable"`
            Variants map[string]interface{}  `json:"variants"`
            Color string  `json:"color"`
            CountryOfOrigin string  `json:"country_of_origin"`
            Pending string  `json:"pending"`
            ItemCode string  `json:"item_code"`
            Name string  `json:"name"`
            Trader []Trader  `json:"trader"`
            HsnCode string  `json:"hsn_code"`
            VerifiedBy VerifiedBy  `json:"verified_by"`
            ImageNature string  `json:"image_nature"`
            Description string  `json:"description"`
            IsActive bool  `json:"is_active"`
            ID string  `json:"id"`
            PrimaryColor string  `json:"primary_color"`
            Brand Brand  `json:"brand"`
            TemplateTag string  `json:"template_tag"`
            Highlights []string  `json:"highlights"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            Sizes []map[string]interface{}  `json:"sizes"`
            L3Mapping []string  `json:"l3_mapping"`
            VariantGroup map[string]interface{}  `json:"variant_group"`
            AllSizes []map[string]interface{}  `json:"all_sizes"`
         
    }
    
    // ProductListingResponse used by Catalog
    type ProductListingResponse struct {

        
            Page Page  `json:"page"`
            Items []Product  `json:"items"`
         
    }
    
    // BulkJob used by Catalog
    type BulkJob struct {

        
            CreatedBy UserInfo1  `json:"created_by"`
            Cancelled float64  `json:"cancelled"`
            ModifiedOn string  `json:"modified_on"`
            CancelledRecords []map[string]interface{}  `json:"cancelled_records"`
            Stage string  `json:"stage"`
            TemplateTag string  `json:"template_tag"`
            ModifiedBy UserInfo1  `json:"modified_by"`
            CreatedOn string  `json:"created_on"`
            FailedRecords []map[string]interface{}  `json:"failed_records"`
            FilePath string  `json:"file_path"`
            Failed float64  `json:"failed"`
            CompanyID float64  `json:"company_id"`
            Succeed float64  `json:"succeed"`
            TrackingURL string  `json:"tracking_url"`
            CustomTemplateTag string  `json:"custom_template_tag"`
            Total float64  `json:"total"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // BulkResponse used by Catalog
    type BulkResponse struct {

        
            CreatedBy UserInfo1  `json:"created_by"`
            ModifiedOn string  `json:"modified_on"`
            ModifiedBy UserInfo1  `json:"modified_by"`
            CreatedOn string  `json:"created_on"`
            BatchID string  `json:"batch_id"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // UserDetail1 used by Catalog
    type UserDetail1 struct {

        
            Username string  `json:"username"`
            FullName string  `json:"full_name"`
            UserID string  `json:"user_id"`
         
    }
    
    // ProductBulkRequest used by Catalog
    type ProductBulkRequest struct {

        
            CancelledRecords []string  `json:"cancelled_records"`
            TemplateTag string  `json:"template_tag"`
            Stage string  `json:"stage"`
            ModifiedOn string  `json:"modified_on"`
            CreatedBy UserDetail1  `json:"created_by"`
            Cancelled float64  `json:"cancelled"`
            FailedRecords []string  `json:"failed_records"`
            CreatedOn string  `json:"created_on"`
            ModifiedBy UserDetail1  `json:"modified_by"`
            FilePath string  `json:"file_path"`
            Failed float64  `json:"failed"`
            CompanyID float64  `json:"company_id"`
            Succeed float64  `json:"succeed"`
            Template ProductTemplate  `json:"template"`
            Total float64  `json:"total"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // ProductBulkRequestList used by Catalog
    type ProductBulkRequestList struct {

        
            Page Page  `json:"page"`
            Items ProductBulkRequest  `json:"items"`
         
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
    
    // ProductBulkAssets used by Catalog
    type ProductBulkAssets struct {

        
            User map[string]interface{}  `json:"user"`
            CompanyID float64  `json:"company_id"`
            URL string  `json:"url"`
         
    }
    
    // UserCommon used by Catalog
    type UserCommon struct {

        
            CompanyID float64  `json:"company_id"`
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
         
    }
    
    // Items used by Catalog
    type Items struct {

        
            IsActive bool  `json:"is_active"`
            CreatedBy UserCommon  `json:"created_by"`
            CancelledRecords []string  `json:"cancelled_records"`
            ModifiedOn string  `json:"modified_on"`
            Stage string  `json:"stage"`
            Cancelled float64  `json:"cancelled"`
            ModifiedBy UserCommon  `json:"modified_by"`
            CreatedOn string  `json:"created_on"`
            FailedRecords []string  `json:"failed_records"`
            FilePath string  `json:"file_path"`
            Failed float64  `json:"failed"`
            CompanyID float64  `json:"company_id"`
            TrackingURL string  `json:"tracking_url"`
            Total float64  `json:"total"`
            ID string  `json:"id"`
            Succeed float64  `json:"succeed"`
            Retry float64  `json:"retry"`
         
    }
    
    // BulkAssetResponse used by Catalog
    type BulkAssetResponse struct {

        
            Page Page  `json:"page"`
            Items []Items  `json:"items"`
         
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
    
    // ItemQuery used by Catalog
    type ItemQuery struct {

        
            ItemCode string  `json:"item_code"`
            UID float64  `json:"uid"`
            BrandUID float64  `json:"brand_uid"`
         
    }
    
    // GTIN used by Catalog
    type GTIN struct {

        
            GtinValue interface{}  `json:"gtin_value"`
            Primary bool  `json:"primary"`
            GtinType string  `json:"gtin_type"`
         
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
            Quantity float64  `json:"quantity"`
            SizeDistribution SizeDistribution  `json:"size_distribution"`
         
    }
    
    // InvSize used by Catalog
    type InvSize struct {

        
            ItemHeight float64  `json:"item_height"`
            PriceEffective float64  `json:"price_effective"`
            ItemWeight float64  `json:"item_weight"`
            Currency string  `json:"currency"`
            Identifiers []GTIN  `json:"identifiers"`
            Quantity float64  `json:"quantity"`
            PriceTransfer float64  `json:"price_transfer"`
            IsSet bool  `json:"is_set"`
            ItemLength float64  `json:"item_length"`
            Price float64  `json:"price"`
            ItemWeightUnitOfMeasure string  `json:"item_weight_unit_of_measure"`
            ItemWidth float64  `json:"item_width"`
            Set InventorySet  `json:"set"`
            ExpirationDate string  `json:"expiration_date"`
            ItemDimensionsUnitOfMeasure string  `json:"item_dimensions_unit_of_measure"`
            Size interface{}  `json:"size"`
            StoreCode string  `json:"store_code"`
         
    }
    
    // InventoryRequest used by Catalog
    type InventoryRequest struct {

        
            Item ItemQuery  `json:"item"`
            CompanyID float64  `json:"company_id"`
            Sizes []InvSize  `json:"sizes"`
         
    }
    
    // InventoryResponse used by Catalog
    type InventoryResponse struct {

        
            SellableQuantity float64  `json:"sellable_quantity"`
            PriceEffective float64  `json:"price_effective"`
            Currency string  `json:"currency"`
            Identifiers map[string]interface{}  `json:"identifiers"`
            ItemID float64  `json:"item_id"`
            Quantity float64  `json:"quantity"`
            PriceTransfer float64  `json:"price_transfer"`
            SellerIdentifier string  `json:"seller_identifier"`
            UID string  `json:"uid"`
            Price float64  `json:"price"`
            Store map[string]interface{}  `json:"store"`
            Size string  `json:"size"`
            InventoryUpdatedOn string  `json:"inventory_updated_on"`
         
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
    
    // WeightResponse used by Catalog
    type WeightResponse struct {

        
            Unit string  `json:"unit"`
            Shipping float64  `json:"shipping"`
            IsDefault bool  `json:"is_default"`
         
    }
    
    // PriceMeta used by Catalog
    type PriceMeta struct {

        
            Currency string  `json:"currency"`
            Transfer float64  `json:"transfer"`
            Effective float64  `json:"effective"`
            UpdatedAt string  `json:"updated_at"`
            TpNotes map[string]interface{}  `json:"tp_notes"`
            Marked float64  `json:"marked"`
         
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

        
            Damaged QuantityBase  `json:"damaged"`
            OrderCommitted QuantityBase  `json:"order_committed"`
            NotAvailable QuantityBase  `json:"not_available"`
            Sellable QuantityBase  `json:"sellable"`
         
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
    
    // CompanyMeta used by Catalog
    type CompanyMeta struct {

        
            ID float64  `json:"id"`
         
    }
    
    // DimensionResponse used by Catalog
    type DimensionResponse struct {

        
            Unit string  `json:"unit"`
            Height float64  `json:"height"`
            IsDefault bool  `json:"is_default"`
            Length float64  `json:"length"`
            Width float64  `json:"width"`
         
    }
    
    // ManufacturerResponse used by Catalog
    type ManufacturerResponse struct {

        
            Address string  `json:"address"`
            Name string  `json:"name"`
            IsDefault bool  `json:"is_default"`
         
    }
    
    // InventorySellerResponse used by Catalog
    type InventorySellerResponse struct {

        
            ReturnConfig ReturnConfig1  `json:"return_config"`
            RawMeta map[string]interface{}  `json:"raw_meta"`
            Stage string  `json:"stage"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            SellerIdentifier string  `json:"seller_identifier"`
            Weight WeightResponse  `json:"weight"`
            IsSet bool  `json:"is_set"`
            TotalQuantity float64  `json:"total_quantity"`
            Meta map[string]interface{}  `json:"meta"`
            Tags []string  `json:"tags"`
            AddedOnStore string  `json:"added_on_store"`
            FyndArticleCode string  `json:"fynd_article_code"`
            CreatedBy UserSerializer  `json:"created_by"`
            Fragile bool  `json:"fragile"`
            FyndMeta map[string]interface{}  `json:"fynd_meta"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            UID string  `json:"uid"`
            Price PriceMeta  `json:"price"`
            ExpirationDate string  `json:"expiration_date"`
            Size string  `json:"size"`
            CountryOfOrigin string  `json:"country_of_origin"`
            Identifier map[string]interface{}  `json:"identifier"`
            TraceID string  `json:"trace_id"`
            Trader []Trader1  `json:"trader"`
            Quantities Quantities  `json:"quantities"`
            Set InventorySet  `json:"set"`
            Store StoreMeta  `json:"store"`
            IsActive bool  `json:"is_active"`
            Brand BrandMeta  `json:"brand"`
            TrackInventory bool  `json:"track_inventory"`
            ItemID float64  `json:"item_id"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            FyndItemCode string  `json:"fynd_item_code"`
            Company CompanyMeta  `json:"company"`
            Dimension DimensionResponse  `json:"dimension"`
            Manufacturer ManufacturerResponse  `json:"manufacturer"`
         
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
    
    // WeightResponse1 used by Catalog
    type WeightResponse1 struct {

        
            Unit string  `json:"unit"`
            Shipping float64  `json:"shipping"`
         
    }
    
    // PriceArticle used by Catalog
    type PriceArticle struct {

        
            Currency string  `json:"currency"`
            Transfer float64  `json:"transfer"`
            Effective float64  `json:"effective"`
            TpNotes map[string]interface{}  `json:"tp_notes"`
            Marked float64  `json:"marked"`
         
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

        
            Damaged Quantity  `json:"damaged"`
            OrderCommitted Quantity  `json:"order_committed"`
            NotAvailable Quantity  `json:"not_available"`
            Sellable Quantity  `json:"sellable"`
         
    }
    
    // ArticleStoreResponse used by Catalog
    type ArticleStoreResponse struct {

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            StoreCode string  `json:"store_code"`
            StoreType string  `json:"store_type"`
         
    }
    
    // BrandMeta1 used by Catalog
    type BrandMeta1 struct {

        
            Name string  `json:"name"`
            ID float64  `json:"id"`
         
    }
    
    // DateMeta used by Catalog
    type DateMeta struct {

        
            InventoryUpdatedOn string  `json:"inventory_updated_on"`
            CreatedOn string  `json:"created_on"`
            AddedOnStore string  `json:"added_on_store"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // CompanyMeta1 used by Catalog
    type CompanyMeta1 struct {

        
            ID float64  `json:"id"`
         
    }
    
    // DimensionResponse1 used by Catalog
    type DimensionResponse1 struct {

        
            Length float64  `json:"length"`
            Width float64  `json:"width"`
            Unit string  `json:"unit"`
            Height float64  `json:"height"`
         
    }
    
    // ManufacturerResponse1 used by Catalog
    type ManufacturerResponse1 struct {

        
            Address string  `json:"address"`
            Name string  `json:"name"`
            IsDefault bool  `json:"is_default"`
         
    }
    
    // GetInventories used by Catalog
    type GetInventories struct {

        
            ReturnConfig ReturnConfig2  `json:"return_config"`
            Stage string  `json:"stage"`
            SellerIdentifier string  `json:"seller_identifier"`
            Weight WeightResponse1  `json:"weight"`
            IsSet bool  `json:"is_set"`
            TotalQuantity float64  `json:"total_quantity"`
            Tags []string  `json:"tags"`
            CreatedBy UserSerializer  `json:"created_by"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            UID string  `json:"uid"`
            Price PriceArticle  `json:"price"`
            ExpirationDate string  `json:"expiration_date"`
            Size string  `json:"size"`
            CountryOfOrigin string  `json:"country_of_origin"`
            Identifier map[string]interface{}  `json:"identifier"`
            TraceID string  `json:"trace_id"`
            Trader []Trader2  `json:"trader"`
            Quantities QuantitiesArticle  `json:"quantities"`
            Store ArticleStoreResponse  `json:"store"`
            ID string  `json:"id"`
            InventoryUpdatedOn string  `json:"inventory_updated_on"`
            Brand BrandMeta1  `json:"brand"`
            DateMeta DateMeta  `json:"date_meta"`
            TrackInventory bool  `json:"track_inventory"`
            ItemID float64  `json:"item_id"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            Platforms map[string]interface{}  `json:"platforms"`
            Company CompanyMeta1  `json:"company"`
            Dimension DimensionResponse1  `json:"dimension"`
            Manufacturer ManufacturerResponse1  `json:"manufacturer"`
         
    }
    
    // GetInventoriesResponse used by Catalog
    type GetInventoriesResponse struct {

        
            Page Page  `json:"page"`
            Items []GetInventories  `json:"items"`
         
    }
    
    // BulkInventoryGetItems used by Catalog
    type BulkInventoryGetItems struct {

        
            CreatedBy map[string]interface{}  `json:"created_by"`
            CancelledRecords []string  `json:"cancelled_records"`
            Cancelled float64  `json:"cancelled"`
            ModifiedOn string  `json:"modified_on"`
            Stage string  `json:"stage"`
            FailedRecords []string  `json:"failed_records"`
            CreatedOn string  `json:"created_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            FilePath string  `json:"file_path"`
            Failed float64  `json:"failed"`
            CompanyID float64  `json:"company_id"`
            Succeed float64  `json:"succeed"`
            Total float64  `json:"total"`
            ID string  `json:"id"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // BulkInventoryGet used by Catalog
    type BulkInventoryGet struct {

        
            Page Page  `json:"page"`
            Items []BulkInventoryGetItems  `json:"items"`
         
    }
    
    // InventoryJobPayload used by Catalog
    type InventoryJobPayload struct {

        
            PriceEffective float64  `json:"price_effective"`
            Currency string  `json:"currency"`
            SellerIdentifier string  `json:"seller_identifier"`
            Price float64  `json:"price"`
            TraceID string  `json:"trace_id"`
            TotalQuantity float64  `json:"total_quantity"`
            ItemWeightUnitOfMeasure string  `json:"item_weight_unit_of_measure"`
            PriceMarked float64  `json:"price_marked"`
            ExpirationDate string  `json:"expiration_date"`
            Tags []string  `json:"tags"`
            ItemDimensionsUnitOfMeasure string  `json:"item_dimensions_unit_of_measure"`
            Quantity float64  `json:"quantity"`
            StoreCode string  `json:"store_code"`
         
    }
    
    // InventoryBulkRequest used by Catalog
    type InventoryBulkRequest struct {

        
            User map[string]interface{}  `json:"user"`
            CompanyID float64  `json:"company_id"`
            Sizes []InventoryJobPayload  `json:"sizes"`
            BatchID string  `json:"batch_id"`
         
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
            ModifiedOn string  `json:"modified_on"`
            NotificationEmails []string  `json:"notification_emails"`
            CreatedOn string  `json:"created_on"`
            TaskID string  `json:"task_id"`
            SellerID float64  `json:"seller_id"`
            Type string  `json:"type"`
            Status string  `json:"status"`
            Filters map[string]interface{}  `json:"filters"`
         
    }
    
    // InventoryExportQuantityFilter used by Catalog
    type InventoryExportQuantityFilter struct {

        
            Operators string  `json:"operators"`
            Max float64  `json:"max"`
            Min float64  `json:"min"`
         
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

        
            URL string  `json:"url"`
            NotificationEmails []string  `json:"notification_emails"`
            TaskID string  `json:"task_id"`
            SellerID float64  `json:"seller_id"`
            Type string  `json:"type"`
            Status string  `json:"status"`
            Filters InventoryExportAdvanceOption  `json:"filters"`
            CompletedOn string  `json:"completed_on"`
         
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

        
            Filters InventoryExportFilter  `json:"filters"`
            Type string  `json:"type"`
            Data []string  `json:"data"`
            NotificationEmails []string  `json:"notification_emails"`
         
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

        
            CreatedBy UserDetail  `json:"created_by"`
            ModifiedOn string  `json:"modified_on"`
            URL string  `json:"url"`
            NotificationEmails []string  `json:"notification_emails"`
            ID string  `json:"id"`
            CreatedOn string  `json:"created_on"`
            TaskID string  `json:"task_id"`
            CancelledOn string  `json:"cancelled_on"`
            SellerID float64  `json:"seller_id"`
            CancelledBy UserDetail  `json:"cancelled_by"`
            Type string  `json:"type"`
            Status interface{}  `json:"status"`
            Filters InventoryJobFilters  `json:"filters"`
            CompletedOn string  `json:"completed_on"`
         
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

        
            PriceEffective float64  `json:"price_effective"`
            StoreID float64  `json:"store_id"`
            SellerIdentifier string  `json:"seller_identifier"`
            TraceID string  `json:"trace_id"`
            TotalQuantity float64  `json:"total_quantity"`
            PriceMarked float64  `json:"price_marked"`
            ExpirationDate string  `json:"expiration_date"`
            Tags []string  `json:"tags"`
         
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

        
            TaxOnEsp bool  `json:"tax_on_esp"`
            Tax1 float64  `json:"tax1"`
            Threshold2 float64  `json:"threshold2"`
            UID float64  `json:"uid"`
            Threshold1 float64  `json:"threshold1"`
            CompanyID float64  `json:"company_id"`
            Tax2 float64  `json:"tax2"`
            TaxOnMrp bool  `json:"tax_on_mrp"`
            HsnCode string  `json:"hsn_code"`
            Hs2Code string  `json:"hs2_code"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // HsnCodesObject used by Catalog
    type HsnCodesObject struct {

        
            TaxOnEsp bool  `json:"tax_on_esp"`
            Tax1 float64  `json:"tax1"`
            ModifiedOn string  `json:"modified_on"`
            Threshold2 float64  `json:"threshold2"`
            Threshold1 float64  `json:"threshold1"`
            CompanyID float64  `json:"company_id"`
            Tax2 float64  `json:"tax2"`
            TaxOnMrp bool  `json:"tax_on_mrp"`
            HsnCode string  `json:"hsn_code"`
            Hs2Code string  `json:"hs2_code"`
            ID string  `json:"id"`
         
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

        
            Current string  `json:"current"`
            HasNext bool  `json:"has_next"`
            ItemTotal float64  `json:"item_total"`
            HasPrevious bool  `json:"has_previous"`
            Size float64  `json:"size"`
         
    }
    
    // TaxSlab used by Catalog
    type TaxSlab struct {

        
            Threshold float64  `json:"threshold"`
            Rate float64  `json:"rate"`
            Cess float64  `json:"cess"`
            EffectiveDate string  `json:"effective_date"`
         
    }
    
    // HSNDataInsertV2 used by Catalog
    type HSNDataInsertV2 struct {

        
            CreatedBy map[string]interface{}  `json:"created_by"`
            ModifiedOn string  `json:"modified_on"`
            CountryCode string  `json:"country_code"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            CreatedOn string  `json:"created_on"`
            ReportingHsn string  `json:"reporting_hsn"`
            Type string  `json:"type"`
            HsnCode string  `json:"hsn_code"`
            Taxes []TaxSlab  `json:"taxes"`
            Description string  `json:"description"`
            HsnCodeID string  `json:"hsn_code_id"`
         
    }
    
    // HsnCodesListingResponseSchemaV2 used by Catalog
    type HsnCodesListingResponseSchemaV2 struct {

        
            Page PageResponse  `json:"page"`
            Items []HSNDataInsertV2  `json:"items"`
         
    }
    
    // BrandItem used by Catalog
    type BrandItem struct {

        
            Logo Media  `json:"logo"`
            Departments []string  `json:"departments"`
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            Discount string  `json:"discount"`
            Banners ImageUrls  `json:"banners"`
            Action Action  `json:"action"`
         
    }
    
    // BrandListingResponse used by Catalog
    type BrandListingResponse struct {

        
            Page Page  `json:"page"`
            Items []BrandItem  `json:"items"`
         
    }
    
    // Department used by Catalog
    type Department struct {

        
            PriorityOrder float64  `json:"priority_order"`
            Logo Media  `json:"logo"`
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // DepartmentResponse used by Catalog
    type DepartmentResponse struct {

        
            Items []Department  `json:"items"`
         
    }
    
    // ThirdLevelChild used by Catalog
    type ThirdLevelChild struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Slug string  `json:"slug"`
            Childs []map[string]interface{}  `json:"childs"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            Banners ImageUrls  `json:"banners"`
            Action Action  `json:"action"`
         
    }
    
    // SecondLevelChild used by Catalog
    type SecondLevelChild struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Slug string  `json:"slug"`
            Childs []ThirdLevelChild  `json:"childs"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            Banners ImageUrls  `json:"banners"`
            Action Action  `json:"action"`
         
    }
    
    // Child used by Catalog
    type Child struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Slug string  `json:"slug"`
            Childs []SecondLevelChild  `json:"childs"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            Banners ImageUrls  `json:"banners"`
            Action Action  `json:"action"`
         
    }
    
    // CategoryItems used by Catalog
    type CategoryItems struct {

        
            Slug string  `json:"slug"`
            Childs []Child  `json:"childs"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            Banners ImageUrls  `json:"banners"`
            Action Action  `json:"action"`
         
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
            SortOn []ProductSortOn  `json:"sort_on"`
            Items []ProductListingDetail  `json:"items"`
            Operators map[string]interface{}  `json:"operators"`
            Filters []ProductFilters  `json:"filters"`
         
    }
    
    // ProductDetail used by Catalog
    type ProductDetail struct {

        
            ProductOnlineDate string  `json:"product_online_date"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            Tryouts []string  `json:"tryouts"`
            Slug string  `json:"slug"`
            HasVariant bool  `json:"has_variant"`
            UID float64  `json:"uid"`
            ItemType string  `json:"item_type"`
            Type string  `json:"type"`
            ShortDescription string  `json:"short_description"`
            Attributes map[string]interface{}  `json:"attributes"`
            PromoMeta map[string]interface{}  `json:"promo_meta"`
            Similars []string  `json:"similars"`
            Color string  `json:"color"`
            ItemCode string  `json:"item_code"`
            Name string  `json:"name"`
            ImageNature string  `json:"image_nature"`
            Description string  `json:"description"`
            Medias []Media1  `json:"medias"`
            Brand ProductBrand  `json:"brand"`
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
            Highlights []string  `json:"highlights"`
            RatingCount float64  `json:"rating_count"`
            Rating float64  `json:"rating"`
         
    }
    
    // InventoryPage used by Catalog
    type InventoryPage struct {

        
            HasNext bool  `json:"has_next"`
            ItemTotal float64  `json:"item_total"`
            NextID string  `json:"next_id"`
            Type string  `json:"type"`
            HasPrevious bool  `json:"has_previous"`
         
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

        
            ArticleAssignment ArticleAssignment  `json:"article_assignment"`
            Query ArticleQuery  `json:"query"`
            Meta map[string]interface{}  `json:"meta"`
            GroupID string  `json:"group_id"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // AssignStore used by Catalog
    type AssignStore struct {

        
            StoreIds []float64  `json:"store_ids"`
            ChannelType string  `json:"channel_type"`
            ChannelIdentifier string  `json:"channel_identifier"`
            CompanyID float64  `json:"company_id"`
            Pincode string  `json:"pincode"`
            Articles []AssignStoreArticle  `json:"articles"`
            AppID string  `json:"app_id"`
         
    }
    
    // ArticleAssignment1 used by Catalog
    type ArticleAssignment1 struct {

        
            Level string  `json:"level"`
            Strategy string  `json:"strategy"`
         
    }
    
    // StoreAssignResponse used by Catalog
    type StoreAssignResponse struct {

        
            PriceEffective float64  `json:"price_effective"`
            StoreID float64  `json:"store_id"`
            ItemID float64  `json:"item_id"`
            Index float64  `json:"index"`
            SCity string  `json:"s_city"`
            StrategyWiseListing []map[string]interface{}  `json:"strategy_wise_listing"`
            ArticleAssignment ArticleAssignment1  `json:"article_assignment"`
            Quantity float64  `json:"quantity"`
            StorePincode float64  `json:"store_pincode"`
            UID string  `json:"uid"`
            CompanyID float64  `json:"company_id"`
            Status bool  `json:"status"`
            Meta map[string]interface{}  `json:"meta"`
            PriceMarked float64  `json:"price_marked"`
            GroupID string  `json:"group_id"`
            ID string  `json:"_id"`
            Size string  `json:"size"`
         
    }
    
    // Document used by Catalog
    type Document struct {

        
            Verified bool  `json:"verified"`
            URL string  `json:"url"`
            Value string  `json:"value"`
            LegalName string  `json:"legal_name"`
            Type string  `json:"type"`
         
    }
    
    // LocationTimingSerializer used by Catalog
    type LocationTimingSerializer struct {

        
            Minute float64  `json:"minute"`
            Hour float64  `json:"hour"`
         
    }
    
    // LocationDayWiseSerializer used by Catalog
    type LocationDayWiseSerializer struct {

        
            Opening LocationTimingSerializer  `json:"opening"`
            Closing LocationTimingSerializer  `json:"closing"`
            Weekday string  `json:"weekday"`
            Open bool  `json:"open"`
         
    }
    
    // UserSerializer1 used by Catalog
    type UserSerializer1 struct {

        
            Contact string  `json:"contact"`
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
         
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
    
    // SellerPhoneNumber used by Catalog
    type SellerPhoneNumber struct {

        
            CountryCode float64  `json:"country_code"`
            Number string  `json:"number"`
         
    }
    
    // LocationManagerSerializer used by Catalog
    type LocationManagerSerializer struct {

        
            Email string  `json:"email"`
            Name string  `json:"name"`
            MobileNo SellerPhoneNumber  `json:"mobile_no"`
         
    }
    
    // LocationIntegrationType used by Catalog
    type LocationIntegrationType struct {

        
            Inventory string  `json:"inventory"`
            Order string  `json:"order"`
         
    }
    
    // GetAddressSerializer used by Catalog
    type GetAddressSerializer struct {

        
            Address2 string  `json:"address2"`
            Country string  `json:"country"`
            State string  `json:"state"`
            Address1 string  `json:"address1"`
            City string  `json:"city"`
            CountryCode string  `json:"country_code"`
            Longitude float64  `json:"longitude"`
            Pincode float64  `json:"pincode"`
            Landmark string  `json:"landmark"`
            AddressType string  `json:"address_type"`
            Latitude float64  `json:"latitude"`
         
    }
    
    // ProductReturnConfigSerializer used by Catalog
    type ProductReturnConfigSerializer struct {

        
            OnSameStore bool  `json:"on_same_store"`
            StoreUID float64  `json:"store_uid"`
         
    }
    
    // UserSerializer2 used by Catalog
    type UserSerializer2 struct {

        
            Contact string  `json:"contact"`
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
         
    }
    
    // GetCompanySerializer used by Catalog
    type GetCompanySerializer struct {

        
            CreatedBy UserSerializer2  `json:"created_by"`
            BusinessType string  `json:"business_type"`
            ModifiedOn string  `json:"modified_on"`
            Stage string  `json:"stage"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            ModifiedBy UserSerializer2  `json:"modified_by"`
            CreatedOn string  `json:"created_on"`
            VerifiedOn string  `json:"verified_on"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            CompanyType string  `json:"company_type"`
            RejectReason string  `json:"reject_reason"`
            VerifiedBy UserSerializer2  `json:"verified_by"`
         
    }
    
    // GetLocationSerializer used by Catalog
    type GetLocationSerializer struct {

        
            Stage string  `json:"stage"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Documents []Document  `json:"documents"`
            Timing []LocationDayWiseSerializer  `json:"timing"`
            CreatedBy UserSerializer1  `json:"created_by"`
            ModifiedOn string  `json:"modified_on"`
            ModifiedBy UserSerializer1  `json:"modified_by"`
            CreatedOn string  `json:"created_on"`
            VerifiedOn string  `json:"verified_on"`
            Warnings map[string]interface{}  `json:"warnings"`
            PhoneNumber string  `json:"phone_number"`
            UID float64  `json:"uid"`
            GstCredentials InvoiceDetailsSerializer  `json:"gst_credentials"`
            DisplayName string  `json:"display_name"`
            Manager LocationManagerSerializer  `json:"manager"`
            IntegrationType LocationIntegrationType  `json:"integration_type"`
            Address GetAddressSerializer  `json:"address"`
            Name string  `json:"name"`
            StoreType string  `json:"store_type"`
            VerifiedBy UserSerializer1  `json:"verified_by"`
            NotificationEmails []string  `json:"notification_emails"`
            ProductReturnConfig ProductReturnConfigSerializer  `json:"product_return_config"`
            Code string  `json:"code"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            Company GetCompanySerializer  `json:"company"`
         
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

        
            Logo string  `json:"logo"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            IsActive bool  `json:"is_active"`
            AppID string  `json:"app_id"`
         
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
    
    // GetAddressSerializer used by CompanyProfile
    type GetAddressSerializer struct {

        
            State string  `json:"state"`
            Landmark string  `json:"landmark"`
            CountryCode string  `json:"country_code"`
            Address1 string  `json:"address1"`
            AddressType string  `json:"address_type"`
            Address2 string  `json:"address2"`
            Country string  `json:"country"`
            Latitude float64  `json:"latitude"`
            Longitude float64  `json:"longitude"`
            City string  `json:"city"`
            Pincode float64  `json:"pincode"`
         
    }
    
    // CompanyTaxesSerializer used by CompanyProfile
    type CompanyTaxesSerializer struct {

        
            EffectiveDate string  `json:"effective_date"`
            Enable bool  `json:"enable"`
            Rate float64  `json:"rate"`
         
    }
    
    // Document used by CompanyProfile
    type Document struct {

        
            Value string  `json:"value"`
            Type string  `json:"type"`
            LegalName string  `json:"legal_name"`
            Verified bool  `json:"verified"`
            URL string  `json:"url"`
         
    }
    
    // BusinessCountryInfo used by CompanyProfile
    type BusinessCountryInfo struct {

        
            Country string  `json:"country"`
            CountryCode string  `json:"country_code"`
         
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
    
    // GetCompanyProfileSerializerResponse used by CompanyProfile
    type GetCompanyProfileSerializerResponse struct {

        
            BusinessType string  `json:"business_type"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            BusinessDetails BusinessDetails  `json:"business_details"`
            NotificationEmails []string  `json:"notification_emails"`
            FranchiseEnabled bool  `json:"franchise_enabled"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            ModifiedOn string  `json:"modified_on"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            BusinessInfo string  `json:"business_info"`
            VerifiedOn string  `json:"verified_on"`
            Taxes []CompanyTaxesSerializer  `json:"taxes"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Documents []Document  `json:"documents"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            CreatedOn string  `json:"created_on"`
            CreatedBy UserSerializer  `json:"created_by"`
            Warnings map[string]interface{}  `json:"warnings"`
            BusinessCountryInfo BusinessCountryInfo  `json:"business_country_info"`
            ContactDetails ContactDetails  `json:"contact_details"`
            Stage string  `json:"stage"`
            CompanyType string  `json:"company_type"`
            Mode string  `json:"mode"`
         
    }
    
    // ErrorResponse used by CompanyProfile
    type ErrorResponse struct {

        
            Code string  `json:"code"`
            Meta map[string]interface{}  `json:"meta"`
            Status float64  `json:"status"`
            Message string  `json:"message"`
         
    }
    
    // CreateUpdateAddressSerializer used by CompanyProfile
    type CreateUpdateAddressSerializer struct {

        
            State string  `json:"state"`
            CountryCode string  `json:"country_code"`
            Landmark string  `json:"landmark"`
            Address1 string  `json:"address1"`
            AddressType string  `json:"address_type"`
            Country string  `json:"country"`
            Address2 string  `json:"address2"`
            Latitude float64  `json:"latitude"`
            Longitude float64  `json:"longitude"`
            City string  `json:"city"`
            Pincode float64  `json:"pincode"`
         
    }
    
    // CompanyTaxesSerializer1 used by CompanyProfile
    type CompanyTaxesSerializer1 struct {

        
            EffectiveDate string  `json:"effective_date"`
            Enable bool  `json:"enable"`
            Rate float64  `json:"rate"`
         
    }
    
    // UpdateCompany used by CompanyProfile
    type UpdateCompany struct {

        
            BusinessType string  `json:"business_type"`
            Warnings map[string]interface{}  `json:"warnings"`
            Addresses []CreateUpdateAddressSerializer  `json:"addresses"`
            FranchiseEnabled bool  `json:"franchise_enabled"`
            Documents []Document  `json:"documents"`
            Name string  `json:"name"`
            BusinessInfo string  `json:"business_info"`
            ContactDetails ContactDetails  `json:"contact_details"`
            Taxes []CompanyTaxesSerializer1  `json:"taxes"`
            CompanyType string  `json:"company_type"`
            RejectReason string  `json:"reject_reason"`
            BusinessDetails BusinessDetails  `json:"business_details"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            NotificationEmails []string  `json:"notification_emails"`
         
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

        
            Store DocumentsObj  `json:"store"`
            UID float64  `json:"uid"`
            CompanyDocuments DocumentsObj  `json:"company_documents"`
            StoreDocuments DocumentsObj  `json:"store_documents"`
            Product DocumentsObj  `json:"product"`
            Stage string  `json:"stage"`
            Brand DocumentsObj  `json:"brand"`
         
    }
    
    // BrandBannerSerializer used by CompanyProfile
    type BrandBannerSerializer struct {

        
            Landscape string  `json:"landscape"`
            Portrait string  `json:"portrait"`
         
    }
    
    // GetBrandResponseSerializer used by CompanyProfile
    type GetBrandResponseSerializer struct {

        
            ModifiedBy UserSerializer  `json:"modified_by"`
            Banner BrandBannerSerializer  `json:"banner"`
            ModifiedOn string  `json:"modified_on"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            VerifiedOn string  `json:"verified_on"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Synonyms []string  `json:"synonyms"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            Description string  `json:"description"`
            CreatedOn string  `json:"created_on"`
            Logo string  `json:"logo"`
            SlugKey string  `json:"slug_key"`
            CreatedBy UserSerializer  `json:"created_by"`
            Warnings map[string]interface{}  `json:"warnings"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            Stage string  `json:"stage"`
            RejectReason string  `json:"reject_reason"`
            Mode string  `json:"mode"`
         
    }
    
    // CreateUpdateBrandRequestSerializer used by CompanyProfile
    type CreateUpdateBrandRequestSerializer struct {

        
            BrandTier string  `json:"brand_tier"`
            Name string  `json:"name"`
            UID float64  `json:"uid"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            Description string  `json:"description"`
            CompanyID float64  `json:"company_id"`
            Banner BrandBannerSerializer  `json:"banner"`
            Logo string  `json:"logo"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Synonyms []string  `json:"synonyms"`
         
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

        
            BusinessType string  `json:"business_type"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            Details CompanyDetails  `json:"details"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            ModifiedOn string  `json:"modified_on"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            RejectReason string  `json:"reject_reason"`
            BusinessCountryInfo BusinessCountryInfo  `json:"business_country_info"`
            VerifiedOn string  `json:"verified_on"`
            Stage string  `json:"stage"`
            CompanyType string  `json:"company_type"`
            CreatedOn string  `json:"created_on"`
            MarketChannels []string  `json:"market_channels"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            NotificationEmails []string  `json:"notification_emails"`
            CreatedBy UserSerializer  `json:"created_by"`
         
    }
    
    // CompanyBrandSerializer used by CompanyProfile
    type CompanyBrandSerializer struct {

        
            Company CompanySerializer  `json:"company"`
            Warnings map[string]interface{}  `json:"warnings"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            ModifiedOn string  `json:"modified_on"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            UID float64  `json:"uid"`
            RejectReason string  `json:"reject_reason"`
            VerifiedOn string  `json:"verified_on"`
            Stage string  `json:"stage"`
            CreatedOn string  `json:"created_on"`
            CreatedBy UserSerializer  `json:"created_by"`
            Brand GetBrandResponseSerializer  `json:"brand"`
         
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
            Brands []float64  `json:"brands"`
            UID float64  `json:"uid"`
         
    }
    
    // ProductReturnConfigSerializer used by CompanyProfile
    type ProductReturnConfigSerializer struct {

        
            StoreUID float64  `json:"store_uid"`
            OnSameStore bool  `json:"on_same_store"`
         
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
    
    // GetCompanySerializer used by CompanyProfile
    type GetCompanySerializer struct {

        
            BusinessType string  `json:"business_type"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            ModifiedOn string  `json:"modified_on"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            RejectReason string  `json:"reject_reason"`
            VerifiedOn string  `json:"verified_on"`
            Stage string  `json:"stage"`
            CompanyType string  `json:"company_type"`
            CreatedOn string  `json:"created_on"`
            CreatedBy UserSerializer  `json:"created_by"`
         
    }
    
    // LocationManagerSerializer used by CompanyProfile
    type LocationManagerSerializer struct {

        
            Email string  `json:"email"`
            MobileNo SellerPhoneNumber  `json:"mobile_no"`
            Name string  `json:"name"`
         
    }
    
    // LocationTimingSerializer used by CompanyProfile
    type LocationTimingSerializer struct {

        
            Hour float64  `json:"hour"`
            Minute float64  `json:"minute"`
         
    }
    
    // LocationDayWiseSerializer used by CompanyProfile
    type LocationDayWiseSerializer struct {

        
            Weekday string  `json:"weekday"`
            Opening LocationTimingSerializer  `json:"opening"`
            Closing LocationTimingSerializer  `json:"closing"`
            Open bool  `json:"open"`
         
    }
    
    // GetLocationSerializer used by CompanyProfile
    type GetLocationSerializer struct {

        
            Address GetAddressSerializer  `json:"address"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            NotificationEmails []string  `json:"notification_emails"`
            PhoneNumber string  `json:"phone_number"`
            StoreType string  `json:"store_type"`
            ModifiedOn string  `json:"modified_on"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            ProductReturnConfig ProductReturnConfigSerializer  `json:"product_return_config"`
            GstCredentials InvoiceDetailsSerializer  `json:"gst_credentials"`
            VerifiedOn string  `json:"verified_on"`
            Holiday []HolidaySchemaSerializer  `json:"holiday"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Company GetCompanySerializer  `json:"company"`
            Documents []Document  `json:"documents"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            Code string  `json:"code"`
            Manager LocationManagerSerializer  `json:"manager"`
            CreatedOn string  `json:"created_on"`
            CreatedBy UserSerializer  `json:"created_by"`
            Warnings map[string]interface{}  `json:"warnings"`
            Timing []LocationDayWiseSerializer  `json:"timing"`
            DisplayName string  `json:"display_name"`
            Stage string  `json:"stage"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
         
    }
    
    // LocationListSerializer used by CompanyProfile
    type LocationListSerializer struct {

        
            Items []GetLocationSerializer  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // AddressSerializer used by CompanyProfile
    type AddressSerializer struct {

        
            State string  `json:"state"`
            Landmark string  `json:"landmark"`
            CountryCode string  `json:"country_code"`
            Address1 string  `json:"address1"`
            AddressType string  `json:"address_type"`
            Address2 string  `json:"address2"`
            Country string  `json:"country"`
            Latitude float64  `json:"latitude"`
            Longitude float64  `json:"longitude"`
            City string  `json:"city"`
            Pincode float64  `json:"pincode"`
         
    }
    
    // LocationSerializer used by CompanyProfile
    type LocationSerializer struct {

        
            Company float64  `json:"company"`
            Address AddressSerializer  `json:"address"`
            Timing []LocationDayWiseSerializer  `json:"timing"`
            DisplayName string  `json:"display_name"`
            StoreType string  `json:"store_type"`
            Documents []Document  `json:"documents"`
            Name string  `json:"name"`
            UID float64  `json:"uid"`
            Warnings map[string]interface{}  `json:"warnings"`
            ProductReturnConfig ProductReturnConfigSerializer  `json:"product_return_config"`
            Code string  `json:"code"`
            GstCredentials InvoiceDetailsSerializer  `json:"gst_credentials"`
            Stage string  `json:"stage"`
            Manager LocationManagerSerializer  `json:"manager"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            Holiday []HolidaySchemaSerializer  `json:"holiday"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            NotificationEmails []string  `json:"notification_emails"`
         
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
    
    // Ownership used by Cart
    type Ownership struct {

        
            PayableCategory string  `json:"payable_category"`
            PayableBy string  `json:"payable_by"`
         
    }
    
    // Rule used by Cart
    type Rule struct {

        
            DiscountQty float64  `json:"discount_qty"`
            Max float64  `json:"max"`
            Value float64  `json:"value"`
            Min float64  `json:"min"`
            Key float64  `json:"key"`
         
    }
    
    // RuleDefinition used by Cart
    type RuleDefinition struct {

        
            CalculateOn string  `json:"calculate_on"`
            CurrencyCode string  `json:"currency_code"`
            AutoApply bool  `json:"auto_apply"`
            ApplicableOn string  `json:"applicable_on"`
            Type string  `json:"type"`
            IsExact bool  `json:"is_exact"`
            Scope []string  `json:"scope"`
            ValueType string  `json:"value_type"`
         
    }
    
    // Validity used by Cart
    type Validity struct {

        
            Priority float64  `json:"priority"`
         
    }
    
    // DisplayMetaDict used by Cart
    type DisplayMetaDict struct {

        
            Title string  `json:"title"`
            Subtitle string  `json:"subtitle"`
         
    }
    
    // DisplayMeta used by Cart
    type DisplayMeta struct {

        
            Apply DisplayMetaDict  `json:"apply"`
            Subtitle string  `json:"subtitle"`
            Remove DisplayMetaDict  `json:"remove"`
            Auto DisplayMetaDict  `json:"auto"`
            Title string  `json:"title"`
            Description string  `json:"description"`
         
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
    
    // BulkBundleRestriction used by Cart
    type BulkBundleRestriction struct {

        
            MultiStoreAllowed bool  `json:"multi_store_allowed"`
         
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
            Codes []string  `json:"codes"`
            Networks []string  `json:"networks"`
         
    }
    
    // PriceRange used by Cart
    type PriceRange struct {

        
            Min float64  `json:"min"`
            Max float64  `json:"max"`
         
    }
    
    // Restrictions used by Cart
    type Restrictions struct {

        
            CouponAllowed bool  `json:"coupon_allowed"`
            Uses UsesRestriction  `json:"uses"`
            UserGroups []float64  `json:"user_groups"`
            BulkBundle BulkBundleRestriction  `json:"bulk_bundle"`
            Platforms []string  `json:"platforms"`
            PostOrder PostOrder  `json:"post_order"`
            OrderingStores []float64  `json:"ordering_stores"`
            Payments map[string]PaymentModes  `json:"payments"`
            PriceRange PriceRange  `json:"price_range"`
         
    }
    
    // CouponAuthor used by Cart
    type CouponAuthor struct {

        
            ModifiedBy string  `json:"modified_by"`
            CreatedBy string  `json:"created_by"`
         
    }
    
    // CouponSchedule used by Cart
    type CouponSchedule struct {

        
            End string  `json:"end"`
            Cron string  `json:"cron"`
            Duration float64  `json:"duration"`
            Start string  `json:"start"`
            NextSchedule []map[string]interface{}  `json:"next_schedule"`
         
    }
    
    // State used by Cart
    type State struct {

        
            IsDisplay bool  `json:"is_display"`
            IsArchived bool  `json:"is_archived"`
            IsPublic bool  `json:"is_public"`
         
    }
    
    // CouponDateMeta used by Cart
    type CouponDateMeta struct {

        
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
         
    }
    
    // Validation used by Cart
    type Validation struct {

        
            UserRegisteredAfter string  `json:"user_registered_after"`
            Anonymous bool  `json:"anonymous"`
            AppID []string  `json:"app_id"`
         
    }
    
    // Identifier used by Cart
    type Identifier struct {

        
            ArticleID []string  `json:"article_id"`
            StoreID []float64  `json:"store_id"`
            CategoryID []float64  `json:"category_id"`
            CompanyID []float64  `json:"company_id"`
            BrandID []float64  `json:"brand_id"`
            ExcludeBrandID []float64  `json:"exclude_brand_id"`
            UserID []string  `json:"user_id"`
            CollectionID []string  `json:"collection_id"`
            ItemID []float64  `json:"item_id"`
         
    }
    
    // CouponAction used by Cart
    type CouponAction struct {

        
            ActionDate string  `json:"action_date"`
            TxnMode string  `json:"txn_mode"`
         
    }
    
    // CouponAdd used by Cart
    type CouponAdd struct {

        
            Ownership Ownership  `json:"ownership"`
            TypeSlug string  `json:"type_slug"`
            Rule []Rule  `json:"rule"`
            RuleDefinition RuleDefinition  `json:"rule_definition"`
            Validity Validity  `json:"validity"`
            DisplayMeta DisplayMeta  `json:"display_meta"`
            Code string  `json:"code"`
            Restrictions Restrictions  `json:"restrictions"`
            Author CouponAuthor  `json:"author"`
            Schedule CouponSchedule  `json:"_schedule"`
            Tags []string  `json:"tags"`
            State State  `json:"state"`
            DateMeta CouponDateMeta  `json:"date_meta"`
            Validation Validation  `json:"validation"`
            Identifiers Identifier  `json:"identifiers"`
            Action CouponAction  `json:"action"`
         
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

        
            Ownership Ownership  `json:"ownership"`
            TypeSlug string  `json:"type_slug"`
            Rule []Rule  `json:"rule"`
            RuleDefinition RuleDefinition  `json:"rule_definition"`
            Validity Validity  `json:"validity"`
            DisplayMeta DisplayMeta  `json:"display_meta"`
            Code string  `json:"code"`
            Restrictions Restrictions  `json:"restrictions"`
            Author CouponAuthor  `json:"author"`
            Schedule CouponSchedule  `json:"_schedule"`
            Tags []string  `json:"tags"`
            State State  `json:"state"`
            DateMeta CouponDateMeta  `json:"date_meta"`
            Validation Validation  `json:"validation"`
            Identifiers Identifier  `json:"identifiers"`
            Action CouponAction  `json:"action"`
         
    }
    
    // CouponPartialUpdate used by Cart
    type CouponPartialUpdate struct {

        
            Archive bool  `json:"archive"`
            Schedule CouponSchedule  `json:"schedule"`
         
    }
    
    // PromotionDateMeta used by Cart
    type PromotionDateMeta struct {

        
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
         
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
    
    // CompareObject used by Cart
    type CompareObject struct {

        
            Equals float64  `json:"equals"`
            GreaterThanEquals float64  `json:"greater_than_equals"`
            LessThan float64  `json:"less_than"`
            GreaterThan float64  `json:"greater_than"`
            LessThanEquals float64  `json:"less_than_equals"`
         
    }
    
    // ItemCriteria used by Cart
    type ItemCriteria struct {

        
            ItemStore []float64  `json:"item_store"`
            BuyRules []string  `json:"buy_rules"`
            ItemExcludeStore []float64  `json:"item_exclude_store"`
            ItemExcludeCategory []float64  `json:"item_exclude_category"`
            ItemExcludeSku []string  `json:"item_exclude_sku"`
            ItemBrand []float64  `json:"item_brand"`
            ItemSize []string  `json:"item_size"`
            CartTotal CompareObject  `json:"cart_total"`
            AllItems bool  `json:"all_items"`
            ItemExcludeCompany []float64  `json:"item_exclude_company"`
            ItemCategory []float64  `json:"item_category"`
            ItemID []float64  `json:"item_id"`
            ItemCompany []float64  `json:"item_company"`
            ItemSku []string  `json:"item_sku"`
            ItemExcludeBrand []float64  `json:"item_exclude_brand"`
            CartQuantity CompareObject  `json:"cart_quantity"`
         
    }
    
    // PromotionSchedule used by Cart
    type PromotionSchedule struct {

        
            End string  `json:"end"`
            Cron string  `json:"cron"`
            Duration float64  `json:"duration"`
            Published bool  `json:"published"`
            Start string  `json:"start"`
            NextSchedule []map[string]interface{}  `json:"next_schedule"`
         
    }
    
    // DisplayMeta1 used by Cart
    type DisplayMeta1 struct {

        
            OfferText string  `json:"offer_text"`
            Description string  `json:"description"`
            Name string  `json:"name"`
         
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
    
    // PostOrder1 used by Cart
    type PostOrder1 struct {

        
            ReturnAllowed bool  `json:"return_allowed"`
            CancellationAllowed bool  `json:"cancellation_allowed"`
         
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
    
    // Restrictions1 used by Cart
    type Restrictions1 struct {

        
            Uses UsesRestriction1  `json:"uses"`
            UserRegistered UserRegistered  `json:"user_registered"`
            OrderQuantity float64  `json:"order_quantity"`
            UserGroups []float64  `json:"user_groups"`
            Platforms []string  `json:"platforms"`
            PostOrder PostOrder1  `json:"post_order"`
            UserID []string  `json:"user_id"`
            Payments []PromotionPaymentModes  `json:"payments"`
            AnonymousUsers bool  `json:"anonymous_users"`
         
    }
    
    // PromotionAuthor used by Cart
    type PromotionAuthor struct {

        
            ModifiedBy string  `json:"modified_by"`
            CreatedBy string  `json:"created_by"`
         
    }
    
    // DiscountOffer used by Cart
    type DiscountOffer struct {

        
            Code string  `json:"code"`
            DiscountAmount float64  `json:"discount_amount"`
            MinOfferQuantity float64  `json:"min_offer_quantity"`
            DiscountPrice float64  `json:"discount_price"`
            DiscountPercentage float64  `json:"discount_percentage"`
            MaxOfferQuantity float64  `json:"max_offer_quantity"`
            MaxDiscountAmount float64  `json:"max_discount_amount"`
         
    }
    
    // DiscountRule used by Cart
    type DiscountRule struct {

        
            Offer DiscountOffer  `json:"offer"`
            ItemCriteria ItemCriteria  `json:"item_criteria"`
            DiscountType string  `json:"discount_type"`
            BuyCondition string  `json:"buy_condition"`
         
    }
    
    // Ownership1 used by Cart
    type Ownership1 struct {

        
            PayableCategory string  `json:"payable_category"`
            PayableBy string  `json:"payable_by"`
         
    }
    
    // PromotionListItem used by Cart
    type PromotionListItem struct {

        
            ApplyAllDiscount bool  `json:"apply_all_discount"`
            PromotionType string  `json:"promotion_type"`
            DateMeta PromotionDateMeta  `json:"date_meta"`
            Visiblility Visibility  `json:"visiblility"`
            PostOrderAction PromotionAction  `json:"post_order_action"`
            BuyRules map[string]ItemCriteria  `json:"buy_rules"`
            Stackable bool  `json:"stackable"`
            ApplicationID string  `json:"application_id"`
            Schedule PromotionSchedule  `json:"_schedule"`
            DisplayMeta DisplayMeta1  `json:"display_meta"`
            ApplyPriority float64  `json:"apply_priority"`
            Restrictions Restrictions1  `json:"restrictions"`
            Author PromotionAuthor  `json:"author"`
            ApplyExclusive string  `json:"apply_exclusive"`
            PromoGroup string  `json:"promo_group"`
            DiscountRules []DiscountRule  `json:"discount_rules"`
            Ownership Ownership1  `json:"ownership"`
            Code string  `json:"code"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Currency string  `json:"currency"`
            Mode string  `json:"mode"`
         
    }
    
    // PromotionsResponse used by Cart
    type PromotionsResponse struct {

        
            Page Page  `json:"page"`
            Items PromotionListItem  `json:"items"`
         
    }
    
    // PromotionAdd used by Cart
    type PromotionAdd struct {

        
            ApplyAllDiscount bool  `json:"apply_all_discount"`
            PromotionType string  `json:"promotion_type"`
            DateMeta PromotionDateMeta  `json:"date_meta"`
            Visiblility Visibility  `json:"visiblility"`
            PostOrderAction PromotionAction  `json:"post_order_action"`
            BuyRules map[string]ItemCriteria  `json:"buy_rules"`
            Stackable bool  `json:"stackable"`
            ApplicationID string  `json:"application_id"`
            Schedule PromotionSchedule  `json:"_schedule"`
            DisplayMeta DisplayMeta1  `json:"display_meta"`
            ApplyPriority float64  `json:"apply_priority"`
            Restrictions Restrictions1  `json:"restrictions"`
            Author PromotionAuthor  `json:"author"`
            ApplyExclusive string  `json:"apply_exclusive"`
            PromoGroup string  `json:"promo_group"`
            DiscountRules []DiscountRule  `json:"discount_rules"`
            Ownership Ownership1  `json:"ownership"`
            Code string  `json:"code"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Currency string  `json:"currency"`
            Mode string  `json:"mode"`
         
    }
    
    // PromotionUpdate used by Cart
    type PromotionUpdate struct {

        
            ApplyAllDiscount bool  `json:"apply_all_discount"`
            PromotionType string  `json:"promotion_type"`
            DateMeta PromotionDateMeta  `json:"date_meta"`
            Visiblility Visibility  `json:"visiblility"`
            PostOrderAction PromotionAction  `json:"post_order_action"`
            BuyRules map[string]ItemCriteria  `json:"buy_rules"`
            Stackable bool  `json:"stackable"`
            ApplicationID string  `json:"application_id"`
            Schedule PromotionSchedule  `json:"_schedule"`
            DisplayMeta DisplayMeta1  `json:"display_meta"`
            ApplyPriority float64  `json:"apply_priority"`
            Restrictions Restrictions1  `json:"restrictions"`
            Author PromotionAuthor  `json:"author"`
            ApplyExclusive string  `json:"apply_exclusive"`
            PromoGroup string  `json:"promo_group"`
            DiscountRules []DiscountRule  `json:"discount_rules"`
            Ownership Ownership1  `json:"ownership"`
            Code string  `json:"code"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Currency string  `json:"currency"`
            Mode string  `json:"mode"`
         
    }
    
    // PromotionPartialUpdate used by Cart
    type PromotionPartialUpdate struct {

        
            Archive bool  `json:"archive"`
            Schedule PromotionSchedule  `json:"schedule"`
         
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

        
            RevenueEngineCoupon bool  `json:"revenue_engine_coupon"`
            MinCartValue float64  `json:"min_cart_value"`
            GiftPricing float64  `json:"gift_pricing"`
            GiftDisplayText string  `json:"gift_display_text"`
            MaxCartItems float64  `json:"max_cart_items"`
            BulkCoupons bool  `json:"bulk_coupons"`
            Enabled bool  `json:"enabled"`
            DeliveryCharges DeliveryCharges  `json:"delivery_charges"`
         
    }
    
    // CartMetaConfigAdd used by Cart
    type CartMetaConfigAdd struct {

        
            RevenueEngineCoupon bool  `json:"revenue_engine_coupon"`
            MinCartValue float64  `json:"min_cart_value"`
            GiftPricing float64  `json:"gift_pricing"`
            GiftDisplayText string  `json:"gift_display_text"`
            MaxCartItems float64  `json:"max_cart_items"`
            BulkCoupons bool  `json:"bulk_coupons"`
            Enabled bool  `json:"enabled"`
            DeliveryCharges DeliveryCharges  `json:"delivery_charges"`
         
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
    
    // DisplayBreakup used by Cart
    type DisplayBreakup struct {

        
            CurrencyCode string  `json:"currency_code"`
            CurrencySymbol string  `json:"currency_symbol"`
            Display string  `json:"display"`
            Value float64  `json:"value"`
            Message []string  `json:"message"`
            Key string  `json:"key"`
         
    }
    
    // RawBreakup used by Cart
    type RawBreakup struct {

        
            ConvenienceFee float64  `json:"convenience_fee"`
            Coupon float64  `json:"coupon"`
            MrpTotal float64  `json:"mrp_total"`
            Vog float64  `json:"vog"`
            Subtotal float64  `json:"subtotal"`
            Total float64  `json:"total"`
            YouSaved float64  `json:"you_saved"`
            CodCharge float64  `json:"cod_charge"`
            GstCharges float64  `json:"gst_charges"`
            FyndCash float64  `json:"fynd_cash"`
            Discount float64  `json:"discount"`
            DeliveryCharge float64  `json:"delivery_charge"`
         
    }
    
    // LoyaltyPoints used by Cart
    type LoyaltyPoints struct {

        
            Total float64  `json:"total"`
            Description string  `json:"description"`
            Applicable float64  `json:"applicable"`
            IsApplied bool  `json:"is_applied"`
         
    }
    
    // CouponBreakup used by Cart
    type CouponBreakup struct {

        
            Code string  `json:"code"`
            Type string  `json:"type"`
            Value float64  `json:"value"`
            Message string  `json:"message"`
            UID string  `json:"uid"`
            IsApplied bool  `json:"is_applied"`
         
    }
    
    // CartBreakup used by Cart
    type CartBreakup struct {

        
            Display []DisplayBreakup  `json:"display"`
            Raw RawBreakup  `json:"raw"`
            LoyaltyPoints LoyaltyPoints  `json:"loyalty_points"`
            Coupon CouponBreakup  `json:"coupon"`
         
    }
    
    // ProductPrice used by Cart
    type ProductPrice struct {

        
            CurrencyCode string  `json:"currency_code"`
            Marked float64  `json:"marked"`
            CurrencySymbol string  `json:"currency_symbol"`
            AddOn float64  `json:"add_on"`
            Effective float64  `json:"effective"`
            Selling float64  `json:"selling"`
         
    }
    
    // ProductPriceInfo used by Cart
    type ProductPriceInfo struct {

        
            Base ProductPrice  `json:"base"`
            Converted ProductPrice  `json:"converted"`
         
    }
    
    // ProductAvailabilitySize used by Cart
    type ProductAvailabilitySize struct {

        
            Value string  `json:"value"`
            Display string  `json:"display"`
            IsAvailable bool  `json:"is_available"`
         
    }
    
    // ProductAvailability used by Cart
    type ProductAvailability struct {

        
            Sizes []string  `json:"sizes"`
            OtherStoreQuantity float64  `json:"other_store_quantity"`
            IsValid bool  `json:"is_valid"`
            AvailableSizes []ProductAvailabilitySize  `json:"available_sizes"`
            OutOfStock bool  `json:"out_of_stock"`
            Deliverable bool  `json:"deliverable"`
         
    }
    
    // BaseInfo used by Cart
    type BaseInfo struct {

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // CategoryInfo used by Cart
    type CategoryInfo struct {

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // ProductImage used by Cart
    type ProductImage struct {

        
            AspectRatio string  `json:"aspect_ratio"`
            SecureURL string  `json:"secure_url"`
            URL string  `json:"url"`
         
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

        
            Brand BaseInfo  `json:"brand"`
            Categories []CategoryInfo  `json:"categories"`
            Images []ProductImage  `json:"images"`
            Name string  `json:"name"`
            Type string  `json:"type"`
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
            Action ProductAction  `json:"action"`
         
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
    
    // PromoMeta used by Cart
    type PromoMeta struct {

        
            Message string  `json:"message"`
         
    }
    
    // BasePrice used by Cart
    type BasePrice struct {

        
            Effective float64  `json:"effective"`
            CurrencyCode string  `json:"currency_code"`
            Marked float64  `json:"marked"`
            CurrencySymbol string  `json:"currency_symbol"`
         
    }
    
    // ArticlePriceInfo used by Cart
    type ArticlePriceInfo struct {

        
            Base BasePrice  `json:"base"`
            Converted BasePrice  `json:"converted"`
         
    }
    
    // ProductArticle used by Cart
    type ProductArticle struct {

        
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            Size string  `json:"size"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            Price ArticlePriceInfo  `json:"price"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Type string  `json:"type"`
            Seller BaseInfo  `json:"seller"`
            Quantity float64  `json:"quantity"`
            UID string  `json:"uid"`
            ProductGroupTags []string  `json:"product_group_tags"`
            Store BaseInfo  `json:"store"`
         
    }
    
    // Ownership2 used by Cart
    type Ownership2 struct {

        
            PayableCategory string  `json:"payable_category"`
            PayableBy string  `json:"payable_by"`
         
    }
    
    // AppliedPromotion used by Cart
    type AppliedPromotion struct {

        
            PromoID string  `json:"promo_id"`
            ArticleQuantity float64  `json:"article_quantity"`
            Ownership Ownership2  `json:"ownership"`
            Amount float64  `json:"amount"`
            PromotionType string  `json:"promotion_type"`
            MrpPromotion bool  `json:"mrp_promotion"`
            OfferText string  `json:"offer_text"`
         
    }
    
    // CartProductInfo used by Cart
    type CartProductInfo struct {

        
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            IsSet bool  `json:"is_set"`
            Price ProductPriceInfo  `json:"price"`
            CouponMessage string  `json:"coupon_message"`
            Availability ProductAvailability  `json:"availability"`
            Product CartProduct  `json:"product"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            PricePerUnit ProductPriceInfo  `json:"price_per_unit"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
            Quantity float64  `json:"quantity"`
            Message string  `json:"message"`
            PromoMeta PromoMeta  `json:"promo_meta"`
            Key string  `json:"key"`
            Discount string  `json:"discount"`
            Article ProductArticle  `json:"article"`
            Moq map[string]interface{}  `json:"moq"`
            BulkOffer map[string]interface{}  `json:"bulk_offer"`
            PromotionsApplied []AppliedPromotion  `json:"promotions_applied"`
         
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

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
            Errors map[string]interface{}  `json:"errors"`
         
    }
    
    // ShippingAddress used by Cart
    type ShippingAddress struct {

        
            Pincode float64  `json:"pincode"`
            Email string  `json:"email"`
            Landmark string  `json:"landmark"`
            Phone float64  `json:"phone"`
            City string  `json:"city"`
            AreaCode string  `json:"area_code"`
            State string  `json:"state"`
            Meta map[string]interface{}  `json:"meta"`
            AddressType string  `json:"address_type"`
            Name string  `json:"name"`
            AreaCodeSlug string  `json:"area_code_slug"`
            Area string  `json:"area"`
            CountryPhoneCode string  `json:"country_phone_code"`
            CountryIsoCode string  `json:"country_iso_code"`
            Country string  `json:"country"`
            Address string  `json:"address"`
            CountryCode string  `json:"country_code"`
         
    }
    
    // OpenApiCartServiceabilityRequest used by Cart
    type OpenApiCartServiceabilityRequest struct {

        
            ShippingAddress ShippingAddress  `json:"shipping_address"`
            CartItems []CartItem  `json:"cart_items"`
         
    }
    
    // OpenApiCartServiceabilityResponse used by Cart
    type OpenApiCartServiceabilityResponse struct {

        
            BreakupValues CartBreakup  `json:"breakup_values"`
            IsValid bool  `json:"is_valid"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            Message string  `json:"message"`
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
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            CurrentStatus string  `json:"current_status"`
            PaymentGateway string  `json:"payment_gateway"`
            OrderID string  `json:"order_id"`
         
    }
    
    // MultiTenderPaymentMethod used by Cart
    type MultiTenderPaymentMethod struct {

        
            Name string  `json:"name"`
            Meta MultiTenderPaymentMeta  `json:"meta"`
            Mode string  `json:"mode"`
            Amount float64  `json:"amount"`
         
    }
    
    // CartItemMeta used by Cart
    type CartItemMeta struct {

        
            PrimaryItem bool  `json:"primary_item"`
            GroupID string  `json:"group_id"`
         
    }
    
    // OpenApiOrderItem used by Cart
    type OpenApiOrderItem struct {

        
            PaymentMethods []MultiTenderPaymentMethod  `json:"payment_methods"`
            AmountPaid float64  `json:"amount_paid"`
            EmployeeDiscount float64  `json:"employee_discount"`
            Files []OpenApiFiles  `json:"files"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            CodCharges float64  `json:"cod_charges"`
            PriceMarked float64  `json:"price_marked"`
            Meta CartItemMeta  `json:"meta"`
            Size string  `json:"size"`
            Quantity float64  `json:"quantity"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            Discount float64  `json:"discount"`
            ProductID float64  `json:"product_id"`
            CashbackApplied float64  `json:"cashback_applied"`
            DeliveryCharges float64  `json:"delivery_charges"`
            PriceEffective float64  `json:"price_effective"`
         
    }
    
    // OpenApiPlatformCheckoutReq used by Cart
    type OpenApiPlatformCheckoutReq struct {

        
            Files []OpenApiFiles  `json:"files"`
            Coupon string  `json:"coupon"`
            CodCharges float64  `json:"cod_charges"`
            OrderID string  `json:"order_id"`
            BillingAddress ShippingAddress  `json:"billing_address"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            DeliveryCharges float64  `json:"delivery_charges"`
            CashbackApplied float64  `json:"cashback_applied"`
            Gstin string  `json:"gstin"`
            Comment string  `json:"comment"`
            PaymentMethods []MultiTenderPaymentMethod  `json:"payment_methods"`
            EmployeeDiscount map[string]interface{}  `json:"employee_discount"`
            CouponValue float64  `json:"coupon_value"`
            CurrencyCode string  `json:"currency_code"`
            ShippingAddress ShippingAddress  `json:"shipping_address"`
            CartValue float64  `json:"cart_value"`
            CartItems []OpenApiOrderItem  `json:"cart_items"`
            PaymentMode string  `json:"payment_mode"`
            CouponCode string  `json:"coupon_code"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
         
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

        
            IsArchive bool  `json:"is_archive"`
            Coupon map[string]interface{}  `json:"coupon"`
            PickUpCustomerDetails map[string]interface{}  `json:"pick_up_customer_details"`
            Shipments []map[string]interface{}  `json:"shipments"`
            AppID string  `json:"app_id"`
            DeliveryCharges map[string]interface{}  `json:"delivery_charges"`
            BulkCouponDiscount float64  `json:"bulk_coupon_discount"`
            Comment string  `json:"comment"`
            Articles []map[string]interface{}  `json:"articles"`
            ExpireAt string  `json:"expire_at"`
            Meta map[string]interface{}  `json:"meta"`
            CartValue float64  `json:"cart_value"`
            UserID string  `json:"user_id"`
            UID float64  `json:"uid"`
            FcIndexMap []float64  `json:"fc_index_map"`
            PaymentMethods []map[string]interface{}  `json:"payment_methods"`
            ID string  `json:"_id"`
            MergeQty bool  `json:"merge_qty"`
            CodCharges map[string]interface{}  `json:"cod_charges"`
            Promotion map[string]interface{}  `json:"promotion"`
            IsDefault bool  `json:"is_default"`
            OrderID string  `json:"order_id"`
            Cashback map[string]interface{}  `json:"cashback"`
            LastModified string  `json:"last_modified"`
            CreatedOn string  `json:"created_on"`
            Gstin string  `json:"gstin"`
            CheckoutMode string  `json:"checkout_mode"`
            IsActive bool  `json:"is_active"`
            BuyNow bool  `json:"buy_now"`
            Discount float64  `json:"discount"`
            Payments map[string]interface{}  `json:"payments"`
            PaymentMode string  `json:"payment_mode"`
            FyndCredits map[string]interface{}  `json:"fynd_credits"`
         
    }
    
    // AbandonedCartResponse used by Cart
    type AbandonedCartResponse struct {

        
            Result map[string]interface{}  `json:"result"`
            Page Page  `json:"page"`
            Message string  `json:"message"`
            Items []AbandonedCart  `json:"items"`
            Success bool  `json:"success"`
         
    }
    
    // PaymentSelectionLock used by Cart
    type PaymentSelectionLock struct {

        
            Enabled bool  `json:"enabled"`
            PaymentIdentifier string  `json:"payment_identifier"`
            DefaultOptions string  `json:"default_options"`
         
    }
    
    // CartCurrency used by Cart
    type CartCurrency struct {

        
            Code string  `json:"code"`
            Symbol string  `json:"symbol"`
         
    }
    
    // CartDetailResponse used by Cart
    type CartDetailResponse struct {

        
            BreakupValues CartBreakup  `json:"breakup_values"`
            AppliedPromoDetails []AppliedPromotion  `json:"applied_promo_details"`
            IsValid bool  `json:"is_valid"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            BuyNow bool  `json:"buy_now"`
            CouponText string  `json:"coupon_text"`
            Currency CartCurrency  `json:"currency"`
            Message string  `json:"message"`
            Items []CartProductInfo  `json:"items"`
            ID string  `json:"id"`
            LastModified string  `json:"last_modified"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            Gstin string  `json:"gstin"`
            Comment string  `json:"comment"`
            CheckoutMode string  `json:"checkout_mode"`
         
    }
    
    // AddProductCart used by Cart
    type AddProductCart struct {

        
            StoreID float64  `json:"store_id"`
            SellerID float64  `json:"seller_id"`
            ParentItemIdentifiers []map[string]string  `json:"parent_item_identifiers"`
            ArticleID string  `json:"article_id"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Display string  `json:"display"`
            ItemSize string  `json:"item_size"`
            Pos bool  `json:"pos"`
            Quantity float64  `json:"quantity"`
            ArticleAssignment map[string]interface{}  `json:"article_assignment"`
            ItemID float64  `json:"item_id"`
            ProductGroupTags []string  `json:"product_group_tags"`
         
    }
    
    // AddCartRequest used by Cart
    type AddCartRequest struct {

        
            Items []AddProductCart  `json:"items"`
         
    }
    
    // AddCartDetailResponse used by Cart
    type AddCartDetailResponse struct {

        
            Cart CartDetailResponse  `json:"cart"`
            Message string  `json:"message"`
            Success bool  `json:"success"`
            Partial bool  `json:"partial"`
         
    }
    
    // UpdateProductCart used by Cart
    type UpdateProductCart struct {

        
            ArticleID string  `json:"article_id"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ItemSize string  `json:"item_size"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
            Quantity float64  `json:"quantity"`
            ItemIndex float64  `json:"item_index"`
            ItemID float64  `json:"item_id"`
         
    }
    
    // UpdateCartRequest used by Cart
    type UpdateCartRequest struct {

        
            Items []UpdateProductCart  `json:"items"`
            Operation string  `json:"operation"`
         
    }
    
    // UpdateCartDetailResponse used by Cart
    type UpdateCartDetailResponse struct {

        
            Cart CartDetailResponse  `json:"cart"`
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
    
