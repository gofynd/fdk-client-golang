package application



    
    // Meta ...
    type Meta struct {

        
            Source string  `json:"source"`
         
    }
    
    // Media ...
    type Media struct {

        
            Type string  `json:"type"`
            URL string  `json:"url"`
            Alt string  `json:"alt"`
            Meta Meta  `json:"meta"`
         
    }
    
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
    
    // ProductBrand ...
    type ProductBrand struct {

        
            Description string  `json:"description"`
            Name string  `json:"name"`
            Logo Media  `json:"logo"`
            UID float64  `json:"uid"`
            Action ProductListingAction  `json:"action"`
         
    }
    
    // ProductCategoryMap ...
    type ProductCategoryMap struct {

        
            L2 ProductBrand  `json:"l2"`
            L3 ProductBrand  `json:"l3"`
            L1 ProductBrand  `json:"l1"`
         
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

        
            Marked Price  `json:"marked"`
            Effective Price  `json:"effective"`
         
    }
    
    // ProductDetailAttribute ...
    type ProductDetailAttribute struct {

        
            Type string  `json:"type"`
            Key string  `json:"key"`
            Value string  `json:"value"`
         
    }
    
    // ProductDetailGroupedAttribute ...
    type ProductDetailGroupedAttribute struct {

        
            Details []ProductDetailAttribute  `json:"details"`
            Title string  `json:"title"`
         
    }
    
    // CustomMetaFields ...
    type CustomMetaFields struct {

        
            Key string  `json:"key"`
            Value string  `json:"value"`
         
    }
    
    // ApplicationItemSEO ...
    type ApplicationItemSEO struct {

        
            Title string  `json:"title"`
            Description string  `json:"description"`
         
    }
    
    // NetQuantity ...
    type NetQuantity struct {

        
            Unit string  `json:"unit"`
            Value float64  `json:"value"`
         
    }
    
    // ApplicationItemMOQ ...
    type ApplicationItemMOQ struct {

        
            IncrementUnit float64  `json:"increment_unit"`
            Maximum float64  `json:"maximum"`
            Minimum float64  `json:"minimum"`
         
    }
    
    // ProductDetail ...
    type ProductDetail struct {

        
            Description string  `json:"description"`
            IsDependent bool  `json:"is_dependent"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ItemType string  `json:"item_type"`
            ItemCode string  `json:"item_code"`
            CategoryMap ProductCategoryMap  `json:"category_map"`
            Color string  `json:"color"`
            Price ProductListingPrice  `json:"price"`
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
            Highlights []string  `json:"highlights"`
            TeaserTag string  `json:"teaser_tag"`
            Name string  `json:"name"`
            Medias []Media  `json:"medias"`
            CustomMeta []CustomMetaFields  `json:"_custom_meta"`
            Type string  `json:"type"`
            UID float64  `json:"uid"`
            Seo ApplicationItemSEO  `json:"seo"`
            Action ProductListingAction  `json:"action"`
            Tryouts []string  `json:"tryouts"`
            Slug string  `json:"slug"`
            NetQuantity NetQuantity  `json:"net_quantity"`
            RatingCount float64  `json:"rating_count"`
            Similars []string  `json:"similars"`
            ProductOnlineDate string  `json:"product_online_date"`
            Brand ProductBrand  `json:"brand"`
            Tags []string  `json:"tags"`
            HasVariant bool  `json:"has_variant"`
            ImageNature string  `json:"image_nature"`
            Moq ApplicationItemMOQ  `json:"moq"`
            Categories []ProductBrand  `json:"categories"`
            ShortDescription string  `json:"short_description"`
            Attributes map[string]interface{}  `json:"attributes"`
            Discount string  `json:"discount"`
            Rating float64  `json:"rating"`
         
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

        
            Convertable bool  `json:"convertable"`
            Value string  `json:"value"`
         
    }
    
    // ColumnHeaders ...
    type ColumnHeaders struct {

        
            Col3 ColumnHeader  `json:"col_3"`
            Col2 ColumnHeader  `json:"col_2"`
            Col5 ColumnHeader  `json:"col_5"`
            Col1 ColumnHeader  `json:"col_1"`
            Col4 ColumnHeader  `json:"col_4"`
            Col6 ColumnHeader  `json:"col_6"`
         
    }
    
    // SizeChartValues ...
    type SizeChartValues struct {

        
            Col3 string  `json:"col_3"`
            Col2 string  `json:"col_2"`
            Col5 string  `json:"col_5"`
            Col1 string  `json:"col_1"`
            Col4 string  `json:"col_4"`
            Col6 string  `json:"col_6"`
         
    }
    
    // SizeChart ...
    type SizeChart struct {

        
            Description string  `json:"description"`
            Title string  `json:"title"`
            Headers ColumnHeaders  `json:"headers"`
            Sizes []SizeChartValues  `json:"sizes"`
            Unit string  `json:"unit"`
            Image string  `json:"image"`
            SizeTip string  `json:"size_tip"`
         
    }
    
    // Dimension ...
    type Dimension struct {

        
            IsDefault bool  `json:"is_default"`
            Length float64  `json:"length"`
            Width float64  `json:"width"`
            Unit string  `json:"unit"`
            Height float64  `json:"height"`
         
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
            SellerIdentifiers []string  `json:"seller_identifiers"`
            Display string  `json:"display"`
            Quantity float64  `json:"quantity"`
            Dimension Dimension  `json:"dimension"`
            IsAvailable bool  `json:"is_available"`
            Weight Weight  `json:"weight"`
         
    }
    
    // ProductSizes ...
    type ProductSizes struct {

        
            MultiSize bool  `json:"multi_size"`
            Sellable bool  `json:"sellable"`
            Price ProductListingPrice  `json:"price"`
            Stores ProductSizeStores  `json:"stores"`
            Discount string  `json:"discount"`
            SizeChart SizeChart  `json:"size_chart"`
            Sizes []ProductSize  `json:"sizes"`
         
    }
    
    // AttributeDetail ...
    type AttributeDetail struct {

        
            Logo string  `json:"logo"`
            Key string  `json:"key"`
            Display string  `json:"display"`
            Description string  `json:"description"`
         
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
            Title string  `json:"title"`
            Items []ProductDetail  `json:"items"`
         
    }
    
    // ProductFrequentlyComparedSimilarResponse ...
    type ProductFrequentlyComparedSimilarResponse struct {

        
            Similars ProductCompareResponse  `json:"similars"`
         
    }
    
    // ProductVariantItemResponse ...
    type ProductVariantItemResponse struct {

        
            Slug string  `json:"slug"`
            Color string  `json:"color"`
            Value string  `json:"value"`
            ColorName string  `json:"color_name"`
            Name string  `json:"name"`
            Medias []Media  `json:"medias"`
            CustomMeta []CustomMetaFields  `json:"_custom_meta"`
            UID float64  `json:"uid"`
            Action ProductListingAction  `json:"action"`
            IsAvailable bool  `json:"is_available"`
         
    }
    
    // ProductVariantResponse ...
    type ProductVariantResponse struct {

        
            DisplayType string  `json:"display_type"`
            Key string  `json:"key"`
            Items []ProductVariantItemResponse  `json:"items"`
            Header string  `json:"header"`
         
    }
    
    // ProductVariantsResponse ...
    type ProductVariantsResponse struct {

        
            Variants []ProductVariantResponse  `json:"variants"`
         
    }
    
    // StoreDetail ...
    type StoreDetail struct {

        
            City string  `json:"city"`
            Code string  `json:"code"`
            ID float64  `json:"id"`
            Name string  `json:"name"`
         
    }
    
    // Seller ...
    type Seller struct {

        
            Count float64  `json:"count"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // ProductStockPrice ...
    type ProductStockPrice struct {

        
            Marked float64  `json:"marked"`
            Currency string  `json:"currency"`
            Effective float64  `json:"effective"`
         
    }
    
    // CompanyDetail ...
    type CompanyDetail struct {

        
            ID float64  `json:"id"`
            Name string  `json:"name"`
         
    }
    
    // ProductStockStatusItem ...
    type ProductStockStatusItem struct {

        
            Store StoreDetail  `json:"store"`
            Seller Seller  `json:"seller"`
            Identifier map[string]interface{}  `json:"identifier"`
            Price ProductStockPrice  `json:"price"`
            ItemID float64  `json:"item_id"`
            Quantity float64  `json:"quantity"`
            Company CompanyDetail  `json:"company"`
            UID string  `json:"uid"`
            Size string  `json:"size"`
         
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

        
            Value string  `json:"value"`
            IsSelected bool  `json:"is_selected"`
            Name string  `json:"name"`
         
    }
    
    // ProductFiltersKey ...
    type ProductFiltersKey struct {

        
            Logo string  `json:"logo"`
            Display string  `json:"display"`
            Kind string  `json:"kind"`
            Name string  `json:"name"`
         
    }
    
    // ProductFiltersValue ...
    type ProductFiltersValue struct {

        
            Max float64  `json:"max"`
            DisplayFormat string  `json:"display_format"`
            Value string  `json:"value"`
            IsSelected bool  `json:"is_selected"`
            Display string  `json:"display"`
            Count float64  `json:"count"`
            Min float64  `json:"min"`
            SelectedMax float64  `json:"selected_max"`
            CurrencySymbol string  `json:"currency_symbol"`
            SelectedMin float64  `json:"selected_min"`
            QueryFormat string  `json:"query_format"`
            CurrencyCode string  `json:"currency_code"`
         
    }
    
    // ProductFilters ...
    type ProductFilters struct {

        
            Key ProductFiltersKey  `json:"key"`
            Values []ProductFiltersValue  `json:"values"`
         
    }
    
    // ProductVariantListingResponse ...
    type ProductVariantListingResponse struct {

        
            DisplayType string  `json:"display_type"`
            Key string  `json:"key"`
            Items []ProductVariantItemResponse  `json:"items"`
            Total float64  `json:"total"`
            Header string  `json:"header"`
         
    }
    
    // ProductListingDetail ...
    type ProductListingDetail struct {

        
            Description string  `json:"description"`
            IsDependent bool  `json:"is_dependent"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ItemType string  `json:"item_type"`
            ItemCode string  `json:"item_code"`
            CategoryMap ProductCategoryMap  `json:"category_map"`
            Color string  `json:"color"`
            Price ProductListingPrice  `json:"price"`
            Sellable bool  `json:"sellable"`
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
            Highlights []string  `json:"highlights"`
            TeaserTag string  `json:"teaser_tag"`
            Name string  `json:"name"`
            Medias []Media  `json:"medias"`
            CustomMeta []CustomMetaFields  `json:"_custom_meta"`
            Type string  `json:"type"`
            UID float64  `json:"uid"`
            Seo ApplicationItemSEO  `json:"seo"`
            Action ProductListingAction  `json:"action"`
            Variants []ProductVariantListingResponse  `json:"variants"`
            Tryouts []string  `json:"tryouts"`
            Slug string  `json:"slug"`
            NetQuantity NetQuantity  `json:"net_quantity"`
            RatingCount float64  `json:"rating_count"`
            Similars []string  `json:"similars"`
            ProductOnlineDate string  `json:"product_online_date"`
            Brand ProductBrand  `json:"brand"`
            Tags []string  `json:"tags"`
            HasVariant bool  `json:"has_variant"`
            ImageNature string  `json:"image_nature"`
            Sizes []string  `json:"sizes"`
            Moq ApplicationItemMOQ  `json:"moq"`
            Identifiers []string  `json:"identifiers"`
            Categories []ProductBrand  `json:"categories"`
            ShortDescription string  `json:"short_description"`
            Attributes map[string]interface{}  `json:"attributes"`
            Discount string  `json:"discount"`
            Rating float64  `json:"rating"`
         
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

        
            Slug string  `json:"slug"`
            Description string  `json:"description"`
            Name string  `json:"name"`
            Logo Media  `json:"logo"`
            Banners ImageUrls  `json:"banners"`
            UID float64  `json:"uid"`
            Departments []string  `json:"departments"`
            Discount string  `json:"discount"`
            Action ProductListingAction  `json:"action"`
         
    }
    
    // BrandListingResponse ...
    type BrandListingResponse struct {

        
            Items []BrandItem  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // BrandDetailResponse ...
    type BrandDetailResponse struct {

        
            Description string  `json:"description"`
            Name string  `json:"name"`
            Logo Media  `json:"logo"`
            Banners ImageUrls  `json:"banners"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            UID float64  `json:"uid"`
         
    }
    
    // CategoryBanner ...
    type CategoryBanner struct {

        
            Portrait Media  `json:"portrait"`
            Landscape Media  `json:"landscape"`
         
    }
    
    // ThirdLevelChild ...
    type ThirdLevelChild struct {

        
            Slug string  `json:"slug"`
            Name string  `json:"name"`
            Banners ImageUrls  `json:"banners"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            UID float64  `json:"uid"`
            Action ProductListingAction  `json:"action"`
            Childs []map[string]interface{}  `json:"childs"`
         
    }
    
    // SecondLevelChild ...
    type SecondLevelChild struct {

        
            Slug string  `json:"slug"`
            Name string  `json:"name"`
            Banners ImageUrls  `json:"banners"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            UID float64  `json:"uid"`
            Action ProductListingAction  `json:"action"`
            Childs []ThirdLevelChild  `json:"childs"`
         
    }
    
    // Child ...
    type Child struct {

        
            Slug string  `json:"slug"`
            Name string  `json:"name"`
            Banners ImageUrls  `json:"banners"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            UID float64  `json:"uid"`
            Action ProductListingAction  `json:"action"`
            Childs []SecondLevelChild  `json:"childs"`
         
    }
    
    // CategoryItems ...
    type CategoryItems struct {

        
            Slug string  `json:"slug"`
            Name string  `json:"name"`
            Banners CategoryBanner  `json:"banners"`
            UID float64  `json:"uid"`
            Action ProductListingAction  `json:"action"`
            Childs []Child  `json:"childs"`
         
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

        
            Name string  `json:"name"`
            Logo Media  `json:"logo"`
            Banners ImageUrls  `json:"banners"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            UID float64  `json:"uid"`
         
    }
    
    // HomeListingResponse ...
    type HomeListingResponse struct {

        
            Message string  `json:"message"`
            Items []ProductListingDetail  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // Department ...
    type Department struct {

        
            Slug string  `json:"slug"`
            PriorityOrder float64  `json:"priority_order"`
            Name string  `json:"name"`
            Logo Media  `json:"logo"`
            UID float64  `json:"uid"`
         
    }
    
    // DepartmentResponse ...
    type DepartmentResponse struct {

        
            Items []Department  `json:"items"`
         
    }
    
    // AutocompleteItem ...
    type AutocompleteItem struct {

        
            Display string  `json:"display"`
            Logo Media  `json:"logo"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Type string  `json:"type"`
            Action ProductListingAction  `json:"action"`
         
    }
    
    // AutoCompleteResponse ...
    type AutoCompleteResponse struct {

        
            Items []AutocompleteItem  `json:"items"`
         
    }
    
    // CollectionListingFilterType ...
    type CollectionListingFilterType struct {

        
            IsSelected bool  `json:"is_selected"`
            Display string  `json:"display"`
            Name string  `json:"name"`
         
    }
    
    // CollectionListingFilterTag ...
    type CollectionListingFilterTag struct {

        
            IsSelected bool  `json:"is_selected"`
            Display string  `json:"display"`
            Name string  `json:"name"`
         
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
            AllowFacets bool  `json:"allow_facets"`
            Description string  `json:"description"`
            Banners ImageUrls  `json:"banners"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            IsActive bool  `json:"is_active"`
            AppID string  `json:"app_id"`
            Cron map[string]interface{}  `json:"cron"`
            SortOn string  `json:"sort_on"`
            Name string  `json:"name"`
            Type string  `json:"type"`
            UID string  `json:"uid"`
            Action ProductListingAction  `json:"action"`
            Badge map[string]interface{}  `json:"badge"`
            Tag []string  `json:"tag"`
            Slug string  `json:"slug"`
            Meta map[string]interface{}  `json:"meta"`
            Logo Media  `json:"logo"`
            AllowSort bool  `json:"allow_sort"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Schedule map[string]interface{}  `json:"_schedule"`
            Query []CollectionQuery  `json:"query"`
         
    }
    
    // GetCollectionListingResponse ...
    type GetCollectionListingResponse struct {

        
            Filters CollectionListingFilter  `json:"filters"`
            Items []GetCollectionDetailNest  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // CollectionDetailResponse ...
    type CollectionDetailResponse struct {

        
            Priority float64  `json:"priority"`
            AllowFacets bool  `json:"allow_facets"`
            Description string  `json:"description"`
            Banners ImageUrls  `json:"banners"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            IsActive bool  `json:"is_active"`
            AppID string  `json:"app_id"`
            Cron map[string]interface{}  `json:"cron"`
            SortOn string  `json:"sort_on"`
            Name string  `json:"name"`
            Type string  `json:"type"`
            Badge map[string]interface{}  `json:"badge"`
            Tag []string  `json:"tag"`
            Slug string  `json:"slug"`
            Meta map[string]interface{}  `json:"meta"`
            Logo Media  `json:"logo"`
            AllowSort bool  `json:"allow_sort"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Schedule map[string]interface{}  `json:"_schedule"`
            Query []CollectionQuery  `json:"query"`
         
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

        
            Type string  `json:"type"`
            Coordinates []float64  `json:"coordinates"`
         
    }
    
    // Store ...
    type Store struct {

        
            State string  `json:"state"`
            StoreEmail string  `json:"store_email"`
            Address string  `json:"address"`
            Name string  `json:"name"`
            LatLong LatLong  `json:"lat_long"`
            UID float64  `json:"uid"`
            StoreCode string  `json:"store_code"`
            Country string  `json:"country"`
            City string  `json:"city"`
            Pincode float64  `json:"pincode"`
         
    }
    
    // StoreListingResponse ...
    type StoreListingResponse struct {

        
            Items []Store  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // StoreDepartments ...
    type StoreDepartments struct {

        
            Slug string  `json:"slug"`
            PriorityOrder float64  `json:"priority_order"`
            Name string  `json:"name"`
            Logo string  `json:"logo"`
            UID float64  `json:"uid"`
         
    }
    
    // SellerPhoneNumber ...
    type SellerPhoneNumber struct {

        
            Number string  `json:"number"`
            CountryCode float64  `json:"country_code"`
         
    }
    
    // StoreManagerSerializer ...
    type StoreManagerSerializer struct {

        
            Email string  `json:"email"`
            MobileNo SellerPhoneNumber  `json:"mobile_no"`
            Name string  `json:"name"`
         
    }
    
    // StoreAddressSerializer ...
    type StoreAddressSerializer struct {

        
            State string  `json:"state"`
            Latitude float64  `json:"latitude"`
            Country string  `json:"country"`
            Address1 string  `json:"address1"`
            Longitude float64  `json:"longitude"`
            Landmark string  `json:"landmark"`
            City string  `json:"city"`
            Pincode float64  `json:"pincode"`
            Address2 string  `json:"address2"`
         
    }
    
    // CompanyStore ...
    type CompanyStore struct {

        
            CompanyType string  `json:"company_type"`
            BusinessType string  `json:"business_type"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // AppStore ...
    type AppStore struct {

        
            Manager StoreManagerSerializer  `json:"manager"`
            Address StoreAddressSerializer  `json:"address"`
            Name string  `json:"name"`
            Company CompanyStore  `json:"company"`
            UID float64  `json:"uid"`
            Departments []StoreDepartments  `json:"departments"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
         
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

        
            Closing Time  `json:"closing"`
            Open bool  `json:"open"`
            Opening Time  `json:"opening"`
            Weekday string  `json:"weekday"`
         
    }
    
    // StoreDetails ...
    type StoreDetails struct {

        
            Manager StoreManagerSerializer  `json:"manager"`
            Timing []StoreTiming  `json:"timing"`
            Address StoreAddressSerializer  `json:"address"`
            Name string  `json:"name"`
            Company CompanyStore  `json:"company"`
            UID float64  `json:"uid"`
            Departments []StoreDepartments  `json:"departments"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
         
    }
    
    // UserDetail ...
    type UserDetail struct {

        
            Username string  `json:"username"`
            Contact string  `json:"contact"`
            UserID string  `json:"user_id"`
            SuperUser bool  `json:"super_user"`
         
    }
    
    // ProductGroupPrice ...
    type ProductGroupPrice struct {

        
            MaxEffective float64  `json:"max_effective"`
            MinMarked float64  `json:"min_marked"`
            Currency string  `json:"currency"`
            MaxMarked float64  `json:"max_marked"`
            MinEffective float64  `json:"min_effective"`
         
    }
    
    // ProductDetails ...
    type ProductDetails struct {

        
            Description string  `json:"description"`
            CountryOfOrigin string  `json:"country_of_origin"`
            HsnCode float64  `json:"hsn_code"`
            Identifier map[string]interface{}  `json:"identifier"`
            Images []string  `json:"images"`
            BrandUID float64  `json:"brand_uid"`
            OutOfStock bool  `json:"out_of_stock"`
            ItemCode string  `json:"item_code"`
            Media []map[string]interface{}  `json:"media"`
            GroupedAttributes map[string]interface{}  `json:"grouped_attributes"`
            Highlights []string  `json:"highlights"`
            Name string  `json:"name"`
            IsSet bool  `json:"is_set"`
            Slug string  `json:"slug"`
            RatingCount float64  `json:"rating_count"`
            HasVariant bool  `json:"has_variant"`
            ImageNature string  `json:"image_nature"`
            TemplateTag string  `json:"template_tag"`
            ShortDescription string  `json:"short_description"`
            Attributes map[string]interface{}  `json:"attributes"`
            Rating float64  `json:"rating"`
         
    }
    
    // Size ...
    type Size struct {

        
            Quantity float64  `json:"quantity"`
            IsAvailable bool  `json:"is_available"`
            Display string  `json:"display"`
            Value string  `json:"value"`
         
    }
    
    // ProductInGroup ...
    type ProductInGroup struct {

        
            AutoSelect bool  `json:"auto_select"`
            Price ProductGroupPrice  `json:"price"`
            ProductUID float64  `json:"product_uid"`
            MaxQuantity float64  `json:"max_quantity"`
            AutoAddToCart bool  `json:"auto_add_to_cart"`
            AllowRemove bool  `json:"allow_remove"`
            MinQuantity float64  `json:"min_quantity"`
            ProductDetails ProductDetails  `json:"product_details"`
            Sizes []Size  `json:"sizes"`
         
    }
    
    // ProductGroupingModel ...
    type ProductGroupingModel struct {

        
            Slug string  `json:"slug"`
            VerifiedBy UserDetail  `json:"verified_by"`
            CreatedBy UserDetail  `json:"created_by"`
            CompanyID float64  `json:"company_id"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            VerifiedOn string  `json:"verified_on"`
            Name string  `json:"name"`
            Meta map[string]interface{}  `json:"meta"`
            Logo string  `json:"logo"`
            Choice string  `json:"choice"`
            IsActive bool  `json:"is_active"`
            ModifiedOn string  `json:"modified_on"`
            Products []ProductInGroup  `json:"products"`
            ID interface{}  `json:"_id"`
            CreatedOn string  `json:"created_on"`
            PageVisibility []string  `json:"page_visibility"`
            ModifiedBy UserDetail  `json:"modified_by"`
         
    }
    
    // ProductBundle ...
    type ProductBundle struct {

        
            Items []ProductGroupingModel  `json:"items"`
         
    }
    
    // ReturnConfigSchemaV3 ...
    type ReturnConfigSchemaV3 struct {

        
            Returnable bool  `json:"returnable"`
            Time float64  `json:"time"`
            Unit string  `json:"unit"`
         
    }
    
    // SellerV3 ...
    type SellerV3 struct {

        
            Count float64  `json:"count"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // ArticleAssignmentV3 ...
    type ArticleAssignmentV3 struct {

        
            Level string  `json:"level"`
            Strategy string  `json:"strategy"`
         
    }
    
    // ProductStockPriceV3 ...
    type ProductStockPriceV3 struct {

        
            Marked float64  `json:"marked"`
            Currency string  `json:"currency"`
            Effective float64  `json:"effective"`
         
    }
    
    // StrategyWiseListingSchemaV3 ...
    type StrategyWiseListingSchemaV3 struct {

        
            Quantity float64  `json:"quantity"`
            Tat float64  `json:"tat"`
            Pincode float64  `json:"pincode"`
            Distance float64  `json:"distance"`
         
    }
    
    // DetailsSchemaV3 ...
    type DetailsSchemaV3 struct {

        
            Type string  `json:"type"`
            Key string  `json:"key"`
            Value string  `json:"value"`
         
    }
    
    // SellerGroupAttributes ...
    type SellerGroupAttributes struct {

        
            Details []DetailsSchemaV3  `json:"details"`
            Title string  `json:"title"`
         
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

        
            Price float64  `json:"price"`
            Unit string  `json:"unit"`
            CurrencySymbol string  `json:"currency_symbol"`
            CurrencyCode string  `json:"currency_code"`
         
    }
    
    // MarketPlaceSttributesSchemaV3 ...
    type MarketPlaceSttributesSchemaV3 struct {

        
            Details []DetailsSchemaV3  `json:"details"`
            Title string  `json:"title"`
         
    }
    
    // StoreV3 ...
    type StoreV3 struct {

        
            Count float64  `json:"count"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // ProductSizePriceResponseV3 ...
    type ProductSizePriceResponseV3 struct {

        
            ReturnConfig ReturnConfigSchemaV3  `json:"return_config"`
            Quantity float64  `json:"quantity"`
            IsGift bool  `json:"is_gift"`
            ItemType string  `json:"item_type"`
            Seller SellerV3  `json:"seller"`
            ArticleAssignment ArticleAssignmentV3  `json:"article_assignment"`
            Price ProductStockPriceV3  `json:"price"`
            StrategyWiseListing []StrategyWiseListingSchemaV3  `json:"strategy_wise_listing"`
            GroupedAttributes []SellerGroupAttributes  `json:"grouped_attributes"`
            Set ProductSetV3  `json:"set"`
            PricePerUnit ProductStockUnitPriceV3  `json:"price_per_unit"`
            MarketplaceAttributes []MarketPlaceSttributesSchemaV3  `json:"marketplace_attributes"`
            SellerCount float64  `json:"seller_count"`
            LongLat []float64  `json:"long_lat"`
            Store StoreV3  `json:"store"`
            ArticleID string  `json:"article_id"`
            IsCod bool  `json:"is_cod"`
            SpecialBadge string  `json:"special_badge"`
            PricePerPiece ProductStockPriceV3  `json:"price_per_piece"`
            Discount string  `json:"discount"`
            Pincode float64  `json:"pincode"`
         
    }
    
    // ProductSizeSellerFilterSchemaV3 ...
    type ProductSizeSellerFilterSchemaV3 struct {

        
            IsSelected bool  `json:"is_selected"`
            Value string  `json:"value"`
            Name string  `json:"name"`
         
    }
    
    // ProductSizeSellersResponseV3 ...
    type ProductSizeSellersResponseV3 struct {

        
            SortOn []ProductSizeSellerFilterSchemaV3  `json:"sort_on"`
            Items []ProductSizePriceResponseV3  `json:"items"`
            Page Page  `json:"page"`
         
    }
    

    
    // PaymentSelectionLock ...
    type PaymentSelectionLock struct {

        
            DefaultOptions string  `json:"default_options"`
            PaymentIdentifier string  `json:"payment_identifier"`
            Enabled bool  `json:"enabled"`
         
    }
    
    // CartCurrencySchema ...
    type CartCurrencySchema struct {

        
            Symbol string  `json:"symbol"`
            Code string  `json:"code"`
         
    }
    
    // CouponDetails ...
    type CouponDetails struct {

        
            Code string  `json:"code"`
            DiscountTotalQuantity float64  `json:"discount_total_quantity"`
            DiscountSingleQuantity float64  `json:"discount_single_quantity"`
         
    }
    
    // PromoMeta ...
    type PromoMeta struct {

        
            Message string  `json:"message"`
         
    }
    
    // BaseInfo ...
    type BaseInfo struct {

        
            Name string  `json:"name"`
            UID float64  `json:"uid"`
         
    }
    
    // BasePriceSchema ...
    type BasePriceSchema struct {

        
            CurrencyCode string  `json:"currency_code"`
            CurrencySymbol string  `json:"currency_symbol"`
            Marked float64  `json:"marked"`
            Effective float64  `json:"effective"`
         
    }
    
    // ArticlePriceInfo ...
    type ArticlePriceInfo struct {

        
            Converted BasePriceSchema  `json:"converted"`
            Base BasePriceSchema  `json:"base"`
         
    }
    
    // ProductArticle ...
    type ProductArticle struct {

        
            Size string  `json:"size"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Type string  `json:"type"`
            Quantity float64  `json:"quantity"`
            SellerIdentifier string  `json:"seller_identifier"`
            Identifier map[string]interface{}  `json:"identifier"`
            GiftCard map[string]interface{}  `json:"gift_card"`
            IsGiftVisible bool  `json:"is_gift_visible"`
            ProductGroupTags []string  `json:"product_group_tags"`
            Seller BaseInfo  `json:"seller"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            UID string  `json:"uid"`
            Price ArticlePriceInfo  `json:"price"`
            Store BaseInfo  `json:"store"`
            CartItemMeta map[string]interface{}  `json:"cart_item_meta"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
         
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

        
            Query ActionQuery  `json:"query"`
            URL string  `json:"url"`
            Type string  `json:"type"`
         
    }
    
    // ProductImage ...
    type ProductImage struct {

        
            SecureURL string  `json:"secure_url"`
            URL string  `json:"url"`
            AspectRatio string  `json:"aspect_ratio"`
         
    }
    
    // CategoryInfo ...
    type CategoryInfo struct {

        
            Name string  `json:"name"`
            UID float64  `json:"uid"`
         
    }
    
    // Tags ...
    type Tags struct {

        
            Tags map[string]interface{}  `json:"tags"`
         
    }
    
    // CartProduct ...
    type CartProduct struct {

        
            ItemCode string  `json:"item_code"`
            NetQuantity NetQuantity  `json:"net_quantity"`
            Brand BaseInfo  `json:"brand"`
            Type string  `json:"type"`
            Action ProductAction  `json:"action"`
            Slug string  `json:"slug"`
            Images []ProductImage  `json:"images"`
            Categories []CategoryInfo  `json:"categories"`
            UID float64  `json:"uid"`
            TeaserTag Tags  `json:"teaser_tag"`
            Name string  `json:"name"`
            Tags []string  `json:"tags"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // ProductAvailabilitySize ...
    type ProductAvailabilitySize struct {

        
            IsAvailable bool  `json:"is_available"`
            Value string  `json:"value"`
            Display string  `json:"display"`
         
    }
    
    // ProductAvailability ...
    type ProductAvailability struct {

        
            AvailableSizes []ProductAvailabilitySize  `json:"available_sizes"`
            OutOfStock bool  `json:"out_of_stock"`
            Sizes []string  `json:"sizes"`
            OtherStoreQuantity float64  `json:"other_store_quantity"`
            Deliverable bool  `json:"deliverable"`
            IsValid bool  `json:"is_valid"`
         
    }
    
    // CartProductIdentifer ...
    type CartProductIdentifer struct {

        
            Identifier string  `json:"identifier"`
         
    }
    
    // ProductPrice ...
    type ProductPrice struct {

        
            CurrencyCode string  `json:"currency_code"`
            AddOn float64  `json:"add_on"`
            CurrencySymbol string  `json:"currency_symbol"`
            Effective float64  `json:"effective"`
            Selling float64  `json:"selling"`
            Marked float64  `json:"marked"`
         
    }
    
    // ProductPriceInfo ...
    type ProductPriceInfo struct {

        
            Converted ProductPrice  `json:"converted"`
            Base ProductPrice  `json:"base"`
         
    }
    
    // Ownership ...
    type Ownership struct {

        
            PayableBy string  `json:"payable_by"`
            PayableCategory string  `json:"payable_category"`
         
    }
    
    // BuyRulesSchema ...
    type BuyRulesSchema struct {

        
            ItemCriteria map[string]interface{}  `json:"item_criteria"`
            CartConditions map[string]interface{}  `json:"cart_conditions"`
         
    }
    
    // DiscountRulesAppSchema ...
    type DiscountRulesAppSchema struct {

        
            Offer map[string]interface{}  `json:"offer"`
            RawOffer map[string]interface{}  `json:"raw_offer"`
            ItemCriteria map[string]interface{}  `json:"item_criteria"`
            MatchedBuyRules []string  `json:"matched_buy_rules"`
         
    }
    
    // FreeGiftItemSchema ...
    type FreeGiftItemSchema struct {

        
            ItemImagesURL []string  `json:"item_images_url"`
            ItemSlug string  `json:"item_slug"`
            ItemPriceDetails map[string]interface{}  `json:"item_price_details"`
            ItemName string  `json:"item_name"`
            ItemID float64  `json:"item_id"`
            ItemBrandName string  `json:"item_brand_name"`
         
    }
    
    // AppliedFreeArticlesSchema ...
    type AppliedFreeArticlesSchema struct {

        
            ParentItemIdentifier string  `json:"parent_item_identifier"`
            ArticleID string  `json:"article_id"`
            Quantity float64  `json:"quantity"`
            FreeGiftItemDetails FreeGiftItemSchema  `json:"free_gift_item_details"`
         
    }
    
    // AppliedPromotion ...
    type AppliedPromotion struct {

        
            PromotionGroup string  `json:"promotion_group"`
            Amount float64  `json:"amount"`
            Ownership Ownership  `json:"ownership"`
            BuyRules []BuyRulesSchema  `json:"buy_rules"`
            PromotionName string  `json:"promotion_name"`
            DiscountRules []DiscountRulesAppSchema  `json:"discount_rules"`
            MrpPromotion bool  `json:"mrp_promotion"`
            PromoID string  `json:"promo_id"`
            AppliedFreeArticles []AppliedFreeArticlesSchema  `json:"applied_free_articles"`
            ArticleQuantity float64  `json:"article_quantity"`
            OfferText string  `json:"offer_text"`
            PromotionType string  `json:"promotion_type"`
         
    }
    
    // PromiseTimestamp ...
    type PromiseTimestamp struct {

        
            Min float64  `json:"min"`
            Max float64  `json:"max"`
         
    }
    
    // PromiseFormatted ...
    type PromiseFormatted struct {

        
            Min string  `json:"min"`
            Max string  `json:"max"`
         
    }
    
    // ShipmentPromise ...
    type ShipmentPromise struct {

        
            Timestamp PromiseTimestamp  `json:"timestamp"`
            Formatted PromiseFormatted  `json:"formatted"`
         
    }
    
    // CartProductInfo ...
    type CartProductInfo struct {

        
            Coupon CouponDetails  `json:"coupon"`
            PromoMeta PromoMeta  `json:"promo_meta"`
            Message string  `json:"message"`
            Article ProductArticle  `json:"article"`
            Quantity float64  `json:"quantity"`
            Key string  `json:"key"`
            Discount string  `json:"discount"`
            Product CartProduct  `json:"product"`
            IsSet bool  `json:"is_set"`
            CouponMessage string  `json:"coupon_message"`
            Availability ProductAvailability  `json:"availability"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
            Moq map[string]interface{}  `json:"moq"`
            BulkOffer map[string]interface{}  `json:"bulk_offer"`
            PricePerUnit ProductPriceInfo  `json:"price_per_unit"`
            PromotionsApplied []AppliedPromotion  `json:"promotions_applied"`
            Price ProductPriceInfo  `json:"price"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
         
    }
    
    // DisplayBreakupSchema ...
    type DisplayBreakupSchema struct {

        
            Message []string  `json:"message"`
            CurrencyCode string  `json:"currency_code"`
            CurrencySymbol string  `json:"currency_symbol"`
            Key string  `json:"key"`
            Value float64  `json:"value"`
            Display string  `json:"display"`
         
    }
    
    // CouponBreakupSchema ...
    type CouponBreakupSchema struct {

        
            Code string  `json:"code"`
            Type string  `json:"type"`
            Description string  `json:"description"`
            MinimumCartValue float64  `json:"minimum_cart_value"`
            Title string  `json:"title"`
            Value float64  `json:"value"`
            MaxDiscountValue float64  `json:"max_discount_value"`
            CouponType string  `json:"coupon_type"`
            UID string  `json:"uid"`
            SubTitle string  `json:"sub_title"`
            CouponValue float64  `json:"coupon_value"`
            Message string  `json:"message"`
            IsApplied bool  `json:"is_applied"`
         
    }
    
    // RawBreakupSchema ...
    type RawBreakupSchema struct {

        
            ConvenienceFee float64  `json:"convenience_fee"`
            YouSaved float64  `json:"you_saved"`
            GstCharges float64  `json:"gst_charges"`
            CodCharge float64  `json:"cod_charge"`
            MrpTotal float64  `json:"mrp_total"`
            Vog float64  `json:"vog"`
            GiftCard float64  `json:"gift_card"`
            FyndCash float64  `json:"fynd_cash"`
            DeliveryCharge float64  `json:"delivery_charge"`
            Discount float64  `json:"discount"`
            Subtotal float64  `json:"subtotal"`
            Coupon float64  `json:"coupon"`
            Total float64  `json:"total"`
         
    }
    
    // LoyaltyPoints ...
    type LoyaltyPoints struct {

        
            IsApplied bool  `json:"is_applied"`
            Total float64  `json:"total"`
            Description string  `json:"description"`
            Applicable float64  `json:"applicable"`
         
    }
    
    // CartBreakupSchema ...
    type CartBreakupSchema struct {

        
            Display []DisplayBreakupSchema  `json:"display"`
            Coupon CouponBreakupSchema  `json:"coupon"`
            Raw RawBreakupSchema  `json:"raw"`
            LoyaltyPoints LoyaltyPoints  `json:"loyalty_points"`
         
    }
    
    // CartDetailResponse ...
    type CartDetailResponse struct {

        
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            LastModified string  `json:"last_modified"`
            Message string  `json:"message"`
            Currency CartCurrencySchema  `json:"currency"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            PanNo string  `json:"pan_no"`
            PanConfig map[string]interface{}  `json:"pan_config"`
            Gstin string  `json:"gstin"`
            ID string  `json:"id"`
            Comment string  `json:"comment"`
            BuyNow bool  `json:"buy_now"`
            Items []CartProductInfo  `json:"items"`
            CouponText string  `json:"coupon_text"`
            BreakupValues CartBreakupSchema  `json:"breakup_values"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            AppliedPromoDetails []AppliedPromotion  `json:"applied_promo_details"`
            CheckoutMode string  `json:"checkout_mode"`
            IsValid bool  `json:"is_valid"`
         
    }
    
    // AddProductCart ...
    type AddProductCart struct {

        
            ArticleID string  `json:"article_id"`
            StoreID float64  `json:"store_id"`
            ItemSize string  `json:"item_size"`
            Pos bool  `json:"pos"`
            Quantity float64  `json:"quantity"`
            ArticleAssignment map[string]interface{}  `json:"article_assignment"`
            ProductGroupTags []string  `json:"product_group_tags"`
            ItemID float64  `json:"item_id"`
            SellerID float64  `json:"seller_id"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Display string  `json:"display"`
            ParentItemIdentifiers []map[string]string  `json:"parent_item_identifiers"`
         
    }
    
    // AddCartRequest ...
    type AddCartRequest struct {

        
            Items []AddProductCart  `json:"items"`
            NewCart bool  `json:"new_cart"`
         
    }
    
    // AddCartDetailResponse ...
    type AddCartDetailResponse struct {

        
            Cart CartDetailResponse  `json:"cart"`
            Message string  `json:"message"`
            Partial bool  `json:"partial"`
            Success bool  `json:"success"`
         
    }
    
    // UpdateProductCart ...
    type UpdateProductCart struct {

        
            ArticleID string  `json:"article_id"`
            ItemIndex float64  `json:"item_index"`
            ItemSize string  `json:"item_size"`
            Quantity float64  `json:"quantity"`
            ItemID float64  `json:"item_id"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
         
    }
    
    // UpdateCartRequest ...
    type UpdateCartRequest struct {

        
            Items []UpdateProductCart  `json:"items"`
            Operation string  `json:"operation"`
         
    }
    
    // UpdateCartDetailResponse ...
    type UpdateCartDetailResponse struct {

        
            Cart CartDetailResponse  `json:"cart"`
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
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
    
    // PageCouponSchema ...
    type PageCouponSchema struct {

        
            HasNext bool  `json:"has_next"`
            Current float64  `json:"current"`
            HasPrevious bool  `json:"has_previous"`
            TotalItemCount float64  `json:"total_item_count"`
            Total float64  `json:"total"`
         
    }
    
    // Coupon ...
    type Coupon struct {

        
            Description string  `json:"description"`
            MinimumCartValue float64  `json:"minimum_cart_value"`
            ExpiresOn string  `json:"expires_on"`
            Title string  `json:"title"`
            IsApplicable bool  `json:"is_applicable"`
            MaxDiscountValue float64  `json:"max_discount_value"`
            CouponType string  `json:"coupon_type"`
            SubTitle string  `json:"sub_title"`
            CouponCode string  `json:"coupon_code"`
            CouponValue float64  `json:"coupon_value"`
            Message string  `json:"message"`
            IsApplied bool  `json:"is_applied"`
         
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

        
            Name string  `json:"name"`
            UID float64  `json:"uid"`
         
    }
    
    // OfferPrice ...
    type OfferPrice struct {

        
            BulkEffective float64  `json:"bulk_effective"`
            CurrencyCode string  `json:"currency_code"`
            CurrencySymbol string  `json:"currency_symbol"`
            Effective float64  `json:"effective"`
            Marked float64  `json:"marked"`
         
    }
    
    // OfferItemSchema ...
    type OfferItemSchema struct {

        
            Type string  `json:"type"`
            Best bool  `json:"best"`
            Quantity float64  `json:"quantity"`
            AutoApplied bool  `json:"auto_applied"`
            Price OfferPrice  `json:"price"`
            Margin float64  `json:"margin"`
            Total float64  `json:"total"`
         
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

        
            CountryCode string  `json:"country_code"`
            IsActive bool  `json:"is_active"`
            UserID string  `json:"user_id"`
            CreatedByUserID string  `json:"created_by_user_id"`
            Name string  `json:"name"`
            AddressType string  `json:"address_type"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Meta map[string]interface{}  `json:"meta"`
            State string  `json:"state"`
            IsDefaultAddress bool  `json:"is_default_address"`
            Address string  `json:"address"`
            Email string  `json:"email"`
            GoogleMapPoint map[string]interface{}  `json:"google_map_point"`
            Phone string  `json:"phone"`
            AreaCodeSlug string  `json:"area_code_slug"`
            AreaCode string  `json:"area_code"`
            Area string  `json:"area"`
            CountryIsoCode string  `json:"country_iso_code"`
            ID string  `json:"id"`
            Landmark string  `json:"landmark"`
            Tags []string  `json:"tags"`
            CountryPhoneCode string  `json:"country_phone_code"`
            GeoLocation GeoLocation  `json:"geo_location"`
            City string  `json:"city"`
            CheckoutMode string  `json:"checkout_mode"`
            Country string  `json:"country"`
         
    }
    
    // GetAddressesResponse ...
    type GetAddressesResponse struct {

        
            PiiMasking bool  `json:"pii_masking"`
            Address []Address  `json:"address"`
         
    }
    
    // SaveAddressResponse ...
    type SaveAddressResponse struct {

        
            ID string  `json:"id"`
            IsDefaultAddress bool  `json:"is_default_address"`
            Success bool  `json:"success"`
         
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

        
            PaymentIdentifier string  `json:"payment_identifier"`
            MerchantCode string  `json:"merchant_code"`
            ID string  `json:"id"`
            AddressID string  `json:"address_id"`
            AggregatorName string  `json:"aggregator_name"`
            PaymentMode string  `json:"payment_mode"`
         
    }
    
    // CouponValidity ...
    type CouponValidity struct {

        
            DisplayMessageEn string  `json:"display_message_en"`
            Code string  `json:"code"`
            Title string  `json:"title"`
            Discount float64  `json:"discount"`
            Valid bool  `json:"valid"`
         
    }
    
    // PaymentCouponValidateSchema ...
    type PaymentCouponValidateSchema struct {

        
            Success bool  `json:"success"`
            CouponValidity CouponValidity  `json:"coupon_validity"`
            Message string  `json:"message"`
         
    }
    
    // ShipmentResponse ...
    type ShipmentResponse struct {

        
            Promise ShipmentPromise  `json:"promise"`
            DpID string  `json:"dp_id"`
            DpOptions map[string]interface{}  `json:"dp_options"`
            Shipments float64  `json:"shipments"`
            ShipmentType string  `json:"shipment_type"`
            OrderType string  `json:"order_type"`
            FulfillmentID float64  `json:"fulfillment_id"`
            BoxType string  `json:"box_type"`
            FulfillmentType string  `json:"fulfillment_type"`
            Items []CartProductInfo  `json:"items"`
         
    }
    
    // CartShipmentsResponse ...
    type CartShipmentsResponse struct {

        
            Gstin string  `json:"gstin"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            LastModified string  `json:"last_modified"`
            ID string  `json:"id"`
            Comment string  `json:"comment"`
            BuyNow bool  `json:"buy_now"`
            CartID float64  `json:"cart_id"`
            CouponText string  `json:"coupon_text"`
            BreakupValues CartBreakupSchema  `json:"breakup_values"`
            Message string  `json:"message"`
            Currency CartCurrencySchema  `json:"currency"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            Shipments []ShipmentResponse  `json:"shipments"`
            UID string  `json:"uid"`
            CheckoutMode string  `json:"checkout_mode"`
            Error bool  `json:"error"`
            IsValid bool  `json:"is_valid"`
         
    }
    
    // StaffCheckoutSchema ...
    type StaffCheckoutSchema struct {

        
            LastName string  `json:"last_name"`
            User string  `json:"user"`
            FirstName string  `json:"first_name"`
            ID string  `json:"_id"`
            EmployeeCode string  `json:"employee_code"`
         
    }
    
    // CartCheckoutCustomMetaSchema ...
    type CartCheckoutCustomMetaSchema struct {

        
            Value string  `json:"value"`
            Key string  `json:"key"`
         
    }
    
    // CartCheckoutDetailRequest ...
    type CartCheckoutDetailRequest struct {

        
            BillingAddressID string  `json:"billing_address_id"`
            CallbackURL string  `json:"callback_url"`
            PaymentIdentifier string  `json:"payment_identifier"`
            PaymentAutoConfirm bool  `json:"payment_auto_confirm"`
            MerchantCode string  `json:"merchant_code"`
            OrderingStore float64  `json:"ordering_store"`
            ID string  `json:"id"`
            AddressID string  `json:"address_id"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
            PaymentMode string  `json:"payment_mode"`
            OrderType string  `json:"order_type"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            PaymentParams map[string]interface{}  `json:"payment_params"`
            Staff StaffCheckoutSchema  `json:"staff"`
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
            CustomMeta []CartCheckoutCustomMetaSchema  `json:"custom_meta"`
            Aggregator string  `json:"aggregator"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // CheckCartSchema ...
    type CheckCartSchema struct {

        
            CodMessage string  `json:"cod_message"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            LastModified string  `json:"last_modified"`
            CodCharges float64  `json:"cod_charges"`
            DeliveryChargeOrderValue float64  `json:"delivery_charge_order_value"`
            Message string  `json:"message"`
            StoreCode string  `json:"store_code"`
            Currency CartCurrencySchema  `json:"currency"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            OrderID string  `json:"order_id"`
            ErrorMessage string  `json:"error_message"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            DeliveryCharges float64  `json:"delivery_charges"`
            UID string  `json:"uid"`
            Success bool  `json:"success"`
            Gstin string  `json:"gstin"`
            ID string  `json:"id"`
            Comment string  `json:"comment"`
            BuyNow bool  `json:"buy_now"`
            Items []CartProductInfo  `json:"items"`
            CartID float64  `json:"cart_id"`
            StoreEmps []map[string]interface{}  `json:"store_emps"`
            CouponText string  `json:"coupon_text"`
            BreakupValues CartBreakupSchema  `json:"breakup_values"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            CodAvailable bool  `json:"cod_available"`
            UserType string  `json:"user_type"`
            CheckoutMode string  `json:"checkout_mode"`
            IsValid bool  `json:"is_valid"`
         
    }
    
    // CartCheckoutResponseSchema ...
    type CartCheckoutResponseSchema struct {

        
            CallbackURL string  `json:"callback_url"`
            OrderID string  `json:"order_id"`
            Data map[string]interface{}  `json:"data"`
            PaymentConfirmURL string  `json:"payment_confirm_url"`
            Cart CheckCartSchema  `json:"cart"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
            AppInterceptURL string  `json:"app_intercept_url"`
         
    }
    
    // GiftDetailSchema ...
    type GiftDetailSchema struct {

        
            GiftMessage string  `json:"gift_message"`
            IsGiftApplied bool  `json:"is_gift_applied"`
         
    }
    
    // ArticleGiftDetailSchema ...
    type ArticleGiftDetailSchema struct {

        
            ArticleID GiftDetailSchema  `json:"article_id"`
         
    }
    
    // CartMetaRequest ...
    type CartMetaRequest struct {

        
            PickUpCustomerDetails map[string]interface{}  `json:"pick_up_customer_details"`
            Gstin string  `json:"gstin"`
            Comment string  `json:"comment"`
            GiftDetails ArticleGiftDetailSchema  `json:"gift_details"`
            CheckoutMode string  `json:"checkout_mode"`
         
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

        
            Token string  `json:"token"`
            User map[string]interface{}  `json:"user"`
            CreatedOn string  `json:"created_on"`
            Source map[string]interface{}  `json:"source"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // SharedCart ...
    type SharedCart struct {

        
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            LastModified string  `json:"last_modified"`
            Message string  `json:"message"`
            SharedCartDetails SharedCartDetails  `json:"shared_cart_details"`
            Currency CartCurrencySchema  `json:"currency"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            UID string  `json:"uid"`
            Gstin string  `json:"gstin"`
            ID string  `json:"id"`
            Comment string  `json:"comment"`
            BuyNow bool  `json:"buy_now"`
            Items []CartProductInfo  `json:"items"`
            CartID float64  `json:"cart_id"`
            CouponText string  `json:"coupon_text"`
            BreakupValues CartBreakupSchema  `json:"breakup_values"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            CheckoutMode string  `json:"checkout_mode"`
            IsValid bool  `json:"is_valid"`
         
    }
    
    // SharedCartResponse ...
    type SharedCartResponse struct {

        
            Cart SharedCart  `json:"cart"`
            Error string  `json:"error"`
         
    }
    
    // FreeGiftItemsSchema ...
    type FreeGiftItemsSchema struct {

        
            ItemImagesURL []string  `json:"item_images_url"`
            ItemSlug string  `json:"item_slug"`
            ItemPriceDetails map[string]interface{}  `json:"item_price_details"`
            ItemName string  `json:"item_name"`
            ItemID float64  `json:"item_id"`
            ItemBrandName string  `json:"item_brand_name"`
         
    }
    
    // PromotionOffer ...
    type PromotionOffer struct {

        
            PromotionGroup string  `json:"promotion_group"`
            ValidTill string  `json:"valid_till"`
            BuyRules map[string]interface{}  `json:"buy_rules"`
            Description string  `json:"description"`
            ID string  `json:"id"`
            DiscountRules []map[string]interface{}  `json:"discount_rules"`
            FreeGiftItems []FreeGiftItemsSchema  `json:"free_gift_items"`
            OfferText string  `json:"offer_text"`
         
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
    
    // CurrencyInfoSchema ...
    type CurrencyInfoSchema struct {

        
            Symbol string  `json:"symbol"`
            Code string  `json:"code"`
         
    }
    
    // LadderPrice ...
    type LadderPrice struct {

        
            OfferPrice float64  `json:"offer_price"`
            CurrencyCode string  `json:"currency_code"`
            CurrencySymbol string  `json:"currency_symbol"`
            Effective float64  `json:"effective"`
            Marked float64  `json:"marked"`
         
    }
    
    // LadderOfferItemSchema ...
    type LadderOfferItemSchema struct {

        
            MinQuantity float64  `json:"min_quantity"`
            Type string  `json:"type"`
            MaxQuantity float64  `json:"max_quantity"`
            Price LadderPrice  `json:"price"`
            Margin float64  `json:"margin"`
         
    }
    
    // LadderPriceOffer ...
    type LadderPriceOffer struct {

        
            PromotionGroup string  `json:"promotion_group"`
            ValidTill string  `json:"valid_till"`
            BuyRules map[string]interface{}  `json:"buy_rules"`
            Description string  `json:"description"`
            ID string  `json:"id"`
            DiscountRules []map[string]interface{}  `json:"discount_rules"`
            CalculateOn string  `json:"calculate_on"`
            FreeGiftItems []FreeGiftItemsSchema  `json:"free_gift_items"`
            OfferText string  `json:"offer_text"`
            OfferPrices []LadderOfferItemSchema  `json:"offer_prices"`
         
    }
    
    // LadderPriceOffers ...
    type LadderPriceOffers struct {

        
            Currency CurrencyInfoSchema  `json:"currency"`
            AvailableOffers []LadderPriceOffer  `json:"available_offers"`
         
    }
    
    // PaymentMetaSchema ...
    type PaymentMetaSchema struct {

        
            PaymentGateway string  `json:"payment_gateway"`
            PaymentIdentifier string  `json:"payment_identifier"`
            Type string  `json:"type"`
            MerchantCode string  `json:"merchant_code"`
         
    }
    
    // PaymentMethodSchema ...
    type PaymentMethodSchema struct {

        
            Amount float64  `json:"amount"`
            Mode string  `json:"mode"`
            Payment string  `json:"payment"`
            PaymentMeta PaymentMetaSchema  `json:"payment_meta"`
            Name string  `json:"name"`
         
    }
    
    // CartCheckoutDetailV2Request ...
    type CartCheckoutDetailV2Request struct {

        
            CallbackURL string  `json:"callback_url"`
            AddressID string  `json:"address_id"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            Meta map[string]interface{}  `json:"meta"`
            PaymentIdentifier string  `json:"payment_identifier"`
            MerchantCode string  `json:"merchant_code"`
            PaymentMode string  `json:"payment_mode"`
            PaymentMethods []PaymentMethodSchema  `json:"payment_methods"`
            PaymentAutoConfirm bool  `json:"payment_auto_confirm"`
            Staff StaffCheckoutSchema  `json:"staff"`
            CustomMeta map[string]interface{}  `json:"custom_meta"`
            Aggregator string  `json:"aggregator"`
            ID string  `json:"id"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
            CartID string  `json:"cart_id"`
            BillingAddressID string  `json:"billing_address_id"`
            OrderingStore float64  `json:"ordering_store"`
            OrderType string  `json:"order_type"`
            PaymentParams map[string]interface{}  `json:"payment_params"`
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
         
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

        
            Pin string  `json:"pin"`
            Key string  `json:"key"`
            VerifyAPI string  `json:"verify_api"`
            Sdk bool  `json:"sdk"`
            UserID string  `json:"user_id"`
            ConfigType string  `json:"config_type"`
            API string  `json:"api"`
            Secret string  `json:"secret"`
            MerchantID string  `json:"merchant_id"`
            MerchantKey string  `json:"merchant_key"`
         
    }
    
    // AggregatorsConfigDetailResponse ...
    type AggregatorsConfigDetailResponse struct {

        
            Stripe AggregatorConfigDetail  `json:"stripe"`
            Razorpay AggregatorConfigDetail  `json:"razorpay"`
            Payumoney AggregatorConfigDetail  `json:"payumoney"`
            Ccavenue AggregatorConfigDetail  `json:"ccavenue"`
            Env string  `json:"env"`
            Mswipe AggregatorConfigDetail  `json:"mswipe"`
            Juspay AggregatorConfigDetail  `json:"juspay"`
            Simpl AggregatorConfigDetail  `json:"simpl"`
            Rupifi AggregatorConfigDetail  `json:"rupifi"`
            Success bool  `json:"success"`
         
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

        
            Refresh bool  `json:"refresh"`
            CardID string  `json:"card_id"`
            Nickname string  `json:"nickname"`
            NameOnCard string  `json:"name_on_card"`
         
    }
    
    // AttachCardsResponse ...
    type AttachCardsResponse struct {

        
            Data map[string]interface{}  `json:"data"`
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // CardPaymentGateway ...
    type CardPaymentGateway struct {

        
            Aggregator string  `json:"aggregator"`
            API string  `json:"api"`
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

        
            ExpMonth float64  `json:"exp_month"`
            Expired bool  `json:"expired"`
            CompliantWithTokenisationGuidelines bool  `json:"compliant_with_tokenisation_guidelines"`
            CardFingerprint string  `json:"card_fingerprint"`
            Nickname string  `json:"nickname"`
            CardToken string  `json:"card_token"`
            ExpYear float64  `json:"exp_year"`
            AggregatorName string  `json:"aggregator_name"`
            CardName string  `json:"card_name"`
            CardIsin string  `json:"card_isin"`
            CardType string  `json:"card_type"`
            CardID string  `json:"card_id"`
            CardBrand string  `json:"card_brand"`
            CardIssuer string  `json:"card_issuer"`
            CardNumber string  `json:"card_number"`
            CardBrandImage string  `json:"card_brand_image"`
            CardReference string  `json:"card_reference"`
         
    }
    
    // ListCardsResponse ...
    type ListCardsResponse struct {

        
            Data []Card  `json:"data"`
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
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

        
            MerchantParams map[string]interface{}  `json:"merchant_params"`
            OrderItems []map[string]interface{}  `json:"order_items"`
            TransactionAmountInPaise float64  `json:"transaction_amount_in_paise"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
            Payload string  `json:"payload"`
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
            PhoneNumber string  `json:"phone_number"`
            Aggregator string  `json:"aggregator"`
         
    }
    
    // ValidateCustomerResponse ...
    type ValidateCustomerResponse struct {

        
            Error map[string]interface{}  `json:"error"`
            Data map[string]interface{}  `json:"data"`
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // ChargeCustomerRequest ...
    type ChargeCustomerRequest struct {

        
            Verified bool  `json:"verified"`
            Amount float64  `json:"amount"`
            OrderID string  `json:"order_id"`
            TransactionToken string  `json:"transaction_token"`
            Aggregator string  `json:"aggregator"`
         
    }
    
    // ChargeCustomerResponse ...
    type ChargeCustomerResponse struct {

        
            Message string  `json:"message"`
            OrderID string  `json:"order_id"`
            Status string  `json:"status"`
            DeliveryAddressID string  `json:"delivery_address_id"`
            CartID string  `json:"cart_id"`
            Aggregator string  `json:"aggregator"`
            Success bool  `json:"success"`
         
    }
    
    // PaymentInitializationRequest ...
    type PaymentInitializationRequest struct {

        
            RazorpayPaymentID string  `json:"razorpay_payment_id"`
            CustomerID string  `json:"customer_id"`
            Vpa string  `json:"vpa"`
            Timeout float64  `json:"timeout"`
            DeviceID string  `json:"device_id"`
            Amount float64  `json:"amount"`
            Email string  `json:"email"`
            OrderID string  `json:"order_id"`
            Currency string  `json:"currency"`
            MerchantOrderID string  `json:"merchant_order_id"`
            Method string  `json:"method"`
            Aggregator string  `json:"aggregator"`
            Contact string  `json:"contact"`
         
    }
    
    // PaymentInitializationResponse ...
    type PaymentInitializationResponse struct {

        
            PollingURL string  `json:"polling_url"`
            Timeout float64  `json:"timeout"`
            BqrImage string  `json:"bqr_image"`
            Status string  `json:"status"`
            Method string  `json:"method"`
            Aggregator string  `json:"aggregator"`
            AggregatorOrderID string  `json:"aggregator_order_id"`
            Vpa string  `json:"vpa"`
            Amount float64  `json:"amount"`
            MerchantOrderID string  `json:"merchant_order_id"`
            PaymentID string  `json:"payment_id"`
            Contact string  `json:"contact"`
            Success bool  `json:"success"`
            RazorpayPaymentID string  `json:"razorpay_payment_id"`
            UpiPollURL string  `json:"upi_poll_url"`
            Email string  `json:"email"`
            CustomerID string  `json:"customer_id"`
            Currency string  `json:"currency"`
            DeviceID string  `json:"device_id"`
            OrderID string  `json:"order_id"`
            VirtualID string  `json:"virtual_id"`
         
    }
    
    // PaymentStatusUpdateRequest ...
    type PaymentStatusUpdateRequest struct {

        
            MerchantTransactionID string  `json:"merchant_transaction_id"`
            CustomerID string  `json:"customer_id"`
            Vpa string  `json:"vpa"`
            RazorpayPaymentID string  `json:"razorpay_payment_id"`
            DeviceID string  `json:"device_id"`
            Amount float64  `json:"amount"`
            Email string  `json:"email"`
            OrderID string  `json:"order_id"`
            Currency string  `json:"currency"`
            Status string  `json:"status"`
            MerchantOrderID string  `json:"merchant_order_id"`
            PaymentID string  `json:"payment_id"`
            Method string  `json:"method"`
            Aggregator string  `json:"aggregator"`
            Contact string  `json:"contact"`
         
    }
    
    // PaymentStatusUpdateResponse ...
    type PaymentStatusUpdateResponse struct {

        
            AggregatorName string  `json:"aggregator_name"`
            Retry bool  `json:"retry"`
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
            Logos PaymentModeLogo  `json:"logos"`
            PackageName string  `json:"package_name"`
         
    }
    
    // IntentAppErrorList ...
    type IntentAppErrorList struct {

        
            Code string  `json:"code"`
            PackageName string  `json:"package_name"`
         
    }
    
    // PaymentModeList ...
    type PaymentModeList struct {

        
            DisplayPriority float64  `json:"display_priority"`
            Expired bool  `json:"expired"`
            CodLimit float64  `json:"cod_limit"`
            IntentFlow bool  `json:"intent_flow"`
            IntentAppErrorList []string  `json:"intent_app_error_list"`
            CodLimitPerOrder float64  `json:"cod_limit_per_order"`
            AggregatorName string  `json:"aggregator_name"`
            Timeout float64  `json:"timeout"`
            RemainingLimit float64  `json:"remaining_limit"`
            CardNumber string  `json:"card_number"`
            Code string  `json:"code"`
            ExpYear float64  `json:"exp_year"`
            RetryCount float64  `json:"retry_count"`
            CardFingerprint string  `json:"card_fingerprint"`
            Nickname string  `json:"nickname"`
            CardType string  `json:"card_type"`
            LogoURL PaymentModeLogo  `json:"logo_url"`
            CardBrand string  `json:"card_brand"`
            IntentApp []IntentApp  `json:"intent_app"`
            CardIssuer string  `json:"card_issuer"`
            CardBrandImage string  `json:"card_brand_image"`
            ExpMonth float64  `json:"exp_month"`
            CompliantWithTokenisationGuidelines bool  `json:"compliant_with_tokenisation_guidelines"`
            CardIsin string  `json:"card_isin"`
            MerchantCode string  `json:"merchant_code"`
            CardToken string  `json:"card_token"`
            CardName string  `json:"card_name"`
            FyndVpa string  `json:"fynd_vpa"`
            IntentAppErrorDictList []IntentAppErrorList  `json:"intent_app_error_dict_list"`
            DisplayName string  `json:"display_name"`
            Name string  `json:"name"`
            CardID string  `json:"card_id"`
            CardReference string  `json:"card_reference"`
         
    }
    
    // RootPaymentMode ...
    type RootPaymentMode struct {

        
            DisplayPriority float64  `json:"display_priority"`
            AggregatorName string  `json:"aggregator_name"`
            AnonymousEnable bool  `json:"anonymous_enable"`
            AddCardEnabled bool  `json:"add_card_enabled"`
            List []PaymentModeList  `json:"list"`
            DisplayName string  `json:"display_name"`
            Name string  `json:"name"`
            SaveCard bool  `json:"save_card"`
            IsPayByCardPl bool  `json:"is_pay_by_card_pl"`
         
    }
    
    // AggregatorRoute ...
    type AggregatorRoute struct {

        
            Data map[string]interface{}  `json:"data"`
            APILink string  `json:"api_link"`
            PaymentFlow string  `json:"payment_flow"`
            PaymentFlowData string  `json:"payment_flow_data"`
         
    }
    
    // PaymentFlow ...
    type PaymentFlow struct {

        
            Stripe AggregatorRoute  `json:"stripe"`
            Razorpay AggregatorRoute  `json:"razorpay"`
            Ccavenue AggregatorRoute  `json:"ccavenue"`
            BqrRazorpay AggregatorRoute  `json:"bqr_razorpay"`
            UpiRazorpay AggregatorRoute  `json:"upi_razorpay"`
            Mswipe AggregatorRoute  `json:"mswipe"`
            Juspay AggregatorRoute  `json:"juspay"`
            Fynd AggregatorRoute  `json:"fynd"`
            Payubiz AggregatorRoute  `json:"payubiz"`
            Simpl AggregatorRoute  `json:"simpl"`
            Epaylater AggregatorRoute  `json:"epaylater"`
            Rupifi AggregatorRoute  `json:"rupifi"`
            Jiopay AggregatorRoute  `json:"jiopay"`
         
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

        
            KycURL string  `json:"kyc_url"`
            Status string  `json:"status"`
         
    }
    
    // RupifiBannerResponse ...
    type RupifiBannerResponse struct {

        
            Data RupifiBannerData  `json:"data"`
            Success bool  `json:"success"`
         
    }
    
    // EpaylaterBannerData ...
    type EpaylaterBannerData struct {

        
            Message string  `json:"message"`
            Status string  `json:"status"`
            Display bool  `json:"display"`
         
    }
    
    // EpaylaterBannerResponse ...
    type EpaylaterBannerResponse struct {

        
            Data EpaylaterBannerData  `json:"data"`
            Success bool  `json:"success"`
         
    }
    
    // ResendOrCancelPaymentRequest ...
    type ResendOrCancelPaymentRequest struct {

        
            OrderID string  `json:"order_id"`
            RequestType string  `json:"request_type"`
            DeviceID string  `json:"device_id"`
         
    }
    
    // LinkStatus ...
    type LinkStatus struct {

        
            Message string  `json:"message"`
            Status bool  `json:"status"`
         
    }
    
    // ResendOrCancelPaymentResponse ...
    type ResendOrCancelPaymentResponse struct {

        
            Data LinkStatus  `json:"data"`
            Success bool  `json:"success"`
         
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
            Status string  `json:"status"`
            IsValid bool  `json:"is_valid"`
            UpiVpa string  `json:"upi_vpa"`
         
    }
    
    // ValidateVPAResponse ...
    type ValidateVPAResponse struct {

        
            Data ValidateUPI  `json:"data"`
            Success bool  `json:"success"`
         
    }
    
    // TransferItemsDetails ...
    type TransferItemsDetails struct {

        
            ID float64  `json:"id"`
            LogoSmall string  `json:"logo_small"`
            LogoLarge string  `json:"logo_large"`
            DisplayName string  `json:"display_name"`
            Name string  `json:"name"`
         
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

        
            Mobile string  `json:"mobile"`
            TransferMode string  `json:"transfer_mode"`
            Title string  `json:"title"`
            IsActive bool  `json:"is_active"`
            DelightsUserName string  `json:"delights_user_name"`
            BankName string  `json:"bank_name"`
            ID float64  `json:"id"`
            AccountNo string  `json:"account_no"`
            IfscCode string  `json:"ifsc_code"`
            Address string  `json:"address"`
            Email string  `json:"email"`
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
            Comment string  `json:"comment"`
            BranchName string  `json:"branch_name"`
            DisplayName string  `json:"display_name"`
            BeneficiaryID string  `json:"beneficiary_id"`
            Subtitle string  `json:"subtitle"`
            AccountHolder string  `json:"account_holder"`
         
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
            Success bool  `json:"success"`
            BranchName string  `json:"branch_name"`
         
    }
    
    // ErrorCodeDescription ...
    type ErrorCodeDescription struct {

        
            Code string  `json:"code"`
            Description string  `json:"description"`
            Success bool  `json:"success"`
         
    }
    
    // AddBeneficiaryViaOtpVerificationRequest ...
    type AddBeneficiaryViaOtpVerificationRequest struct {

        
            Otp string  `json:"otp"`
            HashKey string  `json:"hash_key"`
            RequestID string  `json:"request_id"`
         
    }
    
    // AddBeneficiaryViaOtpVerificationResponse ...
    type AddBeneficiaryViaOtpVerificationResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // WrongOtpError ...
    type WrongOtpError struct {

        
            Description string  `json:"description"`
            IsVerifiedFlag bool  `json:"is_verified_flag"`
            Success string  `json:"success"`
         
    }
    
    // BeneficiaryModeDetails ...
    type BeneficiaryModeDetails struct {

        
            Mobile string  `json:"mobile"`
            Vpa string  `json:"vpa"`
            BankName string  `json:"bank_name"`
            Wallet string  `json:"wallet"`
            Comment string  `json:"comment"`
            AccountNo string  `json:"account_no"`
            BranchName string  `json:"branch_name"`
            Email string  `json:"email"`
            IfscCode string  `json:"ifsc_code"`
            Address string  `json:"address"`
            AccountHolder string  `json:"account_holder"`
         
    }
    
    // AddBeneficiaryDetailsRequest ...
    type AddBeneficiaryDetailsRequest struct {

        
            Otp string  `json:"otp"`
            TransferMode string  `json:"transfer_mode"`
            Details BeneficiaryModeDetails  `json:"details"`
            Delights bool  `json:"delights"`
            OrderID string  `json:"order_id"`
            RequestID string  `json:"request_id"`
            ShipmentID string  `json:"shipment_id"`
         
    }
    
    // RefundAccountResponse ...
    type RefundAccountResponse struct {

        
            Data map[string]interface{}  `json:"data"`
            IsVerifiedFlag bool  `json:"is_verified_flag"`
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // BankDetailsForOTP ...
    type BankDetailsForOTP struct {

        
            BankName string  `json:"bank_name"`
            AccountNo string  `json:"account_no"`
            BranchName string  `json:"branch_name"`
            IfscCode string  `json:"ifsc_code"`
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

        
            IsVerifiedFlag string  `json:"is_verified_flag"`
            RequestID string  `json:"request_id"`
            Success bool  `json:"success"`
         
    }
    
    // SetDefaultBeneficiaryRequest ...
    type SetDefaultBeneficiaryRequest struct {

        
            OrderID string  `json:"order_id"`
            BeneficiaryID string  `json:"beneficiary_id"`
         
    }
    
    // SetDefaultBeneficiaryResponse ...
    type SetDefaultBeneficiaryResponse struct {

        
            IsBeneficiarySet bool  `json:"is_beneficiary_set"`
            Success bool  `json:"success"`
         
    }
    
    // GetPaymentLinkResponse ...
    type GetPaymentLinkResponse struct {

        
            Message string  `json:"message"`
            PaymentLinkURL string  `json:"payment_link_url"`
            StatusCode float64  `json:"status_code"`
            ExternalOrderID string  `json:"external_order_id"`
            MerchantName string  `json:"merchant_name"`
            PollingTimeout float64  `json:"polling_timeout"`
            PaymentLinkCurrentStatus string  `json:"payment_link_current_status"`
            Amount float64  `json:"amount"`
            Success bool  `json:"success"`
         
    }
    
    // ErrorDescription ...
    type ErrorDescription struct {

        
            Expired bool  `json:"expired"`
            Msg string  `json:"msg"`
            MerchantName string  `json:"merchant_name"`
            InvalidID bool  `json:"invalid_id"`
            Cancelled bool  `json:"cancelled"`
            PaymentTransactionID string  `json:"payment_transaction_id"`
            Amount float64  `json:"amount"`
            MerchantOrderID string  `json:"merchant_order_id"`
         
    }
    
    // ErrorResponse ...
    type ErrorResponse struct {

        
            Error ErrorDescription  `json:"error"`
            StatusCode float64  `json:"status_code"`
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // CreatePaymentLinkMeta ...
    type CreatePaymentLinkMeta struct {

        
            Pincode string  `json:"pincode"`
            Amount string  `json:"amount"`
            CartID string  `json:"cart_id"`
            AssignCardID string  `json:"assign_card_id"`
            CheckoutMode string  `json:"checkout_mode"`
         
    }
    
    // CreatePaymentLinkRequest ...
    type CreatePaymentLinkRequest struct {

        
            MobileNumber string  `json:"mobile_number"`
            ExternalOrderID string  `json:"external_order_id"`
            Meta CreatePaymentLinkMeta  `json:"meta"`
            Amount float64  `json:"amount"`
            Email string  `json:"email"`
            Description string  `json:"description"`
         
    }
    
    // CreatePaymentLinkResponse ...
    type CreatePaymentLinkResponse struct {

        
            Message string  `json:"message"`
            PaymentLinkURL string  `json:"payment_link_url"`
            StatusCode float64  `json:"status_code"`
            PollingTimeout float64  `json:"polling_timeout"`
            PaymentLinkID string  `json:"payment_link_id"`
            Success bool  `json:"success"`
         
    }
    
    // CancelOrResendPaymentLinkRequest ...
    type CancelOrResendPaymentLinkRequest struct {

        
            PaymentLinkID string  `json:"payment_link_id"`
         
    }
    
    // ResendPaymentLinkResponse ...
    type ResendPaymentLinkResponse struct {

        
            StatusCode float64  `json:"status_code"`
            Message string  `json:"message"`
            PollingTimeout float64  `json:"polling_timeout"`
            Success bool  `json:"success"`
         
    }
    
    // CancelPaymentLinkResponse ...
    type CancelPaymentLinkResponse struct {

        
            StatusCode float64  `json:"status_code"`
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // PollingPaymentLinkResponse ...
    type PollingPaymentLinkResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
            AggregatorName string  `json:"aggregator_name"`
            StatusCode float64  `json:"status_code"`
            RedirectURL string  `json:"redirect_url"`
            Amount float64  `json:"amount"`
            OrderID string  `json:"order_id"`
            Status string  `json:"status"`
            PaymentLinkID string  `json:"payment_link_id"`
            HttpStatus float64  `json:"http_status"`
         
    }
    
    // PaymentMethodsMeta ...
    type PaymentMethodsMeta struct {

        
            PaymentGateway string  `json:"payment_gateway"`
            MerchantCode string  `json:"merchant_code"`
            PaymentIdentifier string  `json:"payment_identifier"`
         
    }
    
    // CreateOrderUserPaymentMethods ...
    type CreateOrderUserPaymentMethods struct {

        
            Name string  `json:"name"`
            Meta PaymentMethodsMeta  `json:"meta"`
            Mode string  `json:"mode"`
         
    }
    
    // CreateOrderUserRequest ...
    type CreateOrderUserRequest struct {

        
            PaymentMethods CreateOrderUserPaymentMethods  `json:"payment_methods"`
            Meta map[string]interface{}  `json:"meta"`
            PaymentLinkID string  `json:"payment_link_id"`
            SuccessCallbackURL string  `json:"success_callback_url"`
            FailureCallbackURL string  `json:"failure_callback_url"`
            Currency string  `json:"currency"`
         
    }
    
    // CreateOrderUserData ...
    type CreateOrderUserData struct {

        
            CustomerID string  `json:"customer_id"`
            Aggregator string  `json:"aggregator"`
            Amount float64  `json:"amount"`
            Email string  `json:"email"`
            OrderID string  `json:"order_id"`
            MerchantOrderID string  `json:"merchant_order_id"`
            Method string  `json:"method"`
            Currency string  `json:"currency"`
            Contact string  `json:"contact"`
            CallbackURL string  `json:"callback_url"`
         
    }
    
    // CreateOrderUserResponse ...
    type CreateOrderUserResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
            PaymentConfirmURL string  `json:"payment_confirm_url"`
            StatusCode float64  `json:"status_code"`
            Data CreateOrderUserData  `json:"data"`
            OrderID string  `json:"order_id"`
            CallbackURL string  `json:"callback_url"`
         
    }
    
    // BalanceDetails ...
    type BalanceDetails struct {

        
            Currency string  `json:"currency"`
            Value float64  `json:"value"`
            FormattedValue string  `json:"formatted_value"`
         
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
            Status bool  `json:"status"`
            IsRegistered bool  `json:"is_registered"`
         
    }
    
    // CheckCreditResponse ...
    type CheckCreditResponse struct {

        
            Data CreditDetail  `json:"data"`
            Success bool  `json:"success"`
         
    }
    
    // KYCAddress ...
    type KYCAddress struct {

        
            State string  `json:"state"`
            LandMark string  `json:"land_mark"`
            Pincode string  `json:"pincode"`
            Addressline2 string  `json:"addressline2"`
            City string  `json:"city"`
            Addressline1 string  `json:"addressline1"`
            OwnershipType string  `json:"ownership_type"`
         
    }
    
    // BusinessDetails ...
    type BusinessDetails struct {

        
            EntityType string  `json:"entity_type"`
            Fda string  `json:"fda"`
            BusinessType string  `json:"business_type"`
            Gstin string  `json:"gstin"`
            Name string  `json:"name"`
            ShopAndEstablishment map[string]interface{}  `json:"shop_and_establishment"`
            Address KYCAddress  `json:"address"`
            Pan string  `json:"pan"`
            Fssai string  `json:"fssai"`
            Vintage string  `json:"vintage"`
            BusinessOwnershipType string  `json:"business_ownership_type"`
         
    }
    
    // DeviceDetails ...
    type DeviceDetails struct {

        
            DeviceType string  `json:"device_type"`
            Os string  `json:"os"`
            IdentifierType string  `json:"identifier_type"`
            DeviceModel string  `json:"device_model"`
            DeviceMake string  `json:"device_make"`
            IdentificationNumber string  `json:"identification_number"`
            OsVersion string  `json:"os_version"`
         
    }
    
    // UserPersonalInfoInDetails ...
    type UserPersonalInfoInDetails struct {

        
            MobileVerified bool  `json:"mobile_verified"`
            DrivingLicense string  `json:"driving_license"`
            VoterID string  `json:"voter_id"`
            MiddleName string  `json:"middle_name"`
            Dob string  `json:"dob"`
            Gender string  `json:"gender"`
            LastName string  `json:"last_name"`
            FirstName string  `json:"first_name"`
            Email string  `json:"email"`
            FathersName string  `json:"fathers_name"`
            MothersName string  `json:"mothers_name"`
            Phone string  `json:"phone"`
            Pan string  `json:"pan"`
            AddressAsPerID KYCAddress  `json:"address_as_per_id"`
            Passport string  `json:"passport"`
            EmailVerified bool  `json:"email_verified"`
         
    }
    
    // MarketplaceInfo ...
    type MarketplaceInfo struct {

        
            Name string  `json:"name"`
            MembershipID string  `json:"membership_id"`
            DateOfJoining string  `json:"date_of_joining"`
         
    }
    
    // CustomerOnboardingRequest ...
    type CustomerOnboardingRequest struct {

        
            Source string  `json:"source"`
            BusinessInfo BusinessDetails  `json:"business_info"`
            Device DeviceDetails  `json:"device"`
            PersonalInfo UserPersonalInfoInDetails  `json:"personal_info"`
            Mcc string  `json:"mcc"`
            Aggregator string  `json:"aggregator"`
            MarketplaceInfo MarketplaceInfo  `json:"marketplace_info"`
         
    }
    
    // OnboardSummary ...
    type OnboardSummary struct {

        
            Status bool  `json:"status"`
            Session map[string]interface{}  `json:"session"`
            RedirectURL string  `json:"redirect_url"`
         
    }
    
    // CustomerOnboardingResponse ...
    type CustomerOnboardingResponse struct {

        
            Data OnboardSummary  `json:"data"`
            Success bool  `json:"success"`
         
    }
    

    
    // OrderStatuses ...
    type OrderStatuses struct {

        
            Value float64  `json:"value"`
            Display string  `json:"display"`
            IsSelected bool  `json:"is_selected"`
         
    }
    
    // OrderFilters ...
    type OrderFilters struct {

        
            Statuses []OrderStatuses  `json:"statuses"`
         
    }
    
    // OrderPage ...
    type OrderPage struct {

        
            HasNext bool  `json:"has_next"`
            ItemTotal float64  `json:"item_total"`
            Type string  `json:"type"`
            Current float64  `json:"current"`
            Size float64  `json:"size"`
         
    }
    
    // FulfillingStore ...
    type FulfillingStore struct {

        
            CompanyID float64  `json:"company_id"`
            ID float64  `json:"id"`
            Code string  `json:"code"`
            Name string  `json:"name"`
            CompanyName string  `json:"company_name"`
         
    }
    
    // Invoice ...
    type Invoice struct {

        
            InvoiceURL string  `json:"invoice_url"`
            LabelURL string  `json:"label_url"`
            UpdatedDate string  `json:"updated_date"`
         
    }
    
    // ShipmentPayment ...
    type ShipmentPayment struct {

        
            DisplayName string  `json:"display_name"`
            Status string  `json:"status"`
            Mode string  `json:"mode"`
            Mop string  `json:"mop"`
            Logo string  `json:"logo"`
            PaymentMode string  `json:"payment_mode"`
         
    }
    
    // ShipmentTotalDetails ...
    type ShipmentTotalDetails struct {

        
            Pieces float64  `json:"pieces"`
            Sizes float64  `json:"sizes"`
            TotalPrice float64  `json:"total_price"`
         
    }
    
    // FulfillingCompany ...
    type FulfillingCompany struct {

        
            ID float64  `json:"id"`
            Name string  `json:"name"`
         
    }
    
    // DeliveryAddress ...
    type DeliveryAddress struct {

        
            Landmark string  `json:"landmark"`
            CountryIsoCode string  `json:"country_iso_code"`
            Version string  `json:"version"`
            Email string  `json:"email"`
            AddressCategory string  `json:"address_category"`
            AddressType string  `json:"address_type"`
            Longitude float64  `json:"longitude"`
            UpdatedAt string  `json:"updated_at"`
            Pincode string  `json:"pincode"`
            Country string  `json:"country"`
            Phone string  `json:"phone"`
            Address1 string  `json:"address1"`
            Address2 string  `json:"address2"`
            Latitude float64  `json:"latitude"`
            City string  `json:"city"`
            State string  `json:"state"`
            ContactPerson string  `json:"contact_person"`
            Address string  `json:"address"`
            Area string  `json:"area"`
            CreatedAt string  `json:"created_at"`
            Name string  `json:"name"`
            CountryPhoneCode string  `json:"country_phone_code"`
         
    }
    
    // TimeStampData ...
    type TimeStampData struct {

        
            Max string  `json:"max"`
            Min string  `json:"min"`
         
    }
    
    // Promise ...
    type Promise struct {

        
            Timestamp TimeStampData  `json:"timestamp"`
            ShowPromise bool  `json:"show_promise"`
         
    }
    
    // Prices ...
    type Prices struct {

        
            PriceEffective float64  `json:"price_effective"`
            CashbackApplied float64  `json:"cashback_applied"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            Discount float64  `json:"discount"`
            FyndCredits float64  `json:"fynd_credits"`
            PriceMarked float64  `json:"price_marked"`
            CouponValue float64  `json:"coupon_value"`
            ValueOfGood float64  `json:"value_of_good"`
            TransferPrice float64  `json:"transfer_price"`
            AmountPaid float64  `json:"amount_paid"`
            CurrencySymbol string  `json:"currency_symbol"`
            PromotionEffectiveDiscount float64  `json:"promotion_effective_discount"`
            CodCharges float64  `json:"cod_charges"`
            RefundCredit float64  `json:"refund_credit"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            AmountPaidRoundoff float64  `json:"amount_paid_roundoff"`
            RefundAmount float64  `json:"refund_amount"`
            Cashback float64  `json:"cashback"`
            CurrencyCode string  `json:"currency_code"`
            AddedToFyndCash bool  `json:"added_to_fynd_cash"`
            DeliveryCharge float64  `json:"delivery_charge"`
         
    }
    
    // ShipmentUserInfo ...
    type ShipmentUserInfo struct {

        
            LastName string  `json:"last_name"`
            FirstName string  `json:"first_name"`
            Mobile string  `json:"mobile"`
            Gender string  `json:"gender"`
         
    }
    
    // CurrentStatus ...
    type CurrentStatus struct {

        
            UpdatedAt string  `json:"updated_at"`
            Status string  `json:"status"`
            CreatedAt string  `json:"created_at"`
            JourneyType string  `json:"journey_type"`
            Name string  `json:"name"`
         
    }
    
    // ItemBrand ...
    type ItemBrand struct {

        
            Name string  `json:"name"`
            Logo string  `json:"logo"`
         
    }
    
    // Item ...
    type Item struct {

        
            Code string  `json:"code"`
            ID float64  `json:"id"`
            Brand ItemBrand  `json:"brand"`
            SellerIdentifier string  `json:"seller_identifier"`
            Size string  `json:"size"`
            Name string  `json:"name"`
            Image []string  `json:"image"`
            SlugKey string  `json:"slug_key"`
         
    }
    
    // Identifiers ...
    type Identifiers struct {

        
            Ean string  `json:"ean"`
            SkuCode string  `json:"sku_code"`
         
    }
    
    // FinancialBreakup ...
    type FinancialBreakup struct {

        
            PriceEffective float64  `json:"price_effective"`
            CashbackApplied float64  `json:"cashback_applied"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            Discount float64  `json:"discount"`
            FyndCredits float64  `json:"fynd_credits"`
            PriceMarked float64  `json:"price_marked"`
            CouponValue float64  `json:"coupon_value"`
            ItemName string  `json:"item_name"`
            ValueOfGood float64  `json:"value_of_good"`
            TransferPrice float64  `json:"transfer_price"`
            GstFee float64  `json:"gst_fee"`
            HsnCode string  `json:"hsn_code"`
            AmountPaid float64  `json:"amount_paid"`
            Identifiers Identifiers  `json:"identifiers"`
            PromotionEffectiveDiscount float64  `json:"promotion_effective_discount"`
            CodCharges float64  `json:"cod_charges"`
            RefundCredit float64  `json:"refund_credit"`
            Size string  `json:"size"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            TotalUnits float64  `json:"total_units"`
            AmountPaidRoundoff float64  `json:"amount_paid_roundoff"`
            GstTag string  `json:"gst_tag"`
            RefundAmount float64  `json:"refund_amount"`
            Cashback float64  `json:"cashback"`
            AddedToFyndCash bool  `json:"added_to_fynd_cash"`
            DeliveryCharge float64  `json:"delivery_charge"`
         
    }
    
    // AppliedFreeArticles ...
    type AppliedFreeArticles struct {

        
            Quantity float64  `json:"quantity"`
            ParentItemIdentifier string  `json:"parent_item_identifier"`
            FreeGiftItemDetails map[string]interface{}  `json:"free_gift_item_details"`
            ArticleID string  `json:"article_id"`
         
    }
    
    // AppliedPromos ...
    type AppliedPromos struct {

        
            Amount float64  `json:"amount"`
            PromotionType string  `json:"promotion_type"`
            MrpPromotion bool  `json:"mrp_promotion"`
            PromotionName string  `json:"promotion_name"`
            ArticleQuantity float64  `json:"article_quantity"`
            AppliedFreeArticles []AppliedFreeArticles  `json:"applied_free_articles"`
            PromoID string  `json:"promo_id"`
         
    }
    
    // Bags ...
    type Bags struct {

        
            CurrentStatus CurrentStatus  `json:"current_status"`
            DeliveryDate string  `json:"delivery_date"`
            CanReturn bool  `json:"can_return"`
            ID float64  `json:"id"`
            ParentPromoBags map[string]interface{}  `json:"parent_promo_bags"`
            LineNumber float64  `json:"line_number"`
            Item Item  `json:"item"`
            Meta map[string]interface{}  `json:"meta"`
            SellerIdentifier string  `json:"seller_identifier"`
            FinancialBreakup []FinancialBreakup  `json:"financial_breakup"`
            AppliedPromos []AppliedPromos  `json:"applied_promos"`
            CanCancel bool  `json:"can_cancel"`
            CurrencySymbol string  `json:"currency_symbol"`
            CurrencyCode string  `json:"currency_code"`
            Quantity float64  `json:"quantity"`
            Prices Prices  `json:"prices"`
            ReturnableDate string  `json:"returnable_date"`
         
    }
    
    // NestedTrackingDetails ...
    type NestedTrackingDetails struct {

        
            Time string  `json:"time"`
            IsPassed bool  `json:"is_passed"`
            IsCurrent bool  `json:"is_current"`
            Status string  `json:"status"`
         
    }
    
    // TrackingDetails ...
    type TrackingDetails struct {

        
            Time string  `json:"time"`
            Status string  `json:"status"`
            TrackingDetails []NestedTrackingDetails  `json:"tracking_details"`
            IsPassed bool  `json:"is_passed"`
            IsCurrent bool  `json:"is_current"`
         
    }
    
    // ShipmentStatus ...
    type ShipmentStatus struct {

        
            Title string  `json:"title"`
            HexCode string  `json:"hex_code"`
         
    }
    
    // BreakupValues ...
    type BreakupValues struct {

        
            CurrencySymbol string  `json:"currency_symbol"`
            CurrencyCode string  `json:"currency_code"`
            Value float64  `json:"value"`
            Name string  `json:"name"`
            Display string  `json:"display"`
         
    }
    
    // Shipments ...
    type Shipments struct {

        
            TrakingNo string  `json:"traking_no"`
            FulfillingStore FulfillingStore  `json:"fulfilling_store"`
            Invoice Invoice  `json:"invoice"`
            OrderID string  `json:"order_id"`
            Payment ShipmentPayment  `json:"payment"`
            TotalDetails ShipmentTotalDetails  `json:"total_details"`
            OrderType string  `json:"order_type"`
            ReturnMeta map[string]interface{}  `json:"return_meta"`
            FulfillingCompany FulfillingCompany  `json:"fulfilling_company"`
            DeliveryAddress DeliveryAddress  `json:"delivery_address"`
            ShipmentID string  `json:"shipment_id"`
            SizeInfo map[string]interface{}  `json:"size_info"`
            ShowDownloadInvoice bool  `json:"show_download_invoice"`
            TotalBags float64  `json:"total_bags"`
            BeneficiaryDetails bool  `json:"beneficiary_details"`
            Promise Promise  `json:"promise"`
            RefundDetails map[string]interface{}  `json:"refund_details"`
            Comment string  `json:"comment"`
            ShipmentCreatedAt string  `json:"shipment_created_at"`
            CanBreak map[string]interface{}  `json:"can_break"`
            CanCancel bool  `json:"can_cancel"`
            TrackURL string  `json:"track_url"`
            DpName string  `json:"dp_name"`
            Prices Prices  `json:"prices"`
            UserInfo ShipmentUserInfo  `json:"user_info"`
            Bags []Bags  `json:"bags"`
            DeliveryDate string  `json:"delivery_date"`
            CanReturn bool  `json:"can_return"`
            AwbNo string  `json:"awb_no"`
            TrackingDetails []TrackingDetails  `json:"tracking_details"`
            CustomMeta []map[string]interface{}  `json:"custom_meta"`
            ShowTrackLink bool  `json:"show_track_link"`
            ShipmentStatus ShipmentStatus  `json:"shipment_status"`
            NeedHelpURL string  `json:"need_help_url"`
            BreakupValues []BreakupValues  `json:"breakup_values"`
            ReturnableDate string  `json:"returnable_date"`
         
    }
    
    // UserInfo ...
    type UserInfo struct {

        
            Email string  `json:"email"`
            Name string  `json:"name"`
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

        
            SellerID float64  `json:"seller_id"`
            StoreID float64  `json:"store_id"`
            ArticleAssignment BagsForReorderArticleAssignment  `json:"article_assignment"`
            ItemID float64  `json:"item_id"`
            Quantity float64  `json:"quantity"`
            ItemSize string  `json:"item_size"`
         
    }
    
    // OrderSchema ...
    type OrderSchema struct {

        
            OrderCreatedTime string  `json:"order_created_time"`
            Shipments []Shipments  `json:"shipments"`
            OrderID string  `json:"order_id"`
            BreakupValues []BreakupValues  `json:"breakup_values"`
            UserInfo UserInfo  `json:"user_info"`
            TotalShipmentsInOrder float64  `json:"total_shipments_in_order"`
            BagsForReorder []BagsForReorder  `json:"bags_for_reorder"`
         
    }
    
    // OrderList ...
    type OrderList struct {

        
            Filters OrderFilters  `json:"filters"`
            Page OrderPage  `json:"page"`
            Items []OrderSchema  `json:"items"`
         
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

        
            PresignedURL string  `json:"presigned_url"`
            PresignedType string  `json:"presigned_type"`
            ShipmentID string  `json:"shipment_id"`
            Success bool  `json:"success"`
         
    }
    
    // Track ...
    type Track struct {

        
            UpdatedAt string  `json:"updated_at"`
            Status string  `json:"status"`
            LastLocationRecievedAt string  `json:"last_location_recieved_at"`
            UpdatedTime string  `json:"updated_time"`
            Reason string  `json:"reason"`
            Awb string  `json:"awb"`
            ShipmentType string  `json:"shipment_type"`
            AccountName string  `json:"account_name"`
         
    }
    
    // ShipmentTrack ...
    type ShipmentTrack struct {

        
            Results []Track  `json:"results"`
         
    }
    
    // CustomerDetailsResponse ...
    type CustomerDetailsResponse struct {

        
            Country string  `json:"country"`
            Phone string  `json:"phone"`
            ShipmentID string  `json:"shipment_id"`
            OrderID string  `json:"order_id"`
            Name string  `json:"name"`
         
    }
    
    // SendOtpToCustomerResponse ...
    type SendOtpToCustomerResponse struct {

        
            ResendTimer float64  `json:"resend_timer"`
            Message string  `json:"message"`
            RequestID string  `json:"request_id"`
            Success bool  `json:"success"`
         
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

        
            ID float64  `json:"id"`
            DisplayName string  `json:"display_name"`
         
    }
    
    // BagReasonMeta ...
    type BagReasonMeta struct {

        
            ShowTextArea bool  `json:"show_text_area"`
         
    }
    
    // BagReasons ...
    type BagReasons struct {

        
            QcType []string  `json:"qc_type"`
            DisplayName string  `json:"display_name"`
            QuestionSet []QuestionSet  `json:"question_set"`
            ID float64  `json:"id"`
            Meta BagReasonMeta  `json:"meta"`
            Reasons []BagReasons  `json:"reasons"`
         
    }
    
    // ShipmentBagReasons ...
    type ShipmentBagReasons struct {

        
            Reasons []BagReasons  `json:"reasons"`
            Success bool  `json:"success"`
         
    }
    
    // ShipmentReason ...
    type ShipmentReason struct {

        
            ReasonID float64  `json:"reason_id"`
            ShowTextArea bool  `json:"show_text_area"`
            FeedbackType string  `json:"feedback_type"`
            Priority float64  `json:"priority"`
            ReasonText string  `json:"reason_text"`
            Flow string  `json:"flow"`
         
    }
    
    // ShipmentReasons ...
    type ShipmentReasons struct {

        
            Reasons []ShipmentReason  `json:"reasons"`
         
    }
    
    // EntitiesDataUpdates ...
    type EntitiesDataUpdates struct {

        
            Data map[string]interface{}  `json:"data"`
            Filters []map[string]interface{}  `json:"filters"`
         
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
    
    // DataUpdates ...
    type DataUpdates struct {

        
            Entities []EntitiesDataUpdates  `json:"entities"`
            Products []ProductsDataUpdates  `json:"products"`
         
    }
    
    // Products ...
    type Products struct {

        
            Quantity float64  `json:"quantity"`
            Identifier string  `json:"identifier"`
            LineNumber float64  `json:"line_number"`
         
    }
    
    // EntityReasonData ...
    type EntityReasonData struct {

        
            ReasonText string  `json:"reason_text"`
            ReasonID float64  `json:"reason_id"`
         
    }
    
    // EntitiesReasons ...
    type EntitiesReasons struct {

        
            Data EntityReasonData  `json:"data"`
            Filters []map[string]interface{}  `json:"filters"`
         
    }
    
    // ProductsReasonsData ...
    type ProductsReasonsData struct {

        
            ReasonText string  `json:"reason_text"`
            ReasonID float64  `json:"reason_id"`
         
    }
    
    // ProductsReasonsFilters ...
    type ProductsReasonsFilters struct {

        
            Quantity float64  `json:"quantity"`
            Identifier string  `json:"identifier"`
            LineNumber float64  `json:"line_number"`
         
    }
    
    // ProductsReasons ...
    type ProductsReasons struct {

        
            Data ProductsReasonsData  `json:"data"`
            Filters []ProductsReasonsFilters  `json:"filters"`
         
    }
    
    // ReasonsData ...
    type ReasonsData struct {

        
            Entities []EntitiesReasons  `json:"entities"`
            Products []ProductsReasons  `json:"products"`
         
    }
    
    // ShipmentsRequest ...
    type ShipmentsRequest struct {

        
            DataUpdates DataUpdates  `json:"data_updates"`
            Products []Products  `json:"products"`
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
            LockAfterTransition bool  `json:"lock_after_transition"`
            Statuses []StatuesRequest  `json:"statuses"`
            ForceTransition bool  `json:"force_transition"`
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
            Exception string  `json:"exception"`
            StackTrace string  `json:"stack_trace"`
            Code string  `json:"code"`
            Message string  `json:"message"`
         
    }
    
    // ProductStatus ...
    type ProductStatus struct {

        
            Value string  `json:"value"`
            CreatedAt string  `json:"created_at"`
            Title string  `json:"title"`
            HexCode string  `json:"hex_code"`
         
    }
    
    // ProductBrand ...
    type ProductBrand struct {

        
            ID float64  `json:"id"`
            BrandName string  `json:"brand_name"`
         
    }
    
    // Coupon ...
    type Coupon struct {

        
            CouponType string  `json:"coupon_type"`
            ID float64  `json:"id"`
            PayableCategory string  `json:"payable_category"`
            Value float64  `json:"value"`
            Code string  `json:"code"`
         
    }
    
    // Product ...
    type Product struct {

        
            BagStatus ProductStatus  `json:"bag_status"`
            DeliveryDate string  `json:"delivery_date"`
            CanReturn bool  `json:"can_return"`
            LineNumber float64  `json:"line_number"`
            ParentPromoBags map[string]interface{}  `json:"parent_promo_bags"`
            DocketNumber string  `json:"docket_number"`
            Brand ProductBrand  `json:"brand"`
            Item Item  `json:"item"`
            SellerIdentifier string  `json:"seller_identifier"`
            AppliedPromos []AppliedPromos  `json:"applied_promos"`
            CanCancel bool  `json:"can_cancel"`
            OrderID string  `json:"order_id"`
            Payment ShipmentPayment  `json:"payment"`
            Quantity float64  `json:"quantity"`
            Coupon Coupon  `json:"coupon"`
            ReturnableDate string  `json:"returnable_date"`
         
    }
    
    // ProductListResponse ...
    type ProductListResponse struct {

        
            Filters OrderFilters  `json:"filters"`
            Page OrderPage  `json:"page"`
            Message string  `json:"message"`
            Success bool  `json:"success"`
            Items []Product  `json:"items"`
         
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
    

    
    // DisplayBreakupSchema ...
    type DisplayBreakupSchema struct {

        
            Key string  `json:"key"`
            CurrencySymbol string  `json:"currency_symbol"`
            Display string  `json:"display"`
            Message []string  `json:"message"`
            CurrencyCode string  `json:"currency_code"`
            Value float64  `json:"value"`
         
    }
    
    // RawBreakupSchema ...
    type RawBreakupSchema struct {

        
            DeliveryCharge float64  `json:"delivery_charge"`
            Vog float64  `json:"vog"`
            Discount float64  `json:"discount"`
            YouSaved float64  `json:"you_saved"`
            MrpTotal float64  `json:"mrp_total"`
            FyndCash float64  `json:"fynd_cash"`
            GiftCard float64  `json:"gift_card"`
            Coupon float64  `json:"coupon"`
            CodCharge float64  `json:"cod_charge"`
            Total float64  `json:"total"`
            ConvenienceFee float64  `json:"convenience_fee"`
            GstCharges float64  `json:"gst_charges"`
            Subtotal float64  `json:"subtotal"`
         
    }
    
    // CouponBreakupSchema ...
    type CouponBreakupSchema struct {

        
            Type string  `json:"type"`
            IsApplied bool  `json:"is_applied"`
            CouponValue float64  `json:"coupon_value"`
            UID string  `json:"uid"`
            Title string  `json:"title"`
            MaxDiscountValue float64  `json:"max_discount_value"`
            MinimumCartValue float64  `json:"minimum_cart_value"`
            CouponType string  `json:"coupon_type"`
            Code string  `json:"code"`
            Message string  `json:"message"`
            Value float64  `json:"value"`
            SubTitle string  `json:"sub_title"`
            Description string  `json:"description"`
         
    }
    
    // LoyaltyPoints ...
    type LoyaltyPoints struct {

        
            IsApplied bool  `json:"is_applied"`
            Applicable float64  `json:"applicable"`
            Description string  `json:"description"`
            Total float64  `json:"total"`
         
    }
    
    // CartBreakupSchema ...
    type CartBreakupSchema struct {

        
            Display []DisplayBreakupSchema  `json:"display"`
            Raw RawBreakupSchema  `json:"raw"`
            Coupon CouponBreakupSchema  `json:"coupon"`
            LoyaltyPoints LoyaltyPoints  `json:"loyalty_points"`
         
    }
    
    // BaseInfo ...
    type BaseInfo struct {

        
            Name string  `json:"name"`
            UID float64  `json:"uid"`
         
    }
    
    // BasePriceSchema ...
    type BasePriceSchema struct {

        
            CurrencyCode string  `json:"currency_code"`
            Marked float64  `json:"marked"`
            CurrencySymbol string  `json:"currency_symbol"`
            Effective float64  `json:"effective"`
         
    }
    
    // ArticlePriceInfo ...
    type ArticlePriceInfo struct {

        
            Base BasePriceSchema  `json:"base"`
            Converted BasePriceSchema  `json:"converted"`
         
    }
    
    // ProductArticle ...
    type ProductArticle struct {

        
            Type string  `json:"type"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            Seller BaseInfo  `json:"seller"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            Identifier map[string]interface{}  `json:"identifier"`
            UID string  `json:"uid"`
            CartItemMeta map[string]interface{}  `json:"cart_item_meta"`
            ProductGroupTags []string  `json:"product_group_tags"`
            Size string  `json:"size"`
            Price ArticlePriceInfo  `json:"price"`
            GiftCard map[string]interface{}  `json:"gift_card"`
            SellerIdentifier string  `json:"seller_identifier"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            IsGiftVisible bool  `json:"is_gift_visible"`
            Store BaseInfo  `json:"store"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // CartProductIdentifer ...
    type CartProductIdentifer struct {

        
            Identifier string  `json:"identifier"`
         
    }
    
    // ProductPrice ...
    type ProductPrice struct {

        
            Marked float64  `json:"marked"`
            CurrencySymbol string  `json:"currency_symbol"`
            AddOn float64  `json:"add_on"`
            Selling float64  `json:"selling"`
            CurrencyCode string  `json:"currency_code"`
            Effective float64  `json:"effective"`
         
    }
    
    // ProductPriceInfo ...
    type ProductPriceInfo struct {

        
            Base ProductPrice  `json:"base"`
            Converted ProductPrice  `json:"converted"`
         
    }
    
    // CouponDetails ...
    type CouponDetails struct {

        
            DiscountTotalQuantity float64  `json:"discount_total_quantity"`
            Code string  `json:"code"`
            DiscountSingleQuantity float64  `json:"discount_single_quantity"`
         
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
    
    // ProductImage ...
    type ProductImage struct {

        
            SecureURL string  `json:"secure_url"`
            AspectRatio string  `json:"aspect_ratio"`
            URL string  `json:"url"`
         
    }
    
    // Tags ...
    type Tags struct {

        
            Tags map[string]interface{}  `json:"tags"`
         
    }
    
    // NetQuantity ...
    type NetQuantity struct {

        
            Unit string  `json:"unit"`
            Value string  `json:"value"`
         
    }
    
    // CategoryInfo ...
    type CategoryInfo struct {

        
            Name string  `json:"name"`
            UID float64  `json:"uid"`
         
    }
    
    // CartProduct ...
    type CartProduct struct {

        
            Type string  `json:"type"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
            Brand BaseInfo  `json:"brand"`
            ItemCode string  `json:"item_code"`
            Action ProductAction  `json:"action"`
            Tags []string  `json:"tags"`
            Images []ProductImage  `json:"images"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            TeaserTag Tags  `json:"teaser_tag"`
            NetQuantity NetQuantity  `json:"net_quantity"`
            Categories []CategoryInfo  `json:"categories"`
         
    }
    
    // ProductAvailabilitySize ...
    type ProductAvailabilitySize struct {

        
            Value string  `json:"value"`
            Display string  `json:"display"`
            IsAvailable bool  `json:"is_available"`
         
    }
    
    // ProductAvailability ...
    type ProductAvailability struct {

        
            Sizes []string  `json:"sizes"`
            OtherStoreQuantity float64  `json:"other_store_quantity"`
            Deliverable bool  `json:"deliverable"`
            IsValid bool  `json:"is_valid"`
            AvailableSizes []ProductAvailabilitySize  `json:"available_sizes"`
            OutOfStock bool  `json:"out_of_stock"`
         
    }
    
    // PromoMeta ...
    type PromoMeta struct {

        
            Message string  `json:"message"`
         
    }
    
    // DiscountRulesAppSchema ...
    type DiscountRulesAppSchema struct {

        
            RawOffer map[string]interface{}  `json:"raw_offer"`
            MatchedBuyRules []string  `json:"matched_buy_rules"`
            Offer map[string]interface{}  `json:"offer"`
            ItemCriteria map[string]interface{}  `json:"item_criteria"`
         
    }
    
    // BuyRulesSchema ...
    type BuyRulesSchema struct {

        
            CartConditions map[string]interface{}  `json:"cart_conditions"`
            ItemCriteria map[string]interface{}  `json:"item_criteria"`
         
    }
    
    // FreeGiftItemSchema ...
    type FreeGiftItemSchema struct {

        
            ItemID float64  `json:"item_id"`
            ItemPriceDetails map[string]interface{}  `json:"item_price_details"`
            ItemName string  `json:"item_name"`
            ItemSlug string  `json:"item_slug"`
            ItemBrandName string  `json:"item_brand_name"`
            ItemImagesURL []string  `json:"item_images_url"`
         
    }
    
    // AppliedFreeArticlesSchema ...
    type AppliedFreeArticlesSchema struct {

        
            Quantity float64  `json:"quantity"`
            ArticleID string  `json:"article_id"`
            ParentItemIdentifier string  `json:"parent_item_identifier"`
            FreeGiftItemDetails FreeGiftItemSchema  `json:"free_gift_item_details"`
         
    }
    
    // Ownership ...
    type Ownership struct {

        
            PayableCategory string  `json:"payable_category"`
            PayableBy string  `json:"payable_by"`
         
    }
    
    // AppliedPromotion ...
    type AppliedPromotion struct {

        
            PromotionName string  `json:"promotion_name"`
            ArticleQuantity float64  `json:"article_quantity"`
            MrpPromotion bool  `json:"mrp_promotion"`
            OfferText string  `json:"offer_text"`
            DiscountRules []DiscountRulesAppSchema  `json:"discount_rules"`
            BuyRules []BuyRulesSchema  `json:"buy_rules"`
            AppliedFreeArticles []AppliedFreeArticlesSchema  `json:"applied_free_articles"`
            PromoID string  `json:"promo_id"`
            PromotionType string  `json:"promotion_type"`
            Ownership Ownership  `json:"ownership"`
            Amount float64  `json:"amount"`
            PromotionGroup string  `json:"promotion_group"`
         
    }
    
    // CartProductInfo ...
    type CartProductInfo struct {

        
            Article ProductArticle  `json:"article"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
            Price ProductPriceInfo  `json:"price"`
            Message string  `json:"message"`
            CouponMessage string  `json:"coupon_message"`
            Coupon CouponDetails  `json:"coupon"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            IsSet bool  `json:"is_set"`
            BulkOffer map[string]interface{}  `json:"bulk_offer"`
            Product CartProduct  `json:"product"`
            Quantity float64  `json:"quantity"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            Key string  `json:"key"`
            Availability ProductAvailability  `json:"availability"`
            PricePerUnit ProductPriceInfo  `json:"price_per_unit"`
            Discount string  `json:"discount"`
            PromoMeta PromoMeta  `json:"promo_meta"`
            CustomOrder map[string]interface{}  `json:"custom_order"`
            PromotionsApplied []AppliedPromotion  `json:"promotions_applied"`
            Moq map[string]interface{}  `json:"moq"`
         
    }
    
    // PaymentSelectionLock ...
    type PaymentSelectionLock struct {

        
            Enabled bool  `json:"enabled"`
            DefaultOptions string  `json:"default_options"`
            PaymentIdentifier string  `json:"payment_identifier"`
         
    }
    
    // CartCurrencySchema ...
    type CartCurrencySchema struct {

        
            Code string  `json:"code"`
            Symbol string  `json:"symbol"`
         
    }
    
    // CartDetailResponse ...
    type CartDetailResponse struct {

        
            LastModified string  `json:"last_modified"`
            BreakupValues CartBreakupSchema  `json:"breakup_values"`
            ID string  `json:"id"`
            Message string  `json:"message"`
            PanNo string  `json:"pan_no"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            PanConfig map[string]interface{}  `json:"pan_config"`
            Items []CartProductInfo  `json:"items"`
            CheckoutMode string  `json:"checkout_mode"`
            IsValid bool  `json:"is_valid"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            CouponText string  `json:"coupon_text"`
            Gstin string  `json:"gstin"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            Comment string  `json:"comment"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            AppliedPromoDetails []AppliedPromotion  `json:"applied_promo_details"`
            BuyNow bool  `json:"buy_now"`
            Currency CartCurrencySchema  `json:"currency"`
         
    }
    
    // AddProductCart ...
    type AddProductCart struct {

        
            ParentItemIdentifiers []map[string]string  `json:"parent_item_identifiers"`
            ItemID float64  `json:"item_id"`
            ProductGroupTags []string  `json:"product_group_tags"`
            SellerID float64  `json:"seller_id"`
            StoreID float64  `json:"store_id"`
            ArticleAssignment map[string]interface{}  `json:"article_assignment"`
            Display string  `json:"display"`
            ItemSize string  `json:"item_size"`
            Pos bool  `json:"pos"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            ArticleID string  `json:"article_id"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // AddCartRequest ...
    type AddCartRequest struct {

        
            Items []AddProductCart  `json:"items"`
            NewCart bool  `json:"new_cart"`
         
    }
    
    // AddCartDetailResponse ...
    type AddCartDetailResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
            Partial bool  `json:"partial"`
            Cart CartDetailResponse  `json:"cart"`
         
    }
    
    // UpdateProductCart ...
    type UpdateProductCart struct {

        
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
            ItemID float64  `json:"item_id"`
            ItemSize string  `json:"item_size"`
            ItemIndex float64  `json:"item_index"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            ArticleID string  `json:"article_id"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // UpdateCartRequest ...
    type UpdateCartRequest struct {

        
            Items []UpdateProductCart  `json:"items"`
            Operation string  `json:"operation"`
         
    }
    
    // UpdateCartDetailResponse ...
    type UpdateCartDetailResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
            Cart CartDetailResponse  `json:"cart"`
         
    }
    
    // CartItemCountResponse ...
    type CartItemCountResponse struct {

        
            UserCartItemsCount float64  `json:"user_cart_items_count"`
         
    }
    
    // Coupon ...
    type Coupon struct {

        
            IsApplied bool  `json:"is_applied"`
            CouponValue float64  `json:"coupon_value"`
            IsApplicable bool  `json:"is_applicable"`
            MinimumCartValue float64  `json:"minimum_cart_value"`
            Title string  `json:"title"`
            MaxDiscountValue float64  `json:"max_discount_value"`
            CouponType string  `json:"coupon_type"`
            CouponCode string  `json:"coupon_code"`
            Message string  `json:"message"`
            SubTitle string  `json:"sub_title"`
            ExpiresOn string  `json:"expires_on"`
            Description string  `json:"description"`
         
    }
    
    // PageCouponSchema ...
    type PageCouponSchema struct {

        
            HasNext bool  `json:"has_next"`
            Total float64  `json:"total"`
            HasPrevious bool  `json:"has_previous"`
            Current float64  `json:"current"`
            TotalItemCount float64  `json:"total_item_count"`
         
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

        
            Marked float64  `json:"marked"`
            CurrencySymbol string  `json:"currency_symbol"`
            BulkEffective float64  `json:"bulk_effective"`
            CurrencyCode string  `json:"currency_code"`
            Effective float64  `json:"effective"`
         
    }
    
    // OfferItemSchema ...
    type OfferItemSchema struct {

        
            Type string  `json:"type"`
            Margin float64  `json:"margin"`
            Price OfferPrice  `json:"price"`
            Best bool  `json:"best"`
            Total float64  `json:"total"`
            Quantity float64  `json:"quantity"`
            AutoApplied bool  `json:"auto_applied"`
         
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

        
            Name string  `json:"name"`
            UserID string  `json:"user_id"`
            ID string  `json:"id"`
            Email string  `json:"email"`
            AddressType string  `json:"address_type"`
            Address string  `json:"address"`
            Landmark string  `json:"landmark"`
            CountryPhoneCode string  `json:"country_phone_code"`
            AreaCode string  `json:"area_code"`
            IsActive bool  `json:"is_active"`
            Area string  `json:"area"`
            CheckoutMode string  `json:"checkout_mode"`
            CountryIsoCode string  `json:"country_iso_code"`
            CreatedByUserID string  `json:"created_by_user_id"`
            GeoLocation GeoLocation  `json:"geo_location"`
            Tags []string  `json:"tags"`
            CountryCode string  `json:"country_code"`
            Phone string  `json:"phone"`
            State string  `json:"state"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            IsDefaultAddress bool  `json:"is_default_address"`
            GoogleMapPoint map[string]interface{}  `json:"google_map_point"`
            AreaCodeSlug string  `json:"area_code_slug"`
            Meta map[string]interface{}  `json:"meta"`
            Country string  `json:"country"`
            City string  `json:"city"`
         
    }
    
    // GetAddressesResponse ...
    type GetAddressesResponse struct {

        
            PiiMasking bool  `json:"pii_masking"`
            Address []Address  `json:"address"`
         
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

        
            PaymentMode string  `json:"payment_mode"`
            PaymentIdentifier string  `json:"payment_identifier"`
            ID string  `json:"id"`
            MerchantCode string  `json:"merchant_code"`
            AddressID string  `json:"address_id"`
            AggregatorName string  `json:"aggregator_name"`
         
    }
    
    // CouponValidity ...
    type CouponValidity struct {

        
            Discount float64  `json:"discount"`
            Title string  `json:"title"`
            Valid bool  `json:"valid"`
            Code string  `json:"code"`
            DisplayMessageEn string  `json:"display_message_en"`
         
    }
    
    // PaymentCouponValidateSchema ...
    type PaymentCouponValidateSchema struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
            CouponValidity CouponValidity  `json:"coupon_validity"`
         
    }
    
    // ShipmentResponse ...
    type ShipmentResponse struct {

        
            Shipments float64  `json:"shipments"`
            FulfillmentID float64  `json:"fulfillment_id"`
            DpOptions map[string]interface{}  `json:"dp_options"`
            Promise ShipmentPromise  `json:"promise"`
            BoxType string  `json:"box_type"`
            Items []CartProductInfo  `json:"items"`
            FulfillmentType string  `json:"fulfillment_type"`
            OrderType string  `json:"order_type"`
            ShipmentType string  `json:"shipment_type"`
            DpID string  `json:"dp_id"`
         
    }
    
    // CartShipmentsResponse ...
    type CartShipmentsResponse struct {

        
            LastModified string  `json:"last_modified"`
            Error bool  `json:"error"`
            BreakupValues CartBreakupSchema  `json:"breakup_values"`
            ID string  `json:"id"`
            Gstin string  `json:"gstin"`
            CartID float64  `json:"cart_id"`
            Message string  `json:"message"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            Shipments []ShipmentResponse  `json:"shipments"`
            UID string  `json:"uid"`
            Comment string  `json:"comment"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            CheckoutMode string  `json:"checkout_mode"`
            IsValid bool  `json:"is_valid"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            BuyNow bool  `json:"buy_now"`
            CouponText string  `json:"coupon_text"`
            Currency CartCurrencySchema  `json:"currency"`
         
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
    
    // FilesSchema ...
    type FilesSchema struct {

        
            Values []string  `json:"values"`
            Key string  `json:"key"`
         
    }
    
    // CartCheckoutCustomMetaSchema ...
    type CartCheckoutCustomMetaSchema struct {

        
            Value string  `json:"value"`
            Key string  `json:"key"`
         
    }
    
    // StaffCheckoutSchema ...
    type StaffCheckoutSchema struct {

        
            FirstName string  `json:"first_name"`
            LastName string  `json:"last_name"`
            EmployeeCode string  `json:"employee_code"`
            ID string  `json:"_id"`
            User string  `json:"user"`
         
    }
    
    // CartPosCheckoutDetailRequest ...
    type CartPosCheckoutDetailRequest struct {

        
            ID string  `json:"id"`
            MerchantCode string  `json:"merchant_code"`
            PaymentAutoConfirm bool  `json:"payment_auto_confirm"`
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
            Aggregator string  `json:"aggregator"`
            PickAtStoreUID float64  `json:"pick_at_store_uid"`
            BillingAddressID string  `json:"billing_address_id"`
            AddressID string  `json:"address_id"`
            Files []FilesSchema  `json:"files"`
            OrderingStore float64  `json:"ordering_store"`
            CustomMeta []CartCheckoutCustomMetaSchema  `json:"custom_meta"`
            CallbackURL string  `json:"callback_url"`
            PaymentMode string  `json:"payment_mode"`
            PaymentIdentifier string  `json:"payment_identifier"`
            Pos bool  `json:"pos"`
            Staff StaffCheckoutSchema  `json:"staff"`
            Meta map[string]interface{}  `json:"meta"`
            PaymentParams map[string]interface{}  `json:"payment_params"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            OrderType string  `json:"order_type"`
         
    }
    
    // CheckCartSchema ...
    type CheckCartSchema struct {

        
            LastModified string  `json:"last_modified"`
            BreakupValues CartBreakupSchema  `json:"breakup_values"`
            ID string  `json:"id"`
            CartID float64  `json:"cart_id"`
            CodCharges float64  `json:"cod_charges"`
            CodMessage string  `json:"cod_message"`
            ErrorMessage string  `json:"error_message"`
            Message string  `json:"message"`
            CodAvailable bool  `json:"cod_available"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            OrderID string  `json:"order_id"`
            UserType string  `json:"user_type"`
            Items []CartProductInfo  `json:"items"`
            CheckoutMode string  `json:"checkout_mode"`
            IsValid bool  `json:"is_valid"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            CouponText string  `json:"coupon_text"`
            DeliveryCharges float64  `json:"delivery_charges"`
            Gstin string  `json:"gstin"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            StoreCode string  `json:"store_code"`
            Success bool  `json:"success"`
            DeliveryChargeOrderValue float64  `json:"delivery_charge_order_value"`
            UID string  `json:"uid"`
            Comment string  `json:"comment"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            BuyNow bool  `json:"buy_now"`
            StoreEmps []map[string]interface{}  `json:"store_emps"`
            Currency CartCurrencySchema  `json:"currency"`
         
    }
    
    // CartCheckoutResponseSchema ...
    type CartCheckoutResponseSchema struct {

        
            Success bool  `json:"success"`
            CallbackURL string  `json:"callback_url"`
            Data map[string]interface{}  `json:"data"`
            AppInterceptURL string  `json:"app_intercept_url"`
            OrderID string  `json:"order_id"`
            Cart CheckCartSchema  `json:"cart"`
            Message string  `json:"message"`
            PaymentConfirmURL string  `json:"payment_confirm_url"`
         
    }
    
    // GiftDetailSchema ...
    type GiftDetailSchema struct {

        
            GiftMessage string  `json:"gift_message"`
            IsGiftApplied bool  `json:"is_gift_applied"`
         
    }
    
    // ArticleGiftDetailSchema ...
    type ArticleGiftDetailSchema struct {

        
            ArticleID GiftDetailSchema  `json:"article_id"`
         
    }
    
    // CartMetaRequest ...
    type CartMetaRequest struct {

        
            Comment string  `json:"comment"`
            Gstin string  `json:"gstin"`
            GiftDetails ArticleGiftDetailSchema  `json:"gift_details"`
            PickUpCustomerDetails map[string]interface{}  `json:"pick_up_customer_details"`
            CheckoutMode string  `json:"checkout_mode"`
         
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

        
            Name string  `json:"name"`
            Landmark string  `json:"landmark"`
            UID float64  `json:"uid"`
            ID float64  `json:"id"`
            AreaCode string  `json:"area_code"`
            AddressType string  `json:"address_type"`
            Email string  `json:"email"`
            Phone string  `json:"phone"`
            Area string  `json:"area"`
            Pincode float64  `json:"pincode"`
            AreaCodeSlug string  `json:"area_code_slug"`
            State string  `json:"state"`
            StoreCode string  `json:"store_code"`
            Country string  `json:"country"`
            Address string  `json:"address"`
            City string  `json:"city"`
         
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

        
            Source map[string]interface{}  `json:"source"`
            CreatedOn string  `json:"created_on"`
            User map[string]interface{}  `json:"user"`
            Meta map[string]interface{}  `json:"meta"`
            Token string  `json:"token"`
         
    }
    
    // SharedCart ...
    type SharedCart struct {

        
            LastModified string  `json:"last_modified"`
            BreakupValues CartBreakupSchema  `json:"breakup_values"`
            ID string  `json:"id"`
            CartID float64  `json:"cart_id"`
            Message string  `json:"message"`
            SharedCartDetails SharedCartDetails  `json:"shared_cart_details"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            Items []CartProductInfo  `json:"items"`
            CheckoutMode string  `json:"checkout_mode"`
            IsValid bool  `json:"is_valid"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            CouponText string  `json:"coupon_text"`
            Gstin string  `json:"gstin"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            UID string  `json:"uid"`
            Comment string  `json:"comment"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            BuyNow bool  `json:"buy_now"`
            Currency CartCurrencySchema  `json:"currency"`
         
    }
    
    // SharedCartResponse ...
    type SharedCartResponse struct {

        
            Error string  `json:"error"`
            Cart SharedCart  `json:"cart"`
         
    }
    

    
    // PincodeErrorSchemaResponse ...
    type PincodeErrorSchemaResponse struct {

        
            Message string  `json:"message"`
            Type string  `json:"type"`
            Value string  `json:"value"`
         
    }
    
    // PincodeMetaResponse ...
    type PincodeMetaResponse struct {

        
            Zone string  `json:"zone"`
            InternalZoneID float64  `json:"internal_zone_id"`
         
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
    
    // PincodeParentsResponse ...
    type PincodeParentsResponse struct {

        
            SubType string  `json:"sub_type"`
            UID string  `json:"uid"`
            Name string  `json:"name"`
            DisplayName string  `json:"display_name"`
         
    }
    
    // PincodeDataResponse ...
    type PincodeDataResponse struct {

        
            Meta PincodeMetaResponse  `json:"meta"`
            LatLong PincodeLatLongData  `json:"lat_long"`
            SubType string  `json:"sub_type"`
            UID string  `json:"uid"`
            MetaCode CountryMetaResponse  `json:"meta_code"`
            Parents []PincodeParentsResponse  `json:"parents"`
            DisplayName string  `json:"display_name"`
            Name string  `json:"name"`
            Error PincodeErrorSchemaResponse  `json:"error"`
         
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

        
            ManufacturingTime float64  `json:"manufacturing_time"`
            Category TATCategoryRequest  `json:"category"`
            ManufacturingTimeUnit string  `json:"manufacturing_time_unit"`
         
    }
    
    // TATLocationDetailsRequest ...
    type TATLocationDetailsRequest struct {

        
            FulfillmentID float64  `json:"fulfillment_id"`
            FromPincode string  `json:"from_pincode"`
            Articles []TATArticlesRequest  `json:"articles"`
         
    }
    
    // TATViewRequest ...
    type TATViewRequest struct {

        
            LocationDetails []TATLocationDetailsRequest  `json:"location_details"`
            ToPincode string  `json:"to_pincode"`
            Journey string  `json:"journey"`
            Identifier string  `json:"identifier"`
            Action string  `json:"action"`
            Source string  `json:"source"`
         
    }
    
    // TATErrorSchemaResponse ...
    type TATErrorSchemaResponse struct {

        
            Message string  `json:"message"`
            Type string  `json:"type"`
            Value string  `json:"value"`
         
    }
    
    // TATTimestampResponse ...
    type TATTimestampResponse struct {

        
            Max float64  `json:"max"`
            Min float64  `json:"min"`
         
    }
    
    // TATFormattedResponse ...
    type TATFormattedResponse struct {

        
            Max string  `json:"max"`
            Min string  `json:"min"`
         
    }
    
    // TATPromiseResponse ...
    type TATPromiseResponse struct {

        
            Timestamp TATTimestampResponse  `json:"timestamp"`
            Formatted TATFormattedResponse  `json:"formatted"`
         
    }
    
    // TATArticlesResponse ...
    type TATArticlesResponse struct {

        
            ManufacturingTime float64  `json:"manufacturing_time"`
            ManufacturingTimeUnit string  `json:"manufacturing_time_unit"`
            ManufacturingTimeSeconds float64  `json:"_manufacturing_time_seconds"`
            IsCodAvailable bool  `json:"is_cod_available"`
            Category TATCategoryRequest  `json:"category"`
            Error TATErrorSchemaResponse  `json:"error"`
            Promise TATPromiseResponse  `json:"promise"`
         
    }
    
    // TATLocationDetailsResponse ...
    type TATLocationDetailsResponse struct {

        
            FulfillmentID float64  `json:"fulfillment_id"`
            FromPincode string  `json:"from_pincode"`
            Articles []TATArticlesResponse  `json:"articles"`
         
    }
    
    // TATViewResponse ...
    type TATViewResponse struct {

        
            LocationDetails []TATLocationDetailsResponse  `json:"location_details"`
            ToPincode string  `json:"to_pincode"`
            PaymentMode string  `json:"payment_mode"`
            StormbreakerUUID string  `json:"stormbreaker_uuid"`
            ToCity string  `json:"to_city"`
            Journey string  `json:"journey"`
            Identifier string  `json:"identifier"`
            Action string  `json:"action"`
            Source string  `json:"source"`
            Success bool  `json:"success"`
            Error TATErrorSchemaResponse  `json:"error"`
            RequestUUID string  `json:"request_uuid"`
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

        
            ParentID string  `json:"parent_id"`
            Meta CountryMetaResponse  `json:"meta"`
            SubType string  `json:"sub_type"`
            UID string  `json:"uid"`
            Type string  `json:"type"`
            DisplayName string  `json:"display_name"`
            Logistics LogisticsResponse  `json:"logistics"`
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
         
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

        
            ServiceabilityType string  `json:"serviceability_type"`
            Zones []string  `json:"zones"`
         
    }
    
    // ReAssignStoreRequest ...
    type ReAssignStoreRequest struct {

        
            ToPincode string  `json:"to_pincode"`
            Identifier string  `json:"identifier"`
            Articles []map[string]interface{}  `json:"articles"`
            IgnoredLocations []string  `json:"ignored_locations"`
            Configuration map[string]interface{}  `json:"configuration"`
         
    }
    
    // ReAssignStoreResponse ...
    type ReAssignStoreResponse struct {

        
            Success bool  `json:"success"`
            ToPincode string  `json:"to_pincode"`
            Error map[string]interface{}  `json:"error"`
            Articles []map[string]interface{}  `json:"articles"`
         
    }
    
