package application



    
    // ApplicationItemMOQ ...
    type ApplicationItemMOQ struct {

        
            IncrementUnit float64  `json:"increment_unit"`
            Maximum float64  `json:"maximum"`
            Minimum float64  `json:"minimum"`
         
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
    
    // ApplicationItemSEO ...
    type ApplicationItemSEO struct {

        
            Title interface{}  `json:"title"`
            Description interface{}  `json:"description"`
         
    }
    
    // Meta ...
    type Meta struct {

        
            Source string  `json:"source"`
         
    }
    
    // Media ...
    type Media struct {

        
            URL string  `json:"url"`
            Alt string  `json:"alt"`
            Meta Meta  `json:"meta"`
            Type string  `json:"type"`
         
    }
    
    // ProductBrand ...
    type ProductBrand struct {

        
            Action ProductListingAction  `json:"action"`
            Logo Media  `json:"logo"`
            Name string  `json:"name"`
            Description string  `json:"description"`
            UID float64  `json:"uid"`
         
    }
    
    // ProductCategoryMap ...
    type ProductCategoryMap struct {

        
            L1 ProductBrand  `json:"l1"`
            L3 ProductBrand  `json:"l3"`
            L2 ProductBrand  `json:"l2"`
         
    }
    
    // CustomMetaFields ...
    type CustomMetaFields struct {

        
            Value string  `json:"value"`
            Key string  `json:"key"`
         
    }
    
    // NetQuantity ...
    type NetQuantity struct {

        
            Value float64  `json:"value"`
            Unit interface{}  `json:"unit"`
         
    }
    
    // Price ...
    type Price struct {

        
            CurrencyCode string  `json:"currency_code"`
            Min float64  `json:"min"`
            CurrencySymbol string  `json:"currency_symbol"`
            Max float64  `json:"max"`
         
    }
    
    // ProductListingPrice ...
    type ProductListingPrice struct {

        
            Marked Price  `json:"marked"`
            Effective Price  `json:"effective"`
         
    }
    
    // ProductDetail ...
    type ProductDetail struct {

        
            Discount string  `json:"discount"`
            ImageNature string  `json:"image_nature"`
            Type string  `json:"type"`
            ShortDescription string  `json:"short_description"`
            Color string  `json:"color"`
            Moq ApplicationItemMOQ  `json:"moq"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Action ProductListingAction  `json:"action"`
            HasVariant bool  `json:"has_variant"`
            Slug string  `json:"slug"`
            Tags []string  `json:"tags"`
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
            ItemCode string  `json:"item_code"`
            Seo ApplicationItemSEO  `json:"seo"`
            Similars []string  `json:"similars"`
            CategoryMap ProductCategoryMap  `json:"category_map"`
            ItemType string  `json:"item_type"`
            Attributes map[string]interface{}  `json:"attributes"`
            UID float64  `json:"uid"`
            Rating float64  `json:"rating"`
            TeaserTag string  `json:"teaser_tag"`
            Name string  `json:"name"`
            Categories []ProductBrand  `json:"categories"`
            ProductOnlineDate string  `json:"product_online_date"`
            Highlights []string  `json:"highlights"`
            IsDependent bool  `json:"is_dependent"`
            CustomMeta []CustomMetaFields  `json:"_custom_meta"`
            Tryouts []string  `json:"tryouts"`
            Medias []Media  `json:"medias"`
            NetQuantity NetQuantity  `json:"net_quantity"`
            Description string  `json:"description"`
            RatingCount float64  `json:"rating_count"`
            Price ProductListingPrice  `json:"price"`
            Brand ProductBrand  `json:"brand"`
         
    }
    
    // ErrorResponse ...
    type ErrorResponse struct {

        
            Error string  `json:"error"`
         
    }
    
    // ColumnHeader ...
    type ColumnHeader struct {

        
            Value string  `json:"value"`
            Convertable bool  `json:"convertable"`
         
    }
    
    // ColumnHeaders ...
    type ColumnHeaders struct {

        
            Col1 ColumnHeader  `json:"col_1"`
            Col3 ColumnHeader  `json:"col_3"`
            Col6 ColumnHeader  `json:"col_6"`
            Col2 ColumnHeader  `json:"col_2"`
            Col4 ColumnHeader  `json:"col_4"`
            Col5 ColumnHeader  `json:"col_5"`
         
    }
    
    // SizeChartValues ...
    type SizeChartValues struct {

        
            Col1 string  `json:"col_1"`
            Col3 string  `json:"col_3"`
            Col6 string  `json:"col_6"`
            Col2 string  `json:"col_2"`
            Col4 string  `json:"col_4"`
            Col5 string  `json:"col_5"`
         
    }
    
    // SizeChart ...
    type SizeChart struct {

        
            SizeTip string  `json:"size_tip"`
            Image string  `json:"image"`
            Headers ColumnHeaders  `json:"headers"`
            Unit string  `json:"unit"`
            Description string  `json:"description"`
            Title string  `json:"title"`
            Sizes []SizeChartValues  `json:"sizes"`
         
    }
    
    // Dimension ...
    type Dimension struct {

        
            Width float64  `json:"width"`
            Height float64  `json:"height"`
            Length float64  `json:"length"`
            IsDefault bool  `json:"is_default"`
            Unit string  `json:"unit"`
         
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
            Quantity float64  `json:"quantity"`
            SellerIdentifiers []string  `json:"seller_identifiers"`
            Dimension Dimension  `json:"dimension"`
            Weight Weight  `json:"weight"`
            IsAvailable bool  `json:"is_available"`
            Display string  `json:"display"`
         
    }
    
    // ProductSizeStores ...
    type ProductSizeStores struct {

        
            Count float64  `json:"count"`
         
    }
    
    // ProductSizes ...
    type ProductSizes struct {

        
            SizeChart SizeChart  `json:"size_chart"`
            Discount string  `json:"discount"`
            Sizes []ProductSize  `json:"sizes"`
            Sellable bool  `json:"sellable"`
            Price ProductListingPrice  `json:"price"`
            Stores ProductSizeStores  `json:"stores"`
            MultiSize bool  `json:"multi_size"`
         
    }
    
    // AttributeDetail ...
    type AttributeDetail struct {

        
            Logo string  `json:"logo"`
            Description string  `json:"description"`
            Key string  `json:"key"`
            Display string  `json:"display"`
         
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
            Items []ProductDetail  `json:"items"`
            Subtitle string  `json:"subtitle"`
         
    }
    
    // ProductFrequentlyComparedSimilarResponse ...
    type ProductFrequentlyComparedSimilarResponse struct {

        
            Similars ProductCompareResponse  `json:"similars"`
         
    }
    
    // ProductVariantItemResponse ...
    type ProductVariantItemResponse struct {

        
            Value string  `json:"value"`
            ColorName string  `json:"color_name"`
            Action ProductListingAction  `json:"action"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            CustomMeta []CustomMetaFields  `json:"_custom_meta"`
            Medias []Media  `json:"medias"`
            Color string  `json:"color"`
            IsAvailable bool  `json:"is_available"`
            UID float64  `json:"uid"`
         
    }
    
    // ProductVariantResponse ...
    type ProductVariantResponse struct {

        
            DisplayType string  `json:"display_type"`
            Items []ProductVariantItemResponse  `json:"items"`
            Header string  `json:"header"`
            Key string  `json:"key"`
         
    }
    
    // ProductVariantsResponse ...
    type ProductVariantsResponse struct {

        
            Variants []ProductVariantResponse  `json:"variants"`
         
    }
    
    // StoreDetail ...
    type StoreDetail struct {

        
            City string  `json:"city"`
            ID float64  `json:"id"`
            Code string  `json:"code"`
            Name string  `json:"name"`
         
    }
    
    // Seller ...
    type Seller struct {

        
            UID float64  `json:"uid"`
            Count float64  `json:"count"`
            Name string  `json:"name"`
         
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
    
    // ProductStockStatusItem ...
    type ProductStockStatusItem struct {

        
            Quantity float64  `json:"quantity"`
            Store StoreDetail  `json:"store"`
            Identifier map[string]interface{}  `json:"identifier"`
            Seller Seller  `json:"seller"`
            ItemID float64  `json:"item_id"`
            Company CompanyDetail  `json:"company"`
            Size string  `json:"size"`
            Price ProductStockPrice  `json:"price"`
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

        
            Discount string  `json:"discount"`
            ImageNature string  `json:"image_nature"`
            Type string  `json:"type"`
            ShortDescription string  `json:"short_description"`
            Sizes []string  `json:"sizes"`
            Color string  `json:"color"`
            Moq ApplicationItemMOQ  `json:"moq"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Action ProductListingAction  `json:"action"`
            HasVariant bool  `json:"has_variant"`
            Slug string  `json:"slug"`
            Tags []string  `json:"tags"`
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
            ItemCode string  `json:"item_code"`
            Identifiers []string  `json:"identifiers"`
            Sellable bool  `json:"sellable"`
            Seo ApplicationItemSEO  `json:"seo"`
            Similars []string  `json:"similars"`
            CategoryMap ProductCategoryMap  `json:"category_map"`
            ItemType string  `json:"item_type"`
            Attributes map[string]interface{}  `json:"attributes"`
            UID float64  `json:"uid"`
            Variants []ProductVariantListingResponse  `json:"variants"`
            Rating float64  `json:"rating"`
            TeaserTag string  `json:"teaser_tag"`
            Name string  `json:"name"`
            Categories []ProductBrand  `json:"categories"`
            ProductOnlineDate string  `json:"product_online_date"`
            Highlights []string  `json:"highlights"`
            IsDependent bool  `json:"is_dependent"`
            CustomMeta []CustomMetaFields  `json:"_custom_meta"`
            Tryouts []string  `json:"tryouts"`
            Medias []Media  `json:"medias"`
            NetQuantity NetQuantity  `json:"net_quantity"`
            Description string  `json:"description"`
            RatingCount float64  `json:"rating_count"`
            Price ProductListingPrice  `json:"price"`
            Brand ProductBrand  `json:"brand"`
         
    }
    
    // ProductFiltersValue ...
    type ProductFiltersValue struct {

        
            Value string  `json:"value"`
            Min float64  `json:"min"`
            SelectedMax float64  `json:"selected_max"`
            CurrencyCode string  `json:"currency_code"`
            DisplayFormat string  `json:"display_format"`
            IsSelected bool  `json:"is_selected"`
            Count float64  `json:"count"`
            QueryFormat string  `json:"query_format"`
            Max float64  `json:"max"`
            SelectedMin float64  `json:"selected_min"`
            CurrencySymbol string  `json:"currency_symbol"`
            Display string  `json:"display"`
         
    }
    
    // ProductFiltersKey ...
    type ProductFiltersKey struct {

        
            Logo string  `json:"logo"`
            Kind string  `json:"kind"`
            Display string  `json:"display"`
            Name string  `json:"name"`
         
    }
    
    // ProductFilters ...
    type ProductFilters struct {

        
            Values []ProductFiltersValue  `json:"values"`
            Key ProductFiltersKey  `json:"key"`
         
    }
    
    // ProductSortOn ...
    type ProductSortOn struct {

        
            Value string  `json:"value"`
            IsSelected bool  `json:"is_selected"`
            Name string  `json:"name"`
         
    }
    
    // ProductListingResponse ...
    type ProductListingResponse struct {

        
            Page Page  `json:"page"`
            Items []ProductListingDetail  `json:"items"`
            Filters []ProductFilters  `json:"filters"`
            SortOn []ProductSortOn  `json:"sort_on"`
         
    }
    
    // ImageUrls ...
    type ImageUrls struct {

        
            Landscape Media  `json:"landscape"`
            Portrait Media  `json:"portrait"`
         
    }
    
    // BrandItem ...
    type BrandItem struct {

        
            Action ProductListingAction  `json:"action"`
            Logo Media  `json:"logo"`
            Discount string  `json:"discount"`
            Name string  `json:"name"`
            Departments []string  `json:"departments"`
            Slug string  `json:"slug"`
            Description string  `json:"description"`
            UID float64  `json:"uid"`
            Banners ImageUrls  `json:"banners"`
         
    }
    
    // BrandListingResponse ...
    type BrandListingResponse struct {

        
            Page Page  `json:"page"`
            Items []BrandItem  `json:"items"`
         
    }
    
    // BrandDetailResponse ...
    type BrandDetailResponse struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Logo Media  `json:"logo"`
            Name string  `json:"name"`
            Description string  `json:"description"`
            UID float64  `json:"uid"`
            Banners ImageUrls  `json:"banners"`
         
    }
    
    // DepartmentIdentifier ...
    type DepartmentIdentifier struct {

        
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
         
    }
    
    // CategoryBanner ...
    type CategoryBanner struct {

        
            Landscape Media  `json:"landscape"`
            Portrait Media  `json:"portrait"`
         
    }
    
    // ThirdLevelChild ...
    type ThirdLevelChild struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Action ProductListingAction  `json:"action"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
            Banners ImageUrls  `json:"banners"`
            Childs []map[string]interface{}  `json:"childs"`
         
    }
    
    // SecondLevelChild ...
    type SecondLevelChild struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Action ProductListingAction  `json:"action"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
            Banners ImageUrls  `json:"banners"`
            Childs []ThirdLevelChild  `json:"childs"`
         
    }
    
    // Child ...
    type Child struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Action ProductListingAction  `json:"action"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
            Banners ImageUrls  `json:"banners"`
            Childs []SecondLevelChild  `json:"childs"`
         
    }
    
    // CategoryItems ...
    type CategoryItems struct {

        
            Action ProductListingAction  `json:"action"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
            Banners CategoryBanner  `json:"banners"`
            Childs []Child  `json:"childs"`
         
    }
    
    // DepartmentCategoryTree ...
    type DepartmentCategoryTree struct {

        
            Department string  `json:"department"`
            Items []CategoryItems  `json:"items"`
         
    }
    
    // CategoryListingResponse ...
    type CategoryListingResponse struct {

        
            Departments []DepartmentIdentifier  `json:"departments"`
            Data []DepartmentCategoryTree  `json:"data"`
         
    }
    
    // CategoryMetaResponse ...
    type CategoryMetaResponse struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Logo Media  `json:"logo"`
            Name string  `json:"name"`
            UID float64  `json:"uid"`
            Banners ImageUrls  `json:"banners"`
         
    }
    
    // HomeListingResponse ...
    type HomeListingResponse struct {

        
            Page Page  `json:"page"`
            Message string  `json:"message"`
            Items []ProductListingDetail  `json:"items"`
         
    }
    
    // Department ...
    type Department struct {

        
            Logo Media  `json:"logo"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            PriorityOrder float64  `json:"priority_order"`
            UID float64  `json:"uid"`
         
    }
    
    // DepartmentResponse ...
    type DepartmentResponse struct {

        
            Items []Department  `json:"items"`
         
    }
    
    // AutocompleteItem ...
    type AutocompleteItem struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Action ProductListingAction  `json:"action"`
            Logo Media  `json:"logo"`
            Type string  `json:"type"`
            Display string  `json:"display"`
         
    }
    
    // AutoCompleteResponse ...
    type AutoCompleteResponse struct {

        
            Items []AutocompleteItem  `json:"items"`
         
    }
    
    // CollectionListingFilterTag ...
    type CollectionListingFilterTag struct {

        
            IsSelected bool  `json:"is_selected"`
            Display string  `json:"display"`
            Name string  `json:"name"`
         
    }
    
    // CollectionListingFilterType ...
    type CollectionListingFilterType struct {

        
            IsSelected bool  `json:"is_selected"`
            Display string  `json:"display"`
            Name string  `json:"name"`
         
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

        
            Logo Media  `json:"logo"`
            Badge map[string]interface{}  `json:"badge"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Cron map[string]interface{}  `json:"cron"`
            Type string  `json:"type"`
            Query []CollectionQuery  `json:"query"`
            Banners ImageUrls  `json:"banners"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Action ProductListingAction  `json:"action"`
            Slug string  `json:"slug"`
            Priority float64  `json:"priority"`
            AllowFacets bool  `json:"allow_facets"`
            AllowSort bool  `json:"allow_sort"`
            AppID string  `json:"app_id"`
            Schedule map[string]interface{}  `json:"_schedule"`
            UID string  `json:"uid"`
            Name string  `json:"name"`
            Tag []string  `json:"tag"`
            SortOn string  `json:"sort_on"`
            IsActive bool  `json:"is_active"`
            Description string  `json:"description"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // GetCollectionListingResponse ...
    type GetCollectionListingResponse struct {

        
            Page Page  `json:"page"`
            Filters CollectionListingFilter  `json:"filters"`
            Items []GetCollectionDetailNest  `json:"items"`
         
    }
    
    // CollectionDetailResponse ...
    type CollectionDetailResponse struct {

        
            Logo Media  `json:"logo"`
            Badge map[string]interface{}  `json:"badge"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Cron map[string]interface{}  `json:"cron"`
            Type string  `json:"type"`
            Query []CollectionQuery  `json:"query"`
            Banners ImageUrls  `json:"banners"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Slug string  `json:"slug"`
            Priority float64  `json:"priority"`
            AllowFacets bool  `json:"allow_facets"`
            AllowSort bool  `json:"allow_sort"`
            AppID string  `json:"app_id"`
            Schedule map[string]interface{}  `json:"_schedule"`
            Name string  `json:"name"`
            Tag []string  `json:"tag"`
            SortOn string  `json:"sort_on"`
            IsActive bool  `json:"is_active"`
            Description string  `json:"description"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // GetFollowListingResponse ...
    type GetFollowListingResponse struct {

        
            Items []ProductListingDetail  `json:"items"`
            Page Page  `json:"page"`
         
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

        
            StoreCode string  `json:"store_code"`
            State string  `json:"state"`
            StoreEmail string  `json:"store_email"`
            Name string  `json:"name"`
            Pincode float64  `json:"pincode"`
            Country string  `json:"country"`
            LatLong LatLong  `json:"lat_long"`
            City string  `json:"city"`
            Address string  `json:"address"`
            UID float64  `json:"uid"`
         
    }
    
    // StoreListingResponse ...
    type StoreListingResponse struct {

        
            Page Page  `json:"page"`
            Items []Store  `json:"items"`
         
    }
    
    // StoreDepartments ...
    type StoreDepartments struct {

        
            Logo string  `json:"logo"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            PriorityOrder float64  `json:"priority_order"`
            UID float64  `json:"uid"`
         
    }
    
    // SellerPhoneNumber ...
    type SellerPhoneNumber struct {

        
            Number string  `json:"number"`
            CountryCode float64  `json:"country_code"`
         
    }
    
    // StoreManagerSerializer ...
    type StoreManagerSerializer struct {

        
            MobileNo SellerPhoneNumber  `json:"mobile_no"`
            Email string  `json:"email"`
            Name string  `json:"name"`
         
    }
    
    // CompanyStore ...
    type CompanyStore struct {

        
            BusinessType string  `json:"business_type"`
            UID float64  `json:"uid"`
            CompanyType string  `json:"company_type"`
            Name string  `json:"name"`
         
    }
    
    // StoreAddressSerializer ...
    type StoreAddressSerializer struct {

        
            State string  `json:"state"`
            Latitude float64  `json:"latitude"`
            Landmark string  `json:"landmark"`
            Longitude float64  `json:"longitude"`
            Pincode float64  `json:"pincode"`
            Country string  `json:"country"`
            Address1 string  `json:"address1"`
            City string  `json:"city"`
            Address2 string  `json:"address2"`
         
    }
    
    // AppStore ...
    type AppStore struct {

        
            Name string  `json:"name"`
            Manager StoreManagerSerializer  `json:"manager"`
            Departments []StoreDepartments  `json:"departments"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            Company CompanyStore  `json:"company"`
            Address StoreAddressSerializer  `json:"address"`
            UID float64  `json:"uid"`
         
    }
    
    // ApplicationStoreListing ...
    type ApplicationStoreListing struct {

        
            Page Page  `json:"page"`
            Filters []StoreDepartments  `json:"filters"`
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
            Weekday string  `json:"weekday"`
            Open bool  `json:"open"`
            Closing Time  `json:"closing"`
         
    }
    
    // StoreDetails ...
    type StoreDetails struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Name string  `json:"name"`
            Manager StoreManagerSerializer  `json:"manager"`
            Departments []StoreDepartments  `json:"departments"`
            Timing []StoreTiming  `json:"timing"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            Company CompanyStore  `json:"company"`
            Address StoreAddressSerializer  `json:"address"`
            UID float64  `json:"uid"`
         
    }
    
    // UserDetail ...
    type UserDetail struct {

        
            UserID string  `json:"user_id"`
            SuperUser bool  `json:"super_user"`
            Contact string  `json:"contact"`
            Username string  `json:"username"`
         
    }
    
    // ProductDetails ...
    type ProductDetails struct {

        
            Images []interface{}  `json:"images"`
            ImageNature interface{}  `json:"image_nature"`
            BrandUID float64  `json:"brand_uid"`
            ShortDescription interface{}  `json:"short_description"`
            Media []map[string]interface{}  `json:"media"`
            IsSet bool  `json:"is_set"`
            CountryOfOrigin interface{}  `json:"country_of_origin"`
            HsnCode float64  `json:"hsn_code"`
            OutOfStock bool  `json:"out_of_stock"`
            HasVariant bool  `json:"has_variant"`
            Slug interface{}  `json:"slug"`
            GroupedAttributes map[string]interface{}  `json:"grouped_attributes"`
            ItemCode interface{}  `json:"item_code"`
            Identifier map[string]interface{}  `json:"identifier"`
            Attributes map[string]interface{}  `json:"attributes"`
            Rating float64  `json:"rating"`
            Highlights []interface{}  `json:"highlights"`
            Name interface{}  `json:"name"`
            TemplateTag interface{}  `json:"template_tag"`
            Description interface{}  `json:"description"`
            RatingCount float64  `json:"rating_count"`
         
    }
    
    // Size ...
    type Size struct {

        
            Value interface{}  `json:"value"`
            IsAvailable bool  `json:"is_available"`
            Quantity float64  `json:"quantity"`
            Display interface{}  `json:"display"`
         
    }
    
    // ProductGroupPrice ...
    type ProductGroupPrice struct {

        
            MaxMarked float64  `json:"max_marked"`
            MinEffective float64  `json:"min_effective"`
            MinMarked float64  `json:"min_marked"`
            Currency interface{}  `json:"currency"`
            MaxEffective float64  `json:"max_effective"`
         
    }
    
    // ProductInGroup ...
    type ProductInGroup struct {

        
            AutoAddToCart bool  `json:"auto_add_to_cart"`
            AutoSelect bool  `json:"auto_select"`
            MaxQuantity float64  `json:"max_quantity"`
            ProductDetails ProductDetails  `json:"product_details"`
            ProductUID float64  `json:"product_uid"`
            MinQuantity float64  `json:"min_quantity"`
            AllowRemove bool  `json:"allow_remove"`
            Sizes []Size  `json:"sizes"`
            Price ProductGroupPrice  `json:"price"`
         
    }
    
    // ProductGroupingModel ...
    type ProductGroupingModel struct {

        
            VerifiedOn string  `json:"verified_on"`
            Logo string  `json:"logo"`
            ModifiedOn string  `json:"modified_on"`
            CreatedOn string  `json:"created_on"`
            VerifiedBy UserDetail  `json:"verified_by"`
            Name interface{}  `json:"name"`
            CompanyID float64  `json:"company_id"`
            ModifiedBy UserDetail  `json:"modified_by"`
            Slug interface{}  `json:"slug"`
            Choice interface{}  `json:"choice"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            IsActive bool  `json:"is_active"`
            Products []ProductInGroup  `json:"products"`
            CreatedBy UserDetail  `json:"created_by"`
            ID interface{}  `json:"_id"`
            PageVisibility []interface{}  `json:"page_visibility"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // ProductBundle ...
    type ProductBundle struct {

        
            Items []ProductGroupingModel  `json:"items"`
         
    }
    
    // SellerV3 ...
    type SellerV3 struct {

        
            UID float64  `json:"uid"`
            Count float64  `json:"count"`
            Name string  `json:"name"`
         
    }
    
    // ProductStockUnitPriceV3 ...
    type ProductStockUnitPriceV3 struct {

        
            CurrencyCode string  `json:"currency_code"`
            Price float64  `json:"price"`
            CurrencySymbol string  `json:"currency_symbol"`
            Unit string  `json:"unit"`
         
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
    
    // ProductStockPriceV3 ...
    type ProductStockPriceV3 struct {

        
            Marked float64  `json:"marked"`
            Effective float64  `json:"effective"`
            Currency string  `json:"currency"`
         
    }
    
    // StoreV3 ...
    type StoreV3 struct {

        
            UID float64  `json:"uid"`
            Count float64  `json:"count"`
            Name string  `json:"name"`
         
    }
    
    // StrategyWiseListingSchemaV3 ...
    type StrategyWiseListingSchemaV3 struct {

        
            Pincode float64  `json:"pincode"`
            Tat float64  `json:"tat"`
            Quantity float64  `json:"quantity"`
            Distance float64  `json:"distance"`
         
    }
    
    // ReturnConfigSchemaV3 ...
    type ReturnConfigSchemaV3 struct {

        
            Time float64  `json:"time"`
            Unit string  `json:"unit"`
            Returnable bool  `json:"returnable"`
         
    }
    
    // ProductSizePriceResponseV3 ...
    type ProductSizePriceResponseV3 struct {

        
            Quantity float64  `json:"quantity"`
            Discount string  `json:"discount"`
            Seller SellerV3  `json:"seller"`
            PricePerUnit ProductStockUnitPriceV3  `json:"price_per_unit"`
            MarketplaceAttributes []MarketPlaceSttributesSchemaV3  `json:"marketplace_attributes"`
            ArticleAssignment ArticleAssignmentV3  `json:"article_assignment"`
            ArticleID string  `json:"article_id"`
            LongLat []float64  `json:"long_lat"`
            Pincode float64  `json:"pincode"`
            GroupedAttributes []SellerGroupAttributes  `json:"grouped_attributes"`
            Set ProductSetV3  `json:"set"`
            PricePerPiece ProductStockPriceV3  `json:"price_per_piece"`
            IsGift bool  `json:"is_gift"`
            SpecialBadge string  `json:"special_badge"`
            IsCod bool  `json:"is_cod"`
            SellerCount float64  `json:"seller_count"`
            Store StoreV3  `json:"store"`
            StrategyWiseListing []StrategyWiseListingSchemaV3  `json:"strategy_wise_listing"`
            ReturnConfig ReturnConfigSchemaV3  `json:"return_config"`
            ItemType string  `json:"item_type"`
            Price ProductStockPriceV3  `json:"price"`
         
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
    

    
    // CartCurrency ...
    type CartCurrency struct {

        
            Symbol string  `json:"symbol"`
            Code string  `json:"code"`
         
    }
    
    // RawBreakup ...
    type RawBreakup struct {

        
            FyndCash float64  `json:"fynd_cash"`
            ConvenienceFee float64  `json:"convenience_fee"`
            Subtotal float64  `json:"subtotal"`
            Discount float64  `json:"discount"`
            CodCharge float64  `json:"cod_charge"`
            MrpTotal float64  `json:"mrp_total"`
            GstCharges float64  `json:"gst_charges"`
            DeliveryCharge float64  `json:"delivery_charge"`
            YouSaved float64  `json:"you_saved"`
            Vog float64  `json:"vog"`
            Total float64  `json:"total"`
            Coupon float64  `json:"coupon"`
         
    }
    
    // DisplayBreakup ...
    type DisplayBreakup struct {

        
            Display string  `json:"display"`
            CurrencyCode string  `json:"currency_code"`
            CurrencySymbol string  `json:"currency_symbol"`
            Value float64  `json:"value"`
            Message []string  `json:"message"`
            Key string  `json:"key"`
         
    }
    
    // LoyaltyPoints ...
    type LoyaltyPoints struct {

        
            IsApplied bool  `json:"is_applied"`
            Description string  `json:"description"`
            Total float64  `json:"total"`
            Applicable float64  `json:"applicable"`
         
    }
    
    // CouponBreakup ...
    type CouponBreakup struct {

        
            IsApplied bool  `json:"is_applied"`
            Type string  `json:"type"`
            Code string  `json:"code"`
            UID string  `json:"uid"`
            Value float64  `json:"value"`
            Message string  `json:"message"`
         
    }
    
    // CartBreakup ...
    type CartBreakup struct {

        
            Raw RawBreakup  `json:"raw"`
            Display []DisplayBreakup  `json:"display"`
            LoyaltyPoints LoyaltyPoints  `json:"loyalty_points"`
            Coupon CouponBreakup  `json:"coupon"`
         
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
    
    // PromoMeta ...
    type PromoMeta struct {

        
            Message string  `json:"message"`
         
    }
    
    // ProductPrice ...
    type ProductPrice struct {

        
            Effective float64  `json:"effective"`
            CurrencyCode string  `json:"currency_code"`
            AddOn float64  `json:"add_on"`
            Selling float64  `json:"selling"`
            CurrencySymbol string  `json:"currency_symbol"`
            Marked float64  `json:"marked"`
         
    }
    
    // ProductPriceInfo ...
    type ProductPriceInfo struct {

        
            Base ProductPrice  `json:"base"`
            Converted ProductPrice  `json:"converted"`
         
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

        
            Query ActionQuery  `json:"query"`
            URL string  `json:"url"`
            Type string  `json:"type"`
         
    }
    
    // CategoryInfo ...
    type CategoryInfo struct {

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // BaseInfo ...
    type BaseInfo struct {

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // CartProduct ...
    type CartProduct struct {

        
            Images []ProductImage  `json:"images"`
            Action ProductAction  `json:"action"`
            Slug string  `json:"slug"`
            Type string  `json:"type"`
            UID float64  `json:"uid"`
            Categories []CategoryInfo  `json:"categories"`
            Name string  `json:"name"`
            Brand BaseInfo  `json:"brand"`
         
    }
    
    // Ownership ...
    type Ownership struct {

        
            PayableBy string  `json:"payable_by"`
            PayableCategory string  `json:"payable_category"`
         
    }
    
    // AppliedPromotion ...
    type AppliedPromotion struct {

        
            OfferText string  `json:"offer_text"`
            Ownership Ownership  `json:"ownership"`
            MrpPromotion bool  `json:"mrp_promotion"`
            ArticleQuantity float64  `json:"article_quantity"`
            PromoID string  `json:"promo_id"`
            Amount float64  `json:"amount"`
            PromotionType string  `json:"promotion_type"`
         
    }
    
    // BasePrice ...
    type BasePrice struct {

        
            CurrencyCode string  `json:"currency_code"`
            Effective float64  `json:"effective"`
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

        
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            Price ArticlePriceInfo  `json:"price"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            Store BaseInfo  `json:"store"`
            Quantity float64  `json:"quantity"`
            Size string  `json:"size"`
            Seller BaseInfo  `json:"seller"`
            Type string  `json:"type"`
            UID string  `json:"uid"`
            ProductGroupTags []string  `json:"product_group_tags"`
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
            IsValid bool  `json:"is_valid"`
            OtherStoreQuantity float64  `json:"other_store_quantity"`
            Sizes []string  `json:"sizes"`
            OutOfStock bool  `json:"out_of_stock"`
            Deliverable bool  `json:"deliverable"`
         
    }
    
    // CartProductIdentifer ...
    type CartProductIdentifer struct {

        
            Identifier string  `json:"identifier"`
         
    }
    
    // CartProductInfo ...
    type CartProductInfo struct {

        
            IsSet bool  `json:"is_set"`
            BulkOffer map[string]interface{}  `json:"bulk_offer"`
            PromoMeta PromoMeta  `json:"promo_meta"`
            Moq map[string]interface{}  `json:"moq"`
            Price ProductPriceInfo  `json:"price"`
            Discount string  `json:"discount"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            Quantity float64  `json:"quantity"`
            Product CartProduct  `json:"product"`
            PromotionsApplied []AppliedPromotion  `json:"promotions_applied"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            Article ProductArticle  `json:"article"`
            Availability ProductAvailability  `json:"availability"`
            CouponMessage string  `json:"coupon_message"`
            PricePerUnit ProductPriceInfo  `json:"price_per_unit"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
            Message string  `json:"message"`
            Key string  `json:"key"`
         
    }
    
    // PaymentSelectionLock ...
    type PaymentSelectionLock struct {

        
            Enabled bool  `json:"enabled"`
            DefaultOptions string  `json:"default_options"`
            PaymentIdentifier string  `json:"payment_identifier"`
         
    }
    
    // CartDetailResponse ...
    type CartDetailResponse struct {

        
            CheckoutMode string  `json:"checkout_mode"`
            LastModified string  `json:"last_modified"`
            Comment string  `json:"comment"`
            IsValid bool  `json:"is_valid"`
            Gstin string  `json:"gstin"`
            Message string  `json:"message"`
            Currency CartCurrency  `json:"currency"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            Items []CartProductInfo  `json:"items"`
            ID string  `json:"id"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            CouponText string  `json:"coupon_text"`
            AppliedPromoDetails []AppliedPromotion  `json:"applied_promo_details"`
            BuyNow bool  `json:"buy_now"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
         
    }
    
    // AddProductCart ...
    type AddProductCart struct {

        
            ItemID float64  `json:"item_id"`
            ArticleAssignment map[string]interface{}  `json:"article_assignment"`
            ItemSize string  `json:"item_size"`
            SellerID float64  `json:"seller_id"`
            Display string  `json:"display"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            Quantity float64  `json:"quantity"`
            StoreID float64  `json:"store_id"`
            ProductGroupTags []string  `json:"product_group_tags"`
            Pos bool  `json:"pos"`
            ArticleID string  `json:"article_id"`
         
    }
    
    // AddCartRequest ...
    type AddCartRequest struct {

        
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
            ItemID float64  `json:"item_id"`
            ItemSize string  `json:"item_size"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            Quantity float64  `json:"quantity"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
            ArticleID string  `json:"article_id"`
         
    }
    
    // UpdateCartRequest ...
    type UpdateCartRequest struct {

        
            Items []UpdateProductCart  `json:"items"`
            Operation string  `json:"operation"`
         
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
    
    // PageCoupon ...
    type PageCoupon struct {

        
            HasNext bool  `json:"has_next"`
            Current float64  `json:"current"`
            TotalItemCount float64  `json:"total_item_count"`
            HasPrevious bool  `json:"has_previous"`
            Total float64  `json:"total"`
         
    }
    
    // Coupon ...
    type Coupon struct {

        
            IsApplied bool  `json:"is_applied"`
            CouponCode string  `json:"coupon_code"`
            SubTitle string  `json:"sub_title"`
            Title string  `json:"title"`
            CouponValue float64  `json:"coupon_value"`
            IsApplicable bool  `json:"is_applicable"`
            ExpiresOn string  `json:"expires_on"`
            MinimumCartValue float64  `json:"minimum_cart_value"`
            MaxDiscountValue float64  `json:"max_discount_value"`
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
    
    // OfferPrice ...
    type OfferPrice struct {

        
            Effective float64  `json:"effective"`
            BulkEffective float64  `json:"bulk_effective"`
            CurrencyCode string  `json:"currency_code"`
            CurrencySymbol string  `json:"currency_symbol"`
            Marked float64  `json:"marked"`
         
    }
    
    // OfferItem ...
    type OfferItem struct {

        
            Price OfferPrice  `json:"price"`
            AutoApplied bool  `json:"auto_applied"`
            Best bool  `json:"best"`
            Quantity float64  `json:"quantity"`
            Margin float64  `json:"margin"`
            Type string  `json:"type"`
            Total float64  `json:"total"`
         
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

        
            Landmark string  `json:"landmark"`
            GeoLocation GeoLocation  `json:"geo_location"`
            State string  `json:"state"`
            GoogleMapPoint map[string]interface{}  `json:"google_map_point"`
            Name string  `json:"name"`
            CountryCode string  `json:"country_code"`
            Phone string  `json:"phone"`
            Address string  `json:"address"`
            UserID string  `json:"user_id"`
            AreaCode string  `json:"area_code"`
            CountryIsoCode string  `json:"country_iso_code"`
            IsDefaultAddress bool  `json:"is_default_address"`
            AddressType string  `json:"address_type"`
            CountryPhoneCode string  `json:"country_phone_code"`
            AreaCodeSlug string  `json:"area_code_slug"`
            CheckoutMode string  `json:"checkout_mode"`
            Meta map[string]interface{}  `json:"meta"`
            Country string  `json:"country"`
            Area string  `json:"area"`
            ID string  `json:"id"`
            Tags []string  `json:"tags"`
            City string  `json:"city"`
            IsActive bool  `json:"is_active"`
            Email string  `json:"email"`
         
    }
    
    // GetAddressesResponse ...
    type GetAddressesResponse struct {

        
            Address []Address  `json:"address"`
            PiiMasking bool  `json:"pii_masking"`
         
    }
    
    // SaveAddressResponse ...
    type SaveAddressResponse struct {

        
            IsDefaultAddress bool  `json:"is_default_address"`
            ID string  `json:"id"`
            Success bool  `json:"success"`
         
    }
    
    // UpdateAddressResponse ...
    type UpdateAddressResponse struct {

        
            IsDefaultAddress bool  `json:"is_default_address"`
            ID string  `json:"id"`
            IsUpdated bool  `json:"is_updated"`
            Success bool  `json:"success"`
         
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

        
            AddressID string  `json:"address_id"`
            MerchantCode string  `json:"merchant_code"`
            PaymentMode string  `json:"payment_mode"`
            PaymentIdentifier string  `json:"payment_identifier"`
            ID string  `json:"id"`
            AggregatorName string  `json:"aggregator_name"`
         
    }
    
    // CouponValidity ...
    type CouponValidity struct {

        
            Title string  `json:"title"`
            Discount float64  `json:"discount"`
            DisplayMessageEn string  `json:"display_message_en"`
            Valid bool  `json:"valid"`
            Code string  `json:"code"`
         
    }
    
    // PaymentCouponValidate ...
    type PaymentCouponValidate struct {

        
            Success bool  `json:"success"`
            CouponValidity CouponValidity  `json:"coupon_validity"`
            Message string  `json:"message"`
         
    }
    
    // ShipmentResponse ...
    type ShipmentResponse struct {

        
            Shipments float64  `json:"shipments"`
            DpOptions map[string]interface{}  `json:"dp_options"`
            ShipmentType string  `json:"shipment_type"`
            DpID string  `json:"dp_id"`
            FulfillmentType string  `json:"fulfillment_type"`
            FulfillmentID float64  `json:"fulfillment_id"`
            Promise ShipmentPromise  `json:"promise"`
            Items []CartProductInfo  `json:"items"`
            OrderType string  `json:"order_type"`
            BoxType string  `json:"box_type"`
         
    }
    
    // CartShipmentsResponse ...
    type CartShipmentsResponse struct {

        
            LastModified string  `json:"last_modified"`
            Shipments []ShipmentResponse  `json:"shipments"`
            Comment string  `json:"comment"`
            IsValid bool  `json:"is_valid"`
            Gstin string  `json:"gstin"`
            Currency CartCurrency  `json:"currency"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            CouponText string  `json:"coupon_text"`
            Error bool  `json:"error"`
            UID string  `json:"uid"`
            BuyNow bool  `json:"buy_now"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            CheckoutMode string  `json:"checkout_mode"`
            CartID float64  `json:"cart_id"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            ID string  `json:"id"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            Message string  `json:"message"`
         
    }
    
    // StaffCheckout ...
    type StaffCheckout struct {

        
            ID string  `json:"_id"`
            LastName string  `json:"last_name"`
            User string  `json:"user"`
            EmployeeCode string  `json:"employee_code"`
            FirstName string  `json:"first_name"`
         
    }
    
    // CartCheckoutCustomMeta ...
    type CartCheckoutCustomMeta struct {

        
            Value string  `json:"value"`
            Key string  `json:"key"`
         
    }
    
    // CartCheckoutDetailRequest ...
    type CartCheckoutDetailRequest struct {

        
            AddressID string  `json:"address_id"`
            Aggregator string  `json:"aggregator"`
            OrderingStore float64  `json:"ordering_store"`
            CallbackURL string  `json:"callback_url"`
            MerchantCode string  `json:"merchant_code"`
            Staff StaffCheckout  `json:"staff"`
            Meta map[string]interface{}  `json:"meta"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
            PaymentMode string  `json:"payment_mode"`
            PaymentParams map[string]interface{}  `json:"payment_params"`
            PaymentIdentifier string  `json:"payment_identifier"`
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
            BillingAddressID string  `json:"billing_address_id"`
            PaymentAutoConfirm bool  `json:"payment_auto_confirm"`
            CustomMeta []CartCheckoutCustomMeta  `json:"custom_meta"`
         
    }
    
    // CheckCart ...
    type CheckCart struct {

        
            Success bool  `json:"success"`
            CodCharges float64  `json:"cod_charges"`
            Currency CartCurrency  `json:"currency"`
            Message string  `json:"message"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            OrderID string  `json:"order_id"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            StoreCode string  `json:"store_code"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            UserType string  `json:"user_type"`
            LastModified string  `json:"last_modified"`
            IsValid bool  `json:"is_valid"`
            Gstin string  `json:"gstin"`
            DeliveryCharges float64  `json:"delivery_charges"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            Items []CartProductInfo  `json:"items"`
            CouponText string  `json:"coupon_text"`
            ErrorMessage string  `json:"error_message"`
            UID string  `json:"uid"`
            DeliveryChargeOrderValue float64  `json:"delivery_charge_order_value"`
            BuyNow bool  `json:"buy_now"`
            CodMessage string  `json:"cod_message"`
            CheckoutMode string  `json:"checkout_mode"`
            StoreEmps []map[string]interface{}  `json:"store_emps"`
            CartID float64  `json:"cart_id"`
            CodAvailable bool  `json:"cod_available"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            ID string  `json:"id"`
            Comment string  `json:"comment"`
         
    }
    
    // CartCheckoutResponse ...
    type CartCheckoutResponse struct {

        
            Success bool  `json:"success"`
            AppInterceptURL string  `json:"app_intercept_url"`
            CallbackURL string  `json:"callback_url"`
            OrderID string  `json:"order_id"`
            Cart CheckCart  `json:"cart"`
            PaymentConfirmURL string  `json:"payment_confirm_url"`
            Data map[string]interface{}  `json:"data"`
            Message string  `json:"message"`
         
    }
    
    // CartMetaRequest ...
    type CartMetaRequest struct {

        
            Gstin string  `json:"gstin"`
            CheckoutMode string  `json:"checkout_mode"`
            PickUpCustomerDetails map[string]interface{}  `json:"pick_up_customer_details"`
            Comment string  `json:"comment"`
         
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

        
            Token string  `json:"token"`
            ShareURL string  `json:"share_url"`
         
    }
    
    // SharedCartDetails ...
    type SharedCartDetails struct {

        
            Meta map[string]interface{}  `json:"meta"`
            Source map[string]interface{}  `json:"source"`
            User map[string]interface{}  `json:"user"`
            Token string  `json:"token"`
            CreatedOn string  `json:"created_on"`
         
    }
    
    // SharedCart ...
    type SharedCart struct {

        
            Currency CartCurrency  `json:"currency"`
            Message string  `json:"message"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            LastModified string  `json:"last_modified"`
            IsValid bool  `json:"is_valid"`
            Gstin string  `json:"gstin"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            Items []CartProductInfo  `json:"items"`
            CouponText string  `json:"coupon_text"`
            UID string  `json:"uid"`
            BuyNow bool  `json:"buy_now"`
            SharedCartDetails SharedCartDetails  `json:"shared_cart_details"`
            CheckoutMode string  `json:"checkout_mode"`
            CartID float64  `json:"cart_id"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            ID string  `json:"id"`
            Comment string  `json:"comment"`
         
    }
    
    // SharedCartResponse ...
    type SharedCartResponse struct {

        
            Error string  `json:"error"`
            Cart SharedCart  `json:"cart"`
         
    }
    
    // PromotionOffer ...
    type PromotionOffer struct {

        
            OfferText string  `json:"offer_text"`
            ValidTill string  `json:"valid_till"`
            Description string  `json:"description"`
            ID string  `json:"id"`
            PromotionGroup string  `json:"promotion_group"`
         
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
    
    // CurrencyInfo ...
    type CurrencyInfo struct {

        
            Symbol string  `json:"symbol"`
            Code string  `json:"code"`
         
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
            MaxQuantity float64  `json:"max_quantity"`
            Price LadderPrice  `json:"price"`
            Margin float64  `json:"margin"`
            Type string  `json:"type"`
         
    }
    
    // LadderPriceOffer ...
    type LadderPriceOffer struct {

        
            OfferText string  `json:"offer_text"`
            ValidTill string  `json:"valid_till"`
            Description string  `json:"description"`
            ID string  `json:"id"`
            OfferPrices []LadderOfferItem  `json:"offer_prices"`
            PromotionGroup string  `json:"promotion_group"`
         
    }
    
    // LadderPriceOffers ...
    type LadderPriceOffers struct {

        
            Currency CurrencyInfo  `json:"currency"`
            AvailableOffers []LadderPriceOffer  `json:"available_offers"`
         
    }
    

    
    // ApplicationResponse ...
    type ApplicationResponse struct {

        
            Application Application  `json:"application"`
         
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
    
    // LocationCountry ...
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
    
    // Locations ...
    type Locations struct {

        
            Items []map[string]interface{}  `json:"items"`
         
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
            Phone InformationPhone  `json:"phone"`
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
            Secret string  `json:"secret"`
            VerifyAPI string  `json:"verify_api"`
            MerchantID string  `json:"merchant_id"`
            Pin string  `json:"pin"`
            API string  `json:"api"`
            Key string  `json:"key"`
            ConfigType string  `json:"config_type"`
            Sdk bool  `json:"sdk"`
         
    }
    
    // AggregatorsConfigDetailResponse ...
    type AggregatorsConfigDetailResponse struct {

        
            Success bool  `json:"success"`
            Juspay AggregatorConfigDetail  `json:"juspay"`
            Ccavenue AggregatorConfigDetail  `json:"ccavenue"`
            Mswipe AggregatorConfigDetail  `json:"mswipe"`
            Stripe AggregatorConfigDetail  `json:"stripe"`
            Rupifi AggregatorConfigDetail  `json:"rupifi"`
            Razorpay AggregatorConfigDetail  `json:"razorpay"`
            Payumoney AggregatorConfigDetail  `json:"payumoney"`
            Simpl AggregatorConfigDetail  `json:"simpl"`
            Env string  `json:"env"`
         
    }
    
    // ErrorCodeAndDescription ...
    type ErrorCodeAndDescription struct {

        
            Description string  `json:"description"`
            Code string  `json:"code"`
         
    }
    
    // HttpErrorCodeAndResponse ...
    type HttpErrorCodeAndResponse struct {

        
            Success bool  `json:"success"`
            Error ErrorCodeAndDescription  `json:"error"`
         
    }
    
    // AttachCardRequest ...
    type AttachCardRequest struct {

        
            CardID string  `json:"card_id"`
            NameOnCard string  `json:"name_on_card"`
            Refresh bool  `json:"refresh"`
            Nickname string  `json:"nickname"`
         
    }
    
    // AttachCardsResponse ...
    type AttachCardsResponse struct {

        
            Data map[string]interface{}  `json:"data"`
            Success bool  `json:"success"`
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

        
            Success bool  `json:"success"`
            Cards CardPaymentGateway  `json:"cards"`
            Message string  `json:"message"`
         
    }
    
    // Card ...
    type Card struct {

        
            CardName string  `json:"card_name"`
            CardType string  `json:"card_type"`
            ExpYear float64  `json:"exp_year"`
            ExpMonth float64  `json:"exp_month"`
            CompliantWithTokenisationGuidelines bool  `json:"compliant_with_tokenisation_guidelines"`
            AggregatorName string  `json:"aggregator_name"`
            CardID string  `json:"card_id"`
            CardIssuer string  `json:"card_issuer"`
            CardBrandImage string  `json:"card_brand_image"`
            CardNumber string  `json:"card_number"`
            CardBrand string  `json:"card_brand"`
            CardReference string  `json:"card_reference"`
            CardIsin string  `json:"card_isin"`
            CardFingerprint string  `json:"card_fingerprint"`
            Expired bool  `json:"expired"`
            Nickname string  `json:"nickname"`
            CardToken string  `json:"card_token"`
         
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
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
            MerchantParams map[string]interface{}  `json:"merchant_params"`
            PhoneNumber string  `json:"phone_number"`
            Aggregator string  `json:"aggregator"`
            Payload string  `json:"payload"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
            TransactionAmountInPaise float64  `json:"transaction_amount_in_paise"`
         
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
            OrderID string  `json:"order_id"`
            Verified bool  `json:"verified"`
            Amount float64  `json:"amount"`
            TransactionToken string  `json:"transaction_token"`
         
    }
    
    // ChargeCustomerResponse ...
    type ChargeCustomerResponse struct {

        
            Success bool  `json:"success"`
            OrderID string  `json:"order_id"`
            Aggregator string  `json:"aggregator"`
            DeliveryAddressID string  `json:"delivery_address_id"`
            CartID string  `json:"cart_id"`
            Status string  `json:"status"`
            Message string  `json:"message"`
         
    }
    
    // PaymentInitializationRequest ...
    type PaymentInitializationRequest struct {

        
            Contact string  `json:"contact"`
            Currency string  `json:"currency"`
            OrderID string  `json:"order_id"`
            RazorpayPaymentID string  `json:"razorpay_payment_id"`
            Email string  `json:"email"`
            Aggregator string  `json:"aggregator"`
            Vpa string  `json:"vpa"`
            MerchantOrderID string  `json:"merchant_order_id"`
            Amount float64  `json:"amount"`
            Method string  `json:"method"`
            CustomerID string  `json:"customer_id"`
            Timeout float64  `json:"timeout"`
         
    }
    
    // PaymentInitializationResponse ...
    type PaymentInitializationResponse struct {

        
            Currency string  `json:"currency"`
            Success bool  `json:"success"`
            VirtualID string  `json:"virtual_id"`
            Aggregator string  `json:"aggregator"`
            AggregatorOrderID string  `json:"aggregator_order_id"`
            Vpa string  `json:"vpa"`
            RazorpayPaymentID string  `json:"razorpay_payment_id"`
            BqrImage string  `json:"bqr_image"`
            MerchantOrderID string  `json:"merchant_order_id"`
            UpiPollURL string  `json:"upi_poll_url"`
            PollingURL string  `json:"polling_url"`
            Amount float64  `json:"amount"`
            Status string  `json:"status"`
            Method string  `json:"method"`
            CustomerID string  `json:"customer_id"`
            Timeout float64  `json:"timeout"`
         
    }
    
    // PaymentStatusUpdateRequest ...
    type PaymentStatusUpdateRequest struct {

        
            Contact string  `json:"contact"`
            Currency string  `json:"currency"`
            OrderID string  `json:"order_id"`
            Aggregator string  `json:"aggregator"`
            Email string  `json:"email"`
            Vpa string  `json:"vpa"`
            MerchantOrderID string  `json:"merchant_order_id"`
            Amount float64  `json:"amount"`
            Status string  `json:"status"`
            Method string  `json:"method"`
            CustomerID string  `json:"customer_id"`
         
    }
    
    // PaymentStatusUpdateResponse ...
    type PaymentStatusUpdateResponse struct {

        
            Success bool  `json:"success"`
            AggregatorName string  `json:"aggregator_name"`
            RedirectURL string  `json:"redirect_url"`
            Status string  `json:"status"`
            Retry bool  `json:"retry"`
         
    }
    
    // AggregatorRoute ...
    type AggregatorRoute struct {

        
            APILink string  `json:"api_link"`
            Data map[string]interface{}  `json:"data"`
            PaymentFlowData string  `json:"payment_flow_data"`
            PaymentFlow string  `json:"payment_flow"`
         
    }
    
    // PaymentFlow ...
    type PaymentFlow struct {

        
            Juspay AggregatorRoute  `json:"juspay"`
            Ccavenue AggregatorRoute  `json:"ccavenue"`
            Mswipe AggregatorRoute  `json:"mswipe"`
            UpiRazorpay AggregatorRoute  `json:"upi_razorpay"`
            Stripe AggregatorRoute  `json:"stripe"`
            Rupifi AggregatorRoute  `json:"rupifi"`
            Jiopay AggregatorRoute  `json:"jiopay"`
            BqrRazorpay AggregatorRoute  `json:"bqr_razorpay"`
            Razorpay AggregatorRoute  `json:"razorpay"`
            Payubiz AggregatorRoute  `json:"payubiz"`
            Epaylater AggregatorRoute  `json:"epaylater"`
            Fynd AggregatorRoute  `json:"fynd"`
            Simpl AggregatorRoute  `json:"simpl"`
         
    }
    
    // PaymentModeLogo ...
    type PaymentModeLogo struct {

        
            Large string  `json:"large"`
            Small string  `json:"small"`
         
    }
    
    // IntentApp ...
    type IntentApp struct {

        
            PackageName string  `json:"package_name"`
            Code string  `json:"code"`
            Logos PaymentModeLogo  `json:"logos"`
            DisplayName string  `json:"display_name"`
         
    }
    
    // IntentAppErrorList ...
    type IntentAppErrorList struct {

        
            Code string  `json:"code"`
            PackageName string  `json:"package_name"`
         
    }
    
    // PaymentModeList ...
    type PaymentModeList struct {

        
            CardID string  `json:"card_id"`
            CardIssuer string  `json:"card_issuer"`
            CodLimit float64  `json:"cod_limit"`
            CardBrand string  `json:"card_brand"`
            IntentAppErrorList []string  `json:"intent_app_error_list"`
            CardIsin string  `json:"card_isin"`
            CardFingerprint string  `json:"card_fingerprint"`
            Nickname string  `json:"nickname"`
            Code string  `json:"code"`
            RemainingLimit float64  `json:"remaining_limit"`
            CardType string  `json:"card_type"`
            CompliantWithTokenisationGuidelines bool  `json:"compliant_with_tokenisation_guidelines"`
            ExpMonth float64  `json:"exp_month"`
            MerchantCode string  `json:"merchant_code"`
            CardBrandImage string  `json:"card_brand_image"`
            IntentApp []IntentApp  `json:"intent_app"`
            IntentAppErrorDictList []IntentAppErrorList  `json:"intent_app_error_dict_list"`
            RetryCount float64  `json:"retry_count"`
            Name string  `json:"name"`
            AggregatorName string  `json:"aggregator_name"`
            DisplayName string  `json:"display_name"`
            LogoURL PaymentModeLogo  `json:"logo_url"`
            DisplayPriority float64  `json:"display_priority"`
            IntentFlow bool  `json:"intent_flow"`
            FyndVpa string  `json:"fynd_vpa"`
            Timeout float64  `json:"timeout"`
            CardName string  `json:"card_name"`
            ExpYear float64  `json:"exp_year"`
            CodLimitPerOrder float64  `json:"cod_limit_per_order"`
            CardNumber string  `json:"card_number"`
            CardReference string  `json:"card_reference"`
            Expired bool  `json:"expired"`
            CardToken string  `json:"card_token"`
         
    }
    
    // RootPaymentMode ...
    type RootPaymentMode struct {

        
            List []PaymentModeList  `json:"list"`
            IsPayByCardPl bool  `json:"is_pay_by_card_pl"`
            Name string  `json:"name"`
            DisplayName string  `json:"display_name"`
            AggregatorName string  `json:"aggregator_name"`
            AddCardEnabled bool  `json:"add_card_enabled"`
            AnonymousEnable bool  `json:"anonymous_enable"`
            DisplayPriority float64  `json:"display_priority"`
            SaveCard bool  `json:"save_card"`
         
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

        
            Display bool  `json:"display"`
            Message string  `json:"message"`
            Status string  `json:"status"`
         
    }
    
    // EpaylaterBannerResponse ...
    type EpaylaterBannerResponse struct {

        
            Data EpaylaterBannerData  `json:"data"`
            Success bool  `json:"success"`
         
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

        
            IsValid bool  `json:"is_valid"`
            Status string  `json:"status"`
            CustomerName string  `json:"customer_name"`
            UpiVpa string  `json:"upi_vpa"`
         
    }
    
    // ValidateVPAResponse ...
    type ValidateVPAResponse struct {

        
            Data ValidateUPI  `json:"data"`
            Success bool  `json:"success"`
         
    }
    
    // TransferItemsDetails ...
    type TransferItemsDetails struct {

        
            Name string  `json:"name"`
            DisplayName string  `json:"display_name"`
            ID float64  `json:"id"`
            LogoLarge string  `json:"logo_large"`
            LogoSmall string  `json:"logo_small"`
         
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

        
            Title string  `json:"title"`
            Address string  `json:"address"`
            Comment string  `json:"comment"`
            AccountNo string  `json:"account_no"`
            BranchName string  `json:"branch_name"`
            IsActive bool  `json:"is_active"`
            DisplayName string  `json:"display_name"`
            Email string  `json:"email"`
            BeneficiaryID string  `json:"beneficiary_id"`
            TransferMode string  `json:"transfer_mode"`
            AccountHolder string  `json:"account_holder"`
            BankName string  `json:"bank_name"`
            ModifiedOn string  `json:"modified_on"`
            IfscCode string  `json:"ifsc_code"`
            CreatedOn string  `json:"created_on"`
            Subtitle string  `json:"subtitle"`
            DelightsUserName string  `json:"delights_user_name"`
            Mobile string  `json:"mobile"`
            ID float64  `json:"id"`
         
    }
    
    // OrderBeneficiaryResponse ...
    type OrderBeneficiaryResponse struct {

        
            ShowBeneficiaryDetails bool  `json:"show_beneficiary_details"`
            Beneficiaries []OrderBeneficiaryDetails  `json:"beneficiaries"`
         
    }
    
    // NotFoundResourceError ...
    type NotFoundResourceError struct {

        
            Description string  `json:"description"`
            Success bool  `json:"success"`
            Code string  `json:"code"`
         
    }
    
    // IfscCodeResponse ...
    type IfscCodeResponse struct {

        
            Success bool  `json:"success"`
            BankName string  `json:"bank_name"`
            BranchName string  `json:"branch_name"`
         
    }
    
    // ErrorCodeDescription ...
    type ErrorCodeDescription struct {

        
            Description string  `json:"description"`
            Success bool  `json:"success"`
            Code string  `json:"code"`
         
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

        
            IfscCode string  `json:"ifsc_code"`
            Vpa string  `json:"vpa"`
            Email string  `json:"email"`
            AccountNo string  `json:"account_no"`
            AccountHolder string  `json:"account_holder"`
            BankName string  `json:"bank_name"`
            BranchName string  `json:"branch_name"`
            Mobile string  `json:"mobile"`
            Address string  `json:"address"`
            Wallet string  `json:"wallet"`
            Comment string  `json:"comment"`
         
    }
    
    // AddBeneficiaryDetailsRequest ...
    type AddBeneficiaryDetailsRequest struct {

        
            Details BeneficiaryModeDetails  `json:"details"`
            Otp string  `json:"otp"`
            OrderID string  `json:"order_id"`
            TransferMode string  `json:"transfer_mode"`
            RequestID string  `json:"request_id"`
            Delights bool  `json:"delights"`
            ShipmentID string  `json:"shipment_id"`
         
    }
    
    // RefundAccountResponse ...
    type RefundAccountResponse struct {

        
            Data map[string]interface{}  `json:"data"`
            Success bool  `json:"success"`
            IsVerifiedFlag bool  `json:"is_verified_flag"`
            Message string  `json:"message"`
         
    }
    
    // BankDetailsForOTP ...
    type BankDetailsForOTP struct {

        
            IfscCode string  `json:"ifsc_code"`
            AccountNo string  `json:"account_no"`
            BankName string  `json:"bank_name"`
            AccountHolder string  `json:"account_holder"`
            BranchName string  `json:"branch_name"`
         
    }
    
    // AddBeneficiaryDetailsOTPRequest ...
    type AddBeneficiaryDetailsOTPRequest struct {

        
            Details BankDetailsForOTP  `json:"details"`
            OrderID string  `json:"order_id"`
         
    }
    
    // WalletOtpRequest ...
    type WalletOtpRequest struct {

        
            Mobile string  `json:"mobile"`
            CountryCode string  `json:"country_code"`
         
    }
    
    // WalletOtpResponse ...
    type WalletOtpResponse struct {

        
            Success bool  `json:"success"`
            IsVerifiedFlag string  `json:"is_verified_flag"`
            RequestID string  `json:"request_id"`
         
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

        
            MerchantName string  `json:"merchant_name"`
            Success bool  `json:"success"`
            PollingTimeout float64  `json:"polling_timeout"`
            ExternalOrderID string  `json:"external_order_id"`
            PaymentLinkCurrentStatus string  `json:"payment_link_current_status"`
            Amount float64  `json:"amount"`
            StatusCode float64  `json:"status_code"`
            Message string  `json:"message"`
            PaymentLinkURL string  `json:"payment_link_url"`
         
    }
    
    // ErrorDescription ...
    type ErrorDescription struct {

        
            MerchantName string  `json:"merchant_name"`
            PaymentTransactionID string  `json:"payment_transaction_id"`
            Cancelled bool  `json:"cancelled"`
            Msg string  `json:"msg"`
            InvalidID bool  `json:"invalid_id"`
            MerchantOrderID string  `json:"merchant_order_id"`
            Amount float64  `json:"amount"`
            Expired bool  `json:"expired"`
         
    }
    
    // ErrorResponse ...
    type ErrorResponse struct {

        
            StatusCode float64  `json:"status_code"`
            Success bool  `json:"success"`
            Error ErrorDescription  `json:"error"`
            Message string  `json:"message"`
         
    }
    
    // CreatePaymentLinkMeta ...
    type CreatePaymentLinkMeta struct {

        
            CheckoutMode string  `json:"checkout_mode"`
            AssignCardID string  `json:"assign_card_id"`
            CartID string  `json:"cart_id"`
            Amount string  `json:"amount"`
            Pincode string  `json:"pincode"`
         
    }
    
    // CreatePaymentLinkRequest ...
    type CreatePaymentLinkRequest struct {

        
            Meta CreatePaymentLinkMeta  `json:"meta"`
            ExternalOrderID string  `json:"external_order_id"`
            Email string  `json:"email"`
            Description string  `json:"description"`
            MobileNumber string  `json:"mobile_number"`
            Amount float64  `json:"amount"`
         
    }
    
    // CreatePaymentLinkResponse ...
    type CreatePaymentLinkResponse struct {

        
            Success bool  `json:"success"`
            PollingTimeout float64  `json:"polling_timeout"`
            PaymentLinkID string  `json:"payment_link_id"`
            StatusCode float64  `json:"status_code"`
            Message string  `json:"message"`
            PaymentLinkURL string  `json:"payment_link_url"`
         
    }
    
    // CancelOrResendPaymentLinkRequest ...
    type CancelOrResendPaymentLinkRequest struct {

        
            PaymentLinkID string  `json:"payment_link_id"`
         
    }
    
    // ResendPaymentLinkResponse ...
    type ResendPaymentLinkResponse struct {

        
            StatusCode float64  `json:"status_code"`
            Success bool  `json:"success"`
            PollingTimeout float64  `json:"polling_timeout"`
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

        
            Success bool  `json:"success"`
            AggregatorName string  `json:"aggregator_name"`
            OrderID string  `json:"order_id"`
            PaymentLinkID string  `json:"payment_link_id"`
            RedirectURL string  `json:"redirect_url"`
            Amount float64  `json:"amount"`
            Status string  `json:"status"`
            StatusCode float64  `json:"status_code"`
            HttpStatus float64  `json:"http_status"`
            Message string  `json:"message"`
         
    }
    
    // PaymentMethodsMeta ...
    type PaymentMethodsMeta struct {

        
            PaymentGateway string  `json:"payment_gateway"`
            PaymentIdentifier string  `json:"payment_identifier"`
            MerchantCode string  `json:"merchant_code"`
         
    }
    
    // CreateOrderUserPaymentMethods ...
    type CreateOrderUserPaymentMethods struct {

        
            Meta PaymentMethodsMeta  `json:"meta"`
            Mode string  `json:"mode"`
            Name string  `json:"name"`
         
    }
    
    // CreateOrderUserRequest ...
    type CreateOrderUserRequest struct {

        
            Meta map[string]interface{}  `json:"meta"`
            Currency string  `json:"currency"`
            PaymentMethods CreateOrderUserPaymentMethods  `json:"payment_methods"`
            SuccessCallbackURL string  `json:"success_callback_url"`
            PaymentLinkID string  `json:"payment_link_id"`
            FailureCallbackURL string  `json:"failure_callback_url"`
         
    }
    
    // CreateOrderUserData ...
    type CreateOrderUserData struct {

        
            Contact string  `json:"contact"`
            Currency string  `json:"currency"`
            OrderID string  `json:"order_id"`
            Aggregator string  `json:"aggregator"`
            Email string  `json:"email"`
            MerchantOrderID string  `json:"merchant_order_id"`
            Amount float64  `json:"amount"`
            Method string  `json:"method"`
            CustomerID string  `json:"customer_id"`
            CallbackURL string  `json:"callback_url"`
         
    }
    
    // CreateOrderUserResponse ...
    type CreateOrderUserResponse struct {

        
            Success bool  `json:"success"`
            PaymentConfirmURL string  `json:"payment_confirm_url"`
            OrderID string  `json:"order_id"`
            Data CreateOrderUserData  `json:"data"`
            Message string  `json:"message"`
            StatusCode float64  `json:"status_code"`
            CallbackURL string  `json:"callback_url"`
         
    }
    
    // BalanceDetails ...
    type BalanceDetails struct {

        
            FormattedValue string  `json:"formatted_value"`
            Currency string  `json:"currency"`
            Value float64  `json:"value"`
         
    }
    
    // CreditSummary ...
    type CreditSummary struct {

        
            Balance BalanceDetails  `json:"balance"`
            MerchantCustomerRefID string  `json:"merchant_customer_ref_id"`
            StatusMessage string  `json:"status_message"`
            Status string  `json:"status"`
         
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

        
            IsRegistered bool  `json:"is_registered"`
            SignupURL string  `json:"signup_url"`
            Status bool  `json:"status"`
         
    }
    
    // CheckCreditResponse ...
    type CheckCreditResponse struct {

        
            Data CreditDetail  `json:"data"`
            Success bool  `json:"success"`
         
    }
    
    // MarketplaceInfo ...
    type MarketplaceInfo struct {

        
            DateOfJoining string  `json:"date_of_joining"`
            MembershipID string  `json:"membership_id"`
            Name string  `json:"name"`
         
    }
    
    // DeviceDetails ...
    type DeviceDetails struct {

        
            DeviceModel string  `json:"device_model"`
            DeviceType string  `json:"device_type"`
            IdentifierType string  `json:"identifier_type"`
            DeviceMake string  `json:"device_make"`
            Os string  `json:"os"`
            OsVersion string  `json:"os_version"`
            IdentificationNumber string  `json:"identification_number"`
         
    }
    
    // KYCAddress ...
    type KYCAddress struct {

        
            Addressline2 string  `json:"addressline2"`
            Addressline1 string  `json:"addressline1"`
            LandMark string  `json:"land_mark"`
            City string  `json:"city"`
            State string  `json:"state"`
            OwnershipType string  `json:"ownership_type"`
            Pincode string  `json:"pincode"`
         
    }
    
    // BusinessDetails ...
    type BusinessDetails struct {

        
            BusinessType string  `json:"business_type"`
            Vintage string  `json:"vintage"`
            Name string  `json:"name"`
            EntityType string  `json:"entity_type"`
            Gstin string  `json:"gstin"`
            BusinessOwnershipType string  `json:"business_ownership_type"`
            Pan string  `json:"pan"`
            Address KYCAddress  `json:"address"`
            Fssai string  `json:"fssai"`
            ShopAndEstablishment map[string]interface{}  `json:"shop_and_establishment"`
            Fda string  `json:"fda"`
         
    }
    
    // UserPersonalInfoInDetails ...
    type UserPersonalInfoInDetails struct {

        
            Email string  `json:"email"`
            FathersName string  `json:"fathers_name"`
            MobileVerified bool  `json:"mobile_verified"`
            Phone string  `json:"phone"`
            VoterID string  `json:"voter_id"`
            MothersName string  `json:"mothers_name"`
            Pan string  `json:"pan"`
            LastName string  `json:"last_name"`
            EmailVerified bool  `json:"email_verified"`
            FirstName string  `json:"first_name"`
            AddressAsPerID KYCAddress  `json:"address_as_per_id"`
            MiddleName string  `json:"middle_name"`
            Gender string  `json:"gender"`
            Passport string  `json:"passport"`
            DrivingLicense string  `json:"driving_license"`
            Dob string  `json:"dob"`
         
    }
    
    // CustomerOnboardingRequest ...
    type CustomerOnboardingRequest struct {

        
            Source string  `json:"source"`
            MarketplaceInfo MarketplaceInfo  `json:"marketplace_info"`
            Aggregator string  `json:"aggregator"`
            Device DeviceDetails  `json:"device"`
            Mcc string  `json:"mcc"`
            BusinessInfo BusinessDetails  `json:"business_info"`
            PersonalInfo UserPersonalInfoInDetails  `json:"personal_info"`
         
    }
    
    // OnboardSummary ...
    type OnboardSummary struct {

        
            RedirectURL string  `json:"redirect_url"`
            Status bool  `json:"status"`
            Session map[string]interface{}  `json:"session"`
         
    }
    
    // CustomerOnboardingResponse ...
    type CustomerOnboardingResponse struct {

        
            Data OnboardSummary  `json:"data"`
            Success bool  `json:"success"`
         
    }
    

    
    // OrderPage ...
    type OrderPage struct {

        
            HasNext bool  `json:"has_next"`
            Type string  `json:"type"`
            Current float64  `json:"current"`
            ItemTotal float64  `json:"item_total"`
            Size float64  `json:"size"`
         
    }
    
    // BreakupValues ...
    type BreakupValues struct {

        
            Display string  `json:"display"`
            Value float64  `json:"value"`
            Name string  `json:"name"`
         
    }
    
    // ShipmentPayment ...
    type ShipmentPayment struct {

        
            Mop string  `json:"mop"`
            Status string  `json:"status"`
            Mode string  `json:"mode"`
            Logo string  `json:"logo"`
            PaymentMode string  `json:"payment_mode"`
            DisplayName string  `json:"display_name"`
         
    }
    
    // ShipmentUserInfo ...
    type ShipmentUserInfo struct {

        
            Mobile string  `json:"mobile"`
            Gender string  `json:"gender"`
            FirstName string  `json:"first_name"`
            LastName string  `json:"last_name"`
         
    }
    
    // ShipmentTotalDetails ...
    type ShipmentTotalDetails struct {

        
            TotalPrice float64  `json:"total_price"`
            Sizes float64  `json:"sizes"`
            Pieces float64  `json:"pieces"`
         
    }
    
    // FulfillingStore ...
    type FulfillingStore struct {

        
            ID float64  `json:"id"`
            CompanyID float64  `json:"company_id"`
            Code string  `json:"code"`
            CompanyName string  `json:"company_name"`
            Name string  `json:"name"`
         
    }
    
    // DeliveryAddress ...
    type DeliveryAddress struct {

        
            Latitude float64  `json:"latitude"`
            Phone string  `json:"phone"`
            Longitude float64  `json:"longitude"`
            ContactPerson string  `json:"contact_person"`
            Area string  `json:"area"`
            State string  `json:"state"`
            Pincode string  `json:"pincode"`
            Address2 string  `json:"address2"`
            Landmark string  `json:"landmark"`
            AddressCategory string  `json:"address_category"`
            Address string  `json:"address"`
            Name string  `json:"name"`
            Email string  `json:"email"`
            City string  `json:"city"`
            Country string  `json:"country"`
            Address1 string  `json:"address1"`
            AddressType string  `json:"address_type"`
            Version string  `json:"version"`
            CreatedAt string  `json:"created_at"`
            UpdatedAt string  `json:"updated_at"`
         
    }
    
    // TimeStampData ...
    type TimeStampData struct {

        
            Min string  `json:"min"`
            Max string  `json:"max"`
         
    }
    
    // Promise ...
    type Promise struct {

        
            Timestamp TimeStampData  `json:"timestamp"`
            ShowPromise bool  `json:"show_promise"`
         
    }
    
    // NestedTrackingDetails ...
    type NestedTrackingDetails struct {

        
            IsPassed bool  `json:"is_passed"`
            Time string  `json:"time"`
            Status string  `json:"status"`
            IsCurrent bool  `json:"is_current"`
         
    }
    
    // TrackingDetails ...
    type TrackingDetails struct {

        
            Time string  `json:"time"`
            Status string  `json:"status"`
            TrackingDetails []NestedTrackingDetails  `json:"tracking_details"`
            IsPassed bool  `json:"is_passed"`
            IsCurrent bool  `json:"is_current"`
         
    }
    
    // FulfillingCompany ...
    type FulfillingCompany struct {

        
            ID float64  `json:"id"`
            Name string  `json:"name"`
         
    }
    
    // ShipmentStatus ...
    type ShipmentStatus struct {

        
            Title string  `json:"title"`
            HexCode string  `json:"hex_code"`
         
    }
    
    // Invoice ...
    type Invoice struct {

        
            LabelURL string  `json:"label_url"`
            UpdatedDate string  `json:"updated_date"`
            InvoiceURL string  `json:"invoice_url"`
         
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

        
            FreeGiftItemDetails map[string]interface{}  `json:"free_gift_item_details"`
            ParentItemIdentifier string  `json:"parent_item_identifier"`
            ArticleID string  `json:"article_id"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // AppliedPromos ...
    type AppliedPromos struct {

        
            PromotionName string  `json:"promotion_name"`
            MrpPromotion bool  `json:"mrp_promotion"`
            Amount float64  `json:"amount"`
            PromotionType string  `json:"promotion_type"`
            PromoID string  `json:"promo_id"`
            ArticleQuantity float64  `json:"article_quantity"`
            AppliedFreeArticles []AppliedFreeArticles  `json:"applied_free_articles"`
         
    }
    
    // ItemBrand ...
    type ItemBrand struct {

        
            Name string  `json:"name"`
            Logo string  `json:"logo"`
         
    }
    
    // Item ...
    type Item struct {

        
            Brand ItemBrand  `json:"brand"`
            ID float64  `json:"id"`
            SellerIdentifier string  `json:"seller_identifier"`
            Name string  `json:"name"`
            Code string  `json:"code"`
            SlugKey string  `json:"slug_key"`
            Image []string  `json:"image"`
            Size string  `json:"size"`
         
    }
    
    // Identifiers ...
    type Identifiers struct {

        
            Ean string  `json:"ean"`
            SkuCode string  `json:"sku_code"`
         
    }
    
    // FinancialBreakup ...
    type FinancialBreakup struct {

        
            PromotionEffectiveDiscount float64  `json:"promotion_effective_discount"`
            TotalUnits float64  `json:"total_units"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            TransferPrice float64  `json:"transfer_price"`
            ItemName string  `json:"item_name"`
            AmountPaidRoundoff float64  `json:"amount_paid_roundoff"`
            Identifiers Identifiers  `json:"identifiers"`
            CouponValue float64  `json:"coupon_value"`
            ValueOfGood float64  `json:"value_of_good"`
            AmountPaid float64  `json:"amount_paid"`
            Discount float64  `json:"discount"`
            HsnCode string  `json:"hsn_code"`
            PriceEffective float64  `json:"price_effective"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            CodCharges float64  `json:"cod_charges"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            PriceMarked float64  `json:"price_marked"`
            GstTag string  `json:"gst_tag"`
            Size string  `json:"size"`
            RefundAmount float64  `json:"refund_amount"`
            Cashback float64  `json:"cashback"`
            DeliveryCharge float64  `json:"delivery_charge"`
            FyndCredits float64  `json:"fynd_credits"`
            CashbackApplied float64  `json:"cashback_applied"`
            AddedToFyndCash bool  `json:"added_to_fynd_cash"`
            GstFee float64  `json:"gst_fee"`
            RefundCredit float64  `json:"refund_credit"`
         
    }
    
    // Prices ...
    type Prices struct {

        
            PromotionEffectiveDiscount float64  `json:"promotion_effective_discount"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            TransferPrice float64  `json:"transfer_price"`
            AmountPaidRoundoff float64  `json:"amount_paid_roundoff"`
            CouponValue float64  `json:"coupon_value"`
            ValueOfGood float64  `json:"value_of_good"`
            AmountPaid float64  `json:"amount_paid"`
            Discount float64  `json:"discount"`
            PriceEffective float64  `json:"price_effective"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            CodCharges float64  `json:"cod_charges"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            PriceMarked float64  `json:"price_marked"`
            RefundAmount float64  `json:"refund_amount"`
            Cashback float64  `json:"cashback"`
            DeliveryCharge float64  `json:"delivery_charge"`
            FyndCredits float64  `json:"fynd_credits"`
            CashbackApplied float64  `json:"cashback_applied"`
            AddedToFyndCash bool  `json:"added_to_fynd_cash"`
            RefundCredit float64  `json:"refund_credit"`
         
    }
    
    // Bags ...
    type Bags struct {

        
            ParentPromoBags map[string]interface{}  `json:"parent_promo_bags"`
            ID float64  `json:"id"`
            CurrentStatus CurrentStatus  `json:"current_status"`
            DeliveryDate string  `json:"delivery_date"`
            SellerIdentifier string  `json:"seller_identifier"`
            AppliedPromos []AppliedPromos  `json:"applied_promos"`
            Quantity float64  `json:"quantity"`
            Item Item  `json:"item"`
            LineNumber float64  `json:"line_number"`
            CanCancel bool  `json:"can_cancel"`
            ReturnableDate string  `json:"returnable_date"`
            FinancialBreakup []FinancialBreakup  `json:"financial_breakup"`
            Prices Prices  `json:"prices"`
            CanReturn bool  `json:"can_return"`
         
    }
    
    // Shipments ...
    type Shipments struct {

        
            Payment ShipmentPayment  `json:"payment"`
            TrackURL string  `json:"track_url"`
            ShipmentCreatedAt string  `json:"shipment_created_at"`
            UserInfo ShipmentUserInfo  `json:"user_info"`
            ReturnableDate string  `json:"returnable_date"`
            TotalDetails ShipmentTotalDetails  `json:"total_details"`
            SizeInfo map[string]interface{}  `json:"size_info"`
            FulfillingStore FulfillingStore  `json:"fulfilling_store"`
            RefundDetails map[string]interface{}  `json:"refund_details"`
            DeliveryAddress DeliveryAddress  `json:"delivery_address"`
            CanReturn bool  `json:"can_return"`
            Promise Promise  `json:"promise"`
            TrackingDetails []TrackingDetails  `json:"tracking_details"`
            BeneficiaryDetails bool  `json:"beneficiary_details"`
            NeedHelpURL string  `json:"need_help_url"`
            CanBreak map[string]interface{}  `json:"can_break"`
            FulfillingCompany FulfillingCompany  `json:"fulfilling_company"`
            TotalBags float64  `json:"total_bags"`
            DeliveryDate string  `json:"delivery_date"`
            CustomMeta []map[string]interface{}  `json:"custom_meta"`
            TrakingNo string  `json:"traking_no"`
            OrderID string  `json:"order_id"`
            ShowTrackLink bool  `json:"show_track_link"`
            ShipmentID string  `json:"shipment_id"`
            BreakupValues []BreakupValues  `json:"breakup_values"`
            CanCancel bool  `json:"can_cancel"`
            ShipmentStatus ShipmentStatus  `json:"shipment_status"`
            Comment string  `json:"comment"`
            Invoice Invoice  `json:"invoice"`
            Bags []Bags  `json:"bags"`
            ShowDownloadInvoice bool  `json:"show_download_invoice"`
            OrderType string  `json:"order_type"`
            AwbNo string  `json:"awb_no"`
            Prices Prices  `json:"prices"`
            DpName string  `json:"dp_name"`
         
    }
    
    // BagsForReorderArticleAssignment ...
    type BagsForReorderArticleAssignment struct {

        
            Strategy string  `json:"strategy"`
            Level string  `json:"level"`
         
    }
    
    // BagsForReorder ...
    type BagsForReorder struct {

        
            SellerID float64  `json:"seller_id"`
            Quantity float64  `json:"quantity"`
            ItemSize string  `json:"item_size"`
            StoreID float64  `json:"store_id"`
            ItemID float64  `json:"item_id"`
            ArticleAssignment BagsForReorderArticleAssignment  `json:"article_assignment"`
         
    }
    
    // UserInfo ...
    type UserInfo struct {

        
            Mobile string  `json:"mobile"`
            Email string  `json:"email"`
            Gender string  `json:"gender"`
            Name string  `json:"name"`
         
    }
    
    // OrderSchema ...
    type OrderSchema struct {

        
            OrderCreatedTime string  `json:"order_created_time"`
            OrderID string  `json:"order_id"`
            TotalShipmentsInOrder float64  `json:"total_shipments_in_order"`
            BreakupValues []BreakupValues  `json:"breakup_values"`
            Shipments []Shipments  `json:"shipments"`
            BagsForReorder []BagsForReorder  `json:"bags_for_reorder"`
            UserInfo UserInfo  `json:"user_info"`
         
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

        
            AccountName string  `json:"account_name"`
            Status string  `json:"status"`
            UpdatedTime string  `json:"updated_time"`
            Reason string  `json:"reason"`
            LastLocationRecievedAt string  `json:"last_location_recieved_at"`
            Awb string  `json:"awb"`
            ShipmentType string  `json:"shipment_type"`
            UpdatedAt string  `json:"updated_at"`
         
    }
    
    // ShipmentTrack ...
    type ShipmentTrack struct {

        
            Results []Track  `json:"results"`
         
    }
    
    // CustomerDetailsResponse ...
    type CustomerDetailsResponse struct {

        
            Phone string  `json:"phone"`
            Country string  `json:"country"`
            OrderID string  `json:"order_id"`
            ShipmentID string  `json:"shipment_id"`
            Name string  `json:"name"`
         
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

        
            Meta BagReasonMeta  `json:"meta"`
            ID float64  `json:"id"`
            Reasons []BagReasons  `json:"reasons"`
            DisplayName string  `json:"display_name"`
            QcType []string  `json:"qc_type"`
            QuestionSet []QuestionSet  `json:"question_set"`
         
    }
    
    // ShipmentBagReasons ...
    type ShipmentBagReasons struct {

        
            Success bool  `json:"success"`
            Reasons []BagReasons  `json:"reasons"`
         
    }
    
    // ShipmentReason ...
    type ShipmentReason struct {

        
            ReasonText string  `json:"reason_text"`
            Flow string  `json:"flow"`
            Priority float64  `json:"priority"`
            FeedbackType string  `json:"feedback_type"`
            ReasonID float64  `json:"reason_id"`
            ShowTextArea bool  `json:"show_text_area"`
         
    }
    
    // ShipmentReasons ...
    type ShipmentReasons struct {

        
            Reasons []ShipmentReason  `json:"reasons"`
         
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
    
    // Products ...
    type Products struct {

        
            LineNumber float64  `json:"line_number"`
            Quantity float64  `json:"quantity"`
            Identifier string  `json:"identifier"`
         
    }
    
    // ShipmentsRequest ...
    type ShipmentsRequest struct {

        
            DataUpdates DataUpdates  `json:"data_updates"`
            Reasons ReasonsData  `json:"reasons"`
            Products []Products  `json:"products"`
            Identifier string  `json:"identifier"`
         
    }
    
    // StatuesRequest ...
    type StatuesRequest struct {

        
            Status string  `json:"status"`
            Shipments []ShipmentsRequest  `json:"shipments"`
            ExcludeBagsNextState string  `json:"exclude_bags_next_state"`
         
    }
    
    // UpdateShipmentStatusRequest ...
    type UpdateShipmentStatusRequest struct {

        
            UnlockBeforeTransition bool  `json:"unlock_before_transition"`
            Statuses []StatuesRequest  `json:"statuses"`
            Task bool  `json:"task"`
            ForceTransition bool  `json:"force_transition"`
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

        
            Status float64  `json:"status"`
            Code string  `json:"code"`
            Message string  `json:"message"`
            StackTrace string  `json:"stack_trace"`
            Exception string  `json:"exception"`
         
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
    

    
    // BaseInfo ...
    type BaseInfo struct {

        
            Name string  `json:"name"`
            UID float64  `json:"uid"`
         
    }
    
    // BasePrice ...
    type BasePrice struct {

        
            Marked float64  `json:"marked"`
            Effective float64  `json:"effective"`
            CurrencyCode string  `json:"currency_code"`
            CurrencySymbol string  `json:"currency_symbol"`
         
    }
    
    // ArticlePriceInfo ...
    type ArticlePriceInfo struct {

        
            Converted BasePrice  `json:"converted"`
            Base BasePrice  `json:"base"`
         
    }
    
    // ProductArticle ...
    type ProductArticle struct {

        
            ProductGroupTags []string  `json:"product_group_tags"`
            Store BaseInfo  `json:"store"`
            Type string  `json:"type"`
            Price ArticlePriceInfo  `json:"price"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            Size string  `json:"size"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Seller BaseInfo  `json:"seller"`
            Quantity float64  `json:"quantity"`
            UID string  `json:"uid"`
         
    }
    
    // ProductPrice ...
    type ProductPrice struct {

        
            Effective float64  `json:"effective"`
            CurrencyCode string  `json:"currency_code"`
            Marked float64  `json:"marked"`
            AddOn float64  `json:"add_on"`
            CurrencySymbol string  `json:"currency_symbol"`
            Selling float64  `json:"selling"`
         
    }
    
    // ProductPriceInfo ...
    type ProductPriceInfo struct {

        
            Converted ProductPrice  `json:"converted"`
            Base ProductPrice  `json:"base"`
         
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
    
    // CategoryInfo ...
    type CategoryInfo struct {

        
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
    
    // ProductImage ...
    type ProductImage struct {

        
            SecureURL string  `json:"secure_url"`
            AspectRatio string  `json:"aspect_ratio"`
            URL string  `json:"url"`
         
    }
    
    // CartProduct ...
    type CartProduct struct {

        
            Categories []CategoryInfo  `json:"categories"`
            Type string  `json:"type"`
            Action ProductAction  `json:"action"`
            Name string  `json:"name"`
            Brand BaseInfo  `json:"brand"`
            Slug string  `json:"slug"`
            Images []ProductImage  `json:"images"`
            UID float64  `json:"uid"`
         
    }
    
    // PromoMeta ...
    type PromoMeta struct {

        
            Message string  `json:"message"`
         
    }
    
    // Ownership ...
    type Ownership struct {

        
            PayableCategory string  `json:"payable_category"`
            PayableBy string  `json:"payable_by"`
         
    }
    
    // AppliedPromotion ...
    type AppliedPromotion struct {

        
            ArticleQuantity float64  `json:"article_quantity"`
            MrpPromotion bool  `json:"mrp_promotion"`
            Ownership Ownership  `json:"ownership"`
            PromotionType string  `json:"promotion_type"`
            Amount float64  `json:"amount"`
            OfferText string  `json:"offer_text"`
            PromoID string  `json:"promo_id"`
         
    }
    
    // CartProductIdentifer ...
    type CartProductIdentifer struct {

        
            Identifier string  `json:"identifier"`
         
    }
    
    // ProductAvailabilitySize ...
    type ProductAvailabilitySize struct {

        
            Display string  `json:"display"`
            IsAvailable bool  `json:"is_available"`
            Value string  `json:"value"`
         
    }
    
    // ProductAvailability ...
    type ProductAvailability struct {

        
            IsValid bool  `json:"is_valid"`
            Deliverable bool  `json:"deliverable"`
            OtherStoreQuantity float64  `json:"other_store_quantity"`
            OutOfStock bool  `json:"out_of_stock"`
            AvailableSizes []ProductAvailabilitySize  `json:"available_sizes"`
            Sizes []string  `json:"sizes"`
         
    }
    
    // CartProductInfo ...
    type CartProductInfo struct {

        
            Article ProductArticle  `json:"article"`
            Key string  `json:"key"`
            PricePerUnit ProductPriceInfo  `json:"price_per_unit"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            Product CartProduct  `json:"product"`
            Price ProductPriceInfo  `json:"price"`
            Message string  `json:"message"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            PromoMeta PromoMeta  `json:"promo_meta"`
            PromotionsApplied []AppliedPromotion  `json:"promotions_applied"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
            Availability ProductAvailability  `json:"availability"`
            CouponMessage string  `json:"coupon_message"`
            BulkOffer map[string]interface{}  `json:"bulk_offer"`
            IsSet bool  `json:"is_set"`
            Moq map[string]interface{}  `json:"moq"`
            Discount string  `json:"discount"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // CartCurrency ...
    type CartCurrency struct {

        
            Code string  `json:"code"`
            Symbol string  `json:"symbol"`
         
    }
    
    // PaymentSelectionLock ...
    type PaymentSelectionLock struct {

        
            PaymentIdentifier string  `json:"payment_identifier"`
            DefaultOptions string  `json:"default_options"`
            Enabled bool  `json:"enabled"`
         
    }
    
    // LoyaltyPoints ...
    type LoyaltyPoints struct {

        
            Description string  `json:"description"`
            Applicable float64  `json:"applicable"`
            IsApplied bool  `json:"is_applied"`
            Total float64  `json:"total"`
         
    }
    
    // DisplayBreakup ...
    type DisplayBreakup struct {

        
            CurrencyCode string  `json:"currency_code"`
            Key string  `json:"key"`
            Message []string  `json:"message"`
            CurrencySymbol string  `json:"currency_symbol"`
            Value float64  `json:"value"`
            Display string  `json:"display"`
         
    }
    
    // CouponBreakup ...
    type CouponBreakup struct {

        
            IsApplied bool  `json:"is_applied"`
            Type string  `json:"type"`
            Message string  `json:"message"`
            Code string  `json:"code"`
            Value float64  `json:"value"`
            UID string  `json:"uid"`
         
    }
    
    // RawBreakup ...
    type RawBreakup struct {

        
            FyndCash float64  `json:"fynd_cash"`
            ConvenienceFee float64  `json:"convenience_fee"`
            Total float64  `json:"total"`
            Subtotal float64  `json:"subtotal"`
            Coupon float64  `json:"coupon"`
            YouSaved float64  `json:"you_saved"`
            Vog float64  `json:"vog"`
            CodCharge float64  `json:"cod_charge"`
            Discount float64  `json:"discount"`
            DeliveryCharge float64  `json:"delivery_charge"`
            MrpTotal float64  `json:"mrp_total"`
            GstCharges float64  `json:"gst_charges"`
         
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

        
            CouponText string  `json:"coupon_text"`
            IsValid bool  `json:"is_valid"`
            LastModified string  `json:"last_modified"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            Items []CartProductInfo  `json:"items"`
            ID string  `json:"id"`
            Currency CartCurrency  `json:"currency"`
            Message string  `json:"message"`
            Comment string  `json:"comment"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            AppliedPromoDetails []AppliedPromotion  `json:"applied_promo_details"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            BuyNow bool  `json:"buy_now"`
            CheckoutMode string  `json:"checkout_mode"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            Gstin string  `json:"gstin"`
            RestrictCheckout bool  `json:"restrict_checkout"`
         
    }
    
    // AddProductCart ...
    type AddProductCart struct {

        
            ItemID float64  `json:"item_id"`
            ProductGroupTags []string  `json:"product_group_tags"`
            StoreID float64  `json:"store_id"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Pos bool  `json:"pos"`
            SellerID float64  `json:"seller_id"`
            ItemSize string  `json:"item_size"`
            ArticleAssignment map[string]interface{}  `json:"article_assignment"`
            Display string  `json:"display"`
            ArticleID string  `json:"article_id"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // AddCartRequest ...
    type AddCartRequest struct {

        
            Items []AddProductCart  `json:"items"`
         
    }
    
    // AddCartDetailResponse ...
    type AddCartDetailResponse struct {

        
            Partial bool  `json:"partial"`
            Message string  `json:"message"`
            Cart CartDetailResponse  `json:"cart"`
            Success bool  `json:"success"`
         
    }
    
    // UpdateProductCart ...
    type UpdateProductCart struct {

        
            ItemID float64  `json:"item_id"`
            ItemIndex float64  `json:"item_index"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
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

        
            Message string  `json:"message"`
            Cart CartDetailResponse  `json:"cart"`
            Success bool  `json:"success"`
         
    }
    
    // CartItemCountResponse ...
    type CartItemCountResponse struct {

        
            UserCartItemsCount float64  `json:"user_cart_items_count"`
         
    }
    
    // PageCoupon ...
    type PageCoupon struct {

        
            HasNext bool  `json:"has_next"`
            TotalItemCount float64  `json:"total_item_count"`
            Total float64  `json:"total"`
            HasPrevious bool  `json:"has_previous"`
            Current float64  `json:"current"`
         
    }
    
    // Coupon ...
    type Coupon struct {

        
            SubTitle string  `json:"sub_title"`
            IsApplicable bool  `json:"is_applicable"`
            IsApplied bool  `json:"is_applied"`
            MinimumCartValue float64  `json:"minimum_cart_value"`
            Title string  `json:"title"`
            Message string  `json:"message"`
            MaxDiscountValue float64  `json:"max_discount_value"`
            CouponCode string  `json:"coupon_code"`
            ExpiresOn string  `json:"expires_on"`
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
            Marked float64  `json:"marked"`
            BulkEffective float64  `json:"bulk_effective"`
            CurrencySymbol string  `json:"currency_symbol"`
         
    }
    
    // OfferItem ...
    type OfferItem struct {

        
            Total float64  `json:"total"`
            Type string  `json:"type"`
            Price OfferPrice  `json:"price"`
            Margin float64  `json:"margin"`
            Best bool  `json:"best"`
            AutoApplied bool  `json:"auto_applied"`
            Quantity float64  `json:"quantity"`
         
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

        
            Longitude float64  `json:"longitude"`
            Latitude float64  `json:"latitude"`
         
    }
    
    // Address ...
    type Address struct {

        
            City string  `json:"city"`
            Phone string  `json:"phone"`
            Tags []string  `json:"tags"`
            AreaCodeSlug string  `json:"area_code_slug"`
            Name string  `json:"name"`
            UserID string  `json:"user_id"`
            IsActive bool  `json:"is_active"`
            Meta map[string]interface{}  `json:"meta"`
            CountryCode string  `json:"country_code"`
            AddressType string  `json:"address_type"`
            CheckoutMode string  `json:"checkout_mode"`
            IsDefaultAddress bool  `json:"is_default_address"`
            Area string  `json:"area"`
            Landmark string  `json:"landmark"`
            AreaCode string  `json:"area_code"`
            State string  `json:"state"`
            ID string  `json:"id"`
            CountryPhoneCode string  `json:"country_phone_code"`
            Address string  `json:"address"`
            Country string  `json:"country"`
            GoogleMapPoint map[string]interface{}  `json:"google_map_point"`
            GeoLocation GeoLocation  `json:"geo_location"`
            Email string  `json:"email"`
            CountryIsoCode string  `json:"country_iso_code"`
         
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

        
            IsUpdated bool  `json:"is_updated"`
            ID string  `json:"id"`
            IsDefaultAddress bool  `json:"is_default_address"`
            Success bool  `json:"success"`
         
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
            PaymentIdentifier string  `json:"payment_identifier"`
            PaymentMode string  `json:"payment_mode"`
         
    }
    
    // CouponValidity ...
    type CouponValidity struct {

        
            Title string  `json:"title"`
            DisplayMessageEn string  `json:"display_message_en"`
            Code string  `json:"code"`
            Discount float64  `json:"discount"`
            Valid bool  `json:"valid"`
         
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
            Items []CartProductInfo  `json:"items"`
            DpOptions map[string]interface{}  `json:"dp_options"`
            FulfillmentType string  `json:"fulfillment_type"`
            Shipments float64  `json:"shipments"`
            DpID string  `json:"dp_id"`
            BoxType string  `json:"box_type"`
            OrderType string  `json:"order_type"`
            Promise ShipmentPromise  `json:"promise"`
            ShipmentType string  `json:"shipment_type"`
         
    }
    
    // CartShipmentsResponse ...
    type CartShipmentsResponse struct {

        
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            CartID float64  `json:"cart_id"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            Message string  `json:"message"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            Gstin string  `json:"gstin"`
            CouponText string  `json:"coupon_text"`
            LastModified string  `json:"last_modified"`
            IsValid bool  `json:"is_valid"`
            Currency CartCurrency  `json:"currency"`
            ID string  `json:"id"`
            Comment string  `json:"comment"`
            Error bool  `json:"error"`
            Shipments []ShipmentResponse  `json:"shipments"`
            BuyNow bool  `json:"buy_now"`
            CheckoutMode string  `json:"checkout_mode"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            UID string  `json:"uid"`
         
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
    
    // StaffCheckout ...
    type StaffCheckout struct {

        
            FirstName string  `json:"first_name"`
            LastName string  `json:"last_name"`
            User string  `json:"user"`
            EmployeeCode string  `json:"employee_code"`
            ID string  `json:"_id"`
         
    }
    
    // Files ...
    type Files struct {

        
            Key string  `json:"key"`
            Values []string  `json:"values"`
         
    }
    
    // CartCheckoutCustomMeta ...
    type CartCheckoutCustomMeta struct {

        
            Key string  `json:"key"`
            Value string  `json:"value"`
         
    }
    
    // CartPosCheckoutDetailRequest ...
    type CartPosCheckoutDetailRequest struct {

        
            PaymentAutoConfirm bool  `json:"payment_auto_confirm"`
            Pos bool  `json:"pos"`
            PickAtStoreUID float64  `json:"pick_at_store_uid"`
            Staff StaffCheckout  `json:"staff"`
            PaymentMode string  `json:"payment_mode"`
            Meta map[string]interface{}  `json:"meta"`
            Files []Files  `json:"files"`
            CallbackURL string  `json:"callback_url"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
            OrderType string  `json:"order_type"`
            PaymentParams map[string]interface{}  `json:"payment_params"`
            AddressID string  `json:"address_id"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
            Aggregator string  `json:"aggregator"`
            OrderingStore float64  `json:"ordering_store"`
            MerchantCode string  `json:"merchant_code"`
            BillingAddressID string  `json:"billing_address_id"`
            CustomMeta []CartCheckoutCustomMeta  `json:"custom_meta"`
            PaymentIdentifier string  `json:"payment_identifier"`
         
    }
    
    // CheckCart ...
    type CheckCart struct {

        
            CartID float64  `json:"cart_id"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            UserType string  `json:"user_type"`
            Message string  `json:"message"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            DeliveryCharges float64  `json:"delivery_charges"`
            Success bool  `json:"success"`
            CodMessage string  `json:"cod_message"`
            StoreCode string  `json:"store_code"`
            CouponText string  `json:"coupon_text"`
            LastModified string  `json:"last_modified"`
            Currency CartCurrency  `json:"currency"`
            CodAvailable bool  `json:"cod_available"`
            BuyNow bool  `json:"buy_now"`
            CheckoutMode string  `json:"checkout_mode"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            UID string  `json:"uid"`
            DeliveryChargeOrderValue float64  `json:"delivery_charge_order_value"`
            OrderID string  `json:"order_id"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            Gstin string  `json:"gstin"`
            CodCharges float64  `json:"cod_charges"`
            ErrorMessage string  `json:"error_message"`
            IsValid bool  `json:"is_valid"`
            Items []CartProductInfo  `json:"items"`
            ID string  `json:"id"`
            Comment string  `json:"comment"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            StoreEmps []map[string]interface{}  `json:"store_emps"`
         
    }
    
    // CartCheckoutResponse ...
    type CartCheckoutResponse struct {

        
            OrderID string  `json:"order_id"`
            AppInterceptURL string  `json:"app_intercept_url"`
            CallbackURL string  `json:"callback_url"`
            Message string  `json:"message"`
            PaymentConfirmURL string  `json:"payment_confirm_url"`
            Success bool  `json:"success"`
            Data map[string]interface{}  `json:"data"`
            Cart CheckCart  `json:"cart"`
         
    }
    
    // CartMetaRequest ...
    type CartMetaRequest struct {

        
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
    
    // CartDeliveryModesResponse ...
    type CartDeliveryModesResponse struct {

        
            PickupStores []float64  `json:"pickup_stores"`
            AvailableModes []string  `json:"available_modes"`
         
    }
    
    // PickupStoreDetail ...
    type PickupStoreDetail struct {

        
            StoreCode string  `json:"store_code"`
            City string  `json:"city"`
            ID float64  `json:"id"`
            Phone string  `json:"phone"`
            Area string  `json:"area"`
            Address string  `json:"address"`
            Country string  `json:"country"`
            AreaCodeSlug string  `json:"area_code_slug"`
            AddressType string  `json:"address_type"`
            Landmark string  `json:"landmark"`
            Pincode float64  `json:"pincode"`
            Name string  `json:"name"`
            AreaCode string  `json:"area_code"`
            Email string  `json:"email"`
            UID float64  `json:"uid"`
            State string  `json:"state"`
         
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

        
            Token string  `json:"token"`
            ShareURL string  `json:"share_url"`
         
    }
    
    // SharedCartDetails ...
    type SharedCartDetails struct {

        
            Meta map[string]interface{}  `json:"meta"`
            CreatedOn string  `json:"created_on"`
            User map[string]interface{}  `json:"user"`
            Token string  `json:"token"`
            Source map[string]interface{}  `json:"source"`
         
    }
    
    // SharedCart ...
    type SharedCart struct {

        
            CartID float64  `json:"cart_id"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            Message string  `json:"message"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            CouponText string  `json:"coupon_text"`
            LastModified string  `json:"last_modified"`
            Currency CartCurrency  `json:"currency"`
            BuyNow bool  `json:"buy_now"`
            CheckoutMode string  `json:"checkout_mode"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            UID string  `json:"uid"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            Gstin string  `json:"gstin"`
            SharedCartDetails SharedCartDetails  `json:"shared_cart_details"`
            IsValid bool  `json:"is_valid"`
            Items []CartProductInfo  `json:"items"`
            ID string  `json:"id"`
            Comment string  `json:"comment"`
            RestrictCheckout bool  `json:"restrict_checkout"`
         
    }
    
    // SharedCartResponse ...
    type SharedCartResponse struct {

        
            Error string  `json:"error"`
            Cart SharedCart  `json:"cart"`
         
    }
    

    
    // PincodeLatLongData ...
    type PincodeLatLongData struct {

        
            Coordinates []string  `json:"coordinates"`
            Type string  `json:"type"`
         
    }
    
    // PincodeErrorSchemaResponse ...
    type PincodeErrorSchemaResponse struct {

        
            Message string  `json:"message"`
            Type string  `json:"type"`
            Value string  `json:"value"`
         
    }
    
    // CountryMetaResponse ...
    type CountryMetaResponse struct {

        
            CountryCode string  `json:"country_code"`
            IsdCode string  `json:"isd_code"`
         
    }
    
    // PincodeMetaResponse ...
    type PincodeMetaResponse struct {

        
            InternalZoneID float64  `json:"internal_zone_id"`
            Zone string  `json:"zone"`
         
    }
    
    // PincodeParentsResponse ...
    type PincodeParentsResponse struct {

        
            UID string  `json:"uid"`
            DisplayName string  `json:"display_name"`
            SubType string  `json:"sub_type"`
            Name string  `json:"name"`
         
    }
    
    // PincodeDataResponse ...
    type PincodeDataResponse struct {

        
            LatLong PincodeLatLongData  `json:"lat_long"`
            Error PincodeErrorSchemaResponse  `json:"error"`
            MetaCode CountryMetaResponse  `json:"meta_code"`
            DisplayName string  `json:"display_name"`
            UID string  `json:"uid"`
            SubType string  `json:"sub_type"`
            Meta PincodeMetaResponse  `json:"meta"`
            Name string  `json:"name"`
            Parents []PincodeParentsResponse  `json:"parents"`
         
    }
    
    // PincodeApiResponse ...
    type PincodeApiResponse struct {

        
            Data []PincodeDataResponse  `json:"data"`
            Error PincodeErrorSchemaResponse  `json:"error"`
            Success bool  `json:"success"`
         
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

        
            FulfillmentID float64  `json:"fulfillment_id"`
            FromPincode string  `json:"from_pincode"`
            Articles []TATArticlesRequest  `json:"articles"`
         
    }
    
    // TATViewRequest ...
    type TATViewRequest struct {

        
            Source string  `json:"source"`
            Journey string  `json:"journey"`
            LocationDetails []TATLocationDetailsRequest  `json:"location_details"`
            Action string  `json:"action"`
            ToPincode string  `json:"to_pincode"`
            Identifier string  `json:"identifier"`
         
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

        
            Error TATErrorSchemaResponse  `json:"error"`
            Promise TATPromiseResponse  `json:"promise"`
            IsCodAvailable bool  `json:"is_cod_available"`
            ManufacturingTimeUnit string  `json:"manufacturing_time_unit"`
            Category TATCategoryRequest  `json:"category"`
            ManufacturingTime float64  `json:"manufacturing_time"`
            ManufacturingTimeSeconds float64  `json:"_manufacturing_time_seconds"`
         
    }
    
    // TATLocationDetailsResponse ...
    type TATLocationDetailsResponse struct {

        
            FulfillmentID float64  `json:"fulfillment_id"`
            FromPincode string  `json:"from_pincode"`
            Articles []TATArticlesResponse  `json:"articles"`
         
    }
    
    // TATViewResponse ...
    type TATViewResponse struct {

        
            StormbreakerUUID string  `json:"stormbreaker_uuid"`
            Error TATErrorSchemaResponse  `json:"error"`
            PaymentMode string  `json:"payment_mode"`
            Success bool  `json:"success"`
            ToCity string  `json:"to_city"`
            Source string  `json:"source"`
            Journey string  `json:"journey"`
            IsCodAvailable bool  `json:"is_cod_available"`
            RequestUUID string  `json:"request_uuid"`
            LocationDetails []TATLocationDetailsResponse  `json:"location_details"`
            Action string  `json:"action"`
            ToPincode string  `json:"to_pincode"`
            Identifier string  `json:"identifier"`
         
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

        
            DisplayName string  `json:"display_name"`
            UID string  `json:"uid"`
            ParentID string  `json:"parent_id"`
            Logistics LogisticsResponse  `json:"logistics"`
            SubType string  `json:"sub_type"`
            Meta CountryMetaResponse  `json:"meta"`
            Name string  `json:"name"`
            Type string  `json:"type"`
            IsActive bool  `json:"is_active"`
         
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

        
            Configuration map[string]interface{}  `json:"configuration"`
            Articles []map[string]interface{}  `json:"articles"`
            IgnoredLocations []string  `json:"ignored_locations"`
            ToPincode string  `json:"to_pincode"`
            Identifier string  `json:"identifier"`
         
    }
    
    // ReAssignStoreResponse ...
    type ReAssignStoreResponse struct {

        
            Error map[string]interface{}  `json:"error"`
            Success bool  `json:"success"`
            Articles []map[string]interface{}  `json:"articles"`
            ToPincode string  `json:"to_pincode"`
         
    }
    
