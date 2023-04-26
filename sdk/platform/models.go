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
            Category string  `json:"category"`
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

        
            Items []string  `json:"items"`
         
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
    
    // UpdateUserRequestSchema used by User
    type UpdateUserRequestSchema struct {

        
            FirstName string  `json:"first_name"`
            LastName string  `json:"last_name"`
            Gender string  `json:"gender"`
            ExternalID string  `json:"external_id"`
            Meta map[string]interface{}  `json:"meta"`
         
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

        
            ExcludedFields []string  `json:"excluded_fields"`
            AppID string  `json:"app_id"`
            DisplayFields []string  `json:"display_fields"`
            Aggregators []map[string]interface{}  `json:"aggregators"`
            Created bool  `json:"created"`
            Success bool  `json:"success"`
         
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
            MerchantSalt string  `json:"merchant_salt"`
            Key string  `json:"key"`
            Secret string  `json:"secret"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // PaymentGatewayConfigRequest used by Payment
    type PaymentGatewayConfigRequest struct {

        
            AggregatorName PaymentGatewayConfig  `json:"aggregator_name"`
            AppID string  `json:"app_id"`
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

        
            Success bool  `json:"success"`
            Error ErrorCodeAndDescription  `json:"error"`
         
    }
    
    // PaymentModeLogo used by Payment
    type PaymentModeLogo struct {

        
            Large string  `json:"large"`
            Small string  `json:"small"`
         
    }
    
    // IntentApp used by Payment
    type IntentApp struct {

        
            Logos PaymentModeLogo  `json:"logos"`
            DisplayName string  `json:"display_name"`
            PackageName string  `json:"package_name"`
            Code string  `json:"code"`
         
    }
    
    // IntentAppErrorList used by Payment
    type IntentAppErrorList struct {

        
            Code string  `json:"code"`
            PackageName string  `json:"package_name"`
         
    }
    
    // PaymentModeList used by Payment
    type PaymentModeList struct {

        
            CardIssuer string  `json:"card_issuer"`
            Nickname string  `json:"nickname"`
            CodLimit float64  `json:"cod_limit"`
            CompliantWithTokenisationGuidelines bool  `json:"compliant_with_tokenisation_guidelines"`
            RetryCount float64  `json:"retry_count"`
            CardID string  `json:"card_id"`
            IntentApp []IntentApp  `json:"intent_app"`
            LogoURL PaymentModeLogo  `json:"logo_url"`
            FyndVpa string  `json:"fynd_vpa"`
            CardBrandImage string  `json:"card_brand_image"`
            CardNumber string  `json:"card_number"`
            MerchantCode string  `json:"merchant_code"`
            CardName string  `json:"card_name"`
            AggregatorName string  `json:"aggregator_name"`
            Timeout float64  `json:"timeout"`
            RemainingLimit float64  `json:"remaining_limit"`
            ExpMonth float64  `json:"exp_month"`
            IntentAppErrorDictList []IntentAppErrorList  `json:"intent_app_error_dict_list"`
            CardType string  `json:"card_type"`
            DisplayName string  `json:"display_name"`
            IntentAppErrorList []string  `json:"intent_app_error_list"`
            ExpYear float64  `json:"exp_year"`
            IntentFlow bool  `json:"intent_flow"`
            Code string  `json:"code"`
            CodLimitPerOrder float64  `json:"cod_limit_per_order"`
            CardBrand string  `json:"card_brand"`
            Expired bool  `json:"expired"`
            DisplayPriority float64  `json:"display_priority"`
            CardReference string  `json:"card_reference"`
            CardFingerprint string  `json:"card_fingerprint"`
            CardIsin string  `json:"card_isin"`
            Name string  `json:"name"`
            CardToken string  `json:"card_token"`
         
    }
    
    // RootPaymentMode used by Payment
    type RootPaymentMode struct {

        
            SaveCard bool  `json:"save_card"`
            DisplayName string  `json:"display_name"`
            List []PaymentModeList  `json:"list"`
            AggregatorName string  `json:"aggregator_name"`
            AddCardEnabled bool  `json:"add_card_enabled"`
            Name string  `json:"name"`
            DisplayPriority float64  `json:"display_priority"`
            IsPayByCardPl bool  `json:"is_pay_by_card_pl"`
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

        
            TransferType string  `json:"transfer_type"`
            MoreAttributes map[string]interface{}  `json:"more_attributes"`
            IsDefault bool  `json:"is_default"`
            PayoutsAggregators []map[string]interface{}  `json:"payouts_aggregators"`
            UniqueTransferNo map[string]interface{}  `json:"unique_transfer_no"`
            Customers map[string]interface{}  `json:"customers"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // PayoutBankDetails used by Payment
    type PayoutBankDetails struct {

        
            AccountNo string  `json:"account_no"`
            AccountType string  `json:"account_type"`
            IfscCode string  `json:"ifsc_code"`
            State string  `json:"state"`
            BankName string  `json:"bank_name"`
            City string  `json:"city"`
            Pincode float64  `json:"pincode"`
            Country string  `json:"country"`
            BranchName string  `json:"branch_name"`
            AccountHolder string  `json:"account_holder"`
         
    }
    
    // PayoutRequest used by Payment
    type PayoutRequest struct {

        
            Users map[string]interface{}  `json:"users"`
            TransferType string  `json:"transfer_type"`
            UniqueExternalID string  `json:"unique_external_id"`
            Aggregator string  `json:"aggregator"`
            BankDetails PayoutBankDetails  `json:"bank_details"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // PayoutResponse used by Payment
    type PayoutResponse struct {

        
            Users map[string]interface{}  `json:"users"`
            TransferType string  `json:"transfer_type"`
            Aggregator string  `json:"aggregator"`
            Created bool  `json:"created"`
            Payouts map[string]interface{}  `json:"payouts"`
            UniqueTransferNo string  `json:"unique_transfer_no"`
            Success bool  `json:"success"`
            BankDetails map[string]interface{}  `json:"bank_details"`
            PaymentStatus string  `json:"payment_status"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // UpdatePayoutResponse used by Payment
    type UpdatePayoutResponse struct {

        
            Success bool  `json:"success"`
            IsDefault bool  `json:"is_default"`
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
            IsVerifiedFlag bool  `json:"is_verified_flag"`
            Data map[string]interface{}  `json:"data"`
         
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

        
            Success bool  `json:"success"`
            BranchName string  `json:"branch_name"`
            BankName string  `json:"bank_name"`
         
    }
    
    // OrderBeneficiaryDetails used by Payment
    type OrderBeneficiaryDetails struct {

        
            Address string  `json:"address"`
            IfscCode string  `json:"ifsc_code"`
            Email string  `json:"email"`
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
            AccountHolder string  `json:"account_holder"`
            IsActive bool  `json:"is_active"`
            AccountNo string  `json:"account_no"`
            Comment string  `json:"comment"`
            Title string  `json:"title"`
            ID float64  `json:"id"`
            DisplayName string  `json:"display_name"`
            DelightsUserName string  `json:"delights_user_name"`
            TransferMode string  `json:"transfer_mode"`
            BeneficiaryID string  `json:"beneficiary_id"`
            Mobile string  `json:"mobile"`
            BankName string  `json:"bank_name"`
            BranchName string  `json:"branch_name"`
            Subtitle string  `json:"subtitle"`
         
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
            PaymentID string  `json:"payment_id"`
            OrderID string  `json:"order_id"`
            PaymentGateway string  `json:"payment_gateway"`
         
    }
    
    // MultiTenderPaymentMethod used by Payment
    type MultiTenderPaymentMethod struct {

        
            Amount float64  `json:"amount"`
            Name string  `json:"name"`
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

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
            OrderID string  `json:"order_id"`
         
    }
    
    // CODdata used by Payment
    type CODdata struct {

        
            RemainingLimit float64  `json:"remaining_limit"`
            Limit float64  `json:"limit"`
            UserID string  `json:"user_id"`
            Usages float64  `json:"usages"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // GetUserCODLimitResponse used by Payment
    type GetUserCODLimitResponse struct {

        
            Success bool  `json:"success"`
            UserCodData CODdata  `json:"user_cod_data"`
         
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

        
            AggregatorID float64  `json:"aggregator_id"`
            Models []string  `json:"models"`
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

        
            TerminalSerialNo string  `json:"terminal_serial_no"`
            DeviceTag string  `json:"device_tag"`
            EdcModel string  `json:"edc_model"`
            AggregatorID float64  `json:"aggregator_id"`
            EdcDeviceSerialNo string  `json:"edc_device_serial_no"`
            StoreID float64  `json:"store_id"`
         
    }
    
    // EdcDevice used by Payment
    type EdcDevice struct {

        
            TerminalSerialNo string  `json:"terminal_serial_no"`
            AggregatorName string  `json:"aggregator_name"`
            TerminalUniqueIdentifier string  `json:"terminal_unique_identifier"`
            DeviceTag string  `json:"device_tag"`
            EdcModel string  `json:"edc_model"`
            AggregatorID float64  `json:"aggregator_id"`
            EdcDeviceSerialNo string  `json:"edc_device_serial_no"`
            ApplicationID string  `json:"application_id"`
            MerchantStorePosCode string  `json:"merchant_store_pos_code"`
            StoreID float64  `json:"store_id"`
            IsActive bool  `json:"is_active"`
         
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
            EdcModel string  `json:"edc_model"`
            AggregatorID float64  `json:"aggregator_id"`
            MerchantStorePosCode string  `json:"merchant_store_pos_code"`
            EdcDeviceSerialNo string  `json:"edc_device_serial_no"`
            StoreID float64  `json:"store_id"`
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

        
            Success bool  `json:"success"`
            Items []EdcDevice  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // PaymentInitializationRequest used by Payment
    type PaymentInitializationRequest struct {

        
            MerchantOrderID string  `json:"merchant_order_id"`
            DeviceID string  `json:"device_id"`
            Currency string  `json:"currency"`
            RazorpayPaymentID string  `json:"razorpay_payment_id"`
            Timeout float64  `json:"timeout"`
            CustomerID string  `json:"customer_id"`
            Email string  `json:"email"`
            Contact string  `json:"contact"`
            Method string  `json:"method"`
            Vpa string  `json:"vpa"`
            OrderID string  `json:"order_id"`
            Aggregator string  `json:"aggregator"`
            Amount float64  `json:"amount"`
         
    }
    
    // PaymentInitializationResponse used by Payment
    type PaymentInitializationResponse struct {

        
            MerchantOrderID string  `json:"merchant_order_id"`
            DeviceID string  `json:"device_id"`
            BqrImage string  `json:"bqr_image"`
            Currency string  `json:"currency"`
            RazorpayPaymentID string  `json:"razorpay_payment_id"`
            AggregatorOrderID string  `json:"aggregator_order_id"`
            CustomerID string  `json:"customer_id"`
            Timeout float64  `json:"timeout"`
            Method string  `json:"method"`
            Vpa string  `json:"vpa"`
            Status string  `json:"status"`
            PollingURL string  `json:"polling_url"`
            UpiPollURL string  `json:"upi_poll_url"`
            VirtualID string  `json:"virtual_id"`
            Aggregator string  `json:"aggregator"`
            Success bool  `json:"success"`
            Amount float64  `json:"amount"`
         
    }
    
    // PaymentStatusUpdateRequest used by Payment
    type PaymentStatusUpdateRequest struct {

        
            MerchantOrderID string  `json:"merchant_order_id"`
            DeviceID string  `json:"device_id"`
            Currency string  `json:"currency"`
            CustomerID string  `json:"customer_id"`
            Email string  `json:"email"`
            Contact string  `json:"contact"`
            Method string  `json:"method"`
            Status string  `json:"status"`
            Vpa string  `json:"vpa"`
            OrderID string  `json:"order_id"`
            Aggregator string  `json:"aggregator"`
            Amount float64  `json:"amount"`
         
    }
    
    // PaymentStatusUpdateResponse used by Payment
    type PaymentStatusUpdateResponse struct {

        
            Status string  `json:"status"`
            RedirectURL string  `json:"redirect_url"`
            AggregatorName string  `json:"aggregator_name"`
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

        
            Success bool  `json:"success"`
            Data LinkStatus  `json:"data"`
         
    }
    
    // Code used by Payment
    type Code struct {

        
            Name string  `json:"name"`
            MerchantCode string  `json:"merchant_code"`
            Code string  `json:"code"`
         
    }
    
    // PaymentCode used by Payment
    type PaymentCode struct {

        
            Codes Code  `json:"codes"`
            Name string  `json:"name"`
            Types string  `json:"types"`
            Networks string  `json:"networks"`
         
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
    

    
    // PaymentModeInfo used by Order
    type PaymentModeInfo struct {

        
            Logo string  `json:"logo"`
            Type string  `json:"type"`
         
    }
    
    // ShipmentStatus used by Order
    type ShipmentStatus struct {

        
            OpsStatus string  `json:"ops_status"`
            Title string  `json:"title"`
            ActualStatus string  `json:"actual_status"`
            Status string  `json:"status"`
            HexCode string  `json:"hex_code"`
         
    }
    
    // UserDataInfo used by Order
    type UserDataInfo struct {

        
            AvisUserID string  `json:"avis_user_id"`
            UID float64  `json:"uid"`
            Gender string  `json:"gender"`
            Mobile string  `json:"mobile"`
            FirstName string  `json:"first_name"`
            IsAnonymousUser bool  `json:"is_anonymous_user"`
            LastName string  `json:"last_name"`
            Email string  `json:"email"`
            Name string  `json:"name"`
         
    }
    
    // Prices used by Order
    type Prices struct {

        
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            CodCharges float64  `json:"cod_charges"`
            Discount float64  `json:"discount"`
            GiftPrice float64  `json:"gift_price"`
            AmountPaidRoundoff float64  `json:"amount_paid_roundoff"`
            PriceEffective float64  `json:"price_effective"`
            CouponValue float64  `json:"coupon_value"`
            PriceMarked float64  `json:"price_marked"`
            RefundCredit float64  `json:"refund_credit"`
            DeliveryCharge float64  `json:"delivery_charge"`
            RefundAmount float64  `json:"refund_amount"`
            FyndCredits float64  `json:"fynd_credits"`
            Cashback float64  `json:"cashback"`
            ValueOfGood float64  `json:"value_of_good"`
            PromotionEffectiveDiscount float64  `json:"promotion_effective_discount"`
            CashbackApplied float64  `json:"cashback_applied"`
            AmountPaid float64  `json:"amount_paid"`
            TransferPrice float64  `json:"transfer_price"`
         
    }
    
    // ShipmentItemFulFillingStore used by Order
    type ShipmentItemFulFillingStore struct {

        
            ID string  `json:"id"`
            Code string  `json:"code"`
         
    }
    
    // PlatformItem used by Order
    type PlatformItem struct {

        
            Size string  `json:"size"`
            Image []string  `json:"image"`
            L3Category float64  `json:"l3_category"`
            DepartmentID float64  `json:"department_id"`
            CanReturn bool  `json:"can_return"`
            L3CategoryName string  `json:"l3_category_name"`
            Color string  `json:"color"`
            CanCancel bool  `json:"can_cancel"`
            ID float64  `json:"id"`
            Name string  `json:"name"`
            Images []string  `json:"images"`
            Code string  `json:"code"`
            L1Category []string  `json:"l1_category"`
         
    }
    
    // GSTDetailsData used by Order
    type GSTDetailsData struct {

        
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            GstinCode string  `json:"gstin_code"`
            GstFee float64  `json:"gst_fee"`
            ValueOfGood float64  `json:"value_of_good"`
         
    }
    
    // BagUnit used by Order
    type BagUnit struct {

        
            OrderingChannel string  `json:"ordering_channel"`
            TotalShipmentBags float64  `json:"total_shipment_bags"`
            ItemQuantity float64  `json:"item_quantity"`
            Prices Prices  `json:"prices"`
            Item PlatformItem  `json:"item"`
            Status map[string]interface{}  `json:"status"`
            CanCancel bool  `json:"can_cancel"`
            BagID float64  `json:"bag_id"`
            Gst GSTDetailsData  `json:"gst"`
            CanReturn bool  `json:"can_return"`
            ShipmentID string  `json:"shipment_id"`
         
    }
    
    // ShipmentItem used by Order
    type ShipmentItem struct {

        
            Application map[string]interface{}  `json:"application"`
            PaymentModeInfo PaymentModeInfo  `json:"payment_mode_info"`
            ShipmentStatus ShipmentStatus  `json:"shipment_status"`
            Channel map[string]interface{}  `json:"channel"`
            CreatedAt string  `json:"created_at"`
            User UserDataInfo  `json:"user"`
            Prices Prices  `json:"prices"`
            FulfillingStore ShipmentItemFulFillingStore  `json:"fulfilling_store"`
            Sla map[string]interface{}  `json:"sla"`
            FulfillingCentre string  `json:"fulfilling_centre"`
            Bags []BagUnit  `json:"bags"`
            ShipmentCreatedAt string  `json:"shipment_created_at"`
            TotalShipmentsInOrder float64  `json:"total_shipments_in_order"`
            PaymentMethods map[string]interface{}  `json:"payment_methods"`
            ID string  `json:"id"`
            TotalBagsCount float64  `json:"total_bags_count"`
            ShipmentID string  `json:"shipment_id"`
         
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
            Page map[string]interface{}  `json:"page"`
            AppliedFilters map[string]interface{}  `json:"applied_filters"`
            Filters []FiltersInfo  `json:"filters"`
         
    }
    
    // Error used by Order
    type Error struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // UserDetailsData used by Order
    type UserDetailsData struct {

        
            Landmark string  `json:"landmark"`
            AddressType string  `json:"address_type"`
            Email string  `json:"email"`
            State string  `json:"state"`
            Pincode string  `json:"pincode"`
            City string  `json:"city"`
            Address1 string  `json:"address1"`
            Area string  `json:"area"`
            Phone string  `json:"phone"`
            Country string  `json:"country"`
            Name string  `json:"name"`
            Address string  `json:"address"`
         
    }
    
    // ShipmentPayments used by Order
    type ShipmentPayments struct {

        
            Logo string  `json:"logo"`
            Source string  `json:"source"`
            Mode string  `json:"mode"`
         
    }
    
    // OrderingStoreDetails used by Order
    type OrderingStoreDetails struct {

        
            StoreName string  `json:"store_name"`
            State string  `json:"state"`
            OrderingStoreID float64  `json:"ordering_store_id"`
            Pincode string  `json:"pincode"`
            Meta map[string]interface{}  `json:"meta"`
            City string  `json:"city"`
            Phone string  `json:"phone"`
            ContactPerson string  `json:"contact_person"`
            Country string  `json:"country"`
            Address string  `json:"address"`
            Code string  `json:"code"`
         
    }
    
    // FulfillingStore used by Order
    type FulfillingStore struct {

        
            StoreName string  `json:"store_name"`
            State string  `json:"state"`
            Pincode string  `json:"pincode"`
            Meta map[string]interface{}  `json:"meta"`
            FulfillmentChannel string  `json:"fulfillment_channel"`
            City string  `json:"city"`
            Phone string  `json:"phone"`
            ContactPerson string  `json:"contact_person"`
            Country string  `json:"country"`
            ID float64  `json:"id"`
            Address string  `json:"address"`
            Code string  `json:"code"`
         
    }
    
    // BagConfigs used by Order
    type BagConfigs struct {

        
            AllowForceReturn bool  `json:"allow_force_return"`
            IsReturnable bool  `json:"is_returnable"`
            CanBeCancelled bool  `json:"can_be_cancelled"`
            IsActive bool  `json:"is_active"`
            EnableTracking bool  `json:"enable_tracking"`
            IsCustomerReturnAllowed bool  `json:"is_customer_return_allowed"`
         
    }
    
    // OrderBrandName used by Order
    type OrderBrandName struct {

        
            Company string  `json:"company"`
            Logo string  `json:"logo"`
            BrandName string  `json:"brand_name"`
            CreatedOn string  `json:"created_on"`
            ID float64  `json:"id"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // OrderBagArticle used by Order
    type OrderBagArticle struct {

        
            ReturnConfig map[string]interface{}  `json:"return_config"`
            UID string  `json:"uid"`
            Identifiers map[string]interface{}  `json:"identifiers"`
         
    }
    
    // Identifier used by Order
    type Identifier struct {

        
            Ean string  `json:"ean"`
         
    }
    
    // FinancialBreakup used by Order
    type FinancialBreakup struct {

        
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            AmountPaidRoundoff float64  `json:"amount_paid_roundoff"`
            PriceEffective float64  `json:"price_effective"`
            DeliveryCharge float64  `json:"delivery_charge"`
            CashbackApplied float64  `json:"cashback_applied"`
            AmountPaid float64  `json:"amount_paid"`
            Identifiers Identifier  `json:"identifiers"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            Cashback float64  `json:"cashback"`
            GstFee float64  `json:"gst_fee"`
            ValueOfGood float64  `json:"value_of_good"`
            PromotionEffectiveDiscount float64  `json:"promotion_effective_discount"`
            Discount float64  `json:"discount"`
            AddedToFyndCash bool  `json:"added_to_fynd_cash"`
            GstTag string  `json:"gst_tag"`
            PriceMarked float64  `json:"price_marked"`
            CouponValue float64  `json:"coupon_value"`
            ItemName string  `json:"item_name"`
            TransferPrice float64  `json:"transfer_price"`
            Size string  `json:"size"`
            CodCharges float64  `json:"cod_charges"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            TotalUnits float64  `json:"total_units"`
            RefundCredit float64  `json:"refund_credit"`
            FyndCredits float64  `json:"fynd_credits"`
            HsnCode string  `json:"hsn_code"`
         
    }
    
    // PlatformDeliveryAddress used by Order
    type PlatformDeliveryAddress struct {

        
            Landmark string  `json:"landmark"`
            AddressCategory string  `json:"address_category"`
            Longitude float64  `json:"longitude"`
            Version string  `json:"version"`
            Address2 string  `json:"address2"`
            Email string  `json:"email"`
            CreatedAt string  `json:"created_at"`
            State string  `json:"state"`
            Pincode string  `json:"pincode"`
            UpdatedAt string  `json:"updated_at"`
            Latitude float64  `json:"latitude"`
            City string  `json:"city"`
            Address1 string  `json:"address1"`
            Area string  `json:"area"`
            Phone string  `json:"phone"`
            ContactPerson string  `json:"contact_person"`
            Country string  `json:"country"`
            AddressType string  `json:"address_type"`
         
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

        
            ArticleQuantity float64  `json:"article_quantity"`
            PromoID string  `json:"promo_id"`
            BuyRules []BuyRules  `json:"buy_rules"`
            Amount float64  `json:"amount"`
            PromotionType string  `json:"promotion_type"`
            DiscountRules []DiscountRules  `json:"discount_rules"`
            MrpPromotion bool  `json:"mrp_promotion"`
            PromotionName string  `json:"promotion_name"`
         
    }
    
    // BagGST used by Order
    type BagGST struct {

        
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            GstTag string  `json:"gst_tag"`
            GstinCode string  `json:"gstin_code"`
            GstFee float64  `json:"gst_fee"`
            ValueOfGood float64  `json:"value_of_good"`
            HsnCode string  `json:"hsn_code"`
            IsDefaultHsnCode bool  `json:"is_default_hsn_code"`
         
    }
    
    // B2BPODetails used by Order
    type B2BPODetails struct {

        
            ItemBasePrice float64  `json:"item_base_price"`
            PoTaxAmount float64  `json:"po_tax_amount"`
            PoLineAmount float64  `json:"po_line_amount"`
            TotalGstPercentage float64  `json:"total_gst_percentage"`
            PartialCanRet bool  `json:"partial_can_ret"`
            DockerNumber string  `json:"docker_number"`
         
    }
    
    // GiftCard used by Order
    type GiftCard struct {

        
            GiftPrice float64  `json:"gift_price"`
            DisplayText string  `json:"display_text"`
            IsGiftApplied bool  `json:"is_gift_applied"`
            GiftMessage string  `json:"gift_message"`
         
    }
    
    // BagMeta used by Order
    type BagMeta struct {

        
            B2bPoDetails B2BPODetails  `json:"b2b_po_details"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            GiftCard GiftCard  `json:"gift_card"`
            DocketNumber string  `json:"docket_number"`
            CustomMessage string  `json:"custom_message"`
            CustomJson map[string]interface{}  `json:"custom_json"`
            GroupID string  `json:"group_id"`
            PartialCanRet bool  `json:"partial_can_ret"`
         
    }
    
    // BagStateMapper used by Order
    type BagStateMapper struct {

        
            AppDisplayName string  `json:"app_display_name"`
            DisplayName string  `json:"display_name"`
            AppStateName string  `json:"app_state_name"`
            AppFacing bool  `json:"app_facing"`
            IsActive bool  `json:"is_active"`
            JourneyType string  `json:"journey_type"`
            NotifyCustomer bool  `json:"notify_customer"`
            BsID float64  `json:"bs_id"`
            Name string  `json:"name"`
            StateType string  `json:"state_type"`
         
    }
    
    // CurrentStatus used by Order
    type CurrentStatus struct {

        
            DeliveryPartnerID float64  `json:"delivery_partner_id"`
            CreatedAt string  `json:"created_at"`
            BagID float64  `json:"bag_id"`
            DeliveryAwbNumber string  `json:"delivery_awb_number"`
            CurrentStatusID float64  `json:"current_status_id"`
            UpdatedAt string  `json:"updated_at"`
            StoreID float64  `json:"store_id"`
            KafkaSync bool  `json:"kafka_sync"`
            BagStateMapper BagStateMapper  `json:"bag_state_mapper"`
            StateID float64  `json:"state_id"`
            Status string  `json:"status"`
            StateType string  `json:"state_type"`
            ShipmentID string  `json:"shipment_id"`
         
    }
    
    // OrderBags used by Order
    type OrderBags struct {

        
            Prices Prices  `json:"prices"`
            Item PlatformItem  `json:"item"`
            Quantity float64  `json:"quantity"`
            CanCancel bool  `json:"can_cancel"`
            CanReturn bool  `json:"can_return"`
            BagConfigs BagConfigs  `json:"bag_configs"`
            Brand OrderBrandName  `json:"brand"`
            Article OrderBagArticle  `json:"article"`
            EntityType string  `json:"entity_type"`
            FinancialBreakup FinancialBreakup  `json:"financial_breakup"`
            DisplayName string  `json:"display_name"`
            LineNumber float64  `json:"line_number"`
            DeliveryAddress PlatformDeliveryAddress  `json:"delivery_address"`
            Identifier string  `json:"identifier"`
            GroupID string  `json:"group_id"`
            ParentPromoBags map[string]interface{}  `json:"parent_promo_bags"`
            AppliedPromos []AppliedPromos  `json:"applied_promos"`
            GstDetails BagGST  `json:"gst_details"`
            Meta BagMeta  `json:"meta"`
            SellerIdentifier string  `json:"seller_identifier"`
            BagID float64  `json:"bag_id"`
            CurrentStatus CurrentStatus  `json:"current_status"`
         
    }
    
    // DPDetailsData used by Order
    type DPDetailsData struct {

        
            GstTag string  `json:"gst_tag"`
            EwayBillID string  `json:"eway_bill_id"`
            Pincode string  `json:"pincode"`
            TrackURL string  `json:"track_url"`
            AwbNo string  `json:"awb_no"`
            Country string  `json:"country"`
            ID float64  `json:"id"`
            Name string  `json:"name"`
         
    }
    
    // BagStatusHistory used by Order
    type BagStatusHistory struct {

        
            AppDisplayName string  `json:"app_display_name"`
            DisplayName string  `json:"display_name"`
            Forward bool  `json:"forward"`
            DeliveryPartnerID float64  `json:"delivery_partner_id"`
            BagID float64  `json:"bag_id"`
            StateType string  `json:"state_type"`
            DeliveryAwbNumber string  `json:"delivery_awb_number"`
            CreatedAt string  `json:"created_at"`
            Reasons []map[string]interface{}  `json:"reasons"`
            UpdatedAt string  `json:"updated_at"`
            KafkaSync bool  `json:"kafka_sync"`
            StoreID float64  `json:"store_id"`
            BagStateMapper BagStateMapper  `json:"bag_state_mapper"`
            StateID float64  `json:"state_id"`
            Status string  `json:"status"`
            BshID float64  `json:"bsh_id"`
            ShipmentID string  `json:"shipment_id"`
         
    }
    
    // InvoiceInfo used by Order
    type InvoiceInfo struct {

        
            CreditNoteID string  `json:"credit_note_id"`
            StoreInvoiceID string  `json:"store_invoice_id"`
            LabelURL string  `json:"label_url"`
            UpdatedDate string  `json:"updated_date"`
            InvoiceURL string  `json:"invoice_url"`
         
    }
    
    // TrackingList used by Order
    type TrackingList struct {

        
            IsPassed bool  `json:"is_passed"`
            Time string  `json:"time"`
            IsCurrent bool  `json:"is_current"`
            Text string  `json:"text"`
            Status string  `json:"status"`
         
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

        
            CompanyGst string  `json:"company_gst"`
            CompanyCin string  `json:"company_cin"`
            CompanyName string  `json:"company_name"`
            CompanyID float64  `json:"company_id"`
            CompanyContact ContactDetails  `json:"company_contact"`
            Address map[string]interface{}  `json:"address"`
         
    }
    
    // OrderDetailsData used by Order
    type OrderDetailsData struct {

        
            CodCharges string  `json:"cod_charges"`
            TaxDetails map[string]interface{}  `json:"tax_details"`
            OrderingChannel string  `json:"ordering_channel"`
            OrderValue string  `json:"order_value"`
            OrderingChannelLogo map[string]interface{}  `json:"ordering_channel_logo"`
            AffiliateID string  `json:"affiliate_id"`
            FyndOrderID string  `json:"fynd_order_id"`
            Source string  `json:"source"`
            OrderDate string  `json:"order_date"`
         
    }
    
    // Dimensions used by Order
    type Dimensions struct {

        
            Height float64  `json:"height"`
            Length float64  `json:"length"`
            Unit string  `json:"unit"`
            Width float64  `json:"width"`
            IsDefault bool  `json:"is_default"`
         
    }
    
    // Meta used by Order
    type Meta struct {

        
            LockData map[string]interface{}  `json:"lock_data"`
            Dimension Dimensions  `json:"dimension"`
         
    }
    
    // ShipmentStatusData used by Order
    type ShipmentStatusData struct {

        
            CreatedAt string  `json:"created_at"`
            BagList []string  `json:"bag_list"`
            Status string  `json:"status"`
            ID float64  `json:"id"`
            ShipmentID string  `json:"shipment_id"`
         
    }
    
    // PDFLinks used by Order
    type PDFLinks struct {

        
            InvoiceType string  `json:"invoice_type"`
            B2b string  `json:"b2b"`
            PoInvoice string  `json:"po_invoice"`
            DeliveryChallanA4 string  `json:"delivery_challan_a4"`
            Label string  `json:"label"`
            LabelA4 string  `json:"label_a4"`
            LabelA6 string  `json:"label_a6"`
            LabelType string  `json:"label_type"`
            CreditNoteURL string  `json:"credit_note_url"`
            InvoicePos string  `json:"invoice_pos"`
            LabelPos string  `json:"label_pos"`
            InvoiceA6 string  `json:"invoice_a6"`
            InvoiceA4 string  `json:"invoice_a4"`
            Invoice string  `json:"invoice"`
         
    }
    
    // AffiliateMeta used by Order
    type AffiliateMeta struct {

        
            SizeLevelTotalQty float64  `json:"size_level_total_qty"`
            OrderItemID string  `json:"order_item_id"`
            BoxType string  `json:"box_type"`
            EmployeeDiscount float64  `json:"employee_discount"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            CouponCode string  `json:"coupon_code"`
            Quantity float64  `json:"quantity"`
            DueDate string  `json:"due_date"`
            ChannelOrderID string  `json:"channel_order_id"`
            IsPriority bool  `json:"is_priority"`
            ChannelShipmentID string  `json:"channel_shipment_id"`
         
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
    
    // LockData used by Order
    type LockData struct {

        
            LockMessage string  `json:"lock_message"`
            Mto bool  `json:"mto"`
            Locked bool  `json:"locked"`
         
    }
    
    // EinvoiceInfo used by Order
    type EinvoiceInfo struct {

        
            CreditNote map[string]interface{}  `json:"credit_note"`
            Invoice map[string]interface{}  `json:"invoice"`
         
    }
    
    // ShipmentTimeStamp used by Order
    type ShipmentTimeStamp struct {

        
            TMax string  `json:"t_max"`
            TMin string  `json:"t_min"`
         
    }
    
    // BuyerDetails used by Order
    type BuyerDetails struct {

        
            Address string  `json:"address"`
            Gstin string  `json:"gstin"`
            AjioSiteID string  `json:"ajio_site_id"`
            State string  `json:"state"`
            Pincode float64  `json:"pincode"`
            City string  `json:"city"`
            Name string  `json:"name"`
         
    }
    
    // ShipmentMeta used by Order
    type ShipmentMeta struct {

        
            ReturnAffiliateShipmentID string  `json:"return_affiliate_shipment_id"`
            ShipmentVolumetricWeight float64  `json:"shipment_volumetric_weight"`
            DebugInfo DebugInfo  `json:"debug_info"`
            StoreInvoiceUpdatedDate string  `json:"store_invoice_updated_date"`
            External map[string]interface{}  `json:"external"`
            ReturnAffiliateOrderID string  `json:"return_affiliate_order_id"`
            Formatted Formatted  `json:"formatted"`
            PoNumber string  `json:"po_number"`
            AwbNumber string  `json:"awb_number"`
            EwaybillInfo map[string]interface{}  `json:"ewaybill_info"`
            DpSortKey string  `json:"dp_sort_key"`
            AutoTriggerDpAssignmentAcf bool  `json:"auto_trigger_dp_assignment_acf"`
            B2cBuyerDetails map[string]interface{}  `json:"b2c_buyer_details"`
            DpOptions map[string]interface{}  `json:"dp_options"`
            LockData LockData  `json:"lock_data"`
            AssignDpFromSb bool  `json:"assign_dp_from_sb"`
            PackagingName string  `json:"packaging_name"`
            OrderType string  `json:"order_type"`
            BagWeight map[string]interface{}  `json:"bag_weight"`
            Weight float64  `json:"weight"`
            ReturnAwbNumber string  `json:"return_awb_number"`
            ShipmentWeight float64  `json:"shipment_weight"`
            ReturnStoreNode float64  `json:"return_store_node"`
            EinvoiceInfo EinvoiceInfo  `json:"einvoice_info"`
            DpName string  `json:"dp_name"`
            ForwardAffiliateOrderID string  `json:"forward_affiliate_order_id"`
            FulfilmentPriorityText string  `json:"fulfilment_priority_text"`
            ReturnDetails map[string]interface{}  `json:"return_details"`
            Timestamp ShipmentTimeStamp  `json:"timestamp"`
            B2bBuyerDetails BuyerDetails  `json:"b2b_buyer_details"`
            ForwardAffiliateShipmentID string  `json:"forward_affiliate_shipment_id"`
            MarketplaceStoreID string  `json:"marketplace_store_id"`
            SameStoreAvailable bool  `json:"same_store_available"`
            BoxType string  `json:"box_type"`
            DpID string  `json:"dp_id"`
            DueDate string  `json:"due_date"`
         
    }
    
    // AffiliateDetails used by Order
    type AffiliateDetails struct {

        
            PdfLinks PDFLinks  `json:"pdf_links"`
            AffiliateMeta AffiliateMeta  `json:"affiliate_meta"`
            AffiliateID string  `json:"affiliate_id"`
            ShipmentMeta ShipmentMeta  `json:"shipment_meta"`
            CompanyAffiliateTag string  `json:"company_affiliate_tag"`
            AffiliateStoreID string  `json:"affiliate_store_id"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            AdID string  `json:"ad_id"`
         
    }
    
    // PlatformShipment used by Order
    type PlatformShipment struct {

        
            UserAgent string  `json:"user_agent"`
            BillingDetails UserDetailsData  `json:"billing_details"`
            User UserDataInfo  `json:"user"`
            Prices Prices  `json:"prices"`
            Payments ShipmentPayments  `json:"payments"`
            DeliveryDetails UserDetailsData  `json:"delivery_details"`
            Coupon map[string]interface{}  `json:"coupon"`
            PlatformLogo string  `json:"platform_logo"`
            DeliverySlot map[string]interface{}  `json:"delivery_slot"`
            PackagingType string  `json:"packaging_type"`
            OperationalStatus string  `json:"operational_status"`
            EnableDpTracking bool  `json:"enable_dp_tracking"`
            ShipmentStatus string  `json:"shipment_status"`
            OrderingStore OrderingStoreDetails  `json:"ordering_store"`
            LockStatus bool  `json:"lock_status"`
            JourneyType string  `json:"journey_type"`
            ForwardShipmentID string  `json:"forward_shipment_id"`
            PaymentMethods map[string]interface{}  `json:"payment_methods"`
            FulfillingStore FulfillingStore  `json:"fulfilling_store"`
            Bags []OrderBags  `json:"bags"`
            DpDetails DPDetailsData  `json:"dp_details"`
            BagStatusHistory []BagStatusHistory  `json:"bag_status_history"`
            ShipmentImages []string  `json:"shipment_images"`
            FulfilmentPriority float64  `json:"fulfilment_priority"`
            PickedDate string  `json:"picked_date"`
            PriorityText string  `json:"priority_text"`
            Vertical string  `json:"vertical"`
            PaymentMode string  `json:"payment_mode"`
            TotalBags float64  `json:"total_bags"`
            Invoice InvoiceInfo  `json:"invoice"`
            GstDetails GSTDetailsData  `json:"gst_details"`
            TrackingList []TrackingList  `json:"tracking_list"`
            CustomMeta []map[string]interface{}  `json:"custom_meta"`
            CompanyDetails CompanyDetails  `json:"company_details"`
            InvoiceID string  `json:"invoice_id"`
            TotalItems float64  `json:"total_items"`
            Order OrderDetailsData  `json:"order"`
            Meta Meta  `json:"meta"`
            ShipmentQuantity float64  `json:"shipment_quantity"`
            Status ShipmentStatusData  `json:"status"`
            AffiliateDetails AffiliateDetails  `json:"affiliate_details"`
            ShipmentID string  `json:"shipment_id"`
         
    }
    
    // ShipmentInfoResponse used by Order
    type ShipmentInfoResponse struct {

        
            Success bool  `json:"success"`
            Shipments []PlatformShipment  `json:"shipments"`
            Message string  `json:"message"`
         
    }
    
    // TaxDetails used by Order
    type TaxDetails struct {

        
            Gstin string  `json:"gstin"`
            PanNo string  `json:"pan_no"`
         
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

        
            StaffID float64  `json:"staff_id"`
            FirstName string  `json:"first_name"`
            LastName string  `json:"last_name"`
            User string  `json:"user"`
            EmployeeCode string  `json:"employee_code"`
         
    }
    
    // TransactionData used by Order
    type TransactionData struct {

        
            PaymentID string  `json:"payment_id"`
            UniqueReferenceNumber string  `json:"unique_reference_number"`
            TerminalID string  `json:"terminal_id"`
            Status string  `json:"status"`
            Currency string  `json:"currency"`
            Entity string  `json:"entity"`
            AmountPaid string  `json:"amount_paid"`
            TransactionID string  `json:"transaction_id"`
         
    }
    
    // OrderMeta used by Order
    type OrderMeta struct {

        
            Comment string  `json:"comment"`
            OrderPlatform string  `json:"order_platform"`
            PlatformUserDetails PlatformUserDetails  `json:"platform_user_details"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            CompanyLogo string  `json:"company_logo"`
            OrderingStore float64  `json:"ordering_store"`
            EmployeeID float64  `json:"employee_id"`
            OrderTags []map[string]interface{}  `json:"order_tags"`
            OrderType string  `json:"order_type"`
            Files []map[string]interface{}  `json:"files"`
            BillingStaffDetails BillingStaffDetails  `json:"billing_staff_details"`
            CartID float64  `json:"cart_id"`
            PaymentType string  `json:"payment_type"`
            TransactionData TransactionData  `json:"transaction_data"`
            CurrencySymbol string  `json:"currency_symbol"`
            MongoCartID float64  `json:"mongo_cart_id"`
            Staff map[string]interface{}  `json:"staff"`
            CustomerNote string  `json:"customer_note"`
            OrderChildEntities []string  `json:"order_child_entities"`
         
    }
    
    // OrderDict used by Order
    type OrderDict struct {

        
            TaxDetails TaxDetails  `json:"tax_details"`
            Prices Prices  `json:"prices"`
            Meta OrderMeta  `json:"meta"`
            FyndOrderID string  `json:"fynd_order_id"`
            PaymentMethods map[string]interface{}  `json:"payment_methods"`
            OrderDate string  `json:"order_date"`
         
    }
    
    // ShipmentDetailsResponse used by Order
    type ShipmentDetailsResponse struct {

        
            Success bool  `json:"success"`
            Shipments []PlatformShipment  `json:"shipments"`
            Order OrderDict  `json:"order"`
         
    }
    
    // SubLane used by Order
    type SubLane struct {

        
            Index float64  `json:"index"`
            Actions []map[string]interface{}  `json:"actions"`
            TotalItems float64  `json:"total_items"`
            Value string  `json:"value"`
            Text string  `json:"text"`
         
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

        
            Name string  `json:"name"`
            Display string  `json:"display"`
            Value string  `json:"value"`
         
    }
    
    // PlatformOrderItems used by Order
    type PlatformOrderItems struct {

        
            Shipments []PlatformShipment  `json:"shipments"`
            OrderValue float64  `json:"order_value"`
            Channel PlatformChannel  `json:"channel"`
            OrderID string  `json:"order_id"`
            OrderCreatedTime string  `json:"order_created_time"`
            UserInfo UserDataInfo  `json:"user_info"`
            Meta map[string]interface{}  `json:"meta"`
            BreakupValues []PlatformBreakupValues  `json:"breakup_values"`
            PaymentMode string  `json:"payment_mode"`
            TotalOrderValue float64  `json:"total_order_value"`
         
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

        
            Items []PlatformOrderItems  `json:"items"`
            TotalCount float64  `json:"total_count"`
            Lane string  `json:"lane"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
            Page Page  `json:"page"`
         
    }
    
    // Options used by Order
    type Options struct {

        
            Text string  `json:"text"`
            Value float64  `json:"value"`
         
    }
    
    // MetricsCount used by Order
    type MetricsCount struct {

        
            Key string  `json:"key"`
            Text string  `json:"text"`
            Value float64  `json:"value"`
            Options []Options  `json:"options"`
         
    }
    
    // MetricCountResponse used by Order
    type MetricCountResponse struct {

        
            Items []MetricsCount  `json:"items"`
         
    }
    
    // PlatformTrack used by Order
    type PlatformTrack struct {

        
            Awb string  `json:"awb"`
            LastLocationRecievedAt string  `json:"last_location_recieved_at"`
            Reason string  `json:"reason"`
            ShipmentType string  `json:"shipment_type"`
            UpdatedTime string  `json:"updated_time"`
            UpdatedAt string  `json:"updated_at"`
            Meta map[string]interface{}  `json:"meta"`
            RawStatus string  `json:"raw_status"`
            Status string  `json:"status"`
            AccountName string  `json:"account_name"`
         
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

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // OmsReports used by Order
    type OmsReports struct {

        
            ReportType string  `json:"report_type"`
            DisplayName string  `json:"display_name"`
            S3Key string  `json:"s3_key"`
            ReportCreatedAt string  `json:"report_created_at"`
            ReportName string  `json:"report_name"`
            RequestDetails map[string]interface{}  `json:"request_details"`
            ReportRequestedAt string  `json:"report_requested_at"`
            Status string  `json:"status"`
            ReportID string  `json:"report_id"`
         
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

        
            Message string  `json:"message"`
            Type string  `json:"type"`
            Value string  `json:"value"`
         
    }
    
    // JioCodeUpsertResponse used by Order
    type JioCodeUpsertResponse struct {

        
            TraceID string  `json:"trace_id"`
            Error []NestedErrorSchemaDataSet  `json:"error"`
            Data []map[string]interface{}  `json:"data"`
            Identifier string  `json:"identifier"`
            Success bool  `json:"success"`
         
    }
    
    // BulkInvoicingResponse used by Order
    type BulkInvoicingResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // BulkInvoiceLabelResponse used by Order
    type BulkInvoiceLabelResponse struct {

        
            DoInvoiceLabelGenerated bool  `json:"do_invoice_label_generated"`
            Label map[string]interface{}  `json:"label"`
            StoreName string  `json:"store_name"`
            Data map[string]interface{}  `json:"data"`
            StoreID string  `json:"store_id"`
            InvoiceStatus string  `json:"invoice_status"`
            CompanyID string  `json:"company_id"`
            Invoice map[string]interface{}  `json:"invoice"`
            BatchID string  `json:"batch_id"`
            StoreCode string  `json:"store_code"`
         
    }
    
    // URL used by Order
    type URL struct {

        
            URL string  `json:"url"`
         
    }
    
    // FileUploadResponse used by Order
    type FileUploadResponse struct {

        
            Expiry float64  `json:"expiry"`
            URL string  `json:"url"`
         
    }
    
    // FileResponse used by Order
    type FileResponse struct {

        
            Size float64  `json:"size"`
            ContentType string  `json:"content_type"`
            Operation string  `json:"operation"`
            Cdn URL  `json:"cdn"`
            Namespace string  `json:"namespace"`
            Method string  `json:"method"`
            Upload FileUploadResponse  `json:"upload"`
            FileName string  `json:"file_name"`
            FilePath string  `json:"file_path"`
            Tags []string  `json:"tags"`
         
    }
    
    // bulkListingData used by Order
    type bulkListingData struct {

        
            UploadedOn string  `json:"uploaded_on"`
            SuccessfulShipments []map[string]interface{}  `json:"successful_shipments"`
            ExcelURL string  `json:"excel_url"`
            Failed float64  `json:"failed"`
            Successful float64  `json:"successful"`
            ID string  `json:"id"`
            Total float64  `json:"total"`
            FileName string  `json:"file_name"`
            CompanyID float64  `json:"company_id"`
            UserName string  `json:"user_name"`
            Processing float64  `json:"processing"`
            UserID string  `json:"user_id"`
            ProcessingShipments []string  `json:"processing_shipments"`
            StoreName string  `json:"store_name"`
            FailedShipments []map[string]interface{}  `json:"failed_shipments"`
            StoreID float64  `json:"store_id"`
            Status string  `json:"status"`
            BatchID string  `json:"batch_id"`
            StoreCode string  `json:"store_code"`
         
    }
    
    // BulkListingPage used by Order
    type BulkListingPage struct {

        
            Size float64  `json:"size"`
            HasPrevious bool  `json:"has_previous"`
            Current float64  `json:"current"`
            Total float64  `json:"total"`
            HasNext bool  `json:"has_next"`
            Type string  `json:"type"`
         
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

        
            SuccessfulShipmentIds []string  `json:"successful_shipment_ids"`
            SuccessfulShipmentsCount float64  `json:"successful_shipments_count"`
            FailedShipmentsCount float64  `json:"failed_shipments_count"`
            ProcessingShipmentsCount float64  `json:"processing_shipments_count"`
            CompanyID string  `json:"company_id"`
            BatchID string  `json:"batch_id"`
            TotalShipmentsCount float64  `json:"total_shipments_count"`
         
    }
    
    // BulkActionDetailsResponse used by Order
    type BulkActionDetailsResponse struct {

        
            UserID string  `json:"user_id"`
            UploadedOn string  `json:"uploaded_on"`
            Error []string  `json:"error"`
            Data []BulkActionDetailsDataField  `json:"data"`
            Success string  `json:"success"`
            FailedRecords []string  `json:"failed_records"`
            Status bool  `json:"status"`
            Message string  `json:"message"`
            UploadedBy string  `json:"uploaded_by"`
         
    }
    
    // Attributes used by Order
    type Attributes struct {

        
            Gender []string  `json:"gender"`
            PrimaryColor string  `json:"primary_color"`
            BrandName string  `json:"brand_name"`
            MarketerName string  `json:"marketer_name"`
            PrimaryMaterial string  `json:"primary_material"`
            PrimaryColorHex string  `json:"primary_color_hex"`
            Essential string  `json:"essential"`
            MarketerAddress string  `json:"marketer_address"`
            Name string  `json:"name"`
         
    }
    
    // Item used by Order
    type Item struct {

        
            L2Category []string  `json:"l2_category"`
            Image []string  `json:"image"`
            L3CategoryName string  `json:"l3_category_name"`
            CanCancel bool  `json:"can_cancel"`
            CanReturn bool  `json:"can_return"`
            L1Category []string  `json:"l1_category"`
            Gender string  `json:"gender"`
            L1CategoryID float64  `json:"l1_category_id"`
            SlugKey string  `json:"slug_key"`
            ItemID float64  `json:"item_id"`
            Brand string  `json:"brand"`
            L3Category float64  `json:"l3_category"`
            DepartmentID float64  `json:"department_id"`
            Attributes Attributes  `json:"attributes"`
            LastUpdatedAt string  `json:"last_updated_at"`
            BrandID float64  `json:"brand_id"`
            WebstoreProductURL string  `json:"webstore_product_url"`
            Name string  `json:"name"`
            Code string  `json:"code"`
            Size string  `json:"size"`
            BranchURL string  `json:"branch_url"`
            Meta map[string]interface{}  `json:"meta"`
            L2CategoryID float64  `json:"l2_category_id"`
            Color string  `json:"color"`
         
    }
    
    // ArticleDetails used by Order
    type ArticleDetails struct {

        
            Status map[string]interface{}  `json:"status"`
         
    }
    
    // Dates used by Order
    type Dates struct {

        
            OrderCreated string  `json:"order_created"`
            DeliveryDate interface{}  `json:"delivery_date"`
         
    }
    
    // StoreAddress used by Order
    type StoreAddress struct {

        
            Landmark string  `json:"landmark"`
            Version string  `json:"version"`
            Address2 string  `json:"address2"`
            CreatedAt string  `json:"created_at"`
            CountryCode string  `json:"country_code"`
            Longitude float64  `json:"longitude"`
            Pincode float64  `json:"pincode"`
            UpdatedAt string  `json:"updated_at"`
            Address1 string  `json:"address1"`
            Area string  `json:"area"`
            Phone string  `json:"phone"`
            ContactPerson string  `json:"contact_person"`
            Email string  `json:"email"`
            Latitude float64  `json:"latitude"`
            City string  `json:"city"`
            AddressType string  `json:"address_type"`
            AddressCategory string  `json:"address_category"`
            State string  `json:"state"`
            Country string  `json:"country"`
         
    }
    
    // Document used by Order
    type Document struct {

        
            Value string  `json:"value"`
            Verified bool  `json:"verified"`
            LegalName string  `json:"legal_name"`
            URL string  `json:"url"`
            DsType string  `json:"ds_type"`
         
    }
    
    // StoreDocuments used by Order
    type StoreDocuments struct {

        
            Gst Document  `json:"gst"`
         
    }
    
    // EInvoicePortalDetails used by Order
    type EInvoicePortalDetails struct {

        
            Password string  `json:"password"`
            User string  `json:"user"`
            Username string  `json:"username"`
         
    }
    
    // StoreEinvoice used by Order
    type StoreEinvoice struct {

        
            Password string  `json:"password"`
            Enabled bool  `json:"enabled"`
            User string  `json:"user"`
            Username string  `json:"username"`
         
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

        
            DisplayName string  `json:"display_name"`
            Documents StoreDocuments  `json:"documents"`
            NotificationEmails []string  `json:"notification_emails"`
            EwaybillPortalDetails map[string]interface{}  `json:"ewaybill_portal_details"`
            Timing []map[string]interface{}  `json:"timing"`
            ProductReturnConfig map[string]interface{}  `json:"product_return_config"`
            AdditionalContactDetails map[string]interface{}  `json:"additional_contact_details"`
            Stage string  `json:"stage"`
            GstNumber string  `json:"gst_number"`
            EinvoicePortalDetails EInvoicePortalDetails  `json:"einvoice_portal_details"`
            GstCredentials StoreGstCredentials  `json:"gst_credentials"`
         
    }
    
    // Store used by Order
    type Store struct {

        
            MallArea string  `json:"mall_area"`
            Address2 string  `json:"address2"`
            CreatedAt string  `json:"created_at"`
            IsActive bool  `json:"is_active"`
            LocationType string  `json:"location_type"`
            PackagingMaterialCount float64  `json:"packaging_material_count"`
            Longitude float64  `json:"longitude"`
            MallName string  `json:"mall_name"`
            Pincode string  `json:"pincode"`
            UpdatedAt string  `json:"updated_at"`
            BrandStoreTags []string  `json:"brand_store_tags"`
            Address1 string  `json:"address1"`
            Phone float64  `json:"phone"`
            ContactPerson string  `json:"contact_person"`
            CompanyID float64  `json:"company_id"`
            StoreAddressJson StoreAddress  `json:"store_address_json"`
            StoreEmail string  `json:"store_email"`
            IsEnabledForRecon bool  `json:"is_enabled_for_recon"`
            ParentStoreID float64  `json:"parent_store_id"`
            SID string  `json:"s_id"`
            FulfillmentChannel string  `json:"fulfillment_channel"`
            City string  `json:"city"`
            Latitude float64  `json:"latitude"`
            LoginUsername string  `json:"login_username"`
            BrandID interface{}  `json:"brand_id"`
            Name string  `json:"name"`
            Code string  `json:"code"`
            OrderIntegrationID string  `json:"order_integration_id"`
            VatNo string  `json:"vat_no"`
            IsArchived bool  `json:"is_archived"`
            State string  `json:"state"`
            Meta StoreMeta  `json:"meta"`
            StoreActiveFrom string  `json:"store_active_from"`
            Country string  `json:"country"`
            AlohomoraUserID float64  `json:"alohomora_user_id"`
         
    }
    
    // Brand used by Order
    type Brand struct {

        
            Company string  `json:"company"`
            IsVirtualInvoice bool  `json:"is_virtual_invoice"`
            Logo string  `json:"logo"`
            BrandName string  `json:"brand_name"`
            StartDate string  `json:"start_date"`
            CreditNoteAllowed bool  `json:"credit_note_allowed"`
            CreditNoteExpiryDays float64  `json:"credit_note_expiry_days"`
            CreatedOn float64  `json:"created_on"`
            PickupLocation string  `json:"pickup_location"`
            InvoicePrefix string  `json:"invoice_prefix"`
            BrandID float64  `json:"brand_id"`
            ModifiedOn float64  `json:"modified_on"`
            ScriptLastRan string  `json:"script_last_ran"`
         
    }
    
    // Weight used by Order
    type Weight struct {

        
            Shipping float64  `json:"shipping"`
            Unit string  `json:"unit"`
            IsDefault bool  `json:"is_default"`
         
    }
    
    // ReturnConfig used by Order
    type ReturnConfig struct {

        
            Returnable bool  `json:"returnable"`
            Time float64  `json:"time"`
            Unit string  `json:"unit"`
         
    }
    
    // Article used by Order
    type Article struct {

        
            Size string  `json:"size"`
            UID string  `json:"uid"`
            Identifiers Identifier  `json:"identifiers"`
            ID string  `json:"_id"`
            Dimensions Dimensions  `json:"dimensions"`
            Weight Weight  `json:"weight"`
            EspModified interface{}  `json:"esp_modified"`
            ASet map[string]interface{}  `json:"a_set"`
            ChildDetails map[string]interface{}  `json:"child_details"`
            RawMeta interface{}  `json:"raw_meta"`
            SellerIdentifier string  `json:"seller_identifier"`
            ReturnConfig ReturnConfig  `json:"return_config"`
            Code string  `json:"code"`
            IsSet bool  `json:"is_set"`
         
    }
    
    // BagGSTDetails used by Order
    type BagGSTDetails struct {

        
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            SgstTaxPercentage float64  `json:"sgst_tax_percentage"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            GstTag string  `json:"gst_tag"`
            CgstTaxPercentage float64  `json:"cgst_tax_percentage"`
            HsnCodeID string  `json:"hsn_code_id"`
            GstinCode string  `json:"gstin_code"`
            IgstTaxPercentage float64  `json:"igst_tax_percentage"`
            GstFee float64  `json:"gst_fee"`
            ValueOfGood float64  `json:"value_of_good"`
            HsnCode string  `json:"hsn_code"`
            CgstGstFee string  `json:"cgst_gst_fee"`
            IgstGstFee string  `json:"igst_gst_fee"`
            SgstGstFee string  `json:"sgst_gst_fee"`
            IsDefaultHsnCode bool  `json:"is_default_hsn_code"`
         
    }
    
    // AffiliateBagDetails used by Order
    type AffiliateBagDetails struct {

        
            AffiliateMeta AffiliateMeta  `json:"affiliate_meta"`
            EmployeeDiscount float64  `json:"employee_discount"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
         
    }
    
    // BagReturnableCancelableStatus used by Order
    type BagReturnableCancelableStatus struct {

        
            IsReturnable bool  `json:"is_returnable"`
            CanBeCancelled bool  `json:"can_be_cancelled"`
            IsActive bool  `json:"is_active"`
            EnableTracking bool  `json:"enable_tracking"`
            IsCustomerReturnAllowed bool  `json:"is_customer_return_allowed"`
         
    }
    
    // BagDetailsPlatformResponse used by Order
    type BagDetailsPlatformResponse struct {

        
            CurrentOperationalStatus BagStatusHistory  `json:"current_operational_status"`
            Reasons []map[string]interface{}  `json:"reasons"`
            Prices Prices  `json:"prices"`
            Item Item  `json:"item"`
            BagStatus []BagStatusHistory  `json:"bag_status"`
            Quantity float64  `json:"quantity"`
            NoOfBagsOrder float64  `json:"no_of_bags_order"`
            ArticleDetails ArticleDetails  `json:"article_details"`
            OperationalStatus string  `json:"operational_status"`
            Dates Dates  `json:"dates"`
            RestorePromos map[string]interface{}  `json:"restore_promos"`
            OrderingStore Store  `json:"ordering_store"`
            BagUpdateTime float64  `json:"bag_update_time"`
            OriginalBagList []float64  `json:"original_bag_list"`
            BType string  `json:"b_type"`
            Brand Brand  `json:"brand"`
            QcRequired interface{}  `json:"qc_required"`
            Article Article  `json:"article"`
            JourneyType string  `json:"journey_type"`
            EntityType string  `json:"entity_type"`
            FinancialBreakup []FinancialBreakup  `json:"financial_breakup"`
            RestoreCoupon bool  `json:"restore_coupon"`
            BID float64  `json:"b_id"`
            BagStatusHistory BagStatusHistory  `json:"bag_status_history"`
            DisplayName string  `json:"display_name"`
            LineNumber float64  `json:"line_number"`
            Tags []string  `json:"tags"`
            Identifier string  `json:"identifier"`
            ParentPromoBags map[string]interface{}  `json:"parent_promo_bags"`
            AppliedPromos []map[string]interface{}  `json:"applied_promos"`
            OrderIntegrationID string  `json:"order_integration_id"`
            GstDetails BagGSTDetails  `json:"gst_details"`
            AffiliateBagDetails AffiliateBagDetails  `json:"affiliate_bag_details"`
            Meta BagMeta  `json:"meta"`
            SellerIdentifier string  `json:"seller_identifier"`
            Status BagReturnableCancelableStatus  `json:"status"`
            CurrentStatus BagStatusHistory  `json:"current_status"`
            AffiliateDetails AffiliateDetails  `json:"affiliate_details"`
            ShipmentID string  `json:"shipment_id"`
         
    }
    
    // ErrorResponse used by Order
    type ErrorResponse struct {

        
            Error string  `json:"error"`
            Message string  `json:"message"`
         
    }
    
    // Page1 used by Order
    type Page1 struct {

        
            Size float64  `json:"size"`
            Current float64  `json:"current"`
            HasNext bool  `json:"has_next"`
            PageType string  `json:"page_type"`
            ItemTotal float64  `json:"item_total"`
         
    }
    
    // GetBagsPlatformResponse used by Order
    type GetBagsPlatformResponse struct {

        
            Items []BagDetailsPlatformResponse  `json:"items"`
            Page Page1  `json:"page"`
         
    }
    
    // GeneratePosOrderReceiptResponse used by Order
    type GeneratePosOrderReceiptResponse struct {

        
            Success bool  `json:"success"`
            InvoiceReceipt string  `json:"invoice_receipt"`
            OrderID string  `json:"order_id"`
            PaymentReceipt string  `json:"payment_receipt"`
         
    }
    
    // InvalidateShipmentCachePayload used by Order
    type InvalidateShipmentCachePayload struct {

        
            ShipmentIds []string  `json:"shipment_ids"`
         
    }
    
    // InvalidateShipmentCacheNestedResponse used by Order
    type InvalidateShipmentCacheNestedResponse struct {

        
            ShipmentID string  `json:"shipment_id"`
            Status float64  `json:"status"`
            Error string  `json:"error"`
            Message string  `json:"message"`
         
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
         
    }
    
    // StoreReassign used by Order
    type StoreReassign struct {

        
            ItemID string  `json:"item_id"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            AffiliateID string  `json:"affiliate_id"`
            ReasonIds []float64  `json:"reason_ids"`
            StoreID float64  `json:"store_id"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
            BagID float64  `json:"bag_id"`
            SetID string  `json:"set_id"`
            MongoArticleID string  `json:"mongo_article_id"`
            FyndOrderID string  `json:"fynd_order_id"`
         
    }
    
    // StoreReassignResponse used by Order
    type StoreReassignResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // Entities used by Order
    type Entities struct {

        
            AffiliateOrderID string  `json:"affiliate_order_id"`
            AffiliateID string  `json:"affiliate_id"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
            ID string  `json:"id"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            ReasonText string  `json:"reason_text"`
         
    }
    
    // UpdateShipmentLockPayload used by Order
    type UpdateShipmentLockPayload struct {

        
            EntityType string  `json:"entity_type"`
            ActionType string  `json:"action_type"`
            Entities []Entities  `json:"entities"`
            Action string  `json:"action"`
         
    }
    
    // OriginalFilter used by Order
    type OriginalFilter struct {

        
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            AffiliateID string  `json:"affiliate_id"`
         
    }
    
    // Bags used by Order
    type Bags struct {

        
            BagID float64  `json:"bag_id"`
            IsLocked bool  `json:"is_locked"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
         
    }
    
    // CheckResponse used by Order
    type CheckResponse struct {

        
            IsBagLocked bool  `json:"is_bag_locked"`
            IsShipmentLocked bool  `json:"is_shipment_locked"`
            AffiliateID string  `json:"affiliate_id"`
            OriginalFilter OriginalFilter  `json:"original_filter"`
            ShipmentID string  `json:"shipment_id"`
            LockStatus bool  `json:"lock_status"`
            Status string  `json:"status"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            Bags []Bags  `json:"bags"`
         
    }
    
    // UpdateShipmentLockResponse used by Order
    type UpdateShipmentLockResponse struct {

        
            Success bool  `json:"success"`
            CheckResponse []CheckResponse  `json:"check_response"`
            Message string  `json:"message"`
         
    }
    
    // AnnouncementResponse used by Order
    type AnnouncementResponse struct {

        
            CompanyID float64  `json:"company_id"`
            ToDatetime string  `json:"to_datetime"`
            CreatedAt string  `json:"created_at"`
            Title string  `json:"title"`
            Description string  `json:"description"`
            ID float64  `json:"id"`
            LogoURL string  `json:"logo_url"`
            PlatformName string  `json:"platform_name"`
            PlatformID string  `json:"platform_id"`
            FromDatetime string  `json:"from_datetime"`
         
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

        
            Status bool  `json:"status"`
            CallID string  `json:"call_id"`
         
    }
    
    // ProductsReasonsFilters used by Order
    type ProductsReasonsFilters struct {

        
            Identifier string  `json:"identifier"`
            Quantity float64  `json:"quantity"`
            LineNumber float64  `json:"line_number"`
         
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

        
            LockAfterTransition bool  `json:"lock_after_transition"`
            Task bool  `json:"task"`
            Statuses []StatuesRequest  `json:"statuses"`
            UnlockBeforeTransition bool  `json:"unlock_before_transition"`
            ForceTransition bool  `json:"force_transition"`
         
    }
    
    // ShipmentsResponse used by Order
    type ShipmentsResponse struct {

        
            Identifier string  `json:"identifier"`
            StackTrace string  `json:"stack_trace"`
            FinalState map[string]interface{}  `json:"final_state"`
            Exception string  `json:"exception"`
            Code string  `json:"code"`
            Meta map[string]interface{}  `json:"meta"`
            Status float64  `json:"status"`
            Message string  `json:"message"`
         
    }
    
    // StatuesResponse used by Order
    type StatuesResponse struct {

        
            Shipments []ShipmentsResponse  `json:"shipments"`
         
    }
    
    // UpdateShipmentStatusResponseBody used by Order
    type UpdateShipmentStatusResponseBody struct {

        
            Statuses []StatuesResponse  `json:"statuses"`
         
    }
    
    // AffiliateInventoryStoreConfig used by Order
    type AffiliateInventoryStoreConfig struct {

        
            Store map[string]interface{}  `json:"store"`
         
    }
    
    // AffiliateInventoryLogisticsConfig used by Order
    type AffiliateInventoryLogisticsConfig struct {

        
            DpAssignment bool  `json:"dp_assignment"`
         
    }
    
    // AffiliateInventoryArticleAssignmentConfig used by Order
    type AffiliateInventoryArticleAssignmentConfig struct {

        
            PostOrderReassignment bool  `json:"post_order_reassignment"`
         
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
    
    // AffiliateInventoryConfig used by Order
    type AffiliateInventoryConfig struct {

        
            Inventory AffiliateInventoryStoreConfig  `json:"inventory"`
            Logistics AffiliateInventoryLogisticsConfig  `json:"logistics"`
            ArticleAssignment AffiliateInventoryArticleAssignmentConfig  `json:"article_assignment"`
            Payment AffiliateInventoryPaymentConfig  `json:"payment"`
            Order AffiliateInventoryOrderConfig  `json:"order"`
         
    }
    
    // AffiliateAppConfigMeta used by Order
    type AffiliateAppConfigMeta struct {

        
            Value string  `json:"value"`
            Name string  `json:"name"`
         
    }
    
    // AffiliateAppConfig used by Order
    type AffiliateAppConfig struct {

        
            Name string  `json:"name"`
            Owner string  `json:"owner"`
            CreatedAt string  `json:"created_at"`
            Secret string  `json:"secret"`
            Description string  `json:"description"`
            ID string  `json:"id"`
            UpdatedAt string  `json:"updated_at"`
            Meta []AffiliateAppConfigMeta  `json:"meta"`
            Token string  `json:"token"`
         
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

        
            Affiliate Affiliate  `json:"affiliate"`
            ArticleLookup string  `json:"article_lookup"`
            StoreLookup string  `json:"store_lookup"`
            CreateUser bool  `json:"create_user"`
            AffiliateStoreIDMapping []AffiliateStoreIdMapping  `json:"affiliate_store_id_mapping"`
            BagEndState string  `json:"bag_end_state"`
         
    }
    
    // OrderUser used by Order
    type OrderUser struct {

        
            Address2 string  `json:"address2"`
            FirstName string  `json:"first_name"`
            Mobile float64  `json:"mobile"`
            Country string  `json:"country"`
            LastName string  `json:"last_name"`
            Address1 string  `json:"address1"`
            Pincode string  `json:"pincode"`
            State string  `json:"state"`
            Phone float64  `json:"phone"`
            City string  `json:"city"`
            Email string  `json:"email"`
         
    }
    
    // UserData used by Order
    type UserData struct {

        
            ShippingUser OrderUser  `json:"shipping_user"`
            BillingUser OrderUser  `json:"billing_user"`
         
    }
    
    // ArticleDetails1 used by Order
    type ArticleDetails1 struct {

        
            ID string  `json:"_id"`
            Category map[string]interface{}  `json:"category"`
            BrandID float64  `json:"brand_id"`
            Dimension map[string]interface{}  `json:"dimension"`
            Attributes map[string]interface{}  `json:"attributes"`
            Quantity float64  `json:"quantity"`
            Weight map[string]interface{}  `json:"weight"`
         
    }
    
    // ShipmentDetails used by Order
    type ShipmentDetails struct {

        
            FulfillmentID float64  `json:"fulfillment_id"`
            Shipments float64  `json:"shipments"`
            Meta map[string]interface{}  `json:"meta"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            DpID float64  `json:"dp_id"`
            BoxType string  `json:"box_type"`
            Articles []ArticleDetails1  `json:"articles"`
         
    }
    
    // LocationDetails used by Order
    type LocationDetails struct {

        
            FulfillmentType string  `json:"fulfillment_type"`
            FulfillmentID float64  `json:"fulfillment_id"`
            Articles []ArticleDetails1  `json:"articles"`
         
    }
    
    // ShipmentConfig used by Order
    type ShipmentConfig struct {

        
            ToPincode string  `json:"to_pincode"`
            Source string  `json:"source"`
            Identifier string  `json:"identifier"`
            Action string  `json:"action"`
            Journey string  `json:"journey"`
            Shipment []ShipmentDetails  `json:"shipment"`
            PaymentMode string  `json:"payment_mode"`
            LocationDetails LocationDetails  `json:"location_details"`
         
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
    
    // MarketPlacePdf used by Order
    type MarketPlacePdf struct {

        
            Invoice string  `json:"invoice"`
            Label string  `json:"label"`
         
    }
    
    // AffiliateBag used by Order
    type AffiliateBag struct {

        
            AffiliateMeta map[string]interface{}  `json:"affiliate_meta"`
            ItemID float64  `json:"item_id"`
            AvlQty float64  `json:"avl_qty"`
            AmountPaid float64  `json:"amount_paid"`
            Discount float64  `json:"discount"`
            ID string  `json:"_id"`
            ItemSize string  `json:"item_size"`
            PriceMarked float64  `json:"price_marked"`
            UnitPrice float64  `json:"unit_price"`
            TransferPrice float64  `json:"transfer_price"`
            DeliveryCharge float64  `json:"delivery_charge"`
            PdfLinks MarketPlacePdf  `json:"pdf_links"`
            FyndStoreID string  `json:"fynd_store_id"`
            AffiliateStoreID string  `json:"affiliate_store_id"`
            PriceEffective float64  `json:"price_effective"`
            StoreID float64  `json:"store_id"`
            HsnCodeID string  `json:"hsn_code_id"`
            SellerIdentifier string  `json:"seller_identifier"`
            Quantity float64  `json:"quantity"`
            Sku string  `json:"sku"`
            Identifier map[string]interface{}  `json:"identifier"`
            CompanyID float64  `json:"company_id"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // OrderInfo used by Order
    type OrderInfo struct {

        
            AffiliateOrderID string  `json:"affiliate_order_id"`
            Discount float64  `json:"discount"`
            DeliveryCharges float64  `json:"delivery_charges"`
            User UserData  `json:"user"`
            BillingAddress OrderUser  `json:"billing_address"`
            CodCharges float64  `json:"cod_charges"`
            Items map[string]interface{}  `json:"items"`
            ShippingAddress OrderUser  `json:"shipping_address"`
            Coupon string  `json:"coupon"`
            Shipment ShipmentData  `json:"shipment"`
            OrderPriority OrderPriority  `json:"order_priority"`
            Payment map[string]interface{}  `json:"payment"`
            PaymentMode string  `json:"payment_mode"`
            OrderValue float64  `json:"order_value"`
            Bags []AffiliateBag  `json:"bags"`
         
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

        
            DisplayText string  `json:"display_text"`
            ID float64  `json:"id"`
            Slug string  `json:"slug"`
            Description string  `json:"description"`
         
    }
    
    // GetActionsResponse used by Order
    type GetActionsResponse struct {

        
            Permissions ActionInfo  `json:"permissions"`
         
    }
    
    // HistoryDict used by Order
    type HistoryDict struct {

        
            Createdat string  `json:"createdat"`
            Type string  `json:"type"`
            TicketID string  `json:"ticket_id"`
            L3Detail string  `json:"l3_detail"`
            L2Detail string  `json:"l2_detail"`
            User string  `json:"user"`
            TicketURL string  `json:"ticket_url"`
            BagID float64  `json:"bag_id"`
            L1Detail string  `json:"l1_detail"`
            Message string  `json:"message"`
         
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
    
    // SmsDataPayload used by Order
    type SmsDataPayload struct {

        
            CountryCode string  `json:"country_code"`
            AmountPaid float64  `json:"amount_paid"`
            PhoneNumber float64  `json:"phone_number"`
            ShipmentID float64  `json:"shipment_id"`
            BrandName string  `json:"brand_name"`
            CustomerName string  `json:"customer_name"`
            PaymentMode string  `json:"payment_mode"`
            OrderID string  `json:"order_id"`
            Message string  `json:"message"`
         
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

        
            ShipmentID string  `json:"shipment_id"`
            BagList []float64  `json:"bag_list"`
            ID float64  `json:"id"`
            Remarks string  `json:"remarks"`
            Meta Meta1  `json:"meta"`
            Status string  `json:"status"`
         
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
    
    // ManualAssignDPToShipment used by Order
    type ManualAssignDPToShipment struct {

        
            OrderType string  `json:"order_type"`
            ShipmentIds []string  `json:"shipment_ids"`
            DpID float64  `json:"dp_id"`
            QcRequired string  `json:"qc_required"`
         
    }
    
    // ManualAssignDPToShipmentResponse used by Order
    type ManualAssignDPToShipmentResponse struct {

        
            Success string  `json:"success"`
            Errors []string  `json:"errors"`
         
    }
    
    // ShippingInfo used by Order
    type ShippingInfo struct {

        
            HouseNo string  `json:"house_no"`
            FloorNo string  `json:"floor_no"`
            Country string  `json:"country"`
            Gender string  `json:"gender"`
            PrimaryMobileNumber string  `json:"primary_mobile_number"`
            Slot []map[string]interface{}  `json:"slot"`
            Address1 string  `json:"address1"`
            Pincode string  `json:"pincode"`
            GeoLocation map[string]interface{}  `json:"geo_location"`
            AlternateEmail string  `json:"alternate_email"`
            LastName string  `json:"last_name"`
            City string  `json:"city"`
            FirstName string  `json:"first_name"`
            StateCode string  `json:"state_code"`
            AlternateMobileNumber string  `json:"alternate_mobile_number"`
            Landmark string  `json:"landmark"`
            PrimaryEmail string  `json:"primary_email"`
            Address2 string  `json:"address2"`
            CountryCode string  `json:"country_code"`
            CustomerCode string  `json:"customer_code"`
            ShippingType string  `json:"shipping_type"`
            ExternalCustomerCode string  `json:"external_customer_code"`
            Title string  `json:"title"`
            State string  `json:"state"`
            MiddleName string  `json:"middle_name"`
            AddressType string  `json:"address_type"`
         
    }
    
    // ProcessingDates used by Order
    type ProcessingDates struct {

        
            PackByDate string  `json:"pack_by_date"`
            CustomerPickupSlot map[string]interface{}  `json:"customer_pickup_slot"`
            ConfirmByDate string  `json:"confirm_by_date"`
            DispatchByDate string  `json:"dispatch_by_date"`
            DispatchAfterDate string  `json:"dispatch_after_date"`
            DpPickupSlot map[string]interface{}  `json:"dp_pickup_slot"`
         
    }
    
    // Tax used by Order
    type Tax struct {

        
            Rate float64  `json:"rate"`
            Amount map[string]interface{}  `json:"amount"`
            Name string  `json:"name"`
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

        
            CustomMessasge string  `json:"custom_messasge"`
            ExternalLineID string  `json:"external_line_id"`
            SellerIdentifier string  `json:"seller_identifier"`
            Meta map[string]interface{}  `json:"meta"`
            Charges []Charge  `json:"charges"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // Shipment used by Order
    type Shipment struct {

        
            ExternalShipmentID string  `json:"external_shipment_id"`
            Priority float64  `json:"priority"`
            Meta map[string]interface{}  `json:"meta"`
            ProcessingDates ProcessingDates  `json:"processing_dates"`
            LineItems []LineItem  `json:"line_items"`
            LocationID float64  `json:"location_id"`
         
    }
    
    // TaxInfo used by Order
    type TaxInfo struct {

        
            B2bGstinNumber string  `json:"b2b_gstin_number"`
            Gstin string  `json:"gstin"`
         
    }
    
    // BillingInfo used by Order
    type BillingInfo struct {

        
            HouseNo string  `json:"house_no"`
            FloorNo string  `json:"floor_no"`
            Country string  `json:"country"`
            Gender string  `json:"gender"`
            PrimaryMobileNumber string  `json:"primary_mobile_number"`
            Address1 string  `json:"address1"`
            Pincode string  `json:"pincode"`
            AlternateEmail string  `json:"alternate_email"`
            LastName string  `json:"last_name"`
            City string  `json:"city"`
            FirstName string  `json:"first_name"`
            StateCode string  `json:"state_code"`
            AlternateMobileNumber string  `json:"alternate_mobile_number"`
            PrimaryEmail string  `json:"primary_email"`
            Address2 string  `json:"address2"`
            CountryCode string  `json:"country_code"`
            CustomerCode string  `json:"customer_code"`
            ExternalCustomerCode string  `json:"external_customer_code"`
            Title string  `json:"title"`
            State string  `json:"state"`
            MiddleName string  `json:"middle_name"`
         
    }
    
    // PaymentMethod used by Order
    type PaymentMethod struct {

        
            CollectBy string  `json:"collect_by"`
            Name string  `json:"name"`
            TransactionData map[string]interface{}  `json:"transaction_data"`
            RefundBy string  `json:"refund_by"`
            Mode string  `json:"mode"`
            Meta map[string]interface{}  `json:"meta"`
            Amount float64  `json:"amount"`
         
    }
    
    // PaymentInfo used by Order
    type PaymentInfo struct {

        
            PrimaryMode string  `json:"primary_mode"`
            PaymentMethods []PaymentMethod  `json:"payment_methods"`
         
    }
    
    // CreateOrderAPI used by Order
    type CreateOrderAPI struct {

        
            ShippingInfo ShippingInfo  `json:"shipping_info"`
            Shipments []Shipment  `json:"shipments"`
            TaxInfo TaxInfo  `json:"tax_info"`
            ExternalOrderID string  `json:"external_order_id"`
            BillingInfo BillingInfo  `json:"billing_info"`
            ExternalCreationDate string  `json:"external_creation_date"`
            Meta map[string]interface{}  `json:"meta"`
            CurrencyInfo map[string]interface{}  `json:"currency_info"`
            Charges []Charge  `json:"charges"`
            PaymentInfo PaymentInfo  `json:"payment_info"`
         
    }
    
    // CreateOrderErrorReponse used by Order
    type CreateOrderErrorReponse struct {

        
            Info interface{}  `json:"info"`
            StackTrace string  `json:"stack_trace"`
            Exception string  `json:"exception"`
            Code string  `json:"code"`
            RequestID string  `json:"request_id"`
            Meta string  `json:"meta"`
            Status float64  `json:"status"`
            Message string  `json:"message"`
         
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

        
            ShipmentAssignment string  `json:"shipment_assignment"`
            DpConfiguration DpConfiguration  `json:"dp_configuration"`
            LocationReassignment bool  `json:"location_reassignment"`
            LogoURL map[string]interface{}  `json:"logo_url"`
            LockStates []string  `json:"lock_states"`
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

        
            EndDate string  `json:"end_date"`
            Mobile float64  `json:"mobile"`
            OrderDetails []FyndOrderIdList  `json:"order_details"`
            StartDate string  `json:"start_date"`
         
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

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Words []string  `json:"words"`
            UID string  `json:"uid"`
            AppID string  `json:"app_id"`
            IsActive bool  `json:"is_active"`
            Result map[string]interface{}  `json:"result"`
         
    }
    
    // GetSearchWordsDetailResponse used by Catalog
    type GetSearchWordsDetailResponse struct {

        
            Page Page  `json:"page"`
            Items GetSearchWordsData  `json:"items"`
         
    }
    
    // ErrorResponse used by Catalog
    type ErrorResponse struct {

        
            Error string  `json:"error"`
            Message string  `json:"message"`
            Status float64  `json:"status"`
            Code string  `json:"code"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // SearchKeywordResult used by Catalog
    type SearchKeywordResult struct {

        
            SortOn string  `json:"sort_on"`
            Query map[string]interface{}  `json:"query"`
         
    }
    
    // CreateSearchKeyword used by Catalog
    type CreateSearchKeyword struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Words []string  `json:"words"`
            AppID string  `json:"app_id"`
            IsActive bool  `json:"is_active"`
            Result SearchKeywordResult  `json:"result"`
         
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
    
    // GetAutocompleteWordsData used by Catalog
    type GetAutocompleteWordsData struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Words []string  `json:"words"`
            UID string  `json:"uid"`
            AppID string  `json:"app_id"`
            Results []map[string]interface{}  `json:"results"`
         
    }
    
    // GetAutocompleteWordsResponse used by Catalog
    type GetAutocompleteWordsResponse struct {

        
            Page Page  `json:"page"`
            Items []GetAutocompleteWordsData  `json:"items"`
         
    }
    
    // AutocompletePageAction used by Catalog
    type AutocompletePageAction struct {

        
            URL string  `json:"url"`
            Params map[string]interface{}  `json:"params"`
            Type string  `json:"type"`
            Query map[string]interface{}  `json:"query"`
         
    }
    
    // AutocompleteAction used by Catalog
    type AutocompleteAction struct {

        
            Page AutocompletePageAction  `json:"page"`
            Type string  `json:"type"`
         
    }
    
    // Media used by Catalog
    type Media struct {

        
            URL string  `json:"url"`
            Type string  `json:"type"`
            AspectRatio string  `json:"aspect_ratio"`
         
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

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Words []string  `json:"words"`
            AppID string  `json:"app_id"`
            IsActive bool  `json:"is_active"`
            Results []AutocompleteResult  `json:"results"`
         
    }
    
    // CreateAutocompleteWordsResponse used by Catalog
    type CreateAutocompleteWordsResponse struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Results []map[string]interface{}  `json:"results"`
            Words []string  `json:"words"`
            AppID string  `json:"app_id"`
         
    }
    
    // ProductBundleItem used by Catalog
    type ProductBundleItem struct {

        
            ProductUID float64  `json:"product_uid"`
            AllowRemove bool  `json:"allow_remove"`
            MaxQuantity float64  `json:"max_quantity"`
            AutoSelect bool  `json:"auto_select"`
            MinQuantity float64  `json:"min_quantity"`
            AutoAddToCart bool  `json:"auto_add_to_cart"`
         
    }
    
    // GetProductBundleCreateResponse used by Catalog
    type GetProductBundleCreateResponse struct {

        
            CompanyID float64  `json:"company_id"`
            ID string  `json:"id"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Slug string  `json:"slug"`
            Choice string  `json:"choice"`
            PageVisibility []string  `json:"page_visibility"`
            Logo string  `json:"logo"`
            Meta map[string]interface{}  `json:"meta"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            Name string  `json:"name"`
            Products []ProductBundleItem  `json:"products"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // GetProductBundleListingResponse used by Catalog
    type GetProductBundleListingResponse struct {

        
            Page Page  `json:"page"`
            Items []GetProductBundleCreateResponse  `json:"items"`
         
    }
    
    // ProductBundleRequest used by Catalog
    type ProductBundleRequest struct {

        
            CompanyID float64  `json:"company_id"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Slug string  `json:"slug"`
            Choice string  `json:"choice"`
            PageVisibility []string  `json:"page_visibility"`
            Logo string  `json:"logo"`
            Meta map[string]interface{}  `json:"meta"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            Name string  `json:"name"`
            Products []ProductBundleItem  `json:"products"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // Size used by Catalog
    type Size struct {

        
            Value string  `json:"value"`
            Display string  `json:"display"`
            IsAvailable bool  `json:"is_available"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // LimitedProductData used by Catalog
    type LimitedProductData struct {

        
            Sizes []string  `json:"sizes"`
            Identifier map[string]interface{}  `json:"identifier"`
            CountryOfOrigin string  `json:"country_of_origin"`
            Images []string  `json:"images"`
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
            Quantity float64  `json:"quantity"`
            Name string  `json:"name"`
            ShortDescription string  `json:"short_description"`
            Attributes map[string]interface{}  `json:"attributes"`
            Price map[string]interface{}  `json:"price"`
            ItemCode string  `json:"item_code"`
         
    }
    
    // Price used by Catalog
    type Price struct {

        
            MinMarked float64  `json:"min_marked"`
            MinEffective float64  `json:"min_effective"`
            MaxMarked float64  `json:"max_marked"`
            Currency string  `json:"currency"`
            MaxEffective float64  `json:"max_effective"`
         
    }
    
    // GetProducts used by Catalog
    type GetProducts struct {

        
            Sizes []Size  `json:"sizes"`
            ProductUID float64  `json:"product_uid"`
            AllowRemove bool  `json:"allow_remove"`
            MaxQuantity float64  `json:"max_quantity"`
            ProductDetails LimitedProductData  `json:"product_details"`
            AutoSelect bool  `json:"auto_select"`
            Price Price  `json:"price"`
            MinQuantity float64  `json:"min_quantity"`
            AutoAddToCart bool  `json:"auto_add_to_cart"`
         
    }
    
    // GetProductBundleResponse used by Catalog
    type GetProductBundleResponse struct {

        
            CompanyID float64  `json:"company_id"`
            Slug string  `json:"slug"`
            PageVisibility []string  `json:"page_visibility"`
            Choice string  `json:"choice"`
            Logo string  `json:"logo"`
            Meta map[string]interface{}  `json:"meta"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            Name string  `json:"name"`
            Products []GetProducts  `json:"products"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // ProductBundleUpdateRequest used by Catalog
    type ProductBundleUpdateRequest struct {

        
            CompanyID float64  `json:"company_id"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            ModifiedOn string  `json:"modified_on"`
            Slug string  `json:"slug"`
            Choice string  `json:"choice"`
            PageVisibility []string  `json:"page_visibility"`
            Logo string  `json:"logo"`
            Meta map[string]interface{}  `json:"meta"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            Name string  `json:"name"`
            Products []ProductBundleItem  `json:"products"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // ListSizeGuide used by Catalog
    type ListSizeGuide struct {

        
            Page map[string]interface{}  `json:"page"`
            Items []map[string]interface{}  `json:"items"`
         
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

        
            Active bool  `json:"active"`
            Tag string  `json:"tag"`
            Image string  `json:"image"`
            CompanyID float64  `json:"company_id"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            CreatedOn string  `json:"created_on"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ModifiedOn string  `json:"modified_on"`
            ID string  `json:"id"`
            Description string  `json:"description"`
            Guide Guide  `json:"guide"`
            Title string  `json:"title"`
            Name string  `json:"name"`
            Subtitle string  `json:"subtitle"`
            BrandID float64  `json:"brand_id"`
         
    }
    
    // SuccessResponse used by Catalog
    type SuccessResponse struct {

        
            Success bool  `json:"success"`
         
    }
    
    // SizeGuideResponse used by Catalog
    type SizeGuideResponse struct {

        
            Active bool  `json:"active"`
            CompanyID float64  `json:"company_id"`
            ID string  `json:"id"`
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Guide map[string]interface{}  `json:"guide"`
            Tag string  `json:"tag"`
            Name string  `json:"name"`
            Subtitle string  `json:"subtitle"`
            BrandID float64  `json:"brand_id"`
            Title string  `json:"title"`
         
    }
    
    // SEOData used by Catalog
    type SEOData struct {

        
            Description interface{}  `json:"description"`
            Title interface{}  `json:"title"`
         
    }
    
    // MOQData used by Catalog
    type MOQData struct {

        
            Maximum float64  `json:"maximum"`
            IncrementUnit float64  `json:"increment_unit"`
            Minimum float64  `json:"minimum"`
         
    }
    
    // OwnerAppItemResponse used by Catalog
    type OwnerAppItemResponse struct {

        
            IsCod bool  `json:"is_cod"`
            Seo SEOData  `json:"seo"`
            IsGift bool  `json:"is_gift"`
            AltText map[string]interface{}  `json:"alt_text"`
            Moq MOQData  `json:"moq"`
         
    }
    
    // ApplicationItemSEO used by Catalog
    type ApplicationItemSEO struct {

        
            Description interface{}  `json:"description"`
            Title interface{}  `json:"title"`
         
    }
    
    // MetaFields used by Catalog
    type MetaFields struct {

        
            Key interface{}  `json:"key"`
            Value interface{}  `json:"value"`
         
    }
    
    // ApplicationItemMOQ used by Catalog
    type ApplicationItemMOQ struct {

        
            Maximum float64  `json:"maximum"`
            IncrementUnit float64  `json:"increment_unit"`
            Minimum float64  `json:"minimum"`
         
    }
    
    // ApplicationItemMeta used by Catalog
    type ApplicationItemMeta struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            IsCod bool  `json:"is_cod"`
            Seo ApplicationItemSEO  `json:"seo"`
            CustomMeta []MetaFields  `json:"_custom_meta"`
            IsGift bool  `json:"is_gift"`
            AltText map[string]interface{}  `json:"alt_text"`
            Moq ApplicationItemMOQ  `json:"moq"`
         
    }
    
    // SuccessResponse1 used by Catalog
    type SuccessResponse1 struct {

        
            UID float64  `json:"uid"`
            Success bool  `json:"success"`
         
    }
    
    // GetConfigMetadataResponse used by Catalog
    type GetConfigMetadataResponse struct {

        
            Condition []map[string]interface{}  `json:"condition"`
            Values []map[string]interface{}  `json:"values"`
            Data []map[string]interface{}  `json:"data"`
         
    }
    
    // PageResponseType used by Catalog
    type PageResponseType struct {

        
            HasNext bool  `json:"has_next"`
            TotalCount float64  `json:"total_count"`
            Current float64  `json:"current"`
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

        
            Key string  `json:"key"`
            Priority float64  `json:"priority"`
            Slug string  `json:"slug"`
            DisplayType string  `json:"display_type"`
            Logo string  `json:"logo"`
            Unit string  `json:"unit"`
            Name string  `json:"name"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // AppConfigurationDetail used by Catalog
    type AppConfigurationDetail struct {

        
            Priority float64  `json:"priority"`
            Slug string  `json:"slug"`
            IsDefault bool  `json:"is_default"`
            Attributes []AttributeDetailsGroup  `json:"attributes"`
            TemplateSlugs []string  `json:"template_slugs"`
            Logo string  `json:"logo"`
            AppID string  `json:"app_id"`
            Name string  `json:"name"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // ConfigSuccessResponse used by Catalog
    type ConfigSuccessResponse struct {

        
            Message string  `json:"message"`
         
    }
    
    // AppConfigurationsSort used by Catalog
    type AppConfigurationsSort struct {

        
            Key string  `json:"key"`
            Priority float64  `json:"priority"`
            DefaultKey string  `json:"default_key"`
            IsDefault bool  `json:"is_default"`
            Logo string  `json:"logo"`
            AppID string  `json:"app_id"`
            Name string  `json:"name"`
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

        
            Key string  `json:"key"`
            Display string  `json:"display"`
         
    }
    
    // MetaDataListingSortResponse used by Catalog
    type MetaDataListingSortResponse struct {

        
            Data []MetaDataListingSortMetaResponse  `json:"data"`
         
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
    
    // MetaDataListingResponse used by Catalog
    type MetaDataListingResponse struct {

        
            Sort MetaDataListingSortResponse  `json:"sort"`
            Filter MetaDataListingFilterResponse  `json:"filter"`
         
    }
    
    // GetCatalogConfigurationDetailsProduct used by Catalog
    type GetCatalogConfigurationDetailsProduct struct {

        
            Compare map[string]interface{}  `json:"compare"`
            Similar map[string]interface{}  `json:"similar"`
            Variant map[string]interface{}  `json:"variant"`
            Detail map[string]interface{}  `json:"detail"`
         
    }
    
    // GetCatalogConfigurationMetaData used by Catalog
    type GetCatalogConfigurationMetaData struct {

        
            Listing MetaDataListingResponse  `json:"listing"`
            Product GetCatalogConfigurationDetailsProduct  `json:"product"`
         
    }
    
    // ConfigurationListingSortConfig used by Catalog
    type ConfigurationListingSortConfig struct {

        
            Key string  `json:"key"`
            Priority float64  `json:"priority"`
            Logo string  `json:"logo"`
            Name string  `json:"name"`
            IsActive bool  `json:"is_active"`
         
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

        
            Condition string  `json:"condition"`
            BucketPoints []ConfigurationBucketPoints  `json:"bucket_points"`
            Value string  `json:"value"`
            MapValues []map[string]interface{}  `json:"map_values"`
            Map map[string]interface{}  `json:"map"`
            Sort string  `json:"sort"`
         
    }
    
    // ConfigurationListingFilterConfig used by Catalog
    type ConfigurationListingFilterConfig struct {

        
            Key string  `json:"key"`
            Priority float64  `json:"priority"`
            Logo string  `json:"logo"`
            ValueConfig ConfigurationListingFilterValue  `json:"value_config"`
            Name string  `json:"name"`
            IsActive bool  `json:"is_active"`
            Type string  `json:"type"`
            DisplayName string  `json:"display_name"`
         
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

        
            Key string  `json:"key"`
            Priority float64  `json:"priority"`
            Size ProductSize  `json:"size"`
            Logo string  `json:"logo"`
            Title string  `json:"title"`
            IsActive bool  `json:"is_active"`
            Subtitle string  `json:"subtitle"`
         
    }
    
    // ConfigurationProductSimilar used by Catalog
    type ConfigurationProductSimilar struct {

        
            Config []ConfigurationProductConfig  `json:"config"`
         
    }
    
    // ConfigurationProductVariantConfig used by Catalog
    type ConfigurationProductVariantConfig struct {

        
            Key string  `json:"key"`
            Priority float64  `json:"priority"`
            DisplayType string  `json:"display_type"`
            Size ProductSize  `json:"size"`
            Logo string  `json:"logo"`
            Name string  `json:"name"`
            IsActive bool  `json:"is_active"`
         
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
            ID string  `json:"id"`
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Product ConfigurationProduct  `json:"product"`
            ConfigID string  `json:"config_id"`
            ConfigType string  `json:"config_type"`
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

        
            Listing ConfigurationListing  `json:"listing"`
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Product ConfigurationProduct  `json:"product"`
            ConfigID string  `json:"config_id"`
            ConfigType string  `json:"config_type"`
            AppID string  `json:"app_id"`
            Type string  `json:"type"`
         
    }
    
    // GetCatalogConfigurationDetailsSchemaListing used by Catalog
    type GetCatalogConfigurationDetailsSchemaListing struct {

        
            Sort map[string]interface{}  `json:"sort"`
            Filter map[string]interface{}  `json:"filter"`
         
    }
    
    // EntityConfiguration used by Catalog
    type EntityConfiguration struct {

        
            Listing GetCatalogConfigurationDetailsSchemaListing  `json:"listing"`
            ID string  `json:"id"`
            Product GetCatalogConfigurationDetailsProduct  `json:"product"`
            ConfigID string  `json:"config_id"`
            ConfigType string  `json:"config_type"`
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
            Value string  `json:"value"`
            IsSelected bool  `json:"is_selected"`
         
    }
    
    // ProductFiltersKey used by Catalog
    type ProductFiltersKey struct {

        
            Display string  `json:"display"`
            Operators []string  `json:"operators"`
            Logo string  `json:"logo"`
            Name string  `json:"name"`
            Kind string  `json:"kind"`
         
    }
    
    // ProductFiltersValue used by Catalog
    type ProductFiltersValue struct {

        
            IsSelected bool  `json:"is_selected"`
            Display string  `json:"display"`
            Min float64  `json:"min"`
            SelectedMax float64  `json:"selected_max"`
            QueryFormat string  `json:"query_format"`
            Count float64  `json:"count"`
            Value interface{}  `json:"value"`
            CurrencySymbol string  `json:"currency_symbol"`
            Max float64  `json:"max"`
            DisplayFormat string  `json:"display_format"`
            SelectedMin float64  `json:"selected_min"`
            CurrencyCode string  `json:"currency_code"`
         
    }
    
    // ProductFilters used by Catalog
    type ProductFilters struct {

        
            Key ProductFiltersKey  `json:"key"`
            Values []ProductFiltersValue  `json:"values"`
         
    }
    
    // GetCollectionQueryOptionResponse used by Catalog
    type GetCollectionQueryOptionResponse struct {

        
            Operators map[string]string  `json:"operators"`
            SortOn []ProductSortOn  `json:"sort_on"`
            Filters []ProductFilters  `json:"filters"`
         
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
    
    // Media1 used by Catalog
    type Media1 struct {

        
            URL string  `json:"url"`
            Meta map[string]interface{}  `json:"meta"`
            Type string  `json:"type"`
         
    }
    
    // GetCollectionDetailNest used by Catalog
    type GetCollectionDetailNest struct {

        
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Action Action  `json:"action"`
            Banners ImageUrls  `json:"banners"`
            AppID string  `json:"app_id"`
            IsActive bool  `json:"is_active"`
            AllowFacets bool  `json:"allow_facets"`
            Cron map[string]interface{}  `json:"cron"`
            UID string  `json:"uid"`
            Query []CollectionQuery  `json:"query"`
            Badge map[string]interface{}  `json:"badge"`
            Logo Media1  `json:"logo"`
            Meta map[string]interface{}  `json:"meta"`
            AllowSort bool  `json:"allow_sort"`
            Tag []string  `json:"tag"`
            Name string  `json:"name"`
            Priority float64  `json:"priority"`
            Slug string  `json:"slug"`
            Description string  `json:"description"`
            Schedule map[string]interface{}  `json:"_schedule"`
            Type string  `json:"type"`
         
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
    
    // CollectionBanner used by Catalog
    type CollectionBanner struct {

        
            Landscape CollectionImage  `json:"landscape"`
            Portrait CollectionImage  `json:"portrait"`
         
    }
    
    // CollectionBadge used by Catalog
    type CollectionBadge struct {

        
            Text string  `json:"text"`
            Color string  `json:"color"`
         
    }
    
    // UserInfo used by Catalog
    type UserInfo struct {

        
            Email string  `json:"email"`
            Username string  `json:"username"`
            UID string  `json:"uid"`
            UserID string  `json:"user_id"`
         
    }
    
    // SeoDetail used by Catalog
    type SeoDetail struct {

        
            Description string  `json:"description"`
            Title string  `json:"title"`
         
    }
    
    // NextSchedule used by Catalog
    type NextSchedule struct {

        
            Start string  `json:"start"`
            End string  `json:"end"`
         
    }
    
    // CollectionSchedule used by Catalog
    type CollectionSchedule struct {

        
            Cron string  `json:"cron"`
            End string  `json:"end"`
            NextSchedule []NextSchedule  `json:"next_schedule"`
            Start string  `json:"start"`
            Duration float64  `json:"duration"`
         
    }
    
    // CreateCollection used by Catalog
    type CreateCollection struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            Banners CollectionBanner  `json:"banners"`
            AppID string  `json:"app_id"`
            IsActive bool  `json:"is_active"`
            AllowFacets bool  `json:"allow_facets"`
            IsVisible bool  `json:"is_visible"`
            Query []CollectionQuery  `json:"query"`
            Badge CollectionBadge  `json:"badge"`
            CreatedBy UserInfo  `json:"created_by"`
            Logo CollectionImage  `json:"logo"`
            Meta map[string]interface{}  `json:"meta"`
            AllowSort bool  `json:"allow_sort"`
            Published bool  `json:"published"`
            Name string  `json:"name"`
            SortOn string  `json:"sort_on"`
            ModifiedBy UserInfo  `json:"modified_by"`
            Priority float64  `json:"priority"`
            Slug string  `json:"slug"`
            Seo SeoDetail  `json:"seo"`
            Description string  `json:"description"`
            Tags []string  `json:"tags"`
            Schedule CollectionSchedule  `json:"_schedule"`
            Type string  `json:"type"`
         
    }
    
    // CollectionCreateResponse used by Catalog
    type CollectionCreateResponse struct {

        
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Banners ImageUrls  `json:"banners"`
            AppID string  `json:"app_id"`
            IsActive bool  `json:"is_active"`
            AllowFacets bool  `json:"allow_facets"`
            Cron map[string]interface{}  `json:"cron"`
            Query []CollectionQuery  `json:"query"`
            Badge map[string]interface{}  `json:"badge"`
            Logo BannerImage  `json:"logo"`
            Meta map[string]interface{}  `json:"meta"`
            AllowSort bool  `json:"allow_sort"`
            Tag []string  `json:"tag"`
            Name string  `json:"name"`
            SortOn string  `json:"sort_on"`
            Priority float64  `json:"priority"`
            Slug string  `json:"slug"`
            Description string  `json:"description"`
            Schedule map[string]interface{}  `json:"_schedule"`
            Type string  `json:"type"`
         
    }
    
    // CollectionDetailResponse used by Catalog
    type CollectionDetailResponse struct {

        
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Banners ImageUrls  `json:"banners"`
            AppID string  `json:"app_id"`
            IsActive bool  `json:"is_active"`
            AllowFacets bool  `json:"allow_facets"`
            Cron map[string]interface{}  `json:"cron"`
            UID string  `json:"uid"`
            Query []CollectionQuery  `json:"query"`
            Badge map[string]interface{}  `json:"badge"`
            Logo Media1  `json:"logo"`
            Meta map[string]interface{}  `json:"meta"`
            AllowSort bool  `json:"allow_sort"`
            Tag []string  `json:"tag"`
            Name string  `json:"name"`
            Priority float64  `json:"priority"`
            Slug string  `json:"slug"`
            Description string  `json:"description"`
            Schedule map[string]interface{}  `json:"_schedule"`
            Type string  `json:"type"`
         
    }
    
    // UpdateCollection used by Catalog
    type UpdateCollection struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            Banners CollectionBanner  `json:"banners"`
            IsActive bool  `json:"is_active"`
            AllowFacets bool  `json:"allow_facets"`
            IsVisible bool  `json:"is_visible"`
            Query []CollectionQuery  `json:"query"`
            Badge CollectionBadge  `json:"badge"`
            Logo CollectionImage  `json:"logo"`
            Meta map[string]interface{}  `json:"meta"`
            AllowSort bool  `json:"allow_sort"`
            Name string  `json:"name"`
            SortOn string  `json:"sort_on"`
            ModifiedBy UserInfo  `json:"modified_by"`
            Priority float64  `json:"priority"`
            Slug string  `json:"slug"`
            Seo SeoDetail  `json:"seo"`
            Schedule CollectionSchedule  `json:"_schedule"`
            Description string  `json:"description"`
            Tags []string  `json:"tags"`
            Published bool  `json:"published"`
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

        
            CurrencyCode string  `json:"currency_code"`
            Min float64  `json:"min"`
            CurrencySymbol string  `json:"currency_symbol"`
            Max float64  `json:"max"`
         
    }
    
    // ProductListingPrice used by Catalog
    type ProductListingPrice struct {

        
            Marked Price1  `json:"marked"`
            Effective Price1  `json:"effective"`
         
    }
    
    // ProductBrand used by Catalog
    type ProductBrand struct {

        
            Name string  `json:"name"`
            Action Action  `json:"action"`
            UID float64  `json:"uid"`
            Logo Media1  `json:"logo"`
         
    }
    
    // ProductListingDetail used by Catalog
    type ProductListingDetail struct {

        
            Discount string  `json:"discount"`
            HasVariant bool  `json:"has_variant"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            RatingCount float64  `json:"rating_count"`
            Tryouts []string  `json:"tryouts"`
            UID float64  `json:"uid"`
            ItemType string  `json:"item_type"`
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
            ShortDescription string  `json:"short_description"`
            Color string  `json:"color"`
            Name string  `json:"name"`
            ImageNature string  `json:"image_nature"`
            ItemCode string  `json:"item_code"`
            Sellable bool  `json:"sellable"`
            Rating float64  `json:"rating"`
            Price ProductListingPrice  `json:"price"`
            ProductOnlineDate string  `json:"product_online_date"`
            Similars []string  `json:"similars"`
            Highlights []string  `json:"highlights"`
            Slug string  `json:"slug"`
            Medias []Media1  `json:"medias"`
            Description string  `json:"description"`
            Attributes map[string]interface{}  `json:"attributes"`
            PromoMeta map[string]interface{}  `json:"promo_meta"`
            Type string  `json:"type"`
            Brand ProductBrand  `json:"brand"`
         
    }
    
    // GetCollectionItemsResponse used by Catalog
    type GetCollectionItemsResponse struct {

        
            Page Page  `json:"page"`
            Items []ProductListingDetail  `json:"items"`
            SortOn []ProductSortOn  `json:"sort_on"`
            Filters []ProductFilters  `json:"filters"`
         
    }
    
    // CollectionItem used by Catalog
    type CollectionItem struct {

        
            Action string  `json:"action"`
            ItemID float64  `json:"item_id"`
            Priority float64  `json:"priority"`
         
    }
    
    // CollectionItemUpdate used by Catalog
    type CollectionItemUpdate struct {

        
            AllowFacets bool  `json:"allow_facets"`
            Items []CollectionItem  `json:"items"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Query []CollectionQuery  `json:"query"`
            AllowSort bool  `json:"allow_sort"`
            Type string  `json:"type"`
         
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

        
            ArticleFreshness float64  `json:"article_freshness"`
            TotalArticles float64  `json:"total_articles"`
            AvailableArticles float64  `json:"available_articles"`
            TotalSizes float64  `json:"total_sizes"`
            Name string  `json:"name"`
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

        
            Data CrossSellingData  `json:"data"`
            BrandDistribution CatalogInsightBrand  `json:"brand_distribution"`
         
    }
    
    // OptInPostRequest used by Catalog
    type OptInPostRequest struct {

        
            CompanyID float64  `json:"company_id"`
            StoreIds []float64  `json:"store_ids"`
            Platform string  `json:"platform"`
            BrandIds []float64  `json:"brand_ids"`
            Enabled bool  `json:"enabled"`
            OptLevel string  `json:"opt_level"`
         
    }
    
    // CompanyOptIn used by Catalog
    type CompanyOptIn struct {

        
            CompanyID float64  `json:"company_id"`
            CreatedOn float64  `json:"created_on"`
            ModifiedOn float64  `json:"modified_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            StoreIds []float64  `json:"store_ids"`
            Platform string  `json:"platform"`
            BrandIds []float64  `json:"brand_ids"`
            Enabled bool  `json:"enabled"`
            OptLevel string  `json:"opt_level"`
         
    }
    
    // GetOptInPlatform used by Catalog
    type GetOptInPlatform struct {

        
            Page Page  `json:"page"`
            Items []CompanyOptIn  `json:"items"`
         
    }
    
    // OptinCompanyDetail used by Catalog
    type OptinCompanyDetail struct {

        
            UID float64  `json:"uid"`
            BusinessType string  `json:"business_type"`
            Name string  `json:"name"`
            CompanyType string  `json:"company_type"`
         
    }
    
    // CompanyBrandDetail used by Catalog
    type CompanyBrandDetail struct {

        
            CompanyID float64  `json:"company_id"`
            BrandName string  `json:"brand_name"`
            TotalArticle float64  `json:"total_article"`
            BrandID float64  `json:"brand_id"`
         
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
            Manager map[string]interface{}  `json:"manager"`
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
            StoreType string  `json:"store_type"`
            Address map[string]interface{}  `json:"address"`
            AdditionalContacts []map[string]interface{}  `json:"additional_contacts"`
            UID float64  `json:"uid"`
            StoreCode string  `json:"store_code"`
            Timing map[string]interface{}  `json:"timing"`
            Name string  `json:"name"`
            Documents []map[string]interface{}  `json:"documents"`
            DisplayName string  `json:"display_name"`
         
    }
    
    // OptinStoreDetails used by Catalog
    type OptinStoreDetails struct {

        
            Page Page  `json:"page"`
            Items []StoreDetail  `json:"items"`
         
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

        
            Multi bool  `json:"multi"`
            AllowedValues []string  `json:"allowed_values"`
            Mandatory bool  `json:"mandatory"`
            Range AttributeSchemaRange  `json:"range"`
            Format string  `json:"format"`
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
    
    // AttributeMasterDetails used by Catalog
    type AttributeMasterDetails struct {

        
            DisplayType string  `json:"display_type"`
         
    }
    
    // GenderDetail used by Catalog
    type GenderDetail struct {

        
            ID string  `json:"id"`
            Filters AttributeMasterFilter  `json:"filters"`
            Slug string  `json:"slug"`
            IsNested bool  `json:"is_nested"`
            Departments []string  `json:"departments"`
            Description string  `json:"description"`
            EnabledForEndConsumer bool  `json:"enabled_for_end_consumer"`
            Schema AttributeMaster  `json:"schema"`
            Meta AttributeMasterMeta  `json:"meta"`
            Logo string  `json:"logo"`
            Name string  `json:"name"`
            Details AttributeMasterDetails  `json:"details"`
         
    }
    
    // CategoriesResponse used by Catalog
    type CategoriesResponse struct {

        
            Slug string  `json:"slug"`
            TemplateSlug string  `json:"template_slug"`
            SlugKey string  `json:"slug_key"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // ProdcutTemplateCategoriesResponse used by Catalog
    type ProdcutTemplateCategoriesResponse struct {

        
            Page Page  `json:"page"`
            Items []CategoriesResponse  `json:"items"`
         
    }
    
    // PTErrorResponse used by Catalog
    type PTErrorResponse struct {

        
            Message string  `json:"message"`
            Errors map[string]interface{}  `json:"errors"`
            Status float64  `json:"status"`
            Code string  `json:"code"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // UserSerializer used by Catalog
    type UserSerializer struct {

        
            ID string  `json:"_id"`
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
            UID string  `json:"uid"`
            Contact string  `json:"contact"`
         
    }
    
    // GetDepartment used by Catalog
    type GetDepartment struct {

        
            PageNo float64  `json:"page_no"`
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
            PageSize float64  `json:"page_size"`
            CreatedBy UserSerializer  `json:"created_by"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
            Synonyms []string  `json:"synonyms"`
            Search string  `json:"search"`
            Logo string  `json:"logo"`
            Name string  `json:"name"`
            IsActive bool  `json:"is_active"`
            ItemType string  `json:"item_type"`
            PriorityOrder float64  `json:"priority_order"`
         
    }
    
    // DepartmentsResponse used by Catalog
    type DepartmentsResponse struct {

        
            Page Page  `json:"page"`
            Items []GetDepartment  `json:"items"`
         
    }
    
    // DepartmentErrorResponse used by Catalog
    type DepartmentErrorResponse struct {

        
            Message string  `json:"message"`
            Errors map[string]interface{}  `json:"errors"`
            Status float64  `json:"status"`
            Code string  `json:"code"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // DepartmentCreateUpdate used by Catalog
    type DepartmentCreateUpdate struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
            Synonyms []string  `json:"synonyms"`
            Tags []string  `json:"tags"`
            Logo string  `json:"logo"`
            Cls string  `json:"_cls"`
            Platforms map[string]interface{}  `json:"platforms"`
            Name string  `json:"name"`
            IsActive bool  `json:"is_active"`
            PriorityOrder float64  `json:"priority_order"`
         
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
            Contact string  `json:"contact"`
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
         
    }
    
    // DepartmentModel used by Catalog
    type DepartmentModel struct {

        
            ID interface{}  `json:"_id"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
            ModifiedBy UserDetail  `json:"modified_by"`
            CreatedBy UserDetail  `json:"created_by"`
            Slug interface{}  `json:"slug"`
            UID float64  `json:"uid"`
            VerifiedOn string  `json:"verified_on"`
            VerifiedBy UserDetail  `json:"verified_by"`
            Logo string  `json:"logo"`
            Synonyms []interface{}  `json:"synonyms"`
            Cls interface{}  `json:"_cls"`
            Name interface{}  `json:"name"`
            IsActive bool  `json:"is_active"`
            PriorityOrder float64  `json:"priority_order"`
         
    }
    
    // ProductTemplate used by Catalog
    type ProductTemplate struct {

        
            IsArchived bool  `json:"is_archived"`
            IsExpirable bool  `json:"is_expirable"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Departments []string  `json:"departments"`
            Slug string  `json:"slug"`
            IsPhysical bool  `json:"is_physical"`
            Description string  `json:"description"`
            Attributes []string  `json:"attributes"`
            Logo string  `json:"logo"`
            Tag string  `json:"tag"`
            Name string  `json:"name"`
            IsActive bool  `json:"is_active"`
            Categories []string  `json:"categories"`
         
    }
    
    // TemplatesResponse used by Catalog
    type TemplatesResponse struct {

        
            Page Page  `json:"page"`
            Items ProductTemplate  `json:"items"`
         
    }
    
    // TemplateDetails used by Catalog
    type TemplateDetails struct {

        
            IsArchived bool  `json:"is_archived"`
            IsExpirable bool  `json:"is_expirable"`
            ID string  `json:"id"`
            Departments []string  `json:"departments"`
            Slug string  `json:"slug"`
            IsPhysical bool  `json:"is_physical"`
            Description string  `json:"description"`
            Attributes []string  `json:"attributes"`
            Logo string  `json:"logo"`
            Tag string  `json:"tag"`
            Name string  `json:"name"`
            IsActive bool  `json:"is_active"`
            Categories []string  `json:"categories"`
         
    }
    
    // Properties used by Catalog
    type Properties struct {

        
            Media map[string]interface{}  `json:"media"`
            ReturnConfig map[string]interface{}  `json:"return_config"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            IsActive map[string]interface{}  `json:"is_active"`
            CountryOfOrigin map[string]interface{}  `json:"country_of_origin"`
            IsDependent map[string]interface{}  `json:"is_dependent"`
            CategorySlug map[string]interface{}  `json:"category_slug"`
            ProductPublish map[string]interface{}  `json:"product_publish"`
            Currency map[string]interface{}  `json:"currency"`
            HsnCode map[string]interface{}  `json:"hsn_code"`
            ItemType map[string]interface{}  `json:"item_type"`
            Variants map[string]interface{}  `json:"variants"`
            TraderType map[string]interface{}  `json:"trader_type"`
            BrandUID map[string]interface{}  `json:"brand_uid"`
            Trader map[string]interface{}  `json:"trader"`
            ShortDescription map[string]interface{}  `json:"short_description"`
            MultiSize map[string]interface{}  `json:"multi_size"`
            Name map[string]interface{}  `json:"name"`
            ProductGroupTag map[string]interface{}  `json:"product_group_tag"`
            ItemCode map[string]interface{}  `json:"item_code"`
            Command map[string]interface{}  `json:"command"`
            SizeGuide map[string]interface{}  `json:"size_guide"`
            NoOfBoxes map[string]interface{}  `json:"no_of_boxes"`
            Sizes map[string]interface{}  `json:"sizes"`
            Highlights map[string]interface{}  `json:"highlights"`
            Slug map[string]interface{}  `json:"slug"`
            Description map[string]interface{}  `json:"description"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            Tags map[string]interface{}  `json:"tags"`
         
    }
    
    // GlobalValidation used by Catalog
    type GlobalValidation struct {

        
            Definitions map[string]interface{}  `json:"definitions"`
            Description string  `json:"description"`
            Properties Properties  `json:"properties"`
            Title string  `json:"title"`
            Required []string  `json:"required"`
            Type string  `json:"type"`
         
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
            Username string  `json:"username"`
            UID string  `json:"uid"`
            UserID string  `json:"user_id"`
         
    }
    
    // ProductTemplateExportResponse used by Catalog
    type ProductTemplateExportResponse struct {

        
            CompletedOn string  `json:"completed_on"`
            ModifiedOn string  `json:"modified_on"`
            CreatedBy UserInfo1  `json:"created_by"`
            Filters map[string]interface{}  `json:"filters"`
            SellerID float64  `json:"seller_id"`
            NotificationEmails []string  `json:"notification_emails"`
            Status string  `json:"status"`
            URL string  `json:"url"`
            TaskID string  `json:"task_id"`
            Type string  `json:"type"`
         
    }
    
    // ProductDownloadsResponse used by Catalog
    type ProductDownloadsResponse struct {

        
            Items []ProductTemplateExportResponse  `json:"items"`
         
    }
    
    // ProductTemplateExportFilterRequest used by Catalog
    type ProductTemplateExportFilterRequest struct {

        
            Brands []string  `json:"brands"`
            FromDate string  `json:"from_date"`
            ToDate string  `json:"to_date"`
            CatalogueTypes []string  `json:"catalogue_types"`
            Templates []string  `json:"templates"`
         
    }
    
    // ProductTemplateDownloadsExport used by Catalog
    type ProductTemplateDownloadsExport struct {

        
            NotificationEmails []string  `json:"notification_emails"`
            Type string  `json:"type"`
            Filters ProductTemplateExportFilterRequest  `json:"filters"`
         
    }
    
    // ProductConfigurationDownloads used by Catalog
    type ProductConfigurationDownloads struct {

        
            Multivalue bool  `json:"multivalue"`
            Data []map[string]interface{}  `json:"data"`
         
    }
    
    // Media2 used by Catalog
    type Media2 struct {

        
            Logo string  `json:"logo"`
            Landscape string  `json:"landscape"`
            Portrait string  `json:"portrait"`
         
    }
    
    // Hierarchy used by Catalog
    type Hierarchy struct {

        
            L1 float64  `json:"l1"`
            L2 float64  `json:"l2"`
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
    
    // Category used by Catalog
    type Category struct {

        
            ID string  `json:"id"`
            CreatedOn string  `json:"created_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Priority float64  `json:"priority"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Departments []float64  `json:"departments"`
            Slug string  `json:"slug"`
            Media Media2  `json:"media"`
            Hierarchy []Hierarchy  `json:"hierarchy"`
            Marketplaces CategoryMapping  `json:"marketplaces"`
            Tryouts []string  `json:"tryouts"`
            ModifiedOn string  `json:"modified_on"`
            UID float64  `json:"uid"`
            Synonyms []string  `json:"synonyms"`
            Name string  `json:"name"`
            Level float64  `json:"level"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // CategoryResponse used by Catalog
    type CategoryResponse struct {

        
            Page Page  `json:"page"`
            Items []Category  `json:"items"`
         
    }
    
    // CategoryRequestBody used by Catalog
    type CategoryRequestBody struct {

        
            Priority float64  `json:"priority"`
            Departments []float64  `json:"departments"`
            Slug string  `json:"slug"`
            Media Media2  `json:"media"`
            Hierarchy []Hierarchy  `json:"hierarchy"`
            Marketplaces CategoryMapping  `json:"marketplaces"`
            Tryouts []string  `json:"tryouts"`
            Synonyms []string  `json:"synonyms"`
            Name string  `json:"name"`
            Level float64  `json:"level"`
            IsActive bool  `json:"is_active"`
         
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

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // Image used by Catalog
    type Image struct {

        
            AspectRatioF float64  `json:"aspect_ratio_f"`
            SecureURL string  `json:"secure_url"`
            URL string  `json:"url"`
            AspectRatio string  `json:"aspect_ratio"`
         
    }
    
    // ReturnConfigResponse used by Catalog
    type ReturnConfigResponse struct {

        
            Time float64  `json:"time"`
            Returnable bool  `json:"returnable"`
            Unit string  `json:"unit"`
         
    }
    
    // NetQuantityResponse used by Catalog
    type NetQuantityResponse struct {

        
            Value float64  `json:"value"`
            Unit string  `json:"unit"`
         
    }
    
    // ProductPublish used by Catalog
    type ProductPublish struct {

        
            ProductOnlineDate string  `json:"product_online_date"`
            IsSet bool  `json:"is_set"`
         
    }
    
    // Trader used by Catalog
    type Trader struct {

        
            Name interface{}  `json:"name"`
            Type string  `json:"type"`
            Address []string  `json:"address"`
         
    }
    
    // VerifiedBy used by Catalog
    type VerifiedBy struct {

        
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
         
    }
    
    // Logo used by Catalog
    type Logo struct {

        
            AspectRatioF float64  `json:"aspect_ratio_f"`
            SecureURL string  `json:"secure_url"`
            URL string  `json:"url"`
            AspectRatio string  `json:"aspect_ratio"`
         
    }
    
    // Brand used by Catalog
    type Brand struct {

        
            Name string  `json:"name"`
            UID float64  `json:"uid"`
            Logo Logo  `json:"logo"`
         
    }
    
    // ProductSchemaV2 used by Catalog
    type ProductSchemaV2 struct {

        
            CompanyID float64  `json:"company_id"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Images []Image  `json:"images"`
            Media []Media1  `json:"media"`
            Pending string  `json:"pending"`
            IsSet bool  `json:"is_set"`
            ReturnConfig ReturnConfigResponse  `json:"return_config"`
            VerifiedOn string  `json:"verified_on"`
            NetQuantity NetQuantityResponse  `json:"net_quantity"`
            IsActive bool  `json:"is_active"`
            IsExpirable bool  `json:"is_expirable"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            CountryOfOrigin string  `json:"country_of_origin"`
            IsDependent bool  `json:"is_dependent"`
            ModifiedOn string  `json:"modified_on"`
            CategorySlug string  `json:"category_slug"`
            PrimaryColor string  `json:"primary_color"`
            CreatedOn string  `json:"created_on"`
            AllCompanyIds []float64  `json:"all_company_ids"`
            UID float64  `json:"uid"`
            ProductPublish ProductPublish  `json:"product_publish"`
            Currency string  `json:"currency"`
            Category map[string]interface{}  `json:"category"`
            L3Mapping []string  `json:"l3_mapping"`
            Moq map[string]interface{}  `json:"moq"`
            ItemType string  `json:"item_type"`
            HsnCode string  `json:"hsn_code"`
            Variants map[string]interface{}  `json:"variants"`
            ID string  `json:"id"`
            AllIdentifiers []string  `json:"all_identifiers"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            BrandUID float64  `json:"brand_uid"`
            Trader []Trader  `json:"trader"`
            VerifiedBy VerifiedBy  `json:"verified_by"`
            ShortDescription string  `json:"short_description"`
            MultiSize bool  `json:"multi_size"`
            Color string  `json:"color"`
            Name string  `json:"name"`
            ProductGroupTag []string  `json:"product_group_tag"`
            ImageNature string  `json:"image_nature"`
            ItemCode string  `json:"item_code"`
            SizeGuide string  `json:"size_guide"`
            NoOfBoxes float64  `json:"no_of_boxes"`
            Sizes []map[string]interface{}  `json:"sizes"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Highlights []string  `json:"highlights"`
            Departments []float64  `json:"departments"`
            Slug string  `json:"slug"`
            VariantGroup map[string]interface{}  `json:"variant_group"`
            CategoryUID float64  `json:"category_uid"`
            IsImageLessProduct bool  `json:"is_image_less_product"`
            IsPhysical bool  `json:"is_physical"`
            Description string  `json:"description"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            AllSizes []map[string]interface{}  `json:"all_sizes"`
            Tags []string  `json:"tags"`
            VariantMedia map[string]interface{}  `json:"variant_media"`
            Stage string  `json:"stage"`
            Attributes map[string]interface{}  `json:"attributes"`
            TemplateTag string  `json:"template_tag"`
            Brand Brand  `json:"brand"`
         
    }
    
    // ProductListingResponseV2 used by Catalog
    type ProductListingResponseV2 struct {

        
            Page Page  `json:"page"`
            Items []ProductSchemaV2  `json:"items"`
         
    }
    
    // ReturnConfig used by Catalog
    type ReturnConfig struct {

        
            Time float64  `json:"time"`
            Returnable bool  `json:"returnable"`
            Unit string  `json:"unit"`
         
    }
    
    // NetQuantity used by Catalog
    type NetQuantity struct {

        
            Value float64  `json:"value"`
            Unit interface{}  `json:"unit"`
         
    }
    
    // TeaserTag used by Catalog
    type TeaserTag struct {

        
            Tag string  `json:"tag"`
            URL string  `json:"url"`
         
    }
    
    // TaxIdentifier used by Catalog
    type TaxIdentifier struct {

        
            ReportingHsn string  `json:"reporting_hsn"`
            HsnCode string  `json:"hsn_code"`
            HsnCodeID string  `json:"hsn_code_id"`
         
    }
    
    // ProductPublish1 used by Catalog
    type ProductPublish1 struct {

        
            ProductOnlineDate string  `json:"product_online_date"`
            IsSet bool  `json:"is_set"`
         
    }
    
    // CustomOrder used by Catalog
    type CustomOrder struct {

        
            ManufacturingTimeUnit string  `json:"manufacturing_time_unit"`
            IsCustomOrder bool  `json:"is_custom_order"`
            ManufacturingTime float64  `json:"manufacturing_time"`
         
    }
    
    // ProductCreateUpdateSchemaV2 used by Catalog
    type ProductCreateUpdateSchemaV2 struct {

        
            CompanyID float64  `json:"company_id"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Media []Media1  `json:"media"`
            Requester string  `json:"requester"`
            IsSet bool  `json:"is_set"`
            ReturnConfig ReturnConfig  `json:"return_config"`
            Action string  `json:"action"`
            NetQuantity NetQuantity  `json:"net_quantity"`
            TeaserTag TeaserTag  `json:"teaser_tag"`
            IsActive bool  `json:"is_active"`
            TaxIdentifier TaxIdentifier  `json:"tax_identifier"`
            CountryOfOrigin string  `json:"country_of_origin"`
            IsDependent bool  `json:"is_dependent"`
            CategorySlug string  `json:"category_slug"`
            UID float64  `json:"uid"`
            ProductPublish ProductPublish1  `json:"product_publish"`
            Currency string  `json:"currency"`
            ItemType string  `json:"item_type"`
            Variants map[string]interface{}  `json:"variants"`
            BrandUID float64  `json:"brand_uid"`
            Trader []Trader  `json:"trader"`
            ShortDescription string  `json:"short_description"`
            MultiSize bool  `json:"multi_size"`
            Name string  `json:"name"`
            ProductGroupTag []string  `json:"product_group_tag"`
            ItemCode string  `json:"item_code"`
            SizeGuide string  `json:"size_guide"`
            NoOfBoxes float64  `json:"no_of_boxes"`
            ChangeRequestID interface{}  `json:"change_request_id"`
            Sizes []map[string]interface{}  `json:"sizes"`
            Highlights []string  `json:"highlights"`
            VariantGroup map[string]interface{}  `json:"variant_group"`
            Slug string  `json:"slug"`
            Departments []float64  `json:"departments"`
            BulkJobID string  `json:"bulk_job_id"`
            IsImageLessProduct bool  `json:"is_image_less_product"`
            Description string  `json:"description"`
            CustomOrder CustomOrder  `json:"custom_order"`
            Tags []string  `json:"tags"`
            VariantMedia map[string]interface{}  `json:"variant_media"`
            Attributes map[string]interface{}  `json:"attributes"`
            TemplateTag string  `json:"template_tag"`
         
    }
    
    // ProductVariants used by Catalog
    type ProductVariants struct {

        
            Media []Media1  `json:"media"`
            BrandUID float64  `json:"brand_uid"`
            CategoryUID float64  `json:"category_uid"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            ItemCode string  `json:"item_code"`
         
    }
    
    // ProductVariantsResponse used by Catalog
    type ProductVariantsResponse struct {

        
            Page Page  `json:"page"`
            Variants []ProductVariants  `json:"variants"`
         
    }
    
    // AttributeMasterSerializer used by Catalog
    type AttributeMasterSerializer struct {

        
            Variant bool  `json:"variant"`
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
            Suggestion string  `json:"suggestion"`
            IsNested bool  `json:"is_nested"`
            EnabledForEndConsumer bool  `json:"enabled_for_end_consumer"`
            Unit string  `json:"unit"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Logo string  `json:"logo"`
            Name string  `json:"name"`
            Details AttributeMasterDetails  `json:"details"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            RawKey string  `json:"raw_key"`
            Departments []string  `json:"departments"`
            Slug string  `json:"slug"`
            Filters AttributeMasterFilter  `json:"filters"`
            Description string  `json:"description"`
            Tags []string  `json:"tags"`
            Synonyms map[string]interface{}  `json:"synonyms"`
            Schema AttributeMaster  `json:"schema"`
         
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
            GtinType string  `json:"gtin_type"`
            Primary bool  `json:"primary"`
         
    }
    
    // AllSizes used by Catalog
    type AllSizes struct {

        
            ItemWeight float64  `json:"item_weight"`
            ItemWidth float64  `json:"item_width"`
            ItemDimensionsUnitOfMeasure string  `json:"item_dimensions_unit_of_measure"`
            ItemLength float64  `json:"item_length"`
            ItemWeightUnitOfMeasure interface{}  `json:"item_weight_unit_of_measure"`
            Size interface{}  `json:"size"`
            ItemHeight float64  `json:"item_height"`
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

        
            CompanyID float64  `json:"company_id"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Images []Image  `json:"images"`
            Media []Media1  `json:"media"`
            Pending string  `json:"pending"`
            IsSet bool  `json:"is_set"`
            ReturnConfig ReturnConfigResponse  `json:"return_config"`
            VerifiedOn string  `json:"verified_on"`
            NetQuantity NetQuantityResponse  `json:"net_quantity"`
            IsActive bool  `json:"is_active"`
            IsExpirable bool  `json:"is_expirable"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            CountryOfOrigin string  `json:"country_of_origin"`
            IsDependent bool  `json:"is_dependent"`
            ModifiedOn string  `json:"modified_on"`
            CategorySlug string  `json:"category_slug"`
            PrimaryColor string  `json:"primary_color"`
            CreatedOn string  `json:"created_on"`
            AllCompanyIds []float64  `json:"all_company_ids"`
            UID float64  `json:"uid"`
            ProductPublish ProductPublished  `json:"product_publish"`
            Currency string  `json:"currency"`
            Category map[string]interface{}  `json:"category"`
            L3Mapping []string  `json:"l3_mapping"`
            Moq map[string]interface{}  `json:"moq"`
            ItemType string  `json:"item_type"`
            HsnCode string  `json:"hsn_code"`
            Variants map[string]interface{}  `json:"variants"`
            ID string  `json:"id"`
            AllIdentifiers []string  `json:"all_identifiers"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            BrandUID float64  `json:"brand_uid"`
            Trader []Trader  `json:"trader"`
            VerifiedBy VerifiedBy  `json:"verified_by"`
            ShortDescription string  `json:"short_description"`
            MultiSize bool  `json:"multi_size"`
            Color string  `json:"color"`
            Name string  `json:"name"`
            ProductGroupTag []string  `json:"product_group_tag"`
            ImageNature string  `json:"image_nature"`
            ItemCode string  `json:"item_code"`
            SizeGuide string  `json:"size_guide"`
            NoOfBoxes float64  `json:"no_of_boxes"`
            Sizes []map[string]interface{}  `json:"sizes"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Highlights []string  `json:"highlights"`
            Departments []float64  `json:"departments"`
            Slug string  `json:"slug"`
            VariantGroup map[string]interface{}  `json:"variant_group"`
            CategoryUID float64  `json:"category_uid"`
            IsImageLessProduct bool  `json:"is_image_less_product"`
            IsPhysical bool  `json:"is_physical"`
            Description string  `json:"description"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            AllSizes []map[string]interface{}  `json:"all_sizes"`
            Tags []string  `json:"tags"`
            VariantMedia map[string]interface{}  `json:"variant_media"`
            Stage string  `json:"stage"`
            Attributes map[string]interface{}  `json:"attributes"`
            TemplateTag string  `json:"template_tag"`
            Brand Brand  `json:"brand"`
         
    }
    
    // ProductListingResponse used by Catalog
    type ProductListingResponse struct {

        
            Page Page  `json:"page"`
            Items []Product  `json:"items"`
         
    }
    
    // UserDetail1 used by Catalog
    type UserDetail1 struct {

        
            FullName string  `json:"full_name"`
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
         
    }
    
    // ProductBulkRequest used by Catalog
    type ProductBulkRequest struct {

        
            CompanyID float64  `json:"company_id"`
            IsActive bool  `json:"is_active"`
            CreatedOn string  `json:"created_on"`
            CancelledRecords []string  `json:"cancelled_records"`
            ModifiedOn string  `json:"modified_on"`
            CreatedBy UserDetail1  `json:"created_by"`
            ModifiedBy UserDetail1  `json:"modified_by"`
            Template ProductTemplate  `json:"template"`
            FilePath string  `json:"file_path"`
            Succeed float64  `json:"succeed"`
            Stage string  `json:"stage"`
            FailedRecords []string  `json:"failed_records"`
            TemplateTag string  `json:"template_tag"`
            Failed float64  `json:"failed"`
            Cancelled float64  `json:"cancelled"`
            Total float64  `json:"total"`
         
    }
    
    // ProductBulkRequestList used by Catalog
    type ProductBulkRequestList struct {

        
            Page Page  `json:"page"`
            Items ProductBulkRequest  `json:"items"`
         
    }
    
    // BulkJob used by Catalog
    type BulkJob struct {

        
            CompanyID float64  `json:"company_id"`
            CustomTemplateTag string  `json:"custom_template_tag"`
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
            ModifiedBy UserInfo1  `json:"modified_by"`
            CreatedBy UserInfo1  `json:"created_by"`
            CancelledRecords []map[string]interface{}  `json:"cancelled_records"`
            TemplateTag string  `json:"template_tag"`
            Failed float64  `json:"failed"`
            FilePath string  `json:"file_path"`
            Succeed float64  `json:"succeed"`
            Stage string  `json:"stage"`
            TrackingURL string  `json:"tracking_url"`
            FailedRecords []map[string]interface{}  `json:"failed_records"`
            IsActive bool  `json:"is_active"`
            Cancelled float64  `json:"cancelled"`
            Total float64  `json:"total"`
         
    }
    
    // BulkResponse used by Catalog
    type BulkResponse struct {

        
            BatchID string  `json:"batch_id"`
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
            ModifiedBy UserInfo1  `json:"modified_by"`
            CreatedBy UserInfo1  `json:"created_by"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // BulkProductRequest used by Catalog
    type BulkProductRequest struct {

        
            CompanyID float64  `json:"company_id"`
            TemplateTag string  `json:"template_tag"`
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
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
         
    }
    
    // Items used by Catalog
    type Items struct {

        
            IsActive bool  `json:"is_active"`
            CompanyID float64  `json:"company_id"`
            ID string  `json:"id"`
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
            ModifiedBy UserCommon  `json:"modified_by"`
            CreatedBy UserCommon  `json:"created_by"`
            CancelledRecords []string  `json:"cancelled_records"`
            Retry float64  `json:"retry"`
            FilePath string  `json:"file_path"`
            Succeed float64  `json:"succeed"`
            Stage string  `json:"stage"`
            TrackingURL string  `json:"tracking_url"`
            FailedRecords []string  `json:"failed_records"`
            Failed float64  `json:"failed"`
            Cancelled float64  `json:"cancelled"`
            Total float64  `json:"total"`
         
    }
    
    // BulkAssetResponse used by Catalog
    type BulkAssetResponse struct {

        
            Page Page  `json:"page"`
            Items []Items  `json:"items"`
         
    }
    
    // ProductBulkAssets used by Catalog
    type ProductBulkAssets struct {

        
            URL string  `json:"url"`
            User map[string]interface{}  `json:"user"`
            CompanyID float64  `json:"company_id"`
         
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
            PriceEffective float64  `json:"price_effective"`
            SellableQuantity float64  `json:"sellable_quantity"`
            Size string  `json:"size"`
            UID string  `json:"uid"`
            Store map[string]interface{}  `json:"store"`
            Quantity float64  `json:"quantity"`
            ItemID float64  `json:"item_id"`
            SellerIdentifier string  `json:"seller_identifier"`
            Currency string  `json:"currency"`
            Price float64  `json:"price"`
            Identifiers map[string]interface{}  `json:"identifiers"`
         
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

        
            Name string  `json:"name"`
            Quantity float64  `json:"quantity"`
            SizeDistribution SizeDistribution  `json:"size_distribution"`
         
    }
    
    // GTIN used by Catalog
    type GTIN struct {

        
            GtinValue interface{}  `json:"gtin_value"`
            GtinType string  `json:"gtin_type"`
            Primary bool  `json:"primary"`
         
    }
    
    // InvSize used by Catalog
    type InvSize struct {

        
            ItemWeight float64  `json:"item_weight"`
            Set InventorySet  `json:"set"`
            ItemWidth float64  `json:"item_width"`
            PriceTransfer float64  `json:"price_transfer"`
            PriceEffective float64  `json:"price_effective"`
            ItemDimensionsUnitOfMeasure string  `json:"item_dimensions_unit_of_measure"`
            ItemLength float64  `json:"item_length"`
            ItemWeightUnitOfMeasure string  `json:"item_weight_unit_of_measure"`
            ExpirationDate string  `json:"expiration_date"`
            Size interface{}  `json:"size"`
            IsSet bool  `json:"is_set"`
            Quantity float64  `json:"quantity"`
            ItemHeight float64  `json:"item_height"`
            StoreCode string  `json:"store_code"`
            Currency string  `json:"currency"`
            Price float64  `json:"price"`
            Identifiers []GTIN  `json:"identifiers"`
         
    }
    
    // ItemQuery used by Catalog
    type ItemQuery struct {

        
            UID float64  `json:"uid"`
            ItemCode string  `json:"item_code"`
            BrandUID float64  `json:"brand_uid"`
         
    }
    
    // InventoryRequest used by Catalog
    type InventoryRequest struct {

        
            CompanyID float64  `json:"company_id"`
            Sizes []InvSize  `json:"sizes"`
            Item ItemQuery  `json:"item"`
         
    }
    
    // StoreMeta used by Catalog
    type StoreMeta struct {

        
            ID float64  `json:"id"`
         
    }
    
    // ReturnConfig1 used by Catalog
    type ReturnConfig1 struct {

        
            Time float64  `json:"time"`
            Returnable bool  `json:"returnable"`
            Unit string  `json:"unit"`
         
    }
    
    // DimensionResponse used by Catalog
    type DimensionResponse struct {

        
            Length float64  `json:"length"`
            IsDefault bool  `json:"is_default"`
            Width float64  `json:"width"`
            Unit string  `json:"unit"`
            Height float64  `json:"height"`
         
    }
    
    // ManufacturerResponse used by Catalog
    type ManufacturerResponse struct {

        
            Name string  `json:"name"`
            IsDefault bool  `json:"is_default"`
            Address string  `json:"address"`
         
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
            Currency string  `json:"currency"`
            Marked float64  `json:"marked"`
            TpNotes map[string]interface{}  `json:"tp_notes"`
            Transfer float64  `json:"transfer"`
            UpdatedAt string  `json:"updated_at"`
         
    }
    
    // WeightResponse used by Catalog
    type WeightResponse struct {

        
            IsDefault bool  `json:"is_default"`
            Shipping float64  `json:"shipping"`
            Unit string  `json:"unit"`
         
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
    
    // CompanyMeta used by Catalog
    type CompanyMeta struct {

        
            ID float64  `json:"id"`
         
    }
    
    // BrandMeta used by Catalog
    type BrandMeta struct {

        
            Name string  `json:"name"`
            ID float64  `json:"id"`
         
    }
    
    // InventorySellerResponse used by Catalog
    type InventorySellerResponse struct {

        
            Set InventorySet  `json:"set"`
            Identifier map[string]interface{}  `json:"identifier"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Fragile bool  `json:"fragile"`
            IsSet bool  `json:"is_set"`
            Store StoreMeta  `json:"store"`
            ItemID float64  `json:"item_id"`
            SellerIdentifier string  `json:"seller_identifier"`
            ReturnConfig ReturnConfig1  `json:"return_config"`
            RawMeta map[string]interface{}  `json:"raw_meta"`
            IsActive bool  `json:"is_active"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            CountryOfOrigin string  `json:"country_of_origin"`
            FyndItemCode string  `json:"fynd_item_code"`
            Size string  `json:"size"`
            UID string  `json:"uid"`
            Dimension DimensionResponse  `json:"dimension"`
            TotalQuantity float64  `json:"total_quantity"`
            AddedOnStore string  `json:"added_on_store"`
            FyndArticleCode string  `json:"fynd_article_code"`
            CreatedBy UserSerializer  `json:"created_by"`
            Manufacturer ManufacturerResponse  `json:"manufacturer"`
            Trader []Trader1  `json:"trader"`
            Meta map[string]interface{}  `json:"meta"`
            Price PriceMeta  `json:"price"`
            TraceID string  `json:"trace_id"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            TrackInventory bool  `json:"track_inventory"`
            FyndMeta map[string]interface{}  `json:"fynd_meta"`
            Weight WeightResponse  `json:"weight"`
            Quantities Quantities  `json:"quantities"`
            ExpirationDate string  `json:"expiration_date"`
            Tags []string  `json:"tags"`
            Stage string  `json:"stage"`
            Company CompanyMeta  `json:"company"`
            Brand BrandMeta  `json:"brand"`
         
    }
    
    // InventorySellerIdentifierResponsePaginated used by Catalog
    type InventorySellerIdentifierResponsePaginated struct {

        
            Page Page  `json:"page"`
            Items []InventorySellerResponse  `json:"items"`
         
    }
    
    // ArticleStoreResponse used by Catalog
    type ArticleStoreResponse struct {

        
            Name string  `json:"name"`
            UID float64  `json:"uid"`
            StoreCode string  `json:"store_code"`
            StoreType string  `json:"store_type"`
         
    }
    
    // DateMeta used by Catalog
    type DateMeta struct {

        
            InventoryUpdatedOn string  `json:"inventory_updated_on"`
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
            AddedOnStore string  `json:"added_on_store"`
         
    }
    
    // ReturnConfig2 used by Catalog
    type ReturnConfig2 struct {

        
            Time float64  `json:"time"`
            Returnable bool  `json:"returnable"`
            Unit string  `json:"unit"`
         
    }
    
    // DimensionResponse1 used by Catalog
    type DimensionResponse1 struct {

        
            Length float64  `json:"length"`
            Width float64  `json:"width"`
            Height float64  `json:"height"`
            Unit string  `json:"unit"`
         
    }
    
    // ManufacturerResponse1 used by Catalog
    type ManufacturerResponse1 struct {

        
            Name string  `json:"name"`
            IsDefault bool  `json:"is_default"`
            Address string  `json:"address"`
         
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
            Currency string  `json:"currency"`
            Marked float64  `json:"marked"`
            TpNotes map[string]interface{}  `json:"tp_notes"`
            Transfer float64  `json:"transfer"`
         
    }
    
    // WeightResponse1 used by Catalog
    type WeightResponse1 struct {

        
            Shipping float64  `json:"shipping"`
            Unit string  `json:"unit"`
         
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
    
    // CompanyMeta1 used by Catalog
    type CompanyMeta1 struct {

        
            ID float64  `json:"id"`
         
    }
    
    // BrandMeta1 used by Catalog
    type BrandMeta1 struct {

        
            Name string  `json:"name"`
            ID float64  `json:"id"`
         
    }
    
    // GetInventories used by Catalog
    type GetInventories struct {

        
            Identifier map[string]interface{}  `json:"identifier"`
            IsSet bool  `json:"is_set"`
            Store ArticleStoreResponse  `json:"store"`
            ItemID float64  `json:"item_id"`
            SellerIdentifier string  `json:"seller_identifier"`
            DateMeta DateMeta  `json:"date_meta"`
            ReturnConfig ReturnConfig2  `json:"return_config"`
            Platforms map[string]interface{}  `json:"platforms"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            CountryOfOrigin string  `json:"country_of_origin"`
            Size string  `json:"size"`
            UID string  `json:"uid"`
            Dimension DimensionResponse1  `json:"dimension"`
            TotalQuantity float64  `json:"total_quantity"`
            InventoryUpdatedOn string  `json:"inventory_updated_on"`
            ID string  `json:"id"`
            CreatedBy UserSerializer  `json:"created_by"`
            Manufacturer ManufacturerResponse1  `json:"manufacturer"`
            Trader []Trader2  `json:"trader"`
            Price PriceArticle  `json:"price"`
            TraceID string  `json:"trace_id"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            TrackInventory bool  `json:"track_inventory"`
            Weight WeightResponse1  `json:"weight"`
            Quantities QuantitiesArticle  `json:"quantities"`
            ExpirationDate string  `json:"expiration_date"`
            Tags []string  `json:"tags"`
            Stage string  `json:"stage"`
            Company CompanyMeta1  `json:"company"`
            Brand BrandMeta1  `json:"brand"`
         
    }
    
    // GetInventoriesResponse used by Catalog
    type GetInventoriesResponse struct {

        
            Page Page  `json:"page"`
            Items []GetInventories  `json:"items"`
         
    }
    
    // BulkInventoryGetItems used by Catalog
    type BulkInventoryGetItems struct {

        
            CompanyID float64  `json:"company_id"`
            IsActive bool  `json:"is_active"`
            ID string  `json:"id"`
            CreatedOn string  `json:"created_on"`
            CancelledRecords []string  `json:"cancelled_records"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ModifiedOn string  `json:"modified_on"`
            FilePath string  `json:"file_path"`
            Succeed float64  `json:"succeed"`
            Stage string  `json:"stage"`
            FailedRecords []string  `json:"failed_records"`
            Failed float64  `json:"failed"`
            Cancelled float64  `json:"cancelled"`
            Total float64  `json:"total"`
         
    }
    
    // BulkInventoryGet used by Catalog
    type BulkInventoryGet struct {

        
            Page Page  `json:"page"`
            Items []BulkInventoryGetItems  `json:"items"`
         
    }
    
    // InventoryJobPayload used by Catalog
    type InventoryJobPayload struct {

        
            PriceMarked float64  `json:"price_marked"`
            PriceEffective float64  `json:"price_effective"`
            ItemDimensionsUnitOfMeasure string  `json:"item_dimensions_unit_of_measure"`
            ExpirationDate string  `json:"expiration_date"`
            ItemWeightUnitOfMeasure string  `json:"item_weight_unit_of_measure"`
            Tags []string  `json:"tags"`
            StoreCode string  `json:"store_code"`
            Quantity float64  `json:"quantity"`
            SellerIdentifier string  `json:"seller_identifier"`
            Currency string  `json:"currency"`
            Price float64  `json:"price"`
            TotalQuantity float64  `json:"total_quantity"`
            TraceID string  `json:"trace_id"`
         
    }
    
    // InventoryBulkRequest used by Catalog
    type InventoryBulkRequest struct {

        
            CompanyID float64  `json:"company_id"`
            Sizes []InventoryJobPayload  `json:"sizes"`
            BatchID string  `json:"batch_id"`
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

        
            FromDate string  `json:"from_date"`
            ToDate string  `json:"to_date"`
            StoreIds []float64  `json:"store_ids"`
            Quantity InventoryExportQuantityFilter  `json:"quantity"`
            BrandIds []float64  `json:"brand_ids"`
         
    }
    
    // InventoryExportJob used by Catalog
    type InventoryExportJob struct {

        
            CompletedOn string  `json:"completed_on"`
            Filters InventoryExportAdvanceOption  `json:"filters"`
            SellerID float64  `json:"seller_id"`
            NotificationEmails []string  `json:"notification_emails"`
            Status string  `json:"status"`
            URL string  `json:"url"`
            TaskID string  `json:"task_id"`
            Type string  `json:"type"`
         
    }
    
    // InventoryExportRequest used by Catalog
    type InventoryExportRequest struct {

        
            Store []float64  `json:"store"`
            Type string  `json:"type"`
            Brand []float64  `json:"brand"`
         
    }
    
    // InventoryExportResponse used by Catalog
    type InventoryExportResponse struct {

        
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
            CreatedBy string  `json:"created_by"`
            Filters map[string]interface{}  `json:"filters"`
            SellerID float64  `json:"seller_id"`
            NotificationEmails []string  `json:"notification_emails"`
            Status string  `json:"status"`
            TaskID string  `json:"task_id"`
            Type string  `json:"type"`
         
    }
    
    // InventoryJobFilters used by Catalog
    type InventoryJobFilters struct {

        
            Stores []string  `json:"stores"`
            Brands []string  `json:"brands"`
            ToDate string  `json:"to_date"`
            FromDate string  `json:"from_date"`
            Quantity InventoryExportQuantityFilter  `json:"quantity"`
         
    }
    
    // InventoryJobDetailResponse used by Catalog
    type InventoryJobDetailResponse struct {

        
            ID string  `json:"id"`
            CompletedOn string  `json:"completed_on"`
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
            CreatedBy UserDetail  `json:"created_by"`
            Filters InventoryJobFilters  `json:"filters"`
            SellerID float64  `json:"seller_id"`
            NotificationEmails []string  `json:"notification_emails"`
            Status interface{}  `json:"status"`
            CancelledOn string  `json:"cancelled_on"`
            URL string  `json:"url"`
            TaskID string  `json:"task_id"`
            CancelledBy UserDetail  `json:"cancelled_by"`
            Type string  `json:"type"`
         
    }
    
    // InventoryExportJobListResponse used by Catalog
    type InventoryExportJobListResponse struct {

        
            Items InventoryJobDetailResponse  `json:"items"`
         
    }
    
    // InventoryExportFilter used by Catalog
    type InventoryExportFilter struct {

        
            FromDate string  `json:"from_date"`
            ToDate string  `json:"to_date"`
            StoreIds []float64  `json:"store_ids"`
            Quantity InventoryExportQuantityFilter  `json:"quantity"`
            BrandIds []float64  `json:"brand_ids"`
         
    }
    
    // InventoryCreateRequest used by Catalog
    type InventoryCreateRequest struct {

        
            NotificationEmails []string  `json:"notification_emails"`
            Type string  `json:"type"`
            Data []string  `json:"data"`
            Filters InventoryExportFilter  `json:"filters"`
         
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
            PriceMarked float64  `json:"price_marked"`
            PriceEffective float64  `json:"price_effective"`
            ExpirationDate string  `json:"expiration_date"`
            Tags []string  `json:"tags"`
            SellerIdentifier string  `json:"seller_identifier"`
            TotalQuantity float64  `json:"total_quantity"`
            TraceID string  `json:"trace_id"`
         
    }
    
    // InventoryRequestSchemaV2 used by Catalog
    type InventoryRequestSchemaV2 struct {

        
            CompanyID float64  `json:"company_id"`
            Meta map[string]interface{}  `json:"meta"`
            Payload []InventoryPayload  `json:"payload"`
         
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
    
    // HsnCodesObject used by Catalog
    type HsnCodesObject struct {

        
            CompanyID float64  `json:"company_id"`
            ID string  `json:"id"`
            ModifiedOn string  `json:"modified_on"`
            TaxOnMrp bool  `json:"tax_on_mrp"`
            Threshold2 float64  `json:"threshold2"`
            Tax2 float64  `json:"tax2"`
            Tax1 float64  `json:"tax1"`
            Hs2Code string  `json:"hs2_code"`
            TaxOnEsp bool  `json:"tax_on_esp"`
            Threshold1 float64  `json:"threshold1"`
            HsnCode string  `json:"hsn_code"`
         
    }
    
    // HsnCode used by Catalog
    type HsnCode struct {

        
            Data HsnCodesObject  `json:"data"`
         
    }
    
    // HsnUpsert used by Catalog
    type HsnUpsert struct {

        
            CompanyID float64  `json:"company_id"`
            TaxOnMrp bool  `json:"tax_on_mrp"`
            Threshold2 float64  `json:"threshold2"`
            Tax2 float64  `json:"tax2"`
            UID float64  `json:"uid"`
            Tax1 float64  `json:"tax1"`
            Hs2Code string  `json:"hs2_code"`
            TaxOnEsp bool  `json:"tax_on_esp"`
            Threshold1 float64  `json:"threshold1"`
            IsActive bool  `json:"is_active"`
            HsnCode string  `json:"hsn_code"`
         
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
            Size float64  `json:"size"`
            HasNext bool  `json:"has_next"`
            Current string  `json:"current"`
         
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

        
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Taxes []TaxSlab  `json:"taxes"`
            CreatedOn string  `json:"created_on"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            ModifiedOn string  `json:"modified_on"`
            ReportingHsn string  `json:"reporting_hsn"`
            HsnCodeID string  `json:"hsn_code_id"`
            CountryCode string  `json:"country_code"`
            Description string  `json:"description"`
            Type string  `json:"type"`
            HsnCode string  `json:"hsn_code"`
         
    }
    
    // HsnCodesListingResponseSchemaV2 used by Catalog
    type HsnCodesListingResponseSchemaV2 struct {

        
            Page PageResponse  `json:"page"`
            Items []HSNDataInsertV2  `json:"items"`
         
    }
    
    // BrandItem used by Catalog
    type BrandItem struct {

        
            Departments []string  `json:"departments"`
            Slug string  `json:"slug"`
            Discount string  `json:"discount"`
            UID float64  `json:"uid"`
            Action Action  `json:"action"`
            Logo Media  `json:"logo"`
            Banners ImageUrls  `json:"banners"`
            Name string  `json:"name"`
         
    }
    
    // BrandListingResponse used by Catalog
    type BrandListingResponse struct {

        
            Page Page  `json:"page"`
            Items []BrandItem  `json:"items"`
         
    }
    
    // Department used by Catalog
    type Department struct {

        
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
            Logo Media  `json:"logo"`
            Name string  `json:"name"`
            PriorityOrder float64  `json:"priority_order"`
         
    }
    
    // DepartmentResponse used by Catalog
    type DepartmentResponse struct {

        
            Items []Department  `json:"items"`
         
    }
    
    // ThirdLevelChild used by Catalog
    type ThirdLevelChild struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
            Childs []map[string]interface{}  `json:"childs"`
            Action Action  `json:"action"`
            Banners ImageUrls  `json:"banners"`
            Name string  `json:"name"`
         
    }
    
    // SecondLevelChild used by Catalog
    type SecondLevelChild struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
            Childs []ThirdLevelChild  `json:"childs"`
            Action Action  `json:"action"`
            Banners ImageUrls  `json:"banners"`
            Name string  `json:"name"`
         
    }
    
    // Child used by Catalog
    type Child struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
            Childs []SecondLevelChild  `json:"childs"`
            Action Action  `json:"action"`
            Banners ImageUrls  `json:"banners"`
            Name string  `json:"name"`
         
    }
    
    // CategoryItems used by Catalog
    type CategoryItems struct {

        
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
            Childs []Child  `json:"childs"`
            Action Action  `json:"action"`
            Banners ImageUrls  `json:"banners"`
            Name string  `json:"name"`
         
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

        
            Operators map[string]interface{}  `json:"operators"`
            Items []ProductListingDetail  `json:"items"`
            Filters []ProductFilters  `json:"filters"`
            Page Page  `json:"page"`
            SortOn []ProductSortOn  `json:"sort_on"`
         
    }
    
    // ProductDetail used by Catalog
    type ProductDetail struct {

        
            HasVariant bool  `json:"has_variant"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            RatingCount float64  `json:"rating_count"`
            Tryouts []string  `json:"tryouts"`
            UID float64  `json:"uid"`
            ItemType string  `json:"item_type"`
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
            ShortDescription string  `json:"short_description"`
            Color string  `json:"color"`
            Name string  `json:"name"`
            ImageNature string  `json:"image_nature"`
            ItemCode string  `json:"item_code"`
            Rating float64  `json:"rating"`
            ProductOnlineDate string  `json:"product_online_date"`
            Similars []string  `json:"similars"`
            Highlights []string  `json:"highlights"`
            Slug string  `json:"slug"`
            Medias []Media1  `json:"medias"`
            Description string  `json:"description"`
            Attributes map[string]interface{}  `json:"attributes"`
            PromoMeta map[string]interface{}  `json:"promo_meta"`
            Type string  `json:"type"`
            Brand ProductBrand  `json:"brand"`
         
    }
    
    // InventoryPage used by Catalog
    type InventoryPage struct {

        
            HasPrevious bool  `json:"has_previous"`
            ItemTotal float64  `json:"item_total"`
            NextID string  `json:"next_id"`
            HasNext bool  `json:"has_next"`
            Type string  `json:"type"`
         
    }
    
    // InventoryStockResponse used by Catalog
    type InventoryStockResponse struct {

        
            Page InventoryPage  `json:"page"`
            Items []map[string]interface{}  `json:"items"`
         
    }
    
    // ArticleQuery used by Catalog
    type ArticleQuery struct {

        
            IgnoredStores []float64  `json:"ignored_stores"`
            ItemID float64  `json:"item_id"`
            Size string  `json:"size"`
         
    }
    
    // ArticleAssignment used by Catalog
    type ArticleAssignment struct {

        
            Strategy string  `json:"strategy"`
            Level string  `json:"level"`
         
    }
    
    // AssignStoreArticle used by Catalog
    type AssignStoreArticle struct {

        
            GroupID string  `json:"group_id"`
            Quantity float64  `json:"quantity"`
            Meta map[string]interface{}  `json:"meta"`
            Query ArticleQuery  `json:"query"`
            ArticleAssignment ArticleAssignment  `json:"article_assignment"`
         
    }
    
    // AssignStore used by Catalog
    type AssignStore struct {

        
            CompanyID float64  `json:"company_id"`
            ChannelType string  `json:"channel_type"`
            Articles []AssignStoreArticle  `json:"articles"`
            StoreIds []float64  `json:"store_ids"`
            ChannelIdentifier string  `json:"channel_identifier"`
            AppID string  `json:"app_id"`
            Pincode string  `json:"pincode"`
         
    }
    
    // ArticleAssignment1 used by Catalog
    type ArticleAssignment1 struct {

        
            Strategy string  `json:"strategy"`
            Level string  `json:"level"`
         
    }
    
    // StoreAssignResponse used by Catalog
    type StoreAssignResponse struct {

        
            ID string  `json:"_id"`
            Index float64  `json:"index"`
            CompanyID float64  `json:"company_id"`
            StoreID float64  `json:"store_id"`
            PriceMarked float64  `json:"price_marked"`
            StorePincode float64  `json:"store_pincode"`
            StrategyWiseListing []map[string]interface{}  `json:"strategy_wise_listing"`
            PriceEffective float64  `json:"price_effective"`
            GroupID string  `json:"group_id"`
            Status bool  `json:"status"`
            Size string  `json:"size"`
            UID string  `json:"uid"`
            SCity string  `json:"s_city"`
            Quantity float64  `json:"quantity"`
            ItemID float64  `json:"item_id"`
            Meta map[string]interface{}  `json:"meta"`
            ArticleAssignment ArticleAssignment1  `json:"article_assignment"`
         
    }
    
    // SellerPhoneNumber used by Catalog
    type SellerPhoneNumber struct {

        
            Number string  `json:"number"`
            CountryCode float64  `json:"country_code"`
         
    }
    
    // LocationManagerSerializer used by Catalog
    type LocationManagerSerializer struct {

        
            Name string  `json:"name"`
            Email string  `json:"email"`
            MobileNo SellerPhoneNumber  `json:"mobile_no"`
         
    }
    
    // ProductReturnConfigSerializer used by Catalog
    type ProductReturnConfigSerializer struct {

        
            OnSameStore bool  `json:"on_same_store"`
            StoreUID float64  `json:"store_uid"`
         
    }
    
    // InvoiceCredSerializer used by Catalog
    type InvoiceCredSerializer struct {

        
            Enabled bool  `json:"enabled"`
            Username string  `json:"username"`
            Password string  `json:"password"`
         
    }
    
    // InvoiceDetailsSerializer used by Catalog
    type InvoiceDetailsSerializer struct {

        
            EWaybill InvoiceCredSerializer  `json:"e_waybill"`
            EInvoice InvoiceCredSerializer  `json:"e_invoice"`
         
    }
    
    // UserSerializer1 used by Catalog
    type UserSerializer1 struct {

        
            Contact string  `json:"contact"`
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
         
    }
    
    // LocationIntegrationType used by Catalog
    type LocationIntegrationType struct {

        
            Inventory string  `json:"inventory"`
            Order string  `json:"order"`
         
    }
    
    // GetAddressSerializer used by Catalog
    type GetAddressSerializer struct {

        
            Longitude float64  `json:"longitude"`
            City string  `json:"city"`
            Address2 string  `json:"address2"`
            Country string  `json:"country"`
            CountryCode string  `json:"country_code"`
            Address1 string  `json:"address1"`
            AddressType string  `json:"address_type"`
            State string  `json:"state"`
            Pincode float64  `json:"pincode"`
            Latitude float64  `json:"latitude"`
            Landmark string  `json:"landmark"`
         
    }
    
    // LocationTimingSerializer used by Catalog
    type LocationTimingSerializer struct {

        
            Minute float64  `json:"minute"`
            Hour float64  `json:"hour"`
         
    }
    
    // LocationDayWiseSerializer used by Catalog
    type LocationDayWiseSerializer struct {

        
            Opening LocationTimingSerializer  `json:"opening"`
            Weekday string  `json:"weekday"`
            Closing LocationTimingSerializer  `json:"closing"`
            Open bool  `json:"open"`
         
    }
    
    // Document used by Catalog
    type Document struct {

        
            Verified bool  `json:"verified"`
            Type string  `json:"type"`
            Value string  `json:"value"`
            URL string  `json:"url"`
            LegalName string  `json:"legal_name"`
         
    }
    
    // UserSerializer2 used by Catalog
    type UserSerializer2 struct {

        
            Contact string  `json:"contact"`
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
         
    }
    
    // GetCompanySerializer used by Catalog
    type GetCompanySerializer struct {

        
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
            ModifiedBy UserSerializer2  `json:"modified_by"`
            CreatedBy UserSerializer2  `json:"created_by"`
            BusinessType string  `json:"business_type"`
            UID float64  `json:"uid"`
            VerifiedOn string  `json:"verified_on"`
            VerifiedBy UserSerializer2  `json:"verified_by"`
            RejectReason string  `json:"reject_reason"`
            Stage string  `json:"stage"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            Name string  `json:"name"`
            CompanyType string  `json:"company_type"`
         
    }
    
    // GetLocationSerializer used by Catalog
    type GetLocationSerializer struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            PhoneNumber string  `json:"phone_number"`
            Warnings map[string]interface{}  `json:"warnings"`
            VerifiedOn string  `json:"verified_on"`
            Manager LocationManagerSerializer  `json:"manager"`
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            ProductReturnConfig ProductReturnConfigSerializer  `json:"product_return_config"`
            UID float64  `json:"uid"`
            GstCredentials InvoiceDetailsSerializer  `json:"gst_credentials"`
            CreatedBy UserSerializer1  `json:"created_by"`
            StoreType string  `json:"store_type"`
            VerifiedBy UserSerializer1  `json:"verified_by"`
            Name string  `json:"name"`
            IntegrationType LocationIntegrationType  `json:"integration_type"`
            ModifiedBy UserSerializer1  `json:"modified_by"`
            Address GetAddressSerializer  `json:"address"`
            NotificationEmails []string  `json:"notification_emails"`
            Code string  `json:"code"`
            Timing []LocationDayWiseSerializer  `json:"timing"`
            Stage string  `json:"stage"`
            Documents []Document  `json:"documents"`
            Company GetCompanySerializer  `json:"company"`
            DisplayName string  `json:"display_name"`
         
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

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            UID float64  `json:"uid"`
            Logo string  `json:"logo"`
            AppID string  `json:"app_id"`
            Name string  `json:"name"`
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
    

    
    // GetAddressSerializer used by CompanyProfile
    type GetAddressSerializer struct {

        
            Address1 string  `json:"address1"`
            Pincode float64  `json:"pincode"`
            CountryCode string  `json:"country_code"`
            Latitude float64  `json:"latitude"`
            Address2 string  `json:"address2"`
            Landmark string  `json:"landmark"`
            AddressType string  `json:"address_type"`
            State string  `json:"state"`
            City string  `json:"city"`
            Country string  `json:"country"`
            Longitude float64  `json:"longitude"`
         
    }
    
    // BusinessCountryInfo used by CompanyProfile
    type BusinessCountryInfo struct {

        
            Country string  `json:"country"`
            CountryCode string  `json:"country_code"`
         
    }
    
    // UserSerializer used by CompanyProfile
    type UserSerializer struct {

        
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
            Contact string  `json:"contact"`
         
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
            FranchiseEnabled bool  `json:"franchise_enabled"`
            BusinessCountryInfo BusinessCountryInfo  `json:"business_country_info"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            CreatedBy UserSerializer  `json:"created_by"`
            Name string  `json:"name"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            ContactDetails ContactDetails  `json:"contact_details"`
            UID float64  `json:"uid"`
            Warnings map[string]interface{}  `json:"warnings"`
            CompanyType string  `json:"company_type"`
            Stage string  `json:"stage"`
            NotificationEmails []string  `json:"notification_emails"`
            Taxes []CompanyTaxesSerializer  `json:"taxes"`
            BusinessType string  `json:"business_type"`
            CreatedOn string  `json:"created_on"`
            Mode string  `json:"mode"`
            VerifiedOn string  `json:"verified_on"`
            ModifiedOn string  `json:"modified_on"`
            Documents []Document  `json:"documents"`
            BusinessDetails BusinessDetails  `json:"business_details"`
            BusinessInfo string  `json:"business_info"`
         
    }
    
    // ErrorResponse used by CompanyProfile
    type ErrorResponse struct {

        
            Status float64  `json:"status"`
            Message string  `json:"message"`
            Code string  `json:"code"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // CreateUpdateAddressSerializer used by CompanyProfile
    type CreateUpdateAddressSerializer struct {

        
            Address1 string  `json:"address1"`
            Pincode float64  `json:"pincode"`
            CountryCode string  `json:"country_code"`
            Latitude float64  `json:"latitude"`
            Address2 string  `json:"address2"`
            Landmark string  `json:"landmark"`
            AddressType string  `json:"address_type"`
            State string  `json:"state"`
            City string  `json:"city"`
            Country string  `json:"country"`
            Longitude float64  `json:"longitude"`
         
    }
    
    // CompanyTaxesSerializer1 used by CompanyProfile
    type CompanyTaxesSerializer1 struct {

        
            EffectiveDate string  `json:"effective_date"`
            Enable bool  `json:"enable"`
            Rate float64  `json:"rate"`
         
    }
    
    // UpdateCompany used by CompanyProfile
    type UpdateCompany struct {

        
            Addresses []CreateUpdateAddressSerializer  `json:"addresses"`
            FranchiseEnabled bool  `json:"franchise_enabled"`
            Warnings map[string]interface{}  `json:"warnings"`
            RejectReason string  `json:"reject_reason"`
            CompanyType string  `json:"company_type"`
            Documents []Document  `json:"documents"`
            NotificationEmails []string  `json:"notification_emails"`
            BusinessDetails BusinessDetails  `json:"business_details"`
            BusinessInfo string  `json:"business_info"`
            BusinessType string  `json:"business_type"`
            Name string  `json:"name"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Taxes []CompanyTaxesSerializer1  `json:"taxes"`
            ContactDetails ContactDetails  `json:"contact_details"`
         
    }
    
    // ProfileSuccessResponse used by CompanyProfile
    type ProfileSuccessResponse struct {

        
            UID float64  `json:"uid"`
            Success bool  `json:"success"`
         
    }
    
    // DocumentsObj used by CompanyProfile
    type DocumentsObj struct {

        
            Pending float64  `json:"pending"`
            Verified float64  `json:"verified"`
         
    }
    
    // MetricsSerializer used by CompanyProfile
    type MetricsSerializer struct {

        
            Product DocumentsObj  `json:"product"`
            CompanyDocuments DocumentsObj  `json:"company_documents"`
            StoreDocuments DocumentsObj  `json:"store_documents"`
            Store DocumentsObj  `json:"store"`
            Stage string  `json:"stage"`
            Brand DocumentsObj  `json:"brand"`
            UID float64  `json:"uid"`
         
    }
    
    // BrandBannerSerializer used by CompanyProfile
    type BrandBannerSerializer struct {

        
            Portrait string  `json:"portrait"`
            Landscape string  `json:"landscape"`
         
    }
    
    // GetBrandResponseSerializer used by CompanyProfile
    type GetBrandResponseSerializer struct {

        
            Synonyms []string  `json:"synonyms"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            CreatedBy UserSerializer  `json:"created_by"`
            Name string  `json:"name"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            UID float64  `json:"uid"`
            Logo string  `json:"logo"`
            Warnings map[string]interface{}  `json:"warnings"`
            Stage string  `json:"stage"`
            RejectReason string  `json:"reject_reason"`
            CreatedOn string  `json:"created_on"`
            Mode string  `json:"mode"`
            VerifiedOn string  `json:"verified_on"`
            ModifiedOn string  `json:"modified_on"`
            Description string  `json:"description"`
            SlugKey string  `json:"slug_key"`
            Banner BrandBannerSerializer  `json:"banner"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
         
    }
    
    // CreateUpdateBrandRequestSerializer used by CompanyProfile
    type CreateUpdateBrandRequestSerializer struct {

        
            Logo string  `json:"logo"`
            Synonyms []string  `json:"synonyms"`
            Description string  `json:"description"`
            CompanyID float64  `json:"company_id"`
            BrandTier string  `json:"brand_tier"`
            Banner BrandBannerSerializer  `json:"banner"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Name string  `json:"name"`
            UID float64  `json:"uid"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
         
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

        
            Addresses []GetAddressSerializer  `json:"addresses"`
            BusinessCountryInfo BusinessCountryInfo  `json:"business_country_info"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            RejectReason string  `json:"reject_reason"`
            CompanyType string  `json:"company_type"`
            NotificationEmails []string  `json:"notification_emails"`
            BusinessType string  `json:"business_type"`
            CreatedOn string  `json:"created_on"`
            CreatedBy UserSerializer  `json:"created_by"`
            Name string  `json:"name"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            Details CompanyDetails  `json:"details"`
            MarketChannels []string  `json:"market_channels"`
            Stage string  `json:"stage"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            VerifiedOn string  `json:"verified_on"`
            ModifiedOn string  `json:"modified_on"`
            UID float64  `json:"uid"`
         
    }
    
    // CompanyBrandSerializer used by CompanyProfile
    type CompanyBrandSerializer struct {

        
            Warnings map[string]interface{}  `json:"warnings"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            RejectReason string  `json:"reject_reason"`
            CreatedOn string  `json:"created_on"`
            CreatedBy UserSerializer  `json:"created_by"`
            Brand GetBrandResponseSerializer  `json:"brand"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            Stage string  `json:"stage"`
            VerifiedOn string  `json:"verified_on"`
            ModifiedOn string  `json:"modified_on"`
            UID float64  `json:"uid"`
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

        
            Brands []float64  `json:"brands"`
            UID float64  `json:"uid"`
            Company float64  `json:"company"`
         
    }
    
    // GetCompanySerializer used by CompanyProfile
    type GetCompanySerializer struct {

        
            Addresses []GetAddressSerializer  `json:"addresses"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            RejectReason string  `json:"reject_reason"`
            CompanyType string  `json:"company_type"`
            BusinessType string  `json:"business_type"`
            CreatedOn string  `json:"created_on"`
            CreatedBy UserSerializer  `json:"created_by"`
            Name string  `json:"name"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            Stage string  `json:"stage"`
            VerifiedOn string  `json:"verified_on"`
            ModifiedOn string  `json:"modified_on"`
            UID float64  `json:"uid"`
         
    }
    
    // LocationManagerSerializer used by CompanyProfile
    type LocationManagerSerializer struct {

        
            Name string  `json:"name"`
            Email string  `json:"email"`
            MobileNo SellerPhoneNumber  `json:"mobile_no"`
         
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

        
            EndDate string  `json:"end_date"`
            StartDate string  `json:"start_date"`
         
    }
    
    // HolidaySchemaSerializer used by CompanyProfile
    type HolidaySchemaSerializer struct {

        
            Title string  `json:"title"`
            Date HolidayDateSerializer  `json:"date"`
            HolidayType string  `json:"holiday_type"`
         
    }
    
    // ProductReturnConfigSerializer used by CompanyProfile
    type ProductReturnConfigSerializer struct {

        
            StoreUID float64  `json:"store_uid"`
            OnSameStore bool  `json:"on_same_store"`
         
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
            Weekday string  `json:"weekday"`
            Closing LocationTimingSerializer  `json:"closing"`
         
    }
    
    // GetLocationSerializer used by CompanyProfile
    type GetLocationSerializer struct {

        
            PhoneNumber string  `json:"phone_number"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            CreatedBy UserSerializer  `json:"created_by"`
            Name string  `json:"name"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            UID float64  `json:"uid"`
            Warnings map[string]interface{}  `json:"warnings"`
            DisplayName string  `json:"display_name"`
            Stage string  `json:"stage"`
            Code string  `json:"code"`
            Company GetCompanySerializer  `json:"company"`
            Manager LocationManagerSerializer  `json:"manager"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            Address GetAddressSerializer  `json:"address"`
            NotificationEmails []string  `json:"notification_emails"`
            GstCredentials InvoiceDetailsSerializer  `json:"gst_credentials"`
            CreatedOn string  `json:"created_on"`
            Holiday []HolidaySchemaSerializer  `json:"holiday"`
            VerifiedOn string  `json:"verified_on"`
            ModifiedOn string  `json:"modified_on"`
            Documents []Document  `json:"documents"`
            ProductReturnConfig ProductReturnConfigSerializer  `json:"product_return_config"`
            Timing []LocationDayWiseSerializer  `json:"timing"`
            StoreType string  `json:"store_type"`
         
    }
    
    // LocationListSerializer used by CompanyProfile
    type LocationListSerializer struct {

        
            Items []GetLocationSerializer  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // AddressSerializer used by CompanyProfile
    type AddressSerializer struct {

        
            Address1 string  `json:"address1"`
            Pincode float64  `json:"pincode"`
            CountryCode string  `json:"country_code"`
            Latitude float64  `json:"latitude"`
            Address2 string  `json:"address2"`
            Landmark string  `json:"landmark"`
            AddressType string  `json:"address_type"`
            State string  `json:"state"`
            City string  `json:"city"`
            Country string  `json:"country"`
            Longitude float64  `json:"longitude"`
         
    }
    
    // LocationSerializer used by CompanyProfile
    type LocationSerializer struct {

        
            Manager LocationManagerSerializer  `json:"manager"`
            Warnings map[string]interface{}  `json:"warnings"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            Documents []Document  `json:"documents"`
            DisplayName string  `json:"display_name"`
            Address AddressSerializer  `json:"address"`
            NotificationEmails []string  `json:"notification_emails"`
            ProductReturnConfig ProductReturnConfigSerializer  `json:"product_return_config"`
            GstCredentials InvoiceDetailsSerializer  `json:"gst_credentials"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Name string  `json:"name"`
            Timing []LocationDayWiseSerializer  `json:"timing"`
            Stage string  `json:"stage"`
            Holiday []HolidaySchemaSerializer  `json:"holiday"`
            Code string  `json:"code"`
            UID float64  `json:"uid"`
            StoreType string  `json:"store_type"`
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
            Gstores []GStore  `json:"gstores"`
         
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

        
            Integration string  `json:"integration"`
            IntegrationData map[string]map[string]interface{}  `json:"integration_data"`
            CompanyName string  `json:"company_name"`
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

        
            Integration string  `json:"integration"`
            CompanyName string  `json:"company_name"`
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
            Mapping map[string]PropBeanConfig  `json:"mapping"`
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
    

    
    // PostOrder used by Cart
    type PostOrder struct {

        
            CancellationAllowed bool  `json:"cancellation_allowed"`
            ReturnAllowed bool  `json:"return_allowed"`
         
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
            Uses PaymentAllowValue  `json:"uses"`
         
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
    
    // BulkBundleRestriction used by Cart
    type BulkBundleRestriction struct {

        
            MultiStoreAllowed bool  `json:"multi_store_allowed"`
         
    }
    
    // PriceRange used by Cart
    type PriceRange struct {

        
            Min float64  `json:"min"`
            Max float64  `json:"max"`
         
    }
    
    // Restrictions used by Cart
    type Restrictions struct {

        
            UserGroups []float64  `json:"user_groups"`
            PostOrder PostOrder  `json:"post_order"`
            Payments map[string]PaymentModes  `json:"payments"`
            Platforms []string  `json:"platforms"`
            Uses UsesRestriction  `json:"uses"`
            OrderingStores []float64  `json:"ordering_stores"`
            CouponAllowed bool  `json:"coupon_allowed"`
            BulkBundle BulkBundleRestriction  `json:"bulk_bundle"`
            UserType string  `json:"user_type"`
            PriceRange PriceRange  `json:"price_range"`
         
    }
    
    // DisplayMetaDict used by Cart
    type DisplayMetaDict struct {

        
            Subtitle string  `json:"subtitle"`
            Title string  `json:"title"`
         
    }
    
    // DisplayMeta used by Cart
    type DisplayMeta struct {

        
            Auto DisplayMetaDict  `json:"auto"`
            Description string  `json:"description"`
            Subtitle string  `json:"subtitle"`
            Title string  `json:"title"`
            Apply DisplayMetaDict  `json:"apply"`
            Remove DisplayMetaDict  `json:"remove"`
         
    }
    
    // RuleDefinition used by Cart
    type RuleDefinition struct {

        
            CalculateOn string  `json:"calculate_on"`
            AutoApply bool  `json:"auto_apply"`
            Type string  `json:"type"`
            ValueType string  `json:"value_type"`
            Scope []string  `json:"scope"`
            ApplicableOn string  `json:"applicable_on"`
            CurrencyCode string  `json:"currency_code"`
            IsExact bool  `json:"is_exact"`
         
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
    
    // CouponDateMeta used by Cart
    type CouponDateMeta struct {

        
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // Rule used by Cart
    type Rule struct {

        
            DiscountQty float64  `json:"discount_qty"`
            Key float64  `json:"key"`
            Value float64  `json:"value"`
            Min float64  `json:"min"`
            Max float64  `json:"max"`
         
    }
    
    // Ownership used by Cart
    type Ownership struct {

        
            PayableCategory string  `json:"payable_category"`
            PayableBy string  `json:"payable_by"`
         
    }
    
    // State used by Cart
    type State struct {

        
            IsDisplay bool  `json:"is_display"`
            IsPublic bool  `json:"is_public"`
            IsArchived bool  `json:"is_archived"`
         
    }
    
    // Validity used by Cart
    type Validity struct {

        
            Priority float64  `json:"priority"`
         
    }
    
    // CouponSchedule used by Cart
    type CouponSchedule struct {

        
            Start string  `json:"start"`
            Duration float64  `json:"duration"`
            Cron string  `json:"cron"`
            End string  `json:"end"`
            NextSchedule []map[string]interface{}  `json:"next_schedule"`
         
    }
    
    // Identifier used by Cart
    type Identifier struct {

        
            CollectionID []string  `json:"collection_id"`
            UserID []string  `json:"user_id"`
            CompanyID []float64  `json:"company_id"`
            EmailDomain []string  `json:"email_domain"`
            BrandID []float64  `json:"brand_id"`
            ExcludeBrandID []float64  `json:"exclude_brand_id"`
            ArticleID []string  `json:"article_id"`
            ItemID []float64  `json:"item_id"`
            StoreID []float64  `json:"store_id"`
            CategoryID []float64  `json:"category_id"`
         
    }
    
    // Validation used by Cart
    type Validation struct {

        
            Anonymous bool  `json:"anonymous"`
            AppID []string  `json:"app_id"`
            UserRegisteredAfter string  `json:"user_registered_after"`
         
    }
    
    // CouponAdd used by Cart
    type CouponAdd struct {

        
            Restrictions Restrictions  `json:"restrictions"`
            DisplayMeta DisplayMeta  `json:"display_meta"`
            Tags []string  `json:"tags"`
            RuleDefinition RuleDefinition  `json:"rule_definition"`
            Action CouponAction  `json:"action"`
            Author CouponAuthor  `json:"author"`
            DateMeta CouponDateMeta  `json:"date_meta"`
            Rule []Rule  `json:"rule"`
            TypeSlug string  `json:"type_slug"`
            Code string  `json:"code"`
            Ownership Ownership  `json:"ownership"`
            State State  `json:"state"`
            Validity Validity  `json:"validity"`
            Schedule CouponSchedule  `json:"_schedule"`
            Identifiers Identifier  `json:"identifiers"`
            Validation Validation  `json:"validation"`
         
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

        
            Restrictions Restrictions  `json:"restrictions"`
            DisplayMeta DisplayMeta  `json:"display_meta"`
            Tags []string  `json:"tags"`
            RuleDefinition RuleDefinition  `json:"rule_definition"`
            Action CouponAction  `json:"action"`
            Author CouponAuthor  `json:"author"`
            DateMeta CouponDateMeta  `json:"date_meta"`
            Rule []Rule  `json:"rule"`
            TypeSlug string  `json:"type_slug"`
            Code string  `json:"code"`
            Ownership Ownership  `json:"ownership"`
            State State  `json:"state"`
            Validity Validity  `json:"validity"`
            Schedule CouponSchedule  `json:"_schedule"`
            Identifiers Identifier  `json:"identifiers"`
            Validation Validation  `json:"validation"`
         
    }
    
    // CouponPartialUpdate used by Cart
    type CouponPartialUpdate struct {

        
            Archive bool  `json:"archive"`
            Schedule CouponSchedule  `json:"schedule"`
         
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
            Type string  `json:"type"`
            Uses PaymentAllowValue1  `json:"uses"`
         
    }
    
    // Restrictions1 used by Cart
    type Restrictions1 struct {

        
            UserGroups []float64  `json:"user_groups"`
            UserID []string  `json:"user_id"`
            PostOrder PostOrder1  `json:"post_order"`
            UserRegistered UserRegistered  `json:"user_registered"`
            Platforms []string  `json:"platforms"`
            Uses UsesRestriction1  `json:"uses"`
            AnonymousUsers bool  `json:"anonymous_users"`
            Payments []PromotionPaymentModes  `json:"payments"`
            OrderQuantity float64  `json:"order_quantity"`
         
    }
    
    // DisplayMeta1 used by Cart
    type DisplayMeta1 struct {

        
            OfferText string  `json:"offer_text"`
            Name string  `json:"name"`
            Description string  `json:"description"`
         
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
    
    // PromotionDateMeta used by Cart
    type PromotionDateMeta struct {

        
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // CompareObject used by Cart
    type CompareObject struct {

        
            LessThanEquals float64  `json:"less_than_equals"`
            GreaterThan float64  `json:"greater_than"`
            Equals float64  `json:"equals"`
            GreaterThanEquals float64  `json:"greater_than_equals"`
            LessThan float64  `json:"less_than"`
         
    }
    
    // ItemCriteria used by Cart
    type ItemCriteria struct {

        
            AllItems bool  `json:"all_items"`
            ItemExcludeID []float64  `json:"item_exclude_id"`
            CartTotal CompareObject  `json:"cart_total"`
            CartQuantity CompareObject  `json:"cart_quantity"`
            CartUniqueItemAmount CompareObject  `json:"cart_unique_item_amount"`
            ItemExcludeStore []float64  `json:"item_exclude_store"`
            ItemBrand []float64  `json:"item_brand"`
            ItemSku []string  `json:"item_sku"`
            ItemCompany []float64  `json:"item_company"`
            AvailableZones []string  `json:"available_zones"`
            ItemSize []string  `json:"item_size"`
            ItemExcludeCompany []float64  `json:"item_exclude_company"`
            ItemExcludeSku []string  `json:"item_exclude_sku"`
            ItemExcludeBrand []float64  `json:"item_exclude_brand"`
            ItemCategory []float64  `json:"item_category"`
            ItemExcludeCategory []float64  `json:"item_exclude_category"`
            CartUniqueItemQuantity CompareObject  `json:"cart_unique_item_quantity"`
            ItemID []float64  `json:"item_id"`
            BuyRules []string  `json:"buy_rules"`
            ItemStore []float64  `json:"item_store"`
         
    }
    
    // DiscountOffer used by Cart
    type DiscountOffer struct {

        
            MaxOfferQuantity float64  `json:"max_offer_quantity"`
            PartialCanRet bool  `json:"partial_can_ret"`
            DiscountPercentage float64  `json:"discount_percentage"`
            MaxUsagePerTransaction float64  `json:"max_usage_per_transaction"`
            MaxDiscountAmount float64  `json:"max_discount_amount"`
            DiscountAmount float64  `json:"discount_amount"`
            DiscountPrice float64  `json:"discount_price"`
            Code string  `json:"code"`
            MinOfferQuantity float64  `json:"min_offer_quantity"`
            ApportionDiscount bool  `json:"apportion_discount"`
         
    }
    
    // DiscountRule used by Cart
    type DiscountRule struct {

        
            ItemCriteria ItemCriteria  `json:"item_criteria"`
            Offer DiscountOffer  `json:"offer"`
            BuyCondition string  `json:"buy_condition"`
            DiscountType string  `json:"discount_type"`
         
    }
    
    // PromotionSchedule used by Cart
    type PromotionSchedule struct {

        
            Start string  `json:"start"`
            Published bool  `json:"published"`
            Duration float64  `json:"duration"`
            Cron string  `json:"cron"`
            End string  `json:"end"`
            NextSchedule []map[string]interface{}  `json:"next_schedule"`
         
    }
    
    // PromotionAuthor used by Cart
    type PromotionAuthor struct {

        
            ModifiedBy string  `json:"modified_by"`
            CreatedBy string  `json:"created_by"`
         
    }
    
    // Ownership1 used by Cart
    type Ownership1 struct {

        
            PayableCategory string  `json:"payable_category"`
            PayableBy string  `json:"payable_by"`
         
    }
    
    // PromotionListItem used by Cart
    type PromotionListItem struct {

        
            Restrictions Restrictions1  `json:"restrictions"`
            DisplayMeta DisplayMeta1  `json:"display_meta"`
            Visiblility Visibility  `json:"visiblility"`
            ApplyExclusive string  `json:"apply_exclusive"`
            ApplicationID string  `json:"application_id"`
            PostOrderAction PromotionAction  `json:"post_order_action"`
            ApplyAllDiscount bool  `json:"apply_all_discount"`
            Stackable bool  `json:"stackable"`
            DateMeta PromotionDateMeta  `json:"date_meta"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            DiscountRules []DiscountRule  `json:"discount_rules"`
            Currency string  `json:"currency"`
            CalculateOn string  `json:"calculate_on"`
            PromotionType string  `json:"promotion_type"`
            BuyRules map[string]ItemCriteria  `json:"buy_rules"`
            Schedule PromotionSchedule  `json:"_schedule"`
            Mode string  `json:"mode"`
            ApplyPriority float64  `json:"apply_priority"`
            Author PromotionAuthor  `json:"author"`
            Code string  `json:"code"`
            Ownership Ownership1  `json:"ownership"`
            PromoGroup string  `json:"promo_group"`
         
    }
    
    // PromotionsResponse used by Cart
    type PromotionsResponse struct {

        
            Items PromotionListItem  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // PromotionAdd used by Cart
    type PromotionAdd struct {

        
            Restrictions Restrictions1  `json:"restrictions"`
            DisplayMeta DisplayMeta1  `json:"display_meta"`
            Visiblility Visibility  `json:"visiblility"`
            ApplyExclusive string  `json:"apply_exclusive"`
            ApplicationID string  `json:"application_id"`
            PostOrderAction PromotionAction  `json:"post_order_action"`
            ApplyAllDiscount bool  `json:"apply_all_discount"`
            Stackable bool  `json:"stackable"`
            DateMeta PromotionDateMeta  `json:"date_meta"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            DiscountRules []DiscountRule  `json:"discount_rules"`
            Currency string  `json:"currency"`
            CalculateOn string  `json:"calculate_on"`
            PromotionType string  `json:"promotion_type"`
            BuyRules map[string]ItemCriteria  `json:"buy_rules"`
            Schedule PromotionSchedule  `json:"_schedule"`
            Mode string  `json:"mode"`
            ApplyPriority float64  `json:"apply_priority"`
            Author PromotionAuthor  `json:"author"`
            Code string  `json:"code"`
            Ownership Ownership1  `json:"ownership"`
            PromoGroup string  `json:"promo_group"`
         
    }
    
    // PromotionUpdate used by Cart
    type PromotionUpdate struct {

        
            Restrictions Restrictions1  `json:"restrictions"`
            DisplayMeta DisplayMeta1  `json:"display_meta"`
            Visiblility Visibility  `json:"visiblility"`
            ApplyExclusive string  `json:"apply_exclusive"`
            ApplicationID string  `json:"application_id"`
            PostOrderAction PromotionAction  `json:"post_order_action"`
            ApplyAllDiscount bool  `json:"apply_all_discount"`
            Stackable bool  `json:"stackable"`
            DateMeta PromotionDateMeta  `json:"date_meta"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            DiscountRules []DiscountRule  `json:"discount_rules"`
            Currency string  `json:"currency"`
            CalculateOn string  `json:"calculate_on"`
            PromotionType string  `json:"promotion_type"`
            BuyRules map[string]ItemCriteria  `json:"buy_rules"`
            Schedule PromotionSchedule  `json:"_schedule"`
            Mode string  `json:"mode"`
            ApplyPriority float64  `json:"apply_priority"`
            Author PromotionAuthor  `json:"author"`
            Code string  `json:"code"`
            Ownership Ownership1  `json:"ownership"`
            PromoGroup string  `json:"promo_group"`
         
    }
    
    // PromotionPartialUpdate used by Cart
    type PromotionPartialUpdate struct {

        
            Archive bool  `json:"archive"`
            Schedule PromotionSchedule  `json:"schedule"`
         
    }
    
    // ActivePromosResponse used by Cart
    type ActivePromosResponse struct {

        
            ModifiedOn string  `json:"modified_on"`
            Example string  `json:"example"`
            EntityType string  `json:"entity_type"`
            Description string  `json:"description"`
            EntitySlug string  `json:"entity_slug"`
            Subtitle string  `json:"subtitle"`
            Type string  `json:"type"`
            Title string  `json:"title"`
            CreatedOn string  `json:"created_on"`
            IsHidden bool  `json:"is_hidden"`
         
    }
    
    // CartItem used by Cart
    type CartItem struct {

        
            ProductID string  `json:"product_id"`
            Quantity float64  `json:"quantity"`
            Size string  `json:"size"`
         
    }
    
    // OpenapiCartDetailsRequest used by Cart
    type OpenapiCartDetailsRequest struct {

        
            CartItems CartItem  `json:"cart_items"`
         
    }
    
    // RawBreakup used by Cart
    type RawBreakup struct {

        
            GiftCard float64  `json:"gift_card"`
            Discount float64  `json:"discount"`
            GstCharges float64  `json:"gst_charges"`
            MrpTotal float64  `json:"mrp_total"`
            FyndCash float64  `json:"fynd_cash"`
            DeliveryCharge float64  `json:"delivery_charge"`
            YouSaved float64  `json:"you_saved"`
            ConvenienceFee float64  `json:"convenience_fee"`
            CodCharge float64  `json:"cod_charge"`
            Subtotal float64  `json:"subtotal"`
            Total float64  `json:"total"`
            Coupon float64  `json:"coupon"`
            Vog float64  `json:"vog"`
         
    }
    
    // CouponBreakup used by Cart
    type CouponBreakup struct {

        
            MaxDiscountValue float64  `json:"max_discount_value"`
            CouponValue float64  `json:"coupon_value"`
            Description string  `json:"description"`
            MinimumCartValue float64  `json:"minimum_cart_value"`
            Message string  `json:"message"`
            CouponType string  `json:"coupon_type"`
            IsApplied bool  `json:"is_applied"`
            Type string  `json:"type"`
            Title string  `json:"title"`
            Code string  `json:"code"`
            SubTitle string  `json:"sub_title"`
            Value float64  `json:"value"`
            UID string  `json:"uid"`
         
    }
    
    // LoyaltyPoints used by Cart
    type LoyaltyPoints struct {

        
            Total float64  `json:"total"`
            IsApplied bool  `json:"is_applied"`
            Description string  `json:"description"`
            Applicable float64  `json:"applicable"`
         
    }
    
    // DisplayBreakup used by Cart
    type DisplayBreakup struct {

        
            CurrencySymbol string  `json:"currency_symbol"`
            Message []string  `json:"message"`
            Display string  `json:"display"`
            Key string  `json:"key"`
            Value float64  `json:"value"`
            CurrencyCode string  `json:"currency_code"`
         
    }
    
    // CartBreakup used by Cart
    type CartBreakup struct {

        
            Raw RawBreakup  `json:"raw"`
            Coupon CouponBreakup  `json:"coupon"`
            LoyaltyPoints LoyaltyPoints  `json:"loyalty_points"`
            Display []DisplayBreakup  `json:"display"`
         
    }
    
    // ProductAvailability used by Cart
    type ProductAvailability struct {

        
            IsValid bool  `json:"is_valid"`
            OutOfStock bool  `json:"out_of_stock"`
            Sizes []string  `json:"sizes"`
            OtherStoreQuantity float64  `json:"other_store_quantity"`
            Deliverable bool  `json:"deliverable"`
         
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
    
    // BaseInfo used by Cart
    type BaseInfo struct {

        
            Name string  `json:"name"`
            UID float64  `json:"uid"`
         
    }
    
    // ProductImage used by Cart
    type ProductImage struct {

        
            URL string  `json:"url"`
            AspectRatio string  `json:"aspect_ratio"`
            SecureURL string  `json:"secure_url"`
         
    }
    
    // Tags used by Cart
    type Tags struct {

        
            Tags map[string]interface{}  `json:"tags"`
         
    }
    
    // CategoryInfo used by Cart
    type CategoryInfo struct {

        
            Name string  `json:"name"`
            UID float64  `json:"uid"`
         
    }
    
    // CartProduct used by Cart
    type CartProduct struct {

        
            Slug string  `json:"slug"`
            Tags []string  `json:"tags"`
            Action ProductAction  `json:"action"`
            Type string  `json:"type"`
            Brand BaseInfo  `json:"brand"`
            Images []ProductImage  `json:"images"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Name string  `json:"name"`
            TeaserTag Tags  `json:"teaser_tag"`
            ItemCode string  `json:"item_code"`
            Categories []CategoryInfo  `json:"categories"`
            UID float64  `json:"uid"`
         
    }
    
    // PromoMeta used by Cart
    type PromoMeta struct {

        
            Message string  `json:"message"`
         
    }
    
    // ProductPrice used by Cart
    type ProductPrice struct {

        
            Selling float64  `json:"selling"`
            CurrencySymbol string  `json:"currency_symbol"`
            Effective float64  `json:"effective"`
            Marked float64  `json:"marked"`
            AddOn float64  `json:"add_on"`
            CurrencyCode string  `json:"currency_code"`
         
    }
    
    // ProductPriceInfo used by Cart
    type ProductPriceInfo struct {

        
            Converted ProductPrice  `json:"converted"`
            Base ProductPrice  `json:"base"`
         
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

        
            GiftCard map[string]interface{}  `json:"gift_card"`
            CartItemMeta map[string]interface{}  `json:"cart_item_meta"`
            Identifier map[string]interface{}  `json:"identifier"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            Price ArticlePriceInfo  `json:"price"`
            Seller BaseInfo  `json:"seller"`
            Type string  `json:"type"`
            IsGiftVisible bool  `json:"is_gift_visible"`
            Store BaseInfo  `json:"store"`
            SellerIdentifier string  `json:"seller_identifier"`
            Size string  `json:"size"`
            ProductGroupTags []string  `json:"product_group_tags"`
            Quantity float64  `json:"quantity"`
            UID string  `json:"uid"`
         
    }
    
    // FreeGiftItem used by Cart
    type FreeGiftItem struct {

        
            ItemName string  `json:"item_name"`
            ItemPriceDetails map[string]interface{}  `json:"item_price_details"`
            ItemID float64  `json:"item_id"`
            ItemSlug string  `json:"item_slug"`
            ItemBrandName string  `json:"item_brand_name"`
            ItemImagesURL []string  `json:"item_images_url"`
         
    }
    
    // AppliedFreeArticles used by Cart
    type AppliedFreeArticles struct {

        
            FreeGiftItemDetails FreeGiftItem  `json:"free_gift_item_details"`
            Quantity float64  `json:"quantity"`
            ArticleID string  `json:"article_id"`
            ParentItemIdentifier string  `json:"parent_item_identifier"`
         
    }
    
    // DiscountRulesApp used by Cart
    type DiscountRulesApp struct {

        
            ItemCriteria map[string]interface{}  `json:"item_criteria"`
            Offer map[string]interface{}  `json:"offer"`
            RawOffer map[string]interface{}  `json:"raw_offer"`
            MatchedBuyRules []string  `json:"matched_buy_rules"`
         
    }
    
    // BuyRules used by Cart
    type BuyRules struct {

        
            CartConditions map[string]interface{}  `json:"cart_conditions"`
            ItemCriteria map[string]interface{}  `json:"item_criteria"`
         
    }
    
    // AppliedPromotion used by Cart
    type AppliedPromotion struct {

        
            PromotionName string  `json:"promotion_name"`
            PromotionType string  `json:"promotion_type"`
            ArticleQuantity float64  `json:"article_quantity"`
            MrpPromotion bool  `json:"mrp_promotion"`
            OfferText string  `json:"offer_text"`
            AppliedFreeArticles []AppliedFreeArticles  `json:"applied_free_articles"`
            DiscountRules []DiscountRulesApp  `json:"discount_rules"`
            PromoID string  `json:"promo_id"`
            BuyRules []BuyRules  `json:"buy_rules"`
            PromotionGroup string  `json:"promotion_group"`
            Amount float64  `json:"amount"`
         
    }
    
    // CouponDetails used by Cart
    type CouponDetails struct {

        
            Code string  `json:"code"`
            DiscountSingleQuantity float64  `json:"discount_single_quantity"`
            DiscountTotalQuantity float64  `json:"discount_total_quantity"`
         
    }
    
    // CartProductIdentifer used by Cart
    type CartProductIdentifer struct {

        
            Identifier string  `json:"identifier"`
         
    }
    
    // CartProductInfo used by Cart
    type CartProductInfo struct {

        
            CouponMessage string  `json:"coupon_message"`
            Availability ProductAvailability  `json:"availability"`
            Product CartProduct  `json:"product"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            Discount string  `json:"discount"`
            PromoMeta PromoMeta  `json:"promo_meta"`
            Price ProductPriceInfo  `json:"price"`
            Message string  `json:"message"`
            IsSet bool  `json:"is_set"`
            BulkOffer map[string]interface{}  `json:"bulk_offer"`
            Article ProductArticle  `json:"article"`
            Key string  `json:"key"`
            PricePerUnit ProductPriceInfo  `json:"price_per_unit"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            PromotionsApplied []AppliedPromotion  `json:"promotions_applied"`
            Coupon CouponDetails  `json:"coupon"`
            Quantity float64  `json:"quantity"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
         
    }
    
    // OpenapiCartDetailsResponse used by Cart
    type OpenapiCartDetailsResponse struct {

        
            IsValid bool  `json:"is_valid"`
            Message string  `json:"message"`
            BreakupValues CartBreakup  `json:"breakup_values"`
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

        
            Phone float64  `json:"phone"`
            City string  `json:"city"`
            CountryCode string  `json:"country_code"`
            Meta map[string]interface{}  `json:"meta"`
            AddressType string  `json:"address_type"`
            AreaCode string  `json:"area_code"`
            Country string  `json:"country"`
            Landmark string  `json:"landmark"`
            Address string  `json:"address"`
            Name string  `json:"name"`
            State string  `json:"state"`
            Pincode float64  `json:"pincode"`
            Email string  `json:"email"`
            Area string  `json:"area"`
            AreaCodeSlug string  `json:"area_code_slug"`
         
    }
    
    // OpenApiCartServiceabilityRequest used by Cart
    type OpenApiCartServiceabilityRequest struct {

        
            ShippingAddress ShippingAddress  `json:"shipping_address"`
            CartItems CartItem  `json:"cart_items"`
         
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
    
    // OpenApiCartServiceabilityResponse used by Cart
    type OpenApiCartServiceabilityResponse struct {

        
            IsValid bool  `json:"is_valid"`
            Items []CartProductInfo  `json:"items"`
            Message string  `json:"message"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            BreakupValues CartBreakup  `json:"breakup_values"`
         
    }
    
    // MultiTenderPaymentMeta used by Cart
    type MultiTenderPaymentMeta struct {

        
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            OrderID string  `json:"order_id"`
            CurrentStatus string  `json:"current_status"`
            PaymentGateway string  `json:"payment_gateway"`
            PaymentID string  `json:"payment_id"`
         
    }
    
    // MultiTenderPaymentMethod used by Cart
    type MultiTenderPaymentMethod struct {

        
            Name string  `json:"name"`
            Amount float64  `json:"amount"`
            Meta MultiTenderPaymentMeta  `json:"meta"`
            Mode string  `json:"mode"`
         
    }
    
    // OpenApiFiles used by Cart
    type OpenApiFiles struct {

        
            Key string  `json:"key"`
            Values []string  `json:"values"`
         
    }
    
    // CartItemMeta used by Cart
    type CartItemMeta struct {

        
            GroupID string  `json:"group_id"`
            PrimaryItem bool  `json:"primary_item"`
         
    }
    
    // OpenApiOrderItem used by Cart
    type OpenApiOrderItem struct {

        
            AmountPaid float64  `json:"amount_paid"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            Discount float64  `json:"discount"`
            Quantity float64  `json:"quantity"`
            CodCharges float64  `json:"cod_charges"`
            PriceEffective float64  `json:"price_effective"`
            EmployeeDiscount float64  `json:"employee_discount"`
            Meta CartItemMeta  `json:"meta"`
            PriceMarked float64  `json:"price_marked"`
            ProductID float64  `json:"product_id"`
            Files []OpenApiFiles  `json:"files"`
            Size string  `json:"size"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            PaymentMethods []MultiTenderPaymentMethod  `json:"payment_methods"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            DeliveryCharges float64  `json:"delivery_charges"`
            CashbackApplied float64  `json:"cashback_applied"`
         
    }
    
    // OpenApiPlatformCheckoutReq used by Cart
    type OpenApiPlatformCheckoutReq struct {

        
            Gstin string  `json:"gstin"`
            CouponValue float64  `json:"coupon_value"`
            PaymentMode string  `json:"payment_mode"`
            PaymentMethods []MultiTenderPaymentMethod  `json:"payment_methods"`
            CouponCode string  `json:"coupon_code"`
            CartValue float64  `json:"cart_value"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            Comment string  `json:"comment"`
            Coupon string  `json:"coupon"`
            CashbackApplied float64  `json:"cashback_applied"`
            CurrencyCode string  `json:"currency_code"`
            BillingAddress ShippingAddress  `json:"billing_address"`
            CodCharges float64  `json:"cod_charges"`
            EmployeeDiscount map[string]interface{}  `json:"employee_discount"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            OrderID string  `json:"order_id"`
            ShippingAddress ShippingAddress  `json:"shipping_address"`
            DeliveryCharges float64  `json:"delivery_charges"`
            Files []OpenApiFiles  `json:"files"`
            CartItems []OpenApiOrderItem  `json:"cart_items"`
         
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

        
            Gstin string  `json:"gstin"`
            Discount float64  `json:"discount"`
            PickUpCustomerDetails map[string]interface{}  `json:"pick_up_customer_details"`
            IsDefault bool  `json:"is_default"`
            Articles []map[string]interface{}  `json:"articles"`
            PaymentMode string  `json:"payment_mode"`
            PaymentMethods []map[string]interface{}  `json:"payment_methods"`
            IsArchive bool  `json:"is_archive"`
            Promotion map[string]interface{}  `json:"promotion"`
            FyndCredits map[string]interface{}  `json:"fynd_credits"`
            Comment string  `json:"comment"`
            UserID string  `json:"user_id"`
            CodCharges map[string]interface{}  `json:"cod_charges"`
            Meta map[string]interface{}  `json:"meta"`
            AppID string  `json:"app_id"`
            Payments map[string]interface{}  `json:"payments"`
            CheckoutMode string  `json:"checkout_mode"`
            DeliveryCharges map[string]interface{}  `json:"delivery_charges"`
            Shipments []map[string]interface{}  `json:"shipments"`
            FcIndexMap []float64  `json:"fc_index_map"`
            BulkCouponDiscount float64  `json:"bulk_coupon_discount"`
            LastModified string  `json:"last_modified"`
            CartValue float64  `json:"cart_value"`
            ID string  `json:"_id"`
            Coupon map[string]interface{}  `json:"coupon"`
            IsActive bool  `json:"is_active"`
            Cashback map[string]interface{}  `json:"cashback"`
            MergeQty bool  `json:"merge_qty"`
            OrderID string  `json:"order_id"`
            ExpireAt string  `json:"expire_at"`
            CreatedOn string  `json:"created_on"`
            BuyNow bool  `json:"buy_now"`
            UID float64  `json:"uid"`
         
    }
    
    // AbandonedCartResponse used by Cart
    type AbandonedCartResponse struct {

        
            Items []AbandonedCart  `json:"items"`
            Message string  `json:"message"`
            Success bool  `json:"success"`
            Result map[string]interface{}  `json:"result"`
            Page Page  `json:"page"`
         
    }
    
    // AddProductCart used by Cart
    type AddProductCart struct {

        
            Pos bool  `json:"pos"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            Display string  `json:"display"`
            ArticleID string  `json:"article_id"`
            ItemID float64  `json:"item_id"`
            ItemSize string  `json:"item_size"`
            StoreID float64  `json:"store_id"`
            ProductGroupTags []string  `json:"product_group_tags"`
            Quantity float64  `json:"quantity"`
            SellerID float64  `json:"seller_id"`
            ArticleAssignment map[string]interface{}  `json:"article_assignment"`
         
    }
    
    // AddCartRequest used by Cart
    type AddCartRequest struct {

        
            NewCart bool  `json:"new_cart"`
            Items []AddProductCart  `json:"items"`
         
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

        
            Gstin string  `json:"gstin"`
            IsValid bool  `json:"is_valid"`
            Items []CartProductInfo  `json:"items"`
            CheckoutMode string  `json:"checkout_mode"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            Message string  `json:"message"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            ID string  `json:"id"`
            Comment string  `json:"comment"`
            CouponText string  `json:"coupon_text"`
            LastModified string  `json:"last_modified"`
            PanConfig map[string]interface{}  `json:"pan_config"`
            BuyNow bool  `json:"buy_now"`
            Currency CartCurrency  `json:"currency"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            PanNo string  `json:"pan_no"`
         
    }
    
    // AddCartDetailResponse used by Cart
    type AddCartDetailResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
            Cart CartDetailResponse  `json:"cart"`
            Partial bool  `json:"partial"`
         
    }
    
    // UpdateProductCart used by Cart
    type UpdateProductCart struct {

        
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            ArticleID string  `json:"article_id"`
            ItemID float64  `json:"item_id"`
            ItemSize string  `json:"item_size"`
            Quantity float64  `json:"quantity"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
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
            Message string  `json:"message"`
            Cart CartDetailResponse  `json:"cart"`
         
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

        
            Token string  `json:"token"`
            User map[string]interface{}  `json:"user"`
            Meta map[string]interface{}  `json:"meta"`
            Source map[string]interface{}  `json:"source"`
            CreatedOn string  `json:"created_on"`
         
    }
    
    // SharedCart used by Cart
    type SharedCart struct {

        
            Gstin string  `json:"gstin"`
            IsValid bool  `json:"is_valid"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            Comment string  `json:"comment"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            CheckoutMode string  `json:"checkout_mode"`
            Currency CartCurrency  `json:"currency"`
            CartID float64  `json:"cart_id"`
            SharedCartDetails SharedCartDetails  `json:"shared_cart_details"`
            LastModified string  `json:"last_modified"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            Items []CartProductInfo  `json:"items"`
            Message string  `json:"message"`
            ID string  `json:"id"`
            CouponText string  `json:"coupon_text"`
            BuyNow bool  `json:"buy_now"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            UID string  `json:"uid"`
         
    }
    
    // SharedCartResponse used by Cart
    type SharedCartResponse struct {

        
            Error string  `json:"error"`
            Cart SharedCart  `json:"cart"`
         
    }
    
    // CartList used by Cart
    type CartList struct {

        
            CartID string  `json:"cart_id"`
            UserID string  `json:"user_id"`
            PickUpCustomerDetails map[string]interface{}  `json:"pick_up_customer_details"`
            ItemCounts float64  `json:"item_counts"`
            CreatedOn string  `json:"created_on"`
            CartValue float64  `json:"cart_value"`
         
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

        
            ModifiedOn string  `json:"modified_on"`
            LastName string  `json:"last_name"`
            Gender string  `json:"gender"`
            ExternalID string  `json:"external_id"`
            Mobile string  `json:"mobile"`
            ID string  `json:"_id"`
            CreatedAt string  `json:"created_at"`
            UID string  `json:"uid"`
            FirstName string  `json:"first_name"`
         
    }
    
    // UserCartMappingResponse used by Cart
    type UserCartMappingResponse struct {

        
            Gstin string  `json:"gstin"`
            IsValid bool  `json:"is_valid"`
            User UserInfo  `json:"user"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            PanConfig map[string]interface{}  `json:"pan_config"`
            Comment string  `json:"comment"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            CheckoutMode string  `json:"checkout_mode"`
            Currency CartCurrency  `json:"currency"`
            LastModified string  `json:"last_modified"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            Items []CartProductInfo  `json:"items"`
            Message string  `json:"message"`
            ID string  `json:"id"`
            CouponText string  `json:"coupon_text"`
            BuyNow bool  `json:"buy_now"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            PanNo string  `json:"pan_no"`
         
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

        
            MaxDiscountValue float64  `json:"max_discount_value"`
            CouponValue float64  `json:"coupon_value"`
            MinimumCartValue float64  `json:"minimum_cart_value"`
            Description string  `json:"description"`
            Message string  `json:"message"`
            CouponType string  `json:"coupon_type"`
            IsApplied bool  `json:"is_applied"`
            Title string  `json:"title"`
            SubTitle string  `json:"sub_title"`
            IsApplicable bool  `json:"is_applicable"`
            CouponCode string  `json:"coupon_code"`
            ExpiresOn string  `json:"expires_on"`
         
    }
    
    // PageCoupon used by Cart
    type PageCoupon struct {

        
            HasPrevious bool  `json:"has_previous"`
            Total float64  `json:"total"`
            TotalItemCount float64  `json:"total_item_count"`
            HasNext bool  `json:"has_next"`
            Current float64  `json:"current"`
         
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

        
            GeoLocation GeoLocation  `json:"geo_location"`
            Area string  `json:"area"`
            UserID string  `json:"user_id"`
            Tags []string  `json:"tags"`
            Meta map[string]interface{}  `json:"meta"`
            Landmark string  `json:"landmark"`
            Address string  `json:"address"`
            CheckoutMode string  `json:"checkout_mode"`
            AreaCodeSlug string  `json:"area_code_slug"`
            Phone string  `json:"phone"`
            CartID string  `json:"cart_id"`
            Country string  `json:"country"`
            Email string  `json:"email"`
            IsActive bool  `json:"is_active"`
            IsDefaultAddress bool  `json:"is_default_address"`
            City string  `json:"city"`
            AddressType string  `json:"address_type"`
            AreaCode string  `json:"area_code"`
            GoogleMapPoint map[string]interface{}  `json:"google_map_point"`
            ID string  `json:"id"`
            Name string  `json:"name"`
            State string  `json:"state"`
            CountryCode string  `json:"country_code"`
            CreatedByUserID string  `json:"created_by_user_id"`
         
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

        
            IsDefaultAddress bool  `json:"is_default_address"`
            IsUpdated bool  `json:"is_updated"`
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

        
            CartID string  `json:"cart_id"`
            UserID string  `json:"user_id"`
            BillingAddressID string  `json:"billing_address_id"`
            ID string  `json:"id"`
            CheckoutMode string  `json:"checkout_mode"`
         
    }
    
    // ShipmentResponse used by Cart
    type ShipmentResponse struct {

        
            Shipments float64  `json:"shipments"`
            Items []CartProductInfo  `json:"items"`
            DpOptions map[string]interface{}  `json:"dp_options"`
            Promise ShipmentPromise  `json:"promise"`
            OrderType string  `json:"order_type"`
            BoxType string  `json:"box_type"`
            FulfillmentType string  `json:"fulfillment_type"`
            DpID string  `json:"dp_id"`
            ShipmentType string  `json:"shipment_type"`
            FulfillmentID float64  `json:"fulfillment_id"`
         
    }
    
    // CartShipmentsResponse used by Cart
    type CartShipmentsResponse struct {

        
            Gstin string  `json:"gstin"`
            IsValid bool  `json:"is_valid"`
            CartID float64  `json:"cart_id"`
            Shipments []ShipmentResponse  `json:"shipments"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            LastModified string  `json:"last_modified"`
            Comment string  `json:"comment"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            Message string  `json:"message"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            ID string  `json:"id"`
            CouponText string  `json:"coupon_text"`
            CheckoutMode string  `json:"checkout_mode"`
            BuyNow bool  `json:"buy_now"`
            Currency CartCurrency  `json:"currency"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            UID string  `json:"uid"`
            Error bool  `json:"error"`
         
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
            GiftDetails map[string]interface{}  `json:"gift_details"`
            PickUpCustomerDetails map[string]interface{}  `json:"pick_up_customer_details"`
            CheckoutMode string  `json:"checkout_mode"`
            Comment string  `json:"comment"`
            PanNo string  `json:"pan_no"`
         
    }
    
    // CartMetaResponse used by Cart
    type CartMetaResponse struct {

        
            Message string  `json:"message"`
         
    }
    
    // CartMetaMissingResponse used by Cart
    type CartMetaMissingResponse struct {

        
            Errors []string  `json:"errors"`
         
    }
    
    // Files used by Cart
    type Files struct {

        
            Key string  `json:"key"`
            Values []string  `json:"values"`
         
    }
    
    // StaffCheckout used by Cart
    type StaffCheckout struct {

        
            LastName string  `json:"last_name"`
            User string  `json:"user"`
            ID string  `json:"_id"`
            EmployeeCode string  `json:"employee_code"`
            FirstName string  `json:"first_name"`
         
    }
    
    // PlatformCartCheckoutDetailRequest used by Cart
    type PlatformCartCheckoutDetailRequest struct {

        
            DeviceID string  `json:"device_id"`
            OrderType string  `json:"order_type"`
            PickAtStoreUID float64  `json:"pick_at_store_uid"`
            PaymentMode string  `json:"payment_mode"`
            AddressID string  `json:"address_id"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            UserID string  `json:"user_id"`
            Meta map[string]interface{}  `json:"meta"`
            CheckoutMode string  `json:"checkout_mode"`
            Files []Files  `json:"files"`
            EmployeeCode string  `json:"employee_code"`
            Pos bool  `json:"pos"`
            PaymentAutoConfirm bool  `json:"payment_auto_confirm"`
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
            Staff StaffCheckout  `json:"staff"`
            Aggregator string  `json:"aggregator"`
            OrderingStore float64  `json:"ordering_store"`
            PaymentParams map[string]interface{}  `json:"payment_params"`
            BillingAddressID string  `json:"billing_address_id"`
            PaymentIdentifier string  `json:"payment_identifier"`
            ID string  `json:"id"`
            CallbackURL string  `json:"callback_url"`
            MerchantCode string  `json:"merchant_code"`
         
    }
    
    // CheckCart used by Cart
    type CheckCart struct {

        
            Gstin string  `json:"gstin"`
            IsValid bool  `json:"is_valid"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            DeliveryChargeOrderValue float64  `json:"delivery_charge_order_value"`
            StoreEmps []map[string]interface{}  `json:"store_emps"`
            Comment string  `json:"comment"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            CodCharges float64  `json:"cod_charges"`
            StoreCode string  `json:"store_code"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            ErrorMessage string  `json:"error_message"`
            CheckoutMode string  `json:"checkout_mode"`
            UserType string  `json:"user_type"`
            Currency CartCurrency  `json:"currency"`
            DeliveryCharges float64  `json:"delivery_charges"`
            CartID float64  `json:"cart_id"`
            LastModified string  `json:"last_modified"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            Items []CartProductInfo  `json:"items"`
            CodAvailable bool  `json:"cod_available"`
            Message string  `json:"message"`
            CodMessage string  `json:"cod_message"`
            Success bool  `json:"success"`
            ID string  `json:"id"`
            OrderID string  `json:"order_id"`
            CouponText string  `json:"coupon_text"`
            BuyNow bool  `json:"buy_now"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            UID string  `json:"uid"`
         
    }
    
    // CartCheckoutResponse used by Cart
    type CartCheckoutResponse struct {

        
            Message string  `json:"message"`
            Data map[string]interface{}  `json:"data"`
            Success bool  `json:"success"`
            OrderID string  `json:"order_id"`
            CallbackURL string  `json:"callback_url"`
            PaymentConfirmURL string  `json:"payment_confirm_url"`
            AppInterceptURL string  `json:"app_intercept_url"`
            Cart CheckCart  `json:"cart"`
         
    }
    
    // CartDeliveryModesResponse used by Cart
    type CartDeliveryModesResponse struct {

        
            PickupStores []float64  `json:"pickup_stores"`
            AvailableModes []string  `json:"available_modes"`
         
    }
    
    // PickupStoreDetail used by Cart
    type PickupStoreDetail struct {

        
            Phone string  `json:"phone"`
            City string  `json:"city"`
            StoreCode string  `json:"store_code"`
            AddressType string  `json:"address_type"`
            AreaCode string  `json:"area_code"`
            Country string  `json:"country"`
            Landmark string  `json:"landmark"`
            ID float64  `json:"id"`
            Address string  `json:"address"`
            Name string  `json:"name"`
            State string  `json:"state"`
            Pincode float64  `json:"pincode"`
            Email string  `json:"email"`
            Area string  `json:"area"`
            AreaCodeSlug string  `json:"area_code_slug"`
            UID float64  `json:"uid"`
         
    }
    
    // StoreDetailsResponse used by Cart
    type StoreDetailsResponse struct {

        
            Items []PickupStoreDetail  `json:"items"`
         
    }
    
    // UpdateCartPaymentRequest used by Cart
    type UpdateCartPaymentRequest struct {

        
            PaymentIdentifier string  `json:"payment_identifier"`
            ID string  `json:"id"`
            AggregatorName string  `json:"aggregator_name"`
            MerchantCode string  `json:"merchant_code"`
            PaymentMode string  `json:"payment_mode"`
            AddressID string  `json:"address_id"`
         
    }
    
    // CouponValidity used by Cart
    type CouponValidity struct {

        
            Valid bool  `json:"valid"`
            Discount float64  `json:"discount"`
            Title string  `json:"title"`
            Code string  `json:"code"`
            DisplayMessageEn string  `json:"display_message_en"`
         
    }
    
    // PaymentCouponValidate used by Cart
    type PaymentCouponValidate struct {

        
            Message string  `json:"message"`
            CouponValidity CouponValidity  `json:"coupon_validity"`
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
    
    // APIError used by Partner
    type APIError struct {

        
            Code string  `json:"code"`
            Message string  `json:"message"`
            Info string  `json:"info"`
            RequestID string  `json:"request_id"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // RemoveProxyResponse used by Partner
    type RemoveProxyResponse struct {

        
            Message string  `json:"message"`
            Data map[string]interface{}  `json:"data"`
         
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

        
            Value string  `json:"value"`
            Type string  `json:"type"`
            Message string  `json:"message"`
         
    }
    
    // ApplicationServiceabilityConfig used by Serviceability
    type ApplicationServiceabilityConfig struct {

        
            ChannelID string  `json:"channel_id"`
            ServiceabilityType string  `json:"serviceability_type"`
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
    
    // EntityRegionView_Error used by Serviceability
    type EntityRegionView_Error struct {

        
            Value string  `json:"value"`
            Type string  `json:"type"`
            Message string  `json:"message"`
         
    }
    
    // EntityRegionView_Items used by Serviceability
    type EntityRegionView_Items struct {

        
            SubType string  `json:"sub_type"`
            Name string  `json:"name"`
            UID string  `json:"uid"`
         
    }
    
    // EntityRegionView_page used by Serviceability
    type EntityRegionView_page struct {

        
            Current float64  `json:"current"`
            Type string  `json:"type"`
            ItemTotal float64  `json:"item_total"`
            Size float64  `json:"size"`
            HasNext bool  `json:"has_next"`
         
    }
    
    // EntityRegionView_Response used by Serviceability
    type EntityRegionView_Response struct {

        
            Error EntityRegionView_Error  `json:"error"`
            Data []EntityRegionView_Items  `json:"data"`
            Success bool  `json:"success"`
            Page EntityRegionView_page  `json:"page"`
         
    }
    
    // ListViewSummary used by Serviceability
    type ListViewSummary struct {

        
            TotalPincodesServed float64  `json:"total_pincodes_served"`
            TotalZones float64  `json:"total_zones"`
            TotalActiveZones float64  `json:"total_active_zones"`
         
    }
    
    // ZoneDataItem used by Serviceability
    type ZoneDataItem struct {

        
            Current float64  `json:"current"`
            ItemTotal float64  `json:"item_total"`
            Size float64  `json:"size"`
            HasNext bool  `json:"has_next"`
            Type string  `json:"type"`
         
    }
    
    // ListViewProduct used by Serviceability
    type ListViewProduct struct {

        
            Type string  `json:"type"`
            Count float64  `json:"count"`
         
    }
    
    // ListViewChannels used by Serviceability
    type ListViewChannels struct {

        
            ChannelID string  `json:"channel_id"`
            ChannelType string  `json:"channel_type"`
         
    }
    
    // ListViewItems used by Serviceability
    type ListViewItems struct {

        
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            CompanyID float64  `json:"company_id"`
            Slug string  `json:"slug"`
            StoresCount float64  `json:"stores_count"`
            Product ListViewProduct  `json:"product"`
            ZoneID string  `json:"zone_id"`
            PincodesCount float64  `json:"pincodes_count"`
            Channels ListViewChannels  `json:"channels"`
         
    }
    
    // ListViewResponse used by Serviceability
    type ListViewResponse struct {

        
            Summary []ListViewSummary  `json:"summary"`
            Page []ZoneDataItem  `json:"page"`
            Items []ListViewItems  `json:"items"`
         
    }
    
    // CompanyStoreView_PageItems used by Serviceability
    type CompanyStoreView_PageItems struct {

        
            Current float64  `json:"current"`
            Type string  `json:"type"`
            ItemTotal float64  `json:"item_total"`
            Size float64  `json:"size"`
            HasNext bool  `json:"has_next"`
         
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

        
            Tags []string  `json:"tags"`
            Type string  `json:"type"`
         
    }
    
    // ZoneMappingType used by Serviceability
    type ZoneMappingType struct {

        
            State []string  `json:"state"`
            Country string  `json:"country"`
            Pincode []string  `json:"pincode"`
         
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

        
            Identifier string  `json:"identifier"`
            Data UpdateZoneData  `json:"data"`
         
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

        
            Identifier string  `json:"identifier"`
            Data CreateZoneData  `json:"data"`
         
    }
    
    // ZoneResponse used by Serviceability
    type ZoneResponse struct {

        
            StatusCode float64  `json:"status_code"`
            ZoneID string  `json:"zone_id"`
            Success bool  `json:"success"`
         
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
    
    // PageResponse used by Serviceability
    type PageResponse struct {

        
            Current float64  `json:"current"`
            Type string  `json:"type"`
            ItemTotal float64  `json:"item_total"`
            Size float64  `json:"size"`
            HasNext bool  `json:"has_next"`
         
    }
    
    // OpeningClosing used by Serviceability
    type OpeningClosing struct {

        
            Minute float64  `json:"minute"`
            Hour float64  `json:"hour"`
         
    }
    
    // TimmingResponse used by Serviceability
    type TimmingResponse struct {

        
            Opening OpeningClosing  `json:"opening"`
            Open bool  `json:"open"`
            Weekday string  `json:"weekday"`
            Closing OpeningClosing  `json:"closing"`
         
    }
    
    // ModifiedByResponse used by Serviceability
    type ModifiedByResponse struct {

        
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
         
    }
    
    // Dp used by Serviceability
    type Dp struct {

        
            InternalAccountID string  `json:"internal_account_id"`
            Operations []string  `json:"operations"`
            AssignDpFromSb bool  `json:"assign_dp_from_sb"`
            LmPriority float64  `json:"lm_priority"`
            FmPriority float64  `json:"fm_priority"`
            AreaCode float64  `json:"area_code"`
            ExternalAccountID string  `json:"external_account_id"`
            PaymentMode string  `json:"payment_mode"`
            RvpPriority float64  `json:"rvp_priority"`
            TransportMode string  `json:"transport_mode"`
         
    }
    
    // LogisticsResponse used by Serviceability
    type LogisticsResponse struct {

        
            Override bool  `json:"override"`
            Dp Dp  `json:"dp"`
         
    }
    
    // MobileNo used by Serviceability
    type MobileNo struct {

        
            Number string  `json:"number"`
            CountryCode float64  `json:"country_code"`
         
    }
    
    // ManagerResponse used by Serviceability
    type ManagerResponse struct {

        
            MobileNo MobileNo  `json:"mobile_no"`
            Email string  `json:"email"`
            Name string  `json:"name"`
         
    }
    
    // DocumentsResponse used by Serviceability
    type DocumentsResponse struct {

        
            Value string  `json:"value"`
            LegalName string  `json:"legal_name"`
            Type string  `json:"type"`
            Verified bool  `json:"verified"`
         
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
    
    // WarningsResponse used by Serviceability
    type WarningsResponse struct {

        
            StoreAddress string  `json:"store_address"`
         
    }
    
    // AddressResponse used by Serviceability
    type AddressResponse struct {

        
            Latitude float64  `json:"latitude"`
            Longitude float64  `json:"longitude"`
            Landmark string  `json:"landmark"`
            Country string  `json:"country"`
            Address1 string  `json:"address1"`
            City string  `json:"city"`
            Address2 string  `json:"address2"`
            Pincode float64  `json:"pincode"`
            State string  `json:"state"`
         
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
    
    // CreatedByResponse used by Serviceability
    type CreatedByResponse struct {

        
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
         
    }
    
    // ProductReturnConfigResponse used by Serviceability
    type ProductReturnConfigResponse struct {

        
            OnSameStore bool  `json:"on_same_store"`
         
    }
    
    // ItemResponse used by Serviceability
    type ItemResponse struct {

        
            Name string  `json:"name"`
            ModifiedOn string  `json:"modified_on"`
            DisplayName string  `json:"display_name"`
            Cls string  `json:"_cls"`
            CreatedOn string  `json:"created_on"`
            Timing []TimmingResponse  `json:"timing"`
            VerifiedBy ModifiedByResponse  `json:"verified_by"`
            Logistics LogisticsResponse  `json:"logistics"`
            Stage string  `json:"stage"`
            Code string  `json:"code"`
            Manager ManagerResponse  `json:"manager"`
            Documents []DocumentsResponse  `json:"documents"`
            IntegrationType IntegrationTypeResponse  `json:"integration_type"`
            ContactNumbers []ContactNumberResponse  `json:"contact_numbers"`
            Warnings WarningsResponse  `json:"warnings"`
            StoreType string  `json:"store_type"`
            Address AddressResponse  `json:"address"`
            CompanyID float64  `json:"company_id"`
            VerifiedOn string  `json:"verified_on"`
            GstCredentials GstCredentialsResponse  `json:"gst_credentials"`
            NotificationEmails []string  `json:"notification_emails"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            SubType string  `json:"sub_type"`
            Company float64  `json:"company"`
            CreatedBy CreatedByResponse  `json:"created_by"`
            UID float64  `json:"uid"`
            ModifiedBy ModifiedByResponse  `json:"modified_by"`
            ProductReturnConfig ProductReturnConfigResponse  `json:"product_return_config"`
         
    }
    
    // GetStoresViewResponse used by Serviceability
    type GetStoresViewResponse struct {

        
            Page PageResponse  `json:"page"`
            Items []ItemResponse  `json:"items"`
         
    }
    
