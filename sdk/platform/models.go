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
    

    
    // FilterInfoOption used by Order
    type FilterInfoOption struct {

        
            Name string  `json:"name"`
            Text string  `json:"text"`
            Value string  `json:"value"`
         
    }
    
    // FiltersInfo used by Order
    type FiltersInfo struct {

        
            Options []FilterInfoOption  `json:"options"`
            Type string  `json:"type"`
            Text string  `json:"text"`
            Value string  `json:"value"`
         
    }
    
    // GSTDetailsData used by Order
    type GSTDetailsData struct {

        
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            ValueOfGood float64  `json:"value_of_good"`
            GstinCode string  `json:"gstin_code"`
            GstFee float64  `json:"gst_fee"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
         
    }
    
    // Prices used by Order
    type Prices struct {

        
            CashbackApplied float64  `json:"cashback_applied"`
            AmountPaidRoundoff float64  `json:"amount_paid_roundoff"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            FyndCredits float64  `json:"fynd_credits"`
            PromotionEffectiveDiscount float64  `json:"promotion_effective_discount"`
            RefundAmount float64  `json:"refund_amount"`
            AmountPaid float64  `json:"amount_paid"`
            PriceMarked float64  `json:"price_marked"`
            CodCharges float64  `json:"cod_charges"`
            ValueOfGood float64  `json:"value_of_good"`
            CouponValue float64  `json:"coupon_value"`
            RefundCredit float64  `json:"refund_credit"`
            Cashback float64  `json:"cashback"`
            DeliveryCharge float64  `json:"delivery_charge"`
            Discount float64  `json:"discount"`
            PriceEffective float64  `json:"price_effective"`
         
    }
    
    // PlatformItem used by Order
    type PlatformItem struct {

        
            CanReturn bool  `json:"can_return"`
            CanCancel bool  `json:"can_cancel"`
            L3Category float64  `json:"l3_category"`
            Color string  `json:"color"`
            DepartmentID float64  `json:"department_id"`
            Images []string  `json:"images"`
            ID float64  `json:"id"`
            Size string  `json:"size"`
            Code string  `json:"code"`
            Name string  `json:"name"`
            L3CategoryName string  `json:"l3_category_name"`
            Image []string  `json:"image"`
            L1Category []string  `json:"l1_category"`
         
    }
    
    // BagUnit used by Order
    type BagUnit struct {

        
            TotalShipmentBags float64  `json:"total_shipment_bags"`
            ShipmentID string  `json:"shipment_id"`
            Status map[string]interface{}  `json:"status"`
            Gst GSTDetailsData  `json:"gst"`
            OrderingChannel string  `json:"ordering_channel"`
            BagID float64  `json:"bag_id"`
            ItemQuantity float64  `json:"item_quantity"`
            Prices Prices  `json:"prices"`
            Item PlatformItem  `json:"item"`
         
    }
    
    // ShipmentStatus used by Order
    type ShipmentStatus struct {

        
            OpsStatus string  `json:"ops_status"`
            Status string  `json:"status"`
            HexCode string  `json:"hex_code"`
            Title string  `json:"title"`
            ActualStatus string  `json:"actual_status"`
         
    }
    
    // UserDataInfo used by Order
    type UserDataInfo struct {

        
            AvisUserID string  `json:"avis_user_id"`
            Email string  `json:"email"`
            Gender string  `json:"gender"`
            Mobile string  `json:"mobile"`
            IsAnonymousUser bool  `json:"is_anonymous_user"`
            FirstName string  `json:"first_name"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            LastName string  `json:"last_name"`
         
    }
    
    // ShipmentItemFulFillingStore used by Order
    type ShipmentItemFulFillingStore struct {

        
            Code string  `json:"code"`
            ID string  `json:"id"`
         
    }
    
    // PaymentModeInfo used by Order
    type PaymentModeInfo struct {

        
            Logo string  `json:"logo"`
            Type string  `json:"type"`
         
    }
    
    // ShipmentItem used by Order
    type ShipmentItem struct {

        
            TotalBagsCount float64  `json:"total_bags_count"`
            Channel map[string]interface{}  `json:"channel"`
            Bags []BagUnit  `json:"bags"`
            Sla map[string]interface{}  `json:"sla"`
            ShipmentStatus ShipmentStatus  `json:"shipment_status"`
            PaymentMethods map[string]interface{}  `json:"payment_methods"`
            ID string  `json:"id"`
            CreatedAt string  `json:"created_at"`
            Application map[string]interface{}  `json:"application"`
            ShipmentCreatedAt float64  `json:"shipment_created_at"`
            FulfillingCentre string  `json:"fulfilling_centre"`
            User UserDataInfo  `json:"user"`
            FulfillingStore ShipmentItemFulFillingStore  `json:"fulfilling_store"`
            Prices Prices  `json:"prices"`
            PaymentModeInfo PaymentModeInfo  `json:"payment_mode_info"`
            TotalShipmentsInOrder float64  `json:"total_shipments_in_order"`
         
    }
    
    // ShipmentInternalPlatformViewResponse used by Order
    type ShipmentInternalPlatformViewResponse struct {

        
            Page map[string]interface{}  `json:"page"`
            Filters []FiltersInfo  `json:"filters"`
            AppliedFilters map[string]interface{}  `json:"applied_filters"`
            Items []ShipmentItem  `json:"items"`
         
    }
    
    // Error used by Order
    type Error struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // BagStateMapper used by Order
    type BagStateMapper struct {

        
            JourneyType string  `json:"journey_type"`
            NotifyCustomer bool  `json:"notify_customer"`
            AppDisplayName string  `json:"app_display_name"`
            AppStateName string  `json:"app_state_name"`
            StateType string  `json:"state_type"`
            IsActive bool  `json:"is_active"`
            DisplayName string  `json:"display_name"`
            Name string  `json:"name"`
            BsID float64  `json:"bs_id"`
            AppFacing bool  `json:"app_facing"`
         
    }
    
    // BagStatusHistory used by Order
    type BagStatusHistory struct {

        
            UpdatedAt string  `json:"updated_at"`
            DeliveryPartnerID float64  `json:"delivery_partner_id"`
            ShipmentID string  `json:"shipment_id"`
            Status string  `json:"status"`
            BagStateMapper BagStateMapper  `json:"bag_state_mapper"`
            Reasons []map[string]interface{}  `json:"reasons"`
            StoreID float64  `json:"store_id"`
            BshID float64  `json:"bsh_id"`
            AppDisplayName bool  `json:"app_display_name"`
            StateType string  `json:"state_type"`
            Forward bool  `json:"forward"`
            CreatedAt string  `json:"created_at"`
            BagID float64  `json:"bag_id"`
            DeliveryAwbNumber string  `json:"delivery_awb_number"`
            KafkaSync bool  `json:"kafka_sync"`
            DisplayName bool  `json:"display_name"`
            StateID float64  `json:"state_id"`
         
    }
    
    // UserDetailsData used by Order
    type UserDetailsData struct {

        
            City string  `json:"city"`
            Country string  `json:"country"`
            State string  `json:"state"`
            Email string  `json:"email"`
            Address string  `json:"address"`
            Pincode string  `json:"pincode"`
            Name string  `json:"name"`
            Phone string  `json:"phone"`
         
    }
    
    // DPDetailsData used by Order
    type DPDetailsData struct {

        
            TrackURL string  `json:"track_url"`
            Country string  `json:"country"`
            AwbNo string  `json:"awb_no"`
            EwayBillID string  `json:"eway_bill_id"`
            ID string  `json:"id"`
            GstTag string  `json:"gst_tag"`
            Pincode string  `json:"pincode"`
            Name string  `json:"name"`
         
    }
    
    // TrackingList used by Order
    type TrackingList struct {

        
            Status string  `json:"status"`
            IsCurrent bool  `json:"is_current"`
            IsPassed bool  `json:"is_passed"`
            Time string  `json:"time"`
            Text string  `json:"text"`
         
    }
    
    // OrderBrandName used by Order
    type OrderBrandName struct {

        
            BrandName string  `json:"brand_name"`
            ModifiedOn float64  `json:"modified_on"`
            Company string  `json:"company"`
            Logo string  `json:"logo"`
            CreatedOn float64  `json:"created_on"`
            ID float64  `json:"id"`
         
    }
    
    // PlatformDeliveryAddress used by Order
    type PlatformDeliveryAddress struct {

        
            UpdatedAt string  `json:"updated_at"`
            Address2 string  `json:"address2"`
            Landmark string  `json:"landmark"`
            Country string  `json:"country"`
            State string  `json:"state"`
            Email string  `json:"email"`
            AddressType string  `json:"address_type"`
            ContactPerson string  `json:"contact_person"`
            Version string  `json:"version"`
            CreatedAt string  `json:"created_at"`
            Pincode string  `json:"pincode"`
            AddressCategory string  `json:"address_category"`
            Latitude float64  `json:"latitude"`
            Address1 string  `json:"address1"`
            City string  `json:"city"`
            Longitude float64  `json:"longitude"`
            Area string  `json:"area"`
            Phone string  `json:"phone"`
         
    }
    
    // BagGST used by Order
    type BagGST struct {

        
            HsnCode string  `json:"hsn_code"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            ValueOfGood float64  `json:"value_of_good"`
            GstinCode string  `json:"gstin_code"`
            GstFee float64  `json:"gst_fee"`
            GstTag string  `json:"gst_tag"`
            IsDefaultHsnCode bool  `json:"is_default_hsn_code"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
         
    }
    
    // OrderBagArticle used by Order
    type OrderBagArticle struct {

        
            UID string  `json:"uid"`
            Identifiers map[string]interface{}  `json:"identifiers"`
            ReturnConfig map[string]interface{}  `json:"return_config"`
         
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

        
            PromoID string  `json:"promo_id"`
            DiscountRules []DiscountRules  `json:"discount_rules"`
            MrpPromotion bool  `json:"mrp_promotion"`
            BuyRules []BuyRules  `json:"buy_rules"`
            ArticleQuantity float64  `json:"article_quantity"`
            PromotionName string  `json:"promotion_name"`
            Amount float64  `json:"amount"`
            PromotionType string  `json:"promotion_type"`
         
    }
    
    // Identifier used by Order
    type Identifier struct {

        
            Ean string  `json:"ean"`
         
    }
    
    // FinancialBreakup used by Order
    type FinancialBreakup struct {

        
            CashbackApplied float64  `json:"cashback_applied"`
            AmountPaid float64  `json:"amount_paid"`
            ItemName string  `json:"item_name"`
            AddedToFyndCash bool  `json:"added_to_fynd_cash"`
            GstFee string  `json:"gst_fee"`
            Size string  `json:"size"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            PriceEffective float64  `json:"price_effective"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            HsnCode string  `json:"hsn_code"`
            FyndCredits float64  `json:"fynd_credits"`
            ValueOfGood float64  `json:"value_of_good"`
            CodCharges float64  `json:"cod_charges"`
            PmPriceSplit map[string]interface{}  `json:"pm_price_split"`
            CouponValue float64  `json:"coupon_value"`
            GstTag string  `json:"gst_tag"`
            PromotionEffectiveDiscount float64  `json:"promotion_effective_discount"`
            DeliveryCharge float64  `json:"delivery_charge"`
            Discount float64  `json:"discount"`
            TotalUnits float64  `json:"total_units"`
            TransferPrice float64  `json:"transfer_price"`
            AmountPaidRoundoff float64  `json:"amount_paid_roundoff"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            Identifiers Identifier  `json:"identifiers"`
            PriceMarked float64  `json:"price_marked"`
            RefundCredit float64  `json:"refund_credit"`
            Cashback float64  `json:"cashback"`
         
    }
    
    // BagConfigs used by Order
    type BagConfigs struct {

        
            IsReturnable bool  `json:"is_returnable"`
            IsActive bool  `json:"is_active"`
            IsCustomerReturnAllowed bool  `json:"is_customer_return_allowed"`
            AllowForceReturn bool  `json:"allow_force_return"`
            CanBeCancelled bool  `json:"can_be_cancelled"`
            EnableTracking bool  `json:"enable_tracking"`
         
    }
    
    // OrderBags used by Order
    type OrderBags struct {

        
            ParentPromoBags map[string]interface{}  `json:"parent_promo_bags"`
            Brand OrderBrandName  `json:"brand"`
            DeliveryAddress PlatformDeliveryAddress  `json:"delivery_address"`
            GstDetails BagGST  `json:"gst_details"`
            SellerIdentifier string  `json:"seller_identifier"`
            Article OrderBagArticle  `json:"article"`
            Quantity float64  `json:"quantity"`
            AppliedPromos []AppliedPromos  `json:"applied_promos"`
            LineNumber float64  `json:"line_number"`
            FinancialBreakup []FinancialBreakup  `json:"financial_breakup"`
            EntityType string  `json:"entity_type"`
            BagID float64  `json:"bag_id"`
            CurrentStatus string  `json:"current_status"`
            DisplayName string  `json:"display_name"`
            Prices Prices  `json:"prices"`
            Item PlatformItem  `json:"item"`
            BagConfigs BagConfigs  `json:"bag_configs"`
            Identifier string  `json:"identifier"`
         
    }
    
    // ShipmentPayments used by Order
    type ShipmentPayments struct {

        
            Logo string  `json:"logo"`
            Mode string  `json:"mode"`
            Source string  `json:"source"`
         
    }
    
    // ShipmentStatusData used by Order
    type ShipmentStatusData struct {

        
            ShipmentID string  `json:"shipment_id"`
            Status string  `json:"status"`
            ID float64  `json:"id"`
            CreatedAt string  `json:"created_at"`
            BagList []float64  `json:"bag_list"`
         
    }
    
    // OrderDetailsData used by Order
    type OrderDetailsData struct {

        
            Source string  `json:"source"`
            TaxDetails map[string]interface{}  `json:"tax_details"`
            OrderDate string  `json:"order_date"`
            OrderValue string  `json:"order_value"`
            CodCharges string  `json:"cod_charges"`
            OrderingChannel string  `json:"ordering_channel"`
            OrderingChannelLogo map[string]interface{}  `json:"ordering_channel_logo"`
            FyndOrderID string  `json:"fynd_order_id"`
            AffiliateID string  `json:"affiliate_id"`
         
    }
    
    // FulfillingStore used by Order
    type FulfillingStore struct {

        
            Country string  `json:"country"`
            State string  `json:"state"`
            FulfillmentChannel string  `json:"fulfillment_channel"`
            ContactPerson string  `json:"contact_person"`
            ID string  `json:"id"`
            Address string  `json:"address"`
            Pincode string  `json:"pincode"`
            StoreName string  `json:"store_name"`
            Code string  `json:"code"`
            City string  `json:"city"`
            Meta map[string]interface{}  `json:"meta"`
            Phone string  `json:"phone"`
         
    }
    
    // ShipmentInfoResponse used by Order
    type ShipmentInfoResponse struct {

        
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            JourneyType string  `json:"journey_type"`
            CanCancel bool  `json:"can_cancel"`
            BagStatusHistory []BagStatusHistory  `json:"bag_status_history"`
            Company map[string]interface{}  `json:"company"`
            Vertical string  `json:"vertical"`
            Coupon map[string]interface{}  `json:"coupon"`
            OrderingStore map[string]interface{}  `json:"ordering_store"`
            Escalation map[string]interface{}  `json:"escalation"`
            UserID string  `json:"user_id"`
            RefundDetails map[string]interface{}  `json:"refund_details"`
            RefundText string  `json:"refund_text"`
            BeneficiaryDetails bool  `json:"beneficiary_details"`
            IsPdsr string  `json:"is_pdsr"`
            TrackingURL string  `json:"tracking_url"`
            PayButton string  `json:"pay_button"`
            IsInvoiced bool  `json:"is_invoiced"`
            PlatformLogo bool  `json:"platform_logo"`
            OrderCreatedTime string  `json:"order_created_time"`
            IsFyndCoupon bool  `json:"is_fynd_coupon"`
            UserAgent string  `json:"user_agent"`
            EmailID string  `json:"email_id"`
            FyndstoreEmp map[string]interface{}  `json:"fyndstore_emp"`
            GstDetails GSTDetailsData  `json:"gst_details"`
            IsPackagingOrder bool  `json:"is_packaging_order"`
            CurrentShipmentStatus map[string]interface{}  `json:"current_shipment_status"`
            CustomMeta []map[string]interface{}  `json:"custom_meta"`
            IsFyndStore string  `json:"is_fynd_store"`
            Mid string  `json:"mid"`
            Invoice map[string]interface{}  `json:"invoice"`
            ShipmentQuantity float64  `json:"shipment_quantity"`
            EnableDpTracking string  `json:"enable_dp_tracking"`
            SecuredDeliveryFlag string  `json:"secured_delivery_flag"`
            BillingDetails UserDetailsData  `json:"billing_details"`
            DeliveryDetails UserDetailsData  `json:"delivery_details"`
            DpDetails DPDetailsData  `json:"dp_details"`
            Items []map[string]interface{}  `json:"items"`
            KiranaStoreID string  `json:"kirana_store_id"`
            DeliveryStatus []map[string]interface{}  `json:"delivery_status"`
            ReplacementDetails string  `json:"replacement_details"`
            TrackingList []TrackingList  `json:"tracking_list"`
            DueDate string  `json:"due_date"`
            ForwardOrderStatus []map[string]interface{}  `json:"forward_order_status"`
            LockStatus string  `json:"lock_status"`
            UserInfo map[string]interface{}  `json:"user_info"`
            OrderType string  `json:"order_type"`
            StatusProgress float64  `json:"status_progress"`
            PaymentMode string  `json:"payment_mode"`
            Prices Prices  `json:"prices"`
            Bags []OrderBags  `json:"bags"`
            OperationalStatus string  `json:"operational_status"`
            Payments ShipmentPayments  `json:"payments"`
            TotalBags float64  `json:"total_bags"`
            CanBreak string  `json:"can_break"`
            ShipmentID string  `json:"shipment_id"`
            Status ShipmentStatusData  `json:"status"`
            PackagingType string  `json:"packaging_type"`
            OrderStatus map[string]interface{}  `json:"order_status"`
            CreditNoteID string  `json:"credit_note_id"`
            ChildNodes []string  `json:"child_nodes"`
            CanReturn bool  `json:"can_return"`
            ShipmentImages []string  `json:"shipment_images"`
            BankData map[string]interface{}  `json:"bank_data"`
            ShipmentStatus string  `json:"shipment_status"`
            TotalItems float64  `json:"total_items"`
            DeliverySlot map[string]interface{}  `json:"delivery_slot"`
            ForwardShipmentStatus []map[string]interface{}  `json:"forward_shipment_status"`
            ForwardTrackingList []map[string]interface{}  `json:"forward_tracking_list"`
            Order OrderDetailsData  `json:"order"`
            PriorityText string  `json:"priority_text"`
            FulfillingStore FulfillingStore  `json:"fulfilling_store"`
            GoGreen bool  `json:"go_green"`
            PickedDate string  `json:"picked_date"`
            EnableTracking bool  `json:"enable_tracking"`
            IsNotFyndSource bool  `json:"is_not_fynd_source"`
         
    }
    
    // PlatformShipment used by Order
    type PlatformShipment struct {

        
            JourneyType string  `json:"journey_type"`
            BagStatusHistory []BagStatusHistory  `json:"bag_status_history"`
            Vertical string  `json:"vertical"`
            PlatformLogo string  `json:"platform_logo"`
            UserAgent string  `json:"user_agent"`
            GstDetails GSTDetailsData  `json:"gst_details"`
            CustomMeta []map[string]interface{}  `json:"custom_meta"`
            ShipmentQuantity float64  `json:"shipment_quantity"`
            EnableDpTracking string  `json:"enable_dp_tracking"`
            BillingDetails UserDetailsData  `json:"billing_details"`
            DeliveryDetails UserDetailsData  `json:"delivery_details"`
            DpDetails DPDetailsData  `json:"dp_details"`
            TrackingList []TrackingList  `json:"tracking_list"`
            PaymentMode string  `json:"payment_mode"`
            Prices Prices  `json:"prices"`
            Bags []OrderBags  `json:"bags"`
            OperationalStatus string  `json:"operational_status"`
            Payments ShipmentPayments  `json:"payments"`
            TotalBags float64  `json:"total_bags"`
            ShipmentID string  `json:"shipment_id"`
            Status ShipmentStatusData  `json:"status"`
            PackagingType string  `json:"packaging_type"`
            ShipmentImages []string  `json:"shipment_images"`
            ShipmentStatus string  `json:"shipment_status"`
            TotalItems float64  `json:"total_items"`
            DeliverySlot map[string]interface{}  `json:"delivery_slot"`
            Order OrderDetailsData  `json:"order"`
            PriorityText string  `json:"priority_text"`
            FulfillingStore FulfillingStore  `json:"fulfilling_store"`
            PickedDate string  `json:"picked_date"`
         
    }
    
    // OrderDict used by Order
    type OrderDict struct {

        
            FyndOrderID string  `json:"fynd_order_id"`
            ShipmentCount float64  `json:"shipment_count"`
            OrderDate string  `json:"order_date"`
         
    }
    
    // ShipmentDetailsResponse used by Order
    type ShipmentDetailsResponse struct {

        
            Shipments []PlatformShipment  `json:"shipments"`
            CustomMeta []map[string]interface{}  `json:"custom_meta"`
            Order OrderDict  `json:"order"`
            Success bool  `json:"success"`
         
    }
    
    // SubLane used by Order
    type SubLane struct {

        
            TotalItems float64  `json:"total_items"`
            Index float64  `json:"index"`
            Actions []map[string]interface{}  `json:"actions"`
            Value string  `json:"value"`
            Text string  `json:"text"`
         
    }
    
    // SuperLane used by Order
    type SuperLane struct {

        
            TotalItems float64  `json:"total_items"`
            Options []SubLane  `json:"options"`
            Text string  `json:"text"`
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

        
            Channel PlatformChannel  `json:"channel"`
            TotalOrderValue float64  `json:"total_order_value"`
            OrderValue float64  `json:"order_value"`
            BreakupValues []PlatformBreakupValues  `json:"breakup_values"`
            Shipments []PlatformShipment  `json:"shipments"`
            UserInfo UserDataInfo  `json:"user_info"`
            PaymentMode string  `json:"payment_mode"`
            OrderCreatedTime string  `json:"order_created_time"`
            OrderID string  `json:"order_id"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // OrderListingResponse used by Order
    type OrderListingResponse struct {

        
            Lane string  `json:"lane"`
            TotalCount float64  `json:"total_count"`
            Success bool  `json:"success"`
            Page Page  `json:"page"`
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

        
            Value float64  `json:"value"`
            Key string  `json:"key"`
            Text string  `json:"text"`
            Options []Options  `json:"options"`
         
    }
    
    // MetricCountResponse used by Order
    type MetricCountResponse struct {

        
            Items []MetricsCount  `json:"items"`
         
    }
    
    // PlatformTrack used by Order
    type PlatformTrack struct {

        
            UpdatedAt string  `json:"updated_at"`
            Status string  `json:"status"`
            Reason string  `json:"reason"`
            AccountName string  `json:"account_name"`
            LastLocationRecievedAt string  `json:"last_location_recieved_at"`
            UpdatedTime string  `json:"updated_time"`
            Meta map[string]interface{}  `json:"meta"`
            RawStatus string  `json:"raw_status"`
            Awb string  `json:"awb"`
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

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // OmsReports used by Order
    type OmsReports struct {

        
            Status string  `json:"status"`
            S3Key string  `json:"s3_key"`
            ReportID string  `json:"report_id"`
            ReportName string  `json:"report_name"`
            ReportRequestedAt string  `json:"report_requested_at"`
            ReportCreatedAt string  `json:"report_created_at"`
            RequestDetails map[string]interface{}  `json:"request_details"`
            DisplayName string  `json:"display_name"`
            ReportType string  `json:"report_type"`
         
    }
    
    // JioCodeUpsertDataSet used by Order
    type JioCodeUpsertDataSet struct {

        
            ArticleID string  `json:"article_id"`
            ItemID string  `json:"item_id"`
            JioCode string  `json:"jio_code"`
            CompanyID string  `json:"company_id"`
         
    }
    
    // JioCodeUpsertPayload used by Order
    type JioCodeUpsertPayload struct {

        
            Data []JioCodeUpsertDataSet  `json:"data"`
         
    }
    
    // NestedErrorSchemaDataSet used by Order
    type NestedErrorSchemaDataSet struct {

        
            Type string  `json:"type"`
            Message string  `json:"message"`
            Value string  `json:"value"`
         
    }
    
    // JioCodeUpsertResponse used by Order
    type JioCodeUpsertResponse struct {

        
            TraceID string  `json:"trace_id"`
            Error []NestedErrorSchemaDataSet  `json:"error"`
            Success bool  `json:"success"`
            Data []map[string]interface{}  `json:"data"`
            Identifier string  `json:"identifier"`
         
    }
    
    // BulkInvoicingResponse used by Order
    type BulkInvoicingResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // BulkInvoiceLabelResponse used by Order
    type BulkInvoiceLabelResponse struct {

        
            Label map[string]interface{}  `json:"label"`
            BatchID string  `json:"batch_id"`
            DoInvoiceLabelGenerated bool  `json:"do_invoice_label_generated"`
            StoreID string  `json:"store_id"`
            Invoice map[string]interface{}  `json:"invoice"`
            InvoiceStatus string  `json:"invoice_status"`
            Data map[string]interface{}  `json:"data"`
            CompanyID string  `json:"company_id"`
            StoreName string  `json:"store_name"`
            StoreCode string  `json:"store_code"`
         
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

        
            Namespace string  `json:"namespace"`
            FileName string  `json:"file_name"`
            FilePath string  `json:"file_path"`
            Upload FileUploadResponse  `json:"upload"`
            Cdn URL  `json:"cdn"`
            Tags []string  `json:"tags"`
            Operation string  `json:"operation"`
            ContentType string  `json:"content_type"`
            Size float64  `json:"size"`
            Method string  `json:"method"`
         
    }
    
    // BulkListingPage used by Order
    type BulkListingPage struct {

        
            HasNext bool  `json:"has_next"`
            HasPrevious bool  `json:"has_previous"`
            Type string  `json:"type"`
            Size float64  `json:"size"`
            Current float64  `json:"current"`
            Total float64  `json:"total"`
         
    }
    
    // bulkListingData used by Order
    type bulkListingData struct {

        
            StoreID float64  `json:"store_id"`
            UserID string  `json:"user_id"`
            SuccessfulShipments []map[string]interface{}  `json:"successful_shipments"`
            ExcelURL string  `json:"excel_url"`
            UploadedOn string  `json:"uploaded_on"`
            Failed float64  `json:"failed"`
            FileName string  `json:"file_name"`
            ProcessingShipments []string  `json:"processing_shipments"`
            ID string  `json:"id"`
            CompanyID float64  `json:"company_id"`
            StoreName string  `json:"store_name"`
            Total float64  `json:"total"`
            StoreCode string  `json:"store_code"`
            Status string  `json:"status"`
            BatchID string  `json:"batch_id"`
            UserName string  `json:"user_name"`
            Successful float64  `json:"successful"`
            FailedShipments []map[string]interface{}  `json:"failed_shipments"`
            Processing float64  `json:"processing"`
         
    }
    
    // BulkListingResponse used by Order
    type BulkListingResponse struct {

        
            Error string  `json:"error"`
            Page BulkListingPage  `json:"page"`
            Success bool  `json:"success"`
            Data []bulkListingData  `json:"data"`
         
    }
    
    // ManifestPage used by Order
    type ManifestPage struct {

        
            HasNext bool  `json:"has_next"`
            HasPrevious bool  `json:"has_previous"`
            Type string  `json:"type"`
            Size string  `json:"size"`
            Current string  `json:"current"`
            Total float64  `json:"total"`
         
    }
    
    // DateRange used by Order
    type DateRange struct {

        
            ToDate string  `json:"to_date"`
            FromDate string  `json:"from_date"`
         
    }
    
    // ManifestFilter used by Order
    type ManifestFilter struct {

        
            Lane string  `json:"lane"`
            SalesChannelName string  `json:"sales_channel_name"`
            DpIds string  `json:"dp_ids"`
            SalesChannels string  `json:"sales_channels"`
            DateRange DateRange  `json:"date_range"`
            DpName string  `json:"dp_name"`
            StoreName string  `json:"store_name"`
            Stores string  `json:"stores"`
         
    }
    
    // GeneratedManifestItem used by Order
    type GeneratedManifestItem struct {

        
            Status string  `json:"status"`
            Filters ManifestFilter  `json:"filters"`
            CreatedBy string  `json:"created_by"`
            IsActive bool  `json:"is_active"`
            ManifestID string  `json:"manifest_id"`
            CompanyID float64  `json:"company_id"`
            CreatedAt string  `json:"created_at"`
         
    }
    
    // GeneratedManifestResponse used by Order
    type GeneratedManifestResponse struct {

        
            Page ManifestPage  `json:"page"`
            Items []GeneratedManifestItem  `json:"items"`
         
    }
    
    // ManifestDetailTotalShipmentPricesCount used by Order
    type ManifestDetailTotalShipmentPricesCount struct {

        
            TotalPrice float64  `json:"total_price"`
            ShipmentCount float64  `json:"shipment_count"`
         
    }
    
    // ManifestDetailMeta used by Order
    type ManifestDetailMeta struct {

        
            TotalShipmentPricesCount ManifestDetailTotalShipmentPricesCount  `json:"total_shipment_prices_count"`
            Filters ManifestFilter  `json:"filters"`
         
    }
    
    // ManifestDetail used by Order
    type ManifestDetail struct {

        
            Status string  `json:"status"`
            Filters ManifestFilter  `json:"filters"`
            UID float64  `json:"uid"`
            CreatedBy string  `json:"created_by"`
            UserID float64  `json:"user_id"`
            ID float64  `json:"id"`
            IsActive bool  `json:"is_active"`
            ManifestID string  `json:"manifest_id"`
            CompanyID float64  `json:"company_id"`
            CreatedAt string  `json:"created_at"`
            Meta ManifestDetailMeta  `json:"meta"`
         
    }
    
    // ManifestDetailItem used by Order
    type ManifestDetailItem struct {

        
            ShipmentID string  `json:"shipment_id"`
            InvoiceID string  `json:"invoice_id"`
            Awb string  `json:"awb"`
            OrderID string  `json:"order_id"`
            ItemQty float64  `json:"item_qty"`
         
    }
    
    // ManifestDetailResponse used by Order
    type ManifestDetailResponse struct {

        
            AdditionalShipmentCount float64  `json:"additional_shipment_count"`
            Page ManifestPage  `json:"page"`
            ManifestDetails []ManifestDetail  `json:"manifest_details"`
            Items []ManifestDetailItem  `json:"items"`
         
    }
    
    // QuestionSet used by Order
    type QuestionSet struct {

        
            DisplayName string  `json:"display_name"`
            ID float64  `json:"id"`
         
    }
    
    // Reason used by Order
    type Reason struct {

        
            DisplayName string  `json:"display_name"`
            ID float64  `json:"id"`
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

        
            Status bool  `json:"status"`
            Message string  `json:"message"`
         
    }
    
    // BulkActionDetailsDataField used by Order
    type BulkActionDetailsDataField struct {

        
            ProcessingShipmentsCount float64  `json:"processing_shipments_count"`
            BatchID string  `json:"batch_id"`
            SuccessfulShipmentsCount float64  `json:"successful_shipments_count"`
            FailedShipmentsCount float64  `json:"failed_shipments_count"`
            CompanyID string  `json:"company_id"`
            TotalShipmentsCount float64  `json:"total_shipments_count"`
            SuccessfulShipmentIds []string  `json:"successful_shipment_ids"`
         
    }
    
    // BulkActionDetailsResponse used by Order
    type BulkActionDetailsResponse struct {

        
            Status bool  `json:"status"`
            FailedRecords []string  `json:"failed_records"`
            Error []string  `json:"error"`
            UserID string  `json:"user_id"`
            UploadedOn string  `json:"uploaded_on"`
            Success string  `json:"success"`
            Data []BulkActionDetailsDataField  `json:"data"`
            UploadedBy string  `json:"uploaded_by"`
            Message string  `json:"message"`
         
    }
    
    // Document used by Order
    type Document struct {

        
            DsType string  `json:"ds_type"`
            Value string  `json:"value"`
            URL string  `json:"url"`
            LegalName string  `json:"legal_name"`
            Verified bool  `json:"verified"`
         
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
    
    // StoreEinvoice used by Order
    type StoreEinvoice struct {

        
            Enabled bool  `json:"enabled"`
            User string  `json:"user"`
            Username string  `json:"username"`
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
    
    // StoreMeta used by Order
    type StoreMeta struct {

        
            Timing []map[string]interface{}  `json:"timing"`
            Documents StoreDocuments  `json:"documents"`
            EwaybillPortalDetails map[string]interface{}  `json:"ewaybill_portal_details"`
            AdditionalContactDetails map[string]interface{}  `json:"additional_contact_details"`
            NotificationEmails []string  `json:"notification_emails"`
            Stage string  `json:"stage"`
            ProductReturnConfig map[string]interface{}  `json:"product_return_config"`
            GstNumber string  `json:"gst_number"`
            DisplayName string  `json:"display_name"`
            EinvoicePortalDetails EInvoicePortalDetails  `json:"einvoice_portal_details"`
            GstCredentials StoreGstCredentials  `json:"gst_credentials"`
         
    }
    
    // StoreAddress used by Order
    type StoreAddress struct {

        
            UpdatedAt string  `json:"updated_at"`
            AddressType string  `json:"address_type"`
            Longitude float64  `json:"longitude"`
            Address1 string  `json:"address1"`
            Country string  `json:"country"`
            Email string  `json:"email"`
            ContactPerson string  `json:"contact_person"`
            CreatedAt string  `json:"created_at"`
            Phone string  `json:"phone"`
            Pincode float64  `json:"pincode"`
            Landmark string  `json:"landmark"`
            Area string  `json:"area"`
            Address2 string  `json:"address2"`
            State string  `json:"state"`
            Version string  `json:"version"`
            CountryCode string  `json:"country_code"`
            AddressCategory string  `json:"address_category"`
            Latitude float64  `json:"latitude"`
            City string  `json:"city"`
         
    }
    
    // Store used by Order
    type Store struct {

        
            StoreEmail string  `json:"store_email"`
            UpdatedAt string  `json:"updated_at"`
            IsArchived bool  `json:"is_archived"`
            FulfillmentChannel string  `json:"fulfillment_channel"`
            Code string  `json:"code"`
            Address1 string  `json:"address1"`
            Longitude float64  `json:"longitude"`
            ParentStoreID float64  `json:"parent_store_id"`
            MallName string  `json:"mall_name"`
            BrandID interface{}  `json:"brand_id"`
            Country string  `json:"country"`
            BrandStoreTags []string  `json:"brand_store_tags"`
            IsEnabledForRecon bool  `json:"is_enabled_for_recon"`
            ContactPerson string  `json:"contact_person"`
            PackagingMaterialCount float64  `json:"packaging_material_count"`
            LocationType string  `json:"location_type"`
            CreatedAt string  `json:"created_at"`
            Meta StoreMeta  `json:"meta"`
            Phone float64  `json:"phone"`
            MallArea string  `json:"mall_area"`
            VatNo string  `json:"vat_no"`
            CompanyID float64  `json:"company_id"`
            Pincode string  `json:"pincode"`
            LoginUsername string  `json:"login_username"`
            SID string  `json:"s_id"`
            Name string  `json:"name"`
            Address2 string  `json:"address2"`
            AlohomoraUserID float64  `json:"alohomora_user_id"`
            StoreAddressJson StoreAddress  `json:"store_address_json"`
            State string  `json:"state"`
            IsActive bool  `json:"is_active"`
            OrderIntegrationID string  `json:"order_integration_id"`
            Latitude float64  `json:"latitude"`
            StoreActiveFrom string  `json:"store_active_from"`
            City string  `json:"city"`
         
    }
    
    // Dimensions used by Order
    type Dimensions struct {

        
            Length float64  `json:"length"`
            Unit string  `json:"unit"`
            Height float64  `json:"height"`
            IsDefault bool  `json:"is_default"`
            Width float64  `json:"width"`
         
    }
    
    // ReturnConfig used by Order
    type ReturnConfig struct {

        
            Time float64  `json:"time"`
            Unit string  `json:"unit"`
            Returnable bool  `json:"returnable"`
         
    }
    
    // Weight used by Order
    type Weight struct {

        
            IsDefault bool  `json:"is_default"`
            Shipping float64  `json:"shipping"`
            Unit string  `json:"unit"`
         
    }
    
    // Article used by Order
    type Article struct {

        
            IsSet bool  `json:"is_set"`
            Dimensions Dimensions  `json:"dimensions"`
            ChildDetails map[string]interface{}  `json:"child_details"`
            Identifiers Identifier  `json:"identifiers"`
            UID string  `json:"uid"`
            ReturnConfig ReturnConfig  `json:"return_config"`
            RawMeta interface{}  `json:"raw_meta"`
            Weight Weight  `json:"weight"`
            ASet map[string]interface{}  `json:"a_set"`
            Size string  `json:"size"`
            EspModified interface{}  `json:"esp_modified"`
            Code string  `json:"code"`
            SellerIdentifier string  `json:"seller_identifier"`
            ID string  `json:"_id"`
         
    }
    
    // AffiliateMeta used by Order
    type AffiliateMeta struct {

        
            BoxType string  `json:"box_type"`
            ChannelOrderID string  `json:"channel_order_id"`
            IsPriority bool  `json:"is_priority"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            SizeLevelTotalQty float64  `json:"size_level_total_qty"`
            EmployeeDiscount float64  `json:"employee_discount"`
            DueDate string  `json:"due_date"`
            OrderItemID string  `json:"order_item_id"`
            ChannelShipmentID string  `json:"channel_shipment_id"`
            Quantity float64  `json:"quantity"`
            CouponCode string  `json:"coupon_code"`
         
    }
    
    // LockData used by Order
    type LockData struct {

        
            Locked bool  `json:"locked"`
            LockMessage string  `json:"lock_message"`
            Mto bool  `json:"mto"`
         
    }
    
    // ShipmentTimeStamp used by Order
    type ShipmentTimeStamp struct {

        
            TMin string  `json:"t_min"`
            TMax string  `json:"t_max"`
         
    }
    
    // EInvoice used by Order
    type EInvoice struct {

        
            ErrorMessage string  `json:"error_message"`
            SignedQrCode string  `json:"signed_qr_code"`
            ErrorCode string  `json:"error_code"`
            AcknowledgeNo float64  `json:"acknowledge_no"`
            SignedInvoice string  `json:"signed_invoice"`
            Irn string  `json:"irn"`
            AcknowledgeDate string  `json:"acknowledge_date"`
         
    }
    
    // EinvoiceInfo used by Order
    type EinvoiceInfo struct {

        
            CreditNote EInvoice  `json:"credit_note"`
            Invoice EInvoice  `json:"invoice"`
         
    }
    
    // Formatted used by Order
    type Formatted struct {

        
            FMax string  `json:"f_max"`
            FMin string  `json:"f_min"`
         
    }
    
    // BuyerDetails used by Order
    type BuyerDetails struct {

        
            AjioSiteID string  `json:"ajio_site_id"`
            State string  `json:"state"`
            Address string  `json:"address"`
            Name string  `json:"name"`
            Pincode float64  `json:"pincode"`
            Gstin string  `json:"gstin"`
            City string  `json:"city"`
         
    }
    
    // DebugInfo used by Order
    type DebugInfo struct {

        
            StormbreakerUUID string  `json:"stormbreaker_uuid"`
         
    }
    
    // ShipmentMeta used by Order
    type ShipmentMeta struct {

        
            BoxType string  `json:"box_type"`
            ForwardAffiliateOrderID string  `json:"forward_affiliate_order_id"`
            StoreInvoiceUpdatedDate string  `json:"store_invoice_updated_date"`
            B2cBuyerDetails map[string]interface{}  `json:"b2c_buyer_details"`
            External map[string]interface{}  `json:"external"`
            AutoTriggerDpAssignmentAcf bool  `json:"auto_trigger_dp_assignment_acf"`
            LockData LockData  `json:"lock_data"`
            DpName string  `json:"dp_name"`
            ForwardAffiliateShipmentID string  `json:"forward_affiliate_shipment_id"`
            Timestamp ShipmentTimeStamp  `json:"timestamp"`
            AwbNumber string  `json:"awb_number"`
            ReturnStoreNode float64  `json:"return_store_node"`
            AssignDpFromSb bool  `json:"assign_dp_from_sb"`
            ShipmentWeight float64  `json:"shipment_weight"`
            MarketplaceStoreID string  `json:"marketplace_store_id"`
            SameStoreAvailable bool  `json:"same_store_available"`
            EinvoiceInfo EinvoiceInfo  `json:"einvoice_info"`
            Weight float64  `json:"weight"`
            FulfilmentPriorityText string  `json:"fulfilment_priority_text"`
            DpOptions map[string]interface{}  `json:"dp_options"`
            DpID string  `json:"dp_id"`
            ReturnAffiliateShipmentID string  `json:"return_affiliate_shipment_id"`
            PackagingName string  `json:"packaging_name"`
            DpSortKey string  `json:"dp_sort_key"`
            DueDate string  `json:"due_date"`
            Formatted Formatted  `json:"formatted"`
            ReturnDetails map[string]interface{}  `json:"return_details"`
            OrderType string  `json:"order_type"`
            B2bBuyerDetails BuyerDetails  `json:"b2b_buyer_details"`
            ShipmentVolumetricWeight float64  `json:"shipment_volumetric_weight"`
            PoNumber string  `json:"po_number"`
            ReturnAffiliateOrderID string  `json:"return_affiliate_order_id"`
            DebugInfo DebugInfo  `json:"debug_info"`
            EwaybillInfo map[string]interface{}  `json:"ewaybill_info"`
            BagWeight map[string]interface{}  `json:"bag_weight"`
            ReturnAwbNumber string  `json:"return_awb_number"`
         
    }
    
    // PDFLinks used by Order
    type PDFLinks struct {

        
            Label string  `json:"label"`
            LabelA4 string  `json:"label_a4"`
            InvoiceA4 string  `json:"invoice_a4"`
            InvoiceType string  `json:"invoice_type"`
            LabelA6 string  `json:"label_a6"`
            Invoice string  `json:"invoice"`
            PoInvoice string  `json:"po_invoice"`
            InvoiceA6 string  `json:"invoice_a6"`
            LabelPos string  `json:"label_pos"`
            InvoicePos string  `json:"invoice_pos"`
            CreditNoteURL string  `json:"credit_note_url"`
            B2b string  `json:"b2b"`
            LabelType string  `json:"label_type"`
         
    }
    
    // AffiliateDetails used by Order
    type AffiliateDetails struct {

        
            AdID string  `json:"ad_id"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            AffiliateMeta AffiliateMeta  `json:"affiliate_meta"`
            ShipmentMeta ShipmentMeta  `json:"shipment_meta"`
            AffiliateStoreID string  `json:"affiliate_store_id"`
            PdfLinks PDFLinks  `json:"pdf_links"`
            AffiliateID string  `json:"affiliate_id"`
            CompanyAffiliateTag string  `json:"company_affiliate_tag"`
         
    }
    
    // Dates used by Order
    type Dates struct {

        
            OrderCreated string  `json:"order_created"`
            DeliveryDate interface{}  `json:"delivery_date"`
         
    }
    
    // Brand used by Order
    type Brand struct {

        
            BrandName string  `json:"brand_name"`
            ModifiedOn float64  `json:"modified_on"`
            ScriptLastRan string  `json:"script_last_ran"`
            Company string  `json:"company"`
            PickupLocation string  `json:"pickup_location"`
            CreditNoteExpiryDays float64  `json:"credit_note_expiry_days"`
            CreditNoteAllowed bool  `json:"credit_note_allowed"`
            CreatedOn float64  `json:"created_on"`
            Logo string  `json:"logo"`
            IsVirtualInvoice bool  `json:"is_virtual_invoice"`
            StartDate string  `json:"start_date"`
            BrandID float64  `json:"brand_id"`
            InvoicePrefix string  `json:"invoice_prefix"`
         
    }
    
    // BagGSTDetails used by Order
    type BagGSTDetails struct {

        
            HsnCode string  `json:"hsn_code"`
            CgstGstFee string  `json:"cgst_gst_fee"`
            TaxCollectedAtSource float64  `json:"tax_collected_at_source"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            IgstTaxPercentage float64  `json:"igst_tax_percentage"`
            ValueOfGood float64  `json:"value_of_good"`
            GstinCode string  `json:"gstin_code"`
            IgstGstFee string  `json:"igst_gst_fee"`
            GstFee float64  `json:"gst_fee"`
            SgstGstFee string  `json:"sgst_gst_fee"`
            GstTag string  `json:"gst_tag"`
            HsnCodeID string  `json:"hsn_code_id"`
            CgstTaxPercentage float64  `json:"cgst_tax_percentage"`
            SgstTaxPercentage float64  `json:"sgst_tax_percentage"`
            IsDefaultHsnCode bool  `json:"is_default_hsn_code"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
         
    }
    
    // B2BPODetails used by Order
    type B2BPODetails struct {

        
            DockerNumber string  `json:"docker_number"`
            ItemBasePrice float64  `json:"item_base_price"`
            PoTaxAmount float64  `json:"po_tax_amount"`
            PartialCanRet bool  `json:"partial_can_ret"`
            PoLineAmount float64  `json:"po_line_amount"`
            TotalGstPercentage float64  `json:"total_gst_percentage"`
         
    }
    
    // BagMeta used by Order
    type BagMeta struct {

        
            B2bPoDetails B2BPODetails  `json:"b2b_po_details"`
         
    }
    
    // Attributes used by Order
    type Attributes struct {

        
            BrandName string  `json:"brand_name"`
            MarketerName string  `json:"marketer_name"`
            Essential string  `json:"essential"`
            Gender []string  `json:"gender"`
            PrimaryColor string  `json:"primary_color"`
            MarketerAddress string  `json:"marketer_address"`
            Name string  `json:"name"`
            PrimaryColorHex string  `json:"primary_color_hex"`
            PrimaryMaterial string  `json:"primary_material"`
         
    }
    
    // Item used by Order
    type Item struct {

        
            CanCancel bool  `json:"can_cancel"`
            WebstoreProductURL string  `json:"webstore_product_url"`
            L3Category float64  `json:"l3_category"`
            Attributes Attributes  `json:"attributes"`
            LastUpdatedAt string  `json:"last_updated_at"`
            Code string  `json:"code"`
            ItemID float64  `json:"item_id"`
            SlugKey string  `json:"slug_key"`
            L3CategoryName string  `json:"l3_category_name"`
            BrandID float64  `json:"brand_id"`
            L2Category []string  `json:"l2_category"`
            Color string  `json:"color"`
            Brand string  `json:"brand"`
            DepartmentID float64  `json:"department_id"`
            BranchURL string  `json:"branch_url"`
            L2CategoryID float64  `json:"l2_category_id"`
            Size string  `json:"size"`
            Meta map[string]interface{}  `json:"meta"`
            L1Category []string  `json:"l1_category"`
            Name string  `json:"name"`
            CanReturn bool  `json:"can_return"`
            Gender string  `json:"gender"`
            L1CategoryID float64  `json:"l1_category_id"`
            Image []string  `json:"image"`
         
    }
    
    // AffiliateBagDetails used by Order
    type AffiliateBagDetails struct {

        
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
            EmployeeDiscount float64  `json:"employee_discount"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            AffiliateMeta AffiliateMeta  `json:"affiliate_meta"`
         
    }
    
    // BagReturnableCancelableStatus used by Order
    type BagReturnableCancelableStatus struct {

        
            IsReturnable bool  `json:"is_returnable"`
            IsActive bool  `json:"is_active"`
            IsCustomerReturnAllowed bool  `json:"is_customer_return_allowed"`
            CanBeCancelled bool  `json:"can_be_cancelled"`
            EnableTracking bool  `json:"enable_tracking"`
         
    }
    
    // ArticleDetails used by Order
    type ArticleDetails struct {

        
            Status map[string]interface{}  `json:"status"`
         
    }
    
    // BagDetailsPlatformResponse used by Order
    type BagDetailsPlatformResponse struct {

        
            ParentPromoBags map[string]interface{}  `json:"parent_promo_bags"`
            JourneyType string  `json:"journey_type"`
            RestorePromos map[string]interface{}  `json:"restore_promos"`
            BagStatusHistory BagStatusHistory  `json:"bag_status_history"`
            Reasons []map[string]interface{}  `json:"reasons"`
            OrderingStore Store  `json:"ordering_store"`
            Article Article  `json:"article"`
            AffiliateDetails AffiliateDetails  `json:"affiliate_details"`
            LineNumber float64  `json:"line_number"`
            FinancialBreakup []FinancialBreakup  `json:"financial_breakup"`
            CurrentOperationalStatus BagStatusHistory  `json:"current_operational_status"`
            Dates Dates  `json:"dates"`
            BType string  `json:"b_type"`
            Identifier string  `json:"identifier"`
            Brand Brand  `json:"brand"`
            GstDetails BagGSTDetails  `json:"gst_details"`
            OriginalBagList []float64  `json:"original_bag_list"`
            NoOfBagsOrder float64  `json:"no_of_bags_order"`
            Tags []string  `json:"tags"`
            SellerIdentifier string  `json:"seller_identifier"`
            Meta BagMeta  `json:"meta"`
            RestoreCoupon bool  `json:"restore_coupon"`
            CurrentStatus BagStatusHistory  `json:"current_status"`
            DisplayName string  `json:"display_name"`
            Prices Prices  `json:"prices"`
            Item Item  `json:"item"`
            AffiliateBagDetails AffiliateBagDetails  `json:"affiliate_bag_details"`
            OperationalStatus string  `json:"operational_status"`
            Quantity float64  `json:"quantity"`
            Status BagReturnableCancelableStatus  `json:"status"`
            ShipmentID string  `json:"shipment_id"`
            ArticleDetails ArticleDetails  `json:"article_details"`
            BagUpdateTime float64  `json:"bag_update_time"`
            AppliedPromos []map[string]interface{}  `json:"applied_promos"`
            BID float64  `json:"b_id"`
            EntityType string  `json:"entity_type"`
            BagStatus []BagStatusHistory  `json:"bag_status"`
            QcRequired interface{}  `json:"qc_required"`
            OrderIntegrationID string  `json:"order_integration_id"`
         
    }
    
    // ErrorResponse used by Order
    type ErrorResponse struct {

        
            Error string  `json:"error"`
            Message string  `json:"message"`
         
    }
    
    // Page1 used by Order
    type Page1 struct {

        
            HasNext bool  `json:"has_next"`
            ItemTotal float64  `json:"item_total"`
            PageType string  `json:"page_type"`
            Size float64  `json:"size"`
            Current float64  `json:"current"`
         
    }
    
    // GetBagsPlatformResponse used by Order
    type GetBagsPlatformResponse struct {

        
            Page Page1  `json:"page"`
            Items []BagDetailsPlatformResponse  `json:"items"`
         
    }
    
    // InvalidateShipmentCachePayload used by Order
    type InvalidateShipmentCachePayload struct {

        
            ShipmentIds []string  `json:"shipment_ids"`
         
    }
    
    // InvalidateShipmentCacheNestedResponse used by Order
    type InvalidateShipmentCacheNestedResponse struct {

        
            Message string  `json:"message"`
            Status float64  `json:"status"`
            Error string  `json:"error"`
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
            ErrorTrace string  `json:"error_trace"`
         
    }
    
    // StoreReassign used by Order
    type StoreReassign struct {

        
            AffiliateOrderID string  `json:"affiliate_order_id"`
            SetID string  `json:"set_id"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
            ReasonIds []float64  `json:"reason_ids"`
            MongoArticleID string  `json:"mongo_article_id"`
            BagID float64  `json:"bag_id"`
            ItemID string  `json:"item_id"`
            AffiliateID string  `json:"affiliate_id"`
            StoreID float64  `json:"store_id"`
            FyndOrderID string  `json:"fynd_order_id"`
         
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
            ReasonText string  `json:"reason_text"`
            AffiliateID string  `json:"affiliate_id"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            ID string  `json:"id"`
         
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

        
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            Status string  `json:"status"`
            IsBagLocked bool  `json:"is_bag_locked"`
            OriginalFilter OriginalFilter  `json:"original_filter"`
            LockStatus bool  `json:"lock_status"`
            ShipmentID string  `json:"shipment_id"`
            AffiliateID string  `json:"affiliate_id"`
            Bags []Bags  `json:"bags"`
            IsShipmentLocked bool  `json:"is_shipment_locked"`
         
    }
    
    // UpdateShipmentLockResponse used by Order
    type UpdateShipmentLockResponse struct {

        
            Message string  `json:"message"`
            CheckResponse []CheckResponse  `json:"check_response"`
            Success bool  `json:"success"`
         
    }
    
    // AnnouncementResponse used by Order
    type AnnouncementResponse struct {

        
            CompanyID float64  `json:"company_id"`
            PlatformID string  `json:"platform_id"`
            CreatedAt string  `json:"created_at"`
            Description string  `json:"description"`
            FromDatetime string  `json:"from_datetime"`
            Title string  `json:"title"`
            ToDatetime string  `json:"to_datetime"`
            PlatformName string  `json:"platform_name"`
            LogoURL string  `json:"logo_url"`
            ID float64  `json:"id"`
         
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
    
    // Products used by Order
    type Products struct {

        
            Identifier string  `json:"identifier"`
            LineNumber float64  `json:"line_number"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // ProductsReasonsData used by Order
    type ProductsReasonsData struct {

        
            ReasonID float64  `json:"reason_id"`
            ReasonText string  `json:"reason_text"`
         
    }
    
    // ProductsReasonsFilters used by Order
    type ProductsReasonsFilters struct {

        
            Identifier string  `json:"identifier"`
            LineNumber float64  `json:"line_number"`
            Quantity float64  `json:"quantity"`
         
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

        
            Identifier string  `json:"identifier"`
            LineNumber float64  `json:"line_number"`
         
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
    
    // ShipmentsRequest used by Order
    type ShipmentsRequest struct {

        
            Identifier string  `json:"identifier"`
            Products []Products  `json:"products"`
            Reasons ReasonsData  `json:"reasons"`
            DataUpdates DataUpdates  `json:"data_updates"`
         
    }
    
    // StatuesRequest used by Order
    type StatuesRequest struct {

        
            Status string  `json:"status"`
            ExcludeBagsNextState string  `json:"exclude_bags_next_state"`
            Shipments []ShipmentsRequest  `json:"shipments"`
         
    }
    
    // UpdateShipmentStatusRequest used by Order
    type UpdateShipmentStatusRequest struct {

        
            UnlockBeforeTransition bool  `json:"unlock_before_transition"`
            LockAfterTransition bool  `json:"lock_after_transition"`
            ForceTransition bool  `json:"force_transition"`
            Task bool  `json:"task"`
            Statuses []StatuesRequest  `json:"statuses"`
         
    }
    
    // ShipmentsResponse used by Order
    type ShipmentsResponse struct {

        
            Status float64  `json:"status"`
            Exception string  `json:"exception"`
            Code string  `json:"code"`
            Message string  `json:"message"`
            FinalState map[string]interface{}  `json:"final_state"`
            Meta map[string]interface{}  `json:"meta"`
            Identifier string  `json:"identifier"`
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

        
            Name string  `json:"name"`
            CreatedAt string  `json:"created_at"`
            Description string  `json:"description"`
            UpdatedAt string  `json:"updated_at"`
            Meta []AffiliateAppConfigMeta  `json:"meta"`
            Owner string  `json:"owner"`
            Secret string  `json:"secret"`
            Token string  `json:"token"`
            ID string  `json:"id"`
         
    }
    
    // AffiliateInventoryOrderConfig used by Order
    type AffiliateInventoryOrderConfig struct {

        
            ForceReassignment bool  `json:"force_reassignment"`
         
    }
    
    // AffiliateInventoryArticleAssignmentConfig used by Order
    type AffiliateInventoryArticleAssignmentConfig struct {

        
            PostOrderReassignment bool  `json:"post_order_reassignment"`
         
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
    
    // AffiliateInventoryStoreConfig used by Order
    type AffiliateInventoryStoreConfig struct {

        
            Store map[string]interface{}  `json:"store"`
         
    }
    
    // AffiliateInventoryConfig used by Order
    type AffiliateInventoryConfig struct {

        
            Order AffiliateInventoryOrderConfig  `json:"order"`
            ArticleAssignment AffiliateInventoryArticleAssignmentConfig  `json:"article_assignment"`
            Logistics AffiliateInventoryLogisticsConfig  `json:"logistics"`
            Payment AffiliateInventoryPaymentConfig  `json:"payment"`
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
            Token string  `json:"token"`
            ID string  `json:"id"`
         
    }
    
    // OrderConfig used by Order
    type OrderConfig struct {

        
            StoreLookup string  `json:"store_lookup"`
            AffiliateStoreIDMapping []AffiliateStoreIdMapping  `json:"affiliate_store_id_mapping"`
            CreateUser bool  `json:"create_user"`
            Affiliate Affiliate  `json:"affiliate"`
            ArticleLookup string  `json:"article_lookup"`
            BagEndState string  `json:"bag_end_state"`
         
    }
    
    // OrderUser used by Order
    type OrderUser struct {

        
            City string  `json:"city"`
            FirstName string  `json:"first_name"`
            State string  `json:"state"`
            Mobile float64  `json:"mobile"`
            Address2 string  `json:"address2"`
            Address1 string  `json:"address1"`
            Country string  `json:"country"`
            Phone float64  `json:"phone"`
            Email string  `json:"email"`
            LastName string  `json:"last_name"`
            Pincode string  `json:"pincode"`
         
    }
    
    // UserData used by Order
    type UserData struct {

        
            BillingUser OrderUser  `json:"billing_user"`
            ShippingUser OrderUser  `json:"shipping_user"`
         
    }
    
    // MarketPlacePdf used by Order
    type MarketPlacePdf struct {

        
            Label string  `json:"label"`
            Invoice string  `json:"invoice"`
         
    }
    
    // AffiliateBag used by Order
    type AffiliateBag struct {

        
            PdfLinks MarketPlacePdf  `json:"pdf_links"`
            AffiliateMeta map[string]interface{}  `json:"affiliate_meta"`
            ItemSize string  `json:"item_size"`
            AvlQty float64  `json:"avl_qty"`
            PriceEffective float64  `json:"price_effective"`
            HsnCodeID string  `json:"hsn_code_id"`
            AmountPaid float64  `json:"amount_paid"`
            Sku string  `json:"sku"`
            DeliveryCharge float64  `json:"delivery_charge"`
            UnitPrice float64  `json:"unit_price"`
            AffiliateStoreID string  `json:"affiliate_store_id"`
            CompanyID float64  `json:"company_id"`
            ItemID float64  `json:"item_id"`
            ModifiedOn string  `json:"modified_on"`
            StoreID float64  `json:"store_id"`
            Discount float64  `json:"discount"`
            TransferPrice float64  `json:"transfer_price"`
            FyndStoreID string  `json:"fynd_store_id"`
            Quantity float64  `json:"quantity"`
            ID string  `json:"_id"`
            Identifier map[string]interface{}  `json:"identifier"`
            PriceMarked float64  `json:"price_marked"`
            SellerIdentifier string  `json:"seller_identifier"`
         
    }
    
    // OrderPriority used by Order
    type OrderPriority struct {

        
            FulfilmentPriority float64  `json:"fulfilment_priority"`
            AffiliatePriorityCode string  `json:"affiliate_priority_code"`
            FulfilmentPriorityText string  `json:"fulfilment_priority_text"`
         
    }
    
    // ArticleDetails1 used by Order
    type ArticleDetails1 struct {

        
            BrandID float64  `json:"brand_id"`
            Dimension map[string]interface{}  `json:"dimension"`
            Attributes map[string]interface{}  `json:"attributes"`
            Category map[string]interface{}  `json:"category"`
            Weight map[string]interface{}  `json:"weight"`
            Quantity float64  `json:"quantity"`
            ID string  `json:"_id"`
         
    }
    
    // LocationDetails used by Order
    type LocationDetails struct {

        
            FulfillmentType string  `json:"fulfillment_type"`
            Articles []ArticleDetails1  `json:"articles"`
            FulfillmentID float64  `json:"fulfillment_id"`
         
    }
    
    // ShipmentDetails used by Order
    type ShipmentDetails struct {

        
            FulfillmentID float64  `json:"fulfillment_id"`
            DpID float64  `json:"dp_id"`
            Shipments float64  `json:"shipments"`
            Meta map[string]interface{}  `json:"meta"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            BoxType string  `json:"box_type"`
            Articles []ArticleDetails1  `json:"articles"`
         
    }
    
    // ShipmentConfig used by Order
    type ShipmentConfig struct {

        
            Journey string  `json:"journey"`
            Source string  `json:"source"`
            PaymentMode string  `json:"payment_mode"`
            LocationDetails LocationDetails  `json:"location_details"`
            Action string  `json:"action"`
            Identifier string  `json:"identifier"`
            Shipment []ShipmentDetails  `json:"shipment"`
            ToPincode string  `json:"to_pincode"`
         
    }
    
    // ShipmentData used by Order
    type ShipmentData struct {

        
            ShipmentData ShipmentConfig  `json:"shipment_data"`
         
    }
    
    // OrderInfo used by Order
    type OrderInfo struct {

        
            AffiliateOrderID string  `json:"affiliate_order_id"`
            Discount float64  `json:"discount"`
            ShippingAddress OrderUser  `json:"shipping_address"`
            DeliveryCharges float64  `json:"delivery_charges"`
            Coupon string  `json:"coupon"`
            User UserData  `json:"user"`
            BillingAddress OrderUser  `json:"billing_address"`
            PaymentMode string  `json:"payment_mode"`
            Bags []AffiliateBag  `json:"bags"`
            OrderPriority OrderPriority  `json:"order_priority"`
            Payment map[string]interface{}  `json:"payment"`
            Items map[string]interface{}  `json:"items"`
            CodCharges float64  `json:"cod_charges"`
            Shipment ShipmentData  `json:"shipment"`
            OrderValue float64  `json:"order_value"`
         
    }
    
    // CreateOrderPayload used by Order
    type CreateOrderPayload struct {

        
            OrderConfig OrderConfig  `json:"order_config"`
            AffiliateID string  `json:"affiliate_id"`
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

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // ActionInfo used by Order
    type ActionInfo struct {

        
            DisplayText string  `json:"display_text"`
            Slug string  `json:"slug"`
            Description string  `json:"description"`
            ID float64  `json:"id"`
         
    }
    
    // GetActionsResponse used by Order
    type GetActionsResponse struct {

        
            Permissions ActionInfo  `json:"permissions"`
         
    }
    
    // HistoryDict used by Order
    type HistoryDict struct {

        
            Type string  `json:"type"`
            TicketID string  `json:"ticket_id"`
            BagID float64  `json:"bag_id"`
            L3Detail string  `json:"l3_detail"`
            User string  `json:"user"`
            TicketURL string  `json:"ticket_url"`
            Message string  `json:"message"`
            L1Detail string  `json:"l1_detail"`
            L2Detail string  `json:"l2_detail"`
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
    
    // SmsDataPayload used by Order
    type SmsDataPayload struct {

        
            CountryCode string  `json:"country_code"`
            AmountPaid float64  `json:"amount_paid"`
            PhoneNumber float64  `json:"phone_number"`
            CustomerName string  `json:"customer_name"`
            ShipmentID float64  `json:"shipment_id"`
            OrderID string  `json:"order_id"`
            PaymentMode string  `json:"payment_mode"`
            Message string  `json:"message"`
            BrandName string  `json:"brand_name"`
         
    }
    
    // SendSmsPayload used by Order
    type SendSmsPayload struct {

        
            BagID float64  `json:"bag_id"`
            Data SmsDataPayload  `json:"data"`
            Slug string  `json:"slug"`
         
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

        
            Status string  `json:"status"`
            ShipmentID string  `json:"shipment_id"`
            Meta Meta  `json:"meta"`
            BagList []float64  `json:"bag_list"`
            Remarks string  `json:"remarks"`
            ID float64  `json:"id"`
         
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

        
            ShipmentIds []string  `json:"shipment_ids"`
            DpID float64  `json:"dp_id"`
            OrderType string  `json:"order_type"`
            QcRequired string  `json:"qc_required"`
         
    }
    
    // ManualAssignDPToShipmentResponse used by Order
    type ManualAssignDPToShipmentResponse struct {

        
            Errors []string  `json:"errors"`
            Success string  `json:"success"`
         
    }
    
    // Tax used by Order
    type Tax struct {

        
            Name string  `json:"name"`
            Breakup []map[string]interface{}  `json:"breakup"`
            Amount map[string]interface{}  `json:"amount"`
            Rate float64  `json:"rate"`
         
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
            Quantity float64  `json:"quantity"`
            Charges []Charge  `json:"charges"`
            Meta map[string]interface{}  `json:"meta"`
            SellerIdentifier string  `json:"seller_identifier"`
         
    }
    
    // ProcessingDates used by Order
    type ProcessingDates struct {

        
            DpPickupSlot map[string]interface{}  `json:"dp_pickup_slot"`
            CustomerPickupSlot map[string]interface{}  `json:"customer_pickup_slot"`
            PackByDate string  `json:"pack_by_date"`
            DispatchAfterDate string  `json:"dispatch_after_date"`
            ConfirmByDate string  `json:"confirm_by_date"`
            DispatchByDate string  `json:"dispatch_by_date"`
         
    }
    
    // Shipment used by Order
    type Shipment struct {

        
            ExternalShipmentID float64  `json:"external_shipment_id"`
            LocationID float64  `json:"location_id"`
            LineItems []LineItem  `json:"line_items"`
            Priority float64  `json:"priority"`
            Meta map[string]interface{}  `json:"meta"`
            ProcessingDates ProcessingDates  `json:"processing_dates"`
         
    }
    
    // PaymentMethod used by Order
    type PaymentMethod struct {

        
            Name string  `json:"name"`
            Amount float64  `json:"amount"`
            RefundBy string  `json:"refund_by"`
            TransactionData map[string]interface{}  `json:"transaction_data"`
            Meta map[string]interface{}  `json:"meta"`
            CollectBy string  `json:"collect_by"`
            Mode string  `json:"mode"`
         
    }
    
    // PaymentInfo used by Order
    type PaymentInfo struct {

        
            PrimaryMode string  `json:"primary_mode"`
            PaymentMethods []PaymentMethod  `json:"payment_methods"`
         
    }
    
    // ShippingInfo used by Order
    type ShippingInfo struct {

        
            AlternateMobileNumber string  `json:"alternate_mobile_number"`
            CountryCode string  `json:"country_code"`
            HouseNo string  `json:"house_no"`
            ShippingType string  `json:"shipping_type"`
            State string  `json:"state"`
            Address2 string  `json:"address2"`
            Gender string  `json:"gender"`
            FloorNo string  `json:"floor_no"`
            CustomerCode string  `json:"customer_code"`
            StateCode string  `json:"state_code"`
            AlternateEmail string  `json:"alternate_email"`
            Slot []map[string]interface{}  `json:"slot"`
            MiddleName string  `json:"middle_name"`
            AddressType string  `json:"address_type"`
            PrimaryMobileNumber string  `json:"primary_mobile_number"`
            GeoLocation map[string]interface{}  `json:"geo_location"`
            Pincode string  `json:"pincode"`
            ExternalCustomerCode string  `json:"external_customer_code"`
            PrimaryEmail string  `json:"primary_email"`
            City string  `json:"city"`
            FirstName string  `json:"first_name"`
            Landmark string  `json:"landmark"`
            Title string  `json:"title"`
            Address1 string  `json:"address1"`
            Country string  `json:"country"`
            LastName string  `json:"last_name"`
         
    }
    
    // TaxInfo used by Order
    type TaxInfo struct {

        
            Gstin string  `json:"gstin"`
            B2bGstinNumber string  `json:"b2b_gstin_number"`
         
    }
    
    // BillingInfo used by Order
    type BillingInfo struct {

        
            AlternateMobileNumber string  `json:"alternate_mobile_number"`
            CountryCode string  `json:"country_code"`
            HouseNo string  `json:"house_no"`
            State string  `json:"state"`
            Address2 string  `json:"address2"`
            Gender string  `json:"gender"`
            FloorNo string  `json:"floor_no"`
            CustomerCode string  `json:"customer_code"`
            StateCode string  `json:"state_code"`
            AlternateEmail string  `json:"alternate_email"`
            MiddleName string  `json:"middle_name"`
            PrimaryMobileNumber string  `json:"primary_mobile_number"`
            Pincode string  `json:"pincode"`
            ExternalCustomerCode string  `json:"external_customer_code"`
            PrimaryEmail string  `json:"primary_email"`
            City string  `json:"city"`
            FirstName string  `json:"first_name"`
            Title string  `json:"title"`
            Address1 string  `json:"address1"`
            Country string  `json:"country"`
            LastName string  `json:"last_name"`
         
    }
    
    // CreateOrderAPI used by Order
    type CreateOrderAPI struct {

        
            Shipments []Shipment  `json:"shipments"`
            ExternalCreationDate string  `json:"external_creation_date"`
            PaymentInfo PaymentInfo  `json:"payment_info"`
            Charges []Charge  `json:"charges"`
            ShippingInfo ShippingInfo  `json:"shipping_info"`
            CurrencyInfo map[string]interface{}  `json:"currency_info"`
            Meta map[string]interface{}  `json:"meta"`
            TaxInfo TaxInfo  `json:"tax_info"`
            ExternalOrderID string  `json:"external_order_id"`
            BillingInfo BillingInfo  `json:"billing_info"`
         
    }
    
    // CreateOrderErrorReponse used by Order
    type CreateOrderErrorReponse struct {

        
            Status float64  `json:"status"`
            Exception string  `json:"exception"`
            Code string  `json:"code"`
            StackTrace string  `json:"stack_trace"`
            RequestID string  `json:"request_id"`
            Meta string  `json:"meta"`
            Message string  `json:"message"`
            Info interface{}  `json:"info"`
         
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
            ModeOfPayment string  `json:"mode_of_payment"`
            PaymentMethods []PaymentMethods  `json:"payment_methods"`
         
    }
    
    // CreateChannelConfig used by Order
    type CreateChannelConfig struct {

        
            DpConfiguration DpConfiguration  `json:"dp_configuration"`
            LocationReassignment bool  `json:"location_reassignment"`
            LockStates []string  `json:"lock_states"`
            ShipmentAssignment string  `json:"shipment_assignment"`
            PaymentInfo CreateChannelPaymentInfo  `json:"payment_info"`
            LogoURL map[string]interface{}  `json:"logo_url"`
         
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
            IsInserted bool  `json:"is_inserted"`
            Acknowledged bool  `json:"acknowledged"`
         
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

        
            Mobile float64  `json:"mobile"`
            OrderDetails []FyndOrderIdList  `json:"order_details"`
            StartDate string  `json:"start_date"`
            EndDate string  `json:"end_date"`
         
    }
    

    
    // DeleteResponse used by Catalog
    type DeleteResponse struct {

        
            Message string  `json:"message"`
         
    }
    
    // GetSearchWordsData used by Catalog
    type GetSearchWordsData struct {

        
            Words []string  `json:"words"`
            UID string  `json:"uid"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Result map[string]interface{}  `json:"result"`
            IsActive bool  `json:"is_active"`
            AppID string  `json:"app_id"`
         
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

        
            Words []string  `json:"words"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Result SearchKeywordResult  `json:"result"`
            IsActive bool  `json:"is_active"`
            AppID string  `json:"app_id"`
         
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
            UID string  `json:"uid"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            AppID string  `json:"app_id"`
         
    }
    
    // GetAutocompleteWordsResponse used by Catalog
    type GetAutocompleteWordsResponse struct {

        
            Page Page  `json:"page"`
            Items []GetAutocompleteWordsData  `json:"items"`
         
    }
    
    // Media used by Catalog
    type Media struct {

        
            AspectRatio string  `json:"aspect_ratio"`
            URL string  `json:"url"`
            Type string  `json:"type"`
         
    }
    
    // AutocompletePageAction used by Catalog
    type AutocompletePageAction struct {

        
            URL string  `json:"url"`
            Query map[string]interface{}  `json:"query"`
            Params map[string]interface{}  `json:"params"`
            Type string  `json:"type"`
         
    }
    
    // AutocompleteAction used by Catalog
    type AutocompleteAction struct {

        
            Page AutocompletePageAction  `json:"page"`
            Type string  `json:"type"`
         
    }
    
    // AutocompleteResult used by Catalog
    type AutocompleteResult struct {

        
            Display string  `json:"display"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Logo Media  `json:"logo"`
            Action AutocompleteAction  `json:"action"`
         
    }
    
    // CreateAutocompleteKeyword used by Catalog
    type CreateAutocompleteKeyword struct {

        
            Words []string  `json:"words"`
            Results []AutocompleteResult  `json:"results"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            IsActive bool  `json:"is_active"`
            AppID string  `json:"app_id"`
         
    }
    
    // CreateAutocompleteWordsResponse used by Catalog
    type CreateAutocompleteWordsResponse struct {

        
            Words []string  `json:"words"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            AppID string  `json:"app_id"`
            Results []map[string]interface{}  `json:"results"`
         
    }
    
    // ProductBundleItem used by Catalog
    type ProductBundleItem struct {

        
            AllowRemove bool  `json:"allow_remove"`
            AutoAddToCart bool  `json:"auto_add_to_cart"`
            MinQuantity float64  `json:"min_quantity"`
            MaxQuantity float64  `json:"max_quantity"`
            ProductUID float64  `json:"product_uid"`
            AutoSelect bool  `json:"auto_select"`
         
    }
    
    // ProductBundleRequest used by Catalog
    type ProductBundleRequest struct {

        
            CompanyID float64  `json:"company_id"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Slug string  `json:"slug"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            Products []ProductBundleItem  `json:"products"`
            Logo string  `json:"logo"`
            CreatedOn string  `json:"created_on"`
            PageVisibility []string  `json:"page_visibility"`
            ModifiedOn string  `json:"modified_on"`
            Choice string  `json:"choice"`
            IsActive bool  `json:"is_active"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Meta map[string]interface{}  `json:"meta"`
            Name string  `json:"name"`
         
    }
    
    // GetProductBundleCreateResponse used by Catalog
    type GetProductBundleCreateResponse struct {

        
            CompanyID float64  `json:"company_id"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Slug string  `json:"slug"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            Products []ProductBundleItem  `json:"products"`
            Logo string  `json:"logo"`
            CreatedOn string  `json:"created_on"`
            PageVisibility []string  `json:"page_visibility"`
            ModifiedOn string  `json:"modified_on"`
            Choice string  `json:"choice"`
            IsActive bool  `json:"is_active"`
            ID string  `json:"id"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Meta map[string]interface{}  `json:"meta"`
            Name string  `json:"name"`
         
    }
    
    // GetProductBundleListingResponse used by Catalog
    type GetProductBundleListingResponse struct {

        
            Page Page  `json:"page"`
            Items []GetProductBundleCreateResponse  `json:"items"`
         
    }
    
    // Price used by Catalog
    type Price struct {

        
            Currency string  `json:"currency"`
            MinEffective float64  `json:"min_effective"`
            MinMarked float64  `json:"min_marked"`
            MaxEffective float64  `json:"max_effective"`
            MaxMarked float64  `json:"max_marked"`
         
    }
    
    // LimitedProductData used by Catalog
    type LimitedProductData struct {

        
            Price map[string]interface{}  `json:"price"`
            ShortDescription string  `json:"short_description"`
            UID float64  `json:"uid"`
            Slug string  `json:"slug"`
            CountryOfOrigin string  `json:"country_of_origin"`
            Attributes map[string]interface{}  `json:"attributes"`
            Images []string  `json:"images"`
            Identifier map[string]interface{}  `json:"identifier"`
            ItemCode string  `json:"item_code"`
            Quantity float64  `json:"quantity"`
            Sizes []string  `json:"sizes"`
            Name string  `json:"name"`
         
    }
    
    // Size used by Catalog
    type Size struct {

        
            Quantity float64  `json:"quantity"`
            IsAvailable bool  `json:"is_available"`
            Display string  `json:"display"`
            Value string  `json:"value"`
         
    }
    
    // GetProducts used by Catalog
    type GetProducts struct {

        
            AllowRemove bool  `json:"allow_remove"`
            AutoAddToCart bool  `json:"auto_add_to_cart"`
            Price Price  `json:"price"`
            ProductDetails LimitedProductData  `json:"product_details"`
            MinQuantity float64  `json:"min_quantity"`
            Sizes []Size  `json:"sizes"`
            ProductUID float64  `json:"product_uid"`
            MaxQuantity float64  `json:"max_quantity"`
            AutoSelect bool  `json:"auto_select"`
         
    }
    
    // GetProductBundleResponse used by Catalog
    type GetProductBundleResponse struct {

        
            CompanyID float64  `json:"company_id"`
            Slug string  `json:"slug"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            Products []GetProducts  `json:"products"`
            Logo string  `json:"logo"`
            PageVisibility []string  `json:"page_visibility"`
            Choice string  `json:"choice"`
            IsActive bool  `json:"is_active"`
            Meta map[string]interface{}  `json:"meta"`
            Name string  `json:"name"`
         
    }
    
    // ProductBundleUpdateRequest used by Catalog
    type ProductBundleUpdateRequest struct {

        
            CompanyID float64  `json:"company_id"`
            Slug string  `json:"slug"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            Products []ProductBundleItem  `json:"products"`
            Logo string  `json:"logo"`
            PageVisibility []string  `json:"page_visibility"`
            ModifiedOn string  `json:"modified_on"`
            Choice string  `json:"choice"`
            IsActive bool  `json:"is_active"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Meta map[string]interface{}  `json:"meta"`
            Name string  `json:"name"`
         
    }
    
    // Guide used by Catalog
    type Guide struct {

        
            Meta Meta  `json:"meta"`
         
    }
    
    // ValidateSizeGuide used by Catalog
    type ValidateSizeGuide struct {

        
            Guide Guide  `json:"guide"`
            Active bool  `json:"active"`
            BrandID float64  `json:"brand_id"`
            CompanyID float64  `json:"company_id"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Subtitle string  `json:"subtitle"`
            Description string  `json:"description"`
            CreatedOn string  `json:"created_on"`
            Image string  `json:"image"`
            Title string  `json:"title"`
            ModifiedOn string  `json:"modified_on"`
            Tag string  `json:"tag"`
            ID string  `json:"id"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Name string  `json:"name"`
         
    }
    
    // ListSizeGuide used by Catalog
    type ListSizeGuide struct {

        
            Page map[string]interface{}  `json:"page"`
            Items []map[string]interface{}  `json:"items"`
         
    }
    
    // SizeGuideResponse used by Catalog
    type SizeGuideResponse struct {

        
            Guide map[string]interface{}  `json:"guide"`
            Active bool  `json:"active"`
            BrandID float64  `json:"brand_id"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            CompanyID float64  `json:"company_id"`
            Subtitle string  `json:"subtitle"`
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
            Title string  `json:"title"`
            Tag string  `json:"tag"`
            ID string  `json:"id"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Name string  `json:"name"`
         
    }
    
    // ApplicationItemSEO used by Catalog
    type ApplicationItemSEO struct {

        
            Title interface{}  `json:"title"`
            Description interface{}  `json:"description"`
         
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

        
            IsGift bool  `json:"is_gift"`
            AltText map[string]interface{}  `json:"alt_text"`
            IsCod bool  `json:"is_cod"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Seo ApplicationItemSEO  `json:"seo"`
            Moq ApplicationItemMOQ  `json:"moq"`
            CustomMeta []MetaFields  `json:"_custom_meta"`
         
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
            Maximum float64  `json:"maximum"`
            Minimum float64  `json:"minimum"`
         
    }
    
    // OwnerAppItemResponse used by Catalog
    type OwnerAppItemResponse struct {

        
            IsGift bool  `json:"is_gift"`
            AltText map[string]interface{}  `json:"alt_text"`
            IsCod bool  `json:"is_cod"`
            Seo SEOData  `json:"seo"`
            Moq MOQData  `json:"moq"`
         
    }
    
    // GetConfigMetadataResponse used by Catalog
    type GetConfigMetadataResponse struct {

        
            Data []map[string]interface{}  `json:"data"`
            Values []map[string]interface{}  `json:"values"`
            Condition []map[string]interface{}  `json:"condition"`
         
    }
    
    // AttributeDetailsGroup used by Catalog
    type AttributeDetailsGroup struct {

        
            DisplayType string  `json:"display_type"`
            Slug string  `json:"slug"`
            Logo string  `json:"logo"`
            Key string  `json:"key"`
            IsActive bool  `json:"is_active"`
            Priority float64  `json:"priority"`
            Unit string  `json:"unit"`
            Name string  `json:"name"`
         
    }
    
    // AppConfigurationDetail used by Catalog
    type AppConfigurationDetail struct {

        
            IsDefault bool  `json:"is_default"`
            Slug string  `json:"slug"`
            TemplateSlugs []string  `json:"template_slugs"`
            Logo string  `json:"logo"`
            Attributes []AttributeDetailsGroup  `json:"attributes"`
            IsActive bool  `json:"is_active"`
            Priority float64  `json:"priority"`
            AppID string  `json:"app_id"`
            Name string  `json:"name"`
         
    }
    
    // ConfigErrorResponse used by Catalog
    type ConfigErrorResponse struct {

        
            Message string  `json:"message"`
         
    }
    
    // PageResponseType used by Catalog
    type PageResponseType struct {

        
            HasNext bool  `json:"has_next"`
            Current float64  `json:"current"`
            Next float64  `json:"next"`
            TotalCount float64  `json:"total_count"`
         
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

        
            IsDefault bool  `json:"is_default"`
            DefaultKey string  `json:"default_key"`
            Logo string  `json:"logo"`
            Key string  `json:"key"`
            IsActive bool  `json:"is_active"`
            Priority float64  `json:"priority"`
            AppID string  `json:"app_id"`
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
    
    // GetCatalogConfigurationDetailsProduct used by Catalog
    type GetCatalogConfigurationDetailsProduct struct {

        
            Similar map[string]interface{}  `json:"similar"`
            Detail map[string]interface{}  `json:"detail"`
            Variant map[string]interface{}  `json:"variant"`
            Compare map[string]interface{}  `json:"compare"`
         
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

        
            Key string  `json:"key"`
            FilterTypes []string  `json:"filter_types"`
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
    
    // GetCatalogConfigurationMetaData used by Catalog
    type GetCatalogConfigurationMetaData struct {

        
            Product GetCatalogConfigurationDetailsProduct  `json:"product"`
            Listing MetaDataListingResponse  `json:"listing"`
         
    }
    
    // ConfigurationListingSortConfig used by Catalog
    type ConfigurationListingSortConfig struct {

        
            Logo string  `json:"logo"`
            Key string  `json:"key"`
            IsActive bool  `json:"is_active"`
            Priority float64  `json:"priority"`
            Name string  `json:"name"`
         
    }
    
    // ConfigurationListingSort used by Catalog
    type ConfigurationListingSort struct {

        
            Config []ConfigurationListingSortConfig  `json:"config"`
            DefaultKey string  `json:"default_key"`
         
    }
    
    // ConfigurationBucketPoints used by Catalog
    type ConfigurationBucketPoints struct {

        
            Start float64  `json:"start"`
            Display string  `json:"display"`
            End float64  `json:"end"`
         
    }
    
    // ConfigurationListingFilterValue used by Catalog
    type ConfigurationListingFilterValue struct {

        
            Value string  `json:"value"`
            Map map[string]interface{}  `json:"map"`
            BucketPoints []ConfigurationBucketPoints  `json:"bucket_points"`
            MapValues []map[string]interface{}  `json:"map_values"`
            Sort string  `json:"sort"`
            Condition string  `json:"condition"`
         
    }
    
    // ConfigurationListingFilterConfig used by Catalog
    type ConfigurationListingFilterConfig struct {

        
            Type string  `json:"type"`
            Logo string  `json:"logo"`
            DisplayName string  `json:"display_name"`
            Key string  `json:"key"`
            IsActive bool  `json:"is_active"`
            Priority float64  `json:"priority"`
            ValueConfig ConfigurationListingFilterValue  `json:"value_config"`
            Name string  `json:"name"`
         
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

        
            Size ProductSize  `json:"size"`
            Subtitle string  `json:"subtitle"`
            Logo string  `json:"logo"`
            Key string  `json:"key"`
            Title string  `json:"title"`
            IsActive bool  `json:"is_active"`
            Priority float64  `json:"priority"`
         
    }
    
    // ConfigurationProductSimilar used by Catalog
    type ConfigurationProductSimilar struct {

        
            Config []ConfigurationProductConfig  `json:"config"`
         
    }
    
    // ConfigurationProductVariantConfig used by Catalog
    type ConfigurationProductVariantConfig struct {

        
            Size ProductSize  `json:"size"`
            DisplayType string  `json:"display_type"`
            Logo string  `json:"logo"`
            Key string  `json:"key"`
            IsActive bool  `json:"is_active"`
            Priority float64  `json:"priority"`
            Name string  `json:"name"`
         
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
    
    // AppConfiguration used by Catalog
    type AppConfiguration struct {

        
            CreatedBy map[string]interface{}  `json:"created_by"`
            Listing ConfigurationListing  `json:"listing"`
            ConfigType string  `json:"config_type"`
            Type string  `json:"type"`
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
            ConfigID string  `json:"config_id"`
            Product ConfigurationProduct  `json:"product"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            AppID string  `json:"app_id"`
         
    }
    
    // AppCatalogConfiguration used by Catalog
    type AppCatalogConfiguration struct {

        
            CreatedBy map[string]interface{}  `json:"created_by"`
            Listing ConfigurationListing  `json:"listing"`
            ConfigType string  `json:"config_type"`
            Type string  `json:"type"`
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
            ID string  `json:"id"`
            ConfigID string  `json:"config_id"`
            Product ConfigurationProduct  `json:"product"`
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

        
            Listing GetCatalogConfigurationDetailsSchemaListing  `json:"listing"`
            ConfigType string  `json:"config_type"`
            ID string  `json:"id"`
            ConfigID string  `json:"config_id"`
            Product GetCatalogConfigurationDetailsProduct  `json:"product"`
            AppID string  `json:"app_id"`
         
    }
    
    // GetAppCatalogEntityConfiguration used by Catalog
    type GetAppCatalogEntityConfiguration struct {

        
            IsDefault bool  `json:"is_default"`
            Data EntityConfiguration  `json:"data"`
         
    }
    
    // ProductFiltersKey used by Catalog
    type ProductFiltersKey struct {

        
            Display string  `json:"display"`
            Logo string  `json:"logo"`
            Operators []string  `json:"operators"`
            Kind string  `json:"kind"`
            Name string  `json:"name"`
         
    }
    
    // ProductFiltersValue used by Catalog
    type ProductFiltersValue struct {

        
            SelectedMax float64  `json:"selected_max"`
            Max float64  `json:"max"`
            Display string  `json:"display"`
            CurrencyCode string  `json:"currency_code"`
            Value interface{}  `json:"value"`
            QueryFormat string  `json:"query_format"`
            SelectedMin float64  `json:"selected_min"`
            DisplayFormat string  `json:"display_format"`
            Min float64  `json:"min"`
            Count float64  `json:"count"`
            IsSelected bool  `json:"is_selected"`
            CurrencySymbol string  `json:"currency_symbol"`
         
    }
    
    // ProductFilters used by Catalog
    type ProductFilters struct {

        
            Key ProductFiltersKey  `json:"key"`
            Values []ProductFiltersValue  `json:"values"`
         
    }
    
    // ProductSortOn used by Catalog
    type ProductSortOn struct {

        
            IsSelected bool  `json:"is_selected"`
            Name string  `json:"name"`
            Value string  `json:"value"`
         
    }
    
    // GetCollectionQueryOptionResponse used by Catalog
    type GetCollectionQueryOptionResponse struct {

        
            Filters []ProductFilters  `json:"filters"`
            Operators map[string]string  `json:"operators"`
            SortOn []ProductSortOn  `json:"sort_on"`
         
    }
    
    // UserInfo used by Catalog
    type UserInfo struct {

        
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
            Email string  `json:"email"`
            UID string  `json:"uid"`
         
    }
    
    // CollectionQuery used by Catalog
    type CollectionQuery struct {

        
            Attribute string  `json:"attribute"`
            Op string  `json:"op"`
            Value []interface{}  `json:"value"`
         
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
    
    // SeoDetail used by Catalog
    type SeoDetail struct {

        
            Title string  `json:"title"`
            Description string  `json:"description"`
         
    }
    
    // CollectionBadge used by Catalog
    type CollectionBadge struct {

        
            Color string  `json:"color"`
            Text string  `json:"text"`
         
    }
    
    // CollectionSchedule used by Catalog
    type CollectionSchedule struct {

        
            Start string  `json:"start"`
            NextSchedule []NextSchedule  `json:"next_schedule"`
            Cron string  `json:"cron"`
            End string  `json:"end"`
            Duration float64  `json:"duration"`
         
    }
    
    // CreateCollection used by Catalog
    type CreateCollection struct {

        
            CreatedBy UserInfo  `json:"created_by"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            Description string  `json:"description"`
            IsActive bool  `json:"is_active"`
            Query []CollectionQuery  `json:"query"`
            Name string  `json:"name"`
            Banners CollectionBanner  `json:"banners"`
            Seo SeoDetail  `json:"seo"`
            Badge CollectionBadge  `json:"badge"`
            IsVisible bool  `json:"is_visible"`
            SortOn string  `json:"sort_on"`
            Slug string  `json:"slug"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Logo CollectionImage  `json:"logo"`
            AllowFacets bool  `json:"allow_facets"`
            ModifiedBy UserInfo  `json:"modified_by"`
            Tags []string  `json:"tags"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Meta map[string]interface{}  `json:"meta"`
            AllowSort bool  `json:"allow_sort"`
            Type string  `json:"type"`
            Schedule CollectionSchedule  `json:"_schedule"`
            Published bool  `json:"published"`
            Priority float64  `json:"priority"`
            AppID string  `json:"app_id"`
         
    }
    
    // BannerImage used by Catalog
    type BannerImage struct {

        
            URL string  `json:"url"`
            AspectRatio string  `json:"aspect_ratio"`
            Type string  `json:"type"`
         
    }
    
    // ImageUrls used by Catalog
    type ImageUrls struct {

        
            Landscape BannerImage  `json:"landscape"`
            Portrait BannerImage  `json:"portrait"`
         
    }
    
    // CollectionCreateResponse used by Catalog
    type CollectionCreateResponse struct {

        
            Description string  `json:"description"`
            IsActive bool  `json:"is_active"`
            Query []CollectionQuery  `json:"query"`
            Name string  `json:"name"`
            Cron map[string]interface{}  `json:"cron"`
            Banners ImageUrls  `json:"banners"`
            Tag []string  `json:"tag"`
            Badge map[string]interface{}  `json:"badge"`
            SortOn string  `json:"sort_on"`
            Slug string  `json:"slug"`
            Logo BannerImage  `json:"logo"`
            AllowFacets bool  `json:"allow_facets"`
            Meta map[string]interface{}  `json:"meta"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            AllowSort bool  `json:"allow_sort"`
            Type string  `json:"type"`
            Schedule map[string]interface{}  `json:"_schedule"`
            Priority float64  `json:"priority"`
            AppID string  `json:"app_id"`
         
    }
    
    // CollectionListingFilterTag used by Catalog
    type CollectionListingFilterTag struct {

        
            Display string  `json:"display"`
            IsSelected bool  `json:"is_selected"`
            Name string  `json:"name"`
         
    }
    
    // CollectionListingFilterType used by Catalog
    type CollectionListingFilterType struct {

        
            Display string  `json:"display"`
            IsSelected bool  `json:"is_selected"`
            Name string  `json:"name"`
         
    }
    
    // CollectionListingFilter used by Catalog
    type CollectionListingFilter struct {

        
            Tags []CollectionListingFilterTag  `json:"tags"`
            Type []CollectionListingFilterType  `json:"type"`
         
    }
    
    // Media1 used by Catalog
    type Media1 struct {

        
            URL string  `json:"url"`
            Type string  `json:"type"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // GetCollectionDetailNest used by Catalog
    type GetCollectionDetailNest struct {

        
            UID string  `json:"uid"`
            Description string  `json:"description"`
            IsActive bool  `json:"is_active"`
            Query []CollectionQuery  `json:"query"`
            Action Action  `json:"action"`
            Name string  `json:"name"`
            Cron map[string]interface{}  `json:"cron"`
            Banners ImageUrls  `json:"banners"`
            Tag []string  `json:"tag"`
            Badge map[string]interface{}  `json:"badge"`
            Slug string  `json:"slug"`
            Logo Media1  `json:"logo"`
            AllowFacets bool  `json:"allow_facets"`
            Meta map[string]interface{}  `json:"meta"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            AllowSort bool  `json:"allow_sort"`
            Type string  `json:"type"`
            Schedule map[string]interface{}  `json:"_schedule"`
            Priority float64  `json:"priority"`
            AppID string  `json:"app_id"`
         
    }
    
    // GetCollectionListingResponse used by Catalog
    type GetCollectionListingResponse struct {

        
            Page Page  `json:"page"`
            Filters CollectionListingFilter  `json:"filters"`
            Items []GetCollectionDetailNest  `json:"items"`
         
    }
    
    // CollectionDetailResponse used by Catalog
    type CollectionDetailResponse struct {

        
            AllowSort bool  `json:"allow_sort"`
            Cron map[string]interface{}  `json:"cron"`
            Meta map[string]interface{}  `json:"meta"`
            Banners ImageUrls  `json:"banners"`
            Slug string  `json:"slug"`
            Description string  `json:"description"`
            Type string  `json:"type"`
            Schedule map[string]interface{}  `json:"_schedule"`
            Logo Media1  `json:"logo"`
            Badge map[string]interface{}  `json:"badge"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            IsActive bool  `json:"is_active"`
            Priority float64  `json:"priority"`
            Tag []string  `json:"tag"`
            AllowFacets bool  `json:"allow_facets"`
            Query []CollectionQuery  `json:"query"`
            AppID string  `json:"app_id"`
            Name string  `json:"name"`
         
    }
    
    // UpdateCollection used by Catalog
    type UpdateCollection struct {

        
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            Description string  `json:"description"`
            IsActive bool  `json:"is_active"`
            Query []CollectionQuery  `json:"query"`
            Name string  `json:"name"`
            Banners CollectionBanner  `json:"banners"`
            Seo SeoDetail  `json:"seo"`
            Badge CollectionBadge  `json:"badge"`
            IsVisible bool  `json:"is_visible"`
            SortOn string  `json:"sort_on"`
            Slug string  `json:"slug"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Logo CollectionImage  `json:"logo"`
            AllowFacets bool  `json:"allow_facets"`
            ModifiedBy UserInfo  `json:"modified_by"`
            Tags []string  `json:"tags"`
            Meta map[string]interface{}  `json:"meta"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            AllowSort bool  `json:"allow_sort"`
            Type string  `json:"type"`
            Schedule CollectionSchedule  `json:"_schedule"`
            Published bool  `json:"published"`
            Priority float64  `json:"priority"`
         
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
    
    // Price1 used by Catalog
    type Price1 struct {

        
            Min float64  `json:"min"`
            Max float64  `json:"max"`
            CurrencyCode string  `json:"currency_code"`
            CurrencySymbol string  `json:"currency_symbol"`
         
    }
    
    // ProductListingPrice used by Catalog
    type ProductListingPrice struct {

        
            Effective Price1  `json:"effective"`
            Marked Price1  `json:"marked"`
         
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
    
    // ProductBrand used by Catalog
    type ProductBrand struct {

        
            Action Action  `json:"action"`
            Logo Media1  `json:"logo"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // ProductListingDetail used by Catalog
    type ProductListingDetail struct {

        
            Color string  `json:"color"`
            Price ProductListingPrice  `json:"price"`
            UID float64  `json:"uid"`
            Description string  `json:"description"`
            Sellable bool  `json:"sellable"`
            Medias []Media1  `json:"medias"`
            Similars []string  `json:"similars"`
            Tryouts []string  `json:"tryouts"`
            RatingCount float64  `json:"rating_count"`
            Name string  `json:"name"`
            PromoMeta map[string]interface{}  `json:"promo_meta"`
            ImageNature string  `json:"image_nature"`
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
            Attributes map[string]interface{}  `json:"attributes"`
            ItemCode string  `json:"item_code"`
            Discount string  `json:"discount"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            ShortDescription string  `json:"short_description"`
            Slug string  `json:"slug"`
            ItemType string  `json:"item_type"`
            Brand ProductBrand  `json:"brand"`
            Rating float64  `json:"rating"`
            Highlights []string  `json:"highlights"`
            HasVariant bool  `json:"has_variant"`
            Type string  `json:"type"`
            ProductOnlineDate string  `json:"product_online_date"`
         
    }
    
    // GetCollectionItemsResponse used by Catalog
    type GetCollectionItemsResponse struct {

        
            Page Page  `json:"page"`
            Filters []ProductFilters  `json:"filters"`
            Items []ProductListingDetail  `json:"items"`
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

        
            AvailableArticles float64  `json:"available_articles"`
            AvailableSizes float64  `json:"available_sizes"`
            ArticleFreshness float64  `json:"article_freshness"`
            TotalSizes float64  `json:"total_sizes"`
            TotalArticles float64  `json:"total_articles"`
            Name string  `json:"name"`
         
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

        
            OptLevel string  `json:"opt_level"`
            CompanyID float64  `json:"company_id"`
            BrandIds []float64  `json:"brand_ids"`
            Platform string  `json:"platform"`
            Enabled bool  `json:"enabled"`
            StoreIds []float64  `json:"store_ids"`
         
    }
    
    // CompanyOptIn used by Catalog
    type CompanyOptIn struct {

        
            OptLevel string  `json:"opt_level"`
            CompanyID float64  `json:"company_id"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            BrandIds []float64  `json:"brand_ids"`
            Platform string  `json:"platform"`
            CreatedOn float64  `json:"created_on"`
            Enabled bool  `json:"enabled"`
            ModifiedOn float64  `json:"modified_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            StoreIds []float64  `json:"store_ids"`
         
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
            Store float64  `json:"store"`
            Company string  `json:"company"`
         
    }
    
    // StoreDetail used by Catalog
    type StoreDetail struct {

        
            StoreCode string  `json:"store_code"`
            CompanyID float64  `json:"company_id"`
            UID float64  `json:"uid"`
            Manager map[string]interface{}  `json:"manager"`
            CreatedOn string  `json:"created_on"`
            DisplayName string  `json:"display_name"`
            ModifiedOn string  `json:"modified_on"`
            AdditionalContacts []map[string]interface{}  `json:"additional_contacts"`
            Documents []map[string]interface{}  `json:"documents"`
            Address map[string]interface{}  `json:"address"`
            Timing map[string]interface{}  `json:"timing"`
            Name string  `json:"name"`
            StoreType string  `json:"store_type"`
         
    }
    
    // OptinStoreDetails used by Catalog
    type OptinStoreDetails struct {

        
            Page Page  `json:"page"`
            Items []StoreDetail  `json:"items"`
         
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

        
            Mandatory bool  `json:"mandatory"`
            Type string  `json:"type"`
            Range AttributeSchemaRange  `json:"range"`
            AllowedValues []string  `json:"allowed_values"`
            Format string  `json:"format"`
            Multi bool  `json:"multi"`
         
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

        
            Slug string  `json:"slug"`
            IsNested bool  `json:"is_nested"`
            Description string  `json:"description"`
            EnabledForEndConsumer bool  `json:"enabled_for_end_consumer"`
            Filters AttributeMasterFilter  `json:"filters"`
            Logo string  `json:"logo"`
            Details AttributeMasterDetails  `json:"details"`
            Schema AttributeMaster  `json:"schema"`
            ID string  `json:"id"`
            Departments []string  `json:"departments"`
            Meta AttributeMasterMeta  `json:"meta"`
            Name string  `json:"name"`
         
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
            Status float64  `json:"status"`
            Code string  `json:"code"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // DepartmentCreateUpdate used by Catalog
    type DepartmentCreateUpdate struct {

        
            UID float64  `json:"uid"`
            Platforms map[string]interface{}  `json:"platforms"`
            Slug string  `json:"slug"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            PriorityOrder float64  `json:"priority_order"`
            Logo string  `json:"logo"`
            Cls string  `json:"_cls"`
            IsActive bool  `json:"is_active"`
            Synonyms []string  `json:"synonyms"`
            Tags []string  `json:"tags"`
            Name string  `json:"name"`
         
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
            UID string  `json:"uid"`
            UserID string  `json:"user_id"`
            Contact string  `json:"contact"`
            ID string  `json:"_id"`
         
    }
    
    // GetDepartment used by Catalog
    type GetDepartment struct {

        
            CreatedBy UserSerializer  `json:"created_by"`
            UID float64  `json:"uid"`
            Slug string  `json:"slug"`
            ItemType string  `json:"item_type"`
            PriorityOrder float64  `json:"priority_order"`
            Logo string  `json:"logo"`
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
            Search string  `json:"search"`
            IsActive bool  `json:"is_active"`
            Synonyms []string  `json:"synonyms"`
            PageNo float64  `json:"page_no"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            PageSize float64  `json:"page_size"`
            Name string  `json:"name"`
         
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
    
    // UserDetail used by Catalog
    type UserDetail struct {

        
            SuperUser bool  `json:"super_user"`
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
            Contact string  `json:"contact"`
         
    }
    
    // DepartmentModel used by Catalog
    type DepartmentModel struct {

        
            CreatedBy UserDetail  `json:"created_by"`
            UID float64  `json:"uid"`
            Slug interface{}  `json:"slug"`
            VerifiedOn string  `json:"verified_on"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            PriorityOrder float64  `json:"priority_order"`
            CreatedOn string  `json:"created_on"`
            Logo string  `json:"logo"`
            ID interface{}  `json:"_id"`
            Cls interface{}  `json:"_cls"`
            ModifiedOn string  `json:"modified_on"`
            IsActive bool  `json:"is_active"`
            Synonyms []interface{}  `json:"synonyms"`
            ModifiedBy UserDetail  `json:"modified_by"`
            Name interface{}  `json:"name"`
            VerifiedBy UserDetail  `json:"verified_by"`
         
    }
    
    // ProductTemplate used by Catalog
    type ProductTemplate struct {

        
            IsPhysical bool  `json:"is_physical"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Slug string  `json:"slug"`
            Description string  `json:"description"`
            Attributes []string  `json:"attributes"`
            Logo string  `json:"logo"`
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
            IsArchived bool  `json:"is_archived"`
            Categories []string  `json:"categories"`
            Tag string  `json:"tag"`
            IsActive bool  `json:"is_active"`
            IsExpirable bool  `json:"is_expirable"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Departments []string  `json:"departments"`
            Name string  `json:"name"`
         
    }
    
    // TemplatesResponse used by Catalog
    type TemplatesResponse struct {

        
            Page Page  `json:"page"`
            Items ProductTemplate  `json:"items"`
         
    }
    
    // Properties used by Catalog
    type Properties struct {

        
            Description map[string]interface{}  `json:"description"`
            TraderType map[string]interface{}  `json:"trader_type"`
            IsActive map[string]interface{}  `json:"is_active"`
            ProductGroupTag map[string]interface{}  `json:"product_group_tag"`
            ReturnConfig map[string]interface{}  `json:"return_config"`
            Sizes map[string]interface{}  `json:"sizes"`
            HsnCode map[string]interface{}  `json:"hsn_code"`
            Name map[string]interface{}  `json:"name"`
            CategorySlug map[string]interface{}  `json:"category_slug"`
            Media map[string]interface{}  `json:"media"`
            BrandUID map[string]interface{}  `json:"brand_uid"`
            ItemCode map[string]interface{}  `json:"item_code"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            ShortDescription map[string]interface{}  `json:"short_description"`
            Slug map[string]interface{}  `json:"slug"`
            ItemType map[string]interface{}  `json:"item_type"`
            SizeGuide map[string]interface{}  `json:"size_guide"`
            ProductPublish map[string]interface{}  `json:"product_publish"`
            NoOfBoxes map[string]interface{}  `json:"no_of_boxes"`
            Variants map[string]interface{}  `json:"variants"`
            Trader map[string]interface{}  `json:"trader"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            MultiSize map[string]interface{}  `json:"multi_size"`
            IsDependent map[string]interface{}  `json:"is_dependent"`
            Highlights map[string]interface{}  `json:"highlights"`
            Tags map[string]interface{}  `json:"tags"`
            Currency map[string]interface{}  `json:"currency"`
            CountryOfOrigin map[string]interface{}  `json:"country_of_origin"`
            Command map[string]interface{}  `json:"command"`
         
    }
    
    // GlobalValidation used by Catalog
    type GlobalValidation struct {

        
            Definitions map[string]interface{}  `json:"definitions"`
            Description string  `json:"description"`
            Type string  `json:"type"`
            Title string  `json:"title"`
            Properties Properties  `json:"properties"`
            Required []string  `json:"required"`
         
    }
    
    // TemplateValidationData used by Catalog
    type TemplateValidationData struct {

        
            GlobalValidation GlobalValidation  `json:"global_validation"`
            TemplateValidation map[string]interface{}  `json:"template_validation"`
         
    }
    
    // TemplateDetails used by Catalog
    type TemplateDetails struct {

        
            IsPhysical bool  `json:"is_physical"`
            Slug string  `json:"slug"`
            Description string  `json:"description"`
            Attributes []string  `json:"attributes"`
            Logo string  `json:"logo"`
            IsArchived bool  `json:"is_archived"`
            Categories []string  `json:"categories"`
            Tag string  `json:"tag"`
            IsActive bool  `json:"is_active"`
            ID string  `json:"id"`
            IsExpirable bool  `json:"is_expirable"`
            Departments []string  `json:"departments"`
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
            Type string  `json:"type"`
            Brand []string  `json:"brand"`
         
    }
    
    // ProductDownloadsItems used by Catalog
    type ProductDownloadsItems struct {

        
            SellerID float64  `json:"seller_id"`
            CreatedBy VerifiedBy  `json:"created_by"`
            CompletedOn string  `json:"completed_on"`
            TriggerOn string  `json:"trigger_on"`
            URL string  `json:"url"`
            Status string  `json:"status"`
            Data ProductDownloadItemsData  `json:"data"`
            ID string  `json:"id"`
            TaskID string  `json:"task_id"`
            TemplateTags map[string]interface{}  `json:"template_tags"`
         
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
    
    // Hierarchy used by Catalog
    type Hierarchy struct {

        
            Department float64  `json:"department"`
            L1 float64  `json:"l1"`
            L2 float64  `json:"l2"`
         
    }
    
    // Media2 used by Catalog
    type Media2 struct {

        
            Landscape string  `json:"landscape"`
            Logo string  `json:"logo"`
            Portrait string  `json:"portrait"`
         
    }
    
    // CategoryMappingValues used by Catalog
    type CategoryMappingValues struct {

        
            CatalogID float64  `json:"catalog_id"`
            Name string  `json:"name"`
         
    }
    
    // CategoryMapping used by Catalog
    type CategoryMapping struct {

        
            Facebook CategoryMappingValues  `json:"facebook"`
            Google CategoryMappingValues  `json:"google"`
            Ajio CategoryMappingValues  `json:"ajio"`
         
    }
    
    // CategoryRequestBody used by Catalog
    type CategoryRequestBody struct {

        
            Hierarchy []Hierarchy  `json:"hierarchy"`
            Slug string  `json:"slug"`
            Media Media2  `json:"media"`
            Level float64  `json:"level"`
            IsActive bool  `json:"is_active"`
            Priority float64  `json:"priority"`
            Tryouts []string  `json:"tryouts"`
            Synonyms []string  `json:"synonyms"`
            Departments []float64  `json:"departments"`
            Marketplaces CategoryMapping  `json:"marketplaces"`
            Name string  `json:"name"`
         
    }
    
    // CategoryCreateResponse used by Catalog
    type CategoryCreateResponse struct {

        
            Message string  `json:"message"`
            UID float64  `json:"uid"`
         
    }
    
    // Category used by Catalog
    type Category struct {

        
            Hierarchy []Hierarchy  `json:"hierarchy"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            UID float64  `json:"uid"`
            Slug string  `json:"slug"`
            Media Media2  `json:"media"`
            Level float64  `json:"level"`
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
            IsActive bool  `json:"is_active"`
            Priority float64  `json:"priority"`
            Tryouts []string  `json:"tryouts"`
            Synonyms []string  `json:"synonyms"`
            ID string  `json:"id"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Departments []float64  `json:"departments"`
            Marketplaces CategoryMapping  `json:"marketplaces"`
            Name string  `json:"name"`
         
    }
    
    // CategoryResponse used by Catalog
    type CategoryResponse struct {

        
            Page Page  `json:"page"`
            Items []Category  `json:"items"`
         
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
    
    // TaxIdentifier used by Catalog
    type TaxIdentifier struct {

        
            ReportingHsn string  `json:"reporting_hsn"`
            HsnCodeID string  `json:"hsn_code_id"`
            HsnCode string  `json:"hsn_code"`
         
    }
    
    // NetQuantity used by Catalog
    type NetQuantity struct {

        
            Unit interface{}  `json:"unit"`
            Value float64  `json:"value"`
         
    }
    
    // TeaserTag used by Catalog
    type TeaserTag struct {

        
            URL string  `json:"url"`
            Tag string  `json:"tag"`
         
    }
    
    // ProductPublish used by Catalog
    type ProductPublish struct {

        
            ProductOnlineDate string  `json:"product_online_date"`
            IsSet bool  `json:"is_set"`
         
    }
    
    // Trader used by Catalog
    type Trader struct {

        
            Address []string  `json:"address"`
            Type string  `json:"type"`
            Name interface{}  `json:"name"`
         
    }
    
    // CustomOrder used by Catalog
    type CustomOrder struct {

        
            IsCustomOrder bool  `json:"is_custom_order"`
            ManufacturingTimeUnit string  `json:"manufacturing_time_unit"`
            ManufacturingTime float64  `json:"manufacturing_time"`
         
    }
    
    // ProductCreateUpdateRequest used by Catalog
    type ProductCreateUpdateRequest struct {

        
            UID float64  `json:"uid"`
            Description string  `json:"description"`
            IsActive bool  `json:"is_active"`
            ProductGroupTag []string  `json:"product_group_tag"`
            ReturnConfig ReturnConfig  `json:"return_config"`
            TemplateTag string  `json:"template_tag"`
            Departments []float64  `json:"departments"`
            Action string  `json:"action"`
            Name interface{}  `json:"name"`
            CompanyID float64  `json:"company_id"`
            CategorySlug string  `json:"category_slug"`
            Media []Media1  `json:"media"`
            BrandUID float64  `json:"brand_uid"`
            BulkJobID string  `json:"bulk_job_id"`
            ItemCode interface{}  `json:"item_code"`
            VariantGroup map[string]interface{}  `json:"variant_group"`
            TaxIdentifier TaxIdentifier  `json:"tax_identifier"`
            IsImageLessProduct bool  `json:"is_image_less_product"`
            NetQuantity NetQuantity  `json:"net_quantity"`
            TeaserTag TeaserTag  `json:"teaser_tag"`
            ShortDescription string  `json:"short_description"`
            Slug string  `json:"slug"`
            ItemType string  `json:"item_type"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            SizeGuide string  `json:"size_guide"`
            ProductPublish ProductPublish  `json:"product_publish"`
            NoOfBoxes float64  `json:"no_of_boxes"`
            Variants map[string]interface{}  `json:"variants"`
            Trader []Trader  `json:"trader"`
            CustomOrder CustomOrder  `json:"custom_order"`
            MultiSize bool  `json:"multi_size"`
            IsDependent bool  `json:"is_dependent"`
            Highlights []string  `json:"highlights"`
            ChangeRequestID interface{}  `json:"change_request_id"`
            Tags []string  `json:"tags"`
            Currency string  `json:"currency"`
            CountryOfOrigin string  `json:"country_of_origin"`
            VariantMedia map[string]interface{}  `json:"variant_media"`
            Requester string  `json:"requester"`
            IsSet bool  `json:"is_set"`
         
    }
    
    // Logo used by Catalog
    type Logo struct {

        
            URL string  `json:"url"`
            AspectRatio string  `json:"aspect_ratio"`
            SecureURL string  `json:"secure_url"`
            AspectRatioF float64  `json:"aspect_ratio_f"`
         
    }
    
    // ProductPublished used by Catalog
    type ProductPublished struct {

        
            ProductOnlineDate float64  `json:"product_online_date"`
            IsSet bool  `json:"is_set"`
         
    }
    
    // Image used by Catalog
    type Image struct {

        
            URL string  `json:"url"`
            AspectRatio string  `json:"aspect_ratio"`
            SecureURL string  `json:"secure_url"`
            AspectRatioF float64  `json:"aspect_ratio_f"`
         
    }
    
    // Product used by Catalog
    type Product struct {

        
            IsPhysical bool  `json:"is_physical"`
            Color string  `json:"color"`
            UID float64  `json:"uid"`
            Description string  `json:"description"`
            IsActive bool  `json:"is_active"`
            IsExpirable bool  `json:"is_expirable"`
            TemplateTag string  `json:"template_tag"`
            Departments []float64  `json:"departments"`
            Sizes []map[string]interface{}  `json:"sizes"`
            HsnCode string  `json:"hsn_code"`
            Name string  `json:"name"`
            ImageNature string  `json:"image_nature"`
            CategorySlug string  `json:"category_slug"`
            Media []Media1  `json:"media"`
            BrandUID float64  `json:"brand_uid"`
            CategoryUID float64  `json:"category_uid"`
            ItemCode string  `json:"item_code"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            ID string  `json:"id"`
            VariantGroup map[string]interface{}  `json:"variant_group"`
            L3Mapping []string  `json:"l3_mapping"`
            ShortDescription string  `json:"short_description"`
            Slug string  `json:"slug"`
            ItemType string  `json:"item_type"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Brand Brand  `json:"brand"`
            SizeGuide string  `json:"size_guide"`
            ProductPublish ProductPublished  `json:"product_publish"`
            Images []Image  `json:"images"`
            Variants map[string]interface{}  `json:"variants"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            MultiSize bool  `json:"multi_size"`
            IsDependent bool  `json:"is_dependent"`
            Highlights []string  `json:"highlights"`
            Currency string  `json:"currency"`
            PrimaryColor string  `json:"primary_color"`
            CountryOfOrigin string  `json:"country_of_origin"`
            IsSet bool  `json:"is_set"`
            AllSizes []map[string]interface{}  `json:"all_sizes"`
         
    }
    
    // ProductListingResponse used by Catalog
    type ProductListingResponse struct {

        
            Page Page  `json:"page"`
            Items []Product  `json:"items"`
         
    }
    
    // ProductCreateUpdateSchemaV2 used by Catalog
    type ProductCreateUpdateSchemaV2 struct {

        
            UID float64  `json:"uid"`
            Description string  `json:"description"`
            IsActive bool  `json:"is_active"`
            ProductGroupTag []string  `json:"product_group_tag"`
            ReturnConfig ReturnConfig  `json:"return_config"`
            TemplateTag string  `json:"template_tag"`
            Departments []float64  `json:"departments"`
            Action string  `json:"action"`
            Name interface{}  `json:"name"`
            CompanyID float64  `json:"company_id"`
            CategorySlug string  `json:"category_slug"`
            Media []Media1  `json:"media"`
            BrandUID float64  `json:"brand_uid"`
            Attributes map[string]interface{}  `json:"attributes"`
            BulkJobID string  `json:"bulk_job_id"`
            ItemCode interface{}  `json:"item_code"`
            VariantGroup map[string]interface{}  `json:"variant_group"`
            TaxIdentifier TaxIdentifier  `json:"tax_identifier"`
            IsImageLessProduct bool  `json:"is_image_less_product"`
            NetQuantity NetQuantity  `json:"net_quantity"`
            TeaserTag TeaserTag  `json:"teaser_tag"`
            ShortDescription string  `json:"short_description"`
            Slug string  `json:"slug"`
            ItemType string  `json:"item_type"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            SizeGuide string  `json:"size_guide"`
            ProductPublish ProductPublish  `json:"product_publish"`
            NoOfBoxes float64  `json:"no_of_boxes"`
            Variants map[string]interface{}  `json:"variants"`
            Trader []Trader  `json:"trader"`
            CustomOrder CustomOrder  `json:"custom_order"`
            MultiSize bool  `json:"multi_size"`
            IsDependent bool  `json:"is_dependent"`
            Highlights []string  `json:"highlights"`
            ChangeRequestID interface{}  `json:"change_request_id"`
            Tags []string  `json:"tags"`
            Currency string  `json:"currency"`
            CountryOfOrigin string  `json:"country_of_origin"`
            VariantMedia map[string]interface{}  `json:"variant_media"`
            Requester string  `json:"requester"`
            IsSet bool  `json:"is_set"`
         
    }
    
    // ProductVariants used by Catalog
    type ProductVariants struct {

        
            UID float64  `json:"uid"`
            Media []Media1  `json:"media"`
            BrandUID float64  `json:"brand_uid"`
            CategoryUID float64  `json:"category_uid"`
            ItemCode string  `json:"item_code"`
            Name string  `json:"name"`
         
    }
    
    // ProductVariantsResponse used by Catalog
    type ProductVariantsResponse struct {

        
            Page Page  `json:"page"`
            Variants []ProductVariants  `json:"variants"`
         
    }
    
    // AttributeMasterSerializer used by Catalog
    type AttributeMasterSerializer struct {

        
            CreatedBy map[string]interface{}  `json:"created_by"`
            Description string  `json:"description"`
            Suggestion string  `json:"suggestion"`
            Details AttributeMasterDetails  `json:"details"`
            Departments []string  `json:"departments"`
            Variant bool  `json:"variant"`
            Name string  `json:"name"`
            EnabledForEndConsumer bool  `json:"enabled_for_end_consumer"`
            Synonyms map[string]interface{}  `json:"synonyms"`
            Slug string  `json:"slug"`
            RawKey string  `json:"raw_key"`
            Filters AttributeMasterFilter  `json:"filters"`
            Logo string  `json:"logo"`
            CreatedOn string  `json:"created_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Tags []string  `json:"tags"`
            Unit string  `json:"unit"`
            IsNested bool  `json:"is_nested"`
            ModifiedOn string  `json:"modified_on"`
            Schema AttributeMaster  `json:"schema"`
         
    }
    
    // ProductAttributesResponse used by Catalog
    type ProductAttributesResponse struct {

        
            Items []AttributeMasterSerializer  `json:"items"`
         
    }
    
    // ReturnConfig1 used by Catalog
    type ReturnConfig1 struct {

        
            Returnable bool  `json:"returnable"`
            Time float64  `json:"time"`
            Unit string  `json:"unit"`
         
    }
    
    // Media3 used by Catalog
    type Media3 struct {

        
            URL string  `json:"url"`
            Type string  `json:"type"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // TaxIdentifier1 used by Catalog
    type TaxIdentifier1 struct {

        
            ReportingHsn string  `json:"reporting_hsn"`
            HsnCodeID string  `json:"hsn_code_id"`
            HsnCode string  `json:"hsn_code"`
         
    }
    
    // TeaserTag1 used by Catalog
    type TeaserTag1 struct {

        
            URL string  `json:"url"`
            Tag string  `json:"tag"`
         
    }
    
    // ProductPublish1 used by Catalog
    type ProductPublish1 struct {

        
            ProductOnlineDate string  `json:"product_online_date"`
            IsSet bool  `json:"is_set"`
         
    }
    
    // Trader1 used by Catalog
    type Trader1 struct {

        
            Address []string  `json:"address"`
            Type string  `json:"type"`
            Name string  `json:"name"`
         
    }
    
    // CustomOrder1 used by Catalog
    type CustomOrder1 struct {

        
            IsCustomOrder bool  `json:"is_custom_order"`
            ManufacturingTimeUnit string  `json:"manufacturing_time_unit"`
            ManufacturingTime float64  `json:"manufacturing_time"`
         
    }
    
    // ProductPatch used by Catalog
    type ProductPatch struct {

        
            UID float64  `json:"uid"`
            Description string  `json:"description"`
            IsActive bool  `json:"is_active"`
            ProductGroupTag []string  `json:"product_group_tag"`
            ReturnConfig ReturnConfig1  `json:"return_config"`
            TemplateTag string  `json:"template_tag"`
            Departments []float64  `json:"departments"`
            Action string  `json:"action"`
            Name string  `json:"name"`
            CompanyID float64  `json:"company_id"`
            CategorySlug string  `json:"category_slug"`
            Media []Media3  `json:"media"`
            BrandUID float64  `json:"brand_uid"`
            BulkJobID string  `json:"bulk_job_id"`
            ItemCode interface{}  `json:"item_code"`
            TaxIdentifier TaxIdentifier1  `json:"tax_identifier"`
            IsImageLessProduct bool  `json:"is_image_less_product"`
            NetQuantity NetQuantity  `json:"net_quantity"`
            TeaserTag TeaserTag1  `json:"teaser_tag"`
            ShortDescription string  `json:"short_description"`
            Slug string  `json:"slug"`
            ItemType string  `json:"item_type"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            SizeGuide string  `json:"size_guide"`
            ProductPublish ProductPublish1  `json:"product_publish"`
            NoOfBoxes float64  `json:"no_of_boxes"`
            Variants map[string]interface{}  `json:"variants"`
            Trader []Trader1  `json:"trader"`
            CustomOrder CustomOrder1  `json:"custom_order"`
            MultiSize bool  `json:"multi_size"`
            IsDependent bool  `json:"is_dependent"`
            Highlights []string  `json:"highlights"`
            ChangeRequestID interface{}  `json:"change_request_id"`
            Tags []string  `json:"tags"`
            Currency string  `json:"currency"`
            CountryOfOrigin string  `json:"country_of_origin"`
            VariantMedia map[string]interface{}  `json:"variant_media"`
            Requester string  `json:"requester"`
            IsSet bool  `json:"is_set"`
         
    }
    
    // SuccessResponseProduct used by Catalog
    type SuccessResponseProduct struct {

        
            Success bool  `json:"success"`
         
    }
    
    // ValidateIdentifier used by Catalog
    type ValidateIdentifier struct {

        
            GtinValue string  `json:"gtin_value"`
            Primary bool  `json:"primary"`
            GtinType string  `json:"gtin_type"`
         
    }
    
    // ALLSizes used by Catalog
    type ALLSizes struct {

        
            ItemLength float64  `json:"item_length"`
            Size interface{}  `json:"size"`
            ItemWidth float64  `json:"item_width"`
            ItemWeight float64  `json:"item_weight"`
            ItemHeight float64  `json:"item_height"`
            ItemDimensionsUnitOfMeasure string  `json:"item_dimensions_unit_of_measure"`
            Identifiers []ValidateIdentifier  `json:"identifiers"`
            ItemWeightUnitOfMeasure interface{}  `json:"item_weight_unit_of_measure"`
         
    }
    
    // ListALLSizes used by Catalog
    type ListALLSizes struct {

        
            AllSizes []ALLSizes  `json:"all_sizes"`
         
    }
    
    // ValidateProduct used by Catalog
    type ValidateProduct struct {

        
            Valid bool  `json:"valid"`
         
    }
    
    // UserInfo1 used by Catalog
    type UserInfo1 struct {

        
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
            Email string  `json:"email"`
            UID string  `json:"uid"`
         
    }
    
    // BulkJob used by Catalog
    type BulkJob struct {

        
            CreatedBy UserInfo1  `json:"created_by"`
            CancelledRecords []map[string]interface{}  `json:"cancelled_records"`
            Total float64  `json:"total"`
            IsActive bool  `json:"is_active"`
            Succeed float64  `json:"succeed"`
            TemplateTag string  `json:"template_tag"`
            CompanyID float64  `json:"company_id"`
            ProductType string  `json:"product_type"`
            Cancelled float64  `json:"cancelled"`
            Failed float64  `json:"failed"`
            FailedRecords []map[string]interface{}  `json:"failed_records"`
            TrackingURL string  `json:"tracking_url"`
            CreatedOn string  `json:"created_on"`
            Department string  `json:"department"`
            ModifiedBy UserInfo1  `json:"modified_by"`
            FilePath string  `json:"file_path"`
            CustomTemplateTag string  `json:"custom_template_tag"`
            Stage string  `json:"stage"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // BulkResponse used by Catalog
    type BulkResponse struct {

        
            CreatedBy UserInfo1  `json:"created_by"`
            BatchID string  `json:"batch_id"`
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
            IsActive bool  `json:"is_active"`
            ModifiedBy UserInfo1  `json:"modified_by"`
         
    }
    
    // UserDetail1 used by Catalog
    type UserDetail1 struct {

        
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
            FullName string  `json:"full_name"`
         
    }
    
    // ProductBulkRequest used by Catalog
    type ProductBulkRequest struct {

        
            Template ProductTemplate  `json:"template"`
            CompanyID float64  `json:"company_id"`
            CreatedBy UserDetail1  `json:"created_by"`
            CancelledRecords []string  `json:"cancelled_records"`
            FailedRecords []string  `json:"failed_records"`
            Cancelled float64  `json:"cancelled"`
            FilePath string  `json:"file_path"`
            CreatedOn string  `json:"created_on"`
            Stage string  `json:"stage"`
            Total float64  `json:"total"`
            ModifiedOn string  `json:"modified_on"`
            Failed float64  `json:"failed"`
            IsActive bool  `json:"is_active"`
            Succeed float64  `json:"succeed"`
            ModifiedBy UserDetail1  `json:"modified_by"`
            TemplateTag string  `json:"template_tag"`
         
    }
    
    // ProductBulkRequestList used by Catalog
    type ProductBulkRequestList struct {

        
            Page Page  `json:"page"`
            Items ProductBulkRequest  `json:"items"`
         
    }
    
    // BulkProductRequest used by Catalog
    type BulkProductRequest struct {

        
            CompanyID float64  `json:"company_id"`
            BatchID string  `json:"batch_id"`
            Data []map[string]interface{}  `json:"data"`
            TemplateTag string  `json:"template_tag"`
         
    }
    
    // NestedTags used by Catalog
    type NestedTags struct {

        
            Tags []string  `json:"tags"`
         
    }
    
    // ProductTagsViewResponse used by Catalog
    type ProductTagsViewResponse struct {

        
            Items NestedTags  `json:"items"`
         
    }
    
    // ProductBulkAssets used by Catalog
    type ProductBulkAssets struct {

        
            URL string  `json:"url"`
            User map[string]interface{}  `json:"user"`
            CompanyID float64  `json:"company_id"`
         
    }
    
    // UserCommon used by Catalog
    type UserCommon struct {

        
            CompanyID float64  `json:"company_id"`
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
         
    }
    
    // Items used by Catalog
    type Items struct {

        
            CreatedBy UserCommon  `json:"created_by"`
            CompanyID float64  `json:"company_id"`
            CancelledRecords []string  `json:"cancelled_records"`
            FailedRecords []string  `json:"failed_records"`
            TrackingURL string  `json:"tracking_url"`
            Cancelled float64  `json:"cancelled"`
            FilePath string  `json:"file_path"`
            CreatedOn string  `json:"created_on"`
            Retry float64  `json:"retry"`
            Stage string  `json:"stage"`
            Total float64  `json:"total"`
            ModifiedOn string  `json:"modified_on"`
            Failed float64  `json:"failed"`
            IsActive bool  `json:"is_active"`
            ID string  `json:"id"`
            Succeed float64  `json:"succeed"`
            ModifiedBy UserCommon  `json:"modified_by"`
         
    }
    
    // BulkAssetResponse used by Catalog
    type BulkAssetResponse struct {

        
            Page Page  `json:"page"`
            Items []Items  `json:"items"`
         
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

        
            Quantity float64  `json:"quantity"`
            SizeDistribution SizeDistribution  `json:"size_distribution"`
            Name string  `json:"name"`
         
    }
    
    // InvSize used by Catalog
    type InvSize struct {

        
            Size interface{}  `json:"size"`
            ItemLength float64  `json:"item_length"`
            Price float64  `json:"price"`
            ItemWidth float64  `json:"item_width"`
            ItemWeight float64  `json:"item_weight"`
            ItemHeight float64  `json:"item_height"`
            StoreCode string  `json:"store_code"`
            PriceEffective float64  `json:"price_effective"`
            ItemDimensionsUnitOfMeasure string  `json:"item_dimensions_unit_of_measure"`
            Currency string  `json:"currency"`
            Identifiers []GTIN  `json:"identifiers"`
            IsSet bool  `json:"is_set"`
            ExpirationDate string  `json:"expiration_date"`
            Set InventorySet  `json:"set"`
            ItemWeightUnitOfMeasure string  `json:"item_weight_unit_of_measure"`
            Quantity float64  `json:"quantity"`
            PriceTransfer float64  `json:"price_transfer"`
         
    }
    
    // ItemQuery used by Catalog
    type ItemQuery struct {

        
            BrandUID float64  `json:"brand_uid"`
            UID float64  `json:"uid"`
            ItemCode string  `json:"item_code"`
         
    }
    
    // InventoryRequest used by Catalog
    type InventoryRequest struct {

        
            CompanyID float64  `json:"company_id"`
            Sizes []InvSize  `json:"sizes"`
            Item ItemQuery  `json:"item"`
         
    }
    
    // InventoryResponse used by Catalog
    type InventoryResponse struct {

        
            Size string  `json:"size"`
            SellerIdentifier string  `json:"seller_identifier"`
            Price float64  `json:"price"`
            Currency string  `json:"currency"`
            UID string  `json:"uid"`
            PriceEffective float64  `json:"price_effective"`
            InventoryUpdatedOn string  `json:"inventory_updated_on"`
            SellableQuantity float64  `json:"sellable_quantity"`
            Identifiers map[string]interface{}  `json:"identifiers"`
            Store map[string]interface{}  `json:"store"`
            Quantity float64  `json:"quantity"`
            ItemID float64  `json:"item_id"`
            PriceTransfer float64  `json:"price_transfer"`
         
    }
    
    // InventoryResponsePaginated used by Catalog
    type InventoryResponsePaginated struct {

        
            Page Page  `json:"page"`
            Items []InventoryResponse  `json:"items"`
         
    }
    
    // PriceMeta used by Catalog
    type PriceMeta struct {

        
            TpNotes map[string]interface{}  `json:"tp_notes"`
            Currency string  `json:"currency"`
            Marked float64  `json:"marked"`
            Effective float64  `json:"effective"`
            Transfer float64  `json:"transfer"`
            UpdatedAt string  `json:"updated_at"`
         
    }
    
    // ReturnConfig2 used by Catalog
    type ReturnConfig2 struct {

        
            Returnable bool  `json:"returnable"`
            Time float64  `json:"time"`
            Unit string  `json:"unit"`
         
    }
    
    // ManufacturerResponse used by Catalog
    type ManufacturerResponse struct {

        
            Address string  `json:"address"`
            IsDefault bool  `json:"is_default"`
            Name string  `json:"name"`
         
    }
    
    // DimensionResponse used by Catalog
    type DimensionResponse struct {

        
            IsDefault bool  `json:"is_default"`
            Width float64  `json:"width"`
            Height float64  `json:"height"`
            Length float64  `json:"length"`
            Unit string  `json:"unit"`
         
    }
    
    // CompanyMeta used by Catalog
    type CompanyMeta struct {

        
            ID float64  `json:"id"`
         
    }
    
    // BrandMeta used by Catalog
    type BrandMeta struct {

        
            ID float64  `json:"id"`
            Name string  `json:"name"`
         
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
            Sellable QuantityBase  `json:"sellable"`
            NotAvailable QuantityBase  `json:"not_available"`
         
    }
    
    // Trader2 used by Catalog
    type Trader2 struct {

        
            Address []string  `json:"address"`
            Type string  `json:"type"`
            Name string  `json:"name"`
         
    }
    
    // WeightResponse used by Catalog
    type WeightResponse struct {

        
            IsDefault bool  `json:"is_default"`
            Unit string  `json:"unit"`
            Shipping float64  `json:"shipping"`
         
    }
    
    // InventorySellerResponse used by Catalog
    type InventorySellerResponse struct {

        
            SellerIdentifier string  `json:"seller_identifier"`
            Price PriceMeta  `json:"price"`
            CreatedBy UserSerializer  `json:"created_by"`
            UID string  `json:"uid"`
            FyndMeta map[string]interface{}  `json:"fynd_meta"`
            ExpirationDate string  `json:"expiration_date"`
            IsActive bool  `json:"is_active"`
            ReturnConfig ReturnConfig2  `json:"return_config"`
            Manufacturer ManufacturerResponse  `json:"manufacturer"`
            Dimension DimensionResponse  `json:"dimension"`
            FyndArticleCode string  `json:"fynd_article_code"`
            FyndItemCode string  `json:"fynd_item_code"`
            Set InventorySet  `json:"set"`
            Fragile bool  `json:"fragile"`
            Identifier map[string]interface{}  `json:"identifier"`
            AddedOnStore string  `json:"added_on_store"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            Company CompanyMeta  `json:"company"`
            Size string  `json:"size"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Brand BrandMeta  `json:"brand"`
            RawMeta map[string]interface{}  `json:"raw_meta"`
            TraceID string  `json:"trace_id"`
            Quantities Quantities  `json:"quantities"`
            Store StoreMeta  `json:"store"`
            Trader []Trader2  `json:"trader"`
            ItemID float64  `json:"item_id"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            Meta map[string]interface{}  `json:"meta"`
            Tags []string  `json:"tags"`
            TrackInventory bool  `json:"track_inventory"`
            CountryOfOrigin string  `json:"country_of_origin"`
            IsSet bool  `json:"is_set"`
            Stage string  `json:"stage"`
            Weight WeightResponse  `json:"weight"`
            TotalQuantity float64  `json:"total_quantity"`
         
    }
    
    // InventorySellerIdentifierResponsePaginated used by Catalog
    type InventorySellerIdentifierResponsePaginated struct {

        
            Page Page  `json:"page"`
            Items []InventorySellerResponse  `json:"items"`
         
    }
    
    // BulkInventoryGetItems used by Catalog
    type BulkInventoryGetItems struct {

        
            CreatedBy map[string]interface{}  `json:"created_by"`
            CompanyID float64  `json:"company_id"`
            CancelledRecords []string  `json:"cancelled_records"`
            FailedRecords []string  `json:"failed_records"`
            Cancelled float64  `json:"cancelled"`
            FilePath string  `json:"file_path"`
            CreatedOn string  `json:"created_on"`
            Stage string  `json:"stage"`
            Total float64  `json:"total"`
            ModifiedOn string  `json:"modified_on"`
            Failed float64  `json:"failed"`
            ID string  `json:"id"`
            IsActive bool  `json:"is_active"`
            Succeed float64  `json:"succeed"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
         
    }
    
    // BulkInventoryGet used by Catalog
    type BulkInventoryGet struct {

        
            Page Page  `json:"page"`
            Items []BulkInventoryGetItems  `json:"items"`
         
    }
    
    // PriceArticle used by Catalog
    type PriceArticle struct {

        
            TpNotes map[string]interface{}  `json:"tp_notes"`
            Currency string  `json:"currency"`
            Marked float64  `json:"marked"`
            Effective float64  `json:"effective"`
            Transfer float64  `json:"transfer"`
         
    }
    
    // ManufacturerResponse1 used by Catalog
    type ManufacturerResponse1 struct {

        
            Address string  `json:"address"`
            IsDefault bool  `json:"is_default"`
            Name string  `json:"name"`
         
    }
    
    // ReturnConfig3 used by Catalog
    type ReturnConfig3 struct {

        
            Returnable bool  `json:"returnable"`
            Time float64  `json:"time"`
            Unit string  `json:"unit"`
         
    }
    
    // DimensionResponse1 used by Catalog
    type DimensionResponse1 struct {

        
            Width float64  `json:"width"`
            Length float64  `json:"length"`
            Unit string  `json:"unit"`
            Height float64  `json:"height"`
         
    }
    
    // CompanyMeta1 used by Catalog
    type CompanyMeta1 struct {

        
            ID float64  `json:"id"`
         
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

        
            Damaged Quantity  `json:"damaged"`
            OrderCommitted Quantity  `json:"order_committed"`
            Sellable Quantity  `json:"sellable"`
            NotAvailable Quantity  `json:"not_available"`
         
    }
    
    // ArticleStoreResponse used by Catalog
    type ArticleStoreResponse struct {

        
            StoreCode string  `json:"store_code"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            StoreType string  `json:"store_type"`
         
    }
    
    // Trader3 used by Catalog
    type Trader3 struct {

        
            Address []string  `json:"address"`
            Type string  `json:"type"`
            Name string  `json:"name"`
         
    }
    
    // WeightResponse1 used by Catalog
    type WeightResponse1 struct {

        
            Unit string  `json:"unit"`
            Shipping float64  `json:"shipping"`
         
    }
    
    // GetInventories used by Catalog
    type GetInventories struct {

        
            SellerIdentifier string  `json:"seller_identifier"`
            Price PriceArticle  `json:"price"`
            CreatedBy UserSerializer  `json:"created_by"`
            UID string  `json:"uid"`
            ExpirationDate string  `json:"expiration_date"`
            DateMeta DateMeta  `json:"date_meta"`
            Manufacturer ManufacturerResponse1  `json:"manufacturer"`
            ReturnConfig ReturnConfig3  `json:"return_config"`
            Dimension DimensionResponse1  `json:"dimension"`
            InventoryUpdatedOn string  `json:"inventory_updated_on"`
            Identifier map[string]interface{}  `json:"identifier"`
            ID string  `json:"id"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            Company CompanyMeta1  `json:"company"`
            Size string  `json:"size"`
            Platforms map[string]interface{}  `json:"platforms"`
            Brand BrandMeta1  `json:"brand"`
            TraceID string  `json:"trace_id"`
            Quantities QuantitiesArticle  `json:"quantities"`
            Store ArticleStoreResponse  `json:"store"`
            Trader []Trader3  `json:"trader"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            Tags []string  `json:"tags"`
            TrackInventory bool  `json:"track_inventory"`
            CountryOfOrigin string  `json:"country_of_origin"`
            IsSet bool  `json:"is_set"`
            Stage string  `json:"stage"`
            Weight WeightResponse1  `json:"weight"`
            TotalQuantity float64  `json:"total_quantity"`
         
    }
    
    // GetInventoriesResponse used by Catalog
    type GetInventoriesResponse struct {

        
            Page Page  `json:"page"`
            Items []GetInventories  `json:"items"`
         
    }
    
    // InventoryJobPayload used by Catalog
    type InventoryJobPayload struct {

        
            SellerIdentifier string  `json:"seller_identifier"`
            StoreCode string  `json:"store_code"`
            Price float64  `json:"price"`
            Currency string  `json:"currency"`
            PriceEffective float64  `json:"price_effective"`
            ItemDimensionsUnitOfMeasure string  `json:"item_dimensions_unit_of_measure"`
            ExpirationDate string  `json:"expiration_date"`
            PriceMarked float64  `json:"price_marked"`
            TraceID string  `json:"trace_id"`
            ItemWeightUnitOfMeasure string  `json:"item_weight_unit_of_measure"`
            Quantity float64  `json:"quantity"`
            Tags []string  `json:"tags"`
            TotalQuantity float64  `json:"total_quantity"`
         
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

        
            Type string  `json:"type"`
            Store []float64  `json:"store"`
            Brand []float64  `json:"brand"`
         
    }
    
    // InventoryExportResponse used by Catalog
    type InventoryExportResponse struct {

        
            SellerID float64  `json:"seller_id"`
            TriggerOn string  `json:"trigger_on"`
            Status string  `json:"status"`
            RequestParams map[string]interface{}  `json:"request_params"`
            TaskID string  `json:"task_id"`
         
    }
    
    // InventoryExportJob used by Catalog
    type InventoryExportJob struct {

        
            SellerID float64  `json:"seller_id"`
            URL string  `json:"url"`
            CompletedOn string  `json:"completed_on"`
            TriggerOn string  `json:"trigger_on"`
            Status string  `json:"status"`
            RequestParams map[string]interface{}  `json:"request_params"`
            TaskID string  `json:"task_id"`
         
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

        
            SellerIdentifier string  `json:"seller_identifier"`
            PriceEffective float64  `json:"price_effective"`
            ExpirationDate string  `json:"expiration_date"`
            PriceMarked float64  `json:"price_marked"`
            TraceID string  `json:"trace_id"`
            StoreID float64  `json:"store_id"`
            Tags []string  `json:"tags"`
            TotalQuantity float64  `json:"total_quantity"`
         
    }
    
    // InventoryRequestSchemaV2 used by Catalog
    type InventoryRequestSchemaV2 struct {

        
            CompanyID float64  `json:"company_id"`
            Payload []InventoryPayload  `json:"payload"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // InventoryFailedReason used by Catalog
    type InventoryFailedReason struct {

        
            Message string  `json:"message"`
            Errors string  `json:"errors"`
         
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
    
    // HsnUpsert used by Catalog
    type HsnUpsert struct {

        
            Threshold2 float64  `json:"threshold2"`
            Tax2 float64  `json:"tax2"`
            CompanyID float64  `json:"company_id"`
            Tax1 float64  `json:"tax1"`
            UID float64  `json:"uid"`
            TaxOnMrp bool  `json:"tax_on_mrp"`
            TaxOnEsp bool  `json:"tax_on_esp"`
            IsActive bool  `json:"is_active"`
            Threshold1 float64  `json:"threshold1"`
            HsnCode string  `json:"hsn_code"`
            Hs2Code string  `json:"hs2_code"`
         
    }
    
    // HsnCodesObject used by Catalog
    type HsnCodesObject struct {

        
            Threshold2 float64  `json:"threshold2"`
            Tax2 float64  `json:"tax2"`
            CompanyID float64  `json:"company_id"`
            Tax1 float64  `json:"tax1"`
            ModifiedOn string  `json:"modified_on"`
            TaxOnMrp bool  `json:"tax_on_mrp"`
            TaxOnEsp bool  `json:"tax_on_esp"`
            ID string  `json:"id"`
            Threshold1 float64  `json:"threshold1"`
            HsnCode string  `json:"hsn_code"`
            Hs2Code string  `json:"hs2_code"`
         
    }
    
    // HsnCode used by Catalog
    type HsnCode struct {

        
            Data HsnCodesObject  `json:"data"`
         
    }
    
    // PageResponse used by Catalog
    type PageResponse struct {

        
            Size float64  `json:"size"`
            Current string  `json:"current"`
            HasNext bool  `json:"has_next"`
            ItemTotal float64  `json:"item_total"`
            HasPrevious bool  `json:"has_previous"`
         
    }
    
    // HsnCodesListingResponse used by Catalog
    type HsnCodesListingResponse struct {

        
            Page PageResponse  `json:"page"`
            Items []HsnCodesObject  `json:"items"`
         
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
            EffectiveDate string  `json:"effective_date"`
            Rate float64  `json:"rate"`
            Threshold float64  `json:"threshold"`
         
    }
    
    // HSNDataInsertV2 used by Catalog
    type HSNDataInsertV2 struct {

        
            CountryCode string  `json:"country_code"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Description string  `json:"description"`
            Type string  `json:"type"`
            ReportingHsn string  `json:"reporting_hsn"`
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Taxes []TaxSlab  `json:"taxes"`
            HsnCode string  `json:"hsn_code"`
         
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
            ItemID float64  `json:"item_id"`
            IgnoredStores []float64  `json:"ignored_stores"`
         
    }
    
    // AssignStoreArticle used by Catalog
    type AssignStoreArticle struct {

        
            ArticleAssignment ArticleAssignment  `json:"article_assignment"`
            GroupID string  `json:"group_id"`
            Quantity float64  `json:"quantity"`
            Query ArticleQuery  `json:"query"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // AssignStore used by Catalog
    type AssignStore struct {

        
            CompanyID float64  `json:"company_id"`
            ChannelIdentifier string  `json:"channel_identifier"`
            Articles []AssignStoreArticle  `json:"articles"`
            ChannelType string  `json:"channel_type"`
            Pincode string  `json:"pincode"`
            AppID string  `json:"app_id"`
            StoreIds []float64  `json:"store_ids"`
         
    }
    
    // ArticleAssignment1 used by Catalog
    type ArticleAssignment1 struct {

        
            Level string  `json:"level"`
            Strategy string  `json:"strategy"`
         
    }
    
    // StoreAssignResponse used by Catalog
    type StoreAssignResponse struct {

        
            Size string  `json:"size"`
            CompanyID float64  `json:"company_id"`
            UID string  `json:"uid"`
            PriceEffective float64  `json:"price_effective"`
            StrategyWiseListing []map[string]interface{}  `json:"strategy_wise_listing"`
            SCity string  `json:"s_city"`
            ID string  `json:"_id"`
            Status bool  `json:"status"`
            Index float64  `json:"index"`
            PriceMarked float64  `json:"price_marked"`
            StorePincode float64  `json:"store_pincode"`
            ArticleAssignment ArticleAssignment1  `json:"article_assignment"`
            GroupID string  `json:"group_id"`
            StoreID float64  `json:"store_id"`
            Quantity float64  `json:"quantity"`
            ItemID float64  `json:"item_id"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // BrandItem used by Catalog
    type BrandItem struct {

        
            UID float64  `json:"uid"`
            Banners ImageUrls  `json:"banners"`
            Slug string  `json:"slug"`
            Logo Media  `json:"logo"`
            Discount string  `json:"discount"`
            Departments []string  `json:"departments"`
            Action Action  `json:"action"`
            Name string  `json:"name"`
         
    }
    
    // BrandListingResponse used by Catalog
    type BrandListingResponse struct {

        
            Page Page  `json:"page"`
            Items []BrandItem  `json:"items"`
         
    }
    
    // Department used by Catalog
    type Department struct {

        
            UID float64  `json:"uid"`
            Slug string  `json:"slug"`
            PriorityOrder float64  `json:"priority_order"`
            Logo Media  `json:"logo"`
            Name string  `json:"name"`
         
    }
    
    // DepartmentResponse used by Catalog
    type DepartmentResponse struct {

        
            Items []Department  `json:"items"`
         
    }
    
    // ThirdLevelChild used by Catalog
    type ThirdLevelChild struct {

        
            UID float64  `json:"uid"`
            Banners ImageUrls  `json:"banners"`
            Slug string  `json:"slug"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Childs []map[string]interface{}  `json:"childs"`
            Action Action  `json:"action"`
            Name string  `json:"name"`
         
    }
    
    // SecondLevelChild used by Catalog
    type SecondLevelChild struct {

        
            UID float64  `json:"uid"`
            Banners ImageUrls  `json:"banners"`
            Slug string  `json:"slug"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Childs []ThirdLevelChild  `json:"childs"`
            Action Action  `json:"action"`
            Name string  `json:"name"`
         
    }
    
    // Child used by Catalog
    type Child struct {

        
            UID float64  `json:"uid"`
            Banners ImageUrls  `json:"banners"`
            Slug string  `json:"slug"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Childs []SecondLevelChild  `json:"childs"`
            Action Action  `json:"action"`
            Name string  `json:"name"`
         
    }
    
    // CategoryItems used by Catalog
    type CategoryItems struct {

        
            UID float64  `json:"uid"`
            Banners ImageUrls  `json:"banners"`
            Slug string  `json:"slug"`
            Childs []Child  `json:"childs"`
            Action Action  `json:"action"`
            Name string  `json:"name"`
         
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

        
            SortOn []ProductSortOn  `json:"sort_on"`
            Filters []ProductFilters  `json:"filters"`
            Items []ProductListingDetail  `json:"items"`
            Operators map[string]interface{}  `json:"operators"`
            Page Page  `json:"page"`
         
    }
    
    // ProductDetail used by Catalog
    type ProductDetail struct {

        
            Color string  `json:"color"`
            UID float64  `json:"uid"`
            Description string  `json:"description"`
            Medias []Media1  `json:"medias"`
            Similars []string  `json:"similars"`
            Tryouts []string  `json:"tryouts"`
            RatingCount float64  `json:"rating_count"`
            Name string  `json:"name"`
            PromoMeta map[string]interface{}  `json:"promo_meta"`
            ImageNature string  `json:"image_nature"`
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
            Attributes map[string]interface{}  `json:"attributes"`
            ItemCode string  `json:"item_code"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            ShortDescription string  `json:"short_description"`
            Slug string  `json:"slug"`
            ItemType string  `json:"item_type"`
            Brand ProductBrand  `json:"brand"`
            Rating float64  `json:"rating"`
            Highlights []string  `json:"highlights"`
            HasVariant bool  `json:"has_variant"`
            Type string  `json:"type"`
            ProductOnlineDate string  `json:"product_online_date"`
         
    }
    
    // InventoryPage used by Catalog
    type InventoryPage struct {

        
            NextID string  `json:"next_id"`
            HasNext bool  `json:"has_next"`
            Type string  `json:"type"`
            ItemTotal float64  `json:"item_total"`
            HasPrevious bool  `json:"has_previous"`
         
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
    
    // LocationIntegrationType used by Catalog
    type LocationIntegrationType struct {

        
            Inventory string  `json:"inventory"`
            Order string  `json:"order"`
         
    }
    
    // GetAddressSerializer used by Catalog
    type GetAddressSerializer struct {

        
            CountryCode string  `json:"country_code"`
            City string  `json:"city"`
            Address2 string  `json:"address2"`
            Pincode float64  `json:"pincode"`
            Address1 string  `json:"address1"`
            Longitude float64  `json:"longitude"`
            State string  `json:"state"`
            Latitude float64  `json:"latitude"`
            AddressType string  `json:"address_type"`
            Country string  `json:"country"`
            Landmark string  `json:"landmark"`
         
    }
    
    // UserSerializer2 used by Catalog
    type UserSerializer2 struct {

        
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
            Contact string  `json:"contact"`
         
    }
    
    // GetCompanySerializer used by Catalog
    type GetCompanySerializer struct {

        
            RejectReason string  `json:"reject_reason"`
            CreatedBy UserSerializer2  `json:"created_by"`
            BusinessType string  `json:"business_type"`
            CompanyType string  `json:"company_type"`
            UID float64  `json:"uid"`
            VerifiedOn string  `json:"verified_on"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            CreatedOn string  `json:"created_on"`
            Stage string  `json:"stage"`
            ModifiedOn string  `json:"modified_on"`
            ModifiedBy UserSerializer2  `json:"modified_by"`
            Name string  `json:"name"`
            VerifiedBy UserSerializer2  `json:"verified_by"`
         
    }
    
    // ProductReturnConfigSerializer used by Catalog
    type ProductReturnConfigSerializer struct {

        
            StoreUID float64  `json:"store_uid"`
            OnSameStore bool  `json:"on_same_store"`
         
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
    
    // LocationTimingSerializer used by Catalog
    type LocationTimingSerializer struct {

        
            Hour float64  `json:"hour"`
            Minute float64  `json:"minute"`
         
    }
    
    // LocationDayWiseSerializer used by Catalog
    type LocationDayWiseSerializer struct {

        
            Open bool  `json:"open"`
            Closing LocationTimingSerializer  `json:"closing"`
            Weekday string  `json:"weekday"`
            Opening LocationTimingSerializer  `json:"opening"`
         
    }
    
    // LocationManagerSerializer used by Catalog
    type LocationManagerSerializer struct {

        
            MobileNo SellerPhoneNumber  `json:"mobile_no"`
            Email string  `json:"email"`
            Name string  `json:"name"`
         
    }
    
    // GetLocationSerializer used by Catalog
    type GetLocationSerializer struct {

        
            CreatedBy UserSerializer1  `json:"created_by"`
            UID float64  `json:"uid"`
            VerifiedOn string  `json:"verified_on"`
            DisplayName string  `json:"display_name"`
            NotificationEmails []string  `json:"notification_emails"`
            PhoneNumber string  `json:"phone_number"`
            Name string  `json:"name"`
            IntegrationType LocationIntegrationType  `json:"integration_type"`
            Address GetAddressSerializer  `json:"address"`
            Company GetCompanySerializer  `json:"company"`
            ProductReturnConfig ProductReturnConfigSerializer  `json:"product_return_config"`
            GstCredentials InvoiceDetailsSerializer  `json:"gst_credentials"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            CreatedOn string  `json:"created_on"`
            Warnings map[string]interface{}  `json:"warnings"`
            Documents []Document  `json:"documents"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            Timing []LocationDayWiseSerializer  `json:"timing"`
            ModifiedBy UserSerializer1  `json:"modified_by"`
            StoreType string  `json:"store_type"`
            Manager LocationManagerSerializer  `json:"manager"`
            Stage string  `json:"stage"`
            ModifiedOn string  `json:"modified_on"`
            Code string  `json:"code"`
            VerifiedBy UserSerializer1  `json:"verified_by"`
         
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

        
            UID float64  `json:"uid"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Logo string  `json:"logo"`
            IsActive bool  `json:"is_active"`
            AppID string  `json:"app_id"`
            Name string  `json:"name"`
         
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

        
            EffectiveDate string  `json:"effective_date"`
            Enable bool  `json:"enable"`
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
    
    // BusinessCountryInfo used by CompanyProfile
    type BusinessCountryInfo struct {

        
            CountryCode string  `json:"country_code"`
            Country string  `json:"country"`
         
    }
    
    // ContactDetails used by CompanyProfile
    type ContactDetails struct {

        
            Phone []SellerPhoneNumber  `json:"phone"`
            Emails []string  `json:"emails"`
         
    }
    
    // GetCompanyProfileSerializerResponse used by CompanyProfile
    type GetCompanyProfileSerializerResponse struct {

        
            VerifiedBy UserSerializer  `json:"verified_by"`
            Mode string  `json:"mode"`
            Taxes []CompanyTaxesSerializer  `json:"taxes"`
            BusinessInfo string  `json:"business_info"`
            Warnings map[string]interface{}  `json:"warnings"`
            BusinessType string  `json:"business_type"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            BusinessDetails BusinessDetails  `json:"business_details"`
            NotificationEmails []string  `json:"notification_emails"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            UID float64  `json:"uid"`
            CreatedOn string  `json:"created_on"`
            VerifiedOn string  `json:"verified_on"`
            ModifiedOn string  `json:"modified_on"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            Documents []Document  `json:"documents"`
            CreatedBy UserSerializer  `json:"created_by"`
            BusinessCountryInfo BusinessCountryInfo  `json:"business_country_info"`
            FranchiseEnabled bool  `json:"franchise_enabled"`
            ContactDetails ContactDetails  `json:"contact_details"`
            Name string  `json:"name"`
            Stage string  `json:"stage"`
            CompanyType string  `json:"company_type"`
         
    }
    
    // CompanyTaxesSerializer1 used by CompanyProfile
    type CompanyTaxesSerializer1 struct {

        
            EffectiveDate string  `json:"effective_date"`
            Enable bool  `json:"enable"`
            Rate float64  `json:"rate"`
         
    }
    
    // CreateUpdateAddressSerializer used by CompanyProfile
    type CreateUpdateAddressSerializer struct {

        
            State string  `json:"state"`
            CountryCode string  `json:"country_code"`
            Address1 string  `json:"address1"`
            Address2 string  `json:"address2"`
            AddressType string  `json:"address_type"`
            Latitude float64  `json:"latitude"`
            Longitude float64  `json:"longitude"`
            Landmark string  `json:"landmark"`
            City string  `json:"city"`
            Country string  `json:"country"`
            Pincode float64  `json:"pincode"`
         
    }
    
    // UpdateCompany used by CompanyProfile
    type UpdateCompany struct {

        
            Warnings map[string]interface{}  `json:"warnings"`
            BusinessType string  `json:"business_type"`
            ContactDetails ContactDetails  `json:"contact_details"`
            FranchiseEnabled bool  `json:"franchise_enabled"`
            Name string  `json:"name"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Taxes []CompanyTaxesSerializer1  `json:"taxes"`
            RejectReason string  `json:"reject_reason"`
            BusinessDetails BusinessDetails  `json:"business_details"`
            CompanyType string  `json:"company_type"`
            NotificationEmails []string  `json:"notification_emails"`
            BusinessInfo string  `json:"business_info"`
            Addresses []CreateUpdateAddressSerializer  `json:"addresses"`
            Documents []Document  `json:"documents"`
         
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

        
            UID float64  `json:"uid"`
            CompanyDocuments DocumentsObj  `json:"company_documents"`
            Stage string  `json:"stage"`
            StoreDocuments DocumentsObj  `json:"store_documents"`
            Store DocumentsObj  `json:"store"`
            Product DocumentsObj  `json:"product"`
            Brand DocumentsObj  `json:"brand"`
         
    }
    
    // BrandBannerSerializer used by CompanyProfile
    type BrandBannerSerializer struct {

        
            Landscape string  `json:"landscape"`
            Portrait string  `json:"portrait"`
         
    }
    
    // GetBrandResponseSerializer used by CompanyProfile
    type GetBrandResponseSerializer struct {

        
            VerifiedBy UserSerializer  `json:"verified_by"`
            Mode string  `json:"mode"`
            Logo string  `json:"logo"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            Banner BrandBannerSerializer  `json:"banner"`
            Warnings map[string]interface{}  `json:"warnings"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            UID float64  `json:"uid"`
            CreatedOn string  `json:"created_on"`
            Description string  `json:"description"`
            VerifiedOn string  `json:"verified_on"`
            RejectReason string  `json:"reject_reason"`
            ModifiedOn string  `json:"modified_on"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            Synonyms []string  `json:"synonyms"`
            CreatedBy UserSerializer  `json:"created_by"`
            Name string  `json:"name"`
            Stage string  `json:"stage"`
            SlugKey string  `json:"slug_key"`
         
    }
    
    // CreateUpdateBrandRequestSerializer used by CompanyProfile
    type CreateUpdateBrandRequestSerializer struct {

        
            Banner BrandBannerSerializer  `json:"banner"`
            UID float64  `json:"uid"`
            Logo string  `json:"logo"`
            BrandTier string  `json:"brand_tier"`
            Name string  `json:"name"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            CompanyID float64  `json:"company_id"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            Description string  `json:"description"`
            Synonyms []string  `json:"synonyms"`
         
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

        
            CreatedBy UserSerializer  `json:"created_by"`
            BusinessCountryInfo BusinessCountryInfo  `json:"business_country_info"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            UID float64  `json:"uid"`
            MarketChannels []string  `json:"market_channels"`
            CompanyType string  `json:"company_type"`
            BusinessType string  `json:"business_type"`
            Name string  `json:"name"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Stage string  `json:"stage"`
            CreatedOn string  `json:"created_on"`
            VerifiedOn string  `json:"verified_on"`
            RejectReason string  `json:"reject_reason"`
            ModifiedOn string  `json:"modified_on"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            NotificationEmails []string  `json:"notification_emails"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            Details CompanyDetails  `json:"details"`
         
    }
    
    // CompanyBrandSerializer used by CompanyProfile
    type CompanyBrandSerializer struct {

        
            CreatedBy UserSerializer  `json:"created_by"`
            Company CompanySerializer  `json:"company"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            UID float64  `json:"uid"`
            Warnings map[string]interface{}  `json:"warnings"`
            Stage string  `json:"stage"`
            CreatedOn string  `json:"created_on"`
            VerifiedOn string  `json:"verified_on"`
            RejectReason string  `json:"reject_reason"`
            ModifiedOn string  `json:"modified_on"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            Brand GetBrandResponseSerializer  `json:"brand"`
         
    }
    
    // CompanyBrandListSerializer used by CompanyProfile
    type CompanyBrandListSerializer struct {

        
            Page Page  `json:"page"`
            Items []CompanyBrandSerializer  `json:"items"`
         
    }
    
    // CompanyBrandPostRequestSerializer used by CompanyProfile
    type CompanyBrandPostRequestSerializer struct {

        
            Company float64  `json:"company"`
            UID float64  `json:"uid"`
            Brands []float64  `json:"brands"`
         
    }
    
    // HolidayDateSerializer used by CompanyProfile
    type HolidayDateSerializer struct {

        
            EndDate string  `json:"end_date"`
            StartDate string  `json:"start_date"`
         
    }
    
    // HolidaySchemaSerializer used by CompanyProfile
    type HolidaySchemaSerializer struct {

        
            Date HolidayDateSerializer  `json:"date"`
            HolidayType string  `json:"holiday_type"`
            Title string  `json:"title"`
         
    }
    
    // LocationSerializer used by CompanyProfile
    type LocationSerializer struct {

        
            Company float64  `json:"company"`
            UID float64  `json:"uid"`
            Holiday []HolidaySchemaSerializer  `json:"holiday"`
            Warnings map[string]interface{}  `json:"warnings"`
            Code string  `json:"code"`
            Name string  `json:"name"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            Stage string  `json:"stage"`
            StoreType string  `json:"store_type"`
            Timing []LocationDayWiseSerializer  `json:"timing"`
            Address GetAddressSerializer  `json:"address"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ProductReturnConfig ProductReturnConfigSerializer  `json:"product_return_config"`
            Manager LocationManagerSerializer  `json:"manager"`
            NotificationEmails []string  `json:"notification_emails"`
            DisplayName string  `json:"display_name"`
            GstCredentials InvoiceDetailsSerializer  `json:"gst_credentials"`
            Documents []Document  `json:"documents"`
         
    }
    
    // BulkLocationSerializer used by CompanyProfile
    type BulkLocationSerializer struct {

        
            Data []LocationSerializer  `json:"data"`
         
    }
    
    // _ArticleQuery used by CompanyProfile
    type _ArticleQuery struct {

        
            Size string  `json:"size"`
            ItemID float64  `json:"item_id"`
            IgnoredStores []float64  `json:"ignored_stores"`
         
    }
    
    // _ArticleAssignment used by CompanyProfile
    type _ArticleAssignment struct {

        
            Strategy string  `json:"strategy"`
            Level string  `json:"level"`
         
    }
    
    // _AssignStoreArticle used by CompanyProfile
    type _AssignStoreArticle struct {

        
            Meta map[string]interface{}  `json:"meta"`
            Query _ArticleQuery  `json:"query"`
            GroupID string  `json:"group_id"`
            ArticleAssignment _ArticleAssignment  `json:"article_assignment"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // AssignStoreRequestValidator used by CompanyProfile
    type AssignStoreRequestValidator struct {

        
            ChannelType string  `json:"channel_type"`
            CompanyID float64  `json:"company_id"`
            StoreIds []float64  `json:"store_ids"`
            AppID string  `json:"app_id"`
            ChannelIdentifier string  `json:"channel_identifier"`
            Articles []_AssignStoreArticle  `json:"articles"`
            Pincode string  `json:"pincode"`
         
    }
    
    // AssignStoreResponseSerializer used by CompanyProfile
    type AssignStoreResponseSerializer struct {

        
            Meta map[string]interface{}  `json:"meta"`
            UID string  `json:"uid"`
            PreiceEffective float64  `json:"preice_effective"`
            Index float64  `json:"index"`
            StoreID float64  `json:"store_id"`
            CompanyID float64  `json:"company_id"`
            SCity string  `json:"s_city"`
            ArticleAssignment _ArticleAssignment  `json:"article_assignment"`
            ID string  `json:"_id"`
            StorePincode string  `json:"store_pincode"`
            ItemID float64  `json:"item_id"`
            PriceMarked float64  `json:"price_marked"`
            Status bool  `json:"status"`
            Size string  `json:"size"`
            Quantity float64  `json:"quantity"`
         
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
    

    
    // DisplayMetaDict used by Cart
    type DisplayMetaDict struct {

        
            Subtitle string  `json:"subtitle"`
            Title string  `json:"title"`
         
    }
    
    // DisplayMeta used by Cart
    type DisplayMeta struct {

        
            Apply DisplayMetaDict  `json:"apply"`
            Remove DisplayMetaDict  `json:"remove"`
            Description string  `json:"description"`
            Title string  `json:"title"`
            Auto DisplayMetaDict  `json:"auto"`
            Subtitle string  `json:"subtitle"`
         
    }
    
    // CouponDateMeta used by Cart
    type CouponDateMeta struct {

        
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
         
    }
    
    // CouponAuthor used by Cart
    type CouponAuthor struct {

        
            CreatedBy string  `json:"created_by"`
            ModifiedBy string  `json:"modified_by"`
         
    }
    
    // CouponSchedule used by Cart
    type CouponSchedule struct {

        
            End string  `json:"end"`
            Start string  `json:"start"`
            NextSchedule []map[string]interface{}  `json:"next_schedule"`
            Duration float64  `json:"duration"`
            Cron string  `json:"cron"`
         
    }
    
    // RuleDefinition used by Cart
    type RuleDefinition struct {

        
            IsExact bool  `json:"is_exact"`
            CalculateOn string  `json:"calculate_on"`
            Scope []string  `json:"scope"`
            ValueType string  `json:"value_type"`
            AutoApply bool  `json:"auto_apply"`
            CurrencyCode string  `json:"currency_code"`
            Type string  `json:"type"`
            ApplicableOn string  `json:"applicable_on"`
         
    }
    
    // Ownership used by Cart
    type Ownership struct {

        
            PayableBy string  `json:"payable_by"`
            PayableCategory string  `json:"payable_category"`
         
    }
    
    // Rule used by Cart
    type Rule struct {

        
            Value float64  `json:"value"`
            Min float64  `json:"min"`
            Key float64  `json:"key"`
            Max float64  `json:"max"`
            DiscountQty float64  `json:"discount_qty"`
         
    }
    
    // Validity used by Cart
    type Validity struct {

        
            Priority float64  `json:"priority"`
         
    }
    
    // Validation used by Cart
    type Validation struct {

        
            Anonymous bool  `json:"anonymous"`
            AppID []string  `json:"app_id"`
            UserRegisteredAfter string  `json:"user_registered_after"`
         
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
    
    // PaymentAllowValue used by Cart
    type PaymentAllowValue struct {

        
            Max float64  `json:"max"`
         
    }
    
    // PaymentModes used by Cart
    type PaymentModes struct {

        
            Networks []string  `json:"networks"`
            Types []string  `json:"types"`
            Uses PaymentAllowValue  `json:"uses"`
            Codes []string  `json:"codes"`
         
    }
    
    // PostOrder used by Cart
    type PostOrder struct {

        
            CancellationAllowed bool  `json:"cancellation_allowed"`
            ReturnAllowed bool  `json:"return_allowed"`
         
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
    
    // Restrictions used by Cart
    type Restrictions struct {

        
            CouponAllowed bool  `json:"coupon_allowed"`
            BulkBundle BulkBundleRestriction  `json:"bulk_bundle"`
            PriceRange PriceRange  `json:"price_range"`
            Payments map[string]PaymentModes  `json:"payments"`
            UserGroups []float64  `json:"user_groups"`
            Platforms []string  `json:"platforms"`
            UserType string  `json:"user_type"`
            PostOrder PostOrder  `json:"post_order"`
            Uses UsesRestriction  `json:"uses"`
            OrderingStores []float64  `json:"ordering_stores"`
         
    }
    
    // State used by Cart
    type State struct {

        
            IsPublic bool  `json:"is_public"`
            IsArchived bool  `json:"is_archived"`
            IsDisplay bool  `json:"is_display"`
         
    }
    
    // CouponAction used by Cart
    type CouponAction struct {

        
            TxnMode string  `json:"txn_mode"`
            ActionDate string  `json:"action_date"`
         
    }
    
    // CouponAdd used by Cart
    type CouponAdd struct {

        
            DisplayMeta DisplayMeta  `json:"display_meta"`
            DateMeta CouponDateMeta  `json:"date_meta"`
            Author CouponAuthor  `json:"author"`
            Schedule CouponSchedule  `json:"_schedule"`
            RuleDefinition RuleDefinition  `json:"rule_definition"`
            Ownership Ownership  `json:"ownership"`
            Rule []Rule  `json:"rule"`
            Validity Validity  `json:"validity"`
            Validation Validation  `json:"validation"`
            Tags []string  `json:"tags"`
            Restrictions Restrictions  `json:"restrictions"`
            Code string  `json:"code"`
            TypeSlug string  `json:"type_slug"`
            State State  `json:"state"`
            Identifiers Identifier  `json:"identifiers"`
            Action CouponAction  `json:"action"`
         
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

        
            DisplayMeta DisplayMeta  `json:"display_meta"`
            DateMeta CouponDateMeta  `json:"date_meta"`
            Author CouponAuthor  `json:"author"`
            Schedule CouponSchedule  `json:"_schedule"`
            RuleDefinition RuleDefinition  `json:"rule_definition"`
            Ownership Ownership  `json:"ownership"`
            Rule []Rule  `json:"rule"`
            Validity Validity  `json:"validity"`
            Validation Validation  `json:"validation"`
            Tags []string  `json:"tags"`
            Restrictions Restrictions  `json:"restrictions"`
            Code string  `json:"code"`
            TypeSlug string  `json:"type_slug"`
            State State  `json:"state"`
            Identifiers Identifier  `json:"identifiers"`
            Action CouponAction  `json:"action"`
         
    }
    
    // CouponPartialUpdate used by Cart
    type CouponPartialUpdate struct {

        
            Archive bool  `json:"archive"`
            Schedule CouponSchedule  `json:"schedule"`
         
    }
    
    // PromotionSchedule used by Cart
    type PromotionSchedule struct {

        
            End string  `json:"end"`
            Start string  `json:"start"`
            Published bool  `json:"published"`
            NextSchedule []map[string]interface{}  `json:"next_schedule"`
            Duration float64  `json:"duration"`
            Cron string  `json:"cron"`
         
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

        
            Codes []string  `json:"codes"`
            Uses PaymentAllowValue1  `json:"uses"`
            Type string  `json:"type"`
         
    }
    
    // PostOrder1 used by Cart
    type PostOrder1 struct {

        
            CancellationAllowed bool  `json:"cancellation_allowed"`
            ReturnAllowed bool  `json:"return_allowed"`
         
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
    
    // Restrictions1 used by Cart
    type Restrictions1 struct {

        
            AnonymousUsers bool  `json:"anonymous_users"`
            Payments []PromotionPaymentModes  `json:"payments"`
            UserGroups []float64  `json:"user_groups"`
            Platforms []string  `json:"platforms"`
            PostOrder PostOrder1  `json:"post_order"`
            Uses UsesRestriction1  `json:"uses"`
            UserRegistered UserRegistered  `json:"user_registered"`
            OrderQuantity float64  `json:"order_quantity"`
            UserID []string  `json:"user_id"`
         
    }
    
    // CompareObject used by Cart
    type CompareObject struct {

        
            GreaterThanEquals float64  `json:"greater_than_equals"`
            LessThanEquals float64  `json:"less_than_equals"`
            Equals float64  `json:"equals"`
            GreaterThan float64  `json:"greater_than"`
            LessThan float64  `json:"less_than"`
         
    }
    
    // ItemCriteria used by Cart
    type ItemCriteria struct {

        
            ItemExcludeID []float64  `json:"item_exclude_id"`
            CartUniqueItemQuantity CompareObject  `json:"cart_unique_item_quantity"`
            AvailableZones []string  `json:"available_zones"`
            ItemExcludeSku []string  `json:"item_exclude_sku"`
            ItemID []float64  `json:"item_id"`
            CartQuantity CompareObject  `json:"cart_quantity"`
            ItemBrand []float64  `json:"item_brand"`
            BuyRules []string  `json:"buy_rules"`
            ProductTags []string  `json:"product_tags"`
            ItemStore []float64  `json:"item_store"`
            ItemExcludeStore []float64  `json:"item_exclude_store"`
            ItemSku []string  `json:"item_sku"`
            ItemExcludeBrand []float64  `json:"item_exclude_brand"`
            ItemCompany []float64  `json:"item_company"`
            CartUniqueItemAmount CompareObject  `json:"cart_unique_item_amount"`
            ItemExcludeCompany []float64  `json:"item_exclude_company"`
            ItemCategory []float64  `json:"item_category"`
            ItemSize []string  `json:"item_size"`
            CartTotal CompareObject  `json:"cart_total"`
            AllItems bool  `json:"all_items"`
            ItemExcludeCategory []float64  `json:"item_exclude_category"`
         
    }
    
    // DisplayMeta1 used by Cart
    type DisplayMeta1 struct {

        
            OfferLabel string  `json:"offer_label"`
            Description string  `json:"description"`
            Name string  `json:"name"`
            OfferText string  `json:"offer_text"`
         
    }
    
    // PromotionAuthor used by Cart
    type PromotionAuthor struct {

        
            CreatedBy string  `json:"created_by"`
            ModifiedBy string  `json:"modified_by"`
         
    }
    
    // DiscountOffer used by Cart
    type DiscountOffer struct {

        
            DiscountPrice float64  `json:"discount_price"`
            DiscountPercentage float64  `json:"discount_percentage"`
            ApportionDiscount bool  `json:"apportion_discount"`
            PartialCanRet bool  `json:"partial_can_ret"`
            MaxOfferQuantity float64  `json:"max_offer_quantity"`
            MinOfferQuantity float64  `json:"min_offer_quantity"`
            MaxDiscountAmount float64  `json:"max_discount_amount"`
            DiscountAmount float64  `json:"discount_amount"`
            Code string  `json:"code"`
            MaxUsagePerTransaction float64  `json:"max_usage_per_transaction"`
         
    }
    
    // DiscountRule used by Cart
    type DiscountRule struct {

        
            DiscountType string  `json:"discount_type"`
            Offer DiscountOffer  `json:"offer"`
            ItemCriteria ItemCriteria  `json:"item_criteria"`
            BuyCondition string  `json:"buy_condition"`
         
    }
    
    // PromotionAction used by Cart
    type PromotionAction struct {

        
            ActionType string  `json:"action_type"`
            ActionDate string  `json:"action_date"`
         
    }
    
    // Visibility used by Cart
    type Visibility struct {

        
            CouponList bool  `json:"coupon_list"`
            Pdp bool  `json:"pdp"`
         
    }
    
    // PromotionDateMeta used by Cart
    type PromotionDateMeta struct {

        
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
         
    }
    
    // PromotionListItem used by Cart
    type PromotionListItem struct {

        
            Schedule PromotionSchedule  `json:"_schedule"`
            Ownership Ownership1  `json:"ownership"`
            PromoGroup string  `json:"promo_group"`
            Restrictions Restrictions1  `json:"restrictions"`
            Stackable bool  `json:"stackable"`
            BuyRules map[string]ItemCriteria  `json:"buy_rules"`
            DisplayMeta DisplayMeta1  `json:"display_meta"`
            Mode string  `json:"mode"`
            Currency string  `json:"currency"`
            ApplyExclusive string  `json:"apply_exclusive"`
            ApplicationID string  `json:"application_id"`
            ApplyPriority float64  `json:"apply_priority"`
            CalculateOn string  `json:"calculate_on"`
            Author PromotionAuthor  `json:"author"`
            DiscountRules []DiscountRule  `json:"discount_rules"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Code string  `json:"code"`
            PostOrderAction PromotionAction  `json:"post_order_action"`
            Visiblility Visibility  `json:"visiblility"`
            PromotionType string  `json:"promotion_type"`
            DateMeta PromotionDateMeta  `json:"date_meta"`
            ApplyAllDiscount bool  `json:"apply_all_discount"`
         
    }
    
    // PromotionsResponse used by Cart
    type PromotionsResponse struct {

        
            Items PromotionListItem  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // PromotionAdd used by Cart
    type PromotionAdd struct {

        
            Schedule PromotionSchedule  `json:"_schedule"`
            Ownership Ownership1  `json:"ownership"`
            PromoGroup string  `json:"promo_group"`
            Restrictions Restrictions1  `json:"restrictions"`
            Stackable bool  `json:"stackable"`
            BuyRules map[string]ItemCriteria  `json:"buy_rules"`
            DisplayMeta DisplayMeta1  `json:"display_meta"`
            Mode string  `json:"mode"`
            Currency string  `json:"currency"`
            ApplyExclusive string  `json:"apply_exclusive"`
            ApplicationID string  `json:"application_id"`
            ApplyPriority float64  `json:"apply_priority"`
            CalculateOn string  `json:"calculate_on"`
            Author PromotionAuthor  `json:"author"`
            DiscountRules []DiscountRule  `json:"discount_rules"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Code string  `json:"code"`
            PostOrderAction PromotionAction  `json:"post_order_action"`
            Visiblility Visibility  `json:"visiblility"`
            PromotionType string  `json:"promotion_type"`
            DateMeta PromotionDateMeta  `json:"date_meta"`
            ApplyAllDiscount bool  `json:"apply_all_discount"`
         
    }
    
    // PromotionUpdate used by Cart
    type PromotionUpdate struct {

        
            Schedule PromotionSchedule  `json:"_schedule"`
            Ownership Ownership1  `json:"ownership"`
            PromoGroup string  `json:"promo_group"`
            Restrictions Restrictions1  `json:"restrictions"`
            Stackable bool  `json:"stackable"`
            BuyRules map[string]ItemCriteria  `json:"buy_rules"`
            DisplayMeta DisplayMeta1  `json:"display_meta"`
            Mode string  `json:"mode"`
            Currency string  `json:"currency"`
            ApplyExclusive string  `json:"apply_exclusive"`
            ApplicationID string  `json:"application_id"`
            ApplyPriority float64  `json:"apply_priority"`
            CalculateOn string  `json:"calculate_on"`
            Author PromotionAuthor  `json:"author"`
            DiscountRules []DiscountRule  `json:"discount_rules"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Code string  `json:"code"`
            PostOrderAction PromotionAction  `json:"post_order_action"`
            Visiblility Visibility  `json:"visiblility"`
            PromotionType string  `json:"promotion_type"`
            DateMeta PromotionDateMeta  `json:"date_meta"`
            ApplyAllDiscount bool  `json:"apply_all_discount"`
         
    }
    
    // PromotionPartialUpdate used by Cart
    type PromotionPartialUpdate struct {

        
            Archive bool  `json:"archive"`
            Schedule PromotionSchedule  `json:"schedule"`
         
    }
    
    // ActivePromosResponse used by Cart
    type ActivePromosResponse struct {

        
            Example string  `json:"example"`
            ModifiedOn string  `json:"modified_on"`
            EntityType string  `json:"entity_type"`
            Description string  `json:"description"`
            CreatedOn string  `json:"created_on"`
            Title string  `json:"title"`
            Type string  `json:"type"`
            IsHidden bool  `json:"is_hidden"`
            Subtitle string  `json:"subtitle"`
            EntitySlug string  `json:"entity_slug"`
         
    }
    
    // CartItem used by Cart
    type CartItem struct {

        
            Quantity float64  `json:"quantity"`
            Size string  `json:"size"`
            ProductID string  `json:"product_id"`
         
    }
    
    // OpenapiCartDetailsRequest used by Cart
    type OpenapiCartDetailsRequest struct {

        
            CartItems CartItem  `json:"cart_items"`
         
    }
    
    // ProductAvailability used by Cart
    type ProductAvailability struct {

        
            Sizes []string  `json:"sizes"`
            OtherStoreQuantity float64  `json:"other_store_quantity"`
            Deliverable bool  `json:"deliverable"`
            OutOfStock bool  `json:"out_of_stock"`
            IsValid bool  `json:"is_valid"`
         
    }
    
    // PromoMeta used by Cart
    type PromoMeta struct {

        
            Message string  `json:"message"`
         
    }
    
    // ProductPrice used by Cart
    type ProductPrice struct {

        
            CurrencyCode string  `json:"currency_code"`
            CurrencySymbol string  `json:"currency_symbol"`
            Effective float64  `json:"effective"`
            Selling float64  `json:"selling"`
            AddOn float64  `json:"add_on"`
            Marked float64  `json:"marked"`
         
    }
    
    // ProductPriceInfo used by Cart
    type ProductPriceInfo struct {

        
            Converted ProductPrice  `json:"converted"`
            Base ProductPrice  `json:"base"`
         
    }
    
    // DiscountRulesApp used by Cart
    type DiscountRulesApp struct {

        
            RawOffer map[string]interface{}  `json:"raw_offer"`
            MatchedBuyRules []string  `json:"matched_buy_rules"`
            Offer map[string]interface{}  `json:"offer"`
            ItemCriteria map[string]interface{}  `json:"item_criteria"`
         
    }
    
    // FreeGiftItem used by Cart
    type FreeGiftItem struct {

        
            ItemPriceDetails map[string]interface{}  `json:"item_price_details"`
            ItemImagesURL []string  `json:"item_images_url"`
            ItemSlug string  `json:"item_slug"`
            ItemName string  `json:"item_name"`
            ItemID float64  `json:"item_id"`
            ItemBrandName string  `json:"item_brand_name"`
         
    }
    
    // AppliedFreeArticles used by Cart
    type AppliedFreeArticles struct {

        
            ArticleID string  `json:"article_id"`
            Quantity float64  `json:"quantity"`
            FreeGiftItemDetails FreeGiftItem  `json:"free_gift_item_details"`
            ParentItemIdentifier string  `json:"parent_item_identifier"`
         
    }
    
    // AppliedPromotion used by Cart
    type AppliedPromotion struct {

        
            MrpPromotion bool  `json:"mrp_promotion"`
            BuyRules []BuyRules  `json:"buy_rules"`
            OfferText string  `json:"offer_text"`
            Amount float64  `json:"amount"`
            DiscountRules []DiscountRulesApp  `json:"discount_rules"`
            AppliedFreeArticles []AppliedFreeArticles  `json:"applied_free_articles"`
            PromotionType string  `json:"promotion_type"`
            ArticleQuantity float64  `json:"article_quantity"`
            PromoID string  `json:"promo_id"`
            PromotionGroup string  `json:"promotion_group"`
            PromotionName string  `json:"promotion_name"`
         
    }
    
    // BasePrice used by Cart
    type BasePrice struct {

        
            CurrencyCode string  `json:"currency_code"`
            Effective float64  `json:"effective"`
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
    
    // ProductArticle used by Cart
    type ProductArticle struct {

        
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            Price ArticlePriceInfo  `json:"price"`
            Seller BaseInfo  `json:"seller"`
            UID string  `json:"uid"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            Store BaseInfo  `json:"store"`
            Size string  `json:"size"`
            Type string  `json:"type"`
            Quantity float64  `json:"quantity"`
            ProductGroupTags []string  `json:"product_group_tags"`
         
    }
    
    // CartProductIdentifer used by Cart
    type CartProductIdentifer struct {

        
            Identifier string  `json:"identifier"`
         
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

        
            Name string  `json:"name"`
            Brand BaseInfo  `json:"brand"`
            Images []ProductImage  `json:"images"`
            Categories []CategoryInfo  `json:"categories"`
            NetQuantity NetQuantity  `json:"net_quantity"`
            UID float64  `json:"uid"`
            Type string  `json:"type"`
            Slug string  `json:"slug"`
            Action ProductAction  `json:"action"`
         
    }
    
    // CartProductInfo used by Cart
    type CartProductInfo struct {

        
            Availability ProductAvailability  `json:"availability"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            PromoMeta PromoMeta  `json:"promo_meta"`
            Price ProductPriceInfo  `json:"price"`
            PromotionsApplied []AppliedPromotion  `json:"promotions_applied"`
            Discount string  `json:"discount"`
            CouponMessage string  `json:"coupon_message"`
            Message string  `json:"message"`
            Article ProductArticle  `json:"article"`
            PricePerUnit ProductPriceInfo  `json:"price_per_unit"`
            IsSet bool  `json:"is_set"`
            Quantity float64  `json:"quantity"`
            Key string  `json:"key"`
            BulkOffer map[string]interface{}  `json:"bulk_offer"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
            Product CartProduct  `json:"product"`
         
    }
    
    // DisplayBreakup used by Cart
    type DisplayBreakup struct {

        
            Value float64  `json:"value"`
            CurrencyCode string  `json:"currency_code"`
            CurrencySymbol string  `json:"currency_symbol"`
            Message []string  `json:"message"`
            Display string  `json:"display"`
            Key string  `json:"key"`
         
    }
    
    // LoyaltyPoints used by Cart
    type LoyaltyPoints struct {

        
            Applicable float64  `json:"applicable"`
            IsApplied bool  `json:"is_applied"`
            Description string  `json:"description"`
            Total float64  `json:"total"`
         
    }
    
    // RawBreakup used by Cart
    type RawBreakup struct {

        
            Subtotal float64  `json:"subtotal"`
            Coupon float64  `json:"coupon"`
            CodCharge float64  `json:"cod_charge"`
            MrpTotal float64  `json:"mrp_total"`
            FyndCash float64  `json:"fynd_cash"`
            ConvenienceFee float64  `json:"convenience_fee"`
            Vog float64  `json:"vog"`
            YouSaved float64  `json:"you_saved"`
            GstCharges float64  `json:"gst_charges"`
            Discount float64  `json:"discount"`
            DeliveryCharge float64  `json:"delivery_charge"`
            Total float64  `json:"total"`
         
    }
    
    // CouponBreakup used by Cart
    type CouponBreakup struct {

        
            Value float64  `json:"value"`
            CouponType string  `json:"coupon_type"`
            MaxDiscountValue float64  `json:"max_discount_value"`
            Message string  `json:"message"`
            UID string  `json:"uid"`
            IsApplied bool  `json:"is_applied"`
            Description string  `json:"description"`
            MinimumCartValue float64  `json:"minimum_cart_value"`
            CouponValue float64  `json:"coupon_value"`
            Title string  `json:"title"`
            Type string  `json:"type"`
            SubTitle string  `json:"sub_title"`
            Code string  `json:"code"`
         
    }
    
    // CartBreakup used by Cart
    type CartBreakup struct {

        
            Display []DisplayBreakup  `json:"display"`
            LoyaltyPoints LoyaltyPoints  `json:"loyalty_points"`
            Raw RawBreakup  `json:"raw"`
            Coupon CouponBreakup  `json:"coupon"`
         
    }
    
    // OpenapiCartDetailsResponse used by Cart
    type OpenapiCartDetailsResponse struct {

        
            Message string  `json:"message"`
            Items []CartProductInfo  `json:"items"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            IsValid bool  `json:"is_valid"`
         
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
            Name string  `json:"name"`
            CountryCode string  `json:"country_code"`
            Area string  `json:"area"`
            AddressType string  `json:"address_type"`
            Email string  `json:"email"`
            Landmark string  `json:"landmark"`
            AreaCodeSlug string  `json:"area_code_slug"`
            Phone float64  `json:"phone"`
            City string  `json:"city"`
            Country string  `json:"country"`
            Meta map[string]interface{}  `json:"meta"`
            State string  `json:"state"`
            Address string  `json:"address"`
            AreaCode string  `json:"area_code"`
         
    }
    
    // OpenApiCartServiceabilityRequest used by Cart
    type OpenApiCartServiceabilityRequest struct {

        
            ShippingAddress ShippingAddress  `json:"shipping_address"`
            CartItems CartItem  `json:"cart_items"`
         
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
            BreakupValues CartBreakup  `json:"breakup_values"`
            Message string  `json:"message"`
            Items []CartProductInfo  `json:"items"`
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

        
            EmployeeDiscount float64  `json:"employee_discount"`
            AmountPaid float64  `json:"amount_paid"`
            Discount float64  `json:"discount"`
            PriceMarked float64  `json:"price_marked"`
            PaymentMethods []MultiTenderPaymentMethod  `json:"payment_methods"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            CodCharges float64  `json:"cod_charges"`
            Meta CartItemMeta  `json:"meta"`
            ProductID float64  `json:"product_id"`
            Size string  `json:"size"`
            DeliveryCharges float64  `json:"delivery_charges"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            PriceEffective float64  `json:"price_effective"`
            Quantity float64  `json:"quantity"`
            CashbackApplied float64  `json:"cashback_applied"`
            Files []OpenApiFiles  `json:"files"`
         
    }
    
    // OpenApiPlatformCheckoutReq used by Cart
    type OpenApiPlatformCheckoutReq struct {

        
            Coupon string  `json:"coupon"`
            EmployeeDiscount map[string]interface{}  `json:"employee_discount"`
            CartItems []OpenApiOrderItem  `json:"cart_items"`
            ShippingAddress ShippingAddress  `json:"shipping_address"`
            PaymentMethods []MultiTenderPaymentMethod  `json:"payment_methods"`
            CouponValue float64  `json:"coupon_value"`
            CodCharges float64  `json:"cod_charges"`
            DeliveryCharges float64  `json:"delivery_charges"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            OrderID string  `json:"order_id"`
            CashbackApplied float64  `json:"cashback_applied"`
            Gstin string  `json:"gstin"`
            CouponCode string  `json:"coupon_code"`
            CurrencyCode string  `json:"currency_code"`
            CartValue float64  `json:"cart_value"`
            Comment string  `json:"comment"`
            BillingAddress ShippingAddress  `json:"billing_address"`
            Files []OpenApiFiles  `json:"files"`
            PaymentMode string  `json:"payment_mode"`
         
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

        
            FcIndexMap []float64  `json:"fc_index_map"`
            AppID string  `json:"app_id"`
            IsArchive bool  `json:"is_archive"`
            BuyNow bool  `json:"buy_now"`
            IsActive bool  `json:"is_active"`
            PaymentMethods []map[string]interface{}  `json:"payment_methods"`
            OrderID string  `json:"order_id"`
            LastModified string  `json:"last_modified"`
            Promotion map[string]interface{}  `json:"promotion"`
            PickUpCustomerDetails map[string]interface{}  `json:"pick_up_customer_details"`
            Payments map[string]interface{}  `json:"payments"`
            Cashback map[string]interface{}  `json:"cashback"`
            UID float64  `json:"uid"`
            Meta map[string]interface{}  `json:"meta"`
            BulkCouponDiscount float64  `json:"bulk_coupon_discount"`
            PaymentMode string  `json:"payment_mode"`
            IsDefault bool  `json:"is_default"`
            Coupon map[string]interface{}  `json:"coupon"`
            ExpireAt string  `json:"expire_at"`
            CodCharges map[string]interface{}  `json:"cod_charges"`
            MergeQty bool  `json:"merge_qty"`
            DeliveryCharges map[string]interface{}  `json:"delivery_charges"`
            Discount float64  `json:"discount"`
            UserID string  `json:"user_id"`
            Gstin string  `json:"gstin"`
            Articles []map[string]interface{}  `json:"articles"`
            Shipments []map[string]interface{}  `json:"shipments"`
            ID string  `json:"_id"`
            CartValue float64  `json:"cart_value"`
            FyndCredits map[string]interface{}  `json:"fynd_credits"`
            Comment string  `json:"comment"`
            CreatedOn string  `json:"created_on"`
            CheckoutMode string  `json:"checkout_mode"`
         
    }
    
    // AbandonedCartResponse used by Cart
    type AbandonedCartResponse struct {

        
            Page Page  `json:"page"`
            Result map[string]interface{}  `json:"result"`
            Message string  `json:"message"`
            Items []AbandonedCart  `json:"items"`
            Success bool  `json:"success"`
         
    }
    
    // CartCurrency used by Cart
    type CartCurrency struct {

        
            Code string  `json:"code"`
            Symbol string  `json:"symbol"`
         
    }
    
    // CartDetailResponse used by Cart
    type CartDetailResponse struct {

        
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            CouponText string  `json:"coupon_text"`
            ID string  `json:"id"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            BuyNow bool  `json:"buy_now"`
            Currency CartCurrency  `json:"currency"`
            Comment string  `json:"comment"`
            Message string  `json:"message"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            CheckoutMode string  `json:"checkout_mode"`
            Items []CartProductInfo  `json:"items"`
            LastModified string  `json:"last_modified"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            Gstin string  `json:"gstin"`
            IsValid bool  `json:"is_valid"`
         
    }
    
    // AddProductCart used by Cart
    type AddProductCart struct {

        
            ArticleAssignment map[string]interface{}  `json:"article_assignment"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            ArticleID string  `json:"article_id"`
            StoreID float64  `json:"store_id"`
            Pos bool  `json:"pos"`
            ItemSize string  `json:"item_size"`
            ItemID float64  `json:"item_id"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            Display string  `json:"display"`
            Quantity float64  `json:"quantity"`
            SellerID float64  `json:"seller_id"`
            ProductGroupTags []string  `json:"product_group_tags"`
         
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

        
            ItemIndex float64  `json:"item_index"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            ArticleID string  `json:"article_id"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            ItemID float64  `json:"item_id"`
            ItemSize string  `json:"item_size"`
            Quantity float64  `json:"quantity"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
         
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
    
