package application



    
    // Meta ...
    type Meta struct {

        
            Source string  `json:"source"`
         
    }
    
    // Media ...
    type Media struct {

        
            Alt string  `json:"alt"`
            Meta Meta  `json:"meta"`
            URL string  `json:"url"`
            Type string  `json:"type"`
         
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
    
    // ProductBrand ...
    type ProductBrand struct {

        
            Logo Media  `json:"logo"`
            Description string  `json:"description"`
            Name string  `json:"name"`
            UID float64  `json:"uid"`
            Action ProductListingAction  `json:"action"`
         
    }
    
    // ApplicationItemMOQ ...
    type ApplicationItemMOQ struct {

        
            Minimum float64  `json:"minimum"`
            Maximum float64  `json:"maximum"`
            IncrementUnit float64  `json:"increment_unit"`
         
    }
    
    // NetQuantity ...
    type NetQuantity struct {

        
            Value float64  `json:"value"`
            Unit string  `json:"unit"`
         
    }
    
    // ProductDetailCustomOrder ...
    type ProductDetailCustomOrder struct {

        
            ManufacturingTimeUnit string  `json:"manufacturing_time_unit"`
            IsCustomOrder bool  `json:"is_custom_order"`
            ManufacturingTime float64  `json:"manufacturing_time"`
         
    }
    
    // CustomMetaFields ...
    type CustomMetaFields struct {

        
            Key string  `json:"key"`
            Value string  `json:"value"`
         
    }
    
    // ProductDetailAttribute ...
    type ProductDetailAttribute struct {

        
            Key string  `json:"key"`
            Value string  `json:"value"`
            Type string  `json:"type"`
         
    }
    
    // ProductDetailGroupedAttribute ...
    type ProductDetailGroupedAttribute struct {

        
            Details []ProductDetailAttribute  `json:"details"`
            Title string  `json:"title"`
         
    }
    
    // ApplicationItemSEO ...
    type ApplicationItemSEO struct {

        
            Description string  `json:"description"`
            Title string  `json:"title"`
         
    }
    
    // ProductCategoryMap ...
    type ProductCategoryMap struct {

        
            L3 ProductBrand  `json:"l3"`
            L2 ProductBrand  `json:"l2"`
            L1 ProductBrand  `json:"l1"`
         
    }
    
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
    
    // ProductDetail ...
    type ProductDetail struct {

        
            Brand ProductBrand  `json:"brand"`
            Color string  `json:"color"`
            ProductGroupTag []string  `json:"product_group_tag"`
            Categories []ProductBrand  `json:"categories"`
            ItemCode string  `json:"item_code"`
            IsDependent bool  `json:"is_dependent"`
            Action ProductListingAction  `json:"action"`
            ShortDescription string  `json:"short_description"`
            Discount string  `json:"discount"`
            Moq ApplicationItemMOQ  `json:"moq"`
            Attributes map[string]interface{}  `json:"attributes"`
            Description string  `json:"description"`
            RatingCount float64  `json:"rating_count"`
            NetQuantity NetQuantity  `json:"net_quantity"`
            ProductOnlineDate string  `json:"product_online_date"`
            Type string  `json:"type"`
            Slug string  `json:"slug"`
            CustomOrder ProductDetailCustomOrder  `json:"custom_order"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Similars []string  `json:"similars"`
            Tryouts []string  `json:"tryouts"`
            Medias []Media  `json:"medias"`
            Tags []string  `json:"tags"`
            Highlights []string  `json:"highlights"`
            ItemType string  `json:"item_type"`
            CustomMeta []CustomMetaFields  `json:"_custom_meta"`
            Name string  `json:"name"`
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
            ImageNature string  `json:"image_nature"`
            Seo ApplicationItemSEO  `json:"seo"`
            Rating float64  `json:"rating"`
            HasVariant bool  `json:"has_variant"`
            CategoryMap ProductCategoryMap  `json:"category_map"`
            Price ProductListingPrice  `json:"price"`
            UID float64  `json:"uid"`
            TeaserTag string  `json:"teaser_tag"`
         
    }
    
    // ErrorResponse ...
    type ErrorResponse struct {

        
            Error string  `json:"error"`
         
    }
    
    // Dimension ...
    type Dimension struct {

        
            Height float64  `json:"height"`
            Length float64  `json:"length"`
            Width float64  `json:"width"`
            Unit string  `json:"unit"`
            IsDefault bool  `json:"is_default"`
         
    }
    
    // Weight ...
    type Weight struct {

        
            Shipping float64  `json:"shipping"`
            IsDefault bool  `json:"is_default"`
            Unit string  `json:"unit"`
         
    }
    
    // ProductSize ...
    type ProductSize struct {

        
            Value string  `json:"value"`
            Dimension Dimension  `json:"dimension"`
            IsAvailable bool  `json:"is_available"`
            Weight Weight  `json:"weight"`
            SellerIdentifiers []string  `json:"seller_identifiers"`
            Display string  `json:"display"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // SizeChartValues ...
    type SizeChartValues struct {

        
            Col5 string  `json:"col_5"`
            Col2 string  `json:"col_2"`
            Col4 string  `json:"col_4"`
            Col1 string  `json:"col_1"`
            Col6 string  `json:"col_6"`
            Col3 string  `json:"col_3"`
         
    }
    
    // ColumnHeader ...
    type ColumnHeader struct {

        
            Value string  `json:"value"`
            Convertable bool  `json:"convertable"`
         
    }
    
    // ColumnHeaders ...
    type ColumnHeaders struct {

        
            Col5 ColumnHeader  `json:"col_5"`
            Col2 ColumnHeader  `json:"col_2"`
            Col4 ColumnHeader  `json:"col_4"`
            Col1 ColumnHeader  `json:"col_1"`
            Col6 ColumnHeader  `json:"col_6"`
            Col3 ColumnHeader  `json:"col_3"`
         
    }
    
    // SizeChart ...
    type SizeChart struct {

        
            Unit string  `json:"unit"`
            SizeTip string  `json:"size_tip"`
            Description string  `json:"description"`
            Sizes []SizeChartValues  `json:"sizes"`
            Title string  `json:"title"`
            Image string  `json:"image"`
            Headers ColumnHeaders  `json:"headers"`
         
    }
    
    // ProductSizeStores ...
    type ProductSizeStores struct {

        
            Count float64  `json:"count"`
         
    }
    
    // ProductSizes ...
    type ProductSizes struct {

        
            Discount string  `json:"discount"`
            Sellable bool  `json:"sellable"`
            Sizes []ProductSize  `json:"sizes"`
            SizeChart SizeChart  `json:"size_chart"`
            Stores ProductSizeStores  `json:"stores"`
            Price ProductListingPrice  `json:"price"`
            MultiSize bool  `json:"multi_size"`
         
    }
    
    // AttributeDetail ...
    type AttributeDetail struct {

        
            Description string  `json:"description"`
            Key string  `json:"key"`
            Logo string  `json:"logo"`
            Display string  `json:"display"`
         
    }
    
    // AttributeMetadata ...
    type AttributeMetadata struct {

        
            Details []AttributeDetail  `json:"details"`
            Title string  `json:"title"`
         
    }
    
    // ProductsComparisonResponse ...
    type ProductsComparisonResponse struct {

        
            AttributesMetadata []AttributeMetadata  `json:"attributes_metadata"`
            Items []ProductDetail  `json:"items"`
         
    }
    
    // ProductCompareResponse ...
    type ProductCompareResponse struct {

        
            AttributesMetadata []AttributeMetadata  `json:"attributes_metadata"`
            Subtitle string  `json:"subtitle"`
            Items []ProductDetail  `json:"items"`
            Title string  `json:"title"`
         
    }
    
    // ProductFrequentlyComparedSimilarResponse ...
    type ProductFrequentlyComparedSimilarResponse struct {

        
            Similars ProductCompareResponse  `json:"similars"`
         
    }
    
    // ProductVariantItemResponse ...
    type ProductVariantItemResponse struct {

        
            Color string  `json:"color"`
            Medias []Media  `json:"medias"`
            Value string  `json:"value"`
            IsAvailable bool  `json:"is_available"`
            Name string  `json:"name"`
            ColorName string  `json:"color_name"`
            CustomMeta []CustomMetaFields  `json:"_custom_meta"`
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
            Action ProductListingAction  `json:"action"`
         
    }
    
    // ProductVariantResponse ...
    type ProductVariantResponse struct {

        
            DisplayType string  `json:"display_type"`
            Key string  `json:"key"`
            Header string  `json:"header"`
            Items []ProductVariantItemResponse  `json:"items"`
         
    }
    
    // ProductVariantsResponse ...
    type ProductVariantsResponse struct {

        
            Variants []ProductVariantResponse  `json:"variants"`
         
    }
    
    // Seller ...
    type Seller struct {

        
            Name string  `json:"name"`
            UID float64  `json:"uid"`
            Count float64  `json:"count"`
         
    }
    
    // CompanyDetail ...
    type CompanyDetail struct {

        
            Name string  `json:"name"`
            ID float64  `json:"id"`
         
    }
    
    // StoreDetail ...
    type StoreDetail struct {

        
            Name string  `json:"name"`
            Code string  `json:"code"`
            ID float64  `json:"id"`
            City string  `json:"city"`
         
    }
    
    // ProductStockPrice ...
    type ProductStockPrice struct {

        
            Currency string  `json:"currency"`
            Effective float64  `json:"effective"`
            Marked float64  `json:"marked"`
         
    }
    
    // ProductStockStatusItem ...
    type ProductStockStatusItem struct {

        
            Identifier map[string]interface{}  `json:"identifier"`
            ItemID float64  `json:"item_id"`
            Seller Seller  `json:"seller"`
            Company CompanyDetail  `json:"company"`
            Store StoreDetail  `json:"store"`
            Size string  `json:"size"`
            Price ProductStockPrice  `json:"price"`
            UID string  `json:"uid"`
            Quantity float64  `json:"quantity"`
         
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

        
            Items []ProductStockStatusItem  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // ProductSortOn ...
    type ProductSortOn struct {

        
            IsSelected bool  `json:"is_selected"`
            Name string  `json:"name"`
            Value string  `json:"value"`
         
    }
    
    // ProductVariantListingResponse ...
    type ProductVariantListingResponse struct {

        
            DisplayType string  `json:"display_type"`
            Header string  `json:"header"`
            Items []ProductVariantItemResponse  `json:"items"`
            Total float64  `json:"total"`
            Key string  `json:"key"`
         
    }
    
    // ProductListingDetail ...
    type ProductListingDetail struct {

        
            Brand ProductBrand  `json:"brand"`
            Color string  `json:"color"`
            ProductGroupTag []string  `json:"product_group_tag"`
            Categories []ProductBrand  `json:"categories"`
            ItemCode string  `json:"item_code"`
            IsDependent bool  `json:"is_dependent"`
            Action ProductListingAction  `json:"action"`
            ShortDescription string  `json:"short_description"`
            Discount string  `json:"discount"`
            Moq ApplicationItemMOQ  `json:"moq"`
            Attributes map[string]interface{}  `json:"attributes"`
            Description string  `json:"description"`
            RatingCount float64  `json:"rating_count"`
            NetQuantity NetQuantity  `json:"net_quantity"`
            ProductOnlineDate string  `json:"product_online_date"`
            Type string  `json:"type"`
            Slug string  `json:"slug"`
            CustomOrder ProductDetailCustomOrder  `json:"custom_order"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Similars []string  `json:"similars"`
            Tryouts []string  `json:"tryouts"`
            Medias []Media  `json:"medias"`
            Tags []string  `json:"tags"`
            Highlights []string  `json:"highlights"`
            Sellable bool  `json:"sellable"`
            ItemType string  `json:"item_type"`
            Sizes []string  `json:"sizes"`
            CustomMeta []CustomMetaFields  `json:"_custom_meta"`
            Name string  `json:"name"`
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
            Variants []ProductVariantListingResponse  `json:"variants"`
            ImageNature string  `json:"image_nature"`
            Seo ApplicationItemSEO  `json:"seo"`
            Rating float64  `json:"rating"`
            Identifiers []string  `json:"identifiers"`
            HasVariant bool  `json:"has_variant"`
            CategoryMap ProductCategoryMap  `json:"category_map"`
            Price ProductListingPrice  `json:"price"`
            UID float64  `json:"uid"`
            TeaserTag string  `json:"teaser_tag"`
         
    }
    
    // ProductFiltersValue ...
    type ProductFiltersValue struct {

        
            CurrencySymbol string  `json:"currency_symbol"`
            Count float64  `json:"count"`
            Max float64  `json:"max"`
            Display string  `json:"display"`
            Value string  `json:"value"`
            SelectedMin float64  `json:"selected_min"`
            CurrencyCode string  `json:"currency_code"`
            DisplayFormat string  `json:"display_format"`
            QueryFormat string  `json:"query_format"`
            Min float64  `json:"min"`
            SelectedMax float64  `json:"selected_max"`
            IsSelected bool  `json:"is_selected"`
         
    }
    
    // ProductFiltersKey ...
    type ProductFiltersKey struct {

        
            Name string  `json:"name"`
            Kind string  `json:"kind"`
            Logo string  `json:"logo"`
            Display string  `json:"display"`
         
    }
    
    // ProductFilters ...
    type ProductFilters struct {

        
            Values []ProductFiltersValue  `json:"values"`
            Key ProductFiltersKey  `json:"key"`
         
    }
    
    // ProductListingResponse ...
    type ProductListingResponse struct {

        
            SortOn []ProductSortOn  `json:"sort_on"`
            Items []ProductListingDetail  `json:"items"`
            Page Page  `json:"page"`
            Filters []ProductFilters  `json:"filters"`
         
    }
    
    // ImageUrls ...
    type ImageUrls struct {

        
            Landscape Media  `json:"landscape"`
            Portrait Media  `json:"portrait"`
         
    }
    
    // BrandItem ...
    type BrandItem struct {

        
            Discount string  `json:"discount"`
            Logo Media  `json:"logo"`
            Banners ImageUrls  `json:"banners"`
            Departments []string  `json:"departments"`
            Description string  `json:"description"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
            Action ProductListingAction  `json:"action"`
         
    }
    
    // BrandListingResponse ...
    type BrandListingResponse struct {

        
            Items []BrandItem  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // BrandDetailResponse ...
    type BrandDetailResponse struct {

        
            Logo Media  `json:"logo"`
            Banners ImageUrls  `json:"banners"`
            Description string  `json:"description"`
            Name string  `json:"name"`
            UID float64  `json:"uid"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // DepartmentIdentifier ...
    type DepartmentIdentifier struct {

        
            UID float64  `json:"uid"`
            Slug string  `json:"slug"`
         
    }
    
    // CategoryBanner ...
    type CategoryBanner struct {

        
            Landscape Media  `json:"landscape"`
            Portrait Media  `json:"portrait"`
         
    }
    
    // ThirdLevelChild ...
    type ThirdLevelChild struct {

        
            Action ProductListingAction  `json:"action"`
            Banners ImageUrls  `json:"banners"`
            Name string  `json:"name"`
            Childs []map[string]interface{}  `json:"childs"`
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // SecondLevelChild ...
    type SecondLevelChild struct {

        
            Action ProductListingAction  `json:"action"`
            Banners ImageUrls  `json:"banners"`
            Name string  `json:"name"`
            Childs []ThirdLevelChild  `json:"childs"`
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // Child ...
    type Child struct {

        
            Action ProductListingAction  `json:"action"`
            Banners ImageUrls  `json:"banners"`
            Name string  `json:"name"`
            Childs []SecondLevelChild  `json:"childs"`
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // CategoryItems ...
    type CategoryItems struct {

        
            Banners CategoryBanner  `json:"banners"`
            Name string  `json:"name"`
            Childs []Child  `json:"childs"`
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
            Action ProductListingAction  `json:"action"`
         
    }
    
    // DepartmentCategoryTree ...
    type DepartmentCategoryTree struct {

        
            Items []CategoryItems  `json:"items"`
            Department string  `json:"department"`
         
    }
    
    // CategoryListingResponse ...
    type CategoryListingResponse struct {

        
            Departments []DepartmentIdentifier  `json:"departments"`
            Data []DepartmentCategoryTree  `json:"data"`
         
    }
    
    // CategoryMetaResponse ...
    type CategoryMetaResponse struct {

        
            Logo Media  `json:"logo"`
            Banners ImageUrls  `json:"banners"`
            Name string  `json:"name"`
            UID float64  `json:"uid"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // HomeListingResponse ...
    type HomeListingResponse struct {

        
            Message string  `json:"message"`
            Items []ProductListingDetail  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // Department ...
    type Department struct {

        
            Logo Media  `json:"logo"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
            PriorityOrder float64  `json:"priority_order"`
         
    }
    
    // DepartmentResponse ...
    type DepartmentResponse struct {

        
            Items []Department  `json:"items"`
         
    }
    
    // AutocompleteItem ...
    type AutocompleteItem struct {

        
            Action ProductListingAction  `json:"action"`
            Logo Media  `json:"logo"`
            Type string  `json:"type"`
            Display string  `json:"display"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // AutoCompleteResponse ...
    type AutoCompleteResponse struct {

        
            Items []AutocompleteItem  `json:"items"`
         
    }
    
    // CollectionQuery ...
    type CollectionQuery struct {

        
            Op string  `json:"op"`
            Value []interface{}  `json:"value"`
            Attribute string  `json:"attribute"`
         
    }
    
    // GetCollectionDetailNest ...
    type GetCollectionDetailNest struct {

        
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Priority float64  `json:"priority"`
            Query []CollectionQuery  `json:"query"`
            Action ProductListingAction  `json:"action"`
            Badge map[string]interface{}  `json:"badge"`
            Banners ImageUrls  `json:"banners"`
            Description string  `json:"description"`
            IsActive bool  `json:"is_active"`
            Slug string  `json:"slug"`
            Type string  `json:"type"`
            AppID string  `json:"app_id"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            AllowSort bool  `json:"allow_sort"`
            Schedule map[string]interface{}  `json:"_schedule"`
            Name string  `json:"name"`
            Cron map[string]interface{}  `json:"cron"`
            Meta map[string]interface{}  `json:"meta"`
            Tag []string  `json:"tag"`
            AllowFacets bool  `json:"allow_facets"`
            Logo Media  `json:"logo"`
            SortOn string  `json:"sort_on"`
            UID string  `json:"uid"`
         
    }
    
    // CollectionListingFilterTag ...
    type CollectionListingFilterTag struct {

        
            Name string  `json:"name"`
            IsSelected bool  `json:"is_selected"`
            Display string  `json:"display"`
         
    }
    
    // CollectionListingFilterType ...
    type CollectionListingFilterType struct {

        
            Name string  `json:"name"`
            IsSelected bool  `json:"is_selected"`
            Display string  `json:"display"`
         
    }
    
    // CollectionListingFilter ...
    type CollectionListingFilter struct {

        
            Tags []CollectionListingFilterTag  `json:"tags"`
            Type []CollectionListingFilterType  `json:"type"`
         
    }
    
    // GetCollectionListingResponse ...
    type GetCollectionListingResponse struct {

        
            Page Page  `json:"page"`
            Items []GetCollectionDetailNest  `json:"items"`
            Filters CollectionListingFilter  `json:"filters"`
         
    }
    
    // CollectionDetailResponse ...
    type CollectionDetailResponse struct {

        
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Priority float64  `json:"priority"`
            Query []CollectionQuery  `json:"query"`
            Badge map[string]interface{}  `json:"badge"`
            Banners ImageUrls  `json:"banners"`
            Description string  `json:"description"`
            IsActive bool  `json:"is_active"`
            Slug string  `json:"slug"`
            Type string  `json:"type"`
            AppID string  `json:"app_id"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            AllowSort bool  `json:"allow_sort"`
            Schedule map[string]interface{}  `json:"_schedule"`
            Name string  `json:"name"`
            Cron map[string]interface{}  `json:"cron"`
            Meta map[string]interface{}  `json:"meta"`
            Tag []string  `json:"tag"`
            AllowFacets bool  `json:"allow_facets"`
            Logo Media  `json:"logo"`
            SortOn string  `json:"sort_on"`
         
    }
    
    // GetFollowListingResponse ...
    type GetFollowListingResponse struct {

        
            Items []ProductListingDetail  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // FollowPostResponse ...
    type FollowPostResponse struct {

        
            ID string  `json:"id"`
            Message string  `json:"message"`
         
    }
    
    // FollowerCountResponse ...
    type FollowerCountResponse struct {

        
            Count float64  `json:"count"`
         
    }
    
    // FollowIdsData ...
    type FollowIdsData struct {

        
            Products []float64  `json:"products"`
            Collections []float64  `json:"collections"`
            Brands []float64  `json:"brands"`
         
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

        
            State string  `json:"state"`
            LatLong LatLong  `json:"lat_long"`
            StoreCode string  `json:"store_code"`
            Name string  `json:"name"`
            StoreEmail string  `json:"store_email"`
            Address string  `json:"address"`
            UID float64  `json:"uid"`
            Pincode float64  `json:"pincode"`
            City string  `json:"city"`
            Country string  `json:"country"`
         
    }
    
    // StoreListingResponse ...
    type StoreListingResponse struct {

        
            Items []Store  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // SellerPhoneNumber ...
    type SellerPhoneNumber struct {

        
            CountryCode float64  `json:"country_code"`
            Number string  `json:"number"`
         
    }
    
    // StoreDepartments ...
    type StoreDepartments struct {

        
            Logo map[string]interface{}  `json:"logo"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
            PriorityOrder float64  `json:"priority_order"`
         
    }
    
    // CompanyStore ...
    type CompanyStore struct {

        
            Name string  `json:"name"`
            UID float64  `json:"uid"`
            BusinessType string  `json:"business_type"`
            CompanyType string  `json:"company_type"`
         
    }
    
    // StoreManagerSerializer ...
    type StoreManagerSerializer struct {

        
            Name string  `json:"name"`
            Email string  `json:"email"`
            MobileNo SellerPhoneNumber  `json:"mobile_no"`
         
    }
    
    // StoreAddressSerializer ...
    type StoreAddressSerializer struct {

        
            Address1 string  `json:"address1"`
            Landmark string  `json:"landmark"`
            Latitude float64  `json:"latitude"`
            State string  `json:"state"`
            Address2 string  `json:"address2"`
            Longitude float64  `json:"longitude"`
            Pincode float64  `json:"pincode"`
            City string  `json:"city"`
            Country string  `json:"country"`
         
    }
    
    // AppStore ...
    type AppStore struct {

        
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            Departments []StoreDepartments  `json:"departments"`
            StoreCode string  `json:"store_code"`
            Company CompanyStore  `json:"company"`
            Name string  `json:"name"`
            Manager StoreManagerSerializer  `json:"manager"`
            Address StoreAddressSerializer  `json:"address"`
            UID float64  `json:"uid"`
         
    }
    
    // ApplicationStoreListing ...
    type ApplicationStoreListing struct {

        
            Filters []map[string]interface{}  `json:"filters"`
            Items []AppStore  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // Time ...
    type Time struct {

        
            Hour float64  `json:"hour"`
            Minute float64  `json:"minute"`
         
    }
    
    // StoreTiming ...
    type StoreTiming struct {

        
            Opening Time  `json:"opening"`
            Weekday string  `json:"weekday"`
            Closing Time  `json:"closing"`
            Open bool  `json:"open"`
         
    }
    
    // StoreDetails ...
    type StoreDetails struct {

        
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            Departments []StoreDepartments  `json:"departments"`
            StoreCode string  `json:"store_code"`
            Company CompanyStore  `json:"company"`
            Timing []StoreTiming  `json:"timing"`
            Name string  `json:"name"`
            Manager StoreManagerSerializer  `json:"manager"`
            Address StoreAddressSerializer  `json:"address"`
            UID float64  `json:"uid"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // ProductDetails ...
    type ProductDetails struct {

        
            OutOfStock bool  `json:"out_of_stock"`
            ItemCode string  `json:"item_code"`
            HsnCode float64  `json:"hsn_code"`
            BrandUID float64  `json:"brand_uid"`
            Identifier map[string]interface{}  `json:"identifier"`
            ShortDescription string  `json:"short_description"`
            TemplateTag string  `json:"template_tag"`
            Attributes map[string]interface{}  `json:"attributes"`
            Description string  `json:"description"`
            RatingCount float64  `json:"rating_count"`
            Slug string  `json:"slug"`
            Images []string  `json:"images"`
            CountryOfOrigin string  `json:"country_of_origin"`
            Media []map[string]interface{}  `json:"media"`
            Highlights []string  `json:"highlights"`
            IsSet bool  `json:"is_set"`
            Name string  `json:"name"`
            GroupedAttributes map[string]interface{}  `json:"grouped_attributes"`
            ImageNature string  `json:"image_nature"`
            Rating float64  `json:"rating"`
            HasVariant bool  `json:"has_variant"`
         
    }
    
    // Size ...
    type Size struct {

        
            IsAvailable bool  `json:"is_available"`
            Quantity float64  `json:"quantity"`
            Value string  `json:"value"`
            Display string  `json:"display"`
         
    }
    
    // ProductGroupPrice ...
    type ProductGroupPrice struct {

        
            Currency string  `json:"currency"`
            MaxEffective float64  `json:"max_effective"`
            MaxMarked float64  `json:"max_marked"`
            MinEffective float64  `json:"min_effective"`
            MinMarked float64  `json:"min_marked"`
         
    }
    
    // ProductInGroup ...
    type ProductInGroup struct {

        
            AllowRemove bool  `json:"allow_remove"`
            ProductDetails ProductDetails  `json:"product_details"`
            AutoSelect bool  `json:"auto_select"`
            ProductUID float64  `json:"product_uid"`
            Sizes []Size  `json:"sizes"`
            AutoAddToCart bool  `json:"auto_add_to_cart"`
            MaxQuantity float64  `json:"max_quantity"`
            Price ProductGroupPrice  `json:"price"`
            MinQuantity float64  `json:"min_quantity"`
         
    }
    
    // UserDetail ...
    type UserDetail struct {

        
            SuperUser bool  `json:"super_user"`
            UserID string  `json:"user_id"`
            Contact string  `json:"contact"`
            Username string  `json:"username"`
         
    }
    
    // ProductGroupingModel ...
    type ProductGroupingModel struct {

        
            Products []ProductInGroup  `json:"products"`
            VerifiedBy UserDetail  `json:"verified_by"`
            Logo string  `json:"logo"`
            PageVisibility []string  `json:"page_visibility"`
            ID interface{}  `json:"_id"`
            CreatedBy UserDetail  `json:"created_by"`
            ModifiedBy UserDetail  `json:"modified_by"`
            VerifiedOn string  `json:"verified_on"`
            Name string  `json:"name"`
            ModifiedOn string  `json:"modified_on"`
            CompanyID float64  `json:"company_id"`
            Slug string  `json:"slug"`
            Meta map[string]interface{}  `json:"meta"`
            CreatedOn string  `json:"created_on"`
            IsActive bool  `json:"is_active"`
            Choice string  `json:"choice"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
         
    }
    
    // ProductBundle ...
    type ProductBundle struct {

        
            Items []ProductGroupingModel  `json:"items"`
         
    }
    
    // ArticleAssignmentV3 ...
    type ArticleAssignmentV3 struct {

        
            Strategy string  `json:"strategy"`
            Level string  `json:"level"`
         
    }
    
    // SellerV3 ...
    type SellerV3 struct {

        
            Name string  `json:"name"`
            UID float64  `json:"uid"`
            Count float64  `json:"count"`
         
    }
    
    // StrategyWiseListingSchemaV3 ...
    type StrategyWiseListingSchemaV3 struct {

        
            Distance float64  `json:"distance"`
            Pincode float64  `json:"pincode"`
            Tat float64  `json:"tat"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // ReturnConfigSchemaV3 ...
    type ReturnConfigSchemaV3 struct {

        
            Returnable bool  `json:"returnable"`
            Time float64  `json:"time"`
            Unit string  `json:"unit"`
         
    }
    
    // DetailsSchemaV3 ...
    type DetailsSchemaV3 struct {

        
            Key string  `json:"key"`
            Value string  `json:"value"`
            Type string  `json:"type"`
         
    }
    
    // MarketPlaceSttributesSchemaV3 ...
    type MarketPlaceSttributesSchemaV3 struct {

        
            Details []DetailsSchemaV3  `json:"details"`
            Title string  `json:"title"`
         
    }
    
    // ProductStockPriceV3 ...
    type ProductStockPriceV3 struct {

        
            Currency string  `json:"currency"`
            Effective float64  `json:"effective"`
            Marked float64  `json:"marked"`
         
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

        
            SizeDistribution ProductSetDistributionV3  `json:"size_distribution"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // StoreV3 ...
    type StoreV3 struct {

        
            Name string  `json:"name"`
            UID float64  `json:"uid"`
            Count float64  `json:"count"`
         
    }
    
    // SellerGroupAttributes ...
    type SellerGroupAttributes struct {

        
            Details []DetailsSchemaV3  `json:"details"`
            Title string  `json:"title"`
         
    }
    
    // ProductStockUnitPriceV3 ...
    type ProductStockUnitPriceV3 struct {

        
            CurrencySymbol string  `json:"currency_symbol"`
            CurrencyCode string  `json:"currency_code"`
            Price float64  `json:"price"`
            Unit string  `json:"unit"`
         
    }
    
    // ProductSizePriceResponseV3 ...
    type ProductSizePriceResponseV3 struct {

        
            ArticleAssignment ArticleAssignmentV3  `json:"article_assignment"`
            Seller SellerV3  `json:"seller"`
            StrategyWiseListing []StrategyWiseListingSchemaV3  `json:"strategy_wise_listing"`
            ReturnConfig ReturnConfigSchemaV3  `json:"return_config"`
            Discount string  `json:"discount"`
            MarketplaceAttributes []MarketPlaceSttributesSchemaV3  `json:"marketplace_attributes"`
            PricePerPiece ProductStockPriceV3  `json:"price_per_piece"`
            Set ProductSetV3  `json:"set"`
            IsCod bool  `json:"is_cod"`
            SpecialBadge string  `json:"special_badge"`
            IsGift bool  `json:"is_gift"`
            ItemType string  `json:"item_type"`
            Store StoreV3  `json:"store"`
            GroupedAttributes []SellerGroupAttributes  `json:"grouped_attributes"`
            PricePerUnit ProductStockUnitPriceV3  `json:"price_per_unit"`
            ArticleID string  `json:"article_id"`
            Quantity float64  `json:"quantity"`
            Price ProductStockPriceV3  `json:"price"`
            SellerCount float64  `json:"seller_count"`
            Pincode float64  `json:"pincode"`
            LongLat []float64  `json:"long_lat"`
         
    }
    
    // ProductSizeSellerFilterSchemaV3 ...
    type ProductSizeSellerFilterSchemaV3 struct {

        
            Name string  `json:"name"`
            Value string  `json:"value"`
            IsSelected bool  `json:"is_selected"`
         
    }
    
    // ProductSizeSellersResponseV3 ...
    type ProductSizeSellersResponseV3 struct {

        
            SortOn []ProductSizeSellerFilterSchemaV3  `json:"sort_on"`
            Items []ProductSizePriceResponseV3  `json:"items"`
            Page Page  `json:"page"`
         
    }
    

    
    // CartCurrencySchema ...
    type CartCurrencySchema struct {

        
            Symbol string  `json:"symbol"`
            Code string  `json:"code"`
         
    }
    
    // PromiseTimestamp ...
    type PromiseTimestamp struct {

        
            Max float64  `json:"max"`
            Min float64  `json:"min"`
         
    }
    
    // PromiseFormatted ...
    type PromiseFormatted struct {

        
            Max string  `json:"max"`
            Min string  `json:"min"`
         
    }
    
    // ShipmentPromise ...
    type ShipmentPromise struct {

        
            Timestamp PromiseTimestamp  `json:"timestamp"`
            Formatted PromiseFormatted  `json:"formatted"`
         
    }
    
    // BuyRulesSchema ...
    type BuyRulesSchema struct {

        
            ItemCriteria map[string]interface{}  `json:"item_criteria"`
            CartConditions map[string]interface{}  `json:"cart_conditions"`
         
    }
    
    // FreeGiftItemSchema ...
    type FreeGiftItemSchema struct {

        
            ItemBrandName string  `json:"item_brand_name"`
            ItemPriceDetails map[string]interface{}  `json:"item_price_details"`
            ItemName string  `json:"item_name"`
            ItemSlug string  `json:"item_slug"`
            ItemID float64  `json:"item_id"`
            ItemImagesURL []string  `json:"item_images_url"`
         
    }
    
    // AppliedFreeArticlesSchema ...
    type AppliedFreeArticlesSchema struct {

        
            ArticleID string  `json:"article_id"`
            ParentItemIdentifier string  `json:"parent_item_identifier"`
            Quantity float64  `json:"quantity"`
            FreeGiftItemDetails FreeGiftItemSchema  `json:"free_gift_item_details"`
         
    }
    
    // Ownership ...
    type Ownership struct {

        
            PayableCategory string  `json:"payable_category"`
            PayableBy string  `json:"payable_by"`
         
    }
    
    // DiscountRulesAppSchema ...
    type DiscountRulesAppSchema struct {

        
            MatchedBuyRules []string  `json:"matched_buy_rules"`
            RawOffer map[string]interface{}  `json:"raw_offer"`
            ItemCriteria map[string]interface{}  `json:"item_criteria"`
            Offer map[string]interface{}  `json:"offer"`
         
    }
    
    // AppliedPromotion ...
    type AppliedPromotion struct {

        
            OfferText string  `json:"offer_text"`
            PromotionType string  `json:"promotion_type"`
            Amount float64  `json:"amount"`
            BuyRules []BuyRulesSchema  `json:"buy_rules"`
            PromotionName string  `json:"promotion_name"`
            MrpPromotion bool  `json:"mrp_promotion"`
            PromotionGroup string  `json:"promotion_group"`
            PromoID string  `json:"promo_id"`
            AppliedFreeArticles []AppliedFreeArticlesSchema  `json:"applied_free_articles"`
            Ownership Ownership  `json:"ownership"`
            DiscountRules []DiscountRulesAppSchema  `json:"discount_rules"`
            ArticleQuantity float64  `json:"article_quantity"`
         
    }
    
    // ProductAvailabilitySize ...
    type ProductAvailabilitySize struct {

        
            Value string  `json:"value"`
            Display string  `json:"display"`
            IsAvailable bool  `json:"is_available"`
         
    }
    
    // ProductAvailability ...
    type ProductAvailability struct {

        
            OutOfStock bool  `json:"out_of_stock"`
            IsValid bool  `json:"is_valid"`
            Deliverable bool  `json:"deliverable"`
            OtherStoreQuantity float64  `json:"other_store_quantity"`
            Sizes []string  `json:"sizes"`
            AvailableSizes []ProductAvailabilitySize  `json:"available_sizes"`
         
    }
    
    // PromoMeta ...
    type PromoMeta struct {

        
            Message string  `json:"message"`
         
    }
    
    // ProductPrice ...
    type ProductPrice struct {

        
            CurrencySymbol string  `json:"currency_symbol"`
            Marked float64  `json:"marked"`
            Effective float64  `json:"effective"`
            CurrencyCode string  `json:"currency_code"`
            Selling float64  `json:"selling"`
            AddOn float64  `json:"add_on"`
         
    }
    
    // ProductPriceInfo ...
    type ProductPriceInfo struct {

        
            Base ProductPrice  `json:"base"`
            Converted ProductPrice  `json:"converted"`
         
    }
    
    // CartProductIdentifer ...
    type CartProductIdentifer struct {

        
            Identifier string  `json:"identifier"`
         
    }
    
    // CouponDetails ...
    type CouponDetails struct {

        
            Code string  `json:"code"`
            DiscountTotalQuantity float64  `json:"discount_total_quantity"`
            DiscountSingleQuantity float64  `json:"discount_single_quantity"`
         
    }
    
    // BaseInfo ...
    type BaseInfo struct {

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // BasePriceSchema ...
    type BasePriceSchema struct {

        
            CurrencySymbol string  `json:"currency_symbol"`
            Marked float64  `json:"marked"`
            Effective float64  `json:"effective"`
            CurrencyCode string  `json:"currency_code"`
         
    }
    
    // ArticlePriceInfo ...
    type ArticlePriceInfo struct {

        
            Base BasePriceSchema  `json:"base"`
            Converted BasePriceSchema  `json:"converted"`
         
    }
    
    // ProductArticle ...
    type ProductArticle struct {

        
            Size string  `json:"size"`
            ProductGroupTags []string  `json:"product_group_tags"`
            CartItemMeta map[string]interface{}  `json:"cart_item_meta"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            SellerIdentifier string  `json:"seller_identifier"`
            Quantity float64  `json:"quantity"`
            Identifier map[string]interface{}  `json:"identifier"`
            GiftCard map[string]interface{}  `json:"gift_card"`
            Store BaseInfo  `json:"store"`
            Type string  `json:"type"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            IsGiftVisible bool  `json:"is_gift_visible"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            Price ArticlePriceInfo  `json:"price"`
            Seller BaseInfo  `json:"seller"`
            UID string  `json:"uid"`
         
    }
    
    // ProductImage ...
    type ProductImage struct {

        
            URL string  `json:"url"`
            AspectRatio string  `json:"aspect_ratio"`
            SecureURL string  `json:"secure_url"`
         
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
    
    // ActionQuery ...
    type ActionQuery struct {

        
            ProductSlug []string  `json:"product_slug"`
         
    }
    
    // ProductAction ...
    type ProductAction struct {

        
            URL string  `json:"url"`
            Type string  `json:"type"`
            Query ActionQuery  `json:"query"`
         
    }
    
    // NetQuantity ...
    type NetQuantity struct {

        
            Value string  `json:"value"`
            Unit string  `json:"unit"`
         
    }
    
    // CartProduct ...
    type CartProduct struct {

        
            Slug string  `json:"slug"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Images []ProductImage  `json:"images"`
            TeaserTag Tags  `json:"teaser_tag"`
            Brand BaseInfo  `json:"brand"`
            Type string  `json:"type"`
            Categories []CategoryInfo  `json:"categories"`
            Action ProductAction  `json:"action"`
            Name string  `json:"name"`
            NetQuantity NetQuantity  `json:"net_quantity"`
            ItemCode string  `json:"item_code"`
            UID float64  `json:"uid"`
            Tags []string  `json:"tags"`
         
    }
    
    // CartProductInfo ...
    type CartProductInfo struct {

        
            Moq map[string]interface{}  `json:"moq"`
            Availability ProductAvailability  `json:"availability"`
            PromoMeta PromoMeta  `json:"promo_meta"`
            Price ProductPriceInfo  `json:"price"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            CouponMessage string  `json:"coupon_message"`
            IsSet bool  `json:"is_set"`
            Discount string  `json:"discount"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
            Coupon CouponDetails  `json:"coupon"`
            PromotionsApplied []AppliedPromotion  `json:"promotions_applied"`
            Article ProductArticle  `json:"article"`
            Key string  `json:"key"`
            BulkOffer map[string]interface{}  `json:"bulk_offer"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            Quantity float64  `json:"quantity"`
            PricePerUnit ProductPriceInfo  `json:"price_per_unit"`
            Message string  `json:"message"`
            Product CartProduct  `json:"product"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
         
    }
    
    // DisplayBreakupSchema ...
    type DisplayBreakupSchema struct {

        
            CurrencySymbol string  `json:"currency_symbol"`
            CurrencyCode string  `json:"currency_code"`
            Key string  `json:"key"`
            Value float64  `json:"value"`
            Message []string  `json:"message"`
            Display string  `json:"display"`
         
    }
    
    // LoyaltyPoints ...
    type LoyaltyPoints struct {

        
            Applicable float64  `json:"applicable"`
            Total float64  `json:"total"`
            IsApplied bool  `json:"is_applied"`
            Description string  `json:"description"`
         
    }
    
    // RawBreakupSchema ...
    type RawBreakupSchema struct {

        
            Subtotal float64  `json:"subtotal"`
            Total float64  `json:"total"`
            ConvenienceFee float64  `json:"convenience_fee"`
            YouSaved float64  `json:"you_saved"`
            GiftCard float64  `json:"gift_card"`
            GstCharges float64  `json:"gst_charges"`
            Vog float64  `json:"vog"`
            MrpTotal float64  `json:"mrp_total"`
            CodCharge float64  `json:"cod_charge"`
            DeliveryCharge float64  `json:"delivery_charge"`
            Coupon float64  `json:"coupon"`
            Discount float64  `json:"discount"`
            FyndCash float64  `json:"fynd_cash"`
         
    }
    
    // CouponBreakupSchema ...
    type CouponBreakupSchema struct {

        
            Description string  `json:"description"`
            Code string  `json:"code"`
            MaxDiscountValue float64  `json:"max_discount_value"`
            Type string  `json:"type"`
            Title string  `json:"title"`
            MinimumCartValue float64  `json:"minimum_cart_value"`
            Value float64  `json:"value"`
            IsApplied bool  `json:"is_applied"`
            SubTitle string  `json:"sub_title"`
            Message string  `json:"message"`
            CouponType string  `json:"coupon_type"`
            UID string  `json:"uid"`
            CouponValue float64  `json:"coupon_value"`
         
    }
    
    // CartBreakupSchema ...
    type CartBreakupSchema struct {

        
            Display []DisplayBreakupSchema  `json:"display"`
            LoyaltyPoints LoyaltyPoints  `json:"loyalty_points"`
            Raw RawBreakupSchema  `json:"raw"`
            Coupon CouponBreakupSchema  `json:"coupon"`
         
    }
    
    // PaymentSelectionLock ...
    type PaymentSelectionLock struct {

        
            DefaultOptions string  `json:"default_options"`
            PaymentIdentifier string  `json:"payment_identifier"`
            Enabled bool  `json:"enabled"`
         
    }
    
    // CartDetailResponse ...
    type CartDetailResponse struct {

        
            CouponText string  `json:"coupon_text"`
            IsValid bool  `json:"is_valid"`
            Currency CartCurrencySchema  `json:"currency"`
            BuyNow bool  `json:"buy_now"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            CheckoutMode string  `json:"checkout_mode"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            ID string  `json:"id"`
            LastModified string  `json:"last_modified"`
            AppliedPromoDetails []AppliedPromotion  `json:"applied_promo_details"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            PanNo string  `json:"pan_no"`
            Items []CartProductInfo  `json:"items"`
            BreakupValues CartBreakupSchema  `json:"breakup_values"`
            Comment string  `json:"comment"`
            Gstin string  `json:"gstin"`
            Message string  `json:"message"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            PanConfig map[string]interface{}  `json:"pan_config"`
         
    }
    
    // AddProductCart ...
    type AddProductCart struct {

        
            ArticleID string  `json:"article_id"`
            ProductGroupTags []string  `json:"product_group_tags"`
            ParentItemIdentifiers []map[string]string  `json:"parent_item_identifiers"`
            StoreID float64  `json:"store_id"`
            Quantity float64  `json:"quantity"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ArticleAssignment map[string]interface{}  `json:"article_assignment"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            Pos bool  `json:"pos"`
            SellerID float64  `json:"seller_id"`
            ItemSize string  `json:"item_size"`
            ItemID float64  `json:"item_id"`
            Display string  `json:"display"`
         
    }
    
    // AddCartRequest ...
    type AddCartRequest struct {

        
            NewCart bool  `json:"new_cart"`
            Items []AddProductCart  `json:"items"`
         
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

        
            ArticleID string  `json:"article_id"`
            ItemIndex float64  `json:"item_index"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            Quantity float64  `json:"quantity"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            ItemSize string  `json:"item_size"`
            ItemID float64  `json:"item_id"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
         
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
    
    // PageCouponSchema ...
    type PageCouponSchema struct {

        
            Total float64  `json:"total"`
            HasNext bool  `json:"has_next"`
            TotalItemCount float64  `json:"total_item_count"`
            HasPrevious bool  `json:"has_previous"`
            Current float64  `json:"current"`
         
    }
    
    // Coupon ...
    type Coupon struct {

        
            ExpiresOn string  `json:"expires_on"`
            Description string  `json:"description"`
            CouponValue float64  `json:"coupon_value"`
            MaxDiscountValue float64  `json:"max_discount_value"`
            IsApplicable bool  `json:"is_applicable"`
            MinimumCartValue float64  `json:"minimum_cart_value"`
            Title string  `json:"title"`
            IsApplied bool  `json:"is_applied"`
            SubTitle string  `json:"sub_title"`
            Message string  `json:"message"`
            CouponType string  `json:"coupon_type"`
            CouponCode string  `json:"coupon_code"`
         
    }
    
    // GetCouponResponse ...
    type GetCouponResponse struct {

        
            Page PageCouponSchema  `json:"page"`
            AvailableCouponList []Coupon  `json:"available_coupon_list"`
         
    }
    
    // ApplyCouponRequest ...
    type ApplyCouponRequest struct {

        
            CouponCode string  `json:"coupon_code"`
         
    }
    
    // OfferSellerSchema ...
    type OfferSellerSchema struct {

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // OfferPrice ...
    type OfferPrice struct {

        
            CurrencySymbol string  `json:"currency_symbol"`
            BulkEffective float64  `json:"bulk_effective"`
            Marked float64  `json:"marked"`
            Effective float64  `json:"effective"`
            CurrencyCode string  `json:"currency_code"`
         
    }
    
    // OfferItemSchema ...
    type OfferItemSchema struct {

        
            AutoApplied bool  `json:"auto_applied"`
            Total float64  `json:"total"`
            Quantity float64  `json:"quantity"`
            Type string  `json:"type"`
            Margin float64  `json:"margin"`
            Best bool  `json:"best"`
            Price OfferPrice  `json:"price"`
         
    }
    
    // BulkPriceOfferSchema ...
    type BulkPriceOfferSchema struct {

        
            Seller OfferSellerSchema  `json:"seller"`
            Offers []OfferItemSchema  `json:"offers"`
         
    }
    
    // BulkPriceResponse ...
    type BulkPriceResponse struct {

        
            Data []BulkPriceOfferSchema  `json:"data"`
         
    }
    
    // RewardPointRequestSchema ...
    type RewardPointRequestSchema struct {

        
            Points bool  `json:"points"`
         
    }
    
    // GeoLocation ...
    type GeoLocation struct {

        
            Latitude float64  `json:"latitude"`
            Longitude float64  `json:"longitude"`
         
    }
    
    // Address ...
    type Address struct {

        
            AreaCode string  `json:"area_code"`
            IsDefaultAddress bool  `json:"is_default_address"`
            AddressType string  `json:"address_type"`
            Tags []string  `json:"tags"`
            CountryIsoCode string  `json:"country_iso_code"`
            Phone string  `json:"phone"`
            UserID string  `json:"user_id"`
            CheckoutMode string  `json:"checkout_mode"`
            Landmark string  `json:"landmark"`
            City string  `json:"city"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ID string  `json:"id"`
            Country string  `json:"country"`
            Name string  `json:"name"`
            Address string  `json:"address"`
            GoogleMapPoint map[string]interface{}  `json:"google_map_point"`
            State string  `json:"state"`
            AreaCodeSlug string  `json:"area_code_slug"`
            IsActive bool  `json:"is_active"`
            Area string  `json:"area"`
            Meta map[string]interface{}  `json:"meta"`
            GeoLocation GeoLocation  `json:"geo_location"`
            CountryCode string  `json:"country_code"`
            CreatedByUserID string  `json:"created_by_user_id"`
            CountryPhoneCode string  `json:"country_phone_code"`
            Email string  `json:"email"`
         
    }
    
    // GetAddressesResponse ...
    type GetAddressesResponse struct {

        
            Address []Address  `json:"address"`
            PiiMasking bool  `json:"pii_masking"`
         
    }
    
    // SaveAddressResponse ...
    type SaveAddressResponse struct {

        
            Success bool  `json:"success"`
            IsDefaultAddress bool  `json:"is_default_address"`
            ID string  `json:"id"`
         
    }
    
    // UpdateAddressResponse ...
    type UpdateAddressResponse struct {

        
            Success bool  `json:"success"`
            IsDefaultAddress bool  `json:"is_default_address"`
            IsUpdated bool  `json:"is_updated"`
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

        
            PaymentIdentifier string  `json:"payment_identifier"`
            PaymentMode string  `json:"payment_mode"`
            ID string  `json:"id"`
            AggregatorName string  `json:"aggregator_name"`
            MerchantCode string  `json:"merchant_code"`
            AddressID string  `json:"address_id"`
         
    }
    
    // CouponValidity ...
    type CouponValidity struct {

        
            Valid bool  `json:"valid"`
            Code string  `json:"code"`
            Title string  `json:"title"`
            DisplayMessageEn string  `json:"display_message_en"`
            Discount float64  `json:"discount"`
         
    }
    
    // PaymentCouponValidateSchema ...
    type PaymentCouponValidateSchema struct {

        
            CouponValidity CouponValidity  `json:"coupon_validity"`
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // ShipmentResponse ...
    type ShipmentResponse struct {

        
            BoxType string  `json:"box_type"`
            Items []CartProductInfo  `json:"items"`
            OrderType string  `json:"order_type"`
            Shipments float64  `json:"shipments"`
            FulfillmentType string  `json:"fulfillment_type"`
            DpID string  `json:"dp_id"`
            ShipmentType string  `json:"shipment_type"`
            DpOptions map[string]interface{}  `json:"dp_options"`
            Promise ShipmentPromise  `json:"promise"`
            FulfillmentID float64  `json:"fulfillment_id"`
         
    }
    
    // CartShipmentsResponse ...
    type CartShipmentsResponse struct {

        
            CouponText string  `json:"coupon_text"`
            IsValid bool  `json:"is_valid"`
            Shipments []ShipmentResponse  `json:"shipments"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            ID string  `json:"id"`
            LastModified string  `json:"last_modified"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            Error bool  `json:"error"`
            Currency CartCurrencySchema  `json:"currency"`
            BuyNow bool  `json:"buy_now"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            BreakupValues CartBreakupSchema  `json:"breakup_values"`
            CartID float64  `json:"cart_id"`
            Comment string  `json:"comment"`
            CheckoutMode string  `json:"checkout_mode"`
            Gstin string  `json:"gstin"`
            Message string  `json:"message"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            UID string  `json:"uid"`
         
    }
    
    // CartCheckoutCustomMetaSchema ...
    type CartCheckoutCustomMetaSchema struct {

        
            Value string  `json:"value"`
            Key string  `json:"key"`
         
    }
    
    // StaffCheckoutSchema ...
    type StaffCheckoutSchema struct {

        
            FirstName string  `json:"first_name"`
            EmployeeCode string  `json:"employee_code"`
            ID string  `json:"_id"`
            LastName string  `json:"last_name"`
            User string  `json:"user"`
         
    }
    
    // CartCheckoutDetailRequest ...
    type CartCheckoutDetailRequest struct {

        
            Meta map[string]interface{}  `json:"meta"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
            CustomMeta []CartCheckoutCustomMetaSchema  `json:"custom_meta"`
            BillingAddressID string  `json:"billing_address_id"`
            OrderType string  `json:"order_type"`
            CallbackURL string  `json:"callback_url"`
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
            OrderingStore float64  `json:"ordering_store"`
            PaymentIdentifier string  `json:"payment_identifier"`
            ID string  `json:"id"`
            PaymentMode string  `json:"payment_mode"`
            PaymentParams map[string]interface{}  `json:"payment_params"`
            PaymentAutoConfirm bool  `json:"payment_auto_confirm"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            Staff StaffCheckoutSchema  `json:"staff"`
            MerchantCode string  `json:"merchant_code"`
            AddressID string  `json:"address_id"`
            Aggregator string  `json:"aggregator"`
         
    }
    
    // CheckCartSchema ...
    type CheckCartSchema struct {

        
            DeliveryChargeOrderValue float64  `json:"delivery_charge_order_value"`
            CouponText string  `json:"coupon_text"`
            IsValid bool  `json:"is_valid"`
            Success bool  `json:"success"`
            CodAvailable bool  `json:"cod_available"`
            Currency CartCurrencySchema  `json:"currency"`
            StoreEmps []map[string]interface{}  `json:"store_emps"`
            BuyNow bool  `json:"buy_now"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            CodMessage string  `json:"cod_message"`
            OrderID string  `json:"order_id"`
            CodCharges float64  `json:"cod_charges"`
            CheckoutMode string  `json:"checkout_mode"`
            StoreCode string  `json:"store_code"`
            UserType string  `json:"user_type"`
            UID string  `json:"uid"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            ID string  `json:"id"`
            LastModified string  `json:"last_modified"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            Items []CartProductInfo  `json:"items"`
            BreakupValues CartBreakupSchema  `json:"breakup_values"`
            CartID float64  `json:"cart_id"`
            Comment string  `json:"comment"`
            DeliveryCharges float64  `json:"delivery_charges"`
            ErrorMessage string  `json:"error_message"`
            Gstin string  `json:"gstin"`
            Message string  `json:"message"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
         
    }
    
    // CartCheckoutResponseSchema ...
    type CartCheckoutResponseSchema struct {

        
            PaymentConfirmURL string  `json:"payment_confirm_url"`
            CallbackURL string  `json:"callback_url"`
            Success bool  `json:"success"`
            Data map[string]interface{}  `json:"data"`
            OrderID string  `json:"order_id"`
            Cart CheckCartSchema  `json:"cart"`
            AppInterceptURL string  `json:"app_intercept_url"`
            Message string  `json:"message"`
         
    }
    
    // GiftDetailSchema ...
    type GiftDetailSchema struct {

        
            IsGiftApplied bool  `json:"is_gift_applied"`
            GiftMessage string  `json:"gift_message"`
         
    }
    
    // ArticleGiftDetailSchema ...
    type ArticleGiftDetailSchema struct {

        
            ArticleID GiftDetailSchema  `json:"article_id"`
         
    }
    
    // CartMetaRequest ...
    type CartMetaRequest struct {

        
            GiftDetails ArticleGiftDetailSchema  `json:"gift_details"`
            Comment string  `json:"comment"`
            CheckoutMode string  `json:"checkout_mode"`
            PickUpCustomerDetails map[string]interface{}  `json:"pick_up_customer_details"`
            Gstin string  `json:"gstin"`
         
    }
    
    // CartMetaResponse ...
    type CartMetaResponse struct {

        
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

        
            Meta map[string]interface{}  `json:"meta"`
            Source map[string]interface{}  `json:"source"`
            CreatedOn string  `json:"created_on"`
            User map[string]interface{}  `json:"user"`
            Token string  `json:"token"`
         
    }
    
    // SharedCart ...
    type SharedCart struct {

        
            CouponText string  `json:"coupon_text"`
            IsValid bool  `json:"is_valid"`
            Currency CartCurrencySchema  `json:"currency"`
            BuyNow bool  `json:"buy_now"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            SharedCartDetails SharedCartDetails  `json:"shared_cart_details"`
            CheckoutMode string  `json:"checkout_mode"`
            UID string  `json:"uid"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            ID string  `json:"id"`
            LastModified string  `json:"last_modified"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            Items []CartProductInfo  `json:"items"`
            BreakupValues CartBreakupSchema  `json:"breakup_values"`
            CartID float64  `json:"cart_id"`
            Comment string  `json:"comment"`
            Gstin string  `json:"gstin"`
            Message string  `json:"message"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
         
    }
    
    // SharedCartResponse ...
    type SharedCartResponse struct {

        
            Cart SharedCart  `json:"cart"`
            Error string  `json:"error"`
         
    }
    
    // FreeGiftItemsSchema ...
    type FreeGiftItemsSchema struct {

        
            ItemBrandName string  `json:"item_brand_name"`
            ItemPriceDetails map[string]interface{}  `json:"item_price_details"`
            ItemName string  `json:"item_name"`
            ItemSlug string  `json:"item_slug"`
            ItemID float64  `json:"item_id"`
            ItemImagesURL []string  `json:"item_images_url"`
         
    }
    
    // PromotionOffer ...
    type PromotionOffer struct {

        
            OfferText string  `json:"offer_text"`
            Description string  `json:"description"`
            BuyRules map[string]interface{}  `json:"buy_rules"`
            PromotionGroup string  `json:"promotion_group"`
            FreeGiftItems []FreeGiftItemsSchema  `json:"free_gift_items"`
            ID string  `json:"id"`
            DiscountRules []map[string]interface{}  `json:"discount_rules"`
            ValidTill string  `json:"valid_till"`
         
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
    
    // LadderPrice ...
    type LadderPrice struct {

        
            CurrencySymbol string  `json:"currency_symbol"`
            OfferPrice float64  `json:"offer_price"`
            Marked float64  `json:"marked"`
            Effective float64  `json:"effective"`
            CurrencyCode string  `json:"currency_code"`
         
    }
    
    // LadderOfferItemSchema ...
    type LadderOfferItemSchema struct {

        
            Type string  `json:"type"`
            MaxQuantity float64  `json:"max_quantity"`
            MinQuantity float64  `json:"min_quantity"`
            Margin float64  `json:"margin"`
            Price LadderPrice  `json:"price"`
         
    }
    
    // LadderPriceOffer ...
    type LadderPriceOffer struct {

        
            OfferText string  `json:"offer_text"`
            CalculateOn string  `json:"calculate_on"`
            Description string  `json:"description"`
            BuyRules map[string]interface{}  `json:"buy_rules"`
            PromotionGroup string  `json:"promotion_group"`
            FreeGiftItems []FreeGiftItemsSchema  `json:"free_gift_items"`
            ID string  `json:"id"`
            DiscountRules []map[string]interface{}  `json:"discount_rules"`
            OfferPrices []LadderOfferItemSchema  `json:"offer_prices"`
            ValidTill string  `json:"valid_till"`
         
    }
    
    // CurrencyInfoSchema ...
    type CurrencyInfoSchema struct {

        
            Symbol string  `json:"symbol"`
            Code string  `json:"code"`
         
    }
    
    // LadderPriceOffers ...
    type LadderPriceOffers struct {

        
            AvailableOffers []LadderPriceOffer  `json:"available_offers"`
            Currency CurrencyInfoSchema  `json:"currency"`
         
    }
    
    // PaymentMetaSchema ...
    type PaymentMetaSchema struct {

        
            PaymentIdentifier string  `json:"payment_identifier"`
            Type string  `json:"type"`
            MerchantCode string  `json:"merchant_code"`
            PaymentGateway string  `json:"payment_gateway"`
         
    }
    
    // PaymentMethodSchema ...
    type PaymentMethodSchema struct {

        
            Amount float64  `json:"amount"`
            Mode string  `json:"mode"`
            PaymentMeta PaymentMetaSchema  `json:"payment_meta"`
            Name string  `json:"name"`
            Payment string  `json:"payment"`
         
    }
    
    // CartCheckoutDetailV2Request ...
    type CartCheckoutDetailV2Request struct {

        
            CallbackURL string  `json:"callback_url"`
            PaymentIdentifier string  `json:"payment_identifier"`
            PaymentAutoConfirm bool  `json:"payment_auto_confirm"`
            Staff StaffCheckoutSchema  `json:"staff"`
            MerchantCode string  `json:"merchant_code"`
            PaymentMethods []PaymentMethodSchema  `json:"payment_methods"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
            BillingAddressID string  `json:"billing_address_id"`
            OrderType string  `json:"order_type"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            AddressID string  `json:"address_id"`
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
            ID string  `json:"id"`
            Meta map[string]interface{}  `json:"meta"`
            CustomMeta map[string]interface{}  `json:"custom_meta"`
            CartID string  `json:"cart_id"`
            OrderingStore float64  `json:"ordering_store"`
            PaymentMode string  `json:"payment_mode"`
            PaymentParams map[string]interface{}  `json:"payment_params"`
            Aggregator string  `json:"aggregator"`
         
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
            Media SlideshowMedia  `json:"media"`
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

        
            Secret string  `json:"secret"`
            Sdk bool  `json:"sdk"`
            UserID string  `json:"user_id"`
            ConfigType string  `json:"config_type"`
            VerifyAPI string  `json:"verify_api"`
            MerchantID string  `json:"merchant_id"`
            MerchantKey string  `json:"merchant_key"`
            API string  `json:"api"`
            Key string  `json:"key"`
            Pin string  `json:"pin"`
         
    }
    
    // AggregatorsConfigDetailResponse ...
    type AggregatorsConfigDetailResponse struct {

        
            Success bool  `json:"success"`
            Simpl AggregatorConfigDetail  `json:"simpl"`
            Ccavenue AggregatorConfigDetail  `json:"ccavenue"`
            Env string  `json:"env"`
            Stripe AggregatorConfigDetail  `json:"stripe"`
            Rupifi AggregatorConfigDetail  `json:"rupifi"`
            Razorpay AggregatorConfigDetail  `json:"razorpay"`
            Juspay AggregatorConfigDetail  `json:"juspay"`
            Payumoney AggregatorConfigDetail  `json:"payumoney"`
            Mswipe AggregatorConfigDetail  `json:"mswipe"`
         
    }
    
    // ErrorCodeAndDescription ...
    type ErrorCodeAndDescription struct {

        
            Code string  `json:"code"`
            Description string  `json:"description"`
         
    }
    
    // HttpErrorCodeAndResponse ...
    type HttpErrorCodeAndResponse struct {

        
            Success bool  `json:"success"`
            Error ErrorCodeAndDescription  `json:"error"`
         
    }
    
    // AttachCardRequest ...
    type AttachCardRequest struct {

        
            NameOnCard string  `json:"name_on_card"`
            Nickname string  `json:"nickname"`
            Refresh bool  `json:"refresh"`
            CardID string  `json:"card_id"`
         
    }
    
    // AttachCardsResponse ...
    type AttachCardsResponse struct {

        
            Success bool  `json:"success"`
            Data map[string]interface{}  `json:"data"`
            Message string  `json:"message"`
         
    }
    
    // CardPaymentGateway ...
    type CardPaymentGateway struct {

        
            CustomerID string  `json:"customer_id"`
            API string  `json:"api"`
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

        
            AggregatorName string  `json:"aggregator_name"`
            CardBrand string  `json:"card_brand"`
            CardNumber string  `json:"card_number"`
            CardID string  `json:"card_id"`
            CardReference string  `json:"card_reference"`
            ExpYear float64  `json:"exp_year"`
            Nickname string  `json:"nickname"`
            CardType string  `json:"card_type"`
            CardIssuer string  `json:"card_issuer"`
            CompliantWithTokenisationGuidelines bool  `json:"compliant_with_tokenisation_guidelines"`
            CardBrandImage string  `json:"card_brand_image"`
            CardFingerprint string  `json:"card_fingerprint"`
            CardIsin string  `json:"card_isin"`
            CardToken string  `json:"card_token"`
            CardName string  `json:"card_name"`
            Expired bool  `json:"expired"`
            ExpMonth float64  `json:"exp_month"`
         
    }
    
    // ListCardsResponse ...
    type ListCardsResponse struct {

        
            Success bool  `json:"success"`
            Data []Card  `json:"data"`
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

        
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
            Payload string  `json:"payload"`
            Aggregator string  `json:"aggregator"`
            MerchantParams map[string]interface{}  `json:"merchant_params"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
            PhoneNumber string  `json:"phone_number"`
            OrderItems []map[string]interface{}  `json:"order_items"`
            TransactionAmountInPaise float64  `json:"transaction_amount_in_paise"`
         
    }
    
    // ValidateCustomerResponse ...
    type ValidateCustomerResponse struct {

        
            Success bool  `json:"success"`
            Data map[string]interface{}  `json:"data"`
            Error map[string]interface{}  `json:"error"`
            Message string  `json:"message"`
         
    }
    
    // ChargeCustomerRequest ...
    type ChargeCustomerRequest struct {

        
            OrderID string  `json:"order_id"`
            Aggregator string  `json:"aggregator"`
            Amount float64  `json:"amount"`
            TransactionToken string  `json:"transaction_token"`
            Verified bool  `json:"verified"`
         
    }
    
    // ChargeCustomerResponse ...
    type ChargeCustomerResponse struct {

        
            Success bool  `json:"success"`
            OrderID string  `json:"order_id"`
            Aggregator string  `json:"aggregator"`
            Message string  `json:"message"`
            DeliveryAddressID string  `json:"delivery_address_id"`
            CartID string  `json:"cart_id"`
            Status string  `json:"status"`
         
    }
    
    // PaymentInitializationRequest ...
    type PaymentInitializationRequest struct {

        
            OrderID string  `json:"order_id"`
            Amount float64  `json:"amount"`
            Aggregator string  `json:"aggregator"`
            RazorpayPaymentID string  `json:"razorpay_payment_id"`
            Vpa string  `json:"vpa"`
            CustomerID string  `json:"customer_id"`
            MerchantOrderID string  `json:"merchant_order_id"`
            Email string  `json:"email"`
            Timeout float64  `json:"timeout"`
            Contact string  `json:"contact"`
            Method string  `json:"method"`
            DeviceID string  `json:"device_id"`
            Currency string  `json:"currency"`
         
    }
    
    // PaymentInitializationResponse ...
    type PaymentInitializationResponse struct {

        
            PaymentID string  `json:"payment_id"`
            Amount float64  `json:"amount"`
            OrderID string  `json:"order_id"`
            Contact string  `json:"contact"`
            Currency string  `json:"currency"`
            Success bool  `json:"success"`
            Vpa string  `json:"vpa"`
            UpiPollURL string  `json:"upi_poll_url"`
            Timeout float64  `json:"timeout"`
            BqrImage string  `json:"bqr_image"`
            AggregatorOrderID string  `json:"aggregator_order_id"`
            Status string  `json:"status"`
            RazorpayPaymentID string  `json:"razorpay_payment_id"`
            MerchantOrderID string  `json:"merchant_order_id"`
            Email string  `json:"email"`
            Method string  `json:"method"`
            DeviceID string  `json:"device_id"`
            PollingURL string  `json:"polling_url"`
            Aggregator string  `json:"aggregator"`
            CustomerID string  `json:"customer_id"`
            VirtualID string  `json:"virtual_id"`
         
    }
    
    // PaymentStatusUpdateRequest ...
    type PaymentStatusUpdateRequest struct {

        
            OrderID string  `json:"order_id"`
            Aggregator string  `json:"aggregator"`
            Amount float64  `json:"amount"`
            Vpa string  `json:"vpa"`
            RazorpayPaymentID string  `json:"razorpay_payment_id"`
            CustomerID string  `json:"customer_id"`
            PaymentID string  `json:"payment_id"`
            MerchantOrderID string  `json:"merchant_order_id"`
            Email string  `json:"email"`
            MerchantTransactionID string  `json:"merchant_transaction_id"`
            Contact string  `json:"contact"`
            Method string  `json:"method"`
            DeviceID string  `json:"device_id"`
            Status string  `json:"status"`
            Currency string  `json:"currency"`
         
    }
    
    // PaymentStatusUpdateResponse ...
    type PaymentStatusUpdateResponse struct {

        
            Retry bool  `json:"retry"`
            Success bool  `json:"success"`
            AggregatorName string  `json:"aggregator_name"`
            RedirectURL string  `json:"redirect_url"`
            Status string  `json:"status"`
         
    }
    
    // AggregatorRoute ...
    type AggregatorRoute struct {

        
            Data map[string]interface{}  `json:"data"`
            PaymentFlow string  `json:"payment_flow"`
            PaymentFlowData string  `json:"payment_flow_data"`
            APILink string  `json:"api_link"`
         
    }
    
    // PaymentFlow ...
    type PaymentFlow struct {

        
            Simpl AggregatorRoute  `json:"simpl"`
            Ccavenue AggregatorRoute  `json:"ccavenue"`
            Fynd AggregatorRoute  `json:"fynd"`
            Payubiz AggregatorRoute  `json:"payubiz"`
            Stripe AggregatorRoute  `json:"stripe"`
            Rupifi AggregatorRoute  `json:"rupifi"`
            Razorpay AggregatorRoute  `json:"razorpay"`
            Juspay AggregatorRoute  `json:"juspay"`
            Epaylater AggregatorRoute  `json:"epaylater"`
            UpiRazorpay AggregatorRoute  `json:"upi_razorpay"`
            BqrRazorpay AggregatorRoute  `json:"bqr_razorpay"`
            Jiopay AggregatorRoute  `json:"jiopay"`
            Mswipe AggregatorRoute  `json:"mswipe"`
         
    }
    
    // IntentAppErrorList ...
    type IntentAppErrorList struct {

        
            Code string  `json:"code"`
            PackageName string  `json:"package_name"`
         
    }
    
    // PaymentModeLogo ...
    type PaymentModeLogo struct {

        
            Large string  `json:"large"`
            Small string  `json:"small"`
         
    }
    
    // IntentApp ...
    type IntentApp struct {

        
            Code string  `json:"code"`
            PackageName string  `json:"package_name"`
            Logos PaymentModeLogo  `json:"logos"`
            DisplayName string  `json:"display_name"`
         
    }
    
    // PaymentModeList ...
    type PaymentModeList struct {

        
            AggregatorName string  `json:"aggregator_name"`
            CardReference string  `json:"card_reference"`
            CardType string  `json:"card_type"`
            IntentAppErrorDictList []IntentAppErrorList  `json:"intent_app_error_dict_list"`
            CodLimit float64  `json:"cod_limit"`
            CodLimitPerOrder float64  `json:"cod_limit_per_order"`
            Expired bool  `json:"expired"`
            DisplayName string  `json:"display_name"`
            Code string  `json:"code"`
            CompliantWithTokenisationGuidelines bool  `json:"compliant_with_tokenisation_guidelines"`
            ExpYear float64  `json:"exp_year"`
            CardIssuer string  `json:"card_issuer"`
            Timeout float64  `json:"timeout"`
            IntentApp []IntentApp  `json:"intent_app"`
            CardToken string  `json:"card_token"`
            CardName string  `json:"card_name"`
            CardBrandImage string  `json:"card_brand_image"`
            RemainingLimit float64  `json:"remaining_limit"`
            RetryCount float64  `json:"retry_count"`
            CardID string  `json:"card_id"`
            Name string  `json:"name"`
            DisplayPriority float64  `json:"display_priority"`
            IntentAppErrorList []string  `json:"intent_app_error_list"`
            CardIsin string  `json:"card_isin"`
            ExpMonth float64  `json:"exp_month"`
            IntentFlow bool  `json:"intent_flow"`
            CardBrand string  `json:"card_brand"`
            CardNumber string  `json:"card_number"`
            Nickname string  `json:"nickname"`
            LogoURL PaymentModeLogo  `json:"logo_url"`
            FyndVpa string  `json:"fynd_vpa"`
            CardFingerprint string  `json:"card_fingerprint"`
            MerchantCode string  `json:"merchant_code"`
         
    }
    
    // RootPaymentMode ...
    type RootPaymentMode struct {

        
            AggregatorName string  `json:"aggregator_name"`
            Name string  `json:"name"`
            SaveCard bool  `json:"save_card"`
            DisplayPriority float64  `json:"display_priority"`
            IsPayByCardPl bool  `json:"is_pay_by_card_pl"`
            List []PaymentModeList  `json:"list"`
            AddCardEnabled bool  `json:"add_card_enabled"`
            AnonymousEnable bool  `json:"anonymous_enable"`
            DisplayName string  `json:"display_name"`
         
    }
    
    // PaymentOptionAndFlow ...
    type PaymentOptionAndFlow struct {

        
            PaymentFlows PaymentFlow  `json:"payment_flows"`
            PaymentOption []RootPaymentMode  `json:"payment_option"`
         
    }
    
    // PaymentModeRouteResponse ...
    type PaymentModeRouteResponse struct {

        
            Success bool  `json:"success"`
            PaymentOptions PaymentOptionAndFlow  `json:"payment_options"`
         
    }
    
    // RupifiBannerData ...
    type RupifiBannerData struct {

        
            KycURL string  `json:"kyc_url"`
            Status string  `json:"status"`
         
    }
    
    // RupifiBannerResponse ...
    type RupifiBannerResponse struct {

        
            Success bool  `json:"success"`
            Data RupifiBannerData  `json:"data"`
         
    }
    
    // EpaylaterBannerData ...
    type EpaylaterBannerData struct {

        
            Message string  `json:"message"`
            Display bool  `json:"display"`
            Status string  `json:"status"`
         
    }
    
    // EpaylaterBannerResponse ...
    type EpaylaterBannerResponse struct {

        
            Success bool  `json:"success"`
            Data EpaylaterBannerData  `json:"data"`
         
    }
    
    // ResendOrCancelPaymentRequest ...
    type ResendOrCancelPaymentRequest struct {

        
            OrderID string  `json:"order_id"`
            RequestType string  `json:"request_type"`
            DeviceID string  `json:"device_id"`
         
    }
    
    // LinkStatus ...
    type LinkStatus struct {

        
            Status bool  `json:"status"`
            Message string  `json:"message"`
         
    }
    
    // ResendOrCancelPaymentResponse ...
    type ResendOrCancelPaymentResponse struct {

        
            Success bool  `json:"success"`
            Data LinkStatus  `json:"data"`
         
    }
    
    // renderHTMLRequest ...
    type renderHTMLRequest struct {

        
            Returntype string  `json:"returntype"`
            Base64Html string  `json:"base64_html"`
         
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

        
            CustomerName string  `json:"customer_name"`
            UpiVpa string  `json:"upi_vpa"`
            Status string  `json:"status"`
            IsValid bool  `json:"is_valid"`
         
    }
    
    // ValidateVPAResponse ...
    type ValidateVPAResponse struct {

        
            Success bool  `json:"success"`
            Data ValidateUPI  `json:"data"`
         
    }
    
    // TransferItemsDetails ...
    type TransferItemsDetails struct {

        
            Name string  `json:"name"`
            ID float64  `json:"id"`
            LogoLarge string  `json:"logo_large"`
            LogoSmall string  `json:"logo_small"`
            DisplayName string  `json:"display_name"`
         
    }
    
    // TransferModeDetails ...
    type TransferModeDetails struct {

        
            Items []TransferItemsDetails  `json:"items"`
            DisplayName string  `json:"display_name"`
         
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

        
            AccountHolder string  `json:"account_holder"`
            BankName string  `json:"bank_name"`
            AccountNo string  `json:"account_no"`
            DisplayName string  `json:"display_name"`
            CreatedOn string  `json:"created_on"`
            TransferMode string  `json:"transfer_mode"`
            BranchName string  `json:"branch_name"`
            ID float64  `json:"id"`
            IsActive bool  `json:"is_active"`
            ModifiedOn string  `json:"modified_on"`
            BeneficiaryID string  `json:"beneficiary_id"`
            Subtitle string  `json:"subtitle"`
            Email string  `json:"email"`
            IfscCode string  `json:"ifsc_code"`
            Mobile string  `json:"mobile"`
            Comment string  `json:"comment"`
            DelightsUserName string  `json:"delights_user_name"`
            Title string  `json:"title"`
            Address string  `json:"address"`
         
    }
    
    // OrderBeneficiaryResponse ...
    type OrderBeneficiaryResponse struct {

        
            ShowBeneficiaryDetails bool  `json:"show_beneficiary_details"`
            Beneficiaries []OrderBeneficiaryDetails  `json:"beneficiaries"`
         
    }
    
    // NotFoundResourceError ...
    type NotFoundResourceError struct {

        
            Code string  `json:"code"`
            Success bool  `json:"success"`
            Description string  `json:"description"`
         
    }
    
    // IfscCodeResponse ...
    type IfscCodeResponse struct {

        
            BranchName string  `json:"branch_name"`
            BankName string  `json:"bank_name"`
            Success bool  `json:"success"`
         
    }
    
    // ErrorCodeDescription ...
    type ErrorCodeDescription struct {

        
            Code string  `json:"code"`
            Success bool  `json:"success"`
            Description string  `json:"description"`
         
    }
    
    // AddBeneficiaryViaOtpVerificationRequest ...
    type AddBeneficiaryViaOtpVerificationRequest struct {

        
            RequestID string  `json:"request_id"`
            HashKey string  `json:"hash_key"`
            Otp string  `json:"otp"`
         
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

        
            Wallet string  `json:"wallet"`
            Vpa string  `json:"vpa"`
            AccountHolder string  `json:"account_holder"`
            BranchName string  `json:"branch_name"`
            BankName string  `json:"bank_name"`
            AccountNo string  `json:"account_no"`
            Email string  `json:"email"`
            Address string  `json:"address"`
            IfscCode string  `json:"ifsc_code"`
            Mobile string  `json:"mobile"`
            Comment string  `json:"comment"`
         
    }
    
    // AddBeneficiaryDetailsRequest ...
    type AddBeneficiaryDetailsRequest struct {

        
            OrderID string  `json:"order_id"`
            Details BeneficiaryModeDetails  `json:"details"`
            ShipmentID string  `json:"shipment_id"`
            Otp string  `json:"otp"`
            RequestID string  `json:"request_id"`
            Delights bool  `json:"delights"`
            TransferMode string  `json:"transfer_mode"`
         
    }
    
    // RefundAccountResponse ...
    type RefundAccountResponse struct {

        
            Success bool  `json:"success"`
            Data map[string]interface{}  `json:"data"`
            Message string  `json:"message"`
            IsVerifiedFlag bool  `json:"is_verified_flag"`
         
    }
    
    // BankDetailsForOTP ...
    type BankDetailsForOTP struct {

        
            AccountHolder string  `json:"account_holder"`
            BranchName string  `json:"branch_name"`
            BankName string  `json:"bank_name"`
            AccountNo string  `json:"account_no"`
            IfscCode string  `json:"ifsc_code"`
         
    }
    
    // AddBeneficiaryDetailsOTPRequest ...
    type AddBeneficiaryDetailsOTPRequest struct {

        
            OrderID string  `json:"order_id"`
            Details BankDetailsForOTP  `json:"details"`
         
    }
    
    // WalletOtpRequest ...
    type WalletOtpRequest struct {

        
            CountryCode string  `json:"country_code"`
            Mobile string  `json:"mobile"`
         
    }
    
    // WalletOtpResponse ...
    type WalletOtpResponse struct {

        
            Success bool  `json:"success"`
            RequestID string  `json:"request_id"`
            IsVerifiedFlag string  `json:"is_verified_flag"`
         
    }
    
    // SetDefaultBeneficiaryRequest ...
    type SetDefaultBeneficiaryRequest struct {

        
            BeneficiaryID string  `json:"beneficiary_id"`
            OrderID string  `json:"order_id"`
         
    }
    
    // SetDefaultBeneficiaryResponse ...
    type SetDefaultBeneficiaryResponse struct {

        
            Success bool  `json:"success"`
            IsBeneficiarySet bool  `json:"is_beneficiary_set"`
         
    }
    
    // GetPaymentLinkResponse ...
    type GetPaymentLinkResponse struct {

        
            Success bool  `json:"success"`
            PaymentLinkURL string  `json:"payment_link_url"`
            Amount float64  `json:"amount"`
            Message string  `json:"message"`
            MerchantName string  `json:"merchant_name"`
            PaymentLinkCurrentStatus string  `json:"payment_link_current_status"`
            PollingTimeout float64  `json:"polling_timeout"`
            ExternalOrderID string  `json:"external_order_id"`
            StatusCode float64  `json:"status_code"`
         
    }
    
    // ErrorDescription ...
    type ErrorDescription struct {

        
            InvalidID bool  `json:"invalid_id"`
            PaymentTransactionID string  `json:"payment_transaction_id"`
            Amount float64  `json:"amount"`
            Msg string  `json:"msg"`
            MerchantName string  `json:"merchant_name"`
            MerchantOrderID string  `json:"merchant_order_id"`
            Cancelled bool  `json:"cancelled"`
            Expired bool  `json:"expired"`
         
    }
    
    // ErrorResponse ...
    type ErrorResponse struct {

        
            Success bool  `json:"success"`
            StatusCode float64  `json:"status_code"`
            Error ErrorDescription  `json:"error"`
            Message string  `json:"message"`
         
    }
    
    // CreatePaymentLinkMeta ...
    type CreatePaymentLinkMeta struct {

        
            Amount string  `json:"amount"`
            AssignCardID string  `json:"assign_card_id"`
            Pincode string  `json:"pincode"`
            CartID string  `json:"cart_id"`
            CheckoutMode string  `json:"checkout_mode"`
         
    }
    
    // CreatePaymentLinkRequest ...
    type CreatePaymentLinkRequest struct {

        
            Amount float64  `json:"amount"`
            MobileNumber string  `json:"mobile_number"`
            Email string  `json:"email"`
            Description string  `json:"description"`
            ExternalOrderID string  `json:"external_order_id"`
            Meta CreatePaymentLinkMeta  `json:"meta"`
         
    }
    
    // CreatePaymentLinkResponse ...
    type CreatePaymentLinkResponse struct {

        
            Success bool  `json:"success"`
            PaymentLinkURL string  `json:"payment_link_url"`
            Message string  `json:"message"`
            PaymentLinkID string  `json:"payment_link_id"`
            PollingTimeout float64  `json:"polling_timeout"`
            StatusCode float64  `json:"status_code"`
         
    }
    
    // CancelOrResendPaymentLinkRequest ...
    type CancelOrResendPaymentLinkRequest struct {

        
            PaymentLinkID string  `json:"payment_link_id"`
         
    }
    
    // ResendPaymentLinkResponse ...
    type ResendPaymentLinkResponse struct {

        
            Success bool  `json:"success"`
            PollingTimeout float64  `json:"polling_timeout"`
            StatusCode float64  `json:"status_code"`
            Message string  `json:"message"`
         
    }
    
    // CancelPaymentLinkResponse ...
    type CancelPaymentLinkResponse struct {

        
            Success bool  `json:"success"`
            StatusCode float64  `json:"status_code"`
            Message string  `json:"message"`
         
    }
    
    // PollingPaymentLinkResponse ...
    type PollingPaymentLinkResponse struct {

        
            Success bool  `json:"success"`
            AggregatorName string  `json:"aggregator_name"`
            Amount float64  `json:"amount"`
            Message string  `json:"message"`
            PaymentLinkID string  `json:"payment_link_id"`
            OrderID string  `json:"order_id"`
            RedirectURL string  `json:"redirect_url"`
            HttpStatus float64  `json:"http_status"`
            StatusCode float64  `json:"status_code"`
            Status string  `json:"status"`
         
    }
    
    // PaymentMethodsMeta ...
    type PaymentMethodsMeta struct {

        
            PaymentGateway string  `json:"payment_gateway"`
            MerchantCode string  `json:"merchant_code"`
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

        
            PaymentLinkID string  `json:"payment_link_id"`
            FailureCallbackURL string  `json:"failure_callback_url"`
            SuccessCallbackURL string  `json:"success_callback_url"`
            Meta map[string]interface{}  `json:"meta"`
            PaymentMethods CreateOrderUserPaymentMethods  `json:"payment_methods"`
            Currency string  `json:"currency"`
         
    }
    
    // CreateOrderUserData ...
    type CreateOrderUserData struct {

        
            OrderID string  `json:"order_id"`
            Amount float64  `json:"amount"`
            Aggregator string  `json:"aggregator"`
            CustomerID string  `json:"customer_id"`
            MerchantOrderID string  `json:"merchant_order_id"`
            CallbackURL string  `json:"callback_url"`
            Email string  `json:"email"`
            Contact string  `json:"contact"`
            Method string  `json:"method"`
            Currency string  `json:"currency"`
         
    }
    
    // CreateOrderUserResponse ...
    type CreateOrderUserResponse struct {

        
            Success bool  `json:"success"`
            OrderID string  `json:"order_id"`
            Message string  `json:"message"`
            PaymentConfirmURL string  `json:"payment_confirm_url"`
            CallbackURL string  `json:"callback_url"`
            Data CreateOrderUserData  `json:"data"`
            StatusCode float64  `json:"status_code"`
         
    }
    
    // BalanceDetails ...
    type BalanceDetails struct {

        
            Value float64  `json:"value"`
            FormattedValue string  `json:"formatted_value"`
            Currency string  `json:"currency"`
         
    }
    
    // CreditSummary ...
    type CreditSummary struct {

        
            Balance BalanceDetails  `json:"balance"`
            StatusMessage string  `json:"status_message"`
            MerchantCustomerRefID string  `json:"merchant_customer_ref_id"`
            Status string  `json:"status"`
         
    }
    
    // CustomerCreditSummaryResponse ...
    type CustomerCreditSummaryResponse struct {

        
            Success bool  `json:"success"`
            Data CreditSummary  `json:"data"`
         
    }
    
    // RedirectURL ...
    type RedirectURL struct {

        
            SignupURL string  `json:"signup_url"`
            Status bool  `json:"status"`
         
    }
    
    // RedirectToAggregatorResponse ...
    type RedirectToAggregatorResponse struct {

        
            Success bool  `json:"success"`
            Data RedirectURL  `json:"data"`
         
    }
    
    // CreditDetail ...
    type CreditDetail struct {

        
            SignupURL string  `json:"signup_url"`
            IsRegistered bool  `json:"is_registered"`
            Status bool  `json:"status"`
         
    }
    
    // CheckCreditResponse ...
    type CheckCreditResponse struct {

        
            Success bool  `json:"success"`
            Data CreditDetail  `json:"data"`
         
    }
    
    // KYCAddress ...
    type KYCAddress struct {

        
            State string  `json:"state"`
            OwnershipType string  `json:"ownership_type"`
            Pincode string  `json:"pincode"`
            City string  `json:"city"`
            LandMark string  `json:"land_mark"`
            Addressline2 string  `json:"addressline2"`
            Addressline1 string  `json:"addressline1"`
         
    }
    
    // UserPersonalInfoInDetails ...
    type UserPersonalInfoInDetails struct {

        
            FirstName string  `json:"first_name"`
            MothersName string  `json:"mothers_name"`
            VoterID string  `json:"voter_id"`
            AddressAsPerID KYCAddress  `json:"address_as_per_id"`
            FathersName string  `json:"fathers_name"`
            MobileVerified bool  `json:"mobile_verified"`
            Dob string  `json:"dob"`
            Email string  `json:"email"`
            DrivingLicense string  `json:"driving_license"`
            Passport string  `json:"passport"`
            Gender string  `json:"gender"`
            LastName string  `json:"last_name"`
            Pan string  `json:"pan"`
            EmailVerified bool  `json:"email_verified"`
            Phone string  `json:"phone"`
            MiddleName string  `json:"middle_name"`
         
    }
    
    // MarketplaceInfo ...
    type MarketplaceInfo struct {

        
            MembershipID string  `json:"membership_id"`
            DateOfJoining string  `json:"date_of_joining"`
            Name string  `json:"name"`
         
    }
    
    // BusinessDetails ...
    type BusinessDetails struct {

        
            Name string  `json:"name"`
            Gstin string  `json:"gstin"`
            ShopAndEstablishment map[string]interface{}  `json:"shop_and_establishment"`
            Vintage string  `json:"vintage"`
            EntityType string  `json:"entity_type"`
            BusinessOwnershipType string  `json:"business_ownership_type"`
            Address KYCAddress  `json:"address"`
            Pan string  `json:"pan"`
            BusinessType string  `json:"business_type"`
            Fda string  `json:"fda"`
            Fssai string  `json:"fssai"`
         
    }
    
    // DeviceDetails ...
    type DeviceDetails struct {

        
            DeviceType string  `json:"device_type"`
            DeviceMake string  `json:"device_make"`
            OsVersion string  `json:"os_version"`
            IdentifierType string  `json:"identifier_type"`
            DeviceModel string  `json:"device_model"`
            Os string  `json:"os"`
            IdentificationNumber string  `json:"identification_number"`
         
    }
    
    // CustomerOnboardingRequest ...
    type CustomerOnboardingRequest struct {

        
            PersonalInfo UserPersonalInfoInDetails  `json:"personal_info"`
            Aggregator string  `json:"aggregator"`
            MarketplaceInfo MarketplaceInfo  `json:"marketplace_info"`
            BusinessInfo BusinessDetails  `json:"business_info"`
            Device DeviceDetails  `json:"device"`
            Source string  `json:"source"`
            Mcc string  `json:"mcc"`
         
    }
    
    // OnboardSummary ...
    type OnboardSummary struct {

        
            RedirectURL string  `json:"redirect_url"`
            Status bool  `json:"status"`
            Session map[string]interface{}  `json:"session"`
         
    }
    
    // CustomerOnboardingResponse ...
    type CustomerOnboardingResponse struct {

        
            Success bool  `json:"success"`
            Data OnboardSummary  `json:"data"`
         
    }
    

    
    // OrderPage ...
    type OrderPage struct {

        
            ItemTotal float64  `json:"item_total"`
            Current float64  `json:"current"`
            Size float64  `json:"size"`
            HasNext bool  `json:"has_next"`
            Type string  `json:"type"`
         
    }
    
    // NestedTrackingDetails ...
    type NestedTrackingDetails struct {

        
            IsPassed bool  `json:"is_passed"`
            Status string  `json:"status"`
            Time string  `json:"time"`
            IsCurrent bool  `json:"is_current"`
         
    }
    
    // TrackingDetails ...
    type TrackingDetails struct {

        
            IsPassed bool  `json:"is_passed"`
            IsCurrent bool  `json:"is_current"`
            Status string  `json:"status"`
            TrackingDetails []NestedTrackingDetails  `json:"tracking_details"`
            Time string  `json:"time"`
         
    }
    
    // ShipmentStatus ...
    type ShipmentStatus struct {

        
            HexCode string  `json:"hex_code"`
            Title string  `json:"title"`
         
    }
    
    // ShipmentUserInfo ...
    type ShipmentUserInfo struct {

        
            LastName string  `json:"last_name"`
            Mobile string  `json:"mobile"`
            FirstName string  `json:"first_name"`
            Gender string  `json:"gender"`
         
    }
    
    // ShipmentPayment ...
    type ShipmentPayment struct {

        
            DisplayName string  `json:"display_name"`
            Mode string  `json:"mode"`
            PaymentMode string  `json:"payment_mode"`
            Logo string  `json:"logo"`
            Status string  `json:"status"`
            Mop string  `json:"mop"`
         
    }
    
    // Invoice ...
    type Invoice struct {

        
            UpdatedDate string  `json:"updated_date"`
            InvoiceURL string  `json:"invoice_url"`
            LabelURL string  `json:"label_url"`
         
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
    
    // DeliveryAddress ...
    type DeliveryAddress struct {

        
            City string  `json:"city"`
            Email string  `json:"email"`
            Version string  `json:"version"`
            Address string  `json:"address"`
            CreatedAt string  `json:"created_at"`
            Phone string  `json:"phone"`
            State string  `json:"state"`
            Area string  `json:"area"`
            CountryPhoneCode string  `json:"country_phone_code"`
            CountryIsoCode string  `json:"country_iso_code"`
            Name string  `json:"name"`
            Address2 string  `json:"address2"`
            ContactPerson string  `json:"contact_person"`
            Address1 string  `json:"address1"`
            UpdatedAt string  `json:"updated_at"`
            AddressCategory string  `json:"address_category"`
            Latitude float64  `json:"latitude"`
            Country string  `json:"country"`
            Longitude float64  `json:"longitude"`
            AddressType string  `json:"address_type"`
            Pincode string  `json:"pincode"`
            Landmark string  `json:"landmark"`
         
    }
    
    // BreakupValues ...
    type BreakupValues struct {

        
            Display string  `json:"display"`
            CurrencySymbol string  `json:"currency_symbol"`
            Name string  `json:"name"`
            CurrencyCode string  `json:"currency_code"`
            Value float64  `json:"value"`
         
    }
    
    // ShipmentTotalDetails ...
    type ShipmentTotalDetails struct {

        
            Sizes float64  `json:"sizes"`
            Pieces float64  `json:"pieces"`
            TotalPrice float64  `json:"total_price"`
         
    }
    
    // FulfillingStore ...
    type FulfillingStore struct {

        
            ID float64  `json:"id"`
            Code string  `json:"code"`
            Name string  `json:"name"`
            CompanyName string  `json:"company_name"`
            CompanyID float64  `json:"company_id"`
         
    }
    
    // AppliedFreeArticles ...
    type AppliedFreeArticles struct {

        
            Quantity float64  `json:"quantity"`
            ArticleID string  `json:"article_id"`
            FreeGiftItemDetails map[string]interface{}  `json:"free_gift_item_details"`
            ParentItemIdentifier string  `json:"parent_item_identifier"`
         
    }
    
    // AppliedPromos ...
    type AppliedPromos struct {

        
            Amount float64  `json:"amount"`
            AppliedFreeArticles []AppliedFreeArticles  `json:"applied_free_articles"`
            ArticleQuantity float64  `json:"article_quantity"`
            PromotionType string  `json:"promotion_type"`
            PromoID string  `json:"promo_id"`
            PromotionName string  `json:"promotion_name"`
            MrpPromotion bool  `json:"mrp_promotion"`
         
    }
    
    // CurrentStatus ...
    type CurrentStatus struct {

        
            UpdatedAt string  `json:"updated_at"`
            Status string  `json:"status"`
            Name string  `json:"name"`
            JourneyType string  `json:"journey_type"`
            CreatedAt string  `json:"created_at"`
         
    }
    
    // Identifiers ...
    type Identifiers struct {

        
            SkuCode string  `json:"sku_code"`
            Ean string  `json:"ean"`
         
    }
    
    // FinancialBreakup ...
    type FinancialBreakup struct {

        
            ValueOfGood float64  `json:"value_of_good"`
            Discount float64  `json:"discount"`
            GstTag string  `json:"gst_tag"`
            ItemName string  `json:"item_name"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            Size string  `json:"size"`
            CashbackApplied float64  `json:"cashback_applied"`
            Identifiers Identifiers  `json:"identifiers"`
            PriceEffective float64  `json:"price_effective"`
            CodCharges float64  `json:"cod_charges"`
            Cashback float64  `json:"cashback"`
            GstFee float64  `json:"gst_fee"`
            FyndCredits float64  `json:"fynd_credits"`
            AddedToFyndCash bool  `json:"added_to_fynd_cash"`
            HsnCode string  `json:"hsn_code"`
            PromotionEffectiveDiscount float64  `json:"promotion_effective_discount"`
            CouponValue float64  `json:"coupon_value"`
            RefundCredit float64  `json:"refund_credit"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            AmountPaidRoundoff float64  `json:"amount_paid_roundoff"`
            TransferPrice float64  `json:"transfer_price"`
            DeliveryCharge float64  `json:"delivery_charge"`
            PriceMarked float64  `json:"price_marked"`
            RefundAmount float64  `json:"refund_amount"`
            TotalUnits float64  `json:"total_units"`
            AmountPaid float64  `json:"amount_paid"`
         
    }
    
    // Prices ...
    type Prices struct {

        
            ValueOfGood float64  `json:"value_of_good"`
            Discount float64  `json:"discount"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            CashbackApplied float64  `json:"cashback_applied"`
            PriceEffective float64  `json:"price_effective"`
            CodCharges float64  `json:"cod_charges"`
            Cashback float64  `json:"cashback"`
            FyndCredits float64  `json:"fynd_credits"`
            AddedToFyndCash bool  `json:"added_to_fynd_cash"`
            PromotionEffectiveDiscount float64  `json:"promotion_effective_discount"`
            CouponValue float64  `json:"coupon_value"`
            CurrencyCode string  `json:"currency_code"`
            RefundCredit float64  `json:"refund_credit"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            AmountPaidRoundoff float64  `json:"amount_paid_roundoff"`
            TransferPrice float64  `json:"transfer_price"`
            DeliveryCharge float64  `json:"delivery_charge"`
            CurrencySymbol string  `json:"currency_symbol"`
            PriceMarked float64  `json:"price_marked"`
            RefundAmount float64  `json:"refund_amount"`
            AmountPaid float64  `json:"amount_paid"`
         
    }
    
    // ItemBrand ...
    type ItemBrand struct {

        
            Logo string  `json:"logo"`
            Name string  `json:"name"`
         
    }
    
    // Item ...
    type Item struct {

        
            Image []string  `json:"image"`
            Brand ItemBrand  `json:"brand"`
            Code string  `json:"code"`
            ID float64  `json:"id"`
            Name string  `json:"name"`
            Size string  `json:"size"`
            SlugKey string  `json:"slug_key"`
            SellerIdentifier string  `json:"seller_identifier"`
         
    }
    
    // Bags ...
    type Bags struct {

        
            Quantity float64  `json:"quantity"`
            AppliedPromos []AppliedPromos  `json:"applied_promos"`
            Meta map[string]interface{}  `json:"meta"`
            CurrentStatus CurrentStatus  `json:"current_status"`
            ID float64  `json:"id"`
            DeliveryDate string  `json:"delivery_date"`
            ParentPromoBags map[string]interface{}  `json:"parent_promo_bags"`
            CanReturn bool  `json:"can_return"`
            CurrencySymbol string  `json:"currency_symbol"`
            ReturnableDate string  `json:"returnable_date"`
            CurrencyCode string  `json:"currency_code"`
            FinancialBreakup []FinancialBreakup  `json:"financial_breakup"`
            Prices Prices  `json:"prices"`
            LineNumber float64  `json:"line_number"`
            CanCancel bool  `json:"can_cancel"`
            SellerIdentifier string  `json:"seller_identifier"`
            Item Item  `json:"item"`
         
    }
    
    // FulfillingCompany ...
    type FulfillingCompany struct {

        
            ID float64  `json:"id"`
            Name string  `json:"name"`
         
    }
    
    // Shipments ...
    type Shipments struct {

        
            CanBreak map[string]interface{}  `json:"can_break"`
            TrakingNo string  `json:"traking_no"`
            TrackURL string  `json:"track_url"`
            TrackingDetails []TrackingDetails  `json:"tracking_details"`
            ShipmentStatus ShipmentStatus  `json:"shipment_status"`
            UserInfo ShipmentUserInfo  `json:"user_info"`
            CustomMeta []map[string]interface{}  `json:"custom_meta"`
            Payment ShipmentPayment  `json:"payment"`
            Comment string  `json:"comment"`
            BeneficiaryDetails bool  `json:"beneficiary_details"`
            OrderType string  `json:"order_type"`
            OrderID string  `json:"order_id"`
            Invoice Invoice  `json:"invoice"`
            ReturnMeta map[string]interface{}  `json:"return_meta"`
            ReturnableDate string  `json:"returnable_date"`
            ShipmentCreatedAt string  `json:"shipment_created_at"`
            Promise Promise  `json:"promise"`
            SizeInfo map[string]interface{}  `json:"size_info"`
            NeedHelpURL string  `json:"need_help_url"`
            ShipmentID string  `json:"shipment_id"`
            DeliveryAddress DeliveryAddress  `json:"delivery_address"`
            BreakupValues []BreakupValues  `json:"breakup_values"`
            TotalBags float64  `json:"total_bags"`
            ShowDownloadInvoice bool  `json:"show_download_invoice"`
            DpName string  `json:"dp_name"`
            CanCancel bool  `json:"can_cancel"`
            TotalDetails ShipmentTotalDetails  `json:"total_details"`
            FulfillingStore FulfillingStore  `json:"fulfilling_store"`
            DeliveryDate string  `json:"delivery_date"`
            CanReturn bool  `json:"can_return"`
            ShowTrackLink bool  `json:"show_track_link"`
            Bags []Bags  `json:"bags"`
            RefundDetails map[string]interface{}  `json:"refund_details"`
            Prices Prices  `json:"prices"`
            FulfillingCompany FulfillingCompany  `json:"fulfilling_company"`
            AwbNo string  `json:"awb_no"`
         
    }
    
    // UserInfo ...
    type UserInfo struct {

        
            Name string  `json:"name"`
            Email string  `json:"email"`
            Mobile string  `json:"mobile"`
            Gender string  `json:"gender"`
         
    }
    
    // BagsForReorderArticleAssignment ...
    type BagsForReorderArticleAssignment struct {

        
            Level string  `json:"level"`
            Strategy string  `json:"strategy"`
         
    }
    
    // BagsForReorder ...
    type BagsForReorder struct {

        
            Quantity float64  `json:"quantity"`
            ItemSize string  `json:"item_size"`
            StoreID float64  `json:"store_id"`
            SellerID float64  `json:"seller_id"`
            ArticleAssignment BagsForReorderArticleAssignment  `json:"article_assignment"`
            ItemID float64  `json:"item_id"`
         
    }
    
    // OrderSchema ...
    type OrderSchema struct {

        
            OrderCreatedTime string  `json:"order_created_time"`
            Shipments []Shipments  `json:"shipments"`
            BreakupValues []BreakupValues  `json:"breakup_values"`
            OrderID string  `json:"order_id"`
            TotalShipmentsInOrder float64  `json:"total_shipments_in_order"`
            UserInfo UserInfo  `json:"user_info"`
            BagsForReorder []BagsForReorder  `json:"bags_for_reorder"`
         
    }
    
    // OrderStatuses ...
    type OrderStatuses struct {

        
            Display string  `json:"display"`
            IsSelected bool  `json:"is_selected"`
            Value float64  `json:"value"`
         
    }
    
    // OrderFilters ...
    type OrderFilters struct {

        
            Statuses []OrderStatuses  `json:"statuses"`
         
    }
    
    // OrderList ...
    type OrderList struct {

        
            Page OrderPage  `json:"page"`
            Items []OrderSchema  `json:"items"`
            Filters OrderFilters  `json:"filters"`
         
    }
    
    // ApefaceApiError ...
    type ApefaceApiError struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
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

        
            Success bool  `json:"success"`
            PresignedType string  `json:"presigned_type"`
            ShipmentID string  `json:"shipment_id"`
            PresignedURL string  `json:"presigned_url"`
         
    }
    
    // Track ...
    type Track struct {

        
            UpdatedTime string  `json:"updated_time"`
            Awb string  `json:"awb"`
            AccountName string  `json:"account_name"`
            Reason string  `json:"reason"`
            UpdatedAt string  `json:"updated_at"`
            Status string  `json:"status"`
            LastLocationRecievedAt string  `json:"last_location_recieved_at"`
            ShipmentType string  `json:"shipment_type"`
         
    }
    
    // ShipmentTrack ...
    type ShipmentTrack struct {

        
            Results []Track  `json:"results"`
         
    }
    
    // CustomerDetailsResponse ...
    type CustomerDetailsResponse struct {

        
            OrderID string  `json:"order_id"`
            Country string  `json:"country"`
            Name string  `json:"name"`
            Phone string  `json:"phone"`
            ShipmentID string  `json:"shipment_id"`
         
    }
    
    // SendOtpToCustomerResponse ...
    type SendOtpToCustomerResponse struct {

        
            Success bool  `json:"success"`
            ResendTimer float64  `json:"resend_timer"`
            RequestID string  `json:"request_id"`
            Message string  `json:"message"`
         
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

        
            DisplayName string  `json:"display_name"`
            Meta BagReasonMeta  `json:"meta"`
            QcType []string  `json:"qc_type"`
            ID float64  `json:"id"`
            Reasons []BagReasons  `json:"reasons"`
            QuestionSet []QuestionSet  `json:"question_set"`
         
    }
    
    // ShipmentBagReasons ...
    type ShipmentBagReasons struct {

        
            Success bool  `json:"success"`
            Reasons []BagReasons  `json:"reasons"`
         
    }
    
    // ShipmentReason ...
    type ShipmentReason struct {

        
            Priority float64  `json:"priority"`
            Flow string  `json:"flow"`
            FeedbackType string  `json:"feedback_type"`
            ReasonID float64  `json:"reason_id"`
            ShowTextArea bool  `json:"show_text_area"`
            ReasonText string  `json:"reason_text"`
         
    }
    
    // ShipmentReasons ...
    type ShipmentReasons struct {

        
            Reasons []ShipmentReason  `json:"reasons"`
         
    }
    
    // ProductsDataUpdatesFilters ...
    type ProductsDataUpdatesFilters struct {

        
            Identifier string  `json:"identifier"`
            LineNumber float64  `json:"line_number"`
         
    }
    
    // ProductsDataUpdates ...
    type ProductsDataUpdates struct {

        
            Data map[string]interface{}  `json:"data"`
            Filters []ProductsDataUpdatesFilters  `json:"filters"`
         
    }
    
    // EntitiesDataUpdates ...
    type EntitiesDataUpdates struct {

        
            Data map[string]interface{}  `json:"data"`
            Filters []map[string]interface{}  `json:"filters"`
         
    }
    
    // DataUpdates ...
    type DataUpdates struct {

        
            Products []ProductsDataUpdates  `json:"products"`
            Entities []EntitiesDataUpdates  `json:"entities"`
         
    }
    
    // Products ...
    type Products struct {

        
            Identifier string  `json:"identifier"`
            LineNumber float64  `json:"line_number"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // ProductsReasonsData ...
    type ProductsReasonsData struct {

        
            ReasonID float64  `json:"reason_id"`
            ReasonText string  `json:"reason_text"`
         
    }
    
    // ProductsReasonsFilters ...
    type ProductsReasonsFilters struct {

        
            Identifier string  `json:"identifier"`
            LineNumber float64  `json:"line_number"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // ProductsReasons ...
    type ProductsReasons struct {

        
            Data ProductsReasonsData  `json:"data"`
            Filters []ProductsReasonsFilters  `json:"filters"`
         
    }
    
    // EntityReasonData ...
    type EntityReasonData struct {

        
            ReasonID float64  `json:"reason_id"`
            ReasonText string  `json:"reason_text"`
         
    }
    
    // EntitiesReasons ...
    type EntitiesReasons struct {

        
            Data EntityReasonData  `json:"data"`
            Filters []map[string]interface{}  `json:"filters"`
         
    }
    
    // ReasonsData ...
    type ReasonsData struct {

        
            Products []ProductsReasons  `json:"products"`
            Entities []EntitiesReasons  `json:"entities"`
         
    }
    
    // ShipmentsRequest ...
    type ShipmentsRequest struct {

        
            Identifier string  `json:"identifier"`
            DataUpdates DataUpdates  `json:"data_updates"`
            Products []Products  `json:"products"`
            Reasons ReasonsData  `json:"reasons"`
         
    }
    
    // StatuesRequest ...
    type StatuesRequest struct {

        
            Shipments []ShipmentsRequest  `json:"shipments"`
            Status string  `json:"status"`
            ExcludeBagsNextState string  `json:"exclude_bags_next_state"`
         
    }
    
    // UpdateShipmentStatusRequest ...
    type UpdateShipmentStatusRequest struct {

        
            UnlockBeforeTransition bool  `json:"unlock_before_transition"`
            LockAfterTransition bool  `json:"lock_after_transition"`
            ForceTransition bool  `json:"force_transition"`
            Task bool  `json:"task"`
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
            Code string  `json:"code"`
            StackTrace string  `json:"stack_trace"`
            Status float64  `json:"status"`
            Message string  `json:"message"`
         
    }
    
    // ProductBrand ...
    type ProductBrand struct {

        
            ID float64  `json:"id"`
            BrandName string  `json:"brand_name"`
         
    }
    
    // Coupon ...
    type Coupon struct {

        
            PayableCategory string  `json:"payable_category"`
            ID float64  `json:"id"`
            Code string  `json:"code"`
            CouponType string  `json:"coupon_type"`
            Value float64  `json:"value"`
         
    }
    
    // ProductStatus ...
    type ProductStatus struct {

        
            HexCode string  `json:"hex_code"`
            Title string  `json:"title"`
            CreatedAt string  `json:"created_at"`
            Value string  `json:"value"`
         
    }
    
    // Product ...
    type Product struct {

        
            AppliedPromos []AppliedPromos  `json:"applied_promos"`
            Quantity float64  `json:"quantity"`
            OrderID string  `json:"order_id"`
            ParentPromoBags map[string]interface{}  `json:"parent_promo_bags"`
            Brand ProductBrand  `json:"brand"`
            DeliveryDate string  `json:"delivery_date"`
            CanReturn bool  `json:"can_return"`
            Coupon Coupon  `json:"coupon"`
            BagStatus ProductStatus  `json:"bag_status"`
            ReturnableDate string  `json:"returnable_date"`
            LineNumber float64  `json:"line_number"`
            DocketNumber string  `json:"docket_number"`
            Payment ShipmentPayment  `json:"payment"`
            CanCancel bool  `json:"can_cancel"`
            SellerIdentifier string  `json:"seller_identifier"`
            Item Item  `json:"item"`
         
    }
    
    // ProductListResponse ...
    type ProductListResponse struct {

        
            Filters OrderFilters  `json:"filters"`
            Success bool  `json:"success"`
            Items []Product  `json:"items"`
            Message string  `json:"message"`
            Page OrderPage  `json:"page"`
         
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
    
    // DiscountRulesAppSchema ...
    type DiscountRulesAppSchema struct {

        
            Offer map[string]interface{}  `json:"offer"`
            RawOffer map[string]interface{}  `json:"raw_offer"`
            ItemCriteria map[string]interface{}  `json:"item_criteria"`
            MatchedBuyRules []string  `json:"matched_buy_rules"`
         
    }
    
    // BuyRulesSchema ...
    type BuyRulesSchema struct {

        
            ItemCriteria map[string]interface{}  `json:"item_criteria"`
            CartConditions map[string]interface{}  `json:"cart_conditions"`
         
    }
    
    // Ownership ...
    type Ownership struct {

        
            PayableBy string  `json:"payable_by"`
            PayableCategory string  `json:"payable_category"`
         
    }
    
    // FreeGiftItemSchema ...
    type FreeGiftItemSchema struct {

        
            ItemSlug string  `json:"item_slug"`
            ItemBrandName string  `json:"item_brand_name"`
            ItemName string  `json:"item_name"`
            ItemID float64  `json:"item_id"`
            ItemPriceDetails map[string]interface{}  `json:"item_price_details"`
            ItemImagesURL []string  `json:"item_images_url"`
         
    }
    
    // AppliedFreeArticlesSchema ...
    type AppliedFreeArticlesSchema struct {

        
            Quantity float64  `json:"quantity"`
            ParentItemIdentifier string  `json:"parent_item_identifier"`
            FreeGiftItemDetails FreeGiftItemSchema  `json:"free_gift_item_details"`
            ArticleID string  `json:"article_id"`
         
    }
    
    // AppliedPromotion ...
    type AppliedPromotion struct {

        
            OfferText string  `json:"offer_text"`
            DiscountRules []DiscountRulesAppSchema  `json:"discount_rules"`
            PromotionName string  `json:"promotion_name"`
            PromotionType string  `json:"promotion_type"`
            PromoID string  `json:"promo_id"`
            PromotionGroup string  `json:"promotion_group"`
            BuyRules []BuyRulesSchema  `json:"buy_rules"`
            ArticleQuantity float64  `json:"article_quantity"`
            Ownership Ownership  `json:"ownership"`
            Amount float64  `json:"amount"`
            AppliedFreeArticles []AppliedFreeArticlesSchema  `json:"applied_free_articles"`
            MrpPromotion bool  `json:"mrp_promotion"`
         
    }
    
    // ProductPrice ...
    type ProductPrice struct {

        
            Effective float64  `json:"effective"`
            Marked float64  `json:"marked"`
            CurrencySymbol string  `json:"currency_symbol"`
            AddOn float64  `json:"add_on"`
            Selling float64  `json:"selling"`
            CurrencyCode string  `json:"currency_code"`
         
    }
    
    // ProductPriceInfo ...
    type ProductPriceInfo struct {

        
            Converted ProductPrice  `json:"converted"`
            Base ProductPrice  `json:"base"`
         
    }
    
    // CouponDetails ...
    type CouponDetails struct {

        
            DiscountTotalQuantity float64  `json:"discount_total_quantity"`
            DiscountSingleQuantity float64  `json:"discount_single_quantity"`
            Code string  `json:"code"`
         
    }
    
    // PromoMeta ...
    type PromoMeta struct {

        
            Message string  `json:"message"`
         
    }
    
    // ProductAvailabilitySize ...
    type ProductAvailabilitySize struct {

        
            Value string  `json:"value"`
            IsAvailable bool  `json:"is_available"`
            Display string  `json:"display"`
         
    }
    
    // ProductAvailability ...
    type ProductAvailability struct {

        
            OtherStoreQuantity float64  `json:"other_store_quantity"`
            OutOfStock bool  `json:"out_of_stock"`
            IsValid bool  `json:"is_valid"`
            AvailableSizes []ProductAvailabilitySize  `json:"available_sizes"`
            Sizes []string  `json:"sizes"`
            Deliverable bool  `json:"deliverable"`
         
    }
    
    // BaseInfo ...
    type BaseInfo struct {

        
            Name string  `json:"name"`
            UID float64  `json:"uid"`
         
    }
    
    // BasePriceSchema ...
    type BasePriceSchema struct {

        
            Marked float64  `json:"marked"`
            CurrencySymbol string  `json:"currency_symbol"`
            Effective float64  `json:"effective"`
            CurrencyCode string  `json:"currency_code"`
         
    }
    
    // ArticlePriceInfo ...
    type ArticlePriceInfo struct {

        
            Converted BasePriceSchema  `json:"converted"`
            Base BasePriceSchema  `json:"base"`
         
    }
    
    // ProductArticle ...
    type ProductArticle struct {

        
            Store BaseInfo  `json:"store"`
            ProductGroupTags []string  `json:"product_group_tags"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            IsGiftVisible bool  `json:"is_gift_visible"`
            Seller BaseInfo  `json:"seller"`
            Quantity float64  `json:"quantity"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            GiftCard map[string]interface{}  `json:"gift_card"`
            Identifier map[string]interface{}  `json:"identifier"`
            Size string  `json:"size"`
            Price ArticlePriceInfo  `json:"price"`
            CartItemMeta map[string]interface{}  `json:"cart_item_meta"`
            Type string  `json:"type"`
            SellerIdentifier string  `json:"seller_identifier"`
            UID string  `json:"uid"`
         
    }
    
    // Tags ...
    type Tags struct {

        
            Tags map[string]interface{}  `json:"tags"`
         
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

        
            URL string  `json:"url"`
            Type string  `json:"type"`
            Query ActionQuery  `json:"query"`
         
    }
    
    // NetQuantity ...
    type NetQuantity struct {

        
            Value string  `json:"value"`
            Unit string  `json:"unit"`
         
    }
    
    // CategoryInfo ...
    type CategoryInfo struct {

        
            Name string  `json:"name"`
            UID float64  `json:"uid"`
         
    }
    
    // CartProduct ...
    type CartProduct struct {

        
            TeaserTag Tags  `json:"teaser_tag"`
            Name string  `json:"name"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Images []ProductImage  `json:"images"`
            Action ProductAction  `json:"action"`
            NetQuantity NetQuantity  `json:"net_quantity"`
            Brand BaseInfo  `json:"brand"`
            Slug string  `json:"slug"`
            Categories []CategoryInfo  `json:"categories"`
            Type string  `json:"type"`
            UID float64  `json:"uid"`
            ItemCode string  `json:"item_code"`
            Tags []string  `json:"tags"`
         
    }
    
    // CartProductIdentifer ...
    type CartProductIdentifer struct {

        
            Identifier string  `json:"identifier"`
         
    }
    
    // CartProductInfo ...
    type CartProductInfo struct {

        
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            BulkOffer map[string]interface{}  `json:"bulk_offer"`
            PromotionsApplied []AppliedPromotion  `json:"promotions_applied"`
            PricePerUnit ProductPriceInfo  `json:"price_per_unit"`
            Price ProductPriceInfo  `json:"price"`
            Key string  `json:"key"`
            Coupon CouponDetails  `json:"coupon"`
            Discount string  `json:"discount"`
            Quantity float64  `json:"quantity"`
            PromoMeta PromoMeta  `json:"promo_meta"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            Availability ProductAvailability  `json:"availability"`
            Article ProductArticle  `json:"article"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            Moq map[string]interface{}  `json:"moq"`
            Product CartProduct  `json:"product"`
            Message string  `json:"message"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
            CouponMessage string  `json:"coupon_message"`
            IsSet bool  `json:"is_set"`
         
    }
    
    // CartCurrencySchema ...
    type CartCurrencySchema struct {

        
            Symbol string  `json:"symbol"`
            Code string  `json:"code"`
         
    }
    
    // PaymentSelectionLock ...
    type PaymentSelectionLock struct {

        
            PaymentIdentifier string  `json:"payment_identifier"`
            DefaultOptions string  `json:"default_options"`
            Enabled bool  `json:"enabled"`
         
    }
    
    // RawBreakupSchema ...
    type RawBreakupSchema struct {

        
            Coupon float64  `json:"coupon"`
            Vog float64  `json:"vog"`
            Discount float64  `json:"discount"`
            FyndCash float64  `json:"fynd_cash"`
            GstCharges float64  `json:"gst_charges"`
            ConvenienceFee float64  `json:"convenience_fee"`
            YouSaved float64  `json:"you_saved"`
            MrpTotal float64  `json:"mrp_total"`
            GiftCard float64  `json:"gift_card"`
            Total float64  `json:"total"`
            CodCharge float64  `json:"cod_charge"`
            Subtotal float64  `json:"subtotal"`
            DeliveryCharge float64  `json:"delivery_charge"`
         
    }
    
    // CouponBreakupSchema ...
    type CouponBreakupSchema struct {

        
            Message string  `json:"message"`
            SubTitle string  `json:"sub_title"`
            Code string  `json:"code"`
            MaxDiscountValue float64  `json:"max_discount_value"`
            Title string  `json:"title"`
            Description string  `json:"description"`
            MinimumCartValue float64  `json:"minimum_cart_value"`
            IsApplied bool  `json:"is_applied"`
            CouponValue float64  `json:"coupon_value"`
            Value float64  `json:"value"`
            Type string  `json:"type"`
            UID string  `json:"uid"`
            CouponType string  `json:"coupon_type"`
         
    }
    
    // LoyaltyPoints ...
    type LoyaltyPoints struct {

        
            Description string  `json:"description"`
            Applicable float64  `json:"applicable"`
            IsApplied bool  `json:"is_applied"`
            Total float64  `json:"total"`
         
    }
    
    // DisplayBreakupSchema ...
    type DisplayBreakupSchema struct {

        
            Message []string  `json:"message"`
            CurrencySymbol string  `json:"currency_symbol"`
            Display string  `json:"display"`
            Key string  `json:"key"`
            Value float64  `json:"value"`
            CurrencyCode string  `json:"currency_code"`
         
    }
    
    // CartBreakupSchema ...
    type CartBreakupSchema struct {

        
            Raw RawBreakupSchema  `json:"raw"`
            Coupon CouponBreakupSchema  `json:"coupon"`
            LoyaltyPoints LoyaltyPoints  `json:"loyalty_points"`
            Display []DisplayBreakupSchema  `json:"display"`
         
    }
    
    // CartDetailResponse ...
    type CartDetailResponse struct {

        
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            PanConfig map[string]interface{}  `json:"pan_config"`
            Comment string  `json:"comment"`
            LastModified string  `json:"last_modified"`
            CouponText string  `json:"coupon_text"`
            Items []CartProductInfo  `json:"items"`
            IsValid bool  `json:"is_valid"`
            ID string  `json:"id"`
            Currency CartCurrencySchema  `json:"currency"`
            BuyNow bool  `json:"buy_now"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            CheckoutMode string  `json:"checkout_mode"`
            BreakupValues CartBreakupSchema  `json:"breakup_values"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            Message string  `json:"message"`
            AppliedPromoDetails []AppliedPromotion  `json:"applied_promo_details"`
            PanNo string  `json:"pan_no"`
            Gstin string  `json:"gstin"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
         
    }
    
    // AddProductCart ...
    type AddProductCart struct {

        
            Pos bool  `json:"pos"`
            ProductGroupTags []string  `json:"product_group_tags"`
            ItemSize string  `json:"item_size"`
            ArticleID string  `json:"article_id"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Quantity float64  `json:"quantity"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            ParentItemIdentifiers []map[string]string  `json:"parent_item_identifiers"`
            ItemID float64  `json:"item_id"`
            StoreID float64  `json:"store_id"`
            ArticleAssignment map[string]interface{}  `json:"article_assignment"`
            Display string  `json:"display"`
            SellerID float64  `json:"seller_id"`
         
    }
    
    // AddCartRequest ...
    type AddCartRequest struct {

        
            Items []AddProductCart  `json:"items"`
            NewCart bool  `json:"new_cart"`
         
    }
    
    // AddCartDetailResponse ...
    type AddCartDetailResponse struct {

        
            Cart CartDetailResponse  `json:"cart"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
            Partial bool  `json:"partial"`
         
    }
    
    // UpdateProductCart ...
    type UpdateProductCart struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ItemSize string  `json:"item_size"`
            ArticleID string  `json:"article_id"`
            Quantity float64  `json:"quantity"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            ItemID float64  `json:"item_id"`
            ItemIndex float64  `json:"item_index"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
         
    }
    
    // UpdateCartRequest ...
    type UpdateCartRequest struct {

        
            Operation string  `json:"operation"`
            Items []UpdateProductCart  `json:"items"`
         
    }
    
    // UpdateCartDetailResponse ...
    type UpdateCartDetailResponse struct {

        
            Cart CartDetailResponse  `json:"cart"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // CartItemCountResponse ...
    type CartItemCountResponse struct {

        
            UserCartItemsCount float64  `json:"user_cart_items_count"`
         
    }
    
    // Coupon ...
    type Coupon struct {

        
            Message string  `json:"message"`
            SubTitle string  `json:"sub_title"`
            IsApplicable bool  `json:"is_applicable"`
            MaxDiscountValue float64  `json:"max_discount_value"`
            CouponCode string  `json:"coupon_code"`
            ExpiresOn string  `json:"expires_on"`
            Description string  `json:"description"`
            MinimumCartValue float64  `json:"minimum_cart_value"`
            IsApplied bool  `json:"is_applied"`
            CouponValue float64  `json:"coupon_value"`
            Title string  `json:"title"`
            CouponType string  `json:"coupon_type"`
         
    }
    
    // PageCouponSchema ...
    type PageCouponSchema struct {

        
            Current float64  `json:"current"`
            TotalItemCount float64  `json:"total_item_count"`
            HasPrevious bool  `json:"has_previous"`
            HasNext bool  `json:"has_next"`
            Total float64  `json:"total"`
         
    }
    
    // GetCouponResponse ...
    type GetCouponResponse struct {

        
            AvailableCouponList []Coupon  `json:"available_coupon_list"`
            Page PageCouponSchema  `json:"page"`
         
    }
    
    // ApplyCouponRequest ...
    type ApplyCouponRequest struct {

        
            CouponCode string  `json:"coupon_code"`
         
    }
    
    // OfferSellerSchema ...
    type OfferSellerSchema struct {

        
            Name string  `json:"name"`
            UID float64  `json:"uid"`
         
    }
    
    // OfferPrice ...
    type OfferPrice struct {

        
            Effective float64  `json:"effective"`
            Marked float64  `json:"marked"`
            CurrencySymbol string  `json:"currency_symbol"`
            BulkEffective float64  `json:"bulk_effective"`
            CurrencyCode string  `json:"currency_code"`
         
    }
    
    // OfferItemSchema ...
    type OfferItemSchema struct {

        
            AutoApplied bool  `json:"auto_applied"`
            Quantity float64  `json:"quantity"`
            Margin float64  `json:"margin"`
            Total float64  `json:"total"`
            Best bool  `json:"best"`
            Price OfferPrice  `json:"price"`
            Type string  `json:"type"`
         
    }
    
    // BulkPriceOfferSchema ...
    type BulkPriceOfferSchema struct {

        
            Seller OfferSellerSchema  `json:"seller"`
            Offers []OfferItemSchema  `json:"offers"`
         
    }
    
    // BulkPriceResponse ...
    type BulkPriceResponse struct {

        
            Data []BulkPriceOfferSchema  `json:"data"`
         
    }
    
    // RewardPointRequestSchema ...
    type RewardPointRequestSchema struct {

        
            Points bool  `json:"points"`
         
    }
    
    // GeoLocation ...
    type GeoLocation struct {

        
            Longitude float64  `json:"longitude"`
            Latitude float64  `json:"latitude"`
         
    }
    
    // Address ...
    type Address struct {

        
            CountryIsoCode string  `json:"country_iso_code"`
            Name string  `json:"name"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            State string  `json:"state"`
            AreaCodeSlug string  `json:"area_code_slug"`
            CreatedByUserID string  `json:"created_by_user_id"`
            IsDefaultAddress bool  `json:"is_default_address"`
            UserID string  `json:"user_id"`
            Country string  `json:"country"`
            GoogleMapPoint map[string]interface{}  `json:"google_map_point"`
            Phone string  `json:"phone"`
            ID string  `json:"id"`
            AddressType string  `json:"address_type"`
            City string  `json:"city"`
            Area string  `json:"area"`
            CountryCode string  `json:"country_code"`
            CountryPhoneCode string  `json:"country_phone_code"`
            CheckoutMode string  `json:"checkout_mode"`
            Meta map[string]interface{}  `json:"meta"`
            Tags []string  `json:"tags"`
            GeoLocation GeoLocation  `json:"geo_location"`
            IsActive bool  `json:"is_active"`
            Address string  `json:"address"`
            Email string  `json:"email"`
            AreaCode string  `json:"area_code"`
            Landmark string  `json:"landmark"`
         
    }
    
    // GetAddressesResponse ...
    type GetAddressesResponse struct {

        
            PiiMasking bool  `json:"pii_masking"`
            Address []Address  `json:"address"`
         
    }
    
    // SaveAddressResponse ...
    type SaveAddressResponse struct {

        
            Success bool  `json:"success"`
            IsDefaultAddress bool  `json:"is_default_address"`
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

        
            BillingAddressID string  `json:"billing_address_id"`
            CartID string  `json:"cart_id"`
            ID string  `json:"id"`
         
    }
    
    // UpdateCartPaymentRequest ...
    type UpdateCartPaymentRequest struct {

        
            AddressID string  `json:"address_id"`
            PaymentMode string  `json:"payment_mode"`
            AggregatorName string  `json:"aggregator_name"`
            ID string  `json:"id"`
            MerchantCode string  `json:"merchant_code"`
            PaymentIdentifier string  `json:"payment_identifier"`
         
    }
    
    // CouponValidity ...
    type CouponValidity struct {

        
            Discount float64  `json:"discount"`
            Code string  `json:"code"`
            DisplayMessageEn string  `json:"display_message_en"`
            Valid bool  `json:"valid"`
            Title string  `json:"title"`
         
    }
    
    // PaymentCouponValidateSchema ...
    type PaymentCouponValidateSchema struct {

        
            Success bool  `json:"success"`
            CouponValidity CouponValidity  `json:"coupon_validity"`
            Message string  `json:"message"`
         
    }
    
    // ShipmentResponse ...
    type ShipmentResponse struct {

        
            ShipmentType string  `json:"shipment_type"`
            OrderType string  `json:"order_type"`
            BoxType string  `json:"box_type"`
            DpID string  `json:"dp_id"`
            DpOptions map[string]interface{}  `json:"dp_options"`
            FulfillmentType string  `json:"fulfillment_type"`
            Items []CartProductInfo  `json:"items"`
            FulfillmentID float64  `json:"fulfillment_id"`
            Shipments float64  `json:"shipments"`
            Promise ShipmentPromise  `json:"promise"`
         
    }
    
    // CartShipmentsResponse ...
    type CartShipmentsResponse struct {

        
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            Currency CartCurrencySchema  `json:"currency"`
            BuyNow bool  `json:"buy_now"`
            Shipments []ShipmentResponse  `json:"shipments"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            Comment string  `json:"comment"`
            CartID float64  `json:"cart_id"`
            CheckoutMode string  `json:"checkout_mode"`
            LastModified string  `json:"last_modified"`
            UID string  `json:"uid"`
            BreakupValues CartBreakupSchema  `json:"breakup_values"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            Message string  `json:"message"`
            Error bool  `json:"error"`
            CouponText string  `json:"coupon_text"`
            Gstin string  `json:"gstin"`
            IsValid bool  `json:"is_valid"`
            ID string  `json:"id"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
         
    }
    
    // UpdateCartShipmentItem ...
    type UpdateCartShipmentItem struct {

        
            ShipmentType string  `json:"shipment_type"`
            Quantity float64  `json:"quantity"`
            ArticleUID string  `json:"article_uid"`
         
    }
    
    // UpdateCartShipmentRequest ...
    type UpdateCartShipmentRequest struct {

        
            Shipments []UpdateCartShipmentItem  `json:"shipments"`
         
    }
    
    // StaffCheckoutSchema ...
    type StaffCheckoutSchema struct {

        
            EmployeeCode string  `json:"employee_code"`
            User string  `json:"user"`
            ID string  `json:"_id"`
            FirstName string  `json:"first_name"`
            LastName string  `json:"last_name"`
         
    }
    
    // FilesSchema ...
    type FilesSchema struct {

        
            Key string  `json:"key"`
            Values []string  `json:"values"`
         
    }
    
    // CartCheckoutCustomMetaSchema ...
    type CartCheckoutCustomMetaSchema struct {

        
            Key string  `json:"key"`
            Value string  `json:"value"`
         
    }
    
    // CartPosCheckoutDetailRequest ...
    type CartPosCheckoutDetailRequest struct {

        
            OrderType string  `json:"order_type"`
            Pos bool  `json:"pos"`
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
            PaymentAutoConfirm bool  `json:"payment_auto_confirm"`
            OrderingStore float64  `json:"ordering_store"`
            Staff StaffCheckoutSchema  `json:"staff"`
            CallbackURL string  `json:"callback_url"`
            Files []FilesSchema  `json:"files"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            AddressID string  `json:"address_id"`
            ID string  `json:"id"`
            BillingAddressID string  `json:"billing_address_id"`
            MerchantCode string  `json:"merchant_code"`
            PaymentMode string  `json:"payment_mode"`
            Aggregator string  `json:"aggregator"`
            Meta map[string]interface{}  `json:"meta"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
            PickAtStoreUID float64  `json:"pick_at_store_uid"`
            CustomMeta []CartCheckoutCustomMetaSchema  `json:"custom_meta"`
            PaymentParams map[string]interface{}  `json:"payment_params"`
            PaymentIdentifier string  `json:"payment_identifier"`
         
    }
    
    // CheckCartSchema ...
    type CheckCartSchema struct {

        
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            CodCharges float64  `json:"cod_charges"`
            OrderID string  `json:"order_id"`
            Comment string  `json:"comment"`
            CartID float64  `json:"cart_id"`
            ErrorMessage string  `json:"error_message"`
            LastModified string  `json:"last_modified"`
            UID string  `json:"uid"`
            CouponText string  `json:"coupon_text"`
            Items []CartProductInfo  `json:"items"`
            IsValid bool  `json:"is_valid"`
            ID string  `json:"id"`
            CodMessage string  `json:"cod_message"`
            UserType string  `json:"user_type"`
            Currency CartCurrencySchema  `json:"currency"`
            DeliveryCharges float64  `json:"delivery_charges"`
            BuyNow bool  `json:"buy_now"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            StoreEmps []map[string]interface{}  `json:"store_emps"`
            CheckoutMode string  `json:"checkout_mode"`
            Success bool  `json:"success"`
            DeliveryChargeOrderValue float64  `json:"delivery_charge_order_value"`
            BreakupValues CartBreakupSchema  `json:"breakup_values"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            Message string  `json:"message"`
            StoreCode string  `json:"store_code"`
            CodAvailable bool  `json:"cod_available"`
            Gstin string  `json:"gstin"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
         
    }
    
    // CartCheckoutResponseSchema ...
    type CartCheckoutResponseSchema struct {

        
            Cart CheckCartSchema  `json:"cart"`
            Message string  `json:"message"`
            PaymentConfirmURL string  `json:"payment_confirm_url"`
            AppInterceptURL string  `json:"app_intercept_url"`
            OrderID string  `json:"order_id"`
            Data map[string]interface{}  `json:"data"`
            Success bool  `json:"success"`
            CallbackURL string  `json:"callback_url"`
         
    }
    
    // GiftDetailSchema ...
    type GiftDetailSchema struct {

        
            IsGiftApplied bool  `json:"is_gift_applied"`
            GiftMessage string  `json:"gift_message"`
         
    }
    
    // ArticleGiftDetailSchema ...
    type ArticleGiftDetailSchema struct {

        
            ArticleID GiftDetailSchema  `json:"article_id"`
         
    }
    
    // CartMetaRequest ...
    type CartMetaRequest struct {

        
            Comment string  `json:"comment"`
            Gstin string  `json:"gstin"`
            PickUpCustomerDetails map[string]interface{}  `json:"pick_up_customer_details"`
            CheckoutMode string  `json:"checkout_mode"`
            GiftDetails ArticleGiftDetailSchema  `json:"gift_details"`
         
    }
    
    // CartMetaResponse ...
    type CartMetaResponse struct {

        
            Message string  `json:"message"`
         
    }
    
    // CartMetaMissingResponse ...
    type CartMetaMissingResponse struct {

        
            Errors []string  `json:"errors"`
         
    }
    
    // CartDeliveryModesResponse ...
    type CartDeliveryModesResponse struct {

        
            PickupStores []float64  `json:"pickup_stores"`
            AvailableModes []string  `json:"available_modes"`
         
    }
    
    // PickupStoreDetail ...
    type PickupStoreDetail struct {

        
            Name string  `json:"name"`
            StoreCode string  `json:"store_code"`
            State string  `json:"state"`
            Address string  `json:"address"`
            Phone string  `json:"phone"`
            Area string  `json:"area"`
            Pincode float64  `json:"pincode"`
            AreaCodeSlug string  `json:"area_code_slug"`
            Landmark string  `json:"landmark"`
            City string  `json:"city"`
            Email string  `json:"email"`
            UID float64  `json:"uid"`
            ID float64  `json:"id"`
            AreaCode string  `json:"area_code"`
            Country string  `json:"country"`
            AddressType string  `json:"address_type"`
         
    }
    
    // StoreDetailsResponse ...
    type StoreDetailsResponse struct {

        
            Items []PickupStoreDetail  `json:"items"`
         
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

        
            CreatedOn string  `json:"created_on"`
            User map[string]interface{}  `json:"user"`
            Source map[string]interface{}  `json:"source"`
            Token string  `json:"token"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // SharedCart ...
    type SharedCart struct {

        
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            Comment string  `json:"comment"`
            CartID float64  `json:"cart_id"`
            LastModified string  `json:"last_modified"`
            UID string  `json:"uid"`
            CouponText string  `json:"coupon_text"`
            Items []CartProductInfo  `json:"items"`
            IsValid bool  `json:"is_valid"`
            ID string  `json:"id"`
            Currency CartCurrencySchema  `json:"currency"`
            BuyNow bool  `json:"buy_now"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            CheckoutMode string  `json:"checkout_mode"`
            BreakupValues CartBreakupSchema  `json:"breakup_values"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            Message string  `json:"message"`
            Gstin string  `json:"gstin"`
            SharedCartDetails SharedCartDetails  `json:"shared_cart_details"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
         
    }
    
    // SharedCartResponse ...
    type SharedCartResponse struct {

        
            Cart SharedCart  `json:"cart"`
            Error string  `json:"error"`
         
    }
    

    
    // PincodeErrorSchemaResponse ...
    type PincodeErrorSchemaResponse struct {

        
            Value string  `json:"value"`
            Message string  `json:"message"`
            Type string  `json:"type"`
         
    }
    
    // PincodeParentsResponse ...
    type PincodeParentsResponse struct {

        
            UID string  `json:"uid"`
            Name string  `json:"name"`
            SubType string  `json:"sub_type"`
            DisplayName string  `json:"display_name"`
         
    }
    
    // PincodeMetaResponse ...
    type PincodeMetaResponse struct {

        
            InternalZoneID float64  `json:"internal_zone_id"`
            Zone string  `json:"zone"`
         
    }
    
    // PincodeLatLongData ...
    type PincodeLatLongData struct {

        
            Type string  `json:"type"`
            Coordinates []string  `json:"coordinates"`
         
    }
    
    // CountryMetaResponse ...
    type CountryMetaResponse struct {

        
            IsdCode string  `json:"isd_code"`
            CountryCode string  `json:"country_code"`
         
    }
    
    // PincodeDataResponse ...
    type PincodeDataResponse struct {

        
            Parents []PincodeParentsResponse  `json:"parents"`
            DisplayName string  `json:"display_name"`
            Error PincodeErrorSchemaResponse  `json:"error"`
            SubType string  `json:"sub_type"`
            UID string  `json:"uid"`
            Name string  `json:"name"`
            Meta PincodeMetaResponse  `json:"meta"`
            LatLong PincodeLatLongData  `json:"lat_long"`
            MetaCode CountryMetaResponse  `json:"meta_code"`
         
    }
    
    // PincodeApiResponse ...
    type PincodeApiResponse struct {

        
            Error PincodeErrorSchemaResponse  `json:"error"`
            Data []PincodeDataResponse  `json:"data"`
            Success bool  `json:"success"`
         
    }
    
    // TATCategoryRequest ...
    type TATCategoryRequest struct {

        
            ID float64  `json:"id"`
            Level string  `json:"level"`
         
    }
    
    // TATArticlesRequest ...
    type TATArticlesRequest struct {

        
            ManufacturingTime float64  `json:"manufacturing_time"`
            Category TATCategoryRequest  `json:"category"`
            ManufacturingTimeUnit string  `json:"manufacturing_time_unit"`
         
    }
    
    // TATLocationDetailsRequest ...
    type TATLocationDetailsRequest struct {

        
            FromPincode string  `json:"from_pincode"`
            Articles []TATArticlesRequest  `json:"articles"`
            FulfillmentID float64  `json:"fulfillment_id"`
         
    }
    
    // TATViewRequest ...
    type TATViewRequest struct {

        
            Source string  `json:"source"`
            Identifier string  `json:"identifier"`
            ToPincode string  `json:"to_pincode"`
            Journey string  `json:"journey"`
            Action string  `json:"action"`
            LocationDetails []TATLocationDetailsRequest  `json:"location_details"`
         
    }
    
    // TATErrorSchemaResponse ...
    type TATErrorSchemaResponse struct {

        
            Value string  `json:"value"`
            Message string  `json:"message"`
            Type string  `json:"type"`
         
    }
    
    // TATFormattedResponse ...
    type TATFormattedResponse struct {

        
            Max string  `json:"max"`
            Min string  `json:"min"`
         
    }
    
    // TATTimestampResponse ...
    type TATTimestampResponse struct {

        
            Max float64  `json:"max"`
            Min float64  `json:"min"`
         
    }
    
    // TATPromiseResponse ...
    type TATPromiseResponse struct {

        
            Formatted TATFormattedResponse  `json:"formatted"`
            Timestamp TATTimestampResponse  `json:"timestamp"`
         
    }
    
    // TATArticlesResponse ...
    type TATArticlesResponse struct {

        
            ManufacturingTimeUnit string  `json:"manufacturing_time_unit"`
            Error TATErrorSchemaResponse  `json:"error"`
            IsCodAvailable bool  `json:"is_cod_available"`
            ManufacturingTimeSeconds float64  `json:"_manufacturing_time_seconds"`
            Promise TATPromiseResponse  `json:"promise"`
            ManufacturingTime float64  `json:"manufacturing_time"`
            Category TATCategoryRequest  `json:"category"`
         
    }
    
    // TATLocationDetailsResponse ...
    type TATLocationDetailsResponse struct {

        
            FromPincode string  `json:"from_pincode"`
            Articles []TATArticlesResponse  `json:"articles"`
            FulfillmentID float64  `json:"fulfillment_id"`
         
    }
    
    // TATViewResponse ...
    type TATViewResponse struct {

        
            StormbreakerUUID string  `json:"stormbreaker_uuid"`
            RequestUUID string  `json:"request_uuid"`
            Error TATErrorSchemaResponse  `json:"error"`
            Source string  `json:"source"`
            IsCodAvailable bool  `json:"is_cod_available"`
            PaymentMode string  `json:"payment_mode"`
            Identifier string  `json:"identifier"`
            ToPincode string  `json:"to_pincode"`
            Journey string  `json:"journey"`
            ToCity string  `json:"to_city"`
            Success bool  `json:"success"`
            Action string  `json:"action"`
            LocationDetails []TATLocationDetailsResponse  `json:"location_details"`
         
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
            DisplayName string  `json:"display_name"`
            ParentID string  `json:"parent_id"`
            Logistics LogisticsResponse  `json:"logistics"`
            SubType string  `json:"sub_type"`
            IsActive bool  `json:"is_active"`
            UID string  `json:"uid"`
            Name string  `json:"name"`
            Meta CountryMetaResponse  `json:"meta"`
         
    }
    
    // CountryListResponse ...
    type CountryListResponse struct {

        
            Results []CountryEntityResponse  `json:"results"`
         
    }
    
    // GetZoneFromPincodeViewRequest ...
    type GetZoneFromPincodeViewRequest struct {

        
            Country string  `json:"country"`
            Pincode string  `json:"pincode"`
         
    }
    
    // GetZoneFromPincodeViewResponse ...
    type GetZoneFromPincodeViewResponse struct {

        
            Zones []string  `json:"zones"`
            ServiceabilityType string  `json:"serviceability_type"`
         
    }
    
    // ReAssignStoreRequest ...
    type ReAssignStoreRequest struct {

        
            Articles []map[string]interface{}  `json:"articles"`
            Identifier string  `json:"identifier"`
            ToPincode string  `json:"to_pincode"`
            Configuration map[string]interface{}  `json:"configuration"`
            IgnoredLocations []string  `json:"ignored_locations"`
         
    }
    
    // ReAssignStoreResponse ...
    type ReAssignStoreResponse struct {

        
            Articles []map[string]interface{}  `json:"articles"`
            ToPincode string  `json:"to_pincode"`
            Error map[string]interface{}  `json:"error"`
            Success bool  `json:"success"`
         
    }
    
