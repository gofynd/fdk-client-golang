package application



    
    // ProductListingActionPage ...
    type ProductListingActionPage struct {

        
            Params map[string]interface{}  `json:"params"`
            Type string  `json:"type"`
            Query map[string]interface{}  `json:"query"`
         
    }
    
    // ProductListingAction ...
    type ProductListingAction struct {

        
            Type string  `json:"type"`
            Page ProductListingActionPage  `json:"page"`
         
    }
    
    // Meta ...
    type Meta struct {

        
            Source string  `json:"source"`
         
    }
    
    // Media ...
    type Media struct {

        
            Alt string  `json:"alt"`
            URL string  `json:"url"`
            Type string  `json:"type"`
            Meta Meta  `json:"meta"`
         
    }
    
    // ProductBrand ...
    type ProductBrand struct {

        
            Action ProductListingAction  `json:"action"`
            Description string  `json:"description"`
            Logo Media  `json:"logo"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // ProductCategoryMap ...
    type ProductCategoryMap struct {

        
            L2 ProductBrand  `json:"l2"`
            L1 ProductBrand  `json:"l1"`
            L3 ProductBrand  `json:"l3"`
         
    }
    
    // Price ...
    type Price struct {

        
            Max float64  `json:"max"`
            Min float64  `json:"min"`
            CurrencySymbol string  `json:"currency_symbol"`
            CurrencyCode string  `json:"currency_code"`
         
    }
    
    // ProductListingPrice ...
    type ProductListingPrice struct {

        
            Effective Price  `json:"effective"`
            Marked Price  `json:"marked"`
         
    }
    
    // CustomMetaFields ...
    type CustomMetaFields struct {

        
            Key string  `json:"key"`
            Value string  `json:"value"`
         
    }
    
    // NetQuantity ...
    type NetQuantity struct {

        
            Unit interface{}  `json:"unit"`
            Value float64  `json:"value"`
         
    }
    
    // ApplicationItemSEO ...
    type ApplicationItemSEO struct {

        
            Title interface{}  `json:"title"`
            Description interface{}  `json:"description"`
         
    }
    
    // ProductDetailAttribute ...
    type ProductDetailAttribute struct {

        
            Key string  `json:"key"`
            Type string  `json:"type"`
            Value string  `json:"value"`
         
    }
    
    // ProductDetailGroupedAttribute ...
    type ProductDetailGroupedAttribute struct {

        
            Title string  `json:"title"`
            Details []ProductDetailAttribute  `json:"details"`
         
    }
    
    // ApplicationItemMOQ ...
    type ApplicationItemMOQ struct {

        
            Minimum float64  `json:"minimum"`
            IncrementUnit float64  `json:"increment_unit"`
            Maximum float64  `json:"maximum"`
         
    }
    
    // ProductDetail ...
    type ProductDetail struct {

        
            Brand ProductBrand  `json:"brand"`
            Tryouts []string  `json:"tryouts"`
            ProductOnlineDate string  `json:"product_online_date"`
            ItemType string  `json:"item_type"`
            Rating float64  `json:"rating"`
            Categories []ProductBrand  `json:"categories"`
            Medias []Media  `json:"medias"`
            Tags []string  `json:"tags"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ItemCode string  `json:"item_code"`
            UID float64  `json:"uid"`
            CategoryMap ProductCategoryMap  `json:"category_map"`
            Price ProductListingPrice  `json:"price"`
            Attributes map[string]interface{}  `json:"attributes"`
            Type string  `json:"type"`
            Color string  `json:"color"`
            Similars []string  `json:"similars"`
            Action ProductListingAction  `json:"action"`
            CustomMeta []CustomMetaFields  `json:"_custom_meta"`
            NetQuantity NetQuantity  `json:"net_quantity"`
            ImageNature string  `json:"image_nature"`
            Discount string  `json:"discount"`
            RatingCount float64  `json:"rating_count"`
            Name string  `json:"name"`
            TeaserTag string  `json:"teaser_tag"`
            Seo ApplicationItemSEO  `json:"seo"`
            IsDependent bool  `json:"is_dependent"`
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
            Highlights []string  `json:"highlights"`
            ShortDescription string  `json:"short_description"`
            Description string  `json:"description"`
            Moq ApplicationItemMOQ  `json:"moq"`
            Slug string  `json:"slug"`
            HasVariant bool  `json:"has_variant"`
         
    }
    
    // ErrorResponse ...
    type ErrorResponse struct {

        
            Error string  `json:"error"`
         
    }
    
    // Weight ...
    type Weight struct {

        
            Unit string  `json:"unit"`
            IsDefault bool  `json:"is_default"`
            Shipping float64  `json:"shipping"`
         
    }
    
    // Dimension ...
    type Dimension struct {

        
            Unit string  `json:"unit"`
            IsDefault bool  `json:"is_default"`
            Height float64  `json:"height"`
            Length float64  `json:"length"`
            Width float64  `json:"width"`
         
    }
    
    // ProductSize ...
    type ProductSize struct {

        
            Quantity float64  `json:"quantity"`
            Value string  `json:"value"`
            Weight Weight  `json:"weight"`
            Display string  `json:"display"`
            Dimension Dimension  `json:"dimension"`
            IsAvailable bool  `json:"is_available"`
            SellerIdentifiers []string  `json:"seller_identifiers"`
         
    }
    
    // ProductSizeStores ...
    type ProductSizeStores struct {

        
            Count float64  `json:"count"`
         
    }
    
    // SizeChartValues ...
    type SizeChartValues struct {

        
            Col5 string  `json:"col_5"`
            Col6 string  `json:"col_6"`
            Col2 string  `json:"col_2"`
            Col4 string  `json:"col_4"`
            Col1 string  `json:"col_1"`
            Col3 string  `json:"col_3"`
         
    }
    
    // ColumnHeader ...
    type ColumnHeader struct {

        
            Convertable bool  `json:"convertable"`
            Value string  `json:"value"`
         
    }
    
    // ColumnHeaders ...
    type ColumnHeaders struct {

        
            Col5 ColumnHeader  `json:"col_5"`
            Col6 ColumnHeader  `json:"col_6"`
            Col2 ColumnHeader  `json:"col_2"`
            Col4 ColumnHeader  `json:"col_4"`
            Col1 ColumnHeader  `json:"col_1"`
            Col3 ColumnHeader  `json:"col_3"`
         
    }
    
    // SizeChart ...
    type SizeChart struct {

        
            Sizes []SizeChartValues  `json:"sizes"`
            SizeTip string  `json:"size_tip"`
            Title string  `json:"title"`
            Unit string  `json:"unit"`
            Description string  `json:"description"`
            Image string  `json:"image"`
            Headers ColumnHeaders  `json:"headers"`
         
    }
    
    // ProductSizes ...
    type ProductSizes struct {

        
            Sizes []ProductSize  `json:"sizes"`
            MultiSize bool  `json:"multi_size"`
            Sellable bool  `json:"sellable"`
            Stores ProductSizeStores  `json:"stores"`
            SizeChart SizeChart  `json:"size_chart"`
            Discount string  `json:"discount"`
            Price ProductListingPrice  `json:"price"`
         
    }
    
    // AttributeDetail ...
    type AttributeDetail struct {

        
            Key string  `json:"key"`
            Logo string  `json:"logo"`
            Display string  `json:"display"`
            Description string  `json:"description"`
         
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
            AttributesMetadata []AttributeMetadata  `json:"attributes_metadata"`
            Subtitle string  `json:"subtitle"`
            Items []ProductDetail  `json:"items"`
         
    }
    
    // ProductFrequentlyComparedSimilarResponse ...
    type ProductFrequentlyComparedSimilarResponse struct {

        
            Similars ProductCompareResponse  `json:"similars"`
         
    }
    
    // ProductVariantItemResponse ...
    type ProductVariantItemResponse struct {

        
            Color string  `json:"color"`
            ColorName string  `json:"color_name"`
            Action ProductListingAction  `json:"action"`
            Value string  `json:"value"`
            CustomMeta []CustomMetaFields  `json:"_custom_meta"`
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
            Medias []Media  `json:"medias"`
            Name string  `json:"name"`
            IsAvailable bool  `json:"is_available"`
         
    }
    
    // ProductVariantResponse ...
    type ProductVariantResponse struct {

        
            Header string  `json:"header"`
            Key string  `json:"key"`
            DisplayType string  `json:"display_type"`
            Items []ProductVariantItemResponse  `json:"items"`
         
    }
    
    // ProductVariantsResponse ...
    type ProductVariantsResponse struct {

        
            Variants []ProductVariantResponse  `json:"variants"`
         
    }
    
    // CompanyDetail ...
    type CompanyDetail struct {

        
            Name string  `json:"name"`
            ID float64  `json:"id"`
         
    }
    
    // Seller ...
    type Seller struct {

        
            Name string  `json:"name"`
            UID float64  `json:"uid"`
            Count float64  `json:"count"`
         
    }
    
    // ProductStockPrice ...
    type ProductStockPrice struct {

        
            Currency string  `json:"currency"`
            Effective float64  `json:"effective"`
            Marked float64  `json:"marked"`
         
    }
    
    // StoreDetail ...
    type StoreDetail struct {

        
            City string  `json:"city"`
            Name string  `json:"name"`
            Code string  `json:"code"`
            ID float64  `json:"id"`
         
    }
    
    // ProductStockStatusItem ...
    type ProductStockStatusItem struct {

        
            Quantity float64  `json:"quantity"`
            Company CompanyDetail  `json:"company"`
            Seller Seller  `json:"seller"`
            Size string  `json:"size"`
            UID string  `json:"uid"`
            Price ProductStockPrice  `json:"price"`
            Identifier map[string]interface{}  `json:"identifier"`
            ItemID float64  `json:"item_id"`
            Store StoreDetail  `json:"store"`
         
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

        
            IsSelected bool  `json:"is_selected"`
            Name string  `json:"name"`
            Value string  `json:"value"`
         
    }
    
    // ProductFiltersValue ...
    type ProductFiltersValue struct {

        
            CurrencyCode string  `json:"currency_code"`
            QueryFormat string  `json:"query_format"`
            DisplayFormat string  `json:"display_format"`
            Max float64  `json:"max"`
            Min float64  `json:"min"`
            SelectedMin float64  `json:"selected_min"`
            SelectedMax float64  `json:"selected_max"`
            Value string  `json:"value"`
            IsSelected bool  `json:"is_selected"`
            Display string  `json:"display"`
            Count float64  `json:"count"`
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

        
            Key string  `json:"key"`
            Header string  `json:"header"`
            Items []ProductVariantItemResponse  `json:"items"`
            Total float64  `json:"total"`
            DisplayType string  `json:"display_type"`
         
    }
    
    // ProductListingDetail ...
    type ProductListingDetail struct {

        
            Brand ProductBrand  `json:"brand"`
            Tryouts []string  `json:"tryouts"`
            ProductOnlineDate string  `json:"product_online_date"`
            ItemType string  `json:"item_type"`
            Rating float64  `json:"rating"`
            Categories []ProductBrand  `json:"categories"`
            Medias []Media  `json:"medias"`
            Tags []string  `json:"tags"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Sellable bool  `json:"sellable"`
            ItemCode string  `json:"item_code"`
            UID float64  `json:"uid"`
            CategoryMap ProductCategoryMap  `json:"category_map"`
            Price ProductListingPrice  `json:"price"`
            Attributes map[string]interface{}  `json:"attributes"`
            Sizes []string  `json:"sizes"`
            Type string  `json:"type"`
            Color string  `json:"color"`
            Similars []string  `json:"similars"`
            Action ProductListingAction  `json:"action"`
            CustomMeta []CustomMetaFields  `json:"_custom_meta"`
            NetQuantity NetQuantity  `json:"net_quantity"`
            ImageNature string  `json:"image_nature"`
            Discount string  `json:"discount"`
            RatingCount float64  `json:"rating_count"`
            Name string  `json:"name"`
            TeaserTag string  `json:"teaser_tag"`
            Seo ApplicationItemSEO  `json:"seo"`
            IsDependent bool  `json:"is_dependent"`
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
            Identifiers []string  `json:"identifiers"`
            Highlights []string  `json:"highlights"`
            ShortDescription string  `json:"short_description"`
            Description string  `json:"description"`
            Moq ApplicationItemMOQ  `json:"moq"`
            Slug string  `json:"slug"`
            HasVariant bool  `json:"has_variant"`
            Variants []ProductVariantListingResponse  `json:"variants"`
         
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

        
            Departments []string  `json:"departments"`
            Action ProductListingAction  `json:"action"`
            Description string  `json:"description"`
            Slug string  `json:"slug"`
            Logo Media  `json:"logo"`
            UID float64  `json:"uid"`
            Discount string  `json:"discount"`
            Banners ImageUrls  `json:"banners"`
            Name string  `json:"name"`
         
    }
    
    // BrandListingResponse ...
    type BrandListingResponse struct {

        
            Page Page  `json:"page"`
            Items []BrandItem  `json:"items"`
         
    }
    
    // BrandDetailResponse ...
    type BrandDetailResponse struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Description string  `json:"description"`
            Logo Media  `json:"logo"`
            UID float64  `json:"uid"`
            Banners ImageUrls  `json:"banners"`
            Name string  `json:"name"`
         
    }
    
    // ThirdLevelChild ...
    type ThirdLevelChild struct {

        
            Childs []map[string]interface{}  `json:"childs"`
            Action ProductListingAction  `json:"action"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
            Banners ImageUrls  `json:"banners"`
            Name string  `json:"name"`
         
    }
    
    // SecondLevelChild ...
    type SecondLevelChild struct {

        
            Childs []ThirdLevelChild  `json:"childs"`
            Action ProductListingAction  `json:"action"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
            Banners ImageUrls  `json:"banners"`
            Name string  `json:"name"`
         
    }
    
    // Child ...
    type Child struct {

        
            Childs []SecondLevelChild  `json:"childs"`
            Action ProductListingAction  `json:"action"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
            Banners ImageUrls  `json:"banners"`
            Name string  `json:"name"`
         
    }
    
    // CategoryBanner ...
    type CategoryBanner struct {

        
            Portrait Media  `json:"portrait"`
            Landscape Media  `json:"landscape"`
         
    }
    
    // CategoryItems ...
    type CategoryItems struct {

        
            Childs []Child  `json:"childs"`
            Action ProductListingAction  `json:"action"`
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
            Banners CategoryBanner  `json:"banners"`
            Name string  `json:"name"`
         
    }
    
    // DepartmentCategoryTree ...
    type DepartmentCategoryTree struct {

        
            Department string  `json:"department"`
            Items []CategoryItems  `json:"items"`
         
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
            UID float64  `json:"uid"`
            Banners ImageUrls  `json:"banners"`
            Name string  `json:"name"`
         
    }
    
    // HomeListingResponse ...
    type HomeListingResponse struct {

        
            Message string  `json:"message"`
            Page Page  `json:"page"`
            Items []ProductListingDetail  `json:"items"`
         
    }
    
    // Department ...
    type Department struct {

        
            PriorityOrder float64  `json:"priority_order"`
            Slug string  `json:"slug"`
            Logo Media  `json:"logo"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // DepartmentResponse ...
    type DepartmentResponse struct {

        
            Items []Department  `json:"items"`
         
    }
    
    // AutocompleteItem ...
    type AutocompleteItem struct {

        
            Type string  `json:"type"`
            Action ProductListingAction  `json:"action"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Logo Media  `json:"logo"`
            Display string  `json:"display"`
         
    }
    
    // AutoCompleteResponse ...
    type AutoCompleteResponse struct {

        
            Items []AutocompleteItem  `json:"items"`
         
    }
    
    // CollectionListingFilterType ...
    type CollectionListingFilterType struct {

        
            IsSelected bool  `json:"is_selected"`
            Name string  `json:"name"`
            Display string  `json:"display"`
         
    }
    
    // CollectionListingFilterTag ...
    type CollectionListingFilterTag struct {

        
            IsSelected bool  `json:"is_selected"`
            Name string  `json:"name"`
            Display string  `json:"display"`
         
    }
    
    // CollectionListingFilter ...
    type CollectionListingFilter struct {

        
            Type []CollectionListingFilterType  `json:"type"`
            Tags []CollectionListingFilterTag  `json:"tags"`
         
    }
    
    // CollectionQuery ...
    type CollectionQuery struct {

        
            Attribute string  `json:"attribute"`
            Op string  `json:"op"`
            Value []interface{}  `json:"value"`
         
    }
    
    // GetCollectionDetailNest ...
    type GetCollectionDetailNest struct {

        
            Priority float64  `json:"priority"`
            AllowSort bool  `json:"allow_sort"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Tag []string  `json:"tag"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Logo Media  `json:"logo"`
            SortOn string  `json:"sort_on"`
            Cron map[string]interface{}  `json:"cron"`
            Schedule map[string]interface{}  `json:"_schedule"`
            AppID string  `json:"app_id"`
            Meta map[string]interface{}  `json:"meta"`
            UID string  `json:"uid"`
            Badge map[string]interface{}  `json:"badge"`
            Type string  `json:"type"`
            AllowFacets bool  `json:"allow_facets"`
            Action ProductListingAction  `json:"action"`
            Query []CollectionQuery  `json:"query"`
            Name string  `json:"name"`
            IsActive bool  `json:"is_active"`
            Description string  `json:"description"`
            Slug string  `json:"slug"`
            Banners ImageUrls  `json:"banners"`
         
    }
    
    // GetCollectionListingResponse ...
    type GetCollectionListingResponse struct {

        
            Filters CollectionListingFilter  `json:"filters"`
            Page Page  `json:"page"`
            Items []GetCollectionDetailNest  `json:"items"`
         
    }
    
    // CollectionDetailResponse ...
    type CollectionDetailResponse struct {

        
            Priority float64  `json:"priority"`
            AllowSort bool  `json:"allow_sort"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Tag []string  `json:"tag"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Logo Media  `json:"logo"`
            SortOn string  `json:"sort_on"`
            Cron map[string]interface{}  `json:"cron"`
            Schedule map[string]interface{}  `json:"_schedule"`
            AppID string  `json:"app_id"`
            Meta map[string]interface{}  `json:"meta"`
            Badge map[string]interface{}  `json:"badge"`
            Type string  `json:"type"`
            AllowFacets bool  `json:"allow_facets"`
            Query []CollectionQuery  `json:"query"`
            Name string  `json:"name"`
            IsActive bool  `json:"is_active"`
            Description string  `json:"description"`
            Slug string  `json:"slug"`
            Banners ImageUrls  `json:"banners"`
         
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

        
            Brands []float64  `json:"brands"`
            Products []float64  `json:"products"`
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

        
            Pincode float64  `json:"pincode"`
            City string  `json:"city"`
            Address string  `json:"address"`
            LatLong LatLong  `json:"lat_long"`
            StoreEmail string  `json:"store_email"`
            State string  `json:"state"`
            StoreCode string  `json:"store_code"`
            UID float64  `json:"uid"`
            Country string  `json:"country"`
            Name string  `json:"name"`
         
    }
    
    // StoreListingResponse ...
    type StoreListingResponse struct {

        
            Page Page  `json:"page"`
            Items []Store  `json:"items"`
         
    }
    
    // StoreDepartments ...
    type StoreDepartments struct {

        
            PriorityOrder float64  `json:"priority_order"`
            Slug string  `json:"slug"`
            Logo string  `json:"logo"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // CompanyStore ...
    type CompanyStore struct {

        
            CompanyType string  `json:"company_type"`
            Name string  `json:"name"`
            BusinessType string  `json:"business_type"`
            UID float64  `json:"uid"`
         
    }
    
    // StoreAddressSerializer ...
    type StoreAddressSerializer struct {

        
            Address1 string  `json:"address1"`
            Pincode float64  `json:"pincode"`
            City string  `json:"city"`
            Latitude float64  `json:"latitude"`
            Landmark string  `json:"landmark"`
            Longitude float64  `json:"longitude"`
            Address2 string  `json:"address2"`
            State string  `json:"state"`
            Country string  `json:"country"`
         
    }
    
    // SellerPhoneNumber ...
    type SellerPhoneNumber struct {

        
            CountryCode float64  `json:"country_code"`
            Number string  `json:"number"`
         
    }
    
    // StoreManagerSerializer ...
    type StoreManagerSerializer struct {

        
            Email string  `json:"email"`
            MobileNo SellerPhoneNumber  `json:"mobile_no"`
            Name string  `json:"name"`
         
    }
    
    // AppStore ...
    type AppStore struct {

        
            Company CompanyStore  `json:"company"`
            Departments []StoreDepartments  `json:"departments"`
            Address StoreAddressSerializer  `json:"address"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            Manager StoreManagerSerializer  `json:"manager"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // ApplicationStoreListing ...
    type ApplicationStoreListing struct {

        
            Filters []StoreDepartments  `json:"filters"`
            Page Page  `json:"page"`
            Items []AppStore  `json:"items"`
         
    }
    
    // Time ...
    type Time struct {

        
            Hour float64  `json:"hour"`
            Minute float64  `json:"minute"`
         
    }
    
    // StoreTiming ...
    type StoreTiming struct {

        
            Closing Time  `json:"closing"`
            Open bool  `json:"open"`
            Opening Time  `json:"opening"`
            Weekday string  `json:"weekday"`
         
    }
    
    // StoreDetails ...
    type StoreDetails struct {

        
            Company CompanyStore  `json:"company"`
            Timing []StoreTiming  `json:"timing"`
            Departments []StoreDepartments  `json:"departments"`
            Address StoreAddressSerializer  `json:"address"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Manager StoreManagerSerializer  `json:"manager"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // UserDetail ...
    type UserDetail struct {

        
            Contact string  `json:"contact"`
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
            SuperUser bool  `json:"super_user"`
         
    }
    
    // Size ...
    type Size struct {

        
            Quantity float64  `json:"quantity"`
            IsAvailable bool  `json:"is_available"`
            Display interface{}  `json:"display"`
            Value interface{}  `json:"value"`
         
    }
    
    // ProductDetails ...
    type ProductDetails struct {

        
            IsSet bool  `json:"is_set"`
            Rating float64  `json:"rating"`
            CountryOfOrigin interface{}  `json:"country_of_origin"`
            TemplateTag interface{}  `json:"template_tag"`
            OutOfStock bool  `json:"out_of_stock"`
            BrandUID float64  `json:"brand_uid"`
            ItemCode interface{}  `json:"item_code"`
            Attributes map[string]interface{}  `json:"attributes"`
            HsnCode float64  `json:"hsn_code"`
            ImageNature interface{}  `json:"image_nature"`
            RatingCount float64  `json:"rating_count"`
            Name interface{}  `json:"name"`
            GroupedAttributes map[string]interface{}  `json:"grouped_attributes"`
            ShortDescription interface{}  `json:"short_description"`
            Highlights []interface{}  `json:"highlights"`
            Description interface{}  `json:"description"`
            Media []map[string]interface{}  `json:"media"`
            Slug interface{}  `json:"slug"`
            HasVariant bool  `json:"has_variant"`
            Identifier map[string]interface{}  `json:"identifier"`
            Images []interface{}  `json:"images"`
         
    }
    
    // ProductGroupPrice ...
    type ProductGroupPrice struct {

        
            MinEffective float64  `json:"min_effective"`
            MinMarked float64  `json:"min_marked"`
            Currency interface{}  `json:"currency"`
            MaxMarked float64  `json:"max_marked"`
            MaxEffective float64  `json:"max_effective"`
         
    }
    
    // ProductInGroup ...
    type ProductInGroup struct {

        
            Sizes []Size  `json:"sizes"`
            ProductDetails ProductDetails  `json:"product_details"`
            AutoAddToCart bool  `json:"auto_add_to_cart"`
            MaxQuantity float64  `json:"max_quantity"`
            ProductUID float64  `json:"product_uid"`
            AllowRemove bool  `json:"allow_remove"`
            Price ProductGroupPrice  `json:"price"`
            AutoSelect bool  `json:"auto_select"`
            MinQuantity float64  `json:"min_quantity"`
         
    }
    
    // ProductGroupingModel ...
    type ProductGroupingModel struct {

        
            CreatedBy UserDetail  `json:"created_by"`
            ID interface{}  `json:"_id"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            Products []ProductInGroup  `json:"products"`
            IsActive bool  `json:"is_active"`
            Choice interface{}  `json:"choice"`
            VerifiedBy UserDetail  `json:"verified_by"`
            PageVisibility []interface{}  `json:"page_visibility"`
            VerifiedOn string  `json:"verified_on"`
            ModifiedOn string  `json:"modified_on"`
            ModifiedBy UserDetail  `json:"modified_by"`
            Slug interface{}  `json:"slug"`
            Logo interface{}  `json:"logo"`
            Meta map[string]interface{}  `json:"meta"`
            CreatedOn string  `json:"created_on"`
            CompanyID float64  `json:"company_id"`
            Name interface{}  `json:"name"`
         
    }
    
    // ProductBundle ...
    type ProductBundle struct {

        
            Items []ProductGroupingModel  `json:"items"`
         
    }
    
    // ProductSetDistributionSizeV2 ...
    type ProductSetDistributionSizeV2 struct {

        
            Pieces float64  `json:"pieces"`
            Size string  `json:"size"`
         
    }
    
    // ProductSetDistributionV2 ...
    type ProductSetDistributionV2 struct {

        
            Sizes []ProductSetDistributionSizeV2  `json:"sizes"`
         
    }
    
    // ProductSetV2 ...
    type ProductSetV2 struct {

        
            Quantity float64  `json:"quantity"`
            SizeDistribution ProductSetDistributionV2  `json:"size_distribution"`
         
    }
    
    // ProductStockUnitPriceV2 ...
    type ProductStockUnitPriceV2 struct {

        
            Price float64  `json:"price"`
            Unit string  `json:"unit"`
            CurrencyCode string  `json:"currency_code"`
            CurrencySymbol string  `json:"currency_symbol"`
         
    }
    
    // StoreV2 ...
    type StoreV2 struct {

        
            Name string  `json:"name"`
            UID float64  `json:"uid"`
            Count float64  `json:"count"`
         
    }
    
    // DetailsSchemaV2 ...
    type DetailsSchemaV2 struct {

        
            Key string  `json:"key"`
            Type string  `json:"type"`
            Value string  `json:"value"`
         
    }
    
    // MarketPlaceSttributesSchemaV2 ...
    type MarketPlaceSttributesSchemaV2 struct {

        
            Title string  `json:"title"`
            Details []DetailsSchemaV2  `json:"details"`
         
    }
    
    // ReturnConfigSchemaV2 ...
    type ReturnConfigSchemaV2 struct {

        
            Time float64  `json:"time"`
            Unit string  `json:"unit"`
            Returnable bool  `json:"returnable"`
         
    }
    
    // ProductStockPriceV2 ...
    type ProductStockPriceV2 struct {

        
            Currency string  `json:"currency"`
            Effective float64  `json:"effective"`
            Marked float64  `json:"marked"`
         
    }
    
    // StrategyWiseListingSchemaV2 ...
    type StrategyWiseListingSchemaV2 struct {

        
            Quantity float64  `json:"quantity"`
            Pincode float64  `json:"pincode"`
            Distance float64  `json:"distance"`
            Tat float64  `json:"tat"`
         
    }
    
    // SellerGroupAttributes ...
    type SellerGroupAttributes struct {

        
            Title string  `json:"title"`
            Details []DetailsSchemaV2  `json:"details"`
         
    }
    
    // ArticleAssignmentV2 ...
    type ArticleAssignmentV2 struct {

        
            Level string  `json:"level"`
            Strategy string  `json:"strategy"`
         
    }
    
    // SellerV2 ...
    type SellerV2 struct {

        
            Name string  `json:"name"`
            UID float64  `json:"uid"`
            Count float64  `json:"count"`
         
    }
    
    // ProductSizePriceResponseV3 ...
    type ProductSizePriceResponseV3 struct {

        
            IsCod bool  `json:"is_cod"`
            ItemType string  `json:"item_type"`
            ArticleID string  `json:"article_id"`
            Set ProductSetV2  `json:"set"`
            PricePerUnit ProductStockUnitPriceV2  `json:"price_per_unit"`
            Store StoreV2  `json:"store"`
            MarketplaceAttributes []MarketPlaceSttributesSchemaV2  `json:"marketplace_attributes"`
            SpecialBadge string  `json:"special_badge"`
            ReturnConfig ReturnConfigSchemaV2  `json:"return_config"`
            LongLat []float64  `json:"long_lat"`
            Price ProductStockPriceV2  `json:"price"`
            SellerCount float64  `json:"seller_count"`
            Quantity float64  `json:"quantity"`
            StrategyWiseListing []StrategyWiseListingSchemaV2  `json:"strategy_wise_listing"`
            PricePerPiece ProductStockPriceV2  `json:"price_per_piece"`
            Pincode float64  `json:"pincode"`
            IsServiceable bool  `json:"is_serviceable"`
            Discount string  `json:"discount"`
            GroupedAttributes []SellerGroupAttributes  `json:"grouped_attributes"`
            IsGift bool  `json:"is_gift"`
            ArticleAssignment ArticleAssignmentV2  `json:"article_assignment"`
            Seller SellerV2  `json:"seller"`
         
    }
    
    // ProductSizeSellerFilterSchemaV2 ...
    type ProductSizeSellerFilterSchemaV2 struct {

        
            IsSelected bool  `json:"is_selected"`
            Name string  `json:"name"`
            Value string  `json:"value"`
         
    }
    
    // ProductSizePriceResponseV2 ...
    type ProductSizePriceResponseV2 struct {

        
            IsCod bool  `json:"is_cod"`
            ItemType string  `json:"item_type"`
            ArticleID string  `json:"article_id"`
            Set ProductSetV2  `json:"set"`
            PricePerUnit ProductStockUnitPriceV2  `json:"price_per_unit"`
            Store StoreV2  `json:"store"`
            MarketplaceAttributes []MarketPlaceSttributesSchemaV2  `json:"marketplace_attributes"`
            SpecialBadge string  `json:"special_badge"`
            ReturnConfig ReturnConfigSchemaV2  `json:"return_config"`
            LongLat []float64  `json:"long_lat"`
            Price ProductStockPriceV2  `json:"price"`
            SellerCount float64  `json:"seller_count"`
            Quantity float64  `json:"quantity"`
            StrategyWiseListing []StrategyWiseListingSchemaV2  `json:"strategy_wise_listing"`
            PricePerPiece ProductStockPriceV2  `json:"price_per_piece"`
            Pincode float64  `json:"pincode"`
            Discount string  `json:"discount"`
            GroupedAttributes []SellerGroupAttributes  `json:"grouped_attributes"`
            IsGift bool  `json:"is_gift"`
            ArticleAssignment ArticleAssignmentV2  `json:"article_assignment"`
            Seller SellerV2  `json:"seller"`
         
    }
    
    // ProductSizeSellersResponseV3 ...
    type ProductSizeSellersResponseV3 struct {

        
            SortOn []ProductSizeSellerFilterSchemaV2  `json:"sort_on"`
            Page Page  `json:"page"`
            Items []ProductSizePriceResponseV2  `json:"items"`
         
    }
    

    
    // DiscountRulesApp ...
    type DiscountRulesApp struct {

        
            MatchedBuyRules []string  `json:"matched_buy_rules"`
            ItemCriteria map[string]interface{}  `json:"item_criteria"`
            RawOffer map[string]interface{}  `json:"raw_offer"`
            Offer map[string]interface{}  `json:"offer"`
         
    }
    
    // BuyRules ...
    type BuyRules struct {

        
            CartConditions map[string]interface{}  `json:"cart_conditions"`
            ItemCriteria map[string]interface{}  `json:"item_criteria"`
         
    }
    
    // Ownership ...
    type Ownership struct {

        
            PayableCategory string  `json:"payable_category"`
            PayableBy string  `json:"payable_by"`
         
    }
    
    // FreeGiftItem ...
    type FreeGiftItem struct {

        
            ItemPriceDetails map[string]interface{}  `json:"item_price_details"`
            ItemImagesURL []string  `json:"item_images_url"`
            ItemName string  `json:"item_name"`
            ItemBrandName string  `json:"item_brand_name"`
            ItemSlug string  `json:"item_slug"`
            ItemID float64  `json:"item_id"`
         
    }
    
    // AppliedFreeArticles ...
    type AppliedFreeArticles struct {

        
            ParentItemIdentifier string  `json:"parent_item_identifier"`
            ArticleID string  `json:"article_id"`
            FreeGiftItemDetails FreeGiftItem  `json:"free_gift_item_details"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // AppliedPromotion ...
    type AppliedPromotion struct {

        
            DiscountRules []DiscountRulesApp  `json:"discount_rules"`
            BuyRules []BuyRules  `json:"buy_rules"`
            Ownership Ownership  `json:"ownership"`
            PromoID string  `json:"promo_id"`
            AppliedFreeArticles []AppliedFreeArticles  `json:"applied_free_articles"`
            MrpPromotion bool  `json:"mrp_promotion"`
            PromotionType string  `json:"promotion_type"`
            Amount float64  `json:"amount"`
            OfferText string  `json:"offer_text"`
            PromotionName string  `json:"promotion_name"`
            ArticleQuantity float64  `json:"article_quantity"`
            PromotionGroup string  `json:"promotion_group"`
         
    }
    
    // ProductAvailability ...
    type ProductAvailability struct {

        
            Deliverable bool  `json:"deliverable"`
            IsValid bool  `json:"is_valid"`
            OutOfStock bool  `json:"out_of_stock"`
            OtherStoreQuantity float64  `json:"other_store_quantity"`
            Sizes []string  `json:"sizes"`
         
    }
    
    // PromoMeta ...
    type PromoMeta struct {

        
            Message string  `json:"message"`
         
    }
    
    // ProductPrice ...
    type ProductPrice struct {

        
            Effective float64  `json:"effective"`
            Selling float64  `json:"selling"`
            CurrencyCode string  `json:"currency_code"`
            AddOn float64  `json:"add_on"`
            CurrencySymbol string  `json:"currency_symbol"`
            Marked float64  `json:"marked"`
         
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
    
    // CategoryInfo ...
    type CategoryInfo struct {

        
            Name string  `json:"name"`
            UID float64  `json:"uid"`
         
    }
    
    // NetQuantity ...
    type NetQuantity struct {

        
            Value string  `json:"value"`
            Unit string  `json:"unit"`
         
    }
    
    // ProductImage ...
    type ProductImage struct {

        
            SecureURL string  `json:"secure_url"`
            AspectRatio string  `json:"aspect_ratio"`
            URL string  `json:"url"`
         
    }
    
    // BaseInfo ...
    type BaseInfo struct {

        
            Name string  `json:"name"`
            UID float64  `json:"uid"`
         
    }
    
    // ActionQuery ...
    type ActionQuery struct {

        
            ProductSlug []string  `json:"product_slug"`
         
    }
    
    // ProductAction ...
    type ProductAction struct {

        
            Query ActionQuery  `json:"query"`
            URL string  `json:"url"`
            Type string  `json:"type"`
         
    }
    
    // CartProduct ...
    type CartProduct struct {

        
            UID float64  `json:"uid"`
            Slug string  `json:"slug"`
            Categories []CategoryInfo  `json:"categories"`
            NetQuantity NetQuantity  `json:"net_quantity"`
            Images []ProductImage  `json:"images"`
            Brand BaseInfo  `json:"brand"`
            Type string  `json:"type"`
            Name string  `json:"name"`
            Action ProductAction  `json:"action"`
         
    }
    
    // BasePrice ...
    type BasePrice struct {

        
            Effective float64  `json:"effective"`
            CurrencySymbol string  `json:"currency_symbol"`
            Marked float64  `json:"marked"`
            CurrencyCode string  `json:"currency_code"`
         
    }
    
    // ArticlePriceInfo ...
    type ArticlePriceInfo struct {

        
            Base BasePrice  `json:"base"`
            Converted BasePrice  `json:"converted"`
         
    }
    
    // ProductArticle ...
    type ProductArticle struct {

        
            Size string  `json:"size"`
            UID string  `json:"uid"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            Seller BaseInfo  `json:"seller"`
            Quantity float64  `json:"quantity"`
            Price ArticlePriceInfo  `json:"price"`
            ProductGroupTags []string  `json:"product_group_tags"`
            Type string  `json:"type"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            Store BaseInfo  `json:"store"`
         
    }
    
    // CartProductInfo ...
    type CartProductInfo struct {

        
            PromotionsApplied []AppliedPromotion  `json:"promotions_applied"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            Availability ProductAvailability  `json:"availability"`
            PromoMeta PromoMeta  `json:"promo_meta"`
            Discount string  `json:"discount"`
            BulkOffer map[string]interface{}  `json:"bulk_offer"`
            Quantity float64  `json:"quantity"`
            Message string  `json:"message"`
            Price ProductPriceInfo  `json:"price"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
            CouponMessage string  `json:"coupon_message"`
            Product CartProduct  `json:"product"`
            Key string  `json:"key"`
            PricePerUnit ProductPriceInfo  `json:"price_per_unit"`
            IsSet bool  `json:"is_set"`
            Article ProductArticle  `json:"article"`
         
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
    
    // PaymentSelectionLock ...
    type PaymentSelectionLock struct {

        
            Enabled bool  `json:"enabled"`
            PaymentIdentifier string  `json:"payment_identifier"`
            DefaultOptions string  `json:"default_options"`
         
    }
    
    // RawBreakup ...
    type RawBreakup struct {

        
            ConvenienceFee float64  `json:"convenience_fee"`
            DeliveryCharge float64  `json:"delivery_charge"`
            Total float64  `json:"total"`
            Subtotal float64  `json:"subtotal"`
            MrpTotal float64  `json:"mrp_total"`
            CodCharge float64  `json:"cod_charge"`
            Vog float64  `json:"vog"`
            Coupon float64  `json:"coupon"`
            FyndCash float64  `json:"fynd_cash"`
            Discount float64  `json:"discount"`
            YouSaved float64  `json:"you_saved"`
            GstCharges float64  `json:"gst_charges"`
         
    }
    
    // DisplayBreakup ...
    type DisplayBreakup struct {

        
            Value float64  `json:"value"`
            Display string  `json:"display"`
            CurrencyCode string  `json:"currency_code"`
            Message []string  `json:"message"`
            CurrencySymbol string  `json:"currency_symbol"`
            Key string  `json:"key"`
         
    }
    
    // CouponBreakup ...
    type CouponBreakup struct {

        
            Title string  `json:"title"`
            Value float64  `json:"value"`
            UID string  `json:"uid"`
            SubTitle string  `json:"sub_title"`
            CouponType string  `json:"coupon_type"`
            Description string  `json:"description"`
            Message string  `json:"message"`
            IsApplied bool  `json:"is_applied"`
            Code string  `json:"code"`
            MaxDiscountValue float64  `json:"max_discount_value"`
            Type string  `json:"type"`
            MinimumCartValue float64  `json:"minimum_cart_value"`
            CouponValue float64  `json:"coupon_value"`
         
    }
    
    // LoyaltyPoints ...
    type LoyaltyPoints struct {

        
            IsApplied bool  `json:"is_applied"`
            Applicable float64  `json:"applicable"`
            Description string  `json:"description"`
            Total float64  `json:"total"`
         
    }
    
    // CartBreakup ...
    type CartBreakup struct {

        
            Raw RawBreakup  `json:"raw"`
            Display []DisplayBreakup  `json:"display"`
            Coupon CouponBreakup  `json:"coupon"`
            LoyaltyPoints LoyaltyPoints  `json:"loyalty_points"`
         
    }
    
    // CartCurrency ...
    type CartCurrency struct {

        
            Code string  `json:"code"`
            Symbol string  `json:"symbol"`
         
    }
    
    // CartDetailResponse ...
    type CartDetailResponse struct {

        
            LastModified string  `json:"last_modified"`
            Comment string  `json:"comment"`
            BuyNow bool  `json:"buy_now"`
            ID string  `json:"id"`
            Items []CartProductInfo  `json:"items"`
            PanNo string  `json:"pan_no"`
            IsValid bool  `json:"is_valid"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            CouponText string  `json:"coupon_text"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            Message string  `json:"message"`
            Gstin string  `json:"gstin"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            Currency CartCurrency  `json:"currency"`
            CheckoutMode string  `json:"checkout_mode"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
         
    }
    
    // AddProductCart ...
    type AddProductCart struct {

        
            Display string  `json:"display"`
            Pos bool  `json:"pos"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            StoreID float64  `json:"store_id"`
            ArticleID string  `json:"article_id"`
            Quantity float64  `json:"quantity"`
            ItemSize string  `json:"item_size"`
            SellerID float64  `json:"seller_id"`
            ProductGroupTags []string  `json:"product_group_tags"`
            ItemID float64  `json:"item_id"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            ArticleAssignment map[string]interface{}  `json:"article_assignment"`
         
    }
    
    // AddCartRequest ...
    type AddCartRequest struct {

        
            NewCart bool  `json:"new_cart"`
            Items []AddProductCart  `json:"items"`
         
    }
    
    // AddCartDetailResponse ...
    type AddCartDetailResponse struct {

        
            Message string  `json:"message"`
            Partial bool  `json:"partial"`
            Cart CartDetailResponse  `json:"cart"`
            Success bool  `json:"success"`
         
    }
    
    // UpdateProductCart ...
    type UpdateProductCart struct {

        
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            ArticleID string  `json:"article_id"`
            Quantity float64  `json:"quantity"`
            ItemSize string  `json:"item_size"`
            ItemIndex float64  `json:"item_index"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
            ItemID float64  `json:"item_id"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
         
    }
    
    // UpdateCartRequest ...
    type UpdateCartRequest struct {

        
            Operation string  `json:"operation"`
            Items []UpdateProductCart  `json:"items"`
         
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
    
    // PageCoupon ...
    type PageCoupon struct {

        
            Total float64  `json:"total"`
            TotalItemCount float64  `json:"total_item_count"`
            Current float64  `json:"current"`
            HasNext bool  `json:"has_next"`
            HasPrevious bool  `json:"has_previous"`
         
    }
    
    // Coupon ...
    type Coupon struct {

        
            Title string  `json:"title"`
            ExpiresOn string  `json:"expires_on"`
            SubTitle string  `json:"sub_title"`
            CouponType string  `json:"coupon_type"`
            Description string  `json:"description"`
            Message string  `json:"message"`
            IsApplied bool  `json:"is_applied"`
            MaxDiscountValue float64  `json:"max_discount_value"`
            MinimumCartValue float64  `json:"minimum_cart_value"`
            IsApplicable bool  `json:"is_applicable"`
            CouponCode string  `json:"coupon_code"`
            CouponValue float64  `json:"coupon_value"`
         
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
    
    // OfferPrice ...
    type OfferPrice struct {

        
            Effective float64  `json:"effective"`
            CurrencyCode string  `json:"currency_code"`
            CurrencySymbol string  `json:"currency_symbol"`
            Marked float64  `json:"marked"`
            BulkEffective float64  `json:"bulk_effective"`
         
    }
    
    // OfferItem ...
    type OfferItem struct {

        
            AutoApplied bool  `json:"auto_applied"`
            Total float64  `json:"total"`
            Quantity float64  `json:"quantity"`
            Best bool  `json:"best"`
            Price OfferPrice  `json:"price"`
            Type string  `json:"type"`
            Margin float64  `json:"margin"`
         
    }
    
    // OfferSeller ...
    type OfferSeller struct {

        
            Name string  `json:"name"`
            UID float64  `json:"uid"`
         
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

        
            ID string  `json:"id"`
            Country string  `json:"country"`
            Meta map[string]interface{}  `json:"meta"`
            GoogleMapPoint map[string]interface{}  `json:"google_map_point"`
            AreaCodeSlug string  `json:"area_code_slug"`
            CheckoutMode string  `json:"checkout_mode"`
            CountryCode string  `json:"country_code"`
            Email string  `json:"email"`
            AreaCode string  `json:"area_code"`
            UserID string  `json:"user_id"`
            GeoLocation GeoLocation  `json:"geo_location"`
            State string  `json:"state"`
            Name string  `json:"name"`
            Address string  `json:"address"`
            CountryPhoneCode string  `json:"country_phone_code"`
            IsDefaultAddress bool  `json:"is_default_address"`
            IsActive bool  `json:"is_active"`
            Landmark string  `json:"landmark"`
            CountryIsoCode string  `json:"country_iso_code"`
            Tags []string  `json:"tags"`
            City string  `json:"city"`
            CreatedByUserID string  `json:"created_by_user_id"`
            Area string  `json:"area"`
            AddressType string  `json:"address_type"`
            Phone string  `json:"phone"`
         
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

        
            ID string  `json:"id"`
            Success bool  `json:"success"`
            IsDefaultAddress bool  `json:"is_default_address"`
            IsUpdated bool  `json:"is_updated"`
         
    }
    
    // DeleteAddressResponse ...
    type DeleteAddressResponse struct {

        
            ID string  `json:"id"`
            IsDeleted bool  `json:"is_deleted"`
         
    }
    
    // SelectCartAddressRequest ...
    type SelectCartAddressRequest struct {

        
            ID string  `json:"id"`
            CartID string  `json:"cart_id"`
            BillingAddressID string  `json:"billing_address_id"`
         
    }
    
    // UpdateCartPaymentRequest ...
    type UpdateCartPaymentRequest struct {

        
            MerchantCode string  `json:"merchant_code"`
            ID string  `json:"id"`
            AggregatorName string  `json:"aggregator_name"`
            AddressID string  `json:"address_id"`
            PaymentMode string  `json:"payment_mode"`
            PaymentIdentifier string  `json:"payment_identifier"`
         
    }
    
    // CouponValidity ...
    type CouponValidity struct {

        
            Title string  `json:"title"`
            Code string  `json:"code"`
            Valid bool  `json:"valid"`
            Discount float64  `json:"discount"`
            DisplayMessageEn string  `json:"display_message_en"`
         
    }
    
    // PaymentCouponValidate ...
    type PaymentCouponValidate struct {

        
            Message string  `json:"message"`
            CouponValidity CouponValidity  `json:"coupon_validity"`
            Success bool  `json:"success"`
         
    }
    
    // ShipmentResponse ...
    type ShipmentResponse struct {

        
            FulfillmentID float64  `json:"fulfillment_id"`
            DpOptions map[string]interface{}  `json:"dp_options"`
            ShipmentType string  `json:"shipment_type"`
            Items []CartProductInfo  `json:"items"`
            Shipments float64  `json:"shipments"`
            Promise ShipmentPromise  `json:"promise"`
            OrderType string  `json:"order_type"`
            FulfillmentType string  `json:"fulfillment_type"`
            DpID string  `json:"dp_id"`
            BoxType string  `json:"box_type"`
         
    }
    
    // CartShipmentsResponse ...
    type CartShipmentsResponse struct {

        
            LastModified string  `json:"last_modified"`
            Comment string  `json:"comment"`
            ID string  `json:"id"`
            BuyNow bool  `json:"buy_now"`
            UID string  `json:"uid"`
            Shipments []ShipmentResponse  `json:"shipments"`
            CartID float64  `json:"cart_id"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            CouponText string  `json:"coupon_text"`
            Gstin string  `json:"gstin"`
            Error bool  `json:"error"`
            CheckoutMode string  `json:"checkout_mode"`
            IsValid bool  `json:"is_valid"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            Message string  `json:"message"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            Currency CartCurrency  `json:"currency"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
         
    }
    
    // StaffCheckout ...
    type StaffCheckout struct {

        
            EmployeeCode string  `json:"employee_code"`
            User string  `json:"user"`
            ID string  `json:"_id"`
            LastName string  `json:"last_name"`
            FirstName string  `json:"first_name"`
         
    }
    
    // CartCheckoutCustomMeta ...
    type CartCheckoutCustomMeta struct {

        
            Key string  `json:"key"`
            Value string  `json:"value"`
         
    }
    
    // CartCheckoutDetailRequest ...
    type CartCheckoutDetailRequest struct {

        
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
            MerchantCode string  `json:"merchant_code"`
            OrderingStore float64  `json:"ordering_store"`
            Meta map[string]interface{}  `json:"meta"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            AddressID string  `json:"address_id"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
            PaymentMode string  `json:"payment_mode"`
            PaymentIdentifier string  `json:"payment_identifier"`
            Aggregator string  `json:"aggregator"`
            Staff StaffCheckout  `json:"staff"`
            CustomMeta []CartCheckoutCustomMeta  `json:"custom_meta"`
            PaymentParams map[string]interface{}  `json:"payment_params"`
            CallbackURL string  `json:"callback_url"`
            PaymentAutoConfirm bool  `json:"payment_auto_confirm"`
            BillingAddressID string  `json:"billing_address_id"`
         
    }
    
    // CheckCart ...
    type CheckCart struct {

        
            LastModified string  `json:"last_modified"`
            Comment string  `json:"comment"`
            ID string  `json:"id"`
            UID string  `json:"uid"`
            DeliveryChargeOrderValue float64  `json:"delivery_charge_order_value"`
            CartID float64  `json:"cart_id"`
            CouponText string  `json:"coupon_text"`
            Gstin string  `json:"gstin"`
            CheckoutMode string  `json:"checkout_mode"`
            Success bool  `json:"success"`
            IsValid bool  `json:"is_valid"`
            UserType string  `json:"user_type"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            CodCharges float64  `json:"cod_charges"`
            DeliveryCharges float64  `json:"delivery_charges"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            BuyNow bool  `json:"buy_now"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            StoreEmps []map[string]interface{}  `json:"store_emps"`
            ErrorMessage string  `json:"error_message"`
            Items []CartProductInfo  `json:"items"`
            CodAvailable bool  `json:"cod_available"`
            OrderID string  `json:"order_id"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            Message string  `json:"message"`
            CodMessage string  `json:"cod_message"`
            StoreCode string  `json:"store_code"`
            Currency CartCurrency  `json:"currency"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
         
    }
    
    // CartCheckoutResponse ...
    type CartCheckoutResponse struct {

        
            Data map[string]interface{}  `json:"data"`
            AppInterceptURL string  `json:"app_intercept_url"`
            OrderID string  `json:"order_id"`
            PaymentConfirmURL string  `json:"payment_confirm_url"`
            Message string  `json:"message"`
            Cart CheckCart  `json:"cart"`
            CallbackURL string  `json:"callback_url"`
            Success bool  `json:"success"`
         
    }
    
    // CartMetaRequest ...
    type CartMetaRequest struct {

        
            Comment string  `json:"comment"`
            PickUpCustomerDetails map[string]interface{}  `json:"pick_up_customer_details"`
            CheckoutMode string  `json:"checkout_mode"`
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

        
            User map[string]interface{}  `json:"user"`
            Token string  `json:"token"`
            Source map[string]interface{}  `json:"source"`
            Meta map[string]interface{}  `json:"meta"`
            CreatedOn string  `json:"created_on"`
         
    }
    
    // SharedCart ...
    type SharedCart struct {

        
            LastModified string  `json:"last_modified"`
            Comment string  `json:"comment"`
            ID string  `json:"id"`
            UID string  `json:"uid"`
            CartID float64  `json:"cart_id"`
            CouponText string  `json:"coupon_text"`
            Gstin string  `json:"gstin"`
            CheckoutMode string  `json:"checkout_mode"`
            IsValid bool  `json:"is_valid"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            BuyNow bool  `json:"buy_now"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            Items []CartProductInfo  `json:"items"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            Message string  `json:"message"`
            SharedCartDetails SharedCartDetails  `json:"shared_cart_details"`
            Currency CartCurrency  `json:"currency"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
         
    }
    
    // SharedCartResponse ...
    type SharedCartResponse struct {

        
            Cart SharedCart  `json:"cart"`
            Error string  `json:"error"`
         
    }
    
    // FreeGiftItems ...
    type FreeGiftItems struct {

        
            ItemPriceDetails map[string]interface{}  `json:"item_price_details"`
            ItemImagesURL []string  `json:"item_images_url"`
            ItemName string  `json:"item_name"`
            ItemBrandName string  `json:"item_brand_name"`
            ItemSlug string  `json:"item_slug"`
            ItemID float64  `json:"item_id"`
         
    }
    
    // PromotionOffer ...
    type PromotionOffer struct {

        
            DiscountRules []map[string]interface{}  `json:"discount_rules"`
            BuyRules map[string]interface{}  `json:"buy_rules"`
            ID string  `json:"id"`
            Description string  `json:"description"`
            FreeGiftItems []FreeGiftItems  `json:"free_gift_items"`
            OfferText string  `json:"offer_text"`
            ValidTill string  `json:"valid_till"`
            PromotionGroup string  `json:"promotion_group"`
         
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

        
            OfferPrice float64  `json:"offer_price"`
            Effective float64  `json:"effective"`
            CurrencyCode string  `json:"currency_code"`
            CurrencySymbol string  `json:"currency_symbol"`
            Marked float64  `json:"marked"`
         
    }
    
    // LadderOfferItem ...
    type LadderOfferItem struct {

        
            MinQuantity float64  `json:"min_quantity"`
            Price LadderPrice  `json:"price"`
            Type string  `json:"type"`
            MaxQuantity float64  `json:"max_quantity"`
            Margin float64  `json:"margin"`
         
    }
    
    // LadderPriceOffer ...
    type LadderPriceOffer struct {

        
            DiscountRules []map[string]interface{}  `json:"discount_rules"`
            BuyRules map[string]interface{}  `json:"buy_rules"`
            ID string  `json:"id"`
            CalculateOn string  `json:"calculate_on"`
            Description string  `json:"description"`
            OfferPrices []LadderOfferItem  `json:"offer_prices"`
            FreeGiftItems []FreeGiftItems  `json:"free_gift_items"`
            OfferText string  `json:"offer_text"`
            ValidTill string  `json:"valid_till"`
            PromotionGroup string  `json:"promotion_group"`
         
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

        
            Pin string  `json:"pin"`
            Sdk bool  `json:"sdk"`
            Key string  `json:"key"`
            MerchantID string  `json:"merchant_id"`
            MerchantKey string  `json:"merchant_key"`
            UserID string  `json:"user_id"`
            ConfigType string  `json:"config_type"`
            VerifyAPI string  `json:"verify_api"`
            Secret string  `json:"secret"`
            API string  `json:"api"`
         
    }
    
    // AggregatorsConfigDetailResponse ...
    type AggregatorsConfigDetailResponse struct {

        
            Env string  `json:"env"`
            Success bool  `json:"success"`
            Stripe AggregatorConfigDetail  `json:"stripe"`
            Payumoney AggregatorConfigDetail  `json:"payumoney"`
            Razorpay AggregatorConfigDetail  `json:"razorpay"`
            Simpl AggregatorConfigDetail  `json:"simpl"`
            Mswipe AggregatorConfigDetail  `json:"mswipe"`
            Rupifi AggregatorConfigDetail  `json:"rupifi"`
            Juspay AggregatorConfigDetail  `json:"juspay"`
            Ccavenue AggregatorConfigDetail  `json:"ccavenue"`
         
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

        
            Nickname string  `json:"nickname"`
            NameOnCard string  `json:"name_on_card"`
            CardID string  `json:"card_id"`
            Refresh bool  `json:"refresh"`
         
    }
    
    // AttachCardsResponse ...
    type AttachCardsResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
            Data map[string]interface{}  `json:"data"`
         
    }
    
    // CardPaymentGateway ...
    type CardPaymentGateway struct {

        
            CustomerID string  `json:"customer_id"`
            API string  `json:"api"`
            Aggregator string  `json:"aggregator"`
         
    }
    
    // ActiveCardPaymentGatewayResponse ...
    type ActiveCardPaymentGatewayResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
            Cards CardPaymentGateway  `json:"cards"`
         
    }
    
    // Card ...
    type Card struct {

        
            AggregatorName string  `json:"aggregator_name"`
            ExpYear float64  `json:"exp_year"`
            CardIssuer string  `json:"card_issuer"`
            CardID string  `json:"card_id"`
            CardType string  `json:"card_type"`
            CardToken string  `json:"card_token"`
            CardBrand string  `json:"card_brand"`
            Nickname string  `json:"nickname"`
            CardReference string  `json:"card_reference"`
            Expired bool  `json:"expired"`
            CardNumber string  `json:"card_number"`
            ExpMonth float64  `json:"exp_month"`
            CardName string  `json:"card_name"`
            CardFingerprint string  `json:"card_fingerprint"`
            CardBrandImage string  `json:"card_brand_image"`
            CompliantWithTokenisationGuidelines bool  `json:"compliant_with_tokenisation_guidelines"`
            CardIsin string  `json:"card_isin"`
         
    }
    
    // ListCardsResponse ...
    type ListCardsResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
            Data []Card  `json:"data"`
         
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
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
            Payload string  `json:"payload"`
            PhoneNumber string  `json:"phone_number"`
            TransactionAmountInPaise float64  `json:"transaction_amount_in_paise"`
            Aggregator string  `json:"aggregator"`
            MerchantParams map[string]interface{}  `json:"merchant_params"`
         
    }
    
    // ValidateCustomerResponse ...
    type ValidateCustomerResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
            Data map[string]interface{}  `json:"data"`
         
    }
    
    // ChargeCustomerRequest ...
    type ChargeCustomerRequest struct {

        
            TransactionToken string  `json:"transaction_token"`
            OrderID string  `json:"order_id"`
            Amount float64  `json:"amount"`
            Verified bool  `json:"verified"`
            Aggregator string  `json:"aggregator"`
         
    }
    
    // ChargeCustomerResponse ...
    type ChargeCustomerResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
            OrderID string  `json:"order_id"`
            Status string  `json:"status"`
            Aggregator string  `json:"aggregator"`
            CartID string  `json:"cart_id"`
            DeliveryAddressID string  `json:"delivery_address_id"`
         
    }
    
    // PaymentInitializationRequest ...
    type PaymentInitializationRequest struct {

        
            Contact string  `json:"contact"`
            MerchantOrderID string  `json:"merchant_order_id"`
            Currency string  `json:"currency"`
            Vpa string  `json:"vpa"`
            RazorpayPaymentID string  `json:"razorpay_payment_id"`
            Email string  `json:"email"`
            OrderID string  `json:"order_id"`
            Amount float64  `json:"amount"`
            CustomerID string  `json:"customer_id"`
            Timeout float64  `json:"timeout"`
            Aggregator string  `json:"aggregator"`
            Method string  `json:"method"`
         
    }
    
    // PaymentInitializationResponse ...
    type PaymentInitializationResponse struct {

        
            MerchantOrderID string  `json:"merchant_order_id"`
            UpiPollURL string  `json:"upi_poll_url"`
            BqrImage string  `json:"bqr_image"`
            Currency string  `json:"currency"`
            Vpa string  `json:"vpa"`
            AggregatorOrderID string  `json:"aggregator_order_id"`
            RazorpayPaymentID string  `json:"razorpay_payment_id"`
            Success bool  `json:"success"`
            Amount float64  `json:"amount"`
            CustomerID string  `json:"customer_id"`
            VirtualID string  `json:"virtual_id"`
            Timeout float64  `json:"timeout"`
            Status string  `json:"status"`
            Aggregator string  `json:"aggregator"`
            PollingURL string  `json:"polling_url"`
            Method string  `json:"method"`
         
    }
    
    // PaymentStatusUpdateRequest ...
    type PaymentStatusUpdateRequest struct {

        
            Contact string  `json:"contact"`
            MerchantOrderID string  `json:"merchant_order_id"`
            Currency string  `json:"currency"`
            Vpa string  `json:"vpa"`
            Email string  `json:"email"`
            OrderID string  `json:"order_id"`
            Amount float64  `json:"amount"`
            CustomerID string  `json:"customer_id"`
            Status string  `json:"status"`
            Aggregator string  `json:"aggregator"`
            Method string  `json:"method"`
         
    }
    
    // PaymentStatusUpdateResponse ...
    type PaymentStatusUpdateResponse struct {

        
            AggregatorName string  `json:"aggregator_name"`
            Success bool  `json:"success"`
            RedirectURL string  `json:"redirect_url"`
            Status string  `json:"status"`
            Retry bool  `json:"retry"`
         
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

        
            Fynd AggregatorRoute  `json:"fynd"`
            Jiopay AggregatorRoute  `json:"jiopay"`
            BqrRazorpay AggregatorRoute  `json:"bqr_razorpay"`
            Stripe AggregatorRoute  `json:"stripe"`
            UpiRazorpay AggregatorRoute  `json:"upi_razorpay"`
            Epaylater AggregatorRoute  `json:"epaylater"`
            Razorpay AggregatorRoute  `json:"razorpay"`
            Simpl AggregatorRoute  `json:"simpl"`
            Mswipe AggregatorRoute  `json:"mswipe"`
            Rupifi AggregatorRoute  `json:"rupifi"`
            Payubiz AggregatorRoute  `json:"payubiz"`
            Juspay AggregatorRoute  `json:"juspay"`
            Ccavenue AggregatorRoute  `json:"ccavenue"`
         
    }
    
    // IntentAppErrorList ...
    type IntentAppErrorList struct {

        
            Code string  `json:"code"`
            PackageName string  `json:"package_name"`
         
    }
    
    // PaymentModeLogo ...
    type PaymentModeLogo struct {

        
            Small string  `json:"small"`
            Large string  `json:"large"`
         
    }
    
    // IntentApp ...
    type IntentApp struct {

        
            Code string  `json:"code"`
            PackageName string  `json:"package_name"`
            DisplayName string  `json:"display_name"`
            Logos PaymentModeLogo  `json:"logos"`
         
    }
    
    // PaymentModeList ...
    type PaymentModeList struct {

        
            IntentAppErrorDictList []IntentAppErrorList  `json:"intent_app_error_dict_list"`
            Name string  `json:"name"`
            CardID string  `json:"card_id"`
            CardBrand string  `json:"card_brand"`
            Code string  `json:"code"`
            CompliantWithTokenisationGuidelines bool  `json:"compliant_with_tokenisation_guidelines"`
            CardIssuer string  `json:"card_issuer"`
            CardType string  `json:"card_type"`
            DisplayName string  `json:"display_name"`
            IntentApp []IntentApp  `json:"intent_app"`
            RetryCount float64  `json:"retry_count"`
            Nickname string  `json:"nickname"`
            RemainingLimit float64  `json:"remaining_limit"`
            Timeout float64  `json:"timeout"`
            ExpMonth float64  `json:"exp_month"`
            CardBrandImage string  `json:"card_brand_image"`
            IntentFlow bool  `json:"intent_flow"`
            CodLimitPerOrder float64  `json:"cod_limit_per_order"`
            AggregatorName string  `json:"aggregator_name"`
            ExpYear float64  `json:"exp_year"`
            IntentAppErrorList []string  `json:"intent_app_error_list"`
            FyndVpa string  `json:"fynd_vpa"`
            DisplayPriority float64  `json:"display_priority"`
            LogoURL PaymentModeLogo  `json:"logo_url"`
            CodLimit float64  `json:"cod_limit"`
            CardToken string  `json:"card_token"`
            CardReference string  `json:"card_reference"`
            Expired bool  `json:"expired"`
            CardNumber string  `json:"card_number"`
            CardName string  `json:"card_name"`
            MerchantCode string  `json:"merchant_code"`
            CardFingerprint string  `json:"card_fingerprint"`
            CardIsin string  `json:"card_isin"`
         
    }
    
    // RootPaymentMode ...
    type RootPaymentMode struct {

        
            AggregatorName string  `json:"aggregator_name"`
            AnonymousEnable bool  `json:"anonymous_enable"`
            Name string  `json:"name"`
            DisplayPriority float64  `json:"display_priority"`
            AddCardEnabled bool  `json:"add_card_enabled"`
            SaveCard bool  `json:"save_card"`
            List []PaymentModeList  `json:"list"`
            IsPayByCardPl bool  `json:"is_pay_by_card_pl"`
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

        
            RequestType string  `json:"request_type"`
            OrderID string  `json:"order_id"`
         
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

        
            UpiVpa string  `json:"upi_vpa"`
            CustomerName string  `json:"customer_name"`
            IsValid bool  `json:"is_valid"`
            Status string  `json:"status"`
         
    }
    
    // ValidateVPAResponse ...
    type ValidateVPAResponse struct {

        
            Success bool  `json:"success"`
            Data ValidateUPI  `json:"data"`
         
    }
    
    // TransferItemsDetails ...
    type TransferItemsDetails struct {

        
            ID float64  `json:"id"`
            Name string  `json:"name"`
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

        
            TransferMode string  `json:"transfer_mode"`
            Enable bool  `json:"enable"`
         
    }
    
    // UpdateRefundTransferModeResponse ...
    type UpdateRefundTransferModeResponse struct {

        
            Success bool  `json:"success"`
         
    }
    
    // OrderBeneficiaryDetails ...
    type OrderBeneficiaryDetails struct {

        
            ID float64  `json:"id"`
            Comment string  `json:"comment"`
            BeneficiaryID string  `json:"beneficiary_id"`
            BankName string  `json:"bank_name"`
            BranchName string  `json:"branch_name"`
            Mobile string  `json:"mobile"`
            AccountNo string  `json:"account_no"`
            Address string  `json:"address"`
            CreatedOn string  `json:"created_on"`
            AccountHolder string  `json:"account_holder"`
            DisplayName string  `json:"display_name"`
            IsActive bool  `json:"is_active"`
            IfscCode string  `json:"ifsc_code"`
            Email string  `json:"email"`
            Subtitle string  `json:"subtitle"`
            ModifiedOn string  `json:"modified_on"`
            TransferMode string  `json:"transfer_mode"`
            DelightsUserName string  `json:"delights_user_name"`
            Title string  `json:"title"`
         
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

        
            BankName string  `json:"bank_name"`
            Success bool  `json:"success"`
            BranchName string  `json:"branch_name"`
         
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
            Otp string  `json:"otp"`
            HashKey string  `json:"hash_key"`
         
    }
    
    // AddBeneficiaryViaOtpVerificationResponse ...
    type AddBeneficiaryViaOtpVerificationResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // WrongOtpError ...
    type WrongOtpError struct {

        
            Success string  `json:"success"`
            Description string  `json:"description"`
            IsVerifiedFlag bool  `json:"is_verified_flag"`
         
    }
    
    // BeneficiaryModeDetails ...
    type BeneficiaryModeDetails struct {

        
            Vpa string  `json:"vpa"`
            AccountNo string  `json:"account_no"`
            Wallet string  `json:"wallet"`
            Address string  `json:"address"`
            Email string  `json:"email"`
            Comment string  `json:"comment"`
            IfscCode string  `json:"ifsc_code"`
            BankName string  `json:"bank_name"`
            BranchName string  `json:"branch_name"`
            AccountHolder string  `json:"account_holder"`
            Mobile string  `json:"mobile"`
         
    }
    
    // AddBeneficiaryDetailsRequest ...
    type AddBeneficiaryDetailsRequest struct {

        
            OrderID string  `json:"order_id"`
            Otp string  `json:"otp"`
            RequestID string  `json:"request_id"`
            Delights bool  `json:"delights"`
            ShipmentID string  `json:"shipment_id"`
            Details BeneficiaryModeDetails  `json:"details"`
            TransferMode string  `json:"transfer_mode"`
         
    }
    
    // RefundAccountResponse ...
    type RefundAccountResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
            IsVerifiedFlag bool  `json:"is_verified_flag"`
            Data map[string]interface{}  `json:"data"`
         
    }
    
    // BankDetailsForOTP ...
    type BankDetailsForOTP struct {

        
            AccountNo string  `json:"account_no"`
            IfscCode string  `json:"ifsc_code"`
            BankName string  `json:"bank_name"`
            BranchName string  `json:"branch_name"`
            AccountHolder string  `json:"account_holder"`
         
    }
    
    // AddBeneficiaryDetailsOTPRequest ...
    type AddBeneficiaryDetailsOTPRequest struct {

        
            OrderID string  `json:"order_id"`
            Details BankDetailsForOTP  `json:"details"`
         
    }
    
    // WalletOtpRequest ...
    type WalletOtpRequest struct {

        
            Mobile string  `json:"mobile"`
            CountryCode string  `json:"country_code"`
         
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

        
            Success bool  `json:"success"`
            IsBeneficiarySet bool  `json:"is_beneficiary_set"`
         
    }
    
    // GetPaymentLinkResponse ...
    type GetPaymentLinkResponse struct {

        
            PaymentLinkURL string  `json:"payment_link_url"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
            StatusCode float64  `json:"status_code"`
            MerchantName string  `json:"merchant_name"`
            ExternalOrderID string  `json:"external_order_id"`
            Amount float64  `json:"amount"`
            PollingTimeout float64  `json:"polling_timeout"`
            PaymentLinkCurrentStatus string  `json:"payment_link_current_status"`
         
    }
    
    // ErrorDescription ...
    type ErrorDescription struct {

        
            MerchantOrderID string  `json:"merchant_order_id"`
            MerchantName string  `json:"merchant_name"`
            InvalidID bool  `json:"invalid_id"`
            PaymentTransactionID string  `json:"payment_transaction_id"`
            Msg string  `json:"msg"`
            Amount float64  `json:"amount"`
            Cancelled bool  `json:"cancelled"`
            Expired bool  `json:"expired"`
         
    }
    
    // ErrorResponse ...
    type ErrorResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
            StatusCode float64  `json:"status_code"`
            Error ErrorDescription  `json:"error"`
         
    }
    
    // CreatePaymentLinkMeta ...
    type CreatePaymentLinkMeta struct {

        
            AssignCardID string  `json:"assign_card_id"`
            Amount string  `json:"amount"`
            Pincode string  `json:"pincode"`
            CheckoutMode string  `json:"checkout_mode"`
            CartID string  `json:"cart_id"`
         
    }
    
    // CreatePaymentLinkRequest ...
    type CreatePaymentLinkRequest struct {

        
            Email string  `json:"email"`
            ExternalOrderID string  `json:"external_order_id"`
            Amount float64  `json:"amount"`
            MobileNumber string  `json:"mobile_number"`
            Meta CreatePaymentLinkMeta  `json:"meta"`
            Description string  `json:"description"`
         
    }
    
    // CreatePaymentLinkResponse ...
    type CreatePaymentLinkResponse struct {

        
            PaymentLinkURL string  `json:"payment_link_url"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
            StatusCode float64  `json:"status_code"`
            PollingTimeout float64  `json:"polling_timeout"`
            PaymentLinkID string  `json:"payment_link_id"`
         
    }
    
    // CancelOrResendPaymentLinkRequest ...
    type CancelOrResendPaymentLinkRequest struct {

        
            PaymentLinkID string  `json:"payment_link_id"`
         
    }
    
    // ResendPaymentLinkResponse ...
    type ResendPaymentLinkResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
            PollingTimeout float64  `json:"polling_timeout"`
            StatusCode float64  `json:"status_code"`
         
    }
    
    // CancelPaymentLinkResponse ...
    type CancelPaymentLinkResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
            StatusCode float64  `json:"status_code"`
         
    }
    
    // PollingPaymentLinkResponse ...
    type PollingPaymentLinkResponse struct {

        
            AggregatorName string  `json:"aggregator_name"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
            StatusCode float64  `json:"status_code"`
            OrderID string  `json:"order_id"`
            Amount float64  `json:"amount"`
            RedirectURL string  `json:"redirect_url"`
            HttpStatus float64  `json:"http_status"`
            Status string  `json:"status"`
            PaymentLinkID string  `json:"payment_link_id"`
         
    }
    
    // PaymentMethodsMeta ...
    type PaymentMethodsMeta struct {

        
            MerchantCode string  `json:"merchant_code"`
            PaymentIdentifier string  `json:"payment_identifier"`
            PaymentGateway string  `json:"payment_gateway"`
         
    }
    
    // CreateOrderUserPaymentMethods ...
    type CreateOrderUserPaymentMethods struct {

        
            Meta PaymentMethodsMeta  `json:"meta"`
            Name string  `json:"name"`
            Mode string  `json:"mode"`
         
    }
    
    // CreateOrderUserRequest ...
    type CreateOrderUserRequest struct {

        
            Currency string  `json:"currency"`
            PaymentMethods CreateOrderUserPaymentMethods  `json:"payment_methods"`
            Meta map[string]interface{}  `json:"meta"`
            PaymentLinkID string  `json:"payment_link_id"`
            SuccessCallbackURL string  `json:"success_callback_url"`
            FailureCallbackURL string  `json:"failure_callback_url"`
         
    }
    
    // CreateOrderUserData ...
    type CreateOrderUserData struct {

        
            Contact string  `json:"contact"`
            MerchantOrderID string  `json:"merchant_order_id"`
            Currency string  `json:"currency"`
            Email string  `json:"email"`
            OrderID string  `json:"order_id"`
            CallbackURL string  `json:"callback_url"`
            Amount float64  `json:"amount"`
            CustomerID string  `json:"customer_id"`
            Aggregator string  `json:"aggregator"`
            Method string  `json:"method"`
         
    }
    
    // CreateOrderUserResponse ...
    type CreateOrderUserResponse struct {

        
            Data CreateOrderUserData  `json:"data"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
            StatusCode float64  `json:"status_code"`
            CallbackURL string  `json:"callback_url"`
            OrderID string  `json:"order_id"`
            PaymentConfirmURL string  `json:"payment_confirm_url"`
         
    }
    
    // BalanceDetails ...
    type BalanceDetails struct {

        
            Currency string  `json:"currency"`
            FormattedValue string  `json:"formatted_value"`
            Value float64  `json:"value"`
         
    }
    
    // CreditSummary ...
    type CreditSummary struct {

        
            MerchantCustomerRefID string  `json:"merchant_customer_ref_id"`
            Balance BalanceDetails  `json:"balance"`
            Status string  `json:"status"`
            StatusMessage string  `json:"status_message"`
         
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
    
    // MarketplaceInfo ...
    type MarketplaceInfo struct {

        
            MembershipID string  `json:"membership_id"`
            DateOfJoining string  `json:"date_of_joining"`
            Name string  `json:"name"`
         
    }
    
    // KYCAddress ...
    type KYCAddress struct {

        
            State string  `json:"state"`
            Addressline2 string  `json:"addressline2"`
            Pincode string  `json:"pincode"`
            OwnershipType string  `json:"ownership_type"`
            Addressline1 string  `json:"addressline1"`
            City string  `json:"city"`
            LandMark string  `json:"land_mark"`
         
    }
    
    // BusinessDetails ...
    type BusinessDetails struct {

        
            Fssai string  `json:"fssai"`
            Name string  `json:"name"`
            Gstin string  `json:"gstin"`
            EntityType string  `json:"entity_type"`
            BusinessOwnershipType string  `json:"business_ownership_type"`
            Address KYCAddress  `json:"address"`
            BusinessType string  `json:"business_type"`
            ShopAndEstablishment map[string]interface{}  `json:"shop_and_establishment"`
            Vintage string  `json:"vintage"`
            Pan string  `json:"pan"`
            Fda string  `json:"fda"`
         
    }
    
    // UserPersonalInfoInDetails ...
    type UserPersonalInfoInDetails struct {

        
            MiddleName string  `json:"middle_name"`
            FathersName string  `json:"fathers_name"`
            EmailVerified bool  `json:"email_verified"`
            MothersName string  `json:"mothers_name"`
            Email string  `json:"email"`
            Phone string  `json:"phone"`
            Passport string  `json:"passport"`
            MobileVerified bool  `json:"mobile_verified"`
            FirstName string  `json:"first_name"`
            Pan string  `json:"pan"`
            AddressAsPerID KYCAddress  `json:"address_as_per_id"`
            Dob string  `json:"dob"`
            DrivingLicense string  `json:"driving_license"`
            VoterID string  `json:"voter_id"`
            LastName string  `json:"last_name"`
            Gender string  `json:"gender"`
         
    }
    
    // DeviceDetails ...
    type DeviceDetails struct {

        
            OsVersion string  `json:"os_version"`
            IdentifierType string  `json:"identifier_type"`
            DeviceModel string  `json:"device_model"`
            Os string  `json:"os"`
            DeviceType string  `json:"device_type"`
            IdentificationNumber string  `json:"identification_number"`
            DeviceMake string  `json:"device_make"`
         
    }
    
    // CustomerOnboardingRequest ...
    type CustomerOnboardingRequest struct {

        
            MarketplaceInfo MarketplaceInfo  `json:"marketplace_info"`
            BusinessInfo BusinessDetails  `json:"business_info"`
            PersonalInfo UserPersonalInfoInDetails  `json:"personal_info"`
            Device DeviceDetails  `json:"device"`
            Source string  `json:"source"`
            Aggregator string  `json:"aggregator"`
            Mcc string  `json:"mcc"`
         
    }
    
    // OnboardSummary ...
    type OnboardSummary struct {

        
            Session map[string]interface{}  `json:"session"`
            Status bool  `json:"status"`
            RedirectURL string  `json:"redirect_url"`
         
    }
    
    // CustomerOnboardingResponse ...
    type CustomerOnboardingResponse struct {

        
            Success bool  `json:"success"`
            Data OnboardSummary  `json:"data"`
         
    }
    

    
    // ShipmentUserInfo ...
    type ShipmentUserInfo struct {

        
            Mobile string  `json:"mobile"`
            LastName string  `json:"last_name"`
            FirstName string  `json:"first_name"`
            Gender string  `json:"gender"`
         
    }
    
    // ShipmentPayment ...
    type ShipmentPayment struct {

        
            Logo string  `json:"logo"`
            DisplayName string  `json:"display_name"`
            Mop string  `json:"mop"`
            Status string  `json:"status"`
            Mode string  `json:"mode"`
            PaymentMode string  `json:"payment_mode"`
         
    }
    
    // Prices ...
    type Prices struct {

        
            AmountPaidRoundoff float64  `json:"amount_paid_roundoff"`
            RefundCredit float64  `json:"refund_credit"`
            Discount float64  `json:"discount"`
            AmountPaid float64  `json:"amount_paid"`
            CashbackApplied float64  `json:"cashback_applied"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            ValueOfGood float64  `json:"value_of_good"`
            RefundAmount float64  `json:"refund_amount"`
            PriceEffective float64  `json:"price_effective"`
            AddedToFyndCash bool  `json:"added_to_fynd_cash"`
            PromotionEffectiveDiscount float64  `json:"promotion_effective_discount"`
            CouponValue float64  `json:"coupon_value"`
            TransferPrice float64  `json:"transfer_price"`
            DeliveryCharge float64  `json:"delivery_charge"`
            CodCharges float64  `json:"cod_charges"`
            PriceMarked float64  `json:"price_marked"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            FyndCredits float64  `json:"fynd_credits"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            Cashback float64  `json:"cashback"`
         
    }
    
    // ShipmentTotalDetails ...
    type ShipmentTotalDetails struct {

        
            Pieces float64  `json:"pieces"`
            Sizes float64  `json:"sizes"`
            TotalPrice float64  `json:"total_price"`
         
    }
    
    // DeliveryAddress ...
    type DeliveryAddress struct {

        
            Latitude float64  `json:"latitude"`
            Address string  `json:"address"`
            Pincode string  `json:"pincode"`
            Address1 string  `json:"address1"`
            Longitude float64  `json:"longitude"`
            Email string  `json:"email"`
            Name string  `json:"name"`
            UpdatedAt string  `json:"updated_at"`
            City string  `json:"city"`
            CreatedAt string  `json:"created_at"`
            Country string  `json:"country"`
            Area string  `json:"area"`
            Version string  `json:"version"`
            AddressType string  `json:"address_type"`
            ContactPerson string  `json:"contact_person"`
            AddressCategory string  `json:"address_category"`
            Phone string  `json:"phone"`
            State string  `json:"state"`
            Address2 string  `json:"address2"`
            Landmark string  `json:"landmark"`
         
    }
    
    // FulfillingCompany ...
    type FulfillingCompany struct {

        
            Name string  `json:"name"`
            ID float64  `json:"id"`
         
    }
    
    // FulfillingStore ...
    type FulfillingStore struct {

        
            ID float64  `json:"id"`
            Code string  `json:"code"`
            CompanyID float64  `json:"company_id"`
            Name string  `json:"name"`
            CompanyName string  `json:"company_name"`
         
    }
    
    // ShipmentStatus ...
    type ShipmentStatus struct {

        
            Title string  `json:"title"`
            HexCode string  `json:"hex_code"`
         
    }
    
    // ItemBrand ...
    type ItemBrand struct {

        
            Logo string  `json:"logo"`
            Name string  `json:"name"`
         
    }
    
    // Item ...
    type Item struct {

        
            Image []string  `json:"image"`
            ID float64  `json:"id"`
            Code string  `json:"code"`
            SellerIdentifier string  `json:"seller_identifier"`
            Brand ItemBrand  `json:"brand"`
            Name string  `json:"name"`
            SlugKey string  `json:"slug_key"`
            Size string  `json:"size"`
         
    }
    
    // AppliedFreeArticles ...
    type AppliedFreeArticles struct {

        
            ParentItemIdentifier string  `json:"parent_item_identifier"`
            ArticleID string  `json:"article_id"`
            FreeGiftItemDetails map[string]interface{}  `json:"free_gift_item_details"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // AppliedPromos ...
    type AppliedPromos struct {

        
            MrpPromotion bool  `json:"mrp_promotion"`
            Amount float64  `json:"amount"`
            PromotionType string  `json:"promotion_type"`
            PromotionName string  `json:"promotion_name"`
            AppliedFreeArticles []AppliedFreeArticles  `json:"applied_free_articles"`
            ArticleQuantity float64  `json:"article_quantity"`
            PromoID string  `json:"promo_id"`
         
    }
    
    // CurrentStatus ...
    type CurrentStatus struct {

        
            JourneyType string  `json:"journey_type"`
            Name string  `json:"name"`
            Status string  `json:"status"`
            UpdatedAt string  `json:"updated_at"`
         
    }
    
    // Identifiers ...
    type Identifiers struct {

        
            SkuCode string  `json:"sku_code"`
            Ean string  `json:"ean"`
         
    }
    
    // FinancialBreakup ...
    type FinancialBreakup struct {

        
            AmountPaidRoundoff float64  `json:"amount_paid_roundoff"`
            RefundCredit float64  `json:"refund_credit"`
            Discount float64  `json:"discount"`
            AmountPaid float64  `json:"amount_paid"`
            CashbackApplied float64  `json:"cashback_applied"`
            ValueOfGood float64  `json:"value_of_good"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            GstFee float64  `json:"gst_fee"`
            RefundAmount float64  `json:"refund_amount"`
            PriceEffective float64  `json:"price_effective"`
            AddedToFyndCash bool  `json:"added_to_fynd_cash"`
            GstTag string  `json:"gst_tag"`
            PromotionEffectiveDiscount float64  `json:"promotion_effective_discount"`
            CouponValue float64  `json:"coupon_value"`
            TransferPrice float64  `json:"transfer_price"`
            DeliveryCharge float64  `json:"delivery_charge"`
            TotalUnits float64  `json:"total_units"`
            CodCharges float64  `json:"cod_charges"`
            ItemName string  `json:"item_name"`
            PriceMarked float64  `json:"price_marked"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            HsnCode string  `json:"hsn_code"`
            FyndCredits float64  `json:"fynd_credits"`
            Identifiers Identifiers  `json:"identifiers"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            Cashback float64  `json:"cashback"`
            Size string  `json:"size"`
         
    }
    
    // Bags ...
    type Bags struct {

        
            Item Item  `json:"item"`
            ID float64  `json:"id"`
            Quantity float64  `json:"quantity"`
            SellerIdentifier string  `json:"seller_identifier"`
            DeliveryDate string  `json:"delivery_date"`
            CanCancel bool  `json:"can_cancel"`
            AppliedPromos []AppliedPromos  `json:"applied_promos"`
            CanReturn bool  `json:"can_return"`
            ParentPromoBags map[string]interface{}  `json:"parent_promo_bags"`
            CurrentStatus CurrentStatus  `json:"current_status"`
            ReturnableDate string  `json:"returnable_date"`
            Prices Prices  `json:"prices"`
            FinancialBreakup []FinancialBreakup  `json:"financial_breakup"`
            LineNumber float64  `json:"line_number"`
         
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

        
            IsPassed bool  `json:"is_passed"`
            IsCurrent bool  `json:"is_current"`
            Time string  `json:"time"`
            Status string  `json:"status"`
         
    }
    
    // TrackingDetails ...
    type TrackingDetails struct {

        
            IsCurrent bool  `json:"is_current"`
            TrackingDetails []NestedTrackingDetails  `json:"tracking_details"`
            IsPassed bool  `json:"is_passed"`
            Status string  `json:"status"`
            Time string  `json:"time"`
         
    }
    
    // Invoice ...
    type Invoice struct {

        
            LabelURL string  `json:"label_url"`
            UpdatedDate string  `json:"updated_date"`
            InvoiceURL string  `json:"invoice_url"`
         
    }
    
    // BreakupValues ...
    type BreakupValues struct {

        
            Display string  `json:"display"`
            Name string  `json:"name"`
            Value float64  `json:"value"`
         
    }
    
    // Shipments ...
    type Shipments struct {

        
            CustomMeta []map[string]interface{}  `json:"custom_meta"`
            UserInfo ShipmentUserInfo  `json:"user_info"`
            CanCancel bool  `json:"can_cancel"`
            CanReturn bool  `json:"can_return"`
            Payment ShipmentPayment  `json:"payment"`
            Prices Prices  `json:"prices"`
            TotalDetails ShipmentTotalDetails  `json:"total_details"`
            DeliveryAddress DeliveryAddress  `json:"delivery_address"`
            ShipmentID string  `json:"shipment_id"`
            ShowDownloadInvoice bool  `json:"show_download_invoice"`
            TrackURL string  `json:"track_url"`
            AwbNo string  `json:"awb_no"`
            ShipmentCreatedAt string  `json:"shipment_created_at"`
            FulfillingCompany FulfillingCompany  `json:"fulfilling_company"`
            Comment string  `json:"comment"`
            RefundDetails map[string]interface{}  `json:"refund_details"`
            FulfillingStore FulfillingStore  `json:"fulfilling_store"`
            OrderID string  `json:"order_id"`
            ShipmentStatus ShipmentStatus  `json:"shipment_status"`
            DpName string  `json:"dp_name"`
            SizeInfo map[string]interface{}  `json:"size_info"`
            Bags []Bags  `json:"bags"`
            TrakingNo string  `json:"traking_no"`
            Promise Promise  `json:"promise"`
            ShowTrackLink bool  `json:"show_track_link"`
            OrderType string  `json:"order_type"`
            CanBreak map[string]interface{}  `json:"can_break"`
            TrackingDetails []TrackingDetails  `json:"tracking_details"`
            BeneficiaryDetails bool  `json:"beneficiary_details"`
            DeliveryDate string  `json:"delivery_date"`
            NeedHelpURL string  `json:"need_help_url"`
            ReturnableDate string  `json:"returnable_date"`
            TotalBags float64  `json:"total_bags"`
            Invoice Invoice  `json:"invoice"`
            BreakupValues []BreakupValues  `json:"breakup_values"`
         
    }
    
    // BagsForReorderArticleAssignment ...
    type BagsForReorderArticleAssignment struct {

        
            Level string  `json:"level"`
            Strategy string  `json:"strategy"`
         
    }
    
    // BagsForReorder ...
    type BagsForReorder struct {

        
            ItemID float64  `json:"item_id"`
            StoreID float64  `json:"store_id"`
            Quantity float64  `json:"quantity"`
            SellerID float64  `json:"seller_id"`
            ItemSize string  `json:"item_size"`
            ArticleAssignment BagsForReorderArticleAssignment  `json:"article_assignment"`
         
    }
    
    // UserInfo ...
    type UserInfo struct {

        
            Mobile string  `json:"mobile"`
            Gender string  `json:"gender"`
            Name string  `json:"name"`
            Email string  `json:"email"`
         
    }
    
    // OrderSchema ...
    type OrderSchema struct {

        
            Shipments []Shipments  `json:"shipments"`
            BagsForReorder []BagsForReorder  `json:"bags_for_reorder"`
            OrderID string  `json:"order_id"`
            UserInfo UserInfo  `json:"user_info"`
            OrderCreatedTime string  `json:"order_created_time"`
            TotalShipmentsInOrder float64  `json:"total_shipments_in_order"`
            BreakupValues []BreakupValues  `json:"breakup_values"`
         
    }
    
    // OrderPage ...
    type OrderPage struct {

        
            Type string  `json:"type"`
            ItemTotal float64  `json:"item_total"`
            Current float64  `json:"current"`
            HasNext bool  `json:"has_next"`
            Size float64  `json:"size"`
         
    }
    
    // OrderStatuses ...
    type OrderStatuses struct {

        
            Display string  `json:"display"`
            Value float64  `json:"value"`
            IsSelected bool  `json:"is_selected"`
         
    }
    
    // OrderFilters ...
    type OrderFilters struct {

        
            Statuses []OrderStatuses  `json:"statuses"`
         
    }
    
    // OrderList ...
    type OrderList struct {

        
            Items []OrderSchema  `json:"items"`
            Page OrderPage  `json:"page"`
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

        
            UpdatedAt string  `json:"updated_at"`
            AccountName string  `json:"account_name"`
            LastLocationRecievedAt string  `json:"last_location_recieved_at"`
            ShipmentType string  `json:"shipment_type"`
            Status string  `json:"status"`
            Reason string  `json:"reason"`
            Awb string  `json:"awb"`
            UpdatedTime string  `json:"updated_time"`
         
    }
    
    // ShipmentTrack ...
    type ShipmentTrack struct {

        
            Results []Track  `json:"results"`
         
    }
    
    // CustomerDetailsResponse ...
    type CustomerDetailsResponse struct {

        
            OrderID string  `json:"order_id"`
            Country string  `json:"country"`
            Phone string  `json:"phone"`
            Name string  `json:"name"`
            ShipmentID string  `json:"shipment_id"`
         
    }
    
    // SendOtpToCustomerResponse ...
    type SendOtpToCustomerResponse struct {

        
            Success bool  `json:"success"`
            RequestID string  `json:"request_id"`
            Message string  `json:"message"`
            ResendTimer float64  `json:"resend_timer"`
         
    }
    
    // VerifyOtp ...
    type VerifyOtp struct {

        
            OtpCode string  `json:"otp_code"`
            RequestID string  `json:"request_id"`
         
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
            Reasons []BagReasons  `json:"reasons"`
            DisplayName string  `json:"display_name"`
            QuestionSet []QuestionSet  `json:"question_set"`
            Meta BagReasonMeta  `json:"meta"`
            QcType []string  `json:"qc_type"`
         
    }
    
    // ShipmentBagReasons ...
    type ShipmentBagReasons struct {

        
            Success bool  `json:"success"`
            Reasons []BagReasons  `json:"reasons"`
         
    }
    
    // ShipmentReason ...
    type ShipmentReason struct {

        
            ShowTextArea bool  `json:"show_text_area"`
            ReasonID float64  `json:"reason_id"`
            FeedbackType string  `json:"feedback_type"`
            Priority float64  `json:"priority"`
            ReasonText string  `json:"reason_text"`
            Flow string  `json:"flow"`
         
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
    
    // ProductsReasonsData ...
    type ProductsReasonsData struct {

        
            ReasonID float64  `json:"reason_id"`
            ReasonText string  `json:"reason_text"`
         
    }
    
    // ProductsReasonsFilters ...
    type ProductsReasonsFilters struct {

        
            LineNumber float64  `json:"line_number"`
            Quantity float64  `json:"quantity"`
            Identifier string  `json:"identifier"`
         
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
    
    // ProductsDataUpdatesFilters ...
    type ProductsDataUpdatesFilters struct {

        
            LineNumber float64  `json:"line_number"`
            Identifier string  `json:"identifier"`
         
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
    
    // ShipmentsRequest ...
    type ShipmentsRequest struct {

        
            Products []Products  `json:"products"`
            Reasons ReasonsData  `json:"reasons"`
            DataUpdates DataUpdates  `json:"data_updates"`
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

        
            Task bool  `json:"task"`
            ForceTransition bool  `json:"force_transition"`
            Statuses []StatuesRequest  `json:"statuses"`
            UnlockBeforeTransition bool  `json:"unlock_before_transition"`
            LockAfterTransition bool  `json:"lock_after_transition"`
         
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
            Status float64  `json:"status"`
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
    

    
    // NetQuantity ...
    type NetQuantity struct {

        
            Value string  `json:"value"`
            Unit string  `json:"unit"`
         
    }
    
    // ActionQuery ...
    type ActionQuery struct {

        
            ProductSlug []string  `json:"product_slug"`
         
    }
    
    // ProductAction ...
    type ProductAction struct {

        
            Type string  `json:"type"`
            Query ActionQuery  `json:"query"`
            URL string  `json:"url"`
         
    }
    
    // BaseInfo ...
    type BaseInfo struct {

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // CategoryInfo ...
    type CategoryInfo struct {

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // ProductImage ...
    type ProductImage struct {

        
            AspectRatio string  `json:"aspect_ratio"`
            SecureURL string  `json:"secure_url"`
            URL string  `json:"url"`
         
    }
    
    // CartProduct ...
    type CartProduct struct {

        
            UID float64  `json:"uid"`
            NetQuantity NetQuantity  `json:"net_quantity"`
            Action ProductAction  `json:"action"`
            Name string  `json:"name"`
            Type string  `json:"type"`
            Brand BaseInfo  `json:"brand"`
            Categories []CategoryInfo  `json:"categories"`
            Slug string  `json:"slug"`
            Images []ProductImage  `json:"images"`
         
    }
    
    // BuyRules ...
    type BuyRules struct {

        
            CartConditions map[string]interface{}  `json:"cart_conditions"`
            ItemCriteria map[string]interface{}  `json:"item_criteria"`
         
    }
    
    // Ownership ...
    type Ownership struct {

        
            PayableCategory string  `json:"payable_category"`
            PayableBy string  `json:"payable_by"`
         
    }
    
    // DiscountRulesApp ...
    type DiscountRulesApp struct {

        
            Offer map[string]interface{}  `json:"offer"`
            ItemCriteria map[string]interface{}  `json:"item_criteria"`
            MatchedBuyRules []string  `json:"matched_buy_rules"`
            RawOffer map[string]interface{}  `json:"raw_offer"`
         
    }
    
    // FreeGiftItem ...
    type FreeGiftItem struct {

        
            ItemName string  `json:"item_name"`
            ItemPriceDetails map[string]interface{}  `json:"item_price_details"`
            ItemSlug string  `json:"item_slug"`
            ItemBrandName string  `json:"item_brand_name"`
            ItemID float64  `json:"item_id"`
            ItemImagesURL []string  `json:"item_images_url"`
         
    }
    
    // AppliedFreeArticles ...
    type AppliedFreeArticles struct {

        
            ParentItemIdentifier string  `json:"parent_item_identifier"`
            FreeGiftItemDetails FreeGiftItem  `json:"free_gift_item_details"`
            ArticleID string  `json:"article_id"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // AppliedPromotion ...
    type AppliedPromotion struct {

        
            PromotionGroup string  `json:"promotion_group"`
            PromoID string  `json:"promo_id"`
            BuyRules []BuyRules  `json:"buy_rules"`
            Amount float64  `json:"amount"`
            ArticleQuantity float64  `json:"article_quantity"`
            Ownership Ownership  `json:"ownership"`
            PromotionName string  `json:"promotion_name"`
            PromotionType string  `json:"promotion_type"`
            DiscountRules []DiscountRulesApp  `json:"discount_rules"`
            MrpPromotion bool  `json:"mrp_promotion"`
            OfferText string  `json:"offer_text"`
            AppliedFreeArticles []AppliedFreeArticles  `json:"applied_free_articles"`
         
    }
    
    // BasePrice ...
    type BasePrice struct {

        
            CurrencySymbol string  `json:"currency_symbol"`
            CurrencyCode string  `json:"currency_code"`
            Marked float64  `json:"marked"`
            Effective float64  `json:"effective"`
         
    }
    
    // ArticlePriceInfo ...
    type ArticlePriceInfo struct {

        
            Base BasePrice  `json:"base"`
            Converted BasePrice  `json:"converted"`
         
    }
    
    // ProductArticle ...
    type ProductArticle struct {

        
            UID string  `json:"uid"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            Seller BaseInfo  `json:"seller"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            Size string  `json:"size"`
            Type string  `json:"type"`
            Store BaseInfo  `json:"store"`
            Quantity float64  `json:"quantity"`
            Price ArticlePriceInfo  `json:"price"`
            ProductGroupTags []string  `json:"product_group_tags"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // ProductAvailability ...
    type ProductAvailability struct {

        
            IsValid bool  `json:"is_valid"`
            OtherStoreQuantity float64  `json:"other_store_quantity"`
            OutOfStock bool  `json:"out_of_stock"`
            Deliverable bool  `json:"deliverable"`
            Sizes []string  `json:"sizes"`
         
    }
    
    // PromoMeta ...
    type PromoMeta struct {

        
            Message string  `json:"message"`
         
    }
    
    // ProductPrice ...
    type ProductPrice struct {

        
            CurrencySymbol string  `json:"currency_symbol"`
            Selling float64  `json:"selling"`
            CurrencyCode string  `json:"currency_code"`
            Effective float64  `json:"effective"`
            AddOn float64  `json:"add_on"`
            Marked float64  `json:"marked"`
         
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
    
    // CartProductInfo ...
    type CartProductInfo struct {

        
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            Product CartProduct  `json:"product"`
            PromotionsApplied []AppliedPromotion  `json:"promotions_applied"`
            Article ProductArticle  `json:"article"`
            Availability ProductAvailability  `json:"availability"`
            CouponMessage string  `json:"coupon_message"`
            PromoMeta PromoMeta  `json:"promo_meta"`
            Key string  `json:"key"`
            PricePerUnit ProductPriceInfo  `json:"price_per_unit"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
            Message string  `json:"message"`
            BulkOffer map[string]interface{}  `json:"bulk_offer"`
            Discount string  `json:"discount"`
            Quantity float64  `json:"quantity"`
            Price ProductPriceInfo  `json:"price"`
            IsSet bool  `json:"is_set"`
         
    }
    
    // CouponBreakup ...
    type CouponBreakup struct {

        
            UID string  `json:"uid"`
            MaxDiscountValue float64  `json:"max_discount_value"`
            Code string  `json:"code"`
            Title string  `json:"title"`
            Value float64  `json:"value"`
            Message string  `json:"message"`
            SubTitle string  `json:"sub_title"`
            Type string  `json:"type"`
            CouponValue float64  `json:"coupon_value"`
            Description string  `json:"description"`
            CouponType string  `json:"coupon_type"`
            IsApplied bool  `json:"is_applied"`
            MinimumCartValue float64  `json:"minimum_cart_value"`
         
    }
    
    // DisplayBreakup ...
    type DisplayBreakup struct {

        
            Display string  `json:"display"`
            CurrencySymbol string  `json:"currency_symbol"`
            Value float64  `json:"value"`
            Key string  `json:"key"`
            Message []string  `json:"message"`
            CurrencyCode string  `json:"currency_code"`
         
    }
    
    // RawBreakup ...
    type RawBreakup struct {

        
            YouSaved float64  `json:"you_saved"`
            FyndCash float64  `json:"fynd_cash"`
            DeliveryCharge float64  `json:"delivery_charge"`
            MrpTotal float64  `json:"mrp_total"`
            Total float64  `json:"total"`
            GstCharges float64  `json:"gst_charges"`
            Coupon float64  `json:"coupon"`
            Vog float64  `json:"vog"`
            CodCharge float64  `json:"cod_charge"`
            Discount float64  `json:"discount"`
            ConvenienceFee float64  `json:"convenience_fee"`
            Subtotal float64  `json:"subtotal"`
         
    }
    
    // LoyaltyPoints ...
    type LoyaltyPoints struct {

        
            IsApplied bool  `json:"is_applied"`
            Description string  `json:"description"`
            Applicable float64  `json:"applicable"`
            Total float64  `json:"total"`
         
    }
    
    // CartBreakup ...
    type CartBreakup struct {

        
            Coupon CouponBreakup  `json:"coupon"`
            Display []DisplayBreakup  `json:"display"`
            Raw RawBreakup  `json:"raw"`
            LoyaltyPoints LoyaltyPoints  `json:"loyalty_points"`
         
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
    
    // CartDetailResponse ...
    type CartDetailResponse struct {

        
            PanNo string  `json:"pan_no"`
            Items []CartProductInfo  `json:"items"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            IsValid bool  `json:"is_valid"`
            BuyNow bool  `json:"buy_now"`
            ID string  `json:"id"`
            Comment string  `json:"comment"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            CheckoutMode string  `json:"checkout_mode"`
            LastModified string  `json:"last_modified"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            Message string  `json:"message"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            Currency CartCurrency  `json:"currency"`
            Gstin string  `json:"gstin"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            CouponText string  `json:"coupon_text"`
         
    }
    
    // AddProductCart ...
    type AddProductCart struct {

        
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            ArticleAssignment map[string]interface{}  `json:"article_assignment"`
            Display string  `json:"display"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            ItemSize string  `json:"item_size"`
            ArticleID string  `json:"article_id"`
            SellerID float64  `json:"seller_id"`
            StoreID float64  `json:"store_id"`
            Quantity float64  `json:"quantity"`
            ItemID float64  `json:"item_id"`
            ProductGroupTags []string  `json:"product_group_tags"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Pos bool  `json:"pos"`
         
    }
    
    // AddCartRequest ...
    type AddCartRequest struct {

        
            Items []AddProductCart  `json:"items"`
            NewCart bool  `json:"new_cart"`
         
    }
    
    // AddCartDetailResponse ...
    type AddCartDetailResponse struct {

        
            Partial bool  `json:"partial"`
            Cart CartDetailResponse  `json:"cart"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // UpdateProductCart ...
    type UpdateProductCart struct {

        
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            ItemIndex float64  `json:"item_index"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            ItemSize string  `json:"item_size"`
            ArticleID string  `json:"article_id"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
            Quantity float64  `json:"quantity"`
            ItemID float64  `json:"item_id"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // UpdateCartRequest ...
    type UpdateCartRequest struct {

        
            Items []UpdateProductCart  `json:"items"`
            Operation string  `json:"operation"`
         
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
    
    // PageCoupon ...
    type PageCoupon struct {

        
            Current float64  `json:"current"`
            Total float64  `json:"total"`
            HasPrevious bool  `json:"has_previous"`
            TotalItemCount float64  `json:"total_item_count"`
            HasNext bool  `json:"has_next"`
         
    }
    
    // Coupon ...
    type Coupon struct {

        
            MaxDiscountValue float64  `json:"max_discount_value"`
            ExpiresOn string  `json:"expires_on"`
            Title string  `json:"title"`
            IsApplicable bool  `json:"is_applicable"`
            MinimumCartValue float64  `json:"minimum_cart_value"`
            SubTitle string  `json:"sub_title"`
            CouponValue float64  `json:"coupon_value"`
            CouponType string  `json:"coupon_type"`
            Description string  `json:"description"`
            Message string  `json:"message"`
            IsApplied bool  `json:"is_applied"`
            CouponCode string  `json:"coupon_code"`
         
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
    
    // OfferPrice ...
    type OfferPrice struct {

        
            CurrencySymbol string  `json:"currency_symbol"`
            BulkEffective float64  `json:"bulk_effective"`
            CurrencyCode string  `json:"currency_code"`
            Effective float64  `json:"effective"`
            Marked float64  `json:"marked"`
         
    }
    
    // OfferItem ...
    type OfferItem struct {

        
            Total float64  `json:"total"`
            Type string  `json:"type"`
            Best bool  `json:"best"`
            Quantity float64  `json:"quantity"`
            Price OfferPrice  `json:"price"`
            Margin float64  `json:"margin"`
            AutoApplied bool  `json:"auto_applied"`
         
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

        
            IsActive bool  `json:"is_active"`
            Country string  `json:"country"`
            Tags []string  `json:"tags"`
            City string  `json:"city"`
            CheckoutMode string  `json:"checkout_mode"`
            IsDefaultAddress bool  `json:"is_default_address"`
            UserID string  `json:"user_id"`
            GoogleMapPoint map[string]interface{}  `json:"google_map_point"`
            CountryCode string  `json:"country_code"`
            CountryIsoCode string  `json:"country_iso_code"`
            AddressType string  `json:"address_type"`
            ID string  `json:"id"`
            AreaCodeSlug string  `json:"area_code_slug"`
            CountryPhoneCode string  `json:"country_phone_code"`
            Email string  `json:"email"`
            CreatedByUserID string  `json:"created_by_user_id"`
            Area string  `json:"area"`
            GeoLocation GeoLocation  `json:"geo_location"`
            Meta map[string]interface{}  `json:"meta"`
            AreaCode string  `json:"area_code"`
            Phone string  `json:"phone"`
            Address string  `json:"address"`
            Name string  `json:"name"`
            State string  `json:"state"`
            Landmark string  `json:"landmark"`
         
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
            IsUpdated bool  `json:"is_updated"`
            Success bool  `json:"success"`
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
            ID string  `json:"id"`
            AggregatorName string  `json:"aggregator_name"`
            AddressID string  `json:"address_id"`
            PaymentMode string  `json:"payment_mode"`
            MerchantCode string  `json:"merchant_code"`
         
    }
    
    // CouponValidity ...
    type CouponValidity struct {

        
            Valid bool  `json:"valid"`
            Code string  `json:"code"`
            DisplayMessageEn string  `json:"display_message_en"`
            Title string  `json:"title"`
            Discount float64  `json:"discount"`
         
    }
    
    // PaymentCouponValidate ...
    type PaymentCouponValidate struct {

        
            CouponValidity CouponValidity  `json:"coupon_validity"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // ShipmentResponse ...
    type ShipmentResponse struct {

        
            ShipmentType string  `json:"shipment_type"`
            Items []CartProductInfo  `json:"items"`
            OrderType string  `json:"order_type"`
            FulfillmentType string  `json:"fulfillment_type"`
            FulfillmentID float64  `json:"fulfillment_id"`
            DpOptions map[string]interface{}  `json:"dp_options"`
            Shipments float64  `json:"shipments"`
            DpID string  `json:"dp_id"`
            BoxType string  `json:"box_type"`
            Promise ShipmentPromise  `json:"promise"`
         
    }
    
    // CartShipmentsResponse ...
    type CartShipmentsResponse struct {

        
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            IsValid bool  `json:"is_valid"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            CheckoutMode string  `json:"checkout_mode"`
            LastModified string  `json:"last_modified"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            CartID float64  `json:"cart_id"`
            Shipments []ShipmentResponse  `json:"shipments"`
            Currency CartCurrency  `json:"currency"`
            Error bool  `json:"error"`
            Gstin string  `json:"gstin"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            UID string  `json:"uid"`
            BuyNow bool  `json:"buy_now"`
            ID string  `json:"id"`
            Comment string  `json:"comment"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            Message string  `json:"message"`
            CouponText string  `json:"coupon_text"`
         
    }
    
    // UpdateCartShipmentItem ...
    type UpdateCartShipmentItem struct {

        
            Quantity float64  `json:"quantity"`
            ShipmentType string  `json:"shipment_type"`
            ArticleUID string  `json:"article_uid"`
         
    }
    
    // UpdateCartShipmentRequest ...
    type UpdateCartShipmentRequest struct {

        
            Shipments []UpdateCartShipmentItem  `json:"shipments"`
         
    }
    
    // Files ...
    type Files struct {

        
            Values []string  `json:"values"`
            Key string  `json:"key"`
         
    }
    
    // StaffCheckout ...
    type StaffCheckout struct {

        
            ID string  `json:"_id"`
            User string  `json:"user"`
            EmployeeCode string  `json:"employee_code"`
            FirstName string  `json:"first_name"`
            LastName string  `json:"last_name"`
         
    }
    
    // CartCheckoutCustomMeta ...
    type CartCheckoutCustomMeta struct {

        
            Value string  `json:"value"`
            Key string  `json:"key"`
         
    }
    
    // CartPosCheckoutDetailRequest ...
    type CartPosCheckoutDetailRequest struct {

        
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            PaymentMode string  `json:"payment_mode"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
            OrderType string  `json:"order_type"`
            BillingAddressID string  `json:"billing_address_id"`
            Pos bool  `json:"pos"`
            CallbackURL string  `json:"callback_url"`
            PaymentIdentifier string  `json:"payment_identifier"`
            Meta map[string]interface{}  `json:"meta"`
            AddressID string  `json:"address_id"`
            PaymentParams map[string]interface{}  `json:"payment_params"`
            Files []Files  `json:"files"`
            OrderingStore float64  `json:"ordering_store"`
            Staff StaffCheckout  `json:"staff"`
            PaymentAutoConfirm bool  `json:"payment_auto_confirm"`
            Aggregator string  `json:"aggregator"`
            PickAtStoreUID float64  `json:"pick_at_store_uid"`
            CustomMeta []CartCheckoutCustomMeta  `json:"custom_meta"`
            MerchantCode string  `json:"merchant_code"`
         
    }
    
    // CheckCart ...
    type CheckCart struct {

        
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            CodMessage string  `json:"cod_message"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            CheckoutMode string  `json:"checkout_mode"`
            CartID float64  `json:"cart_id"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            DeliveryChargeOrderValue float64  `json:"delivery_charge_order_value"`
            ErrorMessage string  `json:"error_message"`
            Items []CartProductInfo  `json:"items"`
            BuyNow bool  `json:"buy_now"`
            ID string  `json:"id"`
            Comment string  `json:"comment"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            CouponText string  `json:"coupon_text"`
            IsValid bool  `json:"is_valid"`
            StoreCode string  `json:"store_code"`
            CodCharges float64  `json:"cod_charges"`
            LastModified string  `json:"last_modified"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            OrderID string  `json:"order_id"`
            Currency CartCurrency  `json:"currency"`
            Gstin string  `json:"gstin"`
            UID string  `json:"uid"`
            StoreEmps []map[string]interface{}  `json:"store_emps"`
            CodAvailable bool  `json:"cod_available"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
            UserType string  `json:"user_type"`
            DeliveryCharges float64  `json:"delivery_charges"`
         
    }
    
    // CartCheckoutResponse ...
    type CartCheckoutResponse struct {

        
            Data map[string]interface{}  `json:"data"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
            OrderID string  `json:"order_id"`
            PaymentConfirmURL string  `json:"payment_confirm_url"`
            AppInterceptURL string  `json:"app_intercept_url"`
            Cart CheckCart  `json:"cart"`
            CallbackURL string  `json:"callback_url"`
         
    }
    
    // CartMetaRequest ...
    type CartMetaRequest struct {

        
            Gstin string  `json:"gstin"`
            Comment string  `json:"comment"`
            CheckoutMode string  `json:"checkout_mode"`
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
    
    // CartDeliveryModesResponse ...
    type CartDeliveryModesResponse struct {

        
            AvailableModes []string  `json:"available_modes"`
            PickupStores []float64  `json:"pickup_stores"`
         
    }
    
    // PickupStoreDetail ...
    type PickupStoreDetail struct {

        
            UID float64  `json:"uid"`
            AddressType string  `json:"address_type"`
            Country string  `json:"country"`
            City string  `json:"city"`
            StoreCode string  `json:"store_code"`
            ID float64  `json:"id"`
            Name string  `json:"name"`
            Area string  `json:"area"`
            State string  `json:"state"`
            AreaCodeSlug string  `json:"area_code_slug"`
            AreaCode string  `json:"area_code"`
            Landmark string  `json:"landmark"`
            Email string  `json:"email"`
            Pincode float64  `json:"pincode"`
            Phone string  `json:"phone"`
            Address string  `json:"address"`
         
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

        
            ShareURL string  `json:"share_url"`
            Token string  `json:"token"`
         
    }
    
    // SharedCartDetails ...
    type SharedCartDetails struct {

        
            User map[string]interface{}  `json:"user"`
            Source map[string]interface{}  `json:"source"`
            Meta map[string]interface{}  `json:"meta"`
            Token string  `json:"token"`
            CreatedOn string  `json:"created_on"`
         
    }
    
    // SharedCart ...
    type SharedCart struct {

        
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            CheckoutMode string  `json:"checkout_mode"`
            SharedCartDetails SharedCartDetails  `json:"shared_cart_details"`
            CartID float64  `json:"cart_id"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            Items []CartProductInfo  `json:"items"`
            BuyNow bool  `json:"buy_now"`
            ID string  `json:"id"`
            Comment string  `json:"comment"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            CouponText string  `json:"coupon_text"`
            IsValid bool  `json:"is_valid"`
            LastModified string  `json:"last_modified"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            Currency CartCurrency  `json:"currency"`
            Gstin string  `json:"gstin"`
            UID string  `json:"uid"`
            Message string  `json:"message"`
         
    }
    
    // SharedCartResponse ...
    type SharedCartResponse struct {

        
            Cart SharedCart  `json:"cart"`
            Error string  `json:"error"`
         
    }
    

    
    // PincodeMetaResponse ...
    type PincodeMetaResponse struct {

        
            InternalZoneID float64  `json:"internal_zone_id"`
            Zone string  `json:"zone"`
         
    }
    
    // PincodeErrorSchemaResponse ...
    type PincodeErrorSchemaResponse struct {

        
            Value string  `json:"value"`
            Type string  `json:"type"`
            Message string  `json:"message"`
         
    }
    
    // PincodeParentsResponse ...
    type PincodeParentsResponse struct {

        
            DisplayName string  `json:"display_name"`
            UID string  `json:"uid"`
            SubType string  `json:"sub_type"`
            Name string  `json:"name"`
         
    }
    
    // PincodeDataResponse ...
    type PincodeDataResponse struct {

        
            Meta PincodeMetaResponse  `json:"meta"`
            Name string  `json:"name"`
            Error PincodeErrorSchemaResponse  `json:"error"`
            UID string  `json:"uid"`
            DisplayName string  `json:"display_name"`
            SubType string  `json:"sub_type"`
            Parents []PincodeParentsResponse  `json:"parents"`
         
    }
    
    // PincodeApiResponse ...
    type PincodeApiResponse struct {

        
            Data []PincodeDataResponse  `json:"data"`
            Success bool  `json:"success"`
            Error PincodeErrorSchemaResponse  `json:"error"`
         
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

        
            Journey string  `json:"journey"`
            Identifier string  `json:"identifier"`
            LocationDetails []TATLocationDetailsRequest  `json:"location_details"`
            Source string  `json:"source"`
            ToPincode string  `json:"to_pincode"`
            Action string  `json:"action"`
         
    }
    
    // TATErrorSchemaResponse ...
    type TATErrorSchemaResponse struct {

        
            Value string  `json:"value"`
            Type string  `json:"type"`
            Message string  `json:"message"`
         
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
    
    // TATArticlesResponse ...
    type TATArticlesResponse struct {

        
            ManufacturingTime float64  `json:"manufacturing_time"`
            Promise TATPromiseResponse  `json:"promise"`
            Error TATErrorSchemaResponse  `json:"error"`
            ManufacturingTimeSeconds float64  `json:"_manufacturing_time_seconds"`
            Category TATCategoryRequest  `json:"category"`
            ManufacturingTimeUnit string  `json:"manufacturing_time_unit"`
            IsCodAvailable bool  `json:"is_cod_available"`
         
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
            ToCity string  `json:"to_city"`
            Error TATErrorSchemaResponse  `json:"error"`
            Journey string  `json:"journey"`
            Success bool  `json:"success"`
            RequestUUID string  `json:"request_uuid"`
            Identifier string  `json:"identifier"`
            LocationDetails []TATLocationDetailsResponse  `json:"location_details"`
            IsCodAvailable bool  `json:"is_cod_available"`
            Source string  `json:"source"`
            PaymentMode string  `json:"payment_mode"`
            ToPincode string  `json:"to_pincode"`
            Action string  `json:"action"`
         
    }
    
    // GetZoneFromPincodeViewRequest ...
    type GetZoneFromPincodeViewRequest struct {

        
            Country string  `json:"country"`
            Pincode string  `json:"pincode"`
         
    }
    
    // GetZoneFromPincodeViewResponse ...
    type GetZoneFromPincodeViewResponse struct {

        
            ServiceabilityType string  `json:"serviceability_type"`
            Zones []string  `json:"zones"`
         
    }
    
    // ReAssignStoreRequest ...
    type ReAssignStoreRequest struct {

        
            Configuration map[string]interface{}  `json:"configuration"`
            Articles []map[string]interface{}  `json:"articles"`
            ToPincode string  `json:"to_pincode"`
            Identifier string  `json:"identifier"`
         
    }
    
    // ReAssignStoreResponse ...
    type ReAssignStoreResponse struct {

        
            Store map[string]interface{}  `json:"store"`
            Article map[string]interface{}  `json:"article"`
            PageSize float64  `json:"page_size"`
            ToPincode string  `json:"to_pincode"`
            Error map[string]interface{}  `json:"error"`
            CustomerDetails map[string]interface{}  `json:"customer_details"`
            AssignedStores []map[string]interface{}  `json:"assigned_stores"`
            Success bool  `json:"success"`
            PageNo float64  `json:"page_no"`
            Articles []map[string]interface{}  `json:"articles"`
            Company map[string]interface{}  `json:"company"`
         
    }
    
