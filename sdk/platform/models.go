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
    
    // TicketCategory used by Lead
    type TicketCategory struct {

        
            Key string  `json:"key"`
            Display string  `json:"display"`
            Form CustomForm  `json:"form"`
            SubCategories []TicketSubCategory  `json:"sub_categories"`
            FeedbackForm TicketFeedbackForm  `json:"feedback_form"`
         
    }
    
    // TicketSubCategory used by Lead
    type TicketSubCategory struct {

        
            Key string  `json:"key"`
            Display string  `json:"display"`
         
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
            ShowSupportDris bool  `json:"show_support_dris"`
            Integration map[string]interface{}  `json:"integration"`
         
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
            TicketID string  `json:"ticket_id"`
            Category TicketCategory  `json:"category"`
            SubCategory TicketSubCategory  `json:"sub_category"`
            Source TicketSourceEnum  `json:"source"`
            Status Status  `json:"status"`
            Priority Priority  `json:"priority"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            AssignedTo map[string]interface{}  `json:"assigned_to"`
            Tags []string  `json:"tags"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            IsFeedbackPending bool  `json:"is_feedback_pending"`
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

        
            Created bool  `json:"created"`
            Success bool  `json:"success"`
            DisplayFields []string  `json:"display_fields"`
            AppID string  `json:"app_id"`
            ExcludedFields []string  `json:"excluded_fields"`
            Aggregators []map[string]interface{}  `json:"aggregators"`
         
    }
    
    // ErrorCodeDescription used by Payment
    type ErrorCodeDescription struct {

        
            Success bool  `json:"success"`
            Description string  `json:"description"`
            Code string  `json:"code"`
         
    }
    
    // PaymentGatewayConfig used by Payment
    type PaymentGatewayConfig struct {

        
            IsActive bool  `json:"is_active"`
            Key string  `json:"key"`
            ConfigType string  `json:"config_type"`
            MerchantSalt string  `json:"merchant_salt"`
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
    
    // PaymentModeLogo used by Payment
    type PaymentModeLogo struct {

        
            Large string  `json:"large"`
            Small string  `json:"small"`
         
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

        
            CardBrandImage string  `json:"card_brand_image"`
            CardIsin string  `json:"card_isin"`
            ExpMonth float64  `json:"exp_month"`
            Expired bool  `json:"expired"`
            CardType string  `json:"card_type"`
            DisplayName string  `json:"display_name"`
            RetryCount float64  `json:"retry_count"`
            Timeout float64  `json:"timeout"`
            CardName string  `json:"card_name"`
            IntentApp []IntentApp  `json:"intent_app"`
            CardBrand string  `json:"card_brand"`
            CardReference string  `json:"card_reference"`
            Nickname string  `json:"nickname"`
            AggregatorName string  `json:"aggregator_name"`
            DisplayPriority float64  `json:"display_priority"`
            CardNumber string  `json:"card_number"`
            MerchantCode string  `json:"merchant_code"`
            CardToken string  `json:"card_token"`
            Name string  `json:"name"`
            CardIssuer string  `json:"card_issuer"`
            IntentAppErrorList []string  `json:"intent_app_error_list"`
            CompliantWithTokenisationGuidelines bool  `json:"compliant_with_tokenisation_guidelines"`
            CardID string  `json:"card_id"`
            ExpYear float64  `json:"exp_year"`
            IntentFlow bool  `json:"intent_flow"`
            IntentAppErrorDictList []IntentAppErrorList  `json:"intent_app_error_dict_list"`
            LogoURL PaymentModeLogo  `json:"logo_url"`
            CardFingerprint string  `json:"card_fingerprint"`
            Code string  `json:"code"`
            FyndVpa string  `json:"fynd_vpa"`
         
    }
    
    // RootPaymentMode used by Payment
    type RootPaymentMode struct {

        
            IsPayByCardPl bool  `json:"is_pay_by_card_pl"`
            AddCardEnabled bool  `json:"add_card_enabled"`
            SaveCard bool  `json:"save_card"`
            List []PaymentModeList  `json:"list"`
            Name string  `json:"name"`
            DisplayName string  `json:"display_name"`
            AnonymousEnable bool  `json:"anonymous_enable"`
            AggregatorName string  `json:"aggregator_name"`
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

        
            IsActive bool  `json:"is_active"`
            IsDefault bool  `json:"is_default"`
            TransferType string  `json:"transfer_type"`
            Customers map[string]interface{}  `json:"customers"`
            UniqueTransferNo map[string]interface{}  `json:"unique_transfer_no"`
            MoreAttributes map[string]interface{}  `json:"more_attributes"`
            PayoutsAggregators []map[string]interface{}  `json:"payouts_aggregators"`
         
    }
    
    // PayoutBankDetails used by Payment
    type PayoutBankDetails struct {

        
            AccountType string  `json:"account_type"`
            Country string  `json:"country"`
            BranchName string  `json:"branch_name"`
            City string  `json:"city"`
            BankName string  `json:"bank_name"`
            AccountNo string  `json:"account_no"`
            State string  `json:"state"`
            AccountHolder string  `json:"account_holder"`
            Pincode float64  `json:"pincode"`
            IfscCode string  `json:"ifsc_code"`
         
    }
    
    // PayoutRequest used by Payment
    type PayoutRequest struct {

        
            IsActive bool  `json:"is_active"`
            BankDetails PayoutBankDetails  `json:"bank_details"`
            TransferType string  `json:"transfer_type"`
            Aggregator string  `json:"aggregator"`
            UniqueExternalID string  `json:"unique_external_id"`
            Users map[string]interface{}  `json:"users"`
         
    }
    
    // PayoutResponse used by Payment
    type PayoutResponse struct {

        
            IsActive bool  `json:"is_active"`
            BankDetails map[string]interface{}  `json:"bank_details"`
            PaymentStatus string  `json:"payment_status"`
            TransferType string  `json:"transfer_type"`
            Created bool  `json:"created"`
            Aggregator string  `json:"aggregator"`
            Success bool  `json:"success"`
            UniqueTransferNo string  `json:"unique_transfer_no"`
            Users map[string]interface{}  `json:"users"`
            Payouts map[string]interface{}  `json:"payouts"`
         
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
            IsVerifiedFlag bool  `json:"is_verified_flag"`
            Data map[string]interface{}  `json:"data"`
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

        
            BranchName string  `json:"branch_name"`
            BankName string  `json:"bank_name"`
            AccountNo string  `json:"account_no"`
            AccountHolder string  `json:"account_holder"`
            IfscCode string  `json:"ifsc_code"`
         
    }
    
    // AddBeneficiaryDetailsOTPRequest used by Payment
    type AddBeneficiaryDetailsOTPRequest struct {

        
            Details BankDetailsForOTP  `json:"details"`
            OrderID string  `json:"order_id"`
         
    }
    
    // IfscCodeResponse used by Payment
    type IfscCodeResponse struct {

        
            Success bool  `json:"success"`
            BankName string  `json:"bank_name"`
            BranchName string  `json:"branch_name"`
         
    }
    
    // OrderBeneficiaryDetails used by Payment
    type OrderBeneficiaryDetails struct {

        
            ID float64  `json:"id"`
            CreatedOn string  `json:"created_on"`
            DisplayName string  `json:"display_name"`
            Comment string  `json:"comment"`
            Mobile string  `json:"mobile"`
            AccountNo string  `json:"account_no"`
            AccountHolder string  `json:"account_holder"`
            IfscCode string  `json:"ifsc_code"`
            BeneficiaryID string  `json:"beneficiary_id"`
            BranchName string  `json:"branch_name"`
            Email string  `json:"email"`
            BankName string  `json:"bank_name"`
            Subtitle string  `json:"subtitle"`
            Title string  `json:"title"`
            Address string  `json:"address"`
            IsActive bool  `json:"is_active"`
            ModifiedOn string  `json:"modified_on"`
            DelightsUserName string  `json:"delights_user_name"`
            TransferMode string  `json:"transfer_mode"`
         
    }
    
    // OrderBeneficiaryResponse used by Payment
    type OrderBeneficiaryResponse struct {

        
            ShowBeneficiaryDetails bool  `json:"show_beneficiary_details"`
            Beneficiaries []OrderBeneficiaryDetails  `json:"beneficiaries"`
         
    }
    
    // MultiTenderPaymentMeta used by Payment
    type MultiTenderPaymentMeta struct {

        
            CurrentStatus string  `json:"current_status"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            PaymentGateway string  `json:"payment_gateway"`
            OrderID string  `json:"order_id"`
            PaymentID string  `json:"payment_id"`
         
    }
    
    // MultiTenderPaymentMethod used by Payment
    type MultiTenderPaymentMethod struct {

        
            Name string  `json:"name"`
            Mode string  `json:"mode"`
            Meta MultiTenderPaymentMeta  `json:"meta"`
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
            OrderID string  `json:"order_id"`
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

        
            ErrorTrace string  `json:"error_trace"`
            Status float64  `json:"status"`
            Message string  `json:"message"`
         
    }
    
    // StoreReassign used by Order
    type StoreReassign struct {

        
            BagID float64  `json:"bag_id"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            StoreID float64  `json:"store_id"`
            ItemID string  `json:"item_id"`
            MongoArticleID string  `json:"mongo_article_id"`
            ReasonIds []float64  `json:"reason_ids"`
            AffiliateID string  `json:"affiliate_id"`
            FyndOrderID string  `json:"fynd_order_id"`
            SetID string  `json:"set_id"`
         
    }
    
    // StoreReassignResponse used by Order
    type StoreReassignResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // Entities used by Order
    type Entities struct {

        
            AffiliateBagID string  `json:"affiliate_bag_id"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            AffiliateID string  `json:"affiliate_id"`
            ReasonText string  `json:"reason_text"`
            ID string  `json:"id"`
         
    }
    
    // UpdateShipmentLockPayload used by Order
    type UpdateShipmentLockPayload struct {

        
            Action string  `json:"action"`
            EntityType string  `json:"entity_type"`
            ActionType string  `json:"action_type"`
            Entities []Entities  `json:"entities"`
         
    }
    
    // OriginalFilter used by Order
    type OriginalFilter struct {

        
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            AffiliateID string  `json:"affiliate_id"`
         
    }
    
    // Bags used by Order
    type Bags struct {

        
            BagID float64  `json:"bag_id"`
            AffiliateBagID string  `json:"affiliate_bag_id"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            IsLocked bool  `json:"is_locked"`
         
    }
    
    // CheckResponse used by Order
    type CheckResponse struct {

        
            LockStatus bool  `json:"lock_status"`
            ShipmentID string  `json:"shipment_id"`
            Status string  `json:"status"`
            IsBagLocked bool  `json:"is_bag_locked"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
            AffiliateID string  `json:"affiliate_id"`
            IsShipmentLocked bool  `json:"is_shipment_locked"`
            OriginalFilter OriginalFilter  `json:"original_filter"`
            Bags []Bags  `json:"bags"`
         
    }
    
    // UpdateShipmentLockResponse used by Order
    type UpdateShipmentLockResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
            CheckResponse []CheckResponse  `json:"check_response"`
         
    }
    
    // AnnouncementResponse used by Order
    type AnnouncementResponse struct {

        
            LogoURL string  `json:"logo_url"`
            Description string  `json:"description"`
            CompanyID float64  `json:"company_id"`
            PlatformName string  `json:"platform_name"`
            CreatedAt string  `json:"created_at"`
            ToDatetime string  `json:"to_datetime"`
            FromDatetime string  `json:"from_datetime"`
            Title string  `json:"title"`
            PlatformID string  `json:"platform_id"`
            ID float64  `json:"id"`
         
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
    
    // ProductsReasonsData used by Order
    type ProductsReasonsData struct {

        
            ReasonID float64  `json:"reason_id"`
            ReasonText string  `json:"reason_text"`
         
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
    
    // ReasonsData used by Order
    type ReasonsData struct {

        
            Entities []EntitiesReasons  `json:"entities"`
            Products []ProductsReasons  `json:"products"`
         
    }
    
    // EntitiesDataUpdates used by Order
    type EntitiesDataUpdates struct {

        
            Data map[string]interface{}  `json:"data"`
            Filters []map[string]interface{}  `json:"filters"`
         
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
    
    // DataUpdates used by Order
    type DataUpdates struct {

        
            Entities []EntitiesDataUpdates  `json:"entities"`
            Products []ProductsDataUpdates  `json:"products"`
         
    }
    
    // Products used by Order
    type Products struct {

        
            LineNumber float64  `json:"line_number"`
            Identifier string  `json:"identifier"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // ShipmentsRequest used by Order
    type ShipmentsRequest struct {

        
            Reasons ReasonsData  `json:"reasons"`
            DataUpdates DataUpdates  `json:"data_updates"`
            Identifier string  `json:"identifier"`
            Products []Products  `json:"products"`
         
    }
    
    // StatuesRequest used by Order
    type StatuesRequest struct {

        
            Shipments []ShipmentsRequest  `json:"shipments"`
            ExcludeBagsNextState string  `json:"exclude_bags_next_state"`
            Status string  `json:"status"`
         
    }
    
    // UpdateShipmentStatusRequest used by Order
    type UpdateShipmentStatusRequest struct {

        
            ForceTransition bool  `json:"force_transition"`
            UnlockBeforeTransition bool  `json:"unlock_before_transition"`
            Task bool  `json:"task"`
            Statuses []StatuesRequest  `json:"statuses"`
            LockAfterTransition bool  `json:"lock_after_transition"`
         
    }
    
    // ShipmentsResponse used by Order
    type ShipmentsResponse struct {

        
            FinalState map[string]interface{}  `json:"final_state"`
            Meta map[string]interface{}  `json:"meta"`
            Exception string  `json:"exception"`
            Status float64  `json:"status"`
            Code string  `json:"code"`
            Message string  `json:"message"`
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
    
    // AffiliateInventoryPaymentConfig used by Order
    type AffiliateInventoryPaymentConfig struct {

        
            ModeOfPayment string  `json:"mode_of_payment"`
            Source string  `json:"source"`
         
    }
    
    // AffiliateInventoryArticleAssignmentConfig used by Order
    type AffiliateInventoryArticleAssignmentConfig struct {

        
            PostOrderReassignment bool  `json:"post_order_reassignment"`
         
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

        
            Payment AffiliateInventoryPaymentConfig  `json:"payment"`
            ArticleAssignment AffiliateInventoryArticleAssignmentConfig  `json:"article_assignment"`
            Inventory AffiliateInventoryStoreConfig  `json:"inventory"`
            Order AffiliateInventoryOrderConfig  `json:"order"`
            Logistics AffiliateInventoryLogisticsConfig  `json:"logistics"`
         
    }
    
    // AffiliateAppConfigMeta used by Order
    type AffiliateAppConfigMeta struct {

        
            Value string  `json:"value"`
            Name string  `json:"name"`
         
    }
    
    // AffiliateAppConfig used by Order
    type AffiliateAppConfig struct {

        
            Meta []AffiliateAppConfigMeta  `json:"meta"`
            Description string  `json:"description"`
            Name string  `json:"name"`
            UpdatedAt string  `json:"updated_at"`
            Secret string  `json:"secret"`
            CreatedAt string  `json:"created_at"`
            Token string  `json:"token"`
            Owner string  `json:"owner"`
            ID string  `json:"id"`
         
    }
    
    // AffiliateConfig used by Order
    type AffiliateConfig struct {

        
            Inventory AffiliateInventoryConfig  `json:"inventory"`
            App AffiliateAppConfig  `json:"app"`
         
    }
    
    // Affiliate used by Order
    type Affiliate struct {

        
            Token string  `json:"token"`
            Config AffiliateConfig  `json:"config"`
            ID string  `json:"id"`
         
    }
    
    // AffiliateStoreIdMapping used by Order
    type AffiliateStoreIdMapping struct {

        
            MarketplaceStoreID string  `json:"marketplace_store_id"`
            StoreID float64  `json:"store_id"`
         
    }
    
    // OrderConfig used by Order
    type OrderConfig struct {

        
            CreateUser bool  `json:"create_user"`
            BagEndState string  `json:"bag_end_state"`
            Affiliate Affiliate  `json:"affiliate"`
            StoreLookup string  `json:"store_lookup"`
            AffiliateStoreIDMapping []AffiliateStoreIdMapping  `json:"affiliate_store_id_mapping"`
            ArticleLookup string  `json:"article_lookup"`
         
    }
    
    // OrderUser used by Order
    type OrderUser struct {

        
            Address1 string  `json:"address1"`
            Mobile float64  `json:"mobile"`
            LastName string  `json:"last_name"`
            State string  `json:"state"`
            FirstName string  `json:"first_name"`
            Country string  `json:"country"`
            Address2 string  `json:"address2"`
            City string  `json:"city"`
            Phone float64  `json:"phone"`
            Email string  `json:"email"`
            Pincode string  `json:"pincode"`
         
    }
    
    // UserData used by Order
    type UserData struct {

        
            ShippingUser OrderUser  `json:"shipping_user"`
            BillingUser OrderUser  `json:"billing_user"`
         
    }
    
    // ArticleDetails1 used by Order
    type ArticleDetails1 struct {

        
            Weight map[string]interface{}  `json:"weight"`
            Quantity float64  `json:"quantity"`
            Dimension map[string]interface{}  `json:"dimension"`
            ID string  `json:"_id"`
            Category map[string]interface{}  `json:"category"`
            BrandID float64  `json:"brand_id"`
            Attributes map[string]interface{}  `json:"attributes"`
         
    }
    
    // LocationDetails used by Order
    type LocationDetails struct {

        
            FulfillmentType string  `json:"fulfillment_type"`
            Articles []ArticleDetails1  `json:"articles"`
            FulfillmentID float64  `json:"fulfillment_id"`
         
    }
    
    // ShipmentDetails used by Order
    type ShipmentDetails struct {

        
            Meta map[string]interface{}  `json:"meta"`
            Articles []ArticleDetails1  `json:"articles"`
            BoxType string  `json:"box_type"`
            DpID float64  `json:"dp_id"`
            Shipments float64  `json:"shipments"`
            FulfillmentID float64  `json:"fulfillment_id"`
            AffiliateShipmentID string  `json:"affiliate_shipment_id"`
         
    }
    
    // ShipmentConfig used by Order
    type ShipmentConfig struct {

        
            Action string  `json:"action"`
            LocationDetails LocationDetails  `json:"location_details"`
            ToPincode string  `json:"to_pincode"`
            Source string  `json:"source"`
            Journey string  `json:"journey"`
            Shipment []ShipmentDetails  `json:"shipment"`
            Identifier string  `json:"identifier"`
            PaymentMode string  `json:"payment_mode"`
         
    }
    
    // ShipmentData used by Order
    type ShipmentData struct {

        
            ShipmentData ShipmentConfig  `json:"shipment_data"`
         
    }
    
    // MarketPlacePdf used by Order
    type MarketPlacePdf struct {

        
            Label string  `json:"label"`
            Invoice string  `json:"invoice"`
         
    }
    
    // AffiliateBag used by Order
    type AffiliateBag struct {

        
            CompanyID float64  `json:"company_id"`
            AffiliateStoreID string  `json:"affiliate_store_id"`
            Sku string  `json:"sku"`
            ID string  `json:"_id"`
            DeliveryCharge float64  `json:"delivery_charge"`
            PriceEffective float64  `json:"price_effective"`
            HsnCodeID string  `json:"hsn_code_id"`
            ItemID float64  `json:"item_id"`
            ItemSize string  `json:"item_size"`
            Discount float64  `json:"discount"`
            SellerIdentifier string  `json:"seller_identifier"`
            AvlQty float64  `json:"avl_qty"`
            ModifiedOn string  `json:"modified_on"`
            PriceMarked float64  `json:"price_marked"`
            StoreID float64  `json:"store_id"`
            TransferPrice float64  `json:"transfer_price"`
            AmountPaid float64  `json:"amount_paid"`
            FyndStoreID string  `json:"fynd_store_id"`
            UnitPrice float64  `json:"unit_price"`
            Quantity float64  `json:"quantity"`
            AffiliateMeta map[string]interface{}  `json:"affiliate_meta"`
            Identifier map[string]interface{}  `json:"identifier"`
            PdfLinks MarketPlacePdf  `json:"pdf_links"`
         
    }
    
    // OrderPriority used by Order
    type OrderPriority struct {

        
            FulfilmentPriorityText string  `json:"fulfilment_priority_text"`
            AffiliatePriorityCode string  `json:"affiliate_priority_code"`
            FulfilmentPriority float64  `json:"fulfilment_priority"`
         
    }
    
    // OrderInfo used by Order
    type OrderInfo struct {

        
            User UserData  `json:"user"`
            Payment map[string]interface{}  `json:"payment"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            Shipment ShipmentData  `json:"shipment"`
            DeliveryCharges float64  `json:"delivery_charges"`
            CodCharges float64  `json:"cod_charges"`
            OrderValue float64  `json:"order_value"`
            BillingAddress OrderUser  `json:"billing_address"`
            Discount float64  `json:"discount"`
            Bags []AffiliateBag  `json:"bags"`
            ShippingAddress OrderUser  `json:"shipping_address"`
            Coupon string  `json:"coupon"`
            OrderPriority OrderPriority  `json:"order_priority"`
            PaymentMode string  `json:"payment_mode"`
            Items map[string]interface{}  `json:"items"`
         
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
            Description string  `json:"description"`
            Slug string  `json:"slug"`
            ID float64  `json:"id"`
         
    }
    
    // GetActionsResponse used by Order
    type GetActionsResponse struct {

        
            Permissions ActionInfo  `json:"permissions"`
         
    }
    
    // HistoryDict used by Order
    type HistoryDict struct {

        
            TicketID string  `json:"ticket_id"`
            BagID float64  `json:"bag_id"`
            L2Detail string  `json:"l2_detail"`
            TicketURL string  `json:"ticket_url"`
            L3Detail string  `json:"l3_detail"`
            Type string  `json:"type"`
            Createdat string  `json:"createdat"`
            Message string  `json:"message"`
            L1Detail string  `json:"l1_detail"`
            User string  `json:"user"`
         
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
            ShipmentID float64  `json:"shipment_id"`
            AmountPaid float64  `json:"amount_paid"`
            CustomerName string  `json:"customer_name"`
            Message string  `json:"message"`
            PhoneNumber float64  `json:"phone_number"`
            OrderID string  `json:"order_id"`
            BrandName string  `json:"brand_name"`
            PaymentMode string  `json:"payment_mode"`
         
    }
    
    // SendSmsPayload used by Order
    type SendSmsPayload struct {

        
            Data SmsDataPayload  `json:"data"`
            BagID float64  `json:"bag_id"`
            Slug string  `json:"slug"`
         
    }
    
    // OrderDetails used by Order
    type OrderDetails struct {

        
            FyndOrderID string  `json:"fynd_order_id"`
            CreatedAt string  `json:"created_at"`
         
    }
    
    // Meta used by Order
    type Meta struct {

        
            StateManagerUsed string  `json:"state_manager_used"`
            KafkaEmissionStatus float64  `json:"kafka_emission_status"`
         
    }
    
    // ShipmentDetail used by Order
    type ShipmentDetail struct {

        
            Meta Meta  `json:"meta"`
            ShipmentID string  `json:"shipment_id"`
            BagList []float64  `json:"bag_list"`
            Status string  `json:"status"`
            Remarks string  `json:"remarks"`
            ID float64  `json:"id"`
         
    }
    
    // OrderStatusData used by Order
    type OrderStatusData struct {

        
            Errors []string  `json:"errors"`
            OrderDetails OrderDetails  `json:"order_details"`
            ShipmentDetails []ShipmentDetail  `json:"shipment_details"`
         
    }
    
    // OrderStatusResult used by Order
    type OrderStatusResult struct {

        
            Success string  `json:"success"`
            Result []OrderStatusData  `json:"result"`
         
    }
    
    // ManualAssignDPToShipment used by Order
    type ManualAssignDPToShipment struct {

        
            OrderType string  `json:"order_type"`
            QcRequired string  `json:"qc_required"`
            ShipmentIds []string  `json:"shipment_ids"`
            DpID float64  `json:"dp_id"`
         
    }
    
    // ManualAssignDPToShipmentResponse used by Order
    type ManualAssignDPToShipmentResponse struct {

        
            Errors []string  `json:"errors"`
            Success string  `json:"success"`
         
    }
    
    // TaxInfo used by Order
    type TaxInfo struct {

        
            B2bGstinNumber string  `json:"b2b_gstin_number"`
            Gstin string  `json:"gstin"`
         
    }
    
    // Tax used by Order
    type Tax struct {

        
            Breakup []map[string]interface{}  `json:"breakup"`
            TaxExempt bool  `json:"tax_exempt"`
            Name string  `json:"name"`
            Amount map[string]interface{}  `json:"amount"`
         
    }
    
    // Charge used by Order
    type Charge struct {

        
            Name string  `json:"name"`
            Type string  `json:"type"`
            Amount map[string]interface{}  `json:"amount"`
            Tax Tax  `json:"tax"`
            Code string  `json:"code"`
         
    }
    
    // LineItem used by Order
    type LineItem struct {

        
            CustomMessasge string  `json:"custom_messasge"`
            Meta map[string]interface{}  `json:"meta"`
            Charges []Charge  `json:"charges"`
            Quantity float64  `json:"quantity"`
            ExternalLineID string  `json:"external_line_id"`
            SellerIdentifier string  `json:"seller_identifier"`
         
    }
    
    // ProcessingDates used by Order
    type ProcessingDates struct {

        
            DispatchByDate string  `json:"dispatch_by_date"`
            DpPickupSlot map[string]interface{}  `json:"dp_pickup_slot"`
            DispatchAfterDate string  `json:"dispatch_after_date"`
            CustomerPickupSlot map[string]interface{}  `json:"customer_pickup_slot"`
            PackByDate string  `json:"pack_by_date"`
            ConfirmByDate string  `json:"confirm_by_date"`
         
    }
    
    // Shipment used by Order
    type Shipment struct {

        
            Meta map[string]interface{}  `json:"meta"`
            LineItems []LineItem  `json:"line_items"`
            ExternalShipmentID float64  `json:"external_shipment_id"`
            ProcessingDates ProcessingDates  `json:"processing_dates"`
            Priority float64  `json:"priority"`
            LocationID float64  `json:"location_id"`
         
    }
    
    // PaymentMethod used by Order
    type PaymentMethod struct {

        
            Meta map[string]interface{}  `json:"meta"`
            Name string  `json:"name"`
            Amount float64  `json:"amount"`
            RefundBy string  `json:"refund_by"`
            CollectBy string  `json:"collect_by"`
            TransactionData map[string]interface{}  `json:"transaction_data"`
            Mode string  `json:"mode"`
         
    }
    
    // PaymentInfo used by Order
    type PaymentInfo struct {

        
            PrimaryMode string  `json:"primary_mode"`
            PaymentMethods []PaymentMethod  `json:"payment_methods"`
         
    }
    
    // ShippingInfo used by Order
    type ShippingInfo struct {

        
            AlternateEmail string  `json:"alternate_email"`
            State string  `json:"state"`
            ExternalCustomerCode string  `json:"external_customer_code"`
            ShippingType string  `json:"shipping_type"`
            HouseNo string  `json:"house_no"`
            FloorNo string  `json:"floor_no"`
            GeoLocation map[string]interface{}  `json:"geo_location"`
            AddressType string  `json:"address_type"`
            Title string  `json:"title"`
            Pincode string  `json:"pincode"`
            CountryCode string  `json:"country_code"`
            Address1 string  `json:"address1"`
            FirstName string  `json:"first_name"`
            Country string  `json:"country"`
            PrimaryEmail string  `json:"primary_email"`
            Slot []map[string]interface{}  `json:"slot"`
            PrimaryMobileNumber string  `json:"primary_mobile_number"`
            Gender string  `json:"gender"`
            LastName string  `json:"last_name"`
            CustomerCode string  `json:"customer_code"`
            StateCode string  `json:"state_code"`
            MiddleName string  `json:"middle_name"`
            Address2 string  `json:"address2"`
            City string  `json:"city"`
            Landmark string  `json:"landmark"`
            AlternateMobileNumber string  `json:"alternate_mobile_number"`
         
    }
    
    // BillingInfo used by Order
    type BillingInfo struct {

        
            AlternateEmail string  `json:"alternate_email"`
            State string  `json:"state"`
            ExternalCustomerCode string  `json:"external_customer_code"`
            HouseNo string  `json:"house_no"`
            FloorNo string  `json:"floor_no"`
            Title string  `json:"title"`
            Pincode string  `json:"pincode"`
            CountryCode string  `json:"country_code"`
            Address1 string  `json:"address1"`
            FirstName string  `json:"first_name"`
            Country string  `json:"country"`
            PrimaryEmail string  `json:"primary_email"`
            PrimaryMobileNumber string  `json:"primary_mobile_number"`
            Gender string  `json:"gender"`
            LastName string  `json:"last_name"`
            CustomerCode string  `json:"customer_code"`
            StateCode string  `json:"state_code"`
            MiddleName string  `json:"middle_name"`
            Address2 string  `json:"address2"`
            City string  `json:"city"`
            AlternateMobileNumber string  `json:"alternate_mobile_number"`
         
    }
    
    // CreateOrderAPI used by Order
    type CreateOrderAPI struct {

        
            Meta map[string]interface{}  `json:"meta"`
            TaxInfo TaxInfo  `json:"tax_info"`
            CurrencyInfo map[string]interface{}  `json:"currency_info"`
            ExternalOrderID string  `json:"external_order_id"`
            Charges []Charge  `json:"charges"`
            ApplicationID string  `json:"application_id"`
            Shipments []Shipment  `json:"shipments"`
            PaymentInfo PaymentInfo  `json:"payment_info"`
            ShippingInfo ShippingInfo  `json:"shipping_info"`
            ExternalCreationDate string  `json:"external_creation_date"`
            BillingInfo BillingInfo  `json:"billing_info"`
         
    }
    
    // CreateOrderErrorReponse used by Order
    type CreateOrderErrorReponse struct {

        
            Meta string  `json:"meta"`
            RequestID string  `json:"request_id"`
            Exception string  `json:"exception"`
            Status float64  `json:"status"`
            Info interface{}  `json:"info"`
            Code string  `json:"code"`
            Message string  `json:"message"`
            StackTrace string  `json:"stack_trace"`
         
    }
    
    // PaymentMethods used by Order
    type PaymentMethods struct {

        
            CollectBy string  `json:"collect_by"`
            RefundBy string  `json:"refund_by"`
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

        
            LogoURL map[string]interface{}  `json:"logo_url"`
            ShipmentAssignment string  `json:"shipment_assignment"`
            PaymentInfo CreateChannelPaymentInfo  `json:"payment_info"`
            LockStates []string  `json:"lock_states"`
            LocationReassignment bool  `json:"location_reassignment"`
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

        
            StartDate string  `json:"start_date"`
            Mobile float64  `json:"mobile"`
            OrderDetails []FyndOrderIdList  `json:"order_details"`
            EndDate string  `json:"end_date"`
         
    }
    

    
    // DeleteResponse used by Catalog
    type DeleteResponse struct {

        
            Message string  `json:"message"`
         
    }
    
    // SearchKeywordResult used by Catalog
    type SearchKeywordResult struct {

        
            SortOn string  `json:"sort_on"`
            Query map[string]interface{}  `json:"query"`
         
    }
    
    // CreateSearchKeyword used by Catalog
    type CreateSearchKeyword struct {

        
            Words []string  `json:"words"`
            Result SearchKeywordResult  `json:"result"`
            AppID string  `json:"app_id"`
            IsActive bool  `json:"is_active"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // GetSearchWordsData used by Catalog
    type GetSearchWordsData struct {

        
            Words []string  `json:"words"`
            Result map[string]interface{}  `json:"result"`
            AppID string  `json:"app_id"`
            UID string  `json:"uid"`
            IsActive bool  `json:"is_active"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
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

        
            Type string  `json:"type"`
            Params map[string]interface{}  `json:"params"`
            URL string  `json:"url"`
            Query map[string]interface{}  `json:"query"`
         
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
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Action AutocompleteAction  `json:"action"`
            Logo Media  `json:"logo"`
         
    }
    
    // CreateAutocompleteKeyword used by Catalog
    type CreateAutocompleteKeyword struct {

        
            Words []string  `json:"words"`
            Results []AutocompleteResult  `json:"results"`
            AppID string  `json:"app_id"`
            IsActive bool  `json:"is_active"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // GetAutocompleteWordsData used by Catalog
    type GetAutocompleteWordsData struct {

        
            Words []string  `json:"words"`
            UID string  `json:"uid"`
            Results []map[string]interface{}  `json:"results"`
            AppID string  `json:"app_id"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // GetAutocompleteWordsResponse used by Catalog
    type GetAutocompleteWordsResponse struct {

        
            Items []GetAutocompleteWordsData  `json:"items"`
            Page Page  `json:"page"`
         
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

        
            AutoAddToCart bool  `json:"auto_add_to_cart"`
            MaxQuantity float64  `json:"max_quantity"`
            ProductUID float64  `json:"product_uid"`
            AutoSelect bool  `json:"auto_select"`
            MinQuantity float64  `json:"min_quantity"`
            AllowRemove bool  `json:"allow_remove"`
         
    }
    
    // ProductBundleRequest used by Catalog
    type ProductBundleRequest struct {

        
            SameStoreAssignment bool  `json:"same_store_assignment"`
            Meta map[string]interface{}  `json:"meta"`
            CompanyID float64  `json:"company_id"`
            Choice string  `json:"choice"`
            Products []ProductBundleItem  `json:"products"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            CreatedOn string  `json:"created_on"`
            IsActive bool  `json:"is_active"`
            ModifiedOn string  `json:"modified_on"`
            Name string  `json:"name"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Slug string  `json:"slug"`
            Logo string  `json:"logo"`
            PageVisibility []string  `json:"page_visibility"`
         
    }
    
    // GetProductBundleCreateResponse used by Catalog
    type GetProductBundleCreateResponse struct {

        
            SameStoreAssignment bool  `json:"same_store_assignment"`
            Meta map[string]interface{}  `json:"meta"`
            CompanyID float64  `json:"company_id"`
            Choice string  `json:"choice"`
            Products []ProductBundleItem  `json:"products"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            CreatedOn string  `json:"created_on"`
            IsActive bool  `json:"is_active"`
            ModifiedOn string  `json:"modified_on"`
            Name string  `json:"name"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Slug string  `json:"slug"`
            Logo string  `json:"logo"`
            PageVisibility []string  `json:"page_visibility"`
            ID string  `json:"id"`
         
    }
    
    // GetProductBundleListingResponse used by Catalog
    type GetProductBundleListingResponse struct {

        
            Items []GetProductBundleCreateResponse  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // ProductBundleUpdateRequest used by Catalog
    type ProductBundleUpdateRequest struct {

        
            SameStoreAssignment bool  `json:"same_store_assignment"`
            Meta map[string]interface{}  `json:"meta"`
            CompanyID float64  `json:"company_id"`
            Choice string  `json:"choice"`
            Products []ProductBundleItem  `json:"products"`
            IsActive bool  `json:"is_active"`
            ModifiedOn string  `json:"modified_on"`
            Name string  `json:"name"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Slug string  `json:"slug"`
            Logo string  `json:"logo"`
            PageVisibility []string  `json:"page_visibility"`
         
    }
    
    // LimitedProductData used by Catalog
    type LimitedProductData struct {

        
            Images []string  `json:"images"`
            Identifier map[string]interface{}  `json:"identifier"`
            Attributes map[string]interface{}  `json:"attributes"`
            Quantity float64  `json:"quantity"`
            CountryOfOrigin string  `json:"country_of_origin"`
            UID float64  `json:"uid"`
            Sizes []string  `json:"sizes"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            ItemCode string  `json:"item_code"`
            ShortDescription string  `json:"short_description"`
            Price map[string]interface{}  `json:"price"`
         
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

        
            MinEffective float64  `json:"min_effective"`
            MaxMarked float64  `json:"max_marked"`
            MaxEffective float64  `json:"max_effective"`
            MinMarked float64  `json:"min_marked"`
            Currency string  `json:"currency"`
         
    }
    
    // GetProducts used by Catalog
    type GetProducts struct {

        
            AutoAddToCart bool  `json:"auto_add_to_cart"`
            MaxQuantity float64  `json:"max_quantity"`
            ProductUID float64  `json:"product_uid"`
            AutoSelect bool  `json:"auto_select"`
            ProductDetails LimitedProductData  `json:"product_details"`
            MinQuantity float64  `json:"min_quantity"`
            AllowRemove bool  `json:"allow_remove"`
            Sizes []Size  `json:"sizes"`
            Price Price  `json:"price"`
         
    }
    
    // GetProductBundleResponse used by Catalog
    type GetProductBundleResponse struct {

        
            SameStoreAssignment bool  `json:"same_store_assignment"`
            CompanyID float64  `json:"company_id"`
            Products []GetProducts  `json:"products"`
            Meta map[string]interface{}  `json:"meta"`
            Choice string  `json:"choice"`
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            PageVisibility []string  `json:"page_visibility"`
            Slug string  `json:"slug"`
            Logo string  `json:"logo"`
         
    }
    
    // Guide used by Catalog
    type Guide struct {

        
            Meta Meta  `json:"meta"`
         
    }
    
    // ValidateSizeGuide used by Catalog
    type ValidateSizeGuide struct {

        
            Image string  `json:"image"`
            Description string  `json:"description"`
            CompanyID float64  `json:"company_id"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Subtitle string  `json:"subtitle"`
            BrandID float64  `json:"brand_id"`
            Tag string  `json:"tag"`
            Active bool  `json:"active"`
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
            Name string  `json:"name"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Guide Guide  `json:"guide"`
            ID string  `json:"id"`
            Title string  `json:"title"`
         
    }
    
    // ListSizeGuide used by Catalog
    type ListSizeGuide struct {

        
            Items []map[string]interface{}  `json:"items"`
            Page map[string]interface{}  `json:"page"`
         
    }
    
    // SizeGuideResponse used by Catalog
    type SizeGuideResponse struct {

        
            CompanyID float64  `json:"company_id"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Subtitle string  `json:"subtitle"`
            BrandID float64  `json:"brand_id"`
            Tag string  `json:"tag"`
            Active bool  `json:"active"`
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Name string  `json:"name"`
            Guide map[string]interface{}  `json:"guide"`
            ID string  `json:"id"`
            Title string  `json:"title"`
         
    }
    
    // MetaFields used by Catalog
    type MetaFields struct {

        
            Key string  `json:"key"`
            Value string  `json:"value"`
         
    }
    
    // ApplicationItemMeta used by Catalog
    type ApplicationItemMeta struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            CustomMeta []MetaFields  `json:"_custom_meta"`
         
    }
    
    // SuccessResponse1 used by Catalog
    type SuccessResponse1 struct {

        
            Success bool  `json:"success"`
            UID float64  `json:"uid"`
         
    }
    
    // MOQ used by Catalog
    type MOQ struct {

        
            Minimum float64  `json:"minimum"`
            Maximum float64  `json:"maximum"`
            IncrementUnit float64  `json:"increment_unit"`
         
    }
    
    // ApplicationItemResponse used by Catalog
    type ApplicationItemResponse struct {

        
            Seo SEO  `json:"seo"`
            AltText map[string]interface{}  `json:"alt_text"`
            Moq MOQ  `json:"moq"`
         
    }
    
    // GetConfigMetadataResponse used by Catalog
    type GetConfigMetadataResponse struct {

        
            Values []map[string]interface{}  `json:"values"`
            Condition []map[string]interface{}  `json:"condition"`
            Data []map[string]interface{}  `json:"data"`
         
    }
    
    // AttributeDetailsGroup used by Catalog
    type AttributeDetailsGroup struct {

        
            Priority float64  `json:"priority"`
            DisplayType string  `json:"display_type"`
            Unit string  `json:"unit"`
            IsActive bool  `json:"is_active"`
            Key string  `json:"key"`
            Name string  `json:"name"`
            Logo string  `json:"logo"`
            Slug string  `json:"slug"`
         
    }
    
    // AppConfigurationDetail used by Catalog
    type AppConfigurationDetail struct {

        
            Priority float64  `json:"priority"`
            TemplateSlugs []string  `json:"template_slugs"`
            Attributes []AttributeDetailsGroup  `json:"attributes"`
            AppID string  `json:"app_id"`
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Logo string  `json:"logo"`
            IsDefault bool  `json:"is_default"`
         
    }
    
    // ConfigErrorResponse used by Catalog
    type ConfigErrorResponse struct {

        
            Message string  `json:"message"`
         
    }
    
    // PageResponseType used by Catalog
    type PageResponseType struct {

        
            TotalCount float64  `json:"total_count"`
            Current float64  `json:"current"`
            HasNext bool  `json:"has_next"`
            Next float64  `json:"next"`
         
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

        
            Priority float64  `json:"priority"`
            AppID string  `json:"app_id"`
            IsActive bool  `json:"is_active"`
            Key string  `json:"key"`
            Name string  `json:"name"`
            Logo string  `json:"logo"`
            DefaultKey string  `json:"default_key"`
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
    
    // MetaDataListingFilterMetaResponse used by Catalog
    type MetaDataListingFilterMetaResponse struct {

        
            Display string  `json:"display"`
            Key string  `json:"key"`
            FilterTypes []string  `json:"filter_types"`
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
            Compare map[string]interface{}  `json:"compare"`
            Similar map[string]interface{}  `json:"similar"`
            Detail map[string]interface{}  `json:"detail"`
         
    }
    
    // GetCatalogConfigurationMetaData used by Catalog
    type GetCatalogConfigurationMetaData struct {

        
            Listing MetaDataListingResponse  `json:"listing"`
            Product GetCatalogConfigurationDetailsProduct  `json:"product"`
         
    }
    
    // ConfigurationBucketPoints used by Catalog
    type ConfigurationBucketPoints struct {

        
            Display string  `json:"display"`
            Start float64  `json:"start"`
            End float64  `json:"end"`
         
    }
    
    // ConfigurationListingFilterValue used by Catalog
    type ConfigurationListingFilterValue struct {

        
            Value string  `json:"value"`
            Map map[string]interface{}  `json:"map"`
            MapValues []map[string]interface{}  `json:"map_values"`
            BucketPoints []ConfigurationBucketPoints  `json:"bucket_points"`
            Condition string  `json:"condition"`
            Sort string  `json:"sort"`
         
    }
    
    // ConfigurationListingFilterConfig used by Catalog
    type ConfigurationListingFilterConfig struct {

        
            Priority float64  `json:"priority"`
            Type string  `json:"type"`
            IsActive bool  `json:"is_active"`
            Key string  `json:"key"`
            Name string  `json:"name"`
            Logo string  `json:"logo"`
            ValueConfig ConfigurationListingFilterValue  `json:"value_config"`
            DisplayName string  `json:"display_name"`
         
    }
    
    // ConfigurationListingFilter used by Catalog
    type ConfigurationListingFilter struct {

        
            AttributeConfig []ConfigurationListingFilterConfig  `json:"attribute_config"`
            AllowSingle bool  `json:"allow_single"`
         
    }
    
    // ConfigurationListingSortConfig used by Catalog
    type ConfigurationListingSortConfig struct {

        
            Priority float64  `json:"priority"`
            IsActive bool  `json:"is_active"`
            Key string  `json:"key"`
            Name string  `json:"name"`
            Logo string  `json:"logo"`
         
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
            DisplayType string  `json:"display_type"`
            IsActive bool  `json:"is_active"`
            Key string  `json:"key"`
            Name string  `json:"name"`
            Logo string  `json:"logo"`
            Size ProductSize  `json:"size"`
         
    }
    
    // ConfigurationProductVariant used by Catalog
    type ConfigurationProductVariant struct {

        
            Config []ConfigurationProductVariantConfig  `json:"config"`
         
    }
    
    // ConfigurationProductConfig used by Catalog
    type ConfigurationProductConfig struct {

        
            Priority float64  `json:"priority"`
            Subtitle string  `json:"subtitle"`
            IsActive bool  `json:"is_active"`
            Key string  `json:"key"`
            Size ProductSize  `json:"size"`
            Logo string  `json:"logo"`
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
    
    // AppConfiguration used by Catalog
    type AppConfiguration struct {

        
            ConfigType string  `json:"config_type"`
            Listing ConfigurationListing  `json:"listing"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            AppID string  `json:"app_id"`
            Type string  `json:"type"`
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            ConfigID string  `json:"config_id"`
            Product ConfigurationProduct  `json:"product"`
         
    }
    
    // AppCatalogConfiguration used by Catalog
    type AppCatalogConfiguration struct {

        
            ConfigType string  `json:"config_type"`
            Listing ConfigurationListing  `json:"listing"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            AppID string  `json:"app_id"`
            Type string  `json:"type"`
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            ConfigID string  `json:"config_id"`
            ID string  `json:"id"`
            Product ConfigurationProduct  `json:"product"`
         
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

        
            ConfigType string  `json:"config_type"`
            Listing GetCatalogConfigurationDetailsSchemaListing  `json:"listing"`
            AppID string  `json:"app_id"`
            ConfigID string  `json:"config_id"`
            ID string  `json:"id"`
            Product GetCatalogConfigurationDetailsProduct  `json:"product"`
         
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

        
            Value interface{}  `json:"value"`
            SelectedMax float64  `json:"selected_max"`
            QueryFormat string  `json:"query_format"`
            Count float64  `json:"count"`
            Display string  `json:"display"`
            Max float64  `json:"max"`
            CurrencyCode string  `json:"currency_code"`
            Min float64  `json:"min"`
            DisplayFormat string  `json:"display_format"`
            CurrencySymbol string  `json:"currency_symbol"`
            IsSelected bool  `json:"is_selected"`
            SelectedMin float64  `json:"selected_min"`
         
    }
    
    // ProductFiltersKey used by Catalog
    type ProductFiltersKey struct {

        
            Kind string  `json:"kind"`
            Display string  `json:"display"`
            Name string  `json:"name"`
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
    
    // CollectionImage used by Catalog
    type CollectionImage struct {

        
            URL string  `json:"url"`
            AspectRatio string  `json:"aspect_ratio"`
         
    }
    
    // CollectionBadge used by Catalog
    type CollectionBadge struct {

        
            Text string  `json:"text"`
            Color string  `json:"color"`
         
    }
    
    // CollectionQuery used by Catalog
    type CollectionQuery struct {

        
            Op string  `json:"op"`
            Value []interface{}  `json:"value"`
            Attribute string  `json:"attribute"`
         
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
    
    // CollectionSchedule used by Catalog
    type CollectionSchedule struct {

        
            Start string  `json:"start"`
            Duration float64  `json:"duration"`
            NextSchedule []NextSchedule  `json:"next_schedule"`
            Cron string  `json:"cron"`
            End string  `json:"end"`
         
    }
    
    // SeoDetail used by Catalog
    type SeoDetail struct {

        
            Description string  `json:"description"`
            Title string  `json:"title"`
         
    }
    
    // CreateCollection used by Catalog
    type CreateCollection struct {

        
            Description string  `json:"description"`
            Published bool  `json:"published"`
            IsVisible bool  `json:"is_visible"`
            Type string  `json:"type"`
            Logo CollectionImage  `json:"logo"`
            Badge CollectionBadge  `json:"badge"`
            AllowFacets bool  `json:"allow_facets"`
            Query []CollectionQuery  `json:"query"`
            CreatedBy UserInfo  `json:"created_by"`
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            Tags []string  `json:"tags"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            SortOn string  `json:"sort_on"`
            Banners CollectionBanner  `json:"banners"`
            ModifiedBy UserInfo  `json:"modified_by"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            Schedule CollectionSchedule  `json:"_schedule"`
            AllowSort bool  `json:"allow_sort"`
            Priority float64  `json:"priority"`
            Seo SeoDetail  `json:"seo"`
            Meta map[string]interface{}  `json:"meta"`
            AppID string  `json:"app_id"`
            Slug string  `json:"slug"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // BannerImage used by Catalog
    type BannerImage struct {

        
            Type string  `json:"type"`
            URL string  `json:"url"`
            AspectRatio string  `json:"aspect_ratio"`
         
    }
    
    // ImageUrls used by Catalog
    type ImageUrls struct {

        
            Portrait BannerImage  `json:"portrait"`
            Landscape BannerImage  `json:"landscape"`
         
    }
    
    // CollectionCreateResponse used by Catalog
    type CollectionCreateResponse struct {

        
            Description string  `json:"description"`
            Type string  `json:"type"`
            Logo BannerImage  `json:"logo"`
            Badge map[string]interface{}  `json:"badge"`
            AllowFacets bool  `json:"allow_facets"`
            Query []CollectionQuery  `json:"query"`
            Tag []string  `json:"tag"`
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            SortOn string  `json:"sort_on"`
            Banners ImageUrls  `json:"banners"`
            Schedule map[string]interface{}  `json:"_schedule"`
            AllowSort bool  `json:"allow_sort"`
            Priority float64  `json:"priority"`
            Meta map[string]interface{}  `json:"meta"`
            AppID string  `json:"app_id"`
            Slug string  `json:"slug"`
            Cron map[string]interface{}  `json:"cron"`
         
    }
    
    // Media1 used by Catalog
    type Media1 struct {

        
            Type string  `json:"type"`
            Meta map[string]interface{}  `json:"meta"`
            URL string  `json:"url"`
         
    }
    
    // GetCollectionDetailNest used by Catalog
    type GetCollectionDetailNest struct {

        
            Description string  `json:"description"`
            UID string  `json:"uid"`
            Type string  `json:"type"`
            Logo Media1  `json:"logo"`
            Badge map[string]interface{}  `json:"badge"`
            AllowFacets bool  `json:"allow_facets"`
            Query []CollectionQuery  `json:"query"`
            Tag []string  `json:"tag"`
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Banners ImageUrls  `json:"banners"`
            Schedule map[string]interface{}  `json:"_schedule"`
            AllowSort bool  `json:"allow_sort"`
            Priority float64  `json:"priority"`
            Meta map[string]interface{}  `json:"meta"`
            Action Action  `json:"action"`
            AppID string  `json:"app_id"`
            Slug string  `json:"slug"`
            Cron map[string]interface{}  `json:"cron"`
         
    }
    
    // CollectionListingFilterType used by Catalog
    type CollectionListingFilterType struct {

        
            Display string  `json:"display"`
            Name string  `json:"name"`
            IsSelected bool  `json:"is_selected"`
         
    }
    
    // CollectionListingFilterTag used by Catalog
    type CollectionListingFilterTag struct {

        
            Display string  `json:"display"`
            Name string  `json:"name"`
            IsSelected bool  `json:"is_selected"`
         
    }
    
    // CollectionListingFilter used by Catalog
    type CollectionListingFilter struct {

        
            Type []CollectionListingFilterType  `json:"type"`
            Tags []CollectionListingFilterTag  `json:"tags"`
         
    }
    
    // GetCollectionListingResponse used by Catalog
    type GetCollectionListingResponse struct {

        
            Items []GetCollectionDetailNest  `json:"items"`
            Page Page  `json:"page"`
            Filters CollectionListingFilter  `json:"filters"`
         
    }
    
    // CollectionDetailResponse used by Catalog
    type CollectionDetailResponse struct {

        
            Priority float64  `json:"priority"`
            AllowFacets bool  `json:"allow_facets"`
            Query []CollectionQuery  `json:"query"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Description string  `json:"description"`
            Meta map[string]interface{}  `json:"meta"`
            AppID string  `json:"app_id"`
            Tag []string  `json:"tag"`
            IsActive bool  `json:"is_active"`
            Banners ImageUrls  `json:"banners"`
            Name string  `json:"name"`
            Type string  `json:"type"`
            Logo Media1  `json:"logo"`
            Slug string  `json:"slug"`
            Badge map[string]interface{}  `json:"badge"`
            Cron map[string]interface{}  `json:"cron"`
            Schedule map[string]interface{}  `json:"_schedule"`
            AllowSort bool  `json:"allow_sort"`
         
    }
    
    // UpdateCollection used by Catalog
    type UpdateCollection struct {

        
            Description string  `json:"description"`
            Published bool  `json:"published"`
            IsVisible bool  `json:"is_visible"`
            Type string  `json:"type"`
            Logo CollectionImage  `json:"logo"`
            Badge CollectionBadge  `json:"badge"`
            AllowFacets bool  `json:"allow_facets"`
            Query []CollectionQuery  `json:"query"`
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            Tags []string  `json:"tags"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            SortOn string  `json:"sort_on"`
            Banners CollectionBanner  `json:"banners"`
            ModifiedBy UserInfo  `json:"modified_by"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            Schedule CollectionSchedule  `json:"_schedule"`
            AllowSort bool  `json:"allow_sort"`
            Priority float64  `json:"priority"`
            Seo SeoDetail  `json:"seo"`
            Meta map[string]interface{}  `json:"meta"`
            Slug string  `json:"slug"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // ItemQueryForUserCollection used by Catalog
    type ItemQueryForUserCollection struct {

        
            ItemID float64  `json:"item_id"`
            Action string  `json:"action"`
         
    }
    
    // CollectionItemRequest used by Catalog
    type CollectionItemRequest struct {

        
            Type string  `json:"type"`
            Item []ItemQueryForUserCollection  `json:"item"`
            Query []CollectionQuery  `json:"query"`
         
    }
    
    // UpdatedResponse used by Catalog
    type UpdatedResponse struct {

        
            ItemsNotUpdated []float64  `json:"items_not_updated"`
            Message string  `json:"message"`
         
    }
    
    // ProductBrand used by Catalog
    type ProductBrand struct {

        
            Name string  `json:"name"`
            Action Action  `json:"action"`
            Logo Media1  `json:"logo"`
            UID float64  `json:"uid"`
         
    }
    
    // ProductDetailAttribute used by Catalog
    type ProductDetailAttribute struct {

        
            Type string  `json:"type"`
            Key string  `json:"key"`
            Value string  `json:"value"`
         
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
            CurrencyCode string  `json:"currency_code"`
            Min float64  `json:"min"`
         
    }
    
    // ProductListingPrice used by Catalog
    type ProductListingPrice struct {

        
            Marked Price1  `json:"marked"`
            Effective Price1  `json:"effective"`
         
    }
    
    // ProductListingDetail used by Catalog
    type ProductListingDetail struct {

        
            Color string  `json:"color"`
            Description string  `json:"description"`
            UID float64  `json:"uid"`
            Type string  `json:"type"`
            Rating float64  `json:"rating"`
            Brand ProductBrand  `json:"brand"`
            Discount string  `json:"discount"`
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
            Medias []Media1  `json:"medias"`
            Sellable bool  `json:"sellable"`
            RatingCount float64  `json:"rating_count"`
            Similars []string  `json:"similars"`
            Name string  `json:"name"`
            Attributes map[string]interface{}  `json:"attributes"`
            ProductOnlineDate string  `json:"product_online_date"`
            PromoMeta map[string]interface{}  `json:"promo_meta"`
            HasVariant bool  `json:"has_variant"`
            ShortDescription string  `json:"short_description"`
            ImageNature string  `json:"image_nature"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            Price ProductListingPrice  `json:"price"`
            ItemType string  `json:"item_type"`
            Tryouts []string  `json:"tryouts"`
            Slug string  `json:"slug"`
            Highlights []string  `json:"highlights"`
            ItemCode string  `json:"item_code"`
         
    }
    
    // GetCollectionItemsResponse used by Catalog
    type GetCollectionItemsResponse struct {

        
            SortOn []ProductSortOn  `json:"sort_on"`
            Items []ProductListingDetail  `json:"items"`
            Page Page  `json:"page"`
            Filters []ProductFilters  `json:"filters"`
         
    }
    
    // CatalogInsightItem used by Catalog
    type CatalogInsightItem struct {

        
            Count float64  `json:"count"`
            OutOfStockCount float64  `json:"out_of_stock_count"`
            SellableCount float64  `json:"sellable_count"`
         
    }
    
    // CatalogInsightBrand used by Catalog
    type CatalogInsightBrand struct {

        
            TotalArticles float64  `json:"total_articles"`
            AvailableSizes float64  `json:"available_sizes"`
            Name string  `json:"name"`
            TotalSizes float64  `json:"total_sizes"`
            ArticleFreshness float64  `json:"article_freshness"`
            AvailableArticles float64  `json:"available_articles"`
         
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

        
            BrandDistribution CatalogInsightBrand  `json:"brand_distribution"`
            Data CrossSellingData  `json:"data"`
         
    }
    
    // OptInPostRequest used by Catalog
    type OptInPostRequest struct {

        
            OptLevel string  `json:"opt_level"`
            CompanyID float64  `json:"company_id"`
            Platform string  `json:"platform"`
            StoreIds []float64  `json:"store_ids"`
            Enabled bool  `json:"enabled"`
            BrandIds []float64  `json:"brand_ids"`
         
    }
    
    // CompanyOptIn used by Catalog
    type CompanyOptIn struct {

        
            BrandIds []float64  `json:"brand_ids"`
            OptLevel string  `json:"opt_level"`
            CompanyID float64  `json:"company_id"`
            Platform string  `json:"platform"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            CreatedOn float64  `json:"created_on"`
            ModifiedOn float64  `json:"modified_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Enabled bool  `json:"enabled"`
            StoreIds []float64  `json:"store_ids"`
         
    }
    
    // GetOptInPlatform used by Catalog
    type GetOptInPlatform struct {

        
            Items []CompanyOptIn  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // OptinCompanyDetail used by Catalog
    type OptinCompanyDetail struct {

        
            BusinessType string  `json:"business_type"`
            Name string  `json:"name"`
            CompanyType string  `json:"company_type"`
            UID float64  `json:"uid"`
         
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

        
            StoreType string  `json:"store_type"`
            Manager map[string]interface{}  `json:"manager"`
            Timing map[string]interface{}  `json:"timing"`
            CompanyID float64  `json:"company_id"`
            UID float64  `json:"uid"`
            Documents []map[string]interface{}  `json:"documents"`
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
            Name string  `json:"name"`
            DisplayName string  `json:"display_name"`
            StoreCode string  `json:"store_code"`
            AdditionalContacts []map[string]interface{}  `json:"additional_contacts"`
            Address map[string]interface{}  `json:"address"`
         
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
    
    // AttributeMasterMandatoryDetails used by Catalog
    type AttributeMasterMandatoryDetails struct {

        
            L3Keys []string  `json:"l3_keys"`
         
    }
    
    // AttributeMasterMeta used by Catalog
    type AttributeMasterMeta struct {

        
            MandatoryDetails AttributeMasterMandatoryDetails  `json:"mandatory_details"`
            Enriched bool  `json:"enriched"`
         
    }
    
    // AttributeSchemaRange used by Catalog
    type AttributeSchemaRange struct {

        
            Max float64  `json:"max"`
            Min float64  `json:"min"`
         
    }
    
    // AttributeMaster used by Catalog
    type AttributeMaster struct {

        
            Multi bool  `json:"multi"`
            Range AttributeSchemaRange  `json:"range"`
            Type string  `json:"type"`
            Mandatory bool  `json:"mandatory"`
            Format string  `json:"format"`
            AllowedValues []string  `json:"allowed_values"`
         
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
            EnabledForEndConsumer bool  `json:"enabled_for_end_consumer"`
            Description string  `json:"description"`
            Meta AttributeMasterMeta  `json:"meta"`
            IsNested bool  `json:"is_nested"`
            Departments []string  `json:"departments"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Logo string  `json:"logo"`
            Schema AttributeMaster  `json:"schema"`
            ID string  `json:"id"`
            Filters AttributeMasterFilter  `json:"filters"`
         
    }
    
    // ProdcutTemplateCategoriesResponse used by Catalog
    type ProdcutTemplateCategoriesResponse struct {

        
            Items []map[string]interface{}  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // PTErrorResponse used by Catalog
    type PTErrorResponse struct {

        
            Code string  `json:"code"`
            Meta map[string]interface{}  `json:"meta"`
            Errors map[string]interface{}  `json:"errors"`
            Message string  `json:"message"`
            Status float64  `json:"status"`
         
    }
    
    // DepartmentCreateUpdate used by Catalog
    type DepartmentCreateUpdate struct {

        
            PriorityOrder float64  `json:"priority_order"`
            Cls string  `json:"_cls"`
            UID float64  `json:"uid"`
            IsActive bool  `json:"is_active"`
            Platforms map[string]interface{}  `json:"platforms"`
            Name string  `json:"name"`
            Logo string  `json:"logo"`
            Slug string  `json:"slug"`
            Synonyms []string  `json:"synonyms"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Tags []string  `json:"tags"`
         
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

        
            ID string  `json:"_id"`
            UID string  `json:"uid"`
            Contact string  `json:"contact"`
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
         
    }
    
    // GetDepartment used by Catalog
    type GetDepartment struct {

        
            PriorityOrder float64  `json:"priority_order"`
            PageSize float64  `json:"page_size"`
            ItemType string  `json:"item_type"`
            CreatedBy UserSerializer  `json:"created_by"`
            Search string  `json:"search"`
            UID float64  `json:"uid"`
            CreatedOn string  `json:"created_on"`
            IsActive bool  `json:"is_active"`
            ModifiedOn string  `json:"modified_on"`
            Name string  `json:"name"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            Slug string  `json:"slug"`
            Logo string  `json:"logo"`
            Synonyms []string  `json:"synonyms"`
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
            Meta map[string]interface{}  `json:"meta"`
            Errors map[string]interface{}  `json:"errors"`
            Message string  `json:"message"`
            Status float64  `json:"status"`
         
    }
    
    // UserDetail used by Catalog
    type UserDetail struct {

        
            UserID string  `json:"user_id"`
            Contact string  `json:"contact"`
            SuperUser bool  `json:"super_user"`
            Username string  `json:"username"`
         
    }
    
    // DepartmentModel used by Catalog
    type DepartmentModel struct {

        
            PriorityOrder float64  `json:"priority_order"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Cls interface{}  `json:"_cls"`
            CreatedBy UserDetail  `json:"created_by"`
            ID interface{}  `json:"_id"`
            UID float64  `json:"uid"`
            CreatedOn string  `json:"created_on"`
            VerifiedOn string  `json:"verified_on"`
            ModifiedOn string  `json:"modified_on"`
            ModifiedBy UserDetail  `json:"modified_by"`
            Name interface{}  `json:"name"`
            VerifiedBy UserDetail  `json:"verified_by"`
            Slug interface{}  `json:"slug"`
            Logo string  `json:"logo"`
            IsActive bool  `json:"is_active"`
            Synonyms []interface{}  `json:"synonyms"`
         
    }
    
    // ProductTemplate used by Catalog
    type ProductTemplate struct {

        
            Attributes []string  `json:"attributes"`
            Description string  `json:"description"`
            IsExpirable bool  `json:"is_expirable"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Departments []string  `json:"departments"`
            IsPhysical bool  `json:"is_physical"`
            Tag string  `json:"tag"`
            Categories []string  `json:"categories"`
            IsActive bool  `json:"is_active"`
            CreatedOn string  `json:"created_on"`
            Name string  `json:"name"`
            IsArchived bool  `json:"is_archived"`
            Slug string  `json:"slug"`
            Logo string  `json:"logo"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // TemplatesResponse used by Catalog
    type TemplatesResponse struct {

        
            Items ProductTemplate  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // Properties used by Catalog
    type Properties struct {

        
            Description map[string]interface{}  `json:"description"`
            IsDependent map[string]interface{}  `json:"is_dependent"`
            Command map[string]interface{}  `json:"command"`
            CountryOfOrigin map[string]interface{}  `json:"country_of_origin"`
            ReturnConfig map[string]interface{}  `json:"return_config"`
            Trader map[string]interface{}  `json:"trader"`
            IsActive map[string]interface{}  `json:"is_active"`
            Name map[string]interface{}  `json:"name"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            CategorySlug map[string]interface{}  `json:"category_slug"`
            Tags map[string]interface{}  `json:"tags"`
            SizeGuide map[string]interface{}  `json:"size_guide"`
            ProductGroupTag map[string]interface{}  `json:"product_group_tag"`
            HsnCode map[string]interface{}  `json:"hsn_code"`
            ProductPublish map[string]interface{}  `json:"product_publish"`
            Sizes map[string]interface{}  `json:"sizes"`
            ShortDescription map[string]interface{}  `json:"short_description"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            NoOfBoxes map[string]interface{}  `json:"no_of_boxes"`
            ItemType map[string]interface{}  `json:"item_type"`
            MultiSize map[string]interface{}  `json:"multi_size"`
            Media map[string]interface{}  `json:"media"`
            Variants map[string]interface{}  `json:"variants"`
            Slug map[string]interface{}  `json:"slug"`
            ItemCode map[string]interface{}  `json:"item_code"`
            Highlights map[string]interface{}  `json:"highlights"`
            BrandUID map[string]interface{}  `json:"brand_uid"`
            TraderType map[string]interface{}  `json:"trader_type"`
            Currency map[string]interface{}  `json:"currency"`
         
    }
    
    // GlobalValidation used by Catalog
    type GlobalValidation struct {

        
            Description string  `json:"description"`
            Properties Properties  `json:"properties"`
            Type string  `json:"type"`
            Required []string  `json:"required"`
            Definitions map[string]interface{}  `json:"definitions"`
            Title string  `json:"title"`
         
    }
    
    // TemplateValidationData used by Catalog
    type TemplateValidationData struct {

        
            TemplateValidation map[string]interface{}  `json:"template_validation"`
            GlobalValidation GlobalValidation  `json:"global_validation"`
         
    }
    
    // TemplateDetails used by Catalog
    type TemplateDetails struct {

        
            Attributes []string  `json:"attributes"`
            Description string  `json:"description"`
            IsExpirable bool  `json:"is_expirable"`
            Departments []string  `json:"departments"`
            IsPhysical bool  `json:"is_physical"`
            Tag string  `json:"tag"`
            Categories []string  `json:"categories"`
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            IsArchived bool  `json:"is_archived"`
            Slug string  `json:"slug"`
            Logo string  `json:"logo"`
            ID string  `json:"id"`
         
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

        
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
         
    }
    
    // ProductDownloadItemsData used by Catalog
    type ProductDownloadItemsData struct {

        
            Type string  `json:"type"`
            Templates []string  `json:"templates"`
            Brand []string  `json:"brand"`
         
    }
    
    // ProductDownloadsItems used by Catalog
    type ProductDownloadsItems struct {

        
            TemplateTags map[string]interface{}  `json:"template_tags"`
            SellerID float64  `json:"seller_id"`
            URL string  `json:"url"`
            CompletedOn string  `json:"completed_on"`
            CreatedBy VerifiedBy  `json:"created_by"`
            TriggerOn string  `json:"trigger_on"`
            TaskID string  `json:"task_id"`
            ID string  `json:"id"`
            Data ProductDownloadItemsData  `json:"data"`
            Status string  `json:"status"`
         
    }
    
    // ProductDownloadsResponse used by Catalog
    type ProductDownloadsResponse struct {

        
            Items ProductDownloadsItems  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // ProductConfigurationDownloads used by Catalog
    type ProductConfigurationDownloads struct {

        
            Multivalue bool  `json:"multivalue"`
            Data []map[string]interface{}  `json:"data"`
         
    }
    
    // Media2 used by Catalog
    type Media2 struct {

        
            Logo string  `json:"logo"`
            Portrait string  `json:"portrait"`
            Landscape string  `json:"landscape"`
         
    }
    
    // CategoryMappingValues used by Catalog
    type CategoryMappingValues struct {

        
            Name string  `json:"name"`
            CatalogID float64  `json:"catalog_id"`
         
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
    
    // CategoryRequestBody used by Catalog
    type CategoryRequestBody struct {

        
            Priority float64  `json:"priority"`
            Level float64  `json:"level"`
            Tryouts []string  `json:"tryouts"`
            Departments []float64  `json:"departments"`
            Media Media2  `json:"media"`
            IsActive bool  `json:"is_active"`
            Marketplaces CategoryMapping  `json:"marketplaces"`
            Name string  `json:"name"`
            Hierarchy []Hierarchy  `json:"hierarchy"`
            Slug string  `json:"slug"`
            Synonyms []string  `json:"synonyms"`
         
    }
    
    // CategoryCreateResponse used by Catalog
    type CategoryCreateResponse struct {

        
            Message string  `json:"message"`
            UID float64  `json:"uid"`
         
    }
    
    // Category used by Catalog
    type Category struct {

        
            Priority float64  `json:"priority"`
            Level float64  `json:"level"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Tryouts []string  `json:"tryouts"`
            Departments []float64  `json:"departments"`
            UID float64  `json:"uid"`
            Media Media2  `json:"media"`
            IsActive bool  `json:"is_active"`
            Marketplaces CategoryMapping  `json:"marketplaces"`
            Name string  `json:"name"`
            CreatedOn string  `json:"created_on"`
            Hierarchy []Hierarchy  `json:"hierarchy"`
            Slug string  `json:"slug"`
            Synonyms []string  `json:"synonyms"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            ModifiedOn string  `json:"modified_on"`
            ID string  `json:"id"`
         
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

        
            Type string  `json:"type"`
            Name string  `json:"name"`
            Address []string  `json:"address"`
         
    }
    
    // CustomOrder used by Catalog
    type CustomOrder struct {

        
            IsCustomOrder bool  `json:"is_custom_order"`
            ManufacturingTimeUnit string  `json:"manufacturing_time_unit"`
            ManufacturingTime float64  `json:"manufacturing_time"`
         
    }
    
    // ProductPublish used by Catalog
    type ProductPublish struct {

        
            ProductOnlineDate string  `json:"product_online_date"`
            IsSet bool  `json:"is_set"`
         
    }
    
    // TaxIdentifier used by Catalog
    type TaxIdentifier struct {

        
            HsnCodeID string  `json:"hsn_code_id"`
            ReportingHsn string  `json:"reporting_hsn"`
            HsnCode string  `json:"hsn_code"`
         
    }
    
    // NetQuantity used by Catalog
    type NetQuantity struct {

        
            Value float64  `json:"value"`
            Unit interface{}  `json:"unit"`
         
    }
    
    // TeaserTag used by Catalog
    type TeaserTag struct {

        
            URL string  `json:"url"`
            Tag string  `json:"tag"`
         
    }
    
    // ProductCreateUpdate used by Catalog
    type ProductCreateUpdate struct {

        
            ChangeRequestID interface{}  `json:"change_request_id"`
            Description string  `json:"description"`
            IsDependent bool  `json:"is_dependent"`
            CountryOfOrigin string  `json:"country_of_origin"`
            ReturnConfig ReturnConfig  `json:"return_config"`
            UID float64  `json:"uid"`
            Trader []Trader  `json:"trader"`
            Departments []float64  `json:"departments"`
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            CustomOrder CustomOrder  `json:"custom_order"`
            CategorySlug string  `json:"category_slug"`
            Tags []string  `json:"tags"`
            SizeGuide string  `json:"size_guide"`
            ProductGroupTag []string  `json:"product_group_tag"`
            ProductPublish ProductPublish  `json:"product_publish"`
            IsImageLessProduct bool  `json:"is_image_less_product"`
            TaxIdentifier TaxIdentifier  `json:"tax_identifier"`
            ShortDescription string  `json:"short_description"`
            VariantMedia map[string]interface{}  `json:"variant_media"`
            BulkJobID string  `json:"bulk_job_id"`
            NetQuantity NetQuantity  `json:"net_quantity"`
            TeaserTag TeaserTag  `json:"teaser_tag"`
            TemplateTag string  `json:"template_tag"`
            IsSet bool  `json:"is_set"`
            CompanyID float64  `json:"company_id"`
            NoOfBoxes float64  `json:"no_of_boxes"`
            ItemType string  `json:"item_type"`
            Action string  `json:"action"`
            MultiSize bool  `json:"multi_size"`
            Media []Media1  `json:"media"`
            Slug string  `json:"slug"`
            ItemCode interface{}  `json:"item_code"`
            Highlights []string  `json:"highlights"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            BrandUID float64  `json:"brand_uid"`
            Variants map[string]interface{}  `json:"variants"`
            Currency string  `json:"currency"`
            Requester string  `json:"requester"`
         
    }
    
    // Image used by Catalog
    type Image struct {

        
            URL string  `json:"url"`
            SecureURL string  `json:"secure_url"`
            AspectRatioF float64  `json:"aspect_ratio_f"`
            AspectRatio string  `json:"aspect_ratio"`
         
    }
    
    // Logo used by Catalog
    type Logo struct {

        
            URL string  `json:"url"`
            AspectRatioF float64  `json:"aspect_ratio_f"`
            SecureURL string  `json:"secure_url"`
            AspectRatio string  `json:"aspect_ratio"`
         
    }
    
    // ProductPublished used by Catalog
    type ProductPublished struct {

        
            ProductOnlineDate float64  `json:"product_online_date"`
            IsSet bool  `json:"is_set"`
         
    }
    
    // Product used by Catalog
    type Product struct {

        
            Images []Image  `json:"images"`
            Color string  `json:"color"`
            CategoryUID float64  `json:"category_uid"`
            Description string  `json:"description"`
            IsDependent bool  `json:"is_dependent"`
            CountryOfOrigin string  `json:"country_of_origin"`
            IsPhysical bool  `json:"is_physical"`
            UID float64  `json:"uid"`
            Brand Brand  `json:"brand"`
            PrimaryColor string  `json:"primary_color"`
            Departments []float64  `json:"departments"`
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            CategorySlug string  `json:"category_slug"`
            SizeGuide string  `json:"size_guide"`
            AllSizes []map[string]interface{}  `json:"all_sizes"`
            ProductPublish ProductPublished  `json:"product_publish"`
            HsnCode string  `json:"hsn_code"`
            Sizes []map[string]interface{}  `json:"sizes"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            ShortDescription string  `json:"short_description"`
            ImageNature string  `json:"image_nature"`
            ID string  `json:"id"`
            TemplateTag string  `json:"template_tag"`
            IsSet bool  `json:"is_set"`
            ItemType string  `json:"item_type"`
            IsExpirable bool  `json:"is_expirable"`
            MultiSize bool  `json:"multi_size"`
            Media []Media1  `json:"media"`
            L3Mapping []string  `json:"l3_mapping"`
            Highlights []string  `json:"highlights"`
            ItemCode string  `json:"item_code"`
            Slug string  `json:"slug"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            BrandUID float64  `json:"brand_uid"`
            Variants map[string]interface{}  `json:"variants"`
            Currency string  `json:"currency"`
         
    }
    
    // ProductListingResponse used by Catalog
    type ProductListingResponse struct {

        
            Items []Product  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // AttributeMasterSerializer used by Catalog
    type AttributeMasterSerializer struct {

        
            EnabledForEndConsumer bool  `json:"enabled_for_end_consumer"`
            Description string  `json:"description"`
            Logo string  `json:"logo"`
            Synonyms map[string]interface{}  `json:"synonyms"`
            Schema AttributeMaster  `json:"schema"`
            Details AttributeMasterDetails  `json:"details"`
            IsNested bool  `json:"is_nested"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            Departments []string  `json:"departments"`
            Variant bool  `json:"variant"`
            Name string  `json:"name"`
            Tags []string  `json:"tags"`
            Filters AttributeMasterFilter  `json:"filters"`
            Suggestion string  `json:"suggestion"`
            Unit string  `json:"unit"`
            CreatedOn string  `json:"created_on"`
            RawKey string  `json:"raw_key"`
            ModifiedOn string  `json:"modified_on"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            Slug string  `json:"slug"`
         
    }
    
    // ProductAttributesResponse used by Catalog
    type ProductAttributesResponse struct {

        
            Items []AttributeMasterSerializer  `json:"items"`
         
    }
    
    // ValidateProduct used by Catalog
    type ValidateProduct struct {

        
            Valid bool  `json:"valid"`
         
    }
    
    // UserInfo1 used by Catalog
    type UserInfo1 struct {

        
            Username string  `json:"username"`
            Email string  `json:"email"`
            UserID string  `json:"user_id"`
            UID string  `json:"uid"`
         
    }
    
    // BulkJob used by Catalog
    type BulkJob struct {

        
            Total float64  `json:"total"`
            CustomTemplateTag string  `json:"custom_template_tag"`
            TemplateTag string  `json:"template_tag"`
            CancelledRecords []map[string]interface{}  `json:"cancelled_records"`
            CompanyID float64  `json:"company_id"`
            CreatedBy UserInfo1  `json:"created_by"`
            FailedRecords []map[string]interface{}  `json:"failed_records"`
            CreatedOn string  `json:"created_on"`
            IsActive bool  `json:"is_active"`
            ModifiedOn string  `json:"modified_on"`
            ModifiedBy UserInfo1  `json:"modified_by"`
            Stage string  `json:"stage"`
            FilePath string  `json:"file_path"`
            Succeed float64  `json:"succeed"`
            Cancelled float64  `json:"cancelled"`
            Failed float64  `json:"failed"`
            TrackingURL string  `json:"tracking_url"`
         
    }
    
    // BulkResponse used by Catalog
    type BulkResponse struct {

        
            CreatedBy UserInfo1  `json:"created_by"`
            CreatedOn string  `json:"created_on"`
            IsActive bool  `json:"is_active"`
            ModifiedOn string  `json:"modified_on"`
            ModifiedBy UserInfo1  `json:"modified_by"`
            BatchID string  `json:"batch_id"`
         
    }
    
    // UserDetail1 used by Catalog
    type UserDetail1 struct {

        
            UserID string  `json:"user_id"`
            FullName string  `json:"full_name"`
            Username string  `json:"username"`
         
    }
    
    // ProductBulkRequest used by Catalog
    type ProductBulkRequest struct {

        
            Total float64  `json:"total"`
            TemplateTag string  `json:"template_tag"`
            CancelledRecords []string  `json:"cancelled_records"`
            CompanyID float64  `json:"company_id"`
            CreatedBy UserDetail1  `json:"created_by"`
            FailedRecords []string  `json:"failed_records"`
            CreatedOn string  `json:"created_on"`
            IsActive bool  `json:"is_active"`
            ModifiedOn string  `json:"modified_on"`
            FilePath string  `json:"file_path"`
            Stage string  `json:"stage"`
            ModifiedBy UserDetail1  `json:"modified_by"`
            Succeed float64  `json:"succeed"`
            Cancelled float64  `json:"cancelled"`
            Failed float64  `json:"failed"`
            Template ProductTemplate  `json:"template"`
         
    }
    
    // ProductBulkRequestList used by Catalog
    type ProductBulkRequestList struct {

        
            Items ProductBulkRequest  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // BulkProductRequest used by Catalog
    type BulkProductRequest struct {

        
            BatchID string  `json:"batch_id"`
            CompanyID float64  `json:"company_id"`
            TemplateTag string  `json:"template_tag"`
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
    
    // ProductBulkAssets used by Catalog
    type ProductBulkAssets struct {

        
            URL string  `json:"url"`
            User map[string]interface{}  `json:"user"`
            CompanyID float64  `json:"company_id"`
         
    }
    
    // UserCommon used by Catalog
    type UserCommon struct {

        
            CompanyID float64  `json:"company_id"`
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
         
    }
    
    // Items used by Catalog
    type Items struct {

        
            Total float64  `json:"total"`
            CancelledRecords []string  `json:"cancelled_records"`
            CompanyID float64  `json:"company_id"`
            CreatedBy UserCommon  `json:"created_by"`
            FailedRecords []string  `json:"failed_records"`
            ID string  `json:"id"`
            CreatedOn string  `json:"created_on"`
            FilePath string  `json:"file_path"`
            ModifiedOn string  `json:"modified_on"`
            ModifiedBy UserCommon  `json:"modified_by"`
            Stage string  `json:"stage"`
            IsActive bool  `json:"is_active"`
            Retry float64  `json:"retry"`
            Succeed float64  `json:"succeed"`
            Cancelled float64  `json:"cancelled"`
            Failed float64  `json:"failed"`
            TrackingURL string  `json:"tracking_url"`
         
    }
    
    // BulkAssetResponse used by Catalog
    type BulkAssetResponse struct {

        
            Items []Items  `json:"items"`
            Page Page  `json:"page"`
         
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

        
            BrandUID float64  `json:"brand_uid"`
            ItemCode string  `json:"item_code"`
            UID float64  `json:"uid"`
         
    }
    
    // GTIN used by Catalog
    type GTIN struct {

        
            GtinType string  `json:"gtin_type"`
            Primary bool  `json:"primary"`
            GtinValue string  `json:"gtin_value"`
         
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

        
            PriceTransfer float64  `json:"price_transfer"`
            PriceEffective float64  `json:"price_effective"`
            ItemHeight float64  `json:"item_height"`
            Identifiers []GTIN  `json:"identifiers"`
            IsSet bool  `json:"is_set"`
            ItemWeightUnitOfMeasure string  `json:"item_weight_unit_of_measure"`
            Quantity float64  `json:"quantity"`
            ItemDimensionsUnitOfMeasure string  `json:"item_dimensions_unit_of_measure"`
            ItemWeight float64  `json:"item_weight"`
            Set InventorySet  `json:"set"`
            Size string  `json:"size"`
            ExpirationDate string  `json:"expiration_date"`
            StoreCode string  `json:"store_code"`
            ItemLength float64  `json:"item_length"`
            ItemWidth float64  `json:"item_width"`
            Currency string  `json:"currency"`
            Price float64  `json:"price"`
         
    }
    
    // InventoryRequest used by Catalog
    type InventoryRequest struct {

        
            Item ItemQuery  `json:"item"`
            CompanyID float64  `json:"company_id"`
            Sizes []InvSize  `json:"sizes"`
         
    }
    
    // InventoryResponse used by Catalog
    type InventoryResponse struct {

        
            ItemID float64  `json:"item_id"`
            PriceTransfer float64  `json:"price_transfer"`
            PriceEffective float64  `json:"price_effective"`
            SellableQuantity float64  `json:"sellable_quantity"`
            Store map[string]interface{}  `json:"store"`
            Quantity float64  `json:"quantity"`
            InventoryUpdatedOn string  `json:"inventory_updated_on"`
            SellerIdentifier string  `json:"seller_identifier"`
            UID string  `json:"uid"`
            Size string  `json:"size"`
            Identifiers map[string]interface{}  `json:"identifiers"`
            Currency string  `json:"currency"`
            Price float64  `json:"price"`
         
    }
    
    // InventoryResponsePaginated used by Catalog
    type InventoryResponsePaginated struct {

        
            Items []InventoryResponse  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // ReturnConfig1 used by Catalog
    type ReturnConfig1 struct {

        
            Time float64  `json:"time"`
            Returnable bool  `json:"returnable"`
            Unit string  `json:"unit"`
         
    }
    
    // BrandMeta used by Catalog
    type BrandMeta struct {

        
            Name string  `json:"name"`
            ID float64  `json:"id"`
         
    }
    
    // WeightResponse used by Catalog
    type WeightResponse struct {

        
            Shipping float64  `json:"shipping"`
            Unit string  `json:"unit"`
            IsDefault bool  `json:"is_default"`
         
    }
    
    // Trader1 used by Catalog
    type Trader1 struct {

        
            Type string  `json:"type"`
            Name string  `json:"name"`
            Address []string  `json:"address"`
         
    }
    
    // ManufacturerResponse used by Catalog
    type ManufacturerResponse struct {

        
            Name string  `json:"name"`
            Address string  `json:"address"`
            IsDefault bool  `json:"is_default"`
         
    }
    
    // QuantityBase used by Catalog
    type QuantityBase struct {

        
            UpdatedAt string  `json:"updated_at"`
            Count float64  `json:"count"`
         
    }
    
    // Quantities used by Catalog
    type Quantities struct {

        
            NotAvailable QuantityBase  `json:"not_available"`
            OrderCommitted QuantityBase  `json:"order_committed"`
            Damaged QuantityBase  `json:"damaged"`
            Sellable QuantityBase  `json:"sellable"`
         
    }
    
    // CompanyMeta used by Catalog
    type CompanyMeta struct {

        
            ID float64  `json:"id"`
         
    }
    
    // DimensionResponse used by Catalog
    type DimensionResponse struct {

        
            Height float64  `json:"height"`
            Length float64  `json:"length"`
            Width float64  `json:"width"`
            Unit string  `json:"unit"`
            IsDefault bool  `json:"is_default"`
         
    }
    
    // PriceMeta used by Catalog
    type PriceMeta struct {

        
            TpNotes map[string]interface{}  `json:"tp_notes"`
            Marked float64  `json:"marked"`
            UpdatedAt string  `json:"updated_at"`
            Currency string  `json:"currency"`
            Effective float64  `json:"effective"`
            Transfer float64  `json:"transfer"`
         
    }
    
    // InventorySellerResponse used by Catalog
    type InventorySellerResponse struct {

        
            Store StoreMeta  `json:"store"`
            FyndMeta map[string]interface{}  `json:"fynd_meta"`
            ReturnConfig ReturnConfig1  `json:"return_config"`
            CountryOfOrigin string  `json:"country_of_origin"`
            UID string  `json:"uid"`
            Brand BrandMeta  `json:"brand"`
            Weight WeightResponse  `json:"weight"`
            Trader []Trader1  `json:"trader"`
            TrackInventory bool  `json:"track_inventory"`
            TotalQuantity float64  `json:"total_quantity"`
            Manufacturer ManufacturerResponse  `json:"manufacturer"`
            CreatedBy UserSerializer  `json:"created_by"`
            Quantities Quantities  `json:"quantities"`
            IsActive bool  `json:"is_active"`
            Company CompanyMeta  `json:"company"`
            Tags []string  `json:"tags"`
            Dimension DimensionResponse  `json:"dimension"`
            Fragile bool  `json:"fragile"`
            Set InventorySet  `json:"set"`
            Size string  `json:"size"`
            TaxIdentifier map[string]interface{}  `json:"tax_identifier"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            Stage string  `json:"stage"`
            TraceID string  `json:"trace_id"`
            Price PriceMeta  `json:"price"`
            ItemID float64  `json:"item_id"`
            AddedOnStore string  `json:"added_on_store"`
            Identifier map[string]interface{}  `json:"identifier"`
            FyndArticleCode string  `json:"fynd_article_code"`
            IsSet bool  `json:"is_set"`
            Meta map[string]interface{}  `json:"meta"`
            SellerIdentifier string  `json:"seller_identifier"`
            FyndItemCode string  `json:"fynd_item_code"`
            RawMeta map[string]interface{}  `json:"raw_meta"`
            ExpirationDate string  `json:"expiration_date"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // InventorySellerIdentifierResponsePaginated used by Catalog
    type InventorySellerIdentifierResponsePaginated struct {

        
            Items []InventorySellerResponse  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // BulkInventoryGetItems used by Catalog
    type BulkInventoryGetItems struct {

        
            Total float64  `json:"total"`
            CancelledRecords []string  `json:"cancelled_records"`
            CompanyID float64  `json:"company_id"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            FailedRecords []string  `json:"failed_records"`
            CreatedOn string  `json:"created_on"`
            FilePath string  `json:"file_path"`
            IsActive bool  `json:"is_active"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            ModifiedOn string  `json:"modified_on"`
            Stage string  `json:"stage"`
            Succeed float64  `json:"succeed"`
            Cancelled float64  `json:"cancelled"`
            Failed float64  `json:"failed"`
            ID string  `json:"id"`
         
    }
    
    // BulkInventoryGet used by Catalog
    type BulkInventoryGet struct {

        
            Items []BulkInventoryGetItems  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // InventoryJobPayload used by Catalog
    type InventoryJobPayload struct {

        
            PriceEffective float64  `json:"price_effective"`
            TotalQuantity float64  `json:"total_quantity"`
            ItemWeightUnitOfMeasure string  `json:"item_weight_unit_of_measure"`
            Quantity float64  `json:"quantity"`
            SellerIdentifier string  `json:"seller_identifier"`
            ItemDimensionsUnitOfMeasure string  `json:"item_dimensions_unit_of_measure"`
            ExpirationDate string  `json:"expiration_date"`
            TraceID string  `json:"trace_id"`
            StoreCode string  `json:"store_code"`
            Currency string  `json:"currency"`
            Tags []string  `json:"tags"`
            PriceMarked float64  `json:"price_marked"`
            Price float64  `json:"price"`
         
    }
    
    // InventoryBulkRequest used by Catalog
    type InventoryBulkRequest struct {

        
            BatchID string  `json:"batch_id"`
            CompanyID float64  `json:"company_id"`
            User map[string]interface{}  `json:"user"`
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

        
            SellerID float64  `json:"seller_id"`
            RequestParams map[string]interface{}  `json:"request_params"`
            TriggerOn string  `json:"trigger_on"`
            TaskID string  `json:"task_id"`
            Status string  `json:"status"`
         
    }
    
    // InventoryExportJob used by Catalog
    type InventoryExportJob struct {

        
            SellerID float64  `json:"seller_id"`
            CompletedOn string  `json:"completed_on"`
            TaskID string  `json:"task_id"`
            RequestParams map[string]interface{}  `json:"request_params"`
            TriggerOn string  `json:"trigger_on"`
            URL string  `json:"url"`
            Status string  `json:"status"`
         
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
            TotalQuantity float64  `json:"total_quantity"`
            SellerIdentifier string  `json:"seller_identifier"`
            ExpirationDate string  `json:"expiration_date"`
            TraceID string  `json:"trace_id"`
            StoreID float64  `json:"store_id"`
            Tags []string  `json:"tags"`
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

        
            Reason InventoryFailedReason  `json:"reason"`
            Data InventoryPayload  `json:"data"`
         
    }
    
    // InventoryUpdateResponse used by Catalog
    type InventoryUpdateResponse struct {

        
            Items []InventoryResponseItem  `json:"items"`
            Message string  `json:"message"`
         
    }
    
    // HsnUpsert used by Catalog
    type HsnUpsert struct {

        
            CompanyID float64  `json:"company_id"`
            Tax2 float64  `json:"tax2"`
            HsnCode string  `json:"hsn_code"`
            UID float64  `json:"uid"`
            Threshold1 float64  `json:"threshold1"`
            IsActive bool  `json:"is_active"`
            Threshold2 float64  `json:"threshold2"`
            Hs2Code string  `json:"hs2_code"`
            TaxOnMrp bool  `json:"tax_on_mrp"`
            TaxOnEsp bool  `json:"tax_on_esp"`
            Tax1 float64  `json:"tax1"`
         
    }
    
    // HsnCodesObject used by Catalog
    type HsnCodesObject struct {

        
            CompanyID float64  `json:"company_id"`
            Tax2 float64  `json:"tax2"`
            HsnCode string  `json:"hsn_code"`
            ID string  `json:"id"`
            Threshold1 float64  `json:"threshold1"`
            ModifiedOn string  `json:"modified_on"`
            Threshold2 float64  `json:"threshold2"`
            Hs2Code string  `json:"hs2_code"`
            TaxOnMrp bool  `json:"tax_on_mrp"`
            TaxOnEsp bool  `json:"tax_on_esp"`
            Tax1 float64  `json:"tax1"`
         
    }
    
    // HsnCode used by Catalog
    type HsnCode struct {

        
            Data HsnCodesObject  `json:"data"`
         
    }
    
    // PageResponse used by Catalog
    type PageResponse struct {

        
            Size float64  `json:"size"`
            Current string  `json:"current"`
            HasPrevious bool  `json:"has_previous"`
            HasNext bool  `json:"has_next"`
            ItemTotal float64  `json:"item_total"`
         
    }
    
    // HsnCodesListingResponse used by Catalog
    type HsnCodesListingResponse struct {

        
            Items []HsnCodesObject  `json:"items"`
            Page PageResponse  `json:"page"`
         
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

        
            EffectiveDate string  `json:"effective_date"`
            Rate float64  `json:"rate"`
            Cess float64  `json:"cess"`
            Threshold float64  `json:"threshold"`
         
    }
    
    // HSNDataInsertV2 used by Catalog
    type HSNDataInsertV2 struct {

        
            Description string  `json:"description"`
            CreatedBy map[string]interface{}  `json:"created_by"`
            CountryCode string  `json:"country_code"`
            HsnCode string  `json:"hsn_code"`
            Type string  `json:"type"`
            CreatedOn string  `json:"created_on"`
            Taxes []TaxSlab  `json:"taxes"`
            ModifiedBy map[string]interface{}  `json:"modified_by"`
            ReportingHsn string  `json:"reporting_hsn"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // HsnCodesListingResponseSchemaV2 used by Catalog
    type HsnCodesListingResponseSchemaV2 struct {

        
            Items []HSNDataInsertV2  `json:"items"`
            Page PageResponse  `json:"page"`
         
    }
    
    // BrandItem used by Catalog
    type BrandItem struct {

        
            Action Action  `json:"action"`
            Departments []string  `json:"departments"`
            UID float64  `json:"uid"`
            Banners ImageUrls  `json:"banners"`
            Name string  `json:"name"`
            Logo Media  `json:"logo"`
            Slug string  `json:"slug"`
            Discount string  `json:"discount"`
         
    }
    
    // BrandListingResponse used by Catalog
    type BrandListingResponse struct {

        
            Items []BrandItem  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // Department used by Catalog
    type Department struct {

        
            PriorityOrder float64  `json:"priority_order"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Logo Media  `json:"logo"`
         
    }
    
    // DepartmentResponse used by Catalog
    type DepartmentResponse struct {

        
            Items []Department  `json:"items"`
         
    }
    
    // ThirdLevelChild used by Catalog
    type ThirdLevelChild struct {

        
            Action Action  `json:"action"`
            UID float64  `json:"uid"`
            Childs []map[string]interface{}  `json:"childs"`
            Banners ImageUrls  `json:"banners"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // SecondLevelChild used by Catalog
    type SecondLevelChild struct {

        
            Action Action  `json:"action"`
            UID float64  `json:"uid"`
            Childs []ThirdLevelChild  `json:"childs"`
            Banners ImageUrls  `json:"banners"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // Child used by Catalog
    type Child struct {

        
            Action Action  `json:"action"`
            UID float64  `json:"uid"`
            Childs []SecondLevelChild  `json:"childs"`
            Banners ImageUrls  `json:"banners"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // CategoryItems used by Catalog
    type CategoryItems struct {

        
            Action Action  `json:"action"`
            UID float64  `json:"uid"`
            Childs []Child  `json:"childs"`
            Banners ImageUrls  `json:"banners"`
            Name string  `json:"name"`
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
            SortOn []ProductSortOn  `json:"sort_on"`
            Items []ProductListingDetail  `json:"items"`
            Operators map[string]interface{}  `json:"operators"`
            Filters []ProductFilters  `json:"filters"`
         
    }
    
    // ProductDetail used by Catalog
    type ProductDetail struct {

        
            Color string  `json:"color"`
            Description string  `json:"description"`
            UID float64  `json:"uid"`
            Type string  `json:"type"`
            Rating float64  `json:"rating"`
            Brand ProductBrand  `json:"brand"`
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
            Medias []Media1  `json:"medias"`
            RatingCount float64  `json:"rating_count"`
            Similars []string  `json:"similars"`
            Name string  `json:"name"`
            Attributes map[string]interface{}  `json:"attributes"`
            ProductOnlineDate string  `json:"product_online_date"`
            PromoMeta map[string]interface{}  `json:"promo_meta"`
            HasVariant bool  `json:"has_variant"`
            ShortDescription string  `json:"short_description"`
            ImageNature string  `json:"image_nature"`
            TeaserTag map[string]interface{}  `json:"teaser_tag"`
            ItemType string  `json:"item_type"`
            Tryouts []string  `json:"tryouts"`
            Slug string  `json:"slug"`
            Highlights []string  `json:"highlights"`
            ItemCode string  `json:"item_code"`
         
    }
    
    // InventoryPage used by Catalog
    type InventoryPage struct {

        
            Type string  `json:"type"`
            NextID string  `json:"next_id"`
            HasPrevious bool  `json:"has_previous"`
            HasNext bool  `json:"has_next"`
            ItemTotal float64  `json:"item_total"`
         
    }
    
    // InventoryStockResponse used by Catalog
    type InventoryStockResponse struct {

        
            Items []map[string]interface{}  `json:"items"`
            Page InventoryPage  `json:"page"`
         
    }
    
    // InvoiceCredSerializer used by Catalog
    type InvoiceCredSerializer struct {

        
            Password string  `json:"password"`
            Enabled bool  `json:"enabled"`
            Username string  `json:"username"`
         
    }
    
    // InvoiceDetailsSerializer used by Catalog
    type InvoiceDetailsSerializer struct {

        
            EInvoice InvoiceCredSerializer  `json:"e_invoice"`
            EWaybill InvoiceCredSerializer  `json:"e_waybill"`
         
    }
    
    // UserSerializer1 used by Catalog
    type UserSerializer1 struct {

        
            UserID string  `json:"user_id"`
            Contact string  `json:"contact"`
            Username string  `json:"username"`
         
    }
    
    // UserSerializer2 used by Catalog
    type UserSerializer2 struct {

        
            UserID string  `json:"user_id"`
            Contact string  `json:"contact"`
            Username string  `json:"username"`
         
    }
    
    // GetAddressSerializer used by Catalog
    type GetAddressSerializer struct {

        
            Latitude float64  `json:"latitude"`
            Address1 string  `json:"address1"`
            Longitude float64  `json:"longitude"`
            Pincode float64  `json:"pincode"`
            CountryCode string  `json:"country_code"`
            City string  `json:"city"`
            Address2 string  `json:"address2"`
            AddressType string  `json:"address_type"`
            Landmark string  `json:"landmark"`
            Country string  `json:"country"`
            State string  `json:"state"`
         
    }
    
    // GetCompanySerializer used by Catalog
    type GetCompanySerializer struct {

        
            BusinessType string  `json:"business_type"`
            CreatedBy UserSerializer2  `json:"created_by"`
            UID float64  `json:"uid"`
            CreatedOn string  `json:"created_on"`
            VerifiedOn string  `json:"verified_on"`
            ModifiedOn string  `json:"modified_on"`
            ModifiedBy UserSerializer2  `json:"modified_by"`
            Name string  `json:"name"`
            VerifiedBy UserSerializer2  `json:"verified_by"`
            CompanyType string  `json:"company_type"`
            Stage string  `json:"stage"`
            RejectReason string  `json:"reject_reason"`
            Addresses []GetAddressSerializer  `json:"addresses"`
         
    }
    
    // LocationIntegrationType used by Catalog
    type LocationIntegrationType struct {

        
            Order string  `json:"order"`
            Inventory string  `json:"inventory"`
         
    }
    
    // ProductReturnConfigSerializer used by Catalog
    type ProductReturnConfigSerializer struct {

        
            OnSameStore bool  `json:"on_same_store"`
            StoreUID float64  `json:"store_uid"`
         
    }
    
    // SellerPhoneNumber used by Catalog
    type SellerPhoneNumber struct {

        
            Number string  `json:"number"`
            CountryCode float64  `json:"country_code"`
         
    }
    
    // LocationManagerSerializer used by Catalog
    type LocationManagerSerializer struct {

        
            Email string  `json:"email"`
            Name string  `json:"name"`
            MobileNo SellerPhoneNumber  `json:"mobile_no"`
         
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
    
    // GetLocationSerializer used by Catalog
    type GetLocationSerializer struct {

        
            GstCredentials InvoiceDetailsSerializer  `json:"gst_credentials"`
            Warnings map[string]interface{}  `json:"warnings"`
            UID float64  `json:"uid"`
            VerifiedBy UserSerializer1  `json:"verified_by"`
            PhoneNumber string  `json:"phone_number"`
            CreatedBy UserSerializer1  `json:"created_by"`
            NotificationEmails []string  `json:"notification_emails"`
            Name string  `json:"name"`
            Company GetCompanySerializer  `json:"company"`
            IntegrationType LocationIntegrationType  `json:"integration_type"`
            ProductReturnConfig ProductReturnConfigSerializer  `json:"product_return_config"`
            StoreType string  `json:"store_type"`
            Code string  `json:"code"`
            CreatedOn string  `json:"created_on"`
            VerifiedOn string  `json:"verified_on"`
            ModifiedOn string  `json:"modified_on"`
            ModifiedBy UserSerializer1  `json:"modified_by"`
            DisplayName string  `json:"display_name"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            Stage string  `json:"stage"`
            Documents []Document  `json:"documents"`
            Address GetAddressSerializer  `json:"address"`
            Manager LocationManagerSerializer  `json:"manager"`
            Timing []LocationDayWiseSerializer  `json:"timing"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
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
    
    // ApplicationStoreJson used by Catalog
    type ApplicationStoreJson struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    

    
    // CompanyTaxesSerializer used by CompanyProfile
    type CompanyTaxesSerializer struct {

        
            Enable bool  `json:"enable"`
            EffectiveDate string  `json:"effective_date"`
            Rate float64  `json:"rate"`
         
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

        
            Taxes []CompanyTaxesSerializer  `json:"taxes"`
            CreatedBy UserSerializer  `json:"created_by"`
            BusinessType string  `json:"business_type"`
            NotificationEmails []string  `json:"notification_emails"`
            BusinessInfo string  `json:"business_info"`
            ModifiedOn string  `json:"modified_on"`
            UID float64  `json:"uid"`
            Mode string  `json:"mode"`
            Documents []Document  `json:"documents"`
            BusinessCountryInfo BusinessCountryInfo  `json:"business_country_info"`
            VerifiedOn string  `json:"verified_on"`
            ContactDetails ContactDetails  `json:"contact_details"`
            BusinessDetails BusinessDetails  `json:"business_details"`
            Stage string  `json:"stage"`
            Warnings map[string]interface{}  `json:"warnings"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            CompanyType string  `json:"company_type"`
            FranchiseEnabled bool  `json:"franchise_enabled"`
            CreatedOn string  `json:"created_on"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            Name string  `json:"name"`
            ModifiedBy UserSerializer  `json:"modified_by"`
         
    }
    
    // CompanyTaxesSerializer1 used by CompanyProfile
    type CompanyTaxesSerializer1 struct {

        
            Enable bool  `json:"enable"`
            EffectiveDate string  `json:"effective_date"`
            Rate float64  `json:"rate"`
         
    }
    
    // CreateUpdateAddressSerializer used by CompanyProfile
    type CreateUpdateAddressSerializer struct {

        
            City string  `json:"city"`
            Address2 string  `json:"address2"`
            AddressType string  `json:"address_type"`
            State string  `json:"state"`
            CountryCode string  `json:"country_code"`
            Longitude float64  `json:"longitude"`
            Pincode float64  `json:"pincode"`
            Landmark string  `json:"landmark"`
            Address1 string  `json:"address1"`
            Country string  `json:"country"`
            Latitude float64  `json:"latitude"`
         
    }
    
    // UpdateCompany used by CompanyProfile
    type UpdateCompany struct {

        
            BusinessInfo string  `json:"business_info"`
            CompanyType string  `json:"company_type"`
            Documents []Document  `json:"documents"`
            FranchiseEnabled bool  `json:"franchise_enabled"`
            Taxes []CompanyTaxesSerializer1  `json:"taxes"`
            RejectReason string  `json:"reject_reason"`
            ContactDetails ContactDetails  `json:"contact_details"`
            Addresses []CreateUpdateAddressSerializer  `json:"addresses"`
            BusinessDetails BusinessDetails  `json:"business_details"`
            Name string  `json:"name"`
            BusinessType string  `json:"business_type"`
            NotificationEmails []string  `json:"notification_emails"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Warnings map[string]interface{}  `json:"warnings"`
         
    }
    
    // ProfileSuccessResponse used by CompanyProfile
    type ProfileSuccessResponse struct {

        
            UID float64  `json:"uid"`
            Success bool  `json:"success"`
         
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
            Brand DocumentsObj  `json:"brand"`
            Store DocumentsObj  `json:"store"`
            StoreDocuments DocumentsObj  `json:"store_documents"`
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

        
            CreatedBy UserSerializer  `json:"created_by"`
            Description string  `json:"description"`
            ModifiedOn string  `json:"modified_on"`
            Synonyms []string  `json:"synonyms"`
            UID float64  `json:"uid"`
            Mode string  `json:"mode"`
            Logo string  `json:"logo"`
            RejectReason string  `json:"reject_reason"`
            VerifiedOn string  `json:"verified_on"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Stage string  `json:"stage"`
            Warnings map[string]interface{}  `json:"warnings"`
            VerifiedBy UserSerializer  `json:"verified_by"`
            SlugKey string  `json:"slug_key"`
            CreatedOn string  `json:"created_on"`
            Banner BrandBannerSerializer  `json:"banner"`
            Name string  `json:"name"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            ModifiedBy UserSerializer  `json:"modified_by"`
         
    }
    
    // CreateUpdateBrandRequestSerializer used by CompanyProfile
    type CreateUpdateBrandRequestSerializer struct {

        
            UID float64  `json:"uid"`
            Logo string  `json:"logo"`
            Banner BrandBannerSerializer  `json:"banner"`
            CompanyID float64  `json:"company_id"`
            Description string  `json:"description"`
            Name string  `json:"name"`
            LocaleLanguage map[string]interface{}  `json:"_locale_language"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            BrandTier string  `json:"brand_tier"`
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

        
            UID float64  `json:"uid"`
            MarketChannels []string  `json:"market_channels"`
            CompanyType string  `json:"company_type"`
            BusinessCountryInfo BusinessCountryInfo  `json:"business_country_info"`
            CreatedOn string  `json:"created_on"`
            RejectReason string  `json:"reject_reason"`
            VerifiedOn string  `json:"verified_on"`
            Addresses []GetAddressSerializer  `json:"addresses"`
            Name string  `json:"name"`
            CreatedBy UserSerializer  `json:"created_by"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            ModifiedOn string  `json:"modified_on"`
            BusinessType string  `json:"business_type"`
            NotificationEmails []string  `json:"notification_emails"`
            Stage string  `json:"stage"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Details CompanyDetails  `json:"details"`
            VerifiedBy UserSerializer  `json:"verified_by"`
         
    }
    
    // CompanyBrandSerializer used by CompanyProfile
    type CompanyBrandSerializer struct {

        
            Company CompanySerializer  `json:"company"`
            UID float64  `json:"uid"`
            CreatedOn string  `json:"created_on"`
            RejectReason string  `json:"reject_reason"`
            VerifiedOn string  `json:"verified_on"`
            Brand GetBrandResponseSerializer  `json:"brand"`
            Warnings map[string]interface{}  `json:"warnings"`
            CreatedBy UserSerializer  `json:"created_by"`
            ModifiedBy UserSerializer  `json:"modified_by"`
            ModifiedOn string  `json:"modified_on"`
            Stage string  `json:"stage"`
            VerifiedBy UserSerializer  `json:"verified_by"`
         
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
    
    // LocationSerializer used by CompanyProfile
    type LocationSerializer struct {

        
            Company float64  `json:"company"`
            UID float64  `json:"uid"`
            Timing []LocationDayWiseSerializer  `json:"timing"`
            ProductReturnConfig ProductReturnConfigSerializer  `json:"product_return_config"`
            GstCredentials InvoiceDetailsSerializer  `json:"gst_credentials"`
            Code string  `json:"code"`
            Documents []Document  `json:"documents"`
            DisplayName string  `json:"display_name"`
            Warnings map[string]interface{}  `json:"warnings"`
            StoreType string  `json:"store_type"`
            Name string  `json:"name"`
            NotificationEmails []string  `json:"notification_emails"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Manager LocationManagerSerializer  `json:"manager"`
            Address GetAddressSerializer  `json:"address"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            Stage string  `json:"stage"`
         
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
    

    
    // RuleDefinition used by Cart
    type RuleDefinition struct {

        
            IsExact bool  `json:"is_exact"`
            ApplicableOn string  `json:"applicable_on"`
            Scope []string  `json:"scope"`
            CurrencyCode string  `json:"currency_code"`
            Type string  `json:"type"`
            CalculateOn string  `json:"calculate_on"`
            AutoApply bool  `json:"auto_apply"`
            ValueType string  `json:"value_type"`
         
    }
    
    // State used by Cart
    type State struct {

        
            IsDisplay bool  `json:"is_display"`
            IsPublic bool  `json:"is_public"`
            IsArchived bool  `json:"is_archived"`
         
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
    
    // CouponAction used by Cart
    type CouponAction struct {

        
            ActionDate string  `json:"action_date"`
            TxnMode string  `json:"txn_mode"`
         
    }
    
    // CouponAuthor used by Cart
    type CouponAuthor struct {

        
            ModifiedBy string  `json:"modified_by"`
            CreatedBy string  `json:"created_by"`
         
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
            Subtitle string  `json:"subtitle"`
            Title string  `json:"title"`
            Auto DisplayMetaDict  `json:"auto"`
         
    }
    
    // Validity used by Cart
    type Validity struct {

        
            Priority float64  `json:"priority"`
         
    }
    
    // Validation used by Cart
    type Validation struct {

        
            AppID []string  `json:"app_id"`
            UserRegisteredAfter string  `json:"user_registered_after"`
            Anonymous bool  `json:"anonymous"`
         
    }
    
    // CouponSchedule used by Cart
    type CouponSchedule struct {

        
            NextSchedule []map[string]interface{}  `json:"next_schedule"`
            Start string  `json:"start"`
            Cron string  `json:"cron"`
            Duration float64  `json:"duration"`
            End string  `json:"end"`
         
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

        
            Networks []string  `json:"networks"`
            Uses PaymentAllowValue  `json:"uses"`
            Codes []string  `json:"codes"`
            Types []string  `json:"types"`
         
    }
    
    // UsesRemaining used by Cart
    type UsesRemaining struct {

        
            Total float64  `json:"total"`
            App float64  `json:"app"`
            User float64  `json:"user"`
         
    }
    
    // UsesRestriction used by Cart
    type UsesRestriction struct {

        
            Remaining UsesRemaining  `json:"remaining"`
            Maximum UsesRemaining  `json:"maximum"`
         
    }
    
    // PostOrder used by Cart
    type PostOrder struct {

        
            ReturnAllowed bool  `json:"return_allowed"`
            CancellationAllowed bool  `json:"cancellation_allowed"`
         
    }
    
    // PriceRange used by Cart
    type PriceRange struct {

        
            Max float64  `json:"max"`
            Min float64  `json:"min"`
         
    }
    
    // Restrictions used by Cart
    type Restrictions struct {

        
            UserGroups []float64  `json:"user_groups"`
            Platforms []string  `json:"platforms"`
            BulkBundle BulkBundleRestriction  `json:"bulk_bundle"`
            Payments map[string]PaymentModes  `json:"payments"`
            OrderingStores []float64  `json:"ordering_stores"`
            CouponAllowed bool  `json:"coupon_allowed"`
            Uses UsesRestriction  `json:"uses"`
            PostOrder PostOrder  `json:"post_order"`
            PriceRange PriceRange  `json:"price_range"`
         
    }
    
    // Rule used by Cart
    type Rule struct {

        
            Min float64  `json:"min"`
            Value float64  `json:"value"`
            Max float64  `json:"max"`
            DiscountQty float64  `json:"discount_qty"`
            Key float64  `json:"key"`
         
    }
    
    // CouponAdd used by Cart
    type CouponAdd struct {

        
            RuleDefinition RuleDefinition  `json:"rule_definition"`
            State State  `json:"state"`
            DateMeta CouponDateMeta  `json:"date_meta"`
            Ownership Ownership  `json:"ownership"`
            Code string  `json:"code"`
            Action CouponAction  `json:"action"`
            Identifiers Identifier  `json:"identifiers"`
            Tags []string  `json:"tags"`
            TypeSlug string  `json:"type_slug"`
            Author CouponAuthor  `json:"author"`
            DisplayMeta DisplayMeta  `json:"display_meta"`
            Validity Validity  `json:"validity"`
            Validation Validation  `json:"validation"`
            Schedule CouponSchedule  `json:"_schedule"`
            Restrictions Restrictions  `json:"restrictions"`
            Rule []Rule  `json:"rule"`
         
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

        
            RuleDefinition RuleDefinition  `json:"rule_definition"`
            State State  `json:"state"`
            DateMeta CouponDateMeta  `json:"date_meta"`
            Ownership Ownership  `json:"ownership"`
            Code string  `json:"code"`
            Action CouponAction  `json:"action"`
            Identifiers Identifier  `json:"identifiers"`
            Tags []string  `json:"tags"`
            TypeSlug string  `json:"type_slug"`
            Author CouponAuthor  `json:"author"`
            DisplayMeta DisplayMeta  `json:"display_meta"`
            Validity Validity  `json:"validity"`
            Validation Validation  `json:"validation"`
            Schedule CouponSchedule  `json:"_schedule"`
            Restrictions Restrictions  `json:"restrictions"`
            Rule []Rule  `json:"rule"`
         
    }
    
    // CouponPartialUpdate used by Cart
    type CouponPartialUpdate struct {

        
            Archive bool  `json:"archive"`
            Schedule CouponSchedule  `json:"schedule"`
         
    }
    
    // CompareObject used by Cart
    type CompareObject struct {

        
            GreaterThan float64  `json:"greater_than"`
            Equals float64  `json:"equals"`
            LessThan float64  `json:"less_than"`
            GreaterThanEquals float64  `json:"greater_than_equals"`
            LessThanEquals float64  `json:"less_than_equals"`
         
    }
    
    // ItemCriteria used by Cart
    type ItemCriteria struct {

        
            AllItems bool  `json:"all_items"`
            AvailableZones []string  `json:"available_zones"`
            CartQuantity CompareObject  `json:"cart_quantity"`
            CartTotal CompareObject  `json:"cart_total"`
            ItemID []float64  `json:"item_id"`
            ItemStore []float64  `json:"item_store"`
            ItemBrand []float64  `json:"item_brand"`
            ItemExcludeCategory []float64  `json:"item_exclude_category"`
            ItemExcludeBrand []float64  `json:"item_exclude_brand"`
            ItemCategory []float64  `json:"item_category"`
            ItemCompany []float64  `json:"item_company"`
            ItemExcludeCompany []float64  `json:"item_exclude_company"`
            CartUniqueItemQuantity CompareObject  `json:"cart_unique_item_quantity"`
            ItemSize []string  `json:"item_size"`
            ItemExcludeID []float64  `json:"item_exclude_id"`
            BuyRules []string  `json:"buy_rules"`
            CartUniqueItemAmount CompareObject  `json:"cart_unique_item_amount"`
            ItemExcludeStore []float64  `json:"item_exclude_store"`
            ItemExcludeSku []string  `json:"item_exclude_sku"`
            ItemSku []string  `json:"item_sku"`
         
    }
    
    // DiscountOffer used by Cart
    type DiscountOffer struct {

        
            DiscountAmount float64  `json:"discount_amount"`
            ApportionDiscount bool  `json:"apportion_discount"`
            MinOfferQuantity float64  `json:"min_offer_quantity"`
            MaxUsagePerTransaction float64  `json:"max_usage_per_transaction"`
            Code string  `json:"code"`
            DiscountPrice float64  `json:"discount_price"`
            DiscountPercentage float64  `json:"discount_percentage"`
            PartialCanRet bool  `json:"partial_can_ret"`
            MaxDiscountAmount float64  `json:"max_discount_amount"`
            MaxOfferQuantity float64  `json:"max_offer_quantity"`
         
    }
    
    // DiscountRule used by Cart
    type DiscountRule struct {

        
            DiscountType string  `json:"discount_type"`
            BuyCondition string  `json:"buy_condition"`
            ItemCriteria ItemCriteria  `json:"item_criteria"`
            Offer DiscountOffer  `json:"offer"`
         
    }
    
    // PromotionDateMeta used by Cart
    type PromotionDateMeta struct {

        
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
         
    }
    
    // Ownership1 used by Cart
    type Ownership1 struct {

        
            PayableCategory string  `json:"payable_category"`
            PayableBy string  `json:"payable_by"`
         
    }
    
    // PromotionSchedule used by Cart
    type PromotionSchedule struct {

        
            NextSchedule []map[string]interface{}  `json:"next_schedule"`
            Start string  `json:"start"`
            Cron string  `json:"cron"`
            Duration float64  `json:"duration"`
            End string  `json:"end"`
            Published bool  `json:"published"`
         
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
    
    // UserRegistered used by Cart
    type UserRegistered struct {

        
            End string  `json:"end"`
            Start string  `json:"start"`
         
    }
    
    // PostOrder1 used by Cart
    type PostOrder1 struct {

        
            ReturnAllowed bool  `json:"return_allowed"`
            CancellationAllowed bool  `json:"cancellation_allowed"`
         
    }
    
    // Restrictions1 used by Cart
    type Restrictions1 struct {

        
            UserGroups []float64  `json:"user_groups"`
            Payments []PromotionPaymentModes  `json:"payments"`
            UserID []string  `json:"user_id"`
            AnonymousUsers bool  `json:"anonymous_users"`
            Uses UsesRestriction1  `json:"uses"`
            OrderQuantity float64  `json:"order_quantity"`
            UserRegistered UserRegistered  `json:"user_registered"`
            PostOrder PostOrder1  `json:"post_order"`
            Platforms []string  `json:"platforms"`
         
    }
    
    // PromotionAuthor used by Cart
    type PromotionAuthor struct {

        
            ModifiedBy string  `json:"modified_by"`
            CreatedBy string  `json:"created_by"`
         
    }
    
    // DisplayMeta1 used by Cart
    type DisplayMeta1 struct {

        
            Description string  `json:"description"`
            OfferText string  `json:"offer_text"`
            Name string  `json:"name"`
         
    }
    
    // PromotionAction used by Cart
    type PromotionAction struct {

        
            ActionType string  `json:"action_type"`
            ActionDate string  `json:"action_date"`
         
    }
    
    // Visibility used by Cart
    type Visibility struct {

        
            Pdp bool  `json:"pdp"`
            CouponList bool  `json:"coupon_list"`
         
    }
    
    // PromotionListItem used by Cart
    type PromotionListItem struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            DiscountRules []DiscountRule  `json:"discount_rules"`
            PromotionType string  `json:"promotion_type"`
            PromoGroup string  `json:"promo_group"`
            DateMeta PromotionDateMeta  `json:"date_meta"`
            Ownership Ownership1  `json:"ownership"`
            Code string  `json:"code"`
            Stackable bool  `json:"stackable"`
            ApplyPriority float64  `json:"apply_priority"`
            Schedule PromotionSchedule  `json:"_schedule"`
            ApplicationID string  `json:"application_id"`
            Restrictions Restrictions1  `json:"restrictions"`
            ApplyExclusive string  `json:"apply_exclusive"`
            Author PromotionAuthor  `json:"author"`
            DisplayMeta DisplayMeta1  `json:"display_meta"`
            Currency string  `json:"currency"`
            BuyRules map[string]ItemCriteria  `json:"buy_rules"`
            PostOrderAction PromotionAction  `json:"post_order_action"`
            Visiblility Visibility  `json:"visiblility"`
            ApplyAllDiscount bool  `json:"apply_all_discount"`
            CalculateOn string  `json:"calculate_on"`
            Mode string  `json:"mode"`
         
    }
    
    // PromotionsResponse used by Cart
    type PromotionsResponse struct {

        
            Page Page  `json:"page"`
            Items PromotionListItem  `json:"items"`
         
    }
    
    // PromotionAdd used by Cart
    type PromotionAdd struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            DiscountRules []DiscountRule  `json:"discount_rules"`
            PromotionType string  `json:"promotion_type"`
            PromoGroup string  `json:"promo_group"`
            DateMeta PromotionDateMeta  `json:"date_meta"`
            Ownership Ownership1  `json:"ownership"`
            Code string  `json:"code"`
            Stackable bool  `json:"stackable"`
            ApplyPriority float64  `json:"apply_priority"`
            Schedule PromotionSchedule  `json:"_schedule"`
            ApplicationID string  `json:"application_id"`
            Restrictions Restrictions1  `json:"restrictions"`
            ApplyExclusive string  `json:"apply_exclusive"`
            Author PromotionAuthor  `json:"author"`
            DisplayMeta DisplayMeta1  `json:"display_meta"`
            Currency string  `json:"currency"`
            BuyRules map[string]ItemCriteria  `json:"buy_rules"`
            PostOrderAction PromotionAction  `json:"post_order_action"`
            Visiblility Visibility  `json:"visiblility"`
            ApplyAllDiscount bool  `json:"apply_all_discount"`
            CalculateOn string  `json:"calculate_on"`
            Mode string  `json:"mode"`
         
    }
    
    // PromotionUpdate used by Cart
    type PromotionUpdate struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            DiscountRules []DiscountRule  `json:"discount_rules"`
            PromotionType string  `json:"promotion_type"`
            PromoGroup string  `json:"promo_group"`
            DateMeta PromotionDateMeta  `json:"date_meta"`
            Ownership Ownership1  `json:"ownership"`
            Code string  `json:"code"`
            Stackable bool  `json:"stackable"`
            ApplyPriority float64  `json:"apply_priority"`
            Schedule PromotionSchedule  `json:"_schedule"`
            ApplicationID string  `json:"application_id"`
            Restrictions Restrictions1  `json:"restrictions"`
            ApplyExclusive string  `json:"apply_exclusive"`
            Author PromotionAuthor  `json:"author"`
            DisplayMeta DisplayMeta1  `json:"display_meta"`
            Currency string  `json:"currency"`
            BuyRules map[string]ItemCriteria  `json:"buy_rules"`
            PostOrderAction PromotionAction  `json:"post_order_action"`
            Visiblility Visibility  `json:"visiblility"`
            ApplyAllDiscount bool  `json:"apply_all_discount"`
            CalculateOn string  `json:"calculate_on"`
            Mode string  `json:"mode"`
         
    }
    
    // PromotionPartialUpdate used by Cart
    type PromotionPartialUpdate struct {

        
            Archive bool  `json:"archive"`
            Schedule PromotionSchedule  `json:"schedule"`
         
    }
    
    // ActivePromosResponse used by Cart
    type ActivePromosResponse struct {

        
            IsHidden bool  `json:"is_hidden"`
            ModifiedOn string  `json:"modified_on"`
            Example string  `json:"example"`
            Description string  `json:"description"`
            EntityType string  `json:"entity_type"`
            Subtitle string  `json:"subtitle"`
            CreatedOn string  `json:"created_on"`
            Type string  `json:"type"`
            Title string  `json:"title"`
            EntitySlug string  `json:"entity_slug"`
         
    }
    
    // CartItem used by Cart
    type CartItem struct {

        
            Quantity float64  `json:"quantity"`
            ProductID string  `json:"product_id"`
            Size string  `json:"size"`
         
    }
    
    // OpenapiCartDetailsRequest used by Cart
    type OpenapiCartDetailsRequest struct {

        
            CartItems CartItem  `json:"cart_items"`
         
    }
    
    // DisplayBreakup used by Cart
    type DisplayBreakup struct {

        
            CurrencyCode string  `json:"currency_code"`
            Display string  `json:"display"`
            CurrencySymbol string  `json:"currency_symbol"`
            Value float64  `json:"value"`
            Message []string  `json:"message"`
            Key string  `json:"key"`
         
    }
    
    // LoyaltyPoints used by Cart
    type LoyaltyPoints struct {

        
            Description string  `json:"description"`
            IsApplied bool  `json:"is_applied"`
            Total float64  `json:"total"`
            Applicable float64  `json:"applicable"`
         
    }
    
    // CouponBreakup used by Cart
    type CouponBreakup struct {

        
            Code string  `json:"code"`
            CouponValue float64  `json:"coupon_value"`
            Description string  `json:"description"`
            IsApplied bool  `json:"is_applied"`
            Value float64  `json:"value"`
            Type string  `json:"type"`
            UID string  `json:"uid"`
            Message string  `json:"message"`
            Title string  `json:"title"`
            SubTitle string  `json:"sub_title"`
            MinimumCartValue float64  `json:"minimum_cart_value"`
            CouponType string  `json:"coupon_type"`
            MaxDiscountValue float64  `json:"max_discount_value"`
         
    }
    
    // RawBreakup used by Cart
    type RawBreakup struct {

        
            CodCharge float64  `json:"cod_charge"`
            Discount float64  `json:"discount"`
            Vog float64  `json:"vog"`
            MrpTotal float64  `json:"mrp_total"`
            YouSaved float64  `json:"you_saved"`
            Subtotal float64  `json:"subtotal"`
            ConvenienceFee float64  `json:"convenience_fee"`
            Coupon float64  `json:"coupon"`
            DeliveryCharge float64  `json:"delivery_charge"`
            Total float64  `json:"total"`
            FyndCash float64  `json:"fynd_cash"`
            GstCharges float64  `json:"gst_charges"`
         
    }
    
    // CartBreakup used by Cart
    type CartBreakup struct {

        
            Display []DisplayBreakup  `json:"display"`
            LoyaltyPoints LoyaltyPoints  `json:"loyalty_points"`
            Coupon CouponBreakup  `json:"coupon"`
            Raw RawBreakup  `json:"raw"`
         
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

        
            Type string  `json:"type"`
            URL string  `json:"url"`
            Query ActionQuery  `json:"query"`
         
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
    
    // ProductImage used by Cart
    type ProductImage struct {

        
            AspectRatio string  `json:"aspect_ratio"`
            SecureURL string  `json:"secure_url"`
            URL string  `json:"url"`
         
    }
    
    // CartProduct used by Cart
    type CartProduct struct {

        
            Action ProductAction  `json:"action"`
            Slug string  `json:"slug"`
            Brand BaseInfo  `json:"brand"`
            Categories []CategoryInfo  `json:"categories"`
            Images []ProductImage  `json:"images"`
            Type string  `json:"type"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // ProductPrice used by Cart
    type ProductPrice struct {

        
            AddOn float64  `json:"add_on"`
            Effective float64  `json:"effective"`
            CurrencyCode string  `json:"currency_code"`
            CurrencySymbol string  `json:"currency_symbol"`
            Selling float64  `json:"selling"`
            Marked float64  `json:"marked"`
         
    }
    
    // ProductPriceInfo used by Cart
    type ProductPriceInfo struct {

        
            Base ProductPrice  `json:"base"`
            Converted ProductPrice  `json:"converted"`
         
    }
    
    // PromoMeta used by Cart
    type PromoMeta struct {

        
            Message string  `json:"message"`
         
    }
    
    // FreeGiftItem used by Cart
    type FreeGiftItem struct {

        
            ItemName string  `json:"item_name"`
            ItemBrandName string  `json:"item_brand_name"`
            ItemPriceDetails map[string]interface{}  `json:"item_price_details"`
            ItemSlug string  `json:"item_slug"`
            ItemID float64  `json:"item_id"`
            ItemImagesURL []string  `json:"item_images_url"`
         
    }
    
    // AppliedFreeArticles used by Cart
    type AppliedFreeArticles struct {

        
            Quantity float64  `json:"quantity"`
            FreeGiftItemDetails FreeGiftItem  `json:"free_gift_item_details"`
            ParentItemIdentifier string  `json:"parent_item_identifier"`
            ArticleID string  `json:"article_id"`
         
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

        
            ArticleQuantity float64  `json:"article_quantity"`
            Amount float64  `json:"amount"`
            BuyRules []BuyRules  `json:"buy_rules"`
            PromoID string  `json:"promo_id"`
            AppliedFreeArticles []AppliedFreeArticles  `json:"applied_free_articles"`
            MrpPromotion bool  `json:"mrp_promotion"`
            PromotionType string  `json:"promotion_type"`
            DiscountRules []DiscountRulesApp  `json:"discount_rules"`
            PromotionName string  `json:"promotion_name"`
            OfferText string  `json:"offer_text"`
            PromotionGroup string  `json:"promotion_group"`
         
    }
    
    // BasePrice used by Cart
    type BasePrice struct {

        
            CurrencyCode string  `json:"currency_code"`
            CurrencySymbol string  `json:"currency_symbol"`
            Effective float64  `json:"effective"`
            Marked float64  `json:"marked"`
         
    }
    
    // ArticlePriceInfo used by Cart
    type ArticlePriceInfo struct {

        
            Base BasePrice  `json:"base"`
            Converted BasePrice  `json:"converted"`
         
    }
    
    // ProductArticle used by Cart
    type ProductArticle struct {

        
            Store BaseInfo  `json:"store"`
            Quantity float64  `json:"quantity"`
            Price ArticlePriceInfo  `json:"price"`
            ProductGroupTags []string  `json:"product_group_tags"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            Size string  `json:"size"`
            Type string  `json:"type"`
            UID string  `json:"uid"`
            Seller BaseInfo  `json:"seller"`
         
    }
    
    // ProductAvailability used by Cart
    type ProductAvailability struct {

        
            OutOfStock bool  `json:"out_of_stock"`
            Deliverable bool  `json:"deliverable"`
            OtherStoreQuantity float64  `json:"other_store_quantity"`
            IsValid bool  `json:"is_valid"`
            Sizes []string  `json:"sizes"`
         
    }
    
    // CartProductInfo used by Cart
    type CartProductInfo struct {

        
            Discount string  `json:"discount"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
            Product CartProduct  `json:"product"`
            Quantity float64  `json:"quantity"`
            Price ProductPriceInfo  `json:"price"`
            CouponMessage string  `json:"coupon_message"`
            BulkOffer map[string]interface{}  `json:"bulk_offer"`
            PromoMeta PromoMeta  `json:"promo_meta"`
            PromotionsApplied []AppliedPromotion  `json:"promotions_applied"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            PricePerUnit ProductPriceInfo  `json:"price_per_unit"`
            Article ProductArticle  `json:"article"`
            Message string  `json:"message"`
            IsSet bool  `json:"is_set"`
            Availability ProductAvailability  `json:"availability"`
            Key string  `json:"key"`
         
    }
    
    // OpenapiCartDetailsResponse used by Cart
    type OpenapiCartDetailsResponse struct {

        
            BreakupValues CartBreakup  `json:"breakup_values"`
            Message string  `json:"message"`
            IsValid bool  `json:"is_valid"`
            Items []CartProductInfo  `json:"items"`
         
    }
    
    // OpenApiErrorResponse used by Cart
    type OpenApiErrorResponse struct {

        
            Errors map[string]interface{}  `json:"errors"`
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // ShippingAddress used by Cart
    type ShippingAddress struct {

        
            Phone float64  `json:"phone"`
            State string  `json:"state"`
            CountryCode string  `json:"country_code"`
            Country string  `json:"country"`
            AddressType string  `json:"address_type"`
            City string  `json:"city"`
            Address string  `json:"address"`
            Meta map[string]interface{}  `json:"meta"`
            Area string  `json:"area"`
            Landmark string  `json:"landmark"`
            Email string  `json:"email"`
            Pincode float64  `json:"pincode"`
            AreaCode string  `json:"area_code"`
            Name string  `json:"name"`
            AreaCodeSlug string  `json:"area_code_slug"`
         
    }
    
    // OpenApiCartServiceabilityRequest used by Cart
    type OpenApiCartServiceabilityRequest struct {

        
            CartItems CartItem  `json:"cart_items"`
            ShippingAddress ShippingAddress  `json:"shipping_address"`
         
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
    
    // OpenApiCartServiceabilityResponse used by Cart
    type OpenApiCartServiceabilityResponse struct {

        
            IsValid bool  `json:"is_valid"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            Message string  `json:"message"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            Items []CartProductInfo  `json:"items"`
         
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
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            AmountPaid float64  `json:"amount_paid"`
            Quantity float64  `json:"quantity"`
            PriceEffective float64  `json:"price_effective"`
            CodCharges float64  `json:"cod_charges"`
            Files []OpenApiFiles  `json:"files"`
            CashbackApplied float64  `json:"cashback_applied"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            Meta CartItemMeta  `json:"meta"`
            ProductID float64  `json:"product_id"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            DeliveryCharges float64  `json:"delivery_charges"`
            EmployeeDiscount float64  `json:"employee_discount"`
            PriceMarked float64  `json:"price_marked"`
            PaymentMethods []MultiTenderPaymentMethod  `json:"payment_methods"`
            Size string  `json:"size"`
         
    }
    
    // OpenApiPlatformCheckoutReq used by Cart
    type OpenApiPlatformCheckoutReq struct {

        
            PaymentMode string  `json:"payment_mode"`
            Comment string  `json:"comment"`
            ShippingAddress ShippingAddress  `json:"shipping_address"`
            CouponValue float64  `json:"coupon_value"`
            CurrencyCode string  `json:"currency_code"`
            Files []OpenApiFiles  `json:"files"`
            OrderID string  `json:"order_id"`
            CashbackApplied float64  `json:"cashback_applied"`
            CartValue float64  `json:"cart_value"`
            Coupon string  `json:"coupon"`
            DeliveryCharges float64  `json:"delivery_charges"`
            EmployeeDiscount map[string]interface{}  `json:"employee_discount"`
            CartItems []OpenApiOrderItem  `json:"cart_items"`
            PaymentMethods []MultiTenderPaymentMethod  `json:"payment_methods"`
            BillingAddress ShippingAddress  `json:"billing_address"`
            CouponCode string  `json:"coupon_code"`
            LoyaltyDiscount float64  `json:"loyalty_discount"`
            CodCharges float64  `json:"cod_charges"`
            AffiliateOrderID string  `json:"affiliate_order_id"`
            Gstin string  `json:"gstin"`
         
    }
    
    // OpenApiCheckoutResponse used by Cart
    type OpenApiCheckoutResponse struct {

        
            OrderRefID string  `json:"order_ref_id"`
            Message string  `json:"message"`
            Success bool  `json:"success"`
            OrderID string  `json:"order_id"`
         
    }
    
    // AbandonedCart used by Cart
    type AbandonedCart struct {

        
            PaymentMode string  `json:"payment_mode"`
            Comment string  `json:"comment"`
            CartValue float64  `json:"cart_value"`
            UID float64  `json:"uid"`
            Coupon map[string]interface{}  `json:"coupon"`
            Shipments []map[string]interface{}  `json:"shipments"`
            FyndCredits map[string]interface{}  `json:"fynd_credits"`
            PaymentMethods []map[string]interface{}  `json:"payment_methods"`
            UserID string  `json:"user_id"`
            ExpireAt string  `json:"expire_at"`
            BulkCouponDiscount float64  `json:"bulk_coupon_discount"`
            IsDefault bool  `json:"is_default"`
            Gstin string  `json:"gstin"`
            BuyNow bool  `json:"buy_now"`
            Discount float64  `json:"discount"`
            IsActive bool  `json:"is_active"`
            OrderID string  `json:"order_id"`
            CreatedOn string  `json:"created_on"`
            LastModified string  `json:"last_modified"`
            AppID string  `json:"app_id"`
            FcIndexMap []float64  `json:"fc_index_map"`
            DeliveryCharges map[string]interface{}  `json:"delivery_charges"`
            Articles []map[string]interface{}  `json:"articles"`
            Cashback map[string]interface{}  `json:"cashback"`
            ID string  `json:"_id"`
            CodCharges map[string]interface{}  `json:"cod_charges"`
            Payments map[string]interface{}  `json:"payments"`
            Meta map[string]interface{}  `json:"meta"`
            CheckoutMode string  `json:"checkout_mode"`
            IsArchive bool  `json:"is_archive"`
            PickUpCustomerDetails map[string]interface{}  `json:"pick_up_customer_details"`
            Promotion map[string]interface{}  `json:"promotion"`
            MergeQty bool  `json:"merge_qty"`
         
    }
    
    // AbandonedCartResponse used by Cart
    type AbandonedCartResponse struct {

        
            Result map[string]interface{}  `json:"result"`
            Message string  `json:"message"`
            Page Page  `json:"page"`
            Success bool  `json:"success"`
            Items []AbandonedCart  `json:"items"`
         
    }
    
    // AddProductCart used by Cart
    type AddProductCart struct {

        
            SellerID float64  `json:"seller_id"`
            ArticleAssignment map[string]interface{}  `json:"article_assignment"`
            Quantity float64  `json:"quantity"`
            Display string  `json:"display"`
            ProductGroupTags []string  `json:"product_group_tags"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            ArticleID string  `json:"article_id"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            ItemID float64  `json:"item_id"`
            StoreID float64  `json:"store_id"`
            Pos bool  `json:"pos"`
            ItemSize string  `json:"item_size"`
         
    }
    
    // AddCartRequest used by Cart
    type AddCartRequest struct {

        
            Items []AddProductCart  `json:"items"`
         
    }
    
    // CartCurrency used by Cart
    type CartCurrency struct {

        
            Symbol string  `json:"symbol"`
            Code string  `json:"code"`
         
    }
    
    // CartDetailResponse used by Cart
    type CartDetailResponse struct {

        
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            CouponText string  `json:"coupon_text"`
            Comment string  `json:"comment"`
            ID string  `json:"id"`
            IsValid bool  `json:"is_valid"`
            LastModified string  `json:"last_modified"`
            Currency CartCurrency  `json:"currency"`
            CheckoutMode string  `json:"checkout_mode"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            Message string  `json:"message"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            Gstin string  `json:"gstin"`
            BuyNow bool  `json:"buy_now"`
            Items []CartProductInfo  `json:"items"`
         
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

        
            Identifiers CartProductIdentifer  `json:"identifiers"`
            Quantity float64  `json:"quantity"`
            ItemIndex float64  `json:"item_index"`
            ArticleID string  `json:"article_id"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            ItemID float64  `json:"item_id"`
            ItemSize string  `json:"item_size"`
         
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
    
    // E used by Rewards
    type E struct {

        
            Code map[string]interface{}  `json:"code"`
            Exception string  `json:"exception"`
            Info string  `json:"info"`
            Message string  `json:"message"`
            RequestID string  `json:"request_id"`
            StackTrace string  `json:"stack_trace"`
            Status float64  `json:"status"`
         
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
    
    // GiveawayResponse used by Rewards
    type GiveawayResponse struct {

        
            Items []Giveaway  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // HistoryPretty used by Rewards
    type HistoryPretty struct {

        
            ID string  `json:"_id"`
            ApplicationID string  `json:"application_id"`
            Claimed bool  `json:"claimed"`
            CreatedAt string  `json:"created_at"`
            ExpiresOn string  `json:"expires_on"`
            Points float64  `json:"points"`
            RemainingPoints float64  `json:"remaining_points"`
            Text1 string  `json:"text_1"`
            Text2 string  `json:"text_2"`
            Text3 string  `json:"text_3"`
            TxnName string  `json:"txn_name"`
            UpdatedAt string  `json:"updated_at"`
            UserID string  `json:"user_id"`
         
    }
    
    // HistoryRes used by Rewards
    type HistoryRes struct {

        
            Items []HistoryPretty  `json:"items"`
            Page Page  `json:"page"`
            Points float64  `json:"points"`
         
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
    
    // Points used by Rewards
    type Points struct {

        
            Available float64  `json:"available"`
         
    }
    
    // Referral used by Rewards
    type Referral struct {

        
            Code string  `json:"code"`
         
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
    
    // RewardsAudience used by Rewards
    type RewardsAudience struct {

        
            HeaderUserID string  `json:"header_user_id"`
            ID string  `json:"id"`
         
    }
    
    // RewardsRule used by Rewards
    type RewardsRule struct {

        
            Amount float64  `json:"amount"`
         
    }
    
    // Schedule used by Rewards
    type Schedule struct {

        
            Cron string  `json:"cron"`
            Duration float64  `json:"duration"`
            End string  `json:"end"`
            Start string  `json:"start"`
         
    }
    
    // ShareMessages used by Rewards
    type ShareMessages struct {

        
            Email string  `json:"email"`
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
    
