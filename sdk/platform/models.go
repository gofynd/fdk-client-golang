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
            Debug Debug  `json:"debug"`
            HasOldPasswordHash bool  `json:"has_old_password_hash"`
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
            Source TagSourceSchema  `json:"__source"`
         
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
            SubNavigation []SubNavigationReference  `json:"sub_navigation"`
         
    }
    
    // SubNavigationReference used by Content
    type SubNavigationReference struct {

        
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
    

    
    // UnauthenticatedUser used by Billing
    type UnauthenticatedUser struct {

        
            Message string  `json:"message"`
         
    }
    
    // UnauthenticatedApplication used by Billing
    type UnauthenticatedApplication struct {

        
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
    

    
    // PaymentGatewayConfigResponse used by Payment
    type PaymentGatewayConfigResponse struct {

        
            Success bool  `json:"success"`
            DisplayFields []string  `json:"display_fields"`
            Created bool  `json:"created"`
            ExcludedFields []string  `json:"excluded_fields"`
            Aggregators []map[string]interface{}  `json:"aggregators"`
            AppID string  `json:"app_id"`
         
    }
    
    // ErrorCodeDescription used by Payment
    type ErrorCodeDescription struct {

        
            Description string  `json:"description"`
            Code string  `json:"code"`
            Success bool  `json:"success"`
         
    }
    
    // PaymentGatewayConfig used by Payment
    type PaymentGatewayConfig struct {

        
            ConfigType string  `json:"config_type"`
            MerchantSalt string  `json:"merchant_salt"`
            IsActive bool  `json:"is_active"`
            Key string  `json:"key"`
            Secret string  `json:"secret"`
         
    }
    
    // PaymentGatewayConfigRequest used by Payment
    type PaymentGatewayConfigRequest struct {

        
            IsActive bool  `json:"is_active"`
            AggregatorName PaymentGatewayConfig  `json:"aggregator_name"`
            AppID string  `json:"app_id"`
         
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

        
            Logos PaymentModeLogo  `json:"logos"`
            DisplayName string  `json:"display_name"`
            Code string  `json:"code"`
            PackageName string  `json:"package_name"`
         
    }
    
    // PaymentModeList used by Payment
    type PaymentModeList struct {

        
            DisplayPriority float64  `json:"display_priority"`
            Expired bool  `json:"expired"`
            Timeout float64  `json:"timeout"`
            CardType string  `json:"card_type"`
            IntentAppErrorDictList []IntentAppErrorList  `json:"intent_app_error_dict_list"`
            CardBrand string  `json:"card_brand"`
            Code string  `json:"code"`
            IntentFlow bool  `json:"intent_flow"`
            IntentAppErrorList []string  `json:"intent_app_error_list"`
            CardIsin string  `json:"card_isin"`
            LogoURL PaymentModeLogo  `json:"logo_url"`
            DisplayName string  `json:"display_name"`
            AggregatorName string  `json:"aggregator_name"`
            RemainingLimit float64  `json:"remaining_limit"`
            CardIssuer string  `json:"card_issuer"`
            CardNumber string  `json:"card_number"`
            MerchantCode string  `json:"merchant_code"`
            CodLimit float64  `json:"cod_limit"`
            CodLimitPerOrder float64  `json:"cod_limit_per_order"`
            CompliantWithTokenisationGuidelines bool  `json:"compliant_with_tokenisation_guidelines"`
            RetryCount float64  `json:"retry_count"`
            IntentApp []IntentApp  `json:"intent_app"`
            ExpMonth float64  `json:"exp_month"`
            ExpYear float64  `json:"exp_year"`
            CardBrandImage string  `json:"card_brand_image"`
            CardToken string  `json:"card_token"`
            CardID string  `json:"card_id"`
            CardReference string  `json:"card_reference"`
            Name string  `json:"name"`
            CardName string  `json:"card_name"`
            FyndVpa string  `json:"fynd_vpa"`
            Nickname string  `json:"nickname"`
            CardFingerprint string  `json:"card_fingerprint"`
         
    }
    
    // RootPaymentMode used by Payment
    type RootPaymentMode struct {

        
            DisplayPriority float64  `json:"display_priority"`
            SaveCard bool  `json:"save_card"`
            DisplayName string  `json:"display_name"`
            AggregatorName string  `json:"aggregator_name"`
            Name string  `json:"name"`
            AnonymousEnable bool  `json:"anonymous_enable"`
            AddCardEnabled bool  `json:"add_card_enabled"`
            IsPayByCardPl bool  `json:"is_pay_by_card_pl"`
            List []PaymentModeList  `json:"list"`
         
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

        
            IsDefault bool  `json:"is_default"`
            UniqueTransferNo map[string]interface{}  `json:"unique_transfer_no"`
            PayoutsAggregators []map[string]interface{}  `json:"payouts_aggregators"`
            MoreAttributes map[string]interface{}  `json:"more_attributes"`
            IsActive bool  `json:"is_active"`
            Customers map[string]interface{}  `json:"customers"`
            TransferType string  `json:"transfer_type"`
         
    }
    
    // PayoutBankDetails used by Payment
    type PayoutBankDetails struct {

        
            BankName string  `json:"bank_name"`
            City string  `json:"city"`
            State string  `json:"state"`
            AccountNo string  `json:"account_no"`
            Pincode float64  `json:"pincode"`
            AccountType string  `json:"account_type"`
            BranchName string  `json:"branch_name"`
            AccountHolder string  `json:"account_holder"`
            Country string  `json:"country"`
            IfscCode string  `json:"ifsc_code"`
         
    }
    
    // PayoutRequest used by Payment
    type PayoutRequest struct {

        
            Users map[string]interface{}  `json:"users"`
            UniqueExternalID string  `json:"unique_external_id"`
            IsActive bool  `json:"is_active"`
            Aggregator string  `json:"aggregator"`
            BankDetails PayoutBankDetails  `json:"bank_details"`
            TransferType string  `json:"transfer_type"`
         
    }
    
    // PayoutResponse used by Payment
    type PayoutResponse struct {

        
            Users map[string]interface{}  `json:"users"`
            UniqueTransferNo string  `json:"unique_transfer_no"`
            PaymentStatus string  `json:"payment_status"`
            Payouts map[string]interface{}  `json:"payouts"`
            Success bool  `json:"success"`
            IsActive bool  `json:"is_active"`
            Created bool  `json:"created"`
            Aggregator string  `json:"aggregator"`
            BankDetails map[string]interface{}  `json:"bank_details"`
            TransferType string  `json:"transfer_type"`
         
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
            UniqueExternalID string  `json:"unique_external_id"`
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

        
            Data map[string]interface{}  `json:"data"`
            Success bool  `json:"success"`
         
    }
    
    // RefundAccountResponse used by Payment
    type RefundAccountResponse struct {

        
            Data map[string]interface{}  `json:"data"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
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
            AccountNo string  `json:"account_no"`
            BranchName string  `json:"branch_name"`
            AccountHolder string  `json:"account_holder"`
            IfscCode string  `json:"ifsc_code"`
         
    }
    
    // AddBeneficiaryDetailsOTPRequest used by Payment
    type AddBeneficiaryDetailsOTPRequest struct {

        
            OrderID string  `json:"order_id"`
            Details BankDetailsForOTP  `json:"details"`
         
    }
    
    // IfscCodeResponse used by Payment
    type IfscCodeResponse struct {

        
            BankName string  `json:"bank_name"`
            BranchName string  `json:"branch_name"`
            Success bool  `json:"success"`
         
    }
    
    // OrderBeneficiaryDetails used by Payment
    type OrderBeneficiaryDetails struct {

        
            Email string  `json:"email"`
            DelightsUserName string  `json:"delights_user_name"`
            ModifiedOn string  `json:"modified_on"`
            Address string  `json:"address"`
            BankName string  `json:"bank_name"`
            DisplayName string  `json:"display_name"`
            Comment string  `json:"comment"`
            BeneficiaryID string  `json:"beneficiary_id"`
            Mobile string  `json:"mobile"`
            BranchName string  `json:"branch_name"`
            AccountHolder string  `json:"account_holder"`
            IfscCode string  `json:"ifsc_code"`
            Title string  `json:"title"`
            AccountNo string  `json:"account_no"`
            TransferMode string  `json:"transfer_mode"`
            Subtitle string  `json:"subtitle"`
            IsActive bool  `json:"is_active"`
            ID float64  `json:"id"`
            CreatedOn string  `json:"created_on"`
         
    }
    
    // OrderBeneficiaryResponse used by Payment
    type OrderBeneficiaryResponse struct {

        
            ShowBeneficiaryDetails bool  `json:"show_beneficiary_details"`
            Beneficiaries []OrderBeneficiaryDetails  `json:"beneficiaries"`
         
    }
    
    // MultiTenderPaymentMeta used by Payment
    type MultiTenderPaymentMeta struct {

        
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            PaymentGateway string  `json:"payment_gateway"`
            OrderID string  `json:"order_id"`
            CurrentStatus string  `json:"current_status"`
            PaymentID string  `json:"payment_id"`
         
    }
    
    // MultiTenderPaymentMethod used by Payment
    type MultiTenderPaymentMethod struct {

        
            Mode string  `json:"mode"`
            Name string  `json:"name"`
            Meta MultiTenderPaymentMeta  `json:"meta"`
            Amount float64  `json:"amount"`
         
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

        
            RemainingLimit float64  `json:"remaining_limit"`
            IsActive bool  `json:"is_active"`
            Usages float64  `json:"usages"`
            Limit float64  `json:"limit"`
            UserID string  `json:"user_id"`
         
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

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    

    
    // PaymentModeInfo used by Order
    type PaymentModeInfo struct {

        
            Logo string  `json:"logo"`
            Type string  `json:"type"`
         
    }
    
    // ShipmentItemFulFillingStore used by Order
    type ShipmentItemFulFillingStore struct {

        
            ID string  `json:"id"`
            Code string  `json:"code"`
         
    }
    
    // Prices used by Order
    type Prices struct {

        
            PriceMarked float64  `json:"price_marked"`
            AmountPaidRoundoff float64  `json:"amount_paid_roundoff"`
            Discount float64  `json:"discount"`
            CashbackApplied float64  `json:"cashback_applied"`
            RefundAmount float64  `json:"refund_amount"`
            CodCharges float64  `json:"cod_charges"`
            FyndCredits float64  `json:"fynd_credits"`
            RefundCredit float64  `json:"refund_credit"`
            PromotionEffectiveDiscount float64  `json:"promotion_effective_discount"`
            CouponValue float64  `json:"coupon_value"`
            Cashback float64  `json:"cashback"`
            AmountPaid float64  `json:"amount_paid"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            DeliveryCharge float64  `json:"delivery_charge"`
            PriceEffective float64  `json:"price_effective"`
            ValueOfGood float64  `json:"value_of_good"`
         
    }
    
    // ShipmentStatus used by Order
    type ShipmentStatus struct {

        
            OpsStatus string  `json:"ops_status"`
            HexCode string  `json:"hex_code"`
            Title string  `json:"title"`
            ActualStatus string  `json:"actual_status"`
            Status string  `json:"status"`
         
    }
    
    // UserDataInfo used by Order
    type UserDataInfo struct {

        
            Gender string  `json:"gender"`
            AvisUserID string  `json:"avis_user_id"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            LastName string  `json:"last_name"`
            Email string  `json:"email"`
            Mobile string  `json:"mobile"`
            FirstName string  `json:"first_name"`
            IsAnonymousUser bool  `json:"is_anonymous_user"`
         
    }
    
    // GSTDetailsData used by Order
    type GSTDetailsData struct {

        
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            GstinCode string  `json:"gstin_code"`
            GstFee float64  `json:"gst_fee"`
            ValueOfGood float64  `json:"value_of_good"`
         
    }
    
    // PlatformItem used by Order
    type PlatformItem struct {

        
            CanCancel bool  `json:"can_cancel"`
            CanReturn bool  `json:"can_return"`
            Size string  `json:"size"`
            L1Category []string  `json:"l1_category"`
            Images []string  `json:"images"`
            Name string  `json:"name"`
            DepartmentID float64  `json:"department_id"`
            Code string  `json:"code"`
            Image []string  `json:"image"`
            Color string  `json:"color"`
            L3Category float64  `json:"l3_category"`
            ID float64  `json:"id"`
            L3CategoryName string  `json:"l3_category_name"`
         
    }
    
    // BagUnit used by Order
    type BagUnit struct {

        
            ShipmentID string  `json:"shipment_id"`
            CanReturn bool  `json:"can_return"`
            CanCancel bool  `json:"can_cancel"`
            TotalShipmentBags float64  `json:"total_shipment_bags"`
            OrderingChannel string  `json:"ordering_channel"`
            Gst GSTDetailsData  `json:"gst"`
            Item PlatformItem  `json:"item"`
            Prices Prices  `json:"prices"`
            BagID float64  `json:"bag_id"`
            Status map[string]interface{}  `json:"status"`
            ItemQuantity float64  `json:"item_quantity"`
         
    }
    
    // ShipmentItem used by Order
    type ShipmentItem struct {

        
            PaymentModeInfo PaymentModeInfo  `json:"payment_mode_info"`
            FulfillingStore ShipmentItemFulFillingStore  `json:"fulfilling_store"`
            Application map[string]interface{}  `json:"application"`
            FulfillingCentre string  `json:"fulfilling_centre"`
            TotalShipmentsInOrder float64  `json:"total_shipments_in_order"`
            PaymentMethods map[string]interface{}  `json:"payment_methods"`
            CreatedAt string  `json:"created_at"`
            Prices Prices  `json:"prices"`
            TotalBagsCount float64  `json:"total_bags_count"`
            ShipmentStatus ShipmentStatus  `json:"shipment_status"`
            User UserDataInfo  `json:"user"`
            ID string  `json:"id"`
            Channel map[string]interface{}  `json:"channel"`
            Sla map[string]interface{}  `json:"sla"`
            Bags []BagUnit  `json:"bags"`
            ShipmentCreatedAt float64  `json:"shipment_created_at"`
         
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
            Value string  `json:"value"`
            Options []FilterInfoOption  `json:"options"`
            Type string  `json:"type"`
         
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
    
    // FulfillingStore used by Order
    type FulfillingStore struct {

        
            ID string  `json:"id"`
            FulfillmentChannel string  `json:"fulfillment_channel"`
            State string  `json:"state"`
            Address string  `json:"address"`
            Phone string  `json:"phone"`
            Meta map[string]interface{}  `json:"meta"`
            Code string  `json:"code"`
            City string  `json:"city"`
            Pincode string  `json:"pincode"`
            StoreName string  `json:"store_name"`
            Country string  `json:"country"`
            ContactPerson string  `json:"contact_person"`
         
    }
    
    // ShipmentStatusData used by Order
    type ShipmentStatusData struct {

        
            ShipmentID string  `json:"shipment_id"`
            BagList []float64  `json:"bag_list"`
            CreatedAt string  `json:"created_at"`
            ID float64  `json:"id"`
            Status string  `json:"status"`
         
    }
    
    // DPDetailsData used by Order
    type DPDetailsData struct {

        
            AwbNo string  `json:"awb_no"`
            GstTag string  `json:"gst_tag"`
            Name string  `json:"name"`
            TrackURL string  `json:"track_url"`
            Pincode string  `json:"pincode"`
            ID string  `json:"id"`
            Country string  `json:"country"`
            EwayBillID string  `json:"eway_bill_id"`
         
    }
    
    // TrackingList used by Order
    type TrackingList struct {

        
            IsPassed bool  `json:"is_passed"`
            IsCurrent bool  `json:"is_current"`
            Text string  `json:"text"`
            Time string  `json:"time"`
            Status string  `json:"status"`
         
    }
    
    // OrderDetailsData used by Order
    type OrderDetailsData struct {

        
            FyndOrderID string  `json:"fynd_order_id"`
            OrderDate string  `json:"order_date"`
            OrderingChannel string  `json:"ordering_channel"`
            Source string  `json:"source"`
            CodCharges string  `json:"cod_charges"`
            AffiliateID string  `json:"affiliate_id"`
            OrderValue string  `json:"order_value"`
            OrderingChannelLogo map[string]interface{}  `json:"ordering_channel_logo"`
            TaxDetails map[string]interface{}  `json:"tax_details"`
         
    }
    
    // ShipmentPayments used by Order
    type ShipmentPayments struct {

        
            Source string  `json:"source"`
            Logo string  `json:"logo"`
            Mode string  `json:"mode"`
         
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

        
            Value float64  `json:"value"`
            Type string  `json:"type"`
         
    }
    
    // AppliedPromos used by Order
    type AppliedPromos struct {

        
            PromotionType string  `json:"promotion_type"`
            PromoID string  `json:"promo_id"`
            ArticleQuantity float64  `json:"article_quantity"`
            BuyRules []BuyRules  `json:"buy_rules"`
            PromotionName string  `json:"promotion_name"`
            DiscountRules []DiscountRules  `json:"discount_rules"`
            Amount float64  `json:"amount"`
            MrpPromotion bool  `json:"mrp_promotion"`
         
    }
    
    // BagGST used by Order
    type BagGST struct {

        
            GstTag string  `json:"gst_tag"`
            IsDefaultHsnCode bool  `json:"is_default_hsn_code"`
            GstinCode string  `json:"gstin_code"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            HsnCode string  `json:"hsn_code"`
            GstFee float64  `json:"gst_fee"`
            ValueOfGood float64  `json:"value_of_good"`
         
    }
    
    // PlatformDeliveryAddress used by Order
    type PlatformDeliveryAddress struct {

        
            Address1 string  `json:"address1"`
            Landmark string  `json:"landmark"`
            State string  `json:"state"`
            Longitude float64  `json:"longitude"`
            Pincode string  `json:"pincode"`
            Phone string  `json:"phone"`
            UpdatedAt string  `json:"updated_at"`
            AddressCategory string  `json:"address_category"`
            Area string  `json:"area"`
            City string  `json:"city"`
            Latitude float64  `json:"latitude"`
            CreatedAt string  `json:"created_at"`
            Email string  `json:"email"`
            AddressType string  `json:"address_type"`
            Version string  `json:"version"`
            Address2 string  `json:"address2"`
            Country string  `json:"country"`
            ContactPerson string  `json:"contact_person"`
         
    }
    
    // OrderBrandName used by Order
    type OrderBrandName struct {

        
            Logo string  `json:"logo"`
            BrandName string  `json:"brand_name"`
            CreatedOn float64  `json:"created_on"`
            Company string  `json:"company"`
            ID float64  `json:"id"`
            ModifiedOn float64  `json:"modified_on"`
         
    }
    
    // OrderBagArticle used by Order
    type OrderBagArticle struct {

        
            UID string  `json:"uid"`
            Identifiers map[string]interface{}  `json:"identifiers"`
            ReturnConfig map[string]interface{}  `json:"return_config"`
         
    }
    
    // BagConfigs used by Order
    type BagConfigs struct {

        
            EnableTracking bool  `json:"enable_tracking"`
            CanBeCancelled bool  `json:"can_be_cancelled"`
            IsReturnable bool  `json:"is_returnable"`
            IsActive bool  `json:"is_active"`
            IsCustomerReturnAllowed bool  `json:"is_customer_return_allowed"`
            AllowForceReturn bool  `json:"allow_force_return"`
         
    }
    
    // Identifier used by Order
    type Identifier struct {

        
            Ean string  `json:"ean"`
         
    }
    
    // FinancialBreakup used by Order
    type FinancialBreakup struct {

        
            AmountPaidRoundoff float64  `json:"amount_paid_roundoff"`
            Discount float64  `json:"discount"`
            CashbackApplied float64  `json:"cashback_applied"`
            GstTag string  `json:"gst_tag"`
            CouponValue float64  `json:"coupon_value"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            PriceEffective float64  `json:"price_effective"`
            PriceMarked float64  `json:"price_marked"`
            DeliveryCharge float64  `json:"delivery_charge"`
            RefundCredit float64  `json:"refund_credit"`
            PromotionEffectiveDiscount float64  `json:"promotion_effective_discount"`
            ItemName string  `json:"item_name"`
            Cashback float64  `json:"cashback"`
            AmountPaid float64  `json:"amount_paid"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            PmPriceSplit map[string]interface{}  `json:"pm_price_split"`
            TransferPrice float64  `json:"transfer_price"`
            TotalUnits float64  `json:"total_units"`
            ValueOfGood float64  `json:"value_of_good"`
            Identifiers Identifier  `json:"identifiers"`
            Size string  `json:"size"`
            FyndCredits float64  `json:"fynd_credits"`
            CodCharges float64  `json:"cod_charges"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            AddedToFyndCash bool  `json:"added_to_fynd_cash"`
            HsnCode string  `json:"hsn_code"`
            GstFee string  `json:"gst_fee"`
         
    }
    
    // OrderBags used by Order
    type OrderBags struct {

        
            AppliedPromos []AppliedPromos  `json:"applied_promos"`
            DisplayName string  `json:"display_name"`
            GstDetails BagGST  `json:"gst_details"`
            DeliveryAddress PlatformDeliveryAddress  `json:"delivery_address"`
            Prices Prices  `json:"prices"`
            BagID float64  `json:"bag_id"`
            Brand OrderBrandName  `json:"brand"`
            SellerIdentifier string  `json:"seller_identifier"`
            EntityType string  `json:"entity_type"`
            CanReturn bool  `json:"can_return"`
            Article OrderBagArticle  `json:"article"`
            Identifier string  `json:"identifier"`
            ParentPromoBags map[string]interface{}  `json:"parent_promo_bags"`
            LineNumber float64  `json:"line_number"`
            BagConfigs BagConfigs  `json:"bag_configs"`
            Quantity float64  `json:"quantity"`
            CanCancel bool  `json:"can_cancel"`
            FinancialBreakup []FinancialBreakup  `json:"financial_breakup"`
            Item PlatformItem  `json:"item"`
            CurrentStatus string  `json:"current_status"`
         
    }
    
    // UserDetailsData used by Order
    type UserDetailsData struct {

        
            State string  `json:"state"`
            Pincode string  `json:"pincode"`
            Address string  `json:"address"`
            Phone string  `json:"phone"`
            Name string  `json:"name"`
            City string  `json:"city"`
            Email string  `json:"email"`
            Country string  `json:"country"`
         
    }
    
    // BagStateMapper used by Order
    type BagStateMapper struct {

        
            StateType string  `json:"state_type"`
            IsActive bool  `json:"is_active"`
            NotifyCustomer bool  `json:"notify_customer"`
            AppDisplayName string  `json:"app_display_name"`
            AppStateName string  `json:"app_state_name"`
            DisplayName string  `json:"display_name"`
            BsID float64  `json:"bs_id"`
            Name string  `json:"name"`
            AppFacing bool  `json:"app_facing"`
            JourneyType string  `json:"journey_type"`
         
    }
    
    // BagStatusHistory used by Order
    type BagStatusHistory struct {

        
            ShipmentID string  `json:"shipment_id"`
            BagStateMapper BagStateMapper  `json:"bag_state_mapper"`
            StateType string  `json:"state_type"`
            Reasons []map[string]interface{}  `json:"reasons"`
            DeliveryPartnerID float64  `json:"delivery_partner_id"`
            AppDisplayName bool  `json:"app_display_name"`
            UpdatedAt string  `json:"updated_at"`
            CreatedAt string  `json:"created_at"`
            DisplayName bool  `json:"display_name"`
            StoreID float64  `json:"store_id"`
            BagID float64  `json:"bag_id"`
            BshID float64  `json:"bsh_id"`
            KafkaSync bool  `json:"kafka_sync"`
            Status string  `json:"status"`
            StateID float64  `json:"state_id"`
            DeliveryAwbNumber string  `json:"delivery_awb_number"`
            Forward bool  `json:"forward"`
         
    }
    
    // ShipmentInfoResponse used by Order
    type ShipmentInfoResponse struct {

        
            ShipmentID string  `json:"shipment_id"`
            EnableTracking bool  `json:"enable_tracking"`
            SecuredDeliveryFlag string  `json:"secured_delivery_flag"`
            ShipmentQuantity float64  `json:"shipment_quantity"`
            GoGreen bool  `json:"go_green"`
            TotalItems float64  `json:"total_items"`
            FulfillingStore FulfillingStore  `json:"fulfilling_store"`
            BeneficiaryDetails bool  `json:"beneficiary_details"`
            ChildNodes []string  `json:"child_nodes"`
            GstDetails GSTDetailsData  `json:"gst_details"`
            Prices Prices  `json:"prices"`
            OrderStatus map[string]interface{}  `json:"order_status"`
            UserID string  `json:"user_id"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            Items []map[string]interface{}  `json:"items"`
            Vertical string  `json:"vertical"`
            PriorityText string  `json:"priority_text"`
            OperationalStatus string  `json:"operational_status"`
            Status ShipmentStatusData  `json:"status"`
            ForwardTrackingList []map[string]interface{}  `json:"forward_tracking_list"`
            Invoice map[string]interface{}  `json:"invoice"`
            DueDate string  `json:"due_date"`
            EnableDpTracking string  `json:"enable_dp_tracking"`
            OrderingStore map[string]interface{}  `json:"ordering_store"`
            PlatformLogo bool  `json:"platform_logo"`
            BankData map[string]interface{}  `json:"bank_data"`
            FyndstoreEmp map[string]interface{}  `json:"fyndstore_emp"`
            CurrentShipmentStatus map[string]interface{}  `json:"current_shipment_status"`
            Company map[string]interface{}  `json:"company"`
            DpDetails DPDetailsData  `json:"dp_details"`
            Escalation map[string]interface{}  `json:"escalation"`
            IsFyndCoupon bool  `json:"is_fynd_coupon"`
            CustomMeta []map[string]interface{}  `json:"custom_meta"`
            TrackingList []TrackingList  `json:"tracking_list"`
            LockStatus string  `json:"lock_status"`
            UserInfo map[string]interface{}  `json:"user_info"`
            Order OrderDetailsData  `json:"order"`
            UserAgent string  `json:"user_agent"`
            CanReturn bool  `json:"can_return"`
            StatusProgress float64  `json:"status_progress"`
            Coupon map[string]interface{}  `json:"coupon"`
            Payments ShipmentPayments  `json:"payments"`
            PackagingType string  `json:"packaging_type"`
            PayButton string  `json:"pay_button"`
            PickedDate string  `json:"picked_date"`
            IsNotFyndSource bool  `json:"is_not_fynd_source"`
            DeliveryStatus []map[string]interface{}  `json:"delivery_status"`
            Mid string  `json:"mid"`
            OrderCreatedTime string  `json:"order_created_time"`
            CanBreak string  `json:"can_break"`
            RefundDetails map[string]interface{}  `json:"refund_details"`
            JourneyType string  `json:"journey_type"`
            Bags []OrderBags  `json:"bags"`
            IsInvoiced bool  `json:"is_invoiced"`
            ShipmentImages []string  `json:"shipment_images"`
            TrackingURL string  `json:"tracking_url"`
            ReplacementDetails string  `json:"replacement_details"`
            OrderType string  `json:"order_type"`
            BillingDetails UserDetailsData  `json:"billing_details"`
            EmailID string  `json:"email_id"`
            CanCancel bool  `json:"can_cancel"`
            DeliverySlot map[string]interface{}  `json:"delivery_slot"`
            IsPackagingOrder bool  `json:"is_packaging_order"`
            ForwardShipmentStatus []map[string]interface{}  `json:"forward_shipment_status"`
            PaymentMode string  `json:"payment_mode"`
            ForwardOrderStatus []map[string]interface{}  `json:"forward_order_status"`
            IsFyndStore string  `json:"is_fynd_store"`
            BagStatusHistory []BagStatusHistory  `json:"bag_status_history"`
            IsPdsr string  `json:"is_pdsr"`
            DeliveryDetails UserDetailsData  `json:"delivery_details"`
            CreditNoteID string  `json:"credit_note_id"`
            RefundText string  `json:"refund_text"`
            KiranaStoreID string  `json:"kirana_store_id"`
            ShipmentStatus string  `json:"shipment_status"`
            TotalBags float64  `json:"total_bags"`
         
    }
    
    // OrderDict used by Order
    type OrderDict struct {

        
            OrderDate string  `json:"order_date"`
            FyndOrderID string  `json:"fynd_order_id"`
            ShipmentCount float64  `json:"shipment_count"`
         
    }
    
    // PlatformShipment used by Order
    type PlatformShipment struct {

        
            ShipmentID string  `json:"shipment_id"`
            ShipmentQuantity float64  `json:"shipment_quantity"`
            TotalItems float64  `json:"total_items"`
            FulfillingStore FulfillingStore  `json:"fulfilling_store"`
            GstDetails GSTDetailsData  `json:"gst_details"`
            Prices Prices  `json:"prices"`
            Vertical string  `json:"vertical"`
            PriorityText string  `json:"priority_text"`
            OperationalStatus string  `json:"operational_status"`
            Status ShipmentStatusData  `json:"status"`
            EnableDpTracking string  `json:"enable_dp_tracking"`
            PlatformLogo string  `json:"platform_logo"`
            DpDetails DPDetailsData  `json:"dp_details"`
            CustomMeta []map[string]interface{}  `json:"custom_meta"`
            TrackingList []TrackingList  `json:"tracking_list"`
            Order OrderDetailsData  `json:"order"`
            UserAgent string  `json:"user_agent"`
            Coupon map[string]interface{}  `json:"coupon"`
            Payments ShipmentPayments  `json:"payments"`
            PackagingType string  `json:"packaging_type"`
            PickedDate string  `json:"picked_date"`
            JourneyType string  `json:"journey_type"`
            Bags []OrderBags  `json:"bags"`
            ShipmentImages []string  `json:"shipment_images"`
            BillingDetails UserDetailsData  `json:"billing_details"`
            DeliverySlot map[string]interface{}  `json:"delivery_slot"`
            PaymentMode string  `json:"payment_mode"`
            BagStatusHistory []BagStatusHistory  `json:"bag_status_history"`
            DeliveryDetails UserDetailsData  `json:"delivery_details"`
            ShipmentStatus string  `json:"shipment_status"`
            TotalBags float64  `json:"total_bags"`
         
    }
    
    // ShipmentDetailsResponse used by Order
    type ShipmentDetailsResponse struct {

        
            Order OrderDict  `json:"order"`
            CustomMeta []map[string]interface{}  `json:"custom_meta"`
            Shipments []PlatformShipment  `json:"shipments"`
            Success bool  `json:"success"`
         
    }
    
    // SubLane used by Order
    type SubLane struct {

        
            TotalItems float64  `json:"total_items"`
            Value string  `json:"value"`
            Index float64  `json:"index"`
            Actions []map[string]interface{}  `json:"actions"`
            Text string  `json:"text"`
         
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

        
            Name string  `json:"name"`
            Value string  `json:"value"`
            Display string  `json:"display"`
         
    }
    
    // PlatformChannel used by Order
    type PlatformChannel struct {

        
            Name string  `json:"name"`
            Logo string  `json:"logo"`
         
    }
    
    // PlatformOrderItems used by Order
    type PlatformOrderItems struct {

        
            PaymentMode string  `json:"payment_mode"`
            Shipments []PlatformShipment  `json:"shipments"`
            BreakupValues []PlatformBreakupValues  `json:"breakup_values"`
            OrderCreatedTime string  `json:"order_created_time"`
            OrderValue float64  `json:"order_value"`
            UserInfo UserDataInfo  `json:"user_info"`
            TotalOrderValue float64  `json:"total_order_value"`
            OrderID string  `json:"order_id"`
            Channel PlatformChannel  `json:"channel"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // OrderListingResponse used by Order
    type OrderListingResponse struct {

        
            Lane string  `json:"lane"`
            Message string  `json:"message"`
            TotalCount float64  `json:"total_count"`
            Items []PlatformOrderItems  `json:"items"`
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
            Value float64  `json:"value"`
            Options []Options  `json:"options"`
            Key string  `json:"key"`
         
    }
    
    // MetricCountResponse used by Order
    type MetricCountResponse struct {

        
            Items []MetricsCount  `json:"items"`
         
    }
    
    // PlatformTrack used by Order
    type PlatformTrack struct {

        
            LastLocationRecievedAt string  `json:"last_location_recieved_at"`
            ShipmentType string  `json:"shipment_type"`
            UpdatedAt string  `json:"updated_at"`
            AccountName string  `json:"account_name"`
            UpdatedTime string  `json:"updated_time"`
            RawStatus string  `json:"raw_status"`
            Awb string  `json:"awb"`
            Reason string  `json:"reason"`
            Status string  `json:"status"`
            Meta map[string]interface{}  `json:"meta"`
         
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

        
            RequestDetails map[string]interface{}  `json:"request_details"`
            ReportID string  `json:"report_id"`
            ReportName string  `json:"report_name"`
            ReportCreatedAt string  `json:"report_created_at"`
            DisplayName string  `json:"display_name"`
            S3Key string  `json:"s3_key"`
            ReportRequestedAt string  `json:"report_requested_at"`
            ReportType string  `json:"report_type"`
            Status string  `json:"status"`
         
    }
    
    // JioCodeUpsertDataSet used by Order
    type JioCodeUpsertDataSet struct {

        
            ItemID string  `json:"item_id"`
            JioCode string  `json:"jio_code"`
            CompanyID string  `json:"company_id"`
            ArticleID string  `json:"article_id"`
         
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
            Success bool  `json:"success"`
            Identifier string  `json:"identifier"`
            Error []NestedErrorSchemaDataSet  `json:"error"`
            TraceID string  `json:"trace_id"`
         
    }
    
    // BulkInvoicingResponse used by Order
    type BulkInvoicingResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // BulkInvoiceLabelResponse used by Order
    type BulkInvoiceLabelResponse struct {

        
            Invoice map[string]interface{}  `json:"invoice"`
            Data map[string]interface{}  `json:"data"`
            DoInvoiceLabelGenerated bool  `json:"do_invoice_label_generated"`
            InvoiceStatus string  `json:"invoice_status"`
            CompanyID string  `json:"company_id"`
            BatchID string  `json:"batch_id"`
            Label map[string]interface{}  `json:"label"`
            StoreID string  `json:"store_id"`
            StoreCode string  `json:"store_code"`
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

        
            Upload FileUploadResponse  `json:"upload"`
            ContentType string  `json:"content_type"`
            FileName string  `json:"file_name"`
            Operation string  `json:"operation"`
            Size float64  `json:"size"`
            Cdn URL  `json:"cdn"`
            Namespace string  `json:"namespace"`
            Method string  `json:"method"`
            FilePath string  `json:"file_path"`
            Tags []string  `json:"tags"`
         
    }
    
    // BulkListingPage used by Order
    type BulkListingPage struct {

        
            HasPrevious bool  `json:"has_previous"`
            Type string  `json:"type"`
            Size float64  `json:"size"`
            Total float64  `json:"total"`
            Current float64  `json:"current"`
            HasNext bool  `json:"has_next"`
         
    }
    
    // bulkListingData used by Order
    type bulkListingData struct {

        
            BatchID string  `json:"batch_id"`
            UserID string  `json:"user_id"`
            StoreCode string  `json:"store_code"`
            ExcelURL string  `json:"excel_url"`
            UserName string  `json:"user_name"`
            Status string  `json:"status"`
            StoreName string  `json:"store_name"`
            SuccessfulShipments []map[string]interface{}  `json:"successful_shipments"`
            Total float64  `json:"total"`
            Successful float64  `json:"successful"`
            StoreID float64  `json:"store_id"`
            Failed float64  `json:"failed"`
            Processing float64  `json:"processing"`
            UploadedOn string  `json:"uploaded_on"`
            ProcessingShipments []string  `json:"processing_shipments"`
            CompanyID float64  `json:"company_id"`
            ID string  `json:"id"`
            FailedShipments []map[string]interface{}  `json:"failed_shipments"`
            FileName string  `json:"file_name"`
         
    }
    
    // BulkListingResponse used by Order
    type BulkListingResponse struct {

        
            Page BulkListingPage  `json:"page"`
            Error string  `json:"error"`
            Data []bulkListingData  `json:"data"`
            Success bool  `json:"success"`
         
    }
    
    // DateRange used by Order
    type DateRange struct {

        
            FromDate string  `json:"from_date"`
            ToDate string  `json:"to_date"`
         
    }
    
    // ManifestFilter used by Order
    type ManifestFilter struct {

        
            Stores string  `json:"stores"`
            Lane string  `json:"lane"`
            SalesChannels string  `json:"sales_channels"`
            SalesChannelName string  `json:"sales_channel_name"`
            DateRange DateRange  `json:"date_range"`
            StoreName string  `json:"store_name"`
            DpIds string  `json:"dp_ids"`
            DpName string  `json:"dp_name"`
         
    }
    
    // GeneratedManifestItem used by Order
    type GeneratedManifestItem struct {

        
            ManifestID string  `json:"manifest_id"`
            Filters ManifestFilter  `json:"filters"`
            IsActive bool  `json:"is_active"`
            CompanyID float64  `json:"company_id"`
            CreatedBy string  `json:"created_by"`
            CreatedAt string  `json:"created_at"`
            Status string  `json:"status"`
         
    }
    
    // ManifestPage used by Order
    type ManifestPage struct {

        
            HasPrevious bool  `json:"has_previous"`
            Type string  `json:"type"`
            Size string  `json:"size"`
            Total float64  `json:"total"`
            Current string  `json:"current"`
            HasNext bool  `json:"has_next"`
         
    }
    
    // GeneratedManifestResponse used by Order
    type GeneratedManifestResponse struct {

        
            Items []GeneratedManifestItem  `json:"items"`
            Page ManifestPage  `json:"page"`
         
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
            ProcessingShipmentsCount float64  `json:"processing_shipments_count"`
            CompanyID string  `json:"company_id"`
            TotalShipmentsCount float64  `json:"total_shipments_count"`
            BatchID string  `json:"batch_id"`
            FailedShipmentsCount float64  `json:"failed_shipments_count"`
         
    }
    
    // BulkActionDetailsResponse used by Order
    type BulkActionDetailsResponse struct {

        
            Data []BulkActionDetailsDataField  `json:"data"`
            UploadedOn string  `json:"uploaded_on"`
            FailedRecords []string  `json:"failed_records"`
            Error []string  `json:"error"`
            UserID string  `json:"user_id"`
            Message string  `json:"message"`
            Status bool  `json:"status"`
            UploadedBy string  `json:"uploaded_by"`
            Success string  `json:"success"`
         
    }
    
    // BagGSTDetails used by Order
    type BagGSTDetails struct {

        
            CgstTaxPercentage float64  `json:"cgst_tax_percentage"`
            SgstGstFee string  `json:"sgst_gst_fee"`
            GstTag string  `json:"gst_tag"`
            IgstGstFee string  `json:"igst_gst_fee"`
            HsnCodeID string  `json:"hsn_code_id"`
            IsDefaultHsnCode bool  `json:"is_default_hsn_code"`
            CgstGstFee string  `json:"cgst_gst_fee"`
            GstinCode string  `json:"gstin_code"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            IgstTaxPercentage float64  `json:"igst_tax_percentage"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            SgstTaxPercentage float64  `json:"sgst_tax_percentage"`
            HsnCode string  `json:"hsn_code"`
            GstFee float64  `json:"gst_fee"`
            ValueOfGood float64  `json:"value_of_good"`
         
    }
    
    // BagReturnableCancelableStatus used by Order
    type BagReturnableCancelableStatus struct {

        
            EnableTracking bool  `json:"enable_tracking"`
            CanBeCancelled bool  `json:"can_be_cancelled"`
            IsReturnable bool  `json:"is_returnable"`
            IsActive bool  `json:"is_active"`
            IsCustomerReturnAllowed bool  `json:"is_customer_return_allowed"`
         
    }
    
    // Formatted used by Order
    type Formatted struct {

        
            FMax string  `json:"f_max"`
            FMin string  `json:"f_min"`
         
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
    
    // EInvoice used by Order
    type EInvoice struct {

        
            ErrorCode string  `json:"error_code"`
            AcknowledgeDate string  `json:"acknowledge_date"`
            SignedQrCode string  `json:"signed_qr_code"`
            SignedInvoice string  `json:"signed_invoice"`
            ErrorMessage string  `json:"error_message"`
            Irn string  `json:"irn"`
            AcknowledgeNo float64  `json:"acknowledge_no"`
         
    }
    
    // EinvoiceInfo used by Order
    type EinvoiceInfo struct {

        
            CreditNote EInvoice  `json:"credit_note"`
            Invoice EInvoice  `json:"invoice"`
         
    }
    
    // BuyerDetails used by Order
    type BuyerDetails struct {

        
            Gstin string  `json:"gstin"`
            State string  `json:"state"`
            Address string  `json:"address"`
            Name string  `json:"name"`
            City string  `json:"city"`
            AjioSiteID string  `json:"ajio_site_id"`
            Pincode float64  `json:"pincode"`
         
    }
    
    // DebugInfo used by Order
    type DebugInfo struct {

        
            StormbreakerUUID string  `json:"stormbreaker_uuid"`
         
    }
    
    // ShipmentMeta used by Order
    type ShipmentMeta struct {

        
            AssignDpFromSb bool  `json:"assign_dp_from_sb"`
            External map[string]interface{}  `json:"external"`
            PackagingName string  `json:"packaging_name"`
            ForwardAffiliateOrderID string  `json:"forward_affiliate_order_id"`
            Formatted Formatted  `json:"formatted"`
            EwaybillInfo map[string]interface{}  `json:"ewaybill_info"`
            LockData LockData  `json:"lock_data"`
            DpName string  `json:"dp_name"`
            Timestamp ShipmentTimeStamp  `json:"timestamp"`
            DueDate string  `json:"due_date"`
            Weight float64  `json:"weight"`
            ReturnStoreNode float64  `json:"return_store_node"`
            AutoTriggerDpAssignmentAcf bool  `json:"auto_trigger_dp_assignment_acf"`
            EinvoiceInfo EinvoiceInfo  `json:"einvoice_info"`
            DpOptions map[string]interface{}  `json:"dp_options"`
            ReturnAffiliateShipmentID string  `json:"return_affiliate_shipment_id"`
            PoNumber string  `json:"po_number"`
            B2bBuyerDetails BuyerDetails  `json:"b2b_buyer_details"`
            DpSortKey string  `json:"dp_sort_key"`
            B2cBuyerDetails map[string]interface{}  `json:"b2c_buyer_details"`
            FulfilmentPriorityText string  `json:"fulfilment_priority_text"`
            DebugInfo DebugInfo  `json:"debug_info"`
            DpID string  `json:"dp_id"`
            ReturnAwbNumber string  `json:"return_awb_number"`
            OrderType string  `json:"order_type"`
            MarketplaceStoreID string  `json:"marketplace_store_id"`
            SameStoreAvailable bool  `json:"same_store_available"`
            AwbNumber string  `json:"awb_number"`
            BoxType string  `json:"box_type"`
            ShipmentWeight float64  `json:"shipment_weight"`
            ForwardAffiliateShipmentID string  `json:"forward_affiliate_shipment_id"`
            ReturnDetails map[string]interface{}  `json:"return_details"`
            ShipmentVolumetricWeight float64  `json:"shipment_volumetric_weight"`
            BagWeight map[string]interface{}  `json:"bag_weight"`
            ReturnAffiliateOrderID string  `json:"return_affiliate_order_id"`
            StoreInvoiceUpdatedDate string  `json:"store_invoice_updated_date"`
         
    }
    
    // PDFLinks used by Order
    type PDFLinks struct {

        
            PoInvoice string  `json:"po_invoice"`
            Invoice string  `json:"invoice"`
            LabelA6 string  `json:"label_a6"`
            InvoiceType string  `json:"invoice_type"`
            InvoiceA4 string  `json:"invoice_a4"`
            LabelType string  `json:"label_type"`
            Label string  `json:"label"`
            B2b string  `json:"b2b"`
            InvoiceA6 string  `json:"invoice_a6"`
            LabelA4 string  `json:"label_a4"`
            InvoicePos string  `json:"invoice_pos"`
            CreditNoteURL string  `json:"credit_note_url"`
            LabelPos string  `json:"label_pos"`
         
    }
    
    // AffiliateMeta used by Order
    type AffiliateMeta struct {

        
            DueDate string  `json:"due_date"`
            OrderItemID string  `json:"order_item_id"`
            ChannelShipmentID string  `json:"channel_shipment_id"`
            IsPriority bool  `json:"is_priority"`
            BoxType string  `json:"box_type"`
            Quantity float64  `json:"quantity"`
            CouponCode string  `json:"coupon_code"`
            SizeLevelTotalQty float64  `json:"size_level_total_qty"`
            ChannelOrderID string  `json:"channel_order_id"`
            EmployeeDiscount float64  `json:"employee_discount"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
         
    }
    
    // AffiliateDetails used by Order
    type AffiliateDetails struct {

        
            AffiliateBagID string  `json:"affiliate_bag_id"`
            AffiliateStoreID string  `json:"affiliate_store_id"`
            CompanyAffiliateTag string  `json:"company_affiliate_tag"`
            AdID string  `json:"ad_id"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            AffiliateID string  `json:"affiliate_id"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            ShipmentMeta ShipmentMeta  `json:"shipment_meta"`
            PdfLinks PDFLinks  `json:"pdf_links"`
            AffiliateMeta AffiliateMeta  `json:"affiliate_meta"`
         
    }
    
    // B2BPODetails used by Order
    type B2BPODetails struct {

        
            PoTaxAmount float64  `json:"po_tax_amount"`
            PoLineAmount float64  `json:"po_line_amount"`
            PartialCanRet bool  `json:"partial_can_ret"`
            ItemBasePrice float64  `json:"item_base_price"`
            TotalGstPercentage float64  `json:"total_gst_percentage"`
            DockerNumber string  `json:"docker_number"`
         
    }
    
    // BagMeta used by Order
    type BagMeta struct {

        
            B2bPoDetails B2BPODetails  `json:"b2b_po_details"`
         
    }
    
    // StoreAddress used by Order
    type StoreAddress struct {

        
            Address1 string  `json:"address1"`
            State string  `json:"state"`
            City string  `json:"city"`
            Latitude float64  `json:"latitude"`
            Email string  `json:"email"`
            Address2 string  `json:"address2"`
            Country string  `json:"country"`
            ContactPerson string  `json:"contact_person"`
            Landmark string  `json:"landmark"`
            CreatedAt string  `json:"created_at"`
            AddressType string  `json:"address_type"`
            Longitude float64  `json:"longitude"`
            CountryCode string  `json:"country_code"`
            Phone string  `json:"phone"`
            Area string  `json:"area"`
            Pincode float64  `json:"pincode"`
            UpdatedAt string  `json:"updated_at"`
            AddressCategory string  `json:"address_category"`
            Version string  `json:"version"`
         
    }
    
    // StoreEwaybill used by Order
    type StoreEwaybill struct {

        
            Enabled bool  `json:"enabled"`
         
    }
    
    // StoreEinvoice used by Order
    type StoreEinvoice struct {

        
            User string  `json:"user"`
            Enabled bool  `json:"enabled"`
            Password string  `json:"password"`
            Username string  `json:"username"`
         
    }
    
    // StoreGstCredentials used by Order
    type StoreGstCredentials struct {

        
            EWaybill StoreEwaybill  `json:"e_waybill"`
            EInvoice StoreEinvoice  `json:"e_invoice"`
         
    }
    
    // Document used by Order
    type Document struct {

        
            Value string  `json:"value"`
            LegalName string  `json:"legal_name"`
            URL string  `json:"url"`
            Verified bool  `json:"verified"`
            DsType string  `json:"ds_type"`
         
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

        
            NotificationEmails []string  `json:"notification_emails"`
            Stage string  `json:"stage"`
            GstNumber string  `json:"gst_number"`
            ProductReturnConfig map[string]interface{}  `json:"product_return_config"`
            AdditionalContactDetails map[string]interface{}  `json:"additional_contact_details"`
            DisplayName string  `json:"display_name"`
            GstCredentials StoreGstCredentials  `json:"gst_credentials"`
            Documents StoreDocuments  `json:"documents"`
            Timing []map[string]interface{}  `json:"timing"`
            EwaybillPortalDetails map[string]interface{}  `json:"ewaybill_portal_details"`
            EinvoicePortalDetails EInvoicePortalDetails  `json:"einvoice_portal_details"`
         
    }
    
    // Store used by Order
    type Store struct {

        
            LoginUsername string  `json:"login_username"`
            VatNo string  `json:"vat_no"`
            Address1 string  `json:"address1"`
            FulfillmentChannel string  `json:"fulfillment_channel"`
            State string  `json:"state"`
            IsActive bool  `json:"is_active"`
            StoreAddressJson StoreAddress  `json:"store_address_json"`
            Name string  `json:"name"`
            Code string  `json:"code"`
            City string  `json:"city"`
            Latitude float64  `json:"latitude"`
            Address2 string  `json:"address2"`
            Country string  `json:"country"`
            ContactPerson string  `json:"contact_person"`
            Meta StoreMeta  `json:"meta"`
            LocationType string  `json:"location_type"`
            ParentStoreID float64  `json:"parent_store_id"`
            BrandStoreTags []string  `json:"brand_store_tags"`
            IsEnabledForRecon bool  `json:"is_enabled_for_recon"`
            CreatedAt string  `json:"created_at"`
            SID string  `json:"s_id"`
            Longitude float64  `json:"longitude"`
            BrandID interface{}  `json:"brand_id"`
            OrderIntegrationID string  `json:"order_integration_id"`
            MallArea string  `json:"mall_area"`
            Phone float64  `json:"phone"`
            CompanyID float64  `json:"company_id"`
            StoreActiveFrom string  `json:"store_active_from"`
            AlohomoraUserID float64  `json:"alohomora_user_id"`
            Pincode string  `json:"pincode"`
            StoreEmail string  `json:"store_email"`
            IsArchived bool  `json:"is_archived"`
            PackagingMaterialCount float64  `json:"packaging_material_count"`
            UpdatedAt string  `json:"updated_at"`
            MallName string  `json:"mall_name"`
         
    }
    
    // Brand used by Order
    type Brand struct {

        
            BrandID float64  `json:"brand_id"`
            ScriptLastRan string  `json:"script_last_ran"`
            Logo string  `json:"logo"`
            CreditNoteExpiryDays float64  `json:"credit_note_expiry_days"`
            BrandName string  `json:"brand_name"`
            PickupLocation string  `json:"pickup_location"`
            CreatedOn float64  `json:"created_on"`
            StartDate string  `json:"start_date"`
            InvoicePrefix string  `json:"invoice_prefix"`
            Company string  `json:"company"`
            IsVirtualInvoice bool  `json:"is_virtual_invoice"`
            CreditNoteAllowed bool  `json:"credit_note_allowed"`
            ModifiedOn float64  `json:"modified_on"`
         
    }
    
    // Weight used by Order
    type Weight struct {

        
            Shipping float64  `json:"shipping"`
            IsDefault bool  `json:"is_default"`
            Unit string  `json:"unit"`
         
    }
    
    // Dimensions used by Order
    type Dimensions struct {

        
            Unit string  `json:"unit"`
            Height float64  `json:"height"`
            Width float64  `json:"width"`
            IsDefault bool  `json:"is_default"`
            Length float64  `json:"length"`
         
    }
    
    // ReturnConfig used by Order
    type ReturnConfig struct {

        
            Time float64  `json:"time"`
            Unit string  `json:"unit"`
            Returnable bool  `json:"returnable"`
         
    }
    
    // Article used by Order
    type Article struct {

        
            Identifiers Identifier  `json:"identifiers"`
            Size string  `json:"size"`
            ASet map[string]interface{}  `json:"a_set"`
            Weight Weight  `json:"weight"`
            RawMeta interface{}  `json:"raw_meta"`
            SellerIdentifier string  `json:"seller_identifier"`
            UID string  `json:"uid"`
            ChildDetails map[string]interface{}  `json:"child_details"`
            Code string  `json:"code"`
            Dimensions Dimensions  `json:"dimensions"`
            EspModified interface{}  `json:"esp_modified"`
            ID string  `json:"_id"`
            IsSet bool  `json:"is_set"`
            ReturnConfig ReturnConfig  `json:"return_config"`
         
    }
    
    // Dates used by Order
    type Dates struct {

        
            OrderCreated string  `json:"order_created"`
            DeliveryDate interface{}  `json:"delivery_date"`
         
    }
    
    // AffiliateBagDetails used by Order
    type AffiliateBagDetails struct {

        
            AffiliateBagID string  `json:"affiliate_bag_id"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            AffiliateMeta AffiliateMeta  `json:"affiliate_meta"`
            EmployeeDiscount float64  `json:"employee_discount"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
         
    }
    
    // ArticleDetails used by Order
    type ArticleDetails struct {

        
            Status map[string]interface{}  `json:"status"`
         
    }
    
    // Attributes used by Order
    type Attributes struct {

        
            PrimaryColor string  `json:"primary_color"`
            MarketerAddress string  `json:"marketer_address"`
            Essential string  `json:"essential"`
            BrandName string  `json:"brand_name"`
            MarketerName string  `json:"marketer_name"`
            Name string  `json:"name"`
            PrimaryColorHex string  `json:"primary_color_hex"`
            Gender []string  `json:"gender"`
            PrimaryMaterial string  `json:"primary_material"`
         
    }
    
    // Item used by Order
    type Item struct {

        
            L2Category []string  `json:"l2_category"`
            LastUpdatedAt string  `json:"last_updated_at"`
            Name string  `json:"name"`
            Code string  `json:"code"`
            Gender string  `json:"gender"`
            L1CategoryID float64  `json:"l1_category_id"`
            Meta map[string]interface{}  `json:"meta"`
            Attributes Attributes  `json:"attributes"`
            L1Category []string  `json:"l1_category"`
            Brand string  `json:"brand"`
            Color string  `json:"color"`
            L3Category float64  `json:"l3_category"`
            BranchURL string  `json:"branch_url"`
            SlugKey string  `json:"slug_key"`
            CanReturn bool  `json:"can_return"`
            BrandID float64  `json:"brand_id"`
            ItemID float64  `json:"item_id"`
            L2CategoryID float64  `json:"l2_category_id"`
            DepartmentID float64  `json:"department_id"`
            Image []string  `json:"image"`
            L3CategoryName string  `json:"l3_category_name"`
            CanCancel bool  `json:"can_cancel"`
            WebstoreProductURL string  `json:"webstore_product_url"`
            Size string  `json:"size"`
         
    }
    
    // BagDetailsPlatformResponse used by Order
    type BagDetailsPlatformResponse struct {

        
            ShipmentID string  `json:"shipment_id"`
            AppliedPromos []map[string]interface{}  `json:"applied_promos"`
            DisplayName string  `json:"display_name"`
            GstDetails BagGSTDetails  `json:"gst_details"`
            Prices Prices  `json:"prices"`
            Status BagReturnableCancelableStatus  `json:"status"`
            AffiliateDetails AffiliateDetails  `json:"affiliate_details"`
            OperationalStatus string  `json:"operational_status"`
            Meta BagMeta  `json:"meta"`
            RestorePromos map[string]interface{}  `json:"restore_promos"`
            OrderingStore Store  `json:"ordering_store"`
            Brand Brand  `json:"brand"`
            SellerIdentifier string  `json:"seller_identifier"`
            EntityType string  `json:"entity_type"`
            BID float64  `json:"b_id"`
            Reasons []map[string]interface{}  `json:"reasons"`
            OrderIntegrationID string  `json:"order_integration_id"`
            Article Article  `json:"article"`
            Dates Dates  `json:"dates"`
            Identifier string  `json:"identifier"`
            ParentPromoBags map[string]interface{}  `json:"parent_promo_bags"`
            LineNumber float64  `json:"line_number"`
            AffiliateBagDetails AffiliateBagDetails  `json:"affiliate_bag_details"`
            JourneyType string  `json:"journey_type"`
            BType string  `json:"b_type"`
            BagStatus []BagStatusHistory  `json:"bag_status"`
            Quantity float64  `json:"quantity"`
            ArticleDetails ArticleDetails  `json:"article_details"`
            BagUpdateTime float64  `json:"bag_update_time"`
            NoOfBagsOrder float64  `json:"no_of_bags_order"`
            OriginalBagList []float64  `json:"original_bag_list"`
            QcRequired interface{}  `json:"qc_required"`
            BagStatusHistory BagStatusHistory  `json:"bag_status_history"`
            FinancialBreakup []FinancialBreakup  `json:"financial_breakup"`
            Item Item  `json:"item"`
            CurrentOperationalStatus BagStatusHistory  `json:"current_operational_status"`
            CurrentStatus BagStatusHistory  `json:"current_status"`
            RestoreCoupon bool  `json:"restore_coupon"`
            Tags []string  `json:"tags"`
         
    }
    
    // ErrorResponse used by Order
    type ErrorResponse struct {

        
            Error string  `json:"error"`
            Message string  `json:"message"`
         
    }
    
    // Page1 used by Order
    type Page1 struct {

        
            Size float64  `json:"size"`
            HasNext bool  `json:"has_next"`
            PageType string  `json:"page_type"`
            Current float64  `json:"current"`
            ItemTotal float64  `json:"item_total"`
         
    }
    
    // GetBagsPlatformResponse used by Order
    type GetBagsPlatformResponse struct {

        
            Items []BagDetailsPlatformResponse  `json:"items"`
            Page Page1  `json:"page"`
         
    }
    

    
    // SearchKeywordResult used by Catalog
    type SearchKeywordResult struct {

        
            SortOn string  `json:"sort_on"`
            Query map[string]interface{}  `json:"query"`
         
    }
    
    // CreateSearchKeyword used by Catalog
    type CreateSearchKeyword struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Result SearchKeywordResult  `json:"result"`
            AppID string  `json:"app_id"`
            IsActive bool  `json:"is_active"`
            Words []string  `json:"words"`
         
    }
    
    // GetSearchWordsData used by Catalog
    type GetSearchWordsData struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Result map[string]interface{}  `json:"result"`
            AppID string  `json:"app_id"`
            UID string  `json:"uid"`
            IsActive bool  `json:"is_active"`
            Words []string  `json:"words"`
         
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
    
    // AutocompletePageAction used by Catalog
    type AutocompletePageAction struct {

        
            URL string  `json:"url"`
            Type string  `json:"type"`
            Params map[string]interface{}  `json:"params"`
            Query map[string]interface{}  `json:"query"`
         
    }
    
    // AutocompleteAction used by Catalog
    type AutocompleteAction struct {

        
            Page AutocompletePageAction  `json:"page"`
            Type string  `json:"type"`
         
    }
    
    // Media used by Catalog
    type Media struct {

        
            AspectRatio string  `json:"aspect_ratio"`
            URL string  `json:"url"`
            Type string  `json:"type"`
         
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

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Results []AutocompleteResult  `json:"results"`
            AppID string  `json:"app_id"`
            IsActive bool  `json:"is_active"`
            Words []string  `json:"words"`
         
    }
    
    // GetAutocompleteWordsData used by Catalog
    type GetAutocompleteWordsData struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Results []map[string]interface{}  `json:"results"`
            AppID string  `json:"app_id"`
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
            Words []string  `json:"words"`
            Results []map[string]interface{}  `json:"results"`
            AppID string  `json:"app_id"`
         
    }
    
    // ProductBundleItem used by Catalog
    type ProductBundleItem struct {

        
            MaxQuantity float64  `json:"max_quantity"`
            ProductUID float64  `json:"product_uid"`
            AutoSelect bool  `json:"auto_select"`
            AllowRemove bool  `json:"allow_remove"`
            AutoAddToCart bool  `json:"auto_add_to_cart"`
            MinQuantity float64  `json:"min_quantity"`
         
    }
    
    // GetProductBundleCreateResponse used by Catalog
    type GetProductBundleCreateResponse struct {

        
            Products []ProductBundleItem  `json:"products"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Name string  `json:"name"`
            ModifiedOn string  `json:"modified_on"`
            PageVisibility []string  `json:"page_visibility"`
            Choice string  `json:"choice"`
            CompanyID float64  `json:"company_id"`
            Meta map[string]interface{}  `json:"meta"`
            Slug string  `json:"slug"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            IsActive bool  `json:"is_active"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            CreatedOn string  `json:"created_on"`
            ID string  `json:"id"`
            Logo string  `json:"logo"`
         
    }
    
    // GetProductBundleListingResponse used by Catalog
    type GetProductBundleListingResponse struct {

        
            Page Page  `json:"page"`
            Items []GetProductBundleCreateResponse  `json:"items"`
         
    }
    
    // ProductBundleRequest used by Catalog
    type ProductBundleRequest struct {

        
            Products []ProductBundleItem  `json:"products"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Name string  `json:"name"`
            ModifiedOn string  `json:"modified_on"`
            PageVisibility []string  `json:"page_visibility"`
            Choice string  `json:"choice"`
            CompanyID float64  `json:"company_id"`
            Meta map[string]interface{}  `json:"meta"`
            Slug string  `json:"slug"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            IsActive bool  `json:"is_active"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            CreatedOn string  `json:"created_on"`
            Logo string  `json:"logo"`
         
    }
    
    // ProductBundleUpdateRequest used by Catalog
    type ProductBundleUpdateRequest struct {

        
            Products []ProductBundleItem  `json:"products"`
            Name string  `json:"name"`
            ModifiedOn string  `json:"modified_on"`
            PageVisibility []string  `json:"page_visibility"`
            Choice string  `json:"choice"`
            CompanyID float64  `json:"company_id"`
            Meta map[string]interface{}  `json:"meta"`
            Slug string  `json:"slug"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            IsActive bool  `json:"is_active"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Logo string  `json:"logo"`
         
    }
    
    // Size used by Catalog
    type Size struct {

        
            IsAvailable bool  `json:"is_available"`
            Value string  `json:"value"`
            Quantity float64  `json:"quantity"`
            Display string  `json:"display"`
         
    }
    
    // LimitedProductData used by Catalog
    type LimitedProductData struct {

        
            Sizes []string  `json:"sizes"`
            ShortDescription string  `json:"short_description"`
            CountryOfOrigin string  `json:"country_of_origin"`
            ItemCode string  `json:"item_code"`
            Name string  `json:"name"`
            Attributes map[string]interface{}  `json:"attributes"`
            Images []string  `json:"images"`
            Slug string  `json:"slug"`
            Quantity float64  `json:"quantity"`
            UID float64  `json:"uid"`
            Identifier map[string]interface{}  `json:"identifier"`
            Price map[string]interface{}  `json:"price"`
         
    }
    
    // Price used by Catalog
    type Price struct {

        
            MinMarked float64  `json:"min_marked"`
            MaxMarked float64  `json:"max_marked"`
            Currency string  `json:"currency"`
            MaxEffective float64  `json:"max_effective"`
            MinEffective float64  `json:"min_effective"`
         
    }
    
    // GetProducts used by Catalog
    type GetProducts struct {

        
            MaxQuantity float64  `json:"max_quantity"`
            Sizes []Size  `json:"sizes"`
            ProductUID float64  `json:"product_uid"`
            AutoSelect bool  `json:"auto_select"`
            AllowRemove bool  `json:"allow_remove"`
            AutoAddToCart bool  `json:"auto_add_to_cart"`
            MinQuantity float64  `json:"min_quantity"`
            ProductDetails LimitedProductData  `json:"product_details"`
            Price Price  `json:"price"`
         
    }
    
    // GetProductBundleResponse used by Catalog
    type GetProductBundleResponse struct {

        
            Products []GetProducts  `json:"products"`
            Name string  `json:"name"`
            IsActive bool  `json:"is_active"`
            Choice string  `json:"choice"`
            CompanyID float64  `json:"company_id"`
            Slug string  `json:"slug"`
            PageVisibility []string  `json:"page_visibility"`
            Meta map[string]interface{}  `json:"meta"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            Logo string  `json:"logo"`
         
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

        
            Subtitle string  `json:"subtitle"`
            Description string  `json:"description"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Guide Guide  `json:"guide"`
            Name string  `json:"name"`
            ModifiedOn string  `json:"modified_on"`
            Title string  `json:"title"`
            CompanyID float64  `json:"company_id"`
            ID string  `json:"id"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Active bool  `json:"active"`
            Image string  `json:"image"`
            CreatedOn string  `json:"created_on"`
            Tag string  `json:"tag"`
            BrandID float64  `json:"brand_id"`
         
    }
    
    // SuccessResponse used by Catalog
    type SuccessResponse struct {

        
            Success bool  `json:"success"`
         
    }
    
    // SizeGuideResponse used by Catalog
    type SizeGuideResponse struct {

        
            CreatedBy map[string]interface{}  `json:"created_by"`
            Subtitle string  `json:"subtitle"`
            Guide map[string]interface{}  `json:"guide"`
            ModifiedOn string  `json:"modified_on"`
            Name string  `json:"name"`
            Title string  `json:"title"`
            CompanyID float64  `json:"company_id"`
            ID string  `json:"id"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Active bool  `json:"active"`
            CreatedOn string  `json:"created_on"`
            Tag string  `json:"tag"`
            BrandID float64  `json:"brand_id"`
         
    }
    
    // ApplicationItemSEO used by Catalog
    type ApplicationItemSEO struct {

        
            Description interface{}  `json:"description"`
            Title interface{}  `json:"title"`
         
    }
    
    // ApplicationItemMOQ used by Catalog
    type ApplicationItemMOQ struct {

        
            IncrementUnit float64  `json:"increment_unit"`
            Maximum float64  `json:"maximum"`
            Minimum float64  `json:"minimum"`
         
    }
    
    // MetaFields used by Catalog
    type MetaFields struct {

        
            Key interface{}  `json:"key"`
            Value interface{}  `json:"value"`
         
    }
    
    // ApplicationItemMeta used by Catalog
    type ApplicationItemMeta struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            IsGift bool  `json:"is_gift"`
            Seo ApplicationItemSEO  `json:"seo"`
            Moq ApplicationItemMOQ  `json:"moq"`
            CustomMeta []MetaFields  `json:"_custom_meta"`
            IsCod bool  `json:"is_cod"`
            AltText map[string]interface{}  `json:"alt_text"`
         
    }
    
    // SuccessResponse1 used by Catalog
    type SuccessResponse1 struct {

        
            UID float64  `json:"uid"`
            Success bool  `json:"success"`
         
    }
    
    // SEOData used by Catalog
    type SEOData struct {

        
            Description interface{}  `json:"description"`
            Title interface{}  `json:"title"`
         
    }
    
    // MOQData used by Catalog
    type MOQData struct {

        
            IncrementUnit float64  `json:"increment_unit"`
            Maximum float64  `json:"maximum"`
            Minimum float64  `json:"minimum"`
         
    }
    
    // OwnerAppItemResponse used by Catalog
    type OwnerAppItemResponse struct {

        
            IsGift bool  `json:"is_gift"`
            Seo SEOData  `json:"seo"`
            Moq MOQData  `json:"moq"`
            IsCod bool  `json:"is_cod"`
            AltText map[string]interface{}  `json:"alt_text"`
         
    }
    
    // GetConfigMetadataResponse used by Catalog
    type GetConfigMetadataResponse struct {

        
            Condition []map[string]interface{}  `json:"condition"`
            Data []map[string]interface{}  `json:"data"`
            Values []map[string]interface{}  `json:"values"`
         
    }
    
    // PageResponseType used by Catalog
    type PageResponseType struct {

        
            HasNext bool  `json:"has_next"`
            TotalCount float64  `json:"total_count"`
            Next float64  `json:"next"`
            Current float64  `json:"current"`
         
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
            Unit string  `json:"unit"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Priority float64  `json:"priority"`
            IsActive bool  `json:"is_active"`
            Key string  `json:"key"`
            Logo string  `json:"logo"`
         
    }
    
    // AppConfigurationDetail used by Catalog
    type AppConfigurationDetail struct {

        
            IsDefault bool  `json:"is_default"`
            Priority float64  `json:"priority"`
            TemplateSlugs []string  `json:"template_slugs"`
            Name string  `json:"name"`
            Attributes []AttributeDetailsGroup  `json:"attributes"`
            Slug string  `json:"slug"`
            AppID string  `json:"app_id"`
            IsActive bool  `json:"is_active"`
            Logo string  `json:"logo"`
         
    }
    
    // ConfigSuccessResponse used by Catalog
    type ConfigSuccessResponse struct {

        
            Message string  `json:"message"`
         
    }
    
    // AppConfigurationsSort used by Catalog
    type AppConfigurationsSort struct {

        
            IsDefault bool  `json:"is_default"`
            DefaultKey string  `json:"default_key"`
            Priority float64  `json:"priority"`
            Name string  `json:"name"`
            AppID string  `json:"app_id"`
            IsActive bool  `json:"is_active"`
            Key string  `json:"key"`
            Logo string  `json:"logo"`
         
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
            Compare map[string]interface{}  `json:"compare"`
            Detail map[string]interface{}  `json:"detail"`
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

        
            Start float64  `json:"start"`
            End float64  `json:"end"`
            Display string  `json:"display"`
         
    }
    
    // ConfigurationListingFilterValue used by Catalog
    type ConfigurationListingFilterValue struct {

        
            Value string  `json:"value"`
            Map map[string]interface{}  `json:"map"`
            MapValues []map[string]interface{}  `json:"map_values"`
            Condition string  `json:"condition"`
            Sort string  `json:"sort"`
            BucketPoints []ConfigurationBucketPoints  `json:"bucket_points"`
         
    }
    
    // ConfigurationListingFilterConfig used by Catalog
    type ConfigurationListingFilterConfig struct {

        
            DisplayName string  `json:"display_name"`
            Type string  `json:"type"`
            Name string  `json:"name"`
            ValueConfig ConfigurationListingFilterValue  `json:"value_config"`
            Priority float64  `json:"priority"`
            IsActive bool  `json:"is_active"`
            Key string  `json:"key"`
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
    
    // ProductSize used by Catalog
    type ProductSize struct {

        
            Max float64  `json:"max"`
            Min float64  `json:"min"`
         
    }
    
    // ConfigurationProductVariantConfig used by Catalog
    type ConfigurationProductVariantConfig struct {

        
            DisplayType string  `json:"display_type"`
            Name string  `json:"name"`
            Priority float64  `json:"priority"`
            IsActive bool  `json:"is_active"`
            Size ProductSize  `json:"size"`
            Key string  `json:"key"`
            Logo string  `json:"logo"`
         
    }
    
    // ConfigurationProductVariant used by Catalog
    type ConfigurationProductVariant struct {

        
            Config []ConfigurationProductVariantConfig  `json:"config"`
         
    }
    
    // ConfigurationProductConfig used by Catalog
    type ConfigurationProductConfig struct {

        
            Subtitle string  `json:"subtitle"`
            Title string  `json:"title"`
            Priority float64  `json:"priority"`
            IsActive bool  `json:"is_active"`
            Size ProductSize  `json:"size"`
            Key string  `json:"key"`
            Logo string  `json:"logo"`
         
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
            CreatedBy map[string]interface{}  `json:"created_by"`
            Product ConfigurationProduct  `json:"product"`
            Type string  `json:"type"`
            ConfigID string  `json:"config_id"`
            ModifiedOn string  `json:"modified_on"`
            ID string  `json:"id"`
            AppID string  `json:"app_id"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
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

        
            Listing ConfigurationListing  `json:"listing"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Product ConfigurationProduct  `json:"product"`
            Type string  `json:"type"`
            ConfigID string  `json:"config_id"`
            ModifiedOn string  `json:"modified_on"`
            AppID string  `json:"app_id"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
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

        
            Listing GetCatalogConfigurationDetailsSchemaListing  `json:"listing"`
            Product GetCatalogConfigurationDetailsProduct  `json:"product"`
            ConfigID string  `json:"config_id"`
            AppID string  `json:"app_id"`
            ID string  `json:"id"`
            ConfigType string  `json:"config_type"`
         
    }
    
    // GetAppCatalogEntityConfiguration used by Catalog
    type GetAppCatalogEntityConfiguration struct {

        
            Data EntityConfiguration  `json:"data"`
            IsDefault bool  `json:"is_default"`
         
    }
    
    // ProductSortOn used by Catalog
    type ProductSortOn struct {

        
            Name string  `json:"name"`
            IsSelected bool  `json:"is_selected"`
            Value string  `json:"value"`
         
    }
    
    // ProductFiltersValue used by Catalog
    type ProductFiltersValue struct {

        
            IsSelected bool  `json:"is_selected"`
            Count float64  `json:"count"`
            SelectedMax float64  `json:"selected_max"`
            Value interface{}  `json:"value"`
            Display string  `json:"display"`
            DisplayFormat string  `json:"display_format"`
            CurrencyCode string  `json:"currency_code"`
            Min float64  `json:"min"`
            SelectedMin float64  `json:"selected_min"`
            CurrencySymbol string  `json:"currency_symbol"`
            Max float64  `json:"max"`
            QueryFormat string  `json:"query_format"`
         
    }
    
    // ProductFiltersKey used by Catalog
    type ProductFiltersKey struct {

        
            Display string  `json:"display"`
            Name string  `json:"name"`
            Kind string  `json:"kind"`
            Operators []string  `json:"operators"`
            Logo string  `json:"logo"`
         
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

        
            Portrait BannerImage  `json:"portrait"`
            Landscape BannerImage  `json:"landscape"`
         
    }
    
    // CollectionQuery used by Catalog
    type CollectionQuery struct {

        
            Op string  `json:"op"`
            Attribute string  `json:"attribute"`
            Value []interface{}  `json:"value"`
         
    }
    
    // GetCollectionDetailNest used by Catalog
    type GetCollectionDetailNest struct {

        
            Type string  `json:"type"`
            AppID string  `json:"app_id"`
            Description string  `json:"description"`
            AllowFacets bool  `json:"allow_facets"`
            Schedule map[string]interface{}  `json:"_schedule"`
            Cron map[string]interface{}  `json:"cron"`
            Action Action  `json:"action"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Priority float64  `json:"priority"`
            UID string  `json:"uid"`
            Name string  `json:"name"`
            Tag []string  `json:"tag"`
            Logo Media1  `json:"logo"`
            Banners ImageUrls  `json:"banners"`
            AllowSort bool  `json:"allow_sort"`
            Meta map[string]interface{}  `json:"meta"`
            Slug string  `json:"slug"`
            Badge map[string]interface{}  `json:"badge"`
            IsActive bool  `json:"is_active"`
            Query []CollectionQuery  `json:"query"`
         
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
    
    // UserInfo used by Catalog
    type UserInfo struct {

        
            UID string  `json:"uid"`
            UserID string  `json:"user_id"`
            Email string  `json:"email"`
            Username string  `json:"username"`
         
    }
    
    // CollectionSchedule used by Catalog
    type CollectionSchedule struct {

        
            Start string  `json:"start"`
            Cron string  `json:"cron"`
            NextSchedule []NextSchedule  `json:"next_schedule"`
            End string  `json:"end"`
            Duration float64  `json:"duration"`
         
    }
    
    // SeoDetail used by Catalog
    type SeoDetail struct {

        
            Description string  `json:"description"`
            Title string  `json:"title"`
         
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
    
    // CollectionBadge used by Catalog
    type CollectionBadge struct {

        
            Text string  `json:"text"`
            Color string  `json:"color"`
         
    }
    
    // CreateCollection used by Catalog
    type CreateCollection struct {

        
            CreatedBy UserInfo  `json:"created_by"`
            SortOn string  `json:"sort_on"`
            Type string  `json:"type"`
            AppID string  `json:"app_id"`
            IsVisible bool  `json:"is_visible"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            Description string  `json:"description"`
            AllowFacets bool  `json:"allow_facets"`
            Schedule CollectionSchedule  `json:"_schedule"`
            Seo SeoDetail  `json:"seo"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Priority float64  `json:"priority"`
            Published bool  `json:"published"`
            Name string  `json:"name"`
            ModifiedBy UserInfo  `json:"modified_by"`
            Logo CollectionImage  `json:"logo"`
            Banners CollectionBanner  `json:"banners"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            AllowSort bool  `json:"allow_sort"`
            Tags []string  `json:"tags"`
            Meta map[string]interface{}  `json:"meta"`
            Slug string  `json:"slug"`
            Badge CollectionBadge  `json:"badge"`
            IsActive bool  `json:"is_active"`
            Query []CollectionQuery  `json:"query"`
         
    }
    
    // CollectionCreateResponse used by Catalog
    type CollectionCreateResponse struct {

        
            SortOn string  `json:"sort_on"`
            Type string  `json:"type"`
            AppID string  `json:"app_id"`
            Description string  `json:"description"`
            AllowFacets bool  `json:"allow_facets"`
            Schedule map[string]interface{}  `json:"_schedule"`
            Cron map[string]interface{}  `json:"cron"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Priority float64  `json:"priority"`
            Name string  `json:"name"`
            Tag []string  `json:"tag"`
            Logo BannerImage  `json:"logo"`
            Banners ImageUrls  `json:"banners"`
            AllowSort bool  `json:"allow_sort"`
            Meta map[string]interface{}  `json:"meta"`
            Slug string  `json:"slug"`
            Badge map[string]interface{}  `json:"badge"`
            IsActive bool  `json:"is_active"`
            Query []CollectionQuery  `json:"query"`
         
    }
    
    // CollectionDetailResponse used by Catalog
    type CollectionDetailResponse struct {

        
            Banners ImageUrls  `json:"banners"`
            Description string  `json:"description"`
            AllowFacets bool  `json:"allow_facets"`
            Schedule map[string]interface{}  `json:"_schedule"`
            Priority float64  `json:"priority"`
            Cron map[string]interface{}  `json:"cron"`
            Badge map[string]interface{}  `json:"badge"`
            Type string  `json:"type"`
            Name string  `json:"name"`
            AllowSort bool  `json:"allow_sort"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Slug string  `json:"slug"`
            AppID string  `json:"app_id"`
            Meta map[string]interface{}  `json:"meta"`
            IsActive bool  `json:"is_active"`
            Tag []string  `json:"tag"`
            Query []CollectionQuery  `json:"query"`
            Logo Media1  `json:"logo"`
         
    }
    
    // UpdateCollection used by Catalog
    type UpdateCollection struct {

        
            SortOn string  `json:"sort_on"`
            Type string  `json:"type"`
            IsVisible bool  `json:"is_visible"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            Description string  `json:"description"`
            AllowFacets bool  `json:"allow_facets"`
            Schedule CollectionSchedule  `json:"_schedule"`
            Seo SeoDetail  `json:"seo"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Priority float64  `json:"priority"`
            Published bool  `json:"published"`
            Name string  `json:"name"`
            ModifiedBy UserInfo  `json:"modified_by"`
            Logo CollectionImage  `json:"logo"`
            Banners CollectionBanner  `json:"banners"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            AllowSort bool  `json:"allow_sort"`
            Tags []string  `json:"tags"`
            Meta map[string]interface{}  `json:"meta"`
            Slug string  `json:"slug"`
            Badge CollectionBadge  `json:"badge"`
            IsActive bool  `json:"is_active"`
            Query []CollectionQuery  `json:"query"`
         
    }
    
    // ProductBrand used by Catalog
    type ProductBrand struct {

        
            Name string  `json:"name"`
            Action Action  `json:"action"`
            UID float64  `json:"uid"`
            Logo Media1  `json:"logo"`
         
    }
    
    // ProductDetailAttribute used by Catalog
    type ProductDetailAttribute struct {

        
            Key string  `json:"key"`
            Value string  `json:"value"`
            Type string  `json:"type"`
         
    }
    
    // ProductDetailGroupedAttribute used by Catalog
    type ProductDetailGroupedAttribute struct {

        
            Title string  `json:"title"`
            Details []ProductDetailAttribute  `json:"details"`
         
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

        
            Effective Price1  `json:"effective"`
            Marked Price1  `json:"marked"`
         
    }
    
    // ProductListingDetail used by Catalog
    type ProductListingDetail struct {

        
            Color string  `json:"color"`
            Type string  `json:"type"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            Attributes map[string]interface{}  `json:"attributes"`
            Description string  `json:"description"`
            Sellable bool  `json:"sellable"`
            RatingCount float64  `json:"rating_count"`
            ImageNature string  `json:"image_nature"`
            UID float64  `json:"uid"`
            Rating float64  `json:"rating"`
            Discount string  `json:"discount"`
            PromoMeta map[string]interface{}  `json:"promo_meta"`
            ItemCode string  `json:"item_code"`
            Brand ProductBrand  `json:"brand"`
            Name string  `json:"name"`
            Similars []string  `json:"similars"`
            HasVariant bool  `json:"has_variant"`
            Medias []Media1  `json:"medias"`
            ShortDescription string  `json:"short_description"`
            ItemType string  `json:"item_type"`
            Slug string  `json:"slug"`
            Tryouts []string  `json:"tryouts"`
            ProductOnlineDate string  `json:"product_online_date"`
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
            Highlights []string  `json:"highlights"`
            Price ProductListingPrice  `json:"price"`
         
    }
    
    // GetCollectionItemsResponse used by Catalog
    type GetCollectionItemsResponse struct {

        
            Page Page  `json:"page"`
            Items []ProductListingDetail  `json:"items"`
            SortOn []ProductSortOn  `json:"sort_on"`
            Filters []ProductFilters  `json:"filters"`
         
    }
    
    // ItemQueryForUserCollection used by Catalog
    type ItemQueryForUserCollection struct {

        
            ItemID float64  `json:"item_id"`
            Action string  `json:"action"`
         
    }
    
    // CollectionItemRequest used by Catalog
    type CollectionItemRequest struct {

        
            Query []CollectionQuery  `json:"query"`
            Type string  `json:"type"`
            Item []ItemQueryForUserCollection  `json:"item"`
         
    }
    
    // UpdatedResponse used by Catalog
    type UpdatedResponse struct {

        
            Message string  `json:"message"`
            ItemsNotUpdated []float64  `json:"items_not_updated"`
         
    }
    
    // CatalogInsightBrand used by Catalog
    type CatalogInsightBrand struct {

        
            AvailableArticles float64  `json:"available_articles"`
            TotalSizes float64  `json:"total_sizes"`
            ArticleFreshness float64  `json:"article_freshness"`
            TotalArticles float64  `json:"total_articles"`
            Name string  `json:"name"`
            AvailableSizes float64  `json:"available_sizes"`
         
    }
    
    // CatalogInsightItem used by Catalog
    type CatalogInsightItem struct {

        
            OutOfStockCount float64  `json:"out_of_stock_count"`
            Count float64  `json:"count"`
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

        
            Platform string  `json:"platform"`
            Enabled bool  `json:"enabled"`
            OptLevel string  `json:"opt_level"`
            CompanyID float64  `json:"company_id"`
            BrandIds []float64  `json:"brand_ids"`
            StoreIds []float64  `json:"store_ids"`
         
    }
    
    // CompanyOptIn used by Catalog
    type CompanyOptIn struct {

        
            CreatedBy map[string]interface{}  `json:"created_by"`
            Platform string  `json:"platform"`
            Enabled bool  `json:"enabled"`
            OptLevel string  `json:"opt_level"`
            ModifiedOn float64  `json:"modified_on"`
            CompanyID float64  `json:"company_id"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            BrandIds []float64  `json:"brand_ids"`
            StoreIds []float64  `json:"store_ids"`
            CreatedOn float64  `json:"created_on"`
         
    }
    
    // GetOptInPlatform used by Catalog
    type GetOptInPlatform struct {

        
            Page Page  `json:"page"`
            Items []CompanyOptIn  `json:"items"`
         
    }
    
    // OptinCompanyDetail used by Catalog
    type OptinCompanyDetail struct {

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            CompanyType string  `json:"company_type"`
            BusinessType string  `json:"business_type"`
         
    }
    
    // CompanyBrandDetail used by Catalog
    type CompanyBrandDetail struct {

        
            TotalArticle float64  `json:"total_article"`
            CompanyID float64  `json:"company_id"`
            BrandName string  `json:"brand_name"`
            BrandID float64  `json:"brand_id"`
         
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

        
            Manager map[string]interface{}  `json:"manager"`
            Documents []map[string]interface{}  `json:"documents"`
            StoreType string  `json:"store_type"`
            DisplayName string  `json:"display_name"`
            AdditionalContacts []map[string]interface{}  `json:"additional_contacts"`
            Address map[string]interface{}  `json:"address"`
            Name string  `json:"name"`
            ModifiedOn string  `json:"modified_on"`
            CompanyID float64  `json:"company_id"`
            StoreCode string  `json:"store_code"`
            UID float64  `json:"uid"`
            Timing map[string]interface{}  `json:"timing"`
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
    
    // AttributeSchemaRange used by Catalog
    type AttributeSchemaRange struct {

        
            Max float64  `json:"max"`
            Min float64  `json:"min"`
         
    }
    
    // AttributeMaster used by Catalog
    type AttributeMaster struct {

        
            Range AttributeSchemaRange  `json:"range"`
            Format string  `json:"format"`
            AllowedValues []string  `json:"allowed_values"`
            Type string  `json:"type"`
            Mandatory bool  `json:"mandatory"`
            Multi bool  `json:"multi"`
         
    }
    
    // GenderDetail used by Catalog
    type GenderDetail struct {

        
            Description string  `json:"description"`
            Details AttributeMasterDetails  `json:"details"`
            EnabledForEndConsumer bool  `json:"enabled_for_end_consumer"`
            Filters AttributeMasterFilter  `json:"filters"`
            Name string  `json:"name"`
            Meta AttributeMasterMeta  `json:"meta"`
            Schema AttributeMaster  `json:"schema"`
            Slug string  `json:"slug"`
            ID string  `json:"id"`
            IsNested bool  `json:"is_nested"`
            Departments []string  `json:"departments"`
            Logo string  `json:"logo"`
         
    }
    
    // ProdcutTemplateCategoriesResponse used by Catalog
    type ProdcutTemplateCategoriesResponse struct {

        
            Page Page  `json:"page"`
            Items []map[string]interface{}  `json:"items"`
         
    }
    
    // PTErrorResponse used by Catalog
    type PTErrorResponse struct {

        
            Message string  `json:"message"`
            Errors map[string]interface{}  `json:"errors"`
            Meta map[string]interface{}  `json:"meta"`
            Status float64  `json:"status"`
            Code string  `json:"code"`
         
    }
    
    // UserSerializer used by Catalog
    type UserSerializer struct {

        
            Username string  `json:"username"`
            ID string  `json:"_id"`
            Contact string  `json:"contact"`
            UserID string  `json:"user_id"`
            UID string  `json:"uid"`
         
    }
    
    // GetDepartment used by Catalog
    type GetDepartment struct {

        
            PageNo float64  `json:"page_no"`
            CreatedBy UserSerializer  `json:"created_by"`
            PriorityOrder float64  `json:"priority_order"`
            Name string  `json:"name"`
            ModifiedOn string  `json:"modified_on"`
            Search string  `json:"search"`
            CreatedOn string  `json:"created_on"`
            ItemType string  `json:"item_type"`
            Slug string  `json:"slug"`
            PageSize float64  `json:"page_size"`
            UID float64  `json:"uid"`
            IsActive bool  `json:"is_active"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            Synonyms []string  `json:"synonyms"`
            Logo string  `json:"logo"`
         
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
            Meta map[string]interface{}  `json:"meta"`
            Status float64  `json:"status"`
            Code string  `json:"code"`
         
    }
    
    // DepartmentCreateUpdate used by Catalog
    type DepartmentCreateUpdate struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            PriorityOrder float64  `json:"priority_order"`
            Cls string  `json:"_cls"`
            Name string  `json:"name"`
            Platforms map[string]interface{}  `json:"platforms"`
            Tags []string  `json:"tags"`
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
            IsActive bool  `json:"is_active"`
            Synonyms []string  `json:"synonyms"`
            Logo string  `json:"logo"`
         
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
            UserID string  `json:"user_id"`
            Contact string  `json:"contact"`
            SuperUser bool  `json:"super_user"`
         
    }
    
    // DepartmentModel used by Catalog
    type DepartmentModel struct {

        
            ID interface{}  `json:"_id"`
            CreatedBy UserDetail  `json:"created_by"`
            Synonyms []interface{}  `json:"synonyms"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            PriorityOrder float64  `json:"priority_order"`
            Cls interface{}  `json:"_cls"`
            ModifiedOn string  `json:"modified_on"`
            VerifiedBy UserDetail  `json:"verified_by"`
            Name interface{}  `json:"name"`
            VerifiedOn string  `json:"verified_on"`
            Slug interface{}  `json:"slug"`
            ModifiedBy UserDetail  `json:"modified_by"`
            IsActive bool  `json:"is_active"`
            CreatedOn string  `json:"created_on"`
            UID float64  `json:"uid"`
            Logo string  `json:"logo"`
         
    }
    
    // ProductTemplate used by Catalog
    type ProductTemplate struct {

        
            Description string  `json:"description"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            IsPhysical bool  `json:"is_physical"`
            Categories []string  `json:"categories"`
            Name string  `json:"name"`
            ModifiedOn string  `json:"modified_on"`
            Attributes []string  `json:"attributes"`
            Departments []string  `json:"departments"`
            Slug string  `json:"slug"`
            IsArchived bool  `json:"is_archived"`
            IsExpirable bool  `json:"is_expirable"`
            IsActive bool  `json:"is_active"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            CreatedOn string  `json:"created_on"`
            Tag string  `json:"tag"`
            Logo string  `json:"logo"`
         
    }
    
    // TemplatesResponse used by Catalog
    type TemplatesResponse struct {

        
            Page Page  `json:"page"`
            Items ProductTemplate  `json:"items"`
         
    }
    
    // Properties used by Catalog
    type Properties struct {

        
            Media map[string]interface{}  `json:"media"`
            MultiSize map[string]interface{}  `json:"multi_size"`
            ProductGroupTag map[string]interface{}  `json:"product_group_tag"`
            HsnCode map[string]interface{}  `json:"hsn_code"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            Command map[string]interface{}  `json:"command"`
            TraderType map[string]interface{}  `json:"trader_type"`
            Description map[string]interface{}  `json:"description"`
            SizeGuide map[string]interface{}  `json:"size_guide"`
            ProductPublish map[string]interface{}  `json:"product_publish"`
            IsActive map[string]interface{}  `json:"is_active"`
            Sizes map[string]interface{}  `json:"sizes"`
            CountryOfOrigin map[string]interface{}  `json:"country_of_origin"`
            CategorySlug map[string]interface{}  `json:"category_slug"`
            ItemCode map[string]interface{}  `json:"item_code"`
            Trader map[string]interface{}  `json:"trader"`
            Name map[string]interface{}  `json:"name"`
            BrandUID map[string]interface{}  `json:"brand_uid"`
            IsDependent map[string]interface{}  `json:"is_dependent"`
            ShortDescription map[string]interface{}  `json:"short_description"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            ReturnConfig map[string]interface{}  `json:"return_config"`
            Currency map[string]interface{}  `json:"currency"`
            Tags map[string]interface{}  `json:"tags"`
            NoOfBoxes map[string]interface{}  `json:"no_of_boxes"`
            ItemType map[string]interface{}  `json:"item_type"`
            Slug map[string]interface{}  `json:"slug"`
            Variants map[string]interface{}  `json:"variants"`
            Highlights map[string]interface{}  `json:"highlights"`
         
    }
    
    // GlobalValidation used by Catalog
    type GlobalValidation struct {

        
            Description string  `json:"description"`
            Required []string  `json:"required"`
            Type string  `json:"type"`
            Title string  `json:"title"`
            Definitions map[string]interface{}  `json:"definitions"`
            Properties Properties  `json:"properties"`
         
    }
    
    // TemplateValidationData used by Catalog
    type TemplateValidationData struct {

        
            TemplateValidation map[string]interface{}  `json:"template_validation"`
            GlobalValidation GlobalValidation  `json:"global_validation"`
         
    }
    
    // TemplateDetails used by Catalog
    type TemplateDetails struct {

        
            Description string  `json:"description"`
            IsPhysical bool  `json:"is_physical"`
            Categories []string  `json:"categories"`
            Name string  `json:"name"`
            Attributes []string  `json:"attributes"`
            ID string  `json:"id"`
            Departments []string  `json:"departments"`
            Slug string  `json:"slug"`
            IsArchived bool  `json:"is_archived"`
            IsExpirable bool  `json:"is_expirable"`
            IsActive bool  `json:"is_active"`
            Tag string  `json:"tag"`
            Logo string  `json:"logo"`
         
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
    
    // VerifiedBy used by Catalog
    type VerifiedBy struct {

        
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
         
    }
    
    // ProductDownloadItemsData used by Catalog
    type ProductDownloadItemsData struct {

        
            Templates []string  `json:"templates"`
            Brand []string  `json:"brand"`
            Type string  `json:"type"`
         
    }
    
    // ProductDownloadsItems used by Catalog
    type ProductDownloadsItems struct {

        
            TaskID string  `json:"task_id"`
            CreatedBy VerifiedBy  `json:"created_by"`
            URL string  `json:"url"`
            CompletedOn string  `json:"completed_on"`
            TriggerOn string  `json:"trigger_on"`
            Data ProductDownloadItemsData  `json:"data"`
            TemplateTags map[string]interface{}  `json:"template_tags"`
            ID string  `json:"id"`
            SellerID float64  `json:"seller_id"`
            Status string  `json:"status"`
         
    }
    
    // ProductDownloadsResponse used by Catalog
    type ProductDownloadsResponse struct {

        
            Page Page  `json:"page"`
            Items ProductDownloadsItems  `json:"items"`
         
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
    
    // Hierarchy used by Catalog
    type Hierarchy struct {

        
            L1 float64  `json:"l1"`
            L2 float64  `json:"l2"`
            Department float64  `json:"department"`
         
    }
    
    // Category used by Catalog
    type Category struct {

        
            Media Media2  `json:"media"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Level float64  `json:"level"`
            Name string  `json:"name"`
            IsActive bool  `json:"is_active"`
            ModifiedOn string  `json:"modified_on"`
            ID string  `json:"id"`
            CreatedOn string  `json:"created_on"`
            Slug string  `json:"slug"`
            Tryouts []string  `json:"tryouts"`
            UID float64  `json:"uid"`
            Priority float64  `json:"priority"`
            Marketplaces CategoryMapping  `json:"marketplaces"`
            Synonyms []string  `json:"synonyms"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Departments []float64  `json:"departments"`
            Hierarchy []Hierarchy  `json:"hierarchy"`
         
    }
    
    // CategoryResponse used by Catalog
    type CategoryResponse struct {

        
            Page Page  `json:"page"`
            Items []Category  `json:"items"`
         
    }
    
    // CategoryRequestBody used by Catalog
    type CategoryRequestBody struct {

        
            Media Media2  `json:"media"`
            Level float64  `json:"level"`
            Name string  `json:"name"`
            IsActive bool  `json:"is_active"`
            Slug string  `json:"slug"`
            Tryouts []string  `json:"tryouts"`
            Priority float64  `json:"priority"`
            Marketplaces CategoryMapping  `json:"marketplaces"`
            Synonyms []string  `json:"synonyms"`
            Departments []float64  `json:"departments"`
            Hierarchy []Hierarchy  `json:"hierarchy"`
         
    }
    
    // CategoryCreateResponse used by Catalog
    type CategoryCreateResponse struct {

        
            UID float64  `json:"uid"`
            Message string  `json:"message"`
         
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
    
    // ProductPublished used by Catalog
    type ProductPublished struct {

        
            IsSet bool  `json:"is_set"`
            ProductOnlineDate float64  `json:"product_online_date"`
         
    }
    
    // Logo used by Catalog
    type Logo struct {

        
            AspectRatio string  `json:"aspect_ratio"`
            AspectRatioF float64  `json:"aspect_ratio_f"`
            URL string  `json:"url"`
            SecureURL string  `json:"secure_url"`
         
    }
    
    // Image used by Catalog
    type Image struct {

        
            AspectRatio string  `json:"aspect_ratio"`
            AspectRatioF float64  `json:"aspect_ratio_f"`
            URL string  `json:"url"`
            SecureURL string  `json:"secure_url"`
         
    }
    
    // Product used by Catalog
    type Product struct {

        
            Media []Media1  `json:"media"`
            TemplateTag string  `json:"template_tag"`
            IsPhysical bool  `json:"is_physical"`
            MultiSize bool  `json:"multi_size"`
            Color string  `json:"color"`
            HsnCode string  `json:"hsn_code"`
            L3Mapping []string  `json:"l3_mapping"`
            Description string  `json:"description"`
            SizeGuide string  `json:"size_guide"`
            AllSizes []map[string]interface{}  `json:"all_sizes"`
            ProductPublish ProductPublished  `json:"product_publish"`
            ImageNature string  `json:"image_nature"`
            IsActive bool  `json:"is_active"`
            UID float64  `json:"uid"`
            VariantGroup map[string]interface{}  `json:"variant_group"`
            IsSet bool  `json:"is_set"`
            Sizes []map[string]interface{}  `json:"sizes"`
            CountryOfOrigin string  `json:"country_of_origin"`
            CategorySlug string  `json:"category_slug"`
            PrimaryColor string  `json:"primary_color"`
            ItemCode string  `json:"item_code"`
            Brand Brand  `json:"brand"`
            Name string  `json:"name"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            CategoryUID float64  `json:"category_uid"`
            Images []Image  `json:"images"`
            BrandUID float64  `json:"brand_uid"`
            Departments []float64  `json:"departments"`
            IsDependent bool  `json:"is_dependent"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ShortDescription string  `json:"short_description"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            Currency string  `json:"currency"`
            ItemType string  `json:"item_type"`
            Slug string  `json:"slug"`
            ID string  `json:"id"`
            IsExpirable bool  `json:"is_expirable"`
            Variants map[string]interface{}  `json:"variants"`
            Highlights []string  `json:"highlights"`
         
    }
    
    // ProductListingResponse used by Catalog
    type ProductListingResponse struct {

        
            Page Page  `json:"page"`
            Items []Product  `json:"items"`
         
    }
    
    // TeaserTag used by Catalog
    type TeaserTag struct {

        
            URL string  `json:"url"`
            Tag string  `json:"tag"`
         
    }
    
    // ProductPublish used by Catalog
    type ProductPublish struct {

        
            IsSet bool  `json:"is_set"`
            ProductOnlineDate string  `json:"product_online_date"`
         
    }
    
    // NetQuantity used by Catalog
    type NetQuantity struct {

        
            Value float64  `json:"value"`
            Unit interface{}  `json:"unit"`
         
    }
    
    // Trader used by Catalog
    type Trader struct {

        
            Name interface{}  `json:"name"`
            Address []string  `json:"address"`
            Type string  `json:"type"`
         
    }
    
    // TaxIdentifier used by Catalog
    type TaxIdentifier struct {

        
            ReportingHsn string  `json:"reporting_hsn"`
            HsnCodeID string  `json:"hsn_code_id"`
            HsnCode string  `json:"hsn_code"`
         
    }
    
    // CustomOrder used by Catalog
    type CustomOrder struct {

        
            ManufacturingTimeUnit string  `json:"manufacturing_time_unit"`
            IsCustomOrder bool  `json:"is_custom_order"`
            ManufacturingTime float64  `json:"manufacturing_time"`
         
    }
    
    // ProductCreateUpdateRequest used by Catalog
    type ProductCreateUpdateRequest struct {

        
            Media []Media1  `json:"media"`
            ChangeRequestID interface{}  `json:"change_request_id"`
            TemplateTag string  `json:"template_tag"`
            MultiSize bool  `json:"multi_size"`
            ProductGroupTag []string  `json:"product_group_tag"`
            TeaserTag TeaserTag  `json:"teaser_tag"`
            BulkJobID string  `json:"bulk_job_id"`
            Description string  `json:"description"`
            SizeGuide string  `json:"size_guide"`
            VariantsGroup map[string]interface{}  `json:"variants_group"`
            ProductPublish ProductPublish  `json:"product_publish"`
            Action string  `json:"action"`
            VariantMedia map[string]interface{}  `json:"variant_media"`
            Requester string  `json:"requester"`
            IsActive bool  `json:"is_active"`
            CompanyID float64  `json:"company_id"`
            UID float64  `json:"uid"`
            IsSet bool  `json:"is_set"`
            NetQuantity NetQuantity  `json:"net_quantity"`
            CategorySlug string  `json:"category_slug"`
            CountryOfOrigin string  `json:"country_of_origin"`
            ItemCode interface{}  `json:"item_code"`
            Trader []Trader  `json:"trader"`
            Name interface{}  `json:"name"`
            TaxIdentifier TaxIdentifier  `json:"tax_identifier"`
            BrandUID float64  `json:"brand_uid"`
            IsDependent bool  `json:"is_dependent"`
            IsImageLessProduct bool  `json:"is_image_less_product"`
            Departments []float64  `json:"departments"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ShortDescription string  `json:"short_description"`
            CustomOrder CustomOrder  `json:"custom_order"`
            ReturnConfig ReturnConfig  `json:"return_config"`
            Currency string  `json:"currency"`
            Tags []string  `json:"tags"`
            NoOfBoxes float64  `json:"no_of_boxes"`
            ItemType string  `json:"item_type"`
            Slug string  `json:"slug"`
            Variants map[string]interface{}  `json:"variants"`
            Highlights []string  `json:"highlights"`
         
    }
    
    // ProductCreateUpdateSchemaV2 used by Catalog
    type ProductCreateUpdateSchemaV2 struct {

        
            Media []Media1  `json:"media"`
            ChangeRequestID interface{}  `json:"change_request_id"`
            TemplateTag string  `json:"template_tag"`
            MultiSize bool  `json:"multi_size"`
            ProductGroupTag []string  `json:"product_group_tag"`
            TeaserTag TeaserTag  `json:"teaser_tag"`
            Attributes map[string]interface{}  `json:"attributes"`
            BulkJobID string  `json:"bulk_job_id"`
            Description string  `json:"description"`
            SizeGuide string  `json:"size_guide"`
            VariantsGroup map[string]interface{}  `json:"variants_group"`
            ProductPublish ProductPublish  `json:"product_publish"`
            Action string  `json:"action"`
            VariantMedia map[string]interface{}  `json:"variant_media"`
            Requester string  `json:"requester"`
            IsActive bool  `json:"is_active"`
            CompanyID float64  `json:"company_id"`
            UID float64  `json:"uid"`
            IsSet bool  `json:"is_set"`
            NetQuantity NetQuantity  `json:"net_quantity"`
            CategorySlug string  `json:"category_slug"`
            CountryOfOrigin string  `json:"country_of_origin"`
            ItemCode interface{}  `json:"item_code"`
            Trader []Trader  `json:"trader"`
            Name interface{}  `json:"name"`
            TaxIdentifier TaxIdentifier  `json:"tax_identifier"`
            BrandUID float64  `json:"brand_uid"`
            IsDependent bool  `json:"is_dependent"`
            IsImageLessProduct bool  `json:"is_image_less_product"`
            Departments []float64  `json:"departments"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ShortDescription string  `json:"short_description"`
            CustomOrder CustomOrder  `json:"custom_order"`
            ReturnConfig ReturnConfig  `json:"return_config"`
            Currency string  `json:"currency"`
            Tags []string  `json:"tags"`
            NoOfBoxes float64  `json:"no_of_boxes"`
            ItemType string  `json:"item_type"`
            Slug string  `json:"slug"`
            Variants map[string]interface{}  `json:"variants"`
            Highlights []string  `json:"highlights"`
         
    }
    
    // ProductVariants used by Catalog
    type ProductVariants struct {

        
            Media []Media1  `json:"media"`
            ItemCode string  `json:"item_code"`
            Name string  `json:"name"`
            CategoryUID float64  `json:"category_uid"`
            BrandUID float64  `json:"brand_uid"`
            UID float64  `json:"uid"`
         
    }
    
    // ProductVariantsResponse used by Catalog
    type ProductVariantsResponse struct {

        
            Variants []ProductVariants  `json:"variants"`
            Page Page  `json:"page"`
         
    }
    
    // AttributeMasterSerializer used by Catalog
    type AttributeMasterSerializer struct {

        
            CreatedBy map[string]interface{}  `json:"created_by"`
            EnabledForEndConsumer bool  `json:"enabled_for_end_consumer"`
            Schema AttributeMaster  `json:"schema"`
            Description string  `json:"description"`
            Unit string  `json:"unit"`
            ModifiedOn string  `json:"modified_on"`
            IsNested bool  `json:"is_nested"`
            CreatedOn string  `json:"created_on"`
            Details AttributeMasterDetails  `json:"details"`
            Suggestion string  `json:"suggestion"`
            Name string  `json:"name"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Departments []string  `json:"departments"`
            Logo string  `json:"logo"`
            Variant bool  `json:"variant"`
            Filters AttributeMasterFilter  `json:"filters"`
            Tags []string  `json:"tags"`
            Slug string  `json:"slug"`
            RawKey string  `json:"raw_key"`
            Synonyms map[string]interface{}  `json:"synonyms"`
         
    }
    
    // ProductAttributesResponse used by Catalog
    type ProductAttributesResponse struct {

        
            Items []AttributeMasterSerializer  `json:"items"`
         
    }
    
    // Media3 used by Catalog
    type Media3 struct {

        
            URL string  `json:"url"`
            Meta map[string]interface{}  `json:"meta"`
            Type string  `json:"type"`
         
    }
    
    // TeaserTag1 used by Catalog
    type TeaserTag1 struct {

        
            URL string  `json:"url"`
            Tag string  `json:"tag"`
         
    }
    
    // ProductPublish1 used by Catalog
    type ProductPublish1 struct {

        
            IsSet bool  `json:"is_set"`
            ProductOnlineDate string  `json:"product_online_date"`
         
    }
    
    // Trader1 used by Catalog
    type Trader1 struct {

        
            Name string  `json:"name"`
            Address []string  `json:"address"`
            Type string  `json:"type"`
         
    }
    
    // TaxIdentifier1 used by Catalog
    type TaxIdentifier1 struct {

        
            ReportingHsn string  `json:"reporting_hsn"`
            HsnCodeID string  `json:"hsn_code_id"`
            HsnCode string  `json:"hsn_code"`
         
    }
    
    // CustomOrder1 used by Catalog
    type CustomOrder1 struct {

        
            ManufacturingTimeUnit string  `json:"manufacturing_time_unit"`
            IsCustomOrder bool  `json:"is_custom_order"`
            ManufacturingTime float64  `json:"manufacturing_time"`
         
    }
    
    // ReturnConfig1 used by Catalog
    type ReturnConfig1 struct {

        
            Time float64  `json:"time"`
            Returnable bool  `json:"returnable"`
            Unit string  `json:"unit"`
         
    }
    
    // ProductPatch used by Catalog
    type ProductPatch struct {

        
            Media []Media3  `json:"media"`
            ChangeRequestID interface{}  `json:"change_request_id"`
            TemplateTag string  `json:"template_tag"`
            MultiSize bool  `json:"multi_size"`
            ProductGroupTag []string  `json:"product_group_tag"`
            TeaserTag TeaserTag1  `json:"teaser_tag"`
            BulkJobID string  `json:"bulk_job_id"`
            Description string  `json:"description"`
            SizeGuide string  `json:"size_guide"`
            ProductPublish ProductPublish1  `json:"product_publish"`
            Action string  `json:"action"`
            VariantMedia map[string]interface{}  `json:"variant_media"`
            Requester string  `json:"requester"`
            IsActive bool  `json:"is_active"`
            CompanyID float64  `json:"company_id"`
            UID float64  `json:"uid"`
            IsSet bool  `json:"is_set"`
            NetQuantity NetQuantity  `json:"net_quantity"`
            CategorySlug string  `json:"category_slug"`
            CountryOfOrigin string  `json:"country_of_origin"`
            ItemCode interface{}  `json:"item_code"`
            Trader []Trader1  `json:"trader"`
            Name string  `json:"name"`
            TaxIdentifier TaxIdentifier1  `json:"tax_identifier"`
            BrandUID float64  `json:"brand_uid"`
            IsDependent bool  `json:"is_dependent"`
            IsImageLessProduct bool  `json:"is_image_less_product"`
            Departments []float64  `json:"departments"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ShortDescription string  `json:"short_description"`
            CustomOrder CustomOrder1  `json:"custom_order"`
            ReturnConfig ReturnConfig1  `json:"return_config"`
            Currency string  `json:"currency"`
            Tags []string  `json:"tags"`
            NoOfBoxes float64  `json:"no_of_boxes"`
            ItemType string  `json:"item_type"`
            Slug string  `json:"slug"`
            Variants map[string]interface{}  `json:"variants"`
            Highlights []string  `json:"highlights"`
         
    }
    
    // ValidateIdentifier used by Catalog
    type ValidateIdentifier struct {

        
            Primary bool  `json:"primary"`
            GtinType string  `json:"gtin_type"`
            GtinValue string  `json:"gtin_value"`
         
    }
    
    // ALLSizes used by Catalog
    type ALLSizes struct {

        
            Identifiers []ValidateIdentifier  `json:"identifiers"`
            ItemHeight float64  `json:"item_height"`
            ItemWidth float64  `json:"item_width"`
            ItemDimensionsUnitOfMeasure string  `json:"item_dimensions_unit_of_measure"`
            ItemLength float64  `json:"item_length"`
            ItemWeightUnitOfMeasure interface{}  `json:"item_weight_unit_of_measure"`
            Size interface{}  `json:"size"`
            ItemWeight float64  `json:"item_weight"`
         
    }
    
    // ListALLSizes used by Catalog
    type ListALLSizes struct {

        
            AllSizes []ALLSizes  `json:"all_sizes"`
         
    }
    
    // ValidateProduct used by Catalog
    type ValidateProduct struct {

        
            Valid bool  `json:"valid"`
         
    }
    
    // UserDetail1 used by Catalog
    type UserDetail1 struct {

        
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
            FullName string  `json:"full_name"`
         
    }
    
    // ProductBulkRequest used by Catalog
    type ProductBulkRequest struct {

        
            TemplateTag string  `json:"template_tag"`
            Stage string  `json:"stage"`
            FailedRecords []string  `json:"failed_records"`
            CreatedBy UserDetail1  `json:"created_by"`
            Cancelled float64  `json:"cancelled"`
            Failed float64  `json:"failed"`
            Succeed float64  `json:"succeed"`
            ModifiedOn string  `json:"modified_on"`
            Template ProductTemplate  `json:"template"`
            CancelledRecords []string  `json:"cancelled_records"`
            CompanyID float64  `json:"company_id"`
            ModifiedBy UserDetail1  `json:"modified_by"`
            IsActive bool  `json:"is_active"`
            CreatedOn string  `json:"created_on"`
            FilePath string  `json:"file_path"`
            Total float64  `json:"total"`
         
    }
    
    // ProductBulkRequestList used by Catalog
    type ProductBulkRequestList struct {

        
            Page Page  `json:"page"`
            Items ProductBulkRequest  `json:"items"`
         
    }
    
    // UserInfo1 used by Catalog
    type UserInfo1 struct {

        
            UID string  `json:"uid"`
            UserID string  `json:"user_id"`
            Email string  `json:"email"`
            Username string  `json:"username"`
         
    }
    
    // BulkJob used by Catalog
    type BulkJob struct {

        
            TrackingURL string  `json:"tracking_url"`
            CreatedBy UserInfo1  `json:"created_by"`
            Stage string  `json:"stage"`
            TemplateTag string  `json:"template_tag"`
            Department string  `json:"department"`
            Total float64  `json:"total"`
            ModifiedOn string  `json:"modified_on"`
            CompanyID float64  `json:"company_id"`
            CreatedOn string  `json:"created_on"`
            FilePath string  `json:"file_path"`
            Succeed float64  `json:"succeed"`
            Failed float64  `json:"failed"`
            ModifiedBy UserInfo1  `json:"modified_by"`
            Cancelled float64  `json:"cancelled"`
            FailedRecords []map[string]interface{}  `json:"failed_records"`
            ProductType string  `json:"product_type"`
            CustomTemplateTag string  `json:"custom_template_tag"`
            CancelledRecords []map[string]interface{}  `json:"cancelled_records"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // BulkResponse used by Catalog
    type BulkResponse struct {

        
            CreatedBy UserInfo1  `json:"created_by"`
            BatchID string  `json:"batch_id"`
            ModifiedOn string  `json:"modified_on"`
            ModifiedBy UserInfo1  `json:"modified_by"`
            IsActive bool  `json:"is_active"`
            CreatedOn string  `json:"created_on"`
         
    }
    
    // BulkProductRequest used by Catalog
    type BulkProductRequest struct {

        
            TemplateTag string  `json:"template_tag"`
            BatchID string  `json:"batch_id"`
            CompanyID float64  `json:"company_id"`
            Data []map[string]interface{}  `json:"data"`
         
    }
    
    // NestedTags used by Catalog
    type NestedTags struct {

        
            Tags []string  `json:"tags"`
         
    }
    
    // ProductTagsViewResponse used by Catalog
    type ProductTagsViewResponse struct {

        
            Items NestedTags  `json:"items"`
         
    }
    
    // UserCommon used by Catalog
    type UserCommon struct {

        
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
            CompanyID float64  `json:"company_id"`
         
    }
    
    // Items used by Catalog
    type Items struct {

        
            TrackingURL string  `json:"tracking_url"`
            CreatedBy UserCommon  `json:"created_by"`
            Stage string  `json:"stage"`
            FailedRecords []string  `json:"failed_records"`
            Cancelled float64  `json:"cancelled"`
            Succeed float64  `json:"succeed"`
            Failed float64  `json:"failed"`
            ModifiedOn string  `json:"modified_on"`
            ID string  `json:"id"`
            CancelledRecords []string  `json:"cancelled_records"`
            CompanyID float64  `json:"company_id"`
            Total float64  `json:"total"`
            ModifiedBy UserCommon  `json:"modified_by"`
            IsActive bool  `json:"is_active"`
            CreatedOn string  `json:"created_on"`
            FilePath string  `json:"file_path"`
            Retry float64  `json:"retry"`
         
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

        
            Identifiers map[string]interface{}  `json:"identifiers"`
            ItemID float64  `json:"item_id"`
            PriceTransfer float64  `json:"price_transfer"`
            Currency string  `json:"currency"`
            Size string  `json:"size"`
            InventoryUpdatedOn string  `json:"inventory_updated_on"`
            SellerIdentifier string  `json:"seller_identifier"`
            Quantity float64  `json:"quantity"`
            UID string  `json:"uid"`
            PriceEffective float64  `json:"price_effective"`
            SellableQuantity float64  `json:"sellable_quantity"`
            Store map[string]interface{}  `json:"store"`
            Price float64  `json:"price"`
         
    }
    
    // InventoryResponsePaginated used by Catalog
    type InventoryResponsePaginated struct {

        
            Page Page  `json:"page"`
            Items []InventoryResponse  `json:"items"`
         
    }
    
    // GTIN used by Catalog
    type GTIN struct {

        
            Primary bool  `json:"primary"`
            GtinType string  `json:"gtin_type"`
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

        
            Name string  `json:"name"`
            SizeDistribution SizeDistribution  `json:"size_distribution"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // InvSize used by Catalog
    type InvSize struct {

        
            ExpirationDate string  `json:"expiration_date"`
            Identifiers []GTIN  `json:"identifiers"`
            ItemWeight float64  `json:"item_weight"`
            PriceTransfer float64  `json:"price_transfer"`
            ItemHeight float64  `json:"item_height"`
            ItemWidth float64  `json:"item_width"`
            ItemDimensionsUnitOfMeasure string  `json:"item_dimensions_unit_of_measure"`
            ItemLength float64  `json:"item_length"`
            Currency string  `json:"currency"`
            Size interface{}  `json:"size"`
            ItemWeightUnitOfMeasure string  `json:"item_weight_unit_of_measure"`
            StoreCode string  `json:"store_code"`
            Quantity float64  `json:"quantity"`
            Set InventorySet  `json:"set"`
            PriceEffective float64  `json:"price_effective"`
            IsSet bool  `json:"is_set"`
            Price float64  `json:"price"`
         
    }
    
    // ItemQuery used by Catalog
    type ItemQuery struct {

        
            UID float64  `json:"uid"`
            ItemCode string  `json:"item_code"`
            BrandUID float64  `json:"brand_uid"`
         
    }
    
    // InventoryRequest used by Catalog
    type InventoryRequest struct {

        
            Sizes []InvSize  `json:"sizes"`
            CompanyID float64  `json:"company_id"`
            Item ItemQuery  `json:"item"`
         
    }
    
    // WeightResponse used by Catalog
    type WeightResponse struct {

        
            IsDefault bool  `json:"is_default"`
            Unit string  `json:"unit"`
            Shipping float64  `json:"shipping"`
         
    }
    
    // QuantityBase used by Catalog
    type QuantityBase struct {

        
            UpdatedAt string  `json:"updated_at"`
            Count float64  `json:"count"`
         
    }
    
    // Quantities used by Catalog
    type Quantities struct {

        
            Damaged QuantityBase  `json:"damaged"`
            NotAvailable QuantityBase  `json:"not_available"`
            OrderCommitted QuantityBase  `json:"order_committed"`
            Sellable QuantityBase  `json:"sellable"`
         
    }
    
    // DimensionResponse used by Catalog
    type DimensionResponse struct {

        
            IsDefault bool  `json:"is_default"`
            Length float64  `json:"length"`
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
    
    // BrandMeta used by Catalog
    type BrandMeta struct {

        
            Name string  `json:"name"`
            ID float64  `json:"id"`
         
    }
    
    // Trader2 used by Catalog
    type Trader2 struct {

        
            Name string  `json:"name"`
            Address []string  `json:"address"`
            Type string  `json:"type"`
         
    }
    
    // CompanyMeta used by Catalog
    type CompanyMeta struct {

        
            ID float64  `json:"id"`
         
    }
    
    // ReturnConfig2 used by Catalog
    type ReturnConfig2 struct {

        
            Time float64  `json:"time"`
            Returnable bool  `json:"returnable"`
            Unit string  `json:"unit"`
         
    }
    
    // PriceMeta used by Catalog
    type PriceMeta struct {

        
            Effective float64  `json:"effective"`
            UpdatedAt string  `json:"updated_at"`
            Transfer float64  `json:"transfer"`
            Currency string  `json:"currency"`
            Marked float64  `json:"marked"`
            TpNotes map[string]interface{}  `json:"tp_notes"`
         
    }
    
    // InventorySellerResponse used by Catalog
    type InventorySellerResponse struct {

        
            CreatedBy UserSerializer  `json:"created_by"`
            Stage string  `json:"stage"`
            Weight WeightResponse  `json:"weight"`
            Quantities Quantities  `json:"quantities"`
            TraceID string  `json:"trace_id"`
            SellerIdentifier string  `json:"seller_identifier"`
            AddedOnStore string  `json:"added_on_store"`
            RawMeta map[string]interface{}  `json:"raw_meta"`
            Dimension DimensionResponse  `json:"dimension"`
            Manufacturer ManufacturerResponse  `json:"manufacturer"`
            TrackInventory bool  `json:"track_inventory"`
            UID string  `json:"uid"`
            Fragile bool  `json:"fragile"`
            Size string  `json:"size"`
            IsSet bool  `json:"is_set"`
            FyndItemCode string  `json:"fynd_item_code"`
            ItemID float64  `json:"item_id"`
            CountryOfOrigin string  `json:"country_of_origin"`
            FyndArticleCode string  `json:"fynd_article_code"`
            Brand BrandMeta  `json:"brand"`
            FyndMeta map[string]interface{}  `json:"fynd_meta"`
            Trader []Trader2  `json:"trader"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            ExpirationDate string  `json:"expiration_date"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Company CompanyMeta  `json:"company"`
            TotalQuantity float64  `json:"total_quantity"`
            ReturnConfig ReturnConfig2  `json:"return_config"`
            Tags []string  `json:"tags"`
            Meta map[string]interface{}  `json:"meta"`
            Set InventorySet  `json:"set"`
            IsActive bool  `json:"is_active"`
            Identifier map[string]interface{}  `json:"identifier"`
            Store StoreMeta  `json:"store"`
            Price PriceMeta  `json:"price"`
         
    }
    
    // InventorySellerIdentifierResponsePaginated used by Catalog
    type InventorySellerIdentifierResponsePaginated struct {

        
            Page Page  `json:"page"`
            Items []InventorySellerResponse  `json:"items"`
         
    }
    
    // BulkInventoryGetItems used by Catalog
    type BulkInventoryGetItems struct {

        
            CreatedBy map[string]interface{}  `json:"created_by"`
            Cancelled float64  `json:"cancelled"`
            FailedRecords []string  `json:"failed_records"`
            Stage string  `json:"stage"`
            Failed float64  `json:"failed"`
            Succeed float64  `json:"succeed"`
            ModifiedOn string  `json:"modified_on"`
            CancelledRecords []string  `json:"cancelled_records"`
            CompanyID float64  `json:"company_id"`
            ID string  `json:"id"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            IsActive bool  `json:"is_active"`
            CreatedOn string  `json:"created_on"`
            FilePath string  `json:"file_path"`
            Total float64  `json:"total"`
         
    }
    
    // BulkInventoryGet used by Catalog
    type BulkInventoryGet struct {

        
            Page Page  `json:"page"`
            Items []BulkInventoryGetItems  `json:"items"`
         
    }
    
    // StoreIntegrationType used by Catalog
    type StoreIntegrationType struct {

        
            Order string  `json:"order"`
            Inventory string  `json:"inventory"`
         
    }
    
    // ArticleStoreResponse used by Catalog
    type ArticleStoreResponse struct {

        
            Name string  `json:"name"`
            StoreType string  `json:"store_type"`
            UID float64  `json:"uid"`
            IntegrationType StoreIntegrationType  `json:"integration_type"`
         
    }
    
    // GetInventories used by Catalog
    type GetInventories struct {

        
            CreatedBy UserSerializer  `json:"created_by"`
            Stage string  `json:"stage"`
            Weight WeightResponse  `json:"weight"`
            TraceID string  `json:"trace_id"`
            SellerIdentifier string  `json:"seller_identifier"`
            AddedOnStore string  `json:"added_on_store"`
            RawMeta map[string]interface{}  `json:"raw_meta"`
            Dimension DimensionResponse  `json:"dimension"`
            Manufacturer ManufacturerResponse  `json:"manufacturer"`
            TrackInventory bool  `json:"track_inventory"`
            InventoryUpdatedOn string  `json:"inventory_updated_on"`
            UID string  `json:"uid"`
            Fragile bool  `json:"fragile"`
            Size string  `json:"size"`
            IsSet bool  `json:"is_set"`
            FyndItemCode string  `json:"fynd_item_code"`
            ItemID float64  `json:"item_id"`
            CountryOfOrigin string  `json:"country_of_origin"`
            FyndArticleCode string  `json:"fynd_article_code"`
            Brand BrandMeta  `json:"brand"`
            FyndMeta map[string]interface{}  `json:"fynd_meta"`
            Trader []Trader2  `json:"trader"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            ExpirationDate string  `json:"expiration_date"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Company CompanyMeta  `json:"company"`
            TotalQuantity float64  `json:"total_quantity"`
            ReturnConfig ReturnConfig2  `json:"return_config"`
            Tags []string  `json:"tags"`
            Meta map[string]interface{}  `json:"meta"`
            Set InventorySet  `json:"set"`
            IsActive bool  `json:"is_active"`
            Identifier map[string]interface{}  `json:"identifier"`
            Store ArticleStoreResponse  `json:"store"`
            Price PriceMeta  `json:"price"`
         
    }
    
    // GetInventoriesResponse used by Catalog
    type GetInventoriesResponse struct {

        
            Page Page  `json:"page"`
            Items []GetInventories  `json:"items"`
         
    }
    
    // InventoryJobPayload used by Catalog
    type InventoryJobPayload struct {

        
            ExpirationDate string  `json:"expiration_date"`
            ItemDimensionsUnitOfMeasure string  `json:"item_dimensions_unit_of_measure"`
            TotalQuantity float64  `json:"total_quantity"`
            TraceID string  `json:"trace_id"`
            Currency string  `json:"currency"`
            Tags []string  `json:"tags"`
            ItemWeightUnitOfMeasure string  `json:"item_weight_unit_of_measure"`
            PriceMarked float64  `json:"price_marked"`
            SellerIdentifier string  `json:"seller_identifier"`
            StoreCode string  `json:"store_code"`
            Quantity float64  `json:"quantity"`
            PriceEffective float64  `json:"price_effective"`
            Price float64  `json:"price"`
         
    }
    
    // InventoryBulkRequest used by Catalog
    type InventoryBulkRequest struct {

        
            Sizes []InventoryJobPayload  `json:"sizes"`
            User map[string]interface{}  `json:"user"`
            CompanyID float64  `json:"company_id"`
            BatchID string  `json:"batch_id"`
         
    }
    
    // InventoryExportJob used by Catalog
    type InventoryExportJob struct {

        
            TaskID string  `json:"task_id"`
            URL string  `json:"url"`
            CompletedOn string  `json:"completed_on"`
            TriggerOn string  `json:"trigger_on"`
            RequestParams map[string]interface{}  `json:"request_params"`
            SellerID float64  `json:"seller_id"`
            Status string  `json:"status"`
         
    }
    
    // InventoryExportRequest used by Catalog
    type InventoryExportRequest struct {

        
            Store []float64  `json:"store"`
            Brand []float64  `json:"brand"`
            Type string  `json:"type"`
         
    }
    
    // InventoryExportResponse used by Catalog
    type InventoryExportResponse struct {

        
            TaskID string  `json:"task_id"`
            TriggerOn string  `json:"trigger_on"`
            RequestParams map[string]interface{}  `json:"request_params"`
            SellerID float64  `json:"seller_id"`
            Status string  `json:"status"`
         
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
            StoreID float64  `json:"store_id"`
            TotalQuantity float64  `json:"total_quantity"`
            TraceID string  `json:"trace_id"`
            Tags []string  `json:"tags"`
            PriceMarked float64  `json:"price_marked"`
            SellerIdentifier string  `json:"seller_identifier"`
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

        
            Data InventoryPayload  `json:"data"`
            Reason InventoryFailedReason  `json:"reason"`
         
    }
    
    // InventoryUpdateResponse used by Catalog
    type InventoryUpdateResponse struct {

        
            Message string  `json:"message"`
            Items []InventoryResponseItem  `json:"items"`
         
    }
    
    // PageResponse used by Catalog
    type PageResponse struct {

        
            HasPrevious bool  `json:"has_previous"`
            ItemTotal float64  `json:"item_total"`
            Current string  `json:"current"`
            HasNext bool  `json:"has_next"`
            Size float64  `json:"size"`
         
    }
    
    // HsnCodesObject used by Catalog
    type HsnCodesObject struct {

        
            Tax1 float64  `json:"tax1"`
            Tax2 float64  `json:"tax2"`
            HsnCode string  `json:"hsn_code"`
            Threshold1 float64  `json:"threshold1"`
            TaxOnMrp bool  `json:"tax_on_mrp"`
            ModifiedOn string  `json:"modified_on"`
            TaxOnEsp bool  `json:"tax_on_esp"`
            CompanyID float64  `json:"company_id"`
            ID string  `json:"id"`
            Hs2Code string  `json:"hs2_code"`
            Threshold2 float64  `json:"threshold2"`
         
    }
    
    // HsnCodesListingResponse used by Catalog
    type HsnCodesListingResponse struct {

        
            Page PageResponse  `json:"page"`
            Items []HsnCodesObject  `json:"items"`
         
    }
    
    // HsnUpsert used by Catalog
    type HsnUpsert struct {

        
            Tax1 float64  `json:"tax1"`
            Tax2 float64  `json:"tax2"`
            HsnCode string  `json:"hsn_code"`
            Threshold1 float64  `json:"threshold1"`
            TaxOnMrp bool  `json:"tax_on_mrp"`
            TaxOnEsp bool  `json:"tax_on_esp"`
            CompanyID float64  `json:"company_id"`
            UID float64  `json:"uid"`
            IsActive bool  `json:"is_active"`
            Hs2Code string  `json:"hs2_code"`
            Threshold2 float64  `json:"threshold2"`
         
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

        
            Description string  `json:"description"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Type string  `json:"type"`
            CountryCode string  `json:"country_code"`
            HsnCode string  `json:"hsn_code"`
            ModifiedOn string  `json:"modified_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            CreatedOn string  `json:"created_on"`
            ReportingHsn string  `json:"reporting_hsn"`
            Taxes []TaxSlab  `json:"taxes"`
         
    }
    
    // HsnCodesListingResponseSchemaV2 used by Catalog
    type HsnCodesListingResponseSchemaV2 struct {

        
            Page PageResponse  `json:"page"`
            Items []HSNDataInsertV2  `json:"items"`
         
    }
    
    // ArticleAssignment used by Catalog
    type ArticleAssignment struct {

        
            Level string  `json:"level"`
            Strategy string  `json:"strategy"`
         
    }
    
    // ArticleQuery used by Catalog
    type ArticleQuery struct {

        
            Size string  `json:"size"`
            IgnoredStores []float64  `json:"ignored_stores"`
            ItemID float64  `json:"item_id"`
         
    }
    
    // AssignStoreArticle used by Catalog
    type AssignStoreArticle struct {

        
            GroupID string  `json:"group_id"`
            Meta map[string]interface{}  `json:"meta"`
            Quantity float64  `json:"quantity"`
            ArticleAssignment ArticleAssignment  `json:"article_assignment"`
            Query ArticleQuery  `json:"query"`
         
    }
    
    // AssignStore used by Catalog
    type AssignStore struct {

        
            ChannelType string  `json:"channel_type"`
            ChannelIdentifier string  `json:"channel_identifier"`
            Articles []AssignStoreArticle  `json:"articles"`
            CompanyID float64  `json:"company_id"`
            AppID string  `json:"app_id"`
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

        
            Index float64  `json:"index"`
            ID string  `json:"_id"`
            GroupID string  `json:"group_id"`
            ItemID float64  `json:"item_id"`
            StrategyWiseListing []map[string]interface{}  `json:"strategy_wise_listing"`
            StoreID float64  `json:"store_id"`
            StorePincode float64  `json:"store_pincode"`
            SCity string  `json:"s_city"`
            Size string  `json:"size"`
            Meta map[string]interface{}  `json:"meta"`
            CompanyID float64  `json:"company_id"`
            ArticleAssignment ArticleAssignment1  `json:"article_assignment"`
            UID string  `json:"uid"`
            PriceEffective float64  `json:"price_effective"`
            Status bool  `json:"status"`
            PriceMarked float64  `json:"price_marked"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // BrandItem used by Catalog
    type BrandItem struct {

        
            Banners ImageUrls  `json:"banners"`
            Action Action  `json:"action"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
            Departments []string  `json:"departments"`
            Logo Media  `json:"logo"`
            Discount string  `json:"discount"`
         
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
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
            Logo Media  `json:"logo"`
         
    }
    
    // DepartmentResponse used by Catalog
    type DepartmentResponse struct {

        
            Items []Department  `json:"items"`
         
    }
    
    // ThirdLevelChild used by Catalog
    type ThirdLevelChild struct {

        
            Banners ImageUrls  `json:"banners"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Action Action  `json:"action"`
            Name string  `json:"name"`
            Childs []map[string]interface{}  `json:"childs"`
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
         
    }
    
    // SecondLevelChild used by Catalog
    type SecondLevelChild struct {

        
            Banners ImageUrls  `json:"banners"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Action Action  `json:"action"`
            Name string  `json:"name"`
            Childs []ThirdLevelChild  `json:"childs"`
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
         
    }
    
    // Child used by Catalog
    type Child struct {

        
            Banners ImageUrls  `json:"banners"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Action Action  `json:"action"`
            Name string  `json:"name"`
            Childs []SecondLevelChild  `json:"childs"`
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
         
    }
    
    // CategoryItems used by Catalog
    type CategoryItems struct {

        
            Banners ImageUrls  `json:"banners"`
            Action Action  `json:"action"`
            Name string  `json:"name"`
            Childs []Child  `json:"childs"`
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
         
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
            Filters []ProductFilters  `json:"filters"`
            Items []ProductListingDetail  `json:"items"`
            Operators map[string]interface{}  `json:"operators"`
         
    }
    
    // ProductDetail used by Catalog
    type ProductDetail struct {

        
            Color string  `json:"color"`
            Type string  `json:"type"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            Attributes map[string]interface{}  `json:"attributes"`
            Description string  `json:"description"`
            RatingCount float64  `json:"rating_count"`
            ImageNature string  `json:"image_nature"`
            UID float64  `json:"uid"`
            Rating float64  `json:"rating"`
            PromoMeta map[string]interface{}  `json:"promo_meta"`
            ItemCode string  `json:"item_code"`
            Brand ProductBrand  `json:"brand"`
            Name string  `json:"name"`
            Similars []string  `json:"similars"`
            HasVariant bool  `json:"has_variant"`
            Medias []Media1  `json:"medias"`
            ShortDescription string  `json:"short_description"`
            ItemType string  `json:"item_type"`
            Slug string  `json:"slug"`
            Tryouts []string  `json:"tryouts"`
            ProductOnlineDate string  `json:"product_online_date"`
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
            Highlights []string  `json:"highlights"`
         
    }
    
    // InventoryPage used by Catalog
    type InventoryPage struct {

        
            ItemTotal float64  `json:"item_total"`
            Type string  `json:"type"`
            HasNext bool  `json:"has_next"`
            HasPrevious bool  `json:"has_previous"`
            NextID string  `json:"next_id"`
         
    }
    
    // InventoryStockResponse used by Catalog
    type InventoryStockResponse struct {

        
            Page InventoryPage  `json:"page"`
            Items []map[string]interface{}  `json:"items"`
         
    }
    
    // UserSerializer1 used by Catalog
    type UserSerializer1 struct {

        
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
            Contact string  `json:"contact"`
         
    }
    
    // GetAddressSerializer used by Catalog
    type GetAddressSerializer struct {

        
            AddressType string  `json:"address_type"`
            Country string  `json:"country"`
            Longitude float64  `json:"longitude"`
            CountryCode string  `json:"country_code"`
            Landmark string  `json:"landmark"`
            Latitude float64  `json:"latitude"`
            City string  `json:"city"`
            State string  `json:"state"`
            Address2 string  `json:"address2"`
            Address1 string  `json:"address1"`
            Pincode float64  `json:"pincode"`
         
    }
    
    // LocationIntegrationType used by Catalog
    type LocationIntegrationType struct {

        
            Order string  `json:"order"`
            Inventory string  `json:"inventory"`
         
    }
    
    // InvoiceCredSerializer used by Catalog
    type InvoiceCredSerializer struct {

        
            Username string  `json:"username"`
            Password string  `json:"password"`
            Enabled bool  `json:"enabled"`
         
    }
    
    // InvoiceDetailsSerializer used by Catalog
    type InvoiceDetailsSerializer struct {

        
            EInvoice InvoiceCredSerializer  `json:"e_invoice"`
            EWaybill InvoiceCredSerializer  `json:"e_waybill"`
         
    }
    
    // SellerPhoneNumber used by Catalog
    type SellerPhoneNumber struct {

        
            Number string  `json:"number"`
            CountryCode float64  `json:"country_code"`
         
    }
    
    // ProductReturnConfigSerializer used by Catalog
    type ProductReturnConfigSerializer struct {

        
            OnSameStore bool  `json:"on_same_store"`
            StoreUID float64  `json:"store_uid"`
         
    }
    
    // LocationTimingSerializer used by Catalog
    type LocationTimingSerializer struct {

        
            Minute float64  `json:"minute"`
            Hour float64  `json:"hour"`
         
    }
    
    // LocationDayWiseSerializer used by Catalog
    type LocationDayWiseSerializer struct {

        
            Closing LocationTimingSerializer  `json:"closing"`
            Opening LocationTimingSerializer  `json:"opening"`
            Weekday string  `json:"weekday"`
            Open bool  `json:"open"`
         
    }
    
    // LocationManagerSerializer used by Catalog
    type LocationManagerSerializer struct {

        
            Name string  `json:"name"`
            MobileNo SellerPhoneNumber  `json:"mobile_no"`
            Email string  `json:"email"`
         
    }
    
    // UserSerializer2 used by Catalog
    type UserSerializer2 struct {

        
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
            Contact string  `json:"contact"`
         
    }
    
    // GetCompanySerializer used by Catalog
    type GetCompanySerializer struct {

        
            BusinessType string  `json:"business_type"`
            CreatedBy UserSerializer2  `json:"created_by"`
            Stage string  `json:"stage"`
            RejectReason string  `json:"reject_reason"`
            CompanyType string  `json:"company_type"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            ModifiedOn string  `json:"modified_on"`
            VerifiedBy UserSerializer2  `json:"verified_by"`
            Name string  `json:"name"`
            VerifiedOn string  `json:"verified_on"`
            UID float64  `json:"uid"`
            CreatedOn string  `json:"created_on"`
            ModifiedBy UserSerializer2  `json:"modified_by"`
         
    }
    
    // GetLocationSerializer used by Catalog
    type GetLocationSerializer struct {

        
            CreatedBy UserSerializer1  `json:"created_by"`
            Stage string  `json:"stage"`
            Address GetAddressSerializer  `json:"address"`
            DisplayName string  `json:"display_name"`
            PhoneNumber string  `json:"phone_number"`
            IntegrationType LocationIntegrationType  `json:"integration_type"`
            GstCredentials InvoiceDetailsSerializer  `json:"gst_credentials"`
            NotificationEmails []string  `json:"notification_emails"`
            ModifiedOn string  `json:"modified_on"`
            Warnings map[string]interface{}  `json:"warnings"`
            UID float64  `json:"uid"`
            CreatedOn string  `json:"created_on"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            Documents []Document  `json:"documents"`
            ProductReturnConfig ProductReturnConfigSerializer  `json:"product_return_config"`
            Name string  `json:"name"`
            VerifiedOn string  `json:"verified_on"`
            ModifiedBy UserSerializer1  `json:"modified_by"`
            Timing []LocationDayWiseSerializer  `json:"timing"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Manager LocationManagerSerializer  `json:"manager"`
            Company GetCompanySerializer  `json:"company"`
            StoreType string  `json:"store_type"`
            VerifiedBy UserSerializer1  `json:"verified_by"`
            Code string  `json:"code"`
         
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
            Name string  `json:"name"`
            AppID string  `json:"app_id"`
            UID float64  `json:"uid"`
            IsActive bool  `json:"is_active"`
            Logo string  `json:"logo"`
         
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
    

    
    // CompanyTaxesSerializer used by CompanyProfile
    type CompanyTaxesSerializer struct {

        
            Enable bool  `json:"enable"`
            Rate float64  `json:"rate"`
            EffectiveDate string  `json:"effective_date"`
         
    }
    
    // ContactDetails used by CompanyProfile
    type ContactDetails struct {

        
            Phone []SellerPhoneNumber  `json:"phone"`
            Emails []string  `json:"emails"`
         
    }
    
    // BusinessCountryInfo used by CompanyProfile
    type BusinessCountryInfo struct {

        
            Country string  `json:"country"`
            CountryCode string  `json:"country_code"`
         
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
            Taxes []CompanyTaxesSerializer  `json:"taxes"`
            CompanyType string  `json:"company_type"`
            Stage string  `json:"stage"`
            UID float64  `json:"uid"`
            Warnings map[string]interface{}  `json:"warnings"`
            Documents []Document  `json:"documents"`
            BusinessType string  `json:"business_type"`
            CreatedOn string  `json:"created_on"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ContactDetails ContactDetails  `json:"contact_details"`
            VerifiedOn string  `json:"verified_on"`
            CreatedBy UserSerializer  `json:"created_by"`
            NotificationEmails []string  `json:"notification_emails"`
            BusinessCountryInfo BusinessCountryInfo  `json:"business_country_info"`
            FranchiseEnabled bool  `json:"franchise_enabled"`
            ModifiedOn string  `json:"modified_on"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            Mode string  `json:"mode"`
            BusinessInfo string  `json:"business_info"`
            Name string  `json:"name"`
            BusinessDetails BusinessDetails  `json:"business_details"`
         
    }
    
    // CompanyTaxesSerializer1 used by CompanyProfile
    type CompanyTaxesSerializer1 struct {

        
            Enable bool  `json:"enable"`
            Rate float64  `json:"rate"`
            EffectiveDate string  `json:"effective_date"`
         
    }
    
    // CreateUpdateAddressSerializer used by CompanyProfile
    type CreateUpdateAddressSerializer struct {

        
            Latitude float64  `json:"latitude"`
            Longitude float64  `json:"longitude"`
            Pincode float64  `json:"pincode"`
            AddressType string  `json:"address_type"`
            Country string  `json:"country"`
            CountryCode string  `json:"country_code"`
            State string  `json:"state"`
            Landmark string  `json:"landmark"`
            Address2 string  `json:"address2"`
            City string  `json:"city"`
            Address1 string  `json:"address1"`
         
    }
    
    // UpdateCompany used by CompanyProfile
    type UpdateCompany struct {

        
            RejectReason string  `json:"reject_reason"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ContactDetails ContactDetails  `json:"contact_details"`
            CompanyType string  `json:"company_type"`
            BusinessInfo string  `json:"business_info"`
            Taxes []CompanyTaxesSerializer1  `json:"taxes"`
            NotificationEmails []string  `json:"notification_emails"`
            FranchiseEnabled bool  `json:"franchise_enabled"`
            Warnings map[string]interface{}  `json:"warnings"`
            Name string  `json:"name"`
            Documents []Document  `json:"documents"`
            BusinessType string  `json:"business_type"`
            Addresses []CreateUpdateAddressSerializer  `json:"addresses"`
            BusinessDetails BusinessDetails  `json:"business_details"`
         
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

        
            Stage string  `json:"stage"`
            UID float64  `json:"uid"`
            Brand DocumentsObj  `json:"brand"`
            CompanyDocuments DocumentsObj  `json:"company_documents"`
            Store DocumentsObj  `json:"store"`
            StoreDocuments DocumentsObj  `json:"store_documents"`
            Product DocumentsObj  `json:"product"`
         
    }
    
    // BrandBannerSerializer used by CompanyProfile
    type BrandBannerSerializer struct {

        
            Landscape string  `json:"landscape"`
            Portrait string  `json:"portrait"`
         
    }
    
    // GetBrandResponseSerializer used by CompanyProfile
    type GetBrandResponseSerializer struct {

        
            VerifiedBy UserSerializer  `json:"verified_by"`
            Description string  `json:"description"`
            Banner BrandBannerSerializer  `json:"banner"`
            SlugKey string  `json:"slug_key"`
            Stage string  `json:"stage"`
            UID float64  `json:"uid"`
            Warnings map[string]interface{}  `json:"warnings"`
            CreatedOn string  `json:"created_on"`
            RejectReason string  `json:"reject_reason"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            VerifiedOn string  `json:"verified_on"`
            CreatedBy UserSerializer  `json:"created_by"`
            ModifiedOn string  `json:"modified_on"`
            Logo string  `json:"logo"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            Mode string  `json:"mode"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            Synonyms []string  `json:"synonyms"`
            Name string  `json:"name"`
         
    }
    
    // CreateUpdateBrandRequestSerializer used by CompanyProfile
    type CreateUpdateBrandRequestSerializer struct {

        
            CompanyID float64  `json:"company_id"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            BrandTier string  `json:"brand_tier"`
            UID float64  `json:"uid"`
            Description string  `json:"description"`
            Banner BrandBannerSerializer  `json:"banner"`
            Synonyms []string  `json:"synonyms"`
            Name string  `json:"name"`
            Logo string  `json:"logo"`
         
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

        
            VerifiedBy UserSerializer  `json:"verified_by"`
            RejectReason string  `json:"reject_reason"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            VerifiedOn string  `json:"verified_on"`
            CompanyType string  `json:"company_type"`
            Stage string  `json:"stage"`
            CreatedBy UserSerializer  `json:"created_by"`
            UID float64  `json:"uid"`
            MarketChannels []string  `json:"market_channels"`
            NotificationEmails []string  `json:"notification_emails"`
            Name string  `json:"name"`
            BusinessCountryInfo BusinessCountryInfo  `json:"business_country_info"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            ModifiedOn string  `json:"modified_on"`
            Details CompanyDetails  `json:"details"`
            BusinessType string  `json:"business_type"`
            CreatedOn string  `json:"created_on"`
            ModifiedBy UserSerializer  `json:"modified_by"`
         
    }
    
    // CompanyBrandSerializer used by CompanyProfile
    type CompanyBrandSerializer struct {

        
            VerifiedBy UserSerializer  `json:"verified_by"`
            RejectReason string  `json:"reject_reason"`
            VerifiedOn string  `json:"verified_on"`
            Stage string  `json:"stage"`
            CreatedBy UserSerializer  `json:"created_by"`
            UID float64  `json:"uid"`
            Company CompanySerializer  `json:"company"`
            Brand GetBrandResponseSerializer  `json:"brand"`
            Warnings map[string]interface{}  `json:"warnings"`
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
            ModifiedBy UserSerializer  `json:"modified_by"`
         
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
    
    // LocationSerializer used by CompanyProfile
    type LocationSerializer struct {

        
            StoreType string  `json:"store_type"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Code string  `json:"code"`
            Timing []LocationDayWiseSerializer  `json:"timing"`
            Stage string  `json:"stage"`
            GstCredentials InvoiceDetailsSerializer  `json:"gst_credentials"`
            Company float64  `json:"company"`
            UID float64  `json:"uid"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            NotificationEmails []string  `json:"notification_emails"`
            DisplayName string  `json:"display_name"`
            Holiday []HolidaySchemaSerializer  `json:"holiday"`
            Manager LocationManagerSerializer  `json:"manager"`
            Warnings map[string]interface{}  `json:"warnings"`
            ProductReturnConfig ProductReturnConfigSerializer  `json:"product_return_config"`
            Address GetAddressSerializer  `json:"address"`
            Name string  `json:"name"`
            Documents []Document  `json:"documents"`
         
    }
    
    // BulkLocationSerializer used by CompanyProfile
    type BulkLocationSerializer struct {

        
            Data []LocationSerializer  `json:"data"`
         
    }
    
    // _ArticleAssignment used by CompanyProfile
    type _ArticleAssignment struct {

        
            Strategy string  `json:"strategy"`
            Level string  `json:"level"`
         
    }
    
    // _ArticleQuery used by CompanyProfile
    type _ArticleQuery struct {

        
            ItemID float64  `json:"item_id"`
            Size string  `json:"size"`
            IgnoredStores []float64  `json:"ignored_stores"`
         
    }
    
    // _AssignStoreArticle used by CompanyProfile
    type _AssignStoreArticle struct {

        
            Meta map[string]interface{}  `json:"meta"`
            GroupID string  `json:"group_id"`
            ArticleAssignment _ArticleAssignment  `json:"article_assignment"`
            Query _ArticleQuery  `json:"query"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // AssignStoreRequestValidator used by CompanyProfile
    type AssignStoreRequestValidator struct {

        
            CompanyID float64  `json:"company_id"`
            Pincode string  `json:"pincode"`
            ChannelType string  `json:"channel_type"`
            ChannelIdentifier string  `json:"channel_identifier"`
            AppID string  `json:"app_id"`
            StoreIds []float64  `json:"store_ids"`
            Articles []_AssignStoreArticle  `json:"articles"`
         
    }
    
    // AssignStoreResponseSerializer used by CompanyProfile
    type AssignStoreResponseSerializer struct {

        
            CompanyID float64  `json:"company_id"`
            SCity string  `json:"s_city"`
            PriceMarked float64  `json:"price_marked"`
            PreiceEffective float64  `json:"preice_effective"`
            Meta map[string]interface{}  `json:"meta"`
            UID string  `json:"uid"`
            ArticleAssignment _ArticleAssignment  `json:"article_assignment"`
            Status bool  `json:"status"`
            Size string  `json:"size"`
            Index float64  `json:"index"`
            Quantity float64  `json:"quantity"`
            StorePincode string  `json:"store_pincode"`
            ItemID float64  `json:"item_id"`
            StoreID float64  `json:"store_id"`
            ID string  `json:"_id"`
         
    }
    

    
    // FailedResponse used by FileStorage
    type FailedResponse struct {

        
            Message string  `json:"message"`
         
    }
    
    // CDN used by FileStorage
    type CDN struct {

        
            URL string  `json:"url"`
         
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
            EmptyCart bool  `json:"empty_cart"`
         
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
    
    // Rule used by Cart
    type Rule struct {

        
            Key float64  `json:"key"`
            Min float64  `json:"min"`
            DiscountQty float64  `json:"discount_qty"`
            Max float64  `json:"max"`
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

        
            ModifiedBy string  `json:"modified_by"`
            CreatedBy string  `json:"created_by"`
         
    }
    
    // CouponDateMeta used by Cart
    type CouponDateMeta struct {

        
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // CouponAction used by Cart
    type CouponAction struct {

        
            ActionDate string  `json:"action_date"`
            TxnMode string  `json:"txn_mode"`
         
    }
    
    // CouponSchedule used by Cart
    type CouponSchedule struct {

        
            End string  `json:"end"`
            NextSchedule []map[string]interface{}  `json:"next_schedule"`
            Duration float64  `json:"duration"`
            Cron string  `json:"cron"`
            Start string  `json:"start"`
         
    }
    
    // RuleDefinition used by Cart
    type RuleDefinition struct {

        
            ApplicableOn string  `json:"applicable_on"`
            Scope []string  `json:"scope"`
            CalculateOn string  `json:"calculate_on"`
            Type string  `json:"type"`
            CurrencyCode string  `json:"currency_code"`
            IsExact bool  `json:"is_exact"`
            AutoApply bool  `json:"auto_apply"`
            ValueType string  `json:"value_type"`
         
    }
    
    // DisplayMetaDict used by Cart
    type DisplayMetaDict struct {

        
            Title string  `json:"title"`
            Subtitle string  `json:"subtitle"`
         
    }
    
    // DisplayMeta used by Cart
    type DisplayMeta struct {

        
            Subtitle string  `json:"subtitle"`
            Auto DisplayMetaDict  `json:"auto"`
            Apply DisplayMetaDict  `json:"apply"`
            Title string  `json:"title"`
            Remove DisplayMetaDict  `json:"remove"`
            Description string  `json:"description"`
         
    }
    
    // State used by Cart
    type State struct {

        
            IsPublic bool  `json:"is_public"`
            IsArchived bool  `json:"is_archived"`
            IsDisplay bool  `json:"is_display"`
         
    }
    
    // UsesRemaining used by Cart
    type UsesRemaining struct {

        
            Total float64  `json:"total"`
            User float64  `json:"user"`
            App float64  `json:"app"`
         
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
    
    // PostOrder used by Cart
    type PostOrder struct {

        
            ReturnAllowed bool  `json:"return_allowed"`
            CancellationAllowed bool  `json:"cancellation_allowed"`
         
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
            Uses PaymentAllowValue  `json:"uses"`
            Networks []string  `json:"networks"`
            Types []string  `json:"types"`
         
    }
    
    // Restrictions used by Cart
    type Restrictions struct {

        
            Uses UsesRestriction  `json:"uses"`
            BulkBundle BulkBundleRestriction  `json:"bulk_bundle"`
            OrderingStores []float64  `json:"ordering_stores"`
            PostOrder PostOrder  `json:"post_order"`
            UserGroups []float64  `json:"user_groups"`
            PriceRange PriceRange  `json:"price_range"`
            Payments map[string]PaymentModes  `json:"payments"`
            Platforms []string  `json:"platforms"`
            CouponAllowed bool  `json:"coupon_allowed"`
            UserType string  `json:"user_type"`
         
    }
    
    // Validity used by Cart
    type Validity struct {

        
            Priority float64  `json:"priority"`
         
    }
    
    // CouponAdd used by Cart
    type CouponAdd struct {

        
            Ownership Ownership  `json:"ownership"`
            Rule []Rule  `json:"rule"`
            Validation Validation  `json:"validation"`
            Author CouponAuthor  `json:"author"`
            DateMeta CouponDateMeta  `json:"date_meta"`
            Action CouponAction  `json:"action"`
            Schedule CouponSchedule  `json:"_schedule"`
            RuleDefinition RuleDefinition  `json:"rule_definition"`
            DisplayMeta DisplayMeta  `json:"display_meta"`
            Identifiers Identifier  `json:"identifiers"`
            TypeSlug string  `json:"type_slug"`
            State State  `json:"state"`
            Tags []string  `json:"tags"`
            Code string  `json:"code"`
            Restrictions Restrictions  `json:"restrictions"`
            Validity Validity  `json:"validity"`
         
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
            Rule []Rule  `json:"rule"`
            Validation Validation  `json:"validation"`
            Author CouponAuthor  `json:"author"`
            DateMeta CouponDateMeta  `json:"date_meta"`
            Action CouponAction  `json:"action"`
            Schedule CouponSchedule  `json:"_schedule"`
            RuleDefinition RuleDefinition  `json:"rule_definition"`
            DisplayMeta DisplayMeta  `json:"display_meta"`
            Identifiers Identifier  `json:"identifiers"`
            TypeSlug string  `json:"type_slug"`
            State State  `json:"state"`
            Tags []string  `json:"tags"`
            Code string  `json:"code"`
            Restrictions Restrictions  `json:"restrictions"`
            Validity Validity  `json:"validity"`
         
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
    
    // PromotionAction used by Cart
    type PromotionAction struct {

        
            ActionDate string  `json:"action_date"`
            ActionType string  `json:"action_type"`
         
    }
    
    // CompareObject used by Cart
    type CompareObject struct {

        
            GreaterThan float64  `json:"greater_than"`
            Equals float64  `json:"equals"`
            LessThan float64  `json:"less_than"`
            LessThanEquals float64  `json:"less_than_equals"`
            GreaterThanEquals float64  `json:"greater_than_equals"`
         
    }
    
    // ItemCriteria used by Cart
    type ItemCriteria struct {

        
            CartUniqueItemQuantity CompareObject  `json:"cart_unique_item_quantity"`
            ItemID []float64  `json:"item_id"`
            ItemSize []string  `json:"item_size"`
            ItemBrand []float64  `json:"item_brand"`
            ItemExcludeID []float64  `json:"item_exclude_id"`
            AvailableZones []string  `json:"available_zones"`
            ItemExcludeCategory []float64  `json:"item_exclude_category"`
            ItemCategory []float64  `json:"item_category"`
            ProductTags []string  `json:"product_tags"`
            CartQuantity CompareObject  `json:"cart_quantity"`
            BuyRules []string  `json:"buy_rules"`
            ItemExcludeStore []float64  `json:"item_exclude_store"`
            CartTotal CompareObject  `json:"cart_total"`
            ItemStore []float64  `json:"item_store"`
            ItemExcludeBrand []float64  `json:"item_exclude_brand"`
            CartUniqueItemAmount CompareObject  `json:"cart_unique_item_amount"`
            ItemExcludeSku []string  `json:"item_exclude_sku"`
            AllItems bool  `json:"all_items"`
            ItemExcludeCompany []float64  `json:"item_exclude_company"`
            ItemSku []string  `json:"item_sku"`
            ItemCompany []float64  `json:"item_company"`
         
    }
    
    // Visibility used by Cart
    type Visibility struct {

        
            CouponList bool  `json:"coupon_list"`
            Pdp bool  `json:"pdp"`
         
    }
    
    // DisplayMeta1 used by Cart
    type DisplayMeta1 struct {

        
            OfferLabel string  `json:"offer_label"`
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

        
            Remaining UsesRemaining1  `json:"remaining"`
            Maximum UsesRemaining1  `json:"maximum"`
         
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

        
            Codes []string  `json:"codes"`
            Uses PaymentAllowValue1  `json:"uses"`
            Type string  `json:"type"`
         
    }
    
    // UserRegistered used by Cart
    type UserRegistered struct {

        
            End string  `json:"end"`
            Start string  `json:"start"`
         
    }
    
    // Restrictions1 used by Cart
    type Restrictions1 struct {

        
            Uses UsesRestriction1  `json:"uses"`
            UserID []string  `json:"user_id"`
            PostOrder PostOrder1  `json:"post_order"`
            UserGroups []float64  `json:"user_groups"`
            OrderQuantity float64  `json:"order_quantity"`
            Payments []PromotionPaymentModes  `json:"payments"`
            Platforms []string  `json:"platforms"`
            UserRegistered UserRegistered  `json:"user_registered"`
            AnonymousUsers bool  `json:"anonymous_users"`
         
    }
    
    // Ownership1 used by Cart
    type Ownership1 struct {

        
            PayableBy string  `json:"payable_by"`
            PayableCategory string  `json:"payable_category"`
         
    }
    
    // DiscountOffer used by Cart
    type DiscountOffer struct {

        
            ApportionDiscount bool  `json:"apportion_discount"`
            DiscountPrice float64  `json:"discount_price"`
            DiscountPercentage float64  `json:"discount_percentage"`
            MaxDiscountAmount float64  `json:"max_discount_amount"`
            MinOfferQuantity float64  `json:"min_offer_quantity"`
            Code string  `json:"code"`
            MaxOfferQuantity float64  `json:"max_offer_quantity"`
            DiscountAmount float64  `json:"discount_amount"`
            MaxUsagePerTransaction float64  `json:"max_usage_per_transaction"`
            PartialCanRet bool  `json:"partial_can_ret"`
         
    }
    
    // DiscountRule used by Cart
    type DiscountRule struct {

        
            BuyCondition string  `json:"buy_condition"`
            DiscountType string  `json:"discount_type"`
            ItemCriteria ItemCriteria  `json:"item_criteria"`
            Offer DiscountOffer  `json:"offer"`
         
    }
    
    // PromotionSchedule used by Cart
    type PromotionSchedule struct {

        
            End string  `json:"end"`
            NextSchedule []map[string]interface{}  `json:"next_schedule"`
            Duration float64  `json:"duration"`
            Published bool  `json:"published"`
            Cron string  `json:"cron"`
            Start string  `json:"start"`
         
    }
    
    // PromotionAuthor used by Cart
    type PromotionAuthor struct {

        
            ModifiedBy string  `json:"modified_by"`
            CreatedBy string  `json:"created_by"`
         
    }
    
    // PromotionListItem used by Cart
    type PromotionListItem struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Stackable bool  `json:"stackable"`
            DateMeta PromotionDateMeta  `json:"date_meta"`
            PostOrderAction PromotionAction  `json:"post_order_action"`
            ApplyExclusive string  `json:"apply_exclusive"`
            Code string  `json:"code"`
            PromoGroup string  `json:"promo_group"`
            BuyRules map[string]ItemCriteria  `json:"buy_rules"`
            ApplyPriority float64  `json:"apply_priority"`
            ApplyAllDiscount bool  `json:"apply_all_discount"`
            Visiblility Visibility  `json:"visiblility"`
            PromotionType string  `json:"promotion_type"`
            DisplayMeta DisplayMeta1  `json:"display_meta"`
            Restrictions Restrictions1  `json:"restrictions"`
            ApplicationID string  `json:"application_id"`
            Ownership Ownership1  `json:"ownership"`
            DiscountRules []DiscountRule  `json:"discount_rules"`
            Schedule PromotionSchedule  `json:"_schedule"`
            Author PromotionAuthor  `json:"author"`
            Mode string  `json:"mode"`
            CalculateOn string  `json:"calculate_on"`
            Currency string  `json:"currency"`
         
    }
    
    // PromotionsResponse used by Cart
    type PromotionsResponse struct {

        
            Page Page  `json:"page"`
            Items PromotionListItem  `json:"items"`
         
    }
    
    // PromotionAdd used by Cart
    type PromotionAdd struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Stackable bool  `json:"stackable"`
            DateMeta PromotionDateMeta  `json:"date_meta"`
            PostOrderAction PromotionAction  `json:"post_order_action"`
            ApplyExclusive string  `json:"apply_exclusive"`
            Code string  `json:"code"`
            PromoGroup string  `json:"promo_group"`
            BuyRules map[string]ItemCriteria  `json:"buy_rules"`
            ApplyPriority float64  `json:"apply_priority"`
            ApplyAllDiscount bool  `json:"apply_all_discount"`
            Visiblility Visibility  `json:"visiblility"`
            PromotionType string  `json:"promotion_type"`
            DisplayMeta DisplayMeta1  `json:"display_meta"`
            Restrictions Restrictions1  `json:"restrictions"`
            ApplicationID string  `json:"application_id"`
            Ownership Ownership1  `json:"ownership"`
            DiscountRules []DiscountRule  `json:"discount_rules"`
            Schedule PromotionSchedule  `json:"_schedule"`
            Author PromotionAuthor  `json:"author"`
            Mode string  `json:"mode"`
            CalculateOn string  `json:"calculate_on"`
            Currency string  `json:"currency"`
         
    }
    
    // PromotionUpdate used by Cart
    type PromotionUpdate struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Stackable bool  `json:"stackable"`
            DateMeta PromotionDateMeta  `json:"date_meta"`
            PostOrderAction PromotionAction  `json:"post_order_action"`
            ApplyExclusive string  `json:"apply_exclusive"`
            Code string  `json:"code"`
            PromoGroup string  `json:"promo_group"`
            BuyRules map[string]ItemCriteria  `json:"buy_rules"`
            ApplyPriority float64  `json:"apply_priority"`
            ApplyAllDiscount bool  `json:"apply_all_discount"`
            Visiblility Visibility  `json:"visiblility"`
            PromotionType string  `json:"promotion_type"`
            DisplayMeta DisplayMeta1  `json:"display_meta"`
            Restrictions Restrictions1  `json:"restrictions"`
            ApplicationID string  `json:"application_id"`
            Ownership Ownership1  `json:"ownership"`
            DiscountRules []DiscountRule  `json:"discount_rules"`
            Schedule PromotionSchedule  `json:"_schedule"`
            Author PromotionAuthor  `json:"author"`
            Mode string  `json:"mode"`
            CalculateOn string  `json:"calculate_on"`
            Currency string  `json:"currency"`
         
    }
    
    // PromotionPartialUpdate used by Cart
    type PromotionPartialUpdate struct {

        
            Schedule PromotionSchedule  `json:"schedule"`
            Archive bool  `json:"archive"`
         
    }
    
    // ActivePromosResponse used by Cart
    type ActivePromosResponse struct {

        
            Subtitle string  `json:"subtitle"`
            IsHidden bool  `json:"is_hidden"`
            EntitySlug string  `json:"entity_slug"`
            EntityType string  `json:"entity_type"`
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
            Example string  `json:"example"`
            Title string  `json:"title"`
            Type string  `json:"type"`
            Description string  `json:"description"`
         
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
    
    // BaseInfo used by Cart
    type BaseInfo struct {

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // ProductImage used by Cart
    type ProductImage struct {

        
            AspectRatio string  `json:"aspect_ratio"`
            SecureURL string  `json:"secure_url"`
            URL string  `json:"url"`
         
    }
    
    // CategoryInfo used by Cart
    type CategoryInfo struct {

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // CartProduct used by Cart
    type CartProduct struct {

        
            Action ProductAction  `json:"action"`
            Slug string  `json:"slug"`
            Brand BaseInfo  `json:"brand"`
            UID float64  `json:"uid"`
            Type string  `json:"type"`
            Images []ProductImage  `json:"images"`
            NetQuantity NetQuantity  `json:"net_quantity"`
            Name string  `json:"name"`
            Categories []CategoryInfo  `json:"categories"`
         
    }
    
    // ProductAvailability used by Cart
    type ProductAvailability struct {

        
            OutOfStock bool  `json:"out_of_stock"`
            Deliverable bool  `json:"deliverable"`
            IsValid bool  `json:"is_valid"`
            Sizes []string  `json:"sizes"`
            OtherStoreQuantity float64  `json:"other_store_quantity"`
         
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

        
            Selling float64  `json:"selling"`
            AddOn float64  `json:"add_on"`
            Effective float64  `json:"effective"`
            CurrencyCode string  `json:"currency_code"`
            CurrencySymbol string  `json:"currency_symbol"`
            Marked float64  `json:"marked"`
         
    }
    
    // ProductPriceInfo used by Cart
    type ProductPriceInfo struct {

        
            Converted ProductPrice  `json:"converted"`
            Base ProductPrice  `json:"base"`
         
    }
    
    // BasePrice used by Cart
    type BasePrice struct {

        
            CurrencyCode string  `json:"currency_code"`
            CurrencySymbol string  `json:"currency_symbol"`
            Marked float64  `json:"marked"`
            Effective float64  `json:"effective"`
         
    }
    
    // ArticlePriceInfo used by Cart
    type ArticlePriceInfo struct {

        
            Converted BasePrice  `json:"converted"`
            Base BasePrice  `json:"base"`
         
    }
    
    // ProductArticle used by Cart
    type ProductArticle struct {

        
            ProductGroupTags []string  `json:"product_group_tags"`
            Quantity float64  `json:"quantity"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            UID string  `json:"uid"`
            Type string  `json:"type"`
            Size string  `json:"size"`
            Price ArticlePriceInfo  `json:"price"`
            Seller BaseInfo  `json:"seller"`
            Store BaseInfo  `json:"store"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
         
    }
    
    // DiscountRulesApp used by Cart
    type DiscountRulesApp struct {

        
            RawOffer map[string]interface{}  `json:"raw_offer"`
            ItemCriteria map[string]interface{}  `json:"item_criteria"`
            MatchedBuyRules []string  `json:"matched_buy_rules"`
            Offer map[string]interface{}  `json:"offer"`
         
    }
    
    // FreeGiftItem used by Cart
    type FreeGiftItem struct {

        
            ItemImagesURL []string  `json:"item_images_url"`
            ItemName string  `json:"item_name"`
            ItemID float64  `json:"item_id"`
            ItemSlug string  `json:"item_slug"`
            ItemBrandName string  `json:"item_brand_name"`
            ItemPriceDetails map[string]interface{}  `json:"item_price_details"`
         
    }
    
    // AppliedFreeArticles used by Cart
    type AppliedFreeArticles struct {

        
            Quantity float64  `json:"quantity"`
            FreeGiftItemDetails FreeGiftItem  `json:"free_gift_item_details"`
            ArticleID string  `json:"article_id"`
            ParentItemIdentifier string  `json:"parent_item_identifier"`
         
    }
    
    // AppliedPromotion used by Cart
    type AppliedPromotion struct {

        
            Amount float64  `json:"amount"`
            PromotionName string  `json:"promotion_name"`
            DiscountRules []DiscountRulesApp  `json:"discount_rules"`
            PromoID string  `json:"promo_id"`
            OfferText string  `json:"offer_text"`
            PromotionType string  `json:"promotion_type"`
            AppliedFreeArticles []AppliedFreeArticles  `json:"applied_free_articles"`
            PromotionGroup string  `json:"promotion_group"`
            ArticleQuantity float64  `json:"article_quantity"`
            BuyRules []BuyRules  `json:"buy_rules"`
            MrpPromotion bool  `json:"mrp_promotion"`
         
    }
    
    // CartProductInfo used by Cart
    type CartProductInfo struct {

        
            Key string  `json:"key"`
            IsSet bool  `json:"is_set"`
            CouponMessage string  `json:"coupon_message"`
            Quantity float64  `json:"quantity"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            Product CartProduct  `json:"product"`
            Availability ProductAvailability  `json:"availability"`
            PromoMeta PromoMeta  `json:"promo_meta"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
            Discount string  `json:"discount"`
            Message string  `json:"message"`
            PricePerUnit ProductPriceInfo  `json:"price_per_unit"`
            Price ProductPriceInfo  `json:"price"`
            Article ProductArticle  `json:"article"`
            PromotionsApplied []AppliedPromotion  `json:"promotions_applied"`
            BulkOffer map[string]interface{}  `json:"bulk_offer"`
         
    }
    
    // DisplayBreakup used by Cart
    type DisplayBreakup struct {

        
            Key string  `json:"key"`
            Display string  `json:"display"`
            Message []string  `json:"message"`
            CurrencyCode string  `json:"currency_code"`
            CurrencySymbol string  `json:"currency_symbol"`
            Value float64  `json:"value"`
         
    }
    
    // LoyaltyPoints used by Cart
    type LoyaltyPoints struct {

        
            Applicable float64  `json:"applicable"`
            Total float64  `json:"total"`
            Description string  `json:"description"`
            IsApplied bool  `json:"is_applied"`
         
    }
    
    // CouponBreakup used by Cart
    type CouponBreakup struct {

        
            CouponValue float64  `json:"coupon_value"`
            MinimumCartValue float64  `json:"minimum_cart_value"`
            CouponType string  `json:"coupon_type"`
            IsApplied bool  `json:"is_applied"`
            UID string  `json:"uid"`
            Title string  `json:"title"`
            Type string  `json:"type"`
            Code string  `json:"code"`
            Description string  `json:"description"`
            MaxDiscountValue float64  `json:"max_discount_value"`
            Message string  `json:"message"`
            SubTitle string  `json:"sub_title"`
            Value float64  `json:"value"`
         
    }
    
    // RawBreakup used by Cart
    type RawBreakup struct {

        
            CodCharge float64  `json:"cod_charge"`
            Subtotal float64  `json:"subtotal"`
            Total float64  `json:"total"`
            GstCharges float64  `json:"gst_charges"`
            DeliveryCharge float64  `json:"delivery_charge"`
            FyndCash float64  `json:"fynd_cash"`
            Vog float64  `json:"vog"`
            Coupon float64  `json:"coupon"`
            Discount float64  `json:"discount"`
            MrpTotal float64  `json:"mrp_total"`
            ConvenienceFee float64  `json:"convenience_fee"`
            YouSaved float64  `json:"you_saved"`
         
    }
    
    // CartBreakup used by Cart
    type CartBreakup struct {

        
            Display []DisplayBreakup  `json:"display"`
            LoyaltyPoints LoyaltyPoints  `json:"loyalty_points"`
            Coupon CouponBreakup  `json:"coupon"`
            Raw RawBreakup  `json:"raw"`
         
    }
    
    // OpenapiCartDetailsResponse used by Cart
    type OpenapiCartDetailsResponse struct {

        
            IsValid bool  `json:"is_valid"`
            Items []CartProductInfo  `json:"items"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            Message string  `json:"message"`
         
    }
    
    // OpenApiErrorResponse used by Cart
    type OpenApiErrorResponse struct {

        
            Message string  `json:"message"`
            Errors map[string]interface{}  `json:"errors"`
            Success bool  `json:"success"`
         
    }
    
    // ShippingAddress used by Cart
    type ShippingAddress struct {

        
            CountryCode string  `json:"country_code"`
            City string  `json:"city"`
            Address string  `json:"address"`
            Meta map[string]interface{}  `json:"meta"`
            Email string  `json:"email"`
            State string  `json:"state"`
            Area string  `json:"area"`
            AreaCodeSlug string  `json:"area_code_slug"`
            AreaCode string  `json:"area_code"`
            Country string  `json:"country"`
            AddressType string  `json:"address_type"`
            Phone float64  `json:"phone"`
            Pincode float64  `json:"pincode"`
            Landmark string  `json:"landmark"`
            Name string  `json:"name"`
         
    }
    
    // OpenApiCartServiceabilityRequest used by Cart
    type OpenApiCartServiceabilityRequest struct {

        
            CartItems CartItem  `json:"cart_items"`
            ShippingAddress ShippingAddress  `json:"shipping_address"`
         
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
    
    // OpenApiCartServiceabilityResponse used by Cart
    type OpenApiCartServiceabilityResponse struct {

        
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            Items []CartProductInfo  `json:"items"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            Message string  `json:"message"`
            IsValid bool  `json:"is_valid"`
         
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
    
    // OpenApiOrderItem used by Cart
    type OpenApiOrderItem struct {

        
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            CashbackApplied float64  `json:"cashback_applied"`
            PaymentMethods []MultiTenderPaymentMethod  `json:"payment_methods"`
            AmountPaid float64  `json:"amount_paid"`
            CodCharges float64  `json:"cod_charges"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            Meta CartItemMeta  `json:"meta"`
            ProductID float64  `json:"product_id"`
            Size string  `json:"size"`
            Quantity float64  `json:"quantity"`
            DeliveryCharges float64  `json:"delivery_charges"`
            PriceEffective float64  `json:"price_effective"`
            Files []OpenApiFiles  `json:"files"`
            Discount float64  `json:"discount"`
            EmployeeDiscount float64  `json:"employee_discount"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            PriceMarked float64  `json:"price_marked"`
         
    }
    
    // OpenApiPlatformCheckoutReq used by Cart
    type OpenApiPlatformCheckoutReq struct {

        
            CartItems []OpenApiOrderItem  `json:"cart_items"`
            CouponValue float64  `json:"coupon_value"`
            OrderID string  `json:"order_id"`
            CouponCode string  `json:"coupon_code"`
            Comment string  `json:"comment"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            Files []OpenApiFiles  `json:"files"`
            EmployeeDiscount map[string]interface{}  `json:"employee_discount"`
            CashbackApplied float64  `json:"cashback_applied"`
            PaymentMethods []MultiTenderPaymentMethod  `json:"payment_methods"`
            CodCharges float64  `json:"cod_charges"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            DeliveryCharges float64  `json:"delivery_charges"`
            CartValue float64  `json:"cart_value"`
            PaymentMode string  `json:"payment_mode"`
            Coupon string  `json:"coupon"`
            Gstin string  `json:"gstin"`
            BillingAddress ShippingAddress  `json:"billing_address"`
            CurrencyCode string  `json:"currency_code"`
            ShippingAddress ShippingAddress  `json:"shipping_address"`
         
    }
    
    // OpenApiCheckoutResponse used by Cart
    type OpenApiCheckoutResponse struct {

        
            Message string  `json:"message"`
            OrderRefID string  `json:"order_ref_id"`
            OrderID string  `json:"order_id"`
            Success bool  `json:"success"`
         
    }
    
    // AbandonedCart used by Cart
    type AbandonedCart struct {

        
            FcIndexMap []float64  `json:"fc_index_map"`
            IsDefault bool  `json:"is_default"`
            UserID string  `json:"user_id"`
            OrderID string  `json:"order_id"`
            Meta map[string]interface{}  `json:"meta"`
            ID string  `json:"_id"`
            Cashback map[string]interface{}  `json:"cashback"`
            CheckoutMode string  `json:"checkout_mode"`
            AppID string  `json:"app_id"`
            MergeQty bool  `json:"merge_qty"`
            IsArchive bool  `json:"is_archive"`
            CodCharges map[string]interface{}  `json:"cod_charges"`
            DeliveryCharges map[string]interface{}  `json:"delivery_charges"`
            Promotion map[string]interface{}  `json:"promotion"`
            CreatedOn string  `json:"created_on"`
            PickUpCustomerDetails map[string]interface{}  `json:"pick_up_customer_details"`
            FyndCredits map[string]interface{}  `json:"fynd_credits"`
            ExpireAt string  `json:"expire_at"`
            Comment string  `json:"comment"`
            UID float64  `json:"uid"`
            Payments map[string]interface{}  `json:"payments"`
            Articles []map[string]interface{}  `json:"articles"`
            PaymentMethods []map[string]interface{}  `json:"payment_methods"`
            IsActive bool  `json:"is_active"`
            CartValue float64  `json:"cart_value"`
            PaymentMode string  `json:"payment_mode"`
            Coupon map[string]interface{}  `json:"coupon"`
            Gstin string  `json:"gstin"`
            Discount float64  `json:"discount"`
            BulkCouponDiscount float64  `json:"bulk_coupon_discount"`
            BuyNow bool  `json:"buy_now"`
            Shipments []map[string]interface{}  `json:"shipments"`
            LastModified string  `json:"last_modified"`
         
    }
    
    // AbandonedCartResponse used by Cart
    type AbandonedCartResponse struct {

        
            Page Page  `json:"page"`
            Result map[string]interface{}  `json:"result"`
            Success bool  `json:"success"`
            Items []AbandonedCart  `json:"items"`
            Message string  `json:"message"`
         
    }
    
    // CartCurrency used by Cart
    type CartCurrency struct {

        
            Symbol string  `json:"symbol"`
            Code string  `json:"code"`
         
    }
    
    // CartDetailResponse used by Cart
    type CartDetailResponse struct {

        
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            Message string  `json:"message"`
            CouponText string  `json:"coupon_text"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            ID string  `json:"id"`
            CheckoutMode string  `json:"checkout_mode"`
            Comment string  `json:"comment"`
            Gstin string  `json:"gstin"`
            Items []CartProductInfo  `json:"items"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            BuyNow bool  `json:"buy_now"`
            Currency CartCurrency  `json:"currency"`
            IsValid bool  `json:"is_valid"`
            LastModified string  `json:"last_modified"`
         
    }
    
    // AddProductCart used by Cart
    type AddProductCart struct {

        
            ProductGroupTags []string  `json:"product_group_tags"`
            Display string  `json:"display"`
            StoreID float64  `json:"store_id"`
            ArticleID string  `json:"article_id"`
            ItemID float64  `json:"item_id"`
            ItemSize string  `json:"item_size"`
            Quantity float64  `json:"quantity"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            Pos bool  `json:"pos"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            ArticleAssignment map[string]interface{}  `json:"article_assignment"`
            SellerID float64  `json:"seller_id"`
         
    }
    
    // AddCartRequest used by Cart
    type AddCartRequest struct {

        
            Items []AddProductCart  `json:"items"`
         
    }
    
    // AddCartDetailResponse used by Cart
    type AddCartDetailResponse struct {

        
            Message string  `json:"message"`
            Partial bool  `json:"partial"`
            Cart CartDetailResponse  `json:"cart"`
            Success bool  `json:"success"`
         
    }
    
    // UpdateProductCart used by Cart
    type UpdateProductCart struct {

        
            ArticleID string  `json:"article_id"`
            ItemID float64  `json:"item_id"`
            ItemSize string  `json:"item_size"`
            Quantity float64  `json:"quantity"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            ItemIndex float64  `json:"item_index"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
         
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
    
    // RewardsRule used by Rewards
    type RewardsRule struct {

        
            Amount float64  `json:"amount"`
         
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
    
    // UserDetails used by Discount
    type UserDetails struct {

        
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
         
    }
    
    // BadRequestObject used by Discount
    type BadRequestObject struct {

        
            Message string  `json:"message"`
         
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
    
    // EntityTypesResponse used by AuditTrail
    type EntityTypesResponse struct {

        
            Items []EntityTypeObj  `json:"items"`
         
    }
    
    // EntityTypeObj used by AuditTrail
    type EntityTypeObj struct {

        
            EntityValue string  `json:"entity_value"`
            DisplayName string  `json:"display_name"`
         
    }
    

    
    // ApplicationServiceabilityConfig used by Logistic
    type ApplicationServiceabilityConfig struct {

        
            ChannelType string  `json:"channel_type"`
            ServiceabilityType string  `json:"serviceability_type"`
            ChannelID string  `json:"channel_id"`
         
    }
    
    // ServiceabilityrErrorResponse used by Logistic
    type ServiceabilityrErrorResponse struct {

        
            Type string  `json:"type"`
            Value string  `json:"value"`
            Message string  `json:"message"`
         
    }
    
    // ApplicationServiceabilityConfigResponse used by Logistic
    type ApplicationServiceabilityConfigResponse struct {

        
            Success bool  `json:"success"`
            Data ApplicationServiceabilityConfig  `json:"data"`
            Error ServiceabilityrErrorResponse  `json:"error"`
         
    }
    
    // EntityRegionView_Request used by Logistic
    type EntityRegionView_Request struct {

        
            SubType []string  `json:"sub_type"`
            ParentID []string  `json:"parent_id"`
         
    }
    
    // EntityRegionView_Error used by Logistic
    type EntityRegionView_Error struct {

        
            Type string  `json:"type"`
            Value string  `json:"value"`
            Message string  `json:"message"`
         
    }
    
    // EntityRegionView_page used by Logistic
    type EntityRegionView_page struct {

        
            Type string  `json:"type"`
            ItemTotal float64  `json:"item_total"`
            Current float64  `json:"current"`
            Size float64  `json:"size"`
            HasNext bool  `json:"has_next"`
         
    }
    
    // EntityRegionView_Items used by Logistic
    type EntityRegionView_Items struct {

        
            Name string  `json:"name"`
            SubType string  `json:"sub_type"`
            UID string  `json:"uid"`
         
    }
    
    // EntityRegionView_Response used by Logistic
    type EntityRegionView_Response struct {

        
            Error EntityRegionView_Error  `json:"error"`
            Page EntityRegionView_page  `json:"page"`
            Data []EntityRegionView_Items  `json:"data"`
            Success bool  `json:"success"`
         
    }
    
    // ListViewSummary used by Logistic
    type ListViewSummary struct {

        
            TotalPincodesServed float64  `json:"total_pincodes_served"`
            TotalZones float64  `json:"total_zones"`
            TotalActiveZones float64  `json:"total_active_zones"`
         
    }
    
    // ZoneDataItem used by Logistic
    type ZoneDataItem struct {

        
            Type string  `json:"type"`
            ItemTotal float64  `json:"item_total"`
            Current float64  `json:"current"`
            Size float64  `json:"size"`
            HasNext bool  `json:"has_next"`
         
    }
    
    // ListViewProduct used by Logistic
    type ListViewProduct struct {

        
            Count float64  `json:"count"`
            Type string  `json:"type"`
         
    }
    
    // ListViewChannels used by Logistic
    type ListViewChannels struct {

        
            ChannelType string  `json:"channel_type"`
            ChannelID string  `json:"channel_id"`
         
    }
    
    // ListViewItems used by Logistic
    type ListViewItems struct {

        
            ZoneID string  `json:"zone_id"`
            Product ListViewProduct  `json:"product"`
            IsActive bool  `json:"is_active"`
            Channels ListViewChannels  `json:"channels"`
            PincodesCount float64  `json:"pincodes_count"`
            CompanyID float64  `json:"company_id"`
            Slug string  `json:"slug"`
            StoresCount float64  `json:"stores_count"`
            Name string  `json:"name"`
         
    }
    
    // ListViewResponse used by Logistic
    type ListViewResponse struct {

        
            Summary []ListViewSummary  `json:"summary"`
            Page []ZoneDataItem  `json:"page"`
            Items []ListViewItems  `json:"items"`
         
    }
    
    // CompanyStoreView_PageItems used by Logistic
    type CompanyStoreView_PageItems struct {

        
            Type string  `json:"type"`
            ItemTotal float64  `json:"item_total"`
            Current float64  `json:"current"`
            Size float64  `json:"size"`
            HasNext bool  `json:"has_next"`
         
    }
    
    // CompanyStoreView_Response used by Logistic
    type CompanyStoreView_Response struct {

        
            Page []CompanyStoreView_PageItems  `json:"page"`
            Items []map[string]interface{}  `json:"items"`
         
    }
    
    // GetZoneDataViewChannels used by Logistic
    type GetZoneDataViewChannels struct {

        
            ChannelType string  `json:"channel_type"`
            ChannelID string  `json:"channel_id"`
         
    }
    
    // ZoneProductTypes used by Logistic
    type ZoneProductTypes struct {

        
            Tags []string  `json:"tags"`
            Type string  `json:"type"`
         
    }
    
    // ZoneMappingType used by Logistic
    type ZoneMappingType struct {

        
            Pincode []string  `json:"pincode"`
            State []string  `json:"state"`
            Country string  `json:"country"`
         
    }
    
    // GetZoneDataViewItems used by Logistic
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
    
    // GetSingleZoneDataViewResponse used by Logistic
    type GetSingleZoneDataViewResponse struct {

        
            Data GetZoneDataViewItems  `json:"data"`
         
    }
    
    // UpdateZoneData used by Logistic
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
    
    // ZoneUpdateRequest used by Logistic
    type ZoneUpdateRequest struct {

        
            Identifier string  `json:"identifier"`
            Data UpdateZoneData  `json:"data"`
         
    }
    
    // ZoneSuccessResponse used by Logistic
    type ZoneSuccessResponse struct {

        
            StatusCode float64  `json:"status_code"`
            Success bool  `json:"success"`
         
    }
    
    // CreateZoneData used by Logistic
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
    
    // ZoneRequest used by Logistic
    type ZoneRequest struct {

        
            Identifier string  `json:"identifier"`
            Data CreateZoneData  `json:"data"`
         
    }
    
    // ZoneResponse used by Logistic
    type ZoneResponse struct {

        
            StatusCode float64  `json:"status_code"`
            ZoneID string  `json:"zone_id"`
            Success bool  `json:"success"`
         
    }
    
    // GetZoneFromPincodeViewRequest used by Logistic
    type GetZoneFromPincodeViewRequest struct {

        
            Pincode string  `json:"pincode"`
            Country string  `json:"country"`
         
    }
    
    // GetZoneFromPincodeViewResponse used by Logistic
    type GetZoneFromPincodeViewResponse struct {

        
            ServiceabilityType string  `json:"serviceability_type"`
            Zones []string  `json:"zones"`
         
    }
    
    // OpeningClosing used by Logistic
    type OpeningClosing struct {

        
            Hour float64  `json:"hour"`
            Minute float64  `json:"minute"`
         
    }
    
    // TimmingResponse used by Logistic
    type TimmingResponse struct {

        
            Opening OpeningClosing  `json:"opening"`
            Open bool  `json:"open"`
            Closing OpeningClosing  `json:"closing"`
            Weekday string  `json:"weekday"`
         
    }
    
    // IntegrationTypeResponse used by Logistic
    type IntegrationTypeResponse struct {

        
            Order string  `json:"order"`
            Inventory string  `json:"inventory"`
         
    }
    
    // Dp used by Logistic
    type Dp struct {

        
            TransportMode string  `json:"transport_mode"`
            FmPriority float64  `json:"fm_priority"`
            AreaCode float64  `json:"area_code"`
            ExternalAccountID string  `json:"external_account_id"`
            PaymentMode string  `json:"payment_mode"`
            LmPriority float64  `json:"lm_priority"`
            InternalAccountID string  `json:"internal_account_id"`
            AssignDpFromSb bool  `json:"assign_dp_from_sb"`
            RvpPriority float64  `json:"rvp_priority"`
            Operations []string  `json:"operations"`
         
    }
    
    // LogisticsResponse used by Logistic
    type LogisticsResponse struct {

        
            Dp Dp  `json:"dp"`
            Override bool  `json:"override"`
         
    }
    
    // ModifiedByResponse used by Logistic
    type ModifiedByResponse struct {

        
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
         
    }
    
    // AddressResponse used by Logistic
    type AddressResponse struct {

        
            Address1 string  `json:"address1"`
            Landmark string  `json:"landmark"`
            Address2 string  `json:"address2"`
            Longitude float64  `json:"longitude"`
            State string  `json:"state"`
            Latitude float64  `json:"latitude"`
            City string  `json:"city"`
            Pincode float64  `json:"pincode"`
            Country string  `json:"country"`
         
    }
    
    // DocumentsResponse used by Logistic
    type DocumentsResponse struct {

        
            LegalName string  `json:"legal_name"`
            Type string  `json:"type"`
            Value string  `json:"value"`
            Verified bool  `json:"verified"`
         
    }
    
    // MobileNo used by Logistic
    type MobileNo struct {

        
            Number string  `json:"number"`
            CountryCode float64  `json:"country_code"`
         
    }
    
    // ManagerResponse used by Logistic
    type ManagerResponse struct {

        
            Email string  `json:"email"`
            MobileNo MobileNo  `json:"mobile_no"`
            Name string  `json:"name"`
         
    }
    
    // EwayBillResponse used by Logistic
    type EwayBillResponse struct {

        
            Enabled bool  `json:"enabled"`
         
    }
    
    // EinvoiceResponse used by Logistic
    type EinvoiceResponse struct {

        
            Enabled bool  `json:"enabled"`
         
    }
    
    // GstCredentialsResponse used by Logistic
    type GstCredentialsResponse struct {

        
            EWaybill EwayBillResponse  `json:"e_waybill"`
            EInvoice EinvoiceResponse  `json:"e_invoice"`
         
    }
    
    // ContactNumberResponse used by Logistic
    type ContactNumberResponse struct {

        
            Number string  `json:"number"`
            CountryCode float64  `json:"country_code"`
         
    }
    
    // CreatedByResponse used by Logistic
    type CreatedByResponse struct {

        
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
         
    }
    
    // ProductReturnConfigResponse used by Logistic
    type ProductReturnConfigResponse struct {

        
            OnSameStore bool  `json:"on_same_store"`
         
    }
    
    // WarningsResponse used by Logistic
    type WarningsResponse struct {

        
            StoreAddress string  `json:"store_address"`
         
    }
    
    // ItemResponse used by Logistic
    type ItemResponse struct {

        
            Timing []TimmingResponse  `json:"timing"`
            ModifiedOn string  `json:"modified_on"`
            DisplayName string  `json:"display_name"`
            IntegrationType IntegrationTypeResponse  `json:"integration_type"`
            Logistics LogisticsResponse  `json:"logistics"`
            StoreType string  `json:"store_type"`
            ModifiedBy ModifiedByResponse  `json:"modified_by"`
            Address AddressResponse  `json:"address"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Documents []DocumentsResponse  `json:"documents"`
            Manager ManagerResponse  `json:"manager"`
            GstCredentials GstCredentialsResponse  `json:"gst_credentials"`
            Code string  `json:"code"`
            Name string  `json:"name"`
            CreatedOn string  `json:"created_on"`
            SubType string  `json:"sub_type"`
            Stage string  `json:"stage"`
            VerifiedOn string  `json:"verified_on"`
            ContactNumbers []ContactNumberResponse  `json:"contact_numbers"`
            CreatedBy CreatedByResponse  `json:"created_by"`
            ProductReturnConfig ProductReturnConfigResponse  `json:"product_return_config"`
            NotificationEmails []string  `json:"notification_emails"`
            Company float64  `json:"company"`
            Cls string  `json:"_cls"`
            UID float64  `json:"uid"`
            VerifiedBy ModifiedByResponse  `json:"verified_by"`
            Warnings WarningsResponse  `json:"warnings"`
            CompanyID float64  `json:"company_id"`
         
    }
    
    // GetStoresViewResponse used by Logistic
    type GetStoresViewResponse struct {

        
            Page PageResponse  `json:"page"`
            Items []ItemResponse  `json:"items"`
         
    }
    
