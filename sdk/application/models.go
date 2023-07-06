package application



    
    // Price ...
    type Price struct {

        
            CurrencySymbol string  `json:"currency_symbol"`
            CurrencyCode string  `json:"currency_code"`
            Max float64  `json:"max"`
            Min float64  `json:"min"`
         
    }
    
    // ProductListingPrice ...
    type ProductListingPrice struct {

        
            Effective Price  `json:"effective"`
            Marked Price  `json:"marked"`
         
    }
    
    // ProductListingActionPage ...
    type ProductListingActionPage struct {

        
            Params map[string]interface{}  `json:"params"`
            Query map[string]interface{}  `json:"query"`
            Type string  `json:"type"`
         
    }
    
    // ProductListingAction ...
    type ProductListingAction struct {

        
            Type string  `json:"type"`
            Page ProductListingActionPage  `json:"page"`
         
    }
    
    // CustomMetaFields ...
    type CustomMetaFields struct {

        
            Value string  `json:"value"`
            Key string  `json:"key"`
         
    }
    
    // ApplicationItemSEO ...
    type ApplicationItemSEO struct {

        
            Title string  `json:"title"`
            Description string  `json:"description"`
         
    }
    
    // Meta ...
    type Meta struct {

        
            Source string  `json:"source"`
         
    }
    
    // Media ...
    type Media struct {

        
            Alt string  `json:"alt"`
            Type string  `json:"type"`
            Meta Meta  `json:"meta"`
            URL string  `json:"url"`
         
    }
    
    // ProductBrand ...
    type ProductBrand struct {

        
            Action ProductListingAction  `json:"action"`
            UID float64  `json:"uid"`
            Logo Media  `json:"logo"`
            Name string  `json:"name"`
            Description string  `json:"description"`
         
    }
    
    // ProductCategoryMap ...
    type ProductCategoryMap struct {

        
            L1 ProductBrand  `json:"l1"`
            L2 ProductBrand  `json:"l2"`
            L3 ProductBrand  `json:"l3"`
         
    }
    
    // ApplicationItemMOQ ...
    type ApplicationItemMOQ struct {

        
            Minimum float64  `json:"minimum"`
            IncrementUnit float64  `json:"increment_unit"`
            Maximum float64  `json:"maximum"`
         
    }
    
    // NetQuantity ...
    type NetQuantity struct {

        
            Value float64  `json:"value"`
            Unit string  `json:"unit"`
         
    }
    
    // ProductDetailCustomOrder ...
    type ProductDetailCustomOrder struct {

        
            ManufacturingTimeUnit string  `json:"manufacturing_time_unit"`
            ManufacturingTime float64  `json:"manufacturing_time"`
            IsCustomOrder bool  `json:"is_custom_order"`
         
    }
    
    // ProductDetailAttribute ...
    type ProductDetailAttribute struct {

        
            Value string  `json:"value"`
            Type string  `json:"type"`
            Key string  `json:"key"`
         
    }
    
    // ProductDetailGroupedAttribute ...
    type ProductDetailGroupedAttribute struct {

        
            Title string  `json:"title"`
            Details []ProductDetailAttribute  `json:"details"`
         
    }
    
    // ProductDetail ...
    type ProductDetail struct {

        
            Price ProductListingPrice  `json:"price"`
            Action ProductListingAction  `json:"action"`
            HasVariant bool  `json:"has_variant"`
            ShortDescription string  `json:"short_description"`
            Discount string  `json:"discount"`
            CustomMeta []CustomMetaFields  `json:"_custom_meta"`
            Tags []string  `json:"tags"`
            Rating float64  `json:"rating"`
            Seo ApplicationItemSEO  `json:"seo"`
            Description string  `json:"description"`
            Slug string  `json:"slug"`
            CategoryMap ProductCategoryMap  `json:"category_map"`
            Highlights []string  `json:"highlights"`
            IsDependent bool  `json:"is_dependent"`
            Moq ApplicationItemMOQ  `json:"moq"`
            ItemType string  `json:"item_type"`
            TeaserTag string  `json:"teaser_tag"`
            Type string  `json:"type"`
            ImageNature string  `json:"image_nature"`
            Tryouts []string  `json:"tryouts"`
            Categories []ProductBrand  `json:"categories"`
            ProductOnlineDate string  `json:"product_online_date"`
            ItemCode string  `json:"item_code"`
            RatingCount float64  `json:"rating_count"`
            Attributes map[string]interface{}  `json:"attributes"`
            NetQuantity NetQuantity  `json:"net_quantity"`
            Brand ProductBrand  `json:"brand"`
            ProductGroupTag []string  `json:"product_group_tag"`
            Medias []Media  `json:"medias"`
            Color string  `json:"color"`
            Name string  `json:"name"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            CustomOrder ProductDetailCustomOrder  `json:"custom_order"`
            Similars []string  `json:"similars"`
            UID float64  `json:"uid"`
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
         
    }
    
    // ErrorResponse ...
    type ErrorResponse struct {

        
            Error string  `json:"error"`
         
    }
    
    // ProductSizeStores ...
    type ProductSizeStores struct {

        
            Count float64  `json:"count"`
         
    }
    
    // ColumnHeader ...
    type ColumnHeader struct {

        
            Value string  `json:"value"`
            Convertable bool  `json:"convertable"`
         
    }
    
    // ColumnHeaders ...
    type ColumnHeaders struct {

        
            Col6 ColumnHeader  `json:"col_6"`
            Col1 ColumnHeader  `json:"col_1"`
            Col3 ColumnHeader  `json:"col_3"`
            Col2 ColumnHeader  `json:"col_2"`
            Col4 ColumnHeader  `json:"col_4"`
            Col5 ColumnHeader  `json:"col_5"`
         
    }
    
    // SizeChartValues ...
    type SizeChartValues struct {

        
            Col6 string  `json:"col_6"`
            Col1 string  `json:"col_1"`
            Col3 string  `json:"col_3"`
            Col2 string  `json:"col_2"`
            Col4 string  `json:"col_4"`
            Col5 string  `json:"col_5"`
         
    }
    
    // SizeChart ...
    type SizeChart struct {

        
            Image string  `json:"image"`
            SizeTip string  `json:"size_tip"`
            Title string  `json:"title"`
            Headers ColumnHeaders  `json:"headers"`
            Sizes []SizeChartValues  `json:"sizes"`
            Unit string  `json:"unit"`
            Description string  `json:"description"`
         
    }
    
    // Weight ...
    type Weight struct {

        
            IsDefault bool  `json:"is_default"`
            Unit string  `json:"unit"`
            Shipping float64  `json:"shipping"`
         
    }
    
    // Dimension ...
    type Dimension struct {

        
            IsDefault bool  `json:"is_default"`
            Height float64  `json:"height"`
            Length float64  `json:"length"`
            Unit string  `json:"unit"`
            Width float64  `json:"width"`
         
    }
    
    // ProductSize ...
    type ProductSize struct {

        
            Display string  `json:"display"`
            Quantity float64  `json:"quantity"`
            Weight Weight  `json:"weight"`
            Dimension Dimension  `json:"dimension"`
            SellerIdentifiers []string  `json:"seller_identifiers"`
            Value string  `json:"value"`
            IsAvailable bool  `json:"is_available"`
         
    }
    
    // ProductSizes ...
    type ProductSizes struct {

        
            Price ProductListingPrice  `json:"price"`
            Stores ProductSizeStores  `json:"stores"`
            Discount string  `json:"discount"`
            SizeChart SizeChart  `json:"size_chart"`
            Sizes []ProductSize  `json:"sizes"`
            Sellable bool  `json:"sellable"`
            MultiSize bool  `json:"multi_size"`
         
    }
    
    // AttributeDetail ...
    type AttributeDetail struct {

        
            Display string  `json:"display"`
            Description string  `json:"description"`
            Logo string  `json:"logo"`
            Key string  `json:"key"`
         
    }
    
    // AttributeMetadata ...
    type AttributeMetadata struct {

        
            Title string  `json:"title"`
            Details []AttributeDetail  `json:"details"`
         
    }
    
    // ProductsComparisonResponse ...
    type ProductsComparisonResponse struct {

        
            AttributesMetadata []AttributeMetadata  `json:"attributes_metadata"`
            Items []ProductDetail  `json:"items"`
         
    }
    
    // ProductCompareResponse ...
    type ProductCompareResponse struct {

        
            Title string  `json:"title"`
            Items []ProductDetail  `json:"items"`
            AttributesMetadata []AttributeMetadata  `json:"attributes_metadata"`
            Subtitle string  `json:"subtitle"`
         
    }
    
    // ProductFrequentlyComparedSimilarResponse ...
    type ProductFrequentlyComparedSimilarResponse struct {

        
            Similars ProductCompareResponse  `json:"similars"`
         
    }
    
    // ProductVariantItemResponse ...
    type ProductVariantItemResponse struct {

        
            Action ProductListingAction  `json:"action"`
            Medias []Media  `json:"medias"`
            CustomMeta []CustomMetaFields  `json:"_custom_meta"`
            Color string  `json:"color"`
            Name string  `json:"name"`
            ColorName string  `json:"color_name"`
            Value string  `json:"value"`
            IsAvailable bool  `json:"is_available"`
            UID float64  `json:"uid"`
            Slug string  `json:"slug"`
         
    }
    
    // ProductVariantResponse ...
    type ProductVariantResponse struct {

        
            Items []ProductVariantItemResponse  `json:"items"`
            Key string  `json:"key"`
            DisplayType string  `json:"display_type"`
            Header string  `json:"header"`
         
    }
    
    // ProductVariantsResponse ...
    type ProductVariantsResponse struct {

        
            Variants []ProductVariantResponse  `json:"variants"`
         
    }
    
    // ProductStockPrice ...
    type ProductStockPrice struct {

        
            Effective float64  `json:"effective"`
            Marked float64  `json:"marked"`
            Currency string  `json:"currency"`
         
    }
    
    // Seller ...
    type Seller struct {

        
            Name string  `json:"name"`
            Count float64  `json:"count"`
            UID float64  `json:"uid"`
         
    }
    
    // CompanyDetail ...
    type CompanyDetail struct {

        
            Name string  `json:"name"`
            ID float64  `json:"id"`
         
    }
    
    // StoreDetail ...
    type StoreDetail struct {

        
            Name string  `json:"name"`
            City string  `json:"city"`
            Code string  `json:"code"`
            ID float64  `json:"id"`
         
    }
    
    // ProductStockStatusItem ...
    type ProductStockStatusItem struct {

        
            Price ProductStockPrice  `json:"price"`
            Size string  `json:"size"`
            Seller Seller  `json:"seller"`
            Quantity float64  `json:"quantity"`
            ItemID float64  `json:"item_id"`
            Identifier map[string]interface{}  `json:"identifier"`
            Company CompanyDetail  `json:"company"`
            Store StoreDetail  `json:"store"`
            UID string  `json:"uid"`
         
    }
    
    // ProductStockStatusResponse ...
    type ProductStockStatusResponse struct {

        
            Items []ProductStockStatusItem  `json:"items"`
         
    }
    
    // Page ...
    type Page struct {

        
            ItemTotal float64  `json:"item_total"`
            NextID string  `json:"next_id"`
            HasPrevious bool  `json:"has_previous"`
            HasNext bool  `json:"has_next"`
            Current float64  `json:"current"`
            Type string  `json:"type"`
            Size float64  `json:"size"`
         
    }
    
    // ProductStockPolling ...
    type ProductStockPolling struct {

        
            Page Page  `json:"page"`
            Items []ProductStockStatusItem  `json:"items"`
         
    }
    
    // ProductSortOn ...
    type ProductSortOn struct {

        
            Name string  `json:"name"`
            IsSelected bool  `json:"is_selected"`
            Value string  `json:"value"`
         
    }
    
    // ProductFiltersValue ...
    type ProductFiltersValue struct {

        
            CurrencyCode string  `json:"currency_code"`
            Min float64  `json:"min"`
            SelectedMin float64  `json:"selected_min"`
            Display string  `json:"display"`
            IsSelected bool  `json:"is_selected"`
            Max float64  `json:"max"`
            CurrencySymbol string  `json:"currency_symbol"`
            SelectedMax float64  `json:"selected_max"`
            QueryFormat string  `json:"query_format"`
            Value string  `json:"value"`
            Count float64  `json:"count"`
            DisplayFormat string  `json:"display_format"`
         
    }
    
    // ProductFiltersKey ...
    type ProductFiltersKey struct {

        
            Name string  `json:"name"`
            Display string  `json:"display"`
            Logo string  `json:"logo"`
            Kind string  `json:"kind"`
         
    }
    
    // ProductFilters ...
    type ProductFilters struct {

        
            Values []ProductFiltersValue  `json:"values"`
            Key ProductFiltersKey  `json:"key"`
         
    }
    
    // ProductVariantListingResponse ...
    type ProductVariantListingResponse struct {

        
            Items []ProductVariantItemResponse  `json:"items"`
            DisplayType string  `json:"display_type"`
            Key string  `json:"key"`
            Total float64  `json:"total"`
            Header string  `json:"header"`
         
    }
    
    // ProductListingDetail ...
    type ProductListingDetail struct {

        
            Price ProductListingPrice  `json:"price"`
            Action ProductListingAction  `json:"action"`
            HasVariant bool  `json:"has_variant"`
            ShortDescription string  `json:"short_description"`
            Discount string  `json:"discount"`
            CustomMeta []CustomMetaFields  `json:"_custom_meta"`
            Tags []string  `json:"tags"`
            Rating float64  `json:"rating"`
            Seo ApplicationItemSEO  `json:"seo"`
            Description string  `json:"description"`
            Slug string  `json:"slug"`
            CategoryMap ProductCategoryMap  `json:"category_map"`
            Highlights []string  `json:"highlights"`
            IsDependent bool  `json:"is_dependent"`
            Moq ApplicationItemMOQ  `json:"moq"`
            ItemType string  `json:"item_type"`
            Sizes []string  `json:"sizes"`
            TeaserTag string  `json:"teaser_tag"`
            Type string  `json:"type"`
            ImageNature string  `json:"image_nature"`
            Tryouts []string  `json:"tryouts"`
            Categories []ProductBrand  `json:"categories"`
            ProductOnlineDate string  `json:"product_online_date"`
            ItemCode string  `json:"item_code"`
            RatingCount float64  `json:"rating_count"`
            Attributes map[string]interface{}  `json:"attributes"`
            Variants []ProductVariantListingResponse  `json:"variants"`
            Identifiers []string  `json:"identifiers"`
            NetQuantity NetQuantity  `json:"net_quantity"`
            Brand ProductBrand  `json:"brand"`
            ProductGroupTag []string  `json:"product_group_tag"`
            Sellable bool  `json:"sellable"`
            Medias []Media  `json:"medias"`
            Color string  `json:"color"`
            Name string  `json:"name"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            CustomOrder ProductDetailCustomOrder  `json:"custom_order"`
            Similars []string  `json:"similars"`
            UID float64  `json:"uid"`
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
         
    }
    
    // ProductListingResponse ...
    type ProductListingResponse struct {

        
            SortOn []ProductSortOn  `json:"sort_on"`
            Filters []ProductFilters  `json:"filters"`
            Page Page  `json:"page"`
            Items []ProductListingDetail  `json:"items"`
         
    }
    
    // ImageUrls ...
    type ImageUrls struct {

        
            Portrait Media  `json:"portrait"`
            Landscape Media  `json:"landscape"`
         
    }
    
    // BrandItem ...
    type BrandItem struct {

        
            Action ProductListingAction  `json:"action"`
            Logo Media  `json:"logo"`
            Banners ImageUrls  `json:"banners"`
            Discount string  `json:"discount"`
            Description string  `json:"description"`
            Departments []string  `json:"departments"`
            Name string  `json:"name"`
            UID float64  `json:"uid"`
            Slug string  `json:"slug"`
         
    }
    
    // BrandListingResponse ...
    type BrandListingResponse struct {

        
            Page Page  `json:"page"`
            Items []BrandItem  `json:"items"`
         
    }
    
    // BrandDetailResponse ...
    type BrandDetailResponse struct {

        
            Logo Media  `json:"logo"`
            Banners ImageUrls  `json:"banners"`
            Description string  `json:"description"`
            Name string  `json:"name"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            UID float64  `json:"uid"`
         
    }
    
    // ThirdLevelChild ...
    type ThirdLevelChild struct {

        
            Action ProductListingAction  `json:"action"`
            Childs []map[string]interface{}  `json:"childs"`
            Banners ImageUrls  `json:"banners"`
            Name string  `json:"name"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            UID float64  `json:"uid"`
            Slug string  `json:"slug"`
         
    }
    
    // SecondLevelChild ...
    type SecondLevelChild struct {

        
            Action ProductListingAction  `json:"action"`
            Childs []ThirdLevelChild  `json:"childs"`
            Banners ImageUrls  `json:"banners"`
            Name string  `json:"name"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            UID float64  `json:"uid"`
            Slug string  `json:"slug"`
         
    }
    
    // Child ...
    type Child struct {

        
            Action ProductListingAction  `json:"action"`
            Childs []SecondLevelChild  `json:"childs"`
            Banners ImageUrls  `json:"banners"`
            Name string  `json:"name"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            UID float64  `json:"uid"`
            Slug string  `json:"slug"`
         
    }
    
    // CategoryBanner ...
    type CategoryBanner struct {

        
            Portrait Media  `json:"portrait"`
            Landscape Media  `json:"landscape"`
         
    }
    
    // CategoryItems ...
    type CategoryItems struct {

        
            Action ProductListingAction  `json:"action"`
            Childs []Child  `json:"childs"`
            Banners CategoryBanner  `json:"banners"`
            Name string  `json:"name"`
            UID float64  `json:"uid"`
            Slug string  `json:"slug"`
         
    }
    
    // DepartmentCategoryTree ...
    type DepartmentCategoryTree struct {

        
            Items []CategoryItems  `json:"items"`
            Department string  `json:"department"`
         
    }
    
    // DepartmentIdentifier ...
    type DepartmentIdentifier struct {

        
            UID float64  `json:"uid"`
            Slug string  `json:"slug"`
         
    }
    
    // CategoryListingResponse ...
    type CategoryListingResponse struct {

        
            Data []DepartmentCategoryTree  `json:"data"`
            Departments []DepartmentIdentifier  `json:"departments"`
         
    }
    
    // CategoryMetaResponse ...
    type CategoryMetaResponse struct {

        
            Logo Media  `json:"logo"`
            Banners ImageUrls  `json:"banners"`
            Name string  `json:"name"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            UID float64  `json:"uid"`
         
    }
    
    // HomeListingResponse ...
    type HomeListingResponse struct {

        
            Message string  `json:"message"`
            Page Page  `json:"page"`
            Items []ProductListingDetail  `json:"items"`
         
    }
    
    // Department ...
    type Department struct {

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            Logo Media  `json:"logo"`
            Slug string  `json:"slug"`
            PriorityOrder float64  `json:"priority_order"`
         
    }
    
    // DepartmentResponse ...
    type DepartmentResponse struct {

        
            Items []Department  `json:"items"`
         
    }
    
    // AutocompleteItem ...
    type AutocompleteItem struct {

        
            Action ProductListingAction  `json:"action"`
            Display string  `json:"display"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Type string  `json:"type"`
            Logo Media  `json:"logo"`
         
    }
    
    // AutoCompleteResponse ...
    type AutoCompleteResponse struct {

        
            Items []AutocompleteItem  `json:"items"`
         
    }
    
    // CollectionListingFilterTag ...
    type CollectionListingFilterTag struct {

        
            Name string  `json:"name"`
            Display string  `json:"display"`
            IsSelected bool  `json:"is_selected"`
         
    }
    
    // CollectionListingFilterType ...
    type CollectionListingFilterType struct {

        
            Name string  `json:"name"`
            Display string  `json:"display"`
            IsSelected bool  `json:"is_selected"`
         
    }
    
    // CollectionListingFilter ...
    type CollectionListingFilter struct {

        
            Tags []CollectionListingFilterTag  `json:"tags"`
            Type []CollectionListingFilterType  `json:"type"`
         
    }
    
    // CollectionQuery ...
    type CollectionQuery struct {

        
            Value []interface{}  `json:"value"`
            Attribute string  `json:"attribute"`
            Op string  `json:"op"`
         
    }
    
    // GetCollectionDetailNest ...
    type GetCollectionDetailNest struct {

        
            Action ProductListingAction  `json:"action"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Meta map[string]interface{}  `json:"meta"`
            AllowSort bool  `json:"allow_sort"`
            Logo Media  `json:"logo"`
            Slug string  `json:"slug"`
            Schedule map[string]interface{}  `json:"_schedule"`
            IsActive bool  `json:"is_active"`
            Description string  `json:"description"`
            Banners ImageUrls  `json:"banners"`
            AllowFacets bool  `json:"allow_facets"`
            Type string  `json:"type"`
            Query []CollectionQuery  `json:"query"`
            Tag []string  `json:"tag"`
            Priority float64  `json:"priority"`
            Badge map[string]interface{}  `json:"badge"`
            AppID string  `json:"app_id"`
            SortOn string  `json:"sort_on"`
            Name string  `json:"name"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Cron map[string]interface{}  `json:"cron"`
            UID string  `json:"uid"`
         
    }
    
    // GetCollectionListingResponse ...
    type GetCollectionListingResponse struct {

        
            Filters CollectionListingFilter  `json:"filters"`
            Page Page  `json:"page"`
            Items []GetCollectionDetailNest  `json:"items"`
         
    }
    
    // CollectionDetailResponse ...
    type CollectionDetailResponse struct {

        
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Meta map[string]interface{}  `json:"meta"`
            AllowSort bool  `json:"allow_sort"`
            Logo Media  `json:"logo"`
            Slug string  `json:"slug"`
            Schedule map[string]interface{}  `json:"_schedule"`
            IsActive bool  `json:"is_active"`
            Description string  `json:"description"`
            Banners ImageUrls  `json:"banners"`
            AllowFacets bool  `json:"allow_facets"`
            Type string  `json:"type"`
            Query []CollectionQuery  `json:"query"`
            Tag []string  `json:"tag"`
            Priority float64  `json:"priority"`
            Badge map[string]interface{}  `json:"badge"`
            AppID string  `json:"app_id"`
            SortOn string  `json:"sort_on"`
            Name string  `json:"name"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Cron map[string]interface{}  `json:"cron"`
         
    }
    
    // GetFollowListingResponse ...
    type GetFollowListingResponse struct {

        
            Page Page  `json:"page"`
            Items []ProductListingDetail  `json:"items"`
         
    }
    
    // FollowPostResponse ...
    type FollowPostResponse struct {

        
            Message string  `json:"message"`
            ID string  `json:"id"`
         
    }
    
    // FollowerCountResponse ...
    type FollowerCountResponse struct {

        
            Count float64  `json:"count"`
         
    }
    
    // FollowIdsData ...
    type FollowIdsData struct {

        
            Products []float64  `json:"products"`
            Brands []float64  `json:"brands"`
            Collections []float64  `json:"collections"`
         
    }
    
    // FollowIdsResponse ...
    type FollowIdsResponse struct {

        
            Data FollowIdsData  `json:"data"`
         
    }
    
    // LatLong ...
    type LatLong struct {

        
            Coordinates []float64  `json:"coordinates"`
            Type string  `json:"type"`
         
    }
    
    // Store ...
    type Store struct {

        
            Country string  `json:"country"`
            Name string  `json:"name"`
            StoreEmail string  `json:"store_email"`
            Address string  `json:"address"`
            Pincode float64  `json:"pincode"`
            State string  `json:"state"`
            LatLong LatLong  `json:"lat_long"`
            StoreCode string  `json:"store_code"`
            City string  `json:"city"`
            UID float64  `json:"uid"`
         
    }
    
    // StoreListingResponse ...
    type StoreListingResponse struct {

        
            Page Page  `json:"page"`
            Items []Store  `json:"items"`
         
    }
    
    // SellerPhoneNumber ...
    type SellerPhoneNumber struct {

        
            Number string  `json:"number"`
            CountryCode float64  `json:"country_code"`
         
    }
    
    // StoreManagerSerializer ...
    type StoreManagerSerializer struct {

        
            Name string  `json:"name"`
            MobileNo SellerPhoneNumber  `json:"mobile_no"`
            Email string  `json:"email"`
         
    }
    
    // StoreDepartments ...
    type StoreDepartments struct {

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            Logo map[string]interface{}  `json:"logo"`
            Slug string  `json:"slug"`
            PriorityOrder float64  `json:"priority_order"`
         
    }
    
    // StoreAddressSerializer ...
    type StoreAddressSerializer struct {

        
            Address1 string  `json:"address1"`
            Country string  `json:"country"`
            Landmark string  `json:"landmark"`
            Address2 string  `json:"address2"`
            Longitude float64  `json:"longitude"`
            Pincode float64  `json:"pincode"`
            State string  `json:"state"`
            Latitude float64  `json:"latitude"`
            City string  `json:"city"`
         
    }
    
    // CompanyStore ...
    type CompanyStore struct {

        
            Name string  `json:"name"`
            CompanyType string  `json:"company_type"`
            UID float64  `json:"uid"`
            BusinessType string  `json:"business_type"`
         
    }
    
    // AppStore ...
    type AppStore struct {

        
            Manager StoreManagerSerializer  `json:"manager"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            Departments []StoreDepartments  `json:"departments"`
            Name string  `json:"name"`
            Address StoreAddressSerializer  `json:"address"`
            Company CompanyStore  `json:"company"`
            StoreCode string  `json:"store_code"`
            UID float64  `json:"uid"`
         
    }
    
    // ApplicationStoreListing ...
    type ApplicationStoreListing struct {

        
            Filters []map[string]interface{}  `json:"filters"`
            Page Page  `json:"page"`
            Items []AppStore  `json:"items"`
         
    }
    
    // Time ...
    type Time struct {

        
            Minute float64  `json:"minute"`
            Hour float64  `json:"hour"`
         
    }
    
    // StoreTiming ...
    type StoreTiming struct {

        
            Opening Time  `json:"opening"`
            Closing Time  `json:"closing"`
            Open bool  `json:"open"`
            Weekday string  `json:"weekday"`
         
    }
    
    // StoreDetails ...
    type StoreDetails struct {

        
            Manager StoreManagerSerializer  `json:"manager"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            Timing []StoreTiming  `json:"timing"`
            Departments []StoreDepartments  `json:"departments"`
            Name string  `json:"name"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Address StoreAddressSerializer  `json:"address"`
            Company CompanyStore  `json:"company"`
            StoreCode string  `json:"store_code"`
            UID float64  `json:"uid"`
         
    }
    
    // UserDetail ...
    type UserDetail struct {

        
            Contact string  `json:"contact"`
            SuperUser bool  `json:"super_user"`
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
         
    }
    
    // ProductGroupPrice ...
    type ProductGroupPrice struct {

        
            MaxEffective float64  `json:"max_effective"`
            Currency string  `json:"currency"`
            MaxMarked float64  `json:"max_marked"`
            MinMarked float64  `json:"min_marked"`
            MinEffective float64  `json:"min_effective"`
         
    }
    
    // ProductDetails ...
    type ProductDetails struct {

        
            HasVariant bool  `json:"has_variant"`
            ShortDescription string  `json:"short_description"`
            Rating float64  `json:"rating"`
            Media []map[string]interface{}  `json:"media"`
            Slug string  `json:"slug"`
            Description string  `json:"description"`
            Highlights []string  `json:"highlights"`
            TemplateTag string  `json:"template_tag"`
            HsnCode float64  `json:"hsn_code"`
            Identifier map[string]interface{}  `json:"identifier"`
            ImageNature string  `json:"image_nature"`
            ItemCode string  `json:"item_code"`
            CountryOfOrigin string  `json:"country_of_origin"`
            OutOfStock bool  `json:"out_of_stock"`
            RatingCount float64  `json:"rating_count"`
            Attributes map[string]interface{}  `json:"attributes"`
            BrandUID float64  `json:"brand_uid"`
            GroupedAttributes map[string]interface{}  `json:"grouped_attributes"`
            Name string  `json:"name"`
            Images []string  `json:"images"`
            IsSet bool  `json:"is_set"`
         
    }
    
    // Size ...
    type Size struct {

        
            Value string  `json:"value"`
            Display string  `json:"display"`
            IsAvailable bool  `json:"is_available"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // ProductInGroup ...
    type ProductInGroup struct {

        
            Price ProductGroupPrice  `json:"price"`
            ProductDetails ProductDetails  `json:"product_details"`
            MaxQuantity float64  `json:"max_quantity"`
            AutoAddToCart bool  `json:"auto_add_to_cart"`
            ProductUID float64  `json:"product_uid"`
            AutoSelect bool  `json:"auto_select"`
            MinQuantity float64  `json:"min_quantity"`
            AllowRemove bool  `json:"allow_remove"`
            Sizes []Size  `json:"sizes"`
         
    }
    
    // ProductGroupingModel ...
    type ProductGroupingModel struct {

        
            VerifiedOn string  `json:"verified_on"`
            VerifiedBy UserDetail  `json:"verified_by"`
            IsActive bool  `json:"is_active"`
            ModifiedBy UserDetail  `json:"modified_by"`
            Choice string  `json:"choice"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            ID interface{}  `json:"_id"`
            Products []ProductInGroup  `json:"products"`
            CreatedBy UserDetail  `json:"created_by"`
            Name string  `json:"name"`
            CompanyID float64  `json:"company_id"`
            CreatedOn string  `json:"created_on"`
            PageVisibility []string  `json:"page_visibility"`
            Meta map[string]interface{}  `json:"meta"`
            ModifiedOn string  `json:"modified_on"`
            Logo string  `json:"logo"`
            Slug string  `json:"slug"`
         
    }
    
    // ProductBundle ...
    type ProductBundle struct {

        
            Items []ProductGroupingModel  `json:"items"`
         
    }
    
    // ProductStockPriceV3 ...
    type ProductStockPriceV3 struct {

        
            Currency string  `json:"currency"`
            Effective float64  `json:"effective"`
            Marked float64  `json:"marked"`
            Selling float64  `json:"selling"`
         
    }
    
    // ArticleAssignmentV3 ...
    type ArticleAssignmentV3 struct {

        
            Level string  `json:"level"`
            Strategy string  `json:"strategy"`
         
    }
    
    // StrategyWiseListingSchemaV3 ...
    type StrategyWiseListingSchemaV3 struct {

        
            Distance float64  `json:"distance"`
            Quantity float64  `json:"quantity"`
            Pincode float64  `json:"pincode"`
            Tat float64  `json:"tat"`
         
    }
    
    // DetailsSchemaV3 ...
    type DetailsSchemaV3 struct {

        
            Value string  `json:"value"`
            Type string  `json:"type"`
            Key string  `json:"key"`
         
    }
    
    // MarketPlaceSttributesSchemaV3 ...
    type MarketPlaceSttributesSchemaV3 struct {

        
            Title string  `json:"title"`
            Details []DetailsSchemaV3  `json:"details"`
         
    }
    
    // DiscountMeta ...
    type DiscountMeta struct {

        
            NumberOfMinutes float64  `json:"number_of_minutes"`
            Start string  `json:"start"`
            End string  `json:"end"`
            Timer bool  `json:"timer"`
         
    }
    
    // ProductSetDistributionSizeV3 ...
    type ProductSetDistributionSizeV3 struct {

        
            Size string  `json:"size"`
            Pieces float64  `json:"pieces"`
         
    }
    
    // ProductSetDistributionV3 ...
    type ProductSetDistributionV3 struct {

        
            Sizes []ProductSetDistributionSizeV3  `json:"sizes"`
         
    }
    
    // ProductSetV3 ...
    type ProductSetV3 struct {

        
            Quantity float64  `json:"quantity"`
            SizeDistribution ProductSetDistributionV3  `json:"size_distribution"`
         
    }
    
    // ProductStockUnitPriceV3 ...
    type ProductStockUnitPriceV3 struct {

        
            Price float64  `json:"price"`
            CurrencyCode string  `json:"currency_code"`
            Unit string  `json:"unit"`
            CurrencySymbol string  `json:"currency_symbol"`
         
    }
    
    // ReturnConfigSchemaV3 ...
    type ReturnConfigSchemaV3 struct {

        
            Time float64  `json:"time"`
            Unit string  `json:"unit"`
            Returnable bool  `json:"returnable"`
         
    }
    
    // SellerV3 ...
    type SellerV3 struct {

        
            Name string  `json:"name"`
            Count float64  `json:"count"`
            UID float64  `json:"uid"`
         
    }
    
    // StoreV3 ...
    type StoreV3 struct {

        
            Name string  `json:"name"`
            Count float64  `json:"count"`
            UID float64  `json:"uid"`
         
    }
    
    // SellerGroupAttributes ...
    type SellerGroupAttributes struct {

        
            Title string  `json:"title"`
            Details []DetailsSchemaV3  `json:"details"`
         
    }
    
    // ProductSizePriceResponseV3 ...
    type ProductSizePriceResponseV3 struct {

        
            Price ProductStockPriceV3  `json:"price"`
            PricePerPiece ProductStockPriceV3  `json:"price_per_piece"`
            ArticleAssignment ArticleAssignmentV3  `json:"article_assignment"`
            SpecialBadge string  `json:"special_badge"`
            Discount string  `json:"discount"`
            StrategyWiseListing []StrategyWiseListingSchemaV3  `json:"strategy_wise_listing"`
            Pincode float64  `json:"pincode"`
            MarketplaceAttributes []MarketPlaceSttributesSchemaV3  `json:"marketplace_attributes"`
            LongLat []float64  `json:"long_lat"`
            DiscountMeta DiscountMeta  `json:"discount_meta"`
            Set ProductSetV3  `json:"set"`
            ItemType string  `json:"item_type"`
            PricePerUnit ProductStockUnitPriceV3  `json:"price_per_unit"`
            ReturnConfig ReturnConfigSchemaV3  `json:"return_config"`
            Seller SellerV3  `json:"seller"`
            Quantity float64  `json:"quantity"`
            Store StoreV3  `json:"store"`
            IsCod bool  `json:"is_cod"`
            SellerCount float64  `json:"seller_count"`
            ArticleID string  `json:"article_id"`
            IsGift bool  `json:"is_gift"`
            GroupedAttributes []SellerGroupAttributes  `json:"grouped_attributes"`
         
    }
    
    // ProductSizeSellerFilterSchemaV3 ...
    type ProductSizeSellerFilterSchemaV3 struct {

        
            Name string  `json:"name"`
            IsSelected bool  `json:"is_selected"`
            Value string  `json:"value"`
         
    }
    
    // ProductSizeSellersResponseV3 ...
    type ProductSizeSellersResponseV3 struct {

        
            SortOn []ProductSizeSellerFilterSchemaV3  `json:"sort_on"`
            Page Page  `json:"page"`
            Items []ProductSizePriceResponseV3  `json:"items"`
         
    }
    

    
    // PromoMeta ...
    type PromoMeta struct {

        
            Message string  `json:"message"`
         
    }
    
    // CartProductIdentifer ...
    type CartProductIdentifer struct {

        
            Identifier string  `json:"identifier"`
         
    }
    
    // NetQuantity ...
    type NetQuantity struct {

        
            Unit string  `json:"unit"`
            Value string  `json:"value"`
         
    }
    
    // BaseInfo ...
    type BaseInfo struct {

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // ProductImage ...
    type ProductImage struct {

        
            URL string  `json:"url"`
            AspectRatio string  `json:"aspect_ratio"`
            SecureURL string  `json:"secure_url"`
         
    }
    
    // ActionQuery ...
    type ActionQuery struct {

        
            ProductSlug []string  `json:"product_slug"`
         
    }
    
    // ProductAction ...
    type ProductAction struct {

        
            Type string  `json:"type"`
            URL string  `json:"url"`
            Query ActionQuery  `json:"query"`
         
    }
    
    // Tags ...
    type Tags struct {

        
            Tags map[string]interface{}  `json:"tags"`
         
    }
    
    // CategoryInfo ...
    type CategoryInfo struct {

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // CartProduct ...
    type CartProduct struct {

        
            ItemCode string  `json:"item_code"`
            NetQuantity NetQuantity  `json:"net_quantity"`
            Slug string  `json:"slug"`
            Tags []string  `json:"tags"`
            UID float64  `json:"uid"`
            Brand BaseInfo  `json:"brand"`
            Type string  `json:"type"`
            Images []ProductImage  `json:"images"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Action ProductAction  `json:"action"`
            TeaserTag Tags  `json:"teaser_tag"`
            Name string  `json:"name"`
            Categories []CategoryInfo  `json:"categories"`
         
    }
    
    // ProductAvailabilitySize ...
    type ProductAvailabilitySize struct {

        
            Meta map[string]interface{}  `json:"meta"`
            Value string  `json:"value"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            IsAvailable bool  `json:"is_available"`
            Display string  `json:"display"`
         
    }
    
    // ProductAvailability ...
    type ProductAvailability struct {

        
            OutOfStock bool  `json:"out_of_stock"`
            AvailableSizes []ProductAvailabilitySize  `json:"available_sizes"`
            IsValid bool  `json:"is_valid"`
            Deliverable bool  `json:"deliverable"`
            OtherStoreQuantity float64  `json:"other_store_quantity"`
            Sizes []string  `json:"sizes"`
         
    }
    
    // Ownership ...
    type Ownership struct {

        
            PayableCategory string  `json:"payable_category"`
            PayableBy string  `json:"payable_by"`
         
    }
    
    // FreeGiftItem ...
    type FreeGiftItem struct {

        
            ItemName string  `json:"item_name"`
            ItemID float64  `json:"item_id"`
            ItemImagesURL []string  `json:"item_images_url"`
            ItemPriceDetails map[string]interface{}  `json:"item_price_details"`
            ItemBrandName string  `json:"item_brand_name"`
            ItemSlug string  `json:"item_slug"`
         
    }
    
    // AppliedFreeArticles ...
    type AppliedFreeArticles struct {

        
            FreeGiftItemDetails FreeGiftItem  `json:"free_gift_item_details"`
            ArticleID string  `json:"article_id"`
            ParentItemIdentifier string  `json:"parent_item_identifier"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // BuyRules ...
    type BuyRules struct {

        
            CartConditions map[string]interface{}  `json:"cart_conditions"`
            ItemCriteria map[string]interface{}  `json:"item_criteria"`
         
    }
    
    // DiscountRulesApp ...
    type DiscountRulesApp struct {

        
            RawOffer map[string]interface{}  `json:"raw_offer"`
            MatchedBuyRules []string  `json:"matched_buy_rules"`
            ItemCriteria map[string]interface{}  `json:"item_criteria"`
            Offer map[string]interface{}  `json:"offer"`
         
    }
    
    // AppliedPromotion ...
    type AppliedPromotion struct {

        
            Ownership Ownership  `json:"ownership"`
            Amount float64  `json:"amount"`
            PromotionType string  `json:"promotion_type"`
            AppliedFreeArticles []AppliedFreeArticles  `json:"applied_free_articles"`
            MrpPromotion bool  `json:"mrp_promotion"`
            BuyRules []BuyRules  `json:"buy_rules"`
            PromotionName string  `json:"promotion_name"`
            ArticleQuantity float64  `json:"article_quantity"`
            PromotionOfferText string  `json:"promotion_offer_text"`
            PromotionGroup string  `json:"promotion_group"`
            PromoID string  `json:"promo_id"`
            DiscountRules []DiscountRulesApp  `json:"discount_rules"`
         
    }
    
    // StoreInfo ...
    type StoreInfo struct {

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            StoreCode string  `json:"store_code"`
         
    }
    
    // BasePrice ...
    type BasePrice struct {

        
            Effective float64  `json:"effective"`
            CurrencyCode string  `json:"currency_code"`
            Marked float64  `json:"marked"`
            CurrencySymbol string  `json:"currency_symbol"`
         
    }
    
    // ArticlePriceInfo ...
    type ArticlePriceInfo struct {

        
            Base BasePrice  `json:"base"`
            Converted BasePrice  `json:"converted"`
         
    }
    
    // ProductArticle ...
    type ProductArticle struct {

        
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            CartItemMeta map[string]interface{}  `json:"cart_item_meta"`
            GiftCard map[string]interface{}  `json:"gift_card"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            IsGiftVisible bool  `json:"is_gift_visible"`
            ProductGroupTags []string  `json:"product_group_tags"`
            UID string  `json:"uid"`
            Store StoreInfo  `json:"store"`
            Type string  `json:"type"`
            Identifier map[string]interface{}  `json:"identifier"`
            Size string  `json:"size"`
            Price ArticlePriceInfo  `json:"price"`
            Seller BaseInfo  `json:"seller"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            MtoQuantity float64  `json:"mto_quantity"`
            SellerIdentifier string  `json:"seller_identifier"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // ProductPrice ...
    type ProductPrice struct {

        
            Marked float64  `json:"marked"`
            Selling float64  `json:"selling"`
            AddOn float64  `json:"add_on"`
            CurrencyCode string  `json:"currency_code"`
            Effective float64  `json:"effective"`
            CurrencySymbol string  `json:"currency_symbol"`
         
    }
    
    // ProductPriceInfo ...
    type ProductPriceInfo struct {

        
            Base ProductPrice  `json:"base"`
            Converted ProductPrice  `json:"converted"`
         
    }
    
    // PromiseFormatted ...
    type PromiseFormatted struct {

        
            Min string  `json:"min"`
            Max string  `json:"max"`
         
    }
    
    // PromiseTimestamp ...
    type PromiseTimestamp struct {

        
            Min float64  `json:"min"`
            Max float64  `json:"max"`
         
    }
    
    // ShipmentPromise ...
    type ShipmentPromise struct {

        
            Formatted PromiseFormatted  `json:"formatted"`
            Timestamp PromiseTimestamp  `json:"timestamp"`
         
    }
    
    // CouponDetails ...
    type CouponDetails struct {

        
            DiscountSingleQuantity float64  `json:"discount_single_quantity"`
            DiscountTotalQuantity float64  `json:"discount_total_quantity"`
            Code string  `json:"code"`
         
    }
    
    // CartProductInfo ...
    type CartProductInfo struct {

        
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            Key string  `json:"key"`
            PromoMeta PromoMeta  `json:"promo_meta"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
            Product CartProduct  `json:"product"`
            Availability ProductAvailability  `json:"availability"`
            PromotionsApplied []AppliedPromotion  `json:"promotions_applied"`
            BulkOffer map[string]interface{}  `json:"bulk_offer"`
            Article ProductArticle  `json:"article"`
            Discount string  `json:"discount"`
            Quantity float64  `json:"quantity"`
            CouponMessage string  `json:"coupon_message"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            PricePerUnit ProductPriceInfo  `json:"price_per_unit"`
            Price ProductPriceInfo  `json:"price"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            Coupon CouponDetails  `json:"coupon"`
            Message string  `json:"message"`
            Moq map[string]interface{}  `json:"moq"`
            IsSet bool  `json:"is_set"`
         
    }
    
    // CouponBreakup ...
    type CouponBreakup struct {

        
            Message string  `json:"message"`
            IsApplied bool  `json:"is_applied"`
            Description string  `json:"description"`
            UID string  `json:"uid"`
            Title string  `json:"title"`
            Type string  `json:"type"`
            Code string  `json:"code"`
            Value float64  `json:"value"`
            MinimumCartValue float64  `json:"minimum_cart_value"`
            CouponValue float64  `json:"coupon_value"`
            MaxDiscountValue float64  `json:"max_discount_value"`
            SubTitle string  `json:"sub_title"`
            CouponType string  `json:"coupon_type"`
         
    }
    
    // RawBreakup ...
    type RawBreakup struct {

        
            Coupon float64  `json:"coupon"`
            Vog float64  `json:"vog"`
            GiftCard float64  `json:"gift_card"`
            GstCharges float64  `json:"gst_charges"`
            FyndCash float64  `json:"fynd_cash"`
            ConvenienceFee float64  `json:"convenience_fee"`
            Total float64  `json:"total"`
            Subtotal float64  `json:"subtotal"`
            MrpTotal float64  `json:"mrp_total"`
            Discount float64  `json:"discount"`
            YouSaved float64  `json:"you_saved"`
            DeliveryCharge float64  `json:"delivery_charge"`
            CodCharge float64  `json:"cod_charge"`
         
    }
    
    // DisplayBreakup ...
    type DisplayBreakup struct {

        
            Message []string  `json:"message"`
            Key string  `json:"key"`
            Display string  `json:"display"`
            Value float64  `json:"value"`
            CurrencyCode string  `json:"currency_code"`
            CurrencySymbol string  `json:"currency_symbol"`
         
    }
    
    // LoyaltyPoints ...
    type LoyaltyPoints struct {

        
            Total float64  `json:"total"`
            IsApplied bool  `json:"is_applied"`
            Applicable float64  `json:"applicable"`
            Description string  `json:"description"`
         
    }
    
    // CartBreakup ...
    type CartBreakup struct {

        
            Coupon CouponBreakup  `json:"coupon"`
            Raw RawBreakup  `json:"raw"`
            Display []DisplayBreakup  `json:"display"`
            LoyaltyPoints LoyaltyPoints  `json:"loyalty_points"`
         
    }
    
    // CartCurrency ...
    type CartCurrency struct {

        
            Code string  `json:"code"`
            Symbol string  `json:"symbol"`
         
    }
    
    // PaymentSelectionLock ...
    type PaymentSelectionLock struct {

        
            DefaultOptions string  `json:"default_options"`
            Enabled bool  `json:"enabled"`
            PaymentIdentifier string  `json:"payment_identifier"`
         
    }
    
    // CartDetailResponse ...
    type CartDetailResponse struct {

        
            LastModified string  `json:"last_modified"`
            CouponText string  `json:"coupon_text"`
            Gstin string  `json:"gstin"`
            Items []CartProductInfo  `json:"items"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            BuyNow bool  `json:"buy_now"`
            AppliedPromoDetails []AppliedPromotion  `json:"applied_promo_details"`
            PanNo string  `json:"pan_no"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            Currency CartCurrency  `json:"currency"`
            Comment string  `json:"comment"`
            Message string  `json:"message"`
            ID string  `json:"id"`
            PanConfig map[string]interface{}  `json:"pan_config"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            CheckoutMode string  `json:"checkout_mode"`
            IsValid bool  `json:"is_valid"`
         
    }
    
    // AddProductCart ...
    type AddProductCart struct {

        
            ParentItemIdentifiers []map[string]string  `json:"parent_item_identifiers"`
            Pos bool  `json:"pos"`
            ArticleID string  `json:"article_id"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            ItemID float64  `json:"item_id"`
            ProductGroupTags []string  `json:"product_group_tags"`
            Display string  `json:"display"`
            Meta map[string]interface{}  `json:"meta"`
            SellerID float64  `json:"seller_id"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            StoreID float64  `json:"store_id"`
            ItemSize string  `json:"item_size"`
            ArticleAssignment map[string]interface{}  `json:"article_assignment"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // AddCartRequest ...
    type AddCartRequest struct {

        
            Items []AddProductCart  `json:"items"`
            NewCart bool  `json:"new_cart"`
         
    }
    
    // AddCartDetailResponse ...
    type AddCartDetailResponse struct {

        
            Message string  `json:"message"`
            Cart CartDetailResponse  `json:"cart"`
            Partial bool  `json:"partial"`
            Success bool  `json:"success"`
         
    }
    
    // UpdateProductCart ...
    type UpdateProductCart struct {

        
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            ArticleID string  `json:"article_id"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            ItemID float64  `json:"item_id"`
            ItemIndex float64  `json:"item_index"`
            Meta map[string]interface{}  `json:"meta"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
            ItemSize string  `json:"item_size"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // UpdateCartRequest ...
    type UpdateCartRequest struct {

        
            Items []UpdateProductCart  `json:"items"`
            Operation string  `json:"operation"`
         
    }
    
    // UpdateCartDetailResponse ...
    type UpdateCartDetailResponse struct {

        
            Message string  `json:"message"`
            Cart CartDetailResponse  `json:"cart"`
            Success bool  `json:"success"`
         
    }
    
    // DeleteCartDetailResponse ...
    type DeleteCartDetailResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // CartItemCountResponse ...
    type CartItemCountResponse struct {

        
            UserCartItemsCount float64  `json:"user_cart_items_count"`
         
    }
    
    // Coupon ...
    type Coupon struct {

        
            CouponCode string  `json:"coupon_code"`
            Message string  `json:"message"`
            IsApplied bool  `json:"is_applied"`
            Description string  `json:"description"`
            CouponValue float64  `json:"coupon_value"`
            Title string  `json:"title"`
            MinimumCartValue float64  `json:"minimum_cart_value"`
            MaxDiscountValue float64  `json:"max_discount_value"`
            CouponType string  `json:"coupon_type"`
            SubTitle string  `json:"sub_title"`
            IsApplicable bool  `json:"is_applicable"`
            ExpiresOn string  `json:"expires_on"`
         
    }
    
    // PageCoupon ...
    type PageCoupon struct {

        
            HasNext bool  `json:"has_next"`
            Total float64  `json:"total"`
            Current float64  `json:"current"`
            TotalItemCount float64  `json:"total_item_count"`
            HasPrevious bool  `json:"has_previous"`
         
    }
    
    // GetCouponResponse ...
    type GetCouponResponse struct {

        
            AvailableCouponList []Coupon  `json:"available_coupon_list"`
            Page PageCoupon  `json:"page"`
         
    }
    
    // ApplyCouponRequest ...
    type ApplyCouponRequest struct {

        
            CouponCode string  `json:"coupon_code"`
         
    }
    
    // OfferPrice ...
    type OfferPrice struct {

        
            Marked float64  `json:"marked"`
            CurrencyCode string  `json:"currency_code"`
            Effective float64  `json:"effective"`
            CurrencySymbol string  `json:"currency_symbol"`
            BulkEffective float64  `json:"bulk_effective"`
         
    }
    
    // OfferItem ...
    type OfferItem struct {

        
            AutoApplied bool  `json:"auto_applied"`
            Type string  `json:"type"`
            Total float64  `json:"total"`
            Price OfferPrice  `json:"price"`
            Best bool  `json:"best"`
            Margin float64  `json:"margin"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // OfferSeller ...
    type OfferSeller struct {

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // BulkPriceOffer ...
    type BulkPriceOffer struct {

        
            Offers []OfferItem  `json:"offers"`
            Seller OfferSeller  `json:"seller"`
         
    }
    
    // BulkPriceResponse ...
    type BulkPriceResponse struct {

        
            Data []BulkPriceOffer  `json:"data"`
         
    }
    
    // RewardPointRequest ...
    type RewardPointRequest struct {

        
            Points bool  `json:"points"`
         
    }
    
    // GeoLocation ...
    type GeoLocation struct {

        
            Latitude float64  `json:"latitude"`
            Longitude float64  `json:"longitude"`
         
    }
    
    // Address ...
    type Address struct {

        
            Tags []string  `json:"tags"`
            GoogleMapPoint map[string]interface{}  `json:"google_map_point"`
            Email string  `json:"email"`
            Meta map[string]interface{}  `json:"meta"`
            CreatedByUserID string  `json:"created_by_user_id"`
            AddressType string  `json:"address_type"`
            CountryPhoneCode string  `json:"country_phone_code"`
            Address string  `json:"address"`
            AreaCode string  `json:"area_code"`
            State string  `json:"state"`
            Phone string  `json:"phone"`
            IsActive bool  `json:"is_active"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Area string  `json:"area"`
            IsDefaultAddress bool  `json:"is_default_address"`
            City string  `json:"city"`
            Name string  `json:"name"`
            ID string  `json:"id"`
            UserID string  `json:"user_id"`
            GeoLocation GeoLocation  `json:"geo_location"`
            Country string  `json:"country"`
            CheckoutMode string  `json:"checkout_mode"`
            CountryCode string  `json:"country_code"`
            AreaCodeSlug string  `json:"area_code_slug"`
            Landmark string  `json:"landmark"`
            CountryIsoCode string  `json:"country_iso_code"`
         
    }
    
    // GetAddressesResponse ...
    type GetAddressesResponse struct {

        
            Address []Address  `json:"address"`
            PiiMasking bool  `json:"pii_masking"`
         
    }
    
    // SaveAddressResponse ...
    type SaveAddressResponse struct {

        
            IsDefaultAddress bool  `json:"is_default_address"`
            Success bool  `json:"success"`
            ID string  `json:"id"`
         
    }
    
    // UpdateAddressResponse ...
    type UpdateAddressResponse struct {

        
            IsUpdated bool  `json:"is_updated"`
            Success bool  `json:"success"`
            IsDefaultAddress bool  `json:"is_default_address"`
            ID string  `json:"id"`
         
    }
    
    // DeleteAddressResponse ...
    type DeleteAddressResponse struct {

        
            IsDeleted bool  `json:"is_deleted"`
            ID string  `json:"id"`
         
    }
    
    // SelectCartAddressRequest ...
    type SelectCartAddressRequest struct {

        
            CartID string  `json:"cart_id"`
            BillingAddressID string  `json:"billing_address_id"`
            ID string  `json:"id"`
         
    }
    
    // UpdateCartPaymentRequest ...
    type UpdateCartPaymentRequest struct {

        
            AggregatorName string  `json:"aggregator_name"`
            AddressID string  `json:"address_id"`
            ID string  `json:"id"`
            MerchantCode string  `json:"merchant_code"`
            PaymentIdentifier string  `json:"payment_identifier"`
            PaymentMode string  `json:"payment_mode"`
         
    }
    
    // CouponValidity ...
    type CouponValidity struct {

        
            Valid bool  `json:"valid"`
            Title string  `json:"title"`
            DisplayMessageEn string  `json:"display_message_en"`
            Code string  `json:"code"`
            NextValidationRequired bool  `json:"next_validation_required"`
            Discount float64  `json:"discount"`
         
    }
    
    // PaymentCouponValidate ...
    type PaymentCouponValidate struct {

        
            Message string  `json:"message"`
            CouponValidity CouponValidity  `json:"coupon_validity"`
            Success bool  `json:"success"`
         
    }
    
    // ShipmentResponse ...
    type ShipmentResponse struct {

        
            Items []CartProductInfo  `json:"items"`
            OrderType string  `json:"order_type"`
            BoxType string  `json:"box_type"`
            DpOptions map[string]interface{}  `json:"dp_options"`
            Shipments float64  `json:"shipments"`
            FulfillmentID float64  `json:"fulfillment_id"`
            FulfillmentType string  `json:"fulfillment_type"`
            ShipmentType string  `json:"shipment_type"`
            DpID string  `json:"dp_id"`
            Promise ShipmentPromise  `json:"promise"`
         
    }
    
    // CartShipmentsResponse ...
    type CartShipmentsResponse struct {

        
            LastModified string  `json:"last_modified"`
            Error bool  `json:"error"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            CouponText string  `json:"coupon_text"`
            Currency CartCurrency  `json:"currency"`
            Gstin string  `json:"gstin"`
            Comment string  `json:"comment"`
            Message string  `json:"message"`
            ID string  `json:"id"`
            UID string  `json:"uid"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            CheckoutMode string  `json:"checkout_mode"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            BuyNow bool  `json:"buy_now"`
            Shipments []ShipmentResponse  `json:"shipments"`
            IsValid bool  `json:"is_valid"`
            CartID float64  `json:"cart_id"`
            RestrictCheckout bool  `json:"restrict_checkout"`
         
    }
    
    // CustomerDetails ...
    type CustomerDetails struct {

        
            Mobile string  `json:"mobile"`
            Name string  `json:"name"`
            Email string  `json:"email"`
         
    }
    
    // CartCheckoutCustomMeta ...
    type CartCheckoutCustomMeta struct {

        
            Key string  `json:"key"`
            Value string  `json:"value"`
         
    }
    
    // StaffCheckout ...
    type StaffCheckout struct {

        
            User string  `json:"user"`
            EmployeeCode string  `json:"employee_code"`
            ID string  `json:"_id"`
            LastName string  `json:"last_name"`
            FirstName string  `json:"first_name"`
         
    }
    
    // CartCheckoutDetailRequest ...
    type CartCheckoutDetailRequest struct {

        
            CustomerDetails CustomerDetails  `json:"customer_details"`
            Meta map[string]interface{}  `json:"meta"`
            CallbackURL string  `json:"callback_url"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            Aggregator string  `json:"aggregator"`
            AddressID string  `json:"address_id"`
            OrderingStore float64  `json:"ordering_store"`
            CustomMeta []CartCheckoutCustomMeta  `json:"custom_meta"`
            PaymentAutoConfirm bool  `json:"payment_auto_confirm"`
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
            OrderType string  `json:"order_type"`
            PaymentIdentifier string  `json:"payment_identifier"`
            ID string  `json:"id"`
            Staff StaffCheckout  `json:"staff"`
            MerchantCode string  `json:"merchant_code"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
            BillingAddressID string  `json:"billing_address_id"`
            PaymentParams map[string]interface{}  `json:"payment_params"`
            PaymentMode string  `json:"payment_mode"`
         
    }
    
    // CheckCart ...
    type CheckCart struct {

        
            LastModified string  `json:"last_modified"`
            CodAvailable bool  `json:"cod_available"`
            DeliveryChargeOrderValue float64  `json:"delivery_charge_order_value"`
            CouponText string  `json:"coupon_text"`
            Gstin string  `json:"gstin"`
            Items []CartProductInfo  `json:"items"`
            DeliveryCharges float64  `json:"delivery_charges"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            BuyNow bool  `json:"buy_now"`
            OrderID string  `json:"order_id"`
            CodCharges float64  `json:"cod_charges"`
            Success bool  `json:"success"`
            UserType string  `json:"user_type"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            Currency CartCurrency  `json:"currency"`
            StoreEmps []map[string]interface{}  `json:"store_emps"`
            Comment string  `json:"comment"`
            CodMessage string  `json:"cod_message"`
            Message string  `json:"message"`
            ErrorMessage string  `json:"error_message"`
            ID string  `json:"id"`
            UID string  `json:"uid"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            StoreCode string  `json:"store_code"`
            CheckoutMode string  `json:"checkout_mode"`
            IsValid bool  `json:"is_valid"`
            CartID float64  `json:"cart_id"`
         
    }
    
    // CartCheckoutResponse ...
    type CartCheckoutResponse struct {

        
            Message string  `json:"message"`
            AppInterceptURL string  `json:"app_intercept_url"`
            OrderID string  `json:"order_id"`
            Success bool  `json:"success"`
            CallbackURL string  `json:"callback_url"`
            Cart CheckCart  `json:"cart"`
            PaymentConfirmURL string  `json:"payment_confirm_url"`
            Data map[string]interface{}  `json:"data"`
         
    }
    
    // GiftDetail ...
    type GiftDetail struct {

        
            GiftMessage string  `json:"gift_message"`
            IsGiftApplied bool  `json:"is_gift_applied"`
         
    }
    
    // ArticleGiftDetail ...
    type ArticleGiftDetail struct {

        
            ArticleID GiftDetail  `json:"article_id"`
         
    }
    
    // CartMetaRequest ...
    type CartMetaRequest struct {

        
            CheckoutMode string  `json:"checkout_mode"`
            Gstin string  `json:"gstin"`
            DeliverySlots map[string]interface{}  `json:"delivery_slots"`
            GiftDetails ArticleGiftDetail  `json:"gift_details"`
            PickUpCustomerDetails map[string]interface{}  `json:"pick_up_customer_details"`
            Comment string  `json:"comment"`
         
    }
    
    // CartMetaResponse ...
    type CartMetaResponse struct {

        
            IsValid bool  `json:"is_valid"`
            Message string  `json:"message"`
         
    }
    
    // CartMetaMissingResponse ...
    type CartMetaMissingResponse struct {

        
            Errors []string  `json:"errors"`
         
    }
    
    // GetShareCartLinkRequest ...
    type GetShareCartLinkRequest struct {

        
            Meta map[string]interface{}  `json:"meta"`
            ID string  `json:"id"`
         
    }
    
    // GetShareCartLinkResponse ...
    type GetShareCartLinkResponse struct {

        
            Token string  `json:"token"`
            ShareURL string  `json:"share_url"`
         
    }
    
    // SharedCartDetails ...
    type SharedCartDetails struct {

        
            Source map[string]interface{}  `json:"source"`
            Token string  `json:"token"`
            User map[string]interface{}  `json:"user"`
            CreatedOn string  `json:"created_on"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // SharedCart ...
    type SharedCart struct {

        
            LastModified string  `json:"last_modified"`
            CouponText string  `json:"coupon_text"`
            Gstin string  `json:"gstin"`
            Items []CartProductInfo  `json:"items"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            BuyNow bool  `json:"buy_now"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            SharedCartDetails SharedCartDetails  `json:"shared_cart_details"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            Currency CartCurrency  `json:"currency"`
            Comment string  `json:"comment"`
            Message string  `json:"message"`
            ID string  `json:"id"`
            UID string  `json:"uid"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            CheckoutMode string  `json:"checkout_mode"`
            IsValid bool  `json:"is_valid"`
            CartID float64  `json:"cart_id"`
         
    }
    
    // SharedCartResponse ...
    type SharedCartResponse struct {

        
            Cart SharedCart  `json:"cart"`
            Error string  `json:"error"`
         
    }
    
    // FreeGiftItems ...
    type FreeGiftItems struct {

        
            ItemName string  `json:"item_name"`
            ItemID float64  `json:"item_id"`
            ItemImagesURL []string  `json:"item_images_url"`
            ItemPriceDetails map[string]interface{}  `json:"item_price_details"`
            ItemBrandName string  `json:"item_brand_name"`
            ItemSlug string  `json:"item_slug"`
         
    }
    
    // PromotionOffer ...
    type PromotionOffer struct {

        
            FreeGiftItems []FreeGiftItems  `json:"free_gift_items"`
            ValidTill string  `json:"valid_till"`
            Description string  `json:"description"`
            ID string  `json:"id"`
            DiscountRules []map[string]interface{}  `json:"discount_rules"`
            BuyRules map[string]interface{}  `json:"buy_rules"`
            PromotionGroup string  `json:"promotion_group"`
            OfferText string  `json:"offer_text"`
         
    }
    
    // PromotionOffersResponse ...
    type PromotionOffersResponse struct {

        
            AvailablePromotions []PromotionOffer  `json:"available_promotions"`
         
    }
    
    // OperationErrorResponse ...
    type OperationErrorResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // CurrencyInfo ...
    type CurrencyInfo struct {

        
            Code string  `json:"code"`
            Symbol string  `json:"symbol"`
         
    }
    
    // LadderPrice ...
    type LadderPrice struct {

        
            Marked float64  `json:"marked"`
            CurrencyCode string  `json:"currency_code"`
            OfferPrice float64  `json:"offer_price"`
            Effective float64  `json:"effective"`
            CurrencySymbol string  `json:"currency_symbol"`
         
    }
    
    // LadderOfferItem ...
    type LadderOfferItem struct {

        
            Type string  `json:"type"`
            Price LadderPrice  `json:"price"`
            MinQuantity float64  `json:"min_quantity"`
            Margin float64  `json:"margin"`
            MaxQuantity float64  `json:"max_quantity"`
         
    }
    
    // LadderPriceOffer ...
    type LadderPriceOffer struct {

        
            FreeGiftItems []FreeGiftItems  `json:"free_gift_items"`
            ValidTill string  `json:"valid_till"`
            Description string  `json:"description"`
            ID string  `json:"id"`
            DiscountRules []map[string]interface{}  `json:"discount_rules"`
            BuyRules map[string]interface{}  `json:"buy_rules"`
            CalculateOn string  `json:"calculate_on"`
            PromotionGroup string  `json:"promotion_group"`
            OfferPrices []LadderOfferItem  `json:"offer_prices"`
            OfferText string  `json:"offer_text"`
         
    }
    
    // LadderPriceOffers ...
    type LadderPriceOffers struct {

        
            Currency CurrencyInfo  `json:"currency"`
            AvailableOffers []LadderPriceOffer  `json:"available_offers"`
         
    }
    
    // PaymentMeta ...
    type PaymentMeta struct {

        
            PaymentGateway string  `json:"payment_gateway"`
            PaymentIdentifier string  `json:"payment_identifier"`
            MerchantCode string  `json:"merchant_code"`
            Type string  `json:"type"`
         
    }
    
    // PaymentMethod ...
    type PaymentMethod struct {

        
            PaymentMeta PaymentMeta  `json:"payment_meta"`
            Payment string  `json:"payment"`
            Amount float64  `json:"amount"`
            Mode string  `json:"mode"`
            Name string  `json:"name"`
         
    }
    
    // CartCheckoutDetailV2Request ...
    type CartCheckoutDetailV2Request struct {

        
            CustomerDetails CustomerDetails  `json:"customer_details"`
            Meta map[string]interface{}  `json:"meta"`
            CallbackURL string  `json:"callback_url"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            Aggregator string  `json:"aggregator"`
            AddressID string  `json:"address_id"`
            OrderingStore float64  `json:"ordering_store"`
            CustomMeta map[string]interface{}  `json:"custom_meta"`
            PaymentMethods []PaymentMethod  `json:"payment_methods"`
            PaymentAutoConfirm bool  `json:"payment_auto_confirm"`
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
            OrderType string  `json:"order_type"`
            PaymentIdentifier string  `json:"payment_identifier"`
            ID string  `json:"id"`
            Staff StaffCheckout  `json:"staff"`
            MerchantCode string  `json:"merchant_code"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
            BillingAddressID string  `json:"billing_address_id"`
            CartID string  `json:"cart_id"`
            PaymentParams map[string]interface{}  `json:"payment_params"`
            PaymentMode string  `json:"payment_mode"`
         
    }
    

    
    // ApplicationResponse ...
    type ApplicationResponse struct {

        
            Application ApplicationData  `json:"application"`
         
    }
    
    // Currency ...
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
    
    // Domain ...
    type Domain struct {

        
            Verified bool  `json:"verified"`
            IsPrimary bool  `json:"is_primary"`
            IsShortlink bool  `json:"is_shortlink"`
            ID string  `json:"_id"`
            Name string  `json:"name"`
            IsPredefined bool  `json:"is_predefined"`
         
    }
    
    // ApplicationWebsite ...
    type ApplicationWebsite struct {

        
            Enabled bool  `json:"enabled"`
            Basepath string  `json:"basepath"`
         
    }
    
    // ApplicationCors ...
    type ApplicationCors struct {

        
            Domains []string  `json:"domains"`
         
    }
    
    // ApplicationAuth ...
    type ApplicationAuth struct {

        
            Enabled bool  `json:"enabled"`
         
    }
    
    // ApplicationRedirections ...
    type ApplicationRedirections struct {

        
            RedirectFrom string  `json:"redirect_from"`
            RedirectTo string  `json:"redirect_to"`
            Type string  `json:"type"`
         
    }
    
    // ApplicationMeta ...
    type ApplicationMeta struct {

        
            Name string  `json:"name"`
            Value string  `json:"value"`
         
    }
    
    // SecureUrl ...
    type SecureUrl struct {

        
            SecureURL string  `json:"secure_url"`
         
    }
    
    // ApplicationData ...
    type ApplicationData struct {

        
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
            Slug string  `json:"slug"`
         
    }
    
    // NotFound ...
    type NotFound struct {

        
            Message string  `json:"message"`
         
    }
    
    // BadRequest ...
    type BadRequest struct {

        
            Message string  `json:"message"`
         
    }
    
    // Page ...
    type Page struct {

        
            ItemTotal float64  `json:"item_total"`
            NextID string  `json:"next_id"`
            HasPrevious bool  `json:"has_previous"`
            HasNext bool  `json:"has_next"`
            Current float64  `json:"current"`
            Type string  `json:"type"`
            Size float64  `json:"size"`
         
    }
    
    // LocationDefaultLanguage ...
    type LocationDefaultLanguage struct {

        
            Name string  `json:"name"`
            Code string  `json:"code"`
         
    }
    
    // LocationDefaultCurrency ...
    type LocationDefaultCurrency struct {

        
            Name string  `json:"name"`
            Symbol string  `json:"symbol"`
            Code string  `json:"code"`
         
    }
    
    // LocationDetails ...
    type LocationDetails struct {

        
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
            StateCode string  `json:"state_code"`
            CountryCode string  `json:"country_code"`
            Latitude string  `json:"latitude"`
            Longitude string  `json:"longitude"`
         
    }
    
    // Locations ...
    type Locations struct {

        
            Items []LocationDetails  `json:"items"`
         
    }
    

    
    // TicketList ...
    type TicketList struct {

        
            Items []Ticket  `json:"items"`
            Filters Filter  `json:"filters"`
            Page Page  `json:"page"`
         
    }
    
    // Page ...
    type Page struct {

        
            ItemTotal float64  `json:"item_total"`
            NextID string  `json:"next_id"`
            HasPrevious bool  `json:"has_previous"`
            HasNext bool  `json:"has_next"`
            Current float64  `json:"current"`
            Type string  `json:"type"`
            Size float64  `json:"size"`
         
    }
    
    // TicketHistoryList ...
    type TicketHistoryList struct {

        
            Items []TicketHistory  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // CustomFormList ...
    type CustomFormList struct {

        
            Items []CustomForm  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // CreateCustomFormPayload ...
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
    
    // EditCustomFormPayload ...
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
    
    // EditTicketPayload ...
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
    
    // AgentChangePayload ...
    type AgentChangePayload struct {

        
            AgentID string  `json:"agent_id"`
         
    }
    
    // CreateVideoRoomResponse ...
    type CreateVideoRoomResponse struct {

        
            UniqueName string  `json:"unique_name"`
         
    }
    
    // CloseVideoRoomResponse ...
    type CloseVideoRoomResponse struct {

        
            Success bool  `json:"success"`
         
    }
    
    // CreateVideoRoomPayload ...
    type CreateVideoRoomPayload struct {

        
            UniqueName string  `json:"unique_name"`
            Notify []NotifyUser  `json:"notify"`
         
    }
    
    // NotifyUser ...
    type NotifyUser struct {

        
            CountryCode string  `json:"country_code"`
            PhoneNumber string  `json:"phone_number"`
         
    }
    
    // Filter ...
    type Filter struct {

        
            Priorities []Priority  `json:"priorities"`
            Categories []TicketCategory  `json:"categories"`
            Statuses []Status  `json:"statuses"`
            Assignees []map[string]interface{}  `json:"assignees"`
         
    }
    
    // TicketHistoryPayload ...
    type TicketHistoryPayload struct {

        
            Value map[string]interface{}  `json:"value"`
            Type HistoryTypeEnum  `json:"type"`
         
    }
    
    // CustomFormSubmissionPayload ...
    type CustomFormSubmissionPayload struct {

        
            Response []map[string]interface{}  `json:"response"`
            Attachments []TicketAsset  `json:"attachments"`
         
    }
    
    // GetTokenForVideoRoomResponse ...
    type GetTokenForVideoRoomResponse struct {

        
            AccessToken string  `json:"access_token"`
         
    }
    
    // GetParticipantsInsideVideoRoomResponse ...
    type GetParticipantsInsideVideoRoomResponse struct {

        
            Participants []Participant  `json:"participants"`
         
    }
    
    // Participant ...
    type Participant struct {

        
            User UserSchema  `json:"user"`
            Identity string  `json:"identity"`
            Status string  `json:"status"`
         
    }
    
    // UserSchema ...
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
    
    // PhoneNumber ...
    type PhoneNumber struct {

        
            Active bool  `json:"active"`
            Primary bool  `json:"primary"`
            Verified bool  `json:"verified"`
            Phone string  `json:"phone"`
            CountryCode float64  `json:"country_code"`
         
    }
    
    // Email ...
    type Email struct {

        
            Primary bool  `json:"primary"`
            Verified bool  `json:"verified"`
            Email string  `json:"email"`
            Active bool  `json:"active"`
         
    }
    
    // Debug ...
    type Debug struct {

        
            Source string  `json:"source"`
            Platform string  `json:"platform"`
         
    }
    
    // SubmitCustomFormResponse ...
    type SubmitCustomFormResponse struct {

        
            Message string  `json:"message"`
            Ticket Ticket  `json:"ticket"`
         
    }
    
    // TicketContext ...
    type TicketContext struct {

        
            ApplicationID string  `json:"application_id"`
            CompanyID string  `json:"company_id"`
         
    }
    
    // CreatedOn ...
    type CreatedOn struct {

        
            UserAgent string  `json:"user_agent"`
         
    }
    
    // TicketAsset ...
    type TicketAsset struct {

        
            Display string  `json:"display"`
            Value string  `json:"value"`
            Type TicketAssetTypeEnum  `json:"type"`
         
    }
    
    // TicketContent ...
    type TicketContent struct {

        
            Title string  `json:"title"`
            Description string  `json:"description"`
            Attachments []TicketAsset  `json:"attachments"`
         
    }
    
    // AddTicketPayload ...
    type AddTicketPayload struct {

        
            CreatedBy map[string]interface{}  `json:"created_by"`
            Status string  `json:"status"`
            Priority PriorityEnum  `json:"priority"`
            Category string  `json:"category"`
            Content TicketContent  `json:"content"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // Priority ...
    type Priority struct {

        
            Key PriorityEnum  `json:"key"`
            Display string  `json:"display"`
            Color string  `json:"color"`
         
    }
    
    // Status ...
    type Status struct {

        
            Key string  `json:"key"`
            Display string  `json:"display"`
            Color string  `json:"color"`
         
    }
    
    // TicketFeedbackForm ...
    type TicketFeedbackForm struct {

        
            Title string  `json:"title"`
            Display []map[string]interface{}  `json:"display"`
         
    }
    
    // TicketFeedbackList ...
    type TicketFeedbackList struct {

        
            Items []TicketFeedback  `json:"items"`
         
    }
    
    // TicketFeedbackPayload ...
    type TicketFeedbackPayload struct {

        
            FormResponse map[string]interface{}  `json:"form_response"`
         
    }
    
    // SubmitButton ...
    type SubmitButton struct {

        
            Title string  `json:"title"`
            TitleColor string  `json:"title_color"`
            BackgroundColor string  `json:"background_color"`
         
    }
    
    // PollForAssignment ...
    type PollForAssignment struct {

        
            Duration float64  `json:"duration"`
            Message string  `json:"message"`
            SuccessMessage string  `json:"success_message"`
            FailureMessage string  `json:"failure_message"`
         
    }
    
    // CustomForm ...
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
    
    // CommunicationDetails ...
    type CommunicationDetails struct {

        
            Type string  `json:"type"`
            Title string  `json:"title"`
            Value string  `json:"value"`
            Description string  `json:"description"`
            Enabled bool  `json:"enabled"`
         
    }
    
    // SupportGeneralConfig ...
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
    
    // FeedbackForm ...
    type FeedbackForm struct {

        
            Inputs map[string]interface{}  `json:"inputs"`
            Title string  `json:"title"`
            Timestamps map[string]interface{}  `json:"timestamps"`
         
    }
    
    // TicketSubCategory ...
    type TicketSubCategory struct {

        
            Key string  `json:"key"`
            Display string  `json:"display"`
            SubCategories *TicketSubCategory  `json:"sub_categories"`
         
    }
    
    // TicketCategory ...
    type TicketCategory struct {

        
            Display string  `json:"display"`
            Key string  `json:"key"`
            SubCategories *TicketCategory  `json:"sub_categories"`
            GroupID float64  `json:"group_id"`
            FeedbackForm FeedbackForm  `json:"feedback_form"`
         
    }
    
    // CategoryData ...
    type CategoryData struct {

        
            List TicketCategory  `json:"list"`
         
    }
    
    // IntegrationConfig ...
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
    
    // FeedbackResponseItem ...
    type FeedbackResponseItem struct {

        
            Display string  `json:"display"`
            Key string  `json:"key"`
            Value string  `json:"value"`
         
    }
    
    // TicketFeedback ...
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
    
    // TicketHistory ...
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
    
    // Ticket ...
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
    

    
    // ThemeReq ...
    type ThemeReq struct {

        
            MarketplaceThemeID string  `json:"marketplace_theme_id"`
         
    }
    
    // ThemeSchema ...
    type ThemeSchema struct {

        
            ID string  `json:"_id"`
            Name string  `json:"name"`
            MarketplaceThemeID MarketplaceThemeId  `json:"marketplace_theme_id"`
            CompanyID float64  `json:"company_id"`
            Meta ThemeMeta  `json:"meta"`
            CreatedAt string  `json:"created_at"`
            UpdatedAt string  `json:"updated_at"`
         
    }
    
    // MarketplaceThemeId ...
    type MarketplaceThemeId struct {

        
            ID string  `json:"_id"`
            IsDefault bool  `json:"is_default"`
         
    }
    
    // ThemeMeta ...
    type ThemeMeta struct {

        
            Payment ThemePayment  `json:"payment"`
            Industry []string  `json:"industry"`
            Description string  `json:"description"`
            Images ThemeImages  `json:"images"`
            Slug string  `json:"slug"`
         
    }
    
    // ThemePayment ...
    type ThemePayment struct {

        
            IsPaid bool  `json:"is_paid"`
            Amount float64  `json:"amount"`
         
    }
    
    // ThemeImages ...
    type ThemeImages struct {

        
            Desktop string  `json:"desktop"`
            Mobile string  `json:"mobile"`
         
    }
    
    // AvailablePageSchema ...
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
    
    // AvailablePageSectionMetaAttributes ...
    type AvailablePageSectionMetaAttributes struct {

        
            Attributes map[string]interface{}  `json:"attributes"`
         
    }
    
    // AvailablePageSeo ...
    type AvailablePageSeo struct {

        
            Title string  `json:"title"`
            Description string  `json:"description"`
            ID string  `json:"_id"`
         
    }
    
    // AvailablePageSchemaSections ...
    type AvailablePageSchemaSections struct {

        
            Name string  `json:"name"`
            Label string  `json:"label"`
            Props map[string]interface{}  `json:"props"`
            Blocks []map[string]interface{}  `json:"blocks"`
            Preset map[string]interface{}  `json:"preset"`
            Predicate AvailablePagePredicate  `json:"predicate"`
         
    }
    
    // AvailablePageScreenPredicate ...
    type AvailablePageScreenPredicate struct {

        
            Mobile bool  `json:"mobile"`
            Desktop bool  `json:"desktop"`
            Tablet bool  `json:"tablet"`
         
    }
    
    // AvailablePageUserPredicate ...
    type AvailablePageUserPredicate struct {

        
            Authenticated bool  `json:"authenticated"`
            Anonymous bool  `json:"anonymous"`
         
    }
    
    // AvailablePageRoutePredicate ...
    type AvailablePageRoutePredicate struct {

        
            Selected string  `json:"selected"`
            ExactURL string  `json:"exact_url"`
            Query map[string]interface{}  `json:"query"`
         
    }
    
    // AvailablePagePredicate ...
    type AvailablePagePredicate struct {

        
            Screen AvailablePageScreenPredicate  `json:"screen"`
            User AvailablePageUserPredicate  `json:"user"`
            Route AvailablePageRoutePredicate  `json:"route"`
         
    }
    
    // MarketplaceThemeResponse ...
    type MarketplaceThemeResponse struct {

        
            Status float64  `json:"status"`
            Body MarketplaceThemeResponseBody  `json:"body"`
         
    }
    
    // MarketplaceThemeResponseBody ...
    type MarketplaceThemeResponseBody struct {

        
            Themes []MarketplaceTheme  `json:"themes"`
            Page PaginationSchema  `json:"page"`
         
    }
    
    // ArrayOfMarketplaceTheme ...
    type ArrayOfMarketplaceTheme struct {

        
            Body []MarketplaceTheme  `json:"body"`
         
    }
    
    // ThemeCreateRequest ...
    type ThemeCreateRequest struct {

        
            Src string  `json:"src"`
            Release Release  `json:"release"`
         
    }
    
    // MarketplaceTheme ...
    type MarketplaceTheme struct {

        
            ID string  `json:"_id"`
            Payment PaymentInfo  `json:"payment"`
            Contact ContactInfo  `json:"contact"`
            Industry []string  `json:"industry"`
            IsUpdate bool  `json:"is_update"`
            IsDefault bool  `json:"is_default"`
            Name string  `json:"name"`
            Tagline string  `json:"tagline"`
            Description string  `json:"description"`
            CatalogSize CatalogSize  `json:"catalog_size"`
            Images MarketplaceThemeImages  `json:"images"`
            Carousel []CarouselItem  `json:"carousel"`
            Src string  `json:"src"`
            Explore ExploreInfo  `json:"explore"`
            Features []Feature  `json:"features"`
            Highlights []Highlight  `json:"highlights"`
            Variations []Variation  `json:"variations"`
            Documentation Documentation  `json:"documentation"`
            Status string  `json:"status"`
            Step float64  `json:"step"`
            Comments Comments  `json:"comments"`
            Release Release  `json:"release"`
            Slug string  `json:"slug"`
            OrganizationID string  `json:"organization_id"`
            UserID string  `json:"user_id"`
            CreatedAt string  `json:"created_at"`
            UpdatedAt string  `json:"updated_at"`
            TemplateThemeID string  `json:"template_theme_id"`
         
    }
    
    // PaymentInfo ...
    type PaymentInfo struct {

        
            IsPaid bool  `json:"is_paid"`
            Amount float64  `json:"amount"`
         
    }
    
    // ContactInfo ...
    type ContactInfo struct {

        
            DeveloperContact []string  `json:"developer_contact"`
            SellerContact string  `json:"seller_contact"`
         
    }
    
    // CatalogSize ...
    type CatalogSize struct {

        
            Min float64  `json:"min"`
            Max float64  `json:"max"`
         
    }
    
    // MarketplaceThemeImages ...
    type MarketplaceThemeImages struct {

        
            Desktop string  `json:"desktop"`
            Mobile string  `json:"mobile"`
         
    }
    
    // CarouselItem ...
    type CarouselItem struct {

        
            Desktop string  `json:"desktop"`
            Mobile string  `json:"mobile"`
         
    }
    
    // ExploreInfo ...
    type ExploreInfo struct {

        
            Title string  `json:"title"`
            Description string  `json:"description"`
         
    }
    
    // Feature ...
    type Feature struct {

        
            Category string  `json:"category"`
            List []FeatureItem  `json:"list"`
         
    }
    
    // FeatureItem ...
    type FeatureItem struct {

        
            Label string  `json:"label"`
            Description string  `json:"description"`
         
    }
    
    // Highlight ...
    type Highlight struct {

        
            Title string  `json:"title"`
            Description string  `json:"description"`
            Image string  `json:"image"`
         
    }
    
    // Variation ...
    type Variation struct {

        
            Name string  `json:"name"`
            Color string  `json:"color"`
            DemoURL string  `json:"demo_url"`
            Images MarketplaceThemeImages  `json:"images"`
         
    }
    
    // Documentation ...
    type Documentation struct {

        
            Notes string  `json:"notes"`
            URL string  `json:"url"`
         
    }
    
    // Comments ...
    type Comments struct {

        
            DeveloperRemark string  `json:"developer_remark"`
            ReviewerFeedback string  `json:"reviewer_feedback"`
         
    }
    
    // Release ...
    type Release struct {

        
            Version string  `json:"version"`
            Notes string  `json:"notes"`
         
    }
    
    // ThemeSlugResponse ...
    type ThemeSlugResponse struct {

        
            Theme MarketplaceTheme  `json:"theme"`
            Organization Organization  `json:"organization"`
            User []ThemeCreator  `json:"user"`
         
    }
    
    // Organization ...
    type Organization struct {

        
            Meta OrganizationMeta  `json:"meta"`
            ID string  `json:"_id"`
         
    }
    
    // OrganizationMeta ...
    type OrganizationMeta struct {

        
            EcommPlatformUsed []string  `json:"ecomm_platform_used"`
            Goals []string  `json:"goals"`
         
    }
    
    // ThemeCreator ...
    type ThemeCreator struct {

        
            ID string  `json:"_id"`
            Gender string  `json:"gender"`
            AccountType string  `json:"account_type"`
            Active bool  `json:"active"`
            FirstName string  `json:"first_name"`
            LastName string  `json:"last_name"`
            PhoneNumbers []PhoneNumber  `json:"phone_numbers"`
            Emails []Email  `json:"emails"`
         
    }
    
    // PhoneNumber ...
    type PhoneNumber struct {

        
            Active bool  `json:"active"`
            Primary bool  `json:"primary"`
            Verified bool  `json:"verified"`
            Phone string  `json:"phone"`
            CountryCode float64  `json:"country_code"`
         
    }
    
    // Email ...
    type Email struct {

        
            Active bool  `json:"active"`
            Primary bool  `json:"primary"`
            Verified bool  `json:"verified"`
            Email string  `json:"email"`
         
    }
    
    // ThemeAndUserDetailsResponse ...
    type ThemeAndUserDetailsResponse struct {

        
            Themes []MarketplaceTheme  `json:"themes"`
            User []ThemeCreator  `json:"user"`
         
    }
    
    // ThemeRejectionReasons ...
    type ThemeRejectionReasons struct {

        
            ID string  `json:"_id"`
            Message string  `json:"message"`
            ThemeID string  `json:"theme_id"`
            OrganizationID string  `json:"organization_id"`
            AdminID string  `json:"admin_id"`
            UserID string  `json:"user_id"`
            Status string  `json:"status"`
            RejectionReasons map[string]map[string]interface{}  `json:"rejection_reasons"`
            CreatedAt string  `json:"created_at"`
            UpdatedAt string  `json:"updated_at"`
         
    }
    
    // RejectionReason ...
    type RejectionReason struct {

        
            Message string  `json:"message"`
         
    }
    
    // ThemeReviewRequest ...
    type ThemeReviewRequest struct {

        
            DynamicProperties map[string]map[string]interface{}  `json:"dynamic_properties"`
         
    }
    
    // UpdateReviewStatusRequest ...
    type UpdateReviewStatusRequest struct {

        
            Status string  `json:"status"`
         
    }
    
    // AllAvailablePageSchema ...
    type AllAvailablePageSchema struct {

        
            Pages []AvailablePageSchema  `json:"pages"`
         
    }
    
    // PaginationSchema ...
    type PaginationSchema struct {

        
            Size float64  `json:"size"`
            ItemTotal float64  `json:"item_total"`
            HasNext bool  `json:"has_next"`
            Type string  `json:"type"`
            Current float64  `json:"current"`
         
    }
    
    // ThemesListingResponseSchema ...
    type ThemesListingResponseSchema struct {

        
            Items []ThemesSchema  `json:"items"`
            Page PaginationSchema  `json:"page"`
         
    }
    
    // AddThemeRequestSchema ...
    type AddThemeRequestSchema struct {

        
            ThemeID string  `json:"theme_id"`
         
    }
    
    // UpgradableThemeSchema ...
    type UpgradableThemeSchema struct {

        
            ParentTheme string  `json:"parent_theme"`
            AppliedTheme string  `json:"applied_theme"`
            Upgrade bool  `json:"upgrade"`
         
    }
    
    // FontsSchema ...
    type FontsSchema struct {

        
            Items FontsSchemaItems  `json:"items"`
            Kind string  `json:"kind"`
         
    }
    
    // BlitzkriegApiErrorSchema ...
    type BlitzkriegApiErrorSchema struct {

        
            Message string  `json:"message"`
         
    }
    
    // BlitzkriegNotFoundSchema ...
    type BlitzkriegNotFoundSchema struct {

        
            Message string  `json:"message"`
         
    }
    
    // BlitzkriegInternalServerErrorSchema ...
    type BlitzkriegInternalServerErrorSchema struct {

        
            Message string  `json:"message"`
         
    }
    
    // FontsSchemaItems ...
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
    
    // FontsSchemaItemsFiles ...
    type FontsSchemaItemsFiles struct {

        
            Regular string  `json:"regular"`
            Italic string  `json:"italic"`
            Bold string  `json:"bold"`
         
    }
    
    // ThemesSchema ...
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
    
    // availableSectionSchema ...
    type availableSectionSchema struct {

        
            Blocks []Blocks  `json:"blocks"`
            Name string  `json:"name"`
            Label string  `json:"label"`
            Props []BlocksProps  `json:"props"`
         
    }
    
    // Information ...
    type Information struct {

        
            Images Images  `json:"images"`
            Features []string  `json:"features"`
            Name string  `json:"name"`
            Description string  `json:"description"`
         
    }
    
    // Images ...
    type Images struct {

        
            Desktop []string  `json:"desktop"`
            Android []string  `json:"android"`
            Ios []string  `json:"ios"`
            Thumbnail []string  `json:"thumbnail"`
         
    }
    
    // Src ...
    type Src struct {

        
            Link string  `json:"link"`
         
    }
    
    // AssetsSchema ...
    type AssetsSchema struct {

        
            UmdJs UmdJs  `json:"umd_js"`
            CommonJs CommonJs  `json:"common_js"`
            Css Css  `json:"css"`
         
    }
    
    // UmdJs ...
    type UmdJs struct {

        
            Link string  `json:"link"`
            Links []string  `json:"links"`
         
    }
    
    // CommonJs ...
    type CommonJs struct {

        
            Link string  `json:"link"`
         
    }
    
    // Css ...
    type Css struct {

        
            Link string  `json:"link"`
            Links []string  `json:"links"`
         
    }
    
    // Sections ...
    type Sections struct {

        
            Attributes string  `json:"attributes"`
         
    }
    
    // Config ...
    type Config struct {

        
            Preset Preset  `json:"preset"`
            GlobalSchema GlobalSchema  `json:"global_schema"`
            Current string  `json:"current"`
            List []ListSchemaItem  `json:"list"`
         
    }
    
    // Preset ...
    type Preset struct {

        
            Pages []AvailablePageSchema  `json:"pages"`
         
    }
    
    // GlobalSchema ...
    type GlobalSchema struct {

        
            Props []GlobalSchemaProps  `json:"props"`
         
    }
    
    // ListSchemaItem ...
    type ListSchemaItem struct {

        
            GlobalConfig map[string]interface{}  `json:"global_config"`
            Page []ConfigPage  `json:"page"`
            Name string  `json:"name"`
         
    }
    
    // Colors ...
    type Colors struct {

        
            BgColor string  `json:"bg_color"`
            PrimaryColor string  `json:"primary_color"`
            SecondaryColor string  `json:"secondary_color"`
            AccentColor string  `json:"accent_color"`
            LinkColor string  `json:"link_color"`
            ButtonSecondaryColor string  `json:"button_secondary_color"`
         
    }
    
    // Custom ...
    type Custom struct {

        
            Props map[string]interface{}  `json:"props"`
         
    }
    
    // ConfigPage ...
    type ConfigPage struct {

        
            Settings map[string]interface{}  `json:"settings"`
            Page string  `json:"page"`
         
    }
    
    // Font ...
    type Font struct {

        
            Family string  `json:"family"`
            Variants Variants  `json:"variants"`
         
    }
    
    // Variants ...
    type Variants struct {

        
            Medium Medium  `json:"medium"`
            SemiBold SemiBold  `json:"semi_bold"`
            Bold Bold  `json:"bold"`
            Light Light  `json:"light"`
            Regular Regular  `json:"regular"`
         
    }
    
    // Medium ...
    type Medium struct {

        
            Name string  `json:"name"`
            File string  `json:"file"`
         
    }
    
    // SemiBold ...
    type SemiBold struct {

        
            Name string  `json:"name"`
            File string  `json:"file"`
         
    }
    
    // Bold ...
    type Bold struct {

        
            Name string  `json:"name"`
            File string  `json:"file"`
         
    }
    
    // Light ...
    type Light struct {

        
            Name string  `json:"name"`
            File string  `json:"file"`
         
    }
    
    // Regular ...
    type Regular struct {

        
            Name string  `json:"name"`
            File string  `json:"file"`
         
    }
    
    // Blocks ...
    type Blocks struct {

        
            Type string  `json:"type"`
            Name string  `json:"name"`
            Props []BlocksProps  `json:"props"`
         
    }
    
    // GlobalSchemaProps ...
    type GlobalSchemaProps struct {

        
            ID string  `json:"id"`
            Label string  `json:"label"`
            Type string  `json:"type"`
            Category string  `json:"category"`
         
    }
    
    // BlocksProps ...
    type BlocksProps struct {

        
            ID string  `json:"id"`
            Label string  `json:"label"`
            Type string  `json:"type"`
         
    }
    
    // ApplyThemeRequestV2 ...
    type ApplyThemeRequestV2 struct {

        
            MarketplaceThemeID string  `json:"marketplace_theme_id"`
         
    }
    
    // ApplyThemeResponseV2 ...
    type ApplyThemeResponseV2 struct {

        
            Font FontV2  `json:"font"`
            Config ConfigV2  `json:"config"`
            Applied bool  `json:"applied"`
            IsPrivate bool  `json:"is_private"`
            Tags []string  `json:"tags"`
            ID string  `json:"_id"`
            ApplicationID string  `json:"application_id"`
            MarketplaceThemeID string  `json:"marketplace_theme_id"`
            Meta MetaV2  `json:"meta"`
            Name string  `json:"name"`
            TemplateThemeID string  `json:"template_theme_id"`
            Version string  `json:"version"`
            Styles map[string]interface{}  `json:"styles"`
            CreatedAt string  `json:"created_at"`
            UpdatedAt string  `json:"updated_at"`
         
    }
    
    // AllThemesApplicationResponseV2 ...
    type AllThemesApplicationResponseV2 struct {

        
            Font FontV2  `json:"font"`
            Config ConfigV2  `json:"config"`
            Applied bool  `json:"applied"`
            IsPrivate bool  `json:"is_private"`
            Tags []string  `json:"tags"`
            ID string  `json:"_id"`
            ApplicationID string  `json:"application_id"`
            MarketplaceThemeID string  `json:"marketplace_theme_id"`
            Meta MetaV2  `json:"meta"`
            Name string  `json:"name"`
            TemplateThemeID string  `json:"template_theme_id"`
            Version string  `json:"version"`
            Styles map[string]interface{}  `json:"styles"`
            CreatedAt string  `json:"created_at"`
            UpdatedAt string  `json:"updated_at"`
            Assets AssetsV2  `json:"assets"`
            AvailableSections []SectionItem  `json:"available_sections"`
         
    }
    
    // ThemeUpgradableResponseV2 ...
    type ThemeUpgradableResponseV2 struct {

        
            Upgrade bool  `json:"upgrade"`
            Versions ThemeVersionsV2  `json:"versions"`
            Message string  `json:"message"`
         
    }
    
    // UpdateThemeNameRequestBodyV2 ...
    type UpdateThemeNameRequestBodyV2 struct {

        
            Name string  `json:"name"`
         
    }
    
    // UpdateThemeRequestBodyV2 ...
    type UpdateThemeRequestBodyV2 struct {

        
            Config ConfigV2  `json:"config"`
            Font FontV2  `json:"font"`
         
    }
    
    // FontV2 ...
    type FontV2 struct {

        
            Variants FontVariantsV2  `json:"variants"`
            Family string  `json:"family"`
         
    }
    
    // FontVariantsV2 ...
    type FontVariantsV2 struct {

        
            Light FontVariantV2  `json:"light"`
            Regular FontVariantV2  `json:"regular"`
            Medium FontVariantV2  `json:"medium"`
            SemiBold FontVariantV2  `json:"semi_bold"`
            Bold FontVariantV2  `json:"bold"`
         
    }
    
    // FontVariantV2 ...
    type FontVariantV2 struct {

        
            Name string  `json:"name"`
            File string  `json:"file"`
         
    }
    
    // ConfigV2 ...
    type ConfigV2 struct {

        
            Current string  `json:"current"`
            List []ConfigurationV2  `json:"list"`
            GlobalSchema GlobalSchemaV2  `json:"global_schema"`
            Preset PresetV2  `json:"preset"`
         
    }
    
    // ConfigurationV2 ...
    type ConfigurationV2 struct {

        
            Name string  `json:"name"`
            GlobalConfig GlobalConfigV2  `json:"global_config"`
            Custom CustomConfigV2  `json:"custom"`
            Page []string  `json:"page"`
         
    }
    
    // GlobalConfigV2 ...
    type GlobalConfigV2 struct {

        
            Statics StaticConfigV2  `json:"statics"`
            Auth AuthConfigV2  `json:"auth"`
            Palette PaletteConfigV2  `json:"palette"`
         
    }
    
    // StaticConfigV2 ...
    type StaticConfigV2 struct {

        
            Props StaticPropsV2  `json:"props"`
         
    }
    
    // AuthConfigV2 ...
    type AuthConfigV2 struct {

        
            ShowHeaderAuth bool  `json:"show_header_auth"`
            ShowFooterAuth bool  `json:"show_footer_auth"`
         
    }
    
    // PaletteConfigV2 ...
    type PaletteConfigV2 struct {

        
            GeneralSetting GeneralSettingV2  `json:"general_setting"`
            AdvanceSetting AdvanceSettingV2  `json:"advance_setting"`
         
    }
    
    // CustomConfigV2 ...
    type CustomConfigV2 struct {

        
            Props CustomPropsV2  `json:"props"`
         
    }
    
    // MetaV2 ...
    type MetaV2 struct {

        
            Payment PaymentV2  `json:"payment"`
            Description string  `json:"description"`
            Industry []string  `json:"industry"`
            Release ReleaseV2  `json:"release"`
            Images ImagesV2  `json:"images"`
            Slug string  `json:"slug"`
            Name string  `json:"name"`
         
    }
    
    // PaymentV2 ...
    type PaymentV2 struct {

        
            IsPaid bool  `json:"is_paid"`
            Amount float64  `json:"amount"`
         
    }
    
    // ReleaseV2 ...
    type ReleaseV2 struct {

        
            Notes string  `json:"notes"`
            Version string  `json:"version"`
         
    }
    
    // ImagesV2 ...
    type ImagesV2 struct {

        
            Desktop string  `json:"desktop"`
            Mobile string  `json:"mobile"`
         
    }
    
    // StaticPropsV2 ...
    type StaticPropsV2 struct {

        
            Colors ColorsV2  `json:"colors"`
            Auth AuthConfigV2  `json:"auth"`
         
    }
    
    // ColorsV2 ...
    type ColorsV2 struct {

        
            PrimaryColor string  `json:"primary_color"`
            SecondaryColor string  `json:"secondary_color"`
            AccentColor string  `json:"accent_color"`
            LinkColor string  `json:"link_color"`
            ButtonSecondaryColor string  `json:"button_secondary_color"`
            BgColor string  `json:"bg_color"`
         
    }
    
    // GeneralSettingV2 ...
    type GeneralSettingV2 struct {

        
            Theme ThemeSettingV2  `json:"theme"`
            Text TextSettingV2  `json:"text"`
            Button ButtonSettingV2  `json:"button"`
            SaleDiscount SaleDiscountSettingV2  `json:"sale_discount"`
            Header HeaderSettingV2  `json:"header"`
            Footer FooterSettingV2  `json:"footer"`
         
    }
    
    // AdvanceSettingV2 ...
    type AdvanceSettingV2 struct {

        
            OverlayPopup OverlayPopupSettingV2  `json:"overlay_popup"`
            DividerStrokeHighlight DividerStrokeHighlightSettingV2  `json:"divider_stroke_highlight"`
            UserAlerts UserAlertsSettingV2  `json:"user_alerts"`
         
    }
    
    // ThemeSettingV2 ...
    type ThemeSettingV2 struct {

        
            PageBackground string  `json:"page_background"`
            ThemeAccent string  `json:"theme_accent"`
         
    }
    
    // TextSettingV2 ...
    type TextSettingV2 struct {

        
            TextHeading string  `json:"text_heading"`
            TextBody string  `json:"text_body"`
            TextLabel string  `json:"text_label"`
            TextSecondary string  `json:"text_secondary"`
         
    }
    
    // ButtonSettingV2 ...
    type ButtonSettingV2 struct {

        
            ButtonPrimary string  `json:"button_primary"`
            ButtonSecondary string  `json:"button_secondary"`
            ButtonLink string  `json:"button_link"`
         
    }
    
    // SaleDiscountSettingV2 ...
    type SaleDiscountSettingV2 struct {

        
            SaleBadgeBackground string  `json:"sale_badge_background"`
            SaleBadgeText string  `json:"sale_badge_text"`
            SaleDiscountText string  `json:"sale_discount_text"`
            SaleTimer string  `json:"sale_timer"`
         
    }
    
    // HeaderSettingV2 ...
    type HeaderSettingV2 struct {

        
            HeaderBackground string  `json:"header_background"`
            HeaderNav string  `json:"header_nav"`
            HeaderIcon string  `json:"header_icon"`
         
    }
    
    // FooterSettingV2 ...
    type FooterSettingV2 struct {

        
            FooterBackground string  `json:"footer_background"`
            FooterBottomBackground string  `json:"footer_bottom_background"`
            FooterHeadingText string  `json:"footer_heading_text"`
            FooterBodyText string  `json:"footer_body_text"`
            FooterIcon string  `json:"footer_icon"`
         
    }
    
    // OverlayPopupSettingV2 ...
    type OverlayPopupSettingV2 struct {

        
            DialogBackgroung string  `json:"dialog_backgroung"`
            Overlay string  `json:"overlay"`
         
    }
    
    // DividerStrokeHighlightSettingV2 ...
    type DividerStrokeHighlightSettingV2 struct {

        
            DividerStrokes string  `json:"divider_strokes"`
            Highlight string  `json:"highlight"`
         
    }
    
    // UserAlertsSettingV2 ...
    type UserAlertsSettingV2 struct {

        
            SuccessBackground string  `json:"success_background"`
            SuccessText string  `json:"success_text"`
            ErrorBackground string  `json:"error_background"`
            ErrorText string  `json:"error_text"`
            InfoBackground string  `json:"info_background"`
            InfoText string  `json:"info_text"`
         
    }
    
    // CustomPropsV2 ...
    type CustomPropsV2 struct {

        
            HeaderBgColor string  `json:"header_bg_color"`
            HeaderTextColor string  `json:"header_text_color"`
            HeaderBorderColor string  `json:"header_border_color"`
            HeaderIconColor string  `json:"header_icon_color"`
            HeaderCartNotificationBgColor string  `json:"header_cart_notification_bg_color"`
            HeaderCartNotificationTextColor string  `json:"header_cart_notification_text_color"`
            HeaderNavHoverColor string  `json:"header_nav_hover_color"`
            ButtonPrimaryColor string  `json:"button_primary_color"`
            ButtonPrimaryLabelColor string  `json:"button_primary_label_color"`
            ButtonAddToCartColor string  `json:"button_add_to_cart_color"`
            ButtonAddToCartLabelColor string  `json:"button_add_to_cart_label_color"`
            ButtonSecondaryColor string  `json:"button_secondary_color"`
            ButtonSecondaryLabelColor string  `json:"button_secondary_label_color"`
            ButtonTertiaryColor string  `json:"button_tertiary_color"`
            ButtonTertiaryLabelColor string  `json:"button_tertiary_label_color"`
            ButtonTertiaryHoverColor string  `json:"button_tertiary_hover_color"`
            ButtonTertiaryHoverTextColor string  `json:"button_tertiary_hover_text_color"`
            TextHeadingLinkColor string  `json:"text_heading_link_color"`
            TextBodyColor string  `json:"text_body_color"`
            TextPriceColor string  `json:"text_price_color"`
            TextSalePriceColor string  `json:"text_sale_price_color"`
            TextStrikethroughPriceColor string  `json:"text_strikethrough_price_color"`
            TextDiscountColor string  `json:"text_discount_color"`
            FooterBgColor string  `json:"footer_bg_color"`
            FooterTextColor string  `json:"footer_text_color"`
            FooterBorderColor string  `json:"footer_border_color"`
            FooterNavHoverColor string  `json:"footer_nav_hover_color"`
            DisableCart bool  `json:"disable_cart"`
            IsMenuBelowLogo bool  `json:"is_menu_below_logo"`
            MenuPosition string  `json:"menu_position"`
         
    }
    
    // GlobalSchemaV2 ...
    type GlobalSchemaV2 struct {

        
            Props []PropV2  `json:"props"`
         
    }
    
    // PropV2 ...
    type PropV2 struct {

        
            Type string  `json:"type"`
            Category string  `json:"category"`
            ID string  `json:"id"`
            Label string  `json:"label"`
            Info string  `json:"info"`
         
    }
    
    // AssetsV2 ...
    type AssetsV2 struct {

        
            UmdJs UMDJs  `json:"umd_js"`
            CommonJs CommonJS  `json:"common_js"`
            Css CSS  `json:"css"`
         
    }
    
    // UMDJs ...
    type UMDJs struct {

        
            Links []string  `json:"links"`
         
    }
    
    // CommonJS ...
    type CommonJS struct {

        
            Link string  `json:"link"`
         
    }
    
    // CSS ...
    type CSS struct {

        
            Links []string  `json:"links"`
         
    }
    
    // SectionItem ...
    type SectionItem struct {

        
            Props []interface{}  `json:"props"`
            Blocks []interface{}  `json:"blocks"`
            Name string  `json:"name"`
            Label string  `json:"label"`
         
    }
    
    // PresetV2 ...
    type PresetV2 struct {

        
            Pages []PageV2  `json:"pages"`
         
    }
    
    // PageV2 ...
    type PageV2 struct {

        
            Sections []SectionV2  `json:"sections"`
            Value string  `json:"value"`
         
    }
    
    // SectionV2 ...
    type SectionV2 struct {

        
            Blocks []BlockV2  `json:"blocks"`
            Predicate PredicateV2  `json:"predicate"`
            Name string  `json:"name"`
            Props SectionPropsV2  `json:"props"`
            Preset SectionPresetV2  `json:"preset"`
         
    }
    
    // BlockV2 ...
    type BlockV2 struct {

        
            Type string  `json:"type"`
            Name string  `json:"name"`
            Props BlockPropsV2  `json:"props"`
         
    }
    
    // PredicateV2 ...
    type PredicateV2 struct {

        
            Screen ScreenV2  `json:"screen"`
            User UserV2  `json:"user"`
            Route RouteV2  `json:"route"`
         
    }
    
    // ScreenV2 ...
    type ScreenV2 struct {

        
            Mobile bool  `json:"mobile"`
            Desktop bool  `json:"desktop"`
            Tablet bool  `json:"tablet"`
         
    }
    
    // UserV2 ...
    type UserV2 struct {

        
            Authenticated bool  `json:"authenticated"`
            Anonymous bool  `json:"anonymous"`
         
    }
    
    // RouteV2 ...
    type RouteV2 struct {

        
            Selected string  `json:"selected"`
            ExactURL string  `json:"exact_url"`
         
    }
    
    // SectionPropsV2 ...
    type SectionPropsV2 struct {

        
            Title TextPropV2  `json:"title"`
            ItemMargin TextPropV2  `json:"item_margin"`
            Autoplay CheckboxPropV2  `json:"autoplay"`
            SlideInterval RangePropV2  `json:"slide_interval"`
         
    }
    
    // SectionPresetV2 ...
    type SectionPresetV2 struct {

        
            Blocks []BlockV2  `json:"blocks"`
         
    }
    
    // BlockPropsV2 ...
    type BlockPropsV2 struct {

        
            Image ImagePickerPropV2  `json:"image"`
            SlideLink UrlPropV2  `json:"slide_link"`
         
    }
    
    // TextPropV2 ...
    type TextPropV2 struct {

        
            Value string  `json:"value"`
            Type string  `json:"type"`
         
    }
    
    // CheckboxPropV2 ...
    type CheckboxPropV2 struct {

        
            Value bool  `json:"value"`
            Type string  `json:"type"`
         
    }
    
    // RangePropV2 ...
    type RangePropV2 struct {

        
            Value float64  `json:"value"`
            Type string  `json:"type"`
         
    }
    
    // ImagePickerPropV2 ...
    type ImagePickerPropV2 struct {

        
            Type string  `json:"type"`
            Value string  `json:"value"`
         
    }
    
    // UrlPropV2 ...
    type UrlPropV2 struct {

        
            Type string  `json:"type"`
            Value string  `json:"value"`
         
    }
    
    // ThemeVersionsV2 ...
    type ThemeVersionsV2 struct {

        
            ParentTheme string  `json:"parent_theme"`
            AppliedTheme string  `json:"applied_theme"`
         
    }
    

    
    // BlockUserRequestSchema ...
    type BlockUserRequestSchema struct {

        
            Status bool  `json:"status"`
            UserID []string  `json:"user_id"`
            Reason string  `json:"reason"`
         
    }
    
    // ArchiveUserRequestSchema ...
    type ArchiveUserRequestSchema struct {

        
            UserID string  `json:"user_id"`
         
    }
    
    // DeleteApplicationUserRequestSchema ...
    type DeleteApplicationUserRequestSchema struct {

        
            UserID string  `json:"user_id"`
            Reason string  `json:"reason"`
            ReasonID string  `json:"reason_id"`
            RequestID string  `json:"request_id"`
            Otp string  `json:"otp"`
         
    }
    
    // UnDeleteUserRequestSchema ...
    type UnDeleteUserRequestSchema struct {

        
            UserID string  `json:"user_id"`
            Reason string  `json:"reason"`
            ReasonID string  `json:"reason_id"`
         
    }
    
    // EditEmailRequestSchema ...
    type EditEmailRequestSchema struct {

        
            Email string  `json:"email"`
         
    }
    
    // SendVerificationLinkMobileRequestSchema ...
    type SendVerificationLinkMobileRequestSchema struct {

        
            Verified bool  `json:"verified"`
            Active bool  `json:"active"`
            CountryCode string  `json:"country_code"`
            Phone string  `json:"phone"`
            Primary bool  `json:"primary"`
         
    }
    
    // EditMobileRequestSchema ...
    type EditMobileRequestSchema struct {

        
            CountryCode string  `json:"country_code"`
            Phone string  `json:"phone"`
         
    }
    
    // EditProfileRequestSchema ...
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
    
    // EditProfileMobileSchema ...
    type EditProfileMobileSchema struct {

        
            Phone string  `json:"phone"`
            CountryCode string  `json:"country_code"`
         
    }
    
    // SendEmailOtpRequestSchema ...
    type SendEmailOtpRequestSchema struct {

        
            Email string  `json:"email"`
            Action string  `json:"action"`
            Token string  `json:"token"`
            RegisterToken string  `json:"register_token"`
            CaptchaCode string  `json:"captcha_code"`
         
    }
    
    // VerifyEmailOtpRequestSchema ...
    type VerifyEmailOtpRequestSchema struct {

        
            Email string  `json:"email"`
            Action string  `json:"action"`
            RegisterToken string  `json:"register_token"`
            Otp string  `json:"otp"`
         
    }
    
    // VerifyOtpRequestSchema ...
    type VerifyOtpRequestSchema struct {

        
            RequestID string  `json:"request_id"`
            RegisterToken string  `json:"register_token"`
            Otp string  `json:"otp"`
         
    }
    
    // SendMobileOtpRequestSchema ...
    type SendMobileOtpRequestSchema struct {

        
            Mobile string  `json:"mobile"`
            CountryCode string  `json:"country_code"`
            Action string  `json:"action"`
            Token string  `json:"token"`
            AndroidHash string  `json:"android_hash"`
            Force string  `json:"force"`
            CaptchaCode string  `json:"captcha_code"`
            Type string  `json:"type"`
         
    }
    
    // UpdatePasswordRequestSchema ...
    type UpdatePasswordRequestSchema struct {

        
            OldPassword string  `json:"old_password"`
            NewPassword string  `json:"new_password"`
         
    }
    
    // FormRegisterRequestSchema ...
    type FormRegisterRequestSchema struct {

        
            FirstName string  `json:"first_name"`
            LastName string  `json:"last_name"`
            Gender string  `json:"gender"`
            Email string  `json:"email"`
            Password string  `json:"password"`
            Phone FormRegisterRequestSchemaPhone  `json:"phone"`
            RegisterToken string  `json:"register_token"`
         
    }
    
    // TokenRequestBodySchema ...
    type TokenRequestBodySchema struct {

        
            Token string  `json:"token"`
         
    }
    
    // ForgotPasswordRequestSchema ...
    type ForgotPasswordRequestSchema struct {

        
            Code string  `json:"code"`
            Password string  `json:"password"`
         
    }
    
    // CodeRequestBodySchema ...
    type CodeRequestBodySchema struct {

        
            Code string  `json:"code"`
         
    }
    
    // SendResetPasswordEmailRequestSchema ...
    type SendResetPasswordEmailRequestSchema struct {

        
            Email string  `json:"email"`
            CaptchaCode string  `json:"captcha_code"`
         
    }
    
    // SendResetPasswordMobileRequestSchema ...
    type SendResetPasswordMobileRequestSchema struct {

        
            CountryCode string  `json:"country_code"`
            Mobile string  `json:"mobile"`
            CaptchaCode string  `json:"captcha_code"`
         
    }
    
    // PasswordLoginRequestSchema ...
    type PasswordLoginRequestSchema struct {

        
            CaptchaCode string  `json:"captcha_code"`
            Password string  `json:"password"`
            Username string  `json:"username"`
         
    }
    
    // SendOtpRequestSchema ...
    type SendOtpRequestSchema struct {

        
            CountryCode string  `json:"country_code"`
            CaptchaCode string  `json:"captcha_code"`
            Mobile string  `json:"mobile"`
            AndroidHash string  `json:"android_hash"`
            Type string  `json:"type"`
         
    }
    
    // OAuthRequestSchema ...
    type OAuthRequestSchema struct {

        
            IsSignedIn bool  `json:"is_signed_in"`
            Oauth2 OAuthRequestSchemaOauth2  `json:"oauth2"`
            Profile OAuthRequestSchemaProfile  `json:"profile"`
         
    }
    
    // OAuthRequestAppleSchema ...
    type OAuthRequestAppleSchema struct {

        
            UserIdentifier string  `json:"user_identifier"`
            Oauth OAuthRequestAppleSchemaOauth  `json:"oauth"`
            Profile OAuthRequestAppleSchemaProfile  `json:"profile"`
         
    }
    
    // UserObjectSchema ...
    type UserObjectSchema struct {

        
            User UserSchema  `json:"user"`
         
    }
    
    // AuthSuccess ...
    type AuthSuccess struct {

        
            RegisterToken string  `json:"register_token"`
            UserExists bool  `json:"user_exists"`
            User UserSchema  `json:"user"`
         
    }
    
    // SendOtpResponse ...
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
    
    // ProfileEditSuccess ...
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
    
    // LoginSuccess ...
    type LoginSuccess struct {

        
            User UserSchema  `json:"user"`
            RequestID string  `json:"request_id"`
            RegisterToken string  `json:"register_token"`
         
    }
    
    // VerifyOtpSuccess ...
    type VerifyOtpSuccess struct {

        
            User UserSchema  `json:"user"`
            UserExists bool  `json:"user_exists"`
            RegisterToken string  `json:"register_token"`
         
    }
    
    // ResetPasswordSuccess ...
    type ResetPasswordSuccess struct {

        
            Status string  `json:"status"`
         
    }
    
    // RegisterFormSuccess ...
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
    
    // VerifyEmailSuccess ...
    type VerifyEmailSuccess struct {

        
            Message string  `json:"message"`
         
    }
    
    // HasPasswordSuccess ...
    type HasPasswordSuccess struct {

        
            Result bool  `json:"result"`
         
    }
    
    // LogoutSuccess ...
    type LogoutSuccess struct {

        
            Logout bool  `json:"logout"`
         
    }
    
    // BlockUserSuccess ...
    type BlockUserSuccess struct {

        
            Success bool  `json:"success"`
         
    }
    
    // ArchiveUserSuccess ...
    type ArchiveUserSuccess struct {

        
            Success bool  `json:"success"`
         
    }
    
    // DeleteUserSuccess ...
    type DeleteUserSuccess struct {

        
            Success bool  `json:"success"`
         
    }
    
    // UnDeleteUserSuccess ...
    type UnDeleteUserSuccess struct {

        
            Success bool  `json:"success"`
         
    }
    
    // OtpSuccess ...
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
    
    // EmailOtpSuccess ...
    type EmailOtpSuccess struct {

        
            Success bool  `json:"success"`
         
    }
    
    // SessionListSuccess ...
    type SessionListSuccess struct {

        
            Sessions []string  `json:"sessions"`
         
    }
    
    // VerifyMobileOTPSuccess ...
    type VerifyMobileOTPSuccess struct {

        
            User UserSchema  `json:"user"`
            VerifyMobileLink bool  `json:"verify_mobile_link"`
         
    }
    
    // VerifyEmailOTPSuccess ...
    type VerifyEmailOTPSuccess struct {

        
            User UserSchema  `json:"user"`
            VerifyEmailLink bool  `json:"verify_email_link"`
         
    }
    
    // SendMobileVerifyLinkSuccess ...
    type SendMobileVerifyLinkSuccess struct {

        
            VerifyMobileLink bool  `json:"verify_mobile_link"`
         
    }
    
    // SendEmailVerifyLinkSuccess ...
    type SendEmailVerifyLinkSuccess struct {

        
            VerifyEmailLink bool  `json:"verify_email_link"`
         
    }
    
    // UserSearchResponseSchema ...
    type UserSearchResponseSchema struct {

        
            Users []UserSchema  `json:"users"`
         
    }
    
    // CustomerListResponseSchema ...
    type CustomerListResponseSchema struct {

        
            Items []UserSchema  `json:"items"`
            Page PaginationSchema  `json:"page"`
         
    }
    
    // PaginationSchema ...
    type PaginationSchema struct {

        
            Size float64  `json:"size"`
            ItemTotal float64  `json:"item_total"`
            HasNext bool  `json:"has_next"`
            Type string  `json:"type"`
            Current float64  `json:"current"`
         
    }
    
    // SessionListResponseSchema ...
    type SessionListResponseSchema struct {

        
            Items []SessionListResponseInfo  `json:"items"`
         
    }
    
    // SessionDeleteResponseSchema ...
    type SessionDeleteResponseSchema struct {

        
            Items []string  `json:"items"`
         
    }
    
    // UnauthorizedSchema ...
    type UnauthorizedSchema struct {

        
            Message string  `json:"message"`
         
    }
    
    // UnauthenticatedSchema ...
    type UnauthenticatedSchema struct {

        
            Authenticated bool  `json:"authenticated"`
         
    }
    
    // NotFoundSchema ...
    type NotFoundSchema struct {

        
            Message string  `json:"message"`
         
    }
    
    // AuthenticationInternalServerErrorSchema ...
    type AuthenticationInternalServerErrorSchema struct {

        
            Message string  `json:"message"`
         
    }
    
    // AuthenticationApiErrorSchema ...
    type AuthenticationApiErrorSchema struct {

        
            Message string  `json:"message"`
         
    }
    
    // ProfileEditSuccessSchema ...
    type ProfileEditSuccessSchema struct {

        
            Email string  `json:"email"`
            VerifyEmailOtp bool  `json:"verify_email_otp"`
            VerifyEmailLink bool  `json:"verify_email_link"`
            VerifyMobileOtp bool  `json:"verify_mobile_otp"`
            User string  `json:"user"`
            RegisterToken string  `json:"register_token"`
            UserExists bool  `json:"user_exists"`
         
    }
    
    // FormRegisterRequestSchemaPhone ...
    type FormRegisterRequestSchemaPhone struct {

        
            CountryCode string  `json:"country_code"`
            Mobile string  `json:"mobile"`
         
    }
    
    // OAuthRequestSchemaOauth2 ...
    type OAuthRequestSchemaOauth2 struct {

        
            AccessToken string  `json:"access_token"`
            Expiry float64  `json:"expiry"`
            RefreshToken string  `json:"refresh_token"`
         
    }
    
    // OAuthRequestSchemaProfile ...
    type OAuthRequestSchemaProfile struct {

        
            LastName string  `json:"last_name"`
            Image string  `json:"image"`
            ID string  `json:"id"`
            Email string  `json:"email"`
            FullName string  `json:"full_name"`
            FirstName string  `json:"first_name"`
         
    }
    
    // OAuthRequestAppleSchemaOauth ...
    type OAuthRequestAppleSchemaOauth struct {

        
            IdentityToken string  `json:"identity_token"`
         
    }
    
    // OAuthRequestAppleSchemaProfile ...
    type OAuthRequestAppleSchemaProfile struct {

        
            LastName string  `json:"last_name"`
            FullName string  `json:"full_name"`
            FirstName string  `json:"first_name"`
         
    }
    
    // AuthSuccessUser ...
    type AuthSuccessUser struct {

        
            FirstName string  `json:"first_name"`
            LastName string  `json:"last_name"`
            Debug AuthSuccessUserDebug  `json:"debug"`
            Active bool  `json:"active"`
            Emails AuthSuccessUserEmails  `json:"emails"`
         
    }
    
    // SessionListResponseInfo ...
    type SessionListResponseInfo struct {

        
            SessionID string  `json:"session_id"`
            UserAgent string  `json:"user_agent"`
            Ip string  `json:"ip"`
            Domain string  `json:"domain"`
            ExpireIn string  `json:"expire_in"`
         
    }
    
    // AuthSuccessUserDebug ...
    type AuthSuccessUserDebug struct {

        
            Platform string  `json:"platform"`
         
    }
    
    // AuthSuccessUserEmails ...
    type AuthSuccessUserEmails struct {

        
            Email string  `json:"email"`
            Verified bool  `json:"verified"`
            Primary bool  `json:"primary"`
            Active bool  `json:"active"`
         
    }
    
    // UserGroupResponseSchema ...
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
    
    // UserGroupListResponseSchema ...
    type UserGroupListResponseSchema struct {

        
            Items []UserGroupResponseSchema  `json:"items"`
            Page PaginationSchema  `json:"page"`
         
    }
    
    // CreateUserGroupSchema ...
    type CreateUserGroupSchema struct {

        
            Name string  `json:"name"`
            Description string  `json:"description"`
            FileURL string  `json:"file_url"`
         
    }
    
    // CreateUserRequestSchema ...
    type CreateUserRequestSchema struct {

        
            PhoneNumber string  `json:"phone_number"`
            Email string  `json:"email"`
            FirstName string  `json:"first_name"`
            LastName string  `json:"last_name"`
            Gender string  `json:"gender"`
            Username string  `json:"username"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // CreateUserResponseSchema ...
    type CreateUserResponseSchema struct {

        
            User UserSchema  `json:"user"`
         
    }
    
    // CreateUserSessionRequestSchema ...
    type CreateUserSessionRequestSchema struct {

        
            Domain string  `json:"domain"`
            MaxAge float64  `json:"max_age"`
            UserID string  `json:"user_id"`
         
    }
    
    // CreateUserSessionResponseSchema ...
    type CreateUserSessionResponseSchema struct {

        
            Domain string  `json:"domain"`
            MaxAge float64  `json:"max_age"`
            Secure bool  `json:"secure"`
            HttpOnly bool  `json:"http_only"`
            Cookie map[string]interface{}  `json:"cookie"`
         
    }
    
    // PlatformSchema ...
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
    
    // LookAndFeel ...
    type LookAndFeel struct {

        
            CardPosition string  `json:"card_position"`
            BackgroundColor string  `json:"background_color"`
         
    }
    
    // Login ...
    type Login struct {

        
            Password bool  `json:"password"`
            Otp bool  `json:"otp"`
            VoiceOtp bool  `json:"voice_otp"`
         
    }
    
    // MetaSchema ...
    type MetaSchema struct {

        
            FyndDefault bool  `json:"fynd_default"`
         
    }
    
    // Social ...
    type Social struct {

        
            AccountKit bool  `json:"account_kit"`
            Facebook bool  `json:"facebook"`
            Google bool  `json:"google"`
            Apple bool  `json:"apple"`
         
    }
    
    // RequiredFields ...
    type RequiredFields struct {

        
            Email PlatformEmail  `json:"email"`
            Mobile PlatformMobile  `json:"mobile"`
         
    }
    
    // PlatformEmail ...
    type PlatformEmail struct {

        
            IsRequired bool  `json:"is_required"`
            Level string  `json:"level"`
         
    }
    
    // PlatformMobile ...
    type PlatformMobile struct {

        
            IsRequired bool  `json:"is_required"`
            Level string  `json:"level"`
         
    }
    
    // RegisterRequiredFields ...
    type RegisterRequiredFields struct {

        
            Email RegisterRequiredFieldsEmail  `json:"email"`
            Mobile RegisterRequiredFieldsMobile  `json:"mobile"`
         
    }
    
    // RegisterRequiredFieldsEmail ...
    type RegisterRequiredFieldsEmail struct {

        
            IsRequired bool  `json:"is_required"`
            Level string  `json:"level"`
         
    }
    
    // RegisterRequiredFieldsMobile ...
    type RegisterRequiredFieldsMobile struct {

        
            IsRequired bool  `json:"is_required"`
            Level string  `json:"level"`
         
    }
    
    // FlashCard ...
    type FlashCard struct {

        
            Text string  `json:"text"`
            TextColor string  `json:"text_color"`
            BackgroundColor string  `json:"background_color"`
         
    }
    
    // SocialTokens ...
    type SocialTokens struct {

        
            Facebook Facebook  `json:"facebook"`
            AccountKit Accountkit  `json:"account_kit"`
            Google Google  `json:"google"`
         
    }
    
    // DeleteAccountReasons ...
    type DeleteAccountReasons struct {

        
            ReasonText string  `json:"reason_text"`
            ReasonID string  `json:"reason_id"`
            ShowTextArea bool  `json:"show_text_area"`
         
    }
    
    // DeleteAccountConsent ...
    type DeleteAccountConsent struct {

        
            ConsentText string  `json:"consent_text"`
         
    }
    
    // Facebook ...
    type Facebook struct {

        
            AppID string  `json:"app_id"`
         
    }
    
    // Accountkit ...
    type Accountkit struct {

        
            AppID string  `json:"app_id"`
         
    }
    
    // Google ...
    type Google struct {

        
            AppID string  `json:"app_id"`
         
    }
    
    // SessionExpiry ...
    type SessionExpiry struct {

        
            Duration float64  `json:"duration"`
            Type string  `json:"type"`
            IsRolling bool  `json:"is_rolling"`
         
    }
    
    // UpdateUserGroupSchema ...
    type UpdateUserGroupSchema struct {

        
            Name string  `json:"name"`
            Description string  `json:"description"`
            FileURL string  `json:"file_url"`
         
    }
    
    // UpdateUserRequestSchema ...
    type UpdateUserRequestSchema struct {

        
            FirstName string  `json:"first_name"`
            LastName string  `json:"last_name"`
            Gender string  `json:"gender"`
            ExternalID string  `json:"external_id"`
            Meta map[string]interface{}  `json:"meta"`
            PhoneNumbers []UserPhoneNumbers  `json:"phone_numbers"`
            Emails []UserEmails  `json:"emails"`
         
    }
    
    // UserEmails ...
    type UserEmails struct {

        
            Active bool  `json:"active"`
            Primary bool  `json:"primary"`
            Verified bool  `json:"verified"`
            Email string  `json:"email"`
         
    }
    
    // UserPhoneNumbers ...
    type UserPhoneNumbers struct {

        
            Active bool  `json:"active"`
            Primary bool  `json:"primary"`
            Verified bool  `json:"verified"`
            Phone string  `json:"phone"`
            CountryCode string  `json:"country_code"`
         
    }
    
    // UserSchema ...
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
    
    // PhoneNumber ...
    type PhoneNumber struct {

        
            Active bool  `json:"active"`
            Primary bool  `json:"primary"`
            Verified bool  `json:"verified"`
            Phone string  `json:"phone"`
            CountryCode float64  `json:"country_code"`
         
    }
    
    // Email ...
    type Email struct {

        
            Primary bool  `json:"primary"`
            Verified bool  `json:"verified"`
            Email string  `json:"email"`
            Active bool  `json:"active"`
         
    }
    

    
    // ApplicationLegal ...
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
    
    // ApplicationLegalFAQ ...
    type ApplicationLegalFAQ struct {

        
            Question string  `json:"question"`
            Answer string  `json:"answer"`
         
    }
    
    // PathMappingSchema ...
    type PathMappingSchema struct {

        
            Application string  `json:"application"`
            ID string  `json:"_id"`
            RedirectFrom string  `json:"redirect_from"`
            RedirectTo string  `json:"redirect_to"`
            UpdatedAt string  `json:"updated_at"`
            CreatedAt string  `json:"created_at"`
            Source PathSourceSchema  `json:"__source"`
         
    }
    
    // PathSourceSchema ...
    type PathSourceSchema struct {

        
            Type string  `json:"type"`
            ID string  `json:"id"`
         
    }
    
    // SeoComponent ...
    type SeoComponent struct {

        
            Seo SeoSchema  `json:"seo"`
         
    }
    
    // SeoSchema ...
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
    
    // CustomMetaTag ...
    type CustomMetaTag struct {

        
            Name string  `json:"name"`
            Content string  `json:"content"`
            ID string  `json:"_id"`
         
    }
    
    // Detail ...
    type Detail struct {

        
            Title string  `json:"title"`
            Description string  `json:"description"`
            ImageURL string  `json:"image_url"`
         
    }
    
    // AnnouncementPageSchema ...
    type AnnouncementPageSchema struct {

        
            PageSlug string  `json:"page_slug"`
            Type string  `json:"type"`
         
    }
    
    // EditorMeta ...
    type EditorMeta struct {

        
            ForegroundColor string  `json:"foreground_color"`
            BackgroundColor string  `json:"background_color"`
            ContentType string  `json:"content_type"`
            Content string  `json:"content"`
         
    }
    
    // AnnouncementAuthorSchema ...
    type AnnouncementAuthorSchema struct {

        
            CreatedBy string  `json:"created_by"`
            ModifiedBy string  `json:"modified_by"`
         
    }
    
    // AdminAnnouncementSchema ...
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
    
    // ScheduleSchema ...
    type ScheduleSchema struct {

        
            Cron string  `json:"cron"`
            Start string  `json:"start"`
            End string  `json:"end"`
            Duration float64  `json:"duration"`
            NextSchedule []map[string]interface{}  `json:"next_schedule"`
         
    }
    
    // NextSchedule ...
    type NextSchedule struct {

        
            Start string  `json:"start"`
            End string  `json:"end"`
         
    }
    
    // AnnouncementSchema ...
    type AnnouncementSchema struct {

        
            Announcement string  `json:"announcement"`
            Schedule ScheduleStartSchema  `json:"schedule"`
         
    }
    
    // ScheduleStartSchema ...
    type ScheduleStartSchema struct {

        
            Start string  `json:"start"`
            End string  `json:"end"`
         
    }
    
    // BlogGetResponse ...
    type BlogGetResponse struct {

        
            Items []BlogSchema  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // ResourceContent ...
    type ResourceContent struct {

        
            Type string  `json:"type"`
            Value string  `json:"value"`
         
    }
    
    // Asset ...
    type Asset struct {

        
            AspectRatio string  `json:"aspect_ratio"`
            ID string  `json:"id"`
            SecureURL string  `json:"secure_url"`
         
    }
    
    // Author ...
    type Author struct {

        
            Designation string  `json:"designation"`
            ID string  `json:"id"`
            Name string  `json:"name"`
         
    }
    
    // BlogSchema ...
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
    
    // SEO ...
    type SEO struct {

        
            Description string  `json:"description"`
            Image SEOImage  `json:"image"`
            Title string  `json:"title"`
         
    }
    
    // SEOImage ...
    type SEOImage struct {

        
            URL string  `json:"url"`
         
    }
    
    // DateMeta ...
    type DateMeta struct {

        
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // BlogRequest ...
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
    
    // GetAnnouncementListSchema ...
    type GetAnnouncementListSchema struct {

        
            Items []AdminAnnouncementSchema  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // CreateAnnouncementSchema ...
    type CreateAnnouncementSchema struct {

        
            Message string  `json:"message"`
            Data AdminAnnouncementSchema  `json:"data"`
         
    }
    
    // DataLoaderResponseSchema ...
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
    
    // DataLoaderResetResponseSchema ...
    type DataLoaderResetResponseSchema struct {

        
            Reset string  `json:"reset"`
         
    }
    
    // Navigation ...
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
    
    // LocaleLanguage ...
    type LocaleLanguage struct {

        
            Hi Language  `json:"hi"`
            Ar Language  `json:"ar"`
            EnUs Language  `json:"en_us"`
         
    }
    
    // Language ...
    type Language struct {

        
            Display string  `json:"display"`
         
    }
    
    // Action ...
    type Action struct {

        
            Page ActionPage  `json:"page"`
            Popup ActionPage  `json:"popup"`
            Type string  `json:"type"`
         
    }
    
    // ActionPage ...
    type ActionPage struct {

        
            Params map[string][]string  `json:"params"`
            Query map[string][]string  `json:"query"`
            URL string  `json:"url"`
            Type PageType  `json:"type"`
         
    }
    
    // NavigationReference ...
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
    
    // LandingPage ...
    type LandingPage struct {

        
            Data LandingPageSchema  `json:"data"`
            Success bool  `json:"success"`
         
    }
    
    // ConfigurationSchema ...
    type ConfigurationSchema struct {

        
            SleepTime float64  `json:"sleep_time"`
            StartOnLaunch bool  `json:"start_on_launch"`
            Duration float64  `json:"duration"`
            SlideDirection string  `json:"slide_direction"`
         
    }
    
    // SlideshowMedia ...
    type SlideshowMedia struct {

        
            Type string  `json:"type"`
            URL string  `json:"url"`
            BgColor string  `json:"bg_color"`
            Duration float64  `json:"duration"`
            AutoDecideDuration bool  `json:"auto_decide_duration"`
            Action Action  `json:"action"`
         
    }
    
    // Slideshow ...
    type Slideshow struct {

        
            Data SlideshowSchema  `json:"data"`
            Success bool  `json:"success"`
         
    }
    
    // AnnouncementsResponseSchema ...
    type AnnouncementsResponseSchema struct {

        
            Announcements map[string][]AnnouncementSchema  `json:"announcements"`
            RefreshRate float64  `json:"refresh_rate"`
            RefreshPages []string  `json:"refresh_pages"`
         
    }
    
    // FaqResponseSchema ...
    type FaqResponseSchema struct {

        
            Faqs []FaqSchema  `json:"faqs"`
         
    }
    
    // UpdateHandpickedSchema ...
    type UpdateHandpickedSchema struct {

        
            Tag HandpickedTagSchema  `json:"tag"`
         
    }
    
    // HandpickedTagSchema ...
    type HandpickedTagSchema struct {

        
            Position string  `json:"position"`
            Attributes map[string]interface{}  `json:"attributes"`
            Name string  `json:"name"`
            URL string  `json:"url"`
            Type string  `json:"type"`
            SubType string  `json:"sub_type"`
            Content string  `json:"content"`
         
    }
    
    // RemoveHandpickedSchema ...
    type RemoveHandpickedSchema struct {

        
            Tags []string  `json:"tags"`
         
    }
    
    // CreateTagSchema ...
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
    
    // CreateTagRequestSchema ...
    type CreateTagRequestSchema struct {

        
            Tags []CreateTagSchema  `json:"tags"`
         
    }
    
    // DataLoaderSchema ...
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
    
    // DataLoaderSourceSchema ...
    type DataLoaderSourceSchema struct {

        
            Type string  `json:"type"`
            ID string  `json:"id"`
         
    }
    
    // DataLoadersSchema ...
    type DataLoadersSchema struct {

        
            Items []DataLoaderSchema  `json:"items"`
         
    }
    
    // TagDeleteSuccessResponse ...
    type TagDeleteSuccessResponse struct {

        
            Success bool  `json:"success"`
         
    }
    
    // ContentAPIError ...
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
    
    // CommonError ...
    type CommonError struct {

        
            Message string  `json:"message"`
         
    }
    
    // CategorySchema ...
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
    
    // ChildrenSchema ...
    type ChildrenSchema struct {

        
            Question string  `json:"question"`
            Answer string  `json:"answer"`
            Slug string  `json:"slug"`
            Application string  `json:"application"`
            ID string  `json:"_id"`
         
    }
    
    // CategoryRequestSchema ...
    type CategoryRequestSchema struct {

        
            Slug string  `json:"slug"`
            Title string  `json:"title"`
         
    }
    
    // FAQCategorySchema ...
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
    
    // FaqSchema ...
    type FaqSchema struct {

        
            Slug string  `json:"slug"`
            Application string  `json:"application"`
            ID string  `json:"_id"`
            Question string  `json:"question"`
            Answer string  `json:"answer"`
            Tags []string  `json:"tags"`
         
    }
    
    // FAQ ...
    type FAQ struct {

        
            Slug string  `json:"slug"`
            Question string  `json:"question"`
            Answer string  `json:"answer"`
         
    }
    
    // CreateFaqResponseSchema ...
    type CreateFaqResponseSchema struct {

        
            Faq FaqSchema  `json:"faq"`
         
    }
    
    // CreateFaqSchema ...
    type CreateFaqSchema struct {

        
            Faq FAQ  `json:"faq"`
         
    }
    
    // GetFaqSchema ...
    type GetFaqSchema struct {

        
            Faqs []FaqSchema  `json:"faqs"`
         
    }
    
    // UpdateFaqCategoryRequestSchema ...
    type UpdateFaqCategoryRequestSchema struct {

        
            Category CategorySchema  `json:"category"`
         
    }
    
    // CreateFaqCategoryRequestSchema ...
    type CreateFaqCategoryRequestSchema struct {

        
            Category CategoryRequestSchema  `json:"category"`
         
    }
    
    // CreateFaqCategorySchema ...
    type CreateFaqCategorySchema struct {

        
            Category CategorySchema  `json:"category"`
         
    }
    
    // GetFaqCategoriesSchema ...
    type GetFaqCategoriesSchema struct {

        
            Categories []CategorySchema  `json:"categories"`
         
    }
    
    // GetFaqCategoryBySlugSchema ...
    type GetFaqCategoryBySlugSchema struct {

        
            Category FAQCategorySchema  `json:"category"`
         
    }
    
    // Page ...
    type Page struct {

        
            ItemTotal float64  `json:"item_total"`
            NextID string  `json:"next_id"`
            HasPrevious bool  `json:"has_previous"`
            HasNext bool  `json:"has_next"`
            Current float64  `json:"current"`
            Type string  `json:"type"`
            Size float64  `json:"size"`
         
    }
    
    // LandingPageGetResponse ...
    type LandingPageGetResponse struct {

        
            Items []LandingPageSchema  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // LandingPageSchema ...
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
    
    // DefaultNavigationResponse ...
    type DefaultNavigationResponse struct {

        
            Items []NavigationSchema  `json:"items"`
         
    }
    
    // NavigationGetResponse ...
    type NavigationGetResponse struct {

        
            Items []NavigationSchema  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // Orientation ...
    type Orientation struct {

        
            Portrait []string  `json:"portrait"`
            Landscape []string  `json:"landscape"`
         
    }
    
    // NavigationSchema ...
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
    
    // NavigationRequest ...
    type NavigationRequest struct {

        
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Platform []string  `json:"platform"`
            Orientation Orientation  `json:"orientation"`
            Navigation []NavigationReference  `json:"navigation"`
         
    }
    
    // CustomPageSchema ...
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
    
    // ContentSchema ...
    type ContentSchema struct {

        
            Type string  `json:"type"`
            Value map[string]interface{}  `json:"value"`
         
    }
    
    // CustomPage ...
    type CustomPage struct {

        
            Data CustomPageSchema  `json:"data"`
         
    }
    
    // FeatureImage ...
    type FeatureImage struct {

        
            SecureURL string  `json:"secure_url"`
         
    }
    
    // PageGetResponse ...
    type PageGetResponse struct {

        
            Items []PageSchema  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // PageSpec ...
    type PageSpec struct {

        
            Specifications []map[string]interface{}  `json:"specifications"`
         
    }
    
    // PageSpecParam ...
    type PageSpecParam struct {

        
            Key string  `json:"key"`
            Required bool  `json:"required"`
         
    }
    
    // PageSpecItem ...
    type PageSpecItem struct {

        
            PageType string  `json:"page_type"`
            DisplayName string  `json:"display_name"`
            Params []PageSpecParam  `json:"params"`
            Query []PageSpecParam  `json:"query"`
         
    }
    
    // PageSchema ...
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
    
    // CreatedBySchema ...
    type CreatedBySchema struct {

        
            ID string  `json:"id"`
         
    }
    
    // PageContent ...
    type PageContent struct {

        
            Type string  `json:"type"`
            Value map[string]interface{}  `json:"value"`
         
    }
    
    // PageMeta ...
    type PageMeta struct {

        
            Key string  `json:"key"`
            Value map[string]interface{}  `json:"value"`
         
    }
    
    // PageRequest ...
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
    
    // CronSchedule ...
    type CronSchedule struct {

        
            Cron string  `json:"cron"`
            Start string  `json:"start"`
            End string  `json:"end"`
            Duration float64  `json:"duration"`
         
    }
    
    // PagePublishRequest ...
    type PagePublishRequest struct {

        
            Publish bool  `json:"publish"`
         
    }
    
    // PageMetaSchema ...
    type PageMetaSchema struct {

        
            SystemPages []NavigationSchema  `json:"system_pages"`
            CustomPages []PageSchema  `json:"custom_pages"`
            ApplicationID string  `json:"application_id"`
         
    }
    
    // SlideshowGetResponse ...
    type SlideshowGetResponse struct {

        
            Items []SlideshowSchema  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // SlideshowSchema ...
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
    
    // SlideshowRequest ...
    type SlideshowRequest struct {

        
            Slug string  `json:"slug"`
            Platform string  `json:"platform"`
            Configuration ConfigurationSchema  `json:"configuration"`
            Media []SlideshowMedia  `json:"media"`
            Active bool  `json:"active"`
         
    }
    
    // Support ...
    type Support struct {

        
            Created bool  `json:"created"`
            ID string  `json:"_id"`
            ConfigType string  `json:"config_type"`
            Application string  `json:"application"`
            CreatedAt string  `json:"created_at"`
            UpdatedAt string  `json:"updated_at"`
            Contact ContactSchema  `json:"contact"`
         
    }
    
    // PhoneProperties ...
    type PhoneProperties struct {

        
            Key string  `json:"key"`
            Code string  `json:"code"`
            Number string  `json:"number"`
         
    }
    
    // PhoneSchema ...
    type PhoneSchema struct {

        
            Active bool  `json:"active"`
            Phone []PhoneProperties  `json:"phone"`
         
    }
    
    // EmailProperties ...
    type EmailProperties struct {

        
            Key string  `json:"key"`
            Value string  `json:"value"`
         
    }
    
    // EmailSchema ...
    type EmailSchema struct {

        
            Active bool  `json:"active"`
            Email []EmailProperties  `json:"email"`
         
    }
    
    // ContactSchema ...
    type ContactSchema struct {

        
            Phone PhoneSchema  `json:"phone"`
            Email EmailSchema  `json:"email"`
         
    }
    
    // TagsSchema ...
    type TagsSchema struct {

        
            Application string  `json:"application"`
            ID string  `json:"_id"`
            Tags []TagSchema  `json:"tags"`
         
    }
    
    // TagSchema ...
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
    
    // TagSourceSchema ...
    type TagSourceSchema struct {

        
            Type string  `json:"type"`
            ID string  `json:"id"`
         
    }
    

    
    // CommunicationConsentReq ...
    type CommunicationConsentReq struct {

        
            Response string  `json:"response"`
            Action string  `json:"action"`
            Channel string  `json:"channel"`
         
    }
    
    // CommunicationConsentRes ...
    type CommunicationConsentRes struct {

        
            AppID string  `json:"app_id"`
            UserID string  `json:"user_id"`
            Channels CommunicationConsentChannels  `json:"channels"`
         
    }
    
    // CommunicationConsentChannelsEmail ...
    type CommunicationConsentChannelsEmail struct {

        
            Response string  `json:"response"`
            DisplayName string  `json:"display_name"`
         
    }
    
    // CommunicationConsentChannelsSms ...
    type CommunicationConsentChannelsSms struct {

        
            Response string  `json:"response"`
            DisplayName string  `json:"display_name"`
         
    }
    
    // CommunicationConsentChannelsWhatsapp ...
    type CommunicationConsentChannelsWhatsapp struct {

        
            Response string  `json:"response"`
            DisplayName string  `json:"display_name"`
            CountryCode string  `json:"country_code"`
            PhoneNumber string  `json:"phone_number"`
         
    }
    
    // CommunicationConsentChannels ...
    type CommunicationConsentChannels struct {

        
            Email CommunicationConsentChannelsEmail  `json:"email"`
            Sms CommunicationConsentChannelsSms  `json:"sms"`
            Whatsapp CommunicationConsentChannelsWhatsapp  `json:"whatsapp"`
         
    }
    
    // CommunicationConsent ...
    type CommunicationConsent struct {

        
            AppID string  `json:"app_id"`
            UserID string  `json:"user_id"`
            Channels CommunicationConsentChannels  `json:"channels"`
         
    }
    
    // BadRequestSchema ...
    type BadRequestSchema struct {

        
            Status string  `json:"status"`
            Message string  `json:"message"`
         
    }
    
    // NotFound ...
    type NotFound struct {

        
            Message string  `json:"message"`
         
    }
    
    // PushtokenReq ...
    type PushtokenReq struct {

        
            Action string  `json:"action"`
            BundleIdentifier string  `json:"bundle_identifier"`
            PushToken string  `json:"push_token"`
            UniqueDeviceID string  `json:"unique_device_id"`
            Type string  `json:"type"`
         
    }
    
    // PushtokenRes ...
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
    
    // Page ...
    type Page struct {

        
            ItemTotal float64  `json:"item_total"`
            NextID string  `json:"next_id"`
            HasPrevious bool  `json:"has_previous"`
            HasNext bool  `json:"has_next"`
            Current float64  `json:"current"`
            Type string  `json:"type"`
            Size float64  `json:"size"`
         
    }
    
    // GenericPage ...
    type GenericPage struct {

        
            Type string  `json:"type"`
            Size float64  `json:"size"`
            Current float64  `json:"current"`
            HasNext bool  `json:"has_next"`
            ItemTotal float64  `json:"item_total"`
         
    }
    
    // GenericSuccess ...
    type GenericSuccess struct {

        
            Success bool  `json:"success"`
         
    }
    
    // GenericError ...
    type GenericError struct {

        
            Message Message  `json:"message"`
            Sentry string  `json:"sentry"`
         
    }
    
    // GenericDelete ...
    type GenericDelete struct {

        
            Message string  `json:"message"`
            Acknowledged bool  `json:"acknowledged"`
            Affected float64  `json:"affected"`
            Operation string  `json:"operation"`
         
    }
    
    // Message ...
    type Message struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
            Info string  `json:"info"`
            Operation string  `json:"operation"`
         
    }
    
    // EnabledObj ...
    type EnabledObj struct {

        
            Enabled bool  `json:"enabled"`
         
    }
    
    // InvalidRangeErrorReqPositive ...
    type InvalidRangeErrorReqPositive struct {

        
            Message string  `json:"message"`
            Code float64  `json:"code"`
            Sentry string  `json:"sentry"`
         
    }
    
    // InvalidInputRequiredByteOrHexError ...
    type InvalidInputRequiredByteOrHexError struct {

        
            Message string  `json:"message"`
            Sentry string  `json:"sentry"`
         
    }
    
    // NameValidatorError ...
    type NameValidatorError struct {

        
            Message NameValidatorErrorMessage  `json:"message"`
            Sentry string  `json:"sentry"`
         
    }
    
    // NameValidatorErrorMessage ...
    type NameValidatorErrorMessage struct {

        
            Name ValidatorErrorBody  `json:"name"`
         
    }
    
    // ApikeyValidatorError ...
    type ApikeyValidatorError struct {

        
            Message ApikeyValidatorErrorMessage  `json:"message"`
            Sentry string  `json:"sentry"`
         
    }
    
    // ApikeyValidatorErrorMessage ...
    type ApikeyValidatorErrorMessage struct {

        
            APIKey ValidatorErrorBody  `json:"api_key"`
         
    }
    
    // FeedidValidatorError ...
    type FeedidValidatorError struct {

        
            Message FeedidValidatorErrorMessage  `json:"message"`
            Sentry string  `json:"sentry"`
         
    }
    
    // FeedidValidatorErrorMessage ...
    type FeedidValidatorErrorMessage struct {

        
            Feedid ValidatorErrorBody  `json:"feedid"`
         
    }
    
    // UsernameValidatorError ...
    type UsernameValidatorError struct {

        
            Message UsernameValidatorErrorMessage  `json:"message"`
            Sentry string  `json:"sentry"`
         
    }
    
    // UsernameValidatorErrorMessage ...
    type UsernameValidatorErrorMessage struct {

        
            Username ValidatorErrorBody  `json:"username"`
         
    }
    
    // PasswordValidatorError ...
    type PasswordValidatorError struct {

        
            Message PasswordValidatorErrorMessage  `json:"message"`
            Sentry string  `json:"sentry"`
         
    }
    
    // PasswordValidatorErrorMessage ...
    type PasswordValidatorErrorMessage struct {

        
            Password ValidatorErrorBody  `json:"password"`
         
    }
    
    // ValidatorErrorBody ...
    type ValidatorErrorBody struct {

        
            Name string  `json:"name"`
            Message string  `json:"message"`
            Properties ValidatorErrorMessageProperties  `json:"properties"`
            Kind string  `json:"kind"`
            Path string  `json:"path"`
         
    }
    
    // ValidatorErrorMessageProperties ...
    type ValidatorErrorMessageProperties struct {

        
            Message string  `json:"message"`
            Type string  `json:"type"`
            Path string  `json:"path"`
         
    }
    
    // CastToStringFail ...
    type CastToStringFail struct {

        
            Message string  `json:"message"`
            Sentry string  `json:"sentry"`
         
    }
    
    // InvalidID ...
    type InvalidID struct {

        
            Message string  `json:"message"`
            Sentry string  `json:"sentry"`
         
    }
    

    
    // QRCodeResp ...
    type QRCodeResp struct {

        
            Link string  `json:"link"`
            Svg string  `json:"svg"`
         
    }
    
    // RedirectDevice ...
    type RedirectDevice struct {

        
            Link string  `json:"link"`
            Type string  `json:"type"`
         
    }
    
    // WebRedirect ...
    type WebRedirect struct {

        
            Link string  `json:"link"`
            Type string  `json:"type"`
         
    }
    
    // Redirects ...
    type Redirects struct {

        
            Ios RedirectDevice  `json:"ios"`
            Android RedirectDevice  `json:"android"`
            Web WebRedirect  `json:"web"`
            ForceWeb bool  `json:"force_web"`
         
    }
    
    // CampaignShortLink ...
    type CampaignShortLink struct {

        
            Source string  `json:"source"`
            Medium string  `json:"medium"`
         
    }
    
    // Attribution ...
    type Attribution struct {

        
            CampaignCookieExpiry string  `json:"campaign_cookie_expiry"`
         
    }
    
    // SocialMediaTags ...
    type SocialMediaTags struct {

        
            Title string  `json:"title"`
            Description string  `json:"description"`
            Image string  `json:"image"`
         
    }
    
    // ShortLinkReq ...
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
    
    // UrlInfo ...
    type UrlInfo struct {

        
            Original string  `json:"original"`
            Short string  `json:"short"`
            Hash string  `json:"hash"`
         
    }
    
    // ShortLinkRes ...
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
    
    // Page ...
    type Page struct {

        
            ItemTotal float64  `json:"item_total"`
            NextID string  `json:"next_id"`
            HasPrevious bool  `json:"has_previous"`
            HasNext bool  `json:"has_next"`
            Current float64  `json:"current"`
            Type string  `json:"type"`
            Size float64  `json:"size"`
         
    }
    
    // ShortLinkList ...
    type ShortLinkList struct {

        
            Items []ShortLinkRes  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // ErrorRes ...
    type ErrorRes struct {

        
            Message string  `json:"message"`
         
    }
    

    
    // FailedResponse ...
    type FailedResponse struct {

        
            Message string  `json:"message"`
         
    }
    
    // CDN ...
    type CDN struct {

        
            URL string  `json:"url"`
            AbsoluteURL string  `json:"absolute_url"`
            RelativeURL string  `json:"relative_url"`
         
    }
    
    // Upload ...
    type Upload struct {

        
            Expiry float64  `json:"expiry"`
            URL string  `json:"url"`
         
    }
    
    // StartResponse ...
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
    
    // StartRequest ...
    type StartRequest struct {

        
            FileName string  `json:"file_name"`
            ContentType string  `json:"content_type"`
            Size float64  `json:"size"`
            Tags []string  `json:"tags"`
            Params map[string]interface{}  `json:"params"`
         
    }
    
    // CompleteResponse ...
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
    
    // Opts ...
    type Opts struct {

        
            Attempts float64  `json:"attempts"`
            Timestamp float64  `json:"timestamp"`
            Delay float64  `json:"delay"`
         
    }
    
    // CopyFileTask ...
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
    
    // BulkUploadResponse ...
    type BulkUploadResponse struct {

        
            TrackingURL string  `json:"tracking_url"`
            Task CopyFileTask  `json:"task"`
         
    }
    
    // ReqConfiguration ...
    type ReqConfiguration struct {

        
            Concurrency float64  `json:"concurrency"`
         
    }
    
    // Destination ...
    type Destination struct {

        
            Namespace string  `json:"namespace"`
            Rewrite string  `json:"rewrite"`
            Basepath string  `json:"basepath"`
         
    }
    
    // BulkRequest ...
    type BulkRequest struct {

        
            Urls []string  `json:"urls"`
            Destination Destination  `json:"destination"`
            Configuration ReqConfiguration  `json:"configuration"`
         
    }
    
    // Urls ...
    type Urls struct {

        
            URL string  `json:"url"`
            SignedURL string  `json:"signed_url"`
            Expiry float64  `json:"expiry"`
         
    }
    
    // SignUrlResponse ...
    type SignUrlResponse struct {

        
            Urls []Urls  `json:"urls"`
         
    }
    
    // SignUrlRequest ...
    type SignUrlRequest struct {

        
            Expiry float64  `json:"expiry"`
            Urls []string  `json:"urls"`
         
    }
    
    // Page ...
    type Page struct {

        
            ItemTotal float64  `json:"item_total"`
            NextID string  `json:"next_id"`
            HasPrevious bool  `json:"has_previous"`
            HasNext bool  `json:"has_next"`
            Current float64  `json:"current"`
            Type string  `json:"type"`
            Size float64  `json:"size"`
         
    }
    
    // DbRecord ...
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
    
    // BrowseResponse ...
    type BrowseResponse struct {

        
            Items []DbRecord  `json:"items"`
            Page Page  `json:"page"`
         
    }
    

    
    // ApplicationAboutResponse ...
    type ApplicationAboutResponse struct {

        
            ApplicationInfo ApplicationInfo  `json:"application_info"`
            CompanyInfo CompanyInfo  `json:"company_info"`
            OwnerInfo OwnerInfo  `json:"owner_info"`
         
    }
    
    // ApplicationInfo ...
    type ApplicationInfo struct {

        
            ID string  `json:"_id"`
            Domain Domain  `json:"domain"`
            Website ApplicationWebsite  `json:"website"`
            Cors ApplicationCors  `json:"cors"`
            Description string  `json:"description"`
            Name string  `json:"name"`
            Meta ApplicationMeta  `json:"meta"`
            Token string  `json:"token"`
            Secret string  `json:"secret"`
            CreatedAt string  `json:"created_at"`
            Banner SecureUrl  `json:"banner"`
            Logo SecureUrl  `json:"logo"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // CompanyInfo ...
    type CompanyInfo struct {

        
            ID string  `json:"_id"`
            UID float64  `json:"uid"`
            CreatedOn string  `json:"created_on"`
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            Addresses []CompanyAboutAddress  `json:"addresses"`
            NotificationEmails []string  `json:"notification_emails"`
         
    }
    
    // OwnerInfo ...
    type OwnerInfo struct {

        
            ID string  `json:"_id"`
            Emails []UserEmail  `json:"emails"`
            PhoneNumbers []UserPhoneNumber  `json:"phone_numbers"`
            FirstName string  `json:"first_name"`
            LastName string  `json:"last_name"`
            ProfilePic string  `json:"profile_pic"`
         
    }
    
    // AppVersionRequest ...
    type AppVersionRequest struct {

        
            Application ApplicationVersionRequest  `json:"application"`
            Device Device  `json:"device"`
            Locale string  `json:"locale"`
            Timezone string  `json:"timezone"`
         
    }
    
    // ApplicationVersionRequest ...
    type ApplicationVersionRequest struct {

        
            ID string  `json:"id"`
            Name string  `json:"name"`
            Namespace string  `json:"namespace"`
            Token string  `json:"token"`
            Version string  `json:"version"`
         
    }
    
    // Device ...
    type Device struct {

        
            Build float64  `json:"build"`
            Model string  `json:"model"`
            Os OS  `json:"os"`
         
    }
    
    // OS ...
    type OS struct {

        
            Name string  `json:"name"`
            Version string  `json:"version"`
         
    }
    
    // SupportedLanguage ...
    type SupportedLanguage struct {

        
            Name string  `json:"name"`
            Code string  `json:"code"`
         
    }
    
    // LanguageResponse ...
    type LanguageResponse struct {

        
            Items []SupportedLanguage  `json:"items"`
         
    }
    
    // AppStaffResponse ...
    type AppStaffResponse struct {

        
            StaffUsers []AppStaff  `json:"staff_users"`
         
    }
    
    // AppStaffListResponse ...
    type AppStaffListResponse struct {

        
            Page Page  `json:"page"`
            Items []AppStaff  `json:"items"`
         
    }
    
    // UpdateDialog ...
    type UpdateDialog struct {

        
            Type string  `json:"type"`
            Interval float64  `json:"interval"`
         
    }
    
    // OrderingStoreSelectRequest ...
    type OrderingStoreSelectRequest struct {

        
            OrderingStore OrderingStoreSelect  `json:"ordering_store"`
         
    }
    
    // OrderingStoreSelect ...
    type OrderingStoreSelect struct {

        
            UID float64  `json:"uid"`
         
    }
    
    // AppStaff ...
    type AppStaff struct {

        
            ID string  `json:"_id"`
            OrderIncent bool  `json:"order_incent"`
            Stores []float64  `json:"stores"`
            Application string  `json:"application"`
            Title string  `json:"title"`
            User string  `json:"user"`
            EmployeeCode string  `json:"employee_code"`
            FirstName string  `json:"first_name"`
            LastName string  `json:"last_name"`
            ProfilePicURL string  `json:"profile_pic_url"`
         
    }
    
    // AppTokenResponse ...
    type AppTokenResponse struct {

        
            Tokens Tokens  `json:"tokens"`
            ID string  `json:"_id"`
            Application string  `json:"application"`
            CreatedAt string  `json:"created_at"`
            UpdatedAt string  `json:"updated_at"`
            V float64  `json:"__v"`
         
    }
    
    // Tokens ...
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
    
    // Firebase ...
    type Firebase struct {

        
            Credentials Credentials  `json:"credentials"`
            Enabled bool  `json:"enabled"`
         
    }
    
    // Credentials ...
    type Credentials struct {

        
            Ios Ios  `json:"ios"`
            Android Android  `json:"android"`
            ProjectID string  `json:"project_id"`
            GcmSenderID string  `json:"gcm_sender_id"`
            ApplicationID string  `json:"application_id"`
            APIKey string  `json:"api_key"`
         
    }
    
    // Ios ...
    type Ios struct {

        
            ApplicationID string  `json:"application_id"`
            APIKey string  `json:"api_key"`
         
    }
    
    // Android ...
    type Android struct {

        
            ApplicationID string  `json:"application_id"`
            APIKey string  `json:"api_key"`
         
    }
    
    // Moengage ...
    type Moengage struct {

        
            Credentials MoengageCredentials  `json:"credentials"`
            Enabled bool  `json:"enabled"`
         
    }
    
    // MoengageCredentials ...
    type MoengageCredentials struct {

        
            AppID string  `json:"app_id"`
         
    }
    
    // Segment ...
    type Segment struct {

        
            Credentials SegmentCredentials  `json:"credentials"`
            Enabled bool  `json:"enabled"`
         
    }
    
    // SegmentCredentials ...
    type SegmentCredentials struct {

        
            WriteKey string  `json:"write_key"`
         
    }
    
    // Gtm ...
    type Gtm struct {

        
            Credentials GtmCredentials  `json:"credentials"`
            Enabled bool  `json:"enabled"`
         
    }
    
    // GtmCredentials ...
    type GtmCredentials struct {

        
            APIKey string  `json:"api_key"`
         
    }
    
    // Freshchat ...
    type Freshchat struct {

        
            Credentials FreshchatCredentials  `json:"credentials"`
            Enabled bool  `json:"enabled"`
         
    }
    
    // FreshchatCredentials ...
    type FreshchatCredentials struct {

        
            AppID string  `json:"app_id"`
            AppKey string  `json:"app_key"`
            WebToken string  `json:"web_token"`
         
    }
    
    // Safetynet ...
    type Safetynet struct {

        
            Credentials SafetynetCredentials  `json:"credentials"`
            Enabled bool  `json:"enabled"`
         
    }
    
    // SafetynetCredentials ...
    type SafetynetCredentials struct {

        
            APIKey string  `json:"api_key"`
         
    }
    
    // FyndRewards ...
    type FyndRewards struct {

        
            Credentials FyndRewardsCredentials  `json:"credentials"`
         
    }
    
    // FyndRewardsCredentials ...
    type FyndRewardsCredentials struct {

        
            PublicKey string  `json:"public_key"`
         
    }
    
    // GoogleMap ...
    type GoogleMap struct {

        
            Credentials GoogleMapCredentials  `json:"credentials"`
         
    }
    
    // GoogleMapCredentials ...
    type GoogleMapCredentials struct {

        
            APIKey string  `json:"api_key"`
         
    }
    
    // RewardPointsConfig ...
    type RewardPointsConfig struct {

        
            Credit Credit  `json:"credit"`
            Debit Debit  `json:"debit"`
         
    }
    
    // Credit ...
    type Credit struct {

        
            Enabled bool  `json:"enabled"`
         
    }
    
    // Debit ...
    type Debit struct {

        
            Enabled bool  `json:"enabled"`
            AutoApply bool  `json:"auto_apply"`
            StrategyChannel string  `json:"strategy_channel"`
         
    }
    
    // ProductDetailFeature ...
    type ProductDetailFeature struct {

        
            Similar []string  `json:"similar"`
            SellerSelection bool  `json:"seller_selection"`
            UpdateProductMeta bool  `json:"update_product_meta"`
            RequestProduct bool  `json:"request_product"`
         
    }
    
    // LaunchPage ...
    type LaunchPage struct {

        
            PageType string  `json:"page_type"`
            Params map[string]interface{}  `json:"params"`
            Query map[string]interface{}  `json:"query"`
         
    }
    
    // LandingPageFeature ...
    type LandingPageFeature struct {

        
            LaunchPage LaunchPage  `json:"launch_page"`
            ContinueAsGuest bool  `json:"continue_as_guest"`
            LoginBtnText string  `json:"login_btn_text"`
            ShowDomainTextbox bool  `json:"show_domain_textbox"`
            ShowRegisterBtn bool  `json:"show_register_btn"`
         
    }
    
    // RegistrationPageFeature ...
    type RegistrationPageFeature struct {

        
            AskStoreAddress bool  `json:"ask_store_address"`
         
    }
    
    // AppFeature ...
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
    
    // HomePageFeature ...
    type HomePageFeature struct {

        
            OrderProcessing bool  `json:"order_processing"`
         
    }
    
    // CommonFeature ...
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
    
    // CommunicationOptinDialogFeature ...
    type CommunicationOptinDialogFeature struct {

        
            Visibility bool  `json:"visibility"`
         
    }
    
    // DeploymentStoreSelectionFeature ...
    type DeploymentStoreSelectionFeature struct {

        
            Enabled bool  `json:"enabled"`
            Type string  `json:"type"`
         
    }
    
    // ListingPriceFeature ...
    type ListingPriceFeature struct {

        
            Value string  `json:"value"`
            Sort string  `json:"sort"`
         
    }
    
    // CurrencyFeature ...
    type CurrencyFeature struct {

        
            Value []string  `json:"value"`
            Type string  `json:"type"`
            DefaultCurrency string  `json:"default_currency"`
         
    }
    
    // RevenueEngineFeature ...
    type RevenueEngineFeature struct {

        
            Enabled bool  `json:"enabled"`
         
    }
    
    // FeedbackFeature ...
    type FeedbackFeature struct {

        
            Enabled bool  `json:"enabled"`
         
    }
    
    // CompareProductsFeature ...
    type CompareProductsFeature struct {

        
            Enabled bool  `json:"enabled"`
         
    }
    
    // CartFeature ...
    type CartFeature struct {

        
            GstInput bool  `json:"gst_input"`
            StaffSelection bool  `json:"staff_selection"`
            PlacingForCustomer bool  `json:"placing_for_customer"`
            GoogleMap bool  `json:"google_map"`
            RevenueEngineCoupon bool  `json:"revenue_engine_coupon"`
         
    }
    
    // QrFeature ...
    type QrFeature struct {

        
            Application bool  `json:"application"`
            Products bool  `json:"products"`
            Collections bool  `json:"collections"`
         
    }
    
    // PcrFeature ...
    type PcrFeature struct {

        
            StaffSelection bool  `json:"staff_selection"`
         
    }
    
    // OrderFeature ...
    type OrderFeature struct {

        
            BuyAgain bool  `json:"buy_again"`
         
    }
    
    // AppFeatureRequest ...
    type AppFeatureRequest struct {

        
            Feature AppFeature  `json:"feature"`
         
    }
    
    // AppFeatureResponse ...
    type AppFeatureResponse struct {

        
            Feature AppFeature  `json:"feature"`
         
    }
    
    // Currency ...
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
    
    // Domain ...
    type Domain struct {

        
            Verified bool  `json:"verified"`
            IsPrimary bool  `json:"is_primary"`
            IsShortlink bool  `json:"is_shortlink"`
            ID string  `json:"_id"`
            Name string  `json:"name"`
            IsPredefined bool  `json:"is_predefined"`
         
    }
    
    // ApplicationWebsite ...
    type ApplicationWebsite struct {

        
            Enabled bool  `json:"enabled"`
            Basepath string  `json:"basepath"`
         
    }
    
    // ApplicationCors ...
    type ApplicationCors struct {

        
            Domains []string  `json:"domains"`
         
    }
    
    // ApplicationAuth ...
    type ApplicationAuth struct {

        
            Enabled bool  `json:"enabled"`
         
    }
    
    // ApplicationRedirections ...
    type ApplicationRedirections struct {

        
            RedirectFrom string  `json:"redirect_from"`
            RedirectTo string  `json:"redirect_to"`
            Type string  `json:"type"`
         
    }
    
    // ApplicationMeta ...
    type ApplicationMeta struct {

        
            Name string  `json:"name"`
            Value string  `json:"value"`
         
    }
    
    // SecureUrl ...
    type SecureUrl struct {

        
            SecureURL string  `json:"secure_url"`
         
    }
    
    // Application ...
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
            Slug string  `json:"slug"`
         
    }
    
    // NotFound ...
    type NotFound struct {

        
            Message string  `json:"message"`
         
    }
    
    // UnhandledError ...
    type UnhandledError struct {

        
            Message string  `json:"message"`
         
    }
    
    // InvalidPayloadRequest ...
    type InvalidPayloadRequest struct {

        
            Message string  `json:"message"`
         
    }
    
    // SuccessMessageResponse ...
    type SuccessMessageResponse struct {

        
            Message string  `json:"message"`
         
    }
    
    // InventoryBrandRule ...
    type InventoryBrandRule struct {

        
            Criteria string  `json:"criteria"`
            Brands []float64  `json:"brands"`
         
    }
    
    // StoreCriteriaRule ...
    type StoreCriteriaRule struct {

        
            Companies []float64  `json:"companies"`
            Brands []float64  `json:"brands"`
         
    }
    
    // InventoryStoreRule ...
    type InventoryStoreRule struct {

        
            Criteria string  `json:"criteria"`
            Rules []StoreCriteriaRule  `json:"rules"`
            Stores []float64  `json:"stores"`
         
    }
    
    // InventoryPaymentConfig ...
    type InventoryPaymentConfig struct {

        
            ModeOfPayment string  `json:"mode_of_payment"`
            Source string  `json:"source"`
         
    }
    
    // StorePriorityRule ...
    type StorePriorityRule struct {

        
            Enabled bool  `json:"enabled"`
            StoretypeOrder []string  `json:"storetype_order"`
         
    }
    
    // ArticleAssignmentRule ...
    type ArticleAssignmentRule struct {

        
            StorePriority StorePriorityRule  `json:"store_priority"`
         
    }
    
    // InventoryArticleAssignment ...
    type InventoryArticleAssignment struct {

        
            PostOrderReassignment bool  `json:"post_order_reassignment"`
            Rules ArticleAssignmentRule  `json:"rules"`
         
    }
    
    // CompanyAboutAddress ...
    type CompanyAboutAddress struct {

        
            Pincode float64  `json:"pincode"`
            Address1 string  `json:"address1"`
            Address2 string  `json:"address2"`
            City string  `json:"city"`
            State string  `json:"state"`
            Country string  `json:"country"`
            AddressType string  `json:"address_type"`
         
    }
    
    // UserEmail ...
    type UserEmail struct {

        
            Active bool  `json:"active"`
            Primary bool  `json:"primary"`
            Verified bool  `json:"verified"`
            Email string  `json:"email"`
         
    }
    
    // UserPhoneNumber ...
    type UserPhoneNumber struct {

        
            Active bool  `json:"active"`
            Primary bool  `json:"primary"`
            Verified bool  `json:"verified"`
            CountryCode float64  `json:"country_code"`
            Phone string  `json:"phone"`
         
    }
    
    // Page ...
    type Page struct {

        
            ItemTotal float64  `json:"item_total"`
            NextID string  `json:"next_id"`
            HasPrevious bool  `json:"has_previous"`
            HasNext bool  `json:"has_next"`
            Current float64  `json:"current"`
            Type string  `json:"type"`
            Size float64  `json:"size"`
         
    }
    
    // ApplicationInformation ...
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
    
    // InformationAddress ...
    type InformationAddress struct {

        
            Loc string  `json:"loc"`
            AddressLine []string  `json:"address_line"`
            Phone []InformationPhone  `json:"phone"`
            City string  `json:"city"`
            Country string  `json:"country"`
            Pincode float64  `json:"pincode"`
         
    }
    
    // InformationPhone ...
    type InformationPhone struct {

        
            Code string  `json:"code"`
            Number string  `json:"number"`
         
    }
    
    // InformationSupport ...
    type InformationSupport struct {

        
            Phone []string  `json:"phone"`
            Email []string  `json:"email"`
            Timing string  `json:"timing"`
         
    }
    
    // SocialLinks ...
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
    
    // FacebookLink ...
    type FacebookLink struct {

        
            Title string  `json:"title"`
            Icon string  `json:"icon"`
            Link string  `json:"link"`
         
    }
    
    // InstagramLink ...
    type InstagramLink struct {

        
            Title string  `json:"title"`
            Icon string  `json:"icon"`
            Link string  `json:"link"`
         
    }
    
    // TwitterLink ...
    type TwitterLink struct {

        
            Title string  `json:"title"`
            Icon string  `json:"icon"`
            Link string  `json:"link"`
         
    }
    
    // PinterestLink ...
    type PinterestLink struct {

        
            Title string  `json:"title"`
            Icon string  `json:"icon"`
            Link string  `json:"link"`
         
    }
    
    // GooglePlusLink ...
    type GooglePlusLink struct {

        
            Title string  `json:"title"`
            Icon string  `json:"icon"`
            Link string  `json:"link"`
         
    }
    
    // YoutubeLink ...
    type YoutubeLink struct {

        
            Title string  `json:"title"`
            Icon string  `json:"icon"`
            Link string  `json:"link"`
         
    }
    
    // LinkedInLink ...
    type LinkedInLink struct {

        
            Title string  `json:"title"`
            Icon string  `json:"icon"`
            Link string  `json:"link"`
         
    }
    
    // VimeoLink ...
    type VimeoLink struct {

        
            Title string  `json:"title"`
            Icon string  `json:"icon"`
            Link string  `json:"link"`
         
    }
    
    // BlogLink ...
    type BlogLink struct {

        
            Title string  `json:"title"`
            Icon string  `json:"icon"`
            Link string  `json:"link"`
         
    }
    
    // Links ...
    type Links struct {

        
            Title string  `json:"title"`
            Link string  `json:"link"`
         
    }
    
    // BusinessHighlights ...
    type BusinessHighlights struct {

        
            ID string  `json:"_id"`
            Title string  `json:"title"`
            Icon string  `json:"icon"`
            SubTitle string  `json:"sub_title"`
         
    }
    
    // ApplicationDetail ...
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
            Slug string  `json:"slug"`
         
    }
    
    // CurrenciesResponse ...
    type CurrenciesResponse struct {

        
            Items []Currency  `json:"items"`
         
    }
    
    // DefaultCurrency ...
    type DefaultCurrency struct {

        
            Ref string  `json:"ref"`
            Code string  `json:"code"`
         
    }
    
    // AppCurrencyResponse ...
    type AppCurrencyResponse struct {

        
            Application string  `json:"application"`
            DefaultCurrency DefaultCurrency  `json:"default_currency"`
            SupportedCurrency []Currency  `json:"supported_currency"`
         
    }
    
    // StoreLatLong ...
    type StoreLatLong struct {

        
            Type string  `json:"type"`
            Coordinates []float64  `json:"coordinates"`
         
    }
    
    // OptedStoreAddress ...
    type OptedStoreAddress struct {

        
            State string  `json:"state"`
            Address1 string  `json:"address1"`
            LatLong StoreLatLong  `json:"lat_long"`
            Address2 string  `json:"address2"`
            Pincode float64  `json:"pincode"`
            Country string  `json:"country"`
            City string  `json:"city"`
         
    }
    
    // OrderingStore ...
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
    
    // OrderingStores ...
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
    
    // OrderingStoresResponse ...
    type OrderingStoresResponse struct {

        
            Page Page  `json:"page"`
            Items []OrderingStore  `json:"items"`
         
    }
    

    
    // AggregatorConfigDetail ...
    type AggregatorConfigDetail struct {

        
            UserID string  `json:"user_id"`
            MerchantKey string  `json:"merchant_key"`
            API string  `json:"api"`
            Pin string  `json:"pin"`
            Key string  `json:"key"`
            ConfigType string  `json:"config_type"`
            VerifyAPI string  `json:"verify_api"`
            Sdk bool  `json:"sdk"`
            Secret string  `json:"secret"`
            MerchantID string  `json:"merchant_id"`
         
    }
    
    // AggregatorsConfigDetailResponse ...
    type AggregatorsConfigDetailResponse struct {

        
            Juspay AggregatorConfigDetail  `json:"juspay"`
            Simpl AggregatorConfigDetail  `json:"simpl"`
            Rupifi AggregatorConfigDetail  `json:"rupifi"`
            Payumoney AggregatorConfigDetail  `json:"payumoney"`
            Razorpay AggregatorConfigDetail  `json:"razorpay"`
            Ccavenue AggregatorConfigDetail  `json:"ccavenue"`
            Mswipe AggregatorConfigDetail  `json:"mswipe"`
            Stripe AggregatorConfigDetail  `json:"stripe"`
            Env string  `json:"env"`
            Success bool  `json:"success"`
         
    }
    
    // ErrorCodeAndDescription ...
    type ErrorCodeAndDescription struct {

        
            Description string  `json:"description"`
            Code string  `json:"code"`
         
    }
    
    // HttpErrorCodeAndResponse ...
    type HttpErrorCodeAndResponse struct {

        
            Error ErrorCodeAndDescription  `json:"error"`
            Success bool  `json:"success"`
         
    }
    
    // AttachCardRequest ...
    type AttachCardRequest struct {

        
            Nickname string  `json:"nickname"`
            Refresh bool  `json:"refresh"`
            NameOnCard string  `json:"name_on_card"`
            CardID string  `json:"card_id"`
         
    }
    
    // AttachCardsResponse ...
    type AttachCardsResponse struct {

        
            Data map[string]interface{}  `json:"data"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // CardPaymentGateway ...
    type CardPaymentGateway struct {

        
            API string  `json:"api"`
            CustomerID string  `json:"customer_id"`
            Aggregator string  `json:"aggregator"`
         
    }
    
    // ActiveCardPaymentGatewayResponse ...
    type ActiveCardPaymentGatewayResponse struct {

        
            Cards CardPaymentGateway  `json:"cards"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // Card ...
    type Card struct {

        
            CardIsin string  `json:"card_isin"`
            CardToken string  `json:"card_token"`
            CardFingerprint string  `json:"card_fingerprint"`
            AggregatorName string  `json:"aggregator_name"`
            CardType string  `json:"card_type"`
            Nickname string  `json:"nickname"`
            CardBrandImage string  `json:"card_brand_image"`
            Expired bool  `json:"expired"`
            CardName string  `json:"card_name"`
            ExpYear float64  `json:"exp_year"`
            CardNumber string  `json:"card_number"`
            ExpMonth float64  `json:"exp_month"`
            CardBrand string  `json:"card_brand"`
            CardID string  `json:"card_id"`
            CardReference string  `json:"card_reference"`
            CompliantWithTokenisationGuidelines bool  `json:"compliant_with_tokenisation_guidelines"`
            CardIssuer string  `json:"card_issuer"`
         
    }
    
    // ListCardsResponse ...
    type ListCardsResponse struct {

        
            Data []Card  `json:"data"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // DeletehCardRequest ...
    type DeletehCardRequest struct {

        
            CardID string  `json:"card_id"`
         
    }
    
    // DeleteCardsResponse ...
    type DeleteCardsResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // ValidateCustomerRequest ...
    type ValidateCustomerRequest struct {

        
            OrderItems []map[string]interface{}  `json:"order_items"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
            Aggregator string  `json:"aggregator"`
            Payload string  `json:"payload"`
            TransactionAmountInPaise float64  `json:"transaction_amount_in_paise"`
            PhoneNumber string  `json:"phone_number"`
            MerchantParams map[string]interface{}  `json:"merchant_params"`
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
         
    }
    
    // ValidateCustomerResponse ...
    type ValidateCustomerResponse struct {

        
            Data map[string]interface{}  `json:"data"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // ChargeCustomerRequest ...
    type ChargeCustomerRequest struct {

        
            Aggregator string  `json:"aggregator"`
            TransactionToken string  `json:"transaction_token"`
            Amount float64  `json:"amount"`
            OrderID string  `json:"order_id"`
            Verified bool  `json:"verified"`
         
    }
    
    // ChargeCustomerResponse ...
    type ChargeCustomerResponse struct {

        
            DeliveryAddressID string  `json:"delivery_address_id"`
            Aggregator string  `json:"aggregator"`
            Message string  `json:"message"`
            CartID string  `json:"cart_id"`
            OrderID string  `json:"order_id"`
            Status string  `json:"status"`
            Success bool  `json:"success"`
         
    }
    
    // PaymentInitializationRequest ...
    type PaymentInitializationRequest struct {

        
            Contact string  `json:"contact"`
            Method string  `json:"method"`
            DeviceID string  `json:"device_id"`
            Email string  `json:"email"`
            CustomerID string  `json:"customer_id"`
            Aggregator string  `json:"aggregator"`
            Vpa string  `json:"vpa"`
            Amount float64  `json:"amount"`
            Timeout float64  `json:"timeout"`
            MerchantOrderID string  `json:"merchant_order_id"`
            OrderID string  `json:"order_id"`
            Currency string  `json:"currency"`
            RazorpayPaymentID string  `json:"razorpay_payment_id"`
         
    }
    
    // PaymentInitializationResponse ...
    type PaymentInitializationResponse struct {

        
            Method string  `json:"method"`
            DeviceID string  `json:"device_id"`
            AggregatorOrderID string  `json:"aggregator_order_id"`
            Aggregator string  `json:"aggregator"`
            CustomerID string  `json:"customer_id"`
            BqrImage string  `json:"bqr_image"`
            Vpa string  `json:"vpa"`
            UpiPollURL string  `json:"upi_poll_url"`
            VirtualID string  `json:"virtual_id"`
            MerchantOrderID string  `json:"merchant_order_id"`
            Amount float64  `json:"amount"`
            PollingURL string  `json:"polling_url"`
            Timeout float64  `json:"timeout"`
            Currency string  `json:"currency"`
            Status string  `json:"status"`
            Success bool  `json:"success"`
            RazorpayPaymentID string  `json:"razorpay_payment_id"`
         
    }
    
    // PaymentStatusUpdateRequest ...
    type PaymentStatusUpdateRequest struct {

        
            Contact string  `json:"contact"`
            Method string  `json:"method"`
            DeviceID string  `json:"device_id"`
            Email string  `json:"email"`
            CustomerID string  `json:"customer_id"`
            Aggregator string  `json:"aggregator"`
            Vpa string  `json:"vpa"`
            Amount float64  `json:"amount"`
            MerchantOrderID string  `json:"merchant_order_id"`
            OrderID string  `json:"order_id"`
            Currency string  `json:"currency"`
            Status string  `json:"status"`
            MerchantTransactionID string  `json:"merchant_transaction_id"`
         
    }
    
    // PaymentStatusUpdateResponse ...
    type PaymentStatusUpdateResponse struct {

        
            Retry bool  `json:"retry"`
            AggregatorName string  `json:"aggregator_name"`
            RedirectURL string  `json:"redirect_url"`
            Status string  `json:"status"`
            Success bool  `json:"success"`
         
    }
    
    // PaymentModeLogo ...
    type PaymentModeLogo struct {

        
            Large string  `json:"large"`
            Small string  `json:"small"`
         
    }
    
    // IntentApp ...
    type IntentApp struct {

        
            DisplayName string  `json:"display_name"`
            Code string  `json:"code"`
            PackageName string  `json:"package_name"`
            Logos PaymentModeLogo  `json:"logos"`
         
    }
    
    // IntentAppErrorList ...
    type IntentAppErrorList struct {

        
            Code string  `json:"code"`
            PackageName string  `json:"package_name"`
         
    }
    
    // PaymentModeList ...
    type PaymentModeList struct {

        
            DisplayPriority float64  `json:"display_priority"`
            Nickname string  `json:"nickname"`
            FyndVpa string  `json:"fynd_vpa"`
            Expired bool  `json:"expired"`
            DisplayName string  `json:"display_name"`
            ExpMonth float64  `json:"exp_month"`
            CardID string  `json:"card_id"`
            CompliantWithTokenisationGuidelines bool  `json:"compliant_with_tokenisation_guidelines"`
            RetryCount float64  `json:"retry_count"`
            IntentAppErrorList []string  `json:"intent_app_error_list"`
            CardIssuer string  `json:"card_issuer"`
            CardIsin string  `json:"card_isin"`
            CodLimitPerOrder float64  `json:"cod_limit_per_order"`
            CardType string  `json:"card_type"`
            IntentApp []IntentApp  `json:"intent_app"`
            Timeout float64  `json:"timeout"`
            ExpYear float64  `json:"exp_year"`
            CardNumber string  `json:"card_number"`
            CardBrand string  `json:"card_brand"`
            Name string  `json:"name"`
            CardToken string  `json:"card_token"`
            Code string  `json:"code"`
            IntentAppErrorDictList []IntentAppErrorList  `json:"intent_app_error_dict_list"`
            AggregatorName string  `json:"aggregator_name"`
            CardBrandImage string  `json:"card_brand_image"`
            CardReference string  `json:"card_reference"`
            CodLimit float64  `json:"cod_limit"`
            CardFingerprint string  `json:"card_fingerprint"`
            MerchantCode string  `json:"merchant_code"`
            IntentFlow bool  `json:"intent_flow"`
            CardName string  `json:"card_name"`
            LogoURL PaymentModeLogo  `json:"logo_url"`
            RemainingLimit float64  `json:"remaining_limit"`
         
    }
    
    // RootPaymentMode ...
    type RootPaymentMode struct {

        
            List []PaymentModeList  `json:"list"`
            DisplayPriority float64  `json:"display_priority"`
            AggregatorName string  `json:"aggregator_name"`
            AddCardEnabled bool  `json:"add_card_enabled"`
            AnonymousEnable bool  `json:"anonymous_enable"`
            IsPayByCardPl bool  `json:"is_pay_by_card_pl"`
            DisplayName string  `json:"display_name"`
            Name string  `json:"name"`
            SaveCard bool  `json:"save_card"`
         
    }
    
    // AggregatorRoute ...
    type AggregatorRoute struct {

        
            Data map[string]interface{}  `json:"data"`
            APILink string  `json:"api_link"`
            PaymentFlowData string  `json:"payment_flow_data"`
            PaymentFlow string  `json:"payment_flow"`
         
    }
    
    // PaymentFlow ...
    type PaymentFlow struct {

        
            Simpl AggregatorRoute  `json:"simpl"`
            Juspay AggregatorRoute  `json:"juspay"`
            Rupifi AggregatorRoute  `json:"rupifi"`
            UpiRazorpay AggregatorRoute  `json:"upi_razorpay"`
            Razorpay AggregatorRoute  `json:"razorpay"`
            Ccavenue AggregatorRoute  `json:"ccavenue"`
            Epaylater AggregatorRoute  `json:"epaylater"`
            Mswipe AggregatorRoute  `json:"mswipe"`
            Stripe AggregatorRoute  `json:"stripe"`
            Payubiz AggregatorRoute  `json:"payubiz"`
            Fynd AggregatorRoute  `json:"fynd"`
            Ajiodhan AggregatorRoute  `json:"ajiodhan"`
            Jiopay AggregatorRoute  `json:"jiopay"`
            BqrRazorpay AggregatorRoute  `json:"bqr_razorpay"`
         
    }
    
    // PaymentOptionAndFlow ...
    type PaymentOptionAndFlow struct {

        
            PaymentOption []RootPaymentMode  `json:"payment_option"`
            PaymentFlows PaymentFlow  `json:"payment_flows"`
         
    }
    
    // PaymentModeRouteResponse ...
    type PaymentModeRouteResponse struct {

        
            PaymentOptions PaymentOptionAndFlow  `json:"payment_options"`
            Success bool  `json:"success"`
         
    }
    
    // RupifiBannerData ...
    type RupifiBannerData struct {

        
            Status string  `json:"status"`
            KycURL string  `json:"kyc_url"`
         
    }
    
    // RupifiBannerResponse ...
    type RupifiBannerResponse struct {

        
            Data RupifiBannerData  `json:"data"`
            Success bool  `json:"success"`
         
    }
    
    // EpaylaterBannerData ...
    type EpaylaterBannerData struct {

        
            Message string  `json:"message"`
            Display bool  `json:"display"`
            Status string  `json:"status"`
         
    }
    
    // EpaylaterBannerResponse ...
    type EpaylaterBannerResponse struct {

        
            Data EpaylaterBannerData  `json:"data"`
            Success bool  `json:"success"`
         
    }
    
    // ResendOrCancelPaymentRequest ...
    type ResendOrCancelPaymentRequest struct {

        
            DeviceID string  `json:"device_id"`
            OrderID string  `json:"order_id"`
            RequestType string  `json:"request_type"`
         
    }
    
    // LinkStatus ...
    type LinkStatus struct {

        
            Status bool  `json:"status"`
            Message string  `json:"message"`
         
    }
    
    // ResendOrCancelPaymentResponse ...
    type ResendOrCancelPaymentResponse struct {

        
            Data LinkStatus  `json:"data"`
            Success bool  `json:"success"`
         
    }
    
    // renderHTMLRequest ...
    type renderHTMLRequest struct {

        
            Base64Html string  `json:"base64_html"`
            Returntype string  `json:"returntype"`
         
    }
    
    // renderHTMLResponse ...
    type renderHTMLResponse struct {

        
            Html string  `json:"html"`
         
    }
    
    // ValidateVPARequest ...
    type ValidateVPARequest struct {

        
            UpiVpa string  `json:"upi_vpa"`
         
    }
    
    // ValidateUPI ...
    type ValidateUPI struct {

        
            UpiVpa string  `json:"upi_vpa"`
            CustomerName string  `json:"customer_name"`
            IsValid bool  `json:"is_valid"`
            Status string  `json:"status"`
         
    }
    
    // ValidateVPAResponse ...
    type ValidateVPAResponse struct {

        
            Data ValidateUPI  `json:"data"`
            Success bool  `json:"success"`
         
    }
    
    // CardDetails ...
    type CardDetails struct {

        
            CardToken string  `json:"card_token"`
            Status bool  `json:"status"`
            BankCode string  `json:"bank_code"`
            ExtendedCardType string  `json:"extended_card_type"`
            CardExpYear string  `json:"card_exp_year"`
            CardExpMonth string  `json:"card_exp_month"`
            IsDomesticCard bool  `json:"is_domestic_card"`
            User string  `json:"user"`
            CardBrand string  `json:"card_brand"`
            Country string  `json:"country"`
            CardObject string  `json:"card_object"`
            Type string  `json:"type"`
            Bank string  `json:"bank"`
            ID string  `json:"id"`
            NameOnCard string  `json:"name_on_card"`
            CardSubType string  `json:"card_sub_type"`
         
    }
    
    // CardDetailsResponse ...
    type CardDetailsResponse struct {

        
            Data CardDetails  `json:"data"`
            Success bool  `json:"success"`
         
    }
    
    // TransferItemsDetails ...
    type TransferItemsDetails struct {

        
            DisplayName string  `json:"display_name"`
            Name string  `json:"name"`
            LogoLarge string  `json:"logo_large"`
            ID float64  `json:"id"`
            LogoSmall string  `json:"logo_small"`
         
    }
    
    // TransferModeDetails ...
    type TransferModeDetails struct {

        
            DisplayName string  `json:"display_name"`
            Items []TransferItemsDetails  `json:"items"`
         
    }
    
    // TransferModeResponse ...
    type TransferModeResponse struct {

        
            Data []TransferModeDetails  `json:"data"`
         
    }
    
    // UpdateRefundTransferModeRequest ...
    type UpdateRefundTransferModeRequest struct {

        
            Enable bool  `json:"enable"`
            TransferMode string  `json:"transfer_mode"`
         
    }
    
    // UpdateRefundTransferModeResponse ...
    type UpdateRefundTransferModeResponse struct {

        
            Success bool  `json:"success"`
         
    }
    
    // OrderBeneficiaryDetails ...
    type OrderBeneficiaryDetails struct {

        
            Email string  `json:"email"`
            Comment string  `json:"comment"`
            DisplayName string  `json:"display_name"`
            CreatedOn string  `json:"created_on"`
            TransferMode string  `json:"transfer_mode"`
            BeneficiaryID string  `json:"beneficiary_id"`
            Title string  `json:"title"`
            ModifiedOn string  `json:"modified_on"`
            IfscCode string  `json:"ifsc_code"`
            Subtitle string  `json:"subtitle"`
            AccountHolder string  `json:"account_holder"`
            Mobile string  `json:"mobile"`
            ID float64  `json:"id"`
            BranchName string  `json:"branch_name"`
            Address string  `json:"address"`
            BankName string  `json:"bank_name"`
            AccountNo string  `json:"account_no"`
            IsActive bool  `json:"is_active"`
            DelightsUserName string  `json:"delights_user_name"`
         
    }
    
    // OrderBeneficiaryResponse ...
    type OrderBeneficiaryResponse struct {

        
            ShowBeneficiaryDetails bool  `json:"show_beneficiary_details"`
            Beneficiaries []OrderBeneficiaryDetails  `json:"beneficiaries"`
         
    }
    
    // NotFoundResourceError ...
    type NotFoundResourceError struct {

        
            Description string  `json:"description"`
            Code string  `json:"code"`
            Success bool  `json:"success"`
         
    }
    
    // IfscCodeResponse ...
    type IfscCodeResponse struct {

        
            BankName string  `json:"bank_name"`
            BranchName string  `json:"branch_name"`
            Success bool  `json:"success"`
         
    }
    
    // ErrorCodeDescription ...
    type ErrorCodeDescription struct {

        
            Description string  `json:"description"`
            Code string  `json:"code"`
            Success bool  `json:"success"`
         
    }
    
    // AddBeneficiaryViaOtpVerificationRequest ...
    type AddBeneficiaryViaOtpVerificationRequest struct {

        
            Otp string  `json:"otp"`
            RequestID string  `json:"request_id"`
            HashKey string  `json:"hash_key"`
         
    }
    
    // AddBeneficiaryViaOtpVerificationResponse ...
    type AddBeneficiaryViaOtpVerificationResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // WrongOtpError ...
    type WrongOtpError struct {

        
            Description string  `json:"description"`
            Success string  `json:"success"`
            IsVerifiedFlag bool  `json:"is_verified_flag"`
         
    }
    
    // BeneficiaryModeDetails ...
    type BeneficiaryModeDetails struct {

        
            BranchName string  `json:"branch_name"`
            Email string  `json:"email"`
            Comment string  `json:"comment"`
            Address string  `json:"address"`
            Vpa string  `json:"vpa"`
            IfscCode string  `json:"ifsc_code"`
            Wallet string  `json:"wallet"`
            BankName string  `json:"bank_name"`
            AccountNo string  `json:"account_no"`
            AccountHolder string  `json:"account_holder"`
            Mobile string  `json:"mobile"`
         
    }
    
    // AddBeneficiaryDetailsRequest ...
    type AddBeneficiaryDetailsRequest struct {

        
            Otp string  `json:"otp"`
            Details BeneficiaryModeDetails  `json:"details"`
            Delights bool  `json:"delights"`
            OrderID string  `json:"order_id"`
            TransferMode string  `json:"transfer_mode"`
            RequestID string  `json:"request_id"`
            ShipmentID string  `json:"shipment_id"`
         
    }
    
    // RefundAccountResponse ...
    type RefundAccountResponse struct {

        
            IsVerifiedFlag bool  `json:"is_verified_flag"`
            Data map[string]interface{}  `json:"data"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // BankDetailsForOTP ...
    type BankDetailsForOTP struct {

        
            BranchName string  `json:"branch_name"`
            IfscCode string  `json:"ifsc_code"`
            BankName string  `json:"bank_name"`
            AccountNo string  `json:"account_no"`
            AccountHolder string  `json:"account_holder"`
         
    }
    
    // AddBeneficiaryDetailsOTPRequest ...
    type AddBeneficiaryDetailsOTPRequest struct {

        
            Details BankDetailsForOTP  `json:"details"`
            OrderID string  `json:"order_id"`
         
    }
    
    // WalletOtpRequest ...
    type WalletOtpRequest struct {

        
            CountryCode string  `json:"country_code"`
            Mobile string  `json:"mobile"`
         
    }
    
    // WalletOtpResponse ...
    type WalletOtpResponse struct {

        
            RequestID string  `json:"request_id"`
            Success bool  `json:"success"`
            IsVerifiedFlag string  `json:"is_verified_flag"`
         
    }
    
    // SetDefaultBeneficiaryRequest ...
    type SetDefaultBeneficiaryRequest struct {

        
            BeneficiaryID string  `json:"beneficiary_id"`
            OrderID string  `json:"order_id"`
         
    }
    
    // SetDefaultBeneficiaryResponse ...
    type SetDefaultBeneficiaryResponse struct {

        
            IsBeneficiarySet bool  `json:"is_beneficiary_set"`
            Success bool  `json:"success"`
         
    }
    
    // GetPaymentLinkResponse ...
    type GetPaymentLinkResponse struct {

        
            PollingTimeout float64  `json:"polling_timeout"`
            PaymentLinkURL string  `json:"payment_link_url"`
            Message string  `json:"message"`
            PaymentLinkCurrentStatus string  `json:"payment_link_current_status"`
            ExternalOrderID string  `json:"external_order_id"`
            MerchantName string  `json:"merchant_name"`
            Amount float64  `json:"amount"`
            StatusCode float64  `json:"status_code"`
            Success bool  `json:"success"`
         
    }
    
    // ErrorDescription ...
    type ErrorDescription struct {

        
            Msg string  `json:"msg"`
            MerchantName string  `json:"merchant_name"`
            InvalidID bool  `json:"invalid_id"`
            MerchantOrderID string  `json:"merchant_order_id"`
            Amount float64  `json:"amount"`
            Expired bool  `json:"expired"`
            Cancelled bool  `json:"cancelled"`
            PaymentTransactionID string  `json:"payment_transaction_id"`
         
    }
    
    // ErrorResponse ...
    type ErrorResponse struct {

        
            StatusCode float64  `json:"status_code"`
            Error ErrorDescription  `json:"error"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // CreatePaymentLinkMeta ...
    type CreatePaymentLinkMeta struct {

        
            AssignCardID string  `json:"assign_card_id"`
            Pincode string  `json:"pincode"`
            CartID string  `json:"cart_id"`
            CheckoutMode string  `json:"checkout_mode"`
            Amount string  `json:"amount"`
         
    }
    
    // CreatePaymentLinkRequest ...
    type CreatePaymentLinkRequest struct {

        
            Email string  `json:"email"`
            Meta CreatePaymentLinkMeta  `json:"meta"`
            ExternalOrderID string  `json:"external_order_id"`
            Amount float64  `json:"amount"`
            Description string  `json:"description"`
            MobileNumber string  `json:"mobile_number"`
         
    }
    
    // CreatePaymentLinkResponse ...
    type CreatePaymentLinkResponse struct {

        
            PollingTimeout float64  `json:"polling_timeout"`
            PaymentLinkURL string  `json:"payment_link_url"`
            Message string  `json:"message"`
            PaymentLinkID string  `json:"payment_link_id"`
            StatusCode float64  `json:"status_code"`
            Success bool  `json:"success"`
         
    }
    
    // CancelOrResendPaymentLinkRequest ...
    type CancelOrResendPaymentLinkRequest struct {

        
            PaymentLinkID string  `json:"payment_link_id"`
         
    }
    
    // ResendPaymentLinkResponse ...
    type ResendPaymentLinkResponse struct {

        
            PollingTimeout float64  `json:"polling_timeout"`
            StatusCode float64  `json:"status_code"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // CancelPaymentLinkResponse ...
    type CancelPaymentLinkResponse struct {

        
            StatusCode float64  `json:"status_code"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // PollingPaymentLinkResponse ...
    type PollingPaymentLinkResponse struct {

        
            Message string  `json:"message"`
            AggregatorName string  `json:"aggregator_name"`
            RedirectURL string  `json:"redirect_url"`
            Amount float64  `json:"amount"`
            PaymentLinkID string  `json:"payment_link_id"`
            HttpStatus float64  `json:"http_status"`
            OrderID string  `json:"order_id"`
            Status string  `json:"status"`
            StatusCode float64  `json:"status_code"`
            Success bool  `json:"success"`
         
    }
    
    // PaymentMethodsMeta ...
    type PaymentMethodsMeta struct {

        
            MerchantCode string  `json:"merchant_code"`
            PaymentGateway string  `json:"payment_gateway"`
            PaymentIdentifier string  `json:"payment_identifier"`
         
    }
    
    // CreateOrderUserPaymentMethods ...
    type CreateOrderUserPaymentMethods struct {

        
            Mode string  `json:"mode"`
            Meta PaymentMethodsMeta  `json:"meta"`
            Name string  `json:"name"`
         
    }
    
    // CreateOrderUserRequest ...
    type CreateOrderUserRequest struct {

        
            SuccessCallbackURL string  `json:"success_callback_url"`
            FailureCallbackURL string  `json:"failure_callback_url"`
            Meta map[string]interface{}  `json:"meta"`
            PaymentLinkID string  `json:"payment_link_id"`
            Currency string  `json:"currency"`
            PaymentMethods CreateOrderUserPaymentMethods  `json:"payment_methods"`
         
    }
    
    // CreateOrderUserData ...
    type CreateOrderUserData struct {

        
            Contact string  `json:"contact"`
            Method string  `json:"method"`
            Email string  `json:"email"`
            CustomerID string  `json:"customer_id"`
            Aggregator string  `json:"aggregator"`
            Amount float64  `json:"amount"`
            MerchantOrderID string  `json:"merchant_order_id"`
            OrderID string  `json:"order_id"`
            Currency string  `json:"currency"`
            CallbackURL string  `json:"callback_url"`
         
    }
    
    // CreateOrderUserResponse ...
    type CreateOrderUserResponse struct {

        
            PaymentConfirmURL string  `json:"payment_confirm_url"`
            Message string  `json:"message"`
            Data CreateOrderUserData  `json:"data"`
            OrderID string  `json:"order_id"`
            StatusCode float64  `json:"status_code"`
            CallbackURL string  `json:"callback_url"`
            Success bool  `json:"success"`
         
    }
    
    // BalanceDetails ...
    type BalanceDetails struct {

        
            Value float64  `json:"value"`
            FormattedValue string  `json:"formatted_value"`
            Currency string  `json:"currency"`
         
    }
    
    // CreditSummary ...
    type CreditSummary struct {

        
            BuyerStatus string  `json:"buyer_status"`
            Balance BalanceDetails  `json:"balance"`
            MerchantCustomerRefID string  `json:"merchant_customer_ref_id"`
            Status string  `json:"status"`
            AmountAvailable BalanceDetails  `json:"amount_available"`
            CreditLineID string  `json:"credit_line_id"`
            StatusMessage string  `json:"status_message"`
         
    }
    
    // CustomerCreditSummaryResponse ...
    type CustomerCreditSummaryResponse struct {

        
            Data CreditSummary  `json:"data"`
            Success bool  `json:"success"`
         
    }
    
    // RedirectURL ...
    type RedirectURL struct {

        
            SignupURL string  `json:"signup_url"`
            Status bool  `json:"status"`
         
    }
    
    // RedirectToAggregatorResponse ...
    type RedirectToAggregatorResponse struct {

        
            Data RedirectURL  `json:"data"`
            Success bool  `json:"success"`
         
    }
    
    // CreditDetail ...
    type CreditDetail struct {

        
            SignupURL string  `json:"signup_url"`
            IsRegistered bool  `json:"is_registered"`
            Status bool  `json:"status"`
         
    }
    
    // CheckCreditResponse ...
    type CheckCreditResponse struct {

        
            Data CreditDetail  `json:"data"`
            Success bool  `json:"success"`
         
    }
    
    // MarketplaceInfo ...
    type MarketplaceInfo struct {

        
            MembershipID string  `json:"membership_id"`
            DateOfJoining string  `json:"date_of_joining"`
            Name string  `json:"name"`
         
    }
    
    // DeviceDetails ...
    type DeviceDetails struct {

        
            IdentifierType string  `json:"identifier_type"`
            DeviceMake string  `json:"device_make"`
            OsVersion string  `json:"os_version"`
            DeviceType string  `json:"device_type"`
            Os string  `json:"os"`
            DeviceModel string  `json:"device_model"`
            IdentificationNumber string  `json:"identification_number"`
         
    }
    
    // KYCAddress ...
    type KYCAddress struct {

        
            Addressline2 string  `json:"addressline2"`
            Pincode string  `json:"pincode"`
            OwnershipType string  `json:"ownership_type"`
            Addressline1 string  `json:"addressline1"`
            LandMark string  `json:"land_mark"`
            State string  `json:"state"`
            City string  `json:"city"`
         
    }
    
    // UserPersonalInfoInDetails ...
    type UserPersonalInfoInDetails struct {

        
            MobileVerified bool  `json:"mobile_verified"`
            Email string  `json:"email"`
            Gender string  `json:"gender"`
            Dob string  `json:"dob"`
            FathersName string  `json:"fathers_name"`
            VoterID string  `json:"voter_id"`
            Passport string  `json:"passport"`
            Phone string  `json:"phone"`
            DrivingLicense string  `json:"driving_license"`
            EmailVerified bool  `json:"email_verified"`
            FirstName string  `json:"first_name"`
            LastName string  `json:"last_name"`
            AddressAsPerID KYCAddress  `json:"address_as_per_id"`
            MothersName string  `json:"mothers_name"`
            Pan string  `json:"pan"`
            MiddleName string  `json:"middle_name"`
         
    }
    
    // BusinessDetails ...
    type BusinessDetails struct {

        
            Vintage string  `json:"vintage"`
            Gstin string  `json:"gstin"`
            Address KYCAddress  `json:"address"`
            BusinessOwnershipType string  `json:"business_ownership_type"`
            Fssai string  `json:"fssai"`
            BusinessType string  `json:"business_type"`
            Name string  `json:"name"`
            ShopAndEstablishment map[string]interface{}  `json:"shop_and_establishment"`
            Fda string  `json:"fda"`
            Pan string  `json:"pan"`
            EntityType string  `json:"entity_type"`
         
    }
    
    // CustomerOnboardingRequest ...
    type CustomerOnboardingRequest struct {

        
            Mcc string  `json:"mcc"`
            MarketplaceInfo MarketplaceInfo  `json:"marketplace_info"`
            Aggregator string  `json:"aggregator"`
            Device DeviceDetails  `json:"device"`
            PersonalInfo UserPersonalInfoInDetails  `json:"personal_info"`
            Source string  `json:"source"`
            BusinessInfo BusinessDetails  `json:"business_info"`
         
    }
    
    // OnboardSummary ...
    type OnboardSummary struct {

        
            Session map[string]interface{}  `json:"session"`
            RedirectURL string  `json:"redirect_url"`
            Status bool  `json:"status"`
         
    }
    
    // CustomerOnboardingResponse ...
    type CustomerOnboardingResponse struct {

        
            Data OnboardSummary  `json:"data"`
            Success bool  `json:"success"`
         
    }
    
    // OutstandingOrderDetailsResponse ...
    type OutstandingOrderDetailsResponse struct {

        
            StatusCode float64  `json:"status_code"`
            Data []map[string]interface{}  `json:"data"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // PaidOrderDetailsResponse ...
    type PaidOrderDetailsResponse struct {

        
            StatusCode float64  `json:"status_code"`
            Data []map[string]interface{}  `json:"data"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    

    
    // BreakupValues ...
    type BreakupValues struct {

        
            Value float64  `json:"value"`
            CurrencyCode string  `json:"currency_code"`
            Name string  `json:"name"`
            CurrencySymbol string  `json:"currency_symbol"`
            Display string  `json:"display"`
         
    }
    
    // DeliveryAddress ...
    type DeliveryAddress struct {

        
            Country string  `json:"country"`
            State string  `json:"state"`
            Landmark string  `json:"landmark"`
            Email string  `json:"email"`
            Longitude float64  `json:"longitude"`
            CountryPhoneCode string  `json:"country_phone_code"`
            Area string  `json:"area"`
            UpdatedAt string  `json:"updated_at"`
            Version string  `json:"version"`
            Pincode string  `json:"pincode"`
            Name string  `json:"name"`
            AddressType string  `json:"address_type"`
            AddressCategory string  `json:"address_category"`
            CreatedAt string  `json:"created_at"`
            Address string  `json:"address"`
            ContactPerson string  `json:"contact_person"`
            Address1 string  `json:"address1"`
            City string  `json:"city"`
            Latitude float64  `json:"latitude"`
            CountryIsoCode string  `json:"country_iso_code"`
            Address2 string  `json:"address2"`
            Phone string  `json:"phone"`
         
    }
    
    // ShipmentPayment ...
    type ShipmentPayment struct {

        
            Logo string  `json:"logo"`
            Mop string  `json:"mop"`
            Mode string  `json:"mode"`
            Status string  `json:"status"`
            PaymentMode string  `json:"payment_mode"`
            DisplayName string  `json:"display_name"`
         
    }
    
    // FulfillingCompany ...
    type FulfillingCompany struct {

        
            ID float64  `json:"id"`
            Name string  `json:"name"`
         
    }
    
    // Invoice ...
    type Invoice struct {

        
            LabelURL string  `json:"label_url"`
            UpdatedDate string  `json:"updated_date"`
            InvoiceURL string  `json:"invoice_url"`
         
    }
    
    // TimeStampData ...
    type TimeStampData struct {

        
            Max string  `json:"max"`
            Min string  `json:"min"`
         
    }
    
    // Promise ...
    type Promise struct {

        
            ShowPromise bool  `json:"show_promise"`
            Timestamp TimeStampData  `json:"timestamp"`
         
    }
    
    // NestedTrackingDetails ...
    type NestedTrackingDetails struct {

        
            IsCurrent bool  `json:"is_current"`
            Status string  `json:"status"`
            Time string  `json:"time"`
            IsPassed bool  `json:"is_passed"`
         
    }
    
    // TrackingDetails ...
    type TrackingDetails struct {

        
            IsCurrent bool  `json:"is_current"`
            Value string  `json:"value"`
            IsPassed bool  `json:"is_passed"`
            Status string  `json:"status"`
            Time string  `json:"time"`
            TrackingDetails []NestedTrackingDetails  `json:"tracking_details"`
         
    }
    
    // ShipmentTotalDetails ...
    type ShipmentTotalDetails struct {

        
            Pieces float64  `json:"pieces"`
            Sizes float64  `json:"sizes"`
            TotalPrice float64  `json:"total_price"`
         
    }
    
    // AppliedFreeArticles ...
    type AppliedFreeArticles struct {

        
            ArticleID string  `json:"article_id"`
            ParentItemIdentifier string  `json:"parent_item_identifier"`
            FreeGiftItemDetails map[string]interface{}  `json:"free_gift_item_details"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // AppliedPromos ...
    type AppliedPromos struct {

        
            AppliedFreeArticles []AppliedFreeArticles  `json:"applied_free_articles"`
            ArticleQuantity float64  `json:"article_quantity"`
            PromotionType string  `json:"promotion_type"`
            Amount float64  `json:"amount"`
            PromotionName string  `json:"promotion_name"`
            MrpPromotion bool  `json:"mrp_promotion"`
            PromoID string  `json:"promo_id"`
         
    }
    
    // Identifiers ...
    type Identifiers struct {

        
            Ean string  `json:"ean"`
            SkuCode string  `json:"sku_code"`
         
    }
    
    // FinancialBreakup ...
    type FinancialBreakup struct {

        
            ValueOfGood float64  `json:"value_of_good"`
            RefundAmount float64  `json:"refund_amount"`
            Size string  `json:"size"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            RefundCredit float64  `json:"refund_credit"`
            PriceMarked float64  `json:"price_marked"`
            FyndCredits float64  `json:"fynd_credits"`
            Cashback float64  `json:"cashback"`
            Identifiers Identifiers  `json:"identifiers"`
            GstTag string  `json:"gst_tag"`
            CashbackApplied float64  `json:"cashback_applied"`
            Discount float64  `json:"discount"`
            CodCharges float64  `json:"cod_charges"`
            CouponValue float64  `json:"coupon_value"`
            AmountPaid float64  `json:"amount_paid"`
            AddedToFyndCash bool  `json:"added_to_fynd_cash"`
            HsnCode string  `json:"hsn_code"`
            ItemName string  `json:"item_name"`
            PriceEffective float64  `json:"price_effective"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            TotalUnits float64  `json:"total_units"`
            TransferPrice float64  `json:"transfer_price"`
            AmountPaidRoundoff float64  `json:"amount_paid_roundoff"`
            DeliveryCharge float64  `json:"delivery_charge"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            GstFee float64  `json:"gst_fee"`
            PromotionEffectiveDiscount float64  `json:"promotion_effective_discount"`
         
    }
    
    // ItemBrand ...
    type ItemBrand struct {

        
            Name string  `json:"name"`
            Logo string  `json:"logo"`
         
    }
    
    // Item ...
    type Item struct {

        
            L3CategoryName string  `json:"l3_category_name"`
            L1Categories []string  `json:"l1_categories"`
            L2Categories []string  `json:"l2_categories"`
            SellerIdentifier string  `json:"seller_identifier"`
            Brand ItemBrand  `json:"brand"`
            Code string  `json:"code"`
            ID float64  `json:"id"`
            SlugKey string  `json:"slug_key"`
            Name string  `json:"name"`
            Size string  `json:"size"`
            Image []string  `json:"image"`
         
    }
    
    // CurrentStatus ...
    type CurrentStatus struct {

        
            UpdatedAt string  `json:"updated_at"`
            Name string  `json:"name"`
            Status string  `json:"status"`
            JourneyType string  `json:"journey_type"`
         
    }
    
    // Prices ...
    type Prices struct {

        
            CurrencyCode string  `json:"currency_code"`
            RefundAmount float64  `json:"refund_amount"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            RefundCredit float64  `json:"refund_credit"`
            PriceMarked float64  `json:"price_marked"`
            FyndCredits float64  `json:"fynd_credits"`
            Cashback float64  `json:"cashback"`
            CashbackApplied float64  `json:"cashback_applied"`
            Discount float64  `json:"discount"`
            CodCharges float64  `json:"cod_charges"`
            CouponValue float64  `json:"coupon_value"`
            AmountPaid float64  `json:"amount_paid"`
            AddedToFyndCash bool  `json:"added_to_fynd_cash"`
            PromotionEffectiveDiscount float64  `json:"promotion_effective_discount"`
            PriceEffective float64  `json:"price_effective"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            TransferPrice float64  `json:"transfer_price"`
            AmountPaidRoundoff float64  `json:"amount_paid_roundoff"`
            DeliveryCharge float64  `json:"delivery_charge"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            CurrencySymbol string  `json:"currency_symbol"`
            ValueOfGood float64  `json:"value_of_good"`
         
    }
    
    // Bags ...
    type Bags struct {

        
            AppliedPromos []AppliedPromos  `json:"applied_promos"`
            CanReturn bool  `json:"can_return"`
            FinancialBreakup []FinancialBreakup  `json:"financial_breakup"`
            ReturnableDate string  `json:"returnable_date"`
            Quantity float64  `json:"quantity"`
            LineNumber float64  `json:"line_number"`
            Item Item  `json:"item"`
            SellerIdentifier string  `json:"seller_identifier"`
            Meta map[string]interface{}  `json:"meta"`
            CurrencyCode string  `json:"currency_code"`
            ID float64  `json:"id"`
            DeliveryDate string  `json:"delivery_date"`
            ParentPromoBags map[string]interface{}  `json:"parent_promo_bags"`
            CanCancel bool  `json:"can_cancel"`
            CurrencySymbol string  `json:"currency_symbol"`
            CurrentStatus CurrentStatus  `json:"current_status"`
            Prices Prices  `json:"prices"`
         
    }
    
    // ShipmentUserInfo ...
    type ShipmentUserInfo struct {

        
            Email string  `json:"email"`
            Gender string  `json:"gender"`
            Mobile string  `json:"mobile"`
            FirstName string  `json:"first_name"`
            Name string  `json:"name"`
            LastName string  `json:"last_name"`
         
    }
    
    // FulfillingStore ...
    type FulfillingStore struct {

        
            CompanyID float64  `json:"company_id"`
            CompanyName string  `json:"company_name"`
            ID float64  `json:"id"`
            Code string  `json:"code"`
            Name string  `json:"name"`
         
    }
    
    // ShipmentStatus ...
    type ShipmentStatus struct {

        
            Value string  `json:"value"`
            HexCode string  `json:"hex_code"`
            Title string  `json:"title"`
         
    }
    
    // Shipments ...
    type Shipments struct {

        
            DeliveryAddress DeliveryAddress  `json:"delivery_address"`
            ReturnableDate string  `json:"returnable_date"`
            ShipmentID string  `json:"shipment_id"`
            Comment string  `json:"comment"`
            ShowTrackLink bool  `json:"show_track_link"`
            ShowDownloadInvoice bool  `json:"show_download_invoice"`
            TotalBags float64  `json:"total_bags"`
            DeliveryDate string  `json:"delivery_date"`
            ShipmentCreatedAt string  `json:"shipment_created_at"`
            BeneficiaryDetails bool  `json:"beneficiary_details"`
            DpName string  `json:"dp_name"`
            Payment ShipmentPayment  `json:"payment"`
            NeedHelpURL string  `json:"need_help_url"`
            FulfillingCompany FulfillingCompany  `json:"fulfilling_company"`
            TrakingNo string  `json:"traking_no"`
            Invoice Invoice  `json:"invoice"`
            Promise Promise  `json:"promise"`
            SizeInfo map[string]interface{}  `json:"size_info"`
            OrderType string  `json:"order_type"`
            TrackingDetails []TrackingDetails  `json:"tracking_details"`
            TotalDetails ShipmentTotalDetails  `json:"total_details"`
            CustomMeta []map[string]interface{}  `json:"custom_meta"`
            CanReturn bool  `json:"can_return"`
            RefundDetails map[string]interface{}  `json:"refund_details"`
            Bags []Bags  `json:"bags"`
            UserInfo ShipmentUserInfo  `json:"user_info"`
            ReturnMeta map[string]interface{}  `json:"return_meta"`
            OrderID string  `json:"order_id"`
            Prices Prices  `json:"prices"`
            BreakupValues []BreakupValues  `json:"breakup_values"`
            TrackURL string  `json:"track_url"`
            CanBreak map[string]interface{}  `json:"can_break"`
            FulfillingStore FulfillingStore  `json:"fulfilling_store"`
            CanCancel bool  `json:"can_cancel"`
            AwbNo string  `json:"awb_no"`
            ShipmentStatus ShipmentStatus  `json:"shipment_status"`
         
    }
    
    // BagsForReorderArticleAssignment ...
    type BagsForReorderArticleAssignment struct {

        
            Level string  `json:"level"`
            Strategy string  `json:"strategy"`
         
    }
    
    // BagsForReorder ...
    type BagsForReorder struct {

        
            ArticleAssignment BagsForReorderArticleAssignment  `json:"article_assignment"`
            Quantity float64  `json:"quantity"`
            ItemID float64  `json:"item_id"`
            StoreID float64  `json:"store_id"`
            ItemSize string  `json:"item_size"`
            SellerID float64  `json:"seller_id"`
         
    }
    
    // UserInfo ...
    type UserInfo struct {

        
            Email string  `json:"email"`
            Gender string  `json:"gender"`
            Mobile string  `json:"mobile"`
            FirstName string  `json:"first_name"`
            Name string  `json:"name"`
            LastName string  `json:"last_name"`
         
    }
    
    // OrderSchema ...
    type OrderSchema struct {

        
            BreakupValues []BreakupValues  `json:"breakup_values"`
            Shipments []Shipments  `json:"shipments"`
            BagsForReorder []BagsForReorder  `json:"bags_for_reorder"`
            UserInfo UserInfo  `json:"user_info"`
            OrderCreatedTime string  `json:"order_created_time"`
            TotalShipmentsInOrder float64  `json:"total_shipments_in_order"`
            OrderID string  `json:"order_id"`
         
    }
    
    // OrderStatuses ...
    type OrderStatuses struct {

        
            IsSelected bool  `json:"is_selected"`
            Value float64  `json:"value"`
            Display string  `json:"display"`
         
    }
    
    // OrderFilters ...
    type OrderFilters struct {

        
            Statuses []OrderStatuses  `json:"statuses"`
         
    }
    
    // OrderPage ...
    type OrderPage struct {

        
            Type string  `json:"type"`
            HasNext bool  `json:"has_next"`
            ItemTotal float64  `json:"item_total"`
            Current float64  `json:"current"`
            Size float64  `json:"size"`
         
    }
    
    // OrderList ...
    type OrderList struct {

        
            Items []OrderSchema  `json:"items"`
            Filters OrderFilters  `json:"filters"`
            Page OrderPage  `json:"page"`
         
    }
    
    // ApefaceApiError ...
    type ApefaceApiError struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // OrderById ...
    type OrderById struct {

        
            Order OrderSchema  `json:"order"`
         
    }
    
    // ShipmentById ...
    type ShipmentById struct {

        
            Shipment Shipments  `json:"shipment"`
         
    }
    
    // ResponseGetInvoiceShipment ...
    type ResponseGetInvoiceShipment struct {

        
            ShipmentID string  `json:"shipment_id"`
            PresignedURL string  `json:"presigned_url"`
            PresignedType string  `json:"presigned_type"`
            Success bool  `json:"success"`
         
    }
    
    // Track ...
    type Track struct {

        
            Awb string  `json:"awb"`
            UpdatedTime string  `json:"updated_time"`
            ShipmentType string  `json:"shipment_type"`
            AccountName string  `json:"account_name"`
            Status string  `json:"status"`
            LastLocationRecievedAt string  `json:"last_location_recieved_at"`
            UpdatedAt string  `json:"updated_at"`
            Reason string  `json:"reason"`
         
    }
    
    // ShipmentTrack ...
    type ShipmentTrack struct {

        
            Results []Track  `json:"results"`
         
    }
    
    // CustomerDetailsResponse ...
    type CustomerDetailsResponse struct {

        
            ShipmentID string  `json:"shipment_id"`
            Country string  `json:"country"`
            Name string  `json:"name"`
            Phone string  `json:"phone"`
            OrderID string  `json:"order_id"`
         
    }
    
    // SendOtpToCustomerResponse ...
    type SendOtpToCustomerResponse struct {

        
            RequestID string  `json:"request_id"`
            ResendTimer float64  `json:"resend_timer"`
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // VerifyOtp ...
    type VerifyOtp struct {

        
            RequestID string  `json:"request_id"`
            OtpCode string  `json:"otp_code"`
         
    }
    
    // VerifyOtpResponse ...
    type VerifyOtpResponse struct {

        
            Success bool  `json:"success"`
         
    }
    
    // BagReasonMeta ...
    type BagReasonMeta struct {

        
            ShowTextArea bool  `json:"show_text_area"`
         
    }
    
    // QuestionSet ...
    type QuestionSet struct {

        
            ID float64  `json:"id"`
            DisplayName string  `json:"display_name"`
         
    }
    
    // BagReasons ...
    type BagReasons struct {

        
            QcType []string  `json:"qc_type"`
            Meta BagReasonMeta  `json:"meta"`
            Reasons []BagReasons  `json:"reasons"`
            ID float64  `json:"id"`
            QuestionSet []QuestionSet  `json:"question_set"`
            DisplayName string  `json:"display_name"`
         
    }
    
    // ShipmentBagReasons ...
    type ShipmentBagReasons struct {

        
            Reasons []BagReasons  `json:"reasons"`
            Success bool  `json:"success"`
         
    }
    
    // ShipmentReason ...
    type ShipmentReason struct {

        
            FeedbackType string  `json:"feedback_type"`
            Priority float64  `json:"priority"`
            ReasonText string  `json:"reason_text"`
            ReasonID float64  `json:"reason_id"`
            ShowTextArea bool  `json:"show_text_area"`
            Flow string  `json:"flow"`
         
    }
    
    // ShipmentReasons ...
    type ShipmentReasons struct {

        
            Reasons []ShipmentReason  `json:"reasons"`
         
    }
    
    // Products ...
    type Products struct {

        
            Quantity float64  `json:"quantity"`
            Identifier string  `json:"identifier"`
            LineNumber float64  `json:"line_number"`
         
    }
    
    // ProductsDataUpdatesFilters ...
    type ProductsDataUpdatesFilters struct {

        
            Identifier string  `json:"identifier"`
            LineNumber float64  `json:"line_number"`
         
    }
    
    // ProductsDataUpdates ...
    type ProductsDataUpdates struct {

        
            Filters []ProductsDataUpdatesFilters  `json:"filters"`
            Data map[string]interface{}  `json:"data"`
         
    }
    
    // EntitiesDataUpdates ...
    type EntitiesDataUpdates struct {

        
            Filters []map[string]interface{}  `json:"filters"`
            Data map[string]interface{}  `json:"data"`
         
    }
    
    // DataUpdates ...
    type DataUpdates struct {

        
            Products []ProductsDataUpdates  `json:"products"`
            Entities []EntitiesDataUpdates  `json:"entities"`
         
    }
    
    // ProductsReasonsFilters ...
    type ProductsReasonsFilters struct {

        
            Quantity float64  `json:"quantity"`
            Identifier string  `json:"identifier"`
            LineNumber float64  `json:"line_number"`
         
    }
    
    // ProductsReasonsData ...
    type ProductsReasonsData struct {

        
            ReasonID float64  `json:"reason_id"`
            ReasonText string  `json:"reason_text"`
         
    }
    
    // ProductsReasons ...
    type ProductsReasons struct {

        
            Filters []ProductsReasonsFilters  `json:"filters"`
            Data ProductsReasonsData  `json:"data"`
         
    }
    
    // EntityReasonData ...
    type EntityReasonData struct {

        
            ReasonID float64  `json:"reason_id"`
            ReasonText string  `json:"reason_text"`
         
    }
    
    // EntitiesReasons ...
    type EntitiesReasons struct {

        
            Filters []map[string]interface{}  `json:"filters"`
            Data EntityReasonData  `json:"data"`
         
    }
    
    // ReasonsData ...
    type ReasonsData struct {

        
            Products []ProductsReasons  `json:"products"`
            Entities []EntitiesReasons  `json:"entities"`
         
    }
    
    // ShipmentsRequest ...
    type ShipmentsRequest struct {

        
            Products []Products  `json:"products"`
            DataUpdates DataUpdates  `json:"data_updates"`
            Identifier string  `json:"identifier"`
            Reasons ReasonsData  `json:"reasons"`
         
    }
    
    // StatuesRequest ...
    type StatuesRequest struct {

        
            ExcludeBagsNextState string  `json:"exclude_bags_next_state"`
            Shipments []ShipmentsRequest  `json:"shipments"`
            Status string  `json:"status"`
         
    }
    
    // UpdateShipmentStatusRequest ...
    type UpdateShipmentStatusRequest struct {

        
            ForceTransition bool  `json:"force_transition"`
            Task bool  `json:"task"`
            LockAfterTransition bool  `json:"lock_after_transition"`
            UnlockBeforeTransition bool  `json:"unlock_before_transition"`
            Statuses []StatuesRequest  `json:"statuses"`
         
    }
    
    // StatusesBodyResponse ...
    type StatusesBodyResponse struct {

        
            Shipments []map[string]interface{}  `json:"shipments"`
         
    }
    
    // ShipmentApplicationStatusResponse ...
    type ShipmentApplicationStatusResponse struct {

        
            Statuses []StatusesBodyResponse  `json:"statuses"`
         
    }
    
    // ErrorResponse ...
    type ErrorResponse struct {

        
            Exception string  `json:"exception"`
            Message string  `json:"message"`
            Code string  `json:"code"`
            Status float64  `json:"status"`
            StackTrace string  `json:"stack_trace"`
         
    }
    

    
    // RewardsArticle ...
    type RewardsArticle struct {

        
            ID string  `json:"id"`
            Points float64  `json:"points"`
            Price float64  `json:"price"`
         
    }
    
    // CatalogueOrderResponse ...
    type CatalogueOrderResponse struct {

        
            Articles []RewardsArticle  `json:"articles"`
         
    }
    
    // CatalogueOrderRequest ...
    type CatalogueOrderRequest struct {

        
            Articles []RewardsArticle  `json:"articles"`
         
    }
    
    // PointsResponse ...
    type PointsResponse struct {

        
            Points float64  `json:"points"`
         
    }
    
    // ReferralDetailsUser ...
    type ReferralDetailsUser struct {

        
            Blocked bool  `json:"blocked"`
            Points float64  `json:"points"`
            Redeemed bool  `json:"redeemed"`
            ReferralCode string  `json:"referral_code"`
         
    }
    
    // Offer ...
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
    
    // Schedule ...
    type Schedule struct {

        
            Duration float64  `json:"duration"`
            End string  `json:"end"`
            Start string  `json:"start"`
            Cron string  `json:"cron"`
         
    }
    
    // Error ...
    type Error struct {

        
            Code float64  `json:"code"`
            Exception string  `json:"exception"`
            Info string  `json:"info"`
            Message string  `json:"message"`
         
    }
    
    // Asset ...
    type Asset struct {

        
            AspectRatio string  `json:"aspect_ratio"`
            ID string  `json:"id"`
            SecureURL string  `json:"secure_url"`
         
    }
    
    // ShareMessages ...
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
    
    // ReferralDetailsResponse ...
    type ReferralDetailsResponse struct {

        
            Referral Offer  `json:"referral"`
            Share ShareMessages  `json:"share"`
            User ReferralDetailsUser  `json:"user"`
            ReferrerInfo string  `json:"referrer_info"`
            TermsConditionsLink string  `json:"terms_conditions_link"`
         
    }
    
    // OrderDiscountRequest ...
    type OrderDiscountRequest struct {

        
            OrderAmount float64  `json:"order_amount"`
            Currency string  `json:"currency"`
         
    }
    
    // OrderDiscountRuleBucket ...
    type OrderDiscountRuleBucket struct {

        
            High float64  `json:"high"`
            Low float64  `json:"low"`
            Max float64  `json:"max"`
            Value float64  `json:"value"`
            ValueType string  `json:"value_type"`
         
    }
    
    // DiscountProperties ...
    type DiscountProperties struct {

        
            Absolute float64  `json:"absolute"`
            Currency string  `json:"currency"`
            DisplayAbsolute string  `json:"display_absolute"`
            DisplayPercent string  `json:"display_percent"`
            Percent float64  `json:"percent"`
         
    }
    
    // OrderDiscountResponse ...
    type OrderDiscountResponse struct {

        
            OrderAmount float64  `json:"order_amount"`
            Points float64  `json:"points"`
            Discount DiscountProperties  `json:"discount"`
            BaseDiscount DiscountProperties  `json:"base_discount"`
            AppliedRuleBucket OrderDiscountRuleBucket  `json:"applied_rule_bucket"`
         
    }
    
    // RedeemReferralCodeRequest ...
    type RedeemReferralCodeRequest struct {

        
            DeviceID string  `json:"device_id"`
            ReferralCode string  `json:"referral_code"`
         
    }
    
    // RedeemReferralCodeResponse ...
    type RedeemReferralCodeResponse struct {

        
            Redeemed bool  `json:"redeemed"`
            Message string  `json:"message"`
            ReferrerInfo string  `json:"referrer_info"`
            ReferrerID string  `json:"referrer_id"`
            Points float64  `json:"points"`
         
    }
    
    // PointsHistoryResponse ...
    type PointsHistoryResponse struct {

        
            Items []PointsHistory  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // PointsHistory ...
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
    
    // Page ...
    type Page struct {

        
            ItemTotal float64  `json:"item_total"`
            NextID string  `json:"next_id"`
            HasPrevious bool  `json:"has_previous"`
            HasNext bool  `json:"has_next"`
            Current float64  `json:"current"`
            Type string  `json:"type"`
            Size float64  `json:"size"`
         
    }
    

    
    // BasePrice ...
    type BasePrice struct {

        
            CurrencyCode string  `json:"currency_code"`
            CurrencySymbol string  `json:"currency_symbol"`
            Effective float64  `json:"effective"`
            Marked float64  `json:"marked"`
         
    }
    
    // ArticlePriceInfo ...
    type ArticlePriceInfo struct {

        
            Converted BasePrice  `json:"converted"`
            Base BasePrice  `json:"base"`
         
    }
    
    // BaseInfo ...
    type BaseInfo struct {

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // StoreInfo ...
    type StoreInfo struct {

        
            StoreCode string  `json:"store_code"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // ProductArticle ...
    type ProductArticle struct {

        
            Price ArticlePriceInfo  `json:"price"`
            Seller BaseInfo  `json:"seller"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            MtoQuantity float64  `json:"mto_quantity"`
            CartItemMeta map[string]interface{}  `json:"cart_item_meta"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            ProductGroupTags []string  `json:"product_group_tags"`
            Size string  `json:"size"`
            Type string  `json:"type"`
            IsGiftVisible bool  `json:"is_gift_visible"`
            Store StoreInfo  `json:"store"`
            SellerIdentifier string  `json:"seller_identifier"`
            GiftCard map[string]interface{}  `json:"gift_card"`
            Identifier map[string]interface{}  `json:"identifier"`
            UID string  `json:"uid"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // CouponDetails ...
    type CouponDetails struct {

        
            DiscountSingleQuantity float64  `json:"discount_single_quantity"`
            Code string  `json:"code"`
            DiscountTotalQuantity float64  `json:"discount_total_quantity"`
         
    }
    
    // ProductPrice ...
    type ProductPrice struct {

        
            Marked float64  `json:"marked"`
            AddOn float64  `json:"add_on"`
            CurrencySymbol string  `json:"currency_symbol"`
            CurrencyCode string  `json:"currency_code"`
            Effective float64  `json:"effective"`
            Selling float64  `json:"selling"`
         
    }
    
    // ProductPriceInfo ...
    type ProductPriceInfo struct {

        
            Converted ProductPrice  `json:"converted"`
            Base ProductPrice  `json:"base"`
         
    }
    
    // CartProductIdentifer ...
    type CartProductIdentifer struct {

        
            Identifier string  `json:"identifier"`
         
    }
    
    // PromoMeta ...
    type PromoMeta struct {

        
            Message string  `json:"message"`
         
    }
    
    // Tags ...
    type Tags struct {

        
            Tags map[string]interface{}  `json:"tags"`
         
    }
    
    // CategoryInfo ...
    type CategoryInfo struct {

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // NetQuantity ...
    type NetQuantity struct {

        
            Value string  `json:"value"`
            Unit string  `json:"unit"`
         
    }
    
    // ProductImage ...
    type ProductImage struct {

        
            AspectRatio string  `json:"aspect_ratio"`
            URL string  `json:"url"`
            SecureURL string  `json:"secure_url"`
         
    }
    
    // ActionQuery ...
    type ActionQuery struct {

        
            ProductSlug []string  `json:"product_slug"`
         
    }
    
    // ProductAction ...
    type ProductAction struct {

        
            URL string  `json:"url"`
            Query ActionQuery  `json:"query"`
            Type string  `json:"type"`
         
    }
    
    // CartProduct ...
    type CartProduct struct {

        
            TeaserTag Tags  `json:"teaser_tag"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Categories []CategoryInfo  `json:"categories"`
            Tags []string  `json:"tags"`
            NetQuantity NetQuantity  `json:"net_quantity"`
            Images []ProductImage  `json:"images"`
            Type string  `json:"type"`
            Brand BaseInfo  `json:"brand"`
            Action ProductAction  `json:"action"`
            ItemCode string  `json:"item_code"`
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // PromiseFormatted ...
    type PromiseFormatted struct {

        
            Max string  `json:"max"`
            Min string  `json:"min"`
         
    }
    
    // PromiseTimestamp ...
    type PromiseTimestamp struct {

        
            Max float64  `json:"max"`
            Min float64  `json:"min"`
         
    }
    
    // ShipmentPromise ...
    type ShipmentPromise struct {

        
            Formatted PromiseFormatted  `json:"formatted"`
            Timestamp PromiseTimestamp  `json:"timestamp"`
         
    }
    
    // BuyRules ...
    type BuyRules struct {

        
            ItemCriteria map[string]interface{}  `json:"item_criteria"`
            CartConditions map[string]interface{}  `json:"cart_conditions"`
         
    }
    
    // Ownership ...
    type Ownership struct {

        
            PayableCategory string  `json:"payable_category"`
            PayableBy string  `json:"payable_by"`
         
    }
    
    // FreeGiftItem ...
    type FreeGiftItem struct {

        
            ItemName string  `json:"item_name"`
            ItemSlug string  `json:"item_slug"`
            ItemImagesURL []string  `json:"item_images_url"`
            ItemPriceDetails map[string]interface{}  `json:"item_price_details"`
            ItemBrandName string  `json:"item_brand_name"`
            ItemID float64  `json:"item_id"`
         
    }
    
    // AppliedFreeArticles ...
    type AppliedFreeArticles struct {

        
            ParentItemIdentifier string  `json:"parent_item_identifier"`
            Quantity float64  `json:"quantity"`
            ArticleID string  `json:"article_id"`
            FreeGiftItemDetails FreeGiftItem  `json:"free_gift_item_details"`
         
    }
    
    // DiscountRulesApp ...
    type DiscountRulesApp struct {

        
            ItemCriteria map[string]interface{}  `json:"item_criteria"`
            RawOffer map[string]interface{}  `json:"raw_offer"`
            MatchedBuyRules []string  `json:"matched_buy_rules"`
            Offer map[string]interface{}  `json:"offer"`
         
    }
    
    // AppliedPromotion ...
    type AppliedPromotion struct {

        
            PromotionOfferText string  `json:"promotion_offer_text"`
            PromoID string  `json:"promo_id"`
            BuyRules []BuyRules  `json:"buy_rules"`
            Ownership Ownership  `json:"ownership"`
            AppliedFreeArticles []AppliedFreeArticles  `json:"applied_free_articles"`
            MrpPromotion bool  `json:"mrp_promotion"`
            Amount float64  `json:"amount"`
            PromotionName string  `json:"promotion_name"`
            ArticleQuantity float64  `json:"article_quantity"`
            PromotionGroup string  `json:"promotion_group"`
            PromotionType string  `json:"promotion_type"`
            DiscountRules []DiscountRulesApp  `json:"discount_rules"`
         
    }
    
    // ProductAvailabilitySize ...
    type ProductAvailabilitySize struct {

        
            IsAvailable bool  `json:"is_available"`
            Value string  `json:"value"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Meta map[string]interface{}  `json:"meta"`
            Display string  `json:"display"`
         
    }
    
    // ProductAvailability ...
    type ProductAvailability struct {

        
            OtherStoreQuantity float64  `json:"other_store_quantity"`
            OutOfStock bool  `json:"out_of_stock"`
            Deliverable bool  `json:"deliverable"`
            Sizes []string  `json:"sizes"`
            AvailableSizes []ProductAvailabilitySize  `json:"available_sizes"`
            IsValid bool  `json:"is_valid"`
         
    }
    
    // CartProductInfo ...
    type CartProductInfo struct {

        
            IsSet bool  `json:"is_set"`
            BulkOffer map[string]interface{}  `json:"bulk_offer"`
            Article ProductArticle  `json:"article"`
            Coupon CouponDetails  `json:"coupon"`
            PricePerUnit ProductPriceInfo  `json:"price_per_unit"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
            PromoMeta PromoMeta  `json:"promo_meta"`
            Quantity float64  `json:"quantity"`
            Product CartProduct  `json:"product"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            CouponMessage string  `json:"coupon_message"`
            Discount string  `json:"discount"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            PromotionsApplied []AppliedPromotion  `json:"promotions_applied"`
            Availability ProductAvailability  `json:"availability"`
            Price ProductPriceInfo  `json:"price"`
            Moq map[string]interface{}  `json:"moq"`
            Key string  `json:"key"`
            Message string  `json:"message"`
         
    }
    
    // PaymentSelectionLock ...
    type PaymentSelectionLock struct {

        
            PaymentIdentifier string  `json:"payment_identifier"`
            DefaultOptions string  `json:"default_options"`
            Enabled bool  `json:"enabled"`
         
    }
    
    // CartCurrency ...
    type CartCurrency struct {

        
            Symbol string  `json:"symbol"`
            Code string  `json:"code"`
         
    }
    
    // LoyaltyPoints ...
    type LoyaltyPoints struct {

        
            Applicable float64  `json:"applicable"`
            Description string  `json:"description"`
            IsApplied bool  `json:"is_applied"`
            Total float64  `json:"total"`
         
    }
    
    // DisplayBreakup ...
    type DisplayBreakup struct {

        
            Value float64  `json:"value"`
            CurrencySymbol string  `json:"currency_symbol"`
            Key string  `json:"key"`
            CurrencyCode string  `json:"currency_code"`
            Message []string  `json:"message"`
            Display string  `json:"display"`
         
    }
    
    // CouponBreakup ...
    type CouponBreakup struct {

        
            Value float64  `json:"value"`
            Description string  `json:"description"`
            IsApplied bool  `json:"is_applied"`
            CouponValue float64  `json:"coupon_value"`
            CouponType string  `json:"coupon_type"`
            Type string  `json:"type"`
            SubTitle string  `json:"sub_title"`
            MinimumCartValue float64  `json:"minimum_cart_value"`
            Code string  `json:"code"`
            Message string  `json:"message"`
            Title string  `json:"title"`
            MaxDiscountValue float64  `json:"max_discount_value"`
            UID string  `json:"uid"`
         
    }
    
    // RawBreakup ...
    type RawBreakup struct {

        
            Coupon float64  `json:"coupon"`
            GstCharges float64  `json:"gst_charges"`
            FyndCash float64  `json:"fynd_cash"`
            CodCharge float64  `json:"cod_charge"`
            DeliveryCharge float64  `json:"delivery_charge"`
            YouSaved float64  `json:"you_saved"`
            ConvenienceFee float64  `json:"convenience_fee"`
            MrpTotal float64  `json:"mrp_total"`
            Discount float64  `json:"discount"`
            Vog float64  `json:"vog"`
            GiftCard float64  `json:"gift_card"`
            Subtotal float64  `json:"subtotal"`
            Total float64  `json:"total"`
         
    }
    
    // CartBreakup ...
    type CartBreakup struct {

        
            LoyaltyPoints LoyaltyPoints  `json:"loyalty_points"`
            Display []DisplayBreakup  `json:"display"`
            Coupon CouponBreakup  `json:"coupon"`
            Raw RawBreakup  `json:"raw"`
         
    }
    
    // CartDetailResponse ...
    type CartDetailResponse struct {

        
            ID string  `json:"id"`
            PanNo string  `json:"pan_no"`
            LastModified string  `json:"last_modified"`
            Items []CartProductInfo  `json:"items"`
            IsValid bool  `json:"is_valid"`
            CouponText string  `json:"coupon_text"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            Gstin string  `json:"gstin"`
            AppliedPromoDetails []AppliedPromotion  `json:"applied_promo_details"`
            BuyNow bool  `json:"buy_now"`
            Currency CartCurrency  `json:"currency"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            CheckoutMode string  `json:"checkout_mode"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            PanConfig map[string]interface{}  `json:"pan_config"`
            Message string  `json:"message"`
            Comment string  `json:"comment"`
         
    }
    
    // AddProductCart ...
    type AddProductCart struct {

        
            Pos bool  `json:"pos"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ArticleAssignment map[string]interface{}  `json:"article_assignment"`
            ParentItemIdentifiers []map[string]string  `json:"parent_item_identifiers"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            ProductGroupTags []string  `json:"product_group_tags"`
            SellerID float64  `json:"seller_id"`
            Meta map[string]interface{}  `json:"meta"`
            Quantity float64  `json:"quantity"`
            StoreID float64  `json:"store_id"`
            ItemSize string  `json:"item_size"`
            ArticleID string  `json:"article_id"`
            Display string  `json:"display"`
            ItemID float64  `json:"item_id"`
         
    }
    
    // AddCartRequest ...
    type AddCartRequest struct {

        
            NewCart bool  `json:"new_cart"`
            Items []AddProductCart  `json:"items"`
         
    }
    
    // AddCartDetailResponse ...
    type AddCartDetailResponse struct {

        
            Cart CartDetailResponse  `json:"cart"`
            Partial bool  `json:"partial"`
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // UpdateProductCart ...
    type UpdateProductCart struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            ItemID float64  `json:"item_id"`
            ItemIndex float64  `json:"item_index"`
            Meta map[string]interface{}  `json:"meta"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
            ItemSize string  `json:"item_size"`
            ArticleID string  `json:"article_id"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // UpdateCartRequest ...
    type UpdateCartRequest struct {

        
            Operation string  `json:"operation"`
            Items []UpdateProductCart  `json:"items"`
         
    }
    
    // UpdateCartDetailResponse ...
    type UpdateCartDetailResponse struct {

        
            Cart CartDetailResponse  `json:"cart"`
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // CartItemCountResponse ...
    type CartItemCountResponse struct {

        
            UserCartItemsCount float64  `json:"user_cart_items_count"`
         
    }
    
    // Coupon ...
    type Coupon struct {

        
            ExpiresOn string  `json:"expires_on"`
            Description string  `json:"description"`
            IsApplied bool  `json:"is_applied"`
            CouponValue float64  `json:"coupon_value"`
            CouponType string  `json:"coupon_type"`
            SubTitle string  `json:"sub_title"`
            MinimumCartValue float64  `json:"minimum_cart_value"`
            CouponCode string  `json:"coupon_code"`
            Message string  `json:"message"`
            Title string  `json:"title"`
            MaxDiscountValue float64  `json:"max_discount_value"`
            IsApplicable bool  `json:"is_applicable"`
         
    }
    
    // PageCoupon ...
    type PageCoupon struct {

        
            TotalItemCount float64  `json:"total_item_count"`
            HasPrevious bool  `json:"has_previous"`
            HasNext bool  `json:"has_next"`
            Current float64  `json:"current"`
            Total float64  `json:"total"`
         
    }
    
    // GetCouponResponse ...
    type GetCouponResponse struct {

        
            AvailableCouponList []Coupon  `json:"available_coupon_list"`
            Page PageCoupon  `json:"page"`
         
    }
    
    // ApplyCouponRequest ...
    type ApplyCouponRequest struct {

        
            CouponCode string  `json:"coupon_code"`
         
    }
    
    // OfferSeller ...
    type OfferSeller struct {

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // OfferPrice ...
    type OfferPrice struct {

        
            Marked float64  `json:"marked"`
            CurrencySymbol string  `json:"currency_symbol"`
            CurrencyCode string  `json:"currency_code"`
            BulkEffective float64  `json:"bulk_effective"`
            Effective float64  `json:"effective"`
         
    }
    
    // OfferItem ...
    type OfferItem struct {

        
            Price OfferPrice  `json:"price"`
            Margin float64  `json:"margin"`
            Type string  `json:"type"`
            AutoApplied bool  `json:"auto_applied"`
            Best bool  `json:"best"`
            Total float64  `json:"total"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // BulkPriceOffer ...
    type BulkPriceOffer struct {

        
            Seller OfferSeller  `json:"seller"`
            Offers []OfferItem  `json:"offers"`
         
    }
    
    // BulkPriceResponse ...
    type BulkPriceResponse struct {

        
            Data []BulkPriceOffer  `json:"data"`
         
    }
    
    // RewardPointRequest ...
    type RewardPointRequest struct {

        
            Points bool  `json:"points"`
         
    }
    
    // GeoLocation ...
    type GeoLocation struct {

        
            Latitude float64  `json:"latitude"`
            Longitude float64  `json:"longitude"`
         
    }
    
    // Address ...
    type Address struct {

        
            ID string  `json:"id"`
            State string  `json:"state"`
            GoogleMapPoint map[string]interface{}  `json:"google_map_point"`
            Tags []string  `json:"tags"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            AddressType string  `json:"address_type"`
            Phone string  `json:"phone"`
            IsActive bool  `json:"is_active"`
            AreaCodeSlug string  `json:"area_code_slug"`
            CountryPhoneCode string  `json:"country_phone_code"`
            Area string  `json:"area"`
            Email string  `json:"email"`
            GeoLocation GeoLocation  `json:"geo_location"`
            Landmark string  `json:"landmark"`
            CheckoutMode string  `json:"checkout_mode"`
            CountryIsoCode string  `json:"country_iso_code"`
            Meta map[string]interface{}  `json:"meta"`
            CountryCode string  `json:"country_code"`
            City string  `json:"city"`
            Country string  `json:"country"`
            IsDefaultAddress bool  `json:"is_default_address"`
            Address string  `json:"address"`
            UserID string  `json:"user_id"`
            CreatedByUserID string  `json:"created_by_user_id"`
            AreaCode string  `json:"area_code"`
            Name string  `json:"name"`
         
    }
    
    // GetAddressesResponse ...
    type GetAddressesResponse struct {

        
            Address []Address  `json:"address"`
            PiiMasking bool  `json:"pii_masking"`
         
    }
    
    // SaveAddressResponse ...
    type SaveAddressResponse struct {

        
            Success bool  `json:"success"`
            ID string  `json:"id"`
            IsDefaultAddress bool  `json:"is_default_address"`
         
    }
    
    // UpdateAddressResponse ...
    type UpdateAddressResponse struct {

        
            Success bool  `json:"success"`
            ID string  `json:"id"`
            IsUpdated bool  `json:"is_updated"`
            IsDefaultAddress bool  `json:"is_default_address"`
         
    }
    
    // DeleteAddressResponse ...
    type DeleteAddressResponse struct {

        
            ID string  `json:"id"`
            IsDeleted bool  `json:"is_deleted"`
         
    }
    
    // SelectCartAddressRequest ...
    type SelectCartAddressRequest struct {

        
            ID string  `json:"id"`
            BillingAddressID string  `json:"billing_address_id"`
            CartID string  `json:"cart_id"`
         
    }
    
    // UpdateCartPaymentRequest ...
    type UpdateCartPaymentRequest struct {

        
            AddressID string  `json:"address_id"`
            ID string  `json:"id"`
            AggregatorName string  `json:"aggregator_name"`
            MerchantCode string  `json:"merchant_code"`
            PaymentMode string  `json:"payment_mode"`
            PaymentIdentifier string  `json:"payment_identifier"`
         
    }
    
    // CouponValidity ...
    type CouponValidity struct {

        
            DisplayMessageEn string  `json:"display_message_en"`
            Valid bool  `json:"valid"`
            Code string  `json:"code"`
            Title string  `json:"title"`
            Discount float64  `json:"discount"`
            NextValidationRequired bool  `json:"next_validation_required"`
         
    }
    
    // PaymentCouponValidate ...
    type PaymentCouponValidate struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
            CouponValidity CouponValidity  `json:"coupon_validity"`
         
    }
    
    // ShipmentResponse ...
    type ShipmentResponse struct {

        
            FulfillmentType string  `json:"fulfillment_type"`
            ShipmentType string  `json:"shipment_type"`
            Promise ShipmentPromise  `json:"promise"`
            BoxType string  `json:"box_type"`
            FulfillmentID float64  `json:"fulfillment_id"`
            Items []CartProductInfo  `json:"items"`
            OrderType string  `json:"order_type"`
            Shipments float64  `json:"shipments"`
            DpOptions map[string]interface{}  `json:"dp_options"`
            DpID string  `json:"dp_id"`
         
    }
    
    // CartShipmentsResponse ...
    type CartShipmentsResponse struct {

        
            ID string  `json:"id"`
            Error bool  `json:"error"`
            BuyNow bool  `json:"buy_now"`
            LastModified string  `json:"last_modified"`
            Currency CartCurrency  `json:"currency"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            CheckoutMode string  `json:"checkout_mode"`
            Shipments []ShipmentResponse  `json:"shipments"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            IsValid bool  `json:"is_valid"`
            CouponText string  `json:"coupon_text"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            Message string  `json:"message"`
            CartID float64  `json:"cart_id"`
            Comment string  `json:"comment"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            Gstin string  `json:"gstin"`
            UID string  `json:"uid"`
         
    }
    
    // UpdateCartShipmentItem ...
    type UpdateCartShipmentItem struct {

        
            ShipmentType string  `json:"shipment_type"`
            ArticleUID string  `json:"article_uid"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // UpdateCartShipmentRequest ...
    type UpdateCartShipmentRequest struct {

        
            Shipments []UpdateCartShipmentItem  `json:"shipments"`
         
    }
    
    // CustomerDetails ...
    type CustomerDetails struct {

        
            Email string  `json:"email"`
            Mobile string  `json:"mobile"`
            Name string  `json:"name"`
         
    }
    
    // StaffCheckout ...
    type StaffCheckout struct {

        
            ID string  `json:"_id"`
            FirstName string  `json:"first_name"`
            LastName string  `json:"last_name"`
            User string  `json:"user"`
            EmployeeCode string  `json:"employee_code"`
         
    }
    
    // Files ...
    type Files struct {

        
            Key string  `json:"key"`
            Values []string  `json:"values"`
         
    }
    
    // CartCheckoutCustomMeta ...
    type CartCheckoutCustomMeta struct {

        
            Value string  `json:"value"`
            Key string  `json:"key"`
         
    }
    
    // CartPosCheckoutDetailRequest ...
    type CartPosCheckoutDetailRequest struct {

        
            CallbackURL string  `json:"callback_url"`
            ID string  `json:"id"`
            OrderingStore float64  `json:"ordering_store"`
            CustomerDetails CustomerDetails  `json:"customer_details"`
            MerchantCode string  `json:"merchant_code"`
            PaymentParams map[string]interface{}  `json:"payment_params"`
            PaymentMode string  `json:"payment_mode"`
            OrderType string  `json:"order_type"`
            PaymentIdentifier string  `json:"payment_identifier"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
            Staff StaffCheckout  `json:"staff"`
            BillingAddressID string  `json:"billing_address_id"`
            Pos bool  `json:"pos"`
            Aggregator string  `json:"aggregator"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            Meta map[string]interface{}  `json:"meta"`
            AddressID string  `json:"address_id"`
            Files []Files  `json:"files"`
            PaymentAutoConfirm bool  `json:"payment_auto_confirm"`
            PickAtStoreUID float64  `json:"pick_at_store_uid"`
            CustomMeta []CartCheckoutCustomMeta  `json:"custom_meta"`
         
    }
    
    // CheckCart ...
    type CheckCart struct {

        
            ID string  `json:"id"`
            LastModified string  `json:"last_modified"`
            Items []CartProductInfo  `json:"items"`
            CodMessage string  `json:"cod_message"`
            CodCharges float64  `json:"cod_charges"`
            IsValid bool  `json:"is_valid"`
            CodAvailable bool  `json:"cod_available"`
            CouponText string  `json:"coupon_text"`
            OrderID string  `json:"order_id"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            Gstin string  `json:"gstin"`
            UID string  `json:"uid"`
            ErrorMessage string  `json:"error_message"`
            BuyNow bool  `json:"buy_now"`
            StoreCode string  `json:"store_code"`
            Currency CartCurrency  `json:"currency"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            CheckoutMode string  `json:"checkout_mode"`
            UserType string  `json:"user_type"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            StoreEmps []map[string]interface{}  `json:"store_emps"`
            Message string  `json:"message"`
            CartID float64  `json:"cart_id"`
            Comment string  `json:"comment"`
            DeliveryChargeOrderValue float64  `json:"delivery_charge_order_value"`
            Success bool  `json:"success"`
            DeliveryCharges float64  `json:"delivery_charges"`
         
    }
    
    // CartCheckoutResponse ...
    type CartCheckoutResponse struct {

        
            CallbackURL string  `json:"callback_url"`
            PaymentConfirmURL string  `json:"payment_confirm_url"`
            AppInterceptURL string  `json:"app_intercept_url"`
            OrderID string  `json:"order_id"`
            Cart CheckCart  `json:"cart"`
            Message string  `json:"message"`
            Data map[string]interface{}  `json:"data"`
            Success bool  `json:"success"`
         
    }
    
    // GiftDetail ...
    type GiftDetail struct {

        
            GiftMessage string  `json:"gift_message"`
            IsGiftApplied bool  `json:"is_gift_applied"`
         
    }
    
    // ArticleGiftDetail ...
    type ArticleGiftDetail struct {

        
            ArticleID GiftDetail  `json:"article_id"`
         
    }
    
    // CartMetaRequest ...
    type CartMetaRequest struct {

        
            CheckoutMode string  `json:"checkout_mode"`
            Comment string  `json:"comment"`
            GiftDetails ArticleGiftDetail  `json:"gift_details"`
            PickUpCustomerDetails map[string]interface{}  `json:"pick_up_customer_details"`
            DeliverySlots map[string]interface{}  `json:"delivery_slots"`
            Gstin string  `json:"gstin"`
         
    }
    
    // CartMetaResponse ...
    type CartMetaResponse struct {

        
            Message string  `json:"message"`
            IsValid bool  `json:"is_valid"`
         
    }
    
    // CartMetaMissingResponse ...
    type CartMetaMissingResponse struct {

        
            Errors []string  `json:"errors"`
         
    }
    
    // CartDeliveryModesResponse ...
    type CartDeliveryModesResponse struct {

        
            AvailableModes []string  `json:"available_modes"`
            PickupStores []float64  `json:"pickup_stores"`
         
    }
    
    // PickupStoreDetail ...
    type PickupStoreDetail struct {

        
            ID float64  `json:"id"`
            AreaCodeSlug string  `json:"area_code_slug"`
            Pincode float64  `json:"pincode"`
            Landmark string  `json:"landmark"`
            StoreCode string  `json:"store_code"`
            Area string  `json:"area"`
            City string  `json:"city"`
            Country string  `json:"country"`
            UID float64  `json:"uid"`
            Address string  `json:"address"`
            State string  `json:"state"`
            Email string  `json:"email"`
            AddressType string  `json:"address_type"`
            Phone string  `json:"phone"`
            AreaCode string  `json:"area_code"`
            Name string  `json:"name"`
         
    }
    
    // StoreDetailsResponse ...
    type StoreDetailsResponse struct {

        
            Items []PickupStoreDetail  `json:"items"`
         
    }
    
    // GetShareCartLinkRequest ...
    type GetShareCartLinkRequest struct {

        
            ID string  `json:"id"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // GetShareCartLinkResponse ...
    type GetShareCartLinkResponse struct {

        
            ShareURL string  `json:"share_url"`
            Token string  `json:"token"`
         
    }
    
    // SharedCartDetails ...
    type SharedCartDetails struct {

        
            CreatedOn string  `json:"created_on"`
            User map[string]interface{}  `json:"user"`
            Source map[string]interface{}  `json:"source"`
            Meta map[string]interface{}  `json:"meta"`
            Token string  `json:"token"`
         
    }
    
    // SharedCart ...
    type SharedCart struct {

        
            ID string  `json:"id"`
            LastModified string  `json:"last_modified"`
            Items []CartProductInfo  `json:"items"`
            IsValid bool  `json:"is_valid"`
            CouponText string  `json:"coupon_text"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            SharedCartDetails SharedCartDetails  `json:"shared_cart_details"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            Gstin string  `json:"gstin"`
            UID string  `json:"uid"`
            BuyNow bool  `json:"buy_now"`
            Currency CartCurrency  `json:"currency"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            CheckoutMode string  `json:"checkout_mode"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            Message string  `json:"message"`
            CartID float64  `json:"cart_id"`
            Comment string  `json:"comment"`
         
    }
    
    // SharedCartResponse ...
    type SharedCartResponse struct {

        
            Cart SharedCart  `json:"cart"`
            Error string  `json:"error"`
         
    }
    

    
    // PincodeErrorSchemaResponse ...
    type PincodeErrorSchemaResponse struct {

        
            Type string  `json:"type"`
            Value string  `json:"value"`
            Message string  `json:"message"`
         
    }
    
    // CountryMetaResponse ...
    type CountryMetaResponse struct {

        
            IsdCode string  `json:"isd_code"`
            CountryCode string  `json:"country_code"`
         
    }
    
    // PincodeParentsResponse ...
    type PincodeParentsResponse struct {

        
            Name string  `json:"name"`
            DisplayName string  `json:"display_name"`
            UID string  `json:"uid"`
            SubType string  `json:"sub_type"`
         
    }
    
    // PincodeLatLongData ...
    type PincodeLatLongData struct {

        
            Type string  `json:"type"`
            Coordinates []string  `json:"coordinates"`
         
    }
    
    // PincodeMetaResponse ...
    type PincodeMetaResponse struct {

        
            InternalZoneID float64  `json:"internal_zone_id"`
            Zone string  `json:"zone"`
         
    }
    
    // PincodeDataResponse ...
    type PincodeDataResponse struct {

        
            MetaCode CountryMetaResponse  `json:"meta_code"`
            Name string  `json:"name"`
            Parents []PincodeParentsResponse  `json:"parents"`
            LatLong PincodeLatLongData  `json:"lat_long"`
            DisplayName string  `json:"display_name"`
            Error PincodeErrorSchemaResponse  `json:"error"`
            Meta PincodeMetaResponse  `json:"meta"`
            UID string  `json:"uid"`
            SubType string  `json:"sub_type"`
         
    }
    
    // PincodeApiResponse ...
    type PincodeApiResponse struct {

        
            Success bool  `json:"success"`
            Error PincodeErrorSchemaResponse  `json:"error"`
            Data []PincodeDataResponse  `json:"data"`
         
    }
    
    // TATCategoryRequest ...
    type TATCategoryRequest struct {

        
            Level string  `json:"level"`
            ID float64  `json:"id"`
         
    }
    
    // TATArticlesRequest ...
    type TATArticlesRequest struct {

        
            Category TATCategoryRequest  `json:"category"`
            ManufacturingTimeUnit string  `json:"manufacturing_time_unit"`
            ManufacturingTime float64  `json:"manufacturing_time"`
         
    }
    
    // TATLocationDetailsRequest ...
    type TATLocationDetailsRequest struct {

        
            Articles []TATArticlesRequest  `json:"articles"`
            FulfillmentID float64  `json:"fulfillment_id"`
            FromPincode string  `json:"from_pincode"`
         
    }
    
    // TATViewRequest ...
    type TATViewRequest struct {

        
            ToPincode string  `json:"to_pincode"`
            LocationDetails []TATLocationDetailsRequest  `json:"location_details"`
            Action string  `json:"action"`
            Identifier string  `json:"identifier"`
            Source string  `json:"source"`
            Journey string  `json:"journey"`
         
    }
    
    // TATFormattedResponse ...
    type TATFormattedResponse struct {

        
            Min string  `json:"min"`
            Max string  `json:"max"`
         
    }
    
    // TATTimestampResponse ...
    type TATTimestampResponse struct {

        
            Min float64  `json:"min"`
            Max float64  `json:"max"`
         
    }
    
    // TATPromiseResponse ...
    type TATPromiseResponse struct {

        
            Formatted TATFormattedResponse  `json:"formatted"`
            Timestamp TATTimestampResponse  `json:"timestamp"`
         
    }
    
    // TATErrorSchemaResponse ...
    type TATErrorSchemaResponse struct {

        
            Type string  `json:"type"`
            Value string  `json:"value"`
            Message string  `json:"message"`
         
    }
    
    // TATArticlesResponse ...
    type TATArticlesResponse struct {

        
            IsCodAvailable bool  `json:"is_cod_available"`
            Promise TATPromiseResponse  `json:"promise"`
            ManufacturingTimeSeconds float64  `json:"_manufacturing_time_seconds"`
            Category TATCategoryRequest  `json:"category"`
            ManufacturingTime float64  `json:"manufacturing_time"`
            ManufacturingTimeUnit string  `json:"manufacturing_time_unit"`
            Error TATErrorSchemaResponse  `json:"error"`
         
    }
    
    // TATLocationDetailsResponse ...
    type TATLocationDetailsResponse struct {

        
            Articles []TATArticlesResponse  `json:"articles"`
            FulfillmentID float64  `json:"fulfillment_id"`
            FromPincode string  `json:"from_pincode"`
         
    }
    
    // TATViewResponse ...
    type TATViewResponse struct {

        
            IsCodAvailable bool  `json:"is_cod_available"`
            ToPincode string  `json:"to_pincode"`
            LocationDetails []TATLocationDetailsResponse  `json:"location_details"`
            StormbreakerUUID string  `json:"stormbreaker_uuid"`
            Action string  `json:"action"`
            ToCity string  `json:"to_city"`
            Identifier string  `json:"identifier"`
            Source string  `json:"source"`
            RequestUUID string  `json:"request_uuid"`
            Journey string  `json:"journey"`
            Success bool  `json:"success"`
            PaymentMode string  `json:"payment_mode"`
            Error TATErrorSchemaResponse  `json:"error"`
         
    }
    
    // DP ...
    type DP struct {

        
            FmPriority float64  `json:"fm_priority"`
            LmPriority float64  `json:"lm_priority"`
            RvpPriority float64  `json:"rvp_priority"`
            PaymentMode string  `json:"payment_mode"`
            Operations []string  `json:"operations"`
            AreaCode string  `json:"area_code"`
            AssignDpFromSb bool  `json:"assign_dp_from_sb"`
            InternalAccountID string  `json:"internal_account_id"`
            ExternalAccountID string  `json:"external_account_id"`
            TransportMode string  `json:"transport_mode"`
         
    }
    
    // LogisticsResponse ...
    type LogisticsResponse struct {

        
            Dp map[string]DP  `json:"dp"`
         
    }
    
    // CountryEntityResponse ...
    type CountryEntityResponse struct {

        
            Type string  `json:"type"`
            ParentID string  `json:"parent_id"`
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            Logistics LogisticsResponse  `json:"logistics"`
            DisplayName string  `json:"display_name"`
            Meta CountryMetaResponse  `json:"meta"`
            UID string  `json:"uid"`
            SubType string  `json:"sub_type"`
         
    }
    
    // CountryListResponse ...
    type CountryListResponse struct {

        
            Results []CountryEntityResponse  `json:"results"`
         
    }
    
    // GetZoneFromPincodeViewRequest ...
    type GetZoneFromPincodeViewRequest struct {

        
            Pincode string  `json:"pincode"`
            Country string  `json:"country"`
         
    }
    
    // GetZoneFromPincodeViewResponse ...
    type GetZoneFromPincodeViewResponse struct {

        
            Zones []string  `json:"zones"`
            ServiceabilityType string  `json:"serviceability_type"`
         
    }
    
    // ReAssignStoreRequest ...
    type ReAssignStoreRequest struct {

        
            Articles []map[string]interface{}  `json:"articles"`
            ToPincode string  `json:"to_pincode"`
            Identifier string  `json:"identifier"`
            IgnoredLocations []float64  `json:"ignored_locations"`
            Configuration map[string]interface{}  `json:"configuration"`
         
    }
    
    // ReAssignStoreResponse ...
    type ReAssignStoreResponse struct {

        
            AssignedStores []map[string]interface{}  `json:"assigned_stores"`
            ToPincode string  `json:"to_pincode"`
            PystormbreakerUUID string  `json:"pystormbreaker_uuid"`
            Success bool  `json:"success"`
            Error map[string]interface{}  `json:"error"`
         
    }
    
