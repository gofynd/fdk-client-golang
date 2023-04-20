package application



    
    // NetQuantity ...
    type NetQuantity struct {

        
            Value float64  `json:"value"`
            Unit interface{}  `json:"unit"`
         
    }
    
    // ApplicationItemSEO ...
    type ApplicationItemSEO struct {

        
            Title interface{}  `json:"title"`
            Description interface{}  `json:"description"`
         
    }
    
    // ProductListingActionPage ...
    type ProductListingActionPage struct {

        
            Type string  `json:"type"`
            Query map[string]interface{}  `json:"query"`
            Params map[string]interface{}  `json:"params"`
         
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

        
            Meta Meta  `json:"meta"`
            URL string  `json:"url"`
            Type string  `json:"type"`
            Alt string  `json:"alt"`
         
    }
    
    // ProductBrand ...
    type ProductBrand struct {

        
            Name string  `json:"name"`
            Action ProductListingAction  `json:"action"`
            UID float64  `json:"uid"`
            Logo Media  `json:"logo"`
            Description string  `json:"description"`
         
    }
    
    // Price ...
    type Price struct {

        
            CurrencySymbol string  `json:"currency_symbol"`
            Min float64  `json:"min"`
            Max float64  `json:"max"`
            CurrencyCode string  `json:"currency_code"`
         
    }
    
    // ProductListingPrice ...
    type ProductListingPrice struct {

        
            Effective Price  `json:"effective"`
            Marked Price  `json:"marked"`
         
    }
    
    // CustomMetaFields ...
    type CustomMetaFields struct {

        
            Value string  `json:"value"`
            Key string  `json:"key"`
         
    }
    
    // ApplicationItemMOQ ...
    type ApplicationItemMOQ struct {

        
            Minimum float64  `json:"minimum"`
            Maximum float64  `json:"maximum"`
            IncrementUnit float64  `json:"increment_unit"`
         
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
    
    // ProductCategoryMap ...
    type ProductCategoryMap struct {

        
            L3 ProductBrand  `json:"l3"`
            L1 ProductBrand  `json:"l1"`
            L2 ProductBrand  `json:"l2"`
         
    }
    
    // ProductDetail ...
    type ProductDetail struct {

        
            NetQuantity NetQuantity  `json:"net_quantity"`
            Seo ApplicationItemSEO  `json:"seo"`
            ProductOnlineDate string  `json:"product_online_date"`
            UID float64  `json:"uid"`
            Tags []string  `json:"tags"`
            Similars []string  `json:"similars"`
            Tryouts []string  `json:"tryouts"`
            Action ProductListingAction  `json:"action"`
            Categories []ProductBrand  `json:"categories"`
            Type string  `json:"type"`
            Medias []Media  `json:"medias"`
            Rating float64  `json:"rating"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ShortDescription string  `json:"short_description"`
            HasVariant bool  `json:"has_variant"`
            Description string  `json:"description"`
            Price ProductListingPrice  `json:"price"`
            IsDependent bool  `json:"is_dependent"`
            Name string  `json:"name"`
            CustomMeta []CustomMetaFields  `json:"_custom_meta"`
            TeaserTag string  `json:"teaser_tag"`
            Brand ProductBrand  `json:"brand"`
            Attributes map[string]interface{}  `json:"attributes"`
            Slug string  `json:"slug"`
            ItemType string  `json:"item_type"`
            Moq ApplicationItemMOQ  `json:"moq"`
            RatingCount float64  `json:"rating_count"`
            ImageNature string  `json:"image_nature"`
            Color string  `json:"color"`
            ItemCode string  `json:"item_code"`
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
            Discount string  `json:"discount"`
            Highlights []string  `json:"highlights"`
            CategoryMap ProductCategoryMap  `json:"category_map"`
         
    }
    
    // ErrorResponse ...
    type ErrorResponse struct {

        
            Error string  `json:"error"`
         
    }
    
    // Weight ...
    type Weight struct {

        
            IsDefault bool  `json:"is_default"`
            Shipping float64  `json:"shipping"`
            Unit string  `json:"unit"`
         
    }
    
    // Dimension ...
    type Dimension struct {

        
            Height float64  `json:"height"`
            Width float64  `json:"width"`
            Unit string  `json:"unit"`
            IsDefault bool  `json:"is_default"`
            Length float64  `json:"length"`
         
    }
    
    // ProductSize ...
    type ProductSize struct {

        
            Value string  `json:"value"`
            SellerIdentifiers []string  `json:"seller_identifiers"`
            Weight Weight  `json:"weight"`
            Dimension Dimension  `json:"dimension"`
            Quantity float64  `json:"quantity"`
            Display string  `json:"display"`
            IsAvailable bool  `json:"is_available"`
         
    }
    
    // ColumnHeader ...
    type ColumnHeader struct {

        
            Convertable bool  `json:"convertable"`
            Value string  `json:"value"`
         
    }
    
    // ColumnHeaders ...
    type ColumnHeaders struct {

        
            Col5 ColumnHeader  `json:"col_5"`
            Col4 ColumnHeader  `json:"col_4"`
            Col1 ColumnHeader  `json:"col_1"`
            Col2 ColumnHeader  `json:"col_2"`
            Col3 ColumnHeader  `json:"col_3"`
            Col6 ColumnHeader  `json:"col_6"`
         
    }
    
    // SizeChartValues ...
    type SizeChartValues struct {

        
            Col5 string  `json:"col_5"`
            Col4 string  `json:"col_4"`
            Col1 string  `json:"col_1"`
            Col2 string  `json:"col_2"`
            Col3 string  `json:"col_3"`
            Col6 string  `json:"col_6"`
         
    }
    
    // SizeChart ...
    type SizeChart struct {

        
            Image string  `json:"image"`
            Unit string  `json:"unit"`
            Headers ColumnHeaders  `json:"headers"`
            Sizes []SizeChartValues  `json:"sizes"`
            SizeTip string  `json:"size_tip"`
            Title string  `json:"title"`
            Description string  `json:"description"`
         
    }
    
    // ProductSizeStores ...
    type ProductSizeStores struct {

        
            Count float64  `json:"count"`
         
    }
    
    // ProductSizes ...
    type ProductSizes struct {

        
            MultiSize bool  `json:"multi_size"`
            Sizes []ProductSize  `json:"sizes"`
            Sellable bool  `json:"sellable"`
            Discount string  `json:"discount"`
            SizeChart SizeChart  `json:"size_chart"`
            Stores ProductSizeStores  `json:"stores"`
            Price ProductListingPrice  `json:"price"`
         
    }
    
    // AttributeDetail ...
    type AttributeDetail struct {

        
            Logo string  `json:"logo"`
            Display string  `json:"display"`
            Key string  `json:"key"`
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
            Items []ProductDetail  `json:"items"`
            Subtitle string  `json:"subtitle"`
            Title string  `json:"title"`
         
    }
    
    // ProductFrequentlyComparedSimilarResponse ...
    type ProductFrequentlyComparedSimilarResponse struct {

        
            Similars ProductCompareResponse  `json:"similars"`
         
    }
    
    // ProductVariantItemResponse ...
    type ProductVariantItemResponse struct {

        
            Value string  `json:"value"`
            Name string  `json:"name"`
            CustomMeta []CustomMetaFields  `json:"_custom_meta"`
            Action ProductListingAction  `json:"action"`
            Color string  `json:"color"`
            Medias []Media  `json:"medias"`
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
            ColorName string  `json:"color_name"`
            IsAvailable bool  `json:"is_available"`
         
    }
    
    // ProductVariantResponse ...
    type ProductVariantResponse struct {

        
            Header string  `json:"header"`
            Items []ProductVariantItemResponse  `json:"items"`
            Key string  `json:"key"`
            DisplayType string  `json:"display_type"`
         
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
    
    // ProductStockPrice ...
    type ProductStockPrice struct {

        
            Currency string  `json:"currency"`
            Marked float64  `json:"marked"`
            Effective float64  `json:"effective"`
         
    }
    
    // Seller ...
    type Seller struct {

        
            Count float64  `json:"count"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // StoreDetail ...
    type StoreDetail struct {

        
            City string  `json:"city"`
            Name string  `json:"name"`
            ID float64  `json:"id"`
            Code string  `json:"code"`
         
    }
    
    // ProductStockStatusItem ...
    type ProductStockStatusItem struct {

        
            Size string  `json:"size"`
            Company CompanyDetail  `json:"company"`
            Price ProductStockPrice  `json:"price"`
            Identifier map[string]interface{}  `json:"identifier"`
            UID string  `json:"uid"`
            Quantity float64  `json:"quantity"`
            Seller Seller  `json:"seller"`
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

        
            Items []ProductStockStatusItem  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // ProductSortOn ...
    type ProductSortOn struct {

        
            IsSelected bool  `json:"is_selected"`
            Value string  `json:"value"`
            Name string  `json:"name"`
         
    }
    
    // ProductVariantListingResponse ...
    type ProductVariantListingResponse struct {

        
            Total float64  `json:"total"`
            DisplayType string  `json:"display_type"`
            Items []ProductVariantItemResponse  `json:"items"`
            Header string  `json:"header"`
            Key string  `json:"key"`
         
    }
    
    // ProductListingDetail ...
    type ProductListingDetail struct {

        
            NetQuantity NetQuantity  `json:"net_quantity"`
            Seo ApplicationItemSEO  `json:"seo"`
            ProductOnlineDate string  `json:"product_online_date"`
            UID float64  `json:"uid"`
            Tags []string  `json:"tags"`
            Similars []string  `json:"similars"`
            Tryouts []string  `json:"tryouts"`
            Action ProductListingAction  `json:"action"`
            Categories []ProductBrand  `json:"categories"`
            Type string  `json:"type"`
            Medias []Media  `json:"medias"`
            Rating float64  `json:"rating"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ShortDescription string  `json:"short_description"`
            HasVariant bool  `json:"has_variant"`
            Description string  `json:"description"`
            Price ProductListingPrice  `json:"price"`
            IsDependent bool  `json:"is_dependent"`
            Name string  `json:"name"`
            CustomMeta []CustomMetaFields  `json:"_custom_meta"`
            Identifiers []string  `json:"identifiers"`
            Variants []ProductVariantListingResponse  `json:"variants"`
            TeaserTag string  `json:"teaser_tag"`
            Sizes []string  `json:"sizes"`
            Brand ProductBrand  `json:"brand"`
            Attributes map[string]interface{}  `json:"attributes"`
            Slug string  `json:"slug"`
            ItemType string  `json:"item_type"`
            Moq ApplicationItemMOQ  `json:"moq"`
            RatingCount float64  `json:"rating_count"`
            ImageNature string  `json:"image_nature"`
            Color string  `json:"color"`
            Sellable bool  `json:"sellable"`
            ItemCode string  `json:"item_code"`
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
            Discount string  `json:"discount"`
            Highlights []string  `json:"highlights"`
            CategoryMap ProductCategoryMap  `json:"category_map"`
         
    }
    
    // ProductFiltersValue ...
    type ProductFiltersValue struct {

        
            SelectedMin float64  `json:"selected_min"`
            Value string  `json:"value"`
            CurrencyCode string  `json:"currency_code"`
            DisplayFormat string  `json:"display_format"`
            Count float64  `json:"count"`
            Min float64  `json:"min"`
            IsSelected bool  `json:"is_selected"`
            CurrencySymbol string  `json:"currency_symbol"`
            Display string  `json:"display"`
            SelectedMax float64  `json:"selected_max"`
            QueryFormat string  `json:"query_format"`
            Max float64  `json:"max"`
         
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
    
    // ProductListingResponse ...
    type ProductListingResponse struct {

        
            SortOn []ProductSortOn  `json:"sort_on"`
            Items []ProductListingDetail  `json:"items"`
            Filters []ProductFilters  `json:"filters"`
            Page Page  `json:"page"`
         
    }
    
    // ImageUrls ...
    type ImageUrls struct {

        
            Portrait Media  `json:"portrait"`
            Landscape Media  `json:"landscape"`
         
    }
    
    // BrandItem ...
    type BrandItem struct {

        
            Departments []string  `json:"departments"`
            Name string  `json:"name"`
            Action ProductListingAction  `json:"action"`
            Banners ImageUrls  `json:"banners"`
            Slug string  `json:"slug"`
            Discount string  `json:"discount"`
            UID float64  `json:"uid"`
            Logo Media  `json:"logo"`
            Description string  `json:"description"`
         
    }
    
    // BrandListingResponse ...
    type BrandListingResponse struct {

        
            Items []BrandItem  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // BrandDetailResponse ...
    type BrandDetailResponse struct {

        
            Name string  `json:"name"`
            Banners ImageUrls  `json:"banners"`
            UID float64  `json:"uid"`
            Logo Media  `json:"logo"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Description string  `json:"description"`
         
    }
    
    // DepartmentIdentifier ...
    type DepartmentIdentifier struct {

        
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
         
    }
    
    // ThirdLevelChild ...
    type ThirdLevelChild struct {

        
            Name string  `json:"name"`
            Childs []map[string]interface{}  `json:"childs"`
            Action ProductListingAction  `json:"action"`
            Banners ImageUrls  `json:"banners"`
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // SecondLevelChild ...
    type SecondLevelChild struct {

        
            Name string  `json:"name"`
            Childs []ThirdLevelChild  `json:"childs"`
            Action ProductListingAction  `json:"action"`
            Banners ImageUrls  `json:"banners"`
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // Child ...
    type Child struct {

        
            Name string  `json:"name"`
            Childs []SecondLevelChild  `json:"childs"`
            Action ProductListingAction  `json:"action"`
            Banners ImageUrls  `json:"banners"`
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // CategoryBanner ...
    type CategoryBanner struct {

        
            Portrait Media  `json:"portrait"`
            Landscape Media  `json:"landscape"`
         
    }
    
    // CategoryItems ...
    type CategoryItems struct {

        
            Name string  `json:"name"`
            Childs []Child  `json:"childs"`
            Action ProductListingAction  `json:"action"`
            Banners CategoryBanner  `json:"banners"`
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
         
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

        
            Name string  `json:"name"`
            Banners ImageUrls  `json:"banners"`
            UID float64  `json:"uid"`
            Logo Media  `json:"logo"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // HomeListingResponse ...
    type HomeListingResponse struct {

        
            Items []ProductListingDetail  `json:"items"`
            Message string  `json:"message"`
            Page Page  `json:"page"`
         
    }
    
    // Department ...
    type Department struct {

        
            Name string  `json:"name"`
            PriorityOrder float64  `json:"priority_order"`
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
            Logo Media  `json:"logo"`
         
    }
    
    // DepartmentResponse ...
    type DepartmentResponse struct {

        
            Items []Department  `json:"items"`
         
    }
    
    // AutocompleteItem ...
    type AutocompleteItem struct {

        
            Action ProductListingAction  `json:"action"`
            Type string  `json:"type"`
            Display string  `json:"display"`
            Logo Media  `json:"logo"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // AutoCompleteResponse ...
    type AutoCompleteResponse struct {

        
            Items []AutocompleteItem  `json:"items"`
         
    }
    
    // CollectionQuery ...
    type CollectionQuery struct {

        
            Value []interface{}  `json:"value"`
            Op string  `json:"op"`
            Attribute string  `json:"attribute"`
         
    }
    
    // GetCollectionDetailNest ...
    type GetCollectionDetailNest struct {

        
            Cron map[string]interface{}  `json:"cron"`
            UID string  `json:"uid"`
            Priority float64  `json:"priority"`
            SortOn string  `json:"sort_on"`
            Badge map[string]interface{}  `json:"badge"`
            Tag []string  `json:"tag"`
            AppID string  `json:"app_id"`
            Action ProductListingAction  `json:"action"`
            Type string  `json:"type"`
            Banners ImageUrls  `json:"banners"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            AllowFacets bool  `json:"allow_facets"`
            Name string  `json:"name"`
            Query []CollectionQuery  `json:"query"`
            Schedule map[string]interface{}  `json:"_schedule"`
            Slug string  `json:"slug"`
            Logo Media  `json:"logo"`
            IsActive bool  `json:"is_active"`
            Meta map[string]interface{}  `json:"meta"`
            AllowSort bool  `json:"allow_sort"`
            Description string  `json:"description"`
         
    }
    
    // CollectionListingFilterTag ...
    type CollectionListingFilterTag struct {

        
            Display string  `json:"display"`
            IsSelected bool  `json:"is_selected"`
            Name string  `json:"name"`
         
    }
    
    // CollectionListingFilterType ...
    type CollectionListingFilterType struct {

        
            Display string  `json:"display"`
            IsSelected bool  `json:"is_selected"`
            Name string  `json:"name"`
         
    }
    
    // CollectionListingFilter ...
    type CollectionListingFilter struct {

        
            Tags []CollectionListingFilterTag  `json:"tags"`
            Type []CollectionListingFilterType  `json:"type"`
         
    }
    
    // GetCollectionListingResponse ...
    type GetCollectionListingResponse struct {

        
            Items []GetCollectionDetailNest  `json:"items"`
            Filters CollectionListingFilter  `json:"filters"`
            Page Page  `json:"page"`
         
    }
    
    // CollectionDetailResponse ...
    type CollectionDetailResponse struct {

        
            Cron map[string]interface{}  `json:"cron"`
            Priority float64  `json:"priority"`
            SortOn string  `json:"sort_on"`
            Badge map[string]interface{}  `json:"badge"`
            Tag []string  `json:"tag"`
            AppID string  `json:"app_id"`
            Type string  `json:"type"`
            Banners ImageUrls  `json:"banners"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            AllowFacets bool  `json:"allow_facets"`
            Name string  `json:"name"`
            Query []CollectionQuery  `json:"query"`
            Schedule map[string]interface{}  `json:"_schedule"`
            Slug string  `json:"slug"`
            Logo Media  `json:"logo"`
            IsActive bool  `json:"is_active"`
            Meta map[string]interface{}  `json:"meta"`
            AllowSort bool  `json:"allow_sort"`
            Description string  `json:"description"`
         
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

        
            Country string  `json:"country"`
            StoreEmail string  `json:"store_email"`
            Name string  `json:"name"`
            Pincode float64  `json:"pincode"`
            City string  `json:"city"`
            UID float64  `json:"uid"`
            Address string  `json:"address"`
            StoreCode string  `json:"store_code"`
            LatLong LatLong  `json:"lat_long"`
            State string  `json:"state"`
         
    }
    
    // StoreListingResponse ...
    type StoreListingResponse struct {

        
            Items []Store  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // StoreDepartments ...
    type StoreDepartments struct {

        
            Name string  `json:"name"`
            PriorityOrder float64  `json:"priority_order"`
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
            Logo string  `json:"logo"`
         
    }
    
    // CompanyStore ...
    type CompanyStore struct {

        
            BusinessType string  `json:"business_type"`
            UID float64  `json:"uid"`
            CompanyType string  `json:"company_type"`
            Name string  `json:"name"`
         
    }
    
    // SellerPhoneNumber ...
    type SellerPhoneNumber struct {

        
            CountryCode float64  `json:"country_code"`
            Number string  `json:"number"`
         
    }
    
    // StoreAddressSerializer ...
    type StoreAddressSerializer struct {

        
            Country string  `json:"country"`
            Address1 string  `json:"address1"`
            Pincode float64  `json:"pincode"`
            Longitude float64  `json:"longitude"`
            Latitude float64  `json:"latitude"`
            City string  `json:"city"`
            Landmark string  `json:"landmark"`
            Address2 string  `json:"address2"`
            State string  `json:"state"`
         
    }
    
    // StoreManagerSerializer ...
    type StoreManagerSerializer struct {

        
            Email string  `json:"email"`
            Name string  `json:"name"`
            MobileNo SellerPhoneNumber  `json:"mobile_no"`
         
    }
    
    // AppStore ...
    type AppStore struct {

        
            Departments []StoreDepartments  `json:"departments"`
            Name string  `json:"name"`
            Company CompanyStore  `json:"company"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            UID float64  `json:"uid"`
            Address StoreAddressSerializer  `json:"address"`
            Manager StoreManagerSerializer  `json:"manager"`
         
    }
    
    // ApplicationStoreListing ...
    type ApplicationStoreListing struct {

        
            Items []AppStore  `json:"items"`
            Filters []StoreDepartments  `json:"filters"`
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
            Weekday string  `json:"weekday"`
            Closing Time  `json:"closing"`
         
    }
    
    // StoreDetails ...
    type StoreDetails struct {

        
            Departments []StoreDepartments  `json:"departments"`
            Name string  `json:"name"`
            Company CompanyStore  `json:"company"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            Timing []StoreTiming  `json:"timing"`
            UID float64  `json:"uid"`
            Address StoreAddressSerializer  `json:"address"`
            Manager StoreManagerSerializer  `json:"manager"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // Size ...
    type Size struct {

        
            Quantity float64  `json:"quantity"`
            Value interface{}  `json:"value"`
            Display interface{}  `json:"display"`
            IsAvailable bool  `json:"is_available"`
         
    }
    
    // ProductGroupPrice ...
    type ProductGroupPrice struct {

        
            Currency interface{}  `json:"currency"`
            MinEffective float64  `json:"min_effective"`
            MinMarked float64  `json:"min_marked"`
            MaxEffective float64  `json:"max_effective"`
            MaxMarked float64  `json:"max_marked"`
         
    }
    
    // ProductDetails ...
    type ProductDetails struct {

        
            Identifier map[string]interface{}  `json:"identifier"`
            HsnCode float64  `json:"hsn_code"`
            BrandUID float64  `json:"brand_uid"`
            OutOfStock bool  `json:"out_of_stock"`
            CountryOfOrigin interface{}  `json:"country_of_origin"`
            Media []map[string]interface{}  `json:"media"`
            Rating float64  `json:"rating"`
            ShortDescription interface{}  `json:"short_description"`
            HasVariant bool  `json:"has_variant"`
            Name interface{}  `json:"name"`
            IsSet bool  `json:"is_set"`
            Attributes map[string]interface{}  `json:"attributes"`
            Slug interface{}  `json:"slug"`
            RatingCount float64  `json:"rating_count"`
            ImageNature interface{}  `json:"image_nature"`
            TemplateTag interface{}  `json:"template_tag"`
            Images []interface{}  `json:"images"`
            ItemCode interface{}  `json:"item_code"`
            GroupedAttributes map[string]interface{}  `json:"grouped_attributes"`
            Highlights []interface{}  `json:"highlights"`
            Description interface{}  `json:"description"`
         
    }
    
    // ProductInGroup ...
    type ProductInGroup struct {

        
            AutoSelect bool  `json:"auto_select"`
            Sizes []Size  `json:"sizes"`
            Price ProductGroupPrice  `json:"price"`
            AutoAddToCart bool  `json:"auto_add_to_cart"`
            ProductUID float64  `json:"product_uid"`
            AllowRemove bool  `json:"allow_remove"`
            MaxQuantity float64  `json:"max_quantity"`
            ProductDetails ProductDetails  `json:"product_details"`
            MinQuantity float64  `json:"min_quantity"`
         
    }
    
    // UserDetail ...
    type UserDetail struct {

        
            UserID string  `json:"user_id"`
            Contact string  `json:"contact"`
            SuperUser bool  `json:"super_user"`
            Username string  `json:"username"`
         
    }
    
    // ProductGroupingModel ...
    type ProductGroupingModel struct {

        
            Meta map[string]interface{}  `json:"meta"`
            Products []ProductInGroup  `json:"products"`
            ModifiedBy UserDetail  `json:"modified_by"`
            Name interface{}  `json:"name"`
            Choice interface{}  `json:"choice"`
            PageVisibility []interface{}  `json:"page_visibility"`
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            VerifiedBy UserDetail  `json:"verified_by"`
            Slug interface{}  `json:"slug"`
            CompanyID float64  `json:"company_id"`
            VerifiedOn string  `json:"verified_on"`
            Logo interface{}  `json:"logo"`
            IsActive bool  `json:"is_active"`
            ID interface{}  `json:"_id"`
            CreatedBy UserDetail  `json:"created_by"`
         
    }
    
    // ProductBundle ...
    type ProductBundle struct {

        
            Items []ProductGroupingModel  `json:"items"`
         
    }
    
    // StrategyWiseListingSchemaV3 ...
    type StrategyWiseListingSchemaV3 struct {

        
            Pincode float64  `json:"pincode"`
            Distance float64  `json:"distance"`
            Quantity float64  `json:"quantity"`
            Tat float64  `json:"tat"`
         
    }
    
    // DetailsSchemaV3 ...
    type DetailsSchemaV3 struct {

        
            Type string  `json:"type"`
            Key string  `json:"key"`
            Value string  `json:"value"`
         
    }
    
    // MarketPlaceSttributesSchemaV3 ...
    type MarketPlaceSttributesSchemaV3 struct {

        
            Details []DetailsSchemaV3  `json:"details"`
            Title string  `json:"title"`
         
    }
    
    // ReturnConfigSchemaV3 ...
    type ReturnConfigSchemaV3 struct {

        
            Time float64  `json:"time"`
            Returnable bool  `json:"returnable"`
            Unit string  `json:"unit"`
         
    }
    
    // ArticleAssignmentV3 ...
    type ArticleAssignmentV3 struct {

        
            Strategy string  `json:"strategy"`
            Level string  `json:"level"`
         
    }
    
    // ProductStockUnitPriceV3 ...
    type ProductStockUnitPriceV3 struct {

        
            CurrencySymbol string  `json:"currency_symbol"`
            Unit string  `json:"unit"`
            CurrencyCode string  `json:"currency_code"`
            Price float64  `json:"price"`
         
    }
    
    // StoreV3 ...
    type StoreV3 struct {

        
            Count float64  `json:"count"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // ProductStockPriceV3 ...
    type ProductStockPriceV3 struct {

        
            Currency string  `json:"currency"`
            Marked float64  `json:"marked"`
            Effective float64  `json:"effective"`
         
    }
    
    // ProductSetDistributionSizeV3 ...
    type ProductSetDistributionSizeV3 struct {

        
            Pieces float64  `json:"pieces"`
            Size string  `json:"size"`
         
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
    
    // SellerGroupAttributes ...
    type SellerGroupAttributes struct {

        
            Details []DetailsSchemaV3  `json:"details"`
            Title string  `json:"title"`
         
    }
    
    // SellerV3 ...
    type SellerV3 struct {

        
            Count float64  `json:"count"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // ProductSizePriceResponseV3 ...
    type ProductSizePriceResponseV3 struct {

        
            StrategyWiseListing []StrategyWiseListingSchemaV3  `json:"strategy_wise_listing"`
            SellerCount float64  `json:"seller_count"`
            MarketplaceAttributes []MarketPlaceSttributesSchemaV3  `json:"marketplace_attributes"`
            Quantity float64  `json:"quantity"`
            ReturnConfig ReturnConfigSchemaV3  `json:"return_config"`
            IsGift bool  `json:"is_gift"`
            ArticleAssignment ArticleAssignmentV3  `json:"article_assignment"`
            ArticleID string  `json:"article_id"`
            PricePerUnit ProductStockUnitPriceV3  `json:"price_per_unit"`
            Store StoreV3  `json:"store"`
            Price ProductStockPriceV3  `json:"price"`
            Set ProductSetV3  `json:"set"`
            Pincode float64  `json:"pincode"`
            ItemType string  `json:"item_type"`
            PricePerPiece ProductStockPriceV3  `json:"price_per_piece"`
            IsCod bool  `json:"is_cod"`
            LongLat []float64  `json:"long_lat"`
            GroupedAttributes []SellerGroupAttributes  `json:"grouped_attributes"`
            SpecialBadge string  `json:"special_badge"`
            Discount string  `json:"discount"`
            Seller SellerV3  `json:"seller"`
         
    }
    
    // ProductSizeSellerFilterSchemaV3 ...
    type ProductSizeSellerFilterSchemaV3 struct {

        
            IsSelected bool  `json:"is_selected"`
            Name string  `json:"name"`
            Value string  `json:"value"`
         
    }
    
    // ProductSizeSellersResponseV3 ...
    type ProductSizeSellersResponseV3 struct {

        
            SortOn []ProductSizeSellerFilterSchemaV3  `json:"sort_on"`
            Items []ProductSizePriceResponseV3  `json:"items"`
            Page Page  `json:"page"`
         
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
    
    // FreeGiftItem ...
    type FreeGiftItem struct {

        
            ItemID float64  `json:"item_id"`
            ItemSlug string  `json:"item_slug"`
            ItemName string  `json:"item_name"`
            ItemPriceDetails map[string]interface{}  `json:"item_price_details"`
            ItemBrandName string  `json:"item_brand_name"`
            ItemImagesURL []string  `json:"item_images_url"`
         
    }
    
    // AppliedFreeArticles ...
    type AppliedFreeArticles struct {

        
            ParentItemIdentifier string  `json:"parent_item_identifier"`
            Quantity float64  `json:"quantity"`
            FreeGiftItemDetails FreeGiftItem  `json:"free_gift_item_details"`
            ArticleID string  `json:"article_id"`
         
    }
    
    // DiscountRulesApp ...
    type DiscountRulesApp struct {

        
            MatchedBuyRules []string  `json:"matched_buy_rules"`
            Offer map[string]interface{}  `json:"offer"`
            RawOffer map[string]interface{}  `json:"raw_offer"`
            ItemCriteria map[string]interface{}  `json:"item_criteria"`
         
    }
    
    // BuyRules ...
    type BuyRules struct {

        
            CartConditions map[string]interface{}  `json:"cart_conditions"`
            ItemCriteria map[string]interface{}  `json:"item_criteria"`
         
    }
    
    // AppliedPromotion ...
    type AppliedPromotion struct {

        
            AppliedFreeArticles []AppliedFreeArticles  `json:"applied_free_articles"`
            PromotionName string  `json:"promotion_name"`
            ArticleQuantity float64  `json:"article_quantity"`
            OfferText string  `json:"offer_text"`
            PromotionGroup string  `json:"promotion_group"`
            DiscountRules []DiscountRulesApp  `json:"discount_rules"`
            MrpPromotion bool  `json:"mrp_promotion"`
            PromotionType string  `json:"promotion_type"`
            PromoID string  `json:"promo_id"`
            Amount float64  `json:"amount"`
            BuyRules []BuyRules  `json:"buy_rules"`
         
    }
    
    // CartProductIdentifer ...
    type CartProductIdentifer struct {

        
            Identifier string  `json:"identifier"`
         
    }
    
    // ProductAvailability ...
    type ProductAvailability struct {

        
            Deliverable bool  `json:"deliverable"`
            IsValid bool  `json:"is_valid"`
            Sizes []string  `json:"sizes"`
            OtherStoreQuantity float64  `json:"other_store_quantity"`
            OutOfStock bool  `json:"out_of_stock"`
         
    }
    
    // ProductPrice ...
    type ProductPrice struct {

        
            CurrencyCode string  `json:"currency_code"`
            CurrencySymbol string  `json:"currency_symbol"`
            Effective float64  `json:"effective"`
            Marked float64  `json:"marked"`
            AddOn float64  `json:"add_on"`
            Selling float64  `json:"selling"`
         
    }
    
    // ProductPriceInfo ...
    type ProductPriceInfo struct {

        
            Converted ProductPrice  `json:"converted"`
            Base ProductPrice  `json:"base"`
         
    }
    
    // BaseInfo ...
    type BaseInfo struct {

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // BasePrice ...
    type BasePrice struct {

        
            Effective float64  `json:"effective"`
            CurrencySymbol string  `json:"currency_symbol"`
            CurrencyCode string  `json:"currency_code"`
            Marked float64  `json:"marked"`
         
    }
    
    // ArticlePriceInfo ...
    type ArticlePriceInfo struct {

        
            Converted BasePrice  `json:"converted"`
            Base BasePrice  `json:"base"`
         
    }
    
    // ProductArticle ...
    type ProductArticle struct {

        
            Store BaseInfo  `json:"store"`
            Identifier map[string]interface{}  `json:"identifier"`
            ProductGroupTags []string  `json:"product_group_tags"`
            Seller BaseInfo  `json:"seller"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            Type string  `json:"type"`
            Quantity float64  `json:"quantity"`
            Price ArticlePriceInfo  `json:"price"`
            SellerIdentifier string  `json:"seller_identifier"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            Size string  `json:"size"`
            UID string  `json:"uid"`
         
    }
    
    // PromoMeta ...
    type PromoMeta struct {

        
            Message string  `json:"message"`
         
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
    
    // ProductImage ...
    type ProductImage struct {

        
            URL string  `json:"url"`
            AspectRatio string  `json:"aspect_ratio"`
            SecureURL string  `json:"secure_url"`
         
    }
    
    // CartProduct ...
    type CartProduct struct {

        
            Categories []CategoryInfo  `json:"categories"`
            Name string  `json:"name"`
            Type string  `json:"type"`
            Action ProductAction  `json:"action"`
            Slug string  `json:"slug"`
            Images []ProductImage  `json:"images"`
            UID float64  `json:"uid"`
            Brand BaseInfo  `json:"brand"`
         
    }
    
    // CartProductInfo ...
    type CartProductInfo struct {

        
            PromotionsApplied []AppliedPromotion  `json:"promotions_applied"`
            CouponMessage string  `json:"coupon_message"`
            BulkOffer map[string]interface{}  `json:"bulk_offer"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
            Availability ProductAvailability  `json:"availability"`
            PricePerUnit ProductPriceInfo  `json:"price_per_unit"`
            Quantity float64  `json:"quantity"`
            Price ProductPriceInfo  `json:"price"`
            Message string  `json:"message"`
            Article ProductArticle  `json:"article"`
            PromoMeta PromoMeta  `json:"promo_meta"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            IsSet bool  `json:"is_set"`
            Discount string  `json:"discount"`
            Product CartProduct  `json:"product"`
            Key string  `json:"key"`
         
    }
    
    // CartCurrency ...
    type CartCurrency struct {

        
            Symbol string  `json:"symbol"`
            Code string  `json:"code"`
         
    }
    
    // PaymentSelectionLock ...
    type PaymentSelectionLock struct {

        
            PaymentIdentifier string  `json:"payment_identifier"`
            DefaultOptions string  `json:"default_options"`
            Enabled bool  `json:"enabled"`
         
    }
    
    // DisplayBreakup ...
    type DisplayBreakup struct {

        
            CurrencyCode string  `json:"currency_code"`
            Value float64  `json:"value"`
            CurrencySymbol string  `json:"currency_symbol"`
            Display string  `json:"display"`
            Message []string  `json:"message"`
            Key string  `json:"key"`
         
    }
    
    // RawBreakup ...
    type RawBreakup struct {

        
            Vog float64  `json:"vog"`
            FyndCash float64  `json:"fynd_cash"`
            ConvenienceFee float64  `json:"convenience_fee"`
            MrpTotal float64  `json:"mrp_total"`
            Subtotal float64  `json:"subtotal"`
            GstCharges float64  `json:"gst_charges"`
            Total float64  `json:"total"`
            YouSaved float64  `json:"you_saved"`
            CodCharge float64  `json:"cod_charge"`
            Discount float64  `json:"discount"`
            DeliveryCharge float64  `json:"delivery_charge"`
            Coupon float64  `json:"coupon"`
         
    }
    
    // LoyaltyPoints ...
    type LoyaltyPoints struct {

        
            Description string  `json:"description"`
            Applicable float64  `json:"applicable"`
            Total float64  `json:"total"`
            IsApplied bool  `json:"is_applied"`
         
    }
    
    // CouponBreakup ...
    type CouponBreakup struct {

        
            Value float64  `json:"value"`
            Description string  `json:"description"`
            MaxDiscountValue float64  `json:"max_discount_value"`
            Title string  `json:"title"`
            IsApplied bool  `json:"is_applied"`
            MinimumCartValue float64  `json:"minimum_cart_value"`
            Type string  `json:"type"`
            Code string  `json:"code"`
            CouponType string  `json:"coupon_type"`
            Message string  `json:"message"`
            CouponValue float64  `json:"coupon_value"`
            UID string  `json:"uid"`
            SubTitle string  `json:"sub_title"`
         
    }
    
    // CartBreakup ...
    type CartBreakup struct {

        
            Display []DisplayBreakup  `json:"display"`
            Raw RawBreakup  `json:"raw"`
            LoyaltyPoints LoyaltyPoints  `json:"loyalty_points"`
            Coupon CouponBreakup  `json:"coupon"`
         
    }
    
    // CartDetailResponse ...
    type CartDetailResponse struct {

        
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            CouponText string  `json:"coupon_text"`
            LastModified string  `json:"last_modified"`
            Comment string  `json:"comment"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            BuyNow bool  `json:"buy_now"`
            Items []CartProductInfo  `json:"items"`
            Currency CartCurrency  `json:"currency"`
            IsValid bool  `json:"is_valid"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            Message string  `json:"message"`
            CheckoutMode string  `json:"checkout_mode"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            Gstin string  `json:"gstin"`
            ID string  `json:"id"`
         
    }
    
    // AddProductCart ...
    type AddProductCart struct {

        
            ItemID float64  `json:"item_id"`
            Pos bool  `json:"pos"`
            ProductGroupTags []string  `json:"product_group_tags"`
            SellerID float64  `json:"seller_id"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            StoreID float64  `json:"store_id"`
            Quantity float64  `json:"quantity"`
            Display string  `json:"display"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            ItemSize string  `json:"item_size"`
            ArticleAssignment map[string]interface{}  `json:"article_assignment"`
            ArticleID string  `json:"article_id"`
         
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
            Identifiers CartProductIdentifer  `json:"identifiers"`
            Quantity float64  `json:"quantity"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            ItemSize string  `json:"item_size"`
            ArticleID string  `json:"article_id"`
         
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
    
    // PageCoupon ...
    type PageCoupon struct {

        
            HasNext bool  `json:"has_next"`
            Current float64  `json:"current"`
            TotalItemCount float64  `json:"total_item_count"`
            Total float64  `json:"total"`
            HasPrevious bool  `json:"has_previous"`
         
    }
    
    // Coupon ...
    type Coupon struct {

        
            IsApplicable bool  `json:"is_applicable"`
            Description string  `json:"description"`
            MaxDiscountValue float64  `json:"max_discount_value"`
            CouponCode string  `json:"coupon_code"`
            MinimumCartValue float64  `json:"minimum_cart_value"`
            Title string  `json:"title"`
            IsApplied bool  `json:"is_applied"`
            CouponType string  `json:"coupon_type"`
            ExpiresOn string  `json:"expires_on"`
            Message string  `json:"message"`
            CouponValue float64  `json:"coupon_value"`
            SubTitle string  `json:"sub_title"`
         
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

        
            CurrencyCode string  `json:"currency_code"`
            CurrencySymbol string  `json:"currency_symbol"`
            BulkEffective float64  `json:"bulk_effective"`
            Effective float64  `json:"effective"`
            Marked float64  `json:"marked"`
         
    }
    
    // OfferItem ...
    type OfferItem struct {

        
            Type string  `json:"type"`
            Best bool  `json:"best"`
            Total float64  `json:"total"`
            Price OfferPrice  `json:"price"`
            Quantity float64  `json:"quantity"`
            Margin float64  `json:"margin"`
            AutoApplied bool  `json:"auto_applied"`
         
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

        
            Tags []string  `json:"tags"`
            GeoLocation GeoLocation  `json:"geo_location"`
            UserID string  `json:"user_id"`
            Address string  `json:"address"`
            Meta map[string]interface{}  `json:"meta"`
            Area string  `json:"area"`
            AddressType string  `json:"address_type"`
            AreaCode string  `json:"area_code"`
            AreaCodeSlug string  `json:"area_code_slug"`
            CheckoutMode string  `json:"checkout_mode"`
            IsDefaultAddress bool  `json:"is_default_address"`
            Email string  `json:"email"`
            City string  `json:"city"`
            Country string  `json:"country"`
            Name string  `json:"name"`
            Landmark string  `json:"landmark"`
            CountryCode string  `json:"country_code"`
            GoogleMapPoint map[string]interface{}  `json:"google_map_point"`
            ID string  `json:"id"`
            State string  `json:"state"`
            Phone string  `json:"phone"`
            IsActive bool  `json:"is_active"`
         
    }
    
    // GetAddressesResponse ...
    type GetAddressesResponse struct {

        
            Address []Address  `json:"address"`
         
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

        
            ID string  `json:"id"`
            CartID string  `json:"cart_id"`
            BillingAddressID string  `json:"billing_address_id"`
         
    }
    
    // UpdateCartPaymentRequest ...
    type UpdateCartPaymentRequest struct {

        
            PaymentMode string  `json:"payment_mode"`
            PaymentIdentifier string  `json:"payment_identifier"`
            MerchantCode string  `json:"merchant_code"`
            AddressID string  `json:"address_id"`
            ID string  `json:"id"`
            AggregatorName string  `json:"aggregator_name"`
         
    }
    
    // CouponValidity ...
    type CouponValidity struct {

        
            Valid bool  `json:"valid"`
            Title string  `json:"title"`
            Code string  `json:"code"`
            DisplayMessageEn string  `json:"display_message_en"`
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
            DpID string  `json:"dp_id"`
            Promise ShipmentPromise  `json:"promise"`
            OrderType string  `json:"order_type"`
            BoxType string  `json:"box_type"`
            FulfillmentID float64  `json:"fulfillment_id"`
            FulfillmentType string  `json:"fulfillment_type"`
            ShipmentType string  `json:"shipment_type"`
            DpOptions map[string]interface{}  `json:"dp_options"`
            Shipments float64  `json:"shipments"`
         
    }
    
    // CartShipmentsResponse ...
    type CartShipmentsResponse struct {

        
            Error bool  `json:"error"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            CouponText string  `json:"coupon_text"`
            Comment string  `json:"comment"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            BuyNow bool  `json:"buy_now"`
            IsValid bool  `json:"is_valid"`
            Message string  `json:"message"`
            Gstin string  `json:"gstin"`
            Shipments []ShipmentResponse  `json:"shipments"`
            ID string  `json:"id"`
            CartID float64  `json:"cart_id"`
            LastModified string  `json:"last_modified"`
            Currency CartCurrency  `json:"currency"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            CheckoutMode string  `json:"checkout_mode"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            UID string  `json:"uid"`
         
    }
    
    // StaffCheckout ...
    type StaffCheckout struct {

        
            ID string  `json:"_id"`
            User string  `json:"user"`
            FirstName string  `json:"first_name"`
            LastName string  `json:"last_name"`
         
    }
    
    // CartCheckoutDetailRequest ...
    type CartCheckoutDetailRequest struct {

        
            PaymentMode string  `json:"payment_mode"`
            PaymentIdentifier string  `json:"payment_identifier"`
            Staff StaffCheckout  `json:"staff"`
            Aggregator string  `json:"aggregator"`
            Meta map[string]interface{}  `json:"meta"`
            PaymentParams map[string]interface{}  `json:"payment_params"`
            MerchantCode string  `json:"merchant_code"`
            PaymentAutoConfirm bool  `json:"payment_auto_confirm"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            AddressID string  `json:"address_id"`
            OrderingStore float64  `json:"ordering_store"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
            BillingAddressID string  `json:"billing_address_id"`
            CallbackURL string  `json:"callback_url"`
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
         
    }
    
    // CheckCart ...
    type CheckCart struct {

        
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            Message string  `json:"message"`
            DeliveryCharges float64  `json:"delivery_charges"`
            Gstin string  `json:"gstin"`
            Currency CartCurrency  `json:"currency"`
            CheckoutMode string  `json:"checkout_mode"`
            OrderID string  `json:"order_id"`
            CodCharges float64  `json:"cod_charges"`
            CouponText string  `json:"coupon_text"`
            Comment string  `json:"comment"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            BuyNow bool  `json:"buy_now"`
            IsValid bool  `json:"is_valid"`
            CodMessage string  `json:"cod_message"`
            StoreEmps []map[string]interface{}  `json:"store_emps"`
            ErrorMessage string  `json:"error_message"`
            ID string  `json:"id"`
            CartID float64  `json:"cart_id"`
            StoreCode string  `json:"store_code"`
            UserType string  `json:"user_type"`
            LastModified string  `json:"last_modified"`
            Items []CartProductInfo  `json:"items"`
            Success bool  `json:"success"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            DeliveryChargeOrderValue float64  `json:"delivery_charge_order_value"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            CodAvailable bool  `json:"cod_available"`
            UID string  `json:"uid"`
         
    }
    
    // CartCheckoutResponse ...
    type CartCheckoutResponse struct {

        
            PaymentConfirmURL string  `json:"payment_confirm_url"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
            Data map[string]interface{}  `json:"data"`
            OrderID string  `json:"order_id"`
            AppInterceptURL string  `json:"app_intercept_url"`
            CallbackURL string  `json:"callback_url"`
            Cart CheckCart  `json:"cart"`
         
    }
    
    // CartMetaRequest ...
    type CartMetaRequest struct {

        
            Comment string  `json:"comment"`
            Gstin string  `json:"gstin"`
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
            CreatedOn string  `json:"created_on"`
            User map[string]interface{}  `json:"user"`
            Token string  `json:"token"`
            Source map[string]interface{}  `json:"source"`
         
    }
    
    // SharedCart ...
    type SharedCart struct {

        
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            Message string  `json:"message"`
            Gstin string  `json:"gstin"`
            SharedCartDetails SharedCartDetails  `json:"shared_cart_details"`
            Currency CartCurrency  `json:"currency"`
            CheckoutMode string  `json:"checkout_mode"`
            CouponText string  `json:"coupon_text"`
            Comment string  `json:"comment"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            BuyNow bool  `json:"buy_now"`
            IsValid bool  `json:"is_valid"`
            ID string  `json:"id"`
            CartID float64  `json:"cart_id"`
            LastModified string  `json:"last_modified"`
            Items []CartProductInfo  `json:"items"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            UID string  `json:"uid"`
         
    }
    
    // SharedCartResponse ...
    type SharedCartResponse struct {

        
            Error string  `json:"error"`
            Cart SharedCart  `json:"cart"`
         
    }
    
    // FreeGiftItems ...
    type FreeGiftItems struct {

        
            ItemID float64  `json:"item_id"`
            ItemSlug string  `json:"item_slug"`
            ItemName string  `json:"item_name"`
            ItemPriceDetails map[string]interface{}  `json:"item_price_details"`
            ItemBrandName string  `json:"item_brand_name"`
            ItemImagesURL []string  `json:"item_images_url"`
         
    }
    
    // PromotionOffer ...
    type PromotionOffer struct {

        
            Description string  `json:"description"`
            OfferText string  `json:"offer_text"`
            PromotionGroup string  `json:"promotion_group"`
            DiscountRules []map[string]interface{}  `json:"discount_rules"`
            ValidTill string  `json:"valid_till"`
            ID string  `json:"id"`
            BuyRules map[string]interface{}  `json:"buy_rules"`
            FreeGiftItems []FreeGiftItems  `json:"free_gift_items"`
         
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

        
            Symbol string  `json:"symbol"`
            Code string  `json:"code"`
         
    }
    
    // LadderPrice ...
    type LadderPrice struct {

        
            CurrencyCode string  `json:"currency_code"`
            OfferPrice float64  `json:"offer_price"`
            CurrencySymbol string  `json:"currency_symbol"`
            Marked float64  `json:"marked"`
            Effective float64  `json:"effective"`
         
    }
    
    // LadderOfferItem ...
    type LadderOfferItem struct {

        
            Type string  `json:"type"`
            MinQuantity float64  `json:"min_quantity"`
            Price LadderPrice  `json:"price"`
            Margin float64  `json:"margin"`
            MaxQuantity float64  `json:"max_quantity"`
         
    }
    
    // LadderPriceOffer ...
    type LadderPriceOffer struct {

        
            Description string  `json:"description"`
            OfferText string  `json:"offer_text"`
            CalculateOn string  `json:"calculate_on"`
            PromotionGroup string  `json:"promotion_group"`
            DiscountRules []map[string]interface{}  `json:"discount_rules"`
            OfferPrices []LadderOfferItem  `json:"offer_prices"`
            ValidTill string  `json:"valid_till"`
            ID string  `json:"id"`
            BuyRules map[string]interface{}  `json:"buy_rules"`
            FreeGiftItems []FreeGiftItems  `json:"free_gift_items"`
         
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

        
            Key string  `json:"key"`
            UserID string  `json:"user_id"`
            Sdk bool  `json:"sdk"`
            ConfigType string  `json:"config_type"`
            VerifyAPI string  `json:"verify_api"`
            API string  `json:"api"`
            MerchantID string  `json:"merchant_id"`
            Secret string  `json:"secret"`
            MerchantKey string  `json:"merchant_key"`
            Pin string  `json:"pin"`
         
    }
    
    // AggregatorsConfigDetailResponse ...
    type AggregatorsConfigDetailResponse struct {

        
            Mswipe AggregatorConfigDetail  `json:"mswipe"`
            Env string  `json:"env"`
            Juspay AggregatorConfigDetail  `json:"juspay"`
            Payumoney AggregatorConfigDetail  `json:"payumoney"`
            Ccavenue AggregatorConfigDetail  `json:"ccavenue"`
            Rupifi AggregatorConfigDetail  `json:"rupifi"`
            Razorpay AggregatorConfigDetail  `json:"razorpay"`
            Stripe AggregatorConfigDetail  `json:"stripe"`
            Success bool  `json:"success"`
            Simpl AggregatorConfigDetail  `json:"simpl"`
         
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

        
            CardID string  `json:"card_id"`
            Refresh bool  `json:"refresh"`
            Nickname string  `json:"nickname"`
            NameOnCard string  `json:"name_on_card"`
         
    }
    
    // AttachCardsResponse ...
    type AttachCardsResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
            Data map[string]interface{}  `json:"data"`
         
    }
    
    // CardPaymentGateway ...
    type CardPaymentGateway struct {

        
            Aggregator string  `json:"aggregator"`
            CustomerID string  `json:"customer_id"`
            API string  `json:"api"`
         
    }
    
    // ActiveCardPaymentGatewayResponse ...
    type ActiveCardPaymentGatewayResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
            Cards CardPaymentGateway  `json:"cards"`
         
    }
    
    // Card ...
    type Card struct {

        
            CardID string  `json:"card_id"`
            CardNumber string  `json:"card_number"`
            CompliantWithTokenisationGuidelines bool  `json:"compliant_with_tokenisation_guidelines"`
            ExpMonth float64  `json:"exp_month"`
            CardToken string  `json:"card_token"`
            CardType string  `json:"card_type"`
            CardIssuer string  `json:"card_issuer"`
            CardBrand string  `json:"card_brand"`
            Nickname string  `json:"nickname"`
            CardReference string  `json:"card_reference"`
            CardName string  `json:"card_name"`
            ExpYear float64  `json:"exp_year"`
            Expired bool  `json:"expired"`
            CardFingerprint string  `json:"card_fingerprint"`
            CardBrandImage string  `json:"card_brand_image"`
            CardIsin string  `json:"card_isin"`
            AggregatorName string  `json:"aggregator_name"`
         
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

        
            Payload string  `json:"payload"`
            MerchantParams map[string]interface{}  `json:"merchant_params"`
            TransactionAmountInPaise float64  `json:"transaction_amount_in_paise"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
            OrderItems []map[string]interface{}  `json:"order_items"`
            Aggregator string  `json:"aggregator"`
            PhoneNumber string  `json:"phone_number"`
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
         
    }
    
    // ValidateCustomerResponse ...
    type ValidateCustomerResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
            Data map[string]interface{}  `json:"data"`
         
    }
    
    // ChargeCustomerRequest ...
    type ChargeCustomerRequest struct {

        
            OrderID string  `json:"order_id"`
            Verified bool  `json:"verified"`
            Amount float64  `json:"amount"`
            Aggregator string  `json:"aggregator"`
            TransactionToken string  `json:"transaction_token"`
         
    }
    
    // ChargeCustomerResponse ...
    type ChargeCustomerResponse struct {

        
            CartID string  `json:"cart_id"`
            Message string  `json:"message"`
            OrderID string  `json:"order_id"`
            Status string  `json:"status"`
            DeliveryAddressID string  `json:"delivery_address_id"`
            Success bool  `json:"success"`
            Aggregator string  `json:"aggregator"`
         
    }
    
    // PaymentInitializationRequest ...
    type PaymentInitializationRequest struct {

        
            Method string  `json:"method"`
            Vpa string  `json:"vpa"`
            RazorpayPaymentID string  `json:"razorpay_payment_id"`
            Email string  `json:"email"`
            OrderID string  `json:"order_id"`
            Timeout float64  `json:"timeout"`
            CustomerID string  `json:"customer_id"`
            Currency string  `json:"currency"`
            Amount float64  `json:"amount"`
            Aggregator string  `json:"aggregator"`
            Contact string  `json:"contact"`
            MerchantOrderID string  `json:"merchant_order_id"`
         
    }
    
    // PaymentInitializationResponse ...
    type PaymentInitializationResponse struct {

        
            BqrImage string  `json:"bqr_image"`
            VirtualID string  `json:"virtual_id"`
            Method string  `json:"method"`
            Vpa string  `json:"vpa"`
            CustomerID string  `json:"customer_id"`
            RazorpayPaymentID string  `json:"razorpay_payment_id"`
            Timeout float64  `json:"timeout"`
            Status string  `json:"status"`
            Currency string  `json:"currency"`
            AggregatorOrderID string  `json:"aggregator_order_id"`
            Amount float64  `json:"amount"`
            UpiPollURL string  `json:"upi_poll_url"`
            Success bool  `json:"success"`
            Aggregator string  `json:"aggregator"`
            PollingURL string  `json:"polling_url"`
            MerchantOrderID string  `json:"merchant_order_id"`
         
    }
    
    // PaymentStatusUpdateRequest ...
    type PaymentStatusUpdateRequest struct {

        
            Method string  `json:"method"`
            Vpa string  `json:"vpa"`
            CustomerID string  `json:"customer_id"`
            Email string  `json:"email"`
            OrderID string  `json:"order_id"`
            Status string  `json:"status"`
            Currency string  `json:"currency"`
            Amount float64  `json:"amount"`
            Aggregator string  `json:"aggregator"`
            Contact string  `json:"contact"`
            MerchantOrderID string  `json:"merchant_order_id"`
         
    }
    
    // PaymentStatusUpdateResponse ...
    type PaymentStatusUpdateResponse struct {

        
            Retry bool  `json:"retry"`
            RedirectURL string  `json:"redirect_url"`
            Status string  `json:"status"`
            Success bool  `json:"success"`
            AggregatorName string  `json:"aggregator_name"`
         
    }
    
    // PaymentModeLogo ...
    type PaymentModeLogo struct {

        
            Small string  `json:"small"`
            Large string  `json:"large"`
         
    }
    
    // IntentApp ...
    type IntentApp struct {

        
            Code string  `json:"code"`
            DisplayName string  `json:"display_name"`
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

        
            CompliantWithTokenisationGuidelines bool  `json:"compliant_with_tokenisation_guidelines"`
            CardNumber string  `json:"card_number"`
            MerchantCode string  `json:"merchant_code"`
            FyndVpa string  `json:"fynd_vpa"`
            Name string  `json:"name"`
            CardReference string  `json:"card_reference"`
            ExpYear float64  `json:"exp_year"`
            CardBrandImage string  `json:"card_brand_image"`
            DisplayName string  `json:"display_name"`
            RetryCount float64  `json:"retry_count"`
            AggregatorName string  `json:"aggregator_name"`
            ExpMonth float64  `json:"exp_month"`
            CardToken string  `json:"card_token"`
            CardType string  `json:"card_type"`
            Expired bool  `json:"expired"`
            CardName string  `json:"card_name"`
            Code string  `json:"code"`
            Timeout float64  `json:"timeout"`
            IntentAppErrorList []string  `json:"intent_app_error_list"`
            DisplayPriority float64  `json:"display_priority"`
            IntentApp []IntentApp  `json:"intent_app"`
            CardIssuer string  `json:"card_issuer"`
            CardBrand string  `json:"card_brand"`
            IntentAppErrorDictList []IntentAppErrorList  `json:"intent_app_error_dict_list"`
            LogoURL PaymentModeLogo  `json:"logo_url"`
            CardFingerprint string  `json:"card_fingerprint"`
            IntentFlow bool  `json:"intent_flow"`
            CardID string  `json:"card_id"`
            Nickname string  `json:"nickname"`
            CardIsin string  `json:"card_isin"`
         
    }
    
    // RootPaymentMode ...
    type RootPaymentMode struct {

        
            SaveCard bool  `json:"save_card"`
            Name string  `json:"name"`
            AnonymousEnable bool  `json:"anonymous_enable"`
            AddCardEnabled bool  `json:"add_card_enabled"`
            List []PaymentModeList  `json:"list"`
            DisplayPriority float64  `json:"display_priority"`
            IsPayByCardPl bool  `json:"is_pay_by_card_pl"`
            DisplayName string  `json:"display_name"`
            AggregatorName string  `json:"aggregator_name"`
         
    }
    
    // AggregatorRoute ...
    type AggregatorRoute struct {

        
            PaymentFlowData map[string]interface{}  `json:"payment_flow_data"`
            Data map[string]interface{}  `json:"data"`
            APILink string  `json:"api_link"`
            PaymentFlow string  `json:"payment_flow"`
         
    }
    
    // PaymentFlow ...
    type PaymentFlow struct {

        
            Mswipe AggregatorRoute  `json:"mswipe"`
            Fynd AggregatorRoute  `json:"fynd"`
            UpiRazorpay AggregatorRoute  `json:"upi_razorpay"`
            Juspay AggregatorRoute  `json:"juspay"`
            Jiopay AggregatorRoute  `json:"jiopay"`
            Epaylater AggregatorRoute  `json:"epaylater"`
            Ccavenue AggregatorRoute  `json:"ccavenue"`
            Rupifi AggregatorRoute  `json:"rupifi"`
            Payubiz AggregatorRoute  `json:"payubiz"`
            Razorpay AggregatorRoute  `json:"razorpay"`
            Stripe AggregatorRoute  `json:"stripe"`
            BqrRazorpay AggregatorRoute  `json:"bqr_razorpay"`
            Simpl AggregatorRoute  `json:"simpl"`
         
    }
    
    // PaymentOptionAndFlow ...
    type PaymentOptionAndFlow struct {

        
            PaymentOption []RootPaymentMode  `json:"payment_option"`
            PaymentFlows PaymentFlow  `json:"payment_flows"`
         
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

        
            IsValid bool  `json:"is_valid"`
            CustomerName string  `json:"customer_name"`
            Status string  `json:"status"`
            UpiVpa string  `json:"upi_vpa"`
         
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
            DisplayName string  `json:"display_name"`
            LogoSmall string  `json:"logo_small"`
            LogoLarge string  `json:"logo_large"`
         
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

        
            Comment string  `json:"comment"`
            TransferMode string  `json:"transfer_mode"`
            Email string  `json:"email"`
            IfscCode string  `json:"ifsc_code"`
            DisplayName string  `json:"display_name"`
            IsActive bool  `json:"is_active"`
            ModifiedOn string  `json:"modified_on"`
            BankName string  `json:"bank_name"`
            CreatedOn string  `json:"created_on"`
            AccountNo string  `json:"account_no"`
            Title string  `json:"title"`
            ID float64  `json:"id"`
            Address string  `json:"address"`
            BranchName string  `json:"branch_name"`
            Mobile string  `json:"mobile"`
            Subtitle string  `json:"subtitle"`
            AccountHolder string  `json:"account_holder"`
            DelightsUserName string  `json:"delights_user_name"`
            BeneficiaryID string  `json:"beneficiary_id"`
         
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

        
            HashKey string  `json:"hash_key"`
            RequestID string  `json:"request_id"`
            Otp string  `json:"otp"`
         
    }
    
    // AddBeneficiaryViaOtpVerificationResponse ...
    type AddBeneficiaryViaOtpVerificationResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // WrongOtpError ...
    type WrongOtpError struct {

        
            Success string  `json:"success"`
            IsVerifiedFlag bool  `json:"is_verified_flag"`
            Description string  `json:"description"`
         
    }
    
    // BeneficiaryModeDetails ...
    type BeneficiaryModeDetails struct {

        
            Comment string  `json:"comment"`
            Vpa string  `json:"vpa"`
            Address string  `json:"address"`
            Email string  `json:"email"`
            IfscCode string  `json:"ifsc_code"`
            Wallet string  `json:"wallet"`
            BranchName string  `json:"branch_name"`
            AccountHolder string  `json:"account_holder"`
            Mobile string  `json:"mobile"`
            BankName string  `json:"bank_name"`
            AccountNo string  `json:"account_no"`
         
    }
    
    // AddBeneficiaryDetailsRequest ...
    type AddBeneficiaryDetailsRequest struct {

        
            RequestID string  `json:"request_id"`
            TransferMode string  `json:"transfer_mode"`
            OrderID string  `json:"order_id"`
            Delights bool  `json:"delights"`
            Otp string  `json:"otp"`
            ShipmentID string  `json:"shipment_id"`
            Details BeneficiaryModeDetails  `json:"details"`
         
    }
    
    // RefundAccountResponse ...
    type RefundAccountResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
            IsVerifiedFlag bool  `json:"is_verified_flag"`
            Data map[string]interface{}  `json:"data"`
         
    }
    
    // BankDetailsForOTP ...
    type BankDetailsForOTP struct {

        
            IfscCode string  `json:"ifsc_code"`
            BranchName string  `json:"branch_name"`
            AccountHolder string  `json:"account_holder"`
            BankName string  `json:"bank_name"`
            AccountNo string  `json:"account_no"`
         
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

        
            Success bool  `json:"success"`
            RequestID string  `json:"request_id"`
            IsVerifiedFlag string  `json:"is_verified_flag"`
         
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

        
            Message string  `json:"message"`
            PaymentLinkURL string  `json:"payment_link_url"`
            MerchantName string  `json:"merchant_name"`
            StatusCode float64  `json:"status_code"`
            PollingTimeout float64  `json:"polling_timeout"`
            Amount float64  `json:"amount"`
            ExternalOrderID string  `json:"external_order_id"`
            Success bool  `json:"success"`
            PaymentLinkCurrentStatus string  `json:"payment_link_current_status"`
         
    }
    
    // ErrorDescription ...
    type ErrorDescription struct {

        
            Cancelled bool  `json:"cancelled"`
            MerchantOrderID string  `json:"merchant_order_id"`
            PaymentTransactionID string  `json:"payment_transaction_id"`
            InvalidID bool  `json:"invalid_id"`
            Expired bool  `json:"expired"`
            Msg string  `json:"msg"`
            Amount float64  `json:"amount"`
            MerchantName string  `json:"merchant_name"`
         
    }
    
    // ErrorResponse ...
    type ErrorResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
            StatusCode float64  `json:"status_code"`
            Error ErrorDescription  `json:"error"`
         
    }
    
    // CreatePaymentLinkMeta ...
    type CreatePaymentLinkMeta struct {

        
            CartID string  `json:"cart_id"`
            AssignCardID string  `json:"assign_card_id"`
            Pincode string  `json:"pincode"`
            CheckoutMode string  `json:"checkout_mode"`
            Amount string  `json:"amount"`
         
    }
    
    // CreatePaymentLinkRequest ...
    type CreatePaymentLinkRequest struct {

        
            MobileNumber string  `json:"mobile_number"`
            Email string  `json:"email"`
            Amount float64  `json:"amount"`
            Meta CreatePaymentLinkMeta  `json:"meta"`
            ExternalOrderID string  `json:"external_order_id"`
            Description string  `json:"description"`
         
    }
    
    // CreatePaymentLinkResponse ...
    type CreatePaymentLinkResponse struct {

        
            Message string  `json:"message"`
            PaymentLinkID string  `json:"payment_link_id"`
            PaymentLinkURL string  `json:"payment_link_url"`
            StatusCode float64  `json:"status_code"`
            PollingTimeout float64  `json:"polling_timeout"`
            Success bool  `json:"success"`
         
    }
    
    // CancelOrResendPaymentLinkRequest ...
    type CancelOrResendPaymentLinkRequest struct {

        
            PaymentLinkID string  `json:"payment_link_id"`
         
    }
    
    // ResendPaymentLinkResponse ...
    type ResendPaymentLinkResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
            StatusCode float64  `json:"status_code"`
            PollingTimeout float64  `json:"polling_timeout"`
         
    }
    
    // CancelPaymentLinkResponse ...
    type CancelPaymentLinkResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
            StatusCode float64  `json:"status_code"`
         
    }
    
    // PollingPaymentLinkResponse ...
    type PollingPaymentLinkResponse struct {

        
            Message string  `json:"message"`
            PaymentLinkID string  `json:"payment_link_id"`
            OrderID string  `json:"order_id"`
            RedirectURL string  `json:"redirect_url"`
            Status string  `json:"status"`
            StatusCode float64  `json:"status_code"`
            HttpStatus float64  `json:"http_status"`
            Amount float64  `json:"amount"`
            Success bool  `json:"success"`
            AggregatorName string  `json:"aggregator_name"`
         
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
            Name string  `json:"name"`
            Meta PaymentMethodsMeta  `json:"meta"`
         
    }
    
    // CreateOrderUserRequest ...
    type CreateOrderUserRequest struct {

        
            PaymentLinkID string  `json:"payment_link_id"`
            PaymentMethods CreateOrderUserPaymentMethods  `json:"payment_methods"`
            SuccessCallbackURL string  `json:"success_callback_url"`
            FailureCallbackURL string  `json:"failure_callback_url"`
            Meta map[string]interface{}  `json:"meta"`
            Currency string  `json:"currency"`
         
    }
    
    // CreateOrderUserData ...
    type CreateOrderUserData struct {

        
            CallbackURL string  `json:"callback_url"`
            Method string  `json:"method"`
            CustomerID string  `json:"customer_id"`
            Email string  `json:"email"`
            OrderID string  `json:"order_id"`
            Aggregator string  `json:"aggregator"`
            Amount float64  `json:"amount"`
            Currency string  `json:"currency"`
            Contact string  `json:"contact"`
            MerchantOrderID string  `json:"merchant_order_id"`
         
    }
    
    // CreateOrderUserResponse ...
    type CreateOrderUserResponse struct {

        
            PaymentConfirmURL string  `json:"payment_confirm_url"`
            Message string  `json:"message"`
            CallbackURL string  `json:"callback_url"`
            OrderID string  `json:"order_id"`
            StatusCode float64  `json:"status_code"`
            Data CreateOrderUserData  `json:"data"`
            Success bool  `json:"success"`
         
    }
    
    // BalanceDetails ...
    type BalanceDetails struct {

        
            Value float64  `json:"value"`
            Currency string  `json:"currency"`
            FormattedValue string  `json:"formatted_value"`
         
    }
    
    // CreditSummary ...
    type CreditSummary struct {

        
            StatusMessage string  `json:"status_message"`
            MerchantCustomerRefID string  `json:"merchant_customer_ref_id"`
            Status string  `json:"status"`
            Balance BalanceDetails  `json:"balance"`
         
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

        
            IsRegistered bool  `json:"is_registered"`
            Status bool  `json:"status"`
            SignupURL string  `json:"signup_url"`
         
    }
    
    // CheckCreditResponse ...
    type CheckCreditResponse struct {

        
            Success bool  `json:"success"`
            Data CreditDetail  `json:"data"`
         
    }
    
    // MarketplaceInfo ...
    type MarketplaceInfo struct {

        
            DateOfJoining string  `json:"date_of_joining"`
            Name string  `json:"name"`
            MembershipID string  `json:"membership_id"`
         
    }
    
    // DeviceDetails ...
    type DeviceDetails struct {

        
            IdentifierType string  `json:"identifier_type"`
            DeviceType string  `json:"device_type"`
            IdentificationNumber string  `json:"identification_number"`
            Os string  `json:"os"`
            DeviceModel string  `json:"device_model"`
            DeviceMake string  `json:"device_make"`
            OsVersion string  `json:"os_version"`
         
    }
    
    // KYCAddress ...
    type KYCAddress struct {

        
            City string  `json:"city"`
            LandMark string  `json:"land_mark"`
            State string  `json:"state"`
            Pincode string  `json:"pincode"`
            OwnershipType string  `json:"ownership_type"`
            Addressline1 string  `json:"addressline1"`
            Addressline2 string  `json:"addressline2"`
         
    }
    
    // UserPersonalInfoInDetails ...
    type UserPersonalInfoInDetails struct {

        
            Dob string  `json:"dob"`
            Phone string  `json:"phone"`
            Email string  `json:"email"`
            LastName string  `json:"last_name"`
            MiddleName string  `json:"middle_name"`
            Gender string  `json:"gender"`
            MothersName string  `json:"mothers_name"`
            VoterID string  `json:"voter_id"`
            EmailVerified bool  `json:"email_verified"`
            Passport string  `json:"passport"`
            Pan string  `json:"pan"`
            MobileVerified bool  `json:"mobile_verified"`
            FirstName string  `json:"first_name"`
            AddressAsPerID KYCAddress  `json:"address_as_per_id"`
            DrivingLicense string  `json:"driving_license"`
            FathersName string  `json:"fathers_name"`
         
    }
    
    // BusinessDetails ...
    type BusinessDetails struct {

        
            Fda string  `json:"fda"`
            Address KYCAddress  `json:"address"`
            Fssai string  `json:"fssai"`
            Name string  `json:"name"`
            Gstin string  `json:"gstin"`
            BusinessType string  `json:"business_type"`
            Vintage string  `json:"vintage"`
            ShopAndEstablishment map[string]interface{}  `json:"shop_and_establishment"`
            EntityType string  `json:"entity_type"`
            Pan string  `json:"pan"`
            BusinessOwnershipType string  `json:"business_ownership_type"`
         
    }
    
    // CustomerOnboardingRequest ...
    type CustomerOnboardingRequest struct {

        
            MarketplaceInfo MarketplaceInfo  `json:"marketplace_info"`
            Device DeviceDetails  `json:"device"`
            PersonalInfo UserPersonalInfoInDetails  `json:"personal_info"`
            BusinessInfo BusinessDetails  `json:"business_info"`
            Mcc string  `json:"mcc"`
            Source string  `json:"source"`
            Aggregator string  `json:"aggregator"`
         
    }
    
    // OnboardSummary ...
    type OnboardSummary struct {

        
            Session map[string]interface{}  `json:"session"`
            RedirectURL string  `json:"redirect_url"`
            Status bool  `json:"status"`
         
    }
    
    // CustomerOnboardingResponse ...
    type CustomerOnboardingResponse struct {

        
            Success bool  `json:"success"`
            Data OnboardSummary  `json:"data"`
         
    }
    

    
    // OrderPage ...
    type OrderPage struct {

        
            Size float64  `json:"size"`
            Current float64  `json:"current"`
            Type string  `json:"type"`
            HasNext bool  `json:"has_next"`
            ItemTotal float64  `json:"item_total"`
         
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
    
    // BreakupValues ...
    type BreakupValues struct {

        
            Value float64  `json:"value"`
            Display string  `json:"display"`
            Name string  `json:"name"`
         
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
            TrackingDetails []NestedTrackingDetails  `json:"tracking_details"`
            IsCurrent bool  `json:"is_current"`
            IsPassed bool  `json:"is_passed"`
            Status string  `json:"status"`
         
    }
    
    // Prices ...
    type Prices struct {

        
            Cashback float64  `json:"cashback"`
            RefundAmount float64  `json:"refund_amount"`
            Discount float64  `json:"discount"`
            PromotionEffectiveDiscount float64  `json:"promotion_effective_discount"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            CodCharges float64  `json:"cod_charges"`
            CashbackApplied float64  `json:"cashback_applied"`
            PriceMarked float64  `json:"price_marked"`
            ValueOfGood float64  `json:"value_of_good"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            RefundCredit float64  `json:"refund_credit"`
            TransferPrice float64  `json:"transfer_price"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            AmountPaidRoundoff float64  `json:"amount_paid_roundoff"`
            CouponValue float64  `json:"coupon_value"`
            FyndCredits float64  `json:"fynd_credits"`
            DeliveryCharge float64  `json:"delivery_charge"`
            AmountPaid float64  `json:"amount_paid"`
            PriceEffective float64  `json:"price_effective"`
            AddedToFyndCash bool  `json:"added_to_fynd_cash"`
         
    }
    
    // ShipmentPayment ...
    type ShipmentPayment struct {

        
            PaymentMode string  `json:"payment_mode"`
            Logo string  `json:"logo"`
            Mode string  `json:"mode"`
            Mop string  `json:"mop"`
            Status string  `json:"status"`
            DisplayName string  `json:"display_name"`
         
    }
    
    // AppliedFreeArticles ...
    type AppliedFreeArticles struct {

        
            FreeGiftItemDetails map[string]interface{}  `json:"free_gift_item_details"`
            ArticleID string  `json:"article_id"`
            ParentItemIdentifier string  `json:"parent_item_identifier"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // AppliedPromos ...
    type AppliedPromos struct {

        
            MrpPromotion bool  `json:"mrp_promotion"`
            PromoID string  `json:"promo_id"`
            Amount float64  `json:"amount"`
            ArticleQuantity float64  `json:"article_quantity"`
            AppliedFreeArticles []AppliedFreeArticles  `json:"applied_free_articles"`
            PromotionType string  `json:"promotion_type"`
            PromotionName string  `json:"promotion_name"`
         
    }
    
    // Identifiers ...
    type Identifiers struct {

        
            SkuCode string  `json:"sku_code"`
            Ean string  `json:"ean"`
         
    }
    
    // FinancialBreakup ...
    type FinancialBreakup struct {

        
            ItemName string  `json:"item_name"`
            Cashback float64  `json:"cashback"`
            RefundAmount float64  `json:"refund_amount"`
            Discount float64  `json:"discount"`
            PromotionEffectiveDiscount float64  `json:"promotion_effective_discount"`
            Identifiers Identifiers  `json:"identifiers"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            HsnCode string  `json:"hsn_code"`
            CodCharges float64  `json:"cod_charges"`
            CashbackApplied float64  `json:"cashback_applied"`
            PriceMarked float64  `json:"price_marked"`
            ValueOfGood float64  `json:"value_of_good"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            RefundCredit float64  `json:"refund_credit"`
            TransferPrice float64  `json:"transfer_price"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            AmountPaidRoundoff float64  `json:"amount_paid_roundoff"`
            TotalUnits float64  `json:"total_units"`
            Size string  `json:"size"`
            CouponValue float64  `json:"coupon_value"`
            FyndCredits float64  `json:"fynd_credits"`
            GstFee float64  `json:"gst_fee"`
            DeliveryCharge float64  `json:"delivery_charge"`
            GstTag string  `json:"gst_tag"`
            AmountPaid float64  `json:"amount_paid"`
            PriceEffective float64  `json:"price_effective"`
            AddedToFyndCash bool  `json:"added_to_fynd_cash"`
         
    }
    
    // ItemBrand ...
    type ItemBrand struct {

        
            Name string  `json:"name"`
            Logo string  `json:"logo"`
         
    }
    
    // Item ...
    type Item struct {

        
            Size string  `json:"size"`
            ID float64  `json:"id"`
            SellerIdentifier string  `json:"seller_identifier"`
            Image []string  `json:"image"`
            Brand ItemBrand  `json:"brand"`
            Code string  `json:"code"`
            Name string  `json:"name"`
            SlugKey string  `json:"slug_key"`
         
    }
    
    // CurrentStatus ...
    type CurrentStatus struct {

        
            JourneyType string  `json:"journey_type"`
            UpdatedAt string  `json:"updated_at"`
            Name string  `json:"name"`
            Status string  `json:"status"`
         
    }
    
    // Bags ...
    type Bags struct {

        
            ReturnableDate string  `json:"returnable_date"`
            CanReturn bool  `json:"can_return"`
            AppliedPromos []AppliedPromos  `json:"applied_promos"`
            FinancialBreakup []FinancialBreakup  `json:"financial_breakup"`
            Prices Prices  `json:"prices"`
            Quantity float64  `json:"quantity"`
            ParentPromoBags map[string]interface{}  `json:"parent_promo_bags"`
            Meta map[string]interface{}  `json:"meta"`
            LineNumber float64  `json:"line_number"`
            Item Item  `json:"item"`
            DeliveryDate string  `json:"delivery_date"`
            SellerIdentifier string  `json:"seller_identifier"`
            CurrentStatus CurrentStatus  `json:"current_status"`
            CanCancel bool  `json:"can_cancel"`
            ID float64  `json:"id"`
         
    }
    
    // DeliveryAddress ...
    type DeliveryAddress struct {

        
            CreatedAt string  `json:"created_at"`
            Email string  `json:"email"`
            Landmark string  `json:"landmark"`
            Address2 string  `json:"address2"`
            Address string  `json:"address"`
            UpdatedAt string  `json:"updated_at"`
            Address1 string  `json:"address1"`
            Pincode string  `json:"pincode"`
            Area string  `json:"area"`
            Version string  `json:"version"`
            Name string  `json:"name"`
            Country string  `json:"country"`
            City string  `json:"city"`
            AddressCategory string  `json:"address_category"`
            AddressType string  `json:"address_type"`
            ContactPerson string  `json:"contact_person"`
            Latitude float64  `json:"latitude"`
            State string  `json:"state"`
            Longitude float64  `json:"longitude"`
            Phone string  `json:"phone"`
         
    }
    
    // FulfillingStore ...
    type FulfillingStore struct {

        
            CompanyName string  `json:"company_name"`
            CompanyID float64  `json:"company_id"`
            Code string  `json:"code"`
            ID float64  `json:"id"`
            Name string  `json:"name"`
         
    }
    
    // ShipmentStatus ...
    type ShipmentStatus struct {

        
            HexCode string  `json:"hex_code"`
            Title string  `json:"title"`
         
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
    
    // ShipmentTotalDetails ...
    type ShipmentTotalDetails struct {

        
            Pieces float64  `json:"pieces"`
            TotalPrice float64  `json:"total_price"`
            Sizes float64  `json:"sizes"`
         
    }
    
    // Invoice ...
    type Invoice struct {

        
            InvoiceURL string  `json:"invoice_url"`
            LabelURL string  `json:"label_url"`
            UpdatedDate string  `json:"updated_date"`
         
    }
    
    // FulfillingCompany ...
    type FulfillingCompany struct {

        
            Name string  `json:"name"`
            ID float64  `json:"id"`
         
    }
    
    // ShipmentUserInfo ...
    type ShipmentUserInfo struct {

        
            LastName string  `json:"last_name"`
            FirstName string  `json:"first_name"`
            Mobile string  `json:"mobile"`
            Gender string  `json:"gender"`
         
    }
    
    // Shipments ...
    type Shipments struct {

        
            CanReturn bool  `json:"can_return"`
            TrackingDetails []TrackingDetails  `json:"tracking_details"`
            Prices Prices  `json:"prices"`
            Payment ShipmentPayment  `json:"payment"`
            Bags []Bags  `json:"bags"`
            DeliveryAddress DeliveryAddress  `json:"delivery_address"`
            TotalBags float64  `json:"total_bags"`
            FulfillingStore FulfillingStore  `json:"fulfilling_store"`
            Comment string  `json:"comment"`
            ShipmentID string  `json:"shipment_id"`
            CustomMeta []map[string]interface{}  `json:"custom_meta"`
            ShipmentStatus ShipmentStatus  `json:"shipment_status"`
            BeneficiaryDetails bool  `json:"beneficiary_details"`
            ShowDownloadInvoice bool  `json:"show_download_invoice"`
            ShowTrackLink bool  `json:"show_track_link"`
            Promise Promise  `json:"promise"`
            TotalDetails ShipmentTotalDetails  `json:"total_details"`
            Invoice Invoice  `json:"invoice"`
            OrderID string  `json:"order_id"`
            FulfillingCompany FulfillingCompany  `json:"fulfilling_company"`
            NeedHelpURL string  `json:"need_help_url"`
            OrderType string  `json:"order_type"`
            CanCancel bool  `json:"can_cancel"`
            RefundDetails map[string]interface{}  `json:"refund_details"`
            ReturnableDate string  `json:"returnable_date"`
            BreakupValues []BreakupValues  `json:"breakup_values"`
            TrakingNo string  `json:"traking_no"`
            CanBreak map[string]interface{}  `json:"can_break"`
            UserInfo ShipmentUserInfo  `json:"user_info"`
            DeliveryDate string  `json:"delivery_date"`
            ShipmentCreatedAt string  `json:"shipment_created_at"`
            TrackURL string  `json:"track_url"`
            SizeInfo map[string]interface{}  `json:"size_info"`
            DpName string  `json:"dp_name"`
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

        
            ItemID float64  `json:"item_id"`
            SellerID float64  `json:"seller_id"`
            ItemSize string  `json:"item_size"`
            StoreID float64  `json:"store_id"`
            Quantity float64  `json:"quantity"`
            ArticleAssignment BagsForReorderArticleAssignment  `json:"article_assignment"`
         
    }
    
    // OrderSchema ...
    type OrderSchema struct {

        
            OrderCreatedTime string  `json:"order_created_time"`
            BreakupValues []BreakupValues  `json:"breakup_values"`
            TotalShipmentsInOrder float64  `json:"total_shipments_in_order"`
            Shipments []Shipments  `json:"shipments"`
            UserInfo UserInfo  `json:"user_info"`
            OrderID string  `json:"order_id"`
            BagsForReorder []BagsForReorder  `json:"bags_for_reorder"`
         
    }
    
    // OrderList ...
    type OrderList struct {

        
            Page OrderPage  `json:"page"`
            Filters OrderFilters  `json:"filters"`
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

        
            ShipmentID string  `json:"shipment_id"`
            PresignedURL string  `json:"presigned_url"`
            Success bool  `json:"success"`
            PresignedType string  `json:"presigned_type"`
         
    }
    
    // Track ...
    type Track struct {

        
            Reason string  `json:"reason"`
            UpdatedAt string  `json:"updated_at"`
            AccountName string  `json:"account_name"`
            Awb string  `json:"awb"`
            Status string  `json:"status"`
            LastLocationRecievedAt string  `json:"last_location_recieved_at"`
            ShipmentType string  `json:"shipment_type"`
            UpdatedTime string  `json:"updated_time"`
         
    }
    
    // ShipmentTrack ...
    type ShipmentTrack struct {

        
            Results []Track  `json:"results"`
         
    }
    
    // CustomerDetailsResponse ...
    type CustomerDetailsResponse struct {

        
            ShipmentID string  `json:"shipment_id"`
            Country string  `json:"country"`
            OrderID string  `json:"order_id"`
            Phone string  `json:"phone"`
            Name string  `json:"name"`
         
    }
    
    // SendOtpToCustomerResponse ...
    type SendOtpToCustomerResponse struct {

        
            Message string  `json:"message"`
            ResendTimer float64  `json:"resend_timer"`
            Success bool  `json:"success"`
            RequestID string  `json:"request_id"`
         
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

        
            QuestionSet []QuestionSet  `json:"question_set"`
            Meta BagReasonMeta  `json:"meta"`
            Reasons []BagReasons  `json:"reasons"`
            QcType []string  `json:"qc_type"`
            ID float64  `json:"id"`
            DisplayName string  `json:"display_name"`
         
    }
    
    // ShipmentBagReasons ...
    type ShipmentBagReasons struct {

        
            Reasons []BagReasons  `json:"reasons"`
            Success bool  `json:"success"`
         
    }
    
    // ShipmentReason ...
    type ShipmentReason struct {

        
            Flow string  `json:"flow"`
            Priority float64  `json:"priority"`
            FeedbackType string  `json:"feedback_type"`
            ReasonText string  `json:"reason_text"`
            ShowTextArea bool  `json:"show_text_area"`
            ReasonID float64  `json:"reason_id"`
         
    }
    
    // ShipmentReasons ...
    type ShipmentReasons struct {

        
            Reasons []ShipmentReason  `json:"reasons"`
         
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
    
    // ReasonsData ...
    type ReasonsData struct {

        
            Products []ProductsReasons  `json:"products"`
            Entities []EntitiesReasons  `json:"entities"`
         
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

        
            Products []ProductsDataUpdates1  `json:"products"`
            Entities []EntitiesDataUpdates  `json:"entities"`
         
    }
    
    // ShipmentsRequest ...
    type ShipmentsRequest struct {

        
            Reasons ReasonsData  `json:"reasons"`
            Products []Products  `json:"products"`
            Identifier string  `json:"identifier"`
            DataUpdates DataUpdates  `json:"data_updates"`
         
    }
    
    // StatuesRequest ...
    type StatuesRequest struct {

        
            Shipments []ShipmentsRequest1  `json:"shipments"`
            Status string  `json:"status"`
            ExcludeBagsNextState string  `json:"exclude_bags_next_state"`
         
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

        
            Exception string  `json:"exception"`
            StackTrace string  `json:"stack_trace"`
            Code string  `json:"code"`
            Status float64  `json:"status"`
            Message string  `json:"message"`
         
    }
    
    // Products1 ...
    type Products1 struct {

        
            Quantity float64  `json:"quantity"`
            Identifier string  `json:"identifier"`
            LineNumber float64  `json:"line_number"`
         
    }
    
    // ProductsReasonsData1 ...
    type ProductsReasonsData1 struct {

        
            ReasonText string  `json:"reason_text"`
            ReasonID float64  `json:"reason_id"`
         
    }
    
    // ProductsReasonsFilters1 ...
    type ProductsReasonsFilters1 struct {

        
            Quantity float64  `json:"quantity"`
            Identifier string  `json:"identifier"`
            LineNumber float64  `json:"line_number"`
         
    }
    
    // ProductsReasons1 ...
    type ProductsReasons1 struct {

        
            Data ProductsReasonsData1  `json:"data"`
            Filters []ProductsReasonsFilters1  `json:"filters"`
         
    }
    
    // EntityReasonData1 ...
    type EntityReasonData1 struct {

        
            ReasonText string  `json:"reason_text"`
            ReasonID float64  `json:"reason_id"`
         
    }
    
    // EntitiesReasons1 ...
    type EntitiesReasons1 struct {

        
            Data EntityReasonData1  `json:"data"`
            Filters []map[string]interface{}  `json:"filters"`
         
    }
    
    // ReasonsData1 ...
    type ReasonsData1 struct {

        
            Products []ProductsReasons1  `json:"products"`
            Entities []EntitiesReasons1  `json:"entities"`
         
    }
    
    // ProductsDataUpdatesFilters1 ...
    type ProductsDataUpdatesFilters1 struct {

        
            Identifier string  `json:"identifier"`
            LineNumber float64  `json:"line_number"`
         
    }
    
    // ProductsDataUpdates1 ...
    type ProductsDataUpdates1 struct {

        
            Data map[string]interface{}  `json:"data"`
            Filters []ProductsDataUpdatesFilters1  `json:"filters"`
         
    }
    
    // ShipmentsRequest1 ...
    type ShipmentsRequest1 struct {

        
            Products []Products1  `json:"products"`
            Reasons ReasonsData1  `json:"reasons"`
            Identifier string  `json:"identifier"`
            DataUpdates DataUpdates  `json:"data_updates"`
         
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
    
    // DisplayBreakup ...
    type DisplayBreakup struct {

        
            Message []string  `json:"message"`
            Display string  `json:"display"`
            Value float64  `json:"value"`
            Key string  `json:"key"`
            CurrencySymbol string  `json:"currency_symbol"`
            CurrencyCode string  `json:"currency_code"`
         
    }
    
    // LoyaltyPoints ...
    type LoyaltyPoints struct {

        
            Total float64  `json:"total"`
            Applicable float64  `json:"applicable"`
            Description string  `json:"description"`
            IsApplied bool  `json:"is_applied"`
         
    }
    
    // RawBreakup ...
    type RawBreakup struct {

        
            Total float64  `json:"total"`
            MrpTotal float64  `json:"mrp_total"`
            ConvenienceFee float64  `json:"convenience_fee"`
            CodCharge float64  `json:"cod_charge"`
            Subtotal float64  `json:"subtotal"`
            YouSaved float64  `json:"you_saved"`
            Discount float64  `json:"discount"`
            Coupon float64  `json:"coupon"`
            FyndCash float64  `json:"fynd_cash"`
            GstCharges float64  `json:"gst_charges"`
            Vog float64  `json:"vog"`
            DeliveryCharge float64  `json:"delivery_charge"`
         
    }
    
    // CouponBreakup ...
    type CouponBreakup struct {

        
            SubTitle string  `json:"sub_title"`
            CouponValue float64  `json:"coupon_value"`
            Message string  `json:"message"`
            Type string  `json:"type"`
            CouponType string  `json:"coupon_type"`
            MaxDiscountValue float64  `json:"max_discount_value"`
            MinimumCartValue float64  `json:"minimum_cart_value"`
            Title string  `json:"title"`
            Value float64  `json:"value"`
            Description string  `json:"description"`
            IsApplied bool  `json:"is_applied"`
            Code string  `json:"code"`
            UID string  `json:"uid"`
         
    }
    
    // CartBreakup ...
    type CartBreakup struct {

        
            Display []DisplayBreakup  `json:"display"`
            LoyaltyPoints LoyaltyPoints  `json:"loyalty_points"`
            Raw RawBreakup  `json:"raw"`
            Coupon CouponBreakup  `json:"coupon"`
         
    }
    
    // PaymentSelectionLock ...
    type PaymentSelectionLock struct {

        
            Enabled bool  `json:"enabled"`
            PaymentIdentifier string  `json:"payment_identifier"`
            DefaultOptions string  `json:"default_options"`
         
    }
    
    // CartCurrency ...
    type CartCurrency struct {

        
            Code string  `json:"code"`
            Symbol string  `json:"symbol"`
         
    }
    
    // ProductPrice ...
    type ProductPrice struct {

        
            Marked float64  `json:"marked"`
            Effective float64  `json:"effective"`
            AddOn float64  `json:"add_on"`
            CurrencySymbol string  `json:"currency_symbol"`
            Selling float64  `json:"selling"`
            CurrencyCode string  `json:"currency_code"`
         
    }
    
    // ProductPriceInfo ...
    type ProductPriceInfo struct {

        
            Converted ProductPrice  `json:"converted"`
            Base ProductPrice  `json:"base"`
         
    }
    
    // ProductAvailability ...
    type ProductAvailability struct {

        
            Sizes []string  `json:"sizes"`
            OtherStoreQuantity float64  `json:"other_store_quantity"`
            OutOfStock bool  `json:"out_of_stock"`
            Deliverable bool  `json:"deliverable"`
            IsValid bool  `json:"is_valid"`
         
    }
    
    // PromoMeta ...
    type PromoMeta struct {

        
            Message string  `json:"message"`
         
    }
    
    // BuyRules ...
    type BuyRules struct {

        
            ItemCriteria map[string]interface{}  `json:"item_criteria"`
            CartConditions map[string]interface{}  `json:"cart_conditions"`
         
    }
    
    // DiscountRulesApp ...
    type DiscountRulesApp struct {

        
            RawOffer map[string]interface{}  `json:"raw_offer"`
            ItemCriteria map[string]interface{}  `json:"item_criteria"`
            MatchedBuyRules []string  `json:"matched_buy_rules"`
            Offer map[string]interface{}  `json:"offer"`
         
    }
    
    // FreeGiftItem ...
    type FreeGiftItem struct {

        
            ItemSlug string  `json:"item_slug"`
            ItemID float64  `json:"item_id"`
            ItemName string  `json:"item_name"`
            ItemBrandName string  `json:"item_brand_name"`
            ItemImagesURL []string  `json:"item_images_url"`
            ItemPriceDetails map[string]interface{}  `json:"item_price_details"`
         
    }
    
    // AppliedFreeArticles ...
    type AppliedFreeArticles struct {

        
            ParentItemIdentifier string  `json:"parent_item_identifier"`
            FreeGiftItemDetails FreeGiftItem  `json:"free_gift_item_details"`
            Quantity float64  `json:"quantity"`
            ArticleID string  `json:"article_id"`
         
    }
    
    // AppliedPromotion ...
    type AppliedPromotion struct {

        
            MrpPromotion bool  `json:"mrp_promotion"`
            PromotionType string  `json:"promotion_type"`
            PromotionName string  `json:"promotion_name"`
            ArticleQuantity float64  `json:"article_quantity"`
            BuyRules []BuyRules  `json:"buy_rules"`
            DiscountRules []DiscountRulesApp  `json:"discount_rules"`
            PromoID string  `json:"promo_id"`
            PromotionGroup string  `json:"promotion_group"`
            OfferText string  `json:"offer_text"`
            AppliedFreeArticles []AppliedFreeArticles  `json:"applied_free_articles"`
            Amount float64  `json:"amount"`
         
    }
    
    // ProductImage ...
    type ProductImage struct {

        
            AspectRatio string  `json:"aspect_ratio"`
            URL string  `json:"url"`
            SecureURL string  `json:"secure_url"`
         
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

        
            Query ActionQuery  `json:"query"`
            URL string  `json:"url"`
            Type string  `json:"type"`
         
    }
    
    // BaseInfo ...
    type BaseInfo struct {

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // CartProduct ...
    type CartProduct struct {

        
            Images []ProductImage  `json:"images"`
            Slug string  `json:"slug"`
            Type string  `json:"type"`
            Categories []CategoryInfo  `json:"categories"`
            Action ProductAction  `json:"action"`
            Brand BaseInfo  `json:"brand"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // CartProductIdentifer ...
    type CartProductIdentifer struct {

        
            Identifier string  `json:"identifier"`
         
    }
    
    // BasePrice ...
    type BasePrice struct {

        
            Marked float64  `json:"marked"`
            CurrencySymbol string  `json:"currency_symbol"`
            Effective float64  `json:"effective"`
            CurrencyCode string  `json:"currency_code"`
         
    }
    
    // ArticlePriceInfo ...
    type ArticlePriceInfo struct {

        
            Converted BasePrice  `json:"converted"`
            Base BasePrice  `json:"base"`
         
    }
    
    // ProductArticle ...
    type ProductArticle struct {

        
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            Quantity float64  `json:"quantity"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            Type string  `json:"type"`
            Price ArticlePriceInfo  `json:"price"`
            Store BaseInfo  `json:"store"`
            Identifier map[string]interface{}  `json:"identifier"`
            ProductGroupTags []string  `json:"product_group_tags"`
            Size string  `json:"size"`
            Seller BaseInfo  `json:"seller"`
            SellerIdentifier string  `json:"seller_identifier"`
            UID string  `json:"uid"`
         
    }
    
    // CartProductInfo ...
    type CartProductInfo struct {

        
            Quantity float64  `json:"quantity"`
            CouponMessage string  `json:"coupon_message"`
            PricePerUnit ProductPriceInfo  `json:"price_per_unit"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            Availability ProductAvailability  `json:"availability"`
            Message string  `json:"message"`
            PromoMeta PromoMeta  `json:"promo_meta"`
            IsSet bool  `json:"is_set"`
            PromotionsApplied []AppliedPromotion  `json:"promotions_applied"`
            Discount string  `json:"discount"`
            Price ProductPriceInfo  `json:"price"`
            Product CartProduct  `json:"product"`
            BulkOffer map[string]interface{}  `json:"bulk_offer"`
            Key string  `json:"key"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
            Article ProductArticle  `json:"article"`
         
    }
    
    // CartDetailResponse ...
    type CartDetailResponse struct {

        
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            Message string  `json:"message"`
            Gstin string  `json:"gstin"`
            CouponText string  `json:"coupon_text"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            Currency CartCurrency  `json:"currency"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            BuyNow bool  `json:"buy_now"`
            Comment string  `json:"comment"`
            ID string  `json:"id"`
            CheckoutMode string  `json:"checkout_mode"`
            LastModified string  `json:"last_modified"`
            IsValid bool  `json:"is_valid"`
            Items []CartProductInfo  `json:"items"`
         
    }
    
    // AddProductCart ...
    type AddProductCart struct {

        
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            Quantity float64  `json:"quantity"`
            ItemID float64  `json:"item_id"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            SellerID float64  `json:"seller_id"`
            ArticleID string  `json:"article_id"`
            Display string  `json:"display"`
            ItemSize string  `json:"item_size"`
            ProductGroupTags []string  `json:"product_group_tags"`
            StoreID float64  `json:"store_id"`
            Pos bool  `json:"pos"`
            ArticleAssignment map[string]interface{}  `json:"article_assignment"`
         
    }
    
    // AddCartRequest ...
    type AddCartRequest struct {

        
            Items []AddProductCart  `json:"items"`
         
    }
    
    // AddCartDetailResponse ...
    type AddCartDetailResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
            Partial bool  `json:"partial"`
            Cart CartDetailResponse  `json:"cart"`
         
    }
    
    // UpdateProductCart ...
    type UpdateProductCart struct {

        
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            Quantity float64  `json:"quantity"`
            ItemIndex float64  `json:"item_index"`
            ItemID float64  `json:"item_id"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            ArticleID string  `json:"article_id"`
            ItemSize string  `json:"item_size"`
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
            Success bool  `json:"success"`
            Cart CartDetailResponse  `json:"cart"`
         
    }
    
    // CartItemCountResponse ...
    type CartItemCountResponse struct {

        
            UserCartItemsCount float64  `json:"user_cart_items_count"`
         
    }
    
    // Coupon ...
    type Coupon struct {

        
            SubTitle string  `json:"sub_title"`
            CouponValue float64  `json:"coupon_value"`
            Message string  `json:"message"`
            MinimumCartValue float64  `json:"minimum_cart_value"`
            MaxDiscountValue float64  `json:"max_discount_value"`
            IsApplicable bool  `json:"is_applicable"`
            Title string  `json:"title"`
            ExpiresOn string  `json:"expires_on"`
            Description string  `json:"description"`
            IsApplied bool  `json:"is_applied"`
            CouponType string  `json:"coupon_type"`
            CouponCode string  `json:"coupon_code"`
         
    }
    
    // PageCoupon ...
    type PageCoupon struct {

        
            Total float64  `json:"total"`
            HasPrevious bool  `json:"has_previous"`
            TotalItemCount float64  `json:"total_item_count"`
            Current float64  `json:"current"`
            HasNext bool  `json:"has_next"`
         
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
            Effective float64  `json:"effective"`
            BulkEffective float64  `json:"bulk_effective"`
            CurrencySymbol string  `json:"currency_symbol"`
            CurrencyCode string  `json:"currency_code"`
         
    }
    
    // OfferItem ...
    type OfferItem struct {

        
            Quantity float64  `json:"quantity"`
            Total float64  `json:"total"`
            Type string  `json:"type"`
            AutoApplied bool  `json:"auto_applied"`
            Price OfferPrice  `json:"price"`
            Margin float64  `json:"margin"`
            Best bool  `json:"best"`
         
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
            UserID string  `json:"user_id"`
            CountryCode string  `json:"country_code"`
            Phone string  `json:"phone"`
            AreaCode string  `json:"area_code"`
            Tags []string  `json:"tags"`
            Email string  `json:"email"`
            City string  `json:"city"`
            GoogleMapPoint map[string]interface{}  `json:"google_map_point"`
            ID string  `json:"id"`
            Address string  `json:"address"`
            AddressType string  `json:"address_type"`
            Landmark string  `json:"landmark"`
            State string  `json:"state"`
            Area string  `json:"area"`
            CheckoutMode string  `json:"checkout_mode"`
            AreaCodeSlug string  `json:"area_code_slug"`
            Country string  `json:"country"`
            Meta map[string]interface{}  `json:"meta"`
            GeoLocation GeoLocation  `json:"geo_location"`
            IsDefaultAddress bool  `json:"is_default_address"`
            Name string  `json:"name"`
         
    }
    
    // GetAddressesResponse ...
    type GetAddressesResponse struct {

        
            Address []Address  `json:"address"`
         
    }
    
    // SaveAddressResponse ...
    type SaveAddressResponse struct {

        
            ID string  `json:"id"`
            Success bool  `json:"success"`
            IsDefaultAddress bool  `json:"is_default_address"`
         
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

        
            AddressID string  `json:"address_id"`
            MerchantCode string  `json:"merchant_code"`
            PaymentMode string  `json:"payment_mode"`
            ID string  `json:"id"`
            PaymentIdentifier string  `json:"payment_identifier"`
            AggregatorName string  `json:"aggregator_name"`
         
    }
    
    // CouponValidity ...
    type CouponValidity struct {

        
            DisplayMessageEn string  `json:"display_message_en"`
            Discount float64  `json:"discount"`
            Title string  `json:"title"`
            Code string  `json:"code"`
            Valid bool  `json:"valid"`
         
    }
    
    // PaymentCouponValidate ...
    type PaymentCouponValidate struct {

        
            Success bool  `json:"success"`
            CouponValidity CouponValidity  `json:"coupon_validity"`
            Message string  `json:"message"`
         
    }
    
    // ShipmentResponse ...
    type ShipmentResponse struct {

        
            BoxType string  `json:"box_type"`
            ShipmentType string  `json:"shipment_type"`
            Shipments float64  `json:"shipments"`
            DpID string  `json:"dp_id"`
            DpOptions map[string]interface{}  `json:"dp_options"`
            FulfillmentID float64  `json:"fulfillment_id"`
            Promise ShipmentPromise  `json:"promise"`
            FulfillmentType string  `json:"fulfillment_type"`
            OrderType string  `json:"order_type"`
            Items []CartProductInfo  `json:"items"`
         
    }
    
    // CartShipmentsResponse ...
    type CartShipmentsResponse struct {

        
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            Shipments []ShipmentResponse  `json:"shipments"`
            Currency CartCurrency  `json:"currency"`
            CheckoutMode string  `json:"checkout_mode"`
            Error bool  `json:"error"`
            Message string  `json:"message"`
            Gstin string  `json:"gstin"`
            CouponText string  `json:"coupon_text"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            CartID float64  `json:"cart_id"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            ID string  `json:"id"`
            Comment string  `json:"comment"`
            UID string  `json:"uid"`
            LastModified string  `json:"last_modified"`
            IsValid bool  `json:"is_valid"`
            BuyNow bool  `json:"buy_now"`
         
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

        
            LastName string  `json:"last_name"`
            ID string  `json:"_id"`
            FirstName string  `json:"first_name"`
            User string  `json:"user"`
         
    }
    
    // Files ...
    type Files struct {

        
            Values []string  `json:"values"`
            Key string  `json:"key"`
         
    }
    
    // CartPosCheckoutDetailRequest ...
    type CartPosCheckoutDetailRequest struct {

        
            PaymentMode string  `json:"payment_mode"`
            Pos bool  `json:"pos"`
            PickAtStoreUID float64  `json:"pick_at_store_uid"`
            BillingAddressID string  `json:"billing_address_id"`
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
            PaymentIdentifier string  `json:"payment_identifier"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            PaymentParams map[string]interface{}  `json:"payment_params"`
            CallbackURL string  `json:"callback_url"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
            AddressID string  `json:"address_id"`
            MerchantCode string  `json:"merchant_code"`
            OrderType string  `json:"order_type"`
            Staff StaffCheckout  `json:"staff"`
            Files []Files  `json:"files"`
            Aggregator string  `json:"aggregator"`
            Meta map[string]interface{}  `json:"meta"`
            OrderingStore float64  `json:"ordering_store"`
            PaymentAutoConfirm bool  `json:"payment_auto_confirm"`
         
    }
    
    // CheckCart ...
    type CheckCart struct {

        
            BreakupValues CartBreakup  `json:"breakup_values"`
            UserType string  `json:"user_type"`
            StoreEmps []map[string]interface{}  `json:"store_emps"`
            CouponText string  `json:"coupon_text"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            ErrorMessage string  `json:"error_message"`
            ID string  `json:"id"`
            IsValid bool  `json:"is_valid"`
            BuyNow bool  `json:"buy_now"`
            Items []CartProductInfo  `json:"items"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            CodAvailable bool  `json:"cod_available"`
            DeliveryCharges float64  `json:"delivery_charges"`
            Currency CartCurrency  `json:"currency"`
            CodCharges float64  `json:"cod_charges"`
            CheckoutMode string  `json:"checkout_mode"`
            OrderID string  `json:"order_id"`
            DeliveryChargeOrderValue float64  `json:"delivery_charge_order_value"`
            StoreCode string  `json:"store_code"`
            Message string  `json:"message"`
            Gstin string  `json:"gstin"`
            CartID float64  `json:"cart_id"`
            Comment string  `json:"comment"`
            CodMessage string  `json:"cod_message"`
            Success bool  `json:"success"`
            UID string  `json:"uid"`
            LastModified string  `json:"last_modified"`
         
    }
    
    // CartCheckoutResponse ...
    type CartCheckoutResponse struct {

        
            Message string  `json:"message"`
            AppInterceptURL string  `json:"app_intercept_url"`
            CallbackURL string  `json:"callback_url"`
            PaymentConfirmURL string  `json:"payment_confirm_url"`
            OrderID string  `json:"order_id"`
            Data map[string]interface{}  `json:"data"`
            Success bool  `json:"success"`
            Cart CheckCart  `json:"cart"`
         
    }
    
    // CartMetaRequest ...
    type CartMetaRequest struct {

        
            PickUpCustomerDetails map[string]interface{}  `json:"pick_up_customer_details"`
            CheckoutMode string  `json:"checkout_mode"`
            Comment string  `json:"comment"`
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

        
            AddressType string  `json:"address_type"`
            Address string  `json:"address"`
            Pincode float64  `json:"pincode"`
            StoreCode string  `json:"store_code"`
            Phone string  `json:"phone"`
            AreaCode string  `json:"area_code"`
            ID float64  `json:"id"`
            Country string  `json:"country"`
            Email string  `json:"email"`
            City string  `json:"city"`
            Landmark string  `json:"landmark"`
            State string  `json:"state"`
            AreaCodeSlug string  `json:"area_code_slug"`
            Area string  `json:"area"`
            UID float64  `json:"uid"`
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
            Meta map[string]interface{}  `json:"meta"`
            User map[string]interface{}  `json:"user"`
            Source map[string]interface{}  `json:"source"`
            Token string  `json:"token"`
         
    }
    
    // SharedCart ...
    type SharedCart struct {

        
            BreakupValues CartBreakup  `json:"breakup_values"`
            CouponText string  `json:"coupon_text"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            ID string  `json:"id"`
            IsValid bool  `json:"is_valid"`
            BuyNow bool  `json:"buy_now"`
            Items []CartProductInfo  `json:"items"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            Currency CartCurrency  `json:"currency"`
            CheckoutMode string  `json:"checkout_mode"`
            SharedCartDetails SharedCartDetails  `json:"shared_cart_details"`
            Message string  `json:"message"`
            Gstin string  `json:"gstin"`
            CartID float64  `json:"cart_id"`
            Comment string  `json:"comment"`
            UID string  `json:"uid"`
            LastModified string  `json:"last_modified"`
         
    }
    
    // SharedCartResponse ...
    type SharedCartResponse struct {

        
            Error string  `json:"error"`
            Cart SharedCart  `json:"cart"`
         
    }
    

    
    // PincodeParentsResponse ...
    type PincodeParentsResponse struct {

        
            Name string  `json:"name"`
            UID string  `json:"uid"`
            SubType string  `json:"sub_type"`
            DisplayName string  `json:"display_name"`
         
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
    
    // PincodeMetaResponse ...
    type PincodeMetaResponse struct {

        
            Zone string  `json:"zone"`
            InternalZoneID float64  `json:"internal_zone_id"`
         
    }
    
    // PincodeDataResponse ...
    type PincodeDataResponse struct {

        
            Parents []PincodeParentsResponse  `json:"parents"`
            DisplayName string  `json:"display_name"`
            SubType string  `json:"sub_type"`
            LatLong PincodeLatLongData  `json:"lat_long"`
            Error PincodeErrorSchemaResponse  `json:"error"`
            Name string  `json:"name"`
            Meta PincodeMetaResponse  `json:"meta"`
            UID string  `json:"uid"`
         
    }
    
    // PincodeApiResponse ...
    type PincodeApiResponse struct {

        
            Data []PincodeDataResponse  `json:"data"`
            Success bool  `json:"success"`
            Error PincodeErrorSchemaResponse  `json:"error"`
         
    }
    
    // TATCategoryRequest ...
    type TATCategoryRequest struct {

        
            Level string  `json:"level"`
            ID float64  `json:"id"`
         
    }
    
    // TATArticlesRequest ...
    type TATArticlesRequest struct {

        
            ManufacturingTimeUnit string  `json:"manufacturing_time_unit"`
            Category TATCategoryRequest  `json:"category"`
            ManufacturingTime float64  `json:"manufacturing_time"`
         
    }
    
    // TATLocationDetailsRequest ...
    type TATLocationDetailsRequest struct {

        
            FromPincode string  `json:"from_pincode"`
            FulfillmentID float64  `json:"fulfillment_id"`
            Articles []TATArticlesRequest  `json:"articles"`
         
    }
    
    // TATViewRequest ...
    type TATViewRequest struct {

        
            Identifier string  `json:"identifier"`
            ToPincode string  `json:"to_pincode"`
            Action string  `json:"action"`
            Source string  `json:"source"`
            LocationDetails []TATLocationDetailsRequest  `json:"location_details"`
            Journey string  `json:"journey"`
         
    }
    
    // TATErrorSchemaResponse ...
    type TATErrorSchemaResponse struct {

        
            Message string  `json:"message"`
            Type string  `json:"type"`
            Value string  `json:"value"`
         
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

        
            ManufacturingTimeSeconds float64  `json:"_manufacturing_time_seconds"`
            Promise TATPromiseResponse  `json:"promise"`
            ManufacturingTime float64  `json:"manufacturing_time"`
            Error TATErrorSchemaResponse  `json:"error"`
            ManufacturingTimeUnit string  `json:"manufacturing_time_unit"`
            IsCodAvailable bool  `json:"is_cod_available"`
            Category TATCategoryRequest  `json:"category"`
         
    }
    
    // TATLocationDetailsResponse ...
    type TATLocationDetailsResponse struct {

        
            FromPincode string  `json:"from_pincode"`
            FulfillmentID float64  `json:"fulfillment_id"`
            Articles []TATArticlesResponse  `json:"articles"`
         
    }
    
    // TATViewResponse ...
    type TATViewResponse struct {

        
            ToCity string  `json:"to_city"`
            PaymentMode string  `json:"payment_mode"`
            Identifier string  `json:"identifier"`
            StormbreakerUUID string  `json:"stormbreaker_uuid"`
            Error TATErrorSchemaResponse  `json:"error"`
            ToPincode string  `json:"to_pincode"`
            Action string  `json:"action"`
            Source string  `json:"source"`
            LocationDetails []TATLocationDetailsResponse  `json:"location_details"`
            IsCodAvailable bool  `json:"is_cod_available"`
            RequestUUID string  `json:"request_uuid"`
            Journey string  `json:"journey"`
            Success bool  `json:"success"`
         
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
    
