package application



    
    // Meta ...
    type Meta struct {

        
            Source string  `json:"source"`
         
    }
    
    // Media ...
    type Media struct {

        
            URL string  `json:"url"`
            Type string  `json:"type"`
            Alt string  `json:"alt"`
            Meta Meta  `json:"meta"`
         
    }
    
    // ProductListingActionPage ...
    type ProductListingActionPage struct {

        
            Query map[string]interface{}  `json:"query"`
            Type string  `json:"type"`
            Params map[string]interface{}  `json:"params"`
         
    }
    
    // ProductListingAction ...
    type ProductListingAction struct {

        
            Page ProductListingActionPage  `json:"page"`
            Type string  `json:"type"`
         
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

        
            L1 ProductBrand  `json:"l1"`
            L3 ProductBrand  `json:"l3"`
            L2 ProductBrand  `json:"l2"`
         
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
    
    // ApplicationItemMOQ ...
    type ApplicationItemMOQ struct {

        
            Maximum float64  `json:"maximum"`
            IncrementUnit float64  `json:"increment_unit"`
            Minimum float64  `json:"minimum"`
         
    }
    
    // Price ...
    type Price struct {

        
            CurrencyCode string  `json:"currency_code"`
            Min float64  `json:"min"`
            Max float64  `json:"max"`
            CurrencySymbol string  `json:"currency_symbol"`
         
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
    
    // ApplicationItemSEO ...
    type ApplicationItemSEO struct {

        
            Description interface{}  `json:"description"`
            Title interface{}  `json:"title"`
         
    }
    
    // NetQuantity ...
    type NetQuantity struct {

        
            Unit interface{}  `json:"unit"`
            Value float64  `json:"value"`
         
    }
    
    // ProductDetail ...
    type ProductDetail struct {

        
            Description string  `json:"description"`
            Highlights []string  `json:"highlights"`
            Medias []Media  `json:"medias"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Color string  `json:"color"`
            CategoryMap ProductCategoryMap  `json:"category_map"`
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
            Tryouts []string  `json:"tryouts"`
            Moq ApplicationItemMOQ  `json:"moq"`
            Tags []string  `json:"tags"`
            Price ProductListingPrice  `json:"price"`
            TeaserTag string  `json:"teaser_tag"`
            CustomMeta []CustomMetaFields  `json:"_custom_meta"`
            UID float64  `json:"uid"`
            ShortDescription string  `json:"short_description"`
            Rating float64  `json:"rating"`
            ItemCode string  `json:"item_code"`
            Slug string  `json:"slug"`
            ProductOnlineDate string  `json:"product_online_date"`
            Similars []string  `json:"similars"`
            Action ProductListingAction  `json:"action"`
            Attributes map[string]interface{}  `json:"attributes"`
            HasVariant bool  `json:"has_variant"`
            Brand ProductBrand  `json:"brand"`
            ItemType string  `json:"item_type"`
            Discount string  `json:"discount"`
            Type string  `json:"type"`
            Seo ApplicationItemSEO  `json:"seo"`
            Name string  `json:"name"`
            Categories []ProductBrand  `json:"categories"`
            NetQuantity NetQuantity  `json:"net_quantity"`
            ImageNature string  `json:"image_nature"`
            IsDependent bool  `json:"is_dependent"`
            RatingCount float64  `json:"rating_count"`
         
    }
    
    // ErrorResponse ...
    type ErrorResponse struct {

        
            Error string  `json:"error"`
         
    }
    
    // SizeChartValues ...
    type SizeChartValues struct {

        
            Col4 string  `json:"col_4"`
            Col2 string  `json:"col_2"`
            Col6 string  `json:"col_6"`
            Col5 string  `json:"col_5"`
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

        
            Col4 ColumnHeader  `json:"col_4"`
            Col2 ColumnHeader  `json:"col_2"`
            Col6 ColumnHeader  `json:"col_6"`
            Col5 ColumnHeader  `json:"col_5"`
            Col1 ColumnHeader  `json:"col_1"`
            Col3 ColumnHeader  `json:"col_3"`
         
    }
    
    // SizeChart ...
    type SizeChart struct {

        
            Description string  `json:"description"`
            SizeTip string  `json:"size_tip"`
            Sizes []SizeChartValues  `json:"sizes"`
            Image string  `json:"image"`
            Unit string  `json:"unit"`
            Headers ColumnHeaders  `json:"headers"`
            Title string  `json:"title"`
         
    }
    
    // ProductSizeStores ...
    type ProductSizeStores struct {

        
            Count float64  `json:"count"`
         
    }
    
    // Weight ...
    type Weight struct {

        
            Shipping float64  `json:"shipping"`
            Unit string  `json:"unit"`
            IsDefault bool  `json:"is_default"`
         
    }
    
    // Dimension ...
    type Dimension struct {

        
            Width float64  `json:"width"`
            Height float64  `json:"height"`
            Unit string  `json:"unit"`
            IsDefault bool  `json:"is_default"`
            Length float64  `json:"length"`
         
    }
    
    // ProductSize ...
    type ProductSize struct {

        
            Display string  `json:"display"`
            Weight Weight  `json:"weight"`
            Quantity float64  `json:"quantity"`
            Dimension Dimension  `json:"dimension"`
            IsAvailable bool  `json:"is_available"`
            SellerIdentifiers []string  `json:"seller_identifiers"`
            Value string  `json:"value"`
         
    }
    
    // ProductSizes ...
    type ProductSizes struct {

        
            SizeChart SizeChart  `json:"size_chart"`
            Stores ProductSizeStores  `json:"stores"`
            Sizes []ProductSize  `json:"sizes"`
            MultiSize bool  `json:"multi_size"`
            Discount string  `json:"discount"`
            Price ProductListingPrice  `json:"price"`
            Sellable bool  `json:"sellable"`
         
    }
    
    // AttributeDetail ...
    type AttributeDetail struct {

        
            Description string  `json:"description"`
            Logo string  `json:"logo"`
            Key string  `json:"key"`
            Display string  `json:"display"`
         
    }
    
    // AttributeMetadata ...
    type AttributeMetadata struct {

        
            Details []AttributeDetail  `json:"details"`
            Title string  `json:"title"`
         
    }
    
    // ProductsComparisonResponse ...
    type ProductsComparisonResponse struct {

        
            Items []ProductDetail  `json:"items"`
            AttributesMetadata []AttributeMetadata  `json:"attributes_metadata"`
         
    }
    
    // ProductCompareResponse ...
    type ProductCompareResponse struct {

        
            Subtitle string  `json:"subtitle"`
            Items []ProductDetail  `json:"items"`
            Title string  `json:"title"`
            AttributesMetadata []AttributeMetadata  `json:"attributes_metadata"`
         
    }
    
    // ProductFrequentlyComparedSimilarResponse ...
    type ProductFrequentlyComparedSimilarResponse struct {

        
            Similars ProductCompareResponse  `json:"similars"`
         
    }
    
    // ProductVariantItemResponse ...
    type ProductVariantItemResponse struct {

        
            Medias []Media  `json:"medias"`
            Color string  `json:"color"`
            Action ProductListingAction  `json:"action"`
            ColorName string  `json:"color_name"`
            Name string  `json:"name"`
            IsAvailable bool  `json:"is_available"`
            Slug string  `json:"slug"`
            CustomMeta []CustomMetaFields  `json:"_custom_meta"`
            UID float64  `json:"uid"`
            Value string  `json:"value"`
         
    }
    
    // ProductVariantResponse ...
    type ProductVariantResponse struct {

        
            Items []ProductVariantItemResponse  `json:"items"`
            Header string  `json:"header"`
            DisplayType string  `json:"display_type"`
            Key string  `json:"key"`
         
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

        
            Effective float64  `json:"effective"`
            Marked float64  `json:"marked"`
            Currency string  `json:"currency"`
         
    }
    
    // ProductStockStatusItem ...
    type ProductStockStatusItem struct {

        
            Store StoreDetail  `json:"store"`
            Identifier map[string]interface{}  `json:"identifier"`
            Size string  `json:"size"`
            ItemID float64  `json:"item_id"`
            Quantity float64  `json:"quantity"`
            Seller Seller  `json:"seller"`
            Company CompanyDetail  `json:"company"`
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

        
            Header string  `json:"header"`
            Total float64  `json:"total"`
            Key string  `json:"key"`
            Items []ProductVariantItemResponse  `json:"items"`
            DisplayType string  `json:"display_type"`
         
    }
    
    // ProductListingDetail ...
    type ProductListingDetail struct {

        
            Description string  `json:"description"`
            Highlights []string  `json:"highlights"`
            Medias []Media  `json:"medias"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Color string  `json:"color"`
            CategoryMap ProductCategoryMap  `json:"category_map"`
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
            Tryouts []string  `json:"tryouts"`
            Moq ApplicationItemMOQ  `json:"moq"`
            Tags []string  `json:"tags"`
            Price ProductListingPrice  `json:"price"`
            TeaserTag string  `json:"teaser_tag"`
            CustomMeta []CustomMetaFields  `json:"_custom_meta"`
            UID float64  `json:"uid"`
            ShortDescription string  `json:"short_description"`
            Sellable bool  `json:"sellable"`
            Rating float64  `json:"rating"`
            ItemCode string  `json:"item_code"`
            Sizes []string  `json:"sizes"`
            Slug string  `json:"slug"`
            ProductOnlineDate string  `json:"product_online_date"`
            Similars []string  `json:"similars"`
            Action ProductListingAction  `json:"action"`
            Attributes map[string]interface{}  `json:"attributes"`
            HasVariant bool  `json:"has_variant"`
            Variants []ProductVariantListingResponse  `json:"variants"`
            Brand ProductBrand  `json:"brand"`
            ItemType string  `json:"item_type"`
            Discount string  `json:"discount"`
            Type string  `json:"type"`
            Identifiers []string  `json:"identifiers"`
            Seo ApplicationItemSEO  `json:"seo"`
            Name string  `json:"name"`
            Categories []ProductBrand  `json:"categories"`
            NetQuantity NetQuantity  `json:"net_quantity"`
            ImageNature string  `json:"image_nature"`
            IsDependent bool  `json:"is_dependent"`
            RatingCount float64  `json:"rating_count"`
         
    }
    
    // ProductFiltersValue ...
    type ProductFiltersValue struct {

        
            Max float64  `json:"max"`
            Display string  `json:"display"`
            CurrencySymbol string  `json:"currency_symbol"`
            DisplayFormat string  `json:"display_format"`
            IsSelected bool  `json:"is_selected"`
            CurrencyCode string  `json:"currency_code"`
            SelectedMax float64  `json:"selected_max"`
            Min float64  `json:"min"`
            Count float64  `json:"count"`
            QueryFormat string  `json:"query_format"`
            SelectedMin float64  `json:"selected_min"`
            Value string  `json:"value"`
         
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

        
            IsSelected bool  `json:"is_selected"`
            Value string  `json:"value"`
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

        
            Description string  `json:"description"`
            Banners ImageUrls  `json:"banners"`
            Name string  `json:"name"`
            Logo Media  `json:"logo"`
            Departments []string  `json:"departments"`
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
            Discount string  `json:"discount"`
            Action ProductListingAction  `json:"action"`
         
    }
    
    // BrandListingResponse ...
    type BrandListingResponse struct {

        
            Page Page  `json:"page"`
            Items []BrandItem  `json:"items"`
         
    }
    
    // BrandDetailResponse ...
    type BrandDetailResponse struct {

        
            Description string  `json:"description"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Banners ImageUrls  `json:"banners"`
            Name string  `json:"name"`
            Logo Media  `json:"logo"`
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
            Banners ImageUrls  `json:"banners"`
            Childs []map[string]interface{}  `json:"childs"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
         
    }
    
    // SecondLevelChild ...
    type SecondLevelChild struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Action ProductListingAction  `json:"action"`
            Banners ImageUrls  `json:"banners"`
            Childs []ThirdLevelChild  `json:"childs"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
         
    }
    
    // Child ...
    type Child struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Action ProductListingAction  `json:"action"`
            Banners ImageUrls  `json:"banners"`
            Childs []SecondLevelChild  `json:"childs"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
         
    }
    
    // CategoryItems ...
    type CategoryItems struct {

        
            Banners CategoryBanner  `json:"banners"`
            Childs []Child  `json:"childs"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
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
            Banners ImageUrls  `json:"banners"`
            Name string  `json:"name"`
            Logo Media  `json:"logo"`
            UID float64  `json:"uid"`
         
    }
    
    // HomeListingResponse ...
    type HomeListingResponse struct {

        
            Page Page  `json:"page"`
            Items []ProductListingDetail  `json:"items"`
            Message string  `json:"message"`
         
    }
    
    // Department ...
    type Department struct {

        
            Name string  `json:"name"`
            Logo Media  `json:"logo"`
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
            Display string  `json:"display"`
            Logo Media  `json:"logo"`
            Type string  `json:"type"`
            Action ProductListingAction  `json:"action"`
         
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

        
            Description string  `json:"description"`
            Query []CollectionQuery  `json:"query"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Badge map[string]interface{}  `json:"badge"`
            Logo Media  `json:"logo"`
            UID string  `json:"uid"`
            AllowSort bool  `json:"allow_sort"`
            AllowFacets bool  `json:"allow_facets"`
            Slug string  `json:"slug"`
            Cron map[string]interface{}  `json:"cron"`
            SortOn string  `json:"sort_on"`
            Schedule map[string]interface{}  `json:"_schedule"`
            Action ProductListingAction  `json:"action"`
            Priority float64  `json:"priority"`
            Banners ImageUrls  `json:"banners"`
            IsActive bool  `json:"is_active"`
            Type string  `json:"type"`
            Meta map[string]interface{}  `json:"meta"`
            AppID string  `json:"app_id"`
            Name string  `json:"name"`
            Tag []string  `json:"tag"`
         
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
    
    // GetCollectionListingResponse ...
    type GetCollectionListingResponse struct {

        
            Page Page  `json:"page"`
            Items []GetCollectionDetailNest  `json:"items"`
            Filters CollectionListingFilter  `json:"filters"`
         
    }
    
    // CollectionDetailResponse ...
    type CollectionDetailResponse struct {

        
            Description string  `json:"description"`
            Query []CollectionQuery  `json:"query"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Badge map[string]interface{}  `json:"badge"`
            Logo Media  `json:"logo"`
            AllowSort bool  `json:"allow_sort"`
            AllowFacets bool  `json:"allow_facets"`
            Slug string  `json:"slug"`
            Cron map[string]interface{}  `json:"cron"`
            SortOn string  `json:"sort_on"`
            Schedule map[string]interface{}  `json:"_schedule"`
            Priority float64  `json:"priority"`
            Banners ImageUrls  `json:"banners"`
            IsActive bool  `json:"is_active"`
            Type string  `json:"type"`
            Meta map[string]interface{}  `json:"meta"`
            AppID string  `json:"app_id"`
            Name string  `json:"name"`
            Tag []string  `json:"tag"`
         
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

        
            Address string  `json:"address"`
            State string  `json:"state"`
            StoreEmail string  `json:"store_email"`
            Name string  `json:"name"`
            StoreCode string  `json:"store_code"`
            Country string  `json:"country"`
            City string  `json:"city"`
            Pincode float64  `json:"pincode"`
            LatLong LatLong  `json:"lat_long"`
            UID float64  `json:"uid"`
         
    }
    
    // StoreListingResponse ...
    type StoreListingResponse struct {

        
            Page Page  `json:"page"`
            Items []Store  `json:"items"`
         
    }
    
    // StoreAddressSerializer ...
    type StoreAddressSerializer struct {

        
            Address2 string  `json:"address2"`
            Landmark string  `json:"landmark"`
            State string  `json:"state"`
            Latitude float64  `json:"latitude"`
            Country string  `json:"country"`
            City string  `json:"city"`
            Address1 string  `json:"address1"`
            Pincode float64  `json:"pincode"`
            Longitude float64  `json:"longitude"`
         
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
    
    // StoreDepartments ...
    type StoreDepartments struct {

        
            Name string  `json:"name"`
            Logo map[string]interface{}  `json:"logo"`
            Slug string  `json:"slug"`
            PriorityOrder float64  `json:"priority_order"`
            UID float64  `json:"uid"`
         
    }
    
    // CompanyStore ...
    type CompanyStore struct {

        
            BusinessType string  `json:"business_type"`
            UID float64  `json:"uid"`
            CompanyType string  `json:"company_type"`
            Name string  `json:"name"`
         
    }
    
    // AppStore ...
    type AppStore struct {

        
            Address StoreAddressSerializer  `json:"address"`
            Manager StoreManagerSerializer  `json:"manager"`
            Name string  `json:"name"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            Departments []StoreDepartments  `json:"departments"`
            Company CompanyStore  `json:"company"`
            UID float64  `json:"uid"`
         
    }
    
    // ApplicationStoreListing ...
    type ApplicationStoreListing struct {

        
            Page Page  `json:"page"`
            Items []AppStore  `json:"items"`
            Filters []map[string]interface{}  `json:"filters"`
         
    }
    
    // Time ...
    type Time struct {

        
            Hour float64  `json:"hour"`
            Minute float64  `json:"minute"`
         
    }
    
    // StoreTiming ...
    type StoreTiming struct {

        
            Closing Time  `json:"closing"`
            Opening Time  `json:"opening"`
            Open bool  `json:"open"`
            Weekday string  `json:"weekday"`
         
    }
    
    // StoreDetails ...
    type StoreDetails struct {

        
            Address StoreAddressSerializer  `json:"address"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Timing []StoreTiming  `json:"timing"`
            Manager StoreManagerSerializer  `json:"manager"`
            Name string  `json:"name"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            Departments []StoreDepartments  `json:"departments"`
            Company CompanyStore  `json:"company"`
            UID float64  `json:"uid"`
         
    }
    
    // ProductGroupPrice ...
    type ProductGroupPrice struct {

        
            MaxMarked float64  `json:"max_marked"`
            MinEffective float64  `json:"min_effective"`
            MaxEffective float64  `json:"max_effective"`
            MinMarked float64  `json:"min_marked"`
            Currency interface{}  `json:"currency"`
         
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

        
            Description interface{}  `json:"description"`
            Images []interface{}  `json:"images"`
            Highlights []interface{}  `json:"highlights"`
            GroupedAttributes map[string]interface{}  `json:"grouped_attributes"`
            IsSet bool  `json:"is_set"`
            ShortDescription interface{}  `json:"short_description"`
            BrandUID float64  `json:"brand_uid"`
            Rating float64  `json:"rating"`
            ItemCode interface{}  `json:"item_code"`
            ImageNature interface{}  `json:"image_nature"`
            Slug interface{}  `json:"slug"`
            TemplateTag interface{}  `json:"template_tag"`
            OutOfStock bool  `json:"out_of_stock"`
            Attributes map[string]interface{}  `json:"attributes"`
            HasVariant bool  `json:"has_variant"`
            Media []map[string]interface{}  `json:"media"`
            CountryOfOrigin interface{}  `json:"country_of_origin"`
            Identifier map[string]interface{}  `json:"identifier"`
            Name interface{}  `json:"name"`
            HsnCode float64  `json:"hsn_code"`
            RatingCount float64  `json:"rating_count"`
         
    }
    
    // ProductInGroup ...
    type ProductInGroup struct {

        
            AutoAddToCart bool  `json:"auto_add_to_cart"`
            Price ProductGroupPrice  `json:"price"`
            ProductUID float64  `json:"product_uid"`
            Sizes []Size  `json:"sizes"`
            AutoSelect bool  `json:"auto_select"`
            MinQuantity float64  `json:"min_quantity"`
            MaxQuantity float64  `json:"max_quantity"`
            ProductDetails ProductDetails  `json:"product_details"`
            AllowRemove bool  `json:"allow_remove"`
         
    }
    
    // UserDetail ...
    type UserDetail struct {

        
            Contact string  `json:"contact"`
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
            SuperUser bool  `json:"super_user"`
         
    }
    
    // ProductGroupingModel ...
    type ProductGroupingModel struct {

        
            Products []ProductInGroup  `json:"products"`
            ID interface{}  `json:"_id"`
            ModifiedBy UserDetail  `json:"modified_by"`
            Choice interface{}  `json:"choice"`
            PageVisibility []interface{}  `json:"page_visibility"`
            Meta map[string]interface{}  `json:"meta"`
            IsActive bool  `json:"is_active"`
            VerifiedOn string  `json:"verified_on"`
            Name interface{}  `json:"name"`
            Logo interface{}  `json:"logo"`
            CreatedOn string  `json:"created_on"`
            Slug interface{}  `json:"slug"`
            ModifiedOn string  `json:"modified_on"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            VerifiedBy UserDetail  `json:"verified_by"`
            CompanyID float64  `json:"company_id"`
            CreatedBy UserDetail  `json:"created_by"`
         
    }
    
    // ProductBundle ...
    type ProductBundle struct {

        
            Items []ProductGroupingModel  `json:"items"`
         
    }
    
    // DetailsSchemaV3 ...
    type DetailsSchemaV3 struct {

        
            Key string  `json:"key"`
            Value string  `json:"value"`
            Type string  `json:"type"`
         
    }
    
    // SellerGroupAttributes ...
    type SellerGroupAttributes struct {

        
            Details []DetailsSchemaV3  `json:"details"`
            Title string  `json:"title"`
         
    }
    
    // ProductStockUnitPriceV3 ...
    type ProductStockUnitPriceV3 struct {

        
            CurrencyCode string  `json:"currency_code"`
            Price float64  `json:"price"`
            Unit string  `json:"unit"`
            CurrencySymbol string  `json:"currency_symbol"`
         
    }
    
    // ProductStockPriceV3 ...
    type ProductStockPriceV3 struct {

        
            Effective float64  `json:"effective"`
            Marked float64  `json:"marked"`
            Currency string  `json:"currency"`
         
    }
    
    // ArticleAssignmentV3 ...
    type ArticleAssignmentV3 struct {

        
            Level string  `json:"level"`
            Strategy string  `json:"strategy"`
         
    }
    
    // SellerV3 ...
    type SellerV3 struct {

        
            UID float64  `json:"uid"`
            Count float64  `json:"count"`
            Name string  `json:"name"`
         
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
    
    // StoreV3 ...
    type StoreV3 struct {

        
            UID float64  `json:"uid"`
            Count float64  `json:"count"`
            Name string  `json:"name"`
         
    }
    
    // MarketPlaceSttributesSchemaV3 ...
    type MarketPlaceSttributesSchemaV3 struct {

        
            Details []DetailsSchemaV3  `json:"details"`
            Title string  `json:"title"`
         
    }
    
    // StrategyWiseListingSchemaV3 ...
    type StrategyWiseListingSchemaV3 struct {

        
            Distance float64  `json:"distance"`
            Quantity float64  `json:"quantity"`
            Tat float64  `json:"tat"`
            Pincode float64  `json:"pincode"`
         
    }
    
    // ReturnConfigSchemaV3 ...
    type ReturnConfigSchemaV3 struct {

        
            Returnable bool  `json:"returnable"`
            Unit string  `json:"unit"`
            Time float64  `json:"time"`
         
    }
    
    // ProductSizePriceResponseV3 ...
    type ProductSizePriceResponseV3 struct {

        
            GroupedAttributes []SellerGroupAttributes  `json:"grouped_attributes"`
            PricePerUnit ProductStockUnitPriceV3  `json:"price_per_unit"`
            Price ProductStockPriceV3  `json:"price"`
            Pincode float64  `json:"pincode"`
            ArticleAssignment ArticleAssignmentV3  `json:"article_assignment"`
            IsCod bool  `json:"is_cod"`
            Seller SellerV3  `json:"seller"`
            Set ProductSetV3  `json:"set"`
            PricePerPiece ProductStockPriceV3  `json:"price_per_piece"`
            Store StoreV3  `json:"store"`
            ArticleID string  `json:"article_id"`
            ItemType string  `json:"item_type"`
            SellerCount float64  `json:"seller_count"`
            MarketplaceAttributes []MarketPlaceSttributesSchemaV3  `json:"marketplace_attributes"`
            StrategyWiseListing []StrategyWiseListingSchemaV3  `json:"strategy_wise_listing"`
            IsGift bool  `json:"is_gift"`
            LongLat []float64  `json:"long_lat"`
            SpecialBadge string  `json:"special_badge"`
            Quantity float64  `json:"quantity"`
            Discount string  `json:"discount"`
            ReturnConfig ReturnConfigSchemaV3  `json:"return_config"`
         
    }
    
    // ProductSizeSellerFilterSchemaV3 ...
    type ProductSizeSellerFilterSchemaV3 struct {

        
            IsSelected bool  `json:"is_selected"`
            Value string  `json:"value"`
            Name string  `json:"name"`
         
    }
    
    // ProductSizeSellersResponseV3 ...
    type ProductSizeSellersResponseV3 struct {

        
            Page Page  `json:"page"`
            Items []ProductSizePriceResponseV3  `json:"items"`
            SortOn []ProductSizeSellerFilterSchemaV3  `json:"sort_on"`
         
    }
    

    
    // Ownership ...
    type Ownership struct {

        
            PayableBy string  `json:"payable_by"`
            PayableCategory string  `json:"payable_category"`
         
    }
    
    // AppliedPromotion ...
    type AppliedPromotion struct {

        
            MrpPromotion bool  `json:"mrp_promotion"`
            PromotionType string  `json:"promotion_type"`
            ArticleQuantity float64  `json:"article_quantity"`
            Ownership Ownership  `json:"ownership"`
            OfferText string  `json:"offer_text"`
            Amount float64  `json:"amount"`
            PromoID string  `json:"promo_id"`
         
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

        
            PaymentIdentifier string  `json:"payment_identifier"`
            Enabled bool  `json:"enabled"`
            DefaultOptions string  `json:"default_options"`
         
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

        
            Message []string  `json:"message"`
            CurrencySymbol string  `json:"currency_symbol"`
            Display string  `json:"display"`
            Key string  `json:"key"`
            CurrencyCode string  `json:"currency_code"`
            Value float64  `json:"value"`
         
    }
    
    // CouponBreakup ...
    type CouponBreakup struct {

        
            Code string  `json:"code"`
            Message string  `json:"message"`
            IsApplied bool  `json:"is_applied"`
            UID string  `json:"uid"`
            Type string  `json:"type"`
            Value float64  `json:"value"`
         
    }
    
    // RawBreakup ...
    type RawBreakup struct {

        
            YouSaved float64  `json:"you_saved"`
            Discount float64  `json:"discount"`
            Coupon float64  `json:"coupon"`
            ConvenienceFee float64  `json:"convenience_fee"`
            GstCharges float64  `json:"gst_charges"`
            MrpTotal float64  `json:"mrp_total"`
            Vog float64  `json:"vog"`
            Subtotal float64  `json:"subtotal"`
            DeliveryCharge float64  `json:"delivery_charge"`
            FyndCash float64  `json:"fynd_cash"`
            CodCharge float64  `json:"cod_charge"`
            Total float64  `json:"total"`
         
    }
    
    // CartBreakup ...
    type CartBreakup struct {

        
            LoyaltyPoints LoyaltyPoints  `json:"loyalty_points"`
            Display []DisplayBreakup  `json:"display"`
            Coupon CouponBreakup  `json:"coupon"`
            Raw RawBreakup  `json:"raw"`
         
    }
    
    // ProductPrice ...
    type ProductPrice struct {

        
            CurrencySymbol string  `json:"currency_symbol"`
            Effective float64  `json:"effective"`
            AddOn float64  `json:"add_on"`
            Selling float64  `json:"selling"`
            CurrencyCode string  `json:"currency_code"`
            Marked float64  `json:"marked"`
         
    }
    
    // ProductPriceInfo ...
    type ProductPriceInfo struct {

        
            Base ProductPrice  `json:"base"`
            Converted ProductPrice  `json:"converted"`
         
    }
    
    // ProductAvailabilitySize ...
    type ProductAvailabilitySize struct {

        
            IsAvailable bool  `json:"is_available"`
            Display string  `json:"display"`
            Value string  `json:"value"`
         
    }
    
    // ProductAvailability ...
    type ProductAvailability struct {

        
            Deliverable bool  `json:"deliverable"`
            OutOfStock bool  `json:"out_of_stock"`
            IsValid bool  `json:"is_valid"`
            OtherStoreQuantity float64  `json:"other_store_quantity"`
            Sizes []string  `json:"sizes"`
            AvailableSizes []ProductAvailabilitySize  `json:"available_sizes"`
         
    }
    
    // PromoMeta ...
    type PromoMeta struct {

        
            Message string  `json:"message"`
         
    }
    
    // BasePrice ...
    type BasePrice struct {

        
            CurrencyCode string  `json:"currency_code"`
            Marked float64  `json:"marked"`
            CurrencySymbol string  `json:"currency_symbol"`
            Effective float64  `json:"effective"`
         
    }
    
    // ArticlePriceInfo ...
    type ArticlePriceInfo struct {

        
            Base BasePrice  `json:"base"`
            Converted BasePrice  `json:"converted"`
         
    }
    
    // BaseInfo ...
    type BaseInfo struct {

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // ProductArticle ...
    type ProductArticle struct {

        
            Price ArticlePriceInfo  `json:"price"`
            Size string  `json:"size"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            UID string  `json:"uid"`
            Store BaseInfo  `json:"store"`
            Seller BaseInfo  `json:"seller"`
            Quantity float64  `json:"quantity"`
            Type string  `json:"type"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ProductGroupTags []string  `json:"product_group_tags"`
         
    }
    
    // CartProductIdentifer ...
    type CartProductIdentifer struct {

        
            Identifier string  `json:"identifier"`
         
    }
    
    // ProductImage ...
    type ProductImage struct {

        
            AspectRatio string  `json:"aspect_ratio"`
            SecureURL string  `json:"secure_url"`
            URL string  `json:"url"`
         
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

        
            Type string  `json:"type"`
            Query ActionQuery  `json:"query"`
            URL string  `json:"url"`
         
    }
    
    // CartProduct ...
    type CartProduct struct {

        
            Brand BaseInfo  `json:"brand"`
            UID float64  `json:"uid"`
            Images []ProductImage  `json:"images"`
            Slug string  `json:"slug"`
            Type string  `json:"type"`
            Categories []CategoryInfo  `json:"categories"`
            Action ProductAction  `json:"action"`
            Name string  `json:"name"`
         
    }
    
    // CartProductInfo ...
    type CartProductInfo struct {

        
            Price ProductPriceInfo  `json:"price"`
            Discount string  `json:"discount"`
            CouponMessage string  `json:"coupon_message"`
            Message string  `json:"message"`
            PromotionsApplied []AppliedPromotion  `json:"promotions_applied"`
            Availability ProductAvailability  `json:"availability"`
            Quantity float64  `json:"quantity"`
            BulkOffer map[string]interface{}  `json:"bulk_offer"`
            IsSet bool  `json:"is_set"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            PromoMeta PromoMeta  `json:"promo_meta"`
            Article ProductArticle  `json:"article"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            Moq map[string]interface{}  `json:"moq"`
            Key string  `json:"key"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
            PricePerUnit ProductPriceInfo  `json:"price_per_unit"`
            Product CartProduct  `json:"product"`
         
    }
    
    // CartCurrency ...
    type CartCurrency struct {

        
            Code string  `json:"code"`
            Symbol string  `json:"symbol"`
         
    }
    
    // CartDetailResponse ...
    type CartDetailResponse struct {

        
            CheckoutMode string  `json:"checkout_mode"`
            IsValid bool  `json:"is_valid"`
            Message string  `json:"message"`
            AppliedPromoDetails []AppliedPromotion  `json:"applied_promo_details"`
            ID string  `json:"id"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            Gstin string  `json:"gstin"`
            BuyNow bool  `json:"buy_now"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            Comment string  `json:"comment"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            Items []CartProductInfo  `json:"items"`
            Currency CartCurrency  `json:"currency"`
            LastModified string  `json:"last_modified"`
            CouponText string  `json:"coupon_text"`
         
    }
    
    // AddProductCart ...
    type AddProductCart struct {

        
            ItemID float64  `json:"item_id"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            ArticleAssignment map[string]interface{}  `json:"article_assignment"`
            Quantity float64  `json:"quantity"`
            ArticleID string  `json:"article_id"`
            ItemSize string  `json:"item_size"`
            Display string  `json:"display"`
            ParentItemIdentifiers []map[string]string  `json:"parent_item_identifiers"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Pos bool  `json:"pos"`
            ProductGroupTags []string  `json:"product_group_tags"`
            StoreID float64  `json:"store_id"`
            SellerID float64  `json:"seller_id"`
         
    }
    
    // AddCartRequest ...
    type AddCartRequest struct {

        
            Items []AddProductCart  `json:"items"`
         
    }
    
    // AddCartDetailResponse ...
    type AddCartDetailResponse struct {

        
            Success bool  `json:"success"`
            Partial bool  `json:"partial"`
            Cart CartDetailResponse  `json:"cart"`
            Message string  `json:"message"`
         
    }
    
    // UpdateProductCart ...
    type UpdateProductCart struct {

        
            ItemID float64  `json:"item_id"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            Quantity float64  `json:"quantity"`
            ArticleID string  `json:"article_id"`
            ItemSize string  `json:"item_size"`
            ItemIndex float64  `json:"item_index"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
         
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

        
            MinimumCartValue float64  `json:"minimum_cart_value"`
            SubTitle string  `json:"sub_title"`
            Message string  `json:"message"`
            IsApplied bool  `json:"is_applied"`
            CouponValue float64  `json:"coupon_value"`
            CouponCode string  `json:"coupon_code"`
            Title string  `json:"title"`
            IsApplicable bool  `json:"is_applicable"`
            ExpiresOn string  `json:"expires_on"`
            MaxDiscountValue float64  `json:"max_discount_value"`
         
    }
    
    // PageCoupon ...
    type PageCoupon struct {

        
            TotalItemCount float64  `json:"total_item_count"`
            Current float64  `json:"current"`
            HasPrevious bool  `json:"has_previous"`
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

        
            BulkEffective float64  `json:"bulk_effective"`
            CurrencySymbol string  `json:"currency_symbol"`
            Effective float64  `json:"effective"`
            CurrencyCode string  `json:"currency_code"`
            Marked float64  `json:"marked"`
         
    }
    
    // OfferItem ...
    type OfferItem struct {

        
            Price OfferPrice  `json:"price"`
            AutoApplied bool  `json:"auto_applied"`
            Best bool  `json:"best"`
            Quantity float64  `json:"quantity"`
            Type string  `json:"type"`
            Margin float64  `json:"margin"`
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

        
            Longitude float64  `json:"longitude"`
            Latitude float64  `json:"latitude"`
         
    }
    
    // Address ...
    type Address struct {

        
            Meta map[string]interface{}  `json:"meta"`
            GoogleMapPoint map[string]interface{}  `json:"google_map_point"`
            ID string  `json:"id"`
            GeoLocation GeoLocation  `json:"geo_location"`
            City string  `json:"city"`
            Phone string  `json:"phone"`
            Landmark string  `json:"landmark"`
            AddressType string  `json:"address_type"`
            CountryIsoCode string  `json:"country_iso_code"`
            UserID string  `json:"user_id"`
            Area string  `json:"area"`
            IsActive bool  `json:"is_active"`
            CheckoutMode string  `json:"checkout_mode"`
            Email string  `json:"email"`
            State string  `json:"state"`
            AreaCode string  `json:"area_code"`
            Country string  `json:"country"`
            AreaCodeSlug string  `json:"area_code_slug"`
            IsDefaultAddress bool  `json:"is_default_address"`
            Tags []string  `json:"tags"`
            CountryPhoneCode string  `json:"country_phone_code"`
            Address string  `json:"address"`
            CountryCode string  `json:"country_code"`
            Name string  `json:"name"`
         
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

        
            BillingAddressID string  `json:"billing_address_id"`
            ID string  `json:"id"`
            CartID string  `json:"cart_id"`
         
    }
    
    // UpdateCartPaymentRequest ...
    type UpdateCartPaymentRequest struct {

        
            AddressID string  `json:"address_id"`
            PaymentIdentifier string  `json:"payment_identifier"`
            MerchantCode string  `json:"merchant_code"`
            ID string  `json:"id"`
            PaymentMode string  `json:"payment_mode"`
            AggregatorName string  `json:"aggregator_name"`
         
    }
    
    // CouponValidity ...
    type CouponValidity struct {

        
            Code string  `json:"code"`
            Discount float64  `json:"discount"`
            DisplayMessageEn string  `json:"display_message_en"`
            Valid bool  `json:"valid"`
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

        
            OrderType string  `json:"order_type"`
            DpID string  `json:"dp_id"`
            Promise ShipmentPromise  `json:"promise"`
            Shipments float64  `json:"shipments"`
            DpOptions map[string]interface{}  `json:"dp_options"`
            ShipmentType string  `json:"shipment_type"`
            BoxType string  `json:"box_type"`
            Items []CartProductInfo  `json:"items"`
            FulfillmentType string  `json:"fulfillment_type"`
            FulfillmentID float64  `json:"fulfillment_id"`
         
    }
    
    // CartShipmentsResponse ...
    type CartShipmentsResponse struct {

        
            Error bool  `json:"error"`
            CheckoutMode string  `json:"checkout_mode"`
            Message string  `json:"message"`
            ID string  `json:"id"`
            Shipments []ShipmentResponse  `json:"shipments"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            Comment string  `json:"comment"`
            Currency CartCurrency  `json:"currency"`
            LastModified string  `json:"last_modified"`
            CouponText string  `json:"coupon_text"`
            IsValid bool  `json:"is_valid"`
            CartID float64  `json:"cart_id"`
            UID string  `json:"uid"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            Gstin string  `json:"gstin"`
            BuyNow bool  `json:"buy_now"`
            RestrictCheckout bool  `json:"restrict_checkout"`
         
    }
    
    // StaffCheckout ...
    type StaffCheckout struct {

        
            EmployeeCode string  `json:"employee_code"`
            FirstName string  `json:"first_name"`
            LastName string  `json:"last_name"`
            ID string  `json:"_id"`
            User string  `json:"user"`
         
    }
    
    // CartCheckoutCustomMeta ...
    type CartCheckoutCustomMeta struct {

        
            Value string  `json:"value"`
            Key string  `json:"key"`
         
    }
    
    // CartCheckoutDetailRequest ...
    type CartCheckoutDetailRequest struct {

        
            OrderType string  `json:"order_type"`
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
            AddressID string  `json:"address_id"`
            CallbackURL string  `json:"callback_url"`
            Meta map[string]interface{}  `json:"meta"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            BillingAddressID string  `json:"billing_address_id"`
            PaymentIdentifier string  `json:"payment_identifier"`
            MerchantCode string  `json:"merchant_code"`
            Staff StaffCheckout  `json:"staff"`
            PaymentMode string  `json:"payment_mode"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
            PaymentAutoConfirm bool  `json:"payment_auto_confirm"`
            Aggregator string  `json:"aggregator"`
            PaymentParams map[string]interface{}  `json:"payment_params"`
            OrderingStore float64  `json:"ordering_store"`
            CustomMeta []CartCheckoutCustomMeta  `json:"custom_meta"`
         
    }
    
    // CheckCart ...
    type CheckCart struct {

        
            OrderID string  `json:"order_id"`
            Message string  `json:"message"`
            ID string  `json:"id"`
            StoreEmps []map[string]interface{}  `json:"store_emps"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            Comment string  `json:"comment"`
            LastModified string  `json:"last_modified"`
            Success bool  `json:"success"`
            IsValid bool  `json:"is_valid"`
            UID string  `json:"uid"`
            CodMessage string  `json:"cod_message"`
            BuyNow bool  `json:"buy_now"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            DeliveryCharges float64  `json:"delivery_charges"`
            ErrorMessage string  `json:"error_message"`
            UserType string  `json:"user_type"`
            CheckoutMode string  `json:"checkout_mode"`
            DeliveryChargeOrderValue float64  `json:"delivery_charge_order_value"`
            CodCharges float64  `json:"cod_charges"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            Items []CartProductInfo  `json:"items"`
            Currency CartCurrency  `json:"currency"`
            CouponText string  `json:"coupon_text"`
            CodAvailable bool  `json:"cod_available"`
            CartID float64  `json:"cart_id"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            Gstin string  `json:"gstin"`
            StoreCode string  `json:"store_code"`
         
    }
    
    // CartCheckoutResponse ...
    type CartCheckoutResponse struct {

        
            OrderID string  `json:"order_id"`
            Success bool  `json:"success"`
            CallbackURL string  `json:"callback_url"`
            Cart CheckCart  `json:"cart"`
            Message string  `json:"message"`
            AppInterceptURL string  `json:"app_intercept_url"`
            PaymentConfirmURL string  `json:"payment_confirm_url"`
            Data map[string]interface{}  `json:"data"`
         
    }
    
    // CartMetaRequest ...
    type CartMetaRequest struct {

        
            Gstin string  `json:"gstin"`
            PickUpCustomerDetails map[string]interface{}  `json:"pick_up_customer_details"`
            CheckoutMode string  `json:"checkout_mode"`
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
            CreatedOn string  `json:"created_on"`
            Token string  `json:"token"`
            User map[string]interface{}  `json:"user"`
         
    }
    
    // SharedCart ...
    type SharedCart struct {

        
            Message string  `json:"message"`
            ID string  `json:"id"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            Comment string  `json:"comment"`
            LastModified string  `json:"last_modified"`
            IsValid bool  `json:"is_valid"`
            UID string  `json:"uid"`
            BuyNow bool  `json:"buy_now"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            SharedCartDetails SharedCartDetails  `json:"shared_cart_details"`
            CheckoutMode string  `json:"checkout_mode"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            Items []CartProductInfo  `json:"items"`
            Currency CartCurrency  `json:"currency"`
            CouponText string  `json:"coupon_text"`
            CartID float64  `json:"cart_id"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            Gstin string  `json:"gstin"`
         
    }
    
    // SharedCartResponse ...
    type SharedCartResponse struct {

        
            Error string  `json:"error"`
            Cart SharedCart  `json:"cart"`
         
    }
    
    // PromotionOffer ...
    type PromotionOffer struct {

        
            ValidTill string  `json:"valid_till"`
            ID string  `json:"id"`
            PromotionGroup string  `json:"promotion_group"`
            OfferText string  `json:"offer_text"`
            Description string  `json:"description"`
         
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

        
            Code string  `json:"code"`
            Symbol string  `json:"symbol"`
         
    }
    
    // LadderPrice ...
    type LadderPrice struct {

        
            CurrencySymbol string  `json:"currency_symbol"`
            Effective float64  `json:"effective"`
            OfferPrice float64  `json:"offer_price"`
            CurrencyCode string  `json:"currency_code"`
            Marked float64  `json:"marked"`
         
    }
    
    // LadderOfferItem ...
    type LadderOfferItem struct {

        
            Price LadderPrice  `json:"price"`
            MinQuantity float64  `json:"min_quantity"`
            MaxQuantity float64  `json:"max_quantity"`
            Type string  `json:"type"`
            Margin float64  `json:"margin"`
         
    }
    
    // LadderPriceOffer ...
    type LadderPriceOffer struct {

        
            ValidTill string  `json:"valid_till"`
            ID string  `json:"id"`
            OfferPrices []LadderOfferItem  `json:"offer_prices"`
            PromotionGroup string  `json:"promotion_group"`
            OfferText string  `json:"offer_text"`
            Description string  `json:"description"`
         
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

        
            Pin string  `json:"pin"`
            Sdk bool  `json:"sdk"`
            MerchantKey string  `json:"merchant_key"`
            ConfigType string  `json:"config_type"`
            API string  `json:"api"`
            Key string  `json:"key"`
            Secret string  `json:"secret"`
            MerchantID string  `json:"merchant_id"`
            VerifyAPI string  `json:"verify_api"`
            UserID string  `json:"user_id"`
         
    }
    
    // AggregatorsConfigDetailResponse ...
    type AggregatorsConfigDetailResponse struct {

        
            Ccavenue AggregatorConfigDetail  `json:"ccavenue"`
            Mswipe AggregatorConfigDetail  `json:"mswipe"`
            Stripe AggregatorConfigDetail  `json:"stripe"`
            Payumoney AggregatorConfigDetail  `json:"payumoney"`
            Razorpay AggregatorConfigDetail  `json:"razorpay"`
            Success bool  `json:"success"`
            Simpl AggregatorConfigDetail  `json:"simpl"`
            Rupifi AggregatorConfigDetail  `json:"rupifi"`
            Env string  `json:"env"`
            Juspay AggregatorConfigDetail  `json:"juspay"`
         
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

        
            NameOnCard string  `json:"name_on_card"`
            CardID string  `json:"card_id"`
            Refresh bool  `json:"refresh"`
            Nickname string  `json:"nickname"`
         
    }
    
    // AttachCardsResponse ...
    type AttachCardsResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
            Data map[string]interface{}  `json:"data"`
         
    }
    
    // CardPaymentGateway ...
    type CardPaymentGateway struct {

        
            CustomerID string  `json:"customer_id"`
            Aggregator string  `json:"aggregator"`
            API string  `json:"api"`
         
    }
    
    // ActiveCardPaymentGatewayResponse ...
    type ActiveCardPaymentGatewayResponse struct {

        
            Message string  `json:"message"`
            Cards CardPaymentGateway  `json:"cards"`
            Success bool  `json:"success"`
         
    }
    
    // Card ...
    type Card struct {

        
            CardName string  `json:"card_name"`
            CardReference string  `json:"card_reference"`
            CardIsin string  `json:"card_isin"`
            CardNumber string  `json:"card_number"`
            CardBrand string  `json:"card_brand"`
            ExpYear float64  `json:"exp_year"`
            AggregatorName string  `json:"aggregator_name"`
            CardType string  `json:"card_type"`
            Expired bool  `json:"expired"`
            Nickname string  `json:"nickname"`
            CardBrandImage string  `json:"card_brand_image"`
            CardIssuer string  `json:"card_issuer"`
            CompliantWithTokenisationGuidelines bool  `json:"compliant_with_tokenisation_guidelines"`
            CardID string  `json:"card_id"`
            ExpMonth float64  `json:"exp_month"`
            CardFingerprint string  `json:"card_fingerprint"`
            CardToken string  `json:"card_token"`
         
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
            BillingAddress map[string]interface{}  `json:"billing_address"`
            TransactionAmountInPaise float64  `json:"transaction_amount_in_paise"`
            MerchantParams map[string]interface{}  `json:"merchant_params"`
            PhoneNumber string  `json:"phone_number"`
            Aggregator string  `json:"aggregator"`
            OrderItems []map[string]interface{}  `json:"order_items"`
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
            TransactionToken string  `json:"transaction_token"`
            Amount float64  `json:"amount"`
            Verified bool  `json:"verified"`
            Aggregator string  `json:"aggregator"`
         
    }
    
    // ChargeCustomerResponse ...
    type ChargeCustomerResponse struct {

        
            OrderID string  `json:"order_id"`
            CartID string  `json:"cart_id"`
            DeliveryAddressID string  `json:"delivery_address_id"`
            Success bool  `json:"success"`
            Aggregator string  `json:"aggregator"`
            Message string  `json:"message"`
            Status string  `json:"status"`
         
    }
    
    // PaymentInitializationRequest ...
    type PaymentInitializationRequest struct {

        
            Contact string  `json:"contact"`
            OrderID string  `json:"order_id"`
            Currency string  `json:"currency"`
            Vpa string  `json:"vpa"`
            Email string  `json:"email"`
            CustomerID string  `json:"customer_id"`
            Amount float64  `json:"amount"`
            RazorpayPaymentID string  `json:"razorpay_payment_id"`
            Aggregator string  `json:"aggregator"`
            MerchantOrderID string  `json:"merchant_order_id"`
            Timeout float64  `json:"timeout"`
            Method string  `json:"method"`
         
    }
    
    // PaymentInitializationResponse ...
    type PaymentInitializationResponse struct {

        
            UpiPollURL string  `json:"upi_poll_url"`
            CustomerID string  `json:"customer_id"`
            PollingURL string  `json:"polling_url"`
            RazorpayPaymentID string  `json:"razorpay_payment_id"`
            Currency string  `json:"currency"`
            Email string  `json:"email"`
            Success bool  `json:"success"`
            MerchantOrderID string  `json:"merchant_order_id"`
            Timeout float64  `json:"timeout"`
            Method string  `json:"method"`
            Vpa string  `json:"vpa"`
            Amount float64  `json:"amount"`
            AggregatorOrderID string  `json:"aggregator_order_id"`
            BqrImage string  `json:"bqr_image"`
            Contact string  `json:"contact"`
            OrderID string  `json:"order_id"`
            DeviceID string  `json:"device_id"`
            PaymentID string  `json:"payment_id"`
            VirtualID string  `json:"virtual_id"`
            Aggregator string  `json:"aggregator"`
            Status string  `json:"status"`
         
    }
    
    // PaymentStatusUpdateRequest ...
    type PaymentStatusUpdateRequest struct {

        
            OrderID string  `json:"order_id"`
            Contact string  `json:"contact"`
            Currency string  `json:"currency"`
            Vpa string  `json:"vpa"`
            Email string  `json:"email"`
            CustomerID string  `json:"customer_id"`
            PaymentID string  `json:"payment_id"`
            Amount float64  `json:"amount"`
            Success bool  `json:"success"`
            RazorpayPaymentID string  `json:"razorpay_payment_id"`
            Aggregator string  `json:"aggregator"`
            MerchantOrderID string  `json:"merchant_order_id"`
            Status string  `json:"status"`
            Method string  `json:"method"`
         
    }
    
    // PaymentStatusUpdateResponse ...
    type PaymentStatusUpdateResponse struct {

        
            Success bool  `json:"success"`
            AggregatorName string  `json:"aggregator_name"`
            RedirectURL string  `json:"redirect_url"`
            Retry bool  `json:"retry"`
            Status string  `json:"status"`
         
    }
    
    // AggregatorRoute ...
    type AggregatorRoute struct {

        
            PaymentFlowData string  `json:"payment_flow_data"`
            PaymentFlow string  `json:"payment_flow"`
            APILink string  `json:"api_link"`
            Data map[string]interface{}  `json:"data"`
         
    }
    
    // PaymentFlow ...
    type PaymentFlow struct {

        
            Ccavenue AggregatorRoute  `json:"ccavenue"`
            BqrRazorpay AggregatorRoute  `json:"bqr_razorpay"`
            Fynd AggregatorRoute  `json:"fynd"`
            Mswipe AggregatorRoute  `json:"mswipe"`
            Stripe AggregatorRoute  `json:"stripe"`
            Payubiz AggregatorRoute  `json:"payubiz"`
            Simpl AggregatorRoute  `json:"simpl"`
            Razorpay AggregatorRoute  `json:"razorpay"`
            Rupifi AggregatorRoute  `json:"rupifi"`
            Jiopay AggregatorRoute  `json:"jiopay"`
            UpiRazorpay AggregatorRoute  `json:"upi_razorpay"`
            Epaylater AggregatorRoute  `json:"epaylater"`
            Juspay AggregatorRoute  `json:"juspay"`
         
    }
    
    // PaymentModeLogo ...
    type PaymentModeLogo struct {

        
            Large string  `json:"large"`
            Small string  `json:"small"`
         
    }
    
    // IntentApp ...
    type IntentApp struct {

        
            Logos PaymentModeLogo  `json:"logos"`
            PackageName string  `json:"package_name"`
            Code string  `json:"code"`
            DisplayName string  `json:"display_name"`
         
    }
    
    // IntentAppErrorList ...
    type IntentAppErrorList struct {

        
            PackageName string  `json:"package_name"`
            Code string  `json:"code"`
         
    }
    
    // PaymentModeList ...
    type PaymentModeList struct {

        
            MerchantCode string  `json:"merchant_code"`
            Code string  `json:"code"`
            ExpMonth float64  `json:"exp_month"`
            CardFingerprint string  `json:"card_fingerprint"`
            CardToken string  `json:"card_token"`
            IntentApp []IntentApp  `json:"intent_app"`
            CardName string  `json:"card_name"`
            IntentAppErrorList []string  `json:"intent_app_error_list"`
            CardIsin string  `json:"card_isin"`
            IntentAppErrorDictList []IntentAppErrorList  `json:"intent_app_error_dict_list"`
            AggregatorName string  `json:"aggregator_name"`
            CardType string  `json:"card_type"`
            CodLimitPerOrder float64  `json:"cod_limit_per_order"`
            Timeout float64  `json:"timeout"`
            Name string  `json:"name"`
            LogoURL PaymentModeLogo  `json:"logo_url"`
            CardNumber string  `json:"card_number"`
            CodLimit float64  `json:"cod_limit"`
            IntentFlow bool  `json:"intent_flow"`
            DisplayPriority float64  `json:"display_priority"`
            RetryCount float64  `json:"retry_count"`
            DisplayName string  `json:"display_name"`
            Nickname string  `json:"nickname"`
            FyndVpa string  `json:"fynd_vpa"`
            CardReference string  `json:"card_reference"`
            CardBrand string  `json:"card_brand"`
            ExpYear float64  `json:"exp_year"`
            RemainingLimit float64  `json:"remaining_limit"`
            Expired bool  `json:"expired"`
            CompliantWithTokenisationGuidelines bool  `json:"compliant_with_tokenisation_guidelines"`
            CardBrandImage string  `json:"card_brand_image"`
            CardIssuer string  `json:"card_issuer"`
            CardID string  `json:"card_id"`
         
    }
    
    // RootPaymentMode ...
    type RootPaymentMode struct {

        
            SaveCard bool  `json:"save_card"`
            AddCardEnabled bool  `json:"add_card_enabled"`
            IsPayByCardPl bool  `json:"is_pay_by_card_pl"`
            DisplayPriority float64  `json:"display_priority"`
            AggregatorName string  `json:"aggregator_name"`
            AnonymousEnable bool  `json:"anonymous_enable"`
            DisplayName string  `json:"display_name"`
            Name string  `json:"name"`
            List []PaymentModeList  `json:"list"`
         
    }
    
    // PaymentOptionAndFlow ...
    type PaymentOptionAndFlow struct {

        
            PaymentFlows PaymentFlow  `json:"payment_flows"`
            PaymentOption []RootPaymentMode  `json:"payment_option"`
         
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

        
            CustomerName string  `json:"customer_name"`
            UpiVpa string  `json:"upi_vpa"`
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

        
            LogoLarge string  `json:"logo_large"`
            ID float64  `json:"id"`
            DisplayName string  `json:"display_name"`
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
            Comment string  `json:"comment"`
            Email string  `json:"email"`
            ID float64  `json:"id"`
            IsActive bool  `json:"is_active"`
            CreatedOn string  `json:"created_on"`
            BankName string  `json:"bank_name"`
            BranchName string  `json:"branch_name"`
            ModifiedOn string  `json:"modified_on"`
            AccountHolder string  `json:"account_holder"`
            AccountNo string  `json:"account_no"`
            DisplayName string  `json:"display_name"`
            DelightsUserName string  `json:"delights_user_name"`
            Subtitle string  `json:"subtitle"`
            Mobile string  `json:"mobile"`
            TransferMode string  `json:"transfer_mode"`
            BeneficiaryID string  `json:"beneficiary_id"`
            Address string  `json:"address"`
            IfscCode string  `json:"ifsc_code"`
         
    }
    
    // OrderBeneficiaryResponse ...
    type OrderBeneficiaryResponse struct {

        
            Beneficiaries []OrderBeneficiaryDetails  `json:"beneficiaries"`
            ShowBeneficiaryDetails bool  `json:"show_beneficiary_details"`
         
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

        
            RequestID string  `json:"request_id"`
            Otp string  `json:"otp"`
            HashKey string  `json:"hash_key"`
         
    }
    
    // AddBeneficiaryViaOtpVerificationResponse ...
    type AddBeneficiaryViaOtpVerificationResponse struct {

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // WrongOtpError ...
    type WrongOtpError struct {

        
            Description string  `json:"description"`
            Success string  `json:"success"`
            IsVerifiedFlag bool  `json:"is_verified_flag"`
         
    }
    
    // BeneficiaryModeDetails ...
    type BeneficiaryModeDetails struct {

        
            Comment string  `json:"comment"`
            Vpa string  `json:"vpa"`
            BankName string  `json:"bank_name"`
            BranchName string  `json:"branch_name"`
            Email string  `json:"email"`
            AccountHolder string  `json:"account_holder"`
            Mobile string  `json:"mobile"`
            AccountNo string  `json:"account_no"`
            Address string  `json:"address"`
            IfscCode string  `json:"ifsc_code"`
            Wallet string  `json:"wallet"`
         
    }
    
    // AddBeneficiaryDetailsRequest ...
    type AddBeneficiaryDetailsRequest struct {

        
            OrderID string  `json:"order_id"`
            Otp string  `json:"otp"`
            Delights bool  `json:"delights"`
            TransferMode string  `json:"transfer_mode"`
            ShipmentID string  `json:"shipment_id"`
            Details BeneficiaryModeDetails  `json:"details"`
            RequestID string  `json:"request_id"`
         
    }
    
    // RefundAccountResponse ...
    type RefundAccountResponse struct {

        
            Message string  `json:"message"`
            IsVerifiedFlag bool  `json:"is_verified_flag"`
            Success bool  `json:"success"`
            Data map[string]interface{}  `json:"data"`
         
    }
    
    // BankDetailsForOTP ...
    type BankDetailsForOTP struct {

        
            BankName string  `json:"bank_name"`
            BranchName string  `json:"branch_name"`
            AccountHolder string  `json:"account_holder"`
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

        
            RequestID string  `json:"request_id"`
            Success bool  `json:"success"`
            IsVerifiedFlag string  `json:"is_verified_flag"`
         
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

        
            ExternalOrderID string  `json:"external_order_id"`
            MerchantName string  `json:"merchant_name"`
            StatusCode float64  `json:"status_code"`
            Success bool  `json:"success"`
            Amount float64  `json:"amount"`
            PaymentLinkCurrentStatus string  `json:"payment_link_current_status"`
            Message string  `json:"message"`
            PollingTimeout float64  `json:"polling_timeout"`
            PaymentLinkURL string  `json:"payment_link_url"`
         
    }
    
    // ErrorDescription ...
    type ErrorDescription struct {

        
            InvalidID bool  `json:"invalid_id"`
            MerchantName string  `json:"merchant_name"`
            Cancelled bool  `json:"cancelled"`
            Msg string  `json:"msg"`
            Amount float64  `json:"amount"`
            Expired bool  `json:"expired"`
            MerchantOrderID string  `json:"merchant_order_id"`
            PaymentTransactionID string  `json:"payment_transaction_id"`
         
    }
    
    // ErrorResponse ...
    type ErrorResponse struct {

        
            Message string  `json:"message"`
            Error ErrorDescription  `json:"error"`
            Success bool  `json:"success"`
            StatusCode float64  `json:"status_code"`
         
    }
    
    // CreatePaymentLinkMeta ...
    type CreatePaymentLinkMeta struct {

        
            CheckoutMode string  `json:"checkout_mode"`
            CartID string  `json:"cart_id"`
            Pincode string  `json:"pincode"`
            AssignCardID string  `json:"assign_card_id"`
            Amount string  `json:"amount"`
         
    }
    
    // CreatePaymentLinkRequest ...
    type CreatePaymentLinkRequest struct {

        
            ExternalOrderID string  `json:"external_order_id"`
            Email string  `json:"email"`
            Meta CreatePaymentLinkMeta  `json:"meta"`
            Amount float64  `json:"amount"`
            Description string  `json:"description"`
            MobileNumber string  `json:"mobile_number"`
         
    }
    
    // CreatePaymentLinkResponse ...
    type CreatePaymentLinkResponse struct {

        
            StatusCode float64  `json:"status_code"`
            Success bool  `json:"success"`
            PaymentLinkID string  `json:"payment_link_id"`
            Message string  `json:"message"`
            PollingTimeout float64  `json:"polling_timeout"`
            PaymentLinkURL string  `json:"payment_link_url"`
         
    }
    
    // CancelOrResendPaymentLinkRequest ...
    type CancelOrResendPaymentLinkRequest struct {

        
            PaymentLinkID string  `json:"payment_link_id"`
         
    }
    
    // ResendPaymentLinkResponse ...
    type ResendPaymentLinkResponse struct {

        
            Message string  `json:"message"`
            PollingTimeout float64  `json:"polling_timeout"`
            Success bool  `json:"success"`
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
            StatusCode float64  `json:"status_code"`
            Success bool  `json:"success"`
            Amount float64  `json:"amount"`
            AggregatorName string  `json:"aggregator_name"`
            RedirectURL string  `json:"redirect_url"`
            PaymentLinkID string  `json:"payment_link_id"`
            Message string  `json:"message"`
            Status string  `json:"status"`
            HttpStatus float64  `json:"http_status"`
         
    }
    
    // PaymentMethodsMeta ...
    type PaymentMethodsMeta struct {

        
            MerchantCode string  `json:"merchant_code"`
            PaymentIdentifier string  `json:"payment_identifier"`
            PaymentGateway string  `json:"payment_gateway"`
         
    }
    
    // CreateOrderUserPaymentMethods ...
    type CreateOrderUserPaymentMethods struct {

        
            Mode string  `json:"mode"`
            Meta PaymentMethodsMeta  `json:"meta"`
            Name string  `json:"name"`
         
    }
    
    // CreateOrderUserRequest ...
    type CreateOrderUserRequest struct {

        
            Currency string  `json:"currency"`
            PaymentMethods CreateOrderUserPaymentMethods  `json:"payment_methods"`
            Meta map[string]interface{}  `json:"meta"`
            SuccessCallbackURL string  `json:"success_callback_url"`
            FailureCallbackURL string  `json:"failure_callback_url"`
            PaymentLinkID string  `json:"payment_link_id"`
         
    }
    
    // CreateOrderUserData ...
    type CreateOrderUserData struct {

        
            CallbackURL string  `json:"callback_url"`
            Contact string  `json:"contact"`
            OrderID string  `json:"order_id"`
            Currency string  `json:"currency"`
            Email string  `json:"email"`
            CustomerID string  `json:"customer_id"`
            Amount float64  `json:"amount"`
            Aggregator string  `json:"aggregator"`
            MerchantOrderID string  `json:"merchant_order_id"`
            Method string  `json:"method"`
         
    }
    
    // CreateOrderUserResponse ...
    type CreateOrderUserResponse struct {

        
            CallbackURL string  `json:"callback_url"`
            OrderID string  `json:"order_id"`
            StatusCode float64  `json:"status_code"`
            Success bool  `json:"success"`
            PaymentConfirmURL string  `json:"payment_confirm_url"`
            Message string  `json:"message"`
            Data CreateOrderUserData  `json:"data"`
         
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
            StatusMessage string  `json:"status_message"`
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

        
            OwnershipType string  `json:"ownership_type"`
            Pincode string  `json:"pincode"`
            LandMark string  `json:"land_mark"`
            State string  `json:"state"`
            Addressline1 string  `json:"addressline1"`
            City string  `json:"city"`
            Addressline2 string  `json:"addressline2"`
         
    }
    
    // BusinessDetails ...
    type BusinessDetails struct {

        
            BusinessOwnershipType string  `json:"business_ownership_type"`
            Fda string  `json:"fda"`
            Pan string  `json:"pan"`
            Fssai string  `json:"fssai"`
            Gstin string  `json:"gstin"`
            Vintage string  `json:"vintage"`
            Address KYCAddress  `json:"address"`
            EntityType string  `json:"entity_type"`
            BusinessType string  `json:"business_type"`
            ShopAndEstablishment map[string]interface{}  `json:"shop_and_establishment"`
            Name string  `json:"name"`
         
    }
    
    // DeviceDetails ...
    type DeviceDetails struct {

        
            OsVersion string  `json:"os_version"`
            DeviceMake string  `json:"device_make"`
            IdentificationNumber string  `json:"identification_number"`
            IdentifierType string  `json:"identifier_type"`
            DeviceModel string  `json:"device_model"`
            DeviceType string  `json:"device_type"`
            Os string  `json:"os"`
         
    }
    
    // MarketplaceInfo ...
    type MarketplaceInfo struct {

        
            MembershipID string  `json:"membership_id"`
            DateOfJoining string  `json:"date_of_joining"`
            Name string  `json:"name"`
         
    }
    
    // UserPersonalInfoInDetails ...
    type UserPersonalInfoInDetails struct {

        
            LastName string  `json:"last_name"`
            MiddleName string  `json:"middle_name"`
            FirstName string  `json:"first_name"`
            MothersName string  `json:"mothers_name"`
            Dob string  `json:"dob"`
            Email string  `json:"email"`
            Gender string  `json:"gender"`
            AddressAsPerID KYCAddress  `json:"address_as_per_id"`
            Pan string  `json:"pan"`
            EmailVerified bool  `json:"email_verified"`
            VoterID string  `json:"voter_id"`
            FathersName string  `json:"fathers_name"`
            Passport string  `json:"passport"`
            DrivingLicense string  `json:"driving_license"`
            MobileVerified bool  `json:"mobile_verified"`
            Phone string  `json:"phone"`
         
    }
    
    // CustomerOnboardingRequest ...
    type CustomerOnboardingRequest struct {

        
            BusinessInfo BusinessDetails  `json:"business_info"`
            Device DeviceDetails  `json:"device"`
            Source string  `json:"source"`
            Mcc string  `json:"mcc"`
            Aggregator string  `json:"aggregator"`
            MarketplaceInfo MarketplaceInfo  `json:"marketplace_info"`
            PersonalInfo UserPersonalInfoInDetails  `json:"personal_info"`
         
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
    

    
    // ShipmentTotalDetails ...
    type ShipmentTotalDetails struct {

        
            Pieces float64  `json:"pieces"`
            Sizes float64  `json:"sizes"`
            TotalPrice float64  `json:"total_price"`
         
    }
    
    // CurrentStatus ...
    type CurrentStatus struct {

        
            JourneyType string  `json:"journey_type"`
            Status string  `json:"status"`
            UpdatedAt string  `json:"updated_at"`
            Name string  `json:"name"`
         
    }
    
    // Prices ...
    type Prices struct {

        
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            PriceEffective float64  `json:"price_effective"`
            AddedToFyndCash bool  `json:"added_to_fynd_cash"`
            Discount float64  `json:"discount"`
            RefundCredit float64  `json:"refund_credit"`
            AmountPaidRoundoff float64  `json:"amount_paid_roundoff"`
            Cashback float64  `json:"cashback"`
            CouponValue float64  `json:"coupon_value"`
            CodCharges float64  `json:"cod_charges"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            TransferPrice float64  `json:"transfer_price"`
            CurrencyCode string  `json:"currency_code"`
            CurrencySymbol string  `json:"currency_symbol"`
            PriceMarked float64  `json:"price_marked"`
            ValueOfGood float64  `json:"value_of_good"`
            AmountPaid float64  `json:"amount_paid"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            FyndCredits float64  `json:"fynd_credits"`
            CashbackApplied float64  `json:"cashback_applied"`
            RefundAmount float64  `json:"refund_amount"`
            PromotionEffectiveDiscount float64  `json:"promotion_effective_discount"`
            DeliveryCharge float64  `json:"delivery_charge"`
         
    }
    
    // Identifiers ...
    type Identifiers struct {

        
            SkuCode string  `json:"sku_code"`
            Ean string  `json:"ean"`
         
    }
    
    // FinancialBreakup ...
    type FinancialBreakup struct {

        
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            Identifiers Identifiers  `json:"identifiers"`
            AddedToFyndCash bool  `json:"added_to_fynd_cash"`
            PriceEffective float64  `json:"price_effective"`
            RefundCredit float64  `json:"refund_credit"`
            Size string  `json:"size"`
            AmountPaidRoundoff float64  `json:"amount_paid_roundoff"`
            Cashback float64  `json:"cashback"`
            Discount float64  `json:"discount"`
            TotalUnits float64  `json:"total_units"`
            CouponValue float64  `json:"coupon_value"`
            CodCharges float64  `json:"cod_charges"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            GstFee float64  `json:"gst_fee"`
            HsnCode string  `json:"hsn_code"`
            TransferPrice float64  `json:"transfer_price"`
            GstTag string  `json:"gst_tag"`
            PriceMarked float64  `json:"price_marked"`
            ItemName string  `json:"item_name"`
            ValueOfGood float64  `json:"value_of_good"`
            AmountPaid float64  `json:"amount_paid"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            FyndCredits float64  `json:"fynd_credits"`
            CashbackApplied float64  `json:"cashback_applied"`
            RefundAmount float64  `json:"refund_amount"`
            PromotionEffectiveDiscount float64  `json:"promotion_effective_discount"`
            DeliveryCharge float64  `json:"delivery_charge"`
         
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

        
            PromotionType string  `json:"promotion_type"`
            PromotionName string  `json:"promotion_name"`
            MrpPromotion bool  `json:"mrp_promotion"`
            PromoID string  `json:"promo_id"`
            Amount float64  `json:"amount"`
            ArticleQuantity float64  `json:"article_quantity"`
            AppliedFreeArticles []AppliedFreeArticles  `json:"applied_free_articles"`
         
    }
    
    // ItemBrand ...
    type ItemBrand struct {

        
            Logo string  `json:"logo"`
            Name string  `json:"name"`
         
    }
    
    // Item ...
    type Item struct {

        
            ID float64  `json:"id"`
            Size string  `json:"size"`
            Code string  `json:"code"`
            SlugKey string  `json:"slug_key"`
            Image []string  `json:"image"`
            Name string  `json:"name"`
            Brand ItemBrand  `json:"brand"`
            SellerIdentifier string  `json:"seller_identifier"`
         
    }
    
    // Bags ...
    type Bags struct {

        
            ReturnableDate string  `json:"returnable_date"`
            CanCancel bool  `json:"can_cancel"`
            ID float64  `json:"id"`
            DeliveryDate string  `json:"delivery_date"`
            CurrencyCode string  `json:"currency_code"`
            CurrentStatus CurrentStatus  `json:"current_status"`
            Prices Prices  `json:"prices"`
            FinancialBreakup []FinancialBreakup  `json:"financial_breakup"`
            LineNumber float64  `json:"line_number"`
            AppliedPromos []AppliedPromos  `json:"applied_promos"`
            CurrencySymbol string  `json:"currency_symbol"`
            CanReturn bool  `json:"can_return"`
            ParentPromoBags map[string]interface{}  `json:"parent_promo_bags"`
            Item Item  `json:"item"`
            Quantity float64  `json:"quantity"`
            SellerIdentifier string  `json:"seller_identifier"`
         
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
    
    // ShipmentUserInfo ...
    type ShipmentUserInfo struct {

        
            LastName string  `json:"last_name"`
            Mobile string  `json:"mobile"`
            Gender string  `json:"gender"`
            FirstName string  `json:"first_name"`
         
    }
    
    // NestedTrackingDetails ...
    type NestedTrackingDetails struct {

        
            Time string  `json:"time"`
            IsCurrent bool  `json:"is_current"`
            Status string  `json:"status"`
            IsPassed bool  `json:"is_passed"`
         
    }
    
    // TrackingDetails ...
    type TrackingDetails struct {

        
            IsPassed bool  `json:"is_passed"`
            Time string  `json:"time"`
            TrackingDetails []NestedTrackingDetails  `json:"tracking_details"`
            Status string  `json:"status"`
            IsCurrent bool  `json:"is_current"`
         
    }
    
    // ShipmentStatus ...
    type ShipmentStatus struct {

        
            Title string  `json:"title"`
            HexCode string  `json:"hex_code"`
         
    }
    
    // ShipmentPayment ...
    type ShipmentPayment struct {

        
            DisplayName string  `json:"display_name"`
            PaymentMode string  `json:"payment_mode"`
            Mode string  `json:"mode"`
            Logo string  `json:"logo"`
            Status string  `json:"status"`
            Mop string  `json:"mop"`
         
    }
    
    // DeliveryAddress ...
    type DeliveryAddress struct {

        
            Email string  `json:"email"`
            Area string  `json:"area"`
            Address2 string  `json:"address2"`
            Pincode string  `json:"pincode"`
            CountryIsoCode string  `json:"country_iso_code"`
            ContactPerson string  `json:"contact_person"`
            Version string  `json:"version"`
            CreatedAt string  `json:"created_at"`
            Name string  `json:"name"`
            CountryPhoneCode string  `json:"country_phone_code"`
            Longitude float64  `json:"longitude"`
            Address string  `json:"address"`
            Address1 string  `json:"address1"`
            State string  `json:"state"`
            City string  `json:"city"`
            UpdatedAt string  `json:"updated_at"`
            AddressType string  `json:"address_type"`
            AddressCategory string  `json:"address_category"`
            Landmark string  `json:"landmark"`
            Country string  `json:"country"`
            Latitude float64  `json:"latitude"`
            Phone string  `json:"phone"`
         
    }
    
    // BreakupValues ...
    type BreakupValues struct {

        
            Display string  `json:"display"`
            CurrencyCode string  `json:"currency_code"`
            CurrencySymbol string  `json:"currency_symbol"`
            Name string  `json:"name"`
            Value float64  `json:"value"`
         
    }
    
    // FulfillingStore ...
    type FulfillingStore struct {

        
            CompanyName string  `json:"company_name"`
            ID float64  `json:"id"`
            CompanyID float64  `json:"company_id"`
            Code string  `json:"code"`
            Name string  `json:"name"`
         
    }
    
    // Invoice ...
    type Invoice struct {

        
            UpdatedDate string  `json:"updated_date"`
            InvoiceURL string  `json:"invoice_url"`
            LabelURL string  `json:"label_url"`
         
    }
    
    // FulfillingCompany ...
    type FulfillingCompany struct {

        
            ID float64  `json:"id"`
            Name string  `json:"name"`
         
    }
    
    // Shipments ...
    type Shipments struct {

        
            TotalDetails ShipmentTotalDetails  `json:"total_details"`
            NeedHelpURL string  `json:"need_help_url"`
            OrderType string  `json:"order_type"`
            Bags []Bags  `json:"bags"`
            Prices Prices  `json:"prices"`
            Promise Promise  `json:"promise"`
            UserInfo ShipmentUserInfo  `json:"user_info"`
            TrackingDetails []TrackingDetails  `json:"tracking_details"`
            Comment string  `json:"comment"`
            TrackURL string  `json:"track_url"`
            ReturnableDate string  `json:"returnable_date"`
            CanCancel bool  `json:"can_cancel"`
            CustomMeta []map[string]interface{}  `json:"custom_meta"`
            DeliveryDate string  `json:"delivery_date"`
            ShowTrackLink bool  `json:"show_track_link"`
            TotalBags float64  `json:"total_bags"`
            RefundDetails map[string]interface{}  `json:"refund_details"`
            ShipmentStatus ShipmentStatus  `json:"shipment_status"`
            BeneficiaryDetails bool  `json:"beneficiary_details"`
            ShowDownloadInvoice bool  `json:"show_download_invoice"`
            AwbNo string  `json:"awb_no"`
            DpName string  `json:"dp_name"`
            Payment ShipmentPayment  `json:"payment"`
            DeliveryAddress DeliveryAddress  `json:"delivery_address"`
            SizeInfo map[string]interface{}  `json:"size_info"`
            BreakupValues []BreakupValues  `json:"breakup_values"`
            FulfillingStore FulfillingStore  `json:"fulfilling_store"`
            TrakingNo string  `json:"traking_no"`
            CanBreak map[string]interface{}  `json:"can_break"`
            OrderID string  `json:"order_id"`
            ShipmentID string  `json:"shipment_id"`
            CanReturn bool  `json:"can_return"`
            Invoice Invoice  `json:"invoice"`
            ShipmentCreatedAt string  `json:"shipment_created_at"`
            FulfillingCompany FulfillingCompany  `json:"fulfilling_company"`
            ReturnMeta map[string]interface{}  `json:"return_meta"`
         
    }
    
    // UserInfo ...
    type UserInfo struct {

        
            Email string  `json:"email"`
            Mobile string  `json:"mobile"`
            Gender string  `json:"gender"`
            Name string  `json:"name"`
         
    }
    
    // BagsForReorderArticleAssignment ...
    type BagsForReorderArticleAssignment struct {

        
            Level string  `json:"level"`
            Strategy string  `json:"strategy"`
         
    }
    
    // BagsForReorder ...
    type BagsForReorder struct {

        
            ItemSize string  `json:"item_size"`
            StoreID float64  `json:"store_id"`
            ItemID float64  `json:"item_id"`
            SellerID float64  `json:"seller_id"`
            Quantity float64  `json:"quantity"`
            ArticleAssignment BagsForReorderArticleAssignment  `json:"article_assignment"`
         
    }
    
    // OrderSchema ...
    type OrderSchema struct {

        
            Shipments []Shipments  `json:"shipments"`
            OrderCreatedTime string  `json:"order_created_time"`
            TotalShipmentsInOrder float64  `json:"total_shipments_in_order"`
            OrderID string  `json:"order_id"`
            UserInfo UserInfo  `json:"user_info"`
            BagsForReorder []BagsForReorder  `json:"bags_for_reorder"`
            BreakupValues []BreakupValues  `json:"breakup_values"`
         
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
    
    // OrderPage ...
    type OrderPage struct {

        
            ItemTotal float64  `json:"item_total"`
            HasNext bool  `json:"has_next"`
            Size float64  `json:"size"`
            Current float64  `json:"current"`
            Type string  `json:"type"`
         
    }
    
    // OrderList ...
    type OrderList struct {

        
            Items []OrderSchema  `json:"items"`
            Filters OrderFilters  `json:"filters"`
            Page OrderPage  `json:"page"`
         
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
            ShipmentID string  `json:"shipment_id"`
            PresignedURL string  `json:"presigned_url"`
            PresignedType string  `json:"presigned_type"`
         
    }
    
    // Track ...
    type Track struct {

        
            ShipmentType string  `json:"shipment_type"`
            LastLocationRecievedAt string  `json:"last_location_recieved_at"`
            UpdatedTime string  `json:"updated_time"`
            AccountName string  `json:"account_name"`
            Reason string  `json:"reason"`
            UpdatedAt string  `json:"updated_at"`
            Status string  `json:"status"`
            Awb string  `json:"awb"`
         
    }
    
    // ShipmentTrack ...
    type ShipmentTrack struct {

        
            Results []Track  `json:"results"`
         
    }
    
    // CustomerDetailsResponse ...
    type CustomerDetailsResponse struct {

        
            Country string  `json:"country"`
            OrderID string  `json:"order_id"`
            ShipmentID string  `json:"shipment_id"`
            Name string  `json:"name"`
            Phone string  `json:"phone"`
         
    }
    
    // SendOtpToCustomerResponse ...
    type SendOtpToCustomerResponse struct {

        
            Success bool  `json:"success"`
            ResendTimer float64  `json:"resend_timer"`
            Message string  `json:"message"`
            RequestID string  `json:"request_id"`
         
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
            ID float64  `json:"id"`
            QuestionSet []QuestionSet  `json:"question_set"`
            QcType []string  `json:"qc_type"`
            Reasons []BagReasons  `json:"reasons"`
         
    }
    
    // ShipmentBagReasons ...
    type ShipmentBagReasons struct {

        
            Success bool  `json:"success"`
            Reasons []BagReasons  `json:"reasons"`
         
    }
    
    // ShipmentReason ...
    type ShipmentReason struct {

        
            FeedbackType string  `json:"feedback_type"`
            ReasonID float64  `json:"reason_id"`
            Flow string  `json:"flow"`
            Priority float64  `json:"priority"`
            ReasonText string  `json:"reason_text"`
            ShowTextArea bool  `json:"show_text_area"`
         
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
            Identifier string  `json:"identifier"`
            Products []Products  `json:"products"`
            Reasons ReasonsData  `json:"reasons"`
         
    }
    
    // StatuesRequest ...
    type StatuesRequest struct {

        
            ExcludeBagsNextState string  `json:"exclude_bags_next_state"`
            Status string  `json:"status"`
            Shipments []ShipmentsRequest  `json:"shipments"`
         
    }
    
    // UpdateShipmentStatusRequest ...
    type UpdateShipmentStatusRequest struct {

        
            UnlockBeforeTransition bool  `json:"unlock_before_transition"`
            ForceTransition bool  `json:"force_transition"`
            Statuses []StatuesRequest  `json:"statuses"`
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

        
            Message string  `json:"message"`
            Exception string  `json:"exception"`
            Code string  `json:"code"`
            StackTrace string  `json:"stack_trace"`
            Status float64  `json:"status"`
         
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
    

    
    // PaymentSelectionLock ...
    type PaymentSelectionLock struct {

        
            Enabled bool  `json:"enabled"`
            DefaultOptions string  `json:"default_options"`
            PaymentIdentifier string  `json:"payment_identifier"`
         
    }
    
    // Ownership ...
    type Ownership struct {

        
            PayableCategory string  `json:"payable_category"`
            PayableBy string  `json:"payable_by"`
         
    }
    
    // AppliedPromotion ...
    type AppliedPromotion struct {

        
            Amount float64  `json:"amount"`
            PromoID string  `json:"promo_id"`
            OfferText string  `json:"offer_text"`
            Ownership Ownership  `json:"ownership"`
            PromotionType string  `json:"promotion_type"`
            ArticleQuantity float64  `json:"article_quantity"`
            MrpPromotion bool  `json:"mrp_promotion"`
         
    }
    
    // ProductPrice ...
    type ProductPrice struct {

        
            Selling float64  `json:"selling"`
            Marked float64  `json:"marked"`
            CurrencyCode string  `json:"currency_code"`
            Effective float64  `json:"effective"`
            AddOn float64  `json:"add_on"`
            CurrencySymbol string  `json:"currency_symbol"`
         
    }
    
    // ProductPriceInfo ...
    type ProductPriceInfo struct {

        
            Converted ProductPrice  `json:"converted"`
            Base ProductPrice  `json:"base"`
         
    }
    
    // CategoryInfo ...
    type CategoryInfo struct {

        
            Name string  `json:"name"`
            UID float64  `json:"uid"`
         
    }
    
    // BaseInfo ...
    type BaseInfo struct {

        
            Name string  `json:"name"`
            UID float64  `json:"uid"`
         
    }
    
    // ProductImage ...
    type ProductImage struct {

        
            AspectRatio string  `json:"aspect_ratio"`
            SecureURL string  `json:"secure_url"`
            URL string  `json:"url"`
         
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

        
            Categories []CategoryInfo  `json:"categories"`
            Brand BaseInfo  `json:"brand"`
            Slug string  `json:"slug"`
            Images []ProductImage  `json:"images"`
            Name string  `json:"name"`
            Action ProductAction  `json:"action"`
            UID float64  `json:"uid"`
            Type string  `json:"type"`
         
    }
    
    // PromoMeta ...
    type PromoMeta struct {

        
            Message string  `json:"message"`
         
    }
    
    // CartProductIdentifer ...
    type CartProductIdentifer struct {

        
            Identifier string  `json:"identifier"`
         
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
    
    // BasePrice ...
    type BasePrice struct {

        
            CurrencySymbol string  `json:"currency_symbol"`
            Marked float64  `json:"marked"`
            CurrencyCode string  `json:"currency_code"`
            Effective float64  `json:"effective"`
         
    }
    
    // ArticlePriceInfo ...
    type ArticlePriceInfo struct {

        
            Converted BasePrice  `json:"converted"`
            Base BasePrice  `json:"base"`
         
    }
    
    // ProductArticle ...
    type ProductArticle struct {

        
            Price ArticlePriceInfo  `json:"price"`
            ProductGroupTags []string  `json:"product_group_tags"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            Seller BaseInfo  `json:"seller"`
            Store BaseInfo  `json:"store"`
            Quantity float64  `json:"quantity"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            UID string  `json:"uid"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Size string  `json:"size"`
            Type string  `json:"type"`
         
    }
    
    // ProductAvailabilitySize ...
    type ProductAvailabilitySize struct {

        
            Value string  `json:"value"`
            IsAvailable bool  `json:"is_available"`
            Display string  `json:"display"`
         
    }
    
    // ProductAvailability ...
    type ProductAvailability struct {

        
            Deliverable bool  `json:"deliverable"`
            OutOfStock bool  `json:"out_of_stock"`
            Sizes []string  `json:"sizes"`
            AvailableSizes []ProductAvailabilitySize  `json:"available_sizes"`
            OtherStoreQuantity float64  `json:"other_store_quantity"`
            IsValid bool  `json:"is_valid"`
         
    }
    
    // CartProductInfo ...
    type CartProductInfo struct {

        
            Message string  `json:"message"`
            PromotionsApplied []AppliedPromotion  `json:"promotions_applied"`
            Moq map[string]interface{}  `json:"moq"`
            Price ProductPriceInfo  `json:"price"`
            Discount string  `json:"discount"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            Product CartProduct  `json:"product"`
            CouponMessage string  `json:"coupon_message"`
            Key string  `json:"key"`
            BulkOffer map[string]interface{}  `json:"bulk_offer"`
            Quantity float64  `json:"quantity"`
            IsSet bool  `json:"is_set"`
            PromoMeta PromoMeta  `json:"promo_meta"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            Article ProductArticle  `json:"article"`
            Availability ProductAvailability  `json:"availability"`
            PricePerUnit ProductPriceInfo  `json:"price_per_unit"`
         
    }
    
    // CartCurrency ...
    type CartCurrency struct {

        
            Symbol string  `json:"symbol"`
            Code string  `json:"code"`
         
    }
    
    // CouponBreakup ...
    type CouponBreakup struct {

        
            Message string  `json:"message"`
            Value float64  `json:"value"`
            Code string  `json:"code"`
            UID string  `json:"uid"`
            IsApplied bool  `json:"is_applied"`
            Type string  `json:"type"`
         
    }
    
    // LoyaltyPoints ...
    type LoyaltyPoints struct {

        
            Applicable float64  `json:"applicable"`
            IsApplied bool  `json:"is_applied"`
            Description string  `json:"description"`
            Total float64  `json:"total"`
         
    }
    
    // RawBreakup ...
    type RawBreakup struct {

        
            DeliveryCharge float64  `json:"delivery_charge"`
            Discount float64  `json:"discount"`
            FyndCash float64  `json:"fynd_cash"`
            Coupon float64  `json:"coupon"`
            ConvenienceFee float64  `json:"convenience_fee"`
            GstCharges float64  `json:"gst_charges"`
            Vog float64  `json:"vog"`
            CodCharge float64  `json:"cod_charge"`
            YouSaved float64  `json:"you_saved"`
            MrpTotal float64  `json:"mrp_total"`
            Subtotal float64  `json:"subtotal"`
            Total float64  `json:"total"`
         
    }
    
    // DisplayBreakup ...
    type DisplayBreakup struct {

        
            Message []string  `json:"message"`
            Display string  `json:"display"`
            CurrencyCode string  `json:"currency_code"`
            Key string  `json:"key"`
            Value float64  `json:"value"`
            CurrencySymbol string  `json:"currency_symbol"`
         
    }
    
    // CartBreakup ...
    type CartBreakup struct {

        
            Coupon CouponBreakup  `json:"coupon"`
            LoyaltyPoints LoyaltyPoints  `json:"loyalty_points"`
            Raw RawBreakup  `json:"raw"`
            Display []DisplayBreakup  `json:"display"`
         
    }
    
    // CartDetailResponse ...
    type CartDetailResponse struct {

        
            Message string  `json:"message"`
            CheckoutMode string  `json:"checkout_mode"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            IsValid bool  `json:"is_valid"`
            Items []CartProductInfo  `json:"items"`
            Currency CartCurrency  `json:"currency"`
            AppliedPromoDetails []AppliedPromotion  `json:"applied_promo_details"`
            Gstin string  `json:"gstin"`
            Comment string  `json:"comment"`
            CouponText string  `json:"coupon_text"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            BuyNow bool  `json:"buy_now"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            ID string  `json:"id"`
            LastModified string  `json:"last_modified"`
         
    }
    
    // AddProductCart ...
    type AddProductCart struct {

        
            Display string  `json:"display"`
            ArticleAssignment map[string]interface{}  `json:"article_assignment"`
            ItemID float64  `json:"item_id"`
            ParentItemIdentifiers []map[string]string  `json:"parent_item_identifiers"`
            ProductGroupTags []string  `json:"product_group_tags"`
            SellerID float64  `json:"seller_id"`
            ItemSize string  `json:"item_size"`
            Quantity float64  `json:"quantity"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            ArticleID string  `json:"article_id"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            StoreID float64  `json:"store_id"`
            Pos bool  `json:"pos"`
         
    }
    
    // AddCartRequest ...
    type AddCartRequest struct {

        
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

        
            ItemID float64  `json:"item_id"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            ItemSize string  `json:"item_size"`
            Quantity float64  `json:"quantity"`
            ItemIndex float64  `json:"item_index"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            ArticleID string  `json:"article_id"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
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
    
    // Coupon ...
    type Coupon struct {

        
            MinimumCartValue float64  `json:"minimum_cart_value"`
            Message string  `json:"message"`
            IsApplicable bool  `json:"is_applicable"`
            ExpiresOn string  `json:"expires_on"`
            MaxDiscountValue float64  `json:"max_discount_value"`
            Title string  `json:"title"`
            CouponCode string  `json:"coupon_code"`
            CouponValue float64  `json:"coupon_value"`
            SubTitle string  `json:"sub_title"`
            IsApplied bool  `json:"is_applied"`
         
    }
    
    // PageCoupon ...
    type PageCoupon struct {

        
            Current float64  `json:"current"`
            TotalItemCount float64  `json:"total_item_count"`
            HasNext bool  `json:"has_next"`
            HasPrevious bool  `json:"has_previous"`
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

        
            Name string  `json:"name"`
            UID float64  `json:"uid"`
         
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
            Price OfferPrice  `json:"price"`
            Quantity float64  `json:"quantity"`
            Margin float64  `json:"margin"`
            Best bool  `json:"best"`
            Type string  `json:"type"`
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

        
            Longitude float64  `json:"longitude"`
            Latitude float64  `json:"latitude"`
         
    }
    
    // Address ...
    type Address struct {

        
            GoogleMapPoint map[string]interface{}  `json:"google_map_point"`
            Address string  `json:"address"`
            CountryCode string  `json:"country_code"`
            Tags []string  `json:"tags"`
            UserID string  `json:"user_id"`
            ID string  `json:"id"`
            GeoLocation GeoLocation  `json:"geo_location"`
            Phone string  `json:"phone"`
            Meta map[string]interface{}  `json:"meta"`
            Area string  `json:"area"`
            AddressType string  `json:"address_type"`
            IsDefaultAddress bool  `json:"is_default_address"`
            AreaCode string  `json:"area_code"`
            Email string  `json:"email"`
            CountryPhoneCode string  `json:"country_phone_code"`
            City string  `json:"city"`
            CountryIsoCode string  `json:"country_iso_code"`
            CheckoutMode string  `json:"checkout_mode"`
            AreaCodeSlug string  `json:"area_code_slug"`
            Country string  `json:"country"`
            State string  `json:"state"`
            Name string  `json:"name"`
            IsActive bool  `json:"is_active"`
            Landmark string  `json:"landmark"`
         
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
            IsUpdated bool  `json:"is_updated"`
            ID string  `json:"id"`
            IsDefaultAddress bool  `json:"is_default_address"`
         
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
            AggregatorName string  `json:"aggregator_name"`
            PaymentIdentifier string  `json:"payment_identifier"`
            PaymentMode string  `json:"payment_mode"`
            ID string  `json:"id"`
            MerchantCode string  `json:"merchant_code"`
         
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

        
            Message string  `json:"message"`
            CouponValidity CouponValidity  `json:"coupon_validity"`
            Success bool  `json:"success"`
         
    }
    
    // ShipmentResponse ...
    type ShipmentResponse struct {

        
            DpID string  `json:"dp_id"`
            Items []CartProductInfo  `json:"items"`
            DpOptions map[string]interface{}  `json:"dp_options"`
            FulfillmentType string  `json:"fulfillment_type"`
            ShipmentType string  `json:"shipment_type"`
            Promise ShipmentPromise  `json:"promise"`
            FulfillmentID float64  `json:"fulfillment_id"`
            OrderType string  `json:"order_type"`
            Shipments float64  `json:"shipments"`
            BoxType string  `json:"box_type"`
         
    }
    
    // CartShipmentsResponse ...
    type CartShipmentsResponse struct {

        
            Message string  `json:"message"`
            Currency CartCurrency  `json:"currency"`
            CouponText string  `json:"coupon_text"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            BuyNow bool  `json:"buy_now"`
            Shipments []ShipmentResponse  `json:"shipments"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            IsValid bool  `json:"is_valid"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            ID string  `json:"id"`
            LastModified string  `json:"last_modified"`
            CheckoutMode string  `json:"checkout_mode"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            Gstin string  `json:"gstin"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            UID string  `json:"uid"`
            CartID float64  `json:"cart_id"`
            Comment string  `json:"comment"`
            Error bool  `json:"error"`
         
    }
    
    // UpdateCartShipmentItem ...
    type UpdateCartShipmentItem struct {

        
            ArticleUID string  `json:"article_uid"`
            Quantity float64  `json:"quantity"`
            ShipmentType string  `json:"shipment_type"`
         
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
            LastName string  `json:"last_name"`
            FirstName string  `json:"first_name"`
         
    }
    
    // CartCheckoutCustomMeta ...
    type CartCheckoutCustomMeta struct {

        
            Value string  `json:"value"`
            Key string  `json:"key"`
         
    }
    
    // CartPosCheckoutDetailRequest ...
    type CartPosCheckoutDetailRequest struct {

        
            AddressID string  `json:"address_id"`
            PaymentIdentifier string  `json:"payment_identifier"`
            Files []Files  `json:"files"`
            Aggregator string  `json:"aggregator"`
            OrderingStore float64  `json:"ordering_store"`
            PickAtStoreUID float64  `json:"pick_at_store_uid"`
            CallbackURL string  `json:"callback_url"`
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
            PaymentAutoConfirm bool  `json:"payment_auto_confirm"`
            Pos bool  `json:"pos"`
            OrderType string  `json:"order_type"`
            Meta map[string]interface{}  `json:"meta"`
            BillingAddressID string  `json:"billing_address_id"`
            MerchantCode string  `json:"merchant_code"`
            Staff StaffCheckout  `json:"staff"`
            PaymentParams map[string]interface{}  `json:"payment_params"`
            PaymentMode string  `json:"payment_mode"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            CustomMeta []CartCheckoutCustomMeta  `json:"custom_meta"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
         
    }
    
    // CheckCart ...
    type CheckCart struct {

        
            Items []CartProductInfo  `json:"items"`
            CodMessage string  `json:"cod_message"`
            DeliveryChargeOrderValue float64  `json:"delivery_charge_order_value"`
            OrderID string  `json:"order_id"`
            BuyNow bool  `json:"buy_now"`
            UserType string  `json:"user_type"`
            StoreEmps []map[string]interface{}  `json:"store_emps"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            IsValid bool  `json:"is_valid"`
            ID string  `json:"id"`
            CodAvailable bool  `json:"cod_available"`
            CodCharges float64  `json:"cod_charges"`
            StoreCode string  `json:"store_code"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            Success bool  `json:"success"`
            CartID float64  `json:"cart_id"`
            UID string  `json:"uid"`
            Comment string  `json:"comment"`
            DeliveryCharges float64  `json:"delivery_charges"`
            Message string  `json:"message"`
            Currency CartCurrency  `json:"currency"`
            CouponText string  `json:"coupon_text"`
            ErrorMessage string  `json:"error_message"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            LastModified string  `json:"last_modified"`
            CheckoutMode string  `json:"checkout_mode"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            Gstin string  `json:"gstin"`
            BreakupValues CartBreakup  `json:"breakup_values"`
         
    }
    
    // CartCheckoutResponse ...
    type CartCheckoutResponse struct {

        
            Message string  `json:"message"`
            Cart CheckCart  `json:"cart"`
            AppInterceptURL string  `json:"app_intercept_url"`
            OrderID string  `json:"order_id"`
            Data map[string]interface{}  `json:"data"`
            CallbackURL string  `json:"callback_url"`
            Success bool  `json:"success"`
            PaymentConfirmURL string  `json:"payment_confirm_url"`
         
    }
    
    // CartMetaRequest ...
    type CartMetaRequest struct {

        
            CheckoutMode string  `json:"checkout_mode"`
            PickUpCustomerDetails map[string]interface{}  `json:"pick_up_customer_details"`
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

        
            AreaCode string  `json:"area_code"`
            Address string  `json:"address"`
            Phone string  `json:"phone"`
            Email string  `json:"email"`
            AreaCodeSlug string  `json:"area_code_slug"`
            StoreCode string  `json:"store_code"`
            Pincode float64  `json:"pincode"`
            Country string  `json:"country"`
            State string  `json:"state"`
            ID float64  `json:"id"`
            Name string  `json:"name"`
            City string  `json:"city"`
            Area string  `json:"area"`
            UID float64  `json:"uid"`
            AddressType string  `json:"address_type"`
            Landmark string  `json:"landmark"`
         
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

        
            CreatedOn string  `json:"created_on"`
            Meta map[string]interface{}  `json:"meta"`
            Source map[string]interface{}  `json:"source"`
            User map[string]interface{}  `json:"user"`
            Token string  `json:"token"`
         
    }
    
    // SharedCart ...
    type SharedCart struct {

        
            Items []CartProductInfo  `json:"items"`
            BuyNow bool  `json:"buy_now"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            IsValid bool  `json:"is_valid"`
            ID string  `json:"id"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            UID string  `json:"uid"`
            CartID float64  `json:"cart_id"`
            Comment string  `json:"comment"`
            Message string  `json:"message"`
            Currency CartCurrency  `json:"currency"`
            CouponText string  `json:"coupon_text"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            SharedCartDetails SharedCartDetails  `json:"shared_cart_details"`
            LastModified string  `json:"last_modified"`
            CheckoutMode string  `json:"checkout_mode"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            Gstin string  `json:"gstin"`
            BreakupValues CartBreakup  `json:"breakup_values"`
         
    }
    
    // SharedCartResponse ...
    type SharedCartResponse struct {

        
            Cart SharedCart  `json:"cart"`
            Error string  `json:"error"`
         
    }
    

    
    // PincodeParentsResponse ...
    type PincodeParentsResponse struct {

        
            UID string  `json:"uid"`
            SubType string  `json:"sub_type"`
            DisplayName string  `json:"display_name"`
            Name string  `json:"name"`
         
    }
    
    // PincodeLatLongData ...
    type PincodeLatLongData struct {

        
            Coordinates []string  `json:"coordinates"`
            Type string  `json:"type"`
         
    }
    
    // CountryMetaResponse ...
    type CountryMetaResponse struct {

        
            CountryCode string  `json:"country_code"`
            IsdCode string  `json:"isd_code"`
         
    }
    
    // PincodeMetaResponse ...
    type PincodeMetaResponse struct {

        
            Zone string  `json:"zone"`
            InternalZoneID float64  `json:"internal_zone_id"`
         
    }
    
    // PincodeErrorSchemaResponse ...
    type PincodeErrorSchemaResponse struct {

        
            Type string  `json:"type"`
            Message string  `json:"message"`
            Value string  `json:"value"`
         
    }
    
    // PincodeDataResponse ...
    type PincodeDataResponse struct {

        
            SubType string  `json:"sub_type"`
            Parents []PincodeParentsResponse  `json:"parents"`
            LatLong PincodeLatLongData  `json:"lat_long"`
            MetaCode CountryMetaResponse  `json:"meta_code"`
            Name string  `json:"name"`
            UID string  `json:"uid"`
            Meta PincodeMetaResponse  `json:"meta"`
            Error PincodeErrorSchemaResponse  `json:"error"`
            DisplayName string  `json:"display_name"`
         
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

        
            Category TATCategoryRequest  `json:"category"`
            ManufacturingTimeUnit string  `json:"manufacturing_time_unit"`
            ManufacturingTime float64  `json:"manufacturing_time"`
         
    }
    
    // TATLocationDetailsRequest ...
    type TATLocationDetailsRequest struct {

        
            FromPincode string  `json:"from_pincode"`
            Articles []TATArticlesRequest  `json:"articles"`
            FulfillmentID float64  `json:"fulfillment_id"`
         
    }
    
    // TATViewRequest ...
    type TATViewRequest struct {

        
            ToPincode string  `json:"to_pincode"`
            Journey string  `json:"journey"`
            Action string  `json:"action"`
            Source string  `json:"source"`
            LocationDetails []TATLocationDetailsRequest  `json:"location_details"`
            Identifier string  `json:"identifier"`
         
    }
    
    // TATTimestampResponse ...
    type TATTimestampResponse struct {

        
            Min float64  `json:"min"`
            Max float64  `json:"max"`
         
    }
    
    // TATFormattedResponse ...
    type TATFormattedResponse struct {

        
            Min string  `json:"min"`
            Max string  `json:"max"`
         
    }
    
    // TATPromiseResponse ...
    type TATPromiseResponse struct {

        
            Timestamp TATTimestampResponse  `json:"timestamp"`
            Formatted TATFormattedResponse  `json:"formatted"`
         
    }
    
    // TATErrorSchemaResponse ...
    type TATErrorSchemaResponse struct {

        
            Type string  `json:"type"`
            Message string  `json:"message"`
            Value string  `json:"value"`
         
    }
    
    // TATArticlesResponse ...
    type TATArticlesResponse struct {

        
            Category TATCategoryRequest  `json:"category"`
            Promise TATPromiseResponse  `json:"promise"`
            ManufacturingTimeUnit string  `json:"manufacturing_time_unit"`
            IsCodAvailable bool  `json:"is_cod_available"`
            ManufacturingTimeSeconds float64  `json:"_manufacturing_time_seconds"`
            Error TATErrorSchemaResponse  `json:"error"`
            ManufacturingTime float64  `json:"manufacturing_time"`
         
    }
    
    // TATLocationDetailsResponse ...
    type TATLocationDetailsResponse struct {

        
            FromPincode string  `json:"from_pincode"`
            Articles []TATArticlesResponse  `json:"articles"`
            FulfillmentID float64  `json:"fulfillment_id"`
         
    }
    
    // TATViewResponse ...
    type TATViewResponse struct {

        
            ToPincode string  `json:"to_pincode"`
            Success bool  `json:"success"`
            RequestUUID string  `json:"request_uuid"`
            ToCity string  `json:"to_city"`
            Journey string  `json:"journey"`
            Action string  `json:"action"`
            StormbreakerUUID string  `json:"stormbreaker_uuid"`
            IsCodAvailable bool  `json:"is_cod_available"`
            Source string  `json:"source"`
            LocationDetails []TATLocationDetailsResponse  `json:"location_details"`
            PaymentMode string  `json:"payment_mode"`
            Error TATErrorSchemaResponse  `json:"error"`
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

        
            SubType string  `json:"sub_type"`
            Logistics LogisticsResponse  `json:"logistics"`
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            UID string  `json:"uid"`
            Meta CountryMetaResponse  `json:"meta"`
            ParentID string  `json:"parent_id"`
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
    
    // ReAssignStoreRequest ...
    type ReAssignStoreRequest struct {

        
            ToPincode string  `json:"to_pincode"`
            Articles []map[string]interface{}  `json:"articles"`
            Configuration map[string]interface{}  `json:"configuration"`
            IgnoredLocations []string  `json:"ignored_locations"`
            Identifier string  `json:"identifier"`
         
    }
    
    // ReAssignStoreResponse ...
    type ReAssignStoreResponse struct {

        
            Success bool  `json:"success"`
            Error map[string]interface{}  `json:"error"`
            ToPincode string  `json:"to_pincode"`
            Articles []map[string]interface{}  `json:"articles"`
         
    }
    
