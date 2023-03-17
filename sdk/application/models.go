package application



    
    // Meta ...
    type Meta struct {

        
            Source string  `json:"source"`
         
    }
    
    // Media ...
    type Media struct {

        
            Alt string  `json:"alt"`
            URL string  `json:"url"`
            Meta Meta  `json:"meta"`
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

        
            Page ProductListingActionPage  `json:"page"`
            Type string  `json:"type"`
         
    }
    
    // ProductBrand ...
    type ProductBrand struct {

        
            Logo Media  `json:"logo"`
            Description string  `json:"description"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            Action ProductListingAction  `json:"action"`
         
    }
    
    // ProductCategoryMap ...
    type ProductCategoryMap struct {

        
            L3 ProductBrand  `json:"l3"`
            L1 ProductBrand  `json:"l1"`
            L2 ProductBrand  `json:"l2"`
         
    }
    
    // ApplicationItemMOQ ...
    type ApplicationItemMOQ struct {

        
            Minimum float64  `json:"minimum"`
            IncrementUnit float64  `json:"increment_unit"`
            Maximum float64  `json:"maximum"`
         
    }
    
    // ApplicationItemSEO ...
    type ApplicationItemSEO struct {

        
            Title string  `json:"title"`
            Description string  `json:"description"`
         
    }
    
    // CustomMetaFields ...
    type CustomMetaFields struct {

        
            Value string  `json:"value"`
            Key string  `json:"key"`
         
    }
    
    // ProductDetailAttribute ...
    type ProductDetailAttribute struct {

        
            Value string  `json:"value"`
            Key string  `json:"key"`
            Type string  `json:"type"`
         
    }
    
    // ProductDetailGroupedAttribute ...
    type ProductDetailGroupedAttribute struct {

        
            Title string  `json:"title"`
            Details []ProductDetailAttribute  `json:"details"`
         
    }
    
    // Price ...
    type Price struct {

        
            Max float64  `json:"max"`
            CurrencyCode string  `json:"currency_code"`
            Min float64  `json:"min"`
            CurrencySymbol string  `json:"currency_symbol"`
         
    }
    
    // ProductListingPrice ...
    type ProductListingPrice struct {

        
            Marked Price  `json:"marked"`
            Effective Price  `json:"effective"`
         
    }
    
    // NetQuantity ...
    type NetQuantity struct {

        
            Value float64  `json:"value"`
            Unit string  `json:"unit"`
         
    }
    
    // ProductDetail ...
    type ProductDetail struct {

        
            Tryouts []string  `json:"tryouts"`
            HasVariant bool  `json:"has_variant"`
            CategoryMap ProductCategoryMap  `json:"category_map"`
            UID float64  `json:"uid"`
            Color string  `json:"color"`
            Slug string  `json:"slug"`
            Name string  `json:"name"`
            Moq ApplicationItemMOQ  `json:"moq"`
            ShortDescription string  `json:"short_description"`
            ProductOnlineDate string  `json:"product_online_date"`
            ItemType string  `json:"item_type"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Seo ApplicationItemSEO  `json:"seo"`
            IsDependent bool  `json:"is_dependent"`
            Attributes map[string]interface{}  `json:"attributes"`
            Similars []string  `json:"similars"`
            Action ProductListingAction  `json:"action"`
            Type string  `json:"type"`
            RatingCount float64  `json:"rating_count"`
            Rating float64  `json:"rating"`
            Medias []Media  `json:"medias"`
            CustomMeta []CustomMetaFields  `json:"_custom_meta"`
            Brand ProductBrand  `json:"brand"`
            TeaserTag string  `json:"teaser_tag"`
            Categories []ProductBrand  `json:"categories"`
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
            Description string  `json:"description"`
            Price ProductListingPrice  `json:"price"`
            ImageNature string  `json:"image_nature"`
            Tags []string  `json:"tags"`
            Discount string  `json:"discount"`
            Highlights []string  `json:"highlights"`
            NetQuantity NetQuantity  `json:"net_quantity"`
            ItemCode string  `json:"item_code"`
         
    }
    
    // ErrorResponse ...
    type ErrorResponse struct {

        
            Error string  `json:"error"`
         
    }
    
    // ProductSizeStores ...
    type ProductSizeStores struct {

        
            Count float64  `json:"count"`
         
    }
    
    // Dimension ...
    type Dimension struct {

        
            Width float64  `json:"width"`
            Length float64  `json:"length"`
            IsDefault bool  `json:"is_default"`
            Unit string  `json:"unit"`
            Height float64  `json:"height"`
         
    }
    
    // Weight ...
    type Weight struct {

        
            IsDefault bool  `json:"is_default"`
            Unit string  `json:"unit"`
            Shipping float64  `json:"shipping"`
         
    }
    
    // ProductSize ...
    type ProductSize struct {

        
            Value string  `json:"value"`
            Display string  `json:"display"`
            Dimension Dimension  `json:"dimension"`
            Weight Weight  `json:"weight"`
            SellerIdentifiers []string  `json:"seller_identifiers"`
            IsAvailable bool  `json:"is_available"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // ColumnHeader ...
    type ColumnHeader struct {

        
            Value string  `json:"value"`
            Convertable bool  `json:"convertable"`
         
    }
    
    // ColumnHeaders ...
    type ColumnHeaders struct {

        
            Col2 ColumnHeader  `json:"col_2"`
            Col4 ColumnHeader  `json:"col_4"`
            Col3 ColumnHeader  `json:"col_3"`
            Col5 ColumnHeader  `json:"col_5"`
            Col1 ColumnHeader  `json:"col_1"`
            Col6 ColumnHeader  `json:"col_6"`
         
    }
    
    // SizeChartValues ...
    type SizeChartValues struct {

        
            Col2 string  `json:"col_2"`
            Col4 string  `json:"col_4"`
            Col3 string  `json:"col_3"`
            Col5 string  `json:"col_5"`
            Col1 string  `json:"col_1"`
            Col6 string  `json:"col_6"`
         
    }
    
    // SizeChart ...
    type SizeChart struct {

        
            SizeTip string  `json:"size_tip"`
            Description string  `json:"description"`
            Headers ColumnHeaders  `json:"headers"`
            Image string  `json:"image"`
            Unit string  `json:"unit"`
            Sizes []SizeChartValues  `json:"sizes"`
            Title string  `json:"title"`
         
    }
    
    // ProductSizes ...
    type ProductSizes struct {

        
            Sellable bool  `json:"sellable"`
            Stores ProductSizeStores  `json:"stores"`
            Price ProductListingPrice  `json:"price"`
            Discount string  `json:"discount"`
            Sizes []ProductSize  `json:"sizes"`
            SizeChart SizeChart  `json:"size_chart"`
            MultiSize bool  `json:"multi_size"`
         
    }
    
    // AttributeDetail ...
    type AttributeDetail struct {

        
            Logo string  `json:"logo"`
            Display string  `json:"display"`
            Description string  `json:"description"`
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

        
            AttributesMetadata []AttributeMetadata  `json:"attributes_metadata"`
            Subtitle string  `json:"subtitle"`
            Title string  `json:"title"`
            Items []ProductDetail  `json:"items"`
         
    }
    
    // ProductFrequentlyComparedSimilarResponse ...
    type ProductFrequentlyComparedSimilarResponse struct {

        
            Similars ProductCompareResponse  `json:"similars"`
         
    }
    
    // ProductVariantItemResponse ...
    type ProductVariantItemResponse struct {

        
            Value string  `json:"value"`
            ColorName string  `json:"color_name"`
            IsAvailable bool  `json:"is_available"`
            Medias []Media  `json:"medias"`
            CustomMeta []CustomMetaFields  `json:"_custom_meta"`
            Color string  `json:"color"`
            UID float64  `json:"uid"`
            Slug string  `json:"slug"`
            Name string  `json:"name"`
            Action ProductListingAction  `json:"action"`
         
    }
    
    // ProductVariantResponse ...
    type ProductVariantResponse struct {

        
            DisplayType string  `json:"display_type"`
            Items []ProductVariantItemResponse  `json:"items"`
            Key string  `json:"key"`
            Header string  `json:"header"`
         
    }
    
    // ProductVariantsResponse ...
    type ProductVariantsResponse struct {

        
            Variants []ProductVariantResponse  `json:"variants"`
         
    }
    
    // StoreDetail ...
    type StoreDetail struct {

        
            Code string  `json:"code"`
            ID float64  `json:"id"`
            Name string  `json:"name"`
            City string  `json:"city"`
         
    }
    
    // CompanyDetail ...
    type CompanyDetail struct {

        
            ID float64  `json:"id"`
            Name string  `json:"name"`
         
    }
    
    // ProductStockPrice ...
    type ProductStockPrice struct {

        
            Marked float64  `json:"marked"`
            Effective float64  `json:"effective"`
            Currency string  `json:"currency"`
         
    }
    
    // Seller ...
    type Seller struct {

        
            Name string  `json:"name"`
            Count float64  `json:"count"`
            UID float64  `json:"uid"`
         
    }
    
    // ProductStockStatusItem ...
    type ProductStockStatusItem struct {

        
            Size string  `json:"size"`
            ItemID float64  `json:"item_id"`
            Store StoreDetail  `json:"store"`
            Company CompanyDetail  `json:"company"`
            Price ProductStockPrice  `json:"price"`
            Seller Seller  `json:"seller"`
            UID string  `json:"uid"`
            Identifier map[string]interface{}  `json:"identifier"`
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
    
    // ProductFiltersValue ...
    type ProductFiltersValue struct {

        
            Max float64  `json:"max"`
            Value string  `json:"value"`
            IsSelected bool  `json:"is_selected"`
            SelectedMin float64  `json:"selected_min"`
            Display string  `json:"display"`
            SelectedMax float64  `json:"selected_max"`
            CurrencyCode string  `json:"currency_code"`
            DisplayFormat string  `json:"display_format"`
            Min float64  `json:"min"`
            Count float64  `json:"count"`
            QueryFormat string  `json:"query_format"`
            CurrencySymbol string  `json:"currency_symbol"`
         
    }
    
    // ProductFiltersKey ...
    type ProductFiltersKey struct {

        
            Kind string  `json:"kind"`
            Logo string  `json:"logo"`
            Name string  `json:"name"`
            Display string  `json:"display"`
         
    }
    
    // ProductFilters ...
    type ProductFilters struct {

        
            Values []ProductFiltersValue  `json:"values"`
            Key ProductFiltersKey  `json:"key"`
         
    }
    
    // ProductVariantListingResponse ...
    type ProductVariantListingResponse struct {

        
            Items []ProductVariantItemResponse  `json:"items"`
            Header string  `json:"header"`
            DisplayType string  `json:"display_type"`
            Key string  `json:"key"`
            Total float64  `json:"total"`
         
    }
    
    // ProductListingDetail ...
    type ProductListingDetail struct {

        
            Tryouts []string  `json:"tryouts"`
            HasVariant bool  `json:"has_variant"`
            CategoryMap ProductCategoryMap  `json:"category_map"`
            UID float64  `json:"uid"`
            Color string  `json:"color"`
            Variants []ProductVariantListingResponse  `json:"variants"`
            Slug string  `json:"slug"`
            Name string  `json:"name"`
            Moq ApplicationItemMOQ  `json:"moq"`
            Identifiers []string  `json:"identifiers"`
            ShortDescription string  `json:"short_description"`
            ProductOnlineDate string  `json:"product_online_date"`
            ItemType string  `json:"item_type"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Seo ApplicationItemSEO  `json:"seo"`
            Sizes []string  `json:"sizes"`
            IsDependent bool  `json:"is_dependent"`
            Attributes map[string]interface{}  `json:"attributes"`
            Similars []string  `json:"similars"`
            Action ProductListingAction  `json:"action"`
            Type string  `json:"type"`
            RatingCount float64  `json:"rating_count"`
            Rating float64  `json:"rating"`
            Medias []Media  `json:"medias"`
            CustomMeta []CustomMetaFields  `json:"_custom_meta"`
            Brand ProductBrand  `json:"brand"`
            TeaserTag string  `json:"teaser_tag"`
            Sellable bool  `json:"sellable"`
            Categories []ProductBrand  `json:"categories"`
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
            Description string  `json:"description"`
            Price ProductListingPrice  `json:"price"`
            ImageNature string  `json:"image_nature"`
            Tags []string  `json:"tags"`
            Discount string  `json:"discount"`
            Highlights []string  `json:"highlights"`
            NetQuantity NetQuantity  `json:"net_quantity"`
            ItemCode string  `json:"item_code"`
         
    }
    
    // ProductSortOn ...
    type ProductSortOn struct {

        
            Value string  `json:"value"`
            IsSelected bool  `json:"is_selected"`
            Name string  `json:"name"`
         
    }
    
    // ProductListingResponse ...
    type ProductListingResponse struct {

        
            Filters []ProductFilters  `json:"filters"`
            Items []ProductListingDetail  `json:"items"`
            Page Page  `json:"page"`
            SortOn []ProductSortOn  `json:"sort_on"`
         
    }
    
    // ImageUrls ...
    type ImageUrls struct {

        
            Landscape Media  `json:"landscape"`
            Portrait Media  `json:"portrait"`
         
    }
    
    // BrandItem ...
    type BrandItem struct {

        
            Logo Media  `json:"logo"`
            Banners ImageUrls  `json:"banners"`
            Description string  `json:"description"`
            Discount string  `json:"discount"`
            UID float64  `json:"uid"`
            Departments []string  `json:"departments"`
            Slug string  `json:"slug"`
            Name string  `json:"name"`
            Action ProductListingAction  `json:"action"`
         
    }
    
    // BrandListingResponse ...
    type BrandListingResponse struct {

        
            Items []BrandItem  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // BrandDetailResponse ...
    type BrandDetailResponse struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Logo Media  `json:"logo"`
            Banners ImageUrls  `json:"banners"`
            Description string  `json:"description"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // CategoryBanner ...
    type CategoryBanner struct {

        
            Landscape Media  `json:"landscape"`
            Portrait Media  `json:"portrait"`
         
    }
    
    // ThirdLevelChild ...
    type ThirdLevelChild struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Banners ImageUrls  `json:"banners"`
            Childs []map[string]interface{}  `json:"childs"`
            UID float64  `json:"uid"`
            Slug string  `json:"slug"`
            Name string  `json:"name"`
            Action ProductListingAction  `json:"action"`
         
    }
    
    // SecondLevelChild ...
    type SecondLevelChild struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Banners ImageUrls  `json:"banners"`
            Childs []ThirdLevelChild  `json:"childs"`
            UID float64  `json:"uid"`
            Slug string  `json:"slug"`
            Name string  `json:"name"`
            Action ProductListingAction  `json:"action"`
         
    }
    
    // Child ...
    type Child struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Banners ImageUrls  `json:"banners"`
            Childs []SecondLevelChild  `json:"childs"`
            UID float64  `json:"uid"`
            Slug string  `json:"slug"`
            Name string  `json:"name"`
            Action ProductListingAction  `json:"action"`
         
    }
    
    // CategoryItems ...
    type CategoryItems struct {

        
            Banners CategoryBanner  `json:"banners"`
            Childs []Child  `json:"childs"`
            UID float64  `json:"uid"`
            Slug string  `json:"slug"`
            Name string  `json:"name"`
            Action ProductListingAction  `json:"action"`
         
    }
    
    // DepartmentCategoryTree ...
    type DepartmentCategoryTree struct {

        
            Items []CategoryItems  `json:"items"`
            Department string  `json:"department"`
         
    }
    
    // DepartmentIdentifier ...
    type DepartmentIdentifier struct {

        
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
         
    }
    
    // CategoryListingResponse ...
    type CategoryListingResponse struct {

        
            Data []DepartmentCategoryTree  `json:"data"`
            Departments []DepartmentIdentifier  `json:"departments"`
         
    }
    
    // CategoryMetaResponse ...
    type CategoryMetaResponse struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Logo Media  `json:"logo"`
            Banners ImageUrls  `json:"banners"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // HomeListingResponse ...
    type HomeListingResponse struct {

        
            Items []ProductListingDetail  `json:"items"`
            Page Page  `json:"page"`
            Message string  `json:"message"`
         
    }
    
    // Department ...
    type Department struct {

        
            PriorityOrder float64  `json:"priority_order"`
            Logo Media  `json:"logo"`
            UID float64  `json:"uid"`
            Slug string  `json:"slug"`
            Name string  `json:"name"`
         
    }
    
    // DepartmentResponse ...
    type DepartmentResponse struct {

        
            Items []Department  `json:"items"`
         
    }
    
    // AutocompleteItem ...
    type AutocompleteItem struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Logo Media  `json:"logo"`
            Display string  `json:"display"`
            Type string  `json:"type"`
            Action ProductListingAction  `json:"action"`
         
    }
    
    // AutoCompleteResponse ...
    type AutoCompleteResponse struct {

        
            Items []AutocompleteItem  `json:"items"`
         
    }
    
    // CollectionListingFilterTag ...
    type CollectionListingFilterTag struct {

        
            IsSelected bool  `json:"is_selected"`
            Name string  `json:"name"`
            Display string  `json:"display"`
         
    }
    
    // CollectionListingFilterType ...
    type CollectionListingFilterType struct {

        
            IsSelected bool  `json:"is_selected"`
            Name string  `json:"name"`
            Display string  `json:"display"`
         
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

        
            Query []CollectionQuery  `json:"query"`
            Banners ImageUrls  `json:"banners"`
            Priority float64  `json:"priority"`
            SortOn string  `json:"sort_on"`
            UID string  `json:"uid"`
            Slug string  `json:"slug"`
            Name string  `json:"name"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            AllowFacets bool  `json:"allow_facets"`
            Meta map[string]interface{}  `json:"meta"`
            Action ProductListingAction  `json:"action"`
            Type string  `json:"type"`
            IsActive bool  `json:"is_active"`
            AppID string  `json:"app_id"`
            Badge map[string]interface{}  `json:"badge"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Logo Media  `json:"logo"`
            AllowSort bool  `json:"allow_sort"`
            Description string  `json:"description"`
            Tag []string  `json:"tag"`
            Schedule map[string]interface{}  `json:"_schedule"`
            Cron map[string]interface{}  `json:"cron"`
         
    }
    
    // GetCollectionListingResponse ...
    type GetCollectionListingResponse struct {

        
            Filters CollectionListingFilter  `json:"filters"`
            Items []GetCollectionDetailNest  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // CollectionDetailResponse ...
    type CollectionDetailResponse struct {

        
            Query []CollectionQuery  `json:"query"`
            Banners ImageUrls  `json:"banners"`
            Priority float64  `json:"priority"`
            SortOn string  `json:"sort_on"`
            Slug string  `json:"slug"`
            Name string  `json:"name"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            AllowFacets bool  `json:"allow_facets"`
            Meta map[string]interface{}  `json:"meta"`
            Type string  `json:"type"`
            IsActive bool  `json:"is_active"`
            AppID string  `json:"app_id"`
            Badge map[string]interface{}  `json:"badge"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Logo Media  `json:"logo"`
            AllowSort bool  `json:"allow_sort"`
            Description string  `json:"description"`
            Tag []string  `json:"tag"`
            Schedule map[string]interface{}  `json:"_schedule"`
            Cron map[string]interface{}  `json:"cron"`
         
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

        
            State string  `json:"state"`
            LatLong LatLong  `json:"lat_long"`
            Pincode float64  `json:"pincode"`
            Country string  `json:"country"`
            UID float64  `json:"uid"`
            City string  `json:"city"`
            StoreEmail string  `json:"store_email"`
            Name string  `json:"name"`
            StoreCode string  `json:"store_code"`
            Address string  `json:"address"`
         
    }
    
    // StoreListingResponse ...
    type StoreListingResponse struct {

        
            Items []Store  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // StoreDepartments ...
    type StoreDepartments struct {

        
            PriorityOrder float64  `json:"priority_order"`
            Logo string  `json:"logo"`
            UID float64  `json:"uid"`
            Slug string  `json:"slug"`
            Name string  `json:"name"`
         
    }
    
    // SellerPhoneNumber ...
    type SellerPhoneNumber struct {

        
            CountryCode float64  `json:"country_code"`
            Number string  `json:"number"`
         
    }
    
    // StoreManagerSerializer ...
    type StoreManagerSerializer struct {

        
            Name string  `json:"name"`
            MobileNo SellerPhoneNumber  `json:"mobile_no"`
            Email string  `json:"email"`
         
    }
    
    // CompanyStore ...
    type CompanyStore struct {

        
            Name string  `json:"name"`
            CompanyType string  `json:"company_type"`
            UID float64  `json:"uid"`
            BusinessType string  `json:"business_type"`
         
    }
    
    // StoreAddressSerializer ...
    type StoreAddressSerializer struct {

        
            Longitude float64  `json:"longitude"`
            Landmark string  `json:"landmark"`
            Address1 string  `json:"address1"`
            State string  `json:"state"`
            Address2 string  `json:"address2"`
            Pincode float64  `json:"pincode"`
            Country string  `json:"country"`
            City string  `json:"city"`
            Latitude float64  `json:"latitude"`
         
    }
    
    // AppStore ...
    type AppStore struct {

        
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            Manager StoreManagerSerializer  `json:"manager"`
            Company CompanyStore  `json:"company"`
            UID float64  `json:"uid"`
            Departments []StoreDepartments  `json:"departments"`
            Name string  `json:"name"`
            Address StoreAddressSerializer  `json:"address"`
         
    }
    
    // ApplicationStoreListing ...
    type ApplicationStoreListing struct {

        
            Filters []StoreDepartments  `json:"filters"`
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

        
            Open bool  `json:"open"`
            Opening Time  `json:"opening"`
            Closing Time  `json:"closing"`
            Weekday string  `json:"weekday"`
         
    }
    
    // StoreDetails ...
    type StoreDetails struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            Manager StoreManagerSerializer  `json:"manager"`
            Company CompanyStore  `json:"company"`
            UID float64  `json:"uid"`
            Departments []StoreDepartments  `json:"departments"`
            Name string  `json:"name"`
            Address StoreAddressSerializer  `json:"address"`
            Timing []StoreTiming  `json:"timing"`
         
    }
    
    // UserDetail ...
    type UserDetail struct {

        
            SuperUser bool  `json:"super_user"`
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
            Contact string  `json:"contact"`
         
    }
    
    // ProductDetails ...
    type ProductDetails struct {

        
            HasVariant bool  `json:"has_variant"`
            BrandUID float64  `json:"brand_uid"`
            IsSet bool  `json:"is_set"`
            HsnCode float64  `json:"hsn_code"`
            Slug string  `json:"slug"`
            Name string  `json:"name"`
            Images []string  `json:"images"`
            Media []map[string]interface{}  `json:"media"`
            ShortDescription string  `json:"short_description"`
            Identifier map[string]interface{}  `json:"identifier"`
            Attributes map[string]interface{}  `json:"attributes"`
            RatingCount float64  `json:"rating_count"`
            Rating float64  `json:"rating"`
            OutOfStock bool  `json:"out_of_stock"`
            GroupedAttributes map[string]interface{}  `json:"grouped_attributes"`
            Description string  `json:"description"`
            CountryOfOrigin string  `json:"country_of_origin"`
            ImageNature string  `json:"image_nature"`
            Highlights []string  `json:"highlights"`
            TemplateTag string  `json:"template_tag"`
            ItemCode string  `json:"item_code"`
         
    }
    
    // ProductGroupPrice ...
    type ProductGroupPrice struct {

        
            MaxEffective float64  `json:"max_effective"`
            MinEffective float64  `json:"min_effective"`
            MaxMarked float64  `json:"max_marked"`
            MinMarked float64  `json:"min_marked"`
            Currency string  `json:"currency"`
         
    }
    
    // Size ...
    type Size struct {

        
            IsAvailable bool  `json:"is_available"`
            Value string  `json:"value"`
            Display string  `json:"display"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // ProductInGroup ...
    type ProductInGroup struct {

        
            MinQuantity float64  `json:"min_quantity"`
            ProductDetails ProductDetails  `json:"product_details"`
            MaxQuantity float64  `json:"max_quantity"`
            ProductUID float64  `json:"product_uid"`
            AutoSelect bool  `json:"auto_select"`
            AutoAddToCart bool  `json:"auto_add_to_cart"`
            Price ProductGroupPrice  `json:"price"`
            Sizes []Size  `json:"sizes"`
            AllowRemove bool  `json:"allow_remove"`
         
    }
    
    // ProductGroupingModel ...
    type ProductGroupingModel struct {

        
            Logo string  `json:"logo"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            CompanyID float64  `json:"company_id"`
            VerifiedBy UserDetail  `json:"verified_by"`
            CreatedBy UserDetail  `json:"created_by"`
            ID interface{}  `json:"_id"`
            PageVisibility []string  `json:"page_visibility"`
            ModifiedBy UserDetail  `json:"modified_by"`
            IsActive bool  `json:"is_active"`
            Products []ProductInGroup  `json:"products"`
            VerifiedOn string  `json:"verified_on"`
            CreatedOn string  `json:"created_on"`
            Slug string  `json:"slug"`
            Name string  `json:"name"`
            Choice string  `json:"choice"`
            Meta map[string]interface{}  `json:"meta"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // ProductBundle ...
    type ProductBundle struct {

        
            Items []ProductGroupingModel  `json:"items"`
         
    }
    
    // StoreV3 ...
    type StoreV3 struct {

        
            Name string  `json:"name"`
            Count float64  `json:"count"`
            UID float64  `json:"uid"`
         
    }
    
    // StrategyWiseListingSchemaV3 ...
    type StrategyWiseListingSchemaV3 struct {

        
            Tat float64  `json:"tat"`
            Pincode float64  `json:"pincode"`
            Distance float64  `json:"distance"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // SellerV3 ...
    type SellerV3 struct {

        
            Name string  `json:"name"`
            Count float64  `json:"count"`
            UID float64  `json:"uid"`
         
    }
    
    // ReturnConfigSchemaV3 ...
    type ReturnConfigSchemaV3 struct {

        
            Returnable bool  `json:"returnable"`
            Unit string  `json:"unit"`
            Time float64  `json:"time"`
         
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
    
    // ProductStockUnitPriceV3 ...
    type ProductStockUnitPriceV3 struct {

        
            CurrencyCode string  `json:"currency_code"`
            Unit string  `json:"unit"`
            Price float64  `json:"price"`
            CurrencySymbol string  `json:"currency_symbol"`
         
    }
    
    // DetailsSchemaV3 ...
    type DetailsSchemaV3 struct {

        
            Value string  `json:"value"`
            Key string  `json:"key"`
            Type string  `json:"type"`
         
    }
    
    // MarketPlaceSttributesSchemaV3 ...
    type MarketPlaceSttributesSchemaV3 struct {

        
            Title string  `json:"title"`
            Details []DetailsSchemaV3  `json:"details"`
         
    }
    
    // ArticleAssignmentV3 ...
    type ArticleAssignmentV3 struct {

        
            Level string  `json:"level"`
            Strategy string  `json:"strategy"`
         
    }
    
    // SellerGroupAttributes ...
    type SellerGroupAttributes struct {

        
            Title string  `json:"title"`
            Details []DetailsSchemaV3  `json:"details"`
         
    }
    
    // ProductStockPriceV3 ...
    type ProductStockPriceV3 struct {

        
            Marked float64  `json:"marked"`
            Effective float64  `json:"effective"`
            Currency string  `json:"currency"`
         
    }
    
    // ProductSizePriceResponseV3 ...
    type ProductSizePriceResponseV3 struct {

        
            LongLat []float64  `json:"long_lat"`
            SellerCount float64  `json:"seller_count"`
            Store StoreV3  `json:"store"`
            Pincode float64  `json:"pincode"`
            Quantity float64  `json:"quantity"`
            ItemType string  `json:"item_type"`
            SpecialBadge string  `json:"special_badge"`
            StrategyWiseListing []StrategyWiseListingSchemaV3  `json:"strategy_wise_listing"`
            Seller SellerV3  `json:"seller"`
            ReturnConfig ReturnConfigSchemaV3  `json:"return_config"`
            Set ProductSetV3  `json:"set"`
            PricePerUnit ProductStockUnitPriceV3  `json:"price_per_unit"`
            MarketplaceAttributes []MarketPlaceSttributesSchemaV3  `json:"marketplace_attributes"`
            ArticleID string  `json:"article_id"`
            ArticleAssignment ArticleAssignmentV3  `json:"article_assignment"`
            GroupedAttributes []SellerGroupAttributes  `json:"grouped_attributes"`
            Price ProductStockPriceV3  `json:"price"`
            PricePerPiece ProductStockPriceV3  `json:"price_per_piece"`
            Discount string  `json:"discount"`
            IsGift bool  `json:"is_gift"`
            IsCod bool  `json:"is_cod"`
         
    }
    
    // ProductSizeSellerFilterSchemaV3 ...
    type ProductSizeSellerFilterSchemaV3 struct {

        
            Value string  `json:"value"`
            IsSelected bool  `json:"is_selected"`
            Name string  `json:"name"`
         
    }
    
    // ProductSizeSellersResponseV3 ...
    type ProductSizeSellersResponseV3 struct {

        
            Items []ProductSizePriceResponseV3  `json:"items"`
            Page Page  `json:"page"`
            SortOn []ProductSizeSellerFilterSchemaV3  `json:"sort_on"`
         
    }
    

    
    // RawBreakup ...
    type RawBreakup struct {

        
            Vog float64  `json:"vog"`
            GstCharges float64  `json:"gst_charges"`
            ConvenienceFee float64  `json:"convenience_fee"`
            Discount float64  `json:"discount"`
            Coupon float64  `json:"coupon"`
            YouSaved float64  `json:"you_saved"`
            Total float64  `json:"total"`
            MrpTotal float64  `json:"mrp_total"`
            Subtotal float64  `json:"subtotal"`
            DeliveryCharge float64  `json:"delivery_charge"`
            CodCharge float64  `json:"cod_charge"`
            FyndCash float64  `json:"fynd_cash"`
         
    }
    
    // DisplayBreakup ...
    type DisplayBreakup struct {

        
            Display string  `json:"display"`
            Value float64  `json:"value"`
            CurrencySymbol string  `json:"currency_symbol"`
            Key string  `json:"key"`
            CurrencyCode string  `json:"currency_code"`
            Message []string  `json:"message"`
         
    }
    
    // LoyaltyPoints ...
    type LoyaltyPoints struct {

        
            IsApplied bool  `json:"is_applied"`
            Total float64  `json:"total"`
            Description string  `json:"description"`
            Applicable float64  `json:"applicable"`
         
    }
    
    // CouponBreakup ...
    type CouponBreakup struct {

        
            Description string  `json:"description"`
            MinimumCartValue float64  `json:"minimum_cart_value"`
            IsApplied bool  `json:"is_applied"`
            Value float64  `json:"value"`
            Title string  `json:"title"`
            CouponValue float64  `json:"coupon_value"`
            Code string  `json:"code"`
            CouponType string  `json:"coupon_type"`
            UID string  `json:"uid"`
            Message string  `json:"message"`
            MaxDiscountValue float64  `json:"max_discount_value"`
            SubTitle string  `json:"sub_title"`
            Type string  `json:"type"`
         
    }
    
    // CartBreakup ...
    type CartBreakup struct {

        
            Raw RawBreakup  `json:"raw"`
            Display []DisplayBreakup  `json:"display"`
            LoyaltyPoints LoyaltyPoints  `json:"loyalty_points"`
            Coupon CouponBreakup  `json:"coupon"`
         
    }
    
    // ProductPrice ...
    type ProductPrice struct {

        
            CurrencySymbol string  `json:"currency_symbol"`
            Marked float64  `json:"marked"`
            Selling float64  `json:"selling"`
            AddOn float64  `json:"add_on"`
            Effective float64  `json:"effective"`
            CurrencyCode string  `json:"currency_code"`
         
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
    
    // BaseInfo ...
    type BaseInfo struct {

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // BasePrice ...
    type BasePrice struct {

        
            Effective float64  `json:"effective"`
            CurrencyCode string  `json:"currency_code"`
            CurrencySymbol string  `json:"currency_symbol"`
            Marked float64  `json:"marked"`
         
    }
    
    // ArticlePriceInfo ...
    type ArticlePriceInfo struct {

        
            Base BasePrice  `json:"base"`
            Converted BasePrice  `json:"converted"`
         
    }
    
    // ProductArticle ...
    type ProductArticle struct {

        
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            Seller BaseInfo  `json:"seller"`
            Quantity float64  `json:"quantity"`
            Price ArticlePriceInfo  `json:"price"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Size string  `json:"size"`
            ProductGroupTags []string  `json:"product_group_tags"`
            UID string  `json:"uid"`
            Store BaseInfo  `json:"store"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            Type string  `json:"type"`
         
    }
    
    // PromoMeta ...
    type PromoMeta struct {

        
            Message string  `json:"message"`
         
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
    
    // NetQuantity ...
    type NetQuantity struct {

        
            Value string  `json:"value"`
            Unit string  `json:"unit"`
         
    }
    
    // CategoryInfo ...
    type CategoryInfo struct {

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // ProductImage ...
    type ProductImage struct {

        
            SecureURL string  `json:"secure_url"`
            URL string  `json:"url"`
            AspectRatio string  `json:"aspect_ratio"`
         
    }
    
    // CartProduct ...
    type CartProduct struct {

        
            Action ProductAction  `json:"action"`
            NetQuantity NetQuantity  `json:"net_quantity"`
            Slug string  `json:"slug"`
            Name string  `json:"name"`
            Categories []CategoryInfo  `json:"categories"`
            Images []ProductImage  `json:"images"`
            Brand BaseInfo  `json:"brand"`
            ItemCode string  `json:"item_code"`
            UID float64  `json:"uid"`
            Type string  `json:"type"`
         
    }
    
    // FreeGiftItem ...
    type FreeGiftItem struct {

        
            ItemPriceDetails map[string]interface{}  `json:"item_price_details"`
            ItemSlug string  `json:"item_slug"`
            ItemName string  `json:"item_name"`
            ItemID float64  `json:"item_id"`
            ItemBrandName string  `json:"item_brand_name"`
            ItemImagesURL []string  `json:"item_images_url"`
         
    }
    
    // AppliedFreeArticles ...
    type AppliedFreeArticles struct {

        
            ArticleID string  `json:"article_id"`
            FreeGiftItemDetails FreeGiftItem  `json:"free_gift_item_details"`
            ParentItemIdentifier string  `json:"parent_item_identifier"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // Ownership ...
    type Ownership struct {

        
            PayableCategory string  `json:"payable_category"`
            PayableBy string  `json:"payable_by"`
         
    }
    
    // DiscountRulesApp ...
    type DiscountRulesApp struct {

        
            RawOffer map[string]interface{}  `json:"raw_offer"`
            MatchedBuyRules []string  `json:"matched_buy_rules"`
            ItemCriteria map[string]interface{}  `json:"item_criteria"`
            Offer map[string]interface{}  `json:"offer"`
         
    }
    
    // BuyRules ...
    type BuyRules struct {

        
            ItemCriteria map[string]interface{}  `json:"item_criteria"`
            CartConditions map[string]interface{}  `json:"cart_conditions"`
         
    }
    
    // AppliedPromotion ...
    type AppliedPromotion struct {

        
            MrpPromotion bool  `json:"mrp_promotion"`
            AppliedFreeArticles []AppliedFreeArticles  `json:"applied_free_articles"`
            Amount float64  `json:"amount"`
            PromotionGroup string  `json:"promotion_group"`
            Ownership Ownership  `json:"ownership"`
            PromotionType string  `json:"promotion_type"`
            PromotionName string  `json:"promotion_name"`
            OfferText string  `json:"offer_text"`
            PromoID string  `json:"promo_id"`
            ArticleQuantity float64  `json:"article_quantity"`
            DiscountRules []DiscountRulesApp  `json:"discount_rules"`
            BuyRules []BuyRules  `json:"buy_rules"`
         
    }
    
    // ProductAvailability ...
    type ProductAvailability struct {

        
            Sizes []string  `json:"sizes"`
            OutOfStock bool  `json:"out_of_stock"`
            OtherStoreQuantity float64  `json:"other_store_quantity"`
            IsValid bool  `json:"is_valid"`
            Deliverable bool  `json:"deliverable"`
         
    }
    
    // CartProductInfo ...
    type CartProductInfo struct {

        
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            Quantity float64  `json:"quantity"`
            PricePerUnit ProductPriceInfo  `json:"price_per_unit"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
            BulkOffer map[string]interface{}  `json:"bulk_offer"`
            Article ProductArticle  `json:"article"`
            Discount string  `json:"discount"`
            Key string  `json:"key"`
            PromoMeta PromoMeta  `json:"promo_meta"`
            Product CartProduct  `json:"product"`
            CouponMessage string  `json:"coupon_message"`
            Price ProductPriceInfo  `json:"price"`
            PromotionsApplied []AppliedPromotion  `json:"promotions_applied"`
            Availability ProductAvailability  `json:"availability"`
            IsSet bool  `json:"is_set"`
            Message string  `json:"message"`
         
    }
    
    // CartCurrency ...
    type CartCurrency struct {

        
            Code string  `json:"code"`
            Symbol string  `json:"symbol"`
         
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
    
    // PaymentSelectionLock ...
    type PaymentSelectionLock struct {

        
            Enabled bool  `json:"enabled"`
            DefaultOptions string  `json:"default_options"`
            PaymentIdentifier string  `json:"payment_identifier"`
         
    }
    
    // CartDetailResponse ...
    type CartDetailResponse struct {

        
            RestrictCheckout bool  `json:"restrict_checkout"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            BuyNow bool  `json:"buy_now"`
            PanNo string  `json:"pan_no"`
            CheckoutMode string  `json:"checkout_mode"`
            LastModified string  `json:"last_modified"`
            CouponText string  `json:"coupon_text"`
            Items []CartProductInfo  `json:"items"`
            Comment string  `json:"comment"`
            Currency CartCurrency  `json:"currency"`
            PanConfig map[string]interface{}  `json:"pan_config"`
            IsValid bool  `json:"is_valid"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            Gstin string  `json:"gstin"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            ID string  `json:"id"`
            Message string  `json:"message"`
         
    }
    
    // AddProductCart ...
    type AddProductCart struct {

        
            ArticleAssignment map[string]interface{}  `json:"article_assignment"`
            Display string  `json:"display"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            Quantity float64  `json:"quantity"`
            ArticleID string  `json:"article_id"`
            Pos bool  `json:"pos"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ProductGroupTags []string  `json:"product_group_tags"`
            ItemID float64  `json:"item_id"`
            SellerID float64  `json:"seller_id"`
            StoreID float64  `json:"store_id"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            ItemSize string  `json:"item_size"`
         
    }
    
    // AddCartRequest ...
    type AddCartRequest struct {

        
            NewCart bool  `json:"new_cart"`
            Items []AddProductCart  `json:"items"`
         
    }
    
    // AddCartDetailResponse ...
    type AddCartDetailResponse struct {

        
            Cart CartDetailResponse  `json:"cart"`
            Success bool  `json:"success"`
            Partial bool  `json:"partial"`
            Message string  `json:"message"`
         
    }
    
    // UpdateProductCart ...
    type UpdateProductCart struct {

        
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            Quantity float64  `json:"quantity"`
            ArticleID string  `json:"article_id"`
            ItemIndex float64  `json:"item_index"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ItemID float64  `json:"item_id"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            ItemSize string  `json:"item_size"`
         
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
    
    // DeleteCartDetailResponse ...
    type DeleteCartDetailResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // CartItemCountResponse ...
    type CartItemCountResponse struct {

        
            UserCartItemsCount float64  `json:"user_cart_items_count"`
         
    }
    
    // PageCoupon ...
    type PageCoupon struct {

        
            TotalItemCount float64  `json:"total_item_count"`
            HasPrevious bool  `json:"has_previous"`
            Current float64  `json:"current"`
            Total float64  `json:"total"`
            HasNext bool  `json:"has_next"`
         
    }
    
    // Coupon ...
    type Coupon struct {

        
            Description string  `json:"description"`
            MinimumCartValue float64  `json:"minimum_cart_value"`
            CouponValue float64  `json:"coupon_value"`
            IsApplied bool  `json:"is_applied"`
            CouponCode string  `json:"coupon_code"`
            CouponType string  `json:"coupon_type"`
            Title string  `json:"title"`
            MaxDiscountValue float64  `json:"max_discount_value"`
            ExpiresOn string  `json:"expires_on"`
            SubTitle string  `json:"sub_title"`
            IsApplicable bool  `json:"is_applicable"`
            Message string  `json:"message"`
         
    }
    
    // GetCouponResponse ...
    type GetCouponResponse struct {

        
            Page PageCoupon  `json:"page"`
            AvailableCouponList []Coupon  `json:"available_coupon_list"`
         
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

        
            CurrencySymbol string  `json:"currency_symbol"`
            Marked float64  `json:"marked"`
            Effective float64  `json:"effective"`
            CurrencyCode string  `json:"currency_code"`
            BulkEffective float64  `json:"bulk_effective"`
         
    }
    
    // OfferItem ...
    type OfferItem struct {

        
            AutoApplied bool  `json:"auto_applied"`
            Quantity float64  `json:"quantity"`
            Price OfferPrice  `json:"price"`
            Best bool  `json:"best"`
            Total float64  `json:"total"`
            Margin float64  `json:"margin"`
            Type string  `json:"type"`
         
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

        
            Longitude float64  `json:"longitude"`
            Latitude float64  `json:"latitude"`
         
    }
    
    // Address ...
    type Address struct {

        
            CountryIsoCode string  `json:"country_iso_code"`
            CountryCode string  `json:"country_code"`
            AreaCodeSlug string  `json:"area_code_slug"`
            IsDefaultAddress bool  `json:"is_default_address"`
            Tags []string  `json:"tags"`
            CountryPhoneCode string  `json:"country_phone_code"`
            Area string  `json:"area"`
            Name string  `json:"name"`
            CheckoutMode string  `json:"checkout_mode"`
            Phone string  `json:"phone"`
            UserID string  `json:"user_id"`
            Email string  `json:"email"`
            Meta map[string]interface{}  `json:"meta"`
            City string  `json:"city"`
            Landmark string  `json:"landmark"`
            GeoLocation GeoLocation  `json:"geo_location"`
            IsActive bool  `json:"is_active"`
            ID string  `json:"id"`
            Country string  `json:"country"`
            CreatedByUserID string  `json:"created_by_user_id"`
            AreaCode string  `json:"area_code"`
            GoogleMapPoint map[string]interface{}  `json:"google_map_point"`
            State string  `json:"state"`
            Address string  `json:"address"`
            AddressType string  `json:"address_type"`
         
    }
    
    // GetAddressesResponse ...
    type GetAddressesResponse struct {

        
            Address []Address  `json:"address"`
         
    }
    
    // SaveAddressResponse ...
    type SaveAddressResponse struct {

        
            IsDefaultAddress bool  `json:"is_default_address"`
            Success bool  `json:"success"`
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

        
            ID string  `json:"id"`
            IsDeleted bool  `json:"is_deleted"`
         
    }
    
    // SelectCartAddressRequest ...
    type SelectCartAddressRequest struct {

        
            BillingAddressID string  `json:"billing_address_id"`
            ID string  `json:"id"`
            CartID string  `json:"cart_id"`
         
    }
    
    // UpdateCartPaymentRequest ...
    type UpdateCartPaymentRequest struct {

        
            MerchantCode string  `json:"merchant_code"`
            AggregatorName string  `json:"aggregator_name"`
            AddressID string  `json:"address_id"`
            PaymentIdentifier string  `json:"payment_identifier"`
            ID string  `json:"id"`
            PaymentMode string  `json:"payment_mode"`
         
    }
    
    // CouponValidity ...
    type CouponValidity struct {

        
            Valid bool  `json:"valid"`
            DisplayMessageEn string  `json:"display_message_en"`
            Discount float64  `json:"discount"`
            Code string  `json:"code"`
            Title string  `json:"title"`
         
    }
    
    // PaymentCouponValidate ...
    type PaymentCouponValidate struct {

        
            Success bool  `json:"success"`
            CouponValidity CouponValidity  `json:"coupon_validity"`
            Message string  `json:"message"`
         
    }
    
    // ShipmentResponse ...
    type ShipmentResponse struct {

        
            ShipmentType string  `json:"shipment_type"`
            DpID string  `json:"dp_id"`
            BoxType string  `json:"box_type"`
            Items []CartProductInfo  `json:"items"`
            OrderType string  `json:"order_type"`
            DpOptions map[string]interface{}  `json:"dp_options"`
            FulfillmentType string  `json:"fulfillment_type"`
            FulfillmentID float64  `json:"fulfillment_id"`
            Shipments float64  `json:"shipments"`
            Promise ShipmentPromise  `json:"promise"`
         
    }
    
    // CartShipmentsResponse ...
    type CartShipmentsResponse struct {

        
            RestrictCheckout bool  `json:"restrict_checkout"`
            Error bool  `json:"error"`
            LastModified string  `json:"last_modified"`
            CouponText string  `json:"coupon_text"`
            Comment string  `json:"comment"`
            Currency CartCurrency  `json:"currency"`
            IsValid bool  `json:"is_valid"`
            CartID float64  `json:"cart_id"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            Shipments []ShipmentResponse  `json:"shipments"`
            ID string  `json:"id"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            BuyNow bool  `json:"buy_now"`
            CheckoutMode string  `json:"checkout_mode"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            UID string  `json:"uid"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            Gstin string  `json:"gstin"`
            Message string  `json:"message"`
         
    }
    
    // CartCheckoutCustomMeta ...
    type CartCheckoutCustomMeta struct {

        
            Key string  `json:"key"`
            Value string  `json:"value"`
         
    }
    
    // StaffCheckout ...
    type StaffCheckout struct {

        
            LastName string  `json:"last_name"`
            FirstName string  `json:"first_name"`
            EmployeeCode string  `json:"employee_code"`
            User string  `json:"user"`
            ID string  `json:"_id"`
         
    }
    
    // CartCheckoutDetailRequest ...
    type CartCheckoutDetailRequest struct {

        
            Meta map[string]interface{}  `json:"meta"`
            BillingAddressID string  `json:"billing_address_id"`
            OrderingStore float64  `json:"ordering_store"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
            MerchantCode string  `json:"merchant_code"`
            PaymentParams map[string]interface{}  `json:"payment_params"`
            AddressID string  `json:"address_id"`
            PaymentAutoConfirm bool  `json:"payment_auto_confirm"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            CallbackURL string  `json:"callback_url"`
            Aggregator string  `json:"aggregator"`
            CustomMeta []CartCheckoutCustomMeta  `json:"custom_meta"`
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
            Staff StaffCheckout  `json:"staff"`
            PaymentIdentifier string  `json:"payment_identifier"`
            PaymentMode string  `json:"payment_mode"`
         
    }
    
    // CheckCart ...
    type CheckCart struct {

        
            CodMessage string  `json:"cod_message"`
            Items []CartProductInfo  `json:"items"`
            StoreEmps []map[string]interface{}  `json:"store_emps"`
            Comment string  `json:"comment"`
            Currency CartCurrency  `json:"currency"`
            CartID float64  `json:"cart_id"`
            UserType string  `json:"user_type"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            Success bool  `json:"success"`
            DeliveryCharges float64  `json:"delivery_charges"`
            CodCharges float64  `json:"cod_charges"`
            CheckoutMode string  `json:"checkout_mode"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            CodAvailable bool  `json:"cod_available"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            DeliveryChargeOrderValue float64  `json:"delivery_charge_order_value"`
            Message string  `json:"message"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            StoreCode string  `json:"store_code"`
            LastModified string  `json:"last_modified"`
            CouponText string  `json:"coupon_text"`
            IsValid bool  `json:"is_valid"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            ID string  `json:"id"`
            BuyNow bool  `json:"buy_now"`
            ErrorMessage string  `json:"error_message"`
            UID string  `json:"uid"`
            OrderID string  `json:"order_id"`
            Gstin string  `json:"gstin"`
         
    }
    
    // CartCheckoutResponse ...
    type CartCheckoutResponse struct {

        
            Success bool  `json:"success"`
            PaymentConfirmURL string  `json:"payment_confirm_url"`
            Data map[string]interface{}  `json:"data"`
            AppInterceptURL string  `json:"app_intercept_url"`
            CallbackURL string  `json:"callback_url"`
            OrderID string  `json:"order_id"`
            Cart CheckCart  `json:"cart"`
            Message string  `json:"message"`
         
    }
    
    // CartMetaRequest ...
    type CartMetaRequest struct {

        
            CheckoutMode string  `json:"checkout_mode"`
            Comment string  `json:"comment"`
            Gstin string  `json:"gstin"`
            PickUpCustomerDetails map[string]interface{}  `json:"pick_up_customer_details"`
         
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

        
            ShareURL string  `json:"share_url"`
            Token string  `json:"token"`
         
    }
    
    // SharedCartDetails ...
    type SharedCartDetails struct {

        
            Meta map[string]interface{}  `json:"meta"`
            Token string  `json:"token"`
            CreatedOn string  `json:"created_on"`
            Source map[string]interface{}  `json:"source"`
            User map[string]interface{}  `json:"user"`
         
    }
    
    // SharedCart ...
    type SharedCart struct {

        
            Items []CartProductInfo  `json:"items"`
            Comment string  `json:"comment"`
            Currency CartCurrency  `json:"currency"`
            CartID float64  `json:"cart_id"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            CheckoutMode string  `json:"checkout_mode"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            Message string  `json:"message"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            SharedCartDetails SharedCartDetails  `json:"shared_cart_details"`
            LastModified string  `json:"last_modified"`
            CouponText string  `json:"coupon_text"`
            IsValid bool  `json:"is_valid"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            ID string  `json:"id"`
            BuyNow bool  `json:"buy_now"`
            UID string  `json:"uid"`
            Gstin string  `json:"gstin"`
         
    }
    
    // SharedCartResponse ...
    type SharedCartResponse struct {

        
            Cart SharedCart  `json:"cart"`
            Error string  `json:"error"`
         
    }
    
    // FreeGiftItems ...
    type FreeGiftItems struct {

        
            ItemPriceDetails map[string]interface{}  `json:"item_price_details"`
            ItemSlug string  `json:"item_slug"`
            ItemName string  `json:"item_name"`
            ItemID float64  `json:"item_id"`
            ItemBrandName string  `json:"item_brand_name"`
            ItemImagesURL []string  `json:"item_images_url"`
         
    }
    
    // PromotionOffer ...
    type PromotionOffer struct {

        
            Description string  `json:"description"`
            PromotionGroup string  `json:"promotion_group"`
            OfferText string  `json:"offer_text"`
            FreeGiftItems []FreeGiftItems  `json:"free_gift_items"`
            ValidTill string  `json:"valid_till"`
            ID string  `json:"id"`
            DiscountRules []map[string]interface{}  `json:"discount_rules"`
            BuyRules map[string]interface{}  `json:"buy_rules"`
         
    }
    
    // PromotionOffersResponse ...
    type PromotionOffersResponse struct {

        
            AvailablePromotions []PromotionOffer  `json:"available_promotions"`
         
    }
    
    // OperationErrorResponse ...
    type OperationErrorResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // LadderPrice ...
    type LadderPrice struct {

        
            OfferPrice float64  `json:"offer_price"`
            CurrencySymbol string  `json:"currency_symbol"`
            Marked float64  `json:"marked"`
            Effective float64  `json:"effective"`
            CurrencyCode string  `json:"currency_code"`
         
    }
    
    // LadderOfferItem ...
    type LadderOfferItem struct {

        
            MinQuantity float64  `json:"min_quantity"`
            Margin float64  `json:"margin"`
            Price LadderPrice  `json:"price"`
            MaxQuantity float64  `json:"max_quantity"`
            Type string  `json:"type"`
         
    }
    
    // LadderPriceOffer ...
    type LadderPriceOffer struct {

        
            Description string  `json:"description"`
            PromotionGroup string  `json:"promotion_group"`
            OfferPrices []LadderOfferItem  `json:"offer_prices"`
            CalculateOn string  `json:"calculate_on"`
            OfferText string  `json:"offer_text"`
            FreeGiftItems []FreeGiftItems  `json:"free_gift_items"`
            ValidTill string  `json:"valid_till"`
            ID string  `json:"id"`
            DiscountRules []map[string]interface{}  `json:"discount_rules"`
            BuyRules map[string]interface{}  `json:"buy_rules"`
         
    }
    
    // CurrencyInfo ...
    type CurrencyInfo struct {

        
            Code string  `json:"code"`
            Symbol string  `json:"symbol"`
         
    }
    
    // LadderPriceOffers ...
    type LadderPriceOffers struct {

        
            AvailableOffers []LadderPriceOffer  `json:"available_offers"`
            Currency CurrencyInfo  `json:"currency"`
         
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

        
            Items []string  `json:"items"`
         
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
            Source TagSourceSchema  `json:"__source"`
         
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

        
            API string  `json:"api"`
            MerchantID string  `json:"merchant_id"`
            ConfigType string  `json:"config_type"`
            MerchantKey string  `json:"merchant_key"`
            Key string  `json:"key"`
            VerifyAPI string  `json:"verify_api"`
            Pin string  `json:"pin"`
            UserID string  `json:"user_id"`
            Sdk bool  `json:"sdk"`
            Secret string  `json:"secret"`
         
    }
    
    // AggregatorsConfigDetailResponse ...
    type AggregatorsConfigDetailResponse struct {

        
            Env string  `json:"env"`
            Rupifi AggregatorConfigDetail  `json:"rupifi"`
            Juspay AggregatorConfigDetail  `json:"juspay"`
            Mswipe AggregatorConfigDetail  `json:"mswipe"`
            Ccavenue AggregatorConfigDetail  `json:"ccavenue"`
            Success bool  `json:"success"`
            Payumoney AggregatorConfigDetail  `json:"payumoney"`
            Razorpay AggregatorConfigDetail  `json:"razorpay"`
            Simpl AggregatorConfigDetail  `json:"simpl"`
            Stripe AggregatorConfigDetail  `json:"stripe"`
         
    }
    
    // ErrorCodeAndDescription ...
    type ErrorCodeAndDescription struct {

        
            Code string  `json:"code"`
            Description string  `json:"description"`
         
    }
    
    // HttpErrorCodeAndResponse ...
    type HttpErrorCodeAndResponse struct {

        
            Error ErrorCodeAndDescription  `json:"error"`
            Success bool  `json:"success"`
         
    }
    
    // AttachCardRequest ...
    type AttachCardRequest struct {

        
            Nickname string  `json:"nickname"`
            CardID string  `json:"card_id"`
            NameOnCard string  `json:"name_on_card"`
            Refresh bool  `json:"refresh"`
         
    }
    
    // AttachCardsResponse ...
    type AttachCardsResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
            Data map[string]interface{}  `json:"data"`
         
    }
    
    // CardPaymentGateway ...
    type CardPaymentGateway struct {

        
            API string  `json:"api"`
            Aggregator string  `json:"aggregator"`
            CustomerID string  `json:"customer_id"`
         
    }
    
    // ActiveCardPaymentGatewayResponse ...
    type ActiveCardPaymentGatewayResponse struct {

        
            Cards CardPaymentGateway  `json:"cards"`
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // Card ...
    type Card struct {

        
            CardReference string  `json:"card_reference"`
            Nickname string  `json:"nickname"`
            CardFingerprint string  `json:"card_fingerprint"`
            ExpMonth float64  `json:"exp_month"`
            CardNumber string  `json:"card_number"`
            CardType string  `json:"card_type"`
            CardIssuer string  `json:"card_issuer"`
            CardIsin string  `json:"card_isin"`
            CardID string  `json:"card_id"`
            Expired bool  `json:"expired"`
            CardToken string  `json:"card_token"`
            ExpYear float64  `json:"exp_year"`
            CardBrand string  `json:"card_brand"`
            CardName string  `json:"card_name"`
            AggregatorName string  `json:"aggregator_name"`
            CompliantWithTokenisationGuidelines bool  `json:"compliant_with_tokenisation_guidelines"`
            CardBrandImage string  `json:"card_brand_image"`
         
    }
    
    // ListCardsResponse ...
    type ListCardsResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
            Data []Card  `json:"data"`
         
    }
    
    // DeletehCardRequest ...
    type DeletehCardRequest struct {

        
            CardID string  `json:"card_id"`
         
    }
    
    // DeleteCardsResponse ...
    type DeleteCardsResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // ValidateCustomerRequest ...
    type ValidateCustomerRequest struct {

        
            Aggregator string  `json:"aggregator"`
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
            MerchantParams map[string]interface{}  `json:"merchant_params"`
            PhoneNumber string  `json:"phone_number"`
            TransactionAmountInPaise float64  `json:"transaction_amount_in_paise"`
            OrderItems []map[string]interface{}  `json:"order_items"`
            Payload string  `json:"payload"`
         
    }
    
    // ValidateCustomerResponse ...
    type ValidateCustomerResponse struct {

        
            Error map[string]interface{}  `json:"error"`
            Message string  `json:"message"`
            Success bool  `json:"success"`
            Data map[string]interface{}  `json:"data"`
         
    }
    
    // ChargeCustomerRequest ...
    type ChargeCustomerRequest struct {

        
            OrderID string  `json:"order_id"`
            Aggregator string  `json:"aggregator"`
            TransactionToken string  `json:"transaction_token"`
            Verified bool  `json:"verified"`
            Amount float64  `json:"amount"`
         
    }
    
    // ChargeCustomerResponse ...
    type ChargeCustomerResponse struct {

        
            OrderID string  `json:"order_id"`
            Aggregator string  `json:"aggregator"`
            Status string  `json:"status"`
            CartID string  `json:"cart_id"`
            Success bool  `json:"success"`
            DeliveryAddressID string  `json:"delivery_address_id"`
            Message string  `json:"message"`
         
    }
    
    // PaymentInitializationRequest ...
    type PaymentInitializationRequest struct {

        
            OrderID string  `json:"order_id"`
            Timeout float64  `json:"timeout"`
            Email string  `json:"email"`
            Aggregator string  `json:"aggregator"`
            Vpa string  `json:"vpa"`
            Contact string  `json:"contact"`
            RazorpayPaymentID string  `json:"razorpay_payment_id"`
            CustomerID string  `json:"customer_id"`
            Method string  `json:"method"`
            DeviceID string  `json:"device_id"`
            MerchantOrderID string  `json:"merchant_order_id"`
            Currency string  `json:"currency"`
            Amount float64  `json:"amount"`
         
    }
    
    // PaymentInitializationResponse ...
    type PaymentInitializationResponse struct {

        
            Timeout float64  `json:"timeout"`
            Aggregator string  `json:"aggregator"`
            UpiPollURL string  `json:"upi_poll_url"`
            Vpa string  `json:"vpa"`
            AggregatorOrderID string  `json:"aggregator_order_id"`
            Status string  `json:"status"`
            BqrImage string  `json:"bqr_image"`
            Success bool  `json:"success"`
            RazorpayPaymentID string  `json:"razorpay_payment_id"`
            CustomerID string  `json:"customer_id"`
            Method string  `json:"method"`
            DeviceID string  `json:"device_id"`
            PollingURL string  `json:"polling_url"`
            MerchantOrderID string  `json:"merchant_order_id"`
            VirtualID string  `json:"virtual_id"`
            Currency string  `json:"currency"`
            Amount float64  `json:"amount"`
         
    }
    
    // PaymentStatusUpdateRequest ...
    type PaymentStatusUpdateRequest struct {

        
            OrderID string  `json:"order_id"`
            Aggregator string  `json:"aggregator"`
            Email string  `json:"email"`
            Vpa string  `json:"vpa"`
            Contact string  `json:"contact"`
            Status string  `json:"status"`
            CustomerID string  `json:"customer_id"`
            Method string  `json:"method"`
            DeviceID string  `json:"device_id"`
            MerchantOrderID string  `json:"merchant_order_id"`
            Currency string  `json:"currency"`
            Amount float64  `json:"amount"`
         
    }
    
    // PaymentStatusUpdateResponse ...
    type PaymentStatusUpdateResponse struct {

        
            Status string  `json:"status"`
            Success bool  `json:"success"`
            RedirectURL string  `json:"redirect_url"`
            AggregatorName string  `json:"aggregator_name"`
            Retry bool  `json:"retry"`
         
    }
    
    // PaymentModeLogo ...
    type PaymentModeLogo struct {

        
            Small string  `json:"small"`
            Large string  `json:"large"`
         
    }
    
    // IntentApp ...
    type IntentApp struct {

        
            Code string  `json:"code"`
            Logos PaymentModeLogo  `json:"logos"`
            DisplayName string  `json:"display_name"`
            PackageName string  `json:"package_name"`
         
    }
    
    // IntentAppErrorList ...
    type IntentAppErrorList struct {

        
            Code string  `json:"code"`
            PackageName string  `json:"package_name"`
         
    }
    
    // PaymentModeList ...
    type PaymentModeList struct {

        
            CardReference string  `json:"card_reference"`
            IntentAppErrorList []string  `json:"intent_app_error_list"`
            ExpMonth float64  `json:"exp_month"`
            CardType string  `json:"card_type"`
            LogoURL PaymentModeLogo  `json:"logo_url"`
            ExpYear float64  `json:"exp_year"`
            IntentFlow bool  `json:"intent_flow"`
            Name string  `json:"name"`
            CardName string  `json:"card_name"`
            DisplayName string  `json:"display_name"`
            CodLimitPerOrder float64  `json:"cod_limit_per_order"`
            CardIssuer string  `json:"card_issuer"`
            DisplayPriority float64  `json:"display_priority"`
            CardBrand string  `json:"card_brand"`
            IntentApp []IntentApp  `json:"intent_app"`
            CardFingerprint string  `json:"card_fingerprint"`
            Nickname string  `json:"nickname"`
            IntentAppErrorDictList []IntentAppErrorList  `json:"intent_app_error_dict_list"`
            Code string  `json:"code"`
            CardNumber string  `json:"card_number"`
            RemainingLimit float64  `json:"remaining_limit"`
            CardIsin string  `json:"card_isin"`
            CardToken string  `json:"card_token"`
            FyndVpa string  `json:"fynd_vpa"`
            CompliantWithTokenisationGuidelines bool  `json:"compliant_with_tokenisation_guidelines"`
            RetryCount float64  `json:"retry_count"`
            Timeout float64  `json:"timeout"`
            CodLimit float64  `json:"cod_limit"`
            MerchantCode string  `json:"merchant_code"`
            CardID string  `json:"card_id"`
            Expired bool  `json:"expired"`
            AggregatorName string  `json:"aggregator_name"`
            CardBrandImage string  `json:"card_brand_image"`
         
    }
    
    // RootPaymentMode ...
    type RootPaymentMode struct {

        
            IsPayByCardPl bool  `json:"is_pay_by_card_pl"`
            SaveCard bool  `json:"save_card"`
            AnonymousEnable bool  `json:"anonymous_enable"`
            DisplayName string  `json:"display_name"`
            List []PaymentModeList  `json:"list"`
            AddCardEnabled bool  `json:"add_card_enabled"`
            DisplayPriority float64  `json:"display_priority"`
            Name string  `json:"name"`
            AggregatorName string  `json:"aggregator_name"`
         
    }
    
    // AggregatorRoute ...
    type AggregatorRoute struct {

        
            PaymentFlow string  `json:"payment_flow"`
            PaymentFlowData string  `json:"payment_flow_data"`
            Data map[string]interface{}  `json:"data"`
            APILink string  `json:"api_link"`
         
    }
    
    // PaymentFlow ...
    type PaymentFlow struct {

        
            Rupifi AggregatorRoute  `json:"rupifi"`
            Payubiz AggregatorRoute  `json:"payubiz"`
            Juspay AggregatorRoute  `json:"juspay"`
            Epaylater AggregatorRoute  `json:"epaylater"`
            Mswipe AggregatorRoute  `json:"mswipe"`
            Ccavenue AggregatorRoute  `json:"ccavenue"`
            BqrRazorpay AggregatorRoute  `json:"bqr_razorpay"`
            Razorpay AggregatorRoute  `json:"razorpay"`
            Simpl AggregatorRoute  `json:"simpl"`
            Stripe AggregatorRoute  `json:"stripe"`
            Jiopay AggregatorRoute  `json:"jiopay"`
            Fynd AggregatorRoute  `json:"fynd"`
            UpiRazorpay AggregatorRoute  `json:"upi_razorpay"`
         
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

        
            Success bool  `json:"success"`
            Data RupifiBannerData  `json:"data"`
         
    }
    
    // EpaylaterBannerData ...
    type EpaylaterBannerData struct {

        
            Display bool  `json:"display"`
            Message string  `json:"message"`
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
            DeviceID string  `json:"device_id"`
            RequestType string  `json:"request_type"`
         
    }
    
    // LinkStatus ...
    type LinkStatus struct {

        
            Message string  `json:"message"`
            Status bool  `json:"status"`
         
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

        
            Status string  `json:"status"`
            IsValid bool  `json:"is_valid"`
            CustomerName string  `json:"customer_name"`
            UpiVpa string  `json:"upi_vpa"`
         
    }
    
    // ValidateVPAResponse ...
    type ValidateVPAResponse struct {

        
            Success bool  `json:"success"`
            Data ValidateUPI  `json:"data"`
         
    }
    
    // TransferItemsDetails ...
    type TransferItemsDetails struct {

        
            LogoLarge string  `json:"logo_large"`
            DisplayName string  `json:"display_name"`
            ID float64  `json:"id"`
            LogoSmall string  `json:"logo_small"`
            Name string  `json:"name"`
         
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

        
            TransferMode string  `json:"transfer_mode"`
            Enable bool  `json:"enable"`
         
    }
    
    // UpdateRefundTransferModeResponse ...
    type UpdateRefundTransferModeResponse struct {

        
            Success bool  `json:"success"`
         
    }
    
    // OrderBeneficiaryDetails ...
    type OrderBeneficiaryDetails struct {

        
            BranchName string  `json:"branch_name"`
            IfscCode string  `json:"ifsc_code"`
            ModifiedOn string  `json:"modified_on"`
            Title string  `json:"title"`
            DisplayName string  `json:"display_name"`
            BeneficiaryID string  `json:"beneficiary_id"`
            Comment string  `json:"comment"`
            CreatedOn string  `json:"created_on"`
            AccountHolder string  `json:"account_holder"`
            DelightsUserName string  `json:"delights_user_name"`
            TransferMode string  `json:"transfer_mode"`
            Mobile string  `json:"mobile"`
            AccountNo string  `json:"account_no"`
            Email string  `json:"email"`
            Address string  `json:"address"`
            Subtitle string  `json:"subtitle"`
            BankName string  `json:"bank_name"`
            ID float64  `json:"id"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // OrderBeneficiaryResponse ...
    type OrderBeneficiaryResponse struct {

        
            ShowBeneficiaryDetails bool  `json:"show_beneficiary_details"`
            Beneficiaries []OrderBeneficiaryDetails  `json:"beneficiaries"`
         
    }
    
    // NotFoundResourceError ...
    type NotFoundResourceError struct {

        
            Code string  `json:"code"`
            Description string  `json:"description"`
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

        
            Code string  `json:"code"`
            Description string  `json:"description"`
            Success bool  `json:"success"`
         
    }
    
    // AddBeneficiaryViaOtpVerificationRequest ...
    type AddBeneficiaryViaOtpVerificationRequest struct {

        
            RequestID string  `json:"request_id"`
            HashKey string  `json:"hash_key"`
            Otp string  `json:"otp"`
         
    }
    
    // AddBeneficiaryViaOtpVerificationResponse ...
    type AddBeneficiaryViaOtpVerificationResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // WrongOtpError ...
    type WrongOtpError struct {

        
            IsVerifiedFlag bool  `json:"is_verified_flag"`
            Description string  `json:"description"`
            Success string  `json:"success"`
         
    }
    
    // BeneficiaryModeDetails ...
    type BeneficiaryModeDetails struct {

        
            AccountHolder string  `json:"account_holder"`
            Email string  `json:"email"`
            Address string  `json:"address"`
            Wallet string  `json:"wallet"`
            Vpa string  `json:"vpa"`
            BankName string  `json:"bank_name"`
            BranchName string  `json:"branch_name"`
            Comment string  `json:"comment"`
            Mobile string  `json:"mobile"`
            IfscCode string  `json:"ifsc_code"`
            AccountNo string  `json:"account_no"`
         
    }
    
    // AddBeneficiaryDetailsRequest ...
    type AddBeneficiaryDetailsRequest struct {

        
            OrderID string  `json:"order_id"`
            Delights bool  `json:"delights"`
            Details BeneficiaryModeDetails  `json:"details"`
            RequestID string  `json:"request_id"`
            TransferMode string  `json:"transfer_mode"`
            ShipmentID string  `json:"shipment_id"`
            Otp string  `json:"otp"`
         
    }
    
    // RefundAccountResponse ...
    type RefundAccountResponse struct {

        
            IsVerifiedFlag bool  `json:"is_verified_flag"`
            Message string  `json:"message"`
            Success bool  `json:"success"`
            Data map[string]interface{}  `json:"data"`
         
    }
    
    // BankDetailsForOTP ...
    type BankDetailsForOTP struct {

        
            AccountHolder string  `json:"account_holder"`
            BankName string  `json:"bank_name"`
            BranchName string  `json:"branch_name"`
            IfscCode string  `json:"ifsc_code"`
            AccountNo string  `json:"account_no"`
         
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

        
            RequestID string  `json:"request_id"`
            IsVerifiedFlag string  `json:"is_verified_flag"`
            Success bool  `json:"success"`
         
    }
    
    // SetDefaultBeneficiaryRequest ...
    type SetDefaultBeneficiaryRequest struct {

        
            OrderID string  `json:"order_id"`
            BeneficiaryID string  `json:"beneficiary_id"`
         
    }
    
    // SetDefaultBeneficiaryResponse ...
    type SetDefaultBeneficiaryResponse struct {

        
            Success bool  `json:"success"`
            IsBeneficiarySet bool  `json:"is_beneficiary_set"`
         
    }
    
    // GetPaymentLinkResponse ...
    type GetPaymentLinkResponse struct {

        
            PaymentLinkURL string  `json:"payment_link_url"`
            Success bool  `json:"success"`
            PaymentLinkCurrentStatus string  `json:"payment_link_current_status"`
            ExternalOrderID string  `json:"external_order_id"`
            Amount float64  `json:"amount"`
            MerchantName string  `json:"merchant_name"`
            Message string  `json:"message"`
            PollingTimeout float64  `json:"polling_timeout"`
            StatusCode float64  `json:"status_code"`
         
    }
    
    // ErrorDescription ...
    type ErrorDescription struct {

        
            Msg string  `json:"msg"`
            PaymentTransactionID string  `json:"payment_transaction_id"`
            InvalidID bool  `json:"invalid_id"`
            Expired bool  `json:"expired"`
            MerchantName string  `json:"merchant_name"`
            MerchantOrderID string  `json:"merchant_order_id"`
            Amount float64  `json:"amount"`
            Cancelled bool  `json:"cancelled"`
         
    }
    
    // ErrorResponse ...
    type ErrorResponse struct {

        
            Error ErrorDescription  `json:"error"`
            Message string  `json:"message"`
            Success bool  `json:"success"`
            StatusCode float64  `json:"status_code"`
         
    }
    
    // CreatePaymentLinkMeta ...
    type CreatePaymentLinkMeta struct {

        
            AssignCardID string  `json:"assign_card_id"`
            CartID string  `json:"cart_id"`
            CheckoutMode string  `json:"checkout_mode"`
            Pincode string  `json:"pincode"`
            Amount string  `json:"amount"`
         
    }
    
    // CreatePaymentLinkRequest ...
    type CreatePaymentLinkRequest struct {

        
            Email string  `json:"email"`
            Meta CreatePaymentLinkMeta  `json:"meta"`
            MobileNumber string  `json:"mobile_number"`
            ExternalOrderID string  `json:"external_order_id"`
            Description string  `json:"description"`
            Amount float64  `json:"amount"`
         
    }
    
    // CreatePaymentLinkResponse ...
    type CreatePaymentLinkResponse struct {

        
            PaymentLinkURL string  `json:"payment_link_url"`
            PaymentLinkID string  `json:"payment_link_id"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
            PollingTimeout float64  `json:"polling_timeout"`
            StatusCode float64  `json:"status_code"`
         
    }
    
    // CancelOrResendPaymentLinkRequest ...
    type CancelOrResendPaymentLinkRequest struct {

        
            PaymentLinkID string  `json:"payment_link_id"`
         
    }
    
    // ResendPaymentLinkResponse ...
    type ResendPaymentLinkResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
            PollingTimeout float64  `json:"polling_timeout"`
            StatusCode float64  `json:"status_code"`
         
    }
    
    // CancelPaymentLinkResponse ...
    type CancelPaymentLinkResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
            StatusCode float64  `json:"status_code"`
         
    }
    
    // PollingPaymentLinkResponse ...
    type PollingPaymentLinkResponse struct {

        
            OrderID string  `json:"order_id"`
            Status string  `json:"status"`
            PaymentLinkID string  `json:"payment_link_id"`
            Success bool  `json:"success"`
            RedirectURL string  `json:"redirect_url"`
            Amount float64  `json:"amount"`
            HttpStatus float64  `json:"http_status"`
            Message string  `json:"message"`
            AggregatorName string  `json:"aggregator_name"`
            StatusCode float64  `json:"status_code"`
         
    }
    
    // PaymentMethodsMeta ...
    type PaymentMethodsMeta struct {

        
            PaymentIdentifier string  `json:"payment_identifier"`
            MerchantCode string  `json:"merchant_code"`
            PaymentGateway string  `json:"payment_gateway"`
         
    }
    
    // CreateOrderUserPaymentMethods ...
    type CreateOrderUserPaymentMethods struct {

        
            Mode string  `json:"mode"`
            Name string  `json:"name"`
            Meta PaymentMethodsMeta  `json:"meta"`
         
    }
    
    // CreateOrderUserRequest ...
    type CreateOrderUserRequest struct {

        
            Meta map[string]interface{}  `json:"meta"`
            PaymentMethods CreateOrderUserPaymentMethods  `json:"payment_methods"`
            PaymentLinkID string  `json:"payment_link_id"`
            FailureCallbackURL string  `json:"failure_callback_url"`
            Currency string  `json:"currency"`
            SuccessCallbackURL string  `json:"success_callback_url"`
         
    }
    
    // CreateOrderUserData ...
    type CreateOrderUserData struct {

        
            OrderID string  `json:"order_id"`
            Email string  `json:"email"`
            Aggregator string  `json:"aggregator"`
            Contact string  `json:"contact"`
            CustomerID string  `json:"customer_id"`
            Method string  `json:"method"`
            CallbackURL string  `json:"callback_url"`
            MerchantOrderID string  `json:"merchant_order_id"`
            Currency string  `json:"currency"`
            Amount float64  `json:"amount"`
         
    }
    
    // CreateOrderUserResponse ...
    type CreateOrderUserResponse struct {

        
            OrderID string  `json:"order_id"`
            Success bool  `json:"success"`
            Data CreateOrderUserData  `json:"data"`
            CallbackURL string  `json:"callback_url"`
            PaymentConfirmURL string  `json:"payment_confirm_url"`
            Message string  `json:"message"`
            StatusCode float64  `json:"status_code"`
         
    }
    
    // BalanceDetails ...
    type BalanceDetails struct {

        
            Currency string  `json:"currency"`
            FormattedValue string  `json:"formatted_value"`
            Value float64  `json:"value"`
         
    }
    
    // CreditSummary ...
    type CreditSummary struct {

        
            StatusMessage string  `json:"status_message"`
            Status string  `json:"status"`
            MerchantCustomerRefID string  `json:"merchant_customer_ref_id"`
            Balance BalanceDetails  `json:"balance"`
         
    }
    
    // CustomerCreditSummaryResponse ...
    type CustomerCreditSummaryResponse struct {

        
            Success bool  `json:"success"`
            Data CreditSummary  `json:"data"`
         
    }
    
    // RedirectURL ...
    type RedirectURL struct {

        
            Status bool  `json:"status"`
            SignupURL string  `json:"signup_url"`
         
    }
    
    // RedirectToAggregatorResponse ...
    type RedirectToAggregatorResponse struct {

        
            Success bool  `json:"success"`
            Data RedirectURL  `json:"data"`
         
    }
    
    // CreditDetail ...
    type CreditDetail struct {

        
            Status bool  `json:"status"`
            SignupURL string  `json:"signup_url"`
            IsRegistered bool  `json:"is_registered"`
         
    }
    
    // CheckCreditResponse ...
    type CheckCreditResponse struct {

        
            Success bool  `json:"success"`
            Data CreditDetail  `json:"data"`
         
    }
    
    // DeviceDetails ...
    type DeviceDetails struct {

        
            Os string  `json:"os"`
            IdentificationNumber string  `json:"identification_number"`
            IdentifierType string  `json:"identifier_type"`
            OsVersion string  `json:"os_version"`
            DeviceMake string  `json:"device_make"`
            DeviceModel string  `json:"device_model"`
            DeviceType string  `json:"device_type"`
         
    }
    
    // KYCAddress ...
    type KYCAddress struct {

        
            State string  `json:"state"`
            LandMark string  `json:"land_mark"`
            Pincode string  `json:"pincode"`
            City string  `json:"city"`
            Addressline2 string  `json:"addressline2"`
            Addressline1 string  `json:"addressline1"`
            OwnershipType string  `json:"ownership_type"`
         
    }
    
    // UserPersonalInfoInDetails ...
    type UserPersonalInfoInDetails struct {

        
            Email string  `json:"email"`
            MobileVerified bool  `json:"mobile_verified"`
            Phone string  `json:"phone"`
            MothersName string  `json:"mothers_name"`
            AddressAsPerID KYCAddress  `json:"address_as_per_id"`
            MiddleName string  `json:"middle_name"`
            Gender string  `json:"gender"`
            LastName string  `json:"last_name"`
            FathersName string  `json:"fathers_name"`
            FirstName string  `json:"first_name"`
            Pan string  `json:"pan"`
            VoterID string  `json:"voter_id"`
            Dob string  `json:"dob"`
            DrivingLicense string  `json:"driving_license"`
            Passport string  `json:"passport"`
            EmailVerified bool  `json:"email_verified"`
         
    }
    
    // MarketplaceInfo ...
    type MarketplaceInfo struct {

        
            Name string  `json:"name"`
            DateOfJoining string  `json:"date_of_joining"`
            MembershipID string  `json:"membership_id"`
         
    }
    
    // BusinessDetails ...
    type BusinessDetails struct {

        
            Address KYCAddress  `json:"address"`
            Gstin string  `json:"gstin"`
            ShopAndEstablishment map[string]interface{}  `json:"shop_and_establishment"`
            Pan string  `json:"pan"`
            BusinessOwnershipType string  `json:"business_ownership_type"`
            Fssai string  `json:"fssai"`
            Vintage string  `json:"vintage"`
            Fda string  `json:"fda"`
            Name string  `json:"name"`
            BusinessType string  `json:"business_type"`
            EntityType string  `json:"entity_type"`
         
    }
    
    // CustomerOnboardingRequest ...
    type CustomerOnboardingRequest struct {

        
            Aggregator string  `json:"aggregator"`
            Device DeviceDetails  `json:"device"`
            PersonalInfo UserPersonalInfoInDetails  `json:"personal_info"`
            Mcc string  `json:"mcc"`
            MarketplaceInfo MarketplaceInfo  `json:"marketplace_info"`
            Source string  `json:"source"`
            BusinessInfo BusinessDetails  `json:"business_info"`
         
    }
    
    // OnboardSummary ...
    type OnboardSummary struct {

        
            Status bool  `json:"status"`
            Session map[string]interface{}  `json:"session"`
            RedirectURL string  `json:"redirect_url"`
         
    }
    
    // CustomerOnboardingResponse ...
    type CustomerOnboardingResponse struct {

        
            Success bool  `json:"success"`
            Data OnboardSummary  `json:"data"`
         
    }
    

    
    // OrderPage ...
    type OrderPage struct {

        
            ItemTotal float64  `json:"item_total"`
            Type string  `json:"type"`
            Size float64  `json:"size"`
            HasNext bool  `json:"has_next"`
            Current float64  `json:"current"`
         
    }
    
    // BreakupValues ...
    type BreakupValues struct {

        
            Display string  `json:"display"`
            Name string  `json:"name"`
            Value float64  `json:"value"`
         
    }
    
    // Invoice ...
    type Invoice struct {

        
            UpdatedDate string  `json:"updated_date"`
            InvoiceURL string  `json:"invoice_url"`
            LabelURL string  `json:"label_url"`
         
    }
    
    // FulfillingCompany ...
    type FulfillingCompany struct {

        
            Name string  `json:"name"`
            ID float64  `json:"id"`
         
    }
    
    // ShipmentUserInfo ...
    type ShipmentUserInfo struct {

        
            LastName string  `json:"last_name"`
            Mobile string  `json:"mobile"`
            FirstName string  `json:"first_name"`
            Gender string  `json:"gender"`
         
    }
    
    // DeliveryAddress ...
    type DeliveryAddress struct {

        
            Latitude float64  `json:"latitude"`
            ContactPerson string  `json:"contact_person"`
            Email string  `json:"email"`
            Pincode string  `json:"pincode"`
            Country string  `json:"country"`
            UpdatedAt string  `json:"updated_at"`
            AddressType string  `json:"address_type"`
            Version string  `json:"version"`
            City string  `json:"city"`
            Longitude float64  `json:"longitude"`
            Name string  `json:"name"`
            State string  `json:"state"`
            Landmark string  `json:"landmark"`
            Phone string  `json:"phone"`
            Address1 string  `json:"address1"`
            Address2 string  `json:"address2"`
            Area string  `json:"area"`
            CreatedAt string  `json:"created_at"`
            AddressCategory string  `json:"address_category"`
            Address string  `json:"address"`
         
    }
    
    // NestedTrackingDetails ...
    type NestedTrackingDetails struct {

        
            Time string  `json:"time"`
            Status string  `json:"status"`
            IsCurrent bool  `json:"is_current"`
            IsPassed bool  `json:"is_passed"`
         
    }
    
    // TrackingDetails ...
    type TrackingDetails struct {

        
            TrackingDetails []NestedTrackingDetails  `json:"tracking_details"`
            IsCurrent bool  `json:"is_current"`
            Status string  `json:"status"`
            Time string  `json:"time"`
            IsPassed bool  `json:"is_passed"`
         
    }
    
    // ShipmentTotalDetails ...
    type ShipmentTotalDetails struct {

        
            Pieces float64  `json:"pieces"`
            TotalPrice float64  `json:"total_price"`
            Sizes float64  `json:"sizes"`
         
    }
    
    // FulfillingStore ...
    type FulfillingStore struct {

        
            ID float64  `json:"id"`
            CompanyID float64  `json:"company_id"`
            Code string  `json:"code"`
            CompanyName string  `json:"company_name"`
            Name string  `json:"name"`
         
    }
    
    // CurrentStatus ...
    type CurrentStatus struct {

        
            JourneyType string  `json:"journey_type"`
            Name string  `json:"name"`
            Status string  `json:"status"`
            UpdatedAt string  `json:"updated_at"`
         
    }
    
    // AppliedFreeArticles ...
    type AppliedFreeArticles struct {

        
            ArticleID string  `json:"article_id"`
            ParentItemIdentifier string  `json:"parent_item_identifier"`
            Quantity float64  `json:"quantity"`
            FreeGiftItemDetails map[string]interface{}  `json:"free_gift_item_details"`
         
    }
    
    // AppliedPromos ...
    type AppliedPromos struct {

        
            ArticleQuantity float64  `json:"article_quantity"`
            AppliedFreeArticles []AppliedFreeArticles  `json:"applied_free_articles"`
            PromotionName string  `json:"promotion_name"`
            MrpPromotion bool  `json:"mrp_promotion"`
            PromoID string  `json:"promo_id"`
            Amount float64  `json:"amount"`
            PromotionType string  `json:"promotion_type"`
         
    }
    
    // Identifiers ...
    type Identifiers struct {

        
            Ean string  `json:"ean"`
            SkuCode string  `json:"sku_code"`
         
    }
    
    // FinancialBreakup ...
    type FinancialBreakup struct {

        
            FyndCredits float64  `json:"fynd_credits"`
            PromotionEffectiveDiscount float64  `json:"promotion_effective_discount"`
            GstTag string  `json:"gst_tag"`
            TransferPrice float64  `json:"transfer_price"`
            PriceMarked float64  `json:"price_marked"`
            AmountPaid float64  `json:"amount_paid"`
            RefundCredit float64  `json:"refund_credit"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            ItemName string  `json:"item_name"`
            Identifiers Identifiers  `json:"identifiers"`
            RefundAmount float64  `json:"refund_amount"`
            Discount float64  `json:"discount"`
            CouponValue float64  `json:"coupon_value"`
            GstFee float64  `json:"gst_fee"`
            CashbackApplied float64  `json:"cashback_applied"`
            Cashback float64  `json:"cashback"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            AddedToFyndCash bool  `json:"added_to_fynd_cash"`
            ValueOfGood float64  `json:"value_of_good"`
            DeliveryCharge float64  `json:"delivery_charge"`
            PriceEffective float64  `json:"price_effective"`
            HsnCode string  `json:"hsn_code"`
            Size string  `json:"size"`
            TotalUnits float64  `json:"total_units"`
            CodCharges float64  `json:"cod_charges"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            AmountPaidRoundoff float64  `json:"amount_paid_roundoff"`
         
    }
    
    // ItemBrand ...
    type ItemBrand struct {

        
            Logo string  `json:"logo"`
            Name string  `json:"name"`
         
    }
    
    // Item ...
    type Item struct {

        
            ID float64  `json:"id"`
            SellerIdentifier string  `json:"seller_identifier"`
            Brand ItemBrand  `json:"brand"`
            Code string  `json:"code"`
            Name string  `json:"name"`
            Size string  `json:"size"`
            SlugKey string  `json:"slug_key"`
            Image []string  `json:"image"`
         
    }
    
    // Prices ...
    type Prices struct {

        
            FyndCredits float64  `json:"fynd_credits"`
            PromotionEffectiveDiscount float64  `json:"promotion_effective_discount"`
            TransferPrice float64  `json:"transfer_price"`
            PriceMarked float64  `json:"price_marked"`
            AmountPaid float64  `json:"amount_paid"`
            RefundCredit float64  `json:"refund_credit"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            RefundAmount float64  `json:"refund_amount"`
            CouponValue float64  `json:"coupon_value"`
            Discount float64  `json:"discount"`
            CashbackApplied float64  `json:"cashback_applied"`
            Cashback float64  `json:"cashback"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            AddedToFyndCash bool  `json:"added_to_fynd_cash"`
            ValueOfGood float64  `json:"value_of_good"`
            DeliveryCharge float64  `json:"delivery_charge"`
            PriceEffective float64  `json:"price_effective"`
            CodCharges float64  `json:"cod_charges"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            AmountPaidRoundoff float64  `json:"amount_paid_roundoff"`
         
    }
    
    // Bags ...
    type Bags struct {

        
            ParentPromoBags map[string]interface{}  `json:"parent_promo_bags"`
            CurrentStatus CurrentStatus  `json:"current_status"`
            AppliedPromos []AppliedPromos  `json:"applied_promos"`
            Quantity float64  `json:"quantity"`
            ID float64  `json:"id"`
            SellerIdentifier string  `json:"seller_identifier"`
            LineNumber float64  `json:"line_number"`
            FinancialBreakup []FinancialBreakup  `json:"financial_breakup"`
            DeliveryDate string  `json:"delivery_date"`
            Item Item  `json:"item"`
            ReturnableDate string  `json:"returnable_date"`
            CanReturn bool  `json:"can_return"`
            Prices Prices  `json:"prices"`
            CanCancel bool  `json:"can_cancel"`
         
    }
    
    // ShipmentStatus ...
    type ShipmentStatus struct {

        
            Title string  `json:"title"`
            HexCode string  `json:"hex_code"`
         
    }
    
    // ShipmentPayment ...
    type ShipmentPayment struct {

        
            Mode string  `json:"mode"`
            DisplayName string  `json:"display_name"`
            Status string  `json:"status"`
            Logo string  `json:"logo"`
            PaymentMode string  `json:"payment_mode"`
            Mop string  `json:"mop"`
         
    }
    
    // TimeStampData ...
    type TimeStampData struct {

        
            Min string  `json:"min"`
            Max string  `json:"max"`
         
    }
    
    // Promise ...
    type Promise struct {

        
            ShowPromise bool  `json:"show_promise"`
            Timestamp TimeStampData  `json:"timestamp"`
         
    }
    
    // Shipments ...
    type Shipments struct {

        
            CanBreak map[string]interface{}  `json:"can_break"`
            BreakupValues []BreakupValues  `json:"breakup_values"`
            ShowTrackLink bool  `json:"show_track_link"`
            DpName string  `json:"dp_name"`
            Invoice Invoice  `json:"invoice"`
            FulfillingCompany FulfillingCompany  `json:"fulfilling_company"`
            Comment string  `json:"comment"`
            UserInfo ShipmentUserInfo  `json:"user_info"`
            DeliveryAddress DeliveryAddress  `json:"delivery_address"`
            CanCancel bool  `json:"can_cancel"`
            TrackingDetails []TrackingDetails  `json:"tracking_details"`
            TotalDetails ShipmentTotalDetails  `json:"total_details"`
            ShipmentID string  `json:"shipment_id"`
            DeliveryDate string  `json:"delivery_date"`
            TotalBags float64  `json:"total_bags"`
            TrakingNo string  `json:"traking_no"`
            ReturnableDate string  `json:"returnable_date"`
            CanReturn bool  `json:"can_return"`
            ShipmentCreatedAt string  `json:"shipment_created_at"`
            RefundDetails map[string]interface{}  `json:"refund_details"`
            CustomMeta []map[string]interface{}  `json:"custom_meta"`
            FulfillingStore FulfillingStore  `json:"fulfilling_store"`
            Bags []Bags  `json:"bags"`
            BeneficiaryDetails bool  `json:"beneficiary_details"`
            ShipmentStatus ShipmentStatus  `json:"shipment_status"`
            AwbNo string  `json:"awb_no"`
            OrderType string  `json:"order_type"`
            Payment ShipmentPayment  `json:"payment"`
            SizeInfo map[string]interface{}  `json:"size_info"`
            Promise Promise  `json:"promise"`
            OrderID string  `json:"order_id"`
            ShowDownloadInvoice bool  `json:"show_download_invoice"`
            TrackURL string  `json:"track_url"`
            Prices Prices  `json:"prices"`
            NeedHelpURL string  `json:"need_help_url"`
         
    }
    
    // BagsForReorderArticleAssignment ...
    type BagsForReorderArticleAssignment struct {

        
            Level string  `json:"level"`
            Strategy string  `json:"strategy"`
         
    }
    
    // BagsForReorder ...
    type BagsForReorder struct {

        
            StoreID float64  `json:"store_id"`
            ItemID float64  `json:"item_id"`
            Quantity float64  `json:"quantity"`
            SellerID float64  `json:"seller_id"`
            ItemSize string  `json:"item_size"`
            ArticleAssignment BagsForReorderArticleAssignment  `json:"article_assignment"`
         
    }
    
    // UserInfo ...
    type UserInfo struct {

        
            Email string  `json:"email"`
            Name string  `json:"name"`
            Mobile string  `json:"mobile"`
            Gender string  `json:"gender"`
         
    }
    
    // OrderSchema ...
    type OrderSchema struct {

        
            TotalShipmentsInOrder float64  `json:"total_shipments_in_order"`
            OrderCreatedTime string  `json:"order_created_time"`
            BreakupValues []BreakupValues  `json:"breakup_values"`
            Shipments []Shipments  `json:"shipments"`
            BagsForReorder []BagsForReorder  `json:"bags_for_reorder"`
            UserInfo UserInfo  `json:"user_info"`
            OrderID string  `json:"order_id"`
         
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

        
            ShipmentID string  `json:"shipment_id"`
            Success bool  `json:"success"`
            PresignedType string  `json:"presigned_type"`
            PresignedURL string  `json:"presigned_url"`
         
    }
    
    // Track ...
    type Track struct {

        
            LastLocationRecievedAt string  `json:"last_location_recieved_at"`
            UpdatedTime string  `json:"updated_time"`
            Status string  `json:"status"`
            Awb string  `json:"awb"`
            ShipmentType string  `json:"shipment_type"`
            Reason string  `json:"reason"`
            AccountName string  `json:"account_name"`
            UpdatedAt string  `json:"updated_at"`
         
    }
    
    // ShipmentTrack ...
    type ShipmentTrack struct {

        
            Results []Track  `json:"results"`
         
    }
    
    // CustomerDetailsResponse ...
    type CustomerDetailsResponse struct {

        
            Phone string  `json:"phone"`
            ShipmentID string  `json:"shipment_id"`
            Name string  `json:"name"`
            Country string  `json:"country"`
            OrderID string  `json:"order_id"`
         
    }
    
    // SendOtpToCustomerResponse ...
    type SendOtpToCustomerResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
            RequestID string  `json:"request_id"`
            ResendTimer float64  `json:"resend_timer"`
         
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
    
    // QuestionSet ...
    type QuestionSet struct {

        
            DisplayName string  `json:"display_name"`
            ID float64  `json:"id"`
         
    }
    
    // BagReasonMeta ...
    type BagReasonMeta struct {

        
            ShowTextArea bool  `json:"show_text_area"`
         
    }
    
    // BagReasons ...
    type BagReasons struct {

        
            ID float64  `json:"id"`
            DisplayName string  `json:"display_name"`
            Reasons []BagReasons  `json:"reasons"`
            QcType []string  `json:"qc_type"`
            QuestionSet []QuestionSet  `json:"question_set"`
            Meta BagReasonMeta  `json:"meta"`
         
    }
    
    // ShipmentBagReasons ...
    type ShipmentBagReasons struct {

        
            Reasons []BagReasons  `json:"reasons"`
            Success bool  `json:"success"`
         
    }
    
    // ShipmentReason ...
    type ShipmentReason struct {

        
            FeedbackType string  `json:"feedback_type"`
            ReasonID float64  `json:"reason_id"`
            Flow string  `json:"flow"`
            Priority float64  `json:"priority"`
            ShowTextArea bool  `json:"show_text_area"`
            ReasonText string  `json:"reason_text"`
         
    }
    
    // ShipmentReasons ...
    type ShipmentReasons struct {

        
            Reasons []ShipmentReason  `json:"reasons"`
         
    }
    
    // Products ...
    type Products struct {

        
            LineNumber float64  `json:"line_number"`
            Quantity float64  `json:"quantity"`
            Identifier string  `json:"identifier"`
         
    }
    
    // ProductsDataUpdatesFilters ...
    type ProductsDataUpdatesFilters struct {

        
            LineNumber float64  `json:"line_number"`
            Identifier string  `json:"identifier"`
         
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

        
            LineNumber float64  `json:"line_number"`
            Quantity float64  `json:"quantity"`
            Identifier string  `json:"identifier"`
         
    }
    
    // ProductsReasonsData ...
    type ProductsReasonsData struct {

        
            ReasonText string  `json:"reason_text"`
            ReasonID float64  `json:"reason_id"`
         
    }
    
    // ProductsReasons ...
    type ProductsReasons struct {

        
            Filters []ProductsReasonsFilters  `json:"filters"`
            Data ProductsReasonsData  `json:"data"`
         
    }
    
    // EntityReasonData ...
    type EntityReasonData struct {

        
            ReasonText string  `json:"reason_text"`
            ReasonID float64  `json:"reason_id"`
         
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
            Reasons ReasonsData  `json:"reasons"`
            Identifier string  `json:"identifier"`
         
    }
    
    // StatuesRequest ...
    type StatuesRequest struct {

        
            Shipments []ShipmentsRequest  `json:"shipments"`
            ExcludeBagsNextState string  `json:"exclude_bags_next_state"`
            Status string  `json:"status"`
         
    }
    
    // UpdateShipmentStatusRequest ...
    type UpdateShipmentStatusRequest struct {

        
            UnlockBeforeTransition bool  `json:"unlock_before_transition"`
            Statuses []StatuesRequest  `json:"statuses"`
            ForceTransition bool  `json:"force_transition"`
            LockAfterTransition bool  `json:"lock_after_transition"`
            Task bool  `json:"task"`
         
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

        
            Status float64  `json:"status"`
            Code string  `json:"code"`
            Exception string  `json:"exception"`
            StackTrace string  `json:"stack_trace"`
            Message string  `json:"message"`
         
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
    
    // CartProductIdentifer ...
    type CartProductIdentifer struct {

        
            Identifier string  `json:"identifier"`
         
    }
    
    // Ownership ...
    type Ownership struct {

        
            PayableCategory string  `json:"payable_category"`
            PayableBy string  `json:"payable_by"`
         
    }
    
    // BuyRules ...
    type BuyRules struct {

        
            ItemCriteria map[string]interface{}  `json:"item_criteria"`
            CartConditions map[string]interface{}  `json:"cart_conditions"`
         
    }
    
    // FreeGiftItem ...
    type FreeGiftItem struct {

        
            ItemImagesURL []string  `json:"item_images_url"`
            ItemPriceDetails map[string]interface{}  `json:"item_price_details"`
            ItemName string  `json:"item_name"`
            ItemID float64  `json:"item_id"`
            ItemBrandName string  `json:"item_brand_name"`
            ItemSlug string  `json:"item_slug"`
         
    }
    
    // AppliedFreeArticles ...
    type AppliedFreeArticles struct {

        
            ParentItemIdentifier string  `json:"parent_item_identifier"`
            FreeGiftItemDetails FreeGiftItem  `json:"free_gift_item_details"`
            ArticleID string  `json:"article_id"`
            Quantity float64  `json:"quantity"`
         
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

        
            Ownership Ownership  `json:"ownership"`
            Amount float64  `json:"amount"`
            PromotionType string  `json:"promotion_type"`
            ArticleQuantity float64  `json:"article_quantity"`
            PromoID string  `json:"promo_id"`
            PromotionName string  `json:"promotion_name"`
            PromotionGroup string  `json:"promotion_group"`
            BuyRules []BuyRules  `json:"buy_rules"`
            OfferText string  `json:"offer_text"`
            AppliedFreeArticles []AppliedFreeArticles  `json:"applied_free_articles"`
            MrpPromotion bool  `json:"mrp_promotion"`
            DiscountRules []DiscountRulesApp  `json:"discount_rules"`
         
    }
    
    // ProductPrice ...
    type ProductPrice struct {

        
            Effective float64  `json:"effective"`
            Selling float64  `json:"selling"`
            Marked float64  `json:"marked"`
            AddOn float64  `json:"add_on"`
            CurrencyCode string  `json:"currency_code"`
            CurrencySymbol string  `json:"currency_symbol"`
         
    }
    
    // ProductPriceInfo ...
    type ProductPriceInfo struct {

        
            Converted ProductPrice  `json:"converted"`
            Base ProductPrice  `json:"base"`
         
    }
    
    // PromoMeta ...
    type PromoMeta struct {

        
            Message string  `json:"message"`
         
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
    
    // ProductArticle ...
    type ProductArticle struct {

        
            Size string  `json:"size"`
            Quantity float64  `json:"quantity"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            Type string  `json:"type"`
            UID string  `json:"uid"`
            Price ArticlePriceInfo  `json:"price"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            Seller BaseInfo  `json:"seller"`
            ProductGroupTags []string  `json:"product_group_tags"`
            Store BaseInfo  `json:"store"`
         
    }
    
    // ProductAvailability ...
    type ProductAvailability struct {

        
            IsValid bool  `json:"is_valid"`
            OtherStoreQuantity float64  `json:"other_store_quantity"`
            Sizes []string  `json:"sizes"`
            OutOfStock bool  `json:"out_of_stock"`
            Deliverable bool  `json:"deliverable"`
         
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
            Query ActionQuery  `json:"query"`
            Type string  `json:"type"`
         
    }
    
    // NetQuantity ...
    type NetQuantity struct {

        
            Value string  `json:"value"`
            Unit string  `json:"unit"`
         
    }
    
    // CategoryInfo ...
    type CategoryInfo struct {

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // CartProduct ...
    type CartProduct struct {

        
            ItemCode string  `json:"item_code"`
            Brand BaseInfo  `json:"brand"`
            Slug string  `json:"slug"`
            Type string  `json:"type"`
            UID float64  `json:"uid"`
            Images []ProductImage  `json:"images"`
            Action ProductAction  `json:"action"`
            NetQuantity NetQuantity  `json:"net_quantity"`
            Name string  `json:"name"`
            Categories []CategoryInfo  `json:"categories"`
         
    }
    
    // CartProductInfo ...
    type CartProductInfo struct {

        
            Identifiers CartProductIdentifer  `json:"identifiers"`
            PromotionsApplied []AppliedPromotion  `json:"promotions_applied"`
            Quantity float64  `json:"quantity"`
            PricePerUnit ProductPriceInfo  `json:"price_per_unit"`
            PromoMeta PromoMeta  `json:"promo_meta"`
            Discount string  `json:"discount"`
            Article ProductArticle  `json:"article"`
            Message string  `json:"message"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            Availability ProductAvailability  `json:"availability"`
            BulkOffer map[string]interface{}  `json:"bulk_offer"`
            IsSet bool  `json:"is_set"`
            Product CartProduct  `json:"product"`
            Price ProductPriceInfo  `json:"price"`
            CouponMessage string  `json:"coupon_message"`
            Key string  `json:"key"`
         
    }
    
    // LoyaltyPoints ...
    type LoyaltyPoints struct {

        
            Applicable float64  `json:"applicable"`
            Description string  `json:"description"`
            IsApplied bool  `json:"is_applied"`
            Total float64  `json:"total"`
         
    }
    
    // RawBreakup ...
    type RawBreakup struct {

        
            DeliveryCharge float64  `json:"delivery_charge"`
            Vog float64  `json:"vog"`
            Total float64  `json:"total"`
            YouSaved float64  `json:"you_saved"`
            Coupon float64  `json:"coupon"`
            CodCharge float64  `json:"cod_charge"`
            ConvenienceFee float64  `json:"convenience_fee"`
            Discount float64  `json:"discount"`
            GstCharges float64  `json:"gst_charges"`
            MrpTotal float64  `json:"mrp_total"`
            Subtotal float64  `json:"subtotal"`
            FyndCash float64  `json:"fynd_cash"`
         
    }
    
    // DisplayBreakup ...
    type DisplayBreakup struct {

        
            Display string  `json:"display"`
            Message []string  `json:"message"`
            Value float64  `json:"value"`
            CurrencyCode string  `json:"currency_code"`
            CurrencySymbol string  `json:"currency_symbol"`
            Key string  `json:"key"`
         
    }
    
    // CouponBreakup ...
    type CouponBreakup struct {

        
            Description string  `json:"description"`
            MaxDiscountValue float64  `json:"max_discount_value"`
            SubTitle string  `json:"sub_title"`
            IsApplied bool  `json:"is_applied"`
            Message string  `json:"message"`
            Type string  `json:"type"`
            Title string  `json:"title"`
            UID string  `json:"uid"`
            Value float64  `json:"value"`
            MinimumCartValue float64  `json:"minimum_cart_value"`
            CouponType string  `json:"coupon_type"`
            CouponValue float64  `json:"coupon_value"`
            Code string  `json:"code"`
         
    }
    
    // CartBreakup ...
    type CartBreakup struct {

        
            LoyaltyPoints LoyaltyPoints  `json:"loyalty_points"`
            Raw RawBreakup  `json:"raw"`
            Display []DisplayBreakup  `json:"display"`
            Coupon CouponBreakup  `json:"coupon"`
         
    }
    
    // PaymentSelectionLock ...
    type PaymentSelectionLock struct {

        
            Enabled bool  `json:"enabled"`
            DefaultOptions string  `json:"default_options"`
            PaymentIdentifier string  `json:"payment_identifier"`
         
    }
    
    // CartCurrency ...
    type CartCurrency struct {

        
            Code string  `json:"code"`
            Symbol string  `json:"symbol"`
         
    }
    
    // CartDetailResponse ...
    type CartDetailResponse struct {

        
            RestrictCheckout bool  `json:"restrict_checkout"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            Gstin string  `json:"gstin"`
            PanNo string  `json:"pan_no"`
            Items []CartProductInfo  `json:"items"`
            CouponText string  `json:"coupon_text"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            ID string  `json:"id"`
            Message string  `json:"message"`
            PanConfig map[string]interface{}  `json:"pan_config"`
            Comment string  `json:"comment"`
            BuyNow bool  `json:"buy_now"`
            CheckoutMode string  `json:"checkout_mode"`
            IsValid bool  `json:"is_valid"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            LastModified string  `json:"last_modified"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            Currency CartCurrency  `json:"currency"`
         
    }
    
    // AddProductCart ...
    type AddProductCart struct {

        
            Display string  `json:"display"`
            ArticleAssignment map[string]interface{}  `json:"article_assignment"`
            Quantity float64  `json:"quantity"`
            StoreID float64  `json:"store_id"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            ArticleID string  `json:"article_id"`
            Pos bool  `json:"pos"`
            SellerID float64  `json:"seller_id"`
            ItemSize string  `json:"item_size"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            ItemID float64  `json:"item_id"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ProductGroupTags []string  `json:"product_group_tags"`
         
    }
    
    // AddCartRequest ...
    type AddCartRequest struct {

        
            NewCart bool  `json:"new_cart"`
            Items []AddProductCart  `json:"items"`
         
    }
    
    // AddCartDetailResponse ...
    type AddCartDetailResponse struct {

        
            Success bool  `json:"success"`
            Cart CartDetailResponse  `json:"cart"`
            Message string  `json:"message"`
            Partial bool  `json:"partial"`
         
    }
    
    // UpdateProductCart ...
    type UpdateProductCart struct {

        
            ItemIndex float64  `json:"item_index"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
            Quantity float64  `json:"quantity"`
            ArticleID string  `json:"article_id"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            ItemSize string  `json:"item_size"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            ItemID float64  `json:"item_id"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // UpdateCartRequest ...
    type UpdateCartRequest struct {

        
            Operation string  `json:"operation"`
            Items []UpdateProductCart  `json:"items"`
         
    }
    
    // UpdateCartDetailResponse ...
    type UpdateCartDetailResponse struct {

        
            Success bool  `json:"success"`
            Cart CartDetailResponse  `json:"cart"`
            Message string  `json:"message"`
         
    }
    
    // CartItemCountResponse ...
    type CartItemCountResponse struct {

        
            UserCartItemsCount float64  `json:"user_cart_items_count"`
         
    }
    
    // Coupon ...
    type Coupon struct {

        
            Description string  `json:"description"`
            MaxDiscountValue float64  `json:"max_discount_value"`
            SubTitle string  `json:"sub_title"`
            CouponType string  `json:"coupon_type"`
            Message string  `json:"message"`
            CouponCode string  `json:"coupon_code"`
            MinimumCartValue float64  `json:"minimum_cart_value"`
            Title string  `json:"title"`
            IsApplied bool  `json:"is_applied"`
            CouponValue float64  `json:"coupon_value"`
            ExpiresOn string  `json:"expires_on"`
            IsApplicable bool  `json:"is_applicable"`
         
    }
    
    // PageCoupon ...
    type PageCoupon struct {

        
            TotalItemCount float64  `json:"total_item_count"`
            HasPrevious bool  `json:"has_previous"`
            Current float64  `json:"current"`
            HasNext bool  `json:"has_next"`
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

        
            Effective float64  `json:"effective"`
            BulkEffective float64  `json:"bulk_effective"`
            Marked float64  `json:"marked"`
            CurrencyCode string  `json:"currency_code"`
            CurrencySymbol string  `json:"currency_symbol"`
         
    }
    
    // OfferItem ...
    type OfferItem struct {

        
            Margin float64  `json:"margin"`
            AutoApplied bool  `json:"auto_applied"`
            Quantity float64  `json:"quantity"`
            Type string  `json:"type"`
            Best bool  `json:"best"`
            Price OfferPrice  `json:"price"`
            Total float64  `json:"total"`
         
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

        
            CountryCode string  `json:"country_code"`
            CheckoutMode string  `json:"checkout_mode"`
            Tags []string  `json:"tags"`
            IsActive bool  `json:"is_active"`
            Landmark string  `json:"landmark"`
            AreaCodeSlug string  `json:"area_code_slug"`
            CountryPhoneCode string  `json:"country_phone_code"`
            Email string  `json:"email"`
            Phone string  `json:"phone"`
            Area string  `json:"area"`
            City string  `json:"city"`
            State string  `json:"state"`
            Country string  `json:"country"`
            CreatedByUserID string  `json:"created_by_user_id"`
            UserID string  `json:"user_id"`
            IsDefaultAddress bool  `json:"is_default_address"`
            AddressType string  `json:"address_type"`
            GoogleMapPoint map[string]interface{}  `json:"google_map_point"`
            AreaCode string  `json:"area_code"`
            Meta map[string]interface{}  `json:"meta"`
            CountryIsoCode string  `json:"country_iso_code"`
            GeoLocation GeoLocation  `json:"geo_location"`
            ID string  `json:"id"`
            Name string  `json:"name"`
            Address string  `json:"address"`
         
    }
    
    // GetAddressesResponse ...
    type GetAddressesResponse struct {

        
            Address []Address  `json:"address"`
         
    }
    
    // SaveAddressResponse ...
    type SaveAddressResponse struct {

        
            IsDefaultAddress bool  `json:"is_default_address"`
            Success bool  `json:"success"`
            ID string  `json:"id"`
         
    }
    
    // UpdateAddressResponse ...
    type UpdateAddressResponse struct {

        
            IsDefaultAddress bool  `json:"is_default_address"`
            Success bool  `json:"success"`
            ID string  `json:"id"`
            IsUpdated bool  `json:"is_updated"`
         
    }
    
    // DeleteAddressResponse ...
    type DeleteAddressResponse struct {

        
            IsDeleted bool  `json:"is_deleted"`
            ID string  `json:"id"`
         
    }
    
    // SelectCartAddressRequest ...
    type SelectCartAddressRequest struct {

        
            CartID string  `json:"cart_id"`
            ID string  `json:"id"`
            BillingAddressID string  `json:"billing_address_id"`
         
    }
    
    // UpdateCartPaymentRequest ...
    type UpdateCartPaymentRequest struct {

        
            AggregatorName string  `json:"aggregator_name"`
            AddressID string  `json:"address_id"`
            PaymentMode string  `json:"payment_mode"`
            ID string  `json:"id"`
            MerchantCode string  `json:"merchant_code"`
            PaymentIdentifier string  `json:"payment_identifier"`
         
    }
    
    // CouponValidity ...
    type CouponValidity struct {

        
            Valid bool  `json:"valid"`
            DisplayMessageEn string  `json:"display_message_en"`
            Discount float64  `json:"discount"`
            Title string  `json:"title"`
            Code string  `json:"code"`
         
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
            Items []CartProductInfo  `json:"items"`
            Promise ShipmentPromise  `json:"promise"`
            DpID string  `json:"dp_id"`
            ShipmentType string  `json:"shipment_type"`
            FulfillmentID float64  `json:"fulfillment_id"`
            BoxType string  `json:"box_type"`
            OrderType string  `json:"order_type"`
            DpOptions map[string]interface{}  `json:"dp_options"`
            Shipments float64  `json:"shipments"`
         
    }
    
    // CartShipmentsResponse ...
    type CartShipmentsResponse struct {

        
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            Message string  `json:"message"`
            Comment string  `json:"comment"`
            CheckoutMode string  `json:"checkout_mode"`
            IsValid bool  `json:"is_valid"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            Gstin string  `json:"gstin"`
            LastModified string  `json:"last_modified"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            CartID float64  `json:"cart_id"`
            Currency CartCurrency  `json:"currency"`
            Shipments []ShipmentResponse  `json:"shipments"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            Error bool  `json:"error"`
            CouponText string  `json:"coupon_text"`
            ID string  `json:"id"`
            BuyNow bool  `json:"buy_now"`
            UID string  `json:"uid"`
         
    }
    
    // UpdateCartShipmentItem ...
    type UpdateCartShipmentItem struct {

        
            ArticleUID string  `json:"article_uid"`
            ShipmentType string  `json:"shipment_type"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // UpdateCartShipmentRequest ...
    type UpdateCartShipmentRequest struct {

        
            Shipments []UpdateCartShipmentItem  `json:"shipments"`
         
    }
    
    // StaffCheckout ...
    type StaffCheckout struct {

        
            LastName string  `json:"last_name"`
            FirstName string  `json:"first_name"`
            User string  `json:"user"`
            EmployeeCode string  `json:"employee_code"`
            ID string  `json:"_id"`
         
    }
    
    // CartCheckoutCustomMeta ...
    type CartCheckoutCustomMeta struct {

        
            Value string  `json:"value"`
            Key string  `json:"key"`
         
    }
    
    // Files ...
    type Files struct {

        
            Values []string  `json:"values"`
            Key string  `json:"key"`
         
    }
    
    // CartPosCheckoutDetailRequest ...
    type CartPosCheckoutDetailRequest struct {

        
            Aggregator string  `json:"aggregator"`
            PaymentMode string  `json:"payment_mode"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
            PaymentParams map[string]interface{}  `json:"payment_params"`
            CallbackURL string  `json:"callback_url"`
            OrderingStore float64  `json:"ordering_store"`
            OrderType string  `json:"order_type"`
            Staff StaffCheckout  `json:"staff"`
            BillingAddressID string  `json:"billing_address_id"`
            Pos bool  `json:"pos"`
            PickAtStoreUID float64  `json:"pick_at_store_uid"`
            PaymentAutoConfirm bool  `json:"payment_auto_confirm"`
            MerchantCode string  `json:"merchant_code"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
            PaymentIdentifier string  `json:"payment_identifier"`
            Meta map[string]interface{}  `json:"meta"`
            AddressID string  `json:"address_id"`
            CustomMeta []CartCheckoutCustomMeta  `json:"custom_meta"`
            Files []Files  `json:"files"`
         
    }
    
    // CheckCart ...
    type CheckCart struct {

        
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            StoreEmps []map[string]interface{}  `json:"store_emps"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            Message string  `json:"message"`
            Comment string  `json:"comment"`
            CheckoutMode string  `json:"checkout_mode"`
            IsValid bool  `json:"is_valid"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            Gstin string  `json:"gstin"`
            LastModified string  `json:"last_modified"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            CartID float64  `json:"cart_id"`
            OrderID string  `json:"order_id"`
            Currency CartCurrency  `json:"currency"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            DeliveryChargeOrderValue float64  `json:"delivery_charge_order_value"`
            StoreCode string  `json:"store_code"`
            DeliveryCharges float64  `json:"delivery_charges"`
            BuyNow bool  `json:"buy_now"`
            UID string  `json:"uid"`
            CodAvailable bool  `json:"cod_available"`
            UserType string  `json:"user_type"`
            Success bool  `json:"success"`
            CodCharges float64  `json:"cod_charges"`
            ErrorMessage string  `json:"error_message"`
            CodMessage string  `json:"cod_message"`
            Items []CartProductInfo  `json:"items"`
            CouponText string  `json:"coupon_text"`
            ID string  `json:"id"`
         
    }
    
    // CartCheckoutResponse ...
    type CartCheckoutResponse struct {

        
            AppInterceptURL string  `json:"app_intercept_url"`
            Cart CheckCart  `json:"cart"`
            CallbackURL string  `json:"callback_url"`
            Message string  `json:"message"`
            PaymentConfirmURL string  `json:"payment_confirm_url"`
            Success bool  `json:"success"`
            Data map[string]interface{}  `json:"data"`
            OrderID string  `json:"order_id"`
         
    }
    
    // CartMetaRequest ...
    type CartMetaRequest struct {

        
            CheckoutMode string  `json:"checkout_mode"`
            Comment string  `json:"comment"`
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
    
    // CartDeliveryModesResponse ...
    type CartDeliveryModesResponse struct {

        
            AvailableModes []string  `json:"available_modes"`
            PickupStores []float64  `json:"pickup_stores"`
         
    }
    
    // PickupStoreDetail ...
    type PickupStoreDetail struct {

        
            AreaCode string  `json:"area_code"`
            StoreCode string  `json:"store_code"`
            State string  `json:"state"`
            Pincode float64  `json:"pincode"`
            ID float64  `json:"id"`
            UID float64  `json:"uid"`
            Phone string  `json:"phone"`
            Area string  `json:"area"`
            AddressType string  `json:"address_type"`
            City string  `json:"city"`
            Landmark string  `json:"landmark"`
            Name string  `json:"name"`
            Address string  `json:"address"`
            AreaCodeSlug string  `json:"area_code_slug"`
            Country string  `json:"country"`
            Email string  `json:"email"`
         
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

        
            Meta map[string]interface{}  `json:"meta"`
            Token string  `json:"token"`
            User map[string]interface{}  `json:"user"`
            CreatedOn string  `json:"created_on"`
            Source map[string]interface{}  `json:"source"`
         
    }
    
    // SharedCart ...
    type SharedCart struct {

        
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            Message string  `json:"message"`
            Comment string  `json:"comment"`
            CheckoutMode string  `json:"checkout_mode"`
            SharedCartDetails SharedCartDetails  `json:"shared_cart_details"`
            IsValid bool  `json:"is_valid"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            Gstin string  `json:"gstin"`
            LastModified string  `json:"last_modified"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            CartID float64  `json:"cart_id"`
            Currency CartCurrency  `json:"currency"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            BuyNow bool  `json:"buy_now"`
            UID string  `json:"uid"`
            Items []CartProductInfo  `json:"items"`
            CouponText string  `json:"coupon_text"`
            ID string  `json:"id"`
         
    }
    
    // SharedCartResponse ...
    type SharedCartResponse struct {

        
            Cart SharedCart  `json:"cart"`
            Error string  `json:"error"`
         
    }
    

    
    // PincodeErrorSchemaResponse ...
    type PincodeErrorSchemaResponse struct {

        
            Value string  `json:"value"`
            Type string  `json:"type"`
            Message string  `json:"message"`
         
    }
    
    // PincodeParentsResponse ...
    type PincodeParentsResponse struct {

        
            UID string  `json:"uid"`
            SubType string  `json:"sub_type"`
            Name string  `json:"name"`
            DisplayName string  `json:"display_name"`
         
    }
    
    // CountryMetaResponse ...
    type CountryMetaResponse struct {

        
            IsdCode string  `json:"isd_code"`
            CountryCode string  `json:"country_code"`
         
    }
    
    // PincodeMetaResponse ...
    type PincodeMetaResponse struct {

        
            InternalZoneID float64  `json:"internal_zone_id"`
            Zone string  `json:"zone"`
         
    }
    
    // PincodeDataResponse ...
    type PincodeDataResponse struct {

        
            Parents []PincodeParentsResponse  `json:"parents"`
            MetaCode CountryMetaResponse  `json:"meta_code"`
            Meta PincodeMetaResponse  `json:"meta"`
            Name string  `json:"name"`
            SubType string  `json:"sub_type"`
            UID string  `json:"uid"`
            Error PincodeErrorSchemaResponse  `json:"error"`
            DisplayName string  `json:"display_name"`
         
    }
    
    // PincodeApiResponse ...
    type PincodeApiResponse struct {

        
            Success bool  `json:"success"`
            Error PincodeErrorSchemaResponse  `json:"error"`
            Data []PincodeDataResponse  `json:"data"`
         
    }
    
    // TATCategoryRequest ...
    type TATCategoryRequest struct {

        
            ID float64  `json:"id"`
            Level string  `json:"level"`
         
    }
    
    // TATArticlesRequest ...
    type TATArticlesRequest struct {

        
            Category TATCategoryRequest  `json:"category"`
            ManufacturingTime float64  `json:"manufacturing_time"`
            ManufacturingTimeUnit string  `json:"manufacturing_time_unit"`
         
    }
    
    // TATLocationDetailsRequest ...
    type TATLocationDetailsRequest struct {

        
            Articles []TATArticlesRequest  `json:"articles"`
            FulfillmentID float64  `json:"fulfillment_id"`
            FromPincode string  `json:"from_pincode"`
         
    }
    
    // TATViewRequest ...
    type TATViewRequest struct {

        
            Source string  `json:"source"`
            Action string  `json:"action"`
            LocationDetails []TATLocationDetailsRequest  `json:"location_details"`
            Identifier string  `json:"identifier"`
            ToPincode string  `json:"to_pincode"`
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

        
            Value string  `json:"value"`
            Type string  `json:"type"`
            Message string  `json:"message"`
         
    }
    
    // TATArticlesResponse ...
    type TATArticlesResponse struct {

        
            Promise TATPromiseResponse  `json:"promise"`
            ManufacturingTime float64  `json:"manufacturing_time"`
            ManufacturingTimeUnit string  `json:"manufacturing_time_unit"`
            ManufacturingTimeSeconds float64  `json:"_manufacturing_time_seconds"`
            Category TATCategoryRequest  `json:"category"`
            Error TATErrorSchemaResponse  `json:"error"`
            IsCodAvailable bool  `json:"is_cod_available"`
         
    }
    
    // TATLocationDetailsResponse ...
    type TATLocationDetailsResponse struct {

        
            Articles []TATArticlesResponse  `json:"articles"`
            FulfillmentID float64  `json:"fulfillment_id"`
            FromPincode string  `json:"from_pincode"`
         
    }
    
    // TATViewResponse ...
    type TATViewResponse struct {

        
            Source string  `json:"source"`
            RequestUUID string  `json:"request_uuid"`
            StormbreakerUUID string  `json:"stormbreaker_uuid"`
            ToCity string  `json:"to_city"`
            Action string  `json:"action"`
            LocationDetails []TATLocationDetailsResponse  `json:"location_details"`
            Identifier string  `json:"identifier"`
            Success bool  `json:"success"`
            PaymentMode string  `json:"payment_mode"`
            ToPincode string  `json:"to_pincode"`
            Journey string  `json:"journey"`
            Error TATErrorSchemaResponse  `json:"error"`
            IsCodAvailable bool  `json:"is_cod_available"`
         
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

        
            IsActive bool  `json:"is_active"`
            Logistics LogisticsResponse  `json:"logistics"`
            Meta CountryMetaResponse  `json:"meta"`
            Name string  `json:"name"`
            SubType string  `json:"sub_type"`
            ParentID string  `json:"parent_id"`
            UID string  `json:"uid"`
            Type string  `json:"type"`
            DisplayName string  `json:"display_name"`
         
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

        
            ServiceabilityType string  `json:"serviceability_type"`
            Zones []string  `json:"zones"`
         
    }
    
